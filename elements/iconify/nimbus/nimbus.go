package nimbus

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = "0.3.2"
	hAttr          = "1em"
	viewbox        = "0 0 0 0"
	fill           = "currentColor"
)

type NimbusIcon struct {
	*SVGSVGElement
}

type NimbusIconFn func(children ...ElementRenderer) *NimbusIcon

var IconLookup = map[string]NimbusIconFn{
	"accordion":           Accordion,
	"alignCenter":         AlignCenter,
	"alignLeft":           AlignLeft,
	"alignRight":          AlignRight,
	"apps":                Apps,
	"archive":             Archive,
	"arrowLeft":           ArrowLeft,
	"arrowRight":          ArrowRight,
	"arrowsHorizontal":    ArrowsHorizontal,
	"arrowsVertical":      ArrowsVertical,
	"backspace":           Backspace,
	"barcode":             Barcode,
	"bold":                Bold,
	"boxPacked":           BoxPacked,
	"boxUnpacked":         BoxUnpacked,
	"briefcase":           Briefcase,
	"browser":             Browser,
	"browserSearch":       BrowserSearch,
	"calendar":            Calendar,
	"calendarDays":        CalendarDays,
	"camera":              Camera,
	"cash":                Cash,
	"chatDots":            ChatDots,
	"check":               Check,
	"checkCircle":         CheckCircle,
	"chevronDown":         ChevronDown,
	"chevronLeft":         ChevronLeft,
	"chevronRight":        ChevronRight,
	"chevronUp":           ChevronUp,
	"christ":              Christ,
	"clock":               Clock,
	"close":               Close,
	"code":                Code,
	"cog":                 Cog,
	"colorPalette":        ColorPalette,
	"copy":                Copy,
	"creditCard":          CreditCard,
	"desktop":             Desktop,
	"discountCircle":      DiscountCircle,
	"diskette":            Diskette,
	"download":            Download,
	"drag":                Drag,
	"dragDots":            DragDots,
	"drink":               Drink,
	"drop":                Drop,
	"drums":               Drums,
	"duplicate":           Duplicate,
	"ecosystem":           Ecosystem,
	"edit":                Edit,
	"ellipsis":            Ellipsis,
	"exclamationCircle":   ExclamationCircle,
	"exclamationTriangle": ExclamationTriangle,
	"externalLink":        ExternalLink,
	"eye":                 Eye,
	"eyeOff":              EyeOff,
	"file":                File,
	"fileAlt":             FileAlt,
	"fingerprint":         Fingerprint,
	"fire":                Fire,
	"flag":                Flag,
	"font":                Font,
	"forbidden":           Forbidden,
	"giftBox":             GiftBox,
	"giftCard":            GiftCard,
	"glasses":             Glasses,
	"globe":               Globe,
	"guitar":              Guitar,
	"heart":               Heart,
	"history":             History,
	"home":                Home,
	"infinite":            Infinite,
	"infoCircle":          InfoCircle,
	"invoice":             Invoice,
	"italic":              Italic,
	"lifeRing":            LifeRing,
	"lightbulb":           Lightbulb,
	"link":                Link,
	"linkOff":             LinkOff,
	"list":                List,
	"location":            Location,
	"lock":                Lock,
	"lockOpen":            LockOpen,
	"logOut":              LogOut,
	"mail":                Mail,
	"marketing":           Marketing,
	"mate":                Mate,
	"menu":                Menu,
	"mobile":              Mobile,
	"money":               Money,
	"moon":                Moon,
	"notification":        Notification,
	"obelisk":             Obelisk,
	"orderedList":         OrderedList,
	"pencil":              Pencil,
	"peso":                Peso,
	"picture":             Picture,
	"pix":                 Pix,
	"planet":              Planet,
	"plusCircle":          PlusCircle,
	"printer":             Printer,
	"pyramid":             Pyramid,
	"qrCode":              QrCode,
	"questionCircle":      QuestionCircle,
	"real":                Real,
	"redo":                Redo,
	"removeFormat":        RemoveFormat,
	"rocket":              Rocket,
	"scooter":             Scooter,
	"search":              Search,
	"share":               Share,
	"shoppingCart":        ShoppingCart,
	"shot":                Shot,
	"sizeHeight":          SizeHeight,
	"sizeWidth":           SizeWidth,
	"sliders":             Sliders,
	"star":                Star,
	"stats":               Stats,
	"stickyNote":          StickyNote,
	"stop":                Stop,
	"store":               Store,
	"sun":                 Sun,
	"tag":                 Tag,
	"telephone":           Telephone,
	"textSize":            TextSize,
	"tiendanube":          Tiendanube,
	"tools":               Tools,
	"transferPeso":        TransferPeso,
	"transferReal":        TransferReal,
	"trash":               Trash,
	"truck":               Truck,
	"undo":                Undo,
	"university":          University,
	"upload":              Upload,
	"user":                User,
	"userCircle":          UserCircle,
	"userGroup":           UserGroup,
	"verticalStacks":      VerticalStacks,
	"volume":              Volume,
	"wallet":              Wallet,
	"whatsapp":            Whatsapp,
}

func Accordion(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.75 5.9V3.49L11 2.5L9.44 3L8 2.51L6.57 3L5 2.5l-2.82.75a1.24 1.24 0 0 0-.93 1.2V5.9A1.25 1.25 0 0 0 0 7.15v1.5A1.25 1.25 0 0 0 1.25 9.9v1.65a1.24 1.24 0 0 0 .93 1.2L5 13.5l1.57-.45l1.43.44l1.43-.44l1.57.45l3.75-1V9.9A1.25 1.25 0 0 0 16 8.65v-1.5a1.25 1.25 0 0 0-1.25-1.25m-11 6l-1.25-.35v-7.1l1.25-.33zm1.25.3V3.8l1 .27v7.86zm2.21-.26V4.06L8 3.82l.79.24v7.88l-.79.24zm2.84 0V4.07L11 3.8v8.4zm3.45-.38l-1.25.33V4.12l1.25.33z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignCenter(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.13 9h9.75v1.25H3.13zM.5 5.75h15V7H.5zm0 6.5h15v1.25H.5zM3.13 2.5h9.75v1.25H3.13z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignLeft(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.5 9h9.75v1.25H.5zm0-3.25h15V7H.5zm0 6.5h15v1.25H.5zm0-9.75h9.75v1.25H.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignRight(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.75 9h9.75v1.25H5.75zM.5 5.75h15V7H.5zm0 6.5h15v1.25H.5zM5.75 2.5h9.75v1.25H5.75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Apps(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.63 2.38V1.5h-1.25v.88h-.88v1.25h.88v.87h1.25v-.87h.87V2.38zm-2.1 2.96a5.79 5.79 0 0 0-2.39.6L10 6a.3.3 0 0 1-.3-.3V2.75A1.25 1.25 0 0 0 8.47 1.5H1.75A1.25 1.25 0 0 0 .5 2.75v3A1.51 1.51 0 0 0 2 7.29a1.46 1.46 0 0 0 .6-.14a4.44 4.44 0 0 1 2-.57A1.52 1.52 0 0 1 6.2 8a1.52 1.52 0 0 1-1.6 1.42a4.44 4.44 0 0 1-2-.57a1.46 1.46 0 0 0-.6-.14a1.51 1.51 0 0 0-1.5 1.51v3a1.25 1.25 0 0 0 1.25 1.28h6.72a1.25 1.25 0 0 0 1.25-1.25v-2.92A.3.3 0 0 1 10 10a.27.27 0 0 1 .12 0a5.79 5.79 0 0 0 2.39.6a2.67 2.67 0 1 0 0-5.32zM13.78 9a1.68 1.68 0 0 1-1.25.45a4.61 4.61 0 0 1-1.87-.49a1.57 1.57 0 0 0-.66-.18a1.55 1.55 0 0 0-1.55 1.55v2.92h-6.7v-3A.26.26 0 0 1 2 10h.09l.21.09a5 5 0 0 0 2.3.59A2.77 2.77 0 0 0 7.45 8A2.77 2.77 0 0 0 4.6 5.33a5 5 0 0 0-2.3.59L2.09 6A.16.16 0 0 1 2 6a.26.26 0 0 1-.25-.26v-3h6.72v2.93A1.55 1.55 0 0 0 10 7.22a1.57 1.57 0 0 0 .64-.14a4.61 4.61 0 0 1 1.87-.49a1.68 1.68 0 0 1 1.27.41a1.38 1.38 0 0 1 0 1.92z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Archive(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.5 2.75a1.25 1.25 0 0 0-1.25-1.25H1.75A1.25 1.25 0 0 0 .5 2.75v2.5h1.25v8A1.25 1.25 0 0 0 3 14.5h10a1.25 1.25 0 0 0 1.25-1.25v-8h1.25zM13 5.25v8H3v-8zM14.25 4H1.75V2.75h12.5z"/><path fill="currentColor" d="M7.38 8.38h1.24a.63.63 0 0 0 0-1.26H7.38a.63.63 0 0 0 0 1.26"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeft(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1.15 6.69l7.71-5.14l.7 1l-7.12 4.78H15.5v1.25H2.74l6.83 4.86l-.72 1L1.15 9A1.42 1.42 0 0 1 .5 7.83a1.4 1.4 0 0 1 .65-1.14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRight(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.85 6.69L7.14 1.55l-.7 1l7.12 4.75H.5v1.28h12.76l-6.83 4.86l.72 1L14.85 9a1.42 1.42 0 0 0 .65-1.15a1.4 1.4 0 0 0-.65-1.16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsHorizontal(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.25 6.12V4.87h-7l3-2.05l-.71-1L.59 4.45A1.29 1.29 0 0 0 0 5.5a1.29 1.29 0 0 0 .59 1l3.93 2.72l.71-1l-3-2.06zm6.16 3.33l-3.93-2.67l-.71 1l3 2.06h-7v1.25h7.03l-3 2l.71 1l3.93-2.67a1.23 1.23 0 0 0 0-2.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsVertical(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.12 6.75H4.87v7.05l-2.05-3l-1 .71l2.67 3.93a1.29 1.29 0 0 0 1 .59a1.29 1.29 0 0 0 1-.59l2.67-3.93l-1-.71l-2.06 3zM9.45.59L6.78 4.52l1 .71l2.06-3v7.02h1.25v-7l2 3l1-.71L11.55.59a1.23 1.23 0 0 0-2.1 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Backspace(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.6 2.5H4.49a1.25 1.25 0 0 0-1 .51L.39 7.26a1.26 1.26 0 0 0 0 1.48L3.48 13a1.26 1.26 0 0 0 1 .51H14.6a1.25 1.25 0 0 0 1.25-1.25V3.75A1.25 1.25 0 0 0 14.6 2.5m0 9.75H4.49L1.4 8l3.09-4.25H14.6z"/><path fill="currentColor" d="m7.86 10.55l1.99-1.72l1.99 1.72l.82-.94L10.81 8l1.85-1.61l-.82-.94l-1.99 1.72l-1.99-1.72l-.82.94L8.9 8L7.04 9.61z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Barcode(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 2.53h2.49v10.95H0zm11 0h2.49v10.95H11zm-6.02 0h1.24v10.95H4.98zm2.49 0h1.24v10.95H7.47zm7.29 0H16v10.95h-1.24z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bold(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.19 7.45A3.49 3.49 0 0 0 13 4.32C13 1.65 10.93 0 7.91 0H2.06v16h6.19c3.5 0 5.69-1.76 5.69-4.69a4 4 0 0 0-2.75-3.86m-7.88-6.2h4.6a4.36 4.36 0 0 1 2.86.87a2.71 2.71 0 0 1 1 2.2c0 .94-.43 2.51-3.36 2.51h-5.1zm8.31 12.54a5.3 5.3 0 0 1-3.37 1H3.31V8.08h4.35a7.29 7.29 0 0 1 3.68.84a2.62 2.62 0 0 1 1.35 2.39a3.05 3.05 0 0 1-1.07 2.48"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxPacked(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.12 4L8.62.85a1.28 1.28 0 0 0-1.24 0L1.88 4a1.25 1.25 0 0 0-.63 1.09V11a1.25 1.25 0 0 0 .63 1l5.5 3.11a1.28 1.28 0 0 0 1.24 0l5.5-3.11a1.25 1.25 0 0 0 .63-1V5.05A1.25 1.25 0 0 0 14.12 4m-6.74 9.71l-2.13-1.2v-5.3l2.13 1.16zM8 7.29L5.92 6.15l4.81-2.67l2.09 1.18zm0-5.35l1.46.82l-4.84 2.69l-1.44-.79zM2.5 5.71l1.5.82v5.27L2.5 11zm6.12 8V8.37l4.88-2.66V11z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxUnpacked(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.66 7l-.91-2.68L8.62.85a1.28 1.28 0 0 0-1.24 0L1.25 4.32L.34 7a1.24 1.24 0 0 0 .58 1.5l.33.18V11a1.25 1.25 0 0 0 .63 1l5.5 3.11a1.28 1.28 0 0 0 1.24 0l5.5-3.11a1.25 1.25 0 0 0 .63-1V8.68l.33-.18a1.24 1.24 0 0 0 .58-1.5M10 9.87l-.48-1.28L14 6.13l.44 1.28zM8 1.94L13.46 5L8 8L2.54 5zM1.52 7.41L2 6.13l4.48 2.46L6 9.87zm1 1.95l4.25 2.32l.62-1.84v3.87L2.5 11zM13.5 11l-4.88 2.71V9.84l.63 1.84l4.25-2.32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Briefcase(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 3.75h-3.49L11.11 2a1 1 0 0 0-1-.77H6A1 1 0 0 0 5 2l-.4 1.73H1a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1v-9a1 1 0 0 0-1-.98M6.17 2.5h3.76l.29 1.25H5.88zM14.75 5v2.5H1.25V5zm-13.5 8.5V8.75H6V9a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-.25h4.75v4.75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Browser(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<ellipse cx="3.72" cy="4.02" fill="currentColor" rx=".67" ry=".62"/><path fill="currentColor" d="M6.29 4.65A.65.65 0 0 0 7 4a.67.67 0 0 0-1.38 0a.65.65 0 0 0 .67.65"/><ellipse cx="8.87" cy="4.02" fill="currentColor" rx=".67" ry=".63"/><path fill="currentColor" d="M14.25 1.5H1.75A1.25 1.25 0 0 0 .5 2.75v10.5a1.25 1.25 0 0 0 1.25 1.25h12.5a1.25 1.25 0 0 0 1.25-1.25V2.75a1.25 1.25 0 0 0-1.25-1.25M1.75 2.75h12.5v2.5H1.75zm0 10.5V6.5h12.5v6.75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrowserSearch(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.24 13.25H1.75V6.5h4.13c-.07.27-.1.55-.1.84c0 2.04 1.77 3.7 3.94 3.7c.99 0 1.9-.35 2.59-.92l2.37 2.07l.82-.94l-2.37-2.07c.33-.54.53-1.17.53-1.84c0-2.04-1.77-3.7-3.94-3.7c-1.35 0-2.54.64-3.25 1.61H1.75v-2.5h12.49v.41h1.25v-.41c0-.69-.56-1.25-1.25-1.25H1.75C1.06 1.5.5 2.06.5 2.75v10.5c0 .69.56 1.25 1.25 1.25h12.49c.69 0 1.25-.56 1.25-1.25v-.41h-1.25zM9.72 4.89c1.48 0 2.69 1.1 2.69 2.45S11.2 9.79 9.72 9.79s-2.69-1.1-2.69-2.45s1.21-2.45 2.69-2.45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.25 2.5h-.77v-1h-1.25v1H3.68v-1H2.43v1h-.68A1.25 1.25 0 0 0 .5 3.75v9.5a1.25 1.25 0 0 0 1.25 1.25h12.5a1.25 1.25 0 0 0 1.25-1.25v-9.5a1.25 1.25 0 0 0-1.25-1.25M1.75 3.75h12.5V5H1.75zm0 9.5v-7h12.5v7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarDays(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.25 2.5h-.75v-1h-1.25v1h-8.5v-1H2.5v1h-.75A1.25 1.25 0 0 0 .5 3.75v9.5a1.25 1.25 0 0 0 1.25 1.25h12.5a1.25 1.25 0 0 0 1.25-1.25v-9.5a1.25 1.25 0 0 0-1.25-1.25M1.75 3.75h12.5V5H1.75zm0 9.5v-7h12.5v7z"/><path fill="currentColor" d="M3 8h1.1v1.2H3zm0 2.4h1.1v1.2H3zM11.8 8h1.1v1.2h-1.1zm0 2.4h1.1v1.2h-1.1zM9.6 8h1.1v1.2H9.6zm0 2.4h1.1v1.2H9.6zM7.4 8h1.1v1.2H7.4zm0 2.4h1.1v1.2H7.4zM5.2 8h1.1v1.2H5.2zm0 2.4h1.1v1.2H5.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Camera(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 3.5h-4l-.22-.65A1.24 1.24 0 0 0 9.6 2H6.4a1.24 1.24 0 0 0-1.18.85L5 3.5H1a1 1 0 0 0-1 1V13a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V4.5a1 1 0 0 0-1-1m-.25 9.25H1.25v-8H5.9l.5-1.5h3.2l.5 1.5h4.65z"/><path fill="currentColor" d="M8 6a2.59 2.59 0 0 0-2.67 2.5A2.59 2.59 0 0 0 8 11a2.59 2.59 0 0 0 2.67-2.5A2.59 2.59 0 0 0 8 6m0 3.75A1.35 1.35 0 0 1 6.58 8.5A1.35 1.35 0 0 1 8 7.25A1.35 1.35 0 0 1 9.42 8.5A1.35 1.35 0 0 1 8 9.75"/><ellipse cx="12.84" cy="6.63" fill="currentColor" rx=".67" ry=".63"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cash(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.75 2.5H1.25A1.25 1.25 0 0 0 0 3.75v8.5a1.25 1.25 0 0 0 1.25 1.25h13.5A1.25 1.25 0 0 0 16 12.25v-8.5a1.25 1.25 0 0 0-1.25-1.25M1.25 3.75h2a1.86 1.86 0 0 1-2 1.75zm0 8.5V10.5a1.86 1.86 0 0 1 2 1.75zm13.5 0H12.8a1.86 1.86 0 0 1 1.95-1.75zm0-3a3.1 3.1 0 0 0-3.2 3h-7.1a3.1 3.1 0 0 0-3.2-3v-2.5a3.1 3.1 0 0 0 3.2-3h7.1a3.1 3.1 0 0 0 3.2 3zm0-3.75a1.86 1.86 0 0 1-1.95-1.75h1.95z"/><path fill="currentColor" d="M8 5a3.1 3.1 0 0 0-3.2 3A3.1 3.1 0 0 0 8 11a3.1 3.1 0 0 0 3.2-3A3.1 3.1 0 0 0 8 5m0 4.75A1.86 1.86 0 0 1 6.05 8A1.86 1.86 0 0 1 8 6.25A1.86 1.86 0 0 1 10 8a1.86 1.86 0 0 1-2 1.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatDots(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 2.19c3.13 0 5.68 2.25 5.68 5s-2.55 5-5.68 5a5.65 5.65 0 0 1-1.89-.29l-.75-.26l-.56.56a13.78 13.78 0 0 1-2 1.55a.13.13 0 0 1-.07 0v-.06a6.58 6.58 0 0 0 .15-4.29a5.25 5.25 0 0 1-.55-2.16c0-2.77 2.55-5 5.68-5M8 .94c-3.83 0-6.93 2.81-6.93 6.27a6.42 6.42 0 0 0 .64 2.64a5.53 5.53 0 0 1-.18 3.48a1.32 1.32 0 0 0 2 1.5a14.93 14.93 0 0 0 2.16-1.71a6.78 6.78 0 0 0 2.31.36c3.83 0 6.93-2.81 6.93-6.27S11.83.94 8 .94"/><ellipse cx="5.2" cy="7.7" fill="currentColor" rx=".8" ry=".75"/><ellipse cx="8" cy="7.7" fill="currentColor" rx=".8" ry=".75"/><ellipse cx="10.8" cy="7.7" fill="currentColor" rx=".8" ry=".75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Check(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m15.12 2.23l-9.79 9.78L.88 7.56L0 8.44l5.33 5.34L16 3.11z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckCircle(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.53 9.02L4.58 7.07l-.88.89l2.83 2.83l.88-.89l4.78-4.77l-.89-.88z"/><path fill="currentColor" d="M8 .5A7.77 7.77 0 0 0 0 8a7.77 7.77 0 0 0 8 7.5A7.77 7.77 0 0 0 16 8A7.77 7.77 0 0 0 8 .5m0 13.75A6.52 6.52 0 0 1 1.25 8A6.52 6.52 0 0 1 8 1.75A6.52 6.52 0 0 1 14.75 8A6.52 6.52 0 0 1 8 14.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDown(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 10.18L2.39 4.52l-.89.87l5.59 5.71a1.18 1.18 0 0 0 .86.39a1.13 1.13 0 0 0 .85-.39l5.7-5.7l-.88-.89z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeft(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.82 8l5.66-5.56l-.87-.89L4.9 7.09a1.18 1.18 0 0 0-.39.91a1.13 1.13 0 0 0 .39.85l5.7 5.7l.89-.88z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRight(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.18 8.05l-5.66 5.56l.87.89l5.71-5.59a1.18 1.18 0 0 0 .39-.86a1.13 1.13 0 0 0-.39-.85L5.4 1.5l-.89.88z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUp(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8.05 5.82l5.56 5.66l.89-.87L8.91 4.9a1.18 1.18 0 0 0-.86-.39a1.13 1.13 0 0 0-.85.39l-5.7 5.7l.88.89z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Christ(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.32 3H10V2a2 2 0 0 0-4 0v1H.68a.69.69 0 0 0-.32 1.29L1.7 5v.58a1.25 1.25 0 0 0 .6 1.07L5 8.28v4H3.75A1.25 1.25 0 0 0 2.5 13.5V16h11v-2.5a1.25 1.25 0 0 0-1.25-1.25H11v-4l2.68-1.58a1.26 1.26 0 0 0 .62-1.08V5l1.34-.71A.69.69 0 0 0 15.32 3M5 6.79L3 5.58V4.25h2zM7.25 2a.75.75 0 0 1 1.5 0v1h-1.5zm-1 2.25H9.7v4.13a3.62 3.62 0 0 1-3.45-3.9zm0 3.75a3.59 3.59 0 0 0 .39.39a5.36 5.36 0 0 0 3.11 1.25v2.61h-3.5zm6 5.47v1.25h-8.5V13.5h8.5zm.8-7.93L11 6.78V4.25h2.05z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clock(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 .5A7.77 7.77 0 0 0 0 8a7.77 7.77 0 0 0 8 7.5A7.77 7.77 0 0 0 16 8A7.77 7.77 0 0 0 8 .5m0 13.75A6.52 6.52 0 0 1 1.25 8A6.52 6.52 0 0 1 8 1.75A6.52 6.52 0 0 1 14.75 8A6.52 6.52 0 0 1 8 14.25"/><path fill="currentColor" d="M8.63 4.25H7.38v4.38H11V7.38H8.63z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Close(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14.41 3.27l-.82-.94L8 7.17L2.41 2.33l-.82.94L7.05 8l-5.46 4.73l.82.94L8 8.83l5.59 4.84l.82-.94L8.95 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Code(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9.539.613l-3.9 14.55l1.209.324L10.746.937L9.54.612zm-5.63 3.434L.598 7.137a1.24 1.24 0 0 0 0 1.821l3.25 3.091l.921-.85l-3.321-3.15l3.321-3.112zm11.893 3.091l-3.31-3.091l-.861.91l3.32 3.091l-3.32 3.111l.92.85l3.251-3.05a1.242 1.242 0 0 0 0-1.821"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cog(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.84 11.49v-.1zm.23.38L1 11.68zM1 11.68l-.11-.19zm-.21-.29c-.06-.1-.11-.21-.16-.31c.05.1.1.21.16.31m.32.54v-.06z"/><path fill="currentColor" d="m11.26 12.63l1.83 1.09a7.34 7.34 0 0 0 1-.94a7.48 7.48 0 0 0 1.56-2.86l-1.74-1A5.29 5.29 0 0 0 14 8a5.29 5.29 0 0 0-.08-.9l1.74-1a7.45 7.45 0 0 0-1.33-2.58a7.54 7.54 0 0 0-1.24-1.22l-1.83 1.04a6 6 0 0 0-1.11-.53v-2A8.55 8.55 0 0 0 7.94.53a8.39 8.39 0 0 0-2.26.3v2a7.23 7.23 0 0 0-1.12.54L2.78 2.28A7.46 7.46 0 0 0 .2 6.06l1.72 1a5.29 5.29 0 0 0-.08.9a5.29 5.29 0 0 0 .08.9l-1.73 1a8 8 0 0 0 .43 1.15c.05.1.1.21.16.31v.1l.11.19l.12.19v.06a7.69 7.69 0 0 0 1.64 1.78l1.81-1.08a7.23 7.23 0 0 0 1.12.54v2a8.39 8.39 0 0 0 2.26.31a8.56 8.56 0 0 0 2.22-.3v-2a6 6 0 0 0 1.2-.48m-2.39 1.52a7.57 7.57 0 0 1-.95.06a7.73 7.73 0 0 1-1-.06v-1.69a4.92 4.92 0 0 1-2.53-1.27l-1.54.92a6.22 6.22 0 0 1-1.08-1.61l1.56-.93a4.27 4.27 0 0 1 0-3.17l-1.56-.92a6.11 6.11 0 0 1 1.12-1.62l1.56.93A5 5 0 0 1 7 3.53V1.82a7.73 7.73 0 0 1 1-.06a7.57 7.57 0 0 1 .95.06v1.72a4.9 4.9 0 0 1 2.4 1.26l1.59-.94a6.31 6.31 0 0 1 1.11 1.62l-1.6.94a4.35 4.35 0 0 1 .3 1.58a4.44 4.44 0 0 1-.29 1.55l1.56.93a6.43 6.43 0 0 1-1.11 1.61l-1.58-.93a5 5 0 0 1-2.49 1.28z"/><path fill="currentColor" d="M7.92 5.49A2.59 2.59 0 0 0 5.25 8a2.59 2.59 0 0 0 2.67 2.51A2.6 2.6 0 0 0 10.6 8a2.6 2.6 0 0 0-2.68-2.51M8 9.2A1.35 1.35 0 0 1 6.55 8A1.35 1.35 0 0 1 8 6.7A1.35 1.35 0 0 1 9.39 8A1.35 1.35 0 0 1 8 9.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ColorPalette(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 .5C3.58.5 0 3.86 0 8s3.58 7.5 8 7.5c4.69 0 1.04-2.83 2.79-4.55c.76-.75 1.63-.87 2.44-.87c.37 0 .73.03 1.06.03c.99 0 1.72-.23 1.72-2.1C16 3.86 12.42.5 8 .5m6.65 8.32c-.05.01-.16.02-.37.02c-.14 0-.29 0-.45-.01c-.19 0-.39-.01-.61-.01c-.89 0-2.19.13-3.32 1.23c-1.17 1.16-.9 2.6-.74 3.47c.03.18.08.44.09.6c-.16.05-.52.13-1.26.13c-3.72 0-6.75-2.8-6.75-6.25S4.28 1.75 8 1.75s6.75 2.8 6.75 6.25c0 .5-.06.74-.1.82"/><path fill="currentColor" d="M5.9 9.47c-1.03 0-1.86.8-1.86 1.79s.84 1.79 1.86 1.79s1.86-.8 1.86-1.79s-.84-1.79-1.86-1.79m0 2.35c-.35 0-.64-.25-.64-.56s.29-.56.64-.56s.64.25.64.56s-.29.56-.64.56m-.2-4.59c0-.99-.84-1.79-1.86-1.79s-1.86.8-1.86 1.79s.84 1.79 1.86 1.79s1.86-.8 1.86-1.79m-1.86.56c-.35 0-.64-.25-.64-.56s.29-.56.64-.56s.64.25.64.56s-.29.56-.64.56M7.37 2.5c-1.03 0-1.86.8-1.86 1.79s.84 1.79 1.86 1.79s1.86-.8 1.86-1.79S8.39 2.5 7.37 2.5m0 2.35c-.35 0-.64-.25-.64-.56s.29-.56.64-.56s.64.25.64.56s-.29.56-.64.56m2.47 1.31c0 .99.84 1.79 1.86 1.79s1.86-.8 1.86-1.79s-.84-1.79-1.86-1.79s-1.86.8-1.86 1.79m2.5 0c0 .31-.29.56-.64.56s-.64-.25-.64-.56s.29-.56.64-.56s.64.25.64.56"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copy(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.49 3L10.74.37A1.22 1.22 0 0 0 9.86 0h-4a1.25 1.25 0 0 0-1.22 1.25v11a1.25 1.25 0 0 0 1.25 1.25h6.72a1.25 1.25 0 0 0 1.25-1.25V3.88a1.22 1.22 0 0 0-.37-.88m-.88 9.25H5.89v-11h2.72v2.63a1.25 1.25 0 0 0 1.25 1.25h2.75zm0-8.37H9.86V1.25z"/><path fill="currentColor" d="M10.11 14.75H3.39v-11H4V2.5h-.61a1.25 1.25 0 0 0-1.25 1.25v11A1.25 1.25 0 0 0 3.39 16h6.72a1.25 1.25 0 0 0 1.25-1.25v-.63h-1.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCard(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.55 9.75h2V11h-2z"/><path fill="currentColor" d="M15.32 2.5H.68a.69.69 0 0 0-.68.71v9.58a.69.69 0 0 0 .68.71h14.64a.69.69 0 0 0 .68-.71V3.21a.69.69 0 0 0-.68-.71m-.57 1.25V5H1.25V3.75zm-13.5 8.5V7.5h13.5v4.75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Desktop(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.25 1.5H1.75A1.25 1.25 0 0 0 .5 2.75v7.5a1.25 1.25 0 0 0 1.25 1.25H6.1v1.75H4v1.25h8v-1.25H9.91V11.5h4.34a1.25 1.25 0 0 0 1.25-1.25v-7.5a1.25 1.25 0 0 0-1.25-1.25M8.66 13.25H7.35V11.5h1.31zm5.59-3H1.75v-7.5h12.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiscountCircle(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.09 4.37a.62.62 0 0 0-.85.23l-3.57 6.18a.62.62 0 0 0 .23.85a.55.55 0 0 0 .31.08a.63.63 0 0 0 .54-.31l3.57-6.18a.61.61 0 0 0-.23-.85M7.35 6.43a2.16 2.16 0 0 0-2.06-2.24a2.17 2.17 0 0 0-2.07 2.24a2.17 2.17 0 0 0 2.07 2.25a2.16 2.16 0 0 0 2.06-2.25m-2.06 1a.93.93 0 0 1-.82-1a.92.92 0 0 1 .82-1a.91.91 0 0 1 .81 1a.92.92 0 0 1-.81 1m5.42-.11a2.16 2.16 0 0 0-2.06 2.24a2.16 2.16 0 0 0 2.06 2.25a2.17 2.17 0 0 0 2.07-2.25a2.17 2.17 0 0 0-2.07-2.24m0 3.24a.92.92 0 0 1-.81-1a.92.92 0 0 1 .81-1a.92.92 0 0 1 .82 1a.93.93 0 0 1-.82 1"/><path fill="currentColor" d="M8 .5A7.76 7.76 0 0 0 0 8a7.76 7.76 0 0 0 8 7.5A7.76 7.76 0 0 0 16 8A7.76 7.76 0 0 0 8 .5m0 13.75A6.52 6.52 0 0 1 1.25 8A6.52 6.52 0 0 1 8 1.75A6.52 6.52 0 0 1 14.75 8A6.52 6.52 0 0 1 8 14.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Diskette(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.18 6.14l-3.81-4.23a1.24 1.24 0 0 0-.93-.41H1.75A1.25 1.25 0 0 0 .5 2.75v10.5a1.25 1.25 0 0 0 1.25 1.25h12.5a1.25 1.25 0 0 0 1.25-1.25V7a1.24 1.24 0 0 0-.32-.86m-6-3.39v2H4.25v-2zm5.05 10.5H1.75V2.75H3v2A1.25 1.25 0 0 0 4.25 6H9.2a1.25 1.25 0 0 0 1.25-1.25v-2L14.25 7z"/><path fill="currentColor" d="M8 7.56a2.43 2.43 0 0 0-2.5 2.35A2.43 2.43 0 0 0 8 12.25a2.43 2.43 0 0 0 2.5-2.34A2.43 2.43 0 0 0 8 7.56M8 11a1.18 1.18 0 0 1-1.25-1.09A1.18 1.18 0 0 1 8 8.81a1.18 1.18 0 0 1 1.25 1.1A1.18 1.18 0 0 1 8 11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Download(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.5 14.75h11V16h-11zm.39-5.45l4.27 3.88a1.26 1.26 0 0 0 1.68 0l4.27-3.88a1.25 1.25 0 0 0-.84-2.18H9.8V1.25A1.25 1.25 0 0 0 8.55 0h-1.1A1.25 1.25 0 0 0 6.2 1.25v5.87H3.73a1.25 1.25 0 0 0-.84 2.18m.84-.93h3.72V1.25h1.1v7.12h3.72L8 12.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Drag(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.46 7l-3.2-2.19l-.71 1l2.29 1.57H8.62V2.16l1.57 2.29l1-.71L9 .54a1.25 1.25 0 0 0-2 0l-2.22 3.2l1 .71l1.59-2.29v5.22H2.16l2.29-1.57l-.71-1L.54 7a1.25 1.25 0 0 0 0 2l3.2 2.19l.71-1l-2.29-1.57h5.21v5.22l-1.56-2.29l-1 .71L7 15.46a1.25 1.25 0 0 0 2.06 0l2.19-3.2l-1-.71l-1.63 2.29V8.62h5.22l-2.29 1.57l.71 1L15.46 9a1.25 1.25 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DragDots(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.375 3.67c0-.645-.56-1.17-1.25-1.17s-1.25.525-1.25 1.17c0 .646.56 1.17 1.25 1.17s1.25-.524 1.25-1.17m0 8.66c0-.646-.56-1.17-1.25-1.17s-1.25.524-1.25 1.17c0 .645.56 1.17 1.25 1.17s1.25-.525 1.25-1.17m-1.25-5.5c.69 0 1.25.525 1.25 1.17c0 .645-.56 1.17-1.25 1.17S4.875 8.645 4.875 8c0-.645.56-1.17 1.25-1.17m5-3.16c0-.645-.56-1.17-1.25-1.17s-1.25.525-1.25 1.17c0 .646.56 1.17 1.25 1.17s1.25-.524 1.25-1.17m-1.25 7.49c.69 0 1.25.524 1.25 1.17c0 .645-.56 1.17-1.25 1.17s-1.25-.525-1.25-1.17c0-.646.56-1.17 1.25-1.17M11.125 8c0-.645-.56-1.17-1.25-1.17s-1.25.525-1.25 1.17c0 .645.56 1.17 1.25 1.17s1.25-.525 1.25-1.17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Drink(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="4.96" cy="9.98" r=".63" fill="currentColor"/><circle cx="5.9" cy="8.55" r=".31" fill="currentColor"/><path fill="currentColor" fill-rule="evenodd" d="M15.15 9v-.09a2.45 2.45 0 0 0 0-.27v-.33a2.21 2.21 0 0 1-.15-.25v-.27c0-.09 0-.17-.08-.25a1.43 1.43 0 0 1-.08-.22v-.05a1.6 1.6 0 0 1-.1-.23a1.86 1.86 0 0 0-.1-.2l-.12-.23v-.09a6.41 6.41 0 0 0-1-1.24h-2.75l.65-2.44a.67.67 0 0 1 .32-.4a.61.61 0 0 1 .51 0l1.87.71l.45-1.16L12.7 1a1.88 1.88 0 0 0-2.49 1.27l-.73 2.78H7.31a3.15 3.15 0 0 0-3.23-2.86a3.13 3.13 0 0 0-3.23 3A3 3 0 0 0 2.79 8a2.21 2.21 0 0 1 0 .25v.31a2.45 2.45 0 0 0 0 .27v.45A6.06 6.06 0 0 0 8.9 15.1a6.06 6.06 0 0 0 6.25-5.85zM8.9 13.85a4.81 4.81 0 0 1-5-4.6A4.41 4.41 0 0 1 5.06 6.3h7.69a4.4 4.4 0 0 1 1.15 3a4.81 4.81 0 0 1-5 4.55M2.1 5.23a1.89 1.89 0 0 1 2-1.79a1.93 1.93 0 0 1 2 1.61H4.53a6.07 6.07 0 0 0-1 1.24a.56.56 0 0 1 0 .08l-.12.23a1.42 1.42 0 0 0-.08.18A1.8 1.8 0 0 1 2.1 5.23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Drop(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 12.09a1.74 1.74 0 0 1-1.25-.49a1.1 1.1 0 0 1-.36-.8a.63.63 0 0 0-.63-.63a.62.62 0 0 0-.62.63a2.33 2.33 0 0 0 .73 1.69a3 3 0 0 0 2.13.85a.63.63 0 1 0 0-1.25"/><path fill="currentColor" d="M9.25 1.25C9.06.23 8.43 0 8 0S6.94.23 6.75 1.25C6 5.2 2.5 7 2.5 10.82A5.35 5.35 0 0 0 8 16a5.35 5.35 0 0 0 5.5-5.18C13.5 7 10 5.2 9.25 1.25M8 14.75a4.1 4.1 0 0 1-4.25-3.93c0-1.66.85-2.9 1.84-4.33A12.85 12.85 0 0 0 8 1.48v-.1v.1a12.85 12.85 0 0 0 2.39 5c1 1.43 1.84 2.67 1.84 4.33A4.1 4.1 0 0 1 8 14.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Drums(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.68 4.8c-.16-.07-.33-.13-.51-.19l1.91-3.3L12 .69L9.9 4.31A11.61 11.61 0 0 0 8 4.16a11.88 11.88 0 0 0-1.82.14L4.1.69L3 1.31l1.91 3.28a4.53 4.53 0 0 0-.59.21c-1.58.64-1.82 1.52-1.82 2v5.9c0 .46.24 1.34 1.82 2a9.94 9.94 0 0 0 3.68.61a9.94 9.94 0 0 0 3.68-.63c1.58-.65 1.82-1.53 1.82-2V6.79c0-.47-.24-1.35-1.82-1.99M4.79 6a6.77 6.77 0 0 1 .76-.25l.79 1.37l1.08-.63l-.56-1A10.6 10.6 0 0 1 8 5.41a11.67 11.67 0 0 1 1.23.07l-.56 1l1.08.63l.78-1.34c1.1.3 1.72.76 1.72 1.06s-.32.54-1 .83a8.72 8.72 0 0 1-3.25.5c-2.64 0-4.25-.89-4.25-1.37c0-.21.32-.54 1.04-.79m6.42 7.57a8.72 8.72 0 0 1-3.21.49c-2.64 0-4.25-.89-4.25-1.37v-1.23a6.18 6.18 0 0 0 .57.27a9.94 9.94 0 0 0 3.68.63a9.94 9.94 0 0 0 3.68-.63a6.18 6.18 0 0 0 .57-.27v1.23c0 .2-.32.54-1.04.83zm0-2.95a8.72 8.72 0 0 1-3.21.49c-2.64 0-4.25-.89-4.25-1.37V8.51a6.18 6.18 0 0 0 .57.27A9.94 9.94 0 0 0 8 9.41a9.94 9.94 0 0 0 3.68-.63a6.18 6.18 0 0 0 .57-.27v1.23c0 .2-.32.54-1.04.83z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Duplicate(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.869 6.092V3.955h1.25v2.137h2.118v1.25h-2.118v2.213h-1.25V7.342H6.636v-1.25z"/><path fill="currentColor" fill-rule="evenodd" d="M15.5 10.76c0 .69-.56 1.25-1.25 1.25H4.623c-.69 0-1.25-.56-1.25-1.25V2.75c0-.69.56-1.25 1.25-1.25h9.627c.69 0 1.25.56 1.25 1.25zm-1.25 0H4.623V2.75h9.627z"/><path fill="currentColor" d="M1.75 13.25h9.615v-.62h1.25v.62c0 .69-.56 1.25-1.25 1.25H1.75c-.69 0-1.25-.56-1.25-1.25v-8C.5 4.56 1.06 4 1.75 4h1.005v1.25H1.75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ecosystem(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 8a2.64 2.64 0 0 0-.53-1.59l1.33-1.32a2.22 2.22 0 0 0 .82.17a1.94 1.94 0 0 0 2-1.88a2 2 0 0 0-4 0a1.76 1.76 0 0 0 .24.88L9.53 5.59A3.14 3.14 0 0 0 8 5.19a3 3 0 0 0-2.93 2.19h-.69A2 2 0 0 0 2.5 6.13A1.94 1.94 0 0 0 .5 8a1.94 1.94 0 0 0 2 1.88a2 2 0 0 0 1.88-1.25h.69A3 3 0 0 0 8 10.82A2.91 2.91 0 0 0 11 8m1.62-5.24a.69.69 0 0 1 .75.62a.76.76 0 0 1-1.5 0a.7.7 0 0 1 .75-.62M2.5 8.63A.7.7 0 0 1 1.75 8a.7.7 0 0 1 .75-.62a.7.7 0 0 1 .75.62a.7.7 0 0 1-.75.63m5.5.94A1.67 1.67 0 0 1 6.25 8A1.66 1.66 0 0 1 8 6.44A1.67 1.67 0 0 1 9.75 8A1.68 1.68 0 0 1 8 9.57"/><path fill="currentColor" d="M2.5 4.38a2 2 0 0 0 .82-.17L5.08 6A3.73 3.73 0 0 1 6 5.13L4.26 3.38a1.76 1.76 0 0 0 .24-.88a1.94 1.94 0 0 0-2-1.87a1.94 1.94 0 0 0-2 1.87a1.94 1.94 0 0 0 2 1.88m0-2.5a.7.7 0 0 1 .75.62a.7.7 0 0 1-.75.63a.7.7 0 0 1-.75-.63a.7.7 0 0 1 .75-.62m11 9.75a2 2 0 0 0-.82.17L10.92 10a3.73 3.73 0 0 1-.93.84l1.74 1.74a1.75 1.75 0 0 0-.23.88a2 2 0 0 0 4 0a1.94 1.94 0 0 0-2-1.83m0 2.5a.7.7 0 0 1-.75-.63a.76.76 0 0 1 1.5 0a.7.7 0 0 1-.75.63"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Edit(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.49 7.3h-1.16v6.35H1.67V3.28H8V2H1.67A1.21 1.21 0 0 0 .5 3.28v10.37a1.21 1.21 0 0 0 1.17 1.25h12.66a1.21 1.21 0 0 0 1.17-1.25z"/><path fill="currentColor" d="M10.56 2.87L6.22 7.22l-.44.44l-.08.08l-1.52 3.16a1.08 1.08 0 0 0 1.45 1.45l3.14-1.53l.53-.53l.43-.43l4.34-4.36l.45-.44l.25-.25a2.18 2.18 0 0 0 0-3.08a2.17 2.17 0 0 0-1.53-.63a2.19 2.19 0 0 0-1.54.63l-.7.69l-.45.44zM5.51 11l1.18-2.43l1.25 1.26zm2-3.36l3.9-3.91l1.3 1.31L8.85 9zm5.68-5.31a.91.91 0 0 1 .65.27a.93.93 0 0 1 0 1.31l-.25.24l-1.3-1.3l.25-.25a.88.88 0 0 1 .69-.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ellipsis(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 2.5a1.22 1.22 0 0 1 1.25 1.17A1.21 1.21 0 0 1 8 4.84a1.21 1.21 0 0 1-1.25-1.17A1.22 1.22 0 0 1 8 2.5m0 8.66a1.17 1.17 0 1 1-1.25 1.17A1.21 1.21 0 0 1 8 11.16m0-4.33a1.17 1.17 0 1 1 0 2.34a1.17 1.17 0 1 1 0-2.34"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationCircle(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.328 8c0 3.377-2.76 6.25-6.328 6.25c-3.567 0-6.328-2.873-6.328-6.25S4.432 1.75 8 1.75c3.567 0 6.328 2.873 6.328 6.25M15.5 8a7.5 7.5 0 1 1-15 0a7.5 7.5 0 0 1 15 0M8.59 3.101l-.013 7.3l-1.172-.002l.013-7.3zm.015 9.249a.673.673 0 0 0-.179-.46a.59.59 0 0 0-.43-.19a.59.59 0 0 0-.432.19a.673.673 0 0 0-.178.46c0 .172.064.338.178.46a.59.59 0 0 0 .431.19a.59.59 0 0 0 .431-.19a.673.673 0 0 0 .179-.46" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationTriangle(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.35 12.81L9 2.08a1.22 1.22 0 0 0-2 0L.65 12.81a1.14 1.14 0 0 0 1 1.69h12.66a1.14 1.14 0 0 0 1.04-1.69m-13.66.55L8 2.64l6.31 10.72z"/><path fill="currentColor" d="M7.32 5.45h1.25V10H7.32z"/><ellipse cx="7.95" cy="11.9" fill="currentColor" rx=".67" ry=".7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExternalLink(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.21 1.5H10v1.25h3.08L7.9 7.21l.82 1l5.53-4.77V7h1.25V2.79a1.29 1.29 0 0 0-1.29-1.29"/><path fill="currentColor" d="M12.25 13.25H1.75v-8.5H7.5V3.5h-6a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-4h-1.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eye(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 5.5A2.59 2.59 0 0 0 5.33 8A2.59 2.59 0 0 0 8 10.5A2.59 2.59 0 0 0 10.67 8A2.59 2.59 0 0 0 8 5.5m0 3.75A1.35 1.35 0 0 1 6.58 8A1.35 1.35 0 0 1 8 6.75A1.35 1.35 0 0 1 9.42 8A1.35 1.35 0 0 1 8 9.25"/><path fill="currentColor" d="M8 2.5A8.11 8.11 0 0 0 0 8a8.11 8.11 0 0 0 8 5.5A8.11 8.11 0 0 0 16 8a8.11 8.11 0 0 0-8-5.5m5.4 7.5A6.91 6.91 0 0 1 8 12.25A6.91 6.91 0 0 1 2.6 10a7.2 7.2 0 0 1-1.27-2A7.2 7.2 0 0 1 2.6 6A6.91 6.91 0 0 1 8 3.75A6.91 6.91 0 0 1 13.4 6a7.2 7.2 0 0 1 1.27 2a7.2 7.2 0 0 1-1.27 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeOff(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 2.5a9.77 9.77 0 0 0-2.53.32l1 1.05A8.78 8.78 0 0 1 8 3.75A6.91 6.91 0 0 1 13.4 6a7.2 7.2 0 0 1 1.27 2a7.2 7.2 0 0 1-1.27 2c-.12.13-.24.26-.37.38l.89.89A8.24 8.24 0 0 0 16 8a8.11 8.11 0 0 0-8-5.5m5 9.56l-.9-.9l-6.97-6.91l-1-1l-1.19-1.19l-.88.88l1 1A8.25 8.25 0 0 0 0 8a8.11 8.11 0 0 0 8 5.5a9.05 9.05 0 0 0 3.82-.79l1.24 1.23l.88-.88l-1-1zM6.66 7.54l1.67 1.67a1.47 1.47 0 0 1-.36 0A1.35 1.35 0 0 1 6.55 8a1.07 1.07 0 0 1 .11-.46M8 12.25A6.91 6.91 0 0 1 2.6 10a7.2 7.2 0 0 1-1.27-2A7.2 7.2 0 0 1 2.6 6A6.49 6.49 0 0 1 4 4.84l1.74 1.79A2.33 2.33 0 0 0 5.3 8A2.59 2.59 0 0 0 8 10.5a2.78 2.78 0 0 0 1.32-.33l1.58 1.58a8 8 0 0 1-2.9.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func File(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.13 4.13L9.37.37A1.26 1.26 0 0 0 8.48 0H3.75A1.25 1.25 0 0 0 2.5 1.25v13.5A1.25 1.25 0 0 0 3.75 16h8.5a1.25 1.25 0 0 0 1.25-1.25V5a1.26 1.26 0 0 0-.37-.87m-.88 10.62h-8.5V1.25h3.48V5a1.25 1.25 0 0 0 1.25 1.27h3.77zm0-9.73H8.48V1.25L12.25 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileAlt(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.13 4.13L9.37.37A1.26 1.26 0 0 0 8.48 0H3.75A1.25 1.25 0 0 0 2.5 1.25v13.5A1.25 1.25 0 0 0 3.75 16h8.5a1.25 1.25 0 0 0 1.25-1.25V5a1.26 1.26 0 0 0-.37-.87m-.88 10.62h-8.5V1.25h3.48V5a1.25 1.25 0 0 0 1.25 1.27h3.77zm0-9.73H8.48V1.25L12.25 5z"/><path fill="currentColor" d="M5 7.5h6v1.25H5zM5 10h6v1.25H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fingerprint(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.57 4.87a.62.62 0 0 0-.05.88c1 1.11 1.1 4.79.51 7.22a.63.63 0 0 0 .46.76h.15a.63.63 0 0 0 .61-.47c.6-2.48.65-6.75-.8-8.36a.63.63 0 0 0-.88-.03M8 4.1a4.67 4.67 0 0 1 1.89.39a.62.62 0 1 0 .5-1.14A5.88 5.88 0 0 0 8 2.85C5 2.85 2.57 5 2.57 7.64a13.5 13.5 0 0 1-.49 4.65a.62.62 0 0 0 .34.81a.52.52 0 0 0 .24.05a.63.63 0 0 0 .58-.38a14.34 14.34 0 0 0 .58-5.13C3.82 5.69 5.7 4.1 8 4.1"/><path fill="currentColor" d="M13.6 2.44A8.27 8.27 0 0 0 8 .29a8.5 8.5 0 0 0-3.7.84A.62.62 0 0 0 4 2a.63.63 0 0 0 .84.29A7.38 7.38 0 0 1 8 1.54c3.65 0 6.73 2.79 6.73 6.1a.63.63 0 0 0 .63.62a.62.62 0 0 0 .64-.62a7 7 0 0 0-2.4-5.2M2.16 3.54a.64.64 0 0 0-.87.16A7.05 7.05 0 0 0 0 7.63a.63.63 0 0 0 .62.63a.62.62 0 0 0 .62-.62a5.74 5.74 0 0 1 1.08-3.23a.62.62 0 0 0-.16-.87M8 5.41a2.9 2.9 0 0 0-2.2.8a.61.61 0 0 0 0 .88a.62.62 0 0 0 .88 0A1.83 1.83 0 0 1 8 6.66a1.62 1.62 0 0 1 1.65 1.18c.33.93 0 4.93-.55 6.4a.63.63 0 0 0 .38.8a.78.78 0 0 0 .21 0a.62.62 0 0 0 .59-.42A18 18 0 0 0 11 11a10.9 10.9 0 0 0-.12-3.53A2.87 2.87 0 0 0 8 5.41M5.68 8.34a.62.62 0 0 0-.57.66a11.09 11.09 0 0 1-.79 4.65a.62.62 0 0 0 .18.86a.6.6 0 0 0 .34.11a.61.61 0 0 0 .52-.29a11.93 11.93 0 0 0 1-5.43a.62.62 0 0 0-.68-.56"/><path fill="currentColor" d="M7.9 7.66a.63.63 0 0 0-.52.72a15.66 15.66 0 0 1-.78 5.95a.63.63 0 0 0 .29.84a.63.63 0 0 0 .83-.29a16.6 16.6 0 0 0 .9-6.71a.62.62 0 0 0-.72-.51"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fire(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.32 15.653a.812.812 0 0 1-.086-.855c.176-.342.245-.733.2-1.118a2.106 2.106 0 0 0-.267-.779a2.027 2.027 0 0 0-.541-.606a3.96 3.96 0 0 1-1.481-2.282c-1.708 2.239-1.053 3.51-.235 4.63a.748.748 0 0 1-.014.901a.87.87 0 0 1-.394.283a.838.838 0 0 1-.478.023c-1.105-.27-2.145-.784-2.85-1.603a4.686 4.686 0 0 1-.906-1.555a4.811 4.811 0 0 1-.263-1.797s-.133-2.463 2.837-4.876c0 0 3.51-2.978 2.292-5.18a.621.621 0 0 1 .112-.653a.558.558 0 0 1 .623-.147l.146.058a7.63 7.63 0 0 1 2.96 3.5c.58 1.413.576 3.06.184 4.527c.325-.292.596-.641.801-1.033l.029-.064c.198-.477.821-.325 1.055-.013c.086.137 2.292 3.343 1.107 6.048a5.516 5.516 0 0 1-1.84 2.027a6.127 6.127 0 0 1-2.138.893a.834.834 0 0 1-.472-.038a.867.867 0 0 1-.381-.29zM7.554 7.892a.422.422 0 0 1 .55.146c.04.059.066.126.075.198l.045.349c.02.511.014 1.045.213 1.536c.206.504.526.95.932 1.298a3.06 3.06 0 0 1 1.16 1.422c.22.564.25 1.19.084 1.773a4.123 4.123 0 0 0 1.39-.757l.103-.084c.336-.277.613-.623.813-1.017c.201-.393.322-.825.354-1.269c.065-1.025-.284-2.054-.827-2.972c-.248.36-.59.639-.985.804c-.247.105-.509.17-.776.19a.792.792 0 0 1-.439-.1a.832.832 0 0 1-.321-.328a.825.825 0 0 1-.035-.729c.412-.972.54-2.05.365-3.097a5.874 5.874 0 0 0-1.642-3.16c-.156 2.205-2.417 4.258-2.881 4.7a3.537 3.537 0 0 1-.224.194c-2.426 1.965-2.26 3.755-2.26 3.834a3.678 3.678 0 0 0 .459 2.043c.365.645.89 1.177 1.52 1.54C4.5 12.808 4.5 10.89 7.183 8.14l.372-.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flag(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.27 1.48a1.3 1.3 0 0 0-.65.18a6.32 6.32 0 0 1-3 1.1c-1.66 0-3.65-1-5.65-1a11.41 11.41 0 0 0-3.18.46v-.09a.63.63 0 0 0-1.25 0v11.76a.63.63 0 1 0 1.25 0V11a10.56 10.56 0 0 1 3.18-.46c1.66 0 4.4 1 6.45 1A5.86 5.86 0 0 0 15 10.11a1.17 1.17 0 0 0 .47-.93V2.66a1.21 1.21 0 0 0-1.2-1.18m0 7.65a4.58 4.58 0 0 1-2.87 1.08a17.73 17.73 0 0 1-3.29-.49a16 16 0 0 0-3.16-.48A12.3 12.3 0 0 0 2 9.57v-6.1A9.85 9.85 0 0 1 4.93 3a11.59 11.59 0 0 1 2.78.48a11.9 11.9 0 0 0 2.87.52a7.5 7.5 0 0 0 3.67-1.27z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Font(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.51 2.6L.25 13.67h1.34l1.49-3.86h4l1.52 3.86h1.34L5.68 2.6a.63.63 0 0 0-1.17 0m-.95 6l1.54-4l1.53 4zm9.35-2.54a2.8 2.8 0 0 0-3 2.08l1.21.31a1.6 1.6 0 0 1 1.78-1.14c.77 0 1.59.26 1.59 1v.75c-.27 0-.63.09-.94.13a9.12 9.12 0 0 0-2.5.52a2.06 2.06 0 0 0-1.41 2.23a1.94 1.94 0 0 0 .88 1.44a3 3 0 0 0 1.62.43a4.39 4.39 0 0 0 1.36-.22a2.92 2.92 0 0 0 1-.52v.61h1.25V8.3c0-1.3-1.14-2.24-2.84-2.24m.22 6.33a2.4 2.4 0 0 1-1.91-.07a.64.64 0 0 1-.32-.52c-.1-.89.82-1.16 2.8-1.38l.76-.1c-.19 1.68-.94 1.94-1.33 2.07"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Forbidden(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 .5A7.77 7.77 0 0 0 0 8a7.77 7.77 0 0 0 8 7.5A7.77 7.77 0 0 0 16 8A7.77 7.77 0 0 0 8 .5M1.25 8A6 6 0 0 1 3 3.85L12.09 13A7.12 7.12 0 0 1 8 14.25A6.52 6.52 0 0 1 1.25 8M13 12.15L3.91 3A7.12 7.12 0 0 1 8 1.75A6.52 6.52 0 0 1 14.75 8A6 6 0 0 1 13 12.15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GiftBox(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.5 6a1.25 1.25 0 0 0-1.25-1.25h-1.06a2.19 2.19 0 0 0 .23-1A2.22 2.22 0 0 0 11.2 1.5A3.78 3.78 0 0 0 8 3.07A3.78 3.78 0 0 0 4.8 1.5a2.22 2.22 0 0 0-2.22 2.22a2.19 2.19 0 0 0 .23 1H1.75A1.25 1.25 0 0 0 .5 6v2.5h1.25v4.79A1.25 1.25 0 0 0 3 14.5h10a1.25 1.25 0 0 0 1.25-1.25V8.46h1.25zm-4.3-3.25a1 1 0 0 1 0 2H8.68a2.36 2.36 0 0 1 2.52-2m-6.4 0a2.36 2.36 0 0 1 2.52 2H4.8a1 1 0 0 1 0-1.95zM1.75 6h4.37v1.21H1.75zm4.37 7.29H3V8.46h3.12zm1.26 0V6h1.24v7.29zm5.62 0H9.88V8.46H13zm1.25-6H9.88V6h4.37z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GiftCard(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.5 11.25h2v1.25h-2z"/><path fill="currentColor" d="M14.75 4H13.3a2.1 2.1 0 0 0-1.87-3a4.24 4.24 0 0 0-2.84 1a3.42 3.42 0 0 0-.59.67A3.42 3.42 0 0 0 7.41 2a4.24 4.24 0 0 0-2.84-1A2.1 2.1 0 0 0 2.7 4H1.25A1.25 1.25 0 0 0 0 5.25v8.5A1.25 1.25 0 0 0 1.25 15h13.5A1.25 1.25 0 0 0 16 13.75v-8.5A1.25 1.25 0 0 0 14.75 4M9.42 2.94a3 3 0 0 1 2-.69a.85.85 0 0 1 0 1.7H8.74a2.32 2.32 0 0 1 .68-1.01m-4.85-.69a3 3 0 0 1 2 .69a2.32 2.32 0 0 1 .68 1H4.57a.85.85 0 0 1 0-1.7zm-3.32 3h13.5V6.5H1.25zm0 8.5V9h13.5v4.75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Glasses(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.86 9.16l-1.43-5.31a1.88 1.88 0 0 0-1.81-1.37h-.44v1.26h.44a.63.63 0 0 1 .6.46L14 7.06a3.9 3.9 0 0 0-1.71-.39a3.64 3.64 0 0 0-3.7 3H7.42a3.64 3.64 0 0 0-3.7-3A3.87 3.87 0 0 0 2 7.06l.78-2.86a.63.63 0 0 1 .6-.46h.44V2.48h-.44a1.88 1.88 0 0 0-1.81 1.37L.14 9.16a3.19 3.19 0 0 0-.14.94a3.59 3.59 0 0 0 3.72 3.42a3.71 3.71 0 0 0 3.62-2.59h1.32a3.71 3.71 0 0 0 3.62 2.59A3.59 3.59 0 0 0 16 10.1a3.19 3.19 0 0 0-.14-.94m-12.14 3.1a2.33 2.33 0 0 1-2.46-2.16a2.33 2.33 0 0 1 2.46-2.17a2.34 2.34 0 0 1 2.47 2.17a2.34 2.34 0 0 1-2.47 2.16m8.56 0a2.34 2.34 0 0 1-2.47-2.16a2.34 2.34 0 0 1 2.47-2.17a2.33 2.33 0 0 1 2.46 2.17a2.33 2.33 0 0 1-2.46 2.16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Globe(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 7.38A7.82 7.82 0 0 0 8 .5a7.82 7.82 0 0 0-8 6.88v1.24a7.82 7.82 0 0 0 8 6.88a7.82 7.82 0 0 0 8-6.88zm-1.25 0h-3a11.34 11.34 0 0 0-.43-2.54a7.6 7.6 0 0 0 1.75-1a6 6 0 0 1 1.65 3.54zm-9.18 0a9.69 9.69 0 0 1 .37-2.14A8.43 8.43 0 0 0 8 5.5a8.49 8.49 0 0 0 2.09-.26a10.2 10.2 0 0 1 .37 2.14zm4.92 1.24a9.59 9.59 0 0 1-.37 2.14a8.53 8.53 0 0 0-4.18 0a9.69 9.69 0 0 1-.37-2.14zm.4-5A11.82 11.82 0 0 0 10 2a6.89 6.89 0 0 1 2 1a6.57 6.57 0 0 1-1.14.66zm-2.6-1.86a10 10 0 0 1 1.38 2.3A7.63 7.63 0 0 1 8 4.25a7.56 7.56 0 0 1-1.67-.19a9.82 9.82 0 0 1 1.38-2.3zm-3.15 1.9A6.57 6.57 0 0 1 4 3a6.89 6.89 0 0 1 2-1a10.38 10.38 0 0 0-.86 1.66M3 3.83a7.6 7.6 0 0 0 1.75 1a11 11 0 0 0-.43 2.54h-3A6 6 0 0 1 3 3.83M1.28 8.62h3a11 11 0 0 0 .43 2.54a7.6 7.6 0 0 0-1.75 1a6 6 0 0 1-1.68-3.54m3.86 3.72A10.38 10.38 0 0 0 6 14a6.89 6.89 0 0 1-2-1a6.57 6.57 0 0 1 1.14-.66m2.57 1.9a9.82 9.82 0 0 1-1.38-2.3a7.43 7.43 0 0 1 3.34 0a9.76 9.76 0 0 1-1.38 2.3zm3.15-1.9a6.57 6.57 0 0 1 1.19.66a7.24 7.24 0 0 1-2 1a11.48 11.48 0 0 0 .81-1.66m2.14-.17a7.6 7.6 0 0 0-1.75-1a10.8 10.8 0 0 0 .43-2.54h3A6 6 0 0 1 13 12.17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Guitar(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.17 3.59L16 1.76L14.24 0l-1.83 1.83v.85l-1.8 1.83a4.18 4.18 0 0 0-2.34-.85a3.73 3.73 0 0 0-2.8 1.07a4.61 4.61 0 0 0-1 1.34a.79.79 0 0 1-.23.34a1.77 1.77 0 0 1-.39 0a2.8 2.8 0 0 0-2.46.84A4.73 4.73 0 0 0 0 10.8A5.46 5.46 0 0 0 5.38 16a4.76 4.76 0 0 0 3.39-1.38a2.8 2.8 0 0 0 .83-2.46a1.76 1.76 0 0 1 0-.39a.85.85 0 0 1 .34-.23a4.42 4.42 0 0 0 1.34-1a3.71 3.71 0 0 0 1.07-2.79a4.11 4.11 0 0 0-.85-2.35l1.83-1.83zm-3.05 4.19a2.46 2.46 0 0 1-.72 1.88a3.23 3.23 0 0 1-1 .78a1.94 1.94 0 0 0-.69.47a1.57 1.57 0 0 0-.34 1.34a1.6 1.6 0 0 1-.48 1.51a3.86 3.86 0 0 1-5.44-.22a4.17 4.17 0 0 1-1.24-2.79a3.55 3.55 0 0 1 1-2.65a1.61 1.61 0 0 1 1.51-.48a1.56 1.56 0 0 0 1.34-.34a2.15 2.15 0 0 0 .48-.69a3 3 0 0 1 .77-1a2.47 2.47 0 0 1 1.88-.72a3 3 0 0 1 1.52.5L7.4 7.73l.88.88l2.34-2.35a3 3 0 0 1 .5 1.52"/><path fill="currentColor" d="M5.83 8.33a1.76 1.76 0 0 0-1.29.51a1.82 1.82 0 0 0 0 2.57a1.88 1.88 0 0 0 1.38.59a1.75 1.75 0 0 0 1.24-.51a1.82 1.82 0 0 0 0-2.57a1.93 1.93 0 0 0-1.33-.59m.47 2.27a.64.64 0 0 1-.9-.9a.56.56 0 0 1 .38-.15a.74.74 0 0 1 .45.2a.6.6 0 0 1 .07.85"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heart(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.37 2.56a3.92 3.92 0 0 0-3-1a4.1 4.1 0 0 0-1.82.52A9.18 9.18 0 0 0 8 3.06a9.35 9.35 0 0 0-1.49-1a3.85 3.85 0 0 0-1.77-.52A4.07 4.07 0 0 0 1.63 2.6A4 4 0 0 0 .35 5.52a3.83 3.83 0 0 0 .88 2.33a33.87 33.87 0 0 0 5.7 6.2l.39-.49l-.38.49a1.67 1.67 0 0 0 1.06.42a1.71 1.71 0 0 0 1.08-.42a37.42 37.42 0 0 0 6.06-6.73a3.5 3.5 0 0 0 .47-1.74a4 4 0 0 0-1.24-3.02m-.26 4.06a37.1 37.1 0 0 1-5.81 6.46a.56.56 0 0 1-.28.13a.51.51 0 0 1-.29-.14a32.77 32.77 0 0 1-5.49-5.94a2.74 2.74 0 0 1-.64-1.61a2.75 2.75 0 0 1 .88-2a2.79 2.79 0 0 1 2.16-.72h.1a2.73 2.73 0 0 1 1.19.38A10.23 10.23 0 0 1 7.24 4l.76.63l.76-.63a9 9 0 0 1 1.34-.86a2.91 2.91 0 0 1 1.26-.39h.1a2.68 2.68 0 0 1 2.07.68a2.78 2.78 0 0 1 .87 2a2.18 2.18 0 0 1-.29 1.19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func History(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.06.56A8.05 8.05 0 0 0 1.24 4.2V1.55H0V5a1.16 1.16 0 0 0 1.15 1.14h3.44V4.9H2.27a6.79 6.79 0 0 1 5.79-3.1a6.48 6.48 0 0 1 6.7 6.2a6.48 6.48 0 0 1-6.7 6.2A6.48 6.48 0 0 1 1.36 8H.12a7.71 7.71 0 0 0 7.94 7.44A7.71 7.71 0 0 0 16 8A7.71 7.71 0 0 0 8.06.56"/><path fill="currentColor" d="M7.44 4.28v4.34h3.6V7.38H8.68v-3.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Home(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.5 14.69h-1.25V7.78a.62.62 0 0 0-.25-.47L8.4 2.7a.65.65 0 0 0-.81 0L2 7.31a.62.62 0 0 0-.22.47v6.91H.5V7.78a1.87 1.87 0 0 1 .68-1.44l5.62-4.6a1.88 1.88 0 0 1 2.39 0l5.63 4.6a1.87 1.87 0 0 1 .68 1.44z"/><path fill="currentColor" d="M11.05 12.11H9.8A1.72 1.72 0 0 0 8 10.49a1.72 1.72 0 0 0-1.8 1.62H5a3 3 0 0 1 3-2.87a3 3 0 0 1 3.05 2.87m-6.1 0H6.2v2.58H4.95zm4.85 0h1.25v2.58H9.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Infinite(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.07 4.25c-1.62 0-3 .94-4.07 2.79c-1.08-1.85-2.45-2.79-4.07-2.79A3.84 3.84 0 0 0 0 8a3.84 3.84 0 0 0 3.93 3.75c1.62 0 3-.94 4.07-2.79c1.08 1.85 2.45 2.79 4.07 2.79A3.84 3.84 0 0 0 16 8a3.84 3.84 0 0 0-3.93-3.75M3.93 10.5A2.6 2.6 0 0 1 1.25 8a2.6 2.6 0 0 1 2.68-2.5c1.25 0 2.29.82 3.18 2.5c-.89 1.68-1.93 2.5-3.18 2.5m8.14 0c-1.25 0-2.29-.82-3.18-2.5c.89-1.68 1.93-2.5 3.18-2.5A2.6 2.6 0 0 1 14.75 8a2.6 2.6 0 0 1-2.68 2.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoCircle(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.37 6.7h1.25v5H7.37z"/><circle cx="8" cy="4.85" r=".65" fill="currentColor"/><path fill="currentColor" d="M8 .5A7.77 7.77 0 0 0 0 8a7.77 7.77 0 0 0 8 7.5A7.77 7.77 0 0 0 16 8A7.77 7.77 0 0 0 8 .5m0 13.75A6.52 6.52 0 0 1 1.25 8A6.52 6.52 0 0 1 8 1.75A6.52 6.52 0 0 1 14.75 8A6.52 6.52 0 0 1 8 14.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Invoice(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.44 7.47h5.26v1.25H5.44zm0 2.36h5.26v1.25H5.44zm0-4.76h5.26v1.25H5.44z"/><path fill="currentColor" d="M11.34 1L9.64.28L8.08 1L6.41.28L4.84 1L2.46 0v16l2.38-1l1.57.69L8.08 15l1.56.69l1.7-.69l2.2 1V0zm.94 13.11l-.92-.41l-1.69.69l-1.57-.72l-1.68.69l-1.55-.69l-1.15.47V1.86l1.15.47l1.55-.69l1.68.69l1.57-.69l1.69.69l.92-.41z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Italic(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 1.25V0H4.71v1.25h3.67l-2.02 13.5H2.5V16h8.79v-1.25H7.62l2.02-13.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LifeRing(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 .5A7.77 7.77 0 0 0 0 8a7.77 7.77 0 0 0 8 7.5A7.77 7.77 0 0 0 16 8A7.77 7.77 0 0 0 8 .5m5.38 3.73l-1.59 1.6A4.42 4.42 0 0 0 10 4.3l1.6-1.58a6.5 6.5 0 0 1 1.78 1.51M8 10.9A3.07 3.07 0 0 1 4.8 8A3.07 3.07 0 0 1 8 5.1A3.07 3.07 0 0 1 11.2 8A3.07 3.07 0 0 1 8 10.9m2.39-8.74L8.65 3.89H7.36L5.64 2.14A7.35 7.35 0 0 1 8 1.75a7 7 0 0 1 2.39.41m-5.95.54L6 4.29a4.45 4.45 0 0 0-1.79 1.53l-1.6-1.58A6.5 6.5 0 0 1 4.44 2.7M1.91 5.3L3.67 7a3.83 3.83 0 0 0-.12 1a4.32 4.32 0 0 0 .1.88l-1.77 1.76a5.85 5.85 0 0 1 0-5.34zm.66 6.41l1.61-1.59A4.34 4.34 0 0 0 6 11.71L4.44 13.3a6.38 6.38 0 0 1-1.87-1.59m3.07 2.15l1.72-1.75H8a5.52 5.52 0 0 0 .71-.05l1.73 1.73a7.27 7.27 0 0 1-2.44.46a7.35 7.35 0 0 1-2.36-.39m6-.6l-1.58-1.58a4.35 4.35 0 0 0 1.77-1.57l1.6 1.6a6.68 6.68 0 0 1-1.79 1.55m2.48-2.63l-1.77-1.76a3.57 3.57 0 0 0 .1-.87a3.75 3.75 0 0 0-.12-1l1.76-1.7a5.83 5.83 0 0 1 0 5.33z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lightbulb(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a5.35 5.35 0 0 0-5.5 5.18c0 3.23 2.59 3.23 2.59 5.17v3.06A2.77 2.77 0 0 0 8 16a2.78 2.78 0 0 0 2.92-2.59v-3.06c0-1.94 2.59-1.94 2.59-5.17A5.35 5.35 0 0 0 8 0m0 14.75a1.54 1.54 0 0 1-1.65-1.34V12.2h3.31v1.21A1.55 1.55 0 0 1 8 14.75m3.1-7.16a3.75 3.75 0 0 0-1.43 2.76V11H6.34v-.6a3.75 3.75 0 0 0-1.43-2.81a2.87 2.87 0 0 1-1.16-2.41A4.1 4.1 0 0 1 8 1.25a4.1 4.1 0 0 1 4.25 3.93a2.87 2.87 0 0 1-1.16 2.41z"/><path fill="currentColor" d="M8 2.66a3 3 0 0 0-2.13.85a2.33 2.33 0 0 0-.73 1.69a.62.62 0 0 0 .62.63a.63.63 0 0 0 .63-.63a1.1 1.1 0 0 1 .36-.8A1.74 1.74 0 0 1 8 3.91a.63.63 0 1 0 0-1.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Link(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 6.1a.31.31 0 0 0-.45.32a2.47 2.47 0 0 0 .51 1.22l.15.13A3 3 0 0 1 9.08 10a3.63 3.63 0 0 1-3.55 3.44a3 3 0 0 1-2.11-.85a3 3 0 0 1-.85-2.22A3.55 3.55 0 0 1 3.63 8a3.66 3.66 0 0 1 1.5-.91A5.19 5.19 0 0 1 5 6v-.16a4.84 4.84 0 0 0-2.31 1.3a4.5 4.5 0 0 0-.2 6.37a4.16 4.16 0 0 0 3 1.22a4.79 4.79 0 0 0 3.38-1.42a4.52 4.52 0 0 0 .21-6.38A4.16 4.16 0 0 0 8 6.1"/><path fill="currentColor" d="M13.46 2.54a4.16 4.16 0 0 0-3-1.22a4.79 4.79 0 0 0-3.37 1.42a4.52 4.52 0 0 0-.21 6.38A4.21 4.21 0 0 0 8 9.9a.31.31 0 0 0 .45-.31a2.41 2.41 0 0 0-.52-1.23l-.15-.13A3 3 0 0 1 6.92 6a3.63 3.63 0 0 1 3.55-3.44a3 3 0 0 1 2.11.85a3 3 0 0 1 .85 2.22A3.55 3.55 0 0 1 12.37 8a3.66 3.66 0 0 1-1.5.91a5.19 5.19 0 0 1 .13 1.14v.16a4.84 4.84 0 0 0 2.31-1.3a4.5 4.5 0 0 0 .15-6.37"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkOff(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 3.63a3.57 3.57 0 0 1 2.49-1.06a2.9 2.9 0 0 1 3 3A3.65 3.65 0 0 1 10.73 9l1 1a4.71 4.71 0 0 0 1.54-1a4.79 4.79 0 0 0 1.42-3.38a4.16 4.16 0 0 0-1.22-3a4.16 4.16 0 0 0-3-1.22a4.79 4.79 0 0 0-3.38 1.34a4.71 4.71 0 0 0-1 1.54l1 1A3.6 3.6 0 0 1 8 3.63m-6.31-.17l2.59 2.59a4.71 4.71 0 0 0-1.54 1a4.79 4.79 0 0 0-1.42 3.38a4.16 4.16 0 0 0 1.22 3a4.16 4.16 0 0 0 3 1.22a4.79 4.79 0 0 0 3.38-1.42a4.71 4.71 0 0 0 1-1.54l2.74 2.74l.89-.88l-11-11zM8 12.37a3.57 3.57 0 0 1-2.49 1.06a2.9 2.9 0 0 1-3-3A3.65 3.65 0 0 1 5.26 7L9 10.74a3.57 3.57 0 0 1-1 1.63"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func List(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.8 3.05H16v1.26H3.8zm0 4.29H16V8.6H3.8zm0 4.36H16v1.26H3.8zM1.25 2.5A1.22 1.22 0 0 1 2.5 3.68a1.22 1.22 0 0 1-1.25 1.17A1.22 1.22 0 0 1 0 3.68A1.22 1.22 0 0 1 1.25 2.5m0 4.3a1.18 1.18 0 1 1 0 2.35a1.18 1.18 0 1 1 0-2.35m0 4.35a1.18 1.18 0 1 1 0 2.35a1.18 1.18 0 1 1 0-2.35"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Location(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 .52a5.39 5.39 0 0 0-5.59 5.15c0 5 3.88 8.61 5.17 9.66a.66.66 0 0 0 .84 0c1.29-1.05 5.17-4.63 5.17-9.66A5.39 5.39 0 0 0 8 .52M8 14a15.2 15.2 0 0 1-2.46-2.76a9.85 9.85 0 0 1-1.88-5.57A4.14 4.14 0 0 1 8 1.77a4.14 4.14 0 0 1 4.34 3.9c0 4.08-2.96 7.16-4.34 8.33"/><path fill="currentColor" d="M8 2.54a2.73 2.73 0 0 0-2.84 2.65A2.74 2.74 0 0 0 8 7.84a2.75 2.75 0 0 0 2.83-2.65A2.74 2.74 0 0 0 8 2.54m0 4.05a1.49 1.49 0 0 1-1.57-1.4A1.49 1.49 0 0 1 8 3.79a1.5 1.5 0 0 1 1.58 1.4A1.5 1.5 0 0 1 8 6.59"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lock(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.25 11.2v1.42h1.5V11.2a1.48 1.48 0 0 0 .85-1.32A1.55 1.55 0 0 0 8 8.38a1.55 1.55 0 0 0-1.6 1.5a1.48 1.48 0 0 0 .85 1.32"/><path fill="currentColor" d="M14.75 5H13.5A5.29 5.29 0 0 0 8 0a5.29 5.29 0 0 0-5.5 5H1.25A1.25 1.25 0 0 0 0 6.25v8.5A1.25 1.25 0 0 0 1.25 16h13.5A1.25 1.25 0 0 0 16 14.75v-8.5A1.25 1.25 0 0 0 14.75 5M8 1.25A4.05 4.05 0 0 1 12.25 5h-8.5A4.05 4.05 0 0 1 8 1.25m6.75 13.5H1.25v-8.5h13.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockOpen(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.25 11.2v1.42h1.5V11.2a1.48 1.48 0 0 0 .85-1.32A1.55 1.55 0 0 0 8 8.38a1.55 1.55 0 0 0-1.6 1.5a1.48 1.48 0 0 0 .85 1.32"/><path fill="currentColor" d="M14.75 5H13.5A5.29 5.29 0 0 0 8 0a5.45 5.45 0 0 0-5.32 3.75H4a4.28 4.28 0 0 1 4-2.5A4.05 4.05 0 0 1 12.25 5h-11A1.25 1.25 0 0 0 0 6.25v8.5A1.25 1.25 0 0 0 1.25 16h13.5A1.25 1.25 0 0 0 16 14.75v-8.5A1.25 1.25 0 0 0 14.75 5m0 9.75H1.25v-8.5h13.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogOut(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.3 3.75h5.88V2.5H1.3A1.25 1.25 0 0 0 .05 3.75v8.5A1.25 1.25 0 0 0 1.3 13.5h5.88v-1.25H1.3z"/><path fill="currentColor" d="m15.4 7l-4-2.74l-.71 1l3.08 2.1H4.71v1.26h9.07l-3.08 2.11l.71 1L15.4 9a1.24 1.24 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mail(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 2.5H1a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1v-9a1 1 0 0 0-1-1m-1.3 1.25L8.42 8.56a.62.62 0 0 1-.84 0L2.3 3.75zm-12.45 8.5V4.48l5.49 5a1.86 1.86 0 0 0 2.52 0l5.49-5v7.77z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Marketing(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.25 2.1a1.25 1.25 0 0 0-1.17-.1L6.91 4.43a1.22 1.22 0 0 1-.46.09H2.5a1.25 1.25 0 0 0-1.25 1.25v.1H0v3h1.25V9a1.25 1.25 0 0 0 1.25 1.22L4 13.4a1.26 1.26 0 0 0 1.13.72h.63A1.25 1.25 0 0 0 7 12.87v-2.53l6.08 2.43a1.27 1.27 0 0 0 .47.09a1.29 1.29 0 0 0 .7-.22a1.25 1.25 0 0 0 .55-1V3.13a1.25 1.25 0 0 0-.55-1.03m-8.5 3.67V9H2.5V5.77zm0 7.1h-.63l-1.23-2.65h1.86zm1.62-3.72A2.29 2.29 0 0 0 7 9V5.7a2.26 2.26 0 0 0 .37-.11l6.18-2.46v8.48zm7.46-3.03v2.5a1.25 1.25 0 0 0 0-2.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mate(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.38 5.71H9.8l1-3.12a.67.67 0 0 1 .3-.36l1.66-.87l-.54-1.11l-1.65.87a1.89 1.89 0 0 0-.9 1.06L8.48 5.7H3.62a5.67 5.67 0 0 0-1.87 4.19A6.06 6.06 0 0 0 8 15.75a6.06 6.06 0 0 0 6.25-5.86a5.67 5.67 0 0 0-1.87-4.18M3.82 7h8.36a4.93 4.93 0 0 1 .73 1.23H3.09A4.93 4.93 0 0 1 3.82 7m9.43 2.94a4.53 4.53 0 0 1-.48 2l-1.44-1.25l-1.66 1.43L8 10.67l-1.67 1.45l-1.66-1.45l-1.44 1.25a4.53 4.53 0 0 1-.48-2v-.46h10.48c.01.12.02.28.02.43zM8 14.75A5.37 5.37 0 0 1 3.94 13l.73-.64l1.66 1.45L8 12.33l1.67 1.45l1.66-1.45l.73.64A5.37 5.37 0 0 1 8 14.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Menu(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.5 7.42h15v1.25H.5zm0 3.6h15v1.25H.5zm0-7.29h15v1.25H.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mobile(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.25 0h-8.5A1.25 1.25 0 0 0 2.5 1.25v13.5A1.25 1.25 0 0 0 3.75 16h8.5a1.25 1.25 0 0 0 1.25-1.25V1.25A1.25 1.25 0 0 0 12.25 0m0 14.75h-8.5V1.25h8.5z"/><ellipse cx="8" cy="12.75" fill="currentColor" rx=".8" ry=".75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Money(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.32 8a3 3 0 0 0-2-.7H5.63A1.59 1.59 0 0 1 4 5.69a2 2 0 0 1 0-.25a1.59 1.59 0 0 1 1.63-1.33h4.62a1.59 1.59 0 0 1 1.57 1.33h1.5a3.08 3.08 0 0 0-3.07-2.83H8.67V.31H7.42v2.3H5.63a3.08 3.08 0 0 0-3.07 2.83a2.09 2.09 0 0 0 0 .25a3.07 3.07 0 0 0 3.07 3.07h4.74A1.59 1.59 0 0 1 12 10.35a1.86 1.86 0 0 1 0 .34a1.59 1.59 0 0 1-1.55 1.24h-4.7a1.59 1.59 0 0 1-1.55-1.24H2.69a3.08 3.08 0 0 0 3.06 2.73h1.67v2.27h1.25v-2.27h1.7a3.08 3.08 0 0 0 3.06-2.73v-.34A3.06 3.06 0 0 0 12.32 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Moon(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.47 2.18a6.49 6.49 0 0 1 3.22 5.5a6.84 6.84 0 0 1-7.08 6.55a7.42 7.42 0 0 1-3.82-1a7 7 0 0 1-1.9-1.65a9.47 9.47 0 0 0 3.34-.74a9.92 9.92 0 0 0 3.3-2.24a10 10 0 0 0 2.94-6.42M10.82.45a.66.66 0 0 0-.66.69a8.63 8.63 0 0 1-2.55 6.54a8.68 8.68 0 0 1-6.09 2.57H.66a.66.66 0 0 0-.58 1a8.45 8.45 0 0 0 7.53 4.39A8.15 8.15 0 0 0 16 7.68A7.85 7.85 0 0 0 11.07.51a.61.61 0 0 0-.25-.06"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Notification(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.58 11.6l-1.33-2.18V6.33A4.36 4.36 0 0 0 10 2.26a2.45 2.45 0 0 0 0-.38A1.94 1.94 0 0 0 8 0a1.94 1.94 0 0 0-2 1.88a1.64 1.64 0 0 0 0 .38a4.36 4.36 0 0 0-2.25 4.07v3.09L2.42 11.6a1.25 1.25 0 0 0 1.06 1.9h1.77A2.68 2.68 0 0 0 8 16a2.68 2.68 0 0 0 2.75-2.5h1.77a1.25 1.25 0 0 0 1.06-1.9M7.25 1.88A.7.7 0 0 1 8 1.25a.7.7 0 0 1 .75.63a6 6 0 0 0-.75 0a5.9 5.9 0 0 0-.75 0M8 14.75a1.44 1.44 0 0 1-1.5-1.25h3A1.44 1.44 0 0 1 8 14.75m-4.52-2.5l1.34-2.17l.18-.31V6.33a4 4 0 0 1 .6-2.12A2.68 2.68 0 0 1 8 3.12a2.68 2.68 0 0 1 2.4 1.09a4 4 0 0 1 .6 2.12v3.44l.18.31l1.34 2.17z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Obelisk(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.25 12.25h-.37L11 2.87a1.21 1.21 0 0 0-.38-.78L8.87.35a1.25 1.25 0 0 0-1.74 0L5.34 2.09a1.21 1.21 0 0 0-.34.78l-.84 9.38h-.41A1.26 1.26 0 0 0 2.5 13.5V16h11v-2.5a1.26 1.26 0 0 0-1.25-1.25M6.2 3L8 1.25L9.8 3l.83 9.27H5.37zm6.05 11.77h-8.5V13.5h8.5z"/><circle cx="8.02" cy="4.04" r=".62" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OrderedList(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.59 3.03h12.2v1.26H3.59zm0 4.29h12.2v1.26H3.59zm0 4.35h12.2v1.26H3.59zM.99 4.79h.49V2.52H.6v.45h.39zm.87 3.88H.91l.14-.11l.3-.24c.35-.28.49-.5.49-.79A.74.74 0 0 0 1 6.8a.77.77 0 0 0-.81.84h.52A.34.34 0 0 1 1 7.25a.31.31 0 0 1 .31.31a.6.6 0 0 1-.22.44l-.87.75v.39h1.64zm-.36 3.56a.52.52 0 0 0 .28-.48a.67.67 0 0 0-.78-.62a.71.71 0 0 0-.77.75h.5a.3.3 0 0 1 .27-.32a.26.26 0 1 1 0 .51H.91v.38H1c.23 0 .37.11.37.29a.29.29 0 0 1-.33.29a.35.35 0 0 1-.36-.35H.21a.76.76 0 0 0 .83.8a.74.74 0 0 0 .83-.72a.53.53 0 0 0-.37-.53"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pencil(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.8 2.2a2.51 2.51 0 0 0-3.54 0l-6.9 6.91l-1.76 3.62a1.26 1.26 0 0 0 1.12 1.8a1.23 1.23 0 0 0 .55-.13l3.62-1.76l6-6l.83-.82l.06-.06a2.52 2.52 0 0 0 .02-3.56m-.89.89a1.25 1.25 0 0 1 0 1.77l-1.77-1.77a1.24 1.24 0 0 1 .86-.37a1.22 1.22 0 0 1 .91.37M2.73 13.27L4.29 10L6 11.71zm4.16-2.4L5.13 9.11L10.26 4L12 5.74z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Peso(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 .52A7.77 7.77 0 0 0 0 8a7.77 7.77 0 0 0 8 7.48A7.77 7.77 0 0 0 16 8A7.77 7.77 0 0 0 8 .52m0 13.71A6.51 6.51 0 0 1 1.25 8A6.51 6.51 0 0 1 8 1.77A6.51 6.51 0 0 1 14.75 8A6.51 6.51 0 0 1 8 14.23"/><path fill="currentColor" d="M9 7.38H7A.34.34 0 0 1 6.69 7a.13.13 0 0 1 .01 0a.34.34 0 0 1 .3-.29h1.94a.34.34 0 0 1 .33.29h1.25a1.59 1.59 0 0 0-1.58-1.54h-.32V4.21H7.37v1.25H7A1.59 1.59 0 0 0 5.44 7a1.55 1.55 0 0 0 .35 1A1.59 1.59 0 0 0 7 8.62h2a.34.34 0 0 1 .33.38a.33.33 0 0 1-.33.29H7.08A.33.33 0 0 1 6.75 9H5.49a1.58 1.58 0 0 0 1.58 1.54h.31v1.25h1.24v-1.25H9A1.58 1.58 0 0 0 10.56 9a1.51 1.51 0 0 0-.34-1A1.59 1.59 0 0 0 9 7.38"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Picture(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 6.73L11 5L7.68 7.53L6 6L2.5 8.49V11h11zm-1.25 3h-8.5v-.59L5.9 7.6l.94.86l.77.7l.83-.63l2.59-2l1.22.84z"/><ellipse cx="3.3" cy="5.75" fill="currentColor" rx=".8" ry=".75"/><path fill="currentColor" d="M14.75 2.5H1.25A1.25 1.25 0 0 0 0 3.75v8.5a1.25 1.25 0 0 0 1.25 1.25h13.5A1.25 1.25 0 0 0 16 12.25v-8.5a1.25 1.25 0 0 0-1.25-1.25m0 9.75H1.25v-8.5h13.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pix(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.917 11.71a2.046 2.046 0 0 1-1.454-.602l-2.1-2.1a.4.4 0 0 0-.551 0l-2.108 2.108a2.044 2.044 0 0 1-1.454.602h-.414l2.66 2.66c.83.83 2.177.83 3.007 0l2.667-2.668zM4.25 4.282c.55 0 1.066.214 1.454.602l2.108 2.108a.39.39 0 0 0 .552 0l2.1-2.1a2.044 2.044 0 0 1 1.453-.602h.253L9.503 1.623a2.127 2.127 0 0 0-3.007 0l-2.66 2.66h.414z"/><path fill="currentColor" d="m14.377 6.496l-1.612-1.612a.307.307 0 0 1-.114.023h-.733c-.379 0-.75.154-1.017.422l-2.1 2.1a1.005 1.005 0 0 1-1.425 0L5.268 5.32a1.448 1.448 0 0 0-1.018-.422h-.9a.306.306 0 0 1-.109-.021L1.623 6.496c-.83.83-.83 2.177 0 3.008l1.618 1.618a.305.305 0 0 1 .108-.022h.901c.38 0 .75-.153 1.018-.421L7.375 8.57a1.034 1.034 0 0 1 1.426 0l2.1 2.1c.267.268.638.421 1.017.421h.733c.04 0 .079.01.114.024l1.612-1.612c.83-.83.83-2.178 0-3.008z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Planet(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 .5A7.76 7.76 0 0 0 0 8a7.76 7.76 0 0 0 8 7.5A7.76 7.76 0 0 0 16 8A7.76 7.76 0 0 0 8 .5m6.71 6.8L13.48 7c-.25-.07-.27-.09-.29-.12c-.15-.2-.32-.47-.48-.73c0-.09-.13-.23-.16-.31s.35-.6.51-.84a2.43 2.43 0 0 1 .59-.45a5.87 5.87 0 0 1 1.06 2.75M8 1.75l-.09.17a.19.19 0 0 1 0-.1c0 .06-.15.15-.25.25l-.3.29a.85.85 0 0 0-.08 1.08h-.12a1.05 1.05 0 0 0-.81.42a1.27 1.27 0 0 0-.2 1.07V5a3 3 0 0 0-.43.11l-.24.08l-.64.21a1.2 1.2 0 0 0-.81.8a1 1 0 0 0 .2.93a5.67 5.67 0 0 0 1.38 1.09a4.17 4.17 0 0 0 1.67.65h1.68a1.2 1.2 0 0 1 1.04.51a.49.49 0 0 1 .13.43a.77.77 0 0 1-.15.35a2.71 2.71 0 0 0-.95 1.61a11.11 11.11 0 0 1-.48 1.38c-.12.31-.23.61-.31.85a3.32 3.32 0 0 1-1-.08a3.28 3.28 0 0 0-.5-2.12a2.24 2.24 0 0 1-.53-1.42a2.11 2.11 0 0 0-1.47-2.29a10.81 10.81 0 0 1-2.9-2.64A6.79 6.79 0 0 1 8 1.75M1.25 8a5.64 5.64 0 0 1 .12-1.16a10.29 10.29 0 0 0 2.94 2.42c.6.22.69.45.69 1.12a3.45 3.45 0 0 0 .86 2.27A3.05 3.05 0 0 1 6 14a6.35 6.35 0 0 1-4.75-6m8.32 6.08c0-.15.12-.32.18-.48a10.2 10.2 0 0 0 .55-1.6a1.55 1.55 0 0 1 .54-.86a1.91 1.91 0 0 0 .57-1.3a1.71 1.71 0 0 0-.47-1.27a2.45 2.45 0 0 0-2-.9H7.35a4.77 4.77 0 0 1-2-1.11l.47-.16l.27-.08a.79.79 0 0 1 .38-.07l.09.15a.64.64 0 0 0 .81.29a.65.65 0 0 0 .34-.8v-.18c-.11-.3-.24-.72-.32-1A1.42 1.42 0 0 0 8.68 4a1 1 0 0 0-.18-1a3.44 3.44 0 0 0 .33-.34a1 1 0 0 0 .22-.8a6.93 6.93 0 0 1 3.73 1.8a3 3 0 0 0-.79.7a9.14 9.14 0 0 0-.64 1.09a1.46 1.46 0 0 0 .24 1.39c.18.31.38.61.56.86a1.58 1.58 0 0 0 1 .58c.22.06 1 .22 1.55.33a6.44 6.44 0 0 1-5.13 5.47"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusCircle(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.64 4.33H7.39v3.05H4.34v1.25h3.05v3.05h1.25V8.63h3.05V7.38H8.64z"/><path fill="currentColor" d="M8 .5A7.77 7.77 0 0 0 0 8a7.77 7.77 0 0 0 8 7.5A7.77 7.77 0 0 0 16 8A7.77 7.77 0 0 0 8 .5m0 13.75A6.52 6.52 0 0 1 1.25 8A6.52 6.52 0 0 1 8 1.75A6.52 6.52 0 0 1 14.75 8A6.52 6.52 0 0 1 8 14.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Printer(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.25 6.57h-.75v-.05a1.26 1.26 0 0 0-.37-.89L9.37 1.87a1.26 1.26 0 0 0-.89-.37H3.75A1.25 1.25 0 0 0 2.5 2.75v3.82h-.75A1.26 1.26 0 0 0 .5 7.85v3.84A1.26 1.26 0 0 0 1.75 13h.75v.28a1.25 1.25 0 0 0 1.25 1.22h8.5a1.25 1.25 0 0 0 1.25-1.25V13h.75a1.26 1.26 0 0 0 1.25-1.28V7.85a1.26 1.26 0 0 0-1.25-1.28M3.75 2.75h4.73l3.77 3.77v.69h-8.5zm8.5 10.5h-8.5v-.92h8.5zm2-1.56h-.75v-.64h-11v.64h-.74V7.85h.74v.64h11v-.67h.74z"/><ellipse cx="12.88" cy="9.77" fill="currentColor" rx=".63" ry=".59"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pyramid(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.77 11V8.5a1.25 1.25 0 0 0-1.25-1.25v-2.5a1.25 1.25 0 0 0-1.27-1.25h-1.87V1.25A1.26 1.26 0 0 0 9.12 0H6.88a1.26 1.26 0 0 0-1.26 1.25V3.5H3.75A1.25 1.25 0 0 0 2.5 4.75v2.5A1.25 1.25 0 0 0 1.23 8.5V11A1.25 1.25 0 0 0 0 12.25V16h16v-3.75A1.25 1.25 0 0 0 14.77 11m-2.52-6.25v2.5h-1.14l-.68-2.5zm-5.37-3.5h2.24V3.5H6.88zm-3.13 3.5h1.83l-.69 2.5H3.75zm-.91 10H1.25v-2.5h2.28zM2.48 11V8.5h2.07L3.87 11zm1.66 3.75l2.07-7.58l.66-2.42h2.26l.69 2.5l.34 1.25l.68 2.5l.34 1.25l.68 2.5zm7.31-6.25h2.07V11h-1.38zm3.3 6.25h-1.59l-.68-2.5h2.27z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QrCode(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.25 1.5h-2.12c-.35 0-.62.28-.62.62s.28.62.62.62h2.12v2.52c0 .35.28.62.62.62s.62-.28.62-.62V2.75c0-.69-.56-1.25-1.25-1.25zM1.12 5.89c.35 0 .62-.28.62-.62V2.75h2.15c.35 0 .62-.28.62-.62s-.28-.62-.62-.62H1.75C1.06 1.51.5 2.07.5 2.76v2.52c0 .35.28.62.62.62zm2.77 7.36H1.75v-2.52c0-.35-.28-.62-.62-.62s-.62.28-.62.62v2.52c0 .69.56 1.25 1.25 1.25h2.13c.35 0 .62-.28.62-.62s-.28-.62-.62-.62zm10.99-3.14c-.35 0-.62.28-.62.62v2.52h-2.12c-.35 0-.62.28-.62.62s.28.62.62.62h2.12c.69 0 1.25-.56 1.25-1.25v-2.52c0-.35-.28-.62-.62-.62z"/><path fill="currentColor" d="M6.12 3.55H4.03c-.69 0-1.25.56-1.25 1.25v1.54c0 .69.56 1.25 1.25 1.25h2.09c.69 0 1.25-.56 1.25-1.25V4.8c0-.69-.56-1.25-1.25-1.25m0 2.79H4.03V4.8h2.09zm7.1 4.86V9.66c0-.69-.56-1.25-1.25-1.25H9.88c-.69 0-1.25.56-1.25 1.25v1.54c0 .69.56 1.25 1.25 1.25h2.09c.69 0 1.25-.56 1.25-1.25M9.88 9.66h2.09v1.54H9.88zM8.63 5.85h2.52V7.1H8.63zm0-1.8h4.59V5.3H8.63zM4.85 8.9h2.52v1.25H4.85zm-2.07 1.8h4.59v1.25H2.78z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionCircle(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 .5A7.77 7.77 0 0 0 0 8a7.77 7.77 0 0 0 8 7.5A7.77 7.77 0 0 0 16 8A7.77 7.77 0 0 0 8 .5m0 13.75A6.52 6.52 0 0 1 1.25 8A6.52 6.52 0 0 1 8 1.75A6.52 6.52 0 0 1 14.75 8A6.52 6.52 0 0 1 8 14.25"/><circle cx="7.98" cy="10.95" r=".76" fill="currentColor"/><path fill="currentColor" d="M9.73 4.75A2.72 2.72 0 0 0 8 4.19a2.28 2.28 0 0 0-2.41 2.17v.11h1.24v-.1A1.12 1.12 0 0 1 8 5.33a1 1 0 0 1 1.12 1c0 .35-.24.73-.78 1.11a2 2 0 0 0-1 1.46v.36h1.24V9a.76.76 0 0 1 .23-.51A3.92 3.92 0 0 1 9.33 8l.17-.14a2 2 0 0 0 .91-1.67a1.85 1.85 0 0 0-.68-1.44"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Real(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.41 7.2a1.7 1.7 0 0 0-1.7-1.7H2.62v5h1.25V8.9h.84a1.59 1.59 0 0 1 1.55 1.6h1.25a2.82 2.82 0 0 0-.77-1.95a1.71 1.71 0 0 0 .67-1.35m-1.7-.45a.45.45 0 0 1 0 .9H3.87v-.9z"/><path fill="currentColor" d="M8 .5A7.77 7.77 0 0 0 0 8a7.77 7.77 0 0 0 8 7.5A7.77 7.77 0 0 0 16 8A7.77 7.77 0 0 0 8 .5m0 13.75A6.52 6.52 0 0 1 1.25 8A6.52 6.52 0 0 1 8 1.75A6.52 6.52 0 0 1 14.75 8A6.52 6.52 0 0 1 8 14.25"/><path fill="currentColor" d="M13.22 8A1.59 1.59 0 0 0 12 7.38h-2A.34.34 0 0 1 9.69 7a.13.13 0 0 1 .01 0a.34.34 0 0 1 .3-.3h1.91a.34.34 0 0 1 .33.28h1.25a1.59 1.59 0 0 0-1.58-1.53h-.32V4.2h-1.22v1.25H10A1.6 1.6 0 0 0 8.44 7a1.55 1.55 0 0 0 .35 1a1.59 1.59 0 0 0 1.21.62h2a.34.34 0 0 1 .34.34V9a.33.33 0 0 1-.33.29h-1.93A.33.33 0 0 1 9.75 9H8.49a1.58 1.58 0 0 0 1.58 1.54h.3v1.26h1.25v-1.26H12A1.58 1.58 0 0 0 13.56 9a1.51 1.51 0 0 0-.34-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Redo(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.94.56a8.05 8.05 0 0 1 6.82 3.64V1.55H16V5a1.16 1.16 0 0 1-1.15 1.15h-3.44V4.9h2.32a6.79 6.79 0 0 0-5.79-3.1A6.48 6.48 0 0 0 1.24 8a6.48 6.48 0 0 0 6.7 6.2a6.48 6.48 0 0 0 6.7-6.2h1.24a7.71 7.71 0 0 1-7.94 7.44A7.71 7.71 0 0 1 0 8A7.71 7.71 0 0 1 7.94.56"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RemoveFormat(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.03 2.49l.19-1.25h3.06L7.7 5.16l1.09 1.09l.75-5.01h3.07l-.18 1.15l1.23.19l.39-2.58h-9.9l-.21 1.4zm3.53 5.29l-1.09-1.1l-4.64-4.63l-.88.87l5.29 5.29l-.97 6.45H4.05v1.25h2.66l.62.09l.01-.09h2.23v-1.25H7.53l.8-5.36l4.56 4.56l.88-.88z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rocket(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.67 10.73a3.52 3.52 0 0 0-.94 1.93a5 5 0 0 0-.07 1.1v.58h.8a5.05 5.05 0 0 0 .88-.08a3.46 3.46 0 0 0 1.93-.94a1.76 1.76 0 0 0-.14-2.48a1.76 1.76 0 0 0-2.46-.11m1.74 1.73a2.26 2.26 0 0 1-1.26.6h-.22v-.22a2.26 2.26 0 0 1 .6-1.26a.36.36 0 0 1 .24-.08a.67.67 0 0 1 .47.22a.54.54 0 0 1 .17.74M14.65 2.24a.91.91 0 0 0-.89-.89A8.75 8.75 0 0 0 7.27 3.5L5.64 5.4l-2.4-.5a1 1 0 0 0-.92.27l-.68.68a1 1 0 0 0-.28.81a1 1 0 0 0 .45.74l2.06 1.32l.13.08l3.2 3.25l.08.08l1.32 2.06a1 1 0 0 0 .74.45h.11a1 1 0 0 0 .7-.29l.68-.68a1 1 0 0 0 .27-.92l-.5-2.39l1.84-1.58a8.79 8.79 0 0 0 2.21-6.54M3.11 6.15l1.32.28l-.64.75l-1-.67zm6.38 7.1l-.67-1l.75-.64l.28 1.32zm2.39-5.11l.18.17zm-.28-.28L7.92 11L5 8.08L8.14 4.4a7.44 7.44 0 0 1 5.26-1.8a7.48 7.48 0 0 1-1.8 5.26"/><path fill="currentColor" d="M11.13 6.63a1.19 1.19 0 0 0-.06-1.7a1.16 1.16 0 1 0-1.64 1.63a1.2 1.2 0 0 0 1.7.07"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Scooter(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.37 7.53a3.67 3.67 0 0 0-2.5 1H7.16a.65.65 0 0 1-.38-.13L5 7V4h2V2.75H4.1a2.91 2.91 0 0 0-.33-.37a2.34 2.34 0 0 0-.77-.5a2.64 2.64 0 0 0-1-.18V4h1.75v3.16a.62.62 0 0 1-.09.32L2.19 9.93A2.38 2.38 0 0 0 0 12.2a2.41 2.41 0 0 0 2.5 2.3A2.41 2.41 0 0 0 5 12.2h11V11a3.54 3.54 0 0 0-3.63-3.47m0 1.25A2.29 2.29 0 0 1 14.75 11H10a2.29 2.29 0 0 1 2.37-2.22M2.5 13.25a1.16 1.16 0 0 1-1.25-1.05a1.16 1.16 0 0 1 1.25-1a1.16 1.16 0 0 1 1.25 1a1.16 1.16 0 0 1-1.25 1.05M4.59 11a2.38 2.38 0 0 0-1.06-.83L4.62 8.3L6 9.36a1.9 1.9 0 0 0 1.13.38H9A3.22 3.22 0 0 0 8.75 11z"/><path fill="currentColor" d="M14.75 1.5H11a1.25 1.25 0 0 0-1.3 1.25v3A1.25 1.25 0 0 0 11 7h3.8A1.25 1.25 0 0 0 16 5.75v-3a1.25 1.25 0 0 0-1.25-1.25m0 4.25H11v-3h3.8zm-2.38 7.5a1.17 1.17 0 0 1-1.25-1.05H9.87a2.41 2.41 0 0 0 2.5 2.3a2.41 2.41 0 0 0 2.5-2.3h-1.25a1.16 1.16 0 0 1-1.25 1.05"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Search(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14.91 13.09l-3.68-3.21a4.86 4.86 0 0 0 .86-2.77A5.34 5.34 0 0 0 6.59 2a5.35 5.35 0 0 0-5.5 5.15a5.34 5.34 0 0 0 5.5 5.15a5.71 5.71 0 0 0 3.82-1.44L14.08 14zM6.59 11a4.09 4.09 0 0 1-4.25-3.9a4.09 4.09 0 0 1 4.25-3.9a4.09 4.09 0 0 1 4.25 3.9A4.08 4.08 0 0 1 6.59 11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Share(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 4.67a2.14 2.14 0 0 0 2.19-2.09A2.14 2.14 0 0 0 8 .5a2.14 2.14 0 0 0-2.24 2.08a1.93 1.93 0 0 0 .22.9L3.3 6.16a2.19 2.19 0 0 0-1.11-.3A2.14 2.14 0 0 0 0 8a2.14 2.14 0 0 0 2.19 2a2.14 2.14 0 0 0 1-.23l2.77 2.78a2.06 2.06 0 0 0-.18.83A2.14 2.14 0 0 0 8 15.5a2.09 2.09 0 1 0 0-4.17a2.25 2.25 0 0 0-1.18.33L4.09 9a1.77 1.77 0 0 0 .15-.31h7.52A2.19 2.19 0 0 0 13.81 10A2.14 2.14 0 0 0 16 8a2.2 2.2 0 0 0-4.3-.52H4.3a1.74 1.74 0 0 0-.14-.38l2.67-2.72A2.33 2.33 0 0 0 8 4.67m0-1.24a.89.89 0 0 0 .94-.85a.89.89 0 0 0-.94-.84a.89.89 0 0 0-1 .84a.89.89 0 0 0 1 .85M2.19 8.8a.9.9 0 0 0 .94-.8a.9.9 0 0 0-.94-.84a.9.9 0 0 0-.95.84a.91.91 0 0 0 .95.8M14.76 8a1 1 0 0 1-1.89 0a1 1 0 0 1 1.89 0m-5.87 5.42a.89.89 0 0 1-.94.84a.85.85 0 1 1 0-1.69a.89.89 0 0 1 .94.85"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingCart(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.35 10.48H4.5l-.24-1.25h9.13a1.24 1.24 0 0 0 1.22-1l.84-4a1.25 1.25 0 0 0-1.22-1.51H3l-.22-1.24H.5v1.25h1.25l1.5 7.84a2 2 0 0 0-1.54 1.93a2.09 2.09 0 0 0 2.16 2a2.08 2.08 0 0 0 2.13-2a2 2 0 0 0-.16-.77h5.49a2 2 0 0 0-.16.77a2.09 2.09 0 0 0 2.16 2a2 2 0 1 0 0-4zM14.23 4l-.84 4H4l-.74-4zM3.87 13.27A.85.85 0 0 1 3 12.5a.85.85 0 0 1 .91-.77a.84.84 0 0 1 .9.77a.84.84 0 0 1-.94.77m9.48 0a.85.85 0 0 1-.91-.77a.92.92 0 0 1 1.81 0a.85.85 0 0 1-.9.77"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shot(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m11.37 9.85l.87-6.4a1.22 1.22 0 0 0-.3-1A1.24 1.24 0 0 0 11 2H2.85a1.24 1.24 0 0 0-1.23 1.41l1.15 8.48A1.27 1.27 0 0 0 4 13h4.84a4.41 4.41 0 0 0 2.66 1h.14a3.9 3.9 0 0 0 2.76-1.12zM11 3.25l-.15 1.09H3l-.15-1.09zM7.23 9.72a4.28 4.28 0 0 0 .57 2H4l-.83-6.13h7.47l-.42 3.11l-1.87-1.87a3.83 3.83 0 0 0-1.12 2.89m4.31 3a3.18 3.18 0 0 1-2.12-.94a3.21 3.21 0 0 1-.95-2.13a2.71 2.71 0 0 1 .11-.86l1.91 1.92l1.91 1.91a2.73 2.73 0 0 1-.76.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SizeHeight(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9 15.46l2.74-4l-1-.71l-2.1 3.09V2.16l2.1 3.09l1-.71L9 .54a1.25 1.25 0 0 0-2 0l-2.74 4l1 .71l2.12-3.09v11.68l-2.11-3.09l-1 .71l2.74 4a1.25 1.25 0 0 0 1.99 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SizeWidth(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.46 7l-4-2.74l-.71 1l3.09 2.11H2.16l3.09-2.1l-.71-1L.54 7a1.25 1.25 0 0 0 0 2l4 2.74l.71-1l-3.09-2.12h11.68l-3.09 2.11l.71 1l4-2.74a1.25 1.25 0 0 0 0-1.99"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sliders(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 8.75h-.88V2.5H1.88v6.25H1a1 1 0 0 0-1 1v1.75a1 1 0 0 0 1 1h.88v1h1.24v-1H4a1 1 0 0 0 1-1V9.75a1 1 0 0 0-1-1m-.25 2.5h-2.5V10h2.5zm5.75-8h-.88V2.5H7.38v.75H6.5a1 1 0 0 0-1 1V6a1 1 0 0 0 1 1h.88v6.5h1.24V7h.88a1 1 0 0 0 1-1V4.25a1 1 0 0 0-1-1m-.25 2.5h-2.5V4.5h2.5zM15 7.35h-.88V2.5h-1.24v4.85H12a1 1 0 0 0-1 1v1.75a1 1 0 0 0 1 1h.88v2.4h1.24v-2.4H15a1 1 0 0 0 1-1V8.35a1 1 0 0 0-1-1m-.25 2.5h-2.5V8.6h2.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Star(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8 3.12l1.43 2.77l.31.59l.66.08l3.12.35L11.26 9l-.5.47l.12.67l.53 2.95l-2.85-1.42l-.56-.28l-.57.28l-2.85 1.44l.53-2.95l.12-.67l-.5-.49l-2.26-2.08l3.12-.36l.66-.08l.31-.59zm0-2a.63.63 0 0 0-.57.33l-2 3.84l-4.51.55a.59.59 0 0 0-.34 1l3.3 3.08l-.76 4.21a.61.61 0 0 0 .62.7a.65.65 0 0 0 .26-.05l4-2l4 2a.65.65 0 0 0 .29.07a.61.61 0 0 0 .62-.7l-.76-4.21l3.31-3.09a.59.59 0 0 0-.34-1l-4.54-.51l-2-3.84A.63.63 0 0 0 8 1.15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stats(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.75 13.25V1.5H.5v12a1.24 1.24 0 0 0 1.22 1H15.5v-1.25z"/><path fill="currentColor" d="M3.15 8H4.4v3.9H3.15zm3.26-4h1.26v7.9H6.41zm3.27 2h1.25v5.9H9.68zm3.27-3.5h1.25v9.4h-1.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StickyNote(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.25 1.5H1.75A1.25 1.25 0 0 0 .5 2.75v10.5a1.25 1.25 0 0 0 1.25 1.25h8.69a1.24 1.24 0 0 0 .93-.41l3.81-4.23A1.24 1.24 0 0 0 15.5 9V2.75a1.25 1.25 0 0 0-1.25-1.25M1.75 13.25V2.75h12.5v5h-3.81A1.25 1.25 0 0 0 9.19 9v4.23zm8.69 0V9h3.81z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stop(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 7.38h10v1.25H3z"/><path fill="currentColor" d="M8 .5A7.76 7.76 0 0 0 0 8a7.76 7.76 0 0 0 8 7.5A7.76 7.76 0 0 0 16 8A7.76 7.76 0 0 0 8 .5m0 13.75A6.52 6.52 0 0 1 1.25 8A6.52 6.52 0 0 1 8 1.75A6.52 6.52 0 0 1 14.75 8A6.52 6.52 0 0 1 8 14.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Store(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.49 12h4.33V7.66H2.49zm1.25-3.09h1.83v1.83H3.74zm7.33-1.25a2.43 2.43 0 0 0-2.43 2.43v3.4H9.9v-3.4a1.18 1.18 0 1 1 2.35 0v3.4h1.25v-3.4a2.43 2.43 0 0 0-2.43-2.43M2.49 5.07H13.5v1.3H2.49z"/><path fill="currentColor" d="M14.12 2.51H1.88A1.88 1.88 0 0 0 0 4.39v9.1h1.25v-9.1a.63.63 0 0 1 .63-.63h12.24a.63.63 0 0 1 .63.63v9.1H16v-9.1a1.88 1.88 0 0 0-1.88-1.88"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sun(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 3.33A4.85 4.85 0 0 0 3 8a4.85 4.85 0 0 0 5 4.67A4.85 4.85 0 0 0 13 8a4.85 4.85 0 0 0-5-4.67m0 8.09A3.6 3.6 0 0 1 4.25 8A3.6 3.6 0 0 1 8 4.58A3.6 3.6 0 0 1 11.75 8A3.6 3.6 0 0 1 8 11.42M7.33 0h1.25v2H7.33zm0 14h1.25v2H7.33zM14 7.33h2v1.25h-2zm-14 0h2v1.25H0zM11.92 2.4h2v1.25h-2zm-9.9 9.9h2v1.25h-2zm10.33-.38h1.25v2h-1.25zm-9.9-9.9H3.7v2H2.45z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tag(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 7.2L8.8 1a1.25 1.25 0 0 0-.88-.34h-6A1.25 1.25 0 0 0 .66 1.91v6A1.25 1.25 0 0 0 1 8.8L7.2 15a1.24 1.24 0 0 0 .88.36A1.28 1.28 0 0 0 9 15l6-6a1.27 1.27 0 0 0 0-1.8m-6.9 6.89L1.91 7.92v-6h6l6.17 6.17z"/><ellipse cx="4.95" cy="4.95" fill="currentColor" rx="1.41" ry="1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Telephone(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.91 2.75h3a11.67 11.67 0 0 1-3.8 7.46a11.71 11.71 0 0 1-7.88 3h-.77V10.4L4 9.52l.84 1.42l.59 1.06l1.07-.62a14.36 14.36 0 0 0 4.77-4.5L12 5.73l-1.22-.64l-1-.52l1.09-1.82m0-1.25a1.26 1.26 0 0 0-1.07.61L8.75 3.93a1.24 1.24 0 0 0 .49 1.75l1 .52a13.06 13.06 0 0 1-4.36 4.1L5 8.89a1.25 1.25 0 0 0-1-.62a1.22 1.22 0 0 0-.41.07L1 9.22a1.25 1.25 0 0 0-.79 1.18v2.85a1.25 1.25 0 0 0 1.25 1.25h.75A13 13 0 0 0 15.14 2.88a1.26 1.26 0 0 0-1.25-1.38z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextSize(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.1 2.75h3.09V4h1.25V1.5H.5V4h1.25V2.75h3.1v10.5H2.7v1.25h5.55v-1.25H6.1z"/><path fill="currentColor" d="M12.4 6.75H8.06v2.5h1.25V8h1.84v5.25H9.63v1.25h4.3v-1.25H12.4V8h1.85v1.25h1.25v-2.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tiendanube(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.25 2.24a5.79 5.79 0 0 0-4 1.63a4.48 4.48 0 1 0 0 8.26a5.76 5.76 0 1 0 4-9.89m0 10.24A4.49 4.49 0 0 1 5.76 8H4.48a5.74 5.74 0 0 0 .89 3.07a3.29 3.29 0 0 1-.88.13a3.2 3.2 0 0 1 0-6.4A3.2 3.2 0 0 1 7.69 8H9a4.42 4.42 0 0 0-1.63-3.43a4.48 4.48 0 1 1 2.88 7.91"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tools(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.25 10.71L11.57 8l2.26-2.26a2.49 2.49 0 0 0 0-3.53a2.5 2.5 0 0 0-3.53 0l-.89.88L8 4.5L5.28 1.75a1.26 1.26 0 0 0-1.76 0L1.75 3.52a1.25 1.25 0 0 0 0 1.77L4.5 8l-.22.22l-.89.88l-1.75 3.66a1.25 1.25 0 0 0 1.67 1.67l3.62-1.75l.49-.49l.39-.39l.19-.23l2.68 2.68a1.26 1.26 0 0 0 1.76 0l1.77-1.77a1.25 1.25 0 0 0 .04-1.77m-2.19-8a1.27 1.27 0 0 1 .89.36a1.25 1.25 0 0 1 0 1.77l-1.77-1.72a1.27 1.27 0 0 1 .88-.36zM2.63 4.4L4.4 2.64l.82.82l-.87.88l.88.88l.88-.88l1 1l-1.73 1.81zm.13 8.91l1.57-3.23L6 11.74zm4.17-2.4L5.16 9.14L10.3 4l1.76 1.76zm4.67 2.45l-2.68-2.67l1.77-1.77l.93.93l-.88.88l.88.89l.89-.89l.86.87z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TransferPeso(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.75 2.5H1.25A1.2 1.2 0 0 0 0 3.64v8.72a1.2 1.2 0 0 0 1.25 1.14h13.5A1.2 1.2 0 0 0 16 12.36V3.64a1.2 1.2 0 0 0-1.25-1.14m0 9.75H1.25v-8.5h13.5z"/><path fill="currentColor" d="M7 8.62h2a.34.34 0 0 1 .33.38a.33.33 0 0 1-.33.29H7.08A.33.33 0 0 1 6.75 9H5.49a1.58 1.58 0 0 0 1.58 1.54h.31v1.26h1.24v-1.26H9A1.58 1.58 0 0 0 10.56 9a1.51 1.51 0 0 0-.34-1A1.59 1.59 0 0 0 9 7.38H7A.34.34 0 0 1 6.69 7a.13.13 0 0 1 .01 0a.34.34 0 0 1 .3-.3h1.94a.34.34 0 0 1 .33.3h1.25a1.59 1.59 0 0 0-1.58-1.55h-.32V4.2H7.37v1.25H7A1.6 1.6 0 0 0 5.44 7a1.55 1.55 0 0 0 .35 1A1.59 1.59 0 0 0 7 8.62"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TransferReal(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 8.52h2a.36.36 0 0 1 .25.11a.35.35 0 0 1 .1.25a.19.19 0 0 1 0 .08a.32.32 0 0 1-.12.2a.35.35 0 0 1-.22.08H10a.35.35 0 0 1-.22-.08a.34.34 0 0 1-.12-.16l-.05-.23h-1.2v.33a1.6 1.6 0 0 0 .52 1a1.57 1.57 0 0 0 1 .41h.43v.94h1.13v-.94h.36a1.6 1.6 0 0 0 1.56-1.41a1.09 1.09 0 0 0 0-.18a1.56 1.56 0 0 0-.15-.67a1.64 1.64 0 0 0-.43-.55a1.56 1.56 0 0 0-1-.36H10a.37.37 0 0 1-.25-.1a.36.36 0 0 1-.11-.26a.09.09 0 0 1 0-.05a.3.3 0 0 1 .12-.22a.33.33 0 0 1 .24-.13h1.92a.35.35 0 0 1 .23.09a.4.4 0 0 1 .12.22v.25h1.2v-.33a1.63 1.63 0 0 0-.5-1a1.6 1.6 0 0 0-1.07-.42h-.31v-1h-1.15v1H10a1.57 1.57 0 0 0-1.07.42a1.58 1.58 0 0 0-.5 1v.12a1.58 1.58 0 0 0 .47 1.12a1.53 1.53 0 0 0 1.1.47m-6.27.29h.84a1.63 1.63 0 0 1 1.59 1.62h1.25a2.86 2.86 0 0 0-.79-2a1.69 1.69 0 0 0 .68-1.34a1.71 1.71 0 0 0-1.71-1.71H2.48v5.05h1.25zm1.86-2.18a.47.47 0 0 1 0 .93H3.73v-.93z"/><path fill="currentColor" d="M14.75 2.5H1.25A1.2 1.2 0 0 0 0 3.64v8.72a1.2 1.2 0 0 0 1.25 1.14h13.5A1.2 1.2 0 0 0 16 12.36V3.64a1.2 1.2 0 0 0-1.25-1.14m0 9.75H1.25v-8.5h13.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trash(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.11 1.88a1.27 1.27 0 0 0-.9-.38h-3v-.3a1.25 1.25 0 0 0-2.5 0v.3h-3a1.25 1.25 0 0 0-1.17 1.29l.41 12A1.26 1.26 0 0 0 4.2 16h7.59A1.24 1.24 0 0 0 13 14.79l.42-12a1.28 1.28 0 0 0-.31-.91M4.2 14.75l-.32-9.5h8.24l-.33 9.5zM12.16 4H3.84V2.75h8.42z"/><path fill="currentColor" d="M5 7h1.25v5H5zm2.37 0h1.25v5H7.37zm2.38 0H11v5H9.75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Truck(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.17 7.36L13 4.92a1.25 1.25 0 0 0-.94-.42h-2V2.75A1.25 1.25 0 0 0 8.82 1.5H1.76A1.25 1.25 0 0 0 .51 2.75v8.5a1.25 1.25 0 0 0 1.25 1.25h.33a2.07 2.07 0 0 0 2.13 2a2.08 2.08 0 0 0 2.14-2H10a2.07 2.07 0 0 0 2.13 2a2.08 2.08 0 0 0 2.14-2a1.26 1.26 0 0 0 1.2-1.25V8.19a1.22 1.22 0 0 0-.3-.83M4.22 13.25a.82.82 0 0 1-.88-.75a.82.82 0 0 1 .88-.75a.83.83 0 0 1 .89.75a.83.83 0 0 1-.89.75m4.6-7.58v5.58H5.89a2.17 2.17 0 0 0-1.67-.75a2.17 2.17 0 0 0-1.66.75h-.8v-8.5h7.06zm1.25.08h2l1.44 1.63h-3.44zm2.08 7.5a.82.82 0 0 1-.88-.75a.82.82 0 0 1 .88-.75a.83.83 0 0 1 .89.75a.83.83 0 0 1-.89.75m0-2.75a2.17 2.17 0 0 0-1.66.75h-.42V8.62h4.17v2.63h-.42a2.17 2.17 0 0 0-1.67-.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Undo(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.06.56A8.05 8.05 0 0 0 1.24 4.2V1.55H0V5a1.16 1.16 0 0 0 1.15 1.14h3.44V4.9H2.27a6.79 6.79 0 0 1 5.79-3.1a6.48 6.48 0 0 1 6.7 6.2a6.48 6.48 0 0 1-6.7 6.2A6.48 6.48 0 0 1 1.36 8H.12a7.71 7.71 0 0 0 7.94 7.44A7.71 7.71 0 0 0 16 8A7.71 7.71 0 0 0 8.06.56"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func University(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 6.28a1.23 1.23 0 0 0-.62-1.07l-6.74-4a1.27 1.27 0 0 0-1.28 0l-6.75 4a1.25 1.25 0 0 0 0 2.15l1.92 1.12v2.81a1.28 1.28 0 0 0 .62 1.09l4.25 2.45a1.28 1.28 0 0 0 1.24 0l4.25-2.45a1.28 1.28 0 0 0 .62-1.09V8.45l1.24-.73v2.72H16zm-3.73 5L8 13.74l-4.22-2.45V9.22l3.58 2.13a1.29 1.29 0 0 0 1.28 0l3.62-2.16zM8 10.27l-6.75-4L8 2.26l6.75 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Upload(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.5 14.75h11V16h-11zm1.23-8.37H6.2v5.87a1.25 1.25 0 0 0 1.25 1.25h1.1a1.25 1.25 0 0 0 1.25-1.25V6.38h2.47a1.25 1.25 0 0 0 .84-2.18L8.84.32a1.26 1.26 0 0 0-1.68 0L2.89 4.2a1.25 1.25 0 0 0 .84 2.18M8 1.25l4.27 3.88H8.55v7.12h-1.1V5.13H3.73z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func User(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 7.83c-3.08 0-5.59 2.17-5.59 4.84V16h1.27v-3.33c0-2 1.94-3.57 4.32-3.57s4.32 1.6 4.32 3.57V16h1.27v-3.33c0-2.67-2.51-4.84-5.59-4.84m.1-1.22a3.22 3.22 0 0 0 3.1-3.31A3.21 3.21 0 0 0 8.1 0A3.21 3.21 0 0 0 5 3.3a3.22 3.22 0 0 0 3.1 3.31m0-5.32a1.92 1.92 0 0 1 1.81 2a1.93 1.93 0 0 1-1.81 2a1.93 1.93 0 0 1-1.8-2a1.92 1.92 0 0 1 1.8-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserCircle(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 .5A7.77 7.77 0 0 0 0 8a7.77 7.77 0 0 0 8 7.5A7.77 7.77 0 0 0 16 8A7.77 7.77 0 0 0 8 .5M5.16 13.67A2.84 2.84 0 0 1 8 11.51a2.82 2.82 0 0 1 2.84 2.16a7.24 7.24 0 0 1-5.68 0m6.84-.61a4.09 4.09 0 0 0-4-2.77a4.09 4.09 0 0 0-3.95 2.78A6.14 6.14 0 0 1 1.25 8A6.52 6.52 0 0 1 8 1.75A6.52 6.52 0 0 1 14.75 8A6.11 6.11 0 0 1 12 13.06"/><path fill="currentColor" d="M8.05 4.3A2.33 2.33 0 0 0 5.8 6.7a2.33 2.33 0 0 0 2.25 2.4a2.33 2.33 0 0 0 2.25-2.4a2.33 2.33 0 0 0-2.25-2.4m0 3.55a1.08 1.08 0 0 1-1-1.15a1.08 1.08 0 0 1 1-1.15a1.08 1.08 0 0 1 1 1.15a1.08 1.08 0 0 1-1 1.15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserGroup(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.1 7.93A2.87 2.87 0 0 0 7.87 5A2.88 2.88 0 0 0 5.1 2a2.88 2.88 0 0 0-2.78 3A2.88 2.88 0 0 0 5.1 7.93m0-4.67A1.62 1.62 0 0 1 6.62 5A1.63 1.63 0 0 1 5.1 6.68A1.63 1.63 0 0 1 3.58 5A1.62 1.62 0 0 1 5.1 3.26m7.19 5.05a2.39 2.39 0 0 0 2.3-2.46a2.39 2.39 0 0 0-2.3-2.47A2.39 2.39 0 0 0 10 5.85a2.39 2.39 0 0 0 2.29 2.46m0-3.68a1.15 1.15 0 0 1 1.05 1.22a1.15 1.15 0 0 1-1.05 1.21a1.15 1.15 0 0 1-1.06-1.21a1.15 1.15 0 0 1 1.06-1.22m-.07 4.93a3.85 3.85 0 0 0-3.07 1.51A5.21 5.21 0 0 0 5.1 9.18A5 5 0 0 0 0 14h1.25a3.72 3.72 0 0 1 3.85-3.57A3.71 3.71 0 0 1 8.94 14h1.25a4.5 4.5 0 0 0-.32-1.69a2.54 2.54 0 0 1 2.35-1.5a2.44 2.44 0 0 1 2.53 2.33V14H16v-.86a3.69 3.69 0 0 0-3.78-3.58"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VerticalStacks(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.82 2.5H1.18C.53 2.5 0 3.03 0 3.68V6.2c0 .65.53 1.18 1.18 1.18h13.64c.65 0 1.18-.53 1.18-1.18V3.68c0-.65-.53-1.18-1.18-1.18m-.07 3.62H1.25V3.75h13.5zm.07 2.5H1.18C.53 8.62 0 9.15 0 9.8v2.52c0 .65.53 1.18 1.18 1.18h13.64c.65 0 1.18-.53 1.18-1.18V9.8c0-.65-.53-1.18-1.18-1.18m-.07 3.63H1.25V9.87h13.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Volume(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.32 7.45h-2v1.06h2V8a4.93 4.93 0 0 0 0-.55M12.91 11l1.56.37a6.28 6.28 0 0 0 .45-1L13.16 10zm2-5.43a6.28 6.28 0 0 0-.45-1l-1.55.34l.25 1zm-4.29-3.06L1.9 6.3v-.69H.65v4.78H1.9v-.71l8.72 3.79v1h1.25V1.5h-1.25zM1.9 8.32v-.65l8.72-3.8v8.24z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wallet(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.25 1.5H1.75A1.25 1.25 0 0 0 .5 2.75v1.11a1.85 1.85 0 0 1 1.25-.48v-.63h12.5v.63a1.85 1.85 0 0 1 1.25.48V2.75a1.25 1.25 0 0 0-1.25-1.25m0 2.5H1.75A1.25 1.25 0 0 0 .5 5.25v1.11a1.85 1.85 0 0 1 1.25-.48v-.63h12.5v.63a1.85 1.85 0 0 1 1.25.48V5.25A1.25 1.25 0 0 0 14.25 4m0 2.5H10.5A2.43 2.43 0 0 1 8 8.84A2.43 2.43 0 0 1 5.5 6.5H1.75A1.25 1.25 0 0 0 .5 7.75v5.5a1.25 1.25 0 0 0 1.25 1.25h12.5a1.25 1.25 0 0 0 1.25-1.25v-5.5a1.25 1.25 0 0 0-1.25-1.25m0 6.75H1.75v-5.5h2.73A3.76 3.76 0 0 0 8 10.09a3.75 3.75 0 0 0 3.52-2.34h2.73z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Whatsapp(children ...ElementRenderer) *NimbusIcon {
	return &NimbusIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.42 9.49c-.19-.09-1.1-.54-1.27-.61s-.29-.09-.42.1s-.48.6-.59.73s-.21.14-.4 0a5.13 5.13 0 0 1-1.49-.92a5.25 5.25 0 0 1-1-1.29c-.11-.18 0-.28.08-.38s.18-.21.28-.32a1.39 1.39 0 0 0 .18-.31a.38.38 0 0 0 0-.33c0-.09-.42-1-.58-1.37s-.3-.32-.41-.32h-.4a.72.72 0 0 0-.5.23a2.1 2.1 0 0 0-.65 1.55A3.59 3.59 0 0 0 5 8.2A8.32 8.32 0 0 0 8.19 11c.44.19.78.3 1.05.39a2.53 2.53 0 0 0 1.17.07a1.93 1.93 0 0 0 1.26-.88a1.67 1.67 0 0 0 .11-.88c-.05-.07-.17-.12-.36-.21"/><path fill="currentColor" d="M13.29 2.68A7.36 7.36 0 0 0 8 .5a7.44 7.44 0 0 0-6.41 11.15l-1 3.85l3.94-1a7.4 7.4 0 0 0 3.55.9H8a7.44 7.44 0 0 0 5.29-12.72M8 14.12a6.12 6.12 0 0 1-3.15-.87l-.22-.13l-2.34.61l.62-2.28l-.14-.23a6.18 6.18 0 0 1 9.6-7.65a6.12 6.12 0 0 1 1.81 4.37A6.19 6.19 0 0 1 8 14.12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
