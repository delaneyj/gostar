package cil

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = "2.0.1"
	hAttr          = "1em"
	viewbox        = "0 0 512 512"
	fill           = "currentColor"
)

type CilIcon struct {
	*SVGSVGElement
}

type CilIconFn func(children ...ElementRenderer) *CilIcon

var IconLookup = map[string]CilIconFn{
	"accountLogout":            AccountLogout,
	"actionRedo":               ActionRedo,
	"actionUndo":               ActionUndo,
	"addressBook":              AddressBook,
	"airplaneMode":             AirplaneMode,
	"airplaneModeOff":          AirplaneModeOff,
	"airplay":                  Airplay,
	"alarm":                    Alarm,
	"album":                    Album,
	"alignCenter":              AlignCenter,
	"alignLeft":                AlignLeft,
	"alignRight":               AlignRight,
	"americanFootball":         AmericanFootball,
	"animal":                   Animal,
	"aperture":                 Aperture,
	"apple":                    Apple,
	"applications":             Applications,
	"applicationsSettings":     ApplicationsSettings,
	"apps":                     Apps,
	"appsSettings":             AppsSettings,
	"arrowBottom":              ArrowBottom,
	"arrowCircleBottom":        ArrowCircleBottom,
	"arrowCircleLeft":          ArrowCircleLeft,
	"arrowCircleRight":         ArrowCircleRight,
	"arrowCircleTop":           ArrowCircleTop,
	"arrowLeft":                ArrowLeft,
	"arrowRight":               ArrowRight,
	"arrowThickBottom":         ArrowThickBottom,
	"arrowThickFromBottom":     ArrowThickFromBottom,
	"arrowThickFromLeft":       ArrowThickFromLeft,
	"arrowThickFromRight":      ArrowThickFromRight,
	"arrowThickFromTop":        ArrowThickFromTop,
	"arrowThickLeft":           ArrowThickLeft,
	"arrowThickRight":          ArrowThickRight,
	"arrowThickToBottom":       ArrowThickToBottom,
	"arrowThickToLeft":         ArrowThickToLeft,
	"arrowThickToRight":        ArrowThickToRight,
	"arrowThickToTop":          ArrowThickToTop,
	"arrowThickTop":            ArrowThickTop,
	"arrowTop":                 ArrowTop,
	"assistiveListeningSystem": AssistiveListeningSystem,
	"asterisk":                 Asterisk,
	"asteriskCircle":           AsteriskCircle,
	"at":                       At,
	"audio":                    Audio,
	"audioDescription":         AudioDescription,
	"audioSpectrum":            AudioSpectrum,
	"avTimer":                  AvTimer,
	"baby":                     Baby,
	"babyCarriage":             BabyCarriage,
	"backspace":                Backspace,
	"badge":                    Badge,
	"balanceScale":             BalanceScale,
	"ban":                      Ban,
	"bank":                     Bank,
	"barChart":                 BarChart,
	"barcode":                  Barcode,
	"baseball":                 Baseball,
	"basket":                   Basket,
	"basketball":               Basketball,
	"bath":                     Bath,
	"bathroom":                 Bathroom,
	"batteryAlert":             BatteryAlert,
	"batteryEmpty":             BatteryEmpty,
	"batteryFive":              BatteryFive,
	"batteryFull":              BatteryFull,
	"batterySlash":             BatterySlash,
	"batteryThree":             BatteryThree,
	"batteryZero":              BatteryZero,
	"beachAccess":              BeachAccess,
	"beaker":                   Beaker,
	"bed":                      Bed,
	"bell":                     Bell,
	"bellExclamation":          BellExclamation,
	"bike":                     Bike,
	"birthdayCake":             BirthdayCake,
	"blind":                    Blind,
	"bluetooth":                Bluetooth,
	"blur":                     Blur,
	"blurCircular":             BlurCircular,
	"blurLinear":               BlurLinear,
	"boatAlt":                  BoatAlt,
	"bold":                     Bold,
	"bolt":                     Bolt,
	"boltCircle":               BoltCircle,
	"book":                     Book,
	"bookmark":                 Bookmark,
	"borderAll":                BorderAll,
	"borderBottom":             BorderBottom,
	"borderClear":              BorderClear,
	"borderHorizontal":         BorderHorizontal,
	"borderInner":              BorderInner,
	"borderLeft":               BorderLeft,
	"borderOuter":              BorderOuter,
	"borderRight":              BorderRight,
	"borderStyle":              BorderStyle,
	"borderTop":                BorderTop,
	"borderVertical":           BorderVertical,
	"bowling":                  Bowling,
	"braille":                  Braille,
	"briefcase":                Briefcase,
	"brightness":               Brightness,
	"britishPound":             BritishPound,
	"browser":                  Browser,
	"brush":                    Brush,
	"brushAlt":                 BrushAlt,
	"bug":                      Bug,
	"building":                 Building,
	"bullhorn":                 Bullhorn,
	"burger":                   Burger,
	"burn":                     Burn,
	"busAlt":                   BusAlt,
	"calculator":               Calculator,
	"calendar":                 Calendar,
	"calendarCheck":            CalendarCheck,
	"camera":                   Camera,
	"cameraControl":            CameraControl,
	"cameraRoll":               CameraRoll,
	"carAlt":                   CarAlt,
	"caretBottom":              CaretBottom,
	"caretLeft":                CaretLeft,
	"caretRight":               CaretRight,
	"caretTop":                 CaretTop,
	"cart":                     Cart,
	"cash":                     Cash,
	"casino":                   Casino,
	"cast":                     Cast,
	"cat":                      Cat,
	"cc":                       Cc,
	"centerFocus":              CenterFocus,
	"chart":                    Chart,
	"chartLine":                ChartLine,
	"chartPie":                 ChartPie,
	"chatBubble":               ChatBubble,
	"check":                    Check,
	"checkAlt":                 CheckAlt,
	"checkCircle":              CheckCircle,
	"chevronBottom":            ChevronBottom,
	"chevronCircleDownAlt":     ChevronCircleDownAlt,
	"chevronCircleLeftAlt":     ChevronCircleLeftAlt,
	"chevronCircleRightAlt":    ChevronCircleRightAlt,
	"chevronCircleUpAlt":       ChevronCircleUpAlt,
	"chevronDoubleDown":        ChevronDoubleDown,
	"chevronDoubleLeft":        ChevronDoubleLeft,
	"chevronDoubleRight":       ChevronDoubleRight,
	"chevronDoubleUp":          ChevronDoubleUp,
	"chevronDoubleUpAlt":       ChevronDoubleUpAlt,
	"chevronLeft":              ChevronLeft,
	"chevronRight":             ChevronRight,
	"chevronTop":               ChevronTop,
	"child":                    Child,
	"childFriendly":            ChildFriendly,
	"circle":                   Circle,
	"clearAll":                 ClearAll,
	"clipboard":                Clipboard,
	"clock":                    Clock,
	"clone":                    Clone,
	"closedCaptioning":         ClosedCaptioning,
	"cloud":                    Cloud,
	"cloudDownload":            CloudDownload,
	"cloudUpload":              CloudUpload,
	"cloudy":                   Cloudy,
	"code":                     Code,
	"coffee":                   Coffee,
	"cog":                      Cog,
	"colorBorder":              ColorBorder,
	"colorFill":                ColorFill,
	"colorPalette":             ColorPalette,
	"columns":                  Columns,
	"command":                  Command,
	"commentBubble":            CommentBubble,
	"commentSquare":            CommentSquare,
	"compass":                  Compass,
	"compress":                 Compress,
	"contact":                  Contact,
	"contrast":                 Contrast,
	"control":                  Control,
	"copy":                     Copy,
	"couch":                    Couch,
	"creditCard":               CreditCard,
	"crop":                     Crop,
	"cropRotate":               CropRotate,
	"cursor":                   Cursor,
	"cursorMove":               CursorMove,
	"cut":                      Cut,
	"dataTransferDown":         DataTransferDown,
	"dataTransferUp":           DataTransferUp,
	"deaf":                     Deaf,
	"delete":                   Delete,
	"description":              Description,
	"devices":                  Devices,
	"dialpad":                  Dialpad,
	"diamond":                  Diamond,
	"dinner":                   Dinner,
	"disabled":                 Disabled,
	"dog":                      Dog,
	"dollar":                   Dollar,
	"door":                     Door,
	"doubleQuoteSansLeft":      DoubleQuoteSansLeft,
	"doubleQuoteSansRight":     DoubleQuoteSansRight,
	"drink":                    Drink,
	"drinkAlcohol":             DrinkAlcohol,
	"drop":                     Drop,
	"dropOne":                  DropOne,
	"eco":                      Eco,
	"education":                Education,
	"elevator":                 Elevator,
	"envelopeClosed":           EnvelopeClosed,
	"envelopeLetter":           EnvelopeLetter,
	"envelopeOpen":             EnvelopeOpen,
	"equalizer":                Equalizer,
	"ethernet":                 Ethernet,
	"euro":                     Euro,
	"excerpt":                  Excerpt,
	"exitToApp":                ExitToApp,
	"expandDown":               ExpandDown,
	"expandLeft":               ExpandLeft,
	"expandRight":              ExpandRight,
	"expandUp":                 ExpandUp,
	"exposure":                 Exposure,
	"externalLink":             ExternalLink,
	"eyedropper":               Eyedropper,
	"face":                     Face,
	"faceDead":                 FaceDead,
	"factory":                  Factory,
	"factorySlash":             FactorySlash,
	"fastfood":                 Fastfood,
	"fax":                      Fax,
	"featuredPlaylist":         FeaturedPlaylist,
	"file":                     File,
	"filter":                   Filter,
	"filterFrames":             FilterFrames,
	"filterPhoto":              FilterPhoto,
	"filterSquare":             FilterSquare,
	"filterX":                  FilterX,
	"findInPage":               FindInPage,
	"fingerprint":              Fingerprint,
	"fire":                     Fire,
	"flagAlt":                  FlagAlt,
	"flightTakeoff":            FlightTakeoff,
	"flip":                     Flip,
	"flipToBack":               FlipToBack,
	"flipToFront":              FlipToFront,
	"flower":                   Flower,
	"folder":                   Folder,
	"folderOpen":               FolderOpen,
	"font":                     Font,
	"football":                 Football,
	"fork":                     Fork,
	"fourK":                    FourK,
	"fridge":                   Fridge,
	"frown":                    Frown,
	"fullscreen":               Fullscreen,
	"fullscreenExit":           FullscreenExit,
	"functions":                Functions,
	"functionsAlt":             FunctionsAlt,
	"gamepad":                  Gamepad,
	"garage":                   Garage,
	"gauge":                    Gauge,
	"gem":                      Gem,
	"gif":                      Gif,
	"gift":                     Gift,
	"globeAlt":                 GlobeAlt,
	"golf":                     Golf,
	"golfAlt":                  GolfAlt,
	"gradient":                 Gradient,
	"grain":                    Grain,
	"graph":                    Graph,
	"grid":                     Grid,
	"gridSlash":                GridSlash,
	"groupIcon":                GroupIcon,
	"hamburgerMenu":            HamburgerMenu,
	"handPointDown":            HandPointDown,
	"handPointLeft":            HandPointLeft,
	"handPointRight":           HandPointRight,
	"handPointUp":              HandPointUp,
	"happy":                    Happy,
	"hd":                       Hd,
	"hdr":                      Hdr,
	"header":                   Header,
	"headphones":               Headphones,
	"healing":                  Healing,
	"heart":                    Heart,
	"highlighter":              Highlighter,
	"highligt":                 Highligt,
	"history":                  History,
	"home":                     Home,
	"hospital":                 Hospital,
	"hotTub":                   HotTub,
	"house":                    House,
	"https":                    Https,
	"image":                    Image,
	"imageBroken":              ImageBroken,
	"imageOne":                 ImageOne,
	"imagePlus":                ImagePlus,
	"inbox":                    Inbox,
	"indentDecrease":           IndentDecrease,
	"indentIncrease":           IndentIncrease,
	"industry":                 Industry,
	"industrySlash":            IndustrySlash,
	"infinity":                 Infinity,
	"info":                     Info,
	"input":                    Input,
	"inputHdmi":                InputHdmi,
	"inputPower":               InputPower,
	"institution":              Institution,
	"italic":                   Italic,
	"justifyCenter":            JustifyCenter,
	"justifyLeft":              JustifyLeft,
	"justifyRight":             JustifyRight,
	"keyboard":                 Keyboard,
	"lan":                      Lan,
	"language":                 Language,
	"laptop":                   Laptop,
	"layers":                   Layers,
	"leaf":                     Leaf,
	"lemon":                    Lemon,
	"levelDown":                LevelDown,
	"levelUp":                  LevelUp,
	"library":                  Library,
	"libraryAdd":               LibraryAdd,
	"libraryBuilding":          LibraryBuilding,
	"lifeRing":                 LifeRing,
	"lightbulb":                Lightbulb,
	"lineSpacing":              LineSpacing,
	"lineStyle":                LineStyle,
	"lineWeight":               LineWeight,
	"link":                     Link,
	"linkAlt":                  LinkAlt,
	"linkBroken":               LinkBroken,
	"list":                     List,
	"listFilter":               ListFilter,
	"listHighPriority":         ListHighPriority,
	"listLowPriority":          ListLowPriority,
	"listNumbered":             ListNumbered,
	"listNumberedRtl":          ListNumberedRtl,
	"listRich":                 ListRich,
	"locationPin":              LocationPin,
	"lockLocked":               LockLocked,
	"lockUnlocked":             LockUnlocked,
	"locomotive":               Locomotive,
	"loop":                     Loop,
	"loopCircular":             LoopCircular,
	"loopOne":                  LoopOne,
	"lowVision":                LowVision,
	"magnifyingGlass":          MagnifyingGlass,
	"map":                      Map,
	"mediaEject":               MediaEject,
	"mediaPause":               MediaPause,
	"mediaPlay":                MediaPlay,
	"mediaRecord":              MediaRecord,
	"mediaSkipBackward":        MediaSkipBackward,
	"mediaSkipForward":         MediaSkipForward,
	"mediaStepBackward":        MediaStepBackward,
	"mediaStepForward":         MediaStepForward,
	"mediaStop":                MediaStop,
	"medicalCross":             MedicalCross,
	"meh":                      Meh,
	"memory":                   Memory,
	"menu":                     Menu,
	"mic":                      Mic,
	"microphone":               Microphone,
	"minus":                    Minus,
	"mobile":                   Mobile,
	"mobileLandscape":          MobileLandscape,
	"money":                    Money,
	"monitor":                  Monitor,
	"moodBad":                  MoodBad,
	"moodGood":                 MoodGood,
	"moodVeryBad":              MoodVeryBad,
	"moodVeryGood":             MoodVeryGood,
	"moon":                     Moon,
	"mouse":                    Mouse,
	"mouthSlash":               MouthSlash,
	"move":                     Move,
	"movie":                    Movie,
	"mug":                      Mug,
	"mugTea":                   MugTea,
	"musicNote":                MusicNote,
	"newspaper":                Newspaper,
	"noteAdd":                  NoteAdd,
	"notes":                    Notes,
	"objectGroup":              ObjectGroup,
	"objectUngroup":            ObjectUngroup,
	"opacity":                  Opacity,
	"opentype":                 Opentype,
	"options":                  Options,
	"optionsHorizontal":        OptionsHorizontal,
	"paint":                    Paint,
	"paintBucket":              PaintBucket,
	"paperPlane":               PaperPlane,
	"paperclip":                Paperclip,
	"paragraph":                Paragraph,
	"paw":                      Paw,
	"pen":                      Pen,
	"penAlt":                   PenAlt,
	"penNib":                   PenNib,
	"pencil":                   Pencil,
	"people":                   People,
	"phone":                    Phone,
	"pin":                      Pin,
	"pizza":                    Pizza,
	"plant":                    Plant,
	"playlistAdd":              PlaylistAdd,
	"plus":                     Plus,
	"pool":                     Pool,
	"powerStandby":             PowerStandby,
	"pregnant":                 Pregnant,
	"print":                    Print,
	"pushchair":                Pushchair,
	"puzzle":                   Puzzle,
	"qrCode":                   QrCode,
	"rain":                     Rain,
	"rectangle":                Rectangle,
	"recycle":                  Recycle,
	"reload":                   Reload,
	"remove":                   Remove,
	"reportSlash":              ReportSlash,
	"resizeBoth":               ResizeBoth,
	"resizeHeight":             ResizeHeight,
	"resizeWidth":              ResizeWidth,
	"restaurant":               Restaurant,
	"room":                     Room,
	"router":                   Router,
	"rowing":                   Rowing,
	"rss":                      Rss,
	"ruble":                    Ruble,
	"running":                  Running,
	"sad":                      Sad,
	"satelite":                 Satelite,
	"save":                     Save,
	"school":                   School,
	"screenDesktop":            ScreenDesktop,
	"screenSmartphone":         ScreenSmartphone,
	"scrubber":                 Scrubber,
	"search":                   Search,
	"send":                     Send,
	"settings":                 Settings,
	"share":                    Share,
	"shareAll":                 ShareAll,
	"shareAlt":                 ShareAlt,
	"shareBoxed":               ShareBoxed,
	"shieldAlt":                ShieldAlt,
	"shortText":                ShortText,
	"shower":                   Shower,
	"signLanguage":             SignLanguage,
	"signalCellularFour":       SignalCellularFour,
	"signalCellularThree":      SignalCellularThree,
	"signalCellularZero":       SignalCellularZero,
	"sim":                      Sim,
	"sitemap":                  Sitemap,
	"smile":                    Smile,
	"smilePlus":                SmilePlus,
	"smoke":                    Smoke,
	"smokeFree":                SmokeFree,
	"smokeSlash":               SmokeSlash,
	"smokingRoom":              SmokingRoom,
	"snowflake":                Snowflake,
	"soccer":                   Soccer,
	"sofa":                     Sofa,
	"sortAlphaDown":            SortAlphaDown,
	"sortAlphaUp":              SortAlphaUp,
	"sortAscending":            SortAscending,
	"sortDescending":           SortDescending,
	"sortNumericDown":          SortNumericDown,
	"sortNumericUp":            SortNumericUp,
	"spa":                      Spa,
	"spaceBar":                 SpaceBar,
	"speak":                    Speak,
	"speaker":                  Speaker,
	"speech":                   Speech,
	"speedometer":              Speedometer,
	"spreadsheet":              Spreadsheet,
	"square":                   Square,
	"star":                     Star,
	"starHalf":                 StarHalf,
	"storage":                  Storage,
	"stream":                   Stream,
	"strikethrough":            Strikethrough,
	"sun":                      Sun,
	"swapHorizontal":           SwapHorizontal,
	"swapVertical":             SwapVertical,
	"swimming":                 Swimming,
	"sync":                     Sync,
	"tablet":                   Tablet,
	"tag":                      Tag,
	"tags":                     Tags,
	"task":                     Task,
	"taxi":                     Taxi,
	"tennis":                   Tennis,
	"tennisBall":               TennisBall,
	"terminal":                 Terminal,
	"terrain":                  Terrain,
	"textIcon":                 TextIcon,
	"textShapes":               TextShapes,
	"textSize":                 TextSize,
	"textSquare":               TextSquare,
	"textStrike":               TextStrike,
	"threeD":                   ThreeD,
	"thumbDown":                ThumbDown,
	"thumbUp":                  ThumbUp,
	"toggleOff":                ToggleOff,
	"toggleOn":                 ToggleOn,
	"toilet":                   Toilet,
	"touchApp":                 TouchApp,
	"transfer":                 Transfer,
	"translate":                Translate,
	"trash":                    Trash,
	"triangle":                 Triangle,
	"truck":                    Truck,
	"tv":                       Tv,
	"underline":                Underline,
	"usb":                      Usb,
	"user":                     User,
	"userFemale":               UserFemale,
	"userFollow":               UserFollow,
	"userPlus":                 UserPlus,
	"userUnfollow":             UserUnfollow,
	"userX":                    UserX,
	"vector":                   Vector,
	"verticalAlignBottom":      VerticalAlignBottom,
	"verticalAlignBottomOne":   VerticalAlignBottomOne,
	"verticalAlignCenter":      VerticalAlignCenter,
	"verticalAlignCenterOne":   VerticalAlignCenterOne,
	"verticalAlignTop":         VerticalAlignTop,
	"verticalAlignTopOne":      VerticalAlignTopOne,
	"video":                    Video,
	"videogame":                Videogame,
	"viewColumn":               ViewColumn,
	"viewModule":               ViewModule,
	"viewQuilt":                ViewQuilt,
	"viewStream":               ViewStream,
	"voice":                    Voice,
	"voiceOverRecord":          VoiceOverRecord,
	"volumeHigh":               VolumeHigh,
	"volumeLow":                VolumeLow,
	"volumeOff":                VolumeOff,
	"walk":                     Walk,
	"wallet":                   Wallet,
	"wallpaper":                Wallpaper,
	"warning":                  Warning,
	"watch":                    Watch,
	"wc":                       Wc,
	"weightlifitng":            Weightlifitng,
	"wheelchair":               Wheelchair,
	"wifiSignalFour":           WifiSignalFour,
	"wifiSignalOff":            WifiSignalOff,
	"wifiSignalOne":            WifiSignalOne,
	"wifiSignalThree":          WifiSignalThree,
	"wifiSignalTwo":            WifiSignalTwo,
	"wifiSignalZero":           WifiSignalZero,
	"window":                   Window,
	"windowMaximize":           WindowMaximize,
	"windowMinimize":           WindowMinimize,
	"windowRestore":            WindowRestore,
	"wrapText":                 WrapText,
	"x":                        X,
	"xcircle":                  Xcircle,
	"yen":                      Yen,
	"zoom":                     Zoom,
	"zoomIn":                   ZoomIn,
	"zoomOut":                  ZoomOut,
}

func AccountLogout(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M77.155 272.034H351.75v-32.001H77.155l75.053-75.053v-.001l-22.628-22.626l-113.681 113.68l.001.001h-.001L129.58 369.715l22.628-22.627v-.001z"/><path fill="currentColor" d="M160 16v32h304v416H160v32h336V16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ActionRedo(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M361.376 495.163L226.753 360.54l22.627-22.627l111.996 111.996l111.997-111.996L496 360.54z"/><path fill="currentColor" d="M377.377 472.52h-32V196.426C345.377 114.584 278.794 48 196.952 48c-83.229 0-148.426 63.106-148.426 143.667h-32c0-48.024 18.85-92.569 53.079-125.429C103.35 33.842 148.576 16 196.952 16c99.487 0 180.425 80.938 180.425 180.426Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ActionUndo(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M495.473 197.262c0-73.061-43.651-136.118-106.242-164.462q-2-.9-4.021-1.762q-4.046-1.715-8.19-3.235q-2.071-.761-4.167-1.47q-4.19-1.422-8.468-2.64a180.951 180.951 0 0 0-98.675 0q-4.278 1.218-8.469 2.64q-2.094.71-4.166 1.47q-4.143 1.519-8.19 3.235q-2.023.857-4.021 1.762c-62.592 28.344-106.242 91.4-106.242 164.462v237.483L38.627 338.75L16 361.377L150.623 496l134.623-134.623l-22.627-22.627l-96 96V197.263c0-72.891 52.814-133.678 122.186-146.1a149.419 149.419 0 0 1 52.479 0c69.371 12.426 122.186 73.213 122.186 146.1h32Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AddressBook(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M496 144.768v-33.536h-39.232V42a25 25 0 0 0-25.179-24.768H80.411A25 25 0 0 0 55.232 42v430a25 25 0 0 0 25.179 24.768h351.178A25 25 0 0 0 456.768 472v-71.232H496v-33.536h-39.232v-94.464H496v-33.536h-39.232v-94.464Zm-72.768 94.464H376v33.536h47.232v94.464H376v33.536h47.232v62.464H88.768V50.768h334.464v60.464H376v33.536h47.232Z"/><path fill="currentColor" d="m313.639 306.925l-28.745-18.685l13.82-33.655v-52.871a65.714 65.714 0 1 0-131.428 0v52.557l12.721 34.684l-27.646 17.97A48.972 48.972 0 0 0 130 348.129V400h206v-51.871a48.972 48.972 0 0 0-22.361-41.204M304 368H162v-19.871a17.084 17.084 0 0 1 7.8-14.373l49.033-31.872l-19.547-53.3v-46.87a33.714 33.714 0 0 1 67.428 0v46.557l-21.5 52.347l50.986 33.138a17.084 17.084 0 0 1 7.8 14.373Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AirplaneMode(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m476.757 337.8l13.329-66.646L320 186.111v-72.222a105.15 105.15 0 0 0-37.937-81a40.705 40.705 0 0 0-52.126 0a105.15 105.15 0 0 0-37.937 81v72.222L21.914 271.154L35.243 337.8l157.189-20.5l7.736 81.224L128 429.45V496h256v-66.55l-72.168-30.929l7.736-81.224ZM352 450.551V464H160v-13.449l74.238-31.818L224 311.24v-30.332L60.757 302.2l-2.671-13.354L224 205.889v-92a73.235 73.235 0 0 1 26.423-56.413a8.707 8.707 0 0 1 11.154 0A73.235 73.235 0 0 1 288 113.889v92l165.914 82.957l-2.671 13.354L288 280.908v30.332l-10.238 107.493Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AirplaneModeOff(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M224 113.889a73.235 73.235 0 0 1 26.423-56.413a8.707 8.707 0 0 1 11.154 0A73.235 73.235 0 0 1 288 113.889v92l165.914 82.957l-2.671 13.354l-86.936-11.339l37.114 37.112l75.336 9.827l13.329-66.646L320 186.111v-72.222a105.155 105.155 0 0 0-37.937-81a40.705 40.705 0 0 0-52.126 0a105.155 105.155 0 0 0-37.937 81v4.675l32 32ZM38.517 16H16.029v22.655L173 195.613L21.914 271.154L35.243 337.8l157.189-20.5l7.735 81.224L128 429.45V496h256v-66.55l-72.167-30.929l5.571-58.507L473.4 496H496v-22.542ZM352 450.55V464H160v-13.45l74.238-31.817L224 311.24v-30.332L60.757 302.2l-2.671-13.354l138.762-69.381L288 310.611v.629l-10.238 107.493Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Airplay(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 47H48a32.036 32.036 0 0 0-32 32v297a32.036 32.036 0 0 0 32 32h76.448l24.89-32H48V79h416l.02 297H362.662l24.89 32H464a32.036 32.036 0 0 0 32-32V79a32.036 32.036 0 0 0-32-32"/><path fill="currentColor" d="M98.834 496h314.332L256 293.939Zm65.431-32L256 346.061L347.735 464Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Alarm(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 80C141.125 80 48 173.125 48 288s93.125 208 208 208s208-93.125 208-208S370.875 80 256 80m124.451 332.451A176 176 0 1 1 432 288a174.849 174.849 0 0 1-51.549 124.451"/><path fill="currentColor" d="M272 160h-32v135.69l86.005 68.804l19.99-24.988L272 280.31zM22.965 114.796l152-104l18.071 26.411l-152 104zm296.002-77.59l18.07-26.41l152 104.002l-18.071 26.41z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Album(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M425.706 86.294A240 240 0 0 0 86.294 425.706A240 240 0 0 0 425.706 86.294M256 464c-114.691 0-208-93.309-208-208S141.309 48 256 48s208 93.309 208 208s-93.309 208-208 208"/><path fill="currentColor" d="M256 152a104 104 0 1 0 104 104a104.118 104.118 0 0 0-104-104m0 176a72 72 0 1 1 72-72a72.081 72.081 0 0 1-72 72"/><path fill="currentColor" d="M240 240h32v32h-32zm16-128V80a174.144 174.144 0 0 0-79.968 19.178A177.573 177.573 0 0 0 115.2 150.39l25.586 19.219A142.923 142.923 0 0 1 256 112"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignCenter(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 64h480v32H16zm80 88h320v32H96zm-80 88h480v32H16zm80 88h320v32H96zm-80 88h480v32H16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignLeft(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 64h480v32H16zm0 88h328v32H16zm0 88h480v32H16zm0 88h328v32H16zm0 88h480v32H16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignRight(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 64h480v32H16zm152 88h328v32H168zM16 240h480v32H16zm152 88h328v32H168zM16 416h480v32H16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AmericanFootball(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m473.274 47.136l-2.267-4.189l-4.188-2.267C416.85 13.637 354.765 6.167 292 19.646c-63.848 13.713-123.787 47.151-173.335 96.7s-82.987 109.487-96.7 173.335C8.489 352.444 15.958 414.529 43 464.5l2.267 4.189l4.189 2.267c33.722 18.25 72.951 27.585 114.194 27.585a288.981 288.981 0 0 0 60.622-6.552c63.848-13.712 123.786-47.151 173.334-96.7S480.6 285.8 494.308 221.952c13.48-62.764 6.01-124.852-21.034-174.816M447.354 66.6c19.78 38.858 25.84 83.807 19.4 129.525L317.829 47.2c45.719-6.44 90.671-.38 129.525 19.4M68.922 445.033C48.72 405.346 42.83 359.306 49.953 312.58L201.374 464c-46.725 7.125-92.766 1.234-132.452-18.967m306.059-72.374c-40.434 40.434-88.112 68.589-136.661 83.034L58.261 275.635c14.446-48.55 42.6-96.228 83.034-136.662c41.156-41.156 89.818-69.583 139.264-83.787L458.769 233.4c-14.204 49.442-42.631 98.1-83.788 139.259"/><path fill="currentColor" d="m347.313 187.313l-22.626-22.626L304 185.373l-28.687-28.686l-22.626 22.626L281.373 208L256 233.373l-28.687-28.686l-22.626 22.626L233.373 256L208 281.373l-28.687-28.686l-22.626 22.626L185.373 304l-20.686 20.687l22.626 22.626L208 326.627l28.687 28.686l22.626-22.626L230.627 304L256 278.627l28.687 28.686l22.626-22.626L278.627 256L304 230.627l28.687 28.686l22.626-22.626L326.627 208z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Animal(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M382.825 304.576a131.562 131.562 0 0 0-253.65 0l-18.248 66.15A80 80 0 0 0 188.046 472h135.908a80 80 0 0 0 77.119-101.274Zm-20.682 116.5A47.638 47.638 0 0 1 323.954 440H188.046a48 48 0 0 1-46.272-60.765l18.248-66.149a99.563 99.563 0 0 1 191.956 0l18.248 66.149a47.636 47.636 0 0 1-8.083 41.845ZM146.1 230.31c2.784-17.4-.908-36.027-10.4-52.463s-23.78-28.947-40.237-35.236c-17.624-6.731-35.6-5.659-50.634 3.017c-29.887 17.256-37.752 59.785-17.529 94.805c9.489 16.436 23.778 28.95 40.235 35.236a64.058 64.058 0 0 0 22.863 4.371a55.133 55.133 0 0 0 27.771-7.389c15.025-8.677 24.945-23.714 27.931-42.341m-31.6-5.058c-1.43 8.929-5.81 15.92-12.333 19.686S87.4 249 78.95 245.775c-9.613-3.671-18.115-11.251-23.941-21.342c-11.2-19.4-8.538-42.8 5.82-51.092a23.483 23.483 0 0 1 11.847-3.058a31.951 31.951 0 0 1 11.368 2.217c9.613 3.673 18.115 11.252 23.941 21.343s8.139 21.248 6.515 31.409m35.066-61.235c11.362 9.083 24.337 13.813 37.458 13.812a54.965 54.965 0 0 0 11.689-1.261c33.723-7.331 54.17-45.443 45.58-84.958c-4.03-18.546-13.828-34.817-27.588-45.818c-14.735-11.78-32.189-16.239-49.147-12.551c-33.722 7.33-54.169 45.442-45.58 84.957c4.031 18.547 13.829 34.818 27.588 45.819m24.788-99.506a22.258 22.258 0 0 1 4.732-.5c5.948 0 12.066 2.327 17.637 6.781c8.037 6.425 13.826 16.234 16.3 27.621c4.76 21.895-4.906 43.368-21.107 46.89c-7.361 1.6-15.305-.628-22.367-6.275c-8.037-6.426-13.826-16.235-16.3-27.621c-4.761-21.901 4.905-43.374 21.105-46.896m292.817 81.117c-15.028-8.676-33.013-9.748-50.634-3.017c-16.457 6.287-30.746 18.8-40.235 35.236s-13.182 35.067-10.4 52.463c2.982 18.627 12.9 33.664 27.931 42.341a55.123 55.123 0 0 0 27.771 7.389a64.054 64.054 0 0 0 22.863-4.371c16.457-6.286 30.746-18.8 40.235-35.236c20.221-35.02 12.356-77.549-17.531-94.805m-10.18 78.805c-5.826 10.091-14.328 17.671-23.941 21.342c-8.446 3.228-16.692 2.931-23.215-.837s-10.9-10.757-12.333-19.686c-1.626-10.161.686-21.314 6.513-31.4s14.328-17.67 23.941-21.343a31.955 31.955 0 0 1 11.368-2.221a23.483 23.483 0 0 1 11.847 3.058c14.358 8.285 17.023 31.682 5.82 51.087m-143.704-47.865a54.965 54.965 0 0 0 11.689 1.261c13.12 0 26.1-4.729 37.458-13.812c13.759-11 23.557-27.272 27.588-45.818c8.589-39.515-11.858-77.627-45.58-84.957c-16.957-3.686-34.412.77-49.147 12.551c-13.76 11-23.558 27.272-27.588 45.817c-8.59 39.515 11.857 77.627 45.58 84.958m-14.31-78.16c2.474-11.387 8.263-21.2 16.3-27.621c5.572-4.454 11.689-6.781 17.637-6.781a22.258 22.258 0 0 1 4.732.5c16.2 3.522 25.866 25 21.107 46.89c-2.476 11.387-8.265 21.2-16.3 27.622c-7.061 5.646-15 7.874-22.367 6.275c-16.203-3.517-25.869-24.993-21.109-46.885"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Aperture(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256.684 16.736A239.3 239.3 0 0 0 87.475 425.245a239.3 239.3 0 0 0 338.419-338.42a237.736 237.736 0 0 0-169.21-70.089m-9.9 32.242L146.033 224.237l-53.985-94.011a207.136 207.136 0 0 1 154.739-81.248Zm56.437 127.035l45.912 79.413l-46.2 80.791h-92.6l-45.859-79.859l46.189-80.345ZM72.648 160.7l100.788 175.517H65.526A207.1 207.1 0 0 1 72.648 160.7m9.791 207.515h202.2l-53.494 93.542a207.584 207.584 0 0 1-148.706-93.54Zm184.818 94.849L367.668 287.48l54.168 93.692a207.167 207.167 0 0 1-154.579 81.894ZM441.125 350.6L340.187 176.013h107.721a207.133 207.133 0 0 1-6.783 174.587M229.063 144.013l53.825-93.627a207.609 207.609 0 0 1 148.147 93.627Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Apple(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M452.415 213.048c-10.609-27.192-27.511-48.256-48.92-61.078a157.712 157.712 0 0 1-17.583 26.757c.038.022.077.041.115.063c31.881 18.323 50.423 65.148 45.091 113.871c-8.833 80.721-33.35 136.043-69.036 155.775c-23.2 12.827-52.133 11-86-5.424l-3.308-1.6h-24.389l-3.307 1.6c-33.867 16.426-62.8 18.251-86 5.424c-35.685-19.732-60.2-75.054-69.036-155.775c-5.332-48.723 13.211-95.549 45.091-113.871a66.626 66.626 0 0 1 33.74-8.768c24.143 0 51.966 11.311 82.2 33.656l1.078.8s43.583-5.299 60.849-13.138c52.021-23.617 63.5-61.156 65.536-66.254a121.886 121.886 0 0 0-1.021-93.559l-4.073-10.169l-10.949.11A122.777 122.777 0 0 0 242.039 159.7c-18.9-10.59-37.278-17.343-54.884-20.14c-24.943-3.965-47.811-.1-67.968 11.486c-22.138 12.724-39.581 34.164-50.442 62c-9.874 25.307-13.608 54.817-10.514 83.094c10.142 92.681 39.659 155.027 85.361 180.3a99.908 99.908 0 0 0 49.1 12.543c19.585 0 40.629-5.194 62.975-15.575h9.83c42.394 19.693 80.085 20.718 112.071 3.032c45.7-25.271 75.221-87.617 85.363-180.3c3.092-28.275-.642-57.785-10.516-83.092M328.019 60.826a90.513 90.513 0 0 1 23.693-6.564a90.8 90.8 0 0 1-75.056 115.205a90.839 90.839 0 0 1 51.363-108.641"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Applications(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M88 32a64 64 0 1 0 64 64a64.072 64.072 0 0 0-64-64m0 96a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32m168-96a64 64 0 1 0 64 64a64.072 64.072 0 0 0-64-64m0 96a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32m168 32a64 64 0 1 0-64-64a64.072 64.072 0 0 0 64 64m0-96a32 32 0 1 1-32 32a32.036 32.036 0 0 1 32-32M88 192a64 64 0 1 0 64 64a64.072 64.072 0 0 0-64-64m0 96a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32m168-96a64 64 0 1 0 64 64a64.072 64.072 0 0 0-64-64m0 96a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32m168-96a64 64 0 1 0 64 64a64.072 64.072 0 0 0-64-64m0 96a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32M88 352a64 64 0 1 0 64 64a64.072 64.072 0 0 0-64-64m0 96a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32m168-96a64 64 0 1 0 64 64a64.072 64.072 0 0 0-64-64m0 96a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32m168-96a64 64 0 1 0 64 64a64.072 64.072 0 0 0-64-64m0 96a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ApplicationsSettings(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M88 160a64 64 0 1 0-64-64a64.072 64.072 0 0 0 64 64m0-96a32 32 0 1 1-32 32a32.036 32.036 0 0 1 32-32m168-32a64 64 0 1 0 64 64a64.072 64.072 0 0 0-64-64m0 96a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32m168 32a64 64 0 1 0-64-64a64.072 64.072 0 0 0 64 64m0-96a32 32 0 1 1-32 32a32.036 32.036 0 0 1 32-32M88 192a64 64 0 1 0 64 64a64.072 64.072 0 0 0-64-64m0 96a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32m168-96a64 64 0 1 0 64 64a64.072 64.072 0 0 0-64-64m0 96a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32m168-96a64 64 0 1 0 64 64a64.072 64.072 0 0 0-64-64m0 96a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32m0 64a64 64 0 1 0 64 64a64.072 64.072 0 0 0-64-64m0 96a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32M56 408h32v32H56zm96 0h32v32h-32zm96 0h32v32h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Apps(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M88 32a64 64 0 1 0 64 64a64.072 64.072 0 0 0-64-64m0 96a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32m168-96a64 64 0 1 0 64 64a64.072 64.072 0 0 0-64-64m0 96a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32m168 32a64 64 0 1 0-64-64a64.072 64.072 0 0 0 64 64m0-96a32 32 0 1 1-32 32a32.036 32.036 0 0 1 32-32M88 192a64 64 0 1 0 64 64a64.072 64.072 0 0 0-64-64m0 96a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32m168-96a64 64 0 1 0 64 64a64.072 64.072 0 0 0-64-64m0 96a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32m168-96a64 64 0 1 0 64 64a64.072 64.072 0 0 0-64-64m0 96a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32M88 352a64 64 0 1 0 64 64a64.072 64.072 0 0 0-64-64m0 96a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32m168-96a64 64 0 1 0 64 64a64.072 64.072 0 0 0-64-64m0 96a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32m168-96a64 64 0 1 0 64 64a64.072 64.072 0 0 0-64-64m0 96a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppsSettings(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M88 160a64 64 0 1 0-64-64a64.072 64.072 0 0 0 64 64m0-96a32 32 0 1 1-32 32a32.036 32.036 0 0 1 32-32m168-32a64 64 0 1 0 64 64a64.072 64.072 0 0 0-64-64m0 96a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32m168 32a64 64 0 1 0-64-64a64.072 64.072 0 0 0 64 64m0-96a32 32 0 1 1-32 32a32.036 32.036 0 0 1 32-32M88 192a64 64 0 1 0 64 64a64.072 64.072 0 0 0-64-64m0 96a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32m168-96a64 64 0 1 0 64 64a64.072 64.072 0 0 0-64-64m0 96a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32m168-96a64 64 0 1 0 64 64a64.072 64.072 0 0 0-64-64m0 96a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32m0 64a64 64 0 1 0 64 64a64.072 64.072 0 0 0-64-64m0 96a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32M56 408h32v32H56zm96 0h32v32h-32zm96 0h32v32h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowBottom(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m367.997 338.75l-95.998 95.997V17.503h-32v417.242l-95.996-95.995l-22.627 22.627L256 496l134.624-134.623z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleBottom(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 16C123.452 16 16 123.452 16 256s107.452 240 240 240s240-107.452 240-240S388.548 16 256 16m147.078 387.078a207.253 207.253 0 1 1 44.589-66.125a207.332 207.332 0 0 1-44.589 66.125"/><path fill="currentColor" d="M272.112 314.481V128h-32v186.481l-75.053-75.052l-22.627 22.627l113.68 113.68l113.681-113.68l-22.627-22.627z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleLeft(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 16.042c-132.548 0-240 107.451-240 240s107.452 240 240 240s240-107.452 240-240s-107.452-240-240-240M403.078 403.12A207.253 207.253 0 1 1 447.667 337a207.364 207.364 0 0 1-44.589 66.12"/><path fill="currentColor" d="m272.614 164.987l-22.628-22.627l-113.681 113.681l113.681 113.681l22.628-22.627l-75.054-75.054H385v-32H197.56z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleRight(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256.25 16.042c-132.548 0-240 107.451-240 240s107.452 240 240 240s240-107.452 240-240s-107.45-240-240-240M403.328 403.12A207.253 207.253 0 1 1 447.917 337a207.364 207.364 0 0 1-44.589 66.12"/><path fill="currentColor" d="m239.637 164.987l75.053 75.054H128.137v32H314.69l-75.053 75.054l22.627 22.627l113.681-113.681L262.264 142.36z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleTop(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 16C123.452 16 16 123.452 16 256s107.452 240 240 240s240-107.452 240-240S388.548 16 256 16m147.078 387.078a207.253 207.253 0 1 1 44.589-66.125a207.332 207.332 0 0 1-44.589 66.125"/><path fill="currentColor" d="m142.319 241.027l22.628 22.627L240 188.602V376h32V188.602l75.053 75.052l22.628-22.627L256 127.347z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeft(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M497.333 239.999H80.092l95.995-95.995l-22.627-22.627L18.837 256L153.46 390.623l22.627-22.627l-95.997-95.997h417.243z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRight(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m359.873 121.377l-22.627 22.627l95.997 95.997H16v32.001h417.24l-95.994 95.994l22.627 22.627L494.498 256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowThickBottom(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M255.682 494.636L16 254.3v-38.276l143.937-.007V16h192v200.007L495.952 216l-.035 38.688ZM54.931 248.022l200.8 201.342L457.328 248l-137.391.008V48h-128v200.015Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowThickFromBottom(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256.2 16L56 215.993v38.632h120v144h160v-144h120V216ZM304 222.625v144h-96v-144H94.639L256.174 61.254l161.21 161.371Zm-248 240h400v32H56z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowThickFromLeft(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M296.007 56h-38.632v120h-144v160h144v120H296l200-199.8Zm-6.632 361.384V304h-144v-96h144V94.639l161.37 161.535ZM17.375 56h32v400h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowThickFromRight(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M254.625 56h-38.632L16 256.2L216 456h38.623V336h144V176h-144Zm112 152v96h-144v113.384l-161.37-161.21l161.37-161.535V208Zm96-152h32v400h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowThickFromTop(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M336 113.375H176v144H56V296l199.8 200L456 296.007v-38.632H336Zm81.361 176L255.826 450.746L94.616 289.375H208v-144h96v144ZM56 17.376h400v32H56z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowThickLeft(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m294.637 496l-38.688-.035L16 255.729L256.334 16.048h38.277l.008 143.937h200.017v192H294.629ZM61.271 255.773l201.364 201.6l-.008-137.391h200.009v-128H262.621l-.008-137.006Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowThickRight(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M254.3 496h-38.275l-.008-143.937H16v-192h200.007L216 16.048l38.688.035l239.948 240.235ZM48 320.063h200.015l.007 137.006l201.342-200.8L248 54.672l.008 137.391H48Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowThickToBottom(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M336 176.005V16H176v160H56v38.623l199.8 200L456 214.637v-38.632Zm-80.174 193.371L94.616 208.005H208V48h96v160h113.361ZM56 464h400v32H56z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowThickToLeft(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M336 176V56h-38.627l-200 199.8L297.363 456H336V336h160V176Zm128 128H304v113.361L142.625 255.826L304 94.616V208h160ZM16 56h32v400H16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowThickToRight(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M176 56v120H16v160h160v120h38.623l200-199.8L214.635 56Zm32 361.384V304H48v-96h160V94.639l161.373 161.535ZM463.998 56h32v400h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowThickToTop(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M56 297.365V336h120v160h160V336h120v-38.626l-199.8-200ZM304 304v160h-96V304H94.639l161.535-161.37L417.384 304ZM56 16.002h400v32H56z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowThickTop(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M352.062 496h-192V295.993L16.047 296l.037-38.688L256.318 17.364L496 257.7v38.278l-143.938.006Zm-160-32h128V263.984l137.007-.006L256.274 62.636L54.672 264l137.39-.008Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowTop(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M390.624 150.625L256 16L121.376 150.625l22.628 22.627l95.997-95.998v417.982h32V77.257l95.995 95.995z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AssistiveListeningSystem(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m277.139 85.414l-8.428 31.021a154.494 154.494 0 0 1 125.555 119.956l31.921-3.6A186.588 186.588 0 0 0 277.139 85.414"/><path fill="currentColor" d="M422.061 89.939A250.681 250.681 0 0 0 294.6 21.146l-8.425 31.005c88.461 17.4 158.108 87.873 174.267 176.767l31.9-3.6a250.6 250.6 0 0 0-70.281-135.379m-188.955 58.814q-1.576-.021-3.146 0A124.083 124.083 0 0 0 107.022 272.83h32.237a92.091 92.091 0 0 1 91.062-92.085q1.05-.009 2.1 0c49.729.69 90.451 41.722 90.774 91.465a91.233 91.233 0 0 1-28.62 67.331a24.021 24.021 0 0 0-7.408 17.333v50.686A56.5 56.5 0 0 1 230.736 464H222.1v32h8.869a88.534 88.534 0 0 0 88.434-88.435v-47.318a124.615 124.615 0 0 0 36.027-88.24a124.438 124.438 0 0 0-122.324-123.254"/><path fill="currentColor" d="M283.141 272.83a52.03 52.03 0 1 0-104.059 0h32a20.03 20.03 0 1 1 40.059 0ZM216 304h32v32h-32zM112 408h32v32h-32zm11.65-67.649l22.626-22.627l88 88l-22.626 22.627z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Asterisk(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M286.903 256L416 177.877v-36.762l-.283-.469L272 227.617V72h-32v155.617L96.283 140.646l-.283.469v36.762L225.097 256L96 334.123v36.762l.283.469L240 284.383V440h32V284.383l143.717 86.971l.283-.469v-36.762z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AsteriskCircle(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M272 227.616V121.388h-32v106.229l-98.623-59.682l-2.431 4.017v31.915L225.096 256l-86.15 52.133v31.915l2.431 4.017L240 284.383v106.229h32V284.384l98.623 59.681l2.431-4.016v-31.915L286.903 256l86.151-52.134v-31.915l-2.431-4.016z"/><path fill="currentColor" d="M425.857 87.379A239.365 239.365 0 0 0 87.344 425.892A239.365 239.365 0 0 0 425.857 87.379M256.6 464c-114.341 0-207.364-93.023-207.364-207.364S142.259 49.271 256.6 49.271s207.365 93.023 207.365 207.365S370.942 464 256.6 464"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func At(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M495.826 232a206.644 206.644 0 0 0-18.882-78.412a227.033 227.033 0 0 0-51.61-71.261C379.708 39.555 319.571 16 256 16A240 240 0 0 0 86.294 425.706a240 240 0 0 0 337.671 1.722l-22.4-22.856A206.824 206.824 0 0 1 256 464c-114.691 0-208-93.309-208-208S141.309 48 256 48c112.748 0 208 87.925 208 192v36c0 28.673-25.122 52-56 52s-56-23.327-56-52V160h-32v26.751a99.988 99.988 0 1 0 12.55 132.437C347.956 343.62 376.01 360 408 360c48.523 0 88-37.682 88-84v-44ZM252 328a68 68 0 1 1 68-68a68.077 68.077 0 0 1-68 68"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Audio(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m110.763 110.763l-22.7-22.7c-.095.1-.193.186-.288.281a238.483 238.483 0 0 0-.7 336.573l22.7-22.7a206.144 206.144 0 0 1 .988-291.462Zm314.306-22.415c-.4-.4-.817-.793-1.223-1.194l-22.7 22.7a206.142 206.142 0 0 1 1.5 292.8l22.7 22.7a238.492 238.492 0 0 0-.281-337Z"/><path fill="currentColor" d="M153.523 153.522a145.746 145.746 0 0 0-.989 205.944l22.617-22.617a113.8 113.8 0 0 1 .989-160.71Zm182.25 21.705a113.8 113.8 0 0 1 1.5 162.05L359.9 359.9a145.746 145.746 0 0 0-1.5-207.285Zm-41.007 41.007a55.914 55.914 0 1 0 17.658 40.759a55.783 55.783 0 0 0-17.658-40.759m-38.342 64.759a24 24 0 1 1 24-24a24 24 0 0 1-24 24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AudioDescription(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M472 64H40a24.028 24.028 0 0 0-24 24v336a24.028 24.028 0 0 0 24 24h432a24.028 24.028 0 0 0 24-24V88a24.028 24.028 0 0 0-24-24m-8 352H48V96h416Z"/><path fill="currentColor" d="M134.2 288h59.6l18.667 56H246.2l-58.668-176h-47.064L81.8 344h33.731Zm29.333-88h.936l18.667 56h-38.27ZM424 263.333v-14.666A80.758 80.758 0 0 0 343.333 168H280v176h63.333A80.758 80.758 0 0 0 424 263.333M312 200h31.333A48.722 48.722 0 0 1 392 248.667v14.666A48.722 48.722 0 0 1 343.333 312H312Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AudioSpectrum(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 160h32v192H16zm360 0h32v192h-32zM104 88h32v328h-32zm184 8h32v320h-32zm176 0h32v320h-32zM192 16h32v480h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AvTimer(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M240 384h32v32h-32zM96 240h32v32H96zm288 0h32v32h-32z"/><path fill="currentColor" d="M414.392 97.608A222.332 222.332 0 0 0 272 32.567V32h-32v96h32V64.672C370.41 72.83 448 155.519 448 256c0 105.869-86.131 192-192 192S64 361.869 64 256a191.61 191.61 0 0 1 56.408-135.942l115.624 145.88l25.078-19.876L124.6 73.828l-12.606 10.59a224 224 0 1 0 302.4 13.19Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Baby(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M425.39 200.035A184.3 184.3 0 0 0 290.812 91.289l26.756-42.809l-27.136-16.96l-35.305 56.488A184.046 184.046 0 0 0 86.61 200.035a71.978 71.978 0 0 0 0 143.93a184.071 184.071 0 0 0 338.78 0a71.978 71.978 0 0 0 0-143.93m27.152 99.975a39.77 39.77 0 0 1-27.76 11.961l-20.725.394l-8.113 19.074a152.066 152.066 0 0 1-279.887 0l-8.114-19.074l-20.725-.394a39.978 39.978 0 0 1 0-79.942l20.725-.394l8.114-19.074a152.067 152.067 0 0 1 279.887 0l8.113 19.074l20.725.394a39.974 39.974 0 0 1 27.76 67.981"/><path fill="currentColor" d="M168 232h40v40h-40zm136 0h40v40h-40zm-48 152a80 80 0 0 0 80-80H176a80 80 0 0 0 80 80"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BabyCarriage(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M445.057 345.134L464 274.1V232c-8.136-93.993-87.933-168-184-168h-32v168H132.158l-17.844-78.768A32.155 32.155 0 0 0 83.038 128H16v32h67.038l40.475 178.67A80 80 0 1 0 224 416q0-4.05-.4-8h104.8q-.395 3.948-.4 8a80 80 0 1 0 117.057-70.866M280 96c78.411 0 143.145 59.678 151.164 136H280ZM144 464a48 48 0 1 1 48-48a48.055 48.055 0 0 1-48 48m194.763-88H213.237a80.166 80.166 0 0 0-57.316-39.108L139.408 264H432v5.9l-17.7 66.368a80.592 80.592 0 0 0-6.3-.271A80.026 80.026 0 0 0 338.763 376M408 464a48 48 0 1 1 48-48a48.055 48.055 0 0 1-48 48"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Backspace(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M227.313 363.313L312 278.627l84.687 84.686l22.626-22.626L334.627 256l84.686-84.687l-22.626-22.626L312 233.373l-84.687-84.686l-22.626 22.626L289.373 256l-84.686 84.687z"/><path fill="currentColor" d="M472 64H194.644a24.091 24.091 0 0 0-17.42 7.492L16 241.623v28.754l161.224 170.131a24.091 24.091 0 0 0 17.42 7.492H472a24.028 24.028 0 0 0 24-24V88a24.028 24.028 0 0 0-24-24m-8 352H198.084L48 257.623v-3.246L198.084 96H464Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Badge(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m328.375 384l3.698 74.999l-75.862-52.719l-76.287 52.769L183.625 384h-32.039l-5.522 112h36.692l73.413-50.78L329.242 496h36.694l-5.522-112zm87.034-229.086l-2.194-48.054L372.7 80.933l-25.932-40.519l-48.055-2.2L256 16.093l-42.713 22.126l-48.055 2.2L139.3 80.933L98.785 106.86l-2.194 48.054l-22.127 42.714l22.127 42.715l2.2 48.053l40.509 25.927l25.928 40.52l48.055 2.195L256 379.164l42.713-22.126l48.055-2.195l25.928-40.52l40.518-25.923l2.195-48.053l22.127-42.715Zm-31.646 76.949L382 270.377l-32.475 20.78l-20.78 32.475l-38.515 1.76L256 343.125l-34.234-17.733l-38.515-1.76l-20.78-32.475L130 270.377l-1.759-38.514l-17.741-34.235l17.737-34.228L130 124.88l32.471-20.78l20.78-32.474l38.515-1.76L256 52.132l34.234 17.733l38.515 1.76l20.78 32.474L382 124.88l1.759 38.515l17.741 34.233Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BalanceScale(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m203.247 386.414l4.753-5.229V355.4L130.125 191h-36.25L16 355.4v27.042l4.234 4.595a124.347 124.347 0 0 0 91.224 39.982h.42a124.343 124.343 0 0 0 91.369-40.605M176 368.608a90.924 90.924 0 0 1-64.231 26.413h-.33A90.907 90.907 0 0 1 48 369.667V362.6l64-135.112L176 362.6ZM418.125 191h-36.25L304 355.4v27.042l4.234 4.595a124.347 124.347 0 0 0 91.224 39.982h.42a124.343 124.343 0 0 0 91.369-40.607l4.753-5.227V355.4ZM464 368.608a90.924 90.924 0 0 1-64.231 26.413h-.33A90.907 90.907 0 0 1 336 369.667V362.6l64-135.112L464 362.6Z"/><path fill="currentColor" d="M272 196.659A56.223 56.223 0 0 0 309.659 159H416v-32H309.659a55.991 55.991 0 0 0-107.318 0H96v32h106.341A56.223 56.223 0 0 0 240 196.659V463H136v32h240v-32H272ZM232 143a24 24 0 1 1 24 24a24 24 0 0 1-24-24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ban(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M425.706 86.294A240 240 0 0 0 86.294 425.705A240 240 0 0 0 425.706 86.294M256 48a207.1 207.1 0 0 1 135.528 50.345L98.345 391.528A207.1 207.1 0 0 1 48 256c0-114.691 93.309-208 208-208m0 416a207.084 207.084 0 0 1-134.986-49.887l293.1-293.1A207.084 207.084 0 0 1 464 256c0 114.691-93.309 208-208 208"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bank(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M247.759 14.358L16 125.946V184h480v-58.362ZM464 152H48v-5.946l200.241-96.412L464 146.362ZM48 408h416v32H48zm-32 56h480v32H16zm40-248h32v160H56zm368 0h32v160h-32zm-96 0h32v160h-32zm-176 0h32v160h-32zm88 0h32v160h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarChart(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M136 416h88a24.028 24.028 0 0 0 24-24V184a24.028 24.028 0 0 0-24-24h-88a24.028 24.028 0 0 0-24 24v208a24.028 24.028 0 0 0 24 24m8-224h72v192h-72ZM424 16h-88a24.028 24.028 0 0 0-24 24v352a24.028 24.028 0 0 0 24 24h88a24.028 24.028 0 0 0 24-24V40a24.028 24.028 0 0 0-24-24m-8 368h-72V48h72Z"/><path fill="currentColor" d="M48 16H16v480h480v-32H48z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Barcode(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 464h480V56H16ZM48 88h416v344H48Z"/><path fill="currentColor" d="M80 128h32v256H80zm64 0h32v192h-32zm64 0h32v256h-32zm64 0h32v192h-32zm64 0h32v192h-32zm64 0h32v256h-32zM144 352h32v32h-32zm128 0h32v32h-32zm64 0h32v32h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Baseball(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M108 200a92 92 0 1 0-92-92a92.1 92.1 0 0 0 92 92m0-152a60 60 0 1 1-60 60a60.068 60.068 0 0 1 60-60m368.937-9.74a76 76 0 0 0-107.48 0l-5.475 5.475a1172.954 1172.954 0 0 0-60.93 65.8l-.318.37l-138.829 183.562l-85.357 85.358a38.263 38.263 0 0 0-46.122 60.229l40.52 40.519a38.272 38.272 0 0 0 60.238-46.13l85.24-85.24l179.9-130.76l.841-.654a1171.036 1171.036 0 0 0 77.771-71.049a76.088 76.088 0 0 0 .001-107.48m-22.629 84.853a1139.855 1139.855 0 0 1-75.23 68.761L197.576 323.8L88.854 432.519l15.572 15.574a6.26 6.26 0 1 1-8.852 8.853l-40.52-40.519a6.26 6.26 0 0 1 8.853-8.854l15.573 15.574L188.1 314.533l139.57-184.541a1140.027 1140.027 0 0 1 58.943-63.63l5.475-5.474a44 44 0 1 1 62.225 62.225Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Basket(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M120 304h32v128h-32zm80 0h32v128h-32zm80 0h32v128h-32zm80 0h32v128h-32z"/><path fill="currentColor" d="M473.681 168L394.062 16h-36.124l79.619 152H74.443l79.619-152h-36.124L38.319 168H16v111.468L58.856 496h394.261L496 281.584V168ZM464 278.416L426.883 464H85.144L48 276.332V272h416ZM464 240H48v-40h416Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Basketball(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 16A240 240 0 0 0 86.294 425.706A240 240 0 0 0 425.706 86.294A238.432 238.432 0 0 0 256 16m-148.239 94.25q4.21 5.825 8.145 12.466c18.663 31.593 29.823 72.765 31.985 117.284H48.609a207.386 207.386 0 0 1 59.152-129.75m0 291.5A207.386 207.386 0 0 1 48.609 272h99.282c-2.162 44.519-13.322 85.691-31.985 117.284q-3.925 6.646-8.145 12.466M240 463.391a206.866 206.866 0 0 1-108.049-40.532a190.612 190.612 0 0 0 11.507-17.3c21.483-36.368 34.233-83.3 36.471-133.559H240ZM240 240h-60.071c-2.238-50.257-14.988-97.191-36.471-133.559a190.612 190.612 0 0 0-11.507-17.3A206.866 206.866 0 0 1 240 48.609Zm32-191.391a206.866 206.866 0 0 1 108.049 40.532a190.612 190.612 0 0 0-11.507 17.3c-21.483 36.368-34.233 83.3-36.471 133.559H272Zm0 414.782V272h60.071c2.238 50.257 14.988 97.191 36.471 133.559a190.612 190.612 0 0 0 11.507 17.3A206.866 206.866 0 0 1 272 463.391m132.239-61.641q-4.21-5.824-8.145-12.466c-18.663-31.593-29.823-72.765-31.985-117.284h99.282a207.386 207.386 0 0 1-59.152 129.75M364.109 240c2.162-44.519 13.322-85.691 31.985-117.284q3.925-6.646 8.145-12.466A207.386 207.386 0 0 1 463.391 240Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bath(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 280H80V100a51.258 51.258 0 0 1 15.113-36.485l.4-.4a51.691 51.691 0 0 1 58.6-10.162a79.1 79.1 0 0 0 11.778 96.627l10.951 10.951l-20.157 20.158l22.626 22.626l20.157-20.157L311.157 71.471l20.157-20.157l-22.627-22.627l-20.158 20.157l-10.951-10.951a79.086 79.086 0 0 0-100.929-8.976A83.61 83.61 0 0 0 72.887 40.485l-.4.4A83.054 83.054 0 0 0 48 100v180H16v32h32v30.7a23.95 23.95 0 0 0 1.232 7.589L79 439.589A23.969 23.969 0 0 0 101.766 456h12.9L103 496h33.333L148 456h208.1l12 40h33.4l-12-40h20.73A23.969 23.969 0 0 0 433 439.589l29.766-89.3A23.982 23.982 0 0 0 464 342.7V312h32v-32ZM188.52 60.52a47.025 47.025 0 0 1 66.431 0L265.9 71.471L199.471 137.9l-10.951-10.949a47.027 47.027 0 0 1 0-66.431M432 341.4L404.468 424H107.532L80 341.4V312h352Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bathroom(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 280H80V100a51.258 51.258 0 0 1 15.113-36.485l.4-.4a51.691 51.691 0 0 1 58.6-10.162a79.1 79.1 0 0 0 11.778 96.627l10.951 10.951l-20.157 20.158l22.626 22.626l20.157-20.157L311.157 71.471l20.157-20.157l-22.627-22.627l-20.158 20.157l-10.951-10.951a79.086 79.086 0 0 0-100.929-8.976A83.61 83.61 0 0 0 72.887 40.485l-.4.4A83.054 83.054 0 0 0 48 100v180H16v32h32v30.7a23.95 23.95 0 0 0 1.232 7.589L79 439.589A23.969 23.969 0 0 0 101.766 456h12.9L103 496h33.333L148 456h208.1l12 40h33.4l-12-40h20.73A23.969 23.969 0 0 0 433 439.589l29.766-89.3A23.982 23.982 0 0 0 464 342.7V312h32v-32ZM188.52 60.52a47.025 47.025 0 0 1 66.431 0L265.9 71.471L199.471 137.9l-10.951-10.949a47.027 47.027 0 0 1 0-66.431M432 341.4L404.468 424H107.532L80 341.4V312h352Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryAlert(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M432 176v-56.41A23.825 23.825 0 0 0 408 96H240v32h160v80h64v104h-64v80H47.986V128H112V96H39.986a23.825 23.825 0 0 0-24 23.59v280.82a23.825 23.825 0 0 0 24 23.59H408a23.825 23.825 0 0 0 24-23.59V344h64V176Z"/><path fill="currentColor" d="M192 64v-8h-32v200h32zm0 240v-16h-32v32h32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryEmpty(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M39.986 424H408a23.825 23.825 0 0 0 24-23.59V344h64V176h-64v-56.41A23.825 23.825 0 0 0 408 96H39.986a23.825 23.825 0 0 0-24 23.59v280.82a23.825 23.825 0 0 0 24 23.59m8-296H400v80h64v104h-64v80H47.986Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryFive(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M80 160h32v200H80zm64 0h32v200h-32zm64 0h32v200h-32zm64 0h32v200h-32zm64 0h32v200h-32z"/><path fill="currentColor" d="M432 176v-56.41A23.825 23.825 0 0 0 408 96H39.986a23.825 23.825 0 0 0-24 23.59v280.82a23.825 23.825 0 0 0 24 23.59H408a23.825 23.825 0 0 0 24-23.59V344h64V176Zm32 136h-64v80H47.986V128H400v80h64Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryFull(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M80 160h32v200H80zm64 0h32v200h-32zm64 0h32v200h-32zm64 0h32v200h-32zm64 0h32v200h-32z"/><path fill="currentColor" d="M432 176v-56.41A23.825 23.825 0 0 0 408 96H39.986a23.825 23.825 0 0 0-24 23.59v280.82a23.825 23.825 0 0 0 24 23.59H408a23.825 23.825 0 0 0 24-23.59V344h64V176Zm32 136h-64v80H47.986V128H400v80h64Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatterySlash(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M432 176v-56.41A23.825 23.825 0 0 0 408 96H163.882l32 32H400v80h64v104h-48v32h80V176Zm-281.373-48l-32-32l-80-80H16v22.627L73.373 96H39.986a23.825 23.825 0 0 0-24 23.59v280.82a23.825 23.825 0 0 0 24 23.59h361.387l72 72H496v-22.627L266.563 243.937ZM47.986 392V128h57.387l264 264Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryThree(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M80 160h32v200H80zm64 0h32v200h-32zm64 0h32v200h-32z"/><path fill="currentColor" d="M432 176v-56.41A23.825 23.825 0 0 0 408 96H39.986a23.825 23.825 0 0 0-24 23.59v280.82a23.825 23.825 0 0 0 24 23.59H408a23.825 23.825 0 0 0 24-23.59V344h64V176Zm32 136h-64v80H47.986V128H400v80h64Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryZero(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M39.986 424H408a23.825 23.825 0 0 0 24-23.59V344h64V176h-64v-56.41A23.825 23.825 0 0 0 408 96H39.986a23.825 23.825 0 0 0-24 23.59v280.82a23.825 23.825 0 0 0 24 23.59m8-296H400v80h64v104h-64v80H47.986Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BeachAccess(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m259.431 268.8l140-140l-27.785-27.785A208.333 208.333 0 0 0 77.019 395.646l27.781 27.785l132-132L401.372 456h45.256ZM224.333 72a175.182 175.182 0 0 1 124.686 51.646l5.157 5.158l-57.058 57.058a477.658 477.658 0 0 0-62.879-53.924c-25.216-17.838-49.439-30.329-71.994-37.131a152.909 152.909 0 0 0-17.092-4.129A175.58 175.58 0 0 1 224.333 72M104.8 378.176l-5.158-5.157a176.637 176.637 0 0 1-32.964-203.866a153.129 153.129 0 0 0 4.129 17.092c6.8 22.556 19.3 46.778 37.131 71.994a477.658 477.658 0 0 0 53.924 62.879Zm79.7-79.7c-11.857-11.634-32.231-32.977-50.438-58.718c-22.872-32.336-46.59-77.9-33.753-115.45c37.421-12.793 82.8 10.736 115.005 33.437c25.864 18.233 47.431 38.815 59.158 50.759Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Beaker(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m80 495.918l352.039.08h.006a24 24 0 0 0 24-24.008L456 354.472a23.9 23.9 0 0 0-4.239-13.613L344 184.511V121h40a24.028 24.028 0 0 0 24-24V16H104v81a24.027 24.027 0 0 0 24 24h39.917v64.621L60.276 340.834A23.9 23.9 0 0 0 56 354.509v117.409a24.029 24.029 0 0 0 24 24m119.917-300.287V89H136V48h240v41h-64v105.47L376.465 288H135.859ZM88 357.011L113.667 320H398.52L424 356.971L424.037 464L88 463.92Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bed(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M416 224H208v120.619h-22.154v-41A87.716 87.716 0 0 0 98.229 216H48v-64H16v344h32v-47.743l416 3.328V496h32V304a80.091 80.091 0 0 0-80-80M48 248h50.229a55.68 55.68 0 0 1 55.617 55.617v41H48Zm416 171.584l-416-3.328v-39.637h416Zm0-74.965H240V256h176a48.055 48.055 0 0 1 48 48Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bell(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m450.27 348.569l-43.67-80.624V184c0-83.813-68.187-152-152-152s-152 68.187-152 152v83.945l-43.672 80.623A24 24 0 0 0 80.031 384h86.935a88.866 88.866 0 0 0-.367 8a88 88 0 0 0 176 0c0-2.7-.129-5.364-.367-8h86.935a24 24 0 0 0 21.1-35.431ZM310.6 392a56 56 0 1 1-111.419-8h110.837a56.14 56.14 0 0 1 .582 8M93.462 352l41.138-75.945V184a120 120 0 0 1 240 0v92.055L415.736 352Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellExclamation(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M451.671 348.569L408 267.945V184c0-83.813-68.187-152-152-152s-152 68.187-152 152v83.945l-43.671 80.623A24 24 0 0 0 81.432 384h86.944a87.762 87.762 0 0 0-.376 8a88 88 0 0 0 176 0c0-2.7-.135-5.364-.376-8h86.944a24 24 0 0 0 21.1-35.431ZM312 392a56 56 0 1 1-111.418-8h110.836a55.85 55.85 0 0 1 .582 8M94.863 352L136 276.055V184a120 120 0 0 1 240 0v92.055L417.137 352Z"/><path fill="currentColor" d="M240 112h32v136h-32zm0 168h32v32h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bike(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M116 468A100 100 0 1 0 16 368a100.113 100.113 0 0 0 100 100m0-168a68 68 0 1 1-68 68a68.077 68.077 0 0 1 68-68m180 68a100 100 0 1 0 100-100a100.113 100.113 0 0 0-100 100m100-68a68 68 0 1 1-68 68a68.077 68.077 0 0 1 68-68"/><circle cx="317.912" cy="94.088" r="34.088" fill="currentColor"/><path fill="currentColor" d="M190.954 266.3L240 314.69V404h32v-92.655a24.154 24.154 0 0 0-7.144-17.084l-51.274-50.588l66.453-66.453l58.165 59.551A24.14 24.14 0 0 0 355.369 244H384v-32h-25.262l-92.487-94.688l-.475.464l-8.4-8.4l-112 112l45.254 45.254Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BirthdayCake(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M422 226.067H312v-96h-36A85.425 85.425 0 0 0 293.054 78.5c0-27.64-13.079-53.611-34.133-67.776l-8.932-6.01l-8.931 6.01C220 24.891 206.925 50.861 206.925 78.5a85.425 85.425 0 0 0 17.059 51.566H184v96H90a58.066 58.066 0 0 0-58 58V464a32.036 32.036 0 0 0 32 32h384a32.036 32.036 0 0 0 32-32V284.067a58.066 58.066 0 0 0-58-58M249.989 45.542c6.99 8.684 11.065 20.466 11.065 32.959s-4.075 24.276-11.065 32.959c-6.989-8.683-11.064-20.466-11.064-32.96S243 54.226 249.989 45.542M216 162.067h64v64h-64Zm-152 122a26.03 26.03 0 0 1 26-26h332a26.03 26.03 0 0 1 26 26v31.577l-21.6 9.531a33.284 33.284 0 0 1-26.809 0L362 308.588l-37.6 16.586a33.283 33.283 0 0 1-26.81 0L260 308.587l-37.6 16.586a33.279 33.279 0 0 1-26.81 0L158 308.588l-37.593 16.585a33.279 33.279 0 0 1-26.81 0L64 312.117ZM448 464H64V347.093l16.678 7.358a65.355 65.355 0 0 0 52.644 0L158 343.563l24.679 10.888a65.353 65.353 0 0 0 52.643 0L260 343.563l24.677 10.888a65.351 65.351 0 0 0 52.642 0L362 343.563l24.678 10.889a65.354 65.354 0 0 0 52.641 0l8.693-3.835L448.02 464Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Blind(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m374.906 253.877l11.36-10.328l-109.911-120.9h-68.492L163.585 151.4A60.364 60.364 0 0 0 136 202.2V274h32v-71.8a28.477 28.477 0 0 1 13.013-23.967L208 160.712v93.836l105.7 220.381h32.942L272.714 315.7V166.214l74.628 82.086l82.409 226.626h25.537ZM214.7 96.861a34.081 34.081 0 1 0-10.871-24.949A33.96 33.96 0 0 0 214.7 96.861"/><path fill="currentColor" d="M132.796 474.929h34.465l57.514-143.785l-18.793-39.181z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bluetooth(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M150.627 376L232 294.627V496h38.627l120-120l-120-120l120-120l-120-120H232v201.373L150.627 136h-45.254l120 120l-120 120ZM264 54.627L345.373 136L264 217.373Zm0 240L345.373 376L264 457.373Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Blur(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M87.371 161.022A71.371 71.371 0 1 0 16 89.652a71.45 71.45 0 0 0 71.371 71.37m0-110.741A39.371 39.371 0 1 1 48 89.652a39.415 39.415 0 0 1 39.371-39.371m186.94 114.693a59.468 59.468 0 1 0-59.468-59.467a59.534 59.534 0 0 0 59.468 59.467m0-86.935a27.468 27.468 0 1 1-27.468 27.468a27.5 27.5 0 0 1 27.468-27.468"/><circle cx="435.516" cy="115.386" r="29.637" fill="currentColor"/><path fill="currentColor" d="M87.371 328.511A71.371 71.371 0 1 0 16 257.141a71.45 71.45 0 0 0 71.371 71.37m0-110.741A39.371 39.371 0 1 1 48 257.141a39.415 39.415 0 0 1 39.371-39.371m186.94 98.838a59.468 59.468 0 1 0-59.468-59.467a59.534 59.534 0 0 0 59.468 59.467m0-86.935a27.468 27.468 0 1 1-27.468 27.468a27.5 27.5 0 0 1 27.468-27.468"/><circle cx="435.516" cy="257.141" r="29.637" fill="currentColor"/><path fill="currentColor" d="M87.371 496A71.371 71.371 0 1 0 16 424.629A71.451 71.451 0 0 0 87.371 496m0-110.742A39.371 39.371 0 1 1 48 424.629a39.415 39.415 0 0 1 39.371-39.371m186.94 82.984a59.468 59.468 0 1 0-59.468-59.468a59.534 59.534 0 0 0 59.468 59.468m0-86.935a27.468 27.468 0 1 1-27.468 27.467a27.5 27.5 0 0 1 27.468-27.467"/><circle cx="435.516" cy="398.895" r="29.637" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BlurCircular(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 16A240 240 0 0 0 86.294 425.705A240 240 0 0 0 425.706 86.294A238.432 238.432 0 0 0 256 16m147.078 387.078c-81.1 81.1-213.058 81.1-294.157 0s-81.1-213.057 0-294.156a208.238 208.238 0 0 1 294.157 0c81.099 81.099 81.099 213.057 0 294.156"/><path fill="currentColor" d="M197.483 242.382a46.332 46.332 0 1 0-32.776-13.555a46.206 46.206 0 0 0 32.776 13.555M187.334 185.9a14.354 14.354 0 1 1 0 20.3a14.311 14.311 0 0 1 0-20.3m127.183 56.482a46.344 46.344 0 1 0-32.777-79.109a46.332 46.332 0 0 0 32.777 79.108Zm-10.15-56.482a14.354 14.354 0 1 1 0 20.3a14.371 14.371 0 0 1 0-20.3m-139.66 97.273a46.353 46.353 0 1 0 65.553 0a46.048 46.048 0 0 0-65.553 0m42.926 42.927a14.347 14.347 0 1 1 0-20.3a14.372 14.372 0 0 1 0 20.3m74.107-42.927a46.354 46.354 0 1 0 65.553 0a46.406 46.406 0 0 0-65.553 0m42.926 42.927a14.354 14.354 0 1 1 4.2-10.15a14.372 14.372 0 0 1-4.2 10.15"/><circle cx="314.517" cy="98.5" r="29.637" fill="currentColor" transform="rotate(-9.217 314.534 98.505)"/><circle cx="197.483" cy="98.5" r="29.637" fill="currentColor" transform="rotate(-9.217 197.493 98.505)"/><circle cx="314.517" cy="413.5" r="29.637" fill="currentColor" transform="rotate(-67.5 314.517 413.5)"/><circle cx="197.483" cy="413.5" r="29.637" fill="currentColor" transform="rotate(-67.5 197.483 413.5)"/><circle cx="413.5" cy="314.517" r="29.637" fill="currentColor" transform="rotate(-22.5 413.5 314.518)"/><circle cx="413.5" cy="197.483" r="29.637" fill="currentColor" transform="rotate(-58.283 413.496 197.483)"/><circle cx="98.5" cy="314.517" r="29.637" fill="currentColor"/><circle cx="98.5" cy="197.483" r="29.637" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BlurLinear(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 16h480v32H16zm0 448h480v32H16zm71.371-232.533A71.371 71.371 0 1 0 16 160.1a71.45 71.45 0 0 0 71.371 71.367m0-110.741A39.371 39.371 0 1 1 48 160.1a39.415 39.415 0 0 1 39.371-39.374m186.94 110.741A59.467 59.467 0 1 0 214.843 172a59.533 59.533 0 0 0 59.468 59.467m0-86.934A27.467 27.467 0 1 1 246.843 172a27.5 27.5 0 0 1 27.468-27.467m161.205 70.935a29.637 29.637 0 1 0-29.637-29.637a29.637 29.637 0 0 0 29.637 29.637M87.371 427.35A71.371 71.371 0 1 0 16 355.979a71.451 71.451 0 0 0 71.371 71.371m0-110.742A39.371 39.371 0 1 1 48 355.979a39.415 39.415 0 0 1 39.371-39.371m186.94 86.935a59.468 59.468 0 1 0-59.468-59.467a59.534 59.534 0 0 0 59.468 59.467m0-86.935a27.468 27.468 0 1 1-27.468 27.468a27.5 27.5 0 0 1 27.468-27.468m161.205 43.274a29.637 29.637 0 1 0-29.637-29.637a29.637 29.637 0 0 0 29.637 29.637"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoatAlt(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M410.866 181.063A32.1 32.1 0 0 0 380.793 160h-39.239l-16-96H221.112l16 96h-39.558l-16-96H77.112l16 96H48v88H16v200h427.727L496 282.466V248h-60.793ZM298.446 96l10.667 64h-39.559l-10.666-64Zm-144 0l10.667 64h-39.559l-10.666-64ZM80 192h300.793l20.363 56H80Zm383.222 88l-42.949 136H48V280Z"/><path fill="currentColor" d="M80 320h32v32H80zm64 0h32v32h-32zm64 0h32v32h-32zm64 0h32v32h-32zm64 0h32v32h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bold(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 496h480V16H16ZM48 48h416v416H48Z"/><path fill="currentColor" d="M341.6 242.986A79.956 79.956 0 0 0 280 112H112v32h40v224h-40v32h208a79.991 79.991 0 0 0 21.6-157.014M208 144h48a48 48 0 0 1 0 96h-48Zm88 224h-88v-96h88a48 48 0 0 1 0 96"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bolt(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m331.464 192l77-176H147.879l-81 288H187.9l-39.53 192h58.851l268.235-304ZM192.779 464h-5.149l39.529-192H109.121l63-224h187.415l-77 176h122.009Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoltCircle(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M425.706 86.294A240 240 0 0 0 86.294 425.706A240 240 0 0 0 425.706 86.294M256 464c-114.691 0-208-93.309-208-208S141.309 48 256 48s208 93.309 208 208s-93.309 208-208 208"/><path fill="currentColor" d="M372.529 120H180.468L121.8 296h76.047l-36.4 104h77.179L384 254.627V208h-65.249ZM352 240v1.373L225.373 368h-18.821l36.4-104H166.2l37.333-112h111.938l-53.778 88Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Book(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M101.667 400H464V16H100.667A60.863 60.863 0 0 0 40 76.667V430.25h.011c0 .151-.011.3-.011.453c0 35.4 27.782 65.3 60.667 65.3H464V464H100.667C85.664 464 72 448.129 72 430.7c0-16.64 13.585-30.7 29.667-30.7M360 48.333v172.816l-48.4-42.49L264 220.9V48.333ZM232 48v216h31.641l48.075-42.659L360.305 264H392V48h40v320H136.08L136 48Zm-131.333 0H104l.076 320h-2.413A59.793 59.793 0 0 0 72 375.883V76.917A28.825 28.825 0 0 1 100.667 48"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookmark(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M424 496h-35.25L256.008 381.19L123.467 496H88V16h336ZM120 48v408.667l135.992-117.8L392 456.5V48Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderAll(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M56 472h400a16 16 0 0 0 16-16V56a16 16 0 0 0-16-16H56a16 16 0 0 0-16 16v400a16 16 0 0 0 16 16M272 72h168v168H272Zm0 200h168v168H272ZM72 72h168v168H72Zm0 200h168v168H72Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderBottom(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M40 440h432v32H40zm0-133.817h32v33.454H40zm0 66.909h32v33.454H40zm0-133.82h32v33.455H40zm0-66.908h32v33.454H40zm0-66.909h32v33.454H40zM40 40h32v32H40zm266.182 0h33.455v32h-33.455zm-133.818 0h33.455v32h-33.455zm200.727 0h33.455v32h-33.455zm-133.819 0h33.455v32h-33.455zm-133.817 0h33.455v32h-33.455zM440 40h32v32h-32zm0 333.091h32v33.454h-32zm0-66.909h32v33.454h-32zm0-200.728h32v33.454h-32zm0 66.909h32v33.454h-32zm0 66.909h32v33.455h-32zm-200.728 0h33.455v33.455h-33.455zm0 66.91h33.455v33.454h-33.455zm0-133.819h33.455v33.454h-33.455zm0-66.909h33.455v33.454h-33.455zm0 267.637h33.455v33.454h-33.455zm66.91-133.819h33.455v33.455h-33.455zm66.909 0h33.455v33.455h-33.455zm-200.727 0h33.455v33.455h-33.455zm-66.909 0h33.455v33.455h-33.455z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderClear(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M440 440h32v32h-32zm-133.818 0h33.455v32h-33.455zm-133.818 0h33.455v32h-33.455zm200.727 0h33.455v32h-33.455zm-267.636 0h33.455v32h-33.455zM40 440h32v32H40zm0-133.817h32v33.454H40zm0 66.909h32v33.454H40zm0-133.82h32v33.455H40zm0-66.908h32v33.454H40zm0-66.909h32v33.454H40zM40 40h32v32H40zm266.182 0h33.455v32h-33.455zm-133.818 0h33.455v32h-33.455zm200.727 0h33.455v32h-33.455zm-133.819 0h33.455v32h-33.455zm-133.817 0h33.455v32h-33.455zM440 40h32v32h-32zm0 333.091h32v33.454h-32zm0-66.909h32v33.454h-32zm0-200.728h32v33.454h-32zm0 66.909h32v33.454h-32zm0 66.909h32v33.455h-32zm-200.728 0h33.455v33.455h-33.455zm0 66.91h33.455v33.454h-33.455zm0-133.819h33.455v33.454h-33.455zm0-66.909h33.455v33.454h-33.455zm0 267.637h33.455v33.454h-33.455zm0 66.909h33.455v32h-33.455zm66.91-200.728h33.455v33.455h-33.455zm66.909 0h33.455v33.455h-33.455zm-200.727 0h33.455v33.455h-33.455zm-66.909 0h33.455v33.455h-33.455z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderHorizontal(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M40 240h432v32H40zm400 200h32v32h-32zm-133.818 0h33.455v32h-33.455zm-133.818 0h33.455v32h-33.455zm200.727 0h33.455v32h-33.455zm-267.636 0h33.455v32h-33.455zM40 440h32v32H40zm0-133.817h32v33.454H40zm0 66.909h32v33.454H40zm0-200.728h32v33.454H40zm0-66.909h32v33.454H40zM40 40h32v32H40zm266.182 0h33.455v32h-33.455zm-133.818 0h33.455v32h-33.455zm200.727 0h33.455v32h-33.455zm-133.819 0h33.455v32h-33.455zm-133.817 0h33.455v32h-33.455zM440 40h32v32h-32zm0 333.091h32v33.454h-32zm0-66.909h32v33.454h-32zm0-200.728h32v33.454h-32zm0 66.909h32v33.454h-32zM239.272 306.182h33.455v33.454h-33.455zm0-133.819h33.455v33.454h-33.455zm0-66.909h33.455v33.454h-33.455zm0 267.637h33.455v33.454h-33.455zm0 66.909h33.455v32h-33.455z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderInner(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M240 272v200h32V272h200v-32H272V40h-32v200H40v32h32zm200 168h32v32h-32zm-133.818 0h33.455v32h-33.455zm-133.818 0h33.455v32h-33.455zm200.727 0h33.455v32h-33.455zm-267.636 0h33.455v32h-33.455zM40 440h32v32H40zm0-133.817h32v33.454H40zm0 66.909h32v33.454H40zm0-200.728h32v33.454H40zm0-66.909h32v33.454H40zM40 40h32v32H40zm266.182 0h33.455v32h-33.455zm-133.818 0h33.455v32h-33.455zm200.727 0h33.455v32h-33.455zm-267.636 0h33.455v32h-33.455zM440 40h32v32h-32zm0 333.091h32v33.454h-32zm0-66.909h32v33.454h-32zm0-200.728h32v33.454h-32zm0 66.909h32v33.454h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderLeft(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M72 440V40H40v432h32zm368 0h32v32h-32zm-133.818 0h33.455v32h-33.455zm-133.818 0h33.455v32h-33.455zm200.727 0h33.455v32h-33.455zm-267.636 0h33.455v32h-33.455zM306.182 40h33.455v32h-33.455zm-133.818 0h33.455v32h-33.455zm200.727 0h33.455v32h-33.455zm-133.819 0h33.455v32h-33.455zm-133.817 0h33.455v32h-33.455zM440 40h32v32h-32zm0 333.091h32v33.454h-32zm0-66.909h32v33.454h-32zm0-200.728h32v33.454h-32zm0 66.909h32v33.454h-32zm0 66.909h32v33.455h-32zm-200.728 0h33.455v33.455h-33.455zm0 66.91h33.455v33.454h-33.455zm0-133.819h33.455v33.454h-33.455zm0-66.909h33.455v33.454h-33.455zm0 267.637h33.455v33.454h-33.455zm0 66.909h33.455v32h-33.455zm66.91-200.728h33.455v33.455h-33.455zm66.909 0h33.455v33.455h-33.455zm-200.727 0h33.455v33.455h-33.455zm-66.909 0h33.455v33.455h-33.455z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderOuter(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M56 40a16 16 0 0 0-16 16v400a16 16 0 0 0 16 16h400a16 16 0 0 0 16-16V56a16 16 0 0 0-16-16Zm384 400H72V72h368Z"/><path fill="currentColor" d="M239.272 239.272h33.455v33.455h-33.455zm0 66.91h33.455v33.454h-33.455zm0-133.819h33.455v33.454h-33.455zm0-66.909h33.455v33.454h-33.455zm0 267.637h33.455v33.454h-33.455zm66.91-133.819h33.455v33.455h-33.455zm66.909 0h33.455v33.455h-33.455zm-200.727 0h33.455v33.455h-33.455zm-66.909 0h33.455v33.455h-33.455z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderRight(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M472 440V40h-32v432h32zm-165.818 0h33.455v32h-33.455zm-133.818 0h33.455v32h-33.455zm200.727 0h33.455v32h-33.455zm-267.636 0h33.455v32h-33.455zM40 440h32v32H40zm0-133.817h32v33.454H40zm0 66.909h32v33.454H40zm0-133.82h32v33.455H40zm0-66.908h32v33.454H40zm0-66.909h32v33.454H40zM40 40h32v32H40zm266.182 0h33.455v32h-33.455zm-133.818 0h33.455v32h-33.455zm200.727 0h33.455v32h-33.455zm-133.819 0h33.455v32h-33.455zm-133.817 0h33.455v32h-33.455zm133.817 199.272h33.455v33.455h-33.455zm0 66.91h33.455v33.454h-33.455zm0-133.819h33.455v33.454h-33.455zm0-66.909h33.455v33.454h-33.455zm0 267.637h33.455v33.454h-33.455zm0 66.909h33.455v32h-33.455zm66.91-200.728h33.455v33.455h-33.455zm66.909 0h33.455v33.455h-33.455zm-200.727 0h33.455v33.455h-33.455zm-66.909 0h33.455v33.455h-33.455z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderStyle(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M440 440h32v32h-32zm-133.818 0h33.455v32h-33.455zm-133.818 0h33.455v32h-33.455zm200.727 0h33.455v32h-33.455zm-267.636 0h33.455v32h-33.455zM472 40H56a16 16 0 0 0-16 16v416h32V72h400Zm-32 333.091h32v33.454h-32zm0-66.909h32v33.454h-32zm0-200.728h32v33.454h-32zm0 66.909h32v33.454h-32zm0 66.909h32v33.455h-32zM239.272 440h33.455v32h-33.455z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderTop(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M40 40h432v32H40zm400 400h32v32h-32zm-133.818 0h33.455v32h-33.455zm-133.818 0h33.455v32h-33.455zm200.727 0h33.455v32h-33.455zm-267.636 0h33.455v32h-33.455zM40 440h32v32H40zm0-133.817h32v33.454H40zm0 66.909h32v33.454H40zm0-133.82h32v33.455H40zm0-66.908h32v33.454H40zm0-66.909h32v33.454H40zm400 267.636h32v33.454h-32zm0-66.909h32v33.454h-32zm0-200.728h32v33.454h-32zm0 66.909h32v33.454h-32zm0 66.909h32v33.455h-32zm-200.728 0h33.455v33.455h-33.455zm0 66.91h33.455v33.454h-33.455zm0-133.819h33.455v33.454h-33.455zm0-66.909h33.455v33.454h-33.455zm0 267.637h33.455v33.454h-33.455zm0 66.909h33.455v32h-33.455zm66.91-200.728h33.455v33.455h-33.455zm66.909 0h33.455v33.455h-33.455zm-200.727 0h33.455v33.455h-33.455zm-66.909 0h33.455v33.455h-33.455z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderVertical(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M272 456V40h-32v432h32zm168-16h32v32h-32zm-133.818 0h33.455v32h-33.455zm-133.818 0h33.455v32h-33.455zm200.727 0h33.455v32h-33.455zm-267.636 0h33.455v32h-33.455zM40 440h32v32H40zm0-133.817h32v33.454H40zm0 66.909h32v33.454H40zm0-133.82h32v33.455H40zm0-66.908h32v33.454H40zm0-66.909h32v33.454H40zM40 40h32v32H40zm266.182 0h33.455v32h-33.455zm-133.818 0h33.455v32h-33.455zm200.727 0h33.455v32h-33.455zm-267.636 0h33.455v32h-33.455zM440 40h32v32h-32zm0 333.091h32v33.454h-32zm0-66.909h32v33.454h-32zm0-200.728h32v33.454h-32zm0 66.909h32v33.454h-32zm0 66.909h32v33.455h-32zm-133.818 0h33.455v33.455h-33.455zm66.909 0h33.455v33.455h-33.455zm-200.727 0h33.455v33.455h-33.455zm-66.909 0h33.455v33.455h-33.455z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bowling(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M312 128a183.645 183.645 0 0 0-139.789 64.5a325.008 325.008 0 0 0-5.325-11.339l11.369-90.005a66.782 66.782 0 1 0-132.51 0l11.369 90.005a323.1 323.1 0 0 0-18.758 239.811L61.843 496h100.314l17.6-56.207A183.469 183.469 0 0 0 312 496c101.458 0 184-82.542 184-184s-82.542-184-184-184M85.91 59.78a34.781 34.781 0 0 1 60.6 27.361L136.294 168H87.706L77.493 87.141A34.8 34.8 0 0 1 85.91 59.78m69.19 351.632L138.643 464H85.357L68.9 411.412A289.239 289.239 0 0 1 83.562 200h56.876q5.58 11.776 10.075 23.879a183.66 183.66 0 0 0 5.278 185.244c-.232.763-.448 1.528-.691 2.289M312 464c-83.813 0-152-68.187-152-152s68.187-152 152-152s152 68.187 152 152s-68.187 152-152 152"/><circle cx="312" cy="224" r="24" fill="currentColor"/><circle cx="384" cy="272" r="24" fill="currentColor"/><circle cx="304" cy="304" r="24" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Braille(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M112 192a48 48 0 1 0-48 48a48.055 48.055 0 0 0 48-48m-64 0a16 16 0 1 1 16 16a16.019 16.019 0 0 1-16-16m272 48a48 48 0 1 0-48-48a48.055 48.055 0 0 0 48 48m0-64a16 16 0 1 1-16 16a16.019 16.019 0 0 1 16-16m-136-64a48 48 0 1 0-48-48a48.055 48.055 0 0 0 48 48m0-64a16 16 0 1 1-16 16a16.019 16.019 0 0 1 16-16m264 96a48 48 0 1 0 48 48a48.055 48.055 0 0 0-48-48m0 64a16 16 0 1 1 16-16a16.019 16.019 0 0 1-16 16m0-96a48 48 0 1 0-48-48a48.055 48.055 0 0 0 48 48m0-64a16 16 0 1 1-16 16a16.019 16.019 0 0 1 16-16m-128 64a48 48 0 1 0-48-48a48.055 48.055 0 0 0 48 48m0-64a16 16 0 1 1-16 16a16.019 16.019 0 0 1 16-16m56 272a47.691 47.691 0 0 0-23.549 6.184a47.958 47.958 0 0 0-64-16A47.991 47.991 0 0 0 232 290.742V248a48 48 0 0 0-96 0v96.038A79.6 79.6 0 0 0 88 328a24.028 24.028 0 0 0-24 24v144h32V360.667A48.078 48.078 0 0 1 136 408v16h32V248a16 16 0 0 1 32 0v176h32v-88a16 16 0 0 1 32 0v88h32v-72a16 16 0 0 1 32 0v72h32v-56a16 16 0 0 1 32 0v128h32V368a48.055 48.055 0 0 0-48-48"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Briefcase(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M471.232 111.731H368V16H144.232v95.731h-103A24.8 24.8 0 0 0 16.464 136.5v334.032A24.8 24.8 0 0 0 41.232 495.3h430A24.8 24.8 0 0 0 496 470.532V136.5a24.8 24.8 0 0 0-24.768-24.769M176.232 48H336v63.731H176.232Zm286.232 97.269v80.286l-53.177 53.176H273V232h-33.536v46.731H103.177L50 225.555v-80.286ZM50 461.764V272.982l39.286 39.287h150.178V359.5H273v-47.231h150.178l39.286-39.287v188.782Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Brightness(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m78.462 253.835l52.522 43.279l6.26 14.276l-8.312 68.617l67.742-6.535l14.52 5.668l42.642 54.4l43.28-52.523l14.274-6.259l68.618 8.313l-6.536-67.743l5.667-14.519l54.4-42.642l-52.522-43.279l-6.26-14.276l8.312-68.618l-67.741 6.536l-14.52-5.667l-42.642-54.4l-43.279 52.522l-14.277 6.26l-68.616-8.312l6.536 67.74l-5.669 14.522Zm81.026-22.854l11.626-29.781l-3.446-35.709l37.779 4.576l29.281-12.839l22.814-27.687l23.479 29.951l29.779 11.621l35.71-3.445l-4.577 37.78l12.839 29.28l27.687 22.814l-29.951 23.479l-11.621 29.779l3.446 35.712l-37.781-4.577l-29.278 12.838l-22.815 27.688l-23.479-29.95l-29.78-11.624l-35.711 3.446l4.577-37.781l-12.839-29.28l-27.687-22.814Z"/><path fill="currentColor" d="M472 16H40a24.028 24.028 0 0 0-24 24v432a24.028 24.028 0 0 0 24 24h432a24.028 24.028 0 0 0 24-24V40a24.028 24.028 0 0 0-24-24m-8 448H48V48h416Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BritishPound(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M375.942 160a88.042 88.042 0 0 0-176.033 3.024v76.588H136v32h63.909v120.3H136v32h240v-32H231.909v-120.3H344v-32H231.909v-76.588a56.046 56.046 0 0 1 112.091 0V168h32v-8Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Browser(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M472 40H40a24.028 24.028 0 0 0-24 24v384a24.028 24.028 0 0 0 24 24h432a24.028 24.028 0 0 0 24-24V64a24.028 24.028 0 0 0-24-24M48 72h96v72H48Zm416 368H48V176h416Zm0-296H176V72h288Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Brush(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M416.941 27.429L185.407 219.3c-1.711-.078-3.421-.13-5.124-.13h-.025a113.878 113.878 0 0 0-108.472 78.655L18.3 461.387A26.873 26.873 0 0 0 43.771 496.6a27.007 27.007 0 0 0 8.4-1.345l163.562-53.483A113.877 113.877 0 0 0 294.388 333.3a114.3 114.3 0 0 0-.263-7.575L485.119 95.668a48.44 48.44 0 0 0-68.178-68.239M205.786 411.355L51.873 461.684L102.2 307.771a81.946 81.946 0 0 1 78.06-56.6c2.271 0 4.559.1 6.841.285l75 75a82.94 82.94 0 0 1 .285 6.842a81.946 81.946 0 0 1-56.6 78.057M460.5 75.227L277.382 295.791L216.6 235.009L437.359 52.067a16.44 16.44 0 0 1 23.141 23.16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrushAlt(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M56 256v240h400V256a48.055 48.055 0 0 0-48-48h-86.617l13.075-106.263l.068-.677a78.777 78.777 0 1 0-157.052 0l.027.338L190.617 208H104a48.055 48.055 0 0 0-48 48m368 208H88V320h336ZM226.8 240L209.348 98.192a46.777 46.777 0 1 1 93.3 0L285.205 240H408a16.019 16.019 0 0 1 16 16v32H88v-32a16.019 16.019 0 0 1 16-16Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bug(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m496.059 182.581l-.025-70.7l-32 .012l.017 48.172l-66.288 23.779l-45.729.007v-30.964A96.554 96.554 0 0 0 329.92 91.3l43.129-43.413h42.84v-32h-56.157l-53.987 54.344a96.815 96.815 0 0 0-100.511-.554l-53.056-53.84l-56.158.05l.029 32l42.748-.038L180.824 90.5a96.564 96.564 0 0 0-22.79 62.39v30.99l-43.235.007L48 160.093v-48.172H16v70.742l80.035 28.509l.007 84.715H16.034v32h80.01v8.01a159.741 159.741 0 0 0 9.7 54.979l-89.71 34.572v70.439h32v-48.476l71.73-27.642a159.794 159.794 0 0 0 249.578 29.044a161.477 161.477 0 0 0 23.058-29.146l71.638 27.727v48.493h32v-70.421l-89.618-34.685a159.178 159.178 0 0 0 9.614-55.1v-7.794h80v-32h-80v-84.6ZM240 463.029C176.991 455.235 128.045 401.2 128.045 335.9l-.01-120.011h30v.007H240Zm-49.966-279.154v-30.988a65 65 0 0 1 130 0v30.968Zm194 151.849A128.282 128.282 0 0 1 272 462.979V215.887h80.032v-.036h32Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Building(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M440 464V16H72v448H16v32h480v-32Zm-32 0H272v-64h-32v64H104V48h304Z"/><path fill="currentColor" d="M160 304h32v32h-32zm80 0h32v32h-32zm80 0h32v32h-32zm-160-96h32v32h-32zm80 0h32v32h-32zm80 0h32v32h-32zm-160-96h32v32h-32zm80 0h32v32h-32zm80 0h32v32h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bullhorn(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m429 15.933l-253.574 120H15.992v211.728l52.192 148.267h114.654l-46.56-128h35.847L431.182 496h64.735L496 15.933Zm-269.009 320h-56.025l.061 36.949l33.119 91.051h-46.3l-42.854-121.74V167.928h112Zm32.125-172.495l175.876-83.233v348.858l-175.875-86.949ZM463.923 464H438.66l-38.668-19.117V65.063l36.2-17.13h27.8Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Burger(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M412.717 89.012c-35.578-20.985-82.545-32.541-132.246-32.541h-40.942c-49.7 0-96.668 11.556-132.246 32.541C69.054 111.56 48 142.453 48 176v16h424v-16c0-33.547-21.054-64.44-59.283-86.988M82.8 160c14.153-40.121 80.185-71.529 156.731-71.529h40.942c76.546 0 142.578 31.408 156.731 71.529ZM48 304h424v32H48zm339.2-79.961c-24.785 0-37.77 6.125-49.227 11.529c-10.034 4.733-17.96 8.471-35.576 8.471s-25.54-3.738-35.574-8.471c-11.456-5.4-24.441-11.529-49.225-11.529s-37.769 6.125-49.225 11.529c-10.033 4.733-17.957 8.471-35.572 8.471s-25.54-3.738-35.573-8.471c-11.456-5.4-24.441-11.529-49.225-11.529v32c17.615 0 25.54 3.738 35.573 8.471c11.456 5.4 24.441 11.529 49.225 11.529s37.768-6.125 49.224-11.529c10.033-4.733 17.958-8.471 35.573-8.471s25.54 3.738 35.573 8.471c11.457 5.4 24.441 11.529 49.226 11.529s37.77-6.125 49.227-11.529c10.034-4.733 17.959-8.471 35.576-8.471s25.542 3.738 35.576 8.471c11.457 5.4 24.442 11.529 49.227 11.529v-32c-17.617 0-25.542-3.738-35.576-8.471c-11.46-5.404-24.445-11.529-49.227-11.529M48 448a24.028 24.028 0 0 0 24 24h376a24.028 24.028 0 0 0 24-24v-80H48Zm32-48h360v40H80Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Burn(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M311.145 257.6c5.5-9.032 11.185-18.371 16.8-28c18.794-32.218 22.243-64.681 10.254-96.489C314.008 68.933 233.142 29.83 214.421 23.591L201.257 19.2l-20.546 41.1l9.915 8.1c.111.089 11.124 9.712 11.858 24.3c.627 12.453-6.2 25.91-20.286 40c-9.782 9.781-20.518 19.239-31.885 29.251C102.487 204.069 48.28 251.82 48.28 342.154q0 1.861.035 3.7a151.362 151.362 0 0 0 46.444 106.49A154.177 154.177 0 0 0 202.583 496h92.1l-11.369-23.072c-46.193-93.751-13.526-147.403 27.831-215.328M202.583 464c-66.2 0-121.05-53.267-122.274-118.739q-.028-1.546-.029-3.107c0-75.878 46.356-116.713 91.186-156.2c11.239-9.9 22.862-20.138 33.359-30.637c20.754-20.753 30.719-42.365 29.619-64.232a64.963 64.963 0 0 0-7.214-26.56c27.84 15.211 67.523 44.053 81.027 79.877c8.544 22.665 5.943 45.26-7.951 69.079c-5.465 9.369-11.072 18.576-16.493 27.481C244.132 306.131 206.49 367.943 244.361 464Z"/><path fill="currentColor" d="M468.243 341.129q-.4-1.586-.834-3.185c-11.546-42.332-75.457-96.762-82.706-102.829l-14.557-12.182l-9.546 16.408c-21.753 37.4-40.421 71.512-48.559 110.212c-9.279 44.134-3.033 88.989 19.1 137.13l4.281 9.317H346.3a125.168 125.168 0 0 0 99.3-48.5a123.175 123.175 0 0 0 22.643-106.371m-47.88 86.683a93.3 93.3 0 0 1-64.44 35.7c-31.541-75.9-12.931-127.635 22.63-191.114c23.642 21.994 52.455 53.69 57.983 73.961q.349 1.283.669 2.554a91.372 91.372 0 0 1-16.842 78.899"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BusAlt(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M441.884 109.647A32.029 32.029 0 0 0 415.669 96H48a32.036 32.036 0 0 0-32 32v248h41.082a67.982 67.982 0 0 0 133.836 0h122.164a67.982 67.982 0 0 0 133.836 0H496V194.521a23.886 23.886 0 0 0-4.338-13.763ZM180 128h144v96H180Zm-132 0h100v96H48Zm76 272a36 36 0 1 1 36-36a36.04 36.04 0 0 1-36 36m256 0a36 36 0 1 1 36-36a36.04 36.04 0 0 1-36 36m84-56h-19.006a68 68 0 0 0-129.988 0H188.994a68 68 0 0 0-129.988 0H48v-88h416Zm0-120H356v-96h59.669L464 197.043Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calculator(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M472 40H40a24.028 24.028 0 0 0-24 24v384a24.028 24.028 0 0 0 24 24h432a24.028 24.028 0 0 0 24-24V64a24.028 24.028 0 0 0-24-24m-8 400H48V72h416Z"/><path fill="currentColor" d="M152 240h32v-40h40v-32h-40v-40h-32v40h-40v32h40zm44.284 45.089L168 313.373l-28.284-28.284l-22.627 22.627L145.373 336l-28.284 28.284l22.627 22.627L168 358.627l28.284 28.284l22.627-22.627L190.627 336l28.284-28.284zM288 168h112v32H288zm0 120h112v32H288zm0 64h112v32H288z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M472 96h-88V40h-32v56H160V40h-32v56H40a24.028 24.028 0 0 0-24 24v336a24.028 24.028 0 0 0 24 24h432a24.028 24.028 0 0 0 24-24V120a24.028 24.028 0 0 0-24-24m-8 352H48V128h80v40h32v-40h192v40h32v-40h80Z"/><path fill="currentColor" d="M112 224h32v32h-32zm88 0h32v32h-32zm80 0h32v32h-32zm88 0h32v32h-32zm-256 72h32v32h-32zm88 0h32v32h-32zm80 0h32v32h-32zm88 0h32v32h-32zm-256 72h32v32h-32zm88 0h32v32h-32zm80 0h32v32h-32zm88 0h32v32h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarCheck(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M472 96h-88V40h-32v56H160V40h-32v56H40a24.028 24.028 0 0 0-24 24v336a24.028 24.028 0 0 0 24 24h432a24.028 24.028 0 0 0 24-24V120a24.028 24.028 0 0 0-24-24m-8 352H48V128h80v40h32v-40h192v40h32v-40h80Z"/><path fill="currentColor" d="m243.397 313.373l-54.385-54.385l-22.627 22.628l77.012 77.011l125.615-125.615l-22.628-22.627z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Camera(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M471.993 112h-89.2l-16.242-46.75a32.023 32.023 0 0 0-30.229-21.5H175.241a31.991 31.991 0 0 0-30.294 21.691L129.1 112H40a24.027 24.027 0 0 0-24 24v312a24.027 24.027 0 0 0 24 24h431.993a24.027 24.027 0 0 0 24-24V136a24.027 24.027 0 0 0-24-24m-8 328H48.007V144h104.01l23.224-68.25h161.081l23.71 68.25h103.961Z"/><path fill="currentColor" d="M256 168a114 114 0 1 0 114 114a114.13 114.13 0 0 0-114-114m0 196a82 82 0 1 1 82-82a82.093 82.093 0 0 1-82 82"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraControl(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M336 256a80 80 0 1 0-80 80a80.091 80.091 0 0 0 80-80m-128 0a48 48 0 1 1 48 48a48.055 48.055 0 0 1-48-48m-48 141.988L245.307 496h21.386L352 397.988V368H160ZM307.825 400L256 459.544L204.175 400ZM245.307 16L160 114.012V144h192v-29.988L266.693 16Zm-41.132 96L256 52.456L307.825 112ZM16 245.307v21.386L114.013 352H144V160h-29.987Zm96 62.519L52.455 256L112 204.174ZM397.987 160H368v192h29.987L496 266.693v-21.386ZM400 307.826V204.174L459.545 256Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraRoll(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M328 104a32.036 32.036 0 0 0-32-32h-16V48a32.036 32.036 0 0 0-32-32H96a32.036 32.036 0 0 0-32 32v24H48a32.036 32.036 0 0 0-32 32v360a32.036 32.036 0 0 0 32 32h248a32.036 32.036 0 0 0 32-32v-24h168V136H328Zm136 64v240H296v56H48V104h48V48h152v56h48v64Z"/><path fill="currentColor" d="M392 200h40v40h-40zm-72 0h40v40h-40zm-72 0h40v40h-40zm144 136h40v40h-40zm-72 0h40v40h-40zm-72 0h40v40h-40zm-72-136h40v40h-40zm0 136h40v40h-40z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CarAlt(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m475.656 223.142l-90.272-13.908l-55.013-89.87a32.07 32.07 0 0 0-27.548-15.291L72.265 105.91A32.018 32.018 0 0 0 42.2 127.684L16 205.375V384h55.006a68 68 0 0 0 129.988 0h102.012a68 68 0 0 0 129.988 0H496V246.862a23.873 23.873 0 0 0-20.344-23.72M224 136.7l79.078-.63l44.1 72.047L224 208.074Zm-151.479 1.208L192 136.956v71.107l-143.119-.051ZM136 400a36 36 0 1 1 36-36a36.04 36.04 0 0 1-36 36m232 0a36 36 0 1 1 36-36a36.04 36.04 0 0 1-36 36m96-48h-29.082a67.982 67.982 0 0 0-133.836 0h-98.164a67.982 67.982 0 0 0-133.836 0H48V240.026l330.526.529L464 253.724Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretBottom(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256.286 408.357L16.333 138.548V104H496v36.45ZM56.892 136l199.466 224.287L457.042 136Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretLeft(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M400.358 496h-36.45L96 256.286L365.811 16.333h34.547ZM144.071 256.358l224.287 200.684V56.892Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretRight(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M148.092 496h-36.45V16.333h34.547L416 256.286Zm-4.45-439.108v400.15l224.287-200.684Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretTop(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M496 400.357H16.333v-36.449L256.047 96L496 365.81Zm-440.708-32h400.149L255.975 144.07Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cart(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M160 96.039v32h304v63.345l-35.5 112.655H149.932L109.932 16H16v32h66.068l40 288.039h329.9L496 196.306V96.039zm16.984 272.305a64.073 64.073 0 0 0-64 64a64 64 0 0 0 128 0a64.072 64.072 0 0 0-64-64m0 96a32 32 0 1 1 32-32a32.038 32.038 0 0 1-32 32m224-96a64.073 64.073 0 0 0-64 64a64 64 0 0 0 128 0a64.072 64.072 0 0 0-64-64m0 96a32 32 0 1 1 32-32a32.038 32.038 0 0 1-32 32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cash(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M432 64H16v320h416Zm-32 288H48V96h352Z"/><path fill="currentColor" d="M464 144v272H96v32h400V144z"/><path fill="currentColor" d="M224 302.46c39.7 0 72-35.137 72-78.326s-32.3-78.326-72-78.326s-72 35.136-72 78.326s32.3 78.326 72 78.326m0-124.652c22.056 0 40 20.782 40 46.326s-17.944 46.326-40 46.326s-40-20.782-40-46.326s17.944-46.326 40-46.326M80 136h32v176H80zm256 0h32v176h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Casino(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M495.24 267.592L445.066 41.083A32.038 32.038 0 0 0 406.9 16.76L180.393 66.934a32 32 0 0 0-24.322 38.166l21.021 94.9H48a32.036 32.036 0 0 0-32 32v232a32.036 32.036 0 0 0 32 32h232a32.036 32.036 0 0 0 32-32V340.957l158.917-35.2a32.038 32.038 0 0 0 24.323-38.165M280 464H48V232h136.181l22.063 99.606a32.031 32.031 0 0 0 31.18 25.092a32.3 32.3 0 0 0 6.984-.769l35.6-7.886L280.02 464Zm184-189.487l-226.513 50.173l-50.173-226.51L413.824 48l50.193 226.505Z"/><path fill="currentColor" d="M80 264h40v40H80zm0 128h40v40H80zm128 0h40v40h-40zm-64-64h40v40h-40zm81.456-205.433l39.054-8.644l8.644 39.055l-39.054 8.644zm152.672 97.223l39.054-8.65l8.65 39.054l-39.054 8.65zm-76.324-48.649l39.053-8.65l8.65 39.053l-39.052 8.65z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cast(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 360v64h64a64 64 0 0 0-64-64M472 80H40a24.028 24.028 0 0 0-24 24v72h32v-64h416v280H264v32h208a24.028 24.028 0 0 0 24-24V104a24.028 24.028 0 0 0-24-24"/><path fill="currentColor" d="M16 216v32c97.047 0 176 78.953 176 176h32c0-114.691-93.309-208-208-208"/><path fill="currentColor" d="M16 288v32a104.118 104.118 0 0 1 104 104h32c0-74.991-61.009-136-136-136"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cat(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M374.762 186.866a54.1 54.1 0 0 0-51.305-36.706H280V21.552l-18.263 2.609c-41.429 5.918-73.7 26.912-95.907 62.4c-16.011 25.581-23.454 53.8-26.908 74.906c-23.847 18.348-44.593 43.611-61.738 75.2c-14.449 26.618-26.41 57.816-35.552 92.729c-15.447 58.99-17.538 107.921-17.619 109.975L24.005 496H56v-55.636c.4-8.231 10.476-188.35 107.032-256.936l5.66-4.021l.93-6.881C174.437 136.9 191.077 78.058 248 59.971V182.16h75.457a22.12 22.12 0 0 1 21 14.974c12.757 37.656 34.677 84.777 68.839 106.921l-10.274 38.528a62.688 62.688 0 0 1-62.54 46.87c-28.668-.86-58.506 2.88-88.689 11.111L240 403.779V496h32v-67.532a265.353 265.353 0 0 1 67.52-7.03a94.97 94.97 0 0 0 94.418-70.61l17.088-64.081l-12.726-5.454c-23.8-10.2-46.364-43.735-63.538-94.427"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cc(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M472 64H40a24.028 24.028 0 0 0-24 24v336a24.028 24.028 0 0 0 24 24h432a24.028 24.028 0 0 0 24-24V88a24.028 24.028 0 0 0-24-24m-8 352H48V96h416Z"/><path fill="currentColor" d="M184 344a87.108 87.108 0 0 0 54.484-18.891l-19.825-25.119A55.41 55.41 0 0 1 184 312a56 56 0 0 1 0-112a55.41 55.41 0 0 1 34.659 12.01l19.825-25.119A87.108 87.108 0 0 0 184 168a88 88 0 0 0 0 176m163.429 0a87.108 87.108 0 0 0 54.484-18.891l-19.825-25.119A55.414 55.414 0 0 1 347.429 312a56 56 0 0 1 0-112a55.414 55.414 0 0 1 34.659 12.01l19.825-25.119A87.108 87.108 0 0 0 347.429 168a88 88 0 0 0 0 176"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CenterFocus(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 496h120v-32H64a16.019 16.019 0 0 1-16-16V328H16v120a48.054 48.054 0 0 0 48 48M48 64a16.019 16.019 0 0 1 16-16h120V16H64a48.054 48.054 0 0 0-48 48v120h32Zm400-48H328v32h120a16.019 16.019 0 0 1 16 16v120h32V64a48.054 48.054 0 0 0-48-48m16 432a16.019 16.019 0 0 1-16 16H328v32h120a48.054 48.054 0 0 0 48-48V328h-32Zm-64-192c0-79.4-64.6-144-144-144s-144 64.6-144 144s64.6 144 144 144s144-64.6 144-144M256 368a112 112 0 1 1 112-112a112.127 112.127 0 0 1-112 112"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chart(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M124 136H36a20.023 20.023 0 0 0-20 20v320a20.023 20.023 0 0 0 20 20h88a20.023 20.023 0 0 0 20-20V156a20.023 20.023 0 0 0-20-20m-12 328H48V168h64Zm188-224h-88a20.023 20.023 0 0 0-20 20v216a20.023 20.023 0 0 0 20 20h88a20.023 20.023 0 0 0 20-20V260a20.023 20.023 0 0 0-20-20m-12 224h-64V272h64ZM476 16h-88a20.023 20.023 0 0 0-20 20v440a20.023 20.023 0 0 0 20 20h88a20.023 20.023 0 0 0 20-20V36a20.023 20.023 0 0 0-20-20m-12 448h-64V48h64Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartLine(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M376 160v32h65.372L252 381.373l-72-72L76.686 412.686l22.628 22.628L180 354.627l72 72l212-211.999V280h32V160z"/><path fill="currentColor" d="M48 104H16v392h480v-32H48z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartPie(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M105.361 398.32A195.891 195.891 0 0 1 343.42 91.125l23.256-23.255A227.875 227.875 0 0 0 82.733 420.948A228.027 228.027 0 0 0 366.24 452.1l-23.312-23.312c-75.028 43.98-173.271 33.829-237.567-30.468"/><path fill="currentColor" d="M468.916 353.07a243.542 243.542 0 0 0 0-186.459a247.667 247.667 0 0 0-2.747-6.354a242.246 242.246 0 0 0-50.059-72.686L404.8 76.257l-11.317 11.314l-172.27 172.269l172.63 172.631l10.957 10.953l11.31-11.314a242.218 242.218 0 0 0 49.452-71.358a249.292 249.292 0 0 0 3.354-7.682m-64.557-231.12a211.57 211.57 0 0 1 0 275.781L266.468 259.84Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubble(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448.205 392.507c30.519-27.2 47.8-63.455 47.8-101.078c0-39.984-18.718-77.378-52.707-105.3C410.218 158.963 366.432 144 320 144s-90.218 14.963-123.293 42.131C162.718 214.051 144 251.445 144 291.429s18.718 77.378 52.707 105.3c33.075 27.168 76.861 42.13 123.293 42.13c6.187 0 12.412-.273 18.585-.816l10.546 9.141A199.849 199.849 0 0 0 480 496h16v-34.057l-4.686-4.685a199.17 199.17 0 0 1-43.109-64.751M370.089 423l-21.161-18.341l-7.056.865A180.275 180.275 0 0 1 320 406.857c-79.4 0-144-51.781-144-115.428S240.6 176 320 176s144 51.781 144 115.429c0 31.71-15.82 61.314-44.546 83.358l-9.215 7.071l4.252 12.035a231.287 231.287 0 0 0 37.882 67.817A167.839 167.839 0 0 1 370.089 423"/><path fill="currentColor" d="M60.185 317.476a220.491 220.491 0 0 0 34.808-63.023l4.22-11.975l-9.207-7.066C62.918 214.626 48 186.728 48 156.857C48 96.833 109.009 48 184 48c55.168 0 102.767 26.43 124.077 64.3c3.957-.192 7.931-.3 11.923-.3q12.027 0 23.834 1.167c-8.235-21.335-22.537-40.811-42.2-56.961C270.072 30.279 228.3 16 184 16S97.928 30.279 66.364 56.206C33.886 82.885 16 118.63 16 156.857c0 35.8 16.352 70.295 45.25 96.243a188.4 188.4 0 0 1-40.563 60.729L16 318.515V352h16a190.643 190.643 0 0 0 85.231-20.125a157.3 157.3 0 0 1-5.071-33.645a158.729 158.729 0 0 1-51.975 19.246"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Check(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m199.066 456l-7.379-7.514l-3.94-3.9l-86.2-86.2l.053-.055l-83.664-83.666l97.614-97.613l83.565 83.565L398.388 61.344L496 158.958L296.729 358.229l-11.26 11.371ZM146.6 358.183l52.459 52.46l.1-.1l.054.054l52.311-52.311l11.259-11.368l187.963-187.96l-52.358-52.358l-199.273 199.271l-83.565-83.565l-52.359 52.359l83.464 83.463Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckAlt(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M200.359 382.269L61.057 251.673l21.886-23.346l116.698 109.404l229.045-229.044l22.628 22.626z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckCircle(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M426.072 86.928A238.75 238.75 0 0 0 88.428 424.572A238.75 238.75 0 0 0 426.072 86.928M257.25 462.5c-114 0-206.75-92.748-206.75-206.75S143.248 49 257.25 49S464 141.748 464 255.75S371.252 462.5 257.25 462.5"/><path fill="currentColor" d="m221.27 305.808l-73.413-73.412l-22.627 22.627l96.04 96.04l167.5-167.499l-22.628-22.627z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronBottom(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m256.072 424l-11.421-11.313l-85.808-85.809l.053-.054L16 183.928l97.122-97.122l142.9 142.9l142.9-142.9l97.122 97.122l-142.801 142.794l-11.361 11.469Zm-.107-45.254l.054.053l51.835-51.835l11.346-11.453l131.583-131.583l-51.868-51.868l-142.9 142.9l-142.9-142.9l-51.861 51.868l142.9 142.9l-.053.054Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronCircleDownAlt(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 496A240 240 0 0 1 86.294 86.294a240 240 0 0 1 339.412 339.412A238.43 238.43 0 0 1 256 496m0-448C141.309 48 48 141.309 48 256s93.309 208 208 208s208-93.309 208-208S370.691 48 256 48"/><path fill="currentColor" d="M256 342.627L132.687 219.313l22.626-22.626L256 297.372l100.687-100.685l22.626 22.626z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronCircleLeftAlt(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 496A240 240 0 0 1 86.294 86.294a240 240 0 0 1 339.412 339.412A238.43 238.43 0 0 1 256 496m0-448C141.309 48 48 141.309 48 256s93.309 208 208 208s208-93.309 208-208S370.691 48 256 48"/><path fill="currentColor" d="M284.687 379.313L161.373 256l123.314-123.313l22.626 22.626L206.628 256l100.685 100.687z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronCircleRightAlt(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 496A240 240 0 0 1 86.294 86.294a240 240 0 0 1 339.412 339.412A238.43 238.43 0 0 1 256 496m0-448C141.309 48 48 141.309 48 256s93.309 208 208 208s208-93.309 208-208S370.691 48 256 48"/><path fill="currentColor" d="m227.313 379.313l-22.626-22.626L305.372 256L204.687 155.313l22.626-22.626L350.627 256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronCircleUpAlt(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 496A240 240 0 0 1 86.294 86.294a240 240 0 0 1 339.412 339.412A238.432 238.432 0 0 1 256 496m0-448C141.309 48 48 141.309 48 256s93.309 208 208 208s208-93.309 208-208S370.691 48 256 48"/><path fill="currentColor" d="M356.686 315.313L256 214.628L155.314 315.313l-22.628-22.626L256 169.373l123.314 123.314z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleDown(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m252.482 345.034l-109.02-109.021L64 315.475l109.087 109.088l-.041.041l68.082 68.082l11.354 11.273l.041.041l68.15-68.148l11.206-11.289l109.087-109.088l-79.466-79.462Zm45.523 68.149l-11.17 11.252l-34.353 34.351l-.039-.038l-34.276-34.277l.041-.042l-108.954-108.954l34.208-34.207l109.02 109.021L361.5 281.268l34.208 34.207Z"/><path fill="currentColor" d="m173.046 203.4l68.082 68.082l11.354 11.272l.041.042l68.15-68.149l11.206-11.289L440.966 94.267L361.5 14.8L252.482 123.825L143.462 14.8L64 94.267l109.087 109.087ZM109.254 94.267l34.208-34.208l109.02 109.021L361.5 60.059l34.208 34.208l-97.707 97.707l-11.182 11.264l-34.341 34.34l-.039-.039l-34.276-34.277l.041-.041Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleLeft(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M276.055 143.463L196.592 64L87.571 173.021l-.041-.041l-68.149 68.148l-4.3 4.328l-7.085 7l68.217 68.216l11.342 11.26l109.037 109.035l79.463-79.467l-109.021-109.016Zm-79.463 252.249l-97.707-97.707l-11.355-11.273l-34.182-34.181l.041-.041l-.067-.068l34.208-34.207l.041.041l109.021-109.021l34.208 34.208l-109.021 109.021L230.8 361.5Z"/><path fill="currentColor" d="M497.263 143.463L417.8 64L308.713 173.088l-.042-.041l-68.081 68.081l-11.19 11.278l-.119.119l68.148 68.148l11.288 11.2L417.8 440.967l79.463-79.467l-109.021-109.016ZM452.008 361.5L417.8 395.712l-97.707-97.707l-11.28-11.2l-34.325-34.324l34.316-34.316l.042.041L417.8 109.255l34.207 34.208l-109.02 109.021Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleRight(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m503.933 255.933l.041-.041l-68.081-68.082l-6.221-6.175l-5.122-5.085L315.516 67.516l-79.462 79.463L345.075 256L236.054 365.021l79.462 79.463l109.021-109.021l.041.041l.067-.067l68.081-68.082L504 256Zm-188.417 143.3l-34.207-34.208L390.33 256L281.309 146.979l34.207-34.208l97.707 97.707l11.355 11.273L458.827 256l-34.249 34.249l-.041-.041l-.067.067Z"/><path fill="currentColor" d="m282.792 256l-.067-.067l.041-.041l-68.082-68.082L203.4 176.6L94.308 67.516l-79.463 79.463L123.866 256L14.845 365.021l79.463 79.463L203.4 335.4l.041.041l68.082-68.082ZM203.3 290.316l-.041-.041L94.308 399.229L60.1 365.021L169.121 256L60.1 146.979l34.208-34.208l97.707 97.707l11.279 11.2L237.619 256Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleUp(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m259.516 173.771l109.021 109.021L448 203.329L338.979 94.308l.041-.041l-68.148-68.149l-4.328-4.3l-7-7.085l-.067.067l-68.151 68.153l-11.338 11.421L71.033 203.329l79.467 79.463ZM214 105.622L225.334 94.2l34.115-34.115l.041.041l.067-.067l34.208 34.208l-.041.041l109.021 109.021l-34.208 34.208l-109.021-109.021L150.5 237.537l-34.207-34.208Z"/><path fill="currentColor" d="m338.953 315.408l-68.081-68.081l-11.289-11.206l-.108-.108l-68.149 68.148l-11.272 11.356L71.033 424.538L150.5 504l109.016-109.021L368.537 504L448 424.538L338.912 315.45Zm63.792 109.13l-34.208 34.207l-109.021-109.021L150.5 458.745l-34.207-34.207L214 326.831l11.256-11.34l34.265-34.265l34.316 34.316l-.041.041Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleUpAlt(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M20.095 20.405L12 12.31l-8.095 8.095l-1.061-1.061l9.155-9.155l9.155 9.155l-1.061 1.061z" fill="currentColor"/><path d="M20.095 12.905L12 4.81l-8.095 8.095l-1.061-1.061l9.155-9.155l9.155 9.155l-1.061 1.061z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeft(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m327.086 496.89l-142.6-142.6l-11.258-11.15l-85.6-85.6l.054-.054l11.259-11.367l85.5-85.5l.054.054l142.6-142.595L424 114.989L281.506 257.483L424 399.978ZM184.64 309.3l11.266 11.159l131.18 131.181l51.658-51.658l-142.493-142.499l142.493-142.494l-51.658-51.658l-142.392 142.394l-.054-.054l-51.813 51.812Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRight(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m183.505 496l-97.268-97.27l143.175-143.174L86.237 112.38l97.268-97.27l143.227 143.228l11.316 11.209l85.9 85.9l.051.05l-11.311 11.419l-85.9 85.9l-.055-.054Zm-52.013-97.27l52.013 52.014L326.629 307.62l.055.054l52.116-52.118l-52.127-52.128l-11.308-11.2l-131.86-131.862l-52.013 52.014l143.175 143.176Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronTop(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M398.986 424.715L256 281.73L113.014 424.715l-97.17-97.169L158.8 184.59l11.29-11.4L256 87.285l5.481 5.531l5.89 5.834l85.907 85.908l-.054.054l142.932 142.934ZM61.1 327.546l51.915 51.915L256 236.474l142.986 142.987l51.914-51.915l-143.037-143.038l.054-.053l-51.812-51.813l-.051.051l-.1-.106l-51.866 51.869l-11.312 11.418Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Child(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M425.39 200.035A184.3 184.3 0 0 0 290.812 91.289l26.756-42.809l-27.136-16.96l-35.305 56.488A184.046 184.046 0 0 0 86.61 200.035a71.978 71.978 0 0 0 0 143.93a184.071 184.071 0 0 0 338.78 0a71.978 71.978 0 0 0 0-143.93m27.152 99.975a39.77 39.77 0 0 1-27.76 11.961l-20.725.394l-8.113 19.074a152.066 152.066 0 0 1-279.887 0l-8.114-19.074l-20.725-.394a39.978 39.978 0 0 1 0-79.942l20.725-.394l8.114-19.074a152.067 152.067 0 0 1 279.887 0l8.113 19.074l20.725.394a39.974 39.974 0 0 1 27.76 67.981"/><path fill="currentColor" d="M168 232h40v40h-40zm136 0h40v40h-40zm-48 152a80 80 0 0 0 80-80H176a80 80 0 0 0 80 80"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChildFriendly(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M445.057 345.134L464 274.1V232c-8.136-93.993-87.933-168-184-168h-32v168H132.158l-17.844-78.768A32.155 32.155 0 0 0 83.038 128H16v32h67.038l40.475 178.67A80 80 0 1 0 224 416q0-4.05-.4-8h104.8q-.395 3.948-.4 8a80 80 0 1 0 117.057-70.866M280 96c78.411 0 143.145 59.678 151.164 136H280ZM144 464a48 48 0 1 1 48-48a48.055 48.055 0 0 1-48 48m194.763-88H213.237a80.166 80.166 0 0 0-57.316-39.108L139.408 264H432v5.9l-17.7 66.368a80.592 80.592 0 0 0-6.3-.271A80.026 80.026 0 0 0 338.763 376M408 464a48 48 0 1 1 48-48a48.055 48.055 0 0 1-48 48"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Circle(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256.6 496A239.364 239.364 0 0 0 425.856 87.379A239.364 239.364 0 0 0 87.344 425.892A237.8 237.8 0 0 0 256.6 496m0-446.729c114.341 0 207.365 93.023 207.365 207.364S370.941 464 256.6 464S49.236 370.977 49.236 256.635S142.259 49.271 256.6 49.271"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClearAll(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M160 112h336v32H160zM88 240h336v32H88zM16 368h336v32H16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clipboard(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M432 56h-56v32h48v376H88V88h48V56H80a24.028 24.028 0 0 0-24 24v392a24.028 24.028 0 0 0 24 24h352a24.028 24.028 0 0 0 24-24V80a24.028 24.028 0 0 0-24-24"/><path fill="currentColor" d="M192 140h128a24.028 24.028 0 0 0 24-24V16H168v100a24.028 24.028 0 0 0 24 24m8-92h112v60H200Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clock(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M271.514 95.5h-32v178.111l115.613 54.948l13.737-28.902l-97.35-46.268z"/><path fill="currentColor" d="M256 16C123.452 16 16 123.452 16 256s107.452 240 240 240s240-107.452 240-240S388.548 16 256 16m0 448c-114.875 0-208-93.125-208-208S141.125 48 256 48s208 93.125 208 208s-93.125 208-208 208"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clone(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M472 16H160a24.027 24.027 0 0 0-24 24v312a24.027 24.027 0 0 0 24 24h312a24.027 24.027 0 0 0 24-24V40a24.027 24.027 0 0 0-24-24m-8 328H168V48h296Z"/><path fill="currentColor" d="M344 464H48V168h56v-32H40a24.027 24.027 0 0 0-24 24v312a24.027 24.027 0 0 0 24 24h312a24.027 24.027 0 0 0 24-24v-64h-32Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClosedCaptioning(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M472 64H40a24.028 24.028 0 0 0-24 24v336a24.028 24.028 0 0 0 24 24h432a24.028 24.028 0 0 0 24-24V88a24.028 24.028 0 0 0-24-24m-8 352H48V96h416Z"/><path fill="currentColor" d="M184 344a87.108 87.108 0 0 0 54.484-18.891l-19.825-25.119A55.41 55.41 0 0 1 184 312a56 56 0 0 1 0-112a55.41 55.41 0 0 1 34.659 12.01l19.825-25.119A87.108 87.108 0 0 0 184 168a88 88 0 0 0 0 176m163.429 0a87.108 87.108 0 0 0 54.484-18.891l-19.825-25.119A55.414 55.414 0 0 1 347.429 312a56 56 0 0 1 0-112a55.414 55.414 0 0 1 34.659 12.01l19.825-25.119A87.108 87.108 0 0 0 347.429 168a88 88 0 0 0 0 176"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloud(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M382.76 432H136c-30.732 0-61.371-12.725-84.061-34.912c-23.221-22.707-36.009-52.35-36.009-83.469A109.4 109.4 0 0 1 49.136 235.2a122.281 122.281 0 0 1 62.794-32.707V200c0-79.4 64.6-144 144-144s144 64.6 144 144v1.453c55.716 7.939 96 53.729 96 112.166c0 31.27-11.375 60.72-32.031 82.927A109.747 109.747 0 0 1 382.76 432M255.93 88a112.127 112.127 0 0 0-112 112v31.405l-14.864 1.059c-45.5 3.239-81.136 38.887-81.136 81.155C47.93 359.635 89.084 400 136 400h246.76c45.515 0 81.17-37.943 81.17-86.381c0-46.566-33.731-80.791-80.2-81.379l-15.8-.2V200a112.127 112.127 0 0 0-112-112"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudDownload(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M272 434.744V209.176h-32v225.568l-51.882-51.882l-22.628 22.627L256 496l90.51-90.511l-22.628-22.627z"/><path fill="currentColor" d="M400 161.176c0-79.4-64.6-144-144-144s-144 64.6-144 144a96 96 0 0 0 0 192h80v-32h-80a64 64 0 0 1 0-128h32v-32a112 112 0 0 1 224 0v32h32a64 64 0 0 1 0 128h-80v32h80a96 96 0 0 0 0-192"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudUpload(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m346.231 284.746l-90.192-90.192l-90.192 90.192l22.627 22.627l51.565-51.565V496h32V255.808l51.565 51.565z"/><path fill="currentColor" d="M400 161.453V160c0-79.4-64.6-144-144-144S112 80.6 112 160v2.491A122.285 122.285 0 0 0 49.206 195.2A109.4 109.4 0 0 0 16 273.619c0 31.119 12.788 60.762 36.01 83.469C74.7 379.275 105.338 392 136.07 392H200v-32h-63.93C89.154 360 48 319.635 48 273.619c0-42.268 35.64-77.916 81.137-81.155L144 191.405V160a112 112 0 0 1 224 0v32.04l15.8.2c46.472.588 80.2 34.813 80.2 81.379C464 322.057 428.346 360 382.83 360H312v32h70.83a109.749 109.749 0 0 0 81.14-35.454c20.655-22.207 32.03-51.657 32.03-82.927c0-58.437-40.284-104.227-96-112.166"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloudy(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M399.667 264.875v-3.813c0-79.4-64.6-144-144-144c-2.2 0-4.391.057-6.569.156a116.689 116.689 0 1 0-136.783 130.226a145.275 145.275 0 0 0-.648 13.618v3.813a116.633 116.633 0 0 0 20.62 231.425h246.759a116.633 116.633 0 0 0 20.621-231.427ZM48.75 132.688a84.677 84.677 0 0 1 168.705-10.47a144.606 144.606 0 0 0-98.59 93.876a84.807 84.807 0 0 1-70.115-83.406M379.046 464.3H132.287a84.619 84.619 0 0 1-3.9-169.148l15.277-.69v-33.4a112 112 0 1 1 224 0v33.4l15.277.69a84.619 84.619 0 0 1-3.9 169.148Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Code(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m388.632 393.82l107.191-137.88l-107.139-137.762l-25.26 19.644l91.864 118.122l-91.92 118.236zm-240.053-19.639L56.712 255.999l91.917-118.176l-25.258-19.646L16.177 255.993l107.137 137.826zM330.529 16h-32.97L178.441 496h32.971z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Coffee(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M191.5 376h97A111.627 111.627 0 0 0 400 264.5V248h24a72 72 0 0 0 0-144H80v160.5A111.627 111.627 0 0 0 191.5 376M400 136h24a40 40 0 0 1 0 80h-24Zm-288 0h256v128.5a79.589 79.589 0 0 1-79.5 79.5h-97a79.589 79.589 0 0 1-79.5-79.5ZM16 416h480v32H16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cog(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M245.151 168a88 88 0 1 0 88 88a88.1 88.1 0 0 0-88-88m0 144a56 56 0 1 1 56-56a56.063 56.063 0 0 1-56 56"/><path fill="currentColor" d="m464.7 322.319l-31.77-26.153a193.081 193.081 0 0 0 0-80.332l31.77-26.153a19.941 19.941 0 0 0 4.606-25.439l-32.612-56.483a19.936 19.936 0 0 0-24.337-8.73l-38.561 14.447a192.038 192.038 0 0 0-69.54-40.192l-6.766-40.571A19.936 19.936 0 0 0 277.762 16H212.54a19.937 19.937 0 0 0-19.728 16.712l-6.762 40.572a192.03 192.03 0 0 0-69.54 40.192L77.945 99.027a19.937 19.937 0 0 0-24.334 8.731L21 164.245a19.94 19.94 0 0 0 4.61 25.438l31.767 26.151a193.081 193.081 0 0 0 0 80.332l-31.77 26.153A19.942 19.942 0 0 0 21 347.758l32.612 56.483a19.937 19.937 0 0 0 24.337 8.73l38.562-14.447a192.03 192.03 0 0 0 69.54 40.192l6.762 40.571A19.937 19.937 0 0 0 212.54 496h65.222a19.936 19.936 0 0 0 19.728-16.712l6.763-40.572a192.038 192.038 0 0 0 69.54-40.192l38.564 14.449a19.938 19.938 0 0 0 24.334-8.731l32.609-56.487a19.939 19.939 0 0 0-4.6-25.436m-50.636 57.12l-48.109-18.024l-7.285 7.334a159.955 159.955 0 0 1-72.625 41.973l-10 2.636L267.6 464h-44.89l-8.442-50.642l-10-2.636a159.955 159.955 0 0 1-72.625-41.973l-7.285-7.334l-48.117 18.024L53.8 340.562l39.629-32.624l-2.7-9.973a160.9 160.9 0 0 1 0-83.93l2.7-9.972L53.8 171.439l22.446-38.878l48.109 18.024l7.285-7.334a159.955 159.955 0 0 1 72.625-41.973l10-2.636L222.706 48H267.6l8.442 50.642l10 2.636a159.955 159.955 0 0 1 72.625 41.973l7.285 7.334l48.109-18.024l22.447 38.877l-39.629 32.625l2.7 9.972a160.9 160.9 0 0 1 0 83.93l-2.7 9.973l39.629 32.623Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ColorBorder(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M143.4 331.909a17.505 17.505 0 0 0 3.285-.311l89.867-21.281l190.2-190.2a62.922 62.922 0 1 0-88.986-88.985l-190.2 190.2l-21.134 89.185l-.145.674a17.435 17.435 0 0 0 17.113 20.718m33.21-94.369L360.4 53.755a30.922 30.922 0 1 1 43.731 43.731L220.342 281.272l-57.314 13.582ZM472 376H40a24.028 24.028 0 0 0-24 24v72a24.028 24.028 0 0 0 24 24h432a24.028 24.028 0 0 0 24-24v-72a24.028 24.028 0 0 0-24-24m-8 88H48v-56h416Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ColorFill(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M40 496h432a24.028 24.028 0 0 0 24-24v-72a24.028 24.028 0 0 0-24-24H40a24.028 24.028 0 0 0-24 24v72a24.028 24.028 0 0 0 24 24m8-88h416v56H48Zm-2.99-168.511a50.2 50.2 0 0 0 36.2 61.07l158.5 40.524a50.174 50.174 0 0 0 61.07-36.2l52.962-207.138L97.969 32.352l-12.01 46.974l-58-14.827l-7.926 31l58 14.828Zm76.035-168.208l193.767 49.54L294.568 200H88.135ZM286.387 232l-16.607 64.956a18.221 18.221 0 0 1-22.141 13.125l-158.5-40.525a18.22 18.22 0 0 1-13.124-22.14L80 231.82v.18ZM420 352c37.5 0 68-32.3 68-72c0-16.181-9.341-38.514-28.559-68.277c-13.285-20.576-26.39-37.021-26.941-37.711L420 158.37l-12.5 15.642c-.551.69-13.656 17.135-26.941 37.711C361.341 241.486 352 263.819 352 280c0 39.7 30.505 72 68 72m0-141.441c17.658 24.641 36 55.426 36 69.441c0 22.056-16.149 40-36 40s-36-17.944-36-40c0-14.015 18.343-44.8 36-69.441"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ColorPalette(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M425.514 82.055C380.778 39.458 320.673 16 256.271 16c-60.023 0-119.856 20.073-164.156 55.071C43.032 109.85 16 164.161 16 224c0 60.1 15.531 98.87 48.876 122.019c28 19.438 68.412 27.731 135.124 27.731h29.75A26.28 26.28 0 0 1 256 400v47.984a32 32 0 0 0 32 32h.032l90.755-.088a32.094 32.094 0 0 0 19.686-6.8c9.725-7.622 34.727-29.4 56.8-66.9C482.3 360.262 496 307.037 496 248c0-63.732-25.032-122.666-70.486-165.945m2.173 307.915c-19.3 32.792-40.663 51.447-48.932 57.926l-90.755.088V400a58.316 58.316 0 0 0-58.25-58.25H200c-59.69 0-94.644-6.585-116.876-22.019C58.833 302.869 48 273.344 48 224C48 113.833 153.9 48 256.271 48C372.755 48 464 135.851 464 248c0 53.253-12.218 101.019-36.313 141.97"/><path fill="currentColor" d="M128 144a56 56 0 1 0 56 56a56.064 56.064 0 0 0-56-56m0 80a24 24 0 1 1 24-24a24.027 24.027 0 0 1-24 24M240 72a56 56 0 1 0 56 56a56.064 56.064 0 0 0-56-56m0 80a24 24 0 1 1 24-24a24.027 24.027 0 0 1-24 24m120-24a56 56 0 1 0 56 56a56.064 56.064 0 0 0-56-56m0 80a24 24 0 1 1 24-24a24.027 24.027 0 0 1-24 24m16 56a56 56 0 1 0 56 56a56.064 56.064 0 0 0-56-56m0 80a24 24 0 1 1 24-24a24.027 24.027 0 0 1-24 24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Columns(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M472 48H40a24.028 24.028 0 0 0-24 24v376a24.028 24.028 0 0 0 24 24h432a24.028 24.028 0 0 0 24-24V72a24.028 24.028 0 0 0-24-24m-8 32v64H48V80ZM48 176h192v264H48Zm224 264V176h192v264Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Command(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M416 176a80 80 0 1 0-80-80v48H176V96a80 80 0 1 0-80 80h48v160H96a80 80 0 1 0 80 80v-48h160v48a80 80 0 1 0 80-80h-48V176Zm-48-80a48 48 0 1 1 48 48h-48ZM144 416a48 48 0 1 1-48-48h48Zm0-272H96a48 48 0 1 1 48-48Zm192 192H176V176h160Zm80 32a48 48 0 1 1-48 48v-48Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentBubble(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M496 496h-16a273.39 273.39 0 0 1-179.025-66.782l-16.827-14.584A291.407 291.407 0 0 1 256 416c-63.527 0-123.385-20.431-168.548-57.529C41.375 320.623 16 270.025 16 216S41.375 111.377 87.452 73.529C132.615 36.431 192.473 16 256 16s123.385 20.431 168.548 57.529C470.625 111.377 496 161.975 496 216a171.161 171.161 0 0 1-21.077 82.151a201.505 201.505 0 0 1-47.065 57.537a285.22 285.22 0 0 0 63.455 97l4.687 4.685ZM294.456 381.222l27.477 23.814a241.379 241.379 0 0 0 135 57.86a317.5 317.5 0 0 1-62.617-105.583l-4.395-12.463l9.209-7.068C440.963 305.678 464 262.429 464 216c0-92.636-93.309-168-208-168S48 123.364 48 216s93.309 168 208 168a259.114 259.114 0 0 0 31.4-1.913Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentSquare(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M496 496h-47.229l-69.522-128H40a24.028 24.028 0 0 1-24-24V40a24.028 24.028 0 0 1 24-24h432a24.028 24.028 0 0 1 24 24ZM48 336h350.284L464 456.993V48H48Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compass(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m209.686 304.963l119.051 59.424a24 24 0 0 0 32.182-32.21l-59.5-118.948l-119.086-59.645a24 24 0 0 0-32.216 32.189Zm67.88-67.892l44.006 87.975l-88.037-43.946l-44.056-88.149Z"/><path fill="currentColor" d="M256 496c132.548 0 240-107.452 240-240S388.548 16 256 16S16 123.452 16 256s107.452 240 240 240M48.353 244C54.269 140.018 136.553 56.476 240 48.606V115h32V48.606C375.447 56.476 457.731 140.018 463.647 244h-63.676v32h63.068C453.474 376.238 372.78 455.727 272 463.394V403h-32v60.394C139.22 455.727 58.526 376.238 48.961 276h63.01v-32Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compress(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M144 144H16v32h160V16h-32zm224 0V16h-32v160h160v-32zm-32 352h32V368h128v-32H336zM16 368h128v128h32V336H16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Contact(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M472 48H40a24.028 24.028 0 0 0-24 24v368a24.028 24.028 0 0 0 24 24h88v-58.822a20.01 20.01 0 0 1 10.284-17.478l91.979-51.123L200 260.919V200a56 56 0 0 1 112 0v60.919l-30.263 75.655l91.979 51.126A20.011 20.011 0 0 1 384 405.178V464h88a24.028 24.028 0 0 0 24-24V72a24.028 24.028 0 0 0-24-24m-8 384h-48v-26.822a52.027 52.027 0 0 0-26.738-45.451L321.915 322.3L344 267.081V200a88 88 0 0 0-176 0v67.081l22.085 55.219l-67.347 37.432A52.027 52.027 0 0 0 96 405.178V432H48V80h416Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Contrast(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 16C123.452 16 16 123.452 16 256s107.452 240 240 240s240-107.452 240-240S388.548 16 256 16m-22 446.849a208.346 208.346 0 0 1-169.667-125.9c-.364-.859-.706-1.724-1.057-2.587L234 429.939Zm0-69.582L50.889 290.76A209.848 209.848 0 0 1 48 256q0-9.912.922-19.67L234 339.939Zm0-90L54.819 202.96a206.385 206.385 0 0 1 9.514-27.913Q67.1 168.5 70.3 162.191L234 253.934Zm0-86.015L86.914 134.819a209.42 209.42 0 0 1 22.008-25.9q3.72-3.72 7.6-7.228L234 166.027Zm0-87.708l-89.648-49.093A206.951 206.951 0 0 1 234 49.151ZM464 256a207.775 207.775 0 0 1-198 207.761V48.239A207.791 207.791 0 0 1 464 256"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Control(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M336 256a80 80 0 1 0-80 80a80.091 80.091 0 0 0 80-80m-128 0a48 48 0 1 1 48 48a48.055 48.055 0 0 1-48-48m-48 141.988L245.307 496h21.386L352 397.988V368H160ZM307.825 400L256 459.544L204.175 400ZM245.307 16L160 114.012V144h192v-29.988L266.693 16Zm-41.132 96L256 52.456L307.825 112ZM16 245.307v21.386L114.013 352H144V160h-29.987Zm96 62.519L52.455 256L112 204.174ZM397.987 160H368v192h29.987L496 266.693v-21.386ZM400 307.826V204.174L459.545 256Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copy(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M408 432h-32v32H112V136h32v-32H80v392h328z"/><path fill="currentColor" d="M176 16v384h320V153.373L358.627 16Zm288 352H208V48h104v152h152Zm0-200H344V48h1.372L464 166.627Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Couch(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 242.025V152a48.055 48.055 0 0 0-48-48H112a48.055 48.055 0 0 0-48 48v90.025A64.115 64.115 0 0 0 16 304v112a32.036 32.036 0 0 0 32 32h16v48h32v-48h320v48h32v-48h16a32.036 32.036 0 0 0 32-32V304a64.115 64.115 0 0 0-48-61.975M112 416H48V304a32 32 0 0 1 64 0Zm256 0H144v-96h224Zm2.025-128h-228.05A64.243 64.243 0 0 0 96 242.025V152a16.019 16.019 0 0 1 16-16h288a16.019 16.019 0 0 1 16 16v90.025A64.243 64.243 0 0 0 370.025 288M464 416h-64V304a32 32 0 0 1 64 0l.02 112Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCard(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M472 72H40a24.028 24.028 0 0 0-24 24v320a24.028 24.028 0 0 0 24 24h432a24.028 24.028 0 0 0 24-24V96a24.028 24.028 0 0 0-24-24m-8 32v64H48v-64ZM48 408V232h416v176Z"/><path fill="currentColor" d="M88 312h64v32H88zm96 0h64v32h-64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crop(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M433.372 56H88V16H56v40H16v32h40v368h368v40h32v-40h40v-32h-40V78.628l40-40V16h-22.628Zm-32 32L88 401.372V88ZM424 424H110.628L424 110.628Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CropRotate(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M49.386 256c0-114.184 92.8-207.1 206.935-207.294v41.759L368 15.965H256.321L256 16c-63.783.094-123.414 25.679-168.525 70.79A237.732 237.732 0 0 0 17.386 256Zm206.601 207.367V432l-111.679 64h111.679c132.4-.36 240-108.214 240-239.965h-32c-.005 114.106-93.249 206.972-208 207.332"/><path fill="currentColor" d="M337.372 152H184v-32h-32v32h-32v32h32v176h168v32h32v-32h32v-32h-32V222.628l-32 32V328H201.373L384 145.372V128h-22.628ZM184 305.372V184h121.372Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cursor(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M446.754 47.9a20.075 20.075 0 0 0-21.307-2.745L32 229.835v37l165.349 66.139L303.317 496h37L453.281 68.369a20.072 20.072 0 0 0-6.527-20.469m-129.63 410.624l-98.473-151.5l-148.9-59.561L415.779 85.044Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CursorMove(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m130.412 323.98l-51.883-51.882H240v161.47l-51.882-51.881l-22.628 22.626l90.51 90.51l90.51-90.51l-22.628-22.626L272 433.568v-161.47h160.667l-51.883 51.882l22.628 22.627l90.509-90.509l-90.509-90.51l-22.628 22.627l51.883 51.883H272V79.432l51.882 51.881l22.628-22.626L256 18.177l-90.51 90.51l22.628 22.626L240 79.432v160.666H78.529l51.883-51.883l-22.628-22.627l-90.51 90.51l90.511 90.51z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cut(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M104 320a88.067 88.067 0 1 0 75.607 43.02L260 282.627L425.373 448H448v-22.627L282.627 260L448 94.627V72h-22.627L260 237.373l-83.523-83.523a88.088 88.088 0 1 0-22.627 22.627L237.373 260l-78.919 78.919A87.57 87.57 0 0 0 104 320m0-160a56 56 0 1 1 56-56a56.063 56.063 0 0 1-56 56m0 304a56 56 0 1 1 56-56a56.063 56.063 0 0 1-56 56"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DataTransferDown(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 464h480v32H16zm379.313-164.687l-22.626-22.626L280 369.373V16h-32v353.373l-92.687-92.686l-22.626 22.626L264 430.627z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DataTransferUp(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 16h480v32H16zm100.687 196.687l22.626 22.626L232 142.627V496h32V142.627l92.687 92.686l22.626-22.626L248 81.373z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Deaf(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256.16 72c65.3.46 118.745 53.964 119.131 119.269a120.126 120.126 0 0 1-33.161 83.545l22.631 22.631a152.141 152.141 0 0 0 42.53-106.365C406.8 108.36 339.105 40.586 256.385 40a150.765 150.765 0 0 0-106.536 42.536l22.629 22.628A119.24 119.24 0 0 1 256.16 72m103.133 265.235l-32-32l-83.993-83.994l-42.29-42.29l-24.843-24.843L38.058 16H15.293v22.489l98.294 98.294A151.914 151.914 0 0 0 103.293 192v32h32v-32A120.77 120.77 0 0 1 139 162.2l28.324 28.324a89.713 89.713 0 0 0-.034 2.381v85.729l5.089 5.082A118.419 118.419 0 0 1 207.293 368h32a150.125 150.125 0 0 0-40-102.453v-43.058l128 128v26.178A63.4 63.4 0 0 1 263.96 440h-24.667v32h24.667a95.452 95.452 0 0 0 95.161-89.684L472.805 496h22.488v-22.766Z"/><path fill="currentColor" d="M254.628 104a87.158 87.158 0 0 0-59.328 23.991l22.62 22.62A55.232 55.232 0 0 1 254.865 136h.428a56 56 0 0 1 56 56h32a88 88 0 0 0-88.665-88"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Delete(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M227.313 363.313L312 278.627l84.687 84.686l22.626-22.626L334.627 256l84.686-84.687l-22.626-22.626L312 233.373l-84.687-84.686l-22.626 22.626L289.373 256l-84.686 84.687z"/><path fill="currentColor" d="M472 64H194.644a24.091 24.091 0 0 0-17.42 7.492L16 241.623v28.754l161.224 170.131a24.091 24.091 0 0 0 17.42 7.492H472a24.028 24.028 0 0 0 24-24V88a24.028 24.028 0 0 0-24-24m-8 352H198.084L48 257.623v-3.246L198.084 96H464Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Description(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M334.627 16H48v480h424V153.373ZM440 166.627V168H320V48h1.373ZM80 464V48h208v152h152v264Z"/><path fill="currentColor" d="M136 296h224v32H136zm0 80h224v32H136z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Devices(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M472 232h-48V120a24.028 24.028 0 0 0-24-24H40a24.028 24.028 0 0 0-24 24v246a24.028 24.028 0 0 0 24 24h172v50h-60v32h152v-32h-60v-50h92v58a24.027 24.027 0 0 0 24 24h112a24.027 24.027 0 0 0 24-24V256a24.027 24.027 0 0 0-24-24m-136 24v102H48V128h344v104h-32a24.027 24.027 0 0 0-24 24m128 184h-96V264h96Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dialpad(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M121 16a48 48 0 1 0 48 48a48.054 48.054 0 0 0-48-48m0 64a16 16 0 1 1 16-16a16.019 16.019 0 0 1-16 16m128-64a48 48 0 1 0 48 48a48.054 48.054 0 0 0-48-48m0 64a16 16 0 1 1 16-16a16.019 16.019 0 0 1-16 16m128 32a48 48 0 1 0-48-48a48.054 48.054 0 0 0 48 48m0-64a16 16 0 1 1-16 16a16.019 16.019 0 0 1 16-16m-256 96a48 48 0 1 0 48 48a48.054 48.054 0 0 0-48-48m0 64a16 16 0 1 1 16-16a16.019 16.019 0 0 1-16 16m128-64a48 48 0 1 0 48 48a48.054 48.054 0 0 0-48-48m0 64a16 16 0 1 1 16-16a16.019 16.019 0 0 1-16 16m128-64a48 48 0 1 0 48 48a48.054 48.054 0 0 0-48-48m0 64a16 16 0 1 1 16-16a16.019 16.019 0 0 1-16 16m-256 64a48 48 0 1 0 48 48a48.054 48.054 0 0 0-48-48m0 64a16 16 0 1 1 16-16a16.019 16.019 0 0 1-16 16m128-64a48 48 0 1 0 48 48a48.054 48.054 0 0 0-48-48m0 64a16 16 0 1 1 16-16a16.019 16.019 0 0 1-16 16m128-64a48 48 0 1 0 48 48a48.054 48.054 0 0 0-48-48m0 64a16 16 0 1 1 16-16a16.019 16.019 0 0 1-16 16m-128 64a48 48 0 1 0 48 48a48.054 48.054 0 0 0-48-48m0 64a16 16 0 1 1 16-16a16.019 16.019 0 0 1-16 16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Diamond(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M408.563 48H103.438L16 179.156v24.786L199.421 480h113.158L496 203.942v-24.786Zm-17.125 32l63.407 95.111H347.739L317.808 80Zm-163.7 0h56.524l29.93 95.111H197.808Zm-107.175 0h73.629l-29.931 95.111H57.155Zm96.016 368L56.525 207.111h106.793L219.813 448Zm36.1 0l-56.492-240.889h119.626L259.318 448Zm42.739 0h-3.234l56.5-240.889h106.791Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dinner(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 344.063c0-109.308-84.755-199.193-192-207.39V80h-32v56.673c-107.245 8.2-192 98.082-192 207.39v33.107h416Zm-32 1.107H80v-1.107c0-97.046 78.953-176 176-176s176 78.953 176 176ZM16 416h480v32H16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Disabled(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="221.912" cy="66.088" r="34.088" fill="currentColor"/><path fill="currentColor" d="m460.12 360.478l-47.943 11.985L393 282.971A24.126 24.126 0 0 0 369.533 264h-88.705l-6.462-56H384v-32H270.674l-4.134-35.826a24 24 0 0 0-26.593-21.091l-39.736 4.585L220.1 296h142.97l24.758 115.537l80.057-20.015Z"/><path fill="currentColor" d="M224 448a120 120 0 0 1-45.248-231.135l-3.779-32.75C115.143 204.558 72 261.334 72 328c0 83.813 68.187 152 152 152a152.06 152.06 0 0 0 130.044-73.378L344 360c-16 48-61.4 88-120 88"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dog(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m393.3 161.33l-58.768-84.892a48.09 48.09 0 0 0-38.775-20.673l-111.527-1.6H184c-57.579 0-101.757 9.631-130.21 56.634C27.3 154.551 16 229.08 16 360v16h36.557L29.024 496h32.61l23.533-120H96a99.521 99.521 0 0 0 70.088-27.992c16.979-16.246 29.226-38.472 35.419-64.274l.056-.232L229.006 152h-32.69l-25.979 124.488C162.425 309.168 138.766 344 96 344H48.06c.869-113.266 11.182-180.419 33.105-216.634c18.4-30.4 45.295-41.191 102.724-41.206l111.408 1.6a16.026 16.026 0 0 1 12.925 6.891L374.7 190.67l89.3 14.884v16.959l-14.892 79.421c-4.395 23.441-11.908 35.249-42.718 38.95l-126.306 21.609L279.249 496h32l.667-106.493l98.7-16.9c22.36-2.749 38.857-9.955 50.426-22.023c9.89-10.318 15.909-23.5 19.519-42.752L496 225.487v-47.041Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dollar(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M296 240h-80a46.222 46.222 0 1 1 0-92.444h128v-32h-68V56h-32v59.556h-28A78.222 78.222 0 0 0 216 272h80a46.274 46.274 0 0 1 46.222 46.222v3.556A46.274 46.274 0 0 1 296 368H160.593v32H244v56h32v-56h20a78.31 78.31 0 0 0 78.222-78.222v-3.556A78.31 78.31 0 0 0 296 240"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Door(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M400 464V64H112v400H16v32h480v-32Zm-32 0H144V96h224Z"/><path fill="currentColor" d="M312 252h32v72h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleQuoteSansLeft(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M232 246.857V16H16v400h38.4ZM48 48h152v185.143L48 377.905Zm232 368h38.4L496 246.857V16H280Zm32-368h152v185.143L312 377.905Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleQuoteSansRight(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M280 185.143V416h216V16h-38.4ZM464 384H312V198.857L464 54.1ZM232 16h-38.4L16 185.143V416h216Zm-32 368H48V198.857L200 54.1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Drink(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m273.637 128l3.809 32h96.43l-8 64H252.573l-7.442-64h.089l-3.809-32l-.045-.389l-9.041-77.745L230.247 32H104v32h97.753l7.442 64h-91.319l40.5 323.969A32.051 32.051 0 0 0 190.125 480h147.75a32.051 32.051 0 0 0 31.753-28.031L410.124 128Zm-119.513 32h58.792l7.442 64h-58.234Zm183.765 288H190.124l-24-192h57.955l13.953 120h32.215l-13.954-120H361.88Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DrinkAlcohol(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M248 440h-72v32h176v-32h-72V318.968l176-192.762V80H72v46.206l176 192.762ZM104 113.794V112h320v1.794L374.508 168H153.492ZM182.709 200h162.582l-80.349 88h-1.884Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Drop(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M394.633 220.663L269.475 25.174a16 16 0 0 0-26.95 0L117.364 220.665A170.531 170.531 0 0 0 84.1 322.3c0 94.785 77.113 171.9 171.9 171.9s171.9-77.113 171.9-171.9a170.519 170.519 0 0 0-33.267-101.637M256 462.2c-77.14 0-139.9-62.758-139.9-139.9a138.758 138.758 0 0 1 27.321-83.058q.319-.432.608-.884L256 63.475l111.967 174.884q.288.453.608.884A138.754 138.754 0 0 1 395.9 322.3c0 77.141-62.76 139.9-139.9 139.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DropOne(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M18.498 10.344L12.631 1.18a.751.751 0 0 0-1.262-.003l-.002.003L5.5 10.344a7.935 7.935 0 0 0-1.559 4.743v.022v-.001c0 4.443 3.615 8.058 8.058 8.058s8.058-3.615 8.058-8.058v-.021a7.957 7.957 0 0 0-1.574-4.764l.015.021zM12 21.666a6.566 6.566 0 0 1-6.558-6.558v-.018c0-1.46.481-2.807 1.293-3.893l-.012.017l.029-.041l5.249-8.198l5.248 8.198l.029.041a6.456 6.456 0 0 1 1.281 3.876v.019v-.001a6.566 6.566 0 0 1-6.558 6.558z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eco(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M389.053 126.3A302.909 302.909 0 0 0 280.012 18.15L272 13.516l-8.012 4.634A301.084 301.084 0 0 0 113.4 279.042c0 3.445.06 6.944.177 10.4c1.592 46.856 19.511 86.283 51.82 114.018c24.724 21.225 56.438 34.182 90.607 37.273V496h32V240H256v168.528c-54.064-6.263-107.873-44.455-110.444-120.174c-.106-3.095-.16-6.228-.16-9.312A270.286 270.286 0 0 1 272 50.673a270.286 270.286 0 0 1 126.6 228.369c0 3.084-.054 6.217-.16 9.313c-2.056 60.573-36.907 97.127-78.444 112.536v33.867a156.188 156.188 0 0 0 58.607-31.3c32.309-27.735 50.228-67.162 51.82-114.017c.117-3.456.177-6.955.177-10.4A300.948 300.948 0 0 0 389.053 126.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Education(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m368 350.643l-112 63l-112-63v-66.562l-32-17.778v103.054l144 81l144-81V266.303l-32 17.778z"/><path fill="currentColor" d="M256 45.977L32 162.125v27.734L256 314.3l192-106.663V296h32V162.125Zm160 142.831l-32 17.777L256 277.7l-128-71.115l-32-17.777l-22.179-12.322L256 82.023l182.179 94.463Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Elevator(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 30.4L96 168.681V232h320v-63.319ZM384 200H128v-16.681L256 72.7l128 110.619ZM96 343.319L256 481.6l160-138.281V280H96ZM128 312h256v16.681L256 439.3L128 328.681Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeClosed(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 112v384h480V112Zm220.8 229.6a32.167 32.167 0 0 0 38.4 0l23.467-17.6L464 448v16H48v-16l165.333-124ZM256 316L48 160v-16h416v16ZM48 200l138.667 104L48 408Zm416 208L325.333 304L464 200Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeLetter(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M400 163.2V68a20.023 20.023 0 0 0-20-20H132a20.023 20.023 0 0 0-20 20v95.2l-96 68.566V496h480V231.766Zm53.679 77.667L400 275.96v-73.44ZM144 80h224v216.883l-57.166 37.378l-46.578-24.152l-50.764 24.507L144 292.425Zm119.744 265.89L464 449.727V464H48v-13.957ZM48 271.575L179.144 351.2L48 414.509Zm295.446 79.6L464 272.347v141.334ZM112 202.52V273l-53.334-32.385Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeOpen(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M274.6 25.623a32.006 32.006 0 0 0-37.2 0L16 183.766V496h480V183.766ZM464 402.693L339.97 322.96L464 226.492ZM256 51.662L454.429 193.4L311.434 304.615L256 268.979l-55.434 35.636L57.571 193.4ZM48 226.492l124.03 96.468L48 402.693ZM464 464H48v-23.265l208-133.714l208 133.714Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Equalizer(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M96 146.025V16H64v130.025a64.009 64.009 0 0 0 0 123.95V496h32V269.975a64.009 64.009 0 0 0 0-123.95M80 240a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32m192 50.025V16h-32v274.025a64.009 64.009 0 0 0 0 123.95V496h32v-82.025a64.009 64.009 0 0 0 0-123.95M256 384a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32M448 82.025V16h-32v66.025a64.009 64.009 0 0 0 0 123.95V496h32V205.975a64.009 64.009 0 0 0 0-123.95M432 176a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ethernet(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m385.128 118.178l-25.26 19.644l91.864 118.122l-91.919 118.236l25.263 19.64l107.192-137.88zm-265.314-.001L12.621 255.993l107.138 137.826l25.263-19.638l-91.866-118.182l91.918-118.176zM160 240h32v32h-32zm80 0h32v32h-32zm80 0h32v32h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Euro(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M316 120h60V88h-60c-73.747 0-136.794 46.657-161.195 112H104v32h42.292a172.176 172.176 0 0 0 0 56H104v32h50.805c24.4 65.343 87.448 112 161.2 112h60v-32H316a140.176 140.176 0 0 1-126.474-80H344v-32H178.815a140.661 140.661 0 0 1 0-56H344v-32H189.526A140.176 140.176 0 0 1 316 120"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Excerpt(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 344h480v32H16zm0-191.667h480v32H16zM16 248h328v32H16zm0-192h384v32H16zm0 384h32v32H16zm224 0h32v32h-32zm-112 0h32v32h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExitToApp(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M472 16H40a24.028 24.028 0 0 0-24 24v160h32V48h416v416H48V304H16v168a24.028 24.028 0 0 0 24 24h432a24.028 24.028 0 0 0 24-24V40a24.028 24.028 0 0 0-24-24"/><path fill="currentColor" d="m209.377 363.306l22.627 22.627L366.627 251.31L232.004 116.687l-22.627 22.626l95.997 95.998H16v32h289.372z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExpandDown(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 16h480v32H16zm0 480h480V368H16Zm32-96h416v64H48ZM416 96H96v37.86l159.923 169.364L416 135.921ZM256.077 256.776L134.478 128h244.813Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExpandLeft(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 16h32v480h-32zm-320 0H16v480h128Zm-32 448H48V48h64ZM416 96h-37.86L208.776 255.923L376.079 416H416Zm-32 283.291L255.224 256.077L384 134.478Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExpandRight(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M40 19h32v480H40zm352 480h128V19H392Zm32-448h64v416h-64ZM120 419h37.86l169.364-159.923L159.921 99H120Zm32-283.291l128.776 123.215L152 380.522Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExpandUp(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 467h480v32H16zm0-320h480V19H16Zm32-96h416v64H48Zm208.077 160.777L96 379.079V419h320v-37.86ZM132.709 387l123.214-128.776L377.522 387Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Exposure(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M456 40H56a16 16 0 0 0-16 16v400a16 16 0 0 0 16 16h400a16 16 0 0 0 16-16V56a16 16 0 0 0-16-16M72 72h345.373L72 417.373Zm368 368H94.627L440 94.627Z"/><path fill="currentColor" d="M336 368v40h32v-40h40v-32h-40v-40h-32v40h-40v32zM112 136h112v32H112z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExternalLink(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M400 464H48V104h192V72H16v424h416V272h-32z"/><path fill="currentColor" d="M304 16v32h137.373L188.687 300.687l22.626 22.626L464 70.627V208h32V16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eyedropper(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.409 431.289v41.468a24.124 24.124 0 0 0 24.1 24.1h33.368a149.445 149.445 0 0 0 106.373-44.065l200.716-200.717l55.846 56.069l22.671-22.582l-55.889-56.114l1.287-1.288l-.021-.02l63.287-63.287a84.853 84.853 0 0 0 0-120l-4-4a84.852 84.852 0 0 0-120 0l-64.326 64.326L230.947 53.1l-22.672 22.58L260.387 128L63.471 324.916a149.449 149.449 0 0 0-44.062 106.373M369.774 63.48a52.854 52.854 0 0 1 74.747 0l4 4a52.913 52.913 0 0 1 0 74.745L387.147 203.6L308.4 124.853Zm-85.72 86.107l78.573 78.573l-69.176 69.176v-.483H136.788ZM104.788 328.853h157.147L160.623 430.165a117.662 117.662 0 0 1-83.746 34.688H51.409v-33.564A117.664 117.664 0 0 1 86.1 347.543Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Face(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 16C123.452 16 16 123.452 16 256s107.452 240 240 240s240-107.452 240-240S388.548 16 256 16m147.078 387.078a207.253 207.253 0 1 1 44.589-66.125a207.332 207.332 0 0 1-44.589 66.125"/><path fill="currentColor" d="M128 192h40v40h-40zm208 0h40v40h-40zM232 306.948a5.049 5.049 0 0 1 .194-1.387L264 194.241V168h-32v21.759l-30.574 107.009A37.052 37.052 0 0 0 237.052 344H296v-32h-58.948a5.057 5.057 0 0 1-5.052-5.052"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceDead(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 16C123.452 16 16 123.452 16 256s107.452 240 240 240s240-107.452 240-240S388.548 16 256 16m147.078 387.078a207.253 207.253 0 1 1 44.589-66.125a207.332 207.332 0 0 1-44.589 66.125"/><path fill="currentColor" d="M168 320h176v32H168zm42.63-91.958l-24.042-21.371l21.37-24.041l-23.916-21.26l-21.371 24.042l-24.041-21.37l-21.26 23.916l24.042 21.371l-21.37 24.041l23.916 21.26l21.371-24.042l24.041 21.37zm173.328-45.412l-23.916-21.26l-21.371 24.042l-24.041-21.37l-21.26 23.916l24.042 21.371l-21.37 24.041l23.916 21.26l21.371-24.042l24.041 21.37l21.26-23.916l-24.042-21.371z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Factory(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M459.26 96L328 225.7V96h-34.525L168 223.267V16H16v480h480V96ZM464 464H48V48h88v216h36.778L296 139.018V264h38.764L464 136.3Z"/><path fill="currentColor" d="M136 328v8h32v-32h-32zm0 48h32v32h-32zm80-48v8h32v-32h-32zm0 48h32v32h-32zm80-48v8h32v-32h-32zm0 48h32v32h-32zm80-72h32v32h-32zm0 72h32v32h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FactorySlash(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M459.26 96L328 225.697V96h-34.55l-64.44 65.128l22.628 22.628L296 138.92v89.198l37.314 37.315L464 136.303v259.815l32 32V96zM168 16H83.882L168 100.118zm-32 288h32v32h-32zm0 72h32v32h-32zm80-72h32v32h-32zm0 72h32v32h-32zm80 0h32v32h-32z"/><path fill="currentColor" d="M38.627 16H16v480h480v-22.627ZM48 464V70.627l88 88V248h32v-57.373L441.373 464Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fastfood(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M474.183 120H360V40h-32v80H200v32h17.92l6.154 48H112a80.091 80.091 0 0 0-80 80v16h280v-16a79.508 79.508 0 0 0-8-34.846a80.248 80.248 0 0 0-47.155-41.185L250.183 152h187.634l-35.9 280H32v32h398.08l40-312H496v-32ZM277.258 264H66.742A48.083 48.083 0 0 1 112 232h120a48.083 48.083 0 0 1 45.258 32"/><path fill="currentColor" d="M304 352h8v-32H32v32h152zm0 56h8v-32H32v32h152z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fax(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M487.944 233.654L424 176.815V89.373L350.627 16H152v128H80a24.028 24.028 0 0 0-24 24v304a24.028 24.028 0 0 0 24 24h416V251.593a24.024 24.024 0 0 0-8.056-17.939M152 464H88V176h64ZM328 48h9.373L392 102.627V112h-64Zm-144 0h112v96h96v72H184Zm280 416H184V248h240v-28.371l40 35.557Z"/><path fill="currentColor" d="M232 344h32v32h-32zm64 0h32v32h-32zm-64 64h32v32h-32zm64 0h32v32h-32zm64-64h32v32h-32zm0 64h32v32h-32zM232 280h160v32H232z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FeaturedPlaylist(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 72H48a32.036 32.036 0 0 0-32 32v304a32.036 32.036 0 0 0 32 32h416a32.036 32.036 0 0 0 32-32V104a32.036 32.036 0 0 0-32-32m0 336H48V104h416l.02 304Z"/><path fill="currentColor" d="M232 184h184v32H232zm-56 72h240v32H176zm0 72h240v32H176zM88.923 144v128.923l99.172-69.42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func File(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M334.627 16H48v480h424V153.373ZM440 166.627V168H320V48h1.373ZM80 464V48h208v152h152v264Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Filter(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M238.627 496H192V253.828l-168-200V16h456v37.612l-160 200v161.015ZM224 464h1.373L288 401.373V242.388L443.51 48H60.9L224 242.172Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterFrames(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 104H128v236.117h256Zm-32 204.117H160V136h192Z"/><path fill="currentColor" d="M181.49 428.117L256 502.628l74.51-74.511H440a32.036 32.036 0 0 0 32-32V48a32.036 32.036 0 0 0-32-32H72a32.036 32.036 0 0 0-32 32v348.117a32.036 32.036 0 0 0 32 32ZM72 48h368v348.117H317.255L256 457.372l-61.255-61.255H72Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterPhoto(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M440 40H72a32.036 32.036 0 0 0-32 32v368a32.036 32.036 0 0 0 32 32h368a32.036 32.036 0 0 0 32-32V72a32.036 32.036 0 0 0-32-32m0 400H72V72h368Z"/><path fill="currentColor" d="m120.614 261.739l.031 2.261l.317 22.978l21.874 7.041l60.183 19.373l18.781 59.642l7.051 22.387h54.3l7.051-22.387l18.782-59.642l60.183-19.373l21.874-7.041l.314-22.978l.031-2.261l.071-5.092l.322-23.282l-22.053-7.465l-56.47-19.127l-23.3-68.995l-7.348-21.763h-53.22l-7.348 21.763l-23.3 69l-56.466 19.122l-22.053 7.469l.322 23.282ZM224 232l28.359-83.989h7.282L288 232l71.46 24.2c-.049 3.519-.053 3.836-.1 7.354L283.428 288l-23.751 75.421h-7.354L228.572 288l-75.93-24.442c-.049-3.518-.053-3.835-.1-7.354Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterSquare(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M472 16H40a24.027 24.027 0 0 0-24 24v432a24.027 24.027 0 0 0 24 24h432a24.028 24.028 0 0 0 24-24V40a24.028 24.028 0 0 0-24-24m-8 448H48V48h416Z"/><path fill="currentColor" d="M208 427h48.627L301 368.076V260.87l114-129.018V85H89v47.176l119 129.018Zm-87-307.328V117h262v2.739L269 248.757v108.618L240.666 395H240V248.69Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterX(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M40 16v37.828L109.024 136h41.791L76.896 48H459.51L304 242.388v158.985L241.373 464H240v-96h-32v128h46.627L336 414.627V253.612l160-200V16z"/><path fill="currentColor" d="m166.403 248.225l60.461-60.462l-22.627-22.628l-60.462 60.462l-60.462-60.462l-22.626 22.628l60.461 60.462l-60.461 60.462l22.626 22.627l60.462-60.462l60.462 60.462l22.627-22.627z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FindInPage(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M334.627 16H48v480h424V153.373ZM440 464H80V48h241.373L440 166.627Z"/><path fill="currentColor" d="M239.861 152a95.861 95.861 0 1 0 53.624 175.284l68.03 68.029l22.627-22.626l-67.5-67.5A95.816 95.816 0 0 0 239.861 152M176 247.861a63.862 63.862 0 1 1 63.861 63.861A63.933 63.933 0 0 1 176 247.861"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fingerprint(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m148.492 332.884l122.084-126.386l23.015 22.232l-122.083 126.387z"/><path fill="currentColor" d="M330.409 282.426A80 80 0 0 0 217.22 169.341l-74.908 74.908A121.268 121.268 0 0 1 56 280h-8v32h8a153.052 153.052 0 0 0 108.94-45.125l74.907-74.907a48 48 0 0 1 67.883 67.883l-68.209 69.395A136.994 136.994 0 0 0 200 425.832V464h32v-38.168a105.179 105.179 0 0 1 30.342-74.155Z"/><path fill="currentColor" d="M397.139 328.853c-19.869 21.6-82.507 85-83.138 85.633l-.1.1a12.1 12.1 0 1 1-17.362-16.87l79.154-81.381a135.935 135.935 0 0 0-.081-192.253c-53.027-53.025-139.307-53.025-192.333 0l-61.4 63.3C89.791 216 78.35 216 48 216v32c36.545 0 55.876-1.578 95.628-37.139l.423-.378l61.938-63.852a104 104 0 0 1 147 147.16L273.6 375.407a44.1 44.1 0 0 0 63.207 61.528c3.185-3.223 63.734-64.51 83.885-86.42C455.612 312.546 480 260.521 480 224h-32c0 28.249-21.866 73.326-50.861 104.853"/><path fill="currentColor" d="M425.12 77.5C389.951 43.5 345.8 24 304 24v32c33.582 0 69.622 16.221 98.88 44.5C431.133 127.815 448 162.019 448 192h32c0-39.055-20-80.79-54.88-114.5m-292.165 59.89C155.02 106.949 191.951 56 272 56V24c-96.376 0-140.988 61.545-164.955 94.61C90.759 141.078 71.446 152 48 152v32c33.476 0 62.853-16.117 84.955-46.61M80 336H48v32h32a56.064 56.064 0 0 1 56 56v40h32v-40a88.1 88.1 0 0 0-88-88"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fire(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M311.145 257.6c5.5-9.032 11.185-18.371 16.8-28c18.794-32.218 22.243-64.681 10.254-96.489C314.008 68.933 233.142 29.83 214.421 23.591L201.257 19.2l-20.546 41.1l9.915 8.1c.111.089 11.124 9.712 11.858 24.3c.627 12.453-6.2 25.91-20.286 40c-9.782 9.781-20.518 19.239-31.885 29.251C102.487 204.069 48.28 251.82 48.28 342.154q0 1.861.035 3.7a151.362 151.362 0 0 0 46.444 106.49A154.177 154.177 0 0 0 202.583 496h92.1l-11.369-23.072c-46.193-93.751-13.526-147.403 27.831-215.328M202.583 464c-66.2 0-121.05-53.267-122.274-118.739q-.028-1.546-.029-3.107c0-75.878 46.356-116.713 91.186-156.2c11.239-9.9 22.862-20.138 33.359-30.637c20.754-20.753 30.719-42.365 29.619-64.232a64.963 64.963 0 0 0-7.214-26.56c27.84 15.211 67.523 44.053 81.027 79.877c8.544 22.665 5.943 45.26-7.951 69.079c-5.465 9.369-11.072 18.576-16.493 27.481C244.132 306.131 206.49 367.943 244.361 464Z"/><path fill="currentColor" d="M468.243 341.129q-.4-1.586-.834-3.185c-11.546-42.332-75.457-96.762-82.706-102.829l-14.557-12.182l-9.546 16.408c-21.753 37.4-40.421 71.512-48.559 110.212c-9.279 44.134-3.033 88.989 19.1 137.13l4.281 9.317H346.3a125.168 125.168 0 0 0 99.3-48.5a123.175 123.175 0 0 0 22.643-106.371m-47.88 86.683a93.3 93.3 0 0 1-64.44 35.7c-31.541-75.9-12.931-127.635 22.63-191.114c23.642 21.994 52.455 53.69 57.983 73.961q.349 1.283.669 2.554a91.372 91.372 0 0 1-16.842 78.899"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlagAlt(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 16v480h32V360h352v-37.238L363.841 208L448 93.238V56H96V16Zm348.159 72l-88 120l88 120H96V88Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlightTakeoff(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 464h480v32H16zm439.688-311.836c-23.388-6.515-48.252-6.053-70.008 1.3l-.894.3l-65.1 30.94l-189.981-75.528a47.719 47.719 0 0 0-49.771 8.862L54.5 140.836a24 24 0 0 0 2.145 37.452l117.767 83.458l-45.173 23.663l-35.775-32.687a48.067 48.067 0 0 0-51.47-8.6l-19.455 8.435a24 24 0 0 0-11.642 33.3l72.821 136.827L480.3 227.21c23.746-11.177 26.641-29.045 21.419-42.059c-5.788-14.428-22.568-26.451-46.031-32.987m10.9 46.133l-.149.07l-369.045 181.9l-54.176-101.8l11.5-4.987a16.021 16.021 0 0 1 17.157 2.867l52.336 47.819l111.329-58.318L83.322 157.974l17.971-16.108a15.908 15.908 0 0 1 16.59-2.954l202.943 80.681l75.95-36.095c15.456-5.009 33.863-5.165 50.662-.413c13.834 3.914 21.182 9.6 23.672 12.582a24.211 24.211 0 0 1-4.52 2.633Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flip(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M296 40h32v32h-32zm72 0h32v32h-32zm-72 400h32v32h-32zm72 0h32v32h-32zm72-400h32v32h-32zm0 65.454h32v33.454h-32zm0 66.909h32v33.454h-32zm0 200.728h32v33.454h-32zm0-133.819h32v33.454h-32zm0 66.91h32v33.454h-32zM440 440h32v32h-32zM40 456a16 16 0 0 0 16 16h168v24h32V16h-32v24H56a16 16 0 0 0-16 16ZM72 72h152v368H72Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlipToBack(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 496h336v-32H64a16.019 16.019 0 0 1-16-16V120H16v328a48.055 48.055 0 0 0 48 48m400-128h32v32h-32z"/><path fill="currentColor" d="M173.091 368h29.091v32h-29.091zm232.727 0h29.091v32h-29.091zm-174.546 0h29.091v32h-29.091zm58.183 0h29.091v32h-29.091zm58.182 0h29.091v32h-29.091zM112 368h32v32h-32zm0-232.728h32v29.091h-32zm0-58.181h32v29.091h-32zm0 232.727h32v29.091h-32zm0-58.181h32v29.091h-32zm0-58.182h32v29.091h-32zM112 16h32v32h-32zm177.454 0h29.091v32h-29.091zm-58.182 0h29.091v32h-29.091zm174.546 0h29.091v32h-29.091zm-232.727 0h29.091v32h-29.091zm174.546 0h29.091v32h-29.091zM464 16h32v32h-32zm0 293.818h32v29.091h-32zm0-232.727h32v29.091h-32zm0 116.363h32v29.091h-32zm0-58.182h32v29.091h-32zm0 116.365h32v29.091h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlipToFront(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 16H144a32.036 32.036 0 0 0-32 32v320a32.036 32.036 0 0 0 32 32h320a32.036 32.036 0 0 0 32-32V48a32.036 32.036 0 0 0-32-32m0 352H144V48h320l.021 320ZM77.846 464h30.77v32h-30.77zM384 464h30.77v32H384zm-183.692 0h30.77v32h-30.77zm-61.231 0h30.769v32h-30.769zm122.462 0h30.769v32h-30.769zm61.23 0h30.77v32h-30.77zM16 464h32v32H16zm0-239.692h32v30.154H16zm0 60.153h32v30.154H16zm0 120.308h32v30.154H16zm0-60.154h32v30.154H16zm0-180.461h32v30.154H16zM16 104h32v30.154H16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flower(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M408 304a168.212 168.212 0 0 0-152 96.5A168.212 168.212 0 0 0 104 304H16v16c0 92.636 75.364 168 168 168h144c92.636 0 168-75.364 168-168v-16ZM184 456c-69.581 0-127.124-52.519-135.064-120H104c69.581 0 127.124 52.519 135.064 120Zm144 0h-55.064c7.94-67.481 65.483-120 135.064-120h55.064c-7.94 67.481-65.483 120-135.064 120"/><path fill="currentColor" d="M169.227 262.773a87.36 87.36 0 0 0 24.547 47.453l4.687 4.686h6.627A87.354 87.354 0 0 0 256 298.716a87.356 87.356 0 0 0 50.912 16.2h6.627l4.687-4.686a87.36 87.36 0 0 0 24.547-47.453a87.36 87.36 0 0 0 47.453-24.547l4.686-4.687v-6.627A87.354 87.354 0 0 0 378.716 176a87.356 87.356 0 0 0 16.2-50.912v-6.627l-4.686-4.687a87.36 87.36 0 0 0-47.453-24.547a87.36 87.36 0 0 0-24.547-47.453l-4.687-4.686h-6.627A87.356 87.356 0 0 0 256 53.284a87.354 87.354 0 0 0-50.912-16.2h-6.627l-4.687 4.686a87.36 87.36 0 0 0-24.547 47.453a87.36 87.36 0 0 0-47.453 24.547l-4.686 4.687v6.627A87.356 87.356 0 0 0 133.284 176a87.354 87.354 0 0 0-16.2 50.912v6.627l4.686 4.687a87.36 87.36 0 0 0 47.457 24.547m-3.736-98.086a55.571 55.571 0 0 1-16-32.8A55.572 55.572 0 0 1 184 120h16v-16a55.572 55.572 0 0 1 11.884-34.506a55.57 55.57 0 0 1 32.8 16L256 96.8l11.313-11.313a55.571 55.571 0 0 1 32.8-16A55.572 55.572 0 0 1 312 104v16h16a55.572 55.572 0 0 1 34.506 11.884a55.571 55.571 0 0 1-16 32.8L335.2 176l11.314 11.314a55.57 55.57 0 0 1 16 32.8A55.572 55.572 0 0 1 328 232h-16v16a55.572 55.572 0 0 1-11.884 34.506a55.571 55.571 0 0 1-32.8-16L256 255.2l-11.314 11.31a55.57 55.57 0 0 1-32.8 16A55.572 55.572 0 0 1 200 248v-16h-16a55.572 55.572 0 0 1-34.506-11.884a55.57 55.57 0 0 1 16-32.8L176.8 176Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folder(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M472 472H40a24.028 24.028 0 0 1-24-24V64a24.028 24.028 0 0 1 24-24h186.667a23.935 23.935 0 0 1 22.154 14.77L269.333 104H472a24.028 24.028 0 0 1 24 24v320a24.028 24.028 0 0 1-24 24M48 440h416V136H248l-26.667-64H48Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOpen(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M488.671 221.645a23.848 23.848 0 0 0-18.917-9.231h-66.2V160a23.138 23.138 0 0 0-23.112-23.111h-136.3L226.37 94.22A23.051 23.051 0 0 0 205.037 80H39.111A23.138 23.138 0 0 0 16 103.111v364.445h420.707l56.33-225.321a23.849 23.849 0 0 0-4.366-20.59M48 435.556V112h151.111l23.7 56.889h148.745v43.525H137.587a23.965 23.965 0 0 0-23.287 18.179L63.063 435.556Zm363.723 0H96.048l47.785-191.142h315.675Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Font(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M221.631 109L109.92 392h58.055l24.079-61h127.892l24.079 61h58.055L290.369 109Zm-8.261 168L256 169l42.63 108Z"/><path fill="currentColor" d="M16 496h480V16H16ZM48 48h416v416H48Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Football(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M425.706 86.294A240 240 0 0 0 86.294 425.706A240 240 0 0 0 425.706 86.294m-4.464 43.535A206.875 206.875 0 0 1 463.824 247.8l-66.14-49.772ZM316.033 56.845l-58.378 43.37l-61.125-43.538a208.143 208.143 0 0 1 119.5.168ZM116.8 198.047L48.156 248.33a206.9 206.9 0 0 1 43.092-119.141ZM86.2 376h79.365l25.035 77.458A208.923 208.923 0 0 1 86.2 376m140.787 85.967L188.85 344H67.562a206.3 206.3 0 0 1-17.324-57.527l104.967-76.9l-39.648-106.858a208.938 208.938 0 0 1 45.714-31.864l96.781 68.934l92.741-68.9a208.922 208.922 0 0 1 45.884 32.048L359.833 209.6l101.951 76.721A206.272 206.272 0 0 1 444.438 344H327.512l-42.045 117.9a208.076 208.076 0 0 1-58.482.064Zm95.606-8.9L350.075 376H425.8a208.961 208.961 0 0 1-103.209 77.069Z"/><path fill="currentColor" d="M346.809 223.427L257.854 158.8L168.9 223.427L202.876 328h109.955ZM289.582 296h-63.457l-19.609-60.351l51.338-37.3l51.337 37.3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fork(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M124 166.291v179.418a76 76 0 1 0 32 0V282h152a80.091 80.091 0 0 0 80-80v-36.689a75.983 75.983 0 1 0-32 1.733V202a48.055 48.055 0 0 1-48 48H156v-83.709a76 76 0 1 0-32 0M324 92a44 44 0 1 1 44 44a44.049 44.049 0 0 1-44-44M184 420a44 44 0 1 1-44-44a44.049 44.049 0 0 1 44 44M140 48a44 44 0 1 1-44 44a44.049 44.049 0 0 1 44-44"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FourK(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M208 184h-32v88h-62.32l13.6-136H95.12l-16.8 168H176v72h32v-72h32v-32h-32zm178.111-48l-52 104H304V136h-32v240h32V272h30.111l52 104h35.778l-60-120l60-120z"/><path fill="currentColor" d="M464 16H48a32.036 32.036 0 0 0-32 32v416a32.036 32.036 0 0 0 32 32h416a32.036 32.036 0 0 0 32-32V48a32.036 32.036 0 0 0-32-32m0 448H48V48h416l.02 416Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fridge(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M376 16H136a48.054 48.054 0 0 0-48 48v400a32.036 32.036 0 0 0 32 32h272a32.036 32.036 0 0 0 32-32V64a48.054 48.054 0 0 0-48-48m16 448H120V240h272Zm0-256H120V64a16.019 16.019 0 0 1 16-16h240a16.019 16.019 0 0 1 16 16Z"/><path fill="currentColor" d="M144 280h32v96h-32zm0-176h32v64h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Frown(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 16C123.452 16 16 123.452 16 256s107.452 240 240 240s240-107.452 240-240S388.548 16 256 16m147.078 387.078a207.253 207.253 0 1 1 44.589-66.125a207.332 207.332 0 0 1-44.589 66.125"/><path fill="currentColor" d="M152 200h40v40h-40zm168 0h40v40h-40zm-64 88a88.1 88.1 0 0 0-88 88h32a56 56 0 0 1 112 0h32a88.1 88.1 0 0 0-88-88"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fullscreen(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M208 48V16H16v192h32V70.627l160.687 160.686l22.626-22.626L70.627 48zm256 256v137.373L299.313 276.687l-22.626 22.626L441.373 464H304v32h192V304z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FullscreenExit(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M204 181.372L38.628 16H16v22.628L181.372 204H44v32h192V44h-32zM326.628 304H464v-32H272v192h32V326.628L473.372 496H496v-22.628z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Functions(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m402.034 206.447l-67.528 65.278l-3.046-14.231l-.044-.2a63.447 63.447 0 0 0-59.129-49.282l-.329-.014h-67.184l18.846-87.688A31.451 31.451 0 0 1 252.549 96h47.637V64h-48l-.659.014A63.452 63.452 0 0 0 192.4 113.3L172.044 208H88v32h77.166l-32.6 151.688A31.451 31.451 0 0 1 103.637 416H56v32h48.329l.329-.014a63.452 63.452 0 0 0 59.13-49.282L197.9 240h73.369a31.453 31.453 0 0 1 28.925 24.3l7.2 33.639L193.531 408h46.042l75.711-73.187l5.07 23.693l.043.2a63.449 63.449 0 0 0 59.13 49.282L432 408v-32h-51.451a31.452 31.452 0 0 1-28.927-24.3l-9.222-43.1l105.676-102.153Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FunctionsAlt(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M128 416h232v-32H160v-9.717L269.834 256L160 137.717V128h200V96H128v54.283L226.166 256L128 361.717z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gamepad(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M368 168h-96v-64h176a24.027 24.027 0 0 0 24-24V16h-32v56H264a24.027 24.027 0 0 0-24 24v72h-96A128.145 128.145 0 0 0 16 296v100.953A91.15 91.15 0 0 0 107.047 488h1.8a90.807 90.807 0 0 0 69.953-32.76L231.5 392h48.628l52.666 68.465A91.046 91.046 0 0 0 496 404.953V296a128.145 128.145 0 0 0-128-128m96 236.953a59.047 59.047 0 0 1-105.849 36L295.878 360h-79.372l-62.294 74.754A58.893 58.893 0 0 1 108.85 456h-1.8A59.113 59.113 0 0 1 48 396.953V296a96.108 96.108 0 0 1 96-96h224a96.108 96.108 0 0 1 96 96Z"/><path fill="currentColor" d="M360 248h32v32h-32zm0 80h32v32h-32zm-40-40h32v32h-32zm80 0h32v32h-32zm-248-40h-32v40H80v32h40v40h32v-40h40v-32h-40z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Garage(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 255.454h-.511L407.067 164.5A48.044 48.044 0 0 0 363.2 136H148.8a48.043 48.043 0 0 0-43.863 28.5l-40.426 90.954H64a32.036 32.036 0 0 0-32 32v112a32.036 32.036 0 0 0 32 32V472a24.028 24.028 0 0 0 24 24h56a24.028 24.028 0 0 0 24-24v-40.546h176V472a24.028 24.028 0 0 0 24 24h56a24.028 24.028 0 0 0 24-24v-40.546a32.036 32.036 0 0 0 32-32v-112a32.036 32.036 0 0 0-32-32M134.175 177.5A16.013 16.013 0 0 1 148.8 168h214.4a16.014 16.014 0 0 1 14.621 9.5l34.646 77.953H99.529ZM136 464H96v-32.546h40Zm280 0h-40v-32.546h40Zm32-64.546H64v-112h384Z"/><path fill="currentColor" d="M96 328h80v32H96zm240 0h80v32h-80zM256 14.758L16 111.121v34.483l240-96.362l240 96.362v-34.483z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gauge(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M425.706 142.294A240 240 0 0 0 16 312v88h144v-32H48v-56c0-114.691 93.309-208 208-208s208 93.309 208 208v56H352v32h144v-88a238.432 238.432 0 0 0-70.294-169.706"/><path fill="currentColor" d="M80 264h32v32H80zm160-136h32v32h-32zm-104 40h32v32h-32zm264 96h32v32h-32zm-102.778 71.1l69.2-144.173l-28.85-13.848l-69.183 144.135a64.141 64.141 0 1 0 28.833 13.886M256 416a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gem(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M408.563 48H103.438L16 179.156v24.786L199.421 480h113.158L496 203.942v-24.786Zm-17.125 32l63.407 95.111H347.739L317.808 80Zm-163.7 0h56.524l29.93 95.111H197.808Zm-107.175 0h73.629l-29.931 95.111H57.155Zm96.016 368L56.525 207.111h106.793L219.813 448Zm36.1 0l-56.492-240.889h119.626L259.318 448Zm42.739 0h-3.234l56.5-240.889h106.791Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gif(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 136v224a32.036 32.036 0 0 0 32 32h96V232h-64v32h32v96H96V136h96v-32H96a32.036 32.036 0 0 0-32 32m176-32h32v288h-32zm208 32v-32H328v288h32V264h72v-32h-72v-96z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gift(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M472 128H371.833A71.962 71.962 0 0 0 256 42.815A71.962 71.962 0 0 0 140.167 128H40a24.028 24.028 0 0 0-24 24v321.556C16 485.932 26.767 496 40 496h432c13.233 0 24-10.068 24-22.444V152a24.028 24.028 0 0 0-24-24M312 48a40 40 0 0 1 0 80h-40V88a40.045 40.045 0 0 1 40-40M160 88a40 40 0 0 1 80 0v40h-40a40.045 40.045 0 0 1-40-40M48 464V256h144v208Zm176 0V256h64v208Zm240 0H320V256h144ZM48 224v-64h416v64Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlobeAlt(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 16C123.452 16 16 123.452 16 256s107.452 240 240 240s240-107.452 240-240S388.548 16 256 16m-80.953 431.667a208.26 208.26 0 0 1-110.714-272.62a200.793 200.793 0 0 1 3.2-7.145L120 215.126v63.235L197.1 360H236v49.047l-47.052 43.915q-7.022-2.394-13.901-5.295m228.031-44.589A207.253 207.253 0 0 1 256 464a210.4 210.4 0 0 1-29.722-2.107L268 422.953V328h-57.1L152 265.639v-64.765l-68.3-61.466a209.259 209.259 0 0 1 91.343-75.075A207.793 207.793 0 0 1 371.3 82.839l-45.564 58.582l15.49 38.725l-10.485 10.485l-78.618-15.723L208 208v88h176l50.345 67.126a208.127 208.127 0 0 1-31.267 39.952M464 256a206.763 206.763 0 0 1-13.873 74.837L400 264H240v-40l19.877-14.908l81.382 16.277l37.515-37.515l-16.51-41.275l34.2-43.977q3.361 3.084 6.61 6.32A207.253 207.253 0 0 1 464 256"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Golf(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M378.841 342.034C339.372 327.825 287.223 320 232 320a537.769 537.769 0 0 0-80 5.793V211.157l152-56v-38.131l-152-71.2V16h-32v400h32v-57.834A501.048 501.048 0 0 1 232 352c51.621 0 99.921 7.153 136 20.143C403.43 384.9 416 399.43 416 408s-12.57 23.1-48 35.857C331.921 456.847 283.621 464 232 464s-99.921-7.153-136-20.143C60.57 431.1 48 416.57 48 408c0-7.89 10.669-20.832 40-32.788v-34.176c-.952.33-1.9.661-2.841 1C28 362.612 16 389.265 16 408s12 45.388 69.159 65.966C124.628 488.175 176.777 496 232 496s107.372-7.825 146.841-22.034C436 453.388 448 426.735 448 408s-12-45.388-69.159-65.966M152 81.163l114.586 53.675L152 177.054Z"/><circle cx="432" cy="280" r="32" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GolfAlt(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M161.846 320h192l4.923-32h-192z"/><path fill="currentColor" d="M467.122 16L450.7 128h-29.525L439.49 16h-32.425l-35.321 216H155.819a48.051 48.051 0 0 0-47.152 39.019l-9.354 49.106A88 88 0 1 0 188.673 432H314.05a144 144 0 0 0 142.589-123.889L499.464 16ZM104 464a56 56 0 1 1 56-56a56.063 56.063 0 0 1-56 56m320.963-160.433A111.266 111.266 0 0 1 400 359.585A113.582 113.582 0 0 1 371.477 384a111.226 111.226 0 0 1-57.427 16H191.633a87.325 87.325 0 0 0-5.657-24h159.255l4.923-32H164.333q-2.626-2.476-5.451-4.735a87.882 87.882 0 0 0-27.783-14.99L139.532 280l.571-2.994A16.015 16.015 0 0 1 155.819 264h243.117l17.006-104h30.068Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gradient(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M440 40H72a16 16 0 0 0-16 16v400a16 16 0 0 0 16 16h368a16 16 0 0 0 16-16V56a16 16 0 0 0-16-16m-16 192h-48v48h48v48h-48v40h-48v-40h-48v40h-48v-40h-48v40h-48v-40H88v-48h48v-48H88V72h336Z"/><path fill="currentColor" d="M136 280h48v48h-48zm48-48h48v48h-48zm48 48h48v48h-48zm48-48h48v48h-48zm48 48h48v48h-48zm-192-96h48v48h-48zm96 0h48v48h-48zm96 0h48v48h-48z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Grain(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M88 16a56 56 0 1 0 56 56a56.063 56.063 0 0 0-56-56m0 80a24 24 0 1 1 24-24a24.028 24.028 0 0 1-24 24m112 42.667a56 56 0 1 0 56 56a56.063 56.063 0 0 0-56-56m0 80a24 24 0 1 1 24-24a24.028 24.028 0 0 1-24 24M88 373.333a56 56 0 1 0-56-56a56.063 56.063 0 0 0 56 56m0-80a24 24 0 1 1-24 24a24.028 24.028 0 0 1 24-24M200 496a56 56 0 1 0-56-56a56.063 56.063 0 0 0 56 56m0-80a24 24 0 1 1-24 24a24.028 24.028 0 0 1 24-24M312 16a56 56 0 1 0 56 56a56.063 56.063 0 0 0-56-56m0 80a24 24 0 1 1 24-24a24.028 24.028 0 0 1-24 24m112 42.667a56 56 0 1 0 56 56a56.063 56.063 0 0 0-56-56m0 80a24 24 0 1 1 24-24a24.028 24.028 0 0 1-24 24M312 373.333a56 56 0 1 0-56-56a56.063 56.063 0 0 0 56 56m0-80a24 24 0 1 1-24 24a24.028 24.028 0 0 1 24-24M424 384a56 56 0 1 0 56 56a56.063 56.063 0 0 0-56-56m0 80a24 24 0 1 1 24-24a24.028 24.028 0 0 1-24 24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Graph(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M432 32a63.973 63.973 0 0 0-34.344 117.963L355.894 296.13A64.372 64.372 0 0 0 352 296a63.659 63.659 0 0 0-38.193 12.678l-77.154-64.295A64 64 0 1 0 131.259 269.7l-45.292 90.588A64.334 64.334 0 0 0 80 360a64.082 64.082 0 1 0 36.243 11.29l42.8-85.589a63.845 63.845 0 0 0 59.982-14.356l74.7 62.252a64 64 0 1 0 92.621-27.56l41.76-146.167c1.289.078 2.585.13 3.894.13a64 64 0 0 0 0-128M80 456a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32m96-200a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32m176 136a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32m80-264a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Grid(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M47.547 63.547v384.906a16 16 0 0 0 16 16h384.906a16 16 0 0 0 16-16V63.547a16 16 0 0 0-16-16H63.547a16 16 0 0 0-16 16m288.6 16h96.3v96.3h-96.3Zm0 128.3h96.3v96.3h-96.3Zm0 128.3h96.3v96.3h-96.3Zm-128.3-256.6h96.3v96.3h-96.3Zm0 128.3h96.3v96.3h-96.3Zm0 128.3h96.3v96.3h-96.3Zm-128.3-256.6h96.3v96.3h-96.3Zm0 128.3h96.3v96.3h-96.3Zm0 128.3h96.3v96.3h-96.3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GridSlash(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464.453 395V63.547a16 16 0 0 0-16-16H117l32 32h26.847v26.847l32 32V79.547h96.3v96.3H245.3l32 32h26.847V234.7l32 32v-58.851h96.3v96.3h-58.841l32 32h26.847V363Zm-128.3-219.149v-96.3h96.3v96.3ZM16 16.667v22.627l31.547 31.547v377.612a16 16 0 0 0 16 16h377.612l32.214 32.214H496V474.04L38.626 16.667Zm320.151 342.778l73.008 73.008h-73.008Zm-128.3-128.3l73.008 73.008h-73.01Zm0 105.008h96.3v96.3h-96.3Zm-128.3-233.31l73.008 73.008H79.547Zm0 105.008h96.3v96.3h-96.3Zm0 128.3h96.3v96.3h-96.3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GroupIcon(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m462.541 316.3l-64.344-42.1l24.774-45.418A79.124 79.124 0 0 0 432.093 192v-72a103.941 103.941 0 0 0-174.609-76.477L279.232 67a71.989 71.989 0 0 1 120.861 53v72a46.809 46.809 0 0 1-5.215 21.452L355.962 284.8l89.058 58.274a42.16 42.16 0 0 1 19.073 35.421V432h-72v32h104v-85.506a74.061 74.061 0 0 0-33.552-62.194"/><path fill="currentColor" d="m318.541 348.3l-64.343-42.1l24.773-45.418A79.124 79.124 0 0 0 288.093 224v-72A104.212 104.212 0 0 0 184.04 47.866C126.723 47.866 80.093 94.581 80.093 152v72a78 78 0 0 0 9.015 36.775l24.908 45.664L50.047 348.3A74.022 74.022 0 0 0 16.5 410.4L16 496h336.093v-85.506a74.061 74.061 0 0 0-33.552-62.194m1.552 115.7H48.186l.31-53.506a42.158 42.158 0 0 1 19.073-35.421l88.682-58.029l-39.051-71.592A46.838 46.838 0 0 1 112.093 224v-72a72 72 0 1 1 144 0v72a46.809 46.809 0 0 1-5.215 21.452L211.962 316.8l89.058 58.274a42.16 42.16 0 0 1 19.073 35.421Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HamburgerMenu(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M80 96h352v32H80zm0 144h352v32H80zm0 144h352v32H80z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandPointDown(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M460.579 101.868L409.06 16H264.892L41.573 263.245a24 24 0 0 0 2.007 34.148A90.41 90.41 0 0 0 173.089 286.6l18.493-22.6H216v180a52 52 0 0 0 104 0V332.162l126.423-35.119A24.071 24.071 0 0 0 464 273.919v-159.7a24.006 24.006 0 0 0-3.421-12.351M432 267.838l-144 40V444a20 20 0 0 1-40 0V232h-71.582l-28.1 34.34a58.439 58.439 0 0 1-77.181 11.91L279.108 48H390.94L432 116.433Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandPointLeft(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M397.784 472h-159.7a24.072 24.072 0 0 1-23.124-17.576L179.838 328H68a52 52 0 0 1 0-104h180v-24.418l-22.6-18.494A90.41 90.41 0 0 1 214.607 51.58a24 24 0 0 1 34.149-2.006L496 272.891v144.168l-85.868 51.521a24 24 0 0 1-12.348 3.42m-153.623-32h151.407L464 398.941V287.109L233.75 79.141a58.437 58.437 0 0 0 11.91 77.18l34.34 28.1V256H68a20 20 0 0 0 0 40h136.162ZM227.307 73.321l.023.02Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandPointRight(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M273.919 472h-159.7a24 24 0 0 1-12.349-3.421L16 417.059V272.891L263.244 49.574a24 24 0 0 1 34.149 2.007A90.409 90.409 0 0 1 286.6 181.088L264 199.582V224h180a52 52 0 0 1 0 104H332.162l-35.119 126.423A24.071 24.071 0 0 1 273.919 472m-157.487-32h151.406l40-144H444a20 20 0 0 0 0-40H232v-71.582l34.34-28.1a58.437 58.437 0 0 0 11.91-77.18L48 287.109v111.832ZM284.693 73.321l-.023.02z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandPointUp(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M417.059 497.692H272.891L49.574 250.448a24 24 0 0 1 2.007-34.148a90.409 90.409 0 0 1 129.507 10.789l18.494 22.6H224v-180a52 52 0 0 1 104 0v111.842l126.423 35.118A24.071 24.071 0 0 1 472 239.773v159.7a24 24 0 0 1-3.421 12.349Zm-129.95-32h111.832L440 397.26V245.854l-144-40V69.692a20 20 0 0 0-40 0v212h-71.582l-28.1-34.34a58.437 58.437 0 0 0-77.18-11.91Zm158.718-218.22l.033.009Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Happy(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 16C123.452 16 16 123.452 16 256s107.452 240 240 240s240-107.452 240-240S388.548 16 256 16m147.078 387.078a207.253 207.253 0 1 1 44.589-66.125a207.332 207.332 0 0 1-44.589 66.125"/><path fill="currentColor" d="M152 200h40v40h-40zm168 0h40v40h-40zm18.289 107.2A83.6 83.6 0 0 1 260.3 360h-8.6a83.6 83.6 0 0 1-77.992-52.8l-1.279-3.2h-34.461L144 319.081A116 116 0 0 0 251.7 392h8.6A116 116 0 0 0 368 319.081L374.032 304h-34.464Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hd(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M364 136h-92v240h92a52.059 52.059 0 0 0 52-52V188a52.059 52.059 0 0 0-52-52m20 188a20.023 20.023 0 0 1-20 20h-60V168h60a20.023 20.023 0 0 1 20 20Zm-176-84h-80V136H96v240h32V272h80v104h32V136h-32z"/><path fill="currentColor" d="M464 16H48a32.036 32.036 0 0 0-32 32v416a32.036 32.036 0 0 0 32 32h416a32.036 32.036 0 0 0 32-32V48a32.036 32.036 0 0 0-32-32m0 448H48V48h416l.02 416Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hdr(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M192 104v288h96a32.036 32.036 0 0 0 32-32V136a32.036 32.036 0 0 0-32-32Zm96 256h-64V136h64Zm-128 32V104h-32v128H64V104H32v288h32V264h64v128zm320-160v-96a32.036 32.036 0 0 0-32-32h-96v288h32V264h29.167L440 331.081V392h32v-67.081L447.632 264H448a32.036 32.036 0 0 0 32-32m-96 0v-96h64v96Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Header(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 496h480V16H16ZM48 48h416v416H48Z"/><path fill="currentColor" d="M288 144h32v96H192v-96h32v-32H96v32h32v224H96v32h128v-32h-32v-96h128v96h-32v32h128v-32h-32V144h32v-32H288z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Headphones(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 216h-48v-32c0-79.4-64.6-144-144-144s-144 64.6-144 144v32H64a48.055 48.055 0 0 0-48 48v128a48.055 48.055 0 0 0 48 48h48a32.036 32.036 0 0 0 32-32V184a112 112 0 0 1 224 0v224a32.036 32.036 0 0 0 32 32h48a48.055 48.055 0 0 0 48-48V264a48.055 48.055 0 0 0-48-48M112 408H64a16.019 16.019 0 0 1-16-16V264a16.019 16.019 0 0 1 16-16h48v56h.008l.012 104Zm352-16a16.019 16.019 0 0 1-16 16h-48V248h48a16.019 16.019 0 0 1 16 16Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Healing(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m93.976 276.6l-55.765 55.767a32 32 0 0 0 0 45.255l96.167 96.168a32.038 32.038 0 0 0 45.254 0l55.768-55.766l-22.616-22.616l-55.764 55.767a.094.094 0 0 1-.015-.013L60.839 355l55.761-55.77Zm379.813-142.222L377.622 38.21a32.037 32.037 0 0 0-45.254 0L276.6 93.976l22.63 22.624L355 60.838l96.166 96.167l-55.766 55.767l22.624 22.628l55.765-55.764a32 32 0 0 0 0-45.255Zm-232.818 73.136l15.515-15.515L272 207.514l-15.514 15.515zm0 96.973l15.515-15.515L272 304.487L256.486 320zm48.484-48.488l15.515-15.515l15.514 15.515l-15.514 15.514zm-96.97.001l15.514-15.515L223.514 256l-15.515 15.515z"/><path fill="currentColor" d="M321.3 462.705a47.988 47.988 0 0 0 67.641-.23l73.539-73.539a48 48 0 0 0 0-67.882l-7.764-7.765l-240-240l-23.77-23.764a48 48 0 0 0-67.882 0l-73.54 73.539a48 48 0 0 0 0 67.882l31.764 31.765l239.766 239.764c.079.078.162.151.246.23M72.152 168.319a16 16 0 0 1 0-22.628l73.539-73.539a16 16 0 0 1 22.628 0l65.054 65.054l-96.167 96.167ZM159.833 256L256 159.833L352.171 256l-96.162 96.172Zm214.961 22.627l65.054 65.054a16 16 0 0 1 0 22.628l-73.539 73.539a16 16 0 0 1-22.628 0l-65.054-65.054Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heart(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M453.122 79.012a128 128 0 0 0-181.087.068l-15.511 15.7l-15.382-15.666l-.1-.1a128 128 0 0 0-181.02 0l-6.91 6.91a128 128 0 0 0 0 181.019l182.373 182.371l20.595 21.578l.491-.492l.533.533l19.296-20.359L460.032 266.94a128.147 128.147 0 0 0 0-181.019ZM437.4 244.313L256.571 425.146L75.738 244.313a96 96 0 0 1 0-135.764l6.911-6.91a96 96 0 0 1 135.713-.051l38.093 38.787l38.274-38.736a96 96 0 0 1 135.765 0l6.91 6.909a96.11 96.11 0 0 1-.004 135.765"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Highlighter(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M398.789 22.31a31.762 31.762 0 0 0-22.771-9.52H376a31.769 31.769 0 0 0-22.783 9.552L87.534 292.234a32.086 32.086 0 0 0 .177 45.076l14.7 14.7L16 439.427V478h106.8l52.8-52.8l12.479 12.48a32 32 0 0 0 46-.77l258.234-276.14a31.913 31.913 0 0 0-.6-44.339ZM109.548 446H54.5l46.552-47.1l27.8 27.8Zm101.16-30.946L175.6 379.947l-24.127 24.127l-27.932-27.932l23.986-24.269l-37.191-37.189l48.338-49.105L257.8 364.7Zm68.958-73.74l-98.541-98.54L376.017 44.791l92.923 94.121Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Highligt(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m104 305.5l88 48V480h128V353.5l88-48V160H104ZM136 192h240v94.5l-88 48V448h-64V334.5l-88-48ZM240 32h32v72h-32zM61.775 72.403l22.627-22.628l50.912 50.912l-22.628 22.627zm314.91 28.281l50.912-50.911L450.224 72.4l-50.912 50.912z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func History(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256.25 16A240 240 0 0 0 88 84.977V16H56v128h128v-32h-77.713A208 208 0 0 1 256.25 48C370.8 48 464 141.2 464 255.75S370.8 463.5 256.25 463.5S48.5 370.3 48.5 255.75h-32a239.75 239.75 0 0 0 409.279 169.529A239.75 239.75 0 0 0 256.25 16"/><path fill="currentColor" d="M240 111.951L239.465 288H368v-32h-96.437L272 112.049z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Home(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469.666 216.45L271.078 33.749a34 34 0 0 0-47.062.98L41.373 217.373L32 226.745V496h176V328h96v168h176V225.958ZM248.038 56.771c.282 0 .108.061-.013.18c-.125-.119-.269-.18.013-.18M448 464H336V328a32 32 0 0 0-32-32h-96a32 32 0 0 0-32 32v136H64V240L248.038 57.356c.013-.012.014-.023.024-.035L448 240Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hospital(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M160 304h32v32h-32zm80 0h32v32h-32zm80 0h32v32h-32z"/><path fill="currentColor" d="M440 464V88h-72V16H144v72H72v376H16v32h480v-32ZM176 48h160v160H176Zm232 416H272v-64h-32v64H104V120h40v120h224V120h40Z"/><path fill="currentColor" d="M272 80h-32v32h-32v32h32v32h32v-32h32v-32h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HotTub(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="111.485" cy="141.912" r="34.088" fill="currentColor"/><path fill="currentColor" d="M192 296h-3.041l-34.909-96H82.287v96H16v152a48.054 48.054 0 0 0 48 48h384a48.054 48.054 0 0 0 48-48V296Zm272 152a16.019 16.019 0 0 1-16 16H64a16.019 16.019 0 0 1-16-16V328h416Z"/><path fill="currentColor" d="M296 216a20.376 20.376 0 0 1 12.362-18.748l11.881-5.091a52.4 52.4 0 0 0 0-96.321l-11.881-5.092A20.376 20.376 0 0 1 296 72V32h-32v40a52.336 52.336 0 0 0 31.757 48.16l11.881 5.092a20.4 20.4 0 0 1 0 37.5l-11.881 5.091A52.338 52.338 0 0 0 264 216v48h32Zm112 0a20.376 20.376 0 0 1 12.362-18.748l11.881-5.091a52.4 52.4 0 0 0 0-96.321l-11.881-5.092A20.376 20.376 0 0 1 408 72V32h-32v40a52.336 52.336 0 0 0 31.757 48.16l11.881 5.092a20.4 20.4 0 0 1 0 37.5l-11.881 5.091A52.338 52.338 0 0 0 376 216v48h32ZM96 360h32v72H96zm96 0h32v72h-32zm96 0h32v72h-32zm96 0h32v72h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func House(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M408 406.545V248H288v158.545ZM320 280h56v94.545h-56Z"/><path fill="currentColor" d="M271.078 33.749a34 34 0 0 0-47.066.984L32 226.745V496h112V336h64v160h272V225.958ZM448 464H240V304H112v160H64V240L249.412 57.356V57.3L448 240Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Https(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M400 200h-12v-56a128 128 0 0 0-256 0v56h-12a24.028 24.028 0 0 0-24 24v248a24.028 24.028 0 0 0 24 24h280a24.028 24.028 0 0 0 24-24V224a24.028 24.028 0 0 0-24-24m-236-56a96 96 0 0 1 192 0v56H164Zm228 320H128V232h264Z"/><path fill="currentColor" d="M240 328h40v40h-40z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Image(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M40 472h432V40H40Zm400-123.858L328.628 236.769l46.6-46.6L440 254.935ZM72 72h368v137.68l-64.769-64.77L306 214.142l-100-100l-134 134Zm0 221.4l134-134l234 234V440H72Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageBroken(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M40 472h432V312h-32v128H72V312H40zm0-206.245l49.373 25.437l53.82-46.829l56.159 50.528L256 247.052l56.648 47.839l56.159-50.528l53.82 46.829L472 265.755V40H40ZM72 72h368v174.244l-12.738 6.564l-58.809-51.171l-56.471 50.806L256 205.167l-55.982 47.276l-56.471-50.806l-58.809 51.171L72 246.244Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageOne(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M1.875 1.875v20.25h20.25V1.875zm18.75 1.5v6.454l-3.036-3.036l-3.245 3.245L9.656 5.35l-6.281 6.281V3.374zm0 12.944l-5.221-5.221l2.185-2.185l3.036 3.036zm-17.25 4.306v-6.872l6.281-6.281l10.969 10.969v2.185z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImagePlus(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M336 72V40H40v432h432V184h-32v25.68l-64.769-64.77L306 214.142l-100-100l-134 134V72Zm39.231 118.166L440 254.935v93.207L328.628 236.769ZM206 159.4l234 234V440H72V293.4Z"/><path fill="currentColor" d="M448 16h-32v48h-48v32h48v48h32V96h48V64h-48z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Inbox(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M32 16v480h448V16Zm416 288H342.111l-32 64H201.889l-32-64H64V144h384Zm0-256v64H64V48ZM64 464V336h86.111l32 64h147.778l32-64H448v128Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IndentDecrease(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M72 63.998h424v32H72zm128 88h296v32H200zm0 88h296v32H200zm0 88h296v32H200zm-128 88h424v32H72zm88-271.089L4.473 256L160 367.091Zm-32 160L59.527 256L128 207.091Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IndentIncrease(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M72 63.998h424v32H72zm128 88h296v32H200zm0 88h296v32H200zm0 88h296v32H200zm-128 88h424v32H72zM16 144.909v222.182L171.527 256Zm32 62.182L116.473 256L48 304.909Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Industry(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M459.26 96L328 225.7V96h-34.525L168 223.267V16H16v480h480V96ZM464 464H48V48h88v216h36.778L296 139.018V264h38.764L464 136.3Z"/><path fill="currentColor" d="M136 328v8h32v-32h-32zm0 48h32v32h-32zm80-48v8h32v-32h-32zm0 48h32v32h-32zm80-48v8h32v-32h-32zm0 48h32v32h-32zm80-72h32v32h-32zm0 72h32v32h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IndustrySlash(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M459.26 96L328 225.697V96h-34.55l-64.44 65.128l22.628 22.628L296 138.92v89.198l37.314 37.315L464 136.303v259.815l32 32V96zM168 16H83.882L168 100.118zm-32 288h32v32h-32zm0 72h32v32h-32zm80-72h32v32h-32zm0 72h32v32h-32zm80 0h32v32h-32z"/><path fill="currentColor" d="M38.627 16H16v480h480v-22.627ZM48 464V70.627l88 88V248h32v-57.373L441.373 464Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Infinity(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M463.2 176.805a112 112 0 0 0-158.39 0l-48.57 48.568l-48.573-48.573a112 112 0 1 0 0 158.392l48.568-48.569l48.57 48.569A112 112 0 1 0 463.2 176.805M185.04 312.569a80 80 0 1 1 0-113.138l55.2 55.2v2.746Zm255.528 0a80 80 0 0 1-113.136 0l-55.2-55.2v-2.744l55.2-55.2a80 80 0 1 1 113.136 113.144"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Info(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 95.998h34.924v34.924H256z"/><path fill="currentColor" d="M16 496h480V16H16ZM48 48h416v416H48Z"/><path fill="currentColor" d="M285.313 359.032a18.123 18.123 0 0 1-15.6 8.966a18.061 18.061 0 0 1-17.327-23.157l35.67-121.277a49.577 49.577 0 0 0-93.356-32.992l-11.718 28.234l29.557 12.266l11.718-28.235a17.577 17.577 0 0 1 33.1 11.7l-35.67 121.277A50.061 50.061 0 0 0 269.709 400a50.227 50.227 0 0 0 43.25-24.853l15.1-25.913l-27.646-16.115Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Input(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M472 96H40a24.028 24.028 0 0 0-24 24v80h32v-72h416v256H48v-72H16v80a24.028 24.028 0 0 0 24 24h432a24.028 24.028 0 0 0 24-24V120a24.028 24.028 0 0 0-24-24"/><path fill="currentColor" d="m212.687 323.078l22.626 22.627l90.511-90.509l-90.511-90.51l-22.626 22.628l51.881 51.882H16v31.999h248.569z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InputHdmi(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M112 16v104H80v131.492l72 157.091V496h208v-87.417l72-157.091V120h-32V16Zm32 32h224v72h-32V80h-32v40h-32V80h-32v40h-32V80h-32v40h-32Zm256 196.508L328 401.6V464H184v-62.4l-72-157.092V152h288Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InputPower(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M152 16v104H88v191.005l104 96v89h128v-89l104-96V120h-64V16h-32v104H184V16Zm240 136v145l-104 96v71h-64v-71l-104-96V152Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Institution(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M247.759 14.358L16 125.946V184h480v-58.362ZM464 152H48v-5.946l200.241-96.412L464 146.362ZM16 496h480V392H16Zm32-72h416v40H48Zm24-216h32v160H72zm336 0h32v160h-32zm-224 0h32v160h-32zm112 0h32v160h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Italic(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 496h480V16H16ZM48 48h416v416H48Z"/><path fill="currentColor" d="M200 143.998h44.442l-42 224H152v32h160v-32h-44.442l42-224H360v-32H200z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func JustifyCenter(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 63.998h480v32H16zm0 88h480v32H16zm0 88h480v32H16zm0 88h480v32H16zm80 88h320v32H96z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func JustifyLeft(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 63.998h480v32H16zm0 88h480v32H16zm0 88h480v32H16zm0 88h480v32H16zm0 88h320v32H16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func JustifyRight(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 63.998h480v32H16zm0 88h480v32H16zm0 88h480v32H16zm0 88h480v32H16zm160 88h320v32H176z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Keyboard(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M472 80H40a24.028 24.028 0 0 0-24 24v240a24.028 24.028 0 0 0 24 24h432a24.028 24.028 0 0 0 24-24V104a24.028 24.028 0 0 0-24-24m-8 256H48V112h416Z"/><path fill="currentColor" d="M144 272h160v32H144zm-64 0h32v32H80zm320 0h32v32h-32zm-64 0h32v32h-32zm32-64h32v32h-32zm-64 0h32v32h-32zm-64 0h32v32h-32zm-64 0h32v32h-32zm-64 0h32v32h-32zm224-64h32v32h-32zm64 0h32v32h-32zm-128 0h32v32h-32zm-64 0h32v32h-32zm-64 0h32v32h-32zm-64 0h32v32H80z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lan(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M496 272v-32H272v-48h56a24.027 24.027 0 0 0 24-24V40a24.027 24.027 0 0 0-24-24H184a24.027 24.027 0 0 0-24 24v128a24.027 24.027 0 0 0 24 24h56v48H16v32h80v48H41.391a24.028 24.028 0 0 0-24 24v128a24.028 24.028 0 0 0 24 24H184a24.027 24.027 0 0 0 24-24V344a24.027 24.027 0 0 0-24-24h-56v-48h256v48h-56a24.027 24.027 0 0 0-24 24v128a24.027 24.027 0 0 0 24 24h144a24.027 24.027 0 0 0 24-24V344a24.027 24.027 0 0 0-24-24h-56v-48ZM192 48h128v112H192Zm-16 416H49.391V352H176Zm288 0H336V352h128Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Language(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M320 24H16v304h176v168h304V192H320ZM148.305 96L98.093 239.3H132l8.166-23.3H192v80H48V56h240v136h-66.668L187.7 96Zm36.317 88h-33.244L168 136.562ZM464 224v240H224V224Z"/><path fill="currentColor" d="M317.432 368.48a136.761 136.761 0 0 0 10.089 14.12q-17.4 9.384-39.521 9.4v32c24.141 0 45.71-6.408 64-18.824C370.29 417.592 391.859 424 416 424v-32q-22.075 0-39.52-9.407a136.574 136.574 0 0 0 10.088-14.113A166.212 166.212 0 0 0 406.662 320H424v-32h-56v-24h-32v24h-56v32h17.338a166.212 166.212 0 0 0 20.094 48.48M373.53 320a133.013 133.013 0 0 1-14.1 31.52a104.39 104.39 0 0 1-7.43 10.448a103.546 103.546 0 0 1-6.93-9.651A132.384 132.384 0 0 1 330.466 320Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Laptop(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M440 368a24.028 24.028 0 0 0 24-24V104a24.028 24.028 0 0 0-24-24H72a24.028 24.028 0 0 0-24 24v240a24.028 24.028 0 0 0 24 24ZM80 112h352v224H80ZM16 400h480v32H16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Layers(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m487.938 162.108l-224-128a16 16 0 0 0-15.876 0l-224 128a16 16 0 0 0 .382 28l224 120a16 16 0 0 0 15.112 0l224-120a16 16 0 0 0 .382-28M256 277.849L65.039 175.548L256 66.428l190.961 109.12Z"/><path fill="currentColor" d="M263.711 394.02L480 275.061v-36.522L256 361.74L32 238.539v36.522L248.289 394.02a16.005 16.005 0 0 0 15.422 0"/><path fill="currentColor" d="m32 362.667l216.471 115.451a16 16 0 0 0 15.058 0L480 362.667V326.4L256 445.867L32 326.4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Leaf(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M389.053 126.3A302.909 302.909 0 0 0 280.012 18.15L272 13.516l-8.012 4.634A301.084 301.084 0 0 0 113.4 279.042c0 3.445.06 6.944.177 10.4c1.592 46.856 19.511 86.283 51.82 114.018c24.724 21.225 56.438 34.182 90.607 37.273V496h32V240H256v168.528c-54.064-6.263-107.873-44.455-110.444-120.174c-.106-3.095-.16-6.228-.16-9.312A270.286 270.286 0 0 1 272 50.673a270.286 270.286 0 0 1 126.6 228.369c0 3.084-.054 6.217-.16 9.313c-2.056 60.573-36.907 97.127-78.444 112.536v33.867a156.188 156.188 0 0 0 58.607-31.3c32.309-27.735 50.228-67.162 51.82-114.017c.117-3.456.177-6.955.177-10.4A300.948 300.948 0 0 0 389.053 126.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lemon(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M466.278 47.947a85.883 85.883 0 0 0-115.124-5.72a196.7 196.7 0 0 0-65.642-11.176A251.264 251.264 0 0 0 34.248 282.316a196.628 196.628 0 0 0 11.176 65.643A85.872 85.872 0 0 0 166.266 468.8a196.664 196.664 0 0 0 65.644 11.177a251.266 251.266 0 0 0 251.264-251.263A196.639 196.639 0 0 0 472 163.065a85.982 85.982 0 0 0-5.722-115.118M438.7 165.563a164.674 164.674 0 0 1 12.471 63.151c0 120.9-98.361 219.264-219.264 219.264a164.7 164.7 0 0 1-63.151-12.471l-10.33-4.263l-7.559 8.23c-.3.331-.607.661-.928.981a53.861 53.861 0 1 1-76.171-76.171c.32-.321.65-.625.981-.93l8.229-7.558l-4.262-10.329a164.713 164.713 0 0 1-12.472-63.152c0-120.9 98.362-219.263 219.265-219.263a164.709 164.709 0 0 1 63.15 12.47l10.324 4.261l7.56-8.222c.3-.326.6-.653.933-.986a53.862 53.862 0 0 1 76.172 76.172c-.324.323-.659.631-.993.939l-8.214 7.559Z"/><path fill="currentColor" d="m120.32 226.743l28.471 14.611A218.891 218.891 0 0 1 269.44 135.047l-10.88-30.094a250.8 250.8 0 0 0-138.24 121.79"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LevelDown(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m171.313 348.686l-22.626 22.628L272 494.627l123.313-123.313l-22.626-22.628L288 433.373V96H72v32h184v305.373z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LevelUp(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M80 384v32h216V78.627l84.687 84.687l22.626-22.628L280 17.373L156.687 140.686l22.626 22.628L264 78.627V384z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Library(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M472 16H168a24 24 0 0 0-24 24v304a24 24 0 0 0 24 24h304a24 24 0 0 0 24-24V40a24 24 0 0 0-24-24m-8 320H176V48h288Z"/><path fill="currentColor" d="M112 400V80H80v328a24 24 0 0 0 24 24h328v-32Z"/><path fill="currentColor" d="M48 464V144H16v328a24 24 0 0 0 24 24h328v-32Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LibraryAdd(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M48 464V144H16v328a24.027 24.027 0 0 0 24 24h328v-32z"/><path fill="currentColor" d="M144 400h-32V80H80v328a24.027 24.027 0 0 0 24 24h328v-32z"/><path fill="currentColor" d="M472 16H168a24.027 24.027 0 0 0-24 24v304a24.027 24.027 0 0 0 24 24h304a24.027 24.027 0 0 0 24-24V40a24.027 24.027 0 0 0-24-24m-8 320H176V48h288Z"/><path fill="currentColor" d="M304 288h32v-84h84v-32h-84V88h-32v84h-84v32h84z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LibraryBuilding(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M247.759 14.358L16 125.946V184h480v-58.362ZM464 152H48v-5.946l200.241-96.412L464 146.362ZM16 496h480V392H16Zm32-72h416v40H48Zm24-216h32v160H72zm336 0h32v160h-32zm-224 0h32v160h-32zm112 0h32v160h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LifeRing(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M425.706 86.294A240 240 0 0 0 86.294 425.706A240 240 0 0 0 425.706 86.294m-80.04 236.745a111.781 111.781 0 0 0 0-134.078l68.218-68.218a207.579 207.579 0 0 1 0 270.514ZM176 256a80 80 0 1 1 80 80a80.091 80.091 0 0 1-80-80M391.257 98.116l-68.218 68.218a111.781 111.781 0 0 0-134.078 0l-68.218-68.218a207.579 207.579 0 0 1 270.514 0M98.116 120.743l68.218 68.218a111.781 111.781 0 0 0 0 134.078l-68.218 68.218a207.579 207.579 0 0 1 0-270.514m22.627 293.141l68.218-68.218a111.781 111.781 0 0 0 134.078 0l68.218 68.218a207.579 207.579 0 0 1-270.514 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lightbulb(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M252.592 496h6.816a108.122 108.122 0 0 0 108-108v-55.692a177.481 177.481 0 0 0 66.37-138.531C433.778 95.751 354.027 16 256 16S78.222 95.751 78.222 193.777a177.477 177.477 0 0 0 66.371 138.531V388a108.121 108.121 0 0 0 107.999 108m6.816-32h-6.816a76.106 76.106 0 0 1-70.631-48H330.04a76.107 76.107 0 0 1-70.632 48m76-80H176.593v-40h158.815ZM110.222 193.777C110.222 113.4 175.618 48 256 48s145.778 65.4 145.778 145.777a146.392 146.392 0 0 1-59.817 117.737l-.661.486H170.7l-.665-.486a146.394 146.394 0 0 1-59.813-117.737"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LineSpacing(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M240 48h256v32H240zm0 128h256v32H240zm0 128h256v32H240zm0 128h256v32H240zM24 368v30.19l88.937 97.728L200 398.089V368h-72.8V144H200v-30.19l-88.937-97.728L24 113.911V144h71.2v224Zm44.538-256l42.791-48.082L155.086 112Zm86.924 288l-42.791 48.082L68.914 400Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LineStyle(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 120h480V16H16Zm32-72h416v40H48Zm192 96H16v96h224Zm-32 64H48v-32h160Zm64 32h224v-96H272Zm32-64h160v32H304Zm-152 96H16v88h136Zm-32 56H48v-24h72Zm72 32h128v-88H192Zm32-56h64v24h-64Zm136 56h136v-88H360Zm32-56h72v24h-72ZM121.442 443.278A52.721 52.721 0 1 0 68.722 496a52.78 52.78 0 0 0 52.72-52.722m-73.442 0A20.721 20.721 0 1 1 68.722 464A20.745 20.745 0 0 1 48 443.278m92.853 0a52.721 52.721 0 1 0 52.721-52.72a52.78 52.78 0 0 0-52.721 52.72m73.442 0a20.721 20.721 0 1 1-20.721-20.72a20.745 20.745 0 0 1 20.726 20.72Zm51.41 0a52.721 52.721 0 1 0 52.721-52.72a52.78 52.78 0 0 0-52.721 52.72m73.442 0a20.721 20.721 0 1 1-20.721-20.72a20.745 20.745 0 0 1 20.721 20.72m104.131-52.72A52.721 52.721 0 1 0 496 443.278a52.78 52.78 0 0 0-52.722-52.72m0 73.442A20.721 20.721 0 1 1 464 443.278A20.745 20.745 0 0 1 443.278 464"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LineWeight(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 464h480v32H16zm0-32h480v-80H16Zm32-48h416v16H48Zm-32-64h480V208H16Zm32-80h416v48H48Zm-32-64h480V16H16ZM48 48h416v96H48Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Link(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M457.47 55.833c-53.026-53.026-139.307-53.026-192.332 0L168.971 152l22.629 22.627l96.165-96.167a104 104 0 0 1 147.078 147.079l-96.167 96.167l22.624 22.627l96.167-96.167C510.5 195.14 510.5 108.86 457.47 55.833m-231.931 379.01a104 104 0 0 1-147.078 0a104 104 0 0 1 0-147.078l90.511-90.511l-22.627-22.627l-90.512 90.511A136 136 0 1 0 248.166 457.47l90.51-90.51l-22.627-22.627Z"/><path fill="currentColor" d="m129.373 361.303l226.274-226.275l22.628 22.628L152 383.93z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkAlt(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M48 256a80.091 80.091 0 0 1 80-80h112v-32H128a112 112 0 0 0 0 224h112v-32H128a80.091 80.091 0 0 1-80-80m336-112H272v32h112a80 80 0 0 1 0 160H272v32h112a112 112 0 0 0 0-224"/><path fill="currentColor" d="M140 240.652h232v32H140z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkBroken(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M288 496h32V368h128v-32H288zm-64-320V16h-32v128H64v32zM78.708 434.573l-1.279-1.273a100.478 100.478 0 0 1 0-142.1l75.2-75.2h-45.257L54.8 268.57a132.478 132.478 0 0 0 0 187.35l1.278 1.278a132.628 132.628 0 0 0 187.352 0l4.57-4.57v-45.255l-27.2 27.2a100.591 100.591 0 0 1-142.092 0M457.2 56.08l-1.278-1.28c-51.653-51.655-135.7-51.653-187.352 0L264 59.372v45.255l27.2-27.2a100.591 100.591 0 0 1 142.095 0l1.279 1.278a100.478 100.478 0 0 1 0 142.1l-75.2 75.2h45.253L457.2 243.43a132.478 132.478 0 0 0 0-187.35"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func List(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M208 80h264v32H208zM40 96a64 64 0 1 0 64-64a64.072 64.072 0 0 0-64 64m64-32a32 32 0 1 1-32 32a32.036 32.036 0 0 1 32-32m104 176h264v32H208zm-104 80a64 64 0 1 0-64-64a64.072 64.072 0 0 0 64 64m0-96a32 32 0 1 1-32 32a32.036 32.036 0 0 1 32-32m104 176h264v32H208zm-104 80a64 64 0 1 0-64-64a64.072 64.072 0 0 0 64 64m0-96a32 32 0 1 1-32 32a32.036 32.036 0 0 1 32-32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListFilter(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M96 197.333h320v32H96zm72 101.334h176v32H168zM216 400h80v32h-80zM48 96h416v32H48z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListHighPriority(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 416h240v32H256zm0-106.667h240v32H256zm0-106.666h240v32H256zM328 96h168v32H328z"/><path fill="currentColor" d="M302 111L167.2 27.216V96h-5.965C121.783 96 84.91 114.755 57.4 148.81C30.7 181.866 16 225.616 16 272s14.7 90.134 41.4 123.19C84.91 429.245 121.783 448 161.231 448H216v-32h-54.769C98.8 416 48 351.4 48 272s50.8-144 113.231-144h5.969v69.228ZM199.2 84.784l42.8 26.607l-42.8 27.381Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListLowPriority(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 56h240v32H256zm0 106.667h240v32H256zm0 106.666h240v32H256zM328 376h168v32H328z"/><path fill="currentColor" d="M161.231 408h5.969v68.783L302 393l-134.8-86.228V376h-5.965C98.8 376 48 311.4 48 232S98.8 88 161.231 88H216V56h-54.769C121.783 56 84.91 74.755 57.4 108.81C30.7 141.866 16 185.616 16 232s14.7 90.134 41.4 123.19C84.91 389.245 121.783 408 161.231 408m37.969-42.772l42.8 27.381l-42.8 26.608Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListNumbered(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M184 80h288v32H184zm0 160h288v32H184zm0 160h288v32H184zm-64-240V40H56v32h32v88zM56 262.111V312h80v-32H91.777L136 257.889V192H56v32h48v14.111zM56 440v32h80V344H56v32h48v16H80v32h24v16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListNumberedRtl(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M40 80h288v32H40zm0 160h288v32H40zm0 160h288v32H40zm400-240V40h-64v32h32v88zm-64 102.111V312h80v-32h-44.223L456 257.889V192h-80v32h48v14.111zM376 440v32h80V344h-80v32h48v16h-24v32h24v16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListRich(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M136 24H16v120h120Zm-32 88H48V56h56Zm32 88H16v120h120Zm-32 88H48v-56h56Zm32 88H16v120h120Zm-32 88H48v-56h56Zm72-440.002h320v32H176zm0 88h256v32H176zm0 88h320v32H176zm0 88h256v32H176zm0 176h256v32H176zm0-88h320v32H176z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationPin(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M253.924 127.592a64 64 0 1 0 64 64a64.073 64.073 0 0 0-64-64m0 96a32 32 0 1 1 32-32a32.037 32.037 0 0 1-32 32"/><path fill="currentColor" d="M376.906 68.515A173.922 173.922 0 0 0 108.2 286.426l120.907 185.613a29.619 29.619 0 0 0 49.635 0l120.911-185.613a173.921 173.921 0 0 0-22.747-217.911m-4.065 200.444l-118.916 182.55l-118.917-182.55c-36.4-55.879-28.593-130.659 18.563-177.817a141.92 141.92 0 0 1 200.708 0c47.156 47.158 54.962 121.938 18.562 177.817"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockLocked(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 200v-56a128 128 0 0 0-256 0v56H88v128c0 92.635 75.364 168 168 168s168-75.365 168-168V200Zm-224-56a96 96 0 0 1 192 0v56H160Zm232 184c0 74.99-61.01 136-136 136s-136-61.01-136-136v-96h272Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockUnlocked(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 200v-56a128 128 0 0 0-217.582-91.43l22.4 22.855A96 96 0 0 1 352 144v56H88v128c0 92.636 75.364 168 168 168s168-75.364 168-168V200Zm8 128c0 74.99-61.009 136-136 136s-136-61.01-136-136v-96h272Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Locomotive(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M328 427.127c0-.175-.012-.347-.013-.522l14.29-.357c0 .294-.022.584-.022.879a52.873 52.873 0 1 0 105.745 0c0-1.051-.04-2.092-.1-3.127H496v-50.9l-24-64v-65.323a47.736 47.736 0 0 0-26.534-42.932L424 190.111V96H312v88h-68.468l-29.059-87.18A47.941 47.941 0 0 0 168.936 64H16v82.832l40 16v50.334l-40 16V432h86.488a52.866 52.866 0 0 0 105.449-2.393l14.372-.36A52.867 52.867 0 0 0 328 427.127M395.127 448A20.873 20.873 0 1 1 416 427.127A20.9 20.9 0 0 1 395.127 448m-240 0A20.873 20.873 0 1 1 176 427.127A20.9 20.9 0 0 1 155.127 448m76.586-51l-32.607.815a52.83 52.83 0 0 0-89.34 2.185H48V250.833l40-16v-93.665l-40-16V96h120.936a15.979 15.979 0 0 1 15.179 10.94L220.468 216H344v-88h48v81.889l39.155 19.577A15.915 15.915 0 0 1 440 243.777V296H96v32h348.912L464 378.9V392h-29.4a52.78 52.78 0 0 0-80.594 1.945l-37.028.926A52.794 52.794 0 0 0 231.713 397m43.414 51A20.873 20.873 0 1 1 296 427.127A20.9 20.9 0 0 1 275.127 448"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Loop(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 210.511V264a112.127 112.127 0 0 1-112 112H78.627l44.686-44.687l-22.626-22.626L56 353.373l-4.415 4.414l-33.566 33.567l74.022 83.276l23.918-21.26L75.63 408H352c79.4 0 144-64.6 144-144v-85.489Z"/><path fill="currentColor" d="M48 256a112.127 112.127 0 0 1 112-112h273.373l-44.686 44.687l22.626 22.626L456 166.627l4.117-4.116l33.864-33.865l-74.022-83.276l-23.918 21.26L436.37 112H160c-79.4 0-144 64.6-144 144v85.787l32-32Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LoopCircular(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M415.313 358.7c36.453-36.452 55.906-85.231 54.779-137.353c-1.112-51.375-21.964-99.908-58.715-136.66l-22.627 22.627a166.816 166.816 0 0 1 49.35 114.725c.937 43.313-15.191 83.81-45.463 114.083l-48.617 49.051l.044-89.165l-32-.016L311.992 440h144.071v-32h-89.614ZM47.937 112h89.614l-48.864 49.3c-36.453 36.451-55.906 85.231-54.779 137.352a198.676 198.676 0 0 0 58.715 136.66l22.627-22.627A166.818 166.818 0 0 1 65.9 297.962c-.937-43.314 15.191-83.811 45.463-114.083l48.617-49.051l-.044 89.165l32 .015L192.008 80H47.937Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LoopOne(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M208 312v32h112v-32h-40V176h-32v24h-32v32h32v80z"/><path fill="currentColor" d="M464 210.511V264a112.127 112.127 0 0 1-112 112H78.627l44.686-44.687l-22.626-22.626L56 353.373l-4.415 4.414l-33.566 33.567l74.022 83.276l23.918-21.26L75.63 408H352c79.4 0 144-64.6 144-144v-85.489Z"/><path fill="currentColor" d="M48 256a112.127 112.127 0 0 1 112-112h273.373l-44.686 44.687l22.626 22.626L456 166.627l4.117-4.116l33.864-33.865l-74.022-83.276l-23.918 21.26L436.37 112H160c-79.4 0-144 64.6-144 144v85.787l32-32Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LowVision(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m397.222 139.1l-.218-.223c-61.693-61.692-154.082-74.423-228.455-38.211l24.213 24.213c60.468-24.573 132.5-12.4 181.509 36.52L464 254.683v3.107l-67.821 70.51l22.63 22.63L496 270.683V241.79Z"/><path fill="currentColor" d="M358.99 290.323A103.984 103.984 0 0 0 221.677 153.01l25.09 25.09a71.974 71.974 0 0 1 87.133 87.133ZM16 16v21.94l99.977 99.978c-.326.321-.656.636-.981.96L16 241.79v28.893l98.778 102.689l.218.222a199.715 199.715 0 0 0 257.5 20.84L474.06 496H496v-23.313L39.313 16Zm65.982 203.355l17.036-17.71L293.14 395.767a168.457 168.457 0 0 1-30.727 4.018ZM48 257.79v-3.107l11.794-12.261l151.439 151.439a166.38 166.38 0 0 1-73.5-42.788Zm278.879 126.462L121.206 178.578l16.523-17.178c.29-.289.586-.567.877-.854l210.9 210.9a166.053 166.053 0 0 1-22.627 12.806"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagnifyingGlass(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m479.6 399.716l-81.084-81.084l-62.368-25.767A175.014 175.014 0 0 0 368 192c0-97.047-78.953-176-176-176S16 94.953 16 192s78.953 176 176 176a175.034 175.034 0 0 0 101.619-32.377l25.7 62.2l81.081 81.088a56 56 0 1 0 79.2-79.195M48 192c0-79.4 64.6-144 144-144s144 64.6 144 144s-64.6 144-144 144S48 271.4 48 192m408.971 264.284a24.028 24.028 0 0 1-33.942 0l-76.572-76.572l-23.894-57.835l57.837 23.894l76.573 76.572a24.028 24.028 0 0 1-.002 33.941"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Map(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M478.465 89.022L329.6 47.382L180.3 89.438L41.459 50.052A20 20 0 0 0 16 69.293v340.6a24.093 24.093 0 0 0 17.449 23.089l146.817 41.65l149.365-42.074l140.983 39.436A20 20 0 0 0 496 452.728V112.135a24.08 24.08 0 0 0-17.535-23.113M163 436.466L48 403.842V85.17l115 32.624Zm150.615-32.647L195 437.231V118.542L313.615 85.13ZM464 436.91L345.615 403.8V85.089L464 118.2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaEject(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M97.088 286.56h317.824a24 24 0 0 0 18.432-39.371L274.433 56.63a24 24 0 0 0-36.866 0L78.656 247.189a24 24 0 0 0 18.432 39.371M256 84.491L397.824 254.56H114.176ZM416 328H96a24.027 24.027 0 0 0-24 24v64a24.027 24.027 0 0 0 24 24h320a24.027 24.027 0 0 0 24-24v-64a24.027 24.027 0 0 0-24-24m-8 80H104v-48h304Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaPause(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M200 48H72a24.028 24.028 0 0 0-24 24v368a24.028 24.028 0 0 0 24 24h128a24.028 24.028 0 0 0 24-24V72a24.028 24.028 0 0 0-24-24m-8 384H80V80h112ZM440 48H312a24.028 24.028 0 0 0-24 24v368a24.028 24.028 0 0 0 24 24h128a24.028 24.028 0 0 0 24-24V72a24.028 24.028 0 0 0-24-24m-8 384H320V80h112Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaPlay(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M444.4 235.236L132.275 49.449A24 24 0 0 0 96 70.072v364.142a24.017 24.017 0 0 0 35.907 20.839L444.03 276.7a24 24 0 0 0 .367-41.461ZM128 420.429V84.144l288.244 171.574Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaRecord(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 72C154.542 72 72 154.542 72 256s82.542 184 184 184s184-82.542 184-184S357.458 72 256 72m0 336c-83.813 0-152-68.187-152-152s68.187-152 152-152s152 68.187 152 152s-68.187 152-152 152"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaSkipBackward(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M455.979 424.271A24.053 24.053 0 0 0 480 400.251V112.015a24 24 0 0 0-38.285-19.286L264 224.369V112.015a24 24 0 0 0-38.285-19.286L31.155 236.847a24 24 0 0 0 0 38.57l194.56 144.119A24 24 0 0 0 264 400.251V287.9l177.715 131.637a23.922 23.922 0 0 0 14.264 4.734M232 384.37L58.88 256.132L232 127.9ZM448 127.9v256.47L274.88 256.132Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaSkipForward(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M32 111.882v288.236A23.979 23.979 0 0 0 70.285 419.4L248 287.763v112.355a23.979 23.979 0 0 0 38.285 19.282l194.56-144.119a24 24 0 0 0 0-38.57L286.285 92.6A24 24 0 0 0 248 111.882v112.355L70.285 92.6A24 24 0 0 0 32 111.882m248 15.881L453.119 256L280 384.237Zm-216 0L237.119 256L64 384.237Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaStepBackward(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M402.8 90.58a23.876 23.876 0 0 0-25.082 2.149L183.155 236.847a24 24 0 0 0 0 38.57l194.56 144.119A24 24 0 0 0 416 400.251V112.015a23.882 23.882 0 0 0-13.2-21.435M384 384.37L210.881 256.133L384 127.9ZM96 88h32v336H96z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaStepForward(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M328.845 236.582L134.285 92.463A24 24 0 0 0 96 111.749v288.236a23.979 23.979 0 0 0 38.285 19.286l194.56-144.118a24 24 0 0 0 0-38.57ZM128 384.1V127.63l173.119 128.237ZM384 88h32v336h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaStop(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M408 80H104a24.027 24.027 0 0 0-24 24v304a24.027 24.027 0 0 0 24 24h304a24.027 24.027 0 0 0 24-24V104a24.027 24.027 0 0 0-24-24m-8 320H112V112h288Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MedicalCross(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M344 16H168v152H16v176h152v152h176V344h152V168H344Zm120 184v112H312v152H200V312H48V200h152V48h112v152Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Meh(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 16C123.452 16 16 123.452 16 256s107.452 240 240 240s240-107.452 240-240S388.548 16 256 16m147.078 387.078a207.253 207.253 0 1 1 44.589-66.125a207.332 207.332 0 0 1-44.589 66.125"/><path fill="currentColor" d="M152 200h40v40h-40zm168 0h40v40h-40zM168 320h176v32H168z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Memory(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M192 320h128V192H192Zm32-96h64v64h-64Z"/><path fill="currentColor" d="M32 288v32h72v88h88v72h32v-72h64v72h32v-72h88v-88h72v-32h-72v-64h72v-32h-72v-88h-88V32h-32v72h-64V32h-32v72h-88v88H32v32h72v64Zm104-152h240v240H136Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Menu(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M80 96h352v32H80zm0 144h352v32H80zm0 144h352v32H80z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mic(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 328a96.108 96.108 0 0 0 96-96V112a96 96 0 0 0-192 0v120a96.108 96.108 0 0 0 96 96m-64-216a64 64 0 0 1 128 0v120a64 64 0 0 1-128 0Z"/><path fill="currentColor" d="M400 176v56c0 79.4-64.6 144-144 144s-144-64.6-144-144v-56H80v56c0 91.653 70.424 167.154 160 175.265V496h32v-88.735c89.576-8.111 160-83.612 160-175.265v-56Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Microphone(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 328a96.108 96.108 0 0 0 96-96V112a96 96 0 0 0-192 0v120a96.108 96.108 0 0 0 96 96m-64-216a64 64 0 0 1 128 0v120a64 64 0 0 1-128 0Z"/><path fill="currentColor" d="M400 176v56c0 79.4-64.6 144-144 144s-144-64.6-144-144v-56H80v56c0 91.653 70.424 167.154 160 175.265V496h32v-88.735c89.576-8.111 160-83.612 160-175.265v-56Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Minus(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M72 240h368v32H72z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mobile(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M380 16H132a32.036 32.036 0 0 0-32 32v416a32.036 32.036 0 0 0 32 32h248a32.036 32.036 0 0 0 32-32V48a32.036 32.036 0 0 0-32-32m0 32v32H132V48Zm0 64l.011 224H132V112Zm0 352H132v-96h248.016v96Z"/><path fill="currentColor" d="M240 400h32v32h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MobileLandscape(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 96H48a32.036 32.036 0 0 0-32 32v256a32.036 32.036 0 0 0 32 32h416a32.036 32.036 0 0 0 32-32V128a32.036 32.036 0 0 0-32-32M48 384V128h48v256.018H48Zm80-256h256v256l-256 .013Zm336 256h-48V128h48Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Money(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M432 64H16v320h416Zm-32 288H48V96h352Z"/><path fill="currentColor" d="M464 144v272H96v32h400V144z"/><path fill="currentColor" d="M224 302.46c39.7 0 72-35.137 72-78.326s-32.3-78.326-72-78.326s-72 35.136-72 78.326s32.3 78.326 72 78.326m0-124.652c22.056 0 40 20.782 40 46.326s-17.944 46.326-40 46.326s-40-20.782-40-46.326s17.944-46.326 40-46.326M80 136h32v176H80zm256 0h32v176h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Monitor(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M472 16H40a24.028 24.028 0 0 0-24 24v336a24.028 24.028 0 0 0 24 24h200v64h-80v32h192v-32h-80v-64h200a24.028 24.028 0 0 0 24-24V40a24.028 24.028 0 0 0-24-24m-8 352H48v-96h416Zm0-128H48V48h416Z"/><path fill="currentColor" d="M400 304h32v32h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoodBad(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 16C123.452 16 16 123.452 16 256s107.452 240 240 240s240-107.452 240-240S388.548 16 256 16m147.078 387.078a207.253 207.253 0 1 1 44.589-66.125a207.332 207.332 0 0 1-44.589 66.125"/><path fill="currentColor" d="M152 200h40v40h-40zm168 0h40v40h-40zm-64 80a104 104 0 0 0-104 104h208a104 104 0 0 0-104-104"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoodGood(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 16C123.452 16 16 123.452 16 256s107.452 240 240 240s240-107.452 240-240S388.548 16 256 16m147.078 387.078a207.253 207.253 0 1 1 44.589-66.125a207.332 207.332 0 0 1-44.589 66.125"/><path fill="currentColor" d="M152 200h40v40h-40zm168 0h40v40h-40zm-64 184a104 104 0 0 0 104-104H152a104 104 0 0 0 104 104"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoodVeryBad(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 16C123.452 16 16 123.452 16 256s107.452 240 240 240s240-107.452 240-240S388.548 16 256 16m147.078 387.078a207.253 207.253 0 1 1 44.589-66.125a207.332 207.332 0 0 1-44.589 66.125"/><path fill="currentColor" d="M256 280a104 104 0 0 0-104 104h208a104 104 0 0 0-104-104m-138.63-92.04l21.261-23.917l72 64l-21.26 23.918zm178.672 45.411l64-72l23.918 21.26l-64 72z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoodVeryGood(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 16C123.452 16 16 123.452 16 256s107.452 240 240 240s240-107.452 240-240S388.548 16 256 16m147.078 387.078a207.253 207.253 0 1 1 44.589-66.125a207.332 207.332 0 0 1-44.589 66.125"/><path fill="currentColor" d="M256 384a104 104 0 0 0 104-104H152a104 104 0 0 0 104 104m-50.243-155.708l20.486-24.584L168 155.173l-58.243 48.535l20.486 24.584L168 196.827zm80-24.584l20.486 24.584L344 196.827l37.757 31.465l20.486-24.584L344 155.173z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Moon(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M268.279 496c-67.574 0-130.978-26.191-178.534-73.745S16 311.293 16 243.718A252.252 252.252 0 0 1 154.183 18.676a24.44 24.44 0 0 1 34.46 28.958a220.12 220.12 0 0 0 54.8 220.923A218.746 218.746 0 0 0 399.085 333.2a220.2 220.2 0 0 0 65.277-9.846a24.439 24.439 0 0 1 28.959 34.461A252.256 252.256 0 0 1 268.279 496M153.31 55.781A219.3 219.3 0 0 0 48 243.718C48 365.181 146.816 464 268.279 464a219.3 219.3 0 0 0 187.938-105.31a252.912 252.912 0 0 1-57.13 6.513a250.539 250.539 0 0 1-178.268-74.016a252.147 252.147 0 0 1-67.509-235.4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mouse(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 80a24.027 24.027 0 0 0 24-24V16h-32v32H264a24.027 24.027 0 0 0-24 24v40h-31.375A88.725 88.725 0 0 0 120 200.625v159.946C120 435.247 180.753 496 255.429 496h1.142C331.247 496 392 435.247 392 360.571V200.625A88.725 88.725 0 0 0 303.375 112H272V80ZM152 200.625A56.689 56.689 0 0 1 208.625 144H240v88h-88Zm208 159.946A103.545 103.545 0 0 1 256.571 464h-1.142A103.545 103.545 0 0 1 152 360.571V264h208ZM303.375 144A56.689 56.689 0 0 1 360 200.625V232h-88v-88Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MouthSlash(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M412.8 107.2c-24.067-32.09-47.141-46.4-74.814-46.4c-22.54 0-44.048 13.373-62.2 38.674a180.547 180.547 0 0 0-15.033 25.326h-9.506a180.547 180.547 0 0 0-15.036-25.33c-18.149-25.3-39.657-38.674-62.2-38.674a71.222 71.222 0 0 0-36.672 9.918L161.4 94.768a38.732 38.732 0 0 1 12.614-1.968c25.175 0 46.513 39.864 51.993 53.8l3.97 10.195h41.106l10.917-.024l3.988-10.146c5.49-13.965 26.828-53.829 52-53.829c11.874 0 27.007 3.992 49.213 33.6c18.589 24.786 65.674 80.813 82.549 100.8l-44.815 120.82q-.4.834-.819 1.663a115.963 115.963 0 0 1-2.73 5.078l23.309 23.309a149.913 149.913 0 0 0 8.023-14.038q.719-1.434 1.41-2.885l.306-.643L496 248.439v-39.757c-14.8-17.525-64.653-76.756-83.2-101.482"/><path fill="currentColor" d="m387.305 320.678l45.426-102.962L331.8 172.6l-7.785 8.808c-.224.254-18.778 20.877-47.754 28.228l25.168 25.164a142.383 142.383 0 0 0 37.883-23.8l51.288 22.934l-27.564 62.475ZM16 16v21.941l77.207 77.207C71.693 143.6 29.973 197.958 16 216.151v39.139l41.441 104.86l.432.994c.461.97.933 1.934 1.411 2.884a151.245 151.245 0 0 0 56.861 61.472a155.361 155.361 0 0 0 80.63 22.5h118.45a155.361 155.361 0 0 0 80.63-22.5q2.391-1.45 4.718-2.986L474.059 496H496v-23.313L39.313 16Zm175.98 197.921l87.261 87.261a807.527 807.527 0 0 0-125.816 6.192l-31.615-66.39l62.17-31.084Zm-29.2-29.2L79.422 226.4l55.192 115.9l11.649-1.665A775.645 775.645 0 0 1 313 334.937l64.233 64.233a123.347 123.347 0 0 1-62 16.83H196.775c-45.934 0-88.675-26.033-108.893-66.328q-.45-.9-.887-1.8L42.214 234.555c14.147-18.422 52.059-67.825 73.836-96.564Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Move(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m403.814 165.089l-22.627 22.627l51.882 51.882H272.402V78.932l51.882 51.881l22.628-22.626l-90.51-90.511l-90.51 90.511l22.628 22.626l51.882-51.881v160.666H78.932l51.882-51.882l-22.627-22.627l-90.51 90.509l90.509 90.509l22.628-22.627l-51.883-51.882h161.471v161.47l-51.882-51.881l-22.628 22.626l90.51 90.511l90.51-90.511l-22.628-22.626l-51.882 51.881v-161.47h160.667l-51.883 51.882l22.628 22.627l90.509-90.509z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Movie(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M136 488h359.985V24H16v464zM408 56h55.985v40H408Zm0 72h55.985v40H408Zm0 72h55.985v40H408Zm0 72h55.985v40H408Zm0 72h55.985v40H408Zm0 72h55.985v40H408ZM136 200V56h239.985v184H136Zm0 216V272h239.985v184H136ZM48 56h56v40H48Zm0 72h56v40H48Zm0 72h56v40H48Zm0 72h56v40H48Zm0 72h56v40H48Zm0 72h56v40H48Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mug(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M416 180h-96v160h96a20.023 20.023 0 0 0 20-20V200a20.023 20.023 0 0 0-20-20m-12 128h-52v-96h52Z"/><path fill="currentColor" d="M436.574 120H352V64H32v344a64.072 64.072 0 0 0 64 64h192a64.072 64.072 0 0 0 64-64v-8h84.574A59.493 59.493 0 0 0 496 340.574V179.426A59.493 59.493 0 0 0 436.574 120M464 340.574A27.457 27.457 0 0 1 436.574 368H320v40a32.036 32.036 0 0 1-32 32H96a32.036 32.036 0 0 1-32-32V96h256v56h116.574A27.457 27.457 0 0 1 464 179.426Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MugTea(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M436.574 120H352V64H32v344a64.072 64.072 0 0 0 64 64h192a64.072 64.072 0 0 0 64-64v-8h84.574A59.493 59.493 0 0 0 496 340.574V179.426A59.493 59.493 0 0 0 436.574 120m-275.2 118.894L192 269.521v57.373h-64v-57.373l30.627-30.627ZM464 340.574A27.457 27.457 0 0 1 436.574 368H320v40a32.036 32.036 0 0 1-32 32H96a32.036 32.036 0 0 1-32-32V96h80v112.266l-48 48v102.628h128V256.266l-48-48V96h144v56h116.574A27.457 27.457 0 0 1 464 179.426Z"/><path fill="currentColor" d="M416 180h-96v160h96a20.023 20.023 0 0 0 20-20V200a20.023 20.023 0 0 0-20-20m-12 128h-52v-96h52Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicNote(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m72 203.517l160-32v209.669h.044Q232 382.588 232 384c0 55.794 48.448 101.186 108 101.186S448 439.794 448 384s-48.448-101.186-108-101.186a111.434 111.434 0 0 0-76 29.367V28.483l-192 38.4Zm268 111.3c41.906 0 76 31.037 76 69.186s-34.094 69.186-76 69.186s-76-31.04-76-69.189s34.094-69.186 76-69.186Zm-236-221.7l128-25.6v71.366l-128 25.6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Newspaper(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M88 72v316a20 20 0 0 1-40 0V152H16v236a52.059 52.059 0 0 0 52 52h376a52.059 52.059 0 0 0 52-52V72Zm376 316a20.023 20.023 0 0 1-20 20H116a51.722 51.722 0 0 0 4-20V104h344Z"/><path fill="currentColor" d="M296 136H152v160h144Zm-32 128h-80v-96h80Zm64-128h104v32H328zm0 64h104v32H328zm0 64h104v32H328z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoteAdd(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24 24v472h294.627L496 318.627V24Zm32 32h408v216H272v192H56Zm249.373 408H304V304h160v1.373Z"/><path fill="currentColor" d="M208 288v-80h80v-32h-80V96h-32v80H96v32h80v80z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Notes(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M112 152h288v32H112zm0 88h288v32H112zm0 88h152v32H112z"/><path fill="currentColor" d="M480 48H32v416h448Zm-32 384H64V80h384Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ObjectGroup(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M304 136H128v152h80v88h176V224h-80ZM160 256v-88h112v88zm192 0v88H240v-56h64v-32Z"/><path fill="currentColor" d="M400 48H112V16H16v96h32v288H16v96h96v-32h288v32h96v-96h-32V112h32V16h-96ZM48 48h32v32H48Zm32 416H48v-32h32Zm384 0h-32v-32h32ZM432 48h32v32h-32Zm0 352h-32v32H112v-32H80V112h32V80h288v32h32Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ObjectUngroup(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M400 168h-72v-56h32V16h-96v32H112V16H16v96h32v152H16v96h96v-32h72v72h-32v96h96v-32h152v32h96v-96h-32V232h32v-96h-96ZM296 48h32v32h-32Zm32 248v32h-32v-32ZM48 48h32v32H48Zm32 280H48v-32h32Zm32-32v-32H80V112h32V80h152v32h32v152h-32v32Zm104 168h-32v-32h32Zm248 0h-32v-32h32Zm-32-296h32v32h-32Zm0 232h-32v32H248v-32h-32v-72h48v32h96v-96h-32v-64h72v32h32Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Opacity(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m285.177 179l15.513-3.914l-7.827-31.028l-15.514 3.913a176.937 176.937 0 0 0-129.3 133.557l-3.407 15.633l31.266 6.814l3.406-15.634A145.559 145.559 0 0 1 285.177 179"/><path fill="currentColor" d="M363.624 147.871C343.733 72.077 274.643 16 192.7 16C95.266 16 16 95.266 16 192.7c0 82.617 57 152.163 133.735 171.4A176.769 176.769 0 0 0 320.7 496c97.431 0 176.7-79.266 176.7-176.695c-.008-81.234-55.76-151.969-133.776-171.434M48 192.7C48 112.91 112.91 48 192.7 48s144.7 64.91 144.7 144.7s-64.911 144.7-144.7 144.7S48 272.481 48 192.7M320.7 464c-60.931 0-115.21-38.854-135.843-94.792c2.6.115 5.214.184 7.843.184a176.862 176.862 0 0 0 32.7-3.047l97.625 97.625c-.778.013-1.552.03-2.325.03m41.528-6.083L260.26 355.954a176.9 176.9 0 0 0 43.662-26.072L408.37 434.33a144.385 144.385 0 0 1-46.147 23.587Zm69.3-45.692L326.851 307.557a177.082 177.082 0 0 0 27.911-44.5L457.67 365.964a144.661 144.661 0 0 1-26.151 46.261Zm33.594-84.073l-99.42-99.42a176.785 176.785 0 0 0 3.7-36.036c0-3.285-.1-6.547-.276-9.787a145.054 145.054 0 0 1 96.276 136.4c-.01 2.967-.111 5.915-.289 8.843Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Opentype(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M366.675 94.818a145.972 145.972 0 0 0-60.087-25.88a163.023 163.023 0 0 0-57.469-.117C208.125 75.986 168.68 98.5 138.051 132.2C107.3 166.046 87.266 209.3 81.642 254c-5.678 45.161 3.671 89.074 26.326 123.649c22.767 34.748 57.389 57.982 97.489 65.421a157.665 157.665 0 0 0 28.692 2.61a166.3 166.3 0 0 0 28.684-2.512c24.911-4.332 48.895-14.088 71.281-29c52.347-34.845 88.324-93.23 96.24-156.176c4.15-32.953.333-65.5-11.038-94.134c-11.178-28.134-29.381-52.01-52.641-69.04M113.392 258c7.808-62.061 48.3-117.877 100.45-144.012a204.61 204.61 0 0 0-19.166 25.174c-20.723 31.884-35.883 72.466-42.685 114.273c-6.841 42.054-4.837 83.062 5.643 115.468a140.759 140.759 0 0 0 11.1 25.7c-40.457-26.582-62.701-78.071-55.342-136.603m227.3-4.566c-8.56 52.618-31.076 101.106-60.232 129.707c-11.193 10.984-22.93 18.378-33.98 21.393a38.951 38.951 0 0 1-21.852.061c-15.612-4.774-28.59-20.946-36.543-45.538c-8.946-27.663-10.546-63.35-4.5-100.485c6.1-37.516 19.575-73.73 37.93-101.97c17.192-26.451 37.167-43.9 56.241-49.133a43.3 43.3 0 0 1 11.417-1.6a35.562 35.562 0 0 1 10.467 1.541l.025.007c8.133 2.488 15.537 8.034 22.007 16.485c20.395 26.647 27.685 76.279 19.016 129.529Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Options(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 144a64 64 0 1 0-64-64a64.072 64.072 0 0 0 64 64m0-96a32 32 0 1 1-32 32a32.036 32.036 0 0 1 32-32m0 320a64 64 0 1 0 64 64a64.072 64.072 0 0 0-64-64m0 96a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32m0-272a64 64 0 1 0 64 64a64.072 64.072 0 0 0-64-64m0 96a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OptionsHorizontal(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M17.25 12a3 3 0 1 0 3-3a3.004 3.004 0 0 0-3 3zm4.5 0a1.5 1.5 0 1 1-1.5-1.5a1.502 1.502 0 0 1 1.5 1.5z" fill="currentColor"/><path d="M6.75 12a3 3 0 1 0-3 3a3.004 3.004 0 0 0 3-3zm-4.5 0a1.5 1.5 0 1 1 1.5 1.5a1.502 1.502 0 0 1-1.5-1.5z" fill="currentColor"/><path d="M15 12a3 3 0 1 0-3 3a3.004 3.004 0 0 0 3-3zm-4.5 0a1.5 1.5 0 1 1 1.5 1.5a1.502 1.502 0 0 1-1.5-1.5z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paint(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M418.968 215.116A71.976 71.976 0 0 0 408.023 72H384V40a24.028 24.028 0 0 0-24-24H72a24.028 24.028 0 0 0-24 24v104a24.028 24.028 0 0 0 24 24h288a24.028 24.028 0 0 0 24-24v-40h24.023a39.977 39.977 0 0 1 6.079 79.489L240 210.273V280h-56v216h144V280h-56v-42.273ZM352 136H80V48h272Zm-56 176v152h-80V312Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaintBucket(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m412.284 294.37l-12.5 15.642c-8.354 10.454-50.027 64.208-50.027 95.883c0 36.451 28.049 66.105 62.526 66.105s62.527-29.654 62.527-66.105c0-31.675-41.673-85.429-50.028-95.883Zm0 145.63c-16.832 0-30.526-15.3-30.526-34.105c0-11.662 15.485-37.883 30.531-59.2c15.043 21.3 30.522 47.509 30.522 59.2c0 18.805-13.695 34.105-30.527 34.105M122.669 51.492L96.133 124.4L30.092 97.205L17.908 126.8l67.271 27.7L26.9 314.606a48.056 48.056 0 0 0 28.689 61.523l184.719 67.232a48 48 0 0 0 61.523-28.688L397.6 151.56Zm149.1 352.236a16 16 0 0 1-20.508 9.563L66.537 346.059a16 16 0 0 1-9.563-20.507L73.553 280H316.8ZM85.2 248l29.594-81.311l36.333 14.961a32.644 32.644 0 1 0 11.236-29.98l-36.615-15.077l16.046-44.085l214.79 78.177L328 249.219V248Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaperPlane(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M474.444 19.857a20.336 20.336 0 0 0-21.592-2.781L33.737 213.8v38.066l176.037 70.414L322.69 496h38.074l120.3-455.4a20.342 20.342 0 0 0-6.62-20.743M337.257 459.693L240.2 310.37l149.353-163.582l-23.631-21.576L215.4 290.069L70.257 232.012L443.7 56.72Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paperclip(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M153.456 472a136 136 0 0 1-96.167-232.166l196.6-196.6l22.631 22.623l-196.6 196.6A104 104 0 0 0 227 409.539l207.912-207.917A72 72 0 0 0 333.088 99.8L125.171 307.716a40 40 0 1 0 56.568 56.568l179.634-179.632L384 207.279L204.367 386.911a72 72 0 1 1-101.823-101.822L310.461 77.172a104 104 0 1 1 147.078 147.077L249.622 432.166A135.1 135.1 0 0 1 153.456 472"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paragraph(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M440 64H208a128 128 0 0 0 0 256h56v128h112V96h64ZM264 288h-56a96 96 0 0 1 0-192h56Zm80 128h-48V96h48Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paw(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M382.825 304.576a131.562 131.562 0 0 0-253.65 0l-18.248 66.15A80 80 0 0 0 188.046 472h135.908a80 80 0 0 0 77.119-101.274Zm-20.682 116.5A47.638 47.638 0 0 1 323.954 440H188.046a48 48 0 0 1-46.272-60.765l18.248-66.149a99.563 99.563 0 0 1 191.956 0l18.248 66.149a47.636 47.636 0 0 1-8.083 41.845ZM146.1 230.31c2.784-17.4-.908-36.027-10.4-52.463s-23.78-28.947-40.237-35.236c-17.624-6.731-35.6-5.659-50.634 3.017c-29.887 17.256-37.752 59.785-17.529 94.805c9.489 16.436 23.778 28.95 40.235 35.236a64.058 64.058 0 0 0 22.863 4.371a55.133 55.133 0 0 0 27.771-7.389c15.025-8.677 24.945-23.714 27.931-42.341m-31.6-5.058c-1.43 8.929-5.81 15.92-12.333 19.686S87.4 249 78.95 245.775c-9.613-3.671-18.115-11.251-23.941-21.342c-11.2-19.4-8.538-42.8 5.82-51.092a23.483 23.483 0 0 1 11.847-3.058a31.951 31.951 0 0 1 11.368 2.217c9.613 3.673 18.115 11.252 23.941 21.343s8.139 21.248 6.515 31.409m35.066-61.235c11.362 9.083 24.337 13.813 37.458 13.812a54.965 54.965 0 0 0 11.689-1.261c33.723-7.331 54.17-45.443 45.58-84.958c-4.03-18.546-13.828-34.817-27.588-45.818c-14.735-11.78-32.189-16.239-49.147-12.551c-33.722 7.33-54.169 45.442-45.58 84.957c4.031 18.547 13.829 34.818 27.588 45.819m24.788-99.506a22.258 22.258 0 0 1 4.732-.5c5.948 0 12.066 2.327 17.637 6.781c8.037 6.425 13.826 16.234 16.3 27.621c4.76 21.895-4.906 43.368-21.107 46.89c-7.361 1.6-15.305-.628-22.367-6.275c-8.037-6.426-13.826-16.235-16.3-27.621c-4.761-21.901 4.905-43.374 21.105-46.896m292.817 81.117c-15.028-8.676-33.013-9.748-50.634-3.017c-16.457 6.287-30.746 18.8-40.235 35.236s-13.182 35.067-10.4 52.463c2.982 18.627 12.9 33.664 27.931 42.341a55.123 55.123 0 0 0 27.771 7.389a64.054 64.054 0 0 0 22.863-4.371c16.457-6.286 30.746-18.8 40.235-35.236c20.221-35.02 12.356-77.549-17.531-94.805m-10.18 78.805c-5.826 10.091-14.328 17.671-23.941 21.342c-8.446 3.228-16.692 2.931-23.215-.837s-10.9-10.757-12.333-19.686c-1.626-10.161.686-21.314 6.513-31.4s14.328-17.67 23.941-21.343a31.955 31.955 0 0 1 11.368-2.221a23.483 23.483 0 0 1 11.847 3.058c14.358 8.285 17.023 31.682 5.82 51.087m-143.704-47.865a54.965 54.965 0 0 0 11.689 1.261c13.12 0 26.1-4.729 37.458-13.812c13.759-11 23.557-27.272 27.588-45.818c8.589-39.515-11.858-77.627-45.58-84.957c-16.957-3.686-34.412.77-49.147 12.551c-13.76 11-23.558 27.272-27.588 45.817c-8.59 39.515 11.857 77.627 45.58 84.958m-14.31-78.16c2.474-11.387 8.263-21.2 16.3-27.621c5.572-4.454 11.689-6.781 17.637-6.781a22.258 22.258 0 0 1 4.732.5c16.2 3.522 25.866 25 21.107 46.89c-2.476 11.387-8.265 21.2-16.3 27.622c-7.061 5.646-15 7.874-22.367 6.275c-16.203-3.517-25.869-24.993-21.109-46.885"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pen(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M345.994 42.019L179.531 208.481a646.3 646.3 0 0 0-154.206 248.04a24.845 24.845 0 0 0 6 25.708l.087.087a24.84 24.84 0 0 0 17.611 7.342a25.172 25.172 0 0 0 8.1-1.344a646.283 646.283 0 0 0 248.04-154.207L471.62 167.646A88.831 88.831 0 0 0 345.994 42.019M282.531 311.48A614.445 614.445 0 0 1 60.419 453.221a614.435 614.435 0 0 1 141.739-222.113l99.162-99.161l80.372 80.372Zm166.462-166.461l-44.674 44.673l-80.372-80.372l44.674-44.674a56.832 56.832 0 1 1 80.372 80.373"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PenAlt(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23.686 456.521a24.841 24.841 0 0 0 6 25.708l.087.087a24.841 24.841 0 0 0 17.612 7.342a25.179 25.179 0 0 0 8.1-1.344a646.28 646.28 0 0 0 248.04-154.207l166.456-166.461A88.832 88.832 0 1 0 344.354 42.019l-9.534 9.534l-20.791-20.791a50.4 50.4 0 0 0-71.274 0L108.687 164.83l22.626 22.627L265.382 53.389a18.4 18.4 0 0 1 26.019 0l20.792 20.791l-134.3 134.3A646.28 646.28 0 0 0 23.686 456.521m343.3-391.875a56.832 56.832 0 1 1 80.373 80.373l-89.493 89.493l-80.372-80.373Zm-112.124 112.12l80.372 80.372l-54.342 54.342A614.383 614.383 0 0 1 58.779 453.221A614.383 614.383 0 0 1 200.52 231.108Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PenNib(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M487.97 155.853a24.035 24.035 0 0 0-7-18.166L376.126 32.842a24 24 0 0 0-35.546 1.772l-89.936 109.922l-130.137 26.514a27.1 27.1 0 0 0-20.025 17.026L17.883 405.84a27.268 27.268 0 0 0 6.205 28.917l53.271 53.271a27.263 27.263 0 0 0 28.915 6.207l217.86-82.635a27.144 27.144 0 0 0 16.95-19.655l28.1-128.7L479.2 173.232a24.041 24.041 0 0 0 8.77-17.379m-177.6 226.741L97.807 463.222l-13.142-13.143l99.36-99.36a56.061 56.061 0 1 0-22.268-22.986l-99.72 99.72l-13.143-13.143l80.576-212.429l124.717-25.41l83.052 83.051ZM187.42 301.9a24 24 0 1 1 24 24a24 24 0 0 1-24-24m168.391-69.065L280.973 158l78.776-96.28l92.343 92.343Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pencil(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m29.663 482.25l.087.087a24.847 24.847 0 0 0 17.612 7.342a25.178 25.178 0 0 0 8.1-1.345l142.006-48.172l272.5-272.5A88.832 88.832 0 0 0 344.334 42.039l-272.5 272.5l-48.168 142.002a24.844 24.844 0 0 0 5.997 25.709m337.3-417.584a56.832 56.832 0 0 1 80.371 80.373L411.5 180.873L331.127 100.5ZM99.744 331.884L308.5 123.127l80.373 80.373l-208.757 208.756l-121.634 41.262Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func People(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m462.541 316.3l-64.344-42.1l24.774-45.418A79.124 79.124 0 0 0 432.093 192v-72a103.941 103.941 0 0 0-174.609-76.477L279.232 67a71.989 71.989 0 0 1 120.861 53v72a46.809 46.809 0 0 1-5.215 21.452L355.962 284.8l89.058 58.274a42.16 42.16 0 0 1 19.073 35.421V432h-72v32h104v-85.506a74.061 74.061 0 0 0-33.552-62.194"/><path fill="currentColor" d="m318.541 348.3l-64.343-42.1l24.773-45.418A79.124 79.124 0 0 0 288.093 224v-72A104.212 104.212 0 0 0 184.04 47.866C126.723 47.866 80.093 94.581 80.093 152v72a78 78 0 0 0 9.015 36.775l24.908 45.664L50.047 348.3A74.022 74.022 0 0 0 16.5 410.4L16 496h336.093v-85.506a74.061 74.061 0 0 0-33.552-62.194m1.552 115.7H48.186l.31-53.506a42.158 42.158 0 0 1 19.073-35.421l88.682-58.029l-39.051-71.592A46.838 46.838 0 0 1 112.093 224v-72a72 72 0 1 1 144 0v72a46.809 46.809 0 0 1-5.215 21.452L211.962 316.8l89.058 58.274a42.16 42.16 0 0 1 19.073 35.421Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Phone(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m449.366 89.648l-.685-.428l-86.593-42.661l-93.463 124.617l43 57.337a88.529 88.529 0 0 1-83.115 83.114l-57.336-43l-124.616 93.461l42.306 85.869l.356.725l.429.684a25.085 25.085 0 0 0 21.393 11.857h22.344a327.836 327.836 0 0 0 327.836-327.837v-22.345a25.084 25.084 0 0 0-11.856-21.393m-20.144 43.738c0 163.125-132.712 295.837-295.836 295.837h-18.08L87 371.76l84.18-63.135l46.867 35.149h5.333a120.535 120.535 0 0 0 120.4-120.4v-5.333l-35.149-46.866L371.759 87l57.463 28.311Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pin(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m477.8 140.2l-106-106a62.132 62.132 0 0 0-93.617 81.24l-4.913 5.533l-97.654 42.9l-42.611-23.026a24.038 24.038 0 0 0-28.86 4.638L62.6 189.487a23.881 23.881 0 0 0 .479 33.449l101.68 101.679L16 473.373V496h22.627l148.758-148.758L288.6 448.457a23.928 23.928 0 0 0 33.275.642l44.425-41.128a23.978 23.978 0 0 0 4.773-29.092l-23.344-42.858l42.131-90.515l8.6-10.318A62.134 62.134 0 0 0 477.8 140.2m-22.628 65.231a30.125 30.125 0 0 1-42.6 0l-16.885-16.886l-33.08 39.678l-50.7 108.933l28.087 51.566l-34.209 31.669L91.2 205.806l32-33.89l50.969 27.543l118.386-52.008l29.177-32.863l-15.158-15.161a30.126 30.126 0 0 1 42.6-42.6l106 106a30.126 30.126 0 0 1 0 42.6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pizza(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M468.285 106.1c-73.4-48.951-142.775-73.181-212.079-74.076c-70.18-.865-139.406 20.697-211.666 66.031a31.842 31.842 0 0 0-10.572 43.3l194.586 331.519A31.724 31.724 0 0 0 256 488.676h.16a31.722 31.722 0 0 0 27.434-15.537L477.96 149.191a31.959 31.959 0 0 0-9.675-43.091m-17.764 26.624l-194.37 323.952L61.547 125.162C127.108 84.028 189.454 64 252.007 64q1.893 0 3.787.024c62.934.812 126.633 23.285 194.737 68.7Z"/><path fill="currentColor" d="M223.576 231.487a27.731 27.731 0 1 0-22.507 32.113a27.731 27.731 0 0 0 22.507-32.113"/><circle cx="307.735" cy="283.762" r="27.731" fill="currentColor" transform="rotate(-40.627 307.731 283.76)"/><path fill="currentColor" d="m213.757 340.32l20.486-24.583l48 40l-20.487 24.584zm36.625-147.275l64-24l11.236 29.963l-64 24z"/><path fill="currentColor" d="M120.5 181.58c33.288-20.885 84.546-45.571 140.818-45.571q1.234 0 2.473.016c46.876.6 89.878 15 135.329 45.313l17.754-26.623c-50.825-33.893-99.337-50-152.671-50.688c-65.052-.805-123.288 26.973-160.703 50.446Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plant(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M389.053 126.3A302.909 302.909 0 0 0 280.012 18.15L272 13.516l-8.012 4.634A301.084 301.084 0 0 0 113.4 279.042c0 3.445.06 6.944.177 10.4c1.592 46.856 19.511 86.283 51.82 114.018c24.724 21.225 56.438 34.182 90.607 37.273V496h32V240H256v168.528c-54.064-6.263-107.873-44.455-110.444-120.174c-.106-3.095-.16-6.228-.16-9.312A270.286 270.286 0 0 1 272 50.673a270.286 270.286 0 0 1 126.6 228.369c0 3.084-.054 6.217-.16 9.313c-2.056 60.573-36.907 97.127-78.444 112.536v33.867a156.188 156.188 0 0 0 58.607-31.3c32.309-27.735 50.228-67.162 51.82-114.017c.117-3.456.177-6.955.177-10.4A300.948 300.948 0 0 0 389.053 126.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlaylistAdd(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 152h208v32H256zm-80 104h288v32H176zm0 104h288v32H176zm16-208h-64V88H96v64H32v32h64v64h32v-64h64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plus(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M440 240H272V72h-32v168H72v32h168v168h32V272h168z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pool(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M468.479 361.5c-9.072-4.233-20.361-9.5-41.054-9.5s-31.983 5.268-41.053 9.5c-7.782 3.631-13.928 6.5-27.523 6.5s-19.739-2.868-27.52-6.5c-9.071-4.232-20.359-9.5-41.052-9.5s-31.981 5.269-41.051 9.5c-7.781 3.631-13.926 6.5-27.519 6.5s-19.737-2.868-27.517-6.5c-9.07-4.232-20.358-9.5-41.05-9.5s-31.981 5.269-41.051 9.5c-7.781 3.631-13.926 6.5-27.519 6.5s-19.738-2.868-27.519-6.5C47.981 357.269 36.692 352 16 352v32c13.593 0 19.738 2.868 27.519 6.5c9.07 4.232 20.359 9.5 41.051 9.5s31.981-5.269 41.051-9.5c7.781-3.631 13.926-6.5 27.519-6.5s19.737 2.868 27.517 6.5c9.07 4.232 20.358 9.5 41.05 9.5s31.981-5.269 41.051-9.5c7.781-3.631 13.927-6.5 27.519-6.5s19.739 2.868 27.52 6.5c9.071 4.232 20.359 9.5 41.052 9.5s31.983-5.268 41.054-9.5c7.781-3.631 13.928-6.5 27.522-6.5s19.741 2.868 27.521 6.5c9.072 4.233 20.361 9.5 41.054 9.5v-32c-13.594 0-19.741-2.868-27.521-6.5M427.425 448c-20.693 0-31.983 5.268-41.053 9.5c-7.782 3.631-13.928 6.5-27.523 6.5s-19.739-2.868-27.52-6.5c-9.071-4.232-20.359-9.5-41.052-9.5s-31.981 5.269-41.051 9.5c-7.781 3.631-13.926 6.5-27.519 6.5s-19.737-2.868-27.517-6.5c-9.07-4.232-20.358-9.5-41.05-9.5s-31.981 5.269-41.051 9.5c-7.781 3.631-13.926 6.5-27.519 6.5s-19.738-2.868-27.519-6.5C47.981 453.269 36.692 448 16 448v32c13.593 0 19.738 2.868 27.519 6.5c9.07 4.232 20.359 9.5 41.051 9.5s31.981-5.269 41.051-9.5c7.781-3.631 13.926-6.5 27.519-6.5s19.737 2.868 27.517 6.5c9.07 4.232 20.358 9.5 41.05 9.5s31.981-5.269 41.051-9.5c7.781-3.631 13.927-6.5 27.519-6.5s19.739 2.868 27.52 6.5c9.071 4.232 20.359 9.5 41.052 9.5s31.983-5.268 41.054-9.5c7.781-3.631 13.928-6.5 27.522-6.5s19.741 2.868 27.521 6.5c9.072 4.233 20.361 9.5 41.054 9.5v-32c-13.594 0-19.741-2.868-27.521-6.5c-9.072-4.232-20.361-9.5-41.054-9.5M248 272h144v56h32V115.878a47.8 47.8 0 0 1 8.446-27.193L482.417 16h-38.834l-37.507 54.556A79.67 79.67 0 0 0 392 115.878V160H248v-44.122a47.8 47.8 0 0 1 8.446-27.193L306.417 16h-38.834l-37.507 54.556A79.67 79.67 0 0 0 216 115.878V328h32Zm0-80h144v48H248Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PowerStandby(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M312 87.666v33.47C381.676 144.582 432 210.522 432 288c0 97.047-78.953 176-176 176S80 385.047 80 288c0-77.478 50.324-143.418 120-166.864v-33.47C112.422 112.179 48 192.7 48 288c0 114.691 93.309 208 208 208s208-93.309 208-208c0-95.3-64.422-175.821-152-200.334"/><path fill="currentColor" d="M240 16h32v272h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pregnant(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M224 144a64 64 0 1 0-64-64a64.072 64.072 0 0 0 64 64m0-96a32 32 0 1 1-32 32a32.036 32.036 0 0 1 32-32m129.959 203.37c-15.021-16.9-35.063-27.659-62.61-33.506L266.551 160h-88.428L152 342.863V400h56v96h96v-96h80v-48c0-44.972-9.826-77.888-30.041-100.63M352 368h-80v96h-32v-96h-56v-22.863L205.877 192h39.572l23.291 54.344l8.629 1.438c24.5 4.083 41.233 11.979 52.672 24.848C344.817 289.253 352 315.215 352 352Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Print(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M420 128.1V16H92v112.1A80.1 80.1 0 0 0 16 208v192h68v-32H48V208a48.054 48.054 0 0 1 48-48h320a48.054 48.054 0 0 1 48 48v160h-44v32h76V208a80.1 80.1 0 0 0-76-79.9m-32-.1H124V48h264Z"/><path fill="currentColor" d="M396 200h32v32h-32zm-280 64H76v32h40v200h272V296h40v-32zm240 200H148V296h208Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pushchair(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M445.057 345.134L464 274.1V232c-8.136-93.993-87.933-168-184-168h-32v168H132.158l-17.844-78.768A32.155 32.155 0 0 0 83.038 128H16v32h67.038l40.475 178.67A80 80 0 1 0 224 416q0-4.05-.4-8h104.8q-.395 3.948-.4 8a80 80 0 1 0 117.057-70.866M280 96c78.411 0 143.145 59.678 151.164 136H280ZM144 464a48 48 0 1 1 48-48a48.055 48.055 0 0 1-48 48m194.763-88H213.237a80.166 80.166 0 0 0-57.316-39.108L139.408 264H432v5.9l-17.7 66.368a80.592 80.592 0 0 0-6.3-.271A80.026 80.026 0 0 0 338.763 376M408 464a48 48 0 1 1 48-48a48.055 48.055 0 0 1-48 48"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Puzzle(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m491.693 256.705l-54.957-49.461l16.407-13.406a80.491 80.491 0 0 0 18.363-21.522c18.148-31.441 12.867-70.042-13.144-96.052s-64.612-31.291-96.051-13.142a80.513 80.513 0 0 0-21.52 18.362l-13.408 16.407l-49.461-54.956l-.579-.611a24.028 24.028 0 0 0-33.941 0l-65.6 65.605l1.19 23.7l33.108 27.056a48.6 48.6 0 0 1 11.079 12.889c10.807 18.722 7.57 41.8-8.056 57.426s-38.7 18.862-57.426 8.058a48.66 48.66 0 0 1-12.9-11.086l-27.047-33.1l-23.7-1.189l-71.26 71.26a24 24 0 0 0 0 33.942l175.357 175.359a80 80 0 0 0 113.138 0L492.3 291.225a24.029 24.029 0 0 0 0-33.94ZM288.657 449.617a48 48 0 0 1-67.883 0L51.069 279.911l53.1-53.095l15.91 19.473l.1.119a80.487 80.487 0 0 0 21.521 18.363c31.441 18.149 70.041 12.867 96.052-13.144s31.291-64.61 13.143-96.05a80.482 80.482 0 0 0-18.363-21.521l-19.591-16.01l47.124-47.124l56.018 62.241l24.282-.579l25.062-30.67a48.611 48.611 0 0 1 12.888-11.078c18.722-10.807 41.8-7.569 57.426 8.056s18.864 38.7 8.057 57.426a48.591 48.591 0 0 1-11.079 12.889l-30.67 25.061l-.58 24.282l62.243 56.018Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QrCode(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M48 48h128V16H16v160h32z"/><path fill="currentColor" d="M176 176V80H80v96zm-64-64h32v32h-32Zm216-64h136v128h32V16H328z"/><path fill="currentColor" d="M432 176V80h-96v96zm-64-64h32v32h-32ZM176 464H48V336H16v160h160z"/><path fill="currentColor" d="M176 336H80v96h96zm-32 64h-32v-32h32Zm320 64H328v32h168V336h-32z"/><path fill="currentColor" d="M272 304h128v64h32v-96H272z"/><path fill="currentColor" d="M432 432v-32H240V272H80v32h128v128zM208 80h32v96h-32z"/><path fill="currentColor" d="M80 240h224V80h-32v128H80zm256-32h96v32h-96zm0 128h32v32h-32zm-64 0h32v32h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rain(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M398.2 137.208a144.013 144.013 0 0 0-284.545.979a122.364 122.364 0 0 0-64.357 32.926A109.4 109.4 0 0 0 16 249.619c0 31.119 12.789 60.762 36.01 83.469q2.84 2.776 5.845 5.347l11.327-33.981C56.091 289.3 48 270.017 48 249.619c0-42.362 35.724-78.015 81.329-81.168l14.055-.972l.814-14.065a111.995 111.995 0 0 1 223.589-.22l.891 14.888l14.913.155c46.592.488 80.409 34.714 80.409 81.382c0 33.152-16.706 61.38-41.84 75.9L409.032 364.9a110.012 110.012 0 0 0 54.938-32.358c20.655-22.203 32.03-51.653 32.03-82.923c0-59.112-41.141-105.219-97.8-112.411"/><path fill="currentColor" d="M126.029 496h33.788l63.336-186.864l-30.306-10.272zm168 0h33.788l63.336-186.864l-30.306-10.272zm-3.919-244.967L225.781 448h33.664l61.084-187.033zm-161.319 0L64.461 448h33.664l61.084-187.033z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rectangle(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M36 416h440a20.023 20.023 0 0 0 20-20V116a20.023 20.023 0 0 0-20-20H36a20.023 20.023 0 0 0-20 20v280a20.023 20.023 0 0 0 20 20m12-288h416v256H48Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Recycle(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M68.328 383.063a31.654 31.654 0 0 1 .207-32.118l50.883-86.406l11.516 50.76l31.207-7.08l-23.884-105.275l-105.274 23.884l7.08 31.207l53.149-12.058l-52.252 88.73a64 64 0 0 0 55.149 96.476h82.435l32-32H96.109a31.655 31.655 0 0 1-27.781-16.12M283.379 79.762l53.747 91.268l-49.053-7.653l-4.934 31.617L389.8 211.635l16.64-106.66l-31.617-4.933l-8.873 56.87l-54.996-93.388a64 64 0 0 0-110.3 0l-39.939 67.82l10.407 45.39l57.106-96.972a32 32 0 0 1 55.148 0ZM470.65 334.707l-47.867-81.283l-41.148-6.812l61.441 104.333a32 32 0 0 1-27.576 48.238H304.046l38.359-38.358l-22.627-22.625l-76.332 76.332l76.332 76.333l22.627-22.628l-37.052-37.051H415.5a64 64 0 0 0 55.149-96.476Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reload(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M265.614 206.387H456V16h-32v133.887l-26.137-26.137c-79.539-79.539-208.96-79.54-288.5 0s-79.539 208.96 0 288.5a204.232 204.232 0 0 0 288.5 0l-22.627-22.627c-67.063 67.063-176.182 67.063-243.244 0s-67.063-176.183 0-243.246s176.182-67.063 243.245 0l28.01 28.01H265.614Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Remove(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M10.655 17.03l3.97-3.97l3.97 3.97l1.061-1.061l-3.97-3.97l3.97-3.97l-1.061-1.061l-3.97 3.97l-3.97-3.97l-1.061 1.061l3.97 3.97l-3.97 3.97l1.061 1.061z" fill="currentColor"/><path d="M22.125 3H9.124a1.126 1.126 0 0 0-.816.351L.751 11.326v1.348l7.557 7.975c.206.216.495.351.816.351h13.001a1.127 1.127 0 0 0 1.125-1.125V4.125A1.127 1.127 0 0 0 22.125 3zm-.375 16.5H9.285L2.25 12.076v-.152L9.285 4.5H21.75z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReportSlash(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M240 352h32v32h-32zM355.411 16H156.589l-34.936 37.771l22.645 22.645L170.58 48h170.84L464 180.53v150.94l-31.063 33.585l22.644 22.645L496 344V168z"/><path fill="currentColor" d="M240 128v44.118l32 32V128zM16 16v22.627l62.164 62.164L16 168v176l140.589 152h198.822l56.681-61.281L473.373 496H496v-22.627L38.627 16Zm325.42 448H170.58L48 331.47V180.53l52.808-57.095L240 262.627V320h32v-25.373l117.447 117.447Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ResizeBoth(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m404.687 212.686l-50.51 50.51l-105.372-105.373l50.841-50.843L272.666 80H80v192l27.313 27.314l50.982-50.981l105.372 105.373l-51.313 51.313L239.333 432H432V240ZM400 400H262.627l46.3-46.294l-150.632-150.628L112 249.373V112h137.373l-45.823 45.823l150.627 150.628L400 262.627Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ResizeHeight(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256.236 504L120 367.764v-38.156h72v-144h-72V146.98L255.766 11.216L392 147.452v38.156h-72v144h72v38.627ZM159.1 361.608l97.137 97.137l97.137-97.137H288v-208h64.9l-97.134-97.137l-97.138 97.137H224v208Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ResizeWidth(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M366.156 392H328v-72H184v72h-38.627L9.607 256.235L145.845 120H184v72h144v-72h38.627l135.766 135.765ZM54.863 256.235L152 353.373V288h208v64.9l97.137-97.137L360 158.627V224H152v-64.9Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Restaurant(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M224 160h-48V48h-32v112H96V48H64v124c0 45.505 34.655 83.393 80 90.715V472h32V262.715c45.345-7.322 80-45.21 80-90.715V48h-32Zm-64 72c-27.811 0-51.524-16.722-60.33-40h120.66c-8.806 23.278-32.519 40-60.33 40M413.567 40.187A138.648 138.648 0 0 0 296 177.224V344h104v128h32V37.351ZM400 312h-72V177.224a105.986 105.986 0 0 1 72-100.911Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Room(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M440 424V88h-88V13.005L88 58.522V424H16v32h86.9L352 490.358V120h56v336h88v-32Zm-120 29.642l-200-27.586V85.478L320 51Z"/><path fill="currentColor" d="M256 232h32v64h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Router(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m443.057 132.943l22.634-22.634a143.764 143.764 0 0 0-211.382 0l22.634 22.634a111.838 111.838 0 0 1 166.114 0"/><path fill="currentColor" d="m299.615 155.615l22.7 22.7a47.913 47.913 0 0 1 75.362 0l22.7-22.7a79.829 79.829 0 0 0-120.77 0ZM472 280h-96v-64h-32v64H40a24.028 24.028 0 0 0-24 24v112a24.028 24.028 0 0 0 24 24h432a24.028 24.028 0 0 0 24-24V304a24.028 24.028 0 0 0-24-24m-8 128H48v-96h416Z"/><path fill="currentColor" d="M96 344h32v32H96zm80 0h32v32h-32zm80 0h32v32h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rowing(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M302.569 106.243A36 36 0 1 0 286.088 76a35.825 35.825 0 0 0 16.481 30.243M419.9 396.687L381.412 358.2l-11.375-11.375l-74.621-19.515L188.173 216h59.9L288 161.935v117.827l32.079 33.538L352 319.683V131.314h-81.167L231.926 184h-74.584l-33.82-35.1l-23.044 22.2l112.877 117.158L85.694 395.762l.779.925H16v32h232v-32H134.274l101.333-85.334l34.619 35.931l20.617 78.738l59.634 59.633a24 24 0 0 0 33.942 0l45.255-45.255l.1-.1a38.176 38.176 0 0 1 27.092-11.614H496v-32Zm-52.451 60.685L319.6 409.526l-12.036-45.963l45.979 12.025l47.844 47.844Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rss(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M66 256v32c88.225 0 160 71.776 160 160h32c0-105.869-86.131-192-192-192"/><path fill="currentColor" d="M66 140v32c152.187 0 276 123.813 276 276h32a305.982 305.982 0 0 0-90.211-217.789A305.987 305.987 0 0 0 66 140"/><path fill="currentColor" d="M456.674 282.955a422.588 422.588 0 0 0-90.861-134.768A422.724 422.724 0 0 0 66 24v32c216.149 0 392 175.851 392 392h32a421.378 421.378 0 0 0-33.326-165.045M90 360a64 64 0 1 0 64 64a64.072 64.072 0 0 0-64-64m0 96a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ruble(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M291.75 308.1a105.8 105.8 0 1 0 0-211.6H136v32h39.943v147.6H136v32h39.943V352H136v32h39.943v56h32v-56H304v-32h-96.057v-43.9Zm-83.807-179.6h83.807a73.8 73.8 0 1 1 0 147.6h-83.807Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Running(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m353.415 200l-30.981-57.855l-60.717-20.239l-.14.432L167.21 149.3A32.133 32.133 0 0 0 144 180.068V264h32v-83.931l73.6-21.028l-32.512 99.633l-.155-.056l-29.464 82.5a16.088 16.088 0 0 1-20.127 9.8l-66.282-22.097l-10.12 30.358l66.282 22.093a48 48 0 0 0 60.378-29.391l24.232-67.849l17.173 5.6l48.3 48.3A15.9 15.9 0 0 1 312 349.255V456h32V349.255a47.694 47.694 0 0 0-14.059-33.942l-48.265-48.264l26.783-82.077l19.269 34.683A24.011 24.011 0 0 0 348.707 232H432v-32Zm-66.587-90.293a36 36 0 1 0-12.916-27.619a35.851 35.851 0 0 0 12.916 27.619"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sad(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 16C123.452 16 16 123.452 16 256s107.452 240 240 240s240-107.452 240-240S388.548 16 256 16m147.078 387.078a207.253 207.253 0 1 1 44.589-66.125a207.332 207.332 0 0 1-44.589 66.125"/><path fill="currentColor" d="M152 200h40v40h-40zm168 0h40v40h-40zm-64 88a88.1 88.1 0 0 0-88 88h32a56 56 0 0 1 112 0h32a88.1 88.1 0 0 0-88-88"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Satelite(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 496h480V16H16Zm448-32h-25.373l-104-104l54.912-54.911L464 379.55ZM48 48h416v286.3l-74.461-74.461L312 337.373l-112-112l-152 152Zm0 374.627l152-152L393.373 464H48Z"/><path fill="currentColor" d="M120 80H80v40a40 40 0 0 0 40-40m-40 83.661V196.6A152.468 152.468 0 0 0 196.6 80h-32.939A120.471 120.471 0 0 1 80 163.661"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Save(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m472.971 122.344l-99.315-99.315A23.838 23.838 0 0 0 356.687 16H56a24.028 24.028 0 0 0-24 24v432a24.028 24.028 0 0 0 24 24h400a24.028 24.028 0 0 0 24-24V139.313a23.838 23.838 0 0 0-7.029-16.969M320 48v96H176V48Zm128 416H64V48h80v128h208V48h1.373L448 142.627Z"/><path fill="currentColor" d="M252 224a92 92 0 1 0 92 92a92.1 92.1 0 0 0-92-92m0 152a60 60 0 1 1 60-60a60.068 60.068 0 0 1-60 60"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func School(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m368 350.643l-112 63l-112-63v-66.562l-32-17.778v103.054l144 81l144-81V266.303l-32 17.778z"/><path fill="currentColor" d="M256 45.977L32 162.125v27.734L256 314.3l192-106.663V296h32V162.125Zm160 142.831l-32 17.777L256 277.7l-128-71.115l-32-17.777l-22.179-12.322L256 82.023l182.179 94.463Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScreenDesktop(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M40 392h200v72h-80v32h192v-32h-80v-72h200a24.028 24.028 0 0 0 24-24V72a24.028 24.028 0 0 0-24-24H40a24.028 24.028 0 0 0-24 24v296a24.028 24.028 0 0 0 24 24m8-312h416v280H48Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScreenSmartphone(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M104 48v416a32.036 32.036 0 0 0 32 32h248a32.036 32.036 0 0 0 32-32V48a32.036 32.036 0 0 0-32-32H136a32.036 32.036 0 0 0-32 32m280.021 416H136V48h248Z"/><path fill="currentColor" d="M216 80h96v32h-96zm32 312h32v32h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Scrubber(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M425.856 87.379A239.364 239.364 0 1 0 87.344 425.892A239.364 239.364 0 1 0 425.856 87.379M256.6 464c-114.341 0-207.365-93.023-207.365-207.365S142.259 49.271 256.6 49.271s207.364 93.023 207.364 207.364S370.941 464 256.6 464"/><path fill="currentColor" d="M256.6 192.635a64 64 0 1 0 64 64a64.073 64.073 0 0 0-64-64m0 96a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Search(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m479.6 399.716l-81.084-81.084l-62.368-25.767A175.014 175.014 0 0 0 368 192c0-97.047-78.953-176-176-176S16 94.953 16 192s78.953 176 176 176a175.034 175.034 0 0 0 101.619-32.377l25.7 62.2l81.081 81.088a56 56 0 1 0 79.2-79.195M48 192c0-79.4 64.6-144 144-144s144 64.6 144 144s-64.6 144-144 144S48 271.4 48 192m408.971 264.284a24.028 24.028 0 0 1-33.942 0l-76.572-76.572l-23.894-57.835l57.837 23.894l76.573 76.572a24.028 24.028 0 0 1-.002 33.941"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Send(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M474.444 19.857a20.336 20.336 0 0 0-21.592-2.781L33.737 213.8v38.066l176.037 70.414L322.69 496h38.074l120.3-455.4a20.342 20.342 0 0 0-6.62-20.743M337.257 459.693L240.2 310.37l149.353-163.582l-23.631-21.576L215.4 290.069L70.257 232.012L443.7 56.72Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Settings(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M245.151 168a88 88 0 1 0 88 88a88.1 88.1 0 0 0-88-88m0 144a56 56 0 1 1 56-56a56.063 56.063 0 0 1-56 56"/><path fill="currentColor" d="m464.7 322.319l-31.77-26.153a193.081 193.081 0 0 0 0-80.332l31.77-26.153a19.941 19.941 0 0 0 4.606-25.439l-32.612-56.483a19.936 19.936 0 0 0-24.337-8.73l-38.561 14.447a192.038 192.038 0 0 0-69.54-40.192l-6.766-40.571A19.936 19.936 0 0 0 277.762 16H212.54a19.937 19.937 0 0 0-19.728 16.712l-6.762 40.572a192.03 192.03 0 0 0-69.54 40.192L77.945 99.027a19.937 19.937 0 0 0-24.334 8.731L21 164.245a19.94 19.94 0 0 0 4.61 25.438l31.767 26.151a193.081 193.081 0 0 0 0 80.332l-31.77 26.153A19.942 19.942 0 0 0 21 347.758l32.612 56.483a19.937 19.937 0 0 0 24.337 8.73l38.562-14.447a192.03 192.03 0 0 0 69.54 40.192l6.762 40.571A19.937 19.937 0 0 0 212.54 496h65.222a19.936 19.936 0 0 0 19.728-16.712l6.763-40.572a192.038 192.038 0 0 0 69.54-40.192l38.564 14.449a19.938 19.938 0 0 0 24.334-8.731l32.609-56.487a19.939 19.939 0 0 0-4.6-25.436m-50.636 57.12l-48.109-18.024l-7.285 7.334a159.955 159.955 0 0 1-72.625 41.973l-10 2.636L267.6 464h-44.89l-8.442-50.642l-10-2.636a159.955 159.955 0 0 1-72.625-41.973l-7.285-7.334l-48.117 18.024L53.8 340.562l39.629-32.624l-2.7-9.973a160.9 160.9 0 0 1 0-83.93l2.7-9.972L53.8 171.439l22.446-38.878l48.109 18.024l7.285-7.334a159.955 159.955 0 0 1 72.625-41.973l10-2.636L222.706 48H267.6l8.442 50.642l10 2.636a159.955 159.955 0 0 1 72.625 41.973l7.285 7.334l48.109-18.024l22.447 38.877l-39.629 32.625l2.7 9.972a160.9 160.9 0 0 1 0 83.93l-2.7 9.973l39.629 32.623Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Share(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M68.146 496H16V235.333a114.169 114.169 0 0 1 12.025-51.309A136.27 136.27 0 0 1 152 104h136.557V16h42.361l163.709 163.71L330.337 344h-41.78v-88h-88.812a85.4 85.4 0 0 0-81.993 62.244ZM152 136a104.217 104.217 0 0 0-94.923 61.443l-.292.614A82.454 82.454 0 0 0 48 235.333v213.81l38.93-139.5A117.5 117.5 0 0 1 199.745 224h120.812v84.525L449.373 179.71L320.557 50.894V136Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareAll(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M242.918 16.173h-42.361v88H152A136.268 136.268 0 0 0 28.025 184.2A114.159 114.159 0 0 0 16 235.506v260.667h52.146l49.606-177.756a85.4 85.4 0 0 1 81.993-62.244h.812v88h41.78l164.29-164.29ZM232.557 308.7v-84.527h-32.812A117.5 117.5 0 0 0 86.93 309.815L48 449.315V235.506a82.454 82.454 0 0 1 8.785-37.276l.292-.614A104.217 104.217 0 0 1 152 136.173h80.557V51.067l128.816 128.816Z"/><path fill="currentColor" d="M330.918 15.509h-43.409l164.373 164.374l-163.626 163.626h42.081l164.29-164.29z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareAlt(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M404 344a75.9 75.9 0 0 0-60.208 29.7l-163.923-93.036a75.693 75.693 0 0 0 0-49.328L343.792 138.3a75.937 75.937 0 1 0-13.776-28.976L163.3 203.946a76 76 0 1 0 0 104.108l166.717 94.623A75.991 75.991 0 1 0 404 344m0-296a44 44 0 1 1-44 44a44.049 44.049 0 0 1 44-44M108 300a44 44 0 1 1 44-44a44.049 44.049 0 0 1-44 44m296 164a44 44 0 1 1 44-44a44.049 44.049 0 0 1-44 44"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareBoxed(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M464 248v216H48V48h216V16H16v480h480V216z"/><path fill="currentColor" d="M106.12 171.135A96.274 96.274 0 0 0 96 214.364v216.181h47.782l41.181-147.564a66.953 66.953 0 0 1 64.283-48.8H304V320h38.6l152.027-151.1L342.656 16H304v88h-93.818a114.4 114.4 0 0 0-104.062 67.135M336 136V54.7l113.373 114.058L336 281.441v-79.259h-86.754a99.055 99.055 0 0 0-95.105 72.2L128 368.051V214.364a64.576 64.576 0 0 1 6.879-29.2l.292-.614A82.356 82.356 0 0 1 210.182 136Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldAlt(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M473.605 88.081c-1.352-.137-135.958-14.259-199.218-68.251L269.9 16h-27.8l-4.488 3.83C174.464 73.727 39.744 87.944 38.4 88.081L24 89.532V104c0 89.133 14.643 165.443 43.523 226.813c38.105 80.973 100.1 133.669 184.267 156.623l4.21 1.148l4.21-1.148c84.165-22.954 146.162-75.65 184.267-156.623C473.357 269.443 488 193.133 488 104V89.532Zm-17.735 30.032q-.237 12.789-.948 25.887H272V57.915c59.921 39.567 149.024 55.322 183.87 60.198M272 320h142.266a288.233 288.233 0 0 1-23.366 40H272Zm0-32v-40h167.9a402.662 402.662 0 0 1-13.236 42.884V288Zm0-72v-40h180.378c-1.4 13.307-3.256 26.682-5.639 40ZM56.13 118.113c34.846-4.876 123.949-20.631 183.87-60.2v392.311C94.012 398.389 58.492 245.387 56.13 118.113M272 450.224V392h92.347c-24.298 24.7-54.639 44.836-92.347 58.224"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShortText(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M128 192h256v32H128zm0 112h128v32H128z"/><path fill="currentColor" d="M48 432h416V88H48Zm32-312h352v280H80Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shower(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m240.801 293.826l-23.851 23.851l23.8 47.6l23.417-24.718zm15.177 101.906l22.299 44.6l23.418-24.719l-22.3-44.599zM438.926 202.62L415.8 227.032l44.423 21.246l23.127-24.412zm-96.323-10.596l42.861 20.499l23.127-24.411l-41.992-20.084zm-34.818 149.022l28.523 38.031l22.325-23.565l-28.523-38.031zm2.848-49.534l-28.936-38.582l-22.857 22.857l29.468 39.29zm-9.89-57.628l36.683 29.347l22.085-23.313l-36.001-28.801zm61.758 49.407l36.721 29.377l22.085-23.313l-36.721-29.376zm-13.814-182.604l-26.24 26.239l-24.718-24.718a111.609 111.609 0 0 0-157.839 0c-.342.341-.673.689-1.009 1.034A77.974 77.974 0 0 0 16 166.988V408h32V166.988a45.975 45.975 0 0 1 72.048-37.868a111.809 111.809 0 0 0 19.842 130.929l24.717 24.717l-23.92 23.921l20 20l208-208ZM185.006 259.911l-22.489-22.489A79.611 79.611 0 0 1 275.1 124.835l22.489 22.49Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignLanguage(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M445.646 241.836V68a8 8 0 0 0-8-8a59.894 59.894 0 0 0-59.9 59.894v41.929l-8.639 8.64l-69.928-69.929a22.751 22.751 0 0 0-26.446-4.134l-56.2-56.2a28 28 0 0 0-39.6 0a28 28 0 0 0-2.711 36.453v.434l-19.916-19.915a28 28 0 0 0-39.6 39.6l22.628 22.628a28 28 0 0 0-39.6 39.6l28.284 28.283a28 28 0 0 0-39.6 0a28 28 0 0 0 0 39.6l14.142 14.142l2.553 2.553a42.93 42.93 0 0 0-28.278 40.3v2.022a42.751 42.751 0 0 0 1.524 11.348A42.927 42.927 0 0 0 40 339.619v2.031a42.927 42.927 0 0 0 36.365 42.367a42.751 42.751 0 0 0-1.524 11.348v2.035a42.906 42.906 0 0 0 42.859 42.854h14.28a42.777 42.777 0 0 0-1.392 10.857v2.032A42.905 42.905 0 0 0 173.444 496h149.107l139.672-31.038V256.906Zm-15.423 197.456L319.037 464H174.46a11.886 11.886 0 0 1-11.872-11.854v-.038a11.886 11.886 0 0 1 11.872-11.854h90.588v-32H118.714a11.873 11.873 0 1 1 0-23.746h146.334v-32H83.873a11.873 11.873 0 0 1 0-23.746h181.175v-32H117.7a10.87 10.87 0 0 1-10.859-10.862v-2.031a10.869 10.869 0 0 1 10.859-10.853h216.523v-48.751l-46.152-46.151a37.778 37.778 0 0 1-4.489-47.926l82.944 82.943l.269.27l63.428 57.661Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignalCellularFour(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384.78 88h-.1v.1L208.78 264h-.1v.1L16 456.78V496h480V16h-39.22Zm-264.1 376H54.034l66.647-66.646Zm88 0h-56v-98.646l56-56Zm88 0h-56V277.354l56-56Zm88 0h-56V189.354l56-56Zm79.32 0h-47.319V101.354L464 54.034Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignalCellularThree(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M208.78 264h-.1v.1L16 456.78V496h480V16h-39.22Zm-88.1 200H54.035l66.646-66.646Zm88 0h-56v-98.646l56-56Zm88 0h-56V277.354l56-56ZM464 464H328.681V189.354L464 54.034Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignalCellularZero(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 456.78V496h480V16h-39.22ZM464 464H54.035L464 54.034Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sim(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 40H230.627A31.791 31.791 0 0 0 208 49.373L97.373 160A31.791 31.791 0 0 0 88 182.627V448a32.036 32.036 0 0 0 32 32h264a32.036 32.036 0 0 0 32-32V72a32.036 32.036 0 0 0-32-32m0 408H120V182.627L230.627 72H384Z"/><path fill="currentColor" d="M208 416h144V216H208Zm32-168h80v136h-80Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sitemap(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M472 328h-24v-64a24.027 24.027 0 0 0-24-24H272v-64h32a24.028 24.028 0 0 0 24-24V80a24.028 24.028 0 0 0-24-24h-96a24.028 24.028 0 0 0-24 24v72a24.028 24.028 0 0 0 24 24h32v64H88a24.027 24.027 0 0 0-24 24v64H40a24.028 24.028 0 0 0-24 24v72a24.028 24.028 0 0 0 24 24h80a24.028 24.028 0 0 0 24-24v-72a24.028 24.028 0 0 0-24-24H96v-56h144v56h-24a24.028 24.028 0 0 0-24 24v72a24.028 24.028 0 0 0 24 24h80a24.028 24.028 0 0 0 24-24v-72a24.028 24.028 0 0 0-24-24h-24v-56h144v56h-24a24.028 24.028 0 0 0-24 24v72a24.028 24.028 0 0 0 24 24h80a24.028 24.028 0 0 0 24-24v-72a24.028 24.028 0 0 0-24-24M216 88h80v56h-80ZM112 360v56H48v-56Zm176 0v56h-64v-56Zm176 56h-64v-56h64Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Smile(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 16C123.452 16 16 123.452 16 256s107.452 240 240 240s240-107.452 240-240S388.548 16 256 16m147.078 387.078a207.253 207.253 0 1 1 44.589-66.125a207.332 207.332 0 0 1-44.589 66.125"/><path fill="currentColor" d="M152 200h40v40h-40zm168 0h40v40h-40zm18.289 107.2A83.6 83.6 0 0 1 260.3 360h-8.6a83.6 83.6 0 0 1-77.992-52.8l-1.279-3.2h-34.461L144 319.081A116 116 0 0 0 251.7 392h8.6A116 116 0 0 0 368 319.081L374.032 304h-34.464Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmilePlus(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M432 80V16h-32v64h-64v32h64v64h32v-64h64V80zM112 256h40v40h-40zm136 0h40v40h-40zm-44.562 128h-6.875a63.691 63.691 0 0 1-59.326-40h-34.47l4.662 11.653A95.541 95.541 0 0 0 196.563 416h6.875a95.54 95.54 0 0 0 89.133-60.347L297.233 344h-34.47a63.691 63.691 0 0 1-59.325 40"/><path fill="currentColor" d="M200 128C98.542 128 16 210.542 16 312s82.542 184 184 184s184-82.542 184-184s-82.542-184-184-184m0 336c-83.813 0-152-68.187-152-152s68.187-152 152-152s152 68.187 152 152s-68.187 152-152 152"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Smoke(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 240v120h344V240zm312 88H48v-56h280Zm56-88h32v120h-32zm56 0h32v120h-32zm-54.572-66.7a31.982 31.982 0 0 1 2.32-38.418a63.745 63.745 0 0 0 3.479-78.69L385.377 48H348.8l-1.82 1.3l18.207 25.49a31.807 31.807 0 0 1-1.736 39.265a64.1 64.1 0 0 0-4.649 76.993L364.77 200h38.46Zm72 0a31.982 31.982 0 0 1 2.32-38.418a63.745 63.745 0 0 0 3.479-78.69L457.377 48H420.8l-1.82 1.3l18.207 25.49a31.807 31.807 0 0 1-1.736 39.265a64.1 64.1 0 0 0-4.649 76.993L436.77 200h38.46Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmokeFree(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m384 316.118l32 32V240h-32zM440 240h32v120h-32zm-76.549-125.945a64.1 64.1 0 0 0-4.649 76.993L364.77 200h38.46l-17.8-26.7a31.982 31.982 0 0 1 2.32-38.418a63.745 63.745 0 0 0 3.479-78.69L385.377 48H348.8l-1.82 1.3l18.208 25.49a31.808 31.808 0 0 1-1.737 39.265m72 0a64.1 64.1 0 0 0-4.649 76.993L436.77 200h38.46l-17.8-26.7a31.982 31.982 0 0 1 2.32-38.418a63.745 63.745 0 0 0 3.479-78.69L457.377 48H420.8l-1.82 1.3l18.208 25.49a31.808 31.808 0 0 1-1.737 39.265M262.627 240l-224-224H16v22.627L217.373 240H16v120h321.373l136 136H496v-22.627ZM48 328v-56h201.373l56 56Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmokeSlash(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m384 316.118l32 32V240h-32zM440 240h32v120h-32zm-76.549-125.945a64.1 64.1 0 0 0-4.649 76.993L364.77 200h38.46l-17.8-26.7a31.982 31.982 0 0 1 2.32-38.418a63.745 63.745 0 0 0 3.479-78.69L385.377 48H348.8l-1.82 1.3l18.208 25.49a31.808 31.808 0 0 1-1.737 39.265m72 0a64.1 64.1 0 0 0-4.649 76.993L436.77 200h38.46l-17.8-26.7a31.982 31.982 0 0 1 2.32-38.418a63.745 63.745 0 0 0 3.479-78.69L457.377 48H420.8l-1.82 1.3l18.208 25.49a31.808 31.808 0 0 1-1.737 39.265M262.627 240l-224-224H16v22.627L217.373 240H16v120h321.373l136 136H496v-22.627ZM48 328v-56h201.373l56 56Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmokingRoom(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 240v120h344V240zm312 88H48v-56h280Zm56-88h32v120h-32zm56 0h32v120h-32zm-54.572-66.7a31.982 31.982 0 0 1 2.32-38.418a63.745 63.745 0 0 0 3.479-78.69L385.377 48H348.8l-1.82 1.3l18.207 25.49a31.807 31.807 0 0 1-1.736 39.265a64.1 64.1 0 0 0-4.649 76.993L364.77 200h38.46Zm72 0a31.982 31.982 0 0 1 2.32-38.418a63.745 63.745 0 0 0 3.479-78.69L457.377 48H420.8l-1.82 1.3l18.207 25.49a31.807 31.807 0 0 1-1.736 39.265a64.1 64.1 0 0 0-4.649 76.993L436.77 200h38.46Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Snowflake(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m135.279 206.728l88.857 51.595l-80.039 46.474l-95.789-25.668v33.13l74.847 20.054l-20.165 75.258l28.69 16.564l23.838-88.966L240 286.115v88.287l-71.177 71.177l28.69 16.565L258 401.657l59.135 59.135l28.691-16.564L272 370.402v-84.287l83.002 48.195l24.277 90.604l28.691-16.565l-20.374-76.036L464 311.841v-33.129l-96.492 25.855l-79.644-46.244l88.463-51.366L464 230.449V197.32l-69.654-18.663l19.23-71.769l-28.69-16.565l-23.69 88.416L272 230.53v-99.962l66.833-66.833l-28.69-16.564l-55.172 55.172l-54.307-54.306l-28.69 16.564L240 132.627v97.903l-90.675-52.65l-23.252-86.777l-28.691 16.564l19.022 70.99l-68.096 18.246v33.129z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Soccer(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M425.706 86.294A240 240 0 0 0 86.294 425.706A240 240 0 0 0 425.706 86.294m-4.464 43.535A206.875 206.875 0 0 1 463.824 247.8l-66.14-49.772ZM316.033 56.845l-58.378 43.37l-61.125-43.538a208.143 208.143 0 0 1 119.5.168ZM116.8 198.047L48.156 248.33a206.9 206.9 0 0 1 43.092-119.141ZM86.2 376h79.365l25.035 77.458A208.923 208.923 0 0 1 86.2 376m140.787 85.967L188.85 344H67.562a206.3 206.3 0 0 1-17.324-57.527l104.967-76.9l-39.648-106.858a208.938 208.938 0 0 1 45.714-31.864l96.781 68.934l92.741-68.9a208.922 208.922 0 0 1 45.884 32.048L359.833 209.6l101.951 76.721A206.272 206.272 0 0 1 444.438 344H327.512l-42.045 117.9a208.076 208.076 0 0 1-58.482.064Zm95.606-8.9L350.075 376H425.8a208.961 208.961 0 0 1-103.209 77.069Z"/><path fill="currentColor" d="M346.809 223.427L257.854 158.8L168.9 223.427L202.876 328h109.955ZM289.582 296h-63.457l-19.609-60.351l51.338-37.3l51.337 37.3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sofa(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 242.025V152a48.055 48.055 0 0 0-48-48H112a48.055 48.055 0 0 0-48 48v90.025A64.115 64.115 0 0 0 16 304v112a32.036 32.036 0 0 0 32 32h16v48h32v-48h320v48h32v-48h16a32.036 32.036 0 0 0 32-32V304a64.115 64.115 0 0 0-48-61.975M112 416H48V304a32 32 0 0 1 64 0Zm256 0H144v-96h224Zm2.025-128h-228.05A64.243 64.243 0 0 0 96 242.025V152a16.019 16.019 0 0 1 16-16h288a16.019 16.019 0 0 1 16 16v90.025A64.243 64.243 0 0 0 370.025 288M464 416h-64V304a32 32 0 0 1 64 0l.02 112Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortAlphaDown(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m206.392 382.863l-51.883 51.882V17.177h-32v417.568l-51.882-51.882L48 405.49L138.509 496l90.51-90.51zm84.921 74.314h144v-32H326.274l109.039-100.732v-43.268h-136v32h101.04l-109.04 100.732zm52.468-408l-58.666 176h33.73l18.667-56h59.6l18.666 56h33.731l-58.666-176Zm4.4 88l18.667-56h.935l18.667 56Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortAlphaUp(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m75.313 130.313l51.883-51.881V496h32V78.432l51.883 51.882l22.627-22.627l-90.51-90.51l-90.509 90.51zM440 280H304v32h101.04L296 412.732V456h144v-32H330.96L440 323.268zM395.532 48h-47.064L289.8 224h33.73l18.67-56h59.6l18.667 56H454.2Zm-42.667 88l18.667-56h.936l18.667 56Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortAscending(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m43.314 130.51l51.882-51.883v417.569h32V78.627l51.883 51.883l22.627-22.627l-90.51-90.511l-90.51 90.511z"/><path fill="currentColor" d="M184 160h120v32H184zm0 72h184v32H184zm0 72h248v32H184zm0 72h312v32H184z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortDescending(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M184 288h120v32H184zm0-72h184v32H184zm0-72h248v32H184zm0-72h312v32H184z"/><path fill="currentColor" d="M95.196 16v417.568l-51.883-51.882l-22.626 22.627l90.509 90.51l90.509-90.51l-22.627-22.626l-51.882 51.881V16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortNumericDown(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M151.196 16v417.568l-51.883-51.881l-22.626 22.626l90.509 90.51l90.51-90.51l-22.627-22.627l-51.883 51.882V16zM432 200h-40V56h-32v32h-32v32h32v80h-40v32h112zm-76.8 232H336v32h19.2a76.887 76.887 0 0 0 76.8-76.8v-51.6a55.663 55.663 0 0 0-55.6-55.6H372a55.663 55.663 0 0 0-55.6 55.6v4.4a55.663 55.663 0 0 0 55.6 55.6h4.4a55.262 55.262 0 0 0 23.474-5.215A44.849 44.849 0 0 1 355.2 432m21.2-68.4H372a23.627 23.627 0 0 1-23.6-23.6v-4.4A23.627 23.627 0 0 1 372 312h4.4a23.627 23.627 0 0 1 23.6 23.6v4.4a23.627 23.627 0 0 1-23.6 23.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortNumericUp(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m102.627 130.313l51.882-51.881V496h32V78.432l51.883 51.882l22.627-22.627l-90.51-90.51L80 107.687zM435.313 200h-40V56h-32v32h-32v32h32v80h-40v32h112zm-76.799 232h-19.2v32h19.2a76.886 76.886 0 0 0 76.8-76.8v-51.6a55.662 55.662 0 0 0-55.6-55.6h-4.4a55.663 55.663 0 0 0-55.6 55.6v4.4a55.663 55.663 0 0 0 55.6 55.6h4.4a55.263 55.263 0 0 0 23.475-5.215A44.85 44.85 0 0 1 358.514 432m21.2-68.4h-4.4a23.627 23.627 0 0 1-23.6-23.6v-4.4a23.627 23.627 0 0 1 23.6-23.6h4.4a23.626 23.626 0 0 1 23.6 23.6v4.4a23.626 23.626 0 0 1-23.601 23.6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spa(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M382.988 237.57a251.854 251.854 0 0 0-102.8-180.91L251.657 36l-26.735 20.79a250.836 250.836 0 0 0-96.279 180.71A176.226 176.226 0 0 0 96 234.451H16V300c0 97.047 78.953 176 176 176h128c97.047 0 176-78.953 176-176v-65.549h-80a176.161 176.161 0 0 0-33.012 3.119M244.568 82.05l7.775-6.05l9.08 6.575a219.732 219.732 0 0 1 90.163 164.079A177.028 177.028 0 0 0 256 337.171a177.022 177.022 0 0 0-95.824-90.6A217.523 217.523 0 0 1 244.568 82.05M240 444h-48c-79.4 0-144-64.6-144-144v-33.549h48c79.4 0 144 64.6 144 144Zm224-144c0 79.4-64.6 144-144 144h-48v-33.549c0-79.4 64.6-144 144-144h48Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpaceBar(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M40 288v128h424V288h-32v96H72v-96z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Speak(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M410.7 286.969c.428-.853.848-1.712 1.252-2.564L456 191.915v-22.891l-3.063-4.209c-.579-.794-58.045-79.741-77.516-105.7c-21.563-28.752-42.328-41.573-67.329-41.573c-19.5 0-39.3 9.269-58.825 27.549A169.483 169.483 0 0 0 236 59.063a169.483 169.483 0 0 0-13.267-13.973c-19.528-18.28-39.32-27.549-58.825-27.549c-25 0-45.766 12.821-67.329 41.573c-19.471 25.96-76.937 104.907-77.516 105.7L16 169.024v22.891l44.037 92.477q.617 1.3 1.262 2.583a134.918 134.918 0 0 0 50.722 54.836a138.545 138.545 0 0 0 71.9 20.065h104.156a138.545 138.545 0 0 0 71.9-20.065a134.924 134.924 0 0 0 50.723-54.842m-226.777 42.907c-39.666 0-76.572-22.473-94.02-57.247a98.988 98.988 0 0 1-.968-1.982L48 184.685v-5.246c12.2-16.749 57.436-78.8 74.179-101.126c19.015-25.354 31.765-28.772 41.729-28.772c23.113 0 47.41 28.439 54.806 39.374L223.468 96h25.064l4.754-7.085c6.854-10.215 31.634-39.374 54.806-39.374c9.964 0 22.714 3.418 41.729 28.773C366.563 100.637 411.8 162.69 424 179.439v5.246l-40.941 85.976c-.314.66-.635 1.317-.959 1.962c-17.451 34.78-54.357 57.253-94.023 57.253Z"/><path fill="currentColor" d="m304.429 112.851l-10.073 11.394c-.233.265-24.313 26.43-58.356 26.43c-34.158 0-58.166-26.219-58.356-26.43l-10.073-11.394l-83.641 66.688l44.393 89.547l11.457-1.637a680.122 680.122 0 0 1 192.44 0l11.457 1.637l44.393-89.547Zm20.815 121.373a711.9 711.9 0 0 0-178.488 0l-22.686-45.763l40.842-32.561c13.575 11.028 38.644 26.778 71.088 26.778s57.513-15.75 71.088-26.778l40.842 32.564Zm17.526 149.198v32a110.961 110.961 0 0 0 107-82.1l-30.92-8.285a78.9 78.9 0 0 1-76.08 58.385"/><path fill="currentColor" d="M480.679 341.605c-16.325 60.868-71.962 105.817-137.909 105.817v32c80.729 0 148.837-55.024 168.82-129.534Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Speaker(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 16H128a32.036 32.036 0 0 0-32 32v416a32.036 32.036 0 0 0 32 32h256a32.036 32.036 0 0 0 32-32V48a32.036 32.036 0 0 0-32-32m0 448H128V48h256l.021 416Z"/><path fill="currentColor" d="M256 240a96 96 0 1 0 96 96a96.108 96.108 0 0 0-96-96m0 160a64 64 0 1 1 64-64a64.072 64.072 0 0 1-64 64m0-200a64 64 0 1 0-64-64a64.072 64.072 0 0 0 64 64m0-96a32 32 0 1 1-32 32a32.036 32.036 0 0 1 32-32"/><path fill="currentColor" d="M240 320h32v32h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Speech(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M144 240h32v32h-32zm96 0h32v32h-32zm96 0h32v32h-32z"/><path fill="currentColor" d="M464 32H48a32.036 32.036 0 0 0-32 32v288a32.036 32.036 0 0 0 32 32h64v112h30.627l112-112H464a32.036 32.036 0 0 0 32-32V64a32.036 32.036 0 0 0-32-32m0 320H241.373L144 449.373V352H48V64h416Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Speedometer(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M425.706 142.294A240 240 0 0 0 16 312v88h144v-32H48v-56c0-114.691 93.309-208 208-208s208 93.309 208 208v56H352v32h144v-88a238.432 238.432 0 0 0-70.294-169.706"/><path fill="currentColor" d="M80 264h32v32H80zm160-136h32v32h-32zm-104 40h32v32h-32zm264 96h32v32h-32zm-102.778 71.1l69.2-144.173l-28.85-13.848l-69.183 144.135a64.141 64.141 0 1 0 28.833 13.886M256 416a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spreadsheet(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 64v392h480V64Zm448 360H48V96h416Z"/><path fill="currentColor" d="M88 136h88v32H88zm0 72h88v32H88zm0 72h88v32H88zm0 72h88v32H88zm200-216h136v32H288zm0 72h136v32H288zm0 72h136v32H288zm0 72h136v32H288zm-72-216h32v32h-32zm0 72h32v32h-32zm0 72h32v32h-32zm0 72h32v32h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Square(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M476 16H36a20.023 20.023 0 0 0-20 20v440a20.023 20.023 0 0 0 20 20h440a20.023 20.023 0 0 0 20-20V36a20.023 20.023 0 0 0-20-20m-12 448H48V48h416Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Star(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M494 198.671a40.536 40.536 0 0 0-32.174-27.592l-115.909-18.837l-53.732-104.414a40.7 40.7 0 0 0-72.37 0l-53.732 104.414l-115.907 18.837a40.7 40.7 0 0 0-22.364 68.827l82.7 83.368l-17.9 116.055a40.672 40.672 0 0 0 58.548 42.538L256 428.977l104.843 52.89a40.69 40.69 0 0 0 58.548-42.538l-17.9-116.055l82.7-83.368A40.538 40.538 0 0 0 494 198.671m-32.53 18.7L367.4 312.2l20.364 132.01a8.671 8.671 0 0 1-12.509 9.088L256 393.136L136.744 453.3a8.671 8.671 0 0 1-12.509-9.088L144.6 312.2l-94.069-94.83a8.7 8.7 0 0 1 4.778-14.706l131.841-21.426l61.119-118.767a8.694 8.694 0 0 1 15.462 0l61.119 118.767l131.841 21.426a8.7 8.7 0 0 1 4.778 14.706Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarHalf(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M135.539 453.9a7.979 7.979 0 0 1-11.51-8.362L144.6 312.2l-95.02-95.789a8 8 0 0 1 4.4-13.53l133.17-21.643L256 76.2V17.833l-18.763 28.624l-69.126 105.455L48.843 171.3a39.847 39.847 0 0 0-31.626 27.122A40.52 40.52 0 0 0 16 203.183v15.276a39.894 39.894 0 0 0 10.862 20.488l83.65 84.327L92.4 440.663a39.979 39.979 0 0 0 57.548 41.812L256 428.977v-35.841Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Storage(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 48v416h480V48Zm448 384H48v-96h416Zm0-128H48v-96h416ZM48 176V80h416v96Z"/><path fill="currentColor" d="M80 112h32v32H80zm0 128h32v32H80zm0 128h32v32H80z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stream(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 112h336v32H16zm144 128h336v32H160zM16 368h336v32H16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Strikethrough(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 496h480V16H16ZM48 48h416v416H48Z"/><path fill="currentColor" d="M344.766 288h-87.741l38.054 25.748a21.894 21.894 0 0 1 9.558 16.187a20.653 20.653 0 0 1-6.058 16.824C294.7 350.584 286.705 357 276.677 357h-104v54h104c21.722 0 42.972-9.165 59.835-25.808a74.638 74.638 0 0 0 21.9-60.218A75.831 75.831 0 0 0 344.766 288m-129.34-93.347c-12.195-8.251-9.725-20.755-8.677-24.389c1.269-4.4 5.647-14.643 18.964-15.248c6.468-.018 118.281 0 118.281 0l.012-54s-119.176-.006-119.795.015c-32.949 1.1-60.171 22.419-69.351 54.278c-7.125 24.726-1.819 49.847 13.481 68.688H80v32h352V224H258.8Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sun(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 104c-83.813 0-152 68.187-152 152s68.187 152 152 152s152-68.187 152-152s-68.187-152-152-152m0 272a120 120 0 1 1 120-120a120.136 120.136 0 0 1-120 120M240 16h32v48h-32zm0 432h32v48h-32zm208-208h48v32h-48zm-432 0h48v32H16zm372.687 171.314l22.627-22.627l32 32l-22.627 22.627zm-320-320l22.628-22.628l32 32l-22.628 22.628zm-.002 329.375l32-32l22.628 22.626l-32 32zm320.002-320.003l32-32l22.628 22.628l-32 32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwapHorizontal(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m364.118 67.313l69.255 69.255H160v32h273.373l-69.255 69.255l22.627 22.627l107.883-107.882L386.745 44.687zM147.882 267.882l-22.627-22.627L17.373 353.137L125.255 461.02l22.627-22.627l-69.255-69.256H352v-32H78.627z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwapVertical(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M384 433.373V160h-32v274.51l-69.823-69.823l-22.627 22.626l107.882 107.883l107.881-107.883l-22.626-22.626zM159.432 17.372L51.55 125.255l22.627 22.627L144 78.059V352h32V79.195l68.687 68.687l22.626-22.627z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Swimming(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M223.308 137.5a16.051 16.051 0 0 1 17.922 3.029l32.883 35.171l-102.055 90.043c9.554 1.9 16.308 5.041 22.132 7.758c7.78 3.631 13.925 6.5 27.517 6.5s19.738-2.868 27.519-6.5l.154-.072a173.415 173.415 0 0 0 41.417-27.1l27.048-23.866l52.9 56.573c6.179-1.137 10.561-3.17 15.624-5.532a99.961 99.961 0 0 1 16.907-6.533L264.388 118.444l-.373-.385a48.151 48.151 0 0 0-54.239-9.555l-64.542 30.12l13.532 29Z"/><circle cx="372" cy="132" r="36" fill="currentColor"/><path fill="currentColor" d="M427.425 376c-20.693 0-31.983 5.268-41.053 9.5c-7.782 3.631-13.928 6.5-27.523 6.5s-19.739-2.868-27.52-6.5c-9.071-4.232-20.359-9.5-41.052-9.5s-31.981 5.269-41.051 9.5c-7.781 3.631-13.926 6.5-27.519 6.5s-19.737-2.868-27.517-6.5c-9.07-4.232-20.358-9.5-41.05-9.5s-31.981 5.269-41.051 9.5c-7.781 3.631-13.926 6.5-27.519 6.5s-19.738-2.868-27.519-6.5C47.981 381.269 36.692 376 16 376v32c13.593 0 19.738 2.868 27.519 6.5c9.07 4.232 20.359 9.5 41.051 9.5s31.981-5.269 41.051-9.5c7.781-3.631 13.926-6.5 27.519-6.5s19.737 2.868 27.517 6.5c9.07 4.232 20.358 9.5 41.05 9.5s31.981-5.269 41.051-9.5c7.781-3.631 13.927-6.5 27.519-6.5s19.739 2.868 27.52 6.5c9.071 4.232 20.359 9.5 41.052 9.5s31.983-5.268 41.054-9.5c7.781-3.631 13.928-6.5 27.522-6.5s19.741 2.868 27.521 6.5c9.072 4.233 20.361 9.5 41.054 9.5v-32c-13.594 0-19.741-2.868-27.521-6.5c-9.072-4.232-20.361-9.5-41.054-9.5m65.875-64.043c-11.741-.359-17.576-3.075-24.82-6.456c-8.666-4.044-19.376-9.02-38.353-9.458c-.889-.02-1.775-.043-2.7-.043s-1.812.023-2.7.043c-18.977.438-29.688 5.414-38.352 9.458c-7.246 3.381-13.08 6.1-24.822 6.456c-.869.026-1.764.043-2.7.043s-1.832-.017-2.7-.043c-11.741-.359-17.575-3.075-24.819-6.456c-8.665-4.043-19.375-9.019-38.351-9.458c-.889-.02-1.774-.043-2.7-.043s-1.811.023-2.7.043c-18.976.438-29.687 5.415-38.351 9.458c-7.244 3.381-13.078 6.1-24.818 6.456c-.869.026-1.764.043-2.7.043s-1.831-.017-2.7-.043c-11.739-.359-17.573-3.075-24.816-6.456c-8.664-4.043-19.374-9.02-38.349-9.458a93.277 93.277 0 0 0-1.95-.033c-.255 0-.5-.012-.753-.012c-.926 0-1.812.023-2.7.043c-18.975.438-29.686 5.415-38.35 9.458c-7.245 3.381-13.078 6.1-24.818 6.456c-.869.026-1.764.043-2.7.043s-1.831-.017-2.7-.043C70.13 311.6 64.3 308.882 57.051 305.5c-8.664-4.043-19.374-9.02-38.35-9.458c-.889-.02-1.775-.043-2.7-.043v32c13.593 0 19.738 2.868 27.519 6.5c9.07 4.232 20.359 9.5 41.051 9.5s31.981-5.269 41.051-9.5c7.781-3.631 13.926-6.5 27.519-6.5s19.737 2.868 27.517 6.5c9.07 4.232 20.358 9.5 41.05 9.5s31.981-5.269 41.051-9.5c7.781-3.631 13.927-6.5 27.519-6.5s19.739 2.868 27.52 6.5c9.071 4.232 20.359 9.5 41.052 9.5s31.983-5.268 41.054-9.5c7.781-3.631 13.928-6.5 27.522-6.5s19.741 2.868 27.521 6.5c9.072 4.233 20.361 9.5 41.054 9.5V312c-.938 0-1.833-.017-2.701-.043"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sync(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m410.168 133.046l-28.958-28.958l82.807-.088l-.034-32L328 72.144V208h32v-79.868l27.541 27.541A152.5 152.5 0 0 1 279.972 416l.056 32a184.5 184.5 0 0 0 130.14-314.954M232.028 104l-.056-32a184.5 184.5 0 0 0-130.14 314.954L130.878 416H48v32h136V312h-32v79.868l-27.541-27.541A152.5 152.5 0 0 1 232.028 104"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tablet(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M440 16H80a24.028 24.028 0 0 0-24 24v432a24.028 24.028 0 0 0 24 24h360a24.028 24.028 0 0 0 24-24V40a24.028 24.028 0 0 0-24-24m-8 448H88v-96h344Zm0-128H88V48h344Z"/><path fill="currentColor" d="M232 400h32v32h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tag(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M485.887 263.261L248 25.373A31.791 31.791 0 0 0 225.373 16H64a48.055 48.055 0 0 0-48 48v161.078A32.115 32.115 0 0 0 26.091 248.4l253.061 237.725a23.815 23.815 0 0 0 16.41 6.51q.447 0 .9-.017a23.828 23.828 0 0 0 16.79-7.734l173.329-188.405a23.941 23.941 0 0 0-.694-33.218M295.171 457.269L48 225.078V64a16.019 16.019 0 0 1 16-16h161.373l232.461 232.462Z"/><path fill="currentColor" d="M148 96a52 52 0 1 0 52 52a52.059 52.059 0 0 0-52-52m0 72a20 20 0 1 1 20-20a20.023 20.023 0 0 1-20 20"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tags(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M511.974 271.891a47.744 47.744 0 0 0-14.706-33.008L311.57 57.98a29.9 29.9 0 0 0-21.2-8.731h-33.142l217.754 212.6l.092.088a15.907 15.907 0 0 1 .741 22.337l-156.4 173.355l21.953 21.356L499.447 305.85a47.748 47.748 0 0 0 12.527-33.959"/><path fill="currentColor" d="M426.9 283.161a23.924 23.924 0 0 0-5.231-24.742c-.106-.111-.2-.231-.306-.34L224.307 57.716l-.094-.094a31.791 31.791 0 0 0-22.628-9.373H60.194A44.244 44.244 0 0 0 16 92.443v141.1a32.121 32.121 0 0 0 10.045 23.28l210.32 200.2l.077.073a23.817 23.817 0 0 0 16.409 6.508q.447 0 .9-.017a24.111 24.111 0 0 0 6.741-1.217a23.854 23.854 0 0 0 10.047-6.517l151.444-164.621a24.033 24.033 0 0 0 4.917-8.071M252.5 428.2L48.077 233.612L48 233.54V92.443a12.208 12.208 0 0 1 12.194-12.194h141.39l191.7 194.918Z"/><path fill="currentColor" d="M140 112a52 52 0 1 0 52 52a52.059 52.059 0 0 0-52-52m0 72a20 20 0 1 1 20-20a20.023 20.023 0 0 1-20 20"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Task(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m222.085 235.644l-62.01-62.01L81.8 251.905l62.009 62.01l-.04.04l66.958 66.957l11.354 11.275l.04.039l66.957-66.957l11.273-11.354l202.277-202.271l-78.272-78.271Zm44.33 66.958l-11.274 11.353l-33.057 33.056l-.04-.039l-33.017-33.017l.04-.04l-62.009-62.01l33.016-33.016l62.01 62.009L424.356 78.627l33.017 33.017Z"/><path fill="currentColor" d="M448 464H48V64h300.22l32-32H16v464h464V179.095l-32 32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Taxi(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M112 304v-96h40v-32H40v32h40v96zm221.483 0L356 269.358L378.517 304h38.166l-41.6-64l41.6-64h-38.166L356 210.642L333.483 176h-38.166l41.6 64l-41.6 64zM440 176h32v128h-32zM40 104h432v32H40zm0 240h432v32H40zm201.337-64l8 24h33.731L240.4 176h-45.952l-39.439 128h33.484l7.4-24Zm-23.617-70.854L230.671 248h-24.923Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tennis(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M144 80a64 64 0 1 0-64 64a64.072 64.072 0 0 0 64-64m-96 0a32 32 0 1 1 32 32a32.036 32.036 0 0 1-32-32m392.65-8.564a133.367 133.367 0 0 0-94.233-39.348h-.622c-37.49 0-74.1 15.969-103.135 45.007c-29.579 29.58-53.748 74.529-64.652 120.24a217.034 217.034 0 0 0-5.459 34.69a119.932 119.932 0 0 1-15.265 51.463A142.364 142.364 0 0 1 134.593 313l-79.285 81.211c-12.389 12.431-13.708 25.214-12.626 33.756c1.254 9.919 6.525 19.771 15.243 28.49S76.5 470.447 86.416 471.7a37.826 37.826 0 0 0 4.754.3c8.188 0 18.755-2.679 29.074-13l79.278-81.2a142.035 142.035 0 0 1 31.969-24.044a118.7 118.7 0 0 1 48.6-13.943a216.365 216.365 0 0 0 34.886-5.562c45.544-10.991 90.409-35.227 120.011-64.83c29.2-29.2 45.182-66.048 45.005-103.757a133.37 133.37 0 0 0-39.343-94.228m-343.094 365c-3.944 3.917-6.256 3.625-7.128 3.517c-2.729-.346-6.328-2.577-9.875-6.124s-5.779-7.146-6.124-9.876c-.111-.877-.405-3.2 3.577-7.186l63.374-64.912l19.61 19.61Zm86.05-87.61l-19.972-19.971a173.6 173.6 0 0 0 21.186-29.091a72.769 72.769 0 0 0 27.88 27.875a173.632 173.632 0 0 0-29.094 21.188ZM412.366 246.8c-25.234 25.233-65.426 46.826-104.891 56.35c-39.275 9.477-72.175 5.684-88.012-10.15c-33.763-33.765-8.923-138.532 45.825-193.28c25.064-25.064 53.5-35.508 80.192-35.508a103.213 103.213 0 0 1 72.543 29.851c35.65 35.65 45.783 101.295-5.657 152.737"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TennisBall(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M425.706 86.294A240 240 0 0 0 86.294 425.706A240 240 0 0 0 425.706 86.294m-143.113 21.431c9.24-17.668 11.486-36.562 6.711-57.046c88.067 14.234 157.784 83.95 172.018 172.017c-20.485-4.774-39.379-2.529-57.047 6.711c-26.085 13.641-45.371 40.8-64.022 67.071c-15.922 22.423-32.385 45.61-52.487 59.911c-21.348 15.189-44.557 18.6-73.038 10.723c-23.426-6.474-43.407-26.086-43.582-26.258c-.2-.2-19.823-20.295-26.258-43.581c-7.872-28.482-4.465-51.691 10.723-73.039c14.3-20.1 37.488-36.565 59.911-52.487c26.268-18.647 53.43-37.937 67.071-64.022M256 464c-114.691 0-208-93.309-208-208c0-114.449 92.917-207.6 207.276-207.991C271.1 90.1 244.08 112.224 197 145.655C149.617 179.3 90.653 221.163 114.044 305.8c8.764 31.706 33.43 56.637 34.476 57.683s25.98 25.709 57.68 34.473a137.081 137.081 0 0 0 36.59 5.291c60.9 0 95.156-48.25 123.553-88.243c33.432-47.085 55.556-74.106 97.646-58.28C463.6 371.083 370.449 464 256 464"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Terminal(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M472 40H40a24.028 24.028 0 0 0-24 24v392a24.028 24.028 0 0 0 24 24h432a24.028 24.028 0 0 0 24-24V64a24.028 24.028 0 0 0-24-24m-8 408H48V72h416Z"/><path fill="currentColor" d="m115.962 282.627l73.445-82.672l-73.451-82.588l-23.912 21.266l54.549 61.333l-54.555 61.407zM216 240h128v32H216z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Terrain(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m40.841 312l103.652-112.88l71.904 71.904l76.29 76.289l22.626-22.626l-77.069-77.07l89.494-95.887L470.836 312H496v-19.864L328.262 104.27L215.603 224.976l-72.096-72.096L16 291.741V312zM16 392h480v32H16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextIcon(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M440 168V88H72v80h64v-48h88v280h-56v32h176v-32h-56V120h88v48h32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextShapes(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m231.359 147l-80.921 205h45.155l15.593-39.5h89.628l15.593 39.5h45.155l-80.921-205Zm-3.594 123.5L256 198.967l28.235 71.533Z"/><path fill="currentColor" d="M384 56H128V16H16v112h40v256H16v112h112v-40h256v40h112V384h-40V128h40V16H384ZM48 96V48h48v48Zm48 368H48v-48h48Zm288-40H128v-40H88V128h40V88h256v40h40v256h-40Zm80-8v48h-48v-48ZM416 48h48v48h-48Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextSize(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M176 184h32v-48h96v232h-40v32h144v-32h-40V136h96v48h32v-80H176z"/><path fill="currentColor" d="M16 280h32v-32h56v152H72v32h112v-32h-32V248h64v32h32v-64H16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextSquare(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 496h480V16H16ZM48 48h416v416H48Z"/><path fill="currentColor" d="M112 176h32v-32h80v224h-32v32h128v-32h-32V144h80v32h32v-64H112z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextStrike(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 496h480V16H16ZM48 48h416v416H48Z"/><path fill="currentColor" d="M344.766 288h-87.741l38.054 25.748a21.894 21.894 0 0 1 9.558 16.187a20.653 20.653 0 0 1-6.058 16.824C294.7 350.584 286.705 357 276.677 357h-104v54h104c21.722 0 42.972-9.165 59.835-25.808a74.638 74.638 0 0 0 21.9-60.218A75.831 75.831 0 0 0 344.766 288m-129.34-93.347c-12.195-8.251-9.725-20.755-8.677-24.389c1.269-4.4 5.647-14.643 18.964-15.248c6.468-.018 118.281 0 118.281 0l.012-54s-119.176-.006-119.795.015c-32.949 1.1-60.171 22.419-69.351 54.278c-7.125 24.726-1.819 49.847 13.481 68.688H80v32h352V224H258.8Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThreeD(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m68.983 382.642l171.35 98.928a32.082 32.082 0 0 0 32 0l171.352-98.929a32.093 32.093 0 0 0 16-27.713V157.071a32.092 32.092 0 0 0-16-27.713L272.334 30.429a32.086 32.086 0 0 0-32 0L68.983 129.358a32.09 32.09 0 0 0-16 27.713v197.858a32.09 32.09 0 0 0 16 27.713M272.333 67.38l155.351 89.691v177.378l-155.351-87.807Zm-16.051 206.947l157.155 88.828l-157.1 90.7l-157.158-90.73ZM84.983 157.071l155.35-89.691v179.2l-155.35 87.81Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbDown(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M127.475 104H51.428a24.086 24.086 0 0 0-23.511 19.176l-1.5 7.31a503.561 503.561 0 0 0 1.477 210.663A24.143 24.143 0 0 0 51.334 360h76.141a24.028 24.028 0 0 0 24-24V128a24.027 24.027 0 0 0-24-24m-8 224H57.811a471.525 471.525 0 0 1-.046-191.082l.188-.918h61.522Zm373.36-109.526L421.583 93.246a24.247 24.247 0 0 0-21.036-12.236h-131.6a24.212 24.212 0 0 0-12.246 3.327l-72.012 42.244v37.1l86.376-50.671h124.947L464 232.5v9.271L444.3 294.3H316l-16 16v48.979l1.418 6.585l10.991 24.341A56.141 56.141 0 0 1 290.7 461.09L207.981 280H184v24l81.007 177.854a24.313 24.313 0 0 0 22.1 14.126a23.923 23.923 0 0 0 9.663-2.034a88.117 88.117 0 0 0 44.8-116.911l-9.57-21.2V326.3h117.7a24.314 24.314 0 0 0 22.661-15.7l22.09-58.906a24.13 24.13 0 0 0 1.542-8.5V230.44a24.256 24.256 0 0 0-3.158-11.966"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbUp(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M126.638 202.672H51.986a24.692 24.692 0 0 0-24.242 19.434a487.088 487.088 0 0 0-1.466 206.535l1.5 7.189a24.94 24.94 0 0 0 24.318 19.78h74.547a24.866 24.866 0 0 0 24.837-24.838V227.509a24.865 24.865 0 0 0-24.842-24.837m-7.163 220.938H57.916l-.309-1.487a455.085 455.085 0 0 1 .158-187.451h61.71Zm374.984-146.326l-22.09-58.906a24.315 24.315 0 0 0-22.662-15.706H332v-29.535l9.573-21.2a88.117 88.117 0 0 0-44.801-116.912a24.3 24.3 0 0 0-31.767 12.1l-80.312 175.812V248h23.731L290.7 67.882a56.141 56.141 0 0 1 21.711 70.885l-10.991 24.341l-1.42 6.584v48.98l16 16h128.3L464 287.2v9.272l-67.988 119.49H271.07l-86.377-50.67v37.1l72.007 42.241a24.222 24.222 0 0 0 12.25 3.329h131.6a24.246 24.246 0 0 0 21.035-12.234l71.25-125.228A24.26 24.26 0 0 0 496 298.531v-12.748a24.144 24.144 0 0 0-1.541-8.499"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToggleOff(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M368 128H144a128 128 0 0 0 0 256h224a128 128 0 0 0 0-256m0 224H144a96 96 0 0 1 0-192h224a96 96 0 0 1 0 192"/><path fill="currentColor" d="M144 192a64 64 0 1 0 64 64a64.072 64.072 0 0 0-64-64m0 96a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToggleOn(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M368 128H144a128 128 0 0 0 0 256h224a128 128 0 0 0 0-256m0 224H144a96 96 0 0 1 0-192h224a96 96 0 0 1 0 192"/><path fill="currentColor" d="M368 192a64 64 0 1 0 64 64a64.072 64.072 0 0 0-64-64m0 96a32 32 0 1 1 32-32a32.036 32.036 0 0 1-32 32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Toilet(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469.613 46.229A77.38 77.38 0 0 0 408 16h-12.381l-56 216H40v32c0 46.482 18.616 88.125 52.417 117.257C124.4 408.82 168.288 424 216 424h38.124l-8 72H448V250.04l35.411-136.585a77.379 77.379 0 0 0-13.798-67.226M416 464H281.876l8-72H216c-84.785 0-144-52.636-144-128h344Zm36.435-358.576L419.619 232h-46.942l47.3-182.436a45.952 45.952 0 0 1 32.46 55.86Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TouchApp(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M454.423 278.957L328 243.839v-8.185a116 116 0 1 0-104 0V312h-24.418l-18.494-22.6a90.414 90.414 0 0 0-126.43-13.367a20.862 20.862 0 0 0-8.026 33.47L215.084 496H472V302.08a24.067 24.067 0 0 0-17.577-23.123M192 132a84 84 0 1 1 136 65.9V132a52 52 0 0 0-104 0v65.9a83.866 83.866 0 0 1-32-65.9m248 332H229.3L79.141 297.75a58.438 58.438 0 0 1 77.181 11.91l28.1 34.34H256V132a20 20 0 0 1 40 0v136.161l144 40Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Transfer(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m356.687 228.687l22.626 22.626L494.627 136L379.313 20.687l-22.626 22.626L433.372 120H16v32h417.372zM496 360H78.628l76.685-76.687l-22.626-22.626L17.373 376l115.314 115.313l22.626-22.626L78.628 392H496z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Translate(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m230.2 356.821l14.425-30.654a573.693 573.693 0 0 1-40.421-44.777q15.881-21.418 29.594-43.958A543.369 543.369 0 0 0 292.212 104H344V72H200V16h-32v56H16v32h59.77a545.123 545.123 0 0 0 71.448 153.959a562.586 562.586 0 0 0 20.566-27.031A512.677 512.677 0 0 1 109.13 104h149.737c-29.727 97.53-84.546 169.208-126.64 213.119c-48.993 51.107-91.952 76.86-92.38 77.114l1.621 2.738L48 408l8.14 13.774c1.873-1.106 46.474-27.729 98.389-81.68q15.38-15.982 29.44-32.931a608.138 608.138 0 0 0 46.231 49.658M333.722 200L328 212.516L209.379 472h35.185l36.571-80h141.73l36.571 80h35.185L370.278 200ZM328 360h-32.236L328 289.484l24-52.5L408.236 360Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trash(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M96 472a23.82 23.82 0 0 0 23.579 24h272.842A23.82 23.82 0 0 0 416 472V152H96Zm32-288h256v280H128Z"/><path fill="currentColor" d="M168 216h32v200h-32zm72 0h32v200h-32zm72 0h32v200h-32zm16-128V40c0-13.458-9.488-24-21.6-24H205.6C193.488 16 184 26.542 184 40v48H64v32h384V88ZM216 48h80v40h-80Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Triangle(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M273.342 66.038a20 20 0 0 0-34.684 0L29.569 430.007a20 20 0 0 0 17.342 29.963h418.178a20 20 0 0 0 17.342-29.962ZM67.644 427.97L256 100.091L444.356 427.97Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Truck(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M441.885 141.649A32.028 32.028 0 0 0 415.669 128H336V80a32.036 32.036 0 0 0-32-32H48a32.036 32.036 0 0 0-32 32v328h53.082a67.982 67.982 0 0 0 133.836 0h106.164a67.982 67.982 0 0 0 133.836 0H496V226.522a23.882 23.882 0 0 0-4.338-13.763ZM47.98 80H304v176H48ZM136 432a36 36 0 1 1 36-36a36.04 36.04 0 0 1-36 36m240 0a36 36 0 1 1 36-36a36.04 36.04 0 0 1-36 36m88-56h-23.006a68 68 0 0 0-129.988 0H200.994a68 68 0 0 0-129.988 0H48v-88h416Zm0-120H336v-96h79.669L464 229.044Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tv(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M472 88.5H302.627l72.5-72.5h-45.254l-72.5 72.5h-2.246l-72.5-72.5h-45.254l72.5 72.5H40a24.028 24.028 0 0 0-24 24v296a24.028 24.028 0 0 0 24 24h112V496h224v-63.5h96a24.028 24.028 0 0 0 24-24v-296a24.028 24.028 0 0 0-24-24M344 464H184v-31.5h160Zm120-63.5H48v-280h416Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Underline(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 496h480V16H16ZM48 48h416v416H48Z"/><path fill="currentColor" d="M296 152h32v88a63.966 63.966 0 0 1-88 59.313V152h24v-32H136v32h32v88a96 96 0 0 0 192 0v-88h32v-32h-96ZM136 368h256v32H136z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Usb(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M320 224h32v40h-80V78.627l17.941 17.941l22.627-22.627L256 17.373l-56.568 56.568l22.627 22.627L240 78.627V264h-72v-42.341a56 56 0 1 0-32 0V296h104v90.341a56 56 0 1 0 32 0V296h112v-72h40V120H320Zm-192-56a24 24 0 1 1 24 24a24.027 24.027 0 0 1-24-24m152 272a24 24 0 1 1-24-24a24.027 24.027 0 0 1 24 24m72-288h40v40h-40Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func User(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m411.6 343.656l-72.823-47.334l27.455-50.334A80.23 80.23 0 0 0 376 207.681V128a112 112 0 0 0-224 0v79.681a80.236 80.236 0 0 0 9.768 38.308l27.455 50.333l-72.823 47.334A79.725 79.725 0 0 0 80 410.732V496h368v-85.268a79.727 79.727 0 0 0-36.4-67.076M416 464H112v-53.268a47.836 47.836 0 0 1 21.841-40.246l97.66-63.479l-41.64-76.341A48.146 48.146 0 0 1 184 207.681V128a80 80 0 0 1 160 0v79.681a48.146 48.146 0 0 1-5.861 22.985L296.5 307.007l97.662 63.479A47.836 47.836 0 0 1 416 410.732Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserFemale(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m403.6 343.656l-72.823-47.334L344.043 272h23.428a48 48 0 0 0 44.119-66.908L361.581 90.57a112.029 112.029 0 0 0-211.162 0L100.41 205.092A48 48 0 0 0 144.529 272h23.428l13.266 24.322l-72.823 47.334A79.725 79.725 0 0 0 72 410.732V496h368v-85.268a79.727 79.727 0 0 0-36.4-67.076M408 464H104v-53.268a47.836 47.836 0 0 1 21.841-40.246l97.66-63.479L186.953 240h-42.424a16 16 0 0 1-14.72-22.27l50.172-114.9l.448-1.143a80.029 80.029 0 0 1 151.142 0l.2.58l50.42 115.463a16 16 0 0 1-14.72 22.27h-42.424L288.5 307.007l97.661 63.479A47.836 47.836 0 0 1 408 410.732Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserFollow(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M208 16A112.127 112.127 0 0 0 96 128v79.681a80.236 80.236 0 0 0 9.768 38.308l27.455 50.333L60.4 343.656A79.725 79.725 0 0 0 24 410.732V496h288v-32H56v-53.268a47.836 47.836 0 0 1 21.841-40.246l97.66-63.479l-41.64-76.341A48.146 48.146 0 0 1 128 207.681V128a80 80 0 0 1 160 0v79.681a48.146 48.146 0 0 1-5.861 22.985L240.5 307.007l71.5 46.476v-38.166l-29.223-19l27.455-50.334A80.23 80.23 0 0 0 320 207.681V128A112.127 112.127 0 0 0 208 16m216 384v-64h-32v64h-64v32h64v64h32v-64h64v-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserPlus(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M208 16A112.127 112.127 0 0 0 96 128v79.681a80.236 80.236 0 0 0 9.768 38.308l27.455 50.333L60.4 343.656A79.725 79.725 0 0 0 24 410.732V496h288v-32H56v-53.268a47.836 47.836 0 0 1 21.841-40.246l97.66-63.479l-41.64-76.341A48.146 48.146 0 0 1 128 207.681V128a80 80 0 0 1 160 0v79.681a48.146 48.146 0 0 1-5.861 22.985L240.5 307.007l71.5 46.476v-38.166l-29.223-19l27.455-50.334A80.23 80.23 0 0 0 320 207.681V128A112.127 112.127 0 0 0 208 16m216 384v-64h-32v64h-64v32h64v64h32v-64h64v-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserUnfollow(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M216 16a112.127 112.127 0 0 0-112 112v79.68a80.239 80.239 0 0 0 9.768 38.308l27.455 50.334L68.4 343.656A79.727 79.727 0 0 0 32 410.732V496h280v-32H64v-53.268a47.838 47.838 0 0 1 21.84-40.246l97.66-63.479l-41.64-76.341A48.149 48.149 0 0 1 136 207.68V128a80 80 0 0 1 160 0v79.68a48.143 48.143 0 0 1-5.861 22.985L248.5 307.007l63.5 41.276v-38.166l-21.223-13.8l27.454-50.334A80.226 80.226 0 0 0 328 207.68V128A112.127 112.127 0 0 0 216 16m267.314 339.314l-22.628-22.628L412 381.372l-48.686-48.686l-22.628 22.628L389.372 404l-48.686 48.686l22.628 22.628L412 426.628l48.686 48.686l22.628-22.628L434.628 404z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserX(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M216 16a112.127 112.127 0 0 0-112 112v79.68a80.239 80.239 0 0 0 9.768 38.308l27.455 50.334L68.4 343.656A79.727 79.727 0 0 0 32 410.732V496h280v-32H64v-53.268a47.838 47.838 0 0 1 21.84-40.246l97.66-63.479l-41.64-76.341A48.149 48.149 0 0 1 136 207.68V128a80 80 0 0 1 160 0v79.68a48.143 48.143 0 0 1-5.861 22.985L248.5 307.007l63.5 41.276v-38.166l-21.223-13.8l27.454-50.334A80.226 80.226 0 0 0 328 207.68V128A112.127 112.127 0 0 0 216 16m267.314 339.314l-22.628-22.628L412 381.372l-48.686-48.686l-22.628 22.628L389.372 404l-48.686 48.686l22.628 22.628L412 426.628l48.686 48.686l22.628-22.628L434.628 404z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vector(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M368 184h47.458c-4.664 69.192-39.8 119.633-95.458 140.869V288H192v36.869C136.344 303.633 101.206 253.192 96.542 184H144V56H16v128h48.471c2.993 50.374 20.242 93.872 50.341 126.537A170.65 170.65 0 0 0 166.7 348.9l-48.9 16.3a57.122 57.122 0 0 0-8.608-9.083a56.63 56.63 0 0 0-79.807 6.693A56.635 56.635 0 0 0 72.535 456q2.4 0 4.827-.2a56.609 56.609 0 0 0 51.731-60.634L192 374.2V416h128v-41.8l62.9 20.963a57.4 57.4 0 0 0 .043 8.939A56.64 56.64 0 0 0 439.274 456q2.4 0 4.828-.2a56.63 56.63 0 1 0-47.982-92.988c-.662.782-1.3 1.582-1.91 2.392L345.3 348.9a170.65 170.65 0 0 0 51.891-38.363c30.1-32.665 47.348-76.163 50.341-126.537H496V56H368ZM48 152V88h64v64Zm43.44 263.27a24.629 24.629 0 0 1-37.62-31.8a24.469 24.469 0 0 1 16.752-8.644q1.053-.087 2.1-.088A24.631 24.631 0 0 1 91.44 415.27M288 384h-64v-64h64Zm132.56-.529a24.467 24.467 0 0 1 16.752-8.644q1.053-.087 2.1-.088a24.635 24.635 0 1 1-18.851 8.733ZM400 88h64v64h-64Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VerticalAlignBottom(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M356.117 316.117L333.49 293.49L272 354.98V56h-32v298.98l-61.49-61.49l-22.627 22.627L256 416.236zM16 464h480v32H16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VerticalAlignBottomOne(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M6.75.75h-6v19.875h6zm-1.5 18.375h-3V2.25h3z" fill="currentColor"/><path d="M15 7.948H9v12.677h6zm-1.5 11.177h-3V9.448h3z" fill="currentColor"/><path d="M17.25.75v19.875h6V.75zm4.5 18.375h-3V2.25h3z" fill="currentColor"/><path d="M.75 21.75h22.5v1.5H.75v-1.5z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VerticalAlignCenter(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M240 16v146.177l-53.491-53.49l-22.627 22.626L256 223.431l92.118-92.118l-22.627-22.626L272 162.177V16zM16 240h480v32H16zm147.882 140.687l22.627 22.626L240 349.823V496h32V349.823l53.491 53.49l22.627-22.626L256 288.569z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VerticalAlignCenterOne(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M17.25 23.25h6v-9h-6zm1.5-7.5h3v6h-3z" fill="currentColor"/><path d="M9 19.875h6V14.25H9zm1.5-4.125h3v2.625h-3z" fill="currentColor"/><path d="M.75 23.25h6v-9h-6zm1.5-7.5h3v6h-3z" fill="currentColor"/><path d="M6.75.75h-6v9h6zm-1.5 7.5h-3v-6h3z" fill="currentColor"/><path d="M15 4.5H9v5.25h6zm-1.5 3.75h-3V6h3z" fill="currentColor"/><path d="M17.25.75v9h6v-9zm4.5 7.5h-3v-6h3z" fill="currentColor"/><path d="M.75 11.25h22.5v1.5H.75v-1.5z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VerticalAlignTop(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 16h480v32H16zm139.883 179.883l22.627 22.627L240 157.02V456h32V157.02l61.49 61.49l22.627-22.627L256 95.764z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VerticalAlignTopOne(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M17.25 23.25h6V3.75h-6zm1.5-18h3v16.5h-3z" fill="currentColor"/><path d="M9 16.125h6V3.75H9zM10.5 5.25h3v9.375h-3z" fill="currentColor"/><path d="M.75 23.25h6V3.75h-6zm1.5-18h3v16.5h-3z" fill="currentColor"/><path d="M.75.75h22.5v1.5H.75V.75z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Video(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m461 132l-101 39.277V72H16v368h344V332.723L461 372h35V132Zm3 206.833l-136-52.89V408H48V104h280v114.057l136-52.89Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Videogame(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 360h480V120H16Zm32-208h416v176H48Z"/><path fill="currentColor" d="M152 184h-32v40H80v32h40v40h32v-40h40v-32h-40zm184 72h32v32h-32zm64-64h32v32h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewColumn(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 64v384h480V64Zm304 32v320H192V96ZM48 96h112v320H48Zm416 320H352V96h112Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewModule(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 64v384h480V64Zm448 176H352V96h112Zm-272 0V96h128v144Zm128 32v144H192V272ZM160 96v144H48V96ZM48 272h112v144H48Zm304 144V272h112v144Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewQuilt(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 64v384h480V64Zm448 176H192V96h272Zm-272 32h120v144H192ZM48 96h112v320H48Zm296 320V272h120v144Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewStream(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 64v384h480V64Zm448 32v144H48V96ZM48 416V272h416v144Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Voice(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M410.7 286.969c.428-.853.848-1.712 1.252-2.564L456 191.915v-22.891l-3.063-4.209c-.579-.794-58.045-79.741-77.516-105.7c-21.563-28.752-42.328-41.573-67.329-41.573c-19.5 0-39.3 9.269-58.825 27.549A169.483 169.483 0 0 0 236 59.063a169.483 169.483 0 0 0-13.267-13.973c-19.528-18.28-39.32-27.549-58.825-27.549c-25 0-45.766 12.821-67.329 41.573c-19.471 25.96-76.937 104.907-77.516 105.7L16 169.024v22.891l44.037 92.477q.617 1.3 1.262 2.583a134.918 134.918 0 0 0 50.722 54.836a138.545 138.545 0 0 0 71.9 20.065h104.156a138.545 138.545 0 0 0 71.9-20.065a134.924 134.924 0 0 0 50.723-54.842m-226.777 42.907c-39.666 0-76.572-22.473-94.02-57.247a98.988 98.988 0 0 1-.968-1.982L48 184.685v-5.246c12.2-16.749 57.436-78.8 74.179-101.126c19.015-25.354 31.765-28.772 41.729-28.772c23.113 0 47.41 28.439 54.806 39.374L223.468 96h25.064l4.754-7.085c6.854-10.215 31.634-39.374 54.806-39.374c9.964 0 22.714 3.418 41.729 28.773C366.563 100.637 411.8 162.69 424 179.439v5.246l-40.941 85.976c-.314.66-.635 1.317-.959 1.962c-17.451 34.78-54.357 57.253-94.023 57.253Z"/><path fill="currentColor" d="m304.429 112.851l-10.073 11.394c-.233.265-24.313 26.43-58.356 26.43c-34.158 0-58.166-26.219-58.356-26.43l-10.073-11.394l-83.641 66.688l44.393 89.547l11.457-1.637a680.122 680.122 0 0 1 192.44 0l11.457 1.637l44.393-89.547Zm20.815 121.373a711.9 711.9 0 0 0-178.488 0l-22.686-45.763l40.842-32.561c13.575 11.028 38.644 26.778 71.088 26.778s57.513-15.75 71.088-26.778l40.842 32.564Zm17.526 149.198v32a110.961 110.961 0 0 0 107-82.1l-30.92-8.285a78.9 78.9 0 0 1-76.08 58.385"/><path fill="currentColor" d="M480.679 341.605c-16.325 60.868-71.962 105.817-137.909 105.817v32c80.729 0 148.837-55.024 168.82-129.534Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VoiceOverRecord(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M342.322 104.837A100.443 100.443 0 0 1 342.457 231l24.943 20.02a132.373 132.373 0 0 0-.178-166.261ZM432.147 32.4l-24.989 20.15a183.6 183.6 0 0 1 .248 230.594l25.03 20.1a216.053 216.053 0 0 0-.289-270.844"/><path fill="currentColor" d="m347.6 343.656l-72.822-47.334l27.455-50.334A80.23 80.23 0 0 0 312 207.681V128a112 112 0 0 0-224 0v79.681a80.236 80.236 0 0 0 9.768 38.308l27.455 50.333L52.4 343.656A79.725 79.725 0 0 0 16 410.732V496h368v-85.268a79.725 79.725 0 0 0-36.4-67.076M352 464H48v-53.268a47.836 47.836 0 0 1 21.841-40.246l97.661-63.48l-41.641-76.34A48.146 48.146 0 0 1 120 207.681V128a80 80 0 0 1 160 0v79.681a48.139 48.139 0 0 1-5.861 22.984L232.5 307.007l97.662 63.479A47.838 47.838 0 0 1 352 410.732Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeHigh(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M145.373 120H16v272h129.373l104 104H288V16h-38.627ZM128 360H48V152h80Zm128 97.373l-96-96V150.627l96-96ZM408 256a88.1 88.1 0 0 0-88-88v32a56 56 0 0 1 0 112v32a88.1 88.1 0 0 0 88-88"/><path fill="currentColor" d="M320 57.445v32.277C401.307 101.4 464 171.512 464 256s-62.693 154.6-144 166.278v32.277C419.005 442.66 496 358.158 496 256S419.005 69.34 320 57.445"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeLow(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M320 168v32a56 56 0 0 1 0 112v32a88 88 0 0 0 0-176m-174.627-48H16v272h129.373l104 104H288V16h-38.627ZM128 360H48V152h80Zm128 97.373l-96-96V150.627l96-96Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeOff(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m490.51 187.313l-22.628-22.626l-64.284 64.284l-64.285-64.284l-22.627 22.626l64.285 64.285l-64.285 64.285l22.627 22.627l64.285-64.284l64.284 64.284l22.628-22.627l-64.285-64.285zM145.373 120H16v272h129.373l104 104H288V16h-38.627ZM128 360H48V152h80Zm128 97.373l-96-96V150.627l96-96Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Walk(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="309.912" cy="82.088" r="34.088" fill="currentColor"/><path fill="currentColor" d="m322 143.462l-60.585-20.64l-77.615 28.226A24.073 24.073 0 0 0 168 173.6V232h32v-52.793l48.811-17.749L158.622 440h33.613l50.082-155.811l7.871 2.568L288 356.079V440h32v-85.96a24.068 24.068 0 0 0-2.931-11.493l-37.56-68.861l28.949-88.715l19.27 34.684A24.011 24.011 0 0 0 348.707 232H400v-32h-46.586Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wallet(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M489.972 119.059a23.839 23.839 0 0 0-17-7.059H456V72a24.027 24.027 0 0 0-24-24H86.627A70.628 70.628 0 0 0 16 118.627v274.746A70.628 70.628 0 0 0 86.627 464h385.4a24.047 24.047 0 0 0 24-23.923l.944-303.995a23.837 23.837 0 0 0-6.999-17.023M464.053 432H86.627A38.627 38.627 0 0 1 48 393.373V118.627A38.627 38.627 0 0 1 86.627 80H424v32H88v32h376.947Z"/><path fill="currentColor" d="M392 264h32v32h-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wallpaper(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M344 112h40v40h-40zM72 72h152V40H40v184h32z"/><path fill="currentColor" d="M288 40v32h152v152h32V40zM72 288H40v184h184v-32H72z"/><path fill="currentColor" d="m280.5 308.873l-91-91l-85.5 85.5v45.254l85.5-85.499L334.372 408h45.255l-76.5-76.5l72.104-72.104L440 324.165V440H288v32h184V312l-96.769-97.857z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Warning(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M240 176h32v176h-32zm0 208h32v32h-32z"/><path fill="currentColor" d="M274.014 16h-36.028L16 445.174V496h480v-50.826ZM464 464H48v-11.041L256 50.826l208 402.133Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Watch(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M349.279 16H162.721l-24.772 132.116a159.7 159.7 0 0 0 0 215.768L162.721 496h186.558l24.772-132.116a159.7 159.7 0 0 0 0-215.768Zm-160 32h133.442l13 69.311a159.752 159.752 0 0 0-159.434 0Zm133.442 416H189.279l-13-69.311a159.752 159.752 0 0 0 159.434 0ZM256 384a128 128 0 1 1 128-128a128.145 128.145 0 0 1-128 128"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wc(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M132 139.824a61.912 61.912 0 1 0-61.912-61.912A61.981 61.981 0 0 0 132 139.824M132 48a29.912 29.912 0 1 1-29.912 29.912A29.947 29.947 0 0 1 132 48m44 104H88a48.053 48.053 0 0 0-48 48v152h32v144h120V352h32V200a48.053 48.053 0 0 0-48-48m16 168h-32v144h-56V320H72V200a16.019 16.019 0 0 1 16-16h88a16.019 16.019 0 0 1 16 16Zm178.088-180.176a61.912 61.912 0 1 0-61.912-61.912a61.981 61.981 0 0 0 61.912 61.912m0-91.824a29.912 29.912 0 1 1-29.912 29.912A29.947 29.947 0 0 1 370.088 48m55.671 145.354a61.586 61.586 0 0 0-115.833-1.392L248 357.1V400h64v96h104v-96h64v-50.7ZM448 368h-64v96h-40v-96h-64v-5.1l59.889-159.7a29.585 29.585 0 0 1 55.645.669L448 354.7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Weightlifitng(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="256" cy="156" r="36" fill="currentColor"/><path fill="currentColor" d="M400 16v56H112V16H16v144h96v-56h25.278l7.524 90.289l79.2 39.6V336h-79.138l-7.413 133.426H169.5L175.136 368h161.728l5.636 101.426h32.05L367.136 336H288V233.889l79.2-39.6L374.722 104H400v56h96V16ZM80 72v56H48V48h32Zm256.8 101.71l-80.8 40.4l-80.8-40.4l-5.811-69.71h173.222ZM464 128h-32V48h32Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wheelchair(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="221.912" cy="66.088" r="34.088" fill="currentColor"/><path fill="currentColor" d="m460.12 360.478l-47.943 11.985L393 282.971A24.126 24.126 0 0 0 369.533 264h-88.705l-6.462-56H384v-32H270.674l-4.134-35.826a24 24 0 0 0-26.593-21.091l-39.736 4.585L220.1 296h142.97l24.758 115.537l80.057-20.015Z"/><path fill="currentColor" d="M224 448a120 120 0 0 1-45.248-231.135l-3.779-32.75C115.143 204.558 72 261.334 72 328c0 83.813 68.187 152 152 152a152.06 152.06 0 0 0 130.044-73.378L344 360c-16 48-61.4 88-120 88"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiSignalFour(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M503.785 124.254a432.019 432.019 0 0 0-495.57 0L8 124.4v27.437L237.778 480h36.444L504 151.841V124.4Zm-100.458-18.088L213.968 390.2l-24.076-34.38L362.877 92.583a395.92 395.92 0 0 1 40.45 13.583M286.4 79.278a400.017 400.017 0 0 1 43.232 5.631l-159.49 242.7l-26.536-37.9ZM123.963 261.664l-22.8-32.563l97.97-146.955a402.727 402.727 0 0 1 49.324-3.946Zm30.751-170.579l-73.253 109.88l-41.6-59.418a398.09 398.09 0 0 1 114.853-50.462M256 450.232l-22.331-31.892l199.315-298.972a401.8 401.8 0 0 1 39.16 22.179Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiSignalOff(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M503.785 124.254a431.821 431.821 0 0 0-332.477-69.728L192.9 83.09a402.723 402.723 0 0 1 63.1-4.962a397.867 397.867 0 0 1 216.144 63.419L359.092 303l20.278 26.832L504 151.842V124.4ZM82.542 16v21.9l25.874 34.237a431.346 431.346 0 0 0-100.2 52.117L8 124.4v27.438L237.778 480h36.444l68.5-97.829L404.564 464h22.986v-22.658L106.1 16Zm239.9 339.34L256 450.232L39.856 141.547a398.932 398.932 0 0 1 88.759-42.683Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiSignalOne(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M237.778 480h36.444L504 151.842V124.4l-.215-.15a432.019 432.019 0 0 0-495.57 0L8 124.4v27.438Zm-58.911-139.925a158.219 158.219 0 0 1 154.266 0L256 450.232ZM256 78.128a397.867 397.867 0 0 1 216.144 63.419L351.561 313.758a190.142 190.142 0 0 0-191.122 0L39.856 141.547A397.867 397.867 0 0 1 256 78.128"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiSignalThree(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M503.785 124.254a432.019 432.019 0 0 0-495.57 0L8 124.4v27.438l33 47.126L237.778 480h36.444l196.783-281.036l33-47.122V124.4ZM189.87 355.789l125.621-192.127a317.213 317.213 0 0 1 42.184 10.981L213.968 390.2Zm-19.77-28.236l-26.453-37.778l89.523-130.843q11.354-.8 22.83-.8q12.2 0 24.276.918Zm-46.064-65.786l-38.248-54.622a317.977 317.977 0 0 1 104.534-42.258Zm109.633 156.574l154.47-231.7a320.459 320.459 0 0 1 38.073 20.509L256 450.232Zm210.91-237.427a351.947 351.947 0 0 0-377.158 0l-27.565-39.367a400.039 400.039 0 0 1 432.288 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiSignalTwo(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M503.785 124.254a432.019 432.019 0 0 0-495.57 0L8 124.4v27.438L86.881 264.5L237.778 480h36.444l150.9-215.5L504 151.842V124.4Zm-313.961 231.47l75.96-117.392a240.089 240.089 0 0 1 43.276 5.686l-95.22 146Zm-19.809-28.291l-38.265-54.649a238.218 238.218 0 0 1 94.9-32.873Zm63.606 90.838l107.373-164.639a239.338 239.338 0 0 1 39.256 19.152L256 450.232Zm165.018-171.748a272.034 272.034 0 0 0-285.278 0l-73.5-104.976a400.039 400.039 0 0 1 432.288 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiSignalZero(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M237.778 480h36.444L504 151.842V124.4l-.215-.15a432.019 432.019 0 0 0-495.57 0L8 124.4v27.438ZM256 78.128a397.867 397.867 0 0 1 216.144 63.419L256 450.232L39.856 141.547A397.867 397.867 0 0 1 256 78.128"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Window(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M424 440V48H88v392H48v32h416v-32ZM120 80h120v120H120Zm0 152h120v144H120Zm272 208H120v-32h272Zm0-64H272V232h120Zm0-176H272V80h120Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowMaximize(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M472 48H40.335a24.027 24.027 0 0 0-24 24v384a24.027 24.027 0 0 0 24 24H472a24.027 24.027 0 0 0 24-24V72a24.027 24.027 0 0 0-24-24m-8 32v71.981l-415.665-.491V80ZM48.335 448V183.49l415.665.491V448Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowMinimize(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M48 448h416v32H48z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowRestore(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M352 153H40.247a24.028 24.028 0 0 0-24 24v281a24.028 24.028 0 0 0 24 24H352a24.028 24.028 0 0 0 24-24V177a24.028 24.028 0 0 0-24-24m-8 32v45.22H48.247V185ZM48.247 450V262.22H344V450Z"/><path fill="currentColor" d="M472 32H152a24.028 24.028 0 0 0-24 24v65h32V64h304v275.143h-56v32h64a24.028 24.028 0 0 0 24-24V56a24.028 24.028 0 0 0-24-24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WrapText(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 232h160v32H16zm0 160h392v32H16zM408 72H16v32h392c30.878 0 56 28.71 56 64s-25.122 64-56 64h-63.2v-69.228L210 249l134.8 83.785V264H408c48.523 0 88-43.065 88-96s-39.477-96-88-96m-95.2 203.217L270 248.609l42.8-27.381Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func X(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m427.314 107.313l-22.628-22.626L256 233.373L107.314 84.687l-22.628 22.626L233.373 256L84.686 404.687l22.628 22.626L256 278.627l148.686 148.686l22.628-22.626L278.627 256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Xcircle(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m348.071 141.302l-87.763 87.763l-87.763-87.763l-22.628 22.627l87.764 87.763l-87.764 87.764l22.628 22.627l87.763-87.763l87.763 87.763l22.628-22.627l-87.764-87.764l87.764-87.763z"/><path fill="currentColor" d="M425.706 86.294A240 240 0 0 0 86.294 425.706A240 240 0 0 0 425.706 86.294M256 464c-114.691 0-208-93.309-208-208S141.309 48 256 48s208 93.309 208 208s-93.309 208-208 208"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Yen(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m396.641 97.81l-25.282-19.62l-114.946 148.122L148.938 78.587l-25.876 18.826L238.438 256h-85.967v32H240v34.823h-87.529v32H240V432h32v-77.177h87.529v-32H272V288h87.529v-32h-85.65z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Zoom(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m479.6 399.716l-81.084-81.084l-62.368-25.767A175.014 175.014 0 0 0 368 192c0-97.047-78.953-176-176-176S16 94.953 16 192s78.953 176 176 176a175.034 175.034 0 0 0 101.619-32.377l25.7 62.2l81.081 81.088a56 56 0 1 0 79.2-79.195M48 192c0-79.4 64.6-144 144-144s144 64.6 144 144s-64.6 144-144 144S48 271.4 48 192m408.971 264.284a24.028 24.028 0 0 1-33.942 0l-76.572-76.572l-23.894-57.835l57.837 23.894l76.573 76.572a24.028 24.028 0 0 1-.002 33.941"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomIn(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M208 96h-32v80H96v32h80v80h32v-80h80v-32h-80z"/><path fill="currentColor" d="m479.6 400.4l-81.084-81.084l-62.368-25.767A175.008 175.008 0 0 0 368 192.687c0-97.047-78.953-176-176-176s-176 78.953-176 176s78.953 176 176 176a175.028 175.028 0 0 0 101.619-32.378l25.7 62.2L400.4 479.6a56 56 0 0 0 79.2-79.2M48 192.687c0-79.4 64.6-144 144-144s144 64.6 144 144s-64.6 144-144 144s-144-64.599-144-144m408.971 264.284a24.029 24.029 0 0 1-33.942 0L346.457 380.4l-23.894-57.835l57.837 23.892l76.573 76.572a24.029 24.029 0 0 1-.002 33.942"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomOut(children ...ElementRenderer) *CilIcon {
	return &CilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M96.344 175.313h192v32h-192z"/><path fill="currentColor" d="m479.6 399.716l-81.084-81.084l-62.368-25.767A175.014 175.014 0 0 0 368 192c0-97.047-78.953-176-176-176S16 94.953 16 192s78.953 176 176 176a175.034 175.034 0 0 0 101.619-32.377l25.7 62.2l81.081 81.088a56 56 0 1 0 79.2-79.195M48 192c0-79.4 64.6-144 144-144s144 64.6 144 144s-64.6 144-144 144S48 271.4 48 192m408.971 264.284a24.028 24.028 0 0 1-33.942 0l-76.572-76.572l-23.894-57.835l57.837 23.894l76.573 76.572a24.028 24.028 0 0 1-.002 33.941"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
