package vaadin

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = "4.3.2"
	hAttr          = "1em"
	viewbox        = "0 0 0 0"
	fill           = "currentColor"
)

type VaadinIcon struct {
	*SVGSVGElement
}

type VaadinIconFn func(children ...ElementRenderer) *VaadinIcon

var IconLookup = map[string]VaadinIconFn{
	"abacus":               Abacus,
	"absolutePosition":     AbsolutePosition,
	"academyCap":           AcademyCap,
	"accessibility":        Accessibility,
	"accordionMenu":        AccordionMenu,
	"addDock":              AddDock,
	"adjust":               Adjust,
	"adobeFlash":           AdobeFlash,
	"airplane":             Airplane,
	"alarm":                Alarm,
	"alignCenter":          AlignCenter,
	"alignJustify":         AlignJustify,
	"alignLeft":            AlignLeft,
	"alignRight":           AlignRight,
	"alt":                  Alt,
	"altA":                 AltA,
	"ambulance":            Ambulance,
	"anchor":               Anchor,
	"angleDoubleDown":      AngleDoubleDown,
	"angleDoubleLeft":      AngleDoubleLeft,
	"angleDoubleRight":     AngleDoubleRight,
	"angleDoubleUp":        AngleDoubleUp,
	"angleDown":            AngleDown,
	"angleLeft":            AngleLeft,
	"angleRight":           AngleRight,
	"angleUp":              AngleUp,
	"archive":              Archive,
	"archives":             Archives,
	"areaSelect":           AreaSelect,
	"arrowBackward":        ArrowBackward,
	"arrowCircleDown":      ArrowCircleDown,
	"arrowCircleDownO":     ArrowCircleDownO,
	"arrowCircleLeft":      ArrowCircleLeft,
	"arrowCircleLeftO":     ArrowCircleLeftO,
	"arrowCircleRight":     ArrowCircleRight,
	"arrowCircleRightO":    ArrowCircleRightO,
	"arrowCircleUp":        ArrowCircleUp,
	"arrowCircleUpO":       ArrowCircleUpO,
	"arrowDown":            ArrowDown,
	"arrowForward":         ArrowForward,
	"arrowLeft":            ArrowLeft,
	"arrowLongDown":        ArrowLongDown,
	"arrowLongLeft":        ArrowLongLeft,
	"arrowRight":           ArrowRight,
	"arrowUp":              ArrowUp,
	"arrows":               Arrows,
	"arrowsCross":          ArrowsCross,
	"arrowsLongH":          ArrowsLongH,
	"arrowsLongRight":      ArrowsLongRight,
	"arrowsLongUp":         ArrowsLongUp,
	"arrowsLongV":          ArrowsLongV,
	"asterisk":             Asterisk,
	"at":                   At,
	"automation":           Automation,
	"backspace":            Backspace,
	"backspaceA":           BackspaceA,
	"backwards":            Backwards,
	"ban":                  Ban,
	"barChart":             BarChart,
	"barChartH":            BarChartH,
	"barChartV":            BarChartV,
	"barcode":              Barcode,
	"bed":                  Bed,
	"bell":                 Bell,
	"bellO":                BellO,
	"bellSlash":            BellSlash,
	"bellSlashO":           BellSlashO,
	"boat":                 Boat,
	"bold":                 Bold,
	"bolt":                 Bolt,
	"bomb":                 Bomb,
	"book":                 Book,
	"bookDollar":           BookDollar,
	"bookPercent":          BookPercent,
	"bookmark":             Bookmark,
	"bookmarkO":            BookmarkO,
	"briefcase":            Briefcase,
	"browser":              Browser,
	"bug":                  Bug,
	"bugO":                 BugO,
	"building":             Building,
	"buildingO":            BuildingO,
	"bullets":              Bullets,
	"bullseye":             Bullseye,
	"buss":                 Buss,
	"button":               Button,
	"calc":                 Calc,
	"calcBook":             CalcBook,
	"calendar":             Calendar,
	"calendarBriefcase":    CalendarBriefcase,
	"calendarClock":        CalendarClock,
	"calendarEnvelope":     CalendarEnvelope,
	"calendarO":            CalendarO,
	"calendarUser":         CalendarUser,
	"camera":               Camera,
	"car":                  Car,
	"caretDown":            CaretDown,
	"caretLeft":            CaretLeft,
	"caretRight":           CaretRight,
	"caretSquareDownO":     CaretSquareDownO,
	"caretSquareLeftO":     CaretSquareLeftO,
	"caretSquareRightO":    CaretSquareRightO,
	"caretSquareUpO":       CaretSquareUpO,
	"caretUp":              CaretUp,
	"cart":                 Cart,
	"cartO":                CartO,
	"cash":                 Cash,
	"chart":                Chart,
	"chartGrid":            ChartGrid,
	"chartLine":            ChartLine,
	"chartThreeD":          ChartThreeD,
	"chartTimeline":        ChartTimeline,
	"chat":                 Chat,
	"check":                Check,
	"checkCircle":          CheckCircle,
	"checkCircleO":         CheckCircleO,
	"checkSquare":          CheckSquare,
	"checkSquareO":         CheckSquareO,
	"chevronCircleDown":    ChevronCircleDown,
	"chevronCircleDownO":   ChevronCircleDownO,
	"chevronCircleLeft":    ChevronCircleLeft,
	"chevronCircleLeftO":   ChevronCircleLeftO,
	"chevronCircleRight":   ChevronCircleRight,
	"chevronCircleRightO":  ChevronCircleRightO,
	"chevronCircleUp":      ChevronCircleUp,
	"chevronCircleUpO":     ChevronCircleUpO,
	"chevronDown":          ChevronDown,
	"chevronDownSmall":     ChevronDownSmall,
	"chevronLeft":          ChevronLeft,
	"chevronLeftSmall":     ChevronLeftSmall,
	"chevronRight":         ChevronRight,
	"chevronRightSmall":    ChevronRightSmall,
	"chevronUp":            ChevronUp,
	"chevronUpSmall":       ChevronUpSmall,
	"child":                Child,
	"circle":               Circle,
	"circleThin":           CircleThin,
	"clipboard":            Clipboard,
	"clipboardCheck":       ClipboardCheck,
	"clipboardCross":       ClipboardCross,
	"clipboardHeart":       ClipboardHeart,
	"clipboardPulse":       ClipboardPulse,
	"clipboardText":        ClipboardText,
	"clipboardUser":        ClipboardUser,
	"clock":                Clock,
	"close":                Close,
	"closeBig":             CloseBig,
	"closeCircle":          CloseCircle,
	"closeCircleO":         CloseCircleO,
	"closeSmall":           CloseSmall,
	"cloud":                Cloud,
	"cloudDownload":        CloudDownload,
	"cloudDownloadO":       CloudDownloadO,
	"cloudO":               CloudO,
	"cloudUpload":          CloudUpload,
	"cloudUploadO":         CloudUploadO,
	"cluster":              Cluster,
	"code":                 Code,
	"coffee":               Coffee,
	"cog":                  Cog,
	"cogO":                 CogO,
	"cogs":                 Cogs,
	"coinPiles":            CoinPiles,
	"coins":                Coins,
	"combobox":             Combobox,
	"comment":              Comment,
	"commentEllipsis":      CommentEllipsis,
	"commentEllipsisO":     CommentEllipsisO,
	"commentO":             CommentO,
	"comments":             Comments,
	"commentsO":            CommentsO,
	"compile":              Compile,
	"compress":             Compress,
	"compressSquare":       CompressSquare,
	"connect":              Connect,
	"connectO":             ConnectO,
	"controller":           Controller,
	"copy":                 Copy,
	"copyO":                CopyO,
	"copyright":            Copyright,
	"cornerLowerLeft":      CornerLowerLeft,
	"cornerLowerRight":     CornerLowerRight,
	"cornerUpperLeft":      CornerUpperLeft,
	"cornerUpperRight":     CornerUpperRight,
	"creditCard":           CreditCard,
	"crop":                 Crop,
	"crossCutlery":         CrossCutlery,
	"crosshairs":           Crosshairs,
	"css":                  Css,
	"ctrl":                 Ctrl,
	"ctrlA":                CtrlA,
	"cube":                 Cube,
	"cubes":                Cubes,
	"curlyBrackets":        CurlyBrackets,
	"cursor":               Cursor,
	"cursorO":              CursorO,
	"cutlery":              Cutlery,
	"dashboard":            Dashboard,
	"database":             Database,
	"dateInput":            DateInput,
	"deindent":             Deindent,
	"del":                  Del,
	"delA":                 DelA,
	"dentalChair":          DentalChair,
	"desktop":              Desktop,
	"diamond":              Diamond,
	"diamondO":             DiamondO,
	"diploma":              Diploma,
	"diplomaScroll":        DiplomaScroll,
	"disc":                 Disc,
	"doctor":               Doctor,
	"doctorBriefcase":      DoctorBriefcase,
	"dollar":               Dollar,
	"dotCircle":            DotCircle,
	"download":             Download,
	"downloadAlt":          DownloadAlt,
	"drop":                 Drop,
	"edit":                 Edit,
	"eject":                Eject,
	"elastic":              Elastic,
	"ellipsisCircle":       EllipsisCircle,
	"ellipsisCircleO":      EllipsisCircleO,
	"ellipsisDotsH":        EllipsisDotsH,
	"ellipsisDotsV":        EllipsisDotsV,
	"ellipsisH":            EllipsisH,
	"ellipsisV":            EllipsisV,
	"enter":                Enter,
	"enterArrow":           EnterArrow,
	"envelope":             Envelope,
	"envelopeO":            EnvelopeO,
	"envelopeOpen":         EnvelopeOpen,
	"envelopeOpenO":        EnvelopeOpenO,
	"envelopes":            Envelopes,
	"envelopesO":           EnvelopesO,
	"eraser":               Eraser,
	"esc":                  Esc,
	"escA":                 EscA,
	"euro":                 Euro,
	"exchange":             Exchange,
	"exclamation":          Exclamation,
	"exclamationCircle":    ExclamationCircle,
	"exclamationCircleO":   ExclamationCircleO,
	"exit":                 Exit,
	"exitO":                ExitO,
	"expand":               Expand,
	"expandFull":           ExpandFull,
	"expandSquare":         ExpandSquare,
	"externalBrowser":      ExternalBrowser,
	"externalLink":         ExternalLink,
	"eye":                  Eye,
	"eyeSlash":             EyeSlash,
	"eyedropper":           Eyedropper,
	"facebook":             Facebook,
	"facebookSquare":       FacebookSquare,
	"factory":              Factory,
	"family":               Family,
	"fastBackward":         FastBackward,
	"fastForward":          FastForward,
	"female":               Female,
	"file":                 File,
	"fileAdd":              FileAdd,
	"fileCode":             FileCode,
	"fileFont":             FileFont,
	"fileMovie":            FileMovie,
	"fileO":                FileO,
	"filePicture":          FilePicture,
	"filePresentation":     FilePresentation,
	"fileProcess":          FileProcess,
	"fileRefresh":          FileRefresh,
	"fileRemove":           FileRemove,
	"fileSearch":           FileSearch,
	"fileSound":            FileSound,
	"fileStart":            FileStart,
	"fileTable":            FileTable,
	"fileText":             FileText,
	"fileTextO":            FileTextO,
	"fileTree":             FileTree,
	"fileTreeSmall":        FileTreeSmall,
	"fileTreeSub":          FileTreeSub,
	"fileZip":              FileZip,
	"fill":                 Fill,
	"film":                 Film,
	"filter":               Filter,
	"fire":                 Fire,
	"flag":                 Flag,
	"flagCheckered":        FlagCheckered,
	"flagO":                FlagO,
	"flash":                Flash,
	"flask":                Flask,
	"flightLanding":        FlightLanding,
	"flightTakeoff":        FlightTakeoff,
	"flipH":                FlipH,
	"flipV":                FlipV,
	"folder":               Folder,
	"folderAdd":            FolderAdd,
	"folderO":              FolderO,
	"folderOpen":           FolderOpen,
	"folderOpenO":          FolderOpenO,
	"folderRemove":         FolderRemove,
	"folderSearch":         FolderSearch,
	"font":                 Font,
	"form":                 Form,
	"forward":              Forward,
	"frownO":               FrownO,
	"funcion":              Funcion,
	"funnel":               Funnel,
	"gamepad":              Gamepad,
	"gavel":                Gavel,
	"gift":                 Gift,
	"glass":                Glass,
	"glasses":              Glasses,
	"globe":                Globe,
	"globeWire":            GlobeWire,
	"golf":                 Golf,
	"googlePlus":           GooglePlus,
	"googlePlusSquare":     GooglePlusSquare,
	"grab":                 Grab,
	"grid":                 Grid,
	"gridBevel":            GridBevel,
	"gridBig":              GridBig,
	"gridBigO":             GridBigO,
	"gridH":                GridH,
	"gridSmall":            GridSmall,
	"gridSmallO":           GridSmallO,
	"gridV":                GridV,
	"groupIcon":            GroupIcon,
	"hammer":               Hammer,
	"hand":                 Hand,
	"handleCorner":         HandleCorner,
	"handsUp":              HandsUp,
	"handshake":            Handshake,
	"harddrive":            Harddrive,
	"harddriveO":           HarddriveO,
	"hash":                 Hash,
	"header":               Header,
	"headphones":           Headphones,
	"headset":              Headset,
	"healthCard":           HealthCard,
	"heart":                Heart,
	"heartO":               HeartO,
	"home":                 Home,
	"homeO":                HomeO,
	"hospital":             Hospital,
	"hourglass":            Hourglass,
	"hourglassEmpty":       HourglassEmpty,
	"hourglassEnd":         HourglassEnd,
	"hourglassStart":       HourglassStart,
	"inbox":                Inbox,
	"indent":               Indent,
	"info":                 Info,
	"infoCircle":           InfoCircle,
	"infoCircleO":          InfoCircleO,
	"input":                Input,
	"insert":               Insert,
	"institution":          Institution,
	"invoice":              Invoice,
	"italic":               Italic,
	"key":                  Key,
	"keyO":                 KeyO,
	"keyboard":             Keyboard,
	"keyboardO":            KeyboardO,
	"laptop":               Laptop,
	"layout":               Layout,
	"levelDown":            LevelDown,
	"levelDownBold":        LevelDownBold,
	"levelLeft":            LevelLeft,
	"levelLeftBold":        LevelLeftBold,
	"levelRight":           LevelRight,
	"levelRightBold":       LevelRightBold,
	"levelUp":              LevelUp,
	"levelUpBold":          LevelUpBold,
	"lifebuoy":             Lifebuoy,
	"lightbulb":            Lightbulb,
	"lineBarChart":         LineBarChart,
	"lineChart":            LineChart,
	"lineH":                LineH,
	"lineV":                LineV,
	"lines":                Lines,
	"linesList":            LinesList,
	"link":                 Link,
	"list":                 List,
	"listOl":               ListOl,
	"listSelect":           ListSelect,
	"listUl":               ListUl,
	"locationArrow":        LocationArrow,
	"locationArrowCircle":  LocationArrowCircle,
	"locationArrowCircleO": LocationArrowCircleO,
	"lock":                 Lock,
	"magic":                Magic,
	"magnet":               Magnet,
	"mailbox":              Mailbox,
	"male":                 Male,
	"mapMarker":            MapMarker,
	"margin":               Margin,
	"marginBottom":         MarginBottom,
	"marginLeft":           MarginLeft,
	"marginRight":          MarginRight,
	"marginTop":            MarginTop,
	"medal":                Medal,
	"megafone":             Megafone,
	"mehO":                 MehO,
	"menu":                 Menu,
	"microphone":           Microphone,
	"minus":                Minus,
	"minusCircle":          MinusCircle,
	"minusCircleO":         MinusCircleO,
	"minusSquareO":         MinusSquareO,
	"mobile":               Mobile,
	"mobileBrowser":        MobileBrowser,
	"mobileRetro":          MobileRetro,
	"modal":                Modal,
	"modalList":            ModalList,
	"money":                Money,
	"moneyDeposit":         MoneyDeposit,
	"moneyExchange":        MoneyExchange,
	"moneyWithdraw":        MoneyWithdraw,
	"moon":                 Moon,
	"moonO":                MoonO,
	"morning":              Morning,
	"movie":                Movie,
	"music":                Music,
	"mute":                 Mute,
	"nativeButton":         NativeButton,
	"newspaper":            Newspaper,
	"notebook":             Notebook,
	"nurse":                Nurse,
	"office":               Office,
	"openBook":             OpenBook,
	"option":               Option,
	"optionA":              OptionA,
	"options":              Options,
	"orientation":          Orientation,
	"out":                  Out,
	"outbox":               Outbox,
	"package":              Package,
	"padding":              Padding,
	"paddingBottom":        PaddingBottom,
	"paddingLeft":          PaddingLeft,
	"paddingRight":         PaddingRight,
	"paddingTop":           PaddingTop,
	"paintRoll":            PaintRoll,
	"paintbrush":           Paintbrush,
	"palete":               Palete,
	"panel":                Panel,
	"paperclip":            Paperclip,
	"paperplane":           Paperplane,
	"paperplaneO":          PaperplaneO,
	"paragraph":            Paragraph,
	"password":             Password,
	"paste":                Paste,
	"pause":                Pause,
	"pencil":               Pencil,
	"phone":                Phone,
	"phoneLandline":        PhoneLandline,
	"picture":              Picture,
	"pieBarChart":          PieBarChart,
	"pieChart":             PieChart,
	"piggyBank":            PiggyBank,
	"piggyBankCoin":        PiggyBankCoin,
	"pill":                 Pill,
	"pills":                Pills,
	"pin":                  Pin,
	"pinPost":              PinPost,
	"play":                 Play,
	"playCircle":           PlayCircle,
	"playCircleO":          PlayCircleO,
	"plug":                 Plug,
	"plus":                 Plus,
	"plusCircle":           PlusCircle,
	"plusCircleO":          PlusCircleO,
	"plusMinus":            PlusMinus,
	"plusSquareO":          PlusSquareO,
	"pointer":              Pointer,
	"powerOff":             PowerOff,
	"presentation":         Presentation,
	"print":                Print,
	"progressbar":          Progressbar,
	"puzzlePiece":          PuzzlePiece,
	"pyramidChart":         PyramidChart,
	"qrcode":               Qrcode,
	"question":             Question,
	"questionCircle":       QuestionCircle,
	"questionCircleO":      QuestionCircleO,
	"quoteLeft":            QuoteLeft,
	"quoteRight":           QuoteRight,
	"random":               Random,
	"raster":               Raster,
	"rasterLowerLeft":      RasterLowerLeft,
	"records":              Records,
	"recycle":              Recycle,
	"refresh":              Refresh,
	"reply":                Reply,
	"replyAll":             ReplyAll,
	"resizeH":              ResizeH,
	"resizeV":              ResizeV,
	"retweet":              Retweet,
	"rhombus":              Rhombus,
	"road":                 Road,
	"roadBranch":           RoadBranch,
	"roadBranches":         RoadBranches,
	"roadSplit":            RoadSplit,
	"rocket":               Rocket,
	"rotateLeft":           RotateLeft,
	"rotateRight":          RotateRight,
	"rss":                  Rss,
	"rssSquare":            RssSquare,
	"safe":                 Safe,
	"safeLock":             SafeLock,
	"scale":                Scale,
	"scaleUnbalance":       ScaleUnbalance,
	"scatterChart":         ScatterChart,
	"scissors":             Scissors,
	"screwdriver":          Screwdriver,
	"search":               Search,
	"searchMinus":          SearchMinus,
	"searchPlus":           SearchPlus,
	"select":               Select,
	"server":               Server,
	"share":                Share,
	"shareSquare":          ShareSquare,
	"shield":               Shield,
	"shift":                Shift,
	"shiftArrow":           ShiftArrow,
	"shop":                 Shop,
	"signIn":               SignIn,
	"signInAlt":            SignInAlt,
	"signOut":              SignOut,
	"signOutAlt":           SignOutAlt,
	"signal":               Signal,
	"sitemap":              Sitemap,
	"slider":               Slider,
	"sliders":              Sliders,
	"smileyO":              SmileyO,
	"sort":                 Sort,
	"soundDisable":         SoundDisable,
	"sparkLine":            SparkLine,
	"specialist":           Specialist,
	"spinner":              Spinner,
	"spinnerArc":           SpinnerArc,
	"spinnerThird":         SpinnerThird,
	"splineAreaChart":      SplineAreaChart,
	"splineChart":          SplineChart,
	"split":                Split,
	"splitH":               SplitH,
	"splitV":               SplitV,
	"spoon":                Spoon,
	"squareShadow":         SquareShadow,
	"star":                 Star,
	"starHalfLeft":         StarHalfLeft,
	"starHalfLeftO":        StarHalfLeftO,
	"starHalfRight":        StarHalfRight,
	"starHalfRightO":       StarHalfRightO,
	"starO":                StarO,
	"startCog":             StartCog,
	"stepBackward":         StepBackward,
	"stepForward":          StepForward,
	"stethoscope":          Stethoscope,
	"stock":                Stock,
	"stop":                 Stop,
	"stopCog":              StopCog,
	"stopwatch":            Stopwatch,
	"storage":              Storage,
	"strikethrough":        Strikethrough,
	"subscript":            Subscript,
	"suitcase":             Suitcase,
	"sunDown":              SunDown,
	"sunO":                 SunO,
	"sunRise":              SunRise,
	"superscript":          Superscript,
	"sword":                Sword,
	"tab":                  Tab,
	"tabA":                 TabA,
	"table":                Table,
	"tablet":               Tablet,
	"tabs":                 Tabs,
	"tag":                  Tag,
	"tags":                 Tags,
	"tasks":                Tasks,
	"taxi":                 Taxi,
	"teeth":                Teeth,
	"terminal":             Terminal,
	"textHeight":           TextHeight,
	"textInput":            TextInput,
	"textLabel":            TextLabel,
	"textWidth":            TextWidth,
	"thinSquare":           ThinSquare,
	"thumbsDown":           ThumbsDown,
	"thumbsDownO":          ThumbsDownO,
	"thumbsUp":             ThumbsUp,
	"thumbsUpO":            ThumbsUpO,
	"ticket":               Ticket,
	"timeBackward":         TimeBackward,
	"timeForward":          TimeForward,
	"timer":                Timer,
	"toolbox":              Toolbox,
	"tools":                Tools,
	"tooth":                Tooth,
	"touch":                Touch,
	"train":                Train,
	"trash":                Trash,
	"treeTable":            TreeTable,
	"trendindDown":         TrendindDown,
	"trendingUp":           TrendingUp,
	"trophy":               Trophy,
	"truck":                Truck,
	"twinColSelect":        TwinColSelect,
	"twitter":              Twitter,
	"twitterSquare":        TwitterSquare,
	"umbrella":             Umbrella,
	"underline":            Underline,
	"unlink":               Unlink,
	"unlock":               Unlock,
	"upload":               Upload,
	"uploadAlt":            UploadAlt,
	"user":                 User,
	"userCard":             UserCard,
	"userCheck":            UserCheck,
	"userClock":            UserClock,
	"userHeart":            UserHeart,
	"userStar":             UserStar,
	"users":                Users,
	"vaadinH":              VaadinH,
	"vaadinV":              VaadinV,
	"viewport":             Viewport,
	"vimeo":                Vimeo,
	"vimeoSquare":          VimeoSquare,
	"volume":               Volume,
	"volumeDown":           VolumeDown,
	"volumeOff":            VolumeOff,
	"volumeUp":             VolumeUp,
	"wallet":               Wallet,
	"warning":              Warning,
	"workplace":            Workplace,
	"wrench":               Wrench,
	"youtube":              Youtube,
	"youtubeSquare":        YoutubeSquare,
}

func Abacus(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v16h16V0zm14 2v3h-.1c-.2-.6-.8-1-1.4-1s-1.2.4-1.4 1H7.9c-.2-.6-.7-1-1.4-1s-1.2.4-1.4 1h-.2c-.2-.6-.7-1-1.4-1s-1.2.4-1.4 1H2V2zm-.1 8c-.2-.6-.8-1-1.4-1s-1.2.4-1.4 1h-.2c-.2-.6-.8-1-1.4-1s-1.2.4-1.4 1H4.9c-.2-.6-.7-1-1.4-1s-1.2.4-1.4 1H2V6h.1c.2.6.8 1 1.4 1s1.2-.4 1.4-1h.2c.2.6.8 1 1.4 1s1.2-.4 1.4-1h3.2c.2.6.8 1 1.4 1s1.2-.4 1.4-1h.1zM2 14v-3h.1c.2.6.8 1 1.4 1s1.2-.4 1.4-1h3.2c.2.6.8 1 1.4 1s1.2-.4 1.4-1h.2c.2.6.8 1 1.4 1s1.2-.4 1.4-1h.1v3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AbsolutePosition(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v16h16V0zm15 15H1V9h3v1l3-2l-3-2v1H1V1h6v3H6l2 3l2-3H9V1h6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AcademyCap(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.09 12.79a1 1 0 0 0-.086-1.638L15 5.33L14 6v5.15a1.001 1.001 0 0 0-.092 1.629l-.378.502a2.48 2.48 0 0 0-.53 1.498v1.222h.815a.88.88 0 0 0 .853-.664l.331-1.336v2h1v-1.21a2.486 2.486 0 0 0-.534-1.505zM8 0L0 4l8 5l8-5z"/><path fill="currentColor" d="M8 10L3 6.67v1.71C3 9.29 5.94 12 8 12s5-2.71 5-3.62V6.67z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Accessibility(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.4 10h-.5c.1.3.1.7.1 1c0 2.2-1.8 4-4 4s-4-1.8-4-4c0-2.1 1.6-3.8 3.7-4l-.2-1C2.9 6.4 1 8.4 1 11c0 2.8 2.2 5 5 5c2.4 0 4.4-1.7 4.9-3.9z"/><path fill="currentColor" d="M13.1 13L12 8H7.9l-.2-1H11V6H7.5l-.6-2.5c.9-.1 1.6-.8 1.6-1.7C8.5.8 7.7 0 6.7 0S5 .8 5 1.8c0 .6.3 1.2.8 1.5L7.1 9h4.1l1.2 5H15v-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AccordionMenu(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 4v8h16V4zm15 7H1V7h14zM0 0h16v3H0zm0 13h16v3H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AddDock(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 11v5h16v-5zm12 4H9v-3h3zm0-8V5c0-5-8-5-8-5s5 0 5 5v2H7l3.5 3L14 7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Adjust(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8M2 8c0-3.3 2.7-6 6-6v12c-3.3 0-6-2.7-6-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdobeFlash(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v16h16V0zm13 4.4C10 4.4 9.7 7 9.7 7H11v2H8.6C6.8 14.8 3 14 3 14v-2.5s2.5.6 3.9-4C8.7 1.4 13 2 13 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Airplane(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.3 6.5c.5-.5.9-.8 1.2-1.1c1.6-1.6 3.2-4.1 2.2-5.1s-3.4.6-5 2.2c-.3.3-.6.7-1.1 1.2L2.6.5C1.9.2 1.1.3.6.8l-.6.5L6.6 7c-1.3 1.6-2.7 3.1-3.4 4l-1.1-.6c-.5-.3-1.2-.3-1.6.2l-.3.3L3 13l2 2.8l.3-.3c.4-.4.5-1.1.2-1.6L5 12.8c.9-.7 2.4-2.1 4-3.4l5.7 6.6l.5-.5c.5-.5.6-1.3.3-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Alarm(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 5H7v5h4V9l-2.93.07zM5.46.87A2.099 2.099 0 0 0 2.651.335L1.66 1.1a2.095 2.095 0 0 0-.207 2.844zm8.88.23l-1-.77a2.097 2.097 0 0 0-2.796.534L14.54 3.94c.287-.356.46-.813.46-1.312c0-.602-.253-1.145-.659-1.528z"/><path fill="currentColor" d="M12.87 14A6.979 6.979 0 0 0 15 9.002A7.052 7.052 0 0 0 8.003 2A7.051 7.051 0 0 0 1 8.997a6.98 6.98 0 0 0 2.128 5.001l-.938.942a.63.63 0 0 0 .882.878l.998-.999c1.092.758 2.446 1.211 3.905 1.211s2.813-.453 3.928-1.226l.977 1.015a.63.63 0 0 0 .878-.882zm-10-5a5.18 5.18 0 0 1 5.127-5.13a5.181 5.181 0 0 1 5.133 5.127a5.181 5.181 0 0 1-5.127 5.133A5.181 5.181 0 0 1 2.87 9.003z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignCenter(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 0h6v3H5zM1 4h14v3H1zm2 4h10v3H3zm-3 4h16v3H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignJustify(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h16v3H0zm0 4h16v3H0zm0 8h16v3H0zm0-4h16v3H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignLeft(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h11v3H0zm0 4h15v3H0zm0 4h13v3H0zm0 4h16v3H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignRight(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 0h11v3H5zM1 4h15v3H1zm2 4h13v3H3zm-3 4h16v3H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Alt(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.89 9h2.22L5 5.1z"/><path fill="currentColor" d="M0 0v16h16V0zm7 12l-.61-2H3.61L3 12H2l2.27-8h1.46L8 12zm3 0H9V3h1zm4-5h-1v3.5s0 .5 1 .5v1c-1 0-2-.44-2-1.44V7h-.5V6h.5V5h1v1h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AltA(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 7V6h-1V5h-1v1h-.5v1h.5v3.56c0 1 .56 1.44 2 1.44v-1a.899.899 0 0 1-.998-.495L13 7zM9 3h1v9H9zm-6 9l.57-2h2.82L7 12h1L5.73 4H4.27L2 12zm2-6.9L6.11 9H3.89z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ambulance(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.18 14a2 2 0 1 1-3.999.001A2 2 0 0 1 6.18 14M14 14a2 2 0 1 1-3.999.001A2 2 0 0 1 14 14M5 6H4v1H3v1h1v1h1V8h1V7H5z"/><path fill="currentColor" d="m15.76 8.64l-3-4.53A2.501 2.501 0 0 0 10.682 3H8V2a1 1 0 0 0-2 0v1H1.5A1.5 1.5 0 0 0 0 4.5V13h1.37a3.238 3.238 0 0 1 2.812-2a3.236 3.236 0 0 1 2.81 1.978l2.188.021a3.238 3.238 0 0 1 2.812-2a3.003 3.003 0 0 1 2.822 1.979l1.187.021v-3.57a1.427 1.427 0 0 0-.243-.795zm-8.84-.52a2.5 2.5 0 1 1-3.017-2.997a2.48 2.48 0 0 1 3.014 3.014zM10 8V5h.859a2.25 2.25 0 0 1 1.866.992L14.05 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Anchor(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 9v2s-.8 1.7-4 1.9V6h2.2c.2.3.5.5.8.5c.6 0 1-.4 1-1s-.4-1-1-1c-.4 0-.7.2-.8.5H9V3.7c.6-.3 1-1 1-1.7c0-1.1-.9-2-2-2S6 .9 6 2c0 .7.4 1.4 1 1.7V5H4.8c-.1-.3-.4-.5-.8-.5c-.6 0-1 .4-1 1s.4 1 1 1c.4 0 .7-.2.8-.5H7v7c-3.3-.3-4-2-4-2V9H0s2.8 7 8 7c5 0 8-7 8-7zM8 1c.6 0 1 .4 1 1s-.4 1-1 1s-1-.4-1-1s.4-1 1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleDoubleDown(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 2v2l5 5l5-5V2L8 7z"/><path fill="currentColor" d="M3 7v2l5 5l5-5V7l-5 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleDoubleLeft(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 3h-2L7 8l5 5h2L9 8z"/><path fill="currentColor" d="M9 3H7L2 8l5 5h2L4 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleDoubleRight(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 13h2l5-5l-5-5H2l5 5z"/><path fill="currentColor" d="M7 13h2l5-5l-5-5H7l5 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleDoubleUp(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 14v-2L8 7l-5 5v2l5-5z"/><path fill="currentColor" d="M13 9V7L8 2L3 7v2l5-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleDown(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 4v2l-5 5l-5-5V4l5 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleLeft(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 13h-2L5 8l5-5h2L7 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleRight(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 13h2l5-5l-5-5H4l5 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleUp(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 12v-2l5-5l5 5v2L8 7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Archive(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 1h16v3H0zm1 4v11h14V5zm10 4H5V7h6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Archives(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 2H5v4h6zM9 4H7V3h2z"/><path fill="currentColor" d="M3 0v16h2v-1h6v1h2V0zm9 14H4V8h8zm0-7H4V1h8z"/><path fill="currentColor" d="M11 9H5v4h6zm-2 2H7v-1h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AreaSelect(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7.9 7.9l2.1 7.5l1.7-2.6l3.2 3.2l1.1-1.1l-3.3-3.2l2.7-1.6z"/><path fill="currentColor" d="M8 12H1V3h12v5.4l1 .2V2H0v11h8.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowBackward(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 7.9L6 3v3h2c8 0 8 8 8 8s-1-4-7.8-4H6v2.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleDown(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 8c0 4.4 3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8s-8 3.6-8 8m9 1.6l1.8-1.8l1.4 1.4L8 13.4L3.8 9.2l1.4-1.4L7 9.6V3h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleDownO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 8c0-3.9 3.1-7 7-7s7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7M0 8c0 4.4 3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8s-8 3.6-8 8"/><path fill="currentColor" d="m9 9.6l1.8-1.8l1.4 1.4L8 13.4L3.8 9.2l1.4-1.4L7 9.6V3h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleLeft(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8M6.4 9l1.8 1.8l-1.4 1.4L2.6 8l4.2-4.2l1.4 1.4L6.4 7H13v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleLeftO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1c3.9 0 7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7m0-1C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8"/><path fill="currentColor" d="m6.4 9l1.8 1.8l-1.4 1.4L2.6 8l4.2-4.2l1.4 1.4L6.4 7H13v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleRight(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 16c4.4 0 8-3.6 8-8s-3.6-8-8-8s-8 3.6-8 8s3.6 8 8 8m1.6-9L7.8 5.2l1.4-1.4L13.4 8l-4.2 4.2l-1.4-1.4L9.6 9H3V7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleRightO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 15c-3.9 0-7-3.1-7-7s3.1-7 7-7s7 3.1 7 7s-3.1 7-7 7m0 1c4.4 0 8-3.6 8-8s-3.6-8-8-8s-8 3.6-8 8s3.6 8 8 8"/><path fill="currentColor" d="M9.6 7L7.8 5.2l1.4-1.4L13.4 8l-4.2 4.2l-1.4-1.4L9.6 9H3V7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleUp(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 8c0-4.4-3.6-8-8-8S0 3.6 0 8s3.6 8 8 8s8-3.6 8-8M7 6.4L5.2 8.2L3.8 6.8L8 2.6l4.2 4.2l-1.4 1.4L9 6.4V13H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleUpO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 8c0 3.9-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7s7 3.1 7 7m1 0c0-4.4-3.6-8-8-8S0 3.6 0 8s3.6 8 8 8s8-3.6 8-8"/><path fill="currentColor" d="M7 6.4L5.2 8.2L3.8 6.8L8 2.6l4.2 4.2l-1.4 1.4L9 6.4V13H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDown(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 8.6L9 12.2V0H7v12.2L3.5 8.6l-1.4 1.5L8 16l5.9-5.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowForward(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 7.9L10 3v3H8c-8 0-8 8-8 8s1-4 7.8-4H10v2.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeft(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.4 12.5L3.8 9H16V7H3.8l3.6-3.5l-1.5-1.4L0 8l5.9 5.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongDown(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 1h2v11h2l-3 3l-3-3h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongLeft(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 7v2H4v2L1 8l3-3v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRight(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.6 3.5L12.1 7H0v2h12.1l-3.5 3.5l1.4 1.4L16 8l-6-5.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUp(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.4 7.4L7 3.8V16h2V3.8l3.5 3.6l1.4-1.5L8 0L2 5.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Arrows(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16 8l-3-3v2H9V3h2L8 0L5 3h2v4H3V5L0 8l3 3V9h4v4H5l3 3l3-3H9V9h4v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsCross(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 5V1h-4l1.3 1.3L8 6.6L3.7 2.3L5 1H1v4l1.3-1.3L6.6 8l-4.3 4.3L1 11v4h4l-1.3-1.3L8 9.4l4.3 4.3L11 15h4v-4l-1.3 1.3L9.4 8l4.3-4.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsLongH(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16 8l-3-3v2H3V5L0 8l3 3V9h10v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsLongRight(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 9V7h11V5l3 3l-3 3V9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsLongUp(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 15H7V4H5l3-3l3 3H9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsLongV(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 3h2L8 0L5 3h2v10H5l3 3l3-3H9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Asterisk(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.9 5.7l-2-3.4L10 4.5V0H6v4.5L2 2.3L0 5.7L3.9 8L0 10.3l2 3.4l4-2.2V16h4v-4.5l3.9 2.2l2-3.4l-4-2.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func At(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 12.2c-2.3 0-4.2-1.9-4.2-4.2s1.9-4.2 4.2-4.2s4.2 1.9 4.2 4.2c.1 2.3-1.9 4.2-4.2 4.2m0-7C6 5.2 4.8 6.5 4.8 8s1.2 2.8 2.8 2.8s2.8-1.2 2.8-2.8S9 5.2 7.5 5.2"/><path fill="currentColor" d="M8 16c-4.4 0-8-3.6-8-8s3.6-8 8-8s8 3.6 8 8c0 1.5-.4 3-1.2 4.2c-.3.5-1.1 1.2-2.3 1.2c-.8 0-1.3-.3-1.6-.6c-.7-.7-.6-1.8-.6-1.9V4h1.5v7c0 .2 0 .6.2.8c0 0 .2.2.5.2c.7 0 1.1-.5 1.1-.5c.6-1 1-2.2 1-3.4c0-3.6-2.9-6.5-6.5-6.5S1.5 4.4 1.5 8s2.9 6.5 6.5 6.5c.7 0 1.3-.1 1.9-.3l.4 1.4c-.7.3-1.5.4-2.3.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Automation(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 12a2 2 0 1 1-3.999.001A2 2 0 0 1 14 12"/><path fill="currentColor" d="M11.7 16c-.8 0-1.6-.2-2.3-.7L3.2 12c-.5-.4-.9-.6-1.3-1C.7 9.8 0 8.1 0 6.4s.7-3.3 1.9-4.5C3.1.7 4.7 0 6.4 0S9.7.7 11 1.9c.4.4.6.7 1 1.2l3.5 6.4c1 1.7.7 3.8-.7 5.2c-.9.9-1.9 1.3-3.1 1.3M6.4 1C5 1 3.6 1.6 2.6 2.6S1 5 1 6.4c0 1.5.6 2.8 1.6 3.8c.3.3.6.5 1.1.8l6.3 3.4c.6.4 1.2.5 1.8.5c.9 0 1.7-.3 2.3-1c1.1-1.1 1.3-2.7.5-4l-3.5-6.4c-.3-.4-.5-.7-.8-1C9.2 1.6 7.9 1 6.4 1"/><path fill="currentColor" d="M11 7V6l-1.4-.5c-.1-.2-.1-.3-.2-.5l.6-1.3l-.7-.7l-1.3.6c-.2-.1-.3-.1-.5-.2L7 2H6l-.5 1.4c-.2.1-.3.1-.5.2L3.7 3l-.7.7l.6 1.3c-.1.2-.1.3-.2.5L2 6v1l1.4.5c.1.2.1.3.2.5L3 9.3l.7.7L5 9.4c.2.1.3.2.5.2L6 11h1l.5-1.4c.2-.1.3-.1.5-.2l1.3.6l.7-.7L9.4 8c.1-.2.2-.3.2-.5zM6.5 8C5.7 8 5 7.3 5 6.5S5.7 5 6.5 5S8 5.7 8 6.5S7.3 8 6.5 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Backspace(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 2v12h16V2zm13 7H6v2L3 8l3-3v2h7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BackspaceA(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 12L0 8l5-4v2h11v4H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Backwards(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 15V1L8 8zm-8 0V1L0 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ban(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8m0 2c1.3 0 2.5.4 3.5 1.1l-8.4 8.4C2.4 10.5 2 9.3 2 8c0-3.3 2.7-6 6-6m0 12c-1.3 0-2.5-.4-3.5-1.1l8.4-8.4c.7 1 1.1 2.2 1.1 3.5c0 3.3-2.7 6-6 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarChart(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 15h15v1H0zm0-4h3v3H0zm4-2h3v5H4zm4-4h3v9H8zm4-5h3v14h-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarChartH(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 15V0H0v16h16v-1z"/><path fill="currentColor" d="M2 8h4v6H2zm5-6h4v12H7zm5 4h4v8h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarChartV(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 15V0H0v16h16v-1z"/><path fill="currentColor" d="M8 0v4H2V0zm6 5v4H2V5zm-4 5v4H2v-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Barcode(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 3h1v10H0zm8 0h2v10H8zm3 0h1v10h-1zm2 0h1v10h-1zm2 0h1v10h-1zM2 3h3v10H2zm4 0h1v10H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bed(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.28 7H7L5.85 5.32a3.315 3.315 0 0 0-2.295-1.319L3 4v1.54a1.248 1.248 0 0 0 1.232 1.461L4.282 7zM13 7v-.29A1.71 1.71 0 0 0 11.322 5H6.63C7.13 5.62 8 7 8 7z"/><path fill="currentColor" d="M15 5.1a1 1 0 0 0-1 1V8H2V4a1 1 0 0 0-2 0v9h2v-2h12v2h2V6.1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bell(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 14h4s.1 2-2 2s-2-2-2-2m6.7-2.6c-.5-.2-.7-.7-.7-1.2V5s-.2-2.4-3-2.9V1s.1-1-1-1s-1 1-1 1v1.1C4.2 2.6 4 5 4 5v5.2c0 .5-.3 1-.7 1.2L2 12v1h12v-1zM6 4.8V12H4c.8 0 1-1 1-1V5s0-.8.7-1.4C6.4 2.9 7 3 7 3s-1 .7-1 1.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.7 11.4c-.5-.2-.7-.7-.7-1.2V5s0-2.4-3-2.9V1s.1-1-1-1s-1 1-1 1v1.1C4 2.6 4 5 4 5v5.2c0 .5-.3 1-.7 1.2L2 12v2h4s-.1 2 2 2s2-2 2-2h4v-2zM13 13H3v-.4l.7-.4c.8-.3 1.3-1.1 1.3-2V5c0-.1 0-1.6 2.2-1.9l.8-.2l.8.1c2 .4 2.2 1.7 2.2 2v5.2c0 .9.5 1.7 1.3 2.1l.7.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellSlash(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.2 0l-3.6 3.6C11.1 3 10.4 2.3 9 2.1V1s.1-1-1-1s-1 1-1 1v1.1C4.2 2.6 4 5 4 5v5.2c0 .5-.3 1-.7 1.2L2 12v1h.3L0 15.3v.7h.7L16 .6V0zM6 4.8v4.5l-1 1V5s0-.8.7-1.4C6.4 2.9 7 3 7 3s-1 .7-1 1.8M8 16c2.1 0 2-2 2-2H6s-.1 2 2 2m4-5.8V5.6l-6 6l-.3.4l-1 1H14v-1l-1.3-.6c-.4-.3-.7-.7-.7-1.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellSlashO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.2 0l-3.6 3.6C11.2 3 10.4 2.3 9 2.1V1s.1-1-1-1s-1 1-1 1v1.1C4 2.6 4 5 4 5v5.2c0 .5-.3 1-.7 1.2L2 12v1.3l-2 2v.7h.7L16 .6V0zM5 10.3c0-.1 0-.1 0 0V5c0-.1.1-1.6 2.2-1.9l.8-.2l.8.1c1.2.2 1.8.8 2 1.3zm7-.1V5.6l-1 1v3.5c0 .9.5 1.7 1.3 2.1l.7.4v.4H4.7l-1 1h2.4s-.1 2 2 2s2-2 2-2H14v-2l-1.3-.6c-.4-.3-.7-.7-.7-1.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Boat(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.5 9.6c1.1.7 2.5 1.9 2.5 3.3V14h.1s.9 0 2-1c1 1 2 1 2 1s1 0 2-1c1 1 1.9 1 1.9 1h.1v-1.1c0-1.4 1.4-2.6 2.5-3.3c.6-.4.5-1.2-.2-1.4L13 7.8V4h-1V3H9V1H7v2H4v1H3v3.8l-1.3.4c-.8.2-.8 1-.2 1.4M4 5h1V4h6v1h1v2.5l-3.3-1c-.5-.1-1-.1-1.5 0L4 7.5zm10 9c-1 1-2 1-2 1s-1 0-2-1c-1 1-2 1-2 1s-1 0-2-1c-1 1-2 1-2 1s-1 0-2-1c-1 1-2 1-2 1v1h16v-1s-1 0-2-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bold(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 7.5s2-.8 2-3.6C13-.2 7.9 0 6 0H2v16h4c3.7 0 8 0 8-4.4c0-3.8-3-4.1-3-4.1M9 4.4C9 6.2 7.5 6 6 6V3c1.8 0 3 .1 3 1.4M6 13V9c1.8 0 4-.3 4 2.2c0 1.9-2.5 1.8-4 1.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bolt(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.99 0L.98 9.38L7 8.96L2.04 16L15 6l-7.01.47L15 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bomb(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1h1v1h-1zm0 4h1v1h-1zm2-2h1v1h-1zm-4 0h1v1h-1zm4.6-.9l.7-.7l-.7-.7l-1.4 1.4l.7.7zm-.7 2.1l-.7.7l1.4 1.4l.7-.7l-.7-.7zm-2.8-1.4l.7-.7L10.4.7l-.7.7l.7.7z"/><path fill="currentColor" d="m10.4 6.4l2-2l-.7-.7l-2 2L9 5l-.7.8C7.5 5.3 6.5 5 5.5 5C2.5 5 0 7.5 0 10.5S2.5 16 5.5 16s5.5-2.5 5.5-5.5c0-1-.3-1.9-.7-2.8L11 7zM6 7.2C4 7.2 2.6 9 2.6 10h-1C1.6 8 4 6.2 6 6.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Book(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.6 2.5C11 1.3 11 0 11 0H2v12.5C2 14.4 3.6 16 5.5 16H14V3s-1-.2-1.4-.5M4 2h5v2H4zm9 13H5.5c-1 0-1.8-.6-2-1.3c-.1-.4 0-.7.4-.7H11V2.7c.4.6 1.2 1.1 2 1.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookDollar(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.9 2.5C11.3 1.3 11.5 0 11.5 0H2v12.5C2 14.4 4.1 16 6 16h8V3s-.8-.2-1.1-.5M7 6.3c-.9-.3-2.3-.8-2.3-1.9C4.8 3.6 6 3 6 2.8V2h1v.7c1 .1 1.8.4 1.9.4l-.3.9s-.7-.3-1.5-.3c-.7 0-1.1.3-1.2.8c0 .3.5.6 1.3.9c1.5.5 1.9 1.1 1.9 1.9C9.1 8 9 8.9 7 9.1v.9H6v-.8c0-.1-1.4-.5-1.5-.5l.5-.9s1.1.5 2 .4s1.3-.6 1.3-1c.1-.3-.4-.6-1.3-.9m6 8.7H6c-1 0-1.8-.6-2-1.3c-.1-.3 0-.7.4-.7H11V2.7c1 .6 2 1.1 2 1.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookPercent(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.6 2.5C11 1.3 11 0 11 0H2v12.5C2 14.4 3.6 16 5.5 16H14V3s-1-.2-1.4-.5m-7.1.7c.8 0 1.5.7 1.5 1.6s-.7 1.4-1.5 1.4S4 5.6 4 4.8s.7-1.6 1.5-1.6M9 3h1l-5 7H4zm1 5.5c0 .8-.7 1.5-1.5 1.5S7 9.3 7 8.5S7.7 7 8.5 7s1.5.7 1.5 1.5m3 6.5H5.5c-1 0-1.8-.6-2-1.3c-.1-.4 0-.7.4-.7H11V2.7c0 .6 1 1.1 2 1.3z"/><path fill="currentColor" d="M9 8.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0M6 4.8a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookmark(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 0v1h10l.1-1zm0 2h10v14l-5-5l-5 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 0v16l5-5l5 5V0zm9 13.7L8 9.8l-4 3.9V3h8zM12 2H4V1h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Briefcase(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 3V1H5v2H0v12h16V3zm-1 0H6V2h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Browser(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1V0H0v15h1v1h15V1zM3 1h9v1H3zM1 1h1v1H1zm0 2h13v11H1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bug(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 6.2a7.596 7.596 0 0 0 3.903-1.129a.31.31 0 0 0 .098-.229L12 4.819a2.914 2.914 0 0 0-1.781-2.522c-.137-.05-.219-.16-.219-.29V.499a.5.5 0 0 0-1 0v1.2a.3.3 0 0 1-.3.3H7.3a.3.3 0 0 1-.3-.3v-1.2a.5.5 0 0 0-1 0v1.5a.3.3 0 0 1-.198.269A2.92 2.92 0 0 0 4 4.812l-.001.029c0 .102.051.193.13.247a5.847 5.847 0 0 0 3.89 1.11zM10 3a1 1 0 1 1 0 2a1 1 0 0 1 0-2M6 3a1 1 0 1 1 0 2a1 1 0 0 1 0-2"/><path fill="currentColor" d="M13 8V7a6.196 6.196 0 0 0 2-4.497a.5.5 0 1 0-1-.003a4.544 4.544 0 0 1-1.105 2.906A4.777 4.777 0 0 1 9.326 7H6.73l-.075.001a4.777 4.777 0 0 1-3.561-1.586A4.54 4.54 0 0 1 2 2.514a.5.5 0 1 0-1-.003a6.192 6.192 0 0 0 1.996 4.486L3 8.001c-3 .06-3 1.42-3 3.47a.5.5 0 0 0 1 0c0-1.72 0-2.4 2-2.47a9.633 9.633 0 0 0 .612 3.136c-.383.006-.696.176-.942.414a3.857 3.857 0 0 0-.711 2.242c0 .2.015.397.044.589L2 15.5a.5.5 0 0 0 1 0v-.14a3.272 3.272 0 0 1 .389-2.096c.165-.152.401-.257.66-.264a3.802 3.802 0 0 0 2.934 1.998L7 14h2v1a3.798 3.798 0 0 0 2.94-1.98c.012-.02.015-.02.018-.02a1 1 0 0 1 .663.251a3.26 3.26 0 0 1 .377 2.127l.002.121a.5.5 0 0 0 1 0v-.14a3.875 3.875 0 0 0-.678-2.802a1.841 1.841 0 0 0-.9-.466c.336-.917.55-1.975.578-3.08c2-.012 2 .708 2 2.458a.5.5 0 0 0 1 0c0-2.03 0-3.39-3-3.47zm-4 5H7v-1h2zm0-2H7v-1h2zm0-2H7V8h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BugO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 8V7a6.196 6.196 0 0 0 2-4.497a.5.5 0 1 0-1-.003a4.544 4.544 0 0 1-1.105 2.906A4.777 4.777 0 0 1 9.326 7H6.73l-.075.001a4.777 4.777 0 0 1-3.561-1.586A4.54 4.54 0 0 1 2 2.514a.5.5 0 1 0-1-.003a6.192 6.192 0 0 0 1.996 4.486L3 8.001c-3 .06-3 1.42-3 3.47a.5.5 0 0 0 1 0c0-1.72 0-2.4 2-2.47a9.633 9.633 0 0 0 .612 3.136c-.383.006-.696.176-.942.414a3.857 3.857 0 0 0-.711 2.242c0 .2.015.397.044.589L2 15.5a.5.5 0 0 0 1 0v-.14a3.272 3.272 0 0 1 .389-2.096c.165-.152.401-.257.66-.264a4.748 4.748 0 0 0 2.92 1.994L7 14h2v1a4.745 4.745 0 0 0 2.939-1.983c.013-.017.016-.017.019-.017a1 1 0 0 1 .663.251a3.26 3.26 0 0 1 .377 2.127l.002.121a.5.5 0 0 0 1 0v-.14a3.875 3.875 0 0 0-.678-2.802a1.841 1.841 0 0 0-.9-.466c.336-.917.55-1.975.578-3.08c2-.012 2 .708 2 2.458a.5.5 0 0 0 1 0c0-2.03 0-3.39-3-3.47zm-7 5.5a3.333 3.333 0 0 1-1.083-.989L4.67 12.1l-.15-.39A8.478 8.478 0 0 1 4 9.013V7.35a5.425 5.425 0 0 0 1.973.647L6 13.57zm3-.5H7v-1h2zm0-2H7v-1h2zm0-2H7V8h2zm3 0a8.642 8.642 0 0 1-.54 2.77l-.13.33l-.24.4c-.285.411-.65.747-1.074.992L10 8a5.515 5.515 0 0 0 2.029-.624z"/><path fill="currentColor" d="M8 6.2a7.596 7.596 0 0 0 3.903-1.129a.31.31 0 0 0 .098-.229L12 4.819a2.914 2.914 0 0 0-1.781-2.522c-.137-.05-.219-.16-.219-.29V.499a.5.5 0 0 0-1 0v1.2a.3.3 0 0 1-.3.3H7.3a.3.3 0 0 1-.3-.3v-1.2a.5.5 0 0 0-1 0v1.5a.3.3 0 0 1-.198.269A2.92 2.92 0 0 0 4 4.812l-.001.029c0 .102.051.193.13.247a5.847 5.847 0 0 0 3.89 1.11zM10 3a1 1 0 1 1 0 2a1 1 0 0 1 0-2M6 3a1 1 0 1 1 0 2a1 1 0 0 1 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Building(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 0v16h4v-3h2v3h4V0zm3 12H4v-2h2zm0-3H4V7h2zm0-3H4V4h2zm0-3H4V1h2zm3 9H7v-2h2zm0-3H7V7h2zm0-3H7V4h2zm0-3H7V1h2zm3 9h-2v-2h2zm0-3h-2V7h2zm0-3h-2V4h2zm0-3h-2V1h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuildingO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 0v16h12V0zm11 15H9v-3H7v3H3V1h10z"/><path fill="currentColor" d="M4 9h2v2H4zm3 0h2v2H7zm3 0h2v2h-2zM4 6h2v2H4zm3 0h2v2H7zm3 0h2v2h-2zM4 3h2v2H4zm3 0h2v2H7zm3 0h2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bullets(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 2.5C0 3.3.7 4 1.5 4S3 3.3 3 2.5S2.3 1 1.5 1S0 1.7 0 2.5m0 5C0 8.3.7 9 1.5 9S3 8.3 3 7.5S2.3 6 1.5 6S0 6.7 0 7.5m0 5c0 .8.7 1.5 1.5 1.5S3 13.3 3 12.5S2.3 11 1.5 11S0 11.7 0 12.5M5 1h11v3H5zm0 5h11v3H5zm0 5h11v3H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bullseye(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8m0 14.9c-3.8 0-6.9-3.1-6.9-6.9S4.2 1.1 8 1.1s6.9 3.1 6.9 6.9s-3.1 6.9-6.9 6.9"/><path fill="currentColor" d="M8 2.3C4.8 2.3 2.3 4.8 2.3 8s2.6 5.7 5.7 5.7s5.7-2.6 5.7-5.7S11.2 2.3 8 2.3m0 10.3c-2.5 0-4.6-2.1-4.6-4.6S5.5 3.4 8 3.4s4.6 2.1 4.6 4.6c0 2.5-2.1 4.6-4.6 4.6"/><path fill="currentColor" d="M8 4.6C6.1 4.6 4.6 6.1 4.6 8s1.5 3.4 3.4 3.4s3.4-1.5 3.4-3.4S9.9 4.6 8 4.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Buss(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.67 4H14V2a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v2h-.68a.32.32 0 0 0-.32.32v2.36c0 .177.143.32.32.32H2v6c0 .55 0 1 1 1v1.5a.5.5 0 0 0 .5.5h2a.5.5 0 0 0 .5-.5V14h4v1.5a.5.5 0 0 0 .5.5h2a.5.5 0 0 0 .5-.5V14c1 0 1-.45 1-1V7h.67a.33.33 0 1 0 0-.66a.33.33 0 0 0 0 .66a.33.33 0 0 0 .33-.33V4.33a.33.33 0 0 0-.33-.33M6 1h4v1H6zM4 12a1 1 0 1 1 0-2a1 1 0 0 1 0 2M3 8V3h10v5zm9 4a1 1 0 1 1 0-2a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Button(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.7 5.3l-1-1c-.2-.2-.4-.3-.7-.3H1c-.6 0-1 .4-1 1v5c0 .3.1.6.3.7l1 1c.2.2.4.3.7.3h13c.6 0 1-.4 1-1V6c0-.3-.1-.5-.3-.7M14 10H1V5h13z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calc(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 3h6v2H9zm0 8h6v2H9zM5 1H3v2H1v2h2v2h2V5h2V3H5zm2 9.4L5.6 9L4 10.6L2.4 9L1 10.4L2.6 12L1 13.6L2.4 15L4 13.4L5.6 15L7 13.6L5.4 12zm6 4.1a1 1 0 1 1-2 0a1 1 0 0 1 2 0m0-5a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalcBook(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.9 0c-1.3 0-2 .4-2.4.8C9.1.4 8.4 0 7 0C3.6 0 3 2 3 2v4H0v10h7v-4.6l1.5-.2s.2-.3.3.7h1.3c.1-1 .4-.7.4-.7l5.5.7V2.1S15.4 0 11.9 0M1 7h5v2H1zm5 3v1H5v-1zm-2 0v1H3v-1zm-2 5H1v-1h1zm0-2H1v-1h1zm0-2H1v-1h1zm2 4H3v-1h1zm0-2H3v-1h1zm2 2H5v-1h1zm0-2H5v-1h1zm3-3.5c-.9-.1-1.3-.3-2-.3V6H4V2.1c0-.4.8-1.5 3-1.5c1.8 0 1.9.8 1.9 1v7.9zm6 .4c-1-.4-1.1-.7-2.5-.7h-.2c-1 0-1.3.2-2.3.4V1.8c0-.2.2-1.1 1.9-1.1c2.3 0 3.1.9 3.1 1.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 1v3h-3V1H5v3H2V1H0v15h16V1zM3 15H1v-2h2zm0-3H1v-2h2zm0-3H1V7h2zm3 6H4v-2h2zm0-3H4v-2h2zm0-3H4V7h2zm3 6H7v-2h2zm0-3H7v-2h2zm0-3H7V7h2zm3 6h-2v-2h2zm0-3h-2v-2h2zm0-3h-2V7h2zm3 6h-2v-2h2zm0-3h-2v-2h2zm0-3h-2V7h2z"/><path fill="currentColor" d="M3 0h1v3H3zm9 0h1v3h-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarBriefcase(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 0h1v3H3zm8 0h1v3h-1z"/><path fill="currentColor" d="M13 1v3h-3V1H5v3H2V1H0v14h5v-1H1V6h13v3h1V1z"/><path fill="currentColor" d="M13 10V8H9v2H6v6h10v-6zm-3-1h2v1h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarClock(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 0h1v3H3zm8 0h1v3h-1z"/><path fill="currentColor" d="M6.6 14H1V6h13v.6c.4.2.7.4 1 .7V1h-2v3h-3V1H5v3H2V1H0v14h7.3c-.3-.3-.5-.6-.7-1"/><path fill="currentColor" d="M14 12h-3V9h1v2h2z"/><path fill="currentColor" d="M11.5 8c1.9 0 3.5 1.6 3.5 3.5S13.4 15 11.5 15S8 13.4 8 11.5S9.6 8 11.5 8m0-1C9 7 7 9 7 11.5S9 16 11.5 16s4.5-2 4.5-4.5S14 7 11.5 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarEnvelope(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 0h1v2H3zm6 0h1v2H9z"/><path fill="currentColor" d="M13 7V1h-2v2H8V1H5v2H2V1H0v12h4v3h12V7zm-9 5H1V5h11v2H4zm1-1.8l2.6 1.5L5 14.3zm.7 4.8l2.8-2.8l1.5.9l1.5-.8l2.8 2.8H5.7zm9.3-.7l-2.6-2.6l2.6-1.4zm0-5.1l-5 2.7L5 9V8h10zm.4.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 1v3h-3V1H5v3H2V1H0v15h16V1zm1 14H1V6h14z"/><path fill="currentColor" d="M3 0h1v3H3zm9 0h1v3h-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarUser(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 0h1v3H3zm8 0h1v3h-1z"/><path fill="currentColor" d="M9 14.1c0-.1 0-.1 0 0L1 14V6h13v1.2c.4.1.7.3 1 .6V1h-2v3h-3V1H5v3H2V1H0v14h9z"/><path fill="currentColor" d="M15 10a2 2 0 1 1-3.999.001A2 2 0 0 1 15 10"/><path fill="currentColor" d="M13.9 12h-1.8c-1.1 0-2.1.9-2.1 2.1V16h6v-1.9c0-1.2-.9-2.1-2.1-2.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Camera(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 9a3 3 0 1 1-6 0a3 3 0 0 1 6 0"/><path fill="currentColor" d="M11 4V1H5v3H0v9h5c.8.6 1.9 1 3 1s2.2-.4 3-1h5V4zM6 2h4v2H6zm2 11c-2.2 0-4-1.8-4-4s1.8-4 4-4s4 1.8 4 4s-1.8 4-4 4m7-7h-2V5h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Car(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15 6.1l-1.4-2.9c-.4-.7-1.1-1.2-1.9-1.2H4.3c-.8 0-1.5.5-1.9 1.2L1 6.1c-.6.1-1 .6-1 1.1v3.5c0 .6.4 1.1 1 1.2v2c0 .6.5 1.1 1.1 1.1H3c.5 0 1-.5 1-1.1V12h8v1.9c0 .6.5 1.1 1.1 1.1h.9c.6 0 1.1-.5 1.1-1.1v-2c.6-.1 1-.6 1-1.2V7.2c-.1-.5-.5-1-1.1-1.1M4 8.4c0 .3-.3.6-.6.6H1.6c-.3 0-.6-.3-.6-.6v-.8c0-.3.3-.6.6-.6h1.8c.3 0 .6.3.6.6zm6 2.6H6v-1h4zM2.1 6l1.2-2.4c.2-.4.6-.6 1-.6h7.4c.4 0 .8.2 1 .6L13.9 6zM15 8.4c0 .3-.3.6-.6.6h-1.8c-.3 0-.6-.3-.6-.6v-.8c0-.3.3-.6.6-.6h1.8c.3 0 .6.3.6.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretDown(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 4h10l-5 7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretLeft(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 3v10L4 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretRight(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 13V3l7 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretSquareDownO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1H1v14h14zm-1 13H2V2h12z"/><path fill="currentColor" d="M4 6h8l-4 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretSquareLeftO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1H1v14h14zm-1 13H2V2h12z"/><path fill="currentColor" d="M10 4v8L5 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretSquareRightO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1H1v14h14zm-1 13H2V2h12z"/><path fill="currentColor" d="M5.9 12V4l5 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretSquareUpO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1H1v14h14zm-1 13H2V2h12z"/><path fill="currentColor" d="M12 10H4l4-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretUp(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 12H3l5-7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cart(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 13.1V12H4.6l.6-1.1l9.2-.9L16 4H3.7L3 1H0v1h2.2l2.1 8.4L3 13v1.5c0 .8.7 1.5 1.5 1.5S6 15.3 6 14.5S5.3 13 4.5 13H12v1.5c0 .8.7 1.5 1.5 1.5s1.5-.7 1.5-1.5c0-.7-.4-1.2-1-1.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CartO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 13.1V12H4.6l.6-1.1l9.2-.9L16 4H3.7L3 1H0v1h2.2l2.1 8.4L3 13v1.5c0 .8.7 1.5 1.5 1.5S6 15.3 6 14.5S5.3 13 4.5 13H12v1.5c0 .8.7 1.5 1.5 1.5s1.5-.7 1.5-1.5c0-.7-.4-1.2-1-1.4M4 5h10.7l-1.1 4l-8.4.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cash(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 14H2v-1h13V6h1z"/><path fill="currentColor" d="M13 4v7H1V4zm1-1H0v9h14z"/><path fill="currentColor" d="M3 6H2v3h1v1h4a2.5 2.5 0 1 1 0-5H3zm8 0V5H7a2.5 2.5 0 1 1 0 5h4V9h1V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chart(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 15h16v1H0z"/><path fill="currentColor" d="M0 0h1v16H0zm9 8L6.1 5L2 9v5h14V.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartGrid(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 9v7h16V9zm5 6H1v-1h4zm0-2H1v-1h4zm0-2H1v-1h4zm5 4H6v-1h4zm0-2H6v-1h4zm0-2H6v-1h4zm5 4h-4v-1h4zm0-2h-4v-1h4zm0-2h-4v-1h4zm1-3H0V0h1v7h15z"/><path fill="currentColor" d="M15 1.57L9.98 4.43L6.02 2.45L2 4.06v1.08l3.98-1.59l4.04 2.02L15 2.72z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartLine(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 16h16V0h-1v2.6L11 6V0h-1v6.4l-4-.9V0H5v5.7L1 8.6V0H0zm5-2H1v-1.7l4-2.9zm5 0H6V8.7l.1-.1l3.9.9zm5 0h-4V9.7h.1L15 6.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartThreeD(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 4V2L8 0L4 2v1L0 5v5l12 6l4-2V6zm-8 6.88l-3-1.5v-3.3l3 1.53zm0-4.39l-2.34-1.2L4 4.12zm4 6.39l-3-1.5V3.07l3 1.54zM5.66 2.29L8 1.12l2.34 1.17L8 3.49zM12 14.88l-3-1.5V7.07l3 1.54zm0-7.39l-2.34-1.2L12 5.12l2.34 1.17z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartTimeline(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 13v-1H1V0H0v13h5v2H0v1h16v-1h-5v-2z"/><path fill="currentColor" d="M9 7L6 4L2 8v3h14V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chat(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 14.2c0-.6 2-1.8 2-3.1c0-1.5-1.4-2.7-3.1-3.2c.7-.8 1.1-1.7 1.1-2.8C14 2.3 11.1 0 7.4 0C3.9 0 0 2.1 0 5.1c0 2.1 1.6 3.6 2.3 4.2c-.1 1.2-.6 1.7-.6 1.7L.5 12H2c1.6 0 2.9-.5 3.7-1.1v.2c0 2 2.2 3.6 5 3.6h.6c.4.5 1.7 1.4 3.4 1.4c.1-.1-.7-.5-.7-1.9M7.4 1C10.5 1 13 2.9 13 5.1s-2.6 4.1-5.8 4.1H6.1l-.1.2c-.3.4-1.5 1.2-3.1 1.5c.1-.4.1-1 .1-1.8v-.3C2 8 .9 6.6.9 5.2C.9 3 4.1 1 7.4 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Check(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.3 14.2L.2 9l1.7-2.4l4.8 3.5l6.6-8.5l2.3 1.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckCircle(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8m-.9 11.7L2.9 7.6l1.4-1.4L7 8.9L12 4l1.4 1.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckCircleO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1c3.9 0 7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7m0-1C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8"/><path fill="currentColor" d="M7.1 11.7L2.9 7.6l1.4-1.4l2.8 2.7L12 4l1.4 1.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckSquare(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 .9L12 2H0v14h14V5.5l1.7-2zM6.5 11.7L2.3 7.5l1.4-1.4l2.7 2.7L13 2.2l1.4 1.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckSquareO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 6.2V14H2V2h10.5l1-1H1v14h14V5.2z"/><path fill="currentColor" d="M7.9 10.9L3.7 6.7l1.5-1.4l2.7 2.8l6.7-6.7L16 2.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronCircleDown(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 8c0 4.4 3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8s-8 3.6-8 8m11.6-2.8L13 6.6l-5 5l-5-5l1.4-1.4L8 8.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronCircleDownO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13 6.6l-5 5l-5-5l1.4-1.4L8 8.8l3.6-3.6z"/><path fill="currentColor" d="M1 8c0-3.9 3.1-7 7-7s7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7M0 8c0 4.4 3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8s-8 3.6-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronCircleLeft(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8m2.8 11.6L9.4 13l-5-5l5-5l1.4 1.4L7.2 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronCircleLeftO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9.4 13l-5-5l5-5l1.4 1.4L7.2 8l3.6 3.6z"/><path fill="currentColor" d="M8 1c3.9 0 7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7m0-1C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronCircleRight(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 16c4.4 0 8-3.6 8-8s-3.6-8-8-8s-8 3.6-8 8s3.6 8 8 8M5.2 4.4L6.6 3l5 5l-5 5l-1.4-1.4L8.8 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronCircleRightO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m6.6 13l5-5l-5-5l-1.4 1.4L8.8 8l-3.6 3.6z"/><path fill="currentColor" d="M8 1c3.9 0 7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7m0-1C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronCircleUp(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 8c0-4.4-3.6-8-8-8S0 3.6 0 8s3.6 8 8 8s8-3.6 8-8M4.4 10.8L3 9.4l5-5l5 5l-1.4 1.4L8 7.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronCircleUpO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m3 9.4l5-5l5 5l-1.4 1.4L8 7.2l-3.6 3.6z"/><path fill="currentColor" d="M15 8c0 3.9-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7s7 3.1 7 7m1 0c0-4.4-3.6-8-8-8S0 3.6 0 8s3.6 8 8 8s8-3.6 8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDown(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8 13.1l-8-8l2.1-2.2L8 8.8l5.9-5.9L16 5.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDownSmall(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 12L1.68 5.68L3.35 4L8 8.65L12.65 4l1.67 1.68z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeft(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m2.9 8l8-8l2.2 2.1L7.2 8l5.9 5.9l-2.2 2.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeftSmall(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m4 8l6.32-6.32L12 3.35L7.35 8L12 12.65l-1.68 1.67z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRight(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.1 8l-8 8l-2.2-2.1L8.8 8L2.9 2.1L5.1 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRightSmall(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 8L5.68 1.68L4 3.35L8.65 8L4 12.65l1.68 1.67z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUp(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8 2.9l8 8l-2.1 2.2L8 7.2l-5.9 5.9L0 10.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUpSmall(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8 4l-6.32 6.32L3.35 12L8 7.35L12.65 12l1.67-1.68z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Child(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 5a2 2 0 1 1-3.999.001A2 2 0 0 1 10 5"/><path fill="currentColor" d="m12.79 10.32l-2.6-2.63A2.311 2.311 0 0 0 8.54 7H7.469c-.648 0-1.235.264-1.659.69l-2.6 2.63a.73.73 0 1 0 .998 1.003L6 9.53V16h1.5v-4h1v4H10V9.53l1.75 1.8a.73.73 0 1 0 1.041-1.009z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Circle(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleThin(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1c3.9 0 7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7m0-1C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clipboard(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 1V0H5v1H3v1H2v14h12v-1h1V1zM6 1h4v2H6zm7 14H3V3h2v1h6V3h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardCheck(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 1V0H5v1H3v1H2v14h12v-1h1V1zM6 1h4v2H6zm7 14H3V3h2v1h6V3h2z"/><path fill="currentColor" d="m7.39 12.47l-3-2.73l1.35-1.48L7.32 9.7l2.87-2.9l1.42 1.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardCross(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 1V0H5v1H3v1H2v14h12v-1h1V1zM6 1h4v2H6zm7 14H3V3h2v1h6V3h2z"/><path fill="currentColor" d="M11 8H9V6H7v2H5v2h2v2h2v-2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardHeart(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.5 7c-.6 0-1.1.6-1.5 1c-.4-.4-.9-1-1.5-1c-1.5 0-2.1 1.9-1 2.9L8 12l2.5-2.1c1.1-1 .5-2.9-1-2.9"/><path fill="currentColor" d="M11 1V0H5v1H3v1H2v14h12v-1h1V1zM6 1h4v2H6zm7 14H3V3h2v1h6V3h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardPulse(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 1V0H5v1H3v1H2v14h12v-1h1V1zM6 1h4v2H6zm7 14H3V3h2v1h6V3h2z"/><path fill="currentColor" d="M9.3 13c-.2 0-.3-.1-.4-.3l-.8-4.8l-.7 3.1c0 .1-.1.2-.3.3c-.1 0-.3 0-.4-.1l-1-1.3H4.4c-.2 0-.4-.2-.4-.4s.2-.4.4-.4H6c.1 0 .2.1.3.1l.6.8l.9-4.3c0-.2.2-.3.4-.3s.3.2.3.4l.9 5.3l.6-1.7c.1-.1.2-.2.3-.2h1.3c.2 0 .4.2.4.4s-.2.4-.4.4h-1l-1 2.9s-.2.1-.3.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardText(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 6h8v1H4zm0 2h8v1H4zm0 2h5v1H4z"/><path fill="currentColor" d="M11 1V0H5v1H3v1H2v14h12v-1h1V1zM6 1h4v2H6zm7 14H3V3h2v1h6V3h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardUser(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 1V0H5v1H3v1H2v14h12v-1h1V1zM6 1h4v2H6zm7 14H3V3h2v1h6V3h2z"/><path fill="currentColor" d="M8 6C5.5 6 6.7 9.2 6.7 9.2c.3.4.7.4.7.6c0 .3-.3.3-.6.4c-.5.1-.9-.1-1.4.8c-.3.4-.4 2-.4 2h6s-.1-1.6-.4-2c-.4-.8-.9-.7-1.4-.8c-.3 0-.6-.1-.6-.4s.3-.2.6-.6C9.3 9.2 10.5 6 8 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clock(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8m0 14c-3.3 0-6-2.7-6-6s2.7-6 6-6s6 2.7 6 6s-2.7 6-6 6"/><path fill="currentColor" d="M8 3H7v6h5V8H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Close(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.1 3.1L12.9.9L8 5.9L3.1.9L.9 3.1l5 4.9l-5 4.9l2.2 2.2l4.9-5l4.9 5l2.2-2.2l-5-4.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloseBig(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16 0l-1 .01L8 7L1 .01L0 0v1l7 7l-7 7v1h1l7-7l7 7h1v-1L9 8l7-7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloseCircle(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8m4.2 10.8l-1.4 1.4L8 9.4l-2.8 2.8l-1.4-1.4L6.6 8L3.8 5.2l1.4-1.4L8 6.6l2.8-2.8l1.4 1.4L9.4 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloseCircleO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1c3.9 0 7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7m0-1C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8"/><path fill="currentColor" d="M12.2 10.8L9.4 8l2.8-2.8l-1.4-1.4L8 6.6L5.2 3.8L3.8 5.2L6.6 8l-2.8 2.8l1.4 1.4L8 9.4l2.8 2.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloseSmall(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12.96 4.46l-1.42-1.42L8 6.59L4.46 3.04L3.04 4.46L6.59 8l-3.55 3.54l1.42 1.42L8 9.41l3.54 3.55l1.42-1.42L9.41 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloud(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 13c1.1 0 2-.9 2-2s-.9-2-2-2h-.1c0-.3.1-.6.1-1c0-2.2-1.8-4-4-4c-.8 0-1.5.2-2.2.6C7.5 3.7 6.6 3 5.5 3C4.1 3 3 4.1 3 5.5c0 .6.2 1.1.6 1.6C3.4 7 3.2 7 3 7c-1.7 0-3 1.3-3 3s1.3 3 3 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudDownload(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 10h-.1c0-.3.1-.6.1-1c0-2.2-1.8-4-4-4V1H6v3.1C5.8 4 5.7 4 5.5 4C4.1 4 3 5.1 3 6.5c0 .6.2 1.1.6 1.6C3.4 8 3.2 8 3 8c-1.7 0-3 1.3-3 3s1.3 3 3 3h11c1.1 0 2-.9 2-2s-.9-2-2-2m-6 1.4L5.1 8H7V2h2v6h1.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudDownloadO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.1 9.8v-.6c0-2.4-1.9-4.3-4.2-4.3c-.3.1-.6.1-.9.1V2H7v2.4c-.4-.3-.9-.4-1.3-.4c-1.6 0-2.9 1.3-2.9 2.9c0 .3.1.6.2.9c-1.6.2-3 1.8-3 3.6C0 13.3 1.5 15 3.3 15h10.3c1.4 0 2.4-1.5 2.4-2.7s-.8-2.3-1.9-2.5m-.5 4.2H3.3C2.1 14 1 12.7 1 11.4s1.1-2.6 2.3-2.6h.4l1.4.2l-.9-1c-.2-.3-.4-.7-.4-1.2c0-1 .8-1.8 1.8-1.8c.5 0 1 .2 1.3.6V8H5l3 4l3-4H9V6.1c.3-.1.6-.1.9-.1c1.8 0 3.2 1.5 3.2 3.3c0 .3 0 .6-.1.9l-.2.6l.8.1c.7 0 1.4.7 1.4 1.5c0 .7-.6 1.6-1.4 1.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.1 8.9v-.6c0-2.4-1.9-4.3-4.2-4.3c-.6 0-1.2.1-1.8.4c-.5-.7-1.5-1.2-2.4-1.2c-1.6 0-2.9 1.2-2.9 2.8c0 .3.1.6.2.9c-1.6.2-3 1.8-3 3.5C0 12.3 1.5 14 3.3 14h10.3c1.4 0 2.4-1.4 2.4-2.6s-.8-2.2-1.9-2.5m-.5 4.1H3.3C2.1 13 1 11.8 1 10.5S2.1 8 3.3 8h.4l1.3.3l-.8-1.2c-.2-.3-.4-.7-.4-1.1c0-1 .8-1.8 1.8-1.8c.8 0 1.5.5 1.7 1.2l.3.6l.5-.3c.5-.3 1.1-.5 1.8-.5c1.8 0 3.2 1.5 3.2 3.3c0 .3 0 .6-.1.9l-.2.6h.8c.7 0 1.4.7 1.4 1.5c0 .6-.6 1.5-1.4 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudUpload(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 10h-.1c0-.3.1-.6.1-1c0-1.6-1-3-2.4-3.6L8 1L5.5 4C4.1 4 3 5.1 3 6.5c0 .6.2 1.1.6 1.6C3.4 8 3.2 8 3 8c-1.7 0-3 1.3-3 3s1.3 3 3 3h11c1.1 0 2-.9 2-2s-.9-2-2-2M9 6v6H7V6H5.1L8 2.6L10.9 6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudUploadO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.1 10.9v-.6c0-2.4-1.9-4.3-4.2-4.3c-.3 0-.6 0-.9.1V4h2L8 0L5 4h2v1.5c-.4-.2-.9-.3-1.3-.3c-1.6 0-2.9 1.2-2.9 2.8c0 .3.1.6.2.9c-1.6.2-3 1.8-3 3.5C0 14.3 1.5 16 3.3 16h10.3c1.4 0 2.4-1.4 2.4-2.6s-.8-2.2-1.9-2.5m-.5 4.1H3.3C2.1 15 1 13.8 1 12.5S2.1 10 3.3 10h.4l1.3.3l-.8-1.2c-.2-.3-.4-.7-.4-1.1c0-1 .8-1.8 1.8-1.8c.5 0 1 .2 1.3.6V10h2V7.2c.3-.1.6-.1.9-.1c1.8 0 3.2 1.5 3.2 3.3c0 .3 0 .6-.1.9l-.2.6h.8c.7 0 1.4.7 1.4 1.5c.1.7-.5 1.6-1.3 1.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cluster(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 12a1.993 1.993 0 0 0-1.008.305L10.78 10.15a3.439 3.439 0 0 0 .74-1.993L13.09 8a1.49 1.49 0 1 0-.089-.768l-1.591.128a3.512 3.512 0 0 0-1.978-2.521L9.74 4H10a2 2 0 1 0-1.01-.265l-.27.855a3.31 3.31 0 0 0-.754-.084c-.83 0-1.59.296-2.181.789L2.791 2.291a1.5 1.5 0 1 0-1.291.71c.281-.001.544-.079.767-.214L5.26 5.791a3.446 3.446 0 0 0-.76 2.168v.203l-.66.11a2 2 0 1 0 .161.786L4 8.999l.63-.097a3.522 3.522 0 0 0 1.466 1.992l-.556 1.188a1.947 1.947 0 0 0-.539-.08h-.017a2 2 0 1 0 1.231.423l.566-1.153c.364.146.787.231 1.229.231c.847 0 1.621-.311 2.216-.824l2.176 2.124a2 2 0 1 0 1.6-.8zm-9 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2m3-4.5a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Code(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.2 14L9.7 2h1.1L6.3 14zm5.9-1h1.2L16 8l-3.7-5H11l3.8 5zm-6.2 0H3.7L0 8l3.7-5H5L1.2 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Coffee(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14 13l-4 1H4l-4-1v-1h14zm.7-10H13V2H1v5c0 1.5.8 2.8 2 3.4v.6h8v-.6c.9-.5 1.6-1.4 1.9-2.4h.1c2.3 0 2.9-2 3-3.5c.1-.8-.5-1.5-1.3-1.5M13 7V4h1.7c.1 0 .2.1.2.1s.1.1.1.3C14.8 7 13.4 7 13 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cog(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 9V7l-1.7-.6c-.2-.6-.4-1.2-.7-1.8l.8-1.6L13 1.6l-1.6.8c-.5-.3-1.1-.6-1.8-.7L9 0H7l-.6 1.7c-.6.2-1.2.4-1.7.7l-1.6-.8l-1.5 1.5l.8 1.6c-.3.5-.5 1.1-.7 1.7L0 7v2l1.7.6c.2.6.4 1.2.7 1.8L1.6 13L3 14.4l1.6-.8c.5.3 1.1.6 1.8.7L7 16h2l.6-1.7c.6-.2 1.2-.4 1.8-.7l1.6.8l1.4-1.4l-.8-1.6c.3-.5.6-1.1.7-1.8zm-8 3c-2.2 0-4-1.8-4-4s1.8-4 4-4s4 1.8 4 4s-1.8 4-4 4"/><path fill="currentColor" d="M10.6 7.9a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CogO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.2 6l-1.1-.2c-.1-.2-.1-.4-.2-.6l.6-.9l.5-.7L12.4 1l-.7.5l-.9.6c-.2-.1-.4-.1-.6-.2L10 .8L9.8 0H6.2L6 .8l-.2 1.1c-.2.1-.4.1-.6.2l-.9-.6l-.7-.4l-2.5 2.5l.5.7l.6.9c-.2.2-.2.4-.3.6L.8 6l-.8.2v3.6l.8.2l1.1.2c.1.2.1.4.2.6l-.6.9l-.5.7L3.6 15l.7-.5l.9-.6c.2.1.4.1.6.2l.2 1.1l.2.8h3.6l.2-.8l.2-1.1c.2-.1.4-.1.6-.2l.9.6l.7.5l2.6-2.6l-.5-.7l-.6-.9c.1-.2.2-.4.2-.6l1.1-.2l.8-.2V6.2zM15 9l-1.7.3c-.1.5-.3 1-.6 1.5l.9 1.4l-1.4 1.4l-1.4-.9c-.5.3-1 .5-1.5.6L9 15H7l-.3-1.7c-.5-.1-1-.3-1.5-.6l-1.4.9l-1.4-1.4l.9-1.4c-.3-.5-.5-1-.6-1.5L1 9V7l1.7-.3c.1-.5.3-1 .6-1.5l-1-1.4l1.4-1.4l1.4.9c.5-.3 1-.5 1.5-.6L7 1h2l.3 1.7c.5.1 1 .3 1.5.6l1.4-.9l1.4 1.4l-.9 1.4c.3.5.5 1 .6 1.5L15 7z"/><path fill="currentColor" d="M8 4.5C6.1 4.5 4.5 6.1 4.5 8s1.6 3.5 3.5 3.5s3.5-1.6 3.5-3.5S9.9 4.5 8 4.5m0 6c-1.4 0-2.5-1.1-2.5-2.5S6.6 5.5 8 5.5s2.5 1.1 2.5 2.5s-1.1 2.5-2.5 2.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cogs(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 7V5l-1.2-.4c-.1-.3-.2-.7-.4-1l.6-1.2l-1.5-1.3l-1.1.5c-.3-.2-.6-.3-1-.4L7 0H5l-.4 1.2c-.3.1-.7.2-1 .4l-1.1-.5l-1.4 1.4l.6 1.2c-.2.3-.3.6-.4 1L0 5v2l1.2.4c.1.3.2.7.4 1l-.5 1.1l1.4 1.4l1.2-.6c.3.2.6.3 1 .4L5 12h2l.4-1.2c.3-.1.7-.2 1-.4l1.2.6L11 9.6l-.6-1.2c.2-.3.3-.6.4-1zM3 6c0-1.7 1.3-3 3-3s3 1.3 3 3s-1.3 3-3 3s-3-1.3-3-3"/><path fill="currentColor" d="M7.5 6a1.5 1.5 0 1 1-3.001-.001A1.5 1.5 0 0 1 7.5 6M16 3V2h-.6c0-.2-.1-.4-.2-.5l.4-.4l-.7-.7l-.4.4c-.2-.1-.3-.2-.5-.2V0h-1v.6c-.2 0-.4.1-.5.2l-.4-.4l-.7.7l.4.4c-.1.2-.2.3-.2.5H11v1h.6c0 .2.1.4.2.5l-.4.4l.7.7l.4-.4c.2.1.3.2.5.2V5h1v-.6c.2 0 .4-.1.5-.2l.4.4l.7-.7l-.4-.4c.1-.2.2-.3.2-.5zm-2.5.5c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1s-.4 1-1 1m1.9 8.3c-.1-.3-.2-.6-.4-.9l.3-.6l-.7-.7l-.5.4c-.3-.2-.6-.3-.9-.4L13 9h-1l-.2.6c-.3.1-.6.2-.9.4l-.6-.3l-.7.7l.3.6c-.2.3-.3.6-.4.9L9 12v1l.6.2c.1.3.2.6.4.9l-.3.6l.7.7l.6-.3c.3.2.6.3.9.4l.1.5h1l.2-.6c.3-.1.6-.2.9-.4l.6.3l.7-.7l-.4-.5c.2-.3.3-.6.4-.9l.6-.2v-1zM12.5 14c-.8 0-1.5-.7-1.5-1.5s.7-1.5 1.5-1.5s1.5.7 1.5 1.5s-.7 1.5-1.5 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoinPiles(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5 0C7.46 0 5 .88 5 2v2c-3 .1-5 .94-5 2v6c0 1.09 2.46 2 5.5 2h.067c.732 0 1.45-.055 2.153-.16c.698 1.305 2.094 2.158 3.69 2.158a4.382 4.382 0 0 0 4.224-3.217c.209-.199.344-.442.367-.717V2c0-1.12-2.46-2-5.5-2zm-5 5C8 5 10 5.45 10 6S8 7 5.5 7S1 6.55 1 6s2-1 4.5-1m0 8c-2.71 0-4.25-.71-4.5-1v-.8a10.405 10.405 0 0 0 4.522.799c.508-.001 1.03-.03 1.544-.085c-.043.371.022.712.123 1.037c-.452.021-.967.051-1.488.051L5.49 13zm1.57-2.09c-.467.057-1.008.09-1.556.09H5.5c-2.709 0-4.249-.71-4.499-1v-.84a10.41 10.41 0 0 0 4.518.84a14.496 14.496 0 0 0 1.897-.128c-.197.306-.291.654-.342 1.015zM5.5 9C2.79 9 1.25 8.29 1 8v-.9a10.41 10.41 0 0 0 4.518.84a10.548 10.548 0 0 0 4.551-.866l-.068.366a4.397 4.397 0 0 0-1.935 1.304C7.314 8.909 6.455 9 5.575 9h-.077zm5.91 6a3.41 3.41 0 1 1 0-6.82a3.41 3.41 0 0 1 0 6.82M15 8c-.175.167-.385.3-.617.386c-.288-.244-.6-.46-.938-.634a7.615 7.615 0 0 0 1.593-.61zm0-2c-.24.31-1.61.94-4 1V6h.011a9.963 9.963 0 0 0 4.053-.855zm0-2c-.25.33-1.79 1-4.5 1h-.23a9.073 9.073 0 0 0-4.169-1H6v-.9a10.41 10.41 0 0 0 4.518.84a10.548 10.548 0 0 0 4.551-.866zm-4.5-1C8 3 6 2.55 6 2s2-1 4.5-1s4.5.45 4.5 1s-2 1-4.5 1"/><path fill="currentColor" d="M10.5 11h.5v3h1V9h-.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Coins(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.5 0A4.5 4.5 0 0 0 7 4.5c.004.261.029.513.074.758a4.647 4.647 0 0 0-1.591-.261a5.51 5.51 0 1 0 5.266 3.884c.23.077.484.099.742.099A4.49 4.49 0 0 0 11.5 0M10 10.5A4.5 4.5 0 1 1 5.5 6a4.51 4.51 0 0 1 4.5 4.499zM12.5 7h-2v-.5h.5v-3h-.5l1-1.5h.5v4.5h.5z"/><path fill="currentColor" d="M5.63 8a1.258 1.258 0 0 1 1.371 1.255L7 9.302C7 11 5.14 12 5.14 12h1.37v-.5H7V13H4v-1s2-1.27 2-2.33C6 9.3 6 9 5.58 9c-.69 0-.65 1-.65 1H4s-.23-2 1.63-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Combobox(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 4H1c-.6 0-1 .4-1 1v6c0 .6.4 1 1 1h14c.6 0 1-.4 1-1V5c0-.6-.4-1-1-1m-5 7H1V5h9zm3-2.6L11 7h4z"/><path fill="currentColor" d="M2 6h1v4H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Comment(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1C3.6 1 0 3.5 0 6.5c0 2 2 3.8 4 4.8c0 2.1-2 2.8-2 2.8c2.8 0 4.4-1.3 5.1-2.1H8c4.4 0 8-2.5 8-5.5S12.4 1 8 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentEllipsis(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1C3.6 1 0 3.5 0 6.5c0 2 2 3.8 4 4.8c0 2.1-2 2.8-2 2.8c2.8 0 4.4-1.3 5.1-2.1H8c4.4 0 8-2.5 8-5.5S12.4 1 8 1M5 8c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1s-.4 1-1 1m3 0c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1s-.4 1-1 1m3 0c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1s-.4 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentEllipsisO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 11.2c0 .1 0 .1 0 0c0 .1 0 .1 0 0M8.3 1C3.9 1 0 3.6 0 6.6c0 2 1.1 3.7 3 4.7s0 .1 0 .1c-.1 1.3-.9 1.7-.9 1.7L.3 14h2c2.5 0 4.3-1.1 5.1-1.9h.8c4.3 0 7.8-2.5 7.8-5.6S12.6 1 8.3 1m-.1 10.1H7.1l-.2.2c-.5.5-1.6 1.4-3.3 1.7c.3-.5.5-1.1.5-2v-.3l-.3-.1C2 9.7 1 8.3 1 6.6C1 4.2 4.5 2 8.3 2C12 2 15 4 15 6.6c0 2.4-3.1 4.5-6.8 4.5"/><path fill="currentColor" d="M6 7a1 1 0 1 1-2 0a1 1 0 0 1 2 0m3 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0m3 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 11.2c0 .1 0 .1 0 0c0 .1 0 .1 0 0M8.3 1C3.9 1 0 3.6 0 6.6c0 2 1.1 3.7 3 4.7s0 .1 0 .1c-.1 1.3-.9 1.7-.9 1.7L.3 14h2c2.5 0 4.3-1.1 5.1-1.9h.8c4.3 0 7.8-2.5 7.8-5.6S12.6 1 8.3 1m-.1 10.1H7l-.2.2c-.5.5-1.6 1.4-3.3 1.7c.3-.5.5-1.1.5-2v-.3l-.3-.1C1.9 9.7 1 8.3 1 6.6C1 4.2 4.5 2 8.3 2C12 2 15 4 15 6.6c0 2.4-3.1 4.5-6.8 4.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Comments(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 11.1c0-1.5-1.5-2.8-3.2-3.3c-1.3 1.5-3.9 2.4-6.4 2.4h-.5c-.1.3-.1.5-.1.8c0 2 2.2 3.6 5 3.6h.6c.4.5 1.7 1.4 3.4 1.4c0 0-.8-.4-.8-1.8c0-.6 2-1.8 2-3.1"/><path fill="currentColor" d="M13 4.6C13 2.1 10.2 0 6.6 0S0 2.1 0 4.6c0 1.7 2 3.2 3 4C3 10.4 1.6 11 1.6 11c2.3 0 3.6-1.1 4.2-1.8h.8c3.5.1 6.4-2 6.4-4.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentsO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.2 14c.6-.5 1.8-1.6 1.8-3.2c0-1.4-1.2-2.6-2.8-3.3c.5-.6.8-1.5.8-2.4C14 2.3 11.1 0 7.4 0C3.9 0 0 2.1 0 5.1c0 2.1 1.6 3.6 2.3 4.2c-.1 1.2-.6 1.7-.6 1.7L.5 12H2c1.2 0 2.2-.3 3-.7c.3 1.9 2.5 3.4 5.3 3.4h.5c.6.5 1.8 1.3 3.5 1.3h1.4l-1.1-.9s-.3-.3-.4-1.1m-3.9-.3C8 13.7 6 12.4 6 10.9v-.2c.2-.2.4-.3.5-.5h.7c2.1 0 4-.7 5.2-1.9c1.5.5 2.6 1.5 2.6 2.5s-.9 2-1.7 2.5l-.3.2v.3c0 .5.2.8.3 1.1c-1-.2-1.7-.7-1.9-1l-.1-.2zM7.4 1C10.5 1 13 2.9 13 5.1s-2.6 4.1-5.8 4.1H6.1l-.1.2c-.3.4-1.5 1.2-3.1 1.5c.1-.4.1-1 .1-1.8v-.3C2 8 .9 6.6.9 5.2C.9 3 4.1 1 7.4 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compile(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 12h4v4H1zm5 0h4v4H6zm5 0h4v4h-4zM1 7h4v4H1zm0-5h4v4H1zm5 5h4v4H6zm1-6h4v4H7zm4 6h4v4h-4zm2-7h3v3h-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compress(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.3 9.3l-5 5l1.4 1.4l5-5L8 12V8H4zm10.4-7.6L14.3.3l-4 4L9 3v4h4l-1.3-1.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompressSquare(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 0H0v12l1-1V1h10zM4 16h12V4l-1 1v10H5z"/><path fill="currentColor" d="M7 9H2l1.8 1.8L0 14.6L1.4 16l3.8-3.8L7 14zm9-7.6L14.6 0l-3.8 3.8L9 2v5h5l-1.8-1.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Connect(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 10c-.8 0-1.4.3-2 .8L6.8 9c.1-.3.2-.7.2-1s-.1-.7-.2-1L10 5.2c.6.5 1.2.8 2 .8c1.7 0 3-1.3 3-3s-1.3-3-3-3s-3 1.3-3 3v.5L5.5 5.4C5.1 5.2 4.6 5 4 5C2.4 5 1 6.3 1 8c0 1.6 1.4 3 3 3c.6 0 1.1-.2 1.5-.4L9 12.5v.5c0 1.7 1.3 3 3 3s3-1.3 3-3s-1.3-3-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ConnectO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 9c-1 0-1.8.4-2.4 1L6.9 8.3c.1-.3.1-.5.1-.8v-.4l2.9-1.3c.6.7 1.5 1.2 2.6 1.2C14.4 7 16 5.4 16 3.5S14.4 0 12.5 0S9 1.6 9 3.5v.4L6.1 5.2C5.5 4.5 4.6 4 3.5 4C1.6 4 0 5.6 0 7.5S1.6 11 3.5 11c1 0 1.8-.4 2.4-1L9 11.7v.8c0 1.9 1.6 3.5 3.5 3.5s3.5-1.6 3.5-3.5S14.4 9 12.5 9m0-8C13.9 1 15 2.1 15 3.5S13.9 6 12.5 6S10 4.9 10 3.5S11.1 1 12.5 1m-9 9C2.1 10 1 8.9 1 7.5S2.1 5 3.5 5S6 6.1 6 7.5S4.9 10 3.5 10m9 5c-1.4 0-2.5-1.1-2.5-2.5s1.1-2.5 2.5-2.5s2.5 1.1 2.5 2.5s-1.1 2.5-2.5 2.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Controller(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.951.249l.981-.195l.195.981l-.981.195zm2.926 14.717l.981-.195l.195.981l-.981.195zM.055 9.071l.981-.195l.195.981l-.981.195zm14.718-2.926l.981-.195l.195.981l-.981.195zm-3.302-4.248l.556-.831l.831.556l-.556.831zM3.139 14.441l.56-.83l.83.56l-.56.83zM1.069 3.989l.56-.83l.83.56l-.56.83zm12.478 8.31l.556-.831l.831.556l-.556.831zM8.875 1.039L9.07.058l.981.195l-.195.981zM5.953 15.745l.195-.981l.981.195l-.195.981zM.061 6.931l.195-.981l.981.195l-.195.981zm14.706 2.923l.195-.981l.981.195l-.195.981zM3.139 1.628l.831-.556l.556.831l-.831.556zm8.338 12.473l.831-.556l.556.831l-.831.556zM1.071 12.033l.831-.556l.556.831l-.831.556zM13.539 3.63l.83-.56l.56.83l-.83.56z"/><path fill="currentColor" d="M14 8a5.99 5.99 0 0 0-2.258-4.681L8.42 8.31l-.84-.59l3.32-5A5.93 5.93 0 0 0 8 1.973a6 6 0 1 0 6 6.029z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copy(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 0v3h3z"/><path fill="currentColor" d="M9 4H5V0H0v12h9zm4 0v3h3z"/><path fill="currentColor" d="M12 4h-2v9H7v3h9V8h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CopyO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 3h-3L7 0H0v13h6v3h10V6zM7 1l2 2H7zM1 12V1h5v3h3v8zm14 3H7v-2h3V4h2v3h3zm-2-9V4l2 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copyright(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1.5c3.6 0 6.5 2.9 6.5 6.5s-2.9 6.5-6.5 6.5S1.5 11.6 1.5 8S4.4 1.5 8 1.5M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8"/><path fill="currentColor" d="M9.9 10.3c-.5.4-1.2.7-1.9.7c-1.7 0-3-1.3-3-3s1.3-3 3-3c.8 0 1.6.3 2.1.9l1.1-1.1c-.8-.8-2-1.3-3.2-1.3c-2.5 0-4.5 2-4.5 4.5s2 4.5 4.5 4.5c1.1 0 2-.4 2.8-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerLowerLeft(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 16L0 0v16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerLowerRight(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 16H0L16 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerUpperLeft(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 16L16 0H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerUpperRight(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 16L0 0h16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCard(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 2v12h16V2zm15 11H1V8h14zm0-8H1V3h14z"/><path fill="currentColor" d="M10 11h3v1h-3zm-8 0h6v1H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crop(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 .7V0h-.7l-3 3H5V0H3v3H0v2h3v8h8v3h2v-3h3v-2h-3V3.7zM5 5h5.3L5 10.3zm6 6H5.7L11 5.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CrossCutlery(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.9 8.6c.6-.1 1.2-.4 1.6-.9l3.1-3.1c.4-.4.4-1 0-1.4l-.1-.2l-3 3c-.2.2-.6.2-.9 0s-.2-.6 0-.9l2.6-2.6c.2-.2.2-.6 0-.9c-.2-.2-.6-.2-.9 0l-2.6 2.6c-.2.2-.6.2-.9 0c-.2-.2-.2-.6 0-.9l3-3l-.1-.1c-.4-.4-1-.4-1.4 0L8.2 3.5c-.4.4-.7 1-.8 1.6L2.5.3c-.4-.4-1-.3-1.3 0L1 .5C-.4 1.9.1 4.7 2.5 7.1l.8.8c.4.4.9.7 1.5.8c-.5.4-.8.8-.8.8L.6 12.9c-.7.7-.7 1.9 0 2.6s1.9.7 2.6 0L6.5 12c.2-.2.7-.8 1.3-1.5c.3.4.5.6.5.6l4.3 4.3c.7.7 1.9.7 2.6 0s.7-1.9 0-2.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crosshairs(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 0h1v4L8 6l-.5-2zm1 16h-1v-4l.5-2l.5 2zM16 7.5v1h-4L10 8l2-.5zm-16 1v-1h4L6 8l-2 .5z"/><path fill="currentColor" d="M8 2.5a5.5 5.5 0 1 1 0 11A5.5 5.5 0 0 1 2.5 8a5.51 5.51 0 0 1 5.499-5.5zM8 1a7 7 0 1 0 0 14A7 7 0 0 0 8 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Css(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.1 11C5.5 11 6 10 6 10l-.8-.5s-.3.5-1 .5S3 9.1 3 7.8C3 6.6 3.6 6 4.2 6c.5 0 .9.4.9.4l.8-.6S5.2 5 4.2 5C3.1 5 2 5.9 2 7.8S2.9 11 4.1 11m4.6-1.1c-.3.1-.7 0-1-.4l-.8.5c.4.6 1 1 1.6 1c.1 0 .3 0 .4-.1c.7-.2 1.1-.8 1.1-1.6c0-1.2-.8-1.6-1.3-1.8c-.5-.3-.7-.4-.7-.8s.1-.7.6-.7c.3 0 .6.4.6.4l.8-.6c-.2-.3-.7-.8-1.4-.8C7.7 5 7 5.6 7 6.6c0 1.1.7 1.5 1.2 1.8c.6.2.8.4.8.9c0 .3 0 .6-.3.6m4 0c-.3.1-.7 0-1-.4l-.8.5c.4.6 1 1 1.6 1c.1 0 .3 0 .4-.1c.7-.2 1.1-.8 1.1-1.6c0-1.2-.8-1.6-1.3-1.8c-.5-.3-.7-.4-.7-.8s.1-.7.6-.7c.3 0 .6.4.6.4l.8-.6c-.2-.3-.7-.8-1.4-.8c-.9 0-1.6.6-1.6 1.6c0 1.1.7 1.5 1.2 1.8c.6.2.8.4.8.9c0 .3 0 .6-.3.6M0 0v16h16V0zm15 15H1V1h14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ctrl(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v16h16V0zm4.19 12C2.16 12 1 10.54 1 8s1.16-4 3.19-4h.029c.539 0 1.052.114 1.515.32l-.424.901a2.719 2.719 0 0 0-1.08-.22h-.042C2.38 5.001 2 6.631 2 8.001s.38 3 2.19 3c.497-.013.96-.145 1.366-.368l.444.898a3.943 3.943 0 0 1-1.806.47zM9 7H8v3.5a.902.902 0 0 0 1.005.499L8.999 12a1.822 1.822 0 0 1-1.998-1.428L6.999 7h-.51V6h.51V5h1v1h1v1zm2 2v3h-1V6h1v.92a2.447 2.447 0 0 1 2.005-.919l-.004 1a1.88 1.88 0 0 0-2 2.006zm4 3h-1V3h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CtrlA(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 7V6H8V5H7v1h-.5v1H7v3.56a1.823 1.823 0 0 0 2.009 1.439L9 11a.899.899 0 0 1-.998-.495L8 7zm5-4h1v9h-1zm-1 3l-.085-.001c-.773 0-1.462.358-1.911.917L11 6.001h-1v6h1v-3a1.88 1.88 0 0 1 2.006-2l-.006-1zm-8.81 6C2.16 12 1 10.54 1 8s1.16-4 3.19-4h.029c.539 0 1.052.114 1.515.32l-.424.901a2.719 2.719 0 0 0-1.08-.22h-.042C2.38 5.001 2 6.631 2 8.001s.38 3 2.19 3c.497-.013.96-.145 1.366-.368l.444.898a3.943 3.943 0 0 1-1.806.47z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cube(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0L0 2v10l8 4l8-4V2zm6.4 2.6L8.5 4.8L1.9 2.6L8 1zM1 11.4V3.3l7 2.4v9.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cubes(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 6V2L8 0L4 2v4L0 8v5l4 2l4-2l4 2l4-2V8zM8.09 1.12L11 2.56l-2.6 1.3l-2.91-1.44zM5 2.78l3 1.5v3.6l-3-1.5zm-1 11.1l-3-1.5v-3.6l3 1.5zm.28-4L1.4 8.42L4 7.12l2.88 1.44zm7.72 4l-3-1.5v-3.6l3 1.5zm.28-4L9.4 8.42l2.6-1.3l2.88 1.44z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CurlyBrackets(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.1 3.1c.2 1.3.4 1.6.4 2.9C2.5 6.8 1 7.5 1 7.5v1s1.5.7 1.5 1.5c0 1.3-.2 1.6-.4 2.9c-.3 2.1.8 3.1 1.8 3.1H6v-2s-1.8.2-1.8-1c0-.9.2-.9.4-2.9c.1-.9-.5-1.6-1.1-2.1c.6-.5 1.2-1.1 1.1-2c-.3-2-.4-2-.4-2.9C4.2 1.9 6 2 6 2V0H3.9C2.8 0 1.8 1 2.1 3.1m11.8 0c-.2 1.3-.4 1.6-.4 2.9c0 .8 1.5 1.5 1.5 1.5v1s-1.5.7-1.5 1.5c0 1.3.2 1.6.4 2.9c.3 2.1-.8 3.1-1.8 3.1H10v-2s1.8.2 1.8-1c0-.9-.2-.9-.4-2.9c-.1-.9.5-1.6 1.1-2.1c-.6-.5-1.2-1.1-1.1-2c.2-2 .4-2 .4-2.9C11.8 1.9 10 2 10 2V0h2.1c1.1 0 2.1 1 1.8 3.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cursor(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 0v13l3.31-3.47L10 16l1.37-.63L8.65 9H13z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CursorO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 2.6L10.75 9H8.29l.63 1.41l1.8 4l-.91.34l-1.88-4.3l-.5-1.11l-1 .71L5 11.07zM4 0v13l3-2.14L9.26 16l2.8-1l-2.23-5H13z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cutlery(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 .8c0-.5-.4-.8-.8-.8H12c-1.7 0-3 1.9-3 4.7v.9c0 1 .5 1.9 1.4 2.4c-.3 1.2-.4 2.5-.4 2.5v4c0 .8.7 1.5 1.5 1.5s1.5-.7 1.5-1.5v-4c0-.4-.1-1.4-.3-2.3c.2-.2.3-.4.3-.7zM7.2 0H7v3.5c0 .3-.2.5-.5.5S6 3.8 6 3.5v-3c0-.3-.2-.5-.5-.5S5 .2 5 .5v3c0 .3-.2.5-.5.5S4 3.8 4 3.5V0h-.2c-.4 0-.8.4-.8.8v3.7c0 1 .6 1.9 1.5 2.3c-.4 1.6-.5 3.7-.5 3.7v4c0 .8.7 1.5 1.5 1.5S7 15.3 7 14.5v-4c0-.5-.1-2.3-.4-3.7C7.4 6.4 8 5.5 8 4.5V.8c0-.4-.4-.8-.8-.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dashboard(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 10.1C16 5.7 12.4 2 8 2s-8 3.7-8 8.1c0 1.4.3 2.9.9 3.9h4.9c.5.6 1.3 1 2.2 1s1.7-.4 2.2-1h4.9c.6-1 .9-2.5.9-3.9M14 7v1l-4.1 3.5c0 .1.1.3.1.5c0 1.1-.9 2-2 2s-2-.9-2-2s.9-2 2-2c.3 0 .6.1.8.2L13 7zm-4-3h1v1h-1zM5 4h1v1H5zm-3 8H1v-1h1zm1-4H2V7h1zm12 4h-1v-1h1z"/><path fill="currentColor" d="M9 12a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Database(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 2.5C14 3.328 11.314 4 8 4s-6-.672-6-1.5S4.686 1 8 1s6 .672 6 1.5"/><path fill="currentColor" d="M8 5c-3.3 0-6-.7-6-1.5v3C2 7.3 4.7 8 8 8s6-.7 6-1.5v-3C14 4.3 11.3 5 8 5"/><path fill="currentColor" d="M8 9c-3.3 0-6-.7-6-1.5v3c0 .8 2.7 1.5 6 1.5s6-.7 6-1.5v-3C14 8.3 11.3 9 8 9"/><path fill="currentColor" d="M8 13c-3.3 0-6-.7-6-1.5v3c0 .8 2.7 1.5 6 1.5s6-.7 6-1.5v-3c0 .8-2.7 1.5-6 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DateInput(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 1v3h-3V1H5v3H2V1H0v15h16V1zm1 14H1V6h14z"/><path fill="currentColor" d="M3 0h1v3H3zm9 0h1v3h-1zM3 8h1v5H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Deindent(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 10.5v-6l-4 3zM0 0h16v3H0zm6 4h10v3H6zm0 4h10v3H6zm-6 4h16v3H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Del(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v16h16V0zm3 12H1V3h2a4.111 4.111 0 0 1 3.999 4.517c.013.1.02.236.02.374A4.11 4.11 0 0 1 3.005 12zm10-3H9v.012c0 .607.211 1.164.564 1.603c.252.244.601.397.986.397l.074-.002a2.4 2.4 0 0 0 1.518-.631l.708.712a3.4 3.4 0 0 1-2.225.92l-.09.002a2.393 2.393 0 0 1-1.696-.702a3.522 3.522 0 0 1-.84-2.289v-.041c0-.968.344-1.855.915-2.547a2.144 2.144 0 0 1 1.578-.623l.086-.002a2.33 2.33 0 0 1 1.641.672c.47.532.762 1.23.78 1.996zm2 3h-1V3h1z"/><path fill="currentColor" d="M3 4H2v7h1c.31 0 3-.12 3-3.5S3.12 4 3 4m7.49 2.8a1.432 1.432 0 0 0-1.339 1.192L11.93 8a1.734 1.734 0 0 0-.431-.831a1.355 1.355 0 0 0-.934-.371l-.079.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DelA(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 3h1v9h-1zM3 12H1V3h2a4.111 4.111 0 0 1 3.999 4.517c.013.1.02.236.02.374A4.11 4.11 0 0 1 3.005 12zm-1-1h1c.31 0 3-.12 3-3.5S3.12 4 3 4H2zm11-2v-.5a3.116 3.116 0 0 0-.783-2.003a2.332 2.332 0 0 0-1.732-.666l-.054-.001c-.594 0-1.132.241-1.521.631A3.978 3.978 0 0 0 8 9.001v.009c0 .881.322 1.686.854 2.306c.43.429 1.03.697 1.692.697l.089-.002a3.398 3.398 0 0 0 2.228-.922l-.712-.708a2.402 2.402 0 0 1-1.515.63l-.076.002c-.385 0-.734-.153-.99-.402A2.54 2.54 0 0 1 9 9.001zm-2.5-2.2l.066-.002c.362 0 .691.141.935.372c.209.224.361.505.427.818l-2.778.011a1.433 1.433 0 0 1 1.337-1.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DentalChair(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.5 8.2c-.3-.1-.6-.2-.8-.2H8V7h3c0-.6-.4-1-1-1H6c0 .6.4 1 1 1v1c-.5 0-1-.2-1.2-.6L4.7 5.6C4.4 5.2 4 5 3.6 5H3v-.7c0-.3-.1-.5-.2-.8l-.3-.7C2.2 2.3 1.6 2 1 2H0l5 7c.4.6 1.1 1 1.8 1H8v1H7v2h-.6c-.9 0-1.8.4-2.4 1H3v1h11v-1h-1c-.6-.6-1.5-1-2.4-1H10v-2H9v-1h1.6c.2 0 .5.1.7.2l1.7.9c.9.5 2 .5 2.9 0h.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Desktop(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 0H0v13h6v2H4v1h8v-1h-2v-2h6zM9 12H7v-1h2zm6-2H1V1h14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Diamond(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 6h4l3 8.6zm16 0h-4l-3 8.6zm-8 9L5 6h6zM4 5H0l2-3zm12 0h-4l2-3zm-6 0H6l2-3zM3.34 2H7L5 5zM9 2h4l-2 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiamondO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 2H3L0 5.5L8 15l8-9.5zM4.64 5H1.75l1.52-1.78zm1.78 0L8 3.16L9.58 5zM10 6l-2 6.68L6 6zM5.26 6l1.89 6.44L1.73 6zm5.49 0h3.53l-5.43 6.44zm.62-1l1.37-1.78L14.25 5h-2.9zM12 3l-1.44 1.81L9.1 3zM5.43 4.83L4 3h2.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Diploma(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 10.58a.371.371 0 0 0-.001-.332l-.479-.698c-.009-.013-.014-.028-.014-.045s.005-.032.014-.045l.48-.7a.371.371 0 0 0-.001-.332a.4.4 0 0 0-.236-.237l-.823-.301a.091.091 0 0 1-.06-.069V6.98a.386.386 0 0 0-.169-.299a.407.407 0 0 0-.231-.071h-.17l-.85.22a.095.095 0 0 1-.071 0l-.549-.65a.428.428 0 0 0-.63 0l-.55.65a.095.095 0 0 1-.071 0h.001l-.9-.23h-.108a.417.417 0 0 0-.234.071a.388.388 0 0 0-.168.298v.841a.092.092 0 0 1-.059.07l-.821.3a.403.403 0 0 0-.338.395c0 .06.014.117.039.167l.479.698c.009.013.014.028.014.045s-.005.032-.014.045l-.48.7a.371.371 0 0 0 .001.332a.4.4 0 0 0 .236.237l.823.301a.091.091 0 0 1 .06.069v.841a.386.386 0 0 0 .169.299a.417.417 0 0 0 .234.071h.168l.31-.07V16l1.53-2l1.47 2v-3.69l.31.08h.118a.417.417 0 0 0 .234-.071a.388.388 0 0 0 .168-.298v-.841a.092.092 0 0 1 .059-.07l.821-.3a.405.405 0 0 0 .289-.227z"/><path fill="currentColor" d="M0 1v12h8l-.11-.05a1.131 1.131 0 0 1-.49-.867V12H1V2h14v10h-1.43v.08a1.134 1.134 0 0 1-.486.868L13 13h3V1z"/><path fill="currentColor" d="M7.43 6.91a1.13 1.13 0 0 1 .486-.908A.184.184 0 0 1 8.001 6H3v1h4.43zM6.42 8H3v1h3.36a.986.986 0 0 1-.047-.837A.292.292 0 0 1 6.42 8M3 4h10v1H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiplomaScroll(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.61 8.41a5.46 5.46 0 0 1-1.454-.424A60.006 60.006 0 0 1 15.61 4.43l.1-.07l.06-.11a2.924 2.924 0 0 0-.765-3.496a2.916 2.916 0 0 0-3.459-.376l-.126.133A54.96 54.96 0 0 1 6.512 6.41a50.219 50.219 0 0 1-6.018 4.592L.1 11.25v.23a4.53 4.53 0 0 0 .7 3.655A2.83 2.83 0 0 0 3.007 16a1.997 1.997 0 0 0 1.778-.902C5.03 14.79 6.85 12.49 8.79 10.39c.268.464.476 1.003.594 1.575a6.29 6.29 0 0 1-.399 3.078L10.67 13L12 14a12.122 12.122 0 0 0-.584-3.336a5.341 5.341 0 0 0-.915-1.214c.406.346.871.643 1.372.874c.94.338 1.989.572 3.076.672L14 9.73L16 8a8.982 8.982 0 0 1-2.777.431c-.216 0-.43-.007-.642-.022zm-.45-7.23a1.89 1.89 0 0 1 .842-.194c.506 0 .966.196 1.309.516a1.926 1.926 0 0 1 .595 2.192c-.486.307-2.346 1.717-4.146 3.307a2.003 2.003 0 0 0-.668-1.298a1.596 1.596 0 0 0-1.24-.372A58.169 58.169 0 0 0 12.16 1.18M2.7 11.81a.458.458 0 0 1 .262-.082l.04.002h.068c.179.052.334.142.461.261l-.871.719a1.28 1.28 0 0 1-.119-.716a.334.334 0 0 1 .158-.183zM4 14.5a1 1 0 0 1-1.005.499a1.852 1.852 0 0 1-1.485-.54a3.432 3.432 0 0 1-.583-1.922c0-.251.027-.495.077-.73l.706-.457c-.094.14-.164.304-.199.481a2.985 2.985 0 0 0 .535 1.958l.354.44l1.7-1.4a2.396 2.396 0 0 1-.106 1.685zm.86-2.45a2.825 2.825 0 0 0-1.54-1.274c-.071-.012-.13-.016-.19-.016s-.119.004-.177.01a1.56 1.56 0 0 0-.35 0a44.978 44.978 0 0 0 3.988-3.052c.398.071.812.25 1.131.533c.297.313.48.739.48 1.209l-.002.094a54.377 54.377 0 0 0-3 3.506a2.75 2.75 0 0 0-.357-1.023z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Disc(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8m7 8c0 1.1-.2 2.1-.7 3l-2.7-1.2c.2-.6.4-1.2.4-1.8c0-2.2-1.8-4-4-4c-.5 0-.9.1-1.4.3L5.4 1.5c.6-.2 1.2-.4 1.8-.5l.3 3H8V1c3.9 0 7 3.1 7 7M8 5c1.7 0 3 1.3 3 3s-1.3 3-3 3s-3-1.3-3-3s1.3-3 3-3M1 8c0-1.1.2-2.1.7-3l2.7 1.2C4.2 6.8 4 7.4 4 8c0 2.2 1.8 4 4 4c.5 0 .9-.1 1.4-.3l1.2 2.8c-.6.2-1.2.4-1.8.5l-.3-3H8v3c-3.9 0-7-3.1-7-7"/><path fill="currentColor" d="M10 8a2 2 0 1 1-3.999.001A2 2 0 0 1 10 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Doctor(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 11.3c-1-1.9-2-1.6-3.1-1.7c.1.3.1.6.1 1c1.6.4 2 2.3 2 3.4v1h-2v-1h1s0-2.5-1.5-2.5S9 13.9 9 14h1v1H8v-1c0-1.1.4-3.1 2-3.4c0-.6-.1-1.1-.2-1.3c-.2-.1-.4-.3-.4-.6c0-.6.8-.4 1.4-1.5c0 0 .9-2.3.6-4.3h-1c0-.2.1-.3.1-.5s0-.3-.1-.5h.8C10.9.9 9.9 0 8 0C6.1 0 5.1.9 4.7 2h.8c0 .2-.1.3-.1.5s0 .3.1.5h-1c-.2 2 .6 4.3.6 4.3c.6 1 1.4.8 1.4 1.5c0 .5-.5.7-1.1.8c-.2.2-.4.6-.4 1.4v1.2c.6.2 1 .8 1 1.4c0 .7-.7 1.4-1.5 1.4S3 14.3 3 13.5c0-.7.4-1.2 1-1.4v-1.2c0-.5.1-.9.2-1.3c-.7.1-1.5.4-2.2 1.7c-.6 1.1-.9 4.7-.9 4.7h13.7c.1 0-.2-3.6-.8-4.7M6.5 2.5C6.5 1.7 7.2 1 8 1s1.5.7 1.5 1.5S8.8 4 8 4s-1.5-.7-1.5-1.5"/><path fill="currentColor" d="M5 13.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoctorBriefcase(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16 12l-1.4-6.7c-.2-.7-.9-1.3-1.7-1.3H11V2.8c0-1-.8-1.8-1.8-1.8H6.8C5.8 1 5 1.8 5 2.8V4H3.1c-.8 0-1.5.6-1.7 1.3L0 12c-.2 1 .6 2 1.7 2h12.5c1.2 0 2-1 1.8-2M6 2.8c0-.4.4-.8.8-.8h2.4c.4 0 .8.4.8.8V4H6zm5 7.2H9v2H7v-2H5V8h2V6h2v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dollar(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.2 6.8c-.1 0-.1-.1-.2-.1V3.6c1.2.1 2.2.6 2.2.6l.9-1.8c-.1 0-1.5-.8-3.1-.8V0H7v1.6c-.8.2-1.4.5-2 .9c-.6.6-1 1.4-1 2.3c0 .7.2 2.3 3 3.6v3.9c-.9-.2-2-.7-2.4-.9l-1 1.7c.2.1 1.8 1 3.4 1.2V16h1v-1.7c2.3-.3 3.6-2.1 3.6-3.8c0-1.5-1-2.7-3.4-3.7M7 6.2c-.8-.5-1-1-1-1.3c0-.4.1-.7.4-.9l.6-.3zm1 6.1V8.9c1.1.5 1.6 1.1 1.6 1.6c0 .6-.3 1.6-1.6 1.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotCircle(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 4C5.8 4 4 5.8 4 8s1.8 4 4 4s4-1.8 4-4s-1.8-4-4-4"/><path fill="currentColor" d="M8 1c3.9 0 7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7m0-1C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Download(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 10h-5.5L8 12.5L5.5 10H0v6h16zM4 14H2v-2h2z"/><path fill="currentColor" d="M10 6V0H6v6H3l5 5l5-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownloadAlt(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 14h16v2H0zm8-1l5-5h-3V0H6v8H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Drop(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0S3 8.2 3 11s2.2 5 5 5s5-2.2 5-5S8 0 8 0m.9 14.9l-.2-1c1.4-.3 2.4-1.7 2.4-3.2c0-.3-.1-1.1-.8-2.6l.9-.4c.6 1.4.8 2.4.8 3c0 2-1.3 3.8-3.1 4.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Edit(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 4s0-1-1-2s-1.9-1-1.9-1L12 2.1V0H0v16h12V8zm-9.7 7.4l-.6-.6l.3-1.1l1.5 1.5zm.9-1.9l-.6-.6l5.2-5.2c.2.1.4.3.6.5zm6.9-7l-.9 1c-.2-.2-.4-.3-.6-.5l.9-.9c.1.1.3.2.6.4M11 15H1V1h10v2.1L5.1 9L4 13.1L8.1 12L11 9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eject(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 11h14L8 1zm0 1h14v3H1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Elastic(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.7 16c-1.7 0-3.1-.8-4-2.1c-1.1-1.7-.9-4 .4-5.8C2 6.8 3.2 6 4.7 5.7c1.2-.3 2.2-1.1 2.5-2.2c.2-.8.7-1.5 1.3-2C9.4.5 10.7 0 12 0c1.1 0 2.2.4 2.9 1.2c1.5 1.6 1.5 4.2-.1 6c-.5.6-1.2 1.1-2 1.4c-1.2.5-2.2 1.6-2.6 3c-.3 1-.8 1.9-1.5 2.6c-1.1 1.2-2.6 1.8-4 1.8M12 1c-1 0-2 .4-2.8 1.2c-.5.5-.8 1-1 1.6c-.5 1.5-1.8 2.5-3.3 2.9c-1.2.2-2.2.9-3 2C.8 10.2.7 12 1.6 13.4c.6 1 1.8 1.6 3.1 1.6c1.2 0 2.4-.5 3.3-1.4c.6-.6 1.1-1.4 1.3-2.2c.4-1.7 1.6-3 3.2-3.6c.6-.2 1.2-.7 1.6-1.2c1.2-1.4 1.3-3.5.1-4.7c-.6-.6-1.4-.9-2.2-.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EllipsisCircle(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8M6 9H4V7h2zm3 0H7V7h2zm3 0h-2V7h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EllipsisCircleO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1c3.9 0 7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7m0-1C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8"/><path fill="currentColor" d="M4 7h2v2H4zm3 0h2v2H7zm3 0h2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EllipsisDotsH(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 8a2 2 0 1 1-3.999.001A2 2 0 0 1 4 8m6 0a2 2 0 1 1-3.999.001A2 2 0 0 1 10 8m6 0a2 2 0 1 1-3.999.001A2 2 0 0 1 16 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EllipsisDotsV(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 2a2 2 0 1 1-3.999.001A2 2 0 0 1 10 2m0 6a2 2 0 1 1-3.999.001A2 2 0 0 1 10 8m0 6a2 2 0 1 1-3.999.001A2 2 0 0 1 10 14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EllipsisH(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 6h4v4H0zm6 0h4v4H6zm6 0h4v4h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EllipsisV(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 0h4v4H6zm0 6h4v4H6zm0 6h4v4H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Enter(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 0v6H1v10h14V0zm8 11H7v2l-3-2.5L7 8v2h4V7h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnterArrow(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 9l7 4v-3h9V3l-3 2v2H7V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Envelope(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 3h16v2.4l-8 4l-8-4z"/><path fill="currentColor" d="m0 14l5.5-4.8L8 10.6l2.5-1.4L16 14zm4.6-5.2L0 6.5V13zm6.8 0L16 6.5V13z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 3v11h16V3zm1 4.1l3.9 2L1 12.5zm.9 5.9l4-3.5L8 10.6l2.1-1.1l4 3.5zm13.1-.5L11.1 9L15 7zm0-6.6L8 9.4L1 5.9V4h14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeOpen(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 3.7v3.7l2-1V5zM2 3.8L0 5v1.5l2 1zM11.2 2L8 0L4.8 2zM13 3H3v4.9l3.4 1.7L8 8.4l1.6 1.2L13 7.9zm3 4.6l-5.5 2.7l5.5 4.4zm-8 2L0 16h16zm-2.5.7L0 7.6v7.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeOpenO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 3.7V3h-1.5L8 0L3.4 3H2v.7L0 5v11h16V5.1zM8 1.2L10.7 3H5.2zM3 4h10v3.7L9.5 9.4L8 8.1L6.5 9.5L3 7.8zM1 5.5l1-.7v2.4l-1-.4zm0 2.4l4.6 2.3l-4.6 4zm.9 7.1L8 9.7l6.1 5.3zm13.1-.8l-4.7-4.1L15 7.8zm0-7.5l-1 .5V4.9l1 .7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Envelopes(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 14H2v-1h13V4h1z"/><path fill="currentColor" d="M14 10.77V5.29L9.32 7.47zM8.28 7.96L7 8.55l-1.28-.59L0 11.99V12l14-.01zM7 7.45l7-3.27V2H0v2.18zm-2.32.02L0 5.29v5.48z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopesO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 2H0v10h14zM5.71 8L7 8.55L8.29 8L13 11H1zM1 9.83v-4l3.64 1.63zm8.36-2.37L13 5.78v4zM13 3v1.68L7 7.45L1 4.68V3z"/><path fill="currentColor" d="M15 4v9H2v1h14V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eraser(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8.1 14l6.4-7.2c.6-.7.6-1.8-.1-2.5l-2.7-2.7c-.3-.4-.8-.6-1.3-.6H8.6c-.5 0-1 .2-1.4.6L.5 9.2c-.6.7-.6 1.9.1 2.5l2.7 2.7c.3.4.8.6 1.3.6H16v-1zm-1.3-.1s0-.1 0 0l-2.7-2.7c-.4-.4-.4-.9 0-1.3L7.5 6h-1l-3 3.3c-.6.7-.6 1.7.1 2.4L5.9 14H4.6c-.2 0-.4-.1-.6-.2L1.2 11c-.3-.3-.3-.8 0-1.1L4.7 6h1.8L10 2h1L7.5 6l3.1 3.7l-3.5 4c-.1.1-.2.1-.3.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Esc(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v16h16V0zm5 4H2v3h3v1H2v3h3v1H1V3h4zm5 6.54c-.067.511-.364.94-.782 1.186a2.426 2.426 0 0 1-1.129.276L7.996 12a6.254 6.254 0 0 1-2.038-.425l.403-.915c.435.202.945.319 1.482.319c.326 0 .643-.043.943-.125a.662.662 0 0 0 .215-.484c.07-.43-.22-.62-1.17-1c-.83-.29-2.04-.76-1.83-2.08c.072-.594.46-1.082.989-1.296a3.252 3.252 0 0 1 2.649.552L9.07 7.3a2.32 2.32 0 0 0-1.663-.368a.617.617 0 0 0-.387.547c-.08.401.14.581 1.15 1.001c.83.3 2.02.75 1.83 2.06m2.67.18c.345.176.752.279 1.183.279c.292 0 .573-.047.835-.134l.311.945c-.383.121-.823.19-1.279.19h-.001l-.089.001c-.583 0-1.124-.18-1.57-.488a2.998 2.998 0 0 1-1.061-2.524a2.866 2.866 0 0 1 1.044-2.516a3.502 3.502 0 0 1 1.72-.446c.443 0 .868.081 1.259.23l-.374.922a2.548 2.548 0 0 0-2.016.066a2.013 2.013 0 0 0-.633 1.764a2.056 2.056 0 0 0 .637 1.708z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EscA(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 12a6.268 6.268 0 0 1-2.043-.425l.403-.915c.435.202.945.319 1.482.319c.326 0 .643-.043.943-.125A.662.662 0 0 0 9 10.37c.07-.43-.22-.62-1.17-1C7 9.08 5.79 8.61 6 7.29c.072-.594.46-1.082.989-1.296a3.252 3.252 0 0 1 2.649.552l-.569.754a2.32 2.32 0 0 0-1.663-.368a.617.617 0 0 0-.387.547c-.08.401.14.581 1.15 1.001c.85.33 2 .77 1.8 2.08c-.067.511-.364.94-.782 1.186A2.42 2.42 0 0 1 7.994 12zm5.71 0l-.089.001c-.583 0-1.124-.18-1.57-.488a2.995 2.995 0 0 1-1.05-2.524a2.866 2.866 0 0 1 1.044-2.516a3.502 3.502 0 0 1 1.72-.446c.443 0 .868.081 1.259.23l-.374.922a2.548 2.548 0 0 0-2.016.066a2.013 2.013 0 0 0-.633 1.764a2.052 2.052 0 0 0 .647 1.748c.346.177.754.28 1.185.28c.292 0 .573-.047.835-.134l.331.905c-.383.121-.823.19-1.279.19h-.012zM5 4V3H1v9h4v-1H2V8h3V7H2V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Euro(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.89 3a5.45 5.45 0 0 1 3.127 1.011L14 1.69a7.369 7.369 0 0 0-3.129-.686a7.482 7.482 0 0 0-7.323 5.947L2 7v1h1.41v.5a3.848 3.848 0 0 0 0 .512L1.999 9v1h1.54c.882 3.353 3.805 5.818 7.331 5.999a7.42 7.42 0 0 0 3.175-.708L14 13a5.426 5.426 0 0 1-3.108 1a5.909 5.909 0 0 1-5.28-3.959L12 10V9H5.41a3.224 3.224 0 0 1 .001-.511L5.41 8H12V7H5.6c.678-2.325 2.788-3.996 5.29-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Exchange(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 5v2H3v2L0 6l3-3v2zM0 12v-2h13V8l3 3l-3 3v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Exclamation(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 0h4v4l-1 7H7L6 4zm4 14a2 2 0 1 1-3.999.001A2 2 0 0 1 10 14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationCircle(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8m1 13H7v-2h2zm0-3H7V3h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationCircleO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1c3.9 0 7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7m0-1C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8"/><path fill="currentColor" d="M7 3h2v7H7zm0 8h2v2H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Exit(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 6h-1.7c-.2 0-.4-.1-.6-.2l-1.3-1.3c-.2-.3-.6-.5-1.1-.5H9c1.1 0 2-.9 2-2s-.9-2-2-2s-2 .9-2 2c0 .7.4 1.4 1 1.7l-.2.3h-2c-1.1 0-2.3.5-3 1.5l-.6.8c-.4.4-.2 1 .2 1.3c.4.2.9.1 1.2-.3l.5-.7c.3-.4.7-.6 1.2-.6h.8l-.7 1.6c-.3.6-.4 1.2-.4 1.9v2c0 .3-.2.5-.5.5H2c-.6 0-1 .4-1 1s.4 1 1 1h3.5c.8 0 1.5-.7 1.5-1.5V10l3.8 4.5c.6.9 1.7 1.5 2.8 1.5h.9L9.1 9.3c-.3-.4-.2-.8 0-1.3l.6-1.5l.7.8c.4.4 1 .7 1.6.7h2c.6 0 1-.4 1-1s-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExitO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0c1.1 0 2 .9 2 2c0 .9-.6 1.7-1.5 1.9V4c.4 0 .7.2 1 .5l1.3 1.3c.1.1.3.2.5.2H15V0z"/><path fill="currentColor" d="M11.8 14.5L8 10v2.5c0 .8-.7 1.5-1.5 1.5H3c-.6 0-1-.4-1-1s.4-1 1-1h2.5c.3 0 .5-.2.5-.5v-2c0-.7.1-1.3.4-2L7.1 6h-.8c-.5 0-.9.2-1.2.6l-.5.7c-.2.4-.7.5-1.2.3c-.4-.3-.6-.9-.2-1.3l.6-.8C4.5 4.5 5.7 4 6.9 4h2l.1-.3c-.6-.3-1-1-1-1.7c0-1.1.9-2 2-2H3v4.9l-.6.8c-.3.4-.5.9-.4 1.5c.1.5.4 1 .9 1.3V11c-1.1 0-2 .9-2 2s.9 2 2 2v1h11.6c-1.1 0-2.1-.6-2.7-1.5"/><path fill="currentColor" d="m11.4 7.3l-.7-.8l-.6 1.5c-.2.5-.3.9 0 1.3l4.9 6.1V8h-2.1c-.6 0-1.1-.2-1.5-.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Expand(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1h-4l1.3 1.3l-4.5 4.5l1.4 1.4l4.5-4.5L15 5zM6.8 7.8l-4.5 4.5L1 11v4h4l-1.3-1.3l4.5-4.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExpandFull(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.3 6.7l1.4-1.4l-3-3L5 1H1v4l1.3-1.3zm1.4 4L5.3 9.3l-3 3L1 11v4h4l-1.3-1.3zm4-1.4l-1.4 1.4l3 3L11 15h4v-4l-1.3 1.3zM11 1l1.3 1.3l-3 3l1.4 1.4l3-3L15 5V1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExpandSquare(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 2H2v9l1-1V3h7zM5 14h9V5l-1 1v7H6z"/><path fill="currentColor" d="M16 0h-5l1.8 1.8l-4.5 4.5l1.4 1.4l4.5-4.5L16 5zM7.7 9.7L6.3 8.3l-4.5 4.5L0 11v5h5l-1.8-1.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExternalBrowser(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 10L8.1 6.8L4.8 10H7v1.8c0 1.7-.9 4.2-4 4.2c4.8 0 6-1.4 6-4.3V10z"/><path fill="currentColor" d="M0 0v13h6v-1H1V3h14v9h-5v1h6V0zm2 2H1V1h1zm11 0H3V1h10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExternalLink(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 16V5l-1 1v9H1V3h9l1-1H0v14z"/><path fill="currentColor" d="M16 0h-5l1.8 1.8L6 8.6L7.4 10l6.8-6.8L16 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eye(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 3.9C1.3 3.9 0 9 0 9s2.2 4.1 7.9 4.1s8.1-4 8.1-4s-1.3-5.2-8-5.2M5.3 5.4c.5-.3 1.3-.3 1.3-.3s-.5.9-.5 1.6c0 .7.2 1.1.2 1.1L5.2 8s-.3-.5-.3-1.2c0-.8.4-1.4.4-1.4m2.6 6.7c-4.1 0-6.2-2.3-6.8-3.2c.3-.7 1.1-2.2 3.1-3.2c-.1.4-.2.8-.2 1.3c0 2.2 1.8 4 4 4s4-1.8 4-4c0-.5-.1-.9-.2-1.3c2 .9 2.8 2.5 3.1 3.2c-.7.9-2.8 3.2-7 3.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeSlash(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12.9 5.2l-.8.8c1.7.9 2.5 2.3 2.8 3c-.7.9-2.8 3.1-7 3.1c-.7 0-1.2-.1-1.8-.2l-.8.8c.8.3 1.7.4 2.6.4c5.7 0 8.1-4 8.1-4s-.6-2.4-3.1-3.9"/><path fill="currentColor" d="M12 7.1c0-.3 0-.6-.1-.8L7.1 11c.3 0 .6.1.9.1c2.2 0 4-1.8 4-4M15.3 0l-4.4 4.4C10.1 4.2 9.1 4 8 4C1.3 4 0 9.1 0 9.1s1 1.8 3.3 3L0 15.3v.7h.7L16 .7V0zM4 11.3C2.4 10.6 1.5 9.5 1.1 9c.3-.7 1.1-2.2 3.1-3.2c-.1.4-.2.8-.2 1.3c0 1.1.5 2.2 1.3 2.9zm2.2-3.4l-1 .2s-.3-.5-.3-1.2c0-.8.4-1.5.4-1.5c.5-.3 1.3-.3 1.3-.3s-.5.9-.5 1.7c-.1.7.1 1.1.1 1.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eyedropper(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1c-1.8-1.8-3.7-.7-4.6.1c-.4.4-.7.9-.7 1.5c0 1.1-1.1 1.8-2.1 1.5L7.5 4l-.7.8l.7.7l-6 6l-.8 2.3l-.7.7L1.5 16l.8-.8l2.3-.8l6-6l.7.7l.7-.6l-.1-.2c-.3-1 .4-2.1 1.5-2.1c.6 0 1.1-.2 1.4-.6c.9-.9 2-2.8.2-4.6M3.9 13.6l-2 .7l-.2.1l.1-.2l.7-2l5.8-5.8l1.5 1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Facebook(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.2 16V8.5h-2V5.8h2V3.5C7.2 1.7 8.4 0 11.1 0c1.1 0 1.9.1 1.9.1l-.1 2.5h-1.7c-1 0-1.1.4-1.1 1.2v2H13l-.1 2.7h-2.8V16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FacebookSquare(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v16h16V0zm12.9 8.4h-2.1V14H8.7V8.4H7.2v-2h1.5V4.7c0-1.5.9-2.7 2.9-2.7c.8 0 1.4.1 1.4.1V4h-1.3c-.7 0-.8.3-.8.9v1.5H13z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Factory(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.4 1.3c-.6.3-.8 1.1-.4 1.5c.5-.9 1.3-.6 2.5.4c.8.7 1.9.1 1.9.1s.2 1.2 1.7 1.4c1.7.2 2.3-.8 2.3-.8s.4 1 1.9.4c1.1-.4.7-1.1.7-1.1s1 0 1-.7c0-.9-1.1-.8-1.1-.8s.2-1-.9-1.1c-1-.1-1.3.5-1.3.5S12.4 0 10.9 0C9.5 0 9 1.3 9 1.3S8.6.7 7.4.7c-.9 0-1.3.7-1.3.7S5 .9 4.4 1.3"/><path fill="currentColor" d="M12 12.1V10l-4 2.1V10H5.6L5 3H3l-.6 7H0v6h16v-6zM6 14H2v-2h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Family(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.5 7.5a1.5 1.5 0 1 1-3.001-.001A1.5 1.5 0 0 1 9.5 7.5M14.27 4h-2.54A1.73 1.73 0 0 0 10 5.73V9a1 1 0 0 0 1 1v6h4v-6a1 1 0 0 0 1-1V5.73A1.73 1.73 0 0 0 14.27 4"/><path fill="currentColor" d="M15 2a2 2 0 1 1-3.999.001A2 2 0 0 1 15 2M4.27 5H1.73a1.73 1.73 0 1 0 .001 3.461A1.73 1.73 0 0 0 1.73 5A1.73 1.73 0 0 0 0 6.73V9a1 1 0 0 0 1 1l-1 3h1v3h4v-3h1l-1-3a1 1 0 0 0 1-1V6.73A1.73 1.73 0 0 0 4.27 5"/><path fill="currentColor" d="M5 3a2 2 0 1 1-3.999.001A2 2 0 0 1 5 3m2 10v3h2v-3a1 1 0 0 0 1-1v-1.54A1.46 1.46 0 0 0 8.54 9H7.46A1.46 1.46 0 0 0 6 10.46V12a1 1 0 0 0 1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastBackward(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 15V1L9 8zm-7 0V1L2 8zM0 1h2v14H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastForward(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 1v14l7-7zm7 0v14l7-7zm7 0h2v14h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Female(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 2a2 2 0 1 1-3.999.001A2 2 0 0 1 10 2"/><path fill="currentColor" d="M10 8V6.5l1.8 1.8c.3.3.7.3 1 0s.3-.8 0-1l-2.6-2.6c-.4-.5-1-.7-1.7-.7h-1c-.7 0-1.3.2-1.7.7L3.2 7.3c-.3.3-.3.8 0 1c.3.3.7.3 1 0L6 6.5V8l-4 5h4v3h4v-3h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func File(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 5h5v11H2V0h7zm1-1V0l4 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileAdd(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 15H2V1h6v4h4v1h1V4L9 0H1v16h12v-2h-1zM9 1l3 3H9z"/><path fill="currentColor" d="M13 7h-2v2H9v2h2v2h2v-2h2V9h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileCode(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0H2v16h12V4zM9 5h4v10H3V1h6zm1-1V1l3 3z"/><path fill="currentColor" d="M6.2 13h-.7l-2-2.5l2-2.5h.7l-2 2.5zm3.6 0h.7l2-2.5l-2-2.5h-.7l2 2.5zm-3.1 1h.6l2.1-7h-.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileFont(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0H2v16h12V4zM9 5h4v10H3V1h6zm1-1V1l3 3z"/><path fill="currentColor" d="M5 7v2h2v5h2V9h2V7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileMovie(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0H2v16h12V4zM9 5h4v10H3V1h6zm1-1V1l3 3z"/><path fill="currentColor" d="M10 10V8H4v5h6v-2l2 2V8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0H2v16h12V4zM9 5h4v10H3V1h6zm1-1V1l3 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilePicture(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0H2v16h12V4zM9 5h4v10H3V1h6zm1-1V1l3 3z"/><path fill="currentColor" d="M4 11.5V14h8v-1.7s.1-1.3-1.3-1.5c-1.3-.2-1.5.4-2.5.5c-.8 0-.6-1.3-2.2-1.3c-1.2 0-2 1.5-2 1.5m8-3a1.5 1.5 0 1 1-3.001-.001A1.5 1.5 0 0 1 12 8.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilePresentation(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0H2v16h12V4zm3 15H3V1h6v4h4zM10 4V1l3 3z"/><path fill="currentColor" d="M9 6H7v1H4v6h2v1h1v-1h2v1h1v-1h2V7H9zm2 2v4H5V8z"/><path fill="currentColor" d="M7 9v2l2-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileProcess(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 0H5v6h.7l.2.7l.1.1V1h5v4h4v9H9l.3.5l-.5.5H16V4zm0 4V1l3 3zm-6.5 7.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/><path fill="currentColor" d="M7.9 12.4L9 12v-1l-1.1-.4c-.1-.3-.2-.6-.4-.9l.5-1l-.7-.7l-1 .5c-.3-.2-.6-.3-.9-.4L5 7H4l-.4 1.1c-.3.1-.6.2-.9.4l-1-.5l-.7.7l.5 1.1c-.2.3-.3.6-.4.9L0 11v1l1.1.4c.1.3.2.6.4.9l-.5 1l.7.7l1.1-.5c.3.2.6.3.9.4L4 16h1l.4-1.1c.3-.1.6-.2.9-.4l1 .5l.7-.7l-.5-1.1c.2-.2.3-.5.4-.8m-3.4 1.1c-1.1 0-2-.9-2-2s.9-2 2-2s2 .9 2 2s-.9 2-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileRefresh(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0H2v16h12V4zm3 15H3V1h6v4h4zM10 4V1l3 3z"/><path fill="currentColor" d="M4.7 7.7L4 7v3h3L5.8 8.8C6.2 8 7.1 7.5 8 7.5c1.4 0 2.5 1.1 2.5 2.5H12c0-2.2-1.8-4-4-4c-1.3 0-2.5.7-3.3 1.7m5.1 4.1c-.5.5-1.1.8-1.8.7c-1 0-1.9-.6-2.3-1.5H4.1c.4 1.7 2 3 3.8 3c1.1 0 2.1-.5 2.8-1.2L12 14v-3H9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileRemove(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 15H2V1h6v4h4v2.59l1-1V4L9 0H1v16h12v-2.59l-1-1zM9 1l3 3H9z"/><path fill="currentColor" d="m15 8l-1-1l-2 2l-2-2l-1 1l2 2l-2 2l1 1l2-2l2 2l1-1l-2-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileSearch(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 13.47V15H2V1h6v4h4v.56c.386.229.716.504.996.825L13 4L9 0H1v16h12v-1.53zM9 1l3 3H9z"/><path fill="currentColor" d="m14.78 12.72l-1.92-1.92a.727.727 0 0 0-.325-.179a3.014 3.014 0 0 0 .468-1.618a3 3 0 1 0-1.371 2.52c.02.136.083.248.169.337l1.92 1.92a.75.75 0 0 0 1.059-1.061zM10 11a2 2 0 1 1-.001-3.999A2 2 0 0 1 10 11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileSound(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.4 10.5c0 1.2-.4 2.2-1 3l.4.5c.7-.9 1.2-2.1 1.2-3.5s-.5-2.6-1.2-3.5l-.4.5c.6.8 1 1.9 1 3"/><path fill="currentColor" d="m9.9 8l-.4.5c.4.5.7 1.2.7 2s-.3 1.5-.7 2l.4.5c.5-.6.8-1.5.8-2.5s-.3-1.8-.8-2.5"/><path fill="currentColor" d="m9.1 9l-.4.5c.2.3.3.6.3 1s-.1.7-.3 1l.4.5c.3-.4.5-.9.5-1.5S9.4 9.4 9.1 9"/><path fill="currentColor" d="M10 0H2v16h12V4zM9 5h4v10H3V1h6zm1-1V1l3 3z"/><path fill="currentColor" d="M6 9H4v3h2l2 2V7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileStart(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0H2v16h12V4zm3 15H3V1h6v4h4zM10 4V1l3 3z"/><path fill="currentColor" d="M5 6v6l6-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileTable(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0H2v16h12V4zM9 5h4v10H3V1h6zm1-1V1l3 3z"/><path fill="currentColor" d="M4 7v6h8V7zm2 5H5v-1h1zm0-2H5V9h1zm3 2H7v-1h2zm0-2H7V9h2zm2 2h-1v-1h1zm0-2h-1V9h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileText(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0v4h4z"/><path fill="currentColor" d="M9 0H2v16h12V5H9zm3 12H4v-1h8zm0-2H4V9h8zm0-3v1H4V7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileTextO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0H2v16h12V4zM9 5h4v10H3V1h6zm1-1V1l3 3z"/><path fill="currentColor" d="M4 7h8v1H4zm0 2h8v1H4zm0 2h8v1H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileTree(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 10V6H5v1H3V4h9V0H0v4h2v10h3v2h11v-4H5v1H3V8h2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileTreeSmall(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 12v2h11V9H5v2H3V7h9V2H0v5h2v5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileTreeSub(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 11v1H7v-2h5V6H4v1H3V5h6V1H0v4h2v3h2v2h2v3h2v2h8v-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileZip(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0H2v16h12V4zM9 15H5v-2.8l.7-2.2h2.4l.9 2.2zm4 0h-3v-3L9 9H7V8H5v1l-1 3v3H3V1h4v1h2v1H7v1h2v1h4zM10 4V1l3 3z"/><path fill="currentColor" d="M5 6h2v1H5zm0-4h2v1H5zm0 2h2v1H5zm2 1h2v1H7zm0 2h2v1H7zm-1 5h2v2H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fill(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 14.5c.468-2.207.985-4.05 1.604-5.846c.411 1.796.928 3.638 1.337 5.521C16 15.328 15.329 16 14.5 16s-1.5-.672-1.5-1.5M8 1L6.56 2.44l-2-2a1.539 1.539 0 0 0-2.121 0a1.496 1.496 0 0 0 .001 2.119l2 2L0 8.999l7 7l8-8zm0 1.41L13.59 8H2.41l2.75-2.75a.49.49 0 0 0 .669-.672l.721-.718l1.54 1.53a.502.502 0 0 0 .71-.71L7.27 3.15zm-4.85-.56a.5.5 0 0 1 .355-.854c.138 0 .263.055.355.144l2 2l-.71.71z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Film(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v16h1v-1h1v1h12v-1h1v1h1V0zm2 14H1v-1h1zm0-2H1v-1h1zm0-2H1V9h1zm0-2H1V7h1zm0-2H1V5h1zm0-2H1V3h1zm0-2H1V1h1zm11 13H3V9h10zm0-8H3V1h10zm2 7h-1v-1h1zm0-2h-1v-1h1zm0-2h-1V9h1zm0-2h-1V7h1zm0-2h-1V5h1zm0-2h-1V3h1zm0-2h-1V1h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Filter(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 2h14v2L9 9v7l-2-2V9L1 4zm0-2h14v1H1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fire(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.9 15.8S1 15.4 1 10.1C1 6 4.1 3.6 4.1 3.6S5.4 5 6.4 5.5C7.4 6.1 7.8 0 7.8 0S15 3.9 15 9.8c0 6.1-4 5.9-4 5.9s1.8-2.4 1.8-5.2c0-3-3.9-6.7-3.9-6.7s-.5 4.4-2.1 5c-1.6-.9-2.5-2.3-2.5-2.3s-3.7 5.8.6 9.3"/><path fill="currentColor" d="M8.2 16.1c-2-.1-3.7-1.4-3.7-3.2s.7-2.6.7-2.6s.5 1 1.1 1.5s1.8.8 2.4.1c.6-.6.8-2.3.8-2.3s1.4 1.1 1.2 3c-.1 2-.9 3.5-2.5 3.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flag(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 2c0-1.1-.9-2-2-2S0 .9 0 2c0 .7.4 1.4 1 1.7V16h2V3.7c.6-.3 1-1 1-1.7m0 2s1-3 3.6-3c2.7 0 2.3 1 4.4 1c2.4 0 4-1 4-1v8s-.7 2-4 2c-2.2 0-2.3-1-5-1c-2.3 0-3 2-3 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlagCheckered(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 0C.9 0 0 .9 0 2c0 .7.4 1.4 1 1.7V16h2V3.7c.6-.3 1-1 1-1.7c0-1.1-.9-2-2-2m10 2c-2.1 0-1.8-1-4.4-1S4 4 4 4v8s.7-2 3-2c2.7 0 2.8 1 5 1c3.3 0 4-2 4-2V1s-1.6 1-4 1m3 2.5c-.2.2-.8.4-2 .6V2.9c.8-.1 1.5-.2 2-.4zM5 7.9V5.3c.4-.6 1.1-1.1 2-1.1V2.1c.2-.1.4-.1.6-.1c1.2 0 1.6.2 2.1.4c.1.1.2.2.3.2v2.2c.5.2 1.1.4 2 .4c.4 0 .7 0 1-.1v2.6c-.3 0-.6.1-1 .1c-1.1 0-1.5-.2-2-.5v2.3C9.3 9.3 8.5 9 7 9V6.8c-.9.2-1.5.6-2 1.1m8 2V7.7c1.1-.2 1.7-.6 2-.8v1.8c-.2.3-.7 1-2 1.2"/><path fill="currentColor" d="M10 7.2V4.8s-1.2-.6-3-.6v2.6c1.7-.4 3 .4 3 .4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlagO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 2c0-1.1-.9-2-2-2S0 .9 0 2c0 .7.4 1.4 1 1.7V16h2V3.7c.6-.3 1-1 1-1.7m3.6 0c1.2 0 1.6.2 2.1.4c.5.3 1.1.6 2.3.6s2.2-.2 3-.5v6.3c-.2.3-.9 1.2-3 1.2c-.9 0-1.3-.2-1.9-.4C9.4 9.3 8.6 9 7 9c-.8 0-1.5.2-2 .5V4.2C5.2 3.7 6 2 7.6 2M16 1s-1.6 1-4 1c-2.1 0-1.8-1-4.4-1S4 4 4 4v8s.7-2 3-2c2.7 0 2.8 1 5 1c3.3 0 4-2 4-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flash(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16 8l-2.2-1.6L14.9 4l-2.7-.2l-.2-2.7l-2.4 1.1L8 0L6.4 2.2L4 1.1l-.2 2.7l-2.7.2l1.1 2.4L0 8l2.2 1.6L1.1 12l2.7.2l.2 2.7l2.4-1.1L8 16l1.6-2.2l2.4 1.1l.2-2.7l2.7-.2l-1.1-2.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flask(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 16h12l-4-8V1h1V0H5v1h1v7zM9 1v7.2l1.9 3.8H5.1L7 8.2V1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlightLanding(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.64 7c-.71-.2-1.89-.43-3.23-.67L6.59 2.09a2.268 2.268 0 0 0-.746-.544L4.65 1c-.09 0-.15 0-.1.11S6 4 6.84 5.7c-1.84-.29-3.5-.53-4.23-.63a.917.917 0 0 1-.608-.406L1.28 3.59a.925.925 0 0 0-.474-.358L0 3l.61 3.26c.067.34.318.609.644.699C2.58 7.34 6.07 8.3 8.78 8.88c6 1.28 6.8 1.28 7.12.91S15.23 7.41 13.64 7M0 13h16v1H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlightTakeoff(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.57 2.26c-.65.29-1.66.85-2.8 1.5L4.31 3a2.172 2.172 0 0 0-.916.064L2.209 3.4c-.1 0-.1.1 0 .14l4.56 2c-1.54.92-2.91 1.76-3.51 2.14a.858.858 0 0 1-.726.088L1.339 7.39a.864.864 0 0 0-.586.002l-.754.308l2.52 2.1a.879.879 0 0 0 .926.128C4.649 9.39 7.819 7.93 10.179 6.7c5.24-2.78 5.82-3.26 5.82-3.7c0-.69-2-1.4-3.43-.74zM0 13h16v1H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlipH(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 15l6-5l-6-4.9zm9-4.9l6 4.9V5zm5 2.8l-3.4-2.8l3.4-3zM7 5h1v1H7zm0-2h1v1H7zm0 4h1v1H7zm0 2h1v1H7zm0 2h1v1H7zm0 2h1v1H7zm0 2h1v1H7z"/><path fill="currentColor" d="M7.5 1c1.3 0 2.6.7 3.6 1.9L10 4h3V1l-1.2 1.2C10.6.8 9.1 0 7.5 0C5.6 0 3.9 1 2.6 2.9l.8.6C4.5 1.9 5.9 1 7.5 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlipV(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1 1l5 6l4.94-6zm4.94 9L1 16h10zm-2.82 5l2.83-3.44l3 3.44zM10 8h1v1h-1zm2 0h1v1h-1zM8 8h1v1H8zM6 8h1v1H6zM4 8h1v1H4zM2 8h1v1H2zM0 8h1v1H0z"/><path fill="currentColor" d="M15 8.47a4.807 4.807 0 0 1-1.879 3.632L12 11v3h3l-1.18-1.18A5.757 5.757 0 0 0 16 8.478V8.47a6.062 6.062 0 0 0-2.884-4.905l-.596.805a5.095 5.095 0 0 1 2.479 4.087z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folder(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 15H0V4h1l1-2h4l1 2h9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderAdd(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 6V4H7L6 2H2L1 4H0v11h14v-1H1V5h.62l1-2h2.57l1.19 2H13v1z"/><path fill="currentColor" d="M14 7h-2v2h-2v2h2v2h2v-2h2V9h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 4L6 2H2L1 4H0v11h16V4zm8 10H1V5h.6l1-2h2.6l1.2 2H15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOpen(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 6V4H7L6 2H2L1 4H0v9.5L3 6z"/><path fill="currentColor" d="M3.7 7L.5 15h12.8l2.5-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOpenO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 6V4H7L6 2H2L1 4H0v11h14l2-9zm.9 1l-1.6 7l-11.9-.1L3.7 7zM1 5h.6l1-2h2.6l1.2 2H13v1H3l-2 5.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderRemove(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 12.41V14H1V5h.62l1-2h2.57l1.19 2H13v2.59l1-1V4H7L6 2H2L1 4H0v11h14v-1.59z"/><path fill="currentColor" d="m16 8l-1-1l-2 2l-2-2l-1 1l2 2l-2 2l1 1l2-2l2 2l1-1l-2-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderSearch(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 13.47V14H1V5h.62l1-2h2.57l1.19 2H13v.91c.385.179.716.407 1.001.681L14 4H7L6 2H2L1 4H0v11h14v-.53z"/><path fill="currentColor" d="m15.78 12.72l-1.92-1.92a.727.727 0 0 0-.325-.179a3.014 3.014 0 0 0 .468-1.618a3 3 0 1 0-1.371 2.52c.02.136.083.248.169.337l1.92 1.92a.75.75 0 0 0 1.059-1.061zM11 11a2 2 0 1 1-.001-3.999A2 2 0 0 1 11 11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Font(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 16h3L9 0H7L1 16h3l1.9-5h4.2zM6.7 9L8 5.4L9.3 9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Form(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 2v2H6V2zm1-1H5v4h11zM0 1h4v4H0zm15 6v2H6V7zm1-1H5v4h11zM0 6h4v4H0zm15 6v2H6v-2zm1-1H5v4h11zM0 11h4v4H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Forward(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 1v14l8-7zm8 0v14l8-7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FrownO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1c3.9 0 7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7m0-1C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8"/><path fill="currentColor" d="M7 6a1 1 0 1 1-2 0a1 1 0 0 1 2 0m4 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0m.3 6.3c-.7-1.1-2-1.8-3.3-1.8s-2.6.7-3.3 1.8l-.8-.6c.9-1.4 2.4-2.2 4.1-2.2s3.2.8 4.1 2.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Funcion(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 0S7.9 0 7.3 3l-.4 2H5l-.5 1h2.2l-1.4 7c-.4 2-1.9 2-1.9 2h-1L2 16h3s2.1 0 2.7-3l1.4-7h2.4l.5-1H9.3l.4-2c.4-2 1.8-2 1.8-2h1l.5-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Funnel(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 11h4v4H6zm7.6-6L16 1H0l2.4 4zM3 6l2.4 4h5.2L13 6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gamepad(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.16 2a9.844 9.844 0 0 1-4.149 1a9.968 9.968 0 0 1-4.229-1.026C1.171 2 .001 3.17.001 5.84v6A2.19 2.19 0 0 0 2.191 14h.232a2.19 2.19 0 0 0 2.074-1.485C4.802 11.6 5.642 10 6.582 10h2.84c.94 0 1.78 1.6 2.08 2.5A2.194 2.194 0 0 0 13.58 14h.232c1.21 0 2.19-.98 2.19-2.19v-6c0-2.64-1.17-3.81-3.84-3.81zM5 7H4v1H3V7H2V6h1V5h1v1h1zm5.06 1.11a1.06 1.06 0 1 1 .001-2.121a1.06 1.06 0 0 1-.001 2.121M13 8a1 1 0 1 1 0-2a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gavel(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.4 4.1c-.4-.4-.4-.9-.1-1.2L8.9.3c.3-.3.8-.3 1.2 0l.1.1c.3.3.3.8 0 1.2L7.6 4.1c-.3.3-.9.3-1.2 0M12 9.7c-.4-.4-.4-.9-.1-1.3l2.6-2.6c.3-.3.8-.3 1.2 0l.1.1c.3.3.3.8 0 1.2l-2.6 2.6c-.4.3-.9.3-1.2 0m-2-2L8.3 6c-.4-.4-.4-1 0-1.4l2.3-2.3c.4-.4 1-.4 1.4 0L13.7 4c.4.4.4 1 0 1.4l-2.3 2.3c-.4.4-1 .4-1.4 0m-6 6.5c.6-.6 4-5.6 4.5-5.3c.4.2 1-.5 1-.5L7.6 6.5s-.7.6-.5 1c.3.5-4.7 3.9-5.3 4.5c0 0-2.8 2.2-1.4 3.6S4 14.2 4 14.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gift(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.1 5c2-.3 3.9-1.1 2.2-3.6c-.7-1-1.4-1.4-2-1.4c-1 0-1.7 1.1-2.3 2.2C7.4 1.1 6.7 0 5.7 0c-.6 0-1.3.4-2 1.4c-1.8 2.5.2 3.3 2.2 3.6H0v3h16V5zm.2-4c.1 0 .5.1 1.2 1c.5.7.6 1.1.5 1.3c-.2.3-1.3.7-3.3.8c0-.2-.1-.4-.2-.6C9.1 2.1 9.8 1 10.3 1M4 3.3c-.1-.2 0-.6.5-1.3c.7-.9 1.1-1 1.2-1c.5 0 1.2 1.1 1.8 2.5c-.1.2-.2.4-.2.6c-2-.1-3.1-.5-3.3-.8M7 7V5h2v2zm2 8H7V9H1v7h14V9H9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Glass(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 15H9V7l6-7H0l6 7v8H4c-2 0-2 1-2 1h11s0-1-2-1m1.9-14l-1.8 2H3.9L2.2 1zM7 15V7h1v8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Glasses(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.5 7h-.7c-.4-1.2-1.5-2-2.8-2s-2.4.9-2.8 2.1c-.3-.4-.7-.6-1.2-.6s-.9.2-1.2.6C6.4 5.9 5.3 5 4 5s-2.4.9-2.8 2H.5c-.3 0-.5.2-.5.5s.2.5.5.5H1c0 1.7 1.3 3 3 3c1.5 0 2.7-1.1 3-2.5c.3 0 .5-.2.5-.5s.2-.5.5-.5s.5.2.5.5s.2.5.5.5c.2 1.4 1.5 2.5 3 2.5c1.7 0 3-1.3 3-3h.5c.3 0 .5-.2.5-.5s-.2-.5-.5-.5M4 10c-1.1 0-2-.9-2-2s.9-2 2-2s2 .9 2 2s-.9 2-2 2m8 0c-1.1 0-2-.9-2-2s.9-2 2-2s2 .9 2 2s-.9 2-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Globe(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8m5.2 5.3c.4 0 .7.3 1.1.3c-.3.4-1.6.4-2-.1c.3-.1.5-.2.9-.2M1 8c0-.4 0-.8.1-1.3c.1 0 .2.1.3.1c0 0 .1.1.1.2c0 .3.3.5.5.5c.8.1 1.1.8 1.8 1c.2.1.1.3 0 .5c-.6.8-.1 1.4.4 1.9c.5.4.5.8.6 1.4c0 .7.1 1.5.4 2.2C2.7 13.3 1 10.9 1 8m7 7c-.7 0-1.5-.1-2.1-.3c-.1-.2-.1-.4 0-.6c.4-.8.8-1.5 1.3-2.2c.2-.2.4-.4.4-.7c0-.2.1-.5.2-.7c.3-.5.2-.8-.2-.9c-.8-.2-1.2-.9-1.8-1.2s-1.2-.5-1.7-.2c-.2.1-.5.2-.5-.1c0-.4-.5-.7-.4-1.1c-.1 0-.2 0-.3.1s-.2.2-.4.1c-.2-.2-.1-.4-.1-.6c.1-.2.2-.3.4-.4c.4-.1.8-.1 1 .4c.3-.9.9-1.4 1.5-1.8c0 0 .8-.7.9-.7s.2.2.4.3c.2 0 .3 0 .3-.2c.1-.5-.2-1.1-.6-1.2c0-.1.1-.1.1-.1c.3-.1.7-.3.6-.6c0-.4-.4-.6-.8-.6c-.2 0-.4 0-.6.1c-.4.2-.9.4-1.5.4C5.2 1.4 6.6 1 8 1h.8c-.6.1-1.2.3-1.6.5c.6.1.7.4.5.9c-.1.2 0 .4.2.5s.4.1.5-.1c.2-.3.6-.4.9-.5c.4-.1.7-.3 1-.7c0-.1.1-.1.2-.2c.6.2 1.2.6 1.8 1c-.1 0-.1.1-.2.1c-.2.2-.5.3-.2.7c.1.2 0 .3-.1.4c-.2.1-.3 0-.4-.1s-.1-.3-.4-.3c-.1.2-.4.3-.4.6c.5 0 .4.4.5.7c-.6.1-.8.4-.5.9c.1.2-.1.3-.2.4c-.4.6-.8 1-.8 1.7s.5 1.4 1.3 1.3c.9-.1.9-.1 1.2.7c0 .1.1.2.1.3c.1.2.2.4.1.6c-.3.8.1 1.4.4 2c.1.2.2.3.3.4c-1.3 1.4-3 2.2-5 2.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlobeWire(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0m6.8 9.5c0 .5-.7.66-2 1a11.68 11.68 0 0 0 .229-1.98l2.001-.02c0 .36-.08.5-.16 1zm-13.6 0c-.1-.5-.15-.64-.2-1h2c.024.723.106 1.411.244 2.079C1.9 10.16 1.2 10 1.2 9.5m0-3c0-.5.7-.66 2-1A11.835 11.835 0 0 0 3 7.489L1 7.5c0-.36.08-.5.16-1zM8.5 5c1.13.013 2.226.107 3.298.277c.047.643.165 1.41.201 2.199L8.5 7.501v-2.5zm0-1V1.06c1.17.27 2.2 1.47 2.84 3.15A24.21 24.21 0 0 0 8.523 4zm-1-2.94V4a25.617 25.617 0 0 0-2.968.214C5.3 2.53 6.33 1.33 7.5 1.06M7.5 5v2.5H4c.031-.806.142-1.571.326-2.307c.932-.08 2.035-.177 3.158-.193zM4 8.5h3.5V11a22.663 22.663 0 0 1-3.298-.277c-.047-.643-.165-1.41-.201-2.199zM7.5 12v2.94c-1.17-.27-2.2-1.47-2.84-3.15a24.21 24.21 0 0 0 2.817.21zm1 2.94V12a25.617 25.617 0 0 0 2.968-.214C10.7 13.47 9.67 14.67 8.5 14.94m0-3.94V8.5H12a11.247 11.247 0 0 1-.326 2.307c-.932.08-2.035.177-3.158.193zM15 7.5h-2a12.229 12.229 0 0 0-.244-2.079c1.354.399 2.014.559 2.014 1.079c.13.5.18.64.23 1m-.7-2.59a9.585 9.585 0 0 0-1.726-.5c-.361-1.019-.809-1.898-1.389-2.672c1.355.726 2.413 1.811 3.067 3.131zM4.84 1.76a8.24 8.24 0 0 0-1.305 2.581c-.699.189-1.299.365-1.874.593c.751-1.39 1.823-2.475 3.139-3.156zm-3.11 9.33c.506.204 1.106.38 1.726.5c.361 1.019.809 1.898 1.389 2.672c-1.367-.722-2.436-1.807-3.097-3.131zm9.44 3.15a8.25 8.25 0 0 0 1.295-2.581c.699-.189 1.299-.365 1.874-.593c-.751 1.39-1.823 2.475-3.139 3.156z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Golf(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 2a2 2 0 1 1-3.999.001A2 2 0 0 1 7 2"/><path fill="currentColor" d="M9.8 1.8c-.2-.5-1.7-.1-2 .5c-.2.3-.2 1.2-1.2 1.9c-.8.5-1.6.5-1.6.5c-.3.6-.1 1.1.2 1.6c.5.9.6 1.8.7 2.8c.1 1.3-.5 2.4-2.3 3.2c-.8.3-1.3.9-1 1.9c0 0 2-.3 3.1-1.2c1.5-1.2 1.8-2.3 1.8-2.3s.1.7 0 1.9c-.1 1-.2 1.5-.4 2.2S7.4 16 8 16s1-.4 1-1l.3-1.9c.3-2.1 0-4.3-.8-6.3c0-.1-.1-.1-.1-.2c-.6-1.6.2-2.6.6-3c.3-.4 1.2-1.2.8-1.8M12 0v10h1V4l3-2zm4 10a1 1 0 1 1-2 0a1 1 0 0 1 2 0M1 8.4l3.7-3.7l-.7-.3L.2 8s-.4.7.1 1.7s1.6.3 1.6.3c.4-.2.2-.4 0-.6s-.9-1-.9-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GooglePlus(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3.9h-2.8V1.3h-.6v2.6H9.9v.8h2.7v2.6h.6V4.7H16zM6.9 9c-.4-.2-1.1-.9-1.1-1.3s.1-.7.8-1.2c.7-.5 1.2-1.2 1.2-2.1c0-1.1-.5-2.1-1.3-2.4h1.3l.9-.7H4.5C2.6 1.3.9 2.7.9 4.4s1.3 3 3.2 3h.4c-.2.2-.2.4-.2.7c0 .5.3.8.6 1.2h-.7c-2.3 0-4.1 1.5-4.1 3s2 2.5 4.3 2.5c2.6 0 4.1-1.5 4.1-3c-.1-1.3-.5-2-1.6-2.8M4.7 6.9c-1.1 0-2.1-1.2-2.3-2.6S2.9 1.8 4 1.8S6.1 3 6.3 4.4S5.8 7 4.7 6.9m-.4 7.2c-1.6 0-2.8-1-2.8-2.2s1.4-2.2 3-2.2c.4 0 .7.1 1 .2c.9.6 1.5.9 1.7 1.6c0 .1.1.3.1.4c0 1.2-.8 2.2-3 2.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GooglePlusSquare(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 3.4c-.8 0-1.3.8-1.2 1.8c.1 1.1.9 1.9 1.7 2c.8 0 1.3-.8 1.2-1.9c-.1-1-.9-1.9-1.7-1.9m.4 5.9c-1.2 0-2.3.7-2.3 1.6s.9 1.7 2.1 1.7c1.7 0 2.3-.7 2.3-1.6v-.3c-.1-.5-.6-.8-1.3-1.2c-.2-.2-.5-.2-.8-.2"/><path fill="currentColor" d="M0 0v16h16V0zm7.9 5.3c0 .7-.4 1.2-.9 1.6s-.6.6-.6.9c0 .3.5.8.8 1c.8.6 1.1 1.1 1.1 2c0 1.1-1.1 2.3-3.1 2.3c-1.7 0-3.2-.7-3.2-1.8C2 10.1 3.3 9 5.1 9h.5c-.2-.3-.4-.6-.4-.9c0-.2.1-.4.2-.6h-.3c-1.4 0-2.4-1-2.4-2.3C2.7 4 4 2.9 5.4 2.9h3.1l-.7.6h-1c.7.2 1.1 1 1.1 1.8m6.1.2h-2.1v2h-.5v-2h-2V5h2V3h.5v2H14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Grab(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.6 4H12c0-.2-.2-.6-.4-.8s-.5-.4-1.1-.4c-.2 0-.4 0-.6.1c-.1-.2-.2-.3-.3-.5c-.2-.2-.5-.4-1.1-.4c-.8 0-1.2.5-1.4 1c-.1 0-.3-.1-.5-.1c-.5 0-.8.2-1.1.4C5 3.9 5 4.7 5 4.8v.4c-.6 0-1.1.2-1.4.5C3 6.4 3 7.3 3 8.5v.7c0 1.4.7 2.1 1.4 2.8l.3.4C6 13.6 7.2 14 9.8 14c2.9 0 4.2-1.6 4.2-5.1V6.4c0-.7-.2-2.1-1.4-2.4m-2.1-.2c.4 0 .5.4.5.6v.8c0 .3.2.5.4.5c.3 0 .5-.1.5-.4c0 0 0-.4.4-.3c.6.2.7 1.1.7 1.3v2.6c0 3.4-1.3 4.1-3.2 4.1c-2.4 0-3.3-.3-4.3-1.3l-.4-.4C4.4 10.6 4 10.2 4 9.2v-.6c0-1 0-1.8.3-2.1c.1-.2.4-.3.7-.3V7l-.3 1.2c0 .1 0 .1.1.1c.1.1.2 0 .2 0l1-1.2V5c0-.1 0-.6.2-.8c.1-.1.2-.2.4-.2c.3 0 .4.2.4.4v.4c0 .2.2.5.5.5S8 5 8 4.8V3.5c0-.1 0-.5.5-.5c.3 0 .5.2.5.5v1.2c0 .3.2.6.5.6s.5-.3.5-.5v-.5c0-.3.2-.5.5-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Grid(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v16h16V0zm5 15H1v-4h4zm0-5H1V6h4zm0-5H1V1h4zm5 10H6v-4h4zm0-5H6V6h4zm0-5H6V1h4zm5 10h-4v-4h4zm0-5h-4V6h4zm0-5h-4V1h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GridBevel(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 2V1H1v13h1v1h13V2zM5 13H2v-3h3zm0-4H2V6h3zm0-4H2V2h3zm4 8H6v-3h3zm0-4H6V6h3zm0-4H6V2h3zm4 8h-3v-3h3zm0-4h-3V6h3zm0-4h-3V2h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GridBig(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h7v7H0zm9 0h7v7H9zM0 9h7v7H0zm9 0h7v7H9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GridBigO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 7h7V0H0zm1-6h5v5H1zm8-1v7h7V0zm6 6h-5V1h5zM0 16h7V9H0zm1-6h5v5H1zm8 6h7V9H9zm1-6h5v5h-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GridH(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v16h16V0zm5 15H1V1h4zm5 0H6V1h4zm5 0h-4V1h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GridSmall(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h4v4H0zm0 6h4v4H0zm0 6h4v4H0zM6 0h4v4H6zm0 6h4v4H6zm0 6h4v4H6zm6-12h4v4h-4zm0 6h4v4h-4zm0 6h4v4h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GridSmallO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 4h4V0H0zm1-3h2v2H1zm-1 9h4V6H0zm1-3h2v2H1zm-1 9h4v-4H0zm1-3h2v2H1zm5-9h4V0H6zm1-3h2v2H7zm-1 9h4V6H6zm1-3h2v2H7zm-1 9h4v-4H6zm1-3h2v2H7zm5-13v4h4V0zm3 3h-2V1h2zm-3 7h4V6h-4zm1-3h2v2h-2zm-1 9h4v-4h-4zm1-3h2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GridV(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 0H0v16h16zM1 5V1h14v4zm0 5V6h14v4zm0 5v-4h14v4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GroupIcon(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 16v-5.3c-.6-.3-1-1-1-1.7V5c0-.7.4-1.3 1-1.7V3c0-1.1-.9-2-2-2s-2 .9-2 2s.9 2 2 2H1c-.5 0-1 .5-1 1v4c0 .5.5 1 1 1v5zM15 5h-2c1.1 0 2-.9 2-2s-.9-2-2-2s-2 .9-2 2v.3c.6.4 1 1 1 1.7v4c0 .7-.4 1.4-1 1.7V16h4v-5c.5 0 1-.5 1-1V6c0-.5-.5-1-1-1m-5-3a2 2 0 1 1-3.999.001A2 2 0 0 1 10 2"/><path fill="currentColor" d="M10 4H6c-.5 0-1 .5-1 1v4c0 .5.5 1 1 1v6h4v-6c.5 0 1-.5 1-1V5c0-.5-.5-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hammer(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m6 2l7 7l3-3l-4.48-4.48S8.55 2.55 7 1zm2.8 3.79L.27 14.31a.998.998 0 0 0 0 1.361a.998.998 0 0 0 1.371.049l8.569-8.519z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hand(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 2.4c-.4-.4-1-.5-1.5-.3c0-.3-.1-.6-.4-.9c-.2-.2-.6-.4-1.1-.4c-.3 0-.5.1-.7.1c0-.2-.1-.3-.2-.5c-.5-.6-1.5-.6-2 0c-.2.2-.4.4-.4.6C7 1 6.8.9 6.6.9c-.5 0-.8.2-1.1.5C5 1.9 5 2.7 5 2.7v3.8c-.3-.3-.8-.8-1.5-.8c-.2 0-.5.1-.7.2c-.4.2-.6.5-.7.9c-.3 1 .6 2.4.6 2.5c.1.1 1.2 2.7 2.2 3.8C5.9 14.3 7 15 9.8 15c2.9 0 4.2-1.6 4.2-5.1V4.4c0-.1.1-1.3-.5-2M8 2c0-.3-.1-1 .5-1c.5 0 .5.5.5 1v4c0 .3.2.5.5.5s.5-.2.5-.5V2.2s0-.4.5-.4c.6 0 .5.9.5.9V6c0 .3.2.5.5.5s.5-.2.5-.5V3.6c0-.1 0-.6.5-.6s.5 1 .5 1v5.9c0 3.4-1.3 4.1-3.2 4.1c-2.4 0-3.3-.5-4.1-1.6c-.9-1-2.1-3.6-2.1-3.7c-.3-.3-.7-1.2-.6-1.6c0-.1.1-.2.2-.3c.1 0 .2-.1.2-.1c.4 0 .8.5.9.7l.6.9c.1.2.4.3.6.2c.4 0 .5-.2.5-.4V2.9c0-.4 0-1 .5-1c.4 0 .5.3.5.8V6c0 .3.2.5.5.5S8 6.3 8 6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandleCorner(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.7 16L16 6.7V5.3L5.3 16zm3 0L16 9.7V8.3L8.3 16zm3 0l3.3-3.3v-1.4L11.3 16zm3 0l.3-.3v-1.4L14.3 16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandsUp(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 2a2 2 0 1 1-3.999.001A2 2 0 0 1 10 2"/><path fill="currentColor" d="M6 16h1.5v-5h1v5H10V7l-.001-.052c0-.521.194-.997.513-1.36L13.79 2.27a.73.73 0 1 0-.998-1.003L10.43 3.65c-.212.216-.508.35-.835.35H6.404c-.327 0-.622-.134-.834-.35L3.25 1.26a.73.73 0 1 0-1.003.998L5.49 5.59c.317.361.511.836.511 1.358L6 7.003z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Handshake(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 3a5.393 5.393 0 0 1-1.902 1.178c-.748.132-2.818-.828-3.838.152c-.17.17-.38.34-.6.51c-.48-.21-1.22-.53-1.76-.84S3 3 3 3L0 6.5s.74 1 1.2 1.66c.3.44.67 1.11.91 1.56l-.34.4a.876.876 0 0 0 .15 1a.833.833 0 0 0 1.002-.002a.62.62 0 0 0 .077.881a.994.994 0 0 0 1.006-.002a.806.806 0 0 0-.003 1.005a1.012 1.012 0 0 0 .892-.114a.822.822 0 0 0 .187.912a1.093 1.093 0 0 0 1.054-.092l.516-.467c.472.47 1.123.761 1.842.761l.061-.001a1.311 1.311 0 0 0 1.094-.791c.146.056.312.094.488.094c.236 0 .455-.068.64-.185c.585-.387.445-.687.445-.687a1.07 1.07 0 0 0 1.229-.279a.996.996 0 0 0 .138-1.215a.036.036 0 0 0 .021.005c.421 0 .787-.232.978-.574a1.564 1.564 0 0 0-.191-1.48l.003.005c.82-.16.79-.57 1.19-1.17a4.725 4.725 0 0 1 1.387-1.208zm-.05 7.06c-.44.44-.78.25-1.53-.32S9.18 8.1 9.18 8.1c.061.305.202.57.401.781c.319.359 1.269 1.179 1.719 1.599c.28.26 1 .78.58 1.18s-.75 0-1.44-.56s-2.23-1.94-2.23-1.94a.937.937 0 0 0 .27.72c.17.2 1.12 1.12 1.52 1.54s.75.67.41 1s-1.03-.19-1.41-.58c-.59-.57-1.76-1.63-1.76-1.63l-.001.053c0 .284.098.544.263.75c.288.378.848.868 1.188 1.248s.54.7 0 1s-1.34-.44-1.69-.8v-.002a.411.411 0 0 0-.1-.269a.896.896 0 0 0-.906-.188A.609.609 0 0 0 6 11.1a.754.754 0 0 0-.912.001a.61.61 0 0 0-.085-.95a1 1 0 0 0-1.174.08a.66.66 0 0 0-.068-.911a.996.996 0 0 0-1.186-.128L1.91 8.069c-.46-.73-1-1.49-1-1.49l2.28-2.77s.81.5 1.48.88c.33.19.9.44 1.33.64c-.68.51-1.25 1-1.08 1.34a1.834 1.834 0 0 0 2.087.036a2.41 2.41 0 0 1 1.343-.403c.347 0 .677.072.976.203c.554.374 1.574 1.294 2.504 1.874c1.17.85 1.4 1.4 1.12 1.68z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Harddrive(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 1H3L.3 9h15.4zM0 10v5h16v-5zm3 3H2v-1h1zm4 0H4v-1h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HarddriveO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 12h1v1H2zm2 0h3v1H4z"/><path fill="currentColor" d="M13 1H3l-3 9v5h16v-5zM3.7 2h8.6l2.7 8H1.1zM1 14v-3h14v3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hash(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 6V4h-2.6l.6-2.8l-2-.4l-.7 3.2h-3L8 1.2L6 .8L5.3 4H2v2h2.9L4 10H1v2h2.6L3 14.8l2 .4l.7-3.2h3L8 14.8l2 .4l.7-3.2H14v-2h-2.9l.9-4zm-6 4H6l1-4h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Header(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 0v7H5V0H2v16h3V9h6v7h3V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Headphones(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 8.3V6c0-3.3-2.7-6-6-6S2 2.7 2 6v2.3c-1.2.5-2 1.7-2 3.1v1.2c0 1.8 1.3 3.2 3 3.4h2V8H4V6c0-2.2 1.8-4 4-4s4 1.8 4 4v2h-1v8h2c1.7-.2 3-1.7 3-3.4v-1.2c0-1.4-.8-2.6-2-3.1M4 15H3V9h1zm9 0h-1V9h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Headset(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.82 8a3.016 3.016 0 0 0-1.799-1.813L13 4.5C13 2 10.53 0 7.5 0S2 2 2 4.5v1.68A3.006 3.006 0 0 0 0 9v1a3 3 0 0 0 3 3h1V6H3V4.5C3 2.57 5 1 7.5 1S12 2.57 12 4.5V6h-1v7h1a3 3 0 0 0 3-3v1.73A3.27 3.27 0 0 1 11.73 15H10a1 1 0 0 0-1-1H8a1 1 0 0 0 0 2h3.73A4.27 4.27 0 0 0 16 11.73V8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HealthCard(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 3v10H1V3zm1-1H0v12h16z"/><path fill="currentColor" d="M9 5h5v1H9zm0 2h5v1H9zm0 2h2v1H9zM6.5 5c-.6 0-1.1.6-1.5 1c-.4-.4-.9-1-1.5-1c-1.5 0-2.1 1.9-1 2.9L5 10l2.5-2.1C8.6 6.9 8 5 6.5 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heart(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2S9 2 8 5C7 2 4 2 4 2C1.8 2 0 3.8 0 6c0 4.1 8 9 8 9s8-5 8-9c0-2.2-1.8-4-4-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.7 2C10.8 2 9 2.5 8 4.1C7 2.5 5.2 2 4.2 2C1.9 2 0 3.9 0 6.2c0 4 7.4 8.5 7.7 8.7l.3.2l.3-.2c.3-.2 7.7-4.8 7.7-8.7C16 3.9 14.1 2 11.7 2M8 13.9c-2.2-1.4-7-5-7-7.7C1 4.4 2.5 3 4.2 3c.1 0 2.5.1 3.3 2.4L8 6.8l.5-1.4C9.3 3.1 11.7 3 11.8 3C13.5 3 15 4.4 15 6.2c0 2.7-4.8 6.3-7 7.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Home(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1.4L6 2.7V1H4v3L0 6.6l.6.8L8 2.6l7.4 4.8l.6-.8z"/><path fill="currentColor" d="M8 4L2 8v7h5v-3h2v3h5V8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 6.6L8 1.4L6 2.7V1H4v3L0 6.6l1.9 2.7l.1-.1V15h5v-4h2v4h5V9.2l.1.1zm-14.6.3L8 2.6l6.6 4.3l-.7 1L8 4L2.1 7.9zM13 14h-3v-4H6v4H3V8.6l5-3.3l5 3.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hospital(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 4V0H8v4H0v12h6v-3h4v3h6V4zM4 11H2V9h2zm0-3H2V6h2zm3 3H5V9h2zm0-3H5V6h2zm3-5V2h1V1h1v1h1v1h-1v1h-1V3zm1 8H9V9h2zm0-3H9V6h2zm3 3h-2V9h2zm0-3h-2V6h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hourglass(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.16 4.6A4.054 4.054 0 0 1 8 7.994c0-1.415.726-2.66 1.825-3.384c.23-.199.426-.395.609-.602L5.56 4.001c.19.214.386.41.593.594z"/><path fill="currentColor" d="M11.18 6.06A4.399 4.399 0 0 0 13 2.5V2h1V0H2v2h1v.5a4.391 4.391 0 0 0 1.808 3.551A2.564 2.564 0 0 1 6 7.99a2.755 2.755 0 0 1-1.209 2.003a4.441 4.441 0 0 0-1.79 3.503v.503h-1v2h12v-2h-1v-.5a4.435 4.435 0 0 0-1.769-3.492a2.762 2.762 0 0 1-1.23-1.996a2.551 2.551 0 0 1 1.169-1.946zM9 8a3.693 3.693 0 0 0 1.519 2.763A3.477 3.477 0 0 1 12 13.495V14h-1.77c-.7-.87-1.71-2-2.23-2s-1.53 1.13-2.23 2H4v-.5a3.472 3.472 0 0 1 1.459-2.723a3.698 3.698 0 0 0 1.54-2.766a3.482 3.482 0 0 0-1.498-2.683a3.438 3.438 0 0 1-1.502-2.827v-.5h8v.5a3.426 3.426 0 0 1-1.479 2.813a3.487 3.487 0 0 0-1.521 2.678z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HourglassEmpty(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.18 6.06A4.399 4.399 0 0 0 13 2.5V2h1V0H2v2h1v.5a4.391 4.391 0 0 0 1.808 3.551A2.564 2.564 0 0 1 6 7.99a2.755 2.755 0 0 1-1.209 2.003a4.441 4.441 0 0 0-1.79 3.503v.503h-1v2h12v-2h-1v-.5a4.435 4.435 0 0 0-1.769-3.492a2.762 2.762 0 0 1-1.23-1.996a2.551 2.551 0 0 1 1.169-1.946zM9 8a3.693 3.693 0 0 0 1.519 2.763A3.477 3.477 0 0 1 12 13.495V14H4v-.5a3.472 3.472 0 0 1 1.459-2.723a3.698 3.698 0 0 0 1.54-2.766a3.482 3.482 0 0 0-1.498-2.683a3.438 3.438 0 0 1-1.502-2.827v-.5h8v.5a3.426 3.426 0 0 1-1.479 2.813a3.487 3.487 0 0 0-1.521 2.678z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HourglassEnd(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.18 6.06A4.399 4.399 0 0 0 13 2.5V2h1V0H2v2h1v.5a4.391 4.391 0 0 0 1.808 3.551A2.564 2.564 0 0 1 6 7.99a2.755 2.755 0 0 1-1.209 2.003a4.441 4.441 0 0 0-1.79 3.503v.503h-1v2h12v-2h-1v-.5a4.435 4.435 0 0 0-1.769-3.492a2.762 2.762 0 0 1-1.23-1.996a2.551 2.551 0 0 1 1.169-1.946zM9 8a3.693 3.693 0 0 0 1.519 2.763A3.477 3.477 0 0 1 12 13.495V14h-1s-1.62-3.5-3-3.5S5 14 5 14H4v-.5a3.472 3.472 0 0 1 1.459-2.723a3.698 3.698 0 0 0 1.54-2.766a3.482 3.482 0 0 0-1.498-2.683a3.438 3.438 0 0 1-1.502-2.827v-.5h8v.5a3.426 3.426 0 0 1-1.479 2.813a3.487 3.487 0 0 0-1.521 2.678z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HourglassStart(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.16 4.6A4.054 4.054 0 0 1 8 7.994c0-1.415.726-2.66 1.825-3.384a2.857 2.857 0 0 0 1.17-1.589L5 3.001c.191.67.603 1.225 1.15 1.594z"/><path fill="currentColor" d="M11.18 6.06A4.399 4.399 0 0 0 13 2.5V2h1V0H2v2h1v.5a4.391 4.391 0 0 0 1.808 3.551A2.564 2.564 0 0 1 6 7.99a2.755 2.755 0 0 1-1.209 2.003a4.441 4.441 0 0 0-1.79 3.503v.503h-1v2h12v-2h-1v-.5a4.435 4.435 0 0 0-1.769-3.492a2.762 2.762 0 0 1-1.23-1.996a2.551 2.551 0 0 1 1.169-1.946zM9 8a3.693 3.693 0 0 0 1.519 2.763A3.477 3.477 0 0 1 12 13.495V14H4v-.5a3.472 3.472 0 0 1 1.459-2.723a3.698 3.698 0 0 0 1.54-2.766a3.482 3.482 0 0 0-1.498-2.683a3.438 3.438 0 0 1-1.502-2.827v-.5h8v.5a3.426 3.426 0 0 1-1.479 2.813a3.487 3.487 0 0 0-1.521 2.678z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Inbox(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 6V0H6v6H4l4 5l4-5z"/><path fill="currentColor" d="M13 1h-2v1h1.3l2.6 8H11v2H5v-2H1.1l2.6-8H5V1H3l-3 9v5h16v-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Indent(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h16v3H0zm6 4h10v3H6zm0 4h10v3H6zm-6 4h16v3H0zm0-7.5v6l4-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Info(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 5h4v11H6zm4-3a2 2 0 1 1-3.999.001A2 2 0 0 1 10 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoCircle(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8m1 13H7V6h2zm0-8H7V3h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoCircleO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1c3.9 0 7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7m0-1C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8"/><path fill="currentColor" d="M7 6h2v7H7zm0-3h2v2H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Input(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 5c0-.6-.4-1-1-1H1c-.6 0-1 .4-1 1v6c0 .6.4 1 1 1h14c.6 0 1-.4 1-1zm-1 6H1V5h14z"/><path fill="currentColor" d="M2 6h1v4H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Insert(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 16V5l-1 1v9H1V3h9l1-1H0v14z"/><path fill="currentColor" d="M16 1.4L14.6 0L7.8 6.8L6 5v5h5L9.2 8.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Institution(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0L0 3v2h16V3zM0 14h16v2H0zm16-7V6H0v1h1v5H0v1h16v-1h-1V7zM4 12H3V7h1zm3 0H6V7h1zm3 0H9V7h1zm3 0h-1V7h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Invoice(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.4 10.2c-.6.1-1.4-.3-1.7-.4l-.5.9s.9.4 1.7.5v.8h1v-.9c.9-.3 1.4-1.1 1.5-1.8c0-.8-.6-1.4-1.9-1.9c-.4-.2-1.1-.5-1.1-.9c0-.5.4-.8 1-.8c.7 0 1.4.3 1.4.3l.4-.9s-.5-.2-1.2-.4V4H4v.7c-.9.2-1.5.8-1.6 1.7c0 1.2 1.3 1.7 1.8 1.9c.6.2 1.3.6 1.3.9c0 .4-.4.9-1.1 1"/><path fill="currentColor" d="M0 2v12h16V2zm15 11H1V3h14z"/><path fill="currentColor" d="M8 5h6v1H8zm0 2h6v1H8zm0 2h3v1H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Italic(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0h3L8 16H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Key(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8.1 7l-.6-.3L15 0h-2L6 6.1C5.7 6 5.4 6 5 6c-2.8 0-5 2.2-5 5s2.3 5 5 5s5-2.2 5-5c0-.6-.1-1.2-.3-1.7L11 8V6h2V4h2l1-1V0zM4 13.2c-.7 0-1.2-.6-1.2-1.2s.6-1.2 1.2-1.2s1.2.6 1.2 1.2s-.5 1.2-1.2 1.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 0L6 6.1C5.7 6 5.4 6 5 6c-2.8 0-5 2.2-5 5s2.3 5 5 5s5-2.2 5-5c0-.3 0-.6-.1-.9L11 9V7h2V5h2l1-1V0zm-1 6h-1.7L12 4.6zm3-2.4l-.4.4h-1.9L15 2zm-7.7 4L8 8l2-1.7v2.3l-.8.8l-.3.4l.1.5c0 .2.1.5.1.7c0 2.2-1.8 4-4 4s-4-1.8-4-4s1.8-4 4-4c.3 0 .5 0 .8.1l.5.1l.4-.3L13.4 1H15z"/><path fill="currentColor" d="M6 11.5a1.5 1.5 0 1 1-3.001-.001A1.5 1.5 0 0 1 6 11.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Keyboard(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 4v9h16V4zm10 2h1v1h-1zM8 6h1v1H8zm2 2v1H9V8zM6 6h1v1H6zm2 2v1H7V8zM4 6h1v1H4zm2 2v1H5V8zM2 6h1v1H2zm1 5H2v-1h1zm0-3h1v1H3zm9 3H4v-1h8zm0-2h-1V8h1zm2 2h-1v-1h1zm0-2h-1V7h-1V6h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyboardO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 5v7H1V5zm1-1H0v9h16z"/><path fill="currentColor" d="M4 10h8v1H4zm-2 0h1v1H2zm11 0h1v1h-1zm-2-2h1v1h-1zM9 8h1v1H9zM7 8h1v1H7zM5 8h1v1H5zM3 8h1v1H3zm7-2h1v1h-1zm2 0v1h1v2h1V6zM8 6h1v1H8zM6 6h1v1H6zM4 6h1v1H4zM2 6h1v1H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Laptop(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 11V2H2v9H0v2h16v-2zm-4 1H6v-1h4zm3-2H3V3h10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Layout(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v16h16V0zm1 3h4v12H1zm14 12H6V3h9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LevelDown(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 1h6v11h2l-3 3l-3-3h2V3H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LevelDownBold(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9 16l4-7h-3V0H3l2 3h2v6H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LevelLeft(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 12V6H4V4L1 7l3 3V8h9v6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LevelLeftBold(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 7l7-4v3h9v7l-3-2V9H7v3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LevelRight(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 12V6h11V4l3 3l-3 3V8H3v6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LevelRightBold(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 7L9 3v3H0v7l3-2V9h6v3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LevelUp(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 15H5V4H3l3-3l3 3H7v9h6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LevelUpBold(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9 0l4 7h-3v9H3l2-3h2V7H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lifebuoy(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8M4 8c0-2.2 1.8-4 4-4s4 1.8 4 4s-1.8 4-4 4s-4-1.8-4-4m8.6 1.8c.3-.5.4-1.2.4-1.8s-.1-1.3-.4-1.8l1.5-1.5c.6 1 .9 2.1.9 3.3s-.3 2.3-.8 3.3zm-1.3-8L9.8 3.4C9.3 3.1 8.6 3 8 3s-1.3.1-1.8.4L4.7 1.8C5.7 1.3 6.8 1 8 1s2.3.3 3.3.8M1.8 4.7l1.5 1.5C3.1 6.7 3 7.4 3 8s.1 1.3.4 1.8l-1.5 1.5C1.3 10.3 1 9.2 1 8s.3-2.3.8-3.3m2.9 9.5l1.5-1.5c.5.2 1.2.3 1.8.3s1.3-.1 1.8-.4l1.5 1.5c-1 .6-2.1.9-3.3.9s-2.3-.3-3.3-.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lightbulb(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a5 5 0 0 0-5 5a4.8 4.8 0 0 0 2.182 3.989L6 11a.5.5 0 0 0 0 1a.5.5 0 0 0 0 1a.5.5 0 0 0 0 1a.5.5 0 0 0 0 1h.41c.342.55.915.929 1.581.999a2.122 2.122 0 0 0 1.594-.99L10 15a.5.5 0 0 0 0-1a.5.5 0 0 0 0-1a.5.5 0 0 0 0-1a.5.5 0 0 0 0-1l.8-2A4.805 4.805 0 0 0 13 5.002A5.001 5.001 0 0 0 8 0m2.25 8.21l-.25.17l-.11.29L9 10.81a.292.292 0 0 1-.27.19H7.22a.29.29 0 0 1-.219-.188L6.13 8.67L6 8.38l-.25-.18A3.914 3.914 0 0 1 4 5.003A4 4 0 1 1 12 5a3.905 3.905 0 0 1-1.736 3.201z"/><path fill="currentColor" d="M10.29 3a3.153 3.153 0 0 0-2.289-1L8 3a2.133 2.133 0 0 1 1.5.62c.278.386.459.858.499 1.37L11 4.999a3.785 3.785 0 0 0-.718-2.011z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LineBarChart(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 11h3v5H5zm-4 3h3v2H1zm12-2h3v4h-3zM9 9h3v7H9zm7-8.93l-5.68 4.97l-5.47-1.7L0 7.1V9l5.15-4l5.53 1.72L16 2.06z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LineChart(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 15V0H0v16h16v-1z"/><path fill="currentColor" d="M9 8L6 5L2 9v2l4-4l3 3l7-7V1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LineH(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 7h16v1H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LineV(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0h1v16H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lines(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 1h16v2H0zm0 4h16v2H0zm0 4h16v2H0zm0 4h16v2H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinesList(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 1h3v2H0zm0 4h3v2H0zm0 4h3v2H0zm0 4h3v2H0zM4 1h12v2H4zm0 4h12v2H4zm0 4h12v2H4zm0 4h12v2H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Link(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.9 1.1c-1.4-1.4-3.7-1.4-5.1 0L5.4 5.4C4 6.9 4 9.1 5.4 10.6c.1.1.3.2.4.3l1.5-1.5c-.1-.1-.3-.2-.4-.3c-.6-.6-.6-1.6 0-2.2l4.4-4.4c.6-.6 1.6-.6 2.2 0s.6 1.6 0 2.2L12.2 6c.4.8.5 1.7.4 2.5l2.3-2.3c1.5-1.4 1.5-3.7 0-5.1"/><path fill="currentColor" d="M10.2 5.1L8.7 6.6s.3.2.4.3c.6.6.6 1.6 0 2.2l-4.4 4.4c-.6.6-1.6.6-2.2 0s-.6-1.6 0-2.2L3.8 10c-.4-.8-.1-1.3-.4-2.5L1.1 9.8c-1.4 1.4-1.4 3.7 0 5.1s3.7 1.4 5.1 0l4.4-4.4c1.4-1.4 1.4-3.7 0-5.1c-.2-.1-.4-.3-.4-.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func List(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h4v3H0zm0 4h4v3H0zm0 8h4v3H0zm0-4h4v3H0zm5-8h11v3H5zm0 4h11v3H5zm0 8h11v3H5zm0-4h11v3H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListOl(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 0h12v4H4zm0 6h12v4H4zm0 6h12v4H4zM1 0L.1.5l.2.7l.7-.3V4h1V0zm1.2 13.9c.3-.2.5-.5.5-.8c0-.5-.4-1-1.3-1c-.5 0-1 .1-1.2.3H.1l.2.8l.1-.1c.1-.1.4-.2.7-.2s.4.1.4.3c0 .4-.5.4-.6.4H.5v.7h.4c.3 0 .6.1.6.4c0 .2-.2.4-.6.4s-.7-.2-.8-.2l-.1-.1v.9h.1c.2.2.6.3 1.1.3c1 0 1.6-.5 1.6-1.2c0-.4-.2-.8-.6-.9M.1 6.4l.3 1s.7-.6 1.2-.3C2.7 7.9 0 9.5 0 9.5v.5h3V9H1.8c.6-.5 1.2-1.2 1-1.9C2.3 5.2.1 6.4.1 6.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListSelect(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 0h12v2H1zm0 8h13v2H1zm0 3h11v2H1zm0 3h14v2H1zM0 3v4h16V3zm11 3H1V4h10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListUl(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 1h3v3H0zm0 5h3v3H0zm0 5h3v3H0zM5 1h11v3H5zm0 5h11v3H5zm0 5h11v3H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationArrow(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 9l16-9l-9 16V9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationArrowCircle(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8M7 14V9H2l10-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationArrowCircleO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 8c0-3.9 3.1-7 7-7s7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7M0 8c0 4.4 3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8s-8 3.6-8 8"/><path fill="currentColor" d="m2 9l10-5l-5 10V9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lock(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 8V4.9C12 2.7 10.4 1 8.2 1h-.3C5.8 1 4 2.7 4 4.9V8H3l.1 5S3 16 8 16s5-3 5-3V8zm-3 6H8v-2c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1zm1-6H6V4.9C6 3.8 6.9 3 7.9 3h.3c1 0 1.8.8 1.8 1.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Magic(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 5h3v1H0zm5-5h1v3H5zm1 11H5V8.5l1 1zm5-5H9.5l-1-1H11zM3.131 7.161l.707.707l-2.97 2.97l-.707-.707zm7-7l.707.707l-2.97 2.97l-.707-.707zM.836.199l3.465 3.465l-.707.707L.129.906zM6.1 4.1L4 6.1l9.8 9.9l2.2-2.1zm0 1.4L8.5 8l-.6.6l-2.5-2.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Magnet(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 0h5v4h-5zm0 5v3c0 1.6-1.4 3-3 3S5 9.6 5 8V5H0v3c0 4.4 3.6 8 8 8s8-3.6 8-8V5zM0 0h5v4H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mailbox(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 1H3l-3 9v5h16v-5zm-2 9v2H5v-2H1.1l2.7-8h8.6l2.7 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Male(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 2a2 2 0 1 1-3.999.001A2 2 0 0 1 10 2"/><path fill="currentColor" d="m12.79 7.32l-2.6-2.63A2.311 2.311 0 0 0 8.54 4H7.469c-.648 0-1.235.264-1.659.69l-2.6 2.63a.73.73 0 1 0 .998 1.003L6 6.53V16h1.5v-5h1v5H10V6.53l1.75 1.8a.73.73 0 1 0 1.041-1.009z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarker(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0C5.2 0 3 2.2 3 5s4 11 5 11s5-8.2 5-11s-2.2-5-5-5m0 8C6.3 8 5 6.7 5 5s1.3-3 3-3s3 1.3 3 3s-1.3 3-3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Margin(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h1v1H0zm2 0h1v1H2zM1 1h1v1H1zM0 2h1v1H0zm2 0h1v1H2zM1 3h1v1H1zM0 4h1v1H0zm1 1h1v1H1zM0 6h1v1H0zm1 1h1v1H1zM0 8h1v1H0zm1 1h1v1H1zm-1 1h1v1H0zm1 1h1v1H1zm-1 1h1v1H0zm1 1h1v1H1zm-1 1h1v1H0zm2 0h1v1H2zm-1 1h1v1H1zm2 0h1v1H3zm2 0h1v1H5zM4 0h1v1H4zM3 1h1v1H3zm2 0h1v1H5zM4 14h1v1H4zM6 0h1v1H6zm2 0h1v1H8zM7 1h1v1H7zM6 14h1v1H6zm2 0h1v1H8zm-1 1h1v1H7zm2 0h1v1H9zm2 0h1v1h-1zM10 0h1v1h-1zM9 1h1v1H9zm2 0h1v1h-1zm-1 13h1v1h-1zm2-14h1v1h-1zm2 0h1v1h-1zm-1 1h1v1h-1z"/><path fill="currentColor" d="M13 2h-1v1h-1V2h-1v1H9V2H8v1H7V2H6v1H5V2H4v1H3v1H2v1h1v1H2v1h1v1H2v1h1v1H2v1h1v1H2v1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v-1h1v-1h-1v-1h1V9h-1V8h1V7h-1V6h1V5h-1V4h1V3h-1zm-1 10H4V4h8zm2-10h1v1h-1zm0 2h1v1h-1zm0 2h1v1h-1zm0 2h1v1h-1zm0 2h1v1h-1zm0 2h1v1h-1z"/><path fill="currentColor" d="M13 13h1v1h-1zm-1 1h1v1h-1zm2 0h1v1h-1zm-1 1h1v1h-1zm2 0h1v1h-1zm0-14h1v1h-1zm0 2h1v1h-1zm0 2h1v1h-1zm0 2h1v1h-1zm0 2h1v1h-1zm0 2h1v1h-1zm0 2h1v1h-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MarginBottom(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v14h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1v1h1v-1h1V0zm15 12H1V1h14zM0 15h1v1H0zm1-1h1v1H1zm1 1h1v1H2zm1-1h1v1H3zm1 1h1v1H4zm1-1h1v1H5zm1 1h1v1H6zm1-1h1v1H7zm1 1h1v1H8zm1-1h1v1H9zm1 1h1v1h-1zm1-1h1v1h-1zm1 1h1v1h-1zm1-1h1v1h-1zm1 1h1v1h-1zm1-1h1v1h-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MarginLeft(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 0v1h1v1H2v1h1v1H2v1h1v1H2v1h1v1H2v1h1v1H2v1h1v1H2v1h1v1H2v1h1v1h13V0zm13 15H4V1h11zM0 0h1v1H0zm1 1h1v1H1zM0 2h1v1H0zm1 1h1v1H1zM0 4h1v1H0zm1 1h1v1H1zM0 6h1v1H0zm1 1h1v1H1zM0 8h1v1H0zm1 1h1v1H1zm-1 1h1v1H0zm1 1h1v1H1zm-1 1h1v1H0zm1 1h1v1H1zm-1 1h1v1H0zm1 1h1v1H1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MarginRight(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 2V1h-1V0H0v16h14v-1h-1v-1h1v-1h-1v-1h1v-1h-1v-1h1V9h-1V8h1V7h-1V6h1V5h-1V4h1V3h-1V2zm-2 13H1V1h11zm3 0h1v1h-1zm-1-1h1v1h-1zm1-1h1v1h-1zm-1-1h1v1h-1zm1-1h1v1h-1zm-1-1h1v1h-1zm1-1h1v1h-1zm-1-1h1v1h-1zm1-1h1v1h-1zm-1-1h1v1h-1zm1-1h1v1h-1zm-1-1h1v1h-1zm1-1h1v1h-1zm-1-1h1v1h-1zm1-1h1v1h-1zm-1-1h1v1h-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MarginTop(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 2v1h-1V2h-1v1h-1V2h-1v1h-1V2H9v1H8V2H7v1H6V2H5v1H4V2H3v1H2V2H1v1H0v13h16V2zm0 13H1V4h14zm0-15h1v1h-1zm-1 1h1v1h-1zm-1-1h1v1h-1zm-1 1h1v1h-1zm-1-1h1v1h-1zm-1 1h1v1h-1zM9 0h1v1H9zM8 1h1v1H8zM7 0h1v1H7zM6 1h1v1H6zM5 0h1v1H5zM4 1h1v1H4zM3 0h1v1H3zM2 1h1v1H2zM1 0h1v1H1zM0 1h1v1H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Medal(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 12.2c-.3 0-.5-.1-.8-.2L8 11.5l-1.2.5c-.2.1-.5.2-.8.2c-.2 0-.3 0-.5-.1L5 16l3-2l3 2l-.6-3.9c-.1.1-.3.1-.4.1m2.9-6.3c-.1-.2-.1-.5 0-.7l.6-1.2c.2-.4 0-.9-.5-1.1l-1.3-.5c-.2-.1-.4-.3-.5-.5L10.7.6c-.1-.4-.4-.6-.7-.6c-.1 0-.3 0-.4.1L8.3.7H8c-.1 0-.2 0-.3-.1L6.4.1C6.3 0 6.1 0 6 0c-.3 0-.6.2-.8.5l-.5 1.4c0 .2-.2.4-.4.5l-1.4.5c-.4.1-.6.6-.4 1.1l.6 1.3c.1.2.1.5 0 .7l-.6 1.2c-.2.4 0 .9.5 1.1l1.3.5c.2.1.4.3.5.5l.5 1.3c.1.4.4.6.7.6c.1 0 .2 0 .3-.1l1.3-.6c.1 0 .2-.1.3-.1s.2 0 .3.1l1.3.6c.1.1.2.1.3.1c.3 0 .6-.2.8-.5l.5-1.3c.1-.2.3-.4.5-.5l1.3-.5c.4-.2.7-.7.5-1.1zM8 9.6c-2.2 0-4-1.8-4-4s1.8-4 4-4s4 1.8 4 4s-1.8 4-4 4"/><path fill="currentColor" d="M11 5.6a3 3 0 1 1-6 0a3 3 0 0 1 6 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Megafone(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.5 5.4L15 5V1c0-.6-.4-1-1-1s-1 .4-1 1v.5C11 2.4 8 4 5 4H2.5C1.1 4 0 5.2 0 6.5c0 .9.5 1.7 1.2 2.1l1.1 5.9c0 .3.3.5.7.5h.2l3.6-.7c.4-.1.6-.4.5-.7c-.3-.6-.8-1.5-1.2-1.8c-.2-.1-.5-.9-.7-1.8H6v-.9c2.7.3 6 1.6 7 2.4v.5c0 .6.4 1 1 1s1-.4 1-1V8l.4-.3c.4-.3.6-.7.6-1.1v-.2c0-.4-.2-.7-.5-1M2 5h3v1H2zm3.6 7.6c.1 0 .3.3.5.7l-2.8.7l-1-5h1.9c.2 1.3.6 3.2 1.4 3.6m7.4-2.3c-1.6-.8-4.4-2-7-2.3V5c2.6-.3 5.4-1.4 7-2.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MehO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1c3.9 0 7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7m0-1C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8"/><path fill="currentColor" d="M7 6a1 1 0 1 1-2 0a1 1 0 0 1 2 0m4 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-7 4h8v1H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Menu(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 1h16v3H0zm0 5h16v3H0zm0 5h16v3H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Microphone(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 10c-1.7 0-3-1.3-3-3V3c0-1.6 1.3-3 3-3c1.6 0 3 1.3 3 3v4c0 1.6-1.4 3-3 3"/><path fill="currentColor" d="M12 5v2.5c0 1.9-1.8 3.5-3.8 3.5h-.4C5.8 11 4 9.4 4 7.5V5c-.6 0-1 .4-1 1v1.5c0 2.2 1.8 4.1 4 4.4V14c-3 0-2.5 2-2.5 2h7s.5-2-2.5-2v-2.1c2.2-.4 4-2.2 4-4.4V6c0-.6-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Minus(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 7h12v2H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusCircle(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8m5 9H3V7h10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusCircleO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1c3.9 0 7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7m0-1C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8"/><path fill="currentColor" d="M3 7h10v2H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusSquareO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 7h8v2H4z"/><path fill="currentColor" d="M15 1H1v14h14zm-1 13H2V2h12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mobile(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 1v14h8V1zm5 13H7v-1h2zm2-2H5V3h6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MobileBrowser(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 0H3v5H0v11h7v-3h9zM6 1h9v1H6zM4 1h1v1H4zm0 14H3v-1h1zm2-2H1V6h5zm9-1H7V5H4V3h11z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MobileRetro(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 0h-1v2H4v14h7zM6 14H5v-1h1zm0-2H5v-1h1zm0-2H5V9h1zm2 4H7v-1h1zm0-2H7v-1h1zm0-2H7V9h1zm2 4H9v-1h1zm0-2H9v-1h1zm0-2H9V9h1zm0-2H5V4h5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Modal(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 1v14h16V1zm15 13H1V4h14zm0-11h-1V2h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ModalList(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 6h2v1H3zm3 0h7v1H6zM3 8h2v1H3zm3 0h7v1H6zm-3 2h2v1H3zm3 0h7v1H6z"/><path fill="currentColor" d="M0 1v14h16V1zm15 13H1V4h14zm0-11h-1V2h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Money(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 4v8H1V4zm1-1H0v10h16z"/><path fill="currentColor" d="M8 5c1.7 0 3 1.3 3 3s-1.3 3-3 3h5v-1h1V6h-1V5zM5 8c0-1.7 1.3-3 3-3H3v1H2v4h1v1h5c-1.7 0-3-1.3-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoneyDeposit(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8 16l-2-3h1v-2h2v2h1zm7-15v8H1V1zm1-1H0v10h16z"/><path fill="currentColor" d="M8 2a3 3 0 1 1 0 6h5V7h1V3h-1V2zM5 5a3 3 0 0 1 3-3H3v1H2v4h1v1h5a3 3 0 0 1-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoneyExchange(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16 14l-3 2v-1H8.25l2-2H13v-1zM0 2l3-2v1h4.75l-2 2H3v1zm9.74-2L0 9.74L6.26 16L16 6.26zM1.39 9.74l8.35-8.35l4.87 4.87l-8.35 8.35z"/><path fill="currentColor" d="m4.17 9.74l-.7.7l2.09 2.09l.7-.7l.74.69l2.74-2.78a2.461 2.461 0 0 1-3.48-3.48L3.48 9z"/><path fill="currentColor" d="m12.52 5.57l-2.09-2.09l-.7.7l-.73-.7l-2.74 2.78a2.461 2.461 0 0 1 3.48 3.48L12.52 7l-.7-.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoneyWithdraw(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8 0l2 3H9v2H7V3H6zm7 7v8H1V7zm1-1H0v10h16z"/><path fill="currentColor" d="M8 8a3 3 0 1 1 0 6h5v-1h1V9h-1V8zm-3 3a3 3 0 0 1 3-3H3v1H2v4h1v1h5a3 3 0 0 1-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Moon(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8m0 15c-3.9 0-7-3.1-7-7c0-2.4 1.2-4.6 3.2-5.9C4.1 2.7 4 3.4 4 4c0 4.9 4 8.9 8.9 9c-1.3 1.3-3 2-4.9 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.2 11.9c-4.5 0-8.1-3.6-8.1-8.1c0-1.4.3-2.7.9-3.8c-3.4.9-6 4.1-6 7.9C0 12.4 3.6 16 8.1 16c3.1 0 5.8-1.8 7.2-4.4c-.6.2-1.3.3-2.1.3M8.1 15C4.2 15 1 11.8 1 7.9c0-2.5 1.3-4.7 3.3-6c-.2.6-.2 1.2-.2 1.9c0 5 4.1 9.1 9.1 9.2c-1.4 1.2-3.2 2-5.1 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Morning(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14 10l-1.58-1.18L13.2 7l-2-.23L11 4.8l-1.82.78L8 4L6.82 5.58L5 4.8l-.23 2L2.8 7l.78 1.82L2 10H0v1h16v-1zM4 10a4.143 4.143 0 0 1 3.993-4A4.143 4.143 0 0 1 12 9.993z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Movie(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 7V4H0v9h12v-3l4 2V5zm-3 4H2V6h7z"/><path fill="currentColor" d="M5 8.4a1 1 0 1 1-2 0a1 1 0 0 1 2 0m3 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Music(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 3v9.4c-.4-.2-.9-.4-1.5-.4c-1.4 0-2.5.9-2.5 2s1.1 2 2.5 2S5 15.1 5 14V6.7l7-2.3v5.1c-.4-.3-.9-.5-1.5-.5C9.1 9 8 9.9 8 11s1.1 2 2.5 2s2.5-.9 2.5-2V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mute(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.2 0L11 4.2V3c0-1.7-1.3-3-3-3S5 1.3 5 3v4c0 .9.4 1.7 1 2.2l-.8.8C4.5 9.4 4 8.5 4 7.5V5c-.6 0-1 .4-1 1v1.5c0 1.3.6 2.4 1.5 3.2L0 15.3v.7h.7L16 .6V0zm-2.7 5.1l-.5.5v1.9c0 1.9-1.8 3.5-3.8 3.5h-.4c-.3 0-.6-.1-.9-.1l-.9.7c.3.1.6.2 1 .3V14c-3 0-2.5 2-2.5 2h7s.5-2-2.5-2v-2.1c2.2-.4 4-2.2 4-4.4V6c0-.4-.2-.7-.5-.9"/><path fill="currentColor" d="M11 7v-.4L7.7 10H8c1.7 0 3-1.4 3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NativeButton(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 12H1c-.6 0-1-.4-1-1V5c0-.6.4-1 1-1h14c.6 0 1 .4 1 1v6c0 .6-.4 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Newspaper(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 4h11v4H2zm0-2h11v1H2zm6 11h3v1H8zm0-2h5v1H8zm0-2h5v1H8zm-6 4h5v1H2zm0-2h5v1H2zm0-2h5v1H2z"/><path fill="currentColor" d="M15 2V0H0v14.5A1.5 1.5 0 0 0 1.5 16h13a1.5 1.5 0 0 0 1.5-1.5V2zM1.5 15a.5.5 0 0 1-.5-.5V1h13v12.5c0 1.5 1 1.5 1 1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Notebook(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 0v1h-.52a.48.48 0 0 0-.48.48a.48.48 0 0 0 .478.52H2v1h-.52a.48.48 0 0 0-.48.48a.48.48 0 0 0 .478.52H2v1h-.52a.48.48 0 0 0-.48.48a.48.48 0 0 0 .478.52H2v1h-.52a.48.48 0 0 0-.48.48a.48.48 0 0 0 .478.52H2v1h-.52a.48.48 0 0 0 0 .96H2v1h-.52a.48.48 0 0 0 0 .96H2v1h-.52a.48.48 0 0 0 0 .96H2v2h12V0zm1.5 14a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0-2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0-2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0-2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0-2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0-2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1m0-2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1M12 6H6V3h6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Nurse(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.2 12a4.076 4.076 0 0 0-3.931-2.37L8 13.53l-3.28-3.9a4.16 4.16 0 0 0-3.909 2.345a9.072 9.072 0 0 0-.808 2.988L2 15v1h12v-1h2a9.199 9.199 0 0 0-.824-3.057z"/><path fill="currentColor" d="M6.57 8.73a.82.82 0 0 1-.685.729L8 12l2.12-2.52a.823.823 0 0 1-.69-.727c0-.613.8-.413 1.43-1.453C10.86 7.27 13.74 0 8 0S5.14 7.27 5.14 7.27c.63 1.05 1.44.85 1.43 1.46"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Office(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 15V4h-3V1H2v14H0v1h7v-3h2v3h7v-1zm-8-4H4V9h2zm0-3H4V6h2zm0-3H4V3h2zm3 6H7V9h2zm0-3H7V6h2zm0-3H7V3h2zm4 6h-2V9h2zm0-3h-2V6h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OpenBook(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 4.7V4a6.804 6.804 0 0 0-4.484-1.999a2.844 2.844 0 0 0-2.513.995a3.02 3.02 0 0 0-2.515-.995A6.804 6.804 0 0 0 1 4v.7L0 5v10l6.7-1.4l.3.4h2l.3-.4L16 15V5zm-9.52 6.61a8.206 8.206 0 0 0-3.526.902L2 4.42A5.22 5.22 0 0 1 5.369 3a4.553 4.553 0 0 1 2.159.701l-.019 7.869a6.548 6.548 0 0 0-2.039-.259zm8.52.88a8.122 8.122 0 0 0-3.468-.88l-.161-.002c-.66 0-1.297.096-1.899.274l.047-7.902a4.484 4.484 0 0 1 2.096-.679a5.216 5.216 0 0 1 3.386 1.422l-.003 7.768z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Option(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 11a1 1 0 0 0 2 0v-1H5a1 1 0 0 0-1 1"/><path fill="currentColor" d="M0 0v16h16V0zm11 9a2 2 0 1 1-2 2v-1H7v1a2 2 0 1 1-2-2h1V7H5a2 2 0 1 1 2-2v1h2V5a2 2 0 1 1 2 2h-1v2z"/><path fill="currentColor" d="M12 5a1 1 0 0 0-2 0v1h1a1 1 0 0 0 1-1M5 4a1 1 0 0 0 0 2h1V5a1 1 0 0 0-1-1m2 3h2v2H7zm3 4a1 1 0 1 0 1-1h-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OptionA(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 10H11V6h1.5A2.5 2.5 0 1 0 10 3.5V5H6V3.5A2.5 2.5 0 1 0 3.5 6H5v4H3.5A2.5 2.5 0 1 0 6 12.5V11h4v1.5a2.5 2.5 0 1 0 2.5-2.5M11 3.5A1.5 1.5 0 1 1 12.5 5H11zm-6 9A1.5 1.5 0 1 1 3.5 11H5zM5 5H3.5A1.5 1.5 0 1 1 5 3.5zm5 5H6V6h4zm2.5 4a1.5 1.5 0 0 1-1.5-1.5V11h1.5a1.5 1.5 0 0 1 0 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Options(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 3.5a1.5 1.5 0 1 1-3.001-.001A1.5 1.5 0 0 1 5 3.5"/><path fill="currentColor" d="M3.5 0C1.6 0 0 1.6 0 3.5S1.6 7 3.5 7S7 5.4 7 3.5S5.4 0 3.5 0m0 6C2.1 6 1 4.9 1 3.5S2.1 1 3.5 1S6 2.1 6 3.5S4.9 6 3.5 6m0 2C1.6 8 0 9.6 0 11.5S1.6 15 3.5 15S7 13.4 7 11.5S5.4 8 3.5 8m0 6C2.1 14 1 12.9 1 11.5S2.1 9 3.5 9S6 10.1 6 11.5S4.9 14 3.5 14M8 2h8v3H8zm0 8h8v3H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Orientation(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 2.1c2 0 3 1.3 3 2.9h-1l1.5 2L16 5h-1c0-2.2-2-3.9-4-3.9V0L9 1.5L11 3z"/><path fill="currentColor" d="M9 9h6v6H8V0H0v16h16V8H9zM7 8H6v1h1v6H1V1h6z"/><path fill="currentColor" d="M2 8h1v1H2zm2 0h1v1H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Out(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.5 8c.3 0 .5.2.5.5v2c0 .3-.2.5-.5.5s-.5-.2-.5-.5v-2c0-.3.2-.5.5-.5m0-1C2.7 7 2 7.7 2 8.5v2c0 .8.7 1.5 1.5 1.5S5 11.3 5 10.5v-2C5 7.7 4.3 7 3.5 7M8 7v3.5c0 .3-.2.5-.5.5s-.5-.2-.5-.5V7H6v3.5c0 .8.7 1.5 1.5 1.5S9 11.3 9 10.5V7zm5 0h-3v1h1v4h1V8h1z"/><path fill="currentColor" d="M15 6V5h-2.4L8.9 2c.1-.2.1-.3.1-.5C9 .7 8.3 0 7.5 0S6 .7 6 1.5c0 .2 0 .3.1.5L2.4 5H0v9h1v1h15V6zM6.7 2.8c.3.1.5.2.8.2s.5-.1.8-.2L11 5H4zM14 13H1V6h13z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Outbox(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 5v6h4V5h2L8 0L4 5z"/><path fill="currentColor" d="M13 2h-2l.9 1h.4l2.6 8H11v2H5v-2H1.1l2.6-8h.4L5 2H3l-3 9v5h16v-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Package(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0L0 2v10l8 4l8-4V2zm0 1l2.1.5l-5.9 1.9l-2.3-.8zm0 13.9l-7-3.5V3.3l3 1v3.4L5 8V4.7l3 1zm.5-10.1l-2.7-.9L12 2l2.4.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Padding(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v16h16V0zm15 3h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v-1h-1v1h-1v-1h-1v1h-1v-1H9v1H8v-1H7v1H6v-1H5v1H4v-1H3v1H2v-1H1v-1h1v-1H1v-1h1v-1H1V9h1V8H1V7h1V6H1V5h1V4H1V3h1V2H1V1h1v1h1V1h1v1h1V1h1v1h1V1h1v1h1V1h1v1h1V1h1v1h1V1h1v1h1z"/><path fill="currentColor" d="M3 2h1v1H3zm1 1h1v1H4zm2 0h1v1H6zM5 2h1v1H5zm2 0h1v1H7zm2 0h1v1H9zM8 3h1v1H8zm2 0h1v1h-1zm2 0h1v1h-1zm-1-1h1v1h-1zm2 0h1v1h-1zm-1 3h1v1h-1zm1-1h1v1h-1zm-1 3h1v1h-1zm1-1h1v1h-1zm-1 3h1v1h-1zm1-1h1v1h-1zm-1 3h1v1h-1zm1-1h1v1h-1zm-1 3h1v1h-1zm1-1h1v1h-1zM2 3h1v1H2zm1 1h1v1H3zM2 5h1v1H2zm1 1h1v1H3zM2 7h1v1H2zm1 1h1v1H3zM2 9h1v1H2zm1 1h1v1H3zm-1 1h1v1H2zm0 2h1v1H2zm1-1h1v1H3zm1-1h1v1H4zm0 2h1v1H4zm1-1h1v1H5zm1 1h1v1H6zm1-1h1v1H7zm2 0h1v1H9zm-1 1h1v1H8zm3-1h1v1h-1zm-1 1h1v1h-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaddingBottom(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 16V0H0v16zM1 13h1v-1H1V1h14v12h-1v1h1v1h-1v-1h-1v1h-1v-1h-1v1h-1v-1H9v1H8v-1H7v1H6v-1H5v1H4v-1H3v1H2v-1H1z"/><path fill="currentColor" d="M12 13h1v1h-1zm1-1h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1H9zm1 1h1v1h-1zm-2 0h1v1H8zm-2 0h1v1H6zm1-1h1v1H7zm-2 0h1v1H5zm-2 0h1v1H3zm1 1h1v1H4zm-2 0h1v1H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaddingLeft(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 16h16V0H0zM3 1v1h1V1h11v14H3v-1H2v1H1v-1h1v-1H1v-1h1v-1H1v-1h1V9H1V8h1V7H1V6h1V5H1V4h1V3H1V2h1V1z"/><path fill="currentColor" d="M2 12h1v1H2zm1 1h1v1H3zm0-2h1v1H3zm0-2h1v1H3zm-1 1h1v1H2zm0-2h1v1H2zm0-2h1v1H2zm1 1h1v1H3zm0-2h1v1H3zm0-2h1v1H3zM2 4h1v1H2zm0-2h1v1H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaddingRight(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 0H0v16h16zm-3 15v-1h-1v1H1V1h12v1h1V1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1h1v1h-1v1z"/><path fill="currentColor" d="M13 3h1v1h-1zm-1-1h1v1h-1zm0 2h1v1h-1zm0 2h1v1h-1zm1-1h1v1h-1zm0 2h1v1h-1zm0 2h1v1h-1zm-1-1h1v1h-1zm0 2h1v1h-1zm0 2h1v1h-1zm1-1h1v1h-1zm0 2h1v1h-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaddingTop(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v16h16V0zm15 3h-1v1h1v11H1V3h1V2H1V1h1v1h1V1h1v1h1V1h1v1h1V1h1v1h1V1h1v1h1V1h1v1h1V1h1v1h1z"/><path fill="currentColor" d="M3 2h1v1H3zM2 3h1v1H2zm2 0h1v1H4zm2 0h1v1H6zM5 2h1v1H5zm2 0h1v1H7zm2 0h1v1H9zM8 3h1v1H8zm2 0h1v1h-1zm2 0h1v1h-1zm-1-1h1v1h-1zm2 0h1v1h-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaintRoll(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 6.9V2h-2V0H1v1H0v3h1v1h13V3h1v3.1l-8 1V9H6v.9s.5 0 .5.9s-.5.6-.5 1.5v2.8s0 .9 1.5.9s1.5-.9 1.5-.9v-2.8c0-.9-.5-.7-.5-1.5s.5-.9.5-.9V9H8V7.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paintbrush(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.6 11.6l-1.2-1.2c-.8-.2-2-.1-2.7 1c-.8 1.1-.3 2.8-1.7 4.6c0 0 3.5 0 4.8-1.3c1.2-1.2 1.2-2.2 1-3zm.2-3.5c-.2.3-.5.7-.7 1c0 .2-.1.3-.2.4L6.4 11c.1-.1.3-.2.4-.3c.3-.2.7-.4 1-.7c.4 0 .6-.2.8-.4L6.4 7.4c-.2.2-.4.4-.6.7m10-7.9c-.3-.3-.7-.3-1-.1c0 0-3 2.5-5.9 5.1c-.4.4-.7.7-1.1 1c-.2.2-.4.4-.6.5l2.1 2.1c.2-.2.4-.4.5-.7c.3-.4.6-.7.9-1.1c2.5-3 5.1-5.9 5.1-5.9c.3-.2.3-.6 0-.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Palete(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.25 0C1.87 0-.86 7.38.24 9.92c.82 1.89 2.62.08 3.34 1c1.88 2.46-2.11 3.81.09 4.68C6.26 16.66 16 16 16 7.07C16 4.38 14.66 0 8.25 0M4.47 9A1.5 1.5 0 1 1 6 7.5A1.5 1.5 0 0 1 4.5 9h-.032zM6 3.5A1.5 1.5 0 1 1 7.5 5h-.032A1.5 1.5 0 0 1 6 3.5M8.47 14A1.5 1.5 0 1 1 10 12.5A1.5 1.5 0 0 1 8.5 14h-.032zm4-3A1.5 1.5 0 1 1 14 9.5a1.5 1.5 0 0 1-1.5 1.5h-.032zm0-5A1.5 1.5 0 1 1 14 4.5A1.5 1.5 0 0 1 12.5 6h-.032z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Panel(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v16h16V0zm13 15H1V3h12zm2 0h-1v-1h1zm0-2h-1V5h1zm0-9h-1V3h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paperclip(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.7 15.3c-.7 0-1.4-.3-1.9-.8c-.9-.9-1.2-2.5 0-3.7l8.9-8.9c1.4-1.4 3.8-1.4 5.2 0s1.4 3.8 0 5.2l-7.4 7.4c-.2.2-.5.2-.7 0s-.2-.5 0-.7l7.4-7.4c1-1 1-2.7 0-3.7s-2.7-1-3.7 0l-8.9 8.9c-.8.8-.6 1.7 0 2.2c.6.6 1.5.8 2.2 0l8.9-8.9c.2-.2.2-.5 0-.7s-.5-.2-.7 0l-7.4 7.4c-.2.2-.5.2-.7 0s-.2-.5 0-.7l7.4-7.4c.6-.6 1.6-.6 2.2 0s.6 1.6 0 2.2l-8.9 8.9c-.6.4-1.3.7-1.9.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paperplane(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 8l4.9 1.4H5v-.1L12.1 4L11 5.2l-6.2 6.6L5 15l2.9-3.2L10 16l6-16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaperplaneO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 0L0 8l4.7 1.6L5 15l2.5-2.8L10 16zM7.5 10.4l4.3-5.9l-6.2 4.3l-3-1L14.2 2L9.7 13.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paragraph(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.5 0C3 0 1 2 1 4.5S3 9 5.5 9H8v7h2V2h1v14h2V2h2V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Password(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 5c0-.6-.4-1-1-1H1c-.6 0-1 .4-1 1v6c0 .6.4 1 1 1h14c.6 0 1-.4 1-1zm-1 6H1V5h14z"/><path fill="currentColor" d="M6 8a1 1 0 1 1-2 0a1 1 0 0 1 2 0m3 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0m3 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paste(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 4h-3V0H0v14h6v2h10V7zM3 1h4v1H3zm12 14H7V5h5v3h3zm-2-8V5l2 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pause(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 1h7v14H0zm9 0h7v14H9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pencil(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 11.9L0 16l4.1-1l9.2-9.2l-3.1-3.1zm.5 3.1l-.4-.5l.4-2l2 2zm9.4-10.6l-8.1 8l-.6-.6l8.1-8zM15.3.7C14.2-.4 12.7.2 12.7.2l-1.5 1.5l3.1 3.1l1.5-1.5c0-.1.6-1.5-.5-2.6m-1.9.9l-.5-.5C13.5.5 14 1 14 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Phone(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.2 10c-1.1-.1-1.7 1.4-2.5 1.8C8.4 12.5 6 10 6 10S3.5 7.6 4.1 6.3c.5-.8 2-1.4 1.9-2.5c-.1-1-2.3-4.6-3.4-3.6C.2 2.4 0 3.3 0 5.1c-.1 3.1 3.9 7 3.9 7c.4.4 3.9 4 7 3.9c1.8 0 2.7-.2 4.9-2.6c1-1.1-2.5-3.3-3.6-3.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneLandline(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.88 3.86l-.61-1.31a1.214 1.214 0 0 0-.792-.658c-1.938-.528-4.161-.851-6.453-.891a27.46 27.46 0 0 0-6.687.934c-.165.048-.453.29-.605.609L.12 3.861a1.195 1.195 0 0 0-.12.52v.87l-.001.041c0 .392.318.71.71.71l.033-.001H3.26a.76.76 0 0 0 .742-.76L4 5.188v-.85a.65.65 0 0 1 .298-.546a6.913 6.913 0 0 1 3.724-.791a6.965 6.965 0 0 1 3.717.788c.143.099.262.3.262.529v.872a.76.76 0 0 0 .739.81h2.521l.031.001a.71.71 0 0 0 .71-.71l-.001-.043V4.38a1.208 1.208 0 0 0-.123-.527z"/><path fill="currentColor" d="M12 8.3a4.73 4.73 0 0 1-1.001-2.92L11 5.296V5h-1v1H6V5H5v.33l.001.08a4.74 4.74 0 0 1-1.009 2.93L1 12v3h14v-3zM8 13a3 3 0 1 1 0-6a3 3 0 0 1 0 6"/><path fill="currentColor" d="M10 10a2 2 0 1 1-3.999.001A2 2 0 0 1 10 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Picture(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 14H0V2h16zM1 13h14V3H1z"/><path fill="currentColor" d="M2 10v2h12v-1s.2-1.7-2-2c-1.9-.3-2.2.6-3.8.6C7.1 9.6 7.3 8 5 8c-1.7 0-3 2-3 2m11-4a2 2 0 1 1-3.999.001A2 2 0 0 1 13 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PieBarChart(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 11h3v5H5zm-4 3h3v2H1zm12-2h3v4h-3zM9 9h3v7H9zM5 0a5 5 0 1 0 .001 10.001A5 5 0 0 0 5 0m0 9a4 4 0 0 1 0-8v4h4a4 4 0 0 1-4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PieChart(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 1c3.2.2 5.7 2.8 6 6H9zm-.5-1H8v8h8v-.5C16 3.4 12.6 0 8.5 0"/><path fill="currentColor" d="M7 9V1c-3.9.3-7 3.5-7 7.5C0 12.6 3.4 16 7.5 16c4 0 7.2-3.1 7.5-7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PiggyBank(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.93 5.75a1.25 1.25 0 0 0-.3-.51a.833.833 0 0 0-.394-.238c.074.117.141.252.191.396c.056.147.092.304.103.467a1.784 1.784 0 0 1 0 .424a1.005 1.005 0 0 0-.142-.292a1.193 1.193 0 0 0-.48-.383a.938.938 0 0 0-1.195.382c-.05.082-.09.171-.12.266c-1.182-1.968-3.309-3.271-5.741-3.271c-.124 0-.247.003-.369.01a8.217 8.217 0 0 0-2.231.313C5.19 2.88 4.62 2 2 2l.8 2.51a5.207 5.207 0 0 0-1.247 1.465L0 6s-.17 4 1 4h.54a5.276 5.276 0 0 0 1.445 1.589L3 13.999h1.08c1.31 0 1.92 0 1.92-.75v-.39a8.256 8.256 0 0 0 3.051-.008L9 13.249c0 .75.62.75 1.94.75H12v-2.39a4.79 4.79 0 0 0 1.903-2.717c.057-.027.114-.024.172-.024s.115-.003.172-.01c.251-.046.48-.144.679-.283a1.65 1.65 0 0 0 .591-.765c.028-.093.049-.191.063-.292l.001-.01c.221-.262.372-.59.419-.951a1.776 1.776 0 0 0-.072-.822zm-12.42 0a.75.75 0 1 1 0 1.5a.75.75 0 0 1 0-1.5M5.88 5a.502.502 0 0 1-.599-.247a.39.39 0 0 1 .296-.503a8.024 8.024 0 0 1 2.009-.22l.101-.001c.672 0 1.324.08 1.949.232c.126.024.262.182.262.372a.385.385 0 0 1-.019.119a.483.483 0 0 1-.346.247H9.38a7.198 7.198 0 0 0-1.706-.2h-.089c-.605 0-1.193.073-1.756.211zm8.7 2.93a1.16 1.16 0 0 1-.453.199L14 8.14v-.45c.165.125.374.2.6.2h.021zm.08-.68a.44.44 0 0 1-.459-.248a.607.607 0 0 1 .001-.566a.332.332 0 0 1 .43-.096a.48.48 0 0 1 .308.448v.001a1.457 1.457 0 0 1-.001.418a1.26 1.26 0 0 1-.282.022z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PiggyBankCoin(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.93 7.75a1.25 1.25 0 0 0-.3-.51a.833.833 0 0 0-.394-.238c.074.117.141.252.191.396c.056.147.092.304.103.467a1.784 1.784 0 0 1 0 .424a1.005 1.005 0 0 0-.142-.292a1.193 1.193 0 0 0-.48-.383a.938.938 0 0 0-1.195.382c-.05.082-.09.171-.12.266c-1.182-1.968-3.309-3.271-5.741-3.271c-.124 0-.247.003-.369.01a8.217 8.217 0 0 0-2.231.313C5.19 4.88 4.62 4 2 4l.8 2.51a5.207 5.207 0 0 0-1.247 1.465L0 8s-.17 4 1 4h.54a5.276 5.276 0 0 0 1.445 1.589L3 16h1.08C5.39 16 6 16 6 15.25v-.39a8.256 8.256 0 0 0 3.051-.008L9 15.25c0 .75.62.75 1.94.75H12v-2.39a4.79 4.79 0 0 0 1.903-2.717c.057-.027.114-.024.172-.024s.115-.003.172-.01c.251-.046.48-.144.679-.283a1.65 1.65 0 0 0 .591-.765c.028-.093.049-.191.063-.292l.001-.01c.221-.262.372-.59.419-.951a1.776 1.776 0 0 0-.072-.822zm-12.42 0a.75.75 0 1 1 0 1.5a.75.75 0 0 1 0-1.5M5.88 7a.502.502 0 0 1-.599-.247a.39.39 0 0 1 .296-.503a8.024 8.024 0 0 1 2.009-.22l.101-.001c.672 0 1.324.08 1.949.232c.126.024.262.182.262.372a.385.385 0 0 1-.019.119a.483.483 0 0 1-.346.247H9.38a7.198 7.198 0 0 0-1.706-.2h-.089c-.605 0-1.193.073-1.756.211zm8.7 2.93a1.16 1.16 0 0 1-.453.199L14 10.13v-.44c.165.125.374.2.6.2h.021zm.08-.68a.44.44 0 0 1-.459-.248a.607.607 0 0 1 .001-.566a.332.332 0 0 1 .43-.096a.48.48 0 0 1 .308.448v.001a1.457 1.457 0 0 1-.001.418a1.26 1.26 0 0 1-.282.022zM8 3H7v-.17h.25V1.74H7l.55-.55h.2v1.64H8z"/><path fill="currentColor" d="M7.5.75a1.5 1.5 0 1 1-.001 3.001A1.5 1.5 0 0 1 7.5.75m0-.75a2.25 2.25 0 1 0 0 4.5a2.25 2.25 0 0 0 0-4.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pill(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14.8 1.4l-.2-.2C13.9.4 12.8 0 11.8 0S9.7.4 8.9 1.2L1.2 8.9c-1.6 1.6-1.6 4.1 0 5.7l.2.2c.7.8 1.8 1.2 2.8 1.2s2.1-.4 2.9-1.2L14.9 7c1.5-1.5 1.5-4.1-.1-5.6m-.7 5l-3.9 3.9l-3.5-3.6l-3.8 3.8c-1.1 1.1-1.1 2.5-1 3.5c-1.2-1.2-1.2-3.1 0-4.3l7.8-7.8c.5-.6 1.3-.9 2.1-.9s1.6.3 2.2.9l.2.2c.5.5.8 1.3.8 2.1s-.3 1.6-.9 2.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pills(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m3.5 8l6.3-6.3c.4-.4 1-.7 1.7-.7s1.3.3 1.8.7c1 1 1 2.6 0 3.5L10.5 8h1.4l2-2c1.4-1.4 1.4-3.6 0-4.9c-.7-.7-1.6-1-2.5-1S9.7.3 9 1L2.7 7.4c-.3.2-.5.5-.7.9c.5-.2 1-.3 1.5-.3"/><path fill="currentColor" d="M7.3 5.6L4.9 8h4.7zM12.5 9h-9C1.6 9 0 10.6 0 12.5S1.6 16 3.5 16h9c1.9 0 3.5-1.6 3.5-3.5S14.4 9 12.5 9m0 6H8v-4H3.5c-1.1 0-2 .6-2.5 1.2C1.2 11 2.2 10 3.5 10h9c1.4 0 2.5 1.1 2.5 2.5S13.9 15 12.5 15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pin(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 6.5V1h1V0H4v1h1v5.5S3 8 3 10c0 .5 1.9.7 4 .7v2.2c0 .7.2 1.4.5 2.1l.5 1l.5-1c.3-.6.5-1.3.5-2.1v-2.2c2.1 0 4-.3 4-.7c0-2-2-3.5-2-3.5m-4 .1s-.5.3-1.6 1.4c-1 1-1.5 1.9-1.5 1.9s.1-1 .8-1.9C5.6 6.9 6 6.6 6 6.6V1h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinPost(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 4V3H9c0-1.69 1-2 1-2V0H5v1s1 .31 1 2H0v12h2v1h14V4zm-1 10H1V4h4v1h2v2h1V5h2V4h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Play(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 1v14l12-7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayCircle(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8M6 12V4l6 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayCircleO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1c3.9 0 7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7m0-1C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8"/><path fill="currentColor" d="M6 4v8l6-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plug(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.7 3.1c-.4-.4-1-.4-1.4 0l-2.8 2.8L9 4.5l2.8-2.8c.4-.4.4-1 0-1.4s-1-.4-1.4 0L7.6 3.1L6.2 1.7L4.8 3.1l.7.7l-1.4 1.4c-1.4 1.4-1.5 3.5-.5 5.1C1.9 11.8 1 14.1 1 16h2c0-1.3.4-3.2 2.1-4.4c1.5.8 3.4.5 4.6-.7l1.4-1.4l.7.7l1.4-1.4l-1.4-1.4l2.8-2.8c.5-.5.5-1.1.1-1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plus(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 7H9V2H7v5H2v2h5v5h2V9h5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusCircle(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8m5 9H9v4H7V9H3V7h4V3h2v4h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusCircleO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1c3.9 0 7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7m0-1C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8"/><path fill="currentColor" d="M13 7H9V3H7v4H3v2h4v4h2V9h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusMinus(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 7h6v2h-6zM4 5H2v2H0v2h2v2h2V9h2V7H4zm2-3l3 12h1L7 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusSquareO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 7H9V4H7v3H4v2h3v3h2V9h3z"/><path fill="currentColor" d="M15 1H1v14h14zm-1 13H2V2h12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pointer(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.6 5H12c0-.2-.2-.6-.4-.8s-.6-.4-1.1-.4c-.2 0-.4 0-.6.1c-.1-.2-.2-.3-.3-.5c-.2-.2-.5-.4-1.1-.4c-.2 0-.4 0-.5.1V1.4C8 .8 7.6 0 6.6 0c-.4 0-.8.2-1.1.4C5 1 5 1.8 5 1.8v4.3c-.6.1-1.1.3-1.4.6C3 7.4 3 8.3 3 9.5v.7c0 1.4.7 2.1 1.4 2.8l.3.4C6 14.6 7.1 15 9.8 15c2.9 0 4.2-1.6 4.2-5.1V7.4c0-.7-.2-2.1-1.4-2.4m.4 2.4V10c0 3.4-1.3 4.1-3.2 4.1c-2.4 0-3.3-.3-4.3-1.3l-.4-.4c-.7-.8-1.1-1.2-1.1-2.2v-.7c0-1 0-1.7.3-2.1c.1-.1.4-.2.7-.2v.5l-.3 1.5c0 .1 0 .1.1.2s.2 0 .2 0l1-1.2V1.8c0-.1 0-.5.2-.7c.1 0 .2-.1.4-.1c.3 0 .4.3.4.4v4.3c0 .3.2.6.5.6S8 6 8 5.8V4.5c0-.1.1-.5.5-.5c.3 0 .5.1.5.4v1.3c0 .3.2.6.5.6s.5-.3.5-.5v-.7c0-.1.1-.3.5-.3c.2 0 .3.1.3.1c.2.1.2.4.2.4v.8c0 .3.2.5.4.5c.3 0 .5-.1.5-.4c0-.1.1-.2.2-.3h.2c.6.2.7 1.2.7 1.5c0-.1 0-.1 0 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PowerOff(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 2.3v3.3c1.2.7 2 2 2 3.4c0 2.2-1.8 4-4 4s-4-1.8-4-4c0-1.5.8-2.8 2-3.4V2.3C3.1 3.2 1 5.8 1 9c0 3.9 3.1 7 7 7s7-3.1 7-7c0-3.2-2.1-5.8-5-6.7"/><path fill="currentColor" d="M7 1h2v7H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Presentation(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 1H9V0H7v1H0v11h5l-2 4h2.2l2-4h1.5l2 4H13l-2-4h5zm-1 10H1V2h14z"/><path fill="currentColor" d="M6 4v5l4-2.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Print(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 10v4h2v2h12v-2h2v-4zm13 5H3v-3h10zm-1-9V2L9.3 0H4v6H0v3h16V6zM9 1l1.3 1H9zm2 6H5V1h3v2h3zm4 1h-1V7h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Progressbar(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 5v6h16V5zm15 5H1V6h14z"/><path fill="currentColor" d="M2 7h7v2H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PuzzlePiece(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.9.9c-1.1-1-2.5-1.3-3.1-.4c-.7 1.1.5 1.7-.3 2.5c-.5.6-2-.8-2-.8l-.8-.8l-1.4 1.4c-.6.7-2.1 1.5-2.6 1.1c-.7-.6.1-1.8-.5-2.6c-.7-1-2.1-.8-3 .3c-1 1.1-1.4 2.4-.5 3c1.1.7 1.9-.3 2.7.5c.4.4-.2 1.7-.5 2.1L.6 9.5L7.1 16l1.7-1.7c.7-.7 1.5-2 1.1-2.4c-.6-.7-1.7.1-2.5-.4c-1-.7-.8-2 .3-3s2.5-1.3 3.1-.4c.7 1.1-.4 1.8.4 2.6c.4.4 1.6-.2 2-.6L15.3 8l-1.1-1.1c-.6-.6-1.9-2-1.4-2.5c.6-.7 1.7.2 2.5-.4c.9-.8.6-2.1-.4-3.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PyramidChart(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.29 5L8 1L5.71 5zm-8 6L0 15h16l-2.29-4zm10.85-1l-2.28-4H5.14l-2.28 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Qrcode(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 0H0v6h6zM5 5H1V1h4z"/><path fill="currentColor" d="M2 2h2v2H2zM0 16h6v-6H0zm1-5h4v4H1z"/><path fill="currentColor" d="M2 12h2v2H2zm8-12v6h6V0zm5 5h-4V1h4z"/><path fill="currentColor" d="M12 2h2v2h-2zM2 7H0v2h3V8H2zm5 2h2v2H7zM3 7h2v1H3zm6 5H7v1h1v1h1v-1zM6 7v1H5v1h2V7zm2-3h1v2H8zm1 4v1h2V7H8v1zM7 6h1v1H7zm2 8h2v2H9zm-2 0h1v2H7zm2-3h1v1H9zm0-8V1H8V0H7v4h1V3zm3 11h1v2h-1zm0-2h2v1h-2zm-1 1h1v1h-1zm-1-1h1v1h-1zm4-2v1h1v1h1v-2h-1zm1 3h-1v3h2v-2h-1zm-5-3v1h3V9h-2v1zm2-3v1h2v1h2V7h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Question(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 11H6c0-3 1.6-4 2.7-4.6c.4-.2.7-.4.9-.6c.5-.5.3-1.2.2-1.4c-.3-.7-1-1.4-2.3-1.4C5.4 3 5 4.9 5 5.3l-3-.4C2.2 3.2 3.7 0 7.5 0c2.3 0 4.3 1.3 5.1 3.2c.7 1.7.4 3.5-.8 4.7c-.5.5-1.1.8-1.6 1.1c-.9.5-1.2 1-1.2 2m.5 3a2 2 0 1 1-3.999.001A2 2 0 0 1 9.5 14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionCircle(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8m.9 13h-2v-2h2zM11 8.1c-.4.4-.8.6-1.2.7c-.6.4-.8.2-.8 1.2H7c0-2 1.2-2.6 2-3c.3-.1.5-.2.7-.4c.1-.1.3-.3.1-.7c-.2-.5-.8-1-1.7-1c-1.4 0-1.6 1.2-1.7 1.5l-2-.3C4.5 5 5.4 2.9 8 2.9c1.6 0 3 .9 3.6 2.2c.4 1.1.2 2.2-.6 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionCircleO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 10H7c0-2 1.2-2.6 2-3c.3-.1.5-.2.7-.4c.1-.1.3-.3.1-.7c-.2-.5-.8-1-1.7-1c-1.4 0-1.6 1.2-1.7 1.5l-2-.3C4.5 5 5.4 2.9 8 2.9c1.6 0 3 .9 3.6 2.2c.4 1.1.2 2.2-.6 3c-.4.4-.8.6-1.2.7c-.6.4-.8.2-.8 1.2"/><path fill="currentColor" d="M8 1c3.9 0 7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7m0-1C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8"/><path fill="currentColor" d="M6.9 11h2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuoteLeft(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 7v7H0V6.9c0-4.8 4.5-5.4 4.5-5.4l.6 1.4s-2 .3-2.4 1.9C2.3 6 3.1 7 3.1 7zm9 0v7H9V6.9c0-4.8 4.5-5.4 4.5-5.4l.6 1.4s-2 .3-2.4 1.9c-.4 1.2.4 2.2.4 2.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuoteRight(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 9V2h7v7.1c0 4.8-4.5 5.4-4.5 5.4l-.6-1.4s2-.3 2.4-1.9c.4-1.2-.4-2.2-.4-2.2zM0 9V2h7v7.1c0 4.8-4.5 5.4-4.5 5.4l-.6-1.4s2-.3 2.4-1.9C4.7 10 3.9 9 3.9 9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Random(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 12h-2c-1 0-1.7-1.2-2.4-2.7c-.3.7-.6 1.5-1 2.3C8.4 13 9.4 14 11 14h2v2l3-3l-3-3zM5.4 6.6c.3-.7.6-1.5 1-2.2C5.6 3 4.5 2 3 2H0v2h3c1 0 1.7 1.2 2.4 2.6"/><path fill="currentColor" d="m16 3l-3-3v2h-2C8.3 2 7.1 5 6 7.7C5.2 9.8 4.3 12 3 12H0v2h3c2.6 0 3.8-2.8 4.9-5.6C8.8 6.2 9.7 4 11 4h2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Raster(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 7h1v1H7zM5 7h1v1H5zM3 7h1v1H3zM1 7h1v1H1zm5-1h1v1H6zM4 6h1v1H4zM2 6h1v1H2zM0 6h1v1H0zm7-1h1v1H7zM5 5h1v1H5zM3 5h1v1H3zM1 5h1v1H1zm5-1h1v1H6zM4 4h1v1H4zM2 4h1v1H2zM0 4h1v1H0zm7-1h1v1H7zM5 3h1v1H5zM3 3h1v1H3zM1 3h1v1H1zm5-1h1v1H6zM4 2h1v1H4zM2 2h1v1H2zM0 2h1v1H0zm7-1h1v1H7zM5 1h1v1H5zM3 1h1v1H3zM1 1h1v1H1zm5-1h1v1H6zM4 0h1v1H4zM2 0h1v1H2zM0 0h1v1H0zm15 7h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1h-1zM9 7h1v1H9zm5-1h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1h-1zM8 6h1v1H8zm7-1h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1h-1zM9 5h1v1H9zm5-1h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1h-1zM8 4h1v1H8zm7-1h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1h-1zM9 3h1v1H9zm5-1h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1h-1zM8 2h1v1H8zm7-1h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1h-1zM9 1h1v1H9zm5-1h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1h-1zM8 0h1v1H8zM7 15h1v1H7zm-2 0h1v1H5zm-2 0h1v1H3zm-2 0h1v1H1zm5-1h1v1H6zm-2 0h1v1H4zm-2 0h1v1H2zm-2 0h1v1H0zm7-1h1v1H7zm-2 0h1v1H5zm-2 0h1v1H3zm-2 0h1v1H1zm5-1h1v1H6zm-2 0h1v1H4zm-2 0h1v1H2zm-2 0h1v1H0zm7-1h1v1H7zm-2 0h1v1H5zm-2 0h1v1H3zm-2 0h1v1H1zm5-1h1v1H6zm-2 0h1v1H4zm-2 0h1v1H2zm-2 0h1v1H0zm7-1h1v1H7zM5 9h1v1H5zM3 9h1v1H3zM1 9h1v1H1zm5-1h1v1H6zM4 8h1v1H4zM2 8h1v1H2zM0 8h1v1H0zm15 7h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1H9zm5-1h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1H8zm7-1h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1H9zm5-1h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1H8zm7-1h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1H9zm5-1h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1H8zm7-1h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1h-1zM9 9h1v1H9zm5-1h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1h-1zM8 8h1v1H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RasterLowerLeft(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 7h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1h-1zM9 7h1v1H9zm5-1h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1h-1zm5-1h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1h-1zm3-1h1v1h-1zm-2 0h1v1h-1zm3-1h1v1h-1zm-2 0h1v1h-1zm1-1h1v1h-1zm1-1h1v1h-1zM7 15h1v1H7zm-2 0h1v1H5zm-2 0h1v1H3zm-2 0h1v1H1zm5-1h1v1H6zm-2 0h1v1H4zm-2 0h1v1H2zm5-1h1v1H7zm-2 0h1v1H5zm-2 0h1v1H3zm3-1h1v1H6zm-2 0h1v1H4zm3-1h1v1H7zm-2 0h1v1H5zm1-1h1v1H6zm1-1h1v1H7zm8 6h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1H9zm5-1h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1H8zm7-1h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1H9zm5-1h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1H8zm7-1h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1H9zm5-1h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1H8zm7-1h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1h-1zM9 9h1v1H9zm5-1h1v1h-1zm-2 0h1v1h-1zm-2 0h1v1h-1zM8 8h1v1H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Records(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 9h4v2H4z"/><path fill="currentColor" d="M16 2h-1V0H5v2H3v1.25L2.4 4H1v1.75L0 7v9h12l4-5zM2 5h8v2H2zm9 10H1V8h10zm1-8h-1V4H4V3h8zm2-2.5l-1 1.25V2H6V1h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Recycle(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8 3.1l1.4 2.2l-1.6 1.1l1.3.3l2.8.6l.6-2.7l.4-1.4l-1.8 1.1l-2-3.3H6.9L4.3 5.3l1.7 1zm8 8.9l-2.7-4.3l-1.7 1l2 3.3H11v-2l-3 3l3 3v-2h3.7zM2.4 12l1.4-2.3l1.7 1.1l-.9-4.2l-2.8.7l-1.3.3l1.6 1L0 12l1.3 2H7v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Refresh(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.6 5.6C3.5 3.5 5.6 2 8 2c3 0 5.4 2.2 5.9 5h2c-.5-3.9-3.8-7-7.9-7c-3 0-5.6 1.6-6.9 4.1L0 3v4h4zM16 9h-4.1l1.5 1.4c-.9 2.1-3 3.6-5.5 3.6C5 14 2.5 11.8 2 9H0c.5 3.9 3.9 7 7.9 7c3 0 5.6-1.7 7-4.1L16 13z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reply(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 8c0-5-4.9-5-4.9-5H6V0L0 6l6 6V9h5.2c3.5 0 1.8 7 1.8 7s3-4.1 3-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReplyAll(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 8c0-5-4.9-5-4.9-5H9V0L3 6l6 6V9h2.2c3.5 0 1.8 7 1.8 7s3-4.1 3-8"/><path fill="currentColor" d="m0 6l6 6v-1.5L1.5 6L6 1.5V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ResizeH(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 7h16v2H0zm7-1h2V3h2L8 0L5 3h2zm2 4H7v3H5l3 3l3-3H9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ResizeV(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 0h2v16H7zM3 5L0 8l3 3V9h3V7H3zm13 3l-3-3v2h-3v2h3v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Retweet(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 1h12v5h2l-3 3l-3-3h2V3H4v2H2zm12 13H2V9H0l3-3l3 3H4v3h8v-2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rhombus(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0L0 8l8 8l8-8zM2 8l6-6l6 6l-6 6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Road(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 11v4h7L12 1H9v3H7V1H4L0 15h7v-4zM7 6h2v3H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RoadBranch(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 4H0v3h3.2L7 10.6c1.6 1.5 3.6 2.4 5.8 2.4H16v-3h-3.2c-1.4 0-2.7-.5-3.7-1.5L7.5 7H16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RoadBranches(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 4V1H0v3h1.7l7.7 9.5c1.3 1.6 3.1 2.5 5 2.5H16v-3h-1.5c-1 0-1.9-.5-2.7-1.4L10.5 10H16V7H8L5.6 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RoadSplit(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 13v-1c0-.2 0-4.1-2.8-5.4C9 5.6 9 3.1 9 3V0H7v3c0 .1 0 2.6-2.2 3.6C2 7.9 2 11.8 2 12v1H0l3 3l3-3H4v-1s0-2.8 1.7-3.6c1.1-.5 1.8-1.3 2.3-2c.5.8 1.2 1.5 2.3 2C12 9.2 12 12 12 12v1h-2l3 3l3-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rocket(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 0s-3.5-.4-6.7 2.8C7.7 4.3 6.4 6.3 5.4 8.1l-2.5-.6l-1.6 1.6l2.8 1.4c-.3.6-.4 1-.4 1l.8.8s.4-.2 1-.4l1.4 2.8l1.6-1.6l-.5-2.5c1.7-1 3.8-2.3 5.3-3.8C16.4 3.6 16 0 16 0m-3.2 4.8c-.4.4-1.1.4-1.6 0c-.4-.4-.4-1.1 0-1.6c.4-.4 1.1-.4 1.6 0c.4.4.4 1.1 0 1.6"/><path fill="currentColor" d="M4 14.2c-.8.8-2.6.4-2.6.4s-.4-1.8.4-2.6s1.5-.9 1.5-.9s-1.3-.3-2.1.6c-1.6 1.6-1 4.2-1 4.2s2.6.6 4.2-1c.9-.9.6-2.2.6-2.2s-.2.7-1 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RotateLeft(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0C5 0 2.4 1.6 1.1 4.1L0 3v4h4L2.5 5.5C3.5 3.5 5.6 2 8 2c3.3 0 6 2.7 6 6s-2.7 6-6 6c-1.8 0-3.4-.8-4.5-2.1L2 13.2C3.4 14.9 5.6 16 8 16c4.4 0 8-3.6 8-8s-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RotateRight(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 7V3l-1.1 1.1C13.6 1.6 11 0 8 0C3.6 0 0 3.6 0 8s3.6 8 8 8c2.4 0 4.6-1.1 6-2.8l-1.5-1.3C11.4 13.2 9.8 14 8 14c-3.3 0-6-2.7-6-6s2.7-6 6-6c2.4 0 4.5 1.5 5.5 3.5L12 7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rss(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.4 13.8a2.2 2.2 0 1 1-4.4 0a2.2 2.2 0 0 1 4.4 0"/><path fill="currentColor" d="M10.6 16H7.5c0-4.1-3.4-7.5-7.5-7.5V5.4c5.9 0 10.6 4.7 10.6 10.6"/><path fill="currentColor" d="M12.8 16C12.8 8.9 7.1 3.2 0 3.2V0c8.8 0 16 7.2 16 16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RssSquare(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v16h16V0zm3.6 14c-.9 0-1.6-.7-1.6-1.6s.7-1.6 1.6-1.6s1.6.7 1.6 1.6S4.6 14 3.6 14m4 0c0-3.1-2.5-5.6-5.6-5.6V6c4.4 0 8 3.6 8 8zm4 0c0-5.3-4.3-9.6-9.6-9.6V2c6.6 0 12 5.4 12 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Safe(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 0v16h3v-1h8v1h3V0zm13 10h-1V5h1zm0-7h-1V2H3v11h10v-1h1v2H2V1h12zM8.5 7.5c0 1.1-.9 2-2 2s-2-.9-2-2s.9-2 2-2s2 .9 2 2"/><path fill="currentColor" d="M7.5 7.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SafeLock(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0m3.13 14.25l-.37-.9l-.92.38l.37.9c-.659.23-1.419.363-2.21.363s-1.551-.133-2.259-.378l.419-.885l-.92-.38l-.37.9a7.061 7.061 0 0 1-3.102-3.08l.882-.41l-.38-.93l-.9.37c-.23-.659-.363-1.419-.363-2.21s.133-1.551.378-2.259l.885.419l.38-.92l-.9-.37a7.054 7.054 0 0 1 3.08-3.092l.41.882l.92-.38l-.37-.9c.659-.23 1.419-.363 2.21-.363s1.551.133 2.259.378l-.419.885l.92.38l.37-.9a7.061 7.061 0 0 1 3.102 3.08l-.882.41l.38.92l.9-.37c.23.659.363 1.419.363 2.21s-.133 1.551-.378 2.259l-.885-.419l-.38.92l.9.37a7.061 7.061 0 0 1-3.08 3.102z"/><path fill="currentColor" d="M10.36 3.62L9.2 6.41A1.986 1.986 0 0 0 8.001 6h.279l1.15-2.77A4.836 4.836 0 0 0 8.003 3h-.071a5.06 5.06 0 1 0 0 10.12a5.06 5.06 0 0 0 2.454-9.486z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Scale(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.81 10l-2.5-5H14a.5.5 0 0 0 0-1h-.79a6.04 6.04 0 0 0-4.198-1.95L9 2a1 1 0 0 0-2 0v.05a6.168 6.168 0 0 0-4.247 1.947L2 4a.5.5 0 0 0 0 1h.69l-2.5 5H0c0 1.1 1.34 2 3 2s3-.9 3-2h-.19L3.26 4.91a.525.525 0 0 0 .159-.148A4.842 4.842 0 0 1 6.994 3.06L7 14H6v1H4v1h8v-1h-2v-1H9V3.06a4.71 4.71 0 0 1 3.524 1.693a.519.519 0 0 0 .193.186L10.19 10H10c0 1.1 1.34 2 3 2s3-.9 3-2zM5 10H1l2-3.94zm6 0l2-3.94L15 10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScaleUnbalance(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.81 9l-2.47-4.93l.83-.15a.509.509 0 1 0-.183-.999l-.777.14a5.96 5.96 0 0 0-4.236-1.178a.958.958 0 0 0-.967-.882h-.008a1 1 0 0 0-1 1v.2a6.332 6.332 0 0 0-4.066 2.697l-.754.153a.503.503 0 1 0 .092 1h.088l.35-.05l-2.52 5h-.19c0 1.1 1.34 2 3 2s3-.9 3-2h-.19l-2.56-5.12h.1a.512.512 0 0 0 .379-.297c.021-.093.701-1.583 3.271-2.363v10.78h-1v1h-2v1h8v-1h-2v-1h-1V2.881a4.617 4.617 0 0 1 3.686 1.065l-.006-.005l-2.49 5.06h-.19c0 1.1 1.34 2 3 2s3-.9 3-2zM5 11H1l2-3.94zm6-2l2-3.94L15 9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScatterChart(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 15V0H0v16h16v-1z"/><path fill="currentColor" d="M5 11a1 1 0 1 1-2 0a1 1 0 0 1 2 0m3-5a1 1 0 1 1-2 0a1 1 0 0 1 2 0m6-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-3 5a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Scissors(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3.1s-2.1-1.1-3.5-1c-.3 0-.5.1-.7.2L7.5 5.7L5.7 4.2c.1-.3.2-.6.3-1C6.1 1.4 4.6-.2 2.7 0C1.5.1.4 1 .1 2.2c-.3 1.3.2 2.5 1.2 3.2L4.6 8l-3.3 2.6c-1 .7-1.5 1.9-1.2 3.2c.3 1.2 1.4 2 2.6 2.2c1.9.2 3.4-1.4 3.2-3.2c0-.3-.1-.7-.3-1l1.8-1.5l4.3 3.4c.2.1.4.2.7.2c1.4.1 3.5-1 3.5-1L10.2 8zM2.8 4.6c-.9-.1-1.6-.9-1.5-1.8s.9-1.6 1.8-1.5c.9.1 1.6.9 1.5 1.8c0 .9-.9 1.6-1.8 1.5m.3 10.1c-.9.1-1.7-.6-1.8-1.5s.6-1.7 1.5-1.8c.9-.1 1.7.6 1.8 1.5s-.6 1.7-1.5 1.8m9.3-11.5h.2c.4 0 .9.1 1.4.2L7.2 9.1L6.3 8zm1.6 9.4c-.5.2-1 .3-1.4.2h-.2l-4-3.2l1-.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Screwdriver(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8 10.8l.9-.8l-.9-.9l5.7-5.7l1.2-.4L16 .8l-.7-.7l-2.3 1l-.5 1.2L6.9 8L6 7.1l-.8.9s.8.6-.1 1.5c-.5.5-1.3-.1-2.8 1.4L.2 13s-.6 1 .6 2.2s2.2.6 2.2.6l2.1-2.1c1.4-1.4.9-2.3 1.3-2.7c.9-.9 1.6-.2 1.6-.2m-3.1-.4l.7.7l-3.8 3.8l-.7-.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Search(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.7 14.3l-4.2-4.2c-.2-.2-.5-.3-.8-.3c.8-1 1.3-2.4 1.3-3.8c0-3.3-2.7-6-6-6S0 2.7 0 6s2.7 6 6 6c1.4 0 2.8-.5 3.8-1.4c0 .3 0 .6.3.8l4.2 4.2c.2.2.5.3.7.3s.5-.1.7-.3c.4-.3.4-.9 0-1.3M6 10.5c-2.5 0-4.5-2-4.5-4.5s2-4.5 4.5-4.5s4.5 2 4.5 4.5s-2 4.5-4.5 4.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchMinus(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.7 14.3l-4.2-4.2c-.2-.2-.5-.3-.8-.3c.8-1 1.3-2.4 1.3-3.8c0-3.3-2.7-6-6-6S0 2.7 0 6s2.7 6 6 6c1.4 0 2.8-.5 3.8-1.4c0 .3 0 .6.3.8l4.2 4.2c.2.2.5.3.7.3s.5-.1.7-.3c.4-.3.4-.9 0-1.3M6 10.5c-2.5 0-4.5-2-4.5-4.5s2-4.5 4.5-4.5s4.5 2 4.5 4.5s-2 4.5-4.5 4.5"/><path fill="currentColor" d="M3 5h6v2H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchPlus(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.7 14.3l-4.2-4.2c-.2-.2-.5-.3-.8-.3c.8-1 1.3-2.4 1.3-3.8c0-3.3-2.7-6-6-6S0 2.7 0 6s2.7 6 6 6c1.4 0 2.8-.5 3.8-1.4c0 .3 0 .6.3.8l4.2 4.2c.2.2.5.3.7.3s.5-.1.7-.3c.4-.3.4-.9 0-1.3M6 10.5c-2.5 0-4.5-2-4.5-4.5s2-4.5 4.5-4.5s4.5 2 4.5 4.5s-2 4.5-4.5 4.5"/><path fill="currentColor" d="M7 3H5v2H3v2h2v2h2V7h2V5H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Select(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 4H1c-.6 0-1 .4-1 1v6c0 .6.4 1 1 1h14c.6 0 1-.4 1-1V5c0-.6-.4-1-1-1m-3 5l-2-2h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Server(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 5v3h10V5zm4 2H4V6h3zM3 4h10l-2-4H5zm0 8h10V9H3zm8-2h1v1h-1zm-2 0h1v1H9zm-6 6h10v-3H3zm1-2h3v1H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Share(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 3H4.9S0 3 0 8c0 3.9 3 8 3 8S1.3 9 4.8 9H10v3l6-6l-6-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareSquare(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 3H7.4S3 2.8 3 7.3C3 10.8 5 14 5 14s-.4-7 2.3-7H11v3l5-5l-5-5z"/><path fill="currentColor" d="M14 9v6H1V2h9V1H0v15h15V8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shield(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 0v7c0 5.6 7 9 7 9s7-3.4 7-9V0zm13 7c0 4.2-4.6 7.1-6 7.9V1h6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shift(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 2v12h16V2zm6 6v3H4V8H2l3-3l3 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShiftArrow(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 2L1 9h4v5h6V9h4zm2 6v5H6V8H3.5L8 3.42L12.5 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shop(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 15h16v1H0zM0 0v6c.005.732.401 1.37.991 1.715L1 14h9V9h3v5h2V7.72c.599-.35.995-.988 1-1.719V0zm4 2h2v4a1 1 0 0 1-2 0zM2 7a1 1 0 0 1-1-1V2h2v4a1 1 0 0 1-1 1m6 5H3V9h5zm1-6a1 1 0 0 1-2 0V2h2zm3 0a1 1 0 0 1-2 0V2h2zm3 0a1 1 0 0 1-2 0V2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignIn(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 1v2l1 1V2h7v12H8v-2l-1 1v2h9V1z"/><path fill="currentColor" d="M10 8L5 4v2H0v4h5v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignInAlt(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h2v16H0zm3 10h8v3l5-5l-5-5v3H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignOut(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 4V1H0v14h9v-3H8v2H1V2h7v2z"/><path fill="currentColor" d="m16 8l-5-4v2H6v4h5v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignOutAlt(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 0h2v16h-2zM8 6H0v4h8v3l5-5l-5-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Signal(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.9 13.2L8 14.3l1.1-1.1c-.3-.3-.7-.5-1.1-.5s-.9.2-1.1.5M8 4.6c2.7 0 5.1 1.1 6.9 2.8L16 6.3C14 4.3 11.1 3 8 3S2 4.3 0 6.3l1.1 1.1C2.9 5.7 5.3 4.6 8 4.6"/><path fill="currentColor" d="m2.3 8.6l1.1 1.1C4.6 8.6 6.2 7.9 8 7.9s3.4.7 4.6 1.9l1.1-1.1C12.3 7.1 10.2 6.2 8 6.2s-4.3.9-5.7 2.4"/><path fill="currentColor" d="M4.6 10.9L5.7 12c.6-.6 1.4-.9 2.3-.9s1.7.4 2.3.9l1.1-1.1C10.6 10 9.3 9.5 8 9.5s-2.6.5-3.4 1.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sitemap(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 12V7.5h-6V4H10V0H6v4h1.5v3.5h-6V12H0v4h4v-4H2.5V8.5h5V12H6v4h4v-4H8.5V8.5h5V12H12v4h4v-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Slider(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 6h-3.6c-.7-1.2-2-2-3.4-2s-2.8.8-3.4 2H0v4h5.6c.7 1.2 2 2 3.4 2s2.8-.8 3.4-2H16zM1 9V7h4.1c0 .3-.1.7-.1 1s.1.7.1 1zm8 2c-1.7 0-3-1.3-3-3s1.3-3 3-3s3 1.3 3 3s-1.3 3-3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sliders(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 0h2v3H7zM6 4v3h1v9h2V7h1V4zM2 0h2v8H2zM1 9v3h1v4h2v-4h1V9zm11-9h2v10h-2zm-1 11v3h1v2h2v-2h1v-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmileyO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1c3.9 0 7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7s3.1-7 7-7m0-1C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8"/><path fill="currentColor" d="M8 13.2c-2 0-3.8-1.2-4.6-3.1l.9-.4c.6 1.5 2.1 2.4 3.7 2.4s3.1-1 3.7-2.4l.9.4c-.8 2-2.6 3.1-4.6 3.1M7 6a1 1 0 1 1-2 0a1 1 0 0 1 2 0m4 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sort(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 7H5l3-4zM5 9h6l-3 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SoundDisable(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 5H0v6h4l5 4V1zm11.9.6l-.8-.7l-2.3 2.4l-2.4-2.4l-.8.7L12 8l-2.4 2.4l.8.7l2.4-2.4l2.3 2.4l.8-.7L13.5 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SparkLine(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 6a2 2 0 0 0-2 2v.16l-.81.25l-2.3-3.48l-1.73 4.32L6 3.44L3.7 8.22L2.06 6.91L0 8v1.08l1.94-1l2.11 1.7l1.56-3.22l1.23 6.19l2.27-5.68l1.68 2.52l1.55-.48A2 2 0 1 0 14.004 6zm0 3a1 1 0 1 1 0-2a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Specialist(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.1 8c.2.6.3 1.1.3 1.1c.8 1.3 1.8 1.1 1.8 1.8c0 .3-.2.6-.5.7L8 13.4l2.3-1.7c-.3-.2-.5-.4-.5-.7c0-.8 1-.5 1.8-1.8c0 0 .2-.4.3-1.1c.3-1.1.6-3.1.5-4.1h-1.5c0-.3.1-.6.1-1h1.1c-.3-1.4-1-2-2.2-2.3C9.4.3 8.7 0 8 0S6.6.3 6.1.7C4.9 1 4.3 1.6 3.9 3H5c0 .4.1.7.2 1H3.6c-.1 1 .2 3 .5 4m7.1.5l-.3.3l-.5.6c-.4.5-.8.8-1.4.9l-.4.1c-.4.1-.9.1-1.4 0l-.4-.1c-.6-.2-1.1-.5-1.5-1.1l-.2-.4l-.3-.3l-.7-.5l3.1-.9c.5-.1 1-.2 1.5 0l3.2.9zM6 3c0-1.1.9-2 2-2s2 .9 2 2s-.9 2-2 2s-2-.9-2-2"/><path fill="currentColor" d="M15.5 14.2c-1.3-2.4-2.6-2-3.9-2.2h-.1L8 14.6L4.5 12h-.1c-1.4.1-2.6-.2-3.9 2.2c-.2.4-.4 1.1-.5 1.8h16c-.1-.7-.3-1.4-.5-1.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spinner(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9.9.2l-.2 1C12.7 2 15 4.7 15 8c0 3.9-3.1 7-7 7s-7-3.1-7-7c0-3.3 2.3-6 5.3-6.8l-.2-1C2.6 1.1 0 4.3 0 8c0 4.4 3.6 8 8 8s8-3.6 8-8c0-3.7-2.6-6.9-6.1-7.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpinnerArc(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 8c0 3.9-3.1 7-7 7s-7-3-7-7H0c0 4 3.6 8 8 8s8-3.6 8-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpinnerThird(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.9 3.1C14.2 4.3 15 6.1 15 8c0 3.9-3.1 7-7 7s-7-3.1-7-7c0-1.9.8-3.7 2.1-4.9l-.8-.8C.9 3.8 0 5.8 0 8c0 4.4 3.6 8 8 8s8-3.6 8-8c0-2.2-.9-4.2-2.3-5.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SplineAreaChart(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 15V0H0v16h16v-1z"/><path fill="currentColor" d="M10 7C8 7 7.92 6 6 6C3.66 6 2 9 2 9v5h14V2c-2 0-3.86 5-6 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SplineChart(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 15V0H0v16h16v-1z"/><path fill="currentColor" d="M12 5c-.69 1-1.41 2-2 2l-.087.001c-.601 0-1.164-.16-1.65-.44a4.519 4.519 0 0 0-2.2-.562h-.067a5.83 5.83 0 0 0-3.991 1.993l-.006 2.347c.77-1.12 2.32-2.84 4-2.84h.048c.579 0 1.121.156 1.587.428a4.682 4.682 0 0 0 2.264.573l.106-.001c1.395 0 2.335-1.32 3.245-2.6s1.75-2.4 2.75-2.4v-1.5c-1.81 0-3 1.61-4 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Split(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 11h6v5H0zm11-1V8l-.64.64a4.427 4.427 0 0 1-1.358-3.658L11.001 5V0h-6v5h2a4.43 4.43 0 0 1-1.358 3.639l-.642-.638v2h2l-.65-.65A5.426 5.426 0 0 0 8 4.982c-.01.149-.016.299-.016.45c0 1.539.639 2.928 1.665 3.917l-.648.652h2zm-1 1h6v5h-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SplitH(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 1v14h16V1zm1 3h6.5v10H1zm14 10H8.5V4H15zm0-11h-1V2h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SplitV(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 1v14h16V1zm14 1h1v1h-1zm1 2v4.5H1V4zM1 14V9.5h14V14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spoon(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5 4.8c0-1.8-.9-4.8-3-4.8s-3 3-3 4.8c0 1.5.8 2.8 2.2 3.1c-.5 1.6-.7 4.6-.7 4.6v2c0 .8.7 1.5 1.5 1.5S9 15.3 9 14.5v-2c0-.6-.2-3.2-.7-4.6c1.4-.3 2.2-1.6 2.2-3.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareShadow(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 2V0H0v14h2v2h14V2zm-1 11H1V1h12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Star(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.9 15.4L8 12.8l-4.9 2.6L4 10L0 6.1l5.5-.8l2.4-5l2.4 5l5.5.8L12 10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarHalfLeft(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.6 5.4l-5.5.8L4 10l-.9 5.5L8 12.9V.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarHalfLeftO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.9 6.2l-5.5-.8L8 .4l-2.4 5l-5.5.8L4 10l-.9 5.4L8 12.9l4.9 2.6L12 10zM8 11.8V2.7l1.8 3.6l4 .6l-2.9 2.8l.7 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarHalfRight(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.5 5.4l5.5.8l-4 3.8l.9 5.5L8 12.9V.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarHalfRightO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.9 6.2l-5.5-.8L8 .4l-2.4 5l-5.5.8L4 10l-.9 5.4L8 12.9l4.9 2.6L12 10zM4.4 13.7l.7-4l-2.9-2.8l4-.6L8 2.7v9.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.9 6.2l-5.5-.8L8 .4l-2.4 5l-5.5.8L4 10l-.9 5.4L8 12.9l4.9 2.6L12 10zM8 11.8l-3.6 1.9l.7-4l-2.9-2.8l4-.6L8 2.7l1.8 3.6l4 .6l-2.9 2.8l.7 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StartCog(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 0v6h1.7l.2.7l.2.6h.1l1.2-.6l1.8 1.8l-.6 1.2v.1l.6.2l.7.2v.2L16 7zm.5 10.5c-.2 0-.4.1-.5.2c-.3.2-.5.5-.5.8s.2.7.5.8c.1.1.3.2.5.2c.6 0 1-.4 1-1s-.4-1-1-1"/><path fill="currentColor" d="M9 12v-1l-1.1-.4c-.1-.3-.2-.6-.4-.9l.5-1l-.7-.7l-1 .5c-.3-.2-.6-.3-.9-.4L5 7H4l-.4 1.1c-.3.1-.6.2-.9.4l-1-.5l-.7.7l.5 1.1c-.2.3-.3.6-.4.9L0 11v1l1.1.4c.1.3.2.6.4.9l-.5 1l.7.7l1.1-.5c.3.2.6.3.9.4L4 16h1l.4-1.1c.3-.1.6-.2.9-.4l1 .5l.7-.7l-.5-1.1c.2-.3.3-.6.4-.9zm-4.5 1.5c-1.1 0-2-.9-2-2s.9-2 2-2s2 .9 2 2s-.9 2-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StepBackward(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 15V1L4 8zM2 1h2v14H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StepForward(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 1v14l10-7zm10 0h2v14h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stethoscope(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.7 15.2c.3.3 1 .8 1.8.8c2.7 0 3.3-2 3.4-3.6c.2-2.3.8-2.2 1.1-2.2c.7 0 .9.4.9 1.1c-.6.4-1 1-1 1.7c0 1.1.9 2 2 2s2-.9 2-2s-.9-2-2-2h-.2c-.2-.9-.7-1.8-1.8-1.8c-1.6 0-2 1.4-2.1 2.9C9.7 14.2 9 15 7.5 15c-.4 0-.8-.2-1-.4c-.6-.5-.5-2.3-.5-2.3c2 0 4-1.8 4.7-4.8l-.2-.1c.3-1.2.5-2.6.5-3.6c0-1.1-.3-1.9-1-2.5S8.5.5 7.9.5C7.7.2 7.4 0 7 0c-.5 0-1 .4-1 1s.4 1 1 1c.4 0 .7-.2.8-.5c.5 0 1 .2 1.5.6s.7.9.7 1.7c0 .9-.2 2.2-.5 3.5l-.2-.1C9 8.3 8 10.8 6 10.8H5c-2 0-3-2.5-3.3-3.6l-.2.1C1.2 6 1 4.7 1 3.8c0-.8.2-1.3.7-1.7c.4-.4 1-.5 1.5-.6c.1.3.4.5.8.5c.6 0 1-.4 1-1s-.4-1-1-1c-.4 0-.7.2-.9.5c-.6 0-1.4.2-2.1.8S0 2.7 0 3.8c0 1 .2 2.4.5 3.7l-.2.1C1 10.5 3 12.3 5 12.3c0 0-.1 2.2.7 2.9M14 14c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1s-.5 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stock(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 6V0H4v6H0v7h16V6zm-5 6H1V7h2v1h2V7h2zM5 6V1h2v1h2V1h2v5zm10 6H9V7h2v1h2V7h2zM0 16h3v-1h10v1h3v-2H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stop(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 1h14v14H1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopCog(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 0v7.2l.5-.5l1.2.6h.1l.2-.6l.3-.7h2.4l.2.7l.2.6h.1l1.2-.6l1.8 1.8l-.6 1.2v.1l.6.2l.7.2v2.4l-.7.2l-.6.2v.1l.6 1.2l-.4.7H16V0z"/><path fill="currentColor" d="M5.5 11.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/><path fill="currentColor" d="M7.9 12.4L9 12v-1l-1.1-.4c-.1-.3-.2-.6-.4-.9l.5-1l-.7-.7l-1 .5c-.3-.2-.6-.3-.9-.4L5 7H4l-.4 1.1c-.3.1-.6.2-.9.4l-1-.5l-.7.7l.5 1.1c-.2.3-.3.6-.4.9L0 11v1l1.1.4c.1.3.2.6.4.9l-.5 1l.7.7l1.1-.5c.3.2.6.3.9.4L4 16h1l.4-1.1c.3-.1.6-.2.9-.4l1 .5l.7-.7l-.5-1.1c.2-.2.3-.5.4-.8m-3.4 1.1c-1.1 0-2-.9-2-2s.9-2 2-2s2 .9 2 2s-.9 2-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stopwatch(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 8.14V4.5h-1v3.64a1.001 1.001 0 0 0 .5 1.866a1 1 0 0 0 .505-1.863z"/><path fill="currentColor" d="M8 2a7 7 0 1 0 0 14A7 7 0 0 0 8 2m0 12.5a5.5 5.5 0 1 1 0-11A5.5 5.5 0 0 1 13.5 9a5.51 5.51 0 0 1-5.499 5.5zM6 0h4v1.5H6z"/><path fill="currentColor" d="m.005 4.438l2.713-2.939L3.82 2.516L1.107 5.455zm12.181-1.919l1.102-1.017l2.713 2.939l-1.102 1.017z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Storage(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 4L7.94 0L0 4v1h1v11h2V7h10v9h2V5h1zM4 6V5h2v1zm3 0V5h2v1zm3 0V5h2v1z"/><path fill="currentColor" d="M6 9H5V8H4v3h3V8H6zm0 4H5v-1H4v3h3v-3H6zm4 0H9v-1H8v3h3v-3h-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Strikethrough(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5 7c-.5-.3-1-.5-1.4-.7c-2-.9-2.1-1.1-2-1.9s.4-1 .6-1.2c.9-.5 2.8-.1 3.5.2L12.3.6C11.9.4 8.6-.8 6.2.6c-.8.5-1.9 1.5-2.1 3.4c-.2 1.3.1 2.3.7 3H0v1h16V7zM7.7 9s.1 0 .1.1c2 .9 2.4 1.2 2.2 2.5c-.2.9-.5 1.1-.8 1.3c-1.1.6-3.3 0-4.4-.5L3.6 15c.3.1 2.3 1 4.5 1c.9 0 1.8-.2 2.6-.6c.9-.5 2-1.4 2.4-3.4c.2-1.3 0-2.3-.4-3.1h-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Subscript(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 15v1h-4v-1s3.3-1.6 2.6-3.2c-.5-1.1-2-.2-2-.2l-.5-.9s1.9-1.4 3.1-.2c2.4 2.3-1.4 4.5-1.4 4.5zM12 3H8.6L6 6L3.4 3H0l4.3 5L0 13h3.4L6 10l2.6 3H12L7.7 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Suitcase(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 3V1H5v2H0v12h16V3zM4 14H3V4h1zm6-11H6V2h4zm3 11h-1V4h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SunDown(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 3H9V1H7v2H6l2 3zm4 10l-1.58-1.18l.78-1.82l-2-.23L11 7.8l-1.82.78L8 7L6.82 8.58L5 7.8l-.23 2l-1.97.2l.78 1.82L2 13H0v1h16v-1zM4 13a4.143 4.143 0 0 1 3.993-4A4.143 4.143 0 0 1 12 12.993z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SunO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16 8l-2.2-1.6L14.9 4l-2.7-.2l-.2-2.7l-2.4 1.1L8 0L6.4 2.2L4 1.1l-.2 2.7l-2.7.2l1.1 2.4L0 8l2.2 1.6L1.1 12l2.7.2l.2 2.7l2.4-1.1L8 16l1.6-2.2l2.4 1.1l.2-2.7l2.7-.2l-1.1-2.4zm-8 5c-2.8 0-5-2.2-5-5s2.2-5 5-5s5 2.2 5 5s-2.2 5-5 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SunRise(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 4h1v2h2V4h1L8 1zm6.42 7.82L13.2 10l-2-.23L11 7.8l-1.82.78L8 7L6.82 8.58L5 7.8l-.23 2l-1.97.2l.78 1.82L2 13H0v1h16v-1h-2zM4 13a4.143 4.143 0 0 1 3.993-4A4.143 4.143 0 0 1 12 12.993z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Superscript(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 5v1h-4V5s3.3-1.6 2.6-3.2c-.5-1.1-2-.2-2-.2l-.5-.9S14-.7 15.2.5C17.6 2.8 13.8 5 13.8 5zm-4-2H8.6L6 6L3.4 3H0l4.3 5L0 13h3.4L6 10l2.6 3H12L7.7 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sword(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.8.5l-.1-.2l-.2-.1c-.1 0-2.5-.8-4.2.9L4.6 7.7c-.9-.6-1.7-1.2-1.8-1l-.4.3c-.2.2.9 1.7 1.8 2.7l-2.5 3.4c-.3-.3-.8-.3-1.1 0l-.3.3c-.3.3-.3.8 0 1.1l1 1c.3.3.8.3 1.1 0l.3-.3c.3-.3.3-.8 0-1.1l3.5-2.5c1 .9 2.5 2 2.7 1.8l.4-.4c.1-.1-.4-1-1.1-1.8l6.7-6.7c1.7-1.5.9-3.9.9-4m-8.1 10l-.8-.8l6.2-6.9L6.2 9l-.7-.7L12 1.8c1-1 2.3-.8 2.9-.7c.1.6.3 1.9-.7 2.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tab(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 2v12h16V2zm13 9h-1V8l-3 3V9H3V7h6V5l3 3V5h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TabA(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 10H0V6h9V4l5 4l-5 4zm5-6h2v8h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Table(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 1v15h16V1zm5 14H1v-2h4zm0-3H1v-2h4zm0-3H1V7h4zm0-3H1V4h4zm5 9H6v-2h4zm0-3H6v-2h4zm0-3H6V7h4zm0-3H6V4h4zm5 9h-4v-2h4zm0-3h-4v-2h4zm0-3h-4V7h4zm0-3h-4V4h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tablet(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 2v12h16V2zm13 11H2V3h11zm2-4h-1V7h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tabs(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 4V2H0v12h16V4zm-4-1h3v1h-3zM6 3h3v1H6zm9 10H1V3h4v2h10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tag(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1H1v7l7 7l7-7zM3.75 5a1.25 1.25 0 1 1 0-2.5a1.25 1.25 0 0 1 0 2.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tags(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 2H7.5l7 7l-5.3 5.2l.8.8l6-6z"/><path fill="currentColor" d="M6 2H0v6l7 7l6-6zM2.8 6c-.7 0-1.3-.6-1.3-1.2s.6-1.2 1.2-1.2S4 4.1 4 4.8S3.4 6 2.8 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tasks(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 0h10v4H6zm0 6h10v4H6zm0 6h10v4H6zM3 1v2H1V1zm1-1H0v4h4zM3 13v2H1v-2zm1-1H0v4h4zm1.3-6.1l-.6-.8l-.9.9H0v4h4V7.2zM2.7 7l-.7.7l-.8-.7zM1 8.2l.9.8H1zM3 9h-.9l.9-.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Taxi(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15 6.1l-1.4-2.9c-.4-.7-1.1-1.2-2-1.2H11V.7c0-.4-.3-.7-.7-.7H5.7c-.4 0-.7.3-.7.7V2h-.7c-.8 0-1.6.5-1.9 1.2L1 6.1c-.6.1-1 .6-1 1.1v3.5c0 .6 0 1.1 1 1.2v2c0 .6.4 1.1 1 1.1h.9c.6 0 1.1-.5 1.1-1.1V12h8v1.9c0 .6.4 1.1 1 1.1h.9c.6 0 1.1-.5 1.1-1.1v-2c1-.1 1-.6 1-1.2V7.2c0-.5-.4-1-1-1.1M4 8.4c0 .3-.3.6-.6.6H1.6c-.3 0-.6-.3-.6-.6v-.8c0-.3.3-.6.6-.6h1.8c.3 0 .6.3.6.6zm6 2.6H6v-1h4zM2.1 6l1.2-2.4c.2-.4.6-.6 1-.6h7.4c.4 0 .8.2 1 .6L13.9 6zM15 8.4c0 .3-.3.6-.6.6h-1.8c-.3 0-.6-.3-.6-.6v-.8c0-.3.3-.6.6-.6h1.8c.3 0 .6.3.6.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Teeth(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.6 7.6c-.1.1-.5.4-1.6.4c1.1 0 1.5.3 1.6.4c.2-.2.6-.4 1.5-.4c-.9 0-1.3-.2-1.5-.4"/><path fill="currentColor" d="M8 0C3.6 0 0 3.6 0 8s3.6 8 8 8s8-3.6 8-8s-3.6-8-8-8m5.1 11.6c-1 0-1.4-.8-1.6-1.6c-.2.9-.6 2-1.8 2c-1.1 0-1.5-1.1-1.7-2c-.2 1-.6 2-1.7 2s-1.6-1.1-1.8-2c-.2.8-.6 1.6-1.6 1.6c-2 0-1.9-3-1.9-3S1.2 8 2.7 8C1.2 8 1 7.5 1 7.5s-.1-3 1.9-3c1 0 1.4.8 1.6 1.6c.2-.9.6-2 1.8-2C7.4 4 7.8 5.1 8 6c.2-1 .6-2 1.7-2s1.6 1.1 1.8 2c.2-.8.6-1.6 1.6-1.6c2 0 1.9 3 1.9 3s-.3.6-1.8.6c-1.2 0-1.6-.3-1.8-.4c-.2.2-.7.4-1.6.4c-1.2 0-1.6-.2-1.8-.4c-.1.1-.6.4-1.6.4c1 0 1.4.3 1.6.4c.2-.2.6-.4 1.8-.4c1 0 1.4.2 1.7.4c0-.1.5-.4 1.7-.4c1.5 0 1.8.6 1.8.6s.1 3-1.9 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Terminal(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 12h9v1H6zm-4.9 1h1.2L6 8L2.3 3H1l3.8 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextHeight(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 3h1l-1.5-3L13 3h1v10h-1l1.5 3l1.5-3h-1zM1 0v3h4v13h3V3h4V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextInput(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 2h1v4H2z"/><path fill="currentColor" d="M1 0C.4 0 0 .4 0 1v14c0 .6.4 1 1 1h15V0zm12 15H1V1h12zm2 0h-1v-1h1zm0-2h-1V3h1zm0-11h-1V1h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextLabel(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 4.9c-1.4 0-2.5.8-2.6.9l1.2 1.6s.7-.5 1.4-.5c1.4 0 1.5 1.2 1.5 1.6c-.4-.1-1.1-.3-2-.1c-1.4.3-2.8 2-2.1 3.9c.7 1.8 3.1 2.1 4.1.6v1h2V8.6c0-2.7-1.9-3.7-3.5-3.7m-1 6.5C11.4 9.5 13 9.5 14 9.6v1c0 1.2-2.3 2.3-2.5.8M6.9 14H9L5.8 2H3.1L0 14h2.1l1-4h2.7zM3.6 8l.8-3.2l.9 3.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextWidth(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 14.5L12 13v1H3v-1l-3 1.5L3 16v-1h9v1zM0 0v3h6v9h3V3h6V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThinSquare(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 1H1v14h14zm-1 13H2V2h12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsDown(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.6 7.8s.5.5.4 1.6c0 1.5-1.6 1.6-1.6 1.6H12c-.2 0-.3.2-.3.4c.3.7.8 2.1.6 3.1c-.3 1.4-1.5 1.5-1.9 1.5c-.1 0-.2-.1-.2-.2l-1-2.8s0-.1-.1-.1l-2.6-2.8c-.1-.1-.2-.1-.3-.1H6V3h.2c.7 0 3.2-2 5.4-2s2.7.3 3.1 1c.4.7.1 1.3.1 1.3s.5.3.6 1c.1.7-.1 1.1-.1 1.1s.5.4.5 1.2c.1.9-.2 1.2-.2 1.2M0 11h5V3H0zm2.5-3.5c.6 0 1 .4 1 1s-.4 1-1 1s-1-.4-1-1s.4-1 1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsDownO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.6 7.3c.1-.3.3-.7.2-1.2c0-.6-.3-1.1-.5-1.3c.1-.3.1-.6 0-1.1s-.4-.8-.6-1c.1-.3.1-.8-.3-1.4C14 .3 13.2 0 10.8 0C9.1 0 7.5.8 6.2 1.5c-.4.2-1 .5-1.2.5H0v9h5v-.9l2.7 2.7l1 2.8c.2.2.4.4.8.4h.1c.5 0 2-.1 2.4-1.9c.2-.9-.1-2.2-.5-3.1h2.3c.7-.1 2.1-.6 2.2-2.1c0-.7-.2-1.3-.4-1.6m-13.1.2c.6 0 1 .4 1 1s-.4 1-1 1s-1-.4-1-1s.4-1 1-1M13.8 10h-2.5c-.3 0-.5.1-.7.4c-.2.2-.2.5-.1.8c.5 1.2.7 2.2.6 2.8c-.2.9-.9 1.1-1.4 1.1l-1-2.7c0-.1-.1-.2-.2-.3L5.6 9.2c-.1-.1-.3-.2-.5-.2H5V3c.4 0 .8-.2 1.7-.6C7.8 1.8 9.4 1 10.8 1c2.5 0 2.7.4 2.9.7c.3.5.1.9.1.9l-.2.4l.4.3s.4.2.5.7c.1.4 0 .7 0 .7l-.3.3l.3.3s.4.3.4.9c0 .5-.2.7-.2.7l-.4.3l.4.4s.4.4.3 1.2c0 1.1-1.1 1.2-1.2 1.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsUp(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.6 8.2s.5-.5.4-1.6C16 5.1 14.4 5 14.4 5H12c-.2 0-.3-.2-.3-.4c.3-.7.8-2.1.6-3.1C12 .1 10.8 0 10.4 0c-.1 0-.2.1-.2.2L9.2 3s0 .1-.1.1L6.5 5.9c-.1.1-.2.1-.3.1H6v7h.2c.7 0 3.2 2 5.4 2s2.7-.3 3.1-1c.4-.7.1-1.3.1-1.3s.5-.3.6-1c.1-.7-.1-1.1-.1-1.1s.5-.4.5-1.2c.1-.9-.2-1.2-.2-1.2M0 14h5V6H0zm2.5-3.5c.6 0 1 .4 1 1s-.4 1-1 1s-1-.4-1-1s.4-1 1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsUpO(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 7.1C16 5.6 14.6 5 13.8 5h-2.2c.4-1 .7-2.2.5-3.1C11.6.1 10.1 0 9.6 0h-.1c-.4 0-.6.2-.8.5l-1 2.8L5 6H0v9h5v-1c.2 0 .7.3 1.2.6c1.2.6 2.9 1.5 4.5 1.5c2.4 0 3.2-.3 3.8-1.3c.3-.6.3-1.1.3-1.4c.2-.2.5-.5.6-1s.1-.8 0-1.1c.2-.3.4-.7.5-1.3c0-.5-.1-.9-.2-1.2c.1-.4.3-.9.3-1.7M2.5 13.5c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1s-.4 1-1 1m12.2-4.4s.2.2.2.7c0 .6-.4.9-.4.9l-.3.3l.2.3s.2.3 0 .7c-.1.4-.5.7-.5.7l-.3.3l.2.4s.2.4-.1.9c-.2.4-.4.7-2.9.7c-1.4 0-3-.8-4.1-1.4c-.8-.4-1.3-.6-1.7-.6V7h.1c.2 0 .4-.1.6-.2L8.5 4c.1-.1.1-.2.2-.3l1-2.7c.5 0 1.2.2 1.4 1.1c.1.6-.1 1.6-.6 2.8c-.1.3-.1.5.1.8c.1.2.4.3.7.3h2.5c.1 0 1.2.2 1.2 1.1c0 .8-.3 1.2-.3 1.2l-.3.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ticket(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 3H2c0 1.1-.9 2-2 2v6c1.1 0 2 .9 2 2h12c0-1.1.9-2 2-2V5c-1.1 0-2-.9-2-2m-1 9H3V4h10z"/><path fill="currentColor" d="M4 5h8v6H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimeBackward(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 4H7v5h4V8H8z"/><path fill="currentColor" d="M8 0C5 0 2.4 1.6 1.1 4.1L0 3v4h4L2.5 5.5C3.5 3.5 5.6 2 8 2c3.3 0 6 2.7 6 6s-2.7 6-6 6c-1.8 0-3.4-.8-4.5-2.1L2 13.2C3.4 14.9 5.6 16 8 16c4.4 0 8-3.6 8-8s-3.6-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimeForward(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 4H7v5h4V8H8z"/><path fill="currentColor" d="M16 7V3l-1.1 1.1C13.6 1.6 11 0 8 0C3.6 0 0 3.6 0 8s3.6 8 8 8c2.4 0 4.6-1.1 6-2.8l-1.5-1.3C11.4 13.2 9.8 14 8 14c-3.3 0-6-2.7-6-6s2.7-6 6-6c2.4 0 4.5 1.5 5.5 3.5L12 7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Timer(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.06 9.06c.271-.271.439-.646.439-1.06s-.168-.789-.439-1.06c-.59-.59-6.72-4.6-6.72-4.6s4 6.13 4.59 6.72a1.497 1.497 0 0 0 2.13 0"/><path fill="currentColor" d="M8 0v3h1V1.59c3.153.495 5.536 3.192 5.536 6.445a6.52 6.52 0 1 1-12.07-3.423L1.55 3.29A7.94 7.94 0 0 0 .017 8a8 8 0 1 0 8-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Toolbox(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 8h6v2h4V8h6v6H0z"/><path fill="currentColor" d="M7 7h2v2H7zm4-3V2H5v2H0v3h6V6h4v1h6V4zM6 4V3h4v1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tools(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.3 8.2l-.9.9l.9.9l-1.2 1.2l4.3 4.3c.6.6 1.5.6 2.1 0s.6-1.5 0-2.1zm3.9 6.8c-.4 0-.8-.3-.8-.8c0-.4.3-.8.8-.8s.8.3.8.8s-.3.8-.8.8M3.6 8l.9-.6L6 5.7l.9.9l.9-.9l-.1-.1c.2-.5.3-1 .3-1.6c0-2.2-1.8-4-4-4c-.6 0-1.1.1-1.6.3l2.9 2.9l-2.1 2.1L.3 2.4C.1 2.9 0 3.4 0 4c0 2.1 1.6 3.7 3.6 4"/><path fill="currentColor" d="m8 10.8l.9-.8l-.9-.9l5.7-5.7l1.2-.4L16 .8l-.7-.7l-2.3 1l-.5 1.2L6.9 8L6 7.1l-.8.9s.8.6-.1 1.5c-.5.5-1.3-.1-2.8 1.4L.2 13s-.6 1 .6 2.2s2.2.6 2.2.6l2.1-2.1c1.4-1.4.9-2.3 1.3-2.7c.9-.9 1.6-.2 1.6-.2m-3.1-.4l.7.7l-3.8 3.8l-.7-.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tooth(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.3 16c-1.2 0-1.7-3.9-1.7-4.1c-.1-1.3-1-2.1-1.6-2.2c-.6 0-1.4.9-1.6 2.2c0 .2-.5 4.1-1.7 4.1s-1.8-4.4-1.9-4.4c-.2-1.4.1-3.4.2-4c-.4-1.2-1.8-5.6-.5-7C3 .2 3.6 0 4.4 0c.6 0 1.3.1 2 .3c.6.1 1.1.2 1.6.2S9 .4 9.6.3c.7-.2 1.4-.3 2-.3c.8 0 1.4.2 1.8.7c1.3 1.4-.1 5.8-.5 7c.1.5.4 2.5.2 3.9c.1 0-.5 4.4-1.8 4.4M8 8.7c1.3.1 2.4 1.4 2.6 3.1c.1 1.2.5 2.4.7 2.9c.3-.6.7-2.1.9-3.3c.2-1.4-.2-3.7-.2-3.7v-.2c.7-2.1 1.4-5.3.8-6.1c-.3-.3-.7-.4-1.2-.4s-1.2.1-1.8.3c-.6.1-1.2.2-1.8.2s-1.2-.1-1.8-.2C5.6 1.1 4.9 1 4.4 1s-.9.1-1.1.4c-.7.7 0 4 .8 6.1v.2s-.4 2.3-.2 3.7c.2 1.2.6 2.7.9 3.3c.2-.6.6-1.7.7-2.9c.1-1.6 1.2-3 2.5-3.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Touch(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.62 6a1.312 1.312 0 0 0-.629.002a1.54 1.54 0 0 0-.402-.843a1.417 1.417 0 0 0-1.008-.42l-.053.001h-.007c-.22 0-.43.044-.621.124a1.2 1.2 0 0 0-.29-.464a1.453 1.453 0 0 0-1.115-.399H8.5a2.5 2.5 0 1 0-3.506.486L5 7.151a2.194 2.194 0 0 0-1.432.581C3 8.351 3 9.311 3 10.511v.72a3.617 3.617 0 0 0 1.402 2.764l.358.356c1.24 1.27 2.38 1.65 5.02 1.65c2.88 0 4.22-1.61 4.22-5.06v-2.51c0-.77-.22-2.12-1.38-2.43zM13 8.35v2.59C13 14.31 11.71 15 9.78 15c-2.6 0-3.4-.39-4.3-1.33l-.36-.37A2.652 2.652 0 0 1 4 11.235V10.5a3.347 3.347 0 0 1 .298-2.09c.186-.132.431-.228.698-.24l.003.7v-.22l-.34 1.5a.186.186 0 1 0 .34.151l1-1.211a.214.214 0 0 0 0-.091V3.39l-.001-.039c0-.256.083-.492.223-.684a.51.51 0 0 1 .399-.157h-.001c.21 0 .38.17.38.38v3.88a.5.5 0 0 0 1 0V5.45A.47.47 0 0 1 8.501 5A.42.42 0 0 1 9 5.412l-.001.029V6.77a.5.5 0 0 0 1 0v-.64a.45.45 0 0 1 .502-.379a.431.431 0 0 1 .338.11c.1.129.16.294.16.473v.836a.5.5 0 0 0 .504.494a.51.51 0 0 0 .496-.39a.513.513 0 0 1 .16-.273a.27.27 0 0 1 .223 0A1.38 1.38 0 0 1 13 8.327l.001-.007z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Train(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 11.2V3.8c0-1-.8-1.8-1.8-1.8H9V1h2V0H5v1h2v1H4.8C3.8 2 3 2.8 3 3.8v7.4c0 1 .8 1.8 1.8 1.8H5l-.7 1H3v1h.7L3 16h2l.6-1h4.9l.6 1h2l-.7-1h.6v-1h-1.3l-.7-1h.2c1 0 1.8-.8 1.8-1.8M4 3.9c0-.5.4-.9.9-.9H11c.6 0 1 .4 1 .9V6c0 .6-.4 1-.9 1H4.9c-.5 0-.9-.4-.9-.9zM4 11c0-.6.4-1 1-1s1 .4 1 1s-.4 1-1 1s-1-.4-1-1m5.9 3H6.1l.6-1h2.6zm.1-3c0-.6.4-1 1-1s1 .4 1 1s-.4 1-1 1s-1-.4-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trash(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 3s0-.51-2-.8v-.7A1.53 1.53 0 0 0 9.47 0h-3A1.5 1.5 0 0 0 5 1.5v.7a3.706 3.706 0 0 0-2.007.806L2 3v1h12V3zM6 1.5a.51.51 0 0 1 .499-.5h3.002a.53.53 0 0 1 .529.499v.561a10.224 10.224 0 0 0-1.527-.059c-.553 0-2.063 0-2.503.07zM2 5v1h1v9c1.234.631 2.692 1 4.236 1h1.529a9.418 9.418 0 0 0 4.289-1.025L13 6h1V5zm4 8.92q-.51-.06-1-.17V7h1zM9 14H7V7h2zm2-.28c-.267.07-.606.136-.95.184L10 7h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TreeTable(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 10V8H4V7h1V5H2v2h1v6h3v-2H4v-1z"/><path fill="currentColor" d="M0 0v16h16V0zm7 15H1V3h6zm4 0H8V3h3zm4 0h-3V3h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrendindDown(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 14h-4l1.29-1.29L9 8.41l-3 3l-6-6V2.59l6 6l3-3l5.71 5.7L15.99 10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrendingUp(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 2h-4l1.29 1.29L9 7.59l-3-3l-6 6v2.82l6-6l3 3l5.71-5.7L15.99 6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trophy(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.7 8c4.2-.3 4.3-2.7 4.3-5h-3V0H3v3H0c0 2.3.1 4.7 4.3 5c.9 1.4 2.1 2 2.7 2v4c-3 0-3 2-3 2h8s0-2-3-2v-4c.6 0 1.8-.6 2.7-2M13 4h2c-.1 1.6-.4 2.7-2.7 2.9c.3-.8.6-1.7.7-2.9M1 4h2c.1 1.2.4 2.1.7 2.9C1.5 6.7 1.1 5.6 1 4m3.5 2.1C4 4.4 4 3 4 3V1h1v2s0 1.7.4 3.1C5.9 7.8 7 9 7 9s-1.8-.2-2.5-2.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Truck(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 3h10v7H6zm9 11a2 2 0 1 1-3.999.001A2 2 0 0 1 15 14m-2-3c1.3 0 2.4.8 2.8 2h.2v-2z"/><path fill="currentColor" d="M5 5H1L0 9v4h1.2c.4-1.2 1.5-2 2.8-2s2.4.8 2.8 2h3.4c.4-1.2 1.5-2 2.8-2H5zM4 9H1l.8-3H4z"/><path fill="currentColor" d="M6 14a2 2 0 1 1-3.999.001A2 2 0 0 1 6 14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwinColSelect(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 2v12h16V2zm7 11H1V3h6zm8 0H9V3h6z"/><path fill="currentColor" d="M10 4h4v1h-4zM2 4h4v1H2zm0 2h4v1H2zm0 2h4v1H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Twitter(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 3c-.6.3-1.2.4-1.9.5c.7-.4 1.2-1 1.4-1.8c-.6.4-1.3.6-2.1.8c-.6-.6-1.5-1-2.4-1c-1.7 0-3.2 1.5-3.2 3.3c0 .3 0 .5.1.7c-2.7-.1-5.2-1.4-6.8-3.4c-.3.5-.4 1-.4 1.7c0 1.1.6 2.1 1.5 2.7c-.5 0-1-.2-1.5-.4C.7 7.7 1.8 9 3.3 9.3c-.3.1-.6.1-.9.1c-.2 0-.4 0-.6-.1c.4 1.3 1.6 2.3 3.1 2.3c-1.1.9-2.5 1.4-4.1 1.4H0c1.5.9 3.2 1.5 5 1.5c6 0 9.3-5 9.3-9.3v-.4C15 4.3 15.6 3.7 16 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwitterSquare(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v16h16V0zm12.8 5.6v.3c0 3.3-2.5 7-7 7c-1.4 0-2.7-.4-3.8-1.1h.6c1.2 0 2.2-.4 3.1-1.1c-1.1 0-2-.7-2.3-1.7h.5c.2 0 .4 0 .6-.1c-1.1-.2-2-1.2-2-2.4c.3.2.7.3 1.1.3c-.7-.4-1.1-1.2-1.1-2c0-.5.1-.9.3-1.2C4 5.1 5.9 6 7.9 6.1c0-.2-.1-.4-.1-.6C7.8 4.1 8.9 3 10.3 3c.7 0 1.3.3 1.8.8c.6-.1 1.1-.3 1.6-.6c-.2.6-.6 1.1-1.1 1.4c.5-.1 1-.2 1.4-.4c-.3.6-.7 1-1.2 1.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Umbrella(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.36.9L5.09.33a.54.54 0 0 0-.773-.259a.55.55 0 0 0-.316.762l.319.577C-1.88 4.9.42 12 .42 12h.06c.25-.12.8-1.64 2.05-2.25s2.78-.09 3-.21l.12-.06a11.75 11.75 0 0 1 1.58-1.97l3.37 7.07a2.364 2.364 0 0 0 1.334 1.335a1.764 1.764 0 0 0 1.237-.069l.359-.176c.263-.145.462-.38.558-.662a2.34 2.34 0 0 0 .183-.913c0-.401-.1-.778-.277-1.108a.628.628 0 0 0-.768-.286a.55.55 0 0 0-.244.752c.041.077.529 1.067-.101 1.337s-1.19-.73-1.19-.73L8.271 7a11.482 11.482 0 0 1 2.532.005l.068-.065c.25-.12.8-1.64 2.05-2.25s2.78-.09 3-.21h.06S12.001-1.93 5.361.9zm2 5.46a3.81 3.81 0 0 0-2.211 2.224c-.55-1.082-.909-2.375-1.007-3.74C4 2.6 4.75 1.92 4.75 1.92l.77-.32a4.989 4.989 0 0 1 2.816 1.195a8.689 8.689 0 0 1 2.233 3.265c-.339-.021-.752-.067-1.175-.067c-.724 0-1.417.134-2.054.379z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Underline(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 15h12v1H2zm9-15v8.4C11 9.9 9.9 11 8.4 11h-.8C6.1 11 5 9.9 5 8.4V0H2v8.4C2 11.5 4.5 14 7.6 14h.9c3.1 0 5.6-2.5 5.6-5.6V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unlink(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0h1v4H8zm0 12h1v4H8zM7 9H3a1 1 0 0 1 0-2h4V5H3a3 3 0 1 0 0 6h4zm6-4H9v2h4a1 1 0 0 1 0 2H9v2h4a3 3 0 1 0 0-6M4.51 15.44L7 12H5.77l-2.08 2.88zm7.98 0L10 12h1.23l2.08 2.88zm0-14.45L10 4h1.23l2.08-2.66zm-7.98 0L7 4H5.77L3.69 1.34z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unlock(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 8V4.9C8 2.7 6.2 1 4.1 1h-.3C1.6 1 0 2.7 0 4.9V7h2V4.9C2 3.8 2.7 3 3.8 3h.3c1 0 1.9.8 1.9 1.9V8H5l.1 5S5 16 10 16s5-3 5-3V8zm3 6h-1v-1.8c-.6 0-1-.6-1-1.1c0-.6.4-1.1 1-1.1s1 .4 1 .9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Upload(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 10v2H5v-2H0v6h16v-6zm-7 4H2v-2h2z"/><path fill="currentColor" d="M13 5L8 0L3 5h3v6h4V5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UploadAlt(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 14h16v2H0zM8 0L3 5h3v8h4V5h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func User(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 0C2.4 0 5.1 7.3 5.1 7.3c.6 1 1.4.8 1.4 1.5c0 .6-.7.8-1.4.9C4 9.7 3 9.5 2 11.3c-.6 1.1-.9 4.7-.9 4.7h13.7s-.3-3.6-.8-4.7c-1-1.9-2-1.6-3.1-1.7c-.7-.1-1.4-.3-1.4-.9s.8-.4 1.4-1.5C10.9 7.3 13.6 0 8 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserCard(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 3v10H1V3zm1-1H0v12h16z"/><path fill="currentColor" d="M8 5h6v1H8zm0 2h6v1H8zm0 2h3v1H8zM5.4 7H5v-.1c.6-.2 1-.8 1-1.4C6 4.7 5.3 4 4.5 4S3 4.7 3 5.5c0 .7.4 1.2 1 1.4V7h-.4C2.7 7 2 7.7 2 8.6V11h5V8.6C7 7.7 6.3 7 5.4 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserCheck(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 14.4c-.8-.8-.8-2 0-2.8s2-.8 2.8 0l.6.6l1.9-2.1c-.7-.4-1.3-.4-2-.4c-.7-.1-1.4-.3-1.4-.9s.8-.4 1.4-1.5c0 0 2.7-7.3-2.9-7.3c-5.5 0-2.8 7.3-2.8 7.3c.6 1 1.4.8 1.4 1.5s-.7.7-1.4.8C4 9.7 3 9.5 2 11.3c-.6 1.1-.9 4.7-.9 4.7h8zm5.3 1.6h2.1s-.1-.9-.2-2z"/><path fill="currentColor" d="M11 16c-.3 0-.5-.1-.7-.3l-2-2c-.4-.4-.4-1 0-1.4s1-.4 1.4 0l1.3 1.3l3.3-3.6c.4-.4 1-.4 1.4-.1c.4.4.4 1 .1 1.4l-4 4.3c-.3.3-.5.4-.8.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserClock(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 13h-3v-3h1v2h2z"/><path fill="currentColor" d="M16 12.5C16 10 14 8 11.5 8c-.7 0-1.4.2-2 .5c.2-.3.8-.3 1.4-1.2c0 0 2.7-7.3-2.9-7.3S5.1 7.3 5.1 7.3c.6 1 1.4.8 1.4 1.5s-.7.7-1.4.8C4 9.7 3 9.5 2 11.3c-.6 1.1-.9 4.7-.9 4.7h10.4C9.6 16 8 14.4 8 12.5S9.6 9 11.5 9s3.5 1.6 3.5 3.5s-1.6 3.5-3.5 3.5h3.4v-.5c.6-.8 1.1-1.8 1.1-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserHeart(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.2 16h.6v-.6zm-5.6-2.1c-.7-.7-1-1.8-.8-2.8S8.6 9.3 9.5 9c0-.1-.1-.2-.1-.2c0-.6.8-.4 1.4-1.5c0 0 2.7-7.3-2.9-7.3c-5.5 0-2.8 7.3-2.8 7.3c.6 1 1.4.8 1.4 1.5s-.7.7-1.4.8C4 9.7 3 9.5 2 11.3c-.6 1.1-.9 4.7-.9 4.7h9.6z"/><path fill="currentColor" d="M14.9 10.1c-.2-.1-.5-.1-.7-.1c-.7 0-1.3.6-1.7 1.1c-.4-.5-1-1.1-1.7-1.1c-.3 0-.5 0-.7.1c-1.2.4-1.4 2-.5 2.9l3 2.9l3-2.9c.8-.9.5-2.5-.7-2.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserStar(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8.92 13.67l-1.61-1.53l-1.5-1.42l2-.29l2.25-.32l.29-.57h-.02a1 1 0 0 1-.979-.794c-.001-.617.799-.417 1.429-1.457c.08-.02 2.82-7.29-2.78-7.29s-2.86 7.27-2.86 7.27c.63 1 1.44.85 1.43 1.45s-.74.8-1.43.87C4 9.719 3 9.459 2 11.349c-.6 1.09-.85 4.65-.85 4.65h7.36v-.17zm2.8 2.33h.56l-.28-.14z"/><path fill="currentColor" d="M12 14.73L14.47 16L14 13.31l2-1.9l-2.76-.39L12 8.57l-1.24 2.45l-2.76.39l2 1.9L9.53 16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Users(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.3 9.7c-.4 0-.9-.2-.9-.6s.5-.3.9-1c0 0 1.8-4.9-1.8-4.9S1.7 8.1 1.7 8.1c.4.7.9.6.9 1s-.5.6-.9.6c-.6.1-1.1 0-1.7.6V16h5c.2-1.7.7-5.2 1.1-6.1l.1-.1c-.2-.1-.5-.1-.9-.1M16 9.5c-.7-.8-1.3-.7-2-.8c-.5-.1-1.1-.2-1.1-.7s.6-.3 1.1-1.2c0 0 2.1-5.9-2.2-5.9c-4.4.1-2.3 6-2.3 6c.5.8 1.1.7 1.1 1.1c0 .5-.6.6-1.1.7c-.9.1-1.7 0-2.5 1.5c-.4.9-1 5.8-1 5.8h10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VaadinH(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.21.35a.79.79 0 0 0-.79.79v.46c0 .5-.32.85-1.07.85H9.8c-1.61 0-1.73 1.19-1.8 1.83c-.06-.64-.18-1.83-1.79-1.83H2.64c-.75 0-1.09-.37-1.09-.86v-.47A.77.77 0 0 0 .78.35a.78.78 0 0 0-.78.78v.011v-.001v1.32C0 4 .7 4.77 2.34 4.77H6c1.09 0 1.19.46 1.19.9v.13a.851.851 0 0 0 1.69.004V5.67c0-.44.1-.9 1.19-.9h3.61C15.29 4.77 16 4 16 2.46V1.14a.79.79 0 0 0-.79-.79m-4 7.03l-.04-.001a1 1 0 0 0-.958.714l-.002.007L8 12.31l-2.3-4.2a1.003 1.003 0 0 0-.963-.731l-.039.001H4.7l-.039-.001a1.05 1.05 0 0 0-.879 1.625L3.78 9l3.29 6.1a.942.942 0 0 0 1.718.006l.002-.006L12.13 9a1.05 1.05 0 0 0-.906-1.58h-.014h.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VaadinV(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.8 7.16h-.13c-.44 0-.9-.1-.9-1.19V2.35C4.77.71 4 0 2.46 0H1.14a.79.79 0 0 0 0 1.58h.46c.5 0 .85.32.85 1.07V6.2c0 1.61 1.19 1.73 1.83 1.8c-.64.06-1.83.18-1.83 1.79v3.55c0 .75-.37 1.09-.86 1.09h-.47a.77.77 0 0 0-.77.77c0 .431.349.78.78.78h1.33c1.54 0 2.31-.7 2.31-2.34v-3.59c0-1.09.46-1.19.9-1.19h.13a.851.851 0 0 0 .004-1.69H5.8zm9.3.03L9 3.87a1.05 1.05 0 0 0-1.58.906v.014v-.001l-.001.04a1 1 0 0 0 .714.958l.007.002l4.21 2.26l-4.24 2.25a1.003 1.003 0 0 0-.73 1.002v-.002l-.001.039a1.05 1.05 0 0 0 1.625.879L9 12.219l6.1-3.29a.942.942 0 0 0 .006-1.718l-.006-.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Viewport(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 4H0V0h4v1H1zm11-3V0h4v4h-1V1zm3 11h1v4h-4v-1h3zM4 15v1H0v-4h1v3z"/><path fill="currentColor" d="M13 3v10H3V3zm1-1H2v12h12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vimeo(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.9 4.4c-.9 5-5.9 9.3-7.4 10.3s-2.9-.4-3.4-1.4C4.6 12 2.9 5.7 2.4 5.1C2 4.6.6 5.7.6 5.7L0 4.8s2.7-3.3 4.8-3.7C7 .7 7 4.5 7.5 6.6c.5 2 .9 3.2 1.3 3.2s1.3-1.1 2.2-2.9c.9-1.7 0-3.3-1.9-2.2c.8-4.3 7.7-5.4 6.8-.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VimeoSquare(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v16h16V0zm13.9 5.3c-.7 3.8-4.4 7-5.5 7.7s-2.2-.3-2.5-1.1c-.4-.9-1.7-5.7-2-6.1c-.4-.3-1.4.5-1.4.5L2 5.6s2-2.4 3.6-2.7s1.6 2.5 2 4.1c.4 1.5.6 2.4 1 2.4c.3 0 1-.9 1.7-2.2s0-2.5-1.4-1.6c.5-3.3 5.7-4.1 5-.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Volume(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m11.8 2.4l-.5 1C12.4 4.8 13 6.6 13 8.5c0 1.7-.5 3.2-1.3 4.6l.7.8c1.1-1.5 1.7-3.4 1.7-5.4c-.1-2.3-.9-4.4-2.3-6.1"/><path fill="currentColor" d="m10.8 4.4l-.5 1.1c.5.9.8 1.9.8 3c0 1-.3 2-.7 2.9l.7.9c.6-1.1 1-2.4 1-3.7c-.1-1.6-.5-3-1.3-4.2M4 5H0v6h4l5 4V1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeDown(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.8 4.4l-.5 1.1c.5.9.8 1.9.8 3c0 1-.3 2-.7 2.9l.7.9c.6-1.1 1-2.4 1-3.7c-.1-1.6-.5-3-1.3-4.2M4 5H0v6h4l5 4V1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeOff(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 5H0v6h4l5 4V1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeUp(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 8.5c0 2.3-.8 4.5-2 6.2l.7.8c1.5-1.9 2.4-4.4 2.4-7c0-3.1-1.2-5.9-3.2-8l-.5 1C14 3.3 15 5.8 15 8.5"/><path fill="currentColor" d="m11.8 2.4l-.5 1C12.4 4.8 13 6.6 13 8.5c0 1.7-.5 3.2-1.3 4.6l.7.8c1.1-1.5 1.7-3.4 1.7-5.4c-.1-2.3-.9-4.4-2.3-6.1"/><path fill="currentColor" d="m10.8 4.4l-.5 1.1c.5.9.8 1.9.8 3c0 1-.3 2-.7 2.9l.7.9c.6-1.1 1-2.4 1-3.7c-.1-1.6-.5-3-1.3-4.2M4 5H0v6h4l5 4V1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wallet(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 4H2.38a1 1 0 0 1-1.19-.982v-.019L14 2.5V1.31A1.18 1.18 0 0 0 12.684.001L1.31 1.85A2.004 2.004 0 0 0 0 3.727v10.772a1.5 1.5 0 0 0 1.5 1.5h13a1.5 1.5 0 0 0 1.5-1.5v-9.01l.001-.041a1.45 1.45 0 0 0-1.45-1.45l-.053.001zM13 11a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 13 11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Warning(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 1L0 15h16zm0 12c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1s-.4 1-1 1m-1-3V6h2v4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Workplace(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 3V0H2v14H0v1h7v-5h2V8h5V3zm-5 7H4V8h2zm0-3H4V5h2zm0-3H4V2h2zm3 3H7V5h2zm0-3H7V2h2zm4 3h-2V5h2zm1 4h2v5H8v-5h2V9h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wrench(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.5 13.4L7.7 5.6c.2-.5.3-1 .3-1.6c0-2.2-1.8-4-4-4c-.6 0-1.1.1-1.6.3l2.9 2.9l-2.1 2.1L.3 2.4C.1 2.9 0 3.4 0 4c0 2.2 1.8 4 4 4c.6 0 1.1-.1 1.6-.3l7.8 7.8c.6.6 1.5.6 2.1 0s.6-1.5 0-2.1M6.8 7.6L5.4 6.2l.9-.9l1.4 1.4zm7.4 7.4c-.4 0-.8-.3-.8-.8c0-.4.3-.8.8-.8s.8.3.8.8s-.3.8-.8.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Youtube(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.6 0h-.9l-.6 2.3L4.5 0h-1c.2.6.4 1.1.6 1.7c.3.8.5 1.5.5 1.9V6h.9V3.6zM9 4.5V3c0-.5-.1-.8-.3-1.1s-.5-.4-.9-.4s-.7.2-.9.5c-.2.2-.3.5-.3 1v1.6c0 .5.1.8.3 1c.2.3.5.4.9.4s.7-.2.9-.5c.2-.1.3-.5.3-1m-.8.2c0 .4-.1.6-.4.6s-.4-.2-.4-.6V2.8c0-.4.1-.6.4-.6s.4.2.4.6zM12 6V1.5h-.8v3.4c-.2.3-.3.4-.5.4c-.1 0-.2-.1-.2-.2V1.5h-.8V5c0 .3 0 .5.1.7c0 .2.2.3.5.3s.6-.2.9-.5V6zm.4 4.5c-.3 0-.4.2-.4.6v.4h.8v-.4c0-.4-.1-.6-.4-.6m-2.9 0c-.1 0-.3.1-.4.2v2.7c.1.1.3.2.4.2c.2 0 .3-.2.3-.6v-1.9c0-.4-.1-.6-.3-.6"/><path fill="currentColor" d="M14.4 8.3C14.2 7.6 13.6 7 13 7c-1.6-.2-3.3-.2-5-.2s-3.3 0-5 .2c-.6 0-1.2.6-1.4 1.3c-.2 1-.2 2.1-.2 3.1s0 2.1.2 3.1c.2.7.7 1.2 1.4 1.3c1.7.2 3.3.2 5 .2s3.3 0 5-.2c.7-.1 1.3-.6 1.4-1.3c.2-1 .2-2.1.2-3.1s0-2.1-.2-3.1m-9.2.9h-1v5.1h-.9V9.2h-.9v-.9h2.8zm2.4 5.1h-.8v-.5c-.3.4-.6.5-.9.5s-.4-.1-.5-.3c0-.1-.1-.3-.1-.7V9.8h.8v3.5c0 .1.1.2.2.2c.2 0 .3-.1.5-.4V9.8h.8zm3-1.4c0 .4 0 .7-.1.9c-.1.3-.3.5-.6.5s-.6-.2-.8-.5v.4h-.8V8.3h.8v1.9c.3-.3.5-.5.8-.5s.5.2.6.5c.1.2.1.5.1.9zm3-.7H12v.8c0 .4.1.6.4.6c.2 0 .3-.1.4-.3v-.5h.8v.6c0 .2-.1.3-.2.5c-.2.3-.5.5-1 .5c-.4 0-.7-.2-1-.5c-.2-.2-.3-.6-.3-1v-1.5c0-.5.1-.8.2-1c.2-.3.5-.5 1-.5c.4 0 .7.2.9.5c.2.2.2.6.2 1v.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YoutubeSquare(children ...ElementRenderer) *VaadinIcon {
	return &VaadinIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.9 6c.2 0 .3-.2.3-.5V4.1c0-.3-.1-.5-.3-.5s-.3.2-.3.5v1.4c0 .3.1.5.3.5m-.8 5.9c-.1.2-.3.3-.4.3s-.1 0-.1-.1V9.4H6V12c0 .2 0 .4.1.5c.1.2.2.2.4.2s.4-.1.7-.4v.4h.6V9.4h-.7zm-3.3-3h.7v3.8h.7V8.9h.7v-.7H3.8zm5.6.4c-.2 0-.4.2-.6.4V8.2h-.6v4.4h.6v-.3c.2.2.4.4.6.4s.4-.1.5-.4c0-.1.1-.4.1-.7v-1.3c0-.3 0-.5-.1-.7c-.1-.1-.2-.3-.5-.3m0 2.4c0 .3-.1.4-.3.4c-.1 0-.2 0-.3-.1v-2c.1-.1.2-.1.3-.1c.2 0 .3.2.3.5zm1.9-2.4c-.3 0-.5.1-.7.3c-.1.2-.2.4-.2.8v1.2c0 .4.1.6.2.8c.2.2.4.3.7.3s.6-.1.7-.4c.1-.1.1-.2.1-.4v-.5h-.6v.4c0 .2-.1.2-.3.2s-.3-.2-.3-.5v-.6h1.2v-.7c0-.4-.1-.6-.2-.8c0 .1-.3-.1-.6-.1m.3 1.3H11v-.3c0-.3.1-.5.3-.5s.3.2.3.5z"/><path fill="currentColor" d="M0 0v16h16V0zm9.3 3.1h.6v2.7c0 .1 0 .2.1.2s.2-.1.4-.3V3.1h.6v3.3h-.6v-.3c-.2.3-.5.4-.7.4s-.3-.1-.4-.2c0-.1-.1-.3-.1-.5V3.1zM7 4.2c0-.3 0-.6.2-.8s.4-.3.7-.3c.3 0 .5.1.7.3c.1.2.2.4.2.8v1.2c0 .4-.1.6-.2.8c-.2.2-.4.3-.7.3s-.5-.1-.7-.3C7 6 7 5.8 7 5.4zM5.3 2l.5 1.8l.5-1.8H7l-.8 2.7v1.8h-.7V4.7c-.1-.4-.2-.8-.4-1.5c-.2-.4-.3-.8-.5-1.2zm7.5 10.9c-.1.5-.6.9-1.1 1c-1.2.1-2.5.1-3.7.1s-2.5 0-3.7-.1c-.5-.1-1-.4-1.1-1c-.2-.8-.2-1.6-.2-2.4c0-.7 0-1.5.2-2.3c.1-.5.6-.9 1.1-1c1.2-.1 2.5-.1 3.7-.1s2.5 0 3.7.1c.5.1 1 .4 1.1 1c.2.8.2 1.6.2 2.3c0 .8 0 1.6-.2 2.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
