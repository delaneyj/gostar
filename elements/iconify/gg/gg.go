package gg

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = "2.1.1"
	hAttr          = "1em"
	viewbox        = "0 0 24 24"
	fill           = "currentColor"
)

type GgIcon struct {
	*SVGSVGElement
}

type GgIconFn func(children ...ElementRenderer) *GgIcon

var IconLookup = map[string]GgIconFn{
	"abstract":              Abstract,
	"add":                   Add,
	"addR":                  AddR,
	"adidas":                Adidas,
	"airplane":              Airplane,
	"alarm":                 Alarm,
	"album":                 Album,
	"alignBottom":           AlignBottom,
	"alignCenter":           AlignCenter,
	"alignLeft":             AlignLeft,
	"alignMiddle":           AlignMiddle,
	"alignRight":            AlignRight,
	"alignTop":              AlignTop,
	"anchor":                Anchor,
	"appleWatch":            AppleWatch,
	"arrangeBack":           ArrangeBack,
	"arrangeFront":          ArrangeFront,
	"arrowAlignH":           ArrowAlignH,
	"arrowAlignV":           ArrowAlignV,
	"arrowBottomLeft":       ArrowBottomLeft,
	"arrowBottomLeftO":      ArrowBottomLeftO,
	"arrowBottomLeftR":      ArrowBottomLeftR,
	"arrowBottomRight":      ArrowBottomRight,
	"arrowBottomRightO":     ArrowBottomRightO,
	"arrowBottomRightR":     ArrowBottomRightR,
	"arrowDown":             ArrowDown,
	"arrowDownO":            ArrowDownO,
	"arrowDownR":            ArrowDownR,
	"arrowLeft":             ArrowLeft,
	"arrowLeftO":            ArrowLeftO,
	"arrowLeftR":            ArrowLeftR,
	"arrowLongDown":         ArrowLongDown,
	"arrowLongDownC":        ArrowLongDownC,
	"arrowLongDownE":        ArrowLongDownE,
	"arrowLongDownL":        ArrowLongDownL,
	"arrowLongDownR":        ArrowLongDownR,
	"arrowLongLeft":         ArrowLongLeft,
	"arrowLongLeftC":        ArrowLongLeftC,
	"arrowLongLeftE":        ArrowLongLeftE,
	"arrowLongLeftL":        ArrowLongLeftL,
	"arrowLongLeftR":        ArrowLongLeftR,
	"arrowLongRight":        ArrowLongRight,
	"arrowLongRightC":       ArrowLongRightC,
	"arrowLongRightE":       ArrowLongRightE,
	"arrowLongRightL":       ArrowLongRightL,
	"arrowLongRightR":       ArrowLongRightR,
	"arrowLongUp":           ArrowLongUp,
	"arrowLongUpC":          ArrowLongUpC,
	"arrowLongUpE":          ArrowLongUpE,
	"arrowLongUpL":          ArrowLongUpL,
	"arrowLongUpR":          ArrowLongUpR,
	"arrowRight":            ArrowRight,
	"arrowRightO":           ArrowRightO,
	"arrowRightR":           ArrowRightR,
	"arrowTopLeft":          ArrowTopLeft,
	"arrowTopLeftO":         ArrowTopLeftO,
	"arrowTopLeftR":         ArrowTopLeftR,
	"arrowTopRight":         ArrowTopRight,
	"arrowTopRightO":        ArrowTopRightO,
	"arrowTopRightR":        ArrowTopRightR,
	"arrowUp":               ArrowUp,
	"arrowUpO":              ArrowUpO,
	"arrowUpR":              ArrowUpR,
	"arrowsBreakeH":         ArrowsBreakeH,
	"arrowsBreakeV":         ArrowsBreakeV,
	"arrowsExchange":        ArrowsExchange,
	"arrowsExchangeAlt":     ArrowsExchangeAlt,
	"arrowsExchangeAltV":    ArrowsExchangeAltV,
	"arrowsExchangeV":       ArrowsExchangeV,
	"arrowsExpandDownLeft":  ArrowsExpandDownLeft,
	"arrowsExpandDownRight": ArrowsExpandDownRight,
	"arrowsExpandLeft":      ArrowsExpandLeft,
	"arrowsExpandLeftAlt":   ArrowsExpandLeftAlt,
	"arrowsExpandRight":     ArrowsExpandRight,
	"arrowsExpandRightAlt":  ArrowsExpandRightAlt,
	"arrowsExpandUpLeft":    ArrowsExpandUpLeft,
	"arrowsExpandUpRight":   ArrowsExpandUpRight,
	"arrowsH":               ArrowsH,
	"arrowsHalt":            ArrowsHalt,
	"arrowsMergeAltH":       ArrowsMergeAltH,
	"arrowsMergeAltV":       ArrowsMergeAltV,
	"arrowsScrollH":         ArrowsScrollH,
	"arrowsScrollV":         ArrowsScrollV,
	"arrowsShrinkH":         ArrowsShrinkH,
	"arrowsShrinkV":         ArrowsShrinkV,
	"arrowsV":               ArrowsV,
	"arrowsValt":            ArrowsValt,
	"assign":                Assign,
	"asterisk":              Asterisk,
	"atlasian":              Atlasian,
	"attachment":            Attachment,
	"attribution":           Attribution,
	"awards":                Awards,
	"backspace":             Backspace,
	"bandAid":               BandAid,
	"battery":               Battery,
	"batteryEmpty":          BatteryEmpty,
	"batteryFull":           BatteryFull,
	"bee":                   Bee,
	"bell":                  Bell,
	"bitbucket":             Bitbucket,
	"block":                 Block,
	"bmw":                   Bmw,
	"board":                 Board,
	"bolt":                  Bolt,
	"bookmark":              Bookmark,
	"borderAll":             BorderAll,
	"borderBottom":          BorderBottom,
	"borderLeft":            BorderLeft,
	"borderRight":           BorderRight,
	"borderStyleDashed":     BorderStyleDashed,
	"borderStyleDotted":     BorderStyleDotted,
	"borderStyleSolid":      BorderStyleSolid,
	"borderTop":             BorderTop,
	"bot":                   Bot,
	"bowl":                  Bowl,
	"box":                   Box,
	"boy":                   Boy,
	"brackets":              Brackets,
	"briefcase":             Briefcase,
	"browse":                Browse,
	"browser":               Browser,
	"brush":                 Brush,
	"bulb":                  Bulb,
	"cplusPlus":             CplusPlus,
	"calculator":            Calculator,
	"calendar":              Calendar,
	"calendarDates":         CalendarDates,
	"calendarDue":           CalendarDue,
	"calendarNext":          CalendarNext,
	"calendarToday":         CalendarToday,
	"calendarTwo":           CalendarTwo,
	"calibrate":             Calibrate,
	"camera":                Camera,
	"cap":                   Cap,
	"captions":              Captions,
	"cardClubs":             CardClubs,
	"cardDiamonds":          CardDiamonds,
	"cardHearts":            CardHearts,
	"cardSpades":            CardSpades,
	"carousel":              Carousel,
	"cast":                  Cast,
	"chanel":                Chanel,
	"chart":                 Chart,
	"check":                 Check,
	"checkO":                CheckO,
	"checkR":                CheckR,
	"chevronDoubleDown":     ChevronDoubleDown,
	"chevronDoubleDownO":    ChevronDoubleDownO,
	"chevronDoubleDownR":    ChevronDoubleDownR,
	"chevronDoubleLeft":     ChevronDoubleLeft,
	"chevronDoubleLeftO":    ChevronDoubleLeftO,
	"chevronDoubleLeftR":    ChevronDoubleLeftR,
	"chevronDoubleRight":    ChevronDoubleRight,
	"chevronDoubleRightO":   ChevronDoubleRightO,
	"chevronDoubleRightR":   ChevronDoubleRightR,
	"chevronDoubleUp":       ChevronDoubleUp,
	"chevronDoubleUpO":      ChevronDoubleUpO,
	"chevronDoubleUpR":      ChevronDoubleUpR,
	"chevronDown":           ChevronDown,
	"chevronDownO":          ChevronDownO,
	"chevronDownR":          ChevronDownR,
	"chevronLeft":           ChevronLeft,
	"chevronLeftO":          ChevronLeftO,
	"chevronLeftR":          ChevronLeftR,
	"chevronRight":          ChevronRight,
	"chevronRightO":         ChevronRightO,
	"chevronRightR":         ChevronRightR,
	"chevronUp":             ChevronUp,
	"chevronUpO":            ChevronUpO,
	"chevronUpR":            ChevronUpR,
	"circleci":              Circleci,
	"clapperBoard":          ClapperBoard,
	"clipboard":             Clipboard,
	"close":                 Close,
	"closeO":                CloseO,
	"closeR":                CloseR,
	"cloud":                 Cloud,
	"code":                  Code,
	"codeClimate":           CodeClimate,
	"codeSlash":             CodeSlash,
	"coffee":                Coffee,
	"collage":               Collage,
	"colorBucket":           ColorBucket,
	"colorPicker":           ColorPicker,
	"comedyCentral":         ComedyCentral,
	"comment":               Comment,
	"community":             Community,
	"components":            Components,
	"compress":              Compress,
	"compressLeft":          CompressLeft,
	"compressRight":         CompressRight,
	"compressV":             CompressV,
	"controller":            Controller,
	"copy":                  Copy,
	"copyright":             Copyright,
	"cornerDoubleDownLeft":  CornerDoubleDownLeft,
	"cornerDoubleDownRight": CornerDoubleDownRight,
	"cornerDoubleLeftDown":  CornerDoubleLeftDown,
	"cornerDoubleLeftUp":    CornerDoubleLeftUp,
	"cornerDoubleRightDown": CornerDoubleRightDown,
	"cornerDoubleRightUp":   CornerDoubleRightUp,
	"cornerDoubleUpLeft":    CornerDoubleUpLeft,
	"cornerDoubleUpRight":   CornerDoubleUpRight,
	"cornerDownLeft":        CornerDownLeft,
	"cornerDownRight":       CornerDownRight,
	"cornerLeftDown":        CornerLeftDown,
	"cornerLeftUp":          CornerLeftUp,
	"cornerRightDown":       CornerRightDown,
	"cornerRightUp":         CornerRightUp,
	"cornerUpLeft":          CornerUpLeft,
	"cornerUpRight":         CornerUpRight,
	"creditCard":            CreditCard,
	"crop":                  Crop,
	"cross":                 Cross,
	"crowdfire":             Crowdfire,
	"crown":                 Crown,
	"danger":                Danger,
	"darkMode":              DarkMode,
	"data":                  Data,
	"database":              Database,
	"debug":                 Debug,
	"designmodo":            Designmodo,
	"desktop":               Desktop,
	"detailsLess":           DetailsLess,
	"detailsMore":           DetailsMore,
	"dialpad":               Dialpad,
	"diceFive":              DiceFive,
	"diceFour":              DiceFour,
	"diceOne":               DiceOne,
	"diceSix":               DiceSix,
	"diceThree":             DiceThree,
	"diceTwo":               DiceTwo,
	"digitalocean":          Digitalocean,
	"disc":                  Disc,
	"displayFlex":           DisplayFlex,
	"displayFullwidth":      DisplayFullwidth,
	"displayGrid":           DisplayGrid,
	"displaySpacing":        DisplaySpacing,
	"distributeHorizontal":  DistributeHorizontal,
	"distributeVertical":    DistributeVertical,
	"dockBottom":            DockBottom,
	"dockLeft":              DockLeft,
	"dockRight":             DockRight,
	"dockWindow":            DockWindow,
	"dolby":                 Dolby,
	"dollar":                Dollar,
	"dribbble":              Dribbble,
	"drive":                 Drive,
	"drop":                  Drop,
	"dropInvert":            DropInvert,
	"dropOpacity":           DropOpacity,
	"duplicate":             Duplicate,
	"editBlackPoint":        EditBlackPoint,
	"editContrast":          EditContrast,
	"editExposure":          EditExposure,
	"editFade":              EditFade,
	"editFlipH":             EditFlipH,
	"editFlipV":             EditFlipV,
	"editHighlight":         EditHighlight,
	"editMarkup":            EditMarkup,
	"editMask":              EditMask,
	"editNoise":             EditNoise,
	"editShadows":           EditShadows,
	"editStraight":          EditStraight,
	"editUnmask":            EditUnmask,
	"eject":                 Eject,
	"enter":                 Enter,
	"erase":                 Erase,
	"ereader":               Ereader,
	"ericsson":              Ericsson,
	"ethernet":              Ethernet,
	"euro":                  Euro,
	"eventbrite":            Eventbrite,
	"expand":                Expand,
	"export":                Export,
	"extension":             Extension,
	"extensionAdd":          ExtensionAdd,
	"extensionAlt":          ExtensionAlt,
	"extensionRemove":       ExtensionRemove,
	"external":              External,
	"eye":                   Eye,
	"eyeAlt":                EyeAlt,
	"facebook":              Facebook,
	"feed":                  Feed,
	"figma":                 Figma,
	"file":                  File,
	"fileAdd":               FileAdd,
	"fileDocument":          FileDocument,
	"fileRemove":            FileRemove,
	"film":                  Film,
	"filters":               Filters,
	"flag":                  Flag,
	"flagAlt":               FlagAlt,
	"folder":                Folder,
	"folderAdd":             FolderAdd,
	"folderRemove":          FolderRemove,
	"fontHeight":            FontHeight,
	"fontSpacing":           FontSpacing,
	"formatBold":            FormatBold,
	"formatCenter":          FormatCenter,
	"formatColor":           FormatColor,
	"formatHeading":         FormatHeading,
	"formatIndentDecrease":  FormatIndentDecrease,
	"formatIndentIncrease":  FormatIndentIncrease,
	"formatItalic":          FormatItalic,
	"formatJustify":         FormatJustify,
	"formatLeft":            FormatLeft,
	"formatLineHeight":      FormatLineHeight,
	"formatRight":           FormatRight,
	"formatSeparator":       FormatSeparator,
	"formatSlash":           FormatSlash,
	"formatStrike":          FormatStrike,
	"formatText":            FormatText,
	"formatUnderline":       FormatUnderline,
	"formatUppercase":       FormatUppercase,
	"framer":                Framer,
	"games":                 Games,
	"genderFemale":          GenderFemale,
	"genderMale":            GenderMale,
	"ghost":                 Ghost,
	"ghostCharacter":        GhostCharacter,
	"gift":                  Gift,
	"girl":                  Girl,
	"gitBranch":             GitBranch,
	"gitCommit":             GitCommit,
	"gitFork":               GitFork,
	"gitPull":               GitPull,
	"gitter":                Gitter,
	"glass":                 Glass,
	"glassAlt":              GlassAlt,
	"globe":                 Globe,
	"globeAlt":              GlobeAlt,
	"google":                Google,
	"googleTasks":           GoogleTasks,
	"gym":                   Gym,
	"hashtag":               Hashtag,
	"headset":               Headset,
	"heart":                 Heart,
	"hello":                 Hello,
	"home":                  Home,
	"homeAlt":               HomeAlt,
	"homeScreen":            HomeScreen,
	"icecream":              Icecream,
	"ifDesign":              IfDesign,
	"image":                 Image,
	"import":                Import,
	"inbox":                 Inbox,
	"indieHackers":          IndieHackers,
	"infinity":              Infinity,
	"info":                  Info,
	"inpicture":             Inpicture,
	"insertAfter":           InsertAfter,
	"insertAfterO":          InsertAfterO,
	"insertAfterR":          InsertAfterR,
	"insertBefore":          InsertBefore,
	"insertBeforeO":         InsertBeforeO,
	"insertBeforeR":         InsertBeforeR,
	"insights":              Insights,
	"instagram":             Instagram,
	"internal":              Internal,
	"key":                   Key,
	"keyboard":              Keyboard,
	"keyhole":               Keyhole,
	"laptop":                Laptop,
	"lastpass":              Lastpass,
	"layoutGrid":            LayoutGrid,
	"layoutGridSmall":       LayoutGridSmall,
	"layoutList":            LayoutList,
	"layoutPin":             LayoutPin,
	"linear":                Linear,
	"link":                  Link,
	"list":                  List,
	"listTree":              ListTree,
	"livePhoto":             LivePhoto,
	"loadbar":               Loadbar,
	"loadbarAlt":            LoadbarAlt,
	"loadbarDoc":            LoadbarDoc,
	"loadbarSound":          LoadbarSound,
	"lock":                  Lock,
	"lockUnlock":            LockUnlock,
	"logIn":                 LogIn,
	"logOff":                LogOff,
	"logOut":                LogOut,
	"loupe":                 Loupe,
	"magnet":                Magnet,
	"mail":                  Mail,
	"mailForward":           MailForward,
	"mailOpen":              MailOpen,
	"mailReply":             MailReply,
	"mathDivide":            MathDivide,
	"mathEqual":             MathEqual,
	"mathMinus":             MathMinus,
	"mathPercent":           MathPercent,
	"mathPlus":              MathPlus,
	"maximize":              Maximize,
	"maximizeAlt":           MaximizeAlt,
	"maze":                  Maze,
	"mediaLive":             MediaLive,
	"mediaPodcast":          MediaPodcast,
	"menu":                  Menu,
	"menuBoxed":             MenuBoxed,
	"menuCake":              MenuCake,
	"menuCheese":            MenuCheese,
	"menuGridO":             MenuGridO,
	"menuGridR":             MenuGridR,
	"menuHotdog":            MenuHotdog,
	"menuLeft":              MenuLeft,
	"menuLeftAlt":           MenuLeftAlt,
	"menuMotion":            MenuMotion,
	"menuOreos":             MenuOreos,
	"menuRight":             MenuRight,
	"menuRightAlt":          MenuRightAlt,
	"menuRound":             MenuRound,
	"mergeHorizontal":       MergeHorizontal,
	"mergeVertical":         MergeVertical,
	"mic":                   Mic,
	"microbit":              Microbit,
	"microsoft":             Microsoft,
	"miniPlayer":            MiniPlayer,
	"minimize":              Minimize,
	"minimizeAlt":           MinimizeAlt,
	"modem":                 Modem,
	"monday":                Monday,
	"moon":                  Moon,
	"more":                  More,
	"moreAlt":               MoreAlt,
	"moreO":                 MoreO,
	"moreR":                 MoreR,
	"moreVertical":          MoreVertical,
	"moreVerticalAlt":       MoreVerticalAlt,
	"moreVerticalO":         MoreVerticalO,
	"moreVerticalR":         MoreVerticalR,
	"mouse":                 Mouse,
	"moveDown":              MoveDown,
	"moveLeft":              MoveLeft,
	"moveRight":             MoveRight,
	"moveTask":              MoveTask,
	"moveUp":                MoveUp,
	"music":                 Music,
	"musicNote":             MusicNote,
	"musicSpeaker":          MusicSpeaker,
	"nametag":               Nametag,
	"notes":                 Notes,
	"notifications":         Notifications,
	"npm":                   Npm,
	"oculus":                Oculus,
	"openCollective":        OpenCollective,
	"options":               Options,
	"organisation":          Organisation,
	"overflow":              Overflow,
	"pacman":                Pacman,
	"password":              Password,
	"pathBack":              PathBack,
	"pathCrop":              PathCrop,
	"pathDivide":            PathDivide,
	"pathExclude":           PathExclude,
	"pathFront":             PathFront,
	"pathIntersect":         PathIntersect,
	"pathOutline":           PathOutline,
	"pathTrim":              PathTrim,
	"pathUnite":             PathUnite,
	"patreon":               Patreon,
	"paypal":                Paypal,
	"pen":                   Pen,
	"pentagonBottomLeft":    PentagonBottomLeft,
	"pentagonBottomRight":   PentagonBottomRight,
	"pentagonDown":          PentagonDown,
	"pentagonLeft":          PentagonLeft,
	"pentagonRight":         PentagonRight,
	"pentagonTopLeft":       PentagonTopLeft,
	"pentagonTopRight":      PentagonTopRight,
	"pentagonUp":            PentagonUp,
	"performance":           Performance,
	"pexels":                Pexels,
	"phone":                 Phone,
	"photoscan":             Photoscan,
	"piano":                 Piano,
	"pill":                  Pill,
	"pin":                   Pin,
	"pinAlt":                PinAlt,
	"pinBottom":             PinBottom,
	"pinTop":                PinTop,
	"playBackwards":         PlayBackwards,
	"playButton":            PlayButton,
	"playButtonO":           PlayButtonO,
	"playButtonR":           PlayButtonR,
	"playForwards":          PlayForwards,
	"playList":              PlayList,
	"playListAdd":           PlayListAdd,
	"playListCheck":         PlayListCheck,
	"playListRemove":        PlayListRemove,
	"playListSearch":        PlayListSearch,
	"playPause":             PlayPause,
	"playPauseO":            PlayPauseO,
	"playPauseR":            PlayPauseR,
	"playStop":              PlayStop,
	"playStopO":             PlayStopO,
	"playStopR":             PlayStopR,
	"playTrackNext":         PlayTrackNext,
	"playTrackNextO":        PlayTrackNextO,
	"playTrackNextR":        PlayTrackNextR,
	"playTrackPrev":         PlayTrackPrev,
	"playTrackPrevO":        PlayTrackPrevO,
	"playTrackPrevR":        PlayTrackPrevR,
	"plug":                  Plug,
	"pocket":                Pocket,
	"pokemon":               Pokemon,
	"polaroid":              Polaroid,
	"poll":                  Poll,
	"presentation":          Presentation,
	"printer":               Printer,
	"productHunt":           ProductHunt,
	"profile":               Profile,
	"pullClear":             PullClear,
	"pushChevronDown":       PushChevronDown,
	"pushChevronDownO":      PushChevronDownO,
	"pushChevronDownR":      PushChevronDownR,
	"pushChevronLeft":       PushChevronLeft,
	"pushChevronLeftO":      PushChevronLeftO,
	"pushChevronLeftR":      PushChevronLeftR,
	"pushChevronRight":      PushChevronRight,
	"pushChevronRightO":     PushChevronRightO,
	"pushChevronRightR":     PushChevronRightR,
	"pushChevronUp":         PushChevronUp,
	"pushChevronUpO":        PushChevronUpO,
	"pushChevronUpR":        PushChevronUpR,
	"pushDown":              PushDown,
	"pushLeft":              PushLeft,
	"pushRight":             PushRight,
	"pushUp":                PushUp,
	"qr":                    Qr,
	"quote":                 Quote,
	"quoteO":                QuoteO,
	"radioCheck":            RadioCheck,
	"radioChecked":          RadioChecked,
	"ratio":                 Ratio,
	"read":                  Read,
	"readme":                Readme,
	"record":                Record,
	"redo":                  Redo,
	"remote":                Remote,
	"remove":                Remove,
	"removeR":               RemoveR,
	"rename":                Rename,
	"reorder":               Reorder,
	"repeat":                Repeat,
	"ring":                  Ring,
	"rowFirst":              RowFirst,
	"rowLast":               RowLast,
	"ruler":                 Ruler,
	"sandClock":             SandClock,
	"scan":                  Scan,
	"screen":                Screen,
	"screenMirror":          ScreenMirror,
	"screenShot":            ScreenShot,
	"screenWide":            ScreenWide,
	"scrollH":               ScrollH,
	"scrollV":               ScrollV,
	"search":                Search,
	"searchFound":           SearchFound,
	"searchLoading":         SearchLoading,
	"select":                Select,
	"selectO":               SelectO,
	"selectR":               SelectR,
	"server":                Server,
	"serverless":            Serverless,
	"shapeCircle":           ShapeCircle,
	"shapeHalfCircle":       ShapeHalfCircle,
	"shapeHexagon":          ShapeHexagon,
	"shapeRhombus":          ShapeRhombus,
	"shapeSquare":           ShapeSquare,
	"shapeTriangle":         ShapeTriangle,
	"shapeZigzag":           ShapeZigzag,
	"share":                 Share,
	"shield":                Shield,
	"shoppingBag":           ShoppingBag,
	"shoppingCart":          ShoppingCart,
	"shortcut":              Shortcut,
	"shutterstock":          Shutterstock,
	"sidebar":               Sidebar,
	"sidebarOpen":           SidebarOpen,
	"sidebarRight":          SidebarRight,
	"signal":                Signal,
	"size":                  Size,
	"sketch":                Sketch,
	"slack":                 Slack,
	"sleep":                 Sleep,
	"smartHomeBoiler":       SmartHomeBoiler,
	"smartHomeCooker":       SmartHomeCooker,
	"smartHomeHeat":         SmartHomeHeat,
	"smartHomeLight":        SmartHomeLight,
	"smartHomeRefrigerator": SmartHomeRefrigerator,
	"smartHomeWashMachine":  SmartHomeWashMachine,
	"smartphone":            Smartphone,
	"smartphoneChip":        SmartphoneChip,
	"smartphoneRam":         SmartphoneRam,
	"smartphoneShake":       SmartphoneShake,
	"smile":                 Smile,
	"smileMouthOpen":        SmileMouthOpen,
	"smileNeutral":          SmileNeutral,
	"smileNoMouth":          SmileNoMouth,
	"smileNone":             SmileNone,
	"smileSad":              SmileSad,
	"smileUpside":           SmileUpside,
	"softwareDownload":      SoftwareDownload,
	"softwareUpload":        SoftwareUpload,
	"sortAz":                SortAz,
	"sortZa":                SortZa,
	"spaceBetween":          SpaceBetween,
	"spaceBetweenV":         SpaceBetweenV,
	"spectrum":              Spectrum,
	"spinner":               Spinner,
	"spinnerAlt":            SpinnerAlt,
	"spinnerTwo":            SpinnerTwo,
	"spinnerTwoAlt":         SpinnerTwoAlt,
	"square":                Square,
	"stack":                 Stack,
	"stark":                 Stark,
	"stopwatch":             Stopwatch,
	"stories":               Stories,
	"studio":                Studio,
	"style":                 Style,
	"sun":                   Sun,
	"support":               Support,
	"swap":                  Swap,
	"swapVertical":          SwapVertical,
	"sweden":                Sweden,
	"swiss":                 Swiss,
	"sync":                  Sync,
	"tab":                   Tab,
	"tag":                   Tag,
	"tally":                 Tally,
	"tapDouble":             TapDouble,
	"tapSingle":             TapSingle,
	"template":              Template,
	"tennis":                Tennis,
	"terminal":              Terminal,
	"terrain":               Terrain,
	"thermometer":           Thermometer,
	"thermostat":            Thermostat,
	"tikcode":               Tikcode,
	"time":                  Time,
	"timelapse":             Timelapse,
	"timer":                 Timer,
	"today":                 Today,
	"toggleOff":             ToggleOff,
	"toggleOn":              ToggleOn,
	"toggleSquare":          ToggleSquare,
	"toggleSquareOff":       ToggleSquareOff,
	"toolbarBottom":         ToolbarBottom,
	"toolbarLeft":           ToolbarLeft,
	"toolbarRight":          ToolbarRight,
	"toolbarTop":            ToolbarTop,
	"toolbox":               Toolbox,
	"touchpad":              Touchpad,
	"track":                 Track,
	"transcript":            Transcript,
	"trash":                 Trash,
	"trashEmpty":            TrashEmpty,
	"tree":                  Tree,
	"trees":                 Trees,
	"trello":                Trello,
	"trending":              Trending,
	"trendingDown":          TrendingDown,
	"trophy":                Trophy,
	"tv":                    Tv,
	"twilio":                Twilio,
	"twitter":               Twitter,
	"uiKit":                 UiKit,
	"umbrella":              Umbrella,
	"unavailable":           Unavailable,
	"unblock":               Unblock,
	"undo":                  Undo,
	"unfold":                Unfold,
	"unsplash":              Unsplash,
	"usb":                   Usb,
	"usbC":                  UsbC,
	"user":                  User,
	"userAdd":               UserAdd,
	"userList":              UserList,
	"userRemove":            UserRemove,
	"userlane":              Userlane,
	"vercel":                Vercel,
	"viewCols":              ViewCols,
	"viewComfortable":       ViewComfortable,
	"viewDay":               ViewDay,
	"viewGrid":              ViewGrid,
	"viewList":              ViewList,
	"viewMonth":             ViewMonth,
	"viewSplit":             ViewSplit,
	"vinyl":                 Vinyl,
	"voicemail":             Voicemail,
	"voicemailO":            VoicemailO,
	"voicemailR":            VoicemailR,
	"volume":                Volume,
	"webcam":                Webcam,
	"website":               Website,
	"windows":               Windows,
	"workAlt":               WorkAlt,
	"yinyang":               Yinyang,
	"youtube":               Youtube,
	"zoomIn":                ZoomIn,
	"zoomOut":               ZoomOut,
}

func Abstract(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M5 5h14v14h-3V8H5z"/><path fill-rule="evenodd" d="M10 19a5 5 0 1 0 0-10a5 5 0 0 0 0 10m0-3a2 2 0 1 0 0-4a2 2 0 0 0 0 4" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Add(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m10-8a8 8 0 1 0 0 16a8 8 0 0 0 0-16"/><path d="M13 7a1 1 0 1 0-2 0v4H7a1 1 0 1 0 0 2h4v4a1 1 0 1 0 2 0v-4h4a1 1 0 1 0 0-2h-4z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AddR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M12 6a1 1 0 0 1 1 1v4h4a1 1 0 1 1 0 2h-4v4a1 1 0 1 1-2 0v-4H7a1 1 0 1 1 0-2h4V7a1 1 0 0 1 1-1"/><path fill-rule="evenodd" d="M5 22a3 3 0 0 1-3-3V5a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v14a3 3 0 0 1-3 3zm-1-3a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Adidas(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1.33 19l-.6-1.036l4.33-2.5L7.103 19zm13.856 0H9.412l-3.619-6.268l4.33-2.5zm8.083 0h-5.774l-6.64-11.5l4.33-2.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Airplane(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.01 5.128h2c1.104 0 2.458.769 3.024 1.718L16.509 11h4.5a1 1 0 1 1 0 2h-4.595l-2.476 4.154c-.565.95-1.919 1.718-3.024 1.718h-2l3.5-5.872h-6.99L3.99 15.453h-2L4.01 12v-.033l-2-3.42h2L5.444 11h7.065z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Alarm(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M5.459 2L1 6.015L2.338 7.5l4.46-4.015zM11 8h2v4h3v2h-5z"/><path fill-rule="evenodd" d="M3 12a9 9 0 1 1 18 0a9 9 0 0 1-18 0m2 0a7 7 0 1 1 14 0a7 7 0 0 1-14 0" clip-rule="evenodd"/><path d="M18.541 2L23 6.015L21.662 7.5l-4.46-4.015z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Album(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 19a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3H5a3 3 0 0 0-3 3zm18 0a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v8.011h2.395L14 9.864l1.605 2.147H18V4h1a1 1 0 0 1 1 1zM16 4h-4v5.336l2-2.676l2 2.676z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignBottom(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-opacity=".5" d="M13 10h4v6h-4z"/><path d="M11 4H7v12h4zm7 14H6v2h12z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignCenter(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-opacity=".5" d="M9 13h6v4H9z"/><path d="M6 7h12v4H6z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignLeft(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-opacity=".5" d="M8 13h6v4H8z"/><path d="M6 6H4v12h2zm14 1H8v4h12z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignMiddle(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-opacity=".5" d="M13 9h4v6h-4z"/><path d="M7 6h4v12H7z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignRight(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path fill-opacity=".5" d="m16 13.004l-6-.013l-.01 4l6 .013z"/><path d="m19.978 18.002l.026-12l-2-.004l-.026 12zM3.996 10.985l12 .026l.009-4l-12-.026z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignTop(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path fill-opacity=".5" d="m13.035 7.988l.002 6l4-.002l-.002-6z"/><path d="M18 4.012L6 4.018v2l12-.006zm-6.963 15.974l-.005-12l-4 .002l.005 12z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Anchor(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 6a3.001 3.001 0 0 1-2 2.83v8.044c1.725-.444 3-2.01 3-3.874h2a6.002 6.002 0 0 1-5 5.917V20a1 1 0 1 1-2 0v-1.083A6.002 6.002 0 0 1 6 13h2a4.002 4.002 0 0 0 3 3.874V8.829A3.001 3.001 0 1 1 15 6m-3 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppleWatch(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.498 5.03c0 .048-.002.096-.005.143A3.001 3.001 0 0 1 18.5 8.005v1h1v4h-1v3a3 3 0 0 1-2.005 2.83a3 3 0 0 1-2.995 3.17h-4a3 3 0 0 1-2.995-3.17a3.001 3.001 0 0 1-2.005-2.83v-8a3 3 0 0 1 2.003-2.83a3 3 0 0 1 3.01-3.18l4 .02a3 3 0 0 1 2.984 3.015m-8-.025h6a1 1 0 0 0-.995-.99l-4-.02a1 1 0 0 0-1.005.995zm7.207 2.021l-4.22-.021H7.5a1 1 0 0 0-1 1v8a1 1 0 0 0 .999 1H15.5a1 1 0 0 0 .999-1v-8a1 1 0 0 0-.795-.979M8.5 19.005a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrangeBack(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 3h8v4h6v6h4v8h-8v-4H7v-6H3zm12 10h-2v2H9v-4h2V9h4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrangeFront(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 3h8v4h6v6h4v8h-8v-4H7v-6H3zm12 6H9v6h6z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowAlignH(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 7h-2v10h2zm-7 .757l1.414 1.415L5.586 11H10v2H5.586l1.828 1.828L6 16.243L1.757 12zm12 8.486l-1.414-1.414L18.414 13H14v-2h4.414l-1.828-1.828L18 7.757L22.243 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowAlignV(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 11v2h10v-2zm.757 7l1.415-1.414L11 18.414V14h2v4.414l1.828-1.828L16.243 18L12 22.243zm8.486-12l-1.414 1.414L13 5.586V10h-2V5.586L9.172 7.414L7.757 6L12 1.757z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowBottomLeft(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m18.243 7.172l-1.415-1.415l-9.07 9.071v-4.585h-2v8h8v-2H9.17z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowBottomLeftO(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M10 10.037H8v6h6v-2h-2.586l5.33-5.33l-1.414-1.414l-5.33 5.33z"/><path fill-rule="evenodd" d="M23 12c0 6.075-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1s11 4.925 11 11m-2 0a9 9 0 1 1-18 0a9 9 0 0 1 18 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowBottomLeftR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M10 10.037H8v6h6v-2h-2.586l5.33-5.33l-1.414-1.414l-5.33 5.33z"/><path fill-rule="evenodd" d="M1 5a4 4 0 0 1 4-4h14a4 4 0 0 1 4 4v14a4 4 0 0 1-4 4H5a4 4 0 0 1-4-4zm4-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowBottomRight(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.757 7.172l1.415-1.415l9.07 9.071v-4.585h2v8h-8v-2h4.586z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowBottomRightO(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M14.037 10.037h2v6h-6v-2h2.585l-5.329-5.33l1.414-1.414l5.33 5.33z"/><path fill-rule="evenodd" d="M23 12c0 6.075-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1s11 4.925 11 11m-2 0a9 9 0 1 1-18 0a9 9 0 0 1 18 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowBottomRightR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M14.037 10.037h2v6h-6v-2h2.585l-5.329-5.33l1.414-1.414l5.33 5.33z"/><path fill-rule="evenodd" d="M1 5a4 4 0 0 1 4-4h14a4 4 0 0 1 4 4v14a4 4 0 0 1-4 4H5a4 4 0 0 1-4-4zm4-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDown(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 3.672h2V16.5l3.243-3.243l1.414 1.415L12 20.328l-5.657-5.656l1.414-1.415L11 16.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownO(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m14.829 12.026l1.414 1.414L12 17.683L7.757 13.44l1.415-1.414L11 13.854V6.317h2v7.537z"/><path fill-rule="evenodd" d="M19.778 19.778c-4.296 4.296-11.26 4.296-15.556 0c-4.296-4.296-4.296-11.26 0-15.556c4.296-4.296 11.26-4.296 15.556 0c4.296 4.296 4.296 11.26 0 15.556m-1.414-1.414A9 9 0 1 1 5.636 5.636a9 9 0 0 1 12.728 12.728" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m14.828 12.026l1.415 1.414L12 17.683L7.757 13.44l1.415-1.414L11 13.854V6.317h2v7.537z"/><path fill-rule="evenodd" d="M1 5a4 4 0 0 1 4-4h14a4 4 0 0 1 4 4v14a4 4 0 0 1-4 4H5a4 4 0 0 1-4-4zm4-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeft(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.328 11v2H7.5l3.243 3.243l-1.415 1.414L3.672 12l5.656-5.657l1.415 1.414L7.5 11z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftO(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m11.948 14.829l-1.414 1.414L6.29 12l4.243-4.243l1.414 1.415L10.12 11h7.537v2H10.12z"/><path fill-rule="evenodd" d="M4.222 19.778c-4.296-4.296-4.296-11.26 0-15.556c4.296-4.296 11.26-4.296 15.556 0c4.296 4.296 4.296 11.26 0 15.556c-4.296 4.296-11.26 4.296-15.556 0m1.414-1.414A9 9 0 1 1 18.364 5.636A9 9 0 0 1 5.636 18.364" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m11.948 14.829l-1.414 1.414L6.29 12l4.243-4.243l1.414 1.415L10.12 11h7.537v2H10.12z"/><path fill-rule="evenodd" d="M23 19a4 4 0 0 1-4 4H5a4 4 0 0 1-4-4V5a4 4 0 0 1 4-4h14a4 4 0 0 1 4 4zm-4 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongDown(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.013 19.162l1.812-1.822l1.418 1.41l-4.231 4.255l-4.255-4.231l1.41-1.418l1.846 1.834L10.998.997l2-.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongDownC(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11 6.85a3.001 3.001 0 1 1 2 0l.012 12.288l1.812-1.823l1.419 1.41l-4.231 4.255l-4.255-4.23l1.41-1.42l1.845 1.835zm.998-1.83a1 1 0 1 1 0-2a1 1 0 0 1 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongDownE(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.998 1.02h-6v6h2.001l.013 12.145l-1.844-1.834l-1.41 1.418l4.254 4.23l4.23-4.254l-1.417-1.41l-1.813 1.823l-.013-12.118h1.999zm-4 2h2v2h-2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongDownL(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.998.972v2h3l-1 .001l.014 16.24l-1.844-1.834l-1.41 1.418l4.254 4.23l4.23-4.254l-1.417-1.41l-1.813 1.823l-.014-16.214h2v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongDownR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.242 4.641L11.999.4L7.756 4.64L11 7.886l.013 11.9l-1.845-1.834l-1.41 1.418l4.255 4.231l4.23-4.254l-1.417-1.41l-1.813 1.822L13 7.883zm-5.657 0l1.414-1.414l1.414 1.414L12 6.056z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongLeft(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1.027 11.993l4.235 4.25L6.68 14.83l-1.821-1.828L22.974 13v-2l-18.12.002L6.69 9.174L5.277 7.757z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongLeftC(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m5.27 7.757l-4.25 4.236l4.236 4.25l1.416-1.412l-1.814-1.82l12.288.042a3.001 3.001 0 0 0 5.834-.975a3 3 0 0 0-5.825-1.025L4.839 11.01l1.843-1.836zm13.71 4.303a1 1 0 1 1 2 .009a1 1 0 0 1-2-.01" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongLeftE(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m5.263 7.757l-4.25 4.236l4.236 4.25l1.417-1.412l-1.815-1.82l12.117.04l-.008 2l6 .027l.026-6l-6-.027l-.009 2l-12.144-.04l1.842-1.837zm15.715 3.312l-.01 2l-2-.01l.01-2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongLeftL(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.208 7.757L.97 12.003l4.246 4.24l1.413-1.416l-1.819-1.815l16.214.018l-.004 2l2 .004l.012-6l-2-.004l-.006 2.989l.001-.99l-16.24-.018l1.838-1.84z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongLeftR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m4.649 7.725l-4.25 4.236l4.235 4.25l1.417-1.412l-1.814-1.82l11.866.04l3.255 3.256l4.243-4.243L19.36 7.79l-3.23 3.23l-11.911-.04l1.843-1.837zm13.295 4.307l1.415-1.414l1.414 1.414l-1.415 1.414z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongRight(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m23.068 11.993l-4.25-4.236l-1.412 1.417l1.835 1.83L.932 11v2l18.305.002l-1.821 1.828l1.416 1.412z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongRightC(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m18.73 7.757l4.25 4.236l-4.236 4.25l-1.416-1.412l1.814-1.82l-12.288.042a3.001 3.001 0 1 1-.009-2l12.316-.043l-1.843-1.836zM5.02 12.06a1 1 0 1 0-2 .009a1 1 0 0 0 2-.01" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongRightE(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m22.987 11.993l-4.236 4.25l-1.417-1.412l1.815-1.82l-12.118.04l.01 2l-6 .027l-.028-6l6-.027l.01 2l12.144-.04l-1.842-1.837l1.412-1.417zm-19.965-.924l.01 2l2-.01l-.01-2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongRightL(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m18.916 7.757l4.25 4.236l-4.236 4.25l-1.416-1.412l1.819-1.825l-16.5.022v2.002h-2v-6h2v1.999l16.51-.023l-1.838-1.832z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongRightR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m19.351 7.725l4.25 4.236l-4.235 4.25l-1.417-1.412l1.814-1.82l-11.866.04l-3.255 3.256l-4.243-4.243L4.642 7.79l3.229 3.23l11.911-.04l-1.843-1.837zm-14.71 5.721l1.415-1.414l-1.414-1.414l-1.415 1.414z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongUp(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12.032 1.017l-4.274 4.21L9.16 6.653l1.859-1.83l-.056 18.155l2 .006l.056-18.113l1.798 1.825l1.425-1.403z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongUpC(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m7.758 5.23l4.274-4.21l4.21 4.275l-1.424 1.403l-1.804-1.83l-.071 12.288a3.001 3.001 0 0 1-1.029 5.824a3 3 0 0 1-.971-5.835l.071-12.315l-1.853 1.826zm4.175 13.75a1 1 0 1 0-.01 2a1 1 0 0 0 .01-2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongUpE(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m12.032 1.013l4.21 4.275l-1.424 1.403l-1.804-1.83l-.07 12.116l1.999.01l-.029 6l-6-.029l.029-6l2 .01l.071-12.145L9.161 6.65L7.758 5.224zm-1.108 19.955l2 .01l.01-2l-2-.01z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongUpL(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12.032 1.019l4.21 4.274l-1.424 1.404l-1.799-1.826l-.051 16.11h1.996v2h-6v-2h2.004l.051-16.157l-1.858 1.83l-1.403-1.425z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongUpR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.793 4.61L12.068.398l4.21 4.275l-1.424 1.403l-1.804-1.831l-.07 11.886l3.227 3.226l-4.243 4.243l-4.242-4.243l3.259-3.26l.07-11.89l-1.854 1.826zm4.171 16.163l1.414-1.415l-1.414-1.414l-1.414 1.414z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRight(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.038 6.343l-1.411 1.418l3.27 3.255l-13.605.013l.002 2l13.568-.013l-3.215 3.23l1.417 1.41l5.644-5.67z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightO(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m12.052 14.829l1.414 1.414L17.71 12l-4.243-4.243l-1.414 1.415L13.88 11H6.343v2h7.537z"/><path fill-rule="evenodd" d="M19.778 19.778c4.296-4.296 4.296-11.26 0-15.556c-4.296-4.296-11.26-4.296-15.556 0c-4.296 4.296-4.296 11.26 0 15.556c4.296 4.296 11.26 4.296 15.556 0m-1.414-1.414A9 9 0 1 0 5.636 5.636a9 9 0 0 0 12.728 12.728" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m12.052 14.829l1.414 1.414L17.71 12l-4.243-4.243l-1.414 1.415L13.88 11H6.343v2h7.537z"/><path fill-rule="evenodd" d="M1 19a4 4 0 0 0 4 4h14a4 4 0 0 0 4-4V5a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4zm4 2h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowTopLeft(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.475 5.495l.004 2l-4.557.01l9.603 9.585l-1.412 1.415l-9.63-9.61l.01 4.614l-2 .004l-.018-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowTopLeftO(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M10 13.963H8v-6h6v2h-2.586l5.33 5.33l-1.414 1.414l-5.33-5.33z"/><path fill-rule="evenodd" d="M23 12c0-6.075-4.925-11-11-11S1 5.925 1 12s4.925 11 11 11s11-4.925 11-11m-2 0a9 9 0 1 0-18 0a9 9 0 0 0 18 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowTopLeftR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M10 13.963H8v-6h6v2h-2.586l5.33 5.33l-1.414 1.414l-5.33-5.33z"/><path fill-rule="evenodd" d="M1 19a4 4 0 0 0 4 4h14a4 4 0 0 0 4-4V5a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4zm4 2h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowTopRight(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.525 5.495l-.004 2l4.557.01l-9.603 9.585l1.413 1.415l9.63-9.61l-.012 4.614l2 .004l.02-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowTopRightO(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M14 13.963h2v-6h-6v2h2.586l-5.33 5.33l1.414 1.414l5.33-5.33z"/><path fill-rule="evenodd" d="M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m2 0a9 9 0 1 1 18 0a9 9 0 0 1-18 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowTopRightR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M14 13.963h2v-6h-6v2h2.586l-5.33 5.33l1.414 1.414l5.33-5.33z"/><path fill-rule="evenodd" d="M23 19a4 4 0 0 1-4 4H5a4 4 0 0 1-4-4V5a4 4 0 0 1 4-4h14a4 4 0 0 1 4 4zm-4 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUp(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m17.657 8.962l-1.418 1.411l-3.255-3.27l-.013 13.605l-2-.002l.013-13.568l-3.23 3.215l-1.41-1.417l5.67-5.644z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpO(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m14.829 11.948l1.414-1.414L12 6.29l-4.243 4.243l1.415 1.414L11 10.12v7.537h2V10.12z"/><path fill-rule="evenodd" d="M19.778 4.222c-4.296-4.296-11.26-4.296-15.556 0c-4.296 4.296-4.296 11.26 0 15.556c4.296 4.296 11.26 4.296 15.556 0c4.296-4.296 4.296-11.26 0-15.556m-1.414 1.414A9 9 0 1 0 5.636 18.364A9 9 0 0 0 18.364 5.636" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m14.854 11.974l1.415-1.414l-4.243-4.243l-4.243 4.243l1.414 1.414l1.829-1.828v7.537h2v-7.537z"/><path fill-rule="evenodd" d="M1 19a4 4 0 0 0 4 4h14a4 4 0 0 0 4-4V5a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4zm4 2h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsBreakeH(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.243 7h2v4h.005v2h-.005v4h-2v-4H4.828l1.829 1.828l-1.414 1.415L1 12l4.243-4.243l1.414 1.415L4.828 11h4.415zm6.01 0h-2v4h-.005v2h.005v4h2v-4h4.414l-1.829 1.829l1.415 1.414L23.495 12l-4.242-4.243l-1.415 1.415L19.668 11h-4.414z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsBreakeV(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16.243 5.243l-1.414 1.414L13 4.828v4.415h4v2H7v-2h4V4.828L9.172 6.657L7.757 5.243L12 1zM7 15.253v-2h10v2h-4v4.414l1.828-1.829l1.415 1.415L12 23.495l-4.243-4.242l1.415-1.415L11 19.668v-4.414z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsExchange(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.993 12.984a1 1 0 0 0-.531 1.848L7.15 17.52a1 1 0 1 0 1.414-1.415l-1.121-1.12h7.55a1 1 0 0 0 0-2H5.015zm14.014-1.968a1 1 0 0 0 .531-1.848L16.85 6.48a1 1 0 0 0-1.414 1.415l1.121 1.12h-7.55a1 1 0 0 0 0 2h9.978z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsExchangeAlt(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.993 11.016a1 1 0 0 1-.531-1.848L7.15 6.48a1 1 0 0 1 1.414 1.415l-1.121 1.12h7.55a1 1 0 0 1 0 2H5.015zm14.014 1.968a1 1 0 0 1 .531 1.848L16.85 17.52a1 1 0 1 1-1.414-1.415l1.121-1.12h-7.55a1 1 0 1 1 0-2h9.978z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsExchangeAltV(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.984 4.993a1 1 0 0 1 1.848-.531L17.52 7.15a1 1 0 1 1-1.415 1.414l-1.12-1.121v7.55a1 1 0 0 1-2 0V5.015zm-1.968 14.014a1 1 0 0 1-1.848.531L6.48 16.85a1 1 0 0 1 1.415-1.414l1.12 1.121v-7.55a1 1 0 0 1 2 0v9.978z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsExchangeV(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.984 15a1 1 0 0 0 1.848.53l2.688-2.687a1 1 0 0 0-1.415-1.414l-1.12 1.12V5a1 1 0 0 0-2 0v9.978zm-1.968-6a1 1 0 0 0-1.848-.53l-2.687 2.687a1 1 0 1 0 1.414 1.414l1.121-1.12V19a1 1 0 1 0 2 0V9.022z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsExpandDownLeft(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M13 5a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2zm2 0h4v4h-4z" clip-rule="evenodd"/><path d="M5 13H3v8h8v-2H6.414l5.364-5.364a1 1 0 0 0-1.414-1.414L5 17.586z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsExpandDownRight(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M11 5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2zM9 5H5v4h4z" clip-rule="evenodd"/><path d="M19 13h2v8h-8v-2h4.586l-5.364-5.364a1 1 0 0 1 1.414-1.414L19 17.586z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsExpandLeft(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.1 4.1v-2h-8v8h2V5.516l5.779 5.778l1.414-1.414L5.515 4.1zm9.8 9.8h2v8h-8v-2h4.585l-5.778-5.779l1.414-1.414l5.778 5.778z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsExpandLeftAlt(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.1 2.1v2H5.516l5.778 5.779l-1.414 1.414L4.1 5.515V10.1h-2v-8zm11.8 11.8h-2v4.585l-5.779-5.778l-1.414 1.414l5.778 5.778H13.9v2h8zm-5.657-4.728l-1.415-1.415l-7.07 7.072l1.414 1.414z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsExpandRight(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.9 4.1v-2h8v8h-2V5.516l-5.779 5.778l-1.414-1.414l5.778-5.78zm-9.8 9.8h-2v8h8v-2H5.516l5.778-5.779l-1.414-1.414l-5.78 5.778z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsExpandRightAlt(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.9 2.1v2h4.585l-5.778 5.78l1.414 1.414L19.9 5.515V10.1h2v-8zM5.515 19.9H10.1v2h-8v-8h2v4.585l5.778-5.778l1.414 1.414zM9.172 7.757L7.757 9.172l7.071 7.07l1.415-1.414z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsExpandUpLeft(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M5 11H3V3h8v2H6.414l5.364 5.364a1 1 0 0 1-1.414 1.414L5 6.414z"/><path fill-rule="evenodd" d="M19 13a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2zm0 2v4h-4v-4z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsExpandUpRight(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M13 5V3h8v8h-2V6.414l-5.364 5.364a1 1 0 0 1-1.414-1.414L17.586 5z"/><path fill-rule="evenodd" d="M5 13a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2zm0 2v4h4v-4z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsH(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.657 9.172L4.243 7.757L0 12l4.243 4.243l1.414-1.415L3.829 13H10v-2H3.83zM14 11v2h6.172l-1.829 1.828l1.414 1.415L24 12l-4.243-4.243l-1.414 1.415L20.172 11z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsHalt(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m4.243 7.757l1.414 1.415L3.828 11h16.344l-1.829-1.828l1.414-1.415L24 12l-4.243 4.243l-1.414-1.415L20.171 13H3.828l1.829 1.828l-1.414 1.415L0 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsMergeAltH(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.503 6h2v5h4.172L5.846 9.172l1.415-1.415L11.503 12l-4.242 4.243l-1.415-1.415L7.675 13H3.503v5h-2zm18.994 0h2v12h-2v-5h-4.172l1.829 1.829l-1.415 1.414L12.497 12l4.242-4.243l1.415 1.415L16.325 11h4.172z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsMergeAltV(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 1.503v2h-5v4.172l1.829-1.829l1.414 1.415L12 11.503L7.757 7.261l1.415-1.415L11 7.675V3.503H6v-2zm0 18.994v2H6v-2h5v-4.172l-1.828 1.829l-1.415-1.415L12 12.497l4.243 4.242l-1.415 1.415L13 16.325v4.172z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsScrollH(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.305 12l2.825-2.825l-1.414-1.414l-2.825 2.825l-.004-.004l-1.414 1.414l.004.004l-.004.004l1.414 1.414l.004-.004l2.825 2.825l1.414-1.414zm-5.195-1.414l.003-.004l1.414 1.414l-.004.004l.004.004l-1.414 1.414l-.004-.004l-2.825 2.825l-1.414-1.414L8.695 12L5.87 9.175l1.414-1.414z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsScrollV(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.414 10.11l.004.003l-1.414 1.414l-.004-.004l-.004.004l-1.414-1.414l.004-.004L7.76 7.284L9.175 5.87L12 8.695l2.825-2.825l1.414 1.414zM12 15.305l2.825 2.825l1.414-1.414l-2.825-2.825l.004-.004l-1.414-1.414l-.004.004l-.004-.004l-1.414 1.414l.004.004l-2.825 2.825l1.414 1.414z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsShrinkH(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 7h2v10H1zm7.448.757l1.414 1.415L8.033 11h7.933l-1.828-1.828l1.414-1.415L19.795 12l-4.243 4.243l-1.414-1.415L15.966 13H8.034l1.828 1.828l-1.414 1.415L4.205 12zM23 7h-2v10h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsShrinkV(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 1v2H7V1zm-.757 7.448l-1.414 1.414L13 8.033v7.934l1.829-1.829l1.414 1.414L12 19.795l-4.243-4.243l1.415-1.414L11 15.966V8.034L9.172 9.862L7.757 8.448L12 4.205zM17 23v-2H7v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsV(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7.757 5.04l1.415 1.415L11 4.627V10h2V4.627l1.828 1.828l1.415-1.414L12 .798zm8.486 13.92l-1.415-1.415L13 19.373V14h-2v5.374l-1.828-1.829l-1.415 1.414L12 23.202z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsValt(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.172 6.455L7.757 5.041L12 .798l4.243 4.243l-1.415 1.414L13 4.627v14.746l1.828-1.828l1.415 1.414L12 23.202l-4.243-4.243l1.415-1.414L11 19.373V4.627z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Assign(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 6h4V4H4v6h2zm4 12H6v-4H4v6h6zm4-12h4v4h2V4h-6zm0 12h4v-4h2v6h-6zm-2-9.5a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Asterisk(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 6h2v4.079l3.341-2.34l1.147 1.639L13.743 12l3.745 2.622l-1.147 1.639L13 13.92V18h-2v-4.079l-3.341 2.34l-1.148-1.639L10.257 12L6.51 9.378l1.15-1.639L11 10.08z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Atlasian(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M8.507 11.556c-.317-.452-.725-.397-.911.122L5 18.908h5.178c.52-2.058.07-4.865-1.097-6.533z" opacity=".8"/><path d="M12.875 7.126c-1.268 1.81-1.676 4.958-.912 7.03l1.75 4.751h5.251l-4.597-12.48c-.19-.519-.602-.572-.919-.12z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Attachment(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 0a5 5 0 0 1 5 5v12a7 7 0 1 1-14 0V9h2v8a5 5 0 0 0 10 0V5a3 3 0 1 0-6 0v12a1 1 0 1 0 2 0V6h2v11a3 3 0 1 1-6 0V5a5 5 0 0 1 5-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Attribution(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 8a2 2 0 0 0 1.732-1H14a2 2 0 1 1 0 4h-4a4 4 0 0 0 0 8h6.268A2 2 0 0 0 20 18a2 2 0 0 0-3.732-1H10a2 2 0 1 1 0-4h4a4 4 0 0 0 0-8H7.732A2 2 0 1 0 6 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Awards(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M19 9a6.992 6.992 0 0 1-3 5.745V22h-2.586L12 20.586L10.586 22H8v-7.255A7 7 0 1 1 19 9m-2 0A5 5 0 1 1 7 9a5 5 0 0 1 10 0m-7 10.757l2-2l2 2V15.71a7.001 7.001 0 0 1-2 .29a7.001 7.001 0 0 1-2-.29z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Backspace(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m17.743 8.464l1.414 1.415L17.036 12l2.121 2.121l-1.414 1.415l-2.122-2.122l-2.121 2.122l-1.414-1.415L14.207 12l-2.121-2.121L13.5 8.465l2.121 2.12z"/><path fill-rule="evenodd" d="m8.586 19l-6.293-6.293a1 1 0 0 1 0-1.414L8.586 5h14v14zm.828-12l-5 5l5 5h11.172V7z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BandAid(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M11.939 9.765a1 1 0 1 1-1.813-.845a1 1 0 0 1 1.813.845M8.92 13.874a1 1 0 1 0 .845-1.813a1 1 0 0 0-.846 1.813m4.955 1.206a1 1 0 1 1-1.813-.845a1 1 0 0 1 1.813.846m.361-3.142a1 1 0 1 0 .845-1.813a1 1 0 0 0-.845 1.813"/><path fill-rule="evenodd" d="M17.071 1.124a6 6 0 0 0-7.973 2.902L4.026 14.902a6 6 0 0 0 10.876 5.072l5.072-10.876a6 6 0 0 0-2.903-7.974m-3.136 16.192l3.38-7.25l-7.25-3.382l-3.38 7.25zm-.846 1.812l-7.25-3.38a4 4 0 1 0 7.25 3.38m3.137-16.191a4 4 0 0 1 1.935 5.316l-7.25-3.381a4 4 0 0 1 5.315-1.935" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Battery(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M6 15a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1h6v6z"/><path fill-rule="evenodd" d="M18 6H5a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3h13a3 3 0 0 0 3-3a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1a3 3 0 0 0-3-3m0 2H5a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h13a1 1 0 0 0 1-1V9a1 1 0 0 0-1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryEmpty(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M18 6H5a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3h13a3 3 0 0 0 3-3a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1a3 3 0 0 0-3-3m0 2H5a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h13a1 1 0 0 0 1-1V9a1 1 0 0 0-1-1" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryFull(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M6 15a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1z"/><path fill-rule="evenodd" d="M18 6H5a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3h13a3 3 0 0 0 3-3a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1a3 3 0 0 0-3-3m0 2H5a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h13a1 1 0 0 0 1-1V9a1 1 0 0 0-1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bee(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M17.951 15.571a5.993 5.993 0 0 1-2.27 4.064a4.016 4.016 0 0 1-1.756 1.96a2 2 0 0 1-3.874 0a4.016 4.016 0 0 1-1.756-1.96a5.993 5.993 0 0 1-2.269-4.047a3.001 3.001 0 0 1-4.11-4.32L6.01 6.39a6 6 0 0 1 11.953-.033l4.12 4.91a3 3 0 0 1-4.132 4.304m-2.326-2.665l-1.678-2h-3.894l-1.678 2zm2.363-.296l1.032 1.229a1 1 0 1 0 1.532-1.286l-2.564-3.055zm-2-3.704v-2a4 4 0 0 0-8 0v2zM4.98 13.839l1.007-1.2V9.527l-2.54 3.027a1 1 0 1 0 1.533 1.285m7.007 5.067a4 4 0 0 1-4-4h8a4 4 0 0 1-4 4" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bell(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14 3v.29c2.892.86 5 3.539 5 6.71v7h1v2H4v-2h1v-7a7.003 7.003 0 0 1 5-6.71V3a2 2 0 1 1 4 0M7 17h10v-7a5 5 0 0 0-10 0zm7 4v-1h-4v1a2 2 0 1 0 4 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bitbucket(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.583 4.635c-.552 0-.915.44-.811.982l2.456 12.766c.104.542.637.982 1.189.982h9.166c.552 0 1.085-.44 1.189-.982l2.456-12.766c.104-.542-.259-.982-.811-.982zm8.962 9.73l.91-4.73h-4.91l.91 4.73z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Block(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M8.465 14.121a1 1 0 1 0 1.414 1.415l5.657-5.657a1 1 0 1 0-1.415-1.415z"/><path fill-rule="evenodd" d="M6.343 17.657A8 8 0 1 0 17.657 6.343A8 8 0 0 0 6.343 17.657m9.9-1.414a6 6 0 1 1-8.486-8.485a6 6 0 0 1 8.486 8.485" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bmw(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0M5 12a7 7 0 0 0 7 7v-7h7a7 7 0 0 0-7-7v7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Board(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6 4a4 4 0 0 0-4 4v8a4 4 0 0 0 4 4h12a4 4 0 0 0 4-4V8a4 4 0 0 0-4-4zm8 2h-4v12h4zm2 0v12h2a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2zM6 18h2V6H6a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bolt(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9 21.5l8.5-8.5l-4.5-3l2-7.5L6.5 11l4.5 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookmark(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M19 20h-1.828l-4.465-4.465a1 1 0 0 0-1.414 0L6.828 20H5V7a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3zM17 7a1 1 0 0 0-1-1H8a1 1 0 0 0-1 1v10l2.879-2.879a3 3 0 0 1 4.242 0L17 17z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderAll(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="3" d="M6.5 6.5h11v11h-11z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderBottom(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-opacity=".3" d="M8 8h8v7h3V5H5v10h3z"/><path d="M5 17h14v3H5z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderLeft(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-opacity=".3" d="M16 8v8H9v3h10V5H9v3z"/><path d="M7 5v14H4V5z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderRight(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-opacity=".3" d="M8 16V8h7V5H5v14h10v-3z"/><path d="M17 19V5h3v14z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderStyleDashed(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 11h4v2H4zm6 0h4v2h-4zm10 0h-4v2h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderStyleDotted(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 11H1v2h2zm4 0H5v2h2zm2 0h2v2H9zm6 0h-2v2h2zm2 0h2v2h-2zm6 0h-2v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderStyleSolid(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 11h20v2H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderTop(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-opacity=".3" d="M8 16h8V9h3v10H5V9h3z"/><path d="M5 7h14V4H5z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bot(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M14.125 13h-4v2h4z"/><path fill-rule="evenodd" d="M8.125 13a2 2 0 1 0 0-4a2 2 0 0 0 0 4m0-1.5a.5.5 0 1 0 0-1a.5.5 0 0 0 0 1m10-.5a2 2 0 1 1-4 0a2 2 0 0 1 4 0m-1.5 0a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0" clip-rule="evenodd"/><path fill-rule="evenodd" d="M2.749 14.666A6 6 0 0 0 8.125 18h8c2.44 0 4.54-1.456 5.478-3.547A2.997 2.997 0 0 0 22.875 12c0-1.013-.503-1.91-1.272-2.452A6.001 6.001 0 0 0 16.125 6h-8A6 6 0 0 0 2.75 9.334a3 3 0 0 0 0 5.332M8.125 8h8c1.384 0 2.603.702 3.322 1.77c.276.69.428 1.442.428 2.23s-.152 1.54-.428 2.23A3.996 3.996 0 0 1 16.125 16h-8a4 4 0 0 1 0-8" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bowl(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M20.547 3.672a1 1 0 0 0-1.414 0l-5.364 5.364H2v2h.008c.218 5.33 4.608 9.585 9.992 9.585c5.384 0 9.774-4.255 9.992-9.585H22v-2h-5.403l3.95-3.95a1 1 0 0 0 0-1.414m-6.37 7.364h5.813a8 8 0 0 1-15.98 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Box(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M10 12a1 1 0 1 0 0 2h4a1 1 0 0 0 0-2z"/><path fill-rule="evenodd" d="M4 2a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h16a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3zm16 2H4a1 1 0 0 0-1 1v3h18V5a1 1 0 0 0-1-1M3 19v-9h18v9a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Boy(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M9 14a1 1 0 1 0 0-2a1 1 0 0 0 0 2m7-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/><path fill-rule="evenodd" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10m0-2a8 8 0 0 0 7.634-10.4c-.835.226-1.713.346-2.619.346a9.996 9.996 0 0 1-8.692-5.053A8 8 0 0 0 12 20" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Brackets(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M11 7v2H9v6h2v2H7V7zm4 8h-2v2h4V7h-4v2h2z"/><path fill-rule="evenodd" d="M3 6a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm3-1h12a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Briefcase(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M14 11h-4v2h4z"/><path fill-rule="evenodd" d="M7 5V4a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v1h3a3 3 0 0 1 3 3v10a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3V8a3 3 0 0 1 3-3zm2-1h6v1H9zM4 7a1 1 0 0 0-1 1v6h18V8a1 1 0 0 0-1-1zM3 18v-2h18v2a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Browse(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M14.364 13.121c.924.924 1.12 2.3.586 3.415l1.535 1.535a1 1 0 0 1-1.414 1.414l-1.535-1.535a3.001 3.001 0 0 1-3.415-4.829a3 3 0 0 1 4.243 0M12.95 15.95a1 1 0 1 0-1.414-1.414a1 1 0 0 0 1.414 1.414" clip-rule="evenodd"/><path d="M8 5h8v2H8zm8 4H8v2h8z"/><path fill-rule="evenodd" d="M4 4a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3v16a3 3 0 0 1-3 3H7a3 3 0 0 1-3-3zm3-1h10a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Browser(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M4 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2m4-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0m2 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path fill-rule="evenodd" d="M3 3a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h18a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3zm18 2H3a1 1 0 0 0-1 1v3h20V6a1 1 0 0 0-1-1M2 18v-7h20v7a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Brush(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 11h3a1 1 0 0 1 1 1v6a3 3 0 0 1-3 3H8a3 3 0 0 1-3-3v-6a1 1 0 0 1 1-1h3V6a3 3 0 1 1 6 0zm-2-5a1 1 0 1 0-2 0v7H7v5a1 1 0 0 0 1 1h1v-3h2v3h2v-3h2v3h1a1 1 0 0 0 1-1v-5h-4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bulb(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M4 9a7.997 7.997 0 0 0 4 6.93V16a4 4 0 1 0 8 0v-.07A8 8 0 1 0 4 9m12 4.472a6 6 0 1 0-8 0h2V16a2 2 0 1 0 4 0v-2.53z" clip-rule="evenodd"/><path d="M10 21.006V21c.588.34 1.271.535 2 .535c.729 0 1.412-.195 2-.535v.006a2 2 0 1 1-4 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CplusPlus(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M12.207 16.278a6 6 0 1 1 .071-8.485l1.414-1.414a8 8 0 1 0-.071 11.314z"/><path d="M15 9h-2v2h-2v2h2v2h2v-2h2v-2h-2zm5 0h2v2h2v2h-2v2h-2v-2h-2v-2h2z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calculator(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M17 5H7v2h10zM7 9h2v2H7zm2 4H7v2h2zm-2 4h2v2H7zm6-8h-2v2h2zm-2 4h2v2h-2zm2 4h-2v2h2zm2-8h2v2h-2zm2 4h-2v6h2z"/><path fill-rule="evenodd" d="M3 3a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v18a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2zm2 0h14v18H5z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M8 9a1 1 0 1 0 0 2h8a1 1 0 1 0 0-2z"/><path fill-rule="evenodd" d="M6 3a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3zM5 18V7h14v11a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarDates(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M8 13a1 1 0 1 1 0-2a1 1 0 0 1 0 2m0 4a1 1 0 1 1 0-2a1 1 0 0 1 0 2m3-1a1 1 0 1 0 2 0a1 1 0 0 0-2 0m5 1a1 1 0 1 1 0-2a1 1 0 0 1 0 2m-5-5a1 1 0 1 0 2 0a1 1 0 0 0-2 0m5 1a1 1 0 1 1 0-2a1 1 0 0 1 0 2M8 7a1 1 0 0 0 0 2h8a1 1 0 1 0 0-2z"/><path fill-rule="evenodd" d="M6 3a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3zm12 2H6a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarDue(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M7 8a1 1 0 0 1 1-1h8a1 1 0 1 1 0 2H8a1 1 0 0 1-1-1m5 8a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/><path fill-rule="evenodd" d="M6 3a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3zm12 2H6a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarNext(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m11.725 16.546l4.5-2.598l-4.5-2.598v1.598h-3.95v2h3.95zM8 7a1 1 0 0 0 0 2h8a1 1 0 1 0 0-2z"/><path fill-rule="evenodd" d="M6 3a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3zm12 2H6a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarToday(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M15 17a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/><path fill-rule="evenodd" d="M6 3a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3zM5 18V7h14v11a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarTwo(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M8 7a1 1 0 0 0 0 2h8a1 1 0 1 0 0-2z"/><path fill-rule="evenodd" d="M3 6a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm3-1h12a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calibrate(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.05 5a8.97 8.97 0 0 1 6.314 2.586l-4.243 4.243A2.99 2.99 0 0 0 12.05 11c-.855 0-1.625.357-2.172.93L5.636 7.687A8.973 8.973 0 0 1 12.05 5m0 14a3 3 0 1 0 0-6a3 3 0 0 0 0 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Camera(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 4.5v2h8v1H3a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3h12a3 3 0 0 0 2.99-2.751L24 17.5v-8l-6.01.751A3 3 0 0 0 15 7.5h-1v-2a1 1 0 0 0-1-1zm14 7.766v2.468l4 .5v-3.468zM16 10.5a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cap(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M8 18v2h8v-2z"/><path fill-rule="evenodd" d="M13.988 3.22a2 2 0 1 0-3.976 0a9.003 9.003 0 0 0-6.94 9.926A3.001 3.001 0 0 0 1 16v4a3 3 0 0 0 3 3h16a3 3 0 0 0 3-3v-4c0-1.333-.87-2.463-2.072-2.854a9.003 9.003 0 0 0-6.94-9.926M12 5a7 7 0 0 0-6.93 8h13.86A7 7 0 0 0 12 5M3 16a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Captions(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M11 8v2H8v4h3v2H6V8zm7 0v2h-3v4h3v2h-5V8z"/><path fill-rule="evenodd" d="M2 5a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1zm2 13V6h16v12z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CardClubs(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M12 11a2 2 0 1 0 0-4a2 2 0 0 0 0 4m-1 2a2 2 0 1 1-4 0a2 2 0 0 1 4 0m4 2a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/><path fill-rule="evenodd" d="M3 4a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v16a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm3-1h12a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CardDiamonds(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M12 7.757L7.757 12L12 16.243L16.243 12z"/><path fill-rule="evenodd" d="M3 4a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v16a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm3-1h12a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CardHearts(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M9.146 12.293a2 2 0 0 1 2.829-2.829L12 9.49l.025-.026a2 2 0 1 1 2.829 2.829l-2.829 2.828l-.025-.025l-.025.025z"/><path fill-rule="evenodd" d="M3 4a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v16a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm3-1h12a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CardSpades(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M9.146 11.707a2 2 0 0 0 2.829 2.829L12 14.51l.025.026a2 2 0 1 0 2.829-2.829L12.025 8.88L12 8.904l-.025-.025z"/><path fill-rule="evenodd" d="M3 20a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V4a3 3 0 0 0-3-3H6a3 3 0 0 0-3 3zm3 1h12a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1H6a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Carousel(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M4 3a3 3 0 0 0-3 3v4a3 3 0 0 0 3 3h16a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3zm16 2H4a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1" clip-rule="evenodd"/><path d="M7 20a1 1 0 0 1 1-1h8a1 1 0 1 1 0 2H8a1 1 0 0 1-1-1m-2-5a1 1 0 1 0 0 2h14a1 1 0 1 0 0-2z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cast(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M20 6H4v2H2V6a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2h-5v-2h5zM2 13a7 7 0 0 1 7 7H7a5 5 0 0 0-5-5zm0 4a3 3 0 0 1 3 3H2z"/><path d="M2 9c6.075 0 11 4.925 11 11h-2a9 9 0 0 0-9-9z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chanel(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.072 3.173a9 9 0 0 0-4.608 2.463l2.13 2.13a5.989 5.989 0 0 1 5.701-1.571a9.002 9.002 0 0 0 0 11.61a5.987 5.987 0 0 1-5.702-1.57l-2.13 2.129A9 9 0 0 0 12 19.974a9.003 9.003 0 0 0 10.536-1.61l-2.13-2.13a5.988 5.988 0 0 1-5.701 1.571A9.012 9.012 0 0 0 16.828 12a9 9 0 0 0-2.123-5.805a5.988 5.988 0 0 1 5.702 1.57l2.13-2.129A9 9 0 0 0 12 4.026a9 9 0 0 0-5.928-.853M12 7.705a5.99 5.99 0 0 0-.806 7.622c.235.352.505.676.806.968a5.987 5.987 0 0 0 0-8.59" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chart(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M22.775 8A9 9 0 0 1 23 10h-9V1a9 9 0 0 1 8.775 7m-2.067 0A6.999 6.999 0 0 0 16 3.292V8z"/><path d="M1 14a9 9 0 0 1 11-8.775V12h6.775A9 9 0 1 1 1 14m15.804 0H10V7.196A6.804 6.804 0 1 0 16.804 14"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Check(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.586 13.414l-2.829-2.828L6.343 12l4.243 4.243l7.07-7.071l-1.413-1.415z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckO(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M10.243 16.314L6 12.07l1.414-1.414l2.829 2.828l5.656-5.657l1.415 1.415z"/><path fill-rule="evenodd" d="M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m11 9a9 9 0 1 1 0-18a9 9 0 0 1 0 18" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M10.243 16.314L6 12.07l1.414-1.414l2.829 2.828l5.656-5.657l1.415 1.415z"/><path fill-rule="evenodd" d="M1 5a4 4 0 0 1 4-4h14a4 4 0 0 1 4 4v14a4 4 0 0 1-4 4H5a4 4 0 0 1-4-4zm4-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleDown(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M7.757 5.636L6.343 7.05L12 12.707l5.657-5.657l-1.414-1.414L12 9.88z"/><path d="m6.343 12.707l1.414-1.414L12 15.536l4.243-4.243l1.414 1.414L12 18.364z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleDownO(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M7.757 8.464L9.172 7.05L12 9.88l2.828-2.829l1.415 1.415L12 12.707z"/><path d="m9.172 11.293l-1.415 1.414L12 16.95l4.243-4.243l-1.415-1.414L12 14.12z"/><path fill-rule="evenodd" d="M23 12c0 6.075-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1s11 4.925 11 11m-2 0a9 9 0 1 1-18 0a9 9 0 0 1 18 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleDownR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M7.757 8.464L9.172 7.05L12 9.88l2.828-2.829l1.415 1.415L12 12.707z"/><path d="m9.172 11.293l-1.415 1.414L12 16.95l4.243-4.243l-1.415-1.414L12 14.12z"/><path fill-rule="evenodd" d="M1 5a4 4 0 0 1 4-4h14a4 4 0 0 1 4 4v14a4 4 0 0 1-4 4H5a4 4 0 0 1-4-4zm4-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleLeft(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M18.364 7.757L16.95 6.343L11.293 12l5.657 5.657l1.414-1.414L14.12 12z"/><path d="m11.293 6.343l1.414 1.414L8.464 12l4.243 4.243l-1.414 1.414L5.636 12z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleLeftO(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m12.707 9.172l-1.414-1.415L7.05 12l4.243 4.243l1.414-1.415L9.88 12z"/><path d="m15.536 7.757l1.414 1.415L14.12 12l2.829 2.828l-1.414 1.415L11.293 12z"/><path fill-rule="evenodd" d="M23 12c0 6.075-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1s11 4.925 11 11m-2 0a9 9 0 1 1-18 0a9 9 0 0 1 18 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleLeftR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m12.707 9.172l-1.414-1.415L7.05 12l4.243 4.243l1.414-1.415L9.88 12z"/><path d="m15.536 7.757l1.414 1.415L14.12 12l2.829 2.828l-1.414 1.415L11.293 12z"/><path fill-rule="evenodd" d="M1 5a4 4 0 0 1 4-4h14a4 4 0 0 1 4 4v14a4 4 0 0 1-4 4H5a4 4 0 0 1-4-4zm4-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleRight(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M5.636 7.757L7.05 6.343L12.707 12L7.05 17.657l-1.414-1.414L9.88 12z"/><path d="m12.707 6.343l-1.414 1.414L15.536 12l-4.243 4.243l1.414 1.414L18.364 12z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleRightO(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M8.464 7.757L7.05 9.172L9.88 12l-2.83 2.828l1.415 1.415L12.707 12z"/><path d="m11.293 9.172l1.414-1.415L16.95 12l-4.243 4.243l-1.414-1.415L14.12 12z"/><path fill-rule="evenodd" d="M1 12c0 6.075 4.925 11 11 11s11-4.925 11-11S18.075 1 12 1S1 5.925 1 12m2 0a9 9 0 1 0 18 0a9 9 0 0 0-18 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleRightR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M8.464 7.757L7.05 9.172L9.88 12l-2.83 2.828l1.415 1.415L12.707 12z"/><path d="m11.293 9.172l1.414-1.415L16.95 12l-4.243 4.243l-1.414-1.415L14.12 12z"/><path fill-rule="evenodd" d="M23 5a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v14a4 4 0 0 0 4 4h14a4 4 0 0 0 4-4zm-4-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleUp(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m17.657 11.293l-1.414 1.414L12 8.464l-4.243 4.243l-1.414-1.414L12 5.636z"/><path d="m17.657 16.95l-1.414 1.414L12 14.12l-4.243 4.243l-1.414-1.414L12 11.293z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleUpO(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m14.828 12.707l1.415-1.414L12 7.05l-4.243 4.243l1.415 1.414L12 9.88z"/><path d="m14.828 16.95l1.415-1.414L12 11.293l-4.243 4.243l1.415 1.414L12 14.12z"/><path fill-rule="evenodd" d="M1 12c0 6.075 4.925 11 11 11s11-4.925 11-11S18.075 1 12 1S1 5.925 1 12m2 0a9 9 0 1 0 18 0a9 9 0 0 0-18 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleUpR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m14.828 12.481l1.415-1.414L12 6.824l-4.243 4.243l1.415 1.414L12 9.653z"/><path d="m14.828 16.724l1.415-1.414L12 11.067L7.757 15.31l1.415 1.414L12 13.896z"/><path fill-rule="evenodd" d="M23 4.774a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v14a4 4 0 0 0 4 4h14a4 4 0 0 0 4-4zm-4-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-14a2 2 0 0 0-2-2" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDown(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.343 7.757L4.93 9.172l7.07 7.07l7.071-7.07l-1.414-1.415L12 13.414z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDownO(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m7.757 10.586l1.415-1.414L12 12l2.829-2.828l1.414 1.414L12 14.828z"/><path fill-rule="evenodd" d="M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m11 9a9 9 0 1 1 0-18a9 9 0 0 1 0 18" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDownR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m7.757 10.586l1.415-1.414L12 12l2.828-2.828l1.415 1.414L12 14.828z"/><path fill-rule="evenodd" d="M1 5a4 4 0 0 1 4-4h14a4 4 0 0 1 4 4v14a4 4 0 0 1-4 4H5a4 4 0 0 1-4-4zm4-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeft(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.243 6.343L14.828 4.93L7.758 12l7.07 7.071l1.415-1.414L10.586 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeftO(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m12 7.757l1.414 1.415L10.586 12l2.828 2.829L12 16.243L7.757 12z"/><path fill-rule="evenodd" d="M12 1c6.075 0 11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1M3 12a9 9 0 1 1 18 0a9 9 0 0 1-18 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeftR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m13 7.757l1.414 1.415L11.586 12l2.828 2.828L13 16.243L8.757 12z"/><path fill-rule="evenodd" d="M19 1a4 4 0 0 1 4 4v14a4 4 0 0 1-4 4H5a4 4 0 0 1-4-4V5a4 4 0 0 1 4-4zm2 4v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRight(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.586 6.343L12 4.93L19.071 12L12 19.071l-1.414-1.414L16.243 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRightO(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M11.086 7.757L15.328 12l-4.242 4.243l-1.414-1.414L12.5 12L9.672 9.172z"/><path fill-rule="evenodd" d="M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11s11-4.925 11-11S18.075 1 12 1m9 11a9 9 0 1 0-18 0a9 9 0 0 0 18 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRightR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m16.485 12.045l-4.242-4.243l-1.415 1.415l2.829 2.828l-2.829 2.829l1.415 1.414z"/><path fill-rule="evenodd" d="M1 4a3 3 0 0 1 3-3h16a3 3 0 0 1 3 3v16a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3zm3-1h16a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUp(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m17.657 16.243l1.414-1.414l-7.07-7.072l-7.072 7.072l1.414 1.414L12 10.586z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUpO(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m14.829 14.828l1.414-1.414L12 9.172l-4.243 4.242l1.415 1.415L12 12z"/><path fill-rule="evenodd" d="M1 12c0 6.075 4.925 11 11 11s11-4.925 11-11S18.075 1 12 1S1 5.925 1 12m11-9a9 9 0 1 0 0 18a9 9 0 0 0 0-18" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUpR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m14.829 14.828l1.414-1.414L12 9.172l-4.243 4.242l1.415 1.415L12 12z"/><path fill-rule="evenodd" d="M1 19a4 4 0 0 0 4 4h14a4 4 0 0 0 4-4V5a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4zm4 2h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Circleci(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 0C5.373 0 0 5.373 0 12s5.373 12 12 12s12-5.373 12-12S18.627 0 12 0M4.144 13.517a8 8 0 1 0-.006-3l2.59-.01a5.478 5.478 0 1 1 .004 3zM9.522 12a2.478 2.478 0 1 0 4.956 0a2.478 2.478 0 0 0-4.956 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClapperBoard(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m20.17 3l-.004.005A3 3 0 0 1 23 6v12a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3V6a3 3 0 0 1 3-3zm-9.694 2h6L13.09 9h-6zM5.09 9l3.387-4H4a1 1 0 0 0-1 1v3zM3 11v7a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1v-7zm18-2V6a1 1 0 0 0-1-1h-1.524L15.09 9z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clipboard(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M8 11a1 1 0 1 0 0 2h7.96a1 1 0 1 0 0-2zm.04 4.067a1 1 0 1 0 0 2H16a1 1 0 1 0 0-2z"/><path fill-rule="evenodd" d="M5 3a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2zm2 2H5v14h14V5h-2v1a3 3 0 0 1-3 3h-4a3 3 0 0 1-3-3zm2 0v1a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1V5z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Close(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.225 4.811a1 1 0 0 0-1.414 1.414L10.586 12L4.81 17.775a1 1 0 1 0 1.414 1.414L12 13.414l5.775 5.775a1 1 0 0 0 1.414-1.414L13.414 12l5.775-5.775a1 1 0 0 0-1.414-1.414L12 10.586z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloseO(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M16.34 9.322a1 1 0 1 0-1.364-1.463l-2.926 2.728L9.322 7.66A1 1 0 0 0 7.86 9.024l2.728 2.926l-2.927 2.728a1 1 0 1 0 1.364 1.462l2.926-2.727l2.728 2.926a1 1 0 1 0 1.462-1.363l-2.727-2.926z"/><path fill-rule="evenodd" d="M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m11 9a9 9 0 1 1 0-18a9 9 0 0 1 0 18" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloseR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M16.396 7.757a1 1 0 0 1 0 1.415l-2.982 2.981l2.676 2.675a1 1 0 1 1-1.415 1.415L12 13.567l-2.675 2.676a1 1 0 0 1-1.415-1.415l2.676-2.675l-2.982-2.981A1 1 0 1 1 9.02 7.757L12 10.74l2.981-2.982a1 1 0 0 1 1.415 0"/><path fill-rule="evenodd" d="M4 1a3 3 0 0 0-3 3v16a3 3 0 0 0 3 3h16a3 3 0 0 0 3-3V4a3 3 0 0 0-3-3zm16 2H4a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloud(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.738 19.996A8 8 0 1 0 8.735 7H7.52a6.5 6.5 0 0 0 0 13h7a5.3 5.3 0 0 0 .219-.004m1.953-2.275c2.35-.769 4.29-3.04 4.29-5.721a6 6 0 0 0-12 0h-2c0-1.06.206-2.074.581-3H7.52a4.5 4.5 0 1 0 0 9h7c.55 0 1.385-.099 2.172-.279" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Code(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9.953 16.912l-1.36 1.449l-6.562-6.16L8.19 5.64l1.458 1.369l-4.79 5.104l5.094 4.781zm4.094 0l1.36 1.449l6.562-6.16L15.81 5.64l-1.458 1.369l4.79 5.104l-5.094 4.781z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeClimate(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9.495 8.11l-6.364 6.365l1.414 1.414l4.95-4.95l4.95 4.95l1.414-1.414zm5.01 0l-1.973 1.974l1.418 1.41l.555-.555l4.95 4.95l1.414-1.414z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeSlash(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.325 3.05L8.667 20.432l1.932.518l4.658-17.382zM7.612 18.36l1.36-1.448l-.001-.019l-5.094-4.78l4.79-5.105l-1.458-1.369l-6.16 6.563zm8.776 0l-1.36-1.448l.001-.019l5.094-4.78l-4.79-5.105l1.458-1.369l6.16 6.563z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Coffee(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M6 2.5a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0v-2a1 1 0 0 0-1-1"/><path fill-rule="evenodd" d="M13 21.5a6.002 6.002 0 0 0 5.917-5H19a4 4 0 0 0 0-8v-1H1v8a6 6 0 0 0 6 6zM3 9.5v6a4 4 0 0 0 4 4h6a4 4 0 0 0 4-4v-6zm18 3a2 2 0 0 1-2 2v-4a2 2 0 0 1 2 2" clip-rule="evenodd"/><path d="M9 3.5a1 1 0 1 1 2 0v2a1 1 0 1 1-2 0zm5-1a1 1 0 0 0-1 1v2a1 1 0 1 0 2 0v-2a1 1 0 0 0-1-1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Collage(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 5a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3v14a3 3 0 0 1-3 3H7a3 3 0 0 1-3-3zm9-1h4a1 1 0 0 1 1 1v8h-5zm0 11v5h4a1 1 0 0 0 1-1v-4zM11 4H7a1 1 0 0 0-1 1v3h5zM6 19v-9h5v10H7a1 1 0 0 1-1-1" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ColorBucket(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M8.203 2.004c1.261 0 2.304 1.103 2.476 2.538l8.483 8.484l-7.778 7.778a3 3 0 0 1-4.243 0L2.9 16.562a3 3 0 0 1 0-4.243l2.804-2.805V4.961c0-1.633 1.12-2.957 2.5-2.957m.5 2.957v1.553l-1 1V4.961c0-.327.224-.591.5-.591c.277 0 .5.264.5.591m0 5.914V9.342l-4.39 4.391a1 1 0 0 0 0 1.414l4.243 4.243a1 1 0 0 0 1.414 0l6.364-6.364l-5.63-5.63v3.48l-.003.128h-2.01a.698.698 0 0 0 .012-.129" clip-rule="evenodd"/><path d="M16.859 16.875a3 3 0 1 0 4.242 0l-2.121-2.121z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ColorPicker(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M20.385 2.879a3 3 0 0 0-4.243 0L14.02 5l-.707-.708a1 1 0 1 0-1.414 1.415l5.657 5.656A1 1 0 0 0 18.97 9.95l-.707-.707l2.122-2.122a3 3 0 0 0 0-4.242"/><path fill-rule="evenodd" d="M11.93 7.091L4.152 14.87a3.001 3.001 0 0 0-.587 3.415L2 19.85l1.414 1.415l1.565-1.566a3.001 3.001 0 0 0 3.415-.586l7.778-7.778zm1.414 4.243L11.93 9.92l-6.364 6.364a1 1 0 0 0 1.414 1.414z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ComedyCentral(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M10.545 19a7 7 0 1 0-4.95-11.95L3.473 4.93l-.018.018A9.969 9.969 0 0 1 10.545 2c5.522 0 10 4.477 10 10s-4.478 10-10 10a9.969 9.969 0 0 1-7.072-2.929l2.122-2.121a6.978 6.978 0 0 0 4.95 2.05"/><path d="M10.545 14c.593 0 1.125-.258 1.492-.668l2.122 2.122a5 5 0 1 1 0-6.909l-2.122 2.123A2 2 0 1 0 10.545 14"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Comment(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M17 9H7V7h10zM7 13h10v-2H7z"/><path fill-rule="evenodd" d="M2 18V2h20v16h-6v4h-2a4 4 0 0 1-4-4zm10-2v2a2 2 0 0 0 2 2v-4h6V4H4v12z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Community(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 17.456a6 6 0 1 1 0-10.912a6 6 0 1 1 0 10.912m-2-1.487a4 4 0 1 1 0-7.938A5.977 5.977 0 0 0 8.5 12a5.98 5.98 0 0 0 1.5 3.969m4-7.938a4 4 0 1 1 0 7.938A5.978 5.978 0 0 0 15.5 12A5.978 5.978 0 0 0 14 8.031m-2 .846c.915.733 1.5 1.86 1.5 3.123c0 1.263-.585 2.39-1.5 3.123A3.993 3.993 0 0 1 10.5 12c0-1.263.585-2.39 1.5-3.123" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Components(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.757 6.343L12 2.1l4.243 4.243L12 10.586zm2.829 0L12 4.93l1.414 1.414L12 7.757zM2.1 12l4.243-4.243L10.586 12l-4.243 4.243zm2.829 0l1.414-1.414L7.757 12l-1.414 1.414zm8.485 0l4.243 4.243L21.9 12l-4.243-4.243zm4.243-1.414L16.243 12l1.414 1.414L19.07 12zm-9.9 7.071L12 13.414l4.243 4.243L12 21.9zm2.829 0L12 16.243l1.414 1.414L12 19.07z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compress(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m19.095 8.43l-1.424-1.404l-4.914 4.985l4.985 4.914l1.404-1.424l-2.502-2.467l6.497.05l.016-2l-6.628-.05zM5.467 15.562l1.416 1.412l4.944-4.956l-4.956-4.943L5.459 8.49l2.591 2.585l-7.206.024l.006 2l7.097-.024z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompressLeft(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7.979 9.457l-3.57-.003l-.002 2l7 .006l.006-7l-2-.002L9.41 8.06L3.096 1.77L1.685 3.185zm11.583 5.095l-.009-2l-7 .028l.028 7l2-.008l-.014-3.601l6.343 6.26l1.405-1.424l-6.324-6.24z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompressRight(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14.567 8.03l6.343-6.26l1.405 1.423l-6.323 6.24l3.57.015l-.009 2l-7-.028l.028-7l2 .008zm-6.588 6.513l-3.57.003l-.002-2l7-.006l.006 7l-2 .002l-.003-3.602l-6.314 6.29l-1.411-1.416z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompressV(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.034 7.356L15.5 4.854l1.424 1.404l-4.913 4.985L7.025 6.33L8.43 4.905l2.604 2.566l.05-6.627l2 .015zm2.529 11.176l1.412-1.416l-4.957-4.943l-4.943 4.957l1.417 1.412l2.584-2.592l.026 7.207l2-.008l-.026-7.096z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Controller(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m14.828 6.343l1.415-1.414L12 .686L7.757 4.93l1.415 1.414L12 3.515zm-9.899 9.9l1.414-1.415L3.515 12l2.828-2.828L4.93 7.757L.686 12zm2.828 2.828L12 23.314l4.243-4.243l-1.415-1.414L12 20.485l-2.828-2.828zm9.9-9.899L20.485 12l-2.828 2.828l1.414 1.415L23.314 12L19.07 7.757z"/><path fill-rule="evenodd" d="M12 8a4 4 0 1 1 0 8a4 4 0 0 1 0-8m0 2a2 2 0 1 1 0 4a2 2 0 0 1 0-4" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copy(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M13 7H7V5h6zm0 4H7V9h6zm-6 4h6v-2H7z"/><path fill-rule="evenodd" d="M3 19V1h14v4h4v18H7v-4zm12-2V3H5v14zm2-10v12H9v2h10V7z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copyright(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m13.392 10.436l1.419-1.418a4 4 0 1 0 .176 5.798l-1.313-1.313h-.21a2 2 0 1 1-.073-3.067"/><path fill-rule="evenodd" d="M12 3a9 9 0 1 1 0 18a9 9 0 0 1 0-18m0 2a7 7 0 1 1 0 14a7 7 0 0 1 0-14" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerDoubleDownLeft(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m11.295 7.694l4.84-5.056l5.057 4.84l-1.383 1.445l-2.462-2.357l.162 6.034a4.8 4.8 0 0 1-4.67 4.927l-5.925.16l2.294 2.246l-1.4 1.43l-5-4.9l4.898-5l1.429 1.4l-2.377 2.427l6.017-.162a2.4 2.4 0 0 0 2.335-2.463l-.158-5.898l-2.212 2.31z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerDoubleDownRight(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.6 7.68L7.638 2.741L2.701 7.704l1.418 1.41L6.522 6.7l-.014 6.036a4.8 4.8 0 0 0 4.788 4.812l5.928.014l-2.238 2.303l1.433 1.394l4.88-5.019l-5.019-4.88l-1.394 1.434l2.436 2.369l-6.02-.015a2.4 2.4 0 0 1-2.394-2.406l.014-5.9l2.268 2.256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerDoubleLeftDown(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.299 7.76l-5.019 4.88l-1.394-1.434l2.436-2.368l-6.02.015a2.4 2.4 0 0 0-2.394 2.406l.014 5.9l2.268-2.256l1.41 1.418l-4.962 4.937l-4.937-4.962l1.418-1.41L6.522 17.3l-.014-6.036a4.8 4.8 0 0 1 4.788-4.812l5.928-.014l-2.238-2.303l1.433-1.394z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerDoubleLeftUp(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M7.784 9.25L6.37 7.836l4.242-4.242l4.243 4.242L13.44 9.25l-2.829-2.828z"/><path d="m13.44 13.493l1.415-1.414l-4.243-4.243L6.37 12.08l1.414 1.414l1.847-1.847v4.76a4 4 0 0 0 4 4h4v-2h-4a2 2 0 0 1-2-2v-4.723z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerDoubleRightDown(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7.694 12.705l-5.056-4.84l4.84-5.057L8.923 4.19L6.566 6.653L12.6 6.49a4.8 4.8 0 0 1 4.927 4.669l.16 5.926l2.246-2.294l1.43 1.4l-4.9 5l-5-4.898l1.4-1.429l2.427 2.378l-.162-6.018a2.4 2.4 0 0 0-2.463-2.335l-5.898.158l2.31 2.212z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerDoubleRightUp(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m16.216 9.25l1.415-1.414l-4.243-4.242l-4.243 4.242L10.56 9.25l2.828-2.828z"/><path d="M10.56 13.493L9.145 12.08l4.243-4.243l4.243 4.243l-1.415 1.414l-1.847-1.847v4.76a4 4 0 0 1-4 4h-4v-2h4a2 2 0 0 0 2-2v-4.723z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerDoubleUpLeft(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M9.25 7.784L7.836 6.369l-4.242 4.243l4.242 4.243L9.25 13.44l-2.828-2.828z"/><path d="m13.493 13.44l-1.414 1.415l-4.243-4.243L12.08 6.37l1.414 1.415l-1.847 1.846h4.76a4 4 0 0 1 4 4v4h-2v-4a2 2 0 0 0-2-2h-4.723z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerDoubleUpRight(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m14.75 7.784l1.414-1.415l4.242 4.243l-4.242 4.243l-1.415-1.415l2.829-2.828z"/><path d="m10.507 13.44l1.414 1.415l4.243-4.243L11.92 6.37l-1.414 1.415l1.847 1.846h-4.76a4 4 0 0 0-4 4v4h2v-4a2 2 0 0 1 2-2h4.723z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerDownLeft(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.15 13.4a2 2 0 0 0 2-2v-8h2v8a4 4 0 0 1-4 4H6.844l3.785 3.785L9.214 20.6L2.85 14.235l6.364-6.364l1.415 1.415L6.514 13.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerDownRight(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.85 13.4a2 2 0 0 1-2-2v-8h-2v8a4 4 0 0 0 4 4h10.306l-3.785 3.785l1.415 1.414l6.364-6.364l-6.364-6.364l-1.415 1.415l4.115 4.115z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerLeftDown(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.601 6.85a2 2 0 0 1 2.002-1.998l8 .007l.002-2l-8-.007a4 4 0 0 0-4.004 3.996l-.01 10.306l-3.78-3.788l-1.416 1.412l6.358 6.37l6.37-6.358l-1.413-1.415l-4.119 4.11z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerLeftUp(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14.71 10.625l1.413-1.415l-6.37-6.358l-6.358 6.37l1.416 1.413l3.78-3.789l.01 10.306a4 4 0 0 0 4.004 3.996l8-.007l-.002-2l-8 .007a2 2 0 0 1-2.002-1.998l-.01-10.636z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerRightDown(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.399 6.85a2 2 0 0 0-2.002-1.998l-8 .007l-.002-2l8-.007a4 4 0 0 1 4.004 3.996l.01 10.306l3.78-3.788l1.416 1.412l-6.358 6.37l-6.37-6.358l1.413-1.415l4.119 4.11z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerRightUp(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.29 10.625L7.877 9.21l6.37-6.358l6.358 6.37l-1.416 1.413l-3.78-3.789l-.01 10.306a4 4 0 0 1-4.004 3.996l-8-.007l.002-2l8 .007a2 2 0 0 0 2.002-1.998l.01-10.636z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerUpLeft(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.628 14.722l-1.412 1.417L2.84 9.79l6.35-6.378l1.417 1.411L6.83 8.615l10.305-.022a4 4 0 0 1 4.009 3.991l.017 8l-2 .005l-.017-8a2 2 0 0 0-2.004-1.996l-10.636.023z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerUpRight(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.372 14.722l1.412 1.417l6.377-6.35l-6.35-6.378l-1.417 1.411l3.776 3.793l-10.305-.022a4 4 0 0 0-4.009 3.991l-.017 8l2 .005l.017-8a2 2 0 0 1 2.004-1.996l10.636.023z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCard(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M4 9a1 1 0 0 1 1-1h4a1 1 0 0 1 0 2H5a1 1 0 0 1-1-1"/><path fill-rule="evenodd" d="M4 3a4 4 0 0 0-4 4v10a4 4 0 0 0 4 4h16a4 4 0 0 0 4-4V7a4 4 0 0 0-4-4zm16 2H4a2 2 0 0 0-2 2v7h20V7a2 2 0 0 0-2-2m2 11H2v1a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crop(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.932 9.009V16H15v4.009h2V16h3.932v-2H17V7.009H9.932V3h-2v4.009H4v2zm2 0V14H15V9.009z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cross(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 6a1 1 0 1 0-2 0v3H7a1 1 0 0 0 0 2h4v7a1 1 0 1 0 2 0v-7h4a1 1 0 1 0 0-2h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crowdfire(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2v2h8v8h2c0 5.523-4.477 10-10 10S2 17.523 2 12m16 0h-2V8h-4V6a6 6 0 1 0 6 6" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crown(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m2.5 6.091l4.72 4.72L12 6.031l4.781 4.78L21.5 6.092v8.877a3 3 0 0 1-3 3h-13a3 3 0 0 1-3-3zm17 4.818v4.06a1 1 0 0 1-1 1h-13a1 1 0 0 1-1-1v-4.061l2.72 2.72L12 8.848l4.781 4.78z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Danger(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M12 6a1 1 0 0 1 1 1v6a1 1 0 1 1-2 0V7a1 1 0 0 1 1-1m0 10a1 1 0 1 0 0 2a1 1 0 0 0 0-2"/><path fill-rule="evenodd" d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2M4 12a8 8 0 1 0 16 0a8 8 0 0 0-16 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DarkMode(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M12 16a4 4 0 0 0 0-8z"/><path fill-rule="evenodd" d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2m0 2v4a4 4 0 1 0 0 8v4a8 8 0 1 0 0-16" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Data(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M12 8.89a3 3 0 0 1 1 5.829v5.17h-2v-5.17a3.001 3.001 0 0 1 1-5.83m0 2a1 1 0 1 1 0 2a1 1 0 0 1 0-2" clip-rule="evenodd"/><path d="M7.05 6.94A6.978 6.978 0 0 0 5 11.89c0 2.177.994 4.122 2.554 5.406l1.423-1.424A4.992 4.992 0 0 1 7 11.89c0-1.38.56-2.63 1.464-3.535zm8.486 1.413A4.985 4.985 0 0 1 17 11.89c0 1.626-.776 3.07-1.977 3.983l1.423 1.424A6.986 6.986 0 0 0 19 11.889a6.978 6.978 0 0 0-2.05-4.95z"/><path d="M1 11.89a10.97 10.97 0 0 1 3.222-7.78l1.414 1.415A8.972 8.972 0 0 0 3 11.89a8.972 8.972 0 0 0 2.636 6.364l-1.414 1.414A10.966 10.966 0 0 1 1 11.89m18.778 7.777A10.965 10.965 0 0 0 23 11.89c0-3.038-1.231-5.788-3.222-7.778l-1.414 1.414A8.972 8.972 0 0 1 21 11.89a8.972 8.972 0 0 1-2.636 6.364z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Database(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M5 9V7h2v2zm4 0h10V7H9zm-4 6v2h2v-2zm14 2H9v-2h10z"/><path fill-rule="evenodd" d="M1 6a3 3 0 0 1 3-3h16a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3zm3-1h16a1 1 0 0 1 1 1v5H3V6a1 1 0 0 1 1-1m-1 8v5a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1v-5z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Debug(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M10 11a1 1 0 0 1 1-1h2a1 1 0 1 1 0 2h-2a1 1 0 0 1-1-1m1 3a1 1 0 1 0 0 2h2a1 1 0 1 0 0-2z"/><path fill-rule="evenodd" d="M9.094 4.75A3.986 3.986 0 0 1 8 2h2a2 2 0 1 0 4 0h2a3.986 3.986 0 0 1-1.094 2.75A6.02 6.02 0 0 1 17.659 8H19a1 1 0 1 1 0 2h-1v2h1a1 1 0 1 1 0 2h-1v2h1a1 1 0 1 1 0 2h-1.341A6.003 6.003 0 0 1 6.34 18H5a1 1 0 1 1 0-2h1v-2H5a1 1 0 1 1 0-2h1v-2H5a1 1 0 1 1 0-2h1.341a6.02 6.02 0 0 1 2.753-3.25M8 16v-6a4 4 0 1 1 8 0v6a4 4 0 0 1-8 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Designmodo(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M7 4.054a5 5 0 1 0 0 10a5 5 0 0 0 0-10m-2 5a2 2 0 1 0 4 0a2 2 0 0 0-4 0" clip-rule="evenodd"/><path d="M22 10.554h-5v-3h5zm-10.45 6.392a3.988 3.988 0 0 0 2.829-1.172l2.121 2.121a6.978 6.978 0 0 1-4.95 2.05a6.978 6.978 0 0 1-4.95-2.05l2.122-2.12a3.987 3.987 0 0 0 2.828 1.17"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Desktop(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M8 15a1 1 0 1 0 0 2a1 1 0 0 0 0-2m3 1a1 1 0 1 1 2 0a1 1 0 0 1-2 0m5-1a1 1 0 1 0 0 2a1 1 0 0 0 0-2"/><path fill-rule="evenodd" d="M4 3a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h16a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3zm16 2H4a1 1 0 0 0-1 1v1h18V6a1 1 0 0 0-1-1M3 18V9h18v9a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DetailsLess(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 9a1 1 0 0 0 0 2h18a1 1 0 1 0 0-2zm0 4a1 1 0 1 0 0 2h12a1 1 0 1 0 0-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DetailsMore(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 8a1 1 0 0 1 1-1h18a1 1 0 1 1 0 2H3a1 1 0 0 1-1-1m0 4a1 1 0 0 1 1-1h18a1 1 0 1 1 0 2H3a1 1 0 0 1-1-1m1 3a1 1 0 1 0 0 2h12a1 1 0 1 0 0-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dialpad(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.5 3h3v3h-3zm0 5h3v3h-3zm0 5v3h3v-3zm5-10h3v3h-3zm0 5v3h3V8zm0 5h3v3h-3zm0 5v3h3v-3zm5-15h3v3h-3zm0 5v3h3V8zm0 5h3v3h-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceFive(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M14.945 7.055a2 2 0 1 1 4 0a2 2 0 0 1-4 0m2 7.837a2 2 0 1 0 0 4a2 2 0 0 0 0-4m-11.89 2a2 2 0 1 1 4 0a2 2 0 0 1-4 0m2-11.837a2 2 0 1 0 0 4a2 2 0 0 0 0-4M10 12a2 2 0 1 1 4 0a2 2 0 0 1-4 0"/><path fill-rule="evenodd" d="M1 4a3 3 0 0 1 3-3h16a3 3 0 0 1 3 3v16a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3zm3-1h16a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceFour(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M16.945 5.055a2 2 0 1 0 0 4a2 2 0 0 0 0-4m-2 11.837a2 2 0 1 1 4 0a2 2 0 0 1-4 0m-7.89-2a2 2 0 1 0 0 4a2 2 0 0 0 0-4m-2-7.837a2 2 0 1 1 4 0a2 2 0 0 1-4 0"/><path fill-rule="evenodd" d="M4 1a3 3 0 0 0-3 3v16a3 3 0 0 0 3 3h16a3 3 0 0 0 3-3V4a3 3 0 0 0-3-3zm16 2H4a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceOne(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M10 12a2 2 0 1 1 4 0a2 2 0 0 1-4 0"/><path fill-rule="evenodd" d="M1 4a3 3 0 0 1 3-3h16a3 3 0 0 1 3 3v16a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3zm3-1h16a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceSix(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M16.945 5.055a2 2 0 1 0 0 4a2 2 0 0 0 0-4m-2 11.837a2 2 0 1 1 4 0a2 2 0 0 1-4 0m-7.89-2a2 2 0 1 0 0 4a2 2 0 0 0 0-4m-2-7.837a2 2 0 1 1 4 0a2 2 0 0 1-4 0m11.89 2.919a2 2 0 1 0 0 4a2 2 0 0 0 0-4m-11.89 2a2 2 0 1 1 4 0a2 2 0 0 1-4 0"/><path fill-rule="evenodd" d="M4 1a3 3 0 0 0-3 3v16a3 3 0 0 0 3 3h16a3 3 0 0 0 3-3V4a3 3 0 0 0-3-3zm16 2H4a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceThree(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M10 12a2 2 0 1 1 4 0a2 2 0 0 1-4 0m6.945 2.892a2 2 0 1 0 0 4a2 2 0 0 0 0-4M5.055 7.055a2 2 0 1 1 4 0a2 2 0 0 1-4 0"/><path fill-rule="evenodd" d="M1 4a3 3 0 0 1 3-3h16a3 3 0 0 1 3 3v16a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3zm3-1h16a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceTwo(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M17.2 14.943a2 2 0 1 0 0 4a2 2 0 0 0 0-4M5.055 7.055a2 2 0 1 1 4 0a2 2 0 0 1-4 0"/><path fill-rule="evenodd" d="M4 1a3 3 0 0 0-3 3v16a3 3 0 0 0 3 3h16a3 3 0 0 0 3-3V4a3 3 0 0 0-3-3zm16 2H4a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Digitalocean(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M12 6a6 6 0 0 0-6 6H1C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11v-5a6 6 0 0 0 0-12"/><path d="M7 18v-5h5v5zm-4 0v4h4v-4zm0 0H1v-2h2z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Disc(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M12 9a3 3 0 1 0 0 6a3 3 0 0 0 0-6m-1 3a1 1 0 1 0 2 0a1 1 0 0 0-2 0" clip-rule="evenodd"/><path d="M5 12a7 7 0 0 1 7-7v2a5 5 0 0 0-5 5zm7 5a5 5 0 0 0 5-5h2a7 7 0 0 1-7 7z"/><path fill-rule="evenodd" d="M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11s11-4.925 11-11S18.075 1 12 1M3 12a9 9 0 1 0 18 0a9 9 0 0 0-18 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DisplayFlex(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M6 17V7h2v10zM16 7v10h2V7z"/><path fill-rule="evenodd" d="M2 3h20v18H2zm2 2v14h16V5z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DisplayFullwidth(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M2 5h20V3H2zm0 16h20v-2H2z"/><path fill-rule="evenodd" d="M2 7v10h20V7zm2 2h16v6H4z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DisplayGrid(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M7 7v4h4V7zm6 0h4v4h-4zm0 6v4h4v-4zm-6 0h4v4H7z"/><path fill-rule="evenodd" d="M3 3h18v18H3zm2 2v14h14V5z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DisplaySpacing(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M3 21V3h2v18z"/><path fill-rule="evenodd" d="M7 3h10v18H7zm2 2v14h6V5z" clip-rule="evenodd"/><path d="M19 3v18h2V3z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DistributeHorizontal(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><path stroke="currentColor" stroke-opacity=".5" stroke-width="2" d="M11 9h2v6h-2z"/><path fill="currentColor" d="M5 5v14h2V5zm12 0v14h2V5z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DistributeVertical(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><path stroke="currentColor" stroke-opacity=".5" stroke-width="2" d="M9 11h6v2H9z"/><path fill="currentColor" d="M19 7H5V5h14zm0 12H5v-2h14z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DockBottom(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 20V4h20v16zM4 6h16v8H4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DockLeft(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 4h20v16H2zm6 2h12v12H8z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DockRight(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 4h20v16H2zm14 14V6H4v12z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DockWindow(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 20V4h20v16zM20 6H6v10h14z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dolby(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 4v16h24V4zm10 8a4 4 0 0 0-4-4H4v8h2a4 4 0 0 0 4-4m8 4h2V8h-2a4 4 0 0 0 0 8" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dollar(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11 19v3h2v-3h1a4 4 0 0 0 0-8h-1V7h2v2h2V5h-4V2h-2v3h-1a4 4 0 1 0 0 8h1v4H9v-2H7v4zm2-2h1a2 2 0 1 0 0-4h-1zm-2-6V7h-1a2 2 0 1 0 0 4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dribbble(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 11.955v.09c.01 2.685.984 5.144 2.592 7.048a11.066 11.066 0 0 0 3.378 2.692A10.922 10.922 0 0 0 12 23c3.69 0 6.955-1.816 8.95-4.604A10.96 10.96 0 0 0 23 12c0-3.26-1.418-6.19-3.672-8.203a10.949 10.949 0 0 0-7.663-2.792A10.944 10.944 0 0 0 4.43 4.019a11.05 11.05 0 0 0-2.76 4.188A10.976 10.976 0 0 0 1 11.955m19.481 3.064c.336-.944.519-1.96.519-3.019a8.971 8.971 0 0 0-2.581-6.309a10.93 10.93 0 0 1-3.152 3.356a11.04 11.04 0 0 1 .738 3.83a11.075 11.075 0 0 1 4.476 2.142m-4.64-.124a9.048 9.048 0 0 1 3.731 1.971a8.995 8.995 0 0 1-6.993 4.116a10.97 10.97 0 0 0 2.393-3.33c.419-.899.706-1.825.87-2.757m-1.845-2.273a9.025 9.025 0 0 0-.495-2.581A10.975 10.975 0 0 1 3.366 9.45A9.002 9.002 0 0 0 3 12a8.96 8.96 0 0 0 1.668 5.22a11.017 11.017 0 0 1 9.328-4.598M6.047 18.75a9.01 9.01 0 0 1 7.812-4.13a9.018 9.018 0 0 1-.7 2.186a8.958 8.958 0 0 1-3.485 3.89a8.98 8.98 0 0 1-3.627-1.946M12 3c1.785 0 3.448.52 4.847 1.415a8.944 8.944 0 0 1-2.479 2.816a10.941 10.941 0 0 0-4.341-4.014A9.031 9.031 0 0 1 12 3M8.806 4.846a8.958 8.958 0 0 1 3.832 3.39a8.979 8.979 0 0 1-7.439-.077a9.084 9.084 0 0 1-1.044-.573a9.044 9.044 0 0 1 3.172-3.28c.5.135.994.314 1.48.54" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Drive(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M19 11a1 1 0 1 0 0 2a1 1 0 0 0 0-2m-5 1a1 1 0 1 1 2 0a1 1 0 0 1-2 0"/><path fill-rule="evenodd" d="M2 8a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h20a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2zm20 2H2v4h20z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Drop(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="2" d="M6.343 19.52a8 8 0 0 1 0-11.313L12 2.55l5.657 5.657A8 8 0 0 1 6.343 19.521Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DropInvert(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1.136L5.636 7.5a9 9 0 0 0 7.227 15.323A9 9 0 0 0 18.364 7.5zM7.05 8.914L12 3.964v16.9a7 7 0 0 1-4.95-11.95" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DropOpacity(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.945 21.956A9 9 0 0 1 5.635 7.5L12 1.136L18.364 7.5a8.97 8.97 0 0 1 1.991 3.012a9.002 9.002 0 0 1-1.991 9.716a8.987 8.987 0 0 1-2.419 1.728M7.05 8.914L12 3.964l4.95 4.95a6.977 6.977 0 0 1 2.048 4.783H5.002A6.976 6.976 0 0 1 7.05 8.914" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Duplicate(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M19 5H7V3h14v14h-2zM9 13v-2h2v2h2v2h-2v2H9v-2H7v-2z"/><path fill-rule="evenodd" d="M3 7h14v14H3zm2 2h10v10H5z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditBlackPoint(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M8 12a4 4 0 1 1 8 0a4 4 0 0 1-8 0m4 1a1 1 0 1 1 0-2a1 1 0 0 1 0 2"/><path d="M3 12a9 9 0 1 1 18 0a9 9 0 0 1-18 0m9 7a7 7 0 1 1 0-14a7 7 0 0 1 0 14"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditContrast(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M12 18a5.978 5.978 0 0 1-4-1.528A5.985 5.985 0 0 1 6 12c0-1.777.773-3.374 2-4.472A5.978 5.978 0 0 1 12 6z"/><path fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m10 8a8 8 0 1 1 0-16a8 8 0 0 1 0 16" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditExposure(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 20a8 8 0 1 1 0-16a8 8 0 0 1 0 16M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m9-5v2H9v2h2v2h2v-2h2V9h-2V7zm-2 8v2h6v-2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditFade(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor"><path stroke-opacity=".3" stroke-width="4" d="M8 12c0-1.48.804-2.773 2-3.465v6.93A3.998 3.998 0 0 1 8 12Z"/><path stroke-opacity=".6" stroke-width="4" d="M14 15.465v-6.93c1.196.692 2 1.984 2 3.465c0 1.48-.804 2.773-2 3.465Z"/><path stroke-width="2" d="M3 12a9 9 0 1 1 18 0a9 9 0 0 1-18 0Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditFlipH(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-opacity=".5" d="M18 7a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1h-3v2h3a3 3 0 0 0 3-3V8a3 3 0 0 0-3-3h-3v2z"/><path d="M13 3h-2v18h2zM5 8a1 1 0 0 1 1-1h3V5H6a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h3v-2H6a1 1 0 0 1-1-1z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditFlipV(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-opacity=".5" d="M17 18a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1v-3H5v3a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3v-3h-2z"/><path d="M16 5a1 1 0 0 1 1 1v3h2V6a3 3 0 0 0-3-3H8a3 3 0 0 0-3 3v3h2V6a1 1 0 0 1 1-1zm5 8v-2H3v2z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditHighlight(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M12 6a5.972 5.972 0 0 0-3.306.992H12v1H7.535a5.996 5.996 0 0 0-1.203 2.034H12v1H6.079a6.042 6.042 0 0 0 .003 1.966H12v1H6.339c.263.748.67 1.429 1.189 2.008H12v1H8.682A6 6 0 1 0 12 6"/><path fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m10 8a8 8 0 1 1 0-16a8 8 0 0 1 0 16" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditMarkup(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 24c6.627 0 12-5.373 12-12S18.627 0 12 0S0 5.373 0 12s5.373 12 12 12m6.58-4.469A9.976 9.976 0 0 0 22 12c0-5.523-4.477-10-10-10S2 6.477 2 12a9.975 9.975 0 0 0 3.333 7.453L7 10.973h2.17l2.83-4.9l2.83 4.9H17zm-2.488 1.596l-.886-8.153H8.794l-.886 8.153A9.964 9.964 0 0 0 12 22a9.965 9.965 0 0 0 4.092-.873" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditMask(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M12 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/><path fill-rule="evenodd" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10m-2.97-2.57a3 3 0 1 1 5.939 0a8.026 8.026 0 0 0 4.462-4.46a3 3 0 1 1 0-5.938a8.026 8.026 0 0 0-4.462-4.463a3 3 0 1 1-5.939 0a8.026 8.026 0 0 0-4.46 4.462A3.015 3.015 0 0 1 5 9a3 3 0 1 1-.43 5.97a8.026 8.026 0 0 0 4.46 4.46" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditNoise(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M10.404 17.766a.59.59 0 0 1 .042.045a5.946 5.946 0 0 1-.733-.248a.75.75 0 0 1 .691.203m-1.968-1.061a.758.758 0 0 1-.086.074a6.035 6.035 0 0 1-1.056-1.038a.75.75 0 1 1 1.142.964m-2.065-3.057a.75.75 0 0 1 .155.835a5.953 5.953 0 0 1-.342-.972a.75.75 0 0 1 .187.137m0-2.973a.748.748 0 0 1-.262.17a5.88 5.88 0 0 1 .344-1.134a.75.75 0 0 1-.082.964m1.004-1.968a.747.747 0 0 1-.181-.293c.265-.353.568-.676.903-.963a.75.75 0 1 1-.722 1.255m1.969-1.939a.754.754 0 0 1-.082-.097a5.958 5.958 0 0 1 1.36-.5a.749.749 0 0 1-1.278.597m3.943 11.106c.31-.068.61-.16.9-.273a.75.75 0 0 0-.9.273m2.244-1.013c.404-.294.77-.638 1.088-1.023a.75.75 0 1 0-1.088 1.024m1.834-2.169a5.93 5.93 0 0 0 .461-1.26a.75.75 0 0 0-.461 1.26m.526-3.799a5.95 5.95 0 0 0-.419-1.334a.75.75 0 0 0 .419 1.335M16.756 8.36a6.04 6.04 0 0 0-.875-.92a.75.75 0 1 0 .875.919m-2.319-1.592a.75.75 0 0 0 .131-.176a5.957 5.957 0 0 0-1.403-.464a.75.75 0 0 0 1.272.64m-1.911.879a.75.75 0 1 1-1.06 1.06a.75.75 0 0 1 1.06-1.06m-2.122 1.969a.75.75 0 1 1-1.06 1.06a.75.75 0 0 1 1.06-1.06m-1.968 3.182a.75.75 0 1 0-1.061-1.061a.75.75 0 0 0 1.06 1.06m1.969.852a.75.75 0 1 1-1.06 1.06a.75.75 0 0 1 1.06-1.06m2.122-.851a.75.75 0 1 0-1.061-1.061a.75.75 0 0 0 1.06 1.06m1.912-3.181a.75.75 0 1 1-1.06 1.06a.75.75 0 0 1 1.06-1.06m2.122 3.153a.75.75 0 1 0-1.061-1.06a.75.75 0 0 0 1.06 1.06m-2.135.922a.75.75 0 1 1-1.06 1.06a.75.75 0 0 1 1.06-1.06m-2.107 3.015a.75.75 0 1 0-1.06-1.06a.75.75 0 0 0 1.06 1.06"/><path fill-rule="evenodd" d="M22 12c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2s10 4.477 10 10m-2 0a8 8 0 1 1-16 0a8 8 0 0 1 16 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditShadows(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M15.306 6.992A6 6 0 1 0 15.318 17H12v-1h4.472a6.01 6.01 0 0 0 1.19-2.008H12v-1h5.918a6.038 6.038 0 0 0 .003-1.966H12v-1h5.668a5.996 5.996 0 0 0-1.203-2.034H12v-1z"/><path fill-rule="evenodd" d="M22 12c0-5.523-4.477-10-10-10S2 6.477 2 12s4.477 10 10 10s10-4.477 10-10m-10 8a8 8 0 1 0 0-16a8 8 0 0 0 0 16" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditStraight(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 4a7 7 0 0 1 7 7H5a7 7 0 0 1 7-7m-7 9H1v-2h4zm14 0a7 7 0 1 1-14 0zm0 0v-2h4v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditUnmask(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 5a3 3 0 1 1-6 0a3 3 0 0 1 6 0m-1 7a2 2 0 1 1-4 0a2 2 0 0 1 4 0m-2 10a3 3 0 1 0 0-6a3 3 0 0 0 0 6m10-10a3 3 0 1 1-6 0a3 3 0 0 1 6 0M5 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eject(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16.95 14.395l1.414-1.415L12 6.617L5.636 12.98l1.414 1.415L12 9.445zM6 17.384h12v-2H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Enter(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M20 5H8v4H6V3h16v18H6v-6h2v4h12z"/><path d="m13.074 16.95l-1.414-1.414L14.196 13H2v-2h12.196L11.66 8.465l1.414-1.415l4.95 4.95z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Erase(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.5 12.9a2 2 0 0 0 0 2.828l3.858 3.858H4.086a1 1 0 1 0 0 2h16a1 1 0 0 0 0-2h-9.13l9.515-9.515a2 2 0 0 0 0-2.828L16.228 3a2 2 0 0 0-2.829 0zm4.326-1.498l-2.912 2.912l4.243 4.242l2.911-2.911zM9.24 9.988l4.243 4.243l5.573-5.574l-4.242-4.243z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ereader(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M16 7a1 1 0 1 0 0 2h3a1 1 0 1 0 0-2zm-1 5a1 1 0 0 1 1-1h3a1 1 0 1 1 0 2h-3a1 1 0 0 1-1-1m1 3a1 1 0 1 0 0 2h3a1 1 0 1 0 0-2z"/><path fill-rule="evenodd" d="M3 3a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h18a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3zm18 2h-8v14h8a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1M3 5h8v14H3a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ericsson(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.717 5.723a2 2 0 1 0 1.69 3.625l10.876-5.071a2 2 0 0 0-1.69-3.625zM4.75 15.381a2 2 0 0 1 .967-2.658l10.876-5.071a2 2 0 1 1 1.69 3.625L7.407 16.348a2 2 0 0 1-2.657-.967m-2 7a2 2 0 0 1 .967-2.658l10.876-5.071a2 2 0 1 1 1.69 3.625L5.407 23.348a2 2 0 0 1-2.657-.967"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ethernet(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 .5v20h7v3h2v-3h7V.5zm14 2H6v6h4v6h4v-6h4zm-12 16v-8h2v6h8v-6h2v8z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Euro(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.554 15.536a4.999 4.999 0 0 1-7.902-1.098h2.38l.696-1.876H10.05a5.047 5.047 0 0 1 0-1.125h4.287l.696-1.874h-4.38a4.998 4.998 0 0 1 7.902-1.099l1.414-1.414A7.003 7.003 0 0 0 8.454 9.562H6.032l-.696 1.875H8.04a7.095 7.095 0 0 0 0 1.126H4.728l-.696 1.874h4.422a7.003 7.003 0 0 0 11.514 2.513z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eventbrite(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.673 17.438a6.002 6.002 0 0 1-6.819-1.234l14.01-6.533a10.047 10.047 0 0 0-.663-1.897C18.867 2.768 12.917.603 7.91 2.937C2.907 5.27.742 11.22 3.076 16.227c2.334 5.005 8.284 7.17 13.289 4.836a9.974 9.974 0 0 0 5.317-6.077h-4.339a5.972 5.972 0 0 1-2.669 2.452M9.602 6.562a6.002 6.002 0 0 0-3.438 6.017l10.257-4.783a6.002 6.002 0 0 0-6.819-1.234" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Expand(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12.306 16.593l-.035 2l-6.999-.122l.123-7l2 .036l-.063 3.585l7.894-7.624l-3.532-.061l.035-2l6.999.122l-.123 7l-2-.036l.064-3.638l-7.948 7.676z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Export(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m16.95 5.968l-1.414 1.414L13 4.846v12.196h-2V4.847L8.465 7.382L7.05 5.968L12 1.018z"/><path d="M5 20.982v-10h4v-2H3v14h18v-14h-6v2h4v10z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Extension(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M13 3h8v8h-8zm2 2h4v4h-4z"/><path d="M17 21v-8h-6V7H3v14zM9 9H5v4h4zM5 19v-4h4v4zm6 0v-4h4v4z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExtensionAdd(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M16 4h2v2h2v2h-2v2h-2V8h-2V6h2z"/><path fill-rule="evenodd" d="M12 12V6H4v14h14v-8zM6 8h4v4H6zm4 6v4H6v-4zm6 0v4h-4v-4z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExtensionAlt(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 5v14h8v-6h6V5zm6 2H7v4h4zm0 6H7v4h4zm2-2h4V7h-4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExtensionRemove(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M12 11V5H4v14h14v-8zM6 7h4v4H6zm4 6v4H6v-4zm6 0v4h-4v-4z" clip-rule="evenodd"/><path d="M20 7h-6v2h6z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func External(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M15.64 7.025h-3.622v-2h7v7h-2v-3.55l-4.914 4.914l-1.414-1.414z"/><path d="M10.982 6.975h-6v12h12v-6h-2v4h-8v-8h4z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eye(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M16 12a4 4 0 1 1-8 0a4 4 0 0 1 8 0m-2 0a2 2 0 1 1-4 0a2 2 0 0 1 4 0"/><path d="M12 3c5.592 0 10.29 3.824 11.622 9c-1.332 5.176-6.03 9-11.622 9S1.71 17.176.378 12C1.71 6.824 6.408 3 12 3m0 16c-4.476 0-8.269-2.942-9.543-7C3.731 7.942 7.524 5 12 5c4.476 0 8.269 2.942 9.543 7c-1.274 4.058-5.067 7-9.543 7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeAlt(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M14 12a2 2 0 1 1-4 0a2 2 0 0 1 4 0"/><path fill-rule="evenodd" d="M12 3C6.408 3 1.71 6.824.378 12C1.71 17.176 6.408 21 12 21s10.29-3.824 11.622-9C22.29 6.824 17.592 3 12 3m4 9a4 4 0 1 1-8 0a4 4 0 0 1 8 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Facebook(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.198 21.5h4v-8.01h3.604l.396-3.98h-4V7.5a1 1 0 0 1 1-1h3v-4h-3a5 5 0 0 0-5 5v2.01h-2l-.396 3.98h2.396z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Feed(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-opacity=".5" d="M12.552 8a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2zm0 9a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2z"/><path fill-opacity=".8" d="M12.552 5a1 1 0 1 0 0 2h8a1 1 0 1 0 0-2zm0 9a1 1 0 1 0 0 2h8a1 1 0 1 0 0-2z"/><path d="M3.448 4.002a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h5a1 1 0 0 0 1-1v-5a1 1 0 0 0-1-1zm0 8.996a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h5a1 1 0 0 0 1-1v-5a1 1 0 0 0-1-1z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Figma(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 2a3 3 0 0 0 0 6h7a3 3 0 1 0 0-6zm7 7a3 3 0 1 0 0 6a3 3 0 0 0 0-6m-10 3a3 3 0 0 1 3-3h3v6h-3a3 3 0 0 1-3-3m3 4a3 3 0 1 0 3 3v-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func File(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 5a3 3 0 0 1 3-3h8a7 7 0 0 1 7 7v10a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm10-1H6a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V9h-6zm5.584 3A5.009 5.009 0 0 0 15 4.1V7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileAdd(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M10 18v-2H8v-2h2v-2h2v2h2v2h-2v2z"/><path fill-rule="evenodd" d="M6 2a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V9a7 7 0 0 0-7-7zm0 2h7v5h6v10a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1m9 .1A5.009 5.009 0 0 1 18.584 7H15z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileDocument(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M7 18h10v-2H7zm10-4H7v-2h10zM7 10h4V8H7z"/><path fill-rule="evenodd" d="M6 2a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V9a7 7 0 0 0-7-7zm0 2h7v5h6v10a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1m9 .1A5.009 5.009 0 0 1 18.584 7H15z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileRemove(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M9 15h6v-2H9z"/><path fill-rule="evenodd" d="M6 2a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V9a7 7 0 0 0-7-7zm0 2h7v5h6v10a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1m9 .1A5.009 5.009 0 0 1 18.584 7H15z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Film(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M6 7a1 1 0 1 0 0 2a1 1 0 0 0 0-2m11 1a1 1 0 1 1 2 0a1 1 0 0 1-2 0M6 11a1 1 0 1 0 0 2a1 1 0 0 0 0-2m11 1a1 1 0 1 1 2 0a1 1 0 0 1-2 0M6 15a1 1 0 1 0 0 2a1 1 0 0 0 0-2m11 1a1 1 0 1 1 2 0a1 1 0 0 1-2 0"/><path fill-rule="evenodd" d="M4 3a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h16a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3zm16 2H4a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Filters(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.708 15.44a6.968 6.968 0 0 0 3.997 1.266a7 7 0 1 0 6.59-9.413A7 7 0 1 0 4.708 15.44m1.148-1.64a4.976 4.976 0 0 0 2.431.886a6.97 6.97 0 0 1 1.256-4.408a6.97 6.97 0 0 1 3.713-2.687a5 5 0 1 0-7.4 6.21m12.289-3.603a4.977 4.977 0 0 0-2.432-.885a6.97 6.97 0 0 1-1.256 4.408a6.97 6.97 0 0 1-3.713 2.687a5 5 0 1 0 7.4-6.21" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flag(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 21h2V11h6v2h8V5h-7V3H4zm8-16H6v4h7v2h5V7h-6z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlagAlt(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m12.439 7l2.4-3H7v6h7.839zM19 12H7v10H5V2h14l-4 5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folder(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 1.5a2 2 0 0 0-2 2v1c0 .057.002.113.007.168A3.001 3.001 0 0 0 0 7.5v12a3 3 0 0 0 3 3h18a3 3 0 0 0 3-3v-12a3 3 0 0 0-3-3h-9.126A4.002 4.002 0 0 0 8 1.5zm5.732 3A2 2 0 0 0 8 3.5H4v1zM3 6.5a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1v-12a1 1 0 0 0-1-1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderAdd(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M11 14.5v2h2v-2h2v-2h-2v-2h-2v2H9v2z"/><path fill-rule="evenodd" d="M4 1.5a2 2 0 0 0-2 2v1c0 .057.002.113.007.168A3.001 3.001 0 0 0 0 7.5v12a3 3 0 0 0 3 3h18a3 3 0 0 0 3-3v-12a3 3 0 0 0-3-3h-9.126A4.002 4.002 0 0 0 8 1.5zm5.732 3A2 2 0 0 0 8 3.5H4v1zM3 6.5a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1v-12a1 1 0 0 0-1-1z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderRemove(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M9 14.5v-2h6v2z"/><path fill-rule="evenodd" d="M4 1.5a2 2 0 0 0-2 2v1c0 .057.002.113.007.168A3.001 3.001 0 0 0 0 7.5v12a3 3 0 0 0 3 3h18a3 3 0 0 0 3-3v-12a3 3 0 0 0-3-3h-9.126A4.002 4.002 0 0 0 8 1.5zm5.732 3A2 2 0 0 0 8 3.5H4v1zM3 6.5a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1v-12a1 1 0 0 0-1-1z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FontHeight(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M21 5V3H3v2zm0 14v2H3v-2z"/><path fill-rule="evenodd" d="M12 7.376a1 1 0 0 0-.967.576l-3.381 7.25a1 1 0 1 0 1.812.846L9.953 15h4.094l.489 1.048a1 1 0 1 0 1.813-.845l-3.381-7.251A1 1 0 0 0 12 7.376M13.115 13h-2.23L12 10.61z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FontSpacing(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M19 21h2V3h-2zM5 21H3V3h2z"/><path fill-rule="evenodd" d="M9.464 16.048L9.953 15h4.094l.489 1.048a1 1 0 1 0 1.813-.845l-3.381-7.25A1 1 0 0 0 12 7.375a1 1 0 0 0-.967.576l-3.381 7.25a1 1 0 0 0 1.812.846M12 10.61L10.885 13h2.23z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatBold(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11 4.504H7v14.992h6a4 4 0 0 0 .604-7.955A4 4 0 0 0 11 4.505Zm-2 2h2a2 2 0 1 1 0 4H9zm0 10.991v-4h4a2 2 0 1 1 0 4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatCenter(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 5a1 1 0 0 0 0 2h16a1 1 0 1 0 0-2zm0 8a1 1 0 1 0 0 2h16a1 1 0 1 0 0-2zm2-3a1 1 0 0 1 1-1h10a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1m1 7a1 1 0 1 0 0 2h10a1 1 0 1 0 0-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatColor(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M12.948 3.21A1 1 0 0 0 12 2.632a1 1 0 0 0-.948.576l-5.917 12.69a1 1 0 1 0 1.813.845l1.45-3.111h7.203l1.451 3.111a1 1 0 0 0 1.813-.845zm1.72 8.422L12 5.909l-2.669 5.723z" clip-rule="evenodd"/><path d="M6 19.368a1 1 0 0 0 0 2h12a1 1 0 1 0 0-2z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatHeading(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 19V5h2v6h8V5h2v14h-2v-6H8v6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatIndentDecrease(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 7H4V5h16zm0 4h-8V9h8zm-8 4h8v-2h-8zM9 9l-5 3l5 3zm-5 8v2h16v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatIndentIncrease(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 7H4V5h16zm0 4h-8V9h8zm-8 4h8v-2h-8zm-8 0l5-3l-5-3zm0 2v2h16v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatItalic(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.49 5.458h6l-.711 1.87h-2l-3.558 9.345h2l-.711 1.87h-6l.711-1.87h2l3.558-9.346h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatJustify(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 5a1 1 0 0 0 0 2h16a1 1 0 1 0 0-2zm0 4a1 1 0 0 0 0 2h16a1 1 0 1 0 0-2zm-1 5a1 1 0 0 1 1-1h16a1 1 0 1 1 0 2H4a1 1 0 0 1-1-1m1 3a1 1 0 1 0 0 2h16a1 1 0 1 0 0-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatLeft(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 5a1 1 0 0 0 0 2h16a1 1 0 1 0 0-2zm0 4a1 1 0 0 0 0 2h8a1 1 0 1 0 0-2zm-1 5a1 1 0 0 1 1-1h16a1 1 0 1 1 0 2H4a1 1 0 0 1-1-1m1 3a1 1 0 1 0 0 2h8a1 1 0 1 0 0-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatLineHeight(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.097 6.997h2.077l-3-3l-3 3h1.923v10.006H1.159l3 3l3-3H5.097zM22.841 7h-14V5h14zm0 4h-14V9h14zm-14 4h14v-2h-14zm14 4h-14v-2h14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatRight(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 5a1 1 0 1 1 0 2H4a1 1 0 0 1 0-2zm0 4a1 1 0 1 1 0 2h-8a1 1 0 1 1 0-2zm1 5a1 1 0 0 0-1-1H4a1 1 0 1 0 0 2h16a1 1 0 0 0 1-1m-1 3a1 1 0 1 1 0 2h-8a1 1 0 1 1 0-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatSeparator(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M16 5a1 1 0 1 0 0-2H8a1 1 0 1 0 0 2zm0 2a1 1 0 1 1 0 2H8a1 1 0 1 1 0-2zm1 5a1 1 0 0 1-1 1H8a1 1 0 1 1 0-2h8a1 1 0 0 1 1 1m-1 9a1 1 0 1 0 0-2H8a1 1 0 1 0 0 2z" opacity=".5"/><path fill-rule="evenodd" d="M21 16a1 1 0 0 1-1 1H4a1 1 0 1 1 0-2h16a1 1 0 0 1 1 1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatSlash(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.526 6.106c.5.233.74.777.537 1.215L9.885 18.424c-.204.438-.775.604-1.276.37c-.5-.233-.74-.778-.536-1.216L13.25 6.476c.204-.438.775-.604 1.276-.37" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatStrike(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 7h4V5H7v2h4v3h2zm-2 12v-5h2v5zm-6-6h14v-2H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatText(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.5 3H3v3.5h1V4h2.5zm2 1V3H11v1zM13 4h2.5V3H13zm4.5-1v1H20v2.5h1V3zM21 8.5h-1V11h1zm0 4.5h-1v2.5h1zm0 4.5h-1V20h-2.5v1H21zM15.5 21v-1H13v1zM11 21v-1H8.5v1zm-4.5 0v-1H4v-2.5H3V21zM3 15.5h1V13H3zM3 11h1V8.5H3zm8-1.5H7v-2h10v2h-4v7h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatUnderline(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 10V4h2v6a4 4 0 0 0 8 0V4h2v6a6 6 0 0 1-12 0m1 8a1 1 0 1 0 0 2h10a1 1 0 1 0 0-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FormatUppercase(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13 9h-3v8H8V9H5V7h8zm5 4h-2v4h-2v-4h-2v-2h6z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Framer(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-opacity=".5" d="M12 21V9H6v6z"/><path d="M18 9V3H6l6 6H6v6h12l-6-6z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Games(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M15.47 11.293a1 1 0 1 0-1.415 1.414a1 1 0 0 0 1.415-1.414m.707-2.121a1 1 0 1 1 1.414 1.414a1 1 0 0 1-1.414-1.414m3.535 2.121a1 1 0 1 0-1.414 1.414a1 1 0 0 0 1.414-1.414m-3.535 2.121a1 1 0 1 1 1.414 1.415a1 1 0 0 1-1.414-1.415M6 13H4v-2h2V9h2v2h2v2H8v2H6z"/><path fill-rule="evenodd" d="M7 5a7 7 0 0 0 0 14h10a7 7 0 1 0 0-14zm10 2H7a5 5 0 0 0 0 10h10a5 5 0 0 0 0-10" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GenderFemale(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 3a5 5 0 0 0-1 9.9V15H8v2h3v4h2v-4h3v-2h-3v-2.1A5.002 5.002 0 0 0 12 3M9 8a3 3 0 1 0 6 0a3 3 0 0 0-6 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GenderMale(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m12.189 7l.002-2l7 .007l-.008 7l-2-.002l.004-3.588l-3.044 3.044a5.002 5.002 0 0 1-7.679 6.336a5 5 0 0 1 6.25-7.735l3.058-3.058zm-4.31 5.14a3 3 0 1 1 4.242 4.244A3 3 0 0 1 7.88 12.14" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ghost(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 4h12v4H3zm18 4h-4V4h4zM3 10h18v4H3zm8 6H3v4h8zm2 0v4h8v-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GhostCharacter(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M10.976 9a1 1 0 1 1-2 0a1 1 0 0 1 2 0m2.995 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path fill-rule="evenodd" d="M19 21V10a7 7 0 1 0-14 0v11h2.828l1.415-1.414L10.657 21h2.686l1.414-1.414L16.172 21zm-2-11a5 5 0 0 0-10 0v9l2.243-2.243L12 19.515l2.757-2.758L17 19z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gift(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M17.535 2.879a3 3 0 0 0-4.242 0l-1.414 1.414c-.06.06-.118.122-.172.186a3.01 3.01 0 0 0-.171-.186L10.12 2.879A3 3 0 1 0 5.88 7.12l.877.88H1v6h2v8h18v-8h2V8h-6.343l.878-.879a3 3 0 0 0 0-4.242M14.707 7.12l1.414-1.414a1 1 0 0 0-1.414-1.414l-1.414 1.414a1 1 0 0 0 1.414 1.414m-4.586-1.414L8.707 4.293a1 1 0 1 0-1.414 1.414l1.414 1.414a1 1 0 1 0 1.414-1.414M21 10v2H3v-2zm-8.083 4H19v6h-6.083zm-1.834 0v6H5v-6z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Girl(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M10 12a1 1 0 1 1-2 0a1 1 0 0 1 2 0m5 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path fill-rule="evenodd" d="M12.024 2H12C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10c0-5.258-4.058-9.568-9.212-9.97v-.002A9.94 9.94 0 0 0 12.025 2M12 20a8 8 0 0 0 7.742-10.022a10.016 10.016 0 0 1-8.981-4.376a7.976 7.976 0 0 1-5.692 2.4A8 8 0 0 0 12 20m-.021-16h.045H12z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitBranch(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 3a2 2 0 0 0-1 3.732v10.536A2 2 0 1 0 10.732 20h1.227a4 4 0 0 0 4-4v-1.268a2 2 0 0 0-1-3.732a2 2 0 0 0-1 3.732V16a2 2 0 0 1-2 2h-1.227a2.01 2.01 0 0 0-.732-.732V6.732A2 2 0 0 0 9 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitCommit(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11 4a1 1 0 1 1 2 0v5.17a3.001 3.001 0 0 1 0 5.66V20a1 1 0 1 1-2 0v-5.17a3.001 3.001 0 0 1 0-5.66zm1 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitFork(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 5a2 2 0 1 1 3 1.732v7.308h1.791a2 2 0 0 0 2-2v-1.256a2 2 0 1 1 2-.024v1.28a4 4 0 0 1-4 4H10v1.228A2 2 0 0 1 9 21a2 2 0 0 1-1-3.732V6.732A2 2 0 0 1 7 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitPull(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 5a2 2 0 1 1 3.763.945h.58a4 4 0 0 1 4 4v1.28a2 2 0 0 1-1.02 3.72a2 2 0 0 1-.98-3.745V9.945a2 2 0 0 0-2-2H10v9.323A2 2 0 0 1 9 21a2 2 0 0 1-1-3.732V6.732A2 2 0 0 1 7 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gitter(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 1.5h2v13H5zm4 3h2v18H9zm6 0h-2v18h2zm2 0h2v10h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Glass(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M17 10a5.002 5.002 0 0 1-4 4.9V17h2v2H9v-2h2v-2.1A5.002 5.002 0 0 1 7 10V5h10zm-2-3H9v3a3 3 0 1 0 6 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlassAlt(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 2h14l-1.64 16.398A4 4 0 0 1 13.38 22h-2.76a4 4 0 0 1-3.98-3.602zm2.51 5l-.3-3h9.58l-.3 3zm.2 2l.92 9.199A2 2 0 0 0 10.62 20h2.76a2 2 0 0 0 1.99-1.801L16.29 9z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Globe(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M6.853 8a5 5 0 1 1 10 0a5 5 0 0 1-10 0m5 3a3 3 0 1 1 0-6a3 3 0 0 1 0 6" clip-rule="evenodd"/><path d="M5 12.13a8.001 8.001 0 0 0 5.941 3.819V18H8.853v2h6v-2h-1.912v-2.073A8.002 8.002 0 0 0 18.63 3.745l-1.704 1.048a6 6 0 1 1-10.221 6.288z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlobeAlt(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 21a9 9 0 1 0 0-18a9 9 0 0 0 0 18m2.806-2.585a7.004 7.004 0 0 0 4.175-5.89c-.823.449-1.861.817-3.044 1.073c-.149 1.87-.554 3.54-1.131 4.817M9.195 5.585a7.016 7.016 0 0 0-3.973 4.659c.232.22.626.49 1.226.757c.45.2.973.379 1.557.529c.054-2.324.498-4.415 1.19-5.945m.906 8.326c.156 1.457.484 2.71.898 3.64c.294.662.593 1.074.823 1.293c.082.078.14.12.178.14a.983.983 0 0 0 .178-.14c.23-.22.529-.63.823-1.292c.414-.932.742-2.184.898-3.641a20.034 20.034 0 0 1-3.798 0m-2.038-.313c.149 1.87.554 3.54 1.131 4.817a7.004 7.004 0 0 1-4.175-5.89c.823.449 1.861.817 3.044 1.073M14 11.89a18.153 18.153 0 0 1-4 0c.014-2.226.423-4.145 1-5.442c.293-.661.592-1.073.822-1.292a.988.988 0 0 1 .178-.14a.988.988 0 0 1 .178.14c.23.22.529.63.823 1.292c.576 1.297.986 3.216.999 5.442m1.995-.36c-.053-2.324-.498-4.415-1.19-5.945a7.016 7.016 0 0 1 3.973 4.659c-.232.22-.626.49-1.226.757c-.45.2-.973.379-1.557.529" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Google(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 12a6 6 0 0 0 11.659 2H12v-4h9.805v4H21.8c-.927 4.564-4.962 8-9.8 8c-5.523 0-10-4.477-10-10S6.477 2 12 2a9.99 9.99 0 0 1 8.282 4.393l-3.278 2.295A6 6 0 0 0 6 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleTasks(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.768 5.714a2 2 0 0 1 3.064 2.572L10.833 19.01a2 2 0 1 1-3.064-2.57zM3 12.74a2 2 0 1 1 4 0a2 2 0 0 1-4 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gym(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.274 9.869l-3.442-4.915l1.639-1.147l3.441 4.915zm-1.884 2.54L16.67 9.95l-8.192 5.736l1.72 2.457l-1.638 1.148l-4.588-6.554L5.61 11.59l1.72 2.458l8.192-5.736l-1.72-2.458l1.638-1.147l4.588 6.554zm2.375-5.326l1.638-1.147l-1.147-1.638l-1.638 1.147zM7.168 19.046l-3.442-4.915l-1.638 1.147l3.441 4.915zm-2.786-.491l-1.638 1.147l-1.147-1.638l1.638-1.147z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hashtag(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 4v4H4v2h4v4H4v2h4v4h2v-4h4v4h2v-4h4v-2h-4v-4h4V8h-4V4h-2v4h-4V4zm6 10v-4h-4v4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Headset(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M17 21a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2h2v-1a7 7 0 1 0-14 0v1h2a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2H3v-9a9 9 0 0 1 18 0v9zm2-6h-2v4h2zM7 15H5v4h2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heart(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m12.012 5.572l-1.087-1.087a5.5 5.5 0 1 0-7.778 7.778l8.839 8.839l.002-.002l.026.026l8.839-8.839a5.5 5.5 0 1 0-7.778-7.778zm-.024 12.7l4.936-4.937l1.45-1.4h.002l1.063-1.062a3.5 3.5 0 1 0-4.95-4.95L12.013 8.4l-.007-.007h-.001L9.511 5.9a3.5 3.5 0 1 0-4.95 4.95l2.54 2.54l.001-.003z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hello(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M17.5 12a5.485 5.485 0 0 1-1.725 4A5.481 5.481 0 0 1 12 17.5c-1.461 0-2.79-.57-3.775-1.5A5.485 5.485 0 0 1 6.5 12z"/><path fill-rule="evenodd" d="M1 7a6 6 0 0 1 6-6h10a6 6 0 0 1 6 6v10a6 6 0 0 1-6 6H1zm2.75 5a8.25 8.25 0 1 1 16.5 0a8.25 8.25 0 0 1-16.5 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Home(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m21 8.772l-6.98-6.979a3 3 0 0 0-4.242 0L3 8.571v14.515h7v-6a2 2 0 1 1 4 0v6h7zm-9.808-5.565L5 9.4v11.686h3v-4a4 4 0 0 1 8 0v4h3V9.6l-6.393-6.394a1 1 0 0 0-1.415 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeAlt(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6 22.879a3 3 0 0 1-3-3v-10c0-.034.002-.068.005-.1H3c0-.577.229-1.13.636-1.536L9.88 2a3 3 0 0 1 4.242 0l6.243 6.243c.407.407.636.96.636 1.535h-.005c.003.033.005.067.005.1v10a3 3 0 0 1-3 3zm6.707-19.465L19 9.707V19.88a1 1 0 0 1-1 1h-3v-5a3 3 0 1 0-6 0v5H6a1 1 0 0 1-1-1V9.707l6.293-6.293a1 1 0 0 1 1.414 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeScreen(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M10 5a1 1 0 1 1-2 0a1 1 0 0 1 2 0M9 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2m1 10a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-1-7a1 1 0 1 0 0-2a1 1 0 0 0 0 2m4-7a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-1 4a1 1 0 1 0 0-2a1 1 0 0 0 0 2m1 10a1 1 0 1 1-2 0a1 1 0 0 1 2 0m2-13a1 1 0 1 0 0-2a1 1 0 0 0 0 2m1 2a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-1 12a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path fill-rule="evenodd" d="M8 1a3 3 0 0 0-3 3v16a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V4a3 3 0 0 0-3-3zm8 2H8a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Icecream(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 17h4V8A7 7 0 1 0 5 8v9h4v3a3 3 0 1 0 6 0zm2-2V8A5 5 0 0 0 7 8v7h4v5a1 1 0 1 0 2 0v-5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IfDesign(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 5h4v14h-4zM5 19v-9h4v9zM7 5a2 2 0 1 0 0 4a2 2 0 0 0 0-4m8 0h4v4h-4zm4 5h-4v4h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Image(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M7 7a3 3 0 1 0 0 6a3 3 0 0 0 0-6m-1 3a1 1 0 1 1 2 0a1 1 0 0 1-2 0"/><path d="M3 3a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h18a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3zm18 2H3a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h4.314l6.878-6.879a3 3 0 0 1 4.243 0L22 15.686V6a1 1 0 0 0-1-1m0 14H10.142l5.465-5.464a1 1 0 0 1 1.414 0l4.886 4.886A1 1 0 0 1 21 19"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Import(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M5 9.982v10h14v-10h-4v-2h6v14H3v-14h6v2z"/><path d="M13 2h-2v12.053l-2.535-2.536l-1.415 1.415l4.95 4.95l4.95-4.95l-1.414-1.415L13 14.054z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Inbox(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 5a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v15a2 2 0 0 1-2 2H8.148a2.006 2.006 0 0 1-.148.005H4a2 2 0 0 1-2-2zm3-1h14a1 1 0 0 1 1 1v9h-4a2 2 0 0 0-2 2v1h-4v-.995a2 2 0 0 0-2-2H4V5a1 1 0 0 1 1-1" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IndieHackers(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 6h3v12H4zm6 0h3v4.5h4V6h3v12h-3v-4.5h-4V18h-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Infinity(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8.121 9.879l2.083 2.083l.007-.006l1.452 1.452l.006.006l2.122 2.122a5 5 0 1 0 0-7.072l-.714.714l1.415 1.414l.713-.713a3 3 0 1 1 0 4.242l-2.072-2.072l-.007.006l-3.59-3.59a5 5 0 1 0 0 7.07l.713-.713l-1.414-1.414l-.714.713a3 3 0 1 1 0-4.242"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Info(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M11 10.98a1 1 0 1 1 2 0v6a1 1 0 1 1-2 0zm1-4.929a1 1 0 1 0 0 2a1 1 0 0 0 0-2"/><path fill-rule="evenodd" d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2M4 12a8 8 0 1 0 16 0a8 8 0 0 0-16 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Inpicture(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M20 12h-6v5h6z"/><path fill-rule="evenodd" d="M1 6a2 2 0 0 1 2-2h18a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H3a2 2 0 0 1-2-2zm2 0h18v12H3z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InsertAfter(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 4a1 1 0 0 1 1 1v3h3a1 1 0 1 1 0 2h-3v3a1 1 0 1 1-2 0v-3H8a1 1 0 0 1 0-2h3V5a1 1 0 0 1 1-1M3 19a1 1 0 0 1 1-1h16a1 1 0 1 1 0 2H4a1 1 0 0 1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InsertAfterO(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M9 8a1 1 0 0 0 0 2h2v2a1 1 0 1 0 2 0v-2h2a1 1 0 1 0 0-2h-2V6a1 1 0 1 0-2 0v2z"/><path fill-rule="evenodd" d="M4 9a8 8 0 1 1 16 0A8 8 0 0 1 4 9m8 6a6 6 0 1 1 0-12a6 6 0 0 1 0 12" clip-rule="evenodd"/><path d="M5 20a1 1 0 1 0 0 2h14a1 1 0 1 0 0-2z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InsertAfterR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M9 8a1 1 0 0 0 0 2h2v2a1 1 0 1 0 2 0v-2h2a1 1 0 1 0 0-2h-2V6a1 1 0 1 0-2 0v2z"/><path fill-rule="evenodd" d="M4 4a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3v10a3 3 0 0 1-3 3H7a3 3 0 0 1-3-3zm3-1h10a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1" clip-rule="evenodd"/><path d="M5 20a1 1 0 1 0 0 2h14a1 1 0 1 0 0-2z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InsertBefore(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 5a1 1 0 0 0 1 1h16a1 1 0 1 0 0-2H4a1 1 0 0 0-1 1m9 15a1 1 0 0 0 1-1v-3h3a1 1 0 1 0 0-2h-3v-3a1 1 0 1 0-2 0v3H8a1 1 0 1 0 0 2h3v3a1 1 0 0 0 1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InsertBeforeO(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M5 3a1 1 0 0 1 0-2h14a1 1 0 1 1 0 2zm4 12a1 1 0 1 1 0-2h2v-2a1 1 0 1 1 2 0v2h2a1 1 0 1 1 0 2h-2v2a1 1 0 1 1-2 0v-2z"/><path fill-rule="evenodd" d="M4 14a8 8 0 1 0 16 0a8 8 0 0 0-16 0m8-6a6 6 0 1 0 0 12a6 6 0 0 0 0-12" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InsertBeforeR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M5 3a1 1 0 0 1 0-2h14a1 1 0 1 1 0 2zm4 12a1 1 0 1 1 0-2h2v-2a1 1 0 1 1 2 0v2h2a1 1 0 1 1 0 2h-2v2a1 1 0 1 1-2 0v-2z"/><path fill-rule="evenodd" d="M4 19a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3H7a3 3 0 0 0-3 3zm13 1a1 1 0 0 0 1-1V9a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Insights(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 8h4v12H5V10h4V4h6zm-2-2h-2v12h2zm2 4v8h2v-8zm-6 2v6H7v-6z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Instagram(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M12 7a5 5 0 1 0 0 10a5 5 0 0 0 0-10m-3 5a3 3 0 1 0 6 0a3 3 0 0 0-6 0" clip-rule="evenodd"/><path d="M18 5a1 1 0 1 0 0 2a1 1 0 0 0 0-2"/><path fill-rule="evenodd" d="M5 1a4 4 0 0 0-4 4v14a4 4 0 0 0 4 4h14a4 4 0 0 0 4-4V5a4 4 0 0 0-4-4zm14 2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Internal(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m20.708 4.412l-10.25 10.287h3.59v2h-7v-7h2v3.58L19.293 3z"/><path d="M11 4.706v2H5v12h12v-6h2v8H3v-16z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Key(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6 8a3 3 0 0 0-3 3v2a3 3 0 1 0 6 0h6v2h2v-2h1v2h2v-4H9a3 3 0 0 0-3-3m1 5v-2a1 1 0 1 0-2 0v2a1 1 0 1 0 2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Keyboard(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M8 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2m1 6a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2zm0-3a1 1 0 1 1-2 0a1 1 0 0 1 2 0m3-3a1 1 0 1 0 0-2a1 1 0 0 0 0 2m1 3a1 1 0 1 1-2 0a1 1 0 0 1 2 0m3-3a1 1 0 1 0 0-2a1 1 0 0 0 0 2m1 3a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/><path fill-rule="evenodd" d="M3 6a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm3-1h12a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Keyhole(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M13 12.83a3.001 3.001 0 1 0-2 0V16a1 1 0 1 0 2 0zM11 10a1 1 0 1 0 2 0a1 1 0 0 0-2 0"/><path d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2M4 12a8 8 0 1 0 16 0a8 8 0 0 0-16 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Laptop(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M3 6a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2zm2 0h14v8H5z" clip-rule="evenodd"/><path d="M2 18a1 1 0 1 0 0 2h20a1 1 0 1 0 0-2z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lastpass(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 6a1 1 0 0 0-1 1v10a1 1 0 1 0 2 0V7a1 1 0 0 0-1-1M4 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4m8-2a2 2 0 1 1-4 0a2 2 0 0 1 4 0m4 2a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutGrid(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 7H7v4h4zm0 6H7v4h4zm2 0h4v4h-4zm4-6h-4v4h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutGridSmall(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 7h2v2H7zm4 0h2v2h-2zm6 0h-2v2h2zM7 11h2v2H7zm6 0h-2v2h2zm2 0h2v2h-2zm-6 4H7v2h2zm2 0h2v2h-2zm6 0h-2v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutList(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 7H7v2h2zm-2 6v-2h2v2zm0 2v2h2v-2zm4 0v2h6v-2zm6-2v-2h-6v2zm0-6v2h-6V7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutPin(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.93 4.016h-2.165a3.001 3.001 0 0 0-5.67 0H6.932a3 3 0 0 0-3 3v2.196a3.001 3.001 0 0 0 0 5.608v2.196a3 3 0 0 0 3 3h2.154a3.001 3.001 0 0 0 5.692 0h2.154a3 3 0 0 0 3-3v-2.171a3.001 3.001 0 0 0 0-5.658V7.016a3 3 0 0 0-3-3m-11 10.853v2.147a1 1 0 0 0 1 1H9.12a3.001 3.001 0 0 1 5.623 0h2.189a1 1 0 0 0 1-1v-2.17a3.001 3.001 0 0 1 0-5.66v-2.17a1 1 0 0 0-1-1h-2.177a3.001 3.001 0 0 1-5.647 0H6.931a1 1 0 0 0-1 1v2.147a3.001 3.001 0 0 1 0 5.706" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Linear(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.035 12.943a8.963 8.963 0 0 0 2.587 5.421a8.963 8.963 0 0 0 5.42 2.587zM3 11.494l9.492 9.492a9.016 9.016 0 0 0 2.378-.459L3.46 9.115A9.016 9.016 0 0 0 3 11.494m.867-3.384l12.009 12.009a8.948 8.948 0 0 0 1.773-1.123L4.99 6.336A8.95 8.95 0 0 0 3.867 8.11m1.796-2.515a9 9 0 0 1 12.728 12.728z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Link(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m14.828 12l1.415 1.414l2.828-2.828a4 4 0 0 0-5.657-5.657l-2.828 2.828L12 9.172l2.828-2.829a2 2 0 1 1 2.829 2.829zM12 14.829l1.414 1.414l-2.828 2.828a4 4 0 0 1-5.657-5.657l2.828-2.828L9.172 12l-2.829 2.829a2 2 0 1 0 2.829 2.828z"/><path d="M14.829 10.586a1 1 0 0 0-1.415-1.415l-4.242 4.243a1 1 0 1 0 1.414 1.414z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func List(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M20 4H4a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1M4 2a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h16a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3zm2 5h2v2H6zm5 0a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2zm-3 4H6v2h2zm2 1a1 1 0 0 1 1-1h6a1 1 0 1 1 0 2h-6a1 1 0 0 1-1-1m-2 3H6v2h2zm2 1a1 1 0 0 1 1-1h6a1 1 0 1 1 0 2h-6a1 1 0 0 1-1-1" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListTree(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9 1H1v8h8V6h2v14h4v3h8v-8h-8v3h-2V6h2v3h8V1h-8v3H9zm12 2h-4v4h4zm-4 14h4v4h-4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LivePhoto(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M12.98 21.953a10.12 10.12 0 0 1-1.96 0l.193-1.991a8.1 8.1 0 0 0 1.574 0zm-3.884-.381l.58-1.914a7.947 7.947 0 0 1-1.446-.6l-.945 1.763c.573.307 1.179.56 1.811.75m-3.44-1.841l1.27-1.545a8.062 8.062 0 0 1-1.111-1.11L4.27 18.343c.415.506.88.97 1.386 1.386m-2.477-3.014l1.763-.944a7.938 7.938 0 0 1-.6-1.447l-1.914.58a9.93 9.93 0 0 0 .751 1.81M2.047 12.98l1.991-.193a8.12 8.12 0 0 1 0-1.574l-1.99-.193a10.123 10.123 0 0 0 0 1.96m.38-3.884l1.914.58c.153-.505.355-.989.6-1.446l-1.763-.945a9.938 9.938 0 0 0-.75 1.811m1.841-3.44l1.545 1.27a8.06 8.06 0 0 1 1.11-1.111L5.657 4.27c-.506.415-.97.88-1.386 1.386m3.014-2.477l.945 1.763a7.938 7.938 0 0 1 1.446-.6l-.58-1.914a9.938 9.938 0 0 0-1.81.751m3.734-1.132a10.123 10.123 0 0 1 1.96 0l-.193 1.991a8.12 8.12 0 0 0-1.574 0zm3.884.381l-.58 1.914c.505.153.989.355 1.447.6l.944-1.763a9.936 9.936 0 0 0-1.811-.75m3.44 1.841l-1.27 1.545c.406.333.778.705 1.111 1.11l1.545-1.269a10.06 10.06 0 0 0-1.386-1.386m2.477 3.015l-1.763.945c.245.457.447.941.6 1.446l1.914-.58a9.937 9.937 0 0 0-.751-1.81m1.132 3.734l-1.991.193a8.1 8.1 0 0 1 0 1.574l1.99.194a10.123 10.123 0 0 0 0-1.961m-.38 3.884l-1.914-.58a7.947 7.947 0 0 1-.6 1.447l1.763.944c.307-.573.56-1.179.75-1.811m-1.841 3.44l-1.545-1.27a8.063 8.063 0 0 1-1.11 1.111l1.269 1.545c.506-.415.97-.88 1.386-1.386m-3.015 2.477l-.944-1.763a7.947 7.947 0 0 1-1.447.6l.58 1.914a9.935 9.935 0 0 0 1.81-.751"/><path fill-rule="evenodd" d="M9 12a3 3 0 1 1 6 0a3 3 0 0 1-6 0m3 1a1 1 0 1 1 0-2a1 1 0 0 1 0 2" clip-rule="evenodd"/><path fill-rule="evenodd" d="M12 6a6 6 0 1 0 0 12a6 6 0 0 0 0-12m-4 6a4 4 0 1 0 8 0a4 4 0 0 0-8 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Loadbar(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<rect width="18" height="4" x="3" y="10" fill="currentColor" rx="2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LoadbarAlt(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><rect width="18" height="4" x="3" y="10" opacity=".3" rx="2"/><rect width="10" height="4" x="7" y="10" rx="2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LoadbarDoc(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M17 5H7a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1M7 3a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3z" clip-rule="evenodd"/><path d="M8 7h8v2H8zm0 4h8v2H8zm0 4h5v2H8z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LoadbarSound(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 6h2v12h-2zm-4 7h2v5H7zm8-4h2v9h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lock(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M18 10.5a3 3 0 0 1 3 3v6a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3v-6a3 3 0 0 1 3-3v-3a6 6 0 1 1 12 0zm-6-7a4 4 0 0 1 4 4v3H8v-3a4 4 0 0 1 4-4m6 9H6a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockUnlock(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M19 7h-2a4 4 0 0 0-8 0v3h9a3 3 0 0 1 3 3v6a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3v-6a3 3 0 0 1 3-3h1V7a6 6 0 1 1 12 0m-1 5H6a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogIn(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M15.486 20h4a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2h-4v2h4v12h-4z"/><path d="m10.158 17.385l-1.42-1.408l3.92-3.953H3.513a1 1 0 1 1 0-2h9.163l-3.98-3.947l1.408-1.42l6.391 6.337z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogOff(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M13 4.009a1 1 0 1 0-2 0l-.003 8.003a1 1 0 0 0 2 0z"/><path d="M4 12.992c0-2.21.895-4.21 2.343-5.657l1.414 1.414a6 6 0 1 0 8.485 0l1.415-1.414A8 8 0 1 1 4 12.992"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LogOut(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M8.514 20h-4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h4v2h-4v12h4z"/><path d="m13.842 17.385l1.42-1.408l-3.92-3.953h9.144a1 1 0 1 0 0-2h-9.162l3.98-3.947l-1.408-1.42l-6.391 6.337z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Loupe(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M11 11V8h2v3h3v2h-3v3h-2v-3H8v-2z"/><path fill-rule="evenodd" d="M3 12a9 9 0 0 0 9 9h6a3 3 0 0 0 3-3v-6a9 9 0 1 0-18 0m9-7a7 7 0 1 1 0 14a7 7 0 0 1 0-14" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Magnet(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-opacity=".5" d="M8 2.5H4v3h4zm12 0h-4v3h4z"/><path d="M8 7.5H4v6a8 8 0 1 0 16 0v-6h-4v6a4 4 0 0 1-8 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mail(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.01 5.838a1 1 0 0 1 1-1H20a1 1 0 0 1 1 1v11.324a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-11c0-.048.003-.094.01-.14zM5 8.062v9.1h14v-9.1l-4.879 4.879a3 3 0 0 1-4.242 0zm1.572-1.256h10.856l-4.72 4.72a1 1 0 0 1-1.415 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailForward(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.638 15.529l1.414 1.414l4.95-4.95l-4.95-4.95l-1.414 1.415l2.498 2.498H7.998a4 4 0 0 0-4 4v2h2v-2a2 2 0 0 1 2-2h8.212z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailOpen(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.05 10.015a2 2 0 0 1 .465-2.1L9.879 1.55a3 3 0 0 1 4.242 0l6.364 6.364a2 2 0 0 1 .465 2.101c.032.099.05.204.05.313v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-11a1 1 0 0 1 .05-.313m1.879-.687l6.364-6.363a1 1 0 0 1 1.414 0l6.364 6.363h-.03v.03l-6.334 6.334a1 1 0 0 1-1.414 0zm14.07 2.9l-4.878 4.879a3 3 0 0 1-4.242 0l-4.88-4.88v9.101h14z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailReply(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.362 15.529l-1.414 1.414l-4.95-4.95l4.95-4.95l1.414 1.415l-2.498 2.498h8.138a4 4 0 0 1 4 4v2h-2v-2a2 2 0 0 0-2-2H7.79z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MathDivide(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 7a2 2 0 1 1-4 0a2 2 0 0 1 4 0M3 12a1 1 0 0 1 1-1h16a1 1 0 1 1 0 2H4a1 1 0 0 1-1-1m9 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MathEqual(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 9a1 1 0 1 0 0 2h14a1 1 0 1 0 0-2zm0 4a1 1 0 1 0 0 2h14a1 1 0 1 0 0-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MathMinus(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 12a1 1 0 0 1 1-1h14a1 1 0 1 1 0 2H5a1 1 0 0 1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MathPercent(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.243 6.343a1 1 0 1 1 1.414 1.414l-9.9 9.9a1 1 0 0 1-1.414-1.414zM9.879 9.879A2 2 0 1 1 7.05 7.05a2 2 0 0 1 2.83 2.83m4.241 7.07a2 2 0 1 0 2.829-2.829a2 2 0 0 0-2.829 2.829"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MathPlus(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 4a1 1 0 0 0-1 1v6H5a1 1 0 1 0 0 2h6v6a1 1 0 1 0 2 0v-6h6a1 1 0 1 0 0-2h-6V5a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Maximize(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 3h6v2H5v4H3zm0 18h6v-2H5v-4H3zm12 0h6v-6h-2v4h-4zm6-18h-6v2h4v4h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MaximizeAlt(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 3h6v2H6.462l4.843 4.843l-1.415 1.414L5 6.367V9H3zm0 18h6v-2H6.376l4.929-4.928l-1.415-1.414L5 17.548V15H3zm12 0h6v-6h-2v2.524l-4.867-4.866l-1.414 1.414L17.647 19H15zm6-18h-6v2h2.562l-4.843 4.843l1.414 1.414L19 6.39V9h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Maze(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.37 9.593L8.779 7L1 14.778l2.593 2.593zM15.222 7L23 14.778l-2.576 2.576l-5.202-5.202l-5.202 5.202l-2.576-2.576z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaLive(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M15 12a3 3 0 1 1-6 0a3 3 0 0 1 6 0m-2 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/><path d="M12 19a7 7 0 1 0 0-14a7 7 0 0 0 0 14m0-2a5 5 0 1 0 0-10a5 5 0 0 0 0 10"/><path d="M12 23c6.075 0 11-4.925 11-11S18.075 1 12 1S1 5.925 1 12s4.925 11 11 11m0-2a9 9 0 1 0 0-18a9 9 0 0 0 0 18"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaPodcast(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M5.636 20.364a9 9 0 1 1 12.728 0l1.414 1.414A10.966 10.966 0 0 0 23 14c0-6.075-4.925-11-11-11S1 7.925 1 14c0 3.038 1.231 5.788 3.222 7.778z"/><path d="M16.95 18.95a7 7 0 1 0-9.9 0l1.415-1.414a5 5 0 1 1 7.071 0z"/><path d="M14.121 16.121a3 3 0 1 0-4.243 0l1.415-1.414a1 1 0 1 1 1.414 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Menu(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 6a1 1 0 0 1 1-1h18a1 1 0 1 1 0 2H3a1 1 0 0 1-1-1m0 6.032a1 1 0 0 1 1-1h18a1 1 0 1 1 0 2H3a1 1 0 0 1-1-1m1 5.033a1 1 0 1 0 0 2h18a1 1 0 0 0 0-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuBoxed(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M8.016 6.982a1.003 1.003 0 0 0 0 2.006h7.95a1.003 1.003 0 0 0 0-2.006zm-1 5.018c0-.552.447-1.003 1-1.003h7.95a1.003 1.003 0 0 1 0 2.006h-7.95c-.553 0-1-.45-1-1.003m1.009 3.012a1.003 1.003 0 0 0 0 2.007h7.95a1.003 1.003 0 0 0 0-2.007z"/><path fill-rule="evenodd" d="M3 6a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm3-1h12a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuCake(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 8a2 2 0 1 0 0-4a2 2 0 0 0 0 4m-7 2a1 1 0 1 0 0 2h14a1 1 0 1 0 0-2zm-1 5a1 1 0 0 1 1-1h14a1 1 0 1 1 0 2H5a1 1 0 0 1-1-1m1 3a1 1 0 1 0 0 2h14a1 1 0 1 0 0-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuCheese(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 6a1 1 0 0 1 1-1h16a1 1 0 1 1 0 2H4a1 1 0 0 1-1-1m0 12a1 1 0 0 1 1-1h16a1 1 0 1 1 0 2H4a1 1 0 0 1-1-1m0-6a1 1 0 0 1 1-1h16a1 1 0 1 1 0 2h-9.737l-2.648 2.648L4.967 13H4a1 1 0 0 1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuGridO(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 6a2 2 0 1 1-4 0a2 2 0 0 1 4 0m0 6a2 2 0 1 1-4 0a2 2 0 0 1 4 0m-2 8a2 2 0 1 0 0-4a2 2 0 0 0 0 4m8-14a2 2 0 1 1-4 0a2 2 0 0 1 4 0m-2 8a2 2 0 1 0 0-4a2 2 0 0 0 0 4m2 4a2 2 0 1 1-4 0a2 2 0 0 1 4 0m4-10a2 2 0 1 0 0-4a2 2 0 0 0 0 4m2 4a2 2 0 1 1-4 0a2 2 0 0 1 4 0m-2 8a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuGridR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 4h4v4H4zm0 6h4v4H4zm4 6H4v4h4zm2-12h4v4h-4zm4 6h-4v4h4zm-4 6h4v4h-4zM20 4h-4v4h4zm-4 6h4v4h-4zm4 6h-4v4h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuHotdog(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 6a3 3 0 0 0-3 3h16a3 3 0 0 0-3-3zm0 12a3 3 0 0 1-3-3h16a3 3 0 0 1-3 3zm-4-7a1 1 0 1 0 0 2h18a1 1 0 1 0 0-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuLeft(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 5.995c0-.55.446-.995.995-.995h8.01a.995.995 0 0 1 0 1.99h-8.01A.995.995 0 0 1 2 5.995M2 12c0-.55.446-.995.995-.995h18.01a.995.995 0 1 1 0 1.99H2.995A.995.995 0 0 1 2 12m.995 5.01a.995.995 0 0 0 0 1.99h12.01a.995.995 0 0 0 0-1.99z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuLeftAlt(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 6a1 1 0 0 1 1-1h14a1 1 0 1 1 0 2H5a1 1 0 0 1-1-1m0 12a1 1 0 0 1 1-1h14a1 1 0 1 1 0 2H5a1 1 0 0 1-1-1m1-7a1 1 0 1 0 0 2h8a1 1 0 1 0 0-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuMotion(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 5a1 1 0 1 0 0 2h8a1 1 0 1 0 0-2zm-5 7a1 1 0 0 1 1-1h8a1 1 0 1 1 0 2H8a1 1 0 0 1-1-1m-4 6a1 1 0 0 1 1-1h8a1 1 0 1 1 0 2H4a1 1 0 0 1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuOreos(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 3a3 3 0 0 0-3 3h16a3 3 0 0 0-3-3zm0 8a3 3 0 0 1-3-3h16a3 3 0 0 1-3 3zm0 2a3 3 0 0 0-3 3h16a3 3 0 0 0-3-3zm0 8a3 3 0 0 1-3-3h16a3 3 0 0 1-3 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuRight(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 18.005c0 .55-.446.995-.995.995h-8.01a.995.995 0 0 1 0-1.99h8.01c.55 0 .995.445.995.995M22 12c0 .55-.446.995-.995.995H2.995a.995.995 0 1 1 0-1.99h18.01c.55 0 .995.446.995.995m-.995-5.01a.995.995 0 0 0 0-1.99H8.995a.995.995 0 1 0 0 1.99z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuRightAlt(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 6a1 1 0 0 1 1-1h14a1 1 0 1 1 0 2H5a1 1 0 0 1-1-1m0 12a1 1 0 0 1 1-1h14a1 1 0 1 1 0 2H5a1 1 0 0 1-1-1m7-7a1 1 0 1 0 0 2h8a1 1 0 1 0 0-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuRound(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M8 6.983a1 1 0 1 0 0 2h8a1 1 0 1 0 0-2zM7 12a1 1 0 0 1 1-1h8a1 1 0 1 1 0 2H8a1 1 0 0 1-1-1m1 3.017a1 1 0 1 0 0 2h8a1 1 0 1 0 0-2z"/><path fill-rule="evenodd" d="M22 12c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2s10 4.477 10 10m-2 0a8 8 0 1 1-16 0a8 8 0 0 1 16 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MergeHorizontal(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12 8.976l4.243-4.243l-1.415-1.414L12 6.147L9.172 3.32L7.757 4.733zM5 12a1 1 0 0 1 1-1h12a1 1 0 0 1 0 2H6a1 1 0 0 1-1-1m7 3.024l-4.243 4.243l1.415 1.414L12 17.853l2.828 2.828l1.415-1.414z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MergeVertical(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.976 12L4.733 7.757L3.32 9.172L6.147 12L3.32 14.828l1.414 1.415zM12 19a1 1 0 0 1-1-1V6a1 1 0 1 1 2 0v12a1 1 0 0 1-1 1m3.024-7l4.243 4.243l1.414-1.415L17.853 12l2.828-2.828l-1.414-1.415z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mic(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M9 4a3 3 0 1 1 6 0v8a3 3 0 1 1-6 0zm4 0v8a1 1 0 1 1-2 0V4a1 1 0 1 1 2 0" clip-rule="evenodd"/><path d="M18 12a6.002 6.002 0 0 1-5 5.917V21h4v2H7v-2h4v-3.083A6.002 6.002 0 0 1 6 12V9h2v3a4 4 0 0 0 8 0V9h2z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Microbit(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M7 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4m12-2a2 2 0 1 1-4 0a2 2 0 0 1 4 0"/><path fill-rule="evenodd" d="M7 5a7 7 0 0 0 0 14h10a7 7 0 1 0 0-14zm10 3H7a4 4 0 1 0 0 8h10a4 4 0 0 0 0-8" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Microsoft(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 3h8v8H3zm0 10h8v8H3zM13 3h8v8h-8zm0 10h8v8h-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MiniPlayer(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 6a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm3-1h12a1 1 0 0 1 1 1v6.268A1.99 1.99 0 0 0 18 12h-4a2 2 0 0 0-2 2v4c0 .364.097.706.268 1H6a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Minimize(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 9H3V7h4V3h2zm0 6H3v2h4v4h2zm12 0h-6v6h2v-4h4zm-6-6h6V7h-4V3h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinimizeAlt(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.074 2l1.414 1.414l-5.85 5.85h2.544v2h-6v-6h2v2.627zm-8.892 10.264v6h-2v-2.422L3.414 21.61L2 20.196l5.932-5.932h-2.75v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Modem(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M18 16.634a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/><path fill-rule="evenodd" d="M5.866 3.134a1 1 0 1 0-1 1.732l13.455 7.768H2v4a4 4 0 0 0 4 4h12a4 4 0 0 0 4-4v-4zM20 14.634H4v2a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Monday(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.779 6.14a3 3 0 0 1 4.915 3.44l-5.736 8.192a3 3 0 0 1-4.915-3.441zm8.489.088a3 3 0 0 1 4.915 3.442l-5.736 8.191a3 3 0 0 1-4.915-3.441zM20.5 18.86a3 3 0 1 0 0-6a3 3 0 0 0 0 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Moon(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.226 2.003a9.971 9.971 0 0 0-7.297 2.926c-3.905 3.905-3.905 10.237 0 14.142c3.905 3.905 10.237 3.905 14.142 0a9.972 9.972 0 0 0 2.926-7.297a10.037 10.037 0 0 0-.337-2.368a14.87 14.87 0 0 1-1.744 1.436c-1.351.949-2.733 1.563-3.986 1.842c-1.906.423-3.214.032-3.93-.684c-.716-.716-1.107-2.024-.684-3.93c.279-1.253.893-2.635 1.842-3.986c.414-.592.893-1.177 1.436-1.744a10.017 10.017 0 0 0-2.368-.337m5.43 15.654a7.964 7.964 0 0 0 2.251-4.438c-3.546 2.045-7.269 2.247-9.321.195c-2.052-2.052-1.85-5.775.195-9.321a8 8 0 1 0 6.876 13.564" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func More(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6m0-2a1 1 0 1 0 0-2a1 1 0 0 0 0 2m7 2a3 3 0 1 0 0-6a3 3 0 0 0 0 6m0-2a1 1 0 1 0 0-2a1 1 0 0 0 0 2m10-1a3 3 0 1 1-6 0a3 3 0 0 1 6 0m-2 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoreAlt(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 12a2 2 0 1 1-4 0a2 2 0 0 1 4 0m6 0a2 2 0 1 1-4 0a2 2 0 0 1 4 0m4 2a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoreO(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M7 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4m7-2a2 2 0 1 1-4 0a2 2 0 0 1 4 0m3 2a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/><path fill-rule="evenodd" d="M24 12c0 6.627-5.373 12-12 12S0 18.627 0 12S5.373 0 12 0s12 5.373 12 12m-2 0c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2s10 4.477 10 10" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoreR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M7 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4m7-2a2 2 0 1 1-4 0a2 2 0 0 1 4 0m3 2a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/><path fill-rule="evenodd" d="M0 5a3 3 0 0 1 3-3h18a3 3 0 0 1 3 3v14a3 3 0 0 1-3 3H3a3 3 0 0 1-3-3zm3-1h18a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoreVertical(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 5a1 1 0 1 0 0-2a1 1 0 0 0 0 2m0 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2m0 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2m3-17a3 3 0 1 1-6 0a3 3 0 0 1 6 0m0 8a3 3 0 1 1-6 0a3 3 0 0 1 6 0m-3 11a3 3 0 1 0 0-6a3 3 0 0 0 0 6" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoreVerticalAlt(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 6a2 2 0 1 1-4 0a2 2 0 0 1 4 0m0 6a2 2 0 1 1-4 0a2 2 0 0 1 4 0m0 6a2 2 0 1 1-4 0a2 2 0 0 1 4 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoreVerticalO(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M12 9a2 2 0 1 0 0-4a2 2 0 0 0 0 4m2 3a2 2 0 1 1-4 0a2 2 0 0 1 4 0m-2 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/><path fill-rule="evenodd" d="M24 12c0 6.627-5.373 12-12 12S0 18.627 0 12S5.373 0 12 0s12 5.373 12 12m-2 0c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2s10 4.477 10 10" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoreVerticalR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M12 9a2 2 0 1 0 0-4a2 2 0 0 0 0 4m2 3a2 2 0 1 1-4 0a2 2 0 0 1 4 0m-2 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/><path fill-rule="evenodd" d="M2 3a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v18a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3zm3-1h14a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V3a1 1 0 0 1 1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mouse(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M12 5a1 1 0 0 0-1 1v4a1 1 0 1 0 2 0V6a1 1 0 0 0-1-1"/><path fill-rule="evenodd" d="M4 8a8 8 0 1 1 16 0v8a8 8 0 1 1-16 0zm14 0v8a6 6 0 0 1-12 0V8a6 6 0 1 1 12 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoveDown(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M7 5h2v8H7zm8 0h2v8h-2z"/><path d="M11 5h2v10h3.035l-4 4.071l-4-4.071H11z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoveLeft(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M19.071 17v-2h-8v2zm0-8V7h-8v2z"/><path d="M19.071 13v-2h-10V7.965l-4.071 4l4.071 4V13z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoveRight(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M5 17v-2h8v2zm0-8V7h8v2z"/><path d="M5 13v-2h10V7.965l4.071 4l-4.071 4V13z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoveTask(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.964 7h-8v2h8zM6 8.829v6.342L9.964 12zM18.964 11h-8v2h8zm-8 4h8v2h-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoveUp(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M17 19.071h-2v-8h2zm-8 0H7v-8h2z"/><path d="M13 19.071h-2v-10H7.965l4-4.071l4 4.071H13z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Music(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M22 6a4 4 0 0 0-4.608-3.953l-7 1.077A4 4 0 0 0 7 7.078v8.763a3.5 3.5 0 1 0 2 3.163V7.078A2 2 0 0 1 10.696 5.1l7-1.077A2 2 0 0 1 20 6.001v6.84a3.5 3.5 0 1 0 2 3.163zm-2 10.004a1.5 1.5 0 1 0-3 0a1.5 1.5 0 0 0 3 0m-13 3a1.5 1.5 0 1 0-3 0a1.5 1.5 0 0 0 3 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicNote(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m12 8.954l5.635-1.127a2.942 2.942 0 0 0-1.154-5.769l-4.07.814A3 3 0 0 0 10 5.814v8.076a4 4 0 1 0 2 3.465zm4.874-4.935l-4.07.814a1 1 0 0 0-.804.98v1.102l5.243-1.049a.942.942 0 0 0-.37-1.847M10 17.354a2 2 0 1 0-4 0a2 2 0 0 0 4 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicSpeaker(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M12 18.939a4 4 0 1 0 0-8a4 4 0 0 0 0 8m0-2a2 2 0 1 0 0-4a2 2 0 0 0 0 4" clip-rule="evenodd"/><path d="M12 9.044a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/><path fill-rule="evenodd" d="M7 1a3 3 0 0 0-3 3v16a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V4a3 3 0 0 0-3-3zm10 2H7a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Nametag(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M4 14v6h6v-2H6v-4z"/><path fill-rule="evenodd" d="M9 9v6h6V9zm4 2h-2v2h2z" clip-rule="evenodd"/><path d="M4 10V4h6v2H6v4zm16 0V4h-6v2h4v4zm0 4v6h-6v-2h4v-4z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Notes(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M6 6a1 1 0 0 1 1-1h10a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1m0 4a1 1 0 0 1 1-1h10a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1m1 3a1 1 0 1 0 0 2h10a1 1 0 1 0 0-2zm-1 5a1 1 0 0 1 1-1h4a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1"/><path fill-rule="evenodd" d="M2 4a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v16a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3zm3-1h14a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Notifications(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M20 7a3 3 0 1 1-6 0a3 3 0 0 1 6 0"/><path d="M12 6H4v14h14v-8h-2v6H6V8h6z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Npm(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 21a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2zm1-3V6h12v12h-3V9h-3v9z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Oculus(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16 10H8a2 2 0 1 0 0 4h8a2 2 0 1 0 0-4M8 6a6 6 0 1 0 0 12h8a6 6 0 0 0 0-12z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OpenCollective(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-opacity=".5" d="m16.682 15.753l2.13 2.13A8.965 8.965 0 0 0 21 12a8.964 8.964 0 0 0-2.123-5.806l-2.133 2.133A5.974 5.974 0 0 1 18 12c0 1.42-.493 2.725-1.318 3.753"/><path d="M15.673 16.745a6 6 0 1 1 .08-9.426l2.13-2.13a9 9 0 1 0-.077 13.689z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Options(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7 3a4.002 4.002 0 0 1 3.874 3H19v2h-8.126A4.002 4.002 0 0 1 3 7a4 4 0 0 1 4-4m0 6a2 2 0 1 0 0-4a2 2 0 0 0 0 4m10 11a4.002 4.002 0 0 1-3.874-3H5v-2h8.126A4.002 4.002 0 0 1 21 16a4 4 0 0 1-4 4m0-2a2 2 0 1 0 0-4a2 2 0 0 0 0 4" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Organisation(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M17 15h2v2h-2zm2-4h-2v2h2z"/><path fill-rule="evenodd" d="M13 7h10v14H1V3h12zM8 5h3v2H8zm3 14v-2H8v2zm0-4v-2H8v2zm0-4V9H8v2zm10 8V9h-8v2h2v2h-2v2h2v2h-2v2zM3 19v-2h3v2zm0-4h3v-2H3zm3-4V9H3v2zM3 7h3V5H3z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Overflow(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M22 11a10 10 0 0 1-20 0z" opacity=".2"/><path d="M20 11a8 8 0 0 1-16 0z" opacity=".3"/><path d="M20 11a8 8 0 0 0-16 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pacman(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M14.064 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path fill-rule="evenodd" d="M13 3c2.152 0 4.128.756 5.677 2.016l1.447 1.447l-1.295 1.295h-.001L14.585 12l3.639 3.638l-.002.002l1.905 1.904l-1.413 1.413l-.002-.002A9 9 0 1 1 13 3m-1.243 9l5.532 5.532a7 7 0 1 1 0-11.065z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Password(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M6 12a1 1 0 1 1-2 0a1 1 0 0 1 2 0m3 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2m5-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0m6-1h-4v2h4z"/><path fill-rule="evenodd" d="M2 6a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h20a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2zm20 2H2v8h20z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PathBack(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9 14H4V4h10v5h5v10H9zM6 6h6v6H6z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PathCrop(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><path stroke="currentColor" stroke-width="2" d="M6 6h8v8H6z" opacity=".5"/><path fill="currentColor" fill-rule="evenodd" d="M9 9h10v10H9zm6 2h2v6h-6v-2h4z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PathDivide(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M5 5h10v4H9v6H5z"/><path d="M9 15v4h10V9h-4v6z"/><path d="M10 10h4v4h-4z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PathExclude(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 5h10v4H9v6H5zm4 10v4h10V9h-4v6z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PathFront(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 5h10v4h4v10H9v-4H5zm6 6v6h6v-6z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PathIntersect(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 5H5v10h4v4h10V9h-4zm-2 2H7v6h2V9h4zm4 10h-6v-2h4v-4h2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PathOutline(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 5h10v4h4v10H9v-4H5zm2 2h6v2H9v4H7zm4 10h6v-6h-2v4h-4zm2-6h-2v2h2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PathTrim(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M5 5h10v3H8v7H5z"/><path d="M19 9H9v10h10z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PathUnite(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 5H5v10h4v4h10V9h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Patreon(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M21 10a6 6 0 1 1-12 0a6 6 0 0 1 12 0" opacity=".5"/><path d="M3 4h4v16H3z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paypal(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.47 3.544h8c1.639 0 2.945.775 3.626 1.971c1.22.872 1.847 2.4 1.512 4.138c-.521 2.712-3.183 4.91-5.944 4.91h-2l-1.134 5.892H6.398l.23-1.199h-3.18zm1.622 1.964h6c1.657 0 2.746 1.32 2.433 2.946c-.313 1.627-1.91 2.946-3.566 2.946h-4l-1.134 5.892h-2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pen(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M21.264 2.293a1 1 0 0 0-1.415 0l-.872.872a3.001 3.001 0 0 0-3.415.587L4.955 14.358l5.657 5.657L21.22 9.41a3 3 0 0 0 .586-3.415l.873-.873a1 1 0 0 0 0-1.414zm-4.268 8.51l-6.384 6.384l-2.828-2.829l6.383-6.383zm1.818-1.818l.99-.99a1 1 0 0 0 0-1.415l-1.413-1.414a1 1 0 0 0-1.415 0l-.99.99z" clip-rule="evenodd"/><path d="m2 22.95l2.122-7.778l5.656 5.657z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PentagonBottomLeft(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.33 9.232L10 5l-5 8.66l3.33 4.232l5.33.768l5-8.66zm2.121 2.326l-3.198-.46l-1.998-2.54l-2.846 4.93l1.998 2.539l3.198.46z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PentagonBottomRight(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.33 8.232L13.66 4l5 8.66l-3.33 4.232l-5.33.768L5 9zm-2.12 2.326l3.197-.46l1.998-2.54l2.846 4.93l-1.998 2.539l-3.198.46z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PentagonDown(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 8L7 6v10l5 2.5l5-2.5V6zm3 .954l-3 1.2l-3-1.2v5.81l3 1.5l3-1.5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PentagonLeft(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m16 12l2-5H8l-2 5l2 5h10zm-.954 3l-1.2-3l1.2-3H9.354l-1.2 3l1.2 3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PentagonRight(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 12L6 7h10l2 5l-2 5H6zm.954 3l1.2-3l-1.2-3h5.692l1.2 3l-1.2 3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PentagonTopLeft(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.33 15.16L11 19.392l-5-8.66L9.33 6.5l5.33-.768l5 8.66zm2.121-2.326l-3.198.46l-1.998 2.54l-2.846-4.93l1.998-2.539l3.198-.46z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PentagonTopRight(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.33 15.16L5 14.392l5-8.66l5.33.768l3.33 4.232l-5 8.66zm3.075.674l-1.998-2.54l-3.198-.46l2.846-4.93l3.198.461l1.998 2.54z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PentagonUp(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m12 16l5 2V8l-5-2l-5 2v10zm-3-.954l3-1.2l3 1.2V9.354l-3-1.2l-3 1.2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Performance(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M11 9v4.17a3.001 3.001 0 1 0 2 0V9zm0 7a1 1 0 1 1 2 0a1 1 0 0 1-2 0" clip-rule="evenodd"/><path d="M12 5a7 7 0 0 1 7 7v1h-2v-1a5 5 0 0 0-10 0v1H5v-1a7 7 0 0 1 7-7"/><path fill-rule="evenodd" d="M12 23c6.075 0 11-4.925 11-11S18.075 1 12 1S1 5.925 1 12s4.925 11 11 11m0-2a9 9 0 1 0 0-18a9 9 0 0 0 0 18" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pexels(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 5a5.001 5.001 0 0 1 0 10v4H6V5zM8 7v10h2v-4h2l.003-.001a2.993 2.993 0 0 0 3.041-3.044l-.007-.39a2.61 2.61 0 0 0-2.711-2.562l-.306.005L12 7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Phone(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M22 12A10.002 10.002 0 0 0 12 2v2a8.003 8.003 0 0 1 7.391 4.938A8 8 0 0 1 20 12zM2 10V5a1 1 0 0 1 1-1h5a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1H6a8 8 0 0 0 8 8v-2a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v5a1 1 0 0 1-1 1h-5C7.373 22 2 16.627 2 10"/><path d="M17.543 9.704A5.99 5.99 0 0 1 18 12h-1.8A4.199 4.199 0 0 0 12 7.8V6a6 6 0 0 1 5.543 3.704"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Photoscan(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M9 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path fill-rule="evenodd" d="M17 3a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2zm-4.535 2H17v11H7v-5.535A4 4 0 0 0 12.465 5M9 5a2 2 0 1 0 0 4a2 2 0 0 0 0-4" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Piano(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M22 21a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H3a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2zM11 5H8.985v8h-1v6H12v-6h-1zm7.015 14H22V5h-2.985v8h-1zm-1-6h-1V5H14v8h-1v6h4.015zm-10.03 6v-6h-1V5H3v14z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pill(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.657 2.757a6 6 0 1 1 8.485 8.486l-9.9 9.9a6 6 0 1 1-8.485-8.486zm7.07 7.071l-4.242 4.243l-5.657-5.657l4.243-4.242a4 4 0 1 1 5.657 5.656" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pin(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M16.272 10.272a4 4 0 1 1-8 0a4 4 0 0 1 8 0m-2 0a2 2 0 1 1-4 0a2 2 0 0 1 4 0"/><path d="M5.794 16.518a9 9 0 1 1 12.724-.312l-6.206 6.518zm11.276-1.691l-4.827 5.07l-5.07-4.827a7 7 0 1 1 9.897-.243"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinAlt(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M12 11a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/><path fill-rule="evenodd" d="M18 9a6.002 6.002 0 0 1-5 5.917V20a1 1 0 1 1-2 0v-5.083A6.002 6.002 0 0 1 12 3a6 6 0 0 1 6 6m-6 4a4 4 0 1 0 0-8a4 4 0 0 0 0 8" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinBottom(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M12 11a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/><path fill-rule="evenodd" d="M18 9a6.002 6.002 0 0 1-5 5.917V20h3a1 1 0 1 1 0 2H8a1 1 0 1 1 0-2h3v-5.083A6.002 6.002 0 0 1 12 3a6 6 0 0 1 6 6m-6 4a4 4 0 1 0 0-8a4 4 0 0 0 0 8" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinTop(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M12 14a2 2 0 1 1 0 4a2 2 0 0 1 0-4"/><path fill-rule="evenodd" d="M8 5a1 1 0 0 1 0-2h8a1 1 0 1 1 0 2h-3v5.083A6.002 6.002 0 0 1 12 22a6 6 0 0 1-1-11.917V5zm4 7a4 4 0 1 1 0 8a4 4 0 0 1 0-8" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayBackwards(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 7h3v10H2zm4 5l7.002-5v10zm15.002-5L14 12l7.002 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayButton(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15 12.33l-6 4.33V8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayButtonO(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M12 21a9 9 0 1 0 0-18a9 9 0 0 0 0 18m0 2c6.075 0 11-4.925 11-11S18.075 1 12 1S1 5.925 1 12s4.925 11 11 11" clip-rule="evenodd"/><path d="m16 12l-6 4.33V7.67z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayButtonR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M19 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2M5 1a4 4 0 0 0-4 4v14a4 4 0 0 0 4 4h14a4 4 0 0 0 4-4V5a4 4 0 0 0-4-4z" clip-rule="evenodd"/><path d="m16 12l-6 4.33V7.67z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayForwards(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.002 17h-3V7h3zm-4-5L10 17V7zM2 17l7.002-5L2 7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayList(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 5H4v2h12zm0 4H4v2h12zM4 13h8v2H4zm16 3l-6-3v6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayListAdd(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M2 5h12v2H2zm0 4h12v2H2zm8 4H2v2h8z"/><path d="M16 9h2v4h4v2h-4v4h-2v-4h-4v-2h4z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayListCheck(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 6H3v2h12zm0 4H3v2h12zM3 14h8v2H3zm8.99 1.025l1.415-1.414l2.121 2.121l4.243-4.242l1.414 1.414l-5.657 5.657z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayListRemove(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.964 4.634h-12v2h12zm0 4h-12v2h12zm-12 4h8v2h-8zm9.001 1.076l1.414-1.415l2.121 2.121l2.121-2.12l1.415 1.413l-2.122 2.122l2.122 2.12l-1.415 1.415l-2.121-2.121l-2.121 2.121l-1.415-1.414l2.122-2.122z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayListSearch(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M15.879 4.879h-12v2h12zm0 4h-12v2h12zm-12 4h8v2h-8z"/><path fill-rule="evenodd" d="M13.757 12.757a3 3 0 0 0 3.415 4.83l1.535 1.534l1.414-1.414l-1.535-1.535a3.001 3.001 0 0 0-4.829-3.415m1.415 2.829a1 1 0 1 0 1.414-1.415a1 1 0 0 0-1.414 1.415" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayPause(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 7H8v10h3zm2 10h3V7h-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayPauseO(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M9 9h2v6H9zm6 6h-2V9h2z"/><path fill-rule="evenodd" d="M23 12c0 6.075-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1s11 4.925 11 11m-2 0a9 9 0 1 1-18 0a9 9 0 0 1 18 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayPauseR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M9 9h2v6H9zm6 6h-2V9h2z"/><path fill-rule="evenodd" d="M1 5a4 4 0 0 1 4-4h14a4 4 0 0 1 4 4v14a4 4 0 0 1-4 4H5a4 4 0 0 1-4-4zm4-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayStop(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 7h10v10H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayStopO(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M15 9H9v6h6z"/><path fill-rule="evenodd" d="M23 12c0 6.075-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1s11 4.925 11 11m-2 0a9 9 0 1 1-18 0a9 9 0 0 1 18 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayStopR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M15 9H9v6h6z"/><path fill-rule="evenodd" d="M1 5a4 4 0 0 1 4-4h14a4 4 0 0 1 4 4v14a4 4 0 0 1-4 4H5a4 4 0 0 1-4-4zm4-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayTrackNext(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m6 17l8-5l-8-5zM18 7h-3v10h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayTrackNextO(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 21a9 9 0 1 1 0-18a9 9 0 0 1 0 18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m13-3a1 1 0 1 1 2 0v6a1 1 0 1 1-2 0zm-1 3l-6 3.464V8.536z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayTrackNextR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M15 9a1 1 0 1 1 2 0v6a1 1 0 1 1-2 0zm-1 3l-6 3.464V8.536z"/><path fill-rule="evenodd" d="M1 5a4 4 0 0 1 4-4h14a4 4 0 0 1 4 4v14a4 4 0 0 1-4 4H5a4 4 0 0 1-4-4zm4-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayTrackPrev(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m18 17l-8-5l8-5zM6 7h3v10H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayTrackPrevO(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M8 8a1 1 0 0 1 1 1v6a1 1 0 1 1-2 0V9a1 1 0 0 1 1-1m8 7.464L10 12l6-3.464z"/><path d="M3 12a9 9 0 1 0 18 0a9 9 0 0 0-18 0m9-11C5.925 1 1 5.925 1 12s4.925 11 11 11s11-4.925 11-11S18.075 1 12 1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayTrackPrevR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M8 8a1 1 0 0 1 1 1v6a1 1 0 1 1-2 0V9a1 1 0 0 1 1-1m8 7.464L10 12l6-3.464z"/><path d="M19 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2M5 1a4 4 0 0 0-4 4v14a4 4 0 0 0 4 4h14a4 4 0 0 0 4-4V5a4 4 0 0 0-4-4z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plug(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9 2a1 1 0 0 0-1 1v2a1 1 0 1 0 2 0V3a1 1 0 0 0-1-1M8 9h8v2a4 4 0 0 1-8 0zm5 7.917A6.002 6.002 0 0 0 18 11V7H6v4a6.002 6.002 0 0 0 5 5.917V22a1 1 0 1 0 2 0zM14 3a1 1 0 1 1 2 0v2a1 1 0 1 1-2 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pocket(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 4h18v7a9 9 0 1 1-18 0zM1 2h22v9c0 6.075-4.925 11-11 11S1 17.075 1 11zm10.293 12.694a1 1 0 0 0 1.414 0l4.243-4.243a1 1 0 1 0-1.414-1.414L12 12.572L8.464 9.037A1 1 0 0 0 7.05 10.45z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pokemon(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 12a9 9 0 1 1 18 0a9 9 0 0 1-18 0m2.07 1a7.002 7.002 0 0 0 13.86 0h-4.1a3.001 3.001 0 0 1-5.66 0zm13.86-2a7.001 7.001 0 0 0-13.86 0h4.1a3.001 3.001 0 0 1 5.66 0zM12 13a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Polaroid(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 4a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2zm2 0h14v11H5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Poll(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M19 4H5a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1M5 2a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3z" clip-rule="evenodd"/><path d="M11 7h2v10h-2zm4 6h2v4h-2zm-8-3h2v7H7z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Presentation(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10m8-10h-2a6 6 0 0 0-12 0H4a8 8 0 1 1 16 0M4.252 14h15.496a8.003 8.003 0 0 1-15.496 0M8 12a4 4 0 1 1 8 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Printer(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 4h8v2H8zm10 2h4v12h-4v4H6v-4H2V6h4V2h12zm2 10h-2v-2H6v2H4V8h16zM8 16h8v4H8zm0-6H6v2h2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ProductHunt(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 19a7 7 0 1 1 0-14a7 7 0 0 1 0 14m-9-7a9 9 0 1 1 18 0a9 9 0 0 1-18 0m6 4V8h4a3 3 0 1 1 0 6h-2v2zm5-5a1 1 0 0 1-1 1h-2v-2h2a1 1 0 0 1 1 1" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Profile(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M16 9a4 4 0 1 1-8 0a4 4 0 0 1 8 0m-2 0a2 2 0 1 1-4 0a2 2 0 0 1 4 0"/><path d="M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11s11-4.925 11-11S18.075 1 12 1M3 12c0 2.09.713 4.014 1.908 5.542A8.986 8.986 0 0 1 12.065 14a8.984 8.984 0 0 1 7.092 3.458A9 9 0 1 0 3 12m9 9a8.963 8.963 0 0 1-5.672-2.012A6.992 6.992 0 0 1 12.065 16a6.991 6.991 0 0 1 5.689 2.92A8.964 8.964 0 0 1 12 21"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PullClear(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M4 6H2v10a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V6h-2v10H4z"/><path d="M6 12h12v2H6zm12-4H6v2h12z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PushChevronDown(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 7.414L6.414 6l5.657 5.657L17.728 6l1.414 1.414l-7.07 7.071zm14 8.929H5v2h14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PushChevronDownO(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M16 14v2H8v-2zM7.757 8.703l1.415-1.415L12 10.117l2.828-2.829l1.415 1.415L12 12.945z"/><path fill-rule="evenodd" d="M12 23c6.075 0 11-4.925 11-11S18.075 1 12 1S1 5.925 1 12s4.925 11 11 11m0-2a9 9 0 1 0 0-18a9 9 0 0 0 0 18" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PushChevronDownR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M9.172 7.288L7.757 8.703L12 12.945l4.243-4.242l-1.415-1.415L12 10.117zM8 14h8v2H8z"/><path fill-rule="evenodd" d="M1 5a4 4 0 0 1 4-4h14a4 4 0 0 1 4 4v14a4 4 0 0 1-4 4H5a4 4 0 0 1-4-4zm4-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PushChevronLeft(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16.929 5l1.414 1.414l-5.657 5.657l5.657 5.657l-1.414 1.414l-7.071-7.07zM8 19V5H6v14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PushChevronLeftO(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M10 16H8V8h2zm5.297-8.243l1.415 1.415L13.883 12l2.829 2.828l-1.415 1.415L11.055 12z"/><path fill-rule="evenodd" d="M1 12c0 6.075 4.925 11 11 11s11-4.925 11-11S18.075 1 12 1S1 5.925 1 12m2 0a9 9 0 1 0 18 0a9 9 0 0 0-18 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PushChevronLeftR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M16.674 9.172L15.26 7.757L11.017 12l4.243 4.243l1.414-1.415L13.846 12zM9.963 8v8h-2V8z"/><path fill-rule="evenodd" d="M1 5a4 4 0 0 1 4-4h14a4 4 0 0 1 4 4v14a4 4 0 0 1-4 4H5a4 4 0 0 1-4-4zm4-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PushChevronRight(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.414 5L6 6.414l5.657 5.657L6 17.728l1.414 1.414l7.071-7.07zm8.929 14V5h2v14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PushChevronRightO(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M14 8h2v8h-2zm-5.297 8.243l-1.415-1.414L10.117 12L7.288 9.172l1.415-1.415L12.945 12z"/><path fill-rule="evenodd" d="M23 12c0-6.075-4.925-11-11-11S1 5.925 1 12s4.925 11 11 11s11-4.925 11-11m-2 0a9 9 0 1 0-18 0a9 9 0 0 0 18 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PushChevronRightR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m7.644 14.828l1.415 1.415L13.3 12L9.06 7.757L7.644 9.172L10.473 12zM14.356 16V8h2v8z"/><path fill-rule="evenodd" d="M1 5a4 4 0 0 1 4-4h14a4 4 0 0 1 4 4v14a4 4 0 0 1-4 4H5a4 4 0 0 1-4-4zm4-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PushChevronUp(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5 16.929l1.414 1.414l5.657-5.657l5.657 5.657l1.414-1.414l-7.07-7.071zM19 8H5V6h14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PushChevronUpO(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18m0-2C5.925 1 1 5.925 1 12s4.925 11 11 11s11-4.925 11-11S18.075 1 12 1m-4 9V8h8v2zm8.243 5.297l-1.414 1.415L12 13.883l-2.828 2.829l-1.415-1.415L12 11.055z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PushChevronUpR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 21h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2m-4-2a4 4 0 0 0 4 4h14a4 4 0 0 0 4-4V5a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4zm8.172-2.288l-1.415-1.415L12 11.055l4.243 4.242l-1.415 1.415L12 13.883zM8 10h8V8H8z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PushDown(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 1h2v14.485l3.243-3.242l1.414 1.414L12 19.314l-5.657-5.657l1.414-1.414L11 15.485zm7 19.288H6v2h12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PushLeft(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.288 11v2H7.802l3.243 3.243l-1.414 1.414L3.974 12L9.63 6.343l1.414 1.414L7.802 11zM3 18V6H1v12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PushRight(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 13v-2h14.485l-3.242-3.243l1.414-1.414L19.314 12l-5.657 5.657l-1.414-1.415L15.485 13zm19.288-7v12h2V6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PushUp(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11 22.288h2V7.802l3.243 3.243l1.414-1.414L12 3.974L6.343 9.63l1.414 1.414L11 7.802zM18 3H6V1h12z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Qr(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9 3H3v6h2V5h4zM3 21v-6h2v4h4v2zM15 3v2h4v4h2V3zm4 12h2v6h-6v-2h4zM7 7h4v4H7zm0 6h4v4H7zm10-6h-4v4h4zm-4 6h4v4h-4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Quote(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.135 9h3L10 14.608H7zm5 0h3L15 14.608h-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuoteO(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M20 5H4v14h16zM4 3a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2z" clip-rule="evenodd"/><path d="M9.067 9.196h3l-2.134 5.608h-3zm5 0h3l-2.134 5.608h-3z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RadioCheck(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 20a8 8 0 1 0 0-16a8 8 0 0 0 0 16m0 2c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RadioChecked(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M12 16a4 4 0 1 0 0-8a4 4 0 0 0 0 8"/><path fill-rule="evenodd" d="M22 12c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2s10 4.477 10 10m-2 0a8 8 0 1 1-16 0a8 8 0 0 1 16 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ratio(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M4 6v6h2V8h4V6zm16 12h-6v-2h4v-4h2z"/><path fill-rule="evenodd" d="M4 2a4 4 0 0 0-4 4v12a4 4 0 0 0 4 4h16a4 4 0 0 0 4-4V6a4 4 0 0 0-4-4zm16 2H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Read(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4m3.465-4a4.002 4.002 0 0 0-7.339 1H2a1 1 0 1 0 0 2h1.126A4.002 4.002 0 0 0 11 12h2a4 4 0 0 0 7.874 1H22a1 1 0 1 0 0-2h-1.126a4.002 4.002 0 0 0-7.339-1zM15 12a2 2 0 1 0 4 0a2 2 0 0 0-4 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Readme(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 5.5h5a2 2 0 0 1 2 2v9a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1v-10a1 1 0 0 1 1-1m10 14c-.35 0-.687-.06-1-.17v.17a1 1 0 1 1-2 0v-.17c-.313.11-.65.17-1 .17H4a3 3 0 0 1-3-3v-10a3 3 0 0 1 3-3h5a3.99 3.99 0 0 1 3 1.354A3.99 3.99 0 0 1 15 3.5h5a3 3 0 0 1 3 3v10a3 3 0 0 1-3 3zm-1-12v9a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1v-10a1 1 0 0 0-1-1h-5a2 2 0 0 0-2 2m-8 0h4v2H5zm10 0h4v2h-4zm4 3h-4v2h4zm-14 0h4v2H5zm14 3h-4v2h4zm-14 0h4v2H5z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Record(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M12 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6"/><path fill-rule="evenodd" d="M22 12c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2s10 4.477 10 10m-2 0a8 8 0 1 1-16 0a8 8 0 0 1 16 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Redo(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.146 11.05l-.174-1.992l2.374-.208a5 5 0 1 0 .82 6.173l2.002.5a7 7 0 1 1-1.315-7.996l-.245-2.803L18.6 4.55l.523 5.977z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Remote(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m17.051 4.322l1.415 1.414l-4.243 4.243l4.243 4.242l-1.415 1.415l-5.656-5.657zM6.949 19.679l-1.415-1.415l4.243-4.242l-4.243-4.243L6.95 8.365l5.656 5.657z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Remove(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M8 11a1 1 0 1 0 0 2h8a1 1 0 1 0 0-2z"/><path fill-rule="evenodd" d="M23 12c0 6.075-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1s11 4.925 11 11m-2 0a9 9 0 1 1-18 0a9 9 0 0 1 18 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RemoveR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M8 11a1 1 0 1 0 0 2h8a1 1 0 1 0 0-2z"/><path fill-rule="evenodd" d="M1 5a4 4 0 0 1 4-4h14a4 4 0 0 1 4 4v14a4 4 0 0 1-4 4H5a4 4 0 0 1-4-4zm4-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rename(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M10 4H8v2H5a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3h3v2h2zM8 8v8H5a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1z" clip-rule="evenodd"/><path d="M19 16h-7v2h7a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3h-7v2h7a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reorder(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-opacity=".5" d="M3 4a1 1 0 0 1 1-1h8a1 1 0 1 1 0 2H4a1 1 0 0 1-1-1m0 8a1 1 0 0 1 1-1h8a1 1 0 1 1 0 2H4a1 1 0 0 1-1-1m0 4a1 1 0 0 1 1-1h8a1 1 0 1 1 0 2H4a1 1 0 0 1-1-1m0 4a1 1 0 0 1 1-1h8a1 1 0 1 1 0 2H4a1 1 0 0 1-1-1"/><path fill-rule="evenodd" d="M15.17 9a3.001 3.001 0 1 0 0-2H4a1 1 0 0 0 0 2zM19 8a1 1 0 1 1-2 0a1 1 0 0 1 2 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Repeat(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m18.37 8l-4.5 2.598V9H6.89v4h-2V7h8.98V5.402zm-8.24 9h8.98v-6h-2v4h-6.98v-1.598L5.63 16l4.5 2.598z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ring(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M6.343 3.686A8.015 8.015 0 0 1 7.936 2.45a8.014 8.014 0 0 1 8.128 0a8.016 8.016 0 0 1 1.593 1.236L12 9.343zM12 6.514L9.413 3.927a6.017 6.017 0 0 1 5.174 0z" clip-rule="evenodd"/><path d="M2 12.658a9.98 9.98 0 0 1 3.692-7.76l1.423 1.424a8 8 0 1 0 9.77 0l1.423-1.424A9.98 9.98 0 0 1 22 12.658c0 5.522-4.477 10-10 10s-10-4.478-10-10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RowFirst(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M6 11a1 1 0 1 0 0 2h8a1 1 0 1 0 0-2zm0 4a1 1 0 1 0 0 2h8a1 1 0 1 0 0-2z" opacity=".5"/><path d="M5 8a1 1 0 0 1 1-1h12a1 1 0 1 1 0 2H6a1 1 0 0 1-1-1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RowLast(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M6 13a1 1 0 1 1 0-2h8a1 1 0 1 1 0 2zm0-4a1 1 0 0 1 0-2h8a1 1 0 1 1 0 2z" opacity=".5"/><path d="M5 16a1 1 0 0 0 1 1h12a1 1 0 1 0 0-2H6a1 1 0 0 0-1 1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ruler(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 6a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h18a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3zm6 2H7v5a1 1 0 1 1-2 0V8H3a1 1 0 0 0-1 1v7a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V9a1 1 0 0 0-1-1h-2v3a1 1 0 1 1-2 0V8h-2v5a1 1 0 1 1-2 0V8h-2v3a1 1 0 1 1-2 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SandClock(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M13 6h-2v1a1 1 0 1 0 2 0z"/><path fill-rule="evenodd" d="M6 2v2h1v3a5 5 0 0 0 5 5a5 5 0 0 0-5 5v3H6v2h12v-2h-1v-3a5 5 0 0 0-5-5a5 5 0 0 0 5-5V4h1V2zm3 2h6v3a3 3 0 1 1-6 0zm0 13v3h6v-3a3 3 0 1 0-6 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Scan(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 3h2v18h-2zM5 8a1 1 0 0 1 1-1h3V5H6a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h3v-2H6a1 1 0 0 1-1-1zm14 0a1 1 0 0 0-1-1h-3V5h3a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3h-3v-2h3a1 1 0 0 0 1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Screen(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11 17H4a3 3 0 0 1-3-3V6a3 3 0 0 1 3-3h16a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3h-7v2h3a1 1 0 1 1 0 2H8a1 1 0 1 1 0-2h3zM4 5h16a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScreenMirror(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M5 8h14v6h-3v2h5V6H3v10h5v-2H5z"/><path d="M16.33 19L12 13l-4.33 6z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScreenShot(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M12 8V6H8v4h2V8zm2 6h2v4h-4v-2h2z"/><path fill-rule="evenodd" d="M4 3a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1zm2 17V4h12v16z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScreenWide(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11 16H3a3 3 0 0 1-3-3V8a3 3 0 0 1 3-3h18a3 3 0 0 1 3 3v5a3 3 0 0 1-3 3h-8v1h2a1 1 0 1 1 0 2H9a1 1 0 1 1 0-2h2zM3 7h18a1 1 0 0 1 1 1v5a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScrollH(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M7.182 9.172L5.768 7.757L1.525 12l4.243 4.243l1.414-1.415L4.353 12zm9.636 5.656l1.414 1.415L22.475 12l-4.243-4.243l-1.414 1.415L19.646 12z"/><path fill-rule="evenodd" d="M15 12a3 3 0 1 1-6 0a3 3 0 0 1 6 0m-2 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScrollV(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m9.172 16.818l-1.415 1.414L12 22.475l4.243-4.243l-1.415-1.414L12 19.647zm5.656-9.636l1.415-1.414L12 1.525L7.757 5.768l1.415 1.414L12 4.354z"/><path fill-rule="evenodd" d="M12 9a3 3 0 1 1 0 6a3 3 0 0 1 0-6m0 2a1 1 0 1 1 0 2a1 1 0 0 1 0-2" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Search(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M18.319 14.433A8.001 8.001 0 0 0 6.343 3.868a8 8 0 0 0 10.564 11.976l.043.045l4.242 4.243a1 1 0 1 0 1.415-1.415l-4.243-4.242a1.116 1.116 0 0 0-.045-.042m-2.076-9.15a6 6 0 1 1-8.485 8.485a6 6 0 0 1 8.485-8.485" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchFound(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M7.665 10.237L9.198 8.95l1.285 1.532l3.064-2.571l1.286 1.532l-4.596 3.857z"/><path fill-rule="evenodd" d="M16.207 4.893a8.001 8.001 0 0 1 .662 10.565c.016.013.03.027.045.042l4.243 4.243a1 1 0 0 1-1.414 1.414L15.5 16.914a1.046 1.046 0 0 1-.042-.045A8.001 8.001 0 0 1 4.893 4.893a8 8 0 0 1 11.314 0m-1.414 9.9a6 6 0 1 0-8.485-8.485a6 6 0 0 0 8.485 8.485" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchLoading(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M8.55 10.55a1 1 0 1 1-2 0a1 1 0 0 1 2 0m2 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2m3 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path fill-rule="evenodd" d="M16.207 4.893a8.001 8.001 0 0 1 .662 10.565c.016.013.03.027.045.042l4.243 4.243a1 1 0 0 1-1.414 1.414L15.5 16.914a.933.933 0 0 1-.042-.045A8.001 8.001 0 0 1 4.893 4.893a8 8 0 0 1 11.314 0m-9.9 9.9a6 6 0 1 0 8.486-8.485a6 6 0 0 0-8.485 8.485" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Select(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m6 9.657l1.414 1.414l4.243-4.243l4.243 4.243l1.414-1.414L11.657 4zm0 4.786l1.414-1.414l4.243 4.243l4.243-4.243l1.414 1.414l-5.657 5.657z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SelectO(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m9.172 11.508l-1.415-1.414L12 5.85l4.243 4.243l-1.415 1.414L12 8.68zm0 .984l-1.415 1.414L12 18.15l4.243-4.243l-1.415-1.414L12 15.32z"/><path fill-rule="evenodd" d="M23 12c0 6.075-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1s11 4.925 11 11m-2 0a9 9 0 1 1-18 0a9 9 0 0 1 18 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SelectR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m9.172 11.508l-1.415-1.414L12 5.85l4.243 4.243l-1.415 1.414L12 8.68zm0 .984l-1.415 1.414L12 18.15l4.243-4.243l-1.415-1.414L12 15.32z"/><path fill-rule="evenodd" d="M1 5a4 4 0 0 1 4-4h14a4 4 0 0 1 4 4v14a4 4 0 0 1-4 4H5a4 4 0 0 1-4-4zm4-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Server(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M9 6a1 1 0 0 0 0 2h6a1 1 0 1 0 0-2zm0 4a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2zm4 7a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/><path fill-rule="evenodd" d="M4 5a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3v14a3 3 0 0 1-3 3H7a3 3 0 0 1-3-3zm3-1h10a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Serverless(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.787 6H5v3h5.695zm-1.82 5H5v3h3.875zm1.037 3l1.092-3H20v3zm-2.856 2H5v3h2.056zm1.036 3l1.092-3H20v3zm3.64-10l1.092-3H20v3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShapeCircle(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 17a5 5 0 1 0 0-10a5 5 0 0 0 0 10m0 3a8 8 0 1 0 0-16a8 8 0 0 0 0 16" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShapeHalfCircle(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 4a8 8 0 1 0 0 16v-3a5 5 0 0 1 0-10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShapeHexagon(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m6 15.235l6 3.333l6-3.333v-6.47l-6-3.333l-6 3.333zM12 2L3 7v10l9 5l9-5V7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShapeRhombus(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 6.343L6.343 12L12 17.657L17.657 12zM2.1 12l9.9 9.9l9.9-9.9L12 2.1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShapeSquare(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M17 7H7v10h10zM4 4v16h16V4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShapeTriangle(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.66 5L3 20h17.32zm0 6l-3.464 6h6.928z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShapeZigzag(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.312 9L1 10.51l3.774 3.28l1.509 1.312l1.312-1.51l1.54-1.77l2.264 1.968l1.51 1.312l1.311-1.51l1.538-1.769l2.263 1.967l1.51 1.312l1.311-1.51l1.969-2.264l-1.51-1.312l-1.968 2.264L15.559 9l-1.312 1.51h.002l-1.538 1.77L8.937 9l-.883 1.016l-1.968 2.264z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Share(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 9a3 3 0 1 0-2.977-2.63l-6.94 3.47a3 3 0 1 0 0 4.319l6.94 3.47a3 3 0 1 0 .895-1.789l-6.94-3.47a3.03 3.03 0 0 0 0-.74l6.94-3.47C16.456 8.68 17.19 9 18 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shield(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7 8v5a5 5 0 0 0 10 0V8zM5 4h14v9a7 7 0 1 1-14 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBag(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 4h14a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1M2 5a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v14a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3zm10 7c-2.761 0-5-2.686-5-6h2c0 2.566 1.67 4 3 4s3-1.434 3-4h2c0 3.314-2.239 6-5 6" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingCart(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M5.792 2H1v2h3.218l2.77 12.678H7V17h13v-.248l2.193-9.661L22.531 6H6.655l-.57-2.611zm14.195 6H7.092l1.529 7h9.777z" clip-rule="evenodd"/><path d="M10 22a2 2 0 1 0 0-4a2 2 0 0 0 0 4m9-2a2 2 0 1 1-4 0a2 2 0 0 1 4 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shortcut(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M16.192 7.707a1 1 0 0 0-1.414 0l-7.07 7.071a1 1 0 1 0 1.413 1.414l7.071-7.07a1 1 0 0 0 0-1.415"/><path fill-rule="evenodd" d="M3 6a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm3-1h12a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shutterstock(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M12 17a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1h-1v4h-4zM11 6a1 1 0 0 1 1 1v1H8v4H7a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1z"/><path fill-rule="evenodd" d="M5 2a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3zm14 2H5a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sidebar(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M21 20H7V4h14zm-2-2H9V6h10z" clip-rule="evenodd"/><path d="M3 20h2V4H3z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SidebarOpen(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 4h18v16H3zm6 2h10v12H9z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SidebarRight(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M3 4h14v16H3zm2 2h10v12H5z" clip-rule="evenodd"/><path d="M21 4h-2v16h2z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Signal(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 7a1 1 0 1 1 2 0v10a1 1 0 1 1-2 0zm-8 8a1 1 0 1 1 2 0v2a1 1 0 1 1-2 0zm5-5a1 1 0 0 0-1 1v6a1 1 0 1 0 2 0v-6a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Size(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 6V4h8v16h-8v-2H8v-2H4V8h4V6zm2 0h4v12h-4zm-2 2h-2v8h2zm-4 2v4H6v-4z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sketch(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M5.209 3h13.694l1.209 7.253l-8.056 10.933L4 10.253zm1.694 2l-.791 4.747l5.944 8.067L18 9.747L17.209 5z" clip-rule="evenodd"/><path d="M8.056 8h8v2h-8z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Slack(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13 10a2 2 0 1 0 4 0V5a2 2 0 1 0-4 0zM5 8a2 2 0 1 0 0 4h5a2 2 0 1 0 0-4zm10 5a2 2 0 1 0 0 4h5a2 2 0 1 0 0-4zm-5 9a2 2 0 0 1-2-2v-5a2 2 0 1 1 4 0v5a2 2 0 0 1-2 2M8 5a2 2 0 1 1 4 0v2h-2a2 2 0 0 1-2-2M3 15a2 2 0 1 0 4 0v-2H5a2 2 0 0 0-2 2m14 5a2 2 0 1 1-4 0v-2h2a2 2 0 0 1 2 2m5-10a2 2 0 1 0-4 0v2h2a2 2 0 0 0 2-2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sleep(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0m-4.101 5A6.977 6.977 0 0 1 12 19a6.977 6.977 0 0 1-4.899-2zm1.427-2a7 7 0 1 0-12.653 0z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmartHomeBoiler(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 5a4 4 0 0 1 4-4h6a4 4 0 0 1 4 4v16h-3.856l.742 2h-2l-.742-2h-2l.742 2h-2l-.742-2H5zm4-2h6a2 2 0 0 1 2 2v2H7V5a2 2 0 0 1 2-2M7 9h10v10H7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmartHomeCooker(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M15 16a3 3 0 1 1-6 0a3 3 0 0 1 6 0m-2 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/><path d="M15 1H9v2h2v2H7a4 4 0 0 0-4 4v10a4 4 0 0 0 4 4h10a4 4 0 0 0 4-4V9a4 4 0 0 0-4-4h-4V3h2zm2 6H7a2 2 0 0 0-2 2h14a2 2 0 0 0-2-2m2 4H5v8a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmartHomeHeat(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M17 11H7a1 1 0 1 0 0 2h3v2H7a3 3 0 1 1 0-6h10a3 3 0 1 1 0 6h-3v-2h3a1 1 0 1 0 0-2"/><path fill-rule="evenodd" d="M0 12a7 7 0 0 1 7-7h10a7 7 0 1 1 0 14H7a7 7 0 0 1-7-7m7-5h10a5 5 0 0 1 0 10H7A5 5 0 0 1 7 7" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmartHomeLight(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M7.034 6.5a5 5 0 0 1 10 0v4a5 5 0 0 1-10 0zm8 0v4a3 3 0 0 1-6 0v-4a3 3 0 0 1 6 0" clip-rule="evenodd"/><path d="M12.034 16.5a1 1 0 0 0-1 1v4a1 1 0 1 0 2 0v-4a1 1 0 0 0-1-1m-4.29-.06a1 1 0 1 1 1.88.684l-1.368 3.759a1 1 0 1 1-1.88-.684zm7.23-.598a1 1 0 0 0-.598 1.282l1.369 3.759a1 1 0 1 0 1.879-.684l-1.368-3.76a1 1 0 0 0-1.282-.597"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmartHomeRefrigerator(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M9 6a1 1 0 0 1 2 0v2a1 1 0 1 1-2 0zm1 7a1 1 0 0 0-1 1v2a1 1 0 1 0 2 0v-2a1 1 0 0 0-1-1"/><path fill-rule="evenodd" d="M5 4a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3v16a3 3 0 0 1-3 3H8a3 3 0 0 1-3-3zm3-1h8a1 1 0 0 1 1 1v6H7V4a1 1 0 0 1 1-1m-1 9h10v8a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmartHomeWashMachine(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6 4h12a1 1 0 0 1 1 1v3H5V5a1 1 0 0 1 1-1m13 15v-9H5v9a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1M3 5a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm4 0a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2zm7 2a1 1 0 1 0 0-2a1 1 0 0 0 0 2m4-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-4 9a2 2 0 1 1-4 0a2 2 0 0 1 4 0m2 0a4 4 0 1 1-8 0a4 4 0 0 1 8 0" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Smartphone(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M13 16h-2v2h2z"/><path fill-rule="evenodd" d="M5 4a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2zm2 0h10v16H7z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmartphoneChip(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M9 22a1 1 0 1 1-2 0a1 1 0 0 1 2 0m4 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0m3 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path d="M9 2a1 1 0 1 1-2 0a1 1 0 0 1 2 0m4 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0m3 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2M9 22a1 1 0 1 1-2 0a1 1 0 0 1 2 0m4 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0m3 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2m7-7a1 1 0 1 0-2 0a1 1 0 0 0 2 0m0-4a1 1 0 1 0-2 0a1 1 0 0 0 2 0m-1-5a1 1 0 1 1 0 2a1 1 0 0 1 0-2M2 15a1 1 0 1 1 0 2a1 1 0 0 1 0-2m0-4a1 1 0 1 1 0 2a1 1 0 0 1 0-2m1-3a1 1 0 1 0-2 0a1 1 0 0 0 2 0m14-2H7a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1M7 4a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3zm7 6h-4v4h4zM8 8v8h8V8z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmartphoneRam(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M5 4a1 1 0 1 1-2 0a1 1 0 0 1 2 0m4 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0m3 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2m5-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0m3 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2M5 20a1 1 0 1 1-2 0a1 1 0 0 1 2 0m4 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0m3 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2m5-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0m3 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2M5 12a1 1 0 1 0 0-2a1 1 0 0 0 0 2m15 1a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/><path fill-rule="evenodd" d="M0 9a3 3 0 0 1 3-3h18a3 3 0 0 1 3 3v6a3 3 0 0 1-3 3H3a3 3 0 0 1-3-3zm3-1h18a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmartphoneShake(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M13 14h-2v2h2z"/><path fill-rule="evenodd" d="M8 7a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2zm2 0h4v10h-4z" clip-rule="evenodd"/><path d="M18 9h2v6h-2zM0 14h2v-4H0zm6 1H4V9h2zm18-5h-2v4h2z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Smile(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M16 13h-2a2 2 0 1 1-4 0H8a4 4 0 0 0 8 0m-6-3a1 1 0 1 1-2 0a1 1 0 0 1 2 0m5 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path fill-rule="evenodd" d="M22 12c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2s10 4.477 10 10m-2 0a8 8 0 1 1-16 0a8 8 0 0 1 16 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmileMouthOpen(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M12 17a4 4 0 0 0 4-4H8a4 4 0 0 0 4 4m-2-7a1 1 0 1 1-2 0a1 1 0 0 1 2 0m5 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path fill-rule="evenodd" d="M22 12c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2s10 4.477 10 10m-2 0a8 8 0 1 1-16 0a8 8 0 0 1 16 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmileNeutral(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M9 11a1 1 0 1 0 0-2a1 1 0 0 0 0 2m0 4a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2zm7-5a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/><path fill-rule="evenodd" d="M22 12c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2s10 4.477 10 10m-2 0a8 8 0 1 1-16 0a8 8 0 0 1 16 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmileNoMouth(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M10 10a1 1 0 1 1-2 0a1 1 0 0 1 2 0m5 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path fill-rule="evenodd" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10m0-2a8 8 0 1 0 0-16a8 8 0 0 0 0 16" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmileNone(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M8 9a1 1 0 0 0 0 2h1a1 1 0 1 0 0-2zm7 0a1 1 0 1 0 0 2h1a1 1 0 1 0 0-2zm-6 6a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2z"/><path fill-rule="evenodd" d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10m0-2a8 8 0 1 0 0-16a8 8 0 0 0 0 16" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmileSad(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M9 11a1 1 0 1 0 0-2a1 1 0 0 0 0 2m5 6a2 2 0 1 0-4 0H8a4 4 0 0 1 8 0zm2-7a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/><path fill-rule="evenodd" d="M22 12c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2s10 4.477 10 10m-2 0a8 8 0 1 1-16 0a8 8 0 0 1 16 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmileUpside(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M16 11h-2a2 2 0 1 0-4 0H8a4 4 0 1 1 8 0m-6 3a1 1 0 1 0-2 0a1 1 0 0 0 2 0m5-1a1 1 0 1 1 0 2a1 1 0 0 1 0-2"/><path fill-rule="evenodd" d="M22 12c0-5.523-4.477-10-10-10S2 6.477 2 12s4.477 10 10 10s10-4.477 10-10m-2 0a8 8 0 1 0-16 0a8 8 0 0 0 16 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SoftwareDownload(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M11 5a1 1 0 1 1 2 0v7.158l3.243-3.243l1.414 1.414L12 15.986L6.343 10.33l1.414-1.414L11 12.158z"/><path d="M4 14h2v4h12v-4h2v4a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SoftwareUpload(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M11 14.986a1 1 0 1 0 2 0V7.828l3.243 3.243l1.414-1.414L12 4L6.343 9.657l1.414 1.414L11 7.83z"/><path d="M4 14h2v4h12v-4h2v4a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortAz(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 8a1 1 0 0 1 1-1h10a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1m2 4a1 1 0 0 1 1-1h6a1 1 0 1 1 0 2H9a1 1 0 0 1-1-1m3 3a1 1 0 1 0 0 2h2a1 1 0 1 0 0-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortZa(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 16a1 1 0 0 0 1 1h10a1 1 0 1 0 0-2H7a1 1 0 0 0-1 1m2-4a1 1 0 0 0 1 1h6a1 1 0 1 0 0-2H9a1 1 0 0 0-1 1m3-3a1 1 0 1 1 0-2h2a1 1 0 1 1 0 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpaceBetween(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 5h-4v14h4v-2h-2V7h2zM5 5h4v14H5v-2h2V7H5zm8 2v10h-2V7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpaceBetweenV(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 5v4h14V5h-2v2H7V5zm0 14v-4h14v4h-2v-2H7v2zm2-8h10v2H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spectrum(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 16h4a8 8 0 0 0-8-8v4a4 4 0 0 1 4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spinner(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M12 19a7 7 0 1 0 0-14a7 7 0 0 0 0 14m0 3c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10" clip-rule="evenodd" opacity=".2"/><path d="M2 12C2 6.477 6.477 2 12 2v3a7 7 0 0 0-7 7z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpinnerAlt(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 12C2 6.477 6.477 2 12 2v3a7 7 0 0 0-7 7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpinnerTwo(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M12 19a7 7 0 1 0 0-14a7 7 0 0 0 0 14m0 3c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10" clip-rule="evenodd" opacity=".2"/><path d="M12 22c5.523 0 10-4.477 10-10h-3a7 7 0 0 1-7 7zM2 12C2 6.477 6.477 2 12 2v3a7 7 0 0 0-7 7z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpinnerTwoAlt(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 22c5.523 0 10-4.477 10-10h-3a7 7 0 0 1-7 7zM2 12C2 6.477 6.477 2 12 2v3a7 7 0 0 0-7 7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Square(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M14 10h-4v4h4z"/><path fill-rule="evenodd" d="M5 9a4 4 0 0 1 4-4h6a4 4 0 0 1 4 4v6a4 4 0 0 1-4 4H9a4 4 0 0 1-4-4zm4-1h6a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1H9a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stack(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M20 4v12h2V2H8v2z"/><path fill-rule="evenodd" d="M2 8v14h14V8zm12 2H4v10h10z" clip-rule="evenodd"/><path d="M17 7H5V5h14v14h-2z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stark(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.172 18.025a8 8 0 0 1 4.935-14.948l-.437 3.126a4.844 4.844 0 0 0-2.988 9.05l6.146-11.278l2.634 1.436a8 8 0 0 1-4.934 14.948l.436-3.126a4.844 4.844 0 0 0 2.988-9.05L9.806 19.46z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stopwatch(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m18.621 2.55l2.829 2.83l-1.414 1.414l-2.829-2.828zM12.823 8.6h-2v4h2z"/><path fill-rule="evenodd" d="M5.186 18.814A9 9 0 1 0 17.914 6.086A9 9 0 0 0 5.186 18.814m1.415-1.415A7 7 0 1 0 16.5 7.5a7 7 0 0 0-9.9 9.9" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stories(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M15 6H9a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1M9 4a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h6a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3z" clip-rule="evenodd"/><path d="M2 6a1 1 0 0 1 2 0v12a1 1 0 1 1-2 0zm18 0a1 1 0 1 1 2 0v12a1 1 0 1 1-2 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Studio(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M17 13h-4v4h4z"/><path fill-rule="evenodd" d="M3 3h18v18H3zm2 2h14v14H5z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Style(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M13 21v-8h8v8zm2-6h4v4h-4zM3 11V3h8v8zm2-6h4v4H5z" clip-rule="evenodd"/><path d="M18 6v6h-2V8h-4V6zm-6 12H6v-6h2v4h4z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sun(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 16a4 4 0 1 0 0-8a4 4 0 0 0 0 8m0 2a6 6 0 1 0 0-12a6 6 0 0 0 0 12M11 0h2v4.062a8.079 8.079 0 0 0-2 0zM7.094 5.68L4.222 2.808L2.808 4.222L5.68 7.094A8.048 8.048 0 0 1 7.094 5.68M4.062 11H0v2h4.062a8.079 8.079 0 0 1 0-2m1.618 5.906l-2.872 2.872l1.414 1.414l2.872-2.872a8.048 8.048 0 0 1-1.414-1.414M11 19.938V24h2v-4.062a8.069 8.069 0 0 1-2 0m5.906-1.618l2.872 2.872l1.414-1.414l-2.872-2.872a8.048 8.048 0 0 1-1.414 1.414M19.938 13H24v-2h-4.062a8.069 8.069 0 0 1 0 2M18.32 7.094l2.872-2.872l-1.414-1.414l-2.872 2.872c.528.41 1.003.886 1.414 1.414" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Support(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.26 21.997a10.276 10.276 0 0 1-.52 0a10.004 10.004 0 0 1-8.983-6.173a10.034 10.034 0 0 1 .017-7.69A10.015 10.015 0 0 1 4.908 4.95l.042-.042a10.015 10.015 0 0 1 3.167-2.126a10.034 10.034 0 0 1 7.753-.006a10.015 10.015 0 0 1 3.186 2.138l.03.03c.913.917 1.65 2.01 2.153 3.223a10.012 10.012 0 0 1 .76 3.985a10.004 10.004 0 0 1-6.226 9.112a10.013 10.013 0 0 1-3.512.733m1.772-6.55l2.874 2.873a8.004 8.004 0 0 1-9.812 0l2.874-2.874a4.007 4.007 0 0 0 4.064 0m-5.478-1.415L5.68 16.906a8.004 8.004 0 0 1 0-9.812l2.874 2.874a4.007 4.007 0 0 0 0 4.064m1.528-1.463a2.01 2.01 0 0 1-.014-1.087a1.99 1.99 0 0 1 .518-.896a1.99 1.99 0 0 1 1.932-.518c.328.088.639.26.896.518a1.99 1.99 0 0 1 .518 1.932c-.088.328-.26.639-.518.896a1.99 1.99 0 0 1-1.932.518a1.991 1.991 0 0 1-.896-.518a1.99 1.99 0 0 1-.504-.845m3.95-4.015a4.007 4.007 0 0 0-4.064 0L7.094 5.68a8.004 8.004 0 0 1 9.812 0zm4.288 8.352a8.004 8.004 0 0 0 0-9.812l-2.874 2.874a4.007 4.007 0 0 1 0 4.064z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Swap(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 13v-1.5h-6v-2h6V8l3 2.5zm-8 4v-1.5h6v-2H8V12l-3 2.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwapVertical(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 16h1.5v-6h2v6H17l-2.5 3zM8 8h1.5v6h2V8H13l-2.5-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sweden(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23 4H10v7h13zm0 9v7H10v-7zM8 13v7H1v-7zm-7-2V4h7v7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Swiss(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 3v18h18V3zm11 4h-4v3H7v4h3v3h4v-3h3v-4h-3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sync(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.56 10.642L6.355 3.95l1.9 1.9a9.004 9.004 0 0 1 11.156 1.256l-1.414 1.415a7.003 7.003 0 0 0-8.28-1.21l1.537 1.538zm14.88 2.716l-1.794 6.692l-1.9-1.9A9.003 9.003 0 0 1 4.59 16.894l1.414-1.415a7.003 7.003 0 0 0 8.28 1.21l-1.537-1.538z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tab(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M19 4a3 3 0 0 1 3 3v10a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V7a3 3 0 0 1 3-3zm1 5.625h-7c-.552 0-1.156-.42-1.348-.938L10.654 6H5a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tag(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M2 8v8a1 1 0 0 0 1 1h13.62a1 1 0 0 0 .76-.35l3.428-4a1 1 0 0 0 0-1.3l-3.428-4a1 1 0 0 0-.76-.35H3a1 1 0 0 0-1 1M0 8v8a3 3 0 0 0 3 3h13.62a3 3 0 0 0 2.278-1.048l3.428-4a3 3 0 0 0 0-3.904l-3.428-4A3 3 0 0 0 16.62 5H3a3 3 0 0 0-3 3"/><path d="M15 13a1 1 0 1 0 0-2a1 1 0 0 0 0 2m0 2a3 3 0 1 0 0-6a3 3 0 0 0 0 6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tally(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.661 2.671a.63.63 0 0 1 .595-.49h2c.276 0 .457.219.405.49L5.34 9.545a.63.63 0 0 1-.595.491h-2a.397.397 0 0 1-.405-.49zm5 0a.63.63 0 0 1 .595-.49h2c.276 0 .457.219.405.49L8.071 21.33a.63.63 0 0 1-.594.491h-2a.397.397 0 0 1-.405-.491zm5 0a.63.63 0 0 1 .595-.49h2c.276 0 .457.219.405.49l-3.59 18.658a.63.63 0 0 1-.594.491h-2a.397.397 0 0 1-.405-.491zm5 0a.63.63 0 0 1 .595-.49h2c.276 0 .457.219.405.49L20.34 9.545a.63.63 0 0 1-.595.491h-2a.397.397 0 0 1-.405-.49z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TapDouble(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M14.924 18v-4a3 3 0 0 0-6 0v4a3 3 0 1 0 6 0m-3-9a5 5 0 0 0-5 5v4a5 5 0 0 0 10 0v-4a5 5 0 0 0-5-5" clip-rule="evenodd"/><path d="M10.924 14a1 1 0 1 1 2 0v3h-2zm1-13a9.97 9.97 0 0 1 7.105 2.963l-1.415 1.414A7.976 7.976 0 0 0 11.924 3c-2.15 0-4.1.847-5.538 2.227L4.97 3.812A9.967 9.967 0 0 1 11.924 1"/><path fill-rule="evenodd" d="M11.923 5a6.97 6.97 0 0 1 4.38 1.539l-1.426 1.426A4.978 4.978 0 0 0 11.923 7c-1.075 0-2.071.34-2.886.917l-1.43-1.429A6.97 6.97 0 0 1 11.924 5" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TapSingle(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M12.05 3.114c2.143 0 4.09.843 5.526 2.216L16.16 6.744a5.98 5.98 0 0 0-4.112-1.63a5.98 5.98 0 0 0-4.21 1.725L6.424 5.425a7.974 7.974 0 0 1 5.625-2.311m-1.072 8.772a1 1 0 1 1 2 0v2h-2z"/><path fill-rule="evenodd" d="M11.977 6.886a5 5 0 0 0-5 5v4a5 5 0 0 0 10 0v-4a5 5 0 0 0-5-5m3 9v-4a3 3 0 0 0-6 0v4a3 3 0 0 0 6 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Template(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M3 3v6h18V3zm16 2H5v2h14zM3 11v10h8V11zm6 2H5v6h4z" clip-rule="evenodd"/><path d="M21 11h-8v2h8zm-8 4h8v2h-8zm8 4h-8v2h8z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tennis(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M19.071 19.071c3.905-3.905 3.905-10.237 0-14.142c-3.905-3.905-10.237-3.905-14.142 0c-3.905 3.905-3.905 10.237 0 14.142c3.905 3.905 10.237 3.905 14.142 0m.872-8.03a7.966 7.966 0 0 0-2.286-4.698a7.966 7.966 0 0 0-4.717-2.288l-.01.056a11.011 11.011 0 0 1-8.819 8.819l-.056.01a7.966 7.966 0 0 0 2.288 4.717a7.966 7.966 0 0 0 4.698 2.286l.012-.07a11.011 11.011 0 0 1 8.819-8.82zm-.071 2.388v-.334a9.013 9.013 0 0 0-6.777 6.777h.334a7.964 7.964 0 0 0 4.228-2.215a7.963 7.963 0 0 0 2.215-4.228m-15.76-2.54v-.223a7.963 7.963 0 0 1 2.231-4.323a7.964 7.964 0 0 1 4.323-2.232h.222a9.013 9.013 0 0 1-6.777 6.777" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Terminal(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="m5.033 14.828l1.415 1.415L10.69 12L6.448 7.757L5.033 9.172L7.862 12zM15 14h-4v2h4z"/><path fill-rule="evenodd" d="M2 2a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h20a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2zm20 2H2v16h20z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Terrain(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8 10l-5 8h10zm2.529.754L13.5 6L21 18h-5.943z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Thermometer(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M16.95 5.636a1 1 0 1 1 1.414 1.414l-7.071 7.071a1 1 0 1 1-1.414-1.414z"/><path fill-rule="evenodd" d="M7.828 17.586a5.002 5.002 0 0 0 6.293-.636l7.071-7.071a5 5 0 1 0-7.07-7.071L7.05 9.878a5.002 5.002 0 0 0-.636 6.294l-3.606 3.606a1 1 0 1 0 1.414 1.415zm4.88-2.05l7.07-7.071a3 3 0 1 0-4.242-4.243l-7.071 7.071a3 3 0 1 0 4.242 4.243" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Thermostat(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M12 19a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path fill-rule="evenodd" d="M15 14a5 5 0 1 1-6 0V4a3 3 0 1 1 6 0zM13 4v11.17A3.001 3.001 0 0 1 12 21a3 3 0 0 1-1-5.83V4a1 1 0 1 1 2 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tikcode(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M9 5H5v4h4zM3 3v8h8V3zm16 2h-4v4h4zm-6-2v8h8V3zM9 15H5v4h4zm-6-2v8h8v-8z" clip-rule="evenodd"/><path d="M13 13h2v8h-2zm3 0h2v8h-2zm3 0h2v8h-2z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Time(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M9 7h2v5h5v2H9z"/><path fill-rule="evenodd" d="M22 12c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2s10 4.477 10 10m-2 0a8 8 0 1 1-16 0a8 8 0 0 1 16 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Timelapse(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M12 19a6.978 6.978 0 0 1-4.95-2.05L12 12V5a7 7 0 1 1 0 14"/><path fill-rule="evenodd" d="M23 12c0 6.075-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1s11 4.925 11 11m-2 0a9 9 0 1 1-18 0a9 9 0 0 1 18 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Timer(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M13 5.07A7.002 7.002 0 0 1 12 19A7 7 0 0 1 7.262 6.847L5.847 5.432A9 9 0 1 0 11 3.055v6.03h2z"/><path d="M7.707 8.707a1 1 0 0 0 0 1.414l2.829 2.829a1 1 0 0 0 1.414-1.414L9.12 8.707a1 1 0 0 0-1.414 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Today(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><rect width="10" height="10" x="7" y="9" opacity=".5" rx="1"/><path fill-rule="evenodd" d="M18 3H6a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1M6 1a3 3 0 0 0-3 3v16a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V4a3 3 0 0 0-3-3z" clip-rule="evenodd"/><path d="M7 6a1 1 0 0 1 1-1h4a1 1 0 1 1 0 2H8a1 1 0 0 1-1-1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToggleOff(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M17 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6"/><path fill-rule="evenodd" d="M0 12a7 7 0 0 1 7-7h10a7 7 0 1 1 0 14H7a7 7 0 0 1-7-7m7-5h10a5 5 0 0 1 0 10H7A5 5 0 0 1 7 7" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToggleOn(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M7 15a3 3 0 1 1 0-6a3 3 0 0 1 0 6"/><path fill-rule="evenodd" d="M24 12a7 7 0 0 0-7-7H7a7 7 0 0 0 0 14h10a7 7 0 0 0 7-7m-7-5H7a5 5 0 0 0 0 10h10a5 5 0 0 0 0-10" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToggleSquare(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M9 9a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1z"/><path fill-rule="evenodd" d="M24 7a2 2 0 0 0-2-2H2a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h20a2 2 0 0 0 2-2zm-2 0H2v10h20z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToggleSquareOff(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M15 9a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1z"/><path fill-rule="evenodd" d="M0 7a2 2 0 0 1 2-2h20a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2zm2 0h20v10H2z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToolbarBottom(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M18 13H6v2h12z"/><path fill-rule="evenodd" d="M2 8a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3zm3-1h14a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToolbarLeft(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M8 9H6v6h2z"/><path fill-rule="evenodd" d="M2 8a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3zm3-1h14a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToolbarRight(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M16 9h2v6h-2z"/><path fill-rule="evenodd" d="M22 8a3 3 0 0 0-3-3H5a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3zm-3-1H5a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V8a1 1 0 0 0-1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToolbarTop(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M18 11H6V9h12z"/><path fill-rule="evenodd" d="M2 16a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V8a3 3 0 0 0-3-3H5a3 3 0 0 0-3 3zm3 1h14a1 1 0 0 0 1-1V8a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Toolbox(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M17 5.5h3a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-12a2 2 0 0 1 2-2h3a3 3 0 0 1 3-3h4a3 3 0 0 1 3 3m-3-1h-4a1 1 0 0 0-1 1h6a1 1 0 0 0-1-1m6 3H4v2h16zm-16 12v-8h3v2h4v-2h2v2h4v-2h3v8z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Touchpad(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M20 21a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3H4a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3zM4 5h16a1 1 0 0 1 1 1v8H3V6a1 1 0 0 1 1-1M3 16v2a1 1 0 0 0 1 1h7v-3zm10 3v-3h8v2a1 1 0 0 1-1 1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Track(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M12 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6"/><path fill-rule="evenodd" d="M12 3a1 1 0 0 1 1 1v1.07A7.004 7.004 0 0 1 18.93 11H20a1 1 0 1 1 0 2h-1.07A7.004 7.004 0 0 1 13 18.93V20a1 1 0 1 1-2 0v-1.07A7.004 7.004 0 0 1 5.07 13H4a1 1 0 1 1 0-2h1.07A7.005 7.005 0 0 1 11 5.07V4a1 1 0 0 1 1-1m-5 9a5 5 0 1 1 10 0a5 5 0 0 1-10 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Transcript(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M5 16a1 1 0 0 1 1-1h8a1 1 0 1 1 0 2H6a1 1 0 0 1-1-1m13-5a1 1 0 1 1 0 2h-8a1 1 0 1 1 0-2zm-2 5a1 1 0 0 1 1-1h1a1 1 0 1 1 0 2h-1a1 1 0 0 1-1-1m-9-5a1 1 0 1 1 0 2H6a1 1 0 1 1 0-2z"/><path fill-rule="evenodd" d="M4 3a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h16a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3zm16 2H4a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trash(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M17 5V4a2 2 0 0 0-2-2H9a2 2 0 0 0-2 2v1H4a1 1 0 0 0 0 2h1v11a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V7h1a1 1 0 1 0 0-2zm-2-1H9v1h6zm2 3H7v11a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1z" clip-rule="evenodd"/><path d="M9 9h2v8H9zm4 0h2v8h-2z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrashEmpty(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M17 6V5a2 2 0 0 0-2-2H9a2 2 0 0 0-2 2v1H4a1 1 0 0 0 0 2h1v11a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V8h1a1 1 0 1 0 0-2zm-2-1H9v1h6zm2 3H7v11a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tree(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11 17.9A5.002 5.002 0 0 1 7 13V7a5 5 0 0 1 10 0v6a5.002 5.002 0 0 1-4 4.9V21a1 1 0 1 1-2 0zM12 4a3 3 0 0 1 3 3v6a3.001 3.001 0 0 1-2 2.83V11a1 1 0 1 0-2 0v4.83A3.001 3.001 0 0 1 9 13V7a3 3 0 0 1 3-3" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trees(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.74 16.319A4.995 4.995 0 0 1 10 17.9V21a1 1 0 1 1-2 0v-3.1A5.002 5.002 0 0 1 4 13V7a5 5 0 0 1 9.98-.453A4 4 0 0 1 20 10v4a4.002 4.002 0 0 1-3 3.874V21a1 1 0 1 1-2 0v-3.126a4.005 4.005 0 0 1-2.26-1.555M12 7v6a3.001 3.001 0 0 1-2 2.83V13a1 1 0 1 0-2 0v2.83A3.001 3.001 0 0 1 6 13V7a3 3 0 0 1 6 0m5 8.732V13a1 1 0 1 0-2 0v2.732A2 2 0 0 1 14 14v-4a2 2 0 1 1 4 0v4a2 2 0 0 1-1 1.732" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trello(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M6 7a1 1 0 0 1 1-1h3a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1zm7 0a1 1 0 0 1 1-1h3a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1h-3a1 1 0 0 1-1-1z"/><path fill-rule="evenodd" d="M2 4a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2zm2 0h16v16H4z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trending(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.414 16.432L0 15.018l7.071-7.071l6.364 6.364l4.243-4.243l-1.743-1.742l6.692-1.793l-1.793 6.692l-1.742-1.742l-5.657 5.656l-6.364-6.364z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrendingDown(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.851 8.106L.437 9.52l7.07 7.072l6.365-6.364l4.243 4.242l-1.743 1.743l6.692 1.793l-1.793-6.692l-1.742 1.742l-5.657-5.657l-6.364 6.364z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trophy(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M13 15.9a5.002 5.002 0 0 0 4-4.9V4H7v7a5.002 5.002 0 0 0 4 4.9V18H9v2h6v-2h-2zM9 6h6v5a3 3 0 1 1-6 0z" clip-rule="evenodd"/><path d="M18 6h2v5h-2zM6 6H4v5h2z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tv(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="m8 6.119l1.413-1.413l2.124 2.124L14.367 4l1.411 1.412l-2.464 2.464H18a2 2 0 0 1 2 2v7a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-7a2 2 0 0 1 2-2h3.757zm10 3.757H6v7h12z" clip-rule="evenodd"/><path d="M8 19.876h8v1H8z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Twilio(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M11 9a2 2 0 1 1-4 0a2 2 0 0 1 4 0m6 0a2 2 0 1 1-4 0a2 2 0 0 1 4 0m-2 8a2 2 0 1 0 0-4a2 2 0 0 0 0 4m-4-2a2 2 0 1 1-4 0a2 2 0 0 1 4 0"/><path fill-rule="evenodd" d="M24 12c0 6.627-5.373 12-12 12S0 18.627 0 12S5.373 0 12 0s12 5.373 12 12m-3 0a9 9 0 1 1-18 0a9 9 0 0 1 18 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Twitter(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 3a2 2 0 0 1 2 2v3h6a2 2 0 1 1 0 4h-6v2a3 3 0 0 0 3 3h3a2 2 0 1 1 0 4h-3a7 7 0 0 1-7-7V5a2 2 0 0 1 2-2" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UiKit(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M14 6h-4a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1m-4-2a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h4a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3z" clip-rule="evenodd"/><path d="m6 7.46l-2.507-.418A3 3 0 0 0 0 10v4.917a3 3 0 0 0 3.493 2.96L6 17.458v-2.027l-2.836.472A1 1 0 0 1 2 14.918v-4.917a1 1 0 0 1 1.164-.987L6 9.487zm12 0l2.507-.418A3 3 0 0 1 24 10v4.917a3 3 0 0 1-3.493 2.96L18 17.458v-2.027l2.836.472A1 1 0 0 0 22 14.918v-4.917a1 1 0 0 0-1.164-.987L18 9.487z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Umbrella(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 9a8 8 0 1 1 16 0v2h-6.981v9.5a2.5 2.5 0 0 1-5 0v-2.643h2V20.5a.5.5 0 1 0 1 0V11H4zm8-6a6 6 0 0 1 6 6H6a6 6 0 0 1 6-6" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unavailable(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M18.364 5.636A9 9 0 1 1 5.636 18.364A9 9 0 0 1 18.364 5.636m-2.172 11.97L6.393 7.808a7.001 7.001 0 0 0 9.8 9.8M16.95 7.05a7.002 7.002 0 0 1 .657 9.142l-9.8-9.799a7.001 7.001 0 0 1 9.143.657" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unblock(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.636 18.364A9 9 0 1 0 18.364 5.636A9 9 0 0 0 5.636 18.364m2.171-.757a7.001 7.001 0 0 0 9.8-9.8l-2.779 2.779a1 1 0 0 1-1.414-1.414l2.778-2.779a7.002 7.002 0 0 0-9.799 9.8l2.779-2.779a1 1 0 0 1 1.414 1.414z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Undo(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.34 4.468h2v2.557a7 7 0 1 1-1.037 10.011l1.619-1.185a5 5 0 1 0 .826-7.384h2.591v2h-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unfold(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 12c0 5.523 4.477 10 10 10s10-4.477 10-10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unsplash(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 4.5H9v4h6zm-11 6h5v4h6v-4h5v9H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Usb(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M10 4.5h1v2h-1zm4 0h-1v2h1z"/><path fill-rule="evenodd" d="M7 8.5v-7h10v7h2v11a3 3 0 0 1-3 3H8a3 3 0 0 1-3-3v-11zm2-5h6v5H9zm8 7H7v9a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1z" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UsbC(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M8 11a1 1 0 1 0 0 2h8a1 1 0 1 0 0-2z"/><path fill-rule="evenodd" d="M3 12a5 5 0 0 1 5-5h8a5 5 0 0 1 0 10H8a5 5 0 0 1-5-5m5-3h8a3 3 0 1 1 0 6H8a3 3 0 1 1 0-6" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func User(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M16 7a4 4 0 1 1-8 0a4 4 0 0 1 8 0m-2 0a2 2 0 1 1-4 0a2 2 0 0 1 4 0" clip-rule="evenodd"/><path d="M16 15a1 1 0 0 0-1-1H9a1 1 0 0 0-1 1v6H6v-6a3 3 0 0 1 3-3h6a3 3 0 0 1 3 3v6h-2z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserAdd(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M8 11a4 4 0 1 0 0-8a4 4 0 0 0 0 8m0-2a2 2 0 1 0 0-4a2 2 0 0 0 0 4" clip-rule="evenodd"/><path d="M11 14a1 1 0 0 1 1 1v6h2v-6a3 3 0 0 0-3-3H5a3 3 0 0 0-3 3v6h2v-6a1 1 0 0 1 1-1zm7-7h2v2h2v2h-2v2h-2v-2h-2V9h2z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserList(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M8 11a4 4 0 1 0 0-8a4 4 0 0 0 0 8m0-2a2 2 0 1 0 0-4a2 2 0 0 0 0 4" clip-rule="evenodd"/><path d="M11 14a1 1 0 0 1 1 1v6h2v-6a3 3 0 0 0-3-3H5a3 3 0 0 0-3 3v6h2v-6a1 1 0 0 1 1-1zm11-3h-6v2h6zm-6 4h6v2h-6zm6-8h-6v2h6z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserRemove(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M8 11a4 4 0 1 0 0-8a4 4 0 0 0 0 8m0-2a2 2 0 1 0 0-4a2 2 0 0 0 0 4" clip-rule="evenodd"/><path d="M11 14a1 1 0 0 1 1 1v6h2v-6a3 3 0 0 0-3-3H5a3 3 0 0 0-3 3v6h2v-6a1 1 0 0 1 1-1zm11-5h-6v2h6z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Userlane(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 4h6v6h-6zM3 12a9 9 0 1 0 18 0h-4a5 5 0 0 1-10 0zm3-2a3 3 0 1 0 0-6a3 3 0 0 0 0 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vercel(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.992 17.023L11.981 6.977L6.008 17.023z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewCols(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 8a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3zm14-1h3a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1h-3zm-2 0h-4v10h4zM8 17V7H5a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewComfortable(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 8a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3zm3-1h14a1 1 0 0 1 1 1v3H4V8a1 1 0 0 1 1-1m-1 6v3a1 1 0 0 0 1 1h3v-4zm6 4h9a1 1 0 0 0 1-1v-3H10z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewDay(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 8a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3zm11-1h6a1 1 0 0 1 1 1v3h-7zm-2 0H5a1 1 0 0 0-1 1v3h7zm-7 6v3a1 1 0 0 0 1 1h6v-4zm9 4h6a1 1 0 0 0 1-1v-3h-7z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewGrid(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 5a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V8a3 3 0 0 0-3-3zm3 2H5a1 1 0 0 0-1 1v1h4zm2 0v2h4V7zm6 0v2h4V8a1 1 0 0 0-1-1zm-2 4h-4v2h4zm2 2v-2h4v2zm-2 2h-4v2h4zm2 2v-2h4v1a1 1 0 0 1-1 1zm-8 0v-2H4v1a1 1 0 0 0 1 1zm0-4v-2H4v2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewList(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 5a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V8a3 3 0 0 0-3-3zm2 2H5a1 1 0 0 0-1 1v1h3zm2 0v2h11V8a1 1 0 0 0-1-1zm-2 4H4v2h3zm2 2v-2h11v2zm-2 2H4v1a1 1 0 0 0 1 1h2zm2 2v-2h11v1a1 1 0 0 1-1 1z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewMonth(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 8a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3zm15-1h2a1 1 0 0 1 1 1v1h-3zm-2 0h-2v2h2zm-4 0H9v2h2zM7 7H5a1 1 0 0 0-1 1v1h3zm-3 4v2h3v-2zm0 4v1a1 1 0 0 0 1 1h2v-2zm5 2h2v-2H9zm4 0h2v-2h-2zm4 0h2a1 1 0 0 0 1-1v-1h-3zm3-4v-2h-3v2zm-9 0H9v-2h2zm4 0h-2v-2h2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewSplit(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 8a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3zm11-1h6a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1h-6zm-2 0H5a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h6z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vinyl(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M12 13a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path fill-rule="evenodd" d="M20 12a8 8 0 1 1-16 0a8 8 0 0 1 16 0m-4 0a4 4 0 1 1-8 0a4 4 0 0 1 8 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Voicemail(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 15a5 5 0 1 0-4 2h12a5 5 0 1 0-4-2zm-4 0a3 3 0 1 0 0-6a3 3 0 0 0 0 6m12 0a3 3 0 1 0 0-6a3 3 0 0 0 0 6" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VoicemailO(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M11 12c0 .35-.06.687-.17 1h2.34A3 3 0 1 1 16 15H8a3 3 0 1 1 3-3m-2 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0m8 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/><path d="M23 12c0 6.075-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1s11 4.925 11 11m-2 0a9 9 0 1 1-18 0a9 9 0 0 1 18 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VoicemailR(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M11.5 12c0 .35-.06.687-.17 1h2.34a3 3 0 1 1 2.83 2h-8a3 3 0 1 1 3-3m-2 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0m8 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/><path d="M1.5 8a3 3 0 0 1 3-3h15a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3h-15a3 3 0 0 1-3-3zm3-1h15a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1h-15a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Volume(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M24 12a8 8 0 0 1-8 8v-2a6 6 0 0 0 0-12V4a8 8 0 0 1 8 8"/><path d="M20 12a4 4 0 0 1-4 4v-2a2 2 0 1 0 0-4V8a4 4 0 0 1 4 4"/><path fill-rule="evenodd" d="m9 16l6 4V4L9 8H5a4 4 0 1 0 0 8zm-4-6h4l4-2.5v9L9 14H5a2 2 0 1 1 0-4" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Webcam(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M13 10a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/><path fill-rule="evenodd" d="M13 14.9A5.002 5.002 0 0 0 12 5a5 5 0 0 0-1 9.9V17H7v2h10v-2h-4zM12 13a3 3 0 1 0 0-6a3 3 0 0 0 0 6" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Website(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path fill-rule="evenodd" d="M14 7a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1V8a1 1 0 0 0-1-1zm3 2h-2v6h2z" clip-rule="evenodd"/><path d="M6 7a1 1 0 0 0 0 2h4a1 1 0 1 0 0-2zm0 4a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2zm-1 5a1 1 0 0 1 1-1h4a1 1 0 1 1 0 2H6a1 1 0 0 1-1-1"/><path fill-rule="evenodd" d="M4 3a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h16a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3zm16 2H4a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Windows(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m3 5.549l7.195-.967v7.029l-7.188.054zm7.195 6.842v7.105l-7.19-.985v-6.12zm.918-7.935L20.998 3v8.533l-9.885.078zM21 12.505L20.998 21l-9.885-1.353v-7.142z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WorkAlt(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M17 7a3 3 0 0 0-3-3h-4a3 3 0 0 0-3 3H6a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-8a3 3 0 0 0-3-3zm-3-1h-4a1 1 0 0 0-1 1h6a1 1 0 0 0-1-1M6 9h12a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1v-8a1 1 0 0 1 1-1" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Yinyang(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor"><path d="M14 16a2 2 0 1 1-4 0a2 2 0 0 1 4 0"/><path fill-rule="evenodd" d="M22 12c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2s10 4.477 10 10m-10 0a4 4 0 0 1 0-8a8 8 0 1 0 0 16a4 4 0 0 0 0-8m2-4a2 2 0 1 1-4 0a2 2 0 0 1 4 0" clip-rule="evenodd"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Youtube(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 7h14a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1M2 8a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3zm8 1l4 3l-4 3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomIn(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.343 15.243a6 6 0 1 0-8.485-8.486a6 6 0 0 0 8.485 8.486m1.414-9.9a8.001 8.001 0 0 1 .662 10.565c.016.013.03.027.046.042l4.242 4.242a1 1 0 0 1-1.414 1.415l-4.243-4.243a.99.99 0 0 1-.042-.045A8.001 8.001 0 0 1 5.444 5.343a8 8 0 0 1 11.313 0M10.1 7h2v3h3v2h-3v3h-2v-3h-3v-2h3z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomOut(children ...ElementRenderer) *GgIcon {
	return &GgIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.343 15.243a6 6 0 1 0-8.485-8.486a6 6 0 0 0 8.485 8.486m1.414-9.9a8.001 8.001 0 0 1 .662 10.565c.016.013.03.027.046.042l4.242 4.242a1 1 0 0 1-1.414 1.415l-4.243-4.243a.99.99 0 0 1-.042-.045A8.001 8.001 0 0 1 5.444 5.343a8 8 0 0 1 11.313 0M7.101 10v2h8v-2z" clip-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
