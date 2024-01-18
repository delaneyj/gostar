package foundation

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = "3.0.0"
	hAttr          = "1em"
	viewbox        = "0 0 100 100"
	fill           = "currentColor"
)

type FoundationIcon struct {
	*SVGSVGElement
}

type FoundationIconFn func(children ...ElementRenderer) *FoundationIcon

var IconLookup = map[string]FoundationIconFn{
	"addressBook":         AddressBook,
	"alert":               Alert,
	"alignCenter":         AlignCenter,
	"alignJustify":        AlignJustify,
	"alignLeft":           AlignLeft,
	"alignRight":          AlignRight,
	"anchor":              Anchor,
	"annotate":            Annotate,
	"archive":             Archive,
	"arrowDown":           ArrowDown,
	"arrowLeft":           ArrowLeft,
	"arrowRight":          ArrowRight,
	"arrowUp":             ArrowUp,
	"arrowsCompress":      ArrowsCompress,
	"arrowsExpand":        ArrowsExpand,
	"arrowsIn":            ArrowsIn,
	"arrowsOut":           ArrowsOut,
	"asl":                 Asl,
	"asterisk":            Asterisk,
	"atSign":              AtSign,
	"backgroundColor":     BackgroundColor,
	"batteryEmpty":        BatteryEmpty,
	"batteryFull":         BatteryFull,
	"batteryHalf":         BatteryHalf,
	"bitcoin":             Bitcoin,
	"bitcoinCircle":       BitcoinCircle,
	"blind":               Blind,
	"bluetooth":           Bluetooth,
	"bold":                Bold,
	"book":                Book,
	"bookBookmark":        BookBookmark,
	"bookmark":            Bookmark,
	"braille":             Braille,
	"burst":               Burst,
	"burstNew":            BurstNew,
	"burstSale":           BurstSale,
	"calendar":            Calendar,
	"camera":              Camera,
	"check":               Check,
	"checkbox":            Checkbox,
	"clipboard":           Clipboard,
	"clipboardNotes":      ClipboardNotes,
	"clipboardPencil":     ClipboardPencil,
	"clock":               Clock,
	"closedCaption":       ClosedCaption,
	"cloud":               Cloud,
	"comment":             Comment,
	"commentMinus":        CommentMinus,
	"commentQuotes":       CommentQuotes,
	"commentVideo":        CommentVideo,
	"comments":            Comments,
	"compass":             Compass,
	"contrast":            Contrast,
	"creditCard":          CreditCard,
	"crop":                Crop,
	"crown":               Crown,
	"cssThree":            CssThree,
	"database":            Database,
	"dieFive":             DieFive,
	"dieFour":             DieFour,
	"dieOne":              DieOne,
	"dieSix":              DieSix,
	"dieThree":            DieThree,
	"dieTwo":              DieTwo,
	"dislike":             Dislike,
	"dollar":              Dollar,
	"dollarBill":          DollarBill,
	"download":            Download,
	"eject":               Eject,
	"elevator":            Elevator,
	"euro":                Euro,
	"eye":                 Eye,
	"fastForward":         FastForward,
	"female":              Female,
	"femaleSymbol":        FemaleSymbol,
	"filter":              Filter,
	"firstAid":            FirstAid,
	"flag":                Flag,
	"folder":              Folder,
	"folderAdd":           FolderAdd,
	"folderLock":          FolderLock,
	"foot":                Foot,
	"foundation":          Foundation,
	"graphBar":            GraphBar,
	"graphHorizontal":     GraphHorizontal,
	"graphPie":            GraphPie,
	"graphTrend":          GraphTrend,
	"guideDog":            GuideDog,
	"hearingAid":          HearingAid,
	"heart":               Heart,
	"home":                Home,
	"htmlFive":            HtmlFive,
	"indentLess":          IndentLess,
	"indentMore":          IndentMore,
	"info":                Info,
	"italic":              Italic,
	"key":                 Key,
	"laptop":              Laptop,
	"layout":              Layout,
	"lightbulb":           Lightbulb,
	"like":                Like,
	"link":                Link,
	"list":                List,
	"listBullet":          ListBullet,
	"listNumber":          ListNumber,
	"listThumbnails":      ListThumbnails,
	"lock":                Lock,
	"loop":                Loop,
	"magnifyingGlass":     MagnifyingGlass,
	"mail":                Mail,
	"male":                Male,
	"maleFemale":          MaleFemale,
	"maleSymbol":          MaleSymbol,
	"map":                 Map,
	"marker":              Marker,
	"megaphone":           Megaphone,
	"microphone":          Microphone,
	"minus":               Minus,
	"minusCircle":         MinusCircle,
	"mobile":              Mobile,
	"mobileSignal":        MobileSignal,
	"monitor":             Monitor,
	"mountains":           Mountains,
	"music":               Music,
	"next":                Next,
	"noDogs":              NoDogs,
	"noSmoking":           NoSmoking,
	"page":                Page,
	"pageAdd":             PageAdd,
	"pageCopy":            PageCopy,
	"pageCsv":             PageCsv,
	"pageDelete":          PageDelete,
	"pageDoc":             PageDoc,
	"pageEdit":            PageEdit,
	"pageExport":          PageExport,
	"pageExportCsv":       PageExportCsv,
	"pageExportDoc":       PageExportDoc,
	"pageExportPdf":       PageExportPdf,
	"pageFilled":          PageFilled,
	"pageMultiple":        PageMultiple,
	"pagePdf":             PagePdf,
	"pageRemove":          PageRemove,
	"pageSearch":          PageSearch,
	"paintBucket":         PaintBucket,
	"paperclip":           Paperclip,
	"pause":               Pause,
	"paw":                 Paw,
	"paypal":              Paypal,
	"pencil":              Pencil,
	"photo":               Photo,
	"play":                Play,
	"playCircle":          PlayCircle,
	"playVideo":           PlayVideo,
	"plus":                Plus,
	"pound":               Pound,
	"power":               Power,
	"previous":            Previous,
	"priceTag":            PriceTag,
	"pricetagMultiple":    PricetagMultiple,
	"print":               Print,
	"prohibited":          Prohibited,
	"projectionScreen":    ProjectionScreen,
	"puzzle":              Puzzle,
	"quote":               Quote,
	"record":              Record,
	"refresh":             Refresh,
	"results":             Results,
	"resultsDemographics": ResultsDemographics,
	"rewind":              Rewind,
	"rewindTen":           RewindTen,
	"rss":                 Rss,
	"safetyCone":          SafetyCone,
	"save":                Save,
	"share":               Share,
	"sheriffBadge":        SheriffBadge,
	"shield":              Shield,
	"shoppingBag":         ShoppingBag,
	"shoppingCart":        ShoppingCart,
	"shuffle":             Shuffle,
	"skull":               Skull,
	"socialAdobe":         SocialAdobe,
	"socialAmazon":        SocialAmazon,
	"socialAndroid":       SocialAndroid,
	"socialApple":         SocialApple,
	"socialBehance":       SocialBehance,
	"socialBing":          SocialBing,
	"socialBlogger":       SocialBlogger,
	"socialDelicious":     SocialDelicious,
	"socialDesignerNews":  SocialDesignerNews,
	"socialDeviantArt":    SocialDeviantArt,
	"socialDigg":          SocialDigg,
	"socialDribbble":      SocialDribbble,
	"socialDrive":         SocialDrive,
	"socialDropbox":       SocialDropbox,
	"socialEvernote":      SocialEvernote,
	"socialFacebook":      SocialFacebook,
	"socialFiveHundredPx": SocialFiveHundredPx,
	"socialFlickr":        SocialFlickr,
	"socialForrst":        SocialForrst,
	"socialFoursquare":    SocialFoursquare,
	"socialGameCenter":    SocialGameCenter,
	"socialGithub":        SocialGithub,
	"socialGooglePlus":    SocialGooglePlus,
	"socialHackerNews":    SocialHackerNews,
	"socialHiFive":        SocialHiFive,
	"socialInstagram":     SocialInstagram,
	"socialJoomla":        SocialJoomla,
	"socialLastfm":        SocialLastfm,
	"socialLinkedin":      SocialLinkedin,
	"socialMedium":        SocialMedium,
	"socialMyspace":       SocialMyspace,
	"socialOrkut":         SocialOrkut,
	"socialPath":          SocialPath,
	"socialPicasa":        SocialPicasa,
	"socialPinterest":     SocialPinterest,
	"socialRdio":          SocialRdio,
	"socialReddit":        SocialReddit,
	"socialSkillshare":    SocialSkillshare,
	"socialSkype":         SocialSkype,
	"socialSmashingMag":   SocialSmashingMag,
	"socialSnapchat":      SocialSnapchat,
	"socialSpotify":       SocialSpotify,
	"socialSquidoo":       SocialSquidoo,
	"socialStackOverflow": SocialStackOverflow,
	"socialSteam":         SocialSteam,
	"socialStumbleupon":   SocialStumbleupon,
	"socialTreehouse":     SocialTreehouse,
	"socialTumblr":        SocialTumblr,
	"socialTwitter":       SocialTwitter,
	"socialVimeo":         SocialVimeo,
	"socialWindows":       SocialWindows,
	"socialXbox":          SocialXbox,
	"socialYahoo":         SocialYahoo,
	"socialYelp":          SocialYelp,
	"socialYoutube":       SocialYoutube,
	"socialZerply":        SocialZerply,
	"socialZurb":          SocialZurb,
	"sound":               Sound,
	"star":                Star,
	"stop":                Stop,
	"strikethrough":       Strikethrough,
	"subscript":           Subscript,
	"superscript":         Superscript,
	"tabletLandscape":     TabletLandscape,
	"tabletPortrait":      TabletPortrait,
	"target":              Target,
	"targetTwo":           TargetTwo,
	"telephone":           Telephone,
	"telephoneAccessible": TelephoneAccessible,
	"textColor":           TextColor,
	"thumbnails":          Thumbnails,
	"ticket":              Ticket,
	"torso":               Torso,
	"torsoBusiness":       TorsoBusiness,
	"torsoFemale":         TorsoFemale,
	"torsos":              Torsos,
	"torsosAll":           TorsosAll,
	"torsosAllFemale":     TorsosAllFemale,
	"torsosFemaleMale":    TorsosFemaleMale,
	"torsosMaleFemale":    TorsosMaleFemale,
	"trash":               Trash,
	"trees":               Trees,
	"trophy":              Trophy,
	"underline":           Underline,
	"universalAccess":     UniversalAccess,
	"unlink":              Unlink,
	"unlock":              Unlock,
	"upload":              Upload,
	"uploadCloud":         UploadCloud,
	"usb":                 Usb,
	"video":               Video,
	"volume":              Volume,
	"volumeNone":          VolumeNone,
	"volumeStrike":        VolumeStrike,
	"web":                 Web,
	"wheelchair":          Wheelchair,
	"widget":              Widget,
	"wrench":              Wrench,
	"x":                   X,
	"xcircle":             Xcircle,
	"yen":                 Yen,
	"zoomIn":              ZoomIn,
	"zoomOut":             ZoomOut,
}

func AddressBook(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M81.195 31.517c1.128 0 2.042-.897 2.042-1.996V23.61c0-1.102-.914-2-2.042-2h-4.586v-1.603a3.207 3.207 0 0 0-3.206-3.207H19.97a3.206 3.206 0 0 0-3.206 3.207v59.986a3.21 3.21 0 0 0 3.206 3.206h53.432a3.21 3.21 0 0 0 3.206-3.206v-1.67h4.586c1.128 0 2.042-.894 2.042-1.996v-5.912c0-1.102-.914-2-2.042-2h-4.586v-5.699h4.586c1.128 0 2.042-.894 2.042-1.991v-5.912c0-1.102-.914-2-2.042-2h-4.586v-5.699h4.586c1.128 0 2.042-.897 2.042-1.996v-5.912c0-1.098-.914-1.996-2.042-1.996h-4.586v-5.695h4.587zM62.391 63.681c0 1.152-.804 2.088-1.795 2.088H32.75c-.992 0-1.795-.935-1.795-2.088v-8.604c0-.856.447-1.625 1.127-1.941l10.9-5.077c-2.583-1.557-4.351-4.689-4.351-8.304c0-5.168 3.599-9.356 8.041-9.356c4.443 0 8.042 4.188 8.042 9.356c0 3.562-1.708 6.655-4.226 8.238l10.789 5.148c.674.325 1.115 1.085 1.115 1.937v8.603z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Alert(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m91.17 81.374l.006-.004l-.139-.24c-.068-.128-.134-.257-.216-.375l-37.69-65.283c-.611-1.109-1.776-1.87-3.133-1.87c-1.47 0-2.731.887-3.285 2.153l-.004-.002L9.312 80.529l.036.021a3.553 3.553 0 0 0-.82 2.257a3.59 3.59 0 0 0 3.588 3.59h75.767a3.59 3.59 0 0 0 3.589-3.589c0-.511-.11-.994-.302-1.434m-41.135-1.757c-2.874 0-5.201-2.257-5.201-5.13c0-2.874 2.326-5.2 5.201-5.2c2.803 0 5.13 2.325 5.13 5.2a5.123 5.123 0 0 1-5.13 5.13m5.13-45.367v28.299h-.002l.002.016c0 1.173-.95 2.094-2.094 2.094l-.014-.001v.001h-6.093c-1.174 0-2.123-.921-2.123-2.094l.002-.016h-.002V34.326c-.001-.026-.008-.051-.008-.077c0-1.117.865-1.996 1.935-2.078v-.016h6.288v.001c1.149.007 2.074.897 2.103 2.039h.005z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignCenter(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M81.232 15.138H18.769a3.407 3.407 0 0 0-3.407 3.407v3.143a3.407 3.407 0 0 0 3.407 3.407h62.463a3.407 3.407 0 0 0 3.407-3.407v-3.143a3.407 3.407 0 0 0-3.407-3.407m-5.008 26.473v-3.143a3.407 3.407 0 0 0-3.407-3.407H27.183a3.407 3.407 0 0 0-3.407 3.407v3.143a3.407 3.407 0 0 0 3.407 3.407h45.634a3.408 3.408 0 0 0 3.407-3.407m1.586 33.294H22.176a3.407 3.407 0 0 0-3.407 3.407v3.143a3.407 3.407 0 0 0 3.407 3.407H77.81a3.407 3.407 0 0 0 3.407-3.407h.015v-3.143h-.015a3.408 3.408 0 0 0-3.407-3.407M72.817 58.39a3.407 3.407 0 0 0-3.407-3.407H30.59a3.407 3.407 0 0 0-3.407 3.407v3.143a3.407 3.407 0 0 0 3.407 3.407h38.82a3.407 3.407 0 0 0 3.407-3.407z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignJustify(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M81.232 15.389H18.769a3.407 3.407 0 0 0-3.407 3.407v3.143a3.407 3.407 0 0 0 3.407 3.407h62.463a3.407 3.407 0 0 0 3.407-3.407v-3.143a3.407 3.407 0 0 0-3.407-3.407m0 19.755H18.769a3.407 3.407 0 0 0-3.407 3.407v3.143a3.407 3.407 0 0 0 3.407 3.407h62.463a3.407 3.407 0 0 0 3.407-3.407v-3.143a3.408 3.408 0 0 0-3.407-3.407m0 19.755H18.769a3.407 3.407 0 0 0-3.407 3.407v3.143a3.407 3.407 0 0 0 3.407 3.407h62.463a3.407 3.407 0 0 0 3.407-3.407v-3.143a3.408 3.408 0 0 0-3.407-3.407m0 19.755H18.769a3.407 3.407 0 0 0-3.407 3.407v3.143a3.407 3.407 0 0 0 3.407 3.407h62.463a3.407 3.407 0 0 0 3.407-3.407v-3.143a3.408 3.408 0 0 0-3.407-3.407"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignLeft(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M81.232 15.138H18.769a3.407 3.407 0 0 0-3.407 3.407v3.143a3.407 3.407 0 0 0 3.407 3.407h62.463a3.407 3.407 0 0 0 3.407-3.407v-3.143a3.407 3.407 0 0 0-3.407-3.407m-62.463 29.87h45.634a3.407 3.407 0 0 0 3.407-3.407v-3.143a3.407 3.407 0 0 0-3.407-3.407H18.769a3.407 3.407 0 0 0-3.407 3.407v3.143a3.406 3.406 0 0 0 3.407 3.407m55.634 29.897H18.769a3.407 3.407 0 0 0-3.407 3.407v3.143a3.407 3.407 0 0 0 3.407 3.407h55.634a3.407 3.407 0 0 0 3.407-3.407h.015v-3.143h-.015a3.408 3.408 0 0 0-3.407-3.407M18.769 64.94h38.82a3.407 3.407 0 0 0 3.407-3.407V58.39a3.407 3.407 0 0 0-3.407-3.407h-38.82a3.407 3.407 0 0 0-3.407 3.407v3.143a3.407 3.407 0 0 0 3.407 3.407"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignRight(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M81.232 15.138H18.769a3.407 3.407 0 0 0-3.407 3.407v3.143a3.407 3.407 0 0 0 3.407 3.407h62.463a3.407 3.407 0 0 0 3.407-3.407v-3.143a3.407 3.407 0 0 0-3.407-3.407m0 19.922H35.597a3.407 3.407 0 0 0-3.407 3.407v3.143a3.407 3.407 0 0 0 3.407 3.407h45.634a3.407 3.407 0 0 0 3.407-3.407v-3.143a3.407 3.407 0 0 0-3.406-3.407m-.015 39.845H25.583a3.407 3.407 0 0 0-3.407 3.407v3.143a3.407 3.407 0 0 0 3.407 3.407h55.634a3.407 3.407 0 0 0 3.407-3.407h.015v-3.143h-.015a3.408 3.408 0 0 0-3.407-3.407m3.422-16.515a3.407 3.407 0 0 0-3.407-3.407h-38.82a3.407 3.407 0 0 0-3.407 3.407v3.143a3.407 3.407 0 0 0 3.407 3.407h38.82a3.407 3.407 0 0 0 3.407-3.407z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Anchor(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m88.189 67.115l.007-.004l-9.393-16.27l-.009.005a1.571 1.571 0 0 0-1.376-.822c-.677 0-1.25.429-1.474 1.027l-9.158 15.862a1.564 1.564 0 0 0-.356.987c0 .873.706 1.578 1.577 1.578h2.39a27.536 27.536 0 0 1-15.42 8.6v-29.2h9.079c.871 0 1.578-.705 1.578-1.578l-.001-.008v-6.793c0-.871-.706-1.577-1.577-1.577h-9.079v-5.36c3.905-1.864 6.611-5.838 6.611-10.454c0-6.401-5.189-11.589-11.589-11.589S38.41 16.707 38.41 23.108c0 4.617 2.705 8.59 6.611 10.454v5.36h-9.08c-.871 0-1.577.706-1.577 1.576h-.001v6.803h.001a1.576 1.576 0 0 0 1.576 1.577h9.08v29.2a27.53 27.53 0 0 1-15.42-8.6h2.389c.872 0 1.578-.705 1.578-1.578c0-.287-.083-.553-.217-.785l.007-.004l-9.393-16.27l-.009.005a1.572 1.572 0 0 0-1.377-.822c-.675 0-1.246.427-1.471 1.022l-9.162 15.869c-.218.271-.354.61-.354.985c0 .873.706 1.578 1.577 1.578h4.219C23.835 80.823 36.017 88.482 50 88.482s26.165-7.658 32.611-19.003h4.218c.872 0 1.577-.705 1.577-1.578a1.56 1.56 0 0 0-.217-.786M50 29.634a6.526 6.526 0 0 1 0-13.051a6.528 6.528 0 0 1 6.526 6.527c0 3.6-2.923 6.524-6.526 6.524"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Annotate(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M82.488 68.439V31.561a10.413 10.413 0 0 0 5.424-9.143c0-5.753-4.664-10.417-10.417-10.417a10.402 10.402 0 0 0-8.577 4.512H31.083a10.402 10.402 0 0 0-8.577-4.512c-5.753 0-10.417 4.664-10.417 10.417c0 3.944 2.192 7.375 5.424 9.143v36.877a10.413 10.413 0 0 0-5.424 9.143c0 5.753 4.664 10.417 10.417 10.417c4.371 0 8.107-2.695 9.653-6.512h35.682c1.546 3.816 5.282 6.512 9.653 6.512c5.753 0 10.417-4.664 10.417-10.417c0-3.943-2.192-7.373-5.423-9.142m-55-.007V31.568a10.456 10.456 0 0 0 4.61-5.081h35.806a10.46 10.46 0 0 0 4.61 5.081v36.864a10.455 10.455 0 0 0-3.473 3.081h-38.08a10.453 10.453 0 0 0-3.473-3.081"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Archive(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M89.148 32.927c.001-.037.011-.07.011-.107a3.972 3.972 0 0 0-1.016-2.642l.02-.011l-7.87-13.627a2.53 2.53 0 0 0-2.468-1.914c-.083 0-.161.016-.242.024v-.024H22.219v.004c-.013 0-.026-.004-.039-.004a2.42 2.42 0 0 0-2.17 1.315l-.008-.005l-8.212 14.231l.015.008a4.068 4.068 0 0 0-.963 2.642c0 .047.012.091.014.138v48.211c-.002.048-.014.093-.014.142c0 2.284 1.817 4.069 4.095 4.066c.043 0 .083-.011.125-.012h69.87c.043.001.083.012.126.012c2.283 0 4.1-1.782 4.1-4.062c0-.036-.01-.068-.011-.104zM63.413 57.492l-12.391 17.43c-.226.318-.59.505-.98.507h-.004c-.386 0-.751-.187-.977-.503L36.59 57.494a1.201 1.201 0 0 1-.091-1.251c.208-.401.62-.654 1.071-.654h5.833l.001-15.654c0-.667.538-1.205 1.203-1.205h10.789c.665 0 1.204.539 1.204 1.204v15.655h5.83a1.206 1.206 0 0 1 .983 1.903M18.376 28.733l5.263-9.119h52.67l5.266 9.119z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDown(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.984 50.868l26.055 36.418a2.516 2.516 0 0 0 2.043 1.051h.006a2.52 2.52 0 0 0 2.048-1.059l25.887-36.417a2.513 2.513 0 0 0 .183-2.612a2.509 2.509 0 0 0-2.236-1.361H63.787l.001-32.709a2.514 2.514 0 0 0-2.516-2.515l-22.541.001a2.515 2.515 0 0 0-2.516 2.516v32.705H24.029c-.94 0-1.803.53-2.237 1.367a2.51 2.51 0 0 0 .192 2.615"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeft(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M49.132 21.984L12.714 48.039a2.516 2.516 0 0 0-1.051 2.043v.006a2.52 2.52 0 0 0 1.059 2.048L49.14 78.023a2.513 2.513 0 0 0 2.612.183a2.508 2.508 0 0 0 1.361-2.236V63.787l32.709.001a2.514 2.514 0 0 0 2.515-2.516l-.001-22.541a2.515 2.515 0 0 0-2.516-2.516H53.114V24.029c0-.94-.53-1.803-1.367-2.237a2.51 2.51 0 0 0-2.615.192"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRight(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m50.868 78.016l36.418-26.055a2.516 2.516 0 0 0 1.051-2.043v-.006a2.52 2.52 0 0 0-1.059-2.048L50.86 21.977a2.513 2.513 0 0 0-2.612-.183a2.509 2.509 0 0 0-1.361 2.236v12.183l-32.709-.001a2.514 2.514 0 0 0-2.515 2.516l.001 22.541a2.515 2.515 0 0 0 2.516 2.516h32.706v12.187c0 .94.53 1.803 1.366 2.237a2.512 2.512 0 0 0 2.616-.193"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUp(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M78.016 49.132L51.961 12.714a2.516 2.516 0 0 0-2.043-1.051h-.006a2.52 2.52 0 0 0-2.048 1.059L21.977 49.14a2.513 2.513 0 0 0-.183 2.612a2.509 2.509 0 0 0 2.236 1.361h12.183l-.001 32.709a2.514 2.514 0 0 0 2.516 2.515l22.541-.001a2.515 2.515 0 0 0 2.516-2.517V53.114h12.187c.94 0 1.803-.53 2.237-1.367a2.513 2.513 0 0 0-.193-2.615"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsCompress(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m39.171 88.427l5.103-30.768a1.746 1.746 0 0 0-.489-1.524l-.002-.003a1.754 1.754 0 0 0-1.531-.486l-30.685 5.186a1.748 1.748 0 0 0-1.375 1.196a1.748 1.748 0 0 0 .43 1.772l6 5.999L.514 85.907a1.752 1.752 0 0 0 0 2.479l11.1 11.101a1.752 1.752 0 0 0 2.479 0l16.108-16.108l6.002 6.002a1.76 1.76 0 0 0 1.774.429a1.75 1.75 0 0 0 1.194-1.383m21.658-76.854l-5.104 30.768a1.75 1.75 0 0 0 .489 1.524l.003.003c.403.4.972.581 1.53.486l30.685-5.187a1.75 1.75 0 0 0 1.376-1.196a1.748 1.748 0 0 0-.431-1.773l-5.999-5.999l16.108-16.107a1.752 1.752 0 0 0 0-2.479L88.386.514a1.752 1.752 0 0 0-2.479 0L69.799 16.621l-6.002-6.001a1.753 1.753 0 0 0-2.968.953"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsExpand(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.128 67.194L.024 97.962c-.093.558.09 1.125.489 1.524l.002.003c.403.4.972.581 1.531.486l30.685-5.186a1.748 1.748 0 0 0 1.376-1.197a1.746 1.746 0 0 0-.431-1.771l-6-5.999l16.109-16.108a1.75 1.75 0 0 0 0-2.478L32.684 56.134a1.752 1.752 0 0 0-2.479 0L14.098 72.242L8.096 66.24a1.752 1.752 0 0 0-2.968.954m89.744-34.388l5.104-30.768a1.748 1.748 0 0 0-.489-1.524l-.002-.003a1.762 1.762 0 0 0-1.531-.487L67.269 5.21a1.755 1.755 0 0 0-.946 2.969l6.001 5.998l-16.109 16.109a1.753 1.753 0 0 0 0 2.48l11.102 11.101a1.752 1.752 0 0 0 2.479 0l16.107-16.108l6.001 6.001a1.76 1.76 0 0 0 1.775.43a1.757 1.757 0 0 0 1.193-1.384"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsIn(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m35.995 89.365l4.69-28.272a1.609 1.609 0 0 0-.449-1.401l-.002-.002a1.608 1.608 0 0 0-1.406-.447L10.63 64.007c-.594.1-1.081.525-1.264 1.1a1.607 1.607 0 0 0 .395 1.628l5.513 5.513L.473 87.05a1.608 1.608 0 0 0 0 2.277l10.2 10.201a1.61 1.61 0 0 0 2.278 0l14.802-14.802l5.515 5.515a1.614 1.614 0 0 0 1.63.394a1.608 1.608 0 0 0 1.097-1.27m28.011-78.73l-4.69 28.273a1.61 1.61 0 0 0 .449 1.4l.002.003c.37.367.893.534 1.406.447l28.196-4.766a1.61 1.61 0 0 0 .869-2.728l-5.513-5.513l14.802-14.802a1.61 1.61 0 0 0 0-2.278l-10.2-10.2a1.61 1.61 0 0 0-2.278 0L72.248 15.274l-5.515-5.515a1.61 1.61 0 0 0-2.727.876m25.359 53.371l-28.271-4.69a1.608 1.608 0 0 0-1.401.449l-.003.003a1.606 1.606 0 0 0-.447 1.406l4.765 28.196c.1.594.525 1.082 1.099 1.264a1.607 1.607 0 0 0 1.629-.395l5.513-5.513l14.803 14.803a1.61 1.61 0 0 0 2.277-.002l10.201-10.2a1.61 1.61 0 0 0 0-2.278L84.727 72.248l5.516-5.515a1.612 1.612 0 0 0-.878-2.727m-78.73-28.011l28.272 4.69a1.605 1.605 0 0 0 1.401-.449l.002-.003a1.61 1.61 0 0 0 .448-1.406L35.993 10.63a1.61 1.61 0 0 0-1.099-1.264a1.604 1.604 0 0 0-1.629.395l-5.513 5.513L12.951.473a1.61 1.61 0 0 0-2.278-.001L.472 10.673a1.61 1.61 0 0 0 0 2.278l14.801 14.802l-5.515 5.515a1.612 1.612 0 0 0 .877 2.727"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsOut(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.712 69.854L.022 98.127a1.606 1.606 0 0 0 .449 1.401l.002.003c.37.368.893.535 1.406.447l28.197-4.765a1.608 1.608 0 0 0 1.264-1.1a1.605 1.605 0 0 0-.396-1.628l-5.513-5.513L40.233 72.17c.63-.629.63-1.647 0-2.277l-10.2-10.201a1.61 1.61 0 0 0-2.278 0L12.954 74.493l-5.515-5.515a1.61 1.61 0 0 0-2.727.876m90.576-39.708l4.69-28.273a1.609 1.609 0 0 0-.449-1.4L99.527.47a1.614 1.614 0 0 0-1.406-.447L69.924 4.788a1.61 1.61 0 0 0-.869 2.728l5.513 5.513L59.766 27.83a1.61 1.61 0 0 0 0 2.278l10.2 10.201a1.61 1.61 0 0 0 2.278 0l14.802-14.802l5.515 5.515a1.613 1.613 0 0 0 2.727-.876M69.854 95.288l28.271 4.69a1.605 1.605 0 0 0 1.401-.449l.002-.003a1.61 1.61 0 0 0 .448-1.406l-4.765-28.196a1.61 1.61 0 0 0-2.728-.869l-5.513 5.513l-14.801-14.803a1.61 1.61 0 0 0-2.277.001L59.691 69.967a1.61 1.61 0 0 0 0 2.277l14.802 14.802l-5.515 5.515a1.612 1.612 0 0 0 .876 2.727M30.146 4.712L1.874.022A1.605 1.605 0 0 0 .473.471L.47.474c-.367.37-.534.892-.447 1.406l4.766 28.197a1.61 1.61 0 0 0 1.099 1.264a1.604 1.604 0 0 0 1.629-.395l5.513-5.513l14.8 14.801c.63.63 1.648.63 2.278.001l10.201-10.201a1.61 1.61 0 0 0 0-2.278L25.507 12.954l5.515-5.515a1.617 1.617 0 0 0 .394-1.63a1.608 1.608 0 0 0-1.27-1.097"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Asl(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M25.352 26.045c.347-.208.731-.37 1.031-.632c3.378-2.916 7.1-5.305 11.381-6.591c2.212-.665 4.583-.822 6.897-1.089c.424-.047 1.128.387 1.316.779c.189.398.079 1.195-.222 1.517c-.538.577-1.262 1.097-2.006 1.345c-2.846.951-5.75 1.731-8.589 2.701c-.673.229-1.164.968-1.761 1.445c-1.382 1.101-2.778 2.193-4.166 3.287c-.219.174-.425.366-.637.551l.127.23c.71-.299 1.404-.641 2.133-.883c2.274-.773 4.514-1.815 6.856-2.161c2.021-.3 4.167.168 6.248.39c.484.055.97.439 1.381.766c1.52 1.196 3.004 2.436 4.511 3.656c.518.42.523.824.133 1.371c-.926 1.293-2.098 1.699-3.469.9c-1.089-.635-2.065-1.469-3.058-2.254c-.528-.418-.999-.57-1.722-.454c-1.379.216-2.796.175-4.191.308c-.513.048-1.091.166-1.499.451a157.742 157.742 0 0 0-5.715 4.25c-.352.273-.506.807-.594 1.381c.188-.084.394-.143.561-.256c1.53-1.043 3.307-1.186 5.062-1.407c.559-.071 1.13-.071 1.685-.165c2.269-.383 4.003.707 5.675 2.002c1.375 1.064 2.769 2.104 4.134 3.179c.803.632 1.214 1.458 1.207 2.514c-.013 1.767.097 3.54-.549 5.243c-.076.199-.116.576-.008.66c1.075.823.371 1.738.089 2.547c-.599 1.737-1.318 3.428-1.995 5.137c-.647 1.627-1.637 2.716-3.481 3.239c-4.077 1.161-8.08 2.583-12.104 3.925c-1.645.551-3.251.766-4.998.344c-3.317-.797-6.583-.447-9.708.938c-1.453.646-2.867 1.382-4.349 2.103c-1.701-2.443-2.704-5.113-3.327-7.945c-.735-3.33-.78-6.69-.372-10.061c.035-.291.299-.699.558-.803c3.312-1.352 4.841-4.121 5.937-7.281c1.302-3.749 1.923-7.824 4.223-11.121c2.549-3.656 4.515-7.73 7.852-10.85c1.645-1.539 2.97-3.417 4.431-5.148c.682-.807.996-.795 1.582.094c.668 1.012.507 2.094-.318 3.133c-1.992 2.516-3.942 5.062-5.896 7.605c-.205.271-.318.61-.474.918c.065.062.131.125.198.192m14.361 28.297c.811 0 1.543.219 1.945-.646c.79-1.7 1.992-3.079 3.444-4.241c.432-.348.389-.635.144-1.1c-.704-1.348-1.305-2.75-1.996-4.104c-.194-.374-.504-.747-.858-.975a59.168 59.168 0 0 0-3.762-2.234c-.341-.186-.872-.248-1.231-.125c-.973.338-1.89.834-2.848 1.221c-.865.35-1.766.602-2.628.953c-2.055.836-2.553 2.734-3.158 4.598c-.405 1.242-.131 2.152.758 3.075c1.176 1.218 2.26 2.532 3.317 3.858c.374.47.714.496 1.203.334c1.887-.625 3.804-.985 5.67-.614M74.49 73.967c-.655.49-1.343.944-1.957 1.481c-3.015 2.646-6.487 4.538-10.255 5.753c-2.191.709-4.569.902-6.881 1.139c-1.312.136-2.023-1.223-1.2-2.281c.511-.652 1.351-1.182 2.15-1.454c2.43-.824 4.931-1.438 7.363-2.261c.881-.299 1.69-.891 2.445-1.459c1.571-1.191 3.085-2.458 4.618-3.695c.225-.182.435-.381.529-.783c-.837.357-1.653.771-2.511 1.064c-2.113.729-4.203 1.715-6.382 2.002c-2.069.274-4.243-.164-6.362-.385c-.479-.049-.947-.461-1.365-.777c-1.328-1.01-2.632-2.049-3.939-3.086c-1.345-1.068-1.142-2.455.542-2.986c.656-.207 1.585-.141 2.18.184c1.187.65 2.215 1.592 3.347 2.354c.317.217.769.348 1.151.332c1.466-.06 2.936-.146 4.391-.33c.665-.081 1.394-.277 1.933-.654a109.935 109.935 0 0 0 5.375-4.027c.368-.291.531-.842.692-1.375c-2.368 1.559-5.022 1.568-7.64 1.78c-1.722.138-3.213-.276-4.568-1.384c-1.756-1.439-3.613-2.756-5.358-4.207c-.414-.345-.79-1.022-.781-1.539c.033-1.757.225-3.513.398-5.264c.057-.582.367-1.139-.247-1.656c-.175-.146-.18-.639-.096-.922c.209-.703.522-1.371.778-2.059c.491-1.312.975-2.627 1.452-3.949c.649-1.801 1.736-2.988 3.709-3.551c4.043-1.165 8.02-2.555 12.019-3.873c1.676-.557 3.299-.788 5.092-.312c3.496.927 6.889.274 10.134-1.2c1.244-.568 2.443-1.232 3.732-1.887c1.471 1.793 2.239 3.92 2.917 6.08c1.224 3.896 1.276 7.885.905 11.908c-.082.898-.825 1.02-1.374 1.251c-1.989.841-3.227 2.423-4.062 4.261c-.906 2-1.605 4.108-2.261 6.207c-.919 2.944-1.89 5.848-3.626 8.441c-1.318 1.967-2.508 4.027-3.907 5.933c-1.142 1.556-2.514 2.942-3.776 4.409c-1.347 1.564-2.683 3.138-4.029 4.703c-.716.83-1.061.798-1.65-.145c-.64-1.025-.456-2.031.413-3.117c1.979-2.482 3.914-4.997 5.849-7.514c.2-.259.275-.615.409-.926c-.101-.075-.199-.147-.296-.224M54.765 51.12c.635 1.527 1.283 3.151 1.998 4.74c.157.352.513.662.851.877c1.237.773 2.529 1.465 3.75 2.264c.445.292.741.361 1.247.142c1.493-.655 3.02-1.235 4.544-1.819c1.476-.56 2.797-1.305 3.38-2.89c.299-.813.575-1.642.788-2.479c.233-.914.105-1.72-.615-2.483c-1.28-1.351-2.428-2.819-3.668-4.209c-.159-.182-.572-.297-.799-.225c-2.13.676-4.274.991-6.516.646c-.408-.066-1.043.256-1.321.604c-1.281 1.596-2.47 3.262-3.639 4.832"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Asterisk(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M82.466 60.066c-.01-.006-.022-.009-.032-.015l.034-.059l-17.31-9.994l17.39-10.04l-.005-.009a2.014 2.014 0 0 0 .687-2.72c-.012-.021-.029-.038-.042-.059l.011-.006l-5.52-9.562l-.007.004a2.013 2.013 0 0 0-2.669-.72l-.004-.006l-17.421 10.062V16.793h-.01a2.013 2.013 0 0 0-2.01-1.953c-.025 0-.048.006-.073.007v-.013h-11.04v.007a2.012 2.012 0 0 0-1.957 1.952h-.006v20.115l-17.45-10.075l-.006.01a2.011 2.011 0 0 0-2.697.764c-.012.021-.018.044-.03.065l-.01-.006l-5.52 9.562l.005.003c-.508.94-.196 2.11.712 2.672l-.002.003l.035.02l.016.011l.018.008l17.411 10.052l-17.028 9.832a2.001 2.001 0 0 0-.426.18a2.013 2.013 0 0 0-.743 2.74l-.007.004l5.711 9.892l.032-.019a2.014 2.014 0 0 0 2.55.455c.011-.006.019-.015.03-.021l.032.056l17.369-10.027v19.593a2.007 2.007 0 0 0-.058.461c0 1.108.895 2.005 2 2.014v.007h11.422v-.037a2.012 2.012 0 0 0 1.67-1.98c0-.012-.003-.023-.004-.035h.066V63.056l16.967 9.796c.112.105.233.202.371.282a2.016 2.016 0 0 0 2.745-.724l.005.003l5.534-9.587l.003-.004l.002-.004l.173-.299l-.03-.018a2.017 2.017 0 0 0-.884-2.435"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AtSign(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M87.5 50.002C87.5 29.293 70.712 12.5 50 12.5S12.5 29.293 12.5 50.002C12.5 70.712 29.288 87.5 50 87.5a37.27 37.27 0 0 0 18.342-4.809a1.58 1.58 0 0 0 1.049-1.486c0-.622-.361-1.153-.882-1.413l.003-.004l-6.529-4.002l-.003.004a1.569 1.569 0 0 0-1.005-.369c-.238 0-.461.056-.663.149l-.014-.012A27.372 27.372 0 0 1 50 77.561c-15.199 0-27.56-12.362-27.56-27.559C22.44 34.807 34.802 22.44 50 22.44c14.322 0 26.121 10.984 27.434 24.967C77.428 57.419 73.059 63 69.631 63c-1.847 0-3.254-1.23-3.254-3.957c0-.527.176-1.672.264-2.111l4.163-19.918h-.018c.012-.071.042-.136.042-.21a1.33 1.33 0 0 0-1.33-1.33h-7.23c-.657 0-1.178.485-1.286 1.112l-.025-.001l-.737 3.549c-1.847-3.342-5.629-5.893-10.994-5.893c-10.202 0-19.877 9.764-19.877 21.549c0 8.531 5.101 14.775 13.632 14.775c4.75 0 9.587-2.727 12.665-7.035l.088.527c.615 3.342 9.843 7.576 15.121 7.576c7.651 0 16.617-5.156 16.617-19.932l-.022-.009c.027-.562.05-1.123.05-1.69m-30.885 6.842c-1.935 2.727-5.101 5.805-9.763 5.805c-4.486 0-7.212-3.166-7.212-7.738c0-6.422 5.013-12.754 12.049-12.754c3.958 0 6.245 2.551 7.124 4.486z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BackgroundColor(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M37.808 77.057L34.411 86h32.764l-3.397-8.943zm2.8-8.741h20.369L50.793 41.327z"/><path d="M79 14H21c-3.85 0-7 3.15-7 7v58c0 3.85 3.15 7 7 7h2.111l20.812-52.766a2.595 2.595 0 0 1 2.476-1.836h8.812c1.152 0 2.118.753 2.462 1.79h.018L78.478 86H79c3.85 0 7-3.15 7-7V21c0-3.85-3.15-7-7-7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryEmpty(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M95.582 41.448h-7.796V24.341c0-.956-.775-1.731-1.733-1.731H4.418c-.957 0-1.732.775-1.732 1.731v51.318a1.73 1.73 0 0 0 1.732 1.731h81.635c.957 0 1.732-.775 1.732-1.731v-18.62h7.796c.957 0 1.732-.775 1.732-1.732V43.18a1.73 1.73 0 0 0-1.731-1.732M77.829 67.434H12.862V32.785h64.967z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryFull(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M95.582 41.448h-7.796V24.341c0-.956-.775-1.731-1.733-1.731H4.418c-.957 0-1.732.775-1.732 1.731v51.318a1.73 1.73 0 0 0 1.732 1.731h81.635c.957 0 1.732-.775 1.732-1.731v-18.62h7.796c.957 0 1.732-.775 1.732-1.732V43.18a1.73 1.73 0 0 0-1.731-1.732M77.829 67.434H12.862V32.785h64.967z"/><path fill="currentColor" d="M17.84 37.382h55.01v25.074H17.84z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryHalf(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M95.582 41.448h-7.796V24.341c0-.956-.775-1.731-1.733-1.731H4.418c-.957 0-1.732.775-1.732 1.731v51.318a1.73 1.73 0 0 0 1.732 1.731h81.635c.957 0 1.732-.775 1.732-1.731v-18.62h7.796c.957 0 1.732-.775 1.732-1.732V43.18a1.73 1.73 0 0 0-1.731-1.732M77.829 67.434H12.862V32.785h64.967z"/><path fill="currentColor" d="M56.581 37.635H17.774v24.73h29.105z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bitcoin(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M66.237 47.445c3.667-1.869 5.959-5.162 5.421-10.644c-.723-7.494-7.193-10.003-15.36-10.715l-.004-9.058c.002-.029.017-.054.017-.084c0-.691-.56-1.25-1.251-1.251h-3.813a1.25 1.25 0 0 0-1.249 1.25h-.029l.005 8.872c-1.664 0-3.364.033-5.052.068l-.003-8.83c.003-.038.022-.071.022-.11c0-.691-.56-1.25-1.251-1.251h-3.813a1.25 1.25 0 0 0-1.249 1.25l.001.003h-.031l.002 9.143c-1.37.028-2.715.057-4.027.057l-.001-.033h-8.726v.013a1.25 1.25 0 0 0-1.251 1.251v4.226c0 .691.56 1.251 1.251 1.251l.005-.001v.019s4.672-.091 4.594-.008c2.562.001 3.397 1.488 3.64 2.771l.005 11.843l.001.045l.005 16.587c-.112.806-.586 2.093-2.376 2.096c.081.071-4.599-.001-4.599-.001l-.001.006c-.008 0-.014-.005-.022-.005c-.601 0-1.079.432-1.2.997l-.051-.003l-1.152 5.169l.027.005c-.005.046-.027.086-.027.134c0 .664.522 1.195 1.176 1.236l-.003.019l8.233-.003c1.532 0 3.04.025 4.52.034l.007 9.262h.003c0 .69.56 1.25 1.251 1.25h3.813a1.251 1.251 0 0 0 1.249-1.251h.002l-.004-9.149c1.735.035 3.414.048 5.054.044l.002 9.106h.003c.001.69.56 1.25 1.25 1.25h3.813a1.251 1.251 0 0 0 1.249-1.251v-.001h.008l-.002-9.247c10.635-.615 18.079-3.297 18.999-13.286c.742-8.041-3.043-11.629-9.081-13.075M45.072 33.32c3.571-.002 14.789-1.142 14.793 6.312c.001 7.148-11.218 6.318-14.789 6.32zm.011 32.935l-.003-13.929c4.288 0 17.733-1.238 17.736 6.955c.006 7.856-13.445 6.967-17.733 6.974"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BitcoinCircle(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m49.805 38.076l-2.049 8.215c2.323.579 9.483 2.941 10.643-1.708c1.209-4.847-6.271-5.928-8.594-6.507m-3.082 12.361l-2.26 9.058c2.789.693 11.392 3.455 12.664-1.655c1.329-5.328-7.615-6.707-10.404-7.403"/><path d="M59.067 13.621C38.981 8.612 18.633 20.84 13.626 40.932c-5.011 20.089 7.216 40.438 27.3 45.447c20.092 5.009 40.44-7.217 45.449-27.307c5.009-20.091-7.217-40.443-27.308-45.451m7.462 31.037c-.541 3.653-2.565 5.422-5.254 6.041c3.691 1.921 5.57 4.869 3.78 9.979c-2.22 6.345-7.497 6.881-14.512 5.553l-1.703 6.824l-4.115-1.025l1.68-6.733a156.615 156.615 0 0 1-3.279-.85l-1.686 6.764l-4.11-1.026l1.703-6.836c-.961-.246-1.937-.508-2.933-.757l-5.354-1.336l2.042-4.709s3.032.807 2.991.747c1.165.289 1.681-.471 1.885-.978l2.691-10.787c.151.037.298.073.434.108a3.497 3.497 0 0 0-.427-.137l1.919-7.701c.05-.874-.251-1.976-1.917-2.392c.065-.044-2.988-.743-2.988-.743l1.095-4.395l5.675 1.417l-.005.021c.853.212 1.732.413 2.628.617l1.686-6.757l4.112 1.025l-1.652 6.625c1.104.252 2.215.506 3.297.775l1.641-6.581l4.114 1.025l-1.685 6.76c5.194 1.789 8.993 4.471 8.247 9.462"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Blind(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><circle cx="41.53" cy="19.773" r="7.152"/><path d="m83.162 84.562l.01-.006l-20.92-36.234a2.978 2.978 0 0 0 .577-1.756a3 3 0 0 0-1.887-2.783L54.907 40.3l-9.159-9.159a7.135 7.135 0 0 0-5.552-2.647a7.125 7.125 0 0 0-5.152 2.199L24.412 41.324l.012.012a.719.719 0 0 0-.285.564c0 .056.02.104.032.156h-.029v11.975c0 .009-.003.017-.003.025a3.003 3.003 0 0 0 6.005 0l-.002-.024v-10.01l3.649-3.649V59.18h.001l-6.954 12.044l-8.891 8.891l.033.033a4.16 4.16 0 0 0-1.339 3.052a4.18 4.18 0 0 0 4.181 4.18c1.317 0 2.476-.621 3.243-1.571l9.596-9.596l-.037-.037l8.096-14.024l7.104 7.104l9.232 15.99l.036-.021c.715 1.278 2.066 2.153 3.635 2.153a4.18 4.18 0 0 0 4.181-4.18c0-.915-.302-1.754-.8-2.443l-9.627-16.674l-.012.007l-8.679-8.678V40.609l4.451 4.451l-.001.002l.007.004l.025.025l.007-.007l6.847 3.954c.054.037.113.064.169.098l.028.016v-.001c.443.258.952.417 1.502.417c.189 0 .372-.022.552-.056l20.816 36.054c.013.032.034.057.05.087l.007.012h.001a1.117 1.117 0 0 0 2.107-.511a1.066 1.066 0 0 0-.196-.592"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bluetooth(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M71.782 64.059L57.807 50.085l14.011-14.011a4.25 4.25 0 0 0 .455-.455l.013-.013l-.001-.001a4.212 4.212 0 0 0 1.028-2.742c0-1.307-.603-2.461-1.531-3.241L52.156 9.995c-.779-.978-1.965-1.617-3.312-1.617a4.25 4.25 0 0 0-4.25 4.25c0 .407.076.793.183 1.167v23.258L34.052 26.329a4.224 4.224 0 0 0-3.115-1.377a4.25 4.25 0 0 0-4.25 4.25c0 1.238.538 2.342 1.383 3.119l16.706 16.706v2.115L27.953 67.966l-.038.038l-.032.032l.003.003a4.232 4.232 0 0 0-1.199 2.949a4.25 4.25 0 0 0 4.25 4.25c1.289 0 2.429-.586 3.209-1.491l10.63-10.63v24.099c-.002.053-.016.102-.016.156s.014.103.016.156v.022h.002a4.24 4.24 0 0 0 4.232 4.072a4.221 4.221 0 0 0 3.156-1.428l.004.004l19.862-19.863l.064-.064l.19-.19l-.018-.018a4.213 4.213 0 0 0 1.045-2.762c-.001-1.308-.603-2.462-1.531-3.242m-8.711 3.263l-9.828 9.828V57.494zm-9.828-24.647V23.057l9.809 9.809z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bold(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M62.73 49.109c5.347-1.103 9.76-5.94 9.76-12.985c0-7.553-5.517-14.428-16.295-14.428H29.011a2.604 2.604 0 0 0-2.604 2.604v51.399a2.604 2.604 0 0 0 2.604 2.604h28.118c10.863 0 16.464-6.79 16.464-15.361c.001-7.043-4.752-12.9-10.863-13.833M38.458 32.305h15.107c4.074 0 6.62 2.461 6.62 5.94c0 3.649-2.546 5.941-6.62 5.941H38.458zm15.615 35.39H38.458v-12.9h15.616c4.668 0 7.214 2.886 7.214 6.45c0 4.074-2.716 6.45-7.215 6.45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Book(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M82.063 22.284h-.006a1.641 1.641 0 0 0-1.612-1.589v-.003h-3.36l-.001 60.404l.001.01c0 .9-.721 1.627-1.617 1.644v.005l-50.906-.002a1.644 1.644 0 0 1-1.619-1.375v-3.602h47.405v-.003c.011 0 .021.003.031.003c.91 0 1.646-.735 1.646-1.646l-.002-.012l.002-62.202c0-.91-.736-1.647-1.646-1.646l-.03.003v-.004l-50.766.001c-.91 0-1.646.736-1.646 1.646c0 .021.006.042.006.063v72.044c-.001.02-.006.039-.006.059s.005.039.006.059v.076h.008a1.64 1.64 0 0 0 1.632 1.511v.001l60.833.002h.029v-.003a1.643 1.643 0 0 0 1.617-1.643l-.001-.01zm-59.147 1.588c0-.908.736-1.646 1.646-1.646h40.84c.909 0 1.646.738 1.646 1.646l-.001 12.361v.002c0 .91-.737 1.646-1.646 1.646H24.562a1.645 1.645 0 0 1-1.646-1.646l.001-12.354z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookBookmark(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M82.063 22.284h-.006a1.641 1.641 0 0 0-1.612-1.589v-.003h-3.36l-.001 60.404l.001.01c0 .9-.721 1.627-1.617 1.644v.005l-50.906-.002a1.644 1.644 0 0 1-1.619-1.375v-3.602h47.405v-.003c.011 0 .021.003.031.003c.91 0 1.646-.735 1.646-1.646l-.002-.012l.002-62.202c0-.91-.736-1.647-1.646-1.646l-.03.003v-.004h-5.756v34.952a.782.782 0 0 1-1.36.53l-5.667-5.667h-.019a.677.677 0 0 0-.513-.241a.677.677 0 0 0-.513.241h-.019l-5.709 5.709h-.009a.78.78 0 0 1-1.318-.571V12.27H19.583c-.91 0-1.646.736-1.646 1.646c0 .021.006.042.006.063v72.044c-.001.02-.006.039-.006.059s.005.039.006.059v.076h.008a1.64 1.64 0 0 0 1.632 1.511v.001l60.833.002h.029v-.003a1.643 1.643 0 0 0 1.617-1.643l-.001-.01z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookmark(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M68.312 13.111H31.687a3.007 3.007 0 0 0-3.007 3.007v68.557c0 1.224.992 2.215 2.216 2.215a2.2 2.2 0 0 0 1.509-.602h.014l16.094-16.094h.035c.355-.414.876-.682 1.465-.682s1.109.268 1.465.682h.035l15.918 15.918a2.2 2.2 0 0 0 1.673.777a2.216 2.216 0 0 0 2.216-2.215V16.118a3.007 3.007 0 0 0-3.008-3.007"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Braille(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="31.102" cy="23.979" r="6.978" fill="currentColor"/><circle cx="68.898" cy="23.979" r="6.978" fill="currentColor"/><circle cx="31.102" cy="50" r="6.978" fill="currentColor"/><circle cx="68.898" cy="50" r="6.978" fill="currentColor"/><circle cx="31.102" cy="76.02" r="6.978" fill="currentColor"/><circle cx="68.898" cy="76.02" r="6.978" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Burst(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M88.558 49.96c0-.885-.435-1.663-1.097-2.151l.014-.024l-9.324-5.383l5.367-9.296l-.018-.011a2.666 2.666 0 0 0-.127-2.408a2.667 2.667 0 0 0-2.025-1.314v-.026H70.58V18.61h-.022a2.667 2.667 0 0 0-1.314-2.022a2.662 2.662 0 0 0-2.412-.125l-.013-.023l-9.481 5.474l-5.25-9.094l-.019.011a2.668 2.668 0 0 0-2.149-1.094c-.885 0-1.664.435-2.151 1.097l-.024-.014l-5.337 9.244l-9.19-5.306l-.011.019a2.666 2.666 0 0 0-2.408.127a2.666 2.666 0 0 0-1.315 2.025h-.027v10.674H18.845v.021a2.667 2.667 0 0 0-2.022 1.314a2.667 2.667 0 0 0-.126 2.41l-.023.014l5.246 9.087l-9.394 5.424l.011.019a2.668 2.668 0 0 0-1.094 2.149c0 .885.435 1.664 1.097 2.151l-.014.024l9.324 5.383l-5.367 9.296l.019.011a2.666 2.666 0 0 0 .127 2.408a2.667 2.667 0 0 0 2.025 1.314v.027H29.42V81.39h.022c.092.816.549 1.58 1.314 2.022a2.665 2.665 0 0 0 2.412.125l.013.023l9.481-5.474l5.25 9.094l.019-.011a2.668 2.668 0 0 0 2.149 1.094c.885 0 1.664-.435 2.151-1.096l.023.013l5.337-9.244l9.191 5.306l.011-.019a2.666 2.666 0 0 0 2.408-.127a2.666 2.666 0 0 0 1.315-2.025h.027V70.398h10.613v-.021a2.667 2.667 0 0 0 2.022-1.314a2.67 2.67 0 0 0 .126-2.411l.023-.013l-5.246-9.087l9.394-5.424l-.011-.019a2.666 2.666 0 0 0 1.094-2.149"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BurstNew(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M88.558 49.96c0-.885-.435-1.663-1.097-2.151l.014-.024l-9.324-5.383l5.367-9.296l-.018-.011a2.666 2.666 0 0 0-.127-2.408a2.667 2.667 0 0 0-2.025-1.314v-.026H70.58V18.61h-.022a2.667 2.667 0 0 0-1.314-2.022a2.662 2.662 0 0 0-2.412-.125l-.013-.023l-9.481 5.474l-5.25-9.094l-.019.011a2.668 2.668 0 0 0-2.149-1.094c-.885 0-1.664.435-2.151 1.097l-.024-.014l-5.337 9.244l-9.19-5.306l-.011.019a2.666 2.666 0 0 0-2.408.127a2.666 2.666 0 0 0-1.315 2.025h-.027v10.674H18.845v.021a2.667 2.667 0 0 0-2.022 1.314a2.667 2.667 0 0 0-.126 2.41l-.023.014l5.246 9.087l-9.394 5.424l.011.019a2.668 2.668 0 0 0-1.094 2.149c0 .885.435 1.664 1.097 2.151l-.014.024l9.324 5.383l-5.367 9.296l.018.01a2.666 2.666 0 0 0 .127 2.408a2.667 2.667 0 0 0 2.025 1.314v.027H29.42V81.39h.022c.092.816.549 1.58 1.314 2.022a2.665 2.665 0 0 0 2.412.125l.013.023l9.481-5.474l5.25 9.094l.019-.011a2.668 2.668 0 0 0 2.149 1.094c.885 0 1.664-.435 2.151-1.096l.023.013l5.337-9.244l9.191 5.306l.011-.019a2.666 2.666 0 0 0 2.408-.127a2.666 2.666 0 0 0 1.315-2.025h.027V70.398h10.613v-.021a2.667 2.667 0 0 0 2.022-1.314a2.67 2.67 0 0 0 .126-2.411l.023-.013l-5.246-9.087l9.394-5.424l-.011-.019a2.666 2.666 0 0 0 1.094-2.149M43.715 61.355l-9.846-4.35l4.345 7.525l-2.456 1.418l-6.662-11.537l2.525-1.459l9.53 4.162l-4.185-7.248l2.457-1.418l6.66 11.537zm4.652-2.686l-6.661-11.538l8.165-4.713l1.248 2.162l-5.709 3.295l1.398 2.422l5.587-3.225l1.248 2.16l-5.587 3.227l1.518 2.629l5.709-3.295l1.248 2.162zm18.906-10.915L60.675 41l2.567 9.08l-2.611 1.508l-9.965-9.629l2.75-1.588l6.838 7.168l-2.617-9.605l1.92-1.108l6.993 7.079l-2.79-9.506l2.75-1.588l3.375 13.436z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BurstSale(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m44.105 55.641l3.598-2.079l-4.666-3.925z"/><path fill="currentColor" d="M88.558 49.96c0-.885-.435-1.663-1.097-2.151l.014-.024l-9.324-5.383l5.367-9.296l-.018-.011a2.666 2.666 0 0 0-.127-2.408a2.667 2.667 0 0 0-2.025-1.314v-.026H70.58V18.61h-.022a2.667 2.667 0 0 0-1.314-2.022a2.662 2.662 0 0 0-2.412-.125l-.013-.023l-9.481 5.474l-5.25-9.094l-.019.011a2.668 2.668 0 0 0-2.149-1.094c-.885 0-1.664.435-2.151 1.097l-.024-.014l-5.337 9.244l-9.19-5.306l-.011.019a2.666 2.666 0 0 0-2.408.127a2.666 2.666 0 0 0-1.315 2.025h-.027v10.674H18.845v.021a2.667 2.667 0 0 0-2.022 1.314a2.667 2.667 0 0 0-.126 2.41l-.023.014l5.246 9.087l-9.394 5.424l.011.019a2.668 2.668 0 0 0-1.094 2.149c0 .885.435 1.664 1.097 2.151l-.014.024l9.324 5.383l-5.367 9.296l.018.01a2.666 2.666 0 0 0 .127 2.408a2.667 2.667 0 0 0 2.025 1.314v.027H29.42V81.39h.022c.092.816.549 1.58 1.314 2.022a2.665 2.665 0 0 0 2.412.125l.013.023l9.481-5.474l5.25 9.094l.019-.011a2.668 2.668 0 0 0 2.149 1.094c.885 0 1.664-.435 2.151-1.096l.023.013l5.337-9.244l9.191 5.306l.011-.019a2.666 2.666 0 0 0 2.408-.127a2.664 2.664 0 0 0 1.315-2.025h.027V70.398h10.613v-.021a2.667 2.667 0 0 0 2.022-1.314a2.667 2.667 0 0 0 .126-2.41l.023-.013l-5.246-9.087l9.394-5.424l-.011-.019a2.67 2.67 0 0 0 1.094-2.15M37.537 65.197c-2.23 1.288-4.252 1.464-5.971 1.002l.241-2.697c1.302.377 2.985.375 4.575-.544c1.367-.789 1.658-1.765 1.269-2.438c-1.159-2.006-6.992 3.23-9.499-1.111c-1.108-1.92-.367-4.471 2.35-6.039c1.833-1.059 3.675-1.383 5.426-.988l-.309 2.623c-1.433-.324-2.908-.004-4.084.674c-1.038.6-1.367 1.389-.967 2.082c1.049 1.816 6.965-3.236 9.451 1.069c1.219 2.109.616 4.58-2.482 6.367m13.943-8.326l-1.854-1.535l-4.947 2.856l.401 2.374l-2.785 1.607L40.08 48.07l3.079-1.777l11.106 8.971zm3.753-2.166l-6.661-11.538l2.474-1.429l5.413 9.375l4.878-2.816l1.249 2.163zm9.012-5.203l-6.661-11.537l8.164-4.715l1.248 2.162l-5.707 3.296l1.398 2.422l5.586-3.226l1.248 2.162l-5.586 3.225l1.518 2.63l5.708-3.296l1.249 2.162z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M80.583 39.494H19.417c-.699 0-1.266.565-1.266 1.264v43.916c0 .698.567 1.264 1.266 1.264h61.165c.703 0 1.267-.566 1.267-1.264V40.758c0-.699-.564-1.264-1.266-1.264m-41.21 37.185c-2.111 0-4.443-.424-6.243-1.136a.632.632 0 0 1-.396-.674l.59-4.294a.624.624 0 0 1 .3-.456a.64.64 0 0 1 .548-.049c1.979.747 3.643 1.053 5.744 1.053c2.202 0 3.736-1.04 3.736-2.528c0-1.847-.891-2.835-5.762-3.113a.63.63 0 0 1-.595-.631v-4.293c0-.332.255-.608.585-.632c4.195-.296 4.68-1.586 4.68-2.53c0-.639 0-1.709-2.84-1.709c-1.519 0-3.132.414-4.792 1.226a.63.63 0 0 1-.904-.481l-.588-4.292a.634.634 0 0 1 .359-.661c1.933-.894 4.372-1.347 7.254-1.347c5.137 0 8.589 2.386 8.589 5.937c0 2.663-1.377 4.623-4.202 5.955c2.425.935 4.903 2.632 4.903 6.222c0 5.122-4.302 8.433-10.966 8.433m26.718-1.023a.633.633 0 0 1-.633.632h-5.227a.632.632 0 0 1-.633-.632V58.434l-3.435 1.31a.628.628 0 0 1-.548-.048a.618.618 0 0 1-.303-.457l-.585-4.294a.629.629 0 0 1 .346-.651l7.493-3.706a.655.655 0 0 1 .282-.065h2.611c.349 0 .633.283.633.633zM80.583 19.72h-6.519v3.051c0 3.235-1.845 6.69-7.03 6.69c-5.185 0-7.031-3.455-7.031-6.69V19.72H40.08v3.051c0 3.235-1.845 6.69-7.03 6.69c-5.186 0-7.036-3.455-7.036-6.69V19.72h-6.597c-.699 0-1.266.565-1.266 1.264v12.859c0 .7.567 1.266 1.266 1.266h61.165c.703 0 1.267-.566 1.267-1.266V20.984c0-.699-.564-1.264-1.266-1.264"/><path fill="currentColor" d="M33.05 25.376c2.095 0 2.946-.755 2.946-2.606v-6.1c0-1.853-.851-2.607-2.946-2.607c-2.101 0-2.952.755-2.952 2.607v6.101c0 1.851.851 2.605 2.952 2.605m33.983 0c2.095 0 2.946-.755 2.946-2.606v-6.1c0-1.853-.851-2.607-2.946-2.607c-2.101 0-2.946.755-2.946 2.607v6.101c-.001 1.851.845 2.605 2.946 2.605"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Camera(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="50.041" cy="53.256" r="9.513" fill="currentColor"/><path fill="currentColor" d="M86.778 24.2v-.037H70.111v-2.792a3.722 3.722 0 0 0-3.722-3.722H33.611a3.722 3.722 0 0 0-3.723 3.722v2.792H13.221v.037A3.72 3.72 0 0 0 9.5 27.922v50.669a3.72 3.72 0 0 0 3.721 3.722v.037h73.557v-.037a3.723 3.723 0 0 0 3.722-3.722V27.922a3.722 3.722 0 0 0-3.722-3.722M50.247 72.542c-10.651 0-19.286-8.634-19.286-19.286c0-10.651 8.635-19.286 19.286-19.286c10.652 0 19.287 8.635 19.287 19.286c.001 10.652-8.635 19.286-19.287 19.286m35.266-33.397h-14.74v-9.957h14.739z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Check(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M88.04 30.319L75.124 17.401a2.42 2.42 0 0 0-3.419 0L37.392 51.714l-9.094-9.093a2.418 2.418 0 0 0-3.419 0L11.96 55.539a2.419 2.419 0 0 0 0 3.419L35.607 82.6a2.416 2.416 0 0 0 1.709.708c.029 0 .055-.016.083-.016c.024 0 .05.014.075.014c.621 0 1.236-.236 1.709-.708l48.857-48.86a2.416 2.416 0 0 0 0-3.419"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Checkbox(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m92.038 24.333l-8.62-8.622a1.615 1.615 0 0 0-2.282 0L49.987 46.86l-6.07-6.068a1.614 1.614 0 0 0-2.282 0l-8.622 8.622a1.611 1.611 0 0 0 0 2.282l15.782 15.778c.302.302.712.473 1.141.473c.019 0 .037-.01.056-.01c.016 0 .033.009.05.009a1.61 1.61 0 0 0 1.141-.473l40.855-40.857c.63-.632.63-1.653 0-2.283"/><path fill="currentColor" d="M72.022 53.625v21.159H27.978V30.74h31.06l9.979-9.978H23.193v.007c-.023 0-.044-.007-.068-.007a5.118 5.118 0 0 0-5.113 5H18v54h.013A5.111 5.111 0 0 0 23 84.749v.013h54v-.013a5.11 5.11 0 0 0 4.987-4.987H82V43.647z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clipboard(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M80.761 17.26h-.019c.001-.049.015-.095.015-.145a5.245 5.245 0 0 0-5.241-5.245c-.04 0-.076.011-.116.012H60.842V8.478c0-.881-.714-1.593-1.593-1.595v-.008H40.8v.006c-.015 0-.028-.004-.043-.004c-.88 0-1.594.713-1.594 1.594l.003.03v3.38H24.501l-.02-.002a5.248 5.248 0 0 0-5.242 5.243c0 .047.013.09.014.137h-.014v70.572h.003l-.003.027a5.24 5.24 0 0 0 5.243 5.238c.078 0 .151-.02.229-.023v.021h50.5c.098.005.191.029.29.029a5.24 5.24 0 0 0 5.239-5.238c0-.019-.005-.036-.005-.055h.028V17.26zm-9.939 65.925H29.178V21.809h4.997v3.069c-.002.033-.01.063-.01.097c0 .998.81 1.814 1.814 1.814l.014-.001v.008H64.02a1.815 1.815 0 0 0 1.813-1.813c0-.015-.004-.029-.004-.044v-3.13h4.992v61.376z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardNotes(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M80.761 17.26h-.019c.001-.049.015-.095.015-.145a5.245 5.245 0 0 0-5.241-5.245c-.04 0-.076.011-.116.012H60.842V8.478c0-.881-.714-1.593-1.593-1.595v-.008H40.8v.006c-.015 0-.028-.004-.043-.004c-.88 0-1.594.713-1.594 1.594l.003.03v3.38H24.501l-.02-.002a5.248 5.248 0 0 0-5.242 5.243c0 .047.013.09.014.137h-.014v70.572h.003l-.003.027a5.24 5.24 0 0 0 5.243 5.238c.078 0 .151-.02.229-.023v.021h50.5c.098.005.191.029.29.029a5.24 5.24 0 0 0 5.239-5.238c0-.019-.005-.036-.005-.055h.028V17.26zm-9.939 65.925H29.178V21.809h4.997v3.069c-.002.033-.01.063-.01.097c0 .998.81 1.814 1.814 1.814l.014-.001v.008H64.02a1.815 1.815 0 0 0 1.813-1.813c0-.015-.004-.029-.004-.044v-3.13h4.992v61.376z"/><path fill="currentColor" d="M39.15 41.119c0-.658-.537-1.195-1.198-1.195c-.031 0-.058.015-.088.018v-.018h-2.502a1.2 1.2 0 0 0-1.199 1.197v2.597c0 .656.537 1.193 1.199 1.193h2.503v-.018c.03.002.056.018.086.018c.663 0 1.2-.537 1.2-1.193v-2.597h-.001zm26.681 0c0-.658-.537-1.195-1.198-1.195H45.337a1.2 1.2 0 0 0-1.199 1.197v2.597c0 .656.537 1.193 1.199 1.193h19.295c.663 0 1.2-.537 1.2-1.193v-2.597h-.001zM39.15 51.094c0-.658-.537-1.195-1.198-1.195c-.031 0-.058.015-.088.018v-.018h-2.502a1.2 1.2 0 0 0-1.199 1.197v2.597c0 .656.537 1.193 1.199 1.193h2.503v-.018c.03.002.056.018.086.018c.663 0 1.2-.537 1.2-1.193v-2.597h-.001zm26.681 0c0-.658-.537-1.195-1.198-1.195H45.337a1.2 1.2 0 0 0-1.199 1.197v2.597c0 .656.537 1.193 1.199 1.193h19.295c.663 0 1.2-.537 1.2-1.193v-2.597h-.001zM39.15 61.069c0-.658-.537-1.195-1.198-1.195c-.031 0-.058.015-.088.018v-.018h-2.502a1.2 1.2 0 0 0-1.199 1.197v2.597c0 .656.537 1.193 1.199 1.193h2.503v-.018c.03.002.056.018.086.018c.663 0 1.2-.537 1.2-1.193v-2.597h-.001zm26.681 0c0-.658-.537-1.195-1.198-1.195H45.337a1.2 1.2 0 0 0-1.199 1.197v2.597c0 .656.537 1.193 1.199 1.193h19.295c.663 0 1.2-.537 1.2-1.193v-2.597h-.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardPencil(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m87.969 43.045l-8.253-8.25a1.227 1.227 0 0 0-1.738 0l-22.583 22.58l.045.045l-.072-.02l-3.091 11.534l.009.002a1.21 1.21 0 0 0 1.669 1.486l.002.009l11.415-3.059l-.007-.027l.019.019l22.584-22.583a1.23 1.23 0 0 0 .001-1.736M56.035 66.697l1.836-6.846l5.012 5.012z"/><path fill="currentColor" d="M71.967 73.23h-7.485a1.23 1.23 0 0 0-1.228 1.229l.002.012v8.715H21.612V21.809h4.997v3.069c-.002.033-.01.063-.01.097c0 .998.81 1.814 1.814 1.814l.014-.001v.008h28.027a1.815 1.815 0 0 0 1.813-1.813c0-.015-.004-.029-.004-.044v-3.13h4.992v8.718l-.002.012a1.227 1.227 0 0 0 1.228 1.229h7.485a1.229 1.229 0 0 0 1.228-1.228V17.26h-.019c.001-.049.015-.095.015-.145a5.245 5.245 0 0 0-5.241-5.245c-.04 0-.076.011-.116.012H53.276V8.478c0-.881-.714-1.593-1.593-1.595v-.008h-18.45v.006c-.015 0-.028-.004-.043-.004c-.88 0-1.594.713-1.594 1.594l.003.03v3.38H16.935l-.02-.002a5.248 5.248 0 0 0-5.242 5.243c0 .047.013.09.014.137h-.014v70.572h.003l-.003.027a5.24 5.24 0 0 0 5.243 5.238c.078 0 .151-.02.229-.023v.021h50.5c.098.005.191.029.29.029a5.24 5.24 0 0 0 5.239-5.238c0-.019-.005-.036-.005-.055h.028V74.458a1.229 1.229 0 0 0-1.23-1.228"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clock(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M50 22.44c15.198 0 27.56 12.367 27.56 27.562c0 15.197-12.362 27.559-27.56 27.559c-15.199 0-27.561-12.362-27.561-27.559C22.439 34.806 34.801 22.44 50 22.44m0-9.94c-20.712 0-37.5 16.792-37.5 37.502S29.288 87.5 50 87.5s37.5-16.788 37.5-37.498C87.5 29.292 70.712 12.5 50 12.5"/><path fill="currentColor" d="m69.195 36.068l-3.897-3.902c-.743-.747-2.077-.729-2.791 0L50.022 44.654l-6.863-6.863c-.743-.743-2.046-.743-2.789 0l-3.892 3.893a1.967 1.967 0 0 0-.585 1.402c0 .525.204 1.025.578 1.394l12.133 12.133c.374.374.869.578 1.396.578h.078a1.96 1.96 0 0 0 1.364-.578l17.754-17.754a1.976 1.976 0 0 0-.001-2.791"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClosedCaption(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m81.615 28.882l-31.178-6.257a2.185 2.185 0 0 0-.875 0l-31.176 6.257a2.188 2.188 0 0 0-1.764 2.136v37.966c0 1.036.738 1.928 1.764 2.135l31.176 6.254c.289.062.588.062.875 0l31.178-6.254a2.187 2.187 0 0 0 1.764-2.135V31.018a2.19 2.19 0 0 0-1.764-2.136M45.643 58.113c-2.142 4.72-5.503 7.108-9.991 7.108c-7.493 0-12.525-6.117-12.525-15.219c0-9.107 5.033-15.225 12.525-15.225c4.462 0 7.803 2.264 9.936 6.729c.191.396.21.854.055 1.265c-.151.417-.467.75-.874.929l-3.621 1.616a1.668 1.668 0 0 1-2.191-.821c-.54-1.172-1.552-2.576-3.305-2.576c-2.774 0-4.709 3.324-4.709 8.084c0 4.758 1.935 8.08 4.709 8.08c1.986 0 2.942-1.953 3.263-2.792c.159-.421.485-.764.904-.934a1.654 1.654 0 0 1 1.309.005l3.663 1.576c.407.172.725.5.885.912a1.62 1.62 0 0 1-.033 1.263m31.083 0c-2.142 4.72-5.501 7.108-9.991 7.108c-7.493 0-12.525-6.117-12.525-15.219c0-9.107 5.033-15.225 12.525-15.225c4.462 0 7.803 2.264 9.936 6.729c.191.396.209.854.055 1.265c-.151.417-.467.75-.874.929l-3.621 1.616a1.669 1.669 0 0 1-2.191-.821c-.54-1.172-1.552-2.576-3.305-2.576c-2.774 0-4.71 3.324-4.71 8.084c0 4.758 1.935 8.08 4.71 8.08c1.986 0 2.942-1.953 3.263-2.792a1.63 1.63 0 0 1 .904-.934a1.654 1.654 0 0 1 1.309.005l3.663 1.576c.407.172.726.5.885.912c.162.408.149.861-.033 1.263"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloud(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M79.437 47.349c-3.51-12.982-14.464-22.083-27.134-22.083c-8.835 0-17.065 4.454-22.414 12.018a20.179 20.179 0 0 0-4.501-.514c-11.889 0-21.563 10.539-21.563 23.498c0 4.647 1.251 9.148 3.612 13.018c.555.906 1.49 1.449 2.49 1.449h80.84c.947 0 1.836-.485 2.403-1.315a17.249 17.249 0 0 0 3.004-9.767c.001-9.658-7.814-17.473-16.737-16.304"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Comment(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M88.466 16.936a3.362 3.362 0 0 0-3.34-3.036H14.781v.009a3.356 3.356 0 0 0-3.247 3.027H11.5v56.342h.068a3.37 3.37 0 0 0 3.214 2.694v.009h11.564v6.744a3.373 3.373 0 0 0 6.172 1.884l8.629-8.629h43.979a3.373 3.373 0 0 0 3.306-2.703h.068V16.936z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentMinus(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M63.185 55.07a1.73 1.73 0 0 1-1.73-1.73V36.445c0-.956.774-1.73 1.73-1.73h17.962V16.936h-.034a3.362 3.362 0 0 0-3.34-3.036H7.429v.009a3.356 3.356 0 0 0-3.247 3.027h-.035v56.342h.068a3.37 3.37 0 0 0 3.214 2.694v.009h11.564v6.744a3.373 3.373 0 0 0 6.173 1.884l8.628-8.628h43.98a3.373 3.373 0 0 0 3.306-2.703h.068V55.07z"/><path fill="currentColor" d="M94.123 39.702h-25.95a1.73 1.73 0 0 0-1.731 1.73v6.92c0 .956.774 1.73 1.731 1.73h25.95a1.73 1.73 0 0 0 1.73-1.73v-6.92a1.73 1.73 0 0 0-1.73-1.73"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentQuotes(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M88.466 16.936a3.362 3.362 0 0 0-3.34-3.036H14.781v.009a3.356 3.356 0 0 0-3.247 3.027H11.5v56.342h.068a3.37 3.37 0 0 0 3.214 2.694v.009h11.564v6.744a3.373 3.373 0 0 0 6.173 1.882l8.627-8.627h43.98a3.373 3.373 0 0 0 3.306-2.703h.068V16.936zM28.399 50.662c-3.055 0-5.787-2.405-5.787-6.632c0-4.878 2.667-9.234 6.699-12.095l.014.012a.713.713 0 0 1 .324-.085a.71.71 0 0 1 .43.156l.008-.007l2.703 1.706l-.016.015c.192.131.325.34.325.59a.719.719 0 0 1-.39.633l.002.002l-.046.028l-.012.007c-2.065 1.124-4.26 3.726-4.839 6.046c.26-.13.78-.195 1.301-.195c2.471 0 4.421 1.885 4.421 4.681c.001 2.797-2.341 5.138-5.137 5.138m13.233 0c-3.055 0-5.787-2.405-5.787-6.632c0-4.878 2.667-9.234 6.699-12.095l.014.012a.713.713 0 0 1 .324-.085a.71.71 0 0 1 .43.156l.008-.007l2.703 1.706l-.016.015c.192.131.325.34.325.59a.719.719 0 0 1-.39.633l.002.002l-.046.028l-.012.007c-2.065 1.124-4.26 3.726-4.839 6.046c.26-.13.78-.195 1.301-.195c2.471 0 4.421 1.885 4.421 4.681c0 2.797-2.341 5.138-5.137 5.138m15.824 7.187l-.014-.012a.713.713 0 0 1-.324.085a.71.71 0 0 1-.43-.156l-.008.007l-2.703-1.706l.016-.015a.713.713 0 0 1-.325-.59c0-.277.161-.511.39-.633l-.002-.002l.046-.028l.012-.007c2.065-1.124 4.26-3.726 4.839-6.046c-.26.13-.78.195-1.301.195c-2.471 0-4.421-1.885-4.421-4.681c0-2.797 2.342-5.137 5.138-5.137c3.055 0 5.787 2.405 5.787 6.632c-.001 4.877-2.668 9.233-6.7 12.094m13.233 0l-.014-.012a.713.713 0 0 1-.324.085a.71.71 0 0 1-.43-.156l-.008.007l-2.703-1.706l.016-.015a.713.713 0 0 1-.325-.59c0-.277.161-.511.39-.633l-.002-.002l.046-.028l.012-.007c2.065-1.124 4.26-3.726 4.839-6.046c-.26.13-.78.195-1.301.195c-2.471 0-4.421-1.885-4.421-4.681c0-2.797 2.342-5.137 5.138-5.137c3.055 0 5.787 2.405 5.787 6.632c-.001 4.877-2.668 9.233-6.7 12.094"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentVideo(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M88.466 16.936a3.362 3.362 0 0 0-3.34-3.036H14.781v.009a3.357 3.357 0 0 0-3.247 3.027H11.5v56.342h.068a3.37 3.37 0 0 0 3.213 2.694v.009h11.565v6.744a3.373 3.373 0 0 0 6.173 1.883l8.628-8.628h43.98a3.373 3.373 0 0 0 3.306-2.703h.067V16.936zM70.974 57.175L58.28 51.172v5.653a2.434 2.434 0 0 1-2.428 2.426H34.374a2.434 2.434 0 0 1-2.427-2.426V33.39a2.433 2.433 0 0 1 2.427-2.425h21.478a2.434 2.434 0 0 1 2.428 2.425v5.503l12.694-5.948z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Comments(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M94.974 17.838a2.526 2.526 0 0 0-2.443-2.277v-.007H39.624a2.53 2.53 0 0 0-2.513 2.284h-.026v13.139h22.342c1.174 0 2.127.896 2.243 2.039h.023v29.232h11.009l6.488 6.487a2.537 2.537 0 0 0 4.644-1.415v-5.072h8.698v-.008a2.534 2.534 0 0 0 2.417-2.027H95V17.838z"/><path fill="currentColor" d="M54.44 35.964H7.204v.006a2.256 2.256 0 0 0-2.181 2.033H5v37.835h.046a2.264 2.264 0 0 0 2.158 1.809v.006h7.765v4.529a2.266 2.266 0 0 0 4.146 1.263l5.792-5.792H54.44c1.096 0 2.011-.78 2.22-1.815h.045V38.003h-.023a2.255 2.255 0 0 0-2.242-2.039"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compass(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M50 12.5c-20.712 0-37.5 16.793-37.5 37.502C12.5 70.712 29.288 87.5 50 87.5s37.5-16.788 37.5-37.498C87.5 29.293 70.712 12.5 50 12.5m0 65.06c-15.199 0-27.56-12.362-27.56-27.558C22.44 34.807 34.801 22.44 50 22.44c15.198 0 27.56 12.367 27.56 27.562C77.56 65.198 65.198 77.56 50 77.56"/><path d="M64.674 32.69L45.129 44.338a1.844 1.844 0 0 0-.646.641L32.69 64.67a1.865 1.865 0 0 0 .291 2.285l.078.078c.359.354.833.539 1.311.539c.328 0 .66-.088.956-.268l19.69-11.789c.262-.16.485-.383.641-.645l11.648-19.545c.437-.729.245-1.738-.349-2.345a1.857 1.857 0 0 0-2.282-.29"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Contrast(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M50 12.5c-20.712 0-37.5 16.793-37.5 37.502C12.5 70.712 29.288 87.5 50 87.5s37.5-16.788 37.5-37.498C87.5 29.293 70.712 12.5 50 12.5m.124 9.943c15.141.067 27.436 12.405 27.436 27.559c0 15.155-12.295 27.488-27.436 27.555z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCard(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M86.941 26.224a2.66 2.66 0 0 0-2.66-2.659c-.036 0-.07.009-.106.011H15.949c-.078-.007-.153-.023-.233-.023a2.659 2.659 0 0 0-2.658 2.659c0 .124.02.243.037.363v7.009h73.846v-7.243h-.011c.001-.041.011-.078.011-.117M13.095 73.78v.01a2.659 2.659 0 0 0 2.659 2.658c.056 0 .109-.013.164-.017v.002h68.459v-.02a2.653 2.653 0 0 0 2.563-2.633h.002V43.582H13.095z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crop(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M88.231 71.062h-7.264V23.959h-.051c.003-.07.021-.135.021-.206a4.773 4.773 0 0 0-4.772-4.772H28.99v-7.068a2.693 2.693 0 0 0-2.693-2.693h-4.571a2.692 2.692 0 0 0-2.693 2.693v7.068h-7.264a2.693 2.693 0 0 0-2.693 2.693v4.571a2.692 2.692 0 0 0 2.693 2.693h7.264v47.103h.052c-.003.07-.021.135-.021.206a4.773 4.773 0 0 0 4.772 4.772H71.01v7.068a2.693 2.693 0 0 0 2.693 2.693h4.571a2.692 2.692 0 0 0 2.693-2.693v-7.068h7.264a2.693 2.693 0 0 0 2.693-2.693v-4.571a2.693 2.693 0 0 0-2.693-2.693m-59.241.001V28.937h42.02v42.125H28.99z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crown(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M84.01 79.998a3.044 3.044 0 0 0-3.042-3.043H19.033a3.044 3.044 0 0 0-3.042 3.043v3.871a3.041 3.041 0 0 0 3.042 3.042h61.935a3.04 3.04 0 0 0 3.042-3.042zM24.619 38.433l-.017-.017l-.002.002a3.24 3.24 0 0 0-2.118-.798a3.264 3.264 0 0 0-3.26 3.193h-.08v29.238h.103c.155 1.659 1.536 2.965 3.237 2.965c.012 0 .024-.003.036-.004v.004h54.963v-.076c.013 0 .024.004.037.004c1.675 0 3.041-1.268 3.229-2.893h.037V41.42a3.267 3.267 0 0 0-3.194-3.945c-.776 0-1.479.282-2.04.733l-.009-.009l-.091.091a3.169 3.169 0 0 0-.312.312l-6.881 6.881l-15.852-15.852l-.031.031a3.248 3.248 0 0 0-2.373-1.033a3.25 3.25 0 0 0-2.377 1.037l-.035-.035l-15.887 15.885l-6.766-6.767a3.502 3.502 0 0 0-.317-.316m48.354 8.394v.095h-.095z"/><circle cx="22.348" cy="27.13" r="6.042"/><circle cx="77.348" cy="27.13" r="6.042"/><circle cx="50.348" cy="19.13" r="6.042"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CssThree(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m18.258 14.29l5.951 64.477L50 85.71l25.79-6.943l5.953-64.477zM68.96 35.187l-18.983 8.116l-.046.019H68.29l-2.1 24.146l-16.184 4.725l-.029-.009v.009l-16.272-4.812l-1.05-12.16h8.077l.525 6.299l8.646 2.183l.074-.021v.01l8.952-2.521l.613-10.148l-9.565-.029l-17.847-.06l-.612-7.611l18.459-7.688l1.076-.447H30.818l-.963-7.788h39.893z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Database(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M32.948 24.348h-.011c-.057-2.327-1.971-4.133-4.3-4.165v-.006h-15.69v.011c-2.307.057-4.067 1.852-4.119 4.161h-.01v51.2h.008c.043 2.3 1.838 4.089 4.121 4.178v.097h15.69c2.401-.001 4.311-1.875 4.311-4.274zM20.882 74.836a6.234 6.234 0 1 1 .001-12.467a6.234 6.234 0 0 1-.001 12.467m7.078-22.41H13.805V25.164H27.96zm34.105-28.078h-.011c-.057-2.327-1.971-4.133-4.3-4.165v-.006h-15.69v.011c-2.307.057-4.067 1.852-4.119 4.161h-.01v51.2h.008c.043 2.3 1.838 4.089 4.121 4.178v.097h15.69c2.401-.001 4.311-1.875 4.311-4.274zM49.999 74.836A6.234 6.234 0 1 1 50 62.369a6.234 6.234 0 0 1 0 12.467zm7.078-22.41H42.923V25.164h14.155v27.262zm34.105-28.078h-.011c-.057-2.327-1.971-4.133-4.3-4.165v-.006h-15.69v.011c-2.307.057-4.067 1.852-4.119 4.161h-.01v51.2h.008c.043 2.3 1.838 4.089 4.121 4.178v.097h15.69c2.401-.001 4.311-1.875 4.311-4.274zM79.117 74.836a6.234 6.234 0 1 1 .001-12.467a6.234 6.234 0 0 1-.001 12.467m7.078-22.41H72.04V25.164h14.155z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DieFive(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M83.921 23.292a7.44 7.44 0 0 0-7.213-7.264v-.011H23.292v.008a7.443 7.443 0 0 0-7.267 7.267h-.009v53.416h.009a7.44 7.44 0 0 0 7.267 7.267v.008h53.416v-.009a7.44 7.44 0 0 0 7.267-7.267h.009V23.292zM32.99 73.676a7.035 7.035 0 1 1 .002-14.068a7.035 7.035 0 0 1-.002 14.068m0-33.606a7.035 7.035 0 1 1 .002-14.068a7.035 7.035 0 0 1-.002 14.068M50 57.035a7.035 7.035 0 1 1 .001-14.069A7.035 7.035 0 0 1 50 57.035m17.009 16.944a7.033 7.033 0 0 1-7.034-7.033a7.034 7.034 0 1 1 14.069 0a7.034 7.034 0 0 1-7.035 7.033m0-33.909a7.033 7.033 0 0 1-7.034-7.033a7.034 7.034 0 1 1 14.069 0a7.034 7.034 0 0 1-7.035 7.033"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DieFour(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M83.921 23.292a7.44 7.44 0 0 0-7.213-7.264v-.011H23.292v.008a7.443 7.443 0 0 0-7.267 7.267h-.009v53.416h.009a7.44 7.44 0 0 0 7.267 7.267v.008h53.416v-.009a7.44 7.44 0 0 0 7.267-7.267h.009V23.292zM32.99 73.676a7.035 7.035 0 1 1 .001-14.069a7.035 7.035 0 0 1-.001 14.069m0-33.606a7.035 7.035 0 1 1 .001-14.069a7.035 7.035 0 0 1-.001 14.069m34.019 33.909a7.033 7.033 0 0 1-7.034-7.033a7.034 7.034 0 1 1 14.069 0a7.034 7.034 0 0 1-7.035 7.033m0-33.909a7.033 7.033 0 0 1-7.034-7.033a7.034 7.034 0 1 1 14.069 0a7.034 7.034 0 0 1-7.035 7.033"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DieOne(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M83.921 23.292a7.44 7.44 0 0 0-7.213-7.264v-.011H23.292v.008a7.443 7.443 0 0 0-7.267 7.267h-.009v53.416h.009a7.44 7.44 0 0 0 7.267 7.267v.008h53.416v-.009a7.44 7.44 0 0 0 7.267-7.267h.009V23.292zM50 57.621a7.621 7.621 0 1 1 0-15.24a7.621 7.621 0 0 1 0 15.24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DieSix(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M83.921 23.292a7.44 7.44 0 0 0-7.213-7.264v-.011H23.292v.008a7.443 7.443 0 0 0-7.267 7.267h-.009v53.416h.009a7.44 7.44 0 0 0 7.267 7.267v.008h53.416v-.009a7.44 7.44 0 0 0 7.267-7.267h.009V23.292zM32.99 73.676a7.035 7.035 0 1 1 .001-14.069a7.035 7.035 0 0 1-.001 14.069m0-16.641a7.035 7.035 0 1 1 .001-14.069a7.035 7.035 0 0 1-.001 14.069m0-16.965a7.035 7.035 0 1 1 .001-14.069a7.035 7.035 0 0 1-.001 14.069m34.019 33.909a7.033 7.033 0 0 1-7.034-7.033a7.034 7.034 0 1 1 14.069 0a7.034 7.034 0 0 1-7.035 7.033m.001-16.944a7.035 7.035 0 1 1 .001-14.069a7.035 7.035 0 0 1-.001 14.069m-.001-16.965a7.033 7.033 0 0 1-7.034-7.033a7.034 7.034 0 1 1 14.069 0a7.034 7.034 0 0 1-7.035 7.033"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DieThree(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M83.921 23.292a7.44 7.44 0 0 0-7.213-7.264v-.011H23.292v.008a7.443 7.443 0 0 0-7.267 7.267h-.009v53.416h.009a7.44 7.44 0 0 0 7.267 7.267v.008h53.416v-.009a7.44 7.44 0 0 0 7.267-7.267h.009V23.292zM32.99 40.07a7.035 7.035 0 1 1 .002-14.068a7.035 7.035 0 0 1-.002 14.068m17.009 16.965a7.033 7.033 0 0 1-7.034-7.033a7.034 7.034 0 1 1 14.069 0a7.034 7.034 0 0 1-7.035 7.033m17.01 16.944a7.033 7.033 0 0 1-7.034-7.033a7.034 7.034 0 1 1 14.069 0a7.034 7.034 0 0 1-7.035 7.033"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DieTwo(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M83.921 23.292a7.44 7.44 0 0 0-7.213-7.264v-.011H23.292v.008a7.443 7.443 0 0 0-7.267 7.267h-.009v53.416h.009a7.44 7.44 0 0 0 7.267 7.267v.008h53.416v-.009a7.44 7.44 0 0 0 7.267-7.267h.009V23.292zM32.99 40.07a7.035 7.035 0 1 1 .001-14.069a7.035 7.035 0 0 1-.001 14.069m34.019 33.908a7.033 7.033 0 0 1-7.034-7.033a7.034 7.034 0 1 1 14.069 0a7.034 7.034 0 0 1-7.035 7.033"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dislike(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M82.533 29.788h-.006a1.27 1.27 0 0 0-.369-.894l.003-.003l-9.764-9.764l-.024.024a1.246 1.246 0 0 0-.967-.272H35.304c-.073-.006-.143-.022-.218-.022a2.53 2.53 0 0 0-2.523 2.406h-.008v.082c0 .017-.005.032-.005.048s.005.032.005.048v36.498h.001l-.001.009a2.537 2.537 0 0 0 2.537 2.537c.047 0 .09-.011.136-.014h9.931c.15.244.334.453.536.637l.109.188v.109h.063l9.709 16.816l.025-.014c.99 1.753 2.847 2.936 4.98 2.936c3.171 0 5.742-2.607 5.742-5.825h.058V60.471h13.684v-.019a2.531 2.531 0 0 0 2.461-2.513h.008zm-55.126-3.915v-.009a2.536 2.536 0 0 0-2.537-2.536c-.047 0-.091.011-.137.014h-4.798v.018a2.531 2.531 0 0 0-2.461 2.513h-.008v27.502h.043a2.523 2.523 0 0 0 2.426 2.287v.022h4.971c1.4 0 2.536-1.136 2.536-2.537c0-.12-.019-.234-.035-.35z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dollar(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M54.284 44.798v-10.11c3.297.807 6.52 2.344 9.157 4.762l.011-.015c.197.157.436.265.708.265c.358 0 .665-.173.877-.428l.015.003l4.262-6.008l-.01-.005c.175-.202.291-.458.291-.746c0-.34-.153-.638-.387-.849c-3.953-3.651-9-5.843-14.924-6.502v-5.806h-.001c0-.637-.516-1.153-1.153-1.153h-4.578c-.637 0-1.153.516-1.153 1.153v5.659c-9.89 1.025-15.75 7.326-15.75 14.725c0 9.963 8.205 12.82 15.75 14.652v11.354c-4.845-.868-8.827-3.379-11.536-6.19c-.019-.021-.039-.039-.06-.058l-.052-.051l-.008.011a1.134 1.134 0 0 0-.719-.273a1.14 1.14 0 0 0-.998.608l-.014-.002l-4.125 6.124l.005.01a1.133 1.133 0 0 0-.292.748c0 .367.182.679.448.89l-.011.016c4.029 4.029 9.67 6.959 17.362 7.619v5.44c0 .637.516 1.153 1.153 1.153h4.578c.637 0 1.153-.517 1.153-1.153h.001V75.2c10.769-1.1 16.117-7.398 16.117-15.531c0-10.035-8.498-12.967-16.117-14.871m-6.886-1.686c-3.003-.951-5.055-2.051-5.055-4.176c0-2.49 1.832-4.248 5.055-4.688zm6.886 22.784V56.08c3.224 1.025 5.495 2.199 5.495 4.615c0 2.345-1.759 4.468-5.495 5.201"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DollarBill(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M93.681 30.128h-.045a3.864 3.864 0 0 0-3.865-3.865H10.185a3.866 3.866 0 0 0-3.866 3.865v39.745a3.865 3.865 0 0 0 3.866 3.865h79.586a3.863 3.863 0 0 0 3.865-3.864v-.001h.045zm-14.977 5.346v-3.431c0-.438.354-.793.793-.793c.031 0 .058.014.088.018v-.018h8.271v.006c.011 0 .02-.006.031-.006c.438 0 .793.355.793.793v8.417a.794.794 0 0 1-.793.793c-.011 0-.02-.006-.031-.006v.006h-3.37a.792.792 0 0 1-.793-.793v-4.194h-4.107v-.018c-.03.003-.057.018-.088.018a.791.791 0 0 1-.794-.792M21.296 64.526v3.431a.792.792 0 0 1-.793.793c-.031 0-.058-.014-.088-.018v.017h-8.271v-.006c-.011 0-.02.006-.031.006a.793.793 0 0 1-.793-.793v-8.418c0-.438.356-.793.793-.793c.011 0 .02.006.031.006v-.006h3.371c.439 0 .793.355.793.793v4.194h4.107v.018c.03-.003.057-.018.088-.018a.794.794 0 0 1 .793.794m.014-32.452v3.371a.793.793 0 0 1-.793.793h-4.194v4.107h-.018c.003.03.018.057.018.088a.793.793 0 0 1-.793.793H12.1a.792.792 0 0 1-.793-.793c0-.031.014-.058.018-.088h-.017v-8.271h.006c0-.011-.006-.02-.006-.031c0-.438.355-.793.793-.793h8.418c.438 0 .793.355.793.793c0 .011-.006.02-.006.031zM50 68.75c-10.356 0-18.75-8.394-18.75-18.749c0-10.354 8.394-18.751 18.75-18.751s18.75 8.396 18.75 18.751S60.356 68.75 50 68.75m38.693-9.094v8.271h-.006c0 .011.006.02.006.031a.793.793 0 0 1-.793.793h-8.418a.794.794 0 0 1-.793-.793c0-.011.006-.02.006-.031h-.006v-3.371c0-.438.355-.793.793-.793h4.194v-4.107h.018c-.003-.03-.018-.057-.018-.088c0-.438.355-.793.793-.793H87.9c.438 0 .793.354.793.793c0 .031-.014.058-.018.088z"/><path fill="currentColor" d="M50 36.22c-7.599 0-13.78 6.184-13.78 13.781c0 2.91.912 5.607 2.457 7.834v-.096c0-.617.322-1.169.812-1.398l5.64-2.627l2.212-1.03c-.999-.602-1.82-1.537-2.374-2.673a7.513 7.513 0 0 1-.761-3.31c0-.694.116-1.351.284-1.981c.729-2.749 2.906-4.76 5.509-4.76c2.653 0 4.863 2.085 5.55 4.915c.141.584.242 1.189.242 1.826a7.574 7.574 0 0 1-.679 3.136c-.536 1.182-1.355 2.163-2.364 2.799l2.31 1.102l5.462 2.605c.485.234.802.782.802 1.396v.095A13.697 13.697 0 0 0 63.779 50c.001-7.597-6.18-13.78-13.779-13.78"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Download(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M84.514 49.615H67.009a4.914 4.914 0 0 0-4.679 3.406a12.922 12.922 0 0 1-12.329 8.983a12.92 12.92 0 0 1-12.329-8.983a4.916 4.916 0 0 0-4.681-3.406H15.486a4.918 4.918 0 0 0-4.919 4.919v28.054a4.92 4.92 0 0 0 4.919 4.917h69.028a4.918 4.918 0 0 0 4.919-4.917V54.534c0-2.719-2.2-4.919-4.919-4.919"/><path fill="currentColor" d="M48.968 52.237c.247.346.651.553 1.076.553h.003c.428 0 .826-.207 1.076-.558l13.604-19.133a1.33 1.33 0 0 0 .096-1.374a1.319 1.319 0 0 0-1.177-.716h-6.399V13.821c0-.735-.593-1.326-1.323-1.326H44.078c-.732 0-1.323.591-1.323 1.326v17.188h-6.404a1.323 1.323 0 0 0-1.076 2.09z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eject(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.194 59.995l69.732-.074v-.014a2.493 2.493 0 0 0 2.361-2.489a2.487 2.487 0 0 0-.802-1.823L51.834 21.02l-.004.004a2.484 2.484 0 0 0-1.902-.892a2.494 2.494 0 0 0-2.02 1.041l-34.46 34.535a2.498 2.498 0 0 0 1.746 4.287m72.114 17.258l-.01-9.803v-.05h-.005a2.534 2.534 0 0 0-2.534-2.485v-.006l-69.751.074v.042a2.53 2.53 0 0 0-2.293 2.516c0 .033.008.063.01.096l.01 9.477c-.006.074-.022.145-.022.22a2.528 2.528 0 0 0 2.311 2.511v.023l69.751-.074a2.536 2.536 0 0 0 2.534-2.539z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Elevator(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M77.741 10.79H4.36a2.521 2.521 0 0 0-2.519 2.521V86.69a2.522 2.522 0 0 0 2.519 2.521h73.381a2.52 2.52 0 0 0 2.521-2.521V13.311a2.52 2.52 0 0 0-2.521-2.521M11.862 20.809h58.381V79.19H11.862zm85.662 43.157h-2.882v-7.737a.595.595 0 0 0-.595-.596h-5.333a.595.595 0 0 0-.595.596v7.736h-2.883a.598.598 0 0 0-.484.942l6.164 8.616a.593.593 0 0 0 .483.248h.001a.594.594 0 0 0 .484-.25l6.125-8.615a.595.595 0 0 0-.485-.94M85.276 36.035h2.882v7.737c0 .329.267.595.596.596l5.333-.001a.594.594 0 0 0 .595-.595v-7.737h2.883a.595.595 0 0 0 .484-.941l-6.164-8.615a.594.594 0 0 0-.483-.25h-.001a.6.6 0 0 0-.485.252l-6.124 8.614a.596.596 0 0 0 .484.94"/><circle cx="27.998" cy="31.993" r="4.955"/><path d="M37.38 38.28v-.024H19.628v.017a2.332 2.332 0 0 0-2.332 2.316h-.002v14.04h.023a2.083 2.083 0 1 0 4.166 0h.005v-9.802h.006a.662.662 0 0 1 .66-.638c.358 0 .646.283.661.638h.006v25.701a2.436 2.436 0 0 0 4.874-.003V57.584h.018a.675.675 0 1 1 1.35 0h.017v12.941a2.438 2.438 0 0 0 4.874 0V44.828a.666.666 0 1 1 1.331 0v9.802h.023a2.083 2.083 0 1 0 4.167 0h.004V40.59a2.33 2.33 0 0 0-2.099-2.31"/><circle cx="53.331" cy="31.993" r="4.955"/><path d="M62.713 38.28v-.024H44.961v.017a2.332 2.332 0 0 0-2.332 2.316h-.002v14.04h.023a2.083 2.083 0 1 0 4.166 0h.005v-9.802h.006a.662.662 0 0 1 .66-.638c.358 0 .646.283.661.638h.006v25.701a2.436 2.436 0 0 0 4.874-.003V57.584h.018a.675.675 0 1 1 1.35 0h.017v12.941a2.438 2.438 0 0 0 4.874 0V44.828a.666.666 0 1 1 1.331 0v9.802h.023a2.083 2.083 0 1 0 4.167 0h.004V40.59a2.331 2.331 0 0 0-2.099-2.31"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Euro(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M77.024 63.084c-.029-.016-.061-.016-.09-.03l.005-.012l-7.807-3.746l-.002.016a1.142 1.142 0 0 0-1.464.46c-.026.045-.028.094-.047.141c-2.274 4.302-7.07 7.714-12.505 7.714c-6.761 0-12.154-3.461-14.891-8.934h19.604c.025.002.047.015.073.015c.637 0 1.153-.516 1.153-1.153v-4.062c0-.622-.494-1.12-1.11-1.145v-.014h-21.65c-.081-.725-.081-1.529-.081-2.334c0-.885 0-1.689.161-2.576h21.513c.005 0 .009.003.015.003s.009-.003.015-.003h.029v-.006a1.148 1.148 0 0 0 1.11-1.145v-4.062c0-.637-.516-1.153-1.153-1.153c-.015 0-.028.008-.044.009H40.305c2.817-5.312 8.21-8.693 14.81-8.693c5.474 0 10.303 3.461 12.557 7.807l.011-.005c.303.5.923.677 1.448.442l.002.016l7.807-3.746l-.005-.012c.029-.014.062-.014.09-.03c.523-.302.705-.949.452-1.486l.015-.007l-.041-.069c-.002-.004-.002-.009-.005-.013c-.003-.005-.008-.007-.01-.012c-3.794-6.728-10.623-13.029-22.321-13.029c-13.12 0-23.744 7.324-27.366 18.836h-4.142c-.015 0-.028-.009-.044-.009c-.637 0-1.153.516-1.153 1.153v4.062c0 .637.516 1.153 1.153 1.153c.005 0 .009-.003.015-.003h2.964c-.081.887-.081 1.691-.081 2.576c0 .805 0 1.609.081 2.334h-2.978v.005c-.637 0-1.153.516-1.153 1.153v4.062c0 .637.516 1.153 1.153 1.153c.026 0 .047-.013.073-.015h4.032c3.541 11.673 14.246 19.08 27.446 19.08c11.679 0 18.505-6.44 22.303-13.078c.008-.012.021-.019.028-.032c.006-.011.006-.024.012-.035l.034-.056l-.013-.006a1.143 1.143 0 0 0-.455-1.479"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eye(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M84.448 47.672L67.712 30.936l-.116.116c-4.632-4.325-10.847-6.977-17.685-6.977c-7.679 0-14.575 3.341-19.322 8.646L15.013 48.298l.012.012a2.659 2.659 0 0 0-.673 1.76a2.66 2.66 0 0 0 1.009 2.079l16.896 16.895l.035-.035a25.828 25.828 0 0 0 17.62 6.915c7.67 0 14.559-3.333 19.305-8.627l.062.062L84.758 51.88a2.69 2.69 0 0 0 .185-.185l.001-.001c.431-.476.7-1.1.7-1.792a2.674 2.674 0 0 0-1.196-2.23M49.912 67.567c-9.702 0-17.567-7.865-17.567-17.567c0-9.701 7.865-17.566 17.567-17.566S67.48 40.299 67.48 50c0 9.702-7.866 17.567-17.568 17.567"/><circle cx="49.99" cy="50.077" r="8.512" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastForward(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M86.346 49.976c0-.702-.422-1.303-1.026-1.57l.017-.03l-37.39-21.588a1.715 1.715 0 0 0-2.92 1.172h-.005v15.249L16.58 26.788a1.715 1.715 0 0 0-2.92 1.172h-.005v44.031a1.718 1.718 0 0 0 2.948 1.197l28.42-16.408v15.212a1.718 1.718 0 0 0 2.948 1.197l37.545-21.677l-.031-.054c.512-.298.861-.847.861-1.482"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Female(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><circle cx="49.244" cy="20.203" r="8.2"/><path d="m72.147 54.101l-5.47-20.414c-.003-.016-.009-.031-.013-.047l-.096-.359l-.014.004c-.499-1.556-1.94-2.69-3.662-2.69c-.013 0-.026.004-.039.004v-.071H37.108v.067c-1.722 0-3.163 1.134-3.662 2.69l-.014-.004l-.096.358c-.004.017-.01.032-.013.049l-5.47 20.413l.052.014c-.025.169-.052.337-.052.513a3.479 3.479 0 0 0 6.721 1.273l.024.006l4.013-14.974v.057h.031c.11-.493.529-.869 1.054-.869c.526 0 .945.377 1.055.869h.033v.161c.002.025.014.046.014.071s-.013.046-.014.071v.153L34.36 65.424a1.833 1.833 0 0 0-.184.688l-.014.051l.008.002c-.001.027-.008.053-.008.081a1.87 1.87 0 0 0 1.871 1.87h4.751v15.849a4.032 4.032 0 0 0 8.064 0V68.116h2.293v15.849a4.033 4.033 0 0 0 8.065 0V68.116h4.765a1.87 1.87 0 0 0 1.871-1.87c0-.028-.007-.054-.008-.081l.008-.002l-.014-.051a1.871 1.871 0 0 0-.184-.688L59.24 41.52l.021-.006c-.027-.095-.059-.188-.059-.291c0-.608.493-1.102 1.102-1.102c.518 0 .932.365 1.05.847l.042-.011l4.006 14.951l.024-.006a3.479 3.479 0 0 0 6.721-1.273c0-.176-.027-.344-.052-.513z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FemaleSymbol(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M68.507 55.658c10.221-10.221 10.221-26.793 0-37.016c-10.222-10.221-26.793-10.221-37.015 0c-10.222 10.223-10.222 26.795 0 37.016A26.054 26.054 0 0 0 44.9 62.823v8.09h-5.299a1.543 1.543 0 0 0-1.544 1.545v6.886c0 .852.692 1.543 1.544 1.544H44.9v6.592c0 .753.54 1.379 1.254 1.515v.029h7.177c.852 0 1.544-.691 1.544-1.544h.001v-6.592h5.523c.853 0 1.544-.691 1.544-1.544v-6.886a1.54 1.54 0 0 0-1.544-1.545h-5.522v-8.048a26.047 26.047 0 0 0 13.63-7.207M50 53.349c-8.946 0-16.199-7.252-16.198-16.198c0-8.946 7.252-16.199 16.199-16.199c8.946 0 16.199 7.252 16.198 16.198c-.001 8.946-7.253 16.199-16.199 16.199"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Filter(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M90 22.292a6.496 6.496 0 0 0-6.258-6.488v-.012H16.848v.018c-.116-.006-.231-.018-.348-.018a6.5 6.5 0 0 0-4.077 11.559l30.14 30.139l-.001 18.599v.154h.015a3.392 3.392 0 0 0 1.713 2.8l-.009.016l7.872 4.545c.066.046.139.079.208.12l.028.016v-.001c.502.29 1.078.469 1.7.469a3.415 3.415 0 0 0 3.416-3.416c0-.09-.02-.175-.026-.263h.026V57.518l30.417-30.416l-.03-.03A6.472 6.472 0 0 0 90 22.292m-57.751 6.5h.014l.001.015z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FirstAid(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M29.365 27.367h6.146v-5.15h28.776v5.258h6.145V18.43h-.002a2.487 2.487 0 0 0-2.477-2.466H31.839a2.488 2.488 0 0 0-2.487 2.488c0 .045.011.086.013.131zm-4.593 56.67h50.436V32.838H24.772zm9.985-29.186a.79.79 0 0 1 .768-.768l10.235.001V43.849c0-.424.344-.768.768-.768h7.17c.424 0 .768.344.768.769v10.234h10.236c.424 0 .768.344.768.768v7.17a.766.766 0 0 1-.768.768H54.466v10.236a.767.767 0 0 1-.768.768h-7.169a.767.767 0 0 1-.768-.768V62.789H35.525a.767.767 0 0 1-.768-.768zM13.835 32.839a3.996 3.996 0 0 0-3.998 3.996c0 .053.013.102.016.154v42.94c-.001.035-.01.067-.01.102a4.001 4.001 0 0 0 3.993 3.998v.008h5.959V32.838zM90.16 80.006V36.814h-.019a3.999 3.999 0 0 0-3.995-3.977h-5.959v51.199h5.943v-.003c.012 0 .023.003.035.003a3.997 3.997 0 0 0 3.998-3.998c0-.01-.003-.021-.003-.032"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flag(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M84.772 23.937c-.21 0-.41.043-.593.118v-.009a24.274 24.274 0 0 1-11.253 2.764c-7.52 0-14.244-3.401-18.724-8.747c-4.048-3.82-9.5-6.167-15.504-6.167a22.519 22.519 0 0 0-15.052 5.745c-.842-1.774-2.635-3.009-4.729-3.009a5.248 5.248 0 0 0-5.248 5.247v62.974a5.248 5.248 0 0 0 10.496 0V54.497a24.288 24.288 0 0 1 10.984-2.614c7.53 0 14.261 3.41 18.741 8.767a22.517 22.517 0 0 0 15.487 6.146c6.49 0 12.328-2.747 16.451-7.128a1.55 1.55 0 0 0 .504-1.142v-33.03a1.56 1.56 0 0 0-1.56-1.559"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folder(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M86.016 24.09H49.314c-1.979-2.369-4.013-4.865-4.541-5.623a3.18 3.18 0 0 0-2.875-1.808H27.367c-.964 0-1.875.436-2.479 1.185l-5.021 6.247h-5.884a5.492 5.492 0 0 0-5.484 5.484V77.86a5.49 5.49 0 0 0 5.484 5.482h72.032a5.49 5.49 0 0 0 5.484-5.482V29.574a5.491 5.491 0 0 0-5.483-5.484"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderAdd(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M86.016 24.091H49.313c-1.979-2.37-4.013-4.865-4.541-5.624a3.18 3.18 0 0 0-2.875-1.808h-14.53c-.964 0-1.875.435-2.479 1.185l-5.021 6.247h-5.884a5.492 5.492 0 0 0-5.484 5.484v48.284a5.49 5.49 0 0 0 5.484 5.482h72.032a5.49 5.49 0 0 0 5.484-5.482V29.575a5.49 5.49 0 0 0-5.483-5.484m-4.473 39.195a.649.649 0 0 1-.195.465a.651.651 0 0 1-.464.194h-8.779v8.78c0 .174-.07.341-.193.466a.666.666 0 0 1-.465.193h-6.15a.66.66 0 0 1-.658-.659v-8.78h-8.78a.656.656 0 0 1-.465-.194a.654.654 0 0 1-.193-.465v-6.149a.678.678 0 0 1 .658-.659h8.779v-8.78c0-.363.294-.658.659-.658l6.15.001c.364 0 .659.293.659.658v8.778h8.78c.364 0 .659.294.659.659z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderLock(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M69.499 52.139a3.338 3.338 0 0 0-3.336 3.332v2.921h6.672v-2.941h-.001l.001-.011a3.339 3.339 0 0 0-3.336-3.301"/><path fill="currentColor" d="M86.016 24.091H49.313c-1.979-2.37-4.013-4.865-4.54-5.624a3.182 3.182 0 0 0-2.874-1.808H27.367c-.963 0-1.875.435-2.478 1.185l-5.021 6.247h-5.885a5.492 5.492 0 0 0-5.484 5.484v48.284a5.49 5.49 0 0 0 5.484 5.482h72.032a5.49 5.49 0 0 0 5.484-5.482V29.575a5.49 5.49 0 0 0-5.483-5.484m-4.473 48.397a.896.896 0 0 1-.898.897H58.67a.896.896 0 0 1-.897-.897V59.289c0-.496.401-.897.897-.897h2.713v-2.921c0-4.473 3.64-8.114 8.115-8.114c4.476 0 8.116 3.641 8.116 8.114l-.001.037v2.884h3.031c.496 0 .898.401.898.897v13.199z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Foot(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="28.549" cy="35.885" r="4.342" fill="currentColor"/><circle cx="65.794" cy="15.531" r="8.767" fill="currentColor"/><circle cx="47.712" cy="18.329" r="5.469" fill="currentColor"/><path fill="currentColor" d="M72.903 74.043a14.561 14.561 0 0 1-.51-14.513l-.034-.051a20.402 20.402 0 0 0 3.436-11.356c0-11.326-9.182-20.507-20.507-20.507c-11.326 0-20.507 9.181-20.507 20.507c0 2.212.36 4.338 1.009 6.335a17.804 17.804 0 0 0 2.068 5.011l-.096.055l14.854 25.729c1.568 4.64 5.947 7.984 11.115 7.984c6.484 0 11.743-5.256 11.743-11.741c0-2.813-.992-5.394-2.643-7.416z"/><circle cx="36.161" cy="26.693" r="4.342" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Foundation(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m78.682 76.043l-2.949-1.574a.44.44 0 0 0-.598.181l-1.189 2.215a.452.452 0 0 0-.034.338a.464.464 0 0 0 .215.263l2.951 1.571a.442.442 0 0 0 .598-.181l1.189-2.214a.44.44 0 0 0-.183-.599m.078-52.556l-2.983-1.747a.44.44 0 0 0-.604.155l-1.277 2.167a.44.44 0 0 0 .157.605l2.983 1.745a.44.44 0 0 0 .224.06a.44.44 0 0 0 .381-.216l1.276-2.166a.436.436 0 0 0-.157-.603m-5.975 49.41l-3.344-1.809V21.964l1.628.957c.069.039.147.06.223.06c.151 0 .296-.08.381-.219l1.276-2.166a.434.434 0 0 0 .049-.335a.452.452 0 0 0-.206-.272l-1.216-.712l-1.373-.945a1.581 1.581 0 0 0-1.819.228l-12.062 6.551l-12.194-6.621c-.001-.002-.005 0-.005 0c-.029-.016-.055-.036-.087-.051c-.136-.062-.288-.059-.434-.089c-.136-.027-.261-.081-.401-.081c-.141 0-.264.054-.401.082c-.143.029-.295.025-.426.088c-.026.012-.046.031-.07.042c-.007.007-.017.003-.023.009l-12.193 6.621l-12.194-6.621a1.971 1.971 0 0 0-1.755-.051c-.552.262-.893.757-.893 1.294v53.523c0 .504.3.975.797 1.245l13.118 7.126c.003.002.009.002.014.005c.03.015.054.03.083.045c.254.12.541.181.829.181c.287 0 .575-.06.831-.181c.023-.011.043-.03.067-.043c.008-.005.019-.002.027-.007l12.191-6.625l12.191 6.625c.004.002.011.002.017.005c.026.015.05.03.079.045c.256.12.544.181.831.181c.289 0 .577-.06.831-.181c.026-.011.042-.03.068-.043c.008-.005.018-.002.028-.007l12.13-6.591l.949.404l.85.452a.426.426 0 0 0 .208.051a.44.44 0 0 0 .391-.23l1.189-2.216a.446.446 0 0 0-.18-.6m-42.74 5.694l-11.351-6.146V22.42l11.351 6.146zm26.236-.072l-13.078-7.086V21.439l13.078 7.119zm28.031-22.202h-2.521a.442.442 0 0 0-.44.442v5.255c0 .243.198.44.44.44h2.521a.441.441 0 0 0 .442-.44v-5.255a.443.443 0 0 0-.442-.442m0 21.021h-2.521a.438.438 0 0 0-.225.064l-.37.221a.463.463 0 0 0-.162.17l-1.189 2.217a.447.447 0 0 0 .183.6l1.556.828c.416.223.883.346 1.313.346c.898 0 1.855-.538 1.855-2.051V77.78a.441.441 0 0 0-.44-.442m-.848-51.096l-1.722-1.011a.442.442 0 0 0-.605.159l-1.276 2.164a.43.43 0 0 0-.047.334a.427.427 0 0 0 .204.271l1.333.778v1.541c0 .243.198.441.44.441h2.521a.442.442 0 0 0 .442-.441v-1.877c-.001-1.57-1.235-2.326-1.29-2.359m.848 9.051h-2.521a.443.443 0 0 0-.44.443v5.256c0 .244.198.442.44.442h2.521a.442.442 0 0 0 .442-.442v-5.256a.444.444 0 0 0-.442-.443m0 31.537h-2.521a.44.44 0 0 0-.44.44v5.257c0 .244.198.442.44.442h2.521a.442.442 0 0 0 .442-.442V67.27a.442.442 0 0 0-.442-.44m0-21.025h-2.521a.442.442 0 0 0-.44.442v5.255c0 .243.198.442.44.442h2.521a.443.443 0 0 0 .442-.442v-5.255a.443.443 0 0 0-.442-.442"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GraphBar(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M46.05 60.163H31.923c-.836 0-1.513.677-1.513 1.513V83.61c0 .836.677 1.513 1.513 1.513H46.05c.836 0 1.512-.677 1.512-1.513V61.675c0-.836-.677-1.512-1.512-1.512m22.027-45.285H53.95c-.836 0-1.513.677-1.513 1.513v67.218c0 .836.677 1.513 1.513 1.513h14.127c.836 0 1.513-.677 1.513-1.513V16.391c0-.836-.677-1.513-1.513-1.513m22.14 20.421H76.09c-.836 0-1.513.677-1.513 1.513v46.797c0 .836.677 1.513 1.513 1.513h14.126c.836 0 1.513-.677 1.513-1.513V36.812c0-.835-.677-1.513-1.512-1.513m-66.307 0H9.783c-.836 0-1.513.677-1.513 1.513v46.797c0 .836.677 1.513 1.513 1.513H23.91c.836 0 1.513-.677 1.513-1.513V36.812c0-.835-.677-1.513-1.513-1.513"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GraphHorizontal(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M86.304 40.985H13.696c-.836 0-1.513.677-1.513 1.513v15.127c0 .836.677 1.513 1.513 1.513h72.608c.836 0 1.513-.677 1.513-1.513V42.497c0-.835-.677-1.512-1.513-1.512M56.261 17.848v-.053H13.696c-.836 0-1.513.677-1.513 1.513v15.127c0 .836.677 1.513 1.513 1.513h42.565v-.053a1.509 1.509 0 0 0 1.135-1.459V19.307c0-.704-.483-1.29-1.135-1.459m9.622 46.205H13.696c-.836 0-1.513.677-1.513 1.513v15.127c0 .836.677 1.513 1.513 1.513h52.187c.836 0 1.513-.677 1.513-1.513V65.566c0-.836-.677-1.513-1.513-1.513"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GraphPie(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M54.193 12.87c-.731 0-1.325.594-1.325 1.326c0 .068.029.127.039.193h-.039v31.438c0 .732.594 1.326 1.325 1.326H85.63v-.039c.066.01.125.039.194.039c.731 0 1.325-.594 1.325-1.326c-.188-18.12-14.836-32.768-32.956-32.957"/><path fill="currentColor" d="M79.485 53.46c0-.732-.593-1.326-1.325-1.326H49.261a1.326 1.326 0 0 1-1.325-1.326V22.015h-.039c.01-.066.039-.125.039-.193c0-.733-.594-1.326-1.326-1.326c-.032 0-.058.016-.089.018v-.009c-.118-.001-.235-.009-.353-.009c-18.4 0-33.317 14.917-33.317 33.317S27.768 87.13 46.168 87.13s33.317-14.917 33.317-33.317c0-.106-.005-.211-.007-.318c0-.013.007-.023.007-.035"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GraphTrend(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M85.088 80.271c2.715 0 4.912-2.2 4.912-4.917V24.645a4.912 4.912 0 0 0-4.912-4.917H14.912A4.913 4.913 0 0 0 10 24.645v50.709a4.912 4.912 0 0 0 4.912 4.917zm-4.913-9.828h-60.35V29.556h60.35z"/><path fill="currentColor" d="M41.564 65.897c-.927 0-1.819-.37-2.481-1.029l-4.896-4.897l-5.96 4.932a3.507 3.507 0 0 1-4.942-.466a3.511 3.511 0 0 1 .466-4.941l8.424-6.969a3.505 3.505 0 0 1 4.719.223l4.072 4.075l8.979-13.989a3.473 3.473 0 0 1 3.181-1.602c1.272.082 2.4.846 2.945 2.002l5.299 11.206L71.602 36.45a3.518 3.518 0 0 1 4.787-1.318a3.517 3.517 0 0 1 1.315 4.79L64.133 63.784a3.537 3.537 0 0 1-3.183 1.769a3.498 3.498 0 0 1-3.041-2.008L52.43 51.954L44.516 64.28a3.512 3.512 0 0 1-2.952 1.617"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GuideDog(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M39.818 24.963L51.575 36.72c.033.035.067.069.102.102l.095.095l.009-.009c.35.285.791.464 1.279.464a2.035 2.035 0 0 0 2.034-2.035c0-.733-.391-1.37-.973-1.728l-11.7-11.699a2.02 2.02 0 0 0-1.523-.698a2.034 2.034 0 0 0-2.034 2.034c0 .725.382 1.358.954 1.717m13.241 40.766V40.69h-26.9c-.147 0-.29.037-.432.091L14.76 34.45l-.009.015a2.545 2.545 0 0 0-1.175-.296a2.568 2.568 0 0 0-1.388 4.734l-.019.033l11.112 6.416c-.429 1.868-.685 4.162-.685 6.651c0 5.261 1.133 9.669 2.664 10.936L18.996 73.79a3.474 3.474 0 0 0-.751 2.146a3.512 3.512 0 0 0 3.512 3.512a3.492 3.492 0 0 0 2.774-1.382l.06.035l7.327-12.691l7.327 12.69l.06-.034a3.492 3.492 0 0 0 2.774 1.382a3.512 3.512 0 0 0 3.512-3.512c0-.812-.287-1.551-.751-2.146l-4.732-8.195c2.467.142 5.146.22 7.951.22c1.721 0 3.393-.03 5-.086"/><path fill="currentColor" d="M70.685 63.13c1.151-.649 2.112-3.049 2.597-6.305L57.059 40.602v22.686L50.995 73.79a3.474 3.474 0 0 0-.751 2.146a3.512 3.512 0 0 0 3.512 3.512a3.492 3.492 0 0 0 2.774-1.382l.06.035l7.327-12.691l7.328 12.69l.06-.034a3.492 3.492 0 0 0 2.774 1.382a3.512 3.512 0 0 0 3.512-3.512c0-.812-.287-1.551-.751-2.146zm18.312-21.454c0-.83-.409-1.56-1.031-2.015l-3.906-3.906V31.69c0-1.528-.879-2.837-2.151-3.49l-5.054-2.918l-.014.024a3.892 3.892 0 0 0-1.406-.503l-2.454-4.25l-5.098 8.83l-8.193 8.192l13.919 13.919a34.687 34.687 0 0 0-.17-3.058l2.746-2.746h9.875v-.027l2.174-2.174l-.01-.01a2.494 2.494 0 0 0 .773-1.803"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HearingAid(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M47.484 28.548c-11.263 0-20.424 9.236-20.424 20.586c0 1.848 1.489 3.349 3.319 3.349c1.833 0 3.323-1.501 3.323-3.349c0-7.654 6.182-13.887 13.782-13.887c7.598 0 13.777 6.232 13.777 13.887c0 1.615-.271 3.192-.729 4.521c-.872 1.553-2.828 4.861-4.005 6.305c-.231.281-.462.549-.694.811c-1.647 1.898-3.704 4.261-3.559 9.344c.092 3.171.146 6.25-.061 7.299l-.062.355c-.151.877-.338 1.963-4.407 2.105a3.305 3.305 0 0 0-2.315 1.064a3.32 3.32 0 0 0-.891 2.398a3.321 3.321 0 0 0 3.318 3.23h.118c6.288-.217 9.894-2.797 10.757-7.844c.294-1.482.339-3.547.189-8.807c-.07-2.453.55-3.166 1.912-4.732c.273-.312.551-.627.824-.965c1.946-2.379 4.73-7.389 4.754-7.432c.081-.148.155-.305.223-.475a20.666 20.666 0 0 0 1.276-7.18c-.002-11.347-9.164-20.583-20.425-20.583"/><path fill="currentColor" d="M52.169 49.685c0 1.537 1.245 2.798 2.775 2.798c1.53 0 2.776-1.261 2.776-2.798c0-5.689-4.594-10.319-10.237-10.319c-5.648 0-10.24 4.63-10.24 10.319c0 1.537 1.244 2.798 2.773 2.798c1.531 0 2.776-1.261 2.776-2.798c0-2.604 2.103-4.721 4.69-4.721c2.584 0 4.687 2.117 4.687 4.721M33.252 59.843L19.534 73.665a3.08 3.08 0 0 0-.902 2.193c0 .828.319 1.604.903 2.193a3.054 3.054 0 0 0 4.352-.002l13.718-13.824c.582-.586.9-1.363.9-2.192a3.07 3.07 0 0 0-.899-2.188c-1.167-1.182-3.195-1.18-4.354-.002m25.859-37.547c-.017 0-.032.004-.048.005v-.005h-1.294v.029a1.275 1.275 0 0 0-1.139 1.232h-.006v2.301h.031a1.263 1.263 0 0 0 1.113 1.107v.029h1.294v-.021c.016 0 .032.005.048.005c4.775 0 8.656 3.862 8.705 8.617h-.011v1.295h.03a1.275 1.275 0 0 0 1.235 1.14v.005h2.299v-.031a1.263 1.263 0 0 0 1.108-1.114h.03v-1.295h-.005c-.048-7.338-6.035-13.299-13.39-13.299"/><path fill="currentColor" d="M81.365 35.597c-.049-12.23-10.012-22.164-22.254-22.164c-.017 0-.032.004-.048.005v-.005h-1.294v.029a1.273 1.273 0 0 0-1.14 1.237h-.005v2.298h.03a1.263 1.263 0 0 0 1.114 1.107v.029h1.294v-.02c.016 0 .032.005.048.005c9.661 0 17.519 7.832 17.568 17.478h-.01v1.295h.03a1.275 1.275 0 0 0 1.232 1.139v.006h2.302v-.031a1.263 1.263 0 0 0 1.105-1.113h.03v-1.295z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heart(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M67.607 13.462A22.1 22.1 0 0 0 50 22.136a22.098 22.098 0 0 0-17.61-8.674c-12.266 0-22.283 10.013-22.33 22.32c-.046 13.245 6.359 21.054 11.507 27.331l1.104 1.349c6.095 7.515 24.992 21.013 25.792 21.584a2.654 2.654 0 0 0 3.077 0c.8-.571 19.697-14.069 25.792-21.584l1.103-1.349c5.147-6.277 11.553-14.086 11.507-27.331c-.048-12.307-10.066-22.32-22.335-22.32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Home(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M83.505 37.85L51.013 12.688a2.508 2.508 0 0 0-3.1.025L16.46 37.874a2.509 2.509 0 0 0-.939 1.956v45.5a2.504 2.504 0 0 0 2.505 2.505h18.697a2.506 2.506 0 0 0 2.505-2.505V57.471h21.54V85.33a2.505 2.505 0 0 0 2.505 2.505h18.7a2.506 2.506 0 0 0 2.505-2.505v-45.5a2.498 2.498 0 0 0-.973-1.98"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HtmlFive(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.706 16.778l5.332 59.802l23.926 6.643l23.993-6.651l5.337-59.793H20.706zM68.172 30.97l-.334 3.718l-.147 1.648H39.624l.67 7.511h26.73l-.179 1.97l-1.724 19.311l-.11 1.238L50 70.526l-.034.011l-15.024-4.171l-1.027-11.517h7.362l.522 5.851l8.168 2.205l.007-.002v-.001l8.181-2.208l.851-9.512h-25.42L31.784 30.97l-.175-1.969h36.739z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IndentLess(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M81.232 15.138H18.769a3.407 3.407 0 0 0-3.407 3.407v3.143a3.407 3.407 0 0 0 3.407 3.407h62.463a3.407 3.407 0 0 0 3.407-3.407v-3.143a3.407 3.407 0 0 0-3.407-3.407m0 59.767H18.769a3.407 3.407 0 0 0-3.407 3.407v3.143a3.407 3.407 0 0 0 3.407 3.407h62.463a3.407 3.407 0 0 0 3.407-3.407v-3.143a3.408 3.408 0 0 0-3.407-3.407m3.407-16.515a3.407 3.407 0 0 0-3.407-3.407h-38.82a3.407 3.407 0 0 0-3.407 3.407v3.143a3.406 3.406 0 0 0 3.407 3.406h38.82a3.407 3.407 0 0 0 3.407-3.406zm0-19.813a3.407 3.407 0 0 0-3.407-3.407h-38.82a3.407 3.407 0 0 0-3.407 3.407v3.143a3.406 3.406 0 0 0 3.407 3.406h38.82a3.407 3.407 0 0 0 3.407-3.406zM15.361 50.008c0 .232.14.431.34.519l-.006.011l12.376 7.146a.567.567 0 0 0 .963-.389h.003V42.722a.57.57 0 0 0-.975-.397L15.636 49.5l.014.024a.56.56 0 0 0-.289.484"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IndentMore(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M81.232 15.138H18.769a3.407 3.407 0 0 0-3.407 3.407v3.143a3.407 3.407 0 0 0 3.407 3.407h62.463a3.407 3.407 0 0 0 3.407-3.407v-3.143a3.407 3.407 0 0 0-3.407-3.407m0 59.767H18.769a3.407 3.407 0 0 0-3.407 3.407v3.143a3.407 3.407 0 0 0 3.407 3.407h62.463a3.407 3.407 0 0 0 3.407-3.407v-3.143a3.408 3.408 0 0 0-3.407-3.407m3.407-16.515a3.407 3.407 0 0 0-3.407-3.407h-38.82a3.407 3.407 0 0 0-3.407 3.407v3.143a3.406 3.406 0 0 0 3.407 3.406h38.82a3.407 3.407 0 0 0 3.407-3.406zm0-19.813a3.407 3.407 0 0 0-3.407-3.407h-38.82a3.407 3.407 0 0 0-3.407 3.407v3.143a3.406 3.406 0 0 0 3.407 3.406h38.82a3.407 3.407 0 0 0 3.407-3.406zM15.93 57.847a.567.567 0 0 0 .406-.172L28.763 50.5l-.014-.024a.56.56 0 0 0 .289-.484a.567.567 0 0 0-.34-.519l.006-.01l-12.376-7.146a.567.567 0 0 0-.963.389h-.003v14.573a.567.567 0 0 0 .568.568"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Info(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M50 12.5c-20.712 0-37.5 16.793-37.5 37.502C12.5 70.712 29.288 87.5 50 87.5s37.5-16.788 37.5-37.498C87.5 29.293 70.712 12.5 50 12.5m3.826 58.36c0 .72-.584 1.304-1.304 1.304h-5.044c-.72 0-1.304-.583-1.304-1.304V46.642c0-.72.584-1.304 1.304-1.304h5.044c.72 0 1.304.583 1.304 1.304zm-3.857-30.927c-2.47 0-4.518-2.048-4.518-4.579a4.512 4.512 0 0 1 4.518-4.518c2.531 0 4.579 1.987 4.579 4.518a4.576 4.576 0 0 1-4.579 4.579"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Italic(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M60.571 24.301a2.604 2.604 0 0 0-2.604-2.604h-4.594a2.598 2.598 0 0 0-2.59 2.463l-.014-.001l-11.276 50.978l-.015.066l-.011.048h.006a2.55 2.55 0 0 0-.045.449a2.595 2.595 0 0 0 2.406 2.584v.02h4.792a2.595 2.595 0 0 0 2.577-2.336l.013.001l11.257-50.972l-.008-.001a2.58 2.58 0 0 0 .106-.695"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Key(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M93.481 44.925a1.88 1.88 0 0 0-1.882-1.709H52.128c-2.601-9.968-11.654-17.354-22.426-17.354C16.907 25.862 6.5 36.269 6.5 49.064c0 12.798 10.407 23.202 23.202 23.202c11.445 0 20.953-8.332 22.84-19.242h8.79v12.517h.014a1.893 1.893 0 0 0 1.886 1.758c.026 0 .049-.006.075-.007v.011h6.1V67.3a1.891 1.891 0 0 0 1.886-1.758h.016V53.025h4.979v19.353h.015a1.898 1.898 0 0 0 1.887 1.76l.021-.002h6.112l.021.002a1.899 1.899 0 0 0 1.887-1.76h.015V53.024h5.355v-.002a1.903 1.903 0 0 0 1.9-1.901H93.5v-6.197h-.019zM29.702 62.532c-7.338 0-13.308-5.974-13.308-13.311c0-7.338 5.97-13.311 13.308-13.311c7.337 0 13.304 5.973 13.304 13.311c.001 7.338-5.967 13.311-13.304 13.311"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Laptop(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.022 67.699h71.957a4.978 4.978 0 0 0 4.976-4.977V18.628a4.976 4.976 0 0 0-4.976-4.976H14.022a4.976 4.976 0 0 0-4.976 4.976v44.094a4.976 4.976 0 0 0 4.976 4.977m4.976-44.095h62.005v34.142H18.998zm76.327 49.082H4.675a2.294 2.294 0 0 0-2.294 2.294v5.555c0 .58.218 1.135.609 1.559l3.25 3.518a2.304 2.304 0 0 0 1.684.735h83.915c.612 0 1.198-.245 1.631-.681l3.485-3.519a2.3 2.3 0 0 0 .664-1.613V74.98a2.295 2.295 0 0 0-2.294-2.294m-36.656 7.969c0 .28-.185.507-.415.507H41.746c-.23 0-.415-.227-.415-.507v-2.264c0-.28.185-.504.415-.504h16.508c.23 0 .415.224.415.504z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Layout(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M27.725 30.515h44.55v23.113h-44.55zm1.068 28.1h11.074v10.87H28.793zm32.122 0h11.074v10.87H60.915zm-16.061 0h11.074v10.87H44.854z"/><path d="M82.69 15.588H17.31a4.513 4.513 0 0 0-4.512 4.512h-.001v60.067h.028c.14 2.366 2.085 4.246 4.485 4.246h65.38c2.401 0 4.344-1.88 4.484-4.246h.027V20.1a4.51 4.51 0 0 0-4.511-4.512m-5.428 58.885H22.738V25.527h54.525v48.946z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lightbulb(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m80.302 18.915l.042-.042l-.008-.008a1.447 1.447 0 0 0-.033-2.003l-2.053-2.053a1.447 1.447 0 0 0-2.003-.033l-.009-.009l-8.842 8.843a1.452 1.452 0 0 0 0 2.053l2.053 2.054h.001a1.452 1.452 0 0 0 2.053 0zm-31.754.359v.001h2.903v-.001c.782 0 1.415-.619 1.446-1.393h.006V5.378c0-.802-.65-1.452-1.452-1.452v-.002h-2.903v.001c-.802 0-1.452.65-1.452 1.452v12.504h.006a1.449 1.449 0 0 0 1.446 1.393m43.065 20.723c0-.783-.619-1.415-1.393-1.446v-.006H77.716c-.802 0-1.452.649-1.452 1.452V42.9c0 .803.65 1.453 1.452 1.453H90.22v-.006a1.45 1.45 0 0 0 1.393-1.447zM50 24.201c-11.756 0-21.285 9.529-21.285 21.285a21.16 21.16 0 0 0 3.968 12.346h-.008c3.174 4.952 5.157 11.512 5.303 18.739c.255.419.711.703 1.237.703c.056 0 .109-.01.163-.017v.017H60.55v-.017c.054.006.107.017.162.017a1.45 1.45 0 0 0 1.303-.822c.175-7.584 2.371-14.429 5.843-19.439c2.154-3.324 3.427-7.27 3.427-11.526c0-11.757-9.529-21.286-21.285-21.286m10.712 65.587c-.056 0-.108.01-.162.017v-.017H39.378v.017c-.054-.006-.107-.017-.163-.017c-.801 0-1.452.649-1.452 1.452v3.384c0 .802.65 1.452 1.452 1.452c.056 0 .109-.01.163-.017v.017H60.55v-.016c.054.006.107.016.162.016c.802 0 1.452-.65 1.452-1.452V91.24c0-.803-.65-1.452-1.452-1.452m0-9.44c-.056 0-.108.01-.162.017v-.017H39.378v.017c-.054-.006-.107-.017-.163-.017a1.45 1.45 0 0 0-1.447 1.403l-.005.057v3.376c0 .802.65 1.452 1.452 1.452c.056 0 .109-.01.163-.017v.017H60.55v-.016c.054.006.107.017.162.017c.802 0 1.452-.65 1.452-1.452V81.8c0-.803-.65-1.452-1.452-1.452M22.344 39.685v-.006H9.839c-.802 0-1.452.65-1.452 1.451v2.904c0 .801.65 1.451 1.452 1.451h12.505v-.005a1.45 1.45 0 0 0 1.393-1.445V41.13a1.45 1.45 0 0 0-1.393-1.445m-2.145-21.27l8.842 8.843l.009-.009c.569.523 1.45.519 2.002-.033h.001l2.053-2.054a1.446 1.446 0 0 0 .033-2.003l.009-.009l-8.843-8.842a1.453 1.453 0 0 0-2.053 0l-2.053 2.053a1.453 1.453 0 0 0 0 2.054"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Like(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M82.488 74.334c0-.12-.019-.234-.035-.35V47.06h-.001l.001-.009a2.536 2.536 0 0 0-2.537-2.536c-.047 0-.09.011-.136.014h-4.799v.018a2.531 2.531 0 0 0-2.461 2.513h-.008v27.502h.043a2.523 2.523 0 0 0 2.426 2.287v.022h4.971c1.4 0 2.536-1.136 2.536-2.537m-14.998 4.22V42.061l.001-.009a2.536 2.536 0 0 0-2.537-2.536c-.046 0-.09.011-.136.014h-9.932a2.96 2.96 0 0 0-.536-.637l-9.88-17.115l-.025.015c-.99-1.753-2.848-2.936-4.98-2.936c-3.171 0-5.742 2.608-5.742 5.825h-.058v14.847H19.981v.018a2.531 2.531 0 0 0-2.461 2.513h-.008v28.152h.006c0 .323.124.646.368.894l-.003.003l9.764 9.764l.024-.024c.281.231.626.318.967.272h36.103c.073.006.143.022.217.022a2.53 2.53 0 0 0 2.523-2.405h.008v-.078c0-.018.005-.035.005-.053s-.003-.035-.004-.053m-46.985-8.342l.042.047l-.047-.047z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Link(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M78.663 21.338c-7.552-7.552-19.648-7.79-27.486-.713l-.019-.019L41.06 30.703a19.887 19.887 0 0 0-4.187 6.176a19.83 19.83 0 0 0-5.419 3.468l-.019-.019l-10.097 10.097c-7.797 7.797-7.797 20.439 0 28.237c7.797 7.798 20.439 7.798 28.237 0l10.098-10.098l-.019-.019a19.825 19.825 0 0 0 3.467-5.419a19.887 19.887 0 0 0 6.176-4.187l10.098-10.098l-.019-.019c7.076-7.837 6.838-19.933-.713-27.484M42.761 71.487l-.001.001c-3.935 3.935-10.314 3.935-14.248 0c-3.935-3.935-3.935-10.314 0-14.248l.001-.001l7.367-7.367a19.847 19.847 0 0 0 5.18 9.068a19.85 19.85 0 0 0 9.067 5.181zm5.473-19.721a10.024 10.024 0 0 1-2.919-6.452a10.027 10.027 0 0 1 6.452 2.919a10.02 10.02 0 0 1 2.919 6.452a10.03 10.03 0 0 1-6.452-2.919m23.875-9.627l-.619.619l-.001.001h-.001l-7.369 7.369a19.845 19.845 0 0 0-5.179-9.068a19.85 19.85 0 0 0-9.069-5.18l7.369-7.369l.001-.001l.001-.001l.619-.619l.029.028c3.959-3.329 9.874-3.134 13.6.591s3.921 9.642.591 13.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func List(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.563 30.277h.012a2.268 2.268 0 0 0 2.246 2.267v.002H80.18v-.001a2.269 2.269 0 0 0 2.259-2.268h.01V19.818a2.269 2.269 0 0 0-2.269-2.265H19.821a2.27 2.27 0 0 0-2.269 2.269c0 .039.01.076.012.115zm62.616 12.227H19.821a2.27 2.27 0 0 0-2.269 2.269c0 .039.01.076.012.115v10.34h.012a2.268 2.268 0 0 0 2.246 2.267v.002h60.359v-.001a2.269 2.269 0 0 0 2.259-2.268h.01V44.769a2.272 2.272 0 0 0-2.271-2.265m0 24.95H19.821a2.27 2.27 0 0 0-2.269 2.269c0 .039.01.076.012.115v10.34h.012a2.268 2.268 0 0 0 2.246 2.267v.002h60.359v-.001a2.27 2.27 0 0 0 2.259-2.269h.01V69.718a2.272 2.272 0 0 0-2.271-2.264"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListBullet(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M88.721 20.13H26.258a3.407 3.407 0 0 0-3.407 3.407v3.143a3.407 3.407 0 0 0 3.407 3.407h62.463a3.407 3.407 0 0 0 3.407-3.407v-3.143a3.407 3.407 0 0 0-3.407-3.407m0 24.892H26.258a3.407 3.407 0 0 0-3.407 3.407v3.143a3.407 3.407 0 0 0 3.407 3.407h62.463a3.407 3.407 0 0 0 3.407-3.407v-3.143a3.407 3.407 0 0 0-3.407-3.407m0 24.891H26.258a3.407 3.407 0 0 0-3.407 3.407v3.143a3.407 3.407 0 0 0 3.407 3.407h62.463a3.407 3.407 0 0 0 3.407-3.407V73.32a3.408 3.408 0 0 0-3.407-3.407"/><circle cx="12.856" cy="25.108" r="4.984" fill="currentColor"/><circle cx="12.856" cy="49.002" r="4.984" fill="currentColor"/><circle cx="12.856" cy="74.891" r="4.984" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListNumber(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.427 22.495v7.747H16.7V19.566h-1.985l-3.361 3.377l1.296 1.361zm-1.599 23.928c.88 0 1.729.448 1.729 1.393c0 1.312-1.28 2.401-5.65 5.634v1.793h8.035v-2.001h-4.354c2.77-2.017 4.274-3.601 4.274-5.426c0-2.129-1.792-3.41-4.082-3.41c-1.489 0-3.073.544-4.114 1.745l1.296 1.505c.721-.753 1.649-1.233 2.866-1.233m1.776 28.313c1.184-.208 2.337-1.088 2.337-2.433c0-1.776-1.537-2.897-4.034-2.897c-1.873 0-3.217.72-4.082 1.697l1.136 1.424a3.852 3.852 0 0 1 2.721-1.104c1.104 0 1.985.416 1.985 1.264c0 .785-.8 1.137-1.985 1.137c-.4 0-1.137 0-1.329-.016v2.049c.16-.016.88-.032 1.329-.032c1.489 0 2.145.384 2.145 1.233c0 .8-.721 1.36-2.017 1.36c-1.041 0-2.209-.448-2.929-1.216l-1.2 1.521c.784.96 2.257 1.712 4.209 1.712c2.561 0 4.21-1.296 4.21-3.137c.002-1.602-1.407-2.45-2.496-2.562m73.322-54.847H25.462a3.407 3.407 0 0 0-3.407 3.407v3.143a3.407 3.407 0 0 0 3.407 3.407h62.463a3.407 3.407 0 0 0 3.407-3.407v-3.143a3.406 3.406 0 0 0-3.406-3.407m0 24.892H25.462a3.407 3.407 0 0 0-3.407 3.407v3.143a3.407 3.407 0 0 0 3.407 3.407h62.463a3.407 3.407 0 0 0 3.407-3.407v-3.143a3.406 3.406 0 0 0-3.406-3.407m0 24.892H25.462a3.407 3.407 0 0 0-3.407 3.407v3.143a3.407 3.407 0 0 0 3.407 3.407h62.463a3.407 3.407 0 0 0 3.407-3.407V73.08a3.406 3.406 0 0 0-3.406-3.407"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListThumbnails(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m80.182 67.457l-35.4-.001v.001l-.009-.001a2.27 2.27 0 0 0-2.269 2.269c0 .042.01.08.012.121V80.18h.011a2.268 2.268 0 0 0 2.254 2.268v.001h35.4v-.002a2.268 2.268 0 0 0 2.254-2.268V69.825c.002-.034.01-.067.01-.101a2.267 2.267 0 0 0-2.263-2.267m-49.904-.003H19.822a2.27 2.27 0 0 0-2.269 2.269v10.454a2.269 2.269 0 0 0 2.268 2.269h10.455a2.27 2.27 0 0 0 2.269-2.269V69.723a2.267 2.267 0 0 0-2.267-2.269m49.904-24.987l-35.4-.001v.001l-.009-.001a2.27 2.27 0 0 0-2.269 2.269c0 .042.01.08.012.121V55.19h.011a2.268 2.268 0 0 0 2.254 2.268v.001h35.4v-.002a2.268 2.268 0 0 0 2.254-2.268V44.835c.002-.034.01-.067.01-.101a2.267 2.267 0 0 0-2.263-2.267m-49.904-.003H19.822a2.27 2.27 0 0 0-2.269 2.269v10.454a2.269 2.269 0 0 0 2.268 2.269h10.455a2.27 2.27 0 0 0 2.269-2.269V44.733a2.267 2.267 0 0 0-2.267-2.269m12.238-12.188h.011a2.268 2.268 0 0 0 2.254 2.268v.001h35.4v-.002a2.268 2.268 0 0 0 2.254-2.268V19.921c.002-.034.01-.067.01-.101a2.269 2.269 0 0 0-2.264-2.268v-.001h-35.4v.001l-.009-.001a2.27 2.27 0 0 0-2.269 2.269c0 .042.01.08.012.121zM30.278 17.551H19.822a2.27 2.27 0 0 0-2.269 2.269v10.454a2.269 2.269 0 0 0 2.268 2.269h10.455a2.27 2.27 0 0 0 2.269-2.269V19.82a2.266 2.266 0 0 0-2.267-2.269"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lock(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M82.105 44.218h-8.858v-8.431c.003-.036.003-.071.003-.102c0-13.073-10.636-23.71-23.713-23.71c-13.073 0-23.71 10.637-23.71 23.71v8.533h-7.931a2.62 2.62 0 0 0-2.621 2.621v38.565a2.62 2.62 0 0 0 2.621 2.621h64.21a2.62 2.62 0 0 0 2.621-2.621V46.839a2.621 2.621 0 0 0-2.622-2.621m-42.314-8.533c0-5.375 4.371-9.741 9.746-9.741c5.341 0 9.695 4.32 9.747 9.649l-.003.031h.003v8.594H39.791z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Loop(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M83.729 23.57a1.514 1.514 0 0 0-2.393-1.223l-5.944 4.262l-.468.336c-6.405-6.391-15.196-10.389-24.938-10.389c-13.284 0-24.878 7.354-30.941 18.201l.024.013a2.54 2.54 0 0 0 1.028 3.272l8.136 4.697a2.546 2.546 0 0 0 3.48-.932c.006-.011.009-.023.015-.034c3.591-6.404 10.438-10.747 18.289-10.747c4.879 0 9.352 1.696 12.914 4.5l-1.001.719l-5.948 4.262a1.511 1.511 0 0 0 .397 2.655l25.447 8.669a1.51 1.51 0 0 0 1.994-1.433zm-3.825 38.388l-.002-.001l-8.136-4.697a2.546 2.546 0 0 0-3.48.932c-.006.011-.009.023-.015.034c-3.591 6.404-10.438 10.747-18.289 10.747c-4.879 0-9.352-1.696-12.914-4.5l1.001-.719l5.948-4.262a1.511 1.511 0 0 0-.397-2.655l-25.447-8.669a1.51 1.51 0 0 0-1.994 1.433l.092 26.828a1.514 1.514 0 0 0 2.393 1.223l5.944-4.262l.468-.336c6.405 6.391 15.196 10.389 24.938 10.389c13.284 0 24.878-7.354 30.941-18.201l-.025-.012a2.54 2.54 0 0 0-1.026-3.272"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagnifyingGlass(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M56.774 10.391c-17.679 0-32.001 14.329-32.001 32a31.845 31.845 0 0 0 4.588 16.517L13.846 74.423l.054.054c-1.656 1.585-2.673 3.835-2.673 6.378c-.001 4.913 3.913 8.755 8.821 8.754c2.507-.001 4.749-1.004 6.349-2.636l.039.039l16.008-16.009a31.865 31.865 0 0 0 14.33 3.388c17.68 0 31.999-14.327 31.999-32c0-17.671-14.32-32-31.999-32m.194 51.417c-11.05 0-20.001-8.954-20.001-20c0-11.044 8.951-20 20.001-20s19.999 8.955 19.999 20c.001 11.046-8.949 20-19.999 20"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mail(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M85.944 20.189H14.056a2.559 2.559 0 0 0-2.556 2.557v5.144c0 .237.257.509.467.619l37.786 21.583a.63.63 0 0 0 .318.083a.634.634 0 0 0 .324-.088L87.039 28.53c.206-.115.752-.419.957-.559c.248-.169.504-.322.504-.625v-4.601a2.559 2.559 0 0 0-2.556-2.556m2.237 15.457a.644.644 0 0 0-.645.004L66.799 47.851a.637.637 0 0 0-.145.985l20.74 22.357a.632.632 0 0 0 .467.204a.64.64 0 0 0 .639-.639V36.201a.638.638 0 0 0-.319-.555M60.823 51.948a.636.636 0 0 0-.791-.118l-8.312 4.891a3.243 3.243 0 0 1-3.208.021l-7.315-4.179a.64.64 0 0 0-.751.086L12.668 78.415a.64.64 0 0 0 .114 1.02c.432.254.849.375 1.273.375h71.153a.635.635 0 0 0 .585-.385a.635.635 0 0 0-.118-.689zm-26.489-2.347a.639.639 0 0 0-.115-1.023L12.453 36.146a.638.638 0 0 0-.953.556v32.62a.637.637 0 0 0 1.073.468z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Male(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><circle cx="49.354" cy="19.938" r="8.272"/><path d="M65.016 30.435v-.04H35.38v.027a3.895 3.895 0 0 0-3.894 3.867h-.003v23.439h.039a3.477 3.477 0 1 0 6.954 0h.008V41.365h.009a1.104 1.104 0 0 1 1.101-1.064c.598 0 1.078.473 1.103 1.064h.009V84.27a4.067 4.067 0 0 0 8.136-.005V62.661h.03a1.127 1.127 0 1 1 2.253 0h.029v21.604a4.07 4.07 0 0 0 8.137 0V41.366a1.11 1.11 0 1 1 2.222 0v16.363h.038a3.478 3.478 0 1 0 6.956 0h.007v-23.44c.002-2.015-1.535-3.653-3.498-3.854"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MaleFemale(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><circle cx="22.336" cy="19.721" r="8.272"/><path d="M37.999 30.217v-.04H8.363v.027a3.895 3.895 0 0 0-3.894 3.867h-.003V57.51h.039a3.477 3.477 0 1 0 6.954 0h.008V41.148h.009a1.104 1.104 0 0 1 1.101-1.064c.598 0 1.078.473 1.103 1.064h.009v42.905a4.067 4.067 0 0 0 8.136-.005V62.444h.03a1.127 1.127 0 1 1 2.253 0h.029v21.604a4.07 4.07 0 0 0 8.137 0v-42.9a1.11 1.11 0 1 1 2.222 0v16.363h.038a3.478 3.478 0 1 0 6.956 0h.007V34.072c.002-2.015-1.535-3.654-3.498-3.855"/><circle cx="72.297" cy="19.768" r="8.32"/><path d="m95.534 54.161l-5.55-20.712c-.003-.016-.009-.032-.013-.048l-.098-.365l-.014.004c-.506-1.579-1.968-2.729-3.715-2.729c-.014 0-.026.004-.04.004v-.072h-26.12v.068c-1.747 0-3.209 1.151-3.715 2.729l-.014-.004l-.098.364c-.004.017-.01.033-.013.05l-5.55 20.711l.052.014c-.025.171-.052.342-.052.52a3.536 3.536 0 0 0 3.536 3.536a3.529 3.529 0 0 0 3.283-2.244l.025.007l4.071-15.193v.057h.032c.111-.5.536-.882 1.07-.882s.959.382 1.07.882h.033v.164c.002.025.015.047.015.072c0 .025-.013.047-.015.072v.155l-6.518 24.326a1.887 1.887 0 0 0-.187.698l-.014.051l.008.002c-.001.028-.008.054-.008.082c0 1.048.85 1.897 1.898 1.897h4.82v16.08a4.091 4.091 0 0 0 8.182 0v-16.08h2.326v16.08a4.09 4.09 0 0 0 8.182 0v-16.08h4.835a1.897 1.897 0 0 0 1.898-1.897c0-.028-.007-.055-.008-.082l.008-.002l-.014-.051a1.887 1.887 0 0 0-.187-.698l-6.499-24.253l.021-.006c-.027-.096-.06-.191-.06-.296c0-.616.5-1.118 1.118-1.118c.525 0 .946.371 1.065.86l.043-.011l4.065 15.169l.024-.007a3.53 3.53 0 0 0 6.819-1.292c0-.178-.027-.349-.052-.52z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MaleSymbol(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m83.417 18.329l.001-.005c0-.853-.69-1.544-1.542-1.545v-.016H61.078v.017a1.544 1.544 0 0 0-1.53 1.543v6.871c0 .847.684 1.534 1.53 1.542v.001h5.151l-8.713 8.713a26.051 26.051 0 0 0-14.761-4.56c-14.455 0-26.173 11.718-26.173 26.173S28.3 83.237 42.755 83.237c14.455 0 26.173-11.719 26.173-26.174a26.043 26.043 0 0 0-4.398-14.521l8.911-8.911v5.473c0 .853.691 1.544 1.544 1.544h6.887c.852 0 1.544-.691 1.544-1.544l-.001-.005v-20.77zM42.755 73.262c-8.946 0-16.198-7.252-16.198-16.199c0-8.946 7.252-16.199 16.198-16.199s16.199 7.253 16.199 16.199c0 8.947-7.253 16.199-16.199 16.199"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Map(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M91.967 7.961c0-.016.005-.031.005-.047a2.733 2.733 0 0 0-2.73-2.733H39.559v.011L8.031 36.721h-.003v55.365h.003v.001a2.735 2.735 0 0 0 2.734 2.731h78.479a2.73 2.73 0 0 0 2.663-2.15h.06v-.536c0-.015.004-.029.004-.044s-.004-.029-.004-.044zm-24.328 7.177H82.01v24.597L63.897 21.621zM39.57 39.453a2.732 2.732 0 0 0 2.722-2.73V15.138H61.88l-27.17 47.06l-16.725-16.725v-6.02zM17.985 84.862V52.527L32.128 66.67L21.626 84.862zm9.4 0l33.93-58.769l20.696 20.696v38.073z"/><path fill="currentColor" d="M62.03 45.576c-6.645 0-12.026 5.387-12.026 12.027c0 2.659.873 5.109 2.334 7.1l7.759 13.439c.047.094.097.186.157.271l.016.027l.004-.002a2.16 2.16 0 0 0 3.405.132l.02.011l.075-.129a2.25 2.25 0 0 0 .287-.497l7.608-13.178a11.962 11.962 0 0 0 2.39-7.175c-.003-6.639-5.384-12.026-12.029-12.026M61.911 63.7a5.924 5.924 0 0 1-5.926-5.925a5.926 5.926 0 1 1 5.926 5.925"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Marker(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M50 10.417c-15.581 0-28.201 12.627-28.201 28.201a28.074 28.074 0 0 0 5.602 16.873L45.49 86.823c.105.202.21.403.339.588l.04.069l.011-.006a5.063 5.063 0 0 0 4.135 2.111c1.556 0 2.912-.708 3.845-1.799l.047.027l.179-.31c.264-.356.498-.736.667-1.155L72.475 55.65a28.074 28.074 0 0 0 5.726-17.032c0-15.574-12.62-28.201-28.201-28.201m-.279 42.498c-7.677 0-13.895-6.221-13.895-13.895c0-7.673 6.218-13.895 13.895-13.895s13.895 6.222 13.895 13.895c0 7.673-6.218 13.895-13.895 13.895"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Megaphone(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M27.328 28.507c-8.252 0-14.941 6.689-14.941 14.941s6.689 14.941 14.941 14.941a2.712 2.712 0 0 0 2.676-2.296h.043V31.227a2.72 2.72 0 0 0-2.719-2.72m60.286-12.193a2.72 2.72 0 0 0-2.719-2.719c-.623 0-1.19.218-1.648.57L59.321 27.979H37.753a2.719 2.719 0 0 0-2.719 2.719v25.809h.034a2.71 2.71 0 0 0 2.685 2.383h19.138l26.16 15.103c.485.45 1.13.731 1.844.731a2.72 2.72 0 0 0 2.719-2.719c0-.103-.019-.201-.03-.301h.03V16.472h-.016c.003-.053.016-.103.016-.158M56.212 79.572l-.038-.066c-.013-.024-.019-.049-.033-.073s-.032-.042-.046-.064L47.79 64.986l-.031.018a2.7 2.7 0 0 0-1.857-1.096v-.03H38.68v.018a2.693 2.693 0 0 0-1.135.347a2.718 2.718 0 0 0-.995 3.714c.064.11.142.206.219.303l-.032.019l9.678 16.764l.002.005a.014.014 0 0 0 .003.004l.097.168l.017-.01a2.71 2.71 0 0 0 3.596.832c.145-.084.276-.183.399-.287l.025.043l4.983-2.876l-.023-.04a2.708 2.708 0 0 0 .691-3.306z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Microphone(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M75.46 49.16h.003V34.787a3.783 3.783 0 0 0-7.565 0v14.256c0 9.885-8.013 17.897-17.898 17.897s-17.898-8.013-17.898-17.897V34.787a3.783 3.783 0 0 0-7.565 0V49.16h.003c.058 12.724 9.447 23.243 21.678 25.065v5.438H32.839v.003a3.78 3.78 0 0 0 0 7.558v.003h34.29a3.782 3.782 0 1 0 0-7.564H53.782v-5.438C66.013 72.403 75.403 61.884 75.46 49.16"/><path d="M37.186 48.941c0 7.088 5.745 12.833 12.833 12.833c7.087 0 12.831-5.746 12.831-12.833c0-.096-.012-.188-.014-.283h.053V25.322h-.053c-.153-6.955-5.826-12.549-12.817-12.549c-6.992 0-12.666 5.594-12.819 12.549h-.052v23.336h.052c-.001.095-.014.187-.014.283"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Minus(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M88.447 38.528H11.554a2.024 2.024 0 0 0-2.024 2.024v18.896c0 1.118.907 2.024 2.024 2.024h76.892a2.024 2.024 0 0 0 2.023-2.024V40.552a2.023 2.023 0 0 0-2.022-2.024"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusCircle(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M67.441 44.796H32.559a.918.918 0 0 0-.918.918v8.572c0 .507.411.918.918.918h34.883a.918.918 0 0 0 .918-.918v-8.572a.92.92 0 0 0-.919-.918"/><path fill="currentColor" d="M50 22.44c15.196 0 27.56 12.366 27.56 27.562C77.56 65.199 65.196 77.56 50 77.56S22.44 65.199 22.44 50.002C22.44 34.806 34.804 22.44 50 22.44m0-9.94c-20.709 0-37.5 16.793-37.5 37.502C12.5 70.712 29.291 87.5 50 87.5s37.5-16.788 37.5-37.498C87.5 29.293 70.709 12.5 50 12.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mobile(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M69.328 17.517H30.567v.01a2.495 2.495 0 0 0-2.396 2.49v59.967a2.495 2.495 0 0 0 2.396 2.489v.011h38.761a2.5 2.5 0 0 0 2.5-2.5V20.017a2.5 2.5 0 0 0-2.5-2.5M50.059 79.9a2.45 2.45 0 1 1 0-4.9a2.45 2.45 0 0 1 0 4.9m11.813-7.395H38.128V27.473h23.743v45.032z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MobileSignal(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M69.328 29.884H30.567v.01a2.495 2.495 0 0 0-2.396 2.49v59.967a2.495 2.495 0 0 0 2.396 2.489v.011h38.761a2.5 2.5 0 0 0 2.5-2.5V32.384a2.5 2.5 0 0 0-2.5-2.5M50.059 92.268a2.45 2.45 0 1 1 0-4.9a2.45 2.45 0 0 1 0 4.9m11.813-7.396H38.128V39.84h23.743v45.032zm7.731-69.176a1.488 1.488 0 0 0-.006-1.85l.024-.024l-.999-.999l-.001-.001l-.078-.078l-.004.004c-10.238-10.154-26.825-10.132-37.032.073c-.015.015-.019.036-.034.051l-.008-.008l-.955.955l-.003.002l-.002.003l-.118.119l.024.024a1.502 1.502 0 0 0 .08 1.979l-.004.004l1.918 1.918l.026-.026c.55.428 1.303.427 1.85-.006l.025.024l1.079-1.078l-.022-.022c.015-.015.036-.019.051-.034c8.053-8.052 21.135-8.074 29.221-.075l-.009.009l.956.957v.001h.001l.12.121l.024-.024a1.503 1.503 0 0 0 1.979-.079l.004.004l.021-.021h.001v-.001L69.5 15.85l.128-.128z"/><path fill="currentColor" d="m62.208 21.235l.024-.024l-1.079-1.079l-.004.004c-6.165-6.078-16.12-6.054-22.251.076c-.015.015-.019.035-.033.05l-.009-.008l-1.079 1.079l.024.024a1.501 1.501 0 0 0 .079 1.979l-.004.004l1.918 1.918l.026-.026c.55.428 1.303.427 1.85-.006l.025.025l1.079-1.078l-.022-.022c.015-.015.036-.019.051-.034c3.977-3.977 10.425-3.995 14.436-.072l-.005.005l.957.957h.001l.121.121l.024-.025a1.503 1.503 0 0 0 1.979-.079l.004.004l.021-.021l1.897-1.897l-.026-.026a1.484 1.484 0 0 0-.004-1.849"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Monitor(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M90.315 12.993H9.684a5.644 5.644 0 0 0-5.644 5.645V67.83a5.644 5.644 0 0 0 5.644 5.645h30.359v8.556h-9.402c-.892 0-1.613.721-1.613 1.612v1.751c0 .892.721 1.613 1.613 1.613h37.901c.891 0 1.613-.721 1.613-1.613v-1.751c0-.892-.722-1.612-1.613-1.612h-8.586v-8.556h30.359a5.643 5.643 0 0 0 5.645-5.645V18.638a5.645 5.645 0 0 0-5.645-5.645M14.091 63.508V22.949h71.818v40.559z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mountains(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m95.294 78.127l.026-.015L62.074 20.53l-.014.008c-.732-1.502-2.26-2.547-4.044-2.547a4.506 4.506 0 0 0-4.107 2.664l-22.125 38.32l-5.417-9.382l-.007.004a2.373 2.373 0 0 0-2.133-1.344c-.964 0-1.791.578-2.164 1.402L6.271 77.009H6.2v.122l-1.529 2.648a1.365 1.365 0 0 0-.285.832c0 .764.619 1.383 1.382 1.383c.152 0 .295-.03.432-.076v.091h87v-.05a2.613 2.613 0 0 0 2.413-2.601a2.588 2.588 0 0 0-.319-1.231"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Music(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M84.105 13.627H32.938a3.196 3.196 0 0 0-3.195 3.195v48.069a16.451 16.451 0 0 0-3.543-.394c-7.456 0-13.5 4.896-13.5 10.938c0 6.041 6.044 10.937 13.5 10.937c7.455 0 13.5-4.896 13.5-10.937V29.257h37.644v26.401a16.451 16.451 0 0 0-3.543-.394c-7.456 0-13.5 4.896-13.5 10.938s6.044 10.937 13.5 10.937c7.455 0 13.5-4.896 13.5-10.937V16.823a3.197 3.197 0 0 0-3.196-3.196"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Next(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M88.765 77.291v-54.72l.001-.005l-.001-.005v-.235h-.024a2.529 2.529 0 0 0-2.513-2.296c-.03 0-.058.008-.087.009h-4.569c-.071-.006-.14-.021-.213-.021a2.538 2.538 0 0 0-2.537 2.537c0 .117.019.23.035.343v23.141l-33.341-19.25a1.715 1.715 0 0 0-2.92 1.172h-.006V43.21L14.148 26.788a1.715 1.715 0 0 0-2.92 1.172h-.005v44.031a1.718 1.718 0 0 0 2.948 1.197L42.59 56.78v15.212a1.718 1.718 0 0 0 2.948 1.197l33.319-19.237v23.382h.013c-.002.039-.012.075-.012.114a2.537 2.537 0 0 0 2.532 2.536h4.852v-.024a2.537 2.537 0 0 0 2.537-2.537c-.001-.046-.012-.088-.014-.132"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoDogs(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M50 12.5c-20.712 0-37.5 16.793-37.5 37.502C12.5 70.712 29.288 87.5 50 87.5s37.5-16.788 37.5-37.498C87.5 29.293 70.712 12.5 50 12.5m0 65.061c-15.199 0-27.56-12.362-27.56-27.559a27.437 27.437 0 0 1 6.4-17.636l12.803 12.803h-4.834v.009a.671.671 0 0 0-.184.042l-6.043-3.489l-.005.008a1.402 1.402 0 0 0-.648-.164a1.418 1.418 0 0 0-.765 2.609l-.01.018l6.124 3.535c-.237 1.03-.378 2.294-.378 3.667c0 2.898.625 5.328 1.469 6.025l-3.454 5.98a1.916 1.916 0 0 0-.413 1.182c0 1.069.866 1.935 1.935 1.935c.626 0 1.176-.302 1.529-.762l.032.018l4.037-6.992l4.037 6.992l.032-.019c.354.461.904.763 1.53.763a1.935 1.935 0 0 0 1.936-1.935a1.92 1.92 0 0 0-.415-1.184l-2.609-4.518a77.618 77.618 0 0 0 8.606.009l-2.604 4.509a1.914 1.914 0 0 0-.414 1.183c0 1.069.866 1.935 1.935 1.935c.626 0 1.176-.302 1.53-.763l.032.019l3.157-5.467l4.167 4.167l.751 1.3l.031-.018c.296.386.73.654 1.232.732l4.665 4.665A27.427 27.427 0 0 1 50 77.561m21.161-9.926l-9.305-9.305l-.459-.795c.919-.518 1.619-3.062 1.619-6.13c0-.688-.036-1.348-.101-1.966l1.515-1.515h5.44v-.015l1.198-1.196l-.006-.005a1.37 1.37 0 0 0-.144-2.106l-2.151-2.15v-2.24c0-.841-.483-1.562-1.181-1.922l-2.789-1.609l-.007.012a1.674 1.674 0 0 0-.118-.064l-1.906-3.301l-3.411 5.911h.199l-5.933 5.932h-4.926l-16.33-16.33a27.426 27.426 0 0 1 17.634-6.4c15.198 0 27.56 12.367 27.56 27.562a27.423 27.423 0 0 1-6.398 17.632"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoSmoking(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m65.551 45.894l-.008.001h-.657l-.008-.001c-.63 0-1.14.51-1.14 1.14l.001 5.932v.154h.03c.077.553.533.986 1.108.986h.668v-.001l.005.001c.575 0 1.031-.433 1.109-.987h.031v-6.086c0-.63-.51-1.139-1.139-1.139m4.24 0l-.008.001h-.656l-.008-.001c-.63 0-1.14.51-1.14 1.14h-.001l.001 6.086h.031c.077.553.533.986 1.108.986h.668v-.001l.005.001c.575 0 1.032-.433 1.109-.987h.031v-6.086a1.14 1.14 0 0 0-1.14-1.139"/><path d="M50 12.5c-20.712 0-37.5 16.793-37.5 37.502C12.5 70.712 29.288 87.5 50 87.5s37.5-16.788 37.5-37.498C87.5 29.293 70.712 12.5 50 12.5m0 65.061c-15.199 0-27.56-12.362-27.56-27.559a27.435 27.435 0 0 1 6.4-17.636l13.529 13.529h-12.15l-.008-.001c-.63 0-1.14.51-1.14 1.14l.001.003h-.002v5.93c0 .63.51 1.14 1.14 1.14h20.37l17.055 17.055A27.44 27.44 0 0 1 50 77.561m21.161-9.926L57.633 54.106h3.463c.629 0 1.14-.51 1.14-1.14v-5.929l.001-.004c0-.63-.511-1.14-1.14-1.14l-.008.001H49.422L32.366 28.839A27.431 27.431 0 0 1 50 22.439c15.198 0 27.56 12.367 27.56 27.562a27.43 27.43 0 0 1-6.399 17.634"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Page(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M80.277 12.498h-.005v-4.02a1.73 1.73 0 0 0-1.73-1.73h-32.87l-25.95 25.95v58.819c0 .956.774 1.73 1.73 1.73h57.09a1.73 1.73 0 0 0 1.73-1.73v-2.448h.005zM29.679 83.292V36.158h17.723a1.73 1.73 0 0 0 1.73-1.73V16.705H70.32v66.587z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PageAdd(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M73.682 78.553h-6.49a1.73 1.73 0 0 0-1.73 1.73h-.007v3.01H24.814V36.16h17.723a1.73 1.73 0 0 0 1.73-1.73V16.707h21.188v20.627h.008a1.73 1.73 0 0 0 1.73 1.726h6.49c.954 0 1.727-.772 1.729-1.726V12.5h-.005V8.48a1.73 1.73 0 0 0-1.73-1.73h-32.87L14.857 32.7v58.819c0 .956.774 1.73 1.73 1.73h57.089a1.73 1.73 0 0 0 1.73-1.73v-2.448h.005v-8.79a1.728 1.728 0 0 0-1.729-1.728"/><path fill="currentColor" d="M83.413 53.562h-7.785v-7.785a1.73 1.73 0 0 0-1.73-1.73h-6.92a1.73 1.73 0 0 0-1.73 1.73v7.785h-7.785a1.73 1.73 0 0 0-1.73 1.73v6.92c0 .955.774 1.73 1.73 1.73h7.785v7.785a1.73 1.73 0 0 0 1.73 1.731h6.92a1.73 1.73 0 0 0 1.73-1.731v-7.785h7.785a1.73 1.73 0 0 0 1.73-1.73v-6.92a1.73 1.73 0 0 0-1.73-1.73"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PageCopy(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M75.272 7.482h-.005v-4.02a1.73 1.73 0 0 0-1.73-1.73h-32.87l-25.95 25.95v58.819c0 .956.774 1.73 1.73 1.73h57.089a1.73 1.73 0 0 0 1.73-1.73v-2.448h.005zM24.674 78.276V31.142h17.723a1.73 1.73 0 0 0 1.73-1.73V11.689h21.188v66.587z"/><path fill="currentColor" d="M83.77 24.857h-3.475v66.911c0 .835-.677 1.513-1.513 1.513H29.306v3.475c0 .836.677 1.513 1.513 1.513H83.77c.836 0 1.513-.677 1.513-1.513V26.37c0-.836-.677-1.513-1.513-1.513"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PageCsv(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M40.416 64.173c.831 0 1.55.623 1.786 1.342l2.408-1.121c-.553-1.273-1.771-2.685-4.193-2.685c-2.893 0-5.079 1.924-5.079 4.775c0 2.837 2.187 4.774 5.079 4.774c2.422 0 3.654-1.467 4.193-2.699l-2.408-1.107c-.235.719-.955 1.342-1.786 1.342c-1.342 0-2.242-1.024-2.242-2.311s.9-2.31 2.242-2.31m8.189.319c0-.235.152-.415.706-.415c.872 0 1.91.304 2.712.913l1.495-1.979c-1.052-.858-2.408-1.287-3.917-1.287c-2.533 0-3.833 1.495-3.833 3.059c0 3.64 5.148 2.74 5.148 3.626c0 .36-.498.499-1.024.499a4.304 4.304 0 0 1-2.976-1.19l-1.453 2.076c.982.886 2.325 1.467 4.291 1.467c2.477 0 3.986-1.176 3.986-3.211c0-3.434-5.135-2.686-5.135-3.558m10.896 3.584l-1.993-6.214h-3.169l3.404 9.231h3.515l3.405-9.231h-3.169z"/><path fill="currentColor" d="M80.277 12.5h-.005V8.48a1.73 1.73 0 0 0-1.73-1.73h-32.87L19.722 32.7v58.819c0 .956.774 1.73 1.73 1.73h57.09a1.73 1.73 0 0 0 1.73-1.73v-2.448h.005zM29.679 83.294V36.16h17.723a1.73 1.73 0 0 0 1.73-1.73V16.707H70.32v66.587z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PageDelete(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M73.603 78.553h-6.49a1.73 1.73 0 0 0-1.73 1.73h-.007v3.01H24.735V36.16h17.723a1.73 1.73 0 0 0 1.73-1.73V16.707h21.188v20.627h.008a1.73 1.73 0 0 0 1.73 1.726h6.49c.954 0 1.727-.772 1.729-1.726V12.5h-.005V8.48a1.73 1.73 0 0 0-1.73-1.73H40.727L14.778 32.7v58.819c0 .956.774 1.73 1.73 1.73h57.089a1.73 1.73 0 0 0 1.73-1.73v-2.448h.005v-8.79a1.728 1.728 0 0 0-1.729-1.728"/><path fill="currentColor" d="m79.21 58.751l5.505-5.505a1.73 1.73 0 0 0 0-2.447l-4.894-4.893a1.729 1.729 0 0 0-2.446 0l-5.505 5.505l-5.505-5.505a1.729 1.729 0 0 0-2.446 0L59.026 50.8a1.73 1.73 0 0 0 0 2.447l5.505 5.505l-5.505 5.505a1.729 1.729 0 0 0 0 2.446l4.894 4.893c.676.676 1.77.676 2.446 0l5.505-5.505l5.505 5.505c.676.676 1.77.676 2.446 0l4.894-4.893a1.729 1.729 0 0 0 0-2.446z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PageDoc(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M80.277 12.5h-.005V8.48a1.73 1.73 0 0 0-1.73-1.73h-32.87L19.722 32.7v58.819c0 .956.774 1.73 1.73 1.73h57.09a1.73 1.73 0 0 0 1.73-1.73v-2.448h.005zM29.679 83.294V36.16h17.723a1.73 1.73 0 0 0 1.73-1.73V16.707H70.32v66.587z"/><path fill="currentColor" d="M38.663 61.862h-4.124v9.231h4.138c2.893 0 5.052-1.675 5.052-4.623s-2.16-4.608-5.066-4.608m0 6.796h-1.329v-4.36h1.342c1.495 0 2.214.927 2.214 2.173c.001 1.162-.829 2.187-2.227 2.187m11.295-6.948c-2.851 0-5.052 1.938-5.052 4.775c0 2.837 2.201 4.774 5.052 4.774c2.851 0 5.051-1.938 5.051-4.774c.001-2.838-2.2-4.775-5.051-4.775m0 7.086c-1.343 0-2.214-1.024-2.214-2.311s.872-2.312 2.214-2.312s2.214 1.024 2.214 2.312s-.871 2.311-2.214 2.311m11.31-4.623c.831 0 1.55.623 1.786 1.342l2.408-1.121c-.554-1.273-1.771-2.685-4.194-2.685c-2.892 0-5.079 1.924-5.079 4.775c0 2.837 2.187 4.774 5.079 4.774c2.422 0 3.654-1.467 4.194-2.699l-2.408-1.107c-.235.719-.955 1.342-1.786 1.342c-1.342 0-2.242-1.024-2.242-2.311s.899-2.31 2.242-2.31"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PageEdit(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M67.041 78.553h-6.49a1.73 1.73 0 0 0-1.73 1.73h-.007v3.01H18.173V36.16h17.723a1.73 1.73 0 0 0 1.73-1.73V16.707h21.188v12.34h.016a1.726 1.726 0 0 0 1.721 1.641h6.49c.925 0 1.674-.728 1.721-1.641h.009V12.499h-.005V8.48a1.73 1.73 0 0 0-1.73-1.73H34.165L8.216 32.7v58.819c0 .956.774 1.73 1.73 1.73h57.089a1.73 1.73 0 0 0 1.73-1.73v-2.448h.005v-8.79a1.728 1.728 0 0 0-1.729-1.728"/><path fill="currentColor" d="M91.277 39.04L79.656 27.419a1.73 1.73 0 0 0-2.447 0L45.404 59.224l.069.069l-.109-.029l-4.351 16.237l.003.001a1.722 1.722 0 0 0 .412 1.765a1.723 1.723 0 0 0 1.948.341l.002.006l16.08-4.309l-.01-.037l.023.024l31.806-31.806a1.729 1.729 0 0 0 0-2.446M46.305 72.353l2.584-9.643l7.059 7.059z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PageExport(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M94.284 65.553L75.825 52.411a1.255 1.255 0 0 0-1.312-.093c-.424.218-.684.694-.685 1.173l.009 6.221H57.231c-.706 0-1.391.497-1.391 1.204v11.442c0 .707.685 1.194 1.391 1.194h16.774v6.27c0 .478.184.917.609 1.136c.425.219.853.182 1.242-.096l18.432-13.228c.335-.239.477-.626.477-1.038v-.002c0-.415-.144-.801-.481-1.041"/><path fill="currentColor" d="M64.06 78.553h-6.49a1.73 1.73 0 0 0-1.73 1.73h-.007v3.01H15.191V36.16h17.723a1.73 1.73 0 0 0 1.73-1.73V16.707h21.188v36.356h.011a1.728 1.728 0 0 0 1.726 1.691h6.49c.943 0 1.705-.754 1.726-1.691h.004V12.501h-.005V8.48a1.73 1.73 0 0 0-1.73-1.73h-32.87L5.235 32.7v58.819c0 .956.774 1.73 1.73 1.73h57.089a1.73 1.73 0 0 0 1.73-1.73v-2.448h.005v-8.79a1.729 1.729 0 0 0-1.729-1.728"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PageExportCsv(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M94.284 65.553L75.825 52.411a1.255 1.255 0 0 0-1.312-.093c-.424.218-.684.694-.685 1.173l.009 6.221H57.231c-.706 0-1.391.497-1.391 1.204v11.442c0 .707.685 1.194 1.391 1.194h16.774v6.27c0 .478.184.917.609 1.136s.853.182 1.242-.097l18.432-13.228c.335-.239.477-.626.477-1.038v-.002c0-.414-.144-.8-.481-1.04"/><path fill="currentColor" d="M64.06 78.553h-6.49a1.73 1.73 0 0 0-1.73 1.73h-.007v3.01H15.191V36.16h17.723a1.73 1.73 0 0 0 1.73-1.73V16.707h21.188v36.356h.011a1.728 1.728 0 0 0 1.726 1.691h6.49c.943 0 1.705-.754 1.726-1.691h.004V12.5h-.005V8.48a1.73 1.73 0 0 0-1.73-1.73h-32.87L5.235 32.7v58.819c0 .956.774 1.73 1.73 1.73h57.089a1.73 1.73 0 0 0 1.73-1.73v-2.448h.005v-8.79a1.729 1.729 0 0 0-1.729-1.728"/><path fill="currentColor" d="M26.18 64.173c.831 0 1.55.623 1.786 1.342l2.408-1.121c-.553-1.273-1.771-2.685-4.193-2.685c-2.893 0-5.079 1.924-5.079 4.775c0 2.837 2.187 4.774 5.079 4.774c2.422 0 3.654-1.467 4.193-2.699l-2.408-1.107c-.235.719-.955 1.342-1.786 1.342c-1.342 0-2.242-1.024-2.242-2.311s.899-2.31 2.242-2.31m9.476 4.734a4.304 4.304 0 0 1-2.976-1.19l-1.453 2.076c.982.886 2.325 1.467 4.291 1.467c2.477 0 3.986-1.176 3.986-3.211c0-3.432-5.135-2.685-5.135-3.557c0-.235.152-.415.706-.415c.872 0 1.91.304 2.712.913l1.495-1.979c-1.052-.858-2.408-1.287-3.917-1.287c-2.533 0-3.833 1.495-3.833 3.059c0 3.64 5.148 2.74 5.148 3.626c0 .359-.498.498-1.024.498m7.615-7.045h-3.169l3.404 9.231h3.516l3.404-9.231h-3.169l-1.993 6.214z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PageExportDoc(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M94.284 65.553L75.825 52.411a1.255 1.255 0 0 0-1.312-.093c-.424.218-.684.694-.685 1.173l.009 6.221H57.231c-.706 0-1.391.497-1.391 1.204v11.442c0 .707.685 1.194 1.391 1.194h16.774v6.27c0 .478.184.917.609 1.136c.425.219.853.182 1.242-.097l18.432-13.228c.335-.239.477-.626.477-1.038v-.002c0-.414-.144-.8-.481-1.04"/><path fill="currentColor" d="M64.06 78.553h-6.49a1.73 1.73 0 0 0-1.73 1.73h-.007v3.01H15.191V36.16h17.723a1.73 1.73 0 0 0 1.73-1.73V16.707h21.188v36.356h.011a1.728 1.728 0 0 0 1.726 1.691h6.49c.943 0 1.705-.754 1.726-1.691h.004V12.5h-.005V8.48a1.73 1.73 0 0 0-1.73-1.73h-32.87L5.235 32.7v58.819c0 .956.774 1.73 1.73 1.73h57.089a1.73 1.73 0 0 0 1.73-1.73v-2.448h.005v-8.79a1.729 1.729 0 0 0-1.729-1.728"/><path fill="currentColor" d="M20.364 61.862v9.231h4.138c2.893 0 5.052-1.675 5.052-4.623s-2.159-4.608-5.065-4.608zm6.352 4.609c0 1.163-.83 2.187-2.228 2.187H23.16v-4.36h1.342c1.495 0 2.214.927 2.214 2.173m4.017.014c0 2.837 2.201 4.774 5.052 4.774c2.851 0 5.051-1.938 5.051-4.774c0-2.837-2.201-4.775-5.051-4.775c-2.852 0-5.052 1.937-5.052 4.775m7.266 0c0 1.287-.872 2.311-2.214 2.311c-1.343 0-2.214-1.024-2.214-2.311s.872-2.312 2.214-2.312s2.214 1.025 2.214 2.312m9.094-2.312c.831 0 1.55.623 1.786 1.342l2.408-1.121c-.554-1.273-1.771-2.685-4.194-2.685c-2.892 0-5.079 1.924-5.079 4.775c0 2.837 2.187 4.774 5.079 4.774c2.422 0 3.654-1.467 4.194-2.699l-2.408-1.107c-.235.719-.955 1.342-1.786 1.342c-1.342 0-2.242-1.024-2.242-2.311s.9-2.31 2.242-2.31"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PageExportPdf(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M94.284 65.553L75.825 52.411a1.255 1.255 0 0 0-1.312-.093c-.424.218-.684.694-.685 1.173l.009 6.221H57.231c-.706 0-1.391.497-1.391 1.204v11.442c0 .707.685 1.194 1.391 1.194h16.774v6.27c0 .478.184.917.609 1.136s.853.182 1.242-.097l18.432-13.228c.335-.239.477-.626.477-1.038v-.002c0-.414-.144-.8-.481-1.04"/><path fill="currentColor" d="M64.06 78.553h-6.49a1.73 1.73 0 0 0-1.73 1.73h-.007v3.01H15.191V36.16h17.723a1.73 1.73 0 0 0 1.73-1.73V16.707h21.188v36.356h.011a1.728 1.728 0 0 0 1.726 1.691h6.49c.943 0 1.705-.754 1.726-1.691h.004V12.5h-.005V8.48a1.73 1.73 0 0 0-1.73-1.73h-32.87L5.235 32.7v58.819c0 .956.774 1.73 1.73 1.73h57.089a1.73 1.73 0 0 0 1.73-1.73v-2.448h.005v-8.79a1.729 1.729 0 0 0-1.729-1.728"/><path fill="currentColor" d="M21.525 61.862v9.231h2.795v-2.906h2.131c2.159 0 3.321-1.439 3.321-3.156c0-1.73-1.162-3.169-3.321-3.169zm5.411 3.169c0 .484-.374.72-.844.72H24.32v-1.453h1.771c.471 0 .845.235.845.733m4.292-3.169v9.231h4.138c2.893 0 5.052-1.675 5.052-4.623s-2.159-4.608-5.065-4.608zm6.352 4.609c0 1.163-.83 2.187-2.228 2.187h-1.329v-4.36h1.342c1.495 0 2.215.927 2.215 2.173m11.536-2.173v-2.436h-7.003v9.231h2.795v-3.446h4.11v-2.436h-4.11v-.913z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PageFilled(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.455 36.16h25.95c.956 0 1.729-.774 1.729-1.73V8.48a1.73 1.73 0 0 0-1.729-1.73h-1.73L19.725 32.7v1.729a1.73 1.73 0 0 0 1.73 1.731"/><path fill="currentColor" d="M78.545 6.75H60.821c-.806 0-1.476.553-1.669 1.298h-.061v36.551h-.002l.002.023a1.73 1.73 0 0 1-1.73 1.73l-.009-.001v.001H21.455c-.873 0-1.586.647-1.706 1.487h-.024V91.52c0 .956.774 1.73 1.73 1.73h57.09a1.73 1.73 0 0 0 1.73-1.73V8.48a1.73 1.73 0 0 0-1.73-1.73"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PageMultiple(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M86.45 23.27h-3.475v66.91c0 .835-.677 1.513-1.513 1.513H31.987v3.475c0 .836.677 1.513 1.513 1.513h52.951c.836 0 1.513-.677 1.513-1.513V24.782a1.513 1.513 0 0 0-1.514-1.512"/><path fill="currentColor" d="M77.988 85.193V14.807c0-.836-.677-1.513-1.513-1.513H73v66.911c0 .836-.677 1.513-1.513 1.513H22.011v3.475c0 .836.677 1.513 1.513 1.513h52.951c.836 0 1.513-.677 1.513-1.513"/><path fill="currentColor" d="M68.013 75.218V4.832c0-.836-.677-1.513-1.513-1.513H13.55c-.836 0-1.513.677-1.513 1.513v70.386c0 .836.677 1.513 1.513 1.513H66.5c.836 0 1.513-.677 1.513-1.513"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PagePdf(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M41.131 61.862h-4.927v9.231H39v-2.906h2.131c2.159 0 3.321-1.439 3.321-3.156c.001-1.73-1.162-3.169-3.321-3.169m-.36 3.889H39v-1.453h1.771c.471 0 .844.235.844.733c.001.485-.373.72-.844.72m9.261-3.889h-4.124v9.231h4.138c2.893 0 5.052-1.675 5.052-4.623s-2.16-4.608-5.066-4.608m0 6.796h-1.329v-4.36h1.342c1.495 0 2.214.927 2.214 2.173c.001 1.162-.829 2.187-2.227 2.187m6.76 2.435h2.796v-3.446h4.11v-2.436h-4.11v-.913h4.207v-2.436h-7.003z"/><path fill="currentColor" d="M80.277 12.5h-.005V8.48a1.73 1.73 0 0 0-1.73-1.73h-32.87L19.722 32.7v58.819c0 .956.774 1.73 1.73 1.73h57.09a1.73 1.73 0 0 0 1.73-1.73v-2.448h.005zM29.679 83.294V36.16h17.723a1.73 1.73 0 0 0 1.73-1.73V16.707H70.32v66.587z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PageRemove(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M75.412 70.658v-.033h-.003a1.728 1.728 0 0 0-1.727-1.696h-6.49c-.944 0-1.708.757-1.727 1.696h-.01v12.668H24.814V36.16h17.723a1.73 1.73 0 0 0 1.73-1.73V16.707h21.188v30.2h.013a1.727 1.727 0 0 0 1.724 1.668h6.49c.935 0 1.69-.742 1.724-1.668h.006V12.501h-.005V8.48a1.73 1.73 0 0 0-1.73-1.73h-32.87L14.857 32.7v58.819c0 .956.774 1.73 1.73 1.73h57.089a1.73 1.73 0 0 0 1.73-1.73v-2.448h.005V70.658z"/><path fill="currentColor" d="M83.413 53.562h-25.95a1.73 1.73 0 0 0-1.73 1.73v6.92c0 .955.774 1.73 1.73 1.73h25.95a1.73 1.73 0 0 0 1.73-1.73v-6.92a1.73 1.73 0 0 0-1.73-1.73"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PageSearch(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M89.307 12.5h-.005V8.48a1.73 1.73 0 0 0-1.73-1.73H54.701L28.752 32.7v5.237h.014a1.724 1.724 0 0 0 1.716 1.588h6.49c.907 0 1.642-.7 1.716-1.588h.021V36.16h17.723a1.73 1.73 0 0 0 1.73-1.73V16.707H79.35v66.587H38.708v-3.012h-.007a1.73 1.73 0 0 0-1.73-1.728h-6.49a1.73 1.73 0 0 0-1.73 1.728v11.239c0 .956.774 1.73 1.73 1.73H87.57a1.73 1.73 0 0 0 1.73-1.73v-2.448h.005z"/><path fill="currentColor" d="M49.6 59.04c0-8.847-7.169-16.021-16.02-16.021c-8.851 0-16.02 7.174-16.02 16.021c0 2.99.834 5.78 2.26 8.175l-7.815 7.815l.031.031a4.38 4.38 0 0 0-1.344 3.194c0 2.463 1.962 4.39 4.423 4.39a4.427 4.427 0 0 0 3.181-1.324l.022.022l8.024-8.024a15.933 15.933 0 0 0 7.238 1.741c8.851-.001 16.02-7.173 16.02-16.02m-25.318-.001a9.298 9.298 0 0 1 9.299-9.298a9.297 9.297 0 0 1 9.298 9.298a9.297 9.297 0 0 1-9.298 9.299a9.297 9.297 0 0 1-9.299-9.299"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaintBucket(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m83.095 47.878l-.004-.003l-32.088-32.09h-.001a2.518 2.518 0 0 0-3.562 0l-5.26 5.26L30.934 9.799a4.924 4.924 0 0 0-3.571-1.534c-2.743-.001-4.966 2.231-4.964 4.986a4.97 4.97 0 0 0 1.558 3.612L35.16 28.068L7.883 55.344a2.518 2.518 0 0 0 0 3.562l32.091 32.092a2.518 2.518 0 0 0 3.562 0l.001-.001L83.095 51.44a2.52 2.52 0 0 0 0-3.562m-19.704 9.228H20.233l29.003-29.004l21.579 21.58zm27.682 16.629l-5.97-10.339c-.031-.058-.061-.117-.098-.171l-.015-.025l-.004.002a1.654 1.654 0 0 0-2.607-.102l-.016-.009l-.059.103a1.705 1.705 0 0 0-.217.375l-5.835 10.105a9.159 9.159 0 0 0-1.829 5.493a9.215 9.215 0 0 0 9.216 9.217a9.216 9.216 0 0 0 9.217-9.217a9.163 9.163 0 0 0-1.783-5.432"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paperclip(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M30.267 88.384c-2.726 0-6.235-.922-9.739-4.426c-4.621-4.621-5.203-9.081-4.878-12.013c.381-3.436 2.14-6.747 5.225-9.833L50.287 32.7c8.215-8.212 13.68-5.208 16.206-2.684c3.056 3.057 5.06 8.338-2.747 16.143L36.668 73.237l-5.271-5.271l27.078-27.078c4.172-4.174 3.211-5.135 2.745-5.601c-.369-.369-1.49-1.485-5.662 2.684L26.146 67.383c-1.279 1.281-2.852 3.252-3.087 5.387c-.213 1.922.682 3.859 2.743 5.916c2.555 2.558 4.487 2.286 5.125 2.194c1.74-.243 3.725-1.442 5.739-3.456c2.495-2.489 32.596-32.591 34.901-34.901c2.849-2.849 4.66-5.62 5.383-8.236c1.048-3.786-.206-7.387-3.83-11.013c-3.825-3.825-9.986-7.829-19.308 1.495L21.601 56.982a3.726 3.726 0 0 1-5.271 0a3.725 3.725 0 0 1 0-5.271l32.213-32.212c9.957-9.96 20.836-10.508 29.849-1.495c6.989 6.989 6.981 13.789 5.747 18.269c-1.087 3.931-3.473 7.697-7.297 11.522c-2.308 2.31-32.414 32.416-34.904 34.902c-3.208 3.213-6.564 5.086-9.967 5.567a11.74 11.74 0 0 1-1.704.12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pause(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M40.531 19.959c-.097 0-.19.016-.283.029v-.029H29.911a2.163 2.163 0 0 0-2.163 2.163v.005h-.001V77.87h.001l-.001.009c0 1.194.969 2.162 2.163 2.162h10.62a2.162 2.162 0 0 0 2.162-2.162l-.001-.009h.001V22.127l.001-.005a2.163 2.163 0 0 0-2.162-2.163m31.72 57.911V22.127l.001-.005a2.163 2.163 0 0 0-2.162-2.163c-.097 0-.19.016-.283.029v-.029H59.47a2.163 2.163 0 0 0-2.163 2.163v.005h-.001V77.87h.001l-.001.009c0 1.194.969 2.162 2.163 2.162h10.618a2.162 2.162 0 0 0 2.162-2.162c.002-.003.001-.006.002-.009"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paw(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M34.848 40.708c0-5.6-4.542-10.141-10.143-10.141c-5.601 0-10.141 4.541-10.141 10.141c0 5.604 4.539 10.143 10.141 10.143s10.143-4.539 10.143-10.143m40.445-8.16c-5.6 0-10.141 4.541-10.141 10.141c0 5.604 4.541 10.141 10.141 10.141c5.601 0 10.142-4.537 10.142-10.141c0-5.6-4.54-10.141-10.142-10.141m-9.211 21.43c-.705-.869-1.703-1.875-2.849-2.93c-3.058-3.963-7.841-6.527-13.233-6.527c-4.799 0-9.113 2.032-12.162 5.27c-1.732 1.507-3.272 2.978-4.252 4.188l-.656.801c-3.06 3.731-6.869 8.373-6.841 16.25c.027 7.315 5.984 13.27 13.278 13.27a13.14 13.14 0 0 0 10.467-5.159a13.137 13.137 0 0 0 10.47 5.159c7.291 0 13.247-5.954 13.275-13.27c.028-7.877-3.782-12.519-6.841-16.25z"/><circle cx="50.703" cy="26.877" r="11.175" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paypal(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m33.469 70.372l2.732-11.775c.365-1.647 2.06-3.01 3.779-3.01h2.267c9.698-.01 17.22-1.993 22.626-5.965c5.398-3.963 8.098-9.217 8.085-15.714c.013-2.854-.481-5.241-1.512-7.148c-.985-1.923-2.463-3.484-4.349-4.663c-1.956-1.207-4.231-2.039-6.883-2.54c-2.668-.479-5.784-.714-9.359-.746l-16.554.013c-1.72.02-3.406 1.369-3.803 3.022L19.335 70.372c-.384 1.66.69 3.032 2.404 3.032h7.931c1.714 0 3.413-1.336 3.799-3.032m6.486-28.161l2.387-10.287c.364-1.657 2.073-3.031 3.787-3.016l2.595-.016c3.031 0 5.335.519 6.89 1.551c1.532 1.004 2.311 2.589 2.284 4.722c.045 3.211-1.172 5.696-3.575 7.45c-2.391 1.748-5.705 2.624-10.006 2.624h-1.944c-1.714 0-2.796-1.359-2.418-3.028"/><path fill="currentColor" d="M79.23 34.544c-.986-1.923-2.463-3.484-4.349-4.663c-.142-.087-.289-.168-.434-.252c.055.101.115.198.167.301c1.031 1.906 1.525 4.295 1.512 7.148c.013 6.497-2.687 11.751-8.085 15.714c-5.406 3.971-12.928 5.955-22.627 5.964H43.15c-1.719 0-3.414 1.363-3.779 3.01l-2.732 11.776c-.385 1.696-2.084 3.032-3.799 3.032h-5.356l-.364 1.582c-.384 1.66.691 3.033 2.404 3.033h7.93c1.715 0 3.414-1.337 3.799-3.033l2.733-11.775c.365-1.646 2.06-3.01 3.779-3.01h2.267c9.698-.01 17.22-1.993 22.626-5.965c5.398-3.963 8.098-9.217 8.084-15.713c.014-2.853-.481-5.242-1.512-7.149"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pencil(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M88.79 29.297L70.702 11.209a2.692 2.692 0 0 0-3.808 0L17.389 60.713l.109.109l-.171-.046l-6.772 25.272l.004.001a2.681 2.681 0 0 0 .642 2.748a2.684 2.684 0 0 0 3.033.531l.002.009l25.027-6.706l-.016-.059l.038.038L88.79 33.105a2.692 2.692 0 0 0 0-3.808m-69.998 51.85l4.022-15.009l10.988 10.988z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Photo(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M93.194 18a4.472 4.472 0 0 0-4.472-4.472c-.228 0-.447.034-.667.067V13.5H11.25v.028A4.472 4.472 0 0 0 6.778 18v64.5h.05c.252 2.231 2.123 3.972 4.421 3.972v.028h76.805v-.095c.219.033.438.067.667.067c2.299 0 4.17-1.74 4.422-3.972h.078V18zm-9.929 58.543H72.404a1.88 1.88 0 0 0-.166-.442l.018-.01l-22.719-39.35l-.009.005a3.073 3.073 0 0 0-2.764-1.74a3.08 3.08 0 0 0-2.807 1.821L28.838 63.013l-3.702-6.411l-.005.003a1.62 1.62 0 0 0-1.457-.918c-.659 0-1.224.395-1.479.958l-5.46 9.457V23.485h66.53z"/><circle cx="68.122" cy="38.584" r="10.1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Play(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M76.982 50c0-.847-.474-1.575-1.167-1.957L26.541 19.595a2.23 2.23 0 0 0-1.279-.404a2.244 2.244 0 0 0-2.244 2.243c0 .087.016.169.026.253h-.026v57.131h.026a2.235 2.235 0 0 0 2.218 1.99a2.22 2.22 0 0 0 1.117-.308l.02.035L75.875 51.97l-.02-.035A2.233 2.233 0 0 0 76.982 50"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayCircle(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M50 12.5c-20.712 0-37.5 16.793-37.5 37.502C12.5 70.712 29.288 87.5 50 87.5s37.5-16.788 37.5-37.498C87.5 29.293 70.712 12.5 50 12.5m0 65.061c-15.199 0-27.56-12.362-27.56-27.559C22.44 34.807 34.802 22.44 50 22.44c15.198 0 27.56 12.367 27.56 27.562c0 15.196-12.362 27.559-27.56 27.559"/><path fill="currentColor" d="m66.352 49.097l.006-.01l-23.367-13.491l-.006.01a1.027 1.027 0 0 0-.521-.157a1.06 1.06 0 0 0-1.059 1.06c0 .043.019.079.024.12h-.024V63.61h.024c.062.526.493.94 1.035.94c.194 0 .365-.066.521-.157l.016.027l23.367-13.49l-.016-.027c.316-.183.538-.511.538-.903s-.222-.719-.538-.903"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayVideo(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M85.527 80.647a4.971 4.971 0 0 0 4.973-4.974V24.327a4.971 4.971 0 0 0-4.973-4.974H14.474A4.972 4.972 0 0 0 9.5 24.327v51.346a4.972 4.972 0 0 0 4.974 4.974zm-4.974-9.948H19.446V29.301h61.107z"/><path fill="currentColor" d="m64.819 50.288l-11.98 6.913l-11.974 6.917V36.462l11.974 6.918z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plus(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M84.437 39.721H60.273V15.563a1.814 1.814 0 0 0-1.812-1.813H41.536a1.813 1.813 0 0 0-1.812 1.813l-.001 24.16l-24.159-.001c-.961 0-1.812.851-1.813 1.813V58.46a1.81 1.81 0 0 0 1.813 1.812h24.16v24.165a1.814 1.814 0 0 0 1.813 1.813H58.46a1.813 1.813 0 0 0 1.813-1.813V60.273l24.163-.001a1.81 1.81 0 0 0 1.813-1.813l.001-16.925a1.813 1.813 0 0 0-1.813-1.813"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pound(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M71.586 72.845c0-.205-.068-.388-.162-.555l.016-.003l-2.974-7.289l-.006-.014l-.005-.013l-.002.002a1.138 1.138 0 0 0-1.891-.312l-.005-.01c-1.046 1.448-3.461 2.897-6.922 2.897c-3.944 0-5.795-2.656-10.464-2.656c-1.207 0-2.736.241-3.944.644c2.897-1.771 5.232-5.151 5.232-8.934c0-.483-.081-1.047-.081-1.529h9.015v-.007a1.153 1.153 0 0 0 1.15-1.153V49.6h-.023a1.142 1.142 0 0 0-1.128-1.043v-.006h-11.59c-1.851-2.897-3.944-5.393-3.944-9.497c0-4.508 3.703-7.486 8.451-7.486c3.691 0 6.9 2.322 8.196 5.684a1.151 1.151 0 0 0 1.514.63l.006-.003l7.132-4.268l-.008-.008c.375-.19.638-.569.638-1.018a1.13 1.13 0 0 0-.193-.602l.022-.013c-3.381-6.6-9.739-9.739-18.674-9.739c-9.498 0-19.559 6.439-19.559 16.662c0 3.943 1.69 6.922 3.703 9.658h-5.474v.014c-.015 0-.028-.009-.044-.009c-.637 0-1.153.516-1.153 1.153v4.312h.022a1.143 1.143 0 0 0 1.131 1.044c.016 0 .028-.008.044-.009v.014h9.659c.563 1.368.966 2.816.966 4.346c0 4.062-3.627 7.57-8.616 9.983l-.022.011l-.136.068l.009.018a1.135 1.135 0 0 0-.563.968c0 .155.033.302.088.438l-.002.001l.003.007l.004.008l2.759 5.759c.006.015.014.028.021.043l.05.104l.017-.005c.201.34.555.579.978.579c.179 0 .343-.049.495-.121l.004.009c3.22-1.771 6.52-3.059 9.337-3.059c4.91 0 7.566 3.541 14.569 3.541c5.876 0 9.659-1.771 11.913-4.024l-.005-.01c.265-.211.446-.523.446-.889"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Power(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m76.587 22.655l-.043.043a2.171 2.171 0 0 0-1.484-.59c-.553 0-1.052.212-1.437.551l-.001-.004l-.09.08c-.006.006-.013.01-.018.016l-4.123 3.632c-.032.029-.067.055-.098.086l-.008.006l.001.001a2.175 2.175 0 0 0 .405 3.364c4.863 4.973 7.869 11.77 7.869 19.257c0 15.196-12.362 27.559-27.56 27.559c-15.199 0-27.561-12.362-27.561-27.559c0-7.561 3.062-14.42 8.01-19.406l-.048-.048c.464-.4.765-.986.765-1.647c0-.591-.237-1.126-.619-1.52l.001-.001l-.008-.006c-.031-.031-.066-.057-.098-.086l-4.123-3.632c-.006-.006-.013-.01-.018-.016l-.09-.08l-.001.004a2.166 2.166 0 0 0-1.437-.551c-.809 0-1.508.445-1.885 1.1C16.458 29.94 12.5 39.054 12.5 49.097c0 20.71 16.788 37.498 37.5 37.498s37.5-16.788 37.5-37.498c0-10.317-4.169-19.662-10.913-26.442"/><path fill="currentColor" d="M47.203 62.733h5.593a2.184 2.184 0 0 0 2.184-2.184V15.594h-.001l.001-.005a2.184 2.184 0 0 0-2.184-2.184h-5.592a2.185 2.185 0 0 0-2.18 2.187l-.001.002v44.955c0 1.206.976 2.182 2.18 2.184"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Previous(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M87.06 26.291c-.483 0-.918.2-1.23.521L57.41 43.22V28.009a1.718 1.718 0 0 0-2.948-1.197L21.144 46.049V22.667h-.013c.002-.039.012-.075.012-.114a2.537 2.537 0 0 0-2.532-2.536h-4.852v.024a2.537 2.537 0 0 0-2.537 2.537c0 .045.011.086.013.131v54.967h.024a2.529 2.529 0 0 0 2.513 2.296c.03 0 .058-.008.087-.009h4.569c.071.006.14.021.213.021c1.4 0 2.537-1.136 2.537-2.537c0-.117-.019-.23-.035-.343V53.962l33.341 19.25a1.715 1.715 0 0 0 2.92-1.172h.005V56.791l28.443 16.422a1.715 1.715 0 0 0 2.92-1.172h.005V28.009a1.717 1.717 0 0 0-1.717-1.718"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PriceTag(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M81.465 66.127c-.022-.038-.052-.069-.076-.105l.002-.001l-26.904-46.599l-.014.008a2.08 2.08 0 0 0-.929-.801l.038-.022l-11.01-6.357v.044a2.082 2.082 0 0 0-1.846-.107l-.012-.021l-.206.119c-.004.003-.009.003-.013.006c-.004.002-.007.006-.012.008L19.3 24.529a2.078 2.078 0 0 0-1.037 1.764l-.007-.004v12.714l.021-.012c-.047.427.03.872.261 1.273c.014.024.032.042.047.065l26.816 46.446c.026.055.044.112.075.166a2.085 2.085 0 0 0 2.772.801l.005.008l32.465-18.743l-.014-.025a2.092 2.092 0 0 0 .761-2.855M36.366 29.361a3.865 3.865 0 1 1 1.415-5.279a3.864 3.864 0 0 1-1.415 5.279"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PricetagMultiple(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M88.286 67.965c-.022-.038-.052-.069-.076-.105l.002-.001L61.307 21.26l-.015.008a2.08 2.08 0 0 0-.929-.801l.038-.022l-7.986-4.612l8.871 15.365l20.178 34.949l-.002.001c.024.036.054.067.076.105a2.091 2.091 0 0 1-.765 2.855l.014.025l-29.861 17.24l1.293 2.239c.026.055.044.112.075.166a2.085 2.085 0 0 0 2.772.801l.005.008l32.465-18.743l-.014-.025a2.09 2.09 0 0 0 .764-2.854"/><path fill="currentColor" d="M73.88 67.143a2.091 2.091 0 0 0 .765-2.855c-.022-.038-.052-.069-.076-.105l.002-.001l-26.905-46.599l-.015.008a2.08 2.08 0 0 0-.929-.801l.038-.022l-11.01-6.357v.044a2.082 2.082 0 0 0-1.846-.107l-.012-.021l-.206.119c-.004.003-.009.003-.013.006s-.007.006-.012.008L12.478 22.69a2.078 2.078 0 0 0-1.037 1.764l-.007-.004v12.714l.021-.012c-.047.427.03.872.261 1.273c.014.024.032.042.047.065l26.815 46.446c.026.055.044.112.075.166a2.085 2.085 0 0 0 2.772.801l.005.008l32.465-18.743zM29.545 27.522a3.865 3.865 0 1 1 1.415-5.279a3.862 3.862 0 0 1-1.415 5.279"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Print(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M90.443 35.141a1.825 1.825 0 0 0-1.818-1.708H75.511V11.857c0-1.012-.819-1.83-1.83-1.83H26.319c-1.011 0-1.83.818-1.83 1.83v21.576H11.375c-.969 0-1.754.755-1.818 1.708h-.012V71.91a1.83 1.83 0 0 0 1.83 1.829h13.114V58.425H75.51v15.314h13.114a1.83 1.83 0 0 0 1.83-1.829V35.141zm-19.919 6.49H29.476V16.844c0-1.012.819-1.83 1.83-1.83h37.387c1.011 0 1.83.818 1.83 1.83v24.787z"/><path fill="currentColor" d="M29.602 88.143c0 1.012.819 1.83 1.83 1.83h37.136c1.011 0 1.83-.818 1.83-1.83v-24.64H29.602z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Prohibited(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M50 12.5c-20.712 0-37.5 16.793-37.5 37.502C12.5 70.712 29.288 87.5 50 87.5s37.5-16.788 37.5-37.498C87.5 29.293 70.712 12.5 50 12.5m0 65.061c-15.199 0-27.561-12.362-27.561-27.559a27.435 27.435 0 0 1 6.4-17.636l38.795 38.795A27.431 27.431 0 0 1 50 77.561m21.161-9.926L32.366 28.839A27.431 27.431 0 0 1 50 22.439c15.198 0 27.56 12.367 27.56 27.562a27.43 27.43 0 0 1-6.399 17.634"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ProjectionScreen(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M88.158 16.509c0-.986-.798-1.785-1.785-1.786H13.628c-.987 0-1.786.799-1.786 1.785v6.386c0 .986.799 1.785 1.786 1.785h72.746c.987 0 1.785-.799 1.785-1.785zm-6.132 13.15H17.974c-.986.001-1.785.8-1.785 1.786l.001.006l-.001.003v37.154c0 2.201 1.857 3.982 4.152 3.982h27.17v5.373a4.067 4.067 0 0 0-1.599 3.222a4.09 4.09 0 1 0 8.178 0a4.065 4.065 0 0 0-1.601-3.223V72.59h27.17c2.294 0 4.152-1.781 4.152-3.982V31.412h-.003a1.786 1.786 0 0 0-1.782-1.753"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Puzzle(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M78.589 48.774V35.887h-.001a3.967 3.967 0 0 0-3.87-3.953v-.038H62.666c-.059.002-.115.018-.175.018s-.116-.016-.175-.018h-.203v-.021c-2.8-.198-5.016-2.508-5.016-5.356c0-1.236.432-2.361 1.131-3.271h-.106a8.168 8.168 0 0 0 1.098-4.092a8.221 8.221 0 0 0-8.223-8.221a8.22 8.22 0 0 0-8.221 8.221c0 1.491.403 2.886 1.098 4.092h-.091a5.35 5.35 0 0 1 1.117 3.256c0 2.847-2.216 5.157-5.017 5.354v.038H25.361v.02a3.972 3.972 0 0 0-3.955 3.971c0 .058.015.112.017.17v13.421h.043a2.245 2.245 0 0 0 2.243 2.244c.226 0 .44-.045.645-.109a9.727 9.727 0 0 1 3.883-.808c5.393 0 9.764 4.371 9.764 9.762c0 5.392-4.371 9.762-9.764 9.762a9.731 9.731 0 0 1-3.788-.763a2.23 2.23 0 0 0-.769-.144a2.24 2.24 0 0 0-2.242 2.244c0 .127.017.25.038.372h-.053v12.884c-.003.058-.017.112-.017.17a3.975 3.975 0 0 0 3.973 3.973c.109 0 .212-.023.319-.032H38.89a2.245 2.245 0 0 0 2.244-2.246c0-.248-.05-.481-.124-.704l-.004-.014c-.017-.051-.027-.106-.048-.155h.008a9.77 9.77 0 0 1-.397-2.741a9.76 9.76 0 0 1 9.762-9.762a9.761 9.761 0 0 1 9.762 9.762a9.76 9.76 0 0 1-.397 2.741h.025a2.245 2.245 0 0 0 2.066 3.123c.013 0 .025-.004.038-.004h12.792a3.974 3.974 0 0 0 3.969-3.933h.005V72.038h-.046a2.234 2.234 0 0 0-2.216-1.964v-.099a9.747 9.747 0 0 1-2.567.353c-5.392 0-9.763-4.37-9.763-9.762c0-5.391 4.371-9.762 9.763-9.762c.89 0 1.748.129 2.567.353v-.094l.025.002c1.01 0 1.853-.67 2.134-1.588h.103v-.605c0-.017.005-.032.005-.049c-.002-.017-.006-.032-.007-.049"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Quote(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.572 43.713c-.809 0-1.617.101-2.02.303c.899-3.602 4.307-7.642 7.513-9.387c.007-.003.012-.008.018-.011c.024-.013.048-.031.071-.044l-.003-.002c.355-.19.605-.552.605-.983a1.11 1.11 0 0 0-.505-.916l.025-.024l-4.196-2.65l-.013.011a1.101 1.101 0 0 0-.668-.242c-.184 0-.35.054-.504.132l-.021-.019c-6.26 4.442-10.401 11.206-10.401 18.78c0 6.562 4.241 10.297 8.985 10.297c4.342 0 7.978-3.634 7.978-7.977c.001-4.34-3.027-7.268-6.864-7.268m20.547 0c-.809 0-1.617.101-2.02.303c.899-3.602 4.307-7.642 7.513-9.387c.007-.003.012-.008.018-.011c.024-.013.048-.031.071-.044l-.003-.002c.355-.19.605-.552.605-.983a1.11 1.11 0 0 0-.505-.916l.025-.024l-4.196-2.65l-.013.011a1.101 1.101 0 0 0-.668-.242c-.184 0-.35.054-.504.132l-.021-.019c-6.26 4.442-10.401 11.206-10.401 18.78c0 6.562 4.241 10.297 8.985 10.297c4.342 0 7.978-3.634 7.978-7.977c.001-4.34-3.027-7.268-6.864-7.268m24.875-2.672c-4.342 0-7.978 3.634-7.978 7.977c0 4.341 3.028 7.269 6.865 7.269c.809 0 1.617-.101 2.02-.303c-.899 3.602-4.307 7.642-7.513 9.387c-.007.003-.012.008-.018.011c-.024.013-.048.031-.071.044l.003.002c-.355.19-.605.552-.605.983c0 .388.208.713.505.916l-.025.024l4.196 2.65l.013-.011c.189.143.413.242.668.242c.184 0 .35-.054.504-.132l.021.019c6.26-4.443 10.401-11.206 10.401-18.78c-.001-6.563-4.242-10.298-8.986-10.298m20.547 0c-4.342 0-7.978 3.634-7.978 7.977c0 4.341 3.028 7.269 6.865 7.269c.809 0 1.617-.101 2.02-.303c-.899 3.602-4.307 7.642-7.513 9.387c-.007.003-.012.008-.018.011c-.024.013-.048.031-.071.044l.003.002c-.355.19-.605.552-.605.983c0 .388.208.713.505.916l-.025.024l4.196 2.65l.013-.011c.189.143.413.242.668.242c.184 0 .35-.054.504-.132l.021.019c6.26-4.443 10.401-11.206 10.401-18.78c-.001-6.563-4.242-10.298-8.986-10.298"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Record(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M50 12.908c-20.485 0-37.092 16.606-37.092 37.092c0 20.485 16.606 37.092 37.092 37.092c20.485 0 37.092-16.606 37.092-37.092c0-20.485-16.607-37.092-37.092-37.092m0 49.283c-6.733 0-12.191-5.458-12.191-12.191c0-6.733 5.458-12.191 12.191-12.191c6.733 0 12.191 5.458 12.191 12.191c0 6.733-5.458 12.191-12.191 12.191"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Refresh(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M84.539 21.586a1.515 1.515 0 0 0-2.393-1.222l-5.944 4.261l-.468.337c-6.405-6.392-15.196-10.389-24.937-10.389c-19.535 0-35.427 15.894-35.427 35.428s15.893 35.428 35.427 35.428a35.42 35.42 0 0 0 29.374-15.618a1.77 1.77 0 0 0-.475-2.462l-8.863-6.151a1.9 1.9 0 0 0-2.628.512c-3.918 5.792-10.41 9.25-17.375 9.25c-11.558 0-20.962-9.402-20.962-20.957s9.404-20.957 20.962-20.957c4.878 0 9.352 1.696 12.914 4.5l-1.001.72l-5.948 4.26a1.513 1.513 0 0 0 .397 2.656l25.446 8.669c.461.161.966.083 1.368-.203c.399-.29.629-.747.627-1.231z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Results(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M92.11 21.929c0-.973-.789-1.761-1.761-1.761l-.006.001v-.005H44.465v.014a1.76 1.76 0 0 0-1.751 1.752h-.001v6.342c-.001.03-.009.057-.009.087c0 .972.788 1.76 1.761 1.761h45.878l.006.001c.973 0 1.761-.789 1.761-1.761zm0 14.901c0-.973-.789-1.761-1.761-1.761l-.006.001v-.005H44.465v.014a1.76 1.76 0 0 0-1.751 1.752h-.001v6.343c-.001.03-.009.057-.009.087c0 .972.788 1.76 1.761 1.761h45.878v-.001l.006.001c.973 0 1.761-.789 1.761-1.761zM32.748 21.925a1.76 1.76 0 0 0-1.761-1.761c-.043 0-.084.01-.126.013H9.777c-.042-.003-.083-.013-.126-.013c-.973 0-1.761.789-1.761 1.761V43.26c0 .972.788 1.761 1.761 1.761h21.336c.973 0 1.761-.789 1.761-1.761zM92.11 56.744c0-.973-.789-1.761-1.761-1.761l-.006.001v-.005H44.465v.014a1.76 1.76 0 0 0-1.751 1.752h-.001v6.343c-.001.03-.009.057-.009.087c0 .972.788 1.761 1.761 1.761h45.878v-.001l.006.001c.973 0 1.761-.789 1.761-1.761zm0 14.901c0-.973-.789-1.761-1.761-1.761l-.006.001v-.005H44.465v.014a1.76 1.76 0 0 0-1.751 1.752h-.001v6.342c-.001.03-.009.057-.009.087c0 .972.788 1.761 1.761 1.761h45.878v-.001l.006.001c.973 0 1.761-.789 1.761-1.761zM32.748 56.74c0-.973-.788-1.762-1.761-1.762c-.043 0-.084.01-.126.013H9.777c-.042-.003-.083-.013-.126-.013c-.973 0-1.761.789-1.761 1.762v21.335c0 .972.789 1.761 1.761 1.761h21.336c.973 0 1.761-.789 1.761-1.761z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ResultsDemographics(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M27.953 46.506a21.02 21.02 0 0 1-2.117-9.192c0-1.743.252-3.534.768-5.468c.231-.87.521-1.702.847-2.509c-1.251-.683-2.626-1.103-4.101-1.103c-5.47 0-9.898 5.153-9.898 11.517c0 4.452 2.176 8.305 5.354 10.222L5.391 56.217c-.836.393-1.387 1.337-1.387 2.392v10.588c0 1.419.991 2.569 2.21 2.569h7.929v-11.11c0-3.237 1.802-6.172 4.599-7.481l10.262-4.779a20.2 20.2 0 0 1-1.051-1.89m32.184-11.705h34.092V34.8l.006.001c.973 0 1.761-.789 1.761-1.761v-6.431c0-.973-.789-1.761-1.761-1.761l-.006.001v-.005H56.133a20.013 20.013 0 0 1 3.526 7.435c.215.889.371 1.72.478 2.522m35.859 31.635c0-.973-.789-1.761-1.761-1.761l-.006.001v-.005H72.007v9.089c0 .293-.016.582-.045.867h22.267v-.001l.006.001c.973 0 1.761-.789 1.761-1.761zm-1.761-21.674l-.006.001v-.005H58.944c-.159.419-.327.836-.514 1.249a20.258 20.258 0 0 1-1.224 2.297l10.288 4.908a7.67 7.67 0 0 1 2.078 1.503h24.657v-.001l.006.001c.973 0 1.761-.789 1.761-1.761v-6.431c0-.973-.789-1.761-1.761-1.761"/><path fill="currentColor" d="m65.323 57.702l-11.551-5.51l-4.885-2.33c2.134-1.344 3.866-3.418 5-5.917a16.045 16.045 0 0 0 1.435-6.631c0-1.348-.213-2.627-.512-3.863c-1.453-5.983-6.126-10.392-11.736-10.392c-5.504 0-10.106 4.251-11.648 10.065c-.356 1.333-.602 2.72-.602 4.189c0 2.552.596 4.93 1.609 7c1.171 2.4 2.906 4.379 5.018 5.651l-4.678 2.178l-11.926 5.554c-1.037.485-1.717 1.654-1.717 2.959V73.76c0 1.756 1.224 3.181 2.735 3.181h42.417c1.511 0 2.735-1.424 2.735-3.181V60.656c.002-1.301-.668-2.458-1.694-2.954"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rewind(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M84.627 26.291c-.483 0-.918.2-1.23.521L54.978 43.22V28.009a1.718 1.718 0 0 0-2.948-1.197L14.485 48.489l.031.054a1.71 1.71 0 0 0-.862 1.481c0 .702.422 1.303 1.026 1.57l-.017.03l37.391 21.588a1.715 1.715 0 0 0 2.92-1.172h.005V56.791l28.443 16.422a1.715 1.715 0 0 0 2.92-1.172h.005V28.009a1.72 1.72 0 0 0-1.72-1.718"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RewindTen(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M78.941 49.005a1.51 1.51 0 0 0 .603-1.242l-.609-26.821a1.517 1.517 0 0 0-.85-1.322a1.518 1.518 0 0 0-1.566.145l-5.861 4.375l-.461.345c-6.527-6.266-15.394-10.094-25.133-9.906c-19.532.376-35.115 16.574-34.739 36.103c.376 19.53 16.572 35.115 36.104 34.739a36.53 36.53 0 0 0 2.1-.105s.469-.024 1.008-.271l.019-.01a2.8 2.8 0 0 0 .719-.495c.033-.031.068-.059.099-.092l.019-.017c.014-.015.019-.03.031-.045a2.58 2.58 0 0 0 .658-1.766c0-.027-.008-.052-.009-.079l-.175-9.105h-.008c0-.028.007-.054.007-.082a2.597 2.597 0 0 0-2.647-2.548v-.003a21.14 21.14 0 0 1-2.067.152c-11.556.223-21.14-8.997-21.362-20.55c-.223-11.554 8.999-21.135 20.555-21.358c4.878-.094 9.383 1.517 12.998 4.251l-.986.738l-5.866 4.375a1.517 1.517 0 0 0-.583 1.46c.092.556.49 1.019 1.031 1.188l25.609 8.177c.463.15.966.063 1.362-.231"/><path fill="currentColor" d="m67.109 59.835l-5.941 6.204l2.384 2.407l3.139-3.322l.269 13.963l4.096-.079l-.37-19.242zm14.416-.566c-5.51.106-7.867 5.115-7.771 10.077c.096 4.962 2.644 9.906 8.154 9.8c5.51-.106 7.866-5.145 7.771-10.107c-.096-4.962-2.644-9.876-8.154-9.77m.313 16.242c-2.741.053-3.832-2.697-3.901-6.245s.917-6.309 3.658-6.362c2.77-.053 3.832 2.668 3.9 6.216c.069 3.548-.916 6.338-3.657 6.391"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rss(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M26.258 64.949a8.78 8.78 0 0 0-8.78 8.784a8.782 8.782 0 1 0 17.564 0a8.78 8.78 0 0 0-8.784-8.784"/><path fill="currentColor" d="M23.536 40.801c-.046 0-.09.006-.135.007v-.007h-3.464v.039a3.437 3.437 0 0 0-3.056 3.344h-.007v6.159h.041a3.43 3.43 0 0 0 3.021 3.002v.039H23.4v-.048c.045.001.09.007.135.007c12.772 0 23.173 10.321 23.311 23.061h-.033v3.464h.039a3.437 3.437 0 0 0 3.344 3.056v.007h6.158v-.041a3.43 3.43 0 0 0 3.002-3.021h.039v-3.464h-.006c-.137-19.657-16.166-35.604-35.853-35.604"/><path fill="currentColor" d="M83.119 76.403C82.98 43.664 56.308 17.07 23.536 17.07c-.046 0-.09.006-.135.007v-.007h-3.464v.039a3.437 3.437 0 0 0-3.056 3.344h-.007v6.159h.041a3.429 3.429 0 0 0 3.021 3.002v.039H23.4v-.048c.045.001.09.007.135.007c25.857 0 46.902 20.967 47.041 46.792h-.035v3.464h.039a3.437 3.437 0 0 0 3.344 3.056v.007h6.159v-.041a3.43 3.43 0 0 0 3.002-3.021h.039v-3.464h-.005z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SafetyCone(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M97.017 82.01v-4.62h-.002l.002-.018a2.437 2.437 0 0 0-2.42-2.417v-.003h-7.512l-33.037-57.22l-.013.008c-.732-1.503-2.26-2.547-4.044-2.547a4.505 4.505 0 0 0-4.108 2.665L12.92 74.951H5.386a2.439 2.439 0 0 0-2.44 2.438v4.863h.012c-.002.039-.012.075-.012.114a2.44 2.44 0 0 0 2.44 2.439h89.21v-.004c.013 0 .024.004.037.003a2.44 2.44 0 0 0 2.42-2.458c0-.114-.02-.224-.036-.336M47.513 25.159h4.977l5.748 9.957H41.765zM36.017 45.072h27.971l5.748 9.957H30.268zM18.766 74.951l5.748-9.957H75.49l5.748 9.957z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Save(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M52.872 17.026h10.352v21.471H52.872z"/><path fill="currentColor" d="m86.217 19.195l.006-.006l-5.211-5.212a2.375 2.375 0 0 0-1.76-.785c-.014 0-.026.004-.04.004h-9.087v.024c-.286.029-.564.128-.779.344a1.246 1.246 0 0 0-.343.779h-.028v25.306h-.009a2.382 2.382 0 0 1-2.378 2.3H32.135v-.012a2.373 2.373 0 0 1-2.259-2.288h-.009V14.343h-.048a1.38 1.38 0 0 0-.367-.735a1.368 1.368 0 0 0-1.119-.387v-.028H15.382a2.385 2.385 0 0 0-2.386 2.388v68.841a2.384 2.384 0 0 0 2.387 2.386h7.57V53.801h.009a2.385 2.385 0 0 1 1.92-2.255v-.013h.132c.108-.015.215-.032.326-.032h49.387v.009a2.375 2.375 0 0 1 2.292 2.291h.009v33.007h7.581a2.386 2.386 0 0 0 2.387-2.386V21.029c.001-.025.007-.048.007-.073a2.37 2.37 0 0 0-.786-1.761"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Share(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M73.213 38.25a11.72 11.72 0 0 0-4.189.782L38.498 21.408c.017-.259.039-.517.039-.78c0-6.489-5.261-11.75-11.75-11.75s-11.75 5.261-11.75 11.75s5.261 11.75 11.75 11.75c2.563 0 4.927-.83 6.858-2.223l28.343 16.364c-.341 1.1-.525 2.27-.525 3.482c0 1.232.191 2.418.543 3.534L33.693 69.881a11.68 11.68 0 0 0-6.906-2.258c-6.489 0-11.75 5.261-11.75 11.75s5.261 11.75 11.75 11.75s11.75-5.261 11.75-11.75c0-.243-.022-.48-.036-.72l30.59-17.661a11.7 11.7 0 0 0 4.122.758c6.489 0 11.75-5.261 11.75-11.75s-5.261-11.75-11.75-11.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SheriffBadge(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M82.066 62.012c-.861 0-1.663.222-2.385.583l-6.813-11.801l6.755-11.701c.438.115.89.196 1.364.196a5.389 5.389 0 0 0 5.382-5.388a5.379 5.379 0 0 0-5.382-5.374a5.368 5.368 0 0 0-4.449 2.349h-15.17l-7.042-12.197a5.348 5.348 0 0 0 .873-2.92a5.388 5.388 0 0 0-5.389-5.385a5.387 5.387 0 0 0-5.382 5.387a5.35 5.35 0 0 0 .835 2.848L38.18 30.876H22.65c-.909-1.686-2.668-2.845-4.716-2.842a5.378 5.378 0 0 0-5.389 5.376a5.388 5.388 0 0 0 5.389 5.386c.861 0 1.663-.223 2.385-.583l6.814 11.802l-6.756 11.702a5.343 5.343 0 0 0-1.364-.196a5.38 5.38 0 1 0 0 10.762a5.377 5.377 0 0 0 4.451-2.35h15.169l6.815 11.805a5.327 5.327 0 0 0-.647 2.505a5.387 5.387 0 0 0 5.389 5.386a5.389 5.389 0 0 0 5.382-5.388a5.335 5.335 0 0 0-.61-2.432l6.857-11.876h15.532c.909 1.685 2.668 2.844 4.717 2.844a5.383 5.383 0 0 0 5.389-5.378a5.393 5.393 0 0 0-5.391-5.387"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shield(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m50.027 10.459l-.018-.032l-33.606 19.404l.076.132v22.893h.014c.286 19.111 14.859 34.755 33.519 36.718c18.66-1.962 33.234-17.606 33.519-36.718V29.953l.066-.114zm-.015 69.097V51.677H26.435V35.651L50.012 22.04v29.637h23.563v1.179h.017c-.278 13.593-10.439 24.798-23.58 26.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBag(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M93.031 38.642c-.213-1.934-1.832-3.443-3.823-3.443h-6.1C76.802 23.367 64.344 15.309 50 15.309c-14.345 0-26.802 8.058-33.109 19.891h-6.099c-1.99 0-3.61 1.51-3.823 3.443h-.043v42.185a3.865 3.865 0 0 0 3.865 3.864h78.418a3.863 3.863 0 0 0 3.865-3.864V38.642zm-64.213-3.443c5.059-6.076 12.675-9.951 21.182-9.951c8.506 0 16.122 3.875 21.181 9.951z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingCart(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="32.961" cy="81.02" r="5.889" fill="currentColor"/><circle cx="74.309" cy="81.02" r="5.889" fill="currentColor"/><path fill="currentColor" d="M73.513 55.216a3.443 3.443 0 0 0 3.11-1.978l.003.002l14.589-25.27l-.018-.01c.214-.449.344-.946.344-1.477a3.451 3.451 0 0 0-3.452-3.452h-53.59v-6.488a3.451 3.451 0 0 0-3.452-3.452H12.186v.028c-.092-.008-.181-.028-.275-.028a3.44 3.44 0 0 0-3.41 3.035h-.042v3.452a3.451 3.451 0 0 0 3.452 3.452c.094 0 .183-.021.275-.028v.028h12.373v43.733h.001a3.45 3.45 0 0 0 3.451 3.442h59.256v-.085a3.433 3.433 0 0 0 3.185-2.922h.049v-.49l.002-.018l-.002-.018V63.67l.002-.018a3.44 3.44 0 0 0-3.236-3.429v-.022H34.498v-4.988h39.015z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shuffle(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.401 62.222H8.022a2.536 2.536 0 0 0-3.013 2h-.047v11h.076a2.527 2.527 0 0 0 2.078 1.906v.094h17v-.061c1.739-.021 3.23-1.008 4.002-2.443L38.63 64.206L28.024 53.599zm75.25-31.446l-18.614-13.26a1.255 1.255 0 0 0-1.319-.094a1.278 1.278 0 0 0-.688 1.146l.002 5.599H59.116v.026a4.577 4.577 0 0 0-2.781 1.138l-.021-.021L45.722 35.9l10.607 10.607l7.341-7.341h10.368l.002 6.005c0 .481.271.924.7 1.146c.429.222.946.183 1.34-.099l18.576-13.346c.338-.241.461-.631.461-1.046v-.003c-.001-.417-.126-.806-.466-1.047"/><path fill="currentColor" d="M94.578 67.126L76.002 53.781a1.289 1.289 0 0 0-2.039 1.047l-.002 6.005H63.592L28.04 25.281c-.772-1.435-2.263-2.421-4.001-2.442v-.061h-17v.094a2.527 2.527 0 0 0-2.078 1.906h-.077v11h.049a2.535 2.535 0 0 0 3.013 2h11.378l36.913 36.913l.021-.021a4.583 4.583 0 0 0 2.782 1.138v.026h14.916l-.002 5.599c0 .484.26.928.688 1.146a1.25 1.25 0 0 0 1.319-.093l18.614-13.26c.34-.242.465-.631.465-1.049v-.004c-.001-.415-.124-.806-.462-1.047"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Skull(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M87.255 72.316c-2.248-2.246-5.832-2.318-8.118-.208l-9-9c2.459-3.77 3.894-8.27 3.894-13.107s-1.435-9.337-3.894-13.107l8.949-8.949c2.287 2.111 5.87 2.042 8.118-.207c2.306-2.3 2.329-6.017.054-8.293c-1.371-1.371-3.264-1.9-5.058-1.603c.33-1.82-.196-3.755-1.592-5.151c-2.274-2.275-5.985-2.252-8.291.055c-2.246 2.248-2.318 5.832-.208 8.118l-9.001 9.001a23.91 23.91 0 0 0-13.107-3.893a23.913 23.913 0 0 0-13.107 3.893l-8.95-8.95c2.11-2.286 2.04-5.869-.206-8.118c-2.305-2.304-6.018-2.328-8.293-.054c-1.372 1.372-1.901 3.266-1.604 5.059c-1.82-.33-3.756.197-5.152 1.591c-2.275 2.275-2.248 5.986.055 8.291c2.249 2.25 5.834 2.319 8.121.209l9 9c-2.46 3.77-3.894 8.27-3.894 13.107s1.434 9.337 3.894 13.107l-8.95 8.95c-2.286-2.11-5.869-2.04-8.118.206c-2.304 2.305-2.327 6.018-.054 8.293c1.372 1.372 3.266 1.901 5.059 1.604c-.33 1.82.197 3.756 1.591 5.152c2.275 2.275 5.986 2.248 8.291-.055c2.249-2.249 2.319-5.834.209-8.121l8.726-8.726v5.946a21.048 21.048 0 0 0 13.382 4.779a21.05 21.05 0 0 0 13.383-4.779V70.41l8.674 8.674c-2.111 2.287-2.043 5.871.206 8.119c2.301 2.306 6.018 2.329 8.293.054c1.371-1.371 1.9-3.265 1.603-5.058c1.82.331 3.755-.196 5.151-1.591c2.274-2.275 2.251-5.985-.056-8.292m-43.81-13.327a4.594 4.594 0 1 1 0-9.189c2.535 0 4.591 2.061 4.591 4.597s-2.057 4.592-4.591 4.592m13.111.067a4.592 4.592 0 0 1 0-9.184a4.593 4.593 0 1 1 0 9.184"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialAdobe(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M59.386 18.139L86 81.861V18.139zm-45.386 0v63.722l26.635-63.722zm24.373 50.904h12.409l5.075 12.814H66.97L50.01 41.622z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialAmazon(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M52.02 35.981c-4.333.371-8.663.795-12.852 2.047c-8.593 2.574-12.979 8.387-12.838 16.969c.128 7.734 4.601 12.865 12.285 13.684c4.953.527 9.708-.293 13.728-3.584c1.151-.941 2.23-1.971 3.249-2.875c1.665 1.838 3.229 3.666 4.906 5.387c1.063 1.09 2.088 1.047 3.231.068c2.144-1.832 4.275-3.68 6.407-5.527c1.152-1 1.302-1.889.268-3.059c-2.529-2.867-3.404-6.16-3.337-9.955a399.95 399.95 0 0 0-.161-20.285c-.171-5.197-2.768-9.066-7.554-11.184c-7.35-3.25-14.798-3.178-22.127.051c-4.949 2.184-7.977 6.076-9.16 11.377c-.388 1.742.076 2.392 1.83 2.602c2.529.303 5.055.627 7.581.957c1.393.182 2.168-.236 2.47-1.631c.899-4.154 5.005-6.209 8.898-5.631c2.513.373 4.387 1.963 4.771 4.371c.312 1.959.289 3.973.413 5.941c-.331.09-.456.145-.586.156c-.474.047-.948.08-1.422.121m-.185 19.772c-1.262 1.906-2.879 3.281-5.155 3.664c-3.688.619-6.401-1.754-6.672-5.738c-.355-5.217 2.309-8.791 7.459-9.781c2.077-.4 4.213-.496 6.422-.742c.021 4.505.473 8.777-2.054 12.597"/><path fill="currentColor" d="M75.38 70.892c-.357.106-.698.262-1.042.408c-4.914 2.096-9.966 3.748-15.213 4.783c-4.954.979-9.942 1.473-14.99 1.006c-6.205-.57-11.993-2.566-17.527-5.324c-3.522-1.754-6.962-3.672-10.447-5.502c-1.093-.574-2.095-.404-2.656.391c-.554.783-.362 1.814.616 2.584c3.468 2.729 6.845 5.594 10.478 8.086c6.773 4.645 14.24 7.418 23.981 7.414c7.914-.199 16.618-2.611 24.675-7.299c1.434-.834 2.808-1.785 4.148-2.766c1.021-.748 1.24-1.848.718-2.789c-.504-.91-1.62-1.324-2.741-.992"/><path fill="currentColor" d="M82.865 63.228c-3.83-.387-7.529.078-10.761 2.443c-.751.547-1.186 1.287-.776 2.203c.391.873 1.19.986 2.077.861c1.678-.236 3.363-.428 5.046-.639l.012.131c.612 0 1.223-.008 1.833.002c.707.012 1.084.32.921 1.08c-.179.828-.3 1.674-.55 2.479c-.481 1.555-1.103 3.068-1.519 4.637c-.131.494.031 1.324.381 1.621c.399.336 1.175.383 1.754.305c.408-.053.806-.477 1.13-.811c2.582-2.672 3.898-5.943 4.29-9.576c.385-3.57-.248-4.373-3.838-4.736"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialAndroid(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m60.623 18.128l3.36-6.136a.668.668 0 1 0-1.169-.644l-3.396 6.197a23.172 23.172 0 0 0-9.432-1.978a23.09 23.09 0 0 0-9.411 1.972l-3.398-6.184a.663.663 0 0 0-.904-.267a.66.66 0 0 0-.265.905l3.362 6.13c-6.606 3.408-11.068 9.893-11.064 17.341l43.365-.004c.002-7.446-4.452-13.92-11.048-17.332m-20.499 9.481a1.818 1.818 0 0 1 .001-3.636c1.006 0 1.819.82 1.82 1.818a1.82 1.82 0 0 1-1.821 1.818m19.746-.004a1.816 1.816 0 0 1-1.818-1.818c0-.996.814-1.816 1.818-1.82a1.829 1.829 0 0 1 1.815 1.822a1.816 1.816 0 0 1-1.815 1.816m23.139 13.433a4.813 4.813 0 1 0-9.627.002l.002 20.159a4.81 4.81 0 0 0 4.816 4.813a4.81 4.81 0 0 0 4.812-4.815zm-11.486-3.907l-43.039.008l.007 31.241a5.078 5.078 0 0 0 .546 2.267c.837 1.697 2.567 2.872 4.591 2.872h3.498l.004 10.666a4.816 4.816 0 0 0 5.781 4.717a4.82 4.82 0 0 0 3.845-4.718l-.003-10.666h6.499v10.666a4.825 4.825 0 0 0 3.848 4.715a4.816 4.816 0 0 0 5.784-4.72l-.003-10.662l3.511-.002a5.073 5.073 0 0 0 2.621-.742c1.5-.897 2.517-2.52 2.518-4.398zm-49.722-.902a4.817 4.817 0 0 0-4.814 4.814l.003 20.161a4.807 4.807 0 0 0 2.164 4.017c.761.504 1.671.8 2.651.802a4.815 4.815 0 0 0 4.813-4.819V41.042a4.817 4.817 0 0 0-4.817-4.813"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialApple(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M82.047 36.841c-4.124-5.139-9.91-8.112-15.372-8.112c-7.226 0-10.279 3.433-15.292 3.433c-5.167 0-9.096-3.428-15.346-3.428c-6.139 0-12.666 3.729-16.804 10.102c-5.823 8.962-4.837 25.824 4.605 40.193C27.215 84.167 31.725 89.945 37.622 90c5.241.048 6.724-3.343 13.835-3.383c7.112-.039 8.457 3.42 13.697 3.371c5.897-.056 10.655-6.454 14.03-11.593c2.419-3.681 3.32-5.541 5.199-9.702c-13.651-5.158-15.846-24.447-2.336-31.852"/><path d="M61.194 22.988c2.625-3.368 4.619-8.126 3.895-12.988c-4.288.296-9.304 3.023-12.232 6.58c-2.658 3.228-4.853 8.015-3.999 12.668c4.681.145 9.525-2.652 12.336-6.26"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialBehance(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M64.165 56.231c-3.279 0-4.361 2.494-4.69 3.652h8.754c-.082-.643-.339-1.773-1.157-2.597c-.697-.699-1.672-1.055-2.907-1.055m-22.503 4.658H33.37v7.067h7.839c1.398-.035 3.13-.485 3.13-3.422c0-3.543-2.403-3.645-2.677-3.645"/><path d="M79.035 14.073h-58a7.004 7.004 0 0 0-6.693 4.946v62.107a7.01 7.01 0 0 0 5.265 4.8h60.855a7.009 7.009 0 0 0 5.196-4.602V18.821a6.998 6.998 0 0 0-6.623-4.748M57.268 44.448c0-.277.226-.502.502-.502h12.465c.274 0 .501.225.501.502v3.721a.503.503 0 0 1-.501.501H57.77a.502.502 0 0 1-.502-.501zm-5.294 20.639c0 4.75-2.513 7.041-4.625 8.12c-2.301 1.179-4.686 1.271-5.364 1.271H26.162a.506.506 0 0 1-.502-.503V43.77c0-.277.228-.503.502-.503h15.721c5.58 0 9.047 3.274 9.047 8.547c0 3.266-1.591 4.833-2.991 5.579a5.07 5.07 0 0 1 2.041 1.329c2.194 2.32 2.006 6.2 1.994 6.365m23.304-.455a.499.499 0 0 1-.497.437H59.648c.109 1.659.766 2.856 1.953 3.563c1.05.623 2.211.691 2.657.691c.093 0 .151-.004.158-.006c1.737 0 2.97-.383 3.697-1.139c.697-.726.678-1.547.673-1.582c0-.138.046-.271.144-.371a.498.498 0 0 1 .362-.155h5.313c.278 0 .504.225.504.503c0 8.131-8.683 8.543-10.429 8.543c-.279 0-.438-.01-.438-.01h-.005c-3.92 0-6.987-1.143-9.091-3.395c-3.297-3.534-2.996-8.478-2.981-8.686c0-.089.014-3.019 1.495-5.992c1.391-2.773 4.352-6.081 10.612-6.081c3.587 0 6.38 1.102 8.287 3.274c3.509 4.001 2.759 10.146 2.719 10.406"/><path d="M43.29 52.132c0-2.609-1.565-2.639-1.63-2.639h-8.29v5.734h7.742c1.151 0 2.178-.236 2.178-3.095"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialBing(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<ellipse cx="50.156" cy="56.277" fill="currentColor" rx="19.383" ry="13.568"/><path fill="currentColor" d="M77.132 38.339c-7.063-4.908-15.964-7.331-24.577-7.779c-10.167-.102-20.041 1.224-29.082 7.514v-19.34h-11.32v35.044c-.129 14.675 9.747 20.961 22.465 25.78c4.621 1.14 9.741 1.709 15.36 1.709c5.526 0 10.588-.569 15.186-1.709c4.597-1.141 8.568-2.973 11.912-4.984c3.344-2.012 5.944-4.365 7.802-7.061c1.857-2.694 2.786-5.572 2.786-8.63c1.43-10.439-5.757-17.226-10.532-20.544M53.313 75.017c-12.718.398-27.624-3.441-29.07-18.361c-.804-8.291 7.453-18.73 25.735-19.374c7.517-.265 25.052 3.927 26.249 18.091c.737 8.731-10.196 19.247-22.914 19.644"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialBlogger(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M74.121 14H25.88C19.351 14 14 19.351 14 25.88v48.24C14 80.65 19.351 86 25.88 86h48.241C80.65 86 86 80.65 86 74.12V25.88C86 19.351 80.65 14 74.121 14m-2.3 43.982c-.044 7.649-6.264 13.913-13.928 13.913H41.886c-7.664 0-13.928-6.264-13.928-13.913V41.725c0-7.665 6.264-13.943 13.928-13.943h9.743c3.596.428 8.829 3.508 10.759 7.605c.531 1.149.81 1.326 1.253 4.73c.236 1.754.354 3.051 1.135 3.773c1.105 1.002 5.203.324 6.014.958l.619.486l.368.767l.133.619z"/><path d="M42.063 44.495h7.724a2.677 2.677 0 0 0 2.668-2.668a2.664 2.664 0 0 0-2.668-2.653h-7.724a2.663 2.663 0 0 0-2.667 2.653a2.676 2.676 0 0 0 2.667 2.668M57.76 55.122H42.063a2.663 2.663 0 0 0-2.667 2.653c0 1.443 1.194 2.652 2.667 2.652H57.76a2.67 2.67 0 0 0 2.653-2.652a2.66 2.66 0 0 0-2.653-2.653"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialDelicious(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M50 49.947V16.053H16v34h34v33.894h34v-34z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialDesignerNews(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m66.746 37.902l.261 15.744l-9.164-14.804l-.582-.94h-5.472v25.001h6.109l-.262-15.181l9.672 15.181h5.547V37.902z"/><path d="M79 14H21c-3.85 0-7 3.15-7 7v58c0 3.85 3.15 7 7 7h31.825L24.843 63.556V36.902h10.529c3.975 0 6.096.493 8.394 1.945c1.259.771 2.411 1.865 3.327 3.156c.377.523.703 1.071.993 1.637l2.703 2.014v-8.753h7.029l.293.474l.104.167l7.625 6.025l-.094-5.65l-.017-1.017h8.126v.01L86 46.753V21c0-3.85-3.15-7-7-7"/><path d="M40.842 38.546c-1.324-.434-2.861-.613-5-.636c-.16-.002-.304-.008-.471-.008h-8.396v25.001h9.822c7.085 0 11.771-5.023 11.771-12.595c0-3.072-.75-5.586-2.286-7.721a11.302 11.302 0 0 0-.964-1.162c-.08-.086-.161-.172-.243-.256a10.696 10.696 0 0 0-.766-.695c-.102-.084-.201-.169-.305-.248a11.864 11.864 0 0 0-1.543-.984a10.38 10.38 0 0 0-.737-.357c-.017-.007-.033-.017-.05-.024a9.759 9.759 0 0 0-.832-.315m1.351 11.95c0 4.387-2.136 6.786-6.109 6.786h-2.737V43.524h2.774c3.899 0 6.072 2.513 6.072 6.972"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialDeviantArt(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M46.347 44.786c2.026-.797 7.307-1.86 11.093-1.328l4.749 9.896l13.052-3.089c.664-.133.996-.996 0-1.893c-8.17-6.575-16.24-7.87-32.181-4.383l9.465 19.727l-33.044 9.316v5.995c0 .219.089.419.233.562a.791.791 0 0 0 .563.233h59.248a.794.794 0 0 0 .796-.795V56.123l-25.206 6.696s-7.075-14.247-8.768-18.033"/><path fill="currentColor" d="M79.524 19.513H20.276a.775.775 0 0 0-.437.148a.782.782 0 0 0-.36.648v25.81c1.266-.977 5.98-4.468 10.528-5.917l-3.886-8.967H37.08s2.89 6.276 2.89 6.376c3.985-1.195 23.114-4.284 40.151 3.089l.199.146V20.31a.781.781 0 0 0-.36-.648a.766.766 0 0 0-.436-.149M27.185 61.89c2.458-.531 11.823-3.056 11.823-3.056l-6.045-12.421c-1.66.465-8.634 3.653-8.103 14.347c0 0-.133 1.661 2.325 1.13"/><path fill="currentColor" d="M80.022 14.133H19.978A5.979 5.979 0 0 0 14 20.11v59.78a5.978 5.978 0 0 0 5.978 5.977h60.044A5.979 5.979 0 0 0 86 79.89V20.11c0-3.3-2.677-5.977-5.978-5.977m2.291 64.893a2.792 2.792 0 0 1-2.789 2.789H20.276a2.793 2.793 0 0 1-2.79-2.789V20.31a2.793 2.793 0 0 1 2.79-2.789h59.248a2.792 2.792 0 0 1 2.789 2.789z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialDigg(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M92.546 38.012c0-.333-.035-.667-.059-1.096h-.903c-6.932 0-13.863.014-20.794-.017c-.825-.004-1.055.236-1.052 1.055c.029 7.811.029 15.623-.001 23.435c-.003.838.26 1.05 1.065 1.042c3.977-.036 7.954-.018 11.931-.014c.33 0 .659.029 1.071.049c0 1.126-.046 2.173.015 3.214c.044.745-.182.972-.944.964c-4.062-.037-8.125.004-12.187-.035c-.772-.008-.97.239-.947.966a67.357 67.357 0 0 1-.002 4.517c-.027.759.225.944.968.941a1974.21 1974.21 0 0 1 20.793.008c.914.006 1.065-.293 1.062-1.111c-.021-11.307-.013-22.613-.016-33.918m-9.588 18.022a61.247 61.247 0 0 0-3.662-.002c-.605.016-.854-.16-.85-.816c.026-3.719.004-7.438.031-11.157c.002-.254.265-.721.419-.727c1.608-.054 3.219-.032 4.909-.032v2.229c0 3.209-.018 6.417.014 9.625c.007.667-.188.904-.861.88M66.026 36.899c-6.902.029-13.805.03-20.707-.001c-.814-.004-1.054.215-1.051 1.04c.03 7.811.03 15.622 0 23.434c-.003.81.203 1.069 1.039 1.061c4.004-.041 8.01.012 12.014-.041c.901-.012 1.142.278 1.068 1.113a14.354 14.354 0 0 0-.002 2.214c.048.701-.195.927-.91.922c-4.033-.03-8.067-.014-12.1-.014h-1.091c0 1.898.038 3.683-.018 5.464c-.024.767.227.944.966.941c6.959-.028 13.918-.016 20.877-.018c.277-.001.554-.024.905-.041c.021-.347.056-.65.056-.953c.002-11.362-.007-22.724.02-34.086c.001-.839-.264-1.038-1.066-1.035m-7.678 18.559c-.002.188-.349.528-.544.534a69.417 69.417 0 0 1-4.256.001c-.176-.004-.493-.316-.494-.488c-.028-4.025-.021-8.053-.021-12.216c1.696 0 3.253-.02 4.808.026c.183.005.509.365.511.562c.028 3.86.028 7.722-.004 11.581M21.554 27.974v7.671c0 1.271 0 1.272-1.312 1.272c-3.949 0-7.898.018-11.847-.016c-.746-.007-.96.202-.958.957c.026 8.607.027 17.217-.002 25.824c-.003.799.232.995 1.006.991c6.932-.027 13.864-.016 20.796-.018c.303 0 .605-.025.927-.04V26.96h-8.611zm-.999 30.307c-1.218-.058-2.441.01-3.66-.04c-.234-.009-.654-.313-.655-.484c-.036-4.766-.029-9.53-.029-14.404h5.343v.908c0 4.342-.02 8.685.018 13.027c.007.801-.248 1.028-1.017.993m20.297-21.37c-2.328.013-4.658.037-6.985-.012c-.781-.017-.952.268-.941.985c.036 2.47.014 4.94.014 7.412c0 6.134.01 12.268-.013 18.401c-.003.713.146 1.036.937.975c.903-.07 1.816-.017 2.725-.017h4.961c.04-.232.067-.315.067-.398c.004-8.889.001-17.776.021-26.665c0-.619-.297-.683-.786-.681m-7.219-3.4c2.442.037 4.886.014 7.329.031c.53.004.69-.209.672-.712c-.03-.823-.007-1.647-.007-2.472V26.96h-8.68c0 2.08-.017 4.064.03 6.049c.003.18.424.499.656.502"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialDribbble(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M50 13.766c-19.979 0-36.234 16.255-36.234 36.233c0 19.98 16.255 36.234 36.234 36.234S86.234 69.98 86.234 50c0-19.979-16.254-36.234-36.234-36.234m23.957 16.709a30.803 30.803 0 0 1 6.989 19.28c-1.021-.218-11.253-2.291-21.551-.997a95.897 95.897 0 0 0-2.72-6.079c11.443-4.674 16.63-11.32 17.282-12.204M50 19.097c7.864 0 15.051 2.95 20.517 7.798c-.557.794-5.216 7.056-16.221 11.181c-5.071-9.318-10.693-16.973-11.548-18.119a30.96 30.96 0 0 1 7.252-.86m-13.157 2.942c.815 1.118 6.346 8.783 11.474 17.898c-14.477 3.845-27.228 3.788-28.612 3.771c2.006-9.603 8.478-17.585 17.138-21.669M19.05 50.046c0-.316.006-.631.016-.945c1.354.028 16.353.221 31.808-4.404a103.124 103.124 0 0 1 2.509 5.253c-.409.115-.816.237-1.22.368c-15.962 5.159-24.456 19.224-25.157 20.424c-4.942-5.488-7.956-12.746-7.956-20.696M50 80.998a30.798 30.798 0 0 1-18.998-6.533c.558-1.139 6.825-13.246 24.281-19.328l.205-.069c4.345 11.291 6.141 20.755 6.602 23.472A30.849 30.849 0 0 1 50 80.998m17.292-5.294c-.316-1.883-1.963-10.943-6.012-22.067c9.698-1.552 18.219.989 19.278 1.322c-1.379 8.613-6.322 16.05-13.266 20.745" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialDrive(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M40.588 59.659L27.579 82.03h46.912L87.5 59.659zm45.364-1.28L62.496 17.753l-25.879-.081L60.073 58.3zm-49.996-39.13L12.5 59.876l12.87 22.452l23.456-40.627z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialDropbox(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M34.063 13.993L11.288 28.864l15.747 12.612L50 27.295zM11.288 54.087l22.775 14.869L50 55.653L27.035 41.476zM50 55.653l15.937 13.303l22.775-14.869l-15.747-12.611zm38.712-26.789L65.937 13.993L50 27.295l22.965 14.181z"/><path fill="currentColor" d="M50.046 58.517L34.063 71.778l-6.84-4.464v5.005l22.823 13.688l22.825-13.688v-5.005l-6.841 4.464z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialEvernote(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23.851 26.072h7.336a.76.76 0 0 0 .759-.758c0-.004-.089-6.286-.089-8.033v-.021c0-1.436.301-2.686.826-3.735l.25-.471a.153.153 0 0 0-.083.048L18.605 27.226a.16.16 0 0 0-.054.083c.295-.147.698-.345.755-.37c1.24-.56 2.745-.867 4.545-.867"/><path fill="currentColor" d="M80.894 24.535c-.582-3.117-2.435-4.654-4.111-5.258c-1.808-.653-5.474-1.33-10.08-1.873c-3.706-.436-8.06-.401-10.691-.32c-.316-2.162-1.833-4.138-3.53-4.822c-4.522-1.818-11.509-1.381-13.303-.878c-1.425.398-3.003 1.216-3.881 2.474c-.589.84-.971 1.917-.974 3.421c0 .852.024 2.856.045 4.639c.022 1.788.045 3.386.045 3.398a2.892 2.892 0 0 1-2.889 2.889h-7.332c-1.564 0-2.759.264-3.672.678a5.102 5.102 0 0 0-2.053 1.637c-.979 1.312-1.15 2.936-1.145 4.592c0 0 .014 1.352.34 3.971c.271 2.026 2.468 16.182 4.555 20.485c.809 1.674 1.349 2.372 2.938 3.109c3.542 1.518 11.633 3.205 15.424 3.688c3.785.484 6.161 1.503 7.577-1.467c.004-.006.283-.738.667-1.81c1.229-3.724 1.4-7.03 1.4-9.42c0-.243.356-.255.356 0c0 1.689-.323 7.662 4.191 9.265c1.781.632 5.478 1.195 9.233 1.635c3.396.391 5.86 1.729 5.86 10.439c0 5.301-1.113 6.027-6.933 6.027c-4.719 0-6.517.122-6.517-3.626c0-3.028 2.997-2.712 5.217-2.712c.992 0 .272-.737.272-2.606c0-1.86 1.163-2.934.063-2.962c-7.678-.211-12.196-.01-12.196 9.594c0 8.717 3.336 10.337 14.234 10.337c8.544 0 11.556-.28 15.084-11.227c.698-2.162 2.386-8.756 3.409-19.829c.647-7.003-.607-28.134-1.603-33.468M66 47.892a15.484 15.484 0 0 0-3.018.182c.268-2.153 1.155-4.793 4.301-4.682c3.483.119 3.972 3.411 3.984 5.642c-1.469-.656-3.287-1.073-5.267-1.142"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialFacebook(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M82.026 14H17.974A3.974 3.974 0 0 0 14 17.974v64.053A3.974 3.974 0 0 0 17.974 86h34.483V58.118h-9.383V47.252h9.383v-8.014c0-9.3 5.68-14.363 13.976-14.363c3.974 0 7.389.295 8.385.428v9.719l-5.754.003c-4.512 0-5.385 2.144-5.385 5.29v6.938h10.76l-1.401 10.866h-9.359V86h18.348A3.974 3.974 0 0 0 86 82.026V17.974A3.974 3.974 0 0 0 82.026 14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialFiveHundredPx(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M65.193 43.131c-1.031 0-2.022.22-2.975.656c-.954.438-1.854.992-2.704 1.661a24.391 24.391 0 0 0-2.434 2.201c-.772.799-1.468 1.532-2.086 2.201a35.8 35.8 0 0 0 2.163 2.278a18.209 18.209 0 0 0 2.472 2.008a13.16 13.16 0 0 0 2.743 1.429c.952.361 1.97.541 3.051.541c1.905 0 3.373-.618 4.404-1.854c1.03-1.236 1.545-2.754 1.545-4.557c0-1.802-.554-3.348-1.661-4.635c-1.107-1.285-2.613-1.929-4.518-1.929m-21.631 4.518a26.225 26.225 0 0 0-2.588-2.201a15.069 15.069 0 0 0-2.82-1.661c-.979-.437-1.958-.656-2.935-.656c-1.803 0-3.233.683-4.288 2.047c-1.056 1.364-1.584 2.896-1.584 4.596c0 1.854.541 3.373 1.622 4.557c1.082 1.185 2.575 1.776 4.481 1.776c.977 0 1.956-.193 2.935-.579a14.388 14.388 0 0 0 2.781-1.468c.875-.592 1.7-1.261 2.472-2.008a32.164 32.164 0 0 0 2.086-2.201a56.593 56.593 0 0 0-2.162-2.202"/><path d="M79 14H21a7 7 0 0 0-7 7v58a7 7 0 0 0 7 7h58a7 7 0 0 0 7-7V21a7 7 0 0 0-7-7m-1.292 41.102c-.618 1.751-1.494 3.282-2.627 4.596c-1.134 1.312-2.538 2.344-4.21 3.089c-1.674.747-3.542 1.12-5.601 1.12c-1.597 0-3.09-.244-4.481-.733a17.68 17.68 0 0 1-3.901-1.931a23.25 23.25 0 0 1-3.438-2.781a52.75 52.75 0 0 1-3.09-3.282a413.808 413.808 0 0 1-3.168 3.282a22.075 22.075 0 0 1-3.284 2.781a16.404 16.404 0 0 1-3.785 1.931c-1.365.489-2.897.733-4.597.733c-2.112 0-3.992-.373-5.64-1.12c-1.649-.745-3.065-1.749-4.249-3.012c-1.186-1.261-2.086-2.768-2.704-4.519c-.618-1.75-.927-3.63-.927-5.639c0-2.008.295-3.873.888-5.6c.591-1.725 1.455-3.217 2.588-4.479c1.132-1.261 2.523-2.265 4.172-3.013c1.648-.745 3.502-1.119 5.562-1.119c1.7 0 3.257.258 4.674.773a16.842 16.842 0 0 1 3.94 2.046c1.21.85 2.343 1.829 3.399 2.935a71.194 71.194 0 0 1 3.129 3.515a76.135 76.135 0 0 1 3.051-3.476a21.6 21.6 0 0 1 3.399-2.975a16.803 16.803 0 0 1 3.94-2.046c1.416-.515 2.948-.773 4.597-.773c2.06 0 3.914.36 5.562 1.081c1.647.722 3.038 1.713 4.172 2.974c1.132 1.262 2.008 2.742 2.627 4.441c.618 1.699.927 3.553.927 5.561s-.306 3.89-.925 5.64"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialFlickr(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M79 14H21c-3.85 0-7 3.15-7 7v58c0 3.85 3.15 7 7 7h58c3.85 0 7-3.15 7-7V21c0-3.85-3.15-7-7-7M35.435 61.5c-6.456 0-11.69-5.233-11.69-11.689s5.234-11.69 11.69-11.69c6.456 0 11.689 5.233 11.689 11.69c0 6.456-5.233 11.689-11.689 11.689m29.633 0c-6.456 0-11.69-5.233-11.69-11.689s5.233-11.69 11.69-11.69c6.456 0 11.689 5.233 11.689 11.69c.001 6.456-5.233 11.689-11.689 11.689"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialForrst(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M84.173 91L49.968 9L15.827 90.853h28.801V73.984l-8.655-5.935l2.119-3.139l6.536 4.464V56.299h8.947v8.306l6.466-3.414l1.864 3.426l-8.33 4.407v4.708l12.537-6.274l1.665 3.377l-14.202 7.123v12.895h17.823V91z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialFoursquare(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M36.217 40.748a2.651 2.651 0 0 0-1.878-.78a2.688 2.688 0 0 0-1.881.776l-6.322 6.315a2.653 2.653 0 0 0-.003 3.759L47.531 72.26c.429.432.986.697 1.579.762c.099.016.201.02.304.02c.007 0 .013 0 .019-.002l.003.002a2.687 2.687 0 0 0 1.899-.781L89.72 33.856a2.654 2.654 0 0 0 .002-3.757l-6.312-6.322a2.626 2.626 0 0 0-1.881-.778a2.639 2.639 0 0 0-1.883.776L49.438 53.994z"/><path fill="currentColor" d="M87.348 43.137L54.835 75.666a7.4 7.4 0 0 1-5.24 2.154l-1.443-.002l-.304-.204a7.39 7.39 0 0 1-3.518-1.97L22.939 54.208a7.346 7.346 0 0 1-2.174-5.238c0-1.981.774-3.845 2.175-5.242l6.33-6.322a7.483 7.483 0 0 1 5.236-2.157c1.952 0 3.862.797 5.234 2.175l9.863 9.88l20.953-20.96l-13.957-13.957c-3.694-3.693-9.738-3.693-13.432 0L12.271 43.283c-3.695 3.693-3.695 9.74 0 13.433l30.897 30.897c3.694 3.693 9.737 3.693 13.432 0l30.897-30.897c3.694-3.692 3.694-9.737 0-13.433z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialGameCenter(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M79 14H21a7 7 0 0 0-7 7v58a7 7 0 0 0 7 7h58a7 7 0 0 0 7-7V21a7 7 0 0 0-7-7m-23.045 5.229c10.658 0 19.298 8.641 19.298 19.299a19.2 19.2 0 0 1-3.482 11.039a16.854 16.854 0 0 0-10.465-3.634c-6.582 0-12.275 3.76-15.082 9.242c-5.718-3.35-9.568-9.541-9.568-16.646c0-10.659 8.64-19.3 19.299-19.3M19.852 46.931c0-11.635 8.755-21.22 20.036-22.54c-3.324 3.774-5.356 8.713-5.356 14.138c0 7.981 4.373 14.928 10.845 18.614a17.449 17.449 0 0 0-.64 2.295a11.95 11.95 0 0 0-3.429-.505c-4.89 0-9.088 2.941-10.937 7.148c-6.323-4.03-10.519-11.098-10.519-19.15m21.454 34.145c-5.632 0-10.197-4.564-10.197-10.197c0-5.631 4.565-10.196 10.197-10.196c5.631 0 10.197 4.565 10.197 10.196c0 5.633-4.566 10.197-10.197 10.197m20-3.25a14.862 14.862 0 0 1-8.782-2.87c.463-1.274.729-2.643.729-4.077c0-4.689-2.708-8.735-6.639-10.691c1.269-6.968 7.357-12.255 14.693-12.255c8.255 0 14.947 6.692 14.947 14.946c-.001 8.256-6.693 14.947-14.948 14.947"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialGithub(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M49.998 11.963C28.461 11.963 11 29.425 11 50.965c0 17.231 11.172 31.849 26.671 37.003c1.952.361 2.662-.84 2.662-1.877c0-.924-.034-3.375-.051-6.633c-10.849 2.359-13.138-5.229-13.138-5.229c-1.774-4.505-4.331-5.703-4.331-5.703c-3.541-2.418.269-2.371.269-2.371c3.914.277 5.974 4.018 5.974 4.018c3.478 5.96 9.129 4.235 11.35 3.243c.353-2.525 1.363-4.24 2.476-5.217c-8.659-.984-17.763-4.33-17.763-19.274c0-4.259 1.519-7.741 4.013-10.468c-.399-.982-1.74-4.947.383-10.319c0 0 3.274-1.048 10.726 4.001c3.109-.869 6.446-1.303 9.763-1.316c3.312.014 6.65.447 9.763 1.316c7.447-5.049 10.716-4.001 10.716-4.001c2.128 5.372.788 9.337.388 10.319c2.5 2.727 4.008 6.209 4.008 10.468c0 14.979-9.117 18.279-17.805 19.241c1.398 1.205 2.646 3.59 2.646 7.229c0 5.211-.047 9.416-.047 10.695c0 1.045.701 2.26 2.681 1.873C77.836 82.798 89 68.191 89 50.965c0-21.54-17.461-39.002-39.002-39.002" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialGooglePlus(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M25.194 51.591c1.757 1.379 3.766 2.068 6.026 2.068c2.855-.105 5.234-1.131 7.138-3.076c.918-1.377 1.508-2.791 1.77-4.24c.157-1.447.236-2.668.236-3.658c0-4.276-1.094-8.588-3.281-12.936c-1.026-2.084-2.376-3.781-4.051-5.09c-1.71-1.236-3.675-1.891-5.897-1.961c-2.939.07-5.384 1.256-7.331 3.553c-1.644 2.404-2.43 5.09-2.358 8.057c0 3.924 1.147 8.008 3.444 12.248a16.741 16.741 0 0 0 4.304 5.035m6.587 16.459a23.029 23.029 0 0 1-2.091-3.299c-.714-1.133-1.07-2.5-1.07-4.096c0-.957.135-1.756.407-2.395a366.3 366.3 0 0 1 .663-1.861a30.289 30.289 0 0 1-3.485.213c-5.169-.062-9.237-1.527-12.205-4.388v19.338a33.422 33.422 0 0 1 8.241-2.561c3.829-.564 7.009-.881 9.54-.951m2.959 3.064a23.873 23.873 0 0 0-2.435-.107c-.553-.07-1.969 0-4.248.213a46.623 46.623 0 0 0-6.89 1.533c-.553.213-1.33.529-2.331.953c-1.002.459-2.021 1.111-3.056 1.959A10.843 10.843 0 0 0 14 77.69v2.949a5.018 5.018 0 0 0 5.013 5.014h27.326c.001-.068.005-.133.005-.199c0-2.752-.899-5.15-2.694-7.197c-1.9-1.94-4.869-4.32-8.91-7.143"/><path fill="currentColor" d="M80.986 14.347H19.013A5.016 5.016 0 0 0 14 19.361v6.723c.316-.336.643-.67.985-1c3.083-2.539 6.288-4.195 9.612-4.973c3.29-.67 6.375-1.004 9.253-1.004h21.692l-6.699 3.91h-6.689c.681.422 1.448 1.041 2.299 1.852c.817.846 1.618 1.887 2.401 3.121c.749 1.164 1.413 2.539 1.992 4.127c.476 1.586.715 3.44.715 5.555c-.064 3.881-.921 6.984-2.569 9.312a29.44 29.44 0 0 1-2.569 3.121a102.307 102.307 0 0 1-3.199 2.91a15.086 15.086 0 0 0-1.793 2.275c-.646.881-.97 1.904-.97 3.068c0 1.129.332 2.062.994 2.803a21.86 21.86 0 0 0 1.641 1.854l3.68 3.016a36.958 36.958 0 0 1 6.016 6.189c1.625 2.293 2.47 5.291 2.536 8.994c0 1.531-.198 3.01-.592 4.439h28.25A5.018 5.018 0 0 0 86 80.639V19.36a5.017 5.017 0 0 0-5.014-5.013m3.945 22.117H73.549v11.383h-5.514V36.464H56.653V30.95h11.382V19.569h5.514V30.95h11.382z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialHackerNews(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M79 14H21a7 7 0 0 0-7 7v58a7 7 0 0 0 7 7h58a7 7 0 0 0 7-7V21a7 7 0 0 0-7-7M53.134 52.654V69.83h-5.267V52.714L35.12 29.494h5.925l5.685 11.132c1.556 3.051 2.752 5.505 4.01 8.318h.12c1.137-2.634 2.514-5.268 4.07-8.318l5.805-11.132h5.924z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialHiFive(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m93.246 44.073l.731-15.993H62.701l-1.564 8.465v-6.128H43.966v7.947h-5.567V40.1c-3.268-2.575-8.316-2.712-12.109-1.3V28.08H3.551v15.041h2.956v12.864H3.55v14.903h60.542v-1.955c3.732 2.115 8.155 2.986 12.412 2.986c5.88 0 10.547-1.768 13.888-4.518c3.461-2.849 5.606-7.444 5.994-12.132c.331-4.007-.623-8.084-3.14-11.196m-45.334-9.71h9.279v6.255h-9.279zm-6.264 32.58H25.361v-7.011h1.557V51.37c0-4.036-4.575-2.609-4.575 4.052v4.51h1.375v7.011H7.497v-7.011h2.955V39.175H7.497v-7.148h14.846V47.18c1.955-3.149 4.606-5.295 8.998-5.295c6.221 0 7.154 5.391 7.154 7.53v10.517h3.152zm18.498 0H42.413v-7.011H45.3V49.484h-2.956V42.31H57.19v17.621h2.956zm16.358 1.031c-6.803 0-12.444-2.532-14.475-6.278c-2.098-3.868-.452-9.839 4.657-9.849c3.333-.006 5.436 1.477 4.732 5.283c-.884 2.564.797 3.414 2.423 3.414c2.464 0 4.336-2.837 4.464-5.733c.12-2.721-1-5.366-3.508-5.395c-1.017-.012-2.227.371-3.379 1.778l-8.77-1.183l3.323-17.985h23.886l-.43 9.416H73.663l-.362 3.016c2.267-1.005 4.984-1.689 7.482-1.689c6.628 0 11.978 3.343 11.667 12.077c0 6.738-5.512 13.128-15.946 13.128"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialInstagram(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M67.961 45.288c0 9.919-8.041 17.96-17.96 17.96s-17.96-8.041-17.96-17.96c0-2.795.657-5.43 1.795-7.788H14v42c0 3.85 3.15 7 7 7h58c3.85 0 7-3.15 7-7v-42H66.165a17.848 17.848 0 0 1 1.796 7.788M79 13.5H21c-3.85 0-7 3.15-7 7v14h21.665c3.278-4.349 8.47-7.172 14.335-7.172s11.057 2.823 14.335 7.172H86v-14c0-3.85-3.15-7-7-7M19.756 31.216H17.15V21.121a4.253 4.253 0 0 1 2.607-3.922zm4.27 0h-2.469V16.868h2.469zm4.269 0h-2.469V16.868h2.469zm4.266-10.095v10.095h-2.466V16.868h2.466zm49.79 5.842a4.252 4.252 0 0 1-4.251 4.253h-6.909a4.252 4.252 0 0 1-4.252-4.253v-5.842a4.252 4.252 0 0 1 4.252-4.253H78.1a4.251 4.251 0 0 1 4.251 4.253z"/><path d="M50.001 61.148c8.736 0 15.818-7.085 15.818-15.819c0-8.735-7.081-15.82-15.818-15.82c-8.738 0-15.82 7.085-15.82 15.82c0 8.733 7.082 15.819 15.82 15.819M50 32.725c6.96 0 12.604 5.643 12.604 12.604c0 6.96-5.643 12.604-12.604 12.604S37.396 52.29 37.396 45.329C37.397 38.367 43.04 32.725 50 32.725"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialJoomla(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M81.433 68.316c-1.245-1.029-2.883-1.582-4.507-2.443c1.247-6.475-1.226-11.717-6.169-16.113c-2.454-2.184-4.671-4.638-6.935-7.027c-2.06-2.176-4.038-4.43-6.348-6.976l-6.986 7.634c.418.318.987.655 1.438 1.109c4.596 4.621 9.167 9.27 13.76 13.895c1.703 1.715 2.339 3.752 1.544 6.012c-.786 2.234-2.489 3.82-4.891 3.749c-1.728-.05-3.432-.926-5.163-1.438l-6.474 6.37c4.344 4.493 9.656 5.624 15.729 4.274c.854 3.877 3.5 5.938 6.805 7.368h4.652c.324-.168.638-.363.975-.501c6.762-2.762 8.183-11.269 2.57-15.913"/><path fill="currentColor" d="M56.4 51.269c-.291.411-.625 1.066-1.122 1.559c-4.63 4.592-9.271 9.173-13.943 13.722c-2.322 2.261-5.26 2.342-7.651.31c-2.307-1.962-2.666-5.088-.869-7.652c.283-.403.6-.783.745-.971l-6.729-6.793c-4.257 4.049-5.51 9.106-4.458 14.934c-4.266 1.211-6.417 4.191-7.282 8.273v1.938c.938 4.087 3.357 6.832 7.365 8.141h4.264c2.657-1.11 5.172-2.475 6.204-5.362c.519-1.451 1.195-1.79 2.67-1.595c4.759.63 8.983-.766 12.37-4.129c5.173-5.138 10.189-10.435 15.56-15.96zM35.946 42.96l6.537 6.358c2.867-2.889 5.67-5.742 8.507-8.561c2.288-2.273 4.571-4.553 6.939-6.742c2.057-1.901 4.531-2.146 6.761-.85c2.125 1.235 3.316 3.697 2.677 6.104c-.318 1.197-1.119 2.266-1.673 3.338l6.65 6.669c4.362-4.219 5.631-9.397 4.411-15.264c6.022-2.594 8.531-5.794 8.055-10.204c-.457-4.231-3.475-7.902-7.386-8.341c-1.958-.22-4.138.099-5.998.776c-3.06 1.114-4.471 3.871-5.263 6.897c-.565-.045-.955-.033-1.325-.112c-4.585-.969-8.977-.134-12.308 2.989c-5.727 5.372-11.017 11.207-16.584 16.943"/><path fill="currentColor" d="M22.564 33.799c-1.655 6.861.583 12.263 5.552 16.803c2.381 2.176 4.605 4.527 6.86 6.837c2.244 2.299 4.438 4.646 6.653 6.97l6.918-6.839c-.526-.554-1-1.065-1.489-1.563c-4.568-4.649-9.164-9.274-13.7-13.955c-2.186-2.256-2.243-5.441-.27-7.695c2.055-2.347 5.236-2.633 7.811-.705c.34.255.69.496.633.455l6.786-6.805c-3.857-3.83-8.686-5.32-14.261-4.3c-2.728-6.083-5.454-8.083-10.412-7.681c-3.768.305-7.274 3.067-8.018 6.787c-1.162 5.824 1.469 9.659 6.937 11.691"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialLastfm(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m78.404 46.408l-5.97-1.306c-4.104-.932-5.316-2.611-5.316-5.41c0-3.171 2.518-5.037 6.622-5.037c4.478 0 6.903 1.68 7.276 5.69l9.328-1.119c-.746-8.396-6.529-11.848-16.044-11.848c-8.395 0-16.604 3.172-16.604 13.34c0 6.344 3.078 10.354 10.82 12.22l6.344 1.492c4.757 1.12 6.342 3.079 6.342 5.784c0 3.451-3.357 4.85-9.701 4.85c-9.421 0-13.338-4.943-15.578-11.753l-3.078-9.327C48.928 31.857 42.678 27.378 30.27 27.378c-13.712 0-20.988 8.676-20.988 23.414c0 14.18 7.276 21.828 20.336 21.828c10.541 0 15.578-4.944 15.578-4.944l-2.985-8.115s-4.851 5.411-12.126 5.411c-6.437 0-11.008-5.598-11.008-14.553c0-11.474 5.783-15.578 11.474-15.578c8.208 0 10.82 5.317 13.059 12.127l2.985 9.328c2.984 9.048 8.581 16.324 24.719 16.324c11.568 0 19.403-3.545 19.403-12.873c.001-7.555-4.291-11.473-12.313-13.339"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialLinkedin(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M80.667 14H19.315C16.381 14 14 16.325 14 19.188v61.617C14 83.672 16.381 86 19.315 86h61.352C83.603 86 86 83.672 86 80.805V19.188C86 16.325 83.603 14 80.667 14M35.354 75.354H24.67V40.995h10.684zm-5.342-39.057a6.19 6.19 0 0 1-6.19-6.194a6.189 6.189 0 1 1 12.379 0a6.194 6.194 0 0 1-6.189 6.194M75.35 75.354H64.683V58.646c0-3.986-.078-9.111-5.551-9.111c-5.558 0-6.405 4.341-6.405 8.822v16.998H42.052v-34.36h10.245v4.692h.146c1.426-2.7 4.91-5.549 10.106-5.549c10.806 0 12.802 7.114 12.802 16.369v18.847z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialMedium(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M81.652 14H18.001A7.06 7.06 0 0 0 14 18.336v62.793C14.917 83.955 17.562 86 20.687 86h58.279C82.851 86 86 82.843 86 78.948V20.517A7.05 7.05 0 0 0 81.652 14M67.867 40.385h-1.4c-.52 0-1.119.682-1.119 1.163v17.437c0 .482.6 1.203 1.119 1.203h1.4v4.048H55.312v-4.048h2.519v-18.32h-.12l-6.198 22.368h-4.798l-6.118-22.368h-.12v18.32h2.639v4.048H32.481v-4.048h1.359c.56 0 1.16-.721 1.16-1.203V41.548c0-.481-.6-1.163-1.16-1.163h-1.359v-4.209h13.275l4.358 16.275h.12l4.398-16.275h13.235z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialMyspace(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M78.841 51.036c7.792 0 14.111-6.294 14.111-14.061c0-7.765-6.319-14.061-14.111-14.061c-7.798 0-14.118 6.297-14.118 14.061c-.001 7.767 6.32 14.061 14.118 14.061"/><ellipse cx="47.046" cy="40.984" fill="currentColor" rx="12.703" ry="12.656"/><path fill="currentColor" d="M18.214 55.984c6.313 0 11.433-5.096 11.433-11.386c0-6.292-5.12-11.39-11.433-11.39c-6.315 0-11.433 5.098-11.433 11.39c0 6.291 5.117 11.386 11.433 11.386m0 2.601c-7.25 0-12.565 6.363-12.565 12.936v4.425c0 .626.512 1.14 1.142 1.14h22.843c.632 0 1.144-.514 1.144-1.14v-4.425c0-6.573-5.315-12.936-12.564-12.936m28.832-2.059c-8.055 0-13.962 7.071-13.962 14.376v4.917c0 .695.569 1.267 1.269 1.267h25.382a1.27 1.27 0 0 0 1.271-1.267v-4.917c.001-7.304-5.905-14.376-13.96-14.376m31.793-2.283c-8.95 0-15.512 7.856-15.512 15.974v5.462c0 .773.632 1.407 1.41 1.407h28.2c.782 0 1.414-.635 1.414-1.407v-5.462c0-8.117-6.562-15.974-15.512-15.974"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialOrkut(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M50.104 34.348c-8.863 0-16.062 7.198-16.062 16.076s7.199 16.062 16.062 16.062c8.878 0 16.077-7.185 16.077-16.062c0-8.878-7.199-16.076-16.077-16.076"/><path d="M79 14H21a7 7 0 0 0-7 7v58a7 7 0 0 0 7 7h58a7 7 0 0 0 7-7V21a7 7 0 0 0-7-7M50.104 74.707c-13.404 0-24.297-10.864-24.297-24.283S36.7 26.113 50.104 26.113c13.419 0 24.297 10.893 24.297 24.311c0 13.419-10.878 24.283-24.297 24.283"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialPath(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M84 41.515c0-16.82-13.663-28.125-34-28.125c-20.015 0-34 11.566-34 28.125c0 5.924 2.539 14.568 7.225 18.93l7.917-8.503c-1.744-1.621-3.519-6.884-3.519-10.427c0-11.408 11.244-16.516 22.377-16.516c10.824 0 22.378 4.341 22.378 16.516c0 13.641-15.881 15.154-21.729 15.242V34.755H38.24v35.283c0 3.574-3.863 5.129-7.984 2.14v12.733c2.565 1.26 4.895 1.699 6.891 1.699c2.235 0 4.067-.555 5.333-1.119c4.89-2.152 8.169-7.258 8.169-12.691v-4.441C70.921 68.137 84 57.659 84 41.515"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialPicasa(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M66.078 14.577a38.608 38.608 0 0 0-16.08-3.495c-4.65 0-9.221.894-13.62 2.542c1.031.936 28.368 25.748 29.7 26.959zM29.924 83.389V49.858A35455.873 35455.873 0 0 1 13.81 64.5c3.126 7.839 8.832 14.487 16.114 18.889M11 50.082c0 3.35.505 6.757 1.416 10.132c1.058-.959 33.894-30.793 34.569-31.408c-.659-.6-14.314-12.995-14.759-13.398C19.134 22.127 11 35.377 11 50.082m23.019 16.952v18.605c3.965 1.776 8.111 2.885 12.378 3.279h7.161c13.471-1.256 25.498-9.561 31.474-21.885zm36.17-50.221v46.111h16.598C88.24 58.767 89 54.448 89 50.082c0-13.553-7.206-26.206-18.811-33.269"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialPinterest(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M50.001 10.999C28.464 10.999 11 28.461 11 50c0 15.97 9.603 29.688 23.345 35.72c-.11-2.723-.021-5.994.679-8.953c.75-3.168 5.018-21.252 5.018-21.252s-1.246-2.49-1.246-6.17c0-5.779 3.35-10.098 7.52-10.098c3.547 0 5.261 2.664 5.261 5.855c0 3.566-2.273 8.896-3.443 13.84c-.978 4.139 2.075 7.512 6.156 7.512c7.389 0 12.365-9.49 12.365-20.736c0-8.545-5.756-14.943-16.227-14.943c-11.831 0-19.201 8.822-19.201 18.676c0 3.4 1.002 5.795 2.572 7.648c.721.853.821 1.196.56 2.176c-.187.717-.617 2.444-.794 3.129c-.259.988-1.059 1.34-1.953.976c-5.449-2.225-7.987-8.192-7.987-14.9c0-11.08 9.343-24.363 27.873-24.363c14.891 0 24.691 10.775 24.691 22.344c0 15.297-8.506 26.727-21.044 26.727c-4.209 0-8.17-2.277-9.527-4.861c0 0-2.264 8.984-2.743 10.722c-.827 3.007-2.446 6.013-3.926 8.354a38.983 38.983 0 0 0 11.053 1.6C71.538 89.001 89 71.538 89 50c0-21.539-17.462-39.001-38.999-39.001"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialRdio(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M90.326 30.746c-9.825.285-21.688-7.424-26.286-10.327c-.426-.27-.861-.518-1.301-.742a31.605 31.605 0 0 0-1.477-.761v28.36h-.014c.018 2.54-.736 5.217-2.324 7.784l-.139.226c-4.594 7.427-14.258 11.756-23.322 8.695C26.66 61.005 24.309 52.48 28.97 44.94l.141-.227c4.589-7.424 14.253-11.754 23.318-8.693c.699.236 1.344.517 1.957.823v-20.51a39.351 39.351 0 0 0-10.313-1.381C23.094 14.952 7.518 30.654 7.518 50v.252c0 19.346 15.449 34.797 36.304 34.797C64.8 85.048 80.378 69.347 80.378 50v-.25c0-1.51-.105-2.988-.287-4.443c10.927-2.849 15.432-14.155 10.235-14.561"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialReddit(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M94.762 48.994c0-5.688-4.63-10.314-10.315-10.314c-2.463 0-4.767.901-6.626 2.477c-.06.037-.122.072-.181.11c-6.707-4.291-15.601-7.031-25.439-7.403l5.872-16.698l14.656 3.504c.012 4.633 3.781 8.4 8.42 8.4c4.642 0 8.422-3.777 8.422-8.421c0-4.646-3.78-8.423-8.422-8.423c-3.529 0-6.544 2.182-7.794 5.26l-17.364-4.15l-7.211 20.49c-10.259.193-19.556 2.969-26.513 7.404c-1.873-1.625-4.21-2.551-6.718-2.551c-5.687 0-10.31 4.627-10.31 10.314c0 3.518 1.815 6.768 4.756 8.66a18.114 18.114 0 0 0-.293 3.123c0 14.886 18.043 26.997 40.219 26.997c22.18 0 40.224-12.111 40.224-26.997a18.14 18.14 0 0 0-.272-3.035c3.02-1.878 4.889-5.175 4.889-8.747M63.598 62.347a6.333 6.333 0 0 1-6.334-6.338a6.337 6.337 0 1 1 6.334 6.338m1.261 10.806c-.19.194-4.733 4.821-15.009 4.821c-10.333 0-14.463-4.689-14.636-4.891a1.623 1.623 0 0 1 .178-2.283a1.63 1.63 0 0 1 2.278.166c.092.104 3.54 3.771 12.18 3.771c8.784 0 12.639-3.798 12.68-3.835a1.62 1.62 0 0 1 2.329 2.251m-34.05-17.144a6.337 6.337 0 1 1 12.672.006a6.337 6.337 0 0 1-12.672-.006"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialSkillshare(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" d="M32.092 55.916h14.575c1.087 0 2.563 2.496 4.734 4.039v-1.776s-.07-1.053.348-1.69c.424-.639 1.308-.809 1.584-.809h14.575s4.401-.74 9.294-3.044l-53.794.514c4.642 2.082 8.684 2.766 8.684 2.766"/><path fill="currentColor" d="M16.868 55.111a24.875 24.875 0 0 1-1.061-.855c.247.396.609.703 1.061.855m2.67-39.55h60.708v35.516c2.128-1.18 3.271-1.984 4.054-2.499v-35.05a2.03 2.03 0 0 0-2.03-2.028H17.508c-1.123 0-2.033.91-2.033 2.028v35.14c.771.486 1.878 1.301 4.063 2.52z"/><path fill="currentColor" d="M87.791 47.978c-.523-.417-1.247-.415-2.396 0c-.335.122-.674.325-1.095.6c-.783.515-1.925 1.319-4.054 2.499c-.657.36-1.373.748-2.241 1.183c-.272.136-.534.247-.803.377c-4.893 2.304-9.294 3.044-9.294 3.044H53.333c-.276 0-1.16.17-1.584.809c-.418.638-.348 1.69-.348 1.69v1.776c.526.373 1.082.707 1.695.918c4.394 1.498 7.367.192 8.664-1.628c1.517-2.134 2.289-3.329 3.955-2.705c1.666.624 0 4.68-2.08 6.763c-2.085 2.084-4.584 3.749-8.12 3.749c-1.486 0-2.926-.223-4.114-.476v15.131c0 4.27 4.538 6.552 7.658 6.552c3.119 0 9.367-1.768 12.387-9.263c3.017-7.493-.205-17.381-.205-17.381s4.888-2.191 9.991-5.414c5.098-3.225 6.455-6.246 6.559-6.67c.1-.427.516-1.136 0-1.554"/><path fill="currentColor" d="M63.635 63.303c2.08-2.083 3.746-6.139 2.08-6.763c-1.666-.624-2.438.571-3.955 2.705c-1.297 1.82-4.27 3.126-8.664 1.628c-.613-.211-1.169-.545-1.695-.918c-2.17-1.543-3.647-4.039-4.734-4.039H32.092s-4.043-.684-8.684-2.766c-.466-.21-.935-.418-1.411-.657a59.295 59.295 0 0 1-2.459-1.306c-2.185-1.219-3.292-2.033-4.063-2.52c-.321-.203-.593-.355-.871-.455c-1.144-.416-1.873-.416-2.394 0c-.518.418-.104 1.128 0 1.553c.08.325.913 2.184 3.596 4.49a26.416 26.416 0 0 0 2.963 2.181c5.098 3.227 9.991 5.412 9.991 5.412s-3.227 9.886-.205 17.381c3.018 7.494 9.263 9.27 12.385 9.27c3.123 0 7.661-2.297 7.661-6.564V65.804s1.152.42 2.799.771c1.188.253 2.628.476 4.114.476c3.537.001 6.036-1.664 8.121-3.748m-2.618-28.922a9.059 9.059 0 0 0-9.058 9.057c0 4.999 4.055 9.056 9.058 9.056c4.998 0 9.052-4.057 9.052-9.056c0-5-4.054-9.057-9.052-9.057"/><path fill="currentColor" d="M39.544 34.381c-5 0-9.053 4.057-9.053 9.057c0 4.999 4.052 9.056 9.053 9.056c5.003 0 9.058-4.057 9.058-9.056c0-5-4.055-9.057-9.058-9.057"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialSkype(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M86.284 58.294c.555-2.541.851-5.176.851-7.883c0-20.283-16.442-36.727-36.727-36.727c-2.14 0-4.236.186-6.279.539a21.169 21.169 0 0 0-11.33-3.27c-11.761 0-21.299 9.534-21.299 21.3c0 3.93 1.069 7.609 2.929 10.768a36.99 36.99 0 0 0-.744 7.391c0 20.282 16.44 36.725 36.723 36.725c2.299 0 4.545-.215 6.729-.615a21.228 21.228 0 0 0 10.065 2.525c11.763 0 21.299-9.531 21.299-21.297a21.178 21.178 0 0 0-2.217-9.456m-17.037 9.67c-1.701 2.401-4.212 4.307-7.462 5.658c-3.223 1.342-7.081 2.021-11.479 2.021c-5.272 0-9.695-.932-13.147-2.762c-2.467-1.336-4.502-3.138-6.046-5.366c-1.561-2.249-2.354-4.483-2.354-6.636c0-1.343.518-2.51 1.534-3.465c1.006-.947 2.294-1.429 3.833-1.429c1.261 0 2.351.38 3.234 1.126c.845.718 1.571 1.776 2.151 3.148c.652 1.488 1.363 2.748 2.112 3.74c.718.949 1.757 1.754 3.081 2.379c1.329.631 3.125.952 5.333.952c3.036 0 5.527-.649 7.396-1.93c1.831-1.249 2.72-2.747 2.72-4.583c0-1.449-.466-2.59-1.421-3.482c-.998-.934-2.315-1.662-3.916-2.16c-1.669-.521-3.938-1.082-6.74-1.67c-3.81-.816-7.047-1.784-9.625-2.88c-2.633-1.12-4.757-2.674-6.308-4.617c-1.576-1.976-2.376-4.452-2.376-7.362c0-2.774.838-5.273 2.494-7.432c1.643-2.143 4.042-3.807 7.127-4.949c3.05-1.129 6.676-1.697 10.776-1.697c3.28 0 6.164.379 8.57 1.125c2.421.758 4.461 1.771 6.066 3.027c1.612 1.262 2.816 2.611 3.567 4.012c.765 1.416 1.152 2.827 1.152 4.191c0 1.312-.507 2.502-1.508 3.543c-1.006 1.045-2.273 1.574-3.771 1.574c-1.36 0-2.425-.33-3.163-.983c-.688-.608-1.402-1.556-2.191-2.919c-.917-1.738-2.023-3.11-3.295-4.078c-1.235-.938-3.299-1.414-6.131-1.414c-2.636 0-4.777.527-6.364 1.57c-1.533 1.004-2.276 2.154-2.276 3.521c0 .837.239 1.535.731 2.136c.525.64 1.258 1.195 2.191 1.664c.965.484 1.958.871 2.954 1.145c1.024.285 2.733.705 5.084 1.246c2.969.641 5.702 1.354 8.116 2.127c2.445.783 4.56 1.749 6.283 2.871c1.76 1.143 3.153 2.616 4.144 4.372c.989 1.765 1.49 3.938 1.49 6.467c0 3.026-.863 5.784-2.566 8.199"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialSmashingMag(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M45.718 55.325c-3.727-1.408-7.229-3.226-10.202-5.401c-.077-.056-.157-.111-.233-.168a31.408 31.408 0 0 1-1.259-.991c-.107-.088-.211-.178-.316-.267a28.525 28.525 0 0 1-1.03-.919c-.062-.058-.127-.115-.188-.173c-.373-.357-.73-.724-1.074-1.096c-.091-.098-.179-.197-.268-.296a22.28 22.28 0 0 1-.877-1.037c-.041-.052-.085-.102-.125-.154a19.66 19.66 0 0 1-.862-1.204c-.07-.104-.137-.21-.204-.315a18.347 18.347 0 0 1-.693-1.171c-.021-.039-.045-.077-.066-.117a16.715 16.715 0 0 1-.75-1.642a15.601 15.601 0 0 1-.463-1.318l-.02-.06a14.956 14.956 0 0 1-.393-1.762c-.08-.482-.144-.968-.176-1.464a25.967 25.967 0 0 1-.047-1.963c.002-.199.018-.392.025-.589c.016-.447.039-.892.081-1.329c.021-.219.051-.435.078-.652a20.54 20.54 0 0 1 .318-1.854c.091-.413.198-.82.315-1.224c.053-.183.104-.366.163-.547c.148-.454.316-.9.499-1.341c.049-.12.092-.242.144-.361a18.909 18.909 0 0 1 3.185-4.943l-15.07 3.608a5.637 5.637 0 0 0-4.171 6.793l12.535 52.421a5.637 5.637 0 0 0 6.793 4.171l6.451-1.543c-2.504-.829-6.833-2.409-10.982-4.623l6.015-13.644s7.079 6.304 17.242 6.304c10.165 0 14.682-11.93-4.375-19.129m42.243 13.307L75.425 16.211a5.636 5.636 0 0 0-6.793-4.171l-10.797 2.581a41.758 41.758 0 0 1 4.282.916l.357.097c.241.066.488.136.738.209c.131.038.261.075.394.115c.353.106.714.218 1.083.339c.116.038.237.081.355.12a53.161 53.161 0 0 1 1.321.464c.285.105.575.215.867.329c.133.052.263.1.397.154c.406.162.818.335 1.235.515c.141.061.285.127.428.19a51.746 51.746 0 0 1 1.39.644c.452.219.906.444 1.365.685l-4.702 13.03s-3.913-4.371-14.405-4.903c-9.974-.504-15.522 10.285 3.931 16.647a47.04 47.04 0 0 1 1.67.585c.15.055.298.112.446.169c.42.16.832.324 1.234.492c.123.051.248.101.369.153a39.35 39.35 0 0 1 1.484.671c.062.029.12.06.181.09a37.569 37.569 0 0 1 1.574.813c.379.208.75.42 1.109.636c.064.038.131.076.194.115c.424.259.833.523 1.229.792l.202.141c.327.226.645.456.953.689c.084.064.169.127.252.192c.357.276.706.556 1.038.841c5.631 4.835 7.638 10.851 7.476 16.562c-.128 4.513-1.707 8.184-3.538 10.953l11.048-2.641a5.638 5.638 0 0 0 4.169-6.793"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialSnapchat(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M71.925 43.847c.846-.21 1.651-.773 2.481-.789c1.279-.032 2.724-.045 3.813.498c1.89.943 2.04 2.999.312 4.25c-1.384.992-3.052 1.647-4.667 2.267c-3.594 1.388-4.072 2.433-2.255 5.837c3.012 5.639 7.229 9.836 13.681 11.39c.668.158 1.741.895 1.708 1.283c-.073.882-.506 2.178-1.174 2.485c-2.17 1.016-4.469 1.878-6.804 2.38c-1.457.3-2.004.838-2.311 2.21c-.656 2.894-1.231 3.254-4.036 2.692c-4.461-.899-8.383-.125-12.184 2.619c-7.67 5.549-13.386 5.525-21.04-.032c-3.732-2.72-7.598-3.481-11.993-2.578c-2.947.603-3.453.247-4.157-2.752c-.291-1.226-.749-1.858-2.178-2.129a27.26 27.26 0 0 1-6.476-2.105c-.826-.38-1.481-1.692-1.643-2.659c-.065-.397 1.316-1.307 2.161-1.542c6.946-1.898 11.176-6.65 14.021-12.949c.672-1.49.105-2.546-1.129-3.238c-1.101-.607-2.348-.951-3.522-1.441c-.741-.312-1.498-.623-2.178-1.044c-1.384-.854-2.546-1.915-1.745-3.752c.652-1.51 2.801-2.311 4.493-1.769c.959.304 1.898.7 2.874.915c1.279.275 2-.089 1.951-1.656c-.113-3.558-.316-7.144-.028-10.686c.607-7.484 4.983-12.423 11.718-15.05c9.735-3.781 21.36-.927 26.61 8.492c1.842 3.303 2.121 6.877 2.275 10.524c-.146 2.287-.263 4.57-.445 6.845c-.112 1.484.665 1.783 1.867 1.484"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialSpotify(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M49.999 10.999C28.461 10.999 11 28.46 11 49.999c0 10.158 3.886 19.408 10.249 26.346a30.148 30.148 0 0 1 9.011-6.986a29.96 29.96 0 0 1 13.665-3.262a30.187 30.187 0 0 1 20.511 8.014a30.194 30.194 0 0 1 6.558 8.75C81.818 75.929 89 63.808 89 49.999c0-21.539-17.461-39-39.001-39m18.468 54.147a2.598 2.598 0 0 1-3.618.654a36.167 36.167 0 0 0-20.71-6.475c-4.305 0-8.518.744-12.52 2.211a2.597 2.597 0 0 1-3.336-1.547a2.6 2.6 0 0 1 1.547-3.336a41.438 41.438 0 0 1 14.309-2.525a41.364 41.364 0 0 1 23.673 7.398a2.602 2.602 0 0 1 .655 3.62m5.673-11.993a2.92 2.92 0 0 1-4.03.93a48.95 48.95 0 0 0-25.971-7.432c-5.119 0-10.167.787-15.005 2.342a2.927 2.927 0 0 1-1.787-5.574A54.95 54.95 0 0 1 44.14 40.8a54.786 54.786 0 0 1 29.07 8.322a2.925 2.925 0 0 1 .93 4.031m2.864-10.402c-.56 0-1.126-.145-1.643-.449a61.707 61.707 0 0 0-31.223-8.469c-5.94 0-11.821.848-17.481 2.516a3.25 3.25 0 1 1-1.837-6.236a68.37 68.37 0 0 1 19.318-2.779c12.149 0 24.085 3.238 34.514 9.365a3.25 3.25 0 0 1-1.648 6.052M54.908 84.444a16.157 16.157 0 0 0-10.983-4.293c-2.569 0-5.029.586-7.312 1.746a16.116 16.116 0 0 0-4.085 2.969a38.83 38.83 0 0 0 17.471 4.135a39.02 39.02 0 0 0 7.907-.805a16.136 16.136 0 0 0-2.998-3.752"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialSquidoo(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M90.624 66.304c-2.183-7.806-8.903-13.12-16.914-13.415c-1.315-.048-2.246-.641-2.768-1.781c-.873-1.911-.825-3.857.09-5.752c.855-1.768 1.788-3.498 2.615-5.279c1.093-2.353 1.786-4.754 1.303-7.436c-.572-3.186-2.348-5.565-4.615-7.671c-3.786-3.518-8.419-5.38-13.33-6.632c-3.276-.829-5.171-2.894-5.898-6.129c-.204-.906-.38-1.816-.571-2.727c-.331.528-.421 1.053-.479 1.581c-.468 4.473 1.55 7.812 5.688 9.432c.487.188.979.361 1.458.567c3.119 1.337 6.103 2.878 8.448 5.421c3.19 3.463 3.458 6.984.775 10.809c-1.561 2.225-3.585 4.088-5.046 6.621c-.067-.357-.093-.508-.122-.656c-.631-3.022-1.328-6.016-3.086-8.656c-3.785-5.672-8.855-9.774-15.038-12.582c-2.406-1.092-4.819-2.156-6.77-4.068c-3.494-3.422-4.804-7.583-4.519-12.355h-.482c.011.157.052.318.026.473c-1.328 7.803.982 14.318 6.875 19.602c2.26 2.022 4.78 3.652 7.396 5.155c3.435 1.975 5.693 4.819 6.158 8.839c.244 2.111.253 4.246.408 6.368c.059.78-.424.508-.768.456c-4.072-.635-7.94-1.886-11.692-3.598c-3.792-1.73-6.243-4.537-7.444-8.492c-.233-.766-.485-1.531-.714-2.298c-.816-2.724-1.428-5.518-2.653-8.107c-1.139-2.408-2.727-4.489-5.382-5.188c-5.015-1.32-9.939-.958-14.629 1.432c-.24.119-.544.219-.412.602c.224.195.454.051.682.012c2.507-.437 4.997-.496 7.539-.085c5.069.816 8.188 3.558 9.154 8.667c.321 1.692.834 3.354 1.12 5.051c1.537 9.107 6.444 15.828 14.341 20.472c1.579.927 3.131 1.905 4.907 2.989c-2.072.373-3.938.721-5.808 1.049c-3.875.676-7.755 1.319-11.715 1.094c-2.149-.118-4.088-.591-4.943-2.9c-.251-.674-.492-1.359-.655-2.057c-.987-4.239-.801-8.553-.628-12.838c.187-4.564-2.917-8.449-7.45-8.964c-.219-.023-.536-.143-.687.122c2.471 1.934 4.377 4.141 4.104 7.482c-.138 1.674-.469 3.336-.768 4.994c-.735 4.078-1.42 8.158-.813 12.322c.587 4.023 2.304 7.344 6.01 9.349c2.296 1.245 4.777 1.817 7.391 1.899c4.685.146 9.281-.672 13.876-1.352c4.19-.619 6.912 1.089 7.659 5.25c.14.788.235 1.59.284 2.389c.3 4.939 1.989 9.313 5.376 12.953c4.094 4.395 9.272 6.291 15.128 5.488c9.154-1.254 14.866-6.853 17.803-15.406c1.414-4.113 1.353-8.34.185-12.516"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialStackOverflow(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M65.269 85.165H23.208l.221-28.143l-6.32-.06L17.074 92h54.853V57.126h-6.658z"/><path d="M27.414 73.601h32.947v7.008H27.414zm.773-12.617l32.97 3.183l-.69 7.153l-32.97-3.183zm3.016-14.269l31.877 8.996l-1.951 6.914l-31.877-8.996zm7.729-16.124l28.352 17.124l-3.715 6.151l-28.352-17.124zM68.509 46.43l-19.29-26.924l5.84-4.185l19.29 26.925zm1.869-37.248l7.086-1.185l5.462 32.67l-7.086 1.184z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialSteam(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M92.43 40.935a7.04 7.04 0 0 1-14.08 0a7.043 7.043 0 0 1 7.039-7.039a7.044 7.044 0 0 1 7.041 7.039m-7.076-13.104c-7.218 0-13.08 5.822-13.145 13.025l-8.19 11.736a9.715 9.715 0 0 0-6.024 1.326L20.901 39.001c-.97-4.4-4.903-7.719-9.586-7.719c-5.406 0-9.815 4.424-9.815 9.828c0 5.41 4.409 9.816 9.815 9.816c1.83 0 3.541-.504 5.009-1.379l37.094 14.889c.959 4.412 4.893 7.733 9.584 7.733c5.083 0 9.275-3.896 9.762-8.858l12.589-9.201c7.258 0 13.146-5.877 13.146-13.135s-5.887-13.144-13.145-13.144m0 4.332c4.863 0 8.813 3.951 8.813 8.812c0 4.863-3.951 8.801-8.813 8.801c-4.861 0-8.813-3.938-8.813-8.801c0-4.861 3.952-8.812 8.813-8.812m-74.039 1.719a7.198 7.198 0 0 1 6.377 3.832l-3.588-1.436v.016a5.723 5.723 0 0 0-7.256 3.248a5.735 5.735 0 0 0 2.978 7.379v.014l3.046 1.215a7.2 7.2 0 0 1-1.557.178A7.198 7.198 0 0 1 4.1 41.111a7.208 7.208 0 0 1 7.215-7.229m51.687 21.254a7.197 7.197 0 0 1 7.216 7.217a7.197 7.197 0 0 1-7.216 7.215a7.192 7.192 0 0 1-6.389-3.844c1.187.48 2.375.953 3.56 1.436c2.941 1.182 6.292-.242 7.473-3.182c1.183-2.941-.254-6.275-3.196-7.459l-3.004-1.205a7.39 7.39 0 0 1 1.556-.178"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialStumbleupon(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M80.801 26.084A38.805 38.805 0 0 1 89 50c0 21.536-17.464 39-38.999 39c-15.153 0-28.29-8.646-34.746-21.271h16.998c8.1 0 13.758-4.366 13.758-11.345c0-14.791-17.063-8.765-17.063-14.014c0-2.529 2.029-3.449 5.742-3.449h11.83c3.279 0 7.136.588 7.136 4.166v11.787c0 8.535 6.665 12.854 13.573 12.854c7.026 0 13.571-4.319 13.571-12.854v-28.79z"/><path d="M50.001 11a38.817 38.817 0 0 1 21.937 6.756V55.15c0 2.592-2.12 4.713-4.708 4.713c-2.591 0-4.709-2.121-4.709-4.713V43.945c0-8.945-4.527-12.666-10.423-12.666H32.909c-6.947 0-12.71 4.422-12.71 11.434c0 3.588 1.537 8.846 8.806 10.15c5.921 1.066 8.258.756 8.258 3.924c0 3.076-3.446 3.076-7.995 3.076H12.261A38.94 38.94 0 0 1 11 50c0-21.531 17.466-39 39.001-39"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialTreehouse(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M49.165 39.829c-4.657 3.569-9.147 7.375-14.374 10.134c-.805.424-1.617.818-2.504 1.064a5.438 5.438 0 0 0-3.995 5.316c.027 2.404 1.728 4.602 4.035 5.211c2.397.632 5.021-.336 6.117-2.566c1.441-2.941 3.917-4.77 6.447-6.576c.98-.698 1.347-.599 2.136.34c.879-1.788 1.739-3.531 2.602-5.285c.268.158.206.327.168.455c-1.433 4.971-2.683 9.996-4.333 14.9c-.362 1.081-.887 1.932-1.946 2.604c-2.8 1.774-3.509 5.462-1.78 8.305c1.687 2.773 5.243 3.746 8.16 2.229c2.886-1.498 4.159-5.016 2.731-7.971c-.511-1.059-.546-2.004-.267-3.055c.733-2.781 1.176-5.658 2.691-8.183c.67-1.116 1.462-2.126 2.563-2.85c1.274-.842 2.244-.669 3.166.536c1.108 1.451 1.659 3.041 1.506 4.908c-.209 2.553 1.348 4.816 3.712 5.617a5.474 5.474 0 0 0 6.253-2.044c1.464-2.086 1.402-4.719-.332-6.697c-1.985-2.269-4.128-4.39-5.936-6.815c-3.092-4.147-3.367-8.433-.753-12.913c1.425-2.437 2.8-4.904 4.218-7.347c1.265-2.175 2.849-4.059 5.091-5.302c1.893-1.051 3.756-1.246 5.676.016c1.229.809 2.566 1.455 3.838 2.201c2.888 1.695 4.422 4.222 4.429 7.576c.021 10.938.021 21.873-.001 32.806c-.008 3.282-1.496 5.829-4.325 7.442A5336.43 5336.43 0 0 1 54.49 90.661c-2.933 1.648-5.955 1.656-8.891.013c-9.965-5.585-19.925-11.18-29.85-16.836c-2.759-1.572-4.228-4.072-4.234-7.291a9862.616 9862.616 0 0 1 0-33.087c.006-3.18 1.436-5.686 4.159-7.238a3074.74 3074.74 0 0 1 30.264-17.06c2.595-1.441 5.385-1.381 8.061-.079c1.551.752 3.02 1.671 4.515 2.532c1.19.682 1.715 1.797 1.872 3.101c.279 2.302-.282 4.427-1.402 6.418c-1.368 2.434-2.693 4.899-4.153 7.278c-2.641 4.301-6.46 6.925-11.48 7.723c-3.206.506-6.39.077-9.571-.28c-2.353-.267-4.523.997-5.311 3.159c-.802 2.211.017 4.577 2.016 5.828c1.977 1.233 4.445.929 6.099-.826c1.041-1.099 2.44-1.039 3.743-1.266c1.015-.178 2.056-.212 3.09-.27c.37-.021.655-.187.957-.345c1.556-.818 3.106-1.635 4.657-2.458c.043.055.089.101.134.152"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialTumblr(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M79 14H21c-3.85 0-7 3.15-7 7v58c0 3.85 3.15 7 7 7h58c3.85 0 7-3.15 7-7V21c0-3.85-3.15-7-7-7M66.715 74.172c-2.319 1.092-4.419 1.861-6.298 2.3c-1.88.44-3.913.659-6.098.659c-2.479 0-4.672-.313-6.579-.939c-1.906-.627-3.532-1.522-4.878-2.678c-1.347-1.161-2.28-2.394-2.799-3.701c-.52-1.305-.78-3.198-.78-5.678V45.102h-5.998v-7.678c2.131-.693 3.959-1.685 5.478-2.979a14.873 14.873 0 0 0 3.659-4.659c.921-1.811 1.553-4.117 1.899-6.916h7.718v13.715h12.877v8.518H52.039v13.916c0 3.145.166 5.164.5 6.059c.332.893.953 1.606 1.858 2.139c1.2.719 2.573 1.08 4.119 1.08c2.746 0 5.479-.894 8.198-2.681v8.556z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialTwitter(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M88.5 26.12a31.562 31.562 0 0 1-9.073 2.486a15.841 15.841 0 0 0 6.945-8.738A31.583 31.583 0 0 1 76.341 23.7a15.783 15.783 0 0 0-11.531-4.988c-8.724 0-15.798 7.072-15.798 15.798c0 1.237.14 2.444.41 3.601c-13.13-.659-24.77-6.949-32.562-16.508a15.73 15.73 0 0 0-2.139 7.943a15.791 15.791 0 0 0 7.028 13.149a15.762 15.762 0 0 1-7.155-1.976c-.002.066-.002.131-.002.199c0 7.652 5.445 14.037 12.671 15.49a15.892 15.892 0 0 1-7.134.27c2.01 6.275 7.844 10.844 14.757 10.972a31.704 31.704 0 0 1-19.62 6.763c-1.275 0-2.532-.074-3.769-.221a44.715 44.715 0 0 0 24.216 7.096c29.058 0 44.948-24.071 44.948-44.945c0-.684-.016-1.367-.046-2.046A32.03 32.03 0 0 0 88.5 26.12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialVimeo(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M79.048 14H20.952C17.128 14 14 17.128 14 20.952v58.096C14 82.872 17.128 86 20.952 86h58.096C82.872 86 86 82.872 86 79.048V20.952C86 17.128 82.872 14 79.048 14M74.41 36.852c-.565 2.938-4.314 13.219-11.107 22.297c-2.97 3.836-6.442 7.604-10.897 10.206c-1.643.96-4.893 2.172-7.052 1.266c-1.512-.629-2.672-1.955-3.601-3.617c-1.672-2.998-2.577-7.097-3.362-10.206c-.725-2.851-1.405-5.558-2.291-8.036c-.516-1.455-.83-2.927-1.43-4.325l-.009-.014l.007.014l-.068-.104c.02.03.041.059.061.091a9.553 9.553 0 0 0-.577-1.148c-1.7-2.828-4.133-.11-6.686-.481c-.689-.1-1.211-.595-1.549-1.121c-.194-.301.151-.612-.849-.864v-.275c4-2.916 7.17-6.455 11.46-8.939c1.345-.782 3.736-1.668 5.591-1.129c1.46.422 2.791 1.779 3.508 3.752h.002c.131 0 .247.651.34.942c.715 2.214 1.099 5.043 1.463 7.194c.577 3.387.481 8.631 2.415 11.59h.002c.255 0 .544.686.875 1.038c3.169 3.345 5.587-1.36 6.878-3.724c1.116-2.052 2.115-4.219 2.288-6.172c.099-1.107.045-2.004-.15-2.699c-.553-1.973-2.255-2.446-4.921-2.355c-.34.012-.694.027-1.064.056c2.424-7.427 9.553-10.53 11.658-10.848c2.593-.39 6.372-.312 8.149 1.907a4.8 4.8 0 0 1 .79 1.436c.456 1.313.391 2.805.126 4.268"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialWindows(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M85.905 47.797V13l-39.97 5.83v28.967zM43.121 19.241l-29.026 4.234v24.322h29.026zM14.095 51.774v24.632l29.026 4.283V51.774zm31.84 29.331L85.905 87V51.774h-39.97z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialXbox(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M39.749 33.478c-11.464-7.078-18.27-10.103-18.27-10.103A38.952 38.952 0 0 0 11 50.019c0 9.229 3.19 17.695 8.522 24.373c0 .001-5.349-14.369 20.227-40.914m10.222-7.543l.032.018a1.098 1.098 0 0 0-.026-.018c15.609-11.279 25.305-5.611 25.305-5.611c-6.803-5.846-15.642-9.381-25.305-9.381c-9.659 0-18.498 3.535-25.301 9.381c-.001 0 9.695-5.801 25.295 5.611m-.125 16.012C28.15 58.006 21.824 76.959 21.824 76.959c7.083 7.455 17.08 12.098 28.152 12.098c11.101 0 21.114-4.658 28.199-12.141c0 0-4.12-17.047-28.329-34.969M89 50.019a38.954 38.954 0 0 0-10.479-26.644s-6.806 3.025-18.27 10.103c25.576 26.545 20.227 40.914 20.227 40.914A38.91 38.91 0 0 0 89 50.019"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialYahoo(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m89.782 77.807l-10.215-1.248l1.136-8.963l10.215 1.246zM100 23.438l-17.481-2.155l.112 40.522l7.607 1.019zm-35.187 7.039c-7.038-5.336-16.118-7.941-27.13-7.941c-11.123 0-20.203 2.605-27.126 7.941C3.519 35.695 0 42.396 0 50.68c0 8.4 3.519 15.213 10.557 20.435c5.742 4.33 12.97 6.865 21.615 7.602h10.981c8.582-.736 15.822-3.271 21.661-7.602c6.923-5.223 10.441-12.035 10.441-20.435c-.001-8.284-3.519-14.985-10.442-20.203m-5.222 13.279l-5.335 1.137c-.793.225-2.951 1.816-6.696 4.881c-3.974 3.178-6.017 5.221-6.356 6.242l-.227 2.84l-.114 1.703l.455 4.312l7.038.109l-.115 2.158H36.776l-10.216.114l.341-2.043l3.292-.109c1.703-.119 2.726-.344 3.064-.686c.228-.227.341-1.59.341-3.973v-1.475l-.113-2.951c-.228-.678-2.044-3.178-5.562-7.492c-3.632-4.426-6.015-7.037-7.151-7.947l-6.583-.904V37.74c.341-.225 3.973-.225 10.669-.113c4.654-.111 8.513-.111 11.579.113l-.228 1.592l-6.925.678c2.157 3.18 5.45 7.605 9.876 13.17c5.789-4.43 8.854-7.268 8.966-8.518l-5.901-.906l-.453-2.041h9.874l9.308.113z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialYelp(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M57.852 61.469c-1.4 1.409-.215 3.981-.215 3.981L68.18 83.056s1.732 2.32 3.231 2.32c1.506 0 2.996-1.236 2.996-1.236l8.335-11.916s.839-1.504.859-2.819c.03-1.871-2.793-2.388-2.793-2.388l-19.737-6.338c.001 0-1.932-.513-3.219.79m-1-8.875c1.011 1.71 3.795 1.212 3.795 1.212l19.691-5.756s2.683-1.09 3.067-2.543c.375-1.459-.444-3.213-.444-3.213l-9.409-11.085s-.816-1.403-2.508-1.543c-1.865-.157-3.015 2.099-3.015 2.099L56.904 49.272s-.983 1.747-.052 3.322m-9.305-6.827c2.32-.572 2.688-3.94 2.688-3.94l-.158-28.035s-.349-3.457-1.903-4.397c-2.438-1.477-3.16-.707-3.857-.602l-16.348 6.074s-1.602.529-2.435 1.865c-1.191 1.891 1.209 4.663 1.209 4.663L43.736 44.56s1.676 1.736 3.811 1.207M43.51 57.112c.06-2.162-2.596-3.461-2.596-3.461l-17.57-8.88s-2.605-1.072-3.869-.324c-.969.573-1.828 1.609-1.911 2.524L16.419 61.06s-.171 2.443.461 3.552c.896 1.573 3.845.479 3.845.479l20.514-4.535c.797-.536 2.195-.586 2.271-3.444m5.102 7.602c-1.761-.904-3.868.971-3.868.971l-13.736 15.12s-1.714 2.313-1.279 3.732c.412 1.333 1.091 1.995 2.053 2.46l13.794 4.354s1.673.347 2.94-.019c1.797-.523 1.466-3.334 1.466-3.334l.311-20.479s-.069-1.971-1.681-2.805"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialYoutube(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M45.19 70.391c-.696.984-1.355 1.469-1.985 1.469c-.42 0-.659-.246-.735-.734c-.03-.101-.03-.484-.03-1.22v-12.76h-3.138V70.85c0 1.225.104 2.051.277 2.576c.313.878 1.011 1.289 2.023 1.289c1.154 0 2.34-.697 3.587-2.129v1.891h3.143v-17.33H45.19zM27.897 54.492h3.694v19.985h3.492V54.492h3.762v-3.274H27.897zM49.281 39.34c1.023 0 1.513-.81 1.513-2.431v-7.368c0-1.619-.49-2.426-1.513-2.426c-1.024 0-1.515.807-1.515 2.426v7.368c.001 1.62.491 2.431 1.515 2.431m8.153 17.595c-1.154 0-2.23.632-3.243 1.88v-7.597H51.05v23.259h3.141v-1.682c1.048 1.299 2.127 1.92 3.243 1.92c1.251 0 2.091-.658 2.511-1.947c.209-.735.316-1.889.316-3.492v-6.901c0-1.64-.107-2.786-.316-3.483c-.421-1.299-1.26-1.957-2.511-1.957m-.313 12.587c0 1.565-.461 2.337-1.366 2.337c-.516 0-1.04-.246-1.565-.771V60.524c.525-.517 1.048-.763 1.565-.763c.906 0 1.366.802 1.366 2.365zm10.314-12.587c-1.602 0-2.854.591-3.76 1.778c-.668.875-.971 2.229-.971 4.08v6.072c0 1.84.342 3.209 1.009 4.074c.907 1.183 2.157 1.775 3.801 1.775c1.64 0 2.927-.621 3.797-1.877a4.51 4.51 0 0 0 .733-1.883c.028-.315.067-1.013.067-2.023v-.46h-3.205c0 1.26-.039 1.957-.068 2.128c-.179.839-.631 1.259-1.403 1.259c-1.078 0-1.602-.8-1.602-2.404v-3.072h6.279v-3.59c0-1.852-.315-3.205-.972-4.08c-.879-1.187-2.129-1.777-3.705-1.777m1.537 6.834h-3.139v-1.603c0-1.602.524-2.404 1.574-2.404c1.039 0 1.565.802 1.565 2.404z"/><path d="M79 14H21c-3.85 0-7 3.15-7 7v58c0 3.85 3.15 7 7 7h58c3.85 0 7-3.15 7-7V21c0-3.85-3.15-7-7-7M56.58 24.47h3.163v12.892c0 .746 0 1.133.039 1.235c.07.492.318.743.744.743c.634 0 1.301-.492 2.004-1.486V24.47h3.174v17.51H62.53v-1.91c-1.263 1.449-2.468 2.153-3.625 2.153c-1.022 0-1.727-.415-2.044-1.303c-.175-.529-.281-1.368-.281-2.604zm-11.987 5.708c0-1.871.319-3.24.994-4.125c.878-1.196 2.114-1.795 3.694-1.795c1.59 0 2.824.599 3.703 1.795c.664.885.985 2.254.985 4.125v6.135c0 1.86-.32 3.24-.985 4.119c-.879 1.193-2.112 1.792-3.703 1.792c-1.581 0-2.817-.599-3.694-1.792c-.675-.879-.994-2.259-.994-4.119zm-8.486-11.696l2.499 9.227l2.399-9.227h3.558l-4.232 13.982v9.516h-3.52v-9.516c-.318-1.688-1.022-4.158-2.15-7.435c-.742-2.18-1.515-4.37-2.256-6.547zm38.966 57.036c-.634 2.763-2.896 4.801-5.616 5.104c-6.444.72-12.965.724-19.457.72c-6.493.004-13.014 0-19.457-.72c-2.721-.304-4.979-2.342-5.614-5.104c-.905-3.936-.905-8.232-.905-12.282c0-4.054.01-8.347.916-12.284c.634-2.762 2.893-4.799 5.614-5.103c6.443-.722 12.964-.724 19.457-.722c6.492-.002 13.013 0 19.457.722c2.719.304 4.98 2.341 5.615 5.103c.905 3.938.9 8.23.9 12.284c.001 4.05-.004 8.346-.91 12.282"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialZerply(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M37.483 61.48c7.221.67 12.822 4.535 18.604 8.086c4.091 2.512 8.593 3.973 13.483 3.91c2.35-.029 4.435-.799 6.224-2.303c.767-.646 1.085-.514 1.312.43c1.756 7.254-4.007 14.668-11.467 14.891c-3.72.111-6.915-1.223-9.895-3.17c-3.4-2.225-6.648-4.686-10.036-6.928c-6.246-4.137-13.147-5.551-20.546-4.674c-2.162.258-2.375.086-2.547-2.199c-.321-4.287 1.237-7.81 4.117-10.93c8.217-8.893 16.293-17.912 23.846-27.389c.647-.812 1.289-1.629 1.92-2.453c.181-.238.557-.445.377-.795c-.175-.34-.564-.185-.859-.17c-3.609.201-7.217.402-10.834.322c-1.868-.045-3.691-.299-5.413-1.059c-2.893-1.271-3.835-3.455-2.733-6.391c.896-2.383 2.238-4.523 3.709-6.594c.416-.584.778-.727 1.449-.381c3.368 1.727 7.045 2.184 10.744 2.477c7.833.621 15.604.076 23.297-1.486c1.005-.203 1.029.293 1.118.982c.456 3.51-.722 6.443-3.028 9.047c-8.268 9.332-16.146 19.008-24.974 27.846c-2.735 2.736-5.469 5.509-7.868 8.931"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialZurb(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M79 14H21a7 7 0 0 0-7 7v58a7 7 0 0 0 7 7h58a7 7 0 0 0 7-7V21a7 7 0 0 0-7-7m-8.647 43.23v5.339a2.257 2.257 0 0 1-1.973 2.019l-30.145.002c-2.232 0-4.135-.72-5.707-2.157c-1.571-1.438-2.357-3.177-2.357-5.221c0-1.626.531-3.071 1.594-4.345c1.064-1.27 2.411-2.154 4.041-2.656l26.341-7.527H30.172v-5.096a2.263 2.263 0 0 1 1.779-2.177h30.338c2.232 0 4.133.714 5.705 2.139c1.572 1.427 2.359 3.13 2.359 5.112c0 1.655-.536 3.111-1.607 4.371c-1.071 1.262-2.449 2.141-4.134 2.635L38.235 57.23z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sound(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M45.697 45.697a6.083 6.083 0 0 0-.002 8.603a6.082 6.082 0 0 0 8.606.001a6.08 6.08 0 0 0 0-8.607a6.08 6.08 0 0 0-8.604.003m30.852 33.468c15.972-16.109 15.934-42.207-.122-58.263c-.023-.023-.05-.037-.073-.059l.006-.006l-1.696-1.698l-.02.02a2.38 2.38 0 0 0-3.136.141l-.003-.003l-.026.026l-.008.007l-.006.008l-2.773 2.772l-.002.002a.002.002 0 0 1-.002.002l-.2.2l.02.02a2.376 2.376 0 0 0 .01 2.951l-.019.019l.19.19v.001h.001L70.196 27l.027-.027c.022.023.037.05.06.073C82.95 39.714 82.99 60.3 70.405 73.02l-.017-.017l-1.504 1.504l-.003.002l-.002.003l-.188.188l.019.019a2.38 2.38 0 0 0 .141 3.136l-.003.003l.031.031l.002.003l.003.002l1.396 1.396l.002.003l.003.002l1.376 1.376l.002.003l.003.002l.198.198l.02-.02a2.377 2.377 0 0 0 2.952-.009l.019.019l1.567-1.568a.018.018 0 0 1 .005-.004l.018-.019l.107-.107z"/><path fill="currentColor" d="M64.923 67.54c9.561-9.699 9.523-25.365-.123-35.01c-.023-.023-.05-.037-.073-.06l.007-.007l-1.697-1.698l-.02.02a2.382 2.382 0 0 0-3.136.141l-.003-.003l-.029.029l-.005.004l-.004.005l-2.774 2.774l-.004.003l-.003.004l-.198.198l.02.02a2.376 2.376 0 0 0 .009 2.951l-.019.019l.189.189l.002.002l.002.002l1.504 1.505l.027-.027c.022.023.037.05.06.073c6.258 6.257 6.293 16.407.119 22.717l-.013-.013l-1.505 1.505l-.002.001l-.001.002l-.189.189l.019.019a2.38 2.38 0 0 0 .141 3.135l-.004.004l2.816 2.815l.201.201l.02-.02a2.378 2.378 0 0 0 2.951-.009l.02.02l1.572-1.572l.125-.125z"/><g fill="currentColor"><path d="M54.305 45.7a6.083 6.083 0 0 0-8.606-.001a6.08 6.08 0 0 0 0 8.605a6.08 6.08 0 0 0 8.605-.001a6.084 6.084 0 0 0 .001-8.603"/><path d="m43.109 63.089l.019-.019l-.188-.188l-.003-.004l-.003-.003l-1.503-1.504l-.027.027c-.022-.023-.037-.05-.059-.072c-6.258-6.258-6.293-16.408-.119-22.718l.013.013l1.697-1.696l-.02-.02a2.38 2.38 0 0 0-.141-3.135l.004-.004l-3.018-3.017l-.02.02a2.376 2.376 0 0 0-2.951.009l-.019-.019l-.191.191l-1.381 1.381l-.125.125l.003.003c-9.562 9.699-9.523 25.365.123 35.011c.022.022.049.037.072.059l-.006.006l1.697 1.698l.02-.02a2.382 2.382 0 0 0 3.135-.141l.003.003l.029-.029l.005-.004l.004-.005l2.775-2.775l.003-.002l.002-.002l.199-.199l-.02-.02a2.374 2.374 0 0 0-.009-2.95"/><path d="m31.483 74.715l.019-.019l-.19-.19l-.001-.001l-.001-.001l-1.506-1.505l-.027.027c-.022-.023-.037-.05-.059-.073C17.05 60.284 17.012 39.7 29.597 26.98l.016.016l1.504-1.504l.003-.002l.002-.003l.188-.188l-.019-.019a2.38 2.38 0 0 0-.141-3.136l.004-.004l-1.434-1.434l-.001-.001l-.001-.001l-1.581-1.581l-.02.021a2.376 2.376 0 0 0-2.951.009l-.02-.02l-1.696 1.697l.003.003c-15.974 16.11-15.936 42.209.121 58.265c.023.023.05.037.073.059l-.007.007l1.697 1.698l.02-.02a2.382 2.382 0 0 0 3.136-.142l.003.003l.033-.033l2.778-2.779l.005-.004l.004-.005l.196-.196l-.02-.02a2.376 2.376 0 0 0-.009-2.951"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Star(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M91.532 39.844a1.997 1.997 0 0 0-1.888-1.343H61.482l-9.597-27.159a2 2 0 0 0-3.771 0l-9.598 27.159H10.357c-.851 0-1.609.539-1.891 1.343a2.003 2.003 0 0 0 .651 2.226l21.986 17.409l-9.84 27.846a2.003 2.003 0 0 0 2.917 2.382l25.818-15.488l25.818 15.488a2.003 2.003 0 0 0 2.918-2.382l-9.84-27.846L90.886 42.07a1.999 1.999 0 0 0 .646-2.226"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stop(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M80.049 22.127a2.162 2.162 0 0 0-2.163-2.156c-.028 0-.054.007-.082.008H22.302c-.063-.006-.125-.019-.19-.019a2.163 2.163 0 0 0-2.163 2.163c0 .101.016.198.03.295V77.87h.001l-.001.009c0 1.194.969 2.162 2.163 2.162c.046 0 .089-.011.134-.013v.002h55.688v-.018a2.158 2.158 0 0 0 2.084-2.142h.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Strikethrough(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M77.5 49.719c-.035 0-.068.009-.103.01v-.01h-55v.01a2.995 2.995 0 0 0-2.897 2.99a2.996 2.996 0 0 0 2.897 2.99v.01h55v-.01c.035.001.068.01.103.01a3 3 0 1 0 0-6m-6.572 9.75h-7.14v.011c-.035-.002-.069-.011-.105-.011c-.863 0-1.562.699-1.562 1.562c0 .126.019.247.047.365h-.018c.092.393.157.802.157 1.249c0 3.819-3.14 7.808-11.288 7.808c-7.741 0-13.842-3.592-17.678-7.653a1.555 1.555 0 0 0-1.237-.617a1.56 1.56 0 0 0-1.275.664l-.001-.002l-.01.015a.48.48 0 0 0-.019.029l-3.425 5.212l.003.001a1.55 1.55 0 0 0-.398 1.033c0 .515.253.969.637 1.253c5.091 5.205 12.61 8.891 22.978 8.891c15.191 0 21.896-8.147 21.896-17.568c0-.172-.011-.335-.018-.501c.007-.06.018-.118.018-.179c0-.863-.699-1.562-1.562-1.562m-39.06-13.792c.269.471.77.792 1.351.792h23.542a.502.502 0 0 0 0-1.004v-.008c-8.471-2.48-17.2-3.403-17.2-8.866c0-4.159 3.734-7.044 9.505-7.044c5.941 0 11.967 2.037 16.465 6.535l.006-.008c.272.231.62.375 1.005.375c.491 0 .923-.231 1.21-.584l.002.003l.028-.039c.028-.037.056-.074.081-.114l3.409-4.788l-.003-.001a1.55 1.55 0 0 0 .398-1.033c0-.473-.215-.892-.547-1.178l.011-.015C65.956 23.606 58.742 20.72 50 20.72c-12.476 0-20.623 7.214-20.623 16.634c0 3.499.939 6.195 2.491 8.323"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Subscript(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M90.694 76.981a.99.99 0 0 0-.703-.291v-.012h-5.208c3.816-2.52 5.064-4.32 5.064-6.504c0-3.216-2.712-5.521-6.865-5.521c-2.138 0-4.685.695-6.521 2.481c-.043.031-.092.048-.13.087c-.027.027-.038.063-.062.093c-.018.018-.038.033-.055.052l.012.014a.967.967 0 0 0 .005 1.096l-.016.01l1.367 1.995h.001a.98.98 0 0 0 1.265.092l.007.009c.036-.03.072-.056.108-.085c.007-.007.017-.009.025-.016l.005-.007c1.184-.965 2.446-1.548 4.088-1.548c1.032 0 1.848.48 1.848 1.248c0 1.2-.696 1.8-8.017 6.96v.022c-.067.041-.139.073-.197.131a.97.97 0 0 0-.267.578h-.024v2.609c0 .241.195.437.437.437c.018 0 .034-.008.052-.01v.004h13.082a.992.992 0 0 0 .994-.994v-2.225a1.013 1.013 0 0 0-.295-.705M70.495 33.34l-3.389-3.39a2.604 2.604 0 0 0-3.683 0l-.016-.016l-13.033 13.034L37.34 29.934a2.604 2.604 0 0 0-3.683 0l-3.389 3.389a2.604 2.604 0 0 0 0 3.683l-.016.016l13.033 13.033L30.252 63.09l.016.016a2.604 2.604 0 0 0 0 3.683l3.389 3.389a2.604 2.604 0 0 0 3.683 0l13.033-13.033l13.033 13.033l.016-.016a2.604 2.604 0 0 0 3.683 0l3.389-3.389a2.604 2.604 0 0 0 0-3.683L57.462 50.056l13.033-13.033a2.603 2.603 0 0 0 0-3.683"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Superscript(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M90.478 31.418a.99.99 0 0 0-.703-.291v-.011h-5.208c3.816-2.52 5.064-4.32 5.064-6.504c0-3.216-2.712-5.521-6.865-5.521c-2.138 0-4.685.695-6.521 2.481c-.043.031-.092.048-.13.087c-.027.027-.038.063-.062.093c-.018.018-.038.033-.055.052l.012.014a.965.965 0 0 0 .005 1.095l-.016.01l1.367 1.995h.001a.982.982 0 0 0 1.265.093l.007.009c.036-.03.072-.056.108-.085c.007-.007.017-.009.025-.016l.005-.007c1.184-.965 2.446-1.548 4.088-1.548c1.032 0 1.848.48 1.848 1.248c0 1.2-.696 1.8-8.017 6.96v.022c-.067.041-.139.073-.197.131a.97.97 0 0 0-.267.578h-.024v2.609c0 .241.195.437.437.437c.018 0 .034-.008.052-.01v.004h13.082a.992.992 0 0 0 .994-.994v-2.225a1.013 1.013 0 0 0-.295-.706m-20.199 2.049l-3.389-3.39a2.604 2.604 0 0 0-3.683 0l-.016-.016l-13.034 13.034l-13.033-13.034a2.604 2.604 0 0 0-3.683 0l-3.389 3.389a2.604 2.604 0 0 0 0 3.683l-.016.016l13.033 13.033l-13.034 13.035l.016.016a2.604 2.604 0 0 0 0 3.683l3.389 3.389a2.604 2.604 0 0 0 3.683 0l13.033-13.033l13.033 13.033l.016-.016a2.604 2.604 0 0 0 3.683 0l3.389-3.389a2.604 2.604 0 0 0 0-3.683L57.246 50.183L70.279 37.15a2.604 2.604 0 0 0 0-3.683"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TabletLandscape(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M92.566 22.081a3.035 3.035 0 0 0-3.035-3.035H10.469a3.035 3.035 0 0 0-3.035 3.035c0 .043.011.084.013.127h-.013v55.838h.013a3.028 3.028 0 0 0 3.022 2.908h79.062a3.028 3.028 0 0 0 3.022-2.908h.013V22.208h-.013c.002-.044.013-.084.013-.127m-9.957 48.916H17.391V29.003H82.61v41.994zm4.978-18.386a2.484 2.484 0 1 1 0-4.968a2.484 2.484 0 0 1 0 4.968"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TabletPortrait(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M77.919 7.434c-.043 0-.084.011-.127.013v-.013H21.954v.013a3.028 3.028 0 0 0-2.908 3.022v79.062a3.028 3.028 0 0 0 2.908 3.022v.013h55.838v-.013c.043.002.083.013.127.013a3.035 3.035 0 0 0 3.035-3.035V10.469a3.035 3.035 0 0 0-3.035-3.035M49.873 90.072a2.485 2.485 0 1 1 .001-4.97a2.485 2.485 0 0 1-.001 4.97m21.124-7.463H29.003V17.391h41.995z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Target(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M49.984 27.452c-12.453 0-22.548 10.095-22.548 22.548c0 12.453 10.095 22.548 22.548 22.548S72.531 62.453 72.531 50c0-12.453-10.095-22.548-22.547-22.548m0 35.156c-6.964 0-12.609-5.645-12.609-12.608c0-6.964 5.645-12.609 12.609-12.609S62.592 43.036 62.592 50c0 6.963-5.645 12.608-12.608 12.608"/><path fill="currentColor" d="M50 12.5c-20.712 0-37.5 16.793-37.5 37.502C12.5 70.712 29.288 87.5 50 87.5s37.5-16.788 37.5-37.498C87.5 29.293 70.712 12.5 50 12.5m0 65.061c-15.199 0-27.56-12.362-27.56-27.559C22.44 34.807 34.802 22.44 50 22.44c15.198 0 27.56 12.367 27.56 27.562c0 15.196-12.362 27.559-27.56 27.559"/><circle cx="49.984" cy="50" r="7.621" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TargetTwo(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M94.988 47.506h-7.58c-1.231-18.719-16.193-33.684-34.914-34.915v-7.58a2.494 2.494 0 0 0-4.987 0v7.58c-18.722 1.231-33.683 16.195-34.914 34.915h-7.58a2.494 2.494 0 0 0 0 4.988h7.579c1.229 18.722 16.191 33.684 34.915 34.914v7.58a2.494 2.494 0 0 0 4.987 0v-7.58c18.723-1.23 33.685-16.193 34.915-34.914h7.579a2.494 2.494 0 0 0 0-4.988M47.506 77.443c-13.212-1.191-23.759-11.739-24.949-24.949h4.99c1.152 10.484 9.475 18.807 19.959 19.959zm0-49.896c-10.484 1.152-18.807 9.475-19.959 19.959h-4.99c1.192-13.208 11.738-23.757 24.948-24.949zm4.988-4.989c13.21 1.191 23.756 11.74 24.948 24.948h-4.989c-1.152-10.485-9.475-18.808-19.959-19.959zm0 54.885v-4.989c10.484-1.152 18.807-9.475 19.959-19.959h4.989c-1.19 13.209-11.737 23.756-24.948 24.948"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Telephone(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M84.96 70.237c-.167-1.032-.814-1.914-1.783-2.438l-14.335-8.446l-.118-.066a4.256 4.256 0 0 0-1.937-.45c-1.201 0-2.348.455-3.144 1.253l-4.231 4.233c-.181.172-.771.421-.95.43c-.049-.004-4.923-.355-13.896-9.329c-8.957-8.955-9.337-13.844-9.34-13.844c.005-.25.251-.838.426-1.02l3.608-3.607c1.271-1.274 1.652-3.386.898-5.022L32.19 16.938c-.579-1.192-1.704-1.928-2.952-1.928c-.883 0-1.735.366-2.401 1.031l-9.835 9.813c-.943.938-1.755 2.578-1.932 3.898c-.086.631-1.831 15.693 18.819 36.346C51.42 83.627 65.09 84.989 68.865 84.989a10.7 10.7 0 0 0 1.376-.071c1.316-.176 2.954-.986 3.891-1.925l9.827-9.826c.802-.806 1.168-1.871 1.001-2.93"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TelephoneAccessible(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.777 37.475c.662.48 1.511.607 2.336.365l12.607-3.262l.104-.023a3.29 3.29 0 0 0 1.321-.83c.665-.661 1.049-1.548 1.047-2.434l-.001-4.684c.006-.193.195-.658.287-.763c.028-.022 2.92-2.528 12.854-2.528c9.913-.002 12.832 2.496 12.834 2.494c.134.141.323.604.328.799l-.001 3.994c.002 1.409.962 2.787 2.281 3.273l12.709 3.893c.981.34 2.012.124 2.704-.57c.489-.486.755-1.156.755-1.893l.014-10.879c.004-1.041-.455-2.398-1.088-3.228c-.302-.397-7.673-9.699-30.535-9.7c-19.405.001-27.724 6.815-29.815 8.906c-.449.449-.68.74-.723.796c-.63.827-1.089 2.183-1.088 3.22v10.883c.002.889.391 1.68 1.07 2.171m4.582 21.101v2.543c0 1.168.95 2.12 2.121 2.12h6.836a2.12 2.12 0 0 0 2.117-2.12v-2.543a2.122 2.122 0 0 0-2.117-2.123H26.48c-1.17 0-2.121.955-2.121 2.123m24.545 2.543v-2.543c0-1.168-.95-2.123-2.119-2.123h-6.838c-1.17 0-2.117.955-2.117 2.123v2.543a2.12 2.12 0 0 0 2.117 2.12h6.838c1.17 0 2.119-.952 2.119-2.12m13.471 0v-2.543c0-1.168-.95-2.123-2.121-2.123h-6.836a2.125 2.125 0 0 0-2.121 2.123v2.543c0 1.168.949 2.12 2.121 2.12h6.836a2.123 2.123 0 0 0 2.121-2.12m13.471 0v-2.543c0-1.168-.95-2.123-2.122-2.123h-6.836a2.126 2.126 0 0 0-2.122 2.123v2.543c0 1.168.95 2.12 2.122 2.12h6.836a2.123 2.123 0 0 0 2.122-2.12m-56.205-10.51h6.84a2.123 2.123 0 0 0 2.117-2.121v-2.543c0-1.166-.95-2.121-2.117-2.121h-6.84a2.124 2.124 0 0 0-2.117 2.121v2.543a2.122 2.122 0 0 0 2.117 2.121m20.306 0a2.123 2.123 0 0 0 2.121-2.121v-2.543c0-1.166-.95-2.121-2.121-2.121h-6.836a2.121 2.121 0 0 0-2.117 2.121v2.543a2.12 2.12 0 0 0 2.117 2.121zm13.471 0a2.124 2.124 0 0 0 2.122-2.121v-2.543c0-1.166-.95-2.121-2.122-2.121h-6.836a2.125 2.125 0 0 0-2.122 2.121v2.543c0 1.168.95 2.121 2.122 2.121zm13.471-6.785h-6.835a2.125 2.125 0 0 0-2.122 2.121v2.543c0 1.168.95 2.121 2.122 2.121h6.835a2.12 2.12 0 0 0 2.117-2.121v-2.543a2.121 2.121 0 0 0-2.117-2.121m13.471 0h-6.84a2.124 2.124 0 0 0-2.117 2.121v2.543c0 1.168.95 2.121 2.117 2.121h6.84a2.123 2.123 0 0 0 2.117-2.121v-2.543a2.125 2.125 0 0 0-2.117-2.121M28.598 73.748v-2.542c0-1.167-.95-2.12-2.117-2.12h-6.84a2.122 2.122 0 0 0-2.117 2.12v2.542c0 1.17.951 2.123 2.117 2.123h6.84a2.123 2.123 0 0 0 2.117-2.123m4.513 2.123h6.836a2.124 2.124 0 0 0 2.121-2.123v-2.542c0-1.167-.95-2.12-2.121-2.12h-6.836a2.12 2.12 0 0 0-2.117 2.12v2.542a2.12 2.12 0 0 0 2.117 2.123m13.472 0h6.836a2.125 2.125 0 0 0 2.122-2.123v-2.542c0-1.167-.95-2.12-2.122-2.12h-6.836a2.123 2.123 0 0 0-2.122 2.12v2.542c0 1.17.95 2.123 2.122 2.123m13.47 0h6.835a2.12 2.12 0 0 0 2.117-2.123v-2.542a2.12 2.12 0 0 0-2.117-2.12h-6.835a2.123 2.123 0 0 0-2.122 2.12v2.542a2.125 2.125 0 0 0 2.122 2.123m20.307-6.785h-6.84a2.122 2.122 0 0 0-2.117 2.12v2.542c0 1.17.95 2.123 2.117 2.123h6.84a2.123 2.123 0 0 0 2.117-2.123v-2.542c0-1.167-.951-2.12-2.117-2.12m-7.933 12.629H26.971a2.124 2.124 0 0 0-2.117 2.121v2.543a2.12 2.12 0 0 0 2.117 2.121h45.456a2.12 2.12 0 0 0 2.118-2.121v-2.543c0-1.166-.95-2.121-2.118-2.121"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextColor(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M84 78.754c-.035 0-.068.009-.104.01v-.01h-68v.01C14.288 78.82 13 80.133 13 81.754s1.288 2.934 2.897 2.99v.01h68v-.01c.035.001.068.01.104.01a3 3 0 1 0-.001-6m-58.629-6.9h5.679A2.596 2.596 0 0 0 33.5 70.1l3.493-9.195h25.97l3.605 9.492l.021-.004a2.59 2.59 0 0 0 2.095 1.437v.024h5.679v-.024c.078.007.153.024.233.024A2.604 2.604 0 0 0 77.2 69.25c0-.266-.051-.517-.125-.759l.012-.002l-.061-.156c-.019-.051-.037-.101-.059-.15L56.876 17.036h-.018a2.596 2.596 0 0 0-2.462-1.79h-8.812a2.596 2.596 0 0 0-2.476 1.836l-20.194 51.3l.012.002c-.096.272-.159.56-.159.865a2.605 2.605 0 0 0 2.604 2.605m24.607-46.678l10.185 26.989h-20.37z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Thumbnails(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M30.278 17.553H19.822a2.27 2.27 0 0 0-2.269 2.269v10.454a2.27 2.27 0 0 0 2.269 2.269h10.455a2.27 2.27 0 0 0 2.269-2.269V19.823a2.269 2.269 0 0 0-2.268-2.27m24.949 0H44.772a2.27 2.27 0 0 0-2.269 2.269v10.454a2.27 2.27 0 0 0 2.269 2.269h10.455a2.27 2.27 0 0 0 2.269-2.269V19.823a2.269 2.269 0 0 0-2.269-2.27m27.219 2.269a2.27 2.27 0 0 0-2.269-2.269H69.721a2.27 2.27 0 0 0-2.269 2.269v10.455a2.27 2.27 0 0 0 2.269 2.27h10.455a2.27 2.27 0 0 0 2.269-2.269zM30.278 42.506H19.822a2.27 2.27 0 0 0-2.269 2.269V55.23a2.27 2.27 0 0 0 2.269 2.269h10.455a2.27 2.27 0 0 0 2.269-2.269V44.776a2.269 2.269 0 0 0-2.268-2.27m24.949 0H44.772a2.27 2.27 0 0 0-2.269 2.269V55.23a2.27 2.27 0 0 0 2.269 2.269h10.455a2.27 2.27 0 0 0 2.269-2.269V44.776a2.269 2.269 0 0 0-2.269-2.27m24.95 0H69.721a2.27 2.27 0 0 0-2.269 2.269V55.23a2.27 2.27 0 0 0 2.269 2.27h10.455a2.27 2.27 0 0 0 2.269-2.269V44.776a2.269 2.269 0 0 0-2.268-2.27M30.278 67.454H19.822a2.27 2.27 0 0 0-2.269 2.269v10.454a2.27 2.27 0 0 0 2.269 2.269h10.455a2.27 2.27 0 0 0 2.269-2.269V69.723a2.268 2.268 0 0 0-2.268-2.269m24.949 0H44.772a2.27 2.27 0 0 0-2.269 2.269v10.454a2.27 2.27 0 0 0 2.269 2.269h10.455a2.27 2.27 0 0 0 2.269-2.269V69.723a2.269 2.269 0 0 0-2.269-2.269m24.95 0H69.721a2.27 2.27 0 0 0-2.269 2.269v10.454a2.27 2.27 0 0 0 2.269 2.27h10.455a2.27 2.27 0 0 0 2.269-2.269V69.723a2.268 2.268 0 0 0-2.268-2.269"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ticket(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M99.483 33.642c-.029-.029-.063-.05-.094-.077L87.158 21.327a6.12 6.12 0 0 1-7.722-.763a6.12 6.12 0 0 1-.761-7.725L66.358.516a1.766 1.766 0 0 0-2.494 0L.52 63.862l-.003.003a1.763 1.763 0 0 0 0 2.494l.001.001l12.325 12.313a6.12 6.12 0 0 1 8.489 8.481l12.166 12.154c.046.06.089.122.144.176a1.762 1.762 0 0 0 2.493 0l.001.001l63.347-63.349a1.763 1.763 0 0 0 0-2.494m-13.965 2.493L36.136 85.518a1.762 1.762 0 0 1-2.493 0l-19.159-19.16a1.762 1.762 0 0 1 0-2.493l49.381-49.382a1.764 1.764 0 0 1 2.493 0l19.16 19.161a1.76 1.76 0 0 1 0 2.491"/><path fill="currentColor" d="m61.923 51.521l-6.821-3.259l.83-7.134a.513.513 0 0 0-.871-.421l-5.097 5.098l-6.653-3.178a.512.512 0 0 0-.683.681l3.179 6.655l-5.098 5.096a.515.515 0 0 0-.099.586a.513.513 0 0 0 .52.284l7.132-.829l3.26 6.822a.514.514 0 0 0 .959-.098l1.87-7.476l7.478-1.871a.52.52 0 0 0 .386-.446a.514.514 0 0 0-.292-.51"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Torso(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m80.161 60.441l-15.66-7.47l-6.622-3.159c2.892-1.822 5.241-4.634 6.778-8.022a21.727 21.727 0 0 0 1.946-8.99c0-1.827-.29-3.562-.694-5.236C63.94 19.453 57.605 13.477 50 13.477c-7.461 0-13.701 5.763-15.792 13.645c-.482 1.808-.815 3.688-.815 5.68c0 3.459.808 6.684 2.181 9.489c1.587 3.254 3.94 5.937 6.804 7.662l-6.342 2.953l-16.168 7.53c-1.404.658-2.327 2.242-2.327 4.011v17.765c0 2.381 1.659 4.312 3.708 4.312h57.505c2.048 0 3.708-1.93 3.708-4.312V64.446c-.002-1.763-.91-3.332-2.301-4.005"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TorsoBusiness(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m80.161 60.441l-15.66-7.47l-6.622-3.159c2.892-1.822 5.241-4.634 6.778-8.021a21.743 21.743 0 0 0 1.945-8.99c0-1.827-.29-3.562-.694-5.236c-1.97-8.112-8.305-14.088-15.91-14.088c-7.461 0-13.7 5.763-15.792 13.644c-.483 1.808-.815 3.688-.815 5.68c0 3.459.808 6.684 2.181 9.489c1.587 3.254 3.94 5.937 6.804 7.662l-6.342 2.953l-16.168 7.53c-1.404.658-2.327 2.242-2.327 4.011v17.765c0 2.381 1.659 4.311 3.707 4.311h24.013V72.92a.78.78 0 0 1 .119-.396l-.01-.006l3.933-6.812l.01.006c.14-.24.389-.41.687-.41c.298 0 .547.169.687.41l.004-.003l.036.063c.005.01.012.018.016.028l3.881 6.721l-.005.003a.783.783 0 0 1 .119.397v13.602h24.013c2.048 0 3.708-1.93 3.708-4.311V64.446c.003-1.763-.905-3.332-2.296-4.005M54.62 55.886l.01.006l-3.934 6.812l-.01-.006a.796.796 0 0 1-.687.409a.796.796 0 0 1-.687-.409l-.005.003l-.04-.069c-.003-.007-.009-.013-.012-.02l-3.881-6.723l.004-.003a.783.783 0 0 1-.119-.397c0-.445.361-.806.806-.806h7.866c.445 0 .806.361.806.806a.762.762 0 0 1-.117.397"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TorsoFemale(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m80.161 60.442l-15.66-7.47l-6.622-3.159c.166-.105.324-.224.487-.335h8.695a2.438 2.438 0 0 0 2.438-2.438v-6.415h-.022c-.266-8.94-3.371-16.805-8.034-21.737c-2.459-2.773-5.646-4.657-9.211-5.22c-.159-.026-.318-.05-.478-.071c-.184-.023-.367-.046-.553-.061a13.701 13.701 0 0 0-1.202-.059h-.002c-.398 0-.791.023-1.183.057c-.188.016-.374.039-.56.062c-.156.02-.311.042-.465.068c-3.536.553-6.701 2.408-9.153 5.141c-4.708 4.927-7.847 12.829-8.115 21.821H30.5v6.415a2.438 2.438 0 0 0 2.438 2.438h8.719c.238.162.475.327.721.475l-6.342 2.953l-16.168 7.53c-1.405.658-2.327 2.242-2.327 4.011v17.765c0 2.381 1.659 4.311 3.708 4.311h57.504c2.049 0 3.708-1.93 3.708-4.311V64.446c-.001-1.763-.909-3.332-2.3-4.004"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Torsos(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M40.348 49.813a13.061 13.061 0 0 0 2.527-2.121c-2.683-4.036-4.239-8.966-4.239-14.191c0-3.492.689-6.798 1.896-9.787c-1.879-1.25-4.026-2.022-6.349-2.022c-7.114 0-12.872 6.702-12.872 14.978c0 5.79 2.83 10.801 6.962 13.295l-17.448 8.121c-1.087.511-1.804 1.739-1.804 3.111v13.771c0 1.846 1.289 3.341 2.874 3.341h11.532v-14.45c0-3.739 2.068-7.121 5.275-8.623z"/><path fill="currentColor" d="M88.772 60.017L73.749 52.85l-6.353-3.03c2.775-1.748 5.028-4.445 6.503-7.695a20.872 20.872 0 0 0 1.866-8.625c0-1.753-.278-3.416-.666-5.023c-1.89-7.782-7.968-13.515-15.263-13.515c-7.158 0-13.144 5.529-15.15 13.09c-.463 1.734-.782 3.538-.782 5.448c0 3.319.775 6.412 2.092 9.103c1.523 3.121 3.78 5.696 6.527 7.35l-6.085 2.834L30.93 60.01c-1.348.631-2.233 2.151-2.233 3.849v17.043c0 2.284 1.591 4.136 3.557 4.136H87.42c1.965 0 3.557-1.852 3.557-4.136V63.859c0-1.691-.871-3.197-2.205-3.842"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TorsosAll(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m96.117 57.583l-16.185-7.719c3.774-2.373 6.338-7.016 6.338-12.354c0-7.754-5.397-14.034-12.058-14.034c-2.289 0-4.403.781-6.223 2.066a24.516 24.516 0 0 1 1.701 8.999c0 4.782-1.411 9.317-3.825 13.066a12.296 12.296 0 0 0 2.809 2.36l-.01.003l10.426 4.971c2.947 1.427 4.853 4.583 4.853 8.045v13.54h11.148c1.489 0 2.694-1.401 2.694-3.13V60.492c0-1.282-.659-2.42-1.668-2.909m-64.55-7.758a12.32 12.32 0 0 0 2.368-1.987c-2.514-3.781-3.972-8.401-3.972-13.297c0-3.271.646-6.37 1.777-9.17c-1.76-1.172-3.773-1.895-5.948-1.895c-6.666 0-12.062 6.28-12.062 14.034c0 5.425 2.651 10.12 6.524 12.457l-16.349 7.61c-1.02.479-1.691 1.629-1.691 2.915v12.903c0 1.729 1.209 3.13 2.694 3.13h10.805v-13.54c0-3.503 1.938-6.672 4.943-8.08z"/><path fill="currentColor" d="m76.94 59.386l-14.077-6.715l-5.952-2.84c2.6-1.637 4.711-4.165 6.093-7.21a19.543 19.543 0 0 0 1.749-8.081c0-1.642-.26-3.201-.623-4.706c-1.771-7.293-7.466-12.665-14.302-12.665c-6.708 0-12.316 5.181-14.196 12.266c-.434 1.625-.733 3.315-.733 5.105c0 3.109.727 6.008 1.961 8.529c1.427 2.925 3.542 5.337 6.116 6.888l-5.702 2.654l-14.534 6.768c-1.262.591-2.092 2.016-2.092 3.606v15.97c0 2.14 1.492 3.875 3.333 3.875h51.691c1.841 0 3.333-1.735 3.333-3.875v-15.97c.001-1.584-.815-2.995-2.065-3.599"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TorsosAllFemale(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M77.067 59.255L62.99 52.54l-5.754-2.745l.301-.208h7.833a2.196 2.196 0 0 0 2.196-2.196v-5.779h-.02c-.228-7.653-2.767-14.425-6.622-18.894c-2.71-3.484-6.605-5.681-10.971-5.681c-4.732 0-8.911 2.586-11.625 6.593c-3.435 4.452-5.664 10.833-5.877 17.982h-.02v5.779c0 1.213.983 2.196 2.196 2.196h7.855c.155.105.312.207.469.309L37.4 52.48l-14.534 6.768c-1.152.54-1.936 1.776-2.065 3.198a4.467 4.467 0 0 0-.046.626v16.005c0 2.146 1.495 3.884 3.34 3.884h51.806c1.846 0 3.34-1.739 3.34-3.884V63.072c0-.805-.212-1.563-.582-2.196c-.359-.711-.907-1.29-1.592-1.621"/><path fill="currentColor" d="M25.646 52.52v-6.081h.027c.243-8.172 2.377-15.627 5.776-21.472c-1.654-1.008-3.513-1.623-5.511-1.623c-6.666 0-12.062 6.28-12.062 14.034c0 5.425 2.651 10.12 6.524 12.457l-16.349 7.61c-1.02.479-1.691 1.629-1.691 2.915v12.903c0 1.729 1.209 3.13 2.694 3.13h10.805v-13.54c0-3.503 1.938-6.672 4.943-8.08zm70.302 5.115L79.6 50.025c3.873-2.337 6.524-7.032 6.524-12.457c0-7.754-5.396-14.034-12.062-14.034c-1.998 0-3.857.615-5.511 1.623c3.399 5.845 5.532 13.3 5.776 21.472h.027v6.081l4.844 2.256c3.005 1.407 4.943 4.576 4.943 8.08v13.54h10.805c1.485 0 2.694-1.401 2.694-3.13V60.55c-.001-1.286-.672-2.436-1.692-2.915"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TorsosFemaleMale(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M34.199 47.04v-6.415h.022c.163-5.466 1.39-10.528 3.382-14.767c.223-.731.474-1.447.756-2.145c-1.879-1.251-4.026-2.022-6.349-2.022c-7.114 0-12.872 6.702-12.872 14.978c0 5.79 2.83 10.801 6.962 13.295L8.652 58.086c-1.088.511-1.804 1.739-1.804 3.111v13.771c0 1.846 1.289 3.341 2.874 3.341h11.532v-14.45c0-3.739 2.068-7.121 5.275-8.623l7.669-3.572V47.04z"/><path fill="currentColor" d="m90.853 60.441l-15.66-7.47l-6.621-3.159c.166-.105.324-.224.487-.335h8.695a2.438 2.438 0 0 0 2.438-2.438v-6.415h-.022c-.266-8.937-3.368-16.798-8.029-21.731c-2.46-2.777-5.65-4.663-9.219-5.226a14.61 14.61 0 0 0-.474-.07c-.185-.023-.369-.046-.555-.062a13.701 13.701 0 0 0-1.202-.059c-.398 0-.793.023-1.186.057c-.187.016-.372.039-.557.062c-.157.02-.313.043-.469.068c-3.532.553-6.694 2.405-9.145 5.134c-4.712 4.927-7.853 12.832-8.121 21.827h-.022v6.415a2.438 2.438 0 0 0 2.438 2.438h8.719c.238.162.475.327.721.475l-6.342 2.953l-16.168 7.53c-1.405.658-2.328 2.242-2.328 4.011v17.765c0 2.381 1.659 4.311 3.708 4.311h57.505c2.049 0 3.708-1.93 3.708-4.311V64.446c0-1.763-.908-3.332-2.299-4.005"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TorsosMaleFemale(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M90.937 63.235c-.185-1.43-.998-2.653-2.166-3.218L73.748 52.85l-6.353-3.03c2.775-1.747 5.028-4.445 6.503-7.695a20.872 20.872 0 0 0 1.866-8.625c0-1.753-.278-3.416-.666-5.023c-1.89-7.782-7.968-13.515-15.263-13.515c-6.494 0-12.011 4.557-14.48 11.047a20.422 20.422 0 0 0-.67 2.043c-.347 1.3-.613 2.64-.725 4.035c-.037.465-.057.936-.057 1.413c0 .415.012.826.036 1.234a20.622 20.622 0 0 0 2.056 7.87c1.523 3.121 3.78 5.695 6.527 7.35l-6.084 2.833l-15.51 7.223c-1.011.474-1.761 1.447-2.074 2.62a4.781 4.781 0 0 0-.159 1.229v17.043c0 2.284 1.591 4.136 3.557 4.136H87.42c1.965 0 3.557-1.852 3.557-4.136V63.859c0-.211-.013-.42-.04-.624"/><path fill="currentColor" d="m39.227 50.335l2.85-1.452c-3.165-4.206-5.065-9.555-5.065-15.383c0-4.041.917-7.85 2.531-11.21c-1.566-.772-3.238-1.212-4.983-1.212c-8.365 0-15.186 9.612-15.544 21.665H19v5.119c0 1.074.871 1.945 1.945 1.945h7.089c.081.051.159.108.24.157l-17.448 8.122c-1.087.511-1.804 1.739-1.804 3.111v13.771c0 1.846 1.289 3.341 2.875 3.341h11.531v-14.45c0-3.739 2.068-7.121 5.275-8.623z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trash(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M75.834 33.388h-51.67a2.372 2.372 0 0 0-2.375 2.373v49.887a2.376 2.376 0 0 0 2.375 2.377h51.67a2.374 2.374 0 0 0 2.375-2.377V35.76a2.37 2.37 0 0 0-2.375-2.372m3.17-16.036H59.402v-2.999a2.373 2.373 0 0 0-2.373-2.377H42.971a2.375 2.375 0 0 0-2.375 2.377v2.999h-19.6a2.372 2.372 0 0 0-2.375 2.373v6.932a2.372 2.372 0 0 0 2.375 2.373h58.008a2.37 2.37 0 0 0 2.375-2.373v-6.932a2.37 2.37 0 0 0-2.375-2.373"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trees(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m91.963 80.982l.023-.013l-7.285-12.617h2.867v-.013c.598 0 1.083-.484 1.083-1.082c0-.185-.059-.351-.14-.503l.019-.011l-6.737-11.669h1.639v-.009a.773.773 0 0 0 .773-.772a.758.758 0 0 0-.1-.359l.013-.008l-9.802-16.979l-.01.006a1.322 1.322 0 0 0-1.186-.754c-.524 0-.968.311-1.185.752l-.005-.003l-9.802 16.978l.002.001a.75.75 0 0 0-.105.366c0 .426.346.772.773.772v.009h1.661l-6.737 11.669l.003.001a1.06 1.06 0 0 0-.147.513c0 .598.485 1.082 1.083 1.082v.013h2.894l-2.1 3.638l-8.399-14.548h4.046v-.018c.844 0 1.528-.685 1.528-1.528c0-.26-.071-.502-.186-.717l.015-.009l-9.507-16.467h2.313v-.012a1.09 1.09 0 0 0 1.091-1.092c0-.186-.059-.353-.141-.506l.019-.011L36.4 13.125l-.005.003a1.873 1.873 0 0 0-1.683-1.06c-.758 0-1.408.452-1.704 1.1L19.201 37.082l.003.002a1.06 1.06 0 0 0-.148.516a1.09 1.09 0 0 0 1.09 1.092v.012h2.345l-9.395 16.272a1.516 1.516 0 0 0-.316.92c0 .844.685 1.528 1.528 1.528v.018h4.084L8.252 75.007c-.24.314-.387.702-.387 1.128c0 1.032.838 1.87 1.871 1.87v.021h19.779v8.43c0 .815.661 1.477 1.476 1.477h7.383c.815 0 1.477-.661 1.477-1.477v-8.43h16.12l-1.699 2.943l.003.002c-.104.189-.18.396-.18.628c0 .732.593 1.325 1.325 1.325v.015h14.016v3.941c0 .578.469 1.046 1.046 1.046h5.232c.578 0 1.046-.468 1.046-1.046v-3.941h14.05v-.015c.732 0 1.326-.593 1.326-1.325a1.295 1.295 0 0 0-.173-.617"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trophy(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M82.296 23.931a1.918 1.918 0 0 0-1.918-1.918h-8.074v-6.998c.004-.061.018-.118.018-.179a2.935 2.935 0 0 0-2.934-2.935c-.036 0-.07.009-.106.011v-.011H30.365v.054a2.925 2.925 0 0 0-2.696 2.911c0 .062.014.119.018.179v6.967H19.62a1.914 1.914 0 0 0-1.909 1.839h-.007v.073l-.001.007l.001.007v26.038l-.001.004l.001.009V50h.001c.01 1.051.863 1.9 1.916 1.9h8.328c1.354 9.109 8.197 16.422 17.069 18.449v12.746h-9.969a2.493 2.493 0 0 0 0 4.988v.017h29.894v-.017a2.493 2.493 0 0 0 0-4.988h-9.969V70.353c8.881-2.02 15.733-9.337 17.087-18.453h8.318c1.028 0 1.86-.81 1.909-1.825h.011V23.931zM27.687 46.913H22.69V27h4.997zm49.623 0h-5.007V27h5.007z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Underline(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M77.5 75.545c-.036 0-.068.009-.103.01v-.01h-55v.01c-1.608.056-2.897 1.368-2.897 2.99s1.288 2.934 2.897 2.99v.01h55v-.01c.035.001.068.01.103.01a3 3 0 0 0 0-6M50 72.12c15.829 0 23.581-9.057 23.581-22.521V21.383a2.928 2.928 0 0 0-2.929-2.928h-3.864a2.928 2.928 0 0 0-2.929 2.928c0 .04.01.076.012.116v27.856c0 8.649-4.814 14.28-13.871 14.28s-13.871-5.631-13.871-14.28V21.49c.001-.036.011-.071.011-.107a2.928 2.928 0 0 0-2.928-2.928h-3.865a2.929 2.929 0 0 0-2.929 2.928v28.216c0 13.464 7.834 22.521 23.582 22.521"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UniversalAccess(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M50 12.5c-20.712 0-37.5 16.793-37.5 37.502C12.5 70.712 29.288 87.5 50 87.5s37.5-16.788 37.5-37.498C87.5 29.293 70.712 12.5 50 12.5m0 65.061c-15.199 0-27.56-12.362-27.56-27.559C22.44 34.807 34.802 22.44 50 22.44c15.198 0 27.56 12.367 27.56 27.562c0 15.196-12.362 27.559-27.56 27.559"/><circle cx="49.63" cy="32.546" r="4.723"/><path d="m65.892 48.189l-8.973-8.974a2.49 2.49 0 0 0-.089-.089l-.061-.061l-.006.006a2.205 2.205 0 0 0-1.457-.556c-.051 0-.098.012-.148.015H44.703l-.008-.001c-.561 0-1.067.214-1.458.557l-.005-.005l-.055.055a1.572 1.572 0 0 0-.1.1l-8.968 8.968l.015.015a1.986 1.986 0 1 0 2.808 2.809l.003.003l6.673-6.673v.004a.634.634 0 0 1 1.086.443v24.466a2.323 2.323 0 0 0 4.645 0V56.935h.017a.643.643 0 1 1 1.286 0h.017v12.334a2.323 2.323 0 0 0 4.645 0V44.875h.02c-.005-.034-.02-.064-.02-.099c0-.35.284-.633.635-.633a.62.62 0 0 1 .402.155v-.007l.021.021c.015.014.029.027.042.042l6.66 6.661l.003-.003a1.986 1.986 0 1 0 2.808-2.809z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unlink(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16.201 20.386l9.719 9.721a2.87 2.87 0 0 0 4.054 0a2.87 2.87 0 0 0 0-4.055l-9.679-9.677l-.006-.008c-.002-.003-.005-.004-.008-.006l-.026-.026l-.003.003a2.865 2.865 0 0 0-4.017.03a2.862 2.862 0 0 0-.03 4.015zm16.697-3.84l3.558 13.277a2.865 2.865 0 0 0 3.51 2.026a2.869 2.869 0 0 0 2.027-3.51l-3.545-13.224l-.001-.007l-.003-.007l-.011-.04l-.004.001a2.864 2.864 0 0 0-3.493-1.982a2.864 2.864 0 0 0-2.035 3.463zM16.054 39.423l-.001.005l.046.012h.002l13.229 3.544a2.866 2.866 0 1 0 1.485-5.535L17.6 33.908c-.005-.002-.01-.005-.016-.006l-.017-.003l-.03-.008l-.001.003a2.865 2.865 0 0 0-3.463 2.035a2.866 2.866 0 0 0 1.981 3.494m67.744 40.19l-9.72-9.72a2.866 2.866 0 1 0-4.053 4.055l9.679 9.677l.006.008l.008.006l.026.026l.003-.003a2.865 2.865 0 0 0 4.017-.03a2.864 2.864 0 0 0 .03-4.016zm-16.697 3.841l-3.559-13.277a2.865 2.865 0 0 0-3.51-2.026a2.865 2.865 0 0 0-2.027 3.509l3.546 13.231v.002l.012.045l.005-.001a2.864 2.864 0 0 0 3.493 1.983a2.865 2.865 0 0 0 2.034-3.463zm16.842-22.877l.001-.004l-.041-.011l-.005-.002l-.005-.001l-13.226-3.544c-1.53-.41-3.1.499-3.511 2.026a2.868 2.868 0 0 0 2.027 3.511l13.224 3.542a.014.014 0 0 0 .006.002l.007.001l.04.011l.001-.004a2.863 2.863 0 0 0 1.482-5.527m-35.731-8.821c-7.552-7.552-19.648-7.79-27.486-.713l-.019-.019L10.61 61.121c-7.797 7.797-7.797 20.44 0 28.237c7.797 7.798 20.439 7.798 28.237 0L48.945 79.26l-.019-.019c7.075-7.837 6.838-19.933-.714-27.485m-6.553 20.802l-.619.619l-.001.001h-.001l-9.005 9.005l-.001.001c-3.935 3.935-10.314 3.935-14.248 0s-3.935-10.314 0-14.248l.001-.001l9.005-9.006l.001-.001l.001-.001l.619-.619l.029.028c3.959-3.329 9.874-3.134 13.6.591c3.726 3.726 3.921 9.642.591 13.6zm47.73-61.917c-7.552-7.552-19.648-7.79-27.486-.713l-.019-.019l-10.097 10.097c-7.797 7.797-7.797 20.44 0 28.237c7.797 7.798 20.439 7.798 28.237 0l10.098-10.098l-.019-.019c7.075-7.837 6.838-19.933-.714-27.485m-6.553 20.802l-.619.619l-.001.001h-.001l-9.005 9.005l-.001.001c-3.935 3.935-10.314 3.935-14.248 0c-3.935-3.935-3.935-10.314 0-14.248l.001-.001l9.005-9.006l.001-.001l.001-.001l.619-.619l.029.028c3.959-3.329 9.874-3.134 13.6.591s3.921 9.642.591 13.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unlock(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M82.105 45.587H41.236V34.945c0-5.207 4.372-9.437 9.747-9.437c5.372 0 9.744 4.23 9.746 9.437c0 1.051.881 1.904 1.966 1.904h10.086c1.087 0 1.965-.853 1.965-1.904a1.8 1.8 0 0 0-.056-.457c-.253-12.456-10.79-22.513-23.707-22.513c-13.074 0-23.71 10.305-23.71 22.97v10.642h-9.377c-1.449 0-2.621 1.135-2.621 2.539v37.361c0 1.403 1.172 2.539 2.621 2.539h64.21c1.449 0 2.621-1.136 2.621-2.539V48.125c-.001-1.403-1.173-2.538-2.622-2.538"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Upload(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M85.886 49.615H67.685c-2.218 0-4.185 1.374-4.865 3.406c-1.803 5.375-6.957 8.983-12.819 8.983c-5.864 0-11.016-3.608-12.818-8.983c-.682-2.032-2.647-3.406-4.867-3.406h-18.2c-2.825 0-5.115 2.2-5.115 4.919v28.054c0 2.714 2.29 4.917 5.115 4.917h71.771c2.827 0 5.114-2.203 5.114-4.917V54.534c-.001-2.719-2.289-4.919-5.115-4.919"/><path fill="currentColor" d="M35.81 34.276h6.653v17.188c0 .735.617 1.326 1.376 1.326h12.317c.762 0 1.376-.591 1.376-1.326V34.276h6.658c.514 0 .986-.279 1.221-.716a1.276 1.276 0 0 0-.102-1.374L51.072 13.048a1.398 1.398 0 0 0-1.119-.553h-.003c-.445 0-.859.207-1.119.558L34.686 32.187a1.288 1.288 0 0 0-.1 1.374c.235.441.71.715 1.224.715"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UploadCloud(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M79.437 36.831c-3.51-12.982-14.464-22.083-27.134-22.083c-8.835 0-17.065 4.454-22.414 12.018a20.172 20.172 0 0 0-4.501-.514C13.5 26.251 3.825 36.79 3.825 49.749c0 4.647 1.251 9.148 3.612 13.018c.555.906 1.49 1.449 2.49 1.449H29.85c.143-9.265 7.688-16.734 16.987-16.734s16.843 7.469 16.987 16.734h26.945c.947 0 1.836-.485 2.403-1.315a17.25 17.25 0 0 0 3.004-9.768c-.001-9.656-7.816-17.471-16.739-16.302"/><path fill="currentColor" d="m57.345 70.33l-9.918-13.861a.963.963 0 0 0-.779-.4h-.002c-.31 0-.598.15-.779.404L36.013 70.33a.963.963 0 0 0-.069.996c.164.32.494.518.853.518h4.634v12.449c0 .532.43.96.958.96h8.58c.53 0 .958-.428.958-.96V71.844h4.638a.959.959 0 0 0 .78-1.514"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Usb(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64.603 35.504a2.777 2.777 0 0 0-2.774-2.72c-.047 0-.091.012-.138.014h-2.039V17.774a1.687 1.687 0 0 0-1.684-1.684H41.765v.03a1.681 1.681 0 0 0-1.386 1.654v15.022h-2.052c-.043-.002-.083-.013-.126-.013a2.776 2.776 0 0 0-2.774 2.72h-.036v45.625h.001a2.782 2.782 0 0 0 2.78 2.781l.014-.001h23.643a2.78 2.78 0 0 0 2.78-2.779V35.504zm-9.938-2.706h-9.329v-11.72h9.329z"/><path fill="currentColor" d="M47.506 27.933h4.988v2.072h-4.988z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Video(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M66.272 61.337v10.565c0 2.497-2.043 4.535-4.539 4.535H14.54c-2.496 0-4.54-2.038-4.54-4.535V28.097c0-2.496 2.043-4.535 4.54-4.535h47.193c2.496 0 4.539 2.038 4.539 4.535v10.432v-.146L90 27.265v45.294z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Volume(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M78.864 17.021c-.026-.026-.056-.042-.082-.067l.008-.008l-1.924-1.923l-.022.022a2.696 2.696 0 0 0-3.553.161l-.004-.004l-3.419 3.418l.023.023a2.693 2.693 0 0 0 .01 3.345l-.022.022l.216.216l1.707 1.708l.031-.031c.025.026.042.057.067.083c14.358 14.358 14.401 37.688.138 52.104l-.019-.019l-1.707 1.707l-.216.216l.022.022a2.698 2.698 0 0 0 .16 3.553l-.004.004l3.42 3.42l.023-.023a2.693 2.693 0 0 0 3.345-.011l.022.022l.216-.216l1.707-1.707l-.004-.004c18.102-18.257 18.058-47.835-.139-66.033"/><path fill="currentColor" d="M69.376 30.198c-.026-.026-.056-.042-.082-.067l.008-.008l-1.925-1.923l-.022.022a2.7 2.7 0 0 0-3.554.16l-.004-.004l-.035.035l-.003.002l-.002.003l-3.149 3.148l-.003.002l-.002.003l-.225.225l.023.023a2.691 2.691 0 0 0 .011 3.344l-.022.022l1.923 1.924l.031-.031c.025.026.042.057.067.083c7.091 7.091 7.132 18.594.135 25.746l-.014-.014l-1.707 1.707h-.002l-.215.215l.022.022a2.698 2.698 0 0 0 .16 3.554l-.004.004l3.42 3.42l.023-.023a2.691 2.691 0 0 0 3.344-.011l.022.022l1.923-1.923l-.004-.004c10.838-10.99 10.794-28.745-.138-39.678m-16.625-6.395a2.04 2.04 0 0 0-1.031.285l-.018-.032l-20.464 11.815v.012l-7.74 4.469H9.016v.04c-.012 0-.024-.004-.037-.004c-.842 0-1.525.684-1.525 1.525v20.66c0 .842.683 1.524 1.525 1.524c.013 0 .024-.003.037-.004v.041h14.482l11.524 6.653v-.031l16.548 9.555a2.071 2.071 0 0 0 3.252-1.698c0-.081-.015-.155-.024-.233h.024V25.64h-.024a2.064 2.064 0 0 0-2.047-1.837"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeNone(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M73.684 23.396h-.024a2.064 2.064 0 0 0-2.047-1.837c-.378 0-.726.108-1.031.285l-.019-.032L50.1 33.627v.012l-7.74 4.469H27.879v.04c-.013 0-.024-.004-.037-.004c-.842 0-1.525.684-1.525 1.525v20.66c0 .842.684 1.524 1.525 1.524c.013 0 .024-.003.037-.004v.041H42.36l11.524 6.653v-.031l16.549 9.555a2.07 2.07 0 0 0 3.251-1.698c0-.081-.015-.155-.024-.233h.024z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeStrike(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m60.416 35.232l-.022.022l1.923 1.924l.031-.031c.025.026.042.057.067.083c4.39 4.39 6.071 10.47 5.059 16.163l7.868 7.868c4.384-10.269 2.403-22.629-5.964-30.995c-.026-.026-.056-.042-.082-.067l.008-.008l-1.924-1.923l-.022.021a2.7 2.7 0 0 0-3.554.161l-.004-.004l-.035.035l-.003.002l-.002.003l-3.149 3.148l-.003.002l-.002.003l-.225.225l.023.022a2.697 2.697 0 0 0 .012 3.346"/><path fill="currentColor" d="M78.868 17.089c-.026-.025-.056-.042-.082-.067l.008-.008l-1.924-1.923l-.022.022a2.696 2.696 0 0 0-3.553.161l-.004-.004l-3.419 3.418l.023.023a2.693 2.693 0 0 0 .01 3.345l-.022.021l.216.216l1.707 1.708l.031-.031c.025.026.042.057.067.083c11.153 11.153 13.663 27.718 7.545 41.315l7.332 7.332c9.752-17.77 7.125-40.573-7.913-55.611m-24.042 8.619h-.024a2.064 2.064 0 0 0-2.047-1.837a2.04 2.04 0 0 0-1.031.285l-.018-.032l-8.558 4.941l11.679 11.679V25.708zM23.502 40.42H9.021v.04c-.012 0-.024-.003-.037-.003c-.842 0-1.525.684-1.525 1.525v20.66c0 .842.683 1.524 1.525 1.524c.013 0 .024-.003.037-.004v.041h14.482l11.524 6.653v-.031l16.548 9.555a2.071 2.071 0 0 0 3.252-1.698c0-.081-.015-.155-.024-.233h.024v-9.413L25.219 39.428zm62.316 38.368L21.172 14.141l-.002.002l-.016-.02a2.487 2.487 0 0 0-3.516 0c-.034.034-.058.074-.09.11l-3.458 3.476l-.012.01a2.485 2.485 0 0 0 0 3.515l64.651 64.651l.001-.001a2.486 2.486 0 0 0 3.507-.009l.004.004l3.432-3.449c.049-.042.103-.076.15-.123a2.488 2.488 0 0 0-.005-3.519"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Web(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M50 12.5c-20.712 0-37.5 16.793-37.5 37.502C12.5 70.712 29.288 87.5 50 87.5s37.5-16.788 37.5-37.498C87.5 29.293 70.712 12.5 50 12.5M22.897 54.969h5.752a37.328 37.328 0 0 0 8.826 19.569c-7.501-3.844-13.015-11.022-14.578-19.569m5.752-9.939h-5.752c1.564-8.546 7.078-15.724 14.579-19.568a37.33 37.33 0 0 0-8.827 19.568M45.03 68.059a27.468 27.468 0 0 1-6.313-13.089h6.313zm0-23.029h-6.313a27.474 27.474 0 0 1 6.313-13.088zm32.072 0h-5.751a37.328 37.328 0 0 0-8.825-19.567c7.499 3.845 13.013 11.023 14.576 19.567M54.97 68.057V54.969h6.313a27.487 27.487 0 0 1-6.313 13.088m0-23.027V31.941a27.468 27.468 0 0 1 6.313 13.089zm7.553 29.509a37.33 37.33 0 0 0 8.828-19.57h5.752c-1.563 8.548-7.078 15.726-14.58 19.57"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wheelchair(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M51.77 75.42a13.883 13.883 0 0 1-7.941 2.487c-7.706 0-13.954-6.248-13.954-13.954c0-2.952.922-5.686 2.487-7.94l-5.348-5.348a21.555 21.555 0 0 0-4.852 13.639c0 11.967 9.701 21.668 21.667 21.668c5.174 0 9.923-1.816 13.649-4.843z"/><circle cx="68.862" cy="21.181" r="7.152" fill="currentColor"/><path fill="currentColor" d="M73.35 45.582H57.609l8.977-8.978c.271-.215.515-.458.73-.729l.073-.073l-.007-.007a4.448 4.448 0 0 0 .919-2.694a4.47 4.47 0 0 0-2.159-3.819l.001-.001L46.04 17.675l-.017.03a4.454 4.454 0 0 0-1.843-.405a4.456 4.456 0 0 0-2.986 1.16l-.006-.006l-.06.06c-.077.072-.15.145-.222.222l-8.865 8.865l.008.008a4.463 4.463 0 0 0-1.237 3.08a4.486 4.486 0 0 0 4.487 4.487a4.463 4.463 0 0 0 3.08-1.237l.008.008l.081-.081l.008-.008l6.457-6.457l5.514 3.184l-13.011 13.012a21.568 21.568 0 0 0-6.854 3.582l5.307 5.307a13.882 13.882 0 0 1 7.94-2.487c7.707 0 13.954 6.248 13.954 13.954c0 2.952-.922 5.686-2.487 7.94l5.666 5.667a21.57 21.57 0 0 0 4.534-13.256c0-3.507-.838-6.816-2.317-9.746h5.683v18.371a4.488 4.488 0 1 0 8.976 0v-22.86a4.488 4.488 0 0 0-4.488-4.487"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Widget(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m88.553 43.441l-10.276-1.013a29.016 29.016 0 0 0-2.936-7.059l6.561-7.993a2.158 2.158 0 0 0-.145-2.893l-6.24-6.243a2.153 2.153 0 0 0-2.894-.142L64.63 24.66a29.042 29.042 0 0 0-7.06-2.936l-1.011-10.278A2.158 2.158 0 0 0 54.412 9.5h-8.823a2.16 2.16 0 0 0-2.148 1.946L42.43 21.727a28.903 28.903 0 0 0-7.054 2.936l-7.998-6.564c-.4-.329-.884-.49-1.369-.49c-.555 0-1.105.212-1.525.633l-6.24 6.241a2.158 2.158 0 0 0-.145 2.893l6.566 7.996a28.914 28.914 0 0 0-2.936 7.057l-10.281 1.013A2.154 2.154 0 0 0 9.5 45.587v8.825c0 1.11.843 2.04 1.947 2.146l10.281 1.013a28.87 28.87 0 0 0 2.936 7.054l-6.562 7.994a2.15 2.15 0 0 0 .14 2.895l6.24 6.242a2.155 2.155 0 0 0 2.894.142l7.993-6.562a29.029 29.029 0 0 0 7.06 2.938l1.011 10.279a2.158 2.158 0 0 0 2.148 1.946h8.823c1.11 0 2.039-.842 2.148-1.946l1.011-10.278a28.853 28.853 0 0 0 7.06-2.939l7.993 6.562a2.156 2.156 0 0 0 2.894-.142l6.24-6.242a2.15 2.15 0 0 0 .14-2.895l-6.557-7.991a29.019 29.019 0 0 0 2.936-7.057l10.276-1.013a2.157 2.157 0 0 0 1.947-2.146v-8.825a2.156 2.156 0 0 0-1.946-2.146M50.002 61.95c-6.603 0-11.953-5.351-11.953-11.95c0-6.602 5.351-11.951 11.953-11.951C56.6 38.049 61.95 43.398 61.95 50c.001 6.599-5.35 11.95-11.948 11.95"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wrench(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m86.257 23.405l-3.866 3.866l-3.737 3.737l-4.759 4.759a9.08 9.08 0 0 1-9.663-9.663l4.759-4.759l3.737-3.737l3.866-3.866a.645.645 0 0 0 0-.911c-.046-.046-.101-.074-.155-.103l.001-.001l-.01-.004a.649.649 0 0 0-.102-.043a21.424 21.424 0 0 0-8.749-1.878c-11.939 0-21.618 9.679-21.618 21.618c0 2.28.358 4.475 1.012 6.538L24.428 61.504c-7.545.122-13.627 6.267-13.627 13.842c0 7.65 6.203 13.853 13.853 13.853c7.574 0 13.72-6.083 13.842-13.628l22.546-22.546a21.602 21.602 0 0 0 6.539 1.012c11.939 0 21.618-9.679 21.618-21.618c0-3.118-.686-6.066-1.877-8.742a.605.605 0 0 0-.05-.118l-.022-.052l-.007.007c-.024-.037-.041-.078-.074-.111a.646.646 0 0 0-.912.002M30.378 75.346a5.724 5.724 0 1 1-11.449 0a5.724 5.724 0 0 1 11.449 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func X(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M84.707 68.752L65.951 49.998l18.75-18.752a1.989 1.989 0 0 0 0-2.813L71.566 15.295a1.99 1.99 0 0 0-2.814 0L49.999 34.047l-18.75-18.752c-.746-.747-2.067-.747-2.814 0L15.297 28.431a1.992 1.992 0 0 0 0 2.814L34.05 49.998L15.294 68.753a1.993 1.993 0 0 0 0 2.814L28.43 84.704a1.988 1.988 0 0 0 2.814 0l18.755-18.755l18.756 18.754c.389.388.896.583 1.407.583s1.019-.195 1.408-.583l13.138-13.137a1.99 1.99 0 0 0-.001-2.814"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Xcircle(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m68.056 59.756l-9.758-9.757l9.755-9.756a1.036 1.036 0 0 0 0-1.464l-6.833-6.835a1.036 1.036 0 0 0-1.464 0L49.999 41.7l-9.755-9.756c-.388-.388-1.075-.388-1.464 0l-6.835 6.835a1.037 1.037 0 0 0 0 1.464l9.756 9.756l-9.758 9.758a1.036 1.036 0 0 0 0 1.463l6.833 6.834a1.033 1.033 0 0 0 1.464 0l9.757-9.757l9.758 9.756c.202.203.466.304.732.304c.266 0 .53-.101.732-.304l6.835-6.834a1.033 1.033 0 0 0 .002-1.463"/><path fill="currentColor" d="M50 22.44c15.196 0 27.56 12.367 27.56 27.562c0 15.197-12.364 27.559-27.56 27.559S22.44 65.199 22.44 50.002C22.44 34.807 34.804 22.44 50 22.44m0-9.94c-20.71 0-37.5 16.793-37.5 37.502C12.5 70.712 29.29 87.5 50 87.5c20.709 0 37.5-16.788 37.5-37.498C87.5 29.293 70.709 12.5 50 12.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Yen(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M76.275 22.811a.268.268 0 0 0 .079-.191c0-.134-.103-.233-.232-.255v-.017H64.121a.905.905 0 0 0-.78.463h-.016l-13.281 21.49l-13.281-21.49h-.014a.904.904 0 0 0-.78-.463H23.917a.272.272 0 0 0-.193.463h-.001l16.42 25.354H25.977v.017c-.025-.002-.046-.015-.072-.015c-.637 0-1.153.516-1.153 1.153v4.12h.014a1.146 1.146 0 0 0 1.139 1.083c.026 0 .047-.012.072-.014v.015h18.352v6.439H25.977v.002l-.012-.002c-.637 0-1.153.516-1.153 1.153v3.969c0 .637.516 1.153 1.153 1.153l.012-.002v.004h18.352v9.199c-.001.021-.012.038-.012.058c0 .637.516 1.153 1.153 1.153h.002l.001.001h9.141v-.001a1.153 1.153 0 0 0 1.151-1.153c0-.011-.006-.021-.006-.031v-9.225h18.433v-.004l.01.002c.637 0 1.153-.516 1.153-1.153v-3.969c0-.637-.516-1.153-1.153-1.153l-.01.002v-.002H55.759v-6.439h18.433v-.003l.01.002c.637 0 1.153-.516 1.153-1.153V49.25h-.014a1.146 1.146 0 0 0-1.139-1.083l-.01.002v-.004H59.864l16.42-25.354z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomIn(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M56.774 10.391c-17.679 0-32 14.329-32 32a31.845 31.845 0 0 0 4.588 16.517L13.846 74.423l.054.054c-1.656 1.585-2.673 3.835-2.673 6.378c-.001 4.913 3.913 8.755 8.821 8.754c2.507-.001 4.749-1.004 6.349-2.636l.039.039l16.008-16.009a31.865 31.865 0 0 0 14.33 3.388c17.68 0 32-14.327 32-32c-.001-17.671-14.321-32-32-32m.194 51.417c-11.05 0-20-8.954-20-20c0-11.044 8.951-20 20-20c11.05 0 20 8.955 20 20c0 11.046-8.95 20-20 20"/><path fill="currentColor" d="M68.751 37.111v-.03l-7.076.008l-.009-7.058a1.57 1.57 0 0 0-1.569-1.57c-.033 0-.064.008-.097.01h-6.084c-.033-.002-.064-.01-.097-.01c-.867 0-1.569.702-1.569 1.57c0 .054.011.105.016.158l-.007 6.911l-7.061.008c-.868 0-1.569.703-1.569 1.57v6.278c0 .867.702 1.569 1.569 1.569c.054 0 .105-.011.158-.016l6.893.007l-.009 7.034h.011l-.002.024c0 .868.703 1.569 1.569 1.569l.01-.001v.011h6.278v-.011a1.568 1.568 0 0 0 1.558-1.556h.03l-.009-7.059l7.028.007v-.011l.024.002c.868 0 1.569-.703 1.569-1.569v-6.278a1.564 1.564 0 0 0-1.555-1.567"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomOut(children ...ElementRenderer) *FoundationIcon {
	return &FoundationIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M56.774 10.391c-17.679 0-32 14.329-32 32a31.845 31.845 0 0 0 4.588 16.517L13.845 74.423l.054.054c-1.656 1.585-2.673 3.835-2.673 6.378c-.001 4.913 3.913 8.755 8.821 8.754c2.507-.001 4.749-1.004 6.349-2.636l.039.039l16.008-16.009a31.865 31.865 0 0 0 14.33 3.388c17.68 0 32-14.327 32-32c0-17.671-14.32-32-31.999-32m.194 51.417c-11.05 0-20-8.954-20-20c0-11.044 8.951-20 20-20c11.05 0 20 8.955 20 20c0 11.046-8.95 20-20 20"/><path fill="currentColor" d="M68.739 37.099H45.197c-.868 0-1.57.702-1.57 1.569v6.278c0 .867.702 1.569 1.57 1.569h23.542a1.57 1.57 0 0 0 1.57-1.569v-6.278a1.57 1.57 0 0 0-1.57-1.569"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
