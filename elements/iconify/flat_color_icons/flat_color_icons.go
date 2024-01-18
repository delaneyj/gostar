package flat_color_icons

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = "1.0.2"
	hAttr          = "1em"
	viewbox        = "0 0 48 48"
	fill           = "currentColor"
)

type FlatColorIconsIcon struct {
	*SVGSVGElement
}

type FlatColorIconsIconFn func(children ...ElementRenderer) *FlatColorIconsIcon

var IconLookup = map[string]FlatColorIconsIconFn{
	"about":                     About,
	"acceptDatabase":            AcceptDatabase,
	"addColumn":                 AddColumn,
	"addDatabase":               AddDatabase,
	"addImage":                  AddImage,
	"addRow":                    AddRow,
	"addressBook":               AddressBook,
	"advance":                   Advance,
	"advertising":               Advertising,
	"alarmClock":                AlarmClock,
	"alphabeticalSortingAz":     AlphabeticalSortingAz,
	"alphabeticalSortingZa":     AlphabeticalSortingZa,
	"androidOs":                 AndroidOs,
	"answers":                   Answers,
	"approval":                  Approval,
	"approve":                   Approve,
	"areaChart":                 AreaChart,
	"assistant":                 Assistant,
	"audioFile":                 AudioFile,
	"automatic":                 Automatic,
	"automotive":                Automotive,
	"badDecision":               BadDecision,
	"barChart":                  BarChart,
	"bbc":                       Bbc,
	"bearish":                   Bearish,
	"binoculars":                Binoculars,
	"biohazard":                 Biohazard,
	"biomass":                   Biomass,
	"biotech":                   Biotech,
	"bookmark":                  Bookmark,
	"briefcase":                 Briefcase,
	"brokenLink":                BrokenLink,
	"bullish":                   Bullish,
	"business":                  Business,
	"businessContact":           BusinessContact,
	"businessman":               Businessman,
	"businesswoman":             Businesswoman,
	"buttingIn":                 ButtingIn,
	"cableRelease":              CableRelease,
	"calculator":                Calculator,
	"calendar":                  Calendar,
	"callTransfer":              CallTransfer,
	"callback":                  Callback,
	"camcorder":                 Camcorder,
	"camcorderPro":              CamcorderPro,
	"camera":                    Camera,
	"cameraAddon":               CameraAddon,
	"cameraIdentification":      CameraIdentification,
	"cancel":                    Cancel,
	"candleSticks":              CandleSticks,
	"capacitor":                 Capacitor,
	"cdLogo":                    CdLogo,
	"cellPhone":                 CellPhone,
	"chargeBattery":             ChargeBattery,
	"checkmark":                 Checkmark,
	"circuit":                   Circuit,
	"clapperboard":              Clapperboard,
	"clearFilters":              ClearFilters,
	"clock":                     Clock,
	"closeUpMode":               CloseUpMode,
	"cloth":                     Cloth,
	"collaboration":             Collaboration,
	"collapse":                  Collapse,
	"collect":                   Collect,
	"comboChart":                ComboChart,
	"commandLine":               CommandLine,
	"comments":                  Comments,
	"compactCamera":             CompactCamera,
	"conferenceCall":            ConferenceCall,
	"contacts":                  Contacts,
	"copyleft":                  Copyleft,
	"copyright":                 Copyright,
	"crystalOscillator":         CrystalOscillator,
	"currencyExchange":          CurrencyExchange,
	"cursor":                    Cursor,
	"customerSupport":           CustomerSupport,
	"dam":                       Dam,
	"dataBackup":                DataBackup,
	"dataConfiguration":         DataConfiguration,
	"dataEncryption":            DataEncryption,
	"dataProtection":            DataProtection,
	"dataRecovery":              DataRecovery,
	"dataSheet":                 DataSheet,
	"database":                  Database,
	"debian":                    Debian,
	"debt":                      Debt,
	"decision":                  Decision,
	"deleteColumn":              DeleteColumn,
	"deleteDatabase":            DeleteDatabase,
	"deleteRow":                 DeleteRow,
	"department":                Department,
	"deployment":                Deployment,
	"diplomaOne":                DiplomaOne,
	"diplomaTwo":                DiplomaTwo,
	"disapprove":                Disapprove,
	"disclaimer":                Disclaimer,
	"dislike":                   Dislike,
	"display":                   Display,
	"doNotInhale":               DoNotInhale,
	"doNotInsert":               DoNotInsert,
	"doNotMix":                  DoNotMix,
	"document":                  Document,
	"donate":                    Donate,
	"doughnutChart":             DoughnutChart,
	"down":                      Down,
	"downLeft":                  DownLeft,
	"downRight":                 DownRight,
	"download":                  Download,
	"dribbble":                  Dribbble,
	"dvdLogo":                   DvdLogo,
	"editImage":                 EditImage,
	"electricalSensor":          ElectricalSensor,
	"electricalThreshold":       ElectricalThreshold,
	"electricity":               Electricity,
	"electroDevices":            ElectroDevices,
	"electronics":               Electronics,
	"emptyBattery":              EmptyBattery,
	"emptyFilter":               EmptyFilter,
	"emptyTrash":                EmptyTrash,
	"endCall":                   EndCall,
	"engineering":               Engineering,
	"enteringHeavenAlive":       EnteringHeavenAlive,
	"expand":                    Expand,
	"expired":                   Expired,
	"export":                    Export,
	"external":                  External,
	"factory":                   Factory,
	"factoryBreakdown":          FactoryBreakdown,
	"faq":                       Faq,
	"feedIn":                    FeedIn,
	"feedback":                  Feedback,
	"file":                      File,
	"filingCabinet":             FilingCabinet,
	"filledFilter":              FilledFilter,
	"film":                      Film,
	"filmReel":                  FilmReel,
	"finePrint":                 FinePrint,
	"flashAuto":                 FlashAuto,
	"flashOff":                  FlashOff,
	"flashOn":                   FlashOn,
	"flowChart":                 FlowChart,
	"folder":                    Folder,
	"frame":                     Frame,
	"fullBattery":               FullBattery,
	"fullTrash":                 FullTrash,
	"gallery":                   Gallery,
	"genealogy":                 Genealogy,
	"genericSortingAsc":         GenericSortingAsc,
	"genericSortingDesc":        GenericSortingDesc,
	"globe":                     Globe,
	"goodDecision":              GoodDecision,
	"google":                    Google,
	"graduationCap":             GraduationCap,
	"grid":                      Grid,
	"headset":                   Headset,
	"heatMap":                   HeatMap,
	"highBattery":               HighBattery,
	"highPriority":              HighPriority,
	"home":                      Home,
	"iconsEightCup":             IconsEightCup,
	"idea":                      Idea,
	"imageFile":                 ImageFile,
	"import":                    Import,
	"inTransit":                 InTransit,
	"info":                      Info,
	"inspection":                Inspection,
	"integratedWebcam":          IntegratedWebcam,
	"internal":                  Internal,
	"invite":                    Invite,
	"ipad":                      Ipad,
	"iphone":                    Iphone,
	"key":                       Key,
	"kindle":                    Kindle,
	"landscape":                 Landscape,
	"leave":                     Leave,
	"left":                      Left,
	"leftDown":                  LeftDown,
	"leftDownTwo":               LeftDownTwo,
	"leftUp":                    LeftUp,
	"leftUpTwo":                 LeftUpTwo,
	"library":                   Library,
	"lightAtTheEndOfTunnel":     LightAtTheEndOfTunnel,
	"like":                      Like,
	"likePlaceholder":           LikePlaceholder,
	"lineChart":                 LineChart,
	"link":                      Link,
	"linux":                     Linux,
	"list":                      List,
	"lock":                      Lock,
	"lockLandscape":             LockLandscape,
	"lockPortrait":              LockPortrait,
	"lowBattery":                LowBattery,
	"lowPriority":               LowPriority,
	"makeDecision":              MakeDecision,
	"manager":                   Manager,
	"mediumPriority":            MediumPriority,
	"menu":                      Menu,
	"middleBattery":             MiddleBattery,
	"mindMap":                   MindMap,
	"minus":                     Minus,
	"missedCall":                MissedCall,
	"mms":                       Mms,
	"moneyTransfer":             MoneyTransfer,
	"multipleCameras":           MultipleCameras,
	"multipleDevices":           MultipleDevices,
	"multipleInputs":            MultipleInputs,
	"multipleSmartphones":       MultipleSmartphones,
	"music":                     Music,
	"negativeDynamic":           NegativeDynamic,
	"neutralDecision":           NeutralDecision,
	"neutralTrading":            NeutralTrading,
	"news":                      News,
	"next":                      Next,
	"nfcSign":                   NfcSign,
	"nightLandscape":            NightLandscape,
	"nightPortrait":             NightPortrait,
	"noIdea":                    NoIdea,
	"noVideo":                   NoVideo,
	"nook":                      Nook,
	"numericalSortingTwelve":    NumericalSortingTwelve,
	"numericalSortingTwentyOne": NumericalSortingTwentyOne,
	"ok":                        Ok,
	"oldTimeCamera":             OldTimeCamera,
	"onlineSupport":             OnlineSupport,
	"openedFolder":              OpenedFolder,
	"orgUnit":                   OrgUnit,
	"organization":              Organization,
	"overtime":                  Overtime,
	"package":                   Package,
	"paid":                      Paid,
	"panorama":                  Panorama,
	"parallelTasks":             ParallelTasks,
	"phone":                     Phone,
	"phoneAndroid":              PhoneAndroid,
	"photoReel":                 PhotoReel,
	"picture":                   Picture,
	"pieChart":                  PieChart,
	"planner":                   Planner,
	"plus":                      Plus,
	"podiumWithAudience":        PodiumWithAudience,
	"podiumWithSpeaker":         PodiumWithSpeaker,
	"podiumWithoutSpeaker":      PodiumWithoutSpeaker,
	"portraitMode":              PortraitMode,
	"positiveDynamic":           PositiveDynamic,
	"previous":                  Previous,
	"print":                     Print,
	"privacy":                   Privacy,
	"process":                   Process,
	"puzzle":                    Puzzle,
	"questions":                 Questions,
	"radarPlot":                 RadarPlot,
	"rating":                    Rating,
	"ratings":                   Ratings,
	"reading":                   Reading,
	"readingEbook":              ReadingEbook,
	"reddit":                    Reddit,
	"redo":                      Redo,
	"refresh":                   Refresh,
	"registeredTrademark":       RegisteredTrademark,
	"removeImage":               RemoveImage,
	"reuse":                     Reuse,
	"right":                     Right,
	"rightDown":                 RightDown,
	"rightDownTwo":              RightDownTwo,
	"rightUp":                   RightUp,
	"rightUpTwo":                RightUpTwo,
	"rotateCamera":              RotateCamera,
	"rotateToLandscape":         RotateToLandscape,
	"rotateToPortrait":          RotateToPortrait,
	"ruler":                     Ruler,
	"rules":                     Rules,
	"safe":                      Safe,
	"salesPerformance":          SalesPerformance,
	"scatterPlot":               ScatterPlot,
	"search":                    Search,
	"selfServiceKiosk":          SelfServiceKiosk,
	"selfie":                    Selfie,
	"serialTasks":               SerialTasks,
	"serviceMark":               ServiceMark,
	"services":                  Services,
	"settings":                  Settings,
	"share":                     Share,
	"shipped":                   Shipped,
	"shop":                      Shop,
	"signature":                 Signature,
	"simCard":                   SimCard,
	"simCardChip":               SimCardChip,
	"slrBackSide":               SlrBackSide,
	"smartphoneTablet":          SmartphoneTablet,
	"sms":                       Sms,
	"soundRecordingCopyright":   SoundRecordingCopyright,
	"speaker":                   Speaker,
	"sportsMode":                SportsMode,
	"stackOfPhotos":             StackOfPhotos,
	"start":                     Start,
	"statistics":                Statistics,
	"steam":                     Steam,
	"stumbleupon":               Stumbleupon,
	"support":                   Support,
	"survey":                    Survey,
	"switchCamera":              SwitchCamera,
	"synchronize":               Synchronize,
	"tabletAndroid":             TabletAndroid,
	"template":                  Template,
	"timeline":                  Timeline,
	"todoList":                  TodoList,
	"touchscreenSmartphone":     TouchscreenSmartphone,
	"trademark":                 Trademark,
	"treeStructure":             TreeStructure,
	"twoSmartphones":            TwoSmartphones,
	"undo":                      Undo,
	"unlock":                    Unlock,
	"up":                        Up,
	"upLeft":                    UpLeft,
	"upRight":                   UpRight,
	"upload":                    Upload,
	"usb":                       Usb,
	"videoCall":                 VideoCall,
	"videoFile":                 VideoFile,
	"videoProjector":            VideoProjector,
	"viewDetails":               ViewDetails,
	"vip":                       Vip,
	"vlc":                       Vlc,
	"voicePresentation":         VoicePresentation,
	"voicemail":                 Voicemail,
	"webcam":                    Webcam,
	"wiFiLogo":                  WiFiLogo,
	"wikipedia":                 Wikipedia,
	"workflow":                  Workflow,
}

func About(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#2196F3" d="M37 40H11l-6 6V12c0-3.3 2.7-6 6-6h26c3.3 0 6 2.7 6 6v22c0 3.3-2.7 6-6 6"/><g fill="#fff"><path d="M22 20h4v11h-4z"/><circle cx="24" cy="15" r="2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AcceptDatabase(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#D1C4E9" d="M38 7H10c-1.1 0-2 .9-2 2v6c0 1.1.9 2 2 2h28c1.1 0 2-.9 2-2V9c0-1.1-.9-2-2-2m0 12H10c-1.1 0-2 .9-2 2v6c0 1.1.9 2 2 2h28c1.1 0 2-.9 2-2v-6c0-1.1-.9-2-2-2m0 12H10c-1.1 0-2 .9-2 2v6c0 1.1.9 2 2 2h28c1.1 0 2-.9 2-2v-6c0-1.1-.9-2-2-2"/><circle cx="38" cy="38" r="10" fill="#43A047"/><path fill="#DCEDC8" d="M42.5 33.3L36.8 39l-2.7-2.7l-2.1 2.2l4.8 4.8l7.8-7.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AddColumn(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#90CAF9" d="M30 5H18c-2.2 0-4 1.8-4 4v30c0 2.2 1.8 4 4 4h12c2.2 0 4-1.8 4-4V9c0-2.2-1.8-4-4-4M18 39V9h12v30z"/><circle cx="38" cy="38" r="10" fill="#43A047"/><g fill="#fff"><path d="M36 32h4v12h-4z"/><path d="M32 36h12v4H32z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AddDatabase(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#D1C4E9" d="M38 7H10c-1.1 0-2 .9-2 2v6c0 1.1.9 2 2 2h28c1.1 0 2-.9 2-2V9c0-1.1-.9-2-2-2m0 12H10c-1.1 0-2 .9-2 2v6c0 1.1.9 2 2 2h28c1.1 0 2-.9 2-2v-6c0-1.1-.9-2-2-2m0 12H10c-1.1 0-2 .9-2 2v6c0 1.1.9 2 2 2h28c1.1 0 2-.9 2-2v-6c0-1.1-.9-2-2-2"/><circle cx="38" cy="38" r="10" fill="#43A047"/><g fill="#fff"><path d="M36 32h4v12h-4z"/><path d="M32 36h12v4H32z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AddImage(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#8CBCD6" d="M40 41H8c-2.2 0-4-1.8-4-4V11c0-2.2 1.8-4 4-4h32c2.2 0 4 1.8 4 4v26c0 2.2-1.8 4-4 4"/><circle cx="35" cy="16" r="3" fill="#B3DDF5"/><path fill="#9AC9E3" d="M20 16L9 32h22z"/><path fill="#B3DDF5" d="m31 22l-8 10h16z"/><circle cx="38" cy="38" r="10" fill="#43A047"/><g fill="#fff"><path d="M36 32h4v12h-4z"/><path d="M32 36h12v4H32z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AddRow(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#90CAF9" d="M43 30V18c0-2.2-1.8-4-4-4H9c-2.2 0-4 1.8-4 4v12c0 2.2 1.8 4 4 4h30c2.2 0 4-1.8 4-4M9 18h30v12H9z"/><circle cx="38" cy="38" r="10" fill="#43A047"/><g fill="#fff"><path d="M32 36h12v4H32z"/><path d="M36 32h4v12h-4z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AddressBook(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#673AB7" d="M38 44H12V4h26c2.2 0 4 1.8 4 4v32c0 2.2-1.8 4-4 4"/><path fill="#311B92" d="M10 4h2v40h-2c-2.2 0-4-1.8-4-4V8c0-2.2 1.8-4 4-4"/><path fill="#fff" d="M36 24.2c-.1 4.8-3.1 6.9-5.3 6.7c-.6-.1-2.1-.1-2.9-1.6c-.8 1-1.8 1.6-3.1 1.6c-2.6 0-3.3-2.5-3.4-3.1c-.1-.7-.2-1.4-.1-2.2c.1-1 1.1-6.5 5.7-6.5c2.2 0 3.5 1.1 3.7 1.3l-.6 6.8c0 .3-.2 1.6 1.1 1.6c2.1 0 2.4-3.9 2.4-4.6c.1-1.2.3-8.2-7-8.2c-6.9 0-7.9 7.4-8 9.2c-.5 8.5 6 8.5 7.2 8.5c1.7 0 3.7-.7 3.9-.8l.4 2c-.3.2-2 1.1-4.4 1.1c-2.2 0-10.1-.4-9.8-10.8c.3-2.1 1.6-11.2 10.8-11.2c9.2 0 9.4 8.1 9.4 10.2m-11.9 1.3c-.1 1 0 1.8.2 2.3c.2.5.6.8 1.2.8c.1 0 .3 0 .4-.1c.2-.1.3-.1.5-.3c.2-.1.3-.3.5-.6c.2-.2.3-.6.4-1l.5-5.4c-.2-.1-.5-.1-.7-.1c-.5 0-.9.1-1.2.3c-.3.2-.6.5-.9.8c-.2.4-.4.8-.6 1.3s-.2 1.3-.3 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Advance(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#1565C0"><path d="M46.1 24L33 35V13zM10 20h4v8h-4zm-6 0h4v8H4zm12 0h4v8h-4z"/><path d="M22 20h14v8H22z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Advertising(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#90CAF9" d="M17.4 33H15v-4h4l.4 1.5c.3 1.3-.7 2.5-2 2.5M37 36s-11.8-7-18-7V15c5.8 0 18-7 18-7z"/><g fill="#283593"><circle cx="9" cy="22" r="5"/><path d="M40 19h-3v6h3c1.7 0 3-1.3 3-3s-1.3-3-3-3M18.6 41.2c-.9.6-2.5 1.2-4.6 1.4c-.6.1-1.2-.3-1.4-1L8.2 27.9S17 21.7 17 29c0 5.5 1.5 8.4 2.2 9.5c.5.7.5 1.6 0 2.3c-.2.2-.4.3-.6.4"/></g><path fill="#3F51B5" d="M9 29h10V15H9c-1.1 0-2 .9-2 2v10c0 1.1.9 2 2 2"/><path fill="#42A5F5" d="M38 38c-1.1 0-2-.9-2-2V8c0-1.1.9-2 2-2s2 .9 2 2v28c0 1.1-.9 2-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlarmClock(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#37474F" d="m38.5 44.6l-4-4l2.1-2.1l4 4c.6.6.6 1.5 0 2.1c-.5.5-1.5.5-2.1 0m-29 0l4-4l-2.1-2.1l-4 4c-.6.6-.6 1.5 0 2.1c.5.5 1.5.5 2.1 0"/><circle cx="24" cy="24" r="20" fill="#C62828"/><circle cx="24" cy="24" r="16" fill="#eee"/><path fill="#E53935" d="m15.096 33.48l-.566-.566l9.191-9.191l.566.565z"/><path d="M23 11h2v13h-2z"/><path d="M31.285 29.654L29.66 31.28l-6.504-6.504l1.626-1.627z"/><circle cx="24" cy="24" r="2"/><circle cx="24" cy="24" r="1" fill="#C62828"/><path fill="#37474F" d="M22 1h4v3h-4zm22.4 15.2c2.5-3.5 2.1-8.4-1-11.5c-3.1-3.1-8-3.5-11.5-1zm-40.8 0c-2.5-3.5-2.1-8.4 1-11.5c3.1-3.1 8-3.5 11.5-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphabeticalSortingAz(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#546E7A" d="M38 33V5h-4v28h-6l8 10l8-10z"/><path fill="#2196F3" d="M16.8 17.2h-5.3l-1.1 3H6.9L12.6 5h2.9l5.7 15.2H18zm-4.6-2.7H16l-1.9-5.7zm.2 26H20V43H8.4v-1.9L16 30.3H8.4v-2.5h11.4v1.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlphabeticalSortingZa(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#2196F3" d="M16.8 40h-5.3l-1.1 3H6.9l5.7-15.2h2.9L21.1 43h-3.2zm-4.6-2.7H16l-1.9-5.7zm.2-19.6H20v2.5H8.4v-1.9L16 7.5H8.4V5h11.4v1.7z"/><path fill="#546E7A" d="M38 33V5h-4v28h-6l8 10l8-10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AndroidOs(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#7CB342" d="M12 29.001a2 2 0 0 1-4 0v-9a2 2 0 0 1 4 0zm28 0a2 2 0 0 1-4 0v-9a2 2 0 0 1 4 0zM22 40a2 2 0 0 1-4 0v-9a2 2 0 0 1 4 0zm8 0a2 2 0 0 1-4 0v-9a2 2 0 0 1 4 0z"/><path fill="#7CB342" d="M14 18.001V33a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V18.001zM24 8c-6 0-9.655 3.645-10 8h20c-.346-4.355-4-8-10-8m-4 5.598a1 1 0 1 1 0-2a1 1 0 0 1 0 2m8 0a1 1 0 1 1 0-2a1 1 0 0 1 0 2"/><path fill="none" stroke="#7CB342" stroke-linecap="round" stroke-width="2" d="m30 7l-1.666 2.499M18 7l1.333 2.082"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Answers(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#42A5F5" d="M36 44H8V8h20l8 8z"/><path fill="#90CAF9" d="M40 40H12V4h20l8 8z"/><path fill="#E1F5FE" d="M38.5 13H31V5.5z"/><path fill="#1976D2" d="M23.4 29.9c0-.2 0-.4.1-.6s.2-.3.3-.5s.3-.2.5-.3s.4-.1.6-.1s.5 0 .7.1s.4.2.5.3s.2.3.3.5s.1.4.1.6s0 .4-.1.6s-.2.3-.3.5s-.3.2-.5.3s-.4.1-.7.1s-.5 0-.6-.1s-.4-.2-.5-.3s-.2-.3-.3-.5s-.1-.4-.1-.6m2.7-3.1h-2.3l-.4-9.8h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Approval(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#8BC34A" d="m24 3l4.7 3.6l5.8-.8l2.2 5.5l5.5 2.2l-.8 5.8L45 24l-3.6 4.7l.8 5.8l-5.5 2.2l-2.2 5.5l-5.8-.8L24 45l-4.7-3.6l-5.8.8l-2.2-5.5l-5.5-2.2l.8-5.8L3 24l3.6-4.7l-.8-5.8l5.5-2.2l2.2-5.5l5.8.8z"/><path fill="#CCFF90" d="M34.6 14.6L21 28.2l-5.6-5.6l-2.8 2.8l8.4 8.4l16.4-16.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Approve(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#FFCC80"><circle cx="38" cy="26" r="4"/><circle cx="10" cy="26" r="4"/><path d="M39 19c0-12.7-30-8.3-30 0v10c0 8.3 6.7 15 15 15s15-6.7 15-15z"/><path d="M24 4C15.2 4 8 11.2 8 20v3.5l2.1.6V19l19.5-6.3l8.2 6.3v5.1l2.1-.6V20C40 12.5 34.6 4 24 4"/></g><path fill="#4CAF50" d="M32.6 18.6L22.3 28.9L17.4 24l-2.8 2.8l7.7 7.7l13.1-13.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AreaChart(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#3F51B5" d="M42 37H6V25l10-15l14 7L42 6z"/><path fill="#00BCD4" d="M42 42H6V32l10-8l14 2l12-9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Assistant(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#FFA726"><circle cx="10" cy="26" r="4"/><circle cx="38" cy="26" r="4"/></g><path fill="#FFB74D" d="M39 19c0-12.7-30-8.3-30 0v10c0 8.3 6.7 15 15 15s15-6.7 15-15z"/><path fill="#FF5722" d="M24 3C14.6 3 7 10.6 7 20v3.4L9 25v-3l21-9.8l9 9.8v3l2-1.6V20c0-8-5.7-17-17-17"/><g fill="#784719"><circle cx="31" cy="26" r="2"/><circle cx="17" cy="26" r="2"/></g><path fill="#757575" d="M43 24c-.6 0-1 .4-1 1v-7c0-8.8-7.2-16-16-16h-7c-.6 0-1 .4-1 1s.4 1 1 1h7c7.7 0 14 6.3 14 14v10c0 .6.4 1 1 1s1-.4 1-1v2c0 3.9-3.1 7-7 7H24c-.6 0-1 .4-1 1s.4 1 1 1h11c5 0 9-4 9-9v-5c0-.6-.4-1-1-1"/><g fill="#37474F"><path d="M43 22h-1c-1.1 0-2 .9-2 2v4c0 1.1.9 2 2 2h1c1.1 0 2-.9 2-2v-4c0-1.1-.9-2-2-2"/><circle cx="24" cy="38" r="2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AudioFile(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" d="M204 0h48v48h-48z"/><path fill="#90CAF9" d="M244 45h-32V3h22l10 10z"/><path fill="#E1F5FE" d="M242.5 14H233V4.5z"/><g fill="#1976D2"><circle cx="227" cy="30" r="4"/><path d="m234 21l-5-2v11h2v-7.1l3 1.1z"/></g><path fill="#90CAF9" d="M40 45H8V3h22l10 10z"/><path fill="#E1F5FE" d="M38.5 14H29V4.5z"/><g fill="#1976D2"><circle cx="23" cy="30" r="4"/><path d="m30 21l-5-2v11h2v-7.1l3 1.1z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Automatic(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#3F51B5" d="M39 43H9c-2.2 0-4-1.8-4-4V9c0-2.2 1.8-4 4-4h30c2.2 0 4 1.8 4 4v30c0 2.2-1.8 4-4 4"/><path fill="#B3E5FC" d="M33.6 25.4c.1-.4.1-.9.1-1.4s0-.9-.1-1.4l2.8-2c.3-.2.4-.6.2-.9l-2.7-4.6c-.2-.3-.5-.4-.8-.3L30 16.3c-.7-.6-1.5-1-2.4-1.4l-.3-3.4c0-.3-.3-.6-.6-.6h-5.3c-.3 0-.6.3-.6.6l-.4 3.5c-.9.3-1.6.8-2.4 1.4L14.9 15c-.3-.1-.7 0-.8.3l-2.7 4.6c-.2.3-.1.7.2.9l2.8 2c-.1.4-.1.9-.1 1.4s0 .9.1 1.4l-2.8 2c-.3.2-.4.6-.2.9l2.7 4.6c.2.3.5.4.8.3L18 32c.7.6 1.5 1 2.4 1.4l.3 3.4c0 .3.3.6.6.6h5.3c.3 0 .6-.3.6-.6l.3-3.4c.9-.3 1.6-.8 2.4-1.4l3.1 1.4c.3.1.7 0 .8-.3l2.7-4.6c.2-.3.1-.7-.2-.9zM24 29c-2.8 0-5-2.2-5-5s2.2-5 5-5s5 2.2 5 5s-2.2 5-5 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Automotive(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="#F44336" stroke-miterlimit="10" stroke-width="4" d="M7 20v-8c0-2.2 1.8-4 4-4h14c1.2 0 2.4.6 3.2 1.6L35 18"/><g fill="#37474F"><circle cx="35" cy="37" r="5"/><circle cx="13" cy="37" r="5"/></g><path fill="#F44336" d="M40.2 17L33 14H7c-1.2 0-2 .8-2 2v10c0 1.2.8 2 2 2h1c0-2.8 2.2-5 5-5s5 2.2 5 5h12c0-2.8 2.2-5 5-5s5 2.2 5 5h1c1.2 0 2-.8 2-2v-5.2c0-1.6-1.2-3.2-2.8-3.8"/><g fill="#546E7A"><circle cx="24" cy="37" r="3"/><circle cx="35" cy="37" r="2"/><circle cx="13" cy="37" r="2"/><path d="M30.4 39c-.3-.6-.4-1.3-.4-2s.2-1.4.4-2H17.6c.3.6.4 1.3.4 2s-.2 1.4-.4 2z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BadDecision(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#FFCC80"><circle cx="38" cy="26" r="4"/><circle cx="10" cy="26" r="4"/><path d="M39 19c0-12.7-30-8.3-30 0v10c0 8.3 6.7 15 15 15s15-6.7 15-15z"/><path d="M24 4C15.2 4 8 11.2 8 20v3.5l2.1.6V19l19.5-6.3l8.2 6.3v5.1l2.1-.6V20C40 12.5 34.6 4 24 4"/></g><path fill="#F44336" d="M16 24h16v4H16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarChart(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#00BCD4" d="M19 22h10v20H19zM6 12h10v30H6zm26-6h10v36H32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bbc(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#37474F" d="M0 17v14h14V17zm34 0v14h14V17zm-17 0v14h14V17z"/><path fill="#FFF" d="M8.2 23.7s1.3-.5 1.3-2c0 0 .2-2.4-3-2.7H3v10h4s3.4 0 3.4-2.8c0 0 .1-1.9-2.2-2.5m-3.4-3.1h1.4c1.5.1 1.4 1.2 1.4 1.2c0 1.4-1.6 1.4-1.6 1.4H4.8zm1.9 6.9H4.8v-2.7h1.9c1.9 0 1.9 1.3 1.9 1.3c-.1 1.4-1.9 1.4-1.9 1.4m18.5-3.8s1.3-.5 1.3-2c0 0 .2-2.4-3-2.7H20v10h4s3.4 0 3.4-2.8c0 0 .1-1.9-2.2-2.5m-3.4-3.1h1.4c1.5.1 1.4 1.2 1.4 1.2c0 1.4-1.6 1.4-1.6 1.4h-1.2zm1.9 6.9h-1.9v-2.7h1.9c1.9 0 1.9 1.3 1.9 1.3c-.1 1.4-1.9 1.4-1.9 1.4m21.6.6s-2.9 1.8-6.3.4c0 0-2.9-1.1-3-4.5c0 0-.1-3.5 3.7-4.7c0 0 1-.4 2.8-.2c0 0 1.1.1 2.7.8v1.8s-1.7-1.1-3.7-1.1c0 0-3.6-.1-3.8 3.5c0 0-.1 3.3 3.7 3.4c0 0 1.6.2 3.8-1.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bearish(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#F44336" d="M40 34h4v10h-4zm-6-5h4v15h-4zm-6 4h4v11h-4zm-6-8h4v19h-4zm-6 3h4v16h-4zm-6-4h4v20h-4zm-6-5h4v25H4z"/><g fill="#D32F2F"><path d="m34 13.2l-4 4l-10-10l-5 5l-7.6-7.6l-2.8 2.8L15 17.8l5-5l10 10l4-4l6.1 6.1l2.8-2.8z"/><path d="M44 26h-9l9-9z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Binoculars(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#37474F"><circle cx="33" cy="16" r="6"/><circle cx="15" cy="16" r="6"/><path d="m46.7 25l-15.3 3H16.7L1.4 25l4.3-7.9C6.8 15.2 8.8 14 11 14h26.2c2.2 0 4.2 1.2 5.3 3.1z"/><circle cx="38" cy="30" r="10"/><circle cx="10" cy="30" r="10"/><circle cx="24" cy="28" r="5"/></g><circle cx="24" cy="28" r="2" fill="#546E7A"/><g fill="#a0f"><circle cx="38" cy="30" r="7"/><circle cx="10" cy="30" r="7"/></g><path fill="#CE93D8" d="M41.7 27.7c-1-1.1-2.3-1.7-3.7-1.7s-2.8.6-3.7 1.7c-.4.4-.3 1 .1 1.4c.4.4 1 .3 1.4-.1c1.2-1.3 3.3-1.3 4.5 0c.2.2.5.3.7.3c.2 0 .5-.1.7-.3c.4-.3.4-.9 0-1.3M10 26c-1.4 0-2.8.6-3.7 1.7c-.4.4-.3 1 .1 1.4c.4.4 1 .3 1.4-.1c1.2-1.3 3.3-1.3 4.5 0c.2.2.5.3.7.3c.2 0 .5-.1.7-.3c.4-.4.4-1 .1-1.4c-1-1-2.4-1.6-3.8-1.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Biohazard(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#00A344" d="M24 13c-7.2 0-13 5.8-13 13s5.8 13 13 13s13-5.8 13-13s-5.8-13-13-13m0 22c-5 0-9-4-9-9s4-9 9-9s9 4 9 9s-4 9-9 9"/><path fill="#00C853" d="M8.5 25.4c4-2.2 9-1.1 11.5 2.5c.1.1.2.1.3.1l1.2-.7c.1-.1.2-.2.1-.3c0-.2-.1-.4-.1-.6v-.6c0-.1 0-.1.1-.2c0-.1 0-.1.1-.2c0 0 0-.1.1-.1c0 0 0-.1.1-.1c0 0 0-.1.1-.1l.1-.1l.1-.1l.1-.1l.1-.1l.1-.1h.1c.2-.1.4-.2.5-.2c.1 0 .2-.1.2-.3v-1.3c0-.1-.1-.2-.2-.2c-4.5-.4-8-4.1-8-8.6c0-4.1 3-7.6 6.9-8.4c.1 0 .2-.1.2-.3v-.5c0-.1-.1-.2-.2-.2c-5.6.9-10 5.8-10 11.7c0 1.3.2 2.6.6 3.8c-1.2.2-2.5.7-3.6 1.3c-5.2 3-7.3 9.2-5.2 14.5c.1.1.2.1.3.1l.3-.2c.1-.1.2-.2.1-.3c-1.2-3.8.3-8.1 4-10.1m30.5-4c-1.2-.7-2.4-1.1-3.6-1.3c.4-1.2.6-2.4.6-3.8c0-5.9-4.4-10.8-10.2-11.7c-.1 0-.2.1-.2.2v.4c0 .1.1.2.2.3c4 .8 6.9 4.3 6.9 8.4c0 4.5-3.5 8.2-8 8.6c-.1 0-.2.1-.2.2V24c0 .1.1.2.2.3c.2.1.4.1.6.2l.1.1c.1 0 .1.1.1.1l.3.3l.1.1l.1.1v.2s0 .1.1.1c0 .1 0 .1.1.2v.7c0 .2 0 .4-.1.6c0 .1 0 .2.1.3l1.2.7c.1.1.2 0 .3-.1c2.6-3.6 7.6-4.8 11.5-2.5c3.6 2.1 5.2 6.3 3.9 10.1c0 .1 0 .2.1.3l.3.2c.1.1.2 0 .3-.1c2.5-5.4.4-11.6-4.8-14.5m-8.2 18.9c-4-2.2-5.5-7.1-3.5-11.1c.1-.1 0-.2-.1-.3l-1.2-.7c-.1-.1-.2 0-.3 0c-.2.1-.3.3-.5.3c-.1 0-.1.1-.2.1s-.1 0-.2.1c-.1 0-.3.1-.4.1h-1.2c-.1 0-.1 0-.2-.1c0 0-.1 0-.1-.1h-.1c-.2-.1-.3-.2-.5-.3c-.1-.1-.2-.1-.3 0l-1.2.7c-.1.1-.1.2-.1.3c1.9 4 .4 8.8-3.5 11.1c-3.6 2.1-8.2 1.3-10.9-1.7c-.1-.1-.2-.1-.3-.1l-.3.2c-.1.1-.1.2-.1.3c3.6 4.5 10.2 5.8 15.4 2.8c1.2-.7 2.2-1.5 3-2.4c.8.9 1.8 1.8 3 2.4c5.2 3 11.7 1.6 15.4-2.8c.1-.1 0-.2-.1-.3l-.3-.3c-.1-.1-.2 0-.3.1c-2.7 2.9-7.3 3.7-10.9 1.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Biomass(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#9CCC65" d="M32 15V7H16v8L6.2 40c-.6 1.5.5 3 2.1 3h31.5c1.6 0 2.6-1.6 2.1-3z"/><path fill="#8BC34A" d="M32 9H16c-1.1 0-2-.9-2-2s.9-2 2-2h16c1.1 0 2 .9 2 2s-.9 2-2 2"/><path fill="#2E7D32" d="M28 30c0 4.4-4 8-4 8s-4-3.6-4-8s4-8 4-8s4 3.6 4 8"/><path fill="#388E3C" d="M31.1 32.6c-2 4-7.1 5.4-7.1 5.4s-2-5 0-8.9s7.1-5.4 7.1-5.4s2 4.9 0 8.9"/><path fill="#43A047" d="M16.9 32.6c2 4 7.1 5.4 7.1 5.4s2-5 0-8.9s-7.1-5.4-7.1-5.4s-2 4.9 0 8.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Biotech(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#AD1457" d="M36 4c0 9.3-6 13.2-12.8 17.8C16.1 26.5 8 31.8 8 44h4c0-10.1 6.5-14.4 13.4-18.9C32.2 20.6 40 15.4 40 4z"/><path fill="#AD1457" d="M38 41H11c-.6 0-1-.4-1-1s.4-1 1-1h27c.6 0 1 .4 1 1s-.4 1-1 1m-2-4H12c-.6 0-1-.4-1-1s.4-1 1-1h24c.6 0 1 .4 1 1s-.4 1-1 1m-2-4H14c-.6 0-1-.4-1-1s.4-1 1-1h20c.6 0 1 .4 1 1s-.4 1-1 1m-5-4H19c-.6 0-1-.4-1-1s.4-1 1-1h10c.6 0 1 .4 1 1s-.4 1-1 1"/><path fill="#E91E63" d="M37 9H10c-.6 0-1-.4-1-1s.4-1 1-1h27c.6 0 1 .4 1 1s-.4 1-1 1m-1 4H12c-.6 0-1-.4-1-1s.4-1 1-1h24c.6 0 1 .4 1 1s-.4 1-1 1m-2 4H14c-.6 0-1-.4-1-1s.4-1 1-1h20c.6 0 1 .4 1 1s-.4 1-1 1m-5 4H19c-.6 0-1-.4-1-1s.4-1 1-1h10c.6 0 1 .4 1 1s-.4 1-1 1"/><path fill="#E91E63" d="M40 44h-4c0-10.1-6.5-14.4-13.4-18.9C15.8 20.6 8 15.4 8 4h4c0 9.3 6 13.2 12.8 17.8C31.9 26.5 40 31.8 40 44"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookmark(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#F44336" d="m37 43l-13-6l-13 6V9c0-2.2 1.8-4 4-4h18c2.2 0 4 1.8 4 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Briefcase(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#424242" d="M27 7h-6c-1.7 0-3 1.3-3 3v3h2v-3c0-.6.4-1 1-1h6c.6 0 1 .4 1 1v3h2v-3c0-1.7-1.3-3-3-3"/><path fill="#E65100" d="M40 43H8c-2.2 0-4-1.8-4-4V15c0-2.2 1.8-4 4-4h32c2.2 0 4 1.8 4 4v24c0 2.2-1.8 4-4 4"/><path fill="#FF6E40" d="M40 28H8c-2.2 0-4-1.8-4-4v-9c0-2.2 1.8-4 4-4h32c2.2 0 4 1.8 4 4v9c0 2.2-1.8 4-4 4"/><path fill="#FFF3E0" d="M26 26h-4c-.6 0-1-.4-1-1v-2c0-.6.4-1 1-1h4c.6 0 1 .4 1 1v2c0 .6-.4 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrokenLink(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#1976D2" d="M17.5 27c-1.1 1.2-2.7 2-4.5 2h-3c-3.3 0-6-2.7-6-6s2.7-6 6-6h3c1.8 0 3.4.8 4.5 2h4.7c-1.5-3.5-5.1-6-9.2-6h-3C4.5 13 0 17.5 0 23s4.5 10 10 10h3c4.1 0 7.6-2.5 9.2-6zM38 13h-3c-4.1 0-7.6 2.5-9.2 6h4.7c1.1-1.2 2.7-2 4.5-2h3c3.3 0 6 2.7 6 6s-2.7 6-6 6h-3c-1.8 0-3.4-.8-4.5-2h-4.7c1.5 3.5 5.1 6 9.2 6h3c5.5 0 10-4.5 10-10s-4.5-10-10-10"/><path fill="#00BCD4" d="M19.5 4L16 6l6.1 8.1l1.3-.8zm9 0L32 6l-6.1 8.1l-1.3-.8zm0 40l3.5-2l-6.1-8.1l-1.3.8zm-9 0L16 42l6.1-8.1l1.3.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bullish(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#4CAF50" d="M40 21h4v23h-4zm-6 7h4v16h-4zm-6-5h4v21h-4zm-6 6h4v15h-4zm-6 3h4v12h-4zm-6-2h4v14h-4zm-6 4h4v10H4z"/><g fill="#388E3C"><path d="M40.1 9.1L34 15.2l-4-4l-10 10l-5-5L4.6 26.6l2.8 2.8l7.6-7.6l5 5l10-10l4 4l8.9-8.9z"/><path d="M44 8h-9l9 9z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Business(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#263238" d="M11 44H9c-.6 0-1-.4-1-1v-2h4v2c0 .6-.4 1-1 1m28 0h-2c-.6 0-1-.4-1-1v-2h4v2c0 .6-.4 1-1 1"/><path fill="#37474F" d="M27 7h-6c-1.7 0-3 1.3-3 3v3h2v-3c0-.6.4-1 1-1h6c.6 0 1 .4 1 1v3h2v-3c0-1.7-1.3-3-3-3"/><path fill="#78909C" d="M40 43H8c-2.2 0-4-1.8-4-4V15c0-2.2 1.8-4 4-4h32c2.2 0 4 1.8 4 4v24c0 2.2-1.8 4-4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BusinessContact(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#673AB7" d="M40 7H8c-2.2 0-4 1.8-4 4v26c0 2.2 1.8 4 4 4h5v-1.3c-.6-.3-1-1-1-1.7c0-1.1.9-2 2-2s2 .9 2 2c0 .7-.4 1.4-1 1.7V41h18v-1.3c-.6-.3-1-1-1-1.7c0-1.1.9-2 2-2s2 .9 2 2c0 .7-.4 1.4-1 1.7V41h5c2.2 0 4-1.8 4-4V11c0-2.2-1.8-4-4-4"/><g fill="#D1C4E9"><circle cx="24" cy="18" r="4"/><path d="M31 28s-1.9-4-7-4s-7 4-7 4v2h14z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Businessman(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#FF9800" d="m24 37l-5-6v-6h10v6z"/><g fill="#FFA726"><circle cx="33" cy="19" r="2"/><circle cx="15" cy="19" r="2"/></g><path fill="#FFB74D" d="M33 13c0-7.6-18-5-18 0v7c0 5 4 9 9 9s9-4 9-9z"/><path fill="#424242" d="M24 4c-6.1 0-10 4.9-10 11v2.3l2 1.7v-5l12-4l4 4v5l2-1.7V15c0-4-1-8-6-9l-1-2z"/><g fill="#784719"><circle cx="28" cy="19" r="1"/><circle cx="20" cy="19" r="1"/></g><path fill="#fff" d="m24 43l-5-12l5 1l5-1z"/><path fill="#D32F2F" d="m23 35l-.7 4.5l1.7 4l1.7-4L25 35l1-1l-2-2l-2 2z"/><path fill="#546E7A" d="m29 31l-5 12l-5-12S8 33 8 44h32c0-11-11-13-11-13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Businesswoman(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#BF360C" d="M16 15h16v18H16z"/><path fill="#78909C" d="M40 44H8c0-11 11-13 11-13h10s11 2 11 13"/><path fill="#FF9800" d="M24 37c-2.2 0-5-6-5-6v-6h10v6s-2.8 6-5 6"/><path fill="#FFB74D" d="M33 14c0-7.6-18-5-18 0v7c0 5 4 9 9 9s9-4 9-9z"/><path fill="#FF5722" d="M24 4C17.9 4 9 7.4 9 27.3l7 4.7V19l12-7l4 5v15l7-6c0-4-.7-20-11-20l-1-2z"/><path fill="#FFB74D" d="M24 38c-4.4 0-5-7-5-7s2.5 4 5 4s5-4 5-4s-.6 7-5 7"/><circle cx="28" cy="21" r="1" fill="#784719"/><circle cx="20" cy="21" r="1" fill="#784719"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ButtingIn(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#CFD8DC" d="M24 3C12.4 3 3 12.4 3 24s9.4 21 21 21c.3 0 .7 0 1-.1V3.1c-.3-.1-.7-.1-1-.1"/><path fill="#37474F" d="M25 3.1V45c4.1-.2 7.9-1.5 11-3.7V6.8c-3.1-2.2-6.9-3.6-11-3.7"/><path fill="#FFB74D" d="M20.5 13c-6.4.3-11.6 5.7-11.5 12.1c0 2.8 1 5.4 2.7 7.5c1.4 1.7 2.3 3.9 2.3 6.1v3.8c3 1.6 6.4 2.5 10 2.5c.3 0 .7 0 1-.1c.7 0 1.3-.1 2-.2v-9.4c3.6-2.1 6-5.9 6-10.4c0-6.7-5.6-12.2-12.5-11.9"/><path fill="#FFB74D" d="m29 38.6l-4-.6v-9h8l-.7 7c-.2 1.6-1.6 2.8-3.3 2.6"/><path fill="#FFB74D" d="m39 29l-7 2l-1-5l1-4z"/><circle cx="29.5" cy="25.5" r="1.5" fill="#784719"/><path fill="#FF5722" d="M21 12c-7.2 0-13 5.8-13 13c0 7.6 5.1 9 6 13l4-3v-8l5-2l1-4c3.2 0 6-3.9 6-6.1c-2.1-1.9-5.6-2.9-9-2.9"/><circle cx="19" cy="27" r="3" fill="#FFB74D"/><path fill="#CFD8DC" d="M45 24c0-7.1-3.6-13.4-9-17.2v34.4c5.4-3.8 9-10.1 9-17.2"/><path fill="#FF9800" d="M20 44.6c1.3.2 2.6.4 4 .4c.3 0 .7 0 1-.1c.7 0 1.3-.1 2-.2v-6.5l-7-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CableRelease(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#37474F" d="M34.9 29.1c-2.7-2.7-7.1-2.7-9.8 0l-4 4c-1.7 1.7-4.5 1.7-6.2 0c-1.7-1.7-1.7-4.5 0-6.2l4.5-4.5l-2.8-2.8l-4.5 4.5c-3.3 3.3-3.3 8.6 0 11.8c3.3 3.3 8.6 3.3 11.8 0l4-4c1.2-1.1 3-1.1 4.2 0c1.1 1.2 1.1 3 0 4.2L27 41.2l2.8 2.8l5.1-5.1c2.7-2.7 2.7-7.1 0-9.8"/><path fill="#0277BD" d="M16.1 22.9c-2.8-2.8-2.8-7.3 0-10l6.8-6.8c2.8-2.8 7.3-2.8 10 0c2.8 2.8 2.8 7.3 0 10l-6.8 6.8c-2.8 2.8-7.2 2.8-10 0"/><circle cx="28" cy="11" r="4" fill="#B3E5FC"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calculator(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#616161" d="M40 16H8v24c0 2.2 1.8 4 4 4h24c2.2 0 4-1.8 4-4z"/><path fill="#424242" d="M36 4H12C9.8 4 8 5.8 8 8v9h32V8c0-2.2-1.8-4-4-4"/><path fill="#9CCC65" d="M36 14H12c-.6 0-1-.4-1-1V8c0-.6.4-1 1-1h24c.6 0 1 .4 1 1v5c0 .6-.4 1-1 1"/><path fill="#33691E" d="M33 10h2v2h-2zm-4 0h2v2h-2z"/><path fill="#FF5252" d="M36 23h-3c-.6 0-1-.4-1-1v-2c0-.6.4-1 1-1h3c.6 0 1 .4 1 1v2c0 .6-.4 1-1 1"/><path fill="#E0E0E0" d="M15 23h-3c-.6 0-1-.4-1-1v-2c0-.6.4-1 1-1h3c.6 0 1 .4 1 1v2c0 .6-.4 1-1 1m7 0h-3c-.6 0-1-.4-1-1v-2c0-.6.4-1 1-1h3c.6 0 1 .4 1 1v2c0 .6-.4 1-1 1m7 0h-3c-.6 0-1-.4-1-1v-2c0-.6.4-1 1-1h3c.6 0 1 .4 1 1v2c0 .6-.4 1-1 1m-14 6h-3c-.6 0-1-.4-1-1v-2c0-.6.4-1 1-1h3c.6 0 1 .4 1 1v2c0 .6-.4 1-1 1m7 0h-3c-.6 0-1-.4-1-1v-2c0-.6.4-1 1-1h3c.6 0 1 .4 1 1v2c0 .6-.4 1-1 1m7 0h-3c-.6 0-1-.4-1-1v-2c0-.6.4-1 1-1h3c.6 0 1 .4 1 1v2c0 .6-.4 1-1 1m-14 6h-3c-.6 0-1-.4-1-1v-2c0-.6.4-1 1-1h3c.6 0 1 .4 1 1v2c0 .6-.4 1-1 1m7 0h-3c-.6 0-1-.4-1-1v-2c0-.6.4-1 1-1h3c.6 0 1 .4 1 1v2c0 .6-.4 1-1 1m7 0h-3c-.6 0-1-.4-1-1v-2c0-.6.4-1 1-1h3c.6 0 1 .4 1 1v2c0 .6-.4 1-1 1m-14 6h-3c-.6 0-1-.4-1-1v-2c0-.6.4-1 1-1h3c.6 0 1 .4 1 1v2c0 .6-.4 1-1 1m7 0h-3c-.6 0-1-.4-1-1v-2c0-.6.4-1 1-1h3c.6 0 1 .4 1 1v2c0 .6-.4 1-1 1m7 0h-3c-.6 0-1-.4-1-1v-2c0-.6.4-1 1-1h3c.6 0 1 .4 1 1v2c0 .6-.4 1-1 1"/><path fill="#BDBDBD" d="M36 29h-3c-.6 0-1-.4-1-1v-2c0-.6.4-1 1-1h3c.6 0 1 .4 1 1v2c0 .6-.4 1-1 1m0 6h-3c-.6 0-1-.4-1-1v-2c0-.6.4-1 1-1h3c.6 0 1 .4 1 1v2c0 .6-.4 1-1 1m0 6h-3c-.6 0-1-.4-1-1v-2c0-.6.4-1 1-1h3c.6 0 1 .4 1 1v2c0 .6-.4 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#CFD8DC" d="M5 38V14h38v24c0 2.2-1.8 4-4 4H9c-2.2 0-4-1.8-4-4"/><path fill="#F44336" d="M43 10v6H5v-6c0-2.2 1.8-4 4-4h30c2.2 0 4 1.8 4 4"/><g fill="#B71C1C"><circle cx="33" cy="10" r="3"/><circle cx="15" cy="10" r="3"/></g><path fill="#B0BEC5" d="M33 3c-1.1 0-2 .9-2 2v5c0 1.1.9 2 2 2s2-.9 2-2V5c0-1.1-.9-2-2-2M15 3c-1.1 0-2 .9-2 2v5c0 1.1.9 2 2 2s2-.9 2-2V5c0-1.1-.9-2-2-2"/><path fill="#90A4AE" d="M13 20h4v4h-4zm6 0h4v4h-4zm6 0h4v4h-4zm6 0h4v4h-4zm-18 6h4v4h-4zm6 0h4v4h-4zm6 0h4v4h-4zm6 0h4v4h-4zm-18 6h4v4h-4zm6 0h4v4h-4zm6 0h4v4h-4zm6 0h4v4h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CallTransfer(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#009688" d="m39.2 8.4l-1.8 1.8c-6.3 6.5-5.4 22 0 27.6l1.8 1.8c.5.5 1.3.5 1.8 0l3.6-3.7c.5-.5.5-1.3 0-1.8l-3.4-3.4h-4.8c-1.3-1.3-1.3-12.1 0-13.4h4.8l3.3-3.4c.5-.5.5-1.3 0-1.8L41 8.4c-.5-.5-1.3-.5-1.8 0m-28 0l-1.8 1.8c-6.3 6.5-5.4 22 0 27.6l1.8 1.8c.5.5 1.3.5 1.8 0l3.6-3.7c.5-.5.5-1.3 0-1.8l-3.4-3.4H8.5c-1.3-1.3-1.3-12.1 0-13.4h4.8l3.3-3.4c.5-.5.5-1.3 0-1.8L13 8.4c-.5-.5-1.3-.5-1.8 0"/><g fill="#2196F3"><path d="m25.3 18.6l5.4 5.4l-5.4 5.4z"/><path d="M16 22h11v4H16z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Callback(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#2196F3" d="M26.4 33.9s4-2.6 4.8-3c.8-.4 1.7-.6 2.2-.2c.8.5 7.5 4.9 8.1 5.3c.6.4.8 1.5.1 2.6c-.8 1.1-4.3 5.5-5.8 5.4c-1.5 0-8.4.4-20.3-11.4C3.6 20.7 4 13.8 4 12.3s4.3-5.1 5.4-5.8c1.1-.8 2.2-.5 2.6.1c.4.6 4.8 7.3 5.3 8.1c.3.5.2 1.4-.2 2.2c-.4.8-3 4.8-3 4.8s.7 2.8 5 7.2c4.4 4.3 7.3 5 7.3 5"/><g fill="#3F51B5"><path d="M35 9H25v4h10c1.1 0 2 .9 2 2v10h4V15c0-3.3-2.7-6-6-6"/><path d="m28 16l-6.7-5L28 6z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Camcorder(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#607D8B" d="M20 42H10c-2.2 0-4-1.8-4-4V15c0-5 4-9 9-9s9 4 9 9v23c0 2.2-1.8 4-4 4"/><circle cx="15" cy="15" r="7" fill="#455A64"/><circle cx="15" cy="15" r="5.2" fill="#42A5F5"/><path fill="#90CAF9" d="M18.3 13c-.8-.9-2-1.5-3.3-1.5s-2.4.5-3.3 1.5c-.3.4-.3.9.1 1.2c.4.3.9.3 1.2-.1c1-1.2 2.9-1.2 3.9 0c.2.2.4.3.7.3c.2 0 .4-.1.6-.2c.4-.3.4-.9.1-1.2"/><path fill="#607D8B" d="M40 31H28c-1.1 0-2-.9-2-2V19c0-1.1.9-2 2-2h12c1.1 0 2 .9 2 2v10c0 1.1-.9 2-2 2"/><path fill="#455A64" d="M24 19h2v10h-2z"/><path fill="#03A9F4" d="M28 19h12v10H28z"/><path fill="#4FC3F7" d="M33 22.2L29 28h8z"/><g fill="#B3E5FC"><circle cx="37.5" cy="21.5" r="1"/><path d="M36 24.2L33 28h6z"/></g><circle cx="15" cy="35" r="3" fill="#455A64"/><circle cx="15" cy="35" r="2" fill="#F44336"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CamcorderPro(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#37474F" d="M27 8h10v4H27z"/><path fill="#607D8B" d="M27 8h-9.7c-1.5 0-2.8.8-3.5 2.1l-3.3 6L14 18l3.3-6H27v7.2h4V12c0-2.2-1.8-4-4-4"/><path fill="#607D8B" d="M30 40H6c-2.2 0-4-1.8-4-4V20c0-2.2 1.8-4 4-4h24c2.2 0 4 1.8 4 4v16c0 2.2-1.8 4-4 4m8-5l5 5h3V18h-3l-5 5z"/><path fill="#546E7A" d="M22 35H8c-1.1 0-2-.9-2-2V23c0-1.1.9-2 2-2h14c1.1 0 2 .9 2 2v10c0 1.1-.9 2-2 2"/><path fill="#455A64" d="M34 23h4v12h-4z"/><path fill="#263238" d="M41 13h-4c-.6 0-1-.4-1-1V8c0-.6.4-1 1-1h4c1.7 0 3 1.3 3 3s-1.3 3-3 3"/><path fill="#03A9F4" d="M8 23h14v10H8z"/><path fill="#4FC3F7" d="M13.5 25.5L9 32h9z"/><g fill="#B3E5FC"><circle cx="19.5" cy="25.5" r="1.5"/><path d="M17.5 27.6L14 32h7z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Camera(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#512DA8" d="M33.9 12.1H14.2L17.6 7c.4-.6 1-.9 1.7-.9h9.6c.7 0 1.3.3 1.7.9z"/><path fill="#8667C4" d="M14 11H8V9.2C8 8.5 8.5 8 9.2 8h3.6c.7 0 1.2.5 1.2 1.2z"/><path fill="#5E35B1" d="M40 42H8c-2.2 0-4-1.8-4-4V14c0-2.2 1.8-4 4-4h32c2.2 0 4 1.8 4 4v24c0 2.2-1.8 4-4 4"/><circle cx="24" cy="26" r="12" fill="#512DA8"/><circle cx="24" cy="26" r="9" fill="#B388FF"/><path fill="#C7A7FF" d="M29 23c-1.2-1.4-3-2.2-4.8-2.2c-1.8 0-3.6.8-4.8 2.2c-.5.5-.4 1.3.1 1.8s1.3.4 1.8-.1c1.5-1.7 4.3-1.7 5.8 0c.3.3.6.4 1 .4c.3 0 .6-.1.9-.3c.4-.4.5-1.3 0-1.8"/><ellipse cx="11" cy="13.5" fill="#8667C4" rx="2" ry="1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraAddon(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#512DA8" d="M33.9 12.1H14.2L17.6 7c.4-.6 1-.9 1.7-.9h9.6c.7 0 1.3.3 1.7.9z"/><path fill="#8667C4" d="M14 11H8V9.2C8 8.5 8.5 8 9.2 8h3.6c.7 0 1.2.5 1.2 1.2z"/><path fill="#5E35B1" d="M40 42H8c-2.2 0-4-1.8-4-4V14c0-2.2 1.8-4 4-4h32c2.2 0 4 1.8 4 4v24c0 2.2-1.8 4-4 4"/><circle cx="24" cy="26" r="12" fill="#512DA8"/><circle cx="24" cy="26" r="9" fill="#B388FF"/><path fill="#C7A7FF" d="M28.8 23c-1.2-1.4-3-2.2-4.8-2.2s-3.6.8-4.8 2.2c-.5.5-.4 1.3.1 1.8s1.3.4 1.8-.1c1.5-1.7 4.3-1.7 5.8 0c.3.3.6.4 1 .4c.3 0 .6-.1.9-.3c.4-.4.5-1.3 0-1.8"/><ellipse cx="11" cy="13.5" fill="#8667C4" rx="2" ry="1.5"/><path fill="#8BC34A" d="M48 33.8c0-1.3-1.1-2.4-2.4-2.4H42c-.4 0-.7-.5-.4-.8c.4-.6.5-1.3.4-2.1c-.2-1.2-1.1-2.1-2.3-2.4c-2-.4-3.7 1-3.7 2.9c0 .6.2 1.1.4 1.6c.2.4 0 .8-.5.8h-3.6c-1.3 0-2.4 1.1-2.4 2.4V37c0 .4.5.7.8.4c.6-.4 1.3-.5 2.1-.4c1.2.2 2.1 1.1 2.4 2.3c.4 1.9-1.1 3.6-2.9 3.6c-.6 0-1.1-.2-1.6-.4c-.4-.2-.8 0-.8.5v2.6c0 1.3 1.1 2.4 2.4 2.4h13.2c1.3 0 2.4-1.1 2.4-2.4V33.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraIdentification(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#512DA8" d="M33.9 12.1H14.2L17.6 7c.4-.6 1-.9 1.7-.9h9.6c.7 0 1.3.3 1.7.9z"/><path fill="#8667C4" d="M14 11H8V9.2C8 8.5 8.5 8 9.2 8h3.6c.7 0 1.2.5 1.2 1.2z"/><path fill="#5E35B1" d="M40 42H8c-2.2 0-4-1.8-4-4V14c0-2.2 1.8-4 4-4h32c2.2 0 4 1.8 4 4v24c0 2.2-1.8 4-4 4"/><circle cx="24" cy="26" r="12" fill="#512DA8"/><circle cx="24" cy="26" r="9" fill="#B388FF"/><g fill="#616161"><path d="m38.912 40.703l1.696-1.697l7.353 7.353l-1.697 1.696z"/><circle cx="35" cy="35" r="10"/></g><path fill="#37474F" d="m42.305 44.106l1.697-1.696l3.96 3.959l-1.698 1.697z"/><circle cx="35" cy="35" r="8" fill="#64B5F6"/><path fill="#BBDEFB" d="M39.3 31.4c-1.1-1.3-2.6-2-4.2-2s-3.2.7-4.2 2c-.2.3-.2.6.1.9c.3.2.6.2.9-.1c.8-1 2-1.5 3.3-1.5s2.5.6 3.3 1.5c.1.1.3.2.5.2c.1 0 .3 0 .4-.1c.1-.2.1-.6-.1-.9"/><path fill="#C7A7FF" d="M29 23c-1.2-1.4-3-2.2-4.8-2.2c-1.8 0-3.6.8-4.8 2.2c-.5.5-.4 1.3.1 1.8s1.3.4 1.8-.1c1.5-1.7 4.3-1.7 5.8 0c.3.3.6.4 1 .4c.3 0 .6-.1.9-.3c.4-.4.5-1.3 0-1.8"/><ellipse cx="11" cy="13.5" fill="#8667C4" rx="2" ry="1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cancel(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#D50000" d="M24 6C14.1 6 6 14.1 6 24s8.1 18 18 18s18-8.1 18-18S33.9 6 24 6m0 4c3.1 0 6 1.1 8.4 2.8L12.8 32.4C11.1 30 10 27.1 10 24c0-7.7 6.3-14 14-14m0 28c-3.1 0-6-1.1-8.4-2.8l19.6-19.6C36.9 18 38 20.9 38 24c0 7.7-6.3 14-14 14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CandleSticks(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#546E7A" d="M38 4h2v20h-2zM15 7h2v17h-2zM8 27h2v17H8zm20-8h2v22h-2z"/><path fill="#4CAF50" d="M36 7h6c1.1 0 2 .9 2 2v10c0 1.1-.9 2-2 2h-6c-1.1 0-2-.9-2-2V9c0-1.1.9-2 2-2m-23 3h6c1.1 0 2 .9 2 2v7c0 1.1-.9 2-2 2h-6c-1.1 0-2-.9-2-2v-7c0-1.1.9-2 2-2"/><path fill="#F44336" d="M6 30h6c1.1 0 2 .9 2 2v7c0 1.1-.9 2-2 2H6c-1.1 0-2-.9-2-2v-7c0-1.1.9-2 2-2m20-8h6c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2h-6c-1.1 0-2-.9-2-2V24c0-1.1.9-2 2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Capacitor(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#FF9800" d="M0 27h25v4H0zm0-10h25v4H0z"/><path fill="#3F51B5" d="M46 35c1.1 0 2-.9 2-2V15c0-1.1-.9-2-2-2H27v22zM21 13c-1.1 0-2 .9-2 2v18c0 1.1.9 2 2 2h2V13z"/><path fill="#303F9F" d="M25 33c1.1 0 2 .9 2 2V13c0 1.1-.9 2-2 2s-2-.9-2-2v22c0-1.1.9-2 2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CdLogo(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#2196F3" d="M17.814 18H20.2c.5 0 .8.4.8.8v16.4c0 .399-.4.8-.8.8h-2.384c-.4 0-.8-.4-.8-.8V18.8c-.1-.5.298-.8.798-.8M14.2 11h-3.3c-.5 0-.9.403-.9.807V17H2.2c-.6 0-1.2.605-1.2 1.21v16.58c0 .606.6 1.21 1.2 1.21h12c.4 0 .8-.305.8-.809V11.807c0-.404-.3-.807-.8-.807M10 31.283c0 .398-.4.8-.8.8H6.8c-.4 0-.8-.399-.8-.8V21.8c0-.4.3-.8.8-.8h2.4c.5 0 .8.4.8.8zM33.2 25c.5 0 .8.6.8.8v9.4c0 .399-.422.8-.8.8h-9.4c-.425 0-.8-.4-.8-.8v-2.386c0-.5.4-.799.8-.799L30 32v-3h-6.2c-.331 0-.8-.4-.8-.801V18.8c0-.5.4-.8.8-.8h9.4c.399 0 .8.4.8.8v2.4c0 .3-.266.8-.8.8H27v3zM48 28v7.2c0 .399-.4.8-.801.8H36.8c-.2 0-.8-.4-.8-.8V18.8c0-.5.432-.8.831-.8H47.2s.8 0 .8.8V25h-4v-2.2s.1-.8-.8-.8h-2.4c-.5 0-.8.4-.8.8v8.4c0 .399.5.8.8.8h2.4c.399 0 .8-.4.8-.8V28z"/><path fill="#0D47A1" d="M45.799 15.98H46.9v-3.099H48v-.901h-3.201v.901h1zm-1.785-1.504h-1.143v.095c0 .382-.096.573-.572.573c-.475 0-.57-.191-.57-.762v-.668c0-.572 0-.762.57-.762c.381 0 .572.095.572.477v.095h1.047v-.095c0-1.047-.381-1.429-1.523-1.429h-.287c-1.141 0-1.523.382-1.523 1.618v.764c0 1.142.381 1.618 1.523 1.618h.383c1.047 0 1.428-.477 1.428-1.43v-.095h.095zM29.516 12l-.799 2.9l-.901-2.9h-1.599v4h1.099v-3l.799 3h1.102l.898-2.9V16h1v-4zm9.197 0h-1.798l-1.199 4h1.199l.199-.7h1.401l.199.7h1.199zm-.799 2.5h-.4l.4-1.7l.398 1.7zM33.92 12h-1.9v4h1.102v-1.1h.799c1.102 0 1.5-.4 1.5-1.4v-.3c-.001-.9-.401-1.2-1.501-1.2m.4 1.6c0 .4-.1.5-.4.5h-.799v-1.2h.7c.4 0 .5.1.5.4v.3zM23.594 12h-.572c-1.143 0-1.523.382-1.523 1.618v.762c0 1.144.381 1.62 1.523 1.62h.572c1.143 0 1.523-.477 1.523-1.62v-.762c0-1.236-.38-1.618-1.523-1.618m.381 2.19c0 .571 0 .763-.571.763h-.19c-.571 0-.571-.191-.571-.763v-.572c0-.57 0-.762.571-.762h.19c.571 0 .571.191.571.762zm-3.553.287h-1.144v.095c0 .382-.095.571-.571.571s-.571-.189-.571-.762v-.666c0-.573 0-.762.571-.762c.381 0 .571.095.571.475v.096h1.048v-.096c0-1.047-.381-1.428-1.523-1.428h-.286c-1.143 0-1.524.381-1.524 1.618v.763c0 1.143.381 1.619 1.524 1.619h.381c1.048 0 1.429-.477 1.429-1.429v-.095h.095z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CellPhone(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#546E7A" d="M12 40V10h20c2.2 0 4 1.8 4 4v26c0 2.2-1.8 4-4 4H16c-2.2 0-4-1.8-4-4"/><path fill="#4FC3F7" d="M32 13H16c-.6 0-1 .4-1 1v8c0 .6.4 1 1 1h16c.6 0 1-.4 1-1v-8c0-.6-.4-1-1-1"/><path fill="#B3E5FC" d="M19 30h-2c-.6 0-1-.4-1-1v-1c0-.6.4-1 1-1h2c.6 0 1 .4 1 1v1c0 .6-.4 1-1 1m6 0h-2c-.6 0-1-.4-1-1v-1c0-.6.4-1 1-1h2c.6 0 1 .4 1 1v1c0 .6-.4 1-1 1m6 0h-2c-.6 0-1-.4-1-1v-1c0-.6.4-1 1-1h2c.6 0 1 .4 1 1v1c0 .6-.4 1-1 1m-12 5h-2c-.6 0-1-.4-1-1v-1c0-.6.4-1 1-1h2c.6 0 1 .4 1 1v1c0 .6-.4 1-1 1m6 0h-2c-.6 0-1-.4-1-1v-1c0-.6.4-1 1-1h2c.6 0 1 .4 1 1v1c0 .6-.4 1-1 1m6 0h-2c-.6 0-1-.4-1-1v-1c0-.6.4-1 1-1h2c.6 0 1 .4 1 1v1c0 .6-.4 1-1 1m-12 5h-2c-.6 0-1-.4-1-1v-1c0-.6.4-1 1-1h2c.6 0 1 .4 1 1v1c0 .6-.4 1-1 1m6 0h-2c-.6 0-1-.4-1-1v-1c0-.6.4-1 1-1h2c.6 0 1 .4 1 1v1c0 .6-.4 1-1 1m6 0h-2c-.6 0-1-.4-1-1v-1c0-.6.4-1 1-1h2c.6 0 1 .4 1 1v1c0 .6-.4 1-1 1"/><path fill="#37474F" d="M16 10h-4V4c0-1.1.9-2 2-2s2 .9 2 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChargeBattery(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#8BC34A"><path d="M34 44H14c-1.1 0-2-.9-2-2V8c0-1.1.9-2 2-2h20c1.1 0 2 .9 2 2v34c0 1.1-.9 2-2 2"/><path d="M28 13h-8c-.6 0-1-.4-1-1V5c0-.6.4-1 1-1h8c.6 0 1 .4 1 1v7c0 .6-.4 1-1 1"/></g><path fill="#FFEB3B" d="M30 24h-5.5l2.2-11L18 26h5.5l-2.2 11z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Checkmark(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#43A047" d="M40.6 12.1L17 35.7l-9.6-9.6L4.6 29L17 41.3l26.4-26.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Circuit(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#4CAF50" d="M6 10v28c0 2.2 1.8 4 4 4h28c2.2 0 4-1.8 4-4V10c0-2.2-1.8-4-4-4H10c-2.2 0-4 1.8-4 4"/><path fill="#FFC107" d="m6.6 8l6 6c-.4.6-.6 1.3-.6 2c0 2.2 1.8 4 4 4s4-1.8 4-4s-1.8-4-4-4c-.7 0-1.4.2-2 .6l-6-6c-.6.3-1.1.8-1.4 1.4m9.4 6.5c.8 0 1.5.7 1.5 1.5s-.7 1.5-1.5 1.5s-1.5-.7-1.5-1.5s.7-1.5 1.5-1.5M41.4 40l-6-6c.4-.6.6-1.3.6-2c0-2.2-1.8-4-4-4s-4 1.8-4 4s1.8 4 4 4c.7 0 1.4-.2 2-.6l6 6c.6-.3 1.1-.8 1.4-1.4M32 33.5c-.8 0-1.5-.7-1.5-1.5s.7-1.5 1.5-1.5s1.5.7 1.5 1.5s-.7 1.5-1.5 1.5M16 36c2.2 0 4-1.8 4-4c0-.7-.2-1.4-.6-2L30 19.4c.6.4 1.3.6 2 .6c2.2 0 4-1.8 4-4s-1.8-4-4-4s-4 1.8-4 4c0 .7.2 1.4.6 2L18 28.6c-.6-.4-1.3-.6-2-.6c-2.2 0-4 1.8-4 4s1.8 4 4 4m16-21.5c.8 0 1.5.7 1.5 1.5s-.7 1.5-1.5 1.5s-1.5-.7-1.5-1.5s.7-1.5 1.5-1.5m-16 16c.8 0 1.5.7 1.5 1.5s-.7 1.5-1.5 1.5s-1.5-.7-1.5-1.5s.7-1.5 1.5-1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clapperboard(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#3F51B5" d="M43.4 8.3L4 15l-.3-2C3.3 10.8 4.8 8.8 7 8.4l31.6-5.3c2.2-.4 4.2 1.1 4.6 3.3zM40 41H8c-2.2 0-4-1.8-4-4V15h40v22c0 2.2-1.8 4-4 4"/><path fill="#9FA8DA" d="m18.8 6.4l4.9 5.3l4-.7l-5-5.3zm-7.9 1.3l4.9 5.3l4-.7l-5-5.2zm15.8-2.6l4.9 5.2l3.9-.6l-4.9-5.3zm7.8-1.3l5 5.2l3.9-.7l-4.9-5.2z"/><circle cx="7.5" cy="11.5" r="1.5" fill="#9FA8DA"/><path fill="#9FA8DA" d="m40 15l-4 6h4l4-6zm-8 0l-4 6h4l4-6zm-8 0l-4 6h4l4-6zm-8 0l-4 6h4l4-6zm-8 0l-4 6h4l4-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClearFilters(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#F57C00" d="M29 23H19L7 9h34z"/><path fill="#FF9800" d="m29 38l-10 6V23h10zM41.5 9h-35C5.7 9 5 8.3 5 7.5S5.7 6 6.5 6h35c.8 0 1.5.7 1.5 1.5S42.3 9 41.5 9"/><circle cx="38" cy="38" r="10" fill="#F44336"/><path fill="#fff" d="M32 36h12v4H32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clock(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="24" cy="24" r="20" fill="#00ACC1"/><circle cx="24" cy="24" r="16" fill="#eee"/><path d="M23 11h2v13h-2z"/><path d="M31.285 29.654L29.66 31.28l-6.504-6.504l1.626-1.627z"/><circle cx="24" cy="24" r="2"/><circle cx="24" cy="24" r="1" fill="#00ACC1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloseUpMode(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#2E7D32" d="M5 24c0 18.9 14.8 19 17 19h4S24.1 26.5 5 24"/><path fill="#388E3C" d="M22 26h4v17h-4z"/><path fill="#C62828" d="M34 16c0 5.1-5.2 8.2-8 8.2s-2-3.1-2-8.2s5-9.2 5-9.2s5 4.1 5 9.2"/><path fill="#C62828" d="M14 16c0 5.1 5.2 8.2 8 8.2s2-3.1 2-8.2s-5-9.2-5-9.2s-5 4.1-5 9.2"/><path fill="#E53935" d="M24 27c-2.2-1.6-1.9-4.5 2.4-8.8C30.8 13.8 32 7 32 7s5 3.4 5 9c0 5.9-5.7 11-13 11"/><path fill="#E53935" d="M24 27c2.2-1.6 1.9-4.5-2.4-8.8C17.2 13.8 16 7 16 7s-5 3.4-5 9c0 5.9 5.7 11 13 11"/><path fill="#F44336" d="M30 16c0 6.1-2.7 11-6 11s-6-4.9-6-11s6-11 6-11s6 4.9 6 11"/><path fill="#4CAF50" d="M22 43h4c2.2 0 17-.1 17-19c-19.1 2.5-21 19-21 19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloth(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#FF5722" d="M6 10v28c0 2.2 1.8 4 4 4h28c2.2 0 4-1.8 4-4V10c0-2.2-1.8-4-4-4H10c-2.2 0-4 1.8-4 4"/><path fill="#BF360C" d="M6 35h36v2H6zm0-4h36v2H6zm.1 8c.2.8.6 1.5 1.2 2h33.2c.6-.5 1-1.2 1.2-2zm0-30h35.7c-.2-.8-.6-1.5-1.2-2H7.4c-.6.5-1.1 1.2-1.3 2M6 23h36v2H6zm0 4h36v2H6zm0-12h36v2H6zm0-4h36v2H6zm0 8h36v2H6z"/><path fill="#FF8A65" d="M27 6h2v5h-2zm0 7h2v6h-2zm0 16h2v6h-2zm4-23h2v1h-2zM19 29h2v6h-2zM31 9h2v6h-2zm-8-3h2v1h-2zm0 19h2v6h-2zm0-16h2v6h-2zm-4 12h2v6h-2zm4-4h2v6h-2zm0 16h2v6h-2zm4-12h2v6h-2zm12 12h2v6h-2zm0-16h2v6h-2zm0 8h2v6h-2zm0-18.9V7h1.6c-.4-.4-1-.7-1.6-.9M31 17h2v6h-2zm9.6 24H39v.9c.6-.2 1.2-.5 1.6-.9M35 13h2v6h-2zm-4 20h2v6h-2zm4-4h2v6h-2zm4-20h2v6h-2zm-4 12h2v6h-2zm-4 4h2v6h-2zm4 12h2v5h-2zm0-31h2v5h-2zm-4 35h2v1h-2zm-8 0h2v1h-2zm4-4h2v5h-2zm-8 0h2v5h-2zM7 17h2v6H7zm2 24H7.4c.5.4 1 .7 1.6.9zM7.4 7H9v-.9c-.6.2-1.2.5-1.6.9M7 33h2v6H7zm0-8h2v6H7zM7 9h2v6H7zm4 20h2v6h-2zm4-12h2v6h-2zm0 16h2v6h-2zm0-24h2v6h-2zm0-3h2v1h-2zm4 0h2v5h-2zm-4 19h2v6h-2zm0 16h2v1h-2zm-4-20h2v6h-2zm0-15h2v5h-2zm0 31h2v5h-2zm8-24h2v6h-2zm-8 0h2v6h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Collaboration(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#1565C0" d="M25 22h13l6 6V11c0-2.2-1.8-4-4-4H25c-2.2 0-4 1.8-4 4v7c0 2.2 1.8 4 4 4"/><path fill="#2196F3" d="M23 19H10l-6 6V8c0-2.2 1.8-4 4-4h15c2.2 0 4 1.8 4 4v7c0 2.2-1.8 4-4 4"/><g fill="#FFA726"><circle cx="12" cy="31" r="5"/><circle cx="36" cy="31" r="5"/></g><path fill="#607D8B" d="M20 42s-2.2-4-8-4s-8 4-8 4v2h16zm24 0s-2.2-4-8-4s-8 4-8 4v2h16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Collapse(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#2196F3" d="M5 30.9L8.1 34L24 18.1L39.9 34l3.1-3.1L24 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Collect(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#009688"><path d="M22 35h4v11h-4z"/><path d="m24 29.6l7 8.4H17z"/></g><g fill="#009688"><path d="M22 2h4v11h-4z"/><path d="M24 18.4L17 10h14z"/></g><g fill="#009688"><path d="M2 22h11v4H2z"/><path d="M18.4 24L10 31V17z"/></g><g fill="#009688"><path d="M35 22h11v4H35z"/><path d="m29.6 24l8.4-7v14z"/></g><circle cx="24" cy="24" r="3" fill="#F44336"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ComboChart(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#00BCD4" d="M37 18h6v24h-6zm-8 8h6v16h-6zm-8-4h6v20h-6zm-8 10h6v10h-6zm-8-4h6v14H5z"/><g fill="#3F51B5"><circle cx="8" cy="16" r="3"/><circle cx="16" cy="18" r="3"/><circle cx="24" cy="11" r="3"/><circle cx="32" cy="13" r="3"/><circle cx="40" cy="9" r="3"/><path d="m39.1 7.2l-7.3 3.7l-8.3-2.1l-8 7l-7-1.7l-1 3.8l9 2.3l8-7l7.7 1.9l8.7-4.3z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommandLine(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#CFD8DC" d="M41 6H7c-.6 0-1 .4-1 1v35h36V7c0-.6-.4-1-1-1"/><path fill="#263238" d="M8 13h32v27H8z"/><path fill="#76FF03" d="M22 27.6c-.1 1.1-.4 1.9-1 2.5c-.6.6-1.4.9-2.5.9s-2-.4-2.6-1.1c-.6-.7-.9-1.8-.9-3.1v-1.6c0-1.3.3-2.4.9-3.1c.6-.7 1.5-1.1 2.6-1.1s1.9.3 2.5.9c.6.6.9 1.4 1 2.6h-2c0-.7-.1-1.2-.3-1.4c-.2-.3-.6-.4-1.1-.4c-.5 0-.9.2-1.2.6c-.2.4-.3 1-.4 1.8v1.8c0 1 .1 1.6.3 2c.2.4.6.5 1.1.5c.5 0 .9-.1 1.1-.4c.2-.3.3-.7.3-1.4zm2-3.6c0-.3.1-.5.3-.7c.2-.2.4-.3.7-.3c.3 0 .5.1.7.3c.2.2.3.4.3.7c0 .3-.1.5-.3.7s-.4.3-.7.3c-.3 0-.5-.1-.7-.3s-.3-.4-.3-.7m0 6c0-.3.1-.5.3-.7c.2-.2.4-.3.7-.3c.3 0 .5.1.7.3c.2.2.3.4.3.7c0 .3-.1.5-.3.7s-.4.3-.7.3c-.3 0-.5-.1-.7-.3s-.3-.4-.3-.7m4-9h2l3 10h-2z"/><g fill="#90A4AE"><circle cx="13.5" cy="9.5" r="1.5"/><circle cx="9.5" cy="9.5" r="1.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Comments(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#8BC34A" d="M37 39H11l-6 6V11c0-3.3 2.7-6 6-6h26c3.3 0 6 2.7 6 6v22c0 3.3-2.7 6-6 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompactCamera(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#607D8B" d="M40 39H8c-2.2 0-4-1.8-4-4V13c0-2.2 1.8-4 4-4h32c2.2 0 4 1.8 4 4v22c0 2.2-1.8 4-4 4"/><circle cx="29" cy="24" r="12" fill="#455A64"/><circle cx="29" cy="24" r="9" fill="#42A5F5"/><path fill="#90CAF9" d="M33.8 21c-1.2-1.4-3-2.2-4.8-2.2s-3.6.8-4.8 2.2c-.5.5-.4 1.3.1 1.8s1.3.4 1.8-.1c1.5-1.7 4.3-1.7 5.8 0c.3.3.6.4 1 .4c.3 0 .6-.1.9-.3c.4-.4.5-1.3 0-1.8"/><path fill="#ADD8FB" d="M8 13h6v3H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ConferenceCall(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="12" cy="21" r="5" fill="#FFA726"/><path fill="#455A64" d="M2 34.7s2.8-6.3 10-6.3s10 6.3 10 6.3V38H2zm44 0s-2.8-6.3-10-6.3s-10 6.3-10 6.3V38h20z"/><circle cx="24" cy="17" r="6" fill="#FFB74D"/><path fill="#607D8B" d="M36 34.1s-3.3-7.5-12-7.5s-12 7.5-12 7.5V38h24z"/><circle cx="36" cy="21" r="5" fill="#FFA726"/><circle cx="12" cy="21" r="5" fill="#FFA726"/><circle cx="36" cy="21" r="5" fill="#FFA726"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Contacts(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#FF7043" d="M38 44H12V4h26c2.2 0 4 1.8 4 4v32c0 2.2-1.8 4-4 4"/><path fill="#BF360C" d="M10 4h2v40h-2c-2.2 0-4-1.8-4-4V8c0-2.2 1.8-4 4-4"/><g fill="#AB300B"><circle cx="26" cy="20" r="4"/><path d="M33 30s-1.9-4-7-4s-7 4-7 4v2h14z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copyleft(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="24" cy="24" r="21" fill="#9C27B0"/><path fill="#E1BEE7" d="M19.3 28.1c.3 1.3.2 4.1 4.6 4.1c.9 0 4.8.2 4.7-7.2v-1.6c0-6.7-3.2-7.2-4.8-7.2c-2.3 0-4.2.6-4.5 4.3h-4.8c.1-1.2.8-8.2 9.3-8.2c4.2 0 9.7 2.5 9.7 11.2V25c0 9.6-6.5 11.2-9.6 11.2c-3.7 0-8.7-1.6-9.3-8.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copyright(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="24" cy="24" r="21" fill="#9C27B0"/><path fill="#E1BEE7" d="M33.5 28.1c-.6 6.4-5.6 8.1-9.3 8.1c-3.1 0-9.6-1.6-9.6-11.2v-1.5c0-8.7 5.5-11.2 9.7-11.2c8.5 0 9.2 7 9.3 8.2h-4.8c-.3-3.6-2.2-4.3-4.5-4.3c-1.6 0-4.8.5-4.8 7.2V25c-.1 7.5 3.8 7.2 4.7 7.2c4.3 0 4.3-2.9 4.6-4.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CrystalOscillator(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#FF9800" d="M3 28h26v4H3zm0-12h26v4H3z"/><path fill="#2196F3" d="M43 11H20v26h23c1.1 0 2-.9 2-2V13c0-1.1-.9-2-2-2"/><path fill="#64B5F6" d="M20 9h-2v30h2c1.1 0 2-.9 2-2V11c0-1.1-.9-2-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CurrencyExchange(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="18" cy="18" r="15" fill="#3F51B5"/><path fill="#FFF59D" d="M20.3 16v1.7h-3.8v1.4h3.8v1.7h-3.8c0 .6.1 1.2.3 1.6c.2.4.4.8.7 1c.3.3.7.4 1.1.6c.4.1.9.2 1.4.2c.4 0 .7 0 1.1-.1c.4-.1.7-.1 1-.3l.4 2.7c-.4.1-.9.2-1.4.2c-.5.1-1 .1-1.5.1c-.9 0-1.8-.1-2.6-.4c-.8-.2-1.5-.6-2-1.1c-.6-.5-1-1.1-1.4-1.9c-.3-.7-.5-1.6-.5-2.6h-1.9v-1.7h1.9v-1.4h-1.9V16h1.9c.1-1 .3-1.8.6-2.6c.4-.7.8-1.4 1.4-1.9c.6-.5 1.3-.9 2.1-1.1c.8-.3 1.7-.4 2.6-.4c.4 0 .9 0 1.3.1s.9.1 1.3.3l-.4 2.7c-.3-.1-.6-.2-1-.3c-.4-.1-.7-.1-1.1-.1c-.5 0-1 .1-1.4.2c-.4.1-.8.3-1 .6c-.3.3-.5.6-.7 1s-.3.9-.3 1.5z"/><circle cx="30" cy="30" r="15" fill="#4CAF50"/><path fill="#fff" d="M28.4 27c.1.2.2.4.4.6c.2.2.4.4.7.5c.3.2.7.3 1.1.5c.7.3 1.4.6 2 .9c.6.3 1.1.7 1.5 1.1c.4.4.8.9 1 1.4c.2.5.4 1.2.4 1.9s-.1 1.3-.3 1.8c-.2.5-.5 1-.9 1.4s-.9.7-1.4.9c-.6.2-1.2.4-1.8.5v2.2h-1.8v-2.2c-.6-.1-1.2-.2-1.8-.4s-1.1-.5-1.5-1c-.5-.4-.8-1-1.1-1.6c-.3-.6-.4-1.4-.4-2.3h3.3c0 .5.1 1 .2 1.3c.1.4.3.6.6.9c.2.2.5.4.8.5c.3.1.6.1.9.1c.4 0 .7 0 .9-.1c.3-.1.5-.2.7-.4c.2-.2.3-.4.4-.6c.1-.2.1-.5.1-.8c0-.3 0-.6-.1-.8c-.1-.2-.2-.5-.4-.7s-.4-.4-.7-.5c-.3-.2-.7-.3-1.1-.5c-.7-.3-1.4-.6-2-.9c-.6-.3-1.1-.7-1.5-1.1c-.4-.4-.8-.9-1-1.4c-.2-.5-.4-1.2-.4-1.9c0-.6.1-1.2.3-1.7c.2-.5.5-1 .9-1.4c.4-.4.9-.7 1.4-1c.5-.2 1.2-.4 1.8-.5v-2.4h1.8v2.4c.6.1 1.2.3 1.8.6c.5.3 1 .6 1.3 1.1c.4.4.7 1 .9 1.6c.2.6.3 1.3.3 2h-3.3c0-.9-.2-1.6-.6-2c-.4-.4-.9-.7-1.5-.7c-.3 0-.6.1-.9.2c-.2.1-.4.2-.6.4c-.2.2-.3.4-.3.6c-.1.2-.1.5-.1.8c-.1.2 0 .5 0 .7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cursor(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#E0E0E0" d="M27.8 39.7c-.1 0-.2 0-.4-.1s-.4-.3-.6-.5l-3.7-8.6l-4.5 4.2c-.1.2-.3.3-.6.3c-.1 0-.3 0-.4-.1c-.3-.1-.6-.5-.6-.9V12c0-.4.2-.8.6-.9c.1-.1.3-.1.4-.1c.2 0 .5.1.7.3l16 15c.3.3.4.7.3 1.1c-.1.4-.5.6-.9.7l-6.3.6l3.9 8.5c.1.2.1.5 0 .8c-.1.2-.3.5-.5.6l-2.9 1.3c-.2-.2-.4-.2-.5-.2"/><path fill="#212121" d="m18 12l16 15l-7.7.7l4.5 9.8l-2.9 1.3l-4.3-9.9L18 34zm0-2c-.3 0-.5.1-.8.2c-.7.3-1.2 1-1.2 1.8v22c0 .8.5 1.5 1.2 1.8c.3.2.6.2.8.2c.5 0 1-.2 1.4-.5l3.4-3.2l3.1 7.3c.2.5.6.9 1.1 1.1c.2.1.5.1.7.1c.3 0 .5-.1.8-.2l2.9-1.3c.5-.2.9-.6 1.1-1.1c.2-.5.2-1.1 0-1.5l-3.3-7.2l4.9-.4c.8-.1 1.5-.6 1.7-1.3c.3-.7.1-1.6-.5-2.1l-16-15c-.3-.5-.8-.7-1.3-.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CustomerSupport(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#FFB74D" d="M29 43v-4.6l2.6.5c2.9.6 5.6-1.5 5.8-4.4L38 28l2.9-1.2c1-.4 1.4-1.6.8-2.6L38 18c-.6-7.6-4.9-15-16-15C10.6 3 5 11.4 5 20c0 3.7 1.3 6.9 3.3 9.6c1.8 2.5 2.7 5.5 2.7 8.5v4.8h18z"/><path fill="#FF9800" d="M29 43v-4.6L22 37v6z"/><circle cx="33.5" cy="21.5" r="1.5" fill="#784719"/><path fill="#FF5722" d="M21.4 3C12.3 3 5 10.3 5 19.4c0 11.1 6 11.4 6 18.6l2.6-.9c2.1-.7 3.9-2.3 4.7-4.4l2.8-6.8L27 23v-6s7-3.8 7-10.3C31 4.2 25.7 3 21.4 3"/><path fill="#546E7A" d="M21 2.1c-.6 0-1 .4-1 1V17c0 .6.4 1 1 1s1-.4 1-1V3.1c0-.6-.4-1-1-1m15.9 29.8c-7.9 0-10.3-4.9-10.4-5.1c-.2-.5-.8-.7-1.3-.5c-.5.2-.7.8-.5 1.3c.1.3 3 6.3 12.2 6.3c.6 0 1-.4 1-1s-.5-1-1-1"/><circle cx="37" cy="33" r="2" fill="#37474F"/><circle cx="21" cy="23" r="7" fill="#37474F"/><circle cx="21" cy="23" r="4" fill="#546E7A"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dam(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#81D4FA" d="M24 28h18v14H24zM6 10h12v32H6z"/><path fill="#1976D2" d="M16 8h-2c0 1.1-.9 2-2 2s-2-.9-2-2H8c0 1.1-.9 2-2 2v2c1.2 0 2.3-.5 3-1.4c.7.8 1.8 1.4 3 1.4s2.3-.5 3-1.4c.7.8 1.8 1.4 3 1.4v-2c-1.1 0-2-.9-2-2m0 6h-2c0 1.1-.9 2-2 2s-2-.9-2-2H8c0 1.1-.9 2-2 2v2c1.2 0 2.3-.5 3-1.4c.7.8 1.8 1.4 3 1.4s2.3-.5 3-1.4c.7.8 1.8 1.4 3 1.4v-2c-1.1 0-2-.9-2-2m0 6h-2c0 1.1-.9 2-2 2s-2-.9-2-2H8c0 1.1-.9 2-2 2v2c1.2 0 2.3-.5 3-1.4c.7.8 1.8 1.4 3 1.4s2.3-.5 3-1.4c.7.8 1.8 1.4 3 1.4v-2c-1.1 0-2-.9-2-2m0 6h-2c0 1.1-.9 2-2 2s-2-.9-2-2H8c0 1.1-.9 2-2 2v2c1.2 0 2.3-.5 3-1.4c.7.8 1.8 1.4 3 1.4s2.3-.5 3-1.4c.7.8 1.8 1.4 3 1.4v-2c-1.1 0-2-.9-2-2m0 6h-2c0 1.1-.9 2-2 2s-2-.9-2-2H8c0 1.1-.9 2-2 2v2c1.2 0 2.3-.5 3-1.4c.7.8 1.8 1.4 3 1.4s2.3-.5 3-1.4c.7.8 1.8 1.4 3 1.4v-2c-1.1 0-2-.9-2-2m0 6h-2c0 1.1-.9 2-2 2s-2-.9-2-2H8c0 1.1-.9 2-2 2v2c1.2 0 2.3-.5 3-1.4c.7.8 1.8 1.4 3 1.4s2.3-.5 3-1.4c.7.8 1.8 1.4 3 1.4v-2c-1.1 0-2-.9-2-2m24-6h-2c0 1.1-.9 2-2 2s-2-.9-2-2h-2c0 1.1-.9 2-2 2s-2-.9-2-2h-2c0 1.1-.9 2-2 2v2c1.2 0 2.3-.5 3-1.4c.7.8 1.8 1.4 3 1.4s2.3-.5 3-1.4c.7.8 1.8 1.4 3 1.4s2.3-.5 3-1.4c.7.8 1.8 1.4 3 1.4v-2c-1.1 0-2-.9-2-2m0-6h-2c0 1.1-.9 2-2 2s-2-.9-2-2h-2c0 1.1-.9 2-2 2s-2-.9-2-2h-2c0 1.1-.9 2-2 2v2c1.2 0 2.3-.5 3-1.4c.7.8 1.8 1.4 3 1.4s2.3-.5 3-1.4c.7.8 1.8 1.4 3 1.4s2.3-.5 3-1.4c.7.8 1.8 1.4 3 1.4v-2c-1.1 0-2-.9-2-2m0 12h-2c0 1.1-.9 2-2 2s-2-.9-2-2h-2c0 1.1-.9 2-2 2v2c1.2 0 2.3-.5 3-1.4c.7.8 1.8 1.4 3 1.4s2.3-.5 3-1.4c.7.8 1.8 1.4 3 1.4v-2c-1.1 0-2-.9-2-2"/><path fill="#455A64" d="M25.1 9.2L31.5 42H18V6h3.2c1.9 0 3.6 1.4 3.9 3.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DataBackup(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#D1C4E9" d="M38 7H10c-1.1 0-2 .9-2 2v6c0 1.1.9 2 2 2h28c1.1 0 2-.9 2-2V9c0-1.1-.9-2-2-2m0 12H10c-1.1 0-2 .9-2 2v6c0 1.1.9 2 2 2h28c1.1 0 2-.9 2-2v-6c0-1.1-.9-2-2-2m0 12H10c-1.1 0-2 .9-2 2v6c0 1.1.9 2 2 2h28c1.1 0 2-.9 2-2v-6c0-1.1-.9-2-2-2"/><g fill="#2196F3"><path d="m31 30l7 5.6V24.4z"/><path d="M38 28c-.3 0-.7 0-1 .1v4c.3-.1.7-.1 1-.1c3.3 0 6 2.7 6 6s-2.7 6-6 6s-6-2.7-6-6c0-.3 0-.6.1-.9l-3.4-2.7c-.4 1.1-.7 2.3-.7 3.6c0 5.5 4.5 10 10 10s10-4.5 10-10s-4.5-10-10-10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DataConfiguration(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#D1C4E9" d="M38 7H10c-1.1 0-2 .9-2 2v6c0 1.1.9 2 2 2h28c1.1 0 2-.9 2-2V9c0-1.1-.9-2-2-2m0 12H10c-1.1 0-2 .9-2 2v6c0 1.1.9 2 2 2h28c1.1 0 2-.9 2-2v-6c0-1.1-.9-2-2-2m0 12H10c-1.1 0-2 .9-2 2v6c0 1.1.9 2 2 2h28c1.1 0 2-.9 2-2v-6c0-1.1-.9-2-2-2"/><path fill="#607D8B" d="M45.2 38.1c.1-.4.1-.8.1-1.1s0-.8-.1-1.1l2.3-1.7c.2-.2.3-.5.2-.7l-2.3-3.9c-.1-.2-.4-.3-.7-.2l-2.6 1.2c-.6-.5-1.3-.9-2-1.2l-.3-2.9c0-.3-.3-.5-.5-.5h-4.5c-.3 0-.5.2-.5.5l-.3 2.9c-.7.3-1.4.7-2 1.2l-2.6-1.2c-.3-.1-.6 0-.7.2l-2.3 3.9c-.1.2-.1.6.2.7l2.3 1.7c-.1.4-.1.8-.1 1.1s0 .8.1 1.1l-2.3 1.7c-.2.2-.3.5-.2.7l2.3 3.9c.1.2.4.3.7.2l2.6-1.2c.6.5 1.3.9 2 1.2l.3 2.9c0 .3.3.5.5.5h4.5c.3 0 .5-.2.5-.5l.3-2.9c.7-.3 1.4-.7 2-1.2l2.6 1.2c.3.1.6 0 .7-.2l2.3-3.9c.1-.2.1-.6-.2-.7zM37 42.2c-2.9 0-5.2-2.3-5.2-5.2c0-2.9 2.3-5.2 5.2-5.2c2.9 0 5.2 2.3 5.2 5.2c0 2.9-2.3 5.2-5.2 5.2"/><path fill="#455A64" d="M37 31c-3.3 0-6 2.7-6 6s2.7 6 6 6s6-2.7 6-6s-2.7-6-6-6m0 9c-1.7 0-3-1.3-3-3s1.3-3 3-3s3 1.3 3 3s-1.3 3-3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DataEncryption(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#D1C4E9" d="M38 7H10c-1.1 0-2 .9-2 2v6c0 1.1.9 2 2 2h28c1.1 0 2-.9 2-2V9c0-1.1-.9-2-2-2m0 12H10c-1.1 0-2 .9-2 2v6c0 1.1.9 2 2 2h25.1c1.3-1.3 4.9-.9 4.9-2v-6c0-1.1-.9-2-2-2m-3.6 12H10c-1.1 0-2 .9-2 2v6c0 1.1.9 2 2 2h28c1.1 0 2-.9 2-2v-2.4c0-3.1-2.5-5.6-5.6-5.6"/><g fill="#FFA000"><path d="m43 46l-2 2h-2l-2-2V35.4h6V40l-1 1l1 1v1l-1 1l1 1z"/><path d="M47.5 28.5c-.3-.9-1-1.6-2-1.8c-1.3-.3-3.3-.7-5.5-.7s-4.2.4-5.5.6c-1 .2-1.7.9-2 1.8c-.2 1-.5 2.2-.5 3.6s.3 2.6.5 3.5c.3.9 1 1.6 2 1.8c1.3.3 3.2.6 5.5.6s4.2-.4 5.5-.6c1-.2 1.7-.9 2-1.8c.3-.9.5-2.1.5-3.5s-.3-2.6-.5-3.5M42.9 31h-5.7c-.6 0-1.1-.5-1.1-1.1v-1.4c0-.3 1.8-.6 4-.6s4 .3 4 .6v1.4c-.1.6-.6 1.1-1.2 1.1"/></g><path fill="#D68600" d="M39 37.1h1V48h-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DataProtection(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#D1C4E9" d="M38 7H10c-1.1 0-2 .9-2 2v6c0 1.1.9 2 2 2h28c1.1 0 2-.9 2-2V9c0-1.1-.9-2-2-2m0 12H10c-1.1 0-2 .9-2 2v6c0 1.1.9 2 2 2h25.1c1.3-1.3 4.9-.9 4.9-2v-6c0-1.1-.9-2-2-2m-3.6 12H10c-1.1 0-2 .9-2 2v6c0 1.1.9 2 2 2h28c1.1 0 2-.9 2-2v-2.4c0-3.1-2.5-5.6-5.6-5.6"/><path fill="#009688" d="M46 25H32c-1.1 0-2 .9-2 2v11.8c0 1.3.6 2.4 1.6 3.2l7.4 5.5l7.4-5.5c1-.8 1.6-1.9 1.6-3.2V27c0-1.1-.9-2-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DataRecovery(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#D1C4E9" d="M38 7H10c-1.1 0-2 .9-2 2v6c0 1.1.9 2 2 2h28c1.1 0 2-.9 2-2V9c0-1.1-.9-2-2-2m0 12H10c-1.1 0-2 .9-2 2v6c0 1.1.9 2 2 2h28c1.1 0 2-.9 2-2v-6c0-1.1-.9-2-2-2m0 12H10c-1.1 0-2 .9-2 2v6c0 1.1.9 2 2 2h28c1.1 0 2-.9 2-2v-6c0-1.1-.9-2-2-2"/><g fill="#F44336"><path d="M35 28h6v20h-6z"/><path d="M28 35h20v6H28z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DataSheet(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#90CAF9" d="M39 16v7h-6v-7h-2v7h-6v-7h-2v7h-7v2h7v6h-7v2h7v6h-7v2h25V16zm0 9v6h-6v-6zm-14 0h6v6h-6zm0 8h6v6h-6zm8 6v-6h6v6z"/><path fill="#00BCD4" d="M40 8H8v32h8V16h24z"/><path fill="#0097A7" d="M7 7v34h10V17h24V7zm2 16v-6h6v6zm6 2v6H9v-6zm2-16h6v6h-6zm8 0h6v6h-6zM15 9v6H9V9zM9 39v-6h6v6zm30-24h-6V9h6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Database(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#D1C4E9" d="M38 7H10c-1.1 0-2 .9-2 2v6c0 1.1.9 2 2 2h28c1.1 0 2-.9 2-2V9c0-1.1-.9-2-2-2m0 12H10c-1.1 0-2 .9-2 2v6c0 1.1.9 2 2 2h28c1.1 0 2-.9 2-2v-6c0-1.1-.9-2-2-2m0 12H10c-1.1 0-2 .9-2 2v6c0 1.1.9 2 2 2h28c1.1 0 2-.9 2-2v-6c0-1.1-.9-2-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Debian(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#E91E63" d="M26.763 24.548c-.614.01.117.317.918.44a9.64 9.64 0 0 0 .6-.515c-.5.119-1.007.121-1.518.075m3.291-.821c.364-.5.631-1.055.723-1.624c-.082.405-.303.755-.51 1.128c-1.146.721-.108-.43 0-.865c-1.232 1.547-.169.927-.213 1.361m1.215-3.159c.073-1.105-.219-.756-.317-.336c.116.062.204.781.317.336m-6.83-15.09c.327.058.706.104.653.183c.357-.079.439-.151-.653-.183m.654.182l-.232.047l.215-.017zm10.201 15.326c.038.991-.29 1.472-.585 2.322l-.529.266c-.435.841.041.535-.268 1.202c-.679.603-2.055 1.883-2.496 2.004c-.321-.009.218-.382.289-.526c-.906.62-.728.934-2.113 1.313l-.041-.09c-3.419 1.607-8.166-1.576-8.103-5.928c-.037.275-.104.209-.18.32c-.175-2.237 1.033-4.486 3.073-5.403c1.995-.987 4.335-.58 5.763.75c-.785-1.028-2.348-2.119-4.199-2.017c-1.814.029-3.51 1.182-4.077 2.434c-.929.585-1.038 2.256-1.441 2.563c-.545 4.003 1.024 5.733 3.68 7.768c.417.282.118.326.175.541a7.173 7.173 0 0 1-2.354-1.801c.353.517.733 1.017 1.223 1.41c-.831-.279-1.942-2.013-2.267-2.084c1.435 2.567 5.818 4.502 8.113 3.541c-1.062.04-2.412.021-3.604-.42c-.501-.257-1.183-.791-1.062-.893c3.133 1.171 6.369.887 9.078-1.286c.689-.537 1.443-1.449 1.662-1.464c-.327.493.057.239-.197.674c.688-1.109-.299-.449.711-1.913l.373.512c-.139-.917 1.143-2.033 1.012-3.489c.291-.445.326.478.015 1.502c.434-1.136.113-1.317.224-2.254c.121.315.279.648.359.981c-.281-1.097.289-1.848.433-2.485c-.142-.063-.435.485-.503-.812c.01-.562.156-.295.214-.435c-.111-.064-.4-.496-.577-1.323c.127-.193.342.506.516.533c-.112-.655-.304-1.159-.313-1.665c-.51-1.061-.181.143-.592-.458c-.543-1.687.449-.39.514-1.156c.82 1.188 1.289 3.029 1.504 3.792a15.505 15.505 0 0 0-.752-2.704c.249.108-.401-1.911.324-.575c-.772-2.848-3.314-5.511-5.65-6.76c.286.262.646.591.517.642c-1.163-.69-.959-.745-1.124-1.041c-.946-.383-1.01.034-1.636 0c-1.786-.943-2.129-.845-3.772-1.437l.078.349c-1.184-.394-1.379.146-2.657.002c-.078-.062.41-.219.811-.278c-1.143.15-1.09-.228-2.208.042c.277-.197.566-.322.861-.486c-.932.059-2.226.542-1.825.103c-1.521.676-4.22 1.63-5.735 3.051l-.047-.322c-.694.835-3.028 2.492-3.215 3.57l-.185.043c-.361.613-.595 1.305-.881 1.935c-.474.806-.692.311-.626.436c-.929 1.883-1.39 3.467-1.79 4.768c.284.424.007 2.558.113 4.264c-.467 8.429 5.916 16.609 12.891 18.5c1.023.365 2.542.354 3.836.39c-1.525-.438-1.722-.232-3.209-.749c-1.074-.506-1.308-1.082-2.066-1.74l.3.53c-1.49-.526-.867-.652-2.078-1.034l.321-.424c-.482-.032-1.279-.811-1.497-1.241l-.528.021c-.634-.783-.972-1.348-.948-1.785l-.17.305c-.194-.332-2.335-2.937-1.224-2.33c-.207-.188-.481-.307-.779-.85l.227-.258c-.535-.686-.983-1.568-.949-1.86c.284.384.482.454.679.522c-1.351-3.349-1.426-.187-2.448-3.409l.216-.019c-.166-.246-.265-.521-.399-.785l.094-.938c-.972-1.125-.272-4.781-.132-6.783c.097-.816.811-1.684 1.354-3.045l-.332-.055c.632-1.104 3.612-4.433 4.99-4.26c.669-.841-.132-.002-.263-.215c1.469-1.52 1.93-1.073 2.92-1.349c1.068-.633-.917.251-.41-.239c1.848-.473 1.31-1.073 3.718-1.311c.254.145-.59.223-.8.41c1.538-.753 4.87-.584 7.034.417c2.511 1.173 5.33 4.642 5.443 7.904l.126.035c-.063 1.298.198 2.798-.257 4.175zm-15.222 4.403l-.086.431c.403.547.724 1.142 1.237 1.567c-.37-.723-.646-1.023-1.151-1.998m.951-.036c-.213-.237-.34-.518-.48-.802c.135.495.411.922.669 1.357zm16.854-3.665l-.088.226a11.004 11.004 0 0 1-1.068 3.412a10.73 10.73 0 0 0 1.156-3.638M24.56 5.185c.414-.154 1.019-.084 1.459-.185c-.573.048-1.144.079-1.706.151zm-14.553 7.738c.095.882-.667 1.229.167.644c.449-1.005-.174-.281-.167-.644m-.979 4.093c.191-.592.226-.943.3-1.285c-.531.679-.244.822-.3 1.285"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Debt(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#FFB74D" d="M10 12c-2.8 0-5-2.2-5-5s2.2-5 5-5s5 2.2 5 5s-2.2 5-5 5"/><path fill="#607D8B" d="M2 22v8l3 2l1 14h8l1-14l3-2v-8c0-4.4-3.6-8-8-8s-8 3.6-8 8"/><g fill="#263238"><path d="M22.4 40.4c-.6 2.5-1 3.6-2.4 3.6c-.6 0-1.2-.5-1.9-1.1c-1-.8-2.2-1.9-4.1-1.9v2c1.1 0 1.9.7 2.8 1.4c.9.7 1.9 1.6 3.2 1.6c3.1 0 3.8-2.9 4.4-5.2c.6-2.6 1-3.8 2.6-3.8v-2c-3.3 0-4.1 3.1-4.6 5.4"/><path d="M14.4 40H10v4h4.1z"/></g><circle cx="36" cy="36" r="10" fill="#4CAF50"/><path fill="#fff" d="M35 34c.1.2.1.3.3.4c.1.1.3.2.5.4c.2.1.5.2.8.3c.5.2.9.4 1.3.6c.4.2.7.4 1 .7c.3.3.5.6.7.9c.2.4.2.8.2 1.3c0 .4-.1.8-.2 1.2c-.1.4-.3.7-.6.9c-.3.3-.6.5-.9.6c-.4.2-.8.3-1.2.3v1.5h-1.2v-1.5c-.4 0-.8-.1-1.2-.3c-.4-.2-.7-.4-1-.6c-.3-.3-.5-.6-.7-1.1c-.2-.4-.3-.9-.3-1.5h2.2c0 .4 0 .7.1.9c.1.2.2.4.4.6c.2.1.3.2.5.3c.2.1.4.1.6.1c.2 0 .4 0 .6-.1c.2-.1.3-.2.4-.3c.1-.1.2-.3.3-.4c.1-.2.1-.3.1-.5s0-.4-.1-.6c-.1-.2-.1-.3-.3-.4c-.1-.1-.3-.3-.5-.4c-.2-.1-.4-.2-.7-.3c-.5-.2-.9-.4-1.3-.6c-.4-.2-.7-.4-1-.7c-.3-.3-.5-.6-.7-.9c-.2-.4-.2-.8-.2-1.3c0-.4.1-.8.2-1.2c.1-.3.3-.7.6-.9c.3-.3.6-.5.9-.6c.4-.2.8-.3 1.2-.3v-1.6H37v1.6c.4.1.8.2 1.2.4c.4.2.6.4.9.7c.2.3.4.6.6 1c.1.4.2.9.2 1.4h-2.2c0-.6-.1-1-.4-1.3c-.2-.3-.6-.4-1-.4c-.2 0-.4 0-.6.1c-.2.1-.3.2-.4.3c-.2 0-.3.1-.3.3s-.1.3-.1.5s0 .4.1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Decision(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#FFCC80"><circle cx="38" cy="26" r="4"/><circle cx="10" cy="26" r="4"/><path d="M39 19c0-12.7-30-8.3-30 0v10c0 8.3 6.7 15 15 15s15-6.7 15-15z"/><path d="M24 4C15.2 4 8 11.2 8 20v3.5l2.1.6V19l19.5-6.3l8.2 6.3v5.1l2.1-.6V20C40 12.5 34.6 4 24 4"/></g><path fill="#0277BD" d="M21.8 29.6c0-6.6 5.1-6.2 5.1-10.2c0-1-.3-3-2.9-3c-2.8 0-3 2.3-3 2.8h-3.8c0-1 .4-6 6.8-6c6.5 0 6.7 5.1 6.7 6c0 4.9-5.4 5.6-5.4 10.3h-3.5zm-.3 4.9c0-.3.1-2.1 2.1-2.1s2.2 1.8 2.2 2.1c0 .6-.3 2-2.2 2c-1.8 0-2.1-1.4-2.1-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeleteColumn(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#90CAF9" d="M30 5H18c-2.2 0-4 1.8-4 4v30c0 2.2 1.8 4 4 4h12c2.2 0 4-1.8 4-4V9c0-2.2-1.8-4-4-4M18 39V9h12v30z"/><circle cx="38" cy="38" r="10" fill="#F44336"/><g fill="#fff"><path d="m43.31 41.181l-2.12 2.122l-8.485-8.484l2.121-2.122z"/><path d="m34.819 43.31l-2.122-2.12l8.484-8.485l2.122 2.121z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeleteDatabase(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#D1C4E9" d="M38 7H10c-1.1 0-2 .9-2 2v6c0 1.1.9 2 2 2h28c1.1 0 2-.9 2-2V9c0-1.1-.9-2-2-2m0 12H10c-1.1 0-2 .9-2 2v6c0 1.1.9 2 2 2h28c1.1 0 2-.9 2-2v-6c0-1.1-.9-2-2-2m0 12H10c-1.1 0-2 .9-2 2v6c0 1.1.9 2 2 2h28c1.1 0 2-.9 2-2v-6c0-1.1-.9-2-2-2"/><circle cx="38" cy="38" r="10" fill="#F44336"/><g fill="#fff"><path d="m43.31 41.181l-2.12 2.122l-8.485-8.484l2.121-2.122z"/><path d="m34.819 43.31l-2.122-2.12l8.484-8.485l2.122 2.121z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeleteRow(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#90CAF9" d="M43 30V18c0-2.2-1.8-4-4-4H9c-2.2 0-4 1.8-4 4v12c0 2.2 1.8 4 4 4h30c2.2 0 4-1.8 4-4M9 18h30v12H9z"/><circle cx="38" cy="38" r="10" fill="#F44336"/><g fill="#fff"><path d="m43.31 41.181l-2.12 2.122l-8.485-8.484l2.121-2.122z"/><path d="m34.819 43.31l-2.122-2.12l8.484-8.485l2.122 2.121z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Department(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#C5CAE9" d="M42 42H6V9l18-7l18 7z"/><path fill="#9FA8DA" d="M6 42h36v2H6z"/><path fill="#BF360C" d="M20 35h8v9h-8z"/><path fill="#1565C0" d="M31 27h6v5h-6zm-10 0h6v5h-6zm-10 0h6v5h-6zm20 8h6v5h-6zm-20 0h6v5h-6zm20-16h6v5h-6zm-10 0h6v5h-6zm-10 0h6v5h-6zm20-8h6v5h-6zm-10 0h6v5h-6zm-10 0h6v5h-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Deployment(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#B0BEC5" d="M37 42H5V32h32c2.8 0 5 2.2 5 5s-2.2 5-5 5"/><path fill="#37474F" d="M10 34c-1.7 0-3 1.3-3 3s1.3 3 3 3s3-1.3 3-3s-1.3-3-3-3m0 4c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1s-.4 1-1 1m9-4c-1.7 0-3 1.3-3 3s1.3 3 3 3s3-1.3 3-3s-1.3-3-3-3m0 4c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1s-.4 1-1 1m18-4c-1.7 0-3 1.3-3 3s1.3 3 3 3s3-1.3 3-3s-1.3-3-3-3m0 4c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1s-.4 1-1 1m-9-4c-1.7 0-3 1.3-3 3s1.3 3 3 3s3-1.3 3-3s-1.3-3-3-3m0 4c-.6 0-1-.4-1-1s.4-1 1-1s1 .4 1 1s-.4 1-1 1"/><path fill="#FF9800" d="M35 31H11c-1.1 0-2-.9-2-2V7c0-1.1.9-2 2-2h24c1.1 0 2 .9 2 2v22c0 1.1-.9 2-2 2"/><path fill="#8A5100" d="M26.5 13h-7c-.8 0-1.5-.7-1.5-1.5s.7-1.5 1.5-1.5h7c.8 0 1.5.7 1.5 1.5s-.7 1.5-1.5 1.5"/><path fill="#607D8B" d="M37 31H5v2h32c2.2 0 4 1.8 4 4s-1.8 4-4 4H5v2h32c3.3 0 6-2.7 6-6s-2.7-6-6-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiplomaOne(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#E8EAF6" d="M4 9h40v30H4z"/><path fill="#5C6BC0" d="M30 34h2.8l-5-5l-2.8 2.8l5 5zm-12 0h-2.8l5-5l2.8 2.8l-5 5z"/><path fill="#9FA8DA" d="M11 15h26v4H11zm13 8c-2.8 0-5 2.2-5 5s2.2 5 5 5s5-2.2 5-5s-2.2-5-5-5m0 8c-1.7 0-3-1.3-3-3s1.3-3 3-3s3 1.3 3 3s-1.3 3-3 3"/><path fill="#9FA8DA" d="M3 8v32h42V8zm40 27c-1.7 0-3 1.3-3 3H8c0-1.7-1.3-3-3-3V13c1.7 0 3-1.3 3-3h32c0 1.7 1.3 3 3 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiplomaTwo(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#FBE9E7" d="M9 4h30v40H9z"/><path fill="#F4511E" d="M30 37h2.8l-5-5l-2.8 2.8l5 5zm-12 0h-2.8l5-5l2.8 2.8l-5 5z"/><path fill="#FF8A65" d="M15 13h18v4H15zm0 7h18v2H15zm9 6c-2.8 0-5 2.2-5 5s2.2 5 5 5s5-2.2 5-5s-2.2-5-5-5m0 8c-1.7 0-3-1.3-3-3s1.3-3 3-3s3 1.3 3 3s-1.3 3-3 3"/><path fill="#FF8A65" d="M8 3v42h32V3zm30 37c-1.7 0-3 1.3-3 3H13c0-1.7-1.3-3-3-3V8c1.7 0 3-1.3 3-3h22c0 1.7 1.3 3 3 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Disapprove(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#FFCC80"><circle cx="38" cy="26" r="4"/><circle cx="10" cy="26" r="4"/><path d="M39 19c0-12.7-30-8.3-30 0v10c0 8.3 6.7 15 15 15s15-6.7 15-15z"/><path d="M24 4C15.2 4 8 11.2 8 20v3.5l2.1.6V19l19.5-6.3l8.2 6.3v5.1l2.1-.6V20C40 12.5 34.6 4 24 4"/></g><g fill="#F44336"><path d="m16.077 20.91l2.828-2.828l13.08 13.08l-2.829 2.827z"/><path d="m29.09 18.077l2.828 2.828l-13.08 13.08l-2.827-2.829z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Disclaimer(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#FFCC80"><path d="M13 22H8v-8.5c0-1.4 1.1-2.5 2.5-2.5s2.5 1.1 2.5 2.5zm7 0h-5V7.5C15 6.1 16.1 5 17.5 5S20 6.1 20 7.5zm7 0h-5V5.5C22 4.1 23.1 3 24.5 3S27 4.1 27 5.5zm7 0h-5V8.5C29 7.1 30.1 6 31.5 6S34 7.1 34 8.5zm-1.9 21l-5-5l10-10c1.4-1.4 3.6-1.4 4.9 0c1.4 1.4 1.4 3.6 0 4.9z"/><path d="M29 21c0 .6-.4 1-1 1s-1-.4-1-1h-5c0 .6-.4 1-1 1s-1-.4-1-1h-5c0 .6-.4 1-1 1s-1-.4-1-1H8v16c0 4.4 3.6 8 8 8h11.2c3.7 0 6.8-3 6.8-6.8V21z"/></g><g fill="#F44336"><path d="m15.413 28.971l2.474-2.474l10.605 10.605l-2.474 2.474z"/><path d="m25.993 26.504l2.475 2.474l-10.605 10.605l-2.475-2.474z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dislike(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#F44336" d="M34 9c-4.2 0-7.9 2.1-10 5.4C21.9 11.1 18.2 9 14 9C7.4 9 2 14.4 2 21c0 11.9 22 24 22 24s22-12 22-24c0-6.6-5.4-12-12-12"/><path fill="#37474F" d="M3.563 6.396L6.39 3.568l37.966 37.966l-2.828 2.828z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Display(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#80DEEA" d="M40 41H8c-2.2 0-4-1.8-4-4V11c0-2.2 1.8-4 4-4h32c2.2 0 4 1.8 4 4v26c0 2.2-1.8 4-4 4"/><path fill="#2962FF" d="M36 17h-5l-2-2l2-2h5l2 2zm0 18h-5l-2-2l2-2h5l2 2zm1-5V18l2-2l2 2v12l-2 2zm-11 0V18l2-2l2 2v12l-2 2zm-9-13h-5l-2-2l2-2h5l2 2zm0 18h-5l-2-2l2-2h5l2 2zm1-5V18l2-2l2 2v12l-2 2zM7 30V18l2-2l2 2v12l-2 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoNotInhale(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<defs><path id="flatColorIconsDoNotInhale0" fill="#FFA726" d="M33.5 20C31.2 17.7 30 13.9 30 9V6H18v3c0 4.9-1.2 8.7-3.5 11c-2.4.2-5.5 2-5.5 5.4c0 4.5 5.1 4.6 6 4.6c1.2 0 6.1 4 8 4h2c1.9 0 6.8-4 8-4c.9 0 6-.1 6-4.6c0-3.4-3.1-5.2-5.5-5.4"/></defs><use href="#flatColorIconsDoNotInhale0"/><use href="#flatColorIconsDoNotInhale0"/><path fill="#FFB74D" d="M26 9V6h-4v3c0 4.9-3 19-3 19s1.6 2 5 2s5-2 5-2s-3-14.1-3-19"/><path fill="#CC861E" d="M23 34c-3.3 0-6.4-3.1-8-4h1.8c3.1 0 4.5 4 6.2 4m2 0c3.3 0 6.4-3.1 8-4h-1.8c-3.1 0-4.5 4-6.2 4"/><path fill="#37474F" d="M3.563 6.396L6.39 3.568l37.966 37.966l-2.828 2.828z"/><g fill="#FF5722"><path d="m18 35l4 4h-8z"/><path d="M17 38h2v4h-2zm13-3l4 4h-8z"/><path d="M29 38h2v4h-2z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoNotInsert(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#7CB342"><path d="m31 29l-7 7l-7-7z"/><path d="M22 7h4v25h-4z"/><path d="M42 18c-3.3 0-6 2.7-6 6v12c0 1.1-.9 2-2 2H14c-1.1 0-2-.9-2-2V24c0-3.3-2.7-6-6-6H4v4h2c1.1 0 2 .9 2 2v12c0 3.3 2.7 6 6 6h20c3.3 0 6-2.7 6-6V24c0-1.1.9-2 2-2h2v-4z"/></g><path fill="#37474F" d="M3.563 6.396L6.39 3.568l37.966 37.966l-2.828 2.828z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoNotMix(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#00BCD4" d="M26.9 42H17v-9.9z"/><path fill="#00BCD4" d="M30 6v20.2L19.8 36.4l2.8 2.8L34 27.8V6z"/><path fill="#2196F3" d="M15.9 31H6v-9.9z"/><path fill="#2196F3" d="M20.2 14L8.8 25.4l2.8 2.8L21.8 18H41v-4z"/><path fill="#37474F" d="M3.563 6.396L6.39 3.568l37.966 37.966l-2.828 2.828z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Document(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#90CAF9" d="M40 45H8V3h22l10 10z"/><path fill="#E1F5FE" d="M38.5 14H29V4.5z"/><path fill="#1976D2" d="M16 21h17v2H16zm0 4h13v2H16zm0 4h17v2H16zm0 4h13v2H16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Donate(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#E69329" d="m11.7 21.6l5.1 9.9l9.5-3.9l4.4-12.7l-14.8.8z"/><circle cx="15" cy="36" r="7.8" fill="#546E7A"/><g fill="#90A4AE"><path d="M15 27c-5 0-9 4-9 9s4 9 9 9s9-4 9-9s-4-9-9-9m0 16c-3.9 0-7-3.1-7-7s3.1-7 7-7s7 3.1 7 7s-3.1 7-7 7"/><path d="M14 33h2v8h-2z"/></g><g fill="#FFB74D"><path d="M12.9 36c1 1.9 3.2 2.7 5.1 1.7l16.5-8.5c1-.5 1.7-1.2 2.2-1.9c1.7-3.2 5.6-10.7 8.2-17.2l-18.2 8.7l-4.8 7.2l-6.8 3.6c-2.6 1.3-3.4 4.2-2.2 6.4"/><path d="M30.2 3L13.7 9.3c-.7.2-1.5 1-2.2 1.7l-5.6 7.5c-1 1.5-1.2 3.4-.5 5.1c.4 1 1.7 3.4 3.1 6.1C10.1 28 12.4 27 15 27c.4 0 .9 0 1.3.1l-2.1-4.2l4.6-4.1h8s15.5-2.2 18.2-8.7z"/></g><path fill="#FFCDD2" d="M18.2 36c-1.3.6-2.8 0-3.3-1.3c-.6-1.3 0-2.8 1.3-3.3c1.2-.6 3.2 4 2 4.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoughnutChart(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#00BCD4" d="M24 30c-3.3 0-6-2.7-6-6s2.7-6 6-6V5C13.5 5 5 13.5 5 24s8.5 19 19 19c4.4 0 8.5-1.5 11.8-4.1l-8-10.2c-1.1.8-2.4 1.3-3.8 1.3"/><path fill="#448AFF" d="M30 24h13c0-10.5-8.5-19-19-19v13c3.3 0 6 2.7 6 6"/><path fill="#3F51B5" d="M43 24H30c0 1.9-.9 3.6-2.3 4.7l8 10.2C40.2 35.4 43 30 43 24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Down(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#3F51B5"><path d="M24 44L12.3 30h23.4z"/><path d="M20 6h8v27h-8z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownLeft(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#3F51B5" d="m4 29l14-11.7v23.4z"/><path fill="#3F51B5" d="M42 21V8h-8v13c0 2.2-1.8 4-4 4H13v8h17c6.6 0 12-5.4 12-12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownRight(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#3F51B5" d="M44 29L30 17.3v23.4z"/><path fill="#3F51B5" d="M6 21V8h8v13c0 2.2 1.8 4 4 4h17v8H18c-6.6 0-12-5.4-12-12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Download(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#1565C0"><path d="M24 37.1L13 24h22zM20 4h8v4h-8zm0 6h8v4h-8z"/><path d="M20 16h8v11h-8zM6 40h36v4H6z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dribbble(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#FF4081" d="M42 37a5 5 0 0 1-5 5H11a5 5 0 0 1-5-5V11a5 5 0 0 1 5-5h26a5 5 0 0 1 5 5z"/><path fill="#FFF" d="M33.061 26.273a1.347 1.347 0 0 0-1.686.895c-.824 2.658-2.316 5.419-2.993 5.57c-.507 0-1.236-.43-1.958-1.44c1.674-3.594 2.551-8.106 2.551-11.118c0-8.413-2.124-10.18-3.908-10.18c-3.757 0-3.8 9.912-3.8 10.012c0 1.166.042 2.248.121 3.256a5.149 5.149 0 0 0-1.77-.319c-3.86 0-5.618 3.809-5.618 7.347C14 33.63 15.871 37 20.046 37c1.972 0 3.634-1.291 4.975-3.221c1.188 1.235 2.432 1.696 3.36 1.696c2.923 0 4.858-5.233 5.556-7.486a1.375 1.375 0 0 0-.876-1.716m-13.013 7.991c-3.031 0-3.36-2.775-3.36-3.969c0-.188.034-4.611 2.932-4.611c1.144 0 2.022.885 2.022.885c.065.07.137.131.212.184c.375 1.904.904 3.426 1.516 4.632c-1.004 1.738-2.167 2.879-3.322 2.879m4.853-6.338c-.559-1.93-.946-4.521-.946-7.914c0-3.126.666-6.068 1.219-7.073c.424.644 1.115 2.65 1.115 7.241c0 2.436-.539 5.266-1.388 7.746"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DvdLogo(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#42A5F5" d="M24.002 27c-12.154 0-22 1.343-22 3.006c0 1.653 9.845 2.994 22 2.994c12.156 0 22-1.341 22-2.994c0-1.663-9.844-3.006-22-3.006m0 3.972c-2.863 0-5.191-.494-5.191-1.104c0-.609 2.329-1.104 5.191-1.104c2.862 0 5.193.495 5.193 1.104c0 .61-2.331 1.104-5.193 1.104"/><path fill="#1565C0" d="m21.29 15l2.371 6.43L29.25 15h9.486c4.647 0 7.906 2.148 7.158 4.904c-.745 2.756-5.178 4.904-9.803 4.904h-6.295s.141-.043.172-.126c.246-.944 1.707-6.264 1.725-6.347c.02-.102-.105-.133-.105-.133h4.572s-.088-.006-.125.133a757.71 757.71 0 0 0-1.145 4.176c-.023.094-.162.139-.162.139h1.094c2.594 0 5.047-.828 5.563-2.748c.473-1.752-1.244-2.746-4.039-2.746h-1.014l-4.375.004l-10.043 9.932l-3.845-9.754s-.036-.066-.065-.147c-.014-.026-.108-.106-.206-.063c-.065.036-.074.117-.066.146c.036.066.042.08.048.111c.569.93.467 2.009.33 2.52c-.774 2.75-5.186 4.904-9.812 4.904H2.002s.149-.043.172-.126c.254-.946 1.717-6.294 1.726-6.347c.018-.09-.099-.133-.099-.133h4.604s-.132.037-.158.131c-.024.078-.954 3.432-1.151 4.178c-.023.094-.178.139-.178.139h1.118c2.597 0 5.032-.828 5.547-2.748c.472-1.752-1.23-2.746-4.021-2.746H4.089s.125-.059.147-.139c.123-.443.497-1.834.515-1.899c.02-.072-.105-.119-.105-.119z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditImage(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#8CBCD6" d="M31 41H8c-2.2 0-4-1.8-4-4V11c0-2.2 1.8-4 4-4h32c2.2 0 4 1.8 4 4v17c0 7.2-5.8 13-13 13"/><circle cx="35" cy="16" r="3" fill="#B3DDF5"/><path fill="#9AC9E3" d="M20 16L9 32h22z"/><path fill="#B3DDF5" d="m31 22l-8 10h16z"/><path fill="#E57373" d="m47.7 29.1l-2.8-2.8c-.4-.4-1.1-.4-1.6 0L42 27.6l4.4 4.4l1.3-1.3c.4-.4.4-1.1 0-1.6"/><path fill="#FF9800" d="M27.467 42.167L39.77 29.865l4.384 4.384L31.85 46.55z"/><path fill="#B0BEC5" d="m46.4 32.038l-2.192 2.192l-4.383-4.384l2.192-2.191z"/><path fill="#FFC107" d="M27.5 42.2L26 48l5.8-1.5z"/><path fill="#37474F" d="m26.7 45l-.7 3l3-.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ElectricalSensor(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="32" cy="24" r="9" fill="#B2EBF2"/><path fill="#4DD0E1" d="M32 12c-6.6 0-12 5.4-12 12s5.4 12 12 12s12-5.4 12-12s-5.4-12-12-12m0 20c-4.4 0-8-3.6-8-8s3.6-8 8-8s8 3.6 8 8s-3.6 8-8 8"/><g fill="#3F51B5"><path d="M25.4 22L19.8 5.1l-6.2 22.6l-2.2-5.7H4v4h4.6l5.8 14.3l5.8-21.4l2.4 7.1H30v-4z"/><circle cx="32" cy="24" r="4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ElectricalThreshold(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#80DEEA" d="M3 12h42v24H3z"/><path fill="#03A9F4" d="M3 23h42v2H3z"/><path fill="none" stroke="#3F51B5" stroke-miterlimit="10" stroke-width="4" d="m4 18l4.5-1.5c.9-.3 1.9.1 2.3.9l8.7 14.3c.7 1.2 2.4 1.3 3.2.2l2.3-2.8c.5-.6 1.4-.9 2.2-.6l3 1c1 .3 2.1-.2 2.5-1.1L37 18.3c.5-1.1 1.9-1.6 2.9-.9l4 2.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Electricity(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#00BCD4" d="M33.7 5L22 17l15 5l-15.7 14.7l5.1 2.8L12 43l2.7-14.8l2.9 5.1L27 24l-15-5L25 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ElectroDevices(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#3F51B5" d="M39 43H9c-2.2 0-4-1.8-4-4V9c0-2.2 1.8-4 4-4h30c2.2 0 4 1.8 4 4v30c0 2.2-1.8 4-4 4"/><path fill="#80DEEA" d="m33.2 5l-9.8 10.1L36 19.3L22.8 31.7l4.3 2.4L15 37l2.3-12.5l2.4 4.3l8-7.8L15 16.8L25.9 5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Electronics(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#FF9800" d="M44 18v-4H34V4h-4v10h-4V4h-4v10h-4V4h-4v10H4v4h10v4H4v4h10v4H4v4h10v10h4V34h4v10h4V34h4v10h4V34h10v-4H34v-4h10v-4H34v-4z"/><path fill="#4CAF50" d="M8 12v24c0 2.2 1.8 4 4 4h24c2.2 0 4-1.8 4-4V12c0-2.2-1.8-4-4-4H12c-2.2 0-4 1.8-4 4"/><path fill="#37474F" d="M31 31H17c-1.1 0-2-.9-2-2V19c0-1.1.9-2 2-2h14c1.1 0 2 .9 2 2v10c0 1.1-.9 2-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmptyBattery(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#CFD8DC"><path d="M34 44H14c-1.1 0-2-.9-2-2V8c0-1.1.9-2 2-2h20c1.1 0 2 .9 2 2v34c0 1.1-.9 2-2 2"/><path d="M28 13h-8c-.6 0-1-.4-1-1V5c0-.6.4-1 1-1h8c.6 0 1 .4 1 1v7c0 .6-.4 1-1 1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmptyFilter(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#FFCC80" d="M29 23H19L7 9h34zm0 15l-10 6V23h10zM41.5 9h-35C5.7 9 5 8.3 5 7.5S5.7 6 6.5 6h35c.8 0 1.5.7 1.5 1.5S42.3 9 41.5 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmptyTrash(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#B39DDB" d="M30.6 44H17.4c-2 0-3.7-1.4-4-3.4L9 11h30l-4.5 29.6c-.3 2-2 3.4-3.9 3.4"/><path fill="#7E57C2" d="M38 13H10c-1.1 0-2-.9-2-2s.9-2 2-2h28c1.1 0 2 .9 2 2s-.9 2-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EndCall(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#F44336" d="m43.5 16.8l-2.3-2.3c-8.1-7.9-27.5-6.8-34.5 0l-2.3 2.3c-.6.6-.6 1.6 0 2.3L9 23.6c.6.6 1.7.6 2.3 0l5.1-4.9l-.4-5.3c1.6-1.6 14.4-1.6 16 0l-.3 5.5l4.9 4.7c.6.6 1.7.6 2.3 0l4.6-4.5c.7-.7.7-1.7 0-2.3"/><g fill="#B71C1C"><path d="M24 40.5L16 31h16z"/><path d="M21 24h6v7.5h-6z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Engineering(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#EF6C00" d="m37.4 24.6l-11.6-2.2l-3.9-11.2l-3.8 1.3L22 23.6l-7.8 9l3 2.6l7.8-9l11.6 2.2z"/><g fill="#FF9800"><path d="M24 19c-2.8 0-5 2.2-5 5s2.2 5 5 5s5-2.2 5-5s-2.2-5-5-5m0 7c-1.1 0-2-.9-2-2s.9-2 2-2s2 .9 2 2s-.9 2-2 2"/><path d="M40.7 27c.2-1 .3-2 .3-3s-.1-2-.3-3l3.3-2.4c.4-.3.6-.9.3-1.4L40 9.8c-.3-.5-.8-.7-1.3-.4L35 11c-1.5-1.3-3.3-2.3-5.2-3l-.4-4.1c-.1-.5-.5-.9-1-.9h-8.6c-.5 0-1 .4-1 .9L18.2 8c-1.9.7-3.7 1.7-5.2 3L9.3 9.3c-.5-.2-1.1 0-1.3.5l-4.3 7.4c-.3.5-.1 1.1.3 1.4L7.3 21c-.2 1-.3 2-.3 3s.1 2 .3 3L4 29.4c-.4.3-.6.9-.3 1.4L8 38.2c.3.5.8.7 1.3.4L13 37c1.5 1.3 3.3 2.3 5.2 3l.4 4.1c.1.5.5.9 1 .9h8.6c.5 0 1-.4 1-.9l.4-4.1c1.9-.7 3.7-1.7 5.2-3l3.7 1.7c.5.2 1.1 0 1.3-.4l4.3-7.4c.3-.5.1-1.1-.3-1.4zM24 35c-6.1 0-11-4.9-11-11s4.9-11 11-11s11 4.9 11 11s-4.9 11-11 11"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnteringHeavenAlive(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#039BE5" d="M17 29h14v2H17zm-4 4h22v2H13zm-4 4h30v2H9zm-4 4h38v2H5z"/><path fill="#81D4FA" d="M35 13c-.4 0-.8 0-1.2.1C32.9 8.5 28.9 5 24 5c-4.1 0-7.6 2.5-9.2 6H14c-4.4 0-8 3.6-8 8s3.6 8 8 8h21c3.9 0 7-3.1 7-7s-3.1-7-7-7"/><path fill="#039BE5" d="M28 21c0-2.2-1.8-4-4-4s-4 1.8-4 4v6h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Expand(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#2196F3" d="M43 17.1L39.9 14L24 29.9L8.1 14L5 17.1L24 36z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Expired(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="17" cy="17" r="14" fill="#00ACC1"/><circle cx="17" cy="17" r="11" fill="#eee"/><path d="M16 8h2v9h-2z"/><path d="m22.655 20.954l-1.697 1.697l-4.808-4.807l1.697-1.697z"/><circle cx="17" cy="17" r="2"/><circle cx="17" cy="17" r="1" fill="#00ACC1"/><path fill="#FFC107" d="m11.9 42l14.4-24.1c.8-1.3 2.7-1.3 3.4 0L44.1 42c.8 1.3-.2 3-1.7 3H13.6c-1.5 0-2.5-1.7-1.7-3"/><path fill="#263238" d="M26.4 39.9c0-.2 0-.4.1-.6s.2-.3.3-.5s.3-.2.5-.3s.4-.1.6-.1s.5 0 .7.1s.4.2.5.3s.2.3.3.5s.1.4.1.6s0 .4-.1.6s-.2.3-.3.5s-.3.2-.5.3s-.4.1-.7.1s-.5 0-.6-.1s-.4-.2-.5-.3s-.2-.3-.3-.5s-.1-.4-.1-.6m2.8-3.1h-2.3l-.4-9.8h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Export(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#FFCCBC" d="M7 40V8c0-2.2 1.8-4 4-4h24c2.2 0 4 1.8 4 4v32c0 2.2-1.8 4-4 4H11c-2.2 0-4-1.8-4-4"/><g fill="#FF5722"><path d="M42.7 24L32 33V15z"/><path d="M14 21h23v6H14z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func External(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="24" cy="31" r="14" fill="#B2DFDB"/><g fill="#009688"><path d="M24 3.3L33 14H15z"/><path d="M21 11h6v23h-6z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Factory(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#BF360C" d="M41.2 5h-7.3L32 43h11z"/><path fill="#E64A19" d="M33 23h-4v-6l-12 6v-6L5 23v20h28z"/><path fill="#FFC107" d="M9 27h4v4H9zm8 0h4v4h-4zm8 0h4v4h-4zM9 35h4v4H9zm8 0h4v4h-4zm8 0h4v4h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FactoryBreakdown(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#E64A19" d="M29 23v-6l-8 4v2h-4v4h-4v-4H5v20h28V23z"/><path fill="#992B0A" d="M25 27h4v4h-4zM9 35h4v4H9zm16 0h4v4h-4zm-8 0h4v4h-4zm0-8h4v4h-4z"/><path fill="#BF360C" d="M41.2 5H38v2h-2v2h-2.3L32 43h11z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Faq(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#558B2F" d="M15 40h23l6 6V25c0-2.2-1.8-4-4-4H15c-2.2 0-4 1.8-4 4v11c0 2.2 1.8 4 4 4"/><path fill="#1B5E20" d="M28.8 32.8h-3.6l-.7 2.1h-2.2l3.7-10h1.9l3.7 10h-2.2zm-3.1-1.6h2.5L27 27.4z"/><path fill="#8BC34A" d="M33 25H10l-6 6V8c0-2.2 1.8-4 4-4h25c2.2 0 4 1.8 4 4v13c0 2.2-1.8 4-4 4"/><path fill="#fff" d="M25.4 14.2c0 1-.2 1.8-.5 2.5s-.7 1.3-1.3 1.7l1.7 1.3l-1.3 1.2l-2.2-1.7c-.2 0-.5.1-.8.1c-.6 0-1.2-.1-1.8-.3c-.5-.2-1-.6-1.4-1c-.4-.4-.7-1-.9-1.6c-.2-.6-.3-1.3-.3-2.1v-.4c0-.8.1-1.5.3-2.1c.2-.6.5-1.2.9-1.6c.4-.4.8-.8 1.4-1c.5-.2 1.1-.3 1.8-.3c.6 0 1.2.1 1.8.3c.5.2 1 .6 1.4 1c.4.4.7 1 .9 1.6c.2.6.3 1.3.3 2.1zm-2.2-.5c0-1.1-.2-1.9-.6-2.4c-.4-.6-.9-.8-1.6-.8c-.7 0-1.3.3-1.6.8c-.4.6-.6 1.4-.6 2.4v.5c0 .5.1 1 .2 1.4c.1.4.2.8.4 1c.2.3.4.5.7.6c.3.1.6.2.9.2c.7 0 1.3-.3 1.6-.8c.4-.6.6-1.4.6-2.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FeedIn(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#2196F3" d="M38 24v12c0 1.1-.9 2-2 2H12c-1.1 0-2-.9-2-2V24c0-3.3-2.7-6-6-6v4c1.1 0 2 .9 2 2v12c0 3.3 2.7 6 6 6h24c3.3 0 6-2.7 6-6V24c0-1.1.9-2 2-2v-4c-3.3 0-6 2.7-6 6"/><g fill="#3F51B5"><path d="M38.6 5.6L29 15.2V28h4V16.8l8.4-8.4zm-32 2.8l8.4 8.4V28h4V15.2L9.4 5.6z"/><path d="m37 27l-6 6l-6-6zm-14 0l-6 6l-6-6z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Feedback(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#78909C" d="M40 41H8c-2.2 0-4-1.8-4-4V16.1c0-1.3.6-2.5 1.7-3.3L24 0l18.3 12.8c1.1.7 1.7 2 1.7 3.3V37c0 2.2-1.8 4-4 4"/><path fill="#fff" d="M12 11h24v22H12z"/><path fill="#9C27B0" d="m24 13.6l-6 7.8h12z"/><path fill="#CFD8DC" d="M40 41H8c-2.2 0-4-1.8-4-4V17l20 13l20-13v20c0 2.2-1.8 4-4 4"/><path fill="#9C27B0" d="m24 28l2-1.3V20h-4v6.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func File(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#90CAF9" d="M40 45H8V3h22l10 10z"/><path fill="#E1F5FE" d="M38.5 14H29V4.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilingCabinet(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#263238" d="M12 44h4v2h-4zm20 0h4v2h-4z"/><path fill="#607D8B" d="M8 41V7c0-2.2 1.8-4 4-4h24c2.2 0 4 1.8 4 4v34c0 2.2-1.8 4-4 4H12c-2.2 0-4-1.8-4-4"/><path fill="#B0BEC5" d="M12 17V8c0-.6.4-1 1-1h22c.6 0 1 .4 1 1v9zm0 2h24v10H12zm0 21v-9h24v9c0 .6-.4 1-1 1H13c-.6 0-1-.4-1-1"/><path fill="#546E7A" d="M20 11h8v2h-8zm0 12h8v2h-8zm0 12h8v2h-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilledFilter(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#F57C00" d="M29 23H19L7 9h34z"/><path fill="#FF9800" d="m29 38l-10 6V23h10zM41.5 9h-35C5.7 9 5 8.3 5 7.5S5.7 6 6.5 6h35c.8 0 1.5.7 1.5 1.5S42.3 9 41.5 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Film(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#3F51B5" d="M45 9H3v30h42zM22 37v-4h4v4zm8 0v-4h4v4zm8 0v-4h4v4zm-24 0v-4h4v4zm-8 0v-4h4v4zm16-22v-4h4v4zm8 0v-4h4v4zm8 0v-4h4v4zm-24 0v-4h4v4zm-8 0v-4h4v4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilmReel(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#3F51B5" d="M43 39V24h-4v15c0 5 4 9 9 9v-4c-2.8 0-5-2.2-5-5"/><circle cx="24" cy="24" r="19" fill="#90A4AE"/><circle cx="24" cy="24" r="2" fill="#37474F"/><g fill="#253278"><circle cx="24" cy="14" r="5"/><circle cx="24" cy="34" r="5"/><circle cx="34" cy="24" r="5"/><circle cx="14" cy="24" r="5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FinePrint(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#90CAF9" d="M33 42H5V4h19l9 9z"/><path fill="#E1F5FE" d="M31.5 14H23V5.5z"/><path fill="#616161" d="m34.505 37.58l1.98-1.98l8.483 8.485l-1.98 1.98z"/><circle cx="28" cy="29" r="11" fill="#616161"/><circle cx="28" cy="29" r="9" fill="#90CAF9"/><path fill="#37474F" d="m36.849 39.88l1.98-1.98l6.15 6.151l-1.98 1.98z"/><path fill="#1976D2" d="M30 31h-9.7c.4 1.6 1.3 3 2.5 4H30zm-9.7-4H30v-4h-7.3c-1.2 1-2 2.4-2.4 4m-.2-7H11v2h7.3c.5-.7 1.1-1.4 1.8-2m-3 4H11v2h5.4c.2-.7.4-1.4.7-2M16 29c0-.3 0-.7.1-1H11v2h5.1c-.1-.3-.1-.7-.1-1m.4 3H11v2h6.1c-.3-.6-.5-1.3-.7-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlashAuto(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#FFC107" d="M33 22h-9.4L30 5H19l-6 21h8.6L17 45z"/><path fill="#F44336" d="M40.8 14.5h-4.3l-.9 2.5H33l4.5-12h2.3l4.5 12h-2.6zm-3.7-2h3L38.6 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlashOff(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#FFC107" d="M33 22h-9.4L30 5H19l-6 21h8.6L17 45z"/><path fill="#37474F" d="M3.563 6.396L6.39 3.568l37.966 37.966l-2.828 2.828z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlashOn(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#FFC107" d="M33 22h-9.4L30 5H19l-6 21h8.6L17 45z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlowChart(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#CFD8DC" d="M35 36h4V22H26v-9h-4v9H9v14h4V26h9v10h4V26h9z"/><path fill="#3F51B5" d="M17 6h14v10H17z"/><path fill="#00BCD4" d="M32 32h10v10H32zM6 32h10v10H6zm13 0h10v10H19z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folder(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#FFA000" d="M40 12H22l-4-4H8c-2.2 0-4 1.8-4 4v8h40v-4c0-2.2-1.8-4-4-4"/><path fill="#FFCA28" d="M40 12H8c-2.2 0-4 1.8-4 4v20c0 2.2 1.8 4 4 4h32c2.2 0 4-1.8 4-4V16c0-2.2-1.8-4-4-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Frame(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#3949AB" d="M40.6 40.1h-1.9l-3.1-.4c-2.4-.3-4.9-.2-7.3.4l-3.6.9c-.5.1-1.1.1-1.6 0l-3.5-1c-2.4-.6-4.8-.7-7.3-.4l-3.1.4H7.4C5.5 40 4 38.5 4 36.6c0-.4.1-.9.2-1.3l.2-.6c1-2.5 1.1-5.3.4-7.9l-.6-2c-.2-.7-.2-1.3 0-2l.3-.8c.9-2.7.8-5.7-.2-8.4l-.1-.3c-.1-.2-.2-.6-.2-1v-1c0-1.9 1.5-3.4 3.4-3.4h1.9l3.1.4c2.4.3 4.9.2 7.3-.4l3.6-.9c.5-.1 1.1-.1 1.6 0l3.5 1c2.4.6 4.8.7 7.3.4l3.1-.4h1.9c1.9 0 3.4 1.5 3.4 3.4v1c0 .4-.1.9-.2 1.3l-.1.3c-1.1 2.7-1.2 5.6-.2 8.4l.3.8c.2.6.2 1.3 0 2l-.6 2c-.7 2.6-.6 5.4.4 7.9l.2.6c.2.4.2.8.2 1.3c-.1 1.6-1.6 3.1-3.5 3.1"/><path fill="#BBDEFB" d="M38 36H10c-.6 0-1-.4-1-1V13c0-.6.4-1 1-1h28c.6 0 1 .4 1 1v22c0 .6-.4 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FullBattery(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#8BC34A"><path d="M34 44H14c-1.1 0-2-.9-2-2V8c0-1.1.9-2 2-2h20c1.1 0 2 .9 2 2v34c0 1.1-.9 2-2 2"/><path d="M28 13h-8c-.6 0-1-.4-1-1V5c0-.6.4-1 1-1h8c.6 0 1 .4 1 1v7c0 .6-.4 1-1 1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FullTrash(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#FF8A65" d="M24 21.3L12.7 10L26 1.7L38.3 10z"/><path fill="#FFAB91" d="M24 21.3L12.7 10L17 4.7L38.3 10z"/><path fill="#B39DDB" d="M30.6 44H17.4c-2 0-3.7-1.4-4-3.4L9 11h30l-4.5 29.6c-.3 2-2 3.4-3.9 3.4"/><path fill="#7E57C2" d="M38 13H10c-1.1 0-2-.9-2-2s.9-2 2-2h28c1.1 0 2 .9 2 2s-.9 2-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gallery(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#E65100" d="M41 42H13c-2.2 0-4-1.8-4-4V18c0-2.2 1.8-4 4-4h28c2.2 0 4 1.8 4 4v20c0 2.2-1.8 4-4 4"/><path fill="#F57C00" d="M35 36H7c-2.2 0-4-1.8-4-4V12c0-2.2 1.8-4 4-4h28c2.2 0 4 1.8 4 4v20c0 2.2-1.8 4-4 4"/><circle cx="30" cy="16" r="3" fill="#FFF9C4"/><path fill="#942A09" d="M17 17.9L8 31h18z"/><path fill="#BF360C" d="M28 23.5L22 31h12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Genealogy(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#CFD8DC" d="M40 9V7h-9v5H15v11H8v2h7v11h16v5h9v-2h-7v-8h7v-2h-9v5H17V14h14v5h9v-2h-7V9z"/><path fill="#00BCD4" d="M4 20h8v8H4z"/><path fill="#3F51B5" d="M36 14h8v8h-8zm0-10h8v8h-8zM20 9h8v8h-8zm0 22h8v8h-8zm16 5h8v8h-8zm0-10h8v8h-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GenericSortingAsc(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#2196F3" d="M6 6h4v4H6zm0 8h12v4H6zm0 8h20v4H6zm0 8h28v4H6zm0 8h36v4H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GenericSortingDesc(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#2196F3" d="M6 38h4v4H6zm0-8h12v4H6zm0-8h20v4H6zm0-8h28v4H6zm0-8h36v4H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Globe(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#7CB342" d="M24 4C13 4 4 13 4 24s9 20 20 20s20-9 20-20S35 4 24 4"/><path fill="#0277BD" d="M45 24c0 11.7-9.5 21-21 21S3 35.7 3 24S12.3 3 24 3s21 9.3 21 21m-21.2 9.7c0-.4-.2-.6-.6-.8c-1.3-.4-2.5-.4-3.6-1.5c-.2-.4-.2-.8-.4-1.3c-.4-.4-1.5-.6-2.1-.8h-4.2c-.6-.2-1.1-1.1-1.5-1.7c0-.2 0-.6-.4-.6c-.4-.2-.8.2-1.3 0c-.2-.2-.2-.4-.2-.6c0-.6.4-1.3.8-1.7c.6-.4 1.3.2 1.9.2c.2 0 .2 0 .4.2c.6.2.8 1 .8 1.7v.4c0 .2.2.2.4.2c.2-1.1.2-2.1.4-3.2c0-1.3 1.3-2.5 2.3-2.9c.4-.2.6.2 1.1 0c1.3-.4 4.4-1.7 3.8-3.4c-.4-1.5-1.7-2.9-3.4-2.7c-.4.2-.6.4-1 .6c-.6.4-1.9 1.7-2.5 1.7c-1.1-.2-1.1-1.7-.8-2.3c.2-.8 2.1-3.6 3.4-3.1l.8.8c.4.2 1.1.2 1.7.2c.2 0 .4 0 .6-.2c.2-.2.2-.2.2-.4c0-.6-.6-1.3-1-1.7c-.4-.4-1.1-.8-1.7-1.1c-2.1-.6-5.5.2-7.1 1.7s-2.9 4-3.8 6.1c-.4 1.3-.8 2.9-1 4.4c-.2 1-.4 1.9.2 2.9c.6 1.3 1.9 2.5 3.2 3.4c.8.6 2.5.6 3.4 1.7c.6.8.4 1.9.4 2.9c0 1.3.8 2.3 1.3 3.4c.2.6.4 1.5.6 2.1c0 .2.2 1.5.2 1.7c1.3.6 2.3 1.3 3.8 1.7c.2 0 1-1.3 1-1.5c.6-.6 1.1-1.5 1.7-1.9c.4-.2.8-.4 1.3-.8c.4-.4.6-1.3.8-1.9c.1-.5.3-1.3.1-1.9m.4-19.4c.2 0 .4-.2.8-.4c.6-.4 1.3-1.1 1.9-1.5c.6-.4 1.3-1.1 1.7-1.5c.6-.4 1.1-1.3 1.3-1.9c.2-.4.8-1.3.6-1.9c-.2-.4-1.3-.6-1.7-.8c-1.7-.4-3.1-.6-4.8-.6c-.6 0-1.5.2-1.7.8c-.2 1.1.6.8 1.5 1.1c0 0 .2 1.7.2 1.9c.2 1-.4 1.7-.4 2.7c0 .6 0 1.7.4 2.1zM41.8 29c.2-.4.2-1.1.4-1.5c.2-1 .2-2.1.2-3.1c0-2.1-.2-4.2-.8-6.1c-.4-.6-.6-1.3-.8-1.9c-.4-1.1-1-2.1-1.9-2.9c-.8-1.1-1.9-4-3.8-3.1c-.6.2-1 1-1.5 1.5c-.4.6-.8 1.3-1.3 1.9c-.2.2-.4.6-.2.8c0 .2.2.2.4.2c.4.2.6.2 1 .4c.2 0 .4.2.2.4c0 0 0 .2-.2.2c-1 1.1-2.1 1.9-3.1 2.9c-.2.2-.4.6-.4.8c0 .2.2.2.2.4s-.2.2-.4.4c-.4.2-.8.4-1.1.6c-.2.4 0 1.1-.2 1.5c-.2 1.1-.8 1.9-1.3 2.9c-.4.6-.6 1.3-1 1.9c0 .8-.2 1.5.2 2.1c1 1.5 2.9.6 4.4 1.3c.4.2.8.2 1.1.6c.6.6.6 1.7.8 2.3c.2.8.4 1.7.8 2.5c.2 1 .6 2.1.8 2.9c1.9-1.5 3.6-3.1 4.8-5.2c1.5-1.3 2.1-3 2.7-4.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoodDecision(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#FFCC80"><circle cx="38" cy="26" r="4"/><circle cx="10" cy="26" r="4"/><path d="M39 19c0-12.7-30-8.3-30 0v10c0 8.3 6.7 15 15 15s15-6.7 15-15z"/><path d="M24 4C15.2 4 8 11.2 8 20v3.5l2.1.6V19l19.5-6.3l8.2 6.3v5.1l2.1-.6V20C40 12.5 34.6 4 24 4"/></g><g fill="#4CAF50"><path d="M22 16h4v18h-4z"/><path d="M15 23h18v4H15z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Google(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#FFC107" d="M43.611 20.083H42V20H24v8h11.303c-1.649 4.657-6.08 8-11.303 8c-6.627 0-12-5.373-12-12s5.373-12 12-12c3.059 0 5.842 1.154 7.961 3.039l5.657-5.657C34.046 6.053 29.268 4 24 4C12.955 4 4 12.955 4 24s8.955 20 20 20s20-8.955 20-20c0-1.341-.138-2.65-.389-3.917"/><path fill="#FF3D00" d="m6.306 14.691l6.571 4.819C14.655 15.108 18.961 12 24 12c3.059 0 5.842 1.154 7.961 3.039l5.657-5.657C34.046 6.053 29.268 4 24 4C16.318 4 9.656 8.337 6.306 14.691"/><path fill="#4CAF50" d="M24 44c5.166 0 9.86-1.977 13.409-5.192l-6.19-5.238A11.91 11.91 0 0 1 24 36c-5.202 0-9.619-3.317-11.283-7.946l-6.522 5.025C9.505 39.556 16.227 44 24 44"/><path fill="#1976D2" d="M43.611 20.083H42V20H24v8h11.303a12.04 12.04 0 0 1-4.087 5.571l.003-.002l6.19 5.238C36.971 39.205 44 34 44 24c0-1.341-.138-2.65-.389-3.917"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GraduationCap(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#37474F"><path d="M9 20h30v13H9z"/><ellipse cx="24" cy="33" rx="15" ry="6"/></g><path fill="#78909C" d="M23.1 8.2L.6 18.1c-.8.4-.8 1.5 0 1.9l22.5 9.9c.6.2 1.2.2 1.8 0L47.4 20c.8-.4.8-1.5 0-1.9L24.9 8.2c-.6-.3-1.2-.3-1.8 0"/><g fill="#37474F"><path d="m43.2 20.4l-20-3.4c-.5-.1-1.1.3-1.2.8c-.1.5.3 1.1.8 1.2L42 22.2V37c0 .6.4 1 1 1s1-.4 1-1V21.4c0-.5-.4-.9-.8-1"/><circle cx="43" cy="37" r="2"/><path d="M46 40c0 1.7-3 6-3 6s-3-4.3-3-6s1.3-3 3-3s3 1.3 3 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Grid(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#90CAF9" d="M7 7v34h34V7zm32 8h-6V9h6zm-14 0V9h6v6zm6 2v6h-6v-6zm-8-2h-6V9h6zm0 2v6h-6v-6zm-8 6H9v-6h6zm0 2v6H9v-6zm2 0h6v6h-6zm6 8v6h-6v-6zm2 0h6v6h-6zm0-2v-6h6v6zm8-6h6v6h-6zm0-2v-6h6v6zM15 9v6H9V9zM9 33h6v6H9zm24 6v-6h6v6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Headset(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#0097A7" d="M24 5C14.1 5 6 13.1 6 23v15h4V23c0-7.7 6.3-14 14-14s14 6.3 14 14v15h4V23c0-9.9-8.1-18-18-18"/><path fill="#37474F" d="M38 43h-4V31h4c2.2 0 4 1.8 4 4v4c0 2.2-1.8 4-4 4m-28 0h4V31h-4c-2.2 0-4 1.8-4 4v4c0 2.2 1.8 4 4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeatMap(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#CFD8DC" d="M9 39V6H7v35h35v-2z"/><g fill="#00BCD4"><circle cx="14" cy="11" r="2"/><circle cx="32" cy="11" r="2"/><circle cx="39" cy="11" r="2"/><circle cx="23" cy="11" r="4"/><circle cx="14" cy="33" r="2"/><circle cx="30" cy="33" r="2"/><circle cx="22" cy="33" r="3"/><circle cx="38" cy="33" r="4"/><circle cx="14" cy="22" r="2"/><circle cx="39" cy="22" r="2"/><circle cx="32" cy="22" r="3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HighBattery(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#CFD8DC"><path d="M34 44H14c-1.1 0-2-.9-2-2V8c0-1.1.9-2 2-2h20c1.1 0 2 .9 2 2v34c0 1.1-.9 2-2 2"/><path d="M28 13h-8c-.6 0-1-.4-1-1V5c0-.6.4-1 1-1h8c.6 0 1 .4 1 1v7c0 .6-.4 1-1 1"/></g><path fill="#8BC34A" d="M34 44H14c-1.1 0-2-.9-2-2V13h24v29c0 1.1-.9 2-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HighPriority(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#F44336" d="m21.2 44.8l-18-18c-1.6-1.6-1.6-4.1 0-5.7l18-18c1.6-1.6 4.1-1.6 5.7 0l18 18c1.6 1.6 1.6 4.1 0 5.7l-18 18c-1.6 1.6-4.2 1.6-5.7 0"/><path fill="#fff" d="M21.6 32.7c0-.3.1-.6.2-.9c.1-.3.3-.5.5-.7c.2-.2.5-.4.8-.5s.6-.2 1-.2s.7.1 1 .2c.3.1.6.3.8.5c.2.2.4.4.5.7c.1.3.2.6.2.9s-.1.6-.2.9s-.3.5-.5.7c-.2.2-.5.4-.8.5c-.3.1-.6.2-1 .2s-.7-.1-1-.2s-.5-.3-.8-.5c-.2-.2-.4-.4-.5-.7s-.2-.5-.2-.9m4.2-4.6h-3.6L21.7 13h4.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Home(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#E8EAF6" d="M42 39H6V23L24 6l18 17z"/><path fill="#C5CAE9" d="m39 21l-5-5V9h5zM6 39h36v5H6z"/><path fill="#B71C1C" d="M24 4.3L4 22.9l2 2.2L24 8.4l18 16.7l2-2.2z"/><path fill="#D84315" d="M18 28h12v16H18z"/><path fill="#01579B" d="M21 17h6v6h-6z"/><path fill="#FF8A65" d="M27.5 35.5c-.3 0-.5.2-.5.5v2c0 .3.2.5.5.5s.5-.2.5-.5v-2c0-.3-.2-.5-.5-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IconsEightCup(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#4CAF50" d="M40 14H8l3.8 28.3c.1 1 1 1.7 2 1.7h20.5c1 0 1.8-.7 2-1.7z"/><g fill="#81C784"><path d="M42 14H6v-3c0-2.2 1.8-4 4-4h28c2.2 0 4 1.8 4 4z"/><path d="M37.2 10H10.8l1.7-4.7c.3-.8 1-1.3 1.9-1.3h19.2c.8 0 1.6.5 1.9 1.3z"/></g><path fill="#E8F5E9" d="M28 28.5c1.2-1.1 2-2.7 2-4.5c0-3.3-2.7-6-6-6s-6 2.7-6 6c0 1.8.8 3.4 2 4.5c-1.2 1.1-2 2.7-2 4.5c0 3.3 2.7 6 6 6s6-2.7 6-6c0-1.8-.8-3.4-2-4.5M24 36c-1.7 0-3-1.3-3-3s1.3-3 3-3s3 1.3 3 3s-1.3 3-3 3m0-9c-1.7 0-3-1.3-3-3s1.3-3 3-3s3 1.3 3 3s-1.3 3-3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Idea(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="24" cy="22" r="20" fill="#FFF59D"/><path fill="#FBC02D" d="M37 22c0-7.7-6.6-13.8-14.5-12.9c-6 .7-10.8 5.5-11.4 11.5c-.5 4.6 1.4 8.7 4.6 11.3c1.4 1.2 2.3 2.9 2.3 4.8v.3h12v-.1c0-1.8.8-3.6 2.2-4.8c2.9-2.4 4.8-6 4.8-10.1"/><path fill="#FFF59D" d="m30.6 20.2l-3-2c-.3-.2-.8-.2-1.1 0L24 19.8l-2.4-1.6c-.3-.2-.8-.2-1.1 0l-3 2c-.2.2-.4.4-.4.7s0 .6.2.8l3.8 4.7V37h2V26c0-.2-.1-.4-.2-.6l-3.3-4.1l1.5-1l2.4 1.6c.3.2.8.2 1.1 0l2.4-1.6l1.5 1l-3.3 4.1c-.1.2-.2.4-.2.6v11h2V26.4l3.8-4.7c.2-.2.3-.5.2-.8s-.2-.6-.4-.7"/><circle cx="24" cy="44" r="3" fill="#5C6BC0"/><path fill="#9FA8DA" d="M26 45h-4c-2.2 0-4-1.8-4-4v-5h12v5c0 2.2-1.8 4-4 4"/><path fill="#5C6BC0" d="m30 41l-11.6 1.6c.3.7.9 1.4 1.6 1.8l9.4-1.3c.4-.6.6-1.3.6-2.1m-12-2.3v2L30 39v-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageFile(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#90CAF9" d="M40 45H8V3h22l10 10z"/><path fill="#E1F5FE" d="M38.5 14H29V4.5z"/><path fill="#1565C0" d="m21 23l-7 10h14z"/><path fill="#1976D2" d="M28 26.4L23 33h10z"/><circle cx="31.5" cy="24.5" r="1.5" fill="#1976D2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Import(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#F8BBD0" d="M7 40V8c0-2.2 1.8-4 4-4h24c2.2 0 4 1.8 4 4v32c0 2.2-1.8 4-4 4H11c-2.2 0-4-1.8-4-4"/><g fill="#E91E63"><path d="M13.3 24L24 15v18z"/><path d="M19 21h23v6H19z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InTransit(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#FFC107" d="M44 36H30V16c0-1.1.9-2 2-2h8c.6 0 1.2.3 1.6.8l6 7.7c.3.4.4.8.4 1.2V32c0 2.2-1.8 4-4 4"/><g fill="#9575CD"><path d="M8 36h22V13c0-2.2-1.8-4-4-4H4v23c0 2.2 1.8 4 4 4"/><path d="M0 9h10v2H0zm0 5h10v2H0zm0 5h10v2H0zm0 5h10v2H0z"/></g><path fill="#7E57C2" d="M4 11h16v2H4zm0 5h12v2H4zm0 5h8v2H4zm0 5h4v2H4z"/><g fill="#37474F"><circle cx="39" cy="36" r="5"/><circle cx="16" cy="36" r="5"/></g><g fill="#78909C"><circle cx="39" cy="36" r="2.5"/><circle cx="16" cy="36" r="2.5"/></g><path fill="#455A64" d="M44 26h-3.6c-.3 0-.5-.1-.7-.3l-1.4-1.4c-.2-.2-.4-.3-.7-.3H34c-.6 0-1-.4-1-1v-6c0-.6.4-1 1-1h5.5c.3 0 .6.1.8.4l4.5 5.4c.1.2.2.4.2.6V25c0 .6-.4 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Info(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="24" cy="24" r="21" fill="#2196F3"/><path fill="#fff" d="M22 22h4v11h-4z"/><circle cx="24" cy="16.5" r="2.5" fill="#fff"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Inspection(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#455A64" d="M36 4H26c0 1.1-.9 2-2 2s-2-.9-2-2H12C9.8 4 8 5.8 8 8v32c0 2.2 1.8 4 4 4h24c2.2 0 4-1.8 4-4V8c0-2.2-1.8-4-4-4"/><path fill="#fff" d="M36 41H12c-.6 0-1-.4-1-1V8c0-.6.4-1 1-1h24c.6 0 1 .4 1 1v32c0 .6-.4 1-1 1"/><g fill="#90A4AE"><path d="M26 4c0 1.1-.9 2-2 2s-2-.9-2-2h-7v4c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2V4z"/><path d="M24 0c-2.2 0-4 1.8-4 4s1.8 4 4 4s4-1.8 4-4s-1.8-4-4-4m0 6c-1.1 0-2-.9-2-2s.9-2 2-2s2 .9 2 2s-.9 2-2 2"/></g><path fill="#4CAF50" d="m30.6 18.6l-9 9l-4.2-4.3l-2.5 2.5l6.8 6.7l11.4-11.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IntegratedWebcam(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#607D8B" d="M38 42H10c-2.2 0-4-1.8-4-4V10c0-2.2 1.8-4 4-4h28c2.2 0 4 1.8 4 4v28c0 2.2-1.8 4-4 4"/><circle cx="24" cy="24" r="12" fill="#455A64"/><circle cx="24" cy="24" r="9" fill="#42A5F5"/><path fill="#90CAF9" d="M28.8 21c-1.2-1.4-3-2.2-4.8-2.2s-3.6.8-4.8 2.2c-.5.5-.4 1.3.1 1.8s1.3.4 1.8-.1c1.5-1.7 4.3-1.7 5.8 0c.3.3.6.4 1 .4c.3 0 .6-.1.9-.3c.4-.4.5-1.3 0-1.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Internal(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="24" cy="30" r="15" fill="#B3E5FC"/><g fill="#1565C0"><path d="M24 38.7L15 28h18z"/><path d="M21 5h6v26h-6z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Invite(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#78909C" d="M40 41H8c-2.2 0-4-1.8-4-4V16.1c0-1.3.6-2.5 1.7-3.3L24 0l18.3 12.8c1.1.7 1.7 2 1.7 3.3V37c0 2.2-1.8 4-4 4"/><path fill="#fff" d="M12 11h24v22H12z"/><path fill="#CFD8DC" d="M40 41H8c-2.2 0-4-1.8-4-4V17l20 13l20-13v20c0 2.2-1.8 4-4 4"/><g fill="#4CAF50"><path d="M22 14h4v12h-4z"/><path d="M18 18h12v4H18z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ipad(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#E38939" d="M8 41V7c0-2.2 1.8-4 4-4h24c2.2 0 4 1.8 4 4v34c0 2.2-1.8 4-4 4H12c-2.2 0-4-1.8-4-4"/><path fill="#FFF3E0" d="M36 6H12c-.6 0-1 .4-1 1v31c0 .6.4 1 1 1h24c.6 0 1-.4 1-1V7c0-.6-.4-1-1-1"/><circle cx="24" cy="42" r="1.5" fill="#A6642A"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Iphone(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#E38939" d="M12 40V8c0-2.2 1.8-4 4-4h16c2.2 0 4 1.8 4 4v32c0 2.2-1.8 4-4 4H16c-2.2 0-4-1.8-4-4"/><path fill="#FFF3E0" d="M32 7H16c-.6 0-1 .4-1 1v29c0 .6.4 1 1 1h16c.6 0 1-.4 1-1V8c0-.6-.4-1-1-1"/><circle cx="24" cy="41" r="1.5" fill="#A6642A"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Key(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#FFA000"><path d="m30 41l-4 4h-4l-4-4V21h12v8l-2 2l2 2v2l-2 2l2 2z"/><path d="M38 7.8c-.5-1.8-2-3.1-3.7-3.6C31.9 3.7 28.2 3 24 3s-7.9.7-10.3 1.2C12 4.7 10.5 6 10 7.8c-.5 1.7-1 4.1-1 6.7c0 2.6.5 5 1 6.7c.5 1.8 1.9 3.1 3.7 3.5c2.4.6 6.1 1.3 10.3 1.3s7.9-.7 10.3-1.2c1.8-.4 3.2-1.8 3.7-3.5s1-4.1 1-6.7c0-2.7-.5-5.1-1-6.8M29 13H19c-1.1 0-2-.9-2-2V9c0-.6 3.1-1 7-1s7 .4 7 1v2c0 1.1-.9 2-2 2"/></g><path fill="#D68600" d="M23 26h2v19h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Kindle(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#37474F" d="M8 41V7c0-2.2 1.8-4 4-4h24c2.2 0 4 1.8 4 4v34c0 2.2-1.8 4-4 4H12c-2.2 0-4-1.8-4-4"/><path fill="#eee" d="M35 6H13c-.6 0-1 .4-1 1v29c0 .6.4 1 1 1h22c.6 0 1-.4 1-1V7c0-.6-.4-1-1-1"/><path fill="#546E7A" d="M20 40h8v2h-8z"/><path fill="#A1A1A1" d="M16 11h16v3H16zm0 7h16v2H16zm0 4h12v2H16zm0 4h16v2H16zm0 4h12v2H16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Landscape(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#FF9800"><path d="m40.997 6.065l7 7l-7 6.999l-7-7z"/><path d="M36 8h10v10H36z"/></g><circle cx="41" cy="13" r="3" fill="#FFEB3B"/><path fill="#2E7D32" d="M16.5 18L0 42h33z"/><path fill="#4CAF50" d="M33.6 24L19.2 42H48z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Leave(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#FFCDD2" d="M5 38V14h38v24c0 2.2-1.8 4-4 4H9c-2.2 0-4-1.8-4-4"/><path fill="#F44336" d="M43 10v6H5v-6c0-2.2 1.8-4 4-4h30c2.2 0 4 1.8 4 4"/><g fill="#B71C1C"><circle cx="33" cy="10" r="3"/><circle cx="15" cy="10" r="3"/></g><path fill="#BDBDBD" d="M33 3c-1.1 0-2 .9-2 2v5c0 1.1.9 2 2 2s2-.9 2-2V5c0-1.1-.9-2-2-2M15 3c-1.1 0-2 .9-2 2v5c0 1.1.9 2 2 2s2-.9 2-2V5c0-1.1-.9-2-2-2"/><path fill="#F44336" d="M22.2 35.3c0-.2 0-.5.1-.7c.1-.2.2-.4.4-.5s.3-.3.5-.3c.2-.1.5-.1.7-.1s.5 0 .7.1l.6.3c.2.1.3.3.4.5c.1.2.1.4.1.7c0 .2 0 .5-.1.7c-.1.2-.2.4-.4.5c-.2.1-.3.3-.6.3s-.3.2-.6.2s-.5 0-.7-.1c-.2-.1-.4-.2-.5-.3c-.2-.1-.3-.3-.4-.5c-.1-.3-.2-.5-.2-.8m3.1-4.3h-2.6l-.4-11h3.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Left(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#3F51B5"><path d="m4 24l14-11.7v23.4z"/><path d="M15 20h27v8H15z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftDown(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#3F51B5" d="M7 41V23l18 18z"/><path fill="#3F51B5" d="m35.355 6.946l5.656 5.656l-23.119 23.12l-5.656-5.657z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftDownTwo(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#3F51B5" d="m19 44l11.7-14H7.3z"/><path fill="#3F51B5" d="M27 6h13v8H27c-2.2 0-4 1.8-4 4v17h-8V18c0-6.6 5.4-12 12-12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftUp(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#3F51B5" d="M7 7h18L7 25z"/><path fill="#3F51B5" d="m41.02 35.322l-5.656 5.656l-23.12-23.119l5.657-5.656z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftUpTwo(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#3F51B5" d="m19 4l11.7 14H7.3z"/><path fill="#3F51B5" d="M27 42h13v-8H27c-2.2 0-4-1.8-4-4V13h-8v17c0 6.6 5.4 12 12 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Library(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#FF9800" d="M1 38h46v2H1zm24-20h4v16h-4zm6 0h4v16h-4zm6 0h4v16h-4zm-18 0h4v16h-4zm-6 0h4v16h-4zm-6 0h4v16H7zm36-2H5v-3l19-9l19 9zM5 34h38v2H5z"/><g fill="#EF6C00"><path d="M25 16h4v2h-4zm6 0h4v2h-4zm6 0h4v2h-4zm-18 0h4v2h-4zm-6 0h4v2h-4zm-6 0h4v2H7zM3 36h42v2H3z"/><circle cx="24" cy="11" r="2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightAtTheEndOfTunnel(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#3F51B5" d="M6 10v28c0 2.2 1.8 4 4 4h28c2.2 0 4-1.8 4-4V10c0-2.2-1.8-4-4-4H10c-2.2 0-4 1.8-4 4"/><path fill="#CCF2F6" d="M27.9 28.9h-5.8l-8.4 7.2l6-7.2v-2.4l-3 .8l3-1.9v-1.5c0-.8.1-1.7.6-2.4l-7.5-8.3l8.7 7.2c.7-.7 1.5-1.1 2.5-1.2l.6-7.3l1.1 7.3c.3 0 .6.1.8.1l1.2-1.2l-.3 1.7c.3.1.4.3.7.6l4.4-2.8l-3.6 3.9c.3.4.6 1 .7 1.7l2.2.1l-2.2.8v1.5l2.6 1.4l-2.6-.3v2.2l6.2 7.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Like(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#F44336" d="M34 9c-4.2 0-7.9 2.1-10 5.4C21.9 11.1 18.2 9 14 9C7.4 9 2 14.4 2 21c0 11.9 22 24 22 24s22-12 22-24c0-6.6-5.4-12-12-12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LikePlaceholder(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#FFCDD2" d="M34 9c-4.2 0-7.9 2.1-10 5.4C21.9 11.1 18.2 9 14 9C7.4 9 2 14.4 2 21c0 11.9 22 24 22 24s22-12 22-24c0-6.6-5.4-12-12-12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LineChart(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#3F51B5"><circle cx="8" cy="38" r="3"/><circle cx="16" cy="40" r="3"/><circle cx="24" cy="33" r="3"/><circle cx="32" cy="35" r="3"/><circle cx="40" cy="31" r="3"/><path d="m39.1 29.2l-7.3 3.7l-8.3-2.1l-8 7l-7-1.7l-1 3.8l9 2.3l8-7l7.7 1.9l8.7-4.3z"/></g><g fill="#00BCD4"><circle cx="8" cy="20" r="3"/><circle cx="16" cy="22" r="3"/><circle cx="24" cy="15" r="3"/><circle cx="32" cy="20" r="3"/><circle cx="40" cy="8" r="3"/><path d="M38.3 6.9c-2.1 3.2-5.3 8-6.9 10.4c-1.2-.7-3.1-2-6.4-4l-1.3-.8l-8.3 7.3l-7-1.7l-1 3.9l9 2.3l7.7-6.7c2.6 1.6 5.8 3.6 6.5 4.1l.5.5l.9-.1c1.1-.1 1.1-.1 9.5-12.9z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Link(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#1976D2" d="M38 13h-3c-5.5 0-10 4.5-10 10s4.5 10 10 10h3c5.5 0 10-4.5 10-10s-4.5-10-10-10m0 16h-3c-3.3 0-6-2.7-6-6s2.7-6 6-6h3c3.3 0 6 2.7 6 6s-2.7 6-6 6M13 13h-3C4.5 13 0 17.5 0 23s4.5 10 10 10h3c5.5 0 10-4.5 10-10s-4.5-10-10-10m0 16h-3c-3.3 0-6-2.7-6-6s2.7-6 6-6h3c3.3 0 6 2.7 6 6s-2.7 6-6 6"/><path fill="#42A5F5" d="M33 21H15c-1.1 0-2 .9-2 2s.9 2 2 2h18c1.1 0 2-.9 2-2s-.9-2-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Linux(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#ECEFF1" d="m20.1 16.2l.1 2.3l-1.6 3l-2.5 4.9l-.5 4.1l1.8 5.8l4.1 2.3h6.2l5.8-4.4l2.6-6.9l-6-7.3l-1.7-4.1z"/><path fill="#263238" d="M34.3 21.9c-1.6-2.3-2.9-3.7-3.6-6.6c-.7-2.9.2-2.1-.4-4.6c-.3-1.3-.8-2.2-1.3-2.9c-.6-.7-1.3-1.1-1.7-1.2c-.9-.5-3-1.3-5.6.1c-2.7 1.4-2.4 4.4-1.9 10.5c0 .4-.1.9-.3 1.3c-.4.9-1.1 1.7-1.7 2.4c-.7 1-1.4 2-1.9 3.1c-1.2 2.3-2.3 5.2-2 6.3c.5-.1 6.8 9.5 6.8 9.7c.4-.1 2.1-.1 3.6-.1c2.1-.1 3.3-.2 5 .2c0-.3-.1-.6-.1-.9c0-.6.1-1.1.2-1.8c.1-.5.2-1 .3-1.6c-1 .9-2.8 1.9-4.5 2.2c-1.5.3-4-.2-5.2-1.7c.1 0 .3 0 .4-.1c.3-.1.6-.2.7-.4c.3-.5.1-1-.1-1.3c-.2-.3-1.7-1.4-2.4-2c-.7-.6-1.1-.9-1.5-1.3l-.8-.8c-.2-.2-.3-.4-.4-.5c-.2-.5-.3-1.1-.2-1.9c.1-1.1.5-2 1-3c.2-.4.7-1.2.7-1.2s-1.7 4.2-.8 5.5c0 0 .1-1.3.5-2.6c.3-.9.8-2.2 1.4-2.9s2.1-3.3 2.2-4.9c0-.7.1-1.4.1-1.9c-.4-.4 6.6-1.4 7-.3c.1.4 1.5 4 2.3 5.9c.4.9.9 1.7 1.2 2.7c.3 1.1.5 2.6.5 4.1c0 .3 0 .8-.1 1.3c.2 0 4.1-4.2-.5-7.7c0 0 2.8 1.3 2.9 3.9c.1 2.1-.8 3.8-1 4.1c.1 0 2.1.9 2.2.9c.4 0 1.2-.3 1.2-.3c.1-.3.4-1.1.4-1.4c.7-2.3-1-6-2.6-8.3"/><g fill="#ECEFF1" transform="translate(0 -2)"><ellipse cx="21.6" cy="15.3" rx="1.3" ry="2"/><ellipse cx="26.1" cy="15.2" rx="1.7" ry="2.3"/></g><g fill="#212121" transform="translate(0 -2)"><ellipse cx="21.7" cy="15.5" rx="1.2" ry=".7" transform="rotate(-97.204 21.677 15.542)"/><ellipse cx="26" cy="15.6" rx="1" ry="1.3"/></g><path fill="#FFC107" d="M39.3 35.6c-.4-.2-1.1-.5-1.7-1.4c-.3-.5-.2-1.9-.7-2.5c-.3-.4-.7-.2-.8-.2c-.9.2-3 1.6-4.4 0c-.2-.2-.5-.5-1-.5s-.7.2-.9.6s-.2.7-.2 1.7c0 .8 0 1.7-.1 2.4c-.2 1.7-.5 2.7-.5 3.7c0 1.1.3 1.8.7 2.1c.3.3.8.5 1.9.5c1.1 0 1.8-.4 2.5-1.1c.5-.5.9-.7 2.3-1.7c1.1-.7 2.8-1.6 3.1-1.9c.2-.2.5-.3.5-.9c0-.5-.4-.7-.7-.8m-20.1.3c-1-1.6-1.1-1.9-1.8-2.9c-.6-1-1.9-2.9-2.7-2.9c-.6 0-.9.3-1.3.7c-.4.4-.8 1.3-1.5 1.8c-.6.5-2.3.4-2.7 1c-.4.6.4 1.5.4 3c0 .6-.5 1-.6 1.4c-.1.5-.2.8 0 1.2c.4.6.9.8 4.3 1.5c1.8.4 3.5 1.4 4.6 1.5c1.1.1 3 0 3-2.7c.1-1.6-.8-2-1.7-3.6m1.9-18.1c-.6-.4-1.1-.8-1.1-1.4c0-.6.4-.8 1-1.3c.1-.1 1.2-1.1 2.3-1.1s2.4.7 2.9.9c.9.2 1.8.4 1.7 1.1c-.1 1-.2 1.2-1.2 1.7c-.7.2-2 1.3-2.9 1.3c-.4 0-1 0-1.4-.1c-.3-.1-.8-.6-1.3-1.1"/><path fill="#634703" d="M20.9 17c.2.2.5.4.8.5c.2.1.5.2.5.2h.9c.5 0 1.2-.2 1.9-.6c.7-.3.8-.5 1.3-.7c.5-.3 1-.6.8-.7c-.2-.1-.4 0-1.1.4c-.6.4-1.1.6-1.7.9c-.3.1-.7.3-1 .3h-.9c-.3 0-.5-.1-.8-.2c-.2-.1-.3-.2-.4-.2c-.2-.1-.6-.5-.8-.6c0 0-.2 0-.1.1zm3-2.2c.1.2.3.2.4.3c.1.1.2.1.2.1c.1-.1 0-.3-.1-.3c0-.2-.5-.2-.5-.1m-1.6.2c0 .1.2.2.2.1c.1-.1.2-.2.3-.2c.2-.1.1-.2-.2-.2c-.2.1-.2.2-.3.3"/><path fill="#455A64" d="M32 32.7v.3c.2.4.7.5 1.1.5c.6 0 1.2-.4 1.5-.8c0-.1.1-.2.2-.3c.2-.3.3-.5.4-.6c0 0-.1-.1-.1-.2c-.1-.2-.4-.4-.8-.5c-.3-.1-.8-.2-1-.2c-.9-.1-1.4.2-1.7.5c0 0 .1 0 .1.1c.2.2.3.4.3.7c.1.2 0 .3 0 .5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func List(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#2196F3" d="M6 22h4v4H6zm0-8h4v4H6zm0 16h4v4H6zM6 6h4v4H6zm0 32h4v4H6zm8-16h28v4H14zm0-8h28v4H14zm0 16h28v4H14zm0-24h28v4H14zm0 32h28v4H14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lock(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#424242" d="M24 4c-5.5 0-10 4.5-10 10v4h4v-4c0-3.3 2.7-6 6-6s6 2.7 6 6v4h4v-4c0-5.5-4.5-10-10-10"/><path fill="#FB8C00" d="M36 44H12c-2.2 0-4-1.8-4-4V22c0-2.2 1.8-4 4-4h24c2.2 0 4 1.8 4 4v18c0 2.2-1.8 4-4 4"/><circle cx="24" cy="31" r="3" fill="#C76E00"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockLandscape(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#37474F" d="M7 10h34c2.2 0 4 1.8 4 4v20c0 2.2-1.8 4-4 4H7c-2.2 0-4-1.8-4-4V14c0-2.2 1.8-4 4-4"/><path fill="#BBDEFB" d="M42 34V14c0-.6-.4-1-1-1H7c-.6 0-1 .4-1 1v20c0 .6.4 1 1 1h34c.6 0 1-.4 1-1"/><g fill="#3F51B5"><path d="M29 31H19c-.6 0-1-.4-1-1v-6c0-.6.4-1 1-1h10c.6 0 1 .4 1 1v6c0 .6-.4 1-1 1"/><path d="M24 17c-2.2 0-4 1.8-4 4v3h2v-3c0-1.1.9-2 2-2s2 .9 2 2v3h2v-3c0-2.2-1.8-4-4-4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockPortrait(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#37474F" d="M10 41V7c0-2.2 1.8-4 4-4h20c2.2 0 4 1.8 4 4v34c0 2.2-1.8 4-4 4H14c-2.2 0-4-1.8-4-4"/><path fill="#BBDEFB" d="M34 6H14c-.6 0-1 .4-1 1v34c0 .6.4 1 1 1h20c.6 0 1-.4 1-1V7c0-.6-.4-1-1-1"/><g fill="#3F51B5"><path d="M29 30H19c-.6 0-1-.4-1-1v-6c0-.6.4-1 1-1h10c.6 0 1 .4 1 1v6c0 .6-.4 1-1 1"/><path d="M24 16c-2.2 0-4 1.8-4 4v3h2v-3c0-1.1.9-2 2-2s2 .9 2 2v3h2v-3c0-2.2-1.8-4-4-4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LowBattery(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#CFD8DC"><path d="M34 44H14c-1.1 0-2-.9-2-2V8c0-1.1.9-2 2-2h20c1.1 0 2 .9 2 2v34c0 1.1-.9 2-2 2"/><path d="M28 13h-8c-.6 0-1-.4-1-1V5c0-.6.4-1 1-1h8c.6 0 1 .4 1 1v7c0 .6-.4 1-1 1"/></g><path fill="#8BC34A" d="M34 44H14c-1.1 0-2-.9-2-2v-9h24v9c0 1.1-.9 2-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LowPriority(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#4CAF50" d="m21.2 44.8l-18-18c-1.6-1.6-1.6-4.1 0-5.7l18-18c1.6-1.6 4.1-1.6 5.7 0l18 18c1.6 1.6 1.6 4.1 0 5.7l-18 18c-1.6 1.6-4.2 1.6-5.7 0"/><g fill="#FFEB3B"><path d="M24 33.4L17 25h14z"/><path d="M22 14.8h4v12.3h-4z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MakeDecision(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#FFCC80"><circle cx="38" cy="26" r="4"/><circle cx="10" cy="26" r="4"/><path d="M39 19c0-12.7-30-8.3-30 0v10c0 8.3 6.7 15 15 15s15-6.7 15-15z"/><path d="M24 4C15.2 4 8 11.2 8 20v3.5l2.1.6V19l19.5-6.3l8.2 6.3v5.1l2.1-.6V20C40 12.5 34.6 4 24 4"/></g><path fill="#FF5722" d="M24 23.5v-11l6.6 5.5z"/><path fill="#FF5722" d="M28.9 24.4c0 .2.1.4.1.6c0 2.8-2.2 5-5 5s-5-2.2-5-5s2.2-5 5-5c.7 0 1.4.2 2 .4v-4.2c-.6-.1-1.3-.2-2-.2c-5 0-9 4-9 9s4 9 9 9s9-4 9-9c0-1.2-.2-2.4-.7-3.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Manager(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#FF9800" d="m24 37l-5-6v-6h10v6z"/><g fill="#FFA726"><circle cx="33" cy="19" r="2"/><circle cx="15" cy="19" r="2"/></g><path fill="#FFB74D" d="M33 13c0-7.6-18-5-18 0v7c0 5 4 9 9 9s9-4 9-9z"/><path fill="#FF5722" d="M24 4c-6.1 0-10 4.9-10 11v2.3l2 1.7v-5l12-4l4 4v5l2-1.7V15c0-4-1-8-6-9l-1-2z"/><g fill="#784719"><circle cx="28" cy="19" r="1"/><circle cx="20" cy="19" r="1"/></g><path fill="#CFD8DC" d="m29 31l-5 1l-5-1S8 33 8 44h32c0-11-11-13-11-13"/><path fill="#3F51B5" d="m23 35l-1 9h4l-1-9l1-1l-2-2l-2 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediumPriority(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#FFC107" d="m21.2 44.8l-18-18c-1.6-1.6-1.6-4.1 0-5.7l18-18c1.6-1.6 4.1-1.6 5.7 0l18 18c1.6 1.6 1.6 4.1 0 5.7l-18 18c-1.6 1.6-4.2 1.6-5.7 0"/><g fill="#37474F"><circle cx="24" cy="24" r="2"/><circle cx="32" cy="24" r="2"/><circle cx="16" cy="24" r="2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Menu(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#607D8B" d="M6 22h36v4H6zm0-12h36v4H6zm0 24h36v4H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MiddleBattery(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#CFD8DC"><path d="M34 44H14c-1.1 0-2-.9-2-2V8c0-1.1.9-2 2-2h20c1.1 0 2 .9 2 2v34c0 1.1-.9 2-2 2"/><path d="M28 13h-8c-.6 0-1-.4-1-1V5c0-.6.4-1 1-1h8c.6 0 1 .4 1 1v7c0 .6-.4 1-1 1"/></g><path fill="#8BC34A" d="M34 44H14c-1.1 0-2-.9-2-2V23h24v19c0 1.1-.9 2-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MindMap(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#CFD8DC" d="m39.4 23l-.8-4L26 21.6V8h-4v12.3l-13.9-9l-2.2 3.4l15.2 9.8L9.4 39.8l3.2 2.4l11.3-14.8l8.4 12.7l3.4-2.2l-8.4-12.5z"/><circle cx="24" cy="24" r="7" fill="#3F51B5"/><g fill="#00BCD4"><circle cx="24" cy="8" r="5"/><circle cx="39" cy="21" r="5"/><circle cx="7" cy="13" r="5"/><circle cx="11" cy="41" r="5"/><circle cx="34" cy="39" r="5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Minus(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#5C6BC0" d="M8 21h32v6H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MissedCall(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#F44336"><path d="M30.3 12.9L24 19.2l-8.3-8.3l-2.8 2.8L24 24.8l9.1-9.1z"/><path d="m36 19l-9-9h9z"/></g><path fill="#009688" d="m44.5 30.8l-2.4-2.4c-8.5-8.3-28.9-7.1-36.2 0l-2.4 2.4c-.7.7-.7 1.7 0 2.4l4.8 4.7c.7.7 1.7.7 2.4 0l5.3-5.1l-.4-5.6c1.7-1.7 15.1-1.7 16.8 0l-.3 5.8l5.1 4.9c.7.7 1.7.7 2.4 0l4.8-4.7c.8-.7.8-1.8.1-2.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mms(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#E91E63" d="M37 39H11l-6 6V11c0-3.3 2.7-6 6-6h26c3.3 0 6 2.7 6 6v22c0 3.3-2.7 6-6 6"/><path fill="#F48FB1" d="M20 16.5L10 31h20z"/><g fill="#F8BBD0"><circle cx="34" cy="15" r="3"/><path d="m30 21l-8 10h16z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoneyTransfer(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#78909C" d="M40 41H8c-2.2 0-4-1.8-4-4V16.1c0-1.3.6-2.5 1.7-3.3L24 0l18.3 12.8c1.1.7 1.7 2 1.7 3.3V37c0 2.2-1.8 4-4 4"/><path fill="#AED581" d="M14 1h20v31H14z"/><g fill="#558B2F"><path d="M13 0v33h22V0zm20 31H15V2h18z"/><path d="M34 3c0 1.7-.3 3-2 3s-3-1.3-3-3s1.3-2 3-2s2 .3 2 2M16 1c1.7 0 3 .3 3 2s-1.3 3-3 3s-2-1.3-2-3s.3-2 2-2"/><circle cx="24" cy="8" r="2"/><circle cx="24" cy="20" r="6"/></g><path fill="#CFD8DC" d="M40 41H8c-2.2 0-4-1.8-4-4V17l20 13l20-13v20c0 2.2-1.8 4-4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MultipleCameras(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#455A64" d="M42 41H12c-2.2 0-4-1.8-4-4V17c0-2.2 1.8-4 4-4h30c2.2 0 4 1.8 4 4v20c0 2.2-1.8 4-4 4"/><path fill="#78909C" d="M36 36H6c-2.2 0-4-1.8-4-4V12c0-2.2 1.8-4 4-4h30c2.2 0 4 1.8 4 4v20c0 2.2-1.8 4-4 4"/><circle cx="26" cy="22" r="10" fill="#455A64"/><circle cx="26" cy="22" r="7" fill="#42A5F5"/><path fill="#90CAF9" d="M29.7 19.7c-1-1.1-2.3-1.7-3.7-1.7s-2.8.6-3.7 1.7c-.4.4-.3 1 .1 1.4c.4.4 1 .3 1.4-.1c1.2-1.3 3.3-1.3 4.5 0c.2.2.5.3.7.3c.2 0 .5-.1.7-.3c.4-.3.4-.9 0-1.3"/><path fill="#ADD8FB" d="M6 12h6v3H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MultipleDevices(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#546E7A" d="M4 28V8c0-2.2 1.8-4 4-4h28c2.2 0 4 1.8 4 4v20c0 2.2-1.8 4-4 4H8c-2.2 0-4-1.8-4-4"/><path fill="#BBDEFB" d="M36 7H8c-.6 0-1 .4-1 1v20c0 .6.4 1 1 1h28c.6 0 1-.4 1-1V8c0-.6-.4-1-1-1"/><path fill="#37474F" d="M38 33H6c-2.2 0-4-1.8-4-4h40c0 2.2-1.8 4-4 4"/><path fill="#E38939" d="M24 40V16c0-2.2 1.8-4 4-4h12c2.2 0 4 1.8 4 4v24c0 2.2-1.8 4-4 4H28c-2.2 0-4-1.8-4-4"/><path fill="#FFF3E0" d="M40 15H28c-.6 0-1 .4-1 1v22c0 .6.4 1 1 1h12c.6 0 1-.4 1-1V16c0-.6-.4-1-1-1"/><circle cx="34" cy="41.5" r="1.5" fill="#A6642A"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MultipleInputs(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#90A4AE" d="M40 35v5H8v-5H4v5c0 2.2 1.8 4 4 4h32c2.2 0 4-1.8 4-4v-5z"/><g fill="#1565C0"><path d="M24 23.4L17 15h14z"/><path d="M22 4h4v14h-4zm9.5 22.9l-.7 1.1l3.5 1.9l.6-1.2c1.6-3 2.6-4.7 3.5-5.2c.9-.5 2.6-.5 5.6-.5v-4c-7.7 0-8.4.4-12.5 7.9"/><path d="m38.4 31l-9 4L28 25zm-21.9-4.1l.6 1.2l-3.5 1.9l-.6-1.2c-1.6-3-2.6-4.7-3.5-5.2C8.7 23 7 23 4 23v-4c7.7 0 8.4.4 12.5 7.9"/><path d="m20 25l-1.4 10l-9-4z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MultipleSmartphones(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#455A64" d="M4 31V8c0-2.2 1.8-4 4-4h12c2.2 0 4 1.8 4 4v23c0 2.2-1.8 4-4 4H8c-2.2 0-4-1.8-4-4"/><path fill="#BBDEFB" d="M20 7H8c-.6 0-1 .4-1 1v21c0 .6.4 1 1 1h12c.6 0 1-.4 1-1V8c0-.6-.4-1-1-1"/><circle cx="14" cy="32.5" r="1.5" fill="#37474F"/><path fill="#546E7A" d="M14 36V13c0-2.2 1.8-4 4-4h12c2.2 0 4 1.8 4 4v23c0 2.2-1.8 4-4 4H18c-2.2 0-4-1.8-4-4"/><path fill="#BBDEFB" d="M30 12H18c-.6 0-1 .4-1 1v21c0 .6.4 1 1 1h12c.6 0 1-.4 1-1V13c0-.6-.4-1-1-1"/><circle cx="24" cy="37.5" r="1.5" fill="#37474F"/><path fill="#E38939" d="M24 40V18c0-2.2 1.8-4 4-4h12c2.2 0 4 1.8 4 4v22c0 2.2-1.8 4-4 4H28c-2.2 0-4-1.8-4-4"/><path fill="#FFF3E0" d="M40 17H28c-.6 0-1 .4-1 1v20c0 .6.4 1 1 1h12c.6 0 1-.4 1-1V18c0-.6-.4-1-1-1"/><circle cx="34" cy="41.5" r="1.5" fill="#A6642A"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Music(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#E91E63"><circle cx="19" cy="33" r="9"/><path d="M24 6v27h4V14l11 3v-7z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NegativeDynamic(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#00BCD4" d="M19 22h10v20H19zM6 8h10v34H6zm26 22h10v12H32z"/><g fill="#3F51B5"><path d="M42 12L32 22h10z"/><path d="m27.561 10.396l2.828-2.828l9.969 9.969l-2.828 2.828z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NeutralDecision(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#FFCC80"><circle cx="38" cy="26" r="4"/><circle cx="10" cy="26" r="4"/><path d="M39 19c0-12.7-30-8.3-30 0v10c0 8.3 6.7 15 15 15s15-6.7 15-15z"/><path d="M24 4C15.2 4 8 11.2 8 20v3.5l2.1.6V19l19.5-6.3l8.2 6.3v5.1l2.1-.6V20C40 12.5 34.6 4 24 4"/></g><g fill="#37474F"><circle cx="24" cy="25" r="2"/><circle cx="32" cy="25" r="2"/><circle cx="16" cy="25" r="2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NeutralTrading(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#1565C0"><path d="M43.4 13L35 20V6z"/><path d="M4 11h34v4H4z"/></g><path fill="#2196F3" d="M40 23h4v19h-4zm-6 0h4v19h-4zm-6 0h4v19h-4zm-6 0h4v19h-4zm-6 0h4v19h-4zm-6 0h4v19h-4zm-6 0h4v19H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func News(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#FF5722" d="M32 15v28H10c-2.2 0-4-1.8-4-4V15z"/><path fill="#FFCCBC" d="M14 5v34c0 2.2-1.8 4-4 4h29c2.2 0 4-1.8 4-4V5z"/><path fill="#FF5722" d="M20 10h18v4H20zm0 7h8v2h-8zm10 0h8v2h-8zm-10 4h8v2h-8zm10 0h8v2h-8zm-10 4h8v2h-8zm10 0h8v2h-8zm-10 4h8v2h-8zm10 0h8v2h-8zm-10 4h8v2h-8zm10 0h8v2h-8zm-10 4h8v2h-8zm10 0h8v2h-8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Next(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#2196F3" d="M17.1 5L14 8.1L29.9 24L14 39.9l3.1 3.1L36 24z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NfcSign(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#2196F3" d="M37 42c-.3 0-.7-.1-1-.3c-1-.5-1.3-1.8-.8-2.7c0-.1 3.7-6.8 3.7-15S35.3 9 35.3 9c-.5-1-.2-2.2.8-2.7s2.2-.2 2.7.8c.2.3 4.3 7.6 4.3 17s-4.1 16.7-4.3 17c-.4.5-1.1.9-1.8.9m-4.2-6.2c.1-.2 2.2-5 2.2-11.8c0-6.8-2.1-11.6-2.2-11.8c-.4-1-1.6-1.5-2.6-1c-1 .4-1.5 1.6-1 2.6c0 0 1.8 4.3 1.8 10.2c0 5.9-1.8 10.2-1.8 10.2c-.4 1 0 2.2 1 2.6c.3.1.5.2.8.2c.8 0 1.5-.4 1.8-1.2M23.3 33c.6-.1 1.1-.5 1.4-1c.1-.2 2.3-3.9 2.3-8s-2.2-7.9-2.3-8c-.6-1-1.8-1.3-2.7-.7c-1 .6-1.3 1.8-.7 2.7c0 0 1.7 3 1.7 6c0 1.3-.3 2.7-.7 3.7l-13-11.2c-.5-.4-1.2-.6-1.8-.4c-.6.2-1.2.6-1.4 1.3c0 .1-1.1 3.1-1.1 6.6c0 3.5 1.1 6.5 1.1 6.7c.4 1 1.5 1.6 2.6 1.2c1-.4 1.6-1.5 1.2-2.6c0 0-.9-2.6-.9-5.3c0-.8.1-1.6.2-2.3l12.5 10.8c.4.3.8.5 1.3.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightLandscape(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#673AB7" d="M16.5 18L0 42h33z"/><path fill="#9575CD" d="M33.6 24L19.2 42H48z"/><path fill="#40C4FF" d="M42.9 6.3C43.6 7.4 44 8.6 44 10c0 3.9-3.1 7-7 7c-.7 0-1.3-.1-1.9-.3c1.2 2 3.4 3.3 5.9 3.3c3.9 0 7-3.1 7-7c0-3.2-2.1-5.9-5.1-6.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightPortrait(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#40C4FF" d="M42.9 6.3C43.6 7.4 44 8.6 44 10c0 3.9-3.1 7-7 7c-.7 0-1.3-.1-1.9-.3c1.2 2 3.4 3.3 5.9 3.3c3.9 0 7-3.1 7-7c0-3.2-2.1-5.9-5.1-6.7"/><g fill="#B39DDB"><circle cx="31" cy="19" r="2"/><circle cx="13" cy="19" r="2"/><path d="m22 37l-5-6v-6h10v6z"/></g><path fill="#D1C4E9" d="M31 13c0-7.6-18-5-18 0v7c0 5 4 9 9 9s9-4 9-9z"/><g fill="#673AB7"><circle cx="26" cy="19" r="1"/><circle cx="18" cy="19" r="1"/><path d="M22 4c-6.1 0-10 4.9-10 11v2.3l2 1.7v-5l12-4l4 4v5l2-1.7V15c0-4-1-8-6-9l-1-2zm5 27s-2 1-5 1s-5-1-5-1S6 33 6 44h32c0-11-11-13-11-13"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoIdea(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#FBC02D" d="M37 22c0-7.7-6.6-13.8-14.5-12.9c-6 .7-10.8 5.5-11.4 11.5c-.5 4.6 1.4 8.7 4.6 11.3c1.4 1.2 2.3 2.9 2.3 4.8v.3h12v-.1c0-1.8.8-3.6 2.2-4.8c2.9-2.4 4.8-6 4.8-10.1"/><path fill="#FFF59D" d="m30.6 20.2l-3-2c-.3-.2-.8-.2-1.1 0L24 19.8l-2.4-1.6c-.3-.2-.8-.2-1.1 0l-3 2c-.2.2-.4.4-.4.7s0 .6.2.8l3.8 4.7V37h2V26c0-.2-.1-.4-.2-.6l-3.3-4.1l1.5-1l2.4 1.6c.3.2.8.2 1.1 0l2.4-1.6l1.5 1l-3.3 4.1c-.1.2-.2.4-.2.6v11h2V26.4l3.8-4.7c.2-.2.3-.5.2-.8s-.2-.6-.4-.7"/><circle cx="24" cy="44" r="3" fill="#5C6BC0"/><path fill="#9FA8DA" d="M26 45h-4c-2.2 0-4-1.8-4-4v-5h12v5c0 2.2-1.8 4-4 4"/><path fill="#5C6BC0" d="m30 41l-11.6 1.6c.3.7.9 1.4 1.6 1.8l9.4-1.3c.4-.6.6-1.3.6-2.1m-12-2.3v2L30 39v-2z"/><path fill="#37474F" d="M3.563 6.396L6.39 3.568l37.966 37.966l-2.828 2.828z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoVideo(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#4CAF50" d="M8 12h22c2.2 0 4 1.8 4 4v16c0 2.2-1.8 4-4 4H8c-2.2 0-4-1.8-4-4V16c0-2.2 1.8-4 4-4"/><path fill="#388E3C" d="m44 35l-10-6V19l10-6z"/><path fill="none" stroke="#212121" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="4" d="m5 5l38 38"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Nook(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#90A4AE" d="M8 39V9c0-3.3 2.7-6 6-6h20c3.3 0 6 2.7 6 6v30c0 3.3-2.7 6-6 6H14c-3.3 0-6-2.7-6-6"/><path fill="#ECEFF1" d="M34 7H14c-1.1 0-2 .9-2 2v26c0 1.1.9 2 2 2h20c1.1 0 2-.9 2-2V9c0-1.1-.9-2-2-2"/><path fill="#B0BEC5" d="M16 12h16v3H16zm0 7h16v2H16zm0 4h12v2H16zm0 4h16v2H16zm0 4h12v2H16z"/><path fill="none" stroke="#eee" stroke-miterlimit="10" stroke-width="2" d="M22 43v-1c0-1.1.9-2 2-2s2 .9 2 2v1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NumericalSortingTwelve(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#546E7A" d="M38 33V5h-4v28h-6l8 10l8-10z"/><path fill="#2196F3" d="M16.4 20h-3V8.6L9.9 9.7V7.3L16 5.1h.3V20zm3 23H9.2v-2l4.8-5.1c.4-.4.7-.8.9-1.1c.2-.3.5-.6.6-.9c.2-.3.3-.5.3-.8c.1-.2.1-.5.1-.7c0-.7-.2-1.2-.5-1.6c-.3-.4-.8-.6-1.4-.6c-.3 0-.7.1-.9.2c-.3.1-.5.3-.7.5c-.2.2-.3.5-.4.8s-.1.6-.1 1h-3c0-.7.1-1.3.4-1.9c.2-.6.6-1.1 1-1.6c.5-.4 1-.8 1.6-1.1c.6-.3 1.4-.4 2.2-.4c.8 0 1.5.1 2.1.3c.6.2 1.1.5 1.5.8s.7.8.9 1.3s.3 1.1.3 1.8c0 .5-.1 1-.2 1.4s-.4 1.2-.7 1.7s-.6.9-1 1.4c-.4.5-.9 1-1.4 1.5L13 40.6h6.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NumericalSortingTwentyOne(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#546E7A" d="M38 33V5h-4v28h-6l8 10l8-10z"/><path fill="#2196F3" d="M19.2 20H9v-2l4.8-5.1c.4-.4.7-.8.9-1.1c.2-.3.5-.6.6-.9c.2-.3.3-.5.3-.8c.1-.2.1-.5.1-.7c0-.7-.2-1.2-.5-1.6c-.3-.4-.8-.6-1.4-.6c-.3 0-.7.1-.9.2c-.3.1-.5.3-.7.5c-.2.2-.3.5-.4.8s-.1.6-.1 1h-3c0-.7.1-1.3.4-1.9c.2-.6.6-1.1 1-1.6c.5-.4 1-.8 1.6-1.1c.6-.3 1.4-.4 2.2-.4c.8 0 1.5.1 2.1.3c.6.2 1.1.5 1.5.8s.7.8.9 1.3c.2.5.3 1.1.3 1.8c0 .5-.1 1-.2 1.4s-.4.9-.7 1.4s-.6.9-1 1.4c-.4.5-.9 1-1.4 1.5l-2.6 2.8h6.4zm-3 23h-3V31.6l-3.5 1.1v-2.4l6.2-2.2h.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ok(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="24" cy="24" r="21" fill="#4CAF50"/><path fill="#CCFF90" d="M34.6 14.6L21 28.2l-5.6-5.6l-2.8 2.8l8.4 8.4l16.4-16.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OldTimeCamera(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#546E7A" d="M14 13H8v-1.8c0-.7.5-1.2 1.2-1.2h3.6c.7 0 1.2.5 1.2 1.2z"/><path fill="#5E35B1" d="M40 40H8c-2.2 0-4-1.8-4-4V22h40v14c0 2.2-1.8 4-4 4"/><path fill="#42257A" d="M12.7 22c-.4 1.3-.7 2.6-.7 4c0 6.6 5.4 12 12 12s12-5.4 12-12c0-1.4-.3-2.7-.7-4z"/><path fill="#78909C" d="M8 12h32c2.2 0 4 1.8 4 4v6H4v-6c0-2.2 1.8-4 4-4"/><path fill="#78909C" d="M33.9 13.1H14.2L17.6 8c.4-.6 1-.9 1.7-.9h9.6c.7 0 1.3.3 1.7.9z"/><path fill="#455A64" d="M35.3 22c-1.6-4.7-6.1-8-11.3-8s-9.7 3.3-11.3 8z"/><circle cx="24" cy="26" r="9" fill="#B388FF"/><path fill="#C7A7FF" d="M29 23c-1.2-1.4-3-2.2-4.8-2.2c-1.8 0-3.6.8-4.8 2.2c-.5.5-.4 1.3.1 1.8s1.3.4 1.8-.1c1.5-1.7 4.3-1.7 5.8 0c.3.3.6.4 1 .4c.3 0 .6-.1.9-.3c.4-.4.5-1.3 0-1.8"/><path fill="#DBE2E5" d="M36 15h5v4h-5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OnlineSupport(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#BF360C" d="M13 30h22v12H13z"/><g fill="#FFA726"><circle cx="10" cy="26" r="4"/><circle cx="38" cy="26" r="4"/></g><path fill="#FFB74D" d="M39 19c0-12.7-30-8.3-30 0v10c0 8.3 6.7 15 15 15s15-6.7 15-15z"/><g fill="#784719"><circle cx="30" cy="26" r="2"/><circle cx="18" cy="26" r="2"/></g><path fill="#FF5722" d="M24 2C15.5 2 3 7.8 3 35.6L13 42V24l16.8-9.8L35 21v21l10-8.2c0-5.6-.9-29-15.4-29L28.2 2z"/><path fill="#757575" d="M45 24c-.6 0-1 .4-1 1v-7c0-8.8-7.2-16-16-16h-9c-.6 0-1 .4-1 1s.4 1 1 1h9c7.7 0 14 6.3 14 14v10c0 .6.4 1 1 1s1-.4 1-1v2c0 3.9-3.1 7-7 7H24c-.6 0-1 .4-1 1s.4 1 1 1h13c5 0 9-4 9-9v-5c0-.6-.4-1-1-1"/><g fill="#37474F"><path d="M45 22h-1c-1.1 0-2 .9-2 2v4c0 1.1.9 2 2 2h1c1.1 0 2-.9 2-2v-4c0-1.1-.9-2-2-2"/><circle cx="24" cy="38" r="2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OpenedFolder(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#FFA000" d="M38 12H22l-4-4H8c-2.2 0-4 1.8-4 4v24c0 2.2 1.8 4 4 4h31c1.7 0 3-1.3 3-3V16c0-2.2-1.8-4-4-4"/><path fill="#FFCA28" d="M42.2 18H15.3c-1.9 0-3.6 1.4-3.9 3.3L8 40h31.7c1.9 0 3.6-1.4 3.9-3.3l2.5-14c.5-2.4-1.4-4.7-3.9-4.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OrgUnit(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#90CAF9" d="M10 10v28h28V10zm24 24H14V14h20z"/><path fill="#D81B60" d="M6 6h12v12H6z"/><path fill="#2196F3" d="M30 6h12v12H30zM6 30h12v12H6zm24 0h12v12H30z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Organization(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#90CAF9" d="M42 42H6V10c0-2.2 1.8-4 4-4h28c2.2 0 4 1.8 4 4z"/><path fill="#64B5F6" d="M6 42h36v2H6z"/><path fill="#1565C0" d="M31 27h6v5h-6zm-10 0h6v5h-6zm-10 0h6v5h-6zm20 8h6v5h-6zm-20 0h6v5h-6zm20-16h6v5h-6zm-10 0h6v5h-6zm-10 0h6v5h-6zm20-8h6v5h-6zm-10 0h6v5h-6zm-10 0h6v5h-6zm10 24h6v9h-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Overtime(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#CFD8DC" d="M12 40V20h32v20c0 2.2-1.8 4-4 4H16c-2.2 0-4-1.8-4-4"/><path fill="#78909C" d="M44 16v6H12v-6c0-2.2 1.8-4 4-4h24c2.2 0 4 1.8 4 4"/><g fill="#37474F"><circle cx="37" cy="16" r="3"/><circle cx="20" cy="16" r="3"/></g><path fill="#B0BEC5" d="M37 10c-1.1 0-2 .9-2 2v4c0 1.1.9 2 2 2s2-.9 2-2v-4c0-1.1-.9-2-2-2m-17 0c-1.1 0-2 .9-2 2v4c0 1.1.9 2 2 2s2-.9 2-2v-4c0-1.1-.9-2-2-2"/><path fill="#90A4AE" d="M32 34h4v4h-4zm-6 0h4v4h-4zm-6 0h4v4h-4zm12-6h4v4h-4zm-6 0h4v4h-4zm-6 0h4v4h-4z"/><circle cx="16" cy="15" r="12" fill="#F44336"/><circle cx="16" cy="15" r="9" fill="#eee"/><path d="M15 8h2v7h-2z"/><path d="m20.518 18.1l-1.343 1.344l-3.818-3.818l1.344-1.343z"/><circle cx="16" cy="15" r="1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Package(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#FF9800" d="M38 42H10c-2.2 0-4-1.8-4-4V10c0-2.2 1.8-4 4-4h28c2.2 0 4 1.8 4 4v28c0 2.2-1.8 4-4 4"/><path fill="#8A5100" d="M29.5 16h-11c-.8 0-1.5-.7-1.5-1.5s.7-1.5 1.5-1.5h11c.8 0 1.5.7 1.5 1.5s-.7 1.5-1.5 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paid(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#2E7D32" d="M25.4 5.6c-.8-.8-2-.8-2.8 0l-12 12c-.8.8-.8 2 0 2.8c.4.4.9.6 1.4.6s1-.2 1.4-.6l12-12c.8-.8.8-2 0-2.8"/><path fill="#1B5E20" d="m37.4 17.6l-12-12c-.8-.8-2-.8-2.8 0c-.8.8-.8 2 0 2.8l12 12c.4.4.9.6 1.4.6s1-.2 1.4-.6c.8-.8.8-2 0-2.8"/><path fill="#388E3C" d="M37.4 41H10.6c-1 0-1.8-.7-2-1.6L5 21h38l-3.7 18.4c-.2.9-1 1.6-1.9 1.6"/><path fill="#4CAF50" d="M43 23H5c-1.1 0-2-.9-2-2v-2c0-1.1.9-2 2-2h38c1.1 0 2 .9 2 2v2c0 1.1-.9 2-2 2"/><path fill="#DCEDC8" d="m30.8 24.8l-7.9 7.9l-3.7-3.8l-2.2 2.2l5.9 5.9L33 26.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Panorama(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#F57C00" d="M4 9v32s8.4-3 20-3s20 3 20 3V9s-6.7 3-20 3S4 9 4 9"/><path fill="#942A09" d="M24 34h.4L15 19L6.9 36.2c3.4-.9 9.6-2.2 17.1-2.2"/><path fill="#BF360C" d="M24 34c3.3 0 6.3.2 9 .6l-8-11.8l-7.8 11.5c2.1-.2 4.4-.3 6.8-.3"/><path fill="#E65100" d="M40.7 36L35 26.5l-5 7.8c4.5.4 8.2 1.1 10.7 1.7"/><ellipse cx="36" cy="19.5" fill="#FFF9C4" rx="2" ry="2.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ParallelTasks(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#90CAF9" d="M36 13V9H22v13h-9v4h9v13h14v-4H26v-9h10v-4H26v-9z"/><path fill="#D81B60" d="M6 17h10v14H6z"/><path fill="#2196F3" d="M32 6h10v10H32zm0 26h10v10H32zm0-13h10v10H32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Phone(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#009688" d="M39.1 7h-3.7C22.2 7.2 7.1 24.1 7 35.4v3.7c0 1 .8 1.9 1.9 1.9l7.5-.1c1 0 1.9-.9 1.9-1.9l.2-8.2l-4.7-4c0-2.6 10.5-13.1 13.2-13.2l4.3 4.7l7.9-.2c1 0 1.9-.9 1.9-1.9L41 8.9c0-1.1-.8-1.9-1.9-1.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneAndroid(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#37474F" d="M12 40V8c0-2.2 1.8-4 4-4h16c2.2 0 4 1.8 4 4v32c0 2.2-1.8 4-4 4H16c-2.2 0-4-1.8-4-4"/><path fill="#BBDEFB" d="M32 7H16c-.6 0-1 .4-1 1v29c0 .6.4 1 1 1h16c.6 0 1-.4 1-1V8c0-.6-.4-1-1-1"/><path fill="#78909C" d="M21 40h6v2h-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhotoReel(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#673AB7" d="M10 9c-2.2 0-4 1.8-4 4v26c0 2.2 1.8 4 4 4h16c2.2 0 4-1.8 4-4V13c0-2.2-1.8-4-4-4"/><path fill="#311B92" d="M14 13h2v26h-2zm10-4V7c0-1.2-.8-2-2-2h-8c-1.2 0-2 .8-2 2v2z"/><path fill="#D84315" d="M30 13H16v26h14zm-9 24h-3v-4h3zm0-18h-3v-4h3zm6 18h-3v-4h3zm-3-18v-4h3v4z"/><path fill="#FF5722" d="M30 13v2h3v4h-3v14h3v4h-3v2h12V13zm9 24h-3v-4h3zm0-18h-3v-4h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Picture(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#F57C00" d="M40 41H8c-2.2 0-4-1.8-4-4V11c0-2.2 1.8-4 4-4h32c2.2 0 4 1.8 4 4v26c0 2.2-1.8 4-4 4"/><circle cx="35" cy="16" r="3" fill="#FFF9C4"/><path fill="#942A09" d="M20 16L9 32h22z"/><path fill="#BF360C" d="m31 22l-8 10h16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PieChart(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#00BCD4" d="M24 6C14.1 6 6 14.1 6 24s8.1 18 18 18c5.2 0 9.9-2.2 13.1-5.7L24 24z"/><path fill="#448AFF" d="M42 24c0-9.9-8.1-18-18-18v18z"/><path fill="#3F51B5" d="m24 24l13.1 12.3c3-3.2 4.9-7.5 4.9-12.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Planner(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#CFD8DC" d="M5 38V14h38v24c0 2.2-1.8 4-4 4H9c-2.2 0-4-1.8-4-4"/><path fill="#F44336" d="M43 10v6H5v-6c0-2.2 1.8-4 4-4h30c2.2 0 4 1.8 4 4"/><g fill="#B71C1C"><circle cx="33" cy="10" r="3"/><circle cx="15" cy="10" r="3"/></g><path fill="#B0BEC5" d="M33 3c-1.1 0-2 .9-2 2v5c0 1.1.9 2 2 2s2-.9 2-2V5c0-1.1-.9-2-2-2M15 3c-1.1 0-2 .9-2 2v5c0 1.1.9 2 2 2s2-.9 2-2V5c0-1.1-.9-2-2-2m-2 18h6v6h-6zm8 0h6v6h-6zm8 0h6v6h-6zm-16 8h6v6h-6zm8 0h6v6h-6z"/><path fill="#F44336" d="M29 29h6v6h-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plus(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="24" cy="24" r="21" fill="#4CAF50"/><g fill="#fff"><path d="M21 14h6v20h-6z"/><path d="M14 21h20v6H14z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PodiumWithAudience(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#B0BEC5" d="M41 12H7l-1 4l5 3l-2-3h30l-2 3l5-3z"/><path fill="#78909C" d="M9 16h30l-4 12H13z"/><circle cx="24" cy="28" r="4" fill="#FFB74D"/><circle cx="36" cy="28" r="4" fill="#FFB74D"/><circle cx="12" cy="28" r="4" fill="#FFB74D"/><circle cx="18" cy="37" r="5" fill="#FFB74D"/><circle cx="30" cy="37" r="5" fill="#FFB74D"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PodiumWithSpeaker(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="24" cy="11" r="6" fill="#FFB74D"/><path fill="#607D8B" d="M36 26.1S32.7 19 24 19s-12 7.1-12 7.1V30h24z"/><path fill="#B0BEC5" d="M41 25H7l-1 4l5 3l-2-3h30l-2 3l5-3z"/><path fill="#78909C" d="M9 29h30l-4 12H13z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PodiumWithoutSpeaker(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#B0BEC5" d="M43 16H5l-1 4l5 3l-2-3h34l-2 3l5-3z"/><path fill="#78909C" d="M7 20h34l-4 16H11z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PortraitMode(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#FF9800" d="M22 38c-4.8 0-5-7-5-7v-6h10v6s-.2 7-5 7"/><g fill="#FFA726"><circle cx="31" cy="19" r="2"/><circle cx="13" cy="19" r="2"/></g><path fill="#FFB74D" d="M31 13c0-7.6-18-5-18 0v7c0 5 4 9 9 9s9-4 9-9z"/><path fill="#424242" d="M22 4c-6.1 0-10 4.9-10 11v2.3l2 1.7v-5l12-4l4 4v5l2-1.7V15c0-4-1-8-6-9l-1-2z"/><g fill="#784719"><circle cx="26" cy="19" r="1"/><circle cx="18" cy="19" r="1"/></g><path fill="#009688" d="M27 31s-1.8 2-5 2s-5-2-5-2S6 33 6 44h32c0-11-11-13-11-13"/><g fill="#FF9800"><path d="m40.997 4.065l7 7l-7 6.999l-7-7z"/><path d="M36 6h10v10H36z"/></g><circle cx="41" cy="11" r="3" fill="#FFEB3B"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PositiveDynamic(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#00BCD4" d="M19 22h10v20H19zM32 8h10v34H32zM6 30h10v12H6z"/><g fill="#3F51B5"><path d="m11 8l10 10V8z"/><path d="m9.394 22.437l-2.828-2.828l9.969-9.969l2.828 2.828z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Previous(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#2196F3" d="m30.9 43l3.1-3.1L18.1 24L34 8.1L30.9 5L12 24z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Print(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#424242" d="M9 11h30v3H9z"/><path fill="#616161" d="M4 25h40v-7c0-2.2-1.8-4-4-4H8c-2.2 0-4 1.8-4 4z"/><path fill="#424242" d="M8 36h32c2.2 0 4-1.8 4-4v-8H4v8c0 2.2 1.8 4 4 4"/><circle cx="40" cy="18" r="1" fill="#00E676"/><path fill="#90CAF9" d="M11 4h26v10H11z"/><path fill="#242424" d="M37.5 31h-27c-.8 0-1.5-.7-1.5-1.5s.7-1.5 1.5-1.5h27c.8 0 1.5.7 1.5 1.5s-.7 1.5-1.5 1.5"/><path fill="#90CAF9" d="M11 31h26v11H11z"/><path fill="#42A5F5" d="M11 29h26v2H11z"/><path fill="#1976D2" d="M16 33h17v2H16zm0 4h13v2H16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Privacy(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#424242" d="M24 4c-5.5 0-10 4.5-10 10v4h4v-4c0-3.3 2.7-6 6-6s6 2.7 6 6v4h4v-4c0-5.5-4.5-10-10-10"/><path fill="#FB8C00" d="M36 44H12c-2.2 0-4-1.8-4-4V22c0-2.2 1.8-4 4-4h24c2.2 0 4 1.8 4 4v18c0 2.2-1.8 4-4 4"/><circle cx="24" cy="31" r="6" fill="#EFEBE9"/><circle cx="24" cy="31" r="3" fill="#1E88E5"/><circle cx="26" cy="29" r="1" fill="#fff"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Process(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#9C27B0"><path d="m31 8l11.9 1.6l-9.8 9.8zM17 40L5.1 38.4l9.8-9.8zM8 17L9.6 5.1l9.8 9.8z"/><path d="m9.3 21.2l-4.2.8c-.1.7-.1 1.3-.1 2c0 4.6 1.6 9 4.6 12.4l3-2.6C10.3 31.1 9 27.6 9 24c0-.9.1-1.9.3-2.8M24 5c-5.4 0-10.2 2.3-13.7 5.9l2.8 2.8C15.9 10.8 19.7 9 24 9c.9 0 1.9.1 2.8.3l.7-3.9C26.4 5.1 25.2 5 24 5m14.7 21.8l4.2-.8c.1-.7.1-1.3.1-2c0-4.4-1.5-8.7-4.3-12.1l-3.1 2.5c2.2 2.7 3.4 6.1 3.4 9.5c0 1-.1 2-.3 2.9m-3.8 7.5C32.1 37.2 28.3 39 24 39c-.9 0-1.9-.1-2.8-.3l-.7 3.9c1.2.2 2.4.3 3.5.3c5.4 0 10.2-2.3 13.7-5.9z"/><path d="m40 31l-1.6 11.9l-9.8-9.8z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Puzzle(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#8BC34A" d="M39 15c0-2.2-1.8-4-4-4h-6c-.7 0-1.1-.8-.7-1.4c.6-1 .9-2.2.6-3.5c-.4-2-1.9-3.6-3.8-4C21.8 1.4 19 3.9 19 7c0 1 .3 1.8.7 2.6c.4.6 0 1.4-.8 1.4h-6c-2.2 0-4 1.8-4 4v7c0 .7.8 1.1 1.4.7c1-.6 2.2-.9 3.5-.6c2 .4 3.6 1.9 4 3.8c.7 3.2-1.8 6.1-4.9 6.1c-1 0-1.8-.3-2.6-.7c-.5-.4-1.3 0-1.3.7v6c0 2.2 1.8 4 4 4h22c2.2 0 4-1.8 4-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Questions(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#42A5F5" d="M36 44H8V8h20l8 8z"/><path fill="#90CAF9" d="M40 40H12V4h20l8 8z"/><path fill="#E1F5FE" d="M38.5 13H31V5.5z"/><path fill="#1976D2" d="M24.5 28.3c0-4.7 3.6-4.4 3.6-7.2c0-.7-.2-2.1-2-2.1c-2 0-2.1 1.6-2.1 2h-2.7c0-.7.3-4.2 4.8-4.2c4.6 0 4.7 3.6 4.7 4.3c0 3.5-3.8 4-3.8 7.3h-2.5zm-.2 3.5c0-.2 0-1.5 1.5-1.5c1.4 0 1.5 1.3 1.5 1.5c0 .4-.2 1.4-1.5 1.4s-1.5-1-1.5-1.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RadarPlot(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#CFD8DC" d="M38.4 13L24.1 6.4L4.6 12.1l8.8 13.2l-2.2 15.1h22.7l6.6-13.3zm-6.3 24.5H14.7l1.8-12.9l-7.1-10.7l14.5-4.3L35.6 15l1.8 11.7z"/><g fill="#00BCD4"><circle cx="24" cy="8" r="4"/><circle cx="37" cy="14" r="4"/><circle cx="39" cy="27" r="4"/><circle cx="7" cy="13" r="4"/><circle cx="13" cy="39" r="4"/><circle cx="15" cy="25" r="4"/><circle cx="33" cy="39" r="4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rating(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="24" cy="24" r="21" fill="#F44336"/><path fill="#FFCA28" d="m24 11l3.9 7.9l8.7 1.3l-6.3 6.1l1.5 8.7l-7.8-4.1l-7.8 4.1l1.5-8.7l-6.3-6.1l8.7-1.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ratings(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#42A5F5" d="M36 44H8V8h20l8 8z"/><path fill="#90CAF9" d="M40 40H12V4h20l8 8z"/><path fill="#E1F5FE" d="M38.5 13H31V5.5z"/><path fill="#1976D2" d="M34 20h-7l2.4 2.4l-2.4 2.5l-4-4l-6.1 6l2.2 2.2l3.9-4l4 4l4.6-4.5L34 27z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reading(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#5C6BC0" d="M40 40c-6.9 0-16 4-16 4V22s9-4 18-4z"/><path fill="#7986CB" d="M8 40c6.9 0 16 4 16 4V22s-9-4-18-4z"/><g fill="#FFB74D"><circle cx="24" cy="12" r="8"/><path d="M41 32h1c.6 0 1-.4 1-1v-4c0-.6-.4-1-1-1h-1c-1.7 0-3 1.3-3 3s1.3 3 3 3M7 26H6c-.6 0-1 .4-1 1v4c0 .6.4 1 1 1h1c1.7 0 3-1.3 3-3s-1.3-3-3-3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReadingEbook(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#2196F3" d="M33.5 27c-2.2-3-5.2-5-9.5-5s-7.3 2-9.5 5z"/><path fill="#546E7A" d="M34.1 43H13.9c-1.1 0-1.9-.8-2-1.9l-.8-13c0-1.1.9-2.1 2-2.1h21.8c1.2 0 2.1 1 2 2.1l-.8 13c-.1 1.1-.9 1.9-2 1.9"/><circle cx="34" cy="29" r="1" fill="#B0BEC5"/><g fill="#FFB74D"><circle cx="24" cy="12" r="8"/><path d="M16.1 42.4L15 43.5c-.6.6-1.6.6-2.2 0l-3.3-3.3c-.6-.6-.6-1.6 0-2.2l1.1-1.1c1.3-1.3 3.1-1.3 4.4 0l1.1 1.1c1.2 1.3 1.2 3.2 0 4.4M31.9 38l1.1-1.1c1.3-1.3 3.1-1.3 4.4 0l1.1 1.1c.6.6.6 1.6 0 2.2l-3.3 3.3c-.6.6-1.6.6-2.2 0l-1.1-1.1c-1.2-1.2-1.2-3.1 0-4.4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reddit(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#FFF" d="M12.193 19.555c-1.94-1.741-4.79-1.727-6.365.029c-1.576 1.756-1.301 5.023.926 6.632zm23.614 0c1.939-1.741 4.789-1.727 6.365.029c1.575 1.756 1.302 5.023-.927 6.632z"/><circle cx="38.32" cy="10.475" r="3.5" fill="#FFF"/><ellipse cx="24.085" cy="28.611" fill="#FFF" rx="18.085" ry="12.946"/><g fill="#D84315"><circle cx="30.365" cy="26.39" r="2.884"/><circle cx="17.635" cy="26.39" r="2.884"/></g><g fill="#37474F"><path d="M24.002 34.902c-3.252 0-6.14-.745-8.002-1.902c1.024 2.044 4.196 4 8.002 4c3.802 0 6.976-1.956 7.998-4c-1.857 1.157-4.746 1.902-7.998 1.902m17.828-7.876l-1.17-1.621c.831-.6 1.373-1.556 1.488-2.623c.105-.98-.157-1.903-.721-2.531c-.571-.637-1.391-.99-2.307-.994a4.083 4.083 0 0 0-2.646 1.041l-1.336-1.488a6.136 6.136 0 0 1 3.991-1.553c1.488.007 2.833.596 3.786 1.658c.942 1.05 1.387 2.537 1.221 4.081c-.175 1.63-1.015 3.1-2.306 4.03m-35.661 0c-1.29-.932-2.131-2.401-2.306-4.031c-.166-1.543.279-3.03 1.221-4.079c.953-1.062 2.297-1.651 3.785-1.658h.027c1.441 0 2.849.551 3.965 1.553l-1.336 1.488c-.753-.676-1.689-1.005-2.646-1.041c-.916.004-1.735.357-2.306.994c-.563.628-.826 1.55-.721 2.53c.115 1.067.657 2.023 1.488 2.624zM25 16.84h-2c0-2.885 0-10.548 4.979-10.548c2.154 0 3.193 1.211 3.952 2.096c.629.734.961 1.086 1.616 1.086h1.37v2h-1.37c-1.604 0-2.453-.99-3.135-1.785c-.67-.781-1.198-1.398-2.434-1.398C25.975 8.292 25 11.088 25 16.84"/><path d="M24.085 16.95c9.421 0 17.085 5.231 17.085 11.661c0 6.431-7.664 11.662-17.085 11.662S7 35.042 7 28.611c0-6.43 7.664-11.661 17.085-11.661m0-2C13.544 14.95 5 21.066 5 28.611c0 7.546 8.545 13.662 19.085 13.662c10.54 0 19.085-6.116 19.085-13.662c0-7.545-8.545-13.661-19.085-13.661M38.32 7.975c1.379 0 2.5 1.122 2.5 2.5s-1.121 2.5-2.5 2.5s-2.5-1.122-2.5-2.5s1.121-2.5 2.5-2.5m0-2a4.501 4.501 0 1 0 .002 9.002a4.501 4.501 0 0 0-.002-9.002"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Redo(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#00BCD4"><path d="M43 18L29 6.3v23.4z"/><path d="M20 14h12v8H20c-2.8 0-5 2.2-5 5s2.2 5 5 5h3v8h-3c-7.2 0-13-5.8-13-13s5.8-13 13-13"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Refresh(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#1565C0"><path d="M13 13c0-3.3 2.7-6 6-6h10c3.3 0 6 2.7 6 6h4c0-5.5-4.5-10-10-10H19C13.5 3 9 7.5 9 13v11.2h4z"/><path d="m4.6 22l6.4 8.4l6.4-8.4z"/></g><g fill="#1565C0"><path d="M35 35c0 3.3-2.7 6-6 6H19c-3.3 0-6-2.7-6-6H9c0 5.5 4.5 10 10 10h10c5.5 0 10-4.5 10-10V23h-4z"/><path d="m30.6 26l6.4-8.4l6.4 8.4z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RegisteredTrademark(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="24" cy="24" r="21" fill="#9C27B0"/><path fill="#E1BEE7" d="M25 26.8h-4.5v9h-4V12.5h8.2c1.3 0 2.5.2 3.6.5c1 .3 1.9.8 2.6 1.3c.7.6 1.3 1.3 1.6 2.2s.6 1.9.6 3c0 1.6-.4 2.9-1.1 3.9c-.8 1-1.8 1.9-3.1 2.4l5.2 9.7v.2h-4.3zm-4.5-3.2h4.2c.7 0 1.4-.1 1.9-.3c.5-.2 1-.5 1.4-.8c.4-.3.6-.7.8-1.2c.2-.5.3-1 .3-1.6c0-.6-.1-1.1-.3-1.6c-.2-.5-.4-.9-.8-1.2c-.4-.3-.8-.6-1.4-.8c-.5-.2-1.2-.3-2-.3h-4.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RemoveImage(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#8CBCD6" d="M40 41H8c-2.2 0-4-1.8-4-4V11c0-2.2 1.8-4 4-4h32c2.2 0 4 1.8 4 4v26c0 2.2-1.8 4-4 4"/><circle cx="35" cy="16" r="3" fill="#B3DDF5"/><path fill="#9AC9E3" d="M20 16L9 32h22z"/><path fill="#B3DDF5" d="m31 22l-8 10h16z"/><circle cx="38" cy="38" r="10" fill="#F44336"/><g fill="#fff"><path d="m43.31 41.181l-2.12 2.122l-8.485-8.484l2.121-2.122z"/><path d="m34.819 43.31l-2.122-2.12l8.484-8.485l2.122 2.121z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reuse(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#455A64" d="M12.1 42h5.1l-.7-23.8l-5.6 2z"/><circle cx="36.5" cy="10" r="5" fill="#FFB74D"/><path fill="#607D8B" d="M11 42H6l1.8-23.4l6.4 2.3z"/><path fill="#607D8B" d="M31.7 15.9c-.6-2-1.3-4-2.5-5.8c-1.3-1.6-3.2-3.1-6.1-2c-3.1 1.3-9.2 3.6-11.2 4.5c-2.3 1.1-4.1 2.7-4.1 5.9c0 3.4 4.3 5.3 4.3 5.3l14.7-6.1l1.7 4.5l5.3.1c0 .1-1.5-4.4-2.1-6.4"/><path fill="#B39DDB" d="M37.9 42H30c-1 0-1.8-.7-2-1.7l-2.6-17.1h17l-2.6 17.1c0 1-.9 1.7-1.9 1.7"/><path fill="#7E57C2" d="M42 24H26c-.6 0-1-.4-1-1s.4-1 1-1h16c.6 0 1 .4 1 1s-.4 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Right(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#3F51B5"><path d="M44 24L30 35.7V12.3z"/><path d="M6 20h27v8H6z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightDown(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#3F51B5" d="M41 41H23l18-18z"/><path fill="#3F51B5" d="m6.983 12.607l5.656-5.656l23.119 23.12l-5.656 5.655z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightDownTwo(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#3F51B5" d="M29 44L17.3 30h23.4z"/><path fill="#3F51B5" d="M21 6H8v8h13c2.2 0 4 1.8 4 4v17h8V18c0-6.6-5.4-12-12-12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightUp(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#3F51B5" d="M41 7v18L23 7z"/><path fill="#3F51B5" d="m12.641 40.983l-5.656-5.656l23.12-23.119l5.655 5.656z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightUpTwo(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#3F51B5" d="M29 4L17.3 18h23.4z"/><path fill="#3F51B5" d="M21 42H8v-8h13c2.2 0 4-1.8 4-4V13h8v17c0 6.6-5.4 12-12 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RotateCamera(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#5E35B1"><path d="M33.9 12.1H14.2L17.6 7c.4-.6 1-.9 1.7-.9h9.6c.7 0 1.3.3 1.7.9zM14 11H8V9.2C8 8.5 8.5 8 9.2 8h3.6c.7 0 1.2.5 1.2 1.2z"/><path d="M40 42H8c-2.2 0-4-1.8-4-4V14c0-2.2 1.8-4 4-4h32c2.2 0 4 1.8 4 4v24c0 2.2-1.8 4-4 4"/></g><path fill="#E8EAF6" d="M34 25c0-5.5-4.5-10-10-10s-10 4.5-10 10s4.5 10 10 10v-2c-4.4 0-8-3.6-8-8s3.6-8 8-8s8 3.6 8 8h-3.5l4.5 5.6l4.5-5.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RotateToLandscape(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#37474F" d="M10 41V7c0-2.2 1.8-4 4-4h20c2.2 0 4 1.8 4 4v34c0 2.2-1.8 4-4 4H14c-2.2 0-4-1.8-4-4"/><path fill="#F3E5F5" d="M34 6H14c-.6 0-1 .4-1 1v34c0 .6.4 1 1 1h20c.6 0 1-.4 1-1V7c0-.6-.4-1-1-1"/><path fill="#9C27B0" d="m22 34l5.9-7H16.1z"/><path fill="#9C27B0" d="M26 16c-3.3 0-6 2.7-6 6v6h4v-6c0-1.1.9-2 2-2s2 .9 2 2v2h4v-2c0-3.3-2.7-6-6-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RotateToPortrait(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#37474F" d="M41 38H7c-2.2 0-4-1.8-4-4V14c0-2.2 1.8-4 4-4h34c2.2 0 4 1.8 4 4v20c0 2.2-1.8 4-4 4"/><path fill="#F3E5F5" d="M6 14v20c0 .6.4 1 1 1h34c.6 0 1-.4 1-1V14c0-.6-.4-1-1-1H7c-.6 0-1 .4-1 1"/><path fill="#9C27B0" d="m26 15l-5.9 7h11.8z"/><path fill="#9C27B0" d="M24 21v6c0 1.1-.9 2-2 2s-2-.9-2-2v-2h-4v2c0 3.3 2.7 6 6 6s6-2.7 6-6v-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ruler(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#FFA000" d="M13.324 45.003L3.002 34.68L34.676 3.007L44.998 13.33z"/><path fill="#9E6400" d="m22.803 24.188l-4.666-4.666l1.414-1.414l4.666 4.666zm2.01-5.99l-2.616-2.616l1.414-1.414l2.616 2.616zm5.991-2.01l-4.666-4.666l1.414-1.414l4.666 4.666zm-.649-8.645l1.415-1.414l2.615 2.616l-1.414 1.414zM8.81 34.198l-2.616-2.616l1.414-1.414l2.616 2.616zm5.991-2.01l-4.666-4.666l1.414-1.414l4.666 4.666zm2.011-5.99l-2.616-2.616l1.414-1.414l2.616 2.616z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rules(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#42A5F5" d="M39 45H9s-3-.1-3-8h36c0 7.9-3 8-3 8"/><path fill="#90CAF9" d="M8 3h32v34H8z"/><path fill="#1976D2" d="M18 15h16v2H18zm0 4h16v2H18zm0 4h16v2H18zm0 4h16v2H18zm0 4h16v2H18zm-4-16h2v2h-2zm0 4h2v2h-2zm0 4h2v2h-2zm0 4h2v2h-2zm0 4h2v2h-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Safe(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#455A64" d="M8 39h6v3H8zm26 0h6v3h-6z"/><path fill="#78909C" d="M40 41H8c-2.2 0-4-1.8-4-4V11c0-2.2 1.8-4 4-4h32c2.2 0 4 1.8 4 4v26c0 2.2-1.8 4-4 4"/><path fill="#90A4AE" d="M40 38H8c-.6 0-1-.4-1-1V11c0-.6.4-1 1-1h32c.6 0 1 .4 1 1v26c0 .6-.4 1-1 1"/><path fill="#37474F" d="M29 14c-5.5 0-10 4.5-10 10s4.5 10 10 10s10-4.5 10-10s-4.5-10-10-10m0 17c-3.9 0-7-3.1-7-7s3.1-7 7-7s7 3.1 7 7s-3.1 7-7 7"/><path fill="#B0BEC5" d="m35.3 19.1l.4-.4c.4-.4.4-1 0-1.4s-1-.4-1.4 0l-.4.4c.5.4 1 .9 1.4 1.4m-12.6 0c.4-.5.9-1 1.4-1.4l-.4-.4c-.4-.4-1-.4-1.4 0s-.4 1 0 1.4zM21 24c0-.3 0-.7.1-1h-.6c-.6 0-1 .4-1 1s.4 1 1 1h.6c-.1-.3-.1-.7-.1-1m8-8c.3 0 .7 0 1 .1v-.6c0-.6-.4-1-1-1s-1 .4-1 1v.6c.3-.1.7-.1 1-.1m6.3 12.9c-.4.5-.9 1-1.4 1.4l.4.4c.2.2.5.3.7.3s.5-.1.7-.3c.4-.4.4-1 0-1.4zm-12.6 0l-.4.4c-.4.4-.4 1 0 1.4c.2.2.5.3.7.3s.5-.1.7-.3l.4-.4c-.5-.4-1-.9-1.4-1.4M37.5 23h-.6c0 .3.1.7.1 1s0 .7-.1 1h.6c.6 0 1-.4 1-1s-.4-1-1-1M29 32c-.3 0-.7 0-1-.1v.6c0 .6.4 1 1 1s1-.4 1-1v-.6c-.3.1-.7.1-1 .1"/><path fill="#455A64" d="M12 20c-1.1 0-2 .9-2 2v8c0 1.1.9 2 2 2s2-.9 2-2v-8c0-1.1-.9-2-2-2"/><path fill="#CFD8DC" d="M12 18c-1.1 0-2 .9-2 2v8c0 1.1.9 2 2 2s2-.9 2-2v-8c0-1.1-.9-2-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SalesPerformance(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#FFA000"><path d="M38 13c-3.3 0-6-.9-6-2v2c0 1.1 2.7 2 6 2s6-.9 6-2v-2c0 1.1-2.7 2-6 2"/><path d="M38 10c-3.3 0-6-.9-6-2v2c0 1.1 2.7 2 6 2s6-.9 6-2V8c0 1.1-2.7 2-6 2m0 6c-3.3 0-6-.9-6-2v2c0 1.1 2.7 2 6 2s6-.9 6-2v-2c0 1.1-2.7 2-6 2"/><path d="M38 19c-3.3 0-6-.9-6-2v2c0 1.1 2.7 2 6 2s6-.9 6-2v-2c0 1.1-2.7 2-6 2"/><path d="M38 22c-3.3 0-6-.9-6-2v2c0 1.1 2.7 2 6 2s6-.9 6-2v-2c0 1.1-2.7 2-6 2"/><path d="M38 25c-3.3 0-6-.9-6-2v2c0 1.1 2.7 2 6 2s6-.9 6-2v-2c0 1.1-2.7 2-6 2"/><path d="M38 28c-3.3 0-6-.9-6-2v2c0 1.1 2.7 2 6 2s6-.9 6-2v-2c0 1.1-2.7 2-6 2"/><path d="M38 31c-3.3 0-6-.9-6-2v2c0 1.1 2.7 2 6 2s6-.9 6-2v-2c0 1.1-2.7 2-6 2"/><path d="M38 34c-3.3 0-6-.9-6-2v2c0 1.1 2.7 2 6 2s6-.9 6-2v-2c0 1.1-2.7 2-6 2"/><path d="M38 37c-3.3 0-6-.9-6-2v2c0 1.1 2.7 2 6 2s6-.9 6-2v-2c0 1.1-2.7 2-6 2"/><path d="M38 40c-3.3 0-6-.9-6-2v2c0 1.1 2.7 2 6 2s6-.9 6-2v-2c0 1.1-2.7 2-6 2"/></g><g fill="#FFC107"><ellipse cx="38" cy="8" rx="6" ry="2"/><path d="M38 12c-2.8 0-5.1-.6-5.8-1.5c-.1.2-.2.3-.2.5c0 1.1 2.7 2 6 2s6-.9 6-2c0-.2-.1-.3-.2-.5c-.7.9-3 1.5-5.8 1.5m0 3c-2.8 0-5.1-.6-5.8-1.5c-.1.2-.2.3-.2.5c0 1.1 2.7 2 6 2s6-.9 6-2c0-.2-.1-.3-.2-.5c-.7.9-3 1.5-5.8 1.5m0 3c-2.8 0-5.1-.6-5.8-1.5c-.1.2-.2.3-.2.5c0 1.1 2.7 2 6 2s6-.9 6-2c0-.2-.1-.3-.2-.5c-.7.9-3 1.5-5.8 1.5m0 3c-2.8 0-5.1-.6-5.8-1.5c-.1.2-.2.3-.2.5c0 1.1 2.7 2 6 2s6-.9 6-2c0-.2-.1-.3-.2-.5c-.7.9-3 1.5-5.8 1.5m0 3c-2.8 0-5.1-.6-5.8-1.5c-.1.2-.2.3-.2.5c0 1.1 2.7 2 6 2s6-.9 6-2c0-.2-.1-.3-.2-.5c-.7.9-3 1.5-5.8 1.5m0 3c-2.8 0-5.1-.6-5.8-1.5c-.1.2-.2.3-.2.5c0 1.1 2.7 2 6 2s6-.9 6-2c0-.2-.1-.3-.2-.5c-.7.9-3 1.5-5.8 1.5m0 3c-2.8 0-5.1-.6-5.8-1.5c-.1.2-.2.3-.2.5c0 1.1 2.7 2 6 2s6-.9 6-2c0-.2-.1-.3-.2-.5c-.7.9-3 1.5-5.8 1.5m0 3c-2.8 0-5.1-.6-5.8-1.5c-.1.2-.2.3-.2.5c0 1.1 2.7 2 6 2s6-.9 6-2c0-.2-.1-.3-.2-.5c-.7.9-3 1.5-5.8 1.5m0 3c-2.8 0-5.1-.6-5.8-1.5c-.1.2-.2.3-.2.5c0 1.1 2.7 2 6 2s6-.9 6-2c0-.2-.1-.3-.2-.5c-.7.9-3 1.5-5.8 1.5m0 3c-2.8 0-5.1-.6-5.8-1.5c-.1.2-.2.3-.2.5c0 1.1 2.7 2 6 2s6-.9 6-2c0-.2-.1-.3-.2-.5c-.7.9-3 1.5-5.8 1.5"/></g><g fill="#FFA000"><path d="M10 19c-3.3 0-6-.9-6-2v2c0 1.1 2.7 2 6 2s6-.9 6-2v-2c0 1.1-2.7 2-6 2"/><path d="M10 16c-3.3 0-6-.9-6-2v2c0 1.1 2.7 2 6 2s6-.9 6-2v-2c0 1.1-2.7 2-6 2m0 6c-3.3 0-6-.9-6-2v2c0 1.1 2.7 2 6 2s6-.9 6-2v-2c0 1.1-2.7 2-6 2"/><path d="M10 25c-3.3 0-6-.9-6-2v2c0 1.1 2.7 2 6 2s6-.9 6-2v-2c0 1.1-2.7 2-6 2"/><path d="M10 28c-3.3 0-6-.9-6-2v2c0 1.1 2.7 2 6 2s6-.9 6-2v-2c0 1.1-2.7 2-6 2"/><path d="M10 31c-3.3 0-6-.9-6-2v2c0 1.1 2.7 2 6 2s6-.9 6-2v-2c0 1.1-2.7 2-6 2"/><path d="M10 34c-3.3 0-6-.9-6-2v2c0 1.1 2.7 2 6 2s6-.9 6-2v-2c0 1.1-2.7 2-6 2"/><path d="M10 37c-3.3 0-6-.9-6-2v2c0 1.1 2.7 2 6 2s6-.9 6-2v-2c0 1.1-2.7 2-6 2"/><path d="M10 40c-3.3 0-6-.9-6-2v2c0 1.1 2.7 2 6 2s6-.9 6-2v-2c0 1.1-2.7 2-6 2"/></g><g fill="#FFC107"><ellipse cx="10" cy="14" rx="6" ry="2"/><path d="M10 18c-2.8 0-5.1-.6-5.8-1.5c-.1.2-.2.3-.2.5c0 1.1 2.7 2 6 2s6-.9 6-2c0-.2-.1-.3-.2-.5c-.7.9-3 1.5-5.8 1.5m0 3c-2.8 0-5.1-.6-5.8-1.5c-.1.2-.2.3-.2.5c0 1.1 2.7 2 6 2s6-.9 6-2c0-.2-.1-.3-.2-.5c-.7.9-3 1.5-5.8 1.5m0 3c-2.8 0-5.1-.6-5.8-1.5c-.1.2-.2.3-.2.5c0 1.1 2.7 2 6 2s6-.9 6-2c0-.2-.1-.3-.2-.5c-.7.9-3 1.5-5.8 1.5m0 3c-2.8 0-5.1-.6-5.8-1.5c-.1.2-.2.3-.2.5c0 1.1 2.7 2 6 2s6-.9 6-2c0-.2-.1-.3-.2-.5c-.7.9-3 1.5-5.8 1.5m0 3c-2.8 0-5.1-.6-5.8-1.5c-.1.2-.2.3-.2.5c0 1.1 2.7 2 6 2s6-.9 6-2c0-.2-.1-.3-.2-.5c-.7.9-3 1.5-5.8 1.5m0 3c-2.8 0-5.1-.6-5.8-1.5c-.1.2-.2.3-.2.5c0 1.1 2.7 2 6 2s6-.9 6-2c0-.2-.1-.3-.2-.5c-.7.9-3 1.5-5.8 1.5m0 3c-2.8 0-5.1-.6-5.8-1.5c-.1.2-.2.3-.2.5c0 1.1 2.7 2 6 2s6-.9 6-2c0-.2-.1-.3-.2-.5c-.7.9-3 1.5-5.8 1.5m0 3c-2.8 0-5.1-.6-5.8-1.5c-.1.2-.2.3-.2.5c0 1.1 2.7 2 6 2s6-.9 6-2c0-.2-.1-.3-.2-.5c-.7.9-3 1.5-5.8 1.5"/></g><g fill="#FFA000"><path d="M24 28c-3.3 0-6-.9-6-2v2c0 1.1 2.7 2 6 2s6-.9 6-2v-2c0 1.1-2.7 2-6 2"/><path d="M24 25c-3.3 0-6-.9-6-2v2c0 1.1 2.7 2 6 2s6-.9 6-2v-2c0 1.1-2.7 2-6 2m0 6c-3.3 0-6-.9-6-2v2c0 1.1 2.7 2 6 2s6-.9 6-2v-2c0 1.1-2.7 2-6 2"/><path d="M24 34c-3.3 0-6-.9-6-2v2c0 1.1 2.7 2 6 2s6-.9 6-2v-2c0 1.1-2.7 2-6 2"/><path d="M24 37c-3.3 0-6-.9-6-2v2c0 1.1 2.7 2 6 2s6-.9 6-2v-2c0 1.1-2.7 2-6 2"/><path d="M24 40c-3.3 0-6-.9-6-2v2c0 1.1 2.7 2 6 2s6-.9 6-2v-2c0 1.1-2.7 2-6 2"/></g><g fill="#FFC107"><ellipse cx="24" cy="23" rx="6" ry="2"/><path d="M24 27c-2.8 0-5.1-.6-5.8-1.5c-.1.2-.2.3-.2.5c0 1.1 2.7 2 6 2s6-.9 6-2c0-.2-.1-.3-.2-.5c-.7.9-3 1.5-5.8 1.5m0 3c-2.8 0-5.1-.6-5.8-1.5c-.1.2-.2.3-.2.5c0 1.1 2.7 2 6 2s6-.9 6-2c0-.2-.1-.3-.2-.5c-.7.9-3 1.5-5.8 1.5m0 3c-2.8 0-5.1-.6-5.8-1.5c-.1.2-.2.3-.2.5c0 1.1 2.7 2 6 2s6-.9 6-2c0-.2-.1-.3-.2-.5c-.7.9-3 1.5-5.8 1.5m0 3c-2.8 0-5.1-.6-5.8-1.5c-.1.2-.2.3-.2.5c0 1.1 2.7 2 6 2s6-.9 6-2c0-.2-.1-.3-.2-.5c-.7.9-3 1.5-5.8 1.5m0 3c-2.8 0-5.1-.6-5.8-1.5c-.1.2-.2.3-.2.5c0 1.1 2.7 2 6 2s6-.9 6-2c0-.2-.1-.3-.2-.5c-.7.9-3 1.5-5.8 1.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScatterPlot(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#CFD8DC" d="M9 39V6H7v35h35v-2z"/><g fill="#00BCD4"><circle cx="39" cy="11" r="3"/><circle cx="31" cy="13" r="3"/><circle cx="37" cy="19" r="3"/><circle cx="34" cy="26" r="3"/><circle cx="28" cy="20" r="3"/><circle cx="26" cy="28" r="3"/><circle cx="20" cy="23" r="3"/><circle cx="21" cy="33" r="3"/><circle cx="14" cy="30" r="3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Search(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#616161"><path d="m29.175 31.99l2.828-2.827l12.019 12.019l-2.828 2.827z"/><circle cx="20" cy="20" r="16"/></g><path fill="#37474F" d="m32.45 35.34l2.827-2.828l8.696 8.696l-2.828 2.828z"/><circle cx="20" cy="20" r="13" fill="#64B5F6"/><path fill="#BBDEFB" d="M26.9 14.2c-1.7-2-4.2-3.2-6.9-3.2s-5.2 1.2-6.9 3.2c-.4.4-.3 1.1.1 1.4c.4.4 1.1.3 1.4-.1C16 13.9 17.9 13 20 13s4 .9 5.4 2.5c.2.2.5.4.8.4c.2 0 .5-.1.6-.2c.4-.4.4-1.1.1-1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SelfServiceKiosk(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#546E7A" d="M44 30H4V11c0-2.2 1.8-4 4-4h32c2.2 0 4 1.8 4 4z"/><path fill="#64B5F6" d="M40 27H8c-.6 0-1-.4-1-1V11c0-.6.4-1 1-1h32c.6 0 1 .4 1 1v15c0 .6-.4 1-1 1"/><path fill="#78909C" d="M40 41H8c-2.2 0-4-1.8-4-4v-7h40v7c0 2.2-1.8 4-4 4"/><g fill="#37474F"><path d="M27 34h12v2H27zM9 34h12v2H9z"/><path d="M18 35c0 1.1-1.3 2-3 2s-3-.9-3-2z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Selfie(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#FFB74D" d="M32.9 22c0-.3.1-.7.1-1v-7c0-7.6-18-5-18 0v7c0 .3 0 .7.1 1z"/><path fill="#37474F" d="M40 44H8c-2.2 0-4-1.8-4-4V26c0-2.2 1.8-4 4-4h32c2.2 0 4 1.8 4 4v14c0 2.2-1.8 4-4 4"/><path fill="#BBDEFB" d="M7 26v14c0 .6.4 1 1 1h29c.6 0 1-.4 1-1V26c0-.6-.4-1-1-1H8c-.6 0-1 .4-1 1"/><path fill="#78909C" d="M40 30h2v6h-2z"/><path fill="#BF360C" d="M19 32h8v9h-8z"/><path fill="#FF9800" d="M20.5 37.5h5V41h-5z"/><path fill="#FFB74D" d="M27.5 32c0-3.8-9-2.5-9 0v3.5c0 2.5 2 4.5 4.5 4.5s4.5-2 4.5-4.5z"/><g fill="#784719"><circle cx="28" cy="21" r="1"/><circle cx="20" cy="21" r="1"/><circle cx="25" cy="35.5" r=".5"/><circle cx="21" cy="35.5" r=".5"/></g><path fill="#FF5722" d="M23 27c-3 0-8 1.3-8 11l4 3v-6.5l6-3.5l2 2.5V41l4-3c0-2-.8-10-6-10l-.5-1zm-7-5v-3l12-7l4 5v5h6.8C38.3 15.8 36.1 6 28 6l-1-2h-3C18.5 4 10.7 6.8 9.2 22z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SerialTasks(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#90CAF9" d="M33 9H11v4h22c1.1 0 2 .9 2 2v20H23v4h16V15c0-3.3-2.7-6-6-6"/><path fill="#D81B60" d="M6 6h10v10H6z"/><g fill="#2196F3"><path d="M32 17h10v10H32zM16 32h10v10H16z"/><circle cx="26" cy="11" r="5"/><circle cx="37" cy="37" r="5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ServiceMark(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="24" cy="24" r="21" fill="#9C27B0"/><path fill="#E1BEE7" d="M16.7 28.2c0-3.8-7.3-2.2-7.3-8.1c0-.7.4-4.8 5.5-4.8s5.4 4.5 5.4 5.3h-3.5c0-.4 0-2.5-2-2.5c-1.8 0-1.9 1.7-1.9 2c0 3 7.4 2 7.4 8.1c0 2-1.1 4.8-5.3 4.8c-4.7 0-6-3.4-6-5.7h3.5c0 .5-.2 2.8 2.5 2.8c1.8.1 1.7-1.6 1.7-1.9m10.4-12.6L30.3 28l3.2-12.4H38v17.2h-3.5v-4.6l.3-7.2l-3.4 11.8H29L25.6 21l.3 7.2v4.6h-3.5V15.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Services(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#E65100" d="M25.6 34.4c.1-.4.1-.9.1-1.4s0-.9-.1-1.4l2.8-2c.3-.2.4-.6.2-.9l-2.7-4.6c-.2-.3-.5-.4-.8-.3L22 25.3c-.7-.6-1.5-1-2.4-1.4l-.3-3.4c0-.3-.3-.6-.6-.6h-5.3c-.3 0-.6.3-.6.6l-.4 3.5c-.9.3-1.6.8-2.4 1.4L6.9 24c-.3-.1-.7 0-.8.3l-2.7 4.6c-.2.3-.1.7.2.9l2.8 2c-.1.4-.1.9-.1 1.4s0 .9.1 1.4l-2.8 2c-.3.2-.4.6-.2.9l2.7 4.6c.2.3.5.4.8.3L10 41c.7.6 1.5 1 2.4 1.4l.3 3.4c0 .3.3.6.6.6h5.3c.3 0 .6-.3.6-.6l.3-3.4c.9-.3 1.6-.8 2.4-1.4l3.1 1.4c.3.1.7 0 .8-.3l2.7-4.6c.2-.3.1-.7-.2-.9zM16 38c-2.8 0-5-2.2-5-5s2.2-5 5-5s5 2.2 5 5s-2.2 5-5 5"/><path fill="#FFA000" d="M41.9 15.3c.1-.5.1-.9.1-1.3s0-.8-.1-1.3l2.5-1.8c.3-.2.3-.5.2-.8l-2.5-4.3c-.2-.3-.5-.4-.8-.2l-2.9 1.3c-.7-.5-1.4-.9-2.2-1.3l-.3-3.1c.1-.3-.1-.5-.4-.5h-4.9c-.3 0-.6.2-.6.5l-.3 3.1c-.8.3-1.5.7-2.2 1.3l-2.9-1.3c-.3-.1-.6 0-.8.2l-2.5 4.3c-.2.3-.1.6.2.8l2.5 1.8V14c0 .4 0 .8.1 1.3l-2.5 1.8c-.3.2-.3.5-.2.8l2.5 4.3c.2.3.5.4.8.2l2.9-1.3c.7.5 1.4.9 2.2 1.3l.3 3.1c0 .3.3.5.6.5h4.9c.3 0 .6-.2.6-.5l.3-3.1c.8-.3 1.5-.7 2.2-1.3l2.9 1.3c.3.1.6 0 .8-.2l2.5-4.3c.2-.3.1-.6-.2-.8zM33 19c-2.8 0-5-2.2-5-5s2.2-5 5-5s5 2.2 5 5s-2.2 5-5 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Settings(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#607D8B" d="M39.6 27.2c.1-.7.2-1.4.2-2.2s-.1-1.5-.2-2.2l4.5-3.2c.4-.3.6-.9.3-1.4L40 10.8c-.3-.5-.8-.7-1.3-.4l-5 2.3c-1.2-.9-2.4-1.6-3.8-2.2L29.4 5c-.1-.5-.5-.9-1-.9h-8.6c-.5 0-1 .4-1 .9l-.5 5.5c-1.4.6-2.7 1.3-3.8 2.2l-5-2.3c-.5-.2-1.1 0-1.3.4l-4.3 7.4c-.3.5-.1 1.1.3 1.4l4.5 3.2c-.1.7-.2 1.4-.2 2.2s.1 1.5.2 2.2L4 30.4c-.4.3-.6.9-.3 1.4L8 39.2c.3.5.8.7 1.3.4l5-2.3c1.2.9 2.4 1.6 3.8 2.2l.5 5.5c.1.5.5.9 1 .9h8.6c.5 0 1-.4 1-.9l.5-5.5c1.4-.6 2.7-1.3 3.8-2.2l5 2.3c.5.2 1.1 0 1.3-.4l4.3-7.4c.3-.5.1-1.1-.3-1.4zM24 35c-5.5 0-10-4.5-10-10s4.5-10 10-10s10 4.5 10 10s-4.5 10-10 10"/><path fill="#455A64" d="M24 13c-6.6 0-12 5.4-12 12s5.4 12 12 12s12-5.4 12-12s-5.4-12-12-12m0 17c-2.8 0-5-2.2-5-5s2.2-5 5-5s5 2.2 5 5s-2.2 5-5 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Share(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#1976D2" d="M38.1 31.2L19.4 24l18.7-7.2c1.5-.6 2.3-2.3 1.7-3.9c-.6-1.5-2.3-2.3-3.9-1.7l-26 10C8.8 21.6 8 22.8 8 24s.8 2.4 1.9 2.8l26 10c.4.1.7.2 1.1.2c1.2 0 2.3-.7 2.8-1.9c.6-1.6-.2-3.3-1.7-3.9"/><g fill="#1E88E5"><circle cx="11" cy="24" r="7"/><circle cx="37" cy="14" r="7"/><circle cx="37" cy="34" r="7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shipped(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#8BC34A" d="M43 36H29V14h10.6c.9 0 1.6.6 1.9 1.4L45 26v8c0 1.1-.9 2-2 2"/><path fill="#388E3C" d="M29 36H5c-1.1 0-2-.9-2-2V9c0-1.1.9-2 2-2h22c1.1 0 2 .9 2 2z"/><g fill="#37474F"><circle cx="37" cy="36" r="5"/><circle cx="13" cy="36" r="5"/></g><g fill="#78909C"><circle cx="37" cy="36" r="2"/><circle cx="13" cy="36" r="2"/></g><path fill="#37474F" d="M41 25h-7c-.6 0-1-.4-1-1v-7c0-.6.4-1 1-1h5.3c.4 0 .8.3.9.7l1.7 5.2c0 .1.1.2.1.3V24c0 .6-.4 1-1 1"/><path fill="#DCEDC8" d="m21.8 13.8l-7.9 7.9l-3.7-3.8L8 20.1l5.9 5.9L24 15.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shop(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#CFD8DC" d="M5 19h38v19H5z"/><path fill="#B0BEC5" d="M5 38h38v4H5z"/><path fill="#455A64" d="M27 24h12v18H27z"/><path fill="#E3F2FD" d="M9 24h14v11H9z"/><path fill="#1E88E5" d="M10 25h12v9H10z"/><path fill="#90A4AE" d="M36.5 33.5c-.3 0-.5.2-.5.5v2c0 .3.2.5.5.5s.5-.2.5-.5v-2c0-.3-.2-.5-.5-.5"/><g fill="#558B2F"><circle cx="24" cy="19" r="3"/><circle cx="36" cy="19" r="3"/><circle cx="12" cy="19" r="3"/></g><path fill="#7CB342" d="M40 6H8c-1.1 0-2 .9-2 2v3h36V8c0-1.1-.9-2-2-2m-19 5h6v8h-6zm16 0h-5l1 8h6zm-26 0h5l-1 8H9z"/><g fill="#FFA000"><circle cx="30" cy="19" r="3"/><path d="M45 19c0 1.7-1.3 3-3 3s-3-1.3-3-3s1.3-3 3-3z"/><circle cx="18" cy="19" r="3"/><path d="M3 19c0 1.7 1.3 3 3 3s3-1.3 3-3s-1.3-3-3-3z"/></g><path fill="#FFC107" d="M32 11h-5v8h6zm10 0h-5l2 8h6zm-26 0h5v8h-6zM6 11h5l-2 8H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Signature(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#1565C0" d="M38.8 28.2C41.5 24.8 45 20.1 45 12c0-.6-.4-1-1-1s-1 .4-1 1c0 6.7-2.5 10.7-5 13.9c-.6-1.9-1-4.2-1-6.9c0-.5-.4-1-1-1c-.5 0-1 .4-1 1c-.1 1.7-.6 3.6-1 3.8c-.4 0-.9-1.4-1-2.8c0-.5-.5-.9-1-.9s-1 .3-1 .9c-.3 1.7-1.1 4.1-2 4.1c-.4 0-.6-.1-.7-.3c-.3-.3-.4-1-.4-1.6c0-.4.1-.8.1-1.2c0-.5-.4-1-.9-1s-1 .3-1.1.8c0 .1-.1.5-.1 1.1c-.2 1.7-.8 5.1-2.9 5.1c-.7 0-1.1-.2-1.4-.7c-.5-.8-.5-2.1 0-3.3v-.1c.1-.1.1-.3.2-.4c.8-1.6 1.7-2.5 3.2-2.5c.6 0 1-.4 1-1s-.4-1-1-1c-4.2 0-5.4 4.1-6.6 8c-1.4 4.8-2.7 8-6.4 8c-5.1 0-7-6.6-7-11c0-8.6 4.7-14 9-14c2.9 0 4 2.3 4.1 2.4c.2.5.8.7 1.3.5c.5-.2.7-.8.5-1.3c-.1-.2-1.7-3.6-5.9-3.6C8.6 7 3 13 3 23c0 10.3 5.9 13 9 13c5.1 0 6.8-4.5 8.1-8.5c.7.9 1.7 1.5 2.9 1.5c2.2 0 3.5-1.6 4.2-3.6c.5.4 1.1.6 1.8.6c1.4 0 2.4-1.2 3-2.4c.4.7 1.1 1.2 2 1.2c.6 0 1.1-.3 1.5-.7c.3 1.4.7 2.7 1 3.8c-1.4 1.8-2.5 3.3-2.5 5.1c0 1.7 1.3 3 3 3c1.8 0 3-1.6 3-3c0-1.3-.5-2.7-1.1-4.3c0-.2-.1-.3-.1-.5M37 34c-.7 0-1-.5-1-1c0-.9.5-1.8 1.3-2.9c.4 1.2.7 2.1.7 2.9c0 .3-.3 1-1 1"/><path fill="#90A4AE" d="M3 40h42v2H3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SimCard(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#009688" d="M36 45H12c-2.2 0-4-1.8-4-4V7c0-2.2 1.8-4 4-4h16.3c1.1 0 2.1.4 2.8 1.2l7.7 7.7c.8.8 1.2 1.8 1.2 2.8V41c0 2.2-1.8 4-4 4"/><path fill="#FF9800" d="M32 38H16c-1.1 0-2-.9-2-2V24c0-1.1.9-2 2-2h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2"/><path fill="#FFD54F" d="M29 30v3h5v2h-5v3h-2V22h2v6h5v2zm-15-1v2h5v2h-5v2h5v3h2v-9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SimCardChip(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#FF9800" d="M5 35V13c0-2.2 1.8-4 4-4h30c2.2 0 4 1.8 4 4v22c0 2.2-1.8 4-4 4H9c-2.2 0-4-1.8-4-4"/><path fill="#FFD54F" d="M43 21v-2H31c-1.1 0-2-.9-2-2s.9-2 2-2h1v-2h-1c-2.2 0-4 1.8-4 4s1.8 4 4 4h3v6h-3c-2.8 0-5 2.2-5 5s2.2 5 5 5h2v-2h-2c-1.7 0-3-1.3-3-3s1.3-3 3-3h12v-2h-7v-6zm-26 6h-3v-6h3c2.2 0 4-1.8 4-4s-1.8-4-4-4h-3v2h3c1.1 0 2 .9 2 2s-.9 2-2 2H5v2h7v6H5v2h12c1.7 0 3 1.3 3 3s-1.3 3-3 3h-2v2h2c2.8 0 5-2.2 5-5s-2.2-5-5-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SlrBackSide(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#5E35B1" d="M40 10h-7.6l-2-3c-.4-.6-1-.9-1.7-.9h-9.6c-.7 0-1.3.3-1.7.9l-2 3H8c-2.2 0-4 1.8-4 4v24c0 2.2 1.8 4 4 4h32c2.2 0 4-1.8 4-4V14c0-2.2-1.8-4-4-4"/><path fill="#F57C00" d="M11 16h20c.6 0 1 .4 1 1v16c0 .6-.4 1-1 1H11c-.6 0-1-.4-1-1V17c0-.6.4-1 1-1"/><path fill="#942A09" d="M18.9 22L12 32h13.8z"/><circle cx="27" cy="21" r="2" fill="#FFF9C4"/><path fill="#BF360C" d="m25.2 26l-4.8 6H30z"/><g fill="#8667C4"><path d="M34 10h6v-.8c0-.7-.5-1.2-1.2-1.2h-3.6c-.7 0-1.2.5-1.2 1.2z"/><circle cx="38" cy="18" r="2"/><circle cx="38" cy="24" r="2"/><circle cx="38" cy="30" r="2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmartphoneTablet(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#37474F" d="M4 39V7c0-2.2 1.8-4 4-4h22c2.2 0 4 1.8 4 4v32c0 2.2-1.8 4-4 4H8c-2.2 0-4-1.8-4-4"/><path fill="#BBDEFB" d="M30 6H8c-.6 0-1 .4-1 1v29c0 .6.4 1 1 1h22c.6 0 1-.4 1-1V7c0-.6-.4-1-1-1"/><path fill="#78909C" d="M15 39h6v2h-6z"/><path fill="#E38939" d="M24 41V17c0-2.2 1.8-4 4-4h12c2.2 0 4 1.8 4 4v24c0 2.2-1.8 4-4 4H28c-2.2 0-4-1.8-4-4"/><path fill="#FFF3E0" d="M40 16H28c-.6 0-1 .4-1 1v22c0 .6.4 1 1 1h12c.6 0 1-.4 1-1V17c0-.6-.4-1-1-1"/><circle cx="34" cy="42.5" r="1.5" fill="#A6642A"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sms(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#009688" d="M37 39H11l-6 6V11c0-3.3 2.7-6 6-6h26c3.3 0 6 2.7 6 6v22c0 3.3-2.7 6-6 6"/><g fill="#fff"><circle cx="24" cy="22" r="3"/><circle cx="34" cy="22" r="3"/><circle cx="14" cy="22" r="3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SoundRecordingCopyright(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="24" cy="24" r="21" fill="#9C27B0"/><path fill="#E1BEE7" d="M20.7 27.2v8.4h-3.9V12.9h8.7c1.3 0 2.5.2 3.5.5c1 .4 1.9.9 2.6 1.5c.7.6 1.2 1.4 1.6 2.3c.4.9.6 1.8.6 2.9c0 1.1-.2 2.1-.6 3c-.4.9-.9 1.6-1.6 2.2c-.7.6-1.6 1.1-2.6 1.4c-1 .3-2.2.5-3.5.5zm0-3.2h4.7c.8 0 1.4-.1 2-.3c.5-.2 1-.5 1.4-.8c.4-.3.6-.8.8-1.2c.2-.5.2-1 .2-1.6c0-.5-.1-1-.2-1.5c-.2-.5-.4-.9-.8-1.3s-.8-.7-1.4-.9c-.5-.2-1.2-.3-2-.3h-4.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Speaker(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#81D4FA" d="M28 7.1v2c7.3 1 13 7.3 13 14.9s-5.7 13.9-13 14.9v2c8.4-1 15-8.2 15-16.9S36.4 8.1 28 7.1"/><path fill="#546E7A" d="M14 32H7c-1.1 0-2-.9-2-2V18c0-1.1.9-2 2-2h7z"/><path fill="#78909C" d="M26 42L14 32V16L26 6z"/><path fill="#03A9F4" d="M28 17.3v2.1c1.8.8 3 2.5 3 4.6s-1.2 3.8-3 4.6v2.1c2.9-.9 5-3.5 5-6.7s-2.1-5.8-5-6.7"/><path fill="#4FC3F7" d="M28 12.2v2c4.6.9 8 5 8 9.8s-3.4 8.9-8 9.8v2c5.7-1 10-5.9 10-11.8s-4.3-10.9-10-11.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SportsMode(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="28" cy="9" r="5" fill="#FF9800"/><path fill="#00796B" d="m29 27.3l-9.2-4.1c-1-.5-1.5 1-2 2s-4.1 7.2-3.8 8.3c.3.9 1.1 1.4 1.9 1.4c.2 0 .4 0 .6-.1L28.8 31c.8-.2 1.4-1 1.4-1.8s-.5-1.6-1.2-1.9"/><path fill="#009688" d="m26.8 15.2l-2.2-1c-1.3-.6-2.9 0-3.5 1.3L9.2 41.1c-.5 1 0 2.2 1 2.7c.3.1.6.2.9.2c.8 0 1.5-.4 1.8-1.1c0 0 9.6-13.3 10.4-14.9s4.9-9.3 4.9-9.3c.5-1.3 0-2.9-1.4-3.5"/><path fill="#FF9800" d="M40.5 15.7c-.7-.8-2-1-2.8-.3l-5 4.2l-6.4-3.5c-1.1-.6-2.6-.4-3.3.9c-.8 1.3-.4 2.9.8 3.4l8.3 3.4c.3.1.6.2.9.2c.5 0 .9-.2 1.3-.5l6-5c.8-.7.9-1.9.2-2.8m-28.8 7.4l3.4-5.1l4.6.6l1.5-3.1c.4-.9 1.2-1.4 2.1-1.5h-9.2c-.7 0-1.3.3-1.7.9l-4 6c-.6.9-.4 2.2.6 2.8c.2.2.6.3 1 .3c.6 0 1.3-.3 1.7-.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StackOfPhotos(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#64B5F6" d="m17.474 8.578l26.544 8.904l-8.904 26.544L8.57 35.122z"/><path fill="#1E88E5" d="m19.238 12.504l20.922 6.82l-6.2 19.02l-20.922-6.82z"/><path fill="#90CAF9" d="m10.881 5.778l27.524 5.068l-5.068 27.524l-27.524-5.068z"/><path fill="#42A5F5" d="m13.219 9.444l21.67 3.85l-3.5 19.7l-21.67-3.85z"/><path fill="#BBDEFB" d="M4 4h28v28H4z"/><path fill="#4CAF50" d="M7 7h22v20H7z"/><path fill="#fff" d="M16 13c0-1.1.9-2 2-2s2 .9 2 2s-2 4-2 4s-2-2.9-2-4m4 8c0 1.1-.9 2-2 2s-2-.9-2-2s2-4 2-4s2 2.9 2 4"/><path fill="#fff" d="M13.5 16.7c-1-.6-1.3-1.8-.7-2.7c.6-1 1.8-1.3 2.7-.7c1 .6 2.5 3.7 2.5 3.7s-3.5.3-4.5-.3m9 .6c1 .6 1.3 1.8.7 2.7c-.6 1-1.8 1.3-2.7.7c-1-.5-2.5-3.7-2.5-3.7s3.5-.3 4.5.3"/><path fill="#fff" d="M22.5 16.7c1-.6 1.3-1.8.7-2.7c-.6-1-1.8-1.3-2.7-.7c-1 .5-2.5 3.7-2.5 3.7s3.5.3 4.5-.3m-9 .6c-1 .6-1.3 1.8-.7 2.7c.6 1 1.8 1.3 2.7.7c1-.6 2.5-3.7 2.5-3.7s-3.5-.3-4.5.3"/><circle cx="18" cy="17" r="2" fill="#FFC107"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Start(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#F44336" d="M38 42H10c-2.2 0-4-1.8-4-4V10c0-2.2 1.8-4 4-4h28c2.2 0 4 1.8 4 4v28c0 2.2-1.8 4-4 4"/><path fill="#fff" d="m31 24l-11-8v16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Statistics(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#37474F"><path d="M23 5h2v36h-2z"/><path d="m25.817 32.772l1.414 1.414l-10.04 10.04l-1.414-1.414z"/><path d="m32.259 42.824l-1.414 1.414l-10.04-10.04l1.414-1.414z"/></g><path fill="#CFD8DC" d="M4 8h40v28H4z"/><g fill="#607D8B"><path d="M3 7h42v4H3zm0 28h42v2H3z"/><circle cx="31.5" cy="43.5" r="1.5"/><circle cx="16.5" cy="43.5" r="1.5"/></g><g fill="#C51162"><path d="m31.9 18.9l-5.9 6l-6-6l-8.1 8l2.2 2.2l5.9-6l6 6l8.1-8z"/><path d="m36 24l-7-7h7z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Steam(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#FFF" d="M42 38a4 4 0 0 1-4 4H10a4 4 0 0 1-4-4V10a4 4 0 0 1 4-4h28a4 4 0 0 1 4 4z"/><g fill="#455A64"><path d="M18.459 33.645c-.288 0-.56-.057-.822-.141l-.005.02l-3.67-1.062c.644 1.878 2.406 3.237 4.5 3.237c2.641 0 4.776-2.136 4.776-4.776s-2.135-4.777-4.776-4.777c-1.141 0-2.175.418-2.998 1.087L19 28.255c.029.007.055.018.084.024l.049.016v.002a2.715 2.715 0 0 1-.674 5.348m12.464-9.346a5.84 5.84 0 0 0 5.837-5.838a5.84 5.84 0 0 0-5.837-5.837a5.84 5.84 0 0 0-5.837 5.837a5.84 5.84 0 0 0 5.837 5.838m-.003-9.89a4.053 4.053 0 1 1 .005 8.107a4.053 4.053 0 0 1-.005-8.107"/><path d="M38 6H10a4 4 0 0 0-4 4v14.495l7.027 2.033c1.287-1.59 3.229-2.626 5.434-2.626c.07 0 .135.02.204.021l3.876-5.355c0-.035-.005-.072-.005-.105a8.387 8.387 0 1 1 8.387 8.385c-.044 0-.087-.006-.132-.007l-5.33 3.871c.002.07.021.14.021.211a7.02 7.02 0 0 1-7.021 7.021c-3.593 0-6.52-2.707-6.937-6.188L6 30.158V38a4 4 0 0 0 4 4h28a4 4 0 0 0 4-4V10a4 4 0 0 0-4-4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stumbleupon(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#E64A19" d="M24.001 5c-10.494 0-19 8.506-19 19c0 10.493 8.506 19 19 19c10.493 0 19-8.507 19-19c0-10.494-8.507-19-19-19"/><path fill="#FFF" d="M24.001 19c-.003 0 .003 0 0 0c-.062-.004-1 0-1 1v7.876C22.916 29.888 21.504 33 17.959 33c-3.607 0-4.958-3.065-4.958-4.958V24h4v4c.038.709.629 1 1 1c.665 0 .972-.361 1-1v-8.124c0-2.01 1.332-5 5-5c.045 0 .086.006.131.007c0 0 4.869-.009 4.869 5.117c0 1.104-.896 1.876-2 1.876s-2-.771-2-1.876c0-.876-.96-.997-1-1m11 8.876c0 2.01-1.331 5.124-5 5.124s-5-3.114-5-5.124v-3.439a4.917 4.917 0 0 0 2 .439a4.95 4.95 0 0 0 2-.424V28c.038 1 .663 1 1 1c.247 0 1 0 1-1v-4h4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Support(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#607D8B" d="M44.7 11L36 19.6s-2.6 0-5.2-2.6s-2.6-5.2-2.6-5.2l8.7-8.7c-4.9-1.2-10.8.4-14.4 4c-5.4 5.4-.6 12.3-2 13.7C12.9 28.7 5.1 34.7 4.9 35c-2.3 2.3-2.4 6-.2 8.2c2.2 2.2 5.9 2.1 8.2-.2c.3-.3 6.7-8.4 14.2-15.9c1.4-1.4 8 3.7 13.6-1.8c3.5-3.6 5.2-9.4 4-14.3M9.4 41.1c-1.4 0-2.5-1.1-2.5-2.5C6.9 37.1 8 36 9.4 36s2.5 1.1 2.5 2.5s-1.1 2.6-2.5 2.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Survey(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#455A64" d="M36 4H26c0 1.1-.9 2-2 2s-2-.9-2-2H12C9.8 4 8 5.8 8 8v32c0 2.2 1.8 4 4 4h24c2.2 0 4-1.8 4-4V8c0-2.2-1.8-4-4-4"/><path fill="#fff" d="M36 41H12c-.6 0-1-.4-1-1V8c0-.6.4-1 1-1h24c.6 0 1 .4 1 1v32c0 .6-.4 1-1 1"/><g fill="#90A4AE"><path d="M26 4c0 1.1-.9 2-2 2s-2-.9-2-2h-7v4c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2V4z"/><path d="M24 0c-2.2 0-4 1.8-4 4s1.8 4 4 4s4-1.8 4-4s-1.8-4-4-4m0 6c-1.1 0-2-.9-2-2s.9-2 2-2s2 .9 2 2s-.9 2-2 2"/></g><path fill="#CFD8DC" d="M21 20h12v2H21zm-6-1h4v4h-4z"/><path fill="#03A9F4" d="M21 29h12v2H21zm-6-1h4v4h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwitchCamera(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#5E35B1"><path d="M33.9 12.1H14.2L17.6 7c.4-.6 1-.9 1.7-.9h9.6c.7 0 1.3.3 1.7.9zM14 11H8V9.2C8 8.5 8.5 8 9.2 8h3.6c.7 0 1.2.5 1.2 1.2z"/><path d="M40 42H8c-2.2 0-4-1.8-4-4V14c0-2.2 1.8-4 4-4h32c2.2 0 4 1.8 4 4v24c0 2.2-1.8 4-4 4"/></g><path fill="#E8EAF6" d="M34 25c0-5.5-4.5-10-10-10c-2.4 0-4.6.8-6.3 2.2l1.2 1.6C20.3 17.7 22 17 24 17c4.4 0 8 3.6 8 8h-3.5l4.5 5.6l4.5-5.6zm-4.9 6.2C27.7 32.3 25.9 33 24 33c-4.4 0-8-3.6-8-8h3.5L15 19.4L10.5 25H14c0 5.5 4.5 10 10 10c2.4 0 4.6-.8 6.3-2.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Synchronize(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#FF6F00" d="m38.7 11.9l-3.1 2.5c2.2 2.7 3.4 6.1 3.4 9.5c0 8.3-6.7 15-15 15c-.9 0-1.9-.1-2.8-.3l-.7 3.9c1.2.2 2.4.3 3.5.3c10.5 0 19-8.5 19-19c0-4.2-1.5-8.5-4.3-11.9"/><path fill="#FF6F02" d="m31 8l11.9 1.6l-9.8 9.8z"/><path fill="#FF6F00" d="M24 5C13.5 5 5 13.5 5 24c0 4.6 1.6 9 4.6 12.4l3-2.6C10.3 31.1 9 27.6 9 24c0-8.3 6.7-15 15-15c.9 0 1.9.1 2.8.3l.7-3.9C26.4 5.1 25.2 5 24 5"/><path fill="#FF6F02" d="M17 40L5.1 38.4l9.8-9.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TabletAndroid(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#37474F" d="M8 41V7c0-2.2 1.8-4 4-4h24c2.2 0 4 1.8 4 4v34c0 2.2-1.8 4-4 4H12c-2.2 0-4-1.8-4-4"/><path fill="#BBDEFB" d="M36 6H12c-.6 0-1 .4-1 1v31c0 .6.4 1 1 1h24c.6 0 1-.4 1-1V7c0-.6-.4-1-1-1"/><path fill="#78909C" d="M21 41h6v2h-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Template(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#BBDEFB" d="M4 7h40v34H4z"/><path fill="#3F51B5" d="M9 12h30v5H9z"/><path fill="#2196F3" d="M9 21h13v16H9zm17 0h13v16H26z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Timeline(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#3F51B5" d="M42 29H20.8c-.5 0-1-.2-1.4-.6l-3.7-3.7c-.4-.4-.4-1 0-1.4l3.7-3.7c.4-.4.9-.6 1.4-.6H42c.6 0 1 .4 1 1v8c0 .6-.4 1-1 1"/><path fill="#CFD8DC" d="M9 6h2v36H9z"/><g fill="#90A4AE"><circle cx="10" cy="10" r="3"/><circle cx="10" cy="24" r="3"/><circle cx="10" cy="38" r="3"/></g><path fill="#448AFF" d="M34 43H20.8c-.5 0-1-.2-1.4-.6l-3.7-3.7c-.4-.4-.4-1 0-1.4l3.7-3.7c.4-.4.9-.6 1.4-.6H34c.6 0 1 .4 1 1v8c0 .6-.4 1-1 1"/><path fill="#00BCD4" d="M35 15H20.8c-.5 0-1-.2-1.4-.6l-3.7-3.7c-.4-.4-.4-1 0-1.4l3.7-3.7c.4-.4.9-.6 1.4-.6H35c.6 0 1 .4 1 1v8c0 .6-.4 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TodoList(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#3F51B5" d="m17.8 18.1l-7.4 7.3l-4.2-4.1L4 23.5l6.4 6.4l9.6-9.6zm0-13l-7.4 7.3l-4.2-4.1L4 10.5l6.4 6.4L20 7.3zm0 26l-7.4 7.3l-4.2-4.1L4 36.5l6.4 6.4l9.6-9.6z"/><path fill="#90CAF9" d="M24 22h20v4H24zm0-13h20v4H24zm0 26h20v4H24z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TouchscreenSmartphone(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#E38939" d="M12 40V8c0-2.2 1.8-4 4-4h16c2.2 0 4 1.8 4 4v32c0 2.2-1.8 4-4 4H16c-2.2 0-4-1.8-4-4"/><path fill="#FFF3E0" d="M32 7H16c-.6 0-1 .4-1 1v29c0 .6.4 1 1 1h16c.6 0 1-.4 1-1V8c0-.6-.4-1-1-1"/><circle cx="24" cy="41" r="1.5" fill="#A6642A"/><circle cx="24" cy="23" r="2" fill="#E91E63"/><circle cx="24" cy="23" r="4" fill="none" stroke="#F48FB1" stroke-miterlimit="10" stroke-width="2"/><circle cx="24" cy="23" r="6.5" fill="none" stroke="#F8BBD0" stroke-miterlimit="10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trademark(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="24" cy="24" r="21" fill="#9C27B0"/><path fill="#E1BEE7" d="M20.6 18.5h-4.2v14.2h-3.5V18.5H8.7v-2.9h11.9zm6.5-2.9L30.3 28l3.2-12.4H38v17.1h-3.5v-4.6l.3-7.1l-3.4 11.8H29L25.7 21l.3 7.1v4.6h-3.5V15.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TreeStructure(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#90CAF9" d="m36.9 13.8l-1.8-3.6L7.5 24l27.6 13.8l1.8-3.6L16.5 24z"/><path fill="#D81B60" d="M6 18h12v12H6z"/><path fill="#2196F3" d="M30 6h12v12H30zm0 24h12v12H30z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwoSmartphones(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#37474F" d="M6 36V8c0-2.2 1.8-4 4-4h14c2.2 0 4 1.8 4 4v28c0 2.2-1.8 4-4 4H10c-2.2 0-4-1.8-4-4"/><path fill="#BBDEFB" d="M24 7H10c-.6 0-1 .4-1 1v25c0 .6.4 1 1 1h14c.6 0 1-.4 1-1V8c0-.6-.4-1-1-1"/><path fill="#78909C" d="M14 36h6v2h-6z"/><path fill="#E38939" d="M20 40V12c0-2.2 1.8-4 4-4h14c2.2 0 4 1.8 4 4v28c0 2.2-1.8 4-4 4H24c-2.2 0-4-1.8-4-4"/><path fill="#FFF3E0" d="M38 11H24c-.6 0-1 .4-1 1v25c0 .6.4 1 1 1h14c.6 0 1-.4 1-1V12c0-.6-.4-1-1-1"/><circle cx="31" cy="41" r="1.5" fill="#A6642A"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Undo(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#00BCD4"><path d="M5 18L19 6.3v23.4z"/><path d="M28 14H16v8h12c2.8 0 5 2.2 5 5s-2.2 5-5 5h-3v8h3c7.2 0 13-5.8 13-13s-5.8-13-13-13"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unlock(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#424242" d="M24 4c-5.5 0-10 4.5-10 10v4h4v-4c0-3.3 2.7-6 6-6s6 2.7 6 6h4c0-5.5-4.5-10-10-10"/><path fill="#FB8C00" d="M36 44H12c-2.2 0-4-1.8-4-4V22c0-2.2 1.8-4 4-4h24c2.2 0 4 1.8 4 4v18c0 2.2-1.8 4-4 4"/><circle cx="24" cy="31" r="3" fill="#C76E00"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Up(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#3F51B5"><path d="m24 4l11.7 14H12.3z"/><path d="M20 15h8v27h-8z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpLeft(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#3F51B5" d="m4 19l14 11.7V7.3z"/><path fill="#3F51B5" d="M42 27v13h-8V27c0-2.2-1.8-4-4-4H13v-8h17c6.6 0 12 5.4 12 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpRight(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<defs><path id="flatColorIconsUpRight0" fill="#3F51B5" d="M44 19L30 30.7V7.3z"/><path id="flatColorIconsUpRight1" fill="#3F51B5" d="M6 27v13h8V27c0-2.2 1.8-4 4-4h17v-8H18c-6.6 0-12 5.4-12 12"/></defs><use href="#flatColorIconsUpRight0"/><use href="#flatColorIconsUpRight1"/><use href="#flatColorIconsUpRight0"/><use href="#flatColorIconsUpRight1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Upload(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#009688"><path d="M24 10.9L35 24H13zM20 40h8v4h-8zm0-6h8v4h-8z"/><path d="M20 21h8v11h-8zM6 4h36v4H6z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Usb(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#1565C0" d="M38.701 24.355h-2.189l-.467 2.265h2.51c.233 0 1.545-.167 1.545-1.267c0-1.087-1.399-.998-1.399-.998m.82-4.016h-2.15l-.374 1.796h2.337c.188 0 1.113-.146 1.113-1.006c0-.858-.926-.79-.926-.79m4.543 2.77s1.436-.743 1.436-3.093c0-3.715-4.377-3.516-4.377-3.516h-2.865l.674-3H17.961c-9.344 0-12.158 6.774-12.158 6.774l-.067.226H2.547l-1.047 6h3.37l.001.143S4.586 33.5 15.334 33.5h19.042l.679-3h4.389c4.729 0 5.591-3.354 5.591-4.9c-.003-1.762-.971-2.491-.971-2.491"/><path fill="#FFF" d="M38.701 24.355h-2.189l-.467 2.265h2.51c.233 0 1.545-.167 1.545-1.267c0-1.087-1.399-.998-1.399-.998m.82-4.016h-2.15l-.374 1.796h2.337c.188 0 1.113-.146 1.113-1.006c0-.858-.926-.79-.926-.79M14.022 29.5c-5.306 0-5.306-3.624-5.238-3.986a1911.4 1911.4 0 0 1 1.789-8.014h3.84l-1.358 6.354s-.971 2.728 1.251 2.728c2.081 0 2.336-2.535 2.336-2.535l1.465-6.543h3.839l-1.582 6.979c.001-.003-.106 5.017-6.342 5.017m12.076.021c-2.674 0-4.958-1.262-4.856-4.14h3.438c0 .576.086 1.627 1.633 1.627c.627 0 1.688-.266 1.688-1.133c0-1.631-5.597-.785-5.597-4.57c0-2.063 1.899-3.785 4.989-3.785c4.976 0 4.613 3.749 4.613 3.749h-3.369c0-1.044-.664-1.204-1.463-1.204c-.8 0-1.372.343-1.372.944c0 1.471 5.634.456 5.634 4.531c0 1.765-1.424 3.981-5.338 3.981m13.268-.021h-7.515l2.601-12h6.556c1.113 0 3.43.234 3.43 2.542c0 2.602-2.227 3.013-2.227 3.013s1.764.407 1.764 2.473c0 3.929-4.2 3.972-4.609 3.972"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoCall(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#4CAF50" d="M8 12h22c2.2 0 4 1.8 4 4v16c0 2.2-1.8 4-4 4H8c-2.2 0-4-1.8-4-4V16c0-2.2 1.8-4 4-4"/><path fill="#388E3C" d="m44 35l-10-6V19l10-6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoFile(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#90CAF9" d="M40 45H8V3h22l10 10z"/><path fill="#E1F5FE" d="M38.5 14H29V4.5z"/><path fill="#1976D2" d="m30 28l-10-6v12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoProjector(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#546E7A" d="M5 34h6v3H5zm32 0h6v3h-6z"/><path fill="#78909C" d="M44 35H4c-2.2 0-4-1.8-4-4V17c0-2.2 1.8-4 4-4h40c2.2 0 4 1.8 4 4v14c0 2.2-1.8 4-4 4"/><path fill="#37474F" d="M5 19h2v2H5zm0 4h2v2H5zm0 4h2v2H5zm4-8h2v2H9zm0 4h2v2H9zm0 4h2v2H9zm4-8h2v2h-2zm0 4h2v2h-2zm0 4h2v2h-2zm4-8h2v2h-2zm0 4h2v2h-2zm0 4h2v2h-2zm4-8h2v2h-2zm0 4h2v2h-2zm0 4h2v2h-2z"/><circle cx="37" cy="24" r="8" fill="#37474F"/><circle cx="37" cy="24" r="6" fill="#a0f"/><path fill="#EA80FC" d="M40.7 21.7c-1-1.1-2.3-1.7-3.7-1.7s-2.8.6-3.7 1.7c-.4.4-.3 1 .1 1.4c.4.4 1 .3 1.4-.1c1.2-1.3 3.3-1.3 4.5 0c.2.2.5.3.7.3c.2 0 .5-.1.7-.3c.4-.3.4-.9 0-1.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewDetails(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#BBDEFB" d="M7 4h34v40H7z"/><path fill="#2196F3" d="M13 26h4v4h-4zm0-8h4v4h-4zm0 16h4v4h-4zm0-24h4v4h-4zm8 16h14v4H21zm0-8h14v4H21zm0 16h14v4H21zm0-24h14v4H21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vip(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#880E4F" d="M38 43H10c-2.2 0-4-1.8-4-4V11c0-2.2 1.8-4 4-4h28c2.2 0 4 1.8 4 4v28c0 2.2-1.8 4-4 4"/><path fill="#FFD54F" d="m15.9 28l2.1-9.1h2.8l-3.6 12.6h-2.6L11 18.9h2.8zm9.7 3.5h-2.5V18.9h2.5zm5.6-4.4v4.4h-2.5V18.9H33c3.7 0 4.1 3.4 4.1 4.2c0 1.2-.5 4-4.1 4zm0-2.2h1.7c1.3 0 1.5-1.1 1.5-1.9c0-1.6-.9-2.1-1.5-2.1h-1.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vlc(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#F57C00" d="M36.258 28.837S36.148 28 35.001 28h-3.719c.798 2.671 1.497 5.135 1.497 5.279c0 2.387-3.401 3.393-8.917 3.393c-5.515 0-8.651-.94-8.651-3.326c0-.167.998-2.692 1.791-5.346h-4.063c-.806 0-.937.749-.937.749L8.159 40.986L8.815 42h30.652l.376-1.014z"/><path fill="#E0E0E0" d="M24.001 6c-1.029 0-1.864.179-1.864.398c-.492 1.483-8.122 26.143-8.122 26.774c0 2.388 4.471 3.827 9.985 3.827s9.986-1.439 9.986-3.827c0-.549-7.614-25.268-8.122-26.774c.001-.219-.833-.398-1.863-.398"/><path fill="#FF9800" d="M33.196 30.447C32.032 32.232 28.341 34 24.046 34c-4.34 0-8.156-1.696-9.281-3.51c-.499 1.483-.892 2.647-.892 3.28c0 2.386 4.533 4.229 10.128 4.229s10.131-1.844 10.131-4.229c0-.548-.419-1.815-.936-3.323m-1.809-6.133l-2.074-6.794s-1.857 1.479-5.311 1.479c-3.453 0-5.316-1.479-5.316-1.479l-2.081 6.806S18.673 27 24.002 27c5.373 0 7.385-2.686 7.385-2.686m-4.146-13.505l-1.376-4.41s-.083-.398-1.864-.398c-1.844 0-1.864.398-1.864.398l-1.376 4.407S21.646 12 24 12c2.355 0 3.241-1.191 3.241-1.191"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VoicePresentation(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#2196F3" d="M40 22h-8l-4 4V12c0-2.2 1.8-4 4-4h8c2.2 0 4 1.8 4 4v6c0 2.2-1.8 4-4 4"/><circle cx="17" cy="19" r="8" fill="#FFA726"/><path fill="#607D8B" d="M30 36.7S26.4 30 17 30S4 36.7 4 36.7V40h26z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Voicemail(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#4CAF50" d="M48 24c0-6.1-4.9-11-11-11s-11 4.9-11 11c0 2.7.9 5.1 2.5 7h-9c1.6-1.9 2.5-4.3 2.5-7c0-6.1-4.9-11-11-11S0 17.9 0 24s4.9 11 11 11h27v-.1c5.6-.5 10-5.2 10-10.9M4 24c0-3.9 3.1-7 7-7s7 3.1 7 7s-3.1 7-7 7s-7-3.1-7-7m33 7c-3.9 0-7-3.1-7-7s3.1-7 7-7s7 3.1 7 7s-3.1 7-7 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Webcam(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#455A64" d="M36.5 44h-25c-1.1 0-1.8-1.2-1.3-2.2L13 37h22l2.7 4.8c.6 1-.1 2.2-1.2 2.2"/><circle cx="24" cy="23" r="18" fill="#78909C"/><path fill="#455A64" d="M24 35c-6.6 0-12-5.4-12-12s5.4-12 12-12s12 5.4 12 12s-5.4 12-12 12"/><circle cx="24" cy="23" r="9" fill="#42A5F5"/><path fill="#90CAF9" d="M28.8 20c-1.2-1.4-3-2.2-4.8-2.2s-3.6.8-4.8 2.2c-.5.5-.4 1.3.1 1.8s1.3.4 1.8-.1c1.5-1.7 4.3-1.7 5.8 0c.3.3.6.4 1 .4c.3 0 .6-.1.9-.3c.4-.4.5-1.3 0-1.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WiFiLogo(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="#3F51B5"><path d="M46 26.48c0 4.527-3.268 7.52-7.3 7.52H9.299C5.269 34 2 30.634 2 26.48v-4.96C2 17.366 5.269 14 9.299 14H38.7c4.032 0 7.3 3.366 7.3 7.52z"/><ellipse cx="24" cy="24" rx="14.902" ry="15"/></g><path fill="#FFF" d="M17 19h-2.736l-.837 5.859l-1.039-5.831H9.93l-1.066 5.831l-.81-5.831H5.266L7.597 29h2.459l1.064-6.146L12.209 29h2.484zm2 3h2.508v7H19zm2.5-2.253a1.251 1.251 0 1 1-2.5 0c0-.696.56-1.258 1.25-1.258s1.25.562 1.25 1.258M38.561 16h-7.979S25 16.193 25 21.914v4.336s.101 2.941-3 5.75h16.785S44 32 44 26.447v-4.879S43.772 16 38.561 16m-1.222 5.369h-5.651v2.236h5.094v2.344h-5.094V29H29V19h8.339zm2.911-2.88c.689 0 1.25.562 1.25 1.258c0 .693-.561 1.253-1.25 1.253S39 20.44 39 19.747c0-.696.561-1.258 1.25-1.258M41.508 29H39v-7h2.508z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wikipedia(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#CFD8DC" d="M6 10a4 4 0 0 1 4-4h28a4 4 0 0 1 4 4v28a4 4 0 0 1-4 4H10a4 4 0 0 1-4-4z"/><path fill="#37474F" d="M39 17.271a.342.342 0 0 1-.334.349h-1.799l-8.164 18.179c-.052.12-.17.2-.297.202h-.004a.33.33 0 0 1-.298-.193l-3.874-8.039l-4.18 8.049a.333.333 0 0 1-.303.184a.336.336 0 0 1-.292-.199l-8.252-18.182h-1.87a.345.345 0 0 1-.333-.35v-.921a.34.34 0 0 1 .333-.35h6.657a.34.34 0 0 1 .333.35v.921a.342.342 0 0 1-.333.349h-1.433l5.696 13.748l2.964-5.793l-3.757-7.953h-.904a.342.342 0 0 1-.333-.35v-.922c0-.191.149-.348.333-.348h4.924a.34.34 0 0 1 .333.348v.922c0 .192-.149.35-.333.35h-.867l2.162 4.948l2.572-4.948H25.77a.34.34 0 0 1-.334-.35v-.922a.34.34 0 0 1 .334-.348h4.784c.187 0 .333.156.333.348v.922a.34.34 0 0 1-.333.35h-1.05l-3.757 7.141l3.063 6.584l5.905-13.725h-1.872a.343.343 0 0 1-.334-.35v-.922c0-.191.15-.348.334-.348h5.822a.34.34 0 0 1 .334.348z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Workflow(children ...ElementRenderer) *FlatColorIconsIcon {
	return &FlatColorIconsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="#00BCD4" d="M7 31h10v10H7zm28.3-11.7l-5.6-5.6c-.4-.4-.4-1 0-1.4l5.6-5.6c.4-.4 1-.4 1.4 0l5.6 5.6c.4.4.4 1 0 1.4l-5.6 5.6c-.4.4-1 .4-1.4 0"/><circle cx="12" cy="13" r="6" fill="#3F51B5"/><circle cx="36" cy="36" r="6" fill="#448AFF"/><g fill="#90A4AE"><path d="M11 24h2v5h-2z"/><path d="m12 21l-3 4h6z"/></g><g fill="#90A4AE"><path d="M20 12h5v2h-5z"/><path d="m28 13l-4-3v6z"/></g><g fill="#90A4AE"><path d="M35 21h2v5h-2z"/><path d="m36 29l3-4h-6z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
