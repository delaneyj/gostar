package akar_icons

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = "1.9.31"
	hAttr          = "1em"
	viewbox        = "0 0 24 24"
	fill           = "currentColor"
)

type AkarIconsIcon struct {
	*SVGSVGElement
}

type AkarIconsIconFn func(children ...ElementRenderer) *AkarIconsIcon

var IconLookup = map[string]AkarIconsIconFn{
	"air":                     Air,
	"airplayAudio":            AirplayAudio,
	"airplayVideo":            AirplayVideo,
	"airpods":                 Airpods,
	"alarm":                   Alarm,
	"alignBottom":             AlignBottom,
	"alignHorizontalCenter":   AlignHorizontalCenter,
	"alignLeft":               AlignLeft,
	"alignRight":              AlignRight,
	"alignToBottom":           AlignToBottom,
	"alignToMiddle":           AlignToMiddle,
	"alignToTop":              AlignToTop,
	"alignTop":                AlignTop,
	"alignVerticalCenter":     AlignVerticalCenter,
	"androidFill":             AndroidFill,
	"angularFill":             AngularFill,
	"arrowBack":               ArrowBack,
	"arrowBackThick":          ArrowBackThick,
	"arrowBackThickFill":      ArrowBackThickFill,
	"arrowClockwise":          ArrowClockwise,
	"arrowCounterClockwise":   ArrowCounterClockwise,
	"arrowCycle":              ArrowCycle,
	"arrowDown":               ArrowDown,
	"arrowDownLeft":           ArrowDownLeft,
	"arrowDownRight":          ArrowDownRight,
	"arrowDownThick":          ArrowDownThick,
	"arrowForward":            ArrowForward,
	"arrowForwardThick":       ArrowForwardThick,
	"arrowForwardThickFill":   ArrowForwardThickFill,
	"arrowLeft":               ArrowLeft,
	"arrowLeftThick":          ArrowLeftThick,
	"arrowRepeat":             ArrowRepeat,
	"arrowRight":              ArrowRight,
	"arrowRightLeft":          ArrowRightLeft,
	"arrowRightThick":         ArrowRightThick,
	"arrowShuffle":            ArrowShuffle,
	"arrowUp":                 ArrowUp,
	"arrowUpDown":             ArrowUpDown,
	"arrowUpLeft":             ArrowUpLeft,
	"arrowUpRight":            ArrowUpRight,
	"arrowUpThick":            ArrowUpThick,
	"ascending":               Ascending,
	"attach":                  Attach,
	"augmentedReality":        AugmentedReality,
	"backspace":               Backspace,
	"backspaceFill":           BackspaceFill,
	"bank":                    Bank,
	"basket":                  Basket,
	"batteryCharging":         BatteryCharging,
	"batteryEmpty":            BatteryEmpty,
	"batteryFull":             BatteryFull,
	"batteryLow":              BatteryLow,
	"batteryMedium":           BatteryMedium,
	"behanceFill":             BehanceFill,
	"bell":                    Bell,
	"bicycle":                 Bicycle,
	"bitcoinFill":             BitcoinFill,
	"block":                   Block,
	"bluetooth":               Bluetooth,
	"boat":                    Boat,
	"book":                    Book,
	"bookClose":               BookClose,
	"bookOpen":                BookOpen,
	"bookmark":                Bookmark,
	"bootstrapFill":           BootstrapFill,
	"box":                     Box,
	"briefcase":               Briefcase,
	"bug":                     Bug,
	"cake":                    Cake,
	"calculator":              Calculator,
	"calendar":                Calendar,
	"camera":                  Camera,
	"cart":                    Cart,
	"chatAdd":                 ChatAdd,
	"chatApprove":             ChatApprove,
	"chatBubble":              ChatBubble,
	"chatDots":                ChatDots,
	"chatEdit":                ChatEdit,
	"chatError":               ChatError,
	"chatQuestion":            ChatQuestion,
	"chatRemove":              ChatRemove,
	"check":                   Check,
	"checkBox":                CheckBox,
	"checkBoxFill":            CheckBoxFill,
	"checkIn":                 CheckIn,
	"chess":                   Chess,
	"chevronDown":             ChevronDown,
	"chevronDownSmall":        ChevronDownSmall,
	"chevronHorizontal":       ChevronHorizontal,
	"chevronLeft":             ChevronLeft,
	"chevronLeftSmall":        ChevronLeftSmall,
	"chevronRight":            ChevronRight,
	"chevronRightSmall":       ChevronRightSmall,
	"chevronUp":               ChevronUp,
	"chevronUpSmall":          ChevronUpSmall,
	"chevronVertical":         ChevronVertical,
	"circle":                  Circle,
	"circleAlert":             CircleAlert,
	"circleAlertFill":         CircleAlertFill,
	"circleCheck":             CircleCheck,
	"circleCheckFill":         CircleCheckFill,
	"circleChevronDown":       CircleChevronDown,
	"circleChevronDownFill":   CircleChevronDownFill,
	"circleChevronLeft":       CircleChevronLeft,
	"circleChevronLeftFill":   CircleChevronLeftFill,
	"circleChevronRight":      CircleChevronRight,
	"circleChevronRightFill":  CircleChevronRightFill,
	"circleChevronUp":         CircleChevronUp,
	"circleChevronUpFill":     CircleChevronUpFill,
	"circleFill":              CircleFill,
	"circleMinus":             CircleMinus,
	"circleMinusFill":         CircleMinusFill,
	"circlePlus":              CirclePlus,
	"circlePlusFill":          CirclePlusFill,
	"circleTriangleDown":      CircleTriangleDown,
	"circleTriangleDownFill":  CircleTriangleDownFill,
	"circleTriangleLeft":      CircleTriangleLeft,
	"circleTriangleLeftFill":  CircleTriangleLeftFill,
	"circleTriangleRight":     CircleTriangleRight,
	"circleTriangleRightFill": CircleTriangleRightFill,
	"circleTriangleUp":        CircleTriangleUp,
	"circleTriangleUpFill":    CircleTriangleUpFill,
	"circleX":                 CircleX,
	"circleXfill":             CircleXfill,
	"clipboard":               Clipboard,
	"clock":                   Clock,
	"cloud":                   Cloud,
	"cloudDownload":           CloudDownload,
	"cloudUpload":             CloudUpload,
	"codepenFill":             CodepenFill,
	"coffee":                  Coffee,
	"coin":                    Coin,
	"command":                 Command,
	"comment":                 Comment,
	"commentAdd":              CommentAdd,
	"computing":               Computing,
	"copy":                    Copy,
	"creditCard":              CreditCard,
	"creditCardAltOne":        CreditCardAltOne,
	"cross":                   Cross,
	"crown":                   Crown,
	"cssFill":                 CssFill,
	"cursor":                  Cursor,
	"cut":                     Cut,
	"dashboard":               Dashboard,
	"data":                    Data,
	"dental":                  Dental,
	"descending":              Descending,
	"desktopDevice":           DesktopDevice,
	"devices":                 Devices,
	"diamond":                 Diamond,
	"diceFive":                DiceFive,
	"diceFour":                DiceFour,
	"diceOne":                 DiceOne,
	"diceSix":                 DiceSix,
	"diceThree":               DiceThree,
	"diceTwo":                 DiceTwo,
	"discordFill":             DiscordFill,
	"djangoFill":              DjangoFill,
	"door":                    Door,
	"dotGrid":                 DotGrid,
	"dotGridFill":             DotGridFill,
	"doubleCheck":             DoubleCheck,
	"doubleSword":             DoubleSword,
	"download":                Download,
	"draft":                   Draft,
	"dragHorizontal":          DragHorizontal,
	"dragHorizontalFill":      DragHorizontalFill,
	"dragVertical":            DragVertical,
	"dragVerticalFill":        DragVerticalFill,
	"dribbbleFill":            DribbbleFill,
	"dropboxFill":             DropboxFill,
	"edit":                    Edit,
	"enlarge":                 Enlarge,
	"envelope":                Envelope,
	"equal":                   Equal,
	"equalFill":               EqualFill,
	"eye":                     Eye,
	"eyeClosed":               EyeClosed,
	"eyeOpen":                 EyeOpen,
	"eyeSlashed":              EyeSlashed,
	"faceHappy":               FaceHappy,
	"faceNeutral":             FaceNeutral,
	"faceSad":                 FaceSad,
	"faceVeryHappy":           FaceVeryHappy,
	"faceVerySad":             FaceVerySad,
	"faceWink":                FaceWink,
	"facebookFill":            FacebookFill,
	"figmaFill":               FigmaFill,
	"file":                    File,
	"filter":                  Filter,
	"fire":                    Fire,
	"flag":                    Flag,
	"flashlight":              Flashlight,
	"folder":                  Folder,
	"folderAdd":               FolderAdd,
	"forkLeft":                ForkLeft,
	"forkRight":               ForkRight,
	"frame":                   Frame,
	"fullScreen":              FullScreen,
	"gameController":          GameController,
	"gatsbyFill":              GatsbyFill,
	"gear":                    Gear,
	"gift":                    Gift,
	"githubFill":              GithubFill,
	"githubOutlineFill":       GithubOutlineFill,
	"gitlabFill":              GitlabFill,
	"glasses":                 Glasses,
	"globe":                   Globe,
	"googleContainedFill":     GoogleContainedFill,
	"googleFill":              GoogleFill,
	"graphqlFill":             GraphqlFill,
	"grid":                    Grid,
	"hammer":                  Hammer,
	"hand":                    Hand,
	"hashtag":                 Hashtag,
	"headphone":               Headphone,
	"health":                  Health,
	"heart":                   Heart,
	"height":                  Height,
	"heptagon":                Heptagon,
	"heptagonFill":            HeptagonFill,
	"hexagon":                 Hexagon,
	"hexagonFill":             HexagonFill,
	"history":                 History,
	"home":                    Home,
	"homeAltOne":              HomeAltOne,
	"htmlFill":                HtmlFill,
	"image":                   Image,
	"inbox":                   Inbox,
	"infinite":                Infinite,
	"infinity":                Infinity,
	"info":                    Info,
	"infoFill":                InfoFill,
	"instagramFill":           InstagramFill,
	"jar":                     Jar,
	"javascriptFill":          JavascriptFill,
	"jqueryFill":              JqueryFill,
	"key":                     Key,
	"keyCap":                  KeyCap,
	"language":                Language,
	"laptopDevice":            LaptopDevice,
	"leaf":                    Leaf,
	"lifesaver":               Lifesaver,
	"lightBulb":               LightBulb,
	"linkChain":               LinkChain,
	"linkOff":                 LinkOff,
	"linkOn":                  LinkOn,
	"linkOut":                 LinkOut,
	"linkedinBoxFill":         LinkedinBoxFill,
	"linkedinFill":            LinkedinFill,
	"linkedinVoneFill":        LinkedinVoneFill,
	"linkedinVtwoFill":        LinkedinVtwoFill,
	"linkedinvOneFill":        LinkedinvOneFill,
	"linkedinvTwoFill":        LinkedinvTwoFill,
	"location":                Location,
	"lockOff":                 LockOff,
	"lockOn":                  LockOn,
	"map":                     Map,
	"mastodonFill":            MastodonFill,
	"mediumFill":              MediumFill,
	"mention":                 Mention,
	"microphone":              Microphone,
	"miniplayer":              Miniplayer,
	"minus":                   Minus,
	"mobileDevice":            MobileDevice,
	"money":                   Money,
	"moon":                    Moon,
	"moonFill":                MoonFill,
	"moreHorizontal":          MoreHorizontal,
	"moreHorizontalFill":      MoreHorizontalFill,
	"moreVertical":            MoreVertical,
	"moreVerticalFill":        MoreVerticalFill,
	"music":                   Music,
	"musicAlbum":              MusicAlbum,
	"musicAlbumFill":          MusicAlbumFill,
	"musicNote":               MusicNote,
	"network":                 Network,
	"newspaper":               Newspaper,
	"nextjsFill":              NextjsFill,
	"nodeFill":                NodeFill,
	"normalScreen":            NormalScreen,
	"npmFill":                 NpmFill,
	"octagon":                 Octagon,
	"octagonFill":             OctagonFill,
	"octocatFill":             OctocatFill,
	"openEnvelope":            OpenEnvelope,
	"oval":                    Oval,
	"panelBottom":             PanelBottom,
	"panelLeft":               PanelLeft,
	"panelRight":              PanelRight,
	"panelSplit":              PanelSplit,
	"panelSplitColumn":        PanelSplitColumn,
	"panelSplitRow":           PanelSplitRow,
	"panelTop":                PanelTop,
	"paper":                   Paper,
	"paperAirplane":           PaperAirplane,
	"parallelogram":           Parallelogram,
	"pause":                   Pause,
	"pencil":                  Pencil,
	"pentagon":                Pentagon,
	"pentagonFill":            PentagonFill,
	"peopleGroup":             PeopleGroup,
	"peopleMultiple":          PeopleMultiple,
	"percentage":              Percentage,
	"person":                  Person,
	"personAdd":               PersonAdd,
	"personCheck":             PersonCheck,
	"personCross":             PersonCross,
	"phone":                   Phone,
	"phpFill":                 PhpFill,
	"pin":                     Pin,
	"pinterestFill":           PinterestFill,
	"plane":                   Plane,
	"planeFill":               PlaneFill,
	"planet":                  Planet,
	"plant":                   Plant,
	"play":                    Play,
	"plus":                    Plus,
	"pointerDownFill":         PointerDownFill,
	"pointerHand":             PointerHand,
	"pointerLeftFill":         PointerLeftFill,
	"pointerRightFill":        PointerRightFill,
	"pointerUpFill":           PointerUpFill,
	"pointingUp":              PointingUp,
	"postgresqlFill":          PostgresqlFill,
	"priceCut":                PriceCut,
	"productHuntFill":         ProductHuntFill,
	"pythonFill":              PythonFill,
	"question":                Question,
	"questionFill":            QuestionFill,
	"radio":                   Radio,
	"radioFill":               RadioFill,
	"radish":                  Radish,
	"reactFill":               ReactFill,
	"receipt":                 Receipt,
	"reciept":                 Reciept,
	"redditFill":              RedditFill,
	"reduce":                  Reduce,
	"reduxFill":               ReduxFill,
	"reply":                   Reply,
	"ribbon":                  Ribbon,
	"rockOn":                  RockOn,
	"rss":                     Rss,
	"sassFill":                SassFill,
	"save":                    Save,
	"schedule":                Schedule,
	"scissor":                 Scissor,
	"search":                  Search,
	"send":                    Send,
	"settingsHorizontal":      SettingsHorizontal,
	"settingsVertical":        SettingsVertical,
	"shareArrow":              ShareArrow,
	"shareBox":                ShareBox,
	"shield":                  Shield,
	"shippingBoxOne":          ShippingBoxOne,
	"shippingBoxTwo":          ShippingBoxTwo,
	"shippingBoxVone":         ShippingBoxVone,
	"shippingBoxVtwo":         ShippingBoxVtwo,
	"shoppingBag":             ShoppingBag,
	"sidebarLeft":             SidebarLeft,
	"sidebarRight":            SidebarRight,
	"signOut":                 SignOut,
	"slackFill":               SlackFill,
	"slice":                   Slice,
	"snapchatFill":            SnapchatFill,
	"sort":                    Sort,
	"soundDown":               SoundDown,
	"soundOff":                SoundOff,
	"soundOn":                 SoundOn,
	"soundUp":                 SoundUp,
	"soundcloudFill":          SoundcloudFill,
	"sparkles":                Sparkles,
	"spotifyFill":             SpotifyFill,
	"square":                  Square,
	"squareFill":              SquareFill,
	"stackOverflowFill":       StackOverflowFill,
	"star":                    Star,
	"statisticDown":           StatisticDown,
	"statisticUp":             StatisticUp,
	"stop":                    Stop,
	"stopFill":                StopFill,
	"sun":                     Sun,
	"sunFill":                 SunFill,
	"sword":                   Sword,
	"tabletDevice":            TabletDevice,
	"tag":                     Tag,
	"telegramFill":            TelegramFill,
	"telescope":               Telescope,
	"tetragon":                Tetragon,
	"tetragonFill":            TetragonFill,
	"textAlignCenter":         TextAlignCenter,
	"textAlignJustified":      TextAlignJustified,
	"textAlignLeft":           TextAlignLeft,
	"textAlignRight":          TextAlignRight,
	"threadsFill":             ThreadsFill,
	"threeLineHorizontal":     ThreeLineHorizontal,
	"threeLineVertical":       ThreeLineVertical,
	"thumbsDown":              ThumbsDown,
	"thumbsUp":                ThumbsUp,
	"thunder":                 Thunder,
	"ticket":                  Ticket,
	"tiktokFill":              TiktokFill,
	"toggleOff":               ToggleOff,
	"toggleOffFill":           ToggleOffFill,
	"toggleOn":                ToggleOn,
	"toggleOnFill":            ToggleOnFill,
	"togoCup":                 TogoCup,
	"trash":                   Trash,
	"trashBin":                TrashBin,
	"trashCan":                TrashCan,
	"triangle":                Triangle,
	"triangleAlert":           TriangleAlert,
	"triangleAlertFill":       TriangleAlertFill,
	"triangleDown":            TriangleDown,
	"triangleDownFill":        TriangleDownFill,
	"triangleFill":            TriangleFill,
	"triangleLeft":            TriangleLeft,
	"triangleLeftFill":        TriangleLeftFill,
	"triangleRight":           TriangleRight,
	"triangleRightFill":       TriangleRightFill,
	"triangleUp":              TriangleUp,
	"triangleUpFill":          TriangleUpFill,
	"trophy":                  Trophy,
	"truck":                   Truck,
	"tumblrFill":              TumblrFill,
	"twitchFill":              TwitchFill,
	"twitterFill":             TwitterFill,
	"twoLineHorizontal":       TwoLineHorizontal,
	"twoLineVertical":         TwoLineVertical,
	"typescriptFill":          TypescriptFill,
	"umbrella":                Umbrella,
	"unsplashFill":            UnsplashFill,
	"utensils":                Utensils,
	"vapeKit":                 VapeKit,
	"vercelFill":              VercelFill,
	"victoryHand":             VictoryHand,
	"video":                   Video,
	"vimeoFill":               VimeoFill,
	"vkFill":                  VkFill,
	"vrAr":                    VrAr,
	"vscodeFill":              VscodeFill,
	"vueFill":                 VueFill,
	"wallet":                  Wallet,
	"watchDevice":             WatchDevice,
	"water":                   Water,
	"whatsappFill":            WhatsappFill,
	"width":                   Width,
	"wifi":                    Wifi,
	"wineGlass":               WineGlass,
	"xfill":                   Xfill,
	"xsmall":                  Xsmall,
	"yarnFill":                YarnFill,
	"yelpFill":                YelpFill,
	"youtubeFill":             YoutubeFill,
	"zoomFill":                ZoomFill,
	"zoomIn":                  ZoomIn,
	"zoomOut":                 ZoomOut,
}

func Air(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8h7a3 3 0 1 0-3-3M4 16h11a3 3 0 1 1-3 3M2 12h17a3 3 0 1 0-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AirplayAudio(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m12 16l6 5H6z"/><path d="M4 18a9.956 9.956 0 0 1-2-6C2 6.477 6.477 2 12 2s10 4.477 10 10a9.956 9.956 0 0 1-2 6"/><path d="M17.123 15.125a6 6 0 1 0-10.247-.002"/><path d="M14 12a2 2 0 1 0-4 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AirplayVideo(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 17H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-1"/><path d="m12 16l6 5H6z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Airpods(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><path d="M14 7c0 2.21 1.644 4 4 4s4-1.79 4-4s-1.644-4-4-4s-4 1.79-4 4Zm-4 0c0 2.21-1.644 4-4 4S2 9.21 2 7s1.644-4 4-4s4 1.79 4 4Z"/><path stroke-linecap="round" d="M14 7v12a2 2 0 0 0 2 2v0a2 2 0 0 0 2-2v-8"/><path d="M14 17h4M6 17h4"/><path stroke-linecap="round" d="M10 7v12a2 2 0 0 1-2 2v0a2 2 0 0 1-2-2v-8"/><path d="M20 4a5.408 5.408 0 0 0 0 6M4 4a5.408 5.408 0 0 1 0 6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Alarm(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="13" r="9"/><path d="M15.5 9.5L12 13m7 6l1 3M5 19l-1 3M2 5l3-3m14 0l3 3M12 4V2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignBottom(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="5.759" stroke-width="2"><path d="M21 22H3"/><path stroke-linejoin="round" d="M6 18V2h4v16zm8 0V8h4v10z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignHorizontalCenter(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="5.759" stroke-width="2"><path d="M12 3v18"/><path stroke-linejoin="round" d="M16 6h4v4h-4m-8 0H4V6h4m8 8h2v4h-2m-8-4H6v4h2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignLeft(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="5.759" stroke-width="2"><path d="M2 3v18"/><path stroke-linejoin="round" d="M6 6h16v4H6zm0 8h10v4H6z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignRight(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="5.759" stroke-width="2"><path d="M22 3v18"/><path stroke-linejoin="round" d="M2 6h16v4H2zm6 8h10v4H8z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignToBottom(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 22h18M12 2v15m-7-7l7 7l7-7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignToMiddle(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M21 12H3"/><path stroke-linejoin="round" d="M12 2v6m0 14v-6M9 5l3 3l3-3M9 19l3-3l3 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignToTop(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 22V7m-7 7l7-7l7 7M3 2h18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignTop(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="5.759" stroke-width="2"><path d="M21 2H3"/><path stroke-linejoin="round" d="M6 22V6h4v16zm8-6V6h4v10z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignVerticalCenter(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="5.759" stroke-width="2"><path d="M21 12H3"/><path stroke-linejoin="round" d="M6 16v4h4v-4m4 0v2h4v-2m-4-8V6h4v2m-8 0V4H6v4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AndroidFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.532 15.106a1.003 1.003 0 1 1 .001-2.007a1.003 1.003 0 0 1 0 2.007m-11.044 0a1.003 1.003 0 1 1 .001-2.007a1.003 1.003 0 0 1 0 2.007m11.4-6.018l2.006-3.459a.413.413 0 1 0-.721-.407l-2.027 3.5a12.243 12.243 0 0 0-5.13-1.108c-1.85 0-3.595.398-5.141 1.098l-2.027-3.5a.413.413 0 1 0-.72.407l1.995 3.458C2.696 10.947.345 14.417 0 18.523h24c-.334-4.096-2.675-7.565-6.112-9.435"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngularFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.174 12.594h3.652L12 8.095z"/><path fill="currentColor" d="M12 1L2 4.652l1.525 13.541L12 23l8.475-4.807L22 4.652zm6.24 16.786h-2.33l-1.257-3.212H9.347L8.09 17.786H5.76L12 3.431z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowBack(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m8 5l-5 5l5 5"/><path d="M3 10h8c5.523 0 10 4.477 10 10v1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowBackThick(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m2 11l7-9v5c11.953 0 13.332 9.678 13 15c-.502-2.685-.735-7-13-7v5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowBackThickFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 2a1 1 0 0 0-1.79-.614l-7 9a1 1 0 0 0 0 1.228l7 9A1 1 0 0 0 10 20v-3.99c5.379.112 7.963 1.133 9.261 2.243c1.234 1.055 1.46 2.296 1.695 3.596l.061.335a1 1 0 0 0 1.981-.122c.171-2.748-.086-6.73-2.027-10.061C19.087 8.768 15.695 6.282 10 6.022z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowClockwise(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19.734 16.06a8.923 8.923 0 0 1-3.915 3.978a8.706 8.706 0 0 1-5.471.832a8.795 8.795 0 0 1-4.887-2.64a9.067 9.067 0 0 1-2.388-5.079a9.136 9.136 0 0 1 1.044-5.53a8.904 8.904 0 0 1 4.069-3.815a8.7 8.7 0 0 1 5.5-.608c1.85.401 3.366 1.313 4.62 2.755c.151.16.735.806 1.22 1.781"/><path d="m15.069 7.813l5.04.907L21 3.59"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCounterClockwise(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4.266 16.06a8.923 8.923 0 0 0 3.915 3.978a8.706 8.706 0 0 0 5.471.832a8.796 8.796 0 0 0 4.887-2.64a9.067 9.067 0 0 0 2.388-5.079a9.137 9.137 0 0 0-1.044-5.53a8.904 8.904 0 0 0-4.068-3.815a8.7 8.7 0 0 0-5.5-.608c-1.85.401-3.367 1.313-4.62 2.755a7.62 7.62 0 0 0-1.22 1.781"/><path d="m8.931 7.813l-5.04.907L3 3.59"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCycle(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 12c0 6-4.39 10-9.806 10C7.792 22 4.24 19.665 3 16m-1-4C2 6 6.39 2 11.807 2C16.208 2 19.758 4.335 21 8"/><path d="m7 17l-4-1l-1 4M17 7l4 1l1-4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDown(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 20V4m-7 9l7 7l7-7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownLeft(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 8v10h10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownRight(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 18L6 6m2 12h10V8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownThick(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 21l9-7h-4.99L16 3H8v11H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowForward(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m16 5l5 5l-5 5"/><path d="M21 10h-8C7.477 10 3 14.477 3 20v1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowForwardThick(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m22 11l-7-9v5C3.047 7 1.668 16.678 2 22c.502-2.685.735-7 13-7v5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowForwardThickFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.676 1.054a1 1 0 0 1 1.113.332l7 9a1 1 0 0 1 0 1.228l-7 9A1 1 0 0 1 14 20v-3.99c-5.379.112-7.963 1.133-9.261 2.243c-1.234 1.055-1.46 2.296-1.695 3.596l-.061.335a1 1 0 0 1-1.981-.122c-.172-2.748.086-6.73 2.027-10.061C4.913 8.768 8.305 6.282 14 6.022V2a1 1 0 0 1 .675-.946" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeft(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m11 5l-7 7l7 7m-7-7h16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftThick(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 12l7-9v4.99L21 8v8H10v5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRepeat(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m18 2l3 3l-3 3M6 22l-3-3l3-3"/><path d="M21 5H10a7 7 0 0 0-7 7m0 7h11a7 7 0 0 0 7-7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRight(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 12h16m-7-7l7 7l-7 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightLeft(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 6H3m18 0l-4 4m4-4l-4-4M3 18h18M3 18l4 4m-4-4l4-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightThick(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m21 12l-7-9v4.99L3 8v8h11v5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowShuffle(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 19h3.908a2 2 0 0 0 1.682-.919L11.5 12l3.91-6.082A2 2 0 0 1 17.092 5H22m0 14h-4.908a2 2 0 0 1-1.682-.919L13.429 15M2 5h3.908a2 2 0 0 1 1.682.918L9.571 9"/><path d="m19 2l3 3l-3 3m0 8l3 3l-3 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUp(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 20V4m-7 7l7-7l7 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpDown(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 3v18M6 3l4 4M6 3L2 7m16 14V3m0 18l4-4m-4 4l-4-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpLeft(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6 6l12 12M16 6H6v10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpRight(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 6L6 18M8 6h10v10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpThick(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 3l9 7h-4.99L16 21H8V10H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ascending(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 3v18M6 3l3 3.333M6 3L3 6.333"/><path stroke-miterlimit="5.759" d="M14 21h8l-1-4h-7zm0-7h6l-1-4h-5zm0-7h4l-1-4h-3z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Attach(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 7.91V16a6 6 0 0 0 6 6v0a6 6 0 0 0 6-6V6a4 4 0 0 0-4-4v0a4 4 0 0 0-4 4v9.182a2 2 0 0 0 2 2v0a2 2 0 0 0 2-2V8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AugmentedReality(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path stroke-linecap="round" d="M10.971 6.617a2 2 0 0 1 2.058 0l3.486 2.092a1 1 0 0 1 .485.857v4.302a2 2 0 0 1-.971 1.715l-3 1.8a2 2 0 0 1-2.058 0l-3-1.8A2 2 0 0 1 7 13.868V9.566a1 1 0 0 1 .486-.857z"/><path d="m7 9l5 2.759m0 0L17 9m-5 2.759V17"/><path stroke-linecap="round" d="M6 2H4a2 2 0 0 0-2 2v2m16 16h2a2 2 0 0 0 2-2v-2m0-12V4a2 2 0 0 0-2-2h-2M2 18v2a2 2 0 0 0 2 2h2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Backspace(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="m17 15l-6-6m6 0l-6 6"/><path stroke-linejoin="round" d="M7.4 4.8A2 2 0 0 1 9 4h11a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H9a2 2 0 0 1-1.6-.8l-4.5-6a2 2 0 0 1 0-2.4z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BackspaceFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.6 4.2A3 3 0 0 1 9 3h11a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H9a3 3 0 0 1-2.4-1.2l-4.5-6a3 3 0 0 1 0-3.6zm11.107 5.507a1 1 0 0 0-1.414-1.414L14 10.586l-2.293-2.293a1 1 0 1 0-1.414 1.414L12.586 12l-2.293 2.293a1 1 0 0 0 1.414 1.414L14 13.414l2.293 2.293a1 1 0 0 0 1.414-1.414L15.414 12z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bank(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 11V8l-8-5l-8 5v3h8zM3 21h18M5 20v-5h2m10 5v-5h2m-8 5v-5h2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Basket(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linejoin="round" d="M2.31 11.243A1 1 0 0 1 3.28 10h17.44a1 1 0 0 1 .97 1.242l-1.811 7.243A2 2 0 0 1 17.939 20H6.061a2 2 0 0 1-1.94-1.515z"/><path stroke-linecap="round" d="M9 14v2m6-2v2m-9-6l4-6m8 6l-4-6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryCharging(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="12" x="2" y="6" rx="2"/><path d="M20 10h1.5a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-.5.5H20zm-8.6-1L9 12h4l-2.4 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryEmpty(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="12" x="2" y="6" rx="2"/><path d="M20 10h1.5a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-.5.5H20z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryFull(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="12" x="2" y="6" rx="2"/><path d="M7 10v4m4-4v4m4-4v4m5-4h1.5a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-.5.5H20z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryLow(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="12" x="2" y="6" rx="2"/><path d="M7 10v4m13-4h1.5a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-.5.5H20z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryMedium(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="12" x="2" y="6" rx="2"/><path d="M20 10h1.5a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-.5.5H20zM7 10v4m4-4v4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BehanceFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#akarIconsBehanceFill0)"><path fill="currentColor" d="M22 7h-7V5h7zm1.726 10c-.442 1.297-2.029 3-5.101 3c-3.074 0-5.564-1.729-5.564-5.675c0-3.91 2.325-5.92 5.466-5.92c3.082 0 4.964 1.782 5.375 4.426c.078.506.109 1.188.095 2.14H15.97c.13 3.211 3.483 3.312 4.588 2.029zm-7.686-4h4.965c-.105-1.547-1.136-2.219-2.477-2.219c-1.466 0-2.277.768-2.488 2.219m-9.574 6.988H0V5.021h6.953c5.476.081 5.58 5.444 2.72 6.906c3.461 1.26 3.577 8.061-3.207 8.061M3 11h3.584c2.508 0 2.906-3-.312-3H3zm3.391 3H3v3.016h3.341c3.055 0 2.868-3.016.05-3.016"/></g><defs><clipPath id="akarIconsBehanceFill0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bell(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12.721 5.003L11.255 5c-3.344-.008-6.247 2.709-6.27 6v3.79c0 .79-.1 1.561-.531 2.218l-.287.438C3.73 18.11 4.2 19 4.985 19h14.03c.785 0 1.254-.89.818-1.554l-.287-.438c-.43-.657-.531-1.429-.531-2.219v-3.788c-.04-3.292-2.95-5.99-6.294-5.998M15 19a3 3 0 1 1-6 0"/><path d="M12 2a2 2 0 0 1 2 2v1h-4V4a2 2 0 0 1 2-2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bicycle(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="6" cy="15" r="4"/><circle cx="18" cy="15" r="4"/><path d="m6 15l2-7h7.5M6 5h3m9 10L15 5h4m0 0h.5A1.5 1.5 0 0 1 21 6.5v0A1.5 1.5 0 0 1 19.5 8H19"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BitcoinFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g fill="currentColor" clip-path="url(#akarIconsBitcoinFill0)"><path d="M11.385 15.275c1.111-.004 3.54-.013 3.512-1.558c-.027-1.58-2.36-1.485-3.497-1.438c-.127.005-.24.01-.332.011l.052 2.987c.075-.002.165-.002.265-.002m-.118-4.353c.927-.001 2.95-.003 2.926-1.408c-.026-1.437-1.969-1.352-2.918-1.31c-.107.005-.2.009-.278.01l.047 2.709z"/><path fill-rule="evenodd" d="M9.096 23.641c6.43 1.603 12.942-2.31 14.545-8.738C25.244 8.474 21.33 1.962 14.9.36C8.474-1.244 1.962 2.67.36 9.1c-1.603 6.428 2.31 12.94 8.737 14.542m4.282-17.02c1.754.124 3.15.638 3.333 2.242c.136 1.174-.344 1.889-1.123 2.303c1.3.288 2.125 1.043 1.995 2.771c-.161 2.145-1.748 2.748-4.026 2.919l.038 2.25l-1.356.024l-.039-2.22c-.351.006-.711.01-1.084.008l.04 2.23l-1.356.024l-.04-2.254l-.383.003c-.194.001-.39.002-.586.006l-1.766.03l.241-1.624s1.004-.002.986-.017c.384-.008.481-.285.502-.459L8.693 11.3l.097-.002h.046a1.101 1.101 0 0 0-.144-.007l-.044-2.54c-.057-.274-.241-.59-.79-.58c.015-.02-.986.017-.986.017L6.846 6.74l1.872-.032v.007c.281-.005.57-.015.863-.026L9.543 4.46l1.356-.023l.038 2.184c.362-.013.726-.027 1.083-.033l-.038-2.17l1.357-.024z" clip-rule="evenodd"/></g><defs><clipPath id="akarIconsBitcoinFill0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Block(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M5 19L19 5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bluetooth(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 7l12 10l-6 5V2l6 5L5 17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Boat(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m16.926 19.382l-4.302 1.413a2 2 0 0 1-1.248 0L7.074 19.38a4 4 0 0 1-2.623-2.794L3 11l9 1l9-1l-1.451 5.587a4 4 0 0 1-2.623 2.794"/><path d="M6.497 7.257A2 2 0 0 1 8.354 6h7.292a2 2 0 0 1 1.857 1.257L19 11l-7 1l-7-1zM12 3v3m-1-3h2m-1 9v4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Book(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 6s1.5-2 5-2s5 2 5 2v14s-1.5-1-5-1s-5 1-5 1zm10 0s1.5-2 5-2s5 2 5 2v14s-1.5-1-5-1s-5 1-5 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookClose(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4.222v15.556C4 21.005 5.023 22 6.286 22h11.428C18.977 22 20 21.005 20 19.778V8.444a2 2 0 0 0-2-2H6.286C5.023 6.444 4 5.45 4 4.222m0 0C4 2.995 5.023 2 6.286 2h9.143c1.262 0 2.285.995 2.285 2.222v2.222"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookOpen(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 6s1.5-2 5-2s5 2 5 2v14s-1.5-1-5-1s-5 1-5 1zm10 0s1.5-2 5-2s5 2 5 2v14s-1.5-1-5-1s-5 1-5 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookmark(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><path d="M4 4v18l8-8l8 8V4a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2z" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BootstrapFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.985 2c-1.37 0-2.383 1.199-2.337 2.498c.043 1.25-.013 2.867-.42 4.186c-.41 1.322-1.1 2.16-2.228 2.268v1.215c1.128.107 1.819.945 2.227 2.268c.408 1.319.464 2.936.42 4.185c-.045 1.3.968 2.499 2.338 2.499h14.032c1.37 0 2.383-1.199 2.337-2.499c-.043-1.249.013-2.866.42-4.185c.409-1.323 1.098-2.16 2.226-2.268v-1.215c-1.128-.108-1.817-.946-2.226-2.268c-.407-1.32-.463-2.937-.42-4.186C21.4 3.198 20.386 2 19.017 2zM16.27 13.769c0 1.79-1.335 2.875-3.55 2.875H8.949a.407.407 0 0 1-.407-.407V6.881a.407.407 0 0 1 .407-.406h3.75c1.847 0 3.06 1 3.06 2.537c0 1.078-.816 2.043-1.855 2.213v.056c1.415.155 2.367 1.135 2.367 2.488M12.31 7.764h-2.15v3.038h1.811c1.4 0 2.172-.564 2.172-1.572c0-.944-.664-1.466-1.833-1.466m-2.15 4.243v3.347h2.23c1.457 0 2.23-.585 2.23-1.684c0-1.1-.794-1.663-2.324-1.663z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Box(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<rect width="18" height="18" x="3" y="3" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" rx="4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Briefcase(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="12" x="3" y="7" rx="2"/><path d="M9 6a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v1H9zm1 6l.211.106a4 4 0 0 0 3.578 0L14 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bug(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 9a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v5a7 7 0 0 1-7 7v0a7 7 0 0 1-7-7zm3-3v-.425c0-.981.384-1.96 1.326-2.238c1.525-.45 3.823-.45 5.348 0C15.616 3.615 16 4.594 16 5.575V6m2.5 1.5L22 4M5.5 7.5L2 4m4 14l-4 3m3-9H1.5m21 0H19m-1 6l4 3m-10-8v8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cake(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" d="M3 13a4 4 0 0 1 4-4h10a4 4 0 0 1 4 4v7a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/><path d="m3 13l2.914 2.331c1.187.95 2.9.855 3.975-.22v0a2.985 2.985 0 0 1 4.222 0v0a2.985 2.985 0 0 0 3.975.22L21 13"/><path stroke-linejoin="round" d="M12 6a2 2 0 0 1-2-2c0-.876.677-1.576 1.273-2.217L12 1l.727.783C13.323 2.424 14 3.124 14 4a2 2 0 0 1-2 2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calculator(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" d="M2 6a4 4 0 0 1 4-4h12a4 4 0 0 1 4 4v5H2z"/><path d="M18.5 16.5h-3"/><path stroke-linejoin="round" d="M12 11h10v7a4 4 0 0 1-4 4h-6zm0 0H2v7a4 4 0 0 0 4 4h6z"/><path d="M5.5 18L7 16.5m0 0L8.5 15M7 16.5L8.5 18M7 16.5L5.5 15"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="18" x="2" y="4" rx="4"/><path d="M8 2v4m8-4v4M2 10h20"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Camera(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m6.233 5.834l.445-2.226A2 2 0 0 1 8.64 2h6.72a2 2 0 0 1 1.962 1.608l.445 2.226a1.879 1.879 0 0 0 1.387 1.454A3.758 3.758 0 0 1 22 10.934V18a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4v-7.066a3.758 3.758 0 0 1 2.846-3.646a1.879 1.879 0 0 0 1.387-1.454"/><circle cx="12" cy="14" r="4"/><path d="M11 6h2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cart(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path d="M5 7h13.79a2 2 0 0 1 1.99 2.199l-.6 6A2 2 0 0 1 18.19 17H8.64a2 2 0 0 1-1.962-1.608z"/><path stroke-linecap="round" d="m5 7l-.81-3.243A1 1 0 0 0 3.22 3H2m6 18h2m6 0h2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatAdd(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M12 8v3m0 0v3m0-3h3m-3 0H9"/><path stroke-linejoin="round" d="M14 19c3.771 0 5.657 0 6.828-1.172C22 16.657 22 14.771 22 11c0-3.771 0-5.657-1.172-6.828C19.657 3 17.771 3 14 3h-4C6.229 3 4.343 3 3.172 4.172C2 5.343 2 7.229 2 11c0 3.771 0 5.657 1.172 6.828c.653.654 1.528.943 2.828 1.07"/><path d="M14 19c-1.236 0-2.598.5-3.841 1.145c-1.998 1.037-2.997 1.556-3.489 1.225c-.492-.33-.399-1.355-.212-3.404L6.5 17.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatApprove(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" d="M14 19c3.771 0 5.657 0 6.828-1.172C22 16.657 22 14.771 22 11c0-3.771 0-5.657-1.172-6.828C19.657 3 17.771 3 14 3h-4C6.229 3 4.343 3 3.172 4.172C2 5.343 2 7.229 2 11c0 3.771 0 5.657 1.172 6.828c.653.654 1.528.943 2.828 1.07"/><path stroke-linejoin="round" d="m9 11l2.25 2L15 9"/><path d="M14 19c-1.236 0-2.598.5-3.841 1.145c-1.998 1.037-2.997 1.556-3.489 1.225c-.492-.33-.399-1.355-.212-3.404L6.5 17.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubble(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" d="M14 19c3.771 0 5.657 0 6.828-1.172C22 16.657 22 14.771 22 11c0-3.771 0-5.657-1.172-6.828C19.657 3 17.771 3 14 3h-4C6.229 3 4.343 3 3.172 4.172C2 5.343 2 7.229 2 11c0 3.771 0 5.657 1.172 6.828c.653.654 1.528.943 2.828 1.07"/><path d="M14 19c-1.236 0-2.598.5-3.841 1.145c-1.998 1.037-2.997 1.556-3.489 1.225c-.492-.33-.399-1.355-.212-3.404L6.5 17.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatDots(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" d="M14 19c3.771 0 5.657 0 6.828-1.172C22 16.657 22 14.771 22 11c0-3.771 0-5.657-1.172-6.828C19.657 3 17.771 3 14 3h-4C6.229 3 4.343 3 3.172 4.172C2 5.343 2 7.229 2 11c0 3.771 0 5.657 1.172 6.828c.653.654 1.528.943 2.828 1.07"/><path d="M12 11v.01M8 11v.01m8-.01v.01M14 19c-1.236 0-2.598.5-3.841 1.145c-1.998 1.037-2.997 1.556-3.489 1.225c-.492-.33-.399-1.355-.212-3.404L6.5 17.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatEdit(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M14 19c3.771 0 5.657 0 6.828-1.172C22 16.657 22 14.771 22 11c0-3.771 0-5.657-1.172-6.828C19.657 3 17.771 3 14 3h-4C6.229 3 4.343 3 3.172 4.172C2 5.343 2 7.229 2 11c0 3.771 0 5.657 1.172 6.828c.653.654 1.528.943 2.828 1.07"/><path stroke-linecap="round" stroke-linejoin="round" d="M15.207 6.793a1 1 0 0 0-1.418.003l-4.55 4.597a2 2 0 0 0-.54 1.015l-.18.896a1 1 0 0 0 1.177 1.177l.896-.18a2 2 0 0 0 1.015-.54l4.597-4.55a1 1 0 0 0 .003-1.418z"/><path d="m12.5 9.5l1 1"/><path stroke-linecap="round" d="M14 19c-1.236 0-2.598.5-3.841 1.145c-1.998 1.037-2.997 1.556-3.489 1.225c-.492-.33-.399-1.355-.212-3.404L6.5 17.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatError(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" d="M14 19c3.771 0 5.657 0 6.828-1.172C22 16.657 22 14.771 22 11c0-3.771 0-5.657-1.172-6.828C19.657 3 17.771 3 14 3h-4C6.229 3 4.343 3 3.172 4.172C2 5.343 2 7.229 2 11c0 3.771 0 5.657 1.172 6.828c.653.654 1.528.943 2.828 1.07"/><path d="M14 19c-1.236 0-2.598.5-3.841 1.145c-1.998 1.037-2.997 1.556-3.489 1.225c-.492-.33-.399-1.355-.212-3.404L6.5 17.5m3.379-8.621L12 11m0 0l2.121 2.121M12 11l2.121-2.121M12 11l-2.121 2.121"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatQuestion(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" d="M14 19c3.771 0 5.657 0 6.828-1.172C22 16.657 22 14.771 22 11c0-3.771 0-5.657-1.172-6.828C19.657 3 17.771 3 14 3h-4C6.229 3 4.343 3 3.172 4.172C2 5.343 2 7.229 2 11c0 3.771 0 5.657 1.172 6.828c.653.654 1.528.943 2.828 1.07"/><path d="M10 8.484C10.5 7.494 11 7 12 7c1.246 0 2 .989 2 1.978s-.5 1.033-2 2.022v1m0 2.5v.5m2 4c-1.236 0-2.598.5-3.841 1.145c-1.998 1.037-2.997 1.556-3.489 1.225c-.492-.33-.399-1.355-.212-3.404L6.5 17.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatRemove(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M15 11H9"/><path stroke-linejoin="round" d="M14 19c3.771 0 5.657 0 6.828-1.172C22 16.657 22 14.771 22 11c0-3.771 0-5.657-1.172-6.828C19.657 3 17.771 3 14 3h-4C6.229 3 4.343 3 3.172 4.172C2 5.343 2 7.229 2 11c0 3.771 0 5.657 1.172 6.828c.653.654 1.528.943 2.828 1.07"/><path d="M14 19c-1.236 0-2.598.5-3.841 1.145c-1.998 1.037-2.997 1.556-3.489 1.225c-.492-.33-.399-1.355-.212-3.404L6.5 17.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Check(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 12l6 6L20 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckBox(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="4"/><path d="m9 12l2.25 2L15 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckBoxFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7 2a5 5 0 0 0-5 5v10a5 5 0 0 0 5 5h10a5 5 0 0 0 5-5V7a5 5 0 0 0-5-5zm8.73 8.684a1 1 0 1 0-1.46-1.368l-3.083 3.29l-1.523-1.353a1 1 0 0 0-1.328 1.494l2.25 2a1 1 0 0 0 1.393-.063z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckIn(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 2a6 6 0 0 0-6 6c0 1.419.302 2.348 1.125 3.375L12 17l4.875-5.625C17.698 10.348 18 9.419 18 8a6 6 0 0 0-6-6"/><path d="M5 15.143C3.149 15.87 2 16.881 2 18c0 2.21 4.477 4 10 4s10-1.79 10-4c0-1.119-1.149-2.13-3-2.857"/><circle cx="12" cy="8" r="1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chess(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.533 3.81L8 2l1 4l-5.37 4.475A1.75 1.75 0 0 0 3 11.82v0c0 .617.537 1.088 1.127.986L9 12l-2.097 7h10.614l1.283-5.745c.913-4.088-1.386-8.21-5.267-9.445M4 21a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v1H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDown(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 9l8 8l8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDownSmall(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6 10l6 6l6-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronHorizontal(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m17 8l4 4l-4 4M7 8l-4 4l4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeft(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 4l-8 8l8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeftSmall(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14 6l-6 6l6 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRight(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m8 4l8 8l-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRightSmall(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9 6l6 6l-6 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUp(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 15l8-8l8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUpSmall(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6 14l6-6l6 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronVertical(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m8 17l4 4l4-4M8 7l4-4l4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Circle(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="12" cy="12" r="10" fill="none" stroke="currentColor" stroke-width="2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleAlert(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><path stroke-linecap="round" d="M12 7v6m0 3.5v.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleAlertFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11s11-4.925 11-11S18.075 1 12 1m1 6a1 1 0 1 0-2 0v6a1 1 0 1 0 2 0zm0 9.5a1 1 0 1 0-2 0v.5a1 1 0 1 0 2 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleCheck(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="m8 12.5l3 3l5-6"/><circle cx="12" cy="12" r="10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleCheckFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11s11-4.925 11-11S18.075 1 12 1m4.768 9.14a1 1 0 1 0-1.536-1.28l-4.3 5.159l-2.225-2.226a1 1 0 0 0-1.414 1.414l3 3a1 1 0 0 0 1.475-.067z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleChevronDown(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><path stroke-linecap="round" stroke-linejoin="round" d="m8 10.5l4 4l4-4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleChevronDownFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11s11-4.925 11-11S18.075 1 12 1M8.707 9.793a1 1 0 0 0-1.414 1.414l4 4a1 1 0 0 0 1.414 0l4-4a1 1 0 0 0-1.414-1.414L12 13.086z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleChevronLeft(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><path stroke-linecap="round" stroke-linejoin="round" d="m13.5 8l-4 4l4 4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleChevronLeftFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11s11-4.925 11-11S18.075 1 12 1m2.207 7.707a1 1 0 0 0-1.414-1.414l-4 4a1 1 0 0 0 0 1.414l4 4a1 1 0 0 0 1.414-1.414L10.914 12z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleChevronRight(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><path stroke-linecap="round" stroke-linejoin="round" d="m10.5 8l4 4l-4 4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleChevronRightFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11s11-4.925 11-11S18.075 1 12 1M9.793 8.707a1 1 0 0 1 1.414-1.414l4 4a1 1 0 0 1 0 1.414l-4 4a1 1 0 0 1-1.414-1.414L13.086 12z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleChevronUp(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><path stroke-linecap="round" stroke-linejoin="round" d="m8 13.5l4-4l4 4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleChevronUpFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11s11-4.925 11-11S18.075 1 12 1M8.707 14.207a1 1 0 0 1-1.414-1.414l4-4a1 1 0 0 1 1.414 0l4 4a1 1 0 0 1-1.414 1.414L12 10.914z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="12" cy="12" r="11" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleMinus(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" d="M16 12H8"/><circle cx="12" cy="12" r="10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleMinusFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11s11-4.925 11-11S18.075 1 12 1M8 11a1 1 0 1 0 0 2h8a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CirclePlus(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" d="M12 8v4m0 0v4m0-4h4m-4 0H8"/><circle cx="12" cy="12" r="10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CirclePlusFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11s11-4.925 11-11S18.075 1 12 1m1 15a1 1 0 1 1-2 0v-3H8a1 1 0 1 1 0-2h3V8a1 1 0 1 1 2 0v3h3a1 1 0 1 1 0 2h-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleTriangleDown(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><path stroke-linecap="round" stroke-linejoin="round" d="m12 16l-4-6h8z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleTriangleDownFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11s11-4.925 11-11S18.075 1 12 1m-3.2 9a.809.809 0 0 0-.705.396a.71.71 0 0 0 .04.77l3.2 4.5A.815.815 0 0 0 12 16a.815.815 0 0 0 .666-.334l3.2-4.5a.71.71 0 0 0 .04-.77A.809.809 0 0 0 15.2 10z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleTriangleLeft(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><path stroke-linecap="round" stroke-linejoin="round" d="m8 12l6-4v8z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleTriangleLeftFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11s11-4.925 11-11S18.075 1 12 1m2 7.8a.809.809 0 0 0-.396-.705a.71.71 0 0 0-.77.04l-4.5 3.2A.815.815 0 0 0 8 12c0 .268.125.517.334.666l4.5 3.2a.71.71 0 0 0 .77.04A.809.809 0 0 0 14 15.2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleTriangleRight(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><path stroke-linecap="round" stroke-linejoin="round" d="m16 12l-6-4v8z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleTriangleRightFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11s11-4.925 11-11S18.075 1 12 1m-2 7.8c0-.295.152-.566.396-.705a.71.71 0 0 1 .77.04l4.5 3.2A.815.815 0 0 1 16 12a.815.815 0 0 1-.334.666l-4.5 3.2a.71.71 0 0 1-.77.04A.809.809 0 0 1 10 15.2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleTriangleUp(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><path stroke-linecap="round" stroke-linejoin="round" d="m12 8l-4 6h8z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleTriangleUpFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11s11-4.925 11-11S18.075 1 12 1M8.8 14a.809.809 0 0 1-.705-.396a.71.71 0 0 1 .04-.77l3.2-4.5A.815.815 0 0 1 12 8c.268 0 .517.125.666.334l3.2 4.5a.71.71 0 0 1 .04.77a.809.809 0 0 1-.706.396z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleX(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" d="M15 15L9 9m6 0l-6 6"/><circle cx="12" cy="12" r="10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleXfill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11s11-4.925 11-11S18.075 1 12 1m3.707 8.707a1 1 0 0 0-1.414-1.414L12 10.586L9.707 8.293a1 1 0 1 0-1.414 1.414L10.586 12l-2.293 2.293a1 1 0 1 0 1.414 1.414L12 13.414l2.293 2.293a1 1 0 0 0 1.414-1.414L13.414 12z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clipboard(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" d="M15.5 4H18a2 2 0 0 1 2 2v13a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2.5"/><path stroke-linejoin="round" d="M8.621 3.515A2 2 0 0 1 10.561 2h2.877a2 2 0 0 1 1.94 1.515L16 6H8z"/><path d="M9 12h6m-6 4h6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clock(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m15 16l-2.414-2.414A2 2 0 0 1 12 12.172V6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloud(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5.034 11.117A4.002 4.002 0 0 0 6 19h11a5 5 0 1 0-1.17-9.862L14.5 9.5"/><path d="M15.83 9.138a5.5 5.5 0 0 0-10.796 1.98S5.187 12 5.5 12.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudDownload(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 22v-9m0 9l-2.5-2m2.5 2l2.5-2M5.034 9.117A4.002 4.002 0 0 0 6 17h1"/><path d="M15.83 7.138a5.5 5.5 0 0 0-10.796 1.98S5.187 10 5.5 10.5"/><path d="M17 17a5 5 0 1 0-1.17-9.862L14.5 7.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudUpload(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 12v9m0-9l-2.5 2m2.5-2l2.5 2M5.034 9.117A4.002 4.002 0 0 0 6 17h1"/><path d="M15.83 7.138a5.5 5.5 0 0 0-10.796 1.98S5.187 10 5.5 10.5"/><path d="M17 17a5 5 0 1 0-1.17-9.862L14.5 7.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodepenFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#akarIconsCodepenFill0)"><path fill="currentColor" fill-rule="evenodd" d="M11.372.19c.38-.253.875-.253 1.256 0L23.492 7.4c.317.21.508.565.508.946v7.308c0 .38-.19.736-.508.947l-10.864 7.21c-.38.252-.875.252-1.256 0L.508 16.6A1.136 1.136 0 0 1 0 15.654V8.346c0-.38.19-.736.508-.947zm-9.1 10.273v3.058l2.288-1.54zm4.337 2.878L3.18 15.648l7.684 5.1v-4.583zm6.527 2.824v4.582l7.684-5.1l-3.43-2.306zm6.304-4.183l2.288 1.54v-3.06zm1.37-3.636l-3.41 2.263l-4.264-2.868V3.253zm-9.946-5.093V7.74l-4.263 2.868L3.19 8.346zM12 9.715l-3.35 2.254L12 14.192l3.35-2.223z" clip-rule="evenodd"/></g><defs><clipPath id="akarIconsCodepenFill0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Coffee(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11 7c4.418 0 8 .895 8 2s-3.582 2-8 2s-8-.895-8-2c0-.357.375-.693 1.033-.984"/><path d="M3 9v9.343c0 1.061.44 2.08 1.409 2.513C5.624 21.399 7.711 22 11 22c3.29 0 5.377-.601 6.591-1.144c.968-.433 1.409-1.452 1.409-2.513V9m0 1a3 3 0 0 1 3 3v0a3 3 0 0 1-3 3v0M7 3v4m4-5v2m4 0v3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Coin(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><ellipse cx="9.5" cy="10" stroke-linecap="round" stroke-linejoin="round" rx="9.5" ry="10" transform="matrix(-1 0 0 1 20 2)"/><path stroke-linecap="round" stroke-linejoin="round" d="M13 8.8a3.583 3.583 0 0 0-2.25-.8C8.679 8 7 9.79 7 12s1.679 4 3.75 4c.844 0 1.623-.298 2.25-.8"/><path d="M10 2c4.333 0 13 1 13 10s-8.667 10-13 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Command(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="2" d="M9.3 17.85a3.15 3.15 0 1 1-3.15-3.15h11.7a3.15 3.15 0 1 1-3.15 3.15V6.15a3.15 3.15 0 1 1 3.15 3.15H6.15A3.15 3.15 0 1 1 9.3 6.15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Comment(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" d="M14 19c3.771 0 5.657 0 6.828-1.172C22 16.657 22 14.771 22 11c0-3.771 0-5.657-1.172-6.828C19.657 3 17.771 3 14 3h-4C6.229 3 4.343 3 3.172 4.172C2 5.343 2 7.229 2 11c0 3.771 0 5.657 1.172 6.828c.653.654 1.528.943 2.828 1.07"/><path d="M14 19c-1.236 0-2.598.5-3.841 1.145c-1.998 1.037-2.997 1.556-3.489 1.225c-.492-.33-.399-1.355-.212-3.404L6.5 17.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentAdd(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M12 8v3m0 0v3m0-3h3m-3 0H9"/><path stroke-linejoin="round" d="M14 19c3.771 0 5.657 0 6.828-1.172C22 16.657 22 14.771 22 11c0-3.771 0-5.657-1.172-6.828C19.657 3 17.771 3 14 3h-4C6.229 3 4.343 3 3.172 4.172C2 5.343 2 7.229 2 11c0 3.771 0 5.657 1.172 6.828c.653.654 1.528.943 2.828 1.07"/><path d="M14 19c-1.236 0-2.598.5-3.841 1.145c-1.998 1.037-2.997 1.556-3.489 1.225c-.492-.33-.399-1.355-.212-3.404L6.5 17.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Computing(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="14" height="14" x="5" y="5" rx="2"/><path d="M8 5V2m8 3V3l1-1m-1 20v-3m-7 3v-3M5 8H2m20 0h-3m3 8h-3M5 16H3l-1 1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copy(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 4v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V7.242a2 2 0 0 0-.602-1.43L16.083 2.57A2 2 0 0 0 14.685 2H10a2 2 0 0 0-2 2"/><path d="M16 18v2a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V9a2 2 0 0 1 2-2h2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCard(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="16" x="2" y="4" rx="2"/><path d="M12 10.016A2.794 2.794 0 0 0 9.857 9C8.28 9 7 10.343 7 12s1.28 3 2.857 3c.854 0 1.62-.393 2.143-1.016M17 12c0 1.657-1.28 3-2.857 3c-1.578 0-2.857-1.343-2.857-3s1.279-3 2.857-3C15.72 9 17 10.343 17 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCardAltOne(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="16" x="2" y="4" rx="2"/><path d="M6 8h2v2H6z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cross(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M20 20L4 4m16 0L4 20"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crown(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m2 8l1.304 1.043a4 4 0 0 0 5.995-1.181L12 3l2.701 4.862a4 4 0 0 0 5.996 1.18L22 8l-1.754 8.77a2.564 2.564 0 0 1-1.367 1.79v0a15.381 15.381 0 0 1-13.758 0v0a2.564 2.564 0 0 1-1.367-1.79z"/><path d="M8 15c2.596 1.333 5.404 1.333 8 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CssFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.502 0h2.578v1.078h-1.5v1.078h1.5v1.078H7.502zm3.093 0h2.579v.938h-1.5v.187h1.5v2.156h-2.579v-.984h1.5v-.188h-1.5zm3.095 0h2.577v.938h-1.5v.187h1.5v2.156H13.69v-.984h1.5v-.188h-1.5z"/><path fill="currentColor" fill-rule="evenodd" d="m11.991 24l-6.944-1.928L3 4.717h18L18.954 22.07zM7.047 12.573l.191 2.128h7.377l-.247 2.76l-2.374.642h-.002l-2.37-.64l-.152-1.697H7.333l.298 3.342l4.36 1.21l4.367-1.21l.532-5.964l.052-.571l.384-4.309H6.664l.194 2.129h8.136l-.194 2.18z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cursor(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="m3 3l7 19l2.051-6.154a6 6 0 0 1 3.795-3.795L22 10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cut(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="7" cy="18" r="3"/><path d="M15 15L7 3m2 12l3-4.5M17 3l-3 4.5"/><circle cx="17" cy="18" r="3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dashboard(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 5a2 2 0 0 1 2-2h6v18H4a2 2 0 0 1-2-2zm12-2h6a2 2 0 0 1 2 2v5h-8zm0 11h8v5a2 2 0 0 1-2 2h-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Data(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><ellipse cx="12" cy="6" rx="8" ry="3"/><path d="M6.037 12C4.77 12.53 4 13.232 4 14c0 1.657 3.582 3 8 3s8-1.343 8-3c0-.768-.77-1.47-2.037-2"/><path d="M4 6v4c0 1.657 3.582 3 8 3s8-1.343 8-3V6M4 14v4c0 1.657 3.582 3 8 3s8-1.343 8-3v-4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dental(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.363C9 2.732 3 1.23 3 8.277c0 5.492 1.188 9.756 3.005 12.141c.645.847 2.216.584 2.888-.265a1.22 1.22 0 0 0 .174-.328l1.063-2.8c.654-1.72 3.086-1.72 3.74 0l1.063 2.8c.045.116.097.23.174.328c.672.85 2.243 1.112 2.888.265C19.812 18.033 21 13.77 21 8.277c0-7.046-6-5.545-9-3.914m0 0L15 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Descending(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 21V3m0 18l3-3.333M6 21l-3-3.333"/><path stroke-miterlimit="5.759" d="M14 3h8l-1 4h-7zm0 7h6l-1 4h-5zm0 7h4l-1 4h-3z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DesktopDevice(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 6a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2zm6 14h8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Devices(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 15H4V6a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v2M2 18h12"/><path d="M14 9.2c0-.663.597-1.2 1.333-1.2h5.334C21.403 8 22 8.537 22 9.2v9.6c0 .663-.597 1.2-1.333 1.2h-5.334C14.597 20 14 19.463 14 18.8zm4 7.8h.01"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Diamond(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="2" d="M4.511 13.542c-.681-.852-.681-2.232 0-3.084l6.256-7.82c.68-.85 1.785-.85 2.467 0l6.255 7.82c.681.852.681 2.232 0 3.084l-6.256 7.82c-.68.85-1.785.85-2.466 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceFive(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" stroke-linejoin="round" rx="2"/><path d="M8 8h.5m7 0h.5m-4 4h.5m-.25-.25v.5M8 16h.5m7 0h.5M8.25 7.75v.5m0 7.5v.5m7.5-8.5v.5m0 7.5v.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceFour(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" stroke-linejoin="round" rx="2"/><path d="M8 8h.5m7 0h.5m-8 8h.5m7 0h.5M8.25 7.75v.5m0 7.5v.5m7.5-8.5v.5m0 7.5v.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceOne(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" stroke-linejoin="round" rx="2"/><path d="M12.25 11.75v.5M12 12h.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceSix(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" stroke-linejoin="round" rx="2"/><path d="M8 8h.5m-.25-.25v.5m0 3.5v.5m0 3.5v.5M15.5 8h.5m-.25-.25v.5m0 3.5v.5m0 3.5v.5M8 12h.5m7 0h.5m-8 4h.5m7 0h.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceThree(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" stroke-linejoin="round" rx="2"/><path d="M12.25 11.75v.5m-4-4.5v.5m7.5 7.5v.5M8 8h.5m3.5 4h.5m3 4h.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceTwo(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" stroke-linejoin="round" rx="2"/><path d="M8.25 7.75v.5m7.5 7.5v.5M8 8h.5m7 8h.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiscordFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.317 4.492c-1.53-.69-3.17-1.2-4.885-1.49a.075.075 0 0 0-.079.036c-.21.369-.444.85-.608 1.23a18.566 18.566 0 0 0-5.487 0a12.36 12.36 0 0 0-.617-1.23A.077.077 0 0 0 8.562 3c-1.714.29-3.354.8-4.885 1.491a.07.07 0 0 0-.032.027C.533 9.093-.32 13.555.099 17.961a.08.08 0 0 0 .031.055a20.03 20.03 0 0 0 5.993 2.98a.078.078 0 0 0 .084-.026a13.83 13.83 0 0 0 1.226-1.963a.074.074 0 0 0-.041-.104a13.201 13.201 0 0 1-1.872-.878a.075.075 0 0 1-.008-.125c.126-.093.252-.19.372-.287a.075.075 0 0 1 .078-.01c3.927 1.764 8.18 1.764 12.061 0a.075.075 0 0 1 .079.009c.12.098.245.195.372.288a.075.075 0 0 1-.006.125c-.598.344-1.22.635-1.873.877a.075.075 0 0 0-.041.105c.36.687.772 1.341 1.225 1.962a.077.077 0 0 0 .084.028a19.963 19.963 0 0 0 6.002-2.981a.076.076 0 0 0 .032-.054c.5-5.094-.838-9.52-3.549-13.442a.06.06 0 0 0-.031-.028M8.02 15.278c-1.182 0-2.157-1.069-2.157-2.38c0-1.312.956-2.38 2.157-2.38c1.21 0 2.176 1.077 2.157 2.38c0 1.312-.956 2.38-2.157 2.38m7.975 0c-1.183 0-2.157-1.069-2.157-2.38c0-1.312.955-2.38 2.157-2.38c1.21 0 2.176 1.077 2.157 2.38c0 1.312-.946 2.38-2.157 2.38"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DjangoFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.172 0h-4.176v5.932a7.21 7.21 0 0 0-1.816-.2C4.816 5.731 2 8.305 2 12.273c0 4.118 2.655 6.263 7.755 6.268c1.703 0 3.278-.15 5.417-.53zM9.734 8.977c.516 0 .92.05 1.408.2v6.248c-.596.075-.972.1-1.434.1c-2.14 0-3.305-1.142-3.305-3.21c0-2.125 1.22-3.338 3.331-3.338" clip-rule="evenodd"/><path fill="currentColor" d="M22 15.233V6.215h-4.17v7.675c0 3.387-.188 4.674-.785 5.786c-.57 1.087-1.462 1.8-3.305 2.606L17.615 24c1.843-.862 2.735-1.643 3.412-2.88c.726-1.288.973-2.782.973-5.887M21.585 0h-4.176v3.993h4.176z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Door(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 2h11a3 3 0 0 1 3 3v14a1 1 0 0 1-1 1h-3"/><path d="m5 2l7.588 1.518A3 3 0 0 1 15 6.459V20.78a1 1 0 0 1-1.196.98l-7.196-1.438A2 2 0 0 1 5 18.36zm7 10v2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotGrid(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><circle cx="4" cy="4" r="1" transform="rotate(90 4 4)"/><circle cx="12" cy="4" r="1" transform="rotate(90 12 4)"/><circle cx="20" cy="4" r="1" transform="rotate(90 20 4)"/><circle cx="4" cy="12" r="1" transform="rotate(90 4 12)"/><circle cx="12" cy="12" r="1" transform="rotate(90 12 12)"/><circle cx="20" cy="12" r="1" transform="rotate(90 20 12)"/><circle cx="4" cy="20" r="1" transform="rotate(90 4 20)"/><circle cx="12" cy="20" r="1" transform="rotate(90 12 20)"/><circle cx="20" cy="20" r="1" transform="rotate(90 20 20)"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotGridFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 2a2 2 0 1 1 0 4a2 2 0 0 1 0-4m8 0a2 2 0 1 1 0 4a2 2 0 0 1 0-4m10 2a2 2 0 1 0-4 0a2 2 0 0 0 4 0M4 10a2 2 0 1 1 0 4a2 2 0 0 1 0-4m10 2a2 2 0 1 0-4 0a2 2 0 0 0 4 0m6-2a2 2 0 1 1 0 4a2 2 0 0 1 0-4M6 20a2 2 0 1 0-4 0a2 2 0 0 0 4 0m6-2a2 2 0 1 1 0 4a2 2 0 0 1 0-4m10 2a2 2 0 1 0-4 0a2 2 0 0 0 4 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleCheck(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m2 12l5.25 5l2.625-3M8 12l5.25 5L22 7m-6 0l-3.5 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleSword(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 19.2L3.8 21m.9-7.2l.9 3.6m0 0l3.6.9m-3.6-.9l-2.7 2.7M16.4 3.9l-9 9l.45 2.25l2.25.45l9-9L20 3z"/><path d="M22 19.2L20.2 21m-.9-7.2l-.9 3.6m0 0l2.7 2.7m-2.7-2.7l-1.8.45l-1.8.45M9.3 11L4.9 6.6L4 3l3.6.9L12 8.3m.1 5.5l1.8 1.8l2.25-.45l.45-2.25l-1.8-1.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Download(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15V3m0 12l-4-4m4 4l4-4M2 17l.621 2.485A2 2 0 0 0 4.561 21h14.878a2 2 0 0 0 1.94-1.515L22 17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Draft(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M13.5 6v4H16m-1.315-8H10a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V7.242a2 2 0 0 0-.602-1.43L16.083 2.57A2 2 0 0 0 14.685 2"/><path d="M16 18v2a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V9a2 2 0 0 1 2-2h2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DragHorizontal(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><circle cx="20" cy="8" r="1" transform="rotate(-180 20 8)"/><circle cx="20" cy="16" r="1" transform="rotate(-180 20 16)"/><circle cx="12" cy="8" r="1" transform="rotate(-180 12 8)"/><circle cx="12" cy="16" r="1" transform="rotate(-180 12 16)"/><circle cx="4" cy="8" r="1" transform="rotate(-180 4 8)"/><circle cx="4" cy="16" r="1" transform="rotate(-180 4 16)"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DragHorizontalFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M22 8a2 2 0 1 1-4 0a2 2 0 0 1 4 0m0 8a2 2 0 1 1-4 0a2 2 0 0 1 4 0m-8 0a2 2 0 1 1-4 0a2 2 0 0 1 4 0M6 8a2 2 0 1 1-4 0a2 2 0 0 1 4 0m0 8a2 2 0 1 1-4 0a2 2 0 0 1 4 0m8-8a2 2 0 1 1-4 0a2 2 0 0 1 4 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DragVertical(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><circle cx="8" cy="4" r="1" transform="rotate(90 8 4)"/><circle cx="16" cy="4" r="1" transform="rotate(90 16 4)"/><circle cx="8" cy="12" r="1" transform="rotate(90 8 12)"/><circle cx="16" cy="12" r="1" transform="rotate(90 16 12)"/><circle cx="8" cy="20" r="1" transform="rotate(90 8 20)"/><circle cx="16" cy="20" r="1" transform="rotate(90 16 20)"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DragVerticalFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16 2a2 2 0 1 1 0 4a2 2 0 0 1 0-4m0 8a2 2 0 1 1 0 4a2 2 0 0 1 0-4m-8 8a2 2 0 1 1 0 4a2 2 0 0 1 0-4m8 0a2 2 0 1 1 0 4a2 2 0 0 1 0-4M8 2a2 2 0 1 1 0 4a2 2 0 0 1 0-4m0 8a2 2 0 1 1 0 4a2 2 0 0 1 0-4" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DribbbleFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#akarIconsDribbbleFill0)"><path fill="currentColor" fill-rule="evenodd" d="M12 0C5.375 0 0 5.375 0 12s5.375 12 12 12c6.612 0 12-5.375 12-12S18.612 0 12 0m7.926 5.531a10.202 10.202 0 0 1 2.317 6.378c-.338-.065-3.722-.755-7.132-.325c-.079-.17-.144-.352-.222-.534a30.53 30.53 0 0 0-.676-1.484c3.774-1.536 5.492-3.748 5.713-4.035M12 1.771c2.603 0 4.985.975 6.794 2.576c-.182.26-1.731 2.33-5.375 3.696c-1.68-3.084-3.54-5.61-3.827-6A10.424 10.424 0 0 1 12 1.77m-4.36.962c.273.365 2.095 2.903 3.8 5.922c-4.79 1.276-9.02 1.25-9.475 1.25c.664-3.176 2.812-5.818 5.675-7.172m-5.896 9.28V11.7c.443.013 5.414.078 10.53-1.458c.299.573.572 1.158.832 1.744c-.13.039-.273.078-.403.117c-5.284 1.705-8.096 6.365-8.33 6.755a10.226 10.226 0 0 1-2.629-6.846M12 22.256a10.18 10.18 0 0 1-6.286-2.16c.182-.378 2.264-4.387 8.043-6.404c.026-.013.04-.013.065-.026c1.445 3.735 2.03 6.872 2.187 7.77c-1.237.534-2.59.82-4.009.82m5.714-1.757c-.104-.625-.651-3.618-1.992-7.302c3.215-.507 6.026.326 6.378.443c-.443 2.85-2.083 5.31-4.386 6.859" clip-rule="evenodd"/></g><defs><clipPath id="akarIconsDribbbleFill0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DropboxFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m6 2l6 3.75L6 9.5L0 5.75zm12 0l6 3.75l-6 3.75l-6-3.75zM0 13.25L6 9.5l6 3.75L6 17zM18 9.5l6 3.75L18 17l-6-3.75zM6 18.25l6-3.75l6 3.75L12 22z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Edit(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m16.475 5.408l2.117 2.117m-.756-3.982L12.109 9.27a2.118 2.118 0 0 0-.58 1.082L11 13l2.648-.53c.41-.082.786-.283 1.082-.579l5.727-5.727a1.853 1.853 0 1 0-2.621-2.621"/><path d="M19 15v3a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Enlarge(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.5 9.5L21 3m0 0h-6m6 0v6M3 21l6.5-6.5M3 21v-6m0 6h6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Envelope(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path stroke-linecap="round" d="M2 6a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2z"/><path d="m2 8l7.501 6.001a4 4 0 0 0 4.998 0L22 8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Equal(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" d="M16 14H8m8-4H8"/><circle cx="12" cy="12" r="10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EqualFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m6 2a1 1 0 0 1 1-1h8a1 1 0 1 1 0 2H8a1 1 0 0 1-1-1m1-5a1 1 0 1 0 0 2h8a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eye(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21.257 10.962c.474.62.474 1.457 0 2.076C19.764 14.987 16.182 19 12 19c-4.182 0-7.764-4.013-9.257-5.962a1.692 1.692 0 0 1 0-2.076C4.236 9.013 7.818 5 12 5c4.182 0 7.764 4.013 9.257 5.962"/><circle cx="12" cy="12" r="3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeClosed(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 10s3.5 4 10 4s10-4 10-4M4 11.645L2 14m20 0l-1.996-2.352M8.914 13.68L8 16.5m7.063-2.812L16 16.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeOpen(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21.257 10.962c.474.62.474 1.457 0 2.076C19.764 14.987 16.182 19 12 19c-4.182 0-7.764-4.013-9.257-5.962a1.692 1.692 0 0 1 0-2.076C4.236 9.013 7.818 5 12 5c4.182 0 7.764 4.013 9.257 5.962"/><circle cx="12" cy="12" r="3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeSlashed(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6.873 17.129c-1.845-1.31-3.305-3.014-4.13-4.09a1.693 1.693 0 0 1 0-2.077C4.236 9.013 7.818 5 12 5c1.876 0 3.63.807 5.13 1.874"/><path d="M14.13 9.887a3 3 0 1 0-4.243 4.242M4 20L20 4M10 18.704A7.124 7.124 0 0 0 12 19c4.182 0 7.764-4.013 9.257-5.962a1.694 1.694 0 0 0-.001-2.078A22.939 22.939 0 0 0 19.57 9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceHappy(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M8 9.05v-.1m8 .1v-.1"/><path stroke-linejoin="round" d="M16 14c-.5 1.5-1.79 3-4 3s-3.5-1.5-4-3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceNeutral(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M8 9.05v-.1m8 .1v-.1M8 14h8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceSad(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M8 9.05v-.1m8 .1v-.1"/><path stroke-linejoin="round" d="M16 16c-.5-1.5-1.79-3-4-3s-3.5 1.5-4 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceVeryHappy(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M8 9.05v-.1m8 .1v-.1"/><path stroke-linejoin="round" d="M12 18a4 4 0 0 0 4-4H8a4 4 0 0 0 4 4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceVerySad(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M8 9.05v-.1m8 .1v-.1"/><path stroke-linejoin="round" d="M12 13a4 4 0 0 1 4 4H8a4 4 0 0 1 4-4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceWink(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M7 9h2m7 .05v-.1"/><path stroke-linejoin="round" d="M16 15c-.5 1-1.79 2-4 2s-3.5-1-4-2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FacebookFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#akarIconsFacebookFill0)"><path fill="currentColor" fill-rule="evenodd" d="M0 12.067C0 18.034 4.333 22.994 10 24v-8.667H7V12h3V9.333c0-3 1.933-4.666 4.667-4.666c.866 0 1.8.133 2.666.266V8H15.8c-1.467 0-1.8.733-1.8 1.667V12h3.2l-.533 3.333H14V24c5.667-1.006 10-5.966 10-11.933C24 5.43 18.6 0 12 0S0 5.43 0 12.067" clip-rule="evenodd"/></g><defs><clipPath id="akarIconsFacebookFill0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FigmaFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#akarIconsFigmaFill0)"><path fill="currentColor" fill-rule="evenodd" d="M8.415 0C5.968 0 4 2.053 4 4.568c0 1.529.728 2.887 1.847 3.716A4.613 4.613 0 0 0 4 12c0 1.53.728 2.887 1.847 3.716A4.613 4.613 0 0 0 4 19.432C4 21.947 5.968 24 8.415 24c2.446 0 4.415-2.053 4.415-4.568V15.57a4.307 4.307 0 0 0 2.755.999C18.032 16.568 20 14.515 20 12a4.61 4.61 0 0 0-1.847-3.716A4.613 4.613 0 0 0 20 4.568C20 2.053 18.032 0 15.585 0zM5.659 4.568c0-1.591 1.242-2.865 2.756-2.865h2.755v5.73H8.415c-1.514 0-2.756-1.275-2.756-2.865m9.926 2.864H12.83v-5.73h2.755c1.515 0 2.756 1.275 2.756 2.866c0 1.59-1.241 2.864-2.756 2.864M5.66 12c0-1.59 1.242-2.865 2.756-2.865h2.755v5.73H8.415C6.9 14.865 5.659 13.59 5.659 12m7.17 0c0-1.59 1.242-2.865 2.756-2.865c1.515 0 2.756 1.274 2.756 2.865c0 1.59-1.241 2.865-2.756 2.865c-1.514 0-2.755-1.274-2.755-2.865m-7.17 7.432c0-1.59 1.242-2.864 2.756-2.864h2.755v2.864c0 1.591-1.24 2.865-2.755 2.865c-1.514 0-2.756-1.274-2.756-2.865" clip-rule="evenodd"/></g><defs><clipPath id="akarIconsFigmaFill0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func File(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path stroke-linecap="round" d="M4 4v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8.342a2 2 0 0 0-.602-1.43l-4.44-4.342A2 2 0 0 0 13.56 2H6a2 2 0 0 0-2 2m5 9h6m-6 4h3"/><path d="M14 2v4a2 2 0 0 0 2 2h4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Filter(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><ellipse cx="12" cy="5" rx="9" ry="2"/><path d="M3 5c0 2.23 3.871 6.674 5.856 8.805A4.197 4.197 0 0 1 10 16.657V19a2 2 0 0 0 2 2v0a2 2 0 0 0 2-2v-2.343c0-1.061.421-2.075 1.144-2.852C17.13 11.674 21 7.231 21 5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fire(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 22c-4.97 0-9-2.582-9-7v-.088C3 12.794 4.338 11.1 6.375 10c1.949-1.052 3.101-2.99 2.813-5l-.563-3l2.086.795c3.757 1.43 6.886 3.912 8.914 7.066A8.495 8.495 0 0 1 21 14.464V15c0 1.562-.504 2.895-1.375 3.965"/><path d="M12 22c-1.657 0-3-1.433-3-3.2c0-1.4 1.016-2.521 1.91-3.548L12 14l1.09 1.252C13.984 16.28 15 17.4 15 18.8c0 1.767-1.343 3.2-3 3.2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flag(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 15h13.865a1 1 0 0 0 .768-1.64L15 9l3.633-4.36A1 1 0 0 0 17.865 3H4v18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flashlight(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 4v1c0 1.636 2 4 3 5l.75 9.01A3.26 3.26 0 0 0 12 22v0a3.26 3.26 0 0 0 3.25-2.99L16 10c1-1 3-3.364 3-5V4m-7 7v2"/><ellipse cx="12" cy="4" rx="7" ry="2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folder(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 19V9a2 2 0 0 0-2-2h-6.764a2 2 0 0 1-1.789-1.106l-.894-1.788A2 2 0 0 0 8.763 3H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderAdd(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" d="M22 19V9a2 2 0 0 0-2-2h-6.764a2 2 0 0 1-1.789-1.106l-.894-1.788A2 2 0 0 0 8.763 3H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2"/><path d="M12 11v3m0 0v3m0-3h3m-3 0H9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ForkLeft(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M15 3v18m0-18l-4 4m4-4l4 4M8 10l-4 4l4 4"/><path d="M15 21a7 7 0 0 0-7-7H4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ForkRight(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M9 3v18M9 3l4 4M9 3L5 7m11 3l4 4l-4 4"/><path d="M9 21a7 7 0 0 1 7-7h4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Frame(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 3v18M18 3v18M3 6h18M3 18h18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FullScreen(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 7V2h5m15 5V2h-5M7 22H2v-5m15 5h5v-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GameController(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m9 15l-2.968 2.968A2.362 2.362 0 0 1 2 16.298V15l1.357-6.784A4 4 0 0 1 7.279 5h9.442a4 4 0 0 1 3.922 3.216L22 15v1.297a2.362 2.362 0 0 1-4.032 1.67L15 15z"/><path d="m9 5l1 2h4l1-2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GatsbyFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#akarIconsGatsbyFill0)"><path fill="currentColor" d="M12 0C5.323 0 0 5.317 0 12s5.317 12 12 12s12-5.323 12-12S18.683 0 12 0M2.608 12.101l9.29 9.286c-5.114.005-9.29-4.171-9.29-9.286m11.477 9.083L2.821 9.909C3.76 5.733 7.515 2.603 12 2.603a9.493 9.493 0 0 1 7.616 3.861l-1.355 1.147A7.666 7.666 0 0 0 11.9 4.267A7.57 7.57 0 0 0 4.693 9.38l9.814 9.819c2.4-.837 4.277-2.923 4.8-5.43h-4.07V12h6.155c0 4.485-3.13 8.245-7.307 9.184"/></g><defs><clipPath id="akarIconsGatsbyFill0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gear(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><path d="M14 3.269C14 2.568 13.432 2 12.731 2H11.27C10.568 2 10 2.568 10 3.269c0 .578-.396 1.074-.935 1.286c-.085.034-.17.07-.253.106c-.531.23-1.162.16-1.572-.249a1.269 1.269 0 0 0-1.794 0L4.412 5.446a1.269 1.269 0 0 0 0 1.794c.41.41.48 1.04.248 1.572a7.946 7.946 0 0 0-.105.253c-.212.539-.708.935-1.286.935C2.568 10 2 10.568 2 11.269v1.462C2 13.432 2.568 14 3.269 14c.578 0 1.074.396 1.286.935c.034.085.07.17.105.253c.231.531.161 1.162-.248 1.572a1.269 1.269 0 0 0 0 1.794l1.034 1.034a1.269 1.269 0 0 0 1.794 0c.41-.41 1.04-.48 1.572-.249c.083.037.168.072.253.106c.539.212.935.708.935 1.286c0 .701.568 1.269 1.269 1.269h1.462c.701 0 1.269-.568 1.269-1.269c0-.578.396-1.074.935-1.287c.085-.033.17-.068.253-.104c.531-.232 1.162-.161 1.571.248a1.269 1.269 0 0 0 1.795 0l1.034-1.034a1.269 1.269 0 0 0 0-1.794c-.41-.41-.48-1.04-.249-1.572c.037-.083.072-.168.106-.253c.212-.539.708-.935 1.286-.935c.701 0 1.269-.568 1.269-1.269V11.27c0-.701-.568-1.269-1.269-1.269c-.578 0-1.074-.396-1.287-.935a7.755 7.755 0 0 0-.105-.253c-.23-.531-.16-1.162.249-1.572a1.269 1.269 0 0 0 0-1.794l-1.034-1.034a1.269 1.269 0 0 0-1.794 0c-.41.41-1.04.48-1.572.249a7.913 7.913 0 0 0-.253-.106C14.396 4.343 14 3.847 14 3.27z"/><path d="M16 12a4 4 0 1 1-8 0a4 4 0 0 1 8 0Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gift(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="14" x="3" y="8" rx="2"/><path d="M12 5a3 3 0 1 0-3 3m6 0a3 3 0 1 0-3-3m0 0v17m9-7H3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GithubFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#akarIconsGithubFill0)"><path fill="currentColor" fill-rule="evenodd" d="M12 0C5.37 0 0 5.37 0 12c0 5.31 3.435 9.795 8.205 11.385c.6.105.825-.255.825-.57c0-.285-.015-1.23-.015-2.235c-3.015.555-3.795-.735-4.035-1.41c-.135-.345-.72-1.41-1.23-1.695c-.42-.225-1.02-.78-.015-.795c.945-.015 1.62.87 1.845 1.23c1.08 1.815 2.805 1.305 3.495.99c.105-.78.42-1.305.765-1.605c-2.67-.3-5.46-1.335-5.46-5.925c0-1.305.465-2.385 1.23-3.225c-.12-.3-.54-1.53.12-3.18c0 0 1.005-.315 3.3 1.23c.96-.27 1.98-.405 3-.405s2.04.135 3 .405c2.295-1.56 3.3-1.23 3.3-1.23c.66 1.65.24 2.88.12 3.18c.765.84 1.23 1.905 1.23 3.225c0 4.605-2.805 5.625-5.475 5.925c.435.375.81 1.095.81 2.22c0 1.605-.015 2.895-.015 3.3c0 .315.225.69.825.57A12.02 12.02 0 0 0 24 12c0-6.63-5.37-12-12-12" clip-rule="evenodd"/></g><defs><clipPath id="akarIconsGithubFill0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GithubOutlineFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1.545c-6.055 0-10.957 4.877-10.957 10.883c0 4.41 2.643 8.205 6.447 9.912c.398.179.787.108 1.091-.12a1.38 1.38 0 0 0 .532-1.107v-.57l-1.357-.184a.489.489 0 0 1-.03-.005c-.748-.146-1.253-.409-1.623-.788c-.311-.319-.501-.701-.662-1.025l-.056-.113a11.645 11.645 0 0 0-.487-.91c-.135-.214-.24-.324-.344-.393c-.264-.175-.518-.472-.518-.843a.72.72 0 0 1 .26-.56a.844.844 0 0 1 .546-.182c.222 0 .431.068.605.146c.178.08.355.186.52.296c.402.268.798.585 1.133.976c.349.406.596.664.968.717c.306.044.618.045.861.034a2.558 2.558 0 0 1 .166-.56a8.347 8.347 0 0 1-.877-.256c-.654-.232-1.383-.593-1.873-1.14c-.538-.602-.871-1.139-1.057-1.767c-.181-.612-.211-1.275-.211-2.091c0-1.25.541-2.303.947-2.862a8.655 8.655 0 0 1-.335-1.501a3.515 3.515 0 0 1 .001-.968c.05-.293.175-.653.504-.87c.317-.211.689-.202.979-.148c.301.057.614.182.902.325c.469.232.935.542 1.284.805c.602-.177 1.667-.405 2.63-.426h.022c.964.021 1.981.249 2.568.425c.35-.262.815-.573 1.283-.804c.288-.143.6-.268.902-.325c.29-.054.662-.063.98.147c.328.218.454.578.503.871c.051.305.039.646 0 .968a8.661 8.661 0 0 1-.334 1.501c.406.56.947 1.613.947 2.862c0 .816-.03 1.479-.21 2.09c-.187.63-.52 1.166-1.058 1.768c-.49.547-1.219.908-1.873 1.14a8.45 8.45 0 0 1-1.062.297c.123.372.167.673.167.846v3.09c0 .47.219.871.53 1.105c.302.229.69.3 1.088.126c3.84-1.692 6.514-5.497 6.514-9.93c0-6.005-4.9-10.882-10.956-10.882M9.476 18.71c-.26.037-.001 0-.001 0h-.003l-.008.002l-.029.004a6.523 6.523 0 0 1-.447.037a5.782 5.782 0 0 1-1.066-.043c-.791-.112-1.272-.672-1.583-1.036l-.03-.034a3.6 3.6 0 0 0-.327-.333c.103.19.21.402.325.63l.01.02l.051.104c.175.348.29.58.481.774c.186.19.476.374 1.062.49l1.794.243a.522.522 0 0 1 .452.518v1.027c0 .8-.375 1.513-.95 1.945a2.087 2.087 0 0 1-2.143.236C2.902 21.427 0 17.27 0 12.428C0 5.836 5.377.5 12 .5s12 5.336 12 11.928c0 4.867-2.939 9.035-7.137 10.886a2.09 2.09 0 0 1-2.137-.247a2.426 2.426 0 0 1-.946-1.942v-3.09c0-.07-.049-.474-.322-1.017a.524.524 0 0 1 .408-.755a7.187 7.187 0 0 0 1.504-.356c.594-.21 1.127-.498 1.444-.852c.471-.526.706-.93.835-1.367c.134-.453.168-.98.168-1.793c0-1.141-.6-2.1-.876-2.409a.523.523 0 0 1-.104-.523a7.82 7.82 0 0 0 .375-1.554c.032-.275.034-.505.006-.671a.457.457 0 0 0-.054-.173a.486.486 0 0 0-.207.009c-.17.032-.387.112-.633.234c-.49.242-1 .6-1.316.856a.52.52 0 0 1-.505.085A9.306 9.306 0 0 0 12 7.296c-1.007.024-2.169.31-2.566.453a.52.52 0 0 1-.505-.085a7.775 7.775 0 0 0-1.316-.856a2.607 2.607 0 0 0-.634-.234a.486.486 0 0 0-.206-.009a.456.456 0 0 0-.054.173c-.028.166-.026.396.006.671A7.89 7.89 0 0 0 7.1 8.963c.063.18.024.38-.104.523c-.276.309-.876 1.268-.876 2.409c0 .812.034 1.34.168 1.793c.13.437.364.84.834 1.367c.318.354.85.642 1.445.852a7.194 7.194 0 0 0 1.503.356a.522.522 0 0 1 .246.945c-.165.118-.274.33-.335.575a1.944 1.944 0 0 0-.056.385v.019a.523.523 0 0 1-.449.523" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitlabFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m22.749 9.769l-.031-.08l-3.027-7.9a.788.788 0 0 0-.782-.495a.81.81 0 0 0-.456.17a.81.81 0 0 0-.268.408L16.14 8.125H7.865L5.822 1.872a.794.794 0 0 0-.269-.409a.81.81 0 0 0-.926-.05c-.14.09-.25.22-.312.376L1.283 9.684l-.03.08a5.62 5.62 0 0 0 1.864 6.496l.01.008l.028.02l4.61 3.453l2.282 1.726l1.39 1.049a.935.935 0 0 0 1.13 0l1.389-1.05l2.281-1.726l4.639-3.473l.011-.01A5.623 5.623 0 0 0 22.75 9.77"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Glasses(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="6" cy="14" r="4"/><circle cx="18" cy="14" r="4"/><path d="m10 14l.211-.106a4 4 0 0 1 3.578 0L14 14m5-8l2.838 6.623a2 2 0 0 1 .162.788V14M5 6l-2.838 6.623A2 2 0 0 0 2 13.41V14"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Globe(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><ellipse cx="12" cy="12" rx="10" ry="4" transform="rotate(90 12 12)"/><path d="M2 12h20"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleContainedFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#akarIconsGoogleContainedFill0)"><path fill="currentColor" fill-rule="evenodd" d="M12 0C5.372 0 0 5.373 0 12s5.372 12 12 12c6.627 0 12-5.373 12-12S18.627 0 12 0m.14 19.018c-3.868 0-7-3.14-7-7.018c0-3.878 3.132-7.018 7-7.018c1.89 0 3.47.697 4.682 1.829l-1.974 1.978v-.004c-.735-.702-1.667-1.062-2.708-1.062c-2.31 0-4.187 1.956-4.187 4.273c0 2.315 1.877 4.277 4.187 4.277c2.096 0 3.522-1.202 3.816-2.852H12.14v-2.737h6.585c.088.47.135.96.135 1.474c0 4.01-2.677 6.86-6.72 6.86" clip-rule="evenodd"/></g><defs><clipPath id="akarIconsGoogleContainedFill0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.456 10.154c.123.659.19 1.348.19 2.067c0 5.624-3.764 9.623-9.449 9.623A9.841 9.841 0 0 1 2.353 12a9.841 9.841 0 0 1 9.844-9.844c2.658 0 4.879.978 6.583 2.566l-2.775 2.775V7.49c-1.033-.984-2.344-1.489-3.808-1.489c-3.248 0-5.888 2.744-5.888 5.993c0 3.248 2.64 5.999 5.888 5.999c2.947 0 4.953-1.686 5.365-4h-5.365v-3.839z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GraphqlFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#akarIconsGraphqlFill0)"><path fill="currentColor" fill-rule="evenodd" d="M11.384 4.187a2.132 2.132 0 0 0 1.171.018l6.326 11.182a2.126 2.126 0 0 0-.593 1.038H5.71a2.148 2.148 0 0 0-.64-1.079zm-.921-.565c.018.02.038.039.057.058L4.193 14.865a2.11 2.11 0 0 0-.09-.022V9.156A2.122 2.122 0 0 0 5.481 8.14a2.142 2.142 0 0 0 .19-1.694zm3.57-.835a2.134 2.134 0 1 0-4.079-.04L5.159 5.573A2.131 2.131 0 0 0 1.785 6a2.146 2.146 0 0 0 1.322 3.143v5.715A2.147 2.147 0 0 0 1.79 18a2.13 2.13 0 0 0 3.368.43l4.795 2.826a2.151 2.151 0 0 0-.086.605A2.134 2.134 0 0 0 12 24a2.138 2.138 0 0 0 2.012-2.848l4.751-2.8a2.128 2.128 0 0 0 3.44-.352a2.138 2.138 0 0 0-1.384-3.159v-5.68A2.147 2.147 0 0 0 22.215 6a2.133 2.133 0 0 0-3.401-.398zm-.599.937l.075-.072l4.808 2.833a2.142 2.142 0 0 0 .203 1.654a2.12 2.12 0 0 0 1.303 1v5.72a2.074 2.074 0 0 0-.078.023zm4.895 13.858l-4.774 2.814a2.124 2.124 0 0 0-1.554-.674c-.606 0-1.152.252-1.54.658l-4.785-2.82c.012-.039.023-.078.032-.118H18.29c.012.047.025.094.04.14" clip-rule="evenodd"/></g><defs><clipPath id="akarIconsGraphqlFill0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Grid(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="7" height="7" x="3" y="3" rx="1"/><rect width="7" height="7" x="3" y="14" rx="1"/><rect width="7" height="7" x="14" y="3" rx="1"/><rect width="7" height="7" x="14" y="14" rx="1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hammer(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 22l-2-2l1-1l2 2zm5-9l2 2m1-3l-8 8M20 4l-1 1M9.707 7.707a1 1 0 0 1 0-1.414l4.086-4.086a1 1 0 0 1 1.414 0l6.586 6.586a1 1 0 0 1 0 1.414l-4.086 4.086a1 1 0 0 1-1.414 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hand(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 16V8.5c0-.828-.641-1.5-1.48-1.5C18 7 17 7.3 17 8.5v-3c0-.828-.641-1.5-1.48-1.5c-.507 0-1.52.3-1.52 1.5v-2c0-.828-.641-1.5-1.48-1.5c-.84 0-1.52.672-1.52 1.5v2C11 4.3 10.007 4 9.5 4C8.66 4 8 4.691 8 5.52V14m3-8.5V11m3-5.5V11m3-5.5V11"/><path d="M20 16c0 4-3.134 6-7 6s-5.196-1-8.196-6l-1.571-2.605c-.536-.868-.107-1.994.881-2.314a1.657 1.657 0 0 1 1.818.552L8 14.033"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hashtag(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 3L6 21M18 3l-4 18M4 8h17M3 16h17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Headphone(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 13.565C2 11.512 4 11 6 11v9a4 4 0 0 1-4-4zm20 0C22 11.512 20 11 18 11v9a4 4 0 0 0 4-4zM6 20V10a6 6 0 1 1 12 0v10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Health(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 3a1 1 0 0 0-1-1H9a1 1 0 0 0-1 1v5H3a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h5v5a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1v-5h5a1 1 0 0 0 1-1V9a1 1 0 0 0-1-1h-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heart(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 3C4.239 3 2 5.216 2 7.95c0 2.207.875 7.445 9.488 12.74a.985.985 0 0 0 1.024 0C21.126 15.395 22 10.157 22 7.95C22 5.216 19.761 3 17 3s-5 3-5 3s-2.239-3-5-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Height(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 22V2m0 20l-4-4m4 4l4-4M12 2L8 6m4-4l4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heptagon(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="2" d="M11.114 2.438a2 2 0 0 1 1.772 0l6.275 3.1a2 2 0 0 1 1.066 1.358l1.569 7.047a2 2 0 0 1-.374 1.662l-4.371 5.623a2 2 0 0 1-1.579.772H8.528a2 2 0 0 1-1.579-.772l-4.371-5.623a2 2 0 0 1-.374-1.662l1.569-7.047a2 2 0 0 1 1.066-1.359z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeptagonFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.114 1.438a2 2 0 0 1 1.772 0l7.077 3.495a2 2 0 0 1 1.066 1.359l1.767 7.937a2 2 0 0 1-.374 1.662l-4.926 6.337a2 2 0 0 1-1.579.772H8.083a2 2 0 0 1-1.579-.772L1.578 15.89a2 2 0 0 1-.374-1.662l1.767-7.937a2 2 0 0 1 1.066-1.359z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hexagon(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="2" d="M11.029 2.54a2 2 0 0 1 1.942 0l7 3.888A2 2 0 0 1 21 8.177v7.646a2 2 0 0 1-1.029 1.749l-7 3.888a2 2 0 0 1-1.942 0l-7-3.889A2 2 0 0 1 3 15.824V8.177a2 2 0 0 1 1.029-1.749z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HexagonFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.036 1.53a2 2 0 0 1 1.928 0l8 4.4A2 2 0 0 1 22 7.683v8.635a2 2 0 0 1-1.036 1.752l-8 4.4a2 2 0 0 1-1.928 0l-8-4.4A2 2 0 0 1 2 16.317V7.683A2 2 0 0 1 3.036 5.93z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func History(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.266 16.06a8.923 8.923 0 0 0 3.915 3.978a8.706 8.706 0 0 0 5.471.832a8.795 8.795 0 0 0 4.887-2.64a9.067 9.067 0 0 0 2.388-5.079a9.135 9.135 0 0 0-1.044-5.53a8.903 8.903 0 0 0-4.069-3.815a8.7 8.7 0 0 0-5.5-.608c-1.85.401-3.366 1.313-4.62 2.755c-.151.16-.735.806-1.22 1.781M7.5 8l-3.609.72L3 5m9 4v4l3 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Home(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 19v-6.733a4 4 0 0 0-1.245-2.9L13.378 3.31a2 2 0 0 0-2.755 0L4.245 9.367A4 4 0 0 0 3 12.267V19a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeAltOne(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 19v-6.733a4 4 0 0 0-1.245-2.9L13.378 3.31a2 2 0 0 0-2.755 0L4.245 9.367A4 4 0 0 0 3 12.267V19a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2"/><path d="M9 15a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v6H9z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HtmlFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g fill="currentColor" clip-path="url(#akarIconsHtmlFill0)"><path d="M5.08 0h1.082v1.069h.99V0h1.082v3.236H7.152V2.153h-.99v1.083H5.081zm4.576 1.073h-.952V0h2.987v1.073h-.953v2.163H9.656zM12.165 0h1.128l.694 1.137L14.68 0h1.128v3.236h-1.077V1.632l-.744 1.151h-.019l-.745-1.15v1.603h-1.058zm4.181 0h1.083v2.167h1.52v1.07h-2.603z"/><path fill-rule="evenodd" d="M5.046 22.072L3 4.717h18L18.953 22.07L11.99 24zm4.137-9.5l-.194-2.18h8.145l.19-2.128H6.664l.574 6.437h7.377l-.247 2.76l-2.374.642h-.002l-2.37-.64l-.152-1.697H7.332l.298 3.342l4.36 1.21l4.367-1.21l.532-5.964l.052-.571z" clip-rule="evenodd"/></g><defs><clipPath id="akarIconsHtmlFill0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Image(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 6a4 4 0 0 1 4-4h12a4 4 0 0 1 4 4v12a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4z"/><circle cx="8.5" cy="8.5" r="2.5"/><path d="M14.526 12.621L6 22h12.133A3.867 3.867 0 0 0 22 18.133V18c0-.466-.175-.645-.49-.99l-4.03-4.395a2 2 0 0 0-2.954.006"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Inbox(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path stroke-linecap="round" d="M3.304 6.132A4 4 0 0 1 7.209 3h9.582a4 4 0 0 1 3.905 3.132l.147.662a23.997 23.997 0 0 1 0 10.412l-.147.662A4 4 0 0 1 16.791 21H7.21a4 4 0 0 1-3.905-3.132l-.147-.662a24 24 0 0 1 0-10.412z"/><path d="M2.5 13h6.338c0 1 .973 3 3.405 3c2.433 0 3.406-2 3.406-3H21.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Infinite(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5.636 16C2.91 16 2 14 2 12s.91-4 3.636-4c3.637 0 9.091 8 12.728 8C21.09 16 22 14 22 12s-.91-4-3.636-4c-3.637 0-9.091 8-12.728 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Infinity(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5.636 16C2.91 16 2 14 2 12s.91-4 3.636-4c3.637 0 9.091 8 12.728 8C21.09 16 22 14 22 12s-.91-4-3.636-4c-3.637 0-9.091 8-12.728 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Info(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><path stroke-linecap="round" d="M12 7h.01"/><path stroke-linecap="round" stroke-linejoin="round" d="M10 11h2v5m-2 0h4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11s11-4.925 11-11S18.075 1 12 1m-.5 5a1 1 0 1 0 0 2h.5a1 1 0 1 0 0-2zM10 10a1 1 0 1 0 0 2h1v3h-1a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2h-1v-4a1 1 0 0 0-1-1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InstagramFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.465 1.066C8.638 1.012 9.012 1 12 1c2.988 0 3.362.013 4.534.066c1.172.053 1.972.24 2.672.511c.733.277 1.398.71 1.948 1.27c.56.549.992 1.213 1.268 1.947c.272.7.458 1.5.512 2.67C22.988 8.639 23 9.013 23 12c0 2.988-.013 3.362-.066 4.535c-.053 1.17-.24 1.97-.512 2.67a5.396 5.396 0 0 1-1.268 1.949c-.55.56-1.215.992-1.948 1.268c-.7.272-1.5.458-2.67.512c-1.174.054-1.548.066-4.536.066c-2.988 0-3.362-.013-4.535-.066c-1.17-.053-1.97-.24-2.67-.512a5.397 5.397 0 0 1-1.949-1.268a5.392 5.392 0 0 1-1.269-1.948c-.271-.7-.457-1.5-.511-2.67C1.012 15.361 1 14.987 1 12c0-2.988.013-3.362.066-4.534c.053-1.172.24-1.972.511-2.672a5.396 5.396 0 0 1 1.27-1.948a5.392 5.392 0 0 1 1.947-1.269c.7-.271 1.5-.457 2.67-.511m8.98 1.98c-1.16-.053-1.508-.064-4.445-.064c-2.937 0-3.285.011-4.445.064c-1.073.049-1.655.228-2.043.379c-.513.2-.88.437-1.265.822a3.412 3.412 0 0 0-.822 1.265c-.151.388-.33.97-.379 2.043c-.053 1.16-.064 1.508-.064 4.445c0 2.937.011 3.285.064 4.445c.049 1.073.228 1.655.379 2.043c.176.477.457.91.822 1.265c.355.365.788.646 1.265.822c.388.151.97.33 2.043.379c1.16.053 1.507.064 4.445.064c2.938 0 3.285-.011 4.445-.064c1.073-.049 1.655-.228 2.043-.379c.513-.2.88-.437 1.265-.822c.365-.355.646-.788.822-1.265c.151-.388.33-.97.379-2.043c.053-1.16.064-1.508.064-4.445c0-2.937-.011-3.285-.064-4.445c-.049-1.073-.228-1.655-.379-2.043c-.2-.513-.437-.88-.822-1.265a3.413 3.413 0 0 0-1.265-.822c-.388-.151-.97-.33-2.043-.379m-5.85 12.345a3.669 3.669 0 0 0 4-5.986a3.67 3.67 0 1 0-4 5.986M8.002 8.002a5.654 5.654 0 1 1 7.996 7.996a5.654 5.654 0 0 1-7.996-7.996m10.906-.814a1.337 1.337 0 1 0-1.89-1.89a1.337 1.337 0 0 0 1.89 1.89" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Jar(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m3 9l1.164 8.926a3.875 3.875 0 0 0 2.844 3.243v0c3.27.872 6.713.872 9.984 0v0a3.875 3.875 0 0 0 2.844-3.243L21 9"/><path d="M5.035 7.266C3.763 7.661 3 8.165 3 8.714C3 9.977 7.03 11 12 11s9-1.023 9-2.286c0-.55-.764-1.054-2.037-1.448"/><path d="m9 4l-3 .51C4.159 4.874 3 5.407 3 6c0 1.105 4.03 2 9 2s9-.895 9-2c0-.592-1.159-1.125-3-1.49L15 4m0 0a2 2 0 0 0-2-2h-2a2 2 0 0 0-2 2v0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func JavascriptFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#akarIconsJavascriptFill0)"><path fill="currentColor" fill-rule="evenodd" d="M0 0h24v24H0zm18.347 20.12c-1.113 0-1.742-.58-2.225-1.37l-1.833 1.065c.662 1.308 2.015 2.306 4.11 2.306c2.142 0 3.737-1.112 3.737-3.143c0-1.883-1.082-2.72-2.998-3.543l-.564-.241c-.968-.42-1.387-.693-1.387-1.37c0-.547.42-.966 1.08-.966c.647 0 1.064.273 1.451.966l1.756-1.127c-.743-1.307-1.773-1.806-3.207-1.806c-2.014 0-3.303 1.288-3.303 2.98c0 1.835 1.08 2.704 2.708 3.397l.564.242c1.029.45 1.642.724 1.642 1.497c0 .646-.597 1.113-1.531 1.113m-8.74-.015c-.775 0-1.098-.53-1.452-1.16l-1.836 1.112c.532 1.126 1.578 2.06 3.383 2.06c1.999 0 3.368-1.063 3.368-3.398v-7.7h-2.255v7.67c0 1.127-.468 1.416-1.209 1.416" clip-rule="evenodd"/></g><defs><clipPath id="akarIconsJavascriptFill0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func JqueryFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g fill="currentColor" clip-path="url(#akarIconsJqueryFill0)"><path d="M13.301 5.601c-.628-1.528-.54-3.247.367-4.63L14.346 0c-1.254 2.35-.06 5.236 1.86 6.749a6.343 6.343 0 0 0 1.698.95c.164.06.35.124.52.167c2.146.54 3.895.167 4.936-1.389c-.032.048-.073.124-.125.22c-.464.86-1.822 3.377-5.76 2.587c-.075-.015-.147-.04-.219-.063l-.064-.022l-.046-.015a2.119 2.119 0 0 0-.064-.021a5.8 5.8 0 0 1-.174-.063c-1.585-.604-2.944-1.876-3.607-3.499"/><path d="M8.12 2.966c-1.537 2.295-1.453 5.368-.254 7.796a10.242 10.242 0 0 0 .762 1.301c.209.313.439.66.708.902c.111.127.227.25.345.371l.068.07l.023.023a10.22 10.22 0 0 0 .367.353l.001.001l.002.002a10.623 10.623 0 0 0 .465.397l.04.032c.14.112.283.222.43.327l.006.004l.007.005c.047.034.095.067.143.099l.053.036a3.109 3.109 0 0 1 .093.064c.105.07.211.137.319.203l.03.017l.015.01a11.466 11.466 0 0 0 .351.201l.031.016l.04.022c.052.028.104.057.158.084l.03.015a9.941 9.941 0 0 0 .47.224l.032.013a10.315 10.315 0 0 0 .451.186l.03.011c.103.04.206.076.31.112l.052.018l.088.029c.039.012.078.026.116.039c.109.038.218.075.332.095C22.168 17.408 24 11.068 24 11.068c-1.651 2.468-4.849 3.646-8.261 2.726a9.54 9.54 0 0 1-.45-.135a2.811 2.811 0 0 1-.12-.04l-.015-.005a8.82 8.82 0 0 1-.314-.112l-.003-.002a10.08 10.08 0 0 1-.512-.21l-.056-.025a9.272 9.272 0 0 1-.58-.287l-.042-.022l-.12-.064a9.64 9.64 0 0 1-.325-.19a10.318 10.318 0 0 1-.62-.41a10.976 10.976 0 0 1-.429-.326l-.035-.029a6.304 6.304 0 0 1-.061-.05c-1.494-1.224-2.678-2.897-3.24-4.793c-.59-1.968-.463-4.176.559-5.968z"/><path d="M1.524 5.637C-.6 8.807-.336 12.932 1.287 16.3c.03.065.063.129.095.192l.024.048l.03.059a2.987 2.987 0 0 0 .08.155l.013.025l.07.13l.015.027c.05.092.102.184.155.276l.017.03a9.464 9.464 0 0 0 .251.413l.018.03l.059.093c.084.13.17.261.26.39v.001a.085.085 0 0 0 .007.009l.024.033l.02.028c.078.11.157.22.238.328l.09.118a18.083 18.083 0 0 0 .608.746l.006.007l.005.005l.007.008a15.627 15.627 0 0 0 .658.697l.03.03l.071.07a12.947 12.947 0 0 0 .346.326l.033.03l.026.022a18.342 18.342 0 0 0 .38.33l.056.047c.085.07.171.14.258.208l.137.109c.096.073.192.145.289.215l.016.012l.09.066l.028.02c.088.064.177.124.267.184l.014.01a5.661 5.661 0 0 1 .12.083a15.577 15.577 0 0 0 .507.319l.054.032a14.516 14.516 0 0 0 .459.26l.053.029l.059.032c.059.032.117.064.177.095l.025.012l.03.015l.034.017c.02.01.041.02.062.032c.124.061.248.121.374.18l.047.021l.032.015a13.238 13.238 0 0 0 .953.39l.01.003l.042.016a13.91 13.91 0 0 0 .69.228c.116.039.233.077.353.1c10.277 1.946 13.262-6.41 13.262-6.41c-2.507 3.39-6.957 4.285-11.174 3.289a5.368 5.368 0 0 1-.47-.137l-.118-.037c-.152-.048-.303-.1-.453-.153l-.062-.023c-.135-.05-.267-.1-.4-.154l-.112-.046a12.71 12.71 0 0 1-.431-.186l-.046-.02a17.296 17.296 0 0 1-.404-.194l-.041-.022a5.56 5.56 0 0 1-.064-.031l-.214-.113l-.068-.036a6.057 6.057 0 0 0-.067-.036l-.12-.066c-.114-.063-.226-.13-.338-.197l-.043-.025l-.07-.04a16.203 16.203 0 0 1-.516-.332l-.043-.03a12.657 12.657 0 0 1-1.225-.929a15.269 15.269 0 0 1-.352-.307l-.016-.015a16.412 16.412 0 0 1-.448-.428a11.175 11.175 0 0 1-.25-.255l-.034-.034a13.54 13.54 0 0 1-.38-.414l-.006-.006l-.002-.003c-.106-.12-.21-.244-.312-.368l-.082-.101a13.975 13.975 0 0 1-.307-.395c-.088-.12-.175-.238-.26-.359c-2.34-3.314-3.181-7.886-1.31-11.64z"/></g><defs><clipPath id="akarIconsJqueryFill0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Key(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m8 18l2-2h2l1.36-1.36a6.5 6.5 0 1 0-3.997-3.992L2 18v4h4l2-2z"/><circle cx="17" cy="7" r="1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyCap(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 3s3 .5 5 .5s5-.5 5-.5l1 9s-3 1-6 1s-6-1-6-1z"/><path d="m3.869 8.147l-.862 10.294c-.03.366.04.733.205 1.06l.197.393A2.01 2.01 0 0 0 5.206 21h13.588a2.01 2.01 0 0 0 1.797-1.106l.197-.392c.165-.328.236-.695.205-1.06l-.862-10.295a3.99 3.99 0 0 0-.79-2.068L17 3s-3 .5-5 .5S6.978 3 6.978 3l-2.32 3.08a3.99 3.99 0 0 0-.79 2.067M6 12l-2.5 8M18 12l2.5 8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Language(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" d="M14 19c3.771 0 5.657 0 6.828-1.172C22 16.657 22 14.771 22 11c0-3.771 0-5.657-1.172-6.828C19.657 3 17.771 3 14 3h-4C6.229 3 4.343 3 3.172 4.172C2 5.343 2 7.229 2 11c0 3.771 0 5.657 1.172 6.828c.653.654 1.528.943 2.828 1.07"/><path d="M14 19c-1.236 0-2.598.5-3.841 1.145c-1.998 1.037-2.997 1.556-3.489 1.225c-.492-.33-.399-1.355-.212-3.404L6.5 17.5"/><path stroke-linejoin="round" d="m5.5 13.5l1-2m0 0l1.106-2.211a1 1 0 0 1 1.788 0L10.5 11.5m-4 0h4m0 0l1 2m1-6h1.982V9c0 .5-.496 1.5-1.487 1.5m3.964-3v2m0 0v4m0-4H18.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LaptopDevice(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 7a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v9H4zM2 19h20"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Leaf(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 18C19.955 18 20.917 7.83 20.994 2.997a.983.983 0 0 0-1.006-.988C3 2.321 3 10.557 3 18v4"/><path d="M3 18s0-6 8-7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lifesaver(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g stroke="currentColor" stroke-linejoin="round" stroke-width="2" clip-path="url(#akarIconsLifesaver0)"><circle cx="12" cy="12" r="10" stroke-linecap="round" transform="rotate(45 12 12)"/><circle cx="12" cy="12" r="4" stroke-linecap="round" transform="rotate(45 12 12)"/><path d="m19.071 4.929l-4.243 4.243m-5.656 5.656l-4.243 4.243m14.142 0l-4.243-4.243M9.172 9.172L4.929 4.929"/></g><defs><clipPath id="akarIconsLifesaver0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightBulb(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 22h4M5 9a7 7 0 0 1 14 0a6.972 6.972 0 0 1-3 5.734l-.542 2.566a2 2 0 0 1-1.977 1.7h-2.962a2 2 0 0 1-1.977-1.7L8 14.745A6.992 6.992 0 0 1 5 9m3 6h8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkChain(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M13.544 10.456a4.368 4.368 0 0 0-6.176 0l-3.089 3.088a4.367 4.367 0 1 0 6.177 6.177L12 18.177"/><path d="M10.456 13.544a4.368 4.368 0 0 0 6.176 0l3.089-3.088a4.367 4.367 0 1 0-6.177-6.177L12 5.823"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkOff(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 6h1a6 6 0 0 1 0 12h-1m-6 0H8A6 6 0 0 1 8 6h1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkOn(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m0-6h1a6 6 0 0 1 0 12h-1m-6 0H8A6 6 0 0 1 8 6h1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkOut(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.5 10.5L21 3m-5 0h5v5m0 6v5a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkedinBoxFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 2.838A1.838 1.838 0 0 1 2.838 1H21.16A1.837 1.837 0 0 1 23 2.838V21.16A1.838 1.838 0 0 1 21.161 23H2.838A1.838 1.838 0 0 1 1 21.161zm8.708 6.55h2.979v1.496c.43-.86 1.53-1.634 3.183-1.634c3.169 0 3.92 1.713 3.92 4.856v5.822h-3.207v-5.106c0-1.79-.43-2.8-1.522-2.8c-1.515 0-2.145 1.089-2.145 2.8v5.106H9.708zm-5.5 10.403h3.208V9.25H4.208zM7.875 5.812a2.063 2.063 0 1 1-4.125 0a2.063 2.063 0 0 1 4.125 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkedinFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.429 8.969h3.714v1.85c.535-1.064 1.907-2.02 3.968-2.02c3.951 0 4.889 2.118 4.889 6.004V22h-4v-6.312c0-2.213-.535-3.461-1.897-3.461c-1.889 0-2.674 1.345-2.674 3.46V22h-4zM2.57 21.83h4V8.799h-4zM7.143 4.55a2.53 2.53 0 0 1-.753 1.802a2.573 2.573 0 0 1-1.82.748a2.59 2.59 0 0 1-1.818-.747A2.548 2.548 0 0 1 2 4.55c0-.677.27-1.325.753-1.803A2.583 2.583 0 0 1 4.571 2c.682 0 1.336.269 1.819.747c.482.478.753 1.126.753 1.803" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkedinVoneFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 2.838A1.838 1.838 0 0 1 2.838 1H21.16A1.837 1.837 0 0 1 23 2.838V21.16A1.838 1.838 0 0 1 21.161 23H2.838A1.838 1.838 0 0 1 1 21.161zm8.708 6.55h2.979v1.496c.43-.86 1.53-1.634 3.183-1.634c3.169 0 3.92 1.713 3.92 4.856v5.822h-3.207v-5.106c0-1.79-.43-2.8-1.522-2.8c-1.515 0-2.145 1.089-2.145 2.8v5.106H9.708zm-5.5 10.403h3.208V9.25H4.208zM7.875 5.812a2.063 2.063 0 1 1-4.125 0a2.063 2.063 0 0 1 4.125 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkedinVtwoFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.429 8.969h3.714v1.85c.535-1.064 1.907-2.02 3.968-2.02c3.951 0 4.889 2.118 4.889 6.004V22h-4v-6.312c0-2.213-.535-3.461-1.897-3.461c-1.889 0-2.674 1.345-2.674 3.46V22h-4zM2.57 21.83h4V8.799h-4zM7.143 4.55a2.53 2.53 0 0 1-.753 1.802a2.573 2.573 0 0 1-1.82.748a2.59 2.59 0 0 1-1.818-.747A2.548 2.548 0 0 1 2 4.55c0-.677.27-1.325.753-1.803A2.583 2.583 0 0 1 4.571 2c.682 0 1.336.269 1.819.747c.482.478.753 1.126.753 1.803" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkedinvOneFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 2.838A1.838 1.838 0 0 1 2.838 1H21.16A1.837 1.837 0 0 1 23 2.838V21.16A1.838 1.838 0 0 1 21.161 23H2.838A1.838 1.838 0 0 1 1 21.161zm8.708 6.55h2.979v1.496c.43-.86 1.53-1.634 3.183-1.634c3.169 0 3.92 1.713 3.92 4.856v5.822h-3.207v-5.106c0-1.79-.43-2.8-1.522-2.8c-1.515 0-2.145 1.089-2.145 2.8v5.106H9.708zm-5.5 10.403h3.208V9.25H4.208zM7.875 5.812a2.063 2.063 0 1 1-4.125 0a2.063 2.063 0 0 1 4.125 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkedinvTwoFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.429 8.969h3.714v1.85c.535-1.064 1.907-2.02 3.968-2.02c3.951 0 4.889 2.118 4.889 6.004V22h-4v-6.312c0-2.213-.535-3.461-1.897-3.461c-1.889 0-2.674 1.345-2.674 3.46V22h-4zM2.57 21.83h4V8.799h-4zM7.143 4.55a2.53 2.53 0 0 1-.753 1.802a2.573 2.573 0 0 1-1.82.748a2.59 2.59 0 0 1-1.818-.747A2.548 2.548 0 0 1 2 4.55c0-.677.27-1.325.753-1.803A2.583 2.583 0 0 1 4.571 2c.682 0 1.336.269 1.819.747c.482.478.753 1.126.753 1.803" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Location(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="10" r="3"/><path d="M12 2a8 8 0 0 0-8 8c0 1.892.402 3.13 1.5 4.5L12 22l6.5-7.5c1.098-1.37 1.5-2.608 1.5-4.5a8 8 0 0 0-8-8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockOff(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="12" x="3" y="10" rx="2"/><path d="M6 10V5a3 3 0 0 1 3-3h6a3 3 0 0 1 3 3v1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockOn(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="12" x="3" y="10" rx="2"/><path d="M6 6a3 3 0 0 1 3-3h6a3 3 0 0 1 3 3v4H6z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Map(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path stroke-linecap="round" d="m8.368 4.79l-2.736-.913A2 2 0 0 0 3 5.775v11.783a2 2 0 0 0 1.368 1.898l4 1.333a2 2 0 0 0 1.264 0l4.736-1.578a2 2 0 0 1 1.265 0l2.735.912A2 2 0 0 0 21 18.225V6.442a2 2 0 0 0-1.367-1.898l-4-1.333a2 2 0 0 0-1.265 0L9.631 4.789a2 2 0 0 1-1.264 0"/><path d="M9 5v16m6-18v16"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MastodonFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.158 0c-3.068.025-6.02.357-7.74 1.147c0 0-3.41 1.526-3.41 6.733c0 1.192-.024 2.617.014 4.129c.123 5.092.933 10.11 5.64 11.355c2.171.575 4.035.695 5.535.613c2.722-.151 4.25-.972 4.25-.972l-.09-1.974s-1.945.613-4.13.538c-2.163-.074-4.448-.233-4.798-2.89a5.448 5.448 0 0 1-.048-.745s2.124.519 4.816.642c1.647.076 3.19-.096 4.759-.283c3.007-.36 5.625-2.212 5.954-3.905c.519-2.667.476-6.508.476-6.508c0-5.207-3.411-6.733-3.411-6.733C18.255.357 15.302.025 12.233 0zM8.686 4.068c1.278 0 2.245.491 2.885 1.474l.622 1.043l.623-1.043c.64-.983 1.607-1.474 2.885-1.474c1.105 0 1.995.388 2.675 1.146c.658.757.986 1.781.986 3.07v6.303h-2.497V8.47c0-1.29-.543-1.945-1.628-1.945c-1.2 0-1.802.777-1.802 2.313v3.349h-2.483v-3.35c0-1.535-.601-2.312-1.802-2.312c-1.085 0-1.628.655-1.628 1.945v6.118H5.024V8.283c0-1.288.328-2.312.987-3.07c.68-.757 1.57-1.145 2.675-1.145"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediumFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.372 7.264a.784.784 0 0 0-.252-.658L2.252 4.339V4H8.05l4.482 9.905L16.472 4H22v.339L20.403 5.88a.472.472 0 0 0-.177.452v11.334a.472.472 0 0 0 .177.452l1.56 1.542V20h-7.844v-.339l1.616-1.58c.159-.16.159-.207.159-.451V8.468l-4.492 11.494h-.606L5.566 8.468v7.704c-.043.323.064.65.29.884l2.101 2.568v.338H2v-.338l2.1-2.568a1.03 1.03 0 0 0 .272-.884z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mention(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="4"/><path stroke-linecap="round" stroke-linejoin="round" d="M22 12c0-5.523-4.477-10-10-10S2 6.477 2 12s4.477 10 10 10c2.252 0 4.33-.744 6.001-2"/><path stroke-linecap="round" d="M16 8v4c0 1 .6 3 3 3s3-2 3-3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Microphone(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="8" height="13" x="8" y="2" rx="4"/><path d="M18 16.292A7.98 7.98 0 0 1 12 19a7.98 7.98 0 0 1-6-2.708M12 19v3m-2 0h4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Miniplayer(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="16" x="2" y="4" rx="2"/><rect width="9" height="7" x="13" y="13" rx="2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Minus(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M20 12H4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MobileDevice(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="12" height="20" x="6" y="2" rx="2"/><path d="M11.95 18h.1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Money(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 7a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2z"/><circle cx="12" cy="12" r="3"/><path d="M2 9a4 4 0 0 0 4-4v0m12 14a4 4 0 0 1 4-4v0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Moon(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><path d="M20.958 15.325c.204-.486-.379-.9-.868-.684a7.684 7.684 0 0 1-3.101.648c-4.185 0-7.577-3.324-7.577-7.425a7.28 7.28 0 0 1 1.134-3.91c.284-.448-.057-1.068-.577-.936C5.96 4.041 3 7.613 3 11.862C3 16.909 7.175 21 12.326 21c3.9 0 7.24-2.345 8.632-5.675Z"/><path d="M15.611 3.103c-.53-.354-1.162.278-.809.808l.63.945a2.332 2.332 0 0 1 0 2.588l-.63.945c-.353.53.28 1.162.81.808l.944-.63a2.332 2.332 0 0 1 2.588 0l.945.63c.53.354 1.162-.278.808-.808l-.63-.945a2.332 2.332 0 0 1 0-2.588l.63-.945c.354-.53-.278-1.162-.809-.808l-.944.63a2.332 2.332 0 0 1-2.588 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.958 15.325c.204-.486-.379-.9-.868-.684a7.684 7.684 0 0 1-3.101.648c-4.185 0-7.577-3.324-7.577-7.425a7.28 7.28 0 0 1 1.134-3.91c.284-.448-.057-1.068-.577-.936C5.96 4.041 3 7.613 3 11.862C3 16.909 7.175 21 12.326 21c3.9 0 7.24-2.345 8.632-5.675"/><path fill="currentColor" d="M15.611 3.103c-.53-.354-1.162.278-.809.808l.63.945a2.332 2.332 0 0 1 0 2.588l-.63.945c-.353.53.28 1.162.81.808l.944-.63a2.332 2.332 0 0 1 2.588 0l.945.63c.53.354 1.162-.278.808-.808l-.63-.945a2.332 2.332 0 0 1 0-2.588l.63-.945c.354-.53-.278-1.162-.809-.808l-.944.63a2.332 2.332 0 0 1-2.588 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoreHorizontal(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><circle cx="4" cy="12" r="1"/><circle cx="12" cy="12" r="1"/><circle cx="20" cy="12" r="1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoreHorizontalFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 12a2 2 0 1 1 4 0a2 2 0 0 1-4 0m8 0a2 2 0 1 1 4 0a2 2 0 0 1-4 0m8 0a2 2 0 1 1 4 0a2 2 0 0 1-4 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoreVertical(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="4" r="1" transform="rotate(90 12 4)"/><circle cx="12" cy="12" r="1" transform="rotate(90 12 12)"/><circle cx="12" cy="20" r="1" transform="rotate(90 12 20)"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoreVerticalFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 2a2 2 0 1 1 0 4a2 2 0 0 1 0-4m0 8a2 2 0 1 1 0 4a2 2 0 0 1 0-4m0 8a2 2 0 1 1 0 4a2 2 0 0 1 0-4" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Music(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M8 18V5.716a2 2 0 0 1 1.696-1.977l9-1.385A2 2 0 0 1 21 4.331V16"/><path d="m8 9l13-2"/><path stroke-linecap="round" stroke-linejoin="round" d="M8 18a3 3 0 1 1-6 0c0-1.657 1.343-2 3-2s3 .343 3 2m13-2a3 3 0 1 1-6 0c0-1.657 1.343-2 3-2s3 .343 3 2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicAlbum(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M2 6a4 4 0 0 1 4-4h12a4 4 0 0 1 4 4v12a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4z"/><path stroke-linecap="round" stroke-linejoin="round" d="M9 16.5V8.78a1 1 0 0 1 .757-.97l6-1.5A1 1 0 0 1 17 7.28V15"/><path d="m9 11l8-2"/><circle cx="7.5" cy="16.5" r="1.5"/><circle cx="15.5" cy="15.5" r="1.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicAlbumFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 5a4 4 0 0 1 4.002-4h13.996A4 4 0 0 1 23 5v14a4 4 0 0 1-4 4H5a4 4 0 0 1-4-4zm19 7a8 8 0 1 1-16 0a8 8 0 0 1 16 0m-8 2a2 2 0 1 0 0-4a2 2 0 0 0 0 4" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicNote(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M12 4v14"/><path d="M19 7.674v-.657a4 4 0 0 0-2.901-3.846l-2.824-.807A1 1 0 0 0 12 3.326V7l5.725 1.636A1 1 0 0 0 19 7.674Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 18a3 3 0 1 1-6 0c0-1.657 1.343-2 3-2s3 .343 3 2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Network(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><circle cx="18" cy="5" r="3"/><circle cx="18" cy="19" r="3"/><circle cx="6" cy="12" r="3"/><path stroke-linecap="round" stroke-linejoin="round" d="m15.408 6.512l-6.814 3.975m6.814 7.001l-6.814-3.975"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Newspaper(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" d="M5 21h12a4 4 0 0 0 4-4V5a2 2 0 0 0-2-2H9a2 2 0 0 0-2 2v13c0 1.657-.343 3-2 3"/><path stroke-linejoin="round" d="M3 10a2 2 0 0 1 2-2h2v10.5c0 1.38-.62 2.5-2 2.5s-2-1.12-2-2.5z"/><circle cx="12" cy="8" r="1"/><path d="M11 14h6m-6 3h3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NextjsFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#akarIconsNextjsFill0)"><path fill="currentColor" d="M11.214.006c-.052.005-.216.022-.364.033c-3.408.308-6.6 2.147-8.624 4.974a11.88 11.88 0 0 0-2.118 5.243c-.096.66-.108.854-.108 1.748s.012 1.089.108 1.748c.652 4.507 3.86 8.293 8.209 9.696c.779.251 1.6.422 2.533.526c.364.04 1.936.04 2.3 0c1.611-.179 2.977-.578 4.323-1.265c.207-.105.247-.134.219-.157a211.64 211.64 0 0 1-1.955-2.62l-1.919-2.593l-2.404-3.559a342.499 342.499 0 0 0-2.422-3.556c-.009-.003-.018 1.578-.023 3.51c-.007 3.38-.01 3.516-.052 3.596a.426.426 0 0 1-.206.213c-.075.038-.14.045-.495.045H7.81l-.108-.068a.44.44 0 0 1-.157-.172l-.05-.105l.005-4.704l.007-4.706l.073-.092a.644.644 0 0 1 .174-.143c.096-.047.133-.051.54-.051c.478 0 .558.018.682.154c.035.038 1.337 2 2.895 4.362l4.734 7.172l1.9 2.878l.097-.063a12.318 12.318 0 0 0 2.465-2.163a11.947 11.947 0 0 0 2.825-6.135c.096-.66.108-.854.108-1.748s-.012-1.088-.108-1.748C23.24 5.75 20.032 1.963 15.683.56a12.6 12.6 0 0 0-2.498-.523c-.226-.024-1.776-.05-1.97-.03m4.913 7.26a.473.473 0 0 1 .237.276c.018.06.023 1.365.018 4.305l-.007 4.218l-.743-1.14l-.746-1.14V10.72c0-1.983.009-3.097.023-3.151a.478.478 0 0 1 .232-.296c.097-.05.132-.054.5-.054c.347 0 .408.005.486.047"/></g><defs><clipPath id="akarIconsNextjsFill0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NodeFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 23.956c-.342 0-.66-.089-.957-.243l-3.029-1.738c-.455-.242-.227-.33-.09-.374c.614-.198.728-.242 1.366-.595c.068-.044.16-.022.228.022l2.323 1.343c.09.044.205.044.273 0l9.087-5.084c.09-.044.136-.132.136-.242V6.899c0-.11-.045-.198-.136-.242l-9.087-5.061c-.091-.044-.205-.044-.273 0L2.754 6.657c-.091.044-.137.154-.137.242v10.146c0 .088.046.198.137.242l2.482 1.387c1.344.66 2.186-.11 2.186-.88V7.78c0-.132.114-.264.274-.264h1.161c.137 0 .273.11.273.264v10.013c0 1.739-.979 2.751-2.687 2.751c-.524 0-.934 0-2.095-.55l-2.391-1.32A1.847 1.847 0 0 1 1 17.067V6.921c0-.66.364-1.276.957-1.606L11.044.23a2.095 2.095 0 0 1 1.912 0l9.088 5.084c.592.33.956.946.956 1.606v10.146c0 .66-.364 1.276-.956 1.607l-9.087 5.083a2.4 2.4 0 0 1-.957.198m2.801-6.977c-3.985 0-4.805-1.76-4.805-3.257c0-.132.114-.264.273-.264h1.184c.137 0 .25.088.25.22c.183 1.166.707 1.738 3.121 1.738c1.913 0 2.733-.418 2.733-1.408c0-.572-.228-.99-3.211-1.276c-2.483-.243-4.031-.77-4.031-2.685c0-1.783 1.548-2.84 4.145-2.84c2.915 0 4.35.969 4.532 3.082a.347.347 0 0 1-.069.198c-.045.044-.113.088-.182.088h-1.184a.265.265 0 0 1-.25-.198c-.274-1.21-.98-1.607-2.847-1.607c-2.096 0-2.346.704-2.346 1.233c0 .638.296.836 3.12 1.188c2.801.352 4.122.858 4.122 2.75c-.023 1.938-1.662 3.038-4.555 3.038"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NormalScreen(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 2v5H2m15-5v5h5M2 17h5v5m15-5h-5v5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NpmFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#akarIconsNpmFill0)"><path fill="currentColor" fill-rule="evenodd" d="M24 0H0v24h24zM2.578 2.578H21.42V21.42h-4.75V7.33h-4.751v14.09h-9.34z" clip-rule="evenodd"/></g><defs><clipPath id="akarIconsNpmFill0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Octagon(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="2" d="M11.235 2.317a2 2 0 0 1 1.53 0l5.54 2.295a2 2 0 0 1 1.083 1.082l2.295 5.54a2 2 0 0 1 0 1.531l-2.295 5.54a2 2 0 0 1-1.082 1.083l-5.54 2.295a2 2 0 0 1-1.531 0l-5.54-2.295a2 2 0 0 1-1.083-1.082l-2.295-5.54a2 2 0 0 1 0-1.531l2.295-5.54a2 2 0 0 1 1.082-1.083z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OctagonFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.235 1.317a2 2 0 0 1 1.53 0l6.248 2.588a2 2 0 0 1 1.082 1.082l2.588 6.248a2 2 0 0 1 0 1.53l-2.588 6.248a2 2 0 0 1-1.082 1.082l-6.248 2.588a2 2 0 0 1-1.53 0l-6.248-2.588a2 2 0 0 1-1.082-1.082l-2.588-6.248a2 2 0 0 1 0-1.53l2.588-6.248a2 2 0 0 1 1.082-1.082z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OctocatFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.422 20.081c0 .896.01 1.753.016 2.285a.617.617 0 0 0 .422.58c2.078.686 4.317.718 6.414.091l.292-.087a.67.67 0 0 0 .478-.638c.005-.733.017-2.017.017-3.53c0-1.372-.477-2.25-1.031-2.707c3.399-.366 6.97-1.61 6.97-7.227c0-1.61-.592-2.91-1.566-3.934c.153-.366.688-1.866-.153-3.878c0 0-1.28-.403-4.201 1.5a14.76 14.76 0 0 0-3.82-.494c-1.298 0-2.597.165-3.819.494C5.52.65 4.24 1.036 4.24 1.036c-.84 2.012-.306 3.512-.153 3.878a5.565 5.565 0 0 0-1.566 3.934c0 5.598 3.552 6.86 6.951 7.227c-.439.366-.84 1.006-.973 1.957c-.879.384-3.075 1.006-4.45-1.207c-.286-.44-1.146-1.519-2.349-1.5c-1.28.018-.516.695.02.97c.648.347 1.393 1.646 1.565 2.067c.306.823 1.299 2.396 5.137 1.72"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OpenEnvelope(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path stroke-linecap="round" d="M2 11.083a4 4 0 0 1 1.706-3.277l6-4.2a4 4 0 0 1 4.588 0l6 4.2A4 4 0 0 1 22 11.083V19a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2z"/><path d="m2.5 9.5l7.001 5.501a4 4 0 0 0 4.998 0L21.5 9.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Oval(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<ellipse cx="12" cy="12" fill="none" stroke="currentColor" stroke-width="2" rx="8" ry="10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PanelBottom(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><rect width="20" height="18" x="2" y="3" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path d="M22 15H2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PanelLeft(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><rect width="20" height="18" x="2" y="3" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path d="M9 3v18"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PanelRight(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><rect width="20" height="18" x="2" y="3" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path d="M15 3v18"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PanelSplit(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><rect width="20" height="18" x="2" y="3" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path d="M9 3v18m13-9H9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PanelSplitColumn(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><rect width="20" height="18" x="2" y="3" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path d="M22 12H2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PanelSplitRow(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><rect width="20" height="18" x="2" y="3" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path d="M12 3v18"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PanelTop(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><rect width="20" height="18" x="2" y="3" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path d="M22 9H2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paper(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path stroke-linecap="round" d="M4 4v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8.342a2 2 0 0 0-.602-1.43l-4.44-4.342A2 2 0 0 0 13.56 2H6a2 2 0 0 0-2 2"/><path d="M14 2v4a2 2 0 0 0 2 2h4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaperAirplane(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9.474 16l9.181 3.284a1.1 1.1 0 0 0 1.462-.887L21.99 4.239c.114-.862-.779-1.505-1.567-1.13L2.624 11.605c-.88.42-.814 1.69.106 2.017l2.44.868l1.33.467M13 17.26l-1.99 3.326c-.65.808-1.959.351-1.959-.683V17.24a2 2 0 0 1 .53-1.356l3.871-4.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Parallelogram(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="2" d="M5.586 6.45A2 2 0 0 1 7.509 5h11.84a2 2 0 0 1 1.923 2.55l-2.858 10A2 2 0 0 1 16.491 19H4.651a2 2 0 0 1-1.923-2.55z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pause(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M7 5v14M17 5v14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pencil(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M4.333 16.048L16.57 3.81a2.56 2.56 0 0 1 3.62 3.619L7.951 19.667a2 2 0 0 1-1.022.547L3 21l.786-3.93a2 2 0 0 1 .547-1.022"/><path d="m14.5 6.5l3 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pentagon(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="2" d="M10.75 3a2 2 0 0 1 2.5 0l7.63 6.103a2 2 0 0 1 .63 2.246l-3.031 8.334A2 2 0 0 1 16.599 21H7.401a2 2 0 0 1-1.88-1.317l-3.03-8.334a2 2 0 0 1 .63-2.246z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PentagonFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.126 2.219a3 3 0 0 1 3.748 0l7.63 6.104a3 3 0 0 1 .945 3.367l-3.03 8.335A3 3 0 0 1 16.599 22H7.401a3 3 0 0 1-2.82-1.975l-3.03-8.334a3 3 0 0 1 .945-3.368z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PeopleGroup(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="m16.719 19.752l-.64-5.124A3 3 0 0 0 13.101 12h-2.204a3 3 0 0 0-2.976 2.628l-.641 5.124A2 2 0 0 0 9.266 22h5.468a2 2 0 0 0 1.985-2.248"/><circle cx="12" cy="5" r="3"/><circle cx="4" cy="9" r="2"/><circle cx="20" cy="9" r="2"/><path stroke-linecap="round" stroke-linejoin="round" d="M4 14h-.306a2 2 0 0 0-1.973 1.671l-.333 2A2 2 0 0 0 3.361 20H7m13-6h.306a2 2 0 0 1 1.973 1.671l.333 2A2 2 0 0 1 20.639 20H17"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PeopleMultiple(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><circle cx="7" cy="6" r="3"/><path stroke-linecap="round" stroke-linejoin="round" d="M10 13H5.818a3 3 0 0 0-2.964 2.537L2.36 18.69A2 2 0 0 0 4.337 21H9m12.64-2.309l-.494-3.154A3 3 0 0 0 18.182 13h-2.364a3 3 0 0 0-2.964 2.537l-.493 3.154A2 2 0 0 0 14.337 21h5.326a2 2 0 0 0 1.976-2.309"/><circle cx="17" cy="6" r="3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Percentage(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 19L19 5"/><circle cx="7" cy="7" r="3"/><circle cx="17" cy="17" r="3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Person(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="7" r="5"/><path stroke-linecap="round" stroke-linejoin="round" d="M17 14h.352a3 3 0 0 1 2.976 2.628l.391 3.124A2 2 0 0 1 18.734 22H5.266a2 2 0 0 1-1.985-2.248l.39-3.124A3 3 0 0 1 6.649 14H7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonAdd(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="7" r="5"/><path stroke-linecap="round" stroke-linejoin="round" d="M17 22H5.266a2 2 0 0 1-1.985-2.248l.39-3.124A3 3 0 0 1 6.649 14H7m12 0v4m-2-2h4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonCheck(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="7" r="5"/><path stroke-linecap="round" stroke-linejoin="round" d="M17 22H5.266a2 2 0 0 1-1.985-2.248l.39-3.124A3 3 0 0 1 6.649 14H7m10 2.5l1.5 1.5l2.5-3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonCross(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="7" r="5"/><path stroke-linecap="round" stroke-linejoin="round" d="M17 22H5.266a2 2 0 0 1-1.985-2.248l.39-3.124A3 3 0 0 1 6.649 14H7"/><path stroke-linecap="round" d="m21 18l-3-3m3 0l-3 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Phone(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.554 6.24L7.171 2.335c-.39-.45-1.105-.448-1.558.006L2.831 5.128c-.828.829-1.065 2.06-.586 3.047a29.207 29.207 0 0 0 13.561 13.58c.986.479 2.216.242 3.044-.587l2.808-2.813c.455-.455.456-1.174.002-1.564l-3.92-3.365c-.41-.352-1.047-.306-1.458.106l-1.364 1.366a.462.462 0 0 1-.553.088a14.557 14.557 0 0 1-5.36-5.367a.463.463 0 0 1 .088-.554l1.36-1.361c.412-.414.457-1.054.101-1.465"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhpFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 5.5C5.271 5.5 0 8.355 0 12s5.271 6.5 12 6.5s12-2.855 12-6.5s-5.271-6.5-12-6.5m-1.246 2h1.31l-.416 2h1.17c.742 0 1.24.104 1.524.363c.277.256.361.676.25 1.248l-.52 2.389H12.74l.479-2.209c.058-.303.035-.514-.067-.625c-.101-.111-.324-.166-.658-.166h-1.049l-.633 3H9.5zM5 9.5h2.666c1.271 0 2.041.852 1.74 2.123C9.056 13.1 8.12 13.5 6.396 13.5h-.824L5.311 15H3.986zm10.5 0h2.666c1.271 0 2.041.852 1.74 2.123c-.35 1.477-1.287 1.877-3.01 1.877h-.824l-.261 1.5h-1.325zm-9.365 1l-.377 2h.855c.74 0 1.428-.084 1.543-1.187c.043-.428-.135-.813-.99-.813zm10.5 0l-.377 2h.855c.74 0 1.428-.084 1.543-1.187c.043-.428-.134-.813-.99-.813z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pin(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 10.726l-3 .505L11.23 2l-.504 3M12 16.881l-.77 2.042l7.693-7.692l-2.042.769m-1.804 3.077L21 21M3.538 9.692l6.154-6.154l.236.341a52.22 52.22 0 0 0 7.376 8.518l.235.218l-4.924 4.923l-.218-.234A52.22 52.22 0 0 0 3.88 9.928z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinterestFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<defs><path id="akarIconsPinterestFill0" fill="#fff" d="M0 0h24v24H0z"/></defs><g fill="none"><g clip-path="url(#akarIconsPinterestFill1)"><g clip-path="url(#akarIconsPinterestFill2)"><path fill="currentColor" d="M0 12c0 5.123 3.211 9.497 7.73 11.218c-.11-.937-.227-2.482.025-3.566c.217-.932 1.401-5.938 1.401-5.938s-.357-.715-.357-1.774c0-1.66.962-2.9 2.161-2.9c1.02 0 1.512.765 1.512 1.682c0 1.025-.653 2.557-.99 3.978c-.281 1.189.597 2.159 1.769 2.159c2.123 0 3.756-2.239 3.756-5.471c0-2.861-2.056-4.86-4.991-4.86c-3.398 0-5.393 2.549-5.393 5.184c0 1.027.395 2.127.889 2.726a.36.36 0 0 1 .083.343c-.091.378-.293 1.189-.332 1.355c-.053.218-.173.265-.4.159c-1.492-.694-2.424-2.875-2.424-4.627c0-3.769 2.737-7.229 7.892-7.229c4.144 0 7.365 2.953 7.365 6.899c0 4.117-2.595 7.431-6.199 7.431c-1.211 0-2.348-.63-2.738-1.373c0 0-.599 2.282-.744 2.84c-.282 1.084-1.064 2.456-1.549 3.235C9.584 23.815 10.77 24 12 24c6.627 0 12-5.373 12-12S18.627 0 12 0S0 5.373 0 12"/></g></g><defs><clipPath id="akarIconsPinterestFill1"><use href="#akarIconsPinterestFill0"/></clipPath><clipPath id="akarIconsPinterestFill2"><use href="#akarIconsPinterestFill0"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plane(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 11l-2 4l3.408 1.363a4 4 0 0 1 2.229 2.229L9 22l4-2l-1.21-2.42a2 2 0 0 1 .679-2.56L14 14l4 7l3-4l-2.29-7.469l.715-.714c1.412-1.412 2.71-3.682 1.075-5.317c-1.635-1.635-3.91-.34-5.316 1.077l-.72.708L7 3L3 6l7 4l-1.02 1.531a2 2 0 0 1-2.56.68z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlaneFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m4 11l-2 4l3.408 1.363a4 4 0 0 1 2.229 2.229L9 22l4-2l-1.21-2.42a2 2 0 0 1 .679-2.56L14 14l4 7l3-4l-2.29-7.469l.715-.714c1.412-1.412 2.71-3.682 1.075-5.317c-1.635-1.635-3.91-.34-5.316 1.077l-.72.708L7 3L3 6l7 4l-1.02 1.531a2 2 0 0 1-2.56.68z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Planet(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="8"/><path d="M8.399 4.849C5.372 2.582 2.972 1.489 2.23 2.23c-1.174 1.174 2.248 6.5 7.643 11.895c5.396 5.395 10.722 8.817 11.895 7.643c.431-.43.243-1.421-.435-2.769"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plant(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11.964 6.97s-3.075.306-4.685-1.035C5.669 4.593 6.036 2.03 6.036 2.03s3.075-.306 4.686 1.035c1.61 1.342 1.242 3.905 1.242 3.905"/><path d="M12.036 6.97s3.075.306 4.685-1.035c1.61-1.342 1.243-3.905 1.243-3.905s-3.075-.306-4.685 1.035c-1.61 1.342-1.243 3.905-1.243 3.905M4 11a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1zm1 3h14l-2 8H7z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Play(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 4v16m14-8L6 20m14-8L6 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plus(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M12 20v-8m0 0V4m0 8h8m-8 0H4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PointerDownFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M20.05 17.65a3 3 0 0 0 1.2-2.4v-11a3 3 0 0 0-3-3h-12a3 3 0 0 0-3 3v11a3 3 0 0 0 1.2 2.4l6 4.5a3 3 0 0 0 3.6 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PointerHand(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 16V9.5c0-.828-.641-1.5-1.48-1.5C18 8 17 8.3 17 9.5M8 14V5.52M20 16c0 4-3.134 6-7 6s-5.196-1-8.196-6l-1.571-2.605c-.536-.868-.107-1.994.881-2.314a1.657 1.657 0 0 1 1.818.552L8 14.033"/><path d="M14 11V7.5A1.5 1.5 0 0 1 15.5 6v0A1.5 1.5 0 0 1 17 7.5V11m-6 0V6.5A1.5 1.5 0 0 1 12.5 5v0A1.5 1.5 0 0 1 14 6.5V11m-6 0V2.5A1.5 1.5 0 0 1 9.5 1v0A1.5 1.5 0 0 1 11 2.5V11"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PointerLeftFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.6 4.2A3 3 0 0 1 9 3h11a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H9a3 3 0 0 1-2.4-1.2l-4.5-6a3 3 0 0 1 0-3.6z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PointerRightFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M17.9 4.2A3 3 0 0 0 15.5 3h-11a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h11a3 3 0 0 0 2.4-1.2l4.5-6a3 3 0 0 0 0-3.6z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PointerUpFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M20.05 6.35a3 3 0 0 1 1.2 2.4v11a3 3 0 0 1-3 3h-12a3 3 0 0 1-3-3v-11a3 3 0 0 1 1.2-2.4l6-4.5a3 3 0 0 1 3.6 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PointingUp(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="m19 16l.87-11.735a2.102 2.102 0 0 0-4.181-.433L15 9m-7 6v-3a2 2 0 1 0-4 0v4m8-3v-1.5a2 2 0 1 0-4 0V15"/><path stroke-linecap="round" stroke-linejoin="round" d="M16 13v-2a2 2 0 1 0-4 0v2"/><path d="M19 16c-.536 4-3.358 6-7.5 6C7.358 22 4 20 4 16"/><path stroke-linecap="round" stroke-linejoin="round" d="M13.692 17H11a2 2 0 1 1 0-4h4c2.21 0 4.5 2 3.5 5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PostgresqlFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.805 1a9.85 9.85 0 0 0-2.603.37l-.06.018a10.629 10.629 0 0 0-1.615-.151c-1.113-.019-2.07.243-2.84.68c-.76-.256-2.336-.697-3.997-.609c-1.157.061-2.419.402-3.354 1.36c-.933.958-1.426 2.44-1.322 4.457c.028.557.191 1.464.463 2.64c.27 1.175.652 2.55 1.127 3.805c.475 1.256.996 2.384 1.81 3.15c.406.384.965.707 1.624.68c.463-.018.882-.215 1.243-.506c.176.225.364.323.535.414c.215.114.425.192.642.244a4.61 4.61 0 0 0 1.84.091c.267-.043.548-.127.828-.247c.01.302.022.598.035.898c.038.95.063 1.827.357 2.596c.047.126.176.773.687 1.344c.51.572 1.51.928 2.648.692c.803-.167 1.825-.468 2.503-1.404c.67-.926.973-2.254 1.033-4.409c.015-.116.033-.215.052-.308l.16.014h.018c.857.038 1.787-.08 2.564-.43c.688-.31 1.208-.622 1.587-1.177c.095-.137.199-.303.227-.59c.028-.285-.14-.733-.421-.939c-.563-.414-.916-.257-1.295-.18c-.373.08-.753.124-1.136.133c1.093-1.784 1.876-3.68 2.323-5.358c.264-.99.413-1.903.425-2.701c.012-.798-.055-1.505-.548-2.117c-1.541-1.91-3.708-2.438-5.384-2.456c-.052-.001-.104-.002-.156-.001zm-.044.587c1.585-.015 3.611.417 5.065 2.22c.327.405.424.997.413 1.727c-.012.729-.151 1.601-.405 2.557c-.493 1.852-1.425 4.01-2.738 5.948a.724.724 0 0 0 .15.079c.274.11.898.204 2.145-.044c.313-.065.543-.108.781.068a.478.478 0 0 1 .173.39a.635.635 0 0 1-.123.308c-.24.351-.716.684-1.326.958c-.539.244-1.313.371-1.999.379c-.344.003-.661-.023-.93-.104l-.018-.006c-.104.971-.343 2.89-.498 3.765c-.125.706-.343 1.267-.76 1.687c-.416.42-1.004.673-1.796.838c-.981.204-1.696-.016-2.157-.393c-.46-.375-.671-.874-.798-1.18c-.087-.21-.132-.483-.176-.848a18.073 18.073 0 0 1-.097-1.315a45.725 45.725 0 0 1-.028-2.312c-.41.363-.92.605-1.467.696c-.65.107-1.232.002-1.579-.082a2.19 2.19 0 0 1-.49-.185c-.162-.083-.315-.177-.417-.363a.5.5 0 0 1-.054-.35a.559.559 0 0 1 .206-.303c.188-.148.435-.23.808-.306c.68-.135.917-.228 1.061-.339c.123-.095.262-.287.508-.57a1.166 1.166 0 0 1-.003-.037a2.864 2.864 0 0 1-1.257-.328c-.141.144-.865.887-1.748 1.917c-.371.431-.781.678-1.214.696c-.433.018-.824-.194-1.156-.506c-.665-.626-1.195-1.703-1.657-2.92c-.46-1.218-.836-2.574-1.102-3.729c-.268-1.155-.426-2.086-.448-2.535c-.1-1.909.36-3.195 1.15-4.006c.79-.811 1.872-1.118 2.928-1.177c1.894-.106 3.693.535 4.057.673c.701-.462 1.604-.75 2.733-.732a7.185 7.185 0 0 1 1.588.2l.019-.008c.229-.078.462-.144.698-.196a9.362 9.362 0 0 1 1.957-.23zm.143.614h-.137a8.502 8.502 0 0 0-1.61.176a7.048 7.048 0 0 1 2.692 2.062a7.72 7.72 0 0 1 1.07 1.76c.104.242.174.447.213.605c.02.08.034.147.038.217a.392.392 0 0 1-.011.132l-.006.012c.029.803-.176 1.347-.201 2.113c-.019.556.127 1.209.163 1.92c.034.67-.049 1.405-.497 2.127c.038.044.072.088.108.132c1.185-1.81 2.04-3.814 2.495-5.521c.243-.92.373-1.753.384-2.413c.01-.66-.117-1.139-.279-1.338c-1.268-1.573-2.983-1.974-4.422-1.985zm-4.525.235c-1.117.002-1.919.33-2.526.82c-.627.507-1.047 1.2-1.323 1.911a7.898 7.898 0 0 0-.485 2.213l.013-.007c.337-.184.78-.367 1.254-.473c.475-.106.986-.139 1.449.035c.463.175.846.584.985 1.206c.665 2.986-.207 4.096-.529 4.933a8.628 8.628 0 0 0-.312.929c.04-.01.08-.02.121-.024a1.06 1.06 0 0 1 .51.1c.324.13.546.402.666.714c.031.082.054.17.067.26c.014.038.02.077.019.117a49.059 49.059 0 0 0 .012 3.426c.022.494.054.928.095 1.271c.04.342.098.602.135.69c.12.294.297.678.617.939c.32.26.777.434 1.614.26c.726-.151 1.174-.36 1.474-.663c.298-.301.477-.72.591-1.363c.171-.963.515-3.754.556-4.28c-.018-.395.042-.7.172-.932c.135-.238.343-.384.522-.463c.09-.04.174-.066.243-.085a5.487 5.487 0 0 0-.23-.298a4.065 4.065 0 0 1-.629-1.007a7.578 7.578 0 0 0-.243-.443c-.125-.22-.284-.495-.45-.804c-.333-.619-.695-1.369-.883-2.1c-.187-.729-.215-1.484.265-2.017c.426-.473 1.172-.669 2.293-.559c-.033-.096-.053-.176-.109-.304a7.125 7.125 0 0 0-.983-1.617c-.95-1.178-2.487-2.346-4.863-2.384h-.108zm-6.276.047c-.12 0-.24.004-.36.01c-.954.053-1.856.322-2.501.986c-.647.663-1.072 1.751-.98 3.553c.019.34.172 1.296.434 2.43c.262 1.136.634 2.471 1.08 3.65c.446 1.18.988 2.207 1.502 2.693c.259.243.484.341.688.333c.205-.01.451-.124.753-.475a40.03 40.03 0 0 1 1.71-1.877a3.206 3.206 0 0 1-.932-1.307a3.116 3.116 0 0 1-.17-1.58c.097-.678.11-1.312.099-1.812c-.012-.488-.048-.812-.048-1.015v-.028a8.806 8.806 0 0 1 .559-3.095c.264-.682.658-1.375 1.249-1.936c-.58-.185-1.61-.467-2.725-.52a7.4 7.4 0 0 0-.36-.01zm11.714 4.842c-.641.008-1.001.169-1.19.379c-.268.298-.293.82-.127 1.464c.165.644.507 1.365.829 1.963c.16.3.316.57.442.788c.127.22.22.376.276.51c.052.122.11.23.168.331c.248-.509.293-1.008.267-1.529c-.033-.644-.187-1.303-.164-1.97c.025-.78.184-1.289.198-1.892a5.639 5.639 0 0 0-.699-.044m-7.78.105a2.743 2.743 0 0 0-.582.068a4.49 4.49 0 0 0-1.09.412c-.115.06-.226.13-.33.209l-.02.018c.006.134.033.459.045.936c.01.523-.002 1.19-.106 1.91c-.226 1.568.946 2.866 2.324 2.868c.08-.322.213-.648.345-.992c.384-1.003 1.139-1.734.503-4.589c-.104-.467-.31-.656-.594-.763a1.431 1.431 0 0 0-.495-.077m7.48.187h.048c.062.002.12.009.17.02a.396.396 0 0 1 .13.051a.153.153 0 0 1 .071.1v.008a.215.215 0 0 1-.034.124a.614.614 0 0 1-.104.137a.646.646 0 0 1-.364.195a.57.57 0 0 1-.388-.095a.569.569 0 0 1-.123-.108a.235.235 0 0 1-.06-.116a.151.151 0 0 1 .04-.118a.361.361 0 0 1 .111-.082a1.256 1.256 0 0 1 .504-.118zm-7.388.154c.05 0 .103.005.157.012c.144.02.273.057.371.112c.048.025.09.057.126.097a.214.214 0 0 1 .042.073a.195.195 0 0 1 .009.083a.274.274 0 0 1-.071.141a.608.608 0 0 1-.135.12a.619.619 0 0 1-.424.103a.694.694 0 0 1-.396-.209a.652.652 0 0 1-.112-.15a.25.25 0 0 1-.039-.162c.014-.1.099-.15.18-.18a.842.842 0 0 1 .29-.036zm8.56 6.732h-.003c-.139.05-.253.07-.35.11a.423.423 0 0 0-.225.197c-.06.105-.11.292-.095.61a.49.49 0 0 0 .14.064c.161.048.432.08.735.075c.602-.007 1.344-.143 1.738-.321c.323-.146.623-.336.891-.564c-1.317.264-2.06.194-2.517.011a1.247 1.247 0 0 1-.314-.183zm-7.588.086h-.02c-.05.004-.123.02-.263.172c-.33.358-.444.582-.716.792c-.27.21-.623.321-1.327.461c-.223.044-.35.093-.436.132c.028.022.025.028.066.049c.103.055.236.103.342.13c.303.073.8.159 1.319.073c.518-.086 1.058-.327 1.518-.953c.08-.108.088-.268.023-.44c-.067-.17-.211-.318-.313-.36a.632.632 0 0 0-.193-.054z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PriceCut(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 7v8a2 2 0 0 0 2 2h1M2 7V5a2 2 0 0 1 2-2h2M2 7a4 4 0 0 0 4-4m0 0h13m-9.236 9A3 3 0 0 1 14 7.764M2 20L20 2m1.22 16.047l.549-6.261c.075-.865-.598-1.63-1.504-1.71l-.82-.071m1.776 8.042l-.137 1.566c-.076.864-.872 1.501-1.778 1.422l-1.64-.144m3.555-2.844c-1.813-.158-3.405 1.115-3.556 2.844m0 0L7 19.958m4.347-3.477c.409.52.94.785 1.67.849c1.359.119 2.553-.836 2.666-2.133c.061-.696-.085-1.269-.551-1.743"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ProductHuntFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#akarIconsProductHuntFill0)"><path fill="currentColor" d="M15.402 10.2c0 .992-.808 1.8-1.8 1.8H10.2V8.4h3.402c.992 0 1.8.808 1.8 1.8M24 12c0 6.629-5.371 12-12 12S0 18.629 0 12S5.371 0 12 0s12 5.371 12 12m-6.198-1.8c0-2.318-1.883-4.2-4.2-4.2H7.8v12h2.4v-3.6h3.402c2.317 0 4.2-1.882 4.2-4.2"/></g><defs><clipPath id="akarIconsProductHuntFill0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PythonFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g fill="currentColor" clip-path="url(#akarIconsPythonFill0)"><path d="M11.914 0C5.82 0 6.2 2.656 6.2 2.656l.007 2.752h5.814v.826H3.9S0 5.789 0 11.969c0 6.18 3.403 5.96 3.403 5.96h2.03v-2.867s-.109-3.42 3.35-3.42h5.766s3.24.052 3.24-3.148V3.202S18.28 0 11.913 0M8.708 1.85c.578 0 1.046.47 1.046 1.052c0 .581-.468 1.051-1.046 1.051c-.579 0-1.046-.47-1.046-1.051c0-.582.467-1.052 1.046-1.052"/><path d="M12.087 24c6.092 0 5.712-2.656 5.712-2.656l-.007-2.752h-5.814v-.826h8.123s3.9.445 3.9-5.735c0-6.18-3.404-5.96-3.404-5.96h-2.03v2.867s.109 3.42-3.35 3.42H9.452s-3.24-.052-3.24 3.148v5.292S5.72 24 12.087 24m3.206-1.85c-.579 0-1.046-.47-1.046-1.052c0-.581.467-1.051 1.046-1.051c.578 0 1.046.47 1.046 1.051c0 .582-.468 1.052-1.046 1.052"/></g><defs><clipPath id="akarIconsPythonFill0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Question(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="10"/><path stroke-linecap="round" d="M10 8.484C10.5 7.494 11 7 12 7c1.246 0 2 .989 2 1.978s-.5 1.483-2 2.473V13m0 3.5v.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11s11-4.925 11-11S18.075 1 12 1m-1.108 7.935c.23-.453.4-.668.541-.78c.106-.084.25-.155.567-.155c.625 0 1 .47 1 .978c0 .278-.054.416-.202.592c-.207.246-.59.545-1.348 1.046l-.45.296V13a1 1 0 1 0 2 0v-1.017c.542-.374.997-.732 1.327-1.124c.477-.566.673-1.17.673-1.881C15 7.508 13.867 6 12 6c-.684 0-1.289.176-1.808.587c-.484.383-.814.91-1.084 1.445a1 1 0 1 0 1.784.903M13 16.5a1 1 0 1 0-2 0v.5a1 1 0 1 0 2 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Radio(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12a9 9 0 1 1 18 0a9 9 0 0 1-18 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RadioFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2m0 6a4 4 0 1 0 0 8a4 4 0 0 0 0-8" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Radish(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path stroke-linecap="round" d="M16 10H8c-2.188 0-3.698 1.415-3.935 3.282c-.325 2.56.529 4.105 3.634 5.128c1.394.46 2.579 1.464 3.01 2.868l.223.722l.095-.082a8 8 0 0 1 2.175-1.331l1.921-.789c3.286-1.348 5.22-3.408 4.826-6.516C19.712 11.415 18.188 10 16 10m-4.036-3.03s-3.075.306-4.685-1.035C5.669 4.593 6.036 2.03 6.036 2.03s3.075-.306 4.686 1.035c1.61 1.342 1.242 3.905 1.242 3.905"/><path stroke-linecap="round" d="M12.036 6.97s3.075.306 4.685-1.035c1.61-1.342 1.243-3.905 1.243-3.905s-3.075-.306-4.685 1.035c-1.61 1.342-1.243 3.905-1.243 3.905"/><path d="M19 11.5c-.5 1-3.134 1.5-7 1.5s-6.5-.5-7-1.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReactFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g fill="currentColor" clip-path="url(#akarIconsReactFill0)"><path d="M24 11.689c0-1.59-1.991-3.097-5.044-4.031c.705-3.111.392-5.587-.988-6.38a2.145 2.145 0 0 0-1.096-.273v1.09c.225 0 .406.045.558.128c.665.382.954 1.834.729 3.703c-.054.46-.142.944-.25 1.438a23.706 23.706 0 0 0-3.106-.533a23.857 23.857 0 0 0-2.035-2.446c1.595-1.482 3.092-2.294 4.11-2.294V1c-1.346 0-3.107.959-4.888 2.622C10.21 1.97 8.448 1.02 7.103 1.02v1.09c1.013 0 2.515.808 4.11 2.28c-.685.72-1.37 1.536-2.021 2.441a22.844 22.844 0 0 0-3.111.538a14.683 14.683 0 0 1-.255-1.418c-.23-1.87.054-3.322.715-3.708c.146-.088.337-.128.562-.128v-1.09c-.41 0-.783.088-1.105.273c-1.375.793-1.683 3.263-.974 6.365C1.981 8.603 0 10.104 0 11.689c0 1.59 1.991 3.097 5.044 4.03c-.705 3.112-.392 5.588.988 6.38c.318.186.69.274 1.1.274c1.346 0 3.107-.959 4.888-2.622c1.78 1.653 3.541 2.602 4.887 2.602a2.18 2.18 0 0 0 1.105-.274c1.375-.792 1.683-3.262.974-6.364C22.019 14.781 24 13.274 24 11.689m-6.37-3.263a22.023 22.023 0 0 1-.66 1.932a26.444 26.444 0 0 0-1.345-2.319c.695.103 1.365.23 2.006.387m-2.24 5.21a25.94 25.94 0 0 1-1.179 1.869a25.453 25.453 0 0 1-4.412.005a25.457 25.457 0 0 1-2.201-3.806a26.064 26.064 0 0 1 2.192-3.82a25.455 25.455 0 0 1 4.411-.006c.406.582.803 1.204 1.184 1.86c.372.64.71 1.29 1.018 1.946a27.41 27.41 0 0 1-1.013 1.952M16.97 13c.264.656.49 1.311.675 1.947c-.64.157-1.316.289-2.015.391A27.044 27.044 0 0 0 16.97 13m-4.96 5.22c-.455-.47-.91-.993-1.36-1.566c.44.02.89.035 1.345.035c.46 0 .915-.01 1.36-.035c-.44.573-.895 1.096-1.345 1.566m-3.64-2.882a22.113 22.113 0 0 1-2.006-.386c.181-.631.406-1.282.66-1.932c.201.39.412.782.642 1.174c.23.391.464.773.704 1.144m3.615-10.18c.455.47.91.993 1.36 1.566c-.44-.02-.89-.035-1.345-.035c-.46 0-.915.01-1.36.035c.44-.573.895-1.096 1.345-1.566M8.365 8.04a27.02 27.02 0 0 0-1.34 2.333a20.96 20.96 0 0 1-.675-1.947c.64-.152 1.316-.284 2.015-.386m-4.427 6.124c-1.732-.738-2.852-1.707-2.852-2.475s1.12-1.742 2.852-2.475c.42-.181.88-.343 1.355-.494c.279.958.646 1.956 1.1 2.979a23.165 23.165 0 0 0-1.085 2.964a14.875 14.875 0 0 1-1.37-.499m2.632 6.99c-.665-.38-.954-1.834-.729-3.702c.054-.46.142-.945.25-1.439c.958.235 2.005.416 3.106.534a23.87 23.87 0 0 0 2.035 2.446c-1.595 1.482-3.092 2.294-4.11 2.294a1.167 1.167 0 0 1-.552-.132m11.604-3.727c.23 1.869-.054 3.322-.715 3.708c-.146.088-.337.127-.562.127c-1.013 0-2.515-.807-4.11-2.28a23.01 23.01 0 0 0 2.021-2.44a22.843 22.843 0 0 0 3.111-.538c.113.494.2.968.255 1.423m1.883-3.263c-.42.181-.88.343-1.355.494a23.482 23.482 0 0 0-1.1-2.979c.45-1.017.811-2.01 1.085-2.964a15.3 15.3 0 0 1 1.375.499c1.732.738 2.852 1.707 2.852 2.475c-.005.768-1.125 1.742-2.857 2.475"/><path d="M11.995 13.925a2.236 2.236 0 1 0 0-4.472a2.236 2.236 0 0 0 0 4.472"/></g><defs><clipPath id="akarIconsReactFill0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Receipt(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" d="M19 21H7a4 4 0 0 1-4-4V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v13c0 1.657.343 3 2 3"/><path stroke-linejoin="round" d="M21 10a2 2 0 0 0-2-2h-2v10.5c0 1.38.62 2.5 2 2.5s2-1.12 2-2.5z"/><path d="M13 11H7m6-4H7m3 8H7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reciept(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" d="M19 21H7a4 4 0 0 1-4-4V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v13c0 1.657.343 3 2 3"/><path stroke-linejoin="round" d="M21 10a2 2 0 0 0-2-2h-2v10.5c0 1.38.62 2.5 2 2.5s2-1.12 2-2.5z"/><path d="M13 11H7m6-4H7m3 8H7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RedditFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#akarIconsRedditFill0)"><path fill="currentColor" fill-rule="evenodd" d="M24 12c0 6.627-5.373 12-12 12S0 18.627 0 12S5.373 0 12 0s12 5.373 12 12m-4.312-.942c.194.277.304.604.316.942a1.751 1.751 0 0 1-.972 1.596c.014.176.014.352 0 .528c0 2.688-3.132 4.872-6.996 4.872c-3.864 0-6.996-2.184-6.996-4.872a3.444 3.444 0 0 1 0-.528a1.75 1.75 0 1 1 1.932-2.868a8.568 8.568 0 0 1 4.68-1.476l.888-4.164a.372.372 0 0 1 .444-.288l2.94.588a1.2 1.2 0 1 1-.156.732L13.2 5.58l-.78 3.744a8.544 8.544 0 0 1 4.62 1.476a1.751 1.751 0 0 1 2.648.258M8.206 12.533a1.2 1.2 0 1 0 1.996 1.334a1.2 1.2 0 0 0-1.996-1.334m3.806 4.891c1.065.044 2.113-.234 2.964-.876a.335.335 0 1 0-.468-.48A3.936 3.936 0 0 1 12 16.8a3.924 3.924 0 0 1-2.496-.756a.324.324 0 0 0-.456.456a4.608 4.608 0 0 0 2.964.924m2.081-3.178c.198.132.418.25.655.25a1.199 1.199 0 0 0 1.212-1.248a1.2 1.2 0 1 0-1.867.998" clip-rule="evenodd"/></g><defs><clipPath id="akarIconsRedditFill0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reduce(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14 10l7-7m-7 7h6m-6 0V4M3 21l7-7m0 0v6m0-6H4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReduxFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#akarIconsReduxFill0)"><path fill="currentColor" d="M16.63 16.563c.885-.092 1.557-.863 1.527-1.788a1.723 1.723 0 0 0-1.71-1.665h-.062c-.947.03-1.68.832-1.65 1.788c.032.463.215.863.49 1.14c-1.039 2.067-2.627 3.577-5.01 4.841c-1.618.864-3.298 1.172-4.977.956c-1.375-.185-2.444-.802-3.116-1.819c-.977-1.51-1.068-3.145-.244-4.779c.58-1.171 1.497-2.035 2.077-2.466a16.987 16.987 0 0 1-.397-1.573C-.871 14.436-.412 18.814.93 20.88c1.008 1.542 3.054 2.498 5.315 2.498a7.45 7.45 0 0 0 1.832-.216c3.91-.77 6.872-3.114 8.552-6.598m5.375-3.823c-2.321-2.744-5.742-4.255-9.651-4.255h-.489a1.677 1.677 0 0 0-1.496-.925h-.062c-.946.031-1.68.833-1.649 1.789c.03.925.794 1.664 1.71 1.664h.062a1.721 1.721 0 0 0 1.496-1.048h.55c2.321 0 4.52.678 6.505 2.004c1.527 1.018 2.627 2.343 3.237 3.947c.52 1.294.49 2.558-.06 3.638c-.856 1.634-2.291 2.528-4.185 2.528c-1.221 0-2.382-.37-2.993-.648a18.07 18.07 0 0 1-1.374 1.11c1.313.617 2.657.956 3.94.956c2.932 0 5.1-1.634 5.925-3.268c.885-1.788.824-4.871-1.466-7.492M6.49 17.087c.03.925.794 1.665 1.71 1.665h.061c.947-.03 1.68-.832 1.65-1.788a1.723 1.723 0 0 0-1.71-1.665h-.062a.53.53 0 0 0-.214.03c-1.252-2.096-1.771-4.377-1.588-6.844c.122-1.85.733-3.453 1.802-4.779c.886-1.14 2.596-1.695 3.757-1.726c3.237-.062 4.611 4.008 4.703 5.642c.397.092 1.069.308 1.527.462C17.759 3.09 14.706.5 11.773.5C9.025.5 6.49 2.504 5.482 5.464c-1.405 3.946-.489 7.738 1.222 10.729c-.153.216-.245.555-.214.894"/></g><defs><clipPath id="akarIconsReduxFill0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reply(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><path d="M2 10.981L8.973 2v4.99c11.952 0 13.316 9.688 12.984 15.01l-.007-.041c-.502-2.685-.712-6.986-12.977-6.986v4.99L2 10.98z" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ribbon(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 5v14.586c0 .89 1.077 1.337 1.707.707L12 14l6.293 6.293c.63.63 1.707.184 1.707-.707V5a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RockOn(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="m19 16l.87-11.735a2.102 2.102 0 0 0-4.181-.433L15 9m-7 2l-.713-4.279a2.06 2.06 0 0 0-4.083.525L4 16m8-3v-1.5a2 2 0 1 0-4 0V15"/><path stroke-linecap="round" stroke-linejoin="round" d="M16 13v-2a2 2 0 1 0-4 0v2"/><path d="M19 16c-.536 4-3.358 6-7.5 6C7.358 22 4 20 4 16"/><path stroke-linecap="round" stroke-linejoin="round" d="M13.692 17H11a2 2 0 1 1 0-4h4c2.21 0 4.5 2 3.5 5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rss(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 10.419c6.068-.32 9.9 3.513 9.582 9.581"/><circle cx="5" cy="19" r="1"/><path d="M4 4.03C14.114 3.5 20.501 9.887 19.97 20"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SassFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#akarIconsSassFill0)"><path fill="currentColor" d="M12 0c6.6 0 12 5.4 12 12s-5.4 12-12 12S0 18.6 0 12S5.4 0 12 0M9.6 15.975c.15.675.15 1.275 0 1.8L9.525 18c0 .075-.075.15-.075.15c-.15.3-.3.525-.525.825c-.675.75-1.65 1.05-2.1.825c-.45-.225-.225-1.35.6-2.175c.9-.9 2.1-1.5 2.1-1.5zM19.5 5.1c-.525-2.1-4.05-2.85-7.425-1.65c-1.95.75-4.125 1.875-5.7 3.3c-1.875 1.725-2.1 3.225-2.025 3.825C4.8 12.75 7.8 14.25 9.075 15.3C8.7 15.45 6 16.8 5.4 18.225c-.675 1.5.075 2.55.6 2.625c1.575.45 3.225-.375 4.05-1.65c.825-1.275.75-2.85.375-3.675a3.532 3.532 0 0 1 1.8-.075c2.1.225 2.55 1.575 2.4 2.1c-.15.525-.525.825-.675.975c-.15.075-.225.15-.15.15c0 .075.075.075.225.075c.15 0 1.125-.45 1.125-1.5c.075-1.275-1.2-2.7-3.375-2.7c-.9 0-1.5.075-1.875.225c0-.075-.075-.075-.075-.075c-1.35-1.425-3.825-2.475-3.75-4.425c0-.675.3-2.55 4.8-4.8c3.675-1.875 6.675-1.35 7.2-.225c.75 1.575-1.575 4.575-5.4 5.025c-1.5.15-2.25-.375-2.4-.6c-.225-.225-.225-.225-.3-.225c-.15.075-.075.225 0 .375c.15.3.6.825 1.425 1.125c.675.225 2.4.375 4.5-.45c2.325-.9 4.125-3.375 3.6-5.475z"/></g><defs><clipPath id="akarIconsSassFill0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Save(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M7 21h10a4 4 0 0 0 4-4V7.414a1 1 0 0 0-.293-.707l-3.414-3.414A1 1 0 0 0 16.586 3H7a4 4 0 0 0-4 4v10a4 4 0 0 0 4 4"/><path d="M9 3h6v3a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1zm8 18v-7a1 1 0 0 0-1-1H8a1 1 0 0 0-1 1v7"/><path stroke-linecap="round" d="M11 17h2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Schedule(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 20H6a4 4 0 0 1-4-4V7a4 4 0 0 1 4-4h11a4 4 0 0 1 4 4v3M8 2v2m7-2v2M2 8h19m-2.5 7.643l-1.5 1.5"/><circle cx="17" cy="17" r="5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Scissor(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.252 18.459C7.462 19.764 7.107 21 5.7 21C4.209 21 3 19.757 3 18.223s1.209-2.778 2.7-2.778c1.4 0 2.55 1.095 2.686 2.498a.846.846 0 0 1-.134.515m0 0l1.948-3.476m5.548 3.476C16.538 19.764 16.893 21 18.3 21c1.491 0 2.7-1.243 2.7-2.777s-1.209-2.778-2.7-2.778c-1.4 0-2.55 1.095-2.687 2.498c-.017.182.04.36.135.515m0 0L7.093 3.346a.659.659 0 0 0-1.1-.081c-1.704 2.19-1.534 5.35.395 7.333zm-3.797-6.63l4.953-8.494a.66.66 0 0 1 1.098-.076c1.707 2.194 1.537 5.358-.395 7.345L16.5 11.742"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Search(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="m21 21l-4.486-4.494M19 10.5a8.5 8.5 0 1 1-17 0a8.5 8.5 0 0 1 17 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Send(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.912 12H4L2.023 4.135A.662.662 0 0 1 2 3.995c-.022-.721.772-1.221 1.46-.891L22 12L3.46 20.896c-.68.327-1.464-.159-1.46-.867a.66.66 0 0 1 .033-.186L3.5 15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SettingsHorizontal(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M3 5h4m14 0H11m-8 7h12m6 0h-2M3 19h2m16 0H9"/><circle cx="9" cy="5" r="2"/><circle cx="17" cy="12" r="2"/><circle cx="7" cy="19" r="2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SettingsVertical(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M19 3v4m0 14V11m-7-8v12m0 6v-2M5 3v2m0 16V9"/><circle cx="19" cy="9" r="2" transform="rotate(90 19 9)"/><circle cx="12" cy="17" r="2" transform="rotate(90 12 17)"/><circle cx="5" cy="7" r="2" transform="rotate(90 5 7)"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareArrow(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><path d="M22 10.981L15.027 2v4.99C3.075 6.99 1.711 16.678 2.043 22l.007-.041c.502-2.685.712-6.986 12.977-6.986v4.99L22 10.98z" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareBox(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3v12m0-12L8 7m4-4l4 4M2 17l.621 2.485A2 2 0 0 0 4.561 21h14.878a2 2 0 0 0 1.94-1.515L22 17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shield(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m13.147 21.197l1.67-1.168a13.393 13.393 0 0 0 5.447-13.624a.837.837 0 0 0-.453-.586L12 2L4.19 5.819a.837.837 0 0 0-.454.586a13.393 13.393 0 0 0 5.448 13.624l1.67 1.168a2 2 0 0 0 2.293 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShippingBoxOne(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path stroke-linecap="round" d="M11.029 2.54a2 2 0 0 1 1.942 0l7.515 4.174a1 1 0 0 1 .514.874v8.235a2 2 0 0 1-1.029 1.749l-7 3.888a2 2 0 0 1-1.942 0l-7-3.889A2 2 0 0 1 3 15.824V7.588a1 1 0 0 1 .514-.874z"/><path d="m3 7l9 5m0 0l9-5m-9 5v10"/><path stroke-linecap="round" d="m7.5 9.5l9-5M6 12.328L9 14"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShippingBoxTwo(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path stroke-linecap="round" d="M11.029 2.54a2 2 0 0 1 1.942 0l7.515 4.174a1 1 0 0 1 .514.874v8.235a2 2 0 0 1-1.029 1.749l-7 3.888a2 2 0 0 1-1.942 0l-7-3.889A2 2 0 0 1 3 15.824V7.588a1 1 0 0 1 .514-.874z"/><path stroke-linecap="round" d="m7.5 4.5l9 5V13M6 12.328L9 14"/><path d="m3 7l9 5m0 0l9-5m-9 5v10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShippingBoxVone(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path stroke-linecap="round" d="M11.029 2.54a2 2 0 0 1 1.942 0l7.515 4.174a1 1 0 0 1 .514.874v8.235a2 2 0 0 1-1.029 1.749l-7 3.888a2 2 0 0 1-1.942 0l-7-3.889A2 2 0 0 1 3 15.824V7.588a1 1 0 0 1 .514-.874z"/><path d="m3 7l9 5m0 0l9-5m-9 5v9.5"/><path stroke-linecap="round" d="m7.5 9.5l9-5M6 12.328L9 14"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShippingBoxVtwo(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path stroke-linecap="round" d="M11.029 2.54a2 2 0 0 1 1.942 0l7.515 4.174a1 1 0 0 1 .514.874v8.235a2 2 0 0 1-1.029 1.749l-7 3.888a2 2 0 0 1-1.942 0l-7-3.889A2 2 0 0 1 3 15.824V7.588a1 1 0 0 1 .514-.874z"/><path stroke-linecap="round" d="m7.5 4.5l9 5V13M6 12.328L9 14"/><path d="m3 7l9 5m0 0l9-5m-9 5v9.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBag(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3.977 9.84A2 2 0 0 1 5.971 8h12.058a2 2 0 0 1 1.994 1.84l.803 10A2 2 0 0 1 18.833 22H5.167a2 2 0 0 1-1.993-2.16z"/><path d="M16 11V6a4 4 0 0 0-4-4v0a4 4 0 0 0-4 4v5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SidebarLeft(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><rect width="20" height="18" x="2" y="3" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path d="M9 3v18"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SidebarRight(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><rect width="20" height="18" x="2" y="3" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path d="M15 3v18"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignOut(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 12h9m0 0l-3.333-4M22 12l-3.333 4M14 7V5.174a2 2 0 0 0-2.166-1.993l-8 .666A2 2 0 0 0 2 5.84v12.32a2 2 0 0 0 1.834 1.993l8 .667A2 2 0 0 0 14 18.826V17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SlackFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.066 1a2.2 2.2 0 1 0 .001 4.4h2.2V3.2a2.202 2.202 0 0 0-2.2-2.2m0 5.867H3.2a2.2 2.2 0 0 0 0 4.4h5.866a2.2 2.2 0 1 0 0-4.4M23 9.066a2.2 2.2 0 0 0-4.4 0v2.2h2.2a2.2 2.2 0 0 0 2.2-2.2m-5.867 0V3.2a2.2 2.2 0 0 0-4.4 0v5.866a2.2 2.2 0 1 0 4.4 0M14.933 23a2.2 2.2 0 1 0 0-4.4h-2.2v2.2a2.201 2.201 0 0 0 2.2 2.2m0-5.868H20.8a2.2 2.2 0 0 0 0-4.4h-5.866a2.2 2.2 0 0 0-.001 4.4M1 14.933a2.2 2.2 0 0 0 4.4 0v-2.2H3.2a2.2 2.2 0 0 0-2.2 2.2m5.867 0v5.866a2.2 2.2 0 0 0 4.4.001v-5.866a2.2 2.2 0 0 0-4.4-.001" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Slice(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M13 21H3l10-10z"/><path d="M5 19L21 3v4l-8 8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SnapchatFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#akarIconsSnapchatFill0)"><path fill="currentColor" d="M5.829 4.533c-.6 1.344-.363 3.752-.267 5.436c-.648.359-1.48-.271-1.951-.271c-.49 0-1.075.322-1.167.802c-.066.346.089.85 1.201 1.289c.43.17 1.453.37 1.69.928c.333.784-1.71 4.403-4.918 4.931a.498.498 0 0 0-.416.519c.056.975 2.242 1.357 3.211 1.507c.099.134.179.7.306 1.131c.057.193.204.424.582.424c.493 0 1.312-.38 2.738-.144c1.398.233 2.712 2.215 5.235 2.215c2.345 0 3.744-1.991 5.09-2.215c.779-.129 1.448-.088 2.196.058c.515.101.977.157 1.124-.349c.129-.437.208-.992.305-1.123c.96-.149 3.156-.53 3.211-1.505a.498.498 0 0 0-.416-.519c-3.154-.52-5.259-4.128-4.918-4.931c.236-.557 1.252-.755 1.69-.928c.814-.321 1.222-.716 1.213-1.173c-.011-.585-.715-.934-1.233-.934c-.527 0-1.284.624-1.897.286c.096-1.698.332-4.095-.267-5.438C17.036 1.986 14.511.7 11.987.7C9.479.7 6.973 1.968 5.829 4.533"/></g><defs><clipPath id="akarIconsSnapchatFill0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sort(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M3 6h18M6 12h12m-9 6h6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SoundDown(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M22 12h-6"/><path stroke-linejoin="round" d="M2 14.959V9.04C2 8.466 2.448 8 3 8h3.586a.98.98 0 0 0 .707-.305l3-3.388c.63-.656 1.707-.191 1.707.736v13.914c0 .934-1.09 1.395-1.716.726l-2.99-3.369A.98.98 0 0 0 6.578 16H3c-.552 0-1-.466-1-1.041"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SoundOff(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="m22 15l-6-6m6 0l-6 6"/><path stroke-linejoin="round" d="M2 14.959V9.04C2 8.466 2.448 8 3 8h3.586a.98.98 0 0 0 .707-.305l3-3.388c.63-.656 1.707-.191 1.707.736v13.914c0 .934-1.09 1.395-1.716.726l-2.99-3.369A.98.98 0 0 0 6.578 16H3c-.552 0-1-.466-1-1.041"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SoundOn(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 14.959V9.04C2 8.466 2.448 8 3 8h3.586a.98.98 0 0 0 .707-.305l3-3.388c.63-.656 1.707-.191 1.707.736v13.914c0 .934-1.09 1.395-1.716.726l-2.99-3.369A.98.98 0 0 0 6.578 16H3c-.552 0-1-.466-1-1.041M16 8.5c1.333 1.778 1.333 5.222 0 7M19 5c3.988 3.808 4.012 10.217 0 14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SoundUp(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M19 9v3m0 0v3m0-3h3m-3 0h-3"/><path stroke-linejoin="round" d="M2 14.959V9.04C2 8.466 2.448 8 3 8h3.586a.98.98 0 0 0 .707-.305l3-3.388c.63-.656 1.707-.191 1.707.736v13.914c0 .934-1.09 1.395-1.716.726l-2.99-3.369A.98.98 0 0 0 6.578 16H3c-.552 0-1-.466-1-1.041"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SoundcloudFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.16 16.466c.049 0 .09-.039.098-.093l.27-2.022l-.27-2.069a.103.103 0 0 0-.099-.093c-.05 0-.094.04-.1.093l-.236 2.069l.236 2.021c.006.055.05.094.1.094m-.887-.769c.048 0 .088-.036.095-.09l.209-1.256l-.209-1.28c-.007-.053-.047-.09-.095-.09c-.051 0-.09.037-.098.09L0 14.351l.174 1.256c.008.053.047.09.098.09m1.948-3.8a.122.122 0 0 0-.12-.114a.12.12 0 0 0-.119.114l-.224 2.454l.224 2.364a.12.12 0 0 0 .12.112a.121.121 0 0 0 .12-.113l.254-2.363zm.832 5.026a.143.143 0 0 0 .14-.133l.241-2.439l-.24-2.522a.143.143 0 0 0-.141-.132a.14.14 0 0 0-.14.133l-.213 2.521l.212 2.439a.14.14 0 0 0 .141.133m.958.039a.162.162 0 0 0 .162-.152l.226-2.459l-.226-2.34a.162.162 0 0 0-.162-.151a.16.16 0 0 0-.16.152l-.2 2.34l.2 2.458a.16.16 0 0 0 .16.152m1.36-2.61l-.212-3.805a.184.184 0 0 0-.182-.173a.183.183 0 0 0-.182.173l-.188 3.805l.188 2.458a.183.183 0 0 0 .364 0zm.581 2.635a.201.201 0 0 0 .201-.192v.002l.199-2.444l-.199-4.676a.203.203 0 0 0-.405 0l-.174 4.676l.175 2.443a.201.201 0 0 0 .203.19m.98-7.91a.222.222 0 0 0-.223.212l-.162 5.065l.162 2.418a.221.221 0 0 0 .223.211a.22.22 0 0 0 .223-.211l.185-2.418l-.185-5.065a.22.22 0 0 0-.223-.212m.989 7.911a.24.24 0 0 0 .244-.232v.002l.17-2.404l-.17-5.235a.24.24 0 0 0-.243-.232a.238.238 0 0 0-.243.232l-.153 5.235l.153 2.404c.002.129.11.23.243.23m.997-.002a.26.26 0 0 0 .263-.252v.002l.157-2.381l-.157-5.103a.26.26 0 0 0-.263-.25a.26.26 0 0 0-.264.25l-.138 5.103l.139 2.38c.003.14.119.25.263.25m1.431-2.63l-.142-4.917a.28.28 0 0 0-.284-.27a.28.28 0 0 0-.285.271l-.127 4.916l.127 2.366a.28.28 0 0 0 .285.27a.28.28 0 0 0 .284-.273v.003zm.586 2.64c.165 0 .301-.13.304-.29l.129-2.349l-.129-5.85a.301.301 0 0 0-.304-.291a.303.303 0 0 0-.305.291l-.115 5.848l.115 2.352c.003.158.14.289.305.289m1.009-9.33a.322.322 0 0 0-.327.31l-.133 6.382l.134 2.315a.32.32 0 0 0 .325.308a.32.32 0 0 0 .324-.311v.003l.146-2.315l-.146-6.381a.32.32 0 0 0-.323-.311m.922 9.332l8.182.004C22.678 17 24 15.732 24 14.167c0-1.564-1.322-2.832-2.953-2.832c-.404 0-.79.079-1.142.22C19.673 9.003 17.44 7 14.718 7c-.665 0-1.314.126-1.887.339c-.223.083-.283.168-.285.333v8.989a.349.349 0 0 0 .32.335"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sparkles(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 2l.19.94a4 4 0 0 0 2.57 2.974L8 6l-.24.086A4 4 0 0 0 5.19 9.06L5 10l-.19-.94a4 4 0 0 0-2.57-2.974L2 6l.24-.086A4 4 0 0 0 4.81 2.94zm3 14l.23 1.276a2 2 0 0 0 1.219 1.501L10 19l-.551.223a2 2 0 0 0-1.22 1.5L8 22l-.23-1.276a2 2 0 0 0-1.219-1.501L6 19l.551-.223a2 2 0 0 0 1.22-1.5zm8-13l.556 2.654a8 8 0 0 0 5.213 5.92L23 12l-1.231.426a8 8 0 0 0-5.213 5.92L16 21l-.556-2.654a8 8 0 0 0-5.213-5.92L9 12l1.231-.426a8 8 0 0 0 5.213-5.92z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpotifyFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#akarIconsSpotifyFill0)"><path fill="currentColor" d="M11.995 0C5.381 0 0 5.382 0 11.996C0 18.616 5.381 24 11.995 24C18.615 24 24 18.615 24 11.996C24 5.382 18.615 0 11.995 0M5.908 16.404a14.548 14.548 0 0 1 4.238-.638c2.414 0 4.797.612 6.892 1.77c.125.068.238.292.29.572c.05.28.03.567-.052.716a.61.61 0 0 1-.834.24A13.107 13.107 0 0 0 6.277 18.03a.61.61 0 0 1-.771-.402c-.107-.35.114-1.13.402-1.224m-.523-4.42a18.154 18.154 0 0 1 4.76-.635c2.894 0 5.767.7 8.31 2.026c.179.09.31.244.37.432a.747.747 0 0 1-.052.578c-.227.444-.493.743-.66.743a.769.769 0 0 1-.35-.086a16.33 16.33 0 0 0-7.617-1.854a16.34 16.34 0 0 0-4.366.585a.749.749 0 0 1-.92-.525c-.112-.422.145-1.16.525-1.264M5.25 9.098a.88.88 0 0 1-1.073-.641c-.123-.498.188-1.076.64-1.19a22.365 22.365 0 0 1 5.328-.649c3.45 0 6.756.776 9.824 2.307a.888.888 0 0 1 .4 1.19c-.143.288-.453.598-.795.598a.924.924 0 0 1-.388-.087A20.026 20.026 0 0 0 10.145 8.5c-1.635 0-3.282.201-4.895.598"/></g><defs><clipPath id="akarIconsSpotifyFill0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Square(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<rect width="18" height="18" x="3" y="3" fill="none" stroke="currentColor" stroke-width="2" rx="2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<rect width="20" height="20" x="2" y="2" fill="currentColor" rx="2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StackOverflowFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.372 20.222v-5.358h1.79V22H4v-7.136h1.79v5.358z"/><path fill="currentColor" d="m7.768 14.356l8.79 1.824l.372-1.755L8.14 12.6zm1.162-4.157l8.14 3.764l.744-1.617l-8.14-3.787zm2.256-3.973l6.907 5.705l1.14-1.363l-6.907-5.704zM15.651 2L14.21 3.062l5.35 7.16L21 9.159zm-8.07 16.42h8.977v-1.778H7.581z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Star(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.074 2.633c.32-.844 1.531-.844 1.852 0l2.07 5.734a.99.99 0 0 0 .926.633h5.087c.94 0 1.35 1.17.611 1.743L18 14a.968.968 0 0 0-.322 1.092L19 20.695c.322.9-.72 1.673-1.508 1.119l-4.917-3.12a1 1 0 0 0-1.15 0l-4.917 3.12c-.787.554-1.83-.22-1.508-1.119l1.322-5.603A.968.968 0 0 0 6 14l-3.62-3.257C1.64 10.17 2.052 9 2.99 9h5.087a.989.989 0 0 0 .926-.633z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StatisticDown(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-miterlimit="5.759" d="M3 3v16a2 2 0 0 0 2 2h16"/><path stroke-miterlimit="5.759" d="m7 9l4 4l4-4l6 6"/><path d="M18 15h3v-3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StatisticUp(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path stroke-miterlimit="5.759" d="M3 3v16a2 2 0 0 0 2 2h16"/><path stroke-miterlimit="5.759" d="m7 14l4-4l4 4l6-6"/><path d="M18 8h3v3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stop(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" d="M8 2L2 8.156V16l6 6h8l6-6V8.156L16 2z"/><path d="M16 12H8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 1a1 1 0 0 0-.716.302l-6 6.156A1 1 0 0 0 1 8.156V16a1 1 0 0 0 .293.707l6 6A1 1 0 0 0 8 23h8a1 1 0 0 0 .707-.293l6-6A1 1 0 0 0 23 16V8.156a1 1 0 0 0-.284-.698l-6-6.156A1 1 0 0 0 16 1zm0 10a1 1 0 1 0 0 2h8a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sun(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M12 3V2m0 20v-1m9-9h1M2 12h1m15.5-6.5L20 4M4 20l1.5-1.5M4 4l1.5 1.5m13 13L20 20"/><circle cx="12" cy="12" r="4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SunFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 7a5 5 0 1 0 0 10a5 5 0 0 0 0-10"/><path fill="currentColor" fill-rule="evenodd" d="M12 1a1 1 0 0 1 1 1v1a1 1 0 1 1-2 0V2a1 1 0 0 1 1-1M3.293 3.293a1 1 0 0 1 1.414 0l1.5 1.5a1 1 0 0 1-1.414 1.414l-1.5-1.5a1 1 0 0 1 0-1.414m17.414 0a1 1 0 0 1 0 1.414l-1.5 1.5a1 1 0 1 1-1.414-1.414l1.5-1.5a1 1 0 0 1 1.414 0M1 12a1 1 0 0 1 1-1h1a1 1 0 1 1 0 2H2a1 1 0 0 1-1-1m19 0a1 1 0 0 1 1-1h1a1 1 0 1 1 0 2h-1a1 1 0 0 1-1-1M6.207 17.793a1 1 0 0 1 0 1.414l-1.5 1.5a1 1 0 0 1-1.414-1.414l1.5-1.5a1 1 0 0 1 1.414 0m11.586 0a1 1 0 0 1 1.414 0l1.5 1.5a1 1 0 0 1-1.414 1.414l-1.5-1.5a1 1 0 0 1 0-1.414M12 20a1 1 0 0 1 1 1v1a1 1 0 1 1-2 0v-1a1 1 0 0 1 1-1" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sword(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m18 3l4-1l-1 4l-10 10l-2.5-.5L8 13zM2 20l2 2m1-8l1 4l4 1m-4-1l-3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TabletDevice(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="16" height="20" x="4" y="2" rx="2"/><path d="M11 18h2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tag(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" d="M15.244 21.366a2.164 2.164 0 0 1-3.061 0l-8.549-8.549A2.164 2.164 0 0 1 3 11.287V5.164C3 3.97 3.97 3 5.164 3h6.123c.573 0 1.124.228 1.53.634l8.549 8.549a2.164 2.164 0 0 1 0 3.061z"/><path d="M6.5 6.5L7 7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TelegramFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#akarIconsTelegramFill0)"><path fill="currentColor" fill-rule="evenodd" d="M24 12c0 6.627-5.373 12-12 12S0 18.627 0 12S5.373 0 12 0s12 5.373 12 12M12.43 8.859c-1.167.485-3.5 1.49-6.998 3.014c-.568.226-.866.447-.893.663c-.046.366.412.51 1.034.705c.085.027.173.054.263.084c.613.199 1.437.432 1.865.441c.389.008.823-.152 1.302-.48c3.268-2.207 4.955-3.322 5.061-3.346c.075-.017.179-.039.249.024c.07.062.063.18.056.212c-.046.193-1.84 1.862-2.77 2.726c-.29.269-.495.46-.537.504c-.094.097-.19.19-.282.279c-.57.548-.996.96.024 1.632c.49.323.882.59 1.273.856c.427.291.853.581 1.405.943c.14.092.274.187.405.28c.497.355.944.673 1.496.623c.32-.03.652-.331.82-1.23c.397-2.126 1.179-6.73 1.36-8.628a2.111 2.111 0 0 0-.02-.472a.506.506 0 0 0-.172-.325c-.143-.117-.365-.142-.465-.14c-.451.008-1.143.249-4.476 1.635" clip-rule="evenodd"/></g><defs><clipPath id="akarIconsTelegramFill0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Telescope(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m7 21l6-8l8 8M5.228 7.303l13.532-5.06a1 1 0 0 1 1.108.285l.19.22A8 8 0 0 1 22 7.973v.616a1 1 0 0 1-.65.937l-13.536 5.06z"/><path d="M2.66 11.371a2 2 0 0 1 1.193-2.545l1.694-.624l1.944 5.473l-1.64.612a2 2 0 0 1-2.585-1.205zM13 13v9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tetragon(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="2" d="M3.575 13.388a1.962 1.962 0 0 1 0-2.776l7.037-7.037a1.962 1.962 0 0 1 2.776 0l7.037 7.037a1.963 1.963 0 0 1 0 2.776l-7.037 7.037a1.963 1.963 0 0 1-2.776 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TetragonFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.61 2.576a1.966 1.966 0 0 1 2.78 0l8.034 8.034a1.966 1.966 0 0 1 0 2.78l-8.034 8.034a1.966 1.966 0 0 1-2.78 0L2.576 13.39a1.966 1.966 0 0 1 0-2.78z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextAlignCenter(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M3 6h18M7 12h10M5 18h14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextAlignJustified(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M3 6h18M3 12h18M3 18h18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextAlignLeft(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M3 6h18M3 12h10M3 18h15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextAlignRight(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M3 6h18m-10 6h10M6 18h15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThreadsFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.692 11.124a8.547 8.547 0 0 0-.315-.143c-.185-3.414-2.05-5.368-5.182-5.388h-.042c-1.873 0-3.431.8-4.39 2.255l1.722 1.181c.716-1.087 1.84-1.318 2.669-1.318h.028c1.031.006 1.81.306 2.313.89c.367.426.612 1.015.733 1.757a13.176 13.176 0 0 0-2.96-.143c-2.977.172-4.892 1.909-4.763 4.322c.065 1.223.675 2.277 1.717 2.964c.88.582 2.015.866 3.194.802c1.558-.085 2.78-.68 3.632-1.766c.647-.825 1.056-1.894 1.237-3.241c.742.448 1.292 1.037 1.596 1.745c.516 1.205.546 3.184-1.068 4.797c-1.415 1.414-3.116 2.025-5.686 2.044c-2.851-.02-5.008-.935-6.41-2.717c-1.313-1.67-1.991-4.08-2.016-7.165c.025-3.085.703-5.496 2.016-7.165c1.402-1.782 3.558-2.696 6.41-2.717c2.871.02 5.065.94 6.521 2.73c.714.879 1.252 1.983 1.607 3.27l2.018-.538c-.43-1.585-1.107-2.95-2.027-4.083C18.38 1.2 15.65.024 12.134 0h-.014C8.61.024 5.912 1.205 4.099 3.51c-1.614 2.05-2.446 4.904-2.474 8.482v.016c.028 3.578.86 6.431 2.474 8.482c1.813 2.305 4.511 3.486 8.02 3.51h.015c3.12-.022 5.319-.838 7.13-2.649c2.371-2.368 2.3-5.336 1.518-7.158c-.56-1.307-1.629-2.368-3.09-3.07m-5.387 5.064c-1.305.074-2.66-.512-2.728-1.766c-.05-.93.662-1.969 2.808-2.092c.246-.015.487-.021.724-.021c.78 0 1.508.075 2.171.22c-.247 3.088-1.697 3.59-2.975 3.66"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThreeLineHorizontal(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M5 6h14M5 12h14M5 18h14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThreeLineVertical(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M12 19V5m6 14V5M6 19V5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsDown(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m12.395 18.218l-.23 2.369c-.091.952-.98 1.598-1.878 1.366c-1.351-.35-2.3-1.605-2.3-3.044v-3.035c0-.675 0-1.013-.146-1.26a1.018 1.018 0 0 0-.333-.345c-.24-.151-.567-.151-1.22-.151h-.396c-1.703 0-2.554 0-3.078-.39a2.073 2.073 0 0 1-.78-1.208c-.146-.65.181-1.463.836-3.087l.327-.81c.188-.468.265-.975.225-1.48c-.232-2.874 2.047-5.295 4.833-5.135l10.424.598c1.139.065 1.708.098 2.222.553c.515.455.612.924.805 1.861a14.317 14.317 0 0 1-.055 6.037c-.283 1.249-1.475 1.92-2.706 1.76c-3.264-.42-6.223 2.019-6.55 5.4"/><path d="m17 12.5l.137-.457c.887-2.956.84-6.115-.137-9.043"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsUp(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m11.605 5.782l.23-2.369c.091-.952.98-1.598 1.878-1.366c1.351.35 2.3 1.605 2.3 3.044v3.035c0 .675 0 1.013.146 1.26c.083.141.197.26.333.345c.24.151.567.151 1.22.151h.396c1.703 0 2.554 0 3.078.39c.393.293.67.722.78 1.208c.146.65-.181 1.463-.836 3.087l-.326.81a3.261 3.261 0 0 0-.226 1.48c.232 2.874-2.047 5.295-4.833 5.136l-10.424-.599c-1.139-.065-1.708-.098-2.222-.553c-.515-.455-.612-.924-.805-1.861a14.324 14.324 0 0 1 .055-6.037c.283-1.248 1.475-1.92 2.706-1.76c3.264.42 6.223-2.019 6.55-5.4"/><path d="m7 11.5l-.137.457A14.983 14.983 0 0 0 7 21"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Thunder(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.684 3.603c.521-.659.03-1.603-.836-1.603h-6.716a1.06 1.06 0 0 0-.909.502l-5.082 8.456c-.401.666.103 1.497.908 1.497h3.429l-3.23 8.065c-.467 1.02.795 1.953 1.643 1.215L20 9.331h-6.849z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ticket(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.707 3.293c-.39.39-.369 1.021-.138 1.523a2.83 2.83 0 0 1-3.753 3.753c-.502-.23-1.133-.252-1.523.138l-.586.586a1 1 0 0 0 0 1.414l10.586 10.586a1 1 0 0 0 1.414 0l.586-.586c.39-.39.369-1.021.138-1.523a2.83 2.83 0 0 1 3.753-3.753c.502.23 1.133.252 1.523-.138l.586-.586a1 1 0 0 0 0-1.414L10.707 2.707a1 1 0 0 0-1.414 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TiktokFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.321 5.562a5.124 5.124 0 0 1-.443-.258a6.228 6.228 0 0 1-1.137-.966c-.849-.971-1.166-1.956-1.282-2.645h.004c-.097-.573-.057-.943-.05-.943h-3.865v14.943c0 .2 0 .399-.008.595c0 .024-.003.046-.004.073c0 .01 0 .022-.003.033v.009a3.28 3.28 0 0 1-1.65 2.604a3.226 3.226 0 0 1-1.6.422c-1.8 0-3.26-1.468-3.26-3.281c0-1.814 1.46-3.282 3.26-3.282c.341 0 .68.054 1.004.16l.005-3.936a7.178 7.178 0 0 0-5.532 1.62a7.583 7.583 0 0 0-1.655 2.04c-.163.281-.779 1.412-.853 3.246c-.047 1.04.266 2.12.415 2.565v.01c.093.262.457 1.158 1.049 1.913a7.856 7.856 0 0 0 1.674 1.58v-.01l.009.01c1.87 1.27 3.945 1.187 3.945 1.187c.359-.015 1.562 0 2.928-.647c1.515-.718 2.377-1.787 2.377-1.787a7.43 7.43 0 0 0 1.296-2.153c.35-.92.466-2.022.466-2.462V8.273c.047.028.672.441.672.441s.9.577 2.303.952c1.006.267 2.363.324 2.363.324V6.153c-.475.052-1.44-.098-2.429-.59"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToggleOff(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="10" x="2" y="7" rx="5"/><circle cx="7" cy="12" r="5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToggleOffFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7 7a5 5 0 0 0 0 10h10a5 5 0 0 0 0-10zm0 2.5a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToggleOn(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="10" x="2" y="7" rx="5"/><circle cx="17" cy="12" r="5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToggleOnFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7 7a5 5 0 0 0 0 10h10a5 5 0 0 0 0-10zm10 2.5a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TogoCup(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 5.4A2.4 2.4 0 0 1 5.4 3h13.2A2.4 2.4 0 0 1 21 5.4v0a.6.6 0 0 1-.6.6H3.6a.6.6 0 0 1-.6-.6M5 6h14l-1.555 14.218A2 2 0 0 1 15.457 22H8.543a2 2 0 0 1-1.988-1.782z"/><path d="m6.313 18l-.875-8h13.125l-.875 8z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trash(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><path d="M3 4l.813 8.132c.125 1.243.346 2.475.662 3.684l.667 2.55a4 4 0 0 0 2.66 2.801l.567.18a12 12 0 0 0 7.262 0l.567-.18a4 4 0 0 0 2.66-2.8l.667-2.55c.316-1.21.538-2.442.662-3.685L21 4" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/><ellipse cx="12" cy="4" rx="9" ry="2" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrashBin(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m3 4l2.303 14.077a4 4 0 0 0 2.738 3.166l.328.104a12 12 0 0 0 7.262 0l.328-.104a4 4 0 0 0 2.738-3.166L21 4"/><ellipse cx="12" cy="4" rx="9" ry="2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrashCan(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16l-1.58 14.22A2 2 0 0 1 16.432 22H7.568a2 2 0 0 1-1.988-1.78zm3.345-2.853A2 2 0 0 1 9.154 2h5.692a2 2 0 0 1 1.81 1.147L18 6H6zM2 6h20m-12 5v5m4-5v5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Triangle(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="2" d="M10.293 4.793c.78-1.277 2.634-1.277 3.414 0l7.433 12.164C21.955 18.29 20.996 20 19.434 20H4.566c-1.562 0-2.52-1.71-1.706-3.043z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleAlert(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M12 9v5m0 3.5v.5"/><path stroke-linejoin="round" d="M2.232 19.016L10.35 3.052c.713-1.403 2.59-1.403 3.302 0l8.117 15.964C22.45 20.36 21.544 22 20.116 22H3.883c-1.427 0-2.334-1.64-1.65-2.984"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleAlertFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.543 2.598a2.821 2.821 0 0 0-5.086 0L1.341 18.563C.37 20.469 1.597 23 3.883 23h16.234c2.286 0 3.511-2.53 2.542-4.437zM12 8a1 1 0 0 1 1 1v5a1 1 0 1 1-2 0V9a1 1 0 0 1 1-1m0 8.5a1 1 0 0 1 1 1v.5a1 1 0 1 1-2 0v-.5a1 1 0 0 1 1-1" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleDown(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 17L6 9h12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleDownFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 8a1 1 0 0 0-.8 1.6l6 8a1 1 0 0 0 1.6 0l6-8A1 1 0 0 0 18 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.285 3.858c.777-1.294 2.653-1.294 3.43 0l8.468 14.113c.8 1.333-.16 3.029-1.715 3.029H3.532c-1.554 0-2.514-1.696-1.715-3.029z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleLeft(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 12l8-6v12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleLeftFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 6a1 1 0 0 0-1.6-.8l-8 6a1 1 0 0 0 0 1.6l8 6A1 1 0 0 0 16 18z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleRight(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 12L9 6v12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleRightFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 6a1 1 0 0 1 1.6-.8l8 6a1 1 0 0 1 0 1.6l-8 6A1 1 0 0 1 8 18z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleUp(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 7l-6 8h12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleUpFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 16a1 1 0 0 1-.8-1.6l6-8a1 1 0 0 1 1.6 0l6 8A1 1 0 0 1 18 16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trophy(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path stroke-linecap="round" d="M5 4a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v5a7 7 0 0 1-7 7v0a7 7 0 0 1-7-7zm4 18h6l-3-5z"/><path d="M5 4H4a2 2 0 0 0-2 2v1.239a4 4 0 0 0 2.128 3.535L5.5 11.5M19 4h1a2 2 0 0 1 2 2v.637a5 5 0 0 1-2.66 4.419l-.84.444"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Truck(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 17h6V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v10a1 1 0 0 0 1 1h1m18-1v-4a4 4 0 0 0-4-4h-2v9h5a1 1 0 0 0 1-1m-7 1a3 3 0 1 0 6 0zm7-3h-2"/><circle cx="7" cy="17" r="3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TumblrFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#akarIconsTumblrFill0)"><path fill="currentColor" d="M18.895 22.517c-.798.867-2.646 1.456-4.301 1.483h-.182c-5.557 0-6.766-4.164-6.766-6.594v-6.748H5.458a.454.454 0 0 1-.324-.137a.472.472 0 0 1-.134-.33V7.003a.81.81 0 0 1 .142-.458a.782.782 0 0 1 .376-.29c2.855-1.026 3.748-3.562 3.87-5.49c.035-.516.297-.765.738-.765H13.4a.451.451 0 0 1 .33.134a.468.468 0 0 1 .137.333V5.87h3.823c.121 0 .238.05.324.137a.472.472 0 0 1 .134.33v3.83a.472.472 0 0 1-.134.33a.454.454 0 0 1-.324.138h-3.84v6.245c0 1.568 1.015 2.001 1.64 2.001a4.537 4.537 0 0 0 1.488-.321a.973.973 0 0 1 .595-.106a.483.483 0 0 1 .34.37l1.012 3.014c.068.237.14.498-.03.68"/></g><defs><clipPath id="akarIconsTumblrFill0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwitchFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.547 1L1 4.776v15.433h5.5V23h3.093l2.922-2.791h4.47L23 14.462V1zm18.39 12.478L17.5 16.76H12l-2.922 2.791v-2.79h-4.64V2.97h16.499zM17.5 6.747v5.74h-2.063v-5.74zm-5.5 0v5.74H9.938v-5.74z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwitterFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23.643 4.937c-.835.37-1.732.62-2.675.733a4.67 4.67 0 0 0 2.048-2.578a9.3 9.3 0 0 1-2.958 1.13a4.66 4.66 0 0 0-7.938 4.25a13.229 13.229 0 0 1-9.602-4.868c-.4.69-.63 1.49-.63 2.342A4.66 4.66 0 0 0 3.96 9.824a4.647 4.647 0 0 1-2.11-.583v.06a4.66 4.66 0 0 0 3.737 4.568a4.692 4.692 0 0 1-2.104.08a4.661 4.661 0 0 0 4.352 3.234a9.348 9.348 0 0 1-5.786 1.995a9.5 9.5 0 0 1-1.112-.065a13.175 13.175 0 0 0 7.14 2.093c8.57 0 13.255-7.098 13.255-13.254c0-.2-.005-.402-.014-.602a9.47 9.47 0 0 0 2.323-2.41z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwoLineHorizontal(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M5 9h14M5 15h14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwoLineVertical(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M15 19V5M9 19V5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TypescriptFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#akarIconsTypescriptFill0)"><path fill="currentColor" d="M23.429 0H.57A.571.571 0 0 0 0 .571V23.43a.57.57 0 0 0 .571.571H23.43a.571.571 0 0 0 .571-.571V.57a.571.571 0 0 0-.572-.57m-9.143 12.826h-2.857v8.888H9.143v-8.888H6.286v-1.969h8zm.64 8.38v-2.375s1.298.978 2.855.978s1.497-1.018 1.497-1.158c0-1.477-4.412-1.477-4.412-4.751c0-4.452 6.429-2.695 6.429-2.695l-.08 2.116s-1.078-.719-2.296-.719c-1.218 0-1.657.58-1.657 1.198c0 1.597 4.452 1.438 4.452 4.652c0 4.95-6.788 2.755-6.788 2.755"/></g><defs><clipPath id="akarIconsTypescriptFill0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Umbrella(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4c-4.915 0-8.91 3.378-8.999 8.817a.18.18 0 0 0 .182.183a.188.188 0 0 0 .17-.11C3.876 11.767 4.782 11.5 6 11.5c1.185 0 1.964.227 2.456 1.302c.054.12.172.198.304.198a.366.366 0 0 0 .326-.224C9.56 11.729 10.901 11.5 12 11.5M12 4c4.916 0 8.91 3.378 8.999 8.817a.18.18 0 0 1-.182.183a.188.188 0 0 1-.17-.11c-.524-1.123-1.43-1.39-2.647-1.39c-1.185 0-1.964.227-2.456 1.302a.336.336 0 0 1-.304.198a.366.366 0 0 1-.326-.224C14.44 11.729 13.099 11.5 12 11.5M12 4V2m0 9.5V20a2 2 0 1 1-4 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnsplashFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.25 2h7.5v5.625h-7.5zM2 10.75h6.268v5.675h7.517V10.75H22V22H2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Utensils(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 8V3a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v5a4 4 0 0 0 4 4h1a4 4 0 0 0 4-4m4 8V2h3a4 4 0 0 1 4 4v10h-4m-3 0v5a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-5M5 12v9a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-9M5 6V2m3 4V2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VapeKit(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10a1 1 0 0 1 1-1h6v11a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1zm10 5a3 3 0 0 1 3-3h2a3 3 0 0 1 3 3v6h-8zm2-6h4v3h-4zM6 4h4v5H6zm2 0V2m9 7V6m0 11v-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VercelFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12 1l12 21H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VictoryHand(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="m12 11l-1-7.272c0-.466.185-.913.515-1.243c1.024-1.024 2.777-.44 2.982.994L16 10l1.508-6.328a1.682 1.682 0 0 1 3.276.73L19 16"/><path d="M19 16c-.536 4-3.358 6-7.5 6C7.358 22 4 20 4 16"/><path stroke-linecap="round" stroke-linejoin="round" d="M4 16v-4a2 2 0 1 1 4 0m4 1v-2a2 2 0 1 0-4 0v4"/><path stroke-linecap="round" stroke-linejoin="round" d="M13.692 17H11a2 2 0 1 1 0-4h4c2.21 0 4.5 2 3.5 5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Video(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="16" x="2" y="4" rx="4"/><path d="m15 12l-5-3v6z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VimeoFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.99 7.16c-.092 2.027-1.52 4.8-4.28 8.323C14.857 19.161 12.44 21 10.462 21c-1.225 0-2.26-1.122-3.106-3.359c-.564-2.055-1.127-4.11-1.697-6.16c-.627-2.237-1.3-3.359-2.025-3.359c-.155 0-.707.33-1.645.98L1 7.837c1.035-.906 2.06-1.805 3.066-2.71c1.38-1.185 2.422-1.805 3.112-1.868c1.633-.153 2.64.951 3.02 3.325c.408 2.556.69 4.15.851 4.77c.472 2.124.99 3.183 1.553 3.183c.437 0 1.099-.688 1.979-2.066c.88-1.378 1.351-2.425 1.415-3.143c.126-1.19-.345-1.782-1.415-1.782c-.5 0-1.018.114-1.553.342c1.03-3.353 3.002-4.982 5.913-4.885c2.157.057 3.175 1.446 3.049 4.156"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VkFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M23.45 5.948c.166-.546 0-.948-.795-.948H20.03c-.668 0-.976.347-1.143.73c0 0-1.335 3.196-3.226 5.272c-.612.602-.89.793-1.224.793c-.167 0-.418-.191-.418-.738V5.948c0-.656-.184-.948-.74-.948H9.151c-.417 0-.668.304-.668.593c0 .621.946.765 1.043 2.513v3.798c0 .833-.153.984-.487.984c-.89 0-3.055-3.211-4.34-6.885C4.45 5.288 4.198 5 3.527 5H.9c-.75 0-.9.347-.9.73c0 .682.89 4.07 4.145 8.551C6.315 17.341 9.37 19 12.153 19c1.669 0 1.875-.368 1.875-1.003v-2.313c0-.737.158-.884.687-.884c.39 0 1.057.192 2.615 1.667C19.11 18.216 19.403 19 20.405 19h2.625c.75 0 1.126-.368.91-1.096c-.238-.724-1.088-1.775-2.215-3.022c-.612-.71-1.53-1.475-1.809-1.858c-.389-.491-.278-.71 0-1.147c0 0 3.2-4.426 3.533-5.929" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VrAr(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 10a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v7a2 2 0 0 1-2 2h-3.868a2 2 0 0 1-1.715-.971l-1.56-2.6a1 1 0 0 0-1.714 0l-1.56 2.6A2 2 0 0 1 7.868 19H4a2 2 0 0 1-2-2zm1.813-3.219A4 4 0 0 1 7.14 5h9.718a4 4 0 0 1 3.328 1.781L21 8H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VscodeFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g fill="currentColor" clip-path="url(#akarIconsVscodeFill0)"><path d="M.228 8.37s-.584-.427.117-.995L1.98 5.897s.467-.497.962-.064l15.081 11.542v5.534s-.007.87-1.11.774z"/><path d="M4.116 11.937L.228 15.509s-.4.3 0 .837l1.805 1.66s.429.465 1.062-.065l4.121-3.158zm6.824.029l7.13-5.502l-.047-5.505s-.305-1.202-1.32-.576L7.216 9.11z"/><path d="M16.912 23.69c.414.428.916.288.916.288l5.556-2.767c.711-.49.611-1.098.611-1.098V3.588c0-.726-.735-.977-.735-.977L18.444.264c-1.052-.657-1.741.119-1.741.119s.886-.645 1.32.576v21.85c0 .15-.032.297-.095.43c-.127.259-.402.5-1.062.4z"/></g><defs><clipPath id="akarIconsVscodeFill0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VueFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.114 2H15l-3 4.9L9.429 2H0l12 21L24 2zM3 3.75h2.914L12 14.6l6.086-10.85H21L12 19.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wallet(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 6a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-6zm0 2h20"/><path d="M2 12h7c0 1 .6 3 3 3s3-2 3-3h7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WatchDevice(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 7a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2zm4 15h6M9 2h6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Water(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 22a8 8 0 0 1-8-8c0-3.502 2.71-6.303 5.093-8.87L12 2l2.907 3.13C17.29 7.698 20 10.499 20 14a8 8 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WhatsappFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#akarIconsWhatsappFill0)"><path fill="currentColor" fill-rule="evenodd" d="M17.415 14.382c-.298-.149-1.759-.867-2.031-.967c-.272-.099-.47-.148-.669.15c-.198.297-.767.966-.94 1.164c-.174.199-.347.223-.644.075c-.297-.15-1.255-.463-2.39-1.475c-.883-.788-1.48-1.761-1.653-2.059c-.173-.297-.019-.458.13-.606c.134-.133.297-.347.446-.52c.149-.174.198-.298.297-.497c.1-.198.05-.371-.025-.52c-.074-.149-.668-1.612-.916-2.207c-.241-.579-.486-.5-.668-.51c-.174-.008-.372-.01-.57-.01c-.198 0-.52.074-.792.372c-.273.297-1.04 1.016-1.04 2.479c0 1.462 1.064 2.875 1.213 3.074c.149.198 2.095 3.2 5.076 4.487c.71.306 1.263.489 1.694.625c.712.227 1.36.195 1.872.118c.57-.085 1.758-.719 2.006-1.413c.247-.694.247-1.289.173-1.413c-.074-.124-.272-.198-.57-.347m-5.422 7.403h-.004a9.87 9.87 0 0 1-5.032-1.378l-.36-.214l-3.742.982l.999-3.648l-.235-.374a9.861 9.861 0 0 1-1.511-5.26c.002-5.45 4.436-9.884 9.889-9.884a9.81 9.81 0 0 1 6.988 2.899a9.825 9.825 0 0 1 2.892 6.992c-.002 5.45-4.436 9.885-9.884 9.885m8.412-18.297A11.815 11.815 0 0 0 11.992 0C5.438 0 .102 5.335.1 11.892a11.864 11.864 0 0 0 1.587 5.945L0 24l6.304-1.654a11.881 11.881 0 0 0 5.684 1.448h.005c6.554 0 11.89-5.335 11.892-11.893a11.821 11.821 0 0 0-3.48-8.413" clip-rule="evenodd"/></g><defs><clipPath id="akarIconsWhatsappFill0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Width(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 12H2m20 0l-4 4m4-4l-4-4M2 12l4 4m-4-4l4-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wifi(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 10c6-6.667 14-6.667 20 0M6 14c3.6-4 8.4-4 12 0"/><circle cx="12" cy="18" r="1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WineGlass(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 13c6 0 8-4.477 8-10H4c0 5.523 2 10 8 10m0 0v7M5 8h14M8 22h8l-4-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Xfill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.808 10.469L20.88 2h-1.676l-6.142 7.353L8.158 2H2.5l7.418 11.12L2.5 22h1.676l6.486-7.765L15.842 22H21.5zm-2.296 2.748l-.752-1.107L4.78 3.3h2.575l4.826 7.11l.751 1.107l6.273 9.242h-2.574z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Xsmall(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M17 17L7 7m10 0L7 17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YarnFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g fill="currentColor" clip-path="url(#akarIconsYarnFill0)"><path fill-rule="evenodd" d="M23.994 11.675C23.825 5.23 18.561.013 12.004 0C5.21-.005-.3 5.668.013 12.556c.28 6.216 5.344 11.296 11.71 11.441c6.81.157 12.449-5.4 12.271-12.322M7.92 7.525c-.145.28-.258.575-.336.879c-.025.018-.02.05-.032.073c0 .05 0 .1-.014.147v.328c.017.05.012.1.017.15c.029.065-.019.146.054.198c.088.3.187.594.347.861c.05.084.042.129-.031.194a5.17 5.17 0 0 0-.987 1.183c-.472.802-.648 1.679-.648 2.6a.323.323 0 0 1-.016.104a.275.275 0 0 1-.062.105c-.609.667-.936 1.427-.786 2.348c.076.467.212.91.49 1.298a.395.395 0 0 1 .089.273c-.018.404.144.732.488.939c.58.351 1.207.477 1.875.328a.735.735 0 0 0 .133-.052c.068-.032.137-.063.198-.053c.068.01.128.063.189.115c.035.03.07.06.107.084c.381.228.809.286 1.238.267c1.044-.044 2.085-.107 3.121-.252c.342-.047.67-.133.95-.343a.64.64 0 0 1 .198-.1c1.049-.346 2.052-.798 2.967-1.406c.555-.364 1.137-.662 1.794-.808c.359-.078.62-.315.78-.648c.408-.844-.217-1.79-1.19-1.799c-.658-.005-1.263.2-1.843.49a8.696 8.696 0 0 0-.615.352c-.19.115-.38.231-.579.333l-.029.017c-.04.024-.083.049-.136.046v-.1c.052-1.395-.401-2.607-1.328-3.646c-.056-.062-.046-.102-.005-.162a6.753 6.753 0 0 0 .831-1.6c.344-.973.413-1.972.308-2.987c-.05-.489-.15-.966-.356-1.417c-.18-.385-.515-.567-.93-.474c-.111.023-.149-.01-.194-.102a4.747 4.747 0 0 0-.468-.817a.886.886 0 0 0-.685-.36c-.444-.025-.75.216-1.008.539a3.29 3.29 0 0 0-.458.805c-.036.087-.08.136-.15.162a.382.382 0 0 1-.094.02a3.28 3.28 0 0 0-2.014.97a1.34 1.34 0 0 1-.543.349c-.293.097-.496.299-.637.57" clip-rule="evenodd"/><path d="M14.947 15.813c0 .242-.056.478-.086.713c-.026.207-.005.231.205.195c.472-.082.9-.28 1.311-.515c.441-.251.861-.537 1.332-.726a3.013 3.013 0 0 1 1.134-.243c.343 0 .582.186.624.48c.04.272-.11.533-.384.59c-.781.166-1.456.551-2.113.98c-.882.569-1.842.98-2.841 1.295c-.037.01-.085.019-.108.041c-.236.25-.548.292-.863.333c-.884.116-1.773.168-2.665.22a3.049 3.049 0 0 1-.77-.026c-.41-.08-.583-.233-.645-.57c-.056-.301.086-.587.388-.784l.116-.073a1.306 1.306 0 0 1-.398-.364c-.05-.069-.076-.026-.092.031l-.15.57c-.034.126-.076.25-.129.37c-.152.348-.43.516-.797.55a1.976 1.976 0 0 1-1.065-.215c-.23-.113-.299-.283-.223-.533c.04-.134.116-.25.207-.391c-.317.063-.456-.121-.566-.344a2.374 2.374 0 0 1-.231-1.54c.081-.424.33-.77.63-1.069c.175-.175.228-.344.223-.59c-.04-1.425.498-2.595 1.605-3.499c.097-.081.195-.165.297-.239c.063-.044.07-.073.018-.138c-.286-.36-.509-.75-.61-1.204c-.117-.511.057-.955.314-1.38a.299.299 0 0 1 .158-.11c.323-.116.595-.291.842-.538c.555-.555 1.243-.81 2.03-.776c.1.005.144-.024.175-.118a4.41 4.41 0 0 1 .461-.98a1.26 1.26 0 0 1 .26-.31c.194-.16.36-.137.49.076c.242.386.428.8.616 1.215c.048.105.084.142.194.07a2.03 2.03 0 0 1 .417-.189c.116-.04.184-.013.234.108c.147.354.226.724.27 1.104c.013.105.037.207.024.312c-.008.228.003.457.005.688c0 .102-.036.205-.005.309c-.023.624-.189 1.214-.43 1.786c-.225.535-.553 1.01-.881 1.485c-.129.183-.126.186.042.336c.818.737 1.27 1.663 1.42 2.743c.006.042.004.087.004.131c-.034.1.003.197.005.294c.009.148-.025.296 0 .44"/></g><defs><clipPath id="akarIconsYarnFill0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YelpFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m4.98 11.264l4.701 2.278a1.075 1.075 0 0 1 .6 1.074a1.066 1.066 0 0 1-.81.927L4.393 16.8a1.088 1.088 0 0 1-1.187-.492a1.069 1.069 0 0 1-.146-.429a9.159 9.159 0 0 1 .424-3.996a1.07 1.07 0 0 1 .606-.645a1.085 1.085 0 0 1 .888.026m1.884 9.615l3.5-3.861a1.08 1.08 0 0 1 1.205-.277a1.076 1.076 0 0 1 .673 1.03l-.183 5.195a1.066 1.066 0 0 1-.396.793a1.083 1.083 0 0 1-.861.226a9.401 9.401 0 0 1-3.748-1.506a1.074 1.074 0 0 1-.46-.758a1.065 1.065 0 0 1 .27-.842m8.298-5.139l4.975 1.606a1.08 1.08 0 0 1 .657.596a1.064 1.064 0 0 1-.017.884a9.312 9.312 0 0 1-2.487 3.166a1.082 1.082 0 0 1-1.602-.258l-2.773-4.408a1.065 1.065 0 0 1 .065-1.226a1.078 1.078 0 0 1 1.182-.36m5.059-3.152l-5.029 1.433a1.085 1.085 0 0 1-1.169-.4A1.065 1.065 0 0 1 14 12.393l2.926-4.308a1.075 1.075 0 0 1 .755-.464a1.085 1.085 0 0 1 .85.257a9.222 9.222 0 0 1 2.379 3.25a1.069 1.069 0 0 1-.691 1.46M8.469.468a15.12 15.12 0 0 0-2.585.946a1.077 1.077 0 0 0-.564.65a1.063 1.063 0 0 0 .097.851l4.915 8.456a1.076 1.076 0 0 0 1.212.499a1.066 1.066 0 0 0 .799-1.034V1.072A1.065 1.065 0 0 0 11.622.06a1.084 1.084 0 0 0-.437-.057c-.918.072-1.826.228-2.715.465" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YoutubeFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#akarIconsYoutubeFill0)"><path fill="currentColor" d="M23.5 6.507a2.786 2.786 0 0 0-.766-1.27a3.05 3.05 0 0 0-1.338-.742C19.518 4 11.994 4 11.994 4a76.624 76.624 0 0 0-9.39.47a3.16 3.16 0 0 0-1.338.76c-.37.356-.638.795-.778 1.276A29.09 29.09 0 0 0 0 12c-.012 1.841.151 3.68.488 5.494c.137.479.404.916.775 1.269c.371.353.833.608 1.341.743c1.903.494 9.39.494 9.39.494a76.8 76.8 0 0 0 9.402-.47a3.05 3.05 0 0 0 1.338-.742a2.78 2.78 0 0 0 .765-1.27A28.38 28.38 0 0 0 24 12.023a26.579 26.579 0 0 0-.5-5.517M9.602 15.424V8.577l6.26 3.424z"/></g><defs><clipPath id="akarIconsYoutubeFill0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomFill(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#akarIconsZoomFill0)"><path fill="currentColor" fill-rule="evenodd" d="M24 12c0 6.627-5.373 12-12 12S0 18.627 0 12S5.373 0 12 0s12 5.373 12 12M6 16.2h9V9.6a1.8 1.8 0 0 0-1.8-1.8h-9v6.6A1.8 1.8 0 0 0 6 16.2m10.2-2.4l3.6 2.4V7.8l-3.6 2.4z" clip-rule="evenodd"/></g><defs><clipPath id="akarIconsZoomFill0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomIn(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="m21 21l-4.486-4.494M19 10.5a8.5 8.5 0 1 1-17 0a8.5 8.5 0 0 1 17 0ZM10.5 7v3.5m0 0V14m0-3.5H14m-3.5 0H7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomOut(children ...ElementRenderer) *AkarIconsIcon {
	return &AkarIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="m21 21l-4.486-4.494M19 10.5a8.5 8.5 0 1 1-17 0a8.5 8.5 0 0 1 17 0Zm-6 0H8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
