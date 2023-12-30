package eos_icons

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = "5.4.0"
	hAttr          = "1em"
	viewbox        = "0 0 24 24"
	fill           = "currentColor"
)

type EosIconsIcon struct {
	*SVGSVGElement
}

func Abstract(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3H8a2 2 0 0 0-1.71 1l-4 7a2 2 0 0 0 0 2l4 7A2 2 0 0 0 8 21h8a2 2 0 0 0 1.74-1l4-7a2 2 0 0 0 0-2l-4-7A2 2 0 0 0 16 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AbstractIncomplete(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.72 11.01l-4.01-7A1.968 1.968 0 0 0 15.98 3H8a1.968 1.968 0 0 0-1.73 1.01L4.74 6.68l-.8 1.39l-1.68 2.94A2.033 2.033 0 0 0 2 12a2.004 2.004 0 0 0 .26.99l1.68 2.94l.8 1.39l1.53 2.67A1.968 1.968 0 0 0 8 21h7.98a1.968 1.968 0 0 0 1.73-1.01l4.01-7a2.004 2.004 0 0 0 .26-.99a1.956 1.956 0 0 0-.26-.99m-4.384 5.974l.006-.004l-.004.007ZM20.28 12l-2.935 4.974c-.118.01-.234.026-.355.026A4.994 4.994 0 0 1 12 12.098v-.118A5.004 5.004 0 0 0 7 7a2.96 2.96 0 0 0-.31.03L7.89 5h8.26l4.14 7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AbstractInstance(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.71 11l-4-7A2 2 0 0 0 16 3H8a2 2 0 0 0-1.7 1l-4 7a2 2 0 0 0 0 2l4 7A2 2 0 0 0 8 21h8a2 2 0 0 0 1.74-1l4-7a2 2 0 0 0-.03-2M12 14a2 2 0 1 1 2-2a2 2 0 0 1-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AbstractInstanceOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.732 11l-4-7a2 2 0 0 0-1.74-1h-8a2 2 0 0 0-1.71 1l-4 7a2 2 0 0 0 0 2l4 7a2 2 0 0 0 1.71 1h8a2 2 0 0 0 1.74-1l4-7a2 2 0 0 0 0-2M16 19H8l-4-7l4-7h8l4 7Z"/><circle cx="12" cy="12" r="2" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AbstractOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.732 11l-4.003-7a1.967 1.967 0 0 0-1.72-1H8.005a2.038 2.038 0 0 0-1.733 1l-4.003 7a1.999 1.999 0 0 0 0 2l4.003 7a2.038 2.038 0 0 0 1.733 1h8.006a1.966 1.966 0 0 0 1.719-1l4.003-7a1.999 1.999 0 0 0 0-2M16 19H8l-4-7l4-7h8l4 7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ActionChains(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="20.01" cy="20.01" r="3" fill="currentColor"/><path fill="currentColor" d="M18 3H6.83a3 3 0 1 0 0 2H18a3 3 0 0 1 0 6h-3.17a3 3 0 0 0-5.64 0H6a5 5 0 0 0 0 10h6v3l4-4l-4-4v3H6a3 3 0 1 1 0-6h3.2a3 3 0 0 0 5.63 0H18a5 5 0 0 0 0-10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ActionChainsOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 3H6.834a3 3 0 1 0 0 2H18a3 3 0 0 1 0 6h-3.168a3 3 0 0 0-5.639 0H6a5 5 0 0 0 0 10h6v3l4-4l-4-4v3H6a3 3 0 0 1 0-6h3.203a3 3 0 0 0 5.629 0H18a5 5 0 0 0 0-10M4 5a1 1 0 1 1 1-1a1 1 0 0 1-1 1m8 8a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/><path fill="currentColor" d="M20 17a3 3 0 1 0 3 3a3 3 0 0 0-3-3m0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ActivateSubscriptions(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 6h16v2H4zm2-4h12v2H6zm14 8H4a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-8a2 2 0 0 0-2-2m-5 7h-2v2h-2v-2H9v-2h2v-2h2v2h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ActivateSubscriptionsOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.004 5.999h16v2h-16zm2-3.999h12v2h-12zm13.993 7.991H5.003A2.008 2.008 0 0 0 3 12.004v7.974a2.008 2.008 0 0 0 2.003 2.013h15.994A2.008 2.008 0 0 0 23 19.978v-7.974a2.008 2.008 0 0 0-2.003-2.013M21 20H5v-8h16Z"/><path fill="currentColor" d="M12.004 18.991h2v-1.989h2v-2.004h-2v-2.007h-2v2.01h-2v2.001h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Admin(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1L3 5v6c0 5.55 3.84 10.74 9 12c5.16-1.26 9-6.45 9-12V5Zm0 3.9a3 3 0 1 1-3 3a3 3 0 0 1 3-3m0 7.9c2 0 6 1.09 6 3.08a7.2 7.2 0 0 1-12 0c0-1.99 4-3.08 6-3.08"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdminOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 23C6.443 21.765 2 16.522 2 11V5l10-4l10 4v6c0 5.524-4.443 10.765-10 12M4 6v5a10.58 10.58 0 0 0 8 10a10.58 10.58 0 0 0 8-10V6l-8-3Z"/><circle cx="12" cy="8.5" r="2.5" fill="currentColor"/><path fill="currentColor" d="M7 15a5.782 5.782 0 0 0 5 3a5.782 5.782 0 0 0 5-3c-.025-1.896-3.342-3-5-3c-1.667 0-4.975 1.104-5 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ai(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 11V9h-2V7a2.006 2.006 0 0 0-2-2h-2V3h-2v2h-2V3H9v2H7a2.006 2.006 0 0 0-2 2v2H3v2h2v2H3v2h2v2a2.006 2.006 0 0 0 2 2h2v2h2v-2h2v2h2v-2h2a2.006 2.006 0 0 0 2-2v-2h2v-2h-2v-2Zm-4 6H7V7h10Z"/><path fill="currentColor" d="M11.361 8h-1.345l-2.01 8h1.027l.464-1.875h2.316L12.265 16h1.062Zm-1.729 5.324L10.65 8.95h.046l.983 4.374ZM14.244 8h1v8h-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AiHealing(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 4h2v1h-2z"/><circle cx="9" cy="4" r="1" fill="currentColor"/><circle cx="15" cy="4" r="1" fill="currentColor"/><path fill="currentColor" d="M18 10h-5V8l3-.01A3 3 0 0 0 19 5V3a3.009 3.009 0 0 0-3-3H8a3.009 3.009 0 0 0-3 3v2a3.009 3.009 0 0 0 3 3h3v2H6a3.009 3.009 0 0 0-3 3v8a3.009 3.009 0 0 0 3 3h12a3.009 3.009 0 0 0 3-3v-8a3.009 3.009 0 0 0-3-3M8 6a1.003 1.003 0 0 1-1-1V3a1.003 1.003 0 0 1 1-1h7.99A1.012 1.012 0 0 1 17 3v2a1.012 1.012 0 0 1-1.01 1Zm4.435 14.072L12 20.5l-.435-.432C10.02 18.541 9 17.534 9 16.298a1.712 1.712 0 0 1 1.65-1.798a1.74 1.74 0 0 1 1.35.683a1.74 1.74 0 0 1 1.35-.683A1.712 1.712 0 0 1 15 16.298c0 1.236-1.02 2.243-2.565 3.774"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AiHealingOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12 20.5l-.435-.432C10.02 18.541 9 17.534 9 16.298a1.712 1.712 0 0 1 1.65-1.798a1.74 1.74 0 0 1 1.35.683a1.74 1.74 0 0 1 1.35-.683A1.712 1.712 0 0 1 15 16.298c0 1.236-1.02 2.243-2.565 3.774ZM11 4h2v1h-2z"/><circle cx="9" cy="4" r="1" fill="currentColor"/><circle cx="15" cy="4" r="1" fill="currentColor"/><path fill="currentColor" d="M18 10h-5V8l3-.008A2.998 2.998 0 0 0 19 5V3a3.004 3.004 0 0 0-3-3H8a3.004 3.004 0 0 0-3 3v2a3.004 3.004 0 0 0 3 3h3v2H6a3.005 3.005 0 0 0-3 3v8a3.004 3.004 0 0 0 3 3h12a3.004 3.004 0 0 0 3-3v-8a3.005 3.005 0 0 0-3-3M8 6a1.001 1.001 0 0 1-1-1V3a1.001 1.001 0 0 1 1-1l7.992.002A1.006 1.006 0 0 1 17 3v2a1.008 1.008 0 0 1-1.008 1Zm11 15a1.001 1.001 0 0 1-1 1H6a1.001 1.001 0 0 1-1-1v-8a1.002 1.002 0 0 1 1-1h12a1.001 1.001 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AiOperator(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.356 5H7.01L5 13h1.028l.464-1.875h2.316L9.26 13h1.062Zm-1.729 5.322L7.644 5.95h.045l.984 4.373ZM11.238 13V5h1v8Zm.187 1H4V4h10v4.78a5.504 5.504 0 0 1 4-.786V6h-2V4a2.006 2.006 0 0 0-2-2h-2V0h-2v2H8V0H6v2H4a2.006 2.006 0 0 0-2 2v2H0v2h2v2H0v2h2v2a2.006 2.006 0 0 0 2 2h2v2h2v-2h2v2h2v-1.992A5.547 5.547 0 0 1 11.425 14m2.075-.5A3.5 3.5 0 1 1 17 17a3.499 3.499 0 0 1-3.5-3.5M17 19c-2.336 0-7 1.173-7 3.5V24h14v-1.5c0-2.328-4.664-3.5-7-3.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Api(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.26 10.5h2v1h-2z"/><path fill="currentColor" d="M20 4H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2M8.4 15L8 13.77H6.06L5.62 15H4l2.2-6h1.62L10 15Zm8.36-3.5a1.47 1.47 0 0 1-1.5 1.5h-2v2h-1.5V9h3.5a1.47 1.47 0 0 1 1.5 1.5ZM20 15h-1.5V9H20Z"/><path fill="currentColor" d="M6.43 12.77h1.16l-.58-1.59z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ApiOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.312 9H5.688L3.5 15h1.607l.446-1.226h1.894L7.893 15H9.5Zm-1.394 3.774L6.5 11.18l.582 1.595ZM14.744 9h-3.5v6h1.5v-2h2a1.473 1.473 0 0 0 1.5-1.5v-1a1.473 1.473 0 0 0-1.5-1.5m0 2.5h-2v-1h2ZM18 9h1.5v6H18z"/><path fill="currentColor" d="M22 6v12H2V6zm0-2H2a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h20a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Application(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 3v2h22V3a2 2 0 0 0-2-2H3a2 2 0 0 0-2 2m3 1a1 1 0 1 1 1-1a1 1 0 0 1-1 1m3 0a1 1 0 1 1 1-1a1.003 1.003 0 0 1-1 1M1 6v14a2 2 0 0 0 2 2h18a2 2 0 0 0 2-2V6Zm8.1 9.3l1.3 1.3l-.7.7l-.7.7l-2-2l-2-2l2-2l2-2l.7.7l.7.7l-1.3 1.3L7.8 14Zm7.9.7l-2 2l-.7-.7l-.7-.7l1.3-1.3l1.3-1.3l-1.3-1.3l-1.3-1.3l.7-.7l.7-.7l2 2l2 2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ApplicationIncomplete(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 3v2h22V3a2 2 0 0 0-2-2H3a2 2 0 0 0-2 2m3 1a1 1 0 1 1 1-1a1 1 0 0 1-1 1m3 0a1 1 0 1 1 1-1a1.003 1.003 0 0 1-1 1M1 6v14a2 2 0 0 0 2 2h18a2 2 0 0 0 2-2V6Zm11 13a5 5 0 0 1 0-10v5h5a5 5 0 0 1-5 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ApplicationIncompleteOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 3v2h22l-.01-1.993A2 2 0 0 0 21 1H3a2.001 2.001 0 0 0-2 2m3 1a1 1 0 1 1 1-1a1 1 0 0 1-1 1m3 0a1 1 0 1 1 1-1a1.004 1.004 0 0 1-1 1m5 6a4 4 0 1 0 4 4h-4Z"/><path fill="currentColor" d="M1 6v14a2 2 0 0 0 2 2h18a2 2 0 0 0 2-2V6Zm20 14H3V8h18Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ApplicationInstance(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 6v14a2 2 0 0 0 2 2h18a2.006 2.006 0 0 0 2-2.007V6Zm8 13l.003-4.502L9 10l7 4.5ZM1.01 3.007L1 5h22l-.01-1.993A2 2 0 0 0 21 1H3a2 2 0 0 0-1.99 2.007m2.997 1a1 1 0 1 1 .999-.999a1 1 0 0 1-1 1m2.997 0a1 1 0 1 1 1-1a1.002 1.002 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ApplicationInstanceOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 3v2h22V3a2.001 2.001 0 0 0-2-2H3a2.001 2.001 0 0 0-2 2m3 1a.987.987 0 0 1-.993-.992A1.001 1.001 0 0 1 4 2a1.013 1.013 0 0 1 1.006 1.008A.998.998 0 0 1 4 4m3 0a1 1 0 1 1 1-1a1.004 1.004 0 0 1-1 1M1 6v14a2 2 0 0 0 2 2h18a2 2 0 0 0 2-2V6Zm20 14H3V8h18Z"/><path fill="currentColor" d="M9 10v8l7-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ApplicationOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<defs><path id="eosIconsApplicationOutlined0" fill="currentColor" d="m9 18l.7-.7l.7-.7l-1.3-1.3L7.8 14l1.3-1.3l1.3-1.3l-.7-.7L9 10l-2 2l-2 2l2 2zm4.6-1.4l.7.7l.7.7l2-2l2-2l-2-2l-2-2l-.7.7l-.7.7l1.3 1.3l1.3 1.3l-1.3 1.3z"/></defs><path fill="currentColor" d="M1 3v2h22V3a2.001 2.001 0 0 0-2-2H3a2.001 2.001 0 0 0-2 2m3.007 1.008a1 1 0 1 1 .999-1a1 1 0 0 1-1 1m2.997-.001a1 1 0 1 1 1-1a1.002 1.002 0 0 1-1 1M1 6v14a2 2 0 0 0 2 2h18a2 2 0 0 0 2-2V6Zm20 14H3V8h18Z"/><use href="#eosIconsApplicationOutlined0"/><use href="#eosIconsApplicationOutlined0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ApplicationWindow(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 1H3a2 2 0 0 0-2 2v2h22V3a2 2 0 0 0-2-2M4 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1m3 0a1 1 0 1 1 1-1a1.003 1.003 0 0 1-1 1m16 6V6H1v14a2 2 0 0 0 2 2h18a2 2 0 0 0 2-2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ApplicationWindowOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 3v2h22V3a2.001 2.001 0 0 0-2-2H3a2.001 2.001 0 0 0-2 2m3 1a.987.987 0 0 1-.993-.992A1.001 1.001 0 0 1 4 2a1.013 1.013 0 0 1 1.006 1.008A.998.998 0 0 1 4 4m3 0a1 1 0 1 1 1-1a1.004 1.004 0 0 1-1 1M1 6v14a2 2 0 0 0 2 2h18a2 2 0 0 0 2-2V6Zm20 14H3V8h18Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRotate(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 8l-4 4h3c0 3.31-2.69 6-6 6c-1.01 0-1.97-.25-2.8-.7l-1.46 1.46C8.97 19.54 10.43 20 12 20c4.42 0 8-3.58 8-8h3l-4-4zM6 12c0-3.31 2.69-6 6-6c1.01 0 1.97.25 2.8.7l1.46-1.46C15.03 4.46 13.57 4 12 4c-4.42 0-8 3.58-8 8H1l4 4l4-4H6z"><animateTransform attributeName="transform" attributeType="XML" dur="5s" from="360 12 12" repeatCount="indefinite" to="0 12 12" type="rotate"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AtomElectron(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="12" cy="12" r="2.197" fill="currentColor"/><path fill="currentColor" d="M21.63441 6.43671c-.70909-1.22723-2.41155-1.73142-4.79628-1.42354c-.30258.03942-.61354.09379-.92927.15769C15.00856 2.619 13.6066 1 12 1C10.39089 1 8.98713 2.62441 8.087 5.1834c-2.7384-.54268-4.90657-.15669-5.72186 1.25331C1.55256 7.84376 2.2947 9.90626 4.12193 12C2.2947 14.09374 1.55256 16.15624 2.3651 17.56329c.58143 1.00732 1.85 1.49542 3.52453 1.49542a11.47535 11.47535 0 0 0 2.197-.24305C8.98673 21.37518 10.39069 23 12 23c1.60658 0 3.00854-1.619 3.90884-4.17086c.31573.0639.62669.11827.92927.15769a10.18235 10.18235 0 0 0 1.297.088c1.70461 0 2.92647-.52028 3.49932-1.51151c.81254-1.407.07053-3.46955-1.75643-5.56329C21.70494 9.90626 22.447 7.84376 21.63441 6.43671ZM16.9851 6.13956a9.0809 9.0809 0 0 1 1.15965-.08046c1.26048 0 2.14979.32826 2.507.94617c.50426.87335-.06645 2.44138-1.55187 4.16777a20.50554 20.50554 0 0 0-2.30139-1.95281a19.9752 19.9752 0 0 0-.5471-2.93612C16.49539 6.23744 16.7503 6.17013 16.9851 6.13956Zm-3.07236 9.17417c-.64647.37332-1.28408.70346-1.90695.9935c-.63936-.29755-1.2812-.62534-1.919-.9935c-.64888-.37459-1.25525-.76326-1.81959-1.15917c-.06229-.6861-.09688-1.405-.09688-2.15456s.03459-1.46846.09688-2.15456c.56434-.39591 1.17071-.78458 1.81959-1.15917c.63487-.36648 1.27383-.69333 1.91024-.98982c.62668.29132 1.26511.61409 1.91573.98982c.64908.37472 1.25572.76346 1.82019 1.15958c.06222.686.09682 1.40477.09682 2.15415s-.0346 1.46813-.09682 2.15415C15.16846 14.55027 14.56182 14.939 13.91274 15.31373Zm1.628.3351a17.87565 17.87565 0 0 1-.39136 1.82408a18.46424 18.46424 0 0 1-1.76012-.58257c.36293-.18713.72713-.38337 1.092-.594C14.8457 16.086 15.1977 15.86928 15.54078 15.64883Zm-4.92407 1.24372A18.41329 18.41329 0 0 1 8.851 17.474a17.86018 17.86018 0 0 1-.39176-1.82549c.34328.22058.69541.43734 1.06.64787C9.88355 16.50677 10.24989 16.70348 10.61671 16.89255ZM7.07428 13.25036A18.3863 18.3863 0 0 1 5.67548 12a18.38878 18.38878 0 0 1 1.3988-1.25043c-.02005.41033-.03252.82636-.03252 1.25043S7.05423 12.8401 7.07428 13.25036Zm1.385-4.89886a17.86727 17.86727 0 0 1 .391-1.82227a18.22228 18.22228 0 0 1 1.76937.57681c-.36775.18947-.735.38659-1.10031.59759C9.15468 7.91416 8.80255 8.13092 8.45927 8.3515Zm4.93057-1.242a18.48842 18.48842 0 0 1 1.75958-.58237a17.87565 17.87565 0 0 1 .39136 1.82408c-.34308-.22045-.69508-.43715-1.05948-.64754C14.11663 7.4931 13.75263 7.29652 13.38984 7.10946ZM16.92577 10.75A18.40125 18.40125 0 0 1 18.324 12a18.38809 18.38809 0 0 1-1.39826 1.25c.02-.41013.03251-.826.03251-1.25S16.94581 11.1601 16.92577 10.75ZM12 2.12854c.99773 0 2.05613 1.23433 2.80746 3.31044a20.80159 20.80159 0 0 0-2.8094 1.01141A20.51935 20.51935 0 0 0 9.19055 5.44481C9.94208 3.36508 11.00163 2.12854 12 2.12854ZM3.34774 7.00527c.35133-.60825 1.2519-.93437 2.52043-.93437a10.26212 10.26212 0 0 1 1.88362.20148A19.96778 19.96778 0 0 0 7.20147 9.2205a20.48874 20.48874 0 0 0-2.30166 1.95288C3.41412 9.44692 2.84341 7.87869 3.34774 7.00527Zm0 9.98946c-.50433-.87342.06638-2.44165 1.55207-4.16811A20.50145 20.50145 0 0 0 7.20147 14.7795a19.94422 19.94422 0 0 0 .5518 2.95356C5.50578 18.162 3.85468 17.87245 3.34774 16.99473ZM12 21.87146c-.99866 0-2.05847-1.23708-2.81007-3.31775a20.44828 20.44828 0 0 0 2.81155-1.00255a20.80209 20.80209 0 0 0 2.806 1.00979C14.05615 20.63706 12.99775 21.87146 12 21.87146Zm8.65175-4.87673c-.45914.79277-1.79257 1.10923-3.66667.86571c-.2348-.03057-.48971-.09788-.73369-.14455a19.97619 19.97619 0 0 0 .5471-2.93619A20.50419 20.50419 0 0 0 19.0999 12.827C20.58532 14.55335 21.156 16.12138 20.65177 16.99473Z"/><circle cx="-.5" cy="2" r="1.5" fill="currentColor"><animateMotion dur="2s" path="M14.75 14.1471C12.2277 15.6034 9.69019 16.4332 7.6407 16.6145C5.54599 16.7998 4.15833 16.3018 3.62324 15.375C3.08815 14.4482 3.35067 12.9974 4.55852 11.276C5.74031 9.59178 7.72768 7.80915 10.25 6.35289C12.7723 4.89662 15.3098 4.06682 17.3593 3.88549C19.454 3.70016 20.8417 4.1982 21.3768 5.125C21.9118 6.0518 21.6493 7.50256 20.4415 9.22397C19.2597 10.9082 17.2723 12.6909 14.75 14.1471Z" repeatCount="indefinite"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AugmentedReality(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.5 3.311L19 7.653v8.694l-7.5 4.342L4 16.347V7.653zM11.5 1L2 6.5v11l9.5 5.5l9.5-5.5v-11z"/><path fill="currentColor" d="M11.5 5a.5.5 0 1 0 .5.5a.5.5 0 0 0-.5-.5m0 3.5a.5.5 0 1 0 .5.5a.5.5 0 0 0-.5-.5m0 3.5a.5.5 0 1 0 .5.5a.5.5 0 0 0-.5-.5M8 13a.5.5 0 1 0 .422.231A.498.498 0 0 0 8 13m-2.5 2a.5.5 0 1 0 .421.231a.498.498 0 0 0-.422-.23M15 13a.5.5 0 1 0 .269.079a.5.5 0 0 0-.269-.08M17.535 15a.5.5 0 1 0 .269.079a.5.5 0 0 0-.269-.079M3.382 6.225l-1 1.732l7.027 4.057l1-1.732zM12.5 17.016h-2v5.141h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Autoinstallation(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.424 16.5h5.12v-6h4.448a9.004 9.004 0 1 0-6.458 9.11Zm-1.908-5.526a2.5 2.5 0 1 1 2.5 2.5a2.5 2.5 0 0 1-2.5-2.5"/><path fill="currentColor" d="M22.045 18h-3v-6h-2v6h-3l4 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BackgroundTasks(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 3h8.99v1.5H6zM2.99 6h1.5v1.5h-1.5zm0-3h1.5v1.5h-1.5zm0 6.01H4.5v1.5H2.99z"/><path fill="currentColor" d="M4.5 12h-3V1.49h15V6H18V2a2 2 0 0 0-2-2H2a2 2 0 0 0-2 2v9.48a2 2 0 0 0 2 2h2.5Z"/><path fill="currentColor" d="M22 7.5H8a2 2 0 0 0-2 2V19a2 2 0 0 0 2 2h5.53v1.53H12V24h6v-1.49h-1.5V21H22a2 2 0 0 0 2-2V9.5a2 2 0 0 0-2-2m.51 12h-15V9h15Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BigData(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.965 8a5.35 5.35 0 0 0-3.615-1.96a7.492 7.492 0 0 0-14-2A5.904 5.904 0 0 0 4 4.365A5.98 5.98 0 0 0 4 15.65V8h16v7.9a5.003 5.003 0 0 0 4-4.9a4.908 4.908 0 0 0-1.035-3M18.02 22.003c0 1.103-2.687 1.997-6 1.997s-6-.894-6-1.997v-1.996c0 1.102 2.686 1.996 6 1.996s6-.894 6-1.996"/><path fill="currentColor" d="M12.02 16.684c-3.311 0-6-.898-6-1.996v3.964c0 1.099 2.689 1.997 6 1.997s6-.898 6-1.997v-3.964c0 1.098-2.69 1.996-6 1.996"/><path fill="currentColor" d="M18.02 11.997c0-1.103-2.687-1.997-6-1.997s-6 .894-6 1.997v1.33c0 1.104 2.686 1.998 6 1.998s6-.894 6-1.997"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BigDataOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 6.147a4.973 4.973 0 0 0-.65-.107a7.492 7.492 0 0 0-14-2A5.904 5.904 0 0 0 4 4.365A5.98 5.98 0 0 0 4 15.65v-2.204a3.976 3.976 0 0 1 0-6.901a3.93 3.93 0 0 1 1.56-.515l1.07-.11l.5-.95a5.487 5.487 0 0 1 10.26 1.46l.3 1.5l1.53.11a2.913 2.913 0 0 1 .78.171a2.963 2.963 0 0 1 0 5.604v2.084a4.972 4.972 0 0 0 0-9.752"/><path fill="currentColor" d="M12 11c-3.818 0-6 .758-6 2.167v8.666C6 23.242 9.087 24 12 24s6-.758 6-2.167v-8.666C18 11.758 15.818 11 12 11m0 1c2.711 0 4.91.81 4.91 1.708s-2.2 1.625-4.91 1.625s-4.91-.727-4.91-1.625S9.29 12 12 12m-4.91 7.924v-1.143a12.026 12.026 0 0 0 4.91.886a12.026 12.026 0 0 0 4.91-.886v1.179a12.588 12.588 0 0 1-4.91.79a12.776 12.776 0 0 1-4.91-.826"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Blockchain(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 8L9 9.75v3.5L12 15l3-1.75v-3.5Zm1.517 2.04l-1.53.892l-1.517-.885L12 9.155Zm-3.527.882l1.484.866v1.75l-1.484-.865Zm2.474 2.653v-1.767l1.546-.902v1.767ZM3 0L0 1.75v3.5L3 7l3-1.75v-3.5Zm1.517 2.04l-1.53.892l-1.517-.885L3 1.155ZM.99 2.921l1.484.866v1.75L.99 4.674Zm2.474 2.653V3.808l1.546-.902v1.767ZM3 17l-3 1.75v3.5L3 24l3-1.75v-3.5Zm1.517 2.04l-1.53.892l-1.517-.885L3 18.155Zm-3.527.882l1.484.866v1.75L.99 21.674Zm2.474 2.653v-1.767l1.546-.902v1.767ZM21 0l-3 1.75v3.5L21 7l3-1.75v-3.5Zm1.517 2.04l-1.53.892l-1.517-.885L21 1.155Zm-3.527.882l1.484.866v1.75l-1.484-.865Zm2.474 2.653V3.808l1.546-.902v1.767ZM21 17l-3 1.75v3.5L21 24l3-1.75v-3.5Zm1.517 2.04l-1.53.892l-1.517-.885l1.53-.892Zm-3.527.882l1.484.866v1.75l-1.484-.865Zm2.474 2.653v-1.767l1.546-.902v1.767ZM9 3h6v1H9zm0 17h6v1H9zM3.5 9v6h-1V9zm3.793-.172L5.172 6.707L5.879 6L8 8.12zM16 8.293l2.121-2.121l.707.707L16.707 9zm-7.872 6.586L6.007 17l-.707-.707l2.121-2.121zm8.751-.75L19 16.25l-.707.707l-2.121-2.121zM21.5 9v6h-1V9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bootstrapping(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 22h18v2H3zm9-10c-12 0-9 9-9 9h18s3-9-9-9"/><path fill="currentColor" d="m3 0l1 13c1.22-1.6 3.46-2.51 8-2.51s6.81.93 8 2.53L21 0Zm4 3h10v2H7Zm11 6H6V7h12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BootstrappingOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.667 12c-12 0-9 9-9 9v3h18v-3s3-9-9-9m7 10h-14v-1h14Zm0-2h-14c-.243-1.937-.117-6 7-6s7.243 4.064 7 6M8 6h8v2H8zm1-3h6v2H9z"/><path fill="currentColor" d="M12 0H3l1 13c1-1 2.033-2.5 8-2.5c5.97 0 7 1.5 8 2.5l1-13Zm6.308 9H5.692l-.538-7h13.692Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Branch(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 5a3 3 0 1 0-4 2.816V11H7V7.816a3 3 0 1 0-2 0v8.368a3 3 0 1 0 2 0V13h10a2 2 0 0 0 2-2V7.816A2.991 2.991 0 0 0 21 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BranchOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 2a2.993 2.993 0 0 0-1 5.816V11H7V7.816a3 3 0 1 0-2 0v8.368a3 3 0 1 0 2 0V13h10a2 2 0 0 0 2-2V7.816A2.993 2.993 0 0 0 18 2M6 20a1 1 0 1 1 1-1a1 1 0 0 1-1 1M6 6a1 1 0 1 1 1-1a1 1 0 0 1-1 1m12 0a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BubbleLoading(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="12" cy="2" r="0" fill="currentColor"><animate attributeName="r" begin="0" calcMode="spline" dur="1s" keySplines="0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8" repeatCount="indefinite" values="0;2;0;0"/></circle><circle cx="12" cy="2" r="0" fill="currentColor" transform="rotate(45 12 12)"><animate attributeName="r" begin="0.125s" calcMode="spline" dur="1s" keySplines="0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8" repeatCount="indefinite" values="0;2;0;0"/></circle><circle cx="12" cy="2" r="0" fill="currentColor" transform="rotate(90 12 12)"><animate attributeName="r" begin="0.25s" calcMode="spline" dur="1s" keySplines="0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8" repeatCount="indefinite" values="0;2;0;0"/></circle><circle cx="12" cy="2" r="0" fill="currentColor" transform="rotate(135 12 12)"><animate attributeName="r" begin="0.375s" calcMode="spline" dur="1s" keySplines="0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8" repeatCount="indefinite" values="0;2;0;0"/></circle><circle cx="12" cy="2" r="0" fill="currentColor" transform="rotate(180 12 12)"><animate attributeName="r" begin="0.5s" calcMode="spline" dur="1s" keySplines="0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8" repeatCount="indefinite" values="0;2;0;0"/></circle><circle cx="12" cy="2" r="0" fill="currentColor" transform="rotate(225 12 12)"><animate attributeName="r" begin="0.625s" calcMode="spline" dur="1s" keySplines="0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8" repeatCount="indefinite" values="0;2;0;0"/></circle><circle cx="12" cy="2" r="0" fill="currentColor" transform="rotate(270 12 12)"><animate attributeName="r" begin="0.75s" calcMode="spline" dur="1s" keySplines="0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8" repeatCount="indefinite" values="0;2;0;0"/></circle><circle cx="12" cy="2" r="0" fill="currentColor" transform="rotate(315 12 12)"><animate attributeName="r" begin="0.875s" calcMode="spline" dur="1s" keySplines="0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8" repeatCount="indefinite" values="0;2;0;0"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cleanup(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 15.016h-5v-.013a1.984 1.984 0 0 0-1.978-1.978h-.035l3.857-10.393l-1.732-.643l-4.095 11.036h-.039A1.984 1.984 0 0 0 9 15.003v.013H4v2h16Zm0 3.002H4L2 22h20z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudComputing(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 10v5h12v-5Zm7 3H7v-1h6Zm1.5 0a.5.5 0 1 1 .5-.5a.5.5 0 0 1-.5.5m2 0a.5.5 0 1 1 .5-.5a.5.5 0 0 1-.5.5M6 17v5h12v-5Zm7 3H7v-1h6Zm1.5 0a.5.5 0 1 1 .5-.5a.5.5 0 0 1-.5.5m2 0a.5.5 0 1 1 .5-.5a.5.5 0 0 1-.5.5"/><path fill="currentColor" d="M22.965 8a5.35 5.35 0 0 0-3.615-1.96a7.492 7.492 0 0 0-14-2A5.904 5.904 0 0 0 4 4.365A5.98 5.98 0 0 0 4 15.65V8h16v7.9a5.003 5.003 0 0 0 4-4.9a4.908 4.908 0 0 0-1.035-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudComputingOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 6.147a4.973 4.973 0 0 0-.65-.107a7.492 7.492 0 0 0-14-2A5.904 5.904 0 0 0 4 4.365A5.98 5.98 0 0 0 4 15.65v-2.204a3.976 3.976 0 0 1 0-6.901a3.93 3.93 0 0 1 1.56-.515l1.07-.11l.5-.95a5.487 5.487 0 0 1 10.26 1.46l.3 1.5l1.53.11a2.913 2.913 0 0 1 .78.171a2.963 2.963 0 0 1 0 5.604v2.084a4.972 4.972 0 0 0 0-9.752"/><path fill="currentColor" d="M6 10v5h12v-5Zm7 3H7v-1h6Zm1.5 0a.5.5 0 1 1 .5-.5a.5.5 0 0 1-.5.5m2 0a.5.5 0 1 1 .5-.5a.5.5 0 0 1-.5.5M6 17v5h12v-5Zm7 3H7v-1h6Zm1.5 0a.5.5 0 1 1 .5-.5a.5.5 0 0 1-.5.5m2 0a.5.5 0 1 1 .5-.5a.5.5 0 0 1-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudControllerManager(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 10h8v2H9Zm15 0h-3V8h-2v6h2v-2h3ZM9 20h8v2H9Zm15 0h-3v-2h-2v6h2v-2h3Zm0-3h-8v-2h8ZM9 17h3v2h2v-6h-2v2H9Z"/><path fill="currentColor" d="M19 6H7v10s-7 .17-7-6a6.071 6.071 0 0 1 1.235-4a5.99 5.99 0 0 1 4.459-1.99a4.86 4.86 0 0 1 .572-.99l.272-.36l.23-.26a3.597 3.597 0 0 1 .262-.29A6.897 6.897 0 0 1 12 0c.16 0 .362.01.513.02a7.012 7.012 0 0 1 3.756 1.39a8.546 8.546 0 0 1 .913.81A7.043 7.043 0 0 1 19 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cluster(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 13v10h22V13Zm12 6H4v-2h9Zm3 0a1 1 0 1 1 1-1a1 1 0 0 1-1 1m3 0a1 1 0 1 1 1-1a1 1 0 0 1-1 1M1 1v10h22V1Zm12 6H4V5h9Zm3 0a1 1 0 1 1 1-1a1 1 0 0 1-1 1m3 0a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClusterManagement(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 1v10h22V1Zm12 6H4V5h9Zm3 0a1 1 0 1 1 1-1a1 1 0 0 1-1 1m3 0a1 1 0 1 1 1-1a1 1 0 0 1-1 1m3.69 11.37l1.14-1l-1-1.73l-1.45.49a3.647 3.647 0 0 0-1.08-.63L20 14h-2l-.3 1.49a3.646 3.646 0 0 0-1.08.63l-1.45-.49l-1 1.73l1.14 1a3.337 3.337 0 0 0 0 1.26l-1.14 1l1 1.73l1.45-.49a3.645 3.645 0 0 0 1.08.63L18 24h2l.3-1.49a3.646 3.646 0 0 0 1.08-.63l1.45.49l1-1.73l-1.14-1a3.39 3.39 0 0 0 0-1.27M19 21a2 2 0 1 1 2-2a2.006 2.006 0 0 1-2 2"/><path fill="currentColor" d="M12 19H4v-2h8.294a7.008 7.008 0 0 1 3.114-4H1v10h12.26A6.962 6.962 0 0 1 12 19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClusterManagementOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m22.69 18.37l1.14-1l-1-1.73l-1.45.49a3.647 3.647 0 0 0-1.08-.63L20 14h-2l-.3 1.49a3.646 3.646 0 0 0-1.08.63l-1.45-.49l-1 1.73l1.14 1a3.337 3.337 0 0 0 0 1.26l-1.14 1l1 1.73l1.45-.49a3.645 3.645 0 0 0 1.08.63L18 24h2l.3-1.49a3.646 3.646 0 0 0 1.08-.63l1.45.49l1-1.73l-1.14-1a3.39 3.39 0 0 0 0-1.27M19 21a2 2 0 1 1 2-2a2.006 2.006 0 0 1-2 2m4-10H1V1h22ZM3 9h18V3H3Zm10-4H4v2h9Zm3 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1m3 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1"/><path fill="currentColor" d="M12.294 21H3v-6h10.26a7.026 7.026 0 0 1 2.148-2H1v10h12.26a6.962 6.962 0 0 1-.966-2"/><path fill="currentColor" d="M4 19h8a6.994 6.994 0 0 1 .294-2H4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClusterOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23 11H1V1h22ZM3 9h18V3H3Zm20 14H1V13h22ZM3 21h18v-6H3Z"/><path fill="currentColor" d="M4 5h9v2H4z"/><circle cx="16" cy="6" r="1" fill="currentColor"/><circle cx="19" cy="6" r="1" fill="currentColor"/><path fill="currentColor" d="M4 17h9v2H4z"/><circle cx="16" cy="18" r="1" fill="currentColor"/><circle cx="19" cy="18" r="1" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClusterRole(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 4h-4.18a3 3 0 0 0-5.64 0H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2m-7 0a1 1 0 1 1-1 1a1 1 0 0 1 1-1m7 15H5v-4h14Zm0-6H5V9h14Z"/><circle cx="17" cy="11" r="1" fill="currentColor"/><circle cx="14" cy="11" r="1" fill="currentColor"/><circle cx="14" cy="17" r="1" fill="currentColor"/><circle cx="17" cy="17" r="1" fill="currentColor"/><path fill="currentColor" d="M6 10h5v2H6zm0 6h5v2H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClusterRoleBinding(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 10h5v2H5Zm3.54 12H4a2.006 2.006 0 0 1-2-2V6a2.006 2.006 0 0 1 2-2h4.18a2.988 2.988 0 0 1 5.64 0H18a2.006 2.006 0 0 1 2 2v6.09a5.989 5.989 0 0 0-1-.09h-1V9H4v4h5.69a6.04 6.04 0 0 0-1.878 2H4v4h3.09a5.973 5.973 0 0 0 1.45 3M10 5a1 1 0 1 0 1-1a1.003 1.003 0 0 0-1 1m7 6a1 1 0 1 0-1 1a1 1 0 0 0 1-1M5 18h2a5.96 5.96 0 0 1 .35-2H5Zm9-7a1 1 0 1 0-1 1a1 1 0 0 0 1-1m1 9h-2a2 2 0 0 1 0-4h2v-2h-2a4 4 0 0 0 0 8h2Zm8-2a4 4 0 0 1-4 4h-2v-2h2a2 2 0 0 0 0-4h-2v-2h2a4 4 0 0 1 4 4m-5 1h-4v-2h4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeDeploy(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.95 9h-.27l-4.16 6.02l-.82-.57L20.47 9h-2.34a1.25 1.25 0 1 0-2.5 0h-.43l-1.25 4l-.95-.3L14.15 9h-.9a1.25 1.25 0 0 0-2.5 0h-.9L11 12.7l-.95.3L8.8 9h-.42a1.25 1.25 0 0 0-2.5 0H3.51l3.77 5.45l-.82.57L2.3 9h-.25a10 10 0 0 1 19.9 0"/><path fill="currentColor" d="M11 15.405L9.586 14l-3.618 3.595L4.554 19l5.032 5L11 22.595L7.382 19zm2.021 7.19L14.435 24l3.618-3.595L19.467 19l-5.032-5l-1.414 1.405L16.639 19z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeDeployOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.95 9a10 10 0 0 0-19.9 0h.248l4.162 6.016l.823-.57L3.514 9h2.362a1.25 1.25 0 0 1 2.5 0h.422l1.247 4l.955-.298L9.845 9h.906a1.25 1.25 0 1 1 2.5 0h.904L13 12.702l.955.298l1.247-4h.424a1.25 1.25 0 1 1 2.5 0h2.343L16.7 14.447l.822.569L21.685 9Zm-5.074-3.25a3.23 3.23 0 0 0-2.437 1.123a3.206 3.206 0 0 0-4.875 0a3.207 3.207 0 0 0-4.959.101a7.984 7.984 0 0 1 14.788-.004a3.231 3.231 0 0 0-2.517-1.22"/><path fill="currentColor" d="M11 15.405L9.586 14l-3.618 3.595L4.554 19l5.032 5L11 22.595L7.382 19zm2.021 7.19L14.435 24l3.618-3.595L19.467 19l-5.032-5l-1.414 1.405L16.639 19z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Collocation(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 2H4a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2m0 18H4V4h16Z"/><path fill="currentColor" d="M7.51 18A2.5 2.5 0 0 0 10 16h4.1a2.5 2.5 0 1 0 2.45-3h-.05l-2.42-4.1a2.5 2.5 0 1 0-4.14 0L7.56 13a2.5 2.5 0 0 0 0 5Zm3.17-8.39a2.46 2.46 0 0 0 2.66 0l2.1 3.64A2.49 2.49 0 0 0 14.06 15H10a2.49 2.49 0 0 0-1.38-1.75Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Commit(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 11h-6.141a3.981 3.981 0 0 0-7.718 0H2v2h6.141a3.981 3.981 0 0 0 7.718 0H22Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommitOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.859 11a3.981 3.981 0 0 0-7.718 0H2v2h6.141a3.981 3.981 0 0 0 7.718 0H22v-2ZM12 13a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompareStates(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 16.184V6a2 2 0 0 0-2-2h-5.172l1.586-1.586L14 1l-2.586 2.586L10 5l1.414 1.414L14 9l1.414-1.414L13.828 6H19v10.184a3 3 0 1 0 2 0m-8.414 1.402L10 15l-1.414 1.414L10.172 18H5V7.816a3 3 0 1 0-2 0V18a2 2 0 0 0 2 2h5.172l-1.586 1.586L10 23l2.586-2.586L14 19z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompareStatesOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10 15l-1.414 1.414L10.172 18H5V7.816a3 3 0 1 0-2 0V18a2 2 0 0 0 2 2h5.172l-1.586 1.586L10 23l4-4ZM4 6a1 1 0 1 1 1-1a1 1 0 0 1-1 1m17 10.184V6a2 2 0 0 0-2-2h-5.172l1.586-1.586L14 1l-4 4l4 4l1.414-1.414L13.828 6H19v10.184a3 3 0 1 0 2 0M20 20a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compass(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 10.9c-.61 0-1.1.49-1.1 1.1s.49 1.1 1.1 1.1c.61 0 1.1-.49 1.1-1.1s-.49-1.1-1.1-1.1zM12 2C6.48 2 2 6.48 2 12s4.48 10 10 10s10-4.48 10-10S17.52 2 12 2zm2.19 12.19L6 18l3.81-8.19L18 6l-3.81 8.19z"><animateTransform id="eosIconsCompass0" attributeName="transform" attributeType="XML" begin="0;eosIconsCompass2.end" dur="1s" from="-90 12 12" to="0 12 12" type="rotate"/><animateTransform id="eosIconsCompass1" attributeName="transform" attributeType="XML" begin="eosIconsCompass0.end" dur="1s" from="0 12 12" to="-90 12 12" type="rotate"/><animateTransform id="eosIconsCompass2" attributeName="transform" attributeType="XML" begin="eosIconsCompass1.end" dur="1s" from="-90 12 12" to="270 12 12" type="rotate"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ConfigMap(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.822 13.34l-1.07-.83a3.883 3.883 0 0 0 .04-.49a2.626 2.626 0 0 0-.04-.49l1.06-.83a.26.26 0 0 0 .06-.32l-1-1.73a.254.254 0 0 0-.31-.11l-1.24.5a3.42 3.42 0 0 0-.85-.49l-.19-1.32a.239.239 0 0 0-.24-.21h-2a.257.257 0 0 0-.25.21l-.19 1.32a3.995 3.995 0 0 0-.85.49l-1.24-.5a.262.262 0 0 0-.31.11l-1 1.73a.248.248 0 0 0 .06.32l1.06.83a4.013 4.013 0 0 0 0 .98l-1.06.83a.26.26 0 0 0-.06.32l1 1.73a.254.254 0 0 0 .31.11l1.24-.5a3.42 3.42 0 0 0 .85.49l.19 1.32a.249.249 0 0 0 .25.21h2a.257.257 0 0 0 .25-.21l.19-1.32a3.695 3.695 0 0 0 .84-.49l1.25.5a.262.262 0 0 0 .31-.11l1-1.73a.26.26 0 0 0-.06-.32m-4.78.18a1.5 1.5 0 1 1 1.5-1.5a1.498 1.498 0 0 1-1.5 1.5M19.997.015h-4a.999.999 0 0 0-.995 1L15 3H4.997v3.01l2.012.012v-1.01L14.999 5l-.002 4.015a.999.999 0 0 0 .995 1h6.005a1.003 1.003 0 0 0 1-1v-6Zm-.5 3.5V.765l2.75 2.75Zm.501 10.494h-4a.999.999 0 0 0-.995 1l-.002 4.01l-7.99.004v-1.01H5v3.01l10-.004l-.002 1.99a.999.999 0 0 0 .995 1h6.005a1.003 1.003 0 0 0 1-1v-6Zm-.5 3.5v-2.75l2.75 2.75Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ConfigurationFile(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.507 14.142a1.35 1.35 0 1 0 1.35 1.35a1.348 1.348 0 0 0-1.35-1.35"/><path fill="currentColor" d="M14 2.01H6a1.997 1.997 0 0 0-1.99 2l-.01 16a1.997 1.997 0 0 0 1.99 2H18a2.006 2.006 0 0 0 2-2v-12Zm.863 14.958l-.9 1.557a.236.236 0 0 1-.279.099l-1.125-.45a3.328 3.328 0 0 1-.756.44l-.17 1.189a.231.231 0 0 1-.226.189h-1.8a.224.224 0 0 1-.225-.19l-.17-1.187a3.079 3.079 0 0 1-.766-.441l-1.116.45a.228.228 0 0 1-.279-.1l-.9-1.556a.234.234 0 0 1 .054-.288l.954-.747a3.618 3.618 0 0 1 0-.882l-.954-.747a.223.223 0 0 1-.054-.288l.9-1.557a.236.236 0 0 1 .28-.1l1.115.45a3.593 3.593 0 0 1 .765-.44l.171-1.188a.231.231 0 0 1 .225-.19h1.8a.215.215 0 0 1 .216.19l.171 1.188a3.078 3.078 0 0 1 .765.44l1.116-.45a.228.228 0 0 1 .28.1l.9 1.557a.234.234 0 0 1-.055.288l-.954.747a2.368 2.368 0 0 1 .036.44a3.495 3.495 0 0 1-.036.442l.963.747a.234.234 0 0 1 .054.288M13 9.01v-5.5l5.5 5.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ConfigurationFileOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.691 15.778l-.63-.49a2.29 2.29 0 0 0 .023-.288a1.548 1.548 0 0 0-.024-.289l.625-.49a.153.153 0 0 0 .036-.188l-.59-1.02a.15.15 0 0 0-.183-.065l-.73.295a2.016 2.016 0 0 0-.502-.289l-.112-.778a.14.14 0 0 0-.141-.124h-1.18a.151.151 0 0 0-.147.124l-.112.778a2.355 2.355 0 0 0-.5.29l-.732-.296a.154.154 0 0 0-.183.065l-.59 1.02a.146.146 0 0 0 .036.189l.625.49a2.366 2.366 0 0 0 0 .577l-.625.49a.153.153 0 0 0-.035.188l.59 1.02a.15.15 0 0 0 .182.065l.731-.295a2.016 2.016 0 0 0 .501.289l.112.778a.147.147 0 0 0 .148.124h1.179a.151.151 0 0 0 .147-.124l.112-.778a2.178 2.178 0 0 0 .495-.29l.737.296a.154.154 0 0 0 .183-.065l.59-1.02a.153.153 0 0 0-.036-.189m-2.818.106a.884.884 0 1 1 .885-.884a.883.883 0 0 1-.885.884"/><path fill="currentColor" d="M14 2H6a2.006 2.006 0 0 0-2 2v16a2.006 2.006 0 0 0 2 2h12a2.006 2.006 0 0 0 2-2V8Zm4 18H6V4h7v5h5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Constraint(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 15v-1.5a2.5 2.5 0 0 0-5 0V15a1.075 1.075 0 0 0-1 1v4a1.075 1.075 0 0 0 1 1h5a1.075 1.075 0 0 0 1-1v-4a1.075 1.075 0 0 0-1-1m-1 0h-3v-1.5a1.5 1.5 0 0 1 3 0Zm2-10H2V3h20Zm0 4H2V7h20Zm-9 4H2v-2h11Zm0 4H2v-2h11Zm0 4H2v-2h11Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Container(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 5a1.958 1.958 0 0 0 .796 1.56L5.5 9.986H3.997A1.997 1.997 0 0 0 2 11.983v8.023a1.997 1.997 0 0 0 1.997 1.997h16A2.004 2.004 0 0 0 22 19.999v-8.016a1.997 1.997 0 0 0-1.997-1.997h-1.491L13.21 6.555A1.957 1.957 0 0 0 14 5a2.03 2.03 0 0 0-1-1.721V2h-1v2a1 1 0 1 1-1 1m.995 1.974l.005.001l.017-.002l4.655 3.013h-9.33ZM7 18a1 1 0 0 1-2 0v-4a1 1 0 0 1 2 0Zm4 0a1 1 0 0 1-2 0v-4a1 1 0 0 1 2 0Zm4 0a1 1 0 0 1-2 0v-4a1 1 0 0 1 2 0Zm4 0a1 1 0 0 1-2 0v-4a1 1 0 0 1 2 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ContainerOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 10h-2l-4.79-3.5A1.854 1.854 0 0 0 14 5a2.03 2.03 0 0 0-1-1.721V2h-1v2a1 1 0 1 1-1 1h-1a1.846 1.846 0 0 0 .796 1.5L6 10H4a2.002 2.002 0 0 0-2 2v8a2.002 2.002 0 0 0 2 2h16a1.997 1.997 0 0 0 2-2v-8a2.002 2.002 0 0 0-2-2m-8-3c.006 0 4 3 4 3H8Zm8 13H4v-8h16Z"/><path fill="currentColor" d="M14 19a1 1 0 0 0 1-1v-4a1 1 0 0 0-2 0v4a1 1 0 0 0 1 1m4 0a1 1 0 0 0 1-1v-4a1 1 0 0 0-2 0v4a1 1 0 0 0 1 1M6 19a1 1 0 0 0 1-1v-4a1 1 0 0 0-2 0v4a1 1 0 0 0 1 1m4 0a1 1 0 0 0 1-1v-4a1 1 0 0 0-2 0v4a1 1 0 0 0 1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ContentDeleted(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 19H5V5h6V3H5a2.006 2.006 0 0 0-2 2v14a2.006 2.006 0 0 0 2 2h14a2.006 2.006 0 0 0 2-2v-4h-2Z"/><path fill="currentColor" d="M15 5h6v6a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1zm7-2h-2l-.571-1h-2.858L16 3h-2v1h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ContentLifecycleManagement(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 7h8v2H6zm0 4h12v2H6zm0 4h2.99v2H6z"/><path fill="currentColor" d="m14 3l-3-3v2H4a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h4v-2H4V4h7v2Zm-4 18l3 3v-2h7a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2h-4v2h4v16h-7v-2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ContentModified(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.558 3.592l-1.15-1.15a1.49 1.49 0 0 0-2.12 0L13 7.731V11h3.27l5.288-5.288a1.49 1.49 0 0 0 0-2.12M15.579 9.45h-1.03V8.42L18 4.973l1.03 1.03Z"/><path fill="currentColor" d="M19 19H5V5h6V3H5a2.006 2.006 0 0 0-2 2v14a2.006 2.006 0 0 0 2 2h14a2.006 2.006 0 0 0 2-2v-6h-2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ContentNew(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 5v2h-3v3h-2V7h-3V5h3V2h2v3Zm-3 14H5V5h6V3H5a2.006 2.006 0 0 0-2 2v14a2.006 2.006 0 0 0 2 2h14a2.006 2.006 0 0 0 2-2v-6h-2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Counting(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 2h2v20H2z"/><path fill="currentColor" d="M22 4v2H2V4z"/><path fill="currentColor" d="M20 2h2v20h-2z"/><circle cx="8" cy="5" r="2" fill="currentColor"/><path fill="currentColor" d="M22 10v2H2v-2z"/><circle cx="10" cy="11" r="2" fill="currentColor"/><circle cx="16" cy="11" r="2" fill="currentColor"/><path fill="currentColor" d="M22 16v2H2v-2zm2 5v2H0v-2z"/><circle cx="8" cy="17" r="2" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CriticalBug(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 8h-2.8a6 6 0 0 0-1.82-2L17 4.41L15.6 3l-2.17 2.17A6.07 6.07 0 0 0 12 5a5.92 5.92 0 0 0-1.41.17L8.42 3L7 4.41L8.63 6a6.06 6.06 0 0 0-1.81 2H4v2h2.1a6.64 6.64 0 0 0-.1 1v1H4v2h2v1a6.64 6.64 0 0 0 .09 1H4v2h2.82a6 6 0 0 0 10.38 0H20v-2h-2.08a6.64 6.64 0 0 0 .08-1v-1h2v-2h-2v-1a6.64 6.64 0 0 0-.09-1H20Zm-7 10h-2v-2h2Zm0-4h-2V8h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CriticalBugOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 15h2v2h-2zm0-6h2v5h-2z"/><path fill="currentColor" d="M20 8h-2.81a5.985 5.985 0 0 0-1.82-1.96L17 4.41L15.59 3l-2.17 2.17A6.066 6.066 0 0 0 12 5a5.919 5.919 0 0 0-1.41.17L8.41 3L7 4.41l1.62 1.63A6.062 6.062 0 0 0 6.81 8H4v2h2.09A6.63 6.63 0 0 0 6 11v1H4v2h2v1a6.632 6.632 0 0 0 .09 1H4v2h2.81a5.99 5.99 0 0 0 10.38 0H20v-2h-2.09a6.632 6.632 0 0 0 .09-1v-1h2v-2h-2v-1a6.63 6.63 0 0 0-.09-1H20Zm-4 4v3a4.26 4.26 0 0 1-.07.7l-.1.65l-.37.65a3.993 3.993 0 0 1-6.92 0l-.37-.64l-.1-.65A4.271 4.271 0 0 1 8 15v-4a4.053 4.053 0 0 1 .07-.7l.1-.65l.37-.65a4.1 4.1 0 0 1 1.21-1.31l.57-.39l.74-.18A3.786 3.786 0 0 1 12 7a3.865 3.865 0 0 1 .95.12l.68.16l.61.42a3.894 3.894 0 0 1 1.21 1.31l.38.65l.1.65A4.038 4.038 0 0 1 16 11Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cronjob(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.021 1.021h6v2h-6zM20.04 7.41l1.434-1.434l-1.414-1.414l-1.433 1.433A8.989 8.989 0 0 0 7.53 5.881l1.42 1.44a7.038 7.038 0 0 1 4.06-1.3l.01.001v6.98l4.953 4.958A7.001 7.001 0 0 1 6.01 13.021h3l-4-4l-4 4h3A9 9 0 1 0 20.04 7.41"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CsvFile(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.959 20.046H4V4.004h6.979v4.98h4.98V10h2.051V8.014l-2.05-2.052L14 4.004l-1.99-1.99h-8a1.997 1.997 0 0 0-1.99 2l-.01 16a1.997 1.997 0 0 0 1.99 2h14.01v-2Zm5.301-2.032l1.75-6h-1.5l-1 3.43l-1-3.43h-1.5l1.75 6z"/><path fill="currentColor" d="M10.01 12.014h-3a1.003 1.003 0 0 0-1 1v4a1.003 1.003 0 0 0 1 1h3a1.003 1.003 0 0 0 1-1v-1h-1.5v.5h-2v-3h2v.5h1.5v-1a1.003 1.003 0 0 0-1-1m7 1.506v-1.506h-4a1 1 0 0 0-1 1v1.757a1 1 0 0 0 1 1h2.51v.743h-3.51v1.507h4a1 1 0 0 0 1-1v-1.757a1 1 0 0 0-1-1H13.5v-.743Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Daemon(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.231 9.366q.129-.206.24-.426a4.857 4.857 0 0 0 .2-.455a5.053 5.053 0 0 0-.104-3.84a4.942 4.942 0 0 0-2.772-2.637l-.016-.004L17.764 2a5.033 5.033 0 0 1 .446 4.476l-.002.003l-.001.004A9.041 9.041 0 0 0 15.4 4.655a8.975 8.975 0 0 0-6.804 0A9.041 9.041 0 0 0 5.79 6.483l-.002-.004l-.001-.003a4.872 4.872 0 0 1-.276-2.331A5.21 5.21 0 0 1 6.282 2l-.015.004l-.015.004a5.044 5.044 0 0 0-2.926 6.477a4.854 4.854 0 0 0 .2.455q.11.22.24.426a9 9 0 1 0 16.465 0m-4.734-.378a1.498 1.498 0 1 1-1.061.439a1.494 1.494 0 0 1 1.06-.44m-6.997 0a1.498 1.498 0 1 1-1.06.439a1.494 1.494 0 0 1 1.06-.44m3.498 9.532a5.431 5.431 0 0 1-3.127-.986a5.548 5.548 0 0 1-1.98-2.516h10.215a5.548 5.548 0 0 1-1.98 2.516a5.431 5.431 0 0 1-3.128.986"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DaemonOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.5 9a1.492 1.492 0 0 0-1.064.427A1.522 1.522 0 0 0 14 10.5a1.468 1.468 0 0 0 .436 1.05A1.511 1.511 0 0 0 15.5 12a1.496 1.496 0 0 0 1.058-.45A1.474 1.474 0 0 0 17 10.5a1.527 1.527 0 0 0-.442-1.073A1.476 1.476 0 0 0 15.5 9m-7 0a1.499 1.499 0 1 0 1.06.439A1.496 1.496 0 0 0 8.5 9"/><path fill="currentColor" d="M18 2a16.282 16.282 0 0 1 .51 2.145a5.04 5.04 0 0 1-.083 1.56a8.987 8.987 0 0 0-12.853 0a4.922 4.922 0 0 1-.063-1.56A4.908 4.908 0 0 1 6 2a9.821 9.821 0 0 0-2.546 2.645a4.996 4.996 0 0 0-.128 3.84a4.851 4.851 0 0 0 .2.455l.012.02a9 9 0 1 0 16.923-.003l.01-.017a4.852 4.852 0 0 0 .2-.455a5.053 5.053 0 0 0-.104-3.84A9.599 9.599 0 0 0 18 2m-6 17a7 7 0 1 1 7-7a7.008 7.008 0 0 1-7 7"/><path fill="currentColor" d="M7 14a5.493 5.493 0 0 0 5 3a5.493 5.493 0 0 0 5-3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DaemonSet(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.005 21.995a2 2 0 0 1-2-2V6.005h2v13.99h15.99v2zM16.098 9.994a.896.896 0 1 0-.35-.071a.898.898 0 0 0 .35.07m-5.164 1.817a3.33 3.33 0 0 0 1.189 1.51a3.269 3.269 0 0 0 3.752 0a3.329 3.329 0 0 0 1.189-1.51zm.966-1.817a.896.896 0 1 0-.35-.071a.898.898 0 0 0 .35.07M23 4.064v11.872A2.064 2.064 0 0 1 20.936 18H7.064A2.064 2.064 0 0 1 5 15.936V4.064A2.064 2.064 0 0 1 7.064 2h13.872A2.064 2.064 0 0 1 23 4.064m-4.061 4.355q.038-.061.074-.125t.07-.13l.063-.134q.03-.068.056-.14a3.032 3.032 0 0 0-.062-2.303a3 3 0 0 0-.658-.945a2.949 2.949 0 0 0-1.005-.637l-.005-.002h-.004L17.463 4h-.005a2.978 2.978 0 0 1 .298.616a3.028 3.028 0 0 1-.03 2.07l-.001.002l-.001.002a5.46 5.46 0 0 0-.788-.625a5.368 5.368 0 0 0-5.874 0a5.461 5.461 0 0 0-.788.625v-.001l-.001-.001v-.001l-.001-.002a2.894 2.894 0 0 1-.17-.699a2.967 2.967 0 0 1 .005-.7a3.091 3.091 0 0 1 .16-.67A3.145 3.145 0 0 1 10.57 4l-.004.001l-.005.001l-.004.002h-.005a2.997 2.997 0 0 0-1.01.638a3.049 3.049 0 0 0-.937 2.07a2.974 2.974 0 0 0 .192 1.179q.026.07.057.139t.063.134q.033.066.069.13t.074.125a5.452 5.452 0 0 0-.195.513a5.4 5.4 0 1 0 10.27 0a5.452 5.452 0 0 0-.195-.513"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DaemonSetOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 20V6H1v14a2.001 2.001 0 0 0 2 2h16v-2l-8-.001Z"/><path fill="currentColor" d="M17.293 8.946q.05-.082.095-.17a1.941 1.941 0 0 0 .08-.182a2.021 2.021 0 0 0-.041-1.536a1.977 1.977 0 0 0-1.11-1.055l-.005-.001L16.306 6a2.013 2.013 0 0 1 .178 1.79v.002l-.001.001a3.616 3.616 0 0 0-1.123-.73a3.59 3.59 0 0 0-2.721 0a3.617 3.617 0 0 0-1.123.73v-.001l-.001-.002a1.949 1.949 0 0 1-.11-.932A2.084 2.084 0 0 1 11.712 6l-.006.002l-.006.001a2.018 2.018 0 0 0-1.17 2.59a1.944 1.944 0 0 0 .08.183q.044.088.095.17a3.6 3.6 0 1 0 6.587 0m-1.894-.15a.6.6 0 1 1-.425.175a.598.598 0 0 1 .425-.176m-2.799 0a.6.6 0 1 1-.424.175a.598.598 0 0 1 .424-.176m1.4 3.812a2.172 2.172 0 0 1-1.252-.394a2.22 2.22 0 0 1-.792-1.007h4.086a2.22 2.22 0 0 1-.792 1.007a2.172 2.172 0 0 1-1.25.394"/><path fill="currentColor" d="M22.396 2.604A1.91 1.91 0 0 0 21 2H7a1.91 1.91 0 0 0-1.396.604A1.911 1.911 0 0 0 5 4v12a1.911 1.911 0 0 0 .604 1.396A1.911 1.911 0 0 0 7 18h14a2.067 2.067 0 0 0 2-2V4a1.91 1.91 0 0 0-.604-1.396M21 16H7V4h14Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DataMining(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.403 8.69a7.11 7.11 0 0 0-3.348-4.217l-.869-.496l.496-.868l-1.736-.993l-.496.869L4.977 1l-.992 1.736l3.473 1.985l-.497.868l-.992 1.737L2 14.272l1.736.992l3.97-6.946l.992-1.736l.496-.869l.868.496Z"/><path fill="currentColor" d="M10 10v3h12v-3Zm11 2h-4v-1h4Zm-11 3v3h12v-3Zm11 2h-4v-1h4Zm-11 3v3h12v-3Zm11 2h-4v-1h4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DataScientist(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 5.662v.676c0 1.734 1.365 1.661 1.5 1.661c.146 0 1.5.083 1.5-1.661v-.676c0-1.734-1.365-1.661-1.5-1.661c-.146 0-1.5-.083-1.5 1.661m2.125-.11v.886c0 .703-.219.94-.615.94s-.625-.237-.625-.94v-.885c0-.685.23-.922.615-.922c.396-.01.625.237.625.922M21 12.662v.676c0 1.734 1.365 1.661 1.5 1.661c.146 0 1.5.083 1.5-1.661v-.676c0-1.734-1.365-1.661-1.5-1.661c-.146 0-1.5-.083-1.5 1.661m2.125-.11v.886c0 .703-.219.94-.615.94s-.625-.237-.625-.94v-.885c0-.685.23-.922.615-.922c.396-.01.625.237.625.922M21 19.662v.676c0 1.734 1.365 1.661 1.5 1.661c.146 0 1.5.083 1.5-1.661v-.676c0-1.734-1.365-1.661-1.5-1.661c-.146 0-1.5-.083-1.5 1.661m2.125-.11v.886c0 .703-.219.94-.615.94s-.625-.237-.625-.94v-.885c0-.685.23-.922.615-.922c.396-.01.625.237.625.922M17 9.662v.676c0 1.734 1.365 1.661 1.5 1.661c.146 0 1.5.083 1.5-1.661v-.676c0-1.734-1.365-1.661-1.5-1.661c-.146 0-1.5-.083-1.5 1.661m2.125-.11v.886c0 .703-.219.94-.615.94s-.625-.237-.625-.94v-.885c0-.685.23-.922.615-.922c.396-.01.625.237.625.922M17 1.59v.647l1-.238V5h1V1m-2 14.59v.647l1-.238V19h1.001v-4M16 21v3H0v-3c0-2.66 5.33-4 8-4s8 1.34 8 4"/><circle cx="8" cy="12" r="4" fill="currentColor"/><path fill="currentColor" d="M17.885 23.553c0-.685.23-.922.615-.922c.396-.01.625.237.625.922V24H20v-.338c0-1.734-1.365-1.661-1.5-1.661c-.146 0-1.5-.083-1.5 1.661V24h.885Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DataScientistOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.885 23.553c0-.685.23-.922.615-.922c.396-.01.625.237.625.922V24H20v-.338C20 21.928 18.635 22 18.5 22c-.146 0-1.5-.082-1.5 1.662V24h.885ZM21 5.662v.676C21 8.072 22.365 8 22.5 8c.146 0 1.5.082 1.5-1.662v-.676C24 3.928 22.635 4 22.5 4c-.146 0-1.5-.082-1.5 1.662m2.125-.11v.886c0 .703-.219.94-.615.94s-.625-.237-.625-.94v-.885c0-.685.23-.922.615-.922c.396-.01.625.237.625.922M21 12.662v.676C21 15.072 22.365 15 22.5 15c.146 0 1.5.08 1.5-1.662v-.676C24 10.928 22.635 11 22.5 11c-.146 0-1.5-.08-1.5 1.662m2.125-.11v.886c0 .703-.219.94-.615.94s-.625-.237-.625-.94v-.885c0-.685.23-.922.615-.922c.396-.01.625.237.625.922M21 19.662v.676C21 22.072 22.365 22 22.5 22c.146 0 1.5.082 1.5-1.662v-.676C24 17.928 22.635 18 22.5 18c-.146 0-1.5-.082-1.5 1.662m2.125-.11v.886c0 .703-.219.94-.615.94s-.625-.237-.625-.94v-.885c0-.685.23-.922.615-.922c.396-.01.625.237.625.922M17 9.662v.676C17 12.072 18.365 12 18.5 12c.146 0 1.5.082 1.5-1.662v-.676C20 7.928 18.635 8 18.5 8c-.146 0-1.5-.082-1.5 1.662m2.125-.11v.886c0 .703-.219.94-.615.94s-.625-.237-.625-.94v-.885c0-.685.23-.922.615-.922c.396-.01.625.237.625.922M17 1.59v.647L18 2v3h1V1m-2 14.59v.647L18 16v3h1v-4M8 10a2 2 0 1 1-2 2a2.006 2.006 0 0 1 2-2m0 9c2.7 0 5.8 1.29 6 2v1H2v-1c.2-.72 3.3-2 6-2M8 8a4 4 0 1 0 4 4a3.999 3.999 0 0 0-4-4m0 9c-2.67 0-8 1.34-8 4v3h16v-3c0-2.66-5.33-4-8-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Database(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 16v3c0 1.657-4.03 3-9 3s-9-1.343-9-3v-3c0 1.657 4.03 3 9 3s9-1.343 9-3m-9-1c-4.97 0-9-1.343-9-3v3c0 1.657 4.03 3 9 3s9-1.343 9-3v-3c0 1.657-4.03 3-9 3m0-13C7.03 2 3 3.343 3 5v2c0 1.657 4.03 3 9 3s9-1.343 9-3V5c0-1.657-4.03-3-9-3m0 9c-4.97 0-9-1.343-9-3v3c0 1.657 4.03 3 9 3s9-1.343 9-3V8c0 1.657-4.03 3-9 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabaseMigration(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24 11.021c0 1.013-2.462 1.833-5.5 1.833s-5.5-.82-5.5-1.833V9.188c0 1.012 2.462 1.833 5.5 1.833s5.5-.82 5.5-1.833"/><path fill="currentColor" d="M18.5 6.137c-3.035 0-5.5-.825-5.5-1.833v3.64c0 1.009 2.465 1.834 5.5 1.834S24 8.953 24 7.944v-3.64c0 1.008-2.465 1.833-5.5 1.833"/><path fill="currentColor" d="M24 1.833C24 .821 21.538 0 18.5 0S13 .82 13 1.833v1.223c0 1.012 2.462 1.833 5.5 1.833s5.5-.82 5.5-1.833M11 22.021c0 1.013-2.462 1.833-5.5 1.833S0 23.034 0 22.021v-1.833C0 21.2 2.462 22.02 5.5 22.02s5.5-.82 5.5-1.833Z"/><path fill="currentColor" d="M5.5 17.137c-3.035 0-5.5-.825-5.5-1.833v3.64c0 1.009 2.465 1.834 5.5 1.834s5.5-.825 5.5-1.834v-3.64c0 1.008-2.465 1.833-5.5 1.833"/><path fill="currentColor" d="M11 12.833C11 11.821 8.538 11 5.5 11S0 11.82 0 12.833v1.223c0 1.012 2.462 1.833 5.5 1.833s5.5-.82 5.5-1.833ZM20 21h-2v-2.59L15.41 21L14 19.59L16.59 17H14v-2h6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabaseMigrationOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 21h-2v-2.59L15.41 21L14 19.59L16.59 17H14v-2h6zM5.5 12C2 12 0 12.7 0 14v8c0 1.3 2.83 2 5.5 2s5.5-.7 5.5-2v-8c0-1.3-2-2-5.5-2m0 1c2.485 0 4.5.672 4.5 1.5S7.985 16 5.5 16S1 15.328 1 14.5S3.015 13 5.5 13M1 20.238v-1.056A10.956 10.956 0 0 0 5.5 20a10.957 10.957 0 0 0 4.5-.818v1.089A11.468 11.468 0 0 1 5.5 21a11.64 11.64 0 0 1-4.5-.762M18.5 0C15 0 13 .7 13 2v8c0 1.3 2.83 2 5.5 2s5.5-.7 5.5-2V2c0-1.3-2-2-5.5-2m0 1c2.485 0 4.5.672 4.5 1.5S20.985 4 18.5 4S14 3.328 14 2.5S16.015 1 18.5 1M14 8.238V7.182A10.956 10.956 0 0 0 18.5 8a10.957 10.957 0 0 0 4.5-.818v1.089A11.468 11.468 0 0 1 18.5 9a11.64 11.64 0 0 1-4.5-.762"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabaseOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.5C7.395 1.5 2.5 2.727 2.5 5v14c0 2.273 4.895 3.5 9.5 3.5s9.5-1.227 9.5-3.5V5c0-2.273-4.895-3.5-9.5-3.5m0 16c-5.267 0-8.5-1.456-8.5-2.5v-1.367C5.166 14.854 8.66 15.5 12 15.5s6.834-.646 8.5-1.867V15c0 1.044-3.233 2.5-8.5 2.5m0-4c-5.267 0-8.5-1.456-8.5-2.5V9.632C5.166 10.854 8.66 11.5 12 11.5s6.834-.646 8.5-1.868V11c0 1.044-3.233 2.5-8.5 2.5m8.5 5.5c0 1.044-3.233 2.5-8.5 2.5S3.5 20.044 3.5 19v-1.367C5.166 18.854 8.66 19.5 12 19.5s6.834-.646 8.5-1.867ZM12 9.5C6.733 9.5 3.5 8.044 3.5 7V5c0-1.044 3.233-2.5 8.5-2.5s8.5 1.456 8.5 2.5v2c0 1.044-3.233 2.5-8.5 2.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Deploy(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24.01 20.027v2h-24v-2h4v-1a2.006 2.006 0 0 1-2-2v-10a2.006 2.006 0 0 1 2-2h1.996v2H4.01v10h16v-10h-2.004v-2h2.004a2.006 2.006 0 0 1 2 2l-.01 10a1.997 1.997 0 0 1-1.99 2v1Zm-9-6.012l-3-3l-3 3h2v2.01h2v-2.01Zm.995-7.991a4 4 0 1 1-4-4a4.001 4.001 0 0 1 4 4m-4.4 2.96v-.56a.802.802 0 0 1-.8-.8v-.4L9.06 5.479a2.958 2.958 0 0 0 2.545 3.505m2.658-1.007a2.977 2.977 0 0 0-1.068-4.704a.797.797 0 0 1-.79.75h-.8v.8a.401.401 0 0 1-.4.4h-.8v.8h2.4a.401.401 0 0 1 .4.4v1.2h.4a.787.787 0 0 1 .658.354"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiffModified(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.2 3H4.8A1.8 1.8 0 0 0 3 4.8v14.4A1.8 1.8 0 0 0 4.8 21h14.4a1.8 1.8 0 0 0 1.8-1.8V4.8A1.8 1.8 0 0 0 19.2 3M12 16a4 4 0 1 1 4-4a4 4 0 0 1-4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiffModifiedOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.2 3H4.8A1.8 1.8 0 0 0 3 4.8v14.4A1.8 1.8 0 0 0 4.8 21h14.4a1.8 1.8 0 0 0 1.8-1.8V4.8A1.8 1.8 0 0 0 19.2 3M19 19H5V5h14Z"/><circle cx="12" cy="12" r="4" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dns(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 17.5h2v3h-2z"/><path fill="currentColor" d="M20.998 14H3.002A2.002 2.002 0 0 0 1 16.002v5.996A2.002 2.002 0 0 0 3.002 24h17.996A2.002 2.002 0 0 0 23 21.998v-5.996A2.002 2.002 0 0 0 20.998 14M8 20.5A1.473 1.473 0 0 1 6.5 22H3v-6h3.5A1.473 1.473 0 0 1 8 17.5Zm6.5 1.5h-1.2l-2.55-3.5V22H9.5v-6h1.25l2.5 3.5V16h1.25Zm6.5-4.48h-3.5v.74H20a1 1 0 0 1 1 1V21a1 1 0 0 1-1 1h-4v-1.5h3.51v-.74H17a1 1 0 0 1-1-1V17a1 1 0 0 1 1-1h4ZM4.26 12A8.243 8.243 0 0 1 4 10a8.243 8.243 0 0 1 .26-2h3.38a16.513 16.513 0 0 0-.14 2a16.514 16.514 0 0 0 .14 2h2.02a14.71 14.71 0 0 1-.16-2a14.581 14.581 0 0 1 .16-2h4.68a14.59 14.59 0 0 1 .16 2a14.72 14.72 0 0 1-.16 2h2.02a16.512 16.512 0 0 0 .14-2a16.511 16.511 0 0 0-.14-2h3.38a8.24 8.24 0 0 1 .26 2a8.24 8.24 0 0 1-.26 2h2.059A10 10 0 1 0 2.2 12Zm14.66-6h-2.95a15.65 15.65 0 0 0-1.38-3.56A8.03 8.03 0 0 1 18.92 6M12 2.04A14.086 14.086 0 0 1 13.91 6h-3.82A14.086 14.086 0 0 1 12 2.04m-2.59.4A15.648 15.648 0 0 0 8.03 6H5.08a7.987 7.987 0 0 1 4.33-3.56"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Drone(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24 8V7h-7v1h3v1h-1v2h-3.3l-1.501-2.6H9.855l-1.5 2.6H5V9H4V8h3V7H0v1h3v1H2v5h3v-1h3.167l.216.374A8.002 8.002 0 0 0 4 20.5a.5.5 0 0 0 1 0a7.001 7.001 0 0 1 3.883-6.259l.972 1.683h4.344l.96-1.663A7.002 7.002 0 0 1 19 20.5a.5.5 0 0 0 1 0a8.003 8.003 0 0 0-4.34-7.106l.227-.394H19v1h3V9h-1V8Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EdgeComputing(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.82 10.99L7 9.17V6H6V2h3v4H8v2.76l1.72 1.72A5.61 5.61 0 0 1 12 10a5.551 5.551 0 0 1 1 .09V5.91a1.5 1.5 0 1 1 1 0v4.46A5.447 5.447 0 0 1 16.29 12h2.8a1.5 1.5 0 1 1 0 1h-2.12a5.358 5.358 0 0 1 .54 1.53a3.74 3.74 0 0 1-.26 7.47H7.5a4.499 4.499 0 0 1-2.47-8.26l-1.89-1.89a1.518 1.518 0 1 1 .71-.71l2.13 2.13a4.419 4.419 0 0 1 1.03-.24a5.562 5.562 0 0 1 1.81-2.04"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EdgeComputingOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.09 13a1.5 1.5 0 1 0 0-1h-2.8A5.448 5.448 0 0 0 14 10.37V5.91a1.5 1.5 0 1 0-1 0v4.18a5.551 5.551 0 0 0-1-.09a5.61 5.61 0 0 0-2.28.48L8 8.76V6h1V2H6v4h1v3.17l1.82 1.82a5.562 5.562 0 0 0-1.81 2.04a4.418 4.418 0 0 0-1.03.24l-2.13-2.13a1.519 1.519 0 1 0-.71.71l1.89 1.89A4.499 4.499 0 0 0 7.5 22h9.75a3.74 3.74 0 0 0 .26-7.47a5.358 5.358 0 0 0-.54-1.53Zm-1.72 3.53A1.736 1.736 0 0 1 19 18.25A1.758 1.758 0 0 1 17.25 20H7.5a2.498 2.498 0 0 1-.28-4.98l1.07-.11l.5-.96a3.618 3.618 0 0 1 6.76.97l.3 1.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Endpoints(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.02 10.98h2v2h-2zm3.97 4.03h2v2h-2zm0-4.03h2v2h-2zm4.01 0h2v2h-2zm4.01 0h2v2h-2zm4.01 0h2v2h-2zm-4.01-3.97h2v2h-2zM5.99 17.99h4v4h-4zm8.02-15.98h4v4h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EndpointsConnected(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.02 13.01v-2h-4.01v-5h1v-4h-4v4h1v5h-12v2H7V18H6v4h3.99v-4h-1v-4.99z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EndpointsDisconnected(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m2.39 4.93l6.08 6.08H3.01v2H7V18H6v4h3.99v-4h-1v-4.99h1.48l8.65 8.65l1.28-1.28L3.67 3.66zm14.62 6.08v-5h1v-4h-4v4h1v5h-1.44l2 2h5.44v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Enhancement(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.02 13.99h4v8.02h-4zm7-3.97h4v11.99h-4zM22 6l-4-4l-4 4h2v16.02h4V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Env(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 18a1 1 0 0 1-1 1h-4a3 3 0 0 0-3 3a3 3 0 0 0-3-3H5a1 1 0 0 1-1-1H2a3 3 0 0 0 3 3h4a2 2 0 0 1 2 2h2a2 2 0 0 1 2-2h4a3 3 0 0 0 3-3Zm0-12a1 1 0 0 0-1-1h-4a3 3 0 0 1-3-3a3 3 0 0 1-3 3H5a1 1 0 0 0-1 1H2a3 3 0 0 1 3-3h4a2 2 0 0 0 2-2h2a2 2 0 0 0 2 2h4a3 3 0 0 1 3 3Zm-8 6L9 8H7v8h2v-4l3 4h2V8h-2zm9-4l-2 5.27L17 8h-2l3 8h2l3-8zM1 8v8h5v-2H3v-1h2v-2H3v-1h3V8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileSystem(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 13h6a1 1 0 0 0 1-1V9a1 1 0 0 0-1-1h-3l-1-1h-2a1 1 0 0 0-1 1v1H8V7h2a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1H7L6 1H4a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h2v13h7v1a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1v-3a1 1 0 0 0-1-1h-3l-1-1h-2a1 1 0 0 0-1 1v1H8v-7h5v1a1 1 0 0 0 1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileSystemOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 8h-2.59l-1-1H15a2.006 2.006 0 0 0-2 2v4a2.006 2.006 0 0 0 2 2h6a2.006 2.006 0 0 0 2-2v-3a2.006 2.006 0 0 0-2-2m0 5h-6V9h2l1 1h3Zm0 4h-2.59l-1-1H15a2.006 2.006 0 0 0-2 2v4a2.006 2.006 0 0 0 2 2h6a2.006 2.006 0 0 0 2-2v-3a2.006 2.006 0 0 0-2-2m0 5h-6v-4h2l1 1h3ZM8 19v-7h5v-2H8V8H6v13h7v-2zm2-18H7.41l-1-1H4a2.006 2.006 0 0 0-2 2v4a2.006 2.006 0 0 0 2 2h6a2.006 2.006 0 0 0 2-2V3a2.006 2.006 0 0 0-2-2m0 5H4V2h2l1 1h3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flask(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.642 18.452L19.731 17h-.001L15 9.462V3h2V1H7v2h2v6.462L4.27 17l-.912 1.452A2.317 2.317 0 0 0 5.321 22h13.358a2.317 2.317 0 0 0 1.963-3.548M13 3v7l4.374 7H6.626L11 10V3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fork(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 5a3 3 0 1 0-4 2.816V11H6V7.816a3 3 0 1 0-2 0V11a2 2 0 0 0 2 2h5v4.184a3 3 0 1 0 2 0V13h5a2 2 0 0 0 2-2V7.816A2.991 2.991 0 0 0 22 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ForkOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 2a2.993 2.993 0 0 0-1 5.816V11H6V7.816a3 3 0 1 0-2 0V11a2 2 0 0 0 2 2h5v4.184a3 3 0 1 0 2 0V13h5a2 2 0 0 0 2-2V7.816A2.993 2.993 0 0 0 19 2M5 6a1 1 0 1 1 1-1a1 1 0 0 1-1 1m7 15a1 1 0 1 1 1-1a1 1 0 0 1-1 1m7-15a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Genomic(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.368 16a21.012 21.012 0 0 0-2.751-3.992S10.12 9.153 10.003 9h2.448c-.163.232-.335.45-.513.677l1.326 1.53A23.848 23.848 0 0 0 15.454 8a11.23 11.23 0 0 0 1.534-5c.023-.507.007-1.038-.002-1.575A.958.958 0 0 0 15.993.5a1 1 0 0 0-.997 1.02c.023.494-.006.977-.005 1.48l-6.992.002L8 1.505A.989.989 0 0 0 7 .5a.948.948 0 0 0-.995.955l-.001 1.359a11.615 11.615 0 0 0 2.853 7.918S12.182 14.67 12.415 15h-2.35c.046-.058.302-.366.569-.684l-1.335-1.54l-.442.509A10.421 10.421 0 0 0 7.164 16a11.667 11.667 0 0 0-1.16 5.186l.001 1.359A.948.948 0 0 0 7 23.5a.989.989 0 0 0 1-1.005s0-1.45 0-1.497l6.99.002c-.001.503.028.986.005 1.48a1 1 0 0 0 .997 1.02a.958.958 0 0 0 .993-.925c.009-.537.025-1.068.002-1.575a11.361 11.361 0 0 0-1.62-5m-.915-11a12.711 12.711 0 0 1-.782 2H8.81a9.917 9.917 0 0 1-.613-2Zm-3.169 8.54h.002l.03.037ZM8.198 19a9.917 9.917 0 0 1 .614-2h4.859a12.711 12.711 0 0 1 .782 2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HardwareCircuit(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 0H2a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h20a2 2 0 0 0 2-2V2a2 2 0 0 0-2-2M6 16a1 1 0 1 1 1-1a1.001 1.001 0 0 1-1 1M22 6h-2.184a3 3 0 1 0 0 2H22v4h-4v2h4v2h-2v2h2v4h-8v-1.184a3 3 0 1 0-2 0V22H7v-4.184a3 3 0 1 0-2 0V22H2V2h4v6h2V2h2v10h2V2h10Zm-4 1a1 1 0 1 1-1-1a1.001 1.001 0 0 1 1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hourglass(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g><path fill="currentColor" d="M7 3H17V7.2L12 12L7 7.2V3Z"><animate id="eosIconsHourglass0" fill="freeze" attributeName="opacity" begin="0;eosIconsHourglass1.end" dur="2s" from="1" to="0"/></path><path fill="currentColor" d="M17 21H7V16.8L12 12L17 16.8V21Z"><animate fill="freeze" attributeName="opacity" begin="0;eosIconsHourglass1.end" dur="2s" from="0" to="1"/></path><path fill="currentColor" d="M6 2V8H6.01L6 8.01L10 12L6 16L6.01 16.01H6V22H18V16.01H17.99L18 16L14 12L18 8.01L17.99 8H18V2H6ZM16 16.5V20H8V16.5L12 12.5L16 16.5ZM12 11.5L8 7.5V4H16V7.5L12 11.5Z"/><animateTransform id="eosIconsHourglass1" attributeName="transform" attributeType="XML" begin="eosIconsHourglass0.end" dur="0.5s" from="0 12 12" to="180 12 12" type="rotate"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Inbound(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 2H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h8v-2H4V6l8 5l8-5v8h2V4a2 2 0 0 0-2-2m-8 7L4 4h16Zm6 4v3h4v2h-4v3l-4-4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Infinity(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.385 8.609A5.469 5.469 0 0 0 18.5 7a5.568 5.568 0 0 0-3.928 1.609L12 11.092L9.428 8.61a5.5 5.5 0 1 0 0 7.782L12 13.83l2.572 2.562a5.514 5.514 0 1 0 7.813-7.782m-14.37 6.374a3.541 3.541 0 0 1-4.986 0a3.518 3.518 0 0 1 4.985-4.965l2.486 2.474Zm12.956 0a3.539 3.539 0 0 1-4.985 0l-2.486-2.49l2.486-2.476a3.541 3.541 0 0 1 4.985 0a3.506 3.506 0 0 1 0 4.965"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ingress(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 24h-4v-4h4zm2-13.982h-3v7.996h-2v-7.996H8l1-1.001l1-1.001l1-1.001l1-1.001l1 1l1 1.002l1 1ZM1 15.999V3.001a1.983 1.983 0 0 1 .158-.777a2.02 2.02 0 0 1 1.065-1.066A1.981 1.981 0 0 1 3 1h18a1.981 1.981 0 0 1 .777.158a2.02 2.02 0 0 1 1.065 1.066A1.983 1.983 0 0 1 23 3v12.998a1.983 1.983 0 0 1-.158.777a2.02 2.02 0 0 1-1.065 1.066A1.981 1.981 0 0 1 21 18h-6v-1.99h6v-13H3v13h6V18H3a1.981 1.981 0 0 1-.777-.158a2.02 2.02 0 0 1-1.065-1.066A1.983 1.983 0 0 1 1 16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InitContainer(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m23.915 20.318l-1.07-.83a3.882 3.882 0 0 0 .04-.49a2.626 2.626 0 0 0-.04-.49l1.06-.83a.26.26 0 0 0 .06-.32l-1-1.729a.253.253 0 0 0-.31-.11l-1.24.5a3.418 3.418 0 0 0-.849-.49l-.19-1.32a.239.239 0 0 0-.24-.21h-1.998a.257.257 0 0 0-.25.21l-.19 1.32a3.993 3.993 0 0 0-.85.49l-1.24-.5a.262.262 0 0 0-.309.11l-1 1.73a.248.248 0 0 0 .06.32l1.06.83a4.012 4.012 0 0 0 0 .98l-1.06.83a.26.26 0 0 0-.06.319l1 1.73a.253.253 0 0 0 .31.11l1.24-.5a3.418 3.418 0 0 0 .849.49l.19 1.32a.249.249 0 0 0 .25.21h1.999a.257.257 0 0 0 .25-.21l.19-1.32a3.693 3.693 0 0 0 .839-.49l1.25.5a.262.262 0 0 0 .31-.11l.999-1.73a.26.26 0 0 0-.06-.32m-4.778.18a1.5 1.5 0 1 1 1.5-1.5a1.497 1.497 0 0 1-1.5 1.5"/><path fill="currentColor" d="M10 3a1.96 1.96 0 0 0 .796 1.561L5.5 7.987H3.997A1.997 1.997 0 0 0 2 9.983v8.02A1.997 1.997 0 0 0 3.997 20h8.083a7.014 7.014 0 0 1-.08-1a8.397 8.397 0 0 1 1-4v-3a1 1 0 0 1 2 0v1.252a6.838 6.838 0 0 1 2-.965v-.285A1.002 1.002 0 0 1 18 11a.98.98 0 0 1 1 1a7.262 7.262 0 0 1 3 .685V9.983a1.997 1.997 0 0 0-1.997-1.996h-1.491l-5.302-3.43A1.96 1.96 0 0 0 14 3a2.014 2.014 0 0 0-.986-1.719V0H12v2.006A.992.992 0 1 1 11 3M7 16a1 1 0 0 1-2 0v-3.998a1 1 0 0 1 2 0Zm4 0a1 1 0 0 1-2 0v-4a1 1 0 0 1 2 0Zm1.017-11.025l4.655 3.012h-9.33l4.653-3.011"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InitContainerOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 18H4v-8h16v2.078a6.94 6.94 0 0 1 2 .603V10a2.002 2.002 0 0 0-2-2l-1.488-.015l-5.302-3.43A1.957 1.957 0 0 0 14 3a2.03 2.03 0 0 0-1-1.721V0h-1v2a1 1 0 1 1-1 1h-1a1.957 1.957 0 0 0 .796 1.56L5.5 7.984L4 8a2.002 2.002 0 0 0-2 2v8a2.002 2.002 0 0 0 2 2h8.08a7.013 7.013 0 0 1-.08-1.003A7.063 7.063 0 0 1 12 18m-.005-13.026L12 5c.006 0 .01-.027.017-.027L17 8H7Z"/><path fill="currentColor" d="M15 13.26V12a1 1 0 0 0-2 0v3.408a7.027 7.027 0 0 1 2-2.148M18 11a1 1 0 0 0-1 1v.294A6.994 6.994 0 0 1 19 12a1 1 0 0 0-1-1M6 11a1 1 0 0 0-1 1v4a1 1 0 0 0 2 0v-4a1 1 0 0 0-1-1m4 0a1 1 0 0 0-1 1v4a1 1 0 0 0 2 0v-4a1 1 0 0 0-1-1m13.879 9.319l-1.07-.83a3.885 3.885 0 0 0 .04-.49a2.626 2.626 0 0 0-.04-.49l1.06-.83a.26.26 0 0 0 .06-.32l-1-1.73a.253.253 0 0 0-.31-.11l-1.239.5a3.418 3.418 0 0 0-.85-.49l-.19-1.319a.239.239 0 0 0-.24-.21h-1.998a.257.257 0 0 0-.25.21l-.19 1.32a3.993 3.993 0 0 0-.85.49l-1.239-.5a.262.262 0 0 0-.31.11l-1 1.73a.248.248 0 0 0 .06.32l1.06.829a4.013 4.013 0 0 0 0 .98l-1.06.83a.26.26 0 0 0-.06.32l1 1.73a.253.253 0 0 0 .31.11l1.24-.5a3.418 3.418 0 0 0 .849.49l.19 1.319a.249.249 0 0 0 .25.21H20.1a.257.257 0 0 0 .25-.21l.19-1.32a3.692 3.692 0 0 0 .839-.49l1.25.5a.262.262 0 0 0 .31-.11l.999-1.73a.26.26 0 0 0-.06-.32m-4.758.18a1.5 1.5 0 1 1 1.5-1.5a1.497 1.497 0 0 1-1.5 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Installing(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 16h4v4H4V16z" class="st0"><animate fill="remove" accumulate="none" additive="replace" attributeName="opacity" begin=".2" calcMode="linear" dur="3s" keyTimes="0;0.9;1" repeatCount="indefinite" restart="always" values="1;0;0"/></path><path fill="currentColor" d="M10 16h4v4h-4V16z" class="st0"><animate fill="remove" accumulate="none" additive="replace" attributeName="opacity" begin=".4" calcMode="linear" dur="3s" keyTimes="0;0.9;1" repeatCount="indefinite" restart="always" values="1;0;0"/></path><path fill="currentColor" d="M16 16h4v4h-4V16z" class="st0"><animate fill="remove" accumulate="none" additive="replace" attributeName="opacity" begin=".6" calcMode="linear" dur="3s" keyTimes="0;0.9;1" repeatCount="indefinite" restart="always" values="1;0;0"/></path><path fill="currentColor" d="M4 10h4v4H4V10z" class="st0"><animate fill="remove" accumulate="none" additive="replace" attributeName="opacity" begin=".8" calcMode="linear" dur="3s" keyTimes="0;0.9;1" repeatCount="indefinite" restart="always" values="1;0;0"/></path><path fill="currentColor" d="M10 10h4v4h-4V10z" class="st0"><animate fill="remove" accumulate="none" additive="replace" attributeName="opacity" begin="1" calcMode="linear" dur="3s" keyTimes="0;0.9;1" repeatCount="indefinite" restart="always" values="1;0;0"/></path><path fill="currentColor" d="M16 10h4v4h-4V10z" class="st0"><animate fill="remove" accumulate="none" additive="replace" attributeName="opacity" begin="1.2" calcMode="linear" dur="3s" keyTimes="0;0.9;1" repeatCount="indefinite" restart="always" values="1;0;0"/></path><path fill="currentColor" d="M4 4h4v4H4V4z" class="st0"><animate fill="remove" accumulate="none" additive="replace" attributeName="opacity" begin="1.4" calcMode="linear" dur="3s" keyTimes="0;0.9;1" repeatCount="indefinite" restart="always" values="1;0;0"/></path><path fill="currentColor" d="M10 4h4v4h-4V4z" class="st0"><animate fill="remove" accumulate="none" additive="replace" attributeName="opacity" begin="1.6" calcMode="linear" dur="3s" keyTimes="0;0.9;1" repeatCount="indefinite" restart="always" values="1;0;0"/></path><path fill="currentColor" d="M16 4h4v4h-4V4z" class="st0"><animate fill="remove" accumulate="none" additive="replace" attributeName="opacity" begin="1.8" calcMode="linear" dur="3s" keyTimes="0;0.9;1" repeatCount="indefinite" restart="always" values="1;0;0"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Iot(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 13v9H3v-9Zm18 0v2h-2v7h-2v-7h-2v-2Zm-11-2a7.537 7.537 0 0 1 3.96 1.149l1.447-1.45A9.522 9.522 0 0 0 12 9a9.363 9.363 0 0 0-5.333 1.68l1.449 1.453A7.36 7.36 0 0 1 12 11"/><path fill="currentColor" d="M12 7a11.494 11.494 0 0 1 6.834 2.27l1.427-1.43A13.48 13.48 0 0 0 12 5a13.333 13.333 0 0 0-8.186 2.822l1.426 1.43A11.343 11.343 0 0 1 12 7"/><path fill="currentColor" d="M12 3a15.471 15.471 0 0 1 9.687 3.41l1.427-1.429A17.43 17.43 0 0 0 .96 4.964l1.427 1.429A15.328 15.328 0 0 1 12 3m0 10a4.5 4.5 0 1 0 4.5 4.5A4.5 4.5 0 0 0 12 13m0 7a2.5 2.5 0 1 1 2.5-2.5A2.5 2.5 0 0 1 12 20"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ip(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.26 12A8.243 8.243 0 0 1 4 10a8.243 8.243 0 0 1 .26-2h3.38a16.513 16.513 0 0 0-.14 2a16.514 16.514 0 0 0 .14 2h2.02a14.71 14.71 0 0 1-.16-2a14.581 14.581 0 0 1 .16-2h4.68a14.59 14.59 0 0 1 .16 2a14.72 14.72 0 0 1-.16 2h2.02a16.512 16.512 0 0 0 .14-2a16.511 16.511 0 0 0-.14-2h3.38a8.24 8.24 0 0 1 .26 2a8.24 8.24 0 0 1-.26 2h2.059A10 10 0 1 0 2.2 12Zm14.66-6h-2.95a15.65 15.65 0 0 0-1.38-3.56A8.03 8.03 0 0 1 18.92 6M12 2.04A14.086 14.086 0 0 1 13.91 6h-3.82A14.086 14.086 0 0 1 12 2.04m-2.59.4A15.648 15.648 0 0 0 8.03 6H5.08a7.987 7.987 0 0 1 4.33-3.56m3.339 15.055h2v1h-2z"/><path fill="currentColor" d="M20.998 14H3.002A2.002 2.002 0 0 0 1 16.002v5.996A2.002 2.002 0 0 0 3.002 24h17.996A2.002 2.002 0 0 0 23 21.998v-5.996A2.002 2.002 0 0 0 20.998 14M9.251 22.005h-1.5v-6h1.5Zm6.998-3.51a1.473 1.473 0 0 1-1.5 1.5h-2v2h-1.5v-6h3.5a1.473 1.473 0 0 1 1.5 1.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IpOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.26 12A8.243 8.243 0 0 1 4 10a8.243 8.243 0 0 1 .26-2h3.38a16.513 16.513 0 0 0-.14 2a16.514 16.514 0 0 0 .14 2h2.02a14.71 14.71 0 0 1-.16-2a14.581 14.581 0 0 1 .16-2h4.68a14.59 14.59 0 0 1 .16 2a14.72 14.72 0 0 1-.16 2h2.02a16.512 16.512 0 0 0 .14-2a16.511 16.511 0 0 0-.14-2h3.38a8.24 8.24 0 0 1 .26 2a8.24 8.24 0 0 1-.26 2L22 13v-3a10 10 0 1 0-20 0v3Zm14.66-6h-2.95a15.65 15.65 0 0 0-1.38-3.56A8.03 8.03 0 0 1 18.92 6M12 2.04A14.086 14.086 0 0 1 13.91 6h-3.82A14.086 14.086 0 0 1 12 2.04m-2.59.4A15.648 15.648 0 0 0 8.03 6H5.08a7.987 7.987 0 0 1 4.33-3.56"/><path fill="currentColor" d="M21 14v8H3v-8zm0-2H3a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h18a2 2 0 0 0 2-2v-8a2 2 0 0 0-2-2"/><path fill="currentColor" d="M7 15h2v6H7zm9 0h-4v6h2v-2h2a1.473 1.473 0 0 0 1.5-1.5v-1A1.473 1.473 0 0 0 16 15m0 3h-2v-2h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Job(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 16.28V11l3-3l-4-4l-3 3H8V6h3V2H3v4h3v6H3v4h3v5h4.28A2 2 0 0 0 12 22a2 2 0 1 0-1.72-3H8v-3h3v-4H8V9h7l2 2v5.27a2 2 0 1 0 2 0ZM4 5V3h6v2Zm8 14a1 1 0 1 1-1 1a1 1 0 0 1 1-1m-2-6v2H4v-2Zm8-7.59L20.59 8L18 10.59L15.41 8ZM18 19a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Kubelet(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.803 15H0l1.25-2.75L2.5 9.5l1.25-2.75L5 4h2.803l-1.25 2.75l-1.25 2.75l-1.25 2.75ZM8 1.992a2.61 2.61 0 0 1 .05-.51A3.101 3.101 0 0 1 8.183 1H5.998a.998.998 0 0 0-.92.606a.984.984 0 0 0-.078.386V3h3zm-3 15.01H3.5L2.75 17H2v.002A1.998 1.998 0 0 0 4 19h1.555a4.286 4.286 0 0 1-.231-.458a4.03 4.03 0 0 1-.175-.488a3.848 3.848 0 0 1-.11-.514A3.775 3.775 0 0 1 5 17.003m11-15.01a.986.986 0 0 0-.292-.702a.998.998 0 0 0-.706-.29h-4.004a.998.998 0 0 0-.92.606a.984.984 0 0 0-.078.386V3h6zM15.302 16H7v1.003A1.998 1.998 0 0 0 9 19h3.284l.754-.75l.755-.75l.755-.75Zm.218-1H5l1.25-2.75L7.5 9.5l1.25-2.75L10 4h5.993l.726 1.593l.725 1.594l.725 1.593l.725 1.594q-.254.1-.498.223t-.475.27a5.252 5.252 0 0 0-.445.319a5.052 5.052 0 0 0-.412.366A5.254 5.254 0 0 0 15.52 15m-.51-5.99a1.994 1.994 0 1 0-.157.778a1.994 1.994 0 0 0 .157-.778m0 13.846l1.142-1.136l1.142-1.136l1.143-1.135l1.143-1.136a3.198 3.198 0 0 0 1.808.15a3.247 3.247 0 0 0 .876-.319a3.36 3.36 0 0 0 .782-.58a3.212 3.212 0 0 0 .603-.83a3.26 3.26 0 0 0 .307-.942a3.334 3.334 0 0 0 .015-.98a3.286 3.286 0 0 0-.272-.943l-.54.537l-.54.537l-.54.536l-.54.537l-.377-.374l-.377-.375l-.377-.374l-.376-.375l.54-.549l.54-.549l.54-.55l.54-.548a3.077 3.077 0 0 0-.948-.292a3.214 3.214 0 0 0-.986.017a3.31 3.31 0 0 0-.948.317a3.377 3.377 0 0 0-1.419 1.384a3.202 3.202 0 0 0-.32.87a3.143 3.143 0 0 0 .15 1.798L16.58 17.55l-1.143 1.136l-1.143 1.136l-1.142 1.136a.519.519 0 0 0-.113.165a.469.469 0 0 0 0 .368a.519.519 0 0 0 .113.166l.288.287l.29.287l.288.287l.289.287a.526.526 0 0 0 .173.141a.49.49 0 0 0 .197.053a.43.43 0 0 0 .188-.032a.351.351 0 0 0 .145-.112"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KubeletOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23.97 14.812a3.286 3.286 0 0 0-.271-.943l-2.16 2.147l-1.507-1.498l2.16-2.196a3.293 3.293 0 0 0-4.47 4.094l-3.429 3.407c-.77.872-2.037 1.361-.565 2.409c1.132 1.504 1.47.306 2.424-.512l3.428-3.407a3.29 3.29 0 0 0 4.39-3.501M16 2a1.004 1.004 0 0 0-1-1h-4a1.004 1.004 0 0 0-1 1v1h6Z"/><circle cx="13" cy="9" r="1.5" fill="currentColor"/><path fill="currentColor" d="M9 1H7a1.004 1.004 0 0 0-1 1v1h2V2c0-.83.449-1 1-1M6 4L1 15h2L8 4zM5 17.003V16H3v1.003A1.998 1.998 0 0 0 5 19h2a1.773 1.773 0 0 1-2-1.997M16.06 13H8.1l3.18-7h3.42l2.5 5.467a5.583 5.583 0 0 1 1.694-1.092L15.99 4H10L5 15h10.496a5.018 5.018 0 0 1 .563-2m7.911 1.812a3.286 3.286 0 0 0-.271-.943l-2.16 2.147l-1.507-1.498l2.16-2.196a3.293 3.293 0 0 0-4.47 4.094l-3.429 3.407c-.77.872-2.037 1.361-.565 2.409c1.132 1.504 1.47.306 2.424-.512l3.428-3.407a3.29 3.29 0 0 0 4.39-3.501M14.297 17H9v-1H7v1.003A1.998 1.998 0 0 0 9 19h3.264c.053-.05 2.033-2 2.033-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Loading(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2A10 10 0 1 0 22 12A10 10 0 0 0 12 2Zm0 18a8 8 0 1 1 8-8A8 8 0 0 1 12 20Z" opacity=".5"/><path fill="currentColor" d="M20 12h2A10 10 0 0 0 12 2V4A8 8 0 0 1 20 12Z"><animateTransform attributeName="transform" dur="1s" from="0 12 12" repeatCount="indefinite" to="360 12 12" type="rotate"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockedEnv(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 8h-1V6A5 5 0 0 0 6 6v2H5a2.006 2.006 0 0 0-2 2v10a2.006 2.006 0 0 0 2 2h12a2.006 2.006 0 0 0 2-2V10a2.006 2.006 0 0 0-2-2m-9.72 6.59v.752H5.975v.877H7.45V17H5v-4h2.45v.78H5.975v.81ZM7.9 6a3.1 3.1 0 1 1 6.2 0v2H7.9Zm4.268 11h-.974l-1.63-2.467V17H8.59v-4h.974l1.63 2.479V13h.974Zm3.433 0H14.4L13 13h1l1 3l1-3h1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockedEnvOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.829 14.59v.752H6.524v.877H8V17H5.55v-4H8v.781H6.524v.809z"/><path fill="currentColor" d="M17 8h-1V6A5 5 0 0 0 6 6v2H5a2.006 2.006 0 0 0-2 2v10a2.006 2.006 0 0 0 2 2h12a2.006 2.006 0 0 0 2-2V10a2.006 2.006 0 0 0-2-2M7.9 6a3.1 3.1 0 1 1 6.2 0v2H7.9ZM17 20H5V10h12Z"/><path fill="currentColor" d="M12.168 17h-.974l-1.63-2.467V17H8.59v-4h.974l1.63 2.479V13h.974zm2.933 0h-1.202L12.5 13h1l1 3l1-3h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MachineLearning(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 8h2v1h-2zm0-4h2v1h-2zm0 6h2v1h-2z"/><path fill="currentColor" d="M21 12V9a13.124 13.124 0 0 0-8.354 3h-1.292A13.124 13.124 0 0 0 3 9v3a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1v4a13.153 13.153 0 0 1 9 3.55A13.2 13.2 0 0 1 21 20v-4a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1"/><circle cx="9" cy="4" r="1" fill="currentColor"/><circle cx="15" cy="4" r="1" fill="currentColor"/><path fill="currentColor" d="M16 8H8a3.003 3.003 0 0 1-3-3V3a3.003 3.003 0 0 1 3-3h8a3.003 3.003 0 0 1 3 3v2a3.003 3.003 0 0 1-3 3M8 2a1.001 1.001 0 0 0-1 1v2a1.001 1.001 0 0 0 1 1h8a1.001 1.001 0 0 0 1-1V3a1.001 1.001 0 0 0-1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MachineLearningOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 10a16.84 16.84 0 0 0-9 3.244A16.84 16.84 0 0 0 3 10v2.96a2.004 2.004 0 0 0-2 2.007v1.004c0 1.109 2 2.208 2 2.208v2.007a14.868 14.868 0 0 1 7.417 2.55A15.09 15.09 0 0 1 12 24a15.09 15.09 0 0 1 1.583-1.264A14.868 14.868 0 0 1 21 20.186v-2.208a2.004 2.004 0 0 0 2-2.007v-1.004a2.004 2.004 0 0 0-2-2.007Zm-9 11.422a16.841 16.841 0 0 0-7-2.996v-6.15a14.8 14.8 0 0 1 5.417 2.282A15.09 15.09 0 0 1 12 15.822a15.09 15.09 0 0 1 1.583-1.264A14.8 14.8 0 0 1 19 12.275v6.151a16.841 16.841 0 0 0-7 2.996M11 8h2v1h-2zm0-4h2v1h-2z"/><path fill="currentColor" d="M11 10h2v1h-2zM9 5a1 1 0 0 0 1-1a.983.983 0 0 0-.99-.99A.995.995 0 1 0 9 5"/><circle cx="15" cy="4" r="1" fill="currentColor"/><path fill="currentColor" d="M16 8H8a3.003 3.003 0 0 1-3-3V3a3.003 3.003 0 0 1 3-3h8a3.003 3.003 0 0 1 3 3v2a3.003 3.003 0 0 1-3 3M8 2a1.001 1.001 0 0 0-1 1v2a1.001 1.001 0 0 0 1 1h8a1.001 1.001 0 0 0 1-1V3a1.001 1.001 0 0 0-1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Master(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 20h16v2H4zM21 5a1 1 0 0 0-1 1a.899.899 0 0 0 .064.332a1.104 1.104 0 0 0 .167.283l-2.12 1.692L15.994 10L14.24 6.927l-1.752-3.072a1.045 1.045 0 0 0 .369-.353A.946.946 0 0 0 13 3a1 1 0 1 0-2 0a.946.946 0 0 0 .143.502a1.046 1.046 0 0 0 .369.353L9.757 6.927L8.002 10L5.886 8.308L3.769 6.615a1.104 1.104 0 0 0 .167-.283A.899.899 0 0 0 4 6a1 1 0 1 0-1 1a.147.147 0 0 0 .041-.007a.212.212 0 0 1 .041-.01l.459 5.509L4 18h16l.459-5.508l.459-5.509a.212.212 0 0 1 .04.01A.147.147 0 0 0 21 7a1 1 0 1 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MasterOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 20h18v2H3zm18-1H3L2.147 7.81A2 2 0 1 1 5 6a1.914 1.914 0 0 1-.024.3l2.737 2.189l2.562-4.486A1.948 1.948 0 0 1 10 3a2 2 0 0 1 4 0a1.946 1.946 0 0 1-.276 1.004l2.558 4.485l2.741-2.19A1.906 1.906 0 0 1 19 6a2 2 0 1 1 2.853 1.81ZM4.92 17h14.16l.73-8.77l-4.106 3.281L12 5.017l-3.71 6.494l-4.1-3.28Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Merge(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 11a2.991 2.991 0 0 0-2.815 1.997c-2.995-.014-5.26-.129-6.726-.886a6.205 6.205 0 0 1-3.358-4.326A3.008 3.008 0 1 0 4 7.816v9.368a3 3 0 1 0 2 0v-5.257a8.579 8.579 0 0 0 2.541 1.962c1.847.952 4.36 1.092 7.642 1.108A2.995 2.995 0 1 0 19 11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MergeOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 11a2.991 2.991 0 0 0-2.815 1.997c-2.995-.014-5.26-.129-6.726-.886a6.205 6.205 0 0 1-3.358-4.326A3.008 3.008 0 1 0 4 7.816v9.368a3 3 0 1 0 2 0v-5.257a8.579 8.579 0 0 0 2.541 1.962c1.847.952 4.36 1.092 7.642 1.108A2.995 2.995 0 1 0 19 11M5 21a1 1 0 1 1 1-1a1 1 0 0 1-1 1M5 6a1 1 0 1 1 1-1a1 1 0 0 1-1 1m14 9a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Miscellaneous(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m18 13l.989 7.875A1 1 0 0 1 17.997 22H6.003a1 1 0 0 1-.992-1.125L6 13Zm3-3H3v2h18Zm-6.286-4.196A6.303 6.303 0 0 0 15 3.835C15 2.27 14.552 1 14 1s-1 1.27-1 2.835a7.115 7.115 0 0 0 .115 1.301a4.626 4.626 0 0 0-2.234.001A7.094 7.094 0 0 0 11 3.835C11 2.27 10.552 1 10 1S9 2.27 9 3.835a6.31 6.31 0 0 0 .283 1.971A5.11 5.11 0 0 0 7 9h2a1 1 0 0 1 2 0h2a1 1 0 0 1 2 0h2a5.11 5.11 0 0 0-2.286-3.196"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MiscellaneousOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.36 15L18 21H6l1.64-6zM18 13H6l-2 8a2.185 2.185 0 0 0 2 2h12a2.158 2.158 0 0 0 2-2zm4-3H2v2h20zM10 8a1 1 0 0 1 1 1h2a1 1 0 0 1 2 0h2a5.11 5.11 0 0 0-2.286-3.196A6.303 6.303 0 0 0 15 3.835C15 2.27 14.552 1 14 1s-1 1.27-1 2.835a7.115 7.115 0 0 0 .115 1.301a4.626 4.626 0 0 0-2.234.001A7.094 7.094 0 0 0 11 3.835C11 2.27 10.552 1 10 1S9 2.27 9 3.835a6.31 6.31 0 0 0 .283 1.971A5.11 5.11 0 0 0 7 9h2a1 1 0 0 1 1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ModifiedDate(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 3h-1V1h-2v2H7V1H5v2H4a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2M7.53 21H4v-3.53l9.41-9.41l3.53 3.53ZM19.72 8.81l-1.84 1.84l-3.53-3.53l1.85-1.84a.92.92 0 0 1 1.32 0l2.2 2.2a.94.94 0 0 1 0 1.33"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ModifiedDateOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.372 11.013l.614.614l-6.04 6.04h-.613v-.614zM15.771 7a.667.667 0 0 0-.467.193l-1.22 1.22l2.5 2.5l1.22-1.22a.664.664 0 0 0 0-.94l-1.56-1.56A.655.655 0 0 0 15.772 7m-2.4 2.127L6 16.5V19h2.5l7.372-7.373Z"/><path fill="currentColor" d="M19 1h-2v2H7V1H5v2H4a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2h-1zM4 21V5h16v16Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Molecules(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 9a2.952 2.952 0 0 0-.368.037l-1.478-2.955A2.994 2.994 0 1 0 14.184 3H9.816a2.993 2.993 0 1 0-4.973 3.078l-1.48 2.959A2.963 2.963 0 0 0 3 9a3 3 0 0 0 0 6a2.963 2.963 0 0 0 .364-.037l1.479 2.96A2.994 2.994 0 1 0 9.816 21h4.368a2.993 2.993 0 1 0 4.976-3.074l1.481-2.962A2.963 2.963 0 0 0 21 15a3 3 0 0 0 0-6m-3.632 8.037A2.952 2.952 0 0 0 17 17a2.991 2.991 0 0 0-2.816 2H9.816A2.991 2.991 0 0 0 7 17a2.963 2.963 0 0 0-.364.037l-1.479-2.96a2.983 2.983 0 0 0 0-4.155l1.48-2.959A2.963 2.963 0 0 0 7 7a2.991 2.991 0 0 0 2.816-2h4.368a2.837 2.837 0 0 0 3.175 1.964l1.48 2.962a2.983 2.983 0 0 0 .007 4.156Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoleculesOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 9a2.952 2.952 0 0 0-.368.037l-1.478-2.955A2.994 2.994 0 1 0 14.184 3H9.816a2.993 2.993 0 1 0-4.973 3.078l-1.48 2.959A2.963 2.963 0 0 0 3 9a3 3 0 0 0 0 6a2.963 2.963 0 0 0 .364-.037l1.479 2.96A2.994 2.994 0 1 0 9.816 21h4.368a2.993 2.993 0 1 0 4.976-3.074l1.481-2.962A2.963 2.963 0 0 0 21 15a3 3 0 0 0 0-6M3 13a1 1 0 1 1 1-1a1 1 0 0 1-1 1M17 3a1 1 0 1 1-1 1a1 1 0 0 1 1-1M7 3a1 1 0 1 1-1 1a1 1 0 0 1 1-1m0 18a1 1 0 1 1 1-1a1 1 0 0 1-1 1m10 0a1 1 0 1 1 1-1a1 1 0 0 1-1 1m.368-3.963A2.952 2.952 0 0 0 17 17a2.991 2.991 0 0 0-2.816 2H9.816A2.991 2.991 0 0 0 7 17a2.963 2.963 0 0 0-.364.037l-1.479-2.96a2.983 2.983 0 0 0 0-4.155l1.48-2.959A2.963 2.963 0 0 0 7 7a2.991 2.991 0 0 0 2.816-2h4.368a2.837 2.837 0 0 0 3.175 1.964l1.48 2.962a2.983 2.983 0 0 0 .007 4.156ZM21 13a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Monitoring(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 18h2v3h-2zm5 3v2H8v-2zm4-3H4a3.003 3.003 0 0 1-3-3V4a3.003 3.003 0 0 1 3-3h16a3.003 3.003 0 0 1 3 3v11a3.003 3.003 0 0 1-3 3M4 3a1 1 0 0 0-1 1v11a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1Z"/><path fill="currentColor" d="m16 15l-1.914-6.38L13 13l-1.309-3h-.331L10 14L8.843 9.933L8.309 11H5v-1h2.691L9 7l1.068 3.713L10.64 9h1.669l.487.973L14 4l2 8l.64-2H19v1h-1.64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Move(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.017 12.008L18.019 8.01l-.002 2.993h-5.001v-4.99l2.993.002l-4.002-4.002l-3.998 3.998l2.994.001v4.991h-5L6.002 8.01l-3.998 3.998l4.002 4.002l-.002-2.993h4.999v4.993l-2.994.001l3.998 3.999l4.002-4.002l-2.993.001v-4.992h5l-.001 2.993z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Multistate(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="4" cy="4" r="3" fill="currentColor"/><path fill="currentColor" d="M4 15a3 3 0 1 1 3-3a3 3 0 0 1-3 3m0-5a2 2 0 1 0 2 2a2 2 0 0 0-2-2m8 13a3 3 0 1 1 3-3a3 3 0 0 1-3 3m0-5a2 2 0 1 0 2 2a2 2 0 0 0-2-2m8-11a3 3 0 1 1 3-3a3 3 0 0 1-3 3m0-5a2 2 0 1 0 2 2a2 2 0 0 0-2-2"/><circle cx="4" cy="20" r="3" fill="currentColor"/><circle cx="12" cy="4" r="3" fill="currentColor"/><circle cx="12" cy="12" r="3" fill="currentColor"/><circle cx="20" cy="12" r="3" fill="currentColor"/><circle cx="20" cy="20" r="3" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Namespace(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23 11v2c-.852 0-1 .35-1 1.314v4.498c0 2.514-1.46 3.156-4 3.188v-2c1.404-.032 2 .154 2-1.36v-4.297a3.088 3.088 0 0 1 .294-1.56a1.713 1.713 0 0 1 .997-.783a1.713 1.713 0 0 1-.997-.783A3.088 3.088 0 0 1 20 9.657V5.359c0-1.513-.596-1.327-2-1.359V2c2.54.032 4 .674 4 3.188v4.498c0 .963.148 1.314 1 1.314M1 11v2c.852 0 1 .35 1 1.314v4.498C2 21.326 3.46 21.968 6 22v-2c-1.404-.032-2 .154-2-1.36v-4.297a3.088 3.088 0 0 0-.294-1.56A1.713 1.713 0 0 0 2.709 12a1.713 1.713 0 0 0 .997-.783A3.088 3.088 0 0 0 4 9.657V5.359C4 3.846 4.596 4.032 6 4V2c-2.54.032-4 .674-4 3.188v4.498C2 10.649 1.852 11 1 11m17-2H6v6h12ZM8 11h8v2H8Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Network(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 4a8 8 0 1 0 8 8a8.003 8.003 0 0 0-8-8m-6 8a5.997 5.997 0 0 1 .105-1.095L9.6 14.4v.8a1.605 1.605 0 0 0 1.6 1.6v1.14A6.004 6.004 0 0 1 6 12m9.2 3.2h-.8v-2.4a.802.802 0 0 0-.8-.8H8.8v-1.6h1.6a.802.802 0 0 0 .8-.8V8h1.6a1.6 1.6 0 0 0 1.59-1.5a5.985 5.985 0 0 1 2.137 9.426A1.57 1.57 0 0 0 15.2 15.2M13 1a1 1 0 1 1-1-1a1 1 0 0 1 1 1m5 1a1 1 0 1 1-1-1a1 1 0 0 1 1 1m4 4a1 1 0 1 1-1-1a1 1 0 0 1 1 1m2 6a1 1 0 1 1-1-1a1 1 0 0 1 1 1m-2 6a1 1 0 1 1-1-1a1 1 0 0 1 1 1m-4 4a1 1 0 1 1-1-1a1 1 0 0 1 1 1m-5 1a1 1 0 1 1-1-1a1 1 0 0 1 1 1m-5-1a1 1 0 1 1-1-1a1 1 0 0 1 1 1m-4-4a1 1 0 1 1-1-1a1 1 0 0 1 1 1m-2-6a1 1 0 1 1-1-1a1 1 0 0 1 1 1m2-6a1 1 0 1 1-1-1a1 1 0 0 1 1 1m4-4a1 1 0 1 1-1-1a1 1 0 0 1 1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NetworkFileSystem(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 17h-3v-4h-2v4H8v1H2v2h6v1h8v-1h6v-2h-6zm-1-4h3a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1h-6l-1-1H6a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NetworkFileSystemOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 17h-3v-4h-2v4H8v1H2v2h6v1h8v-1h6v-2h-6z"/><path fill="currentColor" d="M18 4h-5.793l-.853-.854L11.207 3H6a1.502 1.502 0 0 0-1.5 1.5v8A1.502 1.502 0 0 0 6 14h12a1.502 1.502 0 0 0 1.5-1.5v-7A1.502 1.502 0 0 0 18 4m-.5 2v6h-11V5h3.793l.853.854l.147.146z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NetworkPolicy(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 1a1 1 0 1 1 1 1a1 1 0 0 1-1-1M3 7a1 1 0 1 0-1-1a1 1 0 0 0 1 1m4-4a1 1 0 1 0-1-1a1 1 0 0 0 1 1m10 0a1 1 0 1 0-1-1a1 1 0 0 0 1 1m4 4a1 1 0 1 0-1-1a1 1 0 0 0 1 1M10.69 19.883a9.57 9.57 0 0 1-.59-2.201a5.961 5.961 0 0 1-3.995-6.777L9.6 14.4v.8a1.585 1.585 0 0 0 .4 1.045V12H8.8v-1.6h1.6a.802.802 0 0 0 .8-.8V8h1.6a1.6 1.6 0 0 0 1.59-1.5a6.027 6.027 0 0 1 2.353 1.857L18 7.833l1.08.45a7.997 7.997 0 1 0-8.39 11.6M7 21a1 1 0 1 0 1 1a1 1 0 0 0-1-1m5 1a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-9-5a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-2-6a1 1 0 1 0 1 1a1 1 0 0 0-1-1m23 1.5v3.863A8.169 8.169 0 0 1 18 24a8.12 8.12 0 0 1-6-7.596V12.5l6-2.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NetworkPolicyOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="7" cy="2" r="1" fill="currentColor"/><circle cx="3" cy="6" r="1" fill="currentColor"/><circle cx="12" cy="1" r="1" fill="currentColor"/><circle cx="17" cy="2" r="1" fill="currentColor"/><circle cx="21" cy="6" r="1" fill="currentColor"/><path fill="currentColor" d="m17.5 24l-.119-.029A8.623 8.623 0 0 1 11 15.916v-4.218L17.5 9l6.5 2.698v4.218a8.623 8.623 0 0 1-6.381 8.055ZM13 12.865v3.15a6.396 6.396 0 0 0 4.5 5.96a6.396 6.396 0 0 0 4.5-5.96v-3.15l-4.5-1.793Zm10-.503"/><circle cx="12" cy="23" r="1" fill="currentColor"/><path fill="currentColor" d="M10.4 10.4a.802.802 0 0 0 .8-.8V8h1.6a1.6 1.6 0 0 0 1.59-1.5a6.027 6.027 0 0 1 2.353 1.857L18 7.833l1.08.45a7.997 7.997 0 1 0-8.39 11.6a9.57 9.57 0 0 1-.59-2.201a5.961 5.961 0 0 1-3.995-6.777L9.6 14.4v.8a1.585 1.585 0 0 0 .4 1.045V12H8.8v-1.6Z"/><circle cx="3" cy="18" r="1" fill="currentColor"/><circle cx="7" cy="22" r="1" fill="currentColor"/><circle cx="1" cy="12" r="1" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NeuralNetwork(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.564 12.027a1.436 1.436 0 1 0 0-2.87c-.031 0-.06.006-.092.008l-.482-1.9a1.43 1.43 0 1 0-1.918-2.104l-4.89-2.273a1.415 1.415 0 0 0 .08-.452a1.436 1.436 0 1 0-2.871 0a1.404 1.404 0 0 0 .015.146l-3.114.955a1.43 1.43 0 0 0-2.44 1.432l-2.93 2.616a1.42 1.42 0 0 0-.734-.209a1.431 1.431 0 0 0-.43 2.798l-.194 1.866a1.432 1.432 0 1 0 1.097 2.532l.966.425l-.003.03a1.434 1.434 0 0 0 2.845.25h1.104a1.43 1.43 0 0 0 2.785-.128l3.107.102a1.403 1.403 0 0 0 1.664.125a11.69 11.69 0 0 1 4.297 4.039a1.423 1.423 0 0 0-.107.537a1.436 1.436 0 1 0 2.292-1.145a6.723 6.723 0 0 1-1.038-2.422a1.422 1.422 0 0 0 .45-2.472zm.565-1.435a.564.564 0 1 1-.565-.565a.565.565 0 0 1 .565.565M20.985 9.28a1.434 1.434 0 0 0-.844 1.19H19.02a1.427 1.427 0 0 0-.541-1.003l1.407-1.933a1.257 1.257 0 0 0 .647-.04Zm-2.274 5.567l-2.898-.742a1.422 1.422 0 0 0-.345-.878l1.17-1.575a1.376 1.376 0 0 0 1.368.307l1.252 1.934a1.428 1.428 0 0 0-.547.954M8.251 5.385l4.081 2.136a1.289 1.289 0 0 0-.017.785l-2.62 1.347a1.427 1.427 0 0 0-1.077-.497c-.012 0-.023.003-.035.004L7.64 5.79a1.43 1.43 0 0 0 .611-.405M19.4 7.354L18.033 9.23a1.373 1.373 0 0 0-1.361.272l-1.591-1.2a1.102 1.102 0 0 0 .035-.53l3.707-1.06a1.44 1.44 0 0 0 .577.64m-1.804 3.802a.564.564 0 1 1 .565-.564a.565.565 0 0 1-.565.564m-3.899-2.65a.564.564 0 1 1 .565-.565a.565.565 0 0 1-.565.564m-5.08 1.521a.564.564 0 1 1-.564.565a.565.565 0 0 1 .565-.565m1.436.565a1.42 1.42 0 0 0-.098-.511l2.568-1.32a1.434 1.434 0 0 0 1.175.615l.024-.002l-.434 3.66l-3.906.494a1.418 1.418 0 0 0-.24-.054l-.297-1.47a1.435 1.435 0 0 0 1.208-1.412M8.95 14.325a.564.564 0 1 1-.565.565a.565.565 0 0 1 .565-.565m1.076-.375l3.053-.385a1.402 1.402 0 0 0-.005 1.173l-2.713-.088a1.423 1.423 0 0 0-.335-.7m4.213-4.682a1.441 1.441 0 0 0 .629-.5l1.486 1.12a1.38 1.38 0 0 0-.03 1.351l-1.243 1.67a1.4 1.4 0 0 0-1.265-.07Zm5.89-3.708a.564.564 0 1 1-.564.564a.565.565 0 0 1 .565-.564m-1.328.026a1.423 1.423 0 0 0-.107.538c0 .035.008.068.01.103l-3.732 1.066a1.433 1.433 0 0 0-1.274-.788c-.043 0-.085.01-.127.013l-.521-2.67a1.428 1.428 0 0 0 .893-.52ZM12.826 1.87a.564.564 0 1 1-.564.565a.565.565 0 0 1 .564-.565m-1.284 1.192a1.438 1.438 0 0 0 .996.779l.548 2.805a1.44 1.44 0 0 0-.526.429L8.515 4.959a1.356 1.356 0 0 0 .023-.975Zm-4.36.808a.564.564 0 1 1-.564.565a.565.565 0 0 1 .564-.565m-.04 1.996l.95 3.392a1.442 1.442 0 0 0-.672.545l-3.799-.969l.003-.023a1.421 1.421 0 0 0-.324-.896l2.824-2.519a1.426 1.426 0 0 0 1.017.47m-4.806 7.587a.564.564 0 1 1-.564-.564a.565.565 0 0 1 .564.564m-.27-1.405l.189-1.808a1.433 1.433 0 0 0 1.269-.914l3.698.943a1.304 1.304 0 0 0-.014.576l-4.343 1.69a1.431 1.431 0 0 0-.8-.486m.122-3.802a.564.564 0 1 1-.564.565a.565.565 0 0 1 .564-.565m.817 5.93a1.371 1.371 0 0 0 .114-1.202l4.266-1.66a1.434 1.434 0 0 0 .891.667l-1.865 2.583a1.429 1.429 0 0 0-2.683-.07Zm2.054 1.415a.564.564 0 1 1 .565-.565a.565.565 0 0 1-.565.565m3.407-3.306l.179 1.202a1.434 1.434 0 0 0-1.12 1.29h-.7Zm5.916 2.435a.564.564 0 1 1 .565-.565a.565.565 0 0 1-.565.565m6.373 5.795a.564.564 0 1 1 .564-.564a.565.565 0 0 1-.564.564m0-2a1.43 1.43 0 0 0-1.023.431a12.025 12.025 0 0 0-4.21-3.927a1.425 1.425 0 0 0 .199-.358l3.023.723a1.433 1.433 0 0 0 1.344 1.074a5.832 5.832 0 0 0 .775 2.068c-.037-.003-.071-.01-.108-.01m-.626-2.924a.564.564 0 1 1 .564-.565a.565.565 0 0 1-.564.565m.444-1.923a1.303 1.303 0 0 0-.87-.006l-1.249-1.928a1.433 1.433 0 0 0 .52-.763h1.213a1.435 1.435 0 0 0 .883.962Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Node(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="12" cy="3" r="1" fill="currentColor"/><circle cx="20" cy="8" r="1" fill="currentColor"/><circle cx="20" cy="16" r="1" fill="currentColor"/><circle cx="4" cy="8" r="1" fill="currentColor"/><circle cx="4" cy="16" r="1" fill="currentColor"/><path fill="currentColor" d="M20 14v-4a1.992 1.992 0 0 1-1.481-3.333l-4.783-2.69a1.983 1.983 0 0 1-3.472 0l-4.783 2.69A1.992 1.992 0 0 1 4 10v4a1.992 1.992 0 0 1 1.481 3.333l4.783 2.69a1.991 1.991 0 0 1 1.236-.952v-5.142a2 2 0 1 1 1 0v5.142a1.991 1.991 0 0 1 1.236.953l4.783-2.69A1.992 1.992 0 0 1 20 14"/><circle cx="12" cy="21" r="1" fill="currentColor"/><circle cx="12" cy="12" r="1" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NodeOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="12" cy="2" r="1" fill="currentColor"/><circle cx="21" cy="6" r="1" fill="currentColor"/><circle cx="21" cy="17" r="1" fill="currentColor"/><circle cx="12" cy="22" r="1" fill="currentColor"/><path fill="currentColor" d="M21 8a2.006 2.006 0 0 1-2-2a1.94 1.94 0 0 1 .26-.96L14 2a2.005 2.005 0 0 1-1.01 1.73a1.95 1.95 0 0 1-1.98 0A2.005 2.005 0 0 1 10 2L4.74 5.04A1.94 1.94 0 0 1 5 6a2.006 2.006 0 0 1-2 2a2.08 2.08 0 0 1-.5-.07v7.14A2.078 2.078 0 0 1 3 15a1.991 1.991 0 0 1 1.5.69a2.191 2.191 0 0 1 .39.68A1.896 1.896 0 0 1 5 17a1.987 1.987 0 0 1-.58 1.41l3.84 2.21l1.78 1.03a1.278 1.278 0 0 1 .07-.28a.254.254 0 0 0 .02-.07a1.789 1.789 0 0 1 .12-.24c.01-.03.03-.06.04-.09c.05-.07.09-.14.14-.2c.03-.03.05-.06.08-.09a1.213 1.213 0 0 1 .16-.16a.525.525 0 0 1 .11-.1a1.627 1.627 0 0 1 .17-.11a1.19 1.19 0 0 1 .15-.09c.06-.03.11-.05.17-.08l.21-.06a.742.742 0 0 1 .14-.04a1.826 1.826 0 0 1 .76 0a.742.742 0 0 1 .14.04l.21.06c.06.03.11.05.17.08a1.19 1.19 0 0 1 .15.09a1.627 1.627 0 0 1 .17.11a.525.525 0 0 1 .11.1a1.213 1.213 0 0 1 .16.16c.03.03.05.06.08.09c.05.06.09.13.14.2c.01.03.03.06.04.09a1.789 1.789 0 0 1 .12.24a.254.254 0 0 0 .02.07a1.278 1.278 0 0 1 .07.28l1.78-1.03l3.84-2.21A1.987 1.987 0 0 1 19 17a1.896 1.896 0 0 1 .11-.63a2.191 2.191 0 0 1 .39-.68A1.991 1.991 0 0 1 21 15a2.078 2.078 0 0 1 .5.07V7.93A2.08 2.08 0 0 1 21 8m-9 5a1 1 0 1 1 1-1a1.003 1.003 0 0 1-1 1m7.5.3A4.003 4.003 0 0 0 17 17a3.554 3.554 0 0 0 .06.56l-2.47 1.42a3.993 3.993 0 0 0-1.59-.84v-3.32a3 3 0 1 0-2 0v3.32a3.993 3.993 0 0 0-1.59.84l-2.47-1.42A3.554 3.554 0 0 0 7 17a4.003 4.003 0 0 0-2.5-3.7V9.7A4.004 4.004 0 0 0 7 6.04L9.15 4.8a3.984 3.984 0 0 0 5.7 0L17 6.04a4.004 4.004 0 0 0 2.5 3.66Z"/><circle cx="3" cy="17" r="1" fill="currentColor"/><circle cx="3" cy="6" r="1" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Organisms(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.5 11a1.5 1.5 0 0 0-.324.037L19.67 8.43A1.495 1.495 0 1 0 17.092 7h-2.184a1.483 1.483 0 0 0-1.732-.963L11.67 3.428A1.495 1.495 0 1 0 9.092 2H6.908A1.496 1.496 0 1 0 4.33 3.428l-1.506 2.61a1.5 1.5 0 1 0 .012 2.921l1.502 2.603a1.47 1.47 0 0 0-.008 1.866l-1.506 2.61a1.5 1.5 0 1 0 .012 2.921l1.502 2.603A1.495 1.495 0 1 0 6.908 23h2.184a1.496 1.496 0 1 0 2.57-1.438l1.502-2.603A1.478 1.478 0 0 0 14.908 18h2.184a1.496 1.496 0 1 0 2.57-1.438l1.502-2.603A1.499 1.499 0 1 0 21.5 11m-3 5a1.496 1.496 0 0 0-1.408 1h-2.184a1.489 1.489 0 0 0-1.72-.966l-1.51-2.616a1.47 1.47 0 0 0-.016-1.856l1.502-2.603A1.478 1.478 0 0 0 14.908 8h2.184a1.483 1.483 0 0 0 1.732.963l1.506 2.608a1.47 1.47 0 0 0-.008 1.848l-1.51 2.615A1.503 1.503 0 0 0 18.5 16m-8 5a1.496 1.496 0 0 0-1.408 1H6.908a1.489 1.489 0 0 0-1.72-.967l-1.51-2.615a1.47 1.47 0 0 0-.008-1.846l1.506-2.61A1.483 1.483 0 0 0 6.908 13h2.184a1.483 1.483 0 0 0 1.732.963l1.507 2.608a1.47 1.47 0 0 0-.01 1.847l-1.51 2.615A1.503 1.503 0 0 0 10.5 21M3.67 6.572l1.506-2.61A1.483 1.483 0 0 0 6.908 3h2.184a1.483 1.483 0 0 0 1.732.963L12.33 6.57a1.47 1.47 0 0 0-.008 1.847l-1.51 2.615a1.49 1.49 0 0 0-1.72.967H6.908a1.49 1.49 0 0 0-1.72-.967l-1.51-2.615a1.47 1.47 0 0 0-.008-1.846"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OrganismsOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.5 11a1.5 1.5 0 0 0-.324.037L19.67 8.43A1.495 1.495 0 1 0 17.092 7h-2.184a1.483 1.483 0 0 0-1.732-.963L11.67 3.428A1.495 1.495 0 1 0 9.092 2H6.908A1.496 1.496 0 1 0 4.33 3.428l-1.506 2.61a1.5 1.5 0 1 0 .012 2.921l1.502 2.603a1.47 1.47 0 0 0-.008 1.866l-1.506 2.61a1.5 1.5 0 1 0 .012 2.921l1.502 2.603A1.495 1.495 0 1 0 6.908 23h2.184a1.496 1.496 0 1 0 2.57-1.438l1.502-2.603A1.478 1.478 0 0 0 14.908 18h2.184a1.496 1.496 0 1 0 2.57-1.438l1.502-2.603A1.499 1.499 0 1 0 21.5 11m-3-4a.5.5 0 1 1-.5.5a.5.5 0 0 1 .5-.5m-5 0a.5.5 0 1 1-.5.5a.5.5 0 0 1 .5-.5m-3-5a.5.5 0 1 1-.5.5a.5.5 0 0 1 .5-.5m-5 0a.5.5 0 1 1-.5.5a.5.5 0 0 1 .5-.5m-3 6a.5.5 0 1 1 .5-.5a.5.5 0 0 1-.5.5m1.178.418a1.47 1.47 0 0 0-.008-1.846l1.506-2.61A1.483 1.483 0 0 0 6.908 3h2.184a1.483 1.483 0 0 0 1.732.963L12.33 6.57a1.47 1.47 0 0 0-.008 1.847l-1.51 2.615a1.49 1.49 0 0 0-1.72.967H6.908a1.49 1.49 0 0 0-1.72-.967ZM11 12.5a.5.5 0 1 1-.5-.5a.5.5 0 0 1 .5.5m-5 0a.5.5 0 1 1-.5-.5a.5.5 0 0 1 .5.5M2.5 18a.5.5 0 1 1 .5-.5a.5.5 0 0 1-.5.5m3 5a.5.5 0 1 1 .5-.5a.5.5 0 0 1-.5.5m5 0a.5.5 0 1 1 .5-.5a.5.5 0 0 1-.5.5m1.822-4.582l-1.51 2.615a1.489 1.489 0 0 0-1.72.967H6.908a1.489 1.489 0 0 0-1.72-.967l-1.51-2.615a1.47 1.47 0 0 0-.008-1.846l1.506-2.61A1.483 1.483 0 0 0 6.908 13h2.184a1.483 1.483 0 0 0 1.732.963l1.507 2.608a1.47 1.47 0 0 0-.01 1.847M13.5 18a.5.5 0 1 1 .5-.5a.5.5 0 0 1-.5.5m5 0a.5.5 0 1 1 .5-.5a.5.5 0 0 1-.5.5m1.822-4.581l-1.51 2.615a1.489 1.489 0 0 0-1.72.966h-2.184a1.489 1.489 0 0 0-1.72-.966l-1.51-2.616a1.47 1.47 0 0 0-.016-1.856l1.502-2.603A1.478 1.478 0 0 0 14.908 8h2.184a1.483 1.483 0 0 0 1.732.963l1.506 2.608a1.47 1.47 0 0 0-.008 1.848M21.5 13a.5.5 0 1 1 .5-.5a.5.5 0 0 1-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Organization(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.01 10.99h-7v-2h-2v2H3.47v4h2v-2h5.54v2h2v-2h5.5v2h2v-4z"/><circle cx="12.01" cy="4.51" r="2.5" fill="currentColor"/><circle cx="4.47" cy="19.49" r="2.5" fill="currentColor"/><circle cx="12.01" cy="19.49" r="2.5" fill="currentColor"/><circle cx="19.51" cy="19.49" r="2.5" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OrganizationOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 11V9h-2l.025 2H3.5v4h2v-2H11v2h2v-2h5.505l-.005 2h2v-4zm-1-9a2.5 2.5 0 1 0 2.5 2.5A2.5 2.5 0 0 0 12 2m0 3.5a1 1 0 1 1 1-1a1 1 0 0 1-1 1M4.5 17A2.5 2.5 0 1 0 7 19.5A2.5 2.5 0 0 0 4.5 17m0 3.5a1 1 0 1 1 1-1a1 1 0 0 1-1 1M12 17a2.5 2.5 0 1 0 2.5 2.5A2.5 2.5 0 0 0 12 17m0 3.5a1 1 0 1 1 1-1a1 1 0 0 1-1 1m7.5-3.5a2.5 2.5 0 1 0 2.5 2.5a2.5 2.5 0 0 0-2.5-2.5m0 3.5a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Outbound(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 2H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h8v-2H4V6l8 5l8-5v7h2V4a2 2 0 0 0-2-2m-8 7L4 4h16Zm10 8l-4 4v-3h-4v-2h4v-3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Package(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 6v2h2v11a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8h2V6Zm9 12H6v-2h5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PackageOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 16h4v2H8z"/><path fill="currentColor" d="M2 6v2h2v11a2.003 2.003 0 0 0 2 2h12a2.003 2.003 0 0 0 2-2V8h2V6Zm16 13H6V8h12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PackageUpgrade(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.77 14.25L20 16.49V8h2V6H2v2h2v11a2 2 0 0 0 2 2h5ZM6 16h5v2H6Z"/><path fill="currentColor" d="m22.01 20.53l-2.83-2.83l-1.41-1.41l-4.25 4.25l1.41 1.41l2.84-2.83l2.83 2.83z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PackageUpgradeOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m22.243 20.243l-2.829-2.829L18 16l-4.249 4.249l1.415 1.414L18 18.829l2.828 2.828zM7 15h4v2H7z"/><path fill="currentColor" d="M22 6H2v2h2v11a2.003 2.003 0 0 0 2 2h5l7-7l2 2V8h2ZM6 8h12v3l-8 8H6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Packages(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.04 6.01h-16v2h16zm-2.02-4H6.01v2h12.01zM2 10v2h2v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8h2v-2Zm9 9H6v-2h5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PackagesOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 6h16v2H4zm2-4h12v2H6zm-4 8v2h2v8a2.003 2.003 0 0 0 2 2h12a2.006 2.006 0 0 0 2.004-2L20 12h2v-2Zm16 10H6v-8h12Z"/><path fill="currentColor" d="M8 17h4v2H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PatchFixes(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.48 18.71a3.996 3.996 0 0 1-5.163-5.272l2.619 2.619l2.12-2.121l-2.618-2.619a3.988 3.988 0 0 1 5.2 5.308l1.933 1.933A7.96 7.96 0 0 0 20 14A17.115 17.115 0 0 0 13.5.67a21.494 21.494 0 0 1 .74 4.8a3.47 3.47 0 0 1-3.41 3.73A3.64 3.64 0 0 1 7.2 5.47l.03-.36A13.768 13.768 0 0 0 4 14a8 8 0 0 0 12.43 6.66Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PatchFixesOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5.67a21.494 21.494 0 0 1 .74 4.8a3.47 3.47 0 0 1-3.41 3.73A3.64 3.64 0 0 1 7.2 5.47l.03-.36A13.768 13.768 0 0 0 4 14a7.945 7.945 0 0 0 9.076 7.92l.261-.04a7.954 7.954 0 0 0 1-.229l.021-.005a7.924 7.924 0 0 0 .979-.382c.067-.031.133-.064.2-.096c.173-.086.336-.19.503-.286l.01.01c.036-.022.067-.052.103-.074c.084-.05.174-.091.256-.145l.021-.013l.003-.002a7.944 7.944 0 0 0 .701-.533l.027-.023a8.053 8.053 0 0 0 .923-.916l.007-.007c.168-.198.327-.404.476-.617l.004-.004l.004-.006a7.99 7.99 0 0 0 .562-.952l.072-.148a7.933 7.933 0 0 0 .405-1l.01-.029a7.952 7.952 0 0 0 .252-1.062a6.9 6.9 0 0 0 .037-.245A8.024 8.024 0 0 0 20 14A17.115 17.115 0 0 0 13.5.67m1.862 15.763a3.278 3.278 0 0 0-4.443-4.119l2.147 2.16l-1.498 1.507l-2.197-2.16a3.293 3.293 0 0 0 4.094 4.47l1.112 1.119l-.078.037a5.735 5.735 0 0 1-.744.291a6.003 6.003 0 0 1-.771.172l-.177.028A5.942 5.942 0 0 1 12 20a6.007 6.007 0 0 1-6-6a11.875 11.875 0 0 1 .858-4.433A5.562 5.562 0 0 0 10.83 11.2a5.342 5.342 0 0 0 5.281-4.429A14.787 14.787 0 0 1 18 14a5.975 5.975 0 0 1-.066.833l-.029.188a6.247 6.247 0 0 1-.17.74l-.027.082a5.792 5.792 0 0 1-.295.728l-.062.128a5.68 5.68 0 0 1-.409.695l-.016.025c-.08.115-.18.21-.267.319z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Patterns(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="4.02" cy="4.015" r="3" fill="currentColor"/><path fill="currentColor" d="M15.03 12.025a2.991 2.991 0 0 0-2-2.816v-2.38a2.99 2.99 0 0 0 1.807-1.814h2.367a2.99 2.99 0 0 0 1.816 1.816v2.38a3.001 3.001 0 1 0 2-.005V6.83a2.993 2.993 0 1 0-3.816-3.816h-2.367a2.993 2.993 0 1 0-3.807 3.82v2.374a2.983 2.983 0 0 0 0 5.632v2.375a2.988 2.988 0 0 0-1.827 1.823H6.838a2.989 2.989 0 0 0-1.818-1.82v-2.375a2.999 2.999 0 1 0-2-.006v2.38a2.993 2.993 0 1 0 3.815 3.821h2.37a2.993 2.993 0 1 0 3.825-3.817v-2.38a2.991 2.991 0 0 0 2-2.817"/><circle cx="20.02" cy="20.035" r="3" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PatternsOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.98 9.193V6.822a2.992 2.992 0 1 0-3.796-3.81h-2.368a2.992 2.992 0 1 0-3.806 3.813v2.37a2.972 2.972 0 0 0 0 5.624v2.37a2.986 2.986 0 0 0-1.828 1.822H6.818A2.986 2.986 0 0 0 5 17.193v-2.37a3.009 3.009 0 1 0-2-.007v2.377a2.99 2.99 0 1 0 3.815 3.814h2.37a2.992 2.992 0 1 0 3.825-3.81v-2.378a2.982 2.982 0 0 0 0-5.623V6.819a2.987 2.987 0 0 0 1.806-1.81h2.368a2.955 2.955 0 0 0 1.796 1.813v2.377A2.968 2.968 0 0 0 17 12a3 3 0 1 0 3.98-2.807M4 21a1 1 0 1 1 1-1a1 1 0 0 1-1 1m0-8a1 1 0 1 1 1-1a1 1 0 0 1-1 1m8 8a1 1 0 1 1 1-1a1 1 0 0 1-1 1m0-8a1.033 1.033 0 0 1-1.02-1.017A1.003 1.003 0 0 1 12 11a.969.969 0 0 1 .98.983A.999.999 0 0 1 12 13m-.02-8.006a.999.999 0 0 1 0-1.997a.999.999 0 1 1 0 1.997M20 12.98a.984.984 0 0 1-1-.98a1.018 1.018 0 0 1 1.02-1a.99.99 0 0 1-.02 1.981M20 5a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/><path fill="currentColor" d="M20 17a3 3 0 1 0 3 3a3 3 0 0 0-3-3m0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1M4 1a3 3 0 1 0 3 3a3 3 0 0 0-3-3m0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Performance(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 22.007H5v-2h10ZM22 4l-4.735 5.955l-.22.27l-5.63 7.19a2.001 2.001 0 1 1-2.83-2.83ZM2.645 7.234A10.843 10.843 0 0 0 1.19 15H2v-1a9.685 9.685 0 0 1 1.69-5.52ZM12 2a10.958 10.958 0 0 0-8.119 3.597L5.025 6.96A7.428 7.428 0 0 1 10 5a7.432 7.432 0 0 1 4.997 1.978l3.55-2.802A10.936 10.936 0 0 0 12 2m6.83 9.2l-.233.287l-.728.93A10.118 10.118 0 0 1 18 14v1h4.81a10.879 10.879 0 0 0-1.183-7.318Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersistentVolume(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 1H10a2 2 0 0 0-2 2v10.369a6 6 0 0 1 .989.454A5.983 5.983 0 0 1 17.65 21H20a2 2 0 0 0 2-2V3a2 2 0 0 0-2-2M9 3a1 1 0 1 1 1 1a1 1 0 0 1-1-1m6 10a4.973 4.973 0 0 1-2-.422V10h-2.578A4.997 4.997 0 1 1 15 13m5 7a1 1 0 1 1 1-1a1 1 0 0 1-1 1m0-16a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/><path fill="currentColor" d="M12 23a3.973 3.973 0 0 1-3-1.38a3.967 3.967 0 0 1-1-2.617V19a2 2 0 1 0-2 2a1.928 1.928 0 0 0 .79-.175a5.95 5.95 0 0 0 .927 1.765A3.834 3.834 0 0 1 6 23a4 4 0 1 1 3-6.62a3.967 3.967 0 0 1 1 2.617V19a2 2 0 1 0 2-2a1.928 1.928 0 0 0-.79.176a5.95 5.95 0 0 0-.928-1.766A3.834 3.834 0 0 1 12 15a4 4 0 0 1 0 8"/><circle cx="15" cy="8" r="2" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pin(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.27 9.81h-2.82L9.77 4.13l.71-.71l-1.42-1.41l-7.07 7.07l1.42 1.41l.71-.71l5.67 5.68h-.01v2.83l1.42 1.42l3.54-3.55l4.77 4.77l1.41-1.41l-4.77-4.77l3.53-3.53z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9.675 6.83l3.546 3.537a5.004 5.004 0 0 0 2.836 1.415l-4.255 4.244a4.941 4.941 0 0 0-1.418-2.83L6.838 9.66zm.709-3.537l-7.091 7.075a1.002 1.002 0 0 0 1.418 1.415l.709-.708l3.546 3.537a2.993 2.993 0 0 1 0 4.245l1.418 1.415l4.234-4.223L19.582 21H21v-1.415l-4.964-4.952l4.276-4.266l-1.418-1.415a3.01 3.01 0 0 1-4.255 0l-3.546-3.538l.71-.707a1.002 1.002 0 1 0-1.419-1.415"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pipeline(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 9v1H5V4h1V2H1v2h1v7a2 2 0 0 0 2 2h5v1h2V9Zm13 11v-8a2 2 0 0 0-2-2h-5V9h-2v5h2v-1h4v7h-1v2h5v-2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PipelineOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 20v-8a2 2 0 0 0-2-2h-5V9h-2v6h2v-1h3v6h-1v2h6v-2Zm-1 0h-2v-7h-4v-2h5a1 1 0 0 1 1 1ZM9 9v1H6V4h1V2H1v2h1v8a2 2 0 0 0 2 2h5v1h2V9Zm0 4H4a1 1 0 0 1-1-1V4h2v7h4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pod(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 19.003A1.999 1.999 0 0 0 8 21h8a1.999 1.999 0 0 0 2-1.997V18H6Zm9-15.011A.995.995 0 0 0 14.002 3H9.998A.995.995 0 0 0 9 3.992V5h6ZM14.993 6H9L4 17h16Zm-2.983 7.01a2 2 0 1 1 2-2a2 2 0 0 1-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PodAutoscaler(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 11H3.5V8.5L0 12l3.5 3.5V13H5zm19 1l-3.5-3.5V11H19v2h1.5v2.5zM6 18.752V18h12v1.003A1.998 1.998 0 0 1 16 21H8a1.996 1.996 0 0 1-2-1.997Zm9-14.76a.986.986 0 0 0-.292-.702a.998.998 0 0 0-.706-.29H9.998a.998.998 0 0 0-.92.606a.983.983 0 0 0-.078.386V5h6zm3.748 10.258l-1.251-2.75l-1.252-2.75L14.993 6H9L7.75 8.75L6.5 11.5l-1.25 2.75L4 17h16Zm-4.895-2.461a2.007 2.007 0 1 1 .157-.779a2.003 2.003 0 0 1-.157.779"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PodAutoscalerOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 4a1.004 1.004 0 0 0-1-1h-4a1.004 1.004 0 0 0-1 1v1h6Zm-1.3 4l3.2 7H7.1l3.18-7ZM15 6H9L4 17h16Z"/><circle cx="12" cy="11" r="1.5" fill="currentColor"/><path fill="currentColor" d="M16 18v1H8v-1H6v1.003A1.998 1.998 0 0 0 8 21h8a1.998 1.998 0 0 0 2-1.997V18ZM5 11H3.5V8.5L0 12l3.5 3.5V13H5zm19 1l-3.5-3.5V11H19v2h1.5v2.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PodOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 4a1.004 1.004 0 0 0-1-1h-4a1.004 1.004 0 0 0-1 1v1h6Zm-1.3 4l3.2 7H7.1l3.18-7ZM15 6H9L4 17h16Z"/><circle cx="12" cy="11" r="1.5" fill="currentColor"/><path fill="currentColor" d="M16 18v1H8v-1H6v1.003A1.998 1.998 0 0 0 8 21h8a1.998 1.998 0 0 0 2-1.997V18Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PodSecurity(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.993 4.012H6.988V2.998a.998.998 0 0 1 1-.998h4.007a.998.998 0 0 1 .998.998zM12 15.994H2l2.5-5.496L7 5.003h5.993l1.308 2.871l1.309 2.872l-1.805.802L12 12.35zM12 10a2 2 0 1 0-.586 1.414A1.994 1.994 0 0 0 12 10m.085 6.995H3.993V18a2.001 2.001 0 0 0 2.002 2h7.239a7.649 7.649 0 0 1-.737-1.432a7.191 7.191 0 0 1-.412-1.573M22 13v3a6.405 6.405 0 0 1-1.282 3.804A5.776 5.776 0 0 1 17.5 22a5.776 5.776 0 0 1-3.217-2.196A6.405 6.405 0 0 1 13 16v-3l2.25-1l2.25-1l2.25 1Zm-4.5-.905l-1.75.778l-1.75.777v2.85h3.5v4.465a4.788 4.788 0 0 0 2.348-1.678A5.803 5.803 0 0 0 21 16.495h-3.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PodSecurityOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m19.75 12l-2.25-1l-2.25 1L13 13v3a6.405 6.405 0 0 0 1.282 3.804A5.776 5.776 0 0 0 17.5 22a5.776 5.776 0 0 0 3.218-2.196A6.405 6.405 0 0 0 22 16v-3Zm.098 7.287a4.787 4.787 0 0 1-2.348 1.678V16.5H14v-2.85l1.75-.777l1.75-.778v4.4H21a5.803 5.803 0 0 1-1.152 2.792M13 4a1.004 1.004 0 0 0-1-1H8a1.004 1.004 0 0 0-1 1v1h6Z"/><circle cx="10" cy="11" r="1.5" fill="currentColor"/><path fill="currentColor" d="M12 17v-2H5.1l3.18-7h3.42l1.704 3.726l1.825-.811L12.99 6H7L2 17zm2 4a5.797 5.797 0 0 1-.519-.597A7.488 7.488 0 0 1 12.68 19H6v-1H4v1.003A1.998 1.998 0 0 0 6 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Primitive(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="12" cy="12" r="2.197" fill="currentColor"/><path fill="currentColor" d="M21.634 6.437c-.709-1.228-2.411-1.732-4.796-1.424c-.302.04-.613.094-.93.158C15.009 2.619 13.608 1 12 1c-1.61 0-3.013 1.624-3.913 4.183c-2.738-.542-4.907-.156-5.722 1.254c-.812 1.407-.07 3.47 1.757 5.563c-1.827 2.094-2.57 4.156-1.757 5.563c.582 1.008 1.85 1.496 3.525 1.496a11.475 11.475 0 0 0 2.197-.243C8.987 21.376 10.39 23 12 23c1.607 0 3.009-1.619 3.909-4.17c.316.063.627.117.93.157a10.182 10.182 0 0 0 1.296.088c1.705 0 2.927-.52 3.5-1.512c.812-1.407.07-3.47-1.757-5.563c1.827-2.094 2.569-4.156 1.756-5.563m-4.649-.297a9.08 9.08 0 0 1 1.16-.08c1.26 0 2.15.327 2.507.945c.504.874-.067 2.442-1.552 4.168a20.506 20.506 0 0 0-2.302-1.953a19.975 19.975 0 0 0-.547-2.936c.244-.047.5-.114.734-.144m-3.072 9.174c-.647.373-1.284.703-1.907.993a24.552 24.552 0 0 1-1.92-.993a24.472 24.472 0 0 1-1.819-1.16a23.827 23.827 0 0 1 0-4.309a24.472 24.472 0 0 1 3.73-2.149a24.72 24.72 0 0 1 1.916.99c.649.375 1.255.764 1.82 1.16a23.845 23.845 0 0 1 0 4.308c-.565.396-1.171.785-1.82 1.16m1.628.335a17.876 17.876 0 0 1-.392 1.824a18.464 18.464 0 0 1-1.76-.583a25.595 25.595 0 0 0 2.152-1.241m-4.924 1.244a18.413 18.413 0 0 1-1.766.581a17.86 17.86 0 0 1-.392-1.825c.344.22.696.437 1.06.647c.365.21.73.407 1.098.597M7.074 13.25A18.386 18.386 0 0 1 5.675 12a18.389 18.389 0 0 1 1.4-1.25c-.02.41-.033.826-.033 1.25s.012.84.032 1.25M8.46 8.352a17.867 17.867 0 0 1 .391-1.823a18.222 18.222 0 0 1 1.77.577a24.667 24.667 0 0 0-2.16 1.246m4.93-1.242a18.488 18.488 0 0 1 1.76-.583a17.876 17.876 0 0 1 .392 1.824A26.075 26.075 0 0 0 13.39 7.11m3.537 3.64A18.401 18.401 0 0 1 18.324 12a18.388 18.388 0 0 1-1.398 1.25c.02-.41.032-.826.032-1.25s-.012-.84-.032-1.25M12 2.129c.998 0 2.056 1.234 2.807 3.31a20.802 20.802 0 0 0-2.809 1.011a20.52 20.52 0 0 0-2.807-1.005C9.942 3.365 11 2.129 12 2.129M3.348 7.005c.351-.608 1.252-.934 2.52-.934a10.262 10.262 0 0 1 1.884.201a19.968 19.968 0 0 0-.55 2.948A20.489 20.489 0 0 0 4.9 11.173C3.414 9.447 2.843 7.88 3.348 7.005m0 9.99c-.505-.874.066-2.442 1.552-4.168a20.501 20.501 0 0 0 2.3 1.953a19.944 19.944 0 0 0 .552 2.953c-2.247.429-3.898.14-4.405-.738M12 21.87c-.999 0-2.058-1.237-2.81-3.317A20.448 20.448 0 0 0 12 17.55a20.802 20.802 0 0 0 2.806 1.01c-.75 2.076-1.81 3.31-2.807 3.31m8.652-4.876c-.46.792-1.793 1.109-3.667.865c-.235-.03-.49-.097-.734-.144a19.976 19.976 0 0 0 .547-2.936a20.504 20.504 0 0 0 2.302-1.953c1.485 1.726 2.056 3.294 1.552 4.168"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ProductClasses(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 2H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1M5 21a2 2 0 1 1 2-2a2 2 0 0 1-2 2m2-9H3V3h4Z"/><circle cx="5.01" cy="19.01" r="1" fill="currentColor"/><path fill="currentColor" d="M10.01 2v9.01h5V2zM17 2v9.01h5V2zm-6.99 11v9h5v-9zM17 13v9h5v-9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ProductClassesOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="5" cy="19" r="1" fill="currentColor"/><path fill="currentColor" d="M4 4h2v9H4z"/><path fill="currentColor" d="M7 2H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1m0 19H3V3h4Zm7-18v7h-3V3zm1-1h-5v9h5zm6 1v7h-3V3zm1-1h-5v9h5zm-8 12v7h-3v-7zm1-1h-5v9h5zm6 1v7h-3v-7zm1-1h-5v9h5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ProductSubscriptions(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 8H4V6h16Zm-2-6H6v2h12Zm4 10v8a2 2 0 0 1-2 2H4a2.006 2.006 0 0 1-2-2v-8a2.006 2.006 0 0 1 2-2h16a2.006 2.006 0 0 1 2 2m-8.073 5.042l2.323-1.986l-3.059-.256L12 12l-1.191 2.8l-3.059.256l2.323 1.986l-.7 2.958L12 18.428L14.627 20Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ProductSubscriptionsOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.841 15.1L12 13l-.841 2.1L9 15.292l1.64 1.489L10.146 19L12 17.821L13.854 19l-.494-2.219L15 15.292zM6 2h12v2H6zM4 6h16v2H4z"/><path fill="currentColor" d="M20 12v8H4v-8zm0-2H4a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-8a2 2 0 0 0-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Products(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 2H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1M5 21a2 2 0 1 1 2-2a2 2 0 0 1-2 2m2-9H3V3h4Zm-1 7a1 1 0 1 1-1-1a1 1 0 0 1 1 1m8-17h-4a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1m-2 19a2 2 0 1 1 2-2a2 2 0 0 1-2 2m2-9h-4V3h4Zm-1 7a1 1 0 1 1-1-1a1 1 0 0 1 1 1m8-17h-4a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1m-2 19a2 2 0 1 1 2-2a2 2 0 0 1-2 2m2-9h-4V3h4Zm-1 7a1 1 0 1 1-1-1a1 1 0 0 1 1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ProductsOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="5" cy="19" r="1" fill="currentColor"/><path fill="currentColor" d="M4 4h2v9H4z"/><path fill="currentColor" d="M7 2H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1m0 19H3V3h4Z"/><circle cx="12" cy="19" r="1" fill="currentColor"/><path fill="currentColor" d="M11 4h2v9h-2z"/><path fill="currentColor" d="M14 2h-4a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1m0 19h-4V3h4Z"/><circle cx="19" cy="19" r="1" fill="currentColor"/><path fill="currentColor" d="M18 4h2v9h-2z"/><path fill="currentColor" d="M21 2h-4a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1m0 19h-4V3h4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Project(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 4h-4.18a2.988 2.988 0 0 0-5.64 0H5a2.006 2.006 0 0 0-2 2v14a2.006 2.006 0 0 0 2 2h14a2.006 2.006 0 0 0 2-2V6a2.006 2.006 0 0 0-2-2m-7 0a1 1 0 1 1-1 1a1.003 1.003 0 0 1 1-1m-2 5l2.79 2.794l2.52-2.52L14 8h4v4l-1.276-1.311l-3.932 3.935L10 11.83l-2.586 2.584L6 13Zm9 10H5v-2h14Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ProjectOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="12" cy="6" r="1" fill="currentColor"/><path fill="currentColor" d="M6 17h12v2H6zm4-5.17l2.792 2.794l3.932-3.935L18 12V8h-4l1.31 1.275l-2.519 2.519L10 9l-4 4l1.414 1.414z"/><path fill="currentColor" d="M19 3h-3.298a4.497 4.497 0 0 0-.32-.425l-.01-.012a4.426 4.426 0 0 0-2.89-1.518a2.577 2.577 0 0 0-.964 0a4.426 4.426 0 0 0-2.89 1.518l-.01.012a4.497 4.497 0 0 0-.32.424V3H5a3.003 3.003 0 0 0-3 3v14a3.003 3.003 0 0 0 3 3h14a3.003 3.003 0 0 0 3-3V6a3.003 3.003 0 0 0-3-3m1 17a1 1 0 0 1-1 1H5a1.001 1.001 0 0 1-1-1V6a1.001 1.001 0 0 1 1-1h4.55a2.5 2.5 0 0 1 4.9 0H19a1.001 1.001 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Proxy(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.013 9.395l4.57-4.57L21 6.243V2h-4.243l1.411 1.41l-4.834 4.835a3.938 3.938 0 0 0-5.191 2.75H5.72a2 2 0 1 0 .005 2H8.14a3.94 3.94 0 0 0 5.204 2.757l4.83 4.83L16.758 22H21v-4.243l-1.41 1.411l-4.571-4.57a3.967 3.967 0 0 0 .841-1.603L18 13v2l3-3l-3-3v2l-2.143-.005a3.968 3.968 0 0 0-.844-1.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ProxyOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 6.243V2h-4.243l1.411 1.41l-4.834 4.835a3.938 3.938 0 0 0-5.191 2.75H5.72a2 2 0 1 0 .005 2H8.14a3.94 3.94 0 0 0 5.204 2.757l4.83 4.83L16.758 22H21v-4.243l-1.41 1.411l-4.571-4.57a3.967 3.967 0 0 0 .841-1.603L18 13v2l3-3l-3-3v2l-2.143-.005a3.968 3.968 0 0 0-.844-1.6l4.57-4.57ZM12 14a2 2 0 1 1 2-2a2 2 0 0 1-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PullRequest(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 17.184V8.816a3 3 0 1 0-2 0v8.368a3 3 0 1 0 2 0m14 0V7a2 2 0 0 0-2-2h-5.172l1.586-1.586L13 2l-2.586 2.586L9 6l1.414 1.414L13 10l1.414-1.414L12.828 7H18v10.184a3 3 0 1 0 2 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PullRequestOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 17.184V8.816a3 3 0 1 0-2 0v8.368a3 3 0 1 0 2 0M5 5a1 1 0 1 1-1 1a1 1 0 0 1 1-1m0 16a1 1 0 1 1 1-1a1 1 0 0 1-1 1m15-3.816V7a2 2 0 0 0-2-2h-5.172l1.586-1.586L13 2l-2.586 2.586L9 6l1.414 1.414L13 10l1.414-1.414L12.828 7H18v10.184a3 3 0 1 0 2 0M19 21a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PushPin(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m18 13.04l-1.99-1.992V3.017h1v-2h-10v2h1v8.023h-.009L6 13.04v2h5.01L11 22h2l.01-6.96H18z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PushPinOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.996 12.011A3.005 3.005 0 0 1 16 9V4h1l-.013.011A1.023 1.023 0 0 0 18 3a1.007 1.007 0 0 0-1.015-.988L17 2H7a1 1 0 0 0-.003 2H8l-.004 5A2.991 2.991 0 0 1 5 12v2h6v7l1 1l1-1v-7h6v-2ZM9 12a4.941 4.941 0 0 0 1-3V4h4v5a4.99 4.99 0 0 0 1 3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Quota(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 12V3h-1a10 10 0 1 0 10 10v-1Z"/><path fill="currentColor" d="M14 10V1a9 9 0 0 1 9 9Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuotaOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.93 12A8.054 8.054 0 1 1 12 5.07V3h-1a10 10 0 1 0 10 10v-1Z"/><path fill="currentColor" d="M20.364 3.636A9.003 9.003 0 0 0 14 1v9h9a9.003 9.003 0 0 0-2.636-6.364M16 3.294A7.01 7.01 0 0 1 20.706 8H16Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Replica(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 11h-6a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2v-6a2 2 0 0 0-2-2m0 8h-6v-6h6ZM12 3H6a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h3v-2H6V5h6v4h2V5a2 2 0 0 0-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReplicaSet(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 13h-6a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2v-6a2 2 0 0 0-2-2m0 8h-6v-6h6ZM16 7v4h-2V7H8v6h3v2H8a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2M4.003 8.998h-2V3.002a2 2 0 0 1 2-2h5.994v2H4.003Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Repositories(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8 23.01l2.5-1.5l2.5 1.5V19H8zM8 5h2v2H8zm0 5h2v2H8zm0 5.01h2v2H8z"/><path fill="currentColor" d="M18.36 1.05A1.47 1.47 0 0 0 18 1H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2V3h12v16h-3v2h3a2 2 0 0 0 2-2V3a2 2 0 0 0-1.64-1.95"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RoleBinding(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 21h-2a2 2 0 0 1 0-4h2v-2h-2a4 4 0 0 0 0 8h2Zm8-2a4 4 0 0 1-4 4h-2v-2h2a2 2 0 0 0 0-4h-2v-2h2a4 4 0 0 1 4 4"/><path fill="currentColor" d="M14 18h4v2h-4zm-7 1a5.989 5.989 0 0 1 .09-1H3v-1.4c0-2 4-3.1 6-3.1a8.548 8.548 0 0 1 1.35.125A5.954 5.954 0 0 1 13 13h5V4a2.006 2.006 0 0 0-2-2h-4.18a2.988 2.988 0 0 0-5.64 0H2a2.006 2.006 0 0 0-2 2v14a2.006 2.006 0 0 0 2 2h5.09A5.989 5.989 0 0 1 7 19M9 2a1 1 0 1 1-1 1a1.003 1.003 0 0 1 1-1m0 4a3 3 0 1 1-3 3a2.996 2.996 0 0 1 3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RoleBindingOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23 19a4 4 0 0 1-4 4h-2v-2h2a2 2 0 0 0 0-4h-2v-2h2a4 4 0 0 1 4 4M9 19a4 4 0 0 1 4-4h2v2h-2a2 2 0 0 0 0 4h2v2h-2a4 4 0 0 1-4-4"/><path fill="currentColor" d="M14 18h4v2h-4zM9 5a3 3 0 1 0 3 3a3.009 3.009 0 0 0-3-3m0 4a1 1 0 1 1 1-1a1.003 1.003 0 0 1-1 1m-3.69 6A7.011 7.011 0 0 1 9 13.88a5.641 5.641 0 0 1 .778.064A5.965 5.965 0 0 1 13 13h.254A9.398 9.398 0 0 0 9 11.89c-2.03 0-6 1.07-6 3.58V17h4.349a5.986 5.986 0 0 1 1.188-2Z"/><path fill="currentColor" d="M16 2h-4.18a2.988 2.988 0 0 0-5.64 0H2a2.006 2.006 0 0 0-2 2v14a2.006 2.006 0 0 0 2 2h5.141a3.606 3.606 0 0 1 0-2H2V4h14v9h2V4a2.006 2.006 0 0 0-2-2M9 3.25a.756.756 0 0 1-.75-.75a.75.75 0 0 1 1.5 0a.756.756 0 0 1-.75.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RotatingGear(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.43 12.98c.04-.32.07-.64.07-.98s-.03-.66-.07-.98l2.11-1.65c.19-.15.24-.42.12-.64l-2-3.46c-.12-.22-.39-.3-.61-.22l-2.49 1c-.52-.4-1.08-.73-1.69-.98l-.38-2.65C14.46 2.18 14.25 2 14 2h-4c-.25 0-.46.18-.49.42l-.38 2.65c-.61.25-1.17.59-1.69.98l-2.49-1c-.23-.09-.49 0-.61.22l-2 3.46c-.13.22-.07.49.12.64l2.11 1.65c-.04.32-.07.65-.07.98s.03.66.07.98l-2.11 1.65c-.19.15-.24.42-.12.64l2 3.46c.12.22.39.3.61.22l2.49-1c.52.4 1.08.73 1.69.98l.38 2.65c.03.24.24.42.49.42h4c.25 0 .46-.18.49-.42l.38-2.65c.61-.25 1.17-.59 1.69-.98l2.49 1c.23.09.49 0 .61-.22l2-3.46c.12-.22.07-.49-.12-.64l-2.11-1.65zM12 15.5c-1.93 0-3.5-1.57-3.5-3.5s1.57-3.5 3.5-3.5s3.5 1.57 3.5 3.5s-1.57 3.5-3.5 3.5z"><animateTransform attributeName="transform" attributeType="XML" dur="10s" from="0 12 12" repeatCount="indefinite" to="360 12 12" type="rotate"/></path>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Route(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.016 4.502A2.976 2.976 0 0 0 3.038 7.48c0 2.233 2.978 5.53 2.978 5.53s2.978-3.297 2.978-5.53a2.976 2.976 0 0 0-2.978-2.978m0 4.041A1.063 1.063 0 1 1 7.079 7.48a1.064 1.064 0 0 1-1.063 1.063m15.008 2.753v-4.3a4.962 4.962 0 0 0-.204-1.333a4.996 4.996 0 0 0-9.796 1.216v.248l-.01.87v9.952h-.004v.041a2 2 0 0 1-4 0c0-.012.003-.024.004-.037H7.01V16.01h-2v2h.002a3.998 3.998 0 0 0 7.996-.005h.002v-.982h.005V8.997l.01-1.87V6.88a3.001 3.001 0 0 1 6 .123v4.275a1.999 1.999 0 1 0 2 .018"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sandbox(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="9.54" cy="8.99" r="1" fill="currentColor"/><circle cx="11.04" cy="11.99" r="1" fill="currentColor"/><circle cx="8.04" cy="11.99" r="1" fill="currentColor"/><path fill="currentColor" d="M19 9.24V15h-4a3 3 0 0 1-6 0H5V5h9.77l2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V7.24Z"/><path fill="currentColor" d="m17 10.1l-.84-.84l-.26.26l-1.44-1.41l.26-.26l-.78-.78c-.37-.38-.88-.17-1.15.31l-1.17 1.91a2 2 0 0 1 1.49 1.16a1.54 1.54 0 0 1 .6-.13A1.67 1.67 0 0 1 15.39 12l1.25-.74c.5-.26.71-.78.36-1.16"/><path fill="currentColor" d="m15.87 9.53l.26-.27l5.28-5.27l-1.42-1.42l-5.27 5.28l-.26.26z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SandboxOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="9.537" cy="8.987" r="1" fill="currentColor"/><path fill="currentColor" d="M8.037 10.987a1 1 0 0 0-1 1a.984.984 0 0 0 .482.834a.972.972 0 0 0 .518.166a1 1 0 0 0 1-1a1.016 1.016 0 0 0-1-1"/><path fill="currentColor" d="M3 19a2 2 0 0 0 2 2h14a2.006 2.006 0 0 0 2-2V7.244l-2 2V15h-4a3 3 0 0 1-6 0H5L4.99 5h9.769l2-2H5a1.993 1.993 0 0 0-2 2Zm2 0v-2h2.422a5.103 5.103 0 0 0 1.565 2Zm14 0l-3.987-.033A4.998 4.998 0 0 0 16.578 17H19Z"/><path fill="currentColor" d="M13.943 7.071c-.376-.376-.88-.17-1.157.31l-1.164 1.914a2.08 2.08 0 0 1 1.487 1.15a1.618 1.618 0 0 1 .592-.128l.01-.001a1.703 1.703 0 0 1 1.683 1.67l.018.014l.38-.225l.871-.516c.476-.29.687-.78.31-1.157l-.84-.84l5.274-5.274l-1.415-1.415l-5.273 5.274Z"/><circle cx="11.037" cy="11.987" r="1" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SatelliteAlt(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9.74 23.53l-3.78-3.77a.976.976 0 0 1-.95.75a1.003 1.003 0 0 1-1-1a.987.987 0 0 1 .75-.95l-3.51-3.51a6 6 0 0 1 8.49 8.48"/><path fill="currentColor" d="m22.46 12.45l-2.94-2.94l.86-.86a1.956 1.956 0 0 0 0-2.76L18.9 4.41a1.968 1.968 0 0 0-2.77 0l-.85.85l-2.95-2.94a1.672 1.672 0 0 0-2.36 0L6.2 6.09a1.672 1.672 0 0 0 0 2.36l2.95 2.95l-.15.14a1.95 1.95 0 0 0 0 2.77l.74.74l.74.74a1.955 1.955 0 0 0 2.76 0l.15-.15l2.94 2.94a1.678 1.678 0 0 0 2.37 0l3.76-3.76a1.666 1.666 0 0 0 0-2.37M7.38 7.27l3.77-3.77l2.94 2.95l-3.76 3.76ZM17.52 17.4l-2.95-2.94l3.77-3.77l2.94 2.94Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SatelliteAltOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.465 12.45L19.52 9.506l.855-.855a1.954 1.954 0 0 0 0-2.764l-1.48-1.479a1.954 1.954 0 0 0-2.763 0l-.855.855l-2.943-2.944a1.676 1.676 0 0 0-2.367 0L6.2 6.086a1.676 1.676 0 0 0 0 2.367l2.944 2.943l-.148.148a1.954 1.954 0 0 0 0 2.764l.74.739a6 6 0 0 0-8.485 0l3.511 3.512a.986.986 0 0 0-.75.95a1 1 0 0 0 1 1a.986.986 0 0 0 .95-.75l3.774 3.774a6 6 0 0 0 0-8.485l.74.739a1.954 1.954 0 0 0 2.763 0l.148-.148l2.944 2.944a1.676 1.676 0 0 0 2.366 0l3.767-3.767a1.676 1.676 0 0 0 0-2.366M9.359 20.327l-4.901-4.902a4.006 4.006 0 0 1 4.901 4.902M7.385 7.269l3.766-3.766l2.944 2.944l-3.766 3.766Zm4.447 7.045L10.418 12.9l.115-.115l1.184-1.184l3.766-3.766l1.184-1.184l.822-.822l1.414 1.414l-.822.823l-1.184 1.183l-3.766 3.766l-1.183 1.184Zm5.684 3.086l-2.945-2.944l3.767-3.767l2.943 2.944Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Science(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.68 22H5.32a2.318 2.318 0 0 1-1.96-3.55L4.27 17L9 9.46V3H7V1h10v2h-2v6.46L19.73 17l.91 1.45A2.318 2.318 0 0 1 18.68 22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScienceOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13 10l5.949 9.516a.317.317 0 0 1-.268.484H5.319a.317.317 0 0 1-.266-.487L11 10V3h2ZM7 1v2h2v6.46l-5.641 8.99A2.318 2.318 0 0 0 5.319 22h13.362a2.318 2.318 0 0 0 1.96-3.55l-5.64-8.99L15 3h2V1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Scientist(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 17c-2.67 0-8 1.34-8 4v3h16v-3c0-2.66-5.33-4-8-4"/><circle cx="8" cy="12" r="4" fill="currentColor"/><circle cx="20" cy="5" r="1" fill="currentColor"/><circle cx="21.5" cy="3.5" r=".5" fill="currentColor"/><circle cx="20.5" cy="1.5" r=".5" fill="currentColor"/><path fill="currentColor" d="M16 7v1h1l.003 10.031C17 19.712 18 21 19.5 21s2.5-1.207 2.5-3V8h1V7Zm5 8l-.565.424a1.765 1.765 0 0 1-1.05.326H19V15h-1v-2h1v-1h-1v-2h1V9h-1V8h3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScientistOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="20" cy="5" r="1" fill="currentColor"/><circle cx="21.5" cy="3.5" r=".5" fill="currentColor"/><circle cx="20.5" cy="1.5" r=".5" fill="currentColor"/><path fill="currentColor" d="M16 7v1h1l.003 10.031C17 19.712 18 21 19.5 21s2.5-1.207 2.5-3V8h1V7Zm5 8l-.565.424a1.765 1.765 0 0 1-1.05.326H19V15h-1v-2h1v-1h-1v-2h1V9h-1V8h3ZM8 10a2 2 0 1 1-2 2a2.006 2.006 0 0 1 2-2m0 10c2.7 0 5.8 1.29 6 2H2c.23-.72 3.31-2 6-2M8 8a4 4 0 1 0 4 4a3.999 3.999 0 0 0-4-4m0 10c-2.67 0-8 1.34-8 4v2h16v-2c0-2.66-5.33-4-8-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Secret(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 9.5A2.492 2.492 0 0 0 13.511 8H10.49a2.436 2.436 0 1 0 .46 1h2.101a2.5 2.5 0 1 0 4.95.5M20 5h-3V2a2 2 0 0 0-2-2H9a2 2 0 0 0-2 2v3H4v1h16ZM5 17H4a4 4 0 0 0-4 4v3h8l-1-1l2-2Zm15 0h-1l-4 4l2 2l-1 1h8v-3a4 4 0 0 0-4-4m-9 3h2v4h-2zm1-2l-1 2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SecretOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m4.172 19l2 2l-1 1H2v-1a2.002 2.002 0 0 1 2-2zM5 17H4a4 4 0 0 0-4 4v3h8l-1-1l2-2zm14.828 2H20a2.002 2.002 0 0 1 2 2v1h-3.172l-1-1zM19 17l-4 4l2 2l-1 1h8v-3a4 4 0 0 0-4-4Zm-7 1l-1 2v4h2v-4zm5-13V2a2 2 0 0 0-2-2H9a2 2 0 0 0-2 2v3H4v1h16V5Zm-2-1H9V2h6Zm.5 3a2.488 2.488 0 0 0-1.989 1H10.49a2.436 2.436 0 1 0 .46 1h2.101a2.5 2.5 0 1 0 2.45-2m-7 3.5a1 1 0 1 1 1-1a1 1 0 0 1-1 1m7 0a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SecureData(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.269 15.01a4.282 4.282 0 0 1 .78-2h-4.027v-2h7.994v.001a4.903 4.903 0 0 1 1.998.44V3.012a2.001 2.001 0 0 0-2.001-2.001H3.016a2.002 2.002 0 0 0-2.002 2.001V17.01a2.001 2.001 0 0 0 2.002 2.002h9.997v-.801a3.013 3.013 0 0 1 .268-1.2H7.013v-2Zm4.744-6h-7.994v-2h7.994Zm-12-6h12v2h-12Zm-2.011 14h-2v-2h2Zm0-12h-2v-2h2Zm2.005 2h2v2h-2Zm.003 6v-2h2v2Zm14.803 4.001v-1.5a2.818 2.818 0 0 0-5.6 0v1.5a1.29 1.29 0 0 0-1.2 1.2v3.5a1.31 1.31 0 0 0 1.2 1.3h5.5a1.31 1.31 0 0 0 1.3-1.2v-3.5a1.31 1.31 0 0 0-1.2-1.3m-1.3 0h-3v-1.5a1.375 1.375 0 0 1 1.5-1.3a1.375 1.375 0 0 1 1.5 1.3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SecureDataOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.8 17v-1.5a2.818 2.818 0 0 0-5.6 0V17a1.29 1.29 0 0 0-1.2 1.2v3.5a1.31 1.31 0 0 0 1.2 1.3h5.5a1.31 1.31 0 0 0 1.3-1.2v-3.5a1.31 1.31 0 0 0-1.2-1.3m-4.3-1.5a1.375 1.375 0 0 1 1.5-1.3a1.375 1.375 0 0 1 1.5 1.3V17h-3Z"/><path fill="currentColor" d="M24 11V2a2 2 0 0 0-2-2H2a2 2 0 0 0-2 2v18a2 2 0 0 0 2 2h11v-2H2V2h20v9"/><path fill="currentColor" d="M6 4H4v2h2zm14 0H8v2h12zm-7 12H8v2h5zm-7 0H4v2h2zm4-8H8v2h2zm10 0h-8v2h8zm-5.192 4H12v2h2.2v-.5a2.26 2.26 0 0 1 .608-1.5M10 12H8v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SelfHealing(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.048 11.769A8.187 8.187 0 0 1 5.73 5.722L8.167 8.17V2H2l2.276 2.262A10.271 10.271 0 0 0 10.214 22v-2.077a8.218 8.218 0 0 1-7.166-8.154M23 12.283A10.315 10.315 0 0 0 13.786 2v2.077a8.27 8.27 0 0 1 7.166 8.206a8.074 8.074 0 0 1-2.682 5.995l-2.484-2.448V22h6.19l-2.252-2.262A10.12 10.12 0 0 0 23 12.283"/><path fill="currentColor" d="m12.29 16l-.537-.49c-1.91-1.732-3.172-2.875-3.172-4.277a2.02 2.02 0 0 1 2.04-2.04a2.221 2.221 0 0 1 1.67.775a2.221 2.221 0 0 1 1.669-.775a2.02 2.02 0 0 1 2.04 2.04c0 1.402-1.261 2.545-3.172 4.281Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Service(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 21h2v-6H3v3zm9.46 1l-2.93-.64l-2.93-.69a.73.73 0 0 1-.43-.26a.71.71 0 0 1-.17-.47v-4a.77.77 0 0 1 .17-.47a.8.8 0 0 1 .43-.26l4.78-.59l4.79-.6l.15.56l.15.55a.93.93 0 0 1-.15.52a.9.9 0 0 1-.4.37l-2 .5l-2 .5l1.62.66l1.63.65l2.27-.64l2.21-.69a.78.78 0 0 1 .58.08a.73.73 0 0 1 .36.46l.19.77l.2.76a.77.77 0 0 1-.08.56a.78.78 0 0 1-.46.35l-3.63 1l-3.63 1a1.42 1.42 0 0 1-.36 0a1.1 1.1 0 0 1-.36.02M9 4L7 6L5 8l2 2l2 2l.7-.7l.7-.7l-1.3-1.3L7.8 8l1.3-1.3l1.3-1.3l-.7-.7zm6 0l-.7.7l-.7.7l1.3 1.3L16.2 8l-1.3 1.3l-1.3 1.3l.7.7l.7.7l2-2l2-2l-2-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ServiceInstance(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 21h2.03v-6H3v3zm9.45 1l-2.93-.64l-2.93-.64A.76.76 0 0 1 6 20v-4a.71.71 0 0 1 .17-.47a.75.75 0 0 1 .43-.27l4.78-.59l4.79-.67l.15.55l.15.55a1 1 0 0 1-.56.9l-2 .5l-2 .5l1.62.65l1.62.65l2.28-.64l2.22-.66a.73.73 0 0 1 .58.07a.75.75 0 0 1 .36.47l.19.76l.2.77a.74.74 0 0 1-.54.9l-3.63 1l-3.63 1a1.49 1.49 0 0 1-.36.05a1.55 1.55 0 0 1-.37-.02m5.57-13.5l-3.5-2.25L11.02 4v9l3.5-2.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ServiceInstanceOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.977 19.09l-.388-1.532a.747.747 0 0 0-.356-.466a.764.764 0 0 0-.587-.072l-4.552 1.284L12 17l3.915-1.003a1.88 1.88 0 0 0 .558-.893L16.166 14L8 15H3v6h5l4.455.964a3.194 3.194 0 0 0 .727-.017l7.26-1.954a.755.755 0 0 0 .454-.344a.738.738 0 0 0 .081-.56M6 20H4v-4h2zm6.922.982a.557.557 0 0 1-.138.018a.538.538 0 0 1-.115-.012L8 19.963V16.02l7.329-.908L10 17l4.721 2.232l.317.127l.328-.093l4.316-1.217l.264 1.042ZM9 2v11l8.18-5.5Zm2 3.76l2.6 1.74L11 9.25Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ServiceOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 4L7 6L5 8l2 2l2 2l.7-.7l.7-.7l-1.3-1.3L7.8 8l1.3-1.3l1.3-1.3l-.7-.7zm6 0l-.7.7l-.7.7l1.3 1.3L16.2 8l-1.3 1.3l-1.3 1.3l.7.7l.7.7l2-2l2-2l-2-2zm5.977 15.09l-.388-1.532a.747.747 0 0 0-.356-.466a.764.764 0 0 0-.587-.072l-4.552 1.284L12 17l3.915-1.003a1.88 1.88 0 0 0 .558-.893L16.166 14L8 15H3v6h5l4.455.964a3.194 3.194 0 0 0 .727-.017l7.26-1.954a.755.755 0 0 0 .454-.344a.738.738 0 0 0 .081-.56M6 20H4v-4h2zm6.922.982a.557.557 0 0 1-.138.018a.538.538 0 0 1-.115-.012L8 19.963V16.02l7.329-.908L10 17l4.721 2.232l.317.127l.328-.093l4.316-1.217l.264 1.042Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ServicePlan(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 2H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1M5 21a2 2 0 1 1 2-2a2 2 0 0 1-2 2m2-9H3V3h4Z"/><circle cx="5.01" cy="19.01" r="1" fill="currentColor"/><path fill="currentColor" d="M14 2h8v2.02h-8zm-4 0h2.01v2.02H10zm4 4h8v2.02h-8zm-4 0h2.01v2.02H10zm4 4h8v2.02h-8zm-4 0h2.01v2.02H10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ServicePlanOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 2h8v2h-8zm-4 0h2v2h-2zm4 4h8v2h-8zm-4 0h2v2h-2zm4 4h8v2h-8zm-4 0h2v2h-2z"/><circle cx="5" cy="19" r="1" fill="currentColor"/><path fill="currentColor" d="M4 4h2v9H4z"/><path fill="currentColor" d="M7 2H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1m0 19H3V3h4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SnapshotRollback(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 2a2 2 0 0 0-2 2v4h2V4h4V2Zm12 0v2h4v4h2V4a2 2 0 0 0-2-2Zm0 20v-2h4v-4h2v4a2 2 0 0 1-2 2ZM4 20v-4H2v4a2 2 0 0 0 2 2h4v-2Zm8.37-10.48A6 6 0 0 0 8 11.45L6 9.5v5h5l-2-2a4.48 4.48 0 0 1 7.6 1.5l1.4-.44a6 6 0 0 0-5.63-4.04"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Software(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 18.32A7.06 7.06 0 0 1 11.28 16H3V4h18v2.26a7.08 7.08 0 0 1 2 2.15V4a2 2 0 0 0-2-2H3a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h7v2H8v2h8v-2h-2Z"/><path fill="currentColor" d="M17 6a6 6 0 1 0 6 6a6 6 0 0 0-6-6m0 7.5a1.5 1.5 0 1 1 1.5-1.5a1.5 1.5 0 0 1-1.5 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SoftwareOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 18.315A7.037 7.037 0 0 1 11.263 16H3V4h18v2.264a7.046 7.046 0 0 1 2 2.15V4a2.006 2.006 0 0 0-2-2H3a2.006 2.006 0 0 0-2 2v12a2.006 2.006 0 0 0 2 2h7v2H8v2h8v-2h-2Z"/><path fill="currentColor" d="M17 6a6 6 0 1 0 6 6a5.998 5.998 0 0 0-6-6m0 10a4 4 0 1 1 4-4a4.005 4.005 0 0 1-4 4"/><circle cx="17" cy="12" r="1" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spinner(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 2h2v5h-2zm0 15h2v5h-2zm11-6v2h-5v-2zM7 11v2H2v-2zm11.364-6.778l1.414 1.414l-3.536 3.536l-1.414-1.414zM7.758 14.828l1.414 1.414l-3.536 3.536l-1.414-1.414zm12.02 3.536l-1.414 1.414l-3.536-3.536l1.414-1.414zM9.172 7.758L7.758 9.172L4.222 5.636l1.414-1.414z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func State(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 6v12h5V6Zm4 7h-3V7h3Zm5-12H6a2 2 0 0 0-2 2v18a2.005 2.005 0 0 0 2 2h13a2 2 0 0 0 2-2V3a2.006 2.006 0 0 0-2-2m0 20H6V3h13Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StatefulSet(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 17.998V5H4v12.998A2.001 2.001 0 0 0 6 20h13v-2.002z"/><path fill="currentColor" d="M2.01 21.999V9h-2v12.999a2 2 0 0 0 2 2.001h13v-2.001zM15.5 8.763a9.173 9.173 0 0 1-2.085-.23A6.197 6.197 0 0 1 11.75 7.9v3.345h.001c0 .308.366.685 1.017.985a6.645 6.645 0 0 0 2.733.524a6.645 6.645 0 0 0 2.733-.524c.651-.3 1.017-.677 1.017-.985h.002V7.9a6.197 6.197 0 0 1-1.666.634a9.172 9.172 0 0 1-2.084.229zm.008-4.517h-.016a6.642 6.642 0 0 0-2.729.525c-.65.3-1.015.676-1.015.984c0 .307.365.683 1.015.983a6.643 6.643 0 0 0 2.73.525h.015a6.643 6.643 0 0 0 2.729-.525c.65-.3 1.015-.676 1.015-.983c0-.308-.365-.684-1.015-.984a6.642 6.642 0 0 0-2.73-.525"/><path fill="currentColor" d="M21 1H10a2 2 0 0 0-2 2v11a2 2 0 0 0 2 2h11a2 2 0 0 0 2-2V3a2 2 0 0 0-2-2m-.248 4.8v5.445h-.002a2.569 2.569 0 0 1-1.5 2.146a7.658 7.658 0 0 1-3.742.863h-.016a7.658 7.658 0 0 1-3.741-.863a2.569 2.569 0 0 1-1.5-2.146h-.002v-5.49a2.57 2.57 0 0 1 1.502-2.148a7.664 7.664 0 0 1 3.748-.861h.003a7.664 7.664 0 0 1 3.747.861a2.57 2.57 0 0 1 1.503 2.148z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StatefulSetOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 16V4H6v12v-.002A2.151 2.151 0 0 0 8 18h13v-2Z"/><path fill="currentColor" d="M4 20V8H2v12a2.149 2.149 0 0 0 2 2h13v-2Zm7.002-9.137a2.679 2.679 0 0 0 1.571 2.237a8.05 8.05 0 0 0 3.919.9h.016a8.049 8.049 0 0 0 3.919-.9a2.679 2.679 0 0 0 1.571-2.237H22V5.137a2.68 2.68 0 0 0-1.574-2.24A8.055 8.055 0 0 0 16.502 2h-.004a8.055 8.055 0 0 0-3.924.898A2.68 2.68 0 0 0 11 5.138v5.725zm2.632-6.752a6.983 6.983 0 0 1 2.858-.547h.016a6.982 6.982 0 0 1 2.858.547c.681.313 1.063.706 1.063 1.026c0 .32-.382.713-1.063 1.026a6.982 6.982 0 0 1-2.858.547h-.016a6.983 6.983 0 0 1-2.858-.547c-.681-.313-1.063-.705-1.063-1.026c0-.32.382-.713 1.063-1.026M12.571 9.12V7.374a6.508 6.508 0 0 0 1.744.66a9.647 9.647 0 0 0 2.183.24h.003a9.647 9.647 0 0 0 2.184-.24a6.507 6.507 0 0 0 1.744-.66v3.489h-.002c0 .32-.382.714-1.065 1.027a6.985 6.985 0 0 1-2.862.546a6.986 6.986 0 0 1-2.862-.546c-.682-.313-1.065-.707-1.065-1.027h-.002Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StorageClass(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 4H2v4h20ZM6 7H4V5h2Zm14.6 7.8A2.088 2.088 0 0 0 19 14h-7.002a.998.998 0 0 0-.998.998v6.004a.998.998 0 0 0 .998.998H19a1.816 1.816 0 0 0 1.6-.8L23 18ZM11.998 12H19a4.075 4.075 0 0 1 3 1.39V10H2v4h7.184a2.993 2.993 0 0 1 2.814-2M4 13v-2h2v2Zm5 3H2v4h7Zm-5 3v-2h2v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StorageClassOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m23 18l-2.4 3.2a1.816 1.816 0 0 1-1.6.8h-7a1.003 1.003 0 0 1-1-1v-6a1.003 1.003 0 0 1 1-1h7a2.088 2.088 0 0 1 1.6.8ZM21 6H3V2h18ZM2 1v6h20V1"/><path fill="currentColor" d="M4 3h2v2H4zM3 14v-4h18v3.33l1 1.337V9H2v6h7v-1z"/><path fill="currentColor" d="M4 11h2v2H4zm0 8h2v2H4z"/><path fill="currentColor" d="M3 22v-4h6v-1H2v6h7v-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SubscriptionManagement(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 19h5v4l-2.5-1.5L6 23zM20 1h-8v22l2.5-1.5L17 23v-2h3a2 2 0 0 0 2-2V3a2 2 0 0 0-2-2m-6 4h2v2h-2Zm0 5h2v2h-2Zm0 5h2v2h-2ZM4 1a2.006 2.006 0 0 0-2 2v16a2.006 2.006 0 0 0 2 2V3h7V1Z"/><path fill="currentColor" d="M6 5h2v2H6zm0 5h2v2H6zm0 5h2v2H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SubscriptionsCreated(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.01 6.01h16v2h-16zm2-4h12v2h-12zM20 10H4a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-8a2 2 0 0 0-2-2m-9.7 10L7 16.76l1.4-1.4l1.9 1.9l5.3-5.3l1.4 1.4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SubscriptionsCreatedOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 8H4V6h16Zm-2-6H6v2h12Zm-7.688 19.1l-3.3-3.3l1.4-1.4l1.9 1.9l5.3-5.3l1.4 1.4Z"/><path fill="currentColor" d="M22 10H2a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h20a2 2 0 0 0 2-2V12a2 2 0 0 0-2-2m0 12H2V12h20Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Symlink(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8Zm-1 17.47v-2.19c-2.78 0-4.61.85-6 2.72c.56-2.67 2.11-5.33 6-5.87V12l4 3.73ZM13 9V3.5L18.51 9Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SymlinkOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13 12l4 3.5l-4 3.5v-2a6.956 6.956 0 0 0-6 3c.56-2.67 2.11-5.46 6-6Zm0-8l5 5h-5Z"/><path fill="currentColor" d="M14.68 0H4a2.006 2.006 0 0 0-2 2v20a2.006 2.006 0 0 0 2 2h16a2.006 2.006 0 0 0 2-2V7.15ZM20 22H4V2h9.83L20 8Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SystemGroup(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 1H3a2 2 0 0 0-2 2v9h2V3h14Z"/><path fill="currentColor" d="M21 5H7a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h5.53v1.53H11V22h6v-1.48h-1.52V19H21a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2m0 12H7V7h14Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SystemImage(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 3H6a2 2 0 0 0-2 2v15a2 2 0 0 0 2 2h11a2 2 0 0 0 1.41-.58A2 2 0 0 0 19 20V2ZM6 6h8v2H6Zm4 14H8v-2h2Zm0-6h4v2h-4Zm7 6h-6v-2h6Zm0-4h-2v-2h2Zm0-4h-5v-2h5Zm0-4h-2V6h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SystemImageOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 6v2H7V6Zm-2 4v2h6v-2Zm-3 4v2h5v-2Zm3 4v2h6v-2Zm-4 0v2h2v-2Zm8-4v2h2v-2Zm0-8v2h2V6Zm4 16H5V4h14Zm0-20H5a2 2 0 0 0-2 2v18a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SystemOk(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.004 2.004h-18a2.006 2.006 0 0 0-2 2v12a2.006 2.006 0 0 0 2 2h7v2h-2v2h8v-2h-2v-2h7a2.006 2.006 0 0 0 2-2v-12a2.006 2.006 0 0 0-2-2M10.756 13.67l-4.75-4.75L7.42 7.507l3.336 3.336l5.835-5.836l1.415 1.415Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SystemOkOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.005 8.92L7.42 7.508l3.336 3.336l5.835-5.836l1.415 1.415l-7.25 7.248Z"/><path fill="currentColor" d="M21 2H3a2.006 2.006 0 0 0-2 2l.004 12.004A2 2 0 0 0 3 18h7v2l-1.996.004L8 22h8l.004-1.996L14 20v-2h7a2.006 2.006 0 0 0 2-2V4a2.006 2.006 0 0 0-2-2m0 14H3V4h18Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SystemReRegistered(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2H3a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h7v2H8v2h8v-2h-2v-2h7a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2m0 7h-7v2h7v5H3v-5h5v3l4-4l-4-4v3H3V4h18Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SystemWarning(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2H3a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h7v2H8v2h8v-2h-2v-2h7a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2m0 14H3V4h18Z"/><path fill="currentColor" d="M11.01 12.99h2v2h-2zm0-8h2v6h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Templates(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 2H5a2.005 2.005 0 0 0-2 2v16a2.005 2.005 0 0 0 2 2h14a2.005 2.005 0 0 0 2-2V4a2.005 2.005 0 0 0-2-2m0 11.152v5.696L14 16Zm-7 1.694L7 12h10ZM5 4h14v2H5Zm0 4h8v2H5Zm5 8l-5 2.848v-5.704Zm2 1.15L17 20H7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TemplatesOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 0H4a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V2a2 2 0 0 0-2-2m0 22H4V2h16Z"/><path fill="currentColor" d="M6 4h12v2H6zm0 4h7v2H6zm2 12h8l-4-3zm11-1v-6l-5 3zM6 13v6l4-3zm10-1H8l4 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Terminal(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 5H4a2.006 2.006 0 0 0-2 2v10a2.006 2.006 0 0 0 2 2h16a2.006 2.006 0 0 0 2-2V7a2.006 2.006 0 0 0-2-2M6 17l-1.408-1.412L8.17 12L4.594 8.414L6 7l5 5Zm13 0h-7v-2h7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TerminalOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 15h7v2h-7zM4.594 8.414L8.17 12l-3.578 3.588L6 17l5-5l-5-5z"/><path fill="currentColor" d="M22 3H2a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h20a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2m0 16H2V5h20Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TestTube(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="13" cy="11" r="1" fill="currentColor"/><circle cx="11.5" cy="7.5" r=".5" fill="currentColor"/><circle cx="12.5" cy="5.5" r=".5" fill="currentColor"/><path fill="currentColor" d="M5 0v2h2v17a5 5 0 0 0 10 0V2h2V0Zm10 2v13h-3a1 1 0 0 0-2 0H9V2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThreeDprint(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 2v20H2V2zm0-2H2a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h20a2 2 0 0 0 2-2V2a2 2 0 0 0-2-2"/><path fill="currentColor" d="M18 17H6v2h12zm-5 2h-2v3h2z"/><path fill="currentColor" fill-rule="evenodd" d="M9.327 7.089A1.99 1.99 0 0 1 9.272 7H2V5h7.266A1.995 1.995 0 0 1 11 4h2a2.002 2.002 0 0 1 1.731 1H23v2h-8.272a1.997 1.997 0 0 1-1.186.916L12 11l-1.542-3.084a1.995 1.995 0 0 1-1.131-.827M12 7h-1.004h2.008zm1-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThreeDotsLoading(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="18" cy="12" r="0" fill="currentColor"><animate attributeName="r" begin=".67" calcMode="spline" dur="1.5s" keySplines="0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8" repeatCount="indefinite" values="0;2;0;0"/></circle><circle cx="12" cy="12" r="0" fill="currentColor"><animate attributeName="r" begin=".33" calcMode="spline" dur="1.5s" keySplines="0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8" repeatCount="indefinite" values="0;2;0;0"/></circle><circle cx="6" cy="12" r="0" fill="currentColor"><animate attributeName="r" begin="0" calcMode="spline" dur="1.5s" keySplines="0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8" repeatCount="indefinite" values="0;2;0;0"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Timeout(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.774 5.215L6.326 3.767A9.995 9.995 0 0 1 20.23 17.67l-1.445-1.445a7.99 7.99 0 0 0-11.01-11.01m12.588 15.149l-1.279 1.279l-1.413-1.414A9.995 9.995 0 0 1 3.768 6.327l-1.41-1.41l1.278-1.28l1.29 1.292l1.417 1.413L11 11.002V7h1.5v5.25l4.5 2.67l-.75 1.23l-.254-.152Zm-4.136-1.58l-5.793-5.792l-5.218-5.218a7.99 7.99 0 0 0 11.01 11.01"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Troubleshooting(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23 11h-1.05A10.017 10.017 0 0 0 13 2.05V0h-2v2.05A10.017 10.017 0 0 0 2.05 11H0v2h2.05A10.017 10.017 0 0 0 11 21.95V24h2v-2.05A10.017 10.017 0 0 0 21.95 13H24v-2Zm-9.277 0a1.995 1.995 0 0 0-.723-.723V7.881A4.248 4.248 0 0 1 16.119 11Zm2.396 2A4.248 4.248 0 0 1 13 16.119v-2.396a1.995 1.995 0 0 0 .723-.723ZM11 4.07v2.273A5.757 5.757 0 0 0 6.343 11H4.069A8.007 8.007 0 0 1 11 4.07M10.277 11H7.881A4.249 4.249 0 0 1 11 7.881v2.396a1.995 1.995 0 0 0-.723.723M11 13.723v2.396A4.248 4.248 0 0 1 7.881 13h2.396a1.995 1.995 0 0 0 .723.723M4.07 13h2.273A5.757 5.757 0 0 0 11 17.657v2.274A8.007 8.007 0 0 1 4.07 13M13 19.93v-2.273A5.757 5.757 0 0 0 17.657 13h2.274A8.007 8.007 0 0 1 13 19.93M17.657 11A5.757 5.757 0 0 0 13 6.343V4.069A8.007 8.007 0 0 1 19.93 11Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrustedOrganization(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 9a7 7 0 1 0-10.974 5.76L5 20l2.256.093L8.464 22l3.466-6.004c.024 0 .046.004.07.004s.046-.003.07-.004L15.536 22l1.232-1.866L19 20l-3.026-5.24A6.99 6.99 0 0 0 19 9M7 9a5 5 0 1 1 5 5a5 5 0 0 1-5-5"/><circle cx="12" cy="9" r="3" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Typing(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 2H4c-1.1 0-2 .9-2 2v18l4-4h14c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2zm0 14H6l-2 2V4h16v12z"/><circle cx="16" cy="10" r="0" fill="currentColor"><animate attributeName="r" begin=".67" calcMode="spline" dur="1.5s" keySplines="0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8" repeatCount="indefinite" values="0;1.75;0;0"/></circle><circle cx="12" cy="10" r="0" fill="currentColor"><animate attributeName="r" begin=".33" calcMode="spline" dur="1.5s" keySplines="0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8" repeatCount="indefinite" values="0;1.75;0;0"/></circle><circle cx="8" cy="10" r="0" fill="currentColor"><animate attributeName="r" begin="0" calcMode="spline" dur="1.5s" keySplines="0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8;0.2 0.2 0.4 0.8" repeatCount="indefinite" values="0;1.75;0;0"/></circle>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VirtualGuest(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2H3a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h7v2H8v2h8v-2h-2v-2h7a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2m0 14H3V4h18Z"/><path fill="currentColor" d="M6 7v6a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1m11 6H7V8h10Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VirtualHostManager(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.312 18.034L17.297 7.559a1.994 1.994 0 1 0-3.005-2.466L5.984 3.614A1.998 1.998 0 1 0 3 5.7v12.578a1.992 1.992 0 0 0 .694 3.692A1.989 1.989 0 0 0 4 22a1.992 1.992 0 0 0 1.717-.99h12.546a1.99 1.99 0 0 0 3.699-.834c.005-.059.018-.116.018-.176a1.996 1.996 0 0 0-1.668-1.966m-14.538 1.06a1.988 1.988 0 0 0-.172-.28l3.845-4.917a3.512 3.512 0 0 0 .688.07a3.463 3.463 0 0 0 .949-.147l1.242 1.478A2.472 2.472 0 0 0 12 16.51c0 .038.01.073.011.11a2.494 2.494 0 0 0 2.489 2.39a2.468 2.468 0 0 0 1.664-.647l1.903 1.099c-.006.018-.017.035-.022.054H5.935a2.104 2.104 0 0 0-.161-.422m9.225-11.422a1.826 1.826 0 0 0 1.35.288l2.93 10.173a2.028 2.028 0 0 0-.279.141l-2.078-1.2a2.507 2.507 0 0 0 .065-.435c.002-.043.013-.084.013-.129a2.49 2.49 0 0 0-3.995-1.992l-.95-1.13a3.485 3.485 0 0 0 1.285-4.322ZM5.904 4.638l8.178 1.455a1.977 1.977 0 0 0 .228.824L12.782 8.2a3.446 3.446 0 0 0-4.477-.705zm-1.882 1.35a1.982 1.982 0 0 0 1.282-.476l2.221 2.643a3.46 3.46 0 0 0 .923 5.36l-3.642 4.657a1.99 1.99 0 0 0-.784-.17Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VirtualReality(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 10h2v4h-2zM1 10h2v4H1zm17-1a2 2 0 0 0-2-2H8a2 2 0 0 0-2 2a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2a2 2 0 0 0 2 2h1.028l1.602-2.776a1.582 1.582 0 0 1 2.74 0L14.972 17H16a2 2 0 0 0 2-2a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2m-1 3H7v-2h10Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VirtualRealityOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.826 10a3.012 3.012 0 0 0-1.955-1.871A3.005 3.005 0 0 0 16 6H8a3.005 3.005 0 0 0-2.871 2.129A3.011 3.011 0 0 0 3.174 10H2v4h1.174a3.011 3.011 0 0 0 1.955 1.871A3.006 3.006 0 0 0 8 18h1.605l1.891-3.275a.582.582 0 0 1 1.008 0L14.396 18H16a3.006 3.006 0 0 0 2.871-2.129A3.012 3.012 0 0 0 20.826 14H22v-4ZM19 13a1 1 0 0 1-1 1h-1v1a1 1 0 0 1-1 1h-.45l-1.314-2.275a2.582 2.582 0 0 0-4.472 0L8.45 16H8a1.001 1.001 0 0 1-1-1v-1H6a1.001 1.001 0 0 1-1-1v-2a1.001 1.001 0 0 1 1-1h1V9a1.001 1.001 0 0 1 1-1h8a1.001 1.001 0 0 1 1 1v1h1a1.001 1.001 0 0 1 1 1Z"/><path fill="currentColor" d="M15 11H9v-1h6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VirtualSpace(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8"/><circle cx="11.999" cy="7.493" r="1.493" fill="currentColor"/><circle cx="11.999" cy="16.493" r="1.493" fill="currentColor"/><circle cx="16.496" cy="12" r="1.493" fill="currentColor"/><circle cx="7.496" cy="12" r="1.493" fill="currentColor"/><circle cx="11.999" cy="12" r="1" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Volume(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2H7a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2M7 3a1 1 0 1 1-1 1a1 1 0 0 1 1-1m0 18a1 1 0 1 1 1-1a1 1 0 0 1-1 1m5-7a4.973 4.973 0 0 1-2-.422V11H7.422A4.997 4.997 0 1 1 12 14m5 7a1 1 0 1 1 1-1a1 1 0 0 1-1 1m0-16a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/><circle cx="12" cy="9" r="2" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeBinding(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 19a6 6 0 0 1 6-6h4V3a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v16.001A2 2 0 0 0 4 21h2.349A5.976 5.976 0 0 1 6 19m8-17a1 1 0 1 1-1 1a1 1 0 0 1 1-1M9 3a5 5 0 1 1-2 9.578V10H4.422A4.991 4.991 0 0 1 9 3M4 2a1 1 0 1 1-1 1a1 1 0 0 1 1-1m0 18a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/><circle cx="9" cy="8" r="2" fill="currentColor"/><path fill="currentColor" d="M13 18h4v2h-4z"/><path fill="currentColor" d="M18 15h-2v2h2a2 2 0 0 1 0 4h-2v2h2a4 4 0 0 0 0-8m-4 6h-2a2 2 0 0 1 0-4h2v-2h-2a4 4 0 0 0 0 8h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeBindingOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 18h4v2h-4z"/><path fill="currentColor" d="M19 15h-2v2h2a2 2 0 0 1 0 4h-2v2h2a4 4 0 0 0 0-8m-6 2h2v-2h-2a4 4 0 0 0 0 8h2v-2h-2a2 2 0 0 1 0-4M6 4a1 1 0 1 1-1-1a1 1 0 0 1 1 1m10 0a1 1 0 1 1-1-1a1 1 0 0 1 1 1M6 20a1 1 0 1 1-1-1a1 1 0 0 1 1 1"/><path fill="currentColor" d="M9.69 24H3a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v11h-2V2H3v20h4.812a6.04 6.04 0 0 0 1.879 2"/><path fill="currentColor" d="M14.578 13A4.998 4.998 0 1 0 6 14h3v.54A5.969 5.969 0 0 1 13 13ZM10 13a2 2 0 1 1 2-2a2 2 0 0 1-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeOutlined(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 22H5V2h14Zm0-22H5a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V2a2 2 0 0 0-2-2"/><path fill="currentColor" d="M8 4a1 1 0 1 1-1-1a1 1 0 0 1 1 1m10 0a1 1 0 1 1-1-1a1 1 0 0 1 1 1M8 20a1 1 0 1 1-1-1a1 1 0 0 1 1 1m10 0a1 1 0 1 1-1-1a1 1 0 0 1 1 1M12 6a5 5 0 0 0-5 5a4.936 4.936 0 0 0 1 3h3v1.9a5.001 5.001 0 0 0 6-4.9c0-5-5-5-5-5m0 7a2 2 0 1 1 2-2a2 2 0 0 1-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Workload(children ...ElementRenderer) *EosIconsIcon {
	return &EosIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 6H2v14a2 2 0 0 0 2 2h14v-2H4Z"/><path fill="currentColor" d="M10 10h8v2h-8zm0 3h8v2h-8zm3.47-8.01h-2.31L10 7l1.16 2h2.31l1.16-2z"/><path fill="currentColor" d="M20 2H8a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2m0 14H8V4h12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
