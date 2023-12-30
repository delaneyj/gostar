package topcoat

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = "0.2.0"
	hAttr          = "1em"
	viewbox        = "0 0 42 42"
	fill           = "currentColor"
)

type TopcoatIcon struct {
	*SVGSVGElement
}

func Alert(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.295 3.895L1.203 34.555C-.219 37.146.385 39.5 4.228 39.5H36.77c3.854 0 4.447-2.354 3.025-4.945L22.35 3.914c-.354-.691-.868-1.424-1.957-1.414c-1.16.021-1.735.703-2.098 1.395M18.5 13.5h4v14h-4zm0 17h4v4h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDown(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M25.5 2.5h-10c-2.53 0-3 .529-3 3v15H.5l20 20l20-20h-12v-15c0-2.439-.5-3-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeft(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M39.5 26.5v-10c0-2.529-.508-3-2.979-3H21.5v-12l-20 20l20 20v-12h15.021c2.44 0 2.979-.5 2.979-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRight(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.5 26.5v-10c0-2.529.529-3 3-3h15v-12l20 20l-20 20v-12h-15c-2.439 0-3-.5-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUp(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.5 40.5h10c2.529 0 3-.529 3-3v-15h12l-20-20l-20 20h12v15c0 2.439.5 3 3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Attachment(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.938 23.5c-.398.74-.16 1.35.66 2.17l.76.76c.83.83 1.452 1.061 2.171.66c.261-.14.541-.359.841-.66l4.521-4.52c.846-.734 1.718-.43 2.538.39c.84.83 1.35 1.819.59 2.72l-6.778 6.781c-2.41 2.41-6.062 2.98-9.04 0c-2.682-2.681-2.41-6.631 0-9.041L22.758 9.198c2.41-2.411 5.98-3.06 9.041 0c3.05 3.051 2.601 6.441.75 8.292c-1.17 1.17-1.041 2.54-.311 3.34l.75.749c1.141 1.13 2.181.78 3.301-.359c3.84-4.601 3.601-11.461-.72-15.792c-4.581-4.571-12.001-4.571-16.58 0L5.428 18.99c-4.57 4.58-4.57 12.001 0 16.581c4.58 4.571 12.001 4.571 16.581 0l6.78-6.781c2.049-3.11 1.76-7.15-.971-9.88c-2.721-2.721-6.6-2.781-9.699-.771l-4.521 4.521c-.299.3-.52.58-.66.84"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Audio(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M27.5 9.045a13.391 13.391 0 0 1 8.37 12.425a13.39 13.39 0 0 1-8.37 12.424v2.795a16.007 16.007 0 0 0 11-15.219c0-7.099-4.609-13.125-11-15.228zm0 6.599a7.644 7.644 0 0 1 2.68 5.827a7.64 7.64 0 0 1-2.68 5.826v2.844c2.99-1.731 5-4.966 5-8.67a10.01 10.01 0 0 0-5-8.67zm-23 13.835h8l9 11.015c1 1.44 3.34 1.331 4-.302V2.83c-.811-1.632-2.939-1.763-4-.382l-9 11.053h-8c-2.561 0-3 .461-3 2.964v10.012c0 2.442.5 3.002 3 3.002"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Audiooff(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 29.479h8l9 11.015c1 1.44 3.34 1.331 4-.302V2.83c-.811-1.632-2.939-1.763-4-.382l-9 11.053h-8c-2.561 0-3 .461-3 2.964v10.012c0 2.442.5 3.002 3 3.002"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Back(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M27.066 1L7 21.068l19.568 19.569l4.934-4.933l-14.637-14.636L32 5.933z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BackLight(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M31 38.32L13.391 21L31 3.68L28.279 1L8 21.01L28.279 41z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Behance(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.5 1.5v40h40v-40zm6 17h10.392c1.041 0 2.136.171 2.974.942a4.095 4.095 0 0 1 1.248 2.908c-.005.897-.201 1.875-.797 2.7c-.419.562-.817.927-1.817 1.169v.096c1 .189 1.531.597 2.064 1.267c.635.817.864 1.924.881 3.012c-.006 1.154-.399 2.565-1.294 3.481c-.795.779-1.928 1.426-3.13 1.426H6.5zm18 0h8v2h-8zm3.671 3.211c1.056 0 1.993.137 2.81.412c.826.268 1.521.686 2.084 1.252c.562.559.99 1.309 1.278 2.175c.382 1.136.454 2.95.446 3.95H25.99c0 1 .181 1.575.543 2.077c.367.493.973.722 1.811.722c.799-.004 1.731-.206 2.039-1.101c.08-.234.119-.698.119-.698h4.287c-.008 1-.447 2.652-1.596 3.656c-1.271 1.031-3.045 1.295-4.752 1.328a9.51 9.51 0 0 1-2.863-.416a5.57 5.57 0 0 1-2.158-1.273c-.594-.566-1.049-1.283-1.367-2.143c-.318-.865-.479-1.894-.479-3.074c0-1.15.154-2.155.467-3.016c.311-.858.752-1.572 1.322-2.14a5.49 5.49 0 0 1 2.082-1.289c.819-.281 1.728-.422 2.726-.422M11.5 21.5v3h3.494c.305 0 .777-.043.997-.289c.213-.26.509-.502.509-.859v-.244c0-.323-.311-.799-.521-1.076c-.192-.227-.662-.531-.985-.531zm16.844 3.369c-.744 0-1.301.408-1.67.805c-.361.396-.582.826-.66 1.826h4.315c0-1-.173-1.438-.521-1.826c-.346-.397-.833-.805-1.464-.805M11.5 28.5v3h3.852c.301 0 .592-.309.816-.545c.213-.258.332-.77.332-1.127v-.256c0-.326-.111-.534-.332-.815c-.193-.229-.495-.257-.816-.257z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookmark(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.791 2.5c-2.71 0-3.29.63-3.29 3.35v31.589c0 2.182-.07 3.021 1.08 3.062c.779 0 1.52-.681 2.12-1.29c0 0 7.269-8.55 7.749-8.729c.279-.11-.011-.131-.021-.08c-.01.01.021.039.12.08c.48.18 7.75 8.729 7.75 8.729c.6.609 1.34 1.29 2.119 1.29c1.15-.04 1.08-.88 1.08-3.062V5.85c0-2.72-.58-3.35-3.289-3.35z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Brush(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.049 14.741c3.58.43 8.219 4.52 8.6 8.28c0 0 9.449-9.471 9.439-9.5c1.939-1.94 1.87-2.75-.15-4.561l-4.08-3.68c-1.948-1.76-2.698-1.69-4.528.13zm4.95 8.321c-.3-1.351-2.929-4.512-5.759-4.98c-7.1-1.83-13.89 5.88-13.49 11.491c.641 4.581-5.35 5.22-6.25 5.43c5.92 2.062 8.439 2.541 13.919 1.38c4.78-1.009 14.03-5.71 11.58-13.321"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Build(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.5 6.5c0-2.31-.439-3-3-3h-16c-2.56 0-3 .69-3 3v29c0 2.5.62 3 3 3h10v-2l9-8zm-17 2h12v24h-12zm7 28h-3v-2h3zm13.561-8.391L18.05 35.4c-1.04 1.108-2.7 2.01-1.16 3.898c1.28 1.57 2.969.12 3.93-.898l7.99-7.41zm12.048-3.048l-2.27 1.959l-1.91-1.97l.871-2.15l-2.14-2.682s-4.16-3.31-9.36-.05c2.341-.109 4.58 1.181 6.2 4.021l-2.189 2.121l2.75 2.629l1.799-1.619l2.201 1.93l-1.761 1.68l2.14 2.16l5.761-5.58z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M37.499 5.5H35.5v5.52c0 1.23-.234 1.48-1.484 1.48h-6.017c-1.25 0-1.499-.271-1.499-1.5V5.5h-11v5.52c0 1.23-.25 1.48-1.5 1.48H8c-1.25 0-1.5-.271-1.5-1.5V5.5h-2c-2.58 0-3 .561-3 3v28.998c-.01 2.43.6 3.002 3 3.002h32.999c2.55 0 3.001-.531 3.001-3.002V8.5c0-2.35-.381-3-3.001-3m-1.999 30h-29v-20h29z"/><path fill="currentColor" d="M10 10.5h2c1.25 0 1.5-.25 1.5-1.48v-5c0-1.25-.27-1.5-1.5-1.5l-2-.02c-1.189 0-1.5.28-1.5 1.5v5c0 1.229.25 1.5 1.5 1.5m19.999 0h2c1.25 0 1.501-.25 1.501-1.48v-5c0-1.25-.27-1.5-1.5-1.5l-2-.02c-1.189 0-1.5.28-1.5 1.5v5c0 1.229.249 1.5 1.499 1.5m-19.499 9h5v5h-5zm8 0h5v5h-5zm8 0h5v5h-5zm-16 8h5v5h-5zm8 0h5v5h-5zm8 0h5v5h-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Call(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.562 20.766c-1.328-1.922-2.118-4.241-2.281-4.438c1.945-1.356 5.749-3.06 5.962-5.505c.271-3.159-5.081-9.763-6.107-9.823c-2.808.03-7.947 4.782-8.556 6.218c-1.132 2.969-.571 5.732 1.375 9.732c2.478 5.95 11.682 17.237 16.947 20.78c3.484 2.674 6.029 3.724 9.068 3.09c1.413-.268 6.516-4.455 7.027-7.286c.125-1.05-5.807-8.011-8.875-8.287c-2.382-.22-4.666 3.346-6.303 5.089c-.163-.208-1.559-1.297-3.057-3.021c-1.95-2.049-3.762-4.456-5.2-6.549"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Camera(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.41 22.58c0 3.359 2.73 6.09 6.09 6.09c3.359 0 6.09-2.73 6.09-6.09s-2.73-6.09-6.09-6.09a6.095 6.095 0 0 0-6.09 6.09M3.5 36.5h34c2.63 0 3-.37 3-3v-23c0-2.462-.38-3-3-3h-10c0-2.57-.42-3-3-3h-8c-2.55 0-3 .48-3 3h-10c-2.58 0-3 .692-3 3v23c0 2.6.38 3 3 3m7.64-13.92c0-5.17 4.19-9.359 9.36-9.359s9.359 4.189 9.359 9.359s-4.189 9.359-9.359 9.359s-9.36-4.189-9.36-9.359"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cancel(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m21.002 26.588l10.357 10.604c1.039 1.072 1.715 1.083 2.773 0l2.078-2.128c1.018-1.042 1.087-1.726 0-2.839L25.245 21L36.211 9.775c1.027-1.055 1.047-1.767 0-2.84l-2.078-2.127c-1.078-1.104-1.744-1.053-2.773 0L21.002 15.412L10.645 4.809c-1.029-1.053-1.695-1.104-2.773 0L5.794 6.936c-1.048 1.073-1.029 1.785 0 2.84L16.759 21L5.794 32.225c-1.087 1.113-1.029 1.797 0 2.839l2.077 2.128c1.049 1.083 1.725 1.072 2.773 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cart(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M40.5 12.5c0-1.48-.311-2-1.872-2H11.726l-.801-5c-.109-1.46-.85-2-2.421-2H2.501C1.02 3.5.5 3.99.5 5.5v1c0 1.551.52 2 2.001 2h3.722l3.282 19c.35 1.04 1.311 1.95 3.001 2h22.012c1.75 0 2.57-.359 3.002-2zm-7.023 12H13.696l-1.471-9h22.951zm-19.97 12a4 4 0 0 0 4.002 4a4 4 0 1 0 0-8a4 4 0 0 0-4.002 4m13.007 0a4 4 0 0 0 4.002 4a4 4 0 1 0 0-8a4 4 0 0 0-4.002 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chat(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.5 30.49v8.996c0 1.217.938 1.045 1.25.795c.384-.308 10.75-9.791 10.75-9.791h13c2.471 0 3-.529 3-2.999V6.859c0-2.899-.5-3.359-3-3.359h-32c-2.46 0-3 .5-3 2.999v20.992c0 2.429.42 2.999 3 2.999zm2.26-11.496c0 1.52-1.24 2.759-2.76 2.759s-2.76-1.239-2.76-2.759c0-1.519 1.24-2.759 2.76-2.759s2.76 1.241 2.76 2.759m9 0c0 1.52-1.24 2.759-2.76 2.759a2.765 2.765 0 0 1-2.76-2.759A2.765 2.765 0 0 1 20 16.235c1.52 0 2.76 1.241 2.76 2.759m9 0c0 1.52-1.24 2.759-2.76 2.759s-2.76-1.239-2.76-2.759c0-1.519 1.24-2.759 2.76-2.759s2.76 1.241 2.76 2.759"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Checkmark(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m39.04 7.604l-2.398-1.93c-1.182-.95-1.869-.939-2.881.311L16.332 27.494l-8.111-6.739c-1.119-.94-1.819-.89-2.739.26l-1.851 2.41c-.939 1.182-.819 1.853.291 2.78l11.56 9.562c1.19 1 1.86.897 2.78-.222l21.079-25.061c.99-1.19.93-1.901-.301-2.88"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Circle(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.5 5.5C9.454 5.5.5 12.887.5 22s8.954 16.5 20 16.5s20-7.387 20-16.5s-8.954-16.5-20-16.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleOutline(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.5 4.5C9.454 4.5.5 11.887.5 21s8.954 16.5 20 16.5s20-7.387 20-16.5s-8.954-16.5-20-16.5m-.031 30.791c-9.551 0-17.292-6.387-17.292-14.266c0-7.879 7.741-14.266 17.292-14.267c9.551 0 17.293 6.388 17.293 14.267c-.002 7.879-7.742 14.266-17.293 14.266"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloud(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M33.5 34.5c3.689 0 7-3.529 7-7.409c0-3.649-2.65-7-6.02-7.34c.06-.47.09-.98.09-1.49c0-5.99-4.631-10.84-10.36-10.84c-4.56 0-8.42 3.07-9.819 7.34c-.711-.21-1.461-.27-2.221-.32c-5.62-.38-8.34 1.54-8.34 8.792c0 .469.029.92.109 1.35C1.97 25.072.5 26.922.5 29.121c0 2.561 2.54 5.379 5 5.379z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Collapse(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.325 36.209L5.78 40.66c1.134 1.135 1.576 1.105 2.679 0l5.691-5.684l5.35 5.354V22.5H1.655l5.359 5.346l-5.689 5.688c-1.084 1.081-1.116 1.564 0 2.675M39.676 6.79l-4.455-4.451c-1.134-1.134-1.576-1.104-2.679 0l-5.69 5.685L21.5 2.669v17.83h17.846l-5.359-5.347l5.689-5.685c1.084-1.082 1.113-1.565 0-2.677"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Comment(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M29.5 30.5h7c2.529 0 3-.529 3-3v-21c0-2.41-.59-3-3-3h-32c-2.47 0-3 .53-3 3v20.971c0 2.469.41 3.029 3 3.029h13s9.562 8.756 10.75 9.812c.422.375 1.281.172 1.25-.812zm-22-9h21v3h-21zm0-6h13v3h-13zm0-6h26v3h-26z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Computer(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m18.5 35.5l-8 2c-.48.04-1 .52-1 1v1c0 .5.47 1 1 1h20c.43 0 1-.41 1-1v-1c-.02-.52-.55-.98-1-1l-8-2zm19-1c2.59 0 3-.529 3-3v-26c0-2.391-.55-3-3-3h-34c-2.43 0-3 .54-3 3v26c0 2.51.529 3 3 3zm-2-27v22h-30v-22z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Delete(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 16.5v17h3v-17zm6 0v17h3v-17zm6 0v17h3v-17zm3-12c0-2.5-.609-3-3-3h-10c-2.52 0-2.98.55-2.98 3.01L11.5 7.5h-8c-1.48 0-2 .49-2 2v1c0 1.55.52 2 2 2h1v26c0 2.49.55 3 3 3h24c2.5 0 4-.471 4-3v-26h1c1.51 0 2-.48 2-2v-1c0-1.48-.43-2-2-2h-9zm-3 0v3h-10v-3zm-15 8h21v24h-21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Download(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M40.5 38.5v-1c0-1.48-.43-2-2-2h-34c-1.48 0-2 .49-2 2v1c0 1.55.52 2 2 2h34c1.51 0 2-.48 2-2m-15-36h-8c-2.5 0-3 .47-3 3v11h-10l17 17l17-17h-10v-11c0-2.53-.529-3-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dribble(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 1C9.954 1 1 9.954 1 21s8.954 20 20 20s20-8.954 20-20S32.046 1 21 1m0 2.898c4.357 0 8.334 1.63 11.354 4.312c-2.219 2.927-5.59 4.876-8.968 6.195a89.077 89.077 0 0 0-6.415-10.03a17.132 17.132 0 0 1 4.03-.477zm-7.276 1.62c2.23 3.336 4.39 6.429 6.363 9.93c-4.99 1.293-10.629 2.069-15.838 2.082c1.098-5.328 4.677-9.752 9.475-12.011zm20.527 4.67a17.034 17.034 0 0 1 3.851 10.699c-3.956-.78-7.89-.984-11.896-.58c-.45-1.123-.996-2.19-1.519-3.34c3.453-1.393 7.145-3.777 9.564-6.78zm-12.775 7.906c.428.91.924 1.876 1.39 2.863c-5.57 2.456-11.495 5.738-14.57 11.492a17.043 17.043 0 0 1-4.39-11.95c5.965-.028 11.82-.774 17.57-2.405m9.038 4.663a29.15 29.15 0 0 1 7.375.95a17.1 17.1 0 0 1-7.347 11.487c-.918-4.175-1.793-8.17-3.34-12.198a22.966 22.966 0 0 1 3.311-.24zm7.464.31c-.012.098-.023.194-.037.29c.014-.097.026-.193.037-.29m-13.94.71c1.576 4.073 2.813 8.583 3.643 12.972a17.045 17.045 0 0 1-6.68 1.354a17.027 17.027 0 0 1-10.495-3.6c3.097-5.024 7.894-8.826 13.531-10.725z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Email(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M40.5 31.5v-18S22.3 26.2 20.53 26.859C18.79 26.23.5 13.5.5 13.5v18c0 2.5.53 3 3 3h34c2.529 0 3-.439 3-3m-.029-21.529c0-1.821-.531-2.471-2.971-2.471h-34c-2.51 0-3 .78-3 2.6l.03.28s18.069 12.44 20 13.12c2.04-.79 19.97-13.4 19.97-13.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ErrorIcon(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m29.582 8.683l-.129.12L8.3 29.954a3.308 3.308 0 0 0-.547.688c-2.04-2.639-3.233-6-3.233-9.701c0-8.797 6.626-15.482 15.421-15.482c3.691 0 7.014 1.185 9.641 3.224M10.937 33.704c.189-.117.388-.287.606-.507l21.151-21.151l.041-.04c1.74 2.518 2.746 5.602 2.746 8.994c0 8.785-6.696 15.541-15.481 15.541c-3.432 0-6.546-1.035-9.063-2.837M.5 21C.5 31.775 9.235 40.5 20 40.5c10.767 0 19.501-8.725 19.501-19.5s-8.734-19.5-19.5-19.5S.5 10.225.5 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Expand(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.391 16.17l4.439 4.439c1.13 1.131 1.57 1.101 2.67 0l6.67-6.67l5.33 5.34V1.5H22.721l5.34 5.33l-6.67 6.67c-1.08 1.08-1.112 1.561 0 2.67m-1.782 10.66l-4.439-4.439c-1.13-1.131-1.57-1.101-2.67 0l-6.67 6.67l-5.33-5.34V41.5h17.779l-5.34-5.33l6.67-6.67c1.08-1.08 1.112-1.561 0-2.67"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Facebook(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M28.08 6.51c.29-.01 1.109-.03 1.42-.01c1.12 0 2.84-.1 4 0v5h-4c-1.46 0-2 .35-2 2v4h5v5h-5v14h-5v-14h-4v-5h4l-.061-4.42c0-3.061.621-4.92 3.15-6.02c.651-.28 1.641-.53 2.491-.55M.5 1.5v39h39v-39z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Favorite(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M21 1c1.081 0 5.141 12.315 6.201 13.126s13.461 1.053 13.791 2.137c.34 1.087-9.561 8.938-9.961 10.252c-.409 1.307 3.202 13.769 2.331 14.442c-.879.673-11.05-6.79-12.361-6.79S9.52 41.63 8.641 40.957c-.871-.674 2.739-13.136 2.329-14.442c-.399-1.313-10.3-9.165-9.96-10.252c.33-1.084 12.731-1.326 13.791-2.137S19.91 1 21 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Feedback(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.5 25.5v4c.016 2.812 1.344 2.375 2.328 1.531L14.5 25.91v2.59c0 2.43.56 3 3 3h9s5.209 6.125 5.25 6.084c.75.916 2.781.604 2.75-1.084v-5h3c2.45 0 3-.609 3-3v-15c0-2.4-.59-3-3-3h-10v-2c0-2.47-.46-3-3-3h-21c-2.36 0-3 .51-3 3v13c0 2.439.55 4 3 4zm25 3v4.721l-4-4.721h-9c-.75 0-1-.27-1-1v-13c0-.67.31-1 1-1h18c.689 0 1 .37 1 .94V27.5c0 .721-.359 1-1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flickr(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M27.96 13.65c4.06 0 7.36 3.289 7.36 7.35s-3.301 7.35-7.36 7.35s-7.35-3.289-7.35-7.35s3.29-7.35 7.35-7.35m-15.92 0c4.06 0 7.35 3.289 7.35 7.35s-3.29 7.35-7.35 7.35S4.68 25.061 4.68 21s3.3-7.35 7.36-7.35M.5 1.5v39h39v-39z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folder(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.5 7.5v26c0 2.52.51 3 3 3h34c2.471 0 3-.46 3-3v-21c0-2.46-.471-3-3-3h-17v-2c0-2.5-.52-3-3-3h-14c-2.48 0-3 .43-3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Github(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 1C9.954 1 1 9.954 1 21s8.954 20 20 20s20-8.954 20-20S32.046 1 21 1m0 2.178c9.843 0 17.822 7.979 17.822 17.822S30.843 38.822 21 38.822S3.178 30.843 3.178 21S11.157 3.178 21 3.178m-8.846 5.219c-.507 1.476-.684 3.077-.229 4.476a6.665 6.665 0 0 0-1.279 2.1c-.812 2.566-.614 5.848 1.164 7.832c.625.686 1.473 1.248 2.539 1.689s2.438.723 4.113.844c-1.125.526-1.918.848-2.194 2.01c-1.255.838-2.78.639-3.887-.272c-.874-.64-1.271-1.771-2.239-2.192c-.061-.062-.251-.105-.57-.137c-.32-.031-.572.06-.755.273c-.091.09-.085.188.022.295c.732.597 1.439 1.229 1.853 1.986c.396.822.942 1.318 1.37 1.699c1.184.8 2.735 1.086 4.125.578c-.162.979.173 2.74-.148 3.613a1.928 1.928 0 0 1-.525.641c-.203.199-.835.447-.707.752c.061.137.273.221.639.252c.806-.021 1.719-.348 2.191-.982c.168-.229.363-.526.363-.891v-3.742c0-.427.156-.729.339-.914c.183-.182.349-.322.661-.365v4.933c0 .426-.523 1.614-.586 1.735c-.139.236-.45.451-.456.729c0 .093.033.146.14.159a2.077 2.077 0 0 0 1.784-1.165c.415-.63.509-1.005.618-1.733V28h1v4.6c0 .608-.1 1.187.175 1.732c.274.549 1.114.9 1.753 1.053c.336.092.504.127.625.112c.122-.015.146-.067.131-.159c-.072-.275-.463-.513-.625-.729c-.213-.275-.559-.853-.559-1.736v-4.932c.328.074.463.184.662.365c.197.184.338.488.338.915v3.742c0 .365.264.662.432.891c.51.648 1.367.978 2.17.982c.365-.031.623-.115.685-.252s.038-.25-.114-.342s-.331-.229-.561-.41a1.895 1.895 0 0 1-.52-.641c-.137-1.569-.001-3.252-.111-4.84c-.219-1.732-.885-2.304-2.215-2.968c1.586-.12 2.889-.411 3.908-.868c3.053-1.482 3.896-3.775 3.91-6.849c-.049-1.979-.619-3.703-1.922-5.025a6.157 6.157 0 0 0 .161-2.189a7.524 7.524 0 0 0-.481-2.056c-1.127.062-2.072.288-2.834.684s-1.311.731-1.646 1.006c-2.699-.613-5.59-.616-8.18.091a8.024 8.024 0 0 0-4.528-1.78"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GithubText(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m34.492 11.5l3-.012l.008 4.039l-3-.025v8.225c0 1.956.494 2.935 1.744 2.935c.881 0 1.256-.248 2.256-.743v4.378c-1 .584-2.336.876-3.812.876c-2.07 0-3.502-.761-4.299-2.277c-.596-1.14-.889-2.935-.889-5.387l-.005-8.035s-1.968.032-1.968.048l-.027-4.02h2c-.005 0-.033-3.955-.037-4.011l5.037-.015V11.5zM20.5 30.5v-19h5v19zM9.872 11.205c-1.932 0-3.593.656-4.985 1.97c-1.448 1.4-2.173 3.152-2.173 5.254c0 1.4.481 2.712 1.276 3.938c.709 1.138 1.51 1.869 2.51 2.188v.089c-1 .407-1.406 1.43-1.406 3.063c0 1.255.406 2.188 1.406 2.803v.088c-3 .904-4 2.582-4 5.033c0 2.13.86 3.69 2.62 4.688c1.393.787 3.156 1.182 5.314 1.182c5.256 0 7.881-2.264 7.881-6.787c0-2.831-2.035-4.568-6.098-5.21c-.938-.146-1.647-.495-2.13-1.05c-.37-.38-.556-.758-.556-1.14c0-1.08.567-1.706 1.703-1.882c1.732-.263 3.146-1.09 4.239-2.477c1.094-1.387 1.641-3.012 1.641-4.88c0-.584-.102-1.211-.33-1.882c.738-.176 1.715-.336 1.715-.481v-4.51c-2 .7-3.377 1.05-4.74 1.05a7.58 7.58 0 0 0-3.887-1.047m.171 4.026c.852 0 1.521.337 2.003 1.008c.397.613.597 1.344.597 2.19c0 2.043-.867 3.063-2.6 3.063c-1.79 0-2.685-1.008-2.685-3.021c0-2.16.895-3.24 2.685-3.24m.297 17.82c2.386 0 3.581.744 3.581 2.232c0 1.576-1.095 2.363-3.281 2.363c-2.499 0-3.75-.762-3.75-2.277c0-1.547 1.15-2.318 3.45-2.318M25.947 4.76c0 1.802-1.325 3.263-2.961 3.263s-2.961-1.461-2.961-3.262S21.35 1.5 22.986 1.5s2.961 1.46 2.961 3.262z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Googleplus(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.17 32.23c-1.37 0-2.71.25-4.07.438C4.01 33.08 2.12 33.92.5 35.6v4.9h18.41c.06 0 .08-.4.08-.62c-.021-1.43-.971-2.729-2.12-3.72c-1.63-1.381-3.511-2.5-5.32-3.84c-.47-.07-.93-.09-1.38-.09M.5 22.5v9.9c2.439-1.45 4.689-1.621 7.15-1.7c.76 0 1.42-.021 1.979-.05a26.66 26.66 0 0 1-1.14-1.66c-.85-1.35-.65-2.99-.04-4.45c-2.869.37-5.51-.16-7.949-2.04M7.65 6.87c-.7.01-1.42.17-2.09.479C4.2 8.01 3.2 9.09 2.83 10.72c-.28 1.681-.21 3.47.21 5.28c.51 2.061 1.521 4.02 3.351 5.45c1.369.97 3.459 1.11 5.219.64c1.371-.41 2.351-1.71 2.891-3.12c.48-1.53.311-3.27-.109-5.03c-.541-2.25-1.631-4.409-3.551-6.049A5.032 5.032 0 0 0 7.65 6.87M.5 1.5v6.87c2.439-2.35 6.38-3.87 10-3.87h12L19.45 8h-3.46c.84.811 1.609 1.32 2.229 2.35c.79 1.351 1.21 2.801 1.23 4.471c-.011 1.609-.4 3.34-1.4 4.84c-.819 1.2-1.989 2.21-3.22 3.17c-.811.812-1.52 1.479-1.58 2.771c.01.961.641 1.819 1.5 2.439c1.43 1.12 2.66 2.04 4.01 3.2c1.29 1.12 2.49 2.489 3.261 4.17c.66 1.528.67 3.459.211 5.089H39.5v-25h-7v8h-3v-8h-8v-3h8v-8h3v8h7v-11z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GroupIcon(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M29.609 39h.781c4.959-.07 9.689-.6 10.539-1.391c.211-1.719.45-2.549-7.57-6.17c-1.141-.609-.529-2.969-.529-2.969c2.411-1.74 4.062-5.761 4.062-8.991c0-5.57-2.633-8.21-6.502-8.48h-.781c-3.868.27-6.5 2.9-6.5 8.48c0 3.23 1.65 7.25 4.062 8.991c0 0 .61 2.359-.528 2.969c-8.021 3.621-7.782 4.451-7.571 6.17c.848.791 5.578 1.321 10.537 1.391m-18.001-8h.782c4.959-.07 9.689-.6 10.54-1.391c.211-1.719.449-2.549-7.571-6.17c-1.14-.609-.529-2.969-.529-2.969c2.41-1.74 4.061-5.76 4.061-8.99c0-5.57-2.631-8.21-6.5-8.48h-.782c-3.869.27-6.5 2.9-6.5 8.48c0 3.23 1.65 7.25 4.061 8.99c0 0 .61 2.36-.529 2.969C.62 27.06.86 27.89 1.071 29.609C1.919 30.4 6.649 30.93 11.608 31"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Home(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.68 15.726c-.1.039-.16.09-.18.149zm-.18.149v20.627c0 2.509.49 2.998 3 2.998h7c2.51 0 3-.461 3-3v-8h9v8.031c0 2.51.51 2.979 3 2.969c.04.031 7 0 7 0c2.529 0 3-.526 3-2.997V16.495c0-.08-.09-.15-.27-.23L20 1.5L2.68 15.726z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Image(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.5 7.5v27c0 2.52.51 3 3 3h34c2.471 0 3-.46 3-3v-27c0-2.46-.471-3-3-3h-34c-2.48 0-3 .43-3 3m35.29 23H5.23c3.34-4.87 9.279-12.99 10.789-12.99c1.461 0 6.42 6.561 8.661 8.87c0 0 2.881-3.851 4.391-3.851c1.538 0 6.669 7.931 6.719 7.971m-8.979-17c0-2.04 1.649-3.689 3.689-3.689s3.689 1.649 3.689 3.689s-1.649 3.689-3.689 3.689s-3.689-1.649-3.689-3.689"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageOutline(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M35.79 31.5c-.05-.04-5.181-7.971-6.72-7.971c-1.51 0-4.391 3.851-4.391 3.851c-2.24-2.31-7.2-8.87-8.661-8.87c-1.509 0-7.449 8.12-10.789 12.99zm-8.979-17c0 2.04 1.649 3.69 3.689 3.69s3.689-1.65 3.689-3.69s-1.649-3.69-3.689-3.69s-3.689 1.65-3.689 3.69M.5 7.5v27c0 2.52.51 3 3 3h34c2.471 0 3-.46 3-3v-27c0-2.46-.471-3-3-3h-34c-2.48 0-3 .43-3 3m3 0h34v27h-34z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Instagram(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.53 21.5c0 6.609 5.36 11.971 11.97 11.971c6.609 0 11.971-5.361 11.971-11.971c0-1.04.029-2 .029-3h3v15c0 1.98-1.02 3-3 3h-25c-1.98 0-3-1.02-3-3v-15h4c0 1 .03 1.96.03 3M.5 6.68v29.64c0 3.369 1.81 5.18 5.18 5.18h29.64c3.369 0 5.18-1.811 5.18-5.18V6.68c0-3.369-1.811-5.18-5.18-5.18H5.68C2.31 1.5.5 3.311.5 6.68M12.56 21.5c0-4.39 3.55-7.939 7.94-7.939a7.932 7.932 0 0 1 7.939 7.939a7.931 7.931 0 0 1-7.939 7.939a7.932 7.932 0 0 1-7.94-7.939M29.5 8c0-.83.67-1.5 1.5-1.5h3c.83 0 1.5.67 1.5 1.5v3c0 .83-.67 1.5-1.5 1.5h-3c-.83 0-1.5-.67-1.5-1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Like(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.938 10.725C14.51.796 1.5 6.205 1.5 17.021c0 8.122 17.836 20.827 19.438 22.479C22.551 37.848 39.5 25.143 39.5 17.021c0-10.734-12.122-16.225-18.562-6.296"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Linkedin(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.5 17.5h6v17h-6zm20.141-.58c1.689.01 3.85.84 5.17 2.18c1.369 1.53 1.689 3.259 1.689 5.4v10h-6v-9.29c0-1.7-.689-3.771-2.939-3.84c-1.32.021-2.17.78-2.801 2.06c-.18.42-.14.891-.14 1.36l-.12 9.71h-6v-17h6l.12 2.22a6.19 6.19 0 0 1 1.69-1.829c.96-.691 2.081-.952 3.331-.971M9.5 9.35c1.54.021 3.07 1.23 3.13 3.07c.04 1.641-1.39 3.07-3.17 3.07h-.04c-1.53 0-3.03-1.25-3.1-3.07c.03-1.62 1.39-3.029 3.18-3.07M.5 1.5v39h39v-39z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Listview(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M39.5 10.5v-1c0-1.47-.48-2-2-2h-34c-1.53 0-2 .52-2 2v1c0 1.55.52 2 2 2h34c1.5 0 2-.48 2-2m0 11v-1c0-1.47-.48-2-2-2h-34c-1.53 0-2 .52-2 2v1c0 1.55.52 2 2 2h34c1.5 0 2-.48 2-2m0 11v-1c0-1.471-.48-2-2-2h-34c-1.53 0-2 .52-2 2v1c0 1.55.52 2 2 2h34c1.5 0 2-.48 2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Location(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M33 13.924C33 6.893 27.594 1 20.51 1S8 6.897 8 13.93C8 16.25 8.324 18 9.423 20h-.021l10.695 20.621c.402.551.824-.032.824-.032C20.56 41.13 31.616 20 31.616 20h-.009C32.695 18 33 16.246 33 13.924m-18.249-.396c0-3.317 2.579-6.004 5.759-6.004c3.179 0 5.76 2.687 5.76 6.004s-2.581 6.005-5.76 6.005c-3.18 0-5.759-2.687-5.759-6.005"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lock(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" d="M23.68 28.484c0-1.76-1.42-3.18-3.18-3.18s-3.18 1.42-3.18 3.18c0 1.412.91 2.6 2.17 3.02l-1.99 4.98h6l-1.99-4.98a3.178 3.178 0 0 0 2.17-3.02"/><path fill="currentColor" d="M20.5 6.411c5.98 0 8.67 5.073 9 10.073h5c-.09-9-4.67-14.903-14-14.983c-9.27-.09-14 6.983-14 14.983h4.969c.26-5 2.831-10.073 9.031-10.073m14 12.089h-28c-2.41 0-3 .655-3 2.984v17c0 2.49.561 3 3 3h28c2.609 0 3-.471 3-3v-17c0-2.458-.46-2.984-3-2.984m-11 17.984h-6l1.99-4.98a3.176 3.176 0 0 1-2.17-3.02c0-1.76 1.42-3.18 3.18-3.18s3.18 1.42 3.18 3.18c0 1.41-.91 2.6-2.17 3.02z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Minus(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M39.5 21.5v-3c0-1.48-.43-2-2-2h-34c-1.48 0-2 .49-2 2v3c0 1.55.52 2 2 2h34c1.51 0 2-.48 2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Next(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.933 1L34 21.068L14.431 40.637l-4.933-4.933l14.638-14.636L9 5.933z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NextLight(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11 38.32L28.609 21L11 3.68L13.72 1L34 21.01L13.72 41z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Page(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M22.5 1.5h-14c-2.55 0-3 .561-3 3v32c0 2.49.55 3 3 3h24c2.5 0 3-.47 3-3v-22h-13zm13 10l-10-10v10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Path(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M39.497 16.86c-.06 4.266-2.154 8.135-5.24 10.613c-4.01 2.982-7.757 3.977-13.758 4.055v2.399c0 3.082-1.776 5.756-4.397 6.934a6.81 6.81 0 0 1-2.893.64c-1.373-.015-2.71-.464-3.71-1.001v-6.934c1 .645 2.271 1.119 3.345.816c.948-.342 1.656-1.07 1.656-1.948V13.502l6 .002v11.68c2-.022 3.354-.19 4.805-.506c1.9-.453 3.8-1.16 5.17-2.35c1.681-1.565 2.463-3.427 2.476-5.465c-.025-3.019-1.557-5.53-3.933-6.986C26.324 8.37 23.35 7.86 20.486 7.85c-3.087.066-6.12.66-8.593 2.16c-2.524 1.596-3.88 4.236-3.901 6.852c.16 1.887.655 4.511 1.967 5.708l-4.42 4.64c-1.436-1.403-2.38-3.12-3.002-4.854c-.61-1.835-1.03-3.662-1.038-5.494C1.575 12.409 3.6 8.39 6.794 5.769c4.031-3.078 9.093-4.242 13.704-4.266c5.05.084 10.041 1.327 13.757 4.24c3.488 2.891 5.21 7.082 5.242 11.116z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pencil(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.57 29.24L23.33 33s15.8-15.769 15.79-15.779c1.851-1.859 1.83-2.68 0-4.51c.01-.02-1.511-1.51-1.511-1.51zm-1.71-1.711L35.9 9.491l-3.551-3.55L14.31 23.98zM30.85 4.441l-1.5-1.5c-1.91-1.91-2.59-1.931-4.51 0L9.05 18.721l3.76 3.76zm-23.3 15.78L.5 41.5l21.33-7zm-1.5 10.638l5.26 5.262l-7.61 2.26z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Phone(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M31.5 6.65c0-2.589-.561-3.15-3-3.15h-17c-2.47 0-3 .529-3 3.15V36.5c0 2.439.56 3 3 3h17c2.5 0 3-.561 3-3zm-18 1.85h13v24h-13zm8 28h-3v-2h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Picasa(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M30.5 4.45c6 3.727 9.967 10.442 10 17.238c0 2.275-.094 3.812-.863 6.812H30.5zm8.24 26.05c-3.188 6-9.398 10.268-16.2 10.967c-3.489.22-7.04-.662-10.04-1.893V30.5zM.5 21.413c.088-7.625 4.462-14.354 10.951-17.69l7.655 6.837l-17.88 16.023C.74 24.817.5 23.093.5 21.413m10 17.007c-4-2.385-6.71-5.84-8.3-9.656l8.3-7.48zm18-21.836L13.367 2.78c2.346-.787 4.651-1.273 6.947-1.281c2.906 0 6.186.598 8.186 1.793v13.29z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pinterest(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 1c11.882.223 19.922 9.334 20 20c-.213 11.364-9.334 19.922-20 20a20.901 20.901 0 0 1-5.658-.82a28.23 28.23 0 0 0 1.123-2.01c.79-1.604 1.15-3.323 1.555-4.967c.173-.705.418-1.635.734-2.786c.346.662.972 1.238 1.88 1.728c2.366 1.152 5.3.77 7.407-.303c2.956-1.678 4.688-4.286 5.615-7.213c1.787-6.79-.517-13.314-6.674-15.94c-7.424-2.235-15.709.556-18.596 7.192c-1.26 4.436-1.186 10.13 3.196 11.987c.23.087.44.087.626 0c.412-.235.89-2.272.8-2.7c-.028-.13-.13-.31-.303-.54c-2.422-3.278-.946-8.158 1.317-10.604c3.85-3.306 9.84-3.842 13.37-.691c3.131 3.653 2.42 9.338.41 12.959c-1.11 1.77-2.6 3.109-4.6 3.133c-2.106-.047-3.694-1.83-3.154-3.846c.46-2.365 1.733-4.875 1.771-7.084c-.102-1.837-1.02-3.002-2.72-3.023c-2.714.342-3.81 2.879-3.846 5.184c.087 1.107.153 2.235.647 3.153a266.109 266.109 0 0 0-1.339 5.638c-.485 2.529-1.438 5.174-1.577 7.624a23.69 23.69 0 0 0-.02 2.246C5.553 35.802 1.041 28.937.997 21.002C1.216 9.457 10.334 1.079 21 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plugin(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 7.25C10 5.75 9.532 5 7.971 5C6.471 5 6 5.812 6 7.25V8H5c-1.56 0-2 .75-2 2.25v6c0 1.529.44 1.792 2 1.792h14c1.563 0 2.004-.263 2.004-1.792v-6C21.004 8.75 20.563 8 19 8h-1v-.75C18 5.812 17.471 5 15.971 5C14.409 5 14 5.75 14 7.25V8h-4zm0 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plus(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M39.5 22.5v-3c0-1.48-.43-2-2-2h-13v-13c0-1.48-.49-2-2-2h-3c-1.55 0-2 .52-2 2v13h-14c-1.48 0-2 .49-2 2v3c0 1.55.52 2 2 2h14v14c0 1.51.48 2 2 2h3c1.48 0 2-.43 2-2v-14h13c1.51 0 2-.48 2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Preview(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.5 6.5c0-2.38-.5-3-3-3h-16c-2.5 0-3 .69-3 3v29c0 2.5.62 3 3 3h18.939L14 32.5H5.5v-24h12v20l5-4.561zm-9 30h-3v-2h3zm14.34-4.96a1.245 1.245 0 0 0 2.49 0c0-.69-.561-1.25-1.25-1.25c-.68 0-1.24.56-1.24 1.25m2.61 7.021c4.05-.891 10.28-6.199 10.28-6.199S36.1 25.17 29.9 24.4c-.351-.041-1.57-.051-1.801-.03c-5.99.58-10.83 8.22-10.83 8.22s5.259 5.131 10.201 5.96c.58.12 2.38.12 2.98.011m-7.4-7.032c0-3.1 2.671-5.609 5.95-5.609s5.95 2.51 5.95 5.609c0 3.111-2.671 5.621-5.95 5.621s-5.95-2.509-5.95-5.621"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Print(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 30.504V33.5h15v-2.996zm0-5V28.5h15v-2.996zm-9 2.996h3v8.002c-.01 2.43.548 2.998 3 2.998h21c2.452 0 3-.452 3-2.998V28.5h3c2.55.04 3.029-.527 3-2.998v-11c0-2.35-.38-3.002-3-3.002h-3V6.502c.029-2.47-.45-3.002-3-3.002h-21c-2.4 0-3.01.572-3 3.002V11.5h-3c-2.58 0-3 .562-3 3.002v11c-.01 2.43.6 3.038 3 2.998m27 8h-21v-13h21zm1.88-18.498c0-.9.72-1.619 1.62-1.619s1.62.719 1.62 1.619s-.72 1.62-1.62 1.62s-1.62-.72-1.62-1.62M9.5 6.5h21v5.004h-21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Question(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.737 30.521c0 .815.259 1.481.775 2.01c.528.517 1.185.785 1.971.785c.797 0 1.462-.27 1.99-.785c.536-.527.807-1.193.807-2.01c0-.808-.271-1.482-.807-2c-.527-.527-1.193-.787-1.99-.787c-.786 0-1.442.26-1.971.787c-.516.518-.775 1.194-.775 2m.945-5.542h3.751c-.101-.766-.05-1.145.18-1.761a7.932 7.932 0 0 1 .936-1.75c.378-.549.814-1.075 1.305-1.583a18.86 18.86 0 0 0 1.353-1.571a9.394 9.394 0 0 0 1.056-1.731c.287-.618.428-1.294.428-2.06c0-.985-.17-1.84-.498-2.577a4.986 4.986 0 0 0-1.433-1.86c-.618-.508-1.354-.885-2.188-1.134a9.709 9.709 0 0 0-2.756-.379c-1.363 0-2.586.299-3.671.886a9.555 9.555 0 0 0-2.826 2.239l2.378 2.108a6.927 6.927 0 0 1 1.74-1.362a4.395 4.395 0 0 1 2.08-.498c.946 0 1.683.259 2.219.785c.527.528.797 1.215.797 2.08c0 .548-.141 1.055-.408 1.532c-.279.479-.618.965-1.025 1.453c-.408.486-.846.995-1.303 1.502a11.12 11.12 0 0 0-1.225 1.661a7.57 7.57 0 0 0-.815 1.971c-.194.716-.203 1.195-.075 2.049M1.5 21c0 10.766 8.735 19.5 19.5 19.5c10.766 0 19.5-8.734 19.5-19.5S31.766 1.5 21 1.5C10.235 1.5 1.5 10.234 1.5 21m3.592 0c0-8.785 7.123-15.909 15.908-15.909S36.909 12.215 36.909 21S29.785 36.909 21 36.909S5.092 29.785 5.092 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rectangle(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.5 4.5h40v33H.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RectangleOutline(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.5 4.5v33h40v-33zm3 3h34v27h-34z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Refresh(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M26.731 17.5H39.5V4.91l-4.335 4.385c-.34-.399-.651-.79-1.021-1.17c-7.48-7.5-19.572-7.5-27.042 0c-7.479 7.5-7.465 19.661.015 27.161c7.471 7.5 19.597 7.5 27.078 0c0 0 2.505-2.939 3.334-4.51l-4.129-1.84c-.609 1.068-2.39 3.148-2.39 3.148c-5.719 5.74-14.979 5.74-20.689 0c-5.72-5.729-5.719-15.031 0-20.761c5.71-5.74 14.971-5.74 20.689 0c.381.38.729.574 1.061.984z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Retweet(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m24.5 30.5l7.96 7.371L40.5 30.5h-5V9.549c0-2.5-.561-3.049-3-3.049h-18l6.641 6H29.5v18zm-8-18L8.52 5.16L.5 12.5h5v21.049c0 2.5.62 2.951 3 2.951h18.32l-6.32-6h-9v-18z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Roundedrectangle(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.5 4.5c-6.939 0-11 4.47-11 11v11c0 6.971 3.859 11 11 11h18c7.4 0 11-4.029 11-11v-11c0-6.97-4.061-11-11-11z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RoundedrectangleOutline(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.5 4.5c-6.939 0-11 4.47-11 11v11c0 6.971 3.859 11 11 11h18c7.4 0 11-4.029 11-11v-11c0-6.97-4.061-11-11-11zm-1.01 3H30.5c4.59 0 7 2.52 7 7v13c0 4.62-2.45 7-7 7h-20c-4.54 0-7-2.49-7-7V14.37c0-4.38 2.529-6.87 6.99-6.87"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rss(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M40.5 41.5c0-22.092-17.908-40-40-40v8c17.673 0 32 14.326 32 32zm-16 0c0-13.256-10.745-24-24-24v7.998c8.836 0 16.001 7.166 16.001 16.002zm-16 0a8 8 0 0 0-8-8v8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Save(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24.531 1.5H17.5c-2.53 0-3 .566-3 3.016V15.5H5l16.531 16.969L38 15.5H27.5v-11c0-2.453-.453-3-2.969-3M11 36.5l-2.63-6H3l3.94 9.62c.35 1.05 1.091 1.38 3.123 1.38H33c1.969 0 2.38-.3 2.729-1.35l3.94-9.621h-5.13L32 36.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Search(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 17.838c0 8.747 7.131 15.827 15.94 15.827c8.796 0 15.938-7.08 15.938-15.827S25.736 2 16.94 2C8.131 2 1 9.091 1 17.838m5.051 0c0-5.979 4.868-10.817 10.89-10.817c6.01 0 10.888 4.839 10.888 10.817c0 5.979-4.878 10.818-10.888 10.818c-6.022 0-10.89-4.84-10.89-10.818m22.111 14.523l6.855 7.809c1.104 1.102 1.816 1.111 2.938 0l2.201-2.181c1.082-1.081 1.149-1.778 0-2.921l-7.896-6.775z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Settings(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.62 24.5c.4 1.62 1.06 3.13 1.93 4.49l-2.43 2.44c-1.09 1.09-1.08 1.74-.12 2.7l2.37 2.37c.97.971 1.63.95 2.7-.12l2.55-2.56c1.2.688 2.5 1.22 3.88 1.56v3.12c0 1.55.47 2 1.82 2h3.36c1.37 0 1.82-.48 1.82-2v-3.12c1.38-.34 2.68-.87 3.88-1.56l2.61 2.619c1.08 1.068 1.729 1.09 2.699.131l2.381-2.381c.949-.949.97-1.602-.131-2.699l-2.5-2.5a14.665 14.665 0 0 0 1.938-4.49h3.302c1.368 0 1.818-.48 1.818-2v-3c0-1.48-.393-2-1.818-2h-3.302c-.34-1.38-.87-2.68-1.562-3.88l2.382-2.37c1.05-1.05 1.14-1.7.13-2.7l-2.38-2.38c-.95-.95-1.632-.94-2.7.13l-2.26 2.25A14.946 14.946 0 0 0 24.5 6.62V3.5c0-1.48-.391-2-1.82-2h-3.36c-1.35 0-1.82.49-1.82 2v3.12c-1.62.4-3.13 1.06-4.49 1.93L10.75 6.3C9.68 5.23 9 5.22 8.05 6.17L5.67 8.55c-1.01 1-.92 1.65.13 2.7l2.37 2.37c-.68 1.2-1.21 2.5-1.55 3.88h-3.3c-1.35 0-1.82.49-1.82 2v3c0 1.55.47 2 1.82 2zm8.66-3.5c0-3.16 2.56-5.72 5.72-5.72s5.721 2.56 5.721 5.72a5.72 5.72 0 1 1-11.441 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Share(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M30.5 18.5v6l10-9.929L30.5 4.5v5c-15.3.1-15 15-15 15s5.45-7.49 15-6m-8-13h-19c-2.46 0-3 .7-3 3v24c0 2.49.6 3 3 3h29c2.41 0 3-.451 3-3v-8l-5 4.289V30.5h-25v-20h12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tablet(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M35.5 4.5c0-2.38-.5-3-3-3h-25c-2.5 0-3 .561-3 3v34c0 2.41.561 3 3 3h25c2.471 0 3-.561 3-3zm-26 2h21v27h-21zm14 32h-7v-2h7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextIcon(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.5 4.5h37v9h-2.85l-2.15-7h-10v27l6 1.8v2.2h-19v-2.2l6-1.8v-27H7.55l-2.2 7H2.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tileview(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 9.53v-4c0-2.5-.53-3-3-3l-4-.03c-2.38 0-3 .561-3 3v4c0 2.471.5 3 3 3h4c2.5 0 3-.5 3-2.97m0 12.999v-4c0-2.5-.53-3-3-3l-4-.03c-2.38 0-3 .561-3 3v4c0 2.471.5 3 3 3h4c2.5.001 3-.499 3-2.97m0 13v-4c0-2.5-.53-3-3-3l-4-.029c-2.38 0-3 .561-3 3v4c0 2.471.5 3 3 3h4c2.5 0 3-.5 3-2.971m13-25.999v-4c0-2.5-.529-3-3-3l-4-.03c-2.38 0-3 .561-3 3v4c0 2.471.5 3 3 3h4c2.5 0 3-.5 3-2.97m0 12.999v-4c0-2.5-.529-3-3-3l-4-.03c-2.38 0-3 .561-3 3v4c0 2.471.5 3 3 3h4c2.5.001 3-.499 3-2.97m0 13v-4c0-2.5-.529-3-3-3l-4-.029c-2.38 0-3 .561-3 3v4c0 2.471.5 3 3 3h4c2.5 0 3-.5 3-2.971m13-25.999v-4c0-2.5-.529-3-3-3l-4-.03c-2.38 0-3 .561-3 3v4c0 2.471.5 3 3 3h4c2.5 0 3-.5 3-2.97m0 12.999v-4c0-2.5-.529-3-3-3l-4-.03c-2.38 0-3 .561-3 3v4c0 2.471.5 3 3 3h4c2.5.001 3-.499 3-2.97m0 13v-4c0-2.5-.529-3-3-3l-4-.029c-2.38 0-3 .561-3 3v4c0 2.471.5 3 3 3h4c2.5 0 3-.5 3-2.971"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tumblr(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.83 29.38c-.21-.19-.33-.739-.33-.88v-11h7v-4h-7v-6h-4c-.28 3.33-.83 6-5 6v4h3c0 4.17-.05 8.55-.05 12.24c0 1.74 1 3.24 2.409 4.039c1.32.681 1.681 1.051 3.051 1.131c1.51.07 3.17-.01 4.46-.16c1.14-.18 2.37-.65 3.2-1.13V28.8v.04c0 .37-1.021.61-1.7.86c-1.14.38-3.85.87-5.04-.32M40.5 1.5v39h-39v-39z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Twitter(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M41 7.861c-.268.559-.766 1.266-1.495 2.118c-.728.855-1.626 1.605-2.697 2.253c.03.235.052.456.067.661c.115 3.173-.653 6.464-1.717 9.205c-2.065 5.098-5.227 9.201-9.744 11.987c-4.71 2.676-10.196 3.244-15.25 2.759C6.808 36.463 3.5 35.378 1 33.334c4.574.54 8.844-1 12.219-3.488c-3.75.094-6.473-2.592-7.76-5.696c.58.151 1.187.133 1.74.09c.68-.065 1.326-.127 1.96-.266c-2.39-.764-4.481-2.167-5.573-4.282c-.64-1.34-.93-2.616-.936-4.062c1.12.584 2.467 1.154 3.701 1.103c-1.863-1.562-3.357-3.558-3.635-5.849c-.176-1.91.308-3.653 1.004-5.275c2.786 3.031 5.913 5.528 9.52 7.063c2.467 1 4.926 1.536 7.47 1.545c-.294-2.249-.066-4.423 1.003-6.268c1.258-1.985 3.093-3.135 5.15-3.709c2.945-.748 5.817.31 7.626 2.34c1.95-.208 3.819-1.066 5.306-1.942c-.632 1.895-1.862 3.778-3.657 4.68A23.27 23.27 0 0 0 41 7.861"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unlock(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.5 21.501v17c0 2.49.561 2.999 3 2.999h28c2.609 0 3-.471 3-2.999v-17c0-2.459-.46-3.001-3-3.001h-28c-2.41 0-3 .671-3 3.001m20.18 7c0 1.41-.91 2.599-2.17 3.019l1.99 4.98h-6l1.99-4.98a3.176 3.176 0 0 1-2.17-3.019c0-1.76 1.42-3.181 3.18-3.181s3.18 1.421 3.18 3.181m-17.18-12l5.41-.029c.26-4.78 2.439-9.631 8.64-9.631c4.56 0 6.851 2.74 8.24 5.99l4.771-3.359c-1.939-4.631-6.431-7.91-12.811-7.971c-9.27-.09-14.25 6.641-14.25 15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func User(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.474 40.996c.347 0 .705.01 1.051 0c6.773-.1 13.221-.861 14.389-1.973c.27-2.463.598-3.655-10.325-8.822c-1.554-.871-.731-4.246-.731-4.246c3.289-2.482 5.547-8.231 5.547-12.838c0-7.961-3.59-11.726-8.877-12.117h-1.052c-5.278.391-8.867 4.156-8.867 12.117c0 4.607 2.248 10.354 5.539 12.838c0 0 .82 3.375-.725 4.246c-10.934 5.167-10.606 6.36-10.335 8.822c1.167 1.111 7.613 1.873 14.386 1.973"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Videocamera(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 34.5h24c2.529 0 3-.471 3-3v-9l7.52 10.18c1.15 1.12 1.91 1.15 2.48-.25V10.61c-.57-1.4-1.33-1.44-2.48-.32L31.5 21.5v-9c0-2.5-.48-3-3-3h-24c-2.48 0-3 .55-3 3.45V31.5c0 2.5.48 3 3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func View(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.3 20.1c0 3.1 2.6 5.7 5.7 5.7s5.7-2.6 5.7-5.7s-2.6-5.7-5.7-5.7s-5.7 2.6-5.7 5.7m8.1 12.3C30.1 30.9 40.5 22 40.5 22s-7.7-12-18-13.3c-.6-.1-2.6-.1-3-.1c-10 1-18 13.7-18 13.7s8.7 8.6 17 9.9c.9.4 3.9.4 4.9.2M11.1 20.7c0-5.2 4.4-9.4 9.9-9.4s9.9 4.2 9.9 9.4S26.5 30 21 30s-9.9-4.2-9.9-9.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vimeo(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M40.5 1.5v39h-39v-39zm-8.04 10.396c-4.03-3.33-9.38 1.221-10.63 4.86c.77.021 1.66-.06 2.32.19c1.029.55 1.05 1.63.97 2.529c-.46 2.01-1.88 5.669-3.71 5.96c-.38.039-.78-.17-1.19-.642c-1.6-2-1.399-4.63-1.78-6.84c-.34-1.85-.609-5.58-2.799-6.391c-2.41-.299-4.531 1.5-6.08 2.931a62.74 62.74 0 0 1-2.91 2.51v.16c.24.29.43.58.64.851c.36.359 1.021.27 1.42.139c1.47-.379 1.979-.75 2.771.641c1.18 3.21 2.08 6.811 2.859 9.88c.551 1.74 1.27 3.979 3.03 4.699c1.09.461 2.96-.16 3.8-.66c3.55-2.25 6.312-5.85 8.28-9.319c1.44-2.929 5.029-8.748 3.009-11.498"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WthreeC(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M36.294 22.483c.096-.389-.003-.787-.367-.914l-1.507-.686c-.374-.155-.848.054-.914.412c-.291.623-.704 1.133-1.371 1.143c-.58 0-1.053-.313-1.418-.938s-.55-1.44-.55-2.445s.185-1.83.55-2.469c.365-.64.838-.96 1.418-.96c.73.053 1.076.515 1.37 1.096c.157.383.58.473.915.367l1.507-.686c.35-.184.507-.638.32-.96c-.882-1.888-2.255-2.833-4.113-2.833c-1.555 0-2.813.594-3.771 1.782c-.961 1.189-1.441 2.744-1.441 4.663s.479 3.467 1.44 4.64c.96 1.173 2.217 1.759 3.772 1.759c1.859 0 3.246-.989 4.16-2.971m-14.537 2.97c2.423-.035 4.269-1.701 4.298-3.885c-.043-1.152-.507-2.049-1.326-2.742c.72-.682 1.136-1.61 1.143-2.514c-.059-2.47-2.114-3.638-4.114-3.657c-1.738 0-3.078.685-4.022 2.057c-.214.305-.184.61.092.915l1.096 1.004c.352.314.743.256 1.051-.045c.457-.64 1.006-.96 1.646-.96c.518 0 .822.13.913.388c.177.385.126.736 0 1.053c-.471.488-1.376.341-2.043.341c-.434.016-.99.27-.99.686v1.647c0 .472.574.784.99.686c.503 0 1.89-.057 2.096.365c.106.244.213.442.213.595c0 .7-.353 1.05-1.116 1.05c-.73 0-1.328-.335-1.816-1.005c-.299-.361-.762-.365-1.044-.092l-1.14 1.143c-.274.276-.303.565-.09.869c1.037 1.457 2.546 2.09 4.163 2.102zm-4.525-11.938c.083-.439-.24-1.015-.64-1.015H14.58c-.365 0-.579.329-.64.695l-.96 4.963l-1.005-5c-.093-.365-.32-.658-.686-.658H9.826c-.397 0-.625.329-.686.695l-1.006 4.963l-.96-5c-.06-.365-.29-.658-.686-.658H4.523c-.502 0-.742.581-.686 1.015l2.515 11.248c.092.368.32.737.686.737h1.92c.364 0 .594-.297.686-.665l.914-4.626l.913 4.656c.062.366.29.635.686.635h1.92c.367 0 .595-.37.686-.737zM40.5 1.5v30.625l-20 9.375l-20-9.375V1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wifi(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M25 31.041a5.69 5.69 0 0 0-8 0L21 35zm8-7.921c-6.63-6.562-17.37-6.562-24 0l4 3.959c4.42-4.37 11.58-4.37 16 0zm8-7.921c-11.05-10.932-28.95-10.932-40 0l4 3.959c8.84-8.751 23.16-8.751 32 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wordpress(children ...ElementRenderer) *TopcoatIcon {
	return &TopcoatIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.976 41C9.808 40.785 1.078 31.629 1 20.977C1.216 9.817 10.323 1.079 20.976 1C32.138 1.223 40.921 10.31 41 20.977C40.78 32.147 31.644 40.922 20.976 41m0-38.583c-10.37.204-18.533 8.657-18.605 18.56c.202 10.379 8.703 18.534 18.605 18.606c10.38-.202 18.534-8.704 18.606-18.606C39.387 10.6 30.865 2.49 20.976 2.417m-4.798 34.926l5.12-14.72l5.348 14.4c-3.603 1.206-7.01 1.28-10.468.32m-3.932-26.011a31.256 31.256 0 0 1-5.486.273c3.396-4.855 8.834-7.586 14.217-7.634c4.404.083 8.51 1.747 11.521 4.479c-1.396-.104-2.232.51-2.834 1.555c-.875 2.697.974 4.695 2.011 6.72c.982 1.85.904 3.954.366 5.76l-2.56 8.73l-6.172-18.33c.643-.06 1.29-.08 1.874-.183c.691-.154.843-.797.366-1.19c-.152-.12-.32.017-.503.017l-3.703.47h-2.812c-.793 0-3.035-.916-3.223.124c-.098.4.202.616.55.679c.61.079 1.311.12 1.873.179l2.697 7.198l-3.749 11.05l-6.217-18.337c.658-.058 1.324-.08 1.92-.186c.488-.062.701-.306.641-.732a.824.824 0 0 0-.777-.642m-6.858 2.834l8.183 22.125a17.423 17.423 0 0 1-6.948-6.148c-3.04-4.851-3.36-11.006-1.235-15.977m31.977 11.36c-1.343 4.387-4.087 8.02-7.885 10.217c.183-.486.472-1.31.868-2.468l4.755-13.806c.457-1.342.776-2.834.96-4.48c.061-.673.063-1.352-.046-1.965c1.948 4.2 2.5 8.362 1.348 12.502"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
