package prime

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = ""
	hAttr          = "1em"
	viewbox        = "0 0 24 24"
	fill           = "currentColor"
)

type PrimeIcon struct {
	*SVGSVGElement
}

type PrimeIconFn func(children ...ElementRenderer) *PrimeIcon

var IconLookup = map[string]PrimeIconFn{
	"alignCenter":         AlignCenter,
	"alignJustify":        AlignJustify,
	"alignLeft":           AlignLeft,
	"alignRight":          AlignRight,
	"android":             Android,
	"angleDoubleDown":     AngleDoubleDown,
	"angleDoubleLeft":     AngleDoubleLeft,
	"angleDoubleRight":    AngleDoubleRight,
	"angleDoubleUp":       AngleDoubleUp,
	"angleDown":           AngleDown,
	"angleLeft":           AngleLeft,
	"angleRight":          AngleRight,
	"angleUp":             AngleUp,
	"apple":               Apple,
	"arrowCircleDown":     ArrowCircleDown,
	"arrowCircleLeft":     ArrowCircleLeft,
	"arrowCircleRight":    ArrowCircleRight,
	"arrowCircleUp":       ArrowCircleUp,
	"arrowDown":           ArrowDown,
	"arrowDownLeft":       ArrowDownLeft,
	"arrowDownRight":      ArrowDownRight,
	"arrowLeft":           ArrowLeft,
	"arrowRight":          ArrowRight,
	"arrowRightArrowLeft": ArrowRightArrowLeft,
	"arrowUp":             ArrowUp,
	"arrowUpLeft":         ArrowUpLeft,
	"arrowUpRight":        ArrowUpRight,
	"arrowsAlt":           ArrowsAlt,
	"arrowsH":             ArrowsH,
	"arrowsV":             ArrowsV,
	"at":                  At,
	"backward":            Backward,
	"ban":                 Ban,
	"bars":                Bars,
	"bell":                Bell,
	"bitcoin":             Bitcoin,
	"bolt":                Bolt,
	"book":                Book,
	"bookmark":            Bookmark,
	"bookmarkFill":        BookmarkFill,
	"box":                 Box,
	"briefcase":           Briefcase,
	"building":            Building,
	"calculator":          Calculator,
	"calendar":            Calendar,
	"calendarMinus":       CalendarMinus,
	"calendarPlus":        CalendarPlus,
	"calendarTimes":       CalendarTimes,
	"camera":              Camera,
	"car":                 Car,
	"caretDown":           CaretDown,
	"caretLeft":           CaretLeft,
	"caretRight":          CaretRight,
	"caretUp":             CaretUp,
	"cartPlus":            CartPlus,
	"chartBar":            ChartBar,
	"chartLine":           ChartLine,
	"chartPie":            ChartPie,
	"check":               Check,
	"checkCircle":         CheckCircle,
	"checkSquare":         CheckSquare,
	"chevronCircleDown":   ChevronCircleDown,
	"chevronCircleLeft":   ChevronCircleLeft,
	"chevronCircleRight":  ChevronCircleRight,
	"chevronCircleUp":     ChevronCircleUp,
	"chevronDown":         ChevronDown,
	"chevronLeft":         ChevronLeft,
	"chevronRight":        ChevronRight,
	"chevronUp":           ChevronUp,
	"circle":              Circle,
	"circleFill":          CircleFill,
	"circleOff":           CircleOff,
	"circleOn":            CircleOn,
	"clock":               Clock,
	"clone":               Clone,
	"cloud":               Cloud,
	"cloudDownload":       CloudDownload,
	"cloudUpload":         CloudUpload,
	"code":                Code,
	"cog":                 Cog,
	"comment":             Comment,
	"comments":            Comments,
	"compass":             Compass,
	"copy":                Copy,
	"creditCard":          CreditCard,
	"database":            Database,
	"deleteLeft":          DeleteLeft,
	"desktop":             Desktop,
	"directions":          Directions,
	"directionsAlt":       DirectionsAlt,
	"discord":             Discord,
	"dollar":              Dollar,
	"download":            Download,
	"eject":               Eject,
	"ellipsisH":           EllipsisH,
	"ellipsisV":           EllipsisV,
	"envelope":            Envelope,
	"eraser":              Eraser,
	"euro":                Euro,
	"exclamationCircle":   ExclamationCircle,
	"exclamationTriangle": ExclamationTriangle,
	"externalLink":        ExternalLink,
	"eye":                 Eye,
	"eyeSlash":            EyeSlash,
	"facebook":            Facebook,
	"fastBackward":        FastBackward,
	"fastForward":         FastForward,
	"file":                File,
	"fileEdit":            FileEdit,
	"fileExcel":           FileExcel,
	"fileExport":          FileExport,
	"fileImport":          FileImport,
	"fileO":               FileO,
	"filePdf":             FilePdf,
	"fileWord":            FileWord,
	"filter":              Filter,
	"filterFill":          FilterFill,
	"filterSlash":         FilterSlash,
	"flag":                Flag,
	"flagFill":            FlagFill,
	"folder":              Folder,
	"folderOpen":          FolderOpen,
	"forward":             Forward,
	"gift":                Gift,
	"github":              Github,
	"globe":               Globe,
	"google":              Google,
	"hashtag":             Hashtag,
	"heart":               Heart,
	"heartFill":           HeartFill,
	"history":             History,
	"home":                Home,
	"hourglass":           Hourglass,
	"idCard":              IdCard,
	"image":               Image,
	"images":              Images,
	"inbox":               Inbox,
	"info":                Info,
	"infoCircle":          InfoCircle,
	"instagram":           Instagram,
	"key":                 Key,
	"language":            Language,
	"link":                Link,
	"linkedin":            Linkedin,
	"list":                List,
	"lock":                Lock,
	"lockOpen":            LockOpen,
	"map":                 Map,
	"mapMarker":           MapMarker,
	"megaphone":           Megaphone,
	"microphone":          Microphone,
	"microsoft":           Microsoft,
	"minus":               Minus,
	"minusCircle":         MinusCircle,
	"mobile":              Mobile,
	"moneyBill":           MoneyBill,
	"moon":                Moon,
	"palette":             Palette,
	"paperclip":           Paperclip,
	"pause":               Pause,
	"paypal":              Paypal,
	"pencil":              Pencil,
	"percentage":          Percentage,
	"phone":               Phone,
	"play":                Play,
	"plus":                Plus,
	"plusCircle":          PlusCircle,
	"pound":               Pound,
	"powerOff":            PowerOff,
	"prime":               Prime,
	"print":               Print,
	"qrcode":              Qrcode,
	"question":            Question,
	"questionCircle":      QuestionCircle,
	"reddit":              Reddit,
	"refresh":             Refresh,
	"replay":              Replay,
	"reply":               Reply,
	"save":                Save,
	"search":              Search,
	"searchMinus":         SearchMinus,
	"searchPlus":          SearchPlus,
	"send":                Send,
	"server":              Server,
	"shareAlt":            ShareAlt,
	"shield":              Shield,
	"shoppingBag":         ShoppingBag,
	"shoppingCart":        ShoppingCart,
	"signIn":              SignIn,
	"signOut":             SignOut,
	"sitemap":             Sitemap,
	"slack":               Slack,
	"slidersH":            SlidersH,
	"slidersV":            SlidersV,
	"sort":                Sort,
	"sortAlphaAltDown":    SortAlphaAltDown,
	"sortAlphaAltUp":      SortAlphaAltUp,
	"sortAlphaDown":       SortAlphaDown,
	"sortAlphaUp":         SortAlphaUp,
	"sortAlt":             SortAlt,
	"sortAltSlash":        SortAltSlash,
	"sortAmountDown":      SortAmountDown,
	"sortAmountDownAlt":   SortAmountDownAlt,
	"sortAmountUp":        SortAmountUp,
	"sortAmountUpAlt":     SortAmountUpAlt,
	"sortDown":            SortDown,
	"sortNumericAltDown":  SortNumericAltDown,
	"sortNumericAltUp":    SortNumericAltUp,
	"sortNumericDown":     SortNumericDown,
	"sortNumericUp":       SortNumericUp,
	"sortUp":              SortUp,
	"spinner":             Spinner,
	"star":                Star,
	"starFill":            StarFill,
	"stepBackward":        StepBackward,
	"stepBackwardAlt":     StepBackwardAlt,
	"stepForward":         StepForward,
	"stepForwardAlt":      StepForwardAlt,
	"stop":                Stop,
	"stopCircle":          StopCircle,
	"stopwatch":           Stopwatch,
	"sun":                 Sun,
	"sync":                Sync,
	"table":               Table,
	"tablet":              Tablet,
	"tag":                 Tag,
	"tags":                Tags,
	"telegram":            Telegram,
	"thLarge":             ThLarge,
	"thumbsDown":          ThumbsDown,
	"thumbsDownFill":      ThumbsDownFill,
	"thumbsUp":            ThumbsUp,
	"thumbsUpFill":        ThumbsUpFill,
	"ticket":              Ticket,
	"times":               Times,
	"timesCircle":         TimesCircle,
	"trash":               Trash,
	"truck":               Truck,
	"twitter":             Twitter,
	"undo":                Undo,
	"unlock":              Unlock,
	"upload":              Upload,
	"user":                User,
	"userEdit":            UserEdit,
	"userMinus":           UserMinus,
	"userPlus":            UserPlus,
	"users":               Users,
	"verified":            Verified,
	"video":               Video,
	"vimeo":               Vimeo,
	"volumeDown":          VolumeDown,
	"volumeOff":           VolumeOff,
	"volumeUp":            VolumeUp,
	"wallet":              Wallet,
	"whatsapp":            Whatsapp,
	"wifi":                Wifi,
	"windowMaximize":      WindowMaximize,
	"windowMinimize":      WindowMinimize,
	"wrench":              Wrench,
	"youtube":             Youtube,
}

func AlignCenter(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 10.75H7a.75.75 0 0 1 0-1.5h10a.75.75 0 0 1 0 1.5m3-4H4a.75.75 0 0 1 0-1.5h16a.75.75 0 0 1 0 1.5m0 8H4a.75.75 0 0 1 0-1.5h16a.75.75 0 0 1 0 1.5m-3 4H7a.75.75 0 0 1 0-1.5h10a.75.75 0 0 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignJustify(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 10.75H4a.75.75 0 0 1 0-1.5h16a.75.75 0 0 1 0 1.5m0-4H4a.75.75 0 0 1 0-1.5h16a.75.75 0 0 1 0 1.5m0 8H4a.75.75 0 0 1 0-1.5h16a.75.75 0 0 1 0 1.5m0 4H4a.75.75 0 0 1 0-1.5h16a.75.75 0 0 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignLeft(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.93 11h-10a.75.75 0 1 1 0-1.5h10a.75.75 0 0 1 0 1.5m6.14-4h-16a.75.75 0 0 1 0-1.5h16a.75.75 0 0 1 0 1.5m0 8h-16a.75.75 0 0 1 0-1.5h16a.75.75 0 0 1 0 1.5m-6.14 4h-10a.75.75 0 1 1 0-1.5h10a.75.75 0 0 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignRight(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 10.75H10a.75.75 0 0 1 0-1.5h10a.75.75 0 0 1 0 1.5m.07-4h-16a.75.75 0 0 1 0-1.5h16a.75.75 0 0 1 0 1.5m0 8h-16a.75.75 0 0 1 0-1.5h16a.75.75 0 0 1 0 1.5m-.07 4H10a.75.75 0 0 1 0-1.5h10a.75.75 0 0 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Android(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.06 9.33a.94.94 0 0 0-.94.94v4.26a.94.94 0 0 0 1.88 0v-4.26a.94.94 0 0 0-.94-.94m-12.12 0a.94.94 0 0 0-.94.94v4.26a.94.94 0 0 0 1.88 0v-4.26a.94.94 0 0 0-.94-.94m1.62 0v6.4a1.07 1.07 0 0 0 1.07 1.07h.68v2.27a.94.94 0 0 0 1.88 0V16.8h1.62v2.27a.94.94 0 0 0 1.88 0V16.8h.68a1.07 1.07 0 0 0 1.07-1.07v-6.4Z"/><circle cx="9.98" cy="7.07" r=".4" fill="none"/><circle cx="14.02" cy="7.07" r=".4" fill="none"/><path fill="currentColor" d="M16.32 8a4.13 4.13 0 0 0-1.83-2.36l-.15-.09l-.16-.08l.18-.31l.53-1a.14.14 0 0 0-.05-.16h-.07a.17.17 0 0 0-.13.07L14.1 5l-.17.31l-.16-.07l-.17-.06a4.88 4.88 0 0 0-3.2 0l-.16.06l-.17.07L9.9 5l-.54-1a.14.14 0 0 0-.25.14l.53 1l.18.31l-.16.08l-.15.09A4.07 4.07 0 0 0 7.69 8a3.11 3.11 0 0 0-.13.8h8.88a3.6 3.6 0 0 0-.12-.8M10 7.47a.4.4 0 1 1 .4-.4a.4.4 0 0 1-.4.4m4 0a.4.4 0 1 1 .4-.4a.4.4 0 0 1-.4.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleDoubleDown(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 17.25a.74.74 0 0 1-.53-.22L8 13.53a.75.75 0 0 1 1-1.06l3 3l3-3a.75.75 0 0 1 1 1.06L12.53 17a.74.74 0 0 1-.53.25m0-5.5a.74.74 0 0 1-.53-.22L8 8a.75.75 0 0 1 1-1l3 3l3-3a.75.75 0 0 1 1 1l-3.5 3.5a.74.74 0 0 1-.5.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleDoubleLeft(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 16.25a.74.74 0 0 1-.53-.22L7 12.53a.75.75 0 0 1 0-1.06L10.47 8a.75.75 0 0 1 1.06 1l-3 3l3 3a.75.75 0 0 1 0 1.06a.74.74 0 0 1-.53.19m5.5 0A.74.74 0 0 1 16 16l-3.5-3.5a.75.75 0 0 1 0-1.06L16 8a.75.75 0 0 1 1 1l-3 3l3 3a.75.75 0 0 1 0 1a.74.74 0 0 1-.5.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleDoubleRight(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 16.25a.74.74 0 0 1-.53-.22a.75.75 0 0 1 0-1.06l3-3l-3-3A.75.75 0 0 1 13.53 8l3.5 3.5a.75.75 0 0 1 0 1.06L13.53 16a.74.74 0 0 1-.53.25m-5.5 0A.74.74 0 0 1 7 16a.75.75 0 0 1 0-1l3-3l-3-3a.75.75 0 0 1 1-1l3.5 3.5a.75.75 0 0 1 0 1.06L8 16a.74.74 0 0 1-.5.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleDoubleUp(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.5 11.75a.74.74 0 0 1-.53-.22l-3-3l-3 3A.75.75 0 0 1 8 10.47L11.47 7a.75.75 0 0 1 1.06 0l3.5 3.5a.75.75 0 0 1 0 1.06a.74.74 0 0 1-.53.19m0 5.5A.74.74 0 0 1 15 17l-3-3l-3 3a.75.75 0 0 1-1-1l3.5-3.5a.75.75 0 0 1 1.06 0L16 16a.75.75 0 0 1 0 1a.74.74 0 0 1-.5.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleDown(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 14.5a.74.74 0 0 1-.53-.22L8 10.78a.75.75 0 0 1 1-1.06l3 3l3-3a.75.75 0 0 1 1 1.06l-3.5 3.5a.74.74 0 0 1-.5.22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleLeft(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.75 16.25a.74.74 0 0 1-.53-.22l-3.5-3.5a.75.75 0 0 1 0-1.06L13.22 8a.75.75 0 0 1 1.06 1l-3 3l3 3a.75.75 0 0 1 0 1.06a.74.74 0 0 1-.53.19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleRight(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.25 16.25a.74.74 0 0 1-.53-.25a.75.75 0 0 1 0-1.06l3-3l-3-3A.75.75 0 0 1 10.78 8l3.5 3.5a.75.75 0 0 1 0 1.06L10.78 16a.74.74 0 0 1-.53.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleUp(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.5 14.5a.74.74 0 0 1-.53-.22l-3-3l-3 3A.75.75 0 0 1 8 13.22l3.5-3.5a.75.75 0 0 1 1.06 0l3.5 3.5a.75.75 0 0 1 0 1.06a.74.74 0 0 1-.56.22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Apple(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.52 12.46a3.31 3.31 0 0 1 1.78-3a3.85 3.85 0 0 0-3-1.6c-1.27-.1-2.65.74-3.16.74s-1.77-.7-2.73-.7c-2 0-4.11 1.59-4.11 4.76a9 9 0 0 0 .51 2.9c.44 1.28 2.09 4.49 3.81 4.44c.9 0 1.54-.64 2.71-.64s1.72.64 2.73.64c1.73 0 3.23-2.95 3.66-4.26a3.54 3.54 0 0 1-2.2-3.28m-2-5.87A3.37 3.37 0 0 0 15.35 4a3.79 3.79 0 0 0-2.42 1.25A3.41 3.41 0 0 0 12 7.81a3 3 0 0 0 2.5-1.22Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleDown(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21a9 9 0 1 1 9-9a9 9 0 0 1-9 9m0-16.5a7.5 7.5 0 1 0 7.5 7.5A7.5 7.5 0 0 0 12 4.5"/><path fill="currentColor" d="M12 16.75a.74.74 0 0 1-.53-.22l-4-4a.75.75 0 0 1 1.06-1.06L12 14.94l3.47-3.47a.75.75 0 0 1 1.06 1.06l-4 4a.74.74 0 0 1-.53.22"/><path fill="currentColor" d="M12 16.75a.76.76 0 0 1-.75-.75V8a.75.75 0 0 1 1.5 0v8a.76.76 0 0 1-.75.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleLeft(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21a9 9 0 1 1 9-9a9 9 0 0 1-9 9m0-16.5a7.5 7.5 0 1 0 7.5 7.5A7.5 7.5 0 0 0 12 4.5"/><path fill="currentColor" d="M12 16.75a.74.74 0 0 1-.53-.22l-4-4a.75.75 0 0 1 0-1.06l4-4a.75.75 0 0 1 1.06 1.06L9.06 12l3.47 3.47a.75.75 0 0 1 0 1.06a.74.74 0 0 1-.53.22"/><path fill="currentColor" d="M16 12.75H8a.75.75 0 0 1 0-1.5h8a.75.75 0 0 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleRight(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21a9 9 0 1 1 9-9a9 9 0 0 1-9 9m0-16.5a7.5 7.5 0 1 0 7.5 7.5A7.5 7.5 0 0 0 12 4.5"/><path fill="currentColor" d="M12 16.75a.74.74 0 0 1-.53-.22a.75.75 0 0 1 0-1.06L14.94 12l-3.47-3.47a.75.75 0 0 1 1.06-1.06l4 4a.75.75 0 0 1 0 1.06l-4 4a.74.74 0 0 1-.53.22"/><path fill="currentColor" d="M16 12.75H8a.75.75 0 0 1 0-1.5h8a.75.75 0 0 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleUp(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21a9 9 0 1 1 9-9a9 9 0 0 1-9 9m0-16.5a7.5 7.5 0 1 0 7.5 7.5A7.5 7.5 0 0 0 12 4.5"/><path fill="currentColor" d="M16 12.75a.74.74 0 0 1-.53-.22L12 9.06l-3.47 3.47a.75.75 0 0 1-1.06-1.06l4-4a.75.75 0 0 1 1.06 0l4 4a.75.75 0 0 1 0 1.06a.74.74 0 0 1-.53.22"/><path fill="currentColor" d="M12 16.75a.76.76 0 0 1-.75-.75V8a.75.75 0 0 1 1.5 0v8a.76.76 0 0 1-.75.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDown(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 19.75a.74.74 0 0 1-.53-.22l-6-6a.75.75 0 0 1 1.06-1.06L12 17.94l5.47-5.47a.75.75 0 0 1 1.06 1.06l-6 6a.74.74 0 0 1-.53.22"/><path fill="currentColor" d="M12 19.75a.76.76 0 0 1-.75-.75V5a.75.75 0 0 1 1.5 0v14a.76.76 0 0 1-.75.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownLeft(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.54 17.7a.75.75 0 0 0 0-1.5H8.86l8.62-8.62a.75.75 0 1 0-1.06-1.06L7.8 15.14V8.46a.75.75 0 0 0-1.5 0V17a.75.75 0 0 0 .06.29a.76.76 0 0 0 .69.46Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownRight(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.7 8.46a.75.75 0 1 0-1.5 0v6.68L7.58 6.52a.75.75 0 0 0-1.06 1.06l8.62 8.62H8.46a.75.75 0 0 0 0 1.5H17a.75.75 0 0 0 .29-.06a.76.76 0 0 0 .41-.64Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeft(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 18.75a.74.74 0 0 1-.53-.22l-6-6a.75.75 0 0 1 0-1.06l6-6a.75.75 0 0 1 1.06 1.06L6.06 12l5.47 5.47a.75.75 0 0 1 0 1.06a.74.74 0 0 1-.53.22"/><path fill="currentColor" d="M19 12.75H5a.75.75 0 0 1 0-1.5h14a.75.75 0 0 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRight(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 18.75a.74.74 0 0 1-.53-.22a.75.75 0 0 1 0-1.06L17.94 12l-5.47-5.47a.75.75 0 0 1 1.06-1.06l6 6a.75.75 0 0 1 0 1.06l-6 6a.74.74 0 0 1-.53.22"/><path fill="currentColor" d="M19 12.75H5a.75.75 0 0 1 0-1.5h14a.75.75 0 0 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightArrowLeft(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.75 16c0 .41-.34.75-.75.75H6.81l1.22 1.22c.29.29.29.77 0 1.06c-.15.15-.34.22-.53.22s-.38-.07-.53-.22l-2.5-2.5a.776.776 0 0 1-.16-.24a.707.707 0 0 1 0-.57c.04-.09.09-.17.16-.24l2.5-2.5c.29-.29.77-.29 1.06 0s.29.77 0 1.06l-1.22 1.22H19c.41 0 .75.34.75.75ZM5 8.75h12.19l-1.22 1.22c-.29.29-.29.77 0 1.06c.15.15.34.22.53.22s.38-.07.53-.22l2.5-2.5c.07-.07.12-.15.16-.24c.08-.18.08-.39 0-.57a.776.776 0 0 0-.16-.24l-2.5-2.5c-.29-.29-.77-.29-1.06 0s-.29.77 0 1.06l1.22 1.22H5c-.41 0-.75.34-.75.75s.34.75.75.75Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUp(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 11.75a.74.74 0 0 1-.53-.22L12 6.06l-5.47 5.47a.75.75 0 0 1-1.06-1.06l6-6a.75.75 0 0 1 1.06 0l6 6a.75.75 0 0 1 0 1.06a.74.74 0 0 1-.53.22"/><path fill="currentColor" d="M12 19.75a.76.76 0 0 1-.75-.75V5a.75.75 0 0 1 1.5 0v14a.76.76 0 0 1-.75.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpLeft(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.3 15.54a.75.75 0 0 0 1.5 0V8.86l8.62 8.62a.75.75 0 1 0 1.06-1.06L8.86 7.8h6.68a.75.75 0 0 0 0-1.5H7.05a.75.75 0 0 0-.29.06a.76.76 0 0 0-.46.69Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpRight(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.46 6.3a.75.75 0 1 0 0 1.5h6.68l-8.62 8.62a.75.75 0 1 0 1.06 1.06l8.62-8.62v6.68a.75.75 0 0 0 1.5 0V7.05a.75.75 0 0 0-.06-.29a.76.76 0 0 0-.64-.46Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsAlt(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.69 12.29c-.04.09-.09.17-.16.24l-2.5 2.5c-.15.15-.34.22-.53.22s-.38-.07-.53-.22a.754.754 0 0 1 0-1.06l1.22-1.22h-5.44v5.44l1.22-1.22c.29-.29.77-.29 1.06 0s.29.77 0 1.06l-2.5 2.5a.776.776 0 0 1-.53.22a.776.776 0 0 1-.53-.22l-2.5-2.5c-.29-.29-.29-.77 0-1.06s.77-.29 1.06 0l1.22 1.22v-5.44H5.81l1.22 1.22c.29.29.29.77 0 1.06c-.15.15-.34.22-.53.22s-.38-.07-.53-.22l-2.5-2.5a.776.776 0 0 1-.16-.24a.707.707 0 0 1 0-.57c.04-.09.09-.17.16-.24l2.5-2.5c.29-.29.77-.29 1.06 0s.29.77 0 1.06l-1.22 1.22h5.44V5.81l-1.22 1.22c-.29.29-.77.29-1.06 0s-.29-.77 0-1.06l2.5-2.5c.07-.07.15-.12.24-.16c.18-.08.39-.08.57 0c.09.04.17.09.24.16l2.5 2.5c.29.29.29.77 0 1.06c-.15.15-.34.22-.53.22s-.38-.07-.53-.22l-1.22-1.22v5.44h5.44l-1.22-1.22c-.29-.29-.29-.77 0-1.06s.77-.29 1.06 0l2.5 2.5c.07.07.12.15.16.24c.08.18.08.39 0 .57Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsH(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.69 11.71a.78.78 0 0 0-.16-.24l-4-4a.75.75 0 1 0-1.06 1.06l2.72 2.72H5.81l2.72-2.72a.75.75 0 0 0-1.06-1.06l-4 4a.78.78 0 0 0-.16.24a.73.73 0 0 0 0 .58a.78.78 0 0 0 .16.24l4 4a.75.75 0 0 0 1.06 0a.75.75 0 0 0 0-1.06l-2.72-2.72h12.38l-2.72 2.72a.75.75 0 0 0 0 1.06a.75.75 0 0 0 1.06 0l4-4a.78.78 0 0 0 .16-.24a.73.73 0 0 0 0-.58"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsV(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.29 20.69a.78.78 0 0 0 .24-.16l4-4a.75.75 0 0 0-1.06-1.06l-2.72 2.72V5.81l2.72 2.72a.75.75 0 0 0 1.06-1.06l-4-4a.78.78 0 0 0-.24-.16a.73.73 0 0 0-.58 0a.78.78 0 0 0-.24.16l-4 4a.75.75 0 0 0 0 1.06a.75.75 0 0 0 1.06 0l2.72-2.72v12.38l-2.72-2.72a.75.75 0 0 0-1.06 0a.75.75 0 0 0 0 1.06l4 4a.78.78 0 0 0 .24.16a.73.73 0 0 0 .58 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func At(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 3.25A8.75 8.75 0 0 0 3.25 12A8.65 8.65 0 0 0 12 20.75a.75.75 0 0 0 0-1.5A7.17 7.17 0 0 1 4.75 12A7.26 7.26 0 0 1 12 4.75c4.81 0 7.25 2.44 7.25 7.25v1.38a1.46 1.46 0 1 1-2.91 0v-5a.75.75 0 0 0-1.5 0v.34A4.31 4.31 0 0 0 12 7.66a4.34 4.34 0 0 0 0 8.68a4.3 4.3 0 0 0 3.24-1.49a2.95 2.95 0 0 0 5.51-1.47V12c0-5.64-3.11-8.75-8.75-8.75m0 11.59A2.84 2.84 0 1 1 14.84 12A2.85 2.85 0 0 1 12 14.84"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Backward(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.3 4.31a.756.756 0 0 0-.81.14l-7.27 6.82V5c0-.3-.18-.57-.45-.69a.756.756 0 0 0-.81.14l-7.47 7c-.15.14-.24.34-.24.55s.09.41.24.55l7.47 7a.763.763 0 0 0 .81.14c.27-.12.45-.39.45-.69v-6.27l7.27 6.82a.763.763 0 0 0 .81.14c.27-.12.45-.39.45-.69V5c0-.3-.18-.57-.45-.69m-9.58 12.96L5.1 12l5.62-5.27zm8.53 0L13.63 12l5.62-5.27z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ban(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 3a9 9 0 1 0 9 9a9 9 0 0 0-9-9m-7.5 9a7.44 7.44 0 0 1 1.7-4.74L16.74 17.8A7.49 7.49 0 0 1 4.5 12m13.3 4.74L7.26 6.2A7.49 7.49 0 0 1 17.8 16.74"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bars(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 12.75H5a.75.75 0 0 1 0-1.5h14a.75.75 0 0 1 0 1.5m0-4.5H5a.75.75 0 0 1 0-1.5h14a.75.75 0 0 1 0 1.5m0 9H5a.75.75 0 0 1 0-1.5h14a.75.75 0 0 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bell(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.53 16.25c-.09 0-2.11-.36-2.11-6.25c0-4.16-2.42-6.75-6.42-6.75S5.58 5.84 5.58 10c0 6-2.09 6.25-2.08 6.25a.75.75 0 0 0 0 1.5h4.83a3.74 3.74 0 0 0 7.34 0h4.84a.75.75 0 0 0 0-1.5Zm-8.53 3a2.24 2.24 0 0 1-2.11-1.5h4.22a2.24 2.24 0 0 1-2.11 1.5m-6.24-3c.72-1.09 1.32-3 1.32-6.25S8.88 4.75 12 4.75s4.92 1.91 4.92 5.25s.6 5.16 1.32 6.25Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bitcoin(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.52 13.67c-.27 1.09-2.1.5-2.7.35l.48-1.93c.6.15 2.5.44 2.22 1.58m-1.56-4.22l-.44 1.75c.5.12 2.03.63 2.28-.36c.25-1.03-1.34-1.27-1.84-1.39m7.8 4.48c-1.07 4.28-5.41 6.89-9.7 5.83c-4.28-1.07-6.9-5.41-5.83-9.69c1.07-4.29 5.41-6.89 9.7-5.83c4.28 1.07 6.9 5.41 5.83 9.69m-10.39-.59c-.04.11-.15.27-.4.21c-.04 0-.64-.16-.64-.16l-.44 1l1.14.28c.22.06.42.11.63.17l-.36 1.45l.87.22l.36-1.44c.24.07.47.12.7.18l-.36 1.43l.88.22l.36-1.45c1.5.28 2.63.17 3.1-1.19c.38-1.09-.02-1.72-.81-2.13c.57-.13 1-.51 1.12-1.29c.16-1.06-.65-1.64-1.76-2.01l.36-1.45l-.88-.22l-.35 1.4c-.23-.06-.47-.11-.71-.17l.35-1.41l-.87-.22l-.36 1.44c-.19-.04-.38-.09-.56-.13l-1.21-.31l-.23.94s.64.15.64.16c.35.09.41.32.41.51l-.99 3.95Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bolt(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.09 21.5a.67.67 0 0 1-.24 0a.83.83 0 0 1-.59-.81v-.11l.9-6.35H6.82a.8.8 0 0 1-.71-.43a.85.85 0 0 1 0-.86l2-3.49l4.1-6.52a.79.79 0 0 1 .92-.35a.83.83 0 0 1 .59.81v.11l-.9 6.35h4.35a.8.8 0 0 1 .71.43a.85.85 0 0 1 0 .86l-2 3.49l-4.1 6.52a.79.79 0 0 1-.69.35m-3.16-8.81h4a.84.84 0 0 1 .83.85v.11l-.59 4.14l2.5-4l1.44-2.48h-4a.84.84 0 0 1-.83-.85v-.11l.59-4.14l-2.5 4Zm5.51-8.99"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Book(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 3.25H6.75a2.43 2.43 0 0 0-2.5 2.35V18a2.85 2.85 0 0 0 2.94 2.75H19a.76.76 0 0 0 .75-.75V4a.76.76 0 0 0-.75-.75m-.75 16H7.19A1.35 1.35 0 0 1 5.75 18a1.35 1.35 0 0 1 1.44-1.25h11.06Zm0-4H7.19a3 3 0 0 0-1.44.37V5.6a.94.94 0 0 1 1-.85h11.5Z"/><path fill="currentColor" d="M8.75 8.75h6.5a.75.75 0 0 0 0-1.5h-6.5a.75.75 0 0 0 0 1.5m0 3.5h6.5a.75.75 0 0 0 0-1.5h-6.5a.75.75 0 0 0 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookmark(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.75 20.75a.83.83 0 0 1-.43-.13L12 16.91l-5.32 3.71a.75.75 0 0 1-.78 0a.74.74 0 0 1-.4-.62V6a2.75 2.75 0 0 1 2.75-2.75h7.5A2.75 2.75 0 0 1 18.5 6v14a.74.74 0 0 1-.4.66a.73.73 0 0 1-.35.09M12 15.25a.75.75 0 0 1 .43.13L17 18.56V6a1.25 1.25 0 0 0-1.25-1.25h-7.5A1.25 1.25 0 0 0 7 6v12.56l4.57-3.18a.75.75 0 0 1 .43-.13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkFill(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.75 3.25h-7.5A2.75 2.75 0 0 0 5.5 6v14a.75.75 0 0 0 1.18.62L12 16.91l5.32 3.71a.75.75 0 0 0 .43.13a.85.85 0 0 0 .35-.08a.77.77 0 0 0 .4-.67V6a2.75 2.75 0 0 0-2.75-2.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Box(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.73 16.52V7.59a.73.73 0 0 0-.08-.33a.74.74 0 0 0-.36-.36l-8-3.58a.75.75 0 0 0-.62 0l-8 3.58a.8.8 0 0 0-.44.69v8.82a.83.83 0 0 0 .44.69l8 3.58a.72.72 0 0 0 .62 0l8-3.58a.77.77 0 0 0 .44-.58m-16-7.78l6.5 2.92v7.18l-6.5-2.91Zm8 2.92l6.5-2.92v7.19l-6.5 2.91ZM12 4.82l6.17 2.77L12 10.35L5.83 7.59Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Briefcase(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 6.25h-3.75V5a1.89 1.89 0 0 0-2-1.75h-2.5a1.89 1.89 0 0 0-2 1.75v1.25H5A1.76 1.76 0 0 0 3.25 8v10A1.76 1.76 0 0 0 5 19.75h14A1.76 1.76 0 0 0 20.75 18V8A1.76 1.76 0 0 0 19 6.25M10.25 5c0-.08.19-.25.5-.25h2.5c.31 0 .5.17.5.25v1.25h-3.5ZM5 7.75h14a.25.25 0 0 1 .25.25v3.25H4.75V8A.25.25 0 0 1 5 7.75m3.75 5h6.5v1.5h-6.5ZM19 18.25H5a.25.25 0 0 1-.25-.25v-5.25h2.5V15a.76.76 0 0 0 .75.75h8a.76.76 0 0 0 .75-.75v-2.25h2.5V18a.25.25 0 0 1-.25.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Building(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<rect width="2" height="2" x="9" y="6" fill="currentColor" rx=".5"/><rect width="2" height="2" x="13" y="6" fill="currentColor" rx=".5"/><rect width="2" height="2" x="9" y="9.5" fill="currentColor" rx=".5"/><rect width="2" height="2" x="13" y="9.5" fill="currentColor" rx=".5"/><rect width="2" height="2" x="9" y="13" fill="currentColor" rx=".5"/><rect width="2" height="2" x="13" y="13" fill="currentColor" rx=".5"/><path fill="currentColor" d="M18.25 19.25h-.5V4a.76.76 0 0 0-.75-.75H7a.76.76 0 0 0-.75.75v15.25h-.5a.75.75 0 0 0 0 1.5h12.5a.75.75 0 0 0 0-1.5m-2 0H11V17a.5.5 0 0 0-.5-.5h-1a.5.5 0 0 0-.5.5v2.25H7.75V4.75h8.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calculator(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.5 8.25h-7c-.28 0-.5-.22-.5-.5v-1.5c0-.28.22-.5.5-.5h7c.28 0 .5.22.5.5v1.5c0 .28-.22.5-.5.5m.5 3v-1c0-.28-.22-.5-.5-.5h-1c-.28 0-.5.22-.5.5v1c0 .28.22.5.5.5h1c.28 0 .5-.22.5-.5m-6 0v-1c0-.28-.22-.5-.5-.5h-1c-.28 0-.5.22-.5.5v1c0 .28.22.5.5.5h1c.28 0 .5-.22.5-.5m3 0v-1c0-.28-.22-.5-.5-.5h-1c-.28 0-.5.22-.5.5v1c0 .28.22.5.5.5h1c.28 0 .5-.22.5-.5m3 3.25v-1c0-.28-.22-.5-.5-.5h-1c-.28 0-.5.22-.5.5v1c0 .28.22.5.5.5h1c.28 0 .5-.22.5-.5m-6 0v-1c0-.28-.22-.5-.5-.5h-1c-.28 0-.5.22-.5.5v1c0 .28.22.5.5.5h1c.28 0 .5-.22.5-.5m3 0v-1c0-.28-.22-.5-.5-.5h-1c-.28 0-.5.22-.5.5v1c0 .28.22.5.5.5h1c.28 0 .5-.22.5-.5m3 3.25v-1c0-.28-.22-.5-.5-.5h-1c-.28 0-.5.22-.5.5v1c0 .28.22.5.5.5h1c.28 0 .5-.22.5-.5m-3 0v-1c0-.28-.22-.5-.5-.5h-4c-.28 0-.5.22-.5.5v1c0 .28.22.5.5.5h4c.28 0 .5-.22.5-.5m6.25 1.75v-15c0-1.24-1.01-2.25-2.25-2.25H7c-1.24 0-2.25 1.01-2.25 2.25v15c0 1.24 1.01 2.25 2.25 2.25h10c1.24 0 2.25-1.01 2.25-2.25M17 3.75c.41 0 .75.34.75.75v15c0 .41-.34.75-.75.75H7c-.41 0-.75-.34-.75-.75v-15c0-.41.34-.75.75-.75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 4.25h-1.25V3a.75.75 0 0 0-1.5 0v1.25h-4.5V3a.75.75 0 0 0-1.5 0v1.25H7A2.75 2.75 0 0 0 4.25 7v11A2.75 2.75 0 0 0 7 20.75h10A2.75 2.75 0 0 0 19.75 18V7A2.75 2.75 0 0 0 17 4.25M7 5.75h1.25V7a.75.75 0 0 0 1.5 0V5.75h4.5V7a.75.75 0 0 0 1.5 0V5.75H17A1.25 1.25 0 0 1 18.25 7v2.75H5.75V7A1.25 1.25 0 0 1 7 5.75m10 13.5H7A1.25 1.25 0 0 1 5.75 18v-6.75h12.5V18A1.25 1.25 0 0 1 17 19.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarMinus(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 4.75h-1.25V3.5a.75.75 0 0 0-1.5 0v1.25h-4.5V3.5a.75.75 0 0 0-1.5 0v1.25H7A2.75 2.75 0 0 0 4.25 7.5v11A2.75 2.75 0 0 0 7 21.25h10a2.75 2.75 0 0 0 2.75-2.75v-11A2.75 2.75 0 0 0 17 4.75M7 6.25h1.25V7.5a.75.75 0 0 0 1.5 0V6.25h4.5V7.5a.75.75 0 0 0 1.5 0V6.25H17a1.25 1.25 0 0 1 1.25 1.25v2.75H5.75V7.5A1.25 1.25 0 0 1 7 6.25m10 13.5H7a1.25 1.25 0 0 1-1.25-1.25v-6.75h12.5v6.75A1.25 1.25 0 0 1 17 19.75"/><path fill="currentColor" d="M14 15.3h-4a.75.75 0 0 0 0 1.5h4a.75.75 0 0 0 0-1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarPlus(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 4.75h-1.25V3.5a.75.75 0 0 0-1.5 0v1.25h-4.5V3.5a.75.75 0 0 0-1.5 0v1.25H7A2.75 2.75 0 0 0 4.25 7.5v11A2.75 2.75 0 0 0 7 21.25h10a2.75 2.75 0 0 0 2.75-2.75v-11A2.75 2.75 0 0 0 17 4.75M7 6.25h1.25V7.5a.75.75 0 0 0 1.5 0V6.25h4.5V7.5a.75.75 0 0 0 1.5 0V6.25H17a1.25 1.25 0 0 1 1.25 1.25v2.75H5.75V7.5A1.25 1.25 0 0 1 7 6.25m10 13.5H7a1.25 1.25 0 0 1-1.25-1.25v-6.75h12.5v6.75A1.25 1.25 0 0 1 17 19.75"/><path fill="currentColor" d="M14 15.25h-1.25V14a.75.75 0 0 0-1.5 0v1.25H10a.75.75 0 0 0 0 1.5h1.25V18a.75.75 0 0 0 1.5 0v-1.25H14a.75.75 0 0 0 0-1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarTimes(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 4.75h-1.25V3.5a.75.75 0 0 0-1.5 0v1.25h-4.5V3.5a.75.75 0 0 0-1.5 0v1.25H7A2.75 2.75 0 0 0 4.25 7.5v11A2.75 2.75 0 0 0 7 21.25h10a2.75 2.75 0 0 0 2.75-2.75v-11A2.75 2.75 0 0 0 17 4.75M7 6.25h1.25V7.5a.75.75 0 0 0 1.5 0V6.25h4.5V7.5a.75.75 0 0 0 1.5 0V6.25H17a1.25 1.25 0 0 1 1.25 1.25v2.75H5.75V7.5A1.25 1.25 0 0 1 7 6.25m10 13.5H7a1.25 1.25 0 0 1-1.25-1.25v-6.75h12.5v6.75A1.25 1.25 0 0 1 17 19.75"/><path fill="currentColor" d="M13.94 14.06a.74.74 0 0 0-1.06 0l-.88.88l-.88-.88a.75.75 0 0 0-1.06 1.06l.88.88l-.88.88a.74.74 0 0 0 0 1.06a.71.71 0 0 0 .53.22a.74.74 0 0 0 .53-.22l.88-.88l.88.88a.74.74 0 0 0 .53.22a.71.71 0 0 0 .53-.22a.74.74 0 0 0 0-1.06l-.88-.88l.88-.88a.74.74 0 0 0 0-1.06"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Camera(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 19.25H6a2.75 2.75 0 0 1-2.75-2.75v-6A2.75 2.75 0 0 1 6 7.75h.6L7.78 6a2.76 2.76 0 0 1 2.29-1.22h3.86A2.76 2.76 0 0 1 16.22 6l1.18 1.75h.6a2.75 2.75 0 0 1 2.75 2.75v6A2.75 2.75 0 0 1 18 19.25m-12-10a1.25 1.25 0 0 0-1.25 1.25v6A1.25 1.25 0 0 0 6 17.75h12a1.25 1.25 0 0 0 1.25-1.25v-6A1.25 1.25 0 0 0 18 9.25h-1a.75.75 0 0 1-.62-.33L15 6.81a1.24 1.24 0 0 0-1-.56h-3.93a1.24 1.24 0 0 0-1 .56L7.62 8.92a.75.75 0 0 1-.62.33Z"/><path fill="currentColor" d="M12 16.25A3.25 3.25 0 1 1 15.25 13A3.26 3.26 0 0 1 12 16.25m0-5A1.75 1.75 0 1 0 13.75 13A1.76 1.76 0 0 0 12 11.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Car(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m19.78 9.44l-1.84-5a1.75 1.75 0 0 0-1.64-1.19H7.7A1.75 1.75 0 0 0 6.06 4.4l-1.84 5a1.76 1.76 0 0 0-1 1.56v4.5A1.73 1.73 0 0 0 4 16.94S4 17 4 17v2a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-1.75h10V19a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-2.06a1.73 1.73 0 0 0 .76-1.44V11a1.76 1.76 0 0 0-.98-1.56m-.53 6.06a.25.25 0 0 1-.25.25H5a.25.25 0 0 1-.25-.25V11a.25.25 0 0 1 .25-.25h14a.25.25 0 0 1 .25.25ZM7.47 4.91a.25.25 0 0 1 .23-.16h8.6a.25.25 0 0 1 .23.16l1.4 3.84H6.07Z"/><circle cx="8" cy="13.25" r="1.5" fill="currentColor"/><circle cx="16" cy="13.25" r="1.5" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretDown(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 16.75a.74.74 0 0 1-.6-.3l-6-8a.75.75 0 0 1-.07-.79a.76.76 0 0 1 .67-.41h12a.76.76 0 0 1 .67.41a.75.75 0 0 1-.07.79l-6 8a.74.74 0 0 1-.6.3m-4.5-8l4.5 6l4.5-6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretLeft(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 18.75a.74.74 0 0 1-.45-.15l-8-6a.75.75 0 0 1 0-1.2l8-6a.75.75 0 0 1 .79-.07a.76.76 0 0 1 .41.67v12a.76.76 0 0 1-.41.67a.84.84 0 0 1-.34.08M9.25 12l6 4.5v-9Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretRight(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 18.75a.76.76 0 0 1-.33-.08a.75.75 0 0 1-.42-.67V6a.75.75 0 0 1 .42-.67a.74.74 0 0 1 .78.07l8 6a.75.75 0 0 1 0 1.2l-8 6a.74.74 0 0 1-.45.15M8.75 7.5v9l6-4.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretUp(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 16.75H6a.76.76 0 0 1-.67-.41a.75.75 0 0 1 .07-.79l6-8a.77.77 0 0 1 1.2 0l6 8a.75.75 0 0 1 .07.79a.76.76 0 0 1-.67.41m-10.5-1.5h9l-4.5-6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CartPlus(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.39 11.5c0-.41.34-.75.75-.75h1.25V9.5c0-.41.34-.75.75-.75s.75.34.75.75v1.25h1.25c.41 0 .75.34.75.75s-.34.75-.75.75h-1.25v1.25c0 .41-.34.75-.75.75s-.75-.34-.75-.75v-1.25h-1.25c-.41 0-.75-.34-.75-.75m.86 7.25c0 .83-.67 1.5-1.5 1.5s-1.5-.67-1.5-1.5s.67-1.5 1.5-1.5s1.5.67 1.5 1.5m6.5 0c0 .83-.67 1.5-1.5 1.5s-1.5-.67-1.5-1.5s.67-1.5 1.5-1.5s1.5.67 1.5 1.5m2.98-11.07l-2 8a.75.75 0 0 1-.73.57H8c-.36 0-.67-.26-.74-.62L5.37 5.25H4c-.41 0-.75-.34-.75-.75s.34-.75.75-.75h2c.36 0 .67.26.74.62l.43 2.38H20a.754.754 0 0 1 .73.93m-1.69.57H7.44l1.18 6.5h8.79l1.62-6.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartBar(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 20.25a.76.76 0 0 1-.75-.75v-15a.75.75 0 0 1 1.5 0v15a.76.76 0 0 1-.75.75"/><path fill="currentColor" d="M19.5 20.25h-15a.75.75 0 0 1 0-1.5h15a.75.75 0 0 1 0 1.5M8 16.75a.76.76 0 0 1-.75-.75v-4a.75.75 0 0 1 1.5 0v4a.76.76 0 0 1-.75.75m3.5 0a.76.76 0 0 1-.75-.75V8a.75.75 0 0 1 1.5 0v8a.76.76 0 0 1-.75.75m3.5 0a.76.76 0 0 1-.75-.75v-4a.75.75 0 0 1 1.5 0v4a.76.76 0 0 1-.75.75m3.5 0a.76.76 0 0 1-.75-.75V8a.75.75 0 0 1 1.5 0v8a.76.76 0 0 1-.75.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartLine(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 20.25a.76.76 0 0 1-.75-.75v-15a.75.75 0 0 1 1.5 0v15a.76.76 0 0 1-.75.75"/><path fill="currentColor" d="M19.5 20.25h-15a.75.75 0 0 1 0-1.5h15a.75.75 0 0 1 0 1.5m-5.5-5.5a.74.74 0 0 1-.53-.22L11 12.06l-2.47 2.47a.75.75 0 0 1-1.06-1.06l3-3a.75.75 0 0 1 1.06 0L14 12.94l3.47-3.47a.75.75 0 0 1 1.06 1.06l-4 4a.74.74 0 0 1-.53.22"/><path fill="currentColor" d="M18.5 13.84a.76.76 0 0 1-.75-.75v-2.84H15a.75.75 0 0 1 0-1.5h3.5a.76.76 0 0 1 .75.75v3.59a.76.76 0 0 1-.75.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartPie(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.25 4.79V4.5a.76.76 0 0 0-.75-.75a8.8 8.8 0 0 0-7.61 13.13a.75.75 0 0 0 .65.37a.74.74 0 0 0 .37-.1L6.2 17a7.74 7.74 0 1 0 7.05-12.2Zm-1.5.5v6.78l-5.89 3.38a7.28 7.28 0 0 1 5.89-10.16m.75 13.46a6.27 6.27 0 0 1-5-2.51l5.37-3.09a.73.73 0 0 0 .38-.65V6.3a6.25 6.25 0 0 1-.75 12.45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Check(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 18.25a.74.74 0 0 1-.53-.25l-5-5a.75.75 0 1 1 1.06-1L9 16.44L19.47 6a.75.75 0 0 1 1.06 1l-11 11a.74.74 0 0 1-.53.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckCircle(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5 15.25A.74.74 0 0 1 10 15l-3-3a.75.75 0 0 1 1-1l2.47 2.47L19 5a.75.75 0 0 1 1 1l-9 9a.74.74 0 0 1-.5.25"/><path fill="currentColor" d="M12 21a9 9 0 0 1-7.87-4.66a8.67 8.67 0 0 1-1.07-3.41a9 9 0 0 1 4.6-8.81a8.67 8.67 0 0 1 3.41-1.07a8.86 8.86 0 0 1 3.55.34a.75.75 0 1 1-.43 1.43a7.62 7.62 0 0 0-3-.28a7.43 7.43 0 0 0-2.84.89a7.5 7.5 0 0 0-2.2 1.84a7.42 7.42 0 0 0-1.64 5.51a7.43 7.43 0 0 0 .89 2.84a7.5 7.5 0 0 0 1.84 2.2a7.42 7.42 0 0 0 5.51 1.64a7.43 7.43 0 0 0 2.84-.89a7.5 7.5 0 0 0 2.2-1.84a7.42 7.42 0 0 0 1.64-5.51a.75.75 0 1 1 1.57-.15a9 9 0 0 1-4.61 8.81A8.67 8.67 0 0 1 12.93 21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckSquare(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 20.75H6A2.75 2.75 0 0 1 3.25 18V6A2.75 2.75 0 0 1 6 3.25h8.86a.75.75 0 0 1 0 1.5H6A1.25 1.25 0 0 0 4.75 6v12A1.25 1.25 0 0 0 6 19.25h12A1.25 1.25 0 0 0 19.25 18v-7.71a.75.75 0 0 1 1.5 0V18A2.75 2.75 0 0 1 18 20.75"/><path fill="currentColor" d="M10.5 15.25A.74.74 0 0 1 10 15l-3-3a.75.75 0 0 1 1-1l2.47 2.47L19 5a.75.75 0 0 1 1 1l-9 9a.74.74 0 0 1-.5.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronCircleDown(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21a9 9 0 1 1 9-9a9 9 0 0 1-9 9m0-16.5a7.5 7.5 0 1 0 7.5 7.5A7.5 7.5 0 0 0 12 4.5"/><path fill="currentColor" d="M12 14.75a.74.74 0 0 1-.53-.22l-4-4a.75.75 0 0 1 1.06-1.06L12 12.94l3.47-3.47a.75.75 0 0 1 1.06 1.06l-4 4a.74.74 0 0 1-.53.22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronCircleLeft(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21a9 9 0 1 1 9-9a9 9 0 0 1-9 9m0-16.5a7.5 7.5 0 1 0 7.5 7.5A7.5 7.5 0 0 0 12 4.5"/><path fill="currentColor" d="M14 16.75a.74.74 0 0 1-.53-.22l-4-4a.75.75 0 0 1 0-1.06l4-4a.75.75 0 0 1 1.06 1.06L11.06 12l3.47 3.47a.75.75 0 0 1 0 1.06a.74.74 0 0 1-.53.22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronCircleRight(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21a9 9 0 1 1 9-9a9 9 0 0 1-9 9m0-16.5a7.5 7.5 0 1 0 7.5 7.5A7.5 7.5 0 0 0 12 4.5"/><path fill="currentColor" d="M10 16.75a.74.74 0 0 1-.53-.22a.75.75 0 0 1 0-1.06L12.94 12L9.47 8.53a.75.75 0 0 1 1.06-1.06l4 4a.75.75 0 0 1 0 1.06l-4 4a.74.74 0 0 1-.53.22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronCircleUp(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21a9 9 0 1 1 9-9a9 9 0 0 1-9 9m0-16.5a7.5 7.5 0 1 0 7.5 7.5A7.5 7.5 0 0 0 12 4.5"/><path fill="currentColor" d="M16 14.75a.74.74 0 0 1-.53-.22L12 11.06l-3.47 3.47a.75.75 0 0 1-1.06-1.06l4-4a.75.75 0 0 1 1.06 0l4 4a.75.75 0 0 1 0 1.06a.74.74 0 0 1-.53.22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDown(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 15.25a.74.74 0 0 1-.53-.22l-5-5A.75.75 0 0 1 7.53 9L12 13.44L16.47 9a.75.75 0 0 1 1.06 1l-5 5a.74.74 0 0 1-.53.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeft(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 17.75a.74.74 0 0 1-.53-.22l-5-5a.75.75 0 0 1 0-1.06l5-5a.75.75 0 0 1 1.06 1.06L10.06 12l4.47 4.47a.75.75 0 0 1 0 1.06a.74.74 0 0 1-.53.22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRight(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 17.75a.74.74 0 0 1-.53-.22a.75.75 0 0 1 0-1.06L13.94 12L9.47 7.53a.75.75 0 0 1 1.06-1.06l5 5a.75.75 0 0 1 0 1.06l-5 5a.74.74 0 0 1-.53.22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUp(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 15.25a.74.74 0 0 1-.53-.22L12 10.56L7.53 15a.75.75 0 0 1-1.06-1l5-5a.75.75 0 0 1 1.06 0l5 5a.75.75 0 0 1 0 1.06a.74.74 0 0 1-.53.19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Circle(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21a9 9 0 1 1 9-9a9 9 0 0 1-9 9m0-16.5a7.5 7.5 0 1 0 7.5 7.5A7.5 7.5 0 0 0 12 4.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleFill(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="12" cy="12" r="9" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleOff(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21a9 9 0 1 1 9-9a9 9 0 0 1-9 9m0-16.5a7.5 7.5 0 1 0 7.5 7.5A7.5 7.5 0 0 0 12 4.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleOn(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="12" cy="12" r="9" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clock(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21a9 9 0 1 1 9-9a9 9 0 0 1-9 9m0-16.5a7.5 7.5 0 1 0 7.5 7.5A7.5 7.5 0 0 0 12 4.5"/><path fill="currentColor" d="M15 12.75h-3a.76.76 0 0 1-.75-.75V7a.75.75 0 0 1 1.5 0v4.25H15a.75.75 0 0 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clone(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 16.75H6A2.75 2.75 0 0 1 3.25 14V6A2.75 2.75 0 0 1 6 3.25h8A2.75 2.75 0 0 1 16.75 6v8A2.75 2.75 0 0 1 14 16.75m-8-12A1.25 1.25 0 0 0 4.75 6v8A1.25 1.25 0 0 0 6 15.25h8A1.25 1.25 0 0 0 15.25 14V6A1.25 1.25 0 0 0 14 4.75Z"/><path fill="currentColor" d="M18 20.75h-8A2.75 2.75 0 0 1 7.25 18v-2h1.5v2A1.25 1.25 0 0 0 10 19.25h8A1.25 1.25 0 0 0 19.25 18v-8A1.25 1.25 0 0 0 18 8.75h-2v-1.5h2A2.75 2.75 0 0 1 20.75 10v8A2.75 2.75 0 0 1 18 20.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloud(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.5 18.75h-7a6.75 6.75 0 1 1 6.2-9.42a4.43 4.43 0 0 1 .8-.08A5.07 5.07 0 0 1 21.25 14a4.75 4.75 0 0 1-4.75 4.75m-7-12a5.25 5.25 0 0 0 0 10.5h7A3.26 3.26 0 0 0 19.75 14a3.57 3.57 0 0 0-3.25-3.25a3.09 3.09 0 0 0-1 .19a.78.78 0 0 1-.58 0a.73.73 0 0 1-.37-.44A5.24 5.24 0 0 0 9.5 6.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudDownload(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.5 18.25a.75.75 0 0 1 0-1.5c1.66 0 2.25-.83 2.25-3.18a3.57 3.57 0 0 0-3.25-3.25a3.34 3.34 0 0 0-1 .18a.74.74 0 0 1-1-.49a5.25 5.25 0 0 0-10.25 1.56c0 3.22.15 5.18 2.25 5.18a.75.75 0 0 1 0 1.5c-3.75 0-3.75-3.86-3.75-6.68a6.75 6.75 0 0 1 13-2.68a4.4 4.4 0 0 1 .8-.07a5.07 5.07 0 0 1 4.75 4.75c-.05 1.28-.05 4.68-3.8 4.68"/><path fill="currentColor" d="M12 19.18a.74.74 0 0 1-.53-.22l-2.83-2.83a.75.75 0 0 1 1.06-1.06l2.3 2.3l2.3-2.3a.75.75 0 0 1 1.06 1.06L12.53 19a.74.74 0 0 1-.53.18"/><path fill="currentColor" d="M12 19.18a.75.75 0 0 1-.75-.75v-6.36a.75.75 0 0 1 1.5 0v6.36a.75.75 0 0 1-.75.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudUpload(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.5 18.25a.75.75 0 0 1 0-1.5c1.66 0 2.25-.83 2.25-3.18a3.57 3.57 0 0 0-3.25-3.25a3.34 3.34 0 0 0-1 .18a.74.74 0 0 1-1-.49a5.25 5.25 0 0 0-10.25 1.56c0 3.44.76 5.18 2.25 5.18a.75.75 0 0 1 0 1.5c-2.5 0-3.75-2.25-3.75-6.68a6.75 6.75 0 0 1 13-2.68a4.4 4.4 0 0 1 .8-.07a5.07 5.07 0 0 1 4.75 4.75c-.05 1.28-.05 4.68-3.8 4.68"/><path fill="currentColor" d="M14.83 15.65a.77.77 0 0 1-.53-.22l-2.3-2.3l-2.3 2.3a.75.75 0 0 1-1.06-1.06l2.83-2.83a.74.74 0 0 1 1.06 0l2.83 2.83a.75.75 0 0 1 0 1.06a.79.79 0 0 1-.53.22"/><path fill="currentColor" d="M12 19.18a.75.75 0 0 1-.75-.75v-6.36a.75.75 0 0 1 1.5 0v6.36a.75.75 0 0 1-.75.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Code(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.53 6.47a.75.75 0 0 0-1.06 0l-5 5a.75.75 0 0 0 0 1.06l5 5a.75.75 0 0 0 1.06 0a.75.75 0 0 0 0-1.06L5.06 12l4.47-4.47a.75.75 0 0 0 0-1.06m11 5l-5-5a.75.75 0 0 0-1.06 1.06L18.94 12l-4.47 4.47a.75.75 0 0 0 0 1.06a.75.75 0 0 0 1.06 0l5-5a.75.75 0 0 0 0-1.06"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cog(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 14.93A2.93 2.93 0 1 1 14.93 12A2.93 2.93 0 0 1 12 14.93m0-4.36A1.43 1.43 0 1 0 13.43 12A1.43 1.43 0 0 0 12 10.57"/><path fill="currentColor" d="M12.06 20.75a2.22 2.22 0 0 1-2.21-2.24a.48.48 0 0 0-.29-.42a.51.51 0 0 0-.56.09a2.24 2.24 0 0 1-1.56.65a2.26 2.26 0 0 1-1.58-.63a2.22 2.22 0 0 1 0-3.14a.5.5 0 0 0 .14-.52a.56.56 0 0 0-.48-.28a2.27 2.27 0 0 1-2.26-2.2a2.22 2.22 0 0 1 2.23-2.21a.48.48 0 0 0 .42-.29a.51.51 0 0 0-.09-.56a2.24 2.24 0 0 1-.65-1.56a2.2 2.2 0 0 1 .63-1.58a2.22 2.22 0 0 1 3.14 0a.58.58 0 0 0 .55.14a.6.6 0 0 0 .31-.48A2.27 2.27 0 0 1 12 3.25a2.22 2.22 0 0 1 2.2 2.24a.52.52 0 0 0 .28.44a.5.5 0 0 0 .52-.11a2.24 2.24 0 0 1 1.56-.65a2.13 2.13 0 0 1 1.58.63a2.22 2.22 0 0 1 0 3.14a.58.58 0 0 0-.11.55a.6.6 0 0 0 .48.31a2.27 2.27 0 0 1 2.24 2.2a2.22 2.22 0 0 1-2.24 2.2a.52.52 0 0 0-.44.28a.5.5 0 0 0 .11.52a2.24 2.24 0 0 1 .65 1.56a2.2 2.2 0 0 1-.63 1.58a2.22 2.22 0 0 1-3.14 0a.5.5 0 0 0-.52-.11a.56.56 0 0 0-.28.48a2.27 2.27 0 0 1-2.2 2.24m-2.71-4.21a1.83 1.83 0 0 1 .77.16a2 2 0 0 1 1.23 1.8a.74.74 0 0 0 .71.75a.76.76 0 0 0 .7-.77a2.06 2.06 0 0 1 1.18-1.85a2 2 0 0 1 2.18.43a.7.7 0 1 0 1-1a2 2 0 0 1-.42-2.18a2 2 0 0 1 1.82-1.18a.73.73 0 0 0 .74-.7a.75.75 0 0 0-.77-.7a2.1 2.1 0 0 1-1.85-1.22a2.05 2.05 0 0 1 .43-2.2a.7.7 0 1 0-1-1a2 2 0 0 1-3.36-1.4a.73.73 0 0 0-.7-.74a.75.75 0 0 0-.7.77a2.1 2.1 0 0 1-1.22 1.85a2.05 2.05 0 0 1-2.2-.43a.7.7 0 1 0-1 1a2 2 0 0 1 .41 2.18a2 2 0 0 1-1.8 1.23a.74.74 0 0 0-.75.71a.76.76 0 0 0 .77.7a2.06 2.06 0 0 1 1.85 1.18a2 2 0 0 1-.43 2.18a.7.7 0 1 0 1 1a2.06 2.06 0 0 1 1.41-.57m8-2.36"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Comment(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 20.25a.75.75 0 0 1-.72-1L5.38 14a7.76 7.76 0 0 1-.52-2.83a8 8 0 0 1 .62-3.1a8.12 8.12 0 0 1 1.7-2.52a7.83 7.83 0 0 1 2.53-1.7a7.92 7.92 0 0 1 6.19 0a8 8 0 0 1 4.85 7.32a8 8 0 0 1-2.33 5.62a8.12 8.12 0 0 1-2.52 1.7a8 8 0 0 1-5.93.1l-5.25 1.6a.83.83 0 0 1-.22.06m8.3-15.5a6.49 6.49 0 0 0-5.94 3.94a6.55 6.55 0 0 0 0 5a.75.75 0 0 1 0 .51l-1.23 4.17l4.15-1.26a.75.75 0 0 1 .51 0a6.52 6.52 0 0 0 5 0a6.44 6.44 0 0 0 3.43-8.45a6.45 6.45 0 0 0-5.92-3.91"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Comments(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.25 17.31a.75.75 0 0 1-.72-1l1.25-4.12a6.64 6.64 0 0 1-.4-2.19a6.36 6.36 0 0 1 .51-2.51a6.48 6.48 0 0 1 10.53-2.1a6.66 6.66 0 0 1 1.38 2.06a6.46 6.46 0 0 1 0 5a6.66 6.66 0 0 1-1.38 2.06A6.48 6.48 0 0 1 8.59 16l-4.12 1.28a.83.83 0 0 1-.22.03M10.84 5a4.9 4.9 0 0 0-1.93.39A5 5 0 0 0 6.27 8a5 5 0 0 0 0 3.87a.75.75 0 0 1 0 .51l-.92 3l3-.92a.75.75 0 0 1 .51 0a5 5 0 0 0 6.51-2.64A5 5 0 0 0 10.84 5"/><path fill="currentColor" d="M19.75 20.5a.83.83 0 0 1-.22 0l-4.12-1.25a6.42 6.42 0 0 1-8.17-3.48a.73.73 0 0 1 .38-1a.75.75 0 0 1 1 .38a5.06 5.06 0 0 0 1 1.53a5 5 0 0 0 5.44 1.06a.75.75 0 0 1 .51 0l3 .92l-.92-3a.75.75 0 0 1 0-.51a4.9 4.9 0 0 0 .39-1.93a4.93 4.93 0 0 0-1.45-3.51a5.62 5.62 0 0 0-.67-.71a.75.75 0 1 1 .84-1.24a6.69 6.69 0 0 1 1 .79a6.49 6.49 0 0 1 1.38 2.06a6.38 6.38 0 0 1 .51 2.52a6.63 6.63 0 0 1-.4 2.25l1.25 4.12a.75.75 0 0 1-.72 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compass(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.94 7.62l-4.88 2a2.63 2.63 0 0 0-1.48 1.48l-2 4.88a.34.34 0 0 0 .19.44a.36.36 0 0 0 .25 0l4.88-2a2.63 2.63 0 0 0 1.48-1.48l2-4.88a.34.34 0 0 0-.19-.44a.36.36 0 0 0-.25 0M12 13a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/><path fill="currentColor" d="M12 21a9 9 0 1 1 9-9a9 9 0 0 1-9 9m0-16.5a7.5 7.5 0 1 0 7.5 7.5A7.5 7.5 0 0 0 12 4.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copy(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.53 8L14 2.47a.75.75 0 0 0-.53-.22H11A2.75 2.75 0 0 0 8.25 5v1.25H7A2.75 2.75 0 0 0 4.25 9v10A2.75 2.75 0 0 0 7 21.75h7A2.75 2.75 0 0 0 16.75 19v-1.25H17A2.75 2.75 0 0 0 19.75 15V8.5a.75.75 0 0 0-.22-.5m-5.28-3.19l2.94 2.94h-2.94Zm1 14.19A1.25 1.25 0 0 1 14 20.25H7A1.25 1.25 0 0 1 5.75 19V9A1.25 1.25 0 0 1 7 7.75h1.25V15A2.75 2.75 0 0 0 11 17.75h4.25ZM17 16.25h-6A1.25 1.25 0 0 1 9.75 15V5A1.25 1.25 0 0 1 11 3.75h1.75V8.5a.76.76 0 0 0 .75.75h4.75V15A1.25 1.25 0 0 1 17 16.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCard(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 5.25H5A1.76 1.76 0 0 0 3.25 7v10A1.76 1.76 0 0 0 5 18.75h14A1.76 1.76 0 0 0 20.75 17V7A1.76 1.76 0 0 0 19 5.25M5 6.75h14a.25.25 0 0 1 .25.25v2.25H4.75V7A.25.25 0 0 1 5 6.75m14 10.5H5a.25.25 0 0 1-.25-.25v-6.25h14.5V17a.25.25 0 0 1-.25.25"/><path fill="currentColor" d="M9 13H7a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Database(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.48 4.54A19.17 19.17 0 0 0 12 2.75a19.17 19.17 0 0 0-7.48 1.79a1.74 1.74 0 0 0-1 1.59v11.74a1.74 1.74 0 0 0 1 1.59A19.17 19.17 0 0 0 12 21.25a19.17 19.17 0 0 0 7.48-1.79a1.74 1.74 0 0 0 1-1.59V6.13a1.74 1.74 0 0 0-1-1.59m-.48 8a17.76 17.76 0 0 1-7 1.72a17.76 17.76 0 0 1-7-1.72V8.68a18.64 18.64 0 0 0 7 1.57a18.64 18.64 0 0 0 7-1.57ZM5.15 5.9A17.56 17.56 0 0 1 12 4.25a17.67 17.67 0 0 1 6.86 1.65a.26.26 0 0 1 .14.23V7a17.76 17.76 0 0 1-7 1.72A17.76 17.76 0 0 1 5 7v-.9a.25.25 0 0 1 .15-.2m13.7 12.2A17.56 17.56 0 0 1 12 19.75a17.67 17.67 0 0 1-6.86-1.65a.26.26 0 0 1-.14-.23v-3.69a18.64 18.64 0 0 0 7 1.57a18.64 18.64 0 0 0 7-1.57v3.69a.25.25 0 0 1-.15.23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeleteLeft(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.23 9.78L15.01 12l2.22 2.22c.29.29.29.77 0 1.06c-.15.15-.34.22-.53.22s-.38-.07-.53-.22l-2.22-2.22l-2.22 2.22c-.15.15-.34.22-.53.22s-.38-.07-.53-.22a.754.754 0 0 1 0-1.06L12.89 12l-2.22-2.22c-.29-.29-.29-.77 0-1.06s.77-.29 1.06 0l2.22 2.22l2.22-2.22c.29-.29.77-.29 1.06 0s.29.77 0 1.06M21.32 7v10c0 .96-.78 1.75-1.75 1.75H7.64c-.61 0-1.16-.31-1.48-.82l-3.29-5.27c-.25-.4-.25-.92 0-1.33l3.29-5.26c.32-.51.88-.82 1.48-.82h11.94c.96 0 1.75.79 1.75 1.75Zm-1.5 0c0-.14-.11-.25-.25-.25H7.64c-.09 0-.17.04-.21.12L4.22 12l3.21 5.13c.05.07.13.12.21.12h11.94c.14 0 .25-.11.25-.25V7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Desktop(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 16.75H5A1.76 1.76 0 0 1 3.25 15V5A1.76 1.76 0 0 1 5 3.25h14A1.76 1.76 0 0 1 20.75 5v10A1.76 1.76 0 0 1 19 16.75m-14-12a.25.25 0 0 0-.25.25v10a.25.25 0 0 0 .25.25h14a.25.25 0 0 0 .25-.25V5a.25.25 0 0 0-.25-.25Z"/><path fill="currentColor" d="M15 20.75h-3a.76.76 0 0 1-.75-.75v-4a.75.75 0 0 1 1.5 0v3.25H15a.75.75 0 0 1 0 1.5"/><path fill="currentColor" d="M12 20.75H9a.75.75 0 0 1 0-1.5h3a.75.75 0 0 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Directions(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21.25a1.81 1.81 0 0 1-1.28-.53l-7.44-7.44a1.81 1.81 0 0 1 0-2.56l7.44-7.44a1.86 1.86 0 0 1 2.56 0l7.44 7.44a1.81 1.81 0 0 1 0 2.56l-7.44 7.44a1.81 1.81 0 0 1-1.28.53m0-17a.27.27 0 0 0-.21.09l-7.45 7.45a.29.29 0 0 0 0 .42l7.45 7.45a.25.25 0 0 0 .42 0l7.45-7.45a.29.29 0 0 0 0-.42l-7.45-7.45a.27.27 0 0 0-.21-.09"/><path fill="currentColor" d="M13.27 14.42a.74.74 0 0 1-.53-.22a.75.75 0 0 1 0-1.06l1.7-1.68l-1.7-1.68a.75.75 0 1 1 1.06-1.06l2.2 2.21A.75.75 0 0 1 16 12l-2.2 2.2a.74.74 0 0 1-.53.22"/><path fill="currentColor" d="M8.5 15a.76.76 0 0 1-.75-.75v-2.79a.76.76 0 0 1 .75-.75h7a.75.75 0 0 1 0 1.5H9.25v2a.76.76 0 0 1-.75.79"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DirectionsAlt(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21.25a1.81 1.81 0 0 1-1.28-.53l-7.44-7.44a1.81 1.81 0 0 1 0-2.56l7.44-7.44a1.86 1.86 0 0 1 2.56 0l7.44 7.44a1.81 1.81 0 0 1 0 2.56l-7.44 7.44a1.81 1.81 0 0 1-1.28.53m0-17a.27.27 0 0 0-.21.09l-7.45 7.45a.29.29 0 0 0 0 .42l7.45 7.45a.25.25 0 0 0 .42 0l7.45-7.45a.29.29 0 0 0 0-.42l-7.45-7.45a.27.27 0 0 0-.21-.09"/><path fill="currentColor" d="M10.73 14.42a.74.74 0 0 1-.53-.22L8 12a.75.75 0 0 1 0-1.06l2.2-2.22a.75.75 0 0 1 1.06 0a.75.75 0 0 1 0 1.06l-1.7 1.68l1.7 1.68a.75.75 0 0 1 0 1.06a.74.74 0 0 1-.53.22"/><path fill="currentColor" d="M15.5 15a.76.76 0 0 1-.75-.75v-2H8.5a.75.75 0 0 1 0-1.5h7a.76.76 0 0 1 .75.75v2.79a.76.76 0 0 1-.75.71"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Discord(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.59 5.89c-1.23-.57-2.54-.99-3.92-1.23c-.17.3-.37.71-.5 1.04c-1.46-.22-2.91-.22-4.34 0c-.14-.33-.34-.74-.51-1.04c-1.38.24-2.69.66-3.92 1.23c-2.48 3.74-3.15 7.39-2.82 10.98c1.65 1.23 3.24 1.97 4.81 2.46c.39-.53.73-1.1 1.03-1.69c-.57-.21-1.11-.48-1.62-.79c.14-.1.27-.21.4-.31c3.13 1.46 6.52 1.46 9.61 0c.13.11.26.21.4.31c-.51.31-1.06.57-1.62.79c.3.59.64 1.16 1.03 1.69c1.57-.49 3.17-1.23 4.81-2.46c.39-4.17-.67-7.78-2.82-10.98Zm-9.75 8.78c-.94 0-1.71-.87-1.71-1.94s.75-1.94 1.71-1.94s1.72.87 1.71 1.94c0 1.06-.75 1.94-1.71 1.94m6.31 0c-.94 0-1.71-.87-1.71-1.94s.75-1.94 1.71-1.94s1.72.87 1.71 1.94c0 1.06-.75 1.94-1.71 1.94"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dollar(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 20.75a.76.76 0 0 1-.75-.75V4a.75.75 0 0 1 1.5 0v16a.76.76 0 0 1-.75.75"/><path fill="currentColor" d="M13.5 18.75H7a.75.75 0 0 1 0-1.5h6.5A2.54 2.54 0 0 0 16.25 15a2.54 2.54 0 0 0-2.75-2.25h-3A4 4 0 0 1 6.25 9a4 4 0 0 1 4.25-3.75H16a.75.75 0 0 1 0 1.5h-5.5A2.54 2.54 0 0 0 7.75 9a2.54 2.54 0 0 0 2.75 2.25h3A4 4 0 0 1 17.75 15a4 4 0 0 1-4.25 3.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Download(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.22 20.75H5.78A2.64 2.64 0 0 1 3.25 18v-3a.75.75 0 0 1 1.5 0v3a1.16 1.16 0 0 0 1 1.25h12.47a1.16 1.16 0 0 0 1-1.25v-3a.75.75 0 0 1 1.5 0v3a2.64 2.64 0 0 1-2.5 2.75"/><path fill="currentColor" d="M12 15.75a.74.74 0 0 1-.53-.22l-4-4a.75.75 0 0 1 1.06-1.06L12 13.94l3.47-3.47a.75.75 0 0 1 1.06 1.06l-4 4a.74.74 0 0 1-.53.22"/><path fill="currentColor" d="M12 15.75a.76.76 0 0 1-.75-.75V4a.75.75 0 0 1 1.5 0v11a.76.76 0 0 1-.75.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eject(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 12.75H5a.75.75 0 0 1-.53-1.28l7-7a.75.75 0 0 1 1.06 0l7 7a.75.75 0 0 1-.53 1.28m-12.19-1.5h10.38L12 6.06ZM18 19.75H6A1.76 1.76 0 0 1 4.25 18v-2A1.76 1.76 0 0 1 6 14.25h12A1.76 1.76 0 0 1 19.75 16v2A1.76 1.76 0 0 1 18 19.75m-12-4a.25.25 0 0 0-.25.25v2a.25.25 0 0 0 .25.25h12a.25.25 0 0 0 .25-.25v-2a.25.25 0 0 0-.25-.25Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EllipsisH(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="12" cy="12" r="1.75" fill="currentColor"/><circle cx="19" cy="12" r="1.75" fill="currentColor"/><circle cx="5" cy="12" r="1.75" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EllipsisV(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="12" cy="12" r="1.75" fill="currentColor"/><circle cx="12" cy="5" r="1.75" fill="currentColor"/><circle cx="12" cy="19" r="1.75" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Envelope(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 4.25H5A1.76 1.76 0 0 0 3.25 6v12A1.76 1.76 0 0 0 5 19.75h14A1.76 1.76 0 0 0 20.75 18V6A1.76 1.76 0 0 0 19 4.25M5 5.75h14a.25.25 0 0 1 .25.25v1.54L12 11.16L4.75 7.54V6A.25.25 0 0 1 5 5.75m14 12.5H5a.25.25 0 0 1-.25-.25V9.21l6.91 3.46a.76.76 0 0 0 .68 0l6.91-3.46V18a.25.25 0 0 1-.25.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eraser(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.25 9.84c0-.46-.18-.9-.51-1.23l-4.35-4.35c-.68-.68-1.78-.68-2.46 0l-8.67 8.67c-.33.33-.51.77-.51 1.23s.18.9.51 1.23l4.35 4.35c.34.34.78.51 1.23.51c.04 0 .09 0 .13-.01h8.04c.41 0 .75-.34.75-.75s-.34-.75-.75-.75h-5.95l7.67-7.67c.33-.33.51-.77.51-1.23ZM5.32 14.33c-.06-.06-.07-.13-.07-.17s0-.11.07-.17l3.8-3.8l4.69 4.69l-3.8 3.8c-.09.09-.24.09-.34 0l-4.35-4.36Zm13.36-4.31l-3.8 3.8l-4.69-4.69l3.8-3.8c.09-.09.24-.09.34 0l4.35 4.36c.06.06.07.13.07.17s0 .11-.07.17Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Euro(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 20.75h-.15a8.75 8.75 0 0 1-5.93-15a8.54 8.54 0 0 1 6.23-2.46a8.75 8.75 0 0 1 6 2.5a.75.75 0 0 1 0 1.06a.75.75 0 0 1-1.06 0a7.26 7.26 0 1 0-.19 10.53l.22-.21a.79.79 0 0 1 1.09 0a.7.7 0 0 1 .05 1l-.05.05l-.29.28A8.72 8.72 0 0 1 13 20.75"/><path fill="currentColor" d="M17 11.25H3a.75.75 0 0 1 0-1.5h14a.75.75 0 0 1 0 1.5m-1.5 3H3a.75.75 0 0 1 0-1.5h12.5a.75.75 0 0 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationCircle(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21a9 9 0 1 1 9-9a9 9 0 0 1-9 9m0-16.5a7.5 7.5 0 1 0 7.5 7.5A7.5 7.5 0 0 0 12 4.5"/><path fill="currentColor" d="M12 13a.76.76 0 0 1-.75-.75v-3.5a.75.75 0 0 1 1.5 0v3.5A.76.76 0 0 1 12 13m0 3a.76.76 0 0 1-.75-.75v-.5a.75.75 0 0 1 1.5 0v.5A.76.76 0 0 1 12 16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationTriangle(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 18.75H4a.76.76 0 0 1-.65-.37a.77.77 0 0 1 0-.75l8-14a.78.78 0 0 1 1.3 0l8 14a.77.77 0 0 1 0 .75a.76.76 0 0 1-.65.37m-14.71-1.5h13.42L12 5.51Z"/><path fill="currentColor" d="M12 13.25a.76.76 0 0 1-.75-.75V9a.75.75 0 0 1 1.5 0v3.5a.76.76 0 0 1-.75.75m0 3a.76.76 0 0 1-.75-.75V15a.75.75 0 0 1 1.5 0v.5a.76.76 0 0 1-.75.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExternalLink(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 20.75H6A2.75 2.75 0 0 1 3.25 18V6A2.75 2.75 0 0 1 6 3.25h6a.75.75 0 0 1 0 1.5H6A1.25 1.25 0 0 0 4.75 6v12A1.25 1.25 0 0 0 6 19.25h12A1.25 1.25 0 0 0 19.25 18v-6a.75.75 0 0 1 1.5 0v6A2.75 2.75 0 0 1 18 20.75m2-12a.76.76 0 0 1-.75-.75V4.75H16a.75.75 0 0 1 0-1.5h4a.76.76 0 0 1 .75.75v4a.76.76 0 0 1-.75.75"/><path fill="currentColor" d="M13.5 11.25A.74.74 0 0 1 13 11a.75.75 0 0 1 0-1l6.5-6.5a.75.75 0 1 1 1.06 1.06L14 11a.74.74 0 0 1-.5.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eye(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 18.75c-5.83 0-8.57-6.19-8.69-6.45a.78.78 0 0 1 0-.6c.12-.26 2.86-6.45 8.69-6.45s8.57 6.19 8.69 6.45a.78.78 0 0 1 0 .6c-.12.26-2.86 6.45-8.69 6.45M4.83 12c.59 1.15 3 5.25 7.17 5.25s6.58-4.1 7.17-5.25c-.59-1.15-3-5.25-7.17-5.25S5.42 10.85 4.83 12"/><path fill="currentColor" d="M12 15.25A3.25 3.25 0 1 1 15.25 12A3.26 3.26 0 0 1 12 15.25m0-5A1.75 1.75 0 1 0 13.75 12A1.76 1.76 0 0 0 12 10.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeSlash(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.69 11.7c-.12-.26-2.86-6.45-8.69-6.45a7.67 7.67 0 0 0-1.66.18a.75.75 0 0 0 .32 1.46A6.62 6.62 0 0 1 12 6.75c4.18 0 6.58 4.1 7.17 5.25a13.28 13.28 0 0 1-1.26 2a.75.75 0 0 0 .59 1.21a.75.75 0 0 0 .59-.29a13.42 13.42 0 0 0 1.6-2.59a.78.78 0 0 0 0-.63M6.53 5.47a.75.75 0 0 0-1.06 1.06l.92.92a13.16 13.16 0 0 0-3.08 4.26a.76.76 0 0 0 0 .59c.12.26 2.86 6.45 8.69 6.45a7.93 7.93 0 0 0 4.39-1.3l1.08 1.08a.75.75 0 0 0 1.06 0a.75.75 0 0 0 0-1.06Zm3.83 6l2.21 2.22a1.81 1.81 0 0 1-1.81-.4a1.74 1.74 0 0 1-.4-1.87ZM12 17.25c-4.19 0-6.58-4.11-7.17-5.25a12 12 0 0 1 2.62-3.49l1.79 1.79a3.24 3.24 0 0 0 4.46 4.46l1.61 1.61a6.5 6.5 0 0 1-3.31.88"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Facebook(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 12.05a8 8 0 1 0-9.25 8v-5.67h-2v-2.33h2v-1.77a2.83 2.83 0 0 1 3-3.14a11.92 11.92 0 0 1 1.79.16v2h-1a1.16 1.16 0 0 0-1.3 1.26v1.51h2.22l-.36 2.33h-1.85V20A8 8 0 0 0 20 12.05"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastBackward(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.29 4.31a.757.757 0 0 0-.82.16l-6.72 6.72V5c0-.3-.18-.58-.46-.69a.757.757 0 0 0-.82.16l-6.72 6.72V5c0-.41-.34-.75-.75-.75s-.75.34-.75.75v14c0 .41.34.75.75.75s.75-.34.75-.75v-6.19l6.72 6.72a.75.75 0 0 0 1.28-.53v-6.19l6.72 6.72a.75.75 0 0 0 1.28-.53V5c0-.3-.18-.58-.46-.69m-9.04 12.88L6.06 12l5.19-5.19zm8 0L14.06 12l5.19-5.19z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastForward(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 4.25c-.41 0-.75.34-.75.75v6.19l-6.72-6.72a.751.751 0 0 0-1.28.53v6.19L4.53 4.47A.751.751 0 0 0 3.25 5v14c0 .3.18.58.46.69a.75.75 0 0 0 .82-.16l6.72-6.72V19c0 .3.18.58.46.69a.75.75 0 0 0 .82-.16l6.72-6.72V19c0 .41.34.75.75.75s.75-.34.75-.75V5c0-.41-.34-.75-.75-.75M4.75 17.19V6.81L9.94 12zm8 0V6.81L17.94 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func File(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.53 9L13 3.47a.75.75 0 0 0-.53-.22H8A2.75 2.75 0 0 0 5.25 6v12A2.75 2.75 0 0 0 8 20.75h8A2.75 2.75 0 0 0 18.75 18V9.5a.75.75 0 0 0-.22-.5m-5.28-3.19l2.94 2.94h-2.94ZM16 19.25H8A1.25 1.25 0 0 1 6.75 18V6A1.25 1.25 0 0 1 8 4.75h3.75V9.5a.76.76 0 0 0 .75.75h4.75V18A1.25 1.25 0 0 1 16 19.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileEdit(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 18.75H6.5c-.69 0-1.25-.56-1.25-1.25v-12c0-.69.56-1.25 1.25-1.25h3.75V9c0 .41.34.75.75.75h4.8c.1.29.37.5.7.5c.41 0 .75-.34.75-.75V9a.776.776 0 0 0-.22-.53l-5.5-5.5a.776.776 0 0 0-.53-.22H6.5c-1.52 0-2.75 1.23-2.75 2.75v12c0 1.52 1.23 2.75 2.75 2.75H8c.41 0 .75-.34.75-.75s-.34-.75-.75-.75m3.75-13.44l2.94 2.94h-2.94zm7.86 6.06c-.38-.38-.94-.61-1.52-.62c-.6-.03-1.17.2-1.55.59l-6.39 6.4c-.13.13-.2.29-.22.47l-.18 2.23c-.02.22.06.44.22.59c.14.14.33.22.53.22h.07l2.25-.21a.74.74 0 0 0 .46-.22l6.39-6.4c.8-.79.77-2.22-.06-3.05m-1 1.99l-6.2 6.21l-1.09.1l.08-1.06l6.2-6.21c.1-.1.28-.14.46-.15c.2 0 .38.07.49.18c.24.23.27.72.06.93"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileExcel(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.53 9L13 3.47a.75.75 0 0 0-.53-.22H8A2.75 2.75 0 0 0 5.25 6v12A2.75 2.75 0 0 0 8 20.75h8A2.75 2.75 0 0 0 18.75 18V9.5a.75.75 0 0 0-.22-.5m-5.28-3.19l2.94 2.94h-2.94ZM16 19.25H8A1.25 1.25 0 0 1 6.75 18V6A1.25 1.25 0 0 1 8 4.75h3.75V9.5a.76.76 0 0 0 .75.75h4.75V18A1.25 1.25 0 0 1 16 19.25"/><path fill="currentColor" d="M14.47 11.91a.77.77 0 0 0-1.06.12L12 13.8L10.59 12a.75.75 0 1 0-1.18 1L11 15l-1.59 2a.75.75 0 0 0 1.18.94L12 16.2l1.41 1.8a.77.77 0 0 0 .59.28a.75.75 0 0 0 .59-1.28L13 15l1.63-2a.77.77 0 0 0-.16-1.09"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileExport(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 15.75c-.41 0-.75.34-.75.75V18c0 .69-.56 1.25-1.25 1.25h-7.5c-.69 0-1.25-.56-1.25-1.25V6c0-.69.56-1.25 1.25-1.25H9V9.5c0 .41.34.75.75.75h4.75v1.25c0 .41.34.75.75.75s.75-.34.75-.75v-2c0-.2-.08-.39-.22-.53l-5.5-5.5a.75.75 0 0 0-.53-.22h-4C4.23 3.25 3 4.48 3 6v12c0 1.52 1.23 2.75 2.75 2.75h7.5c1.52 0 2.75-1.23 2.75-2.75v-1.5c0-.41-.34-.75-.75-.75M10.5 5.81l2.94 2.94H10.5zm10.44 8.48c-.04.09-.09.17-.16.24l-3 3c-.15.15-.34.22-.53.22s-.38-.07-.53-.22a.754.754 0 0 1 0-1.06l1.72-1.72H9.75c-.41 0-.75-.34-.75-.75s.34-.75.75-.75h8.69l-1.72-1.72c-.29-.29-.29-.77 0-1.06s.77-.29 1.06 0l3 3c.07.07.12.15.16.24c.08.18.08.39 0 .57Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileImport(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.44 14.75H3.75c-.41 0-.75-.34-.75-.75s.34-.75.75-.75h8.69l-1.72-1.72c-.29-.29-.29-.77 0-1.06s.77-.29 1.06 0l3 3c.07.07.12.15.16.24c.08.18.08.39 0 .57c-.04.09-.09.17-.16.24l-3 3c-.15.15-.34.22-.53.22s-.38-.07-.53-.22a.754.754 0 0 1 0-1.06l1.72-1.72ZM21 9.5V18c0 1.52-1.23 2.75-2.75 2.75h-7.5C9.23 20.75 8 19.52 8 18v-1c0-.41.34-.75.75-.75s.75.34.75.75v1c0 .69.56 1.25 1.25 1.25h7.5c.69 0 1.25-.56 1.25-1.25v-7.75h-4.75c-.41 0-.75-.34-.75-.75V4.75h-3.25c-.69 0-1.25.56-1.25 1.25v5c0 .41-.34.75-.75.75S8 11.41 8 11V6c0-1.52 1.23-2.75 2.75-2.75h4c.2 0 .39.08.53.22l5.5 5.5c.14.14.22.33.22.53m-5.5-.75h2.94L15.5 5.81z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileO(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.53 9L13 3.47a.75.75 0 0 0-.53-.22H8A2.75 2.75 0 0 0 5.25 6v12A2.75 2.75 0 0 0 8 20.75h8A2.75 2.75 0 0 0 18.75 18V9.5a.75.75 0 0 0-.22-.5m-5.28-3.19l2.94 2.94h-2.94ZM16 19.25H8A1.25 1.25 0 0 1 6.75 18V6A1.25 1.25 0 0 1 8 4.75h3.75V9.5a.76.76 0 0 0 .75.75h4.75V18A1.25 1.25 0 0 1 16 19.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilePdf(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.53 9L13 3.47a.75.75 0 0 0-.53-.22H8A2.75 2.75 0 0 0 5.25 6v12A2.75 2.75 0 0 0 8 20.75h8A2.75 2.75 0 0 0 18.75 18V9.5a.75.75 0 0 0-.22-.5m-5.28-3.19l2.94 2.94h-2.94ZM16 19.25H8A1.25 1.25 0 0 1 6.75 18V6A1.25 1.25 0 0 1 8 4.75h3.75V9.5a.76.76 0 0 0 .75.75h4.75V18A1.25 1.25 0 0 1 16 19.25"/><path fill="currentColor" d="M13.49 14.85a3.15 3.15 0 0 1-1.31-1.66a4.44 4.44 0 0 0 .19-2a.8.8 0 0 0-1.52-.19a5 5 0 0 0 .25 2.4A29 29 0 0 1 9.83 16c-.71.4-1.68 1-1.83 1.69c-.12.56.93 2 2.72-1.12a18.58 18.58 0 0 1 2.44-.72a4.72 4.72 0 0 0 2 .61a.82.82 0 0 0 .62-1.38c-.42-.43-1.67-.31-2.29-.23m-4.78 3a4.32 4.32 0 0 1 1.09-1.24c-.68 1.08-1.09 1.27-1.09 1.25Zm2.92-6.81c.26 0 .24 1.15.06 1.46a3.07 3.07 0 0 1-.06-1.45Zm-.87 4.88a14.76 14.76 0 0 0 .88-1.92a3.88 3.88 0 0 0 1.08 1.26a12.35 12.35 0 0 0-1.96.67Zm4.7-.18s-.18.22-1.33-.28c1.25-.08 1.46.21 1.33.29Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileWord(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m18.53 8.97l-5.5-5.5a.75.75 0 0 0-.53-.22H8C6.48 3.25 5.25 4.48 5.25 6v12c0 1.52 1.23 2.75 2.75 2.75h8c1.52 0 2.75-1.23 2.75-2.75V9.5c0-.2-.08-.39-.22-.53m-5.28-3.16l2.94 2.94h-2.94zM16 19.25H8c-.69 0-1.25-.56-1.25-1.25V6c0-.69.56-1.25 1.25-1.25h3.75V9.5c0 .41.34.75.75.75h4.75V18c0 .69-.56 1.25-1.25 1.25m.22-6.53l-1.5 5c-.09.3-.37.52-.69.53H14c-.31 0-.58-.19-.7-.47L12 14.52l-1.3 3.26c-.12.3-.41.49-.73.47a.761.761 0 0 1-.69-.53l-1.5-5a.748.748 0 0 1 1.43-.43l.88 2.94l1.2-3.01c.23-.57 1.17-.57 1.39 0l1.2 3.01l.88-2.94c.12-.4.54-.62.93-.5c.4.12.62.54.5.93Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Filter(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 20.5h-4a.76.76 0 0 1-.75-.75V12L3.9 4.69a.73.73 0 0 1-.07-.78a.76.76 0 0 1 .67-.41h15a.76.76 0 0 1 .67.41a.73.73 0 0 1-.07.78L14.75 12v7.75a.76.76 0 0 1-.75.75M10.75 19h2.5v-7.25a.71.71 0 0 1 .15-.44L18 5H6l4.62 6.31a.71.71 0 0 1 .15.44Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterFill(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.17 3.91a.76.76 0 0 0-.67-.41h-15a.76.76 0 0 0-.67.41a.73.73 0 0 0 .07.78L9.25 12v7.75a.76.76 0 0 0 .75.75h4a.76.76 0 0 0 .75-.75V12l5.35-7.31a.73.73 0 0 0 .07-.78"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterSlash(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.17 3.91a.76.76 0 0 0-.67-.41h-7a.75.75 0 0 0 0 1.5H18l-1.85 2.56a.74.74 0 1 0 1.2.88l2.75-3.75a.73.73 0 0 0 .07-.78M8.65 3.72H8.6a.76.76 0 0 0-.19-.12a.79.79 0 0 0-.22-.05H4.5a.76.76 0 0 0-.67.41a.73.73 0 0 0 .07.78L9.25 12v7.75a.76.76 0 0 0 .75.75h4a.76.76 0 0 0 .75-.75V12l.06-.09l3.66 3.62a.74.74 0 0 0 .53.22a.73.73 0 0 0 .53-.22a.75.75 0 0 0 0-1.06Zm4.75 7.59a.71.71 0 0 0-.15.44V19h-2.5v-7.25a.71.71 0 0 0-.15-.44L6 5h1.82l5.91 5.85Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flag(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.42 4.45a.74.74 0 0 0-.69-.08a13.82 13.82 0 0 1-3.73 1a8.46 8.46 0 0 1-2.3-1a8.84 8.84 0 0 0-3-1.16a20 20 0 0 0-5 1.1a.75.75 0 0 0-.51.71V20a.75.75 0 0 0 1.5 0v-5.85a15.79 15.79 0 0 1 3.86-.87a8.54 8.54 0 0 1 2.4 1a9.11 9.11 0 0 0 2.82 1.13H15a16 16 0 0 0 4.21-1.13a.75.75 0 0 0 .48-.7V5.07a.74.74 0 0 0-.27-.62m-1.17 8.63a12.06 12.06 0 0 1-3.25.84a8.46 8.46 0 0 1-2.3-1a8.84 8.84 0 0 0-3-1.16h-.2a16.31 16.31 0 0 0-3.78.8v-7a15.79 15.79 0 0 1 3.86-.87a8.54 8.54 0 0 1 2.4 1a9.11 9.11 0 0 0 2.82 1.13a10.27 10.27 0 0 0 3.42-.73Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlagFill(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.42 4.45a.74.74 0 0 0-.69-.08a13.18 13.18 0 0 1-3.73 1a8.46 8.46 0 0 1-2.3-1a8.76 8.76 0 0 0-3-1.16c-1.29-.12-4.36.89-5 1.1a.75.75 0 0 0-.51.71V20a.75.75 0 0 0 1.5 0v-5.86a16 16 0 0 1 3.86-.85a8.47 8.47 0 0 1 2.4 1a9.11 9.11 0 0 0 2.82 1.13H15a16.37 16.37 0 0 0 4.21-1.13a.76.76 0 0 0 .48-.7V5.07a.74.74 0 0 0-.27-.62"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folder(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.5 19.25h-11a2.75 2.75 0 0 1-2.75-2.75v-9A2.75 2.75 0 0 1 6.5 4.75H9a.77.77 0 0 1 .57.25l2.77 3.24h5.16A2.75 2.75 0 0 1 20.25 11v5.5a2.75 2.75 0 0 1-2.75 2.75m-11-13A1.25 1.25 0 0 0 5.25 7.5v9a1.25 1.25 0 0 0 1.25 1.25h11a1.25 1.25 0 0 0 1.25-1.25V11a1.25 1.25 0 0 0-1.25-1.25H12a.77.77 0 0 1-.57-.26L8.66 6.25Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOpen(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.25 18.5h-1.5v-11a2.71 2.71 0 0 1 2.68-2.75h2.41a.76.76 0 0 1 .58.25l2.67 3.23H16A2.71 2.71 0 0 1 18.72 11v.5h-1.5V11A1.22 1.22 0 0 0 16 9.75h-5.27a.74.74 0 0 1-.57-.27L7.49 6.25H5.43A1.22 1.22 0 0 0 4.25 7.5Z"/><path fill="currentColor" d="M17.12 19.25H3.5a.76.76 0 0 1-.64-.36a.75.75 0 0 1 0-.74l3.38-6.5a.77.77 0 0 1 .67-.4H20.5a.76.76 0 0 1 .64.36a.75.75 0 0 1 0 .74l-3.38 6.5a.77.77 0 0 1-.64.4m-12.39-1.5h11.94l2.6-5H7.33Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Forward(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.51 11.45l-7.47-7a.756.756 0 0 0-.81-.14c-.27.12-.45.39-.45.69v6.27L4.51 4.45a.756.756 0 0 0-.81-.14c-.27.12-.45.39-.45.69v14c0 .3.18.57.45.69c.1.04.2.06.3.06c.19 0 .37-.07.51-.2l7.27-6.82V19c0 .3.18.57.45.69c.1.04.2.06.3.06c.19 0 .37-.07.51-.2l7.47-7c.15-.14.24-.34.24-.55s-.09-.41-.24-.55M4.75 17.27V6.73L10.37 12zm8.53 0V6.73L18.9 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gift(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.5 7.75h-1.4c.4-.48.65-1.08.65-1.75c0-1.52-1.23-2.75-2.75-2.75c-1.68 0-3.16.89-4 2.21a4.75 4.75 0 0 0-4-2.21C6.48 3.25 5.25 4.48 5.25 6c0 .67.25 1.27.65 1.75H4.5c-.69 0-1.25.56-1.25 1.25v2.5c0 .6.43 1.08 1 1.2v6.8c0 .69.56 1.25 1.25 1.25h13c.69 0 1.25-.56 1.25-1.25v-6.8c.57-.12 1-.6 1-1.2V9c0-.69-.56-1.25-1.25-1.25m-.25 3.5h-6.5v-2h6.5zM16 4.75a1.25 1.25 0 0 1 0 2.5h-3.16c.34-1.43 1.63-2.5 3.16-2.5m-8 0c1.53 0 2.82 1.07 3.16 2.5H8a1.25 1.25 0 0 1 0-2.5m-3.25 4.5h6.5v2h-6.5zm1 3.5h5.5v6.5h-5.5zm12.5 6.5h-5.5v-6.5h5.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Github(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.35 16.88c0 .07-.07.12-.17.12S9 17 9 16.88s.08-.12.17-.12s.18.05.18.12m-1-.15c0 .07 0 .15.14.17a.15.15 0 0 0 .2-.07c0-.07 0-.14-.14-.17s-.18 0-.2.07m1.42-.05c-.09 0-.15.08-.14.16s.09.11.19.09s.15-.09.14-.16s-.09-.1-.19-.09M11.9 4A7.83 7.83 0 0 0 4 12.07A8.29 8.29 0 0 0 9.47 20c.41.07.56-.19.56-.4v-2s-2.26.5-2.74-1c0 0-.36-1-.89-1.21c0 0-.74-.52.05-.51a1.67 1.67 0 0 1 1.24.85a1.69 1.69 0 0 0 2.36.69a1.83 1.83 0 0 1 .51-1.11c-1.8-.21-3.62-.47-3.62-3.66A2.54 2.54 0 0 1 7.7 9.7a3.21 3.21 0 0 1 .08-2.24c.68-.22 2.23.89 2.23.89a7.46 7.46 0 0 1 4.05 0s1.55-1.11 2.23-.89a3.14 3.14 0 0 1 .08 2.24a2.61 2.61 0 0 1 .83 1.95c0 3.2-1.9 3.45-3.7 3.66a2 2 0 0 1 .5 1.5v2.77a.42.42 0 0 0 .56.4A8.22 8.22 0 0 0 20 12.07A8 8 0 0 0 11.9 4M7.14 15.41v.17a.12.12 0 0 0 .17 0s0-.11 0-.18s-.13-.03-.17.01m-.35-.27s0 .1.07.13a.09.09 0 0 0 .14 0s0-.1-.07-.13s-.12-.03-.14 0m1 1.18v.21c0 .07.17.08.21 0s0-.14 0-.21s-.13-.05-.17 0Zm-.37-.49v.2c0 .08.14.11.19.08a.16.16 0 0 0 0-.21c-.05-.08-.13-.11-.19-.07"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Globe(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 3a9 9 0 1 0 9 9a9 9 0 0 0-9-9m7.46 8.25H16.7a13 13 0 0 0-2.94-6.53a7.52 7.52 0 0 1 5.7 6.53m-10.65 1.5h6.38A13.18 13.18 0 0 1 12 19.1a13.18 13.18 0 0 1-3.19-6.35m0-1.5A13.18 13.18 0 0 1 12 4.9a13.18 13.18 0 0 1 3.19 6.35Zm1.43-6.53a13 13 0 0 0-2.94 6.53H4.54a7.52 7.52 0 0 1 5.7-6.53m-5.7 8H7.3a13 13 0 0 0 2.94 6.53a7.52 7.52 0 0 1-5.7-6.5Zm9.22 6.53a13 13 0 0 0 2.94-6.53h2.76a7.52 7.52 0 0 1-5.7 6.56Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Google(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m19.76 10.77l-.09-.35h-7.44v3.16h4.45a4.45 4.45 0 0 1-4.36 3.34a5.21 5.21 0 0 1-3.5-1.39A5 5 0 0 1 7.33 12a5.14 5.14 0 0 1 1.46-3.53a5 5 0 0 1 3.48-1.37a4.55 4.55 0 0 1 3 1.16L17.47 6a7.88 7.88 0 0 0-5.27-2a8.14 8.14 0 0 0-5.77 2.35a8.15 8.15 0 0 0-.09 11.21a8.37 8.37 0 0 0 6 2.44a7.45 7.45 0 0 0 5.41-2.27a8 8 0 0 0 2.08-5.54a9.88 9.88 0 0 0-.07-1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hashtag(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 7.25h-3l.77-3.07a.75.75 0 0 0-1.46-.36l-.86 3.43H10l.77-3.07a.75.75 0 0 0-1.46-.36l-.9 3.43H5a.75.75 0 0 0 0 1.5h3l-1.63 6.5H3a.75.75 0 0 0 0 1.5h3l-.77 3.07a.75.75 0 0 0 1.46.36l.86-3.43H14l-.77 3.07a.75.75 0 0 0 1.46.36l.86-3.43H19a.75.75 0 0 0 0-1.5h-3l1.63-6.5H21a.75.75 0 0 0 0-1.5m-5 1.5l-1.63 6.5H8l1.63-6.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heart(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 19.75a.75.75 0 0 1-.53-.22L4.7 12.74a5 5 0 0 1 0-7a4.95 4.95 0 0 1 7 0L12 6l.28-.28a4.92 4.92 0 0 1 3.51-1.46a4.92 4.92 0 0 1 3.51 1.45a5 5 0 0 1 0 7l-6.77 6.79a.75.75 0 0 1-.53.25m-3.79-14a3.44 3.44 0 0 0-2.45 1a3.48 3.48 0 0 0 0 4.91L12 17.94l6.23-6.26a3.47 3.47 0 0 0 0-4.91a3.4 3.4 0 0 0-2.44-1a3.44 3.44 0 0 0-2.45 1l-.81.81a.77.77 0 0 1-1.06 0l-.81-.81a3.44 3.44 0 0 0-2.45-1.02"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartFill(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.3 5.71a4.92 4.92 0 0 0-3.51-1.46a4.92 4.92 0 0 0-3.51 1.46L12 6l-.28-.28a4.95 4.95 0 0 0-7 0a5 5 0 0 0 0 7l6.77 6.79a.75.75 0 0 0 1.06 0l6.77-6.79a5 5 0 0 0-.02-7.01"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func History(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.05 6.08a8.39 8.39 0 0 0-11.84 0L5 7.29V4.38a.75.75 0 0 0-1.5 0v4.74a.76.76 0 0 0 .75.76H9a.75.75 0 0 0 0-1.5H6l1.27-1.24a6.88 6.88 0 0 1 9.72 0c6.19 6.69-3 15.91-9.72 9.72a.75.75 0 0 0-1.06 0a.74.74 0 0 0 0 1.06A8.37 8.37 0 0 0 18.05 6.08"/><path fill="currentColor" d="M12 7.75a.76.76 0 0 0-.75.75V12a.75.75 0 0 0 .22.53L14 15a.74.74 0 0 0 .53.22A.75.75 0 0 0 15 14l-2.28-2.28V8.5a.76.76 0 0 0-.72-.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Home(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 11.75a.74.74 0 0 1-.45-.15L12 5.94L4.45 11.6a.75.75 0 0 1-.9-1.2l8-6a.75.75 0 0 1 .9 0l8 6a.75.75 0 0 1 .15 1a.74.74 0 0 1-.6.35"/><path fill="currentColor" d="M18 19.75H6a.76.76 0 0 1-.75-.75V9.5a.75.75 0 0 1 1.5 0v8.75h10.5V9.5a.75.75 0 0 1 1.5 0V19a.76.76 0 0 1-.75.75"/><path fill="currentColor" d="M14 19.75a.76.76 0 0 1-.75-.75v-6.25h-2.5V19a.75.75 0 0 1-1.5 0v-7a.76.76 0 0 1 .75-.75h4a.76.76 0 0 1 .75.75v7a.76.76 0 0 1-.75.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hourglass(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.25 19.25h-.66c-.38-1.36-1.55-4.74-4.4-7.25c2.85-2.51 4.02-5.89 4.4-7.25h.66c.41 0 .75-.34.75-.75s-.34-.75-.75-.75H5.75c-.41 0-.75.34-.75.75s.34.75.75.75h.66c.38 1.36 1.55 4.74 4.4 7.25c-2.85 2.51-4.02 5.89-4.4 7.25h-.66c-.41 0-.75.34-.75.75s.34.75.75.75h12.5c.41 0 .75-.34.75-.75s-.34-.75-.75-.75M7.98 4.75h8.03c-.45 1.42-1.6 4.25-4.02 6.28c-2.42-2.04-3.57-4.86-4.02-6.28ZM12 12.97c2.42 2.04 3.57 4.86 4.02 6.28H7.98c.45-1.42 1.6-4.25 4.02-6.28"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IdCard(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 18.75H5A1.76 1.76 0 0 1 3.25 17V7A1.76 1.76 0 0 1 5 5.25h14A1.76 1.76 0 0 1 20.75 7v10A1.76 1.76 0 0 1 19 18.75m-14-12a.25.25 0 0 0-.25.25v10a.25.25 0 0 0 .25.25h14a.25.25 0 0 0 .25-.25V7a.25.25 0 0 0-.25-.25Z"/><path fill="currentColor" d="M9 11.75a2 2 0 1 1 2-2a2 2 0 0 1-2 2m0-2.5a.5.5 0 1 0 .5.5a.5.5 0 0 0-.5-.5m3 6.5a.76.76 0 0 1-.75-.75c0-.68-.17-1.25-2.25-1.25s-2.25.57-2.25 1.25a.75.75 0 0 1-1.5 0c0-2.75 2.82-2.75 3.75-2.75s3.75 0 3.75 2.75a.76.76 0 0 1-.75.75m5-5h-3a.75.75 0 0 1 0-1.5h3a.75.75 0 0 1 0 1.5m-1 3h-2a.75.75 0 0 1 0-1.5h2a.75.75 0 0 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Image(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 4.25H6A2.75 2.75 0 0 0 3.25 7v10A2.75 2.75 0 0 0 6 19.75h12A2.75 2.75 0 0 0 20.75 17V7A2.75 2.75 0 0 0 18 4.25M6 5.75h12A1.25 1.25 0 0 1 19.25 7v8.19l-2.72-2.72a.7.7 0 0 0-.56-.22a.79.79 0 0 0-.55.27l-1.29 1.55l-4.6-4.6A.7.7 0 0 0 9 9.25a.79.79 0 0 0-.55.27l-3.7 4.41V7A1.25 1.25 0 0 1 6 5.75M4.75 17v-.73l4.3-5.16l4.12 4.12l-2.52 3H6A1.25 1.25 0 0 1 4.75 17M18 18.25h-5.4l3.45-4.14l3.15 3.15a1.23 1.23 0 0 1-1.2.99"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Images(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.5 3.75h-10A2.75 2.75 0 0 0 5.75 6.5v.25H5.5A2.75 2.75 0 0 0 2.75 9.5v8a2.75 2.75 0 0 0 2.75 2.75h10a2.75 2.75 0 0 0 2.75-2.75v-.25h.25a2.75 2.75 0 0 0 2.75-2.75v-8a2.75 2.75 0 0 0-2.75-2.75M7.25 6.5A1.25 1.25 0 0 1 8.5 5.25h10a1.25 1.25 0 0 1 1.25 1.25v6.2l-2.27-1.91a.74.74 0 0 0-1.05.08l-1.07 1.26l-4-3.88a.7.7 0 0 0-.52-.25a.75.75 0 0 0-.54.26l-3.05 3.63Zm1.25 9.25a1.25 1.25 0 0 1-1.25-1.25v-.3l3.67-4.32l3.46 3.39l-2.1 2.48Zm8.25 1.75a1.25 1.25 0 0 1-1.25 1.25h-10a1.25 1.25 0 0 1-1.25-1.25v-8A1.25 1.25 0 0 1 5.5 8.25h.25v6.25a2.75 2.75 0 0 0 2.75 2.75h8.25Zm1.75-1.75h-4.25l2.84-3.34l2.63 2.23a1.23 1.23 0 0 1-1.22 1.11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Inbox(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 20H6a2.75 2.75 0 0 1-2.75-2.75v-6a.75.75 0 0 1 1.5 0v6A1.25 1.25 0 0 0 6 18.53h12a1.25 1.25 0 0 0 1.25-1.25v-6a.75.75 0 0 1 1.5 0v6A2.75 2.75 0 0 1 18 20"/><path fill="currentColor" d="M12 15.25A3.74 3.74 0 0 1 8.29 12H4a.75.75 0 0 1-.64-1.15l3.73-6A1.69 1.69 0 0 1 8.62 4h6.76a1.74 1.74 0 0 1 1.57 1l3.69 5.94A.75.75 0 0 1 20 12h-4.29A3.74 3.74 0 0 1 12 15.25m-6.65-4.72H9a.75.75 0 0 1 .75.75v.22a2.25 2.25 0 0 0 4.5 0v-.22a.75.75 0 0 1 .75-.75h3.65l-3-4.86c-.08-.14-.16-.2-.26-.2H8.62a.27.27 0 0 0-.23.14Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Info(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 17.75a.76.76 0 0 1-.75-.75v-7a.75.75 0 0 1 1.5 0v7a.76.76 0 0 1-.75.75m0-9.5a.76.76 0 0 1-.75-.75V7a.75.75 0 0 1 1.5 0v.5a.76.76 0 0 1-.75.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoCircle(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 16.75a.76.76 0 0 1-.75-.75v-5a.75.75 0 0 1 1.5 0v5a.76.76 0 0 1-.75.75m0-7.5a.76.76 0 0 1-.75-.75V8a.75.75 0 0 1 1.5 0v.5a.76.76 0 0 1-.75.75"/><path fill="currentColor" d="M12 21a9 9 0 1 1 9-9a9 9 0 0 1-9 9m0-16.5a7.5 7.5 0 1 0 7.5 7.5A7.5 7.5 0 0 0 12 4.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Instagram(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 7.9a4.1 4.1 0 1 0 4.1 4.1A4.09 4.09 0 0 0 12 7.9m0 6.77A2.67 2.67 0 1 1 14.67 12A2.67 2.67 0 0 1 12 14.67m5.23-6.94a1 1 0 1 1-1-1a1 1 0 0 1 1 1m2.71 1a4.71 4.71 0 0 0-1.29-3.35a4.71 4.71 0 0 0-3.35-1.32C14 4 10 4 8.7 4.06a4.73 4.73 0 0 0-3.35 1.29A4.71 4.71 0 0 0 4.06 8.7C4 10 4 14 4.06 15.3a4.71 4.71 0 0 0 1.29 3.35a4.73 4.73 0 0 0 3.35 1.29c1.32.08 5.28.08 6.6 0a4.71 4.71 0 0 0 3.35-1.29a4.71 4.71 0 0 0 1.29-3.35c.06-1.3.06-5.3 0-6.6Zm-1.7 8a2.7 2.7 0 0 1-1.52 1.52a18 18 0 0 1-4.72.32a17.91 17.91 0 0 1-4.71-.32a2.7 2.7 0 0 1-1.52-1.52c-.42-1.06-.33-3.56-.33-4.72s-.09-3.67.33-4.72a2.65 2.65 0 0 1 1.52-1.53A17.91 17.91 0 0 1 12 5.44a18 18 0 0 1 4.72.32a2.7 2.7 0 0 1 1.52 1.52c.42 1.06.32 3.56.32 4.72s.1 3.67-.32 4.72Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Key(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 14.75a5.74 5.74 0 0 1-4.07-1.68A5.77 5.77 0 1 1 15 14.75m0-10A4.25 4.25 0 0 0 12 12a4.25 4.25 0 1 0 3-7.25"/><path fill="currentColor" d="M4.5 20.25A.74.74 0 0 1 4 20a.75.75 0 0 1 0-1l6.46-6.47a.75.75 0 1 1 1.06 1.07L5 20a.74.74 0 0 1-.5.25"/><path fill="currentColor" d="M8 20.75a.74.74 0 0 1-.53-.22l-2-2a.75.75 0 0 1 1.06-1.06l2 2a.75.75 0 0 1 0 1.06a.74.74 0 0 1-.53.22m2-2a.74.74 0 0 1-.53-.22l-2-2a.75.75 0 1 1 1.06-1.06l2 2a.75.75 0 0 1 0 1.06a.74.74 0 0 1-.53.22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Language(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.58 19.37l-2.99-8.36c-.21-.55-.68-.89-1.22-.89s-1 .34-1.23.91l-2.98 8.34a.75.75 0 1 0 1.41.51l.62-1.73h4.35l.62 1.73c.11.31.4.5.71.5c.08 0 .17-.01.25-.04a.75.75 0 0 0 .45-.96Zm-5.84-2.73l1.64-4.59l1.64 4.59zm-2.55-8.79c-2.26 3.57-4.3 5.73-6.78 7.17a.746.746 0 0 1-1.02-.27a.738.738 0 0 1 .27-1.02c2.1-1.22 3.82-2.97 5.75-5.87H4.12c-.41 0-.75-.34-.75-.75s.34-.75.75-.75h3.75V4.38c0-.41.34-.75.75-.75s.75.34.75.75v1.98h3.75c.41 0 .75.34.75.75s-.34.75-.75.75h-.94Zm.04 7.27c-.13 0-.26-.03-.38-.1c-.65-.38-1.28-.8-1.87-1.24a.75.75 0 0 1 .9-1.2c.54.41 1.13.79 1.73 1.14a.752.752 0 0 1-.38 1.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Link(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.23 18.25a4 4 0 0 1-2.83-1.16a4.23 4.23 0 0 1 .15-5.95l3.76-3.79A4.44 4.44 0 0 1 11.42 6a4 4 0 0 1 2.83 1.2a4.25 4.25 0 0 1-.15 6l-1.26 1.26a.75.75 0 1 1-1.06-1.06L13 12.1a2.73 2.73 0 0 0 .14-3.84a2.77 2.77 0 0 0-3.8.15l-3.73 3.78A2.74 2.74 0 0 0 5.46 16a2.5 2.5 0 0 0 2 .71a.74.74 0 0 1 .81.67a.75.75 0 0 1-.67.82Z"/><path fill="currentColor" d="M12.58 18a4 4 0 0 1-2.83-1.2a4.25 4.25 0 0 1 .15-6l1.26-1.26a.75.75 0 1 1 1.06 1.06L11 11.9a2.73 2.73 0 0 0-.14 3.84a2.77 2.77 0 0 0 3.8-.15l3.77-3.78A2.74 2.74 0 0 0 18.54 8a2.5 2.5 0 0 0-2-.71a.74.74 0 0 1-.81-.67a.75.75 0 0 1 .67-.82a4 4 0 0 1 3.2 1.11a4.23 4.23 0 0 1-.15 5.95l-3.76 3.79A4.44 4.44 0 0 1 12.58 18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Linkedin(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.72 4H5.37A1.31 1.31 0 0 0 4 5.25v13.38A1.41 1.41 0 0 0 5.37 20h13.35A1.34 1.34 0 0 0 20 18.63V5.25A1.23 1.23 0 0 0 18.72 4M9 17.34H6.67v-7.13H9ZM7.89 9.13A1.18 1.18 0 0 1 6.67 7.9a1.18 1.18 0 0 1 1.24-1.23A1.18 1.18 0 0 1 9.13 7.9a1.18 1.18 0 0 1-1.24 1.23m9.45 8.21H15v-3.9c0-.93-.33-1.57-1.16-1.57a1.25 1.25 0 0 0-1.17.84a1.43 1.43 0 0 0-.08.57v4.06h-2.3v-7.13h2.3v1a2.32 2.32 0 0 1 2.1-1.21c1.51 0 2.65 1 2.65 3.13Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func List(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 12.75H8a.75.75 0 0 1 0-1.5h11a.75.75 0 0 1 0 1.5m0-4.5H8a.75.75 0 0 1 0-1.5h11a.75.75 0 0 1 0 1.5m0 9H8a.75.75 0 0 1 0-1.5h11a.75.75 0 0 1 0 1.5M5 8.5a1 1 0 0 1-.38-.07a1.46 1.46 0 0 1-.33-.22A1 1 0 0 1 4 7.5a1.05 1.05 0 0 1 .29-.71a.93.93 0 0 1 .33-.21a1 1 0 0 1 .76 0a1 1 0 0 1 .33.21A1.05 1.05 0 0 1 6 7.5a1 1 0 0 1-.29.71a1.46 1.46 0 0 1-.33.22A1 1 0 0 1 5 8.5M5 13a1 1 0 0 1-.38-.08a1.15 1.15 0 0 1-.33-.21A1 1 0 0 1 4 12a1.05 1.05 0 0 1 .29-.71a1.15 1.15 0 0 1 .33-.21A1 1 0 0 1 5.2 11l.18.06l.18.09a1.58 1.58 0 0 1 .15.12A1.05 1.05 0 0 1 6 12a1 1 0 0 1-1 1m0 4.5a1 1 0 0 1-.38-.07a1.46 1.46 0 0 1-.33-.22a1.15 1.15 0 0 1-.21-.33a.94.94 0 0 1 0-.76a1.15 1.15 0 0 1 .21-.33a1 1 0 0 1 1.09-.21a1 1 0 0 1 .33.21a1.15 1.15 0 0 1 .21.33a.94.94 0 0 1 0 .76a1.15 1.15 0 0 1-.21.33a1 1 0 0 1-.71.29"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lock(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 10.25h-.25V8a4.75 4.75 0 0 0-9.5 0v2.25H7A2.75 2.75 0 0 0 4.25 13v5A2.75 2.75 0 0 0 7 20.75h10A2.75 2.75 0 0 0 19.75 18v-5A2.75 2.75 0 0 0 17 10.25M8.75 8a3.25 3.25 0 0 1 6.5 0v2.25h-6.5Zm9.5 10A1.25 1.25 0 0 1 17 19.25H7A1.25 1.25 0 0 1 5.75 18v-5A1.25 1.25 0 0 1 7 11.75h10A1.25 1.25 0 0 1 18.25 13Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockOpen(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3.25A4.75 4.75 0 0 0 11.25 8v2.25H6A2.75 2.75 0 0 0 3.25 13v5A2.75 2.75 0 0 0 6 20.75h7A2.75 2.75 0 0 0 15.75 18v-5A2.75 2.75 0 0 0 13 10.25h-.25V8a3.25 3.25 0 0 1 6.5 0a.75.75 0 0 0 1.5 0A4.75 4.75 0 0 0 16 3.25M14.25 13v5A1.25 1.25 0 0 1 13 19.25H6A1.25 1.25 0 0 1 4.75 18v-5A1.25 1.25 0 0 1 6 11.75h7A1.25 1.25 0 0 1 14.25 13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Map(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.9 4.09a1.75 1.75 0 0 0-1.66-.15l-3.57 1.53l-4.75-2a1.51 1.51 0 0 0-1.18 0L4.38 5.31a1.88 1.88 0 0 0-1.13 1.75v11.25a1.91 1.91 0 0 0 .85 1.6a1.75 1.75 0 0 0 1.66.15l3.57-1.53l4.75 2a1.45 1.45 0 0 0 .59.12a1.52 1.52 0 0 0 .59-.12l4.36-1.87a1.88 1.88 0 0 0 1.13-1.75V5.69a1.91 1.91 0 0 0-.85-1.6m-9.82 1l3.84 1.64v12.13l-3.84-1.64ZM5.17 18.68a.25.25 0 0 1-.25 0a.4.4 0 0 1-.17-.35V7.06A.39.39 0 0 1 5 6.69l3.58-1.55v12.08Zm14.08-1.74a.39.39 0 0 1-.22.37l-3.61 1.55V6.78l3.41-1.46a.25.25 0 0 1 .25 0a.4.4 0 0 1 .17.35Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarker(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21.25a.69.69 0 0 1-.41-.13c-.3-.19-7.34-4.92-7.34-10.67a7.75 7.75 0 0 1 15.5 0c0 5.75-7 10.48-7.34 10.67a.69.69 0 0 1-.41.13m0-17a6.23 6.23 0 0 0-6.25 6.2c0 4.21 4.79 8.06 6.25 9.13c1.46-1.07 6.25-4.92 6.25-9.13A6.23 6.23 0 0 0 12 4.25"/><path fill="currentColor" d="M12 12.75A2.75 2.75 0 1 1 14.75 10A2.75 2.75 0 0 1 12 12.75m0-4A1.25 1.25 0 1 0 13.25 10A1.25 1.25 0 0 0 12 8.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Megaphone(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.17 4.44h-1.34c-1.01 0-1.83.82-1.83 1.83v.45L5.5 8.76v-.51c0-.41-.34-.75-.75-.75S4 7.84 4 8.25v7c0 .41.34.75.75.75s.75-.34.75-.75v-.63l9.5 2.04v.45c0 1.01.82 1.83 1.83 1.83h1.34c1.01 0 1.83-.82 1.83-1.83V6.27c0-1.01-.82-1.83-1.83-1.83M5.5 13.08v-2.79L15 8.25v6.87zm13 4.03c0 .18-.15.33-.33.33h-1.34c-.18 0-.33-.15-.33-.33V6.27c0-.18.15-.33.33-.33h1.34c.18 0 .33.15.33.33zm-5.53.52a2.785 2.785 0 0 1-2.67 1.94c-1.54 0-2.8-1.24-2.8-2.75c0-.11 0-.23.02-.34c.05-.41.43-.7.84-.65c.41.05.7.42.65.83v.15c0 .69.58 1.25 1.3 1.25c.57 0 1.07-.36 1.24-.89c.12-.4.55-.62.94-.49c.4.12.61.55.49.94Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Microphone(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.69 16.15h.62c1.48 0 2.67-1.23 2.67-2.75V6c0-1.52-1.2-2.75-2.67-2.75h-.62c-1.48 0-2.67 1.23-2.67 2.75v7.39c0 1.52 1.2 2.75 2.67 2.75ZM10.52 6c0-.69.53-1.25 1.17-1.25h.62c.65 0 1.17.56 1.17 1.25v7.39c0 .69-.53 1.25-1.17 1.25h-.62c-.65 0-1.17-.56-1.17-1.25zm7.23 3.39v4.23c0 2.81-2.22 5.09-5 5.23v1.9c0 .41-.34.75-.75.75s-.75-.34-.75-.75v-1.9c-2.78-.14-5-2.42-5-5.23V9.39c0-.41.34-.75.75-.75s.75.34.75.75v4.23a3.77 3.77 0 0 0 3.77 3.76h.97a3.77 3.77 0 0 0 3.77-3.76V9.39c0-.41.34-.75.75-.75s.75.34.75.75Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Microsoft(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 4h7.5v7.5H4Zm8.5 0H20v7.5h-7.5ZM4 12.5h7.5V20H4Zm8.5 0H20V20h-7.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Minus(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 13H4a1 1 0 0 1 0-2h16a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusCircle(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21a9 9 0 1 1 9-9a9 9 0 0 1-9 9m0-16.5a7.5 7.5 0 1 0 7.5 7.5A7.5 7.5 0 0 0 12 4.5"/><path fill="currentColor" d="M16 12.75H8a.75.75 0 0 1 0-1.5h8a.75.75 0 0 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mobile(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3.25H8c-.96 0-1.75.79-1.75 1.75v14c0 .96.79 1.75 1.75 1.75h8c.96 0 1.75-.79 1.75-1.75V5c0-.96-.79-1.75-1.75-1.75M16.25 19c0 .14-.11.25-.25.25H8c-.14 0-.25-.11-.25-.25V5c0-.14.11-.25.25-.25h8c.14 0 .25.11.25.25zM13 16c0 .55-.45 1-1 1s-1-.45-1-1s.45-1 1-1s1 .45 1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoneyBill(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 5.25H5A1.76 1.76 0 0 0 3.25 7v10A1.76 1.76 0 0 0 5 18.75h14A1.76 1.76 0 0 0 20.75 17V7A1.76 1.76 0 0 0 19 5.25M19.25 7v1.67h-.21a1.45 1.45 0 0 1-.36 0h-.13a1.8 1.8 0 0 1-.33-.1h-.12a1.29 1.29 0 0 1-.38-.28a1.58 1.58 0 0 1-.47-1.12a1.45 1.45 0 0 1 .06-.41H19a.25.25 0 0 1 .25.24M5 6.75h1.69a1.45 1.45 0 0 1 .06.41a1.58 1.58 0 0 1-.47 1.12a1.49 1.49 0 0 1-.38.28h-.13a1.12 1.12 0 0 1-.31.1h-.14a1.45 1.45 0 0 1-.36 0h-.21V7A.25.25 0 0 1 5 6.75M4.75 17v-1.67h.21a1.45 1.45 0 0 1 .36 0h.13a1.26 1.26 0 0 1 .33.1h.12a1.49 1.49 0 0 1 .38.28a1.58 1.58 0 0 1 .47 1.12a1.45 1.45 0 0 1-.06.41H5a.25.25 0 0 1-.25-.24m3.47.25a2.73 2.73 0 0 0 0-.41a3 3 0 0 0-.91-2.18a2.63 2.63 0 0 0-.49-.39l-.17-.11a3.39 3.39 0 0 0-.36-.16l-.2-.08a3.33 3.33 0 0 0-.53-.11a1.56 1.56 0 0 0-.31 0h-.5v-3.59h.85a2.71 2.71 0 0 0 .4-.1l.22-.07a2.48 2.48 0 0 0 .44-.2l.16-.09a3.42 3.42 0 0 0 .52-.42a3 3 0 0 0 .91-2.18a2.73 2.73 0 0 0 0-.41h7.56a2.73 2.73 0 0 0 0 .41a3 3 0 0 0 .91 2.18a2.9 2.9 0 0 0 .52.42l.16.09a2.48 2.48 0 0 0 .44.2l.22.07a2.71 2.71 0 0 0 .41.09h.84v3.56h-.5a1.62 1.62 0 0 0-.31 0a3.33 3.33 0 0 0-.53.11l-.2.08a3.39 3.39 0 0 0-.39.18l-.17.11a2.63 2.63 0 0 0-.49.39a3 3 0 0 0-.91 2.18a2.73 2.73 0 0 0 0 .41Zm10.78 0h-1.69a1.45 1.45 0 0 1-.06-.41a1.58 1.58 0 0 1 .47-1.12a1.29 1.29 0 0 1 .38-.28h.12a2.07 2.07 0 0 1 .33-.1h.13a1.45 1.45 0 0 1 .36 0h.21V17a.25.25 0 0 1-.25.25"/><path fill="currentColor" d="M12 8.5a3.5 3.5 0 1 0 3.5 3.5A3.5 3.5 0 0 0 12 8.5m0 5.5a2 2 0 1 1 2-2a2 2 0 0 1-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Moon(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.09 20.66a8.68 8.68 0 0 1-1.09-.07A8.8 8.8 0 0 1 3.41 13a8.71 8.71 0 0 1 6.83-9.67a1.23 1.23 0 0 1 1.27.48a1.27 1.27 0 0 1 .05 1.4a5.3 5.3 0 0 0-.66 3.47a5.24 5.24 0 0 0 4.38 4.38a5.19 5.19 0 0 0 3.47-.67a1.27 1.27 0 0 1 1.4.07a1.21 1.21 0 0 1 .48 1.26a8.7 8.7 0 0 1-8.54 6.94M10 5a7.25 7.25 0 1 0 9 9a6.74 6.74 0 0 1-9.61-5A6.75 6.75 0 0 1 10 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Palette(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.07 4.25h-.11c-3.79 0-7.09 2.79-7.69 6.54c-.08.5.06 1.02.4 1.41c.33.39.8.61 1.3.61h.06c1.74 0 3.05.43 3.89 1.27c.85.85 1.27 2.18 1.26 3.95c0 .5.22.97.6 1.3c.32.28.73.43 1.14.42c.09 0 .18 0 .27-.02c3.79-.61 6.6-3.96 6.55-7.8c-.06-4.17-3.5-7.62-7.68-7.68Zm.89 14c-.06 0-.14-.01-.2-.06c-.03-.03-.08-.08-.08-.15c.01-2.19-.56-3.88-1.7-5.02c-1.13-1.13-2.8-1.71-4.95-1.71h-.06c-.08 0-.13-.05-.16-.08c-.02-.03-.08-.1-.06-.2c.49-3.02 3.15-5.28 6.21-5.28h.09c3.31.05 6.15 2.88 6.2 6.2c.04 3.1-2.23 5.81-5.28 6.3ZM9.94 9.02c0 .55-.45 1.01-1.01 1.01s-1.01-.45-1.01-1.01s.45-1.01 1.01-1.01s1.01.45 1.01 1.01m3.02-1.26c0 .55-.45 1.01-1.01 1.01s-1.01-.45-1.01-1.01s.45-1.01 1.01-1.01s1.01.45 1.01 1.01m1.01 1.26c0-.55.45-1.01 1.01-1.01s1.01.45 1.01 1.01s-.45 1.01-1.01 1.01s-1.01-.45-1.01-1.01m3.27 3.02c0 .55-.45 1.01-1.01 1.01s-1.01-.45-1.01-1.01s.45-1.01 1.01-1.01s1.01.45 1.01 1.01m-1.26 3.02c0 .55-.45 1.01-1.01 1.01s-1.01-.45-1.01-1.01s.45-1.01 1.01-1.01s1.01.45 1.01 1.01"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paperclip(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.94 20.74a5.85 5.85 0 0 1-4-1.56a5.23 5.23 0 0 1 0-7.68l7.56-7.14a4.22 4.22 0 0 1 5.69 0A4.1 4.1 0 0 1 19.5 7.3a3.46 3.46 0 0 1-1.1 2.55L10.83 17a2.47 2.47 0 0 1-3.36 0a2.23 2.23 0 0 1 0-3.28l7-6.59a.75.75 0 0 1 1.06 0a.75.75 0 0 1 0 1.06l-7 6.59a.73.73 0 0 0 0 1.1a1 1 0 0 0 1.3 0l7.57-7.13A2 2 0 0 0 18 7.3a2.57 2.57 0 0 0-.84-1.84a2.67 2.67 0 0 0-3.63 0L6 12.59a3.73 3.73 0 0 0 0 5.5a4.4 4.4 0 0 0 6 0L19.49 11a.74.74 0 0 1 1.06 0a.75.75 0 0 1 0 1.06L13 19.18a5.84 5.84 0 0 1-4.06 1.56"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pause(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 19.75a.76.76 0 0 1-.75-.75V5a.75.75 0 0 1 1.5 0v14a.76.76 0 0 1-.75.75m6 0a.76.76 0 0 1-.75-.75V5a.75.75 0 0 1 1.5 0v14a.76.76 0 0 1-.75.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paypal(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.16 13.41c-.12.67-.61 3.83-.75 4.72c0 .07 0 .09-.11.09H5.68a.43.43 0 0 1-.43-.49L7.31 4.61A.73.73 0 0 1 8 4c5.35 0 5.8-.13 7.17.41c2.11.82 2.31 2.8 1.55 4.95s-2.55 3.15-4.93 3.18c-1.52 0-2.44-.24-2.65.86Zm8.64-5.08c-.06-.05-.08-.06-.1 0a8 8 0 0 1-.31 1.19c-1.4 4-5.29 3.67-7.19 3.67a.35.35 0 0 0-.39.33c-.79 5-1 6-1 6a.38.38 0 0 0 .37.46h2.24a.63.63 0 0 0 .61-.53c0-.19 0 .22.51-3.22c.16-.77.5-.69 1-.69c2.49 0 4.44-1 5-4a3.14 3.14 0 0 0-.84-3.27Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pencil(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.21 20.52a.73.73 0 0 1-.52-.21a.75.75 0 0 1-.22-.6l.31-3.84A.73.73 0 0 1 4 15.4L15.06 4.34a3.19 3.19 0 0 1 2.28-.86a3.3 3.3 0 0 1 2.25.91a3.31 3.31 0 0 1 .11 4.5L8.63 20a.77.77 0 0 1-.46.22l-3.89.35Zm1-4.26L5 19l2.74-.25l10.9-10.92A1.72 1.72 0 0 0 17.31 5a1.61 1.61 0 0 0-1.19.42ZM15.59 4.87"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Percentage(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.05 17.7a.74.74 0 0 1-.53-.22a.75.75 0 0 1 0-1.06l9.9-9.9a.75.75 0 1 1 1.06 1.06l-9.9 9.9a.74.74 0 0 1-.53.22m1.45-6.95a2.25 2.25 0 1 1 2.25-2.25a2.25 2.25 0 0 1-2.25 2.25m0-3a.75.75 0 1 0 .75.75a.76.76 0 0 0-.75-.75m7 10a2.25 2.25 0 1 1 2.25-2.25a2.25 2.25 0 0 1-2.25 2.25m0-3a.75.75 0 1 0 .75.75a.76.76 0 0 0-.75-.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Phone(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.4 20.75h-.23a16.73 16.73 0 0 1-7.27-2.58a16.58 16.58 0 0 1-5.06-5.05a16.72 16.72 0 0 1-2.58-7.29A2.3 2.3 0 0 1 3.8 4.1a2.32 2.32 0 0 1 1.6-.84H8a2.36 2.36 0 0 1 2.33 2a9.29 9.29 0 0 0 .53 2.09a2.37 2.37 0 0 1-.54 2.49l-.61.61a12 12 0 0 0 3.77 3.75l.61-.6a2.37 2.37 0 0 1 2.49-.54a9.57 9.57 0 0 0 2.09.53a2.35 2.35 0 0 1 2 2.38v2.4a2.36 2.36 0 0 1-2.35 2.36ZM8 4.75H5.61a.87.87 0 0 0-.61.31a.83.83 0 0 0-.2.62a15.17 15.17 0 0 0 2.31 6.62a15 15 0 0 0 4.59 4.59a15.34 15.34 0 0 0 6.63 2.36A.89.89 0 0 0 19 19a.88.88 0 0 0 .25-.61V16a.86.86 0 0 0-.74-.87a11.42 11.42 0 0 1-2.42-.6a.87.87 0 0 0-.91.19l-1 1a.76.76 0 0 1-.9.12a13.53 13.53 0 0 1-5.11-5.1a.74.74 0 0 1 .12-.9l1-1a.85.85 0 0 0 .19-.9a11.28 11.28 0 0 1-.6-2.42a.87.87 0 0 0-.88-.77"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Play(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 19.75a.75.75 0 0 1-.29-.06a.74.74 0 0 1-.46-.69V5A.75.75 0 0 1 9 4.47l7 7a.75.75 0 0 1 0 1.06l-7 7a.77.77 0 0 1-.5.22m.75-12.94v10.38L14.44 12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plus(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.75 11.25V5a.75.75 0 0 0-1.5 0v6.25H5a.75.75 0 0 0 0 1.5h6.25V19a.76.76 0 0 0 .75.75a.75.75 0 0 0 .75-.75v-6.25H19a.75.75 0 0 0 .75-.75a.76.76 0 0 0-.75-.75Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusCircle(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21a9 9 0 1 1 9-9a9 9 0 0 1-9 9m0-16.5a7.5 7.5 0 1 0 7.5 7.5A7.5 7.5 0 0 0 12 4.5"/><path fill="currentColor" d="M12 16.75a.76.76 0 0 1-.75-.75V8a.75.75 0 0 1 1.5 0v8a.76.76 0 0 1-.75.75"/><path fill="currentColor" d="M16 12.75H8a.75.75 0 0 1 0-1.5h8a.75.75 0 0 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pound(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.54 19.25H8.1l2-2.82a.76.76 0 0 0 .14-.43v-3.25h3.25a.75.75 0 0 0 0-1.5h-3.28V8a3.09 3.09 0 0 1 3.17-3.25a3.14 3.14 0 0 1 3.33 3.41v1a.75.75 0 0 0 1.5 0v-1a4.62 4.62 0 0 0-4.83-4.91A4.57 4.57 0 0 0 8.71 8v3.25H6.46a.75.75 0 0 0 0 1.5h2.25v3l-2.66 3.82a.76.76 0 0 0 0 .78a.74.74 0 0 0 .66.4h10.83a.75.75 0 0 0 0-1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PowerOff(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21A9 9 0 0 1 5.64 5.64a.74.74 0 0 1 1.06 0a.75.75 0 0 1 0 1.06a7.5 7.5 0 1 0 10.6 10.6a7.48 7.48 0 0 0 0-10.6a.75.75 0 0 1 0-1.06a.74.74 0 0 1 1.06 0A9 9 0 0 1 12 21"/><path fill="currentColor" d="M12 12.75a.76.76 0 0 1-.75-.75V4a.75.75 0 0 1 1.5 0v8a.76.76 0 0 1-.75.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Prime(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.36 11.33l-1.1-.25l.86 1.22v3.79l2.93-2.44V9.5l-1.35.48zm-6.84 0L7.17 9.98L5.83 9.5v4.15l2.93 2.44V12.3l.86-1.22z"/><path fill="currentColor" d="M13.16 11.45h-2.44l-.61-.37l-.98 1.47v5.5l.73 1.09l.86.86h2.44l.86-.86l.73-1.09v-5.5l-.98-1.47zm1.96 6.96l1.58-1.59v-1.58l-1.58 1.34zm-7.95-1.59l1.59 1.59v-1.83l-1.59-1.34zm3.91-5.86h.62V4h-1.35l-.97 2.31l-4.41-.36l.74 3.06l5.25 1.95zm3.42-4.65L13.53 4h-1.35v7H13l5.25-2L19 6Z"/><path fill="currentColor" d="M17.32 5.71L15.6 4h-1.71l.86 1.95zM9.98 4H8.27L6.56 5.71l2.57.24z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Print(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 16.75h-2a.75.75 0 0 1 0-1.5h2A1.25 1.25 0 0 0 19.25 14v-4A1.25 1.25 0 0 0 18 8.75H6A1.25 1.25 0 0 0 4.75 10v4A1.25 1.25 0 0 0 6 15.25h2a.75.75 0 0 1 0 1.5H6A2.75 2.75 0 0 1 3.25 14v-4A2.75 2.75 0 0 1 6 7.25h12A2.75 2.75 0 0 1 20.75 10v4A2.75 2.75 0 0 1 18 16.75"/><path fill="currentColor" d="M16 8.75a.76.76 0 0 1-.75-.75V4.75h-6.5V8a.75.75 0 0 1-1.5 0V4.5A1.25 1.25 0 0 1 8.5 3.25h7a1.25 1.25 0 0 1 1.25 1.25V8a.76.76 0 0 1-.75.75m-.5 12h-7a1.25 1.25 0 0 1-1.25-1.25v-7a1.25 1.25 0 0 1 1.25-1.25h7a1.25 1.25 0 0 1 1.25 1.25v7a1.25 1.25 0 0 1-1.25 1.25m-6.75-1.5h6.5v-6.5h-6.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Qrcode(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.38 3.88v7h7v-7Zm5.5 5.5h-4v-4h4Zm-14 1.5h7v-7h-7Zm1.5-5.5h4v4h-4Zm-1.5 14h7v-7h-7Zm1.5-5.5h4v4h-4Zm7-1.5h1.75v1.75h-1.75zm3.5 0h1.75v1.75h-1.75zm-1.75 1.75h1.75v1.75h-1.75zm3.5 0h1.75v1.75h-1.75zm-5.25 1.75h1.75v1.75h-1.75zm3.5 0h1.75v1.75h-1.75zm-1.75 1.75h1.75v1.75h-1.75zm3.5 0h1.75v1.75h-1.75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Question(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.07 4.93A5.75 5.75 0 0 0 6.25 9a.75.75 0 0 0 1.5 0A4.26 4.26 0 1 1 12 13.25a.76.76 0 0 0-.75.75v2a.75.75 0 0 0 1.5 0v-1.3a5.76 5.76 0 0 0 3.32-9.77"/><circle cx="12" cy="19.5" r="1.25" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionCircle(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 3c-4.96 0-9 4.04-9 9s4.04 9 9 9s9-4.04 9-9s-4.04-9-9-9m0 16.5c-4.14 0-7.5-3.36-7.5-7.5S7.86 4.5 12 4.5s7.5 3.36 7.5 7.5s-3.36 7.5-7.5 7.5m2.3-11.8c.61.61.95 1.43.95 2.3s-.34 1.68-.95 2.3c-.43.43-.97.73-1.55.86v.34c0 .41-.34.75-.75.75s-.75-.34-.75-.75v-1c0-.41.34-.75.75-.75A1.739 1.739 0 0 0 13.75 10c0-.47-.18-.91-.51-1.24c-.66-.66-1.81-.66-2.47 0c-.33.33-.51.77-.51 1.24c0 .41-.34.75-.75.75s-.75-.34-.75-.75c0-.87.34-1.68.95-2.3c1.23-1.23 3.37-1.23 4.6 0ZM13 16.25c0 .55-.45 1-1 1s-1-.45-1-1s.45-1 1-1s1 .45 1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reddit(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.24 13.6a.8.8 0 1 1 .8-.8a.8.8 0 0 1-.8.8M20 12a8 8 0 1 1-8-8a8 8 0 0 1 8 8m-4.27-1.33A1.12 1.12 0 0 0 15 11a5.13 5.13 0 0 0-2.77-.85l.56-2.53l1.79.4a.79.79 0 0 0 .79.8a.81.81 0 0 0 .8-.81a.79.79 0 0 0-1.51-.35l-2-.44a.19.19 0 0 0-.22.14l-.62 2.79a5.18 5.18 0 0 0-2.77.85a1.07 1.07 0 1 0-1.25 1.7a2.73 2.73 0 0 0 0 .5c0 1.69 1.91 3.07 4.26 3.07s4.26-1.38 4.26-3.07a2.16 2.16 0 0 0-.06-.51a1.07 1.07 0 0 0-.48-2Zm-2.22 3.75a2.82 2.82 0 0 1-3 0a.2.2 0 0 0-.27 0a.19.19 0 0 0 0 .28a3.17 3.17 0 0 0 3.56 0a.19.19 0 0 0 0-.28a.2.2 0 0 0-.29 0m.25-2.42a.8.8 0 1 0 .8.8a.8.8 0 0 0-.8-.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Refresh(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 20.75a7.25 7.25 0 0 1 0-14.5h2.5a.75.75 0 0 1 0 1.5H12a5.75 5.75 0 1 0 5.75 5.75a.75.75 0 0 1 1.5 0A7.26 7.26 0 0 1 12 20.75"/><path fill="currentColor" d="M12 10.75a.74.74 0 0 1-.53-.22a.75.75 0 0 1 0-1.06L13.94 7l-2.47-2.47a.75.75 0 1 1 1.06-1.06l3 3a.75.75 0 0 1 0 1.06l-3 3a.74.74 0 0 1-.53.22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Replay(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 20.75a7.26 7.26 0 0 1-7.25-7.25a.75.75 0 0 1 1.5 0A5.75 5.75 0 1 0 12 7.75H9.5a.75.75 0 0 1 0-1.5H12a7.25 7.25 0 0 1 0 14.5"/><path fill="currentColor" d="M12 10.75a.74.74 0 0 1-.53-.22l-3-3a.75.75 0 0 1 0-1.06l3-3a.75.75 0 1 1 1.06 1.06L10.06 7l2.47 2.47a.75.75 0 0 1 0 1.06a.74.74 0 0 1-.53.22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reply(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.55 15.59a.75.75 0 0 1-.55-1.28l3.92-3.89L14 6.53a.75.75 0 0 1 0-1.06a.75.75 0 0 1 1.06 0l4.46 4.42a.75.75 0 0 1 0 1.06l-4.46 4.42a.7.7 0 0 1-.51.22"/><path fill="currentColor" d="M5 18.75a.76.76 0 0 1-.75-.75v-7.58A.76.76 0 0 1 5 9.67h14a.75.75 0 0 1 0 1.5H5.75V18a.76.76 0 0 1-.75.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Save(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 20.75H7A2.75 2.75 0 0 1 4.25 18V6A2.75 2.75 0 0 1 7 3.25h7.5a.75.75 0 0 1 .53.22L19.53 8a.75.75 0 0 1 .22.53V18A2.75 2.75 0 0 1 17 20.75m-10-16A1.25 1.25 0 0 0 5.75 6v12A1.25 1.25 0 0 0 7 19.25h10A1.25 1.25 0 0 0 18.25 18V8.81l-4.06-4.06Z"/><path fill="currentColor" d="M16.75 20h-1.5v-6.25h-6.5V20h-1.5v-6.5a1.25 1.25 0 0 1 1.25-1.25h7a1.25 1.25 0 0 1 1.25 1.25ZM12.47 8.75H8.53a1.29 1.29 0 0 1-1.28-1.3V4h1.5v3.25h3.5V4h1.5v3.45a1.29 1.29 0 0 1-1.28 1.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Search(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.77 18.3a7.53 7.53 0 1 1 7.53-7.53a7.53 7.53 0 0 1-7.53 7.53m0-13.55a6 6 0 1 0 6 6a6 6 0 0 0-6-6"/><path fill="currentColor" d="M20 20.75a.74.74 0 0 1-.53-.22l-4.13-4.13a.75.75 0 0 1 1.06-1.06l4.13 4.13a.75.75 0 0 1 0 1.06a.74.74 0 0 1-.53.22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchMinus(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.77 18.3a7.53 7.53 0 1 1 7.53-7.53a7.54 7.54 0 0 1-7.53 7.53m0-13.55a6 6 0 1 0 6 6a6 6 0 0 0-6-6"/><path fill="currentColor" d="M20 20.75a.74.74 0 0 1-.53-.22l-4.13-4.13a.75.75 0 0 1 1.06-1.06l4.13 4.13a.75.75 0 0 1 0 1.06a.74.74 0 0 1-.53.22m-6.75-9.25h-5a.75.75 0 0 1 0-1.5h5a.75.75 0 0 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchPlus(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.77 18.3a7.53 7.53 0 1 1 7.53-7.53a7.53 7.53 0 0 1-7.53 7.53m0-13.55a6 6 0 1 0 6 6a6 6 0 0 0-6-6"/><path fill="currentColor" d="M20 20.75a.74.74 0 0 1-.53-.22l-4.13-4.13a.75.75 0 0 1 1.06-1.06l4.13 4.13a.75.75 0 0 1 0 1.06a.74.74 0 0 1-.53.22M10.75 14a.76.76 0 0 1-.75-.75v-5a.75.75 0 0 1 1.5 0v5a.76.76 0 0 1-.75.75"/><path fill="currentColor" d="M13.25 11.5h-5a.75.75 0 0 1 0-1.5h5a.75.75 0 0 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Send(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.33 3.67a1.45 1.45 0 0 0-1.47-.35L4.23 8.2A1.44 1.44 0 0 0 4 10.85l6.07 3l3 6.09a1.44 1.44 0 0 0 1.29.79h.1a1.43 1.43 0 0 0 1.26-1l4.95-14.59a1.41 1.41 0 0 0-.34-1.47M4.85 9.58l12.77-4.26l-7.09 7.09Zm9.58 9.57l-2.84-5.68l7.09-7.09Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Server(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.5 8.5v-3a1 1 0 0 0-1-1h-15a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h15a1 1 0 0 0 1-1v-3a1 1 0 0 0-1-1a1 1 0 0 0 1-1v-3a1 1 0 0 0-1-1a1 1 0 0 0 1-1m-1 10h-15v-3h15Zm0-5h-15v-3h15Zm0-5h-15v-3h15Z"/><circle cx="6.25" cy="7" r=".75" fill="currentColor"/><circle cx="8.75" cy="7" r=".75" fill="currentColor"/><circle cx="6.25" cy="12" r=".75" fill="currentColor"/><circle cx="8.75" cy="12" r=".75" fill="currentColor"/><circle cx="6.25" cy="17" r=".75" fill="currentColor"/><circle cx="8.75" cy="17" r=".75" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareAlt(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.5 14.25a3.24 3.24 0 0 0-2.62 1.35L9.59 13a3.39 3.39 0 0 0 .16-1a3.39 3.39 0 0 0-.16-1l5.29-2.6a3.23 3.23 0 1 0-.63-1.9a2.94 2.94 0 0 0 .05.5L8.83 9.75a3.19 3.19 0 0 0-2.33-1a3.25 3.25 0 0 0 0 6.5a3.19 3.19 0 0 0 2.33-1L14.3 17a2.94 2.94 0 0 0-.05.51a3.25 3.25 0 1 0 3.25-3.25Zm0-9.5a1.75 1.75 0 1 1-1.75 1.75a1.76 1.76 0 0 1 1.75-1.75m-11 9A1.75 1.75 0 1 1 8.25 12a1.76 1.76 0 0 1-1.75 1.75m11 5.5a1.75 1.75 0 1 1 1.75-1.75a1.76 1.76 0 0 1-1.75 1.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shield(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 20.75a.87.87 0 0 1-.28-.05A14.27 14.27 0 0 1 3.29 6.43a.74.74 0 0 1 .61-.69a27.12 27.12 0 0 0 7.79-2.42a.75.75 0 0 1 .62 0a27.12 27.12 0 0 0 7.79 2.42a.74.74 0 0 1 .61.69a14.27 14.27 0 0 1-8.43 14.27a.87.87 0 0 1-.28.05M4.76 7.11A12.47 12.47 0 0 0 12 19.18a12.47 12.47 0 0 0 7.24-12.07A27.56 27.56 0 0 1 12 4.82a27.56 27.56 0 0 1-7.24 2.29"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBag(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.5 8.25h-3v-.5a4.5 4.5 0 0 0-9 0v.5h-3A1.25 1.25 0 0 0 3.25 9.5V18A2.75 2.75 0 0 0 6 20.75h12A2.75 2.75 0 0 0 20.75 18V9.5a1.25 1.25 0 0 0-1.25-1.25M9 7.75a3 3 0 0 1 6 0v.5H9ZM19.25 18A1.25 1.25 0 0 1 18 19.25H6A1.25 1.25 0 0 1 4.75 18V9.75H7.5V12A.75.75 0 0 0 9 12V9.75h6V12a.75.75 0 0 0 1.5 0V9.75h2.75Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingCart(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.25 18.75c0 .83-.67 1.5-1.5 1.5s-1.5-.67-1.5-1.5s.67-1.5 1.5-1.5s1.5.67 1.5 1.5m5-1.5c-.83 0-1.5.67-1.5 1.5s.67 1.5 1.5 1.5s1.5-.67 1.5-1.5s-.67-1.5-1.5-1.5m4.48-9.57l-2 8a.75.75 0 0 1-.73.57H8c-.36 0-.67-.26-.74-.62L5.37 5.25H4c-.41 0-.75-.34-.75-.75s.34-.75.75-.75h2c.36 0 .67.26.74.62l.43 2.38H20a.754.754 0 0 1 .73.93m-1.69.57H7.44l1.18 6.5h8.79z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignIn(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 20.75h-3a.75.75 0 0 1 0-1.5h3a1.16 1.16 0 0 0 1.25-1V5.78a1.16 1.16 0 0 0-1.25-1h-3a.75.75 0 0 1 0-1.5h3a2.64 2.64 0 0 1 2.75 2.53v12.41A2.64 2.64 0 0 1 18 20.75"/><path fill="currentColor" d="M11 16.75a.74.74 0 0 1-.53-.22a.75.75 0 0 1 0-1.06L13.94 12l-3.47-3.47a.75.75 0 0 1 1.06-1.06l4 4a.75.75 0 0 1 0 1.06l-4 4a.74.74 0 0 1-.53.22"/><path fill="currentColor" d="M15 12.75H4a.75.75 0 0 1 0-1.5h11a.75.75 0 0 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignOut(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 20.75H6a2.64 2.64 0 0 1-2.75-2.53V5.78A2.64 2.64 0 0 1 6 3.25h3a.75.75 0 0 1 0 1.5H6a1.16 1.16 0 0 0-1.25 1v12.47a1.16 1.16 0 0 0 1.25 1h3a.75.75 0 0 1 0 1.5Zm7-4a.74.74 0 0 1-.53-.22a.75.75 0 0 1 0-1.06L18.94 12l-3.47-3.47a.75.75 0 1 1 1.06-1.06l4 4a.75.75 0 0 1 0 1.06l-4 4a.74.74 0 0 1-.53.22"/><path fill="currentColor" d="M20 12.75H9a.75.75 0 0 1 0-1.5h11a.75.75 0 0 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sitemap(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 9.75h-3A1.76 1.76 0 0 1 8.75 8V5a1.76 1.76 0 0 1 1.75-1.75h3A1.76 1.76 0 0 1 15.25 5v3a1.76 1.76 0 0 1-1.75 1.75m-3-5a.25.25 0 0 0-.25.25v3a.25.25 0 0 0 .25.25h3a.25.25 0 0 0 .25-.25V5a.25.25 0 0 0-.25-.25Zm-4.5 15H4A1.76 1.76 0 0 1 2.25 18v-2A1.76 1.76 0 0 1 4 14.25h2A1.76 1.76 0 0 1 7.75 16v2A1.76 1.76 0 0 1 6 19.75m-2-4a.25.25 0 0 0-.25.25v2a.25.25 0 0 0 .25.25h2a.25.25 0 0 0 .25-.25v-2a.25.25 0 0 0-.25-.25Zm9 4h-2A1.76 1.76 0 0 1 9.25 18v-2A1.76 1.76 0 0 1 11 14.25h2A1.76 1.76 0 0 1 14.75 16v2A1.76 1.76 0 0 1 13 19.75m-2-4a.25.25 0 0 0-.25.25v2a.25.25 0 0 0 .25.25h2a.25.25 0 0 0 .25-.25v-2a.25.25 0 0 0-.25-.25Zm9 4h-2A1.76 1.76 0 0 1 16.25 18v-2A1.76 1.76 0 0 1 18 14.25h2A1.76 1.76 0 0 1 21.75 16v2A1.76 1.76 0 0 1 20 19.75m-2-4a.25.25 0 0 0-.25.25v2a.25.25 0 0 0 .25.25h2a.25.25 0 0 0 .25-.25v-2a.25.25 0 0 0-.25-.25Z"/><path fill="currentColor" d="M19 15.75a.76.76 0 0 1-.75-.75v-2a.25.25 0 0 0-.25-.25H6a.25.25 0 0 0-.25.25v2a.75.75 0 0 1-1.5 0v-2A1.76 1.76 0 0 1 6 11.25h12A1.76 1.76 0 0 1 19.75 13v2a.76.76 0 0 1-.75.75"/><path fill="currentColor" d="M12 15.75a.76.76 0 0 1-.75-.75V9a.75.75 0 0 1 1.5 0v6a.76.76 0 0 1-.75.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Slack(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.36 14.11a1.68 1.68 0 1 1-1.68-1.68h1.68Zm.85 0a1.68 1.68 0 0 1 3.36 0v4.21a1.68 1.68 0 0 1-3.36 0Zm1.68-6.75a1.68 1.68 0 1 1 1.68-1.68v1.68Zm0 .85a1.68 1.68 0 0 1 0 3.36H5.68a1.68 1.68 0 0 1 0-3.36Zm6.75 1.68a1.68 1.68 0 1 1 1.68 1.68h-1.68zm-.85 0a1.68 1.68 0 0 1-3.36 0V5.68a1.68 1.68 0 0 1 3.36 0zm-1.68 6.75a1.68 1.68 0 1 1-1.68 1.68v-1.68Zm0-.85a1.68 1.68 0 0 1 0-3.36h4.21a1.68 1.68 0 0 1 0 3.36Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SlidersH(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 8.25h-7a.75.75 0 0 1 0-1.5h7a.75.75 0 0 1 0 1.5m-9 0H5a.75.75 0 0 1 0-1.5h5a.75.75 0 0 1 0 1.5"/><path fill="currentColor" d="M10 9.75A.76.76 0 0 1 9.25 9V6a.75.75 0 0 1 1.5 0v3a.76.76 0 0 1-.75.75m9 7.5h-7a.75.75 0 0 1 0-1.5h7a.75.75 0 0 1 0 1.5m-9 0H5a.75.75 0 0 1 0-1.5h5a.75.75 0 0 1 0 1.5"/><path fill="currentColor" d="M10 18.75a.76.76 0 0 1-.75-.75v-3a.75.75 0 0 1 1.5 0v3a.76.76 0 0 1-.75.75m9-6h-3a.75.75 0 0 1 0-1.5h3a.75.75 0 0 1 0 1.5m-5 0H5a.75.75 0 0 1 0-1.5h9a.75.75 0 0 1 0 1.5"/><path fill="currentColor" d="M14 14.25a.76.76 0 0 1-.75-.75v-3a.75.75 0 0 1 1.5 0v3a.76.76 0 0 1-.75.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SlidersV(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.5 19.75a.76.76 0 0 1-.75-.75v-7a.75.75 0 0 1 1.5 0v7a.76.76 0 0 1-.75.75m0-9a.76.76 0 0 1-.75-.75V5a.75.75 0 0 1 1.5 0v5a.76.76 0 0 1-.75.75"/><path fill="currentColor" d="M18 10.75h-3a.75.75 0 0 1 0-1.5h3a.75.75 0 0 1 0 1.5m-10.5 9a.76.76 0 0 1-.75-.75v-7a.75.75 0 0 1 1.5 0v7a.76.76 0 0 1-.75.75m0-9a.76.76 0 0 1-.75-.75V5a.75.75 0 0 1 1.5 0v5a.76.76 0 0 1-.75.75"/><path fill="currentColor" d="M9 10.75H6a.75.75 0 0 1 0-1.5h3a.75.75 0 0 1 0 1.5m3 9a.76.76 0 0 1-.75-.75v-3a.75.75 0 0 1 1.5 0v3a.76.76 0 0 1-.75.75m0-5a.76.76 0 0 1-.75-.75V5a.75.75 0 0 1 1.5 0v9a.76.76 0 0 1-.75.75"/><path fill="currentColor" d="M13.5 14.75h-3a.75.75 0 0 1 0-1.5h3a.75.75 0 0 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sort(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 10.75H6a.74.74 0 0 1-.69-.46a.75.75 0 0 1 .16-.82l6-6a.75.75 0 0 1 1.06 0l6 6a.75.75 0 0 1 .16.82a.74.74 0 0 1-.69.46M7.81 9.25h8.38L12 5.06ZM12 20.75a.74.74 0 0 1-.53-.22l-6-6a.75.75 0 0 1-.16-.82a.74.74 0 0 1 .69-.46h12a.74.74 0 0 1 .69.46a.75.75 0 0 1-.16.82l-6 6a.74.74 0 0 1-.53.22m-4.19-6L12 18.94l4.19-4.19Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortAlphaAltDown(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.22 15.97L9 17.19V5c0-.41-.34-.75-.75-.75s-.75.34-.75.75v12.19l-1.22-1.22c-.29-.29-.77-.29-1.06 0s-.29.77 0 1.06l2.5 2.5a.776.776 0 0 0 .53.22a.753.753 0 0 0 .53-.22l2.5-2.5c.29-.29.29-.77 0-1.06s-.77-.29-1.06 0m8.74 2.28l-1.71-4.79c-.17-.43-.56-.71-1-.71s-.83.28-1 .73l-1.7 4.77a.75.75 0 1 0 1.41.51l.28-.78h2.03l.28.78c.11.31.4.5.71.5c.08 0 .17-.01.25-.04a.75.75 0 0 0 .45-.96Zm-3.19-1.77l.48-1.34l.48 1.34zm-1.94-5.94c.19.44.59.71 1.05.71h3.13c.41 0 .75-.34.75-.75s-.34-.75-.75-.75h-2.39l2.82-2.93c.34-.36.44-.89.24-1.35c-.19-.44-.59-.71-1.05-.71h-3.11c-.41 0-.75.34-.75.75s.34.75.75.75h2.39l-2.83 2.95c-.34.36-.43.88-.24 1.34Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortAlphaAltUp(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.78 4.47a.776.776 0 0 0-.24-.16a.707.707 0 0 0-.57 0c-.09.04-.17.09-.24.16l-2.5 2.5c-.29.29-.29.77 0 1.06s.77.29 1.06 0l1.22-1.22V19c0 .41.34.75.75.75s.75-.34.75-.75V6.81l1.22 1.22c.15.15.34.22.53.22s.38-.07.53-.22c.29-.29.29-.77 0-1.06l-2.5-2.5Zm10.18 13.78l-1.71-4.79c-.17-.43-.56-.71-1-.71s-.83.28-1 .73l-1.7 4.77a.75.75 0 1 0 1.41.51l.28-.78h2.03l.28.78c.11.31.4.5.71.5c.08 0 .17-.01.25-.04a.75.75 0 0 0 .45-.96Zm-3.19-1.77l.48-1.34l.48 1.34zm-1.94-5.94c.19.44.59.71 1.05.71h3.13c.41 0 .75-.34.75-.75s-.34-.75-.75-.75h-2.39l2.82-2.93c.34-.36.44-.89.24-1.35c-.19-.44-.59-.71-1.05-.71h-3.11c-.41 0-.75.34-.75.75s.34.75.75.75h2.39l-2.83 2.95c-.34.36-.43.88-.24 1.34Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortAlphaDown(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.22 15.97L9 17.19V5c0-.41-.34-.75-.75-.75s-.75.34-.75.75v12.19l-1.22-1.22c-.29-.29-.77-.29-1.06 0s-.29.77 0 1.06l2.5 2.5a.776.776 0 0 0 .53.22a.776.776 0 0 0 .53-.22l2.5-2.5c.29-.29.29-.77 0-1.06s-.77-.29-1.06 0M14 11.21c.39.14.82-.06.96-.45l.28-.78h2.03l.28.78c.11.31.4.5.71.5c.08 0 .17-.01.25-.04a.75.75 0 0 0 .45-.96l-1.71-4.79c-.17-.43-.56-.71-1-.71s-.83.28-1 .73l-1.7 4.77c-.14.39.06.82.45.96Zm2.73-2.73h-.96l.48-1.34zm1.94 4.98c-.19-.44-.59-.71-1.05-.71h-3.11c-.41 0-.75.34-.75.75s.34.75.75.75h2.39l-2.83 2.95c-.34.36-.43.88-.24 1.34c.19.44.59.71 1.05.71h3.13c.41 0 .75-.34.75-.75s-.34-.75-.75-.75h-2.39l2.82-2.93c.34-.36.44-.89.24-1.35Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortAlphaUp(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.78 4.47a.776.776 0 0 0-.24-.16a.707.707 0 0 0-.57 0c-.09.04-.17.09-.24.16l-2.5 2.5c-.29.29-.29.77 0 1.06s.77.29 1.06 0l1.22-1.22V19c0 .41.34.75.75.75s.75-.34.75-.75V6.81l1.22 1.22c.15.15.34.22.53.22s.38-.07.53-.22c.29-.29.29-.77 0-1.06l-2.5-2.5ZM14 11.21c.39.14.82-.06.96-.45l.28-.78h2.03l.28.78c.11.31.4.5.71.5c.08 0 .17-.01.25-.04a.75.75 0 0 0 .45-.96l-1.71-4.79c-.17-.43-.56-.71-1-.71s-.83.28-1 .73l-1.7 4.77c-.14.39.06.82.45.96Zm2.73-2.73h-.96l.48-1.34zm1.94 4.98c-.19-.44-.59-.71-1.05-.71h-3.11c-.41 0-.75.34-.75.75s.34.75.75.75h2.39l-2.83 2.95c-.34.36-.43.88-.24 1.34c.19.44.59.71 1.05.71h3.13c.41 0 .75-.34.75-.75s-.34-.75-.75-.75h-2.39l2.82-2.93c.34-.36.44-.89.24-1.35Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortAlt(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5 8.25c-.19 0-.38-.07-.53-.22L8 6.06L6.03 8.03c-.29.29-.77.29-1.06 0s-.29-.77 0-1.06l2.5-2.5c.29-.29.77-.29 1.06 0l2.5 2.5c.29.29.29.77 0 1.06c-.15.15-.34.22-.53.22"/><path fill="currentColor" d="M8 19.75c-.41 0-.75-.34-.75-.75V5c0-.41.34-.75.75-.75s.75.34.75.75v14c0 .41-.34.75-.75.75m8 0c-.19 0-.38-.07-.53-.22l-2.5-2.5c-.29-.29-.29-.77 0-1.06s.77-.29 1.06 0L16 17.94l1.97-1.97c.29-.29.77-.29 1.06 0s.29.77 0 1.06l-2.5 2.5c-.15.15-.34.22-.53.22"/><path fill="currentColor" d="M16 19.75c-.41 0-.75-.34-.75-.75V5c0-.41.34-.75.75-.75s.75.34.75.75v14c0 .41-.34.75-.75.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortAltSlash(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m6.03 8.03l1.22-1.22v5.69c0 .41.34.75.75.75s.75-.34.75-.75V6.81l1.22 1.22c.15.15.34.22.53.22s.38-.07.53-.22c.29-.29.29-.77 0-1.06l-2.5-2.5a.776.776 0 0 0-.24-.16a.707.707 0 0 0-.57 0c-.09.04-.17.09-.24.16l-2.5 2.5c-.29.29-.29.77 0 1.06s.77.29 1.06 0Zm11.94 7.94l-1.22 1.22V11.5c0-.41-.34-.75-.75-.75s-.75.34-.75.75v5.69l-1.22-1.22c-.29-.29-.77-.29-1.06 0s-.29.77 0 1.06l2.5 2.5a.776.776 0 0 0 .53.22a.776.776 0 0 0 .53-.22l2.5-2.5c.29-.29.29-.77 0-1.06s-.77-.29-1.06 0m1.06-9.94c.29-.29.29-.77 0-1.06s-.77-.29-1.06 0l-1.22 1.22V5c0-.41-.34-.75-.75-.75s-.75.34-.75.75v2.69L4.97 17.97c-.29.29-.29.77 0 1.06c.15.15.34.22.53.22s.38-.07.53-.22l1.22-1.22V19c0 .41.34.75.75.75s.75-.34.75-.75v-2.69z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortAmountDown(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.5 19.75c-.19 0-.38-.07-.53-.22l-2.5-2.5c-.29-.29-.29-.77 0-1.06s.77-.29 1.06 0l1.97 1.97l1.97-1.97c.29-.29.77-.29 1.06 0s.29.77 0 1.06l-2.5 2.5c-.15.15-.34.22-.53.22"/><path fill="currentColor" d="M6.5 19.75c-.41 0-.75-.34-.75-.75V5c0-.41.34-.75.75-.75s.75.34.75.75v14c0 .41-.34.75-.75.75M20 8.25h-8c-.41 0-.75-.34-.75-.75s.34-.75.75-.75h8c.41 0 .75.34.75.75s-.34.75-.75.75m-4 6h-4c-.41 0-.75-.34-.75-.75s.34-.75.75-.75h4c.41 0 .75.34.75.75s-.34.75-.75.75m-2 3h-2c-.41 0-.75-.34-.75-.75s.34-.75.75-.75h2c.41 0 .75.34.75.75s-.34.75-.75.75m4-6h-6c-.41 0-.75-.34-.75-.75s.34-.75.75-.75h6c.41 0 .75.34.75.75s-.34.75-.75.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortAmountDownAlt(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.5 19.75c-.19 0-.38-.07-.53-.22l-2.5-2.5c-.29-.29-.29-.77 0-1.06s.77-.29 1.06 0l1.97 1.97l1.97-1.97c.29-.29.77-.29 1.06 0s.29.77 0 1.06l-2.5 2.5c-.15.15-.34.22-.53.22"/><path fill="currentColor" d="M6.5 19.75c-.41 0-.75-.34-.75-.75V5c0-.41.34-.75.75-.75s.75.34.75.75v14c0 .41-.34.75-.75.75m13.5-2.5h-8c-.41 0-.75-.34-.75-.75s.34-.75.75-.75h8c.41 0 .75.34.75.75s-.34.75-.75.75m-4-6h-4c-.41 0-.75-.34-.75-.75s.34-.75.75-.75h4c.41 0 .75.34.75.75s-.34.75-.75.75m-2-3h-2c-.41 0-.75-.34-.75-.75s.34-.75.75-.75h2c.41 0 .75.34.75.75s-.34.75-.75.75m4 6h-6c-.41 0-.75-.34-.75-.75s.34-.75.75-.75h6c.41 0 .75.34.75.75s-.34.75-.75.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortAmountUp(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 8.25c-.19 0-.38-.07-.53-.22L6.5 6.06L4.53 8.03c-.29.29-.77.29-1.06 0s-.29-.77 0-1.06l2.5-2.5c.29-.29.77-.29 1.06 0l2.5 2.5c.29.29.29.77 0 1.06c-.15.15-.34.22-.53.22"/><path fill="currentColor" d="M6.5 19.75c-.41 0-.75-.34-.75-.75V5c0-.41.34-.75.75-.75s.75.34.75.75v14c0 .41-.34.75-.75.75M20 8.25h-8c-.41 0-.75-.34-.75-.75s.34-.75.75-.75h8c.41 0 .75.34.75.75s-.34.75-.75.75m-4 6h-4c-.41 0-.75-.34-.75-.75s.34-.75.75-.75h4c.41 0 .75.34.75.75s-.34.75-.75.75m-2 3h-2c-.41 0-.75-.34-.75-.75s.34-.75.75-.75h2c.41 0 .75.34.75.75s-.34.75-.75.75m4-6h-6c-.41 0-.75-.34-.75-.75s.34-.75.75-.75h6c.41 0 .75.34.75.75s-.34.75-.75.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortAmountUpAlt(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 8.25c-.19 0-.38-.07-.53-.22L6.5 6.06L4.53 8.03c-.29.29-.77.29-1.06 0s-.29-.77 0-1.06l2.5-2.5c.29-.29.77-.29 1.06 0l2.5 2.5c.29.29.29.77 0 1.06c-.15.15-.34.22-.53.22"/><path fill="currentColor" d="M6.5 19.75c-.41 0-.75-.34-.75-.75V5c0-.41.34-.75.75-.75s.75.34.75.75v14c0 .41-.34.75-.75.75m13.5-2.5h-8c-.41 0-.75-.34-.75-.75s.34-.75.75-.75h8c.41 0 .75.34.75.75s-.34.75-.75.75m-4-6h-4c-.41 0-.75-.34-.75-.75s.34-.75.75-.75h4c.41 0 .75.34.75.75s-.34.75-.75.75m-2-3h-2c-.41 0-.75-.34-.75-.75s.34-.75.75-.75h2c.41 0 .75.34.75.75s-.34.75-.75.75m4 6h-6c-.41 0-.75-.34-.75-.75s.34-.75.75-.75h6c.41 0 .75.34.75.75s-.34.75-.75.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortDown(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 16.25a.74.74 0 0 1-.53-.22l-7-7A.75.75 0 0 1 5 7.75h14A.75.75 0 0 1 19.53 9l-7 7a.74.74 0 0 1-.53.25m-5.19-7L12 14.44l5.19-5.19Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortNumericAltDown(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.38 19.75a.75.75 0 0 1-.54-.22L5.34 17a.75.75 0 0 1 0-1.06a.77.77 0 0 1 1.07 0l2 2l2-2a.77.77 0 0 1 1.07 0a.75.75 0 0 1 0 1.06l-2.5 2.5a.74.74 0 0 1-.6.25"/><path fill="currentColor" d="M8.38 19.75a.76.76 0 0 1-.76-.75V5a.76.76 0 0 1 .76-.75a.75.75 0 0 1 .74.75v14a.75.75 0 0 1-.74.75m8.74-.5a.75.75 0 0 1-.74-.75v-4.06l-.39.22a.75.75 0 0 1-.73-1.32l.66-.36a1.19 1.19 0 0 1 1.22-.11a1.29 1.29 0 0 1 .74 1.18v4.45a.76.76 0 0 1-.76.75m-.5-10A2.25 2.25 0 1 1 18.88 7a2.24 2.24 0 0 1-2.26 2.25m0-3a.75.75 0 0 0 0 1.5a.75.75 0 1 0 0-1.5"/><path fill="currentColor" d="M16.11 11.25h-.49a.75.75 0 0 1 0-1.5h.49a1.28 1.28 0 0 0 1.25-1.19V7a.75.75 0 0 1 .74-.75a.76.76 0 0 1 .76.75v1.64a2.78 2.78 0 0 1-2.75 2.61"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortNumericAltUp(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.88 8.25a.75.75 0 0 1-.54-.25l-2-2l-2 2a.77.77 0 0 1-1 0a.75.75 0 0 1 0-1l2.5-2.5a.77.77 0 0 1 1.07 0l2.5 2.5a.75.75 0 0 1 0 1.06a.74.74 0 0 1-.53.19"/><path fill="currentColor" d="M8.38 19.75a.76.76 0 0 1-.76-.75V5a.76.76 0 0 1 .76-.75a.75.75 0 0 1 .74.75v14a.75.75 0 0 1-.74.75m8.74-.5a.75.75 0 0 1-.74-.75v-4.06l-.39.22a.75.75 0 0 1-.73-1.32l.66-.36a1.19 1.19 0 0 1 1.22-.11a1.29 1.29 0 0 1 .74 1.18v4.45a.76.76 0 0 1-.76.75m-.5-10A2.25 2.25 0 1 1 18.88 7a2.24 2.24 0 0 1-2.26 2.25m0-3a.75.75 0 0 0 0 1.5a.75.75 0 1 0 0-1.5"/><path fill="currentColor" d="M16.11 11.25h-.49a.75.75 0 0 1 0-1.5h.49a1.28 1.28 0 0 0 1.25-1.19V7a.75.75 0 0 1 .74-.75a.76.76 0 0 1 .76.75v1.64a2.78 2.78 0 0 1-2.75 2.61"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortNumericDown(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.38 19.75a.75.75 0 0 1-.54-.22L5.34 17a.75.75 0 0 1 0-1.06a.77.77 0 0 1 1.07 0l2 2l2-2a.77.77 0 0 1 1.07 0a.75.75 0 0 1 0 1.06l-2.5 2.5a.74.74 0 0 1-.6.25"/><path fill="currentColor" d="M8.38 19.75a.76.76 0 0 1-.76-.75V5a.76.76 0 0 1 .76-.75a.75.75 0 0 1 .74.75v14a.75.75 0 0 1-.74.75m8.74-8.5a.75.75 0 0 1-.74-.75V6.44l-.38.22a.75.75 0 0 1-.73-1.32l.65-.34a1.19 1.19 0 0 1 1.22-.11a1.29 1.29 0 0 1 .74 1.18v4.43a.76.76 0 0 1-.76.75m-.5 6A2.25 2.25 0 1 1 18.88 15a2.24 2.24 0 0 1-2.26 2.25m0-3a.75.75 0 0 0 0 1.5a.75.75 0 1 0 0-1.5"/><path fill="currentColor" d="M16.11 19.25h-.49a.75.75 0 0 1 0-1.5h.49a1.28 1.28 0 0 0 1.25-1.19V15a.75.75 0 0 1 .74-.75a.76.76 0 0 1 .76.75v1.64a2.79 2.79 0 0 1-2.75 2.61"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortNumericUp(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.75 8.25c-.19 0-.38-.07-.53-.22L8.25 6.06L6.28 8.03c-.29.29-.77.29-1.06 0s-.29-.77 0-1.06l2.5-2.5c.29-.29.77-.29 1.06 0l2.5 2.5c.29.29.29.77 0 1.06c-.15.15-.34.22-.53.22"/><path fill="currentColor" d="M8.25 19.75c-.41 0-.75-.34-.75-.75V5c0-.41.34-.75.75-.75S9 4.59 9 5v14c0 .41-.34.75-.75.75m8.75-8.5c-.41 0-.75-.34-.75-.75V6.44l-.39.21a.75.75 0 0 1-.73-1.31l.65-.36c.36-.26.82-.3 1.22-.12c.45.21.74.67.74 1.19v4.45c0 .41-.34.75-.75.75Zm-.5 6c-1.24 0-2.25-1.01-2.25-2.25s1.01-2.25 2.25-2.25s2.25 1.01 2.25 2.25s-1.01 2.25-2.25 2.25m0-3c-.41 0-.75.34-.75.75s.34.75.75.75s.75-.34.75-.75s-.34-.75-.75-.75"/><path fill="currentColor" d="M15.98 19.25h-.48c-.41 0-.75-.34-.75-.75s.34-.75.75-.75h.48c.65 0 1.22-.54 1.25-1.19c.01-.29.01-.64.01-1.06V15c0-.41.34-.75.75-.75s.75.34.75.75v.5c0 .45 0 .82-.02 1.14c-.07 1.44-1.31 2.61-2.75 2.61Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortUp(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 16.25H5a.74.74 0 0 1-.69-.46a.75.75 0 0 1 .16-.79l7-7a.75.75 0 0 1 1.06 0l7 7a.75.75 0 0 1 .16.82a.74.74 0 0 1-.69.43m-12.19-1.5h10.38L12 9.56Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spinner(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21a9 9 0 1 1 6.18-15.55a.75.75 0 0 1 0 1.06a.74.74 0 0 1-1.06 0A7.51 7.51 0 1 0 19.5 12a.75.75 0 0 1 1.5 0a9 9 0 0 1-9 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Star(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.2 20.75a.73.73 0 0 1-.35-.09L12 18.11l-4.85 2.55a.75.75 0 0 1-.79-.05a.76.76 0 0 1-.3-.74l.94-5.4l-3.94-3.82a.75.75 0 0 1-.18-.77a.74.74 0 0 1 .6-.51l5.42-.79l2.43-4.91a.78.78 0 0 1 1.34 0l2.43 4.91l5.42.79a.74.74 0 0 1 .6.51a.75.75 0 0 1-.18.77L17 14.47l.93 5.4a.76.76 0 0 1-.3.74a.79.79 0 0 1-.43.14M12 16.52a.85.85 0 0 1 .35.08l3.85 2l-.73-4.29a.78.78 0 0 1 .21-.67l3.12-3l-4.31-.64a.76.76 0 0 1-.56-.41L12 5.69L10.07 9.6a.76.76 0 0 1-.56.41l-4.31.63l3.12 3a.78.78 0 0 1 .21.67l-.73 4.32l3.85-2a.85.85 0 0 1 .35-.11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarFill(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.12 9.88a.74.74 0 0 0-.6-.51l-5.42-.79l-2.43-4.91a.78.78 0 0 0-1.34 0L8.9 8.58l-5.42.79a.74.74 0 0 0-.6.51a.75.75 0 0 0 .18.77L7 14.47l-.93 5.4a.76.76 0 0 0 .3.74a.75.75 0 0 0 .79.05L12 18.11l4.85 2.55a.73.73 0 0 0 .35.09a.79.79 0 0 0 .44-.14a.76.76 0 0 0 .3-.74l-.94-5.4l3.93-3.82a.75.75 0 0 0 .19-.77"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StepBackward(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 19.75a.77.77 0 0 1-.53-.22l-7-7a.75.75 0 0 1 0-1.06l7-7a.75.75 0 0 1 .82-.16a.74.74 0 0 1 .46.69v14a.74.74 0 0 1-.46.69a.75.75 0 0 1-.29.06M10.06 12l5.19 5.19V6.81Z"/><path fill="currentColor" d="M8 19.75a.76.76 0 0 1-.75-.75V5a.75.75 0 0 1 1.5 0v14a.76.76 0 0 1-.75.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StepBackwardAlt(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.29 4.31a.776.776 0 0 0-.82.16l-6.72 6.72V5c0-.41-.34-.75-.75-.75s-.75.34-.75.75v14c0 .41.34.75.75.75s.75-.34.75-.75v-6.19l6.72 6.72c.14.14.34.22.53.22a.753.753 0 0 0 .75-.75V5c0-.3-.18-.58-.46-.69m-1.04 12.88L10.06 12l5.19-5.19z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StepForward(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 19.75a.75.75 0 0 1-.29-.06a.74.74 0 0 1-.46-.69V5a.74.74 0 0 1 .46-.69a.75.75 0 0 1 .82.16l7 7a.75.75 0 0 1 0 1.06l-7 7a.77.77 0 0 1-.53.22m.75-12.94v10.38L13.94 12Z"/><path fill="currentColor" d="M16 19.75a.76.76 0 0 1-.75-.75V5a.75.75 0 0 1 1.5 0v14a.76.76 0 0 1-.75.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StepForwardAlt(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 4.25c-.41 0-.75.34-.75.75v6.19L8.53 4.47A.751.751 0 0 0 7.25 5v14c0 .3.18.58.46.69a.75.75 0 0 0 .82-.16l6.72-6.72V19c0 .41.34.75.75.75s.75-.34.75-.75V5c0-.41-.34-.75-.75-.75M8.75 17.19V6.81L13.94 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stop(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 19.75H7A2.75 2.75 0 0 1 4.25 17V7A2.75 2.75 0 0 1 7 4.25h10A2.75 2.75 0 0 1 19.75 7v10A2.75 2.75 0 0 1 17 19.75m-10-14A1.25 1.25 0 0 0 5.75 7v10A1.25 1.25 0 0 0 7 18.25h10A1.25 1.25 0 0 0 18.25 17V7A1.25 1.25 0 0 0 17 5.75Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopCircle(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21a9 9 0 1 1 9-9a9 9 0 0 1-9 9m0-16.5a7.5 7.5 0 1 0 7.5 7.5A7.5 7.5 0 0 0 12 4.5"/><rect width="8" height="8" x="8" y="8" fill="currentColor" rx="1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stopwatch(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.75 6c-3.86 0-7 3.14-7 7s3.14 7 7 7s7-3.14 7-7s-3.14-7-7-7m0 12.5c-3.03 0-5.5-2.47-5.5-5.5s2.47-5.5 5.5-5.5s5.5 2.47 5.5 5.5s-2.47 5.5-5.5 5.5M8.5 4.75c0-.41.34-.75.75-.75h5c.41 0 .75.34.75.75s-.34.75-.75.75h-5c-.41 0-.75-.34-.75-.75m4 5.25v3c0 .41-.34.75-.75.75S11 13.41 11 13v-3c0-.41.34-.75.75-.75s.75.34.75.75m6.54-1.73c-.15.15-.34.22-.53.22s-.38-.07-.53-.22l-1.5-1.5c-.29-.29-.29-.77 0-1.06s.77-.29 1.06 0l1.5 1.5c.29.29.29.77 0 1.06"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sun(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 17.75A5.75 5.75 0 1 1 17.75 12A5.76 5.76 0 0 1 12 17.75m0-10A4.25 4.25 0 1 0 16.25 12A4.26 4.26 0 0 0 12 7.75M12 5a.76.76 0 0 1-.75-.75v-1.5a.75.75 0 0 1 1.5 0v1.5A.76.76 0 0 1 12 5m0 17a.76.76 0 0 1-.75-.75v-1.5a.75.75 0 0 1 1.5 0v1.5A.76.76 0 0 1 12 22m9.25-9.25h-1.5a.75.75 0 0 1 0-1.5h1.5a.75.75 0 0 1 0 1.5m-17 0h-1.5a.75.75 0 0 1 0-1.5h1.5a.75.75 0 0 1 0 1.5m2.25-5.5A.74.74 0 0 1 6 7L4.91 6A.75.75 0 1 1 6 4.91L7 6a.75.75 0 0 1 0 1a.74.74 0 0 1-.5.25m12.06 12.06a.74.74 0 0 1-.53-.22L17 18a.75.75 0 0 1 1-1l1.09 1a.75.75 0 0 1 0 1.06a.74.74 0 0 1-.53.25M17.5 7.25A.74.74 0 0 1 17 7a.75.75 0 0 1 0-1l1-1.09A.75.75 0 1 1 19.09 6L18 7a.74.74 0 0 1-.5.25M5.44 19.31a.74.74 0 0 1-.53-.22a.75.75 0 0 1 0-1.06L6 17a.75.75 0 0 1 1 1l-1 1.09a.74.74 0 0 1-.56.22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sync(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.43 4.25a.76.76 0 0 0-.75.75v2.43l-.84-.84a7.24 7.24 0 0 0-12 2.78a.74.74 0 0 0 .46 1a.73.73 0 0 0 .25 0a.76.76 0 0 0 .71-.51a5.63 5.63 0 0 1 1.37-2.2a5.76 5.76 0 0 1 8.13 0l.84.84h-2.41a.75.75 0 0 0 0 1.5h4.24a.74.74 0 0 0 .75-.75V5a.75.75 0 0 0-.75-.75m.25 9.43a.76.76 0 0 0-1 .47a5.63 5.63 0 0 1-1.37 2.2a5.76 5.76 0 0 1-8.13 0l-.84-.84h2.47a.75.75 0 0 0 0-1.5H5.57a.74.74 0 0 0-.75.75V19a.75.75 0 0 0 1.5 0v-2.43l.84.84a7.24 7.24 0 0 0 12-2.78a.74.74 0 0 0-.48-.95"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Table(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 4.25H6A2.75 2.75 0 0 0 3.25 7v10A2.75 2.75 0 0 0 6 19.75h12A2.75 2.75 0 0 0 20.75 17V7A2.75 2.75 0 0 0 18 4.25M19.25 7v4.25h-6.5v-5.5H18A1.25 1.25 0 0 1 19.25 7M6 5.75h5.25v5.5h-6.5V7A1.25 1.25 0 0 1 6 5.75M4.75 17v-4.25h6.5v5.5H6A1.25 1.25 0 0 1 4.75 17M18 18.25h-5.25v-5.5h6.5V17A1.25 1.25 0 0 1 18 18.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tablet(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 3.25H6c-.96 0-1.75.79-1.75 1.75v14c0 .96.79 1.75 1.75 1.75h12c.96 0 1.75-.79 1.75-1.75V5c0-.96-.79-1.75-1.75-1.75M18.25 19c0 .14-.11.25-.25.25H6c-.14 0-.25-.11-.25-.25V5c0-.14.11-.25.25-.25h12c.14 0 .25.11.25.25zM13 16c0 .55-.45 1-1 1s-1-.45-1-1s.45-1 1-1s1 .45 1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tag(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.32 19.98c-.58 0-1.16-.22-1.6-.66l-6.48-6.47a.75.75 0 0 1-.22-.53V4.77c0-.41.34-.75.75-.75h7.54c.2 0 .39.08.53.22l6.48 6.48c.87.88.87 2.31 0 3.19l-5.41 5.41c-.44.44-1.02.66-1.6.66Zm-6.8-7.97l6.26 6.25c.3.29.78.29 1.07 0l5.41-5.41c.29-.29.29-.77 0-1.07L12 5.52H5.52V12ZM8.5 9.75a1.25 1.25 0 1 1 0-2.5a1.25 1.25 0 0 1 0 2.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tags(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.07 10.3l-6-6a.75.75 0 0 0-.53-.22H3c-.41 0-.75.34-.75.75v7.88c0 .2.08.39.22.53l6 5.99a2.3 2.3 0 0 0 1.64.68c.62 0 1.21-.24 1.64-.68l.22-.22c.04.08.08.16.15.22a2.326 2.326 0 0 0 3.29 0l5.65-5.65c.9-.9.9-2.37 0-3.28ZM10.7 18.17c-.16.16-.36.24-.58.24s-.43-.09-.58-.24L3.75 12.4V5.58h6.82l5.78 5.78c.32.32.32.84 0 1.16zm9.31-5.65l-5.65 5.65c-.32.32-.85.32-1.17 0a.704.704 0 0 0-.23-.15l4.44-4.44c.9-.91.9-2.38 0-3.28l-4.72-4.72h1.54L20 11.36c.32.32.32.84 0 1.16ZM8.25 8.5a1.25 1.25 0 1 1-2.5 0a1.25 1.25 0 0 1 2.5 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Telegram(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 4a8 8 0 1 0 8 8a8 8 0 0 0-8-8m3.93 5.48l-1.31 6.19c-.1.44-.36.54-.73.34l-2-1.48l-1 .93a.51.51 0 0 1-.4.2l.14-2l3.7-3.35c.17-.14 0-.22-.24-.08l-4.54 2.85l-2-.62c-.43-.13-.44-.43.09-.63l7.71-3c.38-.11.7.11.58.65"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThLarge(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 11.25H6A2.25 2.25 0 0 1 3.75 9V6A2.25 2.25 0 0 1 6 3.75h3A2.25 2.25 0 0 1 11.25 6v3A2.25 2.25 0 0 1 9 11.25m-3-6a.76.76 0 0 0-.75.75v3a.76.76 0 0 0 .75.75h3A.76.76 0 0 0 9.75 9V6A.76.76 0 0 0 9 5.25Zm3 15H6A2.25 2.25 0 0 1 3.75 18v-3A2.25 2.25 0 0 1 6 12.75h3A2.25 2.25 0 0 1 11.25 15v3A2.25 2.25 0 0 1 9 20.25m-3-6a.76.76 0 0 0-.75.75v3a.76.76 0 0 0 .75.75h3a.76.76 0 0 0 .75-.75v-3a.76.76 0 0 0-.75-.75Zm12-3h-3A2.25 2.25 0 0 1 12.75 9V6A2.25 2.25 0 0 1 15 3.75h3A2.25 2.25 0 0 1 20.25 6v3A2.25 2.25 0 0 1 18 11.25m-3-6a.76.76 0 0 0-.75.75v3a.76.76 0 0 0 .75.75h3a.76.76 0 0 0 .75-.75V6a.76.76 0 0 0-.75-.75Zm3 15h-3A2.25 2.25 0 0 1 12.75 18v-3A2.25 2.25 0 0 1 15 12.75h3A2.25 2.25 0 0 1 20.25 15v3A2.25 2.25 0 0 1 18 20.25m-3-6a.76.76 0 0 0-.75.75v3a.76.76 0 0 0 .75.75h3a.76.76 0 0 0 .75-.75v-3a.76.76 0 0 0-.75-.75Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsDown(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.38 3.25H6.81c-1.09 0-2.02.78-2.21 1.86l-1.31 7.5c-.11.66.07 1.33.49 1.84c.43.51 1.05.8 1.72.8h4.03V18c0 1.52 1.23 2.75 2.83 2.75c.7 0 1.33-.42 1.61-1.07l2.54-5.93h1.88c1.31 0 2.37-1.06 2.37-2.37V5.61c0-1.3-1.06-2.36-2.37-2.36Zm-3.12 9.6l-2.67 6.25c-.04.09-.13.16-.32.16c-.69 0-1.24-.56-1.24-1.25v-4.25H5.5c-.23 0-.43-.09-.57-.26a.724.724 0 0 1-.16-.62l1.31-7.5c.06-.36.37-.62.74-.62h8.44zm3.99-1.47c0 .48-.39.87-.87.87h-1.61v-7.5h1.61c.48 0 .87.39.87.86z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsDownFill(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 5.61v5.77c0 .89-.73 1.62-1.62 1.62h-1.61V4h1.61c.9 0 1.62.72 1.62 1.61M5.34 5.24l-1.32 7.5c-.16.92.54 1.76 1.48 1.76h4.78V18c0 1.1.9 2 1.99 2h.09c.4 0 .76-.24.92-.61L16.01 13V4h-9.2c-.73 0-1.35.52-1.48 1.24Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsUp(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.22 9.55c-.43-.51-1.05-.8-1.72-.8h-4.03V6c0-1.52-1.23-2.75-2.83-2.75c-.7 0-1.33.42-1.61 1.07l-2.54 5.93H5.62c-1.31 0-2.37 1.06-2.37 2.37v5.77c0 1.3 1.07 2.36 2.37 2.36h11.56c1.09 0 2.02-.78 2.21-1.86l1.32-7.5c.11-.66-.07-1.33-.5-1.84Zm-14.6 9.7c-.48 0-.87-.39-.87-.86v-5.77c0-.48.39-.87.87-.87h1.61v7.5zm12.3-.62c-.06.36-.37.62-.74.62H8.74v-8.1l2.67-6.25c.04-.09.13-.16.32-.16c.69 0 1.24.56 1.24 1.25v4.25h5.53c.23 0 .43.09.57.26c.14.17.2.39.16.62l-1.32 7.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsUpFill(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.24 11v9H5.63c-.9 0-1.62-.72-1.62-1.61v-5.77c0-.89.73-1.62 1.62-1.62zM18.5 9.5h-4.78V6c0-1.1-.9-2-1.99-2h-.09c-.4 0-.76.24-.92.61L7.99 11v9h9.2c.73 0 1.35-.52 1.48-1.24l1.32-7.5c.16-.92-.54-1.76-1.48-1.76Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ticket(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 18.75H5A1.76 1.76 0 0 1 3.25 17v-2.5a.76.76 0 0 1 .75-.75a1.75 1.75 0 0 0 0-3.5a.76.76 0 0 1-.75-.75V7A1.76 1.76 0 0 1 5 5.25h14A1.76 1.76 0 0 1 20.75 7v2.5a.76.76 0 0 1-.75.75a1.75 1.75 0 0 0 0 3.5a.76.76 0 0 1 .75.75V17A1.76 1.76 0 0 1 19 18.75M4.75 15.16V17a.25.25 0 0 0 .25.25h14a.25.25 0 0 0 .25-.25v-1.84a3.25 3.25 0 0 1 0-6.32V7a.25.25 0 0 0-.25-.25H5a.25.25 0 0 0-.25.25v1.84a3.25 3.25 0 0 1 0 6.32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Times(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.06 12l4.42-4.42a.75.75 0 1 0-1.06-1.06L12 10.94L7.58 6.52a.75.75 0 0 0-1.06 1.06L10.94 12l-4.42 4.42a.75.75 0 0 0 0 1.06a.75.75 0 0 0 1.06 0L12 13.06l4.42 4.42a.75.75 0 0 0 1.06 0a.75.75 0 0 0 0-1.06Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimesCircle(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21a9 9 0 1 1 9-9a9 9 0 0 1-9 9m0-16.5a7.5 7.5 0 1 0 7.5 7.5A7.5 7.5 0 0 0 12 4.5"/><path fill="currentColor" d="M9 15.75a.74.74 0 0 1-.53-.22a.75.75 0 0 1 0-1.06l6-6a.75.75 0 0 1 1.06 1.06l-6 6a.74.74 0 0 1-.53.22"/><path fill="currentColor" d="M15 15.75a.74.74 0 0 1-.53-.22l-6-6a.75.75 0 0 1 1.06-1.06l6 6a.75.75 0 0 1 0 1.06a.74.74 0 0 1-.53.22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trash(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 8.7H4a.75.75 0 1 1 0-1.5h16a.75.75 0 0 1 0 1.5"/><path fill="currentColor" d="M16.44 20.75H7.56A2.4 2.4 0 0 1 5 18.49V8a.75.75 0 0 1 1.5 0v10.49c0 .41.47.76 1 .76h8.88c.56 0 1-.35 1-.76V8A.75.75 0 1 1 19 8v10.49a2.4 2.4 0 0 1-2.56 2.26m.12-13a.74.74 0 0 1-.75-.75V5.51c0-.41-.48-.76-1-.76H9.22c-.55 0-1 .35-1 .76V7a.75.75 0 1 1-1.5 0V5.51a2.41 2.41 0 0 1 2.5-2.26h5.56a2.41 2.41 0 0 1 2.53 2.26V7a.75.75 0 0 1-.75.76Z"/><path fill="currentColor" d="M10.22 17a.76.76 0 0 1-.75-.75v-4.53a.75.75 0 0 1 1.5 0v4.52a.75.75 0 0 1-.75.76m3.56 0a.75.75 0 0 1-.75-.75v-4.53a.75.75 0 0 1 1.5 0v4.52a.76.76 0 0 1-.75.76"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Truck(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.5 6.75h-2V6c0-.96-.79-1.75-1.75-1.75H4.5c-.96 0-1.75.79-1.75 1.75v9c0 .96.79 1.75 1.75 1.75h.28c.13 1.67 1.52 3 3.22 3s3.09-1.33 3.22-3h1.55c.13 1.67 1.52 3 3.22 3s3.09-1.33 3.22-3h.28c.96 0 1.75-.79 1.75-1.75v-4.5c0-2.07-1.68-3.75-3.75-3.75ZM4.25 15V6c0-.14.11-.25.25-.25h9.25c.14 0 .25.11.25.25v7.94s-.05.06-.08.08c-.15.13-.29.27-.42.42c-.04.05-.08.09-.11.14a2.9 2.9 0 0 0-.38.66h-2v-.01a2.91 2.91 0 0 0-.27-.49c-.03-.04-.06-.08-.09-.13c-.08-.12-.17-.22-.27-.33c-.04-.04-.07-.08-.11-.12c-.13-.12-.26-.23-.4-.33c-.02-.01-.03-.03-.05-.04c-.15-.1-.31-.18-.47-.26c-.06-.03-.12-.05-.18-.07c-.13-.05-.25-.09-.38-.12c-.06-.01-.12-.03-.18-.04a2.88 2.88 0 0 0-1.18 0c-.06.01-.12.03-.18.04c-.13.03-.26.07-.38.12c-.06.02-.12.04-.18.07c-.16.07-.32.16-.47.26c-.02.01-.03.02-.05.04c-.14.1-.28.21-.4.33c-.04.04-.08.08-.11.12c-.1.1-.19.21-.27.33c-.03.04-.06.08-.09.13c-.1.16-.19.32-.27.49v.01h-.5c-.14 0-.25-.11-.25-.25ZM8 18.25c-.96 0-1.75-.79-1.75-1.75a1.69 1.69 0 0 1 .24-.85c.01-.03.03-.05.04-.08c.09-.14.2-.27.33-.38c.03-.02.05-.04.08-.06c.13-.1.28-.19.44-.25c.02 0 .05-.02.08-.02c.17-.06.36-.1.55-.1s.37.04.55.1c.03 0 .05.01.08.02c.16.06.31.15.44.25c.03.02.05.04.08.06c.13.11.24.24.33.38c.02.02.03.05.04.08c.09.16.16.32.2.5c.02.11.04.23.04.35c0 .96-.79 1.75-1.75 1.75Zm8 0c-.96 0-1.75-.79-1.75-1.75c0-.12.01-.24.04-.35c.03-.12.06-.24.11-.35v-.02c.05-.1.1-.2.17-.29c.01-.02.03-.03.04-.05a2 2 0 0 1 .19-.22c.02-.02.04-.03.06-.05c.07-.06.15-.12.23-.17c.02-.01.05-.03.08-.04c.09-.05.19-.09.29-.12c.04-.01.08-.03.13-.04c.13-.03.27-.05.41-.05c.19 0 .37.04.55.1c.03 0 .05.01.08.02c.16.06.31.15.44.25c.03.02.05.04.08.06c.13.11.24.24.33.38c.02.02.03.05.04.08c.09.16.16.32.2.5c.02.11.04.23.04.35c0 .96-.79 1.75-1.75 1.75Zm-.5-7v-3h2c1.24 0 2.25 1.01 2.25 2.25v.75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Twitter(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.83 8v.52A11.41 11.41 0 0 1 8.35 20a11.41 11.41 0 0 1-6.2-1.81h1a8.09 8.09 0 0 0 5-1.73a4 4 0 0 1-3.78-2.8a4.81 4.81 0 0 0 .77.06a4.66 4.66 0 0 0 1.06-.13A4 4 0 0 1 3 9.67a4.13 4.13 0 0 0 1.82.51A4.06 4.06 0 0 1 3 6.77a4 4 0 0 1 .54-2A11.47 11.47 0 0 0 11.85 9a4.71 4.71 0 0 1-.1-.92a4 4 0 0 1 7-2.77a7.93 7.93 0 0 0 2.56-1a4 4 0 0 1-1.78 2.22a7.94 7.94 0 0 0 2.33-.62a8.91 8.91 0 0 1-2 2.09Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Undo(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 20.75a7.26 7.26 0 0 1-7.25-7.25a.75.75 0 0 1 1.5 0A5.75 5.75 0 1 0 12 7.75H9.5a.75.75 0 0 1 0-1.5H12a7.25 7.25 0 0 1 0 14.5"/><path fill="currentColor" d="M12 10.75a.74.74 0 0 1-.53-.22l-3-3a.75.75 0 0 1 0-1.06l3-3a.75.75 0 1 1 1.06 1.06L10.06 7l2.47 2.47a.75.75 0 0 1 0 1.06a.74.74 0 0 1-.53.22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unlock(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 10.25H8.75V8a3.25 3.25 0 0 1 6.5 0a.75.75 0 0 0 1.5 0a4.75 4.75 0 0 0-9.5 0v2.25H7A2.75 2.75 0 0 0 4.25 13v5A2.75 2.75 0 0 0 7 20.75h10A2.75 2.75 0 0 0 19.75 18v-5A2.75 2.75 0 0 0 17 10.25M18.25 18A1.25 1.25 0 0 1 17 19.25H7A1.25 1.25 0 0 1 5.75 18v-5A1.25 1.25 0 0 1 7 11.75h10A1.25 1.25 0 0 1 18.25 13Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Upload(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.22 20.75H5.78A2.64 2.64 0 0 1 3.25 18v-3a.75.75 0 0 1 1.5 0v3a1.16 1.16 0 0 0 1 1.25h12.47a1.16 1.16 0 0 0 1-1.25v-3a.75.75 0 0 1 1.5 0v3a2.64 2.64 0 0 1-2.5 2.75M16 8.75a.74.74 0 0 1-.53-.22L12 5.06L8.53 8.53a.75.75 0 0 1-1.06-1.06l4-4a.75.75 0 0 1 1.06 0l4 4a.75.75 0 0 1 0 1.06a.74.74 0 0 1-.53.22"/><path fill="currentColor" d="M12 15.75a.76.76 0 0 1-.75-.75V4a.75.75 0 0 1 1.5 0v11a.76.76 0 0 1-.75.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func User(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 12.25a3.75 3.75 0 1 1 3.75-3.75A3.75 3.75 0 0 1 12 12.25m0-6a2.25 2.25 0 1 0 2.25 2.25A2.25 2.25 0 0 0 12 6.25m7 13a.76.76 0 0 1-.75-.75c0-1.95-1.06-3.25-6.25-3.25s-6.25 1.3-6.25 3.25a.75.75 0 0 1-1.5 0c0-4.75 5.43-4.75 7.75-4.75s7.75 0 7.75 4.75a.76.76 0 0 1-.75.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserEdit(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.56 11.87a3.75 3.75 0 1 1 3.75-3.75a3.76 3.76 0 0 1-3.75 3.75m0-6a2.25 2.25 0 1 0 2.25 2.25a2.25 2.25 0 0 0-2.25-2.25m-7 13a.75.75 0 0 1-.75-.75c0-4.75 5.43-4.75 7.75-4.75c.72 0 1.36 0 1.94.07a.75.75 0 0 1 .69.8a.76.76 0 0 1-.81.69c-.54 0-1.14-.06-1.82-.06c-5.18 0-6.25 1.3-6.25 3.25a.74.74 0 0 1-.75.75m9.11.76a.75.75 0 0 1-.53-.22a.72.72 0 0 1-.22-.59l.16-1.92a.75.75 0 0 1 .21-.47l5.52-5.52a2.06 2.06 0 0 1 2.8 0a2 2 0 0 1 .58 1.44a1.86 1.86 0 0 1-.53 1.33l-5.52 5.52a.74.74 0 0 1-.46.22l-1.94.18Zm.88-2.34l-.06.76l.78-.07l5.33-5.33a.39.39 0 0 0 .09-.27a.59.59 0 0 0-.14-.38a.57.57 0 0 0-.68 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserMinus(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 12.75A3.75 3.75 0 1 1 15.75 9A3.75 3.75 0 0 1 12 12.75m0-6A2.25 2.25 0 1 0 14.25 9A2.25 2.25 0 0 0 12 6.75m7 13a.76.76 0 0 1-.75-.75c0-1.95-1.06-3.25-6.25-3.25S5.75 17.05 5.75 19a.75.75 0 0 1-1.5 0c0-4.75 5.43-4.75 7.75-4.75s7.75 0 7.75 4.75a.76.76 0 0 1-.75.75m1.75-7h-3.5a.75.75 0 0 1 0-1.5h3.5a.75.75 0 0 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserPlus(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 12.75A3.75 3.75 0 1 1 15.75 9A3.75 3.75 0 0 1 12 12.75m0-6A2.25 2.25 0 1 0 14.25 9A2.25 2.25 0 0 0 12 6.75m7 13a.76.76 0 0 1-.75-.75c0-1.95-1.06-3.25-6.25-3.25S5.75 17.05 5.75 19a.75.75 0 0 1-1.5 0c0-4.75 5.43-4.75 7.75-4.75s7.75 0 7.75 4.75a.76.76 0 0 1-.75.75m0-5.25a.76.76 0 0 1-.75-.75v-3.5a.75.75 0 0 1 1.5 0v3.5a.76.76 0 0 1-.75.75"/><path fill="currentColor" d="M20.75 12.75h-3.5a.75.75 0 0 1 0-1.5h3.5a.75.75 0 0 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Users(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 12.25a3.75 3.75 0 1 1 3.75-3.75A3.75 3.75 0 0 1 14 12.25m0-6a2.25 2.25 0 1 0 2.25 2.25A2.25 2.25 0 0 0 14 6.25m7 13a.76.76 0 0 1-.75-.75c0-1.95-1.06-3.25-6.25-3.25s-6.25 1.3-6.25 3.25a.75.75 0 0 1-1.5 0c0-4.75 5.43-4.75 7.75-4.75s7.75 0 7.75 4.75a.76.76 0 0 1-.75.75M8.32 13.06H8a3 3 0 1 1 .58-6a.75.75 0 1 1-.15 1.49a1.46 1.46 0 0 0-1.09.34a1.47 1.47 0 0 0-.54 1a1.49 1.49 0 0 0 1.35 1.64a1.53 1.53 0 0 0 .93-.22a.75.75 0 0 1 .79 1.28a3 3 0 0 1-1.55.47M3 18.5a.76.76 0 0 1-.75-.75c0-2.7.72-4.5 4.25-4.5a.75.75 0 0 1 0 1.5c-2.35 0-2.75.75-2.75 3a.76.76 0 0 1-.75.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Verified(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.01 21c-.49 0-.95-.23-1.33-.43c-.24-.12-.53-.27-.68-.27s-.45.15-.68.27c-.48.25-1.08.55-1.72.38c-.66-.17-1.02-.75-1.32-1.21c-.13-.21-.31-.49-.43-.56c-.12-.07-.44-.08-.71-.1c-.54-.03-1.21-.06-1.69-.53c-.48-.49-.51-1.16-.54-1.7c-.01-.26-.03-.59-.1-.71c-.06-.11-.35-.29-.55-.42c-.46-.3-1.03-.67-1.21-1.32c-.17-.64.13-1.24.38-1.72c.12-.24.27-.53.27-.68s-.15-.45-.27-.68c-.25-.48-.55-1.08-.38-1.72c.17-.66.75-1.02 1.21-1.32c.21-.13.49-.31.56-.43c.07-.12.08-.44.1-.71c.03-.54.06-1.21.53-1.69c.49-.48 1.16-.51 1.7-.54c.26-.01.59-.03.71-.1c.11-.06.29-.35.42-.55c.3-.46.67-1.03 1.32-1.21c.64-.17 1.24.13 1.72.38c.24.12.53.27.68.27s.45-.15.68-.27c.48-.24 1.08-.55 1.72-.38c.66.17 1.02.75 1.32 1.21c.13.21.31.49.43.56c.12.07.44.08.71.1c.54.03 1.21.06 1.69.53c.48.49.51 1.16.54 1.7c.01.26.03.59.1.71c.06.11.35.29.55.42c.46.3 1.03.67 1.21 1.32c.17.64-.13 1.24-.38 1.72c-.12.24-.27.53-.27.68s.15.45.27.68c.25.48.55 1.08.38 1.72c-.17.66-.75 1.02-1.21 1.32c-.21.13-.49.31-.56.43c-.07.12-.08.44-.1.71c-.03.54-.06 1.21-.53 1.69c-.49.48-1.16.51-1.7.54c-.26.01-.59.03-.71.1c-.11.06-.29.35-.42.55c-.3.46-.67 1.03-1.32 1.21c-.13.04-.26.05-.39.05M10 4.5h-.01c-.1.04-.33.38-.44.57c-.24.37-.51.79-.94 1.04c-.44.25-.94.28-1.39.3c-.22.01-.63.03-.72.1c-.06.08-.08.48-.09.7c-.02.45-.05.95-.3 1.38c-.25.43-.67.7-1.04.94c-.19.12-.53.34-.57.45c-.01.11.16.45.26.65c.2.4.44.85.44 1.36s-.23.96-.44 1.36c-.1.2-.28.54-.26.65c.04.1.38.33.57.44c.37.24.79.51 1.04.94c.25.44.28.94.3 1.39c.01.22.03.63.1.72c.08.06.48.08.7.09c.45.02.95.05 1.38.3c.43.25.7.67.94 1.04c.12.19.34.53.45.57c.1.03.45-.16.65-.26c.4-.2.85-.44 1.36-.44s.96.23 1.36.44c.2.1.52.29.65.26c.1-.04.33-.38.44-.57c.24-.37.51-.79.94-1.04c.44-.25.94-.28 1.39-.3c.22-.01.63-.03.72-.1c.06-.08.08-.48.09-.7c.02-.45.05-.95.3-1.38c.25-.43.67-.7 1.04-.94c.19-.12.53-.34.57-.45c.01-.11-.16-.45-.26-.65c-.2-.4-.44-.85-.44-1.36s.23-.96.44-1.36c.1-.2.27-.54.26-.65c-.04-.11-.39-.33-.57-.45c-.37-.24-.79-.51-1.04-.94c-.25-.44-.28-.94-.3-1.39c-.01-.22-.03-.63-.1-.72c-.08-.06-.48-.08-.7-.09c-.45-.02-.95-.05-1.38-.3c-.43-.25-.7-.67-.94-1.04c-.12-.19-.34-.53-.45-.57c-.1-.03-.45.16-.65.26c-.4.2-.85.44-1.36.44s-.96-.23-1.36-.44c-.2-.1-.52-.26-.64-.26Zm.49 11.01a.75.75 0 0 1-.53-.22l-2.51-2.51c-.29-.29-.29-.77 0-1.06s.77-.29 1.06 0l1.98 1.98l4.99-4.99c.29-.29.77-.29 1.06 0s.29.77 0 1.06l-5.52 5.52a.75.75 0 0 1-.53.22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Video(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 18.75H6A2.75 2.75 0 0 1 3.25 16V8A2.75 2.75 0 0 1 6 5.25h7A2.75 2.75 0 0 1 15.75 8v8A2.75 2.75 0 0 1 13 18.75m-7-12A1.25 1.25 0 0 0 4.75 8v8A1.25 1.25 0 0 0 6 17.25h7A1.25 1.25 0 0 0 14.25 16V8A1.25 1.25 0 0 0 13 6.75Z"/><path fill="currentColor" d="M20 16.75a.79.79 0 0 1-.39-.11l-5-3a.75.75 0 0 1-.36-.64v-2a.75.75 0 0 1 .36-.64l5-3a.74.74 0 0 1 .76 0a.75.75 0 0 1 .38.65v8a.75.75 0 0 1-.38.65a.71.71 0 0 1-.37.09m-4.25-4.17l3.5 2.1V9.32l-3.5 2.1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vimeo(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.4 4H5.6A1.6 1.6 0 0 0 4 5.6v12.8A1.6 1.6 0 0 0 5.6 20h12.8a1.6 1.6 0 0 0 1.6-1.6V5.6A1.6 1.6 0 0 0 18.4 4m-.94 5.31q-.08 1.69-2.35 4.63c-1.57 2-2.91 3.06-4 3.06c-.68 0-1.24-.62-1.71-1.87c-.91-3.33-1.3-5.28-2-5.28a3.53 3.53 0 0 0-.91.54l-.54-.7c1.33-1.17 2.6-2.47 3.4-2.54S10.76 7.68 11 9c.74 4.69 1.07 5.4 2.41 3.27a4.63 4.63 0 0 0 .78-1.74c.13-1.19-.92-1.11-1.64-.8q.83-2.8 3.23-2.73q1.78.06 1.68 2.31"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeDown(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 19.75a.81.81 0 0 1-.47-.16l-4.79-3.84H5a.76.76 0 0 1-.75-.75V9A.76.76 0 0 1 5 8.25h4.74l4.79-3.84a.75.75 0 0 1 1.22.59v14a.76.76 0 0 1-.43.68a.71.71 0 0 1-.32.07m-9.25-5.5H10a.78.78 0 0 1 .47.16l3.78 3V6.56l-3.78 3a.78.78 0 0 1-.47.16H5.75Zm12.36 1.13a.78.78 0 0 1-.45-.15a.75.75 0 0 1-.15-1.05a3.6 3.6 0 0 0 0-4.36a.75.75 0 1 1 1.2-.9a5.07 5.07 0 0 1 0 6.16a.77.77 0 0 1-.6.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeOff(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 19.75a.81.81 0 0 1-.47-.16l-4.79-3.84H7a.76.76 0 0 1-.75-.75V9A.76.76 0 0 1 7 8.25h4.74l4.79-3.84a.75.75 0 0 1 1.22.59v14a.77.77 0 0 1-.42.68a.78.78 0 0 1-.33.07m-9.25-5.5H12a.78.78 0 0 1 .47.16l3.78 3V6.56l-3.78 3a.78.78 0 0 1-.47.16H7.75Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeUp(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 19.75a.81.81 0 0 1-.47-.16l-4.79-3.84H3a.76.76 0 0 1-.75-.75V9A.76.76 0 0 1 3 8.25h4.74l4.79-3.84a.75.75 0 0 1 1.22.59v14a.76.76 0 0 1-.43.68a.71.71 0 0 1-.32.07m-9.25-5.5H8a.78.78 0 0 1 .47.16l3.78 3V6.56l-3.78 3a.78.78 0 0 1-.47.19H3.75Zm14.71 3.82a.76.76 0 0 1-.49-.18a.74.74 0 0 1-.07-1.06a7.24 7.24 0 0 0 0-9.66a.75.75 0 1 1 1.12-1a8.7 8.7 0 0 1 0 11.64a.72.72 0 0 1-.56.26"/><path fill="currentColor" d="M16.11 15.38a.78.78 0 0 1-.45-.15a.75.75 0 0 1-.15-1.05a3.6 3.6 0 0 0 0-4.36a.75.75 0 0 1 1.2-.9a5.07 5.07 0 0 1 0 6.16a.77.77 0 0 1-.6.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wallet(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 7.25h-.25V5A1.76 1.76 0 0 0 17 3.25a.67.67 0 0 0-.24 0l-11.9 4h-.27l-.17.06h-.14l-.16.09l-.12.17l-.14.12l-.11.1l-.12.15a.39.39 0 0 0-.08.1a1.62 1.62 0 0 0-.1.18l-.06.11a1.87 1.87 0 0 0-.07.22a.45.45 0 0 1 0 .11a1.87 1.87 0 0 0 0 .34v10A1.76 1.76 0 0 0 5 20.75h14A1.76 1.76 0 0 0 20.75 19V9A1.76 1.76 0 0 0 19 7.25m-1.92-2.49a.26.26 0 0 1 .17.24v2.25H9.62ZM19.25 19a.25.25 0 0 1-.25.25H5a.25.25 0 0 1-.25-.25V9A.25.25 0 0 1 5 8.75h14a.25.25 0 0 1 .25.25Z"/><circle cx="16.5" cy="14" r="1.25" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Whatsapp(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.6 6.32A7.85 7.85 0 0 0 12 4a7.94 7.94 0 0 0-6.88 11.89L4 20l4.2-1.1a7.93 7.93 0 0 0 3.79 1a8 8 0 0 0 8-7.93a8 8 0 0 0-2.39-5.65M12 18.53a6.58 6.58 0 0 1-3.36-.92l-.24-.15l-2.49.66l.66-2.43l-.16-.25a6.6 6.6 0 0 1 10.25-8.17a6.65 6.65 0 0 1 2 4.66a6.66 6.66 0 0 1-6.66 6.6m3.61-4.94c-.2-.1-1.17-.58-1.35-.64s-.32-.1-.45.1a9 9 0 0 1-.63.77c-.11.14-.23.15-.43 0a5.33 5.33 0 0 1-2.69-2.35c-.21-.35.2-.33.58-1.08a.38.38 0 0 0 0-.35c0-.1-.45-1.08-.61-1.47s-.32-.33-.45-.34h-.39a.71.71 0 0 0-.53.25A2.19 2.19 0 0 0 8 10.17a3.82 3.82 0 0 0 .81 2.05a8.89 8.89 0 0 0 3.39 3a3.85 3.85 0 0 0 2.38.5a2 2 0 0 0 1.33-.94a1.62 1.62 0 0 0 .12-.94c-.09-.1-.22-.15-.42-.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wifi(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.35 12.79a.72.72 0 0 1-.49-.19a7.24 7.24 0 0 0-9.66 0a.75.75 0 0 1-1.06-.06a.76.76 0 0 1 .07-1.06a8.7 8.7 0 0 1 11.64 0a.76.76 0 0 1 .07 1.06a.79.79 0 0 1-.57.25"/><path fill="currentColor" d="M20 10a.76.76 0 0 1-.51-.2a10.87 10.87 0 0 0-15 0a.75.75 0 1 1-1-1.1a12.36 12.36 0 0 1 17 0A.75.75 0 0 1 20 10M9.38 15.64a.75.75 0 0 1-.6-.3a.74.74 0 0 1 .15-1.05a5.06 5.06 0 0 1 6.15 0a.75.75 0 0 1 .15 1.05a.76.76 0 0 1-1.05.15a3.59 3.59 0 0 0-4.35 0a.78.78 0 0 1-.45.15M12 18.75a.75.75 0 1 1 0-1.5a.75.75 0 0 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowMaximize(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 20.75h-6a.75.75 0 0 1 0-1.5h6A1.25 1.25 0 0 0 19.25 18V6A1.25 1.25 0 0 0 18 4.75H6A1.25 1.25 0 0 0 4.75 6v6a.75.75 0 0 1-1.5 0V6A2.75 2.75 0 0 1 6 3.25h12A2.75 2.75 0 0 1 20.75 6v12A2.75 2.75 0 0 1 18 20.75"/><path fill="currentColor" d="M16 12.75a.76.76 0 0 1-.75-.75V8.75H12a.75.75 0 0 1 0-1.5h4a.76.76 0 0 1 .75.75v4a.76.76 0 0 1-.75.75"/><path fill="currentColor" d="M11.5 13.25A.74.74 0 0 1 11 13a.75.75 0 0 1 0-1l4.5-4.5a.75.75 0 0 1 1.06 1.06L12 13a.74.74 0 0 1-.5.25M8 20.75H5A1.76 1.76 0 0 1 3.25 19v-3A1.76 1.76 0 0 1 5 14.25h3A1.76 1.76 0 0 1 9.75 16v3A1.76 1.76 0 0 1 8 20.75m-3-5a.25.25 0 0 0-.25.25v3a.25.25 0 0 0 .25.25h3a.25.25 0 0 0 .25-.25v-3a.25.25 0 0 0-.25-.25Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowMinimize(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 3.25H6A2.75 2.75 0 0 0 3.25 6v6a.75.75 0 0 0 1.5 0V6A1.25 1.25 0 0 1 6 4.75h12A1.25 1.25 0 0 1 19.25 6v12A1.25 1.25 0 0 1 18 19.25h-6a.75.75 0 0 0 0 1.5h6A2.75 2.75 0 0 0 20.75 18V6A2.75 2.75 0 0 0 18 3.25"/><path fill="currentColor" d="M11.21 13.19a.75.75 0 0 0 .29.06h4a.75.75 0 0 0 0-1.5h-2.19l3.22-3.22a.75.75 0 0 0-1.06-1.06l-3.22 3.22V8.5a.75.75 0 0 0-1.5 0v4a.75.75 0 0 0 .06.29a.71.71 0 0 0 .4.4M8 14.25H5A1.76 1.76 0 0 0 3.25 16v3A1.76 1.76 0 0 0 5 20.75h3A1.76 1.76 0 0 0 9.75 19v-3A1.76 1.76 0 0 0 8 14.25M8.25 19a.25.25 0 0 1-.25.25H5a.25.25 0 0 1-.25-.25v-3a.25.25 0 0 1 .25-.25h3a.25.25 0 0 1 .25.25Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wrench(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.12 20.75c-.76 0-1.48-.3-2.03-.84a2.86 2.86 0 0 1 0-4.05l5.51-5.51c-.5-1.94.04-4.03 1.46-5.45a5.667 5.667 0 0 1 5.48-1.46c.26.07.46.27.53.53s0 .53-.19.72l-2.45 2.45l.52 1.91l1.91.52l2.45-2.45c.19-.19.47-.26.72-.19c.26.07.46.27.53.53c.53 1.95-.02 4.05-1.46 5.48c-1.42 1.42-3.51 1.96-5.45 1.46l-5.51 5.51c-.54.54-1.26.84-2.02.84m8.56-15.98c-.96.08-1.87.5-2.57 1.2c-1.14 1.14-1.51 2.81-.96 4.35c.1.27.03.58-.18.78l-5.83 5.83c-.53.53-.53 1.4 0 1.93c.26.26.6.4.97.4c.36 0 .71-.14.96-.4l5.83-5.83c.21-.21.51-.27.78-.18c1.54.54 3.21.18 4.35-.96c.7-.7 1.11-1.61 1.2-2.57l-1.63 1.63c-.19.19-.47.26-.73.19l-2.74-.75a.747.747 0 0 1-.53-.53l-.75-2.74c-.07-.26 0-.54.19-.73l1.63-1.63Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Youtube(children ...ElementRenderer) *PrimeIcon {
	return &PrimeIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.67 8.14a2 2 0 0 0-1.42-1.43A48.44 48.44 0 0 0 12 6.38a48.44 48.44 0 0 0-6.25.33a2 2 0 0 0-1.42 1.43A21.27 21.27 0 0 0 4 12a21.42 21.42 0 0 0 .33 3.88a2 2 0 0 0 1.42 1.4a48.44 48.44 0 0 0 6.25.33a48.44 48.44 0 0 0 6.25-.33a2 2 0 0 0 1.42-1.4A21.42 21.42 0 0 0 20 12a21.27 21.27 0 0 0-.33-3.86m-9.31 6.25V9.63L14.55 12l-4.19 2.38Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
