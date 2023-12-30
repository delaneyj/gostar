package bxs

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = "2.1.4"
	hAttr          = "1em"
	viewbox        = "0 0 24 24"
	fill           = "currentColor"
)

type BxsIcon struct {
	*SVGSVGElement
}

func AddToQueue(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 22h12v-2H4V8H2v12c0 1.103.897 2 2 2"/><path fill="currentColor" d="M20 2H8c-1.103 0-2 .897-2 2v12c0 1.103.897 2 2 2h12c1.103 0 2-.897 2-2V4c0-1.103-.897-2-2-2m-2 9h-3v3h-2v-3h-3V9h3V6h2v3h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Adjust(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.579 2 2 6.58 2 12s4.579 10 10 10s10-4.58 10-10S17.421 2 12 2m0 17V5c3.829 0 7 3.169 7 7c0 3.828-3.171 7-7 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdjustAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.071 19.071c3.833-3.833 3.833-10.31 0-14.143s-10.31-3.833-14.143 0s-3.833 10.31 0 14.143s10.31 3.833 14.143 0M7.051 7.051c2.706-2.707 7.191-2.708 9.898 0l-9.898 9.898c-2.708-2.707-2.71-7.19 0-9.898"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Alarm(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 4c-4.878 0-9 4.122-9 9s4.122 9 9 9c4.879 0 9-4.122 9-9s-4.121-9-9-9m5 10h-6V8h2v4h4zm3.292-7.292l-3.01-3l1.412-1.417l3.01 3zM5.282 2.294L6.7 3.706l-2.99 3l-1.417-1.413z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlarmAdd(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 4c-4.879 0-9 4.121-9 9s4.121 9 9 9s9-4.121 9-9s-4.121-9-9-9m4 10h-3v3h-2v-3H8v-2h3V9h2v3h3zm1.284-10.293l1.412-1.416l3.01 3l-1.413 1.417zM5.282 2.294L6.7 3.706l-2.99 3l-1.417-1.413z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlarmExclamation(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m17.284 3.707l1.412-1.416l3.01 3l-1.413 1.417zm-10.586 0l-2.99 2.999L2.29 5.294l2.99-3zM12 4c-4.879 0-9 4.121-9 9s4.121 9 9 9s9-4.121 9-9s-4.121-9-9-9m1 14h-2v-2h2zm0-4h-2V8h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlarmOff(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.292 6.708l-3.01-3l1.412-1.417l3.01 3zm1.415 13.585l-2.287-2.288C20.409 16.563 21 14.837 21 13c0-4.878-4.121-9-9-9c-1.838 0-3.563.59-5.006 1.581L5.91 4.496l.788-.79l-1.416-1.412l-.786.788l-.789-.789l-1.414 1.414l18 18zM17 14h-1.586l-2-2H17zm-6-6h2v3.586l-2-2zm1 14c1.658 0 3.224-.485 4.574-1.305L4.305 8.426A8.794 8.794 0 0 0 3 13c0 4.878 4.122 9 9 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlarmSnooze(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m17.284 3.707l1.412-1.416l3.01 3l-1.413 1.417zm-10.586 0l-2.99 2.999L2.29 5.294l2.99-3zM12 4c-4.878 0-9 4.121-9 9s4.122 9 9 9c4.879 0 9-4.121 9-9s-4.121-9-9-9m4 13H8.131l4-6H8V9h7.868l-1.035 1.554l-.001.001L11.869 15H16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Album(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m-4.607 8.055A4.956 4.956 0 0 0 7 12H5a6.978 6.978 0 0 1 2.051-4.95a6.978 6.978 0 0 1 2.225-1.5l.779 1.842c-.596.252-1.13.612-1.59 1.072s-.82.995-1.072 1.591m4.6 3.945a2.007 2.007 0 1 1 0-4.014a2.007 2.007 0 0 1 0 4.014"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ambulance(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.86 12.48L19.15 8a2 2 0 0 0-1.72-1H15V5a1 1 0 0 0-1-1H4a2 2 0 0 0-2 2v10a2 2 0 0 0 1 1.73a3.49 3.49 0 0 0 7 .27h3.1a3.48 3.48 0 0 0 6.9 0a2 2 0 0 0 2-2v-3a1.07 1.07 0 0 0-.14-.52M6.5 19A1.5 1.5 0 1 1 8 17.5A1.5 1.5 0 0 1 6.5 19m5.5-8h-2v2H8v-2H6V9h2V7h2v2h2zm4.5 8a1.5 1.5 0 1 1 1.5-1.5a1.5 1.5 0 0 1-1.5 1.5M15 12V9h2.43l1.8 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Analyse(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.626 8.878a7.937 7.937 0 0 1 1.71-2.541a7.92 7.92 0 0 1 2.542-1.71a8.12 8.12 0 0 1 6.13-.041A2.49 2.49 0 0 0 17.5 7C18.886 7 20 5.886 20 4.5S18.886 2 17.5 2c-.689 0-1.312.276-1.763.725c-2.431-.973-5.223-.958-7.635.059c-1.19.5-2.26 1.22-3.18 2.139A9.98 9.98 0 0 0 2 12h2c0-1.086.211-2.136.626-3.122m14.747 6.244c-.401.952-.977 1.808-1.71 2.541s-1.589 1.309-2.542 1.71a8.12 8.12 0 0 1-6.13.041A2.488 2.488 0 0 0 6.5 17C5.114 17 4 18.114 4 19.5S5.114 22 6.5 22c.689 0 1.312-.276 1.763-.725A9.973 9.973 0 0 0 12 22a9.983 9.983 0 0 0 9.217-6.102A9.992 9.992 0 0 0 22 12h-2a7.993 7.993 0 0 1-.627 3.122"/><path fill="currentColor" d="M12 7.462c-2.502 0-4.538 2.036-4.538 4.538S9.498 16.538 12 16.538c2.502 0 4.538-2.036 4.538-4.538S14.502 7.462 12 7.462"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Angry(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m-5 8.5l.002-.022l-1.373-.549l.742-1.857l5 2l-.742 1.857l-1.031-.413c-.014.014-.023.031-.037.044A1.499 1.499 0 0 1 7 10.5M8 17s1-3 4-3s4 3 4 3zm8.986-6.507c0 .412-.167.785-.438 1.056a1.488 1.488 0 0 1-2.112 0c-.011-.011-.019-.024-.029-.035l-1.037.415l-.742-1.857l5-2l.742 1.857l-1.386.554a.036.036 0 0 1 .002.01"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Arch(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 8V6H3v14H2v2h6v-7c0-.163.046-4 4-4c3.821 0 3.993 3.602 4 4v7h6v-2h-1zM2 2h20v2H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Archive(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.704 5.29l-2.997-2.997A.996.996 0 0 0 18 2H6a.996.996 0 0 0-.707.293L2.296 5.29A.994.994 0 0 0 2 5.999V19a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V5.999a.994.994 0 0 0-.296-.709M6.414 4h11.172l1 1H5.414zM17 13v1H7v-4h2v2h6v-2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArchiveIn(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.706 5.292l-2.999-2.999A.996.996 0 0 0 18 2H6a.997.997 0 0 0-.707.293L2.294 5.292A.996.996 0 0 0 2 6v13c0 1.103.897 2 2 2h16c1.103 0 2-.897 2-2V6a.994.994 0 0 0-.294-.708M6.414 4h11.172l1 1H5.414zM12 18l-5-5h3v-3h4v3h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArchiveOut(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.706 5.292l-2.999-2.999A.996.996 0 0 0 18 2H6a.996.996 0 0 0-.707.293L2.294 5.292A.994.994 0 0 0 2 6v13c0 1.103.897 2 2 2h16c1.103 0 2-.897 2-2V6a.994.994 0 0 0-.294-.708M6.414 4h11.172l1 1H5.414zM14 14v3h-4v-3H7l5-5l5 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Area(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 19a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2zm9-13h6v6h-2V8h-4zm-6 6h2v4h4v2H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowFromBottom(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 18h12v2H6zm6-14l-6 6h5v6h2v-6h5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowFromLeft(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 6h2v12H4zm10 5H8v2h6v5l6-6l-6-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowFromRight(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 6h2v12h-2zm-8 12v-5h6v-2h-6V6l-6 6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowFromTop(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 4h12v2H6zm5 4v6H6l6 6l6-6h-5V8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowToBottom(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 18h12v2H6zm5-14v6H6l6 6l6-6h-5V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowToLeft(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 6h2v12H4zm10 7h6v-2h-6V6l-6 6l6 6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowToRight(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 6h2v12h-2zm-8 5H4v2h6v5l6-6l-6-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowToTop(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 4h12v2H6zm5 10v6h2v-6h5l-6-6l-6 6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Award(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 8.999a6.99 6.99 0 0 0 2.879 5.646l.001.001a6.972 6.972 0 0 0 1.881.979l.051.019a6.906 6.906 0 0 0 1.163.271a6.79 6.79 0 0 0 1.024.085H12c.35 0 .69-.034 1.027-.084l.182-.028c.336-.059.664-.139.981-.243l.042-.016C17 14.693 19 12.078 19 8.999C19 5.14 15.86 2 12 2S5 5.14 5 8.999M12 4c2.756 0 5 2.242 5 4.999h-2A3.003 3.003 0 0 0 12 6zM7.521 16.795V22L12 20.5l4.479 1.5l.001-5.205a8.932 8.932 0 0 1-8.959 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BabyCarriage(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.666 12.277a7.72 7.72 0 0 0 .171-.665c.003-.017.004-.033.008-.05c.02-.098.029-.199.045-.298c.025-.157.055-.313.07-.471a7.979 7.979 0 0 0-2.303-6.45A7.979 7.979 0 0 0 14 2v8H6.517l-.858-2H2v2h2.341l1.828 4.266A3.504 3.504 0 0 0 4 17.5C4 19.43 5.57 21 7.5 21c1.759 0 3.204-1.309 3.449-3h2.102c.245 1.691 1.69 3 3.449 3c1.93 0 3.5-1.57 3.5-3.5c0-.63-.181-1.213-.473-1.725c.042-.041.089-.077.131-.119c.36-.361.688-.759.977-1.184c.288-.43.536-.886.736-1.359c.016-.037.026-.076.041-.113h.001l.015-.042c.088-.22.168-.441.235-.668zM7.5 19c-.827 0-1.5-.673-1.5-1.5S6.673 16 7.5 16s1.5.673 1.5 1.5S8.327 19 7.5 19m9 0c-.827 0-1.5-.673-1.5-1.5s.673-1.5 1.5-1.5s1.5.673 1.5 1.5s-.673 1.5-1.5 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Backpack(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 15a1 1 0 0 0-1-1H9c-.551 0-1 .448-1 1v2h8zm-8 4h8v3H8z"/><path fill="currentColor" d="M21 12c0-2.967-2.167-5.432-5-5.91V5c0-1.654-1.346-3-3-3h-2C9.346 2 8 3.346 8 5v1.09C5.167 6.568 3 9.033 3 12v8c0 1.103.897 2 2 2h1v-7c0-1.654 1.346-3 3-3h6c1.654 0 3 1.346 3 3v7h1c1.103 0 2-.897 2-2zM10 5c0-.552.449-1 1-1h2a1 1 0 0 1 1 1v1h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Badge(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.965 8.521C19.988 8.347 20 8.173 20 8c0-2.379-2.143-4.288-4.521-3.965C14.786 2.802 13.466 2 12 2s-2.786.802-3.479 2.035C6.138 3.712 4 5.621 4 8c0 .173.012.347.035.521C2.802 9.215 2 10.535 2 12s.802 2.785 2.035 3.479A3.976 3.976 0 0 0 4 16c0 2.379 2.138 4.283 4.521 3.965C9.214 21.198 10.534 22 12 22s2.786-.802 3.479-2.035C17.857 20.283 20 18.379 20 16c0-.173-.012-.347-.035-.521C21.198 14.785 22 13.465 22 12s-.802-2.785-2.035-3.479"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BadgeCheck(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.965 8.521C19.988 8.347 20 8.173 20 8c0-2.379-2.143-4.288-4.521-3.965C14.786 2.802 13.466 2 12 2s-2.786.802-3.479 2.035C6.138 3.712 4 5.621 4 8c0 .173.012.347.035.521C2.802 9.215 2 10.535 2 12s.802 2.785 2.035 3.479A3.976 3.976 0 0 0 4 16c0 2.379 2.138 4.283 4.521 3.965C9.214 21.198 10.534 22 12 22s2.786-.802 3.479-2.035C17.857 20.283 20 18.379 20 16c0-.173-.012-.347-.035-.521C21.198 14.785 22 13.465 22 12s-.802-2.785-2.035-3.479m-9.01 7.895l-3.667-3.714l1.424-1.404l2.257 2.286l4.327-4.294l1.408 1.42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BadgeDollar(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.047 14.668a.994.994 0 0 0 .465.607l1.91 1.104v2.199a1 1 0 0 0 1 1h2.199l1.104 1.91a1.01 1.01 0 0 0 .866.5c.174 0 .347-.046.501-.135L12 20.75l1.91 1.104a1.001 1.001 0 0 0 1.366-.365l1.103-1.91h2.199a1 1 0 0 0 1-1V16.38l1.91-1.104a1 1 0 0 0 .365-1.367L20.75 12l1.104-1.908a1 1 0 0 0-.365-1.366l-1.91-1.104v-2.2a1 1 0 0 0-1-1H16.38l-1.103-1.909a1.008 1.008 0 0 0-.607-.466a.993.993 0 0 0-.759.1L12 3.25l-1.909-1.104a1 1 0 0 0-1.366.365l-1.104 1.91H5.422a1 1 0 0 0-1 1V7.62l-1.91 1.104a1.003 1.003 0 0 0-.365 1.368L3.251 12l-1.104 1.908a1.009 1.009 0 0 0-.1.76M12 13c-3.48 0-4-1.879-4-3c0-1.287 1.029-2.583 3-2.915V6.012h2v1.109c1.734.41 2.4 1.853 2.4 2.879h-1l-1 .018C13.386 9.638 13.185 9 12 9c-1.299 0-2 .515-2 1c0 .374 0 1 2 1c3.48 0 4 1.879 4 3c0 1.287-1.029 2.583-3 2.915V18h-2v-1.08c-2.339-.367-3-2.003-3-2.92h2c.011.143.159 1 2 1c1.38 0 2-.585 2-1c0-.325 0-1-2-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Baguette(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m11.13 4.41l4.23 4.23L14.3 9.7l-4.24-4.24l-1.77 1.77l4.24 4.24l-1.06 1.06l-4.24-4.24l-1.77 1.77L9.7 14.3l-1.06 1.06l-4.23-4.23C1.86 14 1.55 18 3.79 20.21a5.38 5.38 0 0 0 3.85 1.5a8 8 0 0 0 5.6-2.47l6-6c2.87-2.87 3.31-7.11 1-9.45s-6.24-1.93-9.11.62"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ball(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.247 3.034c-.069-.018-1.742-.433-4.052-.433c-2.842 0-6.868.64-9.91 3.687c-5.34 5.349-3.34 13.61-3.252 13.96a1 1 0 0 0 .728.726c.069.018 1.726.426 4.018.426c2.849 0 6.884-.641 9.932-3.688c5.335-5.335 3.351-13.6 3.264-13.949a1.005 1.005 0 0 0-.728-.729m-3.537 9.262l-1.414 1.414l-1.793-1.792l-1.586 1.586l1.792 1.793l-1.414 1.414l-1.792-1.793l-1.793 1.793l-1.414-1.414l1.793-1.793l-1.793-1.794l1.414-1.414l1.793 1.794l1.586-1.586l-1.794-1.793l1.414-1.414l1.794 1.793l1.793-1.793l1.414 1.414l-1.793 1.793z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Balloon(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 19.91L10 22h4l-1-2.09c4-.65 7-5.28 7-9.91a8 8 0 0 0-16 0c0 4.63 3.08 9.26 7 9.91m1-15.66v1.5A4.26 4.26 0 0 0 7.75 10h-1.5A5.76 5.76 0 0 1 12 4.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BandAid(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m11.646 3.868l-7.778 7.778a6.007 6.007 0 0 0 0 8.485a5.984 5.984 0 0 0 4.242 1.754a5.984 5.984 0 0 0 4.243-1.754l7.778-7.778a6.007 6.007 0 0 0 0-8.485a6.008 6.008 0 0 0-8.485 0M9 13a1 1 0 1 1 0-2a1 1 0 0 1 0 2m3 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2m0-6a1 1 0 1 1 0-2a1 1 0 0 1 0 2m3 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bank(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 8v4.001h1V18H2v3h16l3 .001V21h1v-3h-1v-5.999h1V8L12 2zm4 10v-5.999h2V18zm5 0v-5.999h2V18zm7 0h-2v-5.999h2zM14 8a2 2 0 1 1-4.001-.001A2 2 0 0 1 14 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarChartAltTwo(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 21H3a1 1 0 0 1-1-1v-8a1 1 0 0 1 1-1h3a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1m7 0h-3a1 1 0 0 1-1-1V3a1 1 0 0 1 1-1h3a1 1 0 0 1 1 1v17a1 1 0 0 1-1 1m7 0h-3a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1h3a1 1 0 0 1 1 1v11a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarChartSquare(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 19V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2M9 18H7v-6h2zm4 0h-2V7h2zm4 0h-2v-8h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Barcode(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 4H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2M8 17H5V7h3zm2 0H9V7h1zm2 0h-1V7h1zm4 0h-3V7h3zm3 0h-2V7h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Baseball(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.984 12.236a9.966 9.966 0 0 0-2.913-7.308a9.966 9.966 0 0 0-7.308-2.913a9.04 9.04 0 0 1-.673 4.313l-1.84-.78a7.044 7.044 0 0 0 .526-3.284a9.927 9.927 0 0 0-4.847 2.665a9.924 9.924 0 0 0-2.666 4.852a7.082 7.082 0 0 0 2.576-.276l.575 1.916c-1.1.33-2.257.443-3.398.344a9.964 9.964 0 0 0 2.913 7.307a9.965 9.965 0 0 0 7.307 2.913a9.079 9.079 0 0 1 .344-3.398l1.916.575a7.06 7.06 0 0 0-.276 2.576a9.927 9.927 0 0 0 4.853-2.666a9.926 9.926 0 0 0 2.665-4.848a7.056 7.056 0 0 0-3.284.526l-.78-1.841a9.025 9.025 0 0 1 4.31-.673M9.17 9.173a9.017 9.017 0 0 1-2.192 1.612l-.927-1.772a7.01 7.01 0 0 0 2.576-2.314l1.663 1.113c-.328.49-.705.948-1.12 1.361m7.074 7.068a6.991 6.991 0 0 0-1.257 1.708l-1.772-.927a9.025 9.025 0 0 1 2.972-3.312l1.113 1.663a6.987 6.987 0 0 0-1.056.868"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Basket(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.132 2.504L4.42 9H3a1.001 1.001 0 0 0-.965 1.263l2.799 10.263A2.004 2.004 0 0 0 6.764 22h10.473c.898 0 1.692-.605 1.93-1.475l2.799-10.263A.998.998 0 0 0 21 9h-1.42l-3.712-6.496l-1.736.992L17.277 9H6.723l3.145-5.504zM14 13h2v5h-2zm-6 0h2v5H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Basketball(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.328 4.258a9.953 9.953 0 0 0-5.949-2.235a8.99 8.99 0 0 1-1.835 7.107L12 10.586zM7.701 9.115L4.258 5.672a9.938 9.938 0 0 0-2.112 4.704a7.007 7.007 0 0 0 5.555-1.261m12.041-3.443L13.414 12l1.456 1.456a8.993 8.993 0 0 1 7.107-1.835a9.953 9.953 0 0 0-2.235-5.949m2.112 7.952a7.007 7.007 0 0 0-5.555 1.261l3.443 3.443a9.924 9.924 0 0 0 2.112-4.704M9.115 7.701a7.007 7.007 0 0 0 1.261-5.555a9.928 9.928 0 0 0-4.704 2.112zm4.509 14.153a9.936 9.936 0 0 0 4.704-2.111L14.885 16.3a7.003 7.003 0 0 0-1.261 5.554M12 13.414l-6.328 6.328a9.953 9.953 0 0 0 5.949 2.235a8.99 8.99 0 0 1 1.835-7.107zm-7.742 4.914L10.586 12L9.13 10.544a8.993 8.993 0 0 1-7.107 1.835a9.953 9.953 0 0 0 2.235 5.949"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bath(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 10H7V7.113c0-.997.678-1.923 1.661-2.085A2.003 2.003 0 0 1 11 7h2a4.003 4.003 0 0 0-4.398-3.98C6.523 3.222 5 5.089 5 7.178V10H3a1 1 0 0 0-1 1v2c0 2.606 1.674 4.823 4 5.65V22h2v-3h8v3h2v-3.35c2.326-.827 4-3.044 4-5.65v-2a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Battery(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 8a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-2h2v-4h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryCharging(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 10V8a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-2h2v-4zM9 17l2-3.89L7 12l6-5l-1 3.89L16 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryFull(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 18h14a2 2 0 0 0 2-2v-2h2v-4h-2V8a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2m1-9h12v6H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryLow(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 8a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-2h2v-4h-2zM5 15V9h3l4 6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bed(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 9.556V3h-2v2H6V3H4v6.557C2.81 10.25 2 11.526 2 13v4a1 1 0 0 0 1 1h1v4h2v-4h12v4h2v-4h1a1 1 0 0 0 1-1v-4c0-1.474-.811-2.75-2-3.444M11 9H6V7h5zm7 0h-5V7h5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BeenHere(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C7.589 2 4 5.589 4 9.995C3.971 16.44 11.696 21.784 12 22c0 0 8.029-5.56 8-12c0-4.411-3.589-8-8-8m-1 13.414l-3.707-3.707l1.414-1.414L11 12.586l5.293-5.293l1.414 1.414z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Beer(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 6h-2V4a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v15c0 1.654 1.346 3 3 3h10c1.654 0 3-1.346 3-3v-1h2c1.103 0 2-.897 2-2V8c0-1.103-.897-2-2-2M8 17H6V7h2zm6 0h-2V7h2zm6-1h-2V8h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bell(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22a2.98 2.98 0 0 0 2.818-2H9.182A2.98 2.98 0 0 0 12 22m7-7.414V10c0-3.217-2.185-5.927-5.145-6.742C13.562 2.52 12.846 2 12 2s-1.562.52-1.855 1.258C7.185 4.074 5 6.783 5 10v4.586l-1.707 1.707A.996.996 0 0 0 3 17v1a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1v-1a.996.996 0 0 0-.293-.707z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellMinus(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22a2.98 2.98 0 0 0 2.818-2H9.182A2.98 2.98 0 0 0 12 22m8.707-5.707L19 14.586V10c0-3.217-2.185-5.926-5.145-6.743C13.562 2.52 12.846 2 12 2s-1.562.52-1.855 1.258C7.185 4.074 5 6.783 5 10v4.586l-1.707 1.707A.997.997 0 0 0 3 17v1a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1v-1a.997.997 0 0 0-.293-.707M16 12H8v-2h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellOff(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.345 18.931A.993.993 0 0 0 21 18v-1a.996.996 0 0 0-.293-.707L19 14.586V10c0-3.217-2.185-5.927-5.145-6.742C13.562 2.52 12.846 2 12 2s-1.562.52-1.855 1.258c-1.323.364-2.463 1.128-3.346 2.127L3.707 2.293L2.293 3.707l18 18l1.414-1.414zM12 22a2.98 2.98 0 0 0 2.818-2H9.182A2.98 2.98 0 0 0 12 22M5 10v4.586l-1.707 1.707A.996.996 0 0 0 3 17v1a1 1 0 0 0 1 1h10.879L5.068 9.189C5.037 9.457 5 9.724 5 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellPlus(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22a2.98 2.98 0 0 0 2.818-2H9.182A2.98 2.98 0 0 0 12 22m8.707-5.707L19 14.586V10c0-3.217-2.185-5.926-5.145-6.743C13.562 2.52 12.846 2 12 2s-1.562.52-1.855 1.258C7.185 4.074 5 6.783 5 10v4.586l-1.707 1.707A.997.997 0 0 0 3 17v1a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1v-1a.997.997 0 0 0-.293-.707M16 12h-3v3h-2v-3H8v-2h3V7h2v3h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellRing(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.705 3.71l-1.41-1.42C1 5.563 1 7.935 1 11h1l1-.063C3 8.009 3 6.396 5.705 3.71m13.999-1.42l-1.408 1.42C21 6.396 21 8.009 21 11l2-.063c0-3.002 0-5.374-3.296-8.647M12 22a2.98 2.98 0 0 0 2.818-2H9.182A2.98 2.98 0 0 0 12 22m7-7.414V10c0-3.217-2.185-5.927-5.145-6.742C13.562 2.52 12.846 2 12 2s-1.562.52-1.855 1.258C7.184 4.073 5 6.783 5 10v4.586l-1.707 1.707A.996.996 0 0 0 3 17v1a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1v-1a.996.996 0 0 0-.293-.707z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bible(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 22h15v-2H6.012C5.55 19.988 5 19.805 5 19s.55-.988 1.012-1H21V4a2 2 0 0 0-2-2H6c-1.206 0-3 .799-3 3v14c0 2.201 1.794 3 3 3M8 7h3V5h2v2h3v2h-3v6h-2V9H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Binoculars(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.447 5.345A3.27 3.27 0 0 0 16.29 3a3.293 3.293 0 0 0-3.277 3h-2.025a3.297 3.297 0 0 0-3.284-3a3.268 3.268 0 0 0-3.151 2.345l-2.511 8.368A1.027 1.027 0 0 0 2 14v1a5.006 5.006 0 0 0 5.001 5a5.003 5.003 0 0 0 4.576-3h.846a5.003 5.003 0 0 0 4.576 3A5.006 5.006 0 0 0 22 14.999V14c0-.098-.015-.194-.042-.287zM7.001 18A3.005 3.005 0 0 1 4 15c0-.076.017-.147.022-.222A2.995 2.995 0 0 1 7 12a3 3 0 0 1 3 3v.009A3.004 3.004 0 0 1 7.001 18m9.998 0A3.004 3.004 0 0 1 14 15.009V15a3 3 0 0 1 6-.001A3.005 3.005 0 0 1 16.999 18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Blanket(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 2H7C4.243 2 2 4.243 2 7v10c0 2.757 2.243 5 5 5h12c1.654 0 3-1.346 3-3s-1.346-3-3-3H6v2h13a1 1 0 0 1 0 2H7c-1.654 0-3-1.346-3-3s1.346-3 3-3h13c1.103 0 2-.897 2-2V4c0-1.103-.897-2-2-2M4 13h.003L4 13.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bolt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.168 8H13l.806-4.835A1 1 0 0 0 12.819 2H7.667a1 1 0 0 0-.986.835l-1.667 10A1 1 0 0 0 6 14h4v8l8.01-12.459A1 1 0 0 0 17.168 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoltCircle(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m-1 16v-5H7l6-7v5h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bomb(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m18.293 4.293l-1.086 1.086l-1.086-1.086a.999.999 0 0 0-1.414 0l-1.249 1.249A8.427 8.427 0 0 0 10.499 5C5.813 5 2 8.813 2 13.5S5.813 22 10.499 22s8.5-3.813 8.5-8.5a8.42 8.42 0 0 0-.431-2.654L19.914 9.5a.999.999 0 0 0 0-1.414l-1.293-1.293l1.09-1.09C19.94 5.474 20.556 5 21 5h1V3h-1c-1.4 0-2.584 1.167-2.707 1.293M10.499 10c-.935 0-1.813.364-2.475 1.025A3.48 3.48 0 0 0 7 13.5H5c0-1.468.571-2.849 1.609-3.888A5.464 5.464 0 0 1 10.499 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bone(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.852 5.148a3.317 3.317 0 0 0-.96-2.183a3.333 3.333 0 1 0-4.713 4.714l-5.499 5.5a3.333 3.333 0 1 0-4.714 4.713c.606.606 1.39.918 2.183.96c.042.793.354 1.576.96 2.183a3.333 3.333 0 1 0 4.713-4.714l5.499-5.499a3.333 3.333 0 1 0 4.714-4.713a3.313 3.313 0 0 0-2.183-.961"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bong(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m19.293 8.293l-2.069 2.069A7.017 7.017 0 0 0 15 8.681V4h1V2H8v2h1v4.681A7.01 7.01 0 0 0 5 15c0 3.859 3.141 7 7 7s7-3.141 7-7a6.958 6.958 0 0 0-.652-2.934l2.359-2.359zm-8.959 1.998l.666-.235V4h2v6.056l.666.235A5.006 5.006 0 0 1 16.886 14H7.114a5.006 5.006 0 0 1 3.22-3.709"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Book(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.012 18H21V4a2 2 0 0 0-2-2H6c-1.206 0-3 .799-3 3v14c0 2.201 1.794 3 3 3h15v-2H6.012C5.55 19.988 5 19.805 5 19s.55-.988 1.012-1M8 6h9v2H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookAdd(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.012 18H21V4c0-1.103-.897-2-2-2H6c-1.206 0-3 .799-3 3v14c0 2.201 1.794 3 3 3h15v-2H6.012C5.55 19.988 5 19.806 5 19s.55-.988 1.012-1M8 9h3V6h2v3h3v2h-3v3h-2v-3H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 5v14c0 2.201 1.794 3 3 3h15v-2H6.012C5.55 19.988 5 19.806 5 19s.55-.988 1.012-1H21V4c0-1.103-.897-2-2-2H6c-1.206 0-3 .799-3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookBookmark(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 2H6c-1.206 0-3 .799-3 3v14c0 2.201 1.794 3 3 3h15v-2H6.012C5.55 19.988 5 19.806 5 19c0-.101.009-.191.024-.273c.112-.576.584-.717.988-.727H21V4a2 2 0 0 0-2-2m0 9l-2-1l-2 1V4h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookContent(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 3H4c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2h16c1.103 0 2-.897 2-2V5c0-1.103-.897-2-2-2m-1 4v2h-5V7zm-5 4h5v2h-5zM4 19V5h7v14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookHeart(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.012 18H21V4c0-1.103-.897-2-2-2H6c-1.206 0-3 .799-3 3v14c0 2.201 1.794 3 3 3h15v-2H6.012C5.55 19.988 5 19.806 5 19c0-.101.009-.191.024-.273c.112-.576.584-.717.988-.727M8.648 7.642a2.224 2.224 0 0 1 3.125 0l.224.219l.223-.219a2.225 2.225 0 0 1 3.126 0a2.129 2.129 0 0 1 0 3.069L11.998 14l-3.349-3.289a2.128 2.128 0 0 1-.001-3.069"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookOpen(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 3h-7a2.98 2.98 0 0 0-2 .78A2.98 2.98 0 0 0 10 3H3a1 1 0 0 0-1 1v15a1 1 0 0 0 1 1h5.758a2.01 2.01 0 0 1 1.414.586l1.121 1.121c.009.009.021.012.03.021c.086.08.182.15.294.196h.002a.996.996 0 0 0 .762 0h.002c.112-.046.208-.117.294-.196c.009-.009.021-.012.03-.021l1.121-1.121A2.01 2.01 0 0 1 15.242 20H21a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1m-1 15h-4.758a4.03 4.03 0 0 0-2.242.689V6c0-.551.448-1 1-1h6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookReader(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 8v11.529S6.621 19.357 12 22c5.379-2.643 10-2.471 10-2.471V8s-5.454 0-10 2.471C7.454 8 2 8 2 8"/><circle cx="12" cy="5" r="3" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookmark(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 10.132v-6c0-1.103-.897-2-2-2H7c-1.103 0-2 .897-2 2V22l7-4.666L19 22z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.5 2h-12C4.57 2 3 3.57 3 5.5V22l7-3.5l7 3.5v-9h5V5.5C22 3.57 20.43 2 18.5 2m1.5 9h-3V5.5c0-.827.673-1.5 1.5-1.5s1.5.673 1.5 1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkAltMinus(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.5 2h-12C4.57 2 3 3.57 3 5.5V22l7-3.5l7 3.5v-9h5V5.5C22 3.57 20.43 2 18.5 2M13 11H7V9h6zm7 0h-3V5.5c0-.827.673-1.5 1.5-1.5s1.5.673 1.5 1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkAltPlus(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.5 2h-12C4.57 2 3 3.57 3 5.5V22l7-3.5l7 3.5v-9h5V5.5C22 3.57 20.43 2 18.5 2M13 11h-2v2H9v-2H7V9h2V7h2v2h2zm7 0h-3V5.5c0-.827.673-1.5 1.5-1.5s1.5.673 1.5 1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkHeart(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 22V4c0-1.103-.897-2-2-2H7c-1.103 0-2 .897-2 2v18l7-4.666zM8.006 8.056c0-.568.224-1.083.585-1.456c.361-.372.86-.603 1.412-.603c0 0 .996-.003 1.997 1.029c1.001-1.032 1.997-1.029 1.997-1.029c.552 0 1.051.23 1.412.603s.585.888.585 1.456s-.224 1.084-.585 1.456L12 13.203L8.591 9.512a2.083 2.083 0 0 1-.585-1.456"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkMinus(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2H7a2 2 0 0 0-2 2v18l7-4.848L19 22V4a2 2 0 0 0-2-2m-1 9H8V9h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkPlus(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2H7a2 2 0 0 0-2 2v18l7-4.848L19 22V4a2 2 0 0 0-2-2m-1 9h-3v3h-2v-3H8V9h3V6h2v3h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkStar(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 11.222L14.667 13l-.89-3.111L16 8l-2.667-.333L12 5l-1.333 2.667L8 8l2.223 1.889L9.333 13z"/><path fill="currentColor" d="M19 21.723V4a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v17.723l7-4.571zM8 8l2.667-.333L12 5l1.333 2.667L16 8l-2.223 1.889l.89 3.111L12 11.222L9.333 13l.89-3.111z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookmarks(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.999 23V7c0-1.103-.897-2-2-2h-8c-1.103 0-2 .897-2 2v16l6-3.601z"/><path fill="currentColor" d="M15.585 3h1.414c1.103 0 2 .897 2 2v10.443l2 2.489V3c0-1.103-.897-2-2-2h-8c-1.103 0-2 .897-2 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bot(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 10.975V8a2 2 0 0 0-2-2h-6V4.688c.305-.274.5-.668.5-1.11a1.5 1.5 0 0 0-3 0c0 .442.195.836.5 1.11V6H5a2 2 0 0 0-2 2v2.998l-.072.005A.999.999 0 0 0 2 12v2a1 1 0 0 0 1 1v5a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-5a1 1 0 0 0 1-1v-1.938a1.004 1.004 0 0 0-.072-.455c-.202-.488-.635-.605-.928-.632M7 12c0-1.104.672-2 1.5-2s1.5.896 1.5 2s-.672 2-1.5 2S7 13.104 7 12m8.998 6c-1.001-.003-7.997 0-7.998 0v-2s7.001-.002 8.002 0zm-.498-4c-.828 0-1.5-.896-1.5-2s.672-2 1.5-2s1.5.896 1.5 2s-.672 2-1.5 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BowlHot(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 10H3a1 1 0 0 0-1 1a10 10 0 0 0 5 8.66V21a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1v-1.34A10 10 0 0 0 22 11a1 1 0 0 0-1-1M9 9V7.93a4.51 4.51 0 0 0-1.28-3.15A2.49 2.49 0 0 1 7 3V2H5v1a4.51 4.51 0 0 0 1.28 3.17A2.49 2.49 0 0 1 7 7.93V9zm4 0V7.93a4.51 4.51 0 0 0-1.28-3.15A2.49 2.49 0 0 1 11 3V2H9v1a4.51 4.51 0 0 0 1.28 3.15A2.49 2.49 0 0 1 11 7.93V9zm4 0V7.93a4.51 4.51 0 0 0-1.28-3.15A2.49 2.49 0 0 1 15 3V2h-2v1a4.51 4.51 0 0 0 1.28 3.15A2.49 2.49 0 0 1 15 7.93V9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BowlRice(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 10a3.58 3.58 0 0 0-1.8-3a3.66 3.66 0 0 0-3.63-3.13a3.86 3.86 0 0 0-1 .13a3.7 3.7 0 0 0-5.11 0a3.86 3.86 0 0 0-1-.13A3.66 3.66 0 0 0 4.81 7A3.58 3.58 0 0 0 3 10a1 1 0 0 0-1 1a10 10 0 0 0 5 8.66V21a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1v-1.34A10 10 0 0 0 22 11a1 1 0 0 0-1-1M5 10a1.59 1.59 0 0 1 1.11-1.39l.83-.26l-.16-.85a1.64 1.64 0 0 1 1.66-1.62a1.78 1.78 0 0 1 .83.2l.81.45l.5-.77a1.71 1.71 0 0 1 2.84 0l.5.77l.81-.45a1.78 1.78 0 0 1 .83-.2a1.65 1.65 0 0 1 1.67 1.6l-.16.85l.82.28A1.59 1.59 0 0 1 19 10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BowlingBall(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2M6.5 12a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 6.5 12M9 6.5a1.5 1.5 0 1 1 3.001.001A1.5 1.5 0 0 1 9 6.5m2.5 6.5a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 11.5 13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Box(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 3h20v4H2zm17 5H3v11a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V8zm-3 6H8v-2h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Brain(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.299 17.596c.432 1.332 1.745 2.182 3.146 2.182H6.5A2.78 2.78 0 0 0 9.223 22c.457 0 .884-.115 1.262-.313a.992.992 0 0 0 .515-.882V3.027a.997.997 0 0 0-.785-.983a2.324 2.324 0 0 0-1.479.201c-.744.356-1.18 1.151-1.18 1.978v.055a2.778 2.778 0 0 0-2.744 4.433A3.327 3.327 0 0 0 2 12c0 1.178.611 2.211 1.533 2.812c-.43.771-.571 1.746-.234 2.784m15.889-8.885a2.778 2.778 0 0 0-2.744-4.433v-.055c0-.826-.437-1.622-1.181-1.978a2.32 2.32 0 0 0-1.478-.201a.998.998 0 0 0-.785.983v17.777c0 .365.192.712.516.882c.378.199.804.314 1.261.314a2.78 2.78 0 0 0 2.723-2.223h.056c1.4 0 2.714-.85 3.146-2.182c.337-1.038.196-2.013-.234-2.784A3.35 3.35 0 0 0 22 12a3.327 3.327 0 0 0-2.812-3.289"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Briefcase(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 6h-3V4c0-1.103-.897-2-2-2H9c-1.103 0-2 .897-2 2v2H4c-1.103 0-2 .897-2 2v3h20V8c0-1.103-.897-2-2-2M9 4h6v2H9zm5 10h-4v-2H2v7c0 1.103.897 2 2 2h16c1.103 0 2-.897 2-2v-7h-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BriefcaseAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 6h-3V4c0-1.103-.897-2-2-2H9c-1.103 0-2 .897-2 2v2H4c-1.103 0-2 .897-2 2v11c0 1.103.897 2 2 2h16c1.103 0 2-.897 2-2V8c0-1.103-.897-2-2-2M6 8h2v11H6zm12 11h-2V8h2zM15 4v2H9V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BriefcaseAltTwo(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 6h-3V4c0-1.103-.897-2-2-2H9c-1.103 0-2 .897-2 2v2H4c-1.103 0-2 .897-2 2v4h5v-2h2v2h6v-2h2v2h5V8c0-1.103-.897-2-2-2M9 4h6v2H9zm8 11h-2v-2H9v2H7v-2H2v6c0 1.103.897 2 2 2h16c1.103 0 2-.897 2-2v-6h-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Brightness(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.546 11.646L19 9.101V5.5a.5.5 0 0 0-.5-.5h-3.601l-2.546-2.546a.5.5 0 0 0-.707 0L9.101 5H5.5a.5.5 0 0 0-.5.5v3.601l-2.546 2.546a.5.5 0 0 0 0 .707L5 14.899V18.5a.5.5 0 0 0 .5.5h3.601l2.546 2.546a.5.5 0 0 0 .707 0L14.899 19H18.5a.5.5 0 0 0 .5-.5v-3.601l2.546-2.546a.5.5 0 0 0 0-.707M12 16a4 4 0 1 1 0-8a4 4 0 0 1 0 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrightnessHalf(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.546 11.646L19 9.101V5.5a.5.5 0 0 0-.5-.5h-3.601l-2.546-2.546a.5.5 0 0 0-.707 0L9.101 5H5.5a.5.5 0 0 0-.5.5v3.601l-2.546 2.546a.5.5 0 0 0 0 .707L5 14.899V18.5a.5.5 0 0 0 .5.5h3.601l2.546 2.546a.5.5 0 0 0 .707 0L14.899 19H18.5a.5.5 0 0 0 .5-.5v-3.601l2.546-2.546a.5.5 0 0 0 0-.707M12 8a4 4 0 0 1 0 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Brush(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.207 11.278l-2.035-2.035l-1.415-1.415l-5.035-5.035a.999.999 0 0 0-1.414 0L6.151 7.949L4.736 9.363a2.985 2.985 0 0 0-.878 2.122c0 .802.313 1.556.879 2.121l.707.707l-2.122 2.122a2.925 2.925 0 0 0-.873 2.108a2.968 2.968 0 0 0 1.063 2.308a2.92 2.92 0 0 0 1.886.681c.834 0 1.654-.341 2.25-.937l2.039-2.039l.707.706c1.133 1.133 3.107 1.134 4.242.001l.708-.707l.569-.569l.138-.138l5.156-5.157a.999.999 0 0 0 0-1.414m-7.277 5.865l-.708.706a1.021 1.021 0 0 1-1.414 0l-1.414-1.413a.999.999 0 0 0-1.414 0l-2.746 2.745a1.192 1.192 0 0 1-.836.352a.914.914 0 0 1-.595-.208a.981.981 0 0 1-.354-.782a.955.955 0 0 1 .287-.692l2.829-2.829a.999.999 0 0 0 0-1.414l-1.414-1.415c-.189-.188-.293-.438-.293-.706s.104-.519.293-.708l.707-.707l3.536 3.536z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrushAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 8.001h-2V8h-1V4.999a2.92 2.92 0 0 0-.874-2.108a2.943 2.943 0 0 0-2.39-.879C10.202 2.144 9 3.508 9 5.117V8H6c-1.103 0-2 .897-2 2v11a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V10.001c0-1.103-.897-2-2-2M6 12v-2h5V5.117c0-.57.407-1.07 1.002-1.117c.266 0 .512.103.712.307a.956.956 0 0 1 .286.692V10h.995l.005.001h4V12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bug(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.787 7h10.426c-.108-.158-.201-.331-.318-.481l2.813-2.812l-1.414-1.414l-2.846 2.846a6.575 6.575 0 0 0-.723-.454a5.778 5.778 0 0 0-5.45 0c-.25.132-.488.287-.722.453L5.707 2.293L4.293 3.707l2.813 2.812c-.118.151-.21.323-.319.481M5.756 9H2v2h2.307c-.065.495-.107.997-.107 1.5c0 .507.042 1.013.107 1.511H2v2h2.753c.013.039.021.08.034.118c.188.555.421 1.093.695 1.6c.044.081.095.155.141.234l-2.33 2.33l1.414 1.414l2.11-2.111a7.477 7.477 0 0 0 2.068 1.619c.479.253.982.449 1.496.58c.204.052.411.085.618.118V16h2v5.914a6.23 6.23 0 0 0 .618-.118a6.812 6.812 0 0 0 1.496-.58c.465-.246.914-.55 1.333-.904c.258-.218.5-.462.734-.716l2.111 2.111l1.414-1.414l-2.33-2.33c.047-.08.098-.155.142-.236c.273-.505.507-1.043.694-1.599c.013-.039.021-.079.034-.118H22v-2h-2.308c.065-.499.107-1.004.107-1.511c0-.503-.042-1.005-.106-1.5H22V9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BugAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m18 5.414l1.707-1.707l-1.414-1.414l-1.563 1.562C15.483 2.708 13.824 2 12 2s-3.483.708-4.73 1.855L5.707 2.293L4.293 3.707L6 5.414A6.937 6.937 0 0 0 5 9H3v2h2v2H3v2h2c0 3.86 3.141 7 7 7s7-3.14 7-7h2v-2h-2v-2h2V9h-2a6.937 6.937 0 0 0-1-3.586M15 15H9v-2h6zm0-4H9V9h6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Building(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2H7a2 2 0 0 0-2 2v17a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V4a2 2 0 0 0-2-2m-6 14H8v-2h3zm0-4H8v-2h3zm0-4H8V6h3zm5 8h-3v-2h3zm0-4h-3v-2h3zm0-4h-3V6h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuildingHouse(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.991 2H9.01C7.899 2 7 2.899 7 4.01v5.637l-4.702 4.642A1 1 0 0 0 3 16v5a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V4.009C21 2.899 20.102 2 18.991 2m-8.069 13.111V20H5v-5.568l2.987-2.949l2.935 3.003zM13 9h-2V7h2zm4 8h-2v-2h2zm0-4h-2v-2h2zm0-4h-2V7h2z"/><path fill="currentColor" d="M7 15h2v2H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Buildings(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 14.001h2v2H7z"/><path fill="currentColor" d="M19 2h-8a2 2 0 0 0-2 2v6H5c-1.103 0-2 .897-2 2v9a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V4a2 2 0 0 0-2-2M5 20v-8h6v8zm9-12h-2V6h2zm4 8h-2v-2h2zm0-4h-2v-2h2zm0-4h-2V6h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bulb(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 20h6v2H9zm7.906-6.288C17.936 12.506 19 11.259 19 9c0-3.859-3.141-7-7-7S5 5.141 5 9c0 2.285 1.067 3.528 2.101 4.73c.358.418.729.851 1.084 1.349c.144.206.38.996.591 1.921h-.792v2h8.032v-2h-.79c.213-.927.45-1.719.593-1.925c.352-.503.726-.94 1.087-1.363"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bullseye(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m0 17c-3.859 0-7-3.14-7-7s3.141-7 7-7s7 3.14 7 7s-3.141 7-7 7"/><path fill="currentColor" d="M12 7c-2.757 0-5 2.243-5 5s2.243 5 5 5s5-2.243 5-5s-2.243-5-5-5m0 7c-1.103 0-2-.897-2-2s.897-2 2-2s2 .897 2 2s-.897 2-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Buoy(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m7.736 8h-3.16A5.02 5.02 0 0 0 14 7.424V4.263A8.015 8.015 0 0 1 19.736 10M12 15c-1.654 0-3-1.346-3-3s1.346-3 3-3s3 1.346 3 3s-1.346 3-3 3M10 4.263v3.161A5.02 5.02 0 0 0 7.424 10h-3.16A8.015 8.015 0 0 1 10 4.263M4.264 14h3.16A5.02 5.02 0 0 0 10 16.576v3.161A8.015 8.015 0 0 1 4.264 14M14 19.737v-3.161A5.02 5.02 0 0 0 16.576 14h3.16A8.015 8.015 0 0 1 14 19.737"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bus(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 6.021c.003-.146-.007-1.465-1.3-2.735C18.427 2.036 17.143 2 17 2H6.996c-.239 0-1.493.063-2.708 1.302C3.036 4.578 3 5.859 3 6v3H2v3h1v6c0 .734.406 1.373 1 1.721V21a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-1h10v1a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-1.277A1.99 1.99 0 0 0 21 18v-6h1V9h-1zM9 4h6v2H9zM6.5 18a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 6.5 18m4.5-5H5V8h6zm6.5 5a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 17.5 18m1.5-5h-6V8h6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BusSchool(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 11.597V11h1V8h-1V6c0-2.206-1.794-4-4-4H7C4.794 2 3 3.794 3 6v2H2v3h1v.597a3.97 3.97 0 0 0-.999 2.648l.004 3.758c.001.733.404 1.369.995 1.716V21a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-1h12v1a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-1.274a2.02 2.02 0 0 0 .421-.313c.377-.378.585-.881.584-1.415l-.004-3.759A3.965 3.965 0 0 0 21 11.597M8 4h8v2H8zM6.5 17a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 6.5 17m4.5-5H5V8h6zm6.5 5a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 17.5 17m1.5-5h-6V8h6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Business(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 7h-6a1 1 0 0 0-1 1v3h-2V4a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V8a1 1 0 0 0-1-1M8 6h2v2H8zM6 16H4v-2h2zm0-4H4v-2h2zm0-4H4V6h2zm4 8H8v-2h2zm0-4H8v-2h2zm9 4h-2v-2h2zm0-4h-2v-2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cabinet(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 4c0-1.103-.897-2-2-2H5c-1.103 0-2 .897-2 2v7h18zm-5 4H8V5h2v1h4V5h2zM5 22h14c1.103 0 2-.897 2-2v-7H3v7c0 1.103.897 2 2 2m3-6h2v1h4v-1h2v3H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CableCar(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m2 9.76l9-2.45V10H7a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-8a2 2 0 0 0-2-2h-4V6.76l9-2.45V2.24L2 7.69zM11 12v3H7v-3zm6 0v3h-4v-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cake(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.997 15c-1.601 0-2.446-.676-3.125-1.219c-.567-.453-.977-.781-1.878-.781c-.898 0-1.287.311-1.874.78c-.679.544-1.524 1.22-3.125 1.22s-2.444-.676-3.123-1.22C3.285 13.311 2.897 13 2 13v5c0 1.103.897 2 2 2h16c1.103 0 2-.897 2-2v-5c-.899 0-1.288.311-1.876.781c-.68.543-1.525 1.219-3.127 1.219M19 5h-6V2h-2v3H5C3.346 5 2 6.346 2 8v3c1.6 0 2.443.676 3.122 1.22c.587.469.975.78 1.873.78c.899 0 1.287-.311 1.875-.781c.679-.543 1.524-1.219 3.124-1.219c1.602 0 2.447.676 3.127 1.219c.588.47.977.781 1.876.781c.9 0 1.311-.328 1.878-.781C19.554 11.676 20.399 11 22 11V8c0-1.654-1.346-3-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calculator(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 22h12a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2m3-3H7v-2h2zm0-4H7v-2h2zm0-4H7V9h2zm4 8h-2v-2h2zm0-4h-2v-2h2zm0-4h-2V9h2zm4 8h-2v-6h2zm0-8h-2V9h2zM6 4h12v3H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 20V6c0-1.103-.897-2-2-2h-2V2h-2v2H9V2H7v2H5c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2h14c1.103 0 2-.897 2-2M9 18H7v-2h2zm0-4H7v-2h2zm4 4h-2v-2h2zm0-4h-2v-2h2zm4 4h-2v-2h2zm0-4h-2v-2h2zm2-5H5V7h14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 22h14c1.103 0 2-.897 2-2V6c0-1.103-.897-2-2-2h-2V2h-2v2H9V2H7v2H5c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2M5 7h14v2H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarCheck(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 22h14c1.103 0 2-.897 2-2V6c0-1.103-.897-2-2-2h-2V2h-2v2H9V2H7v2H5c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2m6-3.586l-3.707-3.707l1.414-1.414L11 15.586l4.293-4.293l1.414 1.414zM5 7h14v2H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarEdit(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 22h14c1.103 0 2-.897 2-2V6c0-1.103-.897-2-2-2h-2V2h-2v2H9V2H7v2H5c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2m4.799-2.013H8v-1.799l4.977-4.97l1.799 1.799zm5.824-5.817l-1.799-1.799L15.196 11l1.799 1.799zM5 7h14v2H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarEvent(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 4h-2V2h-2v2H9V2H7v2H5c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2h14c1.103 0 2-.897 2-2V6c0-1.103-.897-2-2-2m-1 15h-6v-6h6zm1-10H5V7h14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarExclamation(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 4h-2V2h-2v2H9V2H7v2H5c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2h14c1.103 0 2-.897 2-2V6c0-1.103-.897-2-2-2m-6 16h-2v-2h2zm0-4h-2v-5h2zm6-7H5V7h14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarHeart(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 4h-2V2h-2v2H9V2H7v2H5c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2h14c1.103 0 2-.897 2-2V6c0-1.103-.897-2-2-2m-3.648 11.711L12.002 19l-3.349-3.289a2.129 2.129 0 0 1 0-3.069a2.224 2.224 0 0 1 3.125 0l.224.219l.224-.219a2.225 2.225 0 0 1 3.126 0a2.129 2.129 0 0 1 0 3.069M19 9H5V7h14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarMinus(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 22h14c1.103 0 2-.897 2-2V6c0-1.103-.897-2-2-2h-2V2h-2v2H9V2H7v2H5c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2m11-6H8v-2h8zM5 7h14v2H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarPlus(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 22h14c1.103 0 2-.897 2-2V6c0-1.103-.897-2-2-2h-2V2h-2v2H9V2H7v2H5c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2m11-6h-3v3h-2v-3H8v-2h3v-3h2v3h3zM5 7h14v2H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarStar(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 4h-2V2h-2v2H9V2H7v2H5c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2h14c1.103 0 2-.897 2-2V6c0-1.103-.897-2-2-2m-4.588 15l-2.449-1.288L9.514 19l.468-2.728L8 14.342l2.738-.398l1.225-2.48l1.225 2.48l2.738.398l-1.981 1.931zM19 9H5V7h14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarWeek(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 4h-2V2h-2v2H9V2H7v2H5c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2h14c1.103 0 2-.897 2-2V6c0-1.103-.897-2-2-2m-2 8v2H7v-3h10zm2-3H5V7h14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarX(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 22h14c1.103 0 2-.897 2-2V6c0-1.103-.897-2-2-2h-2V2h-2v2H9V2H7v2H5c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2m10.707-4.707l-1.414 1.414L12 16.414l-2.293 2.293l-1.414-1.414L10.586 15l-2.293-2.293l1.414-1.414L12 13.586l2.293-2.293l1.414 1.414L13.414 15zM5 7h14v2H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Camera(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 9c-1.626 0-3 1.374-3 3s1.374 3 3 3s3-1.374 3-3s-1.374-3-3-3"/><path fill="currentColor" d="M20 5h-2.586l-2.707-2.707A.996.996 0 0 0 14 2h-4a.996.996 0 0 0-.707.293L6.586 5H4c-1.103 0-2 .897-2 2v11c0 1.103.897 2 2 2h16c1.103 0 2-.897 2-2V7c0-1.103-.897-2-2-2m-8 12c-2.71 0-5-2.29-5-5s2.29-5 5-5s5 2.29 5 5s-2.29 5-5 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraHome(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="12" cy="10" r="3" fill="currentColor"/><path fill="currentColor" d="M18.125 2H5.875A1.877 1.877 0 0 0 4 3.875v12.25C4 17.159 4.841 18 5.875 18H11v2H7v2h10v-2h-4v-2h5.125A1.877 1.877 0 0 0 20 16.125V3.875A1.877 1.877 0 0 0 18.125 2M12 15c-2.757 0-5-2.243-5-5s2.243-5 5-5s5 2.243 5 5s-2.243 5-5 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraMovie(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 11c0-.959-.68-1.761-1.581-1.954C16.779 8.445 17 7.75 17 7c0-2.206-1.794-4-4-4c-1.516 0-2.822.857-3.5 2.104C8.822 3.857 7.516 3 6 3C3.794 3 2 4.794 2 7c0 .902.312 1.726.817 2.396A1.993 1.993 0 0 0 2 11v8c0 1.103.897 2 2 2h12c1.103 0 2-.897 2-2v-2.637l4 2v-7l-4 2zm-5-6c1.103 0 2 .897 2 2s-.897 2-2 2s-2-.897-2-2s.897-2 2-2M6 5c1.103 0 2 .897 2 2s-.897 2-2 2s-2-.897-2-2s.897-2 2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraOff(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 20h11.879l-3.083-3.083A4.774 4.774 0 0 1 12 17c-2.71 0-5-2.29-5-5c0-.271.039-.535.083-.796L2.144 6.265C2.054 6.493 2 6.74 2 7v11c0 1.103.897 2 2 2M20 5h-2.586l-2.707-2.707A.996.996 0 0 0 14 2h-4a.996.996 0 0 0-.707.293L6.586 5h-.172L3.707 2.293L2.293 3.707l18 18l1.414-1.414l-.626-.626A1.98 1.98 0 0 0 22 18V7c0-1.103-.897-2-2-2m-5.312 8.274A2.86 2.86 0 0 0 15 12c0-1.626-1.374-3-3-3c-.456 0-.884.12-1.274.312l-1.46-1.46A4.88 4.88 0 0 1 12 7c2.71 0 5 2.29 5 5a4.88 4.88 0 0 1-.852 2.734z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraPlus(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 5h-2.586l-2.707-2.707A.996.996 0 0 0 14 2h-4a.996.996 0 0 0-.707.293L6.586 5H4c-1.103 0-2 .897-2 2v11c0 1.103.897 2 2 2h16c1.103 0 2-.897 2-2V7c0-1.103-.897-2-2-2m-8 12c-2.71 0-5-2.29-5-5c0-2.711 2.29-5 5-5s5 2.289 5 5c0 2.71-2.29 5-5 5"/><path fill="currentColor" d="M13 9h-2v2H9v2h2v2h2v-2h2v-2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Capsule(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.434 20.566c1.335 0 2.591-.52 3.535-1.464l7.134-7.133a5.008 5.008 0 0 0-.001-7.072a4.969 4.969 0 0 0-3.536-1.463c-1.335 0-2.59.52-3.534 1.464l-7.134 7.133a5.01 5.01 0 0 0-.001 7.072a4.971 4.971 0 0 0 3.537 1.463m5.011-14.254a2.979 2.979 0 0 1 2.12-.878c.802 0 1.556.312 2.122.878a3.004 3.004 0 0 1 .001 4.243l-2.893 2.892l-4.242-4.244z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Captions(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 4H4c-1.103 0-2 .897-2 2v12c0 1.103.897 2 2 2h16c1.103 0 2-.897 2-2V6c0-1.103-.897-2-2-2m-9 6H8v4h3v2H8c-1.103 0-2-.897-2-2v-4c0-1.103.897-2 2-2h3zm7 0h-3v4h3v2h-3c-1.103 0-2-.897-2-2v-4c0-1.103.897-2 2-2h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Car(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.772 10.155l-1.368-4.104A2.995 2.995 0 0 0 16.559 4H7.441a2.995 2.995 0 0 0-2.845 2.051l-1.368 4.104A2 2 0 0 0 2 12v5c0 .738.404 1.376 1 1.723V21a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-2h12v2a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-2.277A1.99 1.99 0 0 0 22 17v-5a2 2 0 0 0-1.228-1.845M7.441 6h9.117c.431 0 .813.274.949.684L18.613 10H5.387l1.105-3.316A1 1 0 0 1 7.441 6M5.5 16a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 5.5 16m13 0a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 18.5 16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CarBattery(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 6H4c-1.103 0-2 .897-2 2v10c0 1.103.897 2 2 2h16c1.103 0 2-.897 2-2V8c0-1.103-.897-2-2-2M9 14H4v-2h5zm11 0h-2v2h-2v-2h-2v-2h2v-2h2v2h2zM4 3h4v2H4zm12 0h4v2h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CarCrash(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.634 17.918a1.765 1.765 0 0 0 1.201 1.291l.18.791H4v2h16v-2H6.683a.84.84 0 0 0-.007-.278l-.196-.863l10.357-2.356l.196.863a.886.886 0 0 0 1.06.667l.863-.197a.885.885 0 0 0 .667-1.06l-.251-1.103c.446-.416.67-1.046.525-1.683l-.59-2.59a1.76 1.76 0 0 0-1.262-1.307l-2.049-3.378a2.774 2.774 0 0 0-2.982-1.263l-7.868 1.79a2.769 2.769 0 0 0-2.144 2.43l-.387 3.932a1.76 1.76 0 0 0-.57 1.724zm3.02-.688a1.327 1.327 0 1 1-.59-2.589a1.327 1.327 0 0 1 .59 2.589m11.222-2.552a1.328 1.328 0 1 1-.59-2.587a1.328 1.328 0 0 1 .59 2.587M5.589 9.192l7.869-1.791a.773.773 0 0 1 .83.351l1.585 2.613l-.566.129l-10.046 2.287l-.568.129l.299-3.042a.772.772 0 0 1 .597-.676M18.405 4L17 2l-.5 3L19 9l3 1l-2-2.539l2-.933l-2-.933L22 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CarGarage(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 19.723V21a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-1h12v1a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-1.277A1.99 1.99 0 0 0 22 18v-3c0-.831-.507-1.542-1.228-1.845l-1.368-4.104A2.995 2.995 0 0 0 16.559 7H7.441a2.995 2.995 0 0 0-2.845 2.051l-1.368 4.104A2.001 2.001 0 0 0 2 15v3c0 .738.404 1.376 1 1.723M5.5 18a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 5.5 18m13 0a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 18.5 18M7.441 9h9.117a1 1 0 0 1 .949.684L18.613 13H5.387l1.105-3.316c.137-.409.519-.684.949-.684"/><path fill="currentColor" d="M22 7.388V5.279l-9.684-3.228a.996.996 0 0 0-.658.009L2 5.572V7.7l10.015-3.642z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CarMechanic(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.772 13.155l-1.368-4.104A2.995 2.995 0 0 0 16.559 7H7.441a2.995 2.995 0 0 0-2.845 2.051l-1.368 4.104A2.001 2.001 0 0 0 2 15v3c0 .738.404 1.376 1 1.723V21a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-1h12v1a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-1.277A1.99 1.99 0 0 0 22 18v-3c0-.831-.507-1.542-1.228-1.845M7.441 9h9.117a1 1 0 0 1 .949.684L18.613 13H5.387l1.105-3.316c.137-.409.519-.684.949-.684M5.5 18a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 5.5 18m13 0a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 18.5 18M5.277 5c.347.595.985 1 1.723 1s1.376-.405 1.723-1h6.555c.346.595.984 1 1.722 1s1.376-.405 1.723-1H17V3h1.723c-.347-.595-.985-1-1.723-1s-1.376.405-1.723 1H8.723C8.376 2.405 7.738 2 7 2s-1.376.405-1.723 1H7v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CarWash(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.772 13.155l-1.368-4.104A2.995 2.995 0 0 0 16.559 7H7.441a2.995 2.995 0 0 0-2.845 2.051l-1.368 4.104A2.001 2.001 0 0 0 2 15v3c0 .738.404 1.376 1 1.723V21a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-1h12v1a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-1.277A1.99 1.99 0 0 0 22 18v-3c0-.831-.507-1.542-1.228-1.845M7.441 9h9.117a1 1 0 0 1 .949.684L18.613 13H5.387l1.105-3.316c.137-.409.519-.684.949-.684M5.5 18a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 5.5 18m13 0a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 18.5 18M9 4.378c.005-1.088-1.037-2.051-1.387-2.344a.176.176 0 0 0-.228 0C7.033 2.325 5.995 3.271 6 4.377C6 5.272 6.673 6 7.5 6S9 5.272 9 4.378m4.5 0c.005-1.088-1.037-2.052-1.387-2.344a.176.176 0 0 0-.228 0c-.353.291-1.391 1.238-1.386 2.344C10.5 5.272 11.173 6 12 6s1.5-.728 1.5-1.622m4.5 0c.005-1.088-1.037-2.052-1.387-2.344a.176.176 0 0 0-.228 0c-.352.291-1.39 1.237-1.385 2.343C15 5.272 15.673 6 16.5 6S18 5.272 18 4.378"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Card(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 17c1.103 0 2-.897 2-2V5c0-1.103-.897-2-2-2H6c-1.103 0-2 .897-2 2v10c0 1.103.897 2 2 2zM4 19h16v2H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretDownCircle(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10c5.515 0 10-4.486 10-10S17.515 2 12 2m0 14l-5-6h10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretDownSquare(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 21h14c1.103 0 2-.897 2-2V5c0-1.103-.897-2-2-2H5c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2m12-11l-5 6l-5-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretLeftCircle(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m2 15l-6-5l6-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretLeftSquare(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 21h14c1.103 0 2-.897 2-2V5c0-1.103-.897-2-2-2H5c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2m9-14v10l-6-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretRightCircle(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22c5.514 0 10-4.486 10-10S17.514 2 12 2S2 6.486 2 12s4.486 10 10 10M10 7l6 5l-6 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretRightSquare(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 21h14c1.103 0 2-.897 2-2V5c0-1.103-.897-2-2-2H5c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2m5-14l6 5l-6 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretUpCircle(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.485 2 2 6.486 2 12s4.485 10 10 10c5.514 0 10-4.486 10-10S17.514 2 12 2M7 14l5-6l5 6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretUpSquare(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 21h14c1.103 0 2-.897 2-2V5c0-1.103-.897-2-2-2H5c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2m7-13l5 6H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Carousel(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3H8c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2h8c1.103 0 2-.897 2-2V5c0-1.103-.897-2-2-2M2 7v10c0 1.103.897 2 2 2V5c-1.103 0-2 .897-2 2m18-2v14c1.103 0 2-.897 2-2V7c0-1.103-.897-2-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cart(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.822 7.431A1 1 0 0 0 21 7H7.333L6.179 4.23A1.994 1.994 0 0 0 4.333 3H2v2h2.333l4.744 11.385A1 1 0 0 0 10 17h8c.417 0 .79-.259.937-.648l3-8a1 1 0 0 0-.115-.921"/><circle cx="10.5" cy="19.5" r="1.5" fill="currentColor"/><circle cx="17.5" cy="19.5" r="1.5" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CartAdd(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="10.5" cy="19.5" r="1.5" fill="currentColor"/><circle cx="17.5" cy="19.5" r="1.5" fill="currentColor"/><path fill="currentColor" d="M21 7H7.334L6.18 4.23A1.995 1.995 0 0 0 4.333 3H2v2h2.334l4.743 11.385c.155.372.52.615.923.615h8c.417 0 .79-.259.937-.648l3-8A1.003 1.003 0 0 0 21 7m-4 6h-2v2h-2v-2h-2v-2h2V9h2v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CartAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 4H2v2h2.3l3.521 9.683A2.004 2.004 0 0 0 9.7 17H18v-2H9.7l-.728-2H18c.4 0 .762-.238.919-.606l3-7A.998.998 0 0 0 21 4"/><circle cx="10.5" cy="19.5" r="1.5" fill="currentColor"/><circle cx="16.5" cy="19.5" r="1.5" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CartDownload(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="10.5" cy="19.5" r="1.5" fill="currentColor"/><circle cx="17.5" cy="19.5" r="1.5" fill="currentColor"/><path fill="currentColor" d="M21 7H7.334L6.18 4.23A1.995 1.995 0 0 0 4.333 3H2v2h2.334l4.743 11.385c.155.372.52.615.923.615h8c.417 0 .79-.259.937-.648l3-8A1.003 1.003 0 0 0 21 7m-7 8l-3-3h2V9h2v3h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Castle(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 11h-2V6h1V2h-2v2h-1V2h-2v2h-1V2h-2v2h-1V2H8v2H7V2H5v4h1v5H4V9H2v12h7v-5a3 3 0 0 1 6 0v5h7V9h-2zm-10-1H8V7h2zm6 0h-2V7h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cat(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 14a5 5 0 0 0 2.71-.81L20 13a3.16 3.16 0 0 0 .45-.37l.21-.2a4.48 4.48 0 0 0 .48-.58l.06-.08a4.28 4.28 0 0 0 .41-.76a1.57 1.57 0 0 0 .09-.23a4.21 4.21 0 0 0 .2-.63l.06-.25A5.5 5.5 0 0 0 22 9V2l-3 3h-4l-3-3v7a5 5 0 0 0 5 5m2-7a1 1 0 1 1-1 1a1 1 0 0 1 1-1m-4 0a1 1 0 1 1-1 1a1 1 0 0 1 1-1"/><path fill="currentColor" d="M11 22v-5H8v5H5V11.9a3.49 3.49 0 0 1-2.48-1.64A3.59 3.59 0 0 1 2 8.5A3.65 3.65 0 0 1 6 5a1.89 1.89 0 0 0 2-2a1 1 0 0 1 1-1a1 1 0 0 1 1 1a3.89 3.89 0 0 1-4 4C4.19 7 4 8.16 4 8.51S4.18 10 6 10h5.09A6 6 0 0 0 19 14.65V22h-3v-5h-2v5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Category(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 11h6a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1m10 0h6a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1h-6a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1M4 21h6a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1m13 0c2.206 0 4-1.794 4-4s-1.794-4-4-4s-4 1.794-4 4s1.794 4 4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CategoryAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 3H4a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1m10 10h-6a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1M17 3c-2.206 0-4 1.794-4 4s1.794 4 4 4s4-1.794 4-4s-1.794-4-4-4M7 13c-2.206 0-4 1.794-4 4s1.794 4 4 4s4-1.794 4-4s-1.794-4-4-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cctv(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.618 7.462L6.403 2.085a1.007 1.007 0 0 0-.77-.016a1.002 1.002 0 0 0-.552.537l-3 7a1 1 0 0 0 .525 1.313L9.563 13.9L8.323 17H4v-3H2v8h2v-3h4.323c.823 0 1.552-.494 1.856-1.258l1.222-3.054l3.419 1.465a1 1 0 0 0 1.311-.518l3-6.857a1 1 0 0 0-.513-1.316m1.312 8.91l-1.858-.742l1.998-5l1.858.741z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Certification(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.047 14.667a.992.992 0 0 0 .466.607l1.909 1.104v2.199a1 1 0 0 0 1 1h2.199l1.104 1.91a1.002 1.002 0 0 0 1.366.366L12 20.75l1.91 1.104a1.002 1.002 0 0 0 1.366-.366l1.103-1.909h2.199a1 1 0 0 0 1-1V16.38l1.909-1.104a.999.999 0 0 0 .366-1.366L20.75 12l1.104-1.909a1 1 0 0 0-.366-1.366l-1.909-1.104V5.422a1 1 0 0 0-1-1H16.38l-1.103-1.909a1.004 1.004 0 0 0-.607-.466a.994.994 0 0 0-.759.1L12 3.25l-1.909-1.104a.998.998 0 0 0-1.366.365l-1.104 1.91H5.422a1 1 0 0 0-1 1V7.62L2.513 8.725a1.001 1.001 0 0 0-.365 1.366L3.251 12l-1.104 1.909a1.003 1.003 0 0 0-.1.758"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chalkboard(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 4H4c-1.103 0-2 .897-2 2v10c0 1.103.897 2 2 2h4l-1.8 2.4l1.6 1.2l2.7-3.6h3l2.7 3.6l1.6-1.2L16 18h4c1.103 0 2-.897 2-2V6c0-1.103-.897-2-2-2M5 13h4v2H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chart(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 21c1.103 0 2-.897 2-2V5c0-1.103-.897-2-2-2H5c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2zM9.553 9.658l4 2l1.553-3.105l1.789.895l-2.447 4.895l-4-2l-1.553 3.105l-1.789-.895z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chat(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 18h2v4.081L11.101 18H16c1.103 0 2-.897 2-2V8c0-1.103-.897-2-2-2H4c-1.103 0-2 .897-2 2v8c0 1.103.897 2 2 2"/><path fill="currentColor" d="M20 2H8c-1.103 0-2 .897-2 2h12c1.103 0 2 .897 2 2v8c1.103 0 2-.897 2-2V4c0-1.103-.897-2-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckCircle(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m-1.999 14.413l-3.713-3.705L7.7 11.292l2.299 2.295l5.294-5.294l1.414 1.414z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckShield(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.488 21.754c.294.157.663.156.957-.001c8.012-4.304 8.581-12.713 8.574-15.104a.988.988 0 0 0-.596-.903l-8.05-3.566a1.005 1.005 0 0 0-.813.001L3.566 5.747a.99.99 0 0 0-.592.892c-.034 2.379.445 10.806 8.514 15.115M8.674 10.293l2.293 2.293l4.293-4.293l1.414 1.414l-5.707 5.707l-3.707-3.707z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckSquare(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 3H5c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2h14c1.103 0 2-.897 2-2V5c0-1.103-.897-2-2-2m-7.933 13.481l-3.774-3.774l1.414-1.414l2.226 2.226l4.299-5.159l1.537 1.28z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Checkbox(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 19h10a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckboxChecked(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 5a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2zm4 10.414l-2.707-2.707l1.414-1.414L11 12.586l3.793-3.793l1.414 1.414z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckboxMinus(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 5H7a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2m-1 8H8v-2h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cheese(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.16 2a1 1 0 0 0-.66.13l-12 7a.64.64 0 0 0-.13.1l-.1.08a1.17 1.17 0 0 0-.17.26a.84.84 0 0 0-.1.43v10a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V10a8.08 8.08 0 0 0-6.84-8m0 2.05A6.07 6.07 0 0 1 19.93 9H6.7zM6.5 18A1.5 1.5 0 1 1 8 16.5A1.5 1.5 0 0 1 6.5 18m5-3a1.5 1.5 0 1 1 1.5-1.5a1.5 1.5 0 0 1-1.5 1.5m5.5 3a2 2 0 1 1 2-2a2 2 0 0 1-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chess(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1m-1 6h-4v4h4v4h-4v4h-4v-4H8v4H4v-4h4v-4H4V8h4V4h4v4h4V4h4z"/><path fill="currentColor" d="M8 8h4v4H8zm4 4h4v4h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDown(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.939 7.939L12 12.879l-4.939-4.94l-2.122 2.122L12 17.121l7.061-7.06z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDownCircle(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m0 14.414l-5.707-5.707l1.414-1.414L12 13.586l4.293-4.293l1.414 1.414z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDownSquare(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 21h14c1.103 0 2-.897 2-2V5c0-1.103-.897-2-2-2H5c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2M7.707 9.293L12 13.586l4.293-4.293l1.414 1.414L12 16.414l-5.707-5.707z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeft(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.939 4.939L6.879 12l7.06 7.061l2.122-2.122L11.121 12l4.94-4.939z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeftCircle(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m2.707 14.293l-1.414 1.414L7.586 12l5.707-5.707l1.414 1.414L10.414 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeftSquare(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 3c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2h14c1.103 0 2-.897 2-2V5c0-1.103-.897-2-2-2zm9.707 13.293l-1.414 1.414L7.586 12l5.707-5.707l1.414 1.414L10.414 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRight(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.061 19.061L17.121 12l-7.06-7.061l-2.122 2.122L12.879 12l-4.94 4.939z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRightCircle(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m-1.293 15.707l-1.414-1.414L13.586 12L9.293 7.707l1.414-1.414L16.414 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRightSquare(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 5v14c0 1.103.897 2 2 2h14c1.103 0 2-.897 2-2V5c0-1.103-.897-2-2-2H5c-1.103 0-2 .897-2 2m6.293 2.707l1.414-1.414L16.414 12l-5.707 5.707l-1.414-1.414L13.586 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUp(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12 6.879l-7.061 7.06l2.122 2.122L12 11.121l4.939 4.94l2.122-2.122z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUpCircle(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m4.293 12.707L12 10.414l-4.293 4.293l-1.414-1.414L12 7.586l5.707 5.707z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUpSquare(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 21h14c1.103 0 2-.897 2-2V5c0-1.103-.897-2-2-2H5c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2m7-13.414l5.707 5.707l-1.414 1.414L12 10.414l-4.293 4.293l-1.414-1.414z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronsDown(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.939 10.939L12 15.879l-4.939-4.94l-2.122 2.122L12 20.121l7.061-7.06z"/><path fill="currentColor" d="M16.939 3.939L12 8.879l-4.939-4.94l-2.122 2.122L12 13.121l7.061-7.06z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronsLeft(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8.121 12l4.94-4.939l-2.122-2.122L3.879 12l7.06 7.061l2.122-2.122z"/><path fill="currentColor" d="M17.939 4.939L10.879 12l7.06 7.061l2.122-2.122L15.121 12l4.94-4.939z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronsRight(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.061 4.939l-2.122 2.122L15.879 12l-4.94 4.939l2.122 2.122L20.121 12z"/><path fill="currentColor" d="M6.061 19.061L13.121 12l-7.06-7.061l-2.122 2.122L8.879 12l-4.94 4.939z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronsUp(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12 3.879l-7.061 7.06l2.122 2.122L12 8.121l4.939 4.94l2.122-2.122z"/><path fill="currentColor" d="m4.939 17.939l2.122 2.122L12 15.121l4.939 4.94l2.122-2.122L12 10.879z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chip(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 7a2 2 0 0 0-2-2h-1V2h-2v3h-4V2H8v3H7a2 2 0 0 0-2 2v1H2v2h3v4H2v2h3v1a2 2 0 0 0 2 2h1v3h2v-3h4v3h2v-3h1a2 2 0 0 0 2-2v-1h3v-2h-3v-4h3V8h-3zm-4 8H9V9h6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Church(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 22v-4a2 2 0 0 1 4 0v4h4V12a1 1 0 0 0-.485-.857L13 8.434V6h2V4h-2V2h-2v2H9v2h2v2.434l-4.515 2.709A1 1 0 0 0 6 12v10zm-7 0h2v-8.118l-2.447 1.224A.998.998 0 0 0 2 16v5a1 1 0 0 0 1 1m18.447-6.895L19 13.882V22h2a1 1 0 0 0 1-1v-5c0-.379-.214-.725-.553-.895"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Circle(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleHalf(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2h-1v20h1a10 10 0 0 0 0-20"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleQuarter(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2h-1v11h11v-1A10 10 0 0 0 12 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleThreeQuarter(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2h-1v9H2v1a10 10 0 0 0 17.07 7.07A10 10 0 0 0 12 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func City(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 6h-4V3a1 1 0 0 0-1-1H8a1 1 0 0 0-1 1v7H3a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1M6 18H4v-2h2zm0-4H4v-2h2zm5 4H9v-2h2zm0-4H9v-2h2zm0-4H9V8h2zm0-4H9V4h2zm4 12h-2v-2h2zm0-4h-2v-2h2zm0-4h-2V8h2zm0-4h-2V4h2zm5 12h-2v-2h2zm0-4h-2v-2h2zm0-4h-2V8h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clinic(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.743 12.331l-9-10c-.379-.422-1.107-.422-1.486 0l-9 10a.998.998 0 0 0-.17 1.076c.16.361.518.593.913.593h2v7a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-7h2a.998.998 0 0 0 .743-1.669M16 15h-3v3h-2v-3H8v-2h3v-3h2v3h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloud(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.944 11.112C18.507 7.67 15.56 5 12 5C9.244 5 6.85 6.611 5.757 9.15C3.609 9.792 2 11.82 2 14c0 2.757 2.243 5 5 5h11c2.206 0 4-1.794 4-4a4.01 4.01 0 0 0-3.056-3.888"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudDownload(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.944 11.112C18.507 7.67 15.56 5 12 5C9.244 5 6.85 6.61 5.757 9.149C3.609 9.792 2 11.82 2 14c0 2.657 2.089 4.815 4.708 4.971V19H17.99v-.003L18 19c2.206 0 4-1.794 4-4a4.008 4.008 0 0 0-3.056-3.888M8 12h3V9h2v3h3l-4 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudLightning(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.944 10.112C18.507 6.67 15.56 4 12 4C9.244 4 6.85 5.611 5.757 8.15C3.609 8.792 2 10.82 2 13c0 2.757 2.243 5 5 5h1.333L10 13h4l-2 3h2.975l-1.325 2H18c2.206 0 4-1.794 4-4a4.01 4.01 0 0 0-3.056-3.888M11 18H8.333L8 19h3v3l2.649-4H11.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudRain(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.944 10.112C18.507 6.67 15.56 4 12 4C9.244 4 6.85 5.611 5.757 8.15C3.609 8.792 2 10.82 2 13c0 2.757 2.243 5 5 5h1v3h2v-3h4v3h2v-3h2c2.206 0 4-1.794 4-4a4.01 4.01 0 0 0-3.056-3.888"/><path fill="currentColor" d="M11 19h2v3h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudUpload(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.944 11.112C18.507 7.67 15.56 5 12 5C9.244 5 6.85 6.611 5.757 9.15C3.609 9.792 2 11.82 2 14c0 2.757 2.243 5 5 5h11c2.206 0 4-1.794 4-4a4.01 4.01 0 0 0-3.056-3.888M13 14v3h-2v-3H8l4-5l4 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Coffee(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 2h2v3H5zm4 0h2v3H9zm4 0h2v3h-2zm6 7h-2V8a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v10a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3h2c1.103 0 2-.897 2-2v-5c0-1.103-.897-2-2-2m-2 7v-5h2l.002 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoffeeAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 5h-1V4a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v9a4 4 0 0 0 4 4h6c1.858 0 3.411-1.279 3.858-3H19a3 3 0 0 0 3-3V8a3 3 0 0 0-3-3m1 6a1 1 0 0 1-1 1h-1V7h1a1 1 0 0 1 1 1zm-2 8H3c0 1.654 1.346 3 3 3h11c1.654 0 3-1.346 3-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoffeeBean(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 18a6.06 6.06 0 0 0 5.17-6a7.62 7.62 0 0 1 6.52-7.51l2.59-.37c-.07-.08-.13-.16-.21-.24c-3.26-3.26-9.52-2.28-14 2.18C2.28 9.9 1 15 2.76 18.46z"/><path fill="currentColor" d="M12.73 12a7.63 7.63 0 0 1-6.51 7.52l-2.46.35l.15.17c3.26 3.26 9.52 2.29 14-2.17C21.68 14.11 23 9 21.25 5.59l-3.34.48A6.05 6.05 0 0 0 12.73 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoffeeTogo(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m18.5 5l-1.224-2.447A1 1 0 0 0 16.382 2H7.618a1 1 0 0 0-.894.553L5.5 5H3v2h18V5zM6.734 21.142c.071.492.493.858.991.858h8.551a1 1 0 0 0 .99-.858L19 9H5zM16 12l-.714 5H8.714L8 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cog(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m2.344 15.271l2 3.46a1 1 0 0 0 1.366.365l1.396-.806c.58.457 1.221.832 1.895 1.112V21a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-1.598a8.094 8.094 0 0 0 1.895-1.112l1.396.806c.477.275 1.091.11 1.366-.365l2-3.46a1.004 1.004 0 0 0-.365-1.366l-1.372-.793a7.683 7.683 0 0 0-.002-2.224l1.372-.793c.476-.275.641-.89.365-1.366l-2-3.46a1 1 0 0 0-1.366-.365l-1.396.806A8.034 8.034 0 0 0 15 4.598V3a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1v1.598A8.094 8.094 0 0 0 7.105 5.71L5.71 4.904a.999.999 0 0 0-1.366.365l-2 3.46a1.004 1.004 0 0 0 .365 1.366l1.372.793a7.683 7.683 0 0 0 0 2.224l-1.372.793c-.476.275-.641.89-.365 1.366M12 8c2.206 0 4 1.794 4 4s-1.794 4-4 4s-4-1.794-4-4s1.794-4 4-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Coin(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 5C7.031 5 2 6.546 2 9.5S7.031 14 12 14c4.97 0 10-1.546 10-4.5S16.97 5 12 5m-5 9.938v3c1.237.299 2.605.482 4 .541v-3a21.166 21.166 0 0 1-4-.541m6 .54v3a20.994 20.994 0 0 0 4-.541v-3a20.994 20.994 0 0 1-4 .541m6-1.181v3c1.801-.755 3-1.857 3-3.297v-3c0 1.44-1.199 2.542-3 3.297m-14 3v-3C3.2 13.542 2 12.439 2 11v3c0 1.439 1.2 2.542 3 3.297"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoinStack(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 10c3.976 0 8-1.374 8-4s-4.024-4-8-4s-8 1.374-8 4s4.024 4 8 4"/><path fill="currentColor" d="M4 10c0 2.626 4.024 4 8 4s8-1.374 8-4V8c0 2.626-4.024 4-8 4s-8-1.374-8-4z"/><path fill="currentColor" d="M4 14c0 2.626 4.024 4 8 4s8-1.374 8-4v-2c0 2.626-4.024 4-8 4s-8-1.374-8-4z"/><path fill="currentColor" d="M4 18c0 2.626 4.024 4 8 4s8-1.374 8-4v-2c0 2.626-4.024 4-8 4s-8-1.374-8-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Collection(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 10H5c-1.103 0-2 .897-2 2v8c0 1.103.897 2 2 2h14c1.103 0 2-.897 2-2v-8c0-1.103-.897-2-2-2M5 6h14v2H5zm2-4h10v2H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Color(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.08 11.25A4.84 4.84 0 0 1 8 9.05L4.43 5.49A9.88 9.88 0 0 0 2 11.25zM9.05 8a4.84 4.84 0 0 1 2.2-.91V2a9.88 9.88 0 0 0-5.76 2.43zm3.7-6v5A4.84 4.84 0 0 1 15 8l3.56-3.56A9.88 9.88 0 0 0 12.75 2M8 15a4.84 4.84 0 0 1-.91-2.2H2a9.88 9.88 0 0 0 2.39 5.76zm3.25 1.92a4.84 4.84 0 0 1-2.2-.92l-3.56 3.57A9.88 9.88 0 0 0 11.25 22zM16 9.05a4.84 4.84 0 0 1 .91 2.2h5a9.88 9.88 0 0 0-2.39-5.76zM15 16a4.84 4.84 0 0 1-2.2.91v5a9.88 9.88 0 0 0 5.76-2.39zm1.92-3.25A4.84 4.84 0 0 1 16 15l3.56 3.56A9.88 9.88 0 0 0 22 12.75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ColorFill(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 14c-.092.064-2 2.083-2 3.5c0 1.494.949 2.448 2 2.5c.906.044 2-.891 2-2.5c0-1.5-1.908-3.436-2-3.5M9.586 20c.378.378.88.586 1.414.586s1.036-.208 1.414-.586l7-7l-.707-.707L11 4.586L8.707 2.293L7.293 3.707L9.586 6L4 11.586c-.378.378-.586.88-.586 1.414s.208 1.036.586 1.414zM11 7.414L16.586 13H5.414z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Comment(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 2H4c-1.103 0-2 .897-2 2v18l4-4h14c1.103 0 2-.897 2-2V4c0-1.103-.897-2-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentAdd(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 2H4c-1.103 0-2 .897-2 2v18l4-4h14c1.103 0 2-.897 2-2V4c0-1.103-.897-2-2-2m-3 9h-4v4h-2v-4H7V9h4V5h2v4h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentCheck(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 2H4c-1.103 0-2 .897-2 2v18l4-4h14c1.103 0 2-.897 2-2V4c0-1.103-.897-2-2-2m-9 11.914l-3.707-3.707l1.414-1.414L11 11.086l4.793-4.793l1.414 1.414z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentDetail(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 1.999H4c-1.103 0-2 .897-2 2v18l4-4h14c1.103 0 2-.897 2-2v-12c0-1.103-.897-2-2-2m-6 11H7v-2h7zm3-4H7v-2h10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentDots(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 2H4c-1.103 0-2 .897-2 2v18l4-4h14c1.103 0 2-.897 2-2V4c0-1.103-.897-2-2-2M9 12a2 2 0 1 1 .001-4.001A2 2 0 0 1 9 12m6 0a2 2 0 1 1 .001-4.001A2 2 0 0 1 15 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentEdit(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 2H4c-1.103 0-2 .897-2 2v18l4-4h14c1.103 0 2-.897 2-2V4c0-1.103-.897-2-2-2M8.999 14.987H7v-1.999l5.53-5.522l1.998 1.999zm6.472-6.464l-1.999-1.999l1.524-1.523L16.995 7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentError(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 2H4c-1.103 0-2 .897-2 2v18l4-4h14c1.103 0 2-.897 2-2V4c0-1.103-.897-2-2-2m-7 13h-2v-2h2zm0-4h-2V5h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentMinus(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 2H4c-1.103 0-2 .897-2 2v18l4-4h14c1.103 0 2-.897 2-2V4c0-1.103-.897-2-2-2m-4 9H8V9h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentX(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 2H4c-1.103 0-2 .897-2 2v18l4-4h14c1.103 0 2-.897 2-2V4c0-1.103-.897-2-2-2m-3.294 11.543l-1.414 1.414l-3.293-3.292l-3.292 3.292l-1.414-1.414l3.292-3.292l-3.292-3.293l1.414-1.414l3.292 3.292l3.293-3.292l1.414 1.414l-3.292 3.293z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compass(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m3 13l-8 2l2-8l8-2z"/><circle cx="12" cy="12" r="2" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Component(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m3.553 18.895l4 2a1.001 1.001 0 0 0 .894 0L12 19.118l3.553 1.776a.99.99 0 0 0 .894.001l4-2c.339-.17.553-.516.553-.895v-5c0-.379-.214-.725-.553-.895L17 10.382V6c0-.379-.214-.725-.553-.895l-4-2a1 1 0 0 0-.895 0l-4 2C7.214 5.275 7 5.621 7 6v4.382l-3.447 1.724A.998.998 0 0 0 3 13v5c0 .379.214.725.553.895M8 12.118l2.264 1.132l-2.913 1.457l-2.264-1.132zm4-2.5l3-1.5v2.264l-3 1.5zm6.264 3.632l-2.882 1.441l-2.264-1.132L16 12.118zM8 18.882l-.062-.031V16.65L11 15.118v2.264zm8 0v-2.264l3-1.5v2.264zM12 5.118l2.264 1.132l-2.882 1.441l-2.264-1.132z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Confused(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m-5 8.5a1.5 1.5 0 1 1 3.001.001A1.5 1.5 0 0 1 7 10.5m1.124 6.492l-.248-1.984l8-1l.248 1.984zm7.369-5.006a1.494 1.494 0 1 1 .001-2.987a1.494 1.494 0 0 1-.001 2.987"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Contact(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2H6a2 2 0 0 0-2 2v3H2v2h2v2H2v2h2v2H2v2h2v3a2 2 0 0 0 2 2h15a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1m-8 2.999c1.648 0 3 1.351 3 3A3.012 3.012 0 0 1 13 11c-1.647 0-3-1.353-3-3.001c0-1.649 1.353-3 3-3M19 18H7v-.75c0-2.219 2.705-4.5 6-4.5s6 2.281 6 4.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Conversation(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 14h3.5c.827 0 1.5-.673 1.5-1.5v-9c0-.827-.673-1.5-1.5-1.5h-13C2.673 2 2 2.673 2 3.5V18l5.333-4zm-9-.1l.154-.016L4 14z"/><path fill="currentColor" d="M20.5 8H20v6.001c0 1.1-.893 1.993-1.99 1.999H8v.5c0 .827.673 1.5 1.5 1.5h7.167L22 22V9.5c0-.827-.673-1.5-1.5-1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cookie(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.598 11.064a1.006 1.006 0 0 0-.854-.172A2.938 2.938 0 0 1 20 11c-1.654 0-3-1.346-3.003-2.938c.005-.034.016-.134.017-.168a.998.998 0 0 0-1.254-1.006A3.002 3.002 0 0 1 15 7c-1.654 0-3-1.346-3-3c0-.217.031-.444.099-.716a1 1 0 0 0-1.067-1.236A9.956 9.956 0 0 0 2 12c0 5.514 4.486 10 10 10s10-4.486 10-10c0-.049-.003-.097-.007-.16a1.004 1.004 0 0 0-.395-.776M8.5 6a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3m-2 8a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3m3 4a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3m2.5-6.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0m3.5 6.5a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cool(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m4.095 14.244a5.982 5.982 0 0 1-3.034 1.634a6.05 6.05 0 0 1-2.414 0a5.919 5.919 0 0 1-2.148-.903a6.078 6.078 0 0 1-1.621-1.622l1.658-1.117c.143.211.307.41.488.59a3.988 3.988 0 0 0 1.273.86c.243.102.495.181.749.232a4.108 4.108 0 0 0 1.616 0c.253-.052.505-.131.75-.233c.234-.1.464-.224.679-.368c.208-.142.407-.306.591-.489c.183-.182.347-.381.489-.592l1.658 1.117c-.215.32-.462.62-.734.891M19 10a2 2 0 0 1-2 2h-2a2 2 0 0 1-2-2h-2a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V8h14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copy(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 8H4c-1.103 0-2 .897-2 2v10c0 1.103.897 2 2 2h10c1.103 0 2-.897 2-2V10c0-1.103-.897-2-2-2"/><path fill="currentColor" d="M20 2H10a2 2 0 0 0-2 2v2h8a2 2 0 0 1 2 2v8h2a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CopyAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 2H10a2 2 0 0 0-2 2v2h8a2 2 0 0 1 2 2v8h2a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2"/><path fill="currentColor" d="M4 22h10c1.103 0 2-.897 2-2V10c0-1.103-.897-2-2-2H4c-1.103 0-2 .897-2 2v10c0 1.103.897 2 2 2m2-10h6v2H6zm0 4h6v2H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copyright(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.579 2 2 6.58 2 12s4.579 10 10 10s10-4.58 10-10S17.421 2 12 2m0 13c.992 0 1.85-.265 2.293-.708l1.414 1.415C14.581 16.832 12.901 17 12 17c-2.757 0-5-2.243-5-5s2.243-5 5-5c.901 0 2.582.168 3.707 1.293l-1.414 1.414C13.851 9.264 12.993 9 12 9c-1.626 0-3 1.374-3 3s1.374 3 3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Coupon(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 5H3a1 1 0 0 0-1 1v4h.893c.996 0 1.92.681 2.08 1.664A2.001 2.001 0 0 1 3 14H2v4a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1v-4h-1a2.001 2.001 0 0 1-1.973-2.336c.16-.983 1.084-1.664 2.08-1.664H22V6a1 1 0 0 0-1-1M11 17H9v-2h2zm0-4H9v-2h2zm0-4H9V7h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCard(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 4H4c-1.103 0-2 .897-2 2v2h20V6c0-1.103-.897-2-2-2M2 18c0 1.103.897 2 2 2h16c1.103 0 2-.897 2-2v-6H2zm3-3h6v2H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCardAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 4H4c-1.103 0-2 .897-2 2v12c0 1.103.897 2 2 2h16c1.103 0 2-.897 2-2V6c0-1.103-.897-2-2-2m-7.5 12a2.5 2.5 0 1 1 0-5a2.47 2.47 0 0 1 1.5.512c-.604.456-1 1.173-1 1.988s.396 1.532 1 1.988a2.47 2.47 0 0 1-1.5.512m4 0a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCardFront(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 4H4c-1.103 0-2 .897-2 2v12c0 1.103.897 2 2 2h16c1.103 0 2-.897 2-2V6c0-1.103-.897-2-2-2M5 8.5a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-.5.5h-3a.5.5 0 0 1-.5-.5zm6 7.5H5v-2h6zm8 0h-6v-2h6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CricketBall(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m3.67 16.26l.54.53l-.62.61a9 9 0 0 0 .84 1.11L18.51 4.42a10.93 10.93 0 0 0-1.1-.83l-.62.61l-.53-.53l.48-.48A10 10 0 0 0 3.2 16.74zM14.86 5.07l.53.53L14 7l-.53-.53zm-2.79 2.8l.52.53l-1.39 1.4l-.53-.53zm-2.8 2.8l.53.53l-1.4 1.39l-.53-.53zm-2.8 2.79L7 14l-1.4 1.4l-.53-.53zm.12 6.95l.62-.61l.53.53l-.48.48A10 10 0 0 0 20.8 7.26l-.47.48l-.54-.53l.62-.61a9 9 0 0 0-.84-1.11L5.49 19.58a10.93 10.93 0 0 0 1.1.83M18.4 8.61l.53.53l-1.4 1.4L17 10zm-2.8 2.8l.53.53l-1.4 1.39l-.53-.53zm-2.8 2.79l.53.53l-1.39 1.4l-.54-.53zM10 17l.53.53l-1.4 1.4l-.53-.53z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crop(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 7.5C19 6.121 17.879 5 16.5 5H8V2H5v3H2v3h14v14h3v-3h3v-3h-3z"/><path fill="currentColor" d="M8 10H5v6.5C5 17.879 6.121 19 7.5 19H14v-3H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crown(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21 2l-5 5l-4-5l-4 5l-5-5v13h18zM5 21h14a2 2 0 0 0 2-2v-2H3v2a2 2 0 0 0 2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cube(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.406 6.086l-9-4a1.001 1.001 0 0 0-.813 0l-9 4c-.02.009-.034.024-.054.035c-.028.014-.058.023-.084.04c-.022.015-.039.034-.06.05a.87.87 0 0 0-.19.194c-.02.028-.041.053-.059.081a1.119 1.119 0 0 0-.076.165c-.009.027-.023.052-.031.079A1.013 1.013 0 0 0 2 7v10c0 .396.232.753.594.914l9 4c.13.058.268.086.406.086a.997.997 0 0 0 .402-.096l.004.01l9-4A.999.999 0 0 0 22 17V7a.999.999 0 0 0-.594-.914M12 4.095L18.538 7L12 9.905l-1.308-.581L5.463 7zm1 15.366V11.65l7-3.111v7.812z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CubeAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.895 3.553A1.001 1.001 0 0 0 17 3H7c-.379 0-.725.214-.895.553l-4 8a1 1 0 0 0 0 .895l4 8c.17.338.516.552.895.552h10c.379 0 .725-.214.895-.553l4-8a1 1 0 0 0 0-.895zM19.382 11h-7.764l-3-6h7.764zm-3 8H8.618l3-6h7.764z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cuboid(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.707 2.293A.996.996 0 0 0 16 2H8c-.414 0-.785.255-.934.641l-5 13a.999.999 0 0 0 .227 1.066l5 5A.996.996 0 0 0 8 22h8c.414 0 .785-.255.934-.641l5-13a.999.999 0 0 0-.227-1.066zM18.585 7h-5.171l-3-3h5.172zm-3.272 13h-6.23l4.583-11h5.878z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Customize(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 3H4a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1m10 0h-6a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1M10 13H4a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1m8 1h-2v2h-2v2h2v2h2v-2h2v-2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cylinder(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22c5.131 0 9-1.935 9-4.5V7c0-.051-.024-.097-.033-.146c.016-.117.033-.234.033-.354C21 3.935 17.131 2 12 2S3 3.935 3 6.5v11c0 2.565 3.869 4.5 9 4.5m0-18c4.273 0 7 1.48 7 2.5a.683.683 0 0 1-.025.158c-.004.01-.012.018-.015.027c-.274.848-2.29 1.98-5.496 2.253l-.05.003C12.965 8.979 12.494 9 12 9C7.727 9 5 7.52 5 6.5S7.727 4 12 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dashboard(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 13h6a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1m-1 7a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1zm10 0a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1v-7a1 1 0 0 0-1-1h-6a1 1 0 0 0-1 1zm1-10h6a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1h-6a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Data(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 6c0-2.168-3.663-4-8-4S4 3.832 4 6v2c0 2.168 3.663 4 8 4s8-1.832 8-4zm-8 13c-4.337 0-8-1.832-8-4v3c0 2.168 3.663 4 8 4s8-1.832 8-4v-3c0 2.168-3.663 4-8 4"/><path fill="currentColor" d="M20 10c0 2.168-3.663 4-8 4s-8-1.832-8-4v3c0 2.168 3.663 4 8 4s8-1.832 8-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Detail(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 3H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2m-9 14H5v-2h6zm8-4H5v-2h14zm0-4H5V7h14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Devices(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 3H7c-1.103 0-2 .897-2 2v2H4c-1.103 0-2 .897-2 2v10c0 1.103.897 2 2 2h6c1.103 0 2-.897 2-2h8c1.103 0 2-.897 2-2V5c0-1.103-.897-2-2-2M6 9h4l-.003 9H4V9zm6 8V9c0-1.103-.897-2-2-2H7V5h11v12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Diamond(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.445 3h-8.89c-.345 0-.666.178-.849.47L3.25 9h17.5l-3.456-5.53a1.003 1.003 0 0 0-.849-.47M11.26 21.186a1 1 0 0 0 1.48 0L22 11H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceFive(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 3H5c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2h14c1.103 0 2-.897 2-2V5c0-1.103-.897-2-2-2M8 17.5a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 8 17.5m0-8a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 8 9.5m4 4a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 12 13.5m4 4a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 16 17.5m0-8a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 16 9.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceFour(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 3H5c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2h14c1.103 0 2-.897 2-2V5c0-1.103-.897-2-2-2M8 17.5a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 8 17.5m0-8a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 8 9.5m8 8a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 16 17.5m0-8a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 16 9.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceOne(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 3H5c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2h14c1.103 0 2-.897 2-2V5c0-1.103-.897-2-2-2m-7 10.5a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceSix(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 3H5c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2h14c1.103 0 2-.897 2-2V5c0-1.103-.897-2-2-2M8 17.5a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 8 17.5m0-4a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 8 13.5m0-4a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 8 9.5m8 8a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 16 17.5m0-4a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 16 13.5m0-4a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 16 9.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceThree(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 3H5c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2h14c1.103 0 2-.897 2-2V5c0-1.103-.897-2-2-2M8 9.5a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 8 9.5m4 4a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 12 13.5m4 4a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 16 17.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceTwo(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 3H5c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2h14c1.103 0 2-.897 2-2V5c0-1.103-.897-2-2-2M9.5 13.5a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 9.5 13.5m5 0a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 14.5 13.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DirectionLeft(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.707 11.294l-8.978-9a1.001 1.001 0 0 0-1.415-.002l-9.021 9a1 1 0 0 0 0 1.416l9.021 9c.39.389 1.026.388 1.415-.002l8.978-9a.998.998 0 0 0 0-1.412M15 16h-2v-4h-3v2l-3-3l3-3v2h5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DirectionRight(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m2.295 12.707l8.978 9c.389.39 1.025.391 1.414.002l9.021-9a1 1 0 0 0 0-1.416l-9.021-9a.999.999 0 0 0-1.414.002l-8.978 9a.998.998 0 0 0 0 1.412m6.707-2.706h5v-2l3 3l-3 3v-2h-3v4h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Directions(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 11h-6V8h6a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2H5L2 5l3 3h6v3H5a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h6v5h2v-5h6l3-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Disc(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m0 14c-2.206 0-4-1.794-4-4s1.794-4 4-4s4 1.794 4 4s-1.794 4-4 4"/><circle cx="11.998" cy="11.998" r="2.002" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Discount(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 5H3a1 1 0 0 0-1 1v4h.893c.996 0 1.92.681 2.08 1.664A2.001 2.001 0 0 1 3 14H2v4a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1v-4h-1a2.001 2.001 0 0 1-1.973-2.336c.16-.983 1.084-1.664 2.08-1.664H22V6a1 1 0 0 0-1-1M9 9a1 1 0 1 1 0 2a1 1 0 1 1 0-2m-.8 6.4l6-8l1.6 1.2l-6 8zM15 15a1 1 0 1 1 0-2a1 1 0 1 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dish(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 15c0-4.625-3.507-8.441-8-8.941V4h-2v2.059c-4.493.5-8 4.316-8 8.941v2h18zM2 18h20v2H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dislike(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 3h-1v13h1a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2M4 16h7l-1.122 3.368A2 2 0 0 0 11.775 22H12l5-5.438V3H6l-3.937 8.649l-.063.293V14a2 2 0 0 0 2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dizzy(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2M8 12.414l-1.293 1.293l-1.414-1.414L6.586 11L5.293 9.707l1.414-1.414L8 9.586l1.293-1.293l1.414 1.414L9.414 11l1.293 1.293l-1.414 1.414zM14 18h-4v-2h4zm4.707-5.707l-1.414 1.414L16 12.414l-1.293 1.293l-1.414-1.414L14.586 11l-1.293-1.293l1.414-1.414L16 9.586l1.293-1.293l1.414 1.414L17.414 11z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DockBottom(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 21h14c1.103 0 2-.897 2-2V5c0-1.103-.897-2-2-2H5c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2M19 5l.001 9H5V5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DockLeft(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 19V5c0-1.103-.897-2-2-2H5c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2h14c1.103 0 2-.897 2-2m-11 0V5h9l.002 14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DockRight(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 5c0-1.103-.897-2-2-2H5c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2h14c1.103 0 2-.897 2-2zM5 5h9v14H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DockTop(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 5v14c0 1.103.897 2 2 2h14c1.103 0 2-.897 2-2V5c0-1.103-.897-2-2-2H5c-1.103 0-2 .897-2 2m2 14v-9h14.001l.001 9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dog(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 6h-2l-1.27-1.27A2.49 2.49 0 0 0 16 4h-2.5A2.64 2.64 0 0 0 11 2v6.36a4.38 4.38 0 0 0 1.13 2.72a6.57 6.57 0 0 0 4.13 1.82l3.45-1.38a3 3 0 0 0 1.73-1.84L22 8.15a1.06 1.06 0 0 0 0-.31V7a1 1 0 0 0-1-1m-5 2a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/><path fill="currentColor" d="M11.38 11.74A5.24 5.24 0 0 1 10.07 9H6a1.88 1.88 0 0 1-2-2a1 1 0 0 0-2 0a4.69 4.69 0 0 0 .48 2A3.58 3.58 0 0 0 4 10.53V22h3v-5h6v5h3v-8.13a7.35 7.35 0 0 1-4.62-2.13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DollarCircle(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m1 14.915V18h-2v-1.08c-2.339-.367-3-2.002-3-2.92h2c.011.143.159 1 2 1c1.38 0 2-.585 2-1c0-.324 0-1-2-1c-3.48 0-4-1.88-4-3c0-1.288 1.029-2.584 3-2.915V6.012h2v1.109c1.734.41 2.4 1.853 2.4 2.879h-1l-1 .018C13.386 9.638 13.185 9 12 9c-1.299 0-2 .516-2 1c0 .374 0 1 2 1c3.48 0 4 1.88 4 3c0 1.288-1.029 2.584-3 2.915"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DonateBlood(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.726 13.02L14 16H9v-1h4.065a.5.5 0 0 0 .416-.777l-.888-1.332A1.995 1.995 0 0 0 10.93 12H3a1 1 0 0 0-1 1v6a2 2 0 0 0 2 2h9.639a3 3 0 0 0 2.258-1.024L22 13l-1.452-.484a2.998 2.998 0 0 0-2.822.504M15.403 12a3 3 0 0 0 3-3c0-2.708-3-6-3-6s-3 3.271-3 6a3 3 0 0 0 3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DonateHeart(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.726 13.02L14 16H9v-1h4.065a.5.5 0 0 0 .416-.777l-.888-1.332A1.995 1.995 0 0 0 10.93 12H3a1 1 0 0 0-1 1v6a2 2 0 0 0 2 2h9.639a3 3 0 0 0 2.258-1.024L22 13l-1.452-.484a2.998 2.998 0 0 0-2.822.504m1.532-5.63c.451-.465.73-1.108.73-1.818s-.279-1.353-.73-1.818A2.447 2.447 0 0 0 17.494 3S16.25 2.997 15 4.286C13.75 2.997 12.506 3 12.506 3a2.45 2.45 0 0 0-1.764.753c-.451.466-.73 1.108-.73 1.818s.279 1.354.73 1.818L15 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoorOpen(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 5v14a1 1 0 0 0 1 1h3v-2H7V6h2V4H6a1 1 0 0 0-1 1m14.242-.97l-8-2A1 1 0 0 0 10 3v18a.998.998 0 0 0 1.242.97l8-2A1 1 0 0 0 20 19V5a1 1 0 0 0-.758-.97M15 12.188a1.001 1.001 0 0 1-2 0v-.377a1 1 0 1 1 2 .001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoughnutChart(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 6c2.507.423 4.577 2.493 5 5h4c-.471-4.717-4.283-8.529-9-9z"/><path fill="currentColor" d="M18 13c-.478 2.833-2.982 4.949-5.949 4.949c-3.309 0-6-2.691-6-6C6.051 8.982 8.167 6.478 11 6V2c-5.046.504-8.949 4.773-8.949 9.949c0 5.514 4.486 10 10 10c5.176 0 9.445-3.903 9.949-8.949z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownArrow(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.178 19.569a.998.998 0 0 0 1.644 0l9-13A.999.999 0 0 0 21 5H3a1.002 1.002 0 0 0-.822 1.569z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownArrowAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 12h-5V6h-2v6H6l6 7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownArrowCircle(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m0 15l-5-5h4V7h2v5h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownArrowSquare(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 21a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2zm-8-9V7h2v5h4l-5 5l-5-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Download(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 9h-4V3H9v6H5l7 8zM4 19h16v2H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Downvote(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.901 10.566A1.001 1.001 0 0 0 20 10h-4V3a1 1 0 0 0-1-1H9a1 1 0 0 0-1 1v7H4a1.001 1.001 0 0 0-.781 1.625l8 10a1 1 0 0 0 1.562 0l8-10c.24-.301.286-.712.12-1.059"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Drink(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.832 4.555A1 1 0 0 0 20 3H4a1 1 0 0 0-.832 1.554L11 16.303V20H8v2h8v-2h-3v-3.697zm-2.7.445l-2 3H7.868l-2-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Droplet(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.6 2.4c-.4-.3-.9-.3-1.2 0C9.5 3.9 4 8.5 4 14c0 4.4 3.6 8 8 8s8-3.6 8-8c0-5.4-5.5-10.1-7.4-11.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DropletHalf(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.578 2.184a1.004 1.004 0 0 0-1.156 0C11.119 2.398 4 7.513 4 13.75C4 18.53 7.364 22 12 22s8-3.468 8-8.246c0-6.241-7.119-11.356-7.422-11.57M6 13.75c0-4.283 4.395-8.201 6-9.49V20c-3.533 0-6-2.57-6-6.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dryer(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 22h16a1 1 0 0 0 1-1V5c0-1.654-1.346-3-3-3H6C4.346 2 3 3.346 3 5v16a1 1 0 0 0 1 1M18 3.924a1 1 0 1 1 0 2a1 1 0 0 1 0-2m-3 0a1 1 0 1 1 0 2a1 1 0 0 1 0-2m-3 3.117c3.309 0 6 2.691 6 6s-2.691 6-6 6s-6-2.691-6-6s2.691-6 6-6"/><path fill="currentColor" d="M9.32 12.265c-.415.384-1.041.964-1.041 2.067c0 1.104.626 1.684 1.041 2.068c.352.325.4.398.4.6h2c0-1.104-.626-1.684-1.041-2.068c-.352-.325-.4-.398-.4-.6s.048-.275.4-.6c.414-.384 1.041-.964 1.041-2.068c0-1.103-.626-1.683-1.041-2.066c-.351-.325-.399-.397-.399-.598h-2c0 1.104.627 1.683 1.042 2.066c.351.324.399.396.399.597c-.001.203-.05.276-.401.602m4 0c-.414.384-1.04.964-1.04 2.067s.626 1.684 1.04 2.067c.351.325.399.398.399.601h2c0-1.104-.626-1.684-1.04-2.067c-.351-.325-.399-.398-.399-.601s.049-.275.399-.601c.414-.384 1.04-.964 1.04-2.068c0-1.103-.626-1.682-1.04-2.065c-.35-.324-.399-.397-.399-.598h-2c0 1.103.626 1.683 1.041 2.066c.35.324.398.397.398.598c.001.202-.048.275-.399.601"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Duplicate(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 22h12a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2m2-9h3v-3h2v3h3v2h-3v3H9v-3H6z"/><path fill="currentColor" d="M20 2H8v2h12v12h2V4c0-1.103-.897-2-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Edit(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m18.988 2.012l3 3L19.701 7.3l-3-3zM8 16h3l7.287-7.287l-3-3L8 13z"/><path fill="currentColor" d="M19 19H8.158c-.026 0-.053.01-.079.01c-.033 0-.066-.009-.1-.01H5V5h6.847l2-2H5c-1.103 0-2 .896-2 2v14c0 1.104.897 2 2 2h14a2 2 0 0 0 2-2v-8.668l-2 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16 2.012l3 3L16.713 7.3l-3-3zM4 14v3h3l8.299-8.287l-3-3zm0 6h16v2H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditLocation(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C7.589 2 4 5.589 4 9.995C3.971 16.44 11.696 21.784 12 22c0 0 8.029-5.56 8-12c0-4.411-3.589-8-8-8M9.799 14.987H8v-1.799l4.977-4.97l1.799 1.799zm5.824-5.817l-1.799-1.799L15.196 6l1.799 1.799z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eject(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12 6l-6 8h12zM6 16h12v2H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Envelope(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 4H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2m0 4.7l-8 5.334L4 8.7V6.297l8 5.333l8-5.333z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeOpen(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.555 8.168l-9-6a1 1 0 0 0-1.109 0l-9 6A.995.995 0 0 0 2.004 9H2v11a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V9h-.004a.997.997 0 0 0-.441-.832M20 12.7L12 17l-8-4.3v-2.403l8 4.299l8-4.299z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eraser(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.48 3L7.73 7.75L3 12.59a2 2 0 0 0 0 2.82l4.3 4.3A1 1 0 0 0 8 20h12v-2h-7l7.22-7.22a2 2 0 0 0 0-2.83L15.31 3a2 2 0 0 0-2.83 0M8.41 18l-4-4l4.75-4.84l.74-.75l4.95 4.95l-4.56 4.56l-.07.08z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ErrorAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.707 2.293A.996.996 0 0 0 16 2H8a.996.996 0 0 0-.707.293l-5 5A.996.996 0 0 0 2 8v8c0 .266.105.52.293.707l5 5A.996.996 0 0 0 8 22h8c.266 0 .52-.105.707-.293l5-5A.996.996 0 0 0 22 16V8a.996.996 0 0 0-.293-.707zM13 17h-2v-2h2zm0-4h-2V7h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ErrorCircle(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.953 2C6.465 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.493 2 11.953 2M13 17h-2v-2h2zm0-4h-2V7h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ErrorIcon(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.884 2.532c-.346-.654-1.422-.654-1.768 0l-9 17A.999.999 0 0 0 3 21h18a.998.998 0 0 0 .883-1.467zM13 18h-2v-2h2zm-2-4V9h2l.001 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EvStation(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m19.616 6.48l.014-.017l-4-3.24l-1.26 1.554l2.067 1.674a2.99 2.99 0 0 0-1.394 3.062c.15.899.769 1.676 1.57 2.111c.895.487 1.68.442 2.378.194L18.976 18a.996.996 0 0 1-1.39.922a.995.995 0 0 1-.318-.217a.996.996 0 0 1-.291-.705L17 16a2.98 2.98 0 0 0-.877-2.119A3 3 0 0 0 14 13h-1V5a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h7a2 2 0 0 0 2-2v-4h1c.136 0 .267.027.391.078a1.028 1.028 0 0 1 .531.533A.994.994 0 0 1 15 16l-.024 2c0 .406.079.799.236 1.168c.151.359.368.68.641.951a2.97 2.97 0 0 0 2.123.881c.406 0 .798-.078 1.168-.236c.358-.15.68-.367.951-.641A2.983 2.983 0 0 0 20.976 18L21 9a2.997 2.997 0 0 0-1.384-2.52M6 18l1-5H4l5-7l-1 5h3zm12-8a1 1 0 1 1 0-2a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Exit(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.002 21h14c1.103 0 2-.897 2-2V5c0-1.103-.897-2-2-2h-14c-1.103 0-2 .897-2 2v6.001H10V7l6 5l-6 5v-3.999H3.002V19c0 1.103.897 2 2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Extension(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 10V7c0-1.103-.897-2-2-2h-3c0-1.654-1.346-3-3-3S8 3.346 8 5H5c-1.103 0-2 .897-2 2v4h1a2 2 0 0 1 0 4H3v4c0 1.103.897 2 2 2h4v-1a2 2 0 0 1 4 0v1h4c1.103 0 2-.897 2-2v-3c1.654 0 3-1.346 3-3s-1.346-3-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eyedropper(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m4 15.76l-1 4A1 1 0 0 0 3.75 21a1 1 0 0 0 .49 0l4-1a1 1 0 0 0 .47-.26L17 11.41l1.29 1.3l1.42-1.42l-1.3-1.29L21 7.41a2 2 0 0 0 0-2.82L19.41 3a2 2 0 0 0-2.82 0L14 5.59l-1.3-1.3l-1.42 1.42L12.58 7l-8.29 8.29a1 1 0 0 0-.29.47m1.87.75L14 8.42L15.58 10l-8.09 8.1l-2.12.53z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Face(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m0 18c-4.411 0-8-3.589-8-8c0-1.168.258-2.275.709-3.276c.154.09.308.182.456.276c.396.25.791.5 1.286.688c.494.187 1.088.312 1.879.312c.792 0 1.386-.125 1.881-.313s.891-.437 1.287-.687s.792-.5 1.287-.688c.494-.187 1.088-.312 1.88-.312s1.386.125 1.88.313c.495.187.891.437 1.287.687s.792.5 1.287.688c.178.067.374.122.581.171c.191.682.3 1.398.3 2.141c0 4.411-3.589 8-8 8"/><circle cx="8.5" cy="12.5" r="1.5" fill="currentColor"/><circle cx="15.5" cy="12.5" r="1.5" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceMask(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12c0 2.521.945 4.82 2.49 6.582c1.24 1.52 3.266 3.066 6.439 3.358a9.731 9.731 0 0 0 2.141 0c3.174-.292 5.199-1.839 6.439-3.358A9.948 9.948 0 0 0 22 12c0-5.514-4.486-10-10-10M4.709 8.724c.154.09.308.182.456.276c.396.25.791.5 1.286.688c.494.187 1.088.312 1.879.312c.792 0 1.386-.125 1.881-.313s.891-.437 1.287-.687s.792-.5 1.287-.688c.494-.187 1.088-.312 1.88-.312s1.386.125 1.88.313c.495.187.891.437 1.287.687s.792.5 1.287.688c.178.067.374.122.581.171c.191.682.3 1.398.3 2.141c0 .843-.133 1.654-.375 2.417c-.261.195-.733.474-1.577.756c-.769.256-1.672.458-2.685.602a25.337 25.337 0 0 1-6.727 0c-1.013-.144-1.916-.346-2.685-.602c-.844-.282-1.316-.561-1.577-.756a7.953 7.953 0 0 1 .335-5.693"/><circle cx="8.5" cy="12.5" r="1.5" fill="currentColor"/><circle cx="15.5" cy="12.5" r="1.5" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Factory(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 10V6l-5 4V6l-5 4V4H2v16h20V6zm-8 7H7v-3h2zm5 0h-2v-3h2zm5 0h-2v-3h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastForwardCircle(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.485 2 2 6.485 2 12s4.485 10 10 10c5.514 0 10-4.485 10-10S17.514 2 12 2m1 14v-4l-6 4V8l6 4V8l6 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func File(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 22a2 2 0 0 0 2-2V8l-6-6H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2zM13 4l5 5h-5zM7 8h3v2H7zm0 4h10v2H7zm0 4h10v2H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileArchive(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 2a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8l-6-6h-3v2H9v2h2v2H9v2h2v8H7v-6h2v-2H7V8h2V6H7V4h2V2zm7 2l5 5h-5z"/><path fill="currentColor" d="M8 15h2v2H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileBlank(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 2a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8l-6-6zm8 7h-1V4l5 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileCss(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8zM9.446 17.412c.275 0 .581-.061.762-.132l.138.713c-.168.084-.546.174-1.037.174c-1.397 0-2.117-.869-2.117-2.021C7.191 14.768 8.175 14 9.398 14c.474 0 .833.096.995.18l-.186.726a1.98 1.98 0 0 0-.768-.149c-.726 0-1.29.438-1.29 1.337c.001.808.482 1.318 1.297 1.318m2.491.755c-.461 0-.917-.119-1.145-.245l.186-.756c.246.126.624.252 1.014.252c.42 0 .642-.174.642-.438c0-.252-.192-.396-.678-.57c-.672-.234-1.109-.605-1.109-1.193c0-.689.575-1.217 1.529-1.217c.455 0 .791.096 1.031.203l-.204.738a1.919 1.919 0 0 0-.846-.192c-.396 0-.587.181-.587.39c0 .258.228.372.749.57c.714.264 1.05.636 1.05 1.205c-.001.678-.523 1.253-1.632 1.253m3.24 0c-.461 0-.917-.119-1.145-.245l.186-.756c.246.126.624.252 1.014.252c.42 0 .642-.174.642-.438c0-.252-.192-.396-.678-.57c-.672-.234-1.109-.605-1.109-1.193c0-.689.575-1.217 1.529-1.217c.455 0 .791.096 1.031.203l-.204.738a1.919 1.919 0 0 0-.846-.192c-.396 0-.587.181-.587.39c0 .258.228.372.749.57c.714.264 1.05.636 1.05 1.205c0 .678-.523 1.253-1.632 1.253M14 9h-1V4l5 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileDoc(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.186 14.552c-.617 0-.977.587-.977 1.373c0 .791.371 1.35.983 1.35c.617 0 .971-.588.971-1.374c0-.726-.348-1.349-.977-1.349"/><path fill="currentColor" d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8zM9.155 17.454c-.426.354-1.073.521-1.864.521c-.475 0-.81-.03-1.038-.06v-3.971a8.16 8.16 0 0 1 1.235-.083c.768 0 1.266.138 1.655.432c.42.312.684.81.684 1.522c0 .775-.282 1.309-.672 1.639m2.99.546c-1.2 0-1.901-.906-1.901-2.058c0-1.211.773-2.116 1.967-2.116c1.241 0 1.919.929 1.919 2.045c-.001 1.325-.805 2.129-1.985 2.129m4.655-.762c.275 0 .581-.061.762-.132l.138.713c-.168.084-.546.174-1.037.174c-1.397 0-2.117-.869-2.117-2.021c0-1.379.983-2.146 2.207-2.146c.474 0 .833.096.995.18l-.186.726a1.979 1.979 0 0 0-.768-.15c-.726 0-1.29.438-1.29 1.338c0 .809.48 1.318 1.296 1.318M14 9h-1V4l5 5z"/><path fill="currentColor" d="M7.584 14.563c-.203 0-.335.018-.413.036v2.645c.078.018.204.018.317.018c.828.006 1.367-.449 1.367-1.415c.006-.84-.485-1.284-1.271-1.284"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileExport(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 22a2 2 0 0 0 2-2v-5l-5 4v-3H8v-2h7v-3l5 4V8l-6-6H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2zM13 4l5 5h-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileFind(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 22h12c.178 0 .348-.03.512-.074l-3.759-3.759A4.966 4.966 0 0 1 12 19c-2.757 0-5-2.243-5-5s2.243-5 5-5s5 2.243 5 5a4.964 4.964 0 0 1-.833 2.753l3.759 3.759c.044-.164.074-.334.074-.512V8l-6-6H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2"/><circle cx="12" cy="14" r="3" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileGif(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8zm-2.667 15.772A4.512 4.512 0 0 1 9.984 18c-.737 0-1.271-.186-1.644-.546c-.371-.348-.575-.875-.569-1.469c.006-1.344.983-2.111 2.309-2.111c.521 0 .924.103 1.121.198l-.191.731c-.222-.096-.498-.174-.941-.174c-.762 0-1.338.432-1.338 1.308c0 .833.522 1.325 1.271 1.325c.21 0 .378-.024.45-.061v-.846h-.624v-.713h1.505zm1.634.186h-.918v-4.042h.918zm3.262-3.292h-1.553v.923h1.451v.744h-1.451v1.625h-.918v-4.042h2.471zM14 9h-1V4l5 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileHtml(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8zm-1 2l5 5h-5zM8.531 18h-.76v-1.411H6.515V18h-.767v-3.373h.767v1.296h1.257v-1.296h.76zm3-2.732h-.921V18h-.766v-2.732h-.905v-.641h2.592zM14.818 18l-.05-1.291c-.017-.405-.03-.896-.03-1.387h-.016c-.104.431-.245.911-.375 1.307l-.41 1.316h-.597l-.359-1.307a15.154 15.154 0 0 1-.306-1.316h-.011c-.021.456-.034.976-.059 1.396L12.545 18h-.705l.216-3.373h1.015l.331 1.126c.104.391.21.811.284 1.206h.017c.095-.391.209-.836.32-1.211l.359-1.121h.996L15.563 18zm3.434 0h-2.108v-3.373h.767v2.732h1.342z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileImage(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 22h12a2 2 0 0 0 2-2V8l-6-6H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2m7-18l5 5h-5zm-4.5 7a1.5 1.5 0 1 1-.001 3.001A1.5 1.5 0 0 1 8.5 11m.5 5l1.597 1.363L13 13l4 6H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileImport(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 14V8l-6-6H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-4h-7v3l-5-4l5-4v3zM13 4l5 5h-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileJpg(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8zM9.239 16.446c0 1.152-.551 1.554-1.438 1.554c-.21 0-.486-.036-.665-.097l.101-.737c.127.042.289.072.469.072c.384 0 .623-.174.623-.804v-2.543h.911zm3.294-.365c-.313.293-.773.426-1.313.426c-.12 0-.228-.007-.312-.019v1.445h-.906v-3.988a7.528 7.528 0 0 1 1.236-.083c.563 0 .965.107 1.234.323c.259.204.433.54.433.936s-.133.732-.372.96m4.331 1.667c-.28.096-.815.228-1.349.228c-.737 0-1.271-.186-1.643-.546c-.371-.348-.575-.875-.57-1.469c.007-1.344.983-2.111 2.309-2.111c.521 0 .924.103 1.121.198l-.191.731c-.222-.096-.497-.174-.941-.174c-.761 0-1.338.432-1.338 1.308c0 .833.523 1.325 1.271 1.325c.211 0 .378-.024.451-.061v-.846h-.624v-.713h1.504zM14 9h-1V4l5 5z"/><path fill="currentColor" d="M11.285 14.552c-.186 0-.312.018-.377.036v1.193c.077.018.174.023.307.023c.484 0 .784-.246.784-.659c0-.372-.257-.593-.714-.593"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileJs(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8zm-2.745 14.186c0 1.345-.644 1.814-1.681 1.814c-.245 0-.567-.042-.777-.112l.119-.861c.147.049.336.084.546.084c.448 0 .729-.203.729-.938v-2.97h1.064zm2.043 1.807c-.539 0-1.071-.141-1.337-.287l.217-.883c.287.147.729.294 1.184.294c.491 0 .749-.203.749-.511c0-.295-.224-.463-.791-.666c-.784-.272-1.295-.707-1.295-1.394c0-.806.672-1.422 1.786-1.422c.533 0 .925.112 1.205.238l-.238.861c-.189-.091-.525-.224-.987-.224s-.687.21-.687.455c0 .301.267.435.875.665c.834.309 1.226.742 1.226 1.408c-.002.793-.61 1.466-1.907 1.466M14 9h-1V4l5 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileJson(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.823 15.122c-.517 0-.816.491-.816 1.146c0 .661.311 1.126.82 1.126c.517 0 .812-.49.812-1.146c0-.604-.291-1.126-.816-1.126"/><path fill="currentColor" d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8zM8.022 16.704c0 .961-.461 1.296-1.2 1.296c-.176 0-.406-.029-.557-.08l.086-.615c.104.035.239.06.391.06c.319 0 .52-.145.52-.67v-2.122h.761zm1.459 1.291c-.385 0-.766-.1-.955-.205l.155-.631c.204.105.521.211.846.211c.35 0 .534-.146.534-.365c0-.211-.159-.331-.564-.476c-.562-.195-.927-.506-.927-.996c0-.576.481-1.017 1.277-1.017c.38 0 .659.08.861.171l-.172.615c-.135-.065-.375-.16-.705-.16s-.491.15-.491.325c0 .215.19.311.627.476c.596.22.876.53.876 1.006c.001.566-.436 1.046-1.362 1.046m3.306.005c-1.001 0-1.586-.755-1.586-1.716c0-1.012.646-1.768 1.642-1.768c1.035 0 1.601.776 1.601 1.707C14.443 17.33 13.773 18 12.787 18m4.947-.055h-.802l-.721-1.302a12.64 12.64 0 0 1-.585-1.19l-.016.005c.021.445.031.921.031 1.472v1.016h-.701v-3.373h.891l.701 1.236c.2.354.4.775.552 1.155h.014c-.05-.445-.065-.9-.065-1.406v-.985h.702zM14 9h-1V4l5 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileMd(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.481 14.015c-.238 0-.393.021-.483.042v3.089c.091.021.237.021.371.021c.966.007 1.597-.525 1.597-1.653c.007-.981-.568-1.499-1.485-1.499"/><path fill="currentColor" d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8zm-2.934 15.951l-.07-1.807a53.142 53.142 0 0 1-.042-1.94h-.021a26.098 26.098 0 0 1-.525 1.828l-.574 1.842H9l-.504-1.828a21.996 21.996 0 0 1-.428-1.842h-.013c-.028.638-.049 1.366-.084 1.954l-.084 1.793h-.988L7.2 13.23h1.422l.462 1.576c.147.546.295 1.135.399 1.688h.021a39.87 39.87 0 0 1 .448-1.694l.504-1.569h1.394l.26 4.721h-1.044zm5.25-.56c-.498.413-1.253.609-2.178.609a9.27 9.27 0 0 1-1.212-.07v-4.636a9.535 9.535 0 0 1 1.443-.099c.896 0 1.478.161 1.933.505c.49.364.799.945.799 1.778c0 .904-.33 1.528-.785 1.913M14 9h-1V4l5 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilePdf(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.267 14.68c-.184 0-.308.018-.372.036v1.178c.076.018.171.023.302.023c.479 0 .774-.242.774-.651c0-.366-.254-.586-.704-.586m3.487.012c-.2 0-.33.018-.407.036v2.61c.077.018.201.018.313.018c.817.006 1.349-.444 1.349-1.396c.006-.83-.479-1.268-1.255-1.268"/><path fill="currentColor" d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8zM9.498 16.19c-.309.29-.765.42-1.296.42a2.23 2.23 0 0 1-.308-.018v1.426H7v-3.936A7.558 7.558 0 0 1 8.219 14c.557 0 .953.106 1.22.319c.254.202.426.533.426.923c-.001.392-.131.723-.367.948m3.807 1.355c-.42.349-1.059.515-1.84.515c-.468 0-.799-.03-1.024-.06v-3.917A7.947 7.947 0 0 1 11.66 14c.757 0 1.249.136 1.633.426c.415.308.675.799.675 1.504c0 .763-.279 1.29-.663 1.615M17 14.77h-1.532v.911H16.9v.734h-1.432v1.604h-.906V14.03H17zM14 9h-1V4l5 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilePlus(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 22h12a2 2 0 0 0 2-2V8l-6-6H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2m7-18l5 5h-5zM8 14h3v-3h2v3h3v2h-3v3h-2v-3H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilePng(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.782 14.576c-.186 0-.312.018-.377.036v1.193c.077.018.174.023.306.023c.485 0 .785-.246.785-.659c0-.371-.258-.593-.714-.593"/><path fill="currentColor" d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8zM9.03 16.105c-.313.293-.774.426-1.313.426c-.12 0-.229-.007-.312-.019v1.445h-.906V13.97a7.504 7.504 0 0 1 1.235-.083c.563 0 .966.107 1.235.323c.258.204.432.54.432.936s-.131.731-.371.959m4.302 1.853h-.96l-.863-1.56c-.24-.432-.504-.953-.701-1.427l-.019.006c.024.534.036 1.104.036 1.763v1.218h-.84v-4.042h1.067l.84 1.481c.24.426.479.93.659 1.385h.019a14.746 14.746 0 0 1-.078-1.685v-1.182h.84zm4.169-.186a4.512 4.512 0 0 1-1.349.228c-.737 0-1.271-.186-1.644-.546c-.371-.348-.575-.875-.569-1.469c.006-1.344.983-2.111 2.309-2.111c.521 0 .924.103 1.121.198l-.191.731c-.222-.096-.498-.174-.941-.174c-.762 0-1.338.432-1.338 1.308c0 .833.522 1.325 1.271 1.325c.21 0 .378-.024.45-.061v-.846h-.624v-.713h1.505zM14 9h-1V4l5 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileTxt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8zM9.998 14.768H8.895v3.274h-.917v-3.274H6.893V14h3.105zm2.725 3.274l-.365-.731c-.15-.282-.246-.492-.359-.726h-.013c-.083.233-.185.443-.312.726l-.335.731h-1.045l1.171-2.045L10.336 14h1.05l.354.738c.121.245.21.443.306.671h.013c.096-.258.174-.438.276-.671l.341-.738h1.043l-1.139 1.973l1.198 2.069zm4.384-3.274h-1.104v3.274h-.917v-3.274h-1.085V14h3.105zM14 9h-1V4l5 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Film(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 4v1h-2V3H7v2H5V3H3v18h2v-2h2v2h10v-2h2v2h2V3h-2zM5 7h2v2H5zm0 4h2v2H5zm0 6v-2h2v2zm12 0v-2h2v2zm2-4h-2v-2h2zm-2-4V7h2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 20v-4.586L20.414 8c.375-.375.586-.884.586-1.415V4a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v2.585c0 .531.211 1.04.586 1.415L11 15.414V22z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FirstAid(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 6h-3V4a2 2 0 0 0-2-2H9a2 2 0 0 0-2 2v2H4a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2M9 4h6v2H9zm7 10h-3v3h-2v-3H8v-2h3V9h2v3h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flag(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 4H6V2H4v18H3v2h4v-2H6v-5h13a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlagAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14.303 6l-3-2H6V2H4v20h2v-8h4.697l3 2H20V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlagCheckered(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 2v18H3v2h4v-2H6v-5h13a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1H6V2zm4 3v2h2V5h2v2h2V5h2v2h2v2h-2v2h2v2h-2v-2h-2v2h-2v-2h-2v2H8v-2H6V9h2V7H6V5z"/><path fill="currentColor" d="M8 9h2v2H8zm4 0h2v2h-2zm-2-2h2v2h-2zm4 0h2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flame(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.579 2.393a.982.982 0 0 0-1.153.006C9.592 3.728 4 8.252 4 14c0 3.247 1.948 6.043 4.734 7.296A3.971 3.971 0 0 1 8 19c-.017-3.221 3.558-6.893 3.71-7a.497.497 0 0 1 .579 0c.152.107 3.711 2.974 3.711 7.002c0 .854-.275 1.643-.733 2.294C18.052 20.043 20 17.248 20 14.005c0-5.861-5.582-10.307-7.421-11.612"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flask(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 9.783V4h1V2H8v2h1v5.783l-4.268 9.389a1.992 1.992 0 0 0 .14 1.911A1.99 1.99 0 0 0 6.553 22h10.895a1.99 1.99 0 0 0 1.681-.917c.37-.574.423-1.289.14-1.911zm-4.09.631c.06-.13.09-.271.09-.414V4h2v6c0 .143.03.284.09.414L15.177 15H8.825z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Florist(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.84 21.871L12 22a10.221 10.221 0 0 0-9.013-8.891L2 13l.021.173a10.001 10.001 0 0 0 8.819 8.698m11.139-8.698L22 13l-.987.109c-4.7.523-8.427 4.2-9.013 8.891l1.16-.129a10.001 10.001 0 0 0 8.819-8.698M18.063 5.5a2.5 2.5 0 0 0-3.415-.915c-.062.035-.111.081-.168.121c.005-.069.02-.136.02-.206a2.5 2.5 0 1 0-5 0c0 .07.015.137.021.206c-.057-.04-.107-.086-.168-.121a2.5 2.5 0 0 0-2.5 4.33c.061.035.126.056.188.085c-.062.029-.127.05-.188.085a2.5 2.5 0 0 0 2.5 4.33c.062-.035.111-.081.168-.121c-.006.069-.021.136-.021.206a2.5 2.5 0 1 0 5 0c0-.07-.015-.137-.021-.206c.057.04.106.086.168.121a2.5 2.5 0 0 0 2.5-4.33c-.061-.035-.126-.056-.188-.085c.063-.029.127-.05.188-.085a2.5 2.5 0 0 0 .916-3.415M12 12a3 3 0 1 1 0-6a3 3 0 0 1 0 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folder(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 5h-9.586L8.707 3.293A.997.997 0 0 0 8 3H4c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2h16c1.103 0 2-.897 2-2V7c0-1.103-.897-2-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderMinus(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 5h-9.586L8.707 3.293A.997.997 0 0 0 8 3H4c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2h16c1.103 0 2-.897 2-2V7c0-1.103-.897-2-2-2m-4 9H8v-2h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOpen(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.165 19.551c.186.28.499.449.835.449h15c.4 0 .762-.238.919-.606l3-7A.998.998 0 0 0 21 11h-1V8c0-1.103-.897-2-2-2h-6.655L8.789 4H4c-1.103 0-2 .897-2 2v13h.007a1 1 0 0 0 .158.551M18 8v3H6c-.4 0-.762.238-.919.606L4 14.129V8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderPlus(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 5h-9.586L8.707 3.293A.997.997 0 0 0 8 3H4c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2h16c1.103 0 2-.897 2-2V7c0-1.103-.897-2-2-2m-4 9h-3v3h-2v-3H8v-2h3V9h2v3h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FoodMenu(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 2h2v20H3zm16 0H6v20h13c1.103 0 2-.897 2-2V4c0-1.103-.897-2-2-2m-1 10H9v-2h9zm0-4H9V6h9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fridge(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 2H6c-1.103 0-2 .897-2 2v5h4V6h2v3h10V4c0-1.103-.897-2-2-2m-8 13H8v-5H4v10c0 1.103.897 2 2 2h12c1.103 0 2-.897 2-2V10H10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Game(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22c3.719 0 7.063-2.035 8.809-5.314L13 12l7.809-4.686C19.063 4.035 15.719 2 12 2C6.486 2 2 6.486 2 12s4.486 10 10 10m-.5-16a1.5 1.5 0 1 1-.001 3.001A1.5 1.5 0 0 1 11.5 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GasPump(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m19.616 6.48l.014-.017l-4-3.24l-1.26 1.554l2.067 1.674a2.99 2.99 0 0 0-1.394 3.062c.15.899.769 1.676 1.57 2.111c.895.487 1.68.442 2.378.194L18.976 18a.996.996 0 0 1-1.39.922a.995.995 0 0 1-.318-.217a.996.996 0 0 1-.291-.705L17 16a2.98 2.98 0 0 0-.877-2.119A3 3 0 0 0 14 13h-1V5a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h7a2 2 0 0 0 2-2v-4h1c.136 0 .267.027.391.078a1.028 1.028 0 0 1 .531.533A.994.994 0 0 1 15 16l-.024 2c0 .406.079.799.236 1.168c.151.359.368.68.641.951a2.97 2.97 0 0 0 2.123.881c.406 0 .798-.078 1.168-.236c.358-.15.68-.367.951-.641A2.983 2.983 0 0 0 20.976 18L21 9a2.997 2.997 0 0 0-1.384-2.52M11 8H4V5h7zm7 2a1 1 0 1 1 0-2a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ghost(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 11v8h.051c.245 1.692 1.69 3 3.449 3c1.174 0 2.074-.417 2.672-1.174a3.99 3.99 0 0 0 5.668-.014c.601.762 1.504 1.188 2.66 1.188c1.93 0 3.5-1.57 3.5-3.5V11c0-4.962-4.037-9-9-9s-9 4.038-9 9m6 1c-1.103 0-2-.897-2-2s.897-2 2-2s2 .897 2 2s-.897 2-2 2m6-4c1.103 0 2 .897 2 2s-.897 2-2 2s-2-.897-2-2s.897-2 2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gift(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 12H4v8a2 2 0 0 0 2 2h5V12zm13 0h-5v10h5a2 2 0 0 0 2-2v-8zm.791-5A4.92 4.92 0 0 0 19 5.5C19 3.57 17.43 2 15.5 2c-1.622 0-2.705 1.482-3.404 3.085C11.407 3.57 10.269 2 8.5 2C6.57 2 5 3.57 5 5.5c0 .596.079 1.089.209 1.5H2v4h9V9h2v2h9V7zM7 5.5C7 4.673 7.673 4 8.5 4c.888 0 1.714 1.525 2.198 3H8c-.374 0-1 0-1-1.5M15.5 4c.827 0 1.5.673 1.5 1.5C17 7 16.374 7 16 7h-2.477c.51-1.576 1.251-3 1.977-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Graduation(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 7v1l11 4l9-4V7L11 4z"/><path fill="currentColor" d="M4 11v4.267c0 1.621 4.001 3.893 9 3.734c4-.126 6.586-1.972 7-3.467c.024-.089.037-.178.037-.268V11L13 14l-5-1.667v3.213l-1-.364V12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Grid(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 4h4v4H4zm6 0h4v4h-4zm6 0h4v4h-4zM4 10h4v4H4zm6 0h4v4h-4zm6 0h4v4h-4zM4 16h4v4H4zm6 0h4v4h-4zm6 0h4v4h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GridAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 11h6a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1m10 0h6a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1h-6a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1M4 21h6a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1m10 0h6a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1h-6a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GroupIcon(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.5 12c2.206 0 4-1.794 4-4s-1.794-4-4-4s-4 1.794-4 4s1.794 4 4 4m1.5 1H8c-3.309 0-6 2.691-6 6v1h15v-1c0-3.309-2.691-6-6-6"/><path fill="currentColor" d="M16.604 11.048a5.67 5.67 0 0 0 .751-3.44c-.179-1.784-1.175-3.361-2.803-4.44l-1.105 1.666c1.119.742 1.8 1.799 1.918 2.974a3.693 3.693 0 0 1-1.072 2.986l-1.192 1.192l1.618.475C18.951 13.701 19 17.957 19 18h2c0-1.789-.956-5.285-4.396-6.952"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GuitarAmp(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 6h-2V4c0-1.103-.897-2-2-2H8c-1.103 0-2 .897-2 2v2H4c-1.103 0-2 .897-2 2v12c0 1.103.897 2 2 2h16c1.103 0 2-.897 2-2V8c0-1.103-.897-2-2-2M8 4h8v2H8zM6 19a1 1 0 1 1 0-2a1 1 0 0 1 0 2m0-3a1 1 0 1 1 0-2a1 1 0 0 1 0 2m3 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2m0-3a1 1 0 1 1 0-2a1 1 0 0 1 0 2m3 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2m0-3a1 1 0 1 1 0-2a1 1 0 0 1 0 2m3 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2m0-3a1 1 0 1 1 0-2a1 1 0 0 1 0 2m3 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2m0-3a1 1 0 1 1 0-2a1 1 0 0 1 0 2m2-4H4V8h16z"/><path fill="currentColor" d="M14 9h2v2h-2zm3 0h2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hand(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.5 5A1.5 1.5 0 0 0 19 6.5V11h-1V4.5a1.5 1.5 0 0 0-3 0V11h-1V3.5a1.5 1.5 0 0 0-3 0V11h-1V5.5a1.5 1.5 0 0 0-3 0v10.81l-2.22-3.6a1.5 1.5 0 0 0-2.56 1.58l3.31 5.34A5 5 0 0 0 9.78 22H17a5 5 0 0 0 5-5V6.5A1.5 1.5 0 0 0 20.5 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandDown(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.278 2.231a1.001 1.001 0 0 0-.64-.231H5a2 2 0 0 0-2 2v7.21a2 2 0 0 0 1.779 1.987L12 14v6a2 2 0 0 0 4 0V8l3.03 1.212a2.001 2.001 0 0 0 2.641-1.225l.113-.34a.998.998 0 0 0-.309-1.084z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandLeft(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 3h-7.21a2 2 0 0 0-1.987 1.779L10 12H4a2 2 0 0 0 0 4h12l-1.212 3.03a2.001 2.001 0 0 0 1.225 2.641l.34.113a.998.998 0 0 0 1.084-.309l4.332-5.197c.149-.179.231-.406.231-.64V5a2 2 0 0 0-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandRight(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 8H8l1.212-3.03a2 2 0 0 0-1.225-2.641l-.34-.113a.998.998 0 0 0-1.084.309L2.231 7.722a1.001 1.001 0 0 0-.231.64V19a2 2 0 0 0 2 2h7.21a2 2 0 0 0 1.987-1.779L14 12h6a2 2 0 0 0 0-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandUp(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.221 10.803L12 10V4a2 2 0 0 0-4 0v12l-3.031-1.212a2 2 0 0 0-2.64 1.225l-.113.34a.998.998 0 0 0 .309 1.084l5.197 4.332c.179.149.406.231.64.231H19a2 2 0 0 0 2-2v-7.21a2 2 0 0 0-1.779-1.987"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Happy(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m3.493 7a1.494 1.494 0 1 1-.001 2.987A1.494 1.494 0 0 1 15.493 9M8.5 9a1.5 1.5 0 1 1-.001 3.001A1.5 1.5 0 0 1 8.5 9m3.5 9c-4 0-5-4-5-4h10s-1 4-5 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HappyAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m3.493 6a1.494 1.494 0 1 1-.001 2.987A1.494 1.494 0 0 1 15.493 8M8.5 8a1.5 1.5 0 1 1-.001 3.001A1.5 1.5 0 0 1 8.5 8M12 18c-5 0-6-5-6-5h12s-1 5-6 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HappyBeaming(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2M8 9c2.201 0 3 1.794 3 3H9c-.012-.45-.194-1-1-1s-.988.55-1 1.012L5 12c0-1.206.799-3 3-3m4 9c-4 0-5-4-5-4h10s-1 4-5 4m5-6c-.012-.45-.194-1-1-1s-.988.55-1 1.012L13 12c0-1.206.799-3 3-3s3 1.794 3 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HappyHeartEyes(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2M6.435 8.467A1.49 1.49 0 0 1 8.502 8.4a1.49 1.49 0 0 1 2.065.033c.597.592.604 1.521.018 2.118l-2.05 2.083l-2.082-2.05a1.484 1.484 0 0 1-.018-2.117M12 18c-4 0-5-4-5-4h10s-1 4-5 4m5.585-7.449l-2.05 2.083l-2.082-2.05a1.485 1.485 0 0 1-.019-2.117a1.49 1.49 0 0 1 2.068-.067a1.49 1.49 0 0 1 2.065.033c.597.591.605 1.521.018 2.118"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HardHat(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 18v-3a8 8 0 0 0-5-7.4V13h-1V5h-4v8H9V7.6A8 8 0 0 0 4 15v3H2v2h20v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hdd(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 13H4a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2m-4 5h-2v-2h2zm4 0h-2v-2h2zm.775-7H21l-1.652-7.434A2 2 0 0 0 17.396 2H6.604a2 2 0 0 0-1.952 1.566L3 11z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heart(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.205 4.791a5.938 5.938 0 0 0-4.209-1.754A5.906 5.906 0 0 0 12 4.595a5.904 5.904 0 0 0-3.996-1.558a5.942 5.942 0 0 0-4.213 1.758c-2.353 2.363-2.352 6.059.002 8.412L12 21.414l8.207-8.207c2.354-2.353 2.355-6.049-.002-8.416"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartCircle(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m4.186 10.74L12 16.926L7.814 12.74a2.745 2.745 0 0 1 0-3.907a2.745 2.745 0 0 1 3.906 0l.28.279l.279-.279a2.745 2.745 0 0 1 3.906 0a2.745 2.745 0 0 1 .001 3.907"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartSquare(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 21h16a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1M7.812 8.907a2.746 2.746 0 0 1 3.907 0l.279.279l.278-.279a2.746 2.746 0 0 1 3.907 0a2.745 2.745 0 0 1 0 3.907L11.998 17l-4.187-4.186a2.747 2.747 0 0 1 .001-3.907"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HelpCircle(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m1 16h-2v-2h2zm.976-4.885c-.196.158-.385.309-.535.459c-.408.407-.44.777-.441.793v.133h-2v-.167c0-.118.029-1.177 1.026-2.174c.195-.195.437-.393.691-.599c.734-.595 1.216-1.029 1.216-1.627a1.934 1.934 0 0 0-3.867.001h-2C8.066 7.765 9.831 6 12 6s3.934 1.765 3.934 3.934c0 1.597-1.179 2.55-1.958 3.181"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hide(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.073 12.194L4.212 8.333c-1.52 1.657-2.096 3.317-2.106 3.351L2 12l.105.316C2.127 12.383 4.421 19 12.054 19c.929 0 1.775-.102 2.552-.273l-2.746-2.746a3.987 3.987 0 0 1-3.787-3.787M12.054 5c-1.855 0-3.375.404-4.642.998L3.707 2.293L2.293 3.707l18 18l1.414-1.414l-3.298-3.298c2.638-1.953 3.579-4.637 3.593-4.679l.105-.316l-.105-.316C21.98 11.617 19.687 5 12.054 5m1.906 7.546c.187-.677.028-1.439-.492-1.96s-1.283-.679-1.96-.492L10 8.586A3.955 3.955 0 0 1 12.054 8c2.206 0 4 1.794 4 4a3.94 3.94 0 0 1-.587 2.053z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Home(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.743 12.331l-9-10c-.379-.422-1.107-.422-1.486 0l-9 10a.998.998 0 0 0-.17 1.076c.16.361.518.593.913.593h2v7a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1v-4h4v4a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1v-7h2a.998.998 0 0 0 .743-1.669"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeAltTwo(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.74 2.32a1 1 0 0 0-1.48 0l-9 10A1 1 0 0 0 3 14h2v7a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-7h2a1 1 0 0 0 1-1a1 1 0 0 0-.26-.68z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeCircle(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.743 12.331l-9-10c-.379-.422-1.107-.422-1.486 0l-9 10A1 1 0 0 0 3 14h2v7a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-7h2a.998.998 0 0 0 .743-1.669M12 16a3 3 0 1 1 0-6a3 3 0 0 1 0 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeHeart(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 14h2v7a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-7h2a.998.998 0 0 0 .913-.593a.998.998 0 0 0-.17-1.076l-9-10c-.379-.422-1.107-.422-1.486 0l-9 10A1 1 0 0 0 3 14m5.653-2.359a2.224 2.224 0 0 1 3.125 0l.224.22l.223-.22a2.225 2.225 0 0 1 3.126 0a2.13 2.13 0 0 1 0 3.07L12.002 18l-3.349-3.289a2.13 2.13 0 0 1 0-3.07"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeSmile(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 14h2v7a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-7h2a.998.998 0 0 0 .913-.593a.998.998 0 0 0-.17-1.076l-9-10c-.379-.422-1.107-.422-1.486 0l-9 10A1 1 0 0 0 3 14m5.949-.316C8.98 13.779 9.762 16 12 16c2.269 0 3.042-2.287 3.05-2.311l1.9.621C16.901 14.461 15.703 18 12 18s-4.901-3.539-4.95-3.689z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hot(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.5 8c0 1.5-.5 3.5-2.9 4.3c.7-1.7.8-3.4.3-5c-.7-2.1-3-3.7-4.6-4.6c-.4-.3-1.1.1-1 .7c0 1.1-.3 2.7-2 4.4C4.1 10 3 12.3 3 14.5C3 17.4 5 21 9 21c-4-4-1-7.5-1-7.5c.8 5.9 5 7.5 7 7.5c1.7 0 5-1.2 5-6.4c0-3.1-1.3-5.5-2.4-6.9c-.3-.5-1-.2-1.1.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hotel(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="8" cy="11" r="3" fill="currentColor"/><path fill="currentColor" d="M18.205 7H12v8H4V6H2v14h2v-3h16v3h2v-4c0-.009-.005-.016-.005-.024H22V11c0-2.096-1.698-4-3.795-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hourglass(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 22h1v-2h-1v-1a7.014 7.014 0 0 0-3.433-6.02c-.355-.21-.567-.547-.567-.901v-.158c0-.354.212-.691.566-.9A7.016 7.016 0 0 0 19 5V4h1V2H4v2h1v1a7.016 7.016 0 0 0 3.434 6.021c.354.209.566.545.566.9v.158c0 .354-.212.691-.566.9A7.016 7.016 0 0 0 5 19v1H4v2zM17 4v1a5.005 5.005 0 0 1-1.004 3H8.004A5.005 5.005 0 0 1 7 5V4zM9.45 14.702c.971-.574 1.55-1.554 1.55-2.623V12h2v.079c0 1.068.579 2.049 1.551 2.623A4.98 4.98 0 0 1 16.573 17H7.427a4.977 4.977 0 0 1 2.023-2.298"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HourglassBottom(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 2H4v2h1v1a7.014 7.014 0 0 0 3.433 6.02c.355.21.567.547.567.901v.158c0 .354-.212.691-.566.9A7.016 7.016 0 0 0 5 19v1H4v2h16v-2h-1v-1a7.016 7.016 0 0 0-3.434-6.021c-.354-.208-.566-.545-.566-.9v-.158c0-.354.212-.69.566-.9A7.016 7.016 0 0 0 19 5V4h1V2zm12 3a5.01 5.01 0 0 1-2.45 4.299A3.107 3.107 0 0 0 13.166 11h-2.332a3.114 3.114 0 0 0-1.385-1.702A5.008 5.008 0 0 1 7 5V4h10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HourglassTop(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.566 11.021A7.016 7.016 0 0 0 19 5V4h1V2H4v2h1v1a7.016 7.016 0 0 0 3.434 6.021c.354.208.566.545.566.9v.158c0 .354-.212.69-.566.9A7.016 7.016 0 0 0 5 19v1H4v2h16v-2h-1v-1a7.014 7.014 0 0 0-3.433-6.02c-.355-.21-.567-.547-.567-.901v-.158c0-.355.212-.692.566-.9M17 19v1H7v-1a5.01 5.01 0 0 1 2.45-4.299A3.111 3.111 0 0 0 10.834 13h2.332c.23.691.704 1.3 1.385 1.702A5.008 5.008 0 0 1 17 19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IdCard(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 4H4c-1.103 0-2 .897-2 2v12c0 1.103.897 2 2 2h16c1.103 0 2-.897 2-2V6c0-1.103-.897-2-2-2M8.715 8c1.151 0 2 .849 2 2s-.849 2-2 2s-2-.849-2-2s.848-2 2-2m3.715 8H5v-.465c0-1.373 1.676-2.785 3.715-2.785s3.715 1.412 3.715 2.785zM19 15h-4v-2h4zm0-4h-5V9h5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Image(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.999 4h-16c-1.103 0-2 .897-2 2v12c0 1.103.897 2 2 2h16c1.103 0 2-.897 2-2V6c0-1.103-.897-2-2-2m-13.5 3a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3m5.5 10h-7l4-5l1.5 2l3-4l5.5 7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageAdd(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9 13l3-4l3 4.5V12h4V5c0-1.103-.897-2-2-2H4c-1.103 0-2 .897-2 2v12c0 1.103.897 2 2 2h8v-4H5l3-4z"/><path fill="currentColor" d="M19 14h-2v3h-3v2h3v3h2v-3h3v-2h-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 21h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2m3-7l2.363 2.363L14 11l5 7H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Inbox(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 3H4c-1.103 0-2 .897-2 2v14a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V5c0-1.103-.897-2-2-2m-1 9h-3.142c-.446 1.722-1.997 3-3.858 3s-3.412-1.278-3.858-3H4V5h16v7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoCircle(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m1 15h-2v-6h2zm0-8h-2V7h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoSquare(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 4v16a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1m8 3h2v2h-2zm0 4h2v6h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Injection(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 5.61L9.24 8.35l3.31 3.31l-1.06 1.06l-3.31-3.31l-1.77 1.77l3.31 3.31l-1.06 1.06l-3.31-3.31l-2 2A2 2 0 0 0 3 16.66l1 1.89l-2.25 2.29l1.41 1.41L5.45 20l1.89 1a2 2 0 0 0 1 .26a2 2 0 0 0 1.42-.59L18.39 12zm7.8 3.59l-1.79-1.8l1.42-1.41l1.41 1.41l1.41-1.41l-4.24-4.24l-1.41 1.41l1.41 1.42l-1.41 1.41l-1.8-1.79l-1.74-1.75l-1.41 1.42l1.03 1.03v.01l6.41 6.41h.01l1.03 1.03l1.42-1.41l-1.74-1.74z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Institution(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.857 8.485l-3-5A.997.997 0 0 0 18 3h-4.586l-.707-.707a.999.999 0 0 0-1.414 0L10.586 3H6a.997.997 0 0 0-.857.485l-3 5A1.001 1.001 0 0 0 2.002 9H2v10a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V9h-.002c0-.178-.046-.356-.141-.515M20 18h-6v-4h-4v4H4v-8h2.414l.293-.293l2-2L12 4.414l4.293 4.293l1 1l.293.293H20z"/><circle cx="11.895" cy="9.895" r="2.105" fill="currentColor"/><path fill="currentColor" d="M6 12h2v3H6zm10 0h2v3h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Invader(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 3h2v2H6zm2 16h3v2H8zm8-16h2v2h-2zm-3 16h3v2h-3zm7-8V9h-2V7h-2V5h-2v2h-4V5H8v2H6v2H4v2H2v8h2v-4h2v4h2v-3h8v3h2v-4h2v4h2v-8zm-10 1H8V9h2zm6 0h-2V9h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Joystick(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.986 9.74a3.193 3.193 0 0 0-.008-.088A5.003 5.003 0 0 0 17 5H7a4.97 4.97 0 0 0-4.987 4.737c-.01.079-.013.161-.013.253v6.51c0 .925.373 1.828 1.022 2.476A3.524 3.524 0 0 0 5.5 20c1.8 0 2.504-1 3.5-3c.146-.292.992-2 3-2c1.996 0 2.853 1.707 3 2c1.004 2 1.7 3 3.5 3c.925 0 1.828-.373 2.476-1.022A3.524 3.524 0 0 0 22 16.5V10c0-.095-.004-.18-.014-.26M7 12.031a2 2 0 1 1-.001-3.999A2 2 0 0 1 7 12.031m10-5a1 1 0 1 1 0 2a1 1 0 1 1 0-2m-2 4a1 1 0 1 1 0-2a1 1 0 1 1 0 2m2 2a1 1 0 1 1 0-2a1 1 0 1 1 0 2m2-2a1 1 0 1 1 0-2a1 1 0 1 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func JoystickAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 6H8a6 6 0 0 0 0 12h8a6 6 0 0 0 0-12m-5 7H9v2H7v-2H5v-2h2V9h2v2h2zm3.5 2a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3m3-3a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func JoystickButton(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 8h-4V4a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v4H4a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h4v4a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2v-4h4a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2M7 14l-3-2l3-2zm5 6l-2-3h4zm0-6a2 2 0 1 1 .001-4.001A2 2 0 0 1 12 14m-2-7l2-3l2 3zm7 7v-4l3 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Key(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.433 17.325L3.079 19.8a1 1 0 0 0 1.131 1.131l2.475-.354C7.06 20.524 8 18 8 18s.472.405.665.466c.412.13.813-.274.948-.684L10 16.01s.577.292.786.335c.266.055.524-.109.707-.293a.988.988 0 0 0 .241-.391L12 14.01s.675.187.906.214c.263.03.519-.104.707-.293l1.138-1.137a5.502 5.502 0 0 0 5.581-1.338a5.507 5.507 0 0 0 0-7.778a5.507 5.507 0 0 0-7.778 0a5.5 5.5 0 0 0-1.338 5.581l-7.501 7.5a.994.994 0 0 0-.282.566M18.504 5.506a2.919 2.919 0 0 1 0 4.122l-4.122-4.122a2.919 2.919 0 0 1 4.122 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Keyboard(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 5H3a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h18a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2m-8 2h2v2h-2zm0 4h2v2h-2zM9 7h2v2H9zm0 4h2v2H9zM5 7h2v2H5zm0 4h2v2H5zm12 6H7v-2h10zm2-4h-2v-2h2zm0-4h-2V7h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Label(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.813 4.419A.997.997 0 0 0 16 4H3a1 1 0 0 0-.813 1.581L6.771 12l-4.585 6.419A1 1 0 0 0 3 20h13a.997.997 0 0 0 .813-.419l5-7a.997.997 0 0 0 0-1.162z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Landmark(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 17h-2V9h2V7a.998.998 0 0 0-.594-.914l-9.432-4.191l-8.421 4.21A1 1 0 0 0 2 7v2h2v8H2v4h19zm-5-8v8h-3V9zM7 9h3v8H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Landscape(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="6.5" cy="6.5" r="2.5" fill="currentColor"/><path fill="currentColor" d="m14 7l-5.223 8.487L7 13l-5 7h20z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Laugh(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m-6.447 9.105l2.459-1.229l-1.567-1.044l1.109-1.664l3 2a1 1 0 0 1-.108 1.727l-4 2zM12 18c-4 0-5-4-5-4h10s-1 4-5 4m5.553-5.105l-4-2a1 1 0 0 1-.108-1.727l3-2l1.109 1.664l-1.566 1.044l2.459 1.229z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Layer(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.484 7.125l-9.022-5a1.003 1.003 0 0 0-.968 0l-8.978 4.96a1 1 0 0 0-.003 1.748l9.022 5.04a.995.995 0 0 0 .973.001l8.978-5a1 1 0 0 0-.002-1.749"/><path fill="currentColor" d="m12 15.856l-8.515-4.73l-.971 1.748l9 5a1 1 0 0 0 .971 0l9-5l-.971-1.748z"/><path fill="currentColor" d="m12 19.856l-8.515-4.73l-.971 1.748l9 5a1 1 0 0 0 .971 0l9-5l-.971-1.748z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayerMinus(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.484 11.125l-9.022-5a1 1 0 0 0-.968-.001l-8.978 4.96a1 1 0 0 0-.003 1.749l9.022 5.04a.995.995 0 0 0 .973.001l8.978-5a1 1 0 0 0-.002-1.749"/><path fill="currentColor" d="M20.515 15.126L12 19.856l-8.515-4.73l-.971 1.748l9 5a1 1 0 0 0 .971 0l9-5zM16 4h6v2h-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayerPlus(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m2.513 12.833l9.022 5.04a.995.995 0 0 0 .973.001l8.978-5a1 1 0 0 0-.002-1.749l-9.022-5a1 1 0 0 0-.968-.001l-8.978 4.96a1 1 0 0 0-.003 1.749"/><path fill="currentColor" d="m3.485 15.126l-.971 1.748l9 5a1 1 0 0 0 .971 0l9-5l-.971-1.748L12 19.856zM20 8V6h2V4h-2V2h-2v2h-2v2h2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Layout(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 3H5c-1.103 0-2 .897-2 2v4h18V5c0-1.103-.897-2-2-2M3 19c0 1.103.897 2 2 2h8V11H3zm12 2h4c1.103 0 2-.897 2-2v-8h-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Leaf(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m22 3.41l-.12-1.26l-1.2.4a13.84 13.84 0 0 1-6.41.64a11.87 11.87 0 0 0-6.68.9A7.23 7.23 0 0 0 3.3 9.5a9 9 0 0 0 .39 4.58a16.6 16.6 0 0 1 1.18-2.2a9.85 9.85 0 0 1 4.07-3.43a11.16 11.16 0 0 1 5.06-1A12.08 12.08 0 0 0 9.34 9.2a9.48 9.48 0 0 0-1.86 1.53a11.38 11.38 0 0 0-1.39 1.91a16.39 16.39 0 0 0-1.57 4.54A26.42 26.42 0 0 0 4 22h2a30.69 30.69 0 0 1 .59-4.32a9.25 9.25 0 0 0 4.52 1.11a11 11 0 0 0 4.28-.87C23 14.67 22 3.86 22 3.41"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftArrow(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m4.431 12.822l13 9A1 1 0 0 0 19 21V3a1 1 0 0 0-1.569-.823l-13 9a1.003 1.003 0 0 0 0 1.645"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftArrowAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5 12l7 6v-5h6v-2h-6V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftArrowCircle(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m5 11h-5v4l-5-5l5-5v4h5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftArrowSquare(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 21a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2zM12 7v4h5v2h-5v4l-5-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftDownArrowCircle(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.071 19.071c3.898-3.899 3.898-10.244 0-14.143c-3.899-3.899-10.244-3.898-14.143 0c-3.898 3.899-3.899 10.243 0 14.143c3.9 3.899 10.244 3.899 14.143 0M8.464 8.464l2.829 2.829l3.535-3.536l1.414 1.414l-3.535 3.536l2.828 2.829H8.464z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftTopArrowCircle(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.071 19.071c3.898-3.899 3.898-10.244 0-14.143c-3.899-3.898-10.243-3.898-14.143 0c-3.898 3.899-3.898 10.244 0 14.143c3.9 3.899 10.244 3.899 14.143 0M8.465 8.464h7.07l-2.828 2.829l3.535 3.536l-1.414 1.414l-3.535-3.536l-2.828 2.829z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lemon(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.45 8.74A2.23 2.23 0 0 1 21.64 7a3.51 3.51 0 0 0 .24-2.47a3.55 3.55 0 0 0-2.45-2.45a3.51 3.51 0 0 0-2.43.28a2.23 2.23 0 0 1-1.7.19a10.07 10.07 0 0 0-6.53 0a9.87 9.87 0 0 0-6.23 6.18a10.07 10.07 0 0 0 0 6.53A2.23 2.23 0 0 1 2.36 17a3.51 3.51 0 0 0-.24 2.47a3.55 3.55 0 0 0 2.45 2.45A3.51 3.51 0 0 0 7 21.64a2.23 2.23 0 0 1 1.7-.19A9.83 9.83 0 0 0 12 22a10.33 10.33 0 0 0 3.27-.54a9.87 9.87 0 0 0 6.19-6.19a10.07 10.07 0 0 0-.01-6.53M12 7a5 5 0 0 0-5 5H5a7 7 0 0 1 7-7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Like(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 21h1V8H4a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2M20 8h-7l1.122-3.368A2 2 0 0 0 12.225 2H12L7 7.438V21h11l3.912-8.596L22 12v-2a2 2 0 0 0-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationPlus(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22s8.029-5.56 8-12c0-4.411-3.589-8-8-8S4 5.589 4 9.995C3.971 16.44 11.696 21.784 12 22M8 9h3V6h2v3h3v2h-3v3h-2v-3H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lock(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C9.243 2 7 4.243 7 7v3H6a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8a2 2 0 0 0-2-2h-1V7c0-2.757-2.243-5-5-5M9 7c0-1.654 1.346-3 3-3s3 1.346 3 3v3H9zm4 10.723V20h-2v-2.277a1.993 1.993 0 0 1 .567-3.677A2.001 2.001 0 0 1 14 16a1.99 1.99 0 0 1-1 1.723"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 12c0-1.103-.897-2-2-2h-1V7c0-2.757-2.243-5-5-5S7 4.243 7 7v3H6c-1.103 0-2 .897-2 2v8c0 1.103.897 2 2 2h12c1.103 0 2-.897 2-2zM9 7c0-1.654 1.346-3 3-3s3 1.346 3 3v3H9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockOpen(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 10H9V7c0-1.654 1.346-3 3-3s3 1.346 3 3h2c0-2.757-2.243-5-5-5S7 4.243 7 7v3H6a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8a2 2 0 0 0-2-2m-7.939 5.499A2.002 2.002 0 0 1 14 16a1.99 1.99 0 0 1-1 1.723V20h-2v-2.277a1.992 1.992 0 0 1-.939-2.224"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockOpenAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 8V7c0-2.757-2.243-5-5-5S7 4.243 7 7v3H6c-1.103 0-2 .897-2 2v8c0 1.103.897 2 2 2h12c1.103 0 2-.897 2-2v-8c0-1.103-.897-2-2-2H9V7c0-1.654 1.346-3 3-3s3 1.346 3 3v1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogIn(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.5 2h-13a.5.5 0 0 0-.5.5V11h6V8l5 4l-5 4v-3H5v8.5a.5.5 0 0 0 .5.5h13a.5.5 0 0 0 .5-.5v-19a.5.5 0 0 0-.5-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogInCircle(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 3c-4.625 0-8.442 3.507-8.941 8.001H10v-3l5 4l-5 4v-3H3.06C3.56 17.494 7.376 21 12 21c4.963 0 9-4.037 9-9s-4.037-9-9-9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogOut(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 2H6a1 1 0 0 0-1 1v9l5-4v3h6v2h-6v3l-5-4v9a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogOutCircle(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 3c-4.963 0-9 4.037-9 9v.001l5-4v3h7v2H8v3l-5-4C3.001 16.964 7.037 21 12 21s9-4.037 9-9s-4.037-9-9-9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LowVision(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 4.998c-1.836 0-3.356.389-4.617.971L3.707 2.293L2.293 3.707l3.315 3.315c-2.613 1.952-3.543 4.618-3.557 4.66l-.105.316l.105.316C2.073 12.382 4.367 19 12 19c1.835 0 3.354-.389 4.615-.971l3.678 3.678l1.414-1.414l-3.317-3.317c2.614-1.952 3.545-4.618 3.559-4.66l.105-.316l-.105-.316c-.022-.068-2.316-6.686-9.949-6.686M12.043 7H12a5 5 0 0 1 5 5a4.894 4.894 0 0 1-.852 2.734l-.721-.721A3.919 3.919 0 0 0 16 11.999c0-.474-.099-.925-.255-1.349A.985.985 0 0 1 15 11a1 1 0 0 1-1-1c0-.439.288-.802.682-.936A3.965 3.965 0 0 0 12 7.999c-.735 0-1.419.218-2.015.572l-.72-.72C10.053 7.326 10.982 7 12 7h-.043L12 6.998zm-7.969 4.999c.103-.235.274-.586.521-.989l5.867 5.867c-4.213-.647-5.939-3.842-6.388-4.878m9.247 4.908l-7.48-7.48a8.146 8.146 0 0 1 1.188-.984l8.055 8.055a8.835 8.835 0 0 1-1.763.409"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagicWand(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m11 4l-.5-1l-.5 1l-1 .125l.834.708L9.5 6l1-.666l1 .666l-.334-1.167l.834-.708zm8.334 10.666L18.5 13l-.834 1.666l-1.666.209l1.389 1.181L16.834 18l1.666-1.111L20.166 18l-.555-1.944L21 14.875zM6.667 6.333L6 5l-.667 1.333L4 6.5l1.111.944L4.667 9L6 8.111L7.333 9l-.444-1.556L8 6.5zM3.414 17c0 .534.208 1.036.586 1.414L5.586 20c.378.378.88.586 1.414.586s1.036-.208 1.414-.586L20 8.414c.378-.378.586-.88.586-1.414S20.378 5.964 20 5.586L18.414 4c-.756-.756-2.072-.756-2.828 0L4 15.586c-.378.378-.586.88-.586 1.414M17 5.414L18.586 7L15 10.586L13.414 9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Magnet(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 3H5a1 1 0 0 0-1 1v3h5V4a1 1 0 0 0-1-1m7 1v3h5V4a1 1 0 0 0-1-1h-3a1 1 0 0 0-1 1m0 10a3 3 0 0 1-6 0V9H4v5a8 8 0 0 0 16 0V9h-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Map(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C7.589 2 4 5.589 4 9.995C3.971 16.44 11.696 21.784 12 22c0 0 8.029-5.56 8-12c0-4.411-3.589-8-8-8m0 12c-2.21 0-4-1.79-4-4s1.79-4 4-4s4 1.79 4 4s-1.79 4-4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9 6.882l-7-3.5v13.236l7 3.5l6-3l7 3.5V7.382l-7-3.5zM15 15l-6 3V9l6-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapPin(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12 17l1-2V9.858c1.721-.447 3-2 3-3.858c0-2.206-1.794-4-4-4S8 3.794 8 6c0 1.858 1.279 3.411 3 3.858V15z"/><path fill="currentColor" d="m16.267 10.563l-.533 1.928C18.325 13.207 20 14.584 20 16c0 1.892-3.285 4-8 4s-8-2.108-8-4c0-1.416 1.675-2.793 4.267-3.51l-.533-1.928C4.197 11.54 2 13.623 2 16c0 3.364 4.393 6 10 6s10-2.636 10-6c0-2.377-2.197-4.46-5.733-5.437"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mask(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 6H5C3.346 6 2 7.346 2 9v5c0 2.206 1.794 4 4 4h1.637c1.166 0 2.28-.557 2.981-1.491c.66-.879 2.104-.88 2.764.001A3.744 3.744 0 0 0 16.363 18H18c2.206 0 4-1.794 4-4V9c0-1.654-1.346-3-3-3M7.5 13C6.119 13 5 12.328 5 11.5S6.119 10 7.5 10s2.5.672 2.5 1.5S8.881 13 7.5 13m9 0c-1.381 0-2.5-.672-2.5-1.5s1.119-1.5 2.5-1.5s2.5.672 2.5 1.5s-1.119 1.5-2.5 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Medal(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2h-4v4.059a8.946 8.946 0 0 1 4 1.459zm-6 0H7v5.518a8.946 8.946 0 0 1 4-1.459zm1 20a7 7 0 1 0 0-14a7 7 0 0 0 0 14m-1.225-8.519L12 11l1.225 2.481l2.738.397l-1.981 1.932l.468 2.727L12 17.25l-2.449 1.287l.468-2.727l-1.981-1.932z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Megaphone(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.664 3.478L8 8v7l.748.267l-1.127 2.254a1.999 1.999 0 0 0 1.156 2.792l4.084 1.361a2.015 2.015 0 0 0 2.421-1.003l1.303-2.606l4.079 1.457A1 1 0 0 0 22 18.581V4.419a1 1 0 0 0-1.336-.941m-7.171 16.299L9.41 18.416l1.235-2.471l4.042 1.444zM4 15h2V8H4c-1.103 0-2 .897-2 2v3c0 1.103.897 2 2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Meh(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m-5 8.5a1.5 1.5 0 1 1 3.001.001A1.5 1.5 0 0 1 7 10.5m9 6.5H7.974v-2H16zm-.507-5.014a1.494 1.494 0 1 1 .001-2.987a1.494 1.494 0 0 1-.001 2.987"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MehAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m-6 8h4v2H6zm10 7H7.974v-2H16zm2-5h-4v-2h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MehBlank(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2M8.5 12a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 8.5 12m6.993-.014a1.494 1.494 0 1 1 .001-2.987a1.494 1.494 0 0 1-.001 2.987"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MemoryCard(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 2H6c-1.103 0-2 .897-2 2v16c0 1.103.897 2 2 2h12c1.103 0 2-.897 2-2V7zm-6 8H7V6h2zm3 0h-2V6h2zm3 0h-2V6h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Message(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 2H4c-1.103 0-2 .894-2 1.992v12.016C2 17.106 2.897 18 4 18h3v4l6.351-4H20c1.103 0 2-.894 2-1.992V3.992A1.998 1.998 0 0 0 20 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageAdd(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 2H4c-1.103 0-2 .894-2 1.992v12.016C2 17.106 2.897 18 4 18h3v4l6.351-4H20c1.103 0 2-.894 2-1.992V3.992A1.998 1.998 0 0 0 20 2m-3 9h-4v4h-2v-4H7V9h4V5h2v4h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.999 2h-14c-1.103 0-2 .897-2 2v12c0 1.103.897 2 2 2h3.5l3.5 4l3.5-4h3.5c1.103 0 2-.897 2-2V4c0-1.103-.897-2-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageAltAdd(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8.5 18l3.5 4l3.5-4H19c1.103 0 2-.897 2-2V4c0-1.103-.897-2-2-2H5c-1.103 0-2 .897-2 2v12c0 1.103.897 2 2 2zM7 9h4V5h2v4h4v2h-4v4h-2v-4H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageAltCheck(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 4v12c0 1.103.897 2 2 2h3.5l3.5 4l3.5-4H19c1.103 0 2-.897 2-2V4c0-1.103-.897-2-2-2H5c-1.103 0-2 .897-2 2m5.707 4.293L11 10.586l4.793-4.793l1.414 1.414L11 13.414L7.293 9.707z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageAltDetail(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8.5 18l3.5 4l3.5-4H19c1.103 0 2-.897 2-2V4c0-1.103-.897-2-2-2H5c-1.103 0-2 .897-2 2v12c0 1.103.897 2 2 2zM7 7h10v2H7zm0 4h7v2H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageAltDots(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 2H5c-1.103 0-2 .897-2 2v12c0 1.103.897 2 2 2h3.5l3.5 4l3.5-4H19c1.103 0 2-.897 2-2V4c0-1.103-.897-2-2-2M9 12a2 2 0 1 1 .001-4.001A2 2 0 0 1 9 12m6 0a2 2 0 1 1 .001-4.001A2 2 0 0 1 15 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageAltEdit(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 2c-1.103 0-2 .897-2 2v12c0 1.103.897 2 2 2h3.5l3.5 4l3.5-4H19c1.103 0 2-.897 2-2V4c0-1.103-.897-2-2-2zm4.302 11.987h-1.8v-1.799l4.978-4.97l1.798 1.799zm5.823-5.817l-1.798-1.799L14.698 5l1.8 1.799z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageAltError(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 4v12c0 1.103.897 2 2 2h3.5l3.5 4l3.5-4H19c1.103 0 2-.897 2-2V4c0-1.103-.897-2-2-2H5c-1.103 0-2 .897-2 2m8 1h2v6h-2zm0 8h2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageAltMinus(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 2c-1.103 0-2 .897-2 2v12c0 1.103.897 2 2 2h3.5l3.5 4l3.5-4H19c1.103 0 2-.897 2-2V4c0-1.103-.897-2-2-2zm11 9H8V9h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageAltX(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8.5 18l3.5 4l3.5-4H19c1.103 0 2-.897 2-2V4c0-1.103-.897-2-2-2H5c-1.103 0-2 .897-2 2v12c0 1.103.897 2 2 2zM7.293 6.707l1.414-1.414L12 8.586l3.293-3.293l1.414 1.414L13.414 10l3.293 3.293l-1.414 1.414L12 11.414l-3.293 3.293l-1.414-1.414L10.586 10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageCheck(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 2H4c-1.103 0-2 .894-2 1.992v12.016C2 17.106 2.897 18 4 18h3v4l6.351-4H20c1.103 0 2-.894 2-1.992V3.992A1.998 1.998 0 0 0 20 2m-9 11.914l-3.707-3.707l1.414-1.414L11 11.086l4.793-4.793l1.414 1.414z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageDetail(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 2H4c-1.103 0-2 .894-2 1.992v12.016C2 17.106 2.897 18 4 18h3v4l6.351-4H20c1.103 0 2-.894 2-1.992V3.992A1.998 1.998 0 0 0 20 2m-6 11H7v-2h7zm3-4H7V7h10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageDots(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 2H4c-1.103 0-2 .894-2 1.992v12.017C2 17.106 2.897 18 4 18h3v4l6.351-4H20c1.103 0 2-.894 2-1.992V3.992A1.998 1.998 0 0 0 20 2m-9 8a2 2 0 1 1-2-2c.086 0 .167.015.25.025c.082-.014.164-.025.25-.025A1.5 1.5 0 0 1 11 9.5c0 .086-.012.168-.025.25c.01.083.025.165.025.25m4 2a2 2 0 0 1-2-2c0-.086.015-.167.025-.25A1.592 1.592 0 0 1 13 9.5A1.5 1.5 0 0 1 14.5 8c.086 0 .168.011.25.025c.083-.01.164-.025.25-.025a2 2 0 0 1 0 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageEdit(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 2H4c-1.103 0-2 .894-2 1.992v12.016C2 17.106 2.897 18 4 18h3v4l6.351-4H20c1.103 0 2-.894 2-1.992V3.992A1.998 1.998 0 0 0 20 2M8.999 14.999H7V13l5.53-5.522l1.998 1.999zm6.472-6.464l-1.999-1.999l1.524-1.523l1.999 1.999z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageError(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 2H4c-1.103 0-2 .894-2 1.992v12.016C2 17.106 2.897 18 4 18h3v4l6.351-4H20c1.103 0 2-.894 2-1.992V3.992A1.998 1.998 0 0 0 20 2m-7 13h-2v-2h2zm0-4h-2V5h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageMinus(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 2H4c-1.103 0-2 .894-2 1.992v12.016C2 17.106 2.897 18 4 18h3v4l6.351-4H20c1.103 0 2-.894 2-1.992V3.992A1.998 1.998 0 0 0 20 2m-4 9H8V9h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageRounded(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 5.589 2 10c0 2.908 1.898 5.515 5 6.934V22l5.34-4.005C17.697 17.852 22 14.32 22 10c0-4.411-4.486-8-10-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageRoundedAdd(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 5.589 2 10c0 2.907 1.897 5.515 5 6.934V22l5.34-4.005C17.697 17.853 22 14.32 22 10c0-4.411-4.486-8-10-8m4 9h-3v3h-2v-3H8V9h3V6h2v3h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageRoundedCheck(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 5.589 2 10c0 2.907 1.897 5.515 5 6.934V22l5.34-4.005C17.697 17.853 22 14.32 22 10c0-4.411-4.486-8-10-8m-1 12.414l-3.707-3.707l1.414-1.414L11 11.586l4.793-4.793l1.414 1.414z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageRoundedDetail(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 5.589 2 10c0 2.907 1.897 5.515 5 6.934V22l5.34-4.005C17.697 17.853 22 14.32 22 10c0-4.411-4.486-8-10-8m2 11H7v-2h7zm3-4H7V7h10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageRoundedDots(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 5.589 2 10c0 2.908 1.897 5.516 5 6.934V22l5.34-4.004C17.697 17.852 22 14.32 22 10c0-4.411-4.486-8-10-8m-2.5 9a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3m5 0a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageRoundedEdit(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 5.589 2 10c0 2.907 1.897 5.516 5 6.934V22l5.34-4.005C17.697 17.854 22 14.32 22 10c0-4.411-4.486-8-10-8M9.302 13.986H7.503v-1.798l4.976-4.97l1.798 1.799zm5.823-5.816l-1.799-1.798l1.372-1.371l1.799 1.798z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageRoundedError(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 5.589 2 10c0 2.907 1.897 5.516 5 6.934V22l5.34-4.005C17.697 17.854 22 14.32 22 10c0-4.411-4.486-8-10-8m1 12h-2v-2h2zm0-4h-2V5h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageRoundedMinus(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 5.589 2 10c0 2.907 1.897 5.515 5 6.934V22l5.34-4.005C17.697 17.853 22 14.32 22 10c0-4.411-4.486-8-10-8m4 9H8V9h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageRoundedX(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 5.589 2 10c0 2.907 1.897 5.515 5 6.934V22l5.34-4.005C17.697 17.853 22 14.32 22 10c0-4.411-4.486-8-10-8m3.707 10.293l-1.414 1.414L12 11.414l-2.293 2.293l-1.414-1.414L10.586 10L8.293 7.707l1.414-1.414L12 8.586l2.293-2.293l1.414 1.414L13.414 10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageSquare(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 2H8C4.691 2 2 4.691 2 8v13a1 1 0 0 0 1 1h13c3.309 0 6-2.691 6-6V8c0-3.309-2.691-6-6-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageSquareAdd(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 2H8C4.691 2 2 4.691 2 8v13a1 1 0 0 0 1 1h13c3.309 0 6-2.691 6-6V8c0-3.309-2.691-6-6-6m1 11h-4v4h-2v-4H7v-2h4V7h2v4h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageSquareCheck(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 2H8C4.691 2 2 4.691 2 8v13a1 1 0 0 0 1 1h13c3.309 0 6-2.691 6-6V8c0-3.309-2.691-6-6-6m-5 14.414l-3.707-3.707l1.414-1.414L11 13.586l4.793-4.793l1.414 1.414z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageSquareDetail(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 2H8C4.691 2 2 4.691 2 8v13a1 1 0 0 0 1 1h13c3.309 0 6-2.691 6-6V8c0-3.309-2.691-6-6-6m-2 13H7v-2h7zm3-4H7V9h10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageSquareDots(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 2H8C4.691 2 2 4.691 2 8v13a1 1 0 0 0 1 1h13c3.309 0 6-2.691 6-6V8c0-3.309-2.691-6-6-6m-5 10.5A1.5 1.5 0 0 1 9.5 14c-.086 0-.168-.011-.25-.025c-.083.01-.164.025-.25.025a2 2 0 1 1 2-2c0 .085-.015.167-.025.25c.013.082.025.164.025.25m4 1.5c-.086 0-.167-.015-.25-.025a1.471 1.471 0 0 1-.25.025a1.5 1.5 0 0 1-1.5-1.5c0-.085.012-.168.025-.25c-.01-.083-.025-.164-.025-.25a2 2 0 1 1 2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageSquareEdit(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 2H8C4.691 2 2 4.691 2 8v13a1 1 0 0 0 1 1h13c3.309 0 6-2.691 6-6V8c0-3.309-2.691-6-6-6M8.999 17H7v-1.999l5.53-5.522l1.999 1.999zm6.473-6.465l-1.999-1.999l1.524-1.523l1.999 1.999z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageSquareError(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 2H8C4.691 2 2 4.691 2 8v13a1 1 0 0 0 1 1h13c3.309 0 6-2.691 6-6V8c0-3.309-2.691-6-6-6m-3 16h-2v-2h2zm0-4h-2V6h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageSquareMinus(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 2H8C4.691 2 2 4.691 2 8v13a1 1 0 0 0 1 1h13c3.309 0 6-2.691 6-6V8c0-3.309-2.691-6-6-6m0 11H8v-2h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageSquareX(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 2H8C4.691 2 2 4.691 2 8v13a1 1 0 0 0 1 1h13c3.309 0 6-2.691 6-6V8c0-3.309-2.691-6-6-6m.706 13.293l-1.414 1.414L12 13.415l-3.292 3.292l-1.414-1.414l3.292-3.292l-3.292-3.292l1.414-1.414L12 10.587l3.292-3.292l1.414 1.414l-3.292 3.292z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageX(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 2H4c-1.103 0-2 .894-2 1.992v12.016C2 17.106 2.897 18 4 18h3v4l6.351-4H20c1.103 0 2-.894 2-1.992V3.992A1.998 1.998 0 0 0 20 2m-3.293 11.293l-1.414 1.414L12 11.414l-3.293 3.293l-1.414-1.414L10.586 10L7.293 6.707l1.414-1.414L12 8.586l3.293-3.293l1.414 1.414L13.414 10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Meteor(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.507 2.138a1 1 0 0 0-1.155.102L4.196 9.197c-2.924 2.924-2.924 7.682 0 10.606a7.472 7.472 0 0 0 5.3 2.192c1.924 0 3.85-.734 5.317-2.202l6.903-7.096A1 1 0 0 0 21 11h-3.301l4.175-7.514a1.001 1.001 0 0 0-1.359-1.36l-7.11 3.95l.576-2.879a1.002 1.002 0 0 0-.474-1.059M14 14.5a4.5 4.5 0 0 1-9 0c0-1.57.807-2.949 2.025-3.754c-.01.084-.025.167-.025.254a2 2 0 1 0 3.845-.772C12.669 10.802 14 12.486 14 14.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Microchip(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.999 22h8c1.103 0 2-.897 2-2V4c0-1.103-.897-2-2-2h-8c-1.103 0-2 .897-2 2v16c0 1.103.897 2 2 2m-5-15h2V5h-2v.5h-1v1h1zm18-2h-2v2h2v-.5h1v-1h-1zm-18 6h2V9h-2v.5h-1v1h1zm18-2h-2v2h2v-.5h1v-1h-1zm-18 6h2v-2h-2v.5h-1v1h1zm18-2h-2v2h2v-.5h1v-1h-1zm-18 6h2v-2h-2v.5h-1v1h1zm18-2h-2v2h2v-.5h1v-1h-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Microphone(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 16c2.206 0 4-1.794 4-4V6c0-2.217-1.785-4.021-3.979-4.021a.933.933 0 0 0-.209.025A4.006 4.006 0 0 0 8 6v6c0 2.206 1.794 4 4 4"/><path fill="currentColor" d="M11 19.931V22h2v-2.069c3.939-.495 7-3.858 7-7.931h-2c0 3.309-2.691 6-6 6s-6-2.691-6-6H4c0 4.072 3.061 7.436 7 7.931"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicrophoneAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 12H4c0 4.072 3.061 7.436 7 7.931V22h2v-2.069c3.939-.495 7-3.858 7-7.931h-2c0 3.309-2.691 6-6 6s-6-2.691-6-6"/><path fill="currentColor" d="M8 12c0 2.206 1.794 4 4 4s4-1.794 4-4h-2v-2h2V8h-2V6h2c0-2.217-1.785-4.021-3.979-4.021a.933.933 0 0 0-.209.025A4.006 4.006 0 0 0 8 6h4v2H8v2h4v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicrophoneOff(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.707 20.293l-3.4-3.4A7.93 7.93 0 0 0 20 12h-2a5.945 5.945 0 0 1-1.119 3.467l-1.449-1.45A3.926 3.926 0 0 0 16 12V6c0-2.217-1.785-4.021-3.979-4.021c-.07 0-.14.009-.209.025A4.006 4.006 0 0 0 8 6v.586L3.707 2.293L2.293 3.707l18 18zM6 12H4c0 4.072 3.06 7.436 7 7.931V22h2v-2.069a7.935 7.935 0 0 0 2.241-.63l-1.549-1.548A5.983 5.983 0 0 1 12 18c-3.309 0-6-2.691-6-6"/><path fill="currentColor" d="M8.007 12.067a3.996 3.996 0 0 0 3.926 3.926z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusCircle(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m5 11H7v-2h10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusSquare(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 3a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2zm12 10H7v-2h10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mobile(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 22c1.103 0 2-.897 2-2V4c0-1.103-.897-2-2-2H8c-1.103 0-2 .897-2 2v16c0 1.103.897 2 2 2zm-5-5a1 1 0 1 1 0 2a1 1 0 1 1 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MobileVibration(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.535 2.808a2.003 2.003 0 0 0-2.828 0l-9.899 9.899a2.001 2.001 0 0 0 0 2.828l5.657 5.657c.39.39.902.585 1.414.585s1.024-.195 1.414-.585l9.899-9.899c.78-.779.78-2.049 0-2.828zM8.707 16.707a.999.999 0 1 1-1.414-1.414a.999.999 0 1 1 1.414 1.414m7 5l-1.414-1.414l6-6l1.414 1.415zM8.293 2.293l1.414 1.414l-6 6l-1.414-1.415z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Moon(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 11.807A9.002 9.002 0 0 1 10.049 2a9.942 9.942 0 0 0-5.12 2.735c-3.905 3.905-3.905 10.237 0 14.142c3.906 3.906 10.237 3.905 14.143 0a9.946 9.946 0 0 0 2.735-5.119A9.003 9.003 0 0 1 12 11.807"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mouse(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.975 22H12c3.859 0 7-3.14 7-7V9c0-3.841-3.127-6.974-6.981-7h-.06C8.119 2.022 5 5.157 5 9v6c0 3.86 3.129 7 6.975 7M11 6h2v6h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MouseAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 2v8h6V8c0-3.309-2.691-6-6-6M5 16c0 3.309 2.691 6 6 6h2c3.309 0 6-2.691 6-6v-4H5zm0-8v2h6V2C7.691 2 5 4.691 5 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Movie(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 3H4c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2h16c1.103 0 2-.897 2-2V5c0-1.103-.897-2-2-2m.001 6c-.001 0-.001 0 0 0h-.466l-2.667-4H20zm-5.466 0l-2.667-4h2.596l2.667 4zm-2.404 0H9.535L6.869 5h2.596zM4 5h.465l2.667 4H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoviePlay(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 3H4c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2h16c1.103 0 2-.897 2-2V5c0-1.103-.897-2-2-2m.001 6c-.001 0-.001 0 0 0h-.465l-2.667-4H20zM15.5 15L10 18v-6zm-.964-6l-2.667-4h2.596l2.667 4zm-2.404 0H9.536L6.869 5h2.596zM4 5h.465l2.667 4H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Music(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 18.573c2.206 0 4-1.794 4-4V4.428L19 7.7v7.43a3.953 3.953 0 0 0-2-.557c-2.206 0-4 1.794-4 4s1.794 4 4 4s4-1.794 4-4V7a.998.998 0 0 0-.658-.939l-11-4A.999.999 0 0 0 8 3v8.13a3.953 3.953 0 0 0-2-.557c-2.206 0-4 1.794-4 4s1.794 4 4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Navigation(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.002 9.538c-.023.411.207.794.581.966l7.504 3.442l3.442 7.503c.164.356.52.583.909.583l.057-.002a1 1 0 0 0 .894-.686l5.595-17.032c.117-.358.023-.753-.243-1.02s-.66-.358-1.02-.243L2.688 8.645a.997.997 0 0 0-.686.893"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NetworkChart(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.5 3A2.502 2.502 0 0 0 17 5.5c0 .357.078.696.214 1.005l-1.955 2.199A3.977 3.977 0 0 0 13 8c-.74 0-1.424.216-2.019.566L8.707 6.293l-.023.023C8.88 5.918 9 5.475 9 5a3 3 0 1 0-3 3c.475 0 .917-.12 1.316-.316l-.023.023L9.567 9.98A3.956 3.956 0 0 0 9 12c0 .997.38 1.899.985 2.601l-2.577 2.576A2.472 2.472 0 0 0 6.5 17C5.122 17 4 18.121 4 19.5S5.122 22 6.5 22S9 20.879 9 19.5c0-.321-.066-.626-.177-.909l2.838-2.838c.421.15.867.247 1.339.247c2.206 0 4-1.794 4-4c0-.636-.163-1.229-.428-1.764l2.117-2.383c.256.088.526.147.811.147C20.879 8 22 6.879 22 5.5S20.879 3 19.5 3M13 14c-1.103 0-2-.897-2-2s.897-2 2-2s2 .897 2 2s-.897 2-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func News(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 17a1 1 0 0 1-2 0V5a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v13a2 2 0 0 0 2 2h15c1.654 0 3-1.346 3-3V7h-2zM12 7h3v2h-3zm0 4h3v2h-3zM5 7h5v6H5zm0 10v-2h10v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoEntry(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m5 12H7v-4h10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Note(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h8l8-8V5a2 2 0 0 0-2-2m-7 16v-7h7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Notepad(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 4h-3V2h-2v2h-4V2H8v2H5c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2h14c1.103 0 2-.897 2-2V6c0-1.103-.897-2-2-2m-7 10H7v-2h5zm5-4H7V8h10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Notification(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="18" cy="6" r="3" fill="currentColor"/><path fill="currentColor" d="M13 6c0-.712.153-1.387.422-2H6c-1.103 0-2 .897-2 2v12c0 1.103.897 2 2 2h12c1.103 0 2-.897 2-2v-7.422A4.962 4.962 0 0 1 18 11a5 5 0 0 1-5-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationOff(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="18" cy="6" r="3" fill="currentColor"/><path fill="currentColor" d="M20 18v-7.422A4.962 4.962 0 0 1 18 11a5 5 0 0 1-5-5c0-.712.153-1.387.422-2H6c-.178 0-.347.031-.51.076L3.707 2.293L2.293 3.707l18 18l1.414-1.414l-1.783-1.783c.045-.163.076-.332.076-.51M4 18c0 1.103.897 2 2 2h9.879L4 8.121z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ObjectsHorizontalCenter(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 13h-6v-2h4a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1h-4V2h-2v3H7a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h4v2H5a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h6v3h2v-3h6a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ObjectsHorizontalLeft(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 2h2v20H2z"/><rect width="16" height="6" x="6" y="13" fill="currentColor" rx="1"/><rect width="12" height="6" x="6" y="5" fill="currentColor" rx="1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ObjectsHorizontalRight(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 2h2v20h-2z"/><rect width="16" height="6" x="2" y="13" fill="currentColor" rx="1"/><rect width="12" height="6" x="6" y="5" fill="currentColor" rx="1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ObjectsVerticalBottom(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 20h20v2H2z"/><rect width="6" height="16" x="5" y="2" fill="currentColor" rx="1"/><rect width="6" height="12" x="13" y="6" fill="currentColor" rx="1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ObjectsVerticalCenter(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 7a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1v4h-2V5a1 1 0 0 0-1-1H6a1 1 0 0 0-1 1v6H2v2h3v6a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-6h2v4a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-4h3v-2h-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ObjectsVerticalTop(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 2h20v2H2z"/><rect width="6" height="16" x="5" y="6" fill="currentColor" rx="1"/><rect width="6" height="12" x="13" y="6" fill="currentColor" rx="1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Offer(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.749 12l1.104-1.908a1 1 0 0 0-.365-1.366l-1.91-1.104v-2.2a1 1 0 0 0-1-1h-2.199l-1.103-1.909a1.008 1.008 0 0 0-.607-.466a.993.993 0 0 0-.759.1L12 3.251l-1.91-1.105a1 1 0 0 0-1.366.366L7.62 4.422H5.421a1 1 0 0 0-1 1v2.199l-1.91 1.104a.998.998 0 0 0-.365 1.367L3.25 12l-1.104 1.908a1.004 1.004 0 0 0 .364 1.367l1.91 1.104v2.199a1 1 0 0 0 1 1h2.2l1.104 1.91a1.01 1.01 0 0 0 .866.5c.174 0 .347-.046.501-.135l1.908-1.104l1.91 1.104a1.001 1.001 0 0 0 1.366-.365l1.103-1.91h2.199a1 1 0 0 0 1-1v-2.199l1.91-1.104a1 1 0 0 0 .365-1.367zM9.499 6.99a1.5 1.5 0 1 1-.001 3.001a1.5 1.5 0 0 1 .001-3.001m.3 9.6l-1.6-1.199l6-8l1.6 1.199zm4.7.4a1.5 1.5 0 1 1 .001-3.001a1.5 1.5 0 0 1-.001 3.001"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Package(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.993 7.95a.96.96 0 0 0-.029-.214c-.007-.025-.021-.049-.03-.074c-.021-.057-.04-.113-.07-.165c-.016-.027-.038-.049-.057-.075c-.032-.045-.063-.091-.102-.13c-.023-.022-.053-.04-.078-.061c-.039-.032-.075-.067-.12-.094c-.004-.003-.009-.003-.014-.006l-.008-.006l-8.979-4.99a1.002 1.002 0 0 0-.97-.001l-9.021 4.99c-.003.003-.006.007-.011.01l-.01.004c-.035.02-.061.049-.094.073c-.036.027-.074.051-.106.082c-.03.031-.053.067-.079.102c-.027.035-.057.066-.079.104c-.026.043-.04.092-.059.139c-.014.033-.032.064-.041.1a.975.975 0 0 0-.029.21c-.001.017-.007.032-.007.05V16c0 .363.197.698.515.874l8.978 4.987l.001.001l.002.001l.02.011c.043.024.09.037.135.054c.032.013.063.03.097.039a1.013 1.013 0 0 0 .506 0c.033-.009.064-.026.097-.039c.045-.017.092-.029.135-.054l.02-.011l.002-.001l.001-.001l8.978-4.987c.316-.176.513-.511.513-.874V7.998c0-.017-.006-.031-.007-.048m-10.021 3.922L5.058 8.005L7.82 6.477l6.834 3.905zm.048-7.719L18.941 8l-2.244 1.247l-6.83-3.903zM13 19.301l.002-5.679L16 11.944V15l2-1v-3.175l2-1.119v5.705z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paint(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.084 2.914c-1.178-1.179-3.234-1.179-4.412 0l-8.379 8.379a.999.999 0 0 0 0 1.414l3 3a.997.997 0 0 0 1.414 0l8.379-8.379a3.123 3.123 0 0 0-.002-4.414m-1.412 3L12 13.586L10.414 12l7.672-7.672a1.146 1.146 0 0 1 1.586.002a1.123 1.123 0 0 1 0 1.584M8 15c-1.265-.634-3.5 0-3.5 2c0 1.197.5 2-1.5 3c0 0 3.25 2.25 5.5 0c1.274-1.274 1.494-4-.5-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaintRoll(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 2H7c-1.103 0-2 .897-2 2v3c0 1.103.897 2 2 2h11c1.103 0 2-.897 2-2V4c0-1.103-.897-2-2-2"/><path fill="currentColor" d="M13 15v-2c0-1.103-.897-2-2-2H4V5c-1.103 0-2 .897-2 2v4c0 1.103.897 2 2 2h7v2a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-5a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Palette(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.38 21.646A9.985 9.985 0 0 0 12 22l.141-.001a2.998 2.998 0 0 0 2.515-1.425c.542-.876.6-1.953.153-2.88l-.198-.415c-.453-.942-.097-1.796.388-2.281c.485-.485 1.341-.841 2.28-.388h.001l.413.199a2.99 2.99 0 0 0 2.881-.153A2.997 2.997 0 0 0 22 12.141a9.926 9.926 0 0 0-.353-2.76c-1.038-3.827-4.353-6.754-8.246-7.285c-3.149-.427-6.241.602-8.471 2.833S1.666 10.247 2.096 13.4c.53 3.894 3.458 7.208 7.284 8.246M15.5 6a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3m-5-1a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3M9 15.506a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0m-2.5-6.5a1.5 1.5 0 1 1-.001 3.001A1.5 1.5 0 0 1 6.5 9.006"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaperPlane(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m2.6 13.083l3.452 1.511L16 9.167l-6 7l8.6 3.916a1 1 0 0 0 1.399-.85l1-15a1.002 1.002 0 0 0-1.424-.972l-17 8a1.002 1.002 0 0 0 .025 1.822M8 22.167l4.776-2.316L8 17.623z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Parking(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 3H5v18h4v-5h4.5c3.584 0 6.5-2.916 6.5-6.5S17.084 3 13.5 3m0 9H9V7h4.5C14.879 7 16 8.121 16 9.5S14.879 12 13.5 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Party(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23 7a8.44 8.44 0 0 0-5 1.31c-.36-.41-.73-.82-1.12-1.21l-.29-.27l.14-.12a3.15 3.15 0 0 0 .9-3.49A3.9 3.9 0 0 0 14 1v2a2 2 0 0 1 1.76 1c.17.4 0 .84-.47 1.31l-.23.21a16.71 16.71 0 0 0-3.41-2.2c-2.53-1.14-3.83-.61-4.47 0a2.18 2.18 0 0 0-.46.68l-.18.53L5.1 8.87C6.24 11.71 9 16.76 15 18.94l5-1.66a1 1 0 0 0 .43-.31l.21-.18c1.43-1.44.51-4.21-1.41-6.9A6.63 6.63 0 0 1 23 9zm-3.79 8.37h-.06c-.69.37-3.55-.57-6.79-3.81c-.34-.34-.66-.67-.95-1c-.1-.11-.19-.23-.29-.35l-.53-.64l-.28-.39c-.14-.19-.28-.38-.4-.56s-.16-.26-.24-.39s-.22-.34-.31-.51s-.13-.24-.19-.37s-.17-.28-.23-.42s-.09-.23-.14-.34s-.11-.27-.15-.4S8.6 6 8.58 5.9s-.06-.24-.08-.34a2 2 0 0 1 0-.24a1.15 1.15 0 0 1 0-.26l.11-.31c.17-.18.91-.23 2.23.37a13.83 13.83 0 0 1 2.49 1.54A4.17 4.17 0 0 1 12 7v2a6.43 6.43 0 0 0 3-.94l.49.46c.44.43.83.86 1.19 1.27A5.31 5.31 0 0 0 16 13.2l2-.39a3.23 3.23 0 0 1 0-1.14c1.29 1.97 1.53 3.39 1.21 3.7M4.4 11l-2.23 6.7A3.28 3.28 0 0 0 5.28 22a3.21 3.21 0 0 0 1-.17l6.52-2.17A18.7 18.7 0 0 1 4.4 11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paste(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 4h6v2H9zm11 7h-7a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h7a2 2 0 0 0 2-2v-7a2 2 0 0 0-2-2"/><path fill="currentColor" d="M21 9V6a2 2 0 0 0-2-2h-2a2 2 0 0 0-2-2H9a2 2 0 0 0-2 2H5a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h4v-9a2 2 0 0 1 2-2zM9 6V4h6v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pear(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.12 22a7.71 7.71 0 0 0 6.57-5a7.23 7.23 0 0 0 .46-3.21a3.26 3.26 0 0 1 1-2.57l.61-.61A3.89 3.89 0 0 0 19.43 6l2.28-2.28l-1.42-1.43L18 4.55a3.82 3.82 0 0 0-4.66.57l-.75.75a3.22 3.22 0 0 1-2.52 1a7.05 7.05 0 0 0-3.32.57A7.75 7.75 0 0 0 2 14.11A7.59 7.59 0 0 0 10.12 22M9.5 9.25v1.5a3.75 3.75 0 0 0-3.75 3.75h-1.5A5.26 5.26 0 0 1 9.5 9.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pen(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.587 6.999H7.702a2 2 0 0 0-1.88 1.316l-3.76 10.342c-.133.365-.042.774.232 1.049l.293.293l6.422-6.422c-.001-.026-.008-.052-.008-.078a1.5 1.5 0 1 1 1.5 1.5c-.026 0-.052-.007-.078-.008l-6.422 6.422l.293.293a.997.997 0 0 0 1.049.232l10.342-3.761a2 2 0 0 0 1.316-1.88v-3.885L19 10.414L13.586 5zm8.353 2.062l-5-5l2.12-2.121l5 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pencil(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.707 19.707L18 10.414L13.586 6l-9.293 9.293a1.003 1.003 0 0 0-.263.464L3 21l5.242-1.03c.176-.044.337-.135.465-.263M21 7.414a2 2 0 0 0 0-2.828L19.414 3a2 2 0 0 0-2.828 0L15 4.586L19.414 9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Phone(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.487 17.14l-4.065-3.696a1.001 1.001 0 0 0-1.391.043l-2.393 2.461c-.576-.11-1.734-.471-2.926-1.66c-1.192-1.193-1.553-2.354-1.66-2.926l2.459-2.394a1 1 0 0 0 .043-1.391L6.859 3.513a1 1 0 0 0-1.391-.087l-2.17 1.861a1 1 0 0 0-.29.649c-.015.25-.301 6.172 4.291 10.766C11.305 20.707 16.323 21 17.705 21c.202 0 .326-.006.359-.008a.992.992 0 0 0 .648-.291l1.86-2.171a.997.997 0 0 0-.085-1.39"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneCall(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 10.999h2C22 5.869 18.127 2 12.99 2v2C17.052 4 20 6.943 20 10.999"/><path fill="currentColor" d="M13 8c2.103 0 3 .897 3 3h2c0-3.225-1.775-5-5-5zm3.422 5.443a1.001 1.001 0 0 0-1.391.043l-2.393 2.461c-.576-.11-1.734-.471-2.926-1.66c-1.192-1.193-1.553-2.354-1.66-2.926l2.459-2.394a1 1 0 0 0 .043-1.391L6.859 3.513a1 1 0 0 0-1.391-.087l-2.17 1.861a1 1 0 0 0-.29.649c-.015.25-.301 6.172 4.291 10.766C11.305 20.707 16.323 21 17.705 21c.202 0 .326-.006.359-.008a.992.992 0 0 0 .648-.291l1.86-2.171a1 1 0 0 0-.086-1.391z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneIncoming(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.793 6.793L13 4v7h7l-2.793-2.793l4.5-4.5l-1.414-1.414z"/><path fill="currentColor" d="M16.422 13.443a1.001 1.001 0 0 0-1.391.043l-2.392 2.461c-.576-.11-1.734-.471-2.926-1.66c-1.192-1.193-1.553-2.354-1.66-2.926l2.459-2.394a1 1 0 0 0 .043-1.391L6.86 3.513a1 1 0 0 0-1.391-.087l-2.17 1.861a1.001 1.001 0 0 0-.291.649c-.015.25-.301 6.172 4.291 10.766C11.305 20.707 16.324 21 17.705 21c.203 0 .326-.006.359-.008a.99.99 0 0 0 .648-.291l1.861-2.171a1.001 1.001 0 0 0-.086-1.391z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneOff(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.17 13.42a5.24 5.24 0 0 1-.93-2.06L10.7 9a1 1 0 0 0 0-1.39l-3.65-4.1a1 1 0 0 0-1.4-.08L3.48 5.29a1 1 0 0 0-.29.65a15.25 15.25 0 0 0 3.54 9.92l-4.44 4.43l1.42 1.42l18-18l-1.42-1.42zm7.44.02a1 1 0 0 0-1.39.05L12.82 16a4.07 4.07 0 0 1-.51-.14l-2.66 2.61A15.46 15.46 0 0 0 17.89 21h.36a1 1 0 0 0 .65-.29l1.86-2.17a1 1 0 0 0-.09-1.39z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneOutgoing(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16.793 5.793l-4.5 4.5l1.414 1.414l4.5-4.5L21 10V3h-7z"/><path fill="currentColor" d="M16.422 13.446a1.001 1.001 0 0 0-1.391.043l-2.393 2.461c-.576-.11-1.734-.471-2.926-1.66c-1.192-1.193-1.553-2.354-1.66-2.926l2.459-2.394a1 1 0 0 0 .043-1.391L6.859 3.516a1 1 0 0 0-1.391-.087L3.299 5.29a.996.996 0 0 0-.291.648c-.015.25-.301 6.172 4.291 10.766c4.006 4.006 9.024 4.299 10.406 4.299c.202 0 .326-.006.359-.008a.992.992 0 0 0 .648-.291l1.86-2.171a1 1 0 0 0-.086-1.391z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhotoAlbum(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 2H6c-1.206 0-3 .799-3 3v14c0 2.201 1.794 3 3 3h15v-2H6.012C5.55 19.988 5 19.806 5 19s.55-.988 1.012-1H21V3a1 1 0 0 0-1-1M9.503 5a1.503 1.503 0 1 1 0 3.006a1.503 1.503 0 0 1 0-3.006M12 13H7l3-3l1.5 1.399L14.5 8l3.5 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Piano(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 3H5C3.346 3 2 4.346 2 6v12c0 1.654 1.346 3 3 3h14c1.654 0 3-1.346 3-3V6c0-1.654-1.346-3-3-3m0 16H5a1 1 0 0 1-1-1v-6h2v4h2v-4h1v4h2v-4h1v4h2v-4h2v4h2v-4h2v6a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PieChart(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m0 10V5c4 0 7 3 7 7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PieChartAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.071 4.929A9.97 9.97 0 0 0 12 2a9.936 9.936 0 0 0-7.071 2.929C3.04 6.818 2 9.33 2 12s1.04 5.182 2.929 7.071C6.818 20.96 9.33 22 12 22s5.182-1.04 7.071-2.929A9.936 9.936 0 0 0 22 12a9.97 9.97 0 0 0-2.929-7.071m-1.414 12.728C16.146 19.168 14.137 20 12 20s-4.146-.832-5.657-2.343C4.832 16.146 4 14.137 4 12s.832-4.146 2.343-5.657A7.948 7.948 0 0 1 12 4v8h8a7.948 7.948 0 0 1-2.343 5.657"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PieChartAltTwo(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 2.051V11h8.949c-.47-4.717-4.232-8.479-8.949-8.949m4.969 17.953c2.189-1.637 3.694-4.14 3.98-7.004h-8.183z"/><path fill="currentColor" d="M11 12V2.051C5.954 2.555 2 6.824 2 12c0 5.514 4.486 10 10 10a9.93 9.93 0 0 0 4.255-.964s-5.253-8.915-5.254-9.031A.02.02 0 0 0 11 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pin(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 11.586V6h2V4a2 2 0 0 0-2-2H9a2 2 0 0 0-2 2v2h2v5.586l-2.707 1.707A.996.996 0 0 0 6 14v2a1 1 0 0 0 1 1h4v3l1 2l1-2v-3h4a1 1 0 0 0 1-1v-2a.996.996 0 0 0-.293-.707z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pizza(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.76 2.021a.995.995 0 0 0-.989.703L3.579 19.166a1 1 0 0 0 1.255 1.255l16.442-5.192a.991.991 0 0 0 .702-.988C21.6 7.666 16.334 2.4 9.76 2.021M10 16a2 2 0 1 1 .001-4.001A2 2 0 0 1 10 16m6-2a2 2 0 1 1 .001-4.001A2 2 0 0 1 16 14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plane(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 16.21v-1.895L14 8V4a2 2 0 0 0-4 0v4.105L2 14.42v1.789l8-2.81V18l-3 2v2l5-2l5 2v-2l-3-2v-4.685z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlaneAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.414 13.778L2 15.192l4.949 2.121l2.122 4.95l1.414-1.414l-.707-3.536L13.091 14l3.61 7.704l1.339-1.339l-1.19-10.123l2.828-2.829a2 2 0 1 0-2.828-2.828l-2.903 2.903L3.824 6.297L2.559 7.563l7.644 3.67l-3.253 3.253z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlaneLand(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.842 15.296a1.61 1.61 0 0 0 1.892-1.189v-.001a1.609 1.609 0 0 0-1.177-1.949l-4.576-1.133L9.825 4.21l-2.224-.225l2.931 6.589l-4.449-.449l-2.312-3.829l-1.38.31l1.24 5.52zM3 18h18v2H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlaneTakeOff(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 18h18v2H3zm18.509-9.473a1.61 1.61 0 0 0-2.036-1.019L15 9L7 6L5 7l6 4l-4 2l-4-2l-1 1l4 4l14.547-5.455a1.611 1.611 0 0 0 .962-2.018"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Planet(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.165 15.582c4.587-4.073 8.141-9.424 6.057-11.771c-.661-.744-1.584-1.089-2.575-.983c-.832.091-1.687.502-2.549 1.192a8.922 8.922 0 0 0-8.712.282a8.917 8.917 0 0 0-4.109 5.515a8.892 8.892 0 0 0 .306 5.325c-1.065 1.203-2.054 3.677-.823 5.063c.517.581 1.257.841 2.147.841c2.707 0 6.808-2.399 10.258-5.464m3.699-10.767c.358-.034.632.064.861.323c.231.261.169.946-.252 1.929a9.059 9.059 0 0 0-1.617-1.853c.431-.262.776-.373 1.008-.399M4.633 17.118a8.979 8.979 0 0 0 1.568 1.737c-1.025.303-1.714.283-1.945.021c-.217-.243.002-1.069.377-1.758m16.31-5.869c-1.215 1.797-2.906 3.671-4.778 5.333c-1.934 1.719-4.066 3.208-6.05 4.202a9.082 9.082 0 0 0 1.874.212a8.986 8.986 0 0 0 4.616-1.282a8.915 8.915 0 0 0 4.338-8.465"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Playlist(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 16.493C13 18.427 14.573 20 16.507 20s3.507-1.573 3.507-3.507c0-.177-.027-.347-.053-.517H20V6h2V4h-3a1 1 0 0 0-1 1v8.333a3.465 3.465 0 0 0-1.493-.346A3.51 3.51 0 0 0 13 16.493M2 5h14v2H2z"/><path fill="currentColor" d="M2 9h14v2H2zm0 4h9v2H2zm0 4h9v2H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plug(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 8h2v5c0 2.206 1.794 4 4 4h2v5h2v-5h2c2.206 0 4-1.794 4-4V8h2V6H3zm4-6h2v3H7zm8 0h2v3h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusCircle(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m5 11h-4v4h-2v-4H7v-2h4V7h2v4h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusSquare(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 21h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2m2-10h4V7h2v4h4v2h-4v4h-2v-4H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pointer(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.978 13.21a1 1 0 0 0-.396-1.024l-14-10a.999.999 0 0 0-1.575.931l2 17a1 1 0 0 0 1.767.516l3.612-4.416l3.377 5.46l1.701-1.052l-3.357-5.428l6.089-1.218a.995.995 0 0 0 .782-.769"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Polygon(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.707 7.293l-5-5A.996.996 0 0 0 16 2H8a.996.996 0 0 0-.707.293l-5 5A.996.996 0 0 0 2 8v8c0 .266.105.52.293.707l5 5A.996.996 0 0 0 8 22h8c.266 0 .52-.105.707-.293l5-5A.996.996 0 0 0 22 16V8a.996.996 0 0 0-.293-.707"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Popsicle(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 4a7 7 0 0 0-9.9 0l-7.82 7.82a1 1 0 0 0 0 1.41l3.54 3.54l-3.54 3.53l1.42 1.42l3.53-3.54l3.54 3.54a1 1 0 0 0 1.41 0L20 13.94A7 7 0 0 0 20 4m-2.7 2.7a3.33 3.33 0 0 0-4.6 0l-1.06-1.06a4.76 4.76 0 0 1 6.72 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Printer(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 7h-1V2H6v5H5a3 3 0 0 0-3 3v7a2 2 0 0 0 2 2h2v3h12v-3h2a2 2 0 0 0 2-2v-7a3 3 0 0 0-3-3M8 4h8v3H8zm0 16v-4h8v4zm11-8h-4v-2h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PurchaseTag(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.586 2.586A2 2 0 0 0 11.172 2H4a2 2 0 0 0-2 2v7.172a2 2 0 0 0 .586 1.414l8 8a2 2 0 0 0 2.828 0l7.172-7.172a2 2 0 0 0 0-2.828zM7 9a2 2 0 1 1 .001-4.001A2 2 0 0 1 7 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PurchaseTagAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.707 2.293A.996.996 0 0 0 11 2H6a.996.996 0 0 0-.707.293l-3 3A.996.996 0 0 0 2 6v5c0 .266.105.52.293.707l10 10a.997.997 0 0 0 1.414 0l8-8a.999.999 0 0 0 0-1.414zM8.353 10a1.647 1.647 0 1 1-.001-3.293A1.647 1.647 0 0 1 8.353 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pyramid(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.445 21.832a1 1 0 0 0 1.11 0l9-6A.998.998 0 0 0 21.8 14.4l-9-12c-.377-.504-1.223-.504-1.6 0l-9 12a1 1 0 0 0 .245 1.432zm8.12-7.078L12 19.798V4.667z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuoteAltLeft(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.5 10c-.223 0-.437.034-.65.065c.069-.232.14-.468.254-.68c.114-.308.292-.575.469-.844c.148-.291.409-.488.601-.737c.201-.242.475-.403.692-.604c.213-.21.492-.315.714-.463c.232-.133.434-.28.65-.35l.539-.222l.474-.197l-.485-1.938l-.597.144c-.191.048-.424.104-.689.171c-.271.05-.56.187-.882.312c-.318.142-.686.238-1.028.466c-.344.218-.741.4-1.091.692c-.339.301-.748.562-1.05.945c-.33.358-.656.734-.909 1.162c-.293.408-.492.856-.702 1.299c-.19.443-.343.896-.468 1.336c-.237.882-.343 1.72-.384 2.437c-.034.718-.014 1.315.028 1.747c.015.204.043.402.063.539l.025.168l.026-.006A4.5 4.5 0 1 0 6.5 10m11 0c-.223 0-.437.034-.65.065c.069-.232.14-.468.254-.68c.114-.308.292-.575.469-.844c.148-.291.409-.488.601-.737c.201-.242.475-.403.692-.604c.213-.21.492-.315.714-.463c.232-.133.434-.28.65-.35l.539-.222l.474-.197l-.485-1.938l-.597.144c-.191.048-.424.104-.689.171c-.271.05-.56.187-.882.312c-.317.143-.686.238-1.028.467c-.344.218-.741.4-1.091.692c-.339.301-.748.562-1.05.944c-.33.358-.656.734-.909 1.162c-.293.408-.492.856-.702 1.299c-.19.443-.343.896-.468 1.336c-.237.882-.343 1.72-.384 2.437c-.034.718-.014 1.315.028 1.747c.015.204.043.402.063.539l.025.168l.026-.006A4.5 4.5 0 1 0 17.5 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuoteAltRight(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.95 8.721l-.025-.168l-.026.006A4.5 4.5 0 1 0 17.5 14c.223 0 .437-.034.65-.065c-.069.232-.14.468-.254.68c-.114.308-.292.575-.469.844c-.148.291-.409.488-.601.737c-.201.242-.475.403-.692.604c-.213.21-.492.315-.714.463c-.232.133-.434.28-.65.35l-.539.222l-.474.197l.484 1.939l.597-.144c.191-.048.424-.104.689-.171c.271-.05.56-.187.882-.312c.317-.143.686-.238 1.028-.467c.344-.218.741-.4 1.091-.692c.339-.301.748-.562 1.05-.944c.33-.358.656-.734.909-1.162c.293-.408.492-.856.702-1.299c.19-.443.343-.896.468-1.336c.237-.882.343-1.72.384-2.437c.034-.718.014-1.315-.028-1.747a7.028 7.028 0 0 0-.063-.539m-11 0l-.025-.168l-.026.006A4.5 4.5 0 1 0 6.5 14c.223 0 .437-.034.65-.065c-.069.232-.14.468-.254.68c-.114.308-.292.575-.469.844c-.148.291-.409.488-.601.737c-.201.242-.475.403-.692.604c-.213.21-.492.315-.714.463c-.232.133-.434.28-.65.35l-.539.222c-.301.123-.473.195-.473.195l.484 1.939l.597-.144c.191-.048.424-.104.689-.171c.271-.05.56-.187.882-.312c.317-.143.686-.238 1.028-.467c.344-.218.741-.4 1.091-.692c.339-.301.748-.562 1.05-.944c.33-.358.656-.734.909-1.162c.293-.408.492-.856.702-1.299c.19-.443.343-.896.468-1.336c.237-.882.343-1.72.384-2.437c.034-.718.014-1.315-.028-1.747a7.571 7.571 0 0 0-.064-.537"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuoteLeft(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.691 6.292C5.094 4.771 7.217 4 10 4h1v2.819l-.804.161c-1.37.274-2.323.813-2.833 1.604A2.902 2.902 0 0 0 6.925 10H10a1 1 0 0 1 1 1v7c0 1.103-.897 2-2 2H3a1 1 0 0 1-1-1v-5l.003-2.919c-.009-.111-.199-2.741 1.688-4.789M20 20h-6a1 1 0 0 1-1-1v-5l.003-2.919c-.009-.111-.199-2.741 1.688-4.789C16.094 4.771 18.217 4 21 4h1v2.819l-.804.161c-1.37.274-2.323.813-2.833 1.604A2.902 2.902 0 0 0 17.925 10H21a1 1 0 0 1 1 1v7c0 1.103-.897 2-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuoteRight(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.309 17.708C22.196 15.66 22.006 13.03 22 13V5a1 1 0 0 0-1-1h-6c-1.103 0-2 .897-2 2v7a1 1 0 0 0 1 1h3.078a2.89 2.89 0 0 1-.429 1.396c-.508.801-1.465 1.348-2.846 1.624l-.803.16V20h1c2.783 0 4.906-.771 6.309-2.292m-11.007 0C11.19 15.66 10.999 13.03 10.993 13V5a1 1 0 0 0-1-1h-6c-1.103 0-2 .897-2 2v7a1 1 0 0 0 1 1h3.078a2.89 2.89 0 0 1-.429 1.396c-.508.801-1.465 1.348-2.846 1.624l-.803.16V20h1c2.783 0 4.906-.771 6.309-2.292"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuoteSingleLeft(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.203 6.98l.804-.161V4h-1c-2.784 0-4.906.771-6.309 2.292C6.81 8.34 7 10.97 7.006 11v8a1 1 0 0 0 1 1h7c1.103 0 2-.897 2-2v-7a1 1 0 0 0-1-1h-4.079c.022-.402.123-.912.429-1.396c.509-.801 1.466-1.347 2.847-1.624"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuoteSingleRight(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.804 17.02L8 17.18V20h1c2.783 0 4.906-.771 6.309-2.292C17.196 15.66 17.006 13.03 17 13V5a1 1 0 0 0-1-1H9c-1.103 0-2 .897-2 2v7a1 1 0 0 0 1 1h4.078a2.89 2.89 0 0 1-.429 1.396c-.507.801-1.464 1.347-2.845 1.624"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Radiation(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.912 8.531L7.121 3.877a.501.501 0 0 0-.704-.166a9.982 9.982 0 0 0-4.396 7.604a.505.505 0 0 0 .497.528l5.421.09a4.042 4.042 0 0 1 1.973-3.402m8.109-4.51a.504.504 0 0 0-.729.151L14.499 8.83a4.03 4.03 0 0 1 1.546 3.112l5.419-.09a.507.507 0 0 0 .499-.53a9.986 9.986 0 0 0-3.942-7.301m-4.067 11.511a4.015 4.015 0 0 1-1.962.526a4.016 4.016 0 0 1-1.963-.526l-2.642 4.755a.5.5 0 0 0 .207.692A9.948 9.948 0 0 0 11.992 22a9.94 9.94 0 0 0 4.396-1.021a.5.5 0 0 0 .207-.692z"/><circle cx="12" cy="12" r="3" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Radio(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.249 5.025l-7.897-2.962l-.703 1.873L14.484 5H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V7c0-1.02-.766-1.851-1.751-1.975M10 17H6v-2h4zm6.5 1a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5m3.5-7H4V7h16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Receipt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 12v6a1 1 0 0 1-2 0V4a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v14c0 1.654 1.346 3 3 3h14c1.654 0 3-1.346 3-3v-6zm-6-1v2H6v-2zM6 9V7h8v2zm8 6v2h-3v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rectangle(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 20h18a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Registered(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 9h-3v2h3a1 1 0 0 0 0-2"/><path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m2.13 15l-2.67-4H10v4H8V7h5a3 3 0 0 1 .79 5.88L16.54 17z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rename(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 5.995h-1v12h1c1.103 0 2-.897 2-2v-8c0-1.102-.897-2-2-2"/><path fill="currentColor" d="M17 17.995V4h2.995V2h-8v2H15v1.995H4c-1.103 0-2 .897-2 2v8c0 1.103.897 2 2 2h11V20h-3.005v2h8v-2H17zm-11-4v-4h9v4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Report(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20 8l-6-6H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2zM9 19H7v-9h2zm4 0h-2v-6h2zm4 0h-2v-3h2zM14 9h-1V4l5 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RewindCircle(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.485 2 12s4.486 10 10 10c5.515 0 10-4.485 10-10S17.515 2 12 2m5 14l-6-4v4l-6-4l6-4v4l6-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightArrow(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.536 21.886a1.004 1.004 0 0 0 1.033-.064l13-9a1 1 0 0 0 0-1.644l-13-9A1 1 0 0 0 5 3v18a1 1 0 0 0 .536.886"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightArrowAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m19 12l-7-6v5H6v2h6v5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightArrowCircle(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m0 15v-4H7v-2h5V7l5 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightArrowSquare(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 5v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2m4 6h5V7l5 5l-5 5v-4H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightDownArrowCircle(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.929 4.929c-3.898 3.899-3.898 10.244 0 14.143c3.899 3.898 10.243 3.898 14.143 0c3.898-3.899 3.898-10.244 0-14.143c-3.9-3.899-10.244-3.899-14.143 0m10.606 10.607h-7.07l2.828-2.829l-3.535-3.536l1.414-1.414l3.535 3.536l2.828-2.829z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightTopArrowCircle(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.071 4.929c-3.899-3.898-10.243-3.898-14.143 0c-3.898 3.899-3.898 10.244 0 14.143c3.899 3.898 10.243 3.898 14.143 0c3.899-3.9 3.899-10.244 0-14.143m-3.536 10.607l-2.828-2.829l-3.535 3.536l-1.414-1.414l3.535-3.536l-2.828-2.829h7.07z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rocket(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.78 15.84S18.64 13 19.61 12c3.07-3 1.54-9.18 1.54-9.18S15 1.29 12 4.36C9.66 6.64 8.14 8.22 8.14 8.22S4.3 7.42 2 9.72L14.25 22c2.3-2.33 1.53-6.16 1.53-6.16m-1.5-9a2 2 0 0 1 2.83 0a2 2 0 1 1-2.83 0M3 21a7.81 7.81 0 0 0 5-2l-3-3c-2 1-2 5-2 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ruler(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.875 7H3.125C1.953 7 1 7.897 1 9v6c0 1.103.953 2 2.125 2h17.75C22.047 17 23 16.103 23 15V9c0-1.103-.953-2-2.125-2M7 12H5V9h2zm4 1H9V9h2zm4-1h-2V9h2zm4 1h-2V9h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sad(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m-5 8.5a1.5 1.5 0 1 1 3.001.001A1.5 1.5 0 0 1 7 10.5M8 17s1-3 4-3s4 3 4 3zm7.493-5.014a1.494 1.494 0 1 1 .001-2.987a1.494 1.494 0 0 1-.001 2.987"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Save(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 21h14a2 2 0 0 0 2-2V8l-5-5H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2M7 5h4v2h2V5h2v4H7zm0 8h10v6H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func School(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 10h-2V4h1V2H4v2h1v6H3a1 1 0 0 0-1 1v9h20v-9a1 1 0 0 0-1-1m-7 8v-4h-4v4H7V4h10v14z"/><path fill="currentColor" d="M9 6h2v2H9zm4 0h2v2h-2zm-4 4h2v2H9zm4 0h2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Search(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 2c-4.411 0-8 3.589-8 8s3.589 8 8 8a7.952 7.952 0 0 0 4.897-1.688l4.396 4.396l1.414-1.414l-4.396-4.396A7.952 7.952 0 0 0 18 10c0-4.411-3.589-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchAltTwo(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 16c1.763 0 3.37-.66 4.603-1.739l1.337 2.8s.275.224.653.596c.387.363.896.854 1.384 1.367l1.358 1.392l.604.646l2.121-2.121l-.646-.604l-1.392-1.358a35.13 35.13 0 0 1-1.367-1.384c-.372-.378-.596-.653-.596-.653l-2.8-1.337A6.967 6.967 0 0 0 16 9c0-3.859-3.141-7-7-7S2 5.141 2 9s3.141 7 7 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SelectMultiple(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 2H8a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2m-6.933 12.481l-3.274-3.274l1.414-1.414l1.726 1.726l4.299-5.159l1.537 1.281z"/><path fill="currentColor" d="M4 22h11v-2H4V8H2v12c0 1.103.897 2 2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Send(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.426 11.095l-17-8A1 1 0 0 0 3.03 4.242l1.212 4.849L12 12l-7.758 2.909l-1.212 4.849a.998.998 0 0 0 1.396 1.147l17-8a1 1 0 0 0 0-1.81"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Server(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 3H4a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2m-5 5h-2V6h2zm4 0h-2V6h2zm1 5H4a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2m-5 5h-2v-2h2zm4 0h-2v-2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shapes(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.867 2.504c-.355-.624-1.381-.623-1.736 0l-3.999 7A1 1 0 0 0 13 11h8a1.001 1.001 0 0 0 .868-1.496zM3 22h7a1 1 0 0 0 1-1v-7a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v7a1 1 0 0 0 1 1m14.5-9c-2.481 0-4.5 2.019-4.5 4.5s2.019 4.5 4.5 4.5s4.5-2.019 4.5-4.5s-2.019-4.5-4.5-4.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Share(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 6.914V2.586L6.293 7.293l-3.774 3.774l3.841 3.201L11 18.135V13.9c8.146-.614 11 4.1 11 4.1c0-2.937-.242-5.985-2.551-8.293C16.765 7.022 12.878 6.832 11 6.914"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 12c0 1.654 1.346 3 3 3c.794 0 1.512-.315 2.049-.82l5.991 3.424c-.018.13-.04.26-.04.396c0 1.654 1.346 3 3 3s3-1.346 3-3s-1.346-3-3-3c-.794 0-1.512.315-2.049.82L8.96 12.397c.018-.131.04-.261.04-.397s-.022-.266-.04-.397l5.991-3.423c.537.505 1.255.82 2.049.82c1.654 0 3-1.346 3-3s-1.346-3-3-3s-3 1.346-3 3c0 .136.022.266.04.397L8.049 9.82A2.982 2.982 0 0 0 6 9c-1.654 0-3 1.346-3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shield(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.496 6.106l-7.973-4a.997.997 0 0 0-.895-.002l-8.027 4c-.297.15-.502.437-.544.767c-.013.097-1.145 9.741 8.541 15.008a.995.995 0 0 0 .969-.009c9.307-5.259 8.514-14.573 8.476-14.967a1 1 0 0 0-.547-.797"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldAltTwo(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.881 5.223a.496.496 0 0 0-.747-.412c-.672.392-1.718.898-2.643.898c-.421 0-.849-.064-1.289-.198a5.712 5.712 0 0 1-.808-.309c-1.338-.639-2.567-1.767-3.696-2.889a1.008 1.008 0 0 0-.698-.29a1.008 1.008 0 0 0-.698.29c-1.129 1.122-2.358 2.25-3.696 2.889h-.001a5.655 5.655 0 0 1-.807.309c-.44.134-.869.198-1.289.198c-.925 0-1.971-.507-2.643-.898a.496.496 0 0 0-.747.412c-.061 1.538-.077 4.84.688 7.444c1.399 4.763 4.48 7.976 8.91 9.292l.14.041l.14-.014V22v-.014H12l.143.014l.14-.041c4.43-1.316 7.511-4.529 8.91-9.292c.765-2.604.748-5.906.688-7.444"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldMinus(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.43 5.76l-8-3.56a1 1 0 0 0-.82 0l-8 3.56a1 1 0 0 0-.59.9c0 2.37.44 10.8 8.51 15.11a1 1 0 0 0 1 0c8-4.3 8.58-12.71 8.57-15.1a1 1 0 0 0-.67-.91m-4.43 7H8v-2h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldPlus(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.43 5.76l-8-3.56a1 1 0 0 0-.82 0l-8 3.56a1 1 0 0 0-.59.9c0 2.37.44 10.8 8.51 15.11a1 1 0 0 0 1 0c8-4.3 8.58-12.71 8.57-15.1a1 1 0 0 0-.67-.91m-4.43 7h-3v3h-2v-3H8v-2h3v-3h2v3h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldX(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.492 21.771c.294.157.663.157.957-.001c8.012-4.304 8.581-12.713 8.574-15.104a.988.988 0 0 0-.596-.903l-8.051-3.565a1.005 1.005 0 0 0-.813.001L3.57 5.765a.988.988 0 0 0-.592.891c-.034 2.379.445 10.806 8.514 15.115M8.293 9.707l1.414-1.414L12 10.586l2.293-2.293l1.414 1.414L13.414 12l2.293 2.293l-1.414 1.414L12 13.414l-2.293 2.293l-1.414-1.414L10.586 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ship(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.997 20c-.899 0-1.288-.311-1.876-.781c-.68-.543-1.525-1.219-3.127-1.219c-1.601 0-2.446.676-3.125 1.22c-.587.469-.975.78-1.874.78c-.897 0-1.285-.311-1.872-.78C4.444 18.676 3.601 18 2 18v2c.898 0 1.286.311 1.873.78c.679.544 1.523 1.22 3.122 1.22c1.601 0 2.445-.676 3.124-1.219c.588-.47.976-.781 1.875-.781c.9 0 1.311.328 1.878.781c.679.543 1.524 1.219 3.125 1.219s2.446-.676 3.125-1.219C20.689 20.328 21.1 20 22 20v-2c-1.602 0-2.447.676-3.127 1.219c-.588.47-.977.781-1.876.781M6 8.5L4 9l2 8h.995c1.601 0 2.445-.676 3.124-1.219c.588-.47.976-.781 1.875-.781c.9 0 1.311.328 1.878.781c.679.543 1.524 1.219 3.125 1.219H18l.027-.107l.313-1.252L20 9l-2-.5V5.001a1 1 0 0 0-.804-.981L13 3.181V2h-2v1.181l-4.196.839A1 1 0 0 0 6 5.001zm2-2.681l4-.8l4 .8V8l-4-1l-4 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shocked(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m-5 8.5a1.5 1.5 0 1 1 3.001.001A1.5 1.5 0 0 1 7 10.5m5 7.5c-1.657 0-3-1.119-3-2.5s1.343-2.5 3-2.5s3 1.119 3 2.5s-1.343 2.5-3 2.5m3.493-6.014a1.494 1.494 0 1 1 .001-2.987a1.494 1.494 0 0 1-.001 2.987"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBag(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 22h14a2 2 0 0 0 2-2V9a1 1 0 0 0-1-1h-3v-.777c0-2.609-1.903-4.945-4.5-5.198A5.005 5.005 0 0 0 7 7v1H4a1 1 0 0 0-1 1v11a2 2 0 0 0 2 2m12-12v2h-2v-2zM9 7c0-1.654 1.346-3 3-3s3 1.346 3 3v1H9zm-2 3h2v2H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBagAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 4H3a1 1 0 0 0-1 1v14a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V5a1 1 0 0 0-1-1m-9 9c-3.309 0-6-2.691-6-6h2c0 2.206 1.794 4 4 4s4-1.794 4-4h2c0 3.309-2.691 6-6 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBags(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 20h2V10a1 1 0 0 1 1-1h12V7a1 1 0 0 0-1-1h-3.051c-.252-2.244-2.139-4-4.449-4S6.303 3.756 6.051 6H3a1 1 0 0 0-1 1v11a2 2 0 0 0 2 2m6.5-16c1.207 0 2.218.86 2.45 2h-4.9c.232-1.14 1.243-2 2.45-2"/><path fill="currentColor" d="M21 11H9a1 1 0 0 0-1 1v8a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-8a1 1 0 0 0-1-1m-6 7c-2.757 0-5-2.243-5-5h2c0 1.654 1.346 3 3 3s3-1.346 3-3h2c0 2.757-2.243 5-5 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Show(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 5c-7.633 0-9.927 6.617-9.948 6.684L1.946 12l.105.316C2.073 12.383 4.367 19 12 19s9.927-6.617 9.948-6.684l.106-.316l-.105-.316C21.927 11.617 19.633 5 12 5m0 11c-2.206 0-4-1.794-4-4s1.794-4 4-4s4 1.794 4 4s-1.794 4-4 4"/><path fill="currentColor" d="M12 10c-1.084 0-2 .916-2 2s.916 2 2 2s2-.916 2-2s-.916-2-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shower(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 18.33A6.78 6.78 0 0 0 19.5 15a6.73 6.73 0 0 0-1.5 3.33a1.51 1.51 0 1 0 3 0m-10 2A6.78 6.78 0 0 0 9.5 17A6.73 6.73 0 0 0 8 20.33A1.59 1.59 0 0 0 9.5 22a1.59 1.59 0 0 0 1.5-1.67m5 0A6.78 6.78 0 0 0 14.5 17a6.73 6.73 0 0 0-1.5 3.33A1.59 1.59 0 0 0 14.5 22a1.59 1.59 0 0 0 1.5-1.67m-10-2A6.78 6.78 0 0 0 4.5 15A6.73 6.73 0 0 0 3 18.33A1.59 1.59 0 0 0 4.5 20A1.59 1.59 0 0 0 6 18.33M2 12h20v2H2zm11-7.93V2h-2v2.07A8 8 0 0 0 4.07 11h15.86A8 8 0 0 0 13 4.07"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SkipNextCircle(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m4 14h-2v-4l-6 4V8l6 4V8h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SkipPreviousCircle(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10c5.515 0 10-4.486 10-10S17.515 2 12 2m4 14l-6-4v4H8V8h2v4l6-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Skull(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C5.505 2 2 6.637 2 11c0 2.129 1.009 3.979 3 5.508V21h3v-3h2v3h4v-3h2v3h3v-4.493c1.991-1.528 3-3.379 3-5.507c0-4.363-3.505-9-10-9M8 13c-1.121 0-2-1.098-2-2.5S6.879 8 8 8s2 1.098 2 2.5S9.121 13 8 13m8 0c-1.121 0-2-1.098-2-2.5S14.879 8 16 8s2 1.098 2 2.5s-.879 2.5-2 2.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sleepy(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m-4 9.01l-2-.02C6.017 9.386 7.095 7 10 7v2c-1.924 0-1.998 1.805-2 2.01M12 18c-1.657 0-3-1.119-3-2.5s1.343-2.5 3-2.5s3 1.119 3 2.5s-1.343 2.5-3 2.5m5-7l-1 .008C15.992 10.536 15.826 9 14 9V7c2.935 0 4 2.393 4 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Slideshow(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 3H4c-1.103 0-2 .897-2 2v10c0 1.103.897 2 2 2h7v3H8v2h8v-2h-3v-3h7c1.103 0 2-.897 2-2V5c0-1.103-.897-2-2-2M10 13V7l5 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Smile(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22c5.514 0 10-4.486 10-10S17.514 2 12 2S2 6.486 2 12s4.486 10 10 10m3.493-13a1.494 1.494 0 1 1-.001 2.987A1.494 1.494 0 0 1 15.493 9m-4.301 6.919a4.108 4.108 0 0 0 1.616 0c.253-.052.505-.131.75-.233c.234-.1.464-.224.679-.368c.208-.142.407-.306.591-.489c.183-.182.347-.381.489-.592l1.658 1.117a6.027 6.027 0 0 1-1.619 1.621a6.003 6.003 0 0 1-2.149.904a6.116 6.116 0 0 1-2.414-.001a5.919 5.919 0 0 1-2.148-.903a6.078 6.078 0 0 1-1.621-1.622l1.658-1.117c.143.211.307.41.488.59a3.988 3.988 0 0 0 2.022 1.093M8.5 9a1.5 1.5 0 1 1-.001 3.001A1.5 1.5 0 0 1 8.5 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.227 11h11.547c.862 0 1.32-1.02.747-1.665L12.748 2.84a.998.998 0 0 0-1.494 0L5.479 9.335C4.906 9.98 5.364 11 6.227 11m5.026 10.159a.998.998 0 0 0 1.494 0l5.773-6.495c.574-.644.116-1.664-.747-1.664H6.227c-.862 0-1.32 1.02-.747 1.665z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spa(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 16.114c-3.998-5.951-8.574-7.043-8.78-7.09L2 8.75V10c0 7.29 3.925 12 10 12c5.981 0 10-4.822 10-12V8.75l-1.22.274c-.206.047-4.782 1.139-8.78 7.09"/><path fill="currentColor" d="M11.274 3.767c-1.799 1.898-2.84 3.775-3.443 5.295c1.329.784 2.781 1.943 4.159 3.685c1.364-1.76 2.826-2.925 4.17-3.709c-.605-1.515-1.646-3.383-3.435-5.271L12 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Speaker(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="12" cy="15" r="3" fill="currentColor"/><path fill="currentColor" d="M18 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2m-6 2a2 2 0 1 1-2 2a2 2 0 0 1 2-2m0 16a5 5 0 1 1 5-5a5 5 0 0 1-5 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SprayCan(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.003 3h2v2h-2zM16 3h2v2h-2zm0 3h2v2h-2zm3-3h2v2h-2zm0 3h2v2h-2zm0 3h2v2h-2zM4.012 12v9a1 1 0 0 0 1 1H13a1 1 0 0 0 1-1v-9a4 4 0 0 0-4-4H8.012a4 4 0 0 0-4 4M7.003 2h4v4h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spreadsheet(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 5v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2m7 2h8v2h-8zm0 4h8v2h-8zm0 4h8v2h-8zM6 7h2v2H6zm0 4h2v2H6zm0 4h2v2H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Square(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 3H4a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareRounded(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2H7C4.243 2 2 4.243 2 7v10c0 2.757 2.243 5 5 5h10c2.757 0 5-2.243 5-5V7c0-2.757-2.243-5-5-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Star(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.947 9.179a1.001 1.001 0 0 0-.868-.676l-5.701-.453l-2.467-5.461a.998.998 0 0 0-1.822-.001L8.622 8.05l-5.701.453a1 1 0 0 0-.619 1.713l4.213 4.107l-1.49 6.452a1 1 0 0 0 1.53 1.057L12 18.202l5.445 3.63a1.001 1.001 0 0 0 1.517-1.106l-1.829-6.4l4.536-4.082c.297-.268.406-.686.278-1.065"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarHalf(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.025 20.775A.998.998 0 0 0 6 22a1 1 0 0 0 .555-.168L12 18.202l5.445 3.63a1.001 1.001 0 0 0 1.517-1.106l-1.829-6.4l4.536-4.082a1 1 0 0 0-.59-1.74l-5.701-.454l-2.467-5.461a.998.998 0 0 0-1.822-.001L8.622 8.05l-5.701.453a1 1 0 0 0-.619 1.713l4.214 4.107zM12 5.429l2.042 4.521l.588.047h.001l3.972.315l-3.271 2.944l-.001.002l-.463.416l.171.597v.003l1.253 4.385L12 15.798z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sticker(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.5 11c-4.136 0-7.5 3.364-7.5 7.5c0 .871.157 1.704.432 2.482l9.551-9.551A7.462 7.462 0 0 0 18.5 11"/><path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12c0 4.583 3.158 8.585 7.563 9.69A9.431 9.431 0 0 1 9 18.5C9 13.262 13.262 9 18.5 9c1.12 0 2.191.205 3.19.563C20.585 5.158 16.583 2 12 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stopwatch(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 5c-4.411 0-8 3.589-8 8s3.589 8 8 8s8-3.589 8-8s-3.589-8-8-8m1 8h-2V8h2zM9 2h6v2H9zm9.707 2.293l2 2l-1.414 1.414l-2-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Store(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.999 8a.997.997 0 0 0-.143-.515L19.147 2.97A2.01 2.01 0 0 0 17.433 2H6.565c-.698 0-1.355.372-1.714.971L2.142 7.485A.997.997 0 0 0 1.999 8c0 1.005.386 1.914 1 2.618V20a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1v-5h4v5a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-9.382c.614-.704 1-1.613 1-2.618m-2.016.251A2.002 2.002 0 0 1 17.999 10c-1.103 0-2-.897-2-2c0-.068-.025-.128-.039-.192l.02-.004L15.219 4h2.214zm-9.977-.186L10.818 4h2.361l.813 4.065C13.957 9.138 13.079 10 11.999 10s-1.958-.862-1.993-1.935M6.565 4h2.214l-.76 3.804l.02.004c-.015.064-.04.124-.04.192c0 1.103-.897 2-2 2a2.002 2.002 0 0 1-1.984-1.749zm3.434 12h-4v-3h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StoreAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 2H5C3.346 2 2 3.346 2 5v2.831c0 1.053.382 2.01 1 2.746V20a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1v-5h4v5a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-9.424c.618-.735 1-1.692 1-2.746V5c0-1.654-1.346-3-3-3m1 3v2.831c0 1.14-.849 2.112-1.891 2.167L18 10c-1.103 0-2-.897-2-2V4h3c.552 0 1 .449 1 1M10 8V4h4v4c0 1.103-.897 2-2 2s-2-.897-2-2M4 5c0-.551.448-1 1-1h3v4c0 1.103-.897 2-2 2l-.109-.003C4.849 9.943 4 8.971 4 7.831zm6 11H6v-3h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sun(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.995 12c0 2.761 2.246 5.007 5.007 5.007s5.007-2.246 5.007-5.007s-2.246-5.007-5.007-5.007S6.995 9.239 6.995 12M11 19h2v3h-2zm0-17h2v3h-2zm-9 9h3v2H2zm17 0h3v2h-3zM5.637 19.778l-1.414-1.414l2.121-2.121l1.414 1.414zM16.242 6.344l2.122-2.122l1.414 1.414l-2.122 2.122zM6.344 7.759L4.223 5.637l1.415-1.414l2.12 2.122zm13.434 10.605l-1.414 1.414l-2.122-2.122l1.414-1.414z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sushi(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<ellipse cx="12.07" cy="7" fill="currentColor" rx="3" ry="1.71"/><path fill="currentColor" d="M12.07 22c4.48 0 8-2.2 8-5V7c0-2.8-3.52-5-8-5s-8 2.2-8 5v10c0 2.8 3.51 5 8 5m0-18c3.53 0 6 1.58 6 3a2 2 0 0 1-.29.87c-.68 1-2.53 2-5 2.12h-1.39C8.88 9.83 7 8.89 6.35 7.84a2.16 2.16 0 0 1-.28-.76V7c0-1.42 2.46-3 6-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tshirt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.316 4.055C19.556 3.478 15 1.985 15 2a3 3 0 1 1-6 0c0-.015-4.556 1.478-6.317 2.055A.992.992 0 0 0 2 5.003v3.716a1 1 0 0 0 1.242.97L6 9v12a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1V9l2.758.689A1 1 0 0 0 22 8.719V5.003a.992.992 0 0 0-.684-.948"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tachometer(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 4C6.486 4 2 8.486 2 14a9.89 9.89 0 0 0 1.051 4.445c.17.34.516.555.895.555h16.107c.379 0 .726-.215.896-.555A9.89 9.89 0 0 0 22 14c0-5.514-4.486-10-10-10m5.022 5.022L13.06 15.06a1.53 1.53 0 0 1-2.121.44a1.53 1.53 0 0 1 0-2.561l6.038-3.962a.033.033 0 0 1 .045.01a.034.034 0 0 1 0 .035"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tag(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.842 6.218a1.977 1.977 0 0 0-.424-.628A1.99 1.99 0 0 0 20 5H8c-.297 0-.578.132-.769.359l-5 6c-.309.371-.309.91 0 1.281l5 6c.191.228.472.36.769.36h12a1.977 1.977 0 0 0 1.41-.582A1.99 1.99 0 0 0 22 17V7c0-.266-.052-.525-.158-.782"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.868 11.504l-4-7A1 1 0 0 0 17 4H3a1 1 0 0 0-.868 1.496L5.849 12l-3.717 6.504A1 1 0 0 0 3 20h14a1 1 0 0 0 .868-.504l4-7a.998.998 0 0 0 0-.992"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagX(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.842 6.218a1.977 1.977 0 0 0-.424-.628A1.99 1.99 0 0 0 20 5H8c-.297 0-.578.132-.769.359l-5 6c-.309.371-.309.91 0 1.281l5 6c.191.228.472.36.769.36h12a1.977 1.977 0 0 0 1.41-.582A1.99 1.99 0 0 0 22 17V7c0-.266-.052-.525-.158-.782m-4.135 8.075l-1.414 1.414L14 13.414l-2.293 2.293l-1.414-1.414L12.586 12l-2.293-2.293l1.414-1.414L14 10.586l2.293-2.293l1.414 1.414L15.414 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Taxi(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.772 10.155l-1.368-4.104A2.995 2.995 0 0 0 16.559 4H14V2h-4v2H7.441a2.995 2.995 0 0 0-2.845 2.051l-1.368 4.104A2 2 0 0 0 2 12v5c0 .738.404 1.376 1 1.723V21a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-2h12v2a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-2.277A1.99 1.99 0 0 0 22 17v-5a2 2 0 0 0-1.228-1.845M7.441 6h9.117c.431 0 .813.274.949.684L18.613 10H5.387l1.105-3.316A1 1 0 0 1 7.441 6M5.5 16a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 5.5 16m13 0a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 18.5 16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TennisBall(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.929 19.071a9.953 9.953 0 0 0 6.692 2.906c-.463-2.773.365-5.721 2.5-7.856c2.136-2.135 5.083-2.963 7.856-2.5c-.092-2.433-1.053-4.839-2.906-6.692s-4.26-2.814-6.692-2.906c.463 2.773-.365 5.721-2.5 7.856c-2.136 2.135-5.083 2.963-7.856 2.5a9.944 9.944 0 0 0 2.906 6.692"/><path fill="currentColor" d="M15.535 15.535a6.996 6.996 0 0 0-1.911 6.318a9.929 9.929 0 0 0 8.229-8.229a6.999 6.999 0 0 0-6.318 1.911m-7.07-7.07a6.996 6.996 0 0 0 1.911-6.318a9.929 9.929 0 0 0-8.23 8.229a7 7 0 0 0 6.319-1.911"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Terminal(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 4H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2M6.414 15.707L5 14.293L7.293 12L5 9.707l1.414-1.414L10.121 12zM19 16h-7v-2h7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Thermometer(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 16a3.001 3.001 0 0 0 6 0c0-.353-.072-.686-.184-1H9.184A2.962 2.962 0 0 0 9 16"/><path fill="currentColor" d="M18 6V4h-3.185A2.995 2.995 0 0 0 12 2c-1.654 0-3 1.346-3 3v5.8A6.027 6.027 0 0 0 6 16c0 3.309 2.691 6 6 6s6-2.691 6-6a6.027 6.027 0 0 0-3-5.2V10h3V8h-3V6zm-4.405 6.324A4.033 4.033 0 0 1 16 16c0 2.206-1.794 4-4 4s-4-1.794-4-4c0-1.585.944-3.027 2.405-3.676l.595-.263V5a1 1 0 0 1 2 0v7.061z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Time(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.25 2c-5.514 0-10 4.486-10 10s4.486 10 10 10s10-4.486 10-10s-4.486-10-10-10M18 13h-6.75V6h2v5H18z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimeFive(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m3.293 14.707L11 12.414V6h2v5.586l3.707 3.707z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Timer(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 3h4v2h-4zM3 8h4v2H3zm0 8h4v2H3zm-1-4h3.99v2H2zm19.707-5.293l-1.414-1.414L18.586 7A6.937 6.937 0 0 0 15 6c-3.859 0-7 3.141-7 7s3.141 7 7 7s7-3.141 7-7a6.968 6.968 0 0 0-1.855-4.73zM16 14h-2V8.958h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tired(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m-6.447 9.105l2.459-1.229l-1.567-1.044l1.109-1.664l3 2a1 1 0 0 1-.108 1.727l-4 2zM8 17s1-3 4-3s4 3 4 3zm9.553-4.105l-4-2a1 1 0 0 1-.108-1.727l3-2l1.109 1.664l-1.566 1.044l2.459 1.229z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToTop(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 15h4v6h6v-6h4l-7-8zM4 3h16v2H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToggleLeft(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 6H8c-3.296 0-5.982 2.682-6 5.986v.042A6.01 6.01 0 0 0 8 18h8a6.01 6.01 0 0 0 6-5.994v-.018C21.985 8.685 19.297 6 16 6m-8 9c-1.627 0-3-1.373-3-3s1.373-3 3-3s3 1.373 3 3s-1.373 3-3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToggleRight(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 6H8c-3.296 0-5.982 2.682-6 5.986v.042A6.01 6.01 0 0 0 8 18h8c3.309 0 6-2.691 6-6s-2.691-6-6-6m0 9c-1.627 0-3-1.373-3-3s1.373-3 3-3s3 1.373 3 3s-1.373 3-3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tone(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m-1 9H4.069a7.965 7.965 0 0 1 .52-2H11zm0 4H4.589a7.965 7.965 0 0 1-.52-2H11zm0-10.931V7H5.765A7.996 7.996 0 0 1 11 4.069M5.765 17H11v2.931A7.996 7.996 0 0 1 5.765 17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Torch(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 11.648V20a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2v-8.352c1.067-.552 3-1.928 3-4.648V5H5v2c0 2.72 1.933 4.096 3 4.648M11 11h2v3h-2zM5 2h14v2H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Traffic(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m2.958 16l.043 1.042c.005.12.142 2.255 2.999 3.338v1.12a.5.5 0 0 0 .5.5h11a.5.5 0 0 0 .5-.5v-1.12c2.857-1.083 2.994-3.218 2.999-3.338L21.043 16H18v-1.62c2.857-1.083 2.994-3.218 2.999-3.338L21.043 10H18V8.38c2.857-1.083 2.994-3.218 2.999-3.338L21.043 4H18V2.5a.5.5 0 0 0-.5-.5h-11a.5.5 0 0 0-.5.5V4H2.958l.043 1.042c.005.12.142 2.255 2.999 3.338V10H2.958l.043 1.042c.005.12.142 2.255 2.999 3.338V16zM12 4a2 2 0 1 1-.001 4.001A2 2 0 0 1 12 4m0 6a2 2 0 1 1-.001 4.001A2 2 0 0 1 12 10m0 6a2 2 0 1 1-.001 4.001A2 2 0 0 1 12 16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrafficBarrier(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 6h-2V3h-2v3H7V3H5v3H3a1 1 0 0 0-1 1v7a1 1 0 0 0 1 1h2v6h2v-6h10v6h2v-6h2a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1M4.42 13l2.857-5H9.58l-2.857 5zm7.857-5h2.303l-2.857 5H9.42zm5 0h2.303l-2.857 5H14.42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrafficCone(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.649 16H5.352l-1.06 3H2v2h20v-2h-2.292zM6.057 14h11.886l-1.412-4H7.469zM13 2h-2a1 1 0 0 0-.943.667L8.175 8h7.65l-1.882-5.333A1 1 0 0 0 13 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Train(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.375 2H7.621c-.224 0-1.399.065-2.503 1.351C4.031 4.616 4 5.862 4 6v11a2 2 0 0 0 2 2h1l-2 3h2.353l.667-1h8l.677 1H19l-2-3h1a2 2 0 0 0 2-2V6c.001-.188-.032-1.434-1.129-2.665C17.715 2.037 16.509 2 16.375 2M10 4h4v2h-4zM7.5 17a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 7.5 17m9 0a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 16.5 17m1.5-5H6V8h12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trash(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 7H5v13a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V7zm4 12H8v-9h2zm6 0h-2v-9h2zm.618-15L15 2H9L7.382 4H3v2h18V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrashAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 7H5v13a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V7zm10.618-3L15 2H9L7.382 4H3v2h18V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tree(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20 18l-4-5h3l-4-5h2l-5-6l-5 6h2l-4 5h3l-4 5h7v4h2v-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TreeAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m18 10l-6-8l-6 8h3l-5 8h7v4h2v-4h7l-5-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trophy(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 4h-3V3a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1v1H3a1 1 0 0 0-1 1v3c0 4.31 1.8 6.91 4.82 7A6 6 0 0 0 11 17.91V20H9v2h6v-2h-2v-2.09A6 6 0 0 0 17.18 15c3-.1 4.82-2.7 4.82-7V5a1 1 0 0 0-1-1M4 8V6h2v6.83C4.22 12.08 4 9.3 4 8m14 4.83V6h2v2c0 1.3-.22 4.08-2 4.83"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Truck(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.15 8a2 2 0 0 0-1.72-1H15V5a1 1 0 0 0-1-1H4a2 2 0 0 0-2 2v10a2 2 0 0 0 1 1.73a3.49 3.49 0 0 0 7 .27h3.1a3.48 3.48 0 0 0 6.9 0a2 2 0 0 0 2-2v-3a1.07 1.07 0 0 0-.14-.52zM15 9h2.43l1.8 3H15zM6.5 19A1.5 1.5 0 1 1 8 17.5A1.5 1.5 0 0 1 6.5 19m10 0a1.5 1.5 0 1 1 1.5-1.5a1.5 1.5 0 0 1-1.5 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tv(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 6h-5.586l2.293-2.293l-1.414-1.414L12 5.586L8.707 2.293L7.293 3.707L9.586 6H4c-1.103 0-2 .897-2 2v11c0 1.103.897 2 2 2h16c1.103 0 2-.897 2-2V8c0-1.103-.897-2-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UniversalAccess(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 3.33A1.67 1.67 0 1 1 10.33 7A1.67 1.67 0 0 1 12 5.33m3.33 12.5l-1.66.84l-1.39-3.89h-.56l-1.39 3.89l-1.66-.84l1.66-4.72v-1.66L7 10.33l.56-1.66l3.33 1.11h2.22l3.33-1.11l.56 1.66l-3.33 1.12v1.66z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpArrow(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 19h18a1.002 1.002 0 0 0 .823-1.569l-9-13c-.373-.539-1.271-.539-1.645 0l-9 13A.999.999 0 0 0 3 19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpArrowAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 18v-6h5l-6-7l-6 7h5v6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpArrowCircle(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22c5.514 0 10-4.486 10-10S17.514 2 12 2S2 6.486 2 12s4.486 10 10 10m0-15l5 5h-4v5h-2v-5H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpArrowSquare(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 21h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2m7-14l5 5h-4v5h-2v-5H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpsideDown(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2M8.507 15a1.494 1.494 0 1 1 .001-2.987A1.494 1.494 0 0 1 8.507 15m4.301-6.919a4.108 4.108 0 0 0-1.616 0a4.12 4.12 0 0 0-.751.233c-.234.1-.463.224-.678.368a4.077 4.077 0 0 0-1.08 1.082L7.024 8.646a6.026 6.026 0 0 1 2.639-2.175a6.097 6.097 0 0 1 1.128-.35a6.061 6.061 0 0 1 2.415 0a5.919 5.919 0 0 1 2.148.903a6.078 6.078 0 0 1 1.621 1.622l-1.658 1.117a3.994 3.994 0 0 0-.488-.59a3.988 3.988 0 0 0-2.021-1.092M15.5 15a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 15.5 15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Upvote(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 14h4v7a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1v-7h4a1.001 1.001 0 0 0 .781-1.625l-8-10c-.381-.475-1.181-.475-1.562 0l-8 10A1.001 1.001 0 0 0 4 14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func User(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 6.5C7.5 8.981 9.519 11 12 11s4.5-2.019 4.5-4.5S14.481 2 12 2S7.5 4.019 7.5 6.5M20 21h1v-1c0-3.859-3.141-7-7-7h-4c-3.86 0-7 3.141-7 7v1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserAccount(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 2H8a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2m-6 2.5a2.5 2.5 0 1 1 0 5a2.5 2.5 0 0 1 0-5M19 15H9v-.25C9 12.901 11.254 11 14 11s5 1.901 5 3.75z"/><path fill="currentColor" d="M4 8H2v12c0 1.103.897 2 2 2h12v-2H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserBadge(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.988 22a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2zM9 5h6v2H9zm5.25 6.25A2.26 2.26 0 0 1 12 13.501c-1.235 0-2.25-1.015-2.25-2.251S10.765 9 12 9a2.259 2.259 0 0 1 2.25 2.25M7.5 18.188c0-1.664 2.028-3.375 4.5-3.375s4.5 1.711 4.5 3.375v.563h-9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserCheck(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 12.052c1.995 0 3.5-1.505 3.5-3.5s-1.505-3.5-3.5-3.5s-3.5 1.505-3.5 3.5s1.505 3.5 3.5 3.5M9 13H7c-2.757 0-5 2.243-5 5v1h12v-1c0-2.757-2.243-5-5-5m11.294-4.708l-4.3 4.292l-1.292-1.292l-1.414 1.414l2.706 2.704l5.712-5.702z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserCircle(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.579 2 2 6.579 2 12s4.579 10 10 10s10-4.579 10-10S17.421 2 12 2m0 5c1.727 0 3 1.272 3 3s-1.273 3-3 3c-1.726 0-3-1.272-3-3s1.274-3 3-3m-5.106 9.772c.897-1.32 2.393-2.2 4.106-2.2h2c1.714 0 3.209.88 4.106 2.2C15.828 18.14 14.015 19 12 19s-3.828-.86-5.106-2.228"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserDetail(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 11h7v2h-7zm1 4h6v2h-6zm-2-8h8v2h-8zM4 19h10v-1c0-2.757-2.243-5-5-5H7c-2.757 0-5 2.243-5 5v1zm4-7c1.995 0 3.5-1.505 3.5-3.5S9.995 5 8 5S4.5 6.505 4.5 8.5S6.005 12 8 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserMinus(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 11h8v2h-8zM4.5 8.552c0 1.995 1.505 3.5 3.5 3.5s3.5-1.505 3.5-3.5s-1.505-3.5-3.5-3.5s-3.5 1.505-3.5 3.5M4 19h10v-1c0-2.757-2.243-5-5-5H7c-2.757 0-5 2.243-5 5v1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserPin(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 2H5a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2h4l3 3l3-3h4a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2m-7 3c1.727 0 3 1.272 3 3s-1.273 3-3 3c-1.726 0-3-1.272-3-3s1.274-3 3-3M7.177 16c.558-1.723 2.496-3 4.823-3s4.266 1.277 4.823 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserPlus(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 8.552c0 1.995 1.505 3.5 3.5 3.5s3.5-1.505 3.5-3.5s-1.505-3.5-3.5-3.5s-3.5 1.505-3.5 3.5M19 8h-2v3h-3v2h3v3h2v-3h3v-2h-3zM4 19h10v-1c0-2.757-2.243-5-5-5H7c-2.757 0-5 2.243-5 5v1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserRectangle(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 22h13a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2zm6-17.001c1.647 0 3 1.351 3 3C15 9.647 13.647 11 12 11S9 9.647 9 7.999c0-1.649 1.353-3 3-3M6 17.25c0-2.219 2.705-4.5 6-4.5s6 2.281 6 4.5V18H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserVoice(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 12.052c1.995 0 3.5-1.505 3.5-3.5s-1.505-3.5-3.5-3.5s-3.5 1.505-3.5 3.5s1.505 3.5 3.5 3.5M9 13H7c-2.757 0-5 2.243-5 5v1h12v-1c0-2.757-2.243-5-5-5m9.364-10.364L16.95 4.05C18.271 5.373 19 7.131 19 9s-.729 3.627-2.05 4.95l1.414 1.414C20.064 13.663 21 11.403 21 9s-.936-4.663-2.636-6.364"/><path fill="currentColor" d="M15.535 5.464L14.121 6.88C14.688 7.445 15 8.198 15 9s-.312 1.555-.879 2.12l1.414 1.416C16.479 11.592 17 10.337 17 9s-.521-2.592-1.465-3.536"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserX(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 12.052c1.995 0 3.5-1.505 3.5-3.5s-1.505-3.5-3.5-3.5s-3.5 1.505-3.5 3.5s1.505 3.5 3.5 3.5M9 13H7c-2.757 0-5 2.243-5 5v1h12v-1c0-2.757-2.243-5-5-5m11.293-4.707L18 10.586l-2.293-2.293l-1.414 1.414l2.292 2.292l-2.293 2.293l1.414 1.414l2.293-2.293l2.294 2.294l1.414-1.414L19.414 12l2.293-2.293z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vector(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.277 8c.347.596.985 1 1.723 1a2 2 0 0 0 0-4c-.738 0-1.376.404-1.723 1H16V4a1 1 0 0 0-1-1H9a1 1 0 0 0-1 1v2H5.723C5.376 5.404 4.738 5 4 5a2 2 0 0 0 0 4c.738 0 1.376-.404 1.723-1H8v.368C5.134 9.839 4.319 12.534 4.092 14H3a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1h-.877c.197-.959.718-2.406 2.085-3.418A.984.984 0 0 0 9 11h6a.98.98 0 0 0 .792-.419c1.373 1.013 1.895 2.458 2.089 3.419H17a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1h-1.092c-.227-1.466-1.042-4.161-3.908-5.632V8zM14 9h-4V5h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vial(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 22a4.965 4.965 0 0 0 3.535-1.465l9.193-9.193l.707.708l1.414-1.414l-8.485-8.486l-1.414 1.414l.708.707l-9.193 9.193C2.521 14.408 2 15.664 2 17s.521 2.592 1.465 3.535A4.965 4.965 0 0 0 7 22M18.314 9.928L15.242 13H6.758l7.314-7.314z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Video(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 7c0-1.103-.897-2-2-2H4c-1.103 0-2 .897-2 2v10c0 1.103.897 2 2 2h12c1.103 0 2-.897 2-2v-3.333L22 17V7l-4 3.333z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoOff(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 19h10.879L2.145 6.265A1.977 1.977 0 0 0 2 7v10c0 1.103.897 2 2 2M18 7c0-1.103-.897-2-2-2H6.414L3.707 2.293L2.293 3.707l18 18l1.414-1.414L18 16.586v-2.919L22 17V7l-4 3.333z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoPlus(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 7a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-3.333L22 17V7l-4 3.333zm-4 6h-3v3H9v-3H6v-2h3V8h2v3h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoRecording(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 9c0-1.103-.897-2-2-2h-1.434l-2.418-4.029A2.008 2.008 0 0 0 10.434 2H5v2h5.434l1.8 3H4c-1.103 0-2 .897-2 2v9c0 1.103.897 2 2 2h12c1.103 0 2-.897 2-2v-3l4 2v-7l-4 2zm-7 8H5v-2h6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Videos(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 8H2v12a2 2 0 0 0 2 2h12v-2H4z"/><path fill="currentColor" d="M20 2H8a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2m-9 12V6l7 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Virus(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 11h-.17c-1.053 0-1.958-.669-2.357-1.644l-.021-.049c-.408-.977-.249-2.097.5-2.846l.119-.119a.999.999 0 1 0-1.414-1.414l-.119.119c-.749.749-1.869.908-2.846.5l-.049-.021C13.669 5.128 13 4.218 13 3.165v-.081C13 2.447 12.553 2 12 2s-1 .447-1 1v.036c0 1.096-.66 2.084-1.673 2.503l-.006.003a2.71 2.71 0 0 1-2.953-.588l-.025-.025a.999.999 0 1 0-1.414 1.414l.036.036a2.69 2.69 0 0 1 .583 2.929l-.027.064A2.638 2.638 0 0 1 3.085 11h-.001C2.447 11 2 11.447 2 12s.447 1 1 1h.068a2.66 2.66 0 0 1 2.459 1.644l.021.049a2.69 2.69 0 0 1-.583 2.929l-.036.036a.999.999 0 1 0 1.414 1.414l.036-.036a2.689 2.689 0 0 1 2.929-.583l.143.06A2.505 2.505 0 0 1 11 20.83v.085c0 .638.447 1.085 1 1.085s1-.448 1-1v-.17c0-1.015.611-1.93 1.55-2.318l.252-.104a2.508 2.508 0 0 1 2.736.545l.119.119a.999.999 0 1 0 1.414-1.414l-.119-.119c-.749-.749-.908-1.869-.5-2.846l.021-.049c.399-.975 1.309-1.644 2.362-1.644h.08c.638 0 1.085-.447 1.085-1s-.447-1-1-1M8 13a1 1 0 1 1 0-2a1 1 0 0 1 0 2m5 3.5a1 1 0 1 1 0-2a1 1 0 0 1 0 2m1-4.5a2 2 0 1 1 .001-4.001A2 2 0 0 1 14 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VirusBlock(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.952 17.538c-.749-.749-.908-1.869-.5-2.846l.021-.049c.399-.974 1.309-1.643 2.362-1.643h.08c.638 0 1.085-.447 1.085-1s-.447-1-1-1h-.17c-1.053 0-1.958-.669-2.357-1.644l-.021-.049c-.408-.977-.249-2.097.5-2.846l.119-.119a.999.999 0 1 0-1.414-1.414l-.119.119c-.749.749-1.869.908-2.846.5l-.049-.021C13.669 5.128 13 4.218 13 3.165v-.081C13 2.447 12.553 2 12 2s-1 .447-1 1v.036c0 1.096-.66 2.084-1.673 2.503l-.006.003a2.71 2.71 0 0 1-2.953-.588l-.025-.025l-2.636-2.636l-1.414 1.414l18 18l1.414-1.414l-2.636-2.636zM12 10a2 2 0 1 1 2 2c-.257 0-.501-.053-.728-.142l-1.131-1.131A1.998 1.998 0 0 1 12 10m-4 3a1 1 0 0 1-1-1a.99.99 0 0 1 .244-.635L5.431 9.552A2.634 2.634 0 0 1 3.085 11h-.001C2.447 11 2 11.447 2 12s.447 1 1 1h.068a2.66 2.66 0 0 1 2.459 1.644l.021.049a2.69 2.69 0 0 1-.583 2.929l-.036.036a.999.999 0 1 0 1.414 1.414l.036-.036a2.689 2.689 0 0 1 2.929-.583l.143.06A2.505 2.505 0 0 1 11 20.83v.085c0 .638.447 1.085 1 1.085s1-.448 1-1v-.17c0-.976.568-1.853 1.443-2.266l-5.809-5.809A.98.98 0 0 1 8 13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Volume(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 17h2.697L14 21.868V2.132L6.697 7H4c-1.103 0-2 .897-2 2v6c0 1.103.897 2 2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeFull(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 21c3.527-1.547 5.999-4.909 5.999-9S19.527 4.547 16 3v2c2.387 1.386 3.999 4.047 3.999 7S18.387 17.614 16 19z"/><path fill="currentColor" d="M16 7v10c1.225-1.1 2-3.229 2-5s-.775-3.9-2-5M4 17h2.697L14 21.868V2.132L6.697 7H4c-1.103 0-2 .897-2 2v6c0 1.103.897 2 2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeLow(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 17h2.697L14 21.868V2.132L6.697 7H4c-1.103 0-2 .897-2 2v6c0 1.103.897 2 2 2M16 7v10c1.225-1.1 2-3.229 2-5s-.775-3.9-2-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeMute(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7.727 6.313l-4.02-4.02l-1.414 1.414l18 18l1.414-1.414l-2.02-2.02A9.578 9.578 0 0 0 21.999 12c0-4.091-2.472-7.453-5.999-9v2c2.387 1.386 3.999 4.047 3.999 7a8.13 8.13 0 0 1-1.671 4.914l-1.286-1.286C17.644 14.536 18 13.19 18 12c0-1.771-.775-3.9-2-5v7.586l-2-2V2.132zM4 17h2.697L14 21.868v-3.747L3.102 7.223A1.995 1.995 0 0 0 2 9v6c0 1.103.897 2 2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wallet(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 7V5c0-1.103-.897-2-2-2H5C3.346 3 2 4.346 2 6v12c0 2.201 1.794 3 3 3h15c1.103 0 2-.897 2-2V9c0-1.103-.897-2-2-2m-2 9h-2v-4h2zM5 7a1.001 1.001 0 0 1 0-2h13v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WalletAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 9h8v6h-8z"/><path fill="currentColor" d="M20 3H5C3.346 3 2 4.346 2 6v12c0 1.654 1.346 3 3 3h15c1.103 0 2-.897 2-2v-2h-8c-1.103 0-2-.897-2-2V9c0-1.103.897-2 2-2h8V5c0-1.103-.897-2-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Washer(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 22h16a1 1 0 0 0 1-1V5c0-1.654-1.346-3-3-3H6C4.346 2 3 3.346 3 5v16a1 1 0 0 0 1 1M18 3.924a1 1 0 1 1 0 2a1 1 0 0 1 0-2m-3 0a1 1 0 1 1 0 2a1 1 0 0 1 0-2M12 7c3.309 0 6 2.691 6 6s-2.691 6-6 6s-6-2.691-6-6s2.691-6 6-6"/><path fill="currentColor" d="M12.766 16.929c1.399-.261 2.571-1.315 3.023-2.665a3.853 3.853 0 0 0-.153-2.893a.482.482 0 0 0-.544-.266c-.604.149-1.019.448-1.5.801c-.786.577-1.765 1.294-3.592 1.294c-.813 0-1.45-.146-1.984-.354l-.013.009a4.006 4.006 0 0 0 4.763 4.074"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Watch(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 8h-2v5h5v-2h-3z"/><path fill="currentColor" d="M19.999 12c0-2.953-1.612-5.53-3.999-6.916V3a1 1 0 0 0-1-1H9a1 1 0 0 0-1 1v2.083C5.613 6.469 4.001 9.047 4.001 12a8.003 8.003 0 0 0 4.136 7H8v2.041a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V19h-.139a8 8 0 0 0 4.138-7m-8 5.999A6.005 6.005 0 0 1 6.001 12a6.005 6.005 0 0 1 5.998-5.999c3.31 0 6 2.691 6 5.999a6.005 6.005 0 0 1-6 5.999"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WatchAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 8c0-.909-.613-1.67-1.445-1.912l-1.31-3.443A1 1 0 0 0 14.311 2H8.689a1 1 0 0 0-.934.645l-1.31 3.443A1.996 1.996 0 0 0 5 8v8c0 .909.613 1.67 1.445 1.912l1.31 3.443a1 1 0 0 0 .934.645h5.621c.415 0 .787-.257.935-.645l1.31-3.443A1.996 1.996 0 0 0 18 16v-2h1v-4h-1zm-1.998 8H7V8h9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Webcam(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2c-4.963 0-9 4.037-9 9c0 3.328 1.82 6.232 4.513 7.79l-2.067 1.378A1 1 0 0 0 6 22h12a1 1 0 0 0 .555-1.832l-2.067-1.378C19.18 17.232 21 14.328 21 11c0-4.963-4.037-9-9-9m0 16c-3.859 0-7-3.141-7-7s3.141-7 7-7s7 3.141 7 7s-3.141 7-7 7"/><path fill="currentColor" d="M12 6c-2.757 0-5 2.243-5 5s2.243 5 5 5s5-2.243 5-5s-2.243-5-5-5m-1.5 5a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 10.5 11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Widget(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 11h6a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1m0 10h6a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1m10 0h6a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1h-6a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1m7.293-14.707l-3.586-3.586a.999.999 0 0 0-1.414 0l-3.586 3.586a.999.999 0 0 0 0 1.414l3.586 3.586a.999.999 0 0 0 1.414 0l3.586-3.586a.999.999 0 0 0 0-1.414"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowAlt(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 3H4c-1.103 0-2 .897-2 2v14c0 1.103.897 2 2 2h16c1.103 0 2-.897 2-2V5c0-1.103-.897-2-2-2m-3 3h2v2h-2zm-3 0h2v2h-2zM4 19v-9h16.001l.001 9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wine(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 17.916V20H9v2h6v-2h-2v-2.084c3.162-.402 5.849-2.66 6.713-5.793c.264-.952.312-2.03.143-3.206l-.866-6.059A1 1 0 0 0 18 2H6a1 1 0 0 0-.99.858l-.865 6.058c-.169 1.177-.121 2.255.143 3.206c.863 3.134 3.55 5.392 6.712 5.794M17.133 4l.57 4H6.296l.571-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WinkSmile(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2M8.5 9a1.5 1.5 0 1 1-.001 3.001A1.5 1.5 0 0 1 8.5 9m7.741 7.244a5.982 5.982 0 0 1-3.034 1.634a6.042 6.042 0 0 1-3.541-.349a5.997 5.997 0 0 1-2.642-2.176l1.658-1.117c.143.211.307.41.488.59a3.988 3.988 0 0 0 1.273.86c.243.102.495.181.749.232a4.108 4.108 0 0 0 1.616 0c.253-.052.505-.131.75-.233c.234-.1.464-.224.679-.368c.208-.142.407-.306.591-.489c.183-.182.347-.381.489-.592l1.658 1.117c-.214.32-.461.62-.734.891M13 12s.5-2 2.5-2c1.999 0 2.5 2 2.5 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WinkTongue(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 19v-4h-4v4c0 1.103.897 2 2 2s2-.897 2-2"/><path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12c0 4.434 2.903 8.198 6.906 9.505A3.969 3.969 0 0 1 8 19v-2.499C6.412 15.027 6 13 6 13h12s-.411 2.027-2 3.501V19c0 .953-.349 1.816-.906 2.504C19.097 20.197 22 16.434 22 12c0-5.514-4.486-10-10-10m-3.5 9a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 8.5 11m4.5 0s.5-2 2.5-2c1.999 0 2.5 2 2.5 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wrench(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.512 6.112l-3.89 3.889l-3.535-3.536l3.889-3.889a6.501 6.501 0 0 0-8.484 8.486l-6.276 6.275a.999.999 0 0 0 0 1.414l2.122 2.122a.999.999 0 0 0 1.414 0l6.275-6.276a6.501 6.501 0 0 0 7.071-1.414a6.504 6.504 0 0 0 1.414-7.071"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Xcircle(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2m4.207 12.793l-1.414 1.414L12 13.414l-2.793 2.793l-1.414-1.414L10.586 12L7.793 9.207l1.414-1.414L12 10.586l2.793-2.793l1.414 1.414L13.414 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Xsquare(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2zm-4.793 9.793l-1.414 1.414L12 13.414l-2.793 2.793l-1.414-1.414L10.586 12L7.793 9.207l1.414-1.414L12 10.586l2.793-2.793l1.414 1.414L13.414 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YinYang(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.062 4.938A9.942 9.942 0 0 0 12.016 2h-.026a9.94 9.94 0 0 0-7.071 2.938c-3.898 3.898-3.898 10.243 0 14.143c1.895 1.895 4.405 2.938 7.071 2.938s5.177-1.043 7.071-2.938c3.9-3.899 3.9-10.243.001-14.143M13.5 15a1.5 1.5 0 1 1-.001 3.001A1.5 1.5 0 0 1 13.5 15M6.333 6.353A7.953 7.953 0 0 1 11.99 4l.026.001c1.652.008 3.242 1.066 3.55 2.371c.366 1.552-1.098 3.278-4.018 4.737c-5.113 2.555-5.312 5.333-4.975 6.762l.008.021c-.082-.075-.169-.146-.249-.226c-3.118-3.119-3.118-8.194.001-11.313"/><circle cx="10.5" cy="7.5" r="1.5" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Zap(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.98 11.802a.995.995 0 0 0-.738-.771l-6.86-1.716l2.537-5.921a.998.998 0 0 0-.317-1.192a.996.996 0 0 0-1.234.024l-11 9a1 1 0 0 0 .39 1.744l6.719 1.681l-3.345 5.854A1.001 1.001 0 0 0 8 22a.995.995 0 0 0 .6-.2l12-9a1 1 0 0 0 .38-.998"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomIn(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 2c-4.411 0-8 3.589-8 8s3.589 8 8 8a7.952 7.952 0 0 0 4.897-1.688l4.396 4.396l1.414-1.414l-4.396-4.396A7.952 7.952 0 0 0 18 10c0-4.411-3.589-8-8-8m4 9h-3v3H9v-3H6V9h3V6h2v3h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomOut(children ...ElementRenderer) *BxsIcon {
	return &BxsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 18a7.952 7.952 0 0 0 4.897-1.688l4.396 4.396l1.414-1.414l-4.396-4.396A7.952 7.952 0 0 0 18 10c0-4.411-3.589-8-8-8s-8 3.589-8 8s3.589 8 8 8M6 9h8v2H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
