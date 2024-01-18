package il

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = "0.0.1"
	hAttr          = "1em"
	viewbox        = "0 0 750 850"
	fill           = "currentColor"
)

type IlIcon struct {
	*SVGSVGElement
}

type IlIconFn func(children ...ElementRenderer) *IlIcon

var IconLookup = map[string]IlIconFn{
	"addUser":         AddUser,
	"arrowDown":       ArrowDown,
	"arrowLeft":       ArrowLeft,
	"arrowRight":      ArrowRight,
	"arrowUp":         ArrowUp,
	"attachment":      Attachment,
	"basket":          Basket,
	"behance":         Behance,
	"bell":            Bell,
	"book":            Book,
	"box":             Box,
	"brightness":      Brightness,
	"bucket":          Bucket,
	"calendar":        Calendar,
	"camera":          Camera,
	"card":            Card,
	"cart":            Cart,
	"clock":           Clock,
	"cloud":           Cloud,
	"cog":             Cog,
	"comment":         Comment,
	"compass":         Compass,
	"contrast":        Contrast,
	"controls":        Controls,
	"conversation":    Conversation,
	"cup":             Cup,
	"dashboard":       Dashboard,
	"dialog":          Dialog,
	"dribbble":        Dribbble,
	"drop":            Drop,
	"dropbox":         Dropbox,
	"ellipsis":        Ellipsis,
	"email":           Email,
	"envelope":        Envelope,
	"eye":             Eye,
	"facebook":        Facebook,
	"file":            File,
	"flag":            Flag,
	"folder":          Folder,
	"github":          Github,
	"googlePlus":      GooglePlus,
	"grid":            Grid,
	"heart":           Heart,
	"house":           House,
	"image":           Image,
	"inbox":           Inbox,
	"instagram":       Instagram,
	"layers":          Layers,
	"linkedin":        Linkedin,
	"location":        Location,
	"lock":            Lock,
	"market":          Market,
	"menu":            Menu,
	"mic":             Mic,
	"mobile":          Mobile,
	"money":           Money,
	"moon":            Moon,
	"music":           Music,
	"notification":    Notification,
	"paypal":          Paypal,
	"pencil":          Pencil,
	"pie":             Pie,
	"pin":             Pin,
	"refresh":         Refresh,
	"ribbon":          Ribbon,
	"search":          Search,
	"select":          Select,
	"smallArrowDown":  SmallArrowDown,
	"smallArrowLeft":  SmallArrowLeft,
	"smallArrowRight": SmallArrowRight,
	"smallArrowUp":    SmallArrowUp,
	"tablet":          Tablet,
	"tag":             Tag,
	"thumbs":          Thumbs,
	"triangleDown":    TriangleDown,
	"triangleUp":      TriangleUp,
	"twitter":         Twitter,
	"unlock":          Unlock,
	"url":             Url,
	"user":            User,
	"users":           Users,
	"videocamera":     Videocamera,
	"world":           World,
	"youtube":         Youtube,
}

func AddUser(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M535 518q27 13 42 34t21 44t8 43t0 31t-8 16q-14 12-34 21t-44 16t-49 11l-49 8q-56 7-118 8q-63-1-119-8q-24-3-49-8t-48-11t-44-16t-35-21q-6-5-7.5-15.5T1 639t8-43t21-44t42-34l62-28q29-13 52-28t35-35t13-48v-6q-17-14-25-33l-23-54q-18-20-27-37q-8-14-8-26t17-12q-6-38-3-68q2-25 12-46t36-21q5-19 17-34q11-13 31-24t53-10q30 0 49 13t33 30t23 33t20 23q4 2 7 5t0 8q-4 14-8 41t1 50q18 1 17 12t-10 26q-10 17-28 37q-7 18-11 31t-9 24t-10 17t-15 15v6q0 29 13 48t36 35t51 28t62 28M663 78q11 0 11 11v23q0 12-11 12h-58v58q0 5-3 8t-9 4h-23q-5 0-8-4t-4-8v-58h-57q-12 0-12-12V89q0-11 12-11h57V20q0-5 4-9t8-3h23q12 0 12 12v58z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDown(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M622 106L311 417L0 106l65-65l246 245L556 41z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeft(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M311 630L0 319L311 8l65 66l-245 245l245 245z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRight(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m65 8l311 311L65 630L0 564l245-245L0 74z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUp(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 352L311 41l311 311l-66 65l-245-245L65 417z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Attachment(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M440 240q10 0 17 6t6 17v255q0 50-19 93t-54 75t-78 49t-96 14q-46-3-85-24t-69-53t-45-74t-17-88V186q0-79 34-125t87-52q31-4 60 5t50 29t35 46t12 58v348q0 10-6 16t-17 7h-47q-9 0-16-7t-7-16V150q0-17-10-32t-28-17q-22-3-38 11t-16 35v366q0 27 9 51t25 43t38 33t48 16q32 4 61-6t50-28t35-46t12-58V263q0-10 7-17t16-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Basket(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M718 191q10 0 17 7t6 17v92H0v-92q0-10 7-17t16-7h328L530 13q7-7 17-7t16 7l33 33q7 7 7 16t-7 16L483 191zM46 354h649l-43 258q-2 8-8 14t-15 6H112q-8 0-14-6t-9-14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Behance(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M333 250q38 0 58 20t28 46q10 29 9 66q0 31-7 54t-18 39t-26 28t-30 19q-36 17-81 18H0V7h266q30 0 57 8t46 27t30 44t11 63t-12 56t-27 30q-17 11-38 15m-216-30h140q18 0 32-11t13-52q0-20-5-32t-13-17t-18-6H117zm141 225q8 0 19-3t20-9t17-21t6-36q0-28-8-43t-19-21q-12-8-27-9H117v142zm387-303q39 0 67 10t48 26t32 37t19 43q16 51 11 116H558q0 38 14 56t31 27q19 10 45 9q32 0 49-9t24-19q9-12 9-29h90q0 29-7 50t-20 36t-29 26t-32 16q-39 15-87 14q-39 0-68-9t-51-23t-35-33t-23-37q-20-45-20-102q3-57 25-102q9-19 24-38t35-33t48-23t65-9m80 161q0-23-9-41q-8-16-24-29t-45.5-13t-46.5 13t-28 30t-15 40zm21-199H535V38h211z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bell(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M579 372q0 24 9 45t25 37t37 25t45 9v93H0v-93q24 0 45-9t37-25t25-37t9-45V233q0-48 18-90t50-74t73-50t90-18t90 18t74 50t50 74t18 90zM347 696q-32 0-56-20t-33-49h179q-8 30-32 49t-58 20"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Book(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M116 564q-20 0-33 13t-14 34q0 19 14 32t33 14h417q9 0 16 6t7 17v12q0 24-17 41t-41 17H69q-29 0-49-21T0 680V78q0-29 20-50T69 8h452q14 0 25 10t10 25v498q0 10-7 16t-16 7zm46-313q0 12 12 12h208q12 0 12-12v-69q0-5-3-9t-9-3H174q-5 0-8 3t-4 9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Box(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M729 241q12 16 11 33l-20 261q-2 22-24 22H44q-9 0-15-6t-8-16L1 274q-1-17 10-33L198 14q10-13 27-13h291q16 0 27 13zm-85 73q0-5-3-9t-9-3H109q-5 0-8 3t-4 9l11 140q0 10 11 10h502q10 0 12-10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Brightness(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M46 333q20 0 34 13t13.5 32.5T80 411t-34 14q-19 0-32-14T0 379q0-20 14-33t32-13m63 242q14-14 33-14t32 14t14 33t-14 33t-32 14t-33-14t-14-33t14-33m65-393q-14 14-32 14t-33-14t-14-33t14-32t33-14t32 14t14 32t-14 33m197-81q-20 0-33-14t-14-32q0-20 14-34t33-13.5T404 21t13 34q0 19-13 32t-33 14m324 232q20 0 33 13t13 32.5t-13 32.5t-33 14t-33-14t-13-32q0-20 13-33t33-13M371 657q19 0 33 13t13.5 33t-13.5 33t-33 13.5t-33-13.5t-13.5-33t13.5-33t33-13m0-463q38 0 72 14t59 40t40 59t14 72t-14 72t-40 59t-59 39t-72 15q-39 0-72-15t-59-39t-40-59t-15-72t15-72t40-59t59-40t72-14m0 278q19 0 36-8t30-20t19-29t7-36t-7-36t-19-30t-30-20t-36-7t-36 7t-29 20t-20 30t-8 36t8 36t20 29t29 20t36 8m196 103q14-14 33-14t33 14t14 33t-14 33t-33 14t-33-14t-14-33t14-33m0-458q14-14 33-14t33 14t14 32t-14 33t-33 14t-33-14t-14-33t14-32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bucket(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m579 141l-53 509q-2 21-21 39t-51 31t-74 22t-90 8t-90-8t-74-22t-51-31t-22-39L0 141h1q-1-2-1-5q0-27 23-50t62-41t92-27T290 8t113 10t92 27t62 41t22 50zm-289 76q43 0 81-7t66-17t45-26t16-31t-16-32t-45-25l-66-18q-36-10-81-6q-44 0-82 6t-66 18t-44 25t-17 32t17 31t44 26t66 17t82 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M533 194q-20 0-33-14t-14-33V55q0-20 14-34t33-13.5T566 21t13 34v92q0 20-13 33t-33 14m-325 0q-19 0-32-14t-14-33V55q0-20 14-34t32-13q20 0 34 13t13 34v92q0 20-13 33t-34 14m487-93q20 0 33 13t13 33v556q0 20-13 33t-33 14H46q-19 0-32-14T0 703V147q0-19 14-33t32-13h70v46q0 19 7 36t20 29t29 20t36 8t36-8t30-20t20-29t7-36v-46h139v46q0 19 8 36t20 29t29 20t36 8t36-8t30-20t20-29t7-36v-46zm-46 289q0-11-12-11H104q-11 0-11 11v255q0 12 11 12h533q12 0 12-12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Camera(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M371 264q29 0 49 20t20 50t-20 49t-49 20t-49-20t-21-49t21-50t49-20m347-162q10 0 17 6t6 17v440q0 10-6 16t-17 7H23q-10 0-16-7t-7-16V125q0-10 7-17t16-6h139l16-56q3-16 16-27t29-10h295q17 0 30 10t15 27l16 56zM371 496q33 0 63-13t51-35t35-52t13-62q0-34-13-64t-35-51t-51-35t-63-13.5t-63 13.5t-52 35t-35 51t-13 64q0 33 13 62t35 52t52 35t63 13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Card(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M718 1q10 0 17 6t6 17v510q0 10-6 16t-17 7H23q-10 0-16-7t-7-16V24Q0 14 7 7t16-6zM93 325h208v-69H93zm393 70H93v69h393zM649 94h-70v69h70z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cart(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M278 557q29 0 49 20t20 50t-20 49t-49 20t-49-20t-21-49t21-50t49-20m324 0q29 0 50 20t20 50t-20 49t-50 20t-49-20t-20-49t20-50t49-20M255 418v24h452q11 0 11 11v46q0 12-11 12H174q-12 0-12-12q0-81-21-159L72 79q-3-8-11-8H5q-5 0-5-5V6q0-5 5-5h129q5 0 5 5v35q0 7 7 7h561q11 0 11 11v186q0 15-9 27t-24 17z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clock(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M465 424q10 10 10 24t-10 24.5t-25 10.5t-24-10l-70-70q-10-10-10-24V194q0-15 10-25t24.5-10t24.5 10t10 25v170zM371 8q76 0 144 29t118 79t79 118t29 145t-29 144t-79 118t-118 80t-144.5 29T227 721t-118-80t-80-118T0 379t29-145t80-118t118-79T371 8m0 649q57 0 108-22t89-60t59-88t21.5-108T627 270t-59-88t-89-59t-109-22t-107 22t-88 59t-60 88t-22 108.5T115 487t60 88t88 60t108 22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloud(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M556 123q36 0 68 13t55 38t38 55t13 68t-13 67t-38 56t-55 37t-68 14H139q-29 0-54-11t-44-30t-30-44t-11-54t11-55t30-44t44-29t54-11h12l11 2v-2q0-39 15-73t39-59t59-40t72-14q60 0 107 33t67 87q15-4 35-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cog(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M722 424q12 9 9 24q-11 46-33 86q-3 7-9 10t-15 1q-22-4-43 2t-39 24t-23 40t-2 46q3 16-10 24q-41 24-82 35q-16 5-26-10q-12-20-32-32t-46-12t-45 12t-34 32q-9 15-25 10q-44-12-82-35q-15-8-11-24q5-23-1-46t-23-40t-39-24t-44-2q-16 4-24-11q-22-42-32-86q-5-15 9-24q17-13 28-33t10.5-42t-9.5-40t-25-31q-12-11-8-24q15-48 40-89q7-12 20-10q20 2 40-5t34-22q29-29 27-69q-1-13 11-20q44-26 94-38q13-3 21 6q13 14 31 22t37 8t37-8t30-22q9-9 22-6q50 12 94 38q11 6 11 20q-2 40 27 69q15 15 34 22t39 5q13-2 21 10q25 41 40 89q3 14-8 24q-16 13-25 31t-9 40.5t11 41.5t27 33m-351 76q29 0 55-11t44-30t29-44t11-54t-11-55t-29-44t-44-29t-55-11t-54 11t-44 29t-30 44t-11 55t11 54t30 44t44 30t54 11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Comment(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M718 5q10 0 17 7t6 17v657q0 5-3 6t-7-1L579 538H23q-10 0-16-7t-7-16V29q0-10 7-17t16-7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compass(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M503 217q14-5 25 5t5 25l-69 184q-11 30-41 41l-185 69q-14 5-24-6t-6-24l70-184q11-30 40-41zM403 412q14-14 14-33t-14-33t-32-14t-33 14t-14 33t14 33t33 14t32-14M371 8q76 0 144 29t118 79t79 118t29 145t-29 144t-79 118t-118 80t-144.5 29T227 721t-118-80t-80-118T0 379t29-145t80-118t118-79T371 8m0 649q57 0 108-22t89-60t59-88t21.5-108T627 270t-59-88t-89-59t-109-22t-107 22t-88 59t-60 88t-22 108.5T115 487t60 88t88 60t108 22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Contrast(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M371 8q76 0 144 29t118 79t79 118t29 145t-29 144t-79 118t-118 80t-144.5 29T227 721t-118-80t-80-118T0 379t29-145t80-118t118-79T371 8m0 117q0-11-7-18t-19-5q-53 5-99 28t-80 61t-53 85t-20 103t20 102t53 86t80 60t99 29q11 1 19-6t7-17z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Controls(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M463 194H0v-93h463zm278-93v93h-92v81q0 5-3 8t-9 3h-93q-11 0-11-11V20q0-12 11-12h93q5 0 9 3t3 9v81zM0 564h255v93H0zm440 0h301v93H440v81q0 5-3 8t-8 4h-93q-5 0-8-4t-4-8V483q0-5 4-8t8-3h93q11 0 11 11zM232 333h509v92H232zM46 425H0v-92h46v-82q0-11 12-11h93q11 0 11 11v255q0 12-11 12H58q-12 0-12-12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Conversation(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M185 448h325v163q0 19-13 32t-34 14H185L69 750v-93H46q-19 0-32-14T0 611V309q0-19 14-33t32-13h139zM695 8q20 0 33 13t13 34v301q0 19-13 32t-33 14h-23v93l-116-93H232V55q0-20 13-34t33-13z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cup(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M371 5q76 0 144 12t118 32t79 47t29 59q-6 140-46 249q-17 47-44 91t-66 80t-92 56t-122 22q-70 0-123-22t-92-56t-66-80t-44-91Q6 295 0 155q0-31 29-59t80-47t118-32T371 5m247 162q8-4 8-12t-8-11q-33-17-95-32T371 97t-152 15t-96 32q-7 3-7 11t7 12q33 16 96 31t152 15t152-15t95-31"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dashboard(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M443 108q-29-11-60-16t-66 0q-49 6-90 29t-71 58t-46 80t-17 94v15q0 10-7 16t-17 7H23q-10 0-16-7t-7-16v-14q0-66 23-127t63-108t96-79T303 0q58-8 112 3t100 37q5 4 6 10t-5 10q-14 11-33 25t-29 22q-6 4-11 1m218 89q34 70 34 148v23q0 10-7 16t-17 7h-45q-10 0-17-7t-7-16v-23q0-37-11-75q-2-7 2-11l49-64q4-5 11-5t8 7m-87-64q8-6 16-5t14 6t7 14t-5 17L403 433l-3 3l-3 4q-21 20-50 20t-49-20t-20-49t20-49q2-2 4-3t4-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dialog(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M695 5q20 0 33 13t13 33v394q0 19-13 32t-33 14H463v162L278 491H46q-19 0-32-14T0 445V51q0-20 14-33T46 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dribbble(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M361 273q-21-37-42-71t-37-58t-28-41t-13-17q-68 32-115 91T62 312l29-1q26-1 66-4t93-11t111-23m40 108q5-1 13-3l-12-28q-6-14-14-28q-63 19-122 28t-105 14t-74 4t-31 1q-1 3-1 10q0 63 23 118t60 98q27-46 66-86q33-35 82-71t115-57M180 634l-3-3l-4-3l5 5q1 1 2 1m413-498q-43-38-97-60T379 55q-40 0-76 9l13 18q11 15 28 41l38 59q21 33 42 71q46-17 78-38t52-38t30-29zM379 758q-78 0-147-30t-120-82t-82-120T0 379t30-147t82-121t120-81T379 0t147 30t121 81t81 121t30 147t-30 147t-81 120t-121 82t-147 30m55-326q-71 25-119 59t-79 65t-43 54t-13 24q42 32 92 50t107 19q34 0 66-7t60-19l-6-34q-4-24-13-56l-20-72q-11-40-30-84h-1zm15-131q5 12 11 23t10 24l4 9q2 4 3 8q42-5 82-4t72 4t51 7t21 4q-1-58-20-110t-54-93l-10 13l-33 32q-33 32-55 41t-82 42m48 115q16 44 28 83t19 69t11 51l5 28q55-37 92-93t47-124l-20-6q-17-5-44-8t-63-5t-75 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Drop(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M515 206q45 45 69 103t26 118t-21 117t-67 103t-101 67t-116 23t-115-23t-102-67t-67-103T0 427t25-118t69-103L296 4q8-8 17 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dropbox(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M395 133L161 286L0 150L235 4zm143 493q5 0 10-3l96-63v37L395 754L147 597v-37l96 63q4 3 10 3t12-4l130-109l131 109q6 4 12 4m92-340L395 133L556 4l235 146zm0 0l143 115l-232 152l-146-121zM18 401l143-115l234 146l-145 121z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ellipsis(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M371 229q0 29-20 49t-50 20t-49-20t-20-49t20-50t49-20t50 20t20 50m-232 0q0 29-20 49t-50 20t-49-20t-20-49t20-50t49-20t50 20t20 50m463 0q0 29-20 49t-49 20t-49-20t-21-49t21-50t49-20t49 20t20 50"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Email(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M695 1q20 0 33 13t13 33v464q0 19-13 32t-33 14H46q-19 0-32-14T0 511V47q0-19 14-33T46 1zm-53 97q2-1 2-3t-3-1H101q-2 0-3 1t1 3l259 172q13 10 25 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Envelope(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M348 250q22 6 45 0l348-87v371q0 10-6 16t-17 7H23q-10 0-16-7t-7-16V163zM718 1q10 0 17 6t6 17v70l-348 87q-23 5-45 0L0 94V24Q0 14 7 7t16-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eye(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M370 183q19 0 33 13t13.5 32.5T403 261t-33 13.5t-33-13.5t-13-33t13-32t33-13m0-163q87 0 153 27t110 62t71 68t33 46q3 6 0 12q-5 8-18 26t-34 41t-52 47t-69 43t-87 32t-107 13t-106-13t-88-32t-69-43t-51-47t-35-40t-18-27q-5-6 0-12q7-13 33-46t71-68t111-62t152-27m0 325q24 0 45-9t37-25t25-37t9-45t-9-45t-25-37t-37-25t-45-9t-45 9t-37 25t-25 37t-9 45t9 45t25 37t37 25t45 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Facebook(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M659 8q34 0 59 24t24 58v577q0 35-24 59t-59 24H494V461h83l23-102H494v-69q0-22 13-41t44-19h49V123h-74q-79 0-117 42t-38 110v84h-83v102h83v289H82q-34 0-58-24T0 667V90q0-34 24-58T82 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func File(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M347 240q-9 0-16-7t-7-16V8l232 232zm-69-23q0 29 20 49t49 20h209v440q0 10-6 17t-17 7H23q-10 0-16-7t-7-17V31q0-9 7-16t16-7h255z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flag(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M628 340q10 11 5 25t-22 14H182q-11 0-18 8t-5 19l56 316q2 11-5 19t-18 9h-45q-20 0-23-20L1 35q-2-11 5-19t18-8h535q14 0 21 11t-1 24l-78 136q-9 15 3 27z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folder(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M714 190q11 0 18 7t5 18l-42 428q-2 10-12 10H55q-11 0-11-10L1 215q-1-10 6-18t17-7zm-20-46H45V28q0-10 7-17t16-6h185q10 0 17 6t7 17v23h393q10 0 17 6t7 17z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Github(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M510 383q23 0 39 21t15 53t-15 52t-39 22t-38-22t-16-52t16-53t38-21m186-193q29 31 46 69t16 90q0 74-18 125t-48 84t-64 52t-70 28t-64 10l-44 2H308q-15 0-44-2t-63-10t-70-28t-65-52t-47-84T0 349q0-51 17-90t45-69q-3-6-3-24t1-42t9-56T88 8q22 3 51 14q25 9 59 25t77 46q18-5 46-8t58-2l58 2q28 1 46 8q42-29 77-46t59-25q29-11 51-14q12 30 19 60t9 56t2 42t-4 24M380 614q58 0 109-6t88-23t59-51t21-90q0-27-10-52t-33-45q-19-18-44-24t-56-5t-64 4t-70 3h-2q-36 0-70-3t-64-4t-55 5t-44 24q-23 20-33 45t-11 52q0 57 22 90t58 51t88 23t109 6zM248 383q23 0 39 21t15 53t-15 52t-39 22t-38-22t-16-52t16-53t38-21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GooglePlus(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M39 187q0-43 14-74t35-50t46-32t47-17t35-7t15-1h189v4q0 8-6 13t-16 9t-21 4t-21 3l-17 2q-10 1-14 4q36 19 46 50t10 76q0 51-17 80t-39 50q-13 11-22 21t-9 23t12 25t29 27l28 26q15 14 28 33t21 40t8 49q0 113-96 160q-37 18-72 22t-51 4h-10q-8 0-25-1t-39-6t-43-15t-41-26t-31-43T.5 578.5T13 518t34-42t47-27t50-15t47-7t35-1h2q-12-15-17-30t-7-28t-1-25q-2 1-11 1q-15 0-41-4t-51-21q-61-40-61-132m306 371q-2-41-35-68t-86-27h-13q-27 2-51 11t-42 27t-27 36t-8 40.5T96 617t31 28t44 16t54 4q56-5 90-34t30-73m-50-411q-14-52-38-71t-62-18q-11 0-16 2q-34 9-48.5 50t-1.5 87q11 44 38 71t58 28q5 0 9-1t7-1q17-5 30-19t21-35t9-44t-6-49m445 139v83H616v123h-82V369H410v-83h124V163h82v123z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Grid(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M301 8q10 0 17 7t6 16v278q0 10-6 17t-17 7H23q-10 0-16-7t-7-17V31q0-9 7-16t16-7zM0 448q0-9 7-16t16-7h278q10 0 17 7t6 16v278q0 10-6 17t-17 7H23q-10 0-16-7t-7-17zM417 31q0-9 7-16t16-7h278q10 0 17 7t6 16v278q0 10-6 17t-17 7H440q-10 0-16-7t-7-17zm23 719q-10 0-16-7t-7-17V448q0-9 7-16t16-7h278q10 0 17 7t6 16v278q0 10-6 17t-17 7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heart(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M651 55q29 27 44 61t15 71t-13 71t-42 63L371 605q-7 7-16 7t-17-7L54 321q-27-28-40-63T1 187t14-71t42.5-61T121 17t70-11t70 16t61 42l33 33l33-33q27-27 61-42t70-16t70 11t62 38"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func House(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M733 289q8 6 8 18v406q0 10-6 16t-17 7H487q-10 0-17-7t-7-16V458q0-10-6-17t-17-6H301q-10 0-16 6t-7 17v255q0 10-6 16t-17 7H23q-10 0-16-7t-7-16V307q0-12 9-18L356 7q15-12 29 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Image(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M718 1q10 0 17 6t6 17v510q0 10-6 16t-17 7H23q-10 0-16-7t-7-16V24Q0 14 7 7t16-6zM602 418l-92-162l-93 162l-139-278l-139 278zm-46-185q20 0 33-14t13-33t-13-33t-33-13t-33 13t-13 32.5t13 33.5t33 14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Inbox(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M736 269q5 9 5 21v244q0 10-6 16t-17 7H23q-10 0-16-7t-7-16V290q0-9 5-21L126 26q13-25 42-25h406q28 0 41 25zm-95 33q3 0 5-3t-1-4l-97-195q-3-6-10-6H203q-7 0-10 6L96 295q-2 2 0 4t4 3h117q14 0 21 13l34 67q6 13 20 13h157q14 0 21-13l33-67q6-13 21-13z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Instagram(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M372 182q41 0 77 15t63 42t42 63t15 77t-15 76t-42 63t-63 42t-77 16t-77-16t-62-42t-42-63t-16-76t16-77t42-63t62-42t77-15m0 324q26 0 49-10t41-27t27-41t10-49t-10-50t-27-41t-41-27t-49-10t-49 10t-41 27t-27 41t-10 50t10 49t27 41t41 27t49 10m368-314q9 187 0 374q-2 36-17 68t-39 56t-57 40t-68 17q-47 2-93 3t-94 1t-93-1t-94-3q-36-2-68-17t-56-40t-40-56t-17-68q-8-187 0-374q2-36 17-68t40-57t56-39t68-17q187-9 374 0q36 2 68 17t57 39t39 57t17 68m-70 370q9-183 0-367q-1-22-11-42t-25-36t-36-26t-42-11q-46-2-92-3t-92-1t-92 1t-92 3q-22 1-42 11t-35 26t-26 36t-11 42q-9 184 0 368q1 22 11 42t26 35t35 26t42 11q184 9 368 0q22-1 42-11t36-26t25-36t11-42M569 138q18 0 31 13t13 31t-13 31t-31 13t-31-13t-13-31t13-31t31-13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Layers(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M712 353q6 4 6 10t-6 10L371 571q-12 7-23 0L6 373q-6-3-6-10t6-10l72-41l270 155q11 7 23 0l269-155zm0 163q6 3 6 9t-6 11L371 733q-12 7-23 0L6 536q-6-4-6-11t6-9l72-42l270 155q11 7 23 0l269-155zM6 222q-6-3-6-10t6-10L348 5q11-7 23 0l341 197q6 4 6 10t-6 10L371 420q-12 7-23 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Linkedin(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M165 90q0 35-21 59t-62 24q-37 0-59-24T0 95q0-35 23-61T83 8t60 24t22 58M0 750h165V214H0zm560-528q-32 0-57 8t-45 21t-33 27t-21 27h-4l-9-70H243q0 34 2 74t2 86v355h165V457q0-12 1-22t3-19q4-11 11-23t16-21t22-16t29-6q44 0 64 32t19 83v285h165V445q0-57-14-99t-38-70t-58-41t-72-13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Location(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M671 220q7 3 12 8t5 13t-3 13t-11 9L396 389q-8 5-11 12L259 678q-3 8-9 11t-14 3t-12-4t-9-12L2 36q-5-14 6-24t24-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lock(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M556 294v346q0 22-12 39t-32 26q-72 29-149 29H193q-77 0-149-29q-20-8-32-26T0 640V294h127V155q0-34 14-63t37-50t54-31t66-6q28 4 53 17t41 35t27 47t10 56v134zm-359 0h162V155q0-34-23-58t-58-24t-57 24t-24 58z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Market(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M648 256q0 2 1 3t0 3v20q0 10-7 17t-17 7h-46v324q0 10-6 16t-17 7H93q-10 0-17-7t-7-16V306H23q-10 0-16-7t-7-17v-20q0-4 1-6L60 39q5-16 17-25t28-9h439q16 0 28 9t16 25zm-138 50H139v127q0 5 4 8t8 4h347q5 0 9-4t3-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Menu(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M649 94H0V1h649zm0 231H0v-92h649zm0 232H0v-93h649z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mic(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M555 345q-2 50-20 94t-50 78t-72 58t-89 31v51h81q12 0 12 12v69q0 5-3 8t-9 4H150q-5 0-8-4t-3-8v-69q0-12 11-12h81v-51q-48-8-89-31t-72-58t-49-78t-21-94q0-12 12-12h69q10 0 12 11q2 36 17 68t40 55t58 37t70 14t70-14t57-37t41-55t17-68q0-11 11-11h70q5 0 9 3t2 9m-277 80q-24 0-45-9t-37-25t-25-37t-9-45V124q0-24 9-45t25-37t37-25t45-9t45 9t37 25t25 37t9 45v185q0 24-9 45t-25 37t-37 25t-45 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mobile(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 49q0-19 14-33T46 3h325q19 0 33 13t13 33v579q0 20-13 33t-33 14H46q-19 0-32-14T0 628zm208 579q20 0 34-13t13.5-32.5T242 549t-34-13q-19 0-32 13t-13.5 33t13.5 33t32 13m128-139q11 0 11-11V84q0-5-3-9t-8-3H81q-5 0-8 3t-4 9v394q0 5 4 8t8 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Money(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M713 39q11-3 20 4t8 18v369q0 8-4 14t-12 9q-58 17-116 18t-116-5t-116-18t-116-18t-116-7t-116 16q-11 3-20-5t-9-18V48q0-18 16-23Q74 8 132 6t116 6t117 17l116 18q57 9 116 7t116-15M371 355q17 0 32-9t26-25t18-37t6-45t-6-45t-18-37t-26-25t-32-9t-32 9t-27 25t-17 37l-7 45q-4 24 7 45t17 37t27 25t32 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Moon(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M732 392q3-2 7-1t3 5q-4 76-36 142t-84 114t-122 74t-147 23q-71-4-133-33t-109-77t-77-109T1 397q-4-78 23-147t74-122t114-84T354 8q4 0 6 3t-2 7q-31 40-46 90t-8 106q5 45 25 85t51 71t71 51t85 25q56 7 106-8t90-46"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Music(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M529 9q11-2 20 5t8 18v440q0 24-8 45t-24 38t-35 27t-45 11t-49-10t-41-29t-25-43t-3-52q5-35 30-61t59-34q25-4 48 0V209l-231 50v329q0 24-8 45t-24 38t-35 27t-45 11t-50-10t-40-29t-25-43t-4-52q6-35 31-61t59-34q25-4 48 0V111q0-8 5-14t13-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Notification(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M649 8q19 0 36 7t30 20t19 30t7 36t-7 36t-19 29t-30 20t-36 8q-20 0-36-8t-30-20t-20-29t-7-36t7-36t20-30t30-20t36-7m-93 255h93v463q0 10-7 17t-16 7H23q-10 0-16-7t-7-17V124q0-10 7-17t16-6h463v93H104q-11 0-11 11v440q0 12 11 12h440q12 0 12-12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paypal(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M232 407q-23 0-40 14t-22 38l-35 208H21q-9 0-15-7t-5-16l52-337L96 36q2-12 11-20t21-8h233q55 0 100 16t70 47q18 21 25 38q9 20 9 43v11q-1 6-1 12t-2 14q-1 4-1 7t-1 6q-20 104-84 154t-176 51zm375-189q21 25 26 60t-3 78q-10 52-32 87t-52 58t-69 31t-83 10h-18q-11 0-19 6t-10 18l-2 8l-22 145l-2 6q-2 11-9 18t-19 7H173l45-283q2-11 14-11h68q128 0 205-61t102-177"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pencil(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M626 70q25 25 41 52t23 52t4 46t-17 36L389 544L288 646L0 696l51-287l101-102L440 19q14-14 36-17t47 4t52 23t51 41M259 602q-12-20-30-44t-43-48t-48-43t-44-30L76 542q21 15 42.5 36.5T154 621z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pie(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M440 356q-10 0-16-7t-7-16V11q66 8 124 38t100 75t69 105t30 127zm-93 45q0 10 7 17t17 7h367q-9 72-44 133t-88 106t-121 67t-145 17q-68-6-128-35t-105-76t-74-106T1 403q-5-78 20-148t73-123t113-86T347 9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pin(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M531 99q45 46 68 103t23 117t-23 117t-68 103L327 743q-7 7-16 7t-16-7L91 539q-45-45-68-103T0 319t23-117T91.5 99T194 31T311 8t117 23t103 68M383 391q15-15 22-34t7-38t-7-38t-22-34t-34-22t-38-7t-38 7t-34 22t-22 34t-8 38t8 38t22 34t34 22t38 8t38-8t34-22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Refresh(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M403 3q67 0 127 24t106 68t74 102t31 126q3 72-22 136t-71 112t-108 76t-135 28q-58 0-110-19t-95-51q-4-3-4-9t3-9l50-49q7-7 15-2q30 22 66 34t75 12q54 0 101-22t81-59t50-86t10-104q-5-41-23-78t-46-65t-64-47t-78-24q-54-6-102 9t-86 48t-60 76t-25 97h69L116 466L0 327h70q2-67 29-126t72-103t105-70T403 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ribbon(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M371 8q19 0 33 13t13 34v676q0 12-11 17t-20-3L212 591q-4-3-7 0L31 745q-9 8-20 3T0 731V55q0-20 14-34T46 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Search(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M734 668q8 9 0 17l-49 49q-3 3-8 3t-8-3L519 584q-50 38-112 55t-133 6q-53-8-99-33t-83-61t-59-85T3 366q-10-79 16-150T96 95t121-76T367 3q53 7 100 30t84 59t62 82t33 100q11 69-6 131t-55 114zM325 557q48 0 90-18t74-50t50-74t18-90t-18-90t-50-74t-74-50t-90-18t-90 18t-73 50t-50 74t-18 90t18 90t50 74t73 50t90 18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Select(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 275h232L116 391zm232-92H0L116 67z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmallArrowDown(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M371 148q0 9-7 16L202 326q-7 7-16 7t-17-7L7 164q-7-7-7-17t7-16t16-7t16 7l146 146l146-146q7-7 17-7t16 7q7 9 7 17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmallArrowLeft(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M185 414q-10 0-16-7L7 245q-7-7-7-16t7-17L169 50q7-7 17-7t16 7t7 16t-7 17L56 229l146 146q7 7 7 16t-7 16t-17 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmallArrowRight(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23 44q10 0 16 6l163 162q7 8 7 17t-7 16L39 407q-7 7-16 7t-16-7t-7-16t7-16l146-146L7 83q-7-7-7-17t6.5-16T23 44"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmallArrowUp(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 310q0-10 7-16l162-163q7-7 17-7t16 7l162 163q7 7 7 16t-7 16t-16 7t-17-7L185 181L39 326q-7 7-16 7t-16-6.5T0 310"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tablet(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M510 8q19 0 33 13t13 34v648q0 20-13 33t-33 14H46q-19 0-32-14T0 703V55q0-20 14-34T46 8zM278 703q20 0 33-14t13-33t-13-32t-33-13t-33 13t-13 32.5t13 32.5t33 14M486 89q0-11-11-11H81q-5 0-8 3t-4 8v464q0 5 4 8t8 3h394q11 0 11-11z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tag(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M725 348q7 7 7 16t-7 16L398 708q-29 27-66 27H23q-10 0-16-7t-7-16V402q0-38 27-65L354 10q7-7 17-7t16 7zM185 619q29 0 50-21t20-49t-20-49t-50-20t-49 20t-20 49t20 49t49 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Thumbs(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M139 228q10 0 17 6t6 17v278q0 10-6 16t-17 7H23q-10 0-16-7t-7-16V251q0-10 7-17t16-6zm485 0q22 0 41 9t32 26t18 36t0 43l-48 202q-8 34-35 54t-63 17l-180-15q-18-1-36-6l-111-32q-15-5-24-17t-10-28V248q0-20 16-34L443 14q21-18 46-9l4 1q17 6 26 21t5 33l-31 139q-2 11 5 20t18 9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleDown(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 171h232L116 287z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleUp(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M232 287H0l116-116z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Twitter(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M741 170v22q0 56-14 113t-41 111t-69 101t-94 82t-119 55t-144 20q-72 0-137-20T0 598q19 3 40 3q60 0 114-20t97-53q-56-1-99-34t-60-84q16 3 32 3q22 0 45-5q-29-6-54-21t-43-37t-28-50t-11-59v-2q36 21 77 22q-35-23-55-60t-21-81q0-23 6-45t17-41q63 78 153 125t196 52q-2-9-3-19t-1-19q0-35 14-66t36-54t54-37t65-13q37 0 69 14t55 40q57-12 108-41q-10 29-29 53t-46 40q51-6 97-26q-34 52-84 87"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unlock(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M556 294v346q0 22-12 39t-32 26q-72 29-149 29H193q-77 0-149-29q-20-8-32-26T0 640V294h359V155q0-19-8-36t-21-27t-32-16t-37-1q-29 6-46 29t-18 53v32q0 12-12 12h-46q-5 0-8-4t-4-8v-34q0-34 14-63t37-50t54-31t66-6q28 4 53 17t41 35t27 47t10 56v134z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Url(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 537q3 1 4 4t-1 6l-83 83q-20 20-46 32t-53 14t-54-4t-49-25q-27-21-42-51T1 536t9-61t34-54l129-129q31-31 72-46t84-8q38 6 60 22t35 37q2 5-.5 7.5T418 309t-7 6l-7 7q-10 11-24 11t-25.5-11.5T329 304t-31-5t-30 5t-27 18L93 470q-13 13-19 30t-4 34t10 34t26 27q25 17 54 12t49-24l64-63q2-3 6-1q30 13 62 18M619 31q27 21 42 50t17 60t-9 62t-35 54L505 386q-31 31-72 45t-85 8q-38-5-59-22t-35-37q-2-4 1-7l18-18q11-10 25-10t24 10q12 12 27 18t31 6t31-6t26-18l148-147q13-13 19-31t4-35t-11-33t-26-27q-24-16-53-12t-50 25l-63 63q-3 3-6 1q-28-12-62-19q-4 0-4-4t1-5l83-83q20-20 46-32t52-15t54 5t50 25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func User(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M535 518q27 13 42 34t21 44t8 43t0 31t-8 16q-14 12-34 21t-44 16t-49 11l-49 8q-56 7-118 8q-63-1-119-8q-24-3-49-8t-48-11t-44-16t-35-21q-6-5-7.5-15.5T1 639t8-43t21-44t42-34l62-28q29-13 52-28t35-35t13-48v-6q-17-14-25-33l-23-54q-18-20-27-37q-8-14-8-26t17-12q-6-38-3-68q2-25 12-46t36-21q5-19 17-34q11-13 31-24t53-10q30 0 49 13t33 30t23 33t20 23q4 2 7 5t0 8q-4 14-8 41t1 50q18 1 17 12t-10 26q-10 17-28 37q-7 18-11 31t-9 24t-10 17t-15 15v6q0 29 13 48t36 35t51 28t62 28"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Users(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M535 518q27 13 42 34t21 44t8 43t0 31t-8 16q-14 12-34 21t-44 16t-49 11l-49 8q-56 7-118 8q-63-1-119-8q-24-3-49-8t-48-11t-44-16t-35-21q-6-5-7.5-15.5T1 639t8-43t21-44t42-34l62-28q29-13 52-28t35-35t13-48v-6q-17-14-25-33l-23-54q-18-20-27-37q-8-14-8-26t17-12q-6-38-3-68q2-25 12-46t36-21q5-19 17-34q11-13 31-24t53-10q30 0 49 13t33 30t23 33t20 23q4 2 7 5t0 8q-4 14-8 41t1 50q18 1 17 12t-10 26q-10 17-28 37q-7 18-11 31t-9 24t-10 17t-15 15v6q0 29 13 48t36 35t51 28t62 28m194-54q13 7 13 21v163q0 8-5 14t-14 8q-20 4-38 5t-34 3q0-1 1-2t0-2q1-13 0-39t-11-56t-29-58t-57-45l-46-21q-4-2-10-4q21-13 34-30t12-42v-5q-14-11-20-26l-18-44q-14-17-22-30q-7-11-7-21t14-11q-4-31-2-55q2-20 10-37t28-17q3-15 14-27q9-11 26-20t43-8q24 0 40 10t26 24l19 27q9 13 16 18t6 12q-4 11-7 32t1 41q14 1 14 11t-8 22t-23 29l-17 44q-6 15-20 26v5q0 33 20 51t51 34"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Videocamera(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M726 146q5-2 10 1t5 10v244q0 6-5 9t-10 2l-100-40V186zM510 1q19 0 33 13t13 33v464q0 19-13 32t-33 14H46q-19 0-32-14T0 511V47q0-19 14-33T46 1zM278 441q34 0 64-13t51-35t35-52t12-62q0-34-12-64t-35-51t-51-35t-64-12t-63 12t-52 35t-34 51t-13 64q0 33 13 62t34 52t52 35t63 13m0-255q19 0 36 7t30 20t20 30t7 36t-7 36t-20 29t-30 20t-36 8t-36-8t-29-20t-20-29t-8-36t8-36t20-30t29-20t36-7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func World(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M371 8q76 0 144 29t118 79t79 118t29 145t-29 144t-79 118t-118 80t-144.5 29T227 721t-118-80t-80-118T0 379t29-145t80-118t118-79T371 8M93 379q0 49 16 93t45 80t68 61t85 36q-13-19-9-35q5-20 25-40t44-29.5t41-23.5t17-23t-23-24t-58-22q-17-5-32-14t-28-20t-21-23t-12-24q-5-24-18-31t-27.5 7.5T178 372t-19 10t2-20t21-49t21-53t2-28t-21-9t-37-3q-3 0-4-1q-24 35-37 75t-13 85m512 150q20-32 32-70t12-80q0-36-9-70t-25-63q-1 3-5 9q-9 14-26 34t-31.5 34.5T531 338t-2-14t-8-21t-38-2t-47 15t-34 20t-10 23t10 23t27 10t31-10t28-10t18 10t2 33t-13 48q-10 24-3 38t25 9t43 0t45 19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Youtube(children ...ElementRenderer) *IlIcon {
	return &IlIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M422 2q101 0 171 4t116 16t71 33t37 57t15 86t2 121l-2 121q-1 50-15 86t-37 56t-71 34t-116 16t-171 4t-171-4t-116-16t-71-34t-37-56t-14-86t-3-121t3-121t14-86t37-57t71-33T251 6t171-4m132 331q12-6 12-14t-12-14l-185-86q-12-6-20-1t-9 20v162q0 14 9 19t20 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
