package el

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = "2.0.0"
	hAttr          = "1em"
	viewbox        = "0 0 1200 1200"
	fill           = "currentColor"
)

type ElIcon struct {
	*SVGSVGElement
}

type ElIconFn func(children ...ElementRenderer) *ElIcon

var IconLookup = map[string]ElIconFn{
	"addressBook":      AddressBook,
	"addressBookAlt":   AddressBookAlt,
	"adjust":           Adjust,
	"adjustAlt":        AdjustAlt,
	"adult":            Adult,
	"alignCenter":      AlignCenter,
	"alignJustify":     AlignJustify,
	"alignLeft":        AlignLeft,
	"alignRight":       AlignRight,
	"arrowDown":        ArrowDown,
	"arrowLeft":        ArrowLeft,
	"arrowRight":       ArrowRight,
	"arrowUp":          ArrowUp,
	"asl":              Asl,
	"asterisk":         Asterisk,
	"backward":         Backward,
	"banCircle":        BanCircle,
	"barcode":          Barcode,
	"behance":          Behance,
	"bell":             Bell,
	"blind":            Blind,
	"blogger":          Blogger,
	"bold":             Bold,
	"book":             Book,
	"bookmark":         Bookmark,
	"bookmarkEmpty":    BookmarkEmpty,
	"braille":          Braille,
	"briefcase":        Briefcase,
	"broom":            Broom,
	"brush":            Brush,
	"bulb":             Bulb,
	"bullhorn":         Bullhorn,
	"calendar":         Calendar,
	"calendarSign":     CalendarSign,
	"camera":           Camera,
	"car":              Car,
	"caretDown":        CaretDown,
	"caretLeft":        CaretLeft,
	"caretRight":       CaretRight,
	"caretUp":          CaretUp,
	"cc":               Cc,
	"certificate":      Certificate,
	"check":            Check,
	"checkEmpty":       CheckEmpty,
	"chevronDown":      ChevronDown,
	"chevronLeft":      ChevronLeft,
	"chevronRight":     ChevronRight,
	"chevronUp":        ChevronUp,
	"child":            Child,
	"circleArrowDown":  CircleArrowDown,
	"circleArrowLeft":  CircleArrowLeft,
	"circleArrowRight": CircleArrowRight,
	"circleArrowUp":    CircleArrowUp,
	"cloud":            Cloud,
	"cloudAlt":         CloudAlt,
	"cog":              Cog,
	"cogAlt":           CogAlt,
	"cogs":             Cogs,
	"comment":          Comment,
	"commentAlt":       CommentAlt,
	"compass":          Compass,
	"compassAlt":       CompassAlt,
	"creditCard":       CreditCard,
	"css":              Css,
	"dashboard":        Dashboard,
	"delicious":        Delicious,
	"deviantart":       Deviantart,
	"digg":             Digg,
	"download":         Download,
	"downloadAlt":      DownloadAlt,
	"dribbble":         Dribbble,
	"edit":             Edit,
	"eject":            Eject,
	"envelope":         Envelope,
	"envelopeAlt":      EnvelopeAlt,
	"errorAlt":         ErrorAlt,
	"errorIcon":        ErrorIcon,
	"eur":              Eur,
	"exclamationSign":  ExclamationSign,
	"eyeClose":         EyeClose,
	"eyeOpen":          EyeOpen,
	"facebook":         Facebook,
	"facetimeVideo":    FacetimeVideo,
	"fastBackward":     FastBackward,
	"fastForward":      FastForward,
	"female":           Female,
	"file":             File,
	"fileAlt":          FileAlt,
	"fileEdit":         FileEdit,
	"fileEditAlt":      FileEditAlt,
	"fileNew":          FileNew,
	"fileNewAlt":       FileNewAlt,
	"film":             Film,
	"filter":           Filter,
	"fire":             Fire,
	"flag":             Flag,
	"flagAlt":          FlagAlt,
	"flickr":           Flickr,
	"folder":           Folder,
	"folderClose":      FolderClose,
	"folderOpen":       FolderOpen,
	"folderSign":       FolderSign,
	"font":             Font,
	"fontsize":         Fontsize,
	"fork":             Fork,
	"forward":          Forward,
	"forwardAlt":       ForwardAlt,
	"foursquare":       Foursquare,
	"friendfeed":       Friendfeed,
	"friendfeedRect":   FriendfeedRect,
	"fullscreen":       Fullscreen,
	"gbp":              Gbp,
	"gift":             Gift,
	"github":           Github,
	"githubText":       GithubText,
	"glass":            Glass,
	"glasses":          Glasses,
	"globe":            Globe,
	"globeAlt":         GlobeAlt,
	"googleplus":       Googleplus,
	"graph":            Graph,
	"graphAlt":         GraphAlt,
	"groupAlt":         GroupAlt,
	"groupIcon":        GroupIcon,
	"guidedog":         Guidedog,
	"handDown":         HandDown,
	"handLeft":         HandLeft,
	"handRight":        HandRight,
	"handUp":           HandUp,
	"hdd":              Hdd,
	"headphones":       Headphones,
	"hearingImpaired":  HearingImpaired,
	"heart":            Heart,
	"heartAlt":         HeartAlt,
	"heartEmpty":       HeartEmpty,
	"home":             Home,
	"homeAlt":          HomeAlt,
	"hourglass":        Hourglass,
	"idea":             Idea,
	"ideaAlt":          IdeaAlt,
	"inbox":            Inbox,
	"inboxAlt":         InboxAlt,
	"inboxBox":         InboxBox,
	"indentLeft":       IndentLeft,
	"indentRight":      IndentRight,
	"infoCircle":       InfoCircle,
	"instagram":        Instagram,
	"iphoneHome":       IphoneHome,
	"italic":           Italic,
	"key":              Key,
	"laptop":           Laptop,
	"laptopAlt":        LaptopAlt,
	"lastfm":           Lastfm,
	"leaf":             Leaf,
	"lines":            Lines,
	"link":             Link,
	"linkedin":         Linkedin,
	"list":             List,
	"listAlt":          ListAlt,
	"livejournal":      Livejournal,
	"lock":             Lock,
	"lockAlt":          LockAlt,
	"magic":            Magic,
	"magnet":           Magnet,
	"male":             Male,
	"mapMarker":        MapMarker,
	"mapMarkerAlt":     MapMarkerAlt,
	"mic":              Mic,
	"micAlt":           MicAlt,
	"minus":            Minus,
	"minusSign":        MinusSign,
	"move":             Move,
	"music":            Music,
	"myspace":          Myspace,
	"network":          Network,
	"off":              Off,
	"ok":               Ok,
	"okCircle":         OkCircle,
	"okSign":           OkSign,
	"opensource":       Opensource,
	"paperClip":        PaperClip,
	"paperClipAlt":     PaperClipAlt,
	"path":             Path,
	"pause":            Pause,
	"pauseAlt":         PauseAlt,
	"pencil":           Pencil,
	"pencilAlt":        PencilAlt,
	"person":           Person,
	"phone":            Phone,
	"phoneAlt":         PhoneAlt,
	"photo":            Photo,
	"photoAlt":         PhotoAlt,
	"picasa":           Picasa,
	"picture":          Picture,
	"pinterest":        Pinterest,
	"plane":            Plane,
	"play":             Play,
	"playAlt":          PlayAlt,
	"playCircle":       PlayCircle,
	"plurk":            Plurk,
	"plurkAlt":         PlurkAlt,
	"plus":             Plus,
	"plusSign":         PlusSign,
	"podcast":          Podcast,
	"print":            Print,
	"puzzle":           Puzzle,
	"qrcode":           Qrcode,
	"question":         Question,
	"questionSign":     QuestionSign,
	"quoteAlt":         QuoteAlt,
	"quoteRight":       QuoteRight,
	"quoteRightAlt":    QuoteRightAlt,
	"quotes":           Quotes,
	"random":           Random,
	"record":           Record,
	"reddit":           Reddit,
	"redux":            Redux,
	"refresh":          Refresh,
	"remove":           Remove,
	"removeCircle":     RemoveCircle,
	"removeSign":       RemoveSign,
	"repeat":           Repeat,
	"repeatAlt":        RepeatAlt,
	"resizeFull":       ResizeFull,
	"resizeHorizontal": ResizeHorizontal,
	"resizeSmall":      ResizeSmall,
	"resizeVertical":   ResizeVertical,
	"returnKey":        ReturnKey,
	"retweet":          Retweet,
	"reverseAlt":       ReverseAlt,
	"road":             Road,
	"rss":              Rss,
	"scissors":         Scissors,
	"screen":           Screen,
	"screenAlt":        ScreenAlt,
	"screenshot":       Screenshot,
	"search":           Search,
	"searchAlt":        SearchAlt,
	"share":            Share,
	"shareAlt":         ShareAlt,
	"shoppingCart":     ShoppingCart,
	"shoppingCartSign": ShoppingCartSign,
	"signal":           Signal,
	"skype":            Skype,
	"slideshare":       Slideshare,
	"smiley":           Smiley,
	"smileyAlt":        SmileyAlt,
	"soundcloud":       Soundcloud,
	"speaker":          Speaker,
	"spotify":          Spotify,
	"stackoverflow":    Stackoverflow,
	"star":             Star,
	"starAlt":          StarAlt,
	"starEmpty":        StarEmpty,
	"stepBackward":     StepBackward,
	"stepForward":      StepForward,
	"stop":             Stop,
	"stopAlt":          StopAlt,
	"stumbleupon":      Stumbleupon,
	"tag":              Tag,
	"tags":             Tags,
	"tasks":            Tasks,
	"textHeight":       TextHeight,
	"textWidth":        TextWidth,
	"th":               Th,
	"thLarge":          ThLarge,
	"thList":           ThList,
	"thumbsDown":       ThumbsDown,
	"thumbsUp":         ThumbsUp,
	"time":             Time,
	"timeAlt":          TimeAlt,
	"tint":             Tint,
	"torso":            Torso,
	"trash":            Trash,
	"trashAlt":         TrashAlt,
	"tumblr":           Tumblr,
	"twitter":          Twitter,
	"universalAccess":  UniversalAccess,
	"unlock":           Unlock,
	"unlockAlt":        UnlockAlt,
	"upload":           Upload,
	"usd":              Usd,
	"user":             User,
	"viadeo":           Viadeo,
	"video":            Video,
	"videoAlt":         VideoAlt,
	"videoChat":        VideoChat,
	"viewMode":         ViewMode,
	"vimeo":            Vimeo,
	"vkontakte":        Vkontakte,
	"volumeDown":       VolumeDown,
	"volumeOff":        VolumeOff,
	"volumeUp":         VolumeUp,
	"wthreeC":          WthreeC,
	"warningSign":      WarningSign,
	"website":          Website,
	"websiteAlt":       WebsiteAlt,
	"wheelchair":       Wheelchair,
	"wordpress":        Wordpress,
	"wrench":           Wrench,
	"wrenchAlt":        WrenchAlt,
	"youtube":          Youtube,
	"zoomIn":           ZoomIn,
	"zoomOut":          ZoomOut,
}

func AddressBook(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.008.015v1199.984H1200V987.564h-130.664V870.733H1200V658.377h-130.664V541.545H1200V329.282h-130.664V212.451H1200l-.016-212.45H0zm534.665 209.56c95.784 0 173.373 77.68 173.373 173.466c0 67.881-38.969 126.625-95.78 155.109l222.013 132.926h2.696v187.473H232.28V671.076h2.784L456.982 538.15c-56.812-28.484-95.78-87.229-95.78-155.109c0-95.785 77.68-173.466 173.466-173.466z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AddressBookAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.635 0 0 268.635 0 600s268.635 600 600 600s600-268.635 600-600S931.365 0 600 0M263.688 263.688h672.626V382.75h-73.219v65.5h73.219v119h-73.219v65.5h73.219v119h-73.219v65.5h73.219v119.094H263.688zm299.687 117.468c-53.692 0-97.188 43.558-97.188 97.25c0 38.051 21.842 70.971 53.688 86.938l-124.438 74.5h-1.531v105.094h338.938V639.844h-1.531l-124.438-74.5c31.846-15.967 53.688-48.887 53.688-86.938c-.001-53.692-43.496-97.25-97.188-97.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Adjust(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.641 0 0 268.641 0 600s268.641 600 600 600s600-268.641 600-600S931.359 0 600 0m0 174.975V1025.1c-234.75 0-425.109-190.336-425.109-425.109c0-234.751 190.336-425.016 425.109-425.016"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdjustAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M104.312 0v699.844H0v254.062h367.281V699.844H262.969V0zM520.75 0v451.75H416.344v254.062h367.312V451.75H679.312V0zm416.281 0v201.625H832.719v254.062H1200V201.625h-104.312V0zm-51.093 252.031h260.844v50.438H885.938zm51.093 246.094V1200h158.656V498.125zm-467.437 4.031h260.812v50.438H469.594zm51.156 244.032V1200h158.562V746.188zm-467.531 4.031h260.844v50.438H53.219zm51.125 246.093V1200H263V996.312z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Adult(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M605.096 480c-135.542-2.098-239.082-111.058-239.999-240C367.336 105.667 477.133.942 605.096 0c135.662 5.13 239.036 108.97 240.001 240c-2.668 134.439-111.907 239.09-240.001 240m194.043 49.788c170.592 1.991 257.094 151.63 257.881 301.269V1200H889.784l.001-324.254c-4.072-22.416-19.255-30.018-33.164-27.82c-13.022 2.059-24.929 12.701-25.56 27.82V1200h-464.67V875.746c-3.557-21.334-17.128-29.537-30.331-28.709c-14.138.889-27.853 12.135-28.393 28.709V1200h-164.68V831.057c-.98-159.475 99.901-300.087 259.137-301.269z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignCenter(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M290.625 99.202v178.006h618.75V99.202zm-164.063 274.53v178.006h946.875V373.732zm107.813 274.53v178.006h731.25V648.262zM0 922.792v178.006h1200V922.792z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignJustify(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 99.202v178.006h1200V99.202zm0 274.53v178.006h1200V373.732zm0 274.53v178.006h1200V648.262zm0 274.53v178.006h1200V922.792z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignLeft(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M618.75 99.202v178.006H0V99.202zm328.125 274.53v178.006H0V373.732zM731.25 648.262v178.006H0V648.262zM1200 922.792v178.006H0V922.792z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignRight(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M581.25 99.202v178.006H1200V99.202zm-328.125 274.53v178.006H1200V373.732zm215.625 274.53v178.006H1200V648.262zM0 922.792v178.006h1200V922.792z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDown(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 1200L131.25 496.875h252.539V0h432.422v496.875h252.539z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeft(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 600l703.125 468.75V816.211H1200V383.789H703.125V131.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRight(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1200 600L496.875 131.25v252.539H0v432.422h496.875v252.539z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUp(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M599.992 0L131.243 703.131h252.546V1200h432.422V703.131h252.546z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Asl(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M709.516.007c-.548.012-1.095.039-1.638.063c-32.364 1.444-64.713 31.739-70.976 91.666l-7.872 697.99c-5.189 125.227-142.08 167.029-205.81 97.117c-75.037 97.223-204.71 19.813-213.776-91.321c-14.126-83.898-28.781-167.831-41.943-251.815c-2.542-26.818 2.412-54.485 15.744-79.604l-65.591-225.719c-23.02-68.161-128.592-72.87-116.729 24.78l141.636 793.291c12.173 65.205 48.563 138.618 129.828 143.546h434.102c35.847 0 63.164-6.978 81.965-20.896c161.439-121.679 132.631-373.002 353.995-468.794c51.953-31.315 83.29-129 29.631-168.546c-88.635-79.617-317.355 62.572-357.144 205.92l-35.11-653.348C777.494 32.268 744.809.52 711.151.008a46.896 46.896 0 0 0-1.635-.001M490.48 388.319a88.454 88.454 0 0 0-7.872.438c-34.677 3.444-65.704 27.55-65.969 76.66l27.553 309.24c2.335 52.705 43.063 90.059 80.453 89.817c32.66-.21 62.783-29.105 62.505-101.565L576.632 458.9c-4.699-45.34-46.471-70.979-86.152-70.581m-198.41 68.452c-4.699.062-9.36.59-13.949 1.566c-33.287 7.082-61.577 38.259-54.255 94.517l44.588 266.164c11.993 43.965 41.719 60.28 69.023 54.637c30.867-6.381 58.643-40.821 54.255-95.081l-26.23-257.048c-6.894-43.325-40.535-65.194-73.432-64.755"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Asterisk(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M489.838 29.354v443.603L68.032 335.894L0 545.285l421.829 137.086l-260.743 358.876l178.219 129.398L600.048 811.84l260.673 358.806l178.146-129.398l-260.766-358.783L1200 545.379l-68.032-209.403l-421.899 137.07V29.443H489.84z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Backward(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M627.277 0v545.462L1172.723 0v1200L627.262 654.538V1200L27.277 600z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BanCircle(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1024.263 175.738c-234.317-234.317-614.192-234.317-848.525 0c-234.317 234.317-234.317 614.192 0 848.525c234.317 234.316 614.192 234.316 848.525 0c234.316-234.318 234.316-614.193 0-848.525m-163.489 57.44L233.193 860.743c-125.257-175.737-109.044-421.274 48.624-578.942s403.219-173.881 578.942-48.624zm106.064 106.048c125.248 175.738 109.031 421.29-48.654 578.942c-157.652 157.683-403.205 173.911-578.942 48.639l627.581-627.581z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Barcode(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 99.87v1000.26h100.26V99.87zm186.135 0v1000.26h24.089V99.87zm57.883 0v1000.26h76.219V99.87zm153.658 0v1000.26h32.525V99.87zm118.449 0v1000.26h24.041V99.87zm34.379 0v1000.26h99.626V99.87zm159.411 0v1000.26h50.18V99.87zm149.025 0v1000.26h11.118V99.87zm87.63 0v1000.26h33.501V99.87zm119.376 0v1000.26h24.09V99.87zm34.428 0v1000.26H1200V99.87z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Behance(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v1200h1200V0zm186.476 495.854H491.75c31.229.219 64.056 12.23 89.209 35.376c23.154 22.86 36.754 55.645 37.427 90.747c-.116 26.956-5.251 58.017-23.146 82.765c-12.504 16.834-32.084 29.562-53.686 36.841v2.856c22.748 5.689 45.174 17.843 61.157 37.939c19.012 24.585 25.514 57.747 26.074 90.381c-.228 34.622-12.213 71.699-39.038 99.169c-23.845 23.426-58.079 36.835-94.117 37.573H186.47V495.858zm534.96 1.026h227.048v84.228H721.436zm108.691 109.423c31.679.001 59.783 4.115 84.301 12.378c24.735 8.021 45.552 20.562 62.476 37.573c16.923 16.771 29.698 38.155 38.378 64.16c11.447 34.042 13.618 77.622 13.403 117.408H764.716c0 27.221 5.411 48.359 16.26 63.427c11.065 14.826 29.176 22.193 54.345 22.193c23.913-.112 51.923-7.78 61.158-34.645c2.387-7.048 3.588-14.538 3.589-22.558h128.613c-.267 40.87-13.426 81.462-47.899 111.548c-38.084 30.978-91.314 39.815-142.53 40.796c-31.68 0-60.311-3.991-85.913-12.012c-25.386-8.264-46.955-20.926-64.747-37.939c-17.793-17.013-31.469-38.397-41.016-64.16c-9.547-26.006-14.355-56.729-14.355-92.212c0-34.511 4.659-64.617 13.989-90.381c9.33-25.762 22.557-47.146 39.697-64.16c17.358-17.255 38.174-30.165 62.475-38.671c24.519-8.507 51.795-12.744 81.739-12.744zm-496.51 2.197v87.452h101.222c9.124-.07 17.806-4.249 24.39-11.646c6.394-7.785 9.528-18.041 9.741-28.784v-7.324c-.066-9.694-3.792-19.708-10.107-28.052c-5.739-6.779-14.348-11.156-24.023-11.646zm501.706 92.578c-22.349 0-39.032 5.962-50.097 17.87c-10.85 11.909-17.462 28.071-19.85 48.487h129.492c0-20.659-5.186-36.821-15.601-48.487c-10.413-11.908-25.068-17.87-43.944-17.87M333.617 802.079v87.817h111.914c9.021-.07 17.658-4.211 24.39-11.279c6.363-7.75 9.653-18.055 9.814-28.785v-7.69c-.058-9.781-3.229-19.984-9.814-28.417c-5.78-6.818-14.729-11.093-24.39-11.646z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bell(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M521.329 0v74.121c-163.599 35.891-286.205 181.326-286.954 355.52v319.63L-1.832 953.535v55.591h1203.665v-55.591L965.625 749.271v-319.63c-.753-174.194-123.364-319.629-286.97-355.52V0zm78.661 1024.585c-48.447 0-87.743 39.223-87.743 87.672c0 48.447 39.296 87.743 87.743 87.743c48.448 0 87.744-39.296 87.744-87.743c0-48.448-39.296-87.672-87.744-87.672"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Blind(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M385.312 0c-65.813 0-119.166 53.352-119.166 119.166c0 65.813 53.353 119.166 119.166 119.166c65.814 0 119.167-53.353 119.167-119.166C504.479 53.352 451.126 0 385.312 0m-99.026 299.713L158.228 472.611c-5.122 6.829-7.676 14.053-7.676 21.736L135.2 636.503c-6.928 62.019 81.763 76.349 92.188 8.968l12.806-131.897l33.288-44.802l1.292 185.666l-62.737 225.338l-160.055 203.603c-46.537 72.435 49.647 126.372 98.572 69.158L317.03 938.676c4.471-6.019 7.239-14.069 10.222-21.736l49.97-176.736l112.67 126.766l61.445 286.87c16.961 76.578 134.751 50.207 120.382-23.065L606.446 832.42c-1.555-7.685-3.242-16.207-10.222-21.774L457.943 653.137V458.501l53.77 67.867c5.976 7.683 14.953 13.229 26.903 16.644l126.767 24.32l469.914 618.482c8.392 11.667 32.137.041 23.065-15.353L720.44 568.598c51.496-17.263 23.699-75.92-10.26-81.964l-135.733-26.866l-139.573-175.446c-44.207-43.902-118.137-33.112-148.54 15.352z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Blogger(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v1200h1200V0zm641.742 219.792c115.163 9.225 182.107 82.968 190.065 177.254c6.186 48.52 4.496 84.54-2.563 111.035c65.691 2.152 142.468 31.339 147.363 137.328v104.298c-6.812 91.112-39.518 203.688-207.277 227.562l-280.744 1.025c-170.153 3.09-263.875-71.142-264.749-229.92l2.198-351.932c7.324-102.37 90.271-162.382 187.053-175.709l228.655-.952zM450.507 362.176c-98.028 15.913-102.874 126.92 0 140.624h186.102c98.028-15.913 102.874-126.921 0-140.624zM408.32 669.213c-100.919 15.533-102.489 141.578 3.516 154.614l341.82 1.172c100.061-15.401 99.777-141.28-3.589-154.614l-341.743-1.172z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bold(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v97.286h150.591v1005.478H0V1200h654.399c187.638 0 325.367-28.1 413.15-84.363c88.305-56.263 132.449-143.888 132.45-262.841c0-87.876-31.529-158.346-94.607-211.395c-63.077-53.049-153.749-85.202-272.02-96.457c97.771-10.716 171.354-36.961 220.766-78.755c49.938-41.793 74.904-98.336 74.906-169.603c0-99.131-36.007-173.315-108.019-222.61C949.012 24.68 841.254.001 697.753 0zm454.163 97.286h103.288c93.036 0 160.572 15.786 202.625 47.399c42.574 31.613 63.884 82.244 63.885 151.902c0 69.123-21.817 120.583-65.445 154.34c-43.631 33.222-110.652 49.838-201.064 49.837H454.163zm0 499.935h112.749c99.346 0 173.443 21.162 222.327 63.491c48.883 42.333 73.344 106.353 73.345 192.084c0 85.2-24.206 148.19-72.564 188.914c-48.359 40.724-122.709 61.054-223.107 61.054h-112.75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Book(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M454.771 10.997c-76.209-40.113-226.406 37.395-269.553 105.63c-19.222 30.534-17.862 52.538-17.862 65.022v667.874L730.602 1200l105.917-57.833V491.739L258.215 159.706c31.033-39.057 100.827-86.683 153.16-67.555l515.104 275.498l.001 724.58l106.184-57.936V309.728z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookmark(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M234.375 0v1200L600 834.375L965.625 1200V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkEmpty(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M234.364 0v1200L600 834.364L965.636 1200V0zm75 75h581.271v937.5L600 721.864L309.364 1012.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Braille(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M510.202 1050.982c0 82.301-66.717 149.018-149.018 149.018c-82.3 0-149.018-66.717-149.018-149.018c0-82.3 66.717-149.018 149.018-149.018c82.301.001 149.018 66.718 149.018 149.018m0-450.982c0 82.301-66.717 149.018-149.018 149.018c-82.3 0-149.018-66.717-149.018-149.018s66.717-149.018 149.018-149.018S510.202 517.699 510.202 600m0-450.982c0 82.3-66.717 149.017-149.018 149.017c-82.3 0-149.018-66.717-149.018-149.017C212.167 66.717 278.884 0 361.185 0c82.3 0 149.017 66.717 149.017 149.018m477.631 901.964c0 82.301-66.717 149.018-149.018 149.018s-149.018-66.717-149.018-149.018c0-82.3 66.717-149.018 149.018-149.018s149.018 66.718 149.018 149.018m0-450.982c0 82.301-66.717 149.018-149.018 149.018S689.798 682.301 689.798 600s66.717-149.018 149.018-149.018S987.833 517.699 987.833 600m0-450.982c0 82.3-66.717 149.017-149.018 149.017s-149.018-66.717-149.018-149.017C689.798 66.717 756.515 0 838.815 0s149.018 66.717 149.018 149.018"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Briefcase(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M318.141 73.72v167.244H0v487.372h541.219V628.314h117.562v100.021H1200V240.964H881.808V73.72zm86.456 86.456h390.755v80.788H404.597zM0 822.188v304.092h1200V822.188H658.781v96.617H541.219v-96.617z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Broom(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1093.636 0L683.919 409.716c-29.44-16.92-67.651-12.834-92.811 12.325c-21.045 21.044-27.335 51.226-18.931 77.763c-57.811-29.551-124.29-34.53-191.204-5.992C255.457 556.096 168.858 685.687 0 700.744c12.889 26.536 29.579 56.126 50.049 88.726c71.005 11.18 140.925-11.4 195.559-67.14c-21.221 66.046-73.074 115.338-143.2 141.02c18.195 24.261 39.007 49.729 62.509 76.265c80.851-27.028 109.762-64.018 34.115 36.294c31.904 33.138 66.075 68.279 108.064 97.849c18.654-68.878 68.927-121.768 140.953-148.987c-55.188 51.727-77.289 126.649-63.667 206.934c41.698 28.051 79.998 50.86 114.873 68.297c29.705-127.434 116.39-259.614 206.935-380.972c29.215-61.498 24.481-127.952-6.605-191.407c26.685 8.697 57.168 2.481 78.375-18.726c25.159-25.159 29.244-63.37 12.324-92.811L1200 106.37L1093.639.008z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Brush(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1157.602.013c-46.711 2.677-736.479 591.498-793.123 798.838l130.736 136.944C868.899 624.199 988.915 294.221 1200 .649L1157.617 0zM323.267 840.562C87.09 927.418 235.147 1099.273 0 1183.352c266.294 59.953 384.296-49.748 454.003-205.421L323.267 840.548z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bulb(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 119.608c-158.775 0-287.486 96.415-287.486 215.368l163.322 494.645h248.328l163.322-494.659c0-118.951-128.702-215.369-287.486-215.369zm0 46.996c158.775 0 240.479 128.138 240.479 168.362L724.161 680.393H475.833L359.515 334.966c0-49.63 81.704-168.362 240.479-168.362zM472.215 865.699v85.982h255.57v-85.984h-255.57zm0 128.725v85.982h255.57v-85.982z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bullhorn(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1030.014 96.91C874.836 250.578 646.766 344.637 427.042 357.866H100.749C49.358 358.572 2.898 402.262 0 459.889V607.02c.701 51.369 43.068 99.121 100.699 102.021H231.52c-34.962 118.567 2.593 238.596 43.15 356.415c60.166 56.81 197.785 49.742 242.638-24.195c-95.523-75.091-142.05-277.145-26.889-326.979c196.285 22.271 390.749 122.01 539.594 255.716c51.169-.701 97.705-42.745 100.699-100.062V628.618c102.449-36.383 81.738-164.259 0-190.28V197.659c-.683-51.365-43.073-97.799-100.699-100.7zm2.595 130.135v613.453C888.778 728.424 711.072 653.216 527.741 619.46V448.134c179.894-27.072 355.835-110.077 504.868-221.089"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M266.025 0v188.019h98.741V0zm569.208 0v188.019h98.593V0zm-771 90.368v263.226h1071.533V90.368h-131.341V258.52H764.634V90.368H435.218V258.52H195.574V90.368zm0 332.34V1200h1071.533V422.707zm420.924 179.596c54.544.537 106.162 26.743 126.684 75.505c9.651 29.128 14.79 60.179 2.774 87.296c-15.042 31.283-45.537 47.569-75.455 58.065c56.149 14.65 93.102 49.96 93.439 107.312c-.905 39.339-19.316 74.184-47.116 96.957c-29.559 23.369-65.3 36.423-100.326 36.663c-98.775 0-148.185-44.535-148.185-133.62h87.197c1.853 37.506 19.587 66.032 56.084 66.537c46.396-3.161 64.767-33.646 65.149-72.681c-.021-23.577-8.628-47.739-29.131-56.826c-26.45-10.991-55.977-9.685-81.748-9.711v-67.825c42.127.878 99.996-6.554 100.426-56.777c-1.638-36.941-16.758-63.298-52.616-63.763c-38.196 2.303-59.157 24.017-59.551 58.312h-85.81c1.863-86.473 72.898-125.058 148.185-125.444m300.979 8.373h76.843v445.2h-83.135V697.874c-31.85 27.233-67.103 46.625-105.875 58.164V679.89c45.234-17.541 82.626-40.594 112.167-69.213"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarSign(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.635 0 0 268.635 0 600s268.635 600 600 600s600-268.635 600-600S931.365 0 600 0M389.353 221.555h62.256v118.58h-62.256zm359.032 0h62.184v118.58h-62.184zm-486.336 56.982h82.837v106.055h151.172V278.537H703.85v106.055h151.245V278.537h82.838v166.034H262.048zm0 209.618h675.885v490.283H262.049zm265.499 113.306c-47.488.243-92.282 24.559-93.457 79.102h54.127c.248-21.633 13.479-35.315 37.572-36.768c22.618.293 32.146 16.909 33.179 40.21c-.271 31.679-36.783 36.369-63.354 35.815v42.774c16.255.017 34.88-.78 51.562 6.152c12.933 5.731 18.37 20.944 18.384 35.815c-.242 24.621-11.823 43.855-41.089 45.85c-23.02-.318-34.207-18.311-35.376-41.967h-55.005c0 56.19 31.153 84.301 93.457 84.301c22.093-.151 44.637-8.405 63.281-23.145a81.44 81.44 0 0 0 29.736-61.158c-.214-36.175-23.544-58.435-58.961-67.676c18.871-6.621 38.119-16.889 47.606-36.621c7.579-17.104 4.33-36.705-1.757-55.078c-12.943-30.755-45.503-47.267-79.905-47.606m189.847 5.273c-18.634 18.051-42.221 32.589-70.752 43.652v48.047c24.455-7.278 46.708-19.517 66.797-36.694v225.8h52.441V606.732h-48.486z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Camera(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M390.724 148.172L352.969 300H0v751.828h1200V300H863.484l-37.756-151.828zm212.341 151.537c186.396 0 337.453 151.066 337.453 337.453c0 186.396-151.068 337.453-337.453 337.453c-186.396 0-337.453-151.068-337.453-337.453c0-186.395 151.067-337.453 337.453-337.453m395.719 73.463h170.121v98.779H998.784zm-395.719 26.196c-131.337 0-237.797 106.47-237.797 237.797c0 131.337 106.47 237.797 237.797 237.797c131.337 0 237.797-106.47 237.797-237.797c0-131.337-106.469-237.797-237.797-237.797"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Car(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m214.662 100.752l-75.97 316.401H0v91.456h47.573v590.639h136.569v-96.685h831.716v96.685h136.496V508.609H1200v-91.456h-138.692l-75.971-316.401zm40.474 61.52h689.713l61.187 254.887H193.934l61.186-254.887zm-168.048 423.7l190.44 70.329v91.384l-190.44-70.329zm1025.73 0v91.384l-190.362 70.329v-91.456zM415.833 825.036h368.193v91.384H415.833z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretDown(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 125.407l600 949.186l600-949.186z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretLeft(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1100.007 1200L99.993 600L1100.007 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretRight(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m100 0l1000 600l-1000 600z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretUp(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 1099.996l600-999.993l600 999.993z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cc(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 99.531L0 216.719v766.562l600 117.188l600-117.188V216.719zM381.844 366.094c93.039 0 173.443 54.348 211.125 133H479.344c-25.248-24.406-59.609-39.469-97.5-39.469c-77.505 0-140.312 62.839-140.312 140.344s62.807 140.312 140.312 140.312c37.866 0 72.225-14.966 97.469-39.344h113.625c-37.694 78.636-118.078 132.938-211.094 132.938c-129.175 0-233.875-104.731-233.875-233.906s104.7-233.875 233.875-233.875m459.094 0c93.039 0 173.412 54.348 211.094 133H938.438c-25.249-24.406-59.609-39.469-97.5-39.469c-77.505 0-140.344 62.839-140.344 140.344s62.839 140.312 140.344 140.312c37.866 0 72.162-14.966 97.406-39.344H1052c-37.694 78.636-118.047 132.938-211.062 132.938c-129.175 0-233.875-104.731-233.875-233.906s104.7-233.875 233.875-233.875"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Certificate(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M978.231 1066.743L718.077 939.444l-136.759 255.307l-117.461-264.739l-268.871 107.669l80.19-278.306L0 669.027l240.32-161.649l-152.722-246.09L375.6 291.933L416.791 5.249l200.925 208.6L833.547 20.713l19.832 288.948l289.481-9.217L972.32 534.54L1200 713.555l-281.114 69.706z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Check(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v1200h1200V424.289l-196.875 196.875v381.961h-806.25v-806.25h381.961L775.711 0zm1030.008 15.161l-434.18 434.25L440.7 294.283L281.618 453.438L595.821 767.57l159.082-159.082l434.18-434.25l-159.082-159.081z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckEmpty(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v1200h1200V0zm196.875 196.875h806.25v806.25h-806.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDown(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m600.006 989.352l178.709-178.709L1200 389.357l-178.732-178.709L600.006 631.91L178.721 210.648L0 389.369l421.262 421.262l178.721 178.721z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeft(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m210.632 599.989l178.735 178.735L810.633 1200l178.735-178.721l-421.267-421.29l421.267-421.266L810.645 0L389.378 421.267L210.655 599.989z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRight(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M989.368 600.011L810.633 421.275L389.367 0L210.632 178.721l421.267 421.29l-421.267 421.267L389.355 1200l421.266-421.267L989.345 600.01z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUp(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600.002 210.605L421.285 389.336L0 810.559l178.721 178.836l421.281-421.341l421.281 421.341L1200 810.559L778.733 389.336z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Child(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M604.729 0C469.354 0 359.598 110.317 359.598 246.429c0 136.114 109.765 246.471 245.132 246.471c135.375 0 245.089-110.35 245.089-246.471C849.818 110.314 740.103 0 604.729 0M464.355 287.402h279.366c0 77.145-62.56 139.704-139.704 139.704s-139.662-62.559-139.662-139.704m-52.274 244.336c-127.661 2.621-211.637 85.551-206.878 235.714L204.952 1200h146.945V811.356c1.271-34.026 51.241-45.945 57.673 0V1200h386.133V811.356c1.271-34.026 51.241-45.945 57.673 0V1200h141.672V767.452c-.858-118.72-49.572-236.814-205.581-235.714z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleArrowDown(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1200 600C1200 268.641 931.359 0 600 0S0 268.641 0 600s268.641 600 600 600s600-268.641 600-600m-297.141-66.656L600 987.656L297.141 533.344h163.184V212.32h279.352v321.023z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleArrowLeft(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 1200c331.359 0 600-268.641 600-600S931.359 0 600 0S0 268.641 0 600s268.641 600 600 600m66.656-297.141L212.32 600l454.336-302.859v163.184H987.68v279.352H666.656z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleArrowRight(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.635 0 0 268.635 0 600s268.635 600 600 600s600-268.635 600-600S931.365 0 600 0m-66.647 297.145L987.667 600L533.353 902.855V739.678h-321.02v-279.34h321.02z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleArrowUp(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1200 600c0 331.359-268.641 600-600 600S0 931.359 0 600S268.641 0 600 0s600 268.641 600 600m-297.141 66.633L600 212.32L297.141 666.633h163.184v321.023h279.352V666.633z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloud(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M983.888 575.377c187.925-18.507 293.084 231.644 148.656 358.546H49.759C-89.529 697.252 82.314 382.276 333.563 443.401C535.007 115.536 908.131 291.199 983.88 575.377z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.635 0 0 268.635 0 600s268.635 600 600 600s600-268.635 600-600S931.365 0 600 0m5.347 311.938c112.641-2.458 224.208 89.791 256.494 210.944c128.177-12.623 199.88 157.997 101.366 244.557H224.7c-94.99-161.431 22.206-376.261 193.577-334.58C469.8 349 537.753 313.412 605.33 311.936z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cog(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m506.274 0l-18.238 181.063c-37.264 9.933-72.584 24.601-105.185 43.498L241.97 109.428L109.428 241.97l115.133 140.929a431.034 431.034 0 0 0-43.596 105.186L0 506.274v187.5l181.063 18.189c9.928 37.243 24.616 72.551 43.498 105.137L109.428 958.03l132.542 132.542L382.9 975.438c32.585 18.882 67.893 33.571 105.136 43.498L506.226 1200h187.5l18.189-180.966A430.925 430.925 0 0 0 817.1 975.438l140.931 115.134l132.542-132.542l-115.135-140.882c18.897-32.601 33.565-67.921 43.498-105.185L1200 693.774v-187.5l-180.966-18.189a430.816 430.816 0 0 0-43.596-105.233l115.134-140.881L958.03 109.428L817.148 224.561a430.871 430.871 0 0 0-105.233-43.596L693.726 0zM600 426.544c95.787 0 173.456 77.669 173.456 173.456S695.787 773.456 600 773.456S426.544 695.787 426.544 600S504.213 426.544 600 426.544"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CogAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0m-57.129 234.375h114.258l11.06 110.303c22.729 6.054 44.277 14.988 64.16 26.514l85.84-70.166l80.786 80.786l-70.166 85.84c11.525 19.883 20.46 41.432 26.514 64.16l110.303 11.06V657.13l-110.303 11.133c-6.054 22.709-14.997 44.22-26.514 64.087l70.166 85.84l-80.786 80.786l-85.913-70.166c-19.874 11.516-41.371 20.464-64.087 26.514l-11.06 110.303H542.871l-11.133-110.303a262.31 262.31 0 0 1-64.014-26.514l-85.913 70.166l-80.786-80.786l70.166-85.913a262.317 262.317 0 0 1-26.514-64.014L234.374 657.13V542.871l110.303-11.06c6.05-22.716 14.998-44.213 26.514-64.087l-70.166-85.913l80.786-80.786l85.84 70.166c19.867-11.517 41.378-20.46 64.087-26.514zM600 494.312c-58.374 0-105.688 47.314-105.688 105.688S541.626 705.688 600 705.688S705.688 658.374 705.688 600S658.374 494.312 600 494.312"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cogs(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m910.143 91.119l-16.916 81.053a196.327 196.327 0 0 0-49.691 14.743l-58.229-58.825l-66.309 53.661l45.354 69.303c-10.104 13.862-18.357 29.104-24.623 45.503l-82.85-.374l-8.906 84.87l81.055 16.914c2.723 17.329 7.746 33.897 14.742 49.545l-58.824 58.376l53.66 66.31l69.303-45.43a195.996 195.996 0 0 0 45.504 24.698l-.375 82.774l84.871 8.904l16.838-81.053c17.346-2.722 34.035-7.666 49.695-14.669l58.301 58.825l66.309-53.735l-45.428-69.229a195.81 195.81 0 0 0 24.697-45.503l82.773.3l8.906-84.87l-81.053-16.765c-2.725-17.354-7.66-34.103-14.67-49.77l58.824-58.227l-53.734-66.309l-69.229 45.354c-13.869-10.111-29.17-18.428-45.578-24.697l.373-82.774zm14.068 197.131c2.668.009 5.373.09 8.084.374c43.355 4.555 74.756 43.384 70.201 86.741c-4.555 43.355-43.385 74.83-86.742 70.274c-43.355-4.555-74.83-43.384-70.275-86.739c4.269-40.647 38.724-70.789 78.732-70.65m-608.981 6.96l-11.375 112.711c-23.205 6.187-45.185 15.324-65.486 27.092l-87.714-71.696l-82.55 82.55l71.698 87.788c-11.768 20.308-20.91 42.272-27.092 65.484L0 610.44v116.751l112.71 11.376a268.028 268.028 0 0 0 27.092 65.41l-71.698 87.789l82.55 82.55l87.789-71.697a268.061 268.061 0 0 0 65.411 27.093l11.375 112.71h116.752l11.301-112.71c23.212-6.183 45.178-15.325 65.484-27.093l87.788 71.697l82.55-82.55l-71.697-87.714c11.768-20.302 20.906-42.281 27.092-65.485l112.711-11.376V610.44L634.5 599.138c-6.186-23.225-15.314-45.243-27.092-65.561l71.697-87.714l-82.55-82.549l-87.713 71.696c-20.316-11.775-42.336-20.905-65.562-27.093l-11.301-112.71H315.23zm58.376 265.61c59.649 0 107.996 48.348 107.996 107.996c0 59.647-48.347 107.994-107.996 107.994c-59.648 0-107.996-48.347-107.996-107.994c0-59.648 48.348-107.996 107.996-107.996m495.673 144.219l-11.9 59.273c-12.188 1.993-23.873 5.653-34.877 10.776l-41.012-43.033l-46.553 39.292l31.883 50.667c-7.102 10.143-12.959 21.308-17.363 33.306l-58.15-.301l-6.287 62.118l56.955 12.35a147.66 147.66 0 0 0 10.328 36.298l-41.312 42.659l37.721 48.497l48.721-33.229a135.985 135.985 0 0 0 31.957 18.037l-.225 60.621l59.648 6.511l11.824-59.349c12.189-1.991 23.869-5.579 34.875-10.702l41.014 43.033l46.625-39.291l-31.957-50.668c7.104-10.139 12.959-21.312 17.363-33.306l58.229.226l6.211-62.043l-56.953-12.35c-1.914-12.695-5.402-24.911-10.328-36.372l41.311-42.585l-37.719-48.572l-48.646 33.229c-9.746-7.396-20.498-13.449-32.031-18.036l.299-60.546zm9.879 144.144c1.873.009 3.783.092 5.688.301c30.473 3.331 52.521 31.745 49.32 63.465c-3.201 31.719-30.449 54.748-60.922 51.416s-52.596-31.746-49.395-63.466c3.001-29.738 27.19-51.818 55.309-51.716"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Comment(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 49.109v854.516h247.266v247.266L632.74 903.625H1200V49.109z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 60.509v729.264h211.064v210.992l328.929-210.992h484.071V60.509zm1072.243 160.306v91.767h35.989v520.062h-177.01v139.013l-205.329-131.7l-11.328-7.312H559.78l-143.028 91.768h270.785L952.3 1094.181l70.689 45.311v-215.08H1200V220.815z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compass(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M175.736 1024.264c234.315 234.314 614.213 234.314 848.527 0s234.314-614.212 0-848.527c-234.314-234.314-614.212-234.314-848.527 0c-234.314 234.314-234.314 614.212 0 848.527m58.575-58.575c-201.966-201.966-201.966-529.412 0-731.378s529.412-201.966 731.378 0c201.967 201.967 201.967 529.413 0 731.378c-201.965 201.967-529.411 201.967-731.378 0m58.522-31.281l398.11-243.465l243.465-398.109l-27.241-27.242l-398.11 243.465l-243.465 398.11z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompassAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M175.736 1024.264c234.315 234.314 614.213 234.314 848.527 0s234.314-614.212 0-848.527c-234.314-234.314-614.212-234.314-848.527 0c-234.314 234.314-234.314 614.212 0 848.527m94.931-65.618l-29.261-29.262l261.125-426.853l426.853-261.125l29.261 29.261l-261.124 426.854z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCard(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M-2.031 202.672V452.36h1204.062V202.672zm0 363.781v430.875h1204.062V566.453zm131.562 189.25h453.688v136.938H129.531z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Css(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 23.584v927.1l600 225.732l600-225.732v-927.1zm287.329 368.042c85.63 1.335 145.205 44.392 145.752 130.078H340.21c.063-22.869-9.76-44.64-30.322-53.32c-23.199-8.298-52.19-4.187-67.09 15.454c-15.955 23.851-13.11 50.549-13.11 78.662c-.879 37.461 18.296 69.844 56.323 70.093c38.528-1.939 56.835-20.738 57.202-55.444h89.868c-2.157 92.03-64.41 127.729-145.752 128.101c-102.733.624-156.271-71.726-157.104-156.885c-.605-102.699 72.001-155.903 157.104-156.739m324.976 0c61.678-.651 132.208 26.193 133.96 93.53v5.127h-91.187v-1.685c-.563-22.112-20.637-30.676-39.111-30.762c-13.39.343-39.815 3.222-40.21 19.775c3.371 26.164 68.358 30.933 88.55 35.889c46.904 9.965 88.957 36.56 90.308 85.986c-1.521 80.015-66.448 105.503-133.447 105.762c-79.6-.569-149.041-19.703-149.634-107.886h91.992c-.858 31.54 23.761 38.174 48.779 38.232c14.99-.002 42.551-.543 42.773-20.215c-3.029-24.25-70.219-30.547-89.429-35.449c-49.614-12.438-88.86-38.271-89.209-91.333c3.43-77.204 72.728-96.733 135.865-96.971m318.164 0c66.098.555 133.887 25.562 133.887 98.657h-91.113c.812-23.003-20.373-32.359-39.111-32.446c-13.39.343-39.815 3.222-40.21 19.775c2.864 25.832 68.604 30.992 88.55 35.889c32.667 7.887 66.021 21.116 81.885 50.537c5.302 9.601 8.062 21.406 8.35 35.449c-1.464 80.036-66.423 105.503-133.447 105.762c-79.554-.549-149.03-19.751-149.634-107.886h91.992c-.497 20.146 10.264 33.738 28.638 36.914c16.626 1.562 62.68 7.109 62.988-18.896c-3.621-24.639-69.896-30.464-89.429-35.449c-49.663-12.446-88.86-38.221-89.209-91.333c3.441-77.258 72.647-96.734 135.863-96.973"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dashboard(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 195.373c-331.371 0-600 268.629-600 600c0 73.594 13.256 144.104 37.5 209.253h164.062C168.665 942.111 150 870.923 150 795.373c0-248.528 201.471-450 450-450s450 201.472 450 450c0 75.55-18.665 146.738-51.562 209.253H1162.5c24.244-65.148 37.5-135.659 37.5-209.253c0-331.371-268.629-600-600-600m0 235.62c-41.421 0-75 33.579-75 75c0 41.422 33.579 75 75 75s75-33.578 75-75c0-41.421-33.579-75-75-75m-224.927 73.462c-41.421 0-75 33.579-75 75c0 41.422 33.579 75 75 75s75-33.578 75-75c0-41.421-33.579-75-75-75m449.854 0c-41.422 0-75 33.579-75 75c0 41.422 33.578 75 75 75c41.421 0 75-33.578 75-75c0-41.421-33.579-75-75-75M600 651.672l-58.813 294.141v58.814h117.627v-58.814z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Delicious(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v600h600V0zm600 600v600h600V600z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Deviantart(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" stroke="currentColor" stroke-width="5.851" d="M436.084 351.603h135.93V848.04l-569.089.357c4.97-138.378 79.926-332.443 433.159-378.596l1.73 53.189c-160.578 23.664-216.349 118.607-244.039 236.4h248.219zm198.83 219.814l-2.09 275.938l564.162-.127c5.427-276.043-240.012-401.035-625.802-394.913v55.371c171.055-4.427 385.215 1.916 444.017 249.694l-252.828-2.09V581.863c-45.272-13.217-84.276-13.447-127.459-10.446z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Digg(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1123.2 375.086h-75.429V524.57H600.684V375.086h-76.8v229.027h76.8v441.601h223.544V826.285h76.8v219.43h145.371V604.114h76.801zM677.486 155.658v293.485H970.97V155.658zM1200 303.771v377.143h-76.8v440.229H523.886V680.914H300.343v145.37h71.314v75.43h76.8v219.429H75.429V901.714h76.8v144h219.429v-144H152.229v-75.43h71.314v-145.37H75.429v-76.801h371.657V303.771h153.6V78.857h447.086v224.914zm-828.343 0v-71.313h75.429v71.313zm-296.228 0v300.343H0V303.771zm76.8 0h-76.8v-71.313h76.8zm219.428-71.314H152.229v-76.8h219.429z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Download(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.63 0 0 268.63 0 600s268.63 600 600 600c331.369 0 600-268.631 600-600C1200 268.63 931.369 0 600 0m0 1069.565c-259.37 0-469.565-210.261-469.565-469.565S340.63 130.435 600 130.435c259.369 0 469.565 210.261 469.565 469.565S859.369 1069.565 600 1069.565m117.392-720.652H482.608v266.739H335.87L600 864.13l264.13-248.478H717.391z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownloadAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 1037.516h1200V1200H0zM820.785 0h-441.57v496.632H103.233L600 959.265l496.768-462.633H820.785z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dribbble(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0m0 86.93c130.718 0 250.017 48.884 340.615 129.369c-66.562 87.811-167.722 146.267-269.035 185.867c-55.245-101.044-123.248-207.576-192.476-300.912C517.868 91.891 558.358 86.93 600 86.93m-218.275 48.632c66.894 100.024 131.657 192.828 190.919 297.872c-149.748 38.792-318.895 62.061-475.152 62.462c32.946-159.879 140.295-292.564 284.233-360.334M997.53 275.608c71.56 87.59 114.718 199.269 115.502 321.01c-118.663-23.374-236.688-29.539-356.877-17.438c-13.539-33.691-29.874-65.684-45.555-100.19c103.609-41.788 214.343-113.311 286.93-203.382M614.286 512.804c12.859 27.337 27.733 56.292 41.679 85.904c-167.101 73.684-344.814 172.173-437.12 344.757C136.846 852.525 86.93 732.091 86.93 600c0-5.017.085-10.025.228-15.008c178.986-.868 354.622-23.269 527.128-72.188M885.41 652.697c74.604-.232 151.623 10.19 221.277 28.495c-22.781 143.311-104.915 266.845-220.403 344.644c-27.552-125.258-53.801-245.1-100.189-365.958c31.993-4.77 65.405-7.075 99.315-7.181m223.937 9.309c-.35 2.904-.704 5.812-1.102 8.7c.4-2.895.749-5.789 1.102-8.7M691.109 683.32c47.319 122.162 84.457 257.458 109.309 389.134c-61.575 26.155-129.3 40.615-200.418 40.615c-118.664 0-227.901-40.301-314.817-107.94c92.915-150.784 236.828-264.83 405.926-321.809"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Edit(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v1200h1200V424.292l-196.875 196.875v381.958h-806.25v-806.25h381.958L775.708 0zm1050 0l-76.831 76.831l150 150L1200 150zM936.914 113.086L497.168 552.832l150 150l439.746-439.746zM441.943 622.339c-2.225.034-4.493.195-6.738.366v142.09h142.09c0-38.708-18.492-78.039-47.314-105.542c-23.842-22.751-54.675-37.428-88.038-36.914"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eject(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 65.002L0 711.584h1200zM0 825.623v309.375h1200V825.623z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Envelope(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 146.484v168.677l600 342.114l600-342.114V146.484zm0 276.563v494.604L305.64 597.29zm1200 0L894.36 597.29L1200 917.651zM389.575 645.19L0 1053.516h1200L810.425 645.19L600 765.161z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0M292.969 367.896h614.062v86.353L600 629.297L292.969 454.248zm0 141.577l156.372 89.136l-156.372 163.915zm614.062 0v253.052L750.659 598.608zM492.334 623.145L600 684.521l107.666-61.377l199.365 208.96H292.969z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ErrorAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0M197.314 439.453h805.371v321.094H197.314z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ErrorIcon(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0m0 164.062c240.762 0 435.938 195.176 435.938 435.938S840.762 1035.938 600 1035.938S164.062 840.762 164.062 600S359.238 164.062 600 164.062M281.47 482.153v235.693h637.06V482.153z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eur(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M754.722 199.33c-65.372 0-119.224 19.023-161.554 57.066c-42.331 37.509-69.658 92.164-81.982 163.965h323.912v141.46H499.129l-1.607 28.132v37.776l1.607 26.523h285.332v143.067H512.793C540.12 927.527 626.121 992.63 770.797 992.63c76.624 0 150.301-15.271 221.031-45.813v205.76C929.671 1184.194 851.172 1200 756.33 1200c-131.28 0-239.251-35.633-323.912-106.898c-84.662-71.266-137.978-169.86-159.946-295.78H162.357v-143.07H257.2c-2.144-12.322-3.215-28.934-3.215-49.831l1.607-42.6h-93.235v-141.46h106.899c19.825-129.67 73.141-232.014 159.946-307.031C516.007 37.778 624.513.002 754.721 0c100.736.002 195.043 21.97 282.921 65.909l-78.768 186.47c-36.973-16.611-71.534-29.471-103.684-38.581c-32.15-9.644-65.639-14.467-100.468-14.468"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationSign(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0M381.958 118.213h436.084l-93.75 660.571H475.708zm220.386 705.542c71.262 0 129.053 57.79 129.053 129.053s-57.791 128.979-129.053 128.979s-129.053-57.717-129.053-128.979s57.791-129.053 129.053-129.053"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeClose(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M669.727 273.516c-22.891-2.476-46.15-3.895-69.727-4.248c-103.025.457-209.823 25.517-310.913 73.536c-75.058 37.122-148.173 89.529-211.67 154.174C46.232 529.978 6.431 577.76 0 628.74c.76 44.162 48.153 98.67 77.417 131.764c59.543 62.106 130.754 113.013 211.67 154.174c2.75 1.335 5.51 2.654 8.276 3.955l-75.072 131.102l102.005 60.286l551.416-960.033l-98.186-60.008zm232.836 65.479l-74.927 129.857c34.47 44.782 54.932 100.006 54.932 159.888c0 149.257-126.522 270.264-282.642 270.264c-6.749 0-13.29-.728-19.922-1.172l-49.585 85.84c22.868 2.449 45.99 4.233 69.58 4.541c103.123-.463 209.861-25.812 310.84-73.535c75.058-37.122 148.246-89.529 211.743-154.174c31.186-32.999 70.985-80.782 77.417-131.764c-.76-44.161-48.153-98.669-77.417-131.763c-59.543-62.106-130.827-113.013-211.743-154.175c-2.731-1.324-5.527-2.515-8.276-3.807m-302.636 19.483c6.846 0 13.638.274 20.361.732l-58.081 100.561c-81.514 16.526-142.676 85.88-142.676 168.897c0 20.854 3.841 40.819 10.913 59.325c.008.021-.008.053 0 .074l-58.228 100.854c-34.551-44.823-54.932-100.229-54.932-160.182c.001-149.255 126.524-270.262 282.643-270.261m168.969 212.035L638.013 797.271c81.076-16.837 141.797-85.875 141.797-168.603c0-20.474-4.086-39.939-10.914-58.155"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeOpen(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M779.843 599.925c0 95.331-80.664 172.612-180.169 172.612c-99.504 0-180.168-77.281-180.168-172.612c0-95.332 80.664-172.612 180.168-172.612c99.505-.001 180.169 77.281 180.169 172.612M600 240.521c-103.025.457-209.814 25.538-310.904 73.557C214.038 351.2 140.89 403.574 77.394 468.219C46.208 501.218 6.431 549 0 599.981c.76 44.161 48.13 98.669 77.394 131.763c59.543 62.106 130.786 113.018 211.702 154.179C383.367 931.674 487.712 958.015 600 959.48c103.123-.464 209.888-25.834 310.866-73.557c75.058-37.122 148.243-89.534 211.74-154.179c31.185-32.999 70.962-80.782 77.394-131.763c-.76-44.161-48.13-98.671-77.394-131.764c-59.543-62.106-130.824-112.979-211.74-154.141C816.644 268.36 712.042 242.2 600 240.521m-.076 89.248c156.119 0 282.675 120.994 282.675 270.251c0 149.256-126.556 270.25-282.675 270.25S317.249 749.275 317.249 600.02c0-149.257 126.556-270.251 282.675-270.251"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Facebook(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v1200h1200V0zm863.232 156.592c8.715-.185 17.791.098 27.173.732c34.476.047 70.483 3.155 106.201 6.299l-3.882 142.09h-95.947c-44.988-.996-61.235 16.473-62.695 67.236V484.57h162.524l-6.445 152.197H834.082v423.706H675.513V636.768H565.43V484.57h110.083V353.906c0-94.209 39.829-154.174 118.286-184.57c20.149-7.928 43.288-12.189 69.433-12.744"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FacetimeVideo(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 134.81v930.381h930.381V673.936L1200 1065.19V134.81L930.381 526.064V134.81z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastBackward(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 1200V0h200v550L700 50v500l500-500v1100L700 650v500L200 650v550z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastForward(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1200 0v1200h-200V650l-500 500V650L0 1150V50l500 500V50l500 500V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Female(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C387.136 0 214.538 172.598 214.538 385.462c0 187.751 134.31 344.181 312.055 378.479V888.02H361.278v146.74h165.314V1200h146.74v-165.24h165.314V888.02H673.333V763.94c177.78-34.269 312.129-190.702 312.129-378.479C985.462 172.598 812.864 0 600 0m0 153.278c128.231 0 232.184 103.953 232.184 232.184S728.232 617.647 600 617.647S367.816 513.693 367.816 385.462S471.769 153.278 600 153.278"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func File(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M400.86 0L174.791 218.249v73.93h302.582V0zm179.219 0v391.27H174.791V1200h850.418V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0M496.729 288.794h39.697v151.538H379.468v-38.306zm92.944 0h230.859v622.412H379.468V491.748h210.205z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileEdit(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M284.506 0L58.436 218.249V1200h850.418V817.511L803.861 922.504v173.167H163.428V292.179h197.59v-187.85h442.844v272.996L490.433 690.753L397.32 977.619l286.793-93.188l457.452-457.452L947.885 233.3l-39.031 39.031V0zm254.402 739.154l96.803 96.876l-143.434 46.557z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileEditAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0M456.885 220.898h394.482v172.046l24.683-24.609l122.388 122.314l-289.014 289.087l-181.201 58.887l58.813-181.274l198.047-198.047V286.816H505.225v118.726H380.42v507.642h404.663V803.76l66.284-66.357v241.699H314.062V358.813zm160.766 466.993l-29.443 90.674l90.601-29.443z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileNew(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M255.583 0L29.513 218.249V1200H635.49v-104.328H134.506V292.179h197.589v-187.85h442.844v312.986h104.992V0zm472.725 508.73v249.091H479.145V950.91h249.163V1200h193.016V950.91h249.164V757.821H921.323V508.73z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileNewAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0M448.975 221.924H814.6v244.409h-61.45V283.008H493.799v110.083H378.076v470.581h293.408v61.084H316.553V349.731zM725.83 519.873h113.013v145.898h145.972v113.086H838.843v145.898H725.83V778.857H579.932V665.771H725.83z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Film(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M70.825 0v1200h1058.35V0zm53.98 72.07H254.81v98.877H124.805zm820.312 0h130.005v98.877H945.117zm-634.424 7.325h578.613v523.682H310.693zM124.805 293.628H254.81v98.877H124.805zm820.385 0h130.005v98.877H945.19zM124.805 515.186H254.81v98.877H124.805zm820.385 0h130.005v98.877H945.19zM310.693 639.697h578.613v523.682H310.693zm-185.888 97.046H254.81v98.877H124.805zm820.385 0h130.005v98.877H945.19zM124.805 958.301H254.81v98.877H124.805zm820.385 0h130.005v98.877H945.19z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Filter(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 0l415.869 415.869V1200l368.262-301.318V415.869L1200 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fire(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M381.64 1200C135.779 1061.434 71.049 930.278 108.057 751.148c27.321-132.271 116.782-239.886 125.36-371.903c38.215 69.544 54.183 119.691 58.453 192.364C413.413 422.695 493.731 216.546 498.487 0c0 0 316.575 186.01 337.348 466.98c27.253-57.913 40.972-149.892 13.719-209.504c81.757 59.615 560.293 588.838-64.818 942.524c117.527-228.838 30.32-537.611-173.739-680.218c13.628 61.319-10.265 290.021-100.542 390.515c25.014-167.916-23.8-238.918-23.8-238.918s-16.754 94.054-81.758 189.065c-59.36 86.762-100.49 178.847-23.257 339.556"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flag(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 75.63V1200h159.302V650.952C606.706 528.936 893.764 1025.558 1200 718.693V75.63c-475.667 308.371-726.319-274.04-1200 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlagAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0M410.376 261.328c181.004.811 314.621 181.414 529.028 42.7v362.988c-173.222 173.222-335.647-107.109-588.721-38.232v309.888h-90.088V304.028c54.425-31.422 103.643-42.906 149.781-42.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flickr(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v1200h1200V0zm354.932 373.828c124.93 0 226.172 101.242 226.172 226.172S479.861 826.245 354.932 826.245S128.76 724.93 128.76 600s101.242-226.172 226.172-226.172m490.136 0c124.93 0 226.172 101.242 226.172 226.172S969.998 826.245 845.068 826.245S618.896 724.93 618.896 600s101.243-226.172 226.172-226.172"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folder(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m249.17 117.7l-83.716 108.032H0V1082.3h1200V225.732H591.65L507.935 117.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderClose(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m249.17 117.7l-83.716 108.033H0v144.653h1200V225.733H591.65L507.935 117.7zM0 410.67v671.63h1200V410.67z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOpen(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m249.17 117.7l-83.716 108.032H0v609.972l123.572-465.318H993.75V225.732h-402.1L507.935 117.7zM0 1016.612v65.688h993.75L1200 410.669H175.134z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderSign(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0M390.527 312.012h154.468l50.024 64.453H958.3v511.523H241.699V376.465h98.804z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Font(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M335.151 763.202h435.208L552.753 199.336L335.147 763.202M.004 1192.852v-84.182h104.038L526.542 7.148h133.42l423.294 1101.521H1200v84.182H768.761v-84.182h131.834l-99.271-260.488H302.581l-99.272 260.488h130.244v84.182H0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fontsize(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1041.282 793.277c-11.946-19.626-26.879-32.213-44.799-37.762c-37.166-11.894-78.649-11.087-103.68 16c-27.979 37.062-28.68 92.604 0 124.8c43.744 41.293 121.007 20.886 148.479-19.841zm98.559-331.515C1179.946 498.455 1200 550.507 1200 617.92v389.114c-.653 10.752-9.477 19.059-19.199 19.199h-120.317c-10.753-.65-19.062-9.479-19.2-19.199v-11.521c-40.166 30.48-86.643 44.523-133.117 44.801c-52.74-1.544-98.588-20.205-134.398-53.119c-44.099-44.131-61.091-98.39-61.438-154.238c1.735-61.103 19.384-116.096 62.078-155.517c39.387-33.317 85.968-48.972 133.759-49.279c52.905 0 97.278 14.079 133.117 42.239v-37.12c0-58.879-30.72-88.318-92.158-88.318c-47.785 0-91.305 20.053-130.558 60.159c-10.422 9.223-25.115 5.606-30.72-5.12l-47.359-81.919c-4.267-8.533-2.986-16.213 3.84-23.039c29.014-29.014 62.719-50.986 101.118-65.919c95.159-33.016 222.167-31.196 294.393 32.638m-787.188-63.999l-84.479 302.075h170.237zm351.993 595.191c1.706 5.12 2.562 8.533 2.562 10.24c-.359 11.383-9.095 19.068-19.2 19.197H536.971c-9.387 0-15.359-4.691-17.92-14.079l-42.238-148.478H229.775l-42.239 148.478c-2.56 9.388-8.96 14.079-19.2 14.079H19.858c-15.959-1.327-21.994-11.124-19.199-24.317l250.876-824.308c3.414-9.387 9.813-14.08 19.2-14.08h165.119c9.386 0 15.784 4.693 19.198 14.08z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fork(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M42.203 935.926h186.061V763.958c0-54.408 26.559-114.484 77.32-164.729c50.762-50.242 126.065-104.842 249.904-191.527c124.394-87.076 199.565-135.567 233.346-165.807c33.78-30.24 30.882-25.376 30.882-69.388V0h147.863v172.507c0 66.078-27.619 132.54-80.093 179.516c-52.475 46.975-125.164 91.312-247.208 176.741c-122.601 85.82-195.159 140.381-230.651 175.512c-35.491 35.129-33.5 36.641-33.5 59.685v171.967h194.147L306.276 1200zm587.524 0h189.988V763.958c0-23.043 1.914-24.554-33.577-59.684c-23.477-23.237-65.093-56.146-124.76-99.809c7.49-5.281 13.418-9.555 21.333-15.095c43.674-30.571 75.183-51.648 107.816-73.777c41.578 31.395 73.875 58.12 99.652 83.637c50.763 50.242 77.397 110.319 77.397 164.729v171.968h190.22L893.801 1200z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Forward(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M572.728 1200V654.546L27.272 1200V0l545.455 545.454V0l600 600z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ForwardAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0M300 375.366L600 600L300 824.634zm375 0L975 600L675 824.634z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Foursquare(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v1200h1200V343.799L931.055 570.85c21.353 21.928 21.911 53.832-1.245 78.662L654.492 926.001c-23.001 22.439-54.232 23.539-81.152 0L328.491 681.152c-62.878-3.806-120.201-33.892-166.113-83.716c-40.534-47.473-60.491-107.969-57.129-175.635c5.604-62.7 33.906-119.527 83.716-165.527c43.212-36.103 96.368-56.79 157.324-58.374c69.226.551 133.452 29.644 183.911 84.961c4.228 5.074 8.279 10.397 12.085 15.894s7.384 10.746 10.767 15.82l20.288-20.288c23.001-22.44 54.232-23.539 81.152 0l81.226 81.226L1009.717 0zm1068.091 0L624.097 610.107c-4.228 5.074-9.314 7.69-15.234 7.69s-10.134-2.145-12.671-6.372c-1.691-.846-9.32-10.146-22.852-27.905s-29.988-39.775-49.438-65.991c.176-1.101.286-2.148.439-3.223c9.497-22.87 14.722-47.888 14.722-74.194c0-107.015-86.711-193.799-193.726-193.799s-193.799 86.784-193.799 193.799s86.784 193.726 193.799 193.726c23.367 0 45.761-4.139 66.504-11.719c.037-.014.036.014.073 0c.951-.186 1.894-.32 2.856-.513l183.984 148.389c2.537 2.537 5.007 3.809 7.544 3.809h3.809c4.229 0 7.644-1.271 10.181-3.809L1200 281.616V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Friendfeed(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M885.888.001c-56.654.1-117.869 10.099-180.075 43.254C635.822 18.112 561.916-3.11 459.166 4.506C226.779 26.145 138.079 143.57 124.436 311.039c-49.003 5.582-95.312 18.507-98.674 147.973c-3.216 146.63 42.672 170.527 91.615 186.759l3.529 429.876c6.916 50.756-2.57 107.279 112.753 119.774l172.641 3.529c63.192 8.234 107.033-31.882 130.363-123.304c8.198 74.317 40.861 114.999 102.203 116.245h200.838c45.398-12.1 95.197-16.082 102.166-98.636l3.529-447.484l52.828 3.529c43.146-6.248 86.124-13.375 116.283-84.556V430.853c-8.72-54.658-23.507-48.215-38.749-59.925c41.03-9.078 83.405-13.379 98.636-112.754l-3.491-109.224c1.433-68.531-36.748-107.463-132.428-132.841C994.637 9.708 942.541-.1 885.888.001M485.899 104.869c27.623-.309 65.693 4.183 145.532 29.775l2.479 129.574c-41.544-14.562-77.922-31.363-107.122-27.409c-90.36-4.244-98.837 41.911-104.644 89.7v92.178l231.702 2.479c5.576-122.17-27.277-277.91 175.119-304.844c88.354-11.753 169.529-.5 230.989 25.794l-2.479 137.046c-31.777-25.417-74.674-34.161-119.587-39.875c-124.929-17.565-96.152 94.347-102.165 176.922h166.934v122.064H835.725l2.478 553.104l-184.355-4.994v-548.11l-239.174 4.994v543.116H230.317l4.956-543.116h-89.699l-2.479-137.047l97.171 4.994l-2.478-37.359c4.354-89.286-3.449-243.239 187.885-267.222c23.578 1.623 38.741-1.524 60.226-1.764"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FriendfeedRect(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v1200h1200V0zm809.326 164.868c41.482-.074 79.594 7.105 111.694 11.792c70.058 18.582 98.021 47.087 96.973 97.266l2.563 79.98c-11.152 72.763-42.175 75.896-72.217 82.544c11.16 8.574 22.033 3.852 28.418 43.872v97.998c-22.083 52.119-53.589 57.388-85.181 61.963l-38.672-2.637l-2.563 327.686c-5.103 60.446-41.613 63.357-74.854 72.217h-147.07c-44.915-.912-68.777-30.691-74.78-85.107c-17.082 66.939-49.165 96.337-95.435 90.308l-126.416-2.637c-84.441-9.148-77.553-50.507-82.617-87.671l-2.563-314.795c-35.836-11.885-69.445-29.308-67.09-136.67c2.461-94.796 36.41-104.312 72.29-108.398c9.99-122.622 74.913-208.57 245.068-224.414c75.234-5.577 129.368 9.935 180.615 28.345c45.549-24.278 90.355-31.569 131.837-31.642m-292.895 76.758c-15.731.176-26.828 2.506-44.092 1.318c-140.096 17.56-134.361 130.254-137.549 195.63l1.831 27.393l-71.191-3.662l1.831 100.342h65.698l-3.662 397.632h134.985V562.646l175.122-3.662v401.294l134.985 3.662l-1.831-404.956H894.8v-89.429H772.559c4.402-60.462-16.62-142.354 74.854-129.492c32.886 4.184 64.257 10.54 87.524 29.15l1.831-100.269c-45.001-19.253-104.423-27.575-169.116-18.97c-148.196 19.721-124.164 133.789-128.247 223.242l-169.629-1.831v-67.456c4.252-34.991 10.449-68.806 76.611-65.698c21.38-2.895 48.023 9.405 78.442 20.068l-1.831-94.849c-58.459-18.736-86.342-22.05-106.567-21.824"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fullscreen(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v413.818l144.141-145.386L475.708 600L143.555 932.153L0 789.844V1200h413.818l-145.386-144.141L600 724.292l332.153 332.153L789.844 1200H1200V786.182l-144.141 145.386L724.292 600l332.153-332.153L1200 410.156V0H786.182l145.386 144.141L600 475.708L267.847 143.555L410.156 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gbp(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M676.062 0c105.191 0 210.383 22.119 315.576 66.352L916.385 252.46c-84.694-34.523-158.058-51.786-220.094-51.787c-42.077 0-74.444 12.139-97.101 36.413c-22.657 23.737-33.985 57.992-33.984 102.765v156.17h303.438V673.23H565.206v115.71c-.001 91.707-40.729 158.599-122.186 200.675h580.985V1200H175.994V999.326c55.563-23.735 93.594-50.978 114.093-81.727c21.039-30.748 31.558-73.095 31.559-127.04V673.23H177.612V496.02h144.033V338.235c0-108.428 30.749-191.773 92.246-250.034C475.927 29.4 563.317 0 676.062 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gift(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m750.588 371.765l230.588-145.882L755.293 0L609.412 324.706L524.706 70.588L298.824 221.176l188.234 150.588H58.824v301.178h75.294V1200h931.763V672.941h75.295V371.765zm23.756-272.854l118.195 118.195l-241.1 152.89zm-279.832 59.675l66.739 200.216l-166.848-133.477zm30.194 966.12H209.412v-475.35h315.294zm0-527.06H134.118V447.059h390.588zm465.882 527.06H675.293v-475.35h315.295zm75.293-527.06H675.293V447.059h390.588z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Github(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0m0 65.332c295.289 0 534.668 239.379 534.668 534.668S895.289 1134.668 600 1134.668S65.332 895.289 65.332 600S304.711 65.332 600 65.332M334.644 221.924c-15.228 44.26-20.507 92.298-6.885 134.253c-16.457 18.264-29.235 39.243-38.379 62.988c-24.371 77.026-18.418 175.425 34.937 234.961c18.743 20.548 44.172 37.441 76.172 50.684c31.999 13.241 73.127 21.688 123.413 25.343c-33.751 15.771-57.543 25.406-65.846 60.277c-37.646 25.156-83.427 19.153-116.602-8.203c-26.217-19.143-38.094-53.164-67.163-65.771c-1.828-1.826-7.539-3.188-17.139-4.103c-9.6-.913-17.146 1.811-22.632 8.203c-2.743 2.739-2.541 5.666.659 8.862c22.006 17.897 43.187 36.867 55.591 59.619c11.886 24.657 24.188 42.711 36.987 54.126c35.507 23.993 83.126 36.476 124.806 21.24c-4.86 29.329 8.252 75.13-1.393 101.367c-3.658 7.306-8.89 13.71-15.747 19.188c-6.111 5.962-25.06 13.42-21.24 22.56c1.828 4.108 8.219 6.631 19.189 7.544c24.151-.616 48.193-10.382 62.402-29.443c5.028-6.85 7.544-15.774 7.544-26.733V846.607c0-12.785 2.717-21.913 8.203-27.394c5.484-5.479 11.471-9.16 17.87-10.985v147.948c0 12.785-1.155 23.753-3.441 32.886c-2.286 9.134-4.324 15.536-6.152 19.189c-4.197 7.039-12.221 13.535-12.378 21.899c0 2.738 1.635 4.305 4.834 4.761c24.021-1.042 51.889-15.133 61.67-34.938c7.771-16.438 11.646-33.736 11.646-52.002V802.734h30.176v145.238c0 18.266 4.149 35.563 12.378 52.002s21.962 27.002 41.162 31.567c10.057 2.74 16.85 3.826 20.508 3.369c3.657-.457 5.291-2.021 4.834-4.761c-2.202-8.299-6.791-15.397-11.646-21.899c-6.4-8.219-9.596-25.591-9.596-52.075V808.229c6.4 1.825 12.588 5.507 18.53 10.985c5.942 5.479 8.862 14.607 8.862 27.394v112.279c0 10.959 2.515 19.885 7.544 26.733c15.283 19.437 38.314 29.297 62.401 29.443c10.973-.913 17.361-3.436 19.189-7.544c1.828-4.109.47-7.516-4.102-10.255s-10.281-6.824-17.14-12.305c-6.857-5.479-12.09-11.884-15.747-19.188c-4.104-47.109-.104-97.608-3.441-145.239c-6.589-51.979-26.597-69.079-66.504-88.989c47.543-3.653 86.633-12.376 117.261-26.074c91.558-44.484 116.875-113.258 117.261-205.443c-1.444-59.33-18.57-111.086-57.642-150.732c5.484-21.004 7.12-42.867 4.834-65.698c-2.286-22.831-7.115-43.405-14.43-61.67c-33.824 1.826-62.174 8.635-85.033 20.508c-22.857 11.873-39.31 21.957-49.365 30.176c-81.016-18.393-167.743-18.503-245.436 2.71c-39.748-32.517-89.018-50.632-135.784-53.396"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GithubText(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1023.542 297.679h114.263v127.399c-4.376 0-12.476-.435-24.299-1.312c-11.819-.873-22.981-1.314-33.491-1.314h-56.474v244.294c0 58.669 19.261 87.999 57.789 87.999c27.146 0 51.657-7.436 73.555-22.325v131.337c-32.398 17.519-71.365 26.271-116.899 26.271c-63.919 0-108.132-22.764-132.651-68.293c-18.386-34.153-27.585-87.995-27.585-161.556V425.078h1.315v-2.627l-19.702-1.309c-11.384 0-26.268 1.309-44.654 3.936V297.679h64.356v-51.223c0-24.519-1.315-44.22-3.939-59.105h152.354c-2.623 16.638-3.938 35.464-3.938 56.475zm-420.19 583.154c1.75-17.505 2.62-47.273 2.62-89.308V383.047c0-41.148-.87-69.609-2.62-85.368h148.416c-1.755 16.639-2.631 44.22-2.631 82.748v403.215c0 44.66.876 77.057 2.631 97.191zM273.631 291.138c-59.541 0-110.754 19.703-153.662 59.106c-44.651 42.031-67.017 94.571-67.017 157.617c0 42.031 12.323 81.367 36.841 118.14c21.89 34.146 46.833 56.068 74.854 65.698v2.637c-28.021 12.247-42.041 42.884-42.041 91.919c0 37.655 14.838 65.695 44.604 84.082v2.637c-81.426 27.14-122.095 77.476-122.095 151.025c0 63.926 27.151 110.782 81.445 140.552c42.897 23.645 97.588 35.449 164.136 35.449c161.985 0 243.017-67.886 243.017-203.613c0-84.933-62.647-137.04-187.865-156.299c-28.895-4.375-50.74-14.853-65.625-31.494c-11.384-11.379-17.065-22.747-17.065-34.131c0-32.396 17.497-51.213 52.515-56.47c53.412-7.872 96.951-32.679 130.664-74.268c33.71-41.593 50.61-90.371 50.61-146.411c0-17.513-3.544-36.334-10.547-56.47c22.769-5.259 40.253-10.054 52.515-14.429V291.138c-55.17 21.016-103.725 31.494-145.752 31.494c-36.78-21.014-76.628-31.494-119.532-31.494m5.273 120.776c26.27 0 46.858 10.112 61.743 30.249c12.258 18.393 18.384 40.306 18.384 65.698c0 61.291-26.714 91.919-80.127 91.919c-55.159 0-82.764-30.254-82.764-90.674c.001-64.793 27.605-97.192 82.764-97.192m9.156 534.595c73.556 0 110.376 22.358 110.376 67.017c0 47.287-33.725 70.898-101.147 70.898c-77.051 0-115.576-22.81-115.576-68.335c-.001-46.41 35.422-69.58 106.347-69.58M769.191 97.848c0 54.04-40.869 97.848-91.282 97.848c-50.414 0-91.282-43.808-91.282-97.848S627.495 0 677.909 0c50.413 0 91.282 43.808 91.282 97.848"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Glass(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 47.526l503.929 503.906v493.521H318.627v107.521h562.769v-107.521H696.094v-493.52L1200 47.526z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Glasses(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1118.948 592.892H732.471l49.909 227.792h312.254zm-652.661 0H79.809l24.315 227.792h312.253zm723.047-66.546c8.535 9.384 11.948 19.622 10.237 30.714l-33.272 307.135c-4.563 20.081-18.878 35.561-38.393 35.832h-376.24c-20.231-1.346-33.639-14.035-37.111-31.993l-60.146-267.463c-32.547-20.271-75.345-17.317-110.057-2.56l-60.147 270.022c-4.772 19.452-18.653 31.768-37.112 31.993H70.851c-21.166-1.673-35.301-16.94-37.112-35.832L.465 557.06c-1.706-11.943 1.28-22.182 8.958-30.714l432.548-222.673c20.748-9.255 41.297-.366 49.909 17.916c7.223 20.882 2.533 44.668-16.636 53.749L201.383 499.472h792.152L719.673 375.337c-20.314-11.516-25.569-34.709-17.916-53.749c10.862-19.479 32.864-25.945 51.189-17.916z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Globe(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0M451.254 162.614c21.849 11.764 47.675 22.893 77.508 33.396c29.832 10.504 59.456 15.768 88.867 15.768c21.707-.042 43.36-5.429 61.778-18.921c17.835-10.688 35.096-19.216 57.979-17.021c23.513 2.569 48.274 10.665 73.101 12.006c31.093 16.806 61.356 37.784 90.768 62.994c-14.285.84-29.436 2.344-45.402 4.446c-15.967 2.101-31.493 5.221-46.619 9.422c-15.126 4.201-28.998 9.691-41.604 16.413c-12.604 6.723-21.853 15.106-27.735 25.19c-9.243 15.126-15.747 27.946-19.528 38.45c-7.783 21.779-5.943 56.528-23.937 72.492c-3.781 2.521-6.939 5.453-9.461 8.814c-2.521 3.362-3.573 7.761-3.153 13.222c.42 5.462 3.165 13.245 8.207 23.329c2.521 5.883 4.589 13.024 6.269 21.428c15.964-.051 32.522-2.453 46.657-12.614l83.206 7.562c21.479-24.193 46.508-24.596 68.047 0c6.723 6.722 13.866 16.835 21.429 30.281l-35.296 23.936c-8.404-2.521-18.479-7.559-30.243-15.121c-5.869-3.362-11.772-7.159-17.668-11.36c-33.534-16.276-99.198 2.678-136.132 5.053c-4.815 9.632-8.696 19.063-20.137 21.429c-1.988 6.233-.068 12.919-2.546 18.922c-15.967 25.21-21.417 51.676-16.375 79.407c9.243 43.697 31.912 65.539 68.047 65.539h13.868c17.646 0 30.053.827 37.195 2.508c7.144 1.68 10.715 2.958 10.715 3.799c-4.201 10.084-5.668 18.054-4.407 23.937c4.161 18.874 17.779 33.157 16.375 53.571c-1.305 25.966-9.469 48.354-.607 74.392c10.104 25.258 22.796 52.686 32.143 78.762c3.781 7.563 9.458 12.188 17.021 13.868c15.126 2.521 34.035-7.555 56.725-30.243c16.807-18.488 26.469-38.676 28.989-60.524c3.327-19.34 16.96-36.52 21.429-56.725v-16.376c4.193-8.403 7.766-16.6 10.714-24.582c4.206-11.604 4.666-26.39 5.661-39.703c13.211-13.212 26.102-25.006 35.297-41.604c5.883-10.084 7.574-18.881 5.053-26.444c-.808-1.68-2.914-3.373-6.307-5.053l-18.922-7.561c.323-10.566 19.751-8.993 28.989-7.562l45.365-27.735c-1.681 57.143-12.81 112.406-33.396 165.768c-20.588 53.361-51.04 101.444-91.376 144.301c-53.78 58.824-118.292 100.854-193.503 126.064c-75.21 25.21-152.729 30.247-232.561 15.121c13.744-24.28 21.678-51.639 37.842-75.646c0-12.605 1.88-23.319 5.661-32.144c16.043-37.115 42.22-46.773 72.492-75.607c30.547-31.839 30.261-69.461 31.497-115.35c-.417-29.045-46.857-46.875-68.085-63.032c-49.201-33.157-80.446-81.939-150.607-67.438c-25.075 2.556-31.145 7.429-49.811-8.169l-7.561-3.8l.646-2.508l3.153-6.307c7.53-7.876-3.16-17.779-13.26-14.514c-2.101.42-4.394.646-6.915.646c-2.304-11.257-9.857-21.73-11.36-34.042c11.764 9.243 21.877 16.198 30.281 20.82c8.401 4.623 15.544 7.78 21.429 9.461c5.883 2.521 10.92 3.348 15.121 2.508c9.244-1.681 14.47-10.93 15.729-27.736c1.261-16.806.659-36.13-1.861-57.979c2.521-3.359 4.175-6.743 5.015-10.106c8.882-45.928 32.402-35.856 68.086-49.126c5.883-3.361 7.123-7.572 3.761-12.614c0-.841-.188-1.254-.607-1.254s-.646-.451-.646-1.292c19.398-9.742 30.824-29.689 42.857-47.872c-10.253-16.247-26.169-29.812-44.111-39.096c-9.622-11.868-47.398-4.592-55.471-23.937c-6.723-.841-11.76-2.119-15.122-3.799c-34.032-22.055-48.407-60.622-85.068-76.254c-14.707-1.26-29.218-1.034-43.503.646c42.856-34.453 89.51-60.092 139.931-76.9M144.947 518.085c7.563 12.604 16.811 23.958 27.735 34.042c56.699 52.086 109.994 63.124 182.751 89.477c3.358 2.521 7.983 6.73 13.868 12.613c7.927 6.017 14.635 13.222 22.721 18.921c0 4.201-.639 10.066-1.9 17.63c-1.26 7.562-1.487 19.743-.646 36.55c2.429 46.766 40.996 83.787 51.672 128.571c-9.469 58.039-9.604 115.083-16.375 172.683c-57.143-23.529-107.33-57.139-150.608-100.836c-43.277-43.697-76.698-94.524-100.228-152.508c-16.806-42.017-27.521-84.687-32.143-127.964c-4.628-43.277-3.574-86.322 3.153-129.179"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlobeAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0m-4.834 153.589c33.6 0 65.806 3.38 96.606 10.181c30.801 6.794 60.257 16.637 88.257 29.443c-4-.8-9.201-1.989-15.601-3.589c-32.061-6.359-61.663-9.154-91.26 10.767c-18.21 11.351-38.185 20.091-61.157 19.189c-27.2 0-54.578-4.569-82.178-13.77s-52.229-19.416-73.828-30.615c44-14.4 90.36-21.606 139.161-21.606m-259.497 91.553c7.625.031 14.899.887 21.899 2.637c11.2 2.8 18.852 8.632 22.852 17.432c5.601 11.199 13.365 21.601 23.364 31.201c10 9.6 20.587 17.995 31.787 25.195c6.399 1.6 11.601 3.161 15.601 4.761c11.883 4.181 18.276 15.396 31.201 15.601c9.6-3.2 17.251-.749 22.852 7.251c17.675 8.865 33.334 21.169 44.385 38.379c-10.837 18.037-23.235 37.779-43.213 47.974c.801 0 1.172.26 1.172.659s.445.586 1.245.586c3.204 5.601 1.57 9.983-4.834 13.184c-4 1.6-7.753 2.789-11.353 3.589s-7.426 1.616-11.426 2.417c-10.4 1.6-17.234 3.605-20.435 6.006c0 1.6-1.188 2.417-3.589 2.417c0 .8-2.564 3.363-7.764 7.764s-9.397 13.427-12.598 27.026c-.801 3.2-2.434 6.395-4.834 9.595c2.4 22.4 3.031 39.822 1.831 52.222s-6.215 19.792-15.015 22.192c-3.997.8-9.197-.017-15.601-2.417c-5.601-1.6-12.621-4.608-21.021-9.009s-18.17-11.421-29.37-21.021c1.639 14.528 5.92 29.485 10.767 44.385c10.291-2.021 28.009-.851 20.435 14.429c-5.422 5.492-3.94 9.22 3.589 12.012c3.194.8 6.388 3.623 9.595 8.423c12.415 1.097 25.167 1.381 39.038-1.245c11.601-1.6 23.934-1.973 37.134-1.172c13.2.8 26.237 3.81 39.038 9.009c36.484 15.252 66.089 63.336 106.201 81.006c9.995 6.399 18.205 12.79 24.609 19.189c6.399 6.399 9.595 13.605 9.595 21.606c-.066 27.75-2.509 55.992-9.595 84.009c-2.4 7.2-9.793 17.156-22.192 29.956c-12.397 12.8-26.625 24.835-42.627 36.035a83.303 83.303 0 0 0-16.772 16.772c-4 5.6-7.826 12.807-11.426 21.606s-5.42 19.201-5.42 31.201c-6.399 12-11.786 23.032-16.187 33.032s-8.152 18.136-11.353 24.536c-3.2 8.8-6.468 16.007-9.668 21.606c-31.844-1.852-61.581-13.777-91.187-31.201c3.456-56.88 14.784-113.26 20.435-169.189c1.6-5.6 1.155-10.801-1.245-15.601c-25.727-51.13-55.907-76.68-49.146-147.583c1.6-8 2.344-14.018 2.344-18.018c-10.577-10.422-25.318-18.521-35.962-31.201c-72.703-34.077-148.432-36.685-201.636-112.793c8.8-58.399 28.228-111.967 58.228-160.767s67.407-90.404 112.207-124.805c4.803-.6 9.49-.897 14.065-.878m548.73 12.89c49.6 41.6 88.827 91.973 117.627 151.172s43.14 123.984 43.14 194.385v8.423l-38.379 12.012c-16.24-4.474-33.541 7.29-40.796 22.778l20.435 7.178c3.2 1.6 5.205 3.234 6.006 4.834c4.445 11.388-.252 25.225-3.662 37.207c-6.418 8.463-16.951 16.409-25.195 25.195c-.801 4-1.172 10.017-1.172 18.018c0 12.801-5.2 28.402-15.601 46.802c1.6 5.601 1.6 10.801 0 15.601a134.023 134.023 0 0 1-10.767 29.956c-5.586 9.6-8.854 18.44-9.668 26.44c-3.2 21.601-12.784 41.586-28.784 59.985c-21.601 22.4-40.396 32.356-56.396 29.956c-7.2-1.6-12.587-6.17-16.187-13.77c-11.819-25.037-19.184-51.247-31.787-78.589c-8.522-22.713-2.332-48.874-.586-73.242c.648-18.065-11.363-34.493-15.015-52.148c-.803-6 .758-14.209 4.761-24.609c-.788-.801-4.54-2.248-11.353-4.248c-36.307-10.173-62.698 9.258-93.604-18.604c-12.399-11.199-20.609-27.173-24.609-47.974c-4.8-28.8.772-54.804 16.772-78.003c.801-2.4 1.245-4.777 1.245-7.178c0-5.601.372-9.611 1.172-12.012c3.979-.801 7.359-2.434 10.181-4.834c2.779-2.4 4.971-4.777 6.592-7.178c2.4-3.197 4.033-6.392 4.834-9.595c85.418-7.79 103.662-23.725 180.029 21.606l34.79-24.023c-8-13.6-15.021-23.557-21.021-29.956c-20.257-21.276-45.833-24.939-66.577 0l-84.009-7.251c-10.998 10.789-30.198 11.122-45.63 12.012c0-7.2-2.006-14.406-6.006-21.606c-4.8-9.601-7.178-17.178-7.178-22.778c0-5.596 1.188-9.979 3.589-13.184c25.445-33.53 20.412-35.882 30.615-80.42c3.6-10.4 10.176-23.179 19.775-38.379c22.428-31.82 107.816-49.328 152.419-53.979"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Googleplus(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v211.011c74.675-72.259 199.448-106.512 310.84-107.446h369.727l-98.438 79.395H475.928c25.667 23.617 49.471 48.991 68.408 80.566c24.489 41.642 37.345 90.585 37.793 142.017c-.225 49.503-12.142 104.751-42.92 150.952c-25.129 36.948-61.217 68.971-99.17 98.511c-24.958 25.068-46.733 45.963-48.56 85.693c.297 29.623 19.689 56.336 46.069 75.513c44.153 34.241 81.943 62.855 123.413 98.511c39.778 34.327 76.737 76.818 100.415 128.613c20.229 46.964 20.443 104.049 6.445 156.665H1200V418.286H978.662v223.901h-95.874V418.286H661.377v-92.139h221.411V103.564h95.874v222.583H1200V0zm219.141 164.941a163.752 163.752 0 0 0-64.38 14.722C112.913 200.148 82.18 233.2 70.972 283.374c-8.647 51.898-6.732 106.826 6.445 162.451c15.619 63.448 46.706 123.649 102.905 167.651c42.386 29.937 106.586 34.404 160.62 19.775c42.243-12.571 72.503-52.489 88.916-95.947c14.904-47.123 9.749-100.67-3.223-154.761c-16.857-69.077-50.391-135.771-109.424-186.182c-27.077-21.027-62.066-32.211-98.07-31.42M0 646.069v304.468c60.193-44.472 143.759-49.895 219.434-52.441c23.454 0 43.659-.393 60.718-1.245c-11.94-16.205-23.642-33.285-35.156-51.196c-25.906-41.492-19.517-92.121-.659-136.963C155.81 720.047 76.971 703.727 0 646.069m296.558 299.195c-42.025.153-82.907 7.891-124.512 13.623C107.48 971.406 46.292 997.158 0 1048.975V1200h565.503c1.939-5.889 2.489-12.562 2.49-19.189c-.437-43.961-29.634-83.997-65.259-114.478c-50.033-42.509-108.021-76.984-163.696-118.359c-14.34-1.978-28.472-2.761-42.48-2.71"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Graph(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 145.532v908.936h1200V145.532zM147.51 291.87h904.98v616.26H147.51zm769.922 48.779L763.184 608.13L690.82 457.251l-45.41-94.775l-56.689 88.403L461.938 648.34l-52.075-54.712l-37.866-39.771l-42.114 35.229l-148.902 124.293l75.073 89.941l106.787-89.136l65.698 68.994l51.489 54.053l40.283-62.769l110.303-171.899l75 156.519l48.047 100.269l55.518-96.387l209.766-363.794z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GraphAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0M251.147 335.815h697.705v528.369H251.147zm85.767 85.035v358.3h526.172v-358.3zm447.656 28.344l58.96 34.058l-121.948 211.523l-32.227 56.03l-27.979-58.301l-43.579-91.04l-64.16 99.976l-23.364 36.475l-29.956-31.421l-38.159-40.137l-62.109 51.855l-43.652-52.295l86.572-72.29l24.463-20.435l22.046 23.071l30.249 31.86l73.682-114.844l32.959-51.343l26.44 55.078l42.041 87.744z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GroupAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0m-1.216 333.625c55.343.728 101.183 45.781 116.413 103.191c5.807 23.424 6.462 47.188.608 71.998c-8.827 34.929-26.498 69.048-59.423 90.008l47.986 22.796l114.021 55.205c11.199 4.8 16.793 14.399 16.793 28.8v110.372c0 22.763 1.808 42.393-26.406 50.418H388.792c-27.134-.391-28.258-27.874-27.622-50.418V705.623c0-14.4 6.009-24.415 18.009-30.016l117.591-53.989L542.401 600c-20.8-13.6-37.202-32.383-49.202-56.383c-14.41-31.684-20.123-72.814-9.612-110.411c13.288-50.962 54.904-96.748 115.197-99.581m-195.593 50.38c17.601 0 33.587 5.215 47.986 15.615c-3.993 11.198-7.375 23.009-10.183 35.41c-2.799 12.398-4.217 25.38-4.217 38.981c0 20.001 2.796 39.199 8.396 57.6c5.599 18.399 13.61 35.217 24.013 50.418c-4.801 6.399-11.187 11.993-19.188 16.793l-88.83 40.805c-12 6.4-21.599 15.376-28.799 26.977c-7.2 11.6-10.79 24.619-10.79 39.02v110.372h-87.576c-12.705-.198-21.286-13.002-21.619-26.368V685.221c0-12 4.384-20.013 13.184-24.013L358.777 600c-34.417-21.156-51.021-59.395-52.773-101.976c.606-52.462 34.992-109.661 97.187-114.019m393.58 0c55.291.874 95.229 55.691 97.227 114.02c-.304 38.595-15.369 75.863-50.418 100.798l130.813 62.386c8.8 4.8 13.184 12.812 13.184 24.013v104.407c-.132 12.392-6.82 25.103-21.58 26.367h-90.008V705.623c0-14.4-3.59-27.419-10.79-39.02s-16.8-20.576-28.8-26.976c-37.304-17.339-80.146-29.784-108.017-58.814c20.8-32 31.193-67.601 31.193-106.802c0-24.8-4.384-49.214-13.184-73.214c14.452-9.541 31.558-16.524 50.38-16.792"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GroupIcon(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M596.847 188.488c-103.344 0-187.12 97.81-187.12 218.465c0 83.678 40.296 156.352 99.468 193.047l-68.617 31.801l-182.599 84.688c-17.64 8.821-26.444 23.778-26.444 44.947v201.102c1.451 25.143 16.537 48.577 40.996 48.974h649.62c27.924-2.428 42.05-24.92 42.325-48.974V761.436c0-21.169-8.804-36.126-26.443-44.947l-175.988-84.688l-73.138-34.65c56.744-37.521 95.061-108.624 95.061-190.197c-.001-120.656-83.778-218.466-187.121-218.466m-301.824 76.824c-44.473 1.689-79.719 20.933-106.497 51.596c-29.62 36.918-44.06 80.75-44.339 124.354c1.819 64.478 30.669 125.518 82.029 157.446L21.163 693.997C7.05 699.289 0 711.636 0 731.041v161.398c1.102 21.405 12.216 39.395 33.055 39.703h136.284V761.436c2.255-45.639 23.687-82.529 62.196-100.531l136.247-64.817c10.584-6.175 20.731-14.568 30.433-25.152c-56.176-86.676-63.977-190.491-27.773-281.801c-23.547-14.411-50.01-23.672-75.419-23.823m608.586 0c-29.083.609-55.96 11.319-78.039 26.444c35.217 92.137 25.503 196.016-26.482 276.52c11.467 13.23 23.404 23.377 35.753 30.434l130.965 62.195c39.897 21.881 60.47 59.098 60.866 100.532v170.707h140.235c23.063-1.991 32.893-20.387 33.093-39.704V731.042c0-17.641-7.05-29.987-21.163-37.045l-202.431-96.618c52.498-38.708 78.859-96.72 79.369-156.117c-1.396-47.012-15.757-90.664-44.339-124.354c-29.866-32.399-66.91-51.253-107.827-51.596"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Guidedog(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m623.974 385.737l-135.999-212.98c-7.962-15.269-4.359-33.709 8.98-42.339c15.255-8.012 32.482-4.267 41.057 8.98l135.999 214.263c7.961 15.269 4.363 33.716-8.98 42.34c-15.216 7.334-33.859 2.36-41.057-10.264m570.939 39.773c21.876 39.658-32.348 72.172-57.736 85.962c-22.238-1.718-43.406-4.068-63.509-7.057c-20.101-2.993-37.849-5.773-53.245-8.34c-17.962-3.422-34.641-6.842-50.037-10.265c-1.741 8.555-4.094 17.535-7.057 26.943c-3.009 9.408-6.646 19.245-10.905 29.509c-87.46-18.921-190.461-70.103-238.64-148.829L878.009 225.36l29.51-106.49l62.867 76.981c43.37 12.192 103.54 13.048 121.887 57.735c15.888 26.768 14.601 83.536 32.074 105.207c9.408 9.408 19.674 17.962 30.793 25.66s24.377 21.385 39.773 41.057M579.068 782.188c-38.868-.275-79.431-13.76-118.036-14.113c-22.238-2.564-30.365 7.271-24.378 29.51l53.887 216.828c-.724 34.83-19.645 61.502-47.471 66.076c-27.079 2.602-54.311-21.826-60.302-40.416l-69.283-271.998c-1.71-9.408-7.057-13.898-16.038-13.471s-14.754 4.918-17.321 13.471l-64.15 270.717c-8.458 25.404-28.958 42.031-53.886 42.338c-36.8-.646-62.21-34.828-55.169-66.717l78.264-320.752l6.415-193.734c-67.836-17.21-153.56-63.618-187.319-128.302c-9.523-20.909-2.621-41.189 15.396-51.32c21.296-9.732 45.384-.362 53.886 14.112c33.217 51.991 106.545 85.123 156.527 85.962h382.335c-23.654 118.655-36.891 237.314-33.357 361.809m93.66-339.998c64.765 84.222 158.1 137.001 257.885 157.81c-9.287 27.801-21.692 54.85-30.15 80.188c-3.85 11.547-2.78 28.439 3.207 50.68c24.586 94.191 51.62 194.754 75.698 282.262c9.377 73.299-92.42 71.738-107.773 26.943l-55.169-210.414c-15.224-25.84-53.295-31.871-61.585 0l-50.037 209.131c-25.894 64.572-118.281 21.158-107.771-24.377l47.47-195.018c.856-6.844.428-13.688-1.282-20.527c.825-118.825 4.617-240.396 29.507-356.678"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandDown(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M227.863 0h415.509c43.972 2.192 75.724 37.788 76.252 76.898v186.106c40.854 53.196 97.553 90.897 138.61 142.165c14.432 18.094 30.479 32.095 48.143 42.003c63.281 32.214 136.954 12.695 187.398 58.481c39.482 92.008-15.751 202.231-102.423 233.279c-70.78 24.404-141.366 8.051-203.878-22.294c.182 109.856-.708 219.71-1.292 329.563c-6.678 89.434-66.892 153.202-152.504 153.797c-86.275-.244-149.017-74.623-150.565-153.797v-129.24c-25.417-4.308-48.465-12.709-69.144-25.202c-53.145 7.239-105.247-9.327-142.811-41.356c-306.751-17.576-124.761-419.443-120.194-569.952V76.898C153.061 32.617 188.555.528 227.863 0m0 277.221c-2.591 120.425-184.658 531.909 76.252 492.407c23.47 43.546 86.286 60.301 126.009 33.603c40.005 40.566 90.265 52.118 129.888 12.925c0 38.771.215 77.113.646 115.024c.431 37.91.646 76.252.646 115.023c.933 44.371 33.487 75.122 72.375 75.606c45.976-1.014 75.136-35.914 75.605-75.606c.008-150.385 1.276-298.632 1.292-438.126c100.763 12.555 237.585 132.769 316.963 10.017c9.652-19.652 13.749-35.367 10.017-55.896c-1.724-3.446-3.446-5.385-5.17-5.816c-1.723-.431-4.523-1.076-8.4-1.938c-79.87-13.044-159.73-25.221-212.601-87.237c-52.343-68.263-134.749-108.514-168.013-186.752v-3.231z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandLeft(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1200 972.138V556.629c-2.192-43.973-37.788-75.725-76.898-76.253H936.995c-53.196-40.854-90.897-97.553-142.165-138.61c-18.094-14.432-32.095-30.479-42.003-48.142c-32.214-63.281-12.695-136.954-58.481-187.399c-92.008-39.482-202.231 15.751-233.279 102.423c-24.404 70.78-8.051 141.366 22.294 203.877c-109.856-.182-219.71.708-329.564 1.292C64.363 420.495.594 480.709 0 566.321c.244 86.275 74.623 149.017 153.796 150.565h129.241c4.308 25.417 12.708 48.465 25.202 69.144c-7.239 53.145 9.327 105.247 41.357 142.812c17.576 306.75 419.443 124.761 569.951 120.193h203.555c44.282-2.096 76.37-37.59 76.898-76.897m-277.222 0c-120.425 2.591-531.908 184.658-492.406-76.253c-43.545-23.47-60.301-86.285-33.603-126.009c-40.566-40.005-52.119-90.265-12.924-129.887c-38.772 0-77.114-.216-115.024-.646s-76.252-.646-115.024-.646c-44.371-.933-75.122-33.487-75.606-72.375c1.014-45.975 35.914-75.136 75.606-75.605c150.384-.008 298.632-1.276 438.126-1.292c-12.555-100.763-132.769-237.585-10.017-316.963c19.652-9.652 35.367-13.749 55.896-10.017c3.446 1.723 5.385 3.447 5.816 5.17c.431 1.723 1.076 4.523 1.938 8.4c13.044 79.87 25.221 159.73 87.237 212.601c68.263 52.343 108.514 134.749 186.752 168.014h3.231v415.508z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandRight(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 972.138V556.629c2.192-43.973 37.788-75.724 76.898-76.252h186.106c53.196-40.854 90.897-97.554 142.165-138.611c18.094-14.432 32.095-30.479 42.003-48.142c32.214-63.281 12.695-136.954 58.481-187.399c92.008-39.482 202.231 15.751 233.279 102.423c24.405 70.78 8.052 141.366-22.294 203.877c109.856-.182 219.71.708 329.563 1.292c89.434 6.678 153.202 66.892 153.797 152.504c-.244 86.275-74.623 149.017-153.797 150.565h-129.24c-4.308 25.417-12.709 48.465-25.202 69.144c7.239 53.145-9.327 105.247-41.356 142.812c-17.576 306.75-419.443 124.761-569.952 120.193H76.898C32.617 1046.939.528 1011.445 0 972.138m277.221 0c120.425 2.591 531.909 184.658 492.407-76.252c43.546-23.471 60.301-86.286 33.603-126.01c40.566-40.005 52.118-90.265 12.925-129.887c38.771 0 77.113-.216 115.023-.646c37.911-.431 76.253-.646 115.024-.646c44.371-.933 75.122-33.487 75.606-72.374c-1.014-45.976-35.914-75.137-75.606-75.606c-150.385-.008-298.632-1.276-438.126-1.292c12.555-100.763 132.769-237.585 10.017-316.963c-19.652-9.652-35.367-13.749-55.896-10.017c-3.446 1.723-5.385 3.447-5.816 5.17c-.431 1.723-1.076 4.523-1.938 8.4c-13.044 79.87-25.221 159.73-87.237 212.601c-68.263 52.343-108.514 134.749-186.752 168.014h-3.231z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandUp(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M227.863 1200h415.508c43.973-2.192 75.725-37.788 76.253-76.898V936.995c40.854-53.196 97.553-90.897 138.61-142.165c14.432-18.094 30.479-32.095 48.142-42.003c63.281-32.214 136.955-12.695 187.399-58.481c39.482-92.008-15.751-202.231-102.423-233.279c-70.78-24.404-141.366-8.051-203.878 22.294c.182-109.856-.708-219.71-1.292-329.564C779.505 64.363 719.291.594 633.679 0c-86.275.244-149.017 74.623-150.565 153.796v129.241c-25.417 4.308-48.465 12.708-69.144 25.202c-53.145-7.24-105.247 9.327-142.811 41.356c-306.751 17.576-124.761 419.443-120.194 569.951v203.555c2.096 44.282 37.59 76.37 76.898 76.899m0-277.222c-2.591-120.425-184.658-531.908 76.252-492.406c23.47-43.546 86.285-60.301 126.009-33.603c40.005-40.566 90.265-52.118 129.887-12.925c0-38.771.216-77.113.646-115.024s.646-76.252.646-115.024c.933-44.37 33.487-75.122 72.375-75.605c45.975 1.014 75.136 35.914 75.605 75.605c.008 150.385 1.276 298.632 1.292 438.126c100.763-12.555 237.585-132.769 316.963-10.017c9.652 19.652 13.749 35.367 10.017 55.896c-1.724 3.446-3.446 5.385-5.17 5.816c-1.723.431-4.523 1.076-8.4 1.938c-79.87 13.044-159.73 25.221-212.601 87.237c-52.343 68.263-134.749 108.514-168.013 186.752v3.231H227.863z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hdd(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M240.015 89.722L0 734.546l96.021 375.732H1103.98L1200 734.546L959.985 89.722zM93.75 793.945h1012.573l-65.625 256.934H159.375zm834.229 80.493c-26.508 0-47.974 21.466-47.974 47.974c0 26.509 21.466 48.047 47.974 48.047s48.047-21.538 48.047-48.047c-.001-26.508-21.54-47.974-48.047-47.974"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Headphones(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1050 937.5v-300c0-248.513-201.487-450-450-450s-450 201.487-450 450v300c-11.048 106.906-147.598 78.946-150 0v-300c0-331.388 268.65-600 600-600s600 268.612 600 600v300c-11.048 106.906-147.598 78.946-150 0m-787.5-225h75c20.7 0 37.5 16.763 37.5 37.5v375c0 20.737-16.8 37.5-37.5 37.5h-75c-20.7 0-37.5-16.763-37.5-37.5V750c0-20.737 16.8-37.5 37.5-37.5m600 0h75c20.737 0 37.5 16.763 37.5 37.5v375c0 20.737-16.763 37.5-37.5 37.5h-75c-20.737 0-37.5-16.763-37.5-37.5V750c0-20.737 16.763-37.5 37.5-37.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HearingImpaired(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1159.825 18.085c26.211 30.386 25.351 69.361 0 95.401L995.523 279.114c-30.122 26.332-68.203 25.187-94.076 0c-25.703-29.557-25.893-70.498 0-95.401l164.302-165.627c28.653-24.329 69.465-23.899 94.076-.001M339.64 749.494c28.655-24.326 69.464-23.899 94.077 0c25.217 28.723 26.389 71.743 0 95.401l-299.454 302.104c-30.122 26.332-68.204 25.188-94.077 0c-25.703-29.558-25.892-70.496 0-95.401zm288.854-219.952c-2.398-48.684-41.57-84.218-84.802-84.802c-48.011 2.651-84.196 41.673-84.801 84.802c-2.052 30.882-22.076 51.302-50.351 51.676c-31.261-2.296-50.003-23.781-50.351-51.676c1.746-51.925 20.345-98.006 54.326-132.502c37.862-35.457 83.352-53.979 131.177-54.326c51.85 1.685 96.85 20.496 131.176 54.326c35.505 38.095 53.978 84.438 54.326 132.502c-2.052 30.882-22.076 51.302-50.351 51.676c-31.26-2.296-50.002-23.781-50.349-51.676m-84.802-382.93c101.718 2.181 195.905 42.973 262.354 109.313c69.994 74.183 107.973 170.781 108.651 264.341c-2.101 85.07-17.161 125.399-56.976 190.803c-21.731 51.807-71.849 90.229-98.714 137.139c-15.229 48.481-4.017 101.299-5.963 155.027c-.726 59.921-24.187 126.913-65.588 159.664c-46.38 27.034-94.792 37.101-137.14 37.101c-35.225-2.742-57.584-28.229-59.625-59.626c-1.146-16.341 5.282-31.612 15.899-42.4c26.689-27.214 52.868-16.911 85.464-25.175c11.771-3.74 28.769-8.201 33.126-20.538c7.502-23.134 7.261-51.45 7.287-70.889c2.199-49.368-7.059-94.804 3.312-141.114c7.11-29.836 18.804-58.064 36.438-80.826c11.915-15.676 25.548-29.46 35.775-43.063c14.716-18.206 26.64-37.903 38.426-56.312c12.715-19.453 22.673-38.744 34.45-57.639c9.455-27.191 13.176-55.696 13.25-82.151c-2.398-69.875-27.811-131.796-73.538-178.214c-50.965-47.665-112.463-73.065-176.89-73.539c-69.802 2.337-130.642 27.961-176.89 73.539c-47.712 51.196-73.063 113.548-73.539 178.214c-2.56 35.69-26.989 60.506-59.625 60.951c-35.86-2.689-60.506-28.345-60.951-60.951c2.316-102.082 42.857-196.705 108.651-264.34c73.014-70.834 169.269-108.64 262.356-109.315"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heart(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1176.629 250.347c54.502 168.401 8.89 339.761-87.232 468.872c-63.446 87.553-139.273 163.012-216.796 228.983c-71.322 66.39-230.933 197.753-273.241 201.402c-37.394-7.148-79.353-49.433-109.039-71.196C323.503 951.599 143.93 797.388 52.878 628.779c-76.34-161.871-76.48-362.086 42.333-486.189C249.271 3.702 481.533 30.841 599.359 175.944c31.644-41.05 70.556-73.335 116.737-96.853c187.213-74.728 381.972 1.418 460.533 171.256"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0m132.789 343.503c71.295-1.114 135.772 37.646 166.337 103.724c28.273 87.356 4.612 176.225-45.251 243.199c-32.912 45.417-72.247 84.584-112.462 118.807c-36.997 34.439-119.808 102.591-141.755 104.483c-19.397-3.708-41.173-25.678-56.573-36.968c-86.534-65.781-179.667-145.742-226.899-233.207c-39.601-83.97-39.673-187.864 21.96-252.241c79.917-72.048 200.39-57.946 261.512 17.325c16.415-21.295 36.605-38.066 60.562-50.267c24.279-9.69 48.803-14.483 72.569-14.855"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartEmpty(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M832.486 46.17c-48.316.497-97.873 9.932-146.357 29.285l-2.201.934l-2.201 1.134c-30.735 15.652-58.748 35.232-83.651 58.436c-134.023-115.917-347.394-124.33-493.37 7.271l-1.468 1.334l-1.334 1.401C-26.152 279.721-23.153 485.712 53.607 648.474l.667 1.268l.667 1.334c90.404 167.407 259.216 309.652 413.054 426.596l.4.333l.333.268c9.897 7.256 26.915 22.257 46.562 37.022c19.646 14.766 41.839 30.948 74.312 37.156l7.271 1.4l7.338-.601c34.082-2.939 47.452-15.523 69.843-29.617c22.391-14.095 47.177-31.986 72.645-51.565c50.202-38.595 102.147-83.079 136.351-114.804c.44-.409 1.166-.997 1.601-1.401c.071-.06.13-.14.2-.2c72.052-61.482 143.423-132.361 203.926-215.732c.019-.025.048-.041.065-.066c.114-.153.22-.313.334-.467c95.955-129.193 142.502-303.295 86.987-474.825l-1.001-3.069l-1.334-2.869C1111.497 123.881 977.435 44.681 832.486 46.17m1.735 105.532c104.434-2.402 195.973 53.518 241.415 149.158c41.99 133.897 5.693 270.32-72.512 375.364l-.2.267l-.199.268c-54.48 75.18-120.459 140.988-188.85 199.189l-.867.801l-.867.8c-30.797 28.667-82.925 73.323-130.547 109.935c-23.812 18.306-46.64 34.665-64.373 45.828c-6.668 4.197-13.008 7.059-18.212 9.272c-5.469-2.988-12.41-7.196-19.611-12.607c-14.5-10.897-29.94-24.809-46.763-37.224c-.262-.199-.538-.4-.8-.601c-149.294-113.553-306.55-251.435-382.169-390.44c-61.391-131.153-58.712-288.132 28.017-380.368c117.926-104.852 293.123-82.057 380.5 25.549l42.56 52.366l41.159-53.433c23.682-30.723 51.943-54.21 86.253-72.044c35.963-14.049 71.694-21.29 106.066-22.08"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Home(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0L56.645 422.323V1200h373.829V730.541h339.054V1200h373.828V422.323z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0m0 276.489l292.969 227.71v419.312H691.406V670.386H508.594v253.125H307.031V504.199z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hourglass(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M109.217 0v52.567C106.194 455.64 53.525 319.939 333.036 600c-284.6 280.454-223.819 149.32-223.819 547.434V1200h981.519v-52.566c2.99-402.972 55.98-268.045-223.408-547.434c284.422-279.638 223.408-149.581 223.408-547.433V0zM214.35 105.133h771.252c14.357 296.43-18.396 245.5-230.39 457.495L718.251 600l36.961 37.372c236.757 234.1 230.39 153.272 230.39 457.494H214.35c-14.399-296.413 18.597-245.734 230.39-457.905L481.701 600l-36.961-36.961C208.141 328.703 214.35 409.382 214.35 105.133M321.126 321.15c90.783 77.596 218.086 111.986 244.764 265.708h50.924C614.68 448.637 770.24 398.087 878.416 321.15c-214.421 23.11-301.874 35.51-557.29 0m276.797 439.013c-140.022-.882-279.086 70.906-276.386 218.07h557.289c.103-143.634-140.88-217.187-280.903-218.07"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Idea(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M567.663 0v190.423h64.679V0h-64.685zm-264.11 57.225l-52.992 37.103l109.203 155.946l52.963-37.104zm592.886 0L787.268 213.171l52.971 37.104L949.44 94.328l-52.992-37.103zm-296.45 185.299c-158.227 0-286.493 96.083-286.493 214.625l162.772 492.948h247.47l162.758-492.948c0-118.54-128.258-214.625-286.492-214.625zM85.465 299.673l-22.099 60.814l178.849 65.114l22.181-60.785l-178.935-65.143zm1029.062 0l-178.936 65.148l22.106 60.792l178.936-65.125zM255.756 577.681l-183.9 49.326l16.686 62.431l183.9-49.255l-16.683-62.502zm688.48 0l-16.674 62.501l183.9 49.247l16.674-62.432l-183.9-49.318zM472.66 986.032v85.686h254.687v-85.673H472.661zm0 128.282V1200h254.687v-85.672H472.661z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IdeaAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0m-20.361 221.777h40.723v120.044h-40.723zm-166.553 36.035l68.848 98.291l-33.398 23.438l-68.848-98.291zm373.828 0l33.398 23.438l-68.848 98.291l-33.398-23.438zM600 374.634c99.754 0 180.615 60.549 180.615 135.278L678.003 820.679H521.997L419.385 509.912c0-74.73 80.867-135.278 180.615-135.278m-324.39 36.035l112.793 41.089l-13.989 38.306l-112.72-41.089zm648.78 0l13.916 38.306l-112.793 41.089l-13.916-38.306zM382.983 585.938l10.547 39.404l-115.942 31.055l-10.547-39.404zm434.034 0l115.942 31.055l-10.547 39.404l-115.942-31.055zm-297.29 257.446h160.547v53.979H519.727zm0 80.859h160.547v53.979H519.727z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Inbox(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M478.125 0v280.005h-142.31L600 676.318l264.185-396.313h-142.31V0zM0 621.753V1200h1200V621.753H805.811c0 113.627-92.184 205.811-205.811 205.811S394.189 735.38 394.189 621.753z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0m-73.242 216.064h146.484v168.237h85.474L600 622.412l-158.716-238.11h85.474zM239.502 589.6h236.865c0 68.266 55.367 123.706 123.633 123.706s123.633-55.44 123.633-123.706h236.865v347.461H239.502z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxBox(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m93.75 637.5l277.837-375h456.825l277.838 375H900l-150 150H450l-150-150zm235.538-450L0 637.5v375h1200v-375l-329.287-450z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IndentLeft(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1200 100.232V276.38H0V100.306zM0 374.744l337.939 225.293L0 825.33zm1200 1.318v175.561H506.25V376.062zm0 273.706V825.33H506.25V649.768zM0 923.693l1200 .073v176.001H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IndentRight(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 100.231v176.147h1200V100.305zm1200 274.512L862.061 600.036L1200 825.329zM0 376.062v175.562h693.75V376.062zm0 273.706V825.33h693.75V649.768zm1200 273.926L0 923.768v176.001h1200z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoCircle(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0m-18.823 205.884c27.494.001 49.881 7.415 67.163 22.339c17.281 14.926 25.928 34.609 25.928 58.96c0 24.353-8.646 43.962-25.928 58.887c-17.282 14.926-39.669 22.413-67.163 22.412c-27.494.001-49.882-7.486-67.163-22.412c-16.496-14.925-24.756-34.534-24.756-58.887c0-24.351 8.26-44.034 24.756-58.96c17.281-14.925 39.669-22.338 67.163-22.339m85.986 262.72l23.584 27.1c-12.569 32.993-18.896 69.509-18.896 109.57v288.721c0 10.212.846 17.239 2.417 21.167c2.356 3.928 7.452 7.864 15.308 11.792c14.14 8.641 34.908 16.9 62.402 24.756l-8.203 42.407c-67.557-7.07-115.109-10.547-142.603-10.547c-27.494 0-74.587 3.477-141.357 10.547l-9.448-42.407c28.279-6.284 49.121-14.157 62.476-23.584c11.783-7.069 17.651-18.42 17.651-34.131v-294.58c0-27.493-10.65-41.235-31.86-41.235l-41.235 8.203l-9.375-47.095c28.279-11.782 65.567-23.206 111.914-34.204c46.345-10.998 82.087-16.48 107.225-16.48"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Instagram(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v342.261h355.225c54.647-70.088 140.228-115.283 236.499-115.283c96.277 0 181.926 45.188 236.572 115.283H1200V0zm960.278 50.903h112.207c39.946 0 72.07 32.197 72.07 72.144v93.97c0 39.946-32.124 72.144-72.07 72.144H960.278c-39.946 0-72.07-32.197-72.07-72.144v-93.97c0-39.946 32.124-72.144 72.07-72.144m-828.515 1.099h40.942v237.231h-40.942zm70.385 0h42.188v237.231h-42.188zm71.631 0h40.942v237.232h-40.942zm-172.705 6.372v230.859H58.813V123.047c0-14.37 4.056-27.484 12.158-39.331c8.103-11.839 18.163-20.268 30.103-25.342m490.21 202.588c-145.562 0-263.599 117.507-263.599 262.354s118.037 262.28 263.599 262.28s263.599-117.434 263.599-262.28s-118.037-262.354-263.599-262.354m0 53.247c116.065 0 210.132 93.611 210.132 209.106S707.35 732.422 591.284 732.422s-210.132-93.611-210.132-209.106s94.067-209.107 210.132-209.107M0 390.747V1200h1200V390.747H859.131c20.128 39.854 31.421 84.893 31.421 132.495c0 163.595-133.803 296.191-298.828 296.191S292.969 686.837 292.969 523.242c0-47.614 11.357-92.635 31.494-132.495z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IphoneHome(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0M392.212 342.847h415.503c27.966 1.894 49.067 22.242 49.438 48.047v416.895c-1.982 29.059-23.122 49.012-49.438 49.365H392.212c-29.04-1.98-49.012-23.074-49.365-49.365V390.894c1.602-27.831 24.771-47.717 49.365-48.047m69.946 94.629c-14.911.37-24.526 11.846-24.683 24.756V736.45c.387 14.886 11.808 24.6 24.683 24.756h275.61c14.918-.387 24.602-11.853 24.756-24.756V462.231c-.365-14.945-11.815-24.602-24.756-24.756z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Italic(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M578.299 1114.803h152.713L714.133 1200H245.546l16.879-85.197h152.713L620.897 86.001H468.185L485.867 0h468.587l-17.683 86.001H784.059z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Key(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M328.261 271.758C146.977 271.758 0 418.697 0 599.981c0 181.283 146.977 328.261 328.261 328.261c161.72 0 296.083-116.959 323.206-270.903c.306.017.605.064.912.076h126.386v182.46h139.538v-182.46h65.796v264.068h139.538V657.414H1200V517.878H647.095c-.322.026-.63.048-.95.076c-36.424-141.583-164.926-246.196-317.884-246.196m0 151.853c97.415 0 176.37 78.955 176.37 176.37c0 97.414-78.955 176.407-176.37 176.407s-176.408-78.993-176.408-176.407c0-97.415 78.994-176.37 176.408-176.37"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Laptop(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M120.264 172.705v628.052h959.473V172.705zm119.018 114.478h721.436v399.17H239.282zM0 857.886v106.201l60.938 63.208h1078.125L1200 964.087V857.886zm488.745 34.79h222.51v96.167h-222.51z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LaptopAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0M289.966 323.877h620.068v405.835H289.966zm76.904 73.975v257.959h466.26V397.852zM212.256 766.626h775.488v68.628l-39.404 40.869H251.66l-39.404-40.869zm315.82 22.485v62.183h143.848v-62.183z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lastfm(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0M390.156 322.25c40.347.104 61.878 3.768 97.062 15.969c61.92 21.473 114.651 65.88 150.875 127.031c56.368 107.527 105.136 234.599 144.781 318.281c53.083 73.153 161.504 69.769 201.5 25.438c91.953-101.92-47.67-158.69-96.75-175.125c-54.961-17.18-119.476-46.854-143.938-97.406c-44.839-127.702 37.599-211.409 143.938-208.531c32.991-4.741 95.703 14.375 136.812 80.844l-53.531 29.031c-22.696-24.026-55.044-43.973-86.125-44.25c-69.216-9.748-107.115 78.542-65.281 127.594c14.114 14.532 31.359 22.521 96.531 44.594c93.109 31.288 161.001 62.153 166.125 158.125c6.322 129.559-122.361 207.246-252.75 177.188c-52.51-14.256-95.908-51.567-123.5-106.125c-48.505-94.571-82.44-210.834-132.405-294.283c-41.936-66.639-108.827-103.227-188.625-103.125c-97.803-1.316-183.43 75.955-211.156 153.875c-20.661 83.009-.388 167.228 57.438 224.125c135.674 120.269 285.37 41.503 344.906-43.219c2.516 0 33.688 67.032 33.688 72.438c-40.51 48.755-116.558 90.961-172.406 101.375c-29.79 5.195-77.701 5.248-105.938.125C222.148 886.394 131.069 801.942 100.344 692C54.74 508.422 174.698 361.914 312.812 329.875c23.109-6.102 34.476-7.312 69-7.594a790.74 790.74 0 0 1 8.344-.031"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Leaf(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M131.652 839.52c-7.46-17.981-16.747-35.855-19.464-55.253c-35.786-250.118 168.278-452.905 353.494-529.928c185.214-77.022 398.471 45.498 524.903-86.019c31.587-35.667 63.736-69.443 111.762-69.694c20.018.993 36.688 10.401 45.207 26.999c158.034 323.688-67.753 707.493-322.728 843.865c-161.812 79.98-329.249 97.576-485.976 43.323c-43.582-10.338-78.515-55.918-123.691-56.51c-31.154 17.48-57.029 74.434-78.17 105.797c-24.625 42.738-73.658 54.038-108.937 16.64c-102.899-101.406 110.943-191.407 103.6-239.22m137.506-41.44c22.099 17.854 53.978 12.793 70.95-5.022C475.799 628.65 659.535 556.522 859.36 561.999c29.444 1.709 50.848-22.136 52.74-47.718c.592-30.381-22.788-50.927-48.975-52.742c-241.513-13.425-451.044 93.065-598.993 264.964c-18.738 23.4-15.336 54.535 5.026 71.577"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lines(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v240h1200V0zm0 480v240h1200V480zm0 480v240h1200V960z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Link(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m320.883 1200l400.642-400.664L600.234 678.07l77.836-77.836l121.266 121.289L1200 320.883L879.117 0l-400.64 400.664L599.766 521.93l-77.836 77.836l-121.266-121.289L0 879.117zm0-156.619L156.628 879.127l244.031-244.031l42.973 42.974l-78.556 78.555l78.31 78.31l78.557-78.556l42.973 42.974zm478.476-478.477l-42.974-42.974l78.557-78.555l-78.311-78.31l-78.556 78.556l-42.984-42.984l244.031-244.055l164.273 164.264l-244.052 244.033z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Linkedin(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v1200h1200V0zm294.287 235.913c47.53.535 94.35 33.325 96.387 90.088c1.016 50.475-42.971 88.921-97.632 90.088h-1.318c-47.057-.543-93.012-34.156-95.142-90.088c.671-49.913 42.627-88.904 97.705-90.088M804.199 474.39c52.255.324 101.572 15.826 142.09 57.13c42.106 46.96 55.624 111.71 57.129 177.538v299.414H830.859V729.419c-.384-52.302-18.3-115.877-87.524-117.993c-40.571.432-69.18 24.007-88.77 63.428c-5.348 12.688-6.118 27.273-6.372 41.821v291.797H475.708c.66-145.877 1.567-291.743 1.245-437.622c0-41.438-.399-69.34-1.245-83.716h172.485v73.535c14.641-20.823 30.879-40.571 52.66-56.47c29.545-21.085 65.036-29.168 103.346-29.809m-597.436 12.744h172.485v521.338H206.763z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func List(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 983.901h216.099V1200H0zm0 0h216.099V1200H0zm0-327.966h216.099v216.1H0zm0-327.968h216.099v216.099H0zM0 0h216.099v216.098H0zm317.596 983.901H1200V1200H317.596zm0 0H1200V1200H317.596zm0-327.966H1200v216.1H317.596zm0-327.968H1200v216.099H317.596zm0-327.967H1200v216.098H317.596z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0c-65.168 0-115.356 54.372-115.356 119.385c0 62.619-.439 117.407-.439 117.407h-115.87c-2.181 0-4.291.241-6.372.586h-32.227v112.573h540.527V237.378h-32.227c-2.081-.345-4.191-.586-6.372-.586H715.796s1.318-49.596 1.318-117.041C717.114 57.131 665.168 0 600 0M175.195 114.185V1200h849.609V114.185H755.64v78.662h191.382v928.345h-693.97V192.847H444.36v-78.662zM600 115.649c21.35 0 38.599 17.18 38.599 38.452c0 21.311-17.249 38.525-38.599 38.525s-38.599-17.215-38.599-38.525c0-21.271 17.249-38.452 38.599-38.452M329.736 426.27v38.525h38.599V426.27zm115.869.732v38.525h424.658v-38.525zm-115.869 144.58v38.525h38.599v-38.525zm115.869.732v38.599h424.658v-38.599zM329.736 716.895v38.525h38.599v-38.525zm115.869.805v38.525h424.658V717.7zM329.736 862.28v38.525h38.599V862.28zm115.869.806v38.525h424.658v-38.525zm-115.869 144.507v38.525h38.599v-38.525zm115.869.805v38.525h424.658v-38.525z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Livejournal(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M310.413 1189.541c-104.437-26.895-181.037-122.306-193.227-228.134c-23.08-160.204 44.718-317.021 135.177-445.663c-16.909-37.531-40.052-57.891-73.673-7.597c-61.21 84.591-105.875 180.485-135.924 280.117c-19.07-128.889 74.431-252.393 133.881-343.354c-11.486-52.463-75.317-122.891-72.83-121.87c-33.757-41.903-70.968-126.808-35.385-176.344C116.788 65.129 206.7-2.6 304.651.077c61.017 14.684 90.492 77.374 127.756 122.418l40.317 54.205C608.445 104.275 761.582 41.345 918.113 70.127c115.711 17.944 211.819 113.125 231.811 228.35c39.455 194.442-45.17 390.009-158.971 543.919c-118.005 155.854-276.178 293.805-468.855 344.16c-68.984 14.279-142.735 19.985-211.685 2.985m271.143-77.139c147.709-48.021 259.346-164.772 358.199-279.331c108.022-137.278 180.814-318.714 145.224-494.764C1059.051 225.3 954.5 138.071 838.456 133.587c-112.862-7.301-225.229 27.746-324.56 79c-2.053 16.06 23.763 67.841 48.397 40.632c105.267-54.697 234.001-79.119 348.772-41.607c74.227 27.575 127.934 99.504 130.063 179.14c13.922 168.725-89.033 321.502-208.214 431.129c-38.951 45.985-82.9 63.757-145.065 32.699c-87.803-40.233-179.218-73.61-264.773-118.166c-56.838-53.307-96.694-121.896-144.783-182.938c-16.609 20.17-32.907 55.437-47.431 81.532c-58.382 115.376-97.943 262.342-26.646 381.363c61.213 106.393 198.183 136.349 311.143 112.639c22.409-4.03 44.53-9.623 66.197-16.608m158.243-262.918c-18.934-109.047-30.83-219.678-44.642-329.213c-49.373-80.106-112.427-154.089-160.988-232.934c-51.491-68.835-104.534-144.614-159.538-203.342c-28.913-49.583-87.979-78.293-141.994-49.745c-72.089 30.272-144.603 88.57-159.963 168.942c17.892 63.098 67.167 114.26 102.268 169.411C262.946 488.505 340.641 611.204 427.65 715.27c104.838 45.795 215.022 91.796 314.688 139.569zm-291.438-174.73c-46.225-22.827-65.73-82.577-100.964-120.671c-48.272-67.828-102.071-132.563-142.295-205.558c13.167-48.973 29.363-58.326 57.817-91.87c41.876-29.772 125.264-106.629 178.191-31.077c69.985 88.861 139.687 178.67 198.528 275.384l56.825 234.239l-50.031 39.788c-50.414-28.791-198.071-100.235-198.071-100.235m-48.395-223.082c-40.997-55.332-79.269-113.359-125.588-164.303c-59.311 39.491-21.892 91.014 9.993 140.043c52.833 68.707 100.03 142.015 156.712 207.595c79.699-73.095-25.967-149.079-41.117-183.335m208.954 55.582c-14.482-50.992-57.932-92.273-86.729-137.288c-40.302-54.161-79.581-109.114-121.471-162.072c-27.248 14.682-57.617 16.023-79.293 38.744c41.114 64.485 87.902 131.433 133.357 193.777c44.235 55.052 100.032 88.492 154.136 66.839M154.511 298.603c-45.394-54.645 24.622-113.716 63.375-148.183c35.606-29.96 80.917-52.334 126.248-44.948c28.612 4.663 63.739 41.026 39.437 57.033c-59.549 39.224-135.272 66.623-175.568 134.59c-18.814 35.33-31.462 35.703-53.492 1.508m49.365-38.301c34.689-54.205 98.191-75.646 149.521-109.022c-8.29-60.952-86.97-10.274-115.606 13.074c-28.544 27.404-110.808 88.104-55.515 121.188c9.773-5.995 14.986-16.408 21.6-25.24m-97.487-39.362C95.339 159.938 148.5 107.664 196.44 77.952c27.512-16.751 121.544-49.395 112.197 2.847c-72.169 22.986-136.146 72.475-176.011 137.354c-6.076 10.58-22.031 21.739-26.237 2.787m773.587 492.222c70.807-94.021 131.018-217.054 94.355-336.569C939.06 270.382 810.743 234.948 710.629 255.36c-40.517 14.208-133.658 15.358-128.746 63.391c42.544 61.404 88.479 120.351 133.085 180.253c15.142 111.158 27.554 224.09 44.978 333.839c45.873-32.594 83.998-76.565 120.03-119.681"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lock(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M592.09.004C442.66-.708 311.458 88.89 248.073 228.885c-32.601 77.731-30.662 184.188-29.15 277.147H52.297V1200h1095.406V506.032h-181.2V374.417c0-51.599-10.623-99.226-30.469-145.532C878.015 95.938 741.52.716 592.09.004m0 220.971c87.291 2.317 150.961 67.954 154.76 152.124v132.933H438.575V374.417c8.239-91.624 66.223-155.759 153.515-153.442"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0m-4.321 274.658c81.026.386 155.088 52.056 186.548 124.146c10.762 25.109 16.479 50.903 16.479 78.882v71.339h98.291v376.317H303.003V549.023h90.381c-.819-50.406-1.856-108.07 15.82-150.221c34.37-75.909 105.448-124.53 186.475-124.144m-4.395 119.824c-44.881.944-74.48 35.073-78.81 83.202v71.339h167.14v-72.07c-2.061-45.641-36.604-81.214-83.937-82.471a93.24 93.24 0 0 0-4.393 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Magic(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m338.611 0l-59.2 278.697L1.318 344.979l279.788 58.946l66.542 277.08l59.129-278.768l278.142-66.188l-279.766-58.969zM1072.8 12.671L854.259 230.348l119.2 118.733l218.518-217.685zM774.776 309.509L0 1081.265L119.19 1200l774.776-771.756l-119.2-118.735zm240.236 138.212L983.389 596.62l-148.611 35.462l149.481 31.477l35.529 148.032l31.601-148.898L1200 627.23l-149.482-31.479zM771.247 718.285L736.4 882.49l-163.836 38.977l164.8 34.711l39.177 163.244l34.848-164.205l163.835-39.022l-164.776-34.735z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Magnet(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C293.176 0 43.523 234.737 43.523 523.305c0 15.136.684 30.437 2.042 45.639L56.5 798.955l351.616-28.87l-10.936-229.866c-.474-5.497-.802-11.281-.802-16.914c0-105.6 91.37-191.373 203.621-191.373S803.62 417.773 803.62 523.305c0 5.633-.259 11.351-.802 16.914l-10.863 229.866l351.544 28.87l10.937-230.013a514.916 514.916 0 0 0 2.041-45.639C1156.477 234.737 906.824 0 600 0M419.636 905.905L68.312 936.671l22.965 262.017l351.324-30.766zm360.291 1.313l-22.892 262.09L1108.359 1200l22.893-262.017z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Male(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M727.498 0v167.082h182.347L669.487 405.84c-67.443-43.645-147.834-68.973-234.15-68.973c-238.345 0-431.585 193.202-431.585 431.547C3.751 1006.76 196.992 1200 435.337 1200s431.548-193.24 431.548-431.586c0-90.596-27.903-174.67-75.601-244.09l236.321-236.321V468.75h167.082c0-146.475 1.587-322.656 1.561-468.75zM435.337 510.575c142.412 0 257.839 115.427 257.839 257.839s-115.427 257.878-257.839 257.878s-257.878-115.466-257.878-257.878s115.465-257.839 257.878-257.839"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarker(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C350.178 0 147.656 202.521 147.656 452.344c0 83.547 16.353 169.837 63.281 232.031L600 1200l389.062-515.625c42.625-56.49 63.281-156.356 63.281-232.031C1052.344 202.521 849.822 0 600 0m0 261.987c105.116 0 190.356 85.241 190.356 190.356C790.356 557.46 705.116 642.7 600 642.7s-190.356-85.24-190.356-190.356S494.884 261.987 600 261.987"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarkerAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0m0 203.247c165.185 0 299.121 133.937 299.121 299.121c0 50.037-13.711 116.091-41.896 153.441L600 996.753L342.773 655.811c-31.029-41.123-41.895-98.199-41.895-153.441c.001-165.186 133.937-299.123 299.122-299.123m0 173.291c-69.503 0-125.83 56.327-125.83 125.83s56.327 125.83 125.83 125.83s125.83-56.327 125.83-125.83s-56.327-125.83-125.83-125.83"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mic(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M567.626 0C537 2.389 506.974 8.271 477.905 17.065v160.913h-65.918V43.14c-29.282 14.495-57.049 32.007-82.251 52.808c-.133.11-.233.256-.366.366v657.275c.133.11.233.256.366.366C406.208 817.3 502.917 851.18 600 851.367c99.333-.036 195.305-35.544 270.264-97.412c.135-.11.234-.256.365-.366V96.313c-.133-.11-.232-.256-.365-.366c-25.031-20.735-52.406-37.949-81.078-52.222v134.253h-65.918V17.505C694.014 8.594 663.939 2.485 633.545 0v177.979h-65.918zM199.951 525.22v336.548c24.117 22.258 50.082 42.471 77.637 60.498c63.482 41.534 135.354 71.234 212.549 85.84v-38.817h219.727v38.817c77.195-14.604 149.066-44.306 212.549-85.84c27.555-18.027 53.52-38.24 77.637-60.498V525.22h-77.637v297.876c-87.98 71.849-200.15 114.99-322.412 114.99c-122.262 0-234.432-43.143-322.412-114.99V525.22zm509.912 482.885c-35.59 6.732-72.324 10.254-109.864 10.254c-37.54 0-74.274-3.521-109.863-10.254v109.496H360.131V1200h479.738v-82.396H709.863z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0m-21.021 209.766v115.723h42.848V209.766c19.77 1.616 39.348 5.557 58.374 11.353v104.37h42.847v-87.305c18.649 9.283 36.453 20.498 52.734 33.984c.086.072.133.148.22.22v427.515c-.086.071-.133.147-.22.22c-48.755 40.238-111.174 63.33-175.781 63.354c-63.145-.121-126.043-22.153-175.781-63.354c-.086-.071-.133-.147-.22-.22V272.388c.086-.072.134-.148.22-.22c16.394-13.529 34.421-24.923 53.467-34.351v87.671h42.921V220.825c18.905-5.721 38.452-9.506 58.371-11.059M339.771 551.367h50.537v193.726c57.225 46.73 130.172 74.78 209.692 74.78s152.469-28.05 209.692-74.78V551.367h50.536v218.921c-15.686 14.476-32.614 27.604-50.536 39.331c-41.29 27.015-87.999 46.312-138.208 55.811v71.191h84.521v53.613h-312.01v-53.613h84.521V865.43c-50.209-9.499-96.918-28.796-138.208-55.811c-17.922-11.727-34.852-24.855-50.537-39.331zM528.516 865.43c23.147 4.379 47.067 6.665 71.484 6.665s48.337-2.286 71.484-6.665v-25.269H528.516z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Minus(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 430.078h1200v339.844H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusSign(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0M255.469 502.441h689.137v195.116H255.469z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Move(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 600l206.909 206.909l-.623-144.765h331.567v332.153l-142.933.623L600 1200l206.909-206.909l-144.765.622V662.146H994.3l.622 142.933L1200 600L993.091 393.091l.623 144.763H662.146V205.701l142.933-.623L600 0L393.091 206.909l144.765-.623v331.567H205.701l-.623-142.933z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Music(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m364.798 80.329l-30.419 790.778c-40.935-21.007-92.292-30.096-146.179-22.449C72.007 865.114-11.66 952.766 1.33 1044.407c12.992 91.647 117.714 152.585 233.902 136.124c103.119-14.613 180.524-85.303 187.557-165.075l.022.045c.016-.181.026-.452.042-.656c.173-2.026.271-4.08.346-6.119c4.327-82.368 30.815-708.026 30.815-708.026l652.416-67.219l-29.467 563.479c-41.867-23.303-95.68-33.684-152.264-25.665C808.507 787.768 724.845 875.4 737.836 967.054c12.986 91.639 117.709 152.587 233.901 136.107c105.313-14.906 183.777-88.319 187.895-170.171l.05-.05C1161.89 896.157 1198.7 46.799 1200 16.784z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Myspace(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m330.889 894.143l-.924 74.684H0v-74.684C0 803.62 76.231 731.11 166.756 731.11c90.525-.001 165.299 72.515 164.133 163.033m-14.193-339.729c0 82.965-66.974 150.223-149.939 150.223c-82.967 0-150.508-67.258-150.508-150.223c0-82.967 67.542-150.224 150.508-150.224c82.966 0 149.939 67.257 149.939 150.224m423.391 328.287l-1.063 86.123H358.503v-86.123c0-104.396 87.911-188.014 192.306-188.014c104.395.001 190.623 83.627 189.278 188.014m-16.366-391.78c0 95.678-77.233 173.241-172.912 173.241c-95.678 0-173.567-77.563-173.567-173.241c0-95.677 77.89-173.239 173.566-173.239c95.677 0 172.913 77.563 172.913 173.239m476.261 380.337l-1.206 97.566H767.704v-97.566c0-118.264 99.59-212.99 217.854-212.99c118.264 0 215.948 94.738 214.424 212.99m-18.54-443.828c0 108.39-87.495 196.255-195.886 196.255c-108.389 0-196.626-87.866-196.626-196.255s88.237-196.256 196.626-196.256c108.391 0 195.886 87.868 195.886 196.256"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Network(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.422 630.365C-.4 620.378.049 611.896.049 601.507v-2.748c163.322-14.011 341.241-55.15 473.665-149.787c37.996 17.409 75.363 15.034 111.208-2.748c75.104 75.855 148.807 128.574 247.13 159.405c10.067 25.652 26.086 45.35 48.054 59.091c-26.543 65.961-63.612 124.136-111.209 174.521c-70.346-50.674-163.23-13.979-194.957 59.091c-220.012-2.384-441.761-98.642-572.518-267.967m571.143 354.54c-112.313 58.005-230.856 89.276-351.474 82.451C127.796 989.072 60.567 886.74 26.135 780.151c151.522 130.23 352.912 204.549 546.43 204.754m248.503-16.49c127.807-26.659 245.244-78.05 340.488-156.657c-125.012 325.938-501.479 474.94-810.035 336.676c100.162-14.432 194.251-49.025 274.588-94.817c80.286 46.004 175.832-2.388 194.959-85.202m236.146-335.302c49.196-3.631 97.167-7.251 142.786-15.116c-.089 12.283-1.357 24.374-1.373 35.729c-85.771 109.767-214.696 184.762-343.235 219.87c47.966-58.233 83.545-122.923 108.462-188.264c39.174-5.082 71.173-23.077 93.36-52.219m21.968-87.948c-5.416-40.734-25.791-73.796-57.664-94.819c10.072-93.269 11.733-184.275 4.119-272.089c96.156 99.264 154.383 225.964 170.244 351.792c-34.781 7.329-73.682 12.368-116.699 15.116M410.559 387.133C289.275 463.55 147.263 500.671 6.914 512.185C41.964 293.143 191.16 112.112 391.337 38.09c5.438 71.134 21.91 139.81 48.054 199.257c-41.973 42.622-51.941 97.264-28.832 149.786m236.145-101.69c63.215-78.489 115.77-158.695 145.532-252.851C843.492 50 889.715 72.444 930.903 99.928c14.386 113.183 16.386 225.917 5.491 331.18c-49.729 8.487-88.823 38.744-105.717 82.45c-73.416-26.576-133.514-76.068-186.72-129.174c13.364-34.477 13.869-66.794 2.747-98.941m-127.683-81.077c-25.545-63.148-42.218-124.34-42.562-191.012c76.599-17.623 159.296-17.036 232.027-2.748c-27.786 77.786-71.688 149.88-118.073 208.876c-16.321-6.971-56.075-22.499-71.392-15.116"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Off(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M513.94 0v693.97h172.12V0zM175.708 175.708C67.129 284.287 0 434.314 0 600c0 331.371 268.629 600 600 600s600-268.629 600-600c0-165.686-67.13-315.713-175.708-424.292l-120.85 120.85c77.66 77.658 125.684 184.952 125.684 303.442c0 236.981-192.146 429.126-429.126 429.126c-236.981 0-429.126-192.145-429.126-429.126c0-118.49 48.025-225.784 125.684-303.442z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ok(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1004.237 99.152l-611.44 611.441l-198.305-198.305L0 706.779l198.305 198.306l195.762 195.763L588.56 906.355L1200 294.916z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OkCircle(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.63 0 0 268.63 0 600s268.63 600 600 600c331.369 0 600-268.63 600-600S931.369 0 600 0m0 130.371c259.369 0 469.556 210.325 469.556 469.629c0 259.305-210.187 469.556-469.556 469.556c-259.37 0-469.556-210.251-469.556-469.556C130.445 340.696 340.63 130.371 600 130.371m229.907 184.717L482.153 662.915L369.36 550.122L258.691 660.718l112.793 112.793l111.401 111.401l110.597-110.669l347.826-347.754z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OkSign(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0m257.446 281.03l124.657 124.658l-389.354 389.43L468.823 918.97L344.165 794.312l-126.27-126.344l123.853-123.853l126.27 126.343z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Opensource(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600.073 17.426C268.728 17.426 0 286.008 0 617.353c0 260.49 166.117 482.172 398.146 565.076L546.04 791.313c-74.269-23.013-128.273-92.136-128.273-173.96c0-100.646 81.661-182.307 182.308-182.307c100.646 0 182.16 81.66 182.16 182.307c0 81.888-53.952 151.147-128.273 174.106l147.896 391.115C1033.938 1099.72 1200 877.909 1200 617.353c0-331.345-268.581-599.927-599.927-599.927"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaperClip(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M471.701 1111.207C677.774 906.278 838.97 742.464 1018.27 566.612c59.195-61.825 96.687-122.337 112.472-181.532c27.626-111.707-14.849-208.44-88.793-284.137C970.914 29.909 894.947-3.635 814.047.312s-160.812 44.726-239.74 122.337L71.146 627.78l82.874 80.899l503.16-503.158c51.63-49.087 113.07-95.365 183.505-88.793c107.317 15.066 202.091 146.416 177.587 238.754c-31.071 81.924-73.905 119.289-133.189 178.571c-180.341 179.88-320.956 318.983-496.253 494.279c-65.88 60.199-108.486 76.498-169.692 19.732c-31.571-31.571-45.383-62.484-41.437-92.739c4.379-35.301 24.363-59.717 47.355-82.873l459.748-459.75c21.003-21.04 68.836-61.425 88.793-43.409c15.311 35.521-24.12 69.425-43.408 88.793l-422.26 422.259l80.899 82.874L813.06 540.961c73.069-75.365 125.566-167.46 43.409-252.566c-90.862-77.988-186.583-25.923-254.539 41.437L142.182 789.58c-47.355 47.355-74.323 98.658-80.899 153.908c-5.188 77.454 29.733 139.628 76.953 187.451c44.385 44.143 92.336 68.594 151.936 69.061c71.656-2.869 142.112-50.066 181.529-88.793"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaperClipAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0m143.555 224.194c47.14.435 91.521 21.573 133.229 63.281c46.312 47.408 72.894 107.943 55.591 177.905c-9.887 37.073-33.313 74.951-70.386 113.672c-112.296 110.136-213.272 212.743-342.335 341.09c-24.688 24.255-68.792 53.867-113.672 55.664c-37.326-.292-67.417-15.64-95.215-43.286c-29.573-29.95-51.442-68.896-48.193-117.406c4.119-34.604 21.024-66.729 50.684-96.388l287.988-287.915c42.561-42.188 102.468-74.845 159.375-26.001c51.455 53.303 18.59 111.003-27.173 158.203L467.725 827.49l-50.684-51.93L681.52 511.083c12.08-12.13 36.761-33.344 27.173-55.591c-12.499-11.283-42.438 13.995-55.592 27.173L365.186 770.654c-14.399 14.503-26.921 29.747-29.663 51.854c-2.472 18.948 6.155 38.31 25.928 58.081c38.333 35.553 65.014 25.398 106.273-12.305c109.787-109.787 197.894-196.938 310.841-309.596c37.129-37.129 63.963-60.531 83.423-111.841c15.347-57.831-44.042-140.052-111.255-149.487c-44.113-4.116-82.581 24.848-114.917 55.591L320.654 668.042l-51.855-50.61l315.088-316.406c49.433-48.607 99.479-74.14 150.146-76.611c3.167-.155 6.379-.249 9.522-.221"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Path(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1156.8 460.801c-1.741 127.978-57.59 244.043-148 318.399c-117.563 89.48-265.186 119.294-397.601 121.6v72.001c-3.547 92.429-57.571 172.648-134.399 207.999c-28.249 13.41-59.446 19.088-88 19.2c-40.263-.487-79.431-12.662-112-28.8v-208c27.473 19.356 58.955 32.385 90.4 23.199c27.814-10.322 40.547-32.075 40.8-58.399V350.4h203.2v360c46.688-.646 92.747-5.771 135.2-15.2c55.65-13.509 108.654-34.821 148.8-70.4c49.263-46.933 70.78-102.815 71.2-164c-.765-90.595-46.353-165.932-116-209.6c-78.943-45.167-166.48-60.457-250.4-60.8c-90.493 1.983-179.56 19.778-252 64.8c-73.984 47.895-113.732 127.091-114.4 205.6c4.699 56.578 19.186 135.353 57.6 171.2L161.6 771.2c-42.053-42.06-69.765-93.552-88-145.602c-17.898-55.054-30.173-109.859-30.4-164.8c2.212-133.606 61.584-254.216 155.2-332.8C316.522 35.655 464.832.706 600 0c147.994 2.516 294.258 39.785 403.2 127.2c102.159 86.84 152.695 212.531 153.6 333.601"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pause(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h500v1200H0zm700 0h500v1200H700z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0M330.835 289.38h181.714v621.24H330.835zm356.616 0h181.714v621.24H687.451z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pencil(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1169.166 190.752L1011.645 33.23C968.21-10.204 894.716-7.362 847.613 39.858c-47.104 47.103-50.181 120.715-6.628 164.148l157.521 157.522c43.435 43.434 116.928 40.594 164.149-6.627c47.105-47.22 50.064-120.596 6.511-164.149M164.978 722.374l315.044 315.044l511.976-511.857l-315.044-315.044zM0 1197.544l415.522-83.199L83.199 782.021z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PencilAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0m201.861 187.424c17.59.139 34.667 6.627 47.568 19.529l99.849 99.81c27.6 27.601 25.709 74.104-4.143 104.027c-29.925 29.925-76.502 31.704-104.026 4.18l-99.811-99.81c-27.601-27.525-25.671-74.177 4.18-104.027c15.859-15.899 36.45-23.866 56.383-23.709M637.348 319.301l199.658 199.62l-324.468 324.391L312.88 643.655zM261.056 681.459l210.6 210.601l-263.335 52.735z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Person(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M719.57 272.035c18.461 0 35.451 3.565 50.975 10.698c46.078 23.193 73.285 67.272 73.631 116.424V737.73c-8.643 53.824-70.25 39.188-71.742 0V418.037c-5.49-18.231-40.693-10.672-41.535 0v733.785c-8.951 74.928-105.773 52.934-108.242-1.258V701.23c-6.191-28.982-46.936-15.813-47.829 1.259c1.728 149.353 1.259 298.715 1.259 448.075c-9.656 74.543-106.007 55.47-108.243 1.258l-1.258-733.785c-5.643-17.838-38.263-10.996-39.019 0V737.73c-8.644 53.824-70.25 39.188-71.742 0V399.157c.871-47.056 18.117-94.197 59.156-116.424c13.426-7.133 29.788-10.699 49.087-10.699zm-1.269-153.758c0 65.323-52.955 118.278-118.278 118.278c-65.322 0-118.277-52.955-118.277-118.278C481.745 52.955 534.7 0 600.021 0c65.325 0 118.28 52.955 118.28 118.277"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Phone(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1183.326 997.842l-169.187 167.83c-24.974 25.612-58.077 34.289-90.316 34.328c-142.571-4.271-277.333-74.304-387.981-146.215C354.22 921.655 187.574 757.82 82.984 559.832C42.87 476.809-4.198 370.878.299 278.209c.401-34.86 9.795-69.073 34.346-91.543L203.831 17.565c35.132-29.883 69.107-19.551 91.589 15.257l136.111 258.102c14.326 30.577 6.108 63.339-15.266 85.188l-62.332 62.3c-3.848 5.271-6.298 11.271-6.36 17.801c23.902 92.522 96.313 177.799 160.281 236.486c63.967 58.688 132.725 138.198 221.977 157.021c11.032 3.077 24.545 4.158 32.438-3.179l72.51-73.743c24.996-18.945 61.086-28.205 87.771-12.714h1.272l245.51 144.943c36.041 22.592 39.799 66.252 13.994 92.815"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0M384.375 238.33c12.362-.729 23.536 6.66 32.007 19.775l82.031 155.566c8.637 18.434 3.729 38.172-9.155 51.343l-37.573 37.573c-2.319 3.178-3.845 6.757-3.882 10.693c14.409 55.775 58.117 107.223 96.681 142.603c38.562 35.38 80.009 83.281 133.812 94.629c6.65 1.855 14.797 2.52 19.556-1.903l43.652-44.458c15.068-11.421 36.866-16.956 52.954-7.617h.732l148.021 87.378c21.728 13.619 23.979 39.944 8.423 55.957L849.683 941.016c-15.056 15.44-35.058 20.631-54.491 20.654c-85.948-2.575-167.158-44.759-233.862-88.11c-109.49-79.653-209.923-178.446-272.975-297.803c-24.182-50.05-52.589-113.91-49.878-169.774c.242-21.016 5.928-41.605 20.728-55.151l101.953-101.953c7.942-6.758 15.799-10.111 23.217-10.549"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Photo(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 30.981v1138.037h1200V30.981zm148.096 143.556h903.809v624.097H148.096z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhotoAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0M278.027 294.653h643.945v610.693H278.027zm79.468 77.051v334.863h485.01V371.704z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Picasa(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M910.828 94.9C1088.896 205.959 1198.984 397.994 1200 600.633c0 67.94-11.465 132.908-34.395 194.904H910.828zm228.026 763.058c-94.882 190.637-281.383 312.875-484.075 333.758c-103.906 6.653-216.735-13.006-300.638-49.682V857.958zM0 600.633C2.652 373.328 132.909 172.7 326.115 73.246L554.14 277.067L21.656 754.773C7.219 702.12 0 650.739 0 600.633m290.446 507.007C177.14 1036.634 90.612 933.589 43.313 819.741l247.133-222.929zm556.688-650.955L389.809 45.22C459.643 21.745 531.65 7.251 600 7.002c86.624 0 169.002 17.834 247.134 53.503z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Picture(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 99.243v1001.514h1200V99.243zm92.065 93.604h1015.869v814.307H92.065zm209.4 117.553c-28.65 0-52.756 9.762-72.29 29.297c-19.534 19.535-29.297 43.64-29.297 72.29c0 27.349 9.763 50.777 29.297 70.312c19.534 19.535 43.64 29.297 72.29 29.297c27.348 0 50.778-9.763 70.313-29.297c19.535-19.534 29.297-42.964 29.297-70.312c0-28.65-9.763-52.755-29.297-72.29c-19.534-19.535-42.965-29.297-70.313-29.297m431.689 49.512L502.661 709.57L380.273 607.983L199.878 842.358v68.407h800.244V635.962z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pinterest(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0c356.454 6.666 597.673 280.025 600 600c-6.375 340.923-280.025 597.673-600 600c-59.191-.229-116.981-9.442-169.763-24.622c11.23-18.143 22.462-38.229 33.692-60.259c23.688-48.151 34.503-99.721 46.652-149.028c5.184-21.166 12.527-49.028 22.03-83.585c10.366 19.87 29.157 37.148 56.372 51.835c71.035 34.562 159.039 23.145 222.246-9.071c88.677-50.367 140.639-128.602 168.467-216.413c53.582-203.643-15.51-399.439-200.217-478.187c-222.716-67.084-471.27 16.639-557.883 215.768c-37.816 133.072-35.578 303.927 95.896 359.61c6.915 2.592 13.178 2.592 18.79 0c12.34-7.05 26.651-68.155 23.975-80.993c-.864-3.888-3.888-9.288-9.071-16.198c-72.663-98.369-28.379-244.763 39.524-318.144c115.467-99.179 295.177-115.283 401.08-20.734c93.934 109.584 72.6 280.13 12.311 388.77c-33.312 53.126-78 93.273-138.013 93.952c-63.17-1.398-110.815-54.854-94.601-115.334c13.779-70.979 52.015-146.255 53.132-212.527c-3.053-55.122-30.656-90.083-81.643-90.713c-81.388 10.263-114.242 86.357-115.335 155.508c2.603 33.23 4.618 67.054 19.438 94.602c-17.278 69.979-30.67 126.35-40.173 169.114c-14.567 75.873-43.132 155.229-47.3 228.726c-1.296 24.189-1.513 46.651-.648 67.387C136.638 1044.027 1.314 838.081 0 600C6.476 253.707 280.025 2.327 600 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plane(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M321 1164h120l269.28-480.06H1020s180 0 180-83.94c0-84-180-84-180-84H710.28L441 36H321l149.28 480H255.06L120 395.94H0l96.06 204L0 804h120l135.06-120.06h215.22z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Play(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m100 0l1000 600l-1000 600z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 1200C268.65 1200 0 931.35 0 600S268.65 0 600 0s600 268.65 600 600s-268.65 600-600 600M450 300.45v599.1L900 600z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayCircle(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.65 0 0 268.65 0 600s268.65 600 600 600s600-268.65 600-600S931.35 0 600 0m0 139.16c254.499 0 460.84 206.341 460.84 460.84S854.499 1060.84 600 1060.84S139.16 854.499 139.16 600S345.501 139.16 600 139.16M450 300.439V899.56L900 600z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plurk(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v1200h230.609V284.605h734.963v652.468H416.6V1200H1200V0z"/><path fill="currentColor" d="M425.599 460.097h353.982v282.206H425.599z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlurkAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="m0 359.701l11.802-14.257v45.63c-6.538-8.508-9.17-11.93-11.802-15.355zm1165.094 236.082c-31.162 9.178-59.544 1.592-88.2-11.731c-24.433-11.359-50.03-21.594-76.199-27.412c-22.188-4.931-25.14 1.787-27.675 24.057c-4.337 38.142-7.046 77.103-17.377 113.763c-11.76 41.725-39.022 74.275-81.937 89.46c-14.332 5.071-14.559 12.396-11.485 25.71c9.683 41.937 19.975 84.017 25.07 126.637c3.824 32.019.505 34.678-32.085 40.035c-29.79 4.896-60.351 5.897-90.625 6.688c-14.309.375-19.725-7.656-20.26-24.37c-1.506-47.055-5.006-94.234-11.278-140.883c-3.82-28.417-8.071-28.835-40.109-24.291c11.062 54.54 22.105 108.993 33.582 165.564c-42.317 15.67-84.266 31.608-131.246 20.937c-13.716-3.112-21.191-10.111-22.441-23.854c-3.751-41.138-7.458-82.28-11.088-123.426c-.58-6.569-.673-13.178-1.08-21.613h-39.179c5.34 30.752 10.676 60.616 15.641 90.543c1.742 10.5 2.499 21.16 3.822 31.733c4.561 36.449 4.896 38.1-31.153 44.981c-21.058 4.021-43.838 4.955-64.906 1.519c-27.653-4.506-38.475-21.586-39.388-49.644c-1.214-37.271-2.547-74.541-4.249-111.797c-.279-6.101-2.714-12.106-4.153-18.156l-5.582.47c0 14.098.27 28.205-.055 42.295c-.755 32.652-2.133 65.298-2.536 97.954c-.139 11.359-5.243 16.615-15.593 19.54c-47.242 13.358-93.677 7.234-139.738-5.608c-15.206-4.238-21.062-16.682-19.328-31.123c3.17-26.372 7.682-52.604 12.301-78.776c2.076-11.759 5.748-23.235 8.706-34.837c9.906-38.867 7.792-75.328-13.649-111.301c-26.478-44.426-46.004-92.017-40.565-145.687c1.738-17.151 4.102-34.572 9.172-50.945c4.137-13.358.266-18.605-11.145-24.46c-39.005-20.01-71.495-47.184-85.811-90.43c-21.18-63.975 8.573-132.317 75.524-170.738c-6.087 17.617-11.843 32.652-16.49 48.02c-16.403 54.231 18.823 113.271 74.732 125.205c4.771 1.017 10.52.254 15.128-1.489c5.492-2.081 11.107-5.268 15.19-9.424c63.506-64.683 146.544-86.714 230.783-106.15c44.05-10.163 87.683-24.8 132.319-29.262c143.249-14.324 229.361 24.433 298.878 141.328c16.229 27.291 29.501 56.346 44.567 85.416c39.422-9.78 76.423-19.807 99.018-61.101c25.66-46.911 74.748-53.875 108.261-20.577c20.577 20.443 23.518 44.375 7.441 68.401c-9.921 14.82-20.995 28.871-32.003 43.892c10.624 9.193 20.408 17.076 29.454 25.73c4.754 4.548 8.688 10.151 12.181 15.778c20.686 33.363 14.824 46.491-23.803 54.016c14.354-12.08 13.342-24.491 1.251-37.607c-13.522-14.672-26.713-29.645-40.606-45.106c8.802-11.661 17.159-22.845 25.637-33.936c8.436-11.035 16.995-21.977 25.372-32.789c-2.969-5.162-5.119-9.295-7.622-13.197c-27.018-42.081-61.75-42.226-89.538-.368c-3.688 5.553-6.096 12.342-10.793 16.761c-12.302 11.582-23.998 25.262-38.69 32.578c-44.434 22.121-91.898 29.711-141.84 25.684c-23.61-1.905-47.793 1.525-71.526 4.215c-5.196.587-13.1 7.919-13.632 12.815c-1.165 10.719 1.593 21.867 2.766 33.008c38.236-3.488 73.981-7.883 109.855-9.78c67.44-3.562 134.388-5.896 194.766 35.112c14.463 9.822 36.62 8.316 55.241 12.023m-245.299-41.428c1.49 24.151 4.747 46.718 3.91 69.132c-2.335 62.235-45.979 97.41-107.322 87.696c-58.215-9.213-97.359-40.544-112.997-98.774c-16.518-61.502-15.148-122.975 2.766-183.623c5.138-17.394 15.02-34.307 26.133-48.83c26.712-34.893 65.466-40.18 100.154-13.026c14.262 11.165 25.892 26.837 35.555 42.421c10.453 16.854 16.761 36.227 25.863 53.999c2.62 5.119 8.255 11.431 13.194 12.091c12.748 1.713 25.875.575 41.846.575c-4.091-10.694-6.3-17.526-9.249-24.025c-21.706-47.88-51.68-90.312-88.113-127.841c-50.578-52.099-109.837-46.382-152.079 12.967c-29.438 41.36-43.058 88.411-49.832 137.746c-10.054 73.2-4.625 143.447 38.723 207.156c31.773 46.698 75.5 72.535 130.571 79.417c53.17 6.645 93.084-15.336 118.444-63.321c17.354-32.84 21.041-68.405 20.498-104.808c-.521-34.873-.422-34.873-38.065-38.952m-62.482 417.218c6.123-11.521 12.924-18.672 11.41-22.7c-2.736-7.281-10.44-12.697-16.071-18.892c-3.684 5.522-10.824 11.466-10.222 16.468c.849 7.007 7.466 13.315 14.883 25.124M55.875 368.081l13.644 64.433c-17.818-18.634-23.602-39.844-19-63.287c1.785-.384 3.571-.767 5.356-1.146"/><path d="M219.86 220.696c-4.945 18.438-9.888 36.879-14.831 55.317c-4.989-21.12-4.548-40.778 14.831-55.317"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plus(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M430.078 0v430.078H0v339.844h430.078V1200h339.844V769.922H1200V430.078H769.922V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusSign(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0m-95.801 261.841h191.602v242.358h242.358v191.602H695.801v242.358H504.199V695.801H261.841V504.199h242.358z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Podcast(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 17.653c-331.358 0-600 268.642-600 600c0 254.238 158.227 471.312 381.507 558.643v-52.858c-195.6-84.58-332.482-279.088-332.482-505.783c0-304.348 246.629-551.18 550.975-551.18s550.975 246.832 550.975 551.177c0 226.697-136.882 421.205-332.48 505.785v52.857C1041.866 1089.01 1200 871.957 1200 617.653c0-331.358-268.642-600-600-600m-.202 119.032c-226.008 0-409.146 183.34-409.146 409.348c0 152.788 83.745 285.819 207.801 356.086v-43.78c-102.395-66.167-170.275-181.276-170.275-312.306c0-205.298 166.323-371.823 371.62-371.823c205.298 0 371.823 166.525 371.823 371.823c0 131.03-68.023 246.139-170.479 312.307v43.779c124.115-70.268 208.002-203.3 208.002-356.086c.002-226.01-183.338-409.348-409.346-409.348M600 269.233c-135.062 0-244.52 109.456-244.52 244.52c0 68.818 28.502 130.89 74.244 175.318v-57.498c-23.163-33.464-36.719-74.054-36.719-117.82c0-114.354 92.639-206.994 206.994-206.994c114.354 0 207.195 92.639 207.195 206.994c0 43.767-13.52 84.355-36.719 117.82v57.297c45.622-44.416 74.042-106.4 74.042-175.117c.003-135.064-109.455-244.52-244.517-244.52m0 127.506c-64.605 0-117.014 52.408-117.014 117.015c0 64.605 52.408 116.81 117.014 116.81s117.014-52.206 117.014-116.812S664.605 396.739 600 396.739m0 264.089c-140.271 0-142.031 51.646-142.031 51.646s-.54 150.783 12.105 229.59c12.645 78.809 65.366 240.283 65.366 240.283h129.12s52.741-162.076 65.366-240.89s12.105-228.983 12.105-228.983s-1.76-51.646-142.031-51.646"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Print(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M367.456 73.572v181.2h-47.388L40.21 373.865h1119.58L879.932 254.772h-47.388v-181.2zM0 443.444v364.381h185.742l-65.259 318.603h959.033l-65.259-318.604H1200V443.443zm291.431 137.329h617.065l93.75 457.765H197.681z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Puzzle(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M471.749 0c-94.045 0-170.273 76.305-170.273 170.442c0 46.55 18.624 88.767 48.842 119.528H.001v284.069c31.224-36.108 77.389-59.006 128.842-59.006c94.044 0 170.273 76.304 170.273 170.44c0 94.137-76.229 170.44-170.274 170.44c-51.452 0-97.617-22.897-128.842-59.005V1200h368c-38.914-31.238-63.832-79.2-63.832-133.016c0-94.14 76.229-170.441 170.274-170.441c94.044 0 170.272 76.305 170.272 170.44c0 53.815-24.918 101.776-63.832 133.017h328.253V843.945c30.828 30.945 73.48 50.07 120.59 50.07c94.045 0 170.273-76.305 170.273-170.441c0-94.14-76.23-170.441-170.273-170.441c-47.109 0-89.762 19.125-120.59 50.07V289.969H593.18c30.228-30.763 48.843-72.971 48.843-119.528C642.023 76.304 565.793 0 471.749 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Qrcode(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v545.312h545.312V0zm654.688 0v545.312H1200V0zM108.594 108.594h328.125v328.125H108.594zm654.687 0h328.125v328.125H763.281zM217.969 219.531v108.594h110.156V219.531zm653.906 0v108.594h108.594V219.531zM0 654.688V1200h545.312V654.688zm654.688 0V1200h108.595V873.438h108.594v108.595H1200V654.688h-108.594v108.595H980.469V654.688zM108.594 763.281h328.125v328.125H108.594zm109.375 108.594v110.156h110.156V871.875zm653.906 219.531V1200h108.594v-108.594zm219.531 0V1200H1200v-108.594z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Question(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M585.558 0c-97.152-.233-273.994 93.271-279.6 176.4c-5.606 83.129 11.325 88.252 19.2 92.8h91.2c20.839-47.054 46.22-74.561 112.8-74s139.612 83.846 108.8 157.6s-59.285 99.443-97.2 179.2c-37.914 79.757-50.579 200.231-.8 300.4l112.4 2c-27.82-142.988 119.439-270.381 178-358.4c58.559-88.019 64.125-121.567 64.8-194.4c-.517-69.114-25.544-138.181-80-194.4S682.71.234 585.558 0m5.599 927.601c-75.174 0-136 60.825-136 135.999c0 75.176 60.826 136.4 136 136.4c75.175 0 136-61.226 136-136.4s-60.825-135.999-136-135.999"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionSign(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0m-10.84 150.366c72.81.176 131.308 23.199 172.119 65.332c40.811 42.133 59.524 93.882 59.912 145.679c-.507 54.584-4.675 79.714-48.561 145.679c-43.886 65.966-154.223 161.419-133.373 268.579l-84.229-1.465c-37.308-75.07-27.828-165.374.586-225.146c28.414-59.772 49.711-78.979 72.803-134.253c23.092-55.274-31.621-117.72-81.519-118.141c-49.898-.421-68.904 20.18-84.521 55.444h-68.335c-5.902-3.408-18.63-7.206-14.429-69.507c4.202-62.3 136.738-132.377 209.547-132.201m4.248 695.142c56.338 0 101.88 45.615 101.88 101.953s-45.542 102.173-101.88 102.173s-101.953-45.835-101.953-102.173s45.615-101.953 101.953-101.953"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuoteAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0m-45.923 282.422l25.781 118.727c-75.541 16.721-145.005 38.468-139.38 122.826h88.77v315.968H261.841V544.629c.095-234.691 172.401-253.786 292.236-262.207m358.228 0l25.854 118.727c-75.541 16.721-145.005 38.468-139.38 122.826h88.77v315.968H620.142V544.629c.094-234.691 172.328-253.786 292.163-262.207"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuoteRight(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M681.526 1094.657c212.643-14.942 518.306-48.892 518.474-465.344v-523.97H725.496v560.61h157.559c9.98 149.693-113.285 188.346-247.329 218.017zm-635.724 0c212.644-14.942 518.307-48.894 518.474-465.344v-523.97H89.77v560.61h157.559C257.311 815.647 134.044 854.3 0 883.971z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuoteRightAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 1200c331.371 0 600-268.629 600-600S931.371 0 600 0S0 268.629 0 600s268.629 600 600 600m45.923-282.422L620.142 798.85c75.541-16.721 145.005-38.468 139.38-122.825h-88.77V360.057h267.407v295.314c-.095 234.691-172.401 253.786-292.236 262.207m-358.228 0L261.841 798.85c75.541-16.721 145.005-38.468 139.38-122.825h-88.77V360.057h267.407v295.314c-.094 234.691-172.328 253.786-292.163 262.207"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Quotes(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M518.474 105.344C305.831 120.286.168 154.236 0 570.687v523.97h474.504v-560.61H316.946C306.965 384.354 430.23 345.701 564.274 316.03zm635.724 0c-212.643 14.942-518.306 48.893-518.473 465.343v523.97h474.505v-560.61H952.672C942.689 384.354 1065.956 345.701 1200 316.03z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Random(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M935.926 42.203v186.061H763.958c-54.408 0-114.484 26.559-164.729 77.32c-50.242 50.761-104.842 126.065-191.527 249.904c-87.076 124.394-135.567 199.565-165.807 233.346c-30.24 33.78-25.376 30.882-69.388 30.882H0v147.863h172.507c66.078 0 132.54-27.619 179.515-80.093c46.975-52.475 91.312-125.164 176.742-247.208c85.82-122.601 140.381-195.159 175.512-230.651c35.129-35.491 36.641-33.5 59.685-33.5h171.967v194.147L1200 306.276zM0 228.263v147.863h172.507c44.012 0 39.148-2.975 69.388 30.805c19.456 21.734 51.507 67.826 91.49 125.915c5.419-7.773 7.973-11.521 13.708-19.716c21.78-31.114 41.563-59.187 59.838-84.79c6.36-8.91 11.688-15.939 17.714-24.259c-27.021-39.039-49.525-70.001-72.623-95.803c-46.975-52.474-113.437-80.015-179.515-80.015zm935.926 401.464v189.988H763.958c-23.043 0-24.554 1.915-59.684-33.577c-23.237-23.477-56.146-65.093-99.809-124.76c-5.281 7.49-9.555 13.418-15.095 21.333c-30.571 43.674-51.648 75.183-73.777 107.816c31.395 41.578 58.12 73.875 83.637 99.652c50.242 50.763 110.319 77.397 164.729 77.397h171.968v190.22L1200 893.801z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Record(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 1200C268.65 1200 0 931.35 0 600S268.65 0 600 0s600 268.65 600 600s-268.65 600-600 600m0-900c-165.675 0-300 134.325-300 300s134.325 300 300 300s300-134.325 300-300s-134.325-300-300-300"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reddit(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M799.146 879.54c13.596 22.015-.893 31.198-13.234 41.597c-53.096 39.525-127.408 53.974-188.446 54.201c-57.102-1.121-140.478-14.579-183.403-54.201c-14.659-11.579-23.414-30.209-9.453-42.856c14.785-12.4 28.186.729 39.706 10.714c49.688 34.624 94.54 36.267 153.15 42.228c63.417-3.087 105.899-12.523 160.084-43.486c12.874-11.192 27.546-21.851 41.596-8.197m219.092-785.961c-47.734 0-88.537 29.66-105.018 71.542l-232.56-57.029l-97.074 274.811c-68.908 1.68-133.605 11.342-194.111 28.989c-60.504 17.646-114.264 41.182-161.322 70.594c-25.21-22.689-55.474-34.043-90.768-34.043c-19.327 0-37.185 3.57-53.571 10.714c-16.386 7.144-30.898 17.029-43.503 29.636c-12.605 12.604-22.454 27.305-29.597 44.111C3.571 549.709 0 567.341 0 585.829c0 24.369 5.677 46.663 17.021 66.831s26.645 36.559 45.973 49.164a2915.41 2915.41 0 0 0-2.508 20.781c-.84 7.145-1.254 14.513-1.254 22.075c0 49.579 14.06 96.422 42.211 140.539c28.15 44.117 66.835 82.538 115.995 115.312c49.16 32.773 106.302 58.638 171.428 77.546c65.127 18.907 135.051 28.343 209.841 28.343c73.951 0 143.724-9.436 209.271-28.343c65.547-18.908 122.877-44.772 172.037-77.546c49.158-32.773 87.844-71.194 115.994-115.312c28.15-44.117 42.211-90.96 42.211-140.539c0-13.445-1.277-27.317-3.799-41.604c20.168-12.605 36.146-29.224 47.91-49.812S1200 610.197 1200 585.827c0-18.487-3.572-36.119-10.715-52.926c-7.145-16.807-17.029-31.507-29.635-44.11c-12.605-12.605-27.307-22.492-44.111-29.636c-16.807-7.144-34.891-10.714-54.217-10.714c-33.615 0-63.84 11.354-90.73 34.042c-45.377-28.571-97.068-51.691-155.053-69.34c-57.982-17.646-120.162-27.722-186.551-30.242L708.394 159.8l197.037 46.201v.418c0 62.306 50.498 112.804 112.805 112.804s112.842-50.498 112.842-112.804s-50.535-112.84-112.84-112.84M598.709 425.76c68.066 0 132.353 8.384 192.857 25.19c60.504 16.806 113.248 39.513 158.205 68.085c44.959 28.571 80.484 62.406 106.535 101.48c26.053 39.076 39.059 80.468 39.059 124.164c0 43.698-13.006 84.862-39.059 123.52c-26.049 38.654-61.574 72.265-106.535 100.836c-44.957 28.571-97.701 51.278-158.207 68.085c-60.504 16.807-124.789 25.188-192.855 25.188c-68.067 0-132.354-8.383-192.857-25.188c-60.502-16.806-113.022-39.514-157.561-68.085c-44.537-28.571-79.8-62.182-105.851-100.836c-26.051-38.656-39.096-79.82-39.096-123.52c0-43.696 13.045-85.088 39.096-124.164c26.051-39.074 61.314-72.91 105.851-101.48c44.538-28.572 97.057-51.279 157.561-68.085c60.504-16.806 124.79-25.19 192.857-25.19M781.65 594.947c-46.992 0-85.107 37.816-85.107 84.46c0 46.643 38.115 84.461 85.107 84.461c46.988 0 85.066-37.818 85.066-84.461c.001-46.643-38.075-84.46-85.066-84.46m-353.306 1.329c-46.99 0-85.068 37.817-85.068 84.461c0 46.643 38.078 84.46 85.068 84.46c46.99 0 85.106-37.817 85.106-84.46c.001-46.643-38.116-84.461-85.106-84.461m703.572 108.703a8.992 8.992 0 0 0-1.217.76a7.79 7.79 0 0 1 1.217-.76"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Redux(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M363.858 1007.552L259.669 1123.59H0l259.468-293.53zm322.77 116.038h-222.51l-66.197-112.614l134.429-149.7zM383.214 985.984L20.943 370.06h222.499l168.881 287.129L925.741 76.41H1200z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Refresh(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C308.74 0 66.009 207.555 11.499 482.812h166.553C229.37 297.756 398.603 161.719 600 161.719c121.069 0 230.474 49.195 309.668 128.613l-192.48 192.48H1200V0l-175.781 175.781C915.653 67.181 765.698 0 600 0M0 717.188V1200l175.781-175.781C284.346 1132.819 434.302 1200 600 1200c291.26 0 533.991-207.555 588.501-482.812h-166.553C970.631 902.243 801.396 1038.281 600 1038.281c-121.069 0-230.474-49.195-309.668-128.613l192.48-192.48z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Remove(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 264.84L335.16 600L0 935.16L264.84 1200L600 864.84L935.16 1200L1200 935.16L864.84 600L1200 264.84L935.16 0L600 335.16L264.84 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RemoveCircle(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.63 0 0 268.63 0 600c0 331.369 268.63 600 600 600c331.369 0 600-268.63 600-600S931.369 0 600 0m0 130.371c259.369 0 469.556 210.325 469.556 469.629c0 259.305-210.187 469.556-469.556 469.556c-259.37 0-469.556-210.251-469.556-469.556C130.445 340.696 340.63 130.371 600 130.371M435.425 305.347L305.347 435.425L469.922 600L305.347 764.575l130.078 130.078L600 730.078l164.575 164.575l130.078-130.078L730.078 600l164.575-164.575l-130.078-130.078L600 469.922z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RemoveSign(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0M411.475 262.5L600 451.025L788.525 262.5L937.5 411.475L748.975 600L937.5 788.525L788.525 937.5L600 748.975L411.475 937.5L262.5 788.525L451.025 600L262.5 411.475z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Repeat(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600c222.411 0 416.39-121.104 520.02-300.879L908.79 777.612C847.217 884.405 732.127 956.47 600 956.47c-196.873 0-356.47-159.597-356.47-356.47S403.127 243.53 600 243.53c84.387 0 161.732 29.521 222.729 78.589L665.186 434.18L1200 612.524V53.613l-174.17 123.926C917.124 67.952 766.553 0 600 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RepeatAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0m0 252.466c96.469 0 183.717 39.359 246.68 102.832l100.854-71.777v323.73L637.72 503.979l91.261-64.966c-35.33-28.42-80.104-45.482-128.979-45.482c-114.03 0-206.47 92.438-206.47 206.47S485.97 806.47 600 806.47c76.528 0 143.194-41.709 178.857-103.563l122.313 70.312C841.148 877.346 728.821 947.534 600 947.534c-191.932 0-347.534-155.604-347.534-347.534S408.068 252.466 600 252.466"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ResizeFull(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m670.312 0l177.246 177.295l-241.21 241.211l175.146 175.146l241.211-241.211L1200 529.688V0zM418.506 606.348L177.295 847.559L0 670.312V1200h529.688l-177.246-177.295l241.211-241.211z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ResizeHorizontal(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M304.102 295.898L0 600l304.102 304.102v-203.54h591.797v203.539L1200 600L895.898 295.898v203.539H304.102z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ResizeSmall(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1024.854 0l-241.26 241.26L606.348 63.965v529.688h529.688L958.74 416.406L1200 175.146zM63.965 606.348L241.26 783.594L0 1024.854L175.146 1200l241.26-241.26l177.246 177.295V606.348z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ResizeVertical(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M904.102 304.102L600 0L295.898 304.102h203.539v591.797H295.898L600 1200l304.102-304.102h-203.54V304.102z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReturnKey(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M808.969 133.929v257.06H942.94v267.899H417.981V508.763L0 787.417l417.982 278.654V915.946H1200V133.928H808.969z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Retweet(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M300 225L0 525h225v375h450L525 750H375V525h225zm225 75l150 150h150v225H600l300 300l300-300H975V300z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReverseAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 1200C268.65 1200 0 931.35 0 600S268.65 0 600 0s600 268.65 600 600s-268.65 600-600 600m-75-824.625L225 600l300 224.625zm375 0L600 600l300 224.625z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Road(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1200 1200L957.743 0H658.691l9.164 276.675H532.144L541.308 0H242.256L0 1200h501.562l11.441-345.445h173.992L698.438 1200zM683.573 751.231H516.426l13.479-406.965h140.188z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rss(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1200 1200C1200 537.258 662.742 0 0 0v240.015c530.193 0 959.985 429.792 959.985 959.985zm-480.029 0c0-397.646-322.324-719.971-719.971-719.971v239.94c265.097 0 480.029 214.934 480.029 480.029zm-479.956 0c0-132.549-107.466-240.015-240.015-240.015V1200z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Scissors(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M516.786 78.404c-89.537 0-162.129 72.592-162.129 162.129c7.674 109.36 81.438 140.506 162.946 201.74c60.609 38.349 53.187 81.094 65.404 149.848c-53.95 13.664-118.616 29.877-157.011-2.765l-91.505-65.098c-.131-.108-.277-.2-.409-.307c-27.968-22.936-63.775-36.745-102.764-36.745c-68.543 0-127.057 42.54-150.768 102.66c-3.122 12.076-5.116 12.008-15.66 11.977c-35.857 0-64.89 29.032-64.89 64.894c0 35.858 29.033 64.893 64.893 64.893c30.988-3.452 19.907-7.164 37.461 15.967c29.63 38.812 76.376 63.869 128.966 63.869c66.145 0 122.994-39.62 148.209-96.418c28.644 16.406 53.42 22.574 86.081 27.02H1200c-90.997-165.847-316.997-140.223-465.608-154.451c-30.518-75.807-37.399-175.94-108.188-226.817h-.717c32.815-29.67 53.429-72.545 53.429-120.267c-.001-89.538-72.593-162.13-162.13-162.13zm0 81.064c44.769 0 81.063 36.296 81.063 81.063c0 44.769-36.296 81.064-81.063 81.064c-44.769 0-81.064-36.296-81.064-81.064c0-44.767 36.297-81.063 81.064-81.063M231.32 568.27c44.768 0 81.064 36.296 81.064 81.064c0 44.768-36.296 81.063-81.064 81.063s-81.064-36.297-81.064-81.063c0-44.768 36.297-81.064 81.064-81.064m432.754 33.98c25.438 0 46.059 20.622 46.059 46.06s-20.621 46.06-46.059 46.06s-46.061-20.62-46.061-46.06c.002-25.438 20.623-46.06 46.061-46.06m3.48 222.007c62.491 126.303 105.737 281.315 263.46 297.339l-111.158-297.34zm-487.615 89.765v72.468h157.932v-72.468zm284.237 0v72.468h174.311l-33.367-72.468zm495.497 0l21.494 72.468H1139.1v-72.468z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Screen(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 103.564v792.188h435.718v114.186H329.517v86.498h540.967v-86.498H764.282V895.752H1200V103.564zm147.583 149.561h904.833v493.14H147.583z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScreenAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0M250.049 310.474h699.902v462.013h-254.15v66.576h61.963v50.465H442.236v-50.465h61.963v-66.576h-254.15zm86.133 87.231v287.622h527.637V397.705z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Screenshot(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M509.619 0v156.226C332.04 192.201 192.201 332.04 156.226 509.619H0v180.762h156.226C192.201 867.96 332.04 1007.8 509.619 1043.774V1200h180.762v-156.226C867.96 1007.8 1007.8 867.96 1043.774 690.381H1200V509.619h-156.226C1007.8 332.04 867.96 192.201 690.381 156.226V0zm0 276.636v175.342h180.762V276.636c112.74 31.441 201.543 120.243 232.983 232.983H748.021v180.762h175.343c-31.44 112.74-120.243 201.543-232.983 232.983V748.021H509.619v175.343c-112.74-31.44-201.543-120.243-232.983-232.983h175.342V509.619H276.636c31.44-112.74 120.243-201.542 232.983-232.983"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Search(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M672.633 0C389.326 0 159.634 229.652 159.634 512.958c0 100.662 28.986 194.563 79.083 273.79c-.368.16-.729.305-1.098.471l-223.21 223.172L204.019 1200l231.249-231.288c-.069-.326-.162-.651-.234-.979c71.035 37.19 151.865 58.224 237.6 58.224c283.309 0 512.959-229.69 512.959-512.997C1185.59 229.652 955.939 0 672.633 0m0 227.877c157.441 0 285.041 127.639 285.041 285.081s-127.6 285.081-285.041 285.081c-157.442 0-285.082-127.639-285.082-285.081S515.19 227.877 672.633 227.877"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0m41.162 260.01c160.533 0 290.625 130.166 290.625 290.698c0 160.533-130.092 290.625-290.625 290.625c-48.581 0-94.368-11.885-134.619-32.959c.041.186.106.401.146.586L375.586 939.99L268.213 832.544l126.416-126.489c.209-.094.45-.129.659-.22c-28.387-44.893-44.824-98.088-44.824-155.127c0-160.533 130.165-290.698 290.698-290.698m0 129.126c-89.213 0-161.572 72.359-161.572 161.572s72.359 161.499 161.572 161.499c89.214 0 161.499-72.286 161.499-161.499s-72.285-161.572-161.499-161.572"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Share(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v1200h1200V573.486l-196.875 208.739v220.898h-806.25v-806.25h396.68V0zm857.861 0v225.977c-205.254.005-579.542 2.254-579.542 641.895c42.436-427.736 237.375-430.415 579.542-430.42v246.776L1200 342.09z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M754.553 35.03v294.208C487.317 329.246 0 332.178 0 1164.97c55.25-556.9 309.061-560.402 754.553-560.408v321.292L1200 480.407z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingCart(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1199.398 403.537l-50.323 288.696c-6.206 31.891-31.615 51.282-60.917 51.646H354.5l-14.566 82.106h699.226c36.818 2.798 61.793 28.88 62.242 62.24c-2.678 36.743-28.758 61.786-62.242 62.242H265.773c-41.827-3.762-66.768-37.103-62.242-74.16l33.107-180.104l-50.323-505.88L43.292 145.3C8.341 131.423-5.924 99.805 2.239 67.167c13.573-34.015 46.096-49.556 78.133-41.053l182.752 58.269c24.62 9.229 38.783 29.382 42.377 52.972l10.594 100.646l829.006 92.7c38.09 8.251 58.769 38.422 54.297 72.836m-744.126 677.619c0 52.476-42.54 95.017-95.018 95.017c-52.477 0-95.017-42.541-95.017-95.017c0-52.478 42.541-95.019 95.017-95.019c52.477.001 95.018 42.543 95.018 95.019m567.252 0c0 52.476-42.541 95.017-95.019 95.017c-52.477 0-95.017-42.541-95.017-95.017c0-52.478 42.54-95.019 95.017-95.019c52.478.001 95.019 42.543 95.019 95.019"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingCartSign(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0M297.583 253.491l106.787 33.545c14.137 4.643 23.553 17.771 25.195 31.201l6.006 58.812l483.545 53.979c20.763 4.022 35.353 22.769 32.446 42.041l-30.029 169.188c-3.822 17.697-18.479 29.828-34.79 30.029H457.178l-8.423 47.974h407.959c21.332.751 36.957 16.995 37.207 35.962c-.885 21.638-18.325 35.801-37.207 36.035H405.542c-22.756-1.882-39.462-19.915-35.962-41.968l19.189-105.615l-30.029-295.236l-82.764-26.366c-9.6-3.2-16.806-9.219-21.606-18.019c-9.082-19.032-.599-40.104 15.601-49.219c9.246-4.806 18.276-5.405 27.612-2.343m162.598 559.497c31.066 0 56.25 25.184 56.25 56.25c0 31.065-25.184 56.25-56.25 56.25c-31.064 0-56.25-25.185-56.25-56.25c0-31.066 25.184-56.25 56.25-56.25m330.175 0c31.065 0 56.25 25.184 56.25 56.25c0 31.065-25.185 56.25-56.25 56.25s-56.25-25.185-56.25-56.25c0-31.066 25.185-56.25 56.25-56.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Signal(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M960 1200V0h240v1200zM640 300h240v900H640zM320 600h240v600H320zM0 900h240v300H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Skype(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M934.32 749.374c.628-93.532-54.738-155.718-130.311-190.649c-37.215-15.905-75.342-28.896-111.692-37.873c-49.287-12.466-95.359-20.748-144.432-34.662c-39.77-9.401-88.462-39.91-89.869-75.747c9.04-63.575 86.472-77.962 133.52-78.313c70.348-2.685 114.432 32.604 145.073 84.733c22.44 37.604 40.649 59.956 82.166 60.34c45.188-2.544 80.286-38.254 80.881-79.598c-6.312-93.112-89.408-146.796-165.614-172.035c-98.416-24.873-206.476-25.613-297.85 8.987c-86.917 32.104-146.669 104.271-147.642 190.008c4.361 157.553 162.757 202.847 281.159 228.521c36.062 7.188 70.827 14.699 103.991 25.677c44.354 16.105 81.553 40.426 82.165 87.302c0 28.243-14.122 51.779-42.366 70.609c-60.434 29.738-135.052 45.752-195.145 14.122c-42.897-20.496-63.703-55.097-80.881-93.72c-17.959-39.956-40.076-65.02-82.166-65.477c-44.125 1.646-82.824 33.313-83.449 74.464c6.991 82.326 64.543 150.175 129.667 184.871c123.327 50.646 259.402 60.463 378.729 11.555c92.752-39.804 153.1-120.721 154.066-213.115m223.387-23.108c22.253 46.219 33.38 95.003 33.38 146.356c-1.959 89.621-37.483 172.712-95.646 231.731c-66.881 62.598-147.329 95.036-231.731 95.646c-55.731-.939-108.133-14.064-154.061-38.515c-107.268 18.633-232.149 3.248-323.526-35.307c-138.648-62.618-244.957-173.389-301.059-301.06c-43.081-115.005-54.947-223.053-32.737-333.794C-21.248 358.403.682 197.845 103.038 93.976C216.441-14.467 380.303-28.23 509.371 48.4c183.417-29.618 371.484 35.62 495.561 157.912c106.941 111.563 164.588 257.766 165.614 399.273c-.175 41.777-4.675 82.747-12.839 120.681"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Slideshare(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M26.576 0v247.312c-5.883-3.919-16.366-7.198-21.756-3.26c-5.195 6.009 2.19 14.774 5.447 20.378c32.284 32.494 65.367 53.375 104.349 70.388c-8.818 27.604-11.412 55.262-8.689 82.608c5.795 47.806 38.457 94.232 89.678 94.574c24.498-.561 50.865-18.215 51.097-43.477V360.9c6.216 1.774 12.296 4.395 18.479 5.444v100.539c0 7.248 1.713 13.502 5.154 18.756c10.624 15.354 27.823 24.613 45.925 24.738c53.604-2.141 83.639-46.16 89.694-94.574c2.623-28.619-1.187-57.598-8.705-82.625c35.541-17.946 100.306-44.29 111.692-85.594c.512-5.125-4.312-7.03-8.284-7.067c-6.432.264-11.637 2.655-16.438 6.24L484.221 0zm26.618 27.187h403.84v235.881c-6.732 4.273-13.25 7.104-20.103 10.326l-357.64 3.812c-9.259-3.533-17.301-9.263-26.099-13.049zm133.478 125.178c-33.32 0-60.32 27.016-60.32 60.335c0 33.321 27 60.321 60.32 60.321s60.337-27 60.337-60.321c0-33.319-27.015-60.335-60.337-60.335m142.995 0c-33.32 0-60.32 27.016-60.32 60.335c0 33.321 27 60.321 60.32 60.321s60.336-27 60.336-60.32c0-33.319-27.016-60.336-60.336-60.336"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Smiley(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v1200h1200V0zm392.285 228.223c71.542 0 129.564 57.95 129.564 129.492S463.827 487.28 392.285 487.28S262.72 429.257 262.72 357.715s58.023-129.492 129.565-129.492m403.784 0c71.542 0 129.491 57.95 129.491 129.492S867.611 487.28 796.069 487.28s-129.565-58.023-129.565-129.565s58.023-129.492 129.565-129.492M183.032 636.108h833.936c0 230.285-186.682 417.04-416.968 417.04s-416.968-186.753-416.968-417.04"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmileyAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0M286.157 286.157h627.686v627.686H286.157zm205.152 119.385c-37.424 0-67.749 30.325-67.749 67.749s30.325 67.749 67.749 67.749c37.425 0 67.822-30.325 67.822-67.749s-30.399-67.749-67.822-67.749m211.23 0c-37.424 0-67.749 30.325-67.749 67.749s30.325 67.749 67.749 67.749s67.749-30.325 67.749-67.749s-30.325-67.749-67.749-67.749M381.885 618.896c0 120.463 97.652 218.114 218.115 218.114s218.115-97.651 218.115-218.114z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Soundcloud(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M728.819 340.654c-27.849.194-55.8 4.819-82.565 14.521c-5.576 2.089-14.611 5.147-18.573 14.222l1.22 479.957c2.367 5.222 5.485 8.119 10.692 9.98h426.733c69.749 1.225 124.104-58.435 132.901-131.101c7.534-78.68-40.545-149.141-124.909-161.362c-23.964-3.457-53.819-.022-78.102 9.568c-.812.363-1.046-.553-1.67-6.305c-5.993-52.697-28.083-99.161-63.58-142.807c-48.872-53.839-125.154-87.209-202.147-86.673m-132.206 35.101c-8.142 0-15.934 115.9-17.224 347.712c19.941 177.772 17.697 182.39 36.527 0c-2.673-231.824-11.163-347.715-19.303-347.712m-51.573 27.598c-8.141.253-15.758 107.207-16.716 320.115c19.939 177.772 17.695 182.39 36.526 0c-3.005-213.916-11.671-320.367-19.81-320.115m-149.638 33.075c-.061 0-.126 0-.188.023c-7.658 1.158-14.854 96.507-16.791 287.021c19.941 177.771 17.678 182.39 36.509 0c-3.669-190.688-11.805-286.699-19.53-287.04zm50.243 8.968c-7.892-.162-15.531 92.358-17.484 278.071c19.939 177.772 17.695 182.39 36.526 0c-3.006-185.063-11.15-277.909-19.042-278.071m-99.677 3.733c-8.223.087-15.67 91.61-15.965 274.339c19.94 177.772 17.678 182.39 36.508 0c-3.337-183.073-12.32-274.425-20.543-274.339m149.281 5.44a.53.53 0 0 0-.188.022c-8.058 1.247-15.532 92.121-16.49 268.88c19.94 177.771 17.695 182.39 36.526 0c-3.312-180.327-11.846-268.746-19.849-268.897zm-197.945 11.952c-7.084-.068-13.8 84.201-17.054 256.946c19.941 177.772 17.697 182.39 36.527 0c-4.7-169.796-12.278-256.879-19.473-256.946m-47.673 47.764c-6.73-.411-13.522 68.904-18.46 209.183c19.941 177.772 17.697 182.39 36.527 0c-4.663-138.633-11.336-208.772-18.067-209.183m-96.973 69.49c-3.985.808-7.915 44.779-16.678 133.558c18.855 175.808 16.734 180.374 34.539 0c-9.577-89.276-13.749-134.392-17.861-133.558m-48.084 3.64c-3.985.792-7.915 43.897-16.678 130.913c18.855 172.313 16.734 176.79 34.539 0c-9.427-86.158-13.621-130.354-17.672-130.913c-.066 0-.126-.023-.189 0m95.473 1.988c-4.215.809-8.367 44.8-17.635 133.577c19.941 175.807 17.697 180.372 36.527 0c-9.97-87.903-14.401-133.007-18.686-133.577c-.065 0-.139-.023-.206 0M57.821 610.96c-3.985.666-7.934 36.964-16.697 110.2c18.854 145.034 16.734 148.801 34.539 0c-9.427-72.516-13.603-109.729-17.654-110.2zM13.64 650.771c-3.258.428-6.474 23.76-13.639 70.877c15.415 93.307 13.679 95.729 28.235 0C20.408 674.266 17 650.327 13.64 650.771"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Speaker(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 350.411h282.787L708.197 3.688v1192.623L282.787 849.59H0zm868.033-56.558c83.605 83.607 126.229 185.246 127.868 304.918c0 114.755-42.623 213.114-127.868 295.082l-86.066-88.523c59.018-59.018 88.525-128.688 88.525-209.018c0-81.967-29.509-153.277-88.525-213.934zm147.541-145.082C1138.524 271.722 1200 420.083 1200 593.853c0 173.771-61.476 322.951-184.426 447.541l-90.984-90.982c98.361-96.722 147.541-215.164 147.541-355.327c0-140.164-49.18-259.427-147.541-357.788z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spotify(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.624 0 600c0 144.559 51.154 277.127 136.304 380.713c235.77-249.958 704.305-173.583 820.972 101.294C1104.53 972.67 1200 797.501 1200 600C1200 268.624 931.371 0 600 0m360.645 312.378c55.036 18.106 33.668 111.442-7.837 108.911c-228.644-72.82-404.425-110.314-669.213-69.58c-55.036-18.106-33.668-111.442 7.837-108.911c236.047-20.569 441.27-14.696 669.213 69.58m-78.296 176.587c46.594 15.375 28.573 94.438-6.665 92.285c-193.738-61.68-342.678-93.438-567.041-58.96c-46.644-15.344-28.5-94.429 6.665-92.285c199.76-17.667 374.19-12.286 567.041 58.96m-64.014 174.536c39.224 12.929 24.072 79.521-5.565 77.71c-163.192-52.008-288.674-78.739-477.687-49.658c-39.252-12.952-24.051-79.523 5.64-77.71c168.749-14.709 314.52-10.667 477.612 49.658m-285.939 338.892c-92.553 0-175.839 49.683-214.893 127.075C401.69 1174.479 497.864 1200 600 1200c61.7 0 121.226-9.289 177.246-26.587c-45.149-94.27-148.425-171.02-244.85-171.02"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stackoverflow(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M284.146 943.869v97.896h472.857v-97.896zm8.003-176.091l-6.773 97.896l471.627 32.018l6.156-97.281zm-145.304-40.635V1200h747.459V727.143h-78.192v394.049H225.038V727.143zm181.014-167.471l-23.396 94.817L763.16 768.393l23.396-95.433zM429.45 322.627l-49.872 84.352l407.594 239.508l49.872-84.967zM678.192 97.896l-81.271 54.797l262.903 392.201l81.271-54.182zM995.896 0l-97.281 11.698l57.261 469.164l97.28-11.698z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Star(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M961.359 1173.121L594.085 903.799l-374.367 259.374L362.365 730.65L0 454.756l455.436 2.008L605.848 26.879l138.827 433.765L1200 470.853L830.365 736.927z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0m3.955 209.912l94.556 295.239l309.889 6.958l-251.588 181.055l89.136 296.924l-249.976-183.325l-254.81 176.587l97.119-294.434l-246.68-187.793l310.034 1.392z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarEmpty(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m605.847 26.888l-37.454 107.02l-112.955 322.836l-342.008-1.483L0 454.727l90.281 68.734l272.087 207.211l-107.137 324.737l-35.495 107.73l93.249-64.58l281.11-194.747l275.827 202.228l91.468 67.072l-32.646-108.622l-98.412-327.586l277.607-199.792l92.061-66.3l-113.37-2.493l-341.949-7.717l-104.23-325.686zm-3.205 239.145l77.936 243.538l8.25 25.818l27.065.595l255.646 5.697l-207.509 149.4l-22.021 15.788l7.835 25.938l73.542 244.903l-206.202-151.18l-21.844-16.026l-22.259 15.433l-210.179 145.602l80.13-242.825l8.429-25.701l-21.546-16.44l-203.354-154.86l255.646 1.128l27.066.119l8.962-25.583z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StepBackward(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M250 1200V0h200v550L950 50v1100L450 650v550z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StepForward(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M950 0v1200H750V650l-500 500V50l500 500V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stop(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h1200v1200H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0M300 300h600v600H300z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stumbleupon(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1075.341 231.76c131.618 176.24 160.644 407.959 79.297 601.287c-101.286 224.85-320.8 365.262-553.788 366.953c-229.359-2.552-432.951-132.474-535.091-327.039h261.744c113.074-.812 210-70.317 211.458-175.105c-.053-57.375-21.431-110.063-67.692-138.412c-40.654-23.182-83.378-27.768-126.359-35.408c-26.686-5.588-67.144-10.979-67.692-41.201c2.476-48.569 53.951-52.613 87.677-52.79h181.803c26.044.684 51.756 2.17 74.783 11.588c23.791 11.058 35.874 27.712 36.104 51.502V674.68c1.797 60.14 21.35 111.326 63.824 148.713c43.147 34.728 94.059 49.272 145.056 49.57c54.485-1.484 106.175-16.448 145.055-49.57c44.651-41.657 63.449-93.776 63.824-148.713V231.76zM600.85 0c124.424 1.608 241.951 41.194 337.817 104.292v575.536c-3.203 42.193-33.751 71.574-72.205 72.104c-42.135-3.179-72.947-33.539-73.495-72.104V507.296c-1.57-76.377-22.484-156.604-96.703-184.764c-19.771-7.296-40.833-10.944-63.18-10.944H337.816c-52.301 1.314-100.536 16.079-137.964 47.64c-38.972 35.861-57.663 80.438-58.022 128.756c.58 33.016 8.106 65.892 25.143 91.417c28.273 38.543 68.882 56.633 110.242 64.378c33.069 7.842 67.005 7.889 97.993 18.024c22.555 8.431 29.497 21.535 29.655 42.489c-1.312 24.92-14.557 35.828-35.458 41.846c-29.71 6.937-60.42 5.781-88.322 5.794H19.34C7.024 700.704.12 648.59 0 600C6.339 264.631 280.42 2.327 600.85 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tag(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M30.702 440.542V92.705C33.358 39.559 76.183.64 123.407 0h347.836c55.938 3.476 122.726 25.407 158.596 65.89l514.859 573.089c32.341 40.942 33.256 97.599 0 131.778l-403 403c-39.514 35.748-98.47 34.227-131.779 0L96.591 599.137C59.17 554.944 31.218 495.859 30.702 440.542m150.933-200.734c2.193 51.71 42.984 89.029 88.875 89.641c51.859-2.316 89.028-43.651 89.641-89.641c-2.283-51.883-43.829-88.273-89.641-88.875c-51.735 2.161-88.276 43.163-88.875 88.875"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tags(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 467.478V176.805c3.479-45.64 36.064-76.909 77.47-77.47h290.673c38.559.646 74.477 14.596 105.962 33.613l455.536 500.354c27.468 35.271 28.876 79.164 0 108.844L592.87 1077.637c-34.029 29.361-81.72 32.02-108.842 0L55.062 600.009C24.908 562.69.447 513.568 0 467.478m146.938-115.245c32.17 29.66 78.342 26.916 105.961 0c29.526-31.898 27.06-78.551 0-105.961c-31.94-29.075-78.454-26.768-105.961 0c-29.199 34.102-28.102 76.706 0 105.961m359.5-251.618h111.403c46.704 2.88 101.974 21.285 131.893 55.062L1178.7 634.582c28.985 33.471 27.808 81.543 0 110.123l-335.491 335.492c-46.495 26.355-89.107 24.518-117.806-8.965l329.088-329.089c29.633-32.787 26.469-80.075 0-108.843L670.618 203.787c-45.439-56.713-143.264-100.612-164.18-103.172"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tasks(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 131.213v234.375h1200V131.213zm752.856 58.009h385.62v118.359h-385.62zM0 482.849v234.375h1200V482.85zm487.72 58.008h650.757v118.358H487.72zM0 834.412v234.375h1200V834.412zm894.946 58.008h243.529v118.359H894.946z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextHeight(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M991.699 0L783.398 209.399H922.85V990.6H783.398L991.699 1200L1200 990.601h-139.38V209.399H1200zM5.933 207.642L0 420.703h74.048c0-146.136.1-145.972 186.841-145.972V862.5c0 73.264-.037 73.141-102.539 71.338v58.521h379.468v-58.521c-102.335 0-102.539-.179-102.539-71.338V274.731c186.839 0 187.354.041 187.354 145.972h73.535l-5.347-213.062z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextWidth(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m257.813 0l-5.859 213.062h74.048c0-146.136.027-145.972 186.768-145.972v587.77c0 73.264-.036 73.142-102.538 71.338v58.521h379.54v-58.521c-102.335 0-102.538-.18-102.538-71.338V67.09c186.839 0 187.279.041 187.279 145.972h73.535L942.7 0zm-48.414 783.398L0 991.699L209.399 1200v-139.38H990.6V1200L1200 991.699L990.601 783.398V922.85H209.399z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Th(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v300h300V0zm450 0v300h300V0zm450 0v300h300V0zM0 450v300h300V450zm450 0v300h300V450zm450 0v300h300V450zM0 900v300h300V900zm450 0v300h300V900zm450 0v300h300V900z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThLarge(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v525h525V0zm675 0v525h525V0zM0 675v525h525V675zm675 0v525h525V675z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThList(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v300h300V0zm468.75 0v300H1200V0zM0 450v300h300V450zm468.75 0v300H1200V450zM0 900v300h300V900zm468.75 0v300H1200V900z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsDown(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1141.254 797.17H842.888c6.042 48.99 14.657 96.04 25.361 141.727c5.814 86.875-52.362 149.479-113.38 190.953c-15.913 9.946-32.82 12.434-50.723 7.459c-16.914-5.968-29.348-17.899-37.298-35.804L531.094 797.17H393.846c-36.075-2.562-59.249-27.021-59.675-59.673V134.795c2.563-36.075 27.023-59.249 59.675-59.673h551.979c30.137 2.299 59.875 13.033 84.289 26.853c42.318 25.706 72.122 61.157 85.779 104.429l83.543 519.159c4.976 39.367-23.686 71.109-58.182 71.607m-923.445 11.934H43.264C19.021 807.349.333 789.809 0 767.333V101.975c1.926-25.64 20.297-41.478 43.264-41.771h174.545c25.302 1.761 42.946 18.557 43.263 41.771v665.358c-2.085 24.651-21.115 41.471-43.263 41.771"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsUp(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M717.534 60.514c42.51-.457 82.729 43.455 110.594 74.778c65.435 80.402 29.326 169.5 14.94 267.032h298.455c40.347.301 61.029 30.64 58.225 71.631l-83.567 517.806c-18.937 66.216-92.545 124.507-170.138 132.786H393.96c-32.659-.426-57.13-23.607-59.691-59.691v-602.84c.424-32.654 23.607-57.13 59.691-59.689h137.251L666.999 97.939c7.552-16.428 27.347-37.054 50.535-37.425M217.89 390.388c22.971.293 41.284 16.175 43.21 41.819v665.534c-.333 22.479-18.962 39.989-43.21 41.748H43.286c-22.155-.304-41.2-17.095-43.286-41.748V432.208c.316-23.22 17.976-40.06 43.286-41.819z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Time(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C446.454 0 292.879 58.549 175.728 175.705c-234.304 234.309-234.304 614.254 0 848.562c234.303 234.311 614.241 234.311 848.543 0c234.305-234.309 234.305-614.254 0-848.562C907.121 58.549 753.546 0 600 0m0 156.734c113.438 0 226.909 43.228 313.462 129.782c173.105 173.109 173.105 453.828 0 626.938c-173.104 173.108-453.817 173.108-626.924 0c-173.106-173.109-173.105-453.829 0-626.938C373.091 199.962 486.562 156.734 600 156.734M538.626 257.44v280.291h-155.34V691.17h308.775V257.44z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimeAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1024.236 175.764C915.665 67.192 765.675 0 600 0C268.651 0 0 268.651 0 600c0 331.35 268.651 600 600 600c331.35 0 600-268.651 600-600c0-165.675-67.192-315.665-175.764-424.236m-121.75 86.328l108.13 108.131l-405.851 405.851l-.052-.053l-.052.053l-284.153-284.153l11.859-11.858l84.05-84.05l12.222-12.221l176.074 176.074z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tint(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M175.901 775.15c1.34-84.942 24.997-161.363 67.555-227.438c43.852-65.677 92.564-129.583 133.235-190.657c64.562-99.564 119.289-202.698 152.751-308.503C541.073 15.363 567.479-3.406 600 .513c36.243-2.225 61.341 17.08 70.558 48.04c35.234 114.567 91.42 214.71 153.502 308.503c42.395 69.12 90.82 122.468 132.483 190.657c45.203 71.836 67.148 149.555 67.556 227.438c-2.581 116.192-48.298 224.208-123.477 300.997C816.77 1156.717 706.456 1199.234 600 1200c-116.136-2.474-223.853-48.567-300.247-123.852c-80.401-84.12-123.086-194.44-123.852-300.998m186.155 93.829c2.782 61.17 51.377 105.853 105.836 106.588c61.717-2.731 105.123-52.149 105.837-106.588c0-21.519-5.505-40.283-16.514-56.297c-12.101-15.747-24.03-32.046-33.776-47.664c-15.442-25.218-30.541-50.295-38.282-76.938c-2.001-10.009-7.756-14.262-17.265-12.761c-9.008-2.002-15.013 2.252-18.015 12.761c-8.094 28.541-22.238 53.7-37.156 76.938c-10.493 17.509-23.069 31.136-34.151 47.664c-11.51 16.513-17.015 35.278-16.514 56.297"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Torso(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1129.688 771.053V1200H70.312V771.053L475 598.685c-99.677-63.991-146.501-176.217-147.368-281.578C328.944 158.681 437.638 1.372 600 0c170.532 5.592 271.439 165.247 272.368 317.105c-3.385 113.922-53.606 222.121-143.421 280.264z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trash(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M451.376 0v91.743H104.587V260.55h990.826V91.743H748.624V0zM157.798 388.991V1200h884.404V388.991z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrashAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0m-93.091 224.341h186.182v57.422h217.09v105.688H289.819V281.763h217.09zm-183.764 243.53h553.71v507.788h-553.71z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tumblr(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1200 0v1200H0V0zM698.943 874.63c-20.393-5.684-40.498-12.145-56.448-24.735c-6.343-5.073-9.514-9.726-9.514-13.953V528.964H827.06v-134.46H632.981v-202.96H521.354c-8.321 104.2-63.527 202.96-172.517 202.96v134.46h92.601v332.347c.382 52.267 24.514 97.235 66.598 121.144c39.829 20.617 81.797 25.865 123.679 26.004c41.049-.2 84.731.617 123.679-1.903c34.102-1.224 66.789-9.025 95.771-23.467V838.477v1.269c-16.913 10.994-35.518 20.296-55.813 27.907c-34.048 11.164-64.44 14.827-96.409 6.977"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Twitter(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1200 216.995c-8.026 16.946-22.966 38.351-44.816 64.214s-48.828 48.606-80.937 68.227c.893 7.136 1.562 13.824 2.007 20.067c3.463 96.12-19.561 195.847-51.505 278.93c-61.918 154.473-156.795 278.795-292.308 363.21c-141.314 81.058-305.892 98.298-457.525 83.612C174.25 1083.694 75.033 1050.822 0 988.901c137.233 16.404 265.305-30.28 366.555-105.687c-112.472 2.847-194.191-78.514-232.775-172.575c17.393 4.559 35.588 4.009 52.174 2.677c20.416-1.918 39.796-3.833 58.863-8.026c-71.717-23.139-134.469-65.675-167.224-129.767c-19.213-40.61-27.888-79.282-28.093-123.077c33.605 17.693 74.028 34.937 111.037 33.444c-55.892-47.347-100.7-107.857-109.03-177.258c-5.259-57.866 9.221-110.688 30.101-159.866c83.607 91.885 177.415 167.502 285.618 214.047c73.998 30.302 147.786 46.507 224.081 46.822c-8.807-68.145-1.959-134.05 30.101-189.967c37.73-60.118 92.778-94.998 154.516-112.375c88.337-22.652 174.496 9.345 228.763 70.903c58.521-6.321 114.58-32.283 159.197-58.863c-18.958 57.41-55.874 114.492-109.699 141.806c51.247-9.488 100.149-24.659 145.815-44.144"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UniversalAccess(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0m0 80.167c287.105 0 519.833 232.727 519.833 519.833c0 287.105-232.727 519.871-519.833 519.871S80.129 887.106 80.129 600S312.894 80.167 600 80.167m322.375 460.996c17.269 16.33 16.732 39.747 3.838 56.29c-15.771 16.638-40.629 17.347-56.29 3.838c-48.914-44.386-102.442-85.918-152.237-131.77c-.152 1.96-13.502-8.813-19.189-8.956c-5.118 0-7.677 5.971-7.677 17.91c2.512 49.275-1.148 101.299 4.478 149.04c2.133 17.484 3.198 27.506 3.198 30.064l57.569 327.505c3.299 30.075-13.141 52.539-39.658 57.569c-26.557 7.224-53.846-15.676-57.569-39.658c0 0-46.445-264.445-47.335-268.657s-2.46-27.537-11.514-28.784c-11.634 4.222-10.286 23.812-11.516 28.784c-1.229 4.972-47.335 268.657-47.335 268.657c-7.78 27.743-31.696 44.14-57.568 39.658c-29.562-7.392-44.012-31.018-39.658-57.569L501.481 656.3c8.5-60.769 6.396-115.129 6.396-176.546c0-11.944-2.345-18.128-7.036-18.55c-4.69-.427-10.874 2.771-18.55 9.595L330.051 601.29c-17.788 13.255-42.655 11.046-56.29-3.838c-13.784-19.696-12.902-41.007 3.838-56.29l199.573-176.546c9.384-5.972 18.126-10.021 26.228-12.154c8.104-2.132 18.125-3.198 30.062-3.198h133.05c11.939 0 21.962 1.066 30.063 3.198c8.103 2.133 16.845 6.61 26.227 13.433c63.542 55.737 131.004 115.703 199.573 175.268M688.031 238.326c0 48.328-39.178 87.505-87.504 87.505c-48.328 0-87.506-39.177-87.506-87.505c0-48.327 39.178-87.504 87.506-87.504c48.327-.001 87.504 39.177 87.504 87.504"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unlock(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1130.771 530.771V1200H69.229V530.771l183.334-.002V362.821c2.554-99.036 41.962-191 106.411-256.411c32.479-32.478 70.728-58.333 114.743-77.564C517.736 9.615 564.958 0 615.385 0c98.104 1.998 188.043 41.887 253.205 104.487c68.789 69.893 106.146 162.255 108.333 251.922H764.104c-1.303-41.768-16.416-71.151-43.591-98.717c-30.442-28.599-66.654-43.313-105.128-43.59c-41.768 1.302-77.562 16.416-105.129 43.59c-28.599 30.443-43.312 66.655-43.589 105.128v167.95z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnlockAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0m8.396 212.462c112.097-4.539 194.092 93.25 198.101 195.351H689.894c-4.339-45.75-36.726-79.855-81.497-78.062c-44.771 1.793-80.5 37.062-81.459 81.497v91.983l363.831-.002v366.644H309.194V503.229l100.457.002v-91.983c9.478-113.423 86.65-194.247 198.745-198.786"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Upload(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 1200C268.63 1200 0 931.369 0 600C0 268.63 268.63 0 600 0c331.369 0 600 268.63 600 600s-268.631 600-600 600m0-1069.565c-259.369 0-469.565 210.261-469.565 469.565c0 259.305 210.196 469.564 469.565 469.564S1069.564 859.305 1069.564 600c0-259.304-210.195-469.565-469.564-469.565m117.392 720.652H482.608V584.348H335.87L600 335.87l264.131 248.478H717.392z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Usd(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M960.08 786.845c0 76.036-27.49 137.244-82.485 183.623c-54.995 45.906-109.426 73.398-208.41 82.487v147.043H525.82V1055.83c-116.676-2.39-196.456-22.959-284.445-61.681V804.786c41.598 20.556 91.575 38.734 149.907 54.51c58.817 15.775 88.62 25.1 134.527 27.976V664.914l-25.507-18.65c-94.676-37.303-161.86-77.706-201.553-121.221c-39.208-43.99-58.817-98.262-58.817-162.819c0-69.338 27.018-126.238 81.054-170.712c54.51-44.947 107.756-72.202 204.823-81.764V.007h143.367v106.876c109.504 4.782 185.935 26.782 274.397 65.99l-67.42 167.836c-74.593-30.604-128.55-49.255-206.977-55.942v211.6c93.243 35.859 136.917 66.948 176.136 93.244c39.692 26.296 68.616 55.231 86.793 86.794c18.64 31.55 27.977 68.368 27.977 110.44m-215.908 10.769c0-20.083-8.13-37.054-24.39-50.924c-16.26-13.87-18.099-28.213-50.608-43.042v178.606c65.03-11.006 74.998-39.22 74.998-84.64m-288.335-435.39c0 21.042 7.172 38.498 21.515 52.356c14.829 13.385 15.945 27.255 48.455 41.598V287.631c-61.693 9.09-69.97 33.953-69.97 74.593"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func User(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M939.574 858.383c-157.341-57.318-207.64-105.702-207.64-209.298c0-62.17 51.555-102.462 69.128-155.744c17.575-53.283 27.741-116.367 36.191-162.256c8.451-45.889 11.809-63.638 16.404-112.532C859.276 157.532 818.426 0 600 0C381.639 0 340.659 157.532 346.404 218.553c4.596 48.894 7.972 66.645 16.404 112.532c8.433 45.888 18.5 108.969 36.063 162.256c17.562 53.286 69.19 93.574 69.19 155.744c0 103.596-50.298 151.979-207.638 209.298C102.511 915.83 0 972.479 0 1012.5V1200h1200v-187.5c0-39.957-102.574-96.606-260.426-154.117"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Viadeo(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m551.104 0l161.912 346.347C715.318 258.753 642.125 52.008 551.104 0m506.718 24.356c-2.346.21-4.85 2.874-7.557 8.217c-36.746 67.145-100.152 107.015-159.418 123.69c-60.256 16.911-122.928 57.26-145.92 103.001c-26.859 62.373-15.342 111.877 10.785 168.075c32.949-4.992 65.922-17.72 93.832-31.105c71.59-34.855 133.605-89.476 173.943-153.402c16.221-25.709 19.395-21.096 5.428 7.85c-26.299 54.5-112.469 130.87-220.676 195.659c-24.117 14.44-24.473 14.938-13.646 24.356c34.631 30.131 72.586 48.366 107.037 51.28c149.416.688 213.553-159.392 217.742-270.637c-1.113-72.184-13.322-165.462-57.518-224.711c-1.34-1.685-2.625-2.399-4.032-2.273M513.175 334.902c-238.308 0-432.548 194.24-432.548 432.549C80.626 1005.76 274.868 1200 513.175 1200c238.309 0 432.477-194.24 432.477-432.549c0-59.577-12.064-116.655-33.969-168.441l-90.75 38.369c16.891 39.929 26.266 83.804 26.266 130.072c0 179.681-140.414 325.314-317.957 333.655C734.705 960.734 827.717 597.2 717.859 366.082c8.727 158.612 1.566 590.192-244.007 733.118c-166.393-19.355-294.772-159.99-294.772-331.748c0-185.076 149.02-334.022 334.096-334.022c46.27 0 89.997 9.229 129.926 26.117l38.441-90.603c-51.785-21.905-108.791-34.042-168.368-34.042"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Video(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 145.898v908.203h1200V145.898zm147.144 147.218h905.713v613.77H147.144zm318.237 106.861v408.839L818.848 603.81z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0M242.285 329.297h715.356v541.406H242.285zm87.744 87.744v365.918H869.97V417.041zm189.698 63.721l210.719 121.51l-210.719 122.24z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoChat(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 49.109v854.518h247.266v247.264L632.74 903.625H1200V49.109zm263.086 214.673h483.398v149.049l190.43-149.049v398.439l-190.43-149.049v149.047H263.086z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewMode(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1110.645 0L0 1111.963L89.355 1200L1200 89.429zM0 1.758v281.323h281.25V1.831zm331.421 0v281.323h281.25V1.831zM0 338.452v281.25h281.25v-281.25zm331.421 0v281.25l281.25-281.25zm494.311 297.876L652.295 794.312H1200V636.328zM614.795 838.33v157.982H1200V838.33zm0 202.002v157.91H1200v-157.91z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vimeo(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M952.643 318.393C828.707 215.97 664.09 355.939 625.37 468.076c23.875.411 51.192-2.071 71.67 5.708c31.619 17.104 32.06 50.073 29.81 78.013c-14.099 61.826-57.908 174.325-114.165 183.298c-11.839 1.27-24.101-5.285-36.785-19.661c-49.165-61.587-43.074-142.604-54.546-210.57c-10.732-56.861-18.824-171.623-86.259-196.617c-74.271-9.071-139.494 46.271-187.103 90.062c-29.176 27.062-58.985 52.854-89.429 77.379v5.073c7.243 8.71 13.113 17.589 19.662 26.005c11.113 11.184 31.389 8.446 43.763 4.439c44.995-11.698 60.81-23.164 84.989 19.662c36.439 98.508 64.028 209.531 88.161 303.805c16.867 53.772 39.112 122.449 93.234 144.609c33.438 14.208 91.119-4.886 116.702-20.296c109.478-69.163 194.195-179.831 254.968-286.682c44.368-90.029 154.677-269.147 92.601-353.91M1200 0v1200H0V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vkontakte(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v1200h1200V0zm532.69 320.874c33.977-.352 67.21 3.415 83.423 3.882c60.1 1.73 70.714 32.127 69.945 57.129c-1.124 36.617-24.683 191.51 18.971 206.616c63.609-34.014 106.265-139.659 137.328-206.616c15.714-33.868 13.229-35.551 49.438-35.596l143.848-.22c34.641-.043 33.289-1.258 44.386 12.891c25.81 32.908-81.447 169.794-116.09 207.642c-63.929 69.847-26.63 96.895-16.332 106.494c47.709 44.483 93.809 90.611 122.461 133.154c22.778 33.822 13.37 65.312-15.602 66.504l-154.248 6.372c-51.186 2.104-130.776-89.652-143.48-109.644c-10.274-16.49-59.789-37.157-64.453 4.979l-6.52 58.813c-4.734 42.041-35.775 41.06-60.277 42.113c-169.669 7.301-227.45-57.297-312.892-152.637c-62.621-88.56-122.38-177.629-171.899-284.326c-14.822-31.938-22.673-55.367-24.683-80.2c24.916-19.474 112.373-11.086 171.094-12.378c16.061-.354 26.898 8.849 35.376 27.173c33.081 71.513 69.159 141.739 116.895 203.833c20.789 19.148 39.167 26.189 48.34-12.743c5.326-.49 4.898-64.964 5.2-113.6c.207-33.394-3.716-73.441-60.278-91.919c11.441-29.989 56.365-37.263 100.049-37.716"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeDown(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M170.671 384.805h243.816l366.784-298.94v1028.27l-366.784-298.94H170.671zm748.41-48.764c72.085 72.084 108.835 159.717 110.248 262.898c0 98.938-36.749 183.744-110.248 254.417l-74.205-76.325c50.884-50.884 76.325-110.954 76.325-180.212c0-70.672-25.441-132.156-76.325-184.453z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeOff(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1089.041 681.359c1.719 1.718 2.578 3.741 2.578 6.073c0 2.331-.859 4.356-2.578 6.074l-57.062 57.063c-1.719 1.718-3.682 2.577-5.891 2.577c-1.963 0-4.049-.859-6.258-2.577l-81.359-81.36l-81.359 81.36c-1.719 1.718-3.684 2.577-5.893 2.577c-1.963 0-4.049-.859-6.258-2.577l-57.062-57.063c-1.475-1.473-2.209-3.559-2.209-6.259c0-2.454.734-4.417 2.209-5.89l81.359-81.36l-80.992-81.36c-1.719-1.473-2.576-3.436-2.576-5.89c0-2.7.857-4.786 2.576-6.259l56.693-57.063c1.719-1.718 3.805-2.577 6.26-2.577c2.453 0 4.418.859 5.891 2.577l81.359 81.36l81.359-81.36c1.719-1.718 3.744-2.577 6.074-2.577c2.332 0 4.355.859 6.074 2.577l57.062 57.063c1.719 1.718 2.578 3.743 2.578 6.074s-.859 4.355-2.578 6.073l-81.357 81.362zm-980.66-296.554h243.817l366.784-298.94v1028.27l-366.783-298.94H108.381z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeUp(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 384.805h243.816L610.6 85.866v1028.269l-366.784-298.94H0zm748.41-48.763c72.085 72.084 108.834 159.717 110.247 262.897c0 98.94-36.749 183.745-110.247 254.418l-74.205-76.326c50.883-50.883 76.325-110.953 76.325-180.211c0-70.672-25.442-132.156-76.325-184.453zm127.208-125.088c106.008 106.007 159.011 233.922 159.011 383.745c0 149.825-53.003 278.444-159.011 385.866l-78.445-78.445c84.806-83.393 127.209-185.512 127.209-306.36c0-120.847-42.403-223.675-127.208-308.48zM992.227 94.346c65.019 65.018 115.9 140.636 152.649 226.855C1181.625 407.42 1200 498.586 1200 594.699c0 96.112-18.375 187.633-55.124 274.559s-87.632 162.896-152.649 227.916l-76.324-76.326c117.313-117.314 175.972-259.01 175.972-425.088c0-166.078-58.657-307.774-175.972-425.088z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WthreeC(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1073.828 629.485c2.898-11.642-.086-23.613-10.971-27.43l-45.259-20.57c-11.184-4.669-25.42 1.638-27.428 12.343c-8.735 18.702-21.121 34.011-41.144 34.286c-17.372 0-31.543-9.371-42.515-28.114s-16.458-43.2-16.458-73.371c0-30.172 5.486-54.856 16.458-74.058c10.972-19.2 25.143-28.8 42.515-28.8c21.915 1.592 32.312 15.432 41.144 32.914c4.681 11.483 17.413 14.182 27.428 10.972l45.259-20.571c10.479-5.517 15.148-19.174 9.6-28.8c-26.514-56.687-67.656-85.029-123.429-85.029c-46.628 0-84.344 17.829-113.144 53.485c-28.8 35.658-43.2 82.286-43.2 139.886c0 57.601 14.4 103.999 43.2 139.199c28.8 35.201 66.516 52.801 113.144 52.801c55.773.001 97.372-29.714 124.8-89.143m-436.114 89.144c72.696-1.124 128.054-51.075 128.915-116.572c-1.284-34.587-15.191-61.488-39.771-82.285c21.644-20.438 34.07-48.272 34.284-75.429c-1.799-74.11-63.438-109.123-123.429-109.714c-52.114 0-92.342 20.572-120.686 61.714c-6.4 9.146-5.485 18.288 2.742 27.429l32.914 30.172c10.552 9.405 22.288 7.667 31.543-1.371c13.716-19.2 30.173-28.801 49.371-28.801c15.543 0 24.688 3.886 27.43 11.657c5.288 11.53 3.768 22.05 0 31.543c-14.146 14.646-34.817 10.285-54.856 10.285c-13.034.438-23.115 8.081-23.314 20.571V547.2c1.138 14.162 10.791 23.521 23.314 20.571c15.117-.016 53.453-1.728 59.656 10.972c3.2 7.312 4.801 13.257 4.801 17.828c0 21.027-11.429 31.543-34.286 31.543c-21.942 0-40.228-10.058-54.856-30.172c-8.948-10.827-23.086-10.985-31.542-2.742l-34.286 34.285c-8.229 8.229-9.145 16.914-2.743 26.058c31.079 43.722 76.328 62.701 124.799 63.086M501.943 364.8c2.506-13.146-7.177-25.827-19.2-26.057H422.4c-10.973 0-17.372 5.485-19.2 16.458l-28.8 146.743L344.229 355.2c-2.743-10.972-9.6-16.458-20.571-16.458h-43.886c-11.888 0-18.745 5.485-20.571 16.458l-30.172 146.743l-28.8-146.743c-1.828-10.972-8.685-16.458-20.571-16.458h-58.972c-15.031 1.183-22.255 13.067-20.571 26.057l75.429 331.886c2.743 10.973 9.6 16.457 20.571 16.457h57.6c10.971 0 17.829-5.484 20.571-16.457l27.429-137.143l27.428 137.143c1.828 10.973 8.686 16.457 20.571 16.457h57.602c10.972 0 17.828-5.484 20.57-16.457zM1200 0v918.75L600 1200L0 918.75V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WarningSign(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0L0 1200h1200zm-46.143 416.089h92.284v158.644l-22.559 223.096h-47.168l-22.56-223.096V416.089zm0 469.336h92.284v99.391h-92.284z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Website(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M698.598 922.244h-199.6V722.646h199.6zm252.506 0H751.503V722.646h199.601zm-505.012 0H246.493V722.646h199.599zm505.012-631.262v367.936h-704.61V290.982zm101-105.812H147.896v829.66h904.209zM0 1161.521V38.478h1200v1123.045z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WebsiteAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0M281.154 301.634h637.653v596.732H281.154zm78.611 77.963V820.44h480.471V379.597zm52.393 56.231H786.55v195.479H412.158zm0 229.369h106.041v106.041H412.158zm134.156 0h106.078v106.041H546.314zm134.157 0H786.55v106.041H680.471z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wheelchair(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M476.721 0c-65.602.471-117.577 50.855-120.375 119.094c.477 66.402 51.395 117.868 120.375 120.406c61.426-.832 115.989-51.475 119.094-120.406C595.349 52.962 544.73 3.015 476.721 0m33.281 266.375c-27.32 0-50.372 9.175-69.156 27.531c-15.261 15.317-32.347 49.101-26.875 76.219h-1.281L483.127 709.5c7.265 38.873 39.258 63.938 81.969 66.594c98.116-.27 204.507-2.497 304.781-1.281l169.062 294.563c19.023 40.147 67.734 48.603 103.095 19.219c16.965-16.834 23.826-45.475 12.155-70.438L958.252 677.469c-14.117-29.208-41.97-45.662-76.844-47.375H663.689l-25.594-125.5h170.313c36.818 2.724 63.343-25.375 55.719-64.031c-5.602-18.888-25.174-36.278-48.031-37.155H616.314c-10.068-51.463-8.466-82.573-48.03-117.813c-16.649-12.808-36.084-19.22-58.282-19.22m-183.719 145.25a72.873 72.873 0 0 0-16.063 2.031c-75.444 29.355-143.317 79.801-195.938 149.844c-46.369 65.387-72.293 144.095-74.28 230.531c.772 102.442 44.761 207.567 123.594 286.875c76.408 72.772 182.465 116.699 296.469 119.094c83.152-.458 166.274-25.016 242.063-73.625c68.07-46.237 121.414-112.588 153.656-195.312c5.977-14.516 11.105-30.759 15.375-48.688c6.25-28.186-9.57-51.084-39.688-60.188c-27.394-6.069-51.893 9.046-61.469 38.438a187.143 187.143 0 0 1-11.531 37.125c-21.749 57.507-62.231 107.283-115.906 146c-52.185 35.487-114.684 54.813-182.5 56.344c-80.649-.573-158.332-31.315-223.469-89.625c-55.969-54.264-90.982-131.912-92.844-216.438c.358-60.255 19.762-119.715 56.344-173.531c35.875-49.807 86.096-88.272 147.281-113.344c25.783-8.456 39.673-35.15 30.719-64.031c-8.649-20.761-29.024-31.808-51.813-31.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wordpress(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M599.314 1200C264.241 1193.561 2.343 918.889 0 599.314C6.515 264.491 279.729 2.343 599.314 0C934.157 6.694 1197.654 279.298 1200 599.314c-6.619 335.092-280.682 598.34-600.686 600.686m0-1157.486C288.196 48.64 43.307 302.232 41.143 599.314c6.051 311.365 261.102 556.007 558.171 558.171c311.365-6.053 556.007-261.103 558.172-558.171c-5.873-311.347-261.532-554.638-558.172-556.8m-144 1047.771l153.6-441.6l160.457 432c-108.102 36.189-210.303 38.382-314.057 9.6M337.371 309.943c-56.23 7.644-110.916 10.369-164.571 8.229C274.677 172.521 437.796 90.579 599.314 89.143c132.109 2.5 255.27 52.438 345.6 134.4c-41.893-3.097-66.977 15.276-85.028 46.628c-26.234 80.939 29.201 140.853 60.344 201.601c29.443 55.505 27.108 118.618 10.971 172.8l-76.8 261.943l-185.143-549.944c19.291-1.795 38.708-2.373 56.229-5.486c20.736-4.617 25.286-23.914 10.971-35.657c-4.57-3.656-9.6-5.484-15.086-5.484l-111.085 8.229h-84.343c-23.813 1.406-91.053-21.597-96.688 9.601c-2.911 12.017 6.06 21.436 16.457 23.313c18.346 2.377 39.365 5.076 56.229 6.857l80.914 216.686L470.4 906.515L283.886 356.571c19.736-1.714 39.733-2.312 57.6-5.486c14.63-1.833 21.028-9.146 19.2-21.943c-2.593-11.377-12.734-19.064-23.315-19.199m-205.714 85.029l245.486 663.771c-88.577-43.659-158.606-107.864-208.458-184.458C77.49 728.754 67.897 544.121 131.657 394.972m959.315 340.799c-40.303 131.608-122.616 240.613-236.57 306.515c5.484-14.629 14.171-39.314 26.057-74.058l142.629-414.172c13.714-40.229 23.314-85.028 28.8-134.399c1.853-20.199 1.914-40.564-1.371-58.972c58.437 126.043 75.015 250.892 40.455 375.086"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wrench(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M984.091 19.305C868.695-22.216 734.617 3.236 642.153 95.7c-92.463 92.463-117.916 226.542-76.396 341.937L0 1003.396L196.604 1200l565.759-565.76c115.396 41.521 249.474 16.068 341.937-76.396c92.464-92.463 117.917-226.542 76.396-341.937L981.742 414.861l-141.708-54.896l-54.896-141.708zM236.18 963.82c22.563 22.562 22.563 59.143 0 81.705c-22.562 22.562-59.143 22.562-81.705 0c-22.563-22.563-22.563-59.144 0-81.705c22.562-22.562 59.143-22.562 81.705 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WrenchAlt(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M600 0C268.629 0 0 268.629 0 600s268.629 600 600 600s600-268.629 600-600S931.371 0 600 0m163.403 239.868c22.729-.064 45.519 3.784 67.163 11.572L711.108 370.898l32.959 85.033L829.1 488.89l119.46-119.456c24.923 69.264 9.648 149.801-45.85 205.299s-135.962 70.771-205.225 45.85L357.86 960.133L239.868 842.139l339.551-339.624c-24.922-69.263-9.647-149.726 45.85-205.225c38.155-38.155 88.127-57.281 138.134-57.422m-430.81 578.54c-13.543 13.542-13.543 35.457 0 48.999c13.542 13.542 35.457 13.542 48.999 0c13.542-13.543 13.542-35.457 0-48.999c-14.679-13.344-36.351-12.267-48.999 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Youtube(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1200 1055.438H0V144.562h1200zm-772.708-189.34l419.616-263.539l-419.616-263.54z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomIn(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M614.166 307.121H460.605v153.561H307.044v153.562h153.561v153.56h153.561v-153.56h153.561V460.682H614.166zM1200 1066.325L978.182 844.507c60.81-87.069 96.743-192.795 96.743-307.045C1074.925 240.63 834.219 0 537.463 0S0 240.63 0 537.462c0 296.833 240.629 537.463 537.463 537.463c114.249 0 220.052-35.933 307.121-96.743L1066.402 1200zM537.386 921.441c-211.991 0-383.902-171.836-383.902-383.902c0-212.067 171.911-383.902 383.902-383.902c212.067 0 383.901 171.835 383.901 383.902c.077 212.066-171.757 383.902-383.901 383.902"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomOut(children ...ElementRenderer) *ElIcon {
	return &ElIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1200 1066.325L978.182 844.507c60.81-87.068 96.743-192.795 96.743-307.044C1074.925 240.629 834.219 0 537.463 0S0 240.629 0 537.463c0 296.832 240.629 537.462 537.463 537.462c114.249 0 220.052-35.933 307.121-96.743L1066.402 1200zM537.386 921.44c-211.99 0-383.902-171.834-383.902-383.901s171.912-383.901 383.902-383.901c212.067 0 383.901 171.834 383.901 383.901c.078 212.067-171.757 383.901-383.901 383.901M307.045 614.243h460.682V460.682H307.045z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
