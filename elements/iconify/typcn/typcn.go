package typcn

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = "2.1.2"
	hAttr          = "1em"
	viewbox        = "0 0 24 24"
	fill           = "currentColor"
)

type TypcnIcon struct {
	*SVGSVGElement
}

type TypcnIconFn func(children ...ElementRenderer) *TypcnIcon

var IconLookup = map[string]TypcnIconFn{
	"adjustBrightness":          AdjustBrightness,
	"adjustContrast":            AdjustContrast,
	"anchor":                    Anchor,
	"anchorOutline":             AnchorOutline,
	"archive":                   Archive,
	"arrowBack":                 ArrowBack,
	"arrowBackOutline":          ArrowBackOutline,
	"arrowDown":                 ArrowDown,
	"arrowDownOutline":          ArrowDownOutline,
	"arrowDownThick":            ArrowDownThick,
	"arrowForward":              ArrowForward,
	"arrowForwardOutline":       ArrowForwardOutline,
	"arrowLeft":                 ArrowLeft,
	"arrowLeftOutline":          ArrowLeftOutline,
	"arrowLeftThick":            ArrowLeftThick,
	"arrowLoop":                 ArrowLoop,
	"arrowLoopOutline":          ArrowLoopOutline,
	"arrowMaximise":             ArrowMaximise,
	"arrowMaximiseOutline":      ArrowMaximiseOutline,
	"arrowMinimise":             ArrowMinimise,
	"arrowMinimiseOutline":      ArrowMinimiseOutline,
	"arrowMove":                 ArrowMove,
	"arrowMoveOutline":          ArrowMoveOutline,
	"arrowRepeat":               ArrowRepeat,
	"arrowRepeatOutline":        ArrowRepeatOutline,
	"arrowRight":                ArrowRight,
	"arrowRightOutline":         ArrowRightOutline,
	"arrowRightThick":           ArrowRightThick,
	"arrowShuffle":              ArrowShuffle,
	"arrowSortedDown":           ArrowSortedDown,
	"arrowSortedUp":             ArrowSortedUp,
	"arrowSync":                 ArrowSync,
	"arrowSyncOutline":          ArrowSyncOutline,
	"arrowUnsorted":             ArrowUnsorted,
	"arrowUp":                   ArrowUp,
	"arrowUpOutline":            ArrowUpOutline,
	"arrowUpThick":              ArrowUpThick,
	"at":                        At,
	"attachment":                Attachment,
	"attachmentOutline":         AttachmentOutline,
	"backspace":                 Backspace,
	"backspaceOutline":          BackspaceOutline,
	"batteryCharge":             BatteryCharge,
	"batteryFull":               BatteryFull,
	"batteryHigh":               BatteryHigh,
	"batteryLow":                BatteryLow,
	"batteryMid":                BatteryMid,
	"beaker":                    Beaker,
	"beer":                      Beer,
	"bell":                      Bell,
	"book":                      Book,
	"bookmark":                  Bookmark,
	"briefcase":                 Briefcase,
	"brush":                     Brush,
	"businessCard":              BusinessCard,
	"calculator":                Calculator,
	"calendar":                  Calendar,
	"calendarOutline":           CalendarOutline,
	"camera":                    Camera,
	"cameraOutline":             CameraOutline,
	"cancel":                    Cancel,
	"cancelOutline":             CancelOutline,
	"chartArea":                 ChartArea,
	"chartAreaOutline":          ChartAreaOutline,
	"chartBar":                  ChartBar,
	"chartBarOutline":           ChartBarOutline,
	"chartLine":                 ChartLine,
	"chartLineOutline":          ChartLineOutline,
	"chartPie":                  ChartPie,
	"chartPieOutline":           ChartPieOutline,
	"chevronLeft":               ChevronLeft,
	"chevronLeftOutline":        ChevronLeftOutline,
	"chevronRight":              ChevronRight,
	"chevronRightOutline":       ChevronRightOutline,
	"clipboard":                 Clipboard,
	"cloudStorage":              CloudStorage,
	"cloudStorageOutline":       CloudStorageOutline,
	"code":                      Code,
	"codeOutline":               CodeOutline,
	"coffee":                    Coffee,
	"cog":                       Cog,
	"cogOutline":                CogOutline,
	"compass":                   Compass,
	"contacts":                  Contacts,
	"creditCard":                CreditCard,
	"cssThree":                  CssThree,
	"database":                  Database,
	"delete":                    Delete,
	"deleteOutline":             DeleteOutline,
	"deviceDesktop":             DeviceDesktop,
	"deviceLaptop":              DeviceLaptop,
	"devicePhone":               DevicePhone,
	"deviceTablet":              DeviceTablet,
	"directions":                Directions,
	"divide":                    Divide,
	"divideOutline":             DivideOutline,
	"document":                  Document,
	"documentAdd":               DocumentAdd,
	"documentDelete":            DocumentDelete,
	"documentText":              DocumentText,
	"download":                  Download,
	"downloadOutline":           DownloadOutline,
	"dropbox":                   Dropbox,
	"edit":                      Edit,
	"eject":                     Eject,
	"ejectOutline":              EjectOutline,
	"equals":                    Equals,
	"equalsOutline":             EqualsOutline,
	"export":                    Export,
	"exportOutline":             ExportOutline,
	"eye":                       Eye,
	"eyeOutline":                EyeOutline,
	"feather":                   Feather,
	"film":                      Film,
	"filter":                    Filter,
	"flag":                      Flag,
	"flagOutline":               FlagOutline,
	"flash":                     Flash,
	"flashOutline":              FlashOutline,
	"flowChildren":              FlowChildren,
	"flowMerge":                 FlowMerge,
	"flowParallel":              FlowParallel,
	"flowSwitch":                FlowSwitch,
	"folder":                    Folder,
	"folderAdd":                 FolderAdd,
	"folderDelete":              FolderDelete,
	"folderOpen":                FolderOpen,
	"gift":                      Gift,
	"globe":                     Globe,
	"globeOutline":              GlobeOutline,
	"groupIcon":                 GroupIcon,
	"groupOutline":              GroupOutline,
	"headphones":                Headphones,
	"heart":                     Heart,
	"heartFullOutline":          HeartFullOutline,
	"heartHalfOutline":          HeartHalfOutline,
	"heartOutline":              HeartOutline,
	"home":                      Home,
	"homeOutline":               HomeOutline,
	"htmlFive":                  HtmlFive,
	"image":                     Image,
	"imageOutline":              ImageOutline,
	"infinity":                  Infinity,
	"infinityOutline":           InfinityOutline,
	"info":                      Info,
	"infoLarge":                 InfoLarge,
	"infoLargeOutline":          InfoLargeOutline,
	"infoOutline":               InfoOutline,
	"inputChecked":              InputChecked,
	"inputCheckedOutline":       InputCheckedOutline,
	"key":                       Key,
	"keyOutline":                KeyOutline,
	"keyboard":                  Keyboard,
	"leaf":                      Leaf,
	"lightbulb":                 Lightbulb,
	"link":                      Link,
	"linkOutline":               LinkOutline,
	"location":                  Location,
	"locationArrow":             LocationArrow,
	"locationArrowOutline":      LocationArrowOutline,
	"locationOutline":           LocationOutline,
	"lockClosed":                LockClosed,
	"lockClosedOutline":         LockClosedOutline,
	"lockOpen":                  LockOpen,
	"lockOpenOutline":           LockOpenOutline,
	"mail":                      Mail,
	"map":                       Map,
	"mediaEject":                MediaEject,
	"mediaEjectOutline":         MediaEjectOutline,
	"mediaFastForward":          MediaFastForward,
	"mediaFastForwardOutline":   MediaFastForwardOutline,
	"mediaPause":                MediaPause,
	"mediaPauseOutline":         MediaPauseOutline,
	"mediaPlay":                 MediaPlay,
	"mediaPlayOutline":          MediaPlayOutline,
	"mediaPlayReverse":          MediaPlayReverse,
	"mediaPlayReverseOutline":   MediaPlayReverseOutline,
	"mediaRecord":               MediaRecord,
	"mediaRecordOutline":        MediaRecordOutline,
	"mediaRewind":               MediaRewind,
	"mediaRewindOutline":        MediaRewindOutline,
	"mediaStop":                 MediaStop,
	"mediaStopOutline":          MediaStopOutline,
	"message":                   Message,
	"messageTyping":             MessageTyping,
	"messages":                  Messages,
	"microphone":                Microphone,
	"microphoneOutline":         MicrophoneOutline,
	"minus":                     Minus,
	"minusOutline":              MinusOutline,
	"mortarBoard":               MortarBoard,
	"news":                      News,
	"notes":                     Notes,
	"notesOutline":              NotesOutline,
	"pen":                       Pen,
	"pencil":                    Pencil,
	"phone":                     Phone,
	"phoneOutline":              PhoneOutline,
	"pi":                        Pi,
	"piOutline":                 PiOutline,
	"pin":                       Pin,
	"pinOutline":                PinOutline,
	"pipette":                   Pipette,
	"plane":                     Plane,
	"planeOutline":              PlaneOutline,
	"plug":                      Plug,
	"plus":                      Plus,
	"plusOutline":               PlusOutline,
	"pointOfInterest":           PointOfInterest,
	"pointOfInterestOutline":    PointOfInterestOutline,
	"power":                     Power,
	"powerOutline":              PowerOutline,
	"printer":                   Printer,
	"puzzle":                    Puzzle,
	"puzzleOutline":             PuzzleOutline,
	"radar":                     Radar,
	"radarOutline":              RadarOutline,
	"refresh":                   Refresh,
	"refreshOutline":            RefreshOutline,
	"rss":                       Rss,
	"rssOutline":                RssOutline,
	"scissors":                  Scissors,
	"scissorsOutline":           ScissorsOutline,
	"shoppingBag":               ShoppingBag,
	"shoppingCart":              ShoppingCart,
	"socialAtCircular":          SocialAtCircular,
	"socialDribbble":            SocialDribbble,
	"socialDribbbleCircular":    SocialDribbbleCircular,
	"socialFacebook":            SocialFacebook,
	"socialFacebookCircular":    SocialFacebookCircular,
	"socialFlickr":              SocialFlickr,
	"socialFlickrCircular":      SocialFlickrCircular,
	"socialGithub":              SocialGithub,
	"socialGithubCircular":      SocialGithubCircular,
	"socialGooglePlus":          SocialGooglePlus,
	"socialGooglePlusCircular":  SocialGooglePlusCircular,
	"socialInstagram":           SocialInstagram,
	"socialInstagramCircular":   SocialInstagramCircular,
	"socialLastFm":              SocialLastFm,
	"socialLastFmCircular":      SocialLastFmCircular,
	"socialLinkedin":            SocialLinkedin,
	"socialLinkedinCircular":    SocialLinkedinCircular,
	"socialPinterest":           SocialPinterest,
	"socialPinterestCircular":   SocialPinterestCircular,
	"socialSkype":               SocialSkype,
	"socialSkypeOutline":        SocialSkypeOutline,
	"socialTumbler":             SocialTumbler,
	"socialTumblerCircular":     SocialTumblerCircular,
	"socialTwitter":             SocialTwitter,
	"socialTwitterCircular":     SocialTwitterCircular,
	"socialVimeo":               SocialVimeo,
	"socialVimeoCircular":       SocialVimeoCircular,
	"socialYoutube":             SocialYoutube,
	"socialYoutubeCircular":     SocialYoutubeCircular,
	"sortAlphabetically":        SortAlphabetically,
	"sortAlphabeticallyOutline": SortAlphabeticallyOutline,
	"sortNumerically":           SortNumerically,
	"sortNumericallyOutline":    SortNumericallyOutline,
	"spanner":                   Spanner,
	"spannerOutline":            SpannerOutline,
	"spiral":                    Spiral,
	"star":                      Star,
	"starFullOutline":           StarFullOutline,
	"starHalf":                  StarHalf,
	"starHalfOutline":           StarHalfOutline,
	"starOutline":               StarOutline,
	"starburst":                 Starburst,
	"starburstOutline":          StarburstOutline,
	"stopwatch":                 Stopwatch,
	"support":                   Support,
	"tabsOutline":               TabsOutline,
	"tag":                       Tag,
	"tags":                      Tags,
	"thLarge":                   ThLarge,
	"thLargeOutline":            ThLargeOutline,
	"thList":                    ThList,
	"thListOutline":             ThListOutline,
	"thMenu":                    ThMenu,
	"thMenuOutline":             ThMenuOutline,
	"thSmall":                   ThSmall,
	"thSmallOutline":            ThSmallOutline,
	"thermometer":               Thermometer,
	"thumbsDown":                ThumbsDown,
	"thumbsOk":                  ThumbsOk,
	"thumbsUp":                  ThumbsUp,
	"tick":                      Tick,
	"tickOutline":               TickOutline,
	"ticket":                    Ticket,
	"time":                      Time,
	"times":                     Times,
	"timesOutline":              TimesOutline,
	"trash":                     Trash,
	"tree":                      Tree,
	"upload":                    Upload,
	"uploadOutline":             UploadOutline,
	"user":                      User,
	"userAdd":                   UserAdd,
	"userAddOutline":            UserAddOutline,
	"userDelete":                UserDelete,
	"userDeleteOutline":         UserDeleteOutline,
	"userOutline":               UserOutline,
	"vendorAndroid":             VendorAndroid,
	"vendorApple":               VendorApple,
	"vendorMicrosoft":           VendorMicrosoft,
	"video":                     Video,
	"videoOutline":              VideoOutline,
	"volume":                    Volume,
	"volumeDown":                VolumeDown,
	"volumeMute":                VolumeMute,
	"volumeUp":                  VolumeUp,
	"warning":                   Warning,
	"warningOutline":            WarningOutline,
	"watch":                     Watch,
	"waves":                     Waves,
	"wavesOutline":              WavesOutline,
	"weatherCloudy":             WeatherCloudy,
	"weatherDownpour":           WeatherDownpour,
	"weatherNight":              WeatherNight,
	"weatherPartlySunny":        WeatherPartlySunny,
	"weatherShower":             WeatherShower,
	"weatherSnow":               WeatherSnow,
	"weatherStormy":             WeatherStormy,
	"weatherSunny":              WeatherSunny,
	"weatherWindy":              WeatherWindy,
	"weatherWindyCloudy":        WeatherWindyCloudy,
	"wiFi":                      WiFi,
	"wiFiOutline":               WiFiOutline,
	"wine":                      Wine,
	"world":                     World,
	"worldOutline":              WorldOutline,
	"zoom":                      Zoom,
	"zoomIn":                    ZoomIn,
	"zoomInOutline":             ZoomInOutline,
	"zoomOut":                   ZoomOut,
	"zoomOutOutline":            ZoomOutOutline,
	"zoomOutline":               ZoomOutline,
}

func AdjustBrightness(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 6.934L13 4a1.055 1.055 0 1 0-2 0zM4 11a1.055 1.055 0 1 0 0 2l2.934-1zm8 6.066L11 20a1.055 1.055 0 1 0 2 0zm9.341-5.409A1.055 1.055 0 0 0 20 10.998l-2.934 1l2.934 1a1.057 1.057 0 0 0 1.341-1.341M5.636 7.05l2.781 1.367L7.05 5.636A1.057 1.057 0 1 0 5.636 7.05m-.483 10.382a1.057 1.057 0 0 0 1.896.932l1.367-2.781l-2.781 1.367a1.056 1.056 0 0 0-.482.482m13.21-.483l-2.781-1.367l1.367 2.781a1.056 1.056 0 1 0 1.414-1.414m.481-10.383a1.057 1.057 0 0 0-1.895-.933L15.58 8.416l2.782-1.368c.202-.1.375-.264.482-.482M12 7.5c-2.481 0-4.5 2.019-4.5 4.5s2.019 4.5 4.5 4.5s4.5-2.019 4.5-4.5s-2.019-4.5-4.5-4.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdjustContrast(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 4a8 8 0 1 0 0 16a8 8 0 0 0 0-16m0 14a6 6 0 1 1 0-12a6 6 0 0 1 0 12m0-11v10c2.757 0 5-2.243 5-5s-2.243-5-5-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Anchor(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 13.5a1 1 0 0 0-1 1a5.007 5.007 0 0 1-4 4.898V12h4a1 1 0 1 0 0-2h-4V8.816A2.99 2.99 0 0 0 15 6a3 3 0 1 0-6 0a2.99 2.99 0 0 0 2 2.816V10H7a1 1 0 1 0 0 2h4v7.398A5.008 5.008 0 0 1 7 14.5a1 1 0 1 0-2 0c0 3.859 3.141 7 7 7s7-3.141 7-7a1 1 0 0 0-1-1M12 5c.551 0 1 .449 1 1s-.449 1-1 1s-1-.449-1-1s.449-1 1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AnchorOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="12" cy="6" r="1" fill="currentColor"/><path fill="currentColor" d="M19.793 12.096A3.004 3.004 0 0 0 17 8h-.422A4.949 4.949 0 0 0 17 6c0-2.757-2.243-5-5-5S7 3.243 7 6c0 .703.149 1.381.422 2H7a3.004 3.004 0 0 0-2.793 4.096A3 3 0 0 0 3 14.5c0 4.963 4.037 9 9 9s9-4.037 9-9a3 3 0 0 0-1.207-2.404M12 21.5c-3.859 0-7-3.141-7-7a1 1 0 1 1 2 0a5.007 5.007 0 0 0 4 4.898V12H7a1 1 0 1 1 0-2h4V8.816A2.99 2.99 0 0 1 9 6a3 3 0 1 1 6 0a2.99 2.99 0 0 1-2 2.816V10h4a1 1 0 1 1 0 2h-4v7.398a5.008 5.008 0 0 0 4-4.898a1 1 0 1 1 2 0c0 3.859-3.141 7-7 7M7.321 13H10v4.962A4.015 4.015 0 0 1 8 14.5c0-.597-.263-1.133-.679-1.5m9.358 0A1.996 1.996 0 0 0 16 14.5a4.015 4.015 0 0 1-2 3.462V13z"/><circle cx="12" cy="6" r="1" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Archive(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 12h-3a.5.5 0 0 0 0 1h3a.5.5 0 0 0 0-1m7-7H3a1 1 0 1 0 0 2h17a1 1 0 1 0 0-2m-2 3H5a1 1 0 0 0-1 1v8c0 1.654 1.346 3 3 3h9c1.654 0 3-1.346 3-3V9a1 1 0 0 0-1-1m-2 10H7c-.552 0-1-.449-1-1v-7h11v7c0 .551-.448 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowBack(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 9.059V6.5a1.001 1.001 0 0 0-1.707-.708L4 12l6.293 6.207a.997.997 0 0 0 1.414 0A.999.999 0 0 0 12 17.5v-2.489c2.75.068 5.755.566 8 3.989v-1c0-4.633-3.5-8.443-8-8.941"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowBackOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.164 19.547c-1.641-2.5-3.669-3.285-6.164-3.484V17.5c0 .534-.208 1.036-.586 1.414c-.756.756-2.077.751-2.823.005l-6.293-6.207a1 1 0 0 1 0-1.425l6.288-6.203c.754-.755 2.073-.756 2.829.001c.377.378.585.88.585 1.414v1.704c4.619.933 8 4.997 8 9.796v1a.999.999 0 0 1-1.836.548m-7.141-5.536c2.207.056 4.638.394 6.758 2.121a7.985 7.985 0 0 0-6.893-6.08C11.384 9.996 11 10 11 10V6.503l-5.576 5.496l5.576 5.5V14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDown(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.707 13.293a.999.999 0 0 0-1.414 0L13 15.586V8a1 1 0 1 0-2 0v7.586l-2.293-2.293a.999.999 0 1 0-1.414 1.414L12 19.414l4.707-4.707a.999.999 0 0 0 0-1.414"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12 21.312l-7.121-7.121a3.002 3.002 0 0 1 0-4.242C5.973 8.855 7.857 8.811 9 9.834V5c0-1.654 1.346-3 3-3s3 1.346 3 3v4.834c1.143-1.023 3.027-.979 4.121.115a3.002 3.002 0 0 1 0 4.242zM7 11.07a.999.999 0 0 0-.707 1.707L12 18.484l5.707-5.707a.999.999 0 0 0 0-1.414a1.021 1.021 0 0 0-1.414 0L13 14.656V5a1.001 1.001 0 0 0-2 0v9.656l-3.293-3.293A.991.991 0 0 0 7 11.07"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownThick(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.414 10.656a2 2 0 0 0-2.828 0L14 12.242V5a2 2 0 0 0-4 0v7.242l-1.586-1.586a2 2 0 1 0-2.828 2.828L12 19.898l6.414-6.414a2 2 0 0 0 0-2.828"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowForward(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 5.499a.996.996 0 0 0-1 1v2.559c-4.5.498-8 4.309-8 8.941v1c2.245-3.423 5.25-3.92 8-3.989v2.489a.999.999 0 0 0 1.707.707L20 11.999l-6.293-6.208A.996.996 0 0 0 13 5.499"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowForwardOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 19.999a.999.999 0 0 1-1-1v-1a9.98 9.98 0 0 1 8-9.796V6.499c0-.534.208-1.036.585-1.414c.756-.757 2.075-.756 2.829-.001l6.288 6.203a.996.996 0 0 1 0 1.424l-6.293 6.207c-.746.746-2.067.751-2.823-.005A1.986 1.986 0 0 1 11 17.499v-1.437c-2.495.201-4.523.985-6.164 3.484a1 1 0 0 1-.836.453m8-5.989l1-.01v3.499l5.576-5.5L13 6.503V10s-.384-.004-.891.052a7.982 7.982 0 0 0-6.892 6.08C7.338 14.404 9.768 14.066 12 14.01"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeft(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 11H9.414l2.293-2.293a.999.999 0 1 0-1.414-1.414L5.586 12l4.707 4.707a.997.997 0 0 0 1.414 0a.999.999 0 0 0 0-1.414L9.414 13H17a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.928 21a2.978 2.978 0 0 1-2.121-.879L1.686 13l7.121-7.121c1.133-1.134 3.109-1.134 4.242 0c.566.564.879 1.317.879 2.119c0 .746-.27 1.451-.764 2.002H18c1.654 0 3 1.346 3 3s-1.346 3-3 3h-4.836c.493.549.764 1.252.764 1.998a2.977 2.977 0 0 1-.879 2.124a2.983 2.983 0 0 1-2.121.878m-6.414-8l5.707 5.707a1.023 1.023 0 0 0 1.414 0c.189-.189.293-.441.293-.708s-.104-.517-.291-.705L8.342 14H18a1.001 1.001 0 0 0 0-2H8.342l3.293-3.293a.996.996 0 0 0 .001-1.413a1.023 1.023 0 0 0-1.415-.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftThick(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 11h-7.244l1.586-1.586a2 2 0 1 0-2.828-2.828L3.1 13l6.414 6.414c.39.391.902.586 1.414.586s1.023-.195 1.414-.586a2 2 0 0 0 0-2.828L10.756 15H18a2 2 0 0 0 0-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLoop(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.5 8h-2.086l1.293-1.293a.999.999 0 1 0-1.414-1.414L10.586 9l3.707 3.707a.997.997 0 0 0 1.414 0a.999.999 0 0 0 0-1.414L14.414 10H16.5c1.379 0 2.5 1.346 2.5 3s-1.346 3-3 3H8c-1.654 0-3-1.346-3-3s1.346-3 3-3a1 1 0 1 0 0-2c-2.757 0-5 2.243-5 5s2.243 5 5 5h8c2.757 0 5-2.243 5-5s-2.019-5-4.5-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLoopOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.994 7.187L19 7c0-.801-.312-1.555-.879-2.121C17.555 4.312 16.801 4 16 4s-1.555.312-2.121.879l-2.883 2.883A2.987 2.987 0 0 0 9 7H8c-3.859 0-7 3.14-7 7s3.141 7 7 7h9c3.859 0 7-3.14 7-7c0-3.306-2.14-6.084-5.006-6.813M17 19H8c-2.757 0-5-2.243-5-5s2.243-5 5-5h1a1 1 0 1 1 0 2H8c-1.654 0-3 1.346-3 3s1.346 3 3 3h9c1.654 0 3-1.346 3-3s-1.121-3-2.5-3h-2.086l1.293 1.293a.999.999 0 1 1-1.414 1.414L11.586 10l3.707-3.707a.997.997 0 0 1 1.414 0a.999.999 0 0 1 0 1.414L15.414 9H17.5c2.481 0 4.5 2.243 4.5 5s-2.243 5-5 5m.749-6.971c.7.164 1.251 1 1.251 1.971c0 1.103-.897 2-2 2H8c-1.103 0-2-.897-2-2s.897-2 2-2h1c.856 0 1.588-.541 1.873-1.299l3.713 3.713c.378.378.88.586 1.414.586s1.036-.208 1.414-.586S18 13.534 18 13c0-.345-.087-.677-.251-.971"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowMaximise(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 4a1 1 0 1 0 0 2h1.586l-3.293 3.293a.999.999 0 1 0 1.414 1.414L18 7.414V9a1 1 0 1 0 2 0V4zm-5.707 9.293L6 16.586V15a1 1 0 1 0-2 0v4.999h.996L9 20a1 1 0 0 0 0-2H7.414l3.293-3.292c.391-.391.391-1.023 0-1.414s-1.023-.392-1.414-.001M7 12a1 1 0 0 0 1-1V8h3a1 1 0 1 0 0-2H6.001L6 11a1 1 0 0 0 1 1m10 0a1 1 0 0 0-1 1v3h-3a1 1 0 1 0 0 2h5v-5a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowMaximiseOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 3h-5.243a2.993 2.993 0 0 0-2.815 2H4v7.061l.012.12A2.994 2.994 0 0 0 2 15v7h7c1.311 0 2.593-.826 3-2h7v-7.061l-.012-.12A2.994 2.994 0 0 0 21 10V3zm-2 15h-5a1 1 0 1 1 0-2h3v-3.061a1 1 0 1 1 2 0zM6 7h5.061a1 1 0 1 1 0 2H8v3.061a1 1 0 0 1-2 0zm13 3a1 1 0 1 1-2 0V8.414l-3.293 3.293a.997.997 0 0 1-1.414 0a.999.999 0 0 1 0-1.414L15.586 7H14a1 1 0 1 1 0-2h5zM9 20H4v-5a1 1 0 1 1 2 0v1.586l3.293-3.293a.997.997 0 0 1 1.414 0a.999.999 0 0 1 0 1.414L7.414 18H9a1 1 0 1 1 0 2m2.414-7.414a1.986 1.986 0 0 0-2.437-.297L9 12.061V10h2.061l.229-.023c-.186.307-.29.656-.29 1.023c0 .534.208 1.036.586 1.414a1.986 1.986 0 0 0 2.437.297l-.023.228V15h-1.939c-.122 0-.24.015-.356.036a1.983 1.983 0 0 0-.291-2.45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowMinimise(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.121 13a1 1 0 1 0 0 2h1.465l-3.293 3.293a.999.999 0 1 0 1.414 1.414l3.414-3.414V18c0 .552.447 1 1 1s.879-.448.879-1v-5zM7 11a1 1 0 0 0 1-1V8h2a1 1 0 1 0 0-2H6.001L6 10a1 1 0 0 0 1 1m10 2a1 1 0 0 0-1 1v2h-2a1 1 0 1 0 0 2h4v-4a1 1 0 0 0-1-1m1.293-8.707L15 7.586V6a1 1 0 1 0-2 0v5h5a1 1 0 0 0 0-2h-1.586l3.293-3.292c.391-.391.391-1.023 0-1.414s-1.023-.392-1.414-.001"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowMinimiseOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 6c0-.801-.312-1.555-.879-2.121C20.555 3.312 19.801 3 19 3s-1.555.312-2.121.879l-.883.883A2.987 2.987 0 0 0 14 4a2.98 2.98 0 0 0-2.283 1.077A3.026 3.026 0 0 0 11.061 5H5v6.06c0 .255.042.499.102.736a2.985 2.985 0 0 0-.275 4.135l-.947.948C3.312 17.445 3 18.199 3 19s.312 1.555.879 2.121C4.445 21.688 5.199 22 6 22c.539 0 1.334-.152 2.061-.879l.903-.919A2.99 2.99 0 0 0 11 21c.934 0 1.758-.437 2.309-1.107c.241.063.49.107.752.107H20v-6.061c0-.226-.029-.444-.077-.656A2.98 2.98 0 0 0 21 11c0-.766-.288-1.465-.762-1.996l.883-.883C21.688 7.555 22 6.801 22 6M7 7h4a1 1 0 1 1 0 2H9v2a1 1 0 0 1-2 0zm12.707-.293L16.414 10H18a1 1 0 1 1 0 2h-5V7a1 1 0 1 1 2 0v1.586l3.293-3.293a.997.997 0 0 1 1.414 0a.999.999 0 0 1 0 1.414M12 18a1 1 0 1 1-2 0v-1.707l-3.354 3.414c-.195.195-.39.293-.646.293s-.512-.098-.707-.293a.999.999 0 0 1 0-1.414L8.586 15H7.121a1 1 0 1 1 0-2H12zm0-6H9.722c.173-.295.278-.634.278-1v-1h1.061c.342 0 .658-.094.939-.245zm1 1h2.713c-.433.094-.713.33-.713.939V15h-.939c-.391 0-.752.117-1.061.311zm.061 4a1 1 0 0 1 1-1H16v-2a1 1 0 1 1 2 0v4h-3.939a1 1 0 0 1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowMove(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.707 8.293a.999.999 0 1 0-1.414 1.414L17.586 11H13V6.414l1.293 1.293a.997.997 0 0 0 1.414 0a.999.999 0 0 0 0-1.414L12 2.586L8.293 6.293a.999.999 0 1 0 1.414 1.414L11 6.414V11H6.414l1.293-1.293a.999.999 0 1 0-1.414-1.414L2.586 12l3.707 3.707a.997.997 0 0 0 1.414 0a.999.999 0 0 0 0-1.414L6.414 13H11v4.586l-1.293-1.293a.999.999 0 1 0-1.414 1.414L12 21.414l3.707-3.707a.999.999 0 1 0-1.414-1.414L13 17.586V13h4.586l-1.293 1.293a.999.999 0 1 0 1.414 1.414L21.414 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowMoveOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m22.828 10.586l-9.414-9.414C13.023.781 12.512.586 12 .586s-1.023.195-1.414.586l-9.414 9.414a2 2 0 0 0 0 2.828l9.414 9.414c.391.391.902.586 1.414.586s1.023-.195 1.414-.586l9.414-9.414a2 2 0 0 0 0-2.828M17 16a.999.999 0 0 1-.707-1.707L17.586 13H13v4.586l1.293-1.293a.997.997 0 0 1 1.414 0a.999.999 0 0 1 0 1.414L12 21.414l-3.707-3.707a.999.999 0 1 1 1.414-1.414L11 17.586V13H6.414l1.293 1.293a.999.999 0 1 1-1.414 1.414L2.586 12l3.707-3.707a.997.997 0 0 1 1.414 0a.999.999 0 0 1 0 1.414L6.414 11H11V6.414L9.707 7.707a.997.997 0 0 1-1.414 0a.999.999 0 0 1 0-1.414L12 2.586l3.707 3.707a.999.999 0 1 1-1.414 1.414L13 6.414V11h4.586l-1.293-1.293a.999.999 0 1 1 1.414-1.414L21.414 12l-3.707 3.707A.997.997 0 0 1 17 16m-1.732-2A1.981 1.981 0 0 0 15 15c-.357 0-.699.093-1 .268V14zm-6.536 0H10v1.268A1.981 1.981 0 0 0 9 15c0-.357-.093-.699-.268-1m0-4C8.907 9.699 9 9.357 9 9c.357 0 .699-.093 1-.268V10zm6.536 0H14V8.732c.301.175.643.268 1 .268c0 .357.093.699.268 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRepeat(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.5 7h-2.086l1.293-1.293a.999.999 0 1 0-1.414-1.414L10.586 8l3.707 3.707a.997.997 0 0 0 1.414 0a.999.999 0 0 0 0-1.414L14.414 9H16.5c1.379 0 2.5 1.346 2.5 3s-1.346 3-3 3a1 1 0 1 0 0 2c2.757 0 5-2.243 5-5s-2.019-5-4.5-5m-8.207 5.293a.999.999 0 0 0 0 1.414L9.586 15H7.5C6.121 15 5 13.654 5 12s1.346-3 3-3a1 1 0 1 0 0-2c-2.757 0-5 2.243-5 5s2.019 5 4.5 5h2.086l-1.293 1.293a.999.999 0 1 0 1.414 1.414L13.414 16l-3.707-3.707a.999.999 0 0 0-1.414 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRepeatOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.994 7.187L19 7c0-.801-.312-1.555-.879-2.121C17.555 4.312 16.801 4 16 4s-1.555.312-2.121.879L10.987 7.77A2.974 2.974 0 0 0 9 7H8c-3.859 0-7 3.14-7 7c0 3.306 2.14 6.084 5.006 6.813L6 21c0 .801.312 1.555.879 2.121c.566.567 1.32.879 2.121.879s1.555-.312 2.121-.879l2.892-2.891c.53.473 1.221.77 1.987.77h1c3.859 0 7-3.14 7-7c0-3.306-2.14-6.084-5.006-6.813M17 19h-1a1 1 0 1 1 0-2h1c1.654 0 3-1.346 3-3s-1.121-3-2.5-3h-2.086l1.293 1.293a.999.999 0 1 1-1.414 1.414L11.586 10l3.707-3.707a.997.997 0 0 1 1.414 0a.999.999 0 0 1 0 1.414L15.414 9H17.5c2.481 0 4.5 2.243 4.5 5s-2.243 5-5 5m.749-6.971c.7.164 1.251 1 1.251 1.971c0 1.103-.897 2-2 2h-1c-.857 0-1.584.544-1.868 1.304l-3.718-3.718C10.036 13.208 9.534 13 9 13s-1.036.208-1.414.586S7 14.466 7 15c0 .345.087.677.251.971C6.551 15.807 6 14.971 6 14c0-1.103.897-2 2-2h1c.857 0 1.584-.544 1.868-1.304l3.718 3.718c.378.378.88.586 1.414.586s1.036-.208 1.414-.586S18 13.534 18 13c0-.345-.087-.677-.251-.971M10 10a1 1 0 0 1-1 1H8c-1.654 0-3 1.346-3 3s1.121 3 2.5 3h2.086l-1.293-1.293a.999.999 0 1 1 1.414-1.414L13.414 18l-3.707 3.707a.997.997 0 0 1-1.414 0a.999.999 0 0 1 0-1.414L9.586 19H7.5C5.019 19 3 16.757 3 14s2.243-5 5-5h1a1 1 0 0 1 1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRight(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.293 7.293a.999.999 0 0 0 0 1.414L15.586 11H8a1 1 0 0 0 0 2h7.586l-2.293 2.293a.999.999 0 1 0 1.414 1.414L19.414 12l-4.707-4.707a.999.999 0 0 0-1.414 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21c-.801 0-1.555-.312-2.121-.879S8.999 18.8 9 17.998c0-.746.271-.998.764-1.998H4.928c-1.654 0-3-1.347-3-3c0-1.654 1.346-3 3-3h4.836C9.27 9 9 8.745 9 7.999a2.979 2.979 0 0 1 .88-2.121c1.132-1.132 3.108-1.133 4.241.001L21.242 13l-7.121 7.121A2.978 2.978 0 0 1 12 21m-7.072-9a1.001 1.001 0 0 0 0 2h9.658l-3.293 3.293a.99.99 0 0 0-.293.706c0 .269.104.519.293.708a1.023 1.023 0 0 0 1.414 0L18.414 13l-5.707-5.707a1.023 1.023 0 0 0-1.414 0a.99.99 0 0 0-.293.706c0 .268.104.519.293.708L14.586 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightThick(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.586 6.586a2 2 0 0 0 0 2.828L12.172 11H4.928a2 2 0 0 0 0 4h7.244l-1.586 1.586a2 2 0 1 0 2.828 2.828L19.828 13l-6.414-6.414a2 2 0 0 0-2.828 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowShuffle(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 9h3.5c.736 0 1.393.391 1.851 1.001a8.11 8.11 0 0 1 1.191-1.662C9.739 7.516 8.676 7 7.5 7H4a1 1 0 1 0 0 2m7.685 3.111C12.236 10.454 13.941 9 15.334 9h1.838l-1.293 1.293a.999.999 0 1 0 1.414 1.414L21 8l-3.707-3.707a.999.999 0 1 0-1.414 1.414L17.172 7h-1.838c-2.274 0-4.711 1.967-5.547 4.479l-.472 1.411C8.674 14.816 7.243 16 6.5 16H4a1 1 0 1 0 0 2h2.5c1.837 0 3.863-1.925 4.713-4.479zm4.194 1.182a.999.999 0 0 0 0 1.414L17.172 16h-2.338c-1.268 0-2.33-.891-2.691-2.108a9.335 9.335 0 0 1-1.09 2.185C11.939 17.239 13.296 18 14.834 18h2.338l-1.293 1.293a.999.999 0 1 0 1.414 1.414L21 17l-3.707-3.707a.999.999 0 0 0-1.414 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowSortedDown(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.8 9.7L12 16l6.2-6.3c.2-.2.3-.5.3-.7s-.1-.5-.3-.7c-.2-.2-.4-.3-.7-.3h-11c-.3 0-.5.1-.7.3c-.2.2-.3.4-.3.7s.1.5.3.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowSortedUp(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.2 13.3L12 7l-6.2 6.3c-.2.2-.3.5-.3.7s.1.5.3.7c.2.2.4.3.7.3h11c.3 0 .5-.1.7-.3c.2-.2.3-.5.3-.7s-.1-.5-.3-.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowSync(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.5 12.473c0-1.948-.618-3.397-2.066-4.844a1 1 0 0 0-1.414 1.415c1.079 1.078 1.48 2.007 1.48 3.429a5.46 5.46 0 0 1-1.611 3.888c-1.004 1.003-2.078 1.502-3.428 1.593l1.246-1.247a.999.999 0 1 0-1.414-1.414L8.586 19l3.707 3.707a.997.997 0 0 0 1.414 0a.999.999 0 0 0 0-1.414l-1.337-1.336c1.923-.082 3.542-.792 4.933-2.181a7.447 7.447 0 0 0 2.197-5.303m-13 .027c0-1.469.572-2.85 1.611-3.889c1.009-1.009 2.092-1.508 3.457-1.594l-1.275 1.275A.999.999 0 0 0 11 10a.997.997 0 0 0 .707-.293L15.414 6l-3.707-3.707a.999.999 0 1 0-1.414 1.414l1.311 1.311c-1.914.086-3.525.796-4.907 2.179A7.447 7.447 0 0 0 4.5 12.5c0 1.948.618 3.397 2.066 4.844a.997.997 0 0 0 1.414-.001a1 1 0 0 0 0-1.415C6.901 14.851 6.5 13.922 6.5 12.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowSyncOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.5 12.473c0-2.495-.818-4.426-2.653-6.259a2.99 2.99 0 0 0-1.073-.682l-.946-.946L13.121.879C12.555.312 11.801 0 11 0S9.445.312 8.879.879A2.978 2.978 0 0 0 8 3c0 .277.037.549.11.809a9.445 9.445 0 0 0-2.827 1.974A9.43 9.43 0 0 0 2.5 12.5c0 2.495.818 4.426 2.653 6.259c.299.298.652.521 1.034.669l.985.986l3.707 3.707c.566.567 1.32.879 2.121.879s1.555-.312 2.121-.879c.567-.566.879-1.32.879-2.121c0-.286-.04-.566-.117-.834a9.474 9.474 0 0 0 2.833-1.975a9.432 9.432 0 0 0 2.784-6.718m-9.13 7.484l1.337 1.336a.999.999 0 1 1-1.414 1.414L8.586 19l3.707-3.707a.997.997 0 0 1 1.414 0a.999.999 0 0 1 0 1.414l-1.247 1.247c1.351-.091 2.425-.59 3.428-1.593a5.46 5.46 0 0 0 1.611-3.888c0-1.422-.401-2.351-1.48-3.429a1 1 0 1 1 1.415-1.416c1.448 1.447 2.066 2.896 2.066 4.844c0 2.004-.78 3.887-2.197 5.303c-1.39 1.39-3.01 2.1-4.933 2.182m-.766-14.939l-1.311-1.311a.999.999 0 1 1 1.414-1.414L15.414 6l-3.707 3.707a.997.997 0 0 1-1.414 0a.999.999 0 0 1 0-1.414l1.275-1.275c-1.365.086-2.448.584-3.456 1.593A5.46 5.46 0 0 0 6.5 12.5c0 1.422.401 2.351 1.48 3.429a1 1 0 1 1-1.415 1.416C5.118 15.897 4.5 14.448 4.5 12.5c0-2.004.78-3.887 2.197-5.303c1.382-1.383 2.993-2.093 4.907-2.179M8.688 15.222C7.8 14.335 7.5 13.648 7.5 12.5c0-1.202.468-2.332 1.318-3.181l.187-.179c.033.481.236.93.581 1.274c.378.378.88.586 1.414.586s1.036-.208 1.414-.586l2.339-2.339a1.99 1.99 0 0 0 .56 1.675c.888.887 1.188 1.574 1.188 2.722a4.466 4.466 0 0 1-1.318 3.181l-.188.181a1.987 1.987 0 0 0-.579-1.248C14.036 14.208 13.534 14 13 14s-1.036.208-1.414.586l-2.342 2.342a1.994 1.994 0 0 0-.556-1.706"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUnsorted(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.2 9.3L12 3L5.8 9.3c-.2.2-.3.4-.3.7s.1.5.3.7c.2.2.4.3.7.3h11c.3 0 .5-.1.7-.3c.2-.2.3-.5.3-.7s-.1-.5-.3-.7M5.8 14.7L12 21l6.2-6.3c.2-.2.3-.5.3-.7s-.1-.5-.3-.7c-.2-.2-.4-.3-.7-.3h-11c-.3 0-.5.1-.7.3c-.2.2-.3.5-.3.7s.1.5.3.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUp(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13 5.586l-4.707 4.707a.999.999 0 1 0 1.414 1.414L12 9.414V17a1 1 0 1 0 2 0V9.414l2.293 2.293a.997.997 0 0 0 1.414 0a.999.999 0 0 0 0-1.414z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21c-1.654 0-3-1.346-3-3v-4.764c-1.143 1.024-3.025.979-4.121-.115a3.002 3.002 0 0 1 0-4.242L12 1.758l7.121 7.121a3.002 3.002 0 0 1 0 4.242c-1.094 1.095-2.979 1.14-4.121.115V18c0 1.654-1.346 3-3 3M11 8.414V18a1.001 1.001 0 0 0 2 0V8.414l3.293 3.293a1.023 1.023 0 0 0 1.414 0a.999.999 0 0 0 0-1.414L12 4.586l-5.707 5.707a.999.999 0 0 0 0 1.414a1.023 1.023 0 0 0 1.414 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpThick(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 3.172L5.586 9.586a2 2 0 1 0 2.828 2.828L10 10.828v7.242a2 2 0 0 0 4 0v-7.242l1.586 1.586c.391.391.902.586 1.414.586s1.023-.195 1.414-.586a2 2 0 0 0 0-2.828z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func At(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 4c-4.411 0-8 3.589-8 8s3.589 8 8 8a7.953 7.953 0 0 0 4.499-1.384a1.001 1.001 0 0 0-1.127-1.653A5.951 5.951 0 0 1 12 18c-3.309 0-6-2.691-6-6s2.691-6 6-6s6 2.691 6 6v.5a1 1 0 0 1-2 0v-3a1 1 0 0 0-1-1a.99.99 0 0 0-.938.688A3.466 3.466 0 0 0 12 8.5c-1.93 0-3.5 1.57-3.5 3.5s1.57 3.5 3.5 3.5c1.045 0 1.975-.47 2.616-1.199A2.988 2.988 0 0 0 17 15.5c1.654 0 3-1.346 3-3V12c0-4.411-3.589-8-8-8m0 9.5c-.827 0-1.5-.673-1.5-1.5s.673-1.5 1.5-1.5s1.5.673 1.5 1.5s-.673 1.5-1.5 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Attachment(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.364 6.635a4.007 4.007 0 0 0-5.658 0L8.172 11.17a2.484 2.484 0 0 0-.733 1.77a2.498 2.498 0 0 0 2.501 2.498c.64 0 1.279-.242 1.767-.73l2.122-2.121a2.002 2.002 0 0 0 0-2.828l-3.536 3.535a.5.5 0 0 1-.708-.708l4.535-4.537a2.006 2.006 0 0 1 2.83 0a2.003 2.003 0 0 1 0 2.828l-4.537 4.537l-2.535 2.535a2.003 2.003 0 0 1-2.828 0a2.001 2.001 0 0 1 0-2.828l.095-.096a3.566 3.566 0 0 1-.702-2.125l-.807.807a4.003 4.003 0 0 0 0 5.656c.779.779 1.804 1.17 2.828 1.17s2.049-.391 2.828-1.17l7.072-7.072a4.003 4.003 0 0 0 0-5.656"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AttachmentOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.534 4.466c1.024 0 2.05.39 2.829 1.169a4.001 4.001 0 0 1 0 5.656l-7.071 7.072c-.778.779-1.804 1.17-2.828 1.17s-2.049-.391-2.828-1.17a4.003 4.003 0 0 1 0-5.656l.807-.807c-.004.805.25 1.524.701 2.125l-.094.096a2.001 2.001 0 0 0 0 2.828c.39.39.901.584 1.414.584s1.024-.195 1.414-.584l2.535-2.535l4.537-4.537a2.005 2.005 0 0 0 0-2.828a2.001 2.001 0 0 0-1.417-.584c-.512 0-1.023.195-1.413.584l-4.535 4.537a.5.5 0 1 0 .708.708l2.122-2.121l1.414-1.414c.392.392.586.902.586 1.414c0 .511-.194 1.021-.584 1.41l-2.124 2.125c-.486.487-1.127.729-1.768.729s-1.28-.244-1.769-.729a2.489 2.489 0 0 1-.731-1.769c0-.67.261-1.297.732-1.77l4.534-4.535a3.99 3.99 0 0 1 2.829-1.168m0-2c-1.604 0-3.11.623-4.242 1.755l-7.069 7.073c-1.133 1.131-1.757 2.638-1.757 4.242s.624 3.11 1.757 4.243c1.131 1.132 2.639 1.755 4.241 1.755s3.11-.624 4.242-1.757l7.071-7.071a5.954 5.954 0 0 0 1.757-4.242c0-1.603-.623-3.11-1.755-4.241a5.958 5.958 0 0 0-4.245-1.757"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Backspace(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.5 5h-10c-1.266 0-2.834.807-3.57 1.837L3.32 10.49l-1.199 1.679c-.121.175-.122.492.003.664l1.188 1.664l2.619 3.667C6.666 19.193 8.233 20 9.5 20h10c1.379 0 2.5-1.122 2.5-2.5v-10C22 6.122 20.879 5 19.5 5m-2.293 9.793a.999.999 0 1 1-1.414 1.414L13.5 13.914l-2.293 2.293a.997.997 0 0 1-1.414 0a.999.999 0 0 1 0-1.414l2.293-2.293l-2.293-2.293a.999.999 0 1 1 1.414-1.414l2.293 2.293l2.293-2.293a.999.999 0 1 1 1.414 1.414L14.914 12.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BackspaceOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 21H9c-1.436 0-3.145-.88-3.977-2.046l-2.619-3.667l-1.188-1.661c-.246-.344-.249-.894-.008-1.241l1.204-1.686L5.02 7.046C5.855 5.879 7.566 5 9 5h10c1.654 0 3 1.346 3 3v10c0 1.654-1.346 3-3 3M3.229 12.999l.806 1.125l2.618 3.667C7.104 18.424 8.223 19 9.001 19h10c.552 0 1-.45 1-1.001V8c0-.551-.448-1-1-1h-10c-.776 0-1.897.576-2.351 1.209l-2.608 3.652zM13.707 13l2.646-2.646a.502.502 0 0 0 0-.707a.502.502 0 0 0-.707 0L13 12.293l-2.646-2.646a.5.5 0 0 0-.707.707L12.293 13l-2.646 2.646a.5.5 0 0 0 .707.708L13 13.707l2.646 2.646a.5.5 0 1 0 .708-.706z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryCharge(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 10v6h11v-6zm5.83 4.908L9.62 13L7 13.428l3.223-2.324L11.398 13L14 12.57zM19 10c0-1.654-1.346-3-3-3H5c-1.654 0-3 1.346-3 3v6c0 1.654 1.346 3 3 3h11c1.654 0 3-1.346 3-3a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2m-2 6c0 .552-.449 1-1 1H5c-.551 0-1-.448-1-1v-6c0-.552.449-1 1-1h11c.551 0 1 .448 1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryFull(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 16a1 1 0 0 1-1-1v-4a1 1 0 1 1 2 0v4a1 1 0 0 1-1 1m-3 0a1 1 0 0 1-1-1v-4a1 1 0 1 1 2 0v4a1 1 0 0 1-1 1m9 0a1 1 0 0 1-1-1v-4a1 1 0 1 1 2 0v4a1 1 0 0 1-1 1m-3 0a1 1 0 0 1-1-1v-4a1 1 0 1 1 2 0v4a1 1 0 0 1-1 1m7-6c0-1.654-1.346-3-3-3H5c-1.654 0-3 1.346-3 3v6c0 1.654 1.346 3 3 3h11c1.654 0 3-1.346 3-3a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2m-2 6c0 .552-.449 1-1 1H5c-.551 0-1-.448-1-1v-6c0-.552.449-1 1-1h11c.551 0 1 .448 1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryHigh(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 16a1 1 0 0 1-1-1v-4a1 1 0 1 1 2 0v4a1 1 0 0 1-1 1m-3 0a1 1 0 0 1-1-1v-4a1 1 0 1 1 2 0v4a1 1 0 0 1-1 1m6 0a1 1 0 0 1-1-1v-4a1 1 0 1 1 2 0v4a1 1 0 0 1-1 1m7-6c0-1.654-1.346-3-3-3H5c-1.654 0-3 1.346-3 3v6c0 1.654 1.346 3 3 3h11c1.654 0 3-1.346 3-3a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2m-2 6c0 .552-.449 1-1 1H5c-.551 0-1-.448-1-1v-6c0-.552.449-1 1-1h11c.551 0 1 .448 1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryLow(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 16a1 1 0 0 1-1-1v-4a1 1 0 1 1 2 0v4a1 1 0 0 1-1 1m13-6c0-1.654-1.346-3-3-3H5c-1.654 0-3 1.346-3 3v6c0 1.654 1.346 3 3 3h11c1.654 0 3-1.346 3-3a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2m-2 6c0 .552-.449 1-1 1H5c-.551 0-1-.448-1-1v-6c0-.552.449-1 1-1h11c.551 0 1 .448 1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryMid(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 16a1 1 0 0 1-1-1v-4a1 1 0 1 1 2 0v4a1 1 0 0 1-1 1m-3 0a1 1 0 0 1-1-1v-4a1 1 0 1 1 2 0v4a1 1 0 0 1-1 1m13-6c0-1.654-1.346-3-3-3H5c-1.654 0-3 1.346-3 3v6c0 1.654 1.346 3 3 3h11c1.654 0 3-1.346 3-3a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2m-2 6c0 .552-.449 1-1 1H5c-.551 0-1-.448-1-1v-6c0-.552.449-1 1-1h11c.551 0 1 .448 1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Beaker(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.445 16.809L16.805 7H18a1 1 0 0 0 0-2H6a1 1 0 0 0 0 2h1.135a3.605 3.605 0 0 1-.121.671l-2.459 9.138c-.218.809-.074 1.623.393 2.231c.466.61 1.214.96 2.052.96h10c.838 0 1.586-.35 2.055-.959c.466-.609.609-1.423.39-2.232M14.732 7l1.352 5.018L16 12H8l-.084.018l1.029-3.826c.084-.312.173-.744.192-1.192zm2.734 10.824c-.087.114-.252.176-.466.176H7c-.214 0-.379-.062-.466-.176c-.086-.113-.104-.289-.048-.496l1.197-4.45A.494.494 0 0 0 8 13h8a.493.493 0 0 0 .316-.121l1.197 4.45c.057.206.04.382-.047.495"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Beer(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 16.5c0 .275-.225.5-.5.5s-.5-.225-.5-.5v-6c0-.275.225-.5.5-.5s.5.225.5.5zm2 0c0 .275-.225.5-.5.5s-.5-.225-.5-.5v-6c0-.275.225-.5.5-.5s.5.225.5.5zm2 0c0 .275-.225.5-.5.5s-.5-.225-.5-.5v-6c0-.275.225-.5.5-.5s.5.225.5.5zM18.5 6H18V5a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v13c0 1.656 1.344 3 3 3h7c1.656 0 3-1.344 3-3h.5c1.93 0 3.5-1.57 3.5-3.5v-5C22 7.57 20.43 6 18.5 6M7 5h9v1h-4.444l-.118.332c-.164.458-.663.73-1.117.648l-.348-.058l-.173.307A1.499 1.499 0 0 1 8.5 8C7.673 8 7 7.327 7 6.5zm9 13a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1V8.49c.419.317.936.51 1.5.51c.784 0 1.521-.376 1.989-1c.728 0 1.383-.391 1.736-1H16zm4-3.5c0 .827-.673 1.5-1.5 1.5H17V8h1.5c.827 0 1.5.673 1.5 1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bell(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.715 17.301c-.017-.018-1.717-1.854-1.73-6.32a6.01 6.01 0 0 0-4.019-5.641L14 5c0-1.103-.896-2-2-2s-2 .897-2 2l.034.338c-2.336.816-4.019 3.036-4.019 5.646c0 4.462-1.711 6.296-1.721 6.306A1.002 1.002 0 0 0 5 19h3.143c.447 1.72 1.999 3 3.857 3s3.41-1.28 3.857-3H19c.4 0 .758-.243.915-.61s.076-.799-.2-1.089M12 7c2.189 0 3.978 1.789 3.984 3.987c.002.728.046 1.396.118 2.013h-8.2c.071-.617.113-1.286.113-2.016A3.99 3.99 0 0 1 12 7m0 13a1.993 1.993 0 0 1-1.722-1h3.443A1.99 1.99 0 0 1 12 20m-5.186-3c.352-.736.705-1.731.938-3h8.502c.234 1.269.588 2.264.938 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Book(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 3H7a.997.997 0 0 0-.707.293l-3 3l-.057.062a.996.996 0 0 0-.235.6L3 7.008V18c0 1.654 1.346 3 3 3h9a3.006 3.006 0 0 0 2.829-2h.671c1.402 0 2.5-1.317 2.5-3V6c0-1.654-1.346-3-3-3M6 19c-.551 0-1-.448-1-1V8h2v11zm10-1c0 .552-.449 1-1 1H8V8h7c.551 0 1 .448 1 1zm3-2c0 .62-.324 1-.5 1H18V9c0-1.654-1.346-3-3-3H6.414l1-1H18c.551 0 1 .448 1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookmark(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2H9C7.346 2 6 3.346 6 5v14c0 .514.104.946.308 1.285c.564.935 1.815 1.008 2.813.008l3.172-3.172a1.031 1.031 0 0 1 1.414 0l3.172 3.172c.491.491 1.002.74 1.52.74c.797 0 1.601-.629 1.601-2.033V5c0-1.654-1.346-3-3-3M9 4h8c.551 0 1 .449 1 1v9.905l-2.451-2.247c-1.406-1.289-3.693-1.288-5.099 0L8 14.905V5c0-.551.449-1 1-1m6.121 11.707c-.565-.565-1.318-.876-2.121-.876s-1.556.312-2.121.876L8 18.586v-2.324l3.126-2.866c1.033-.947 2.714-.947 3.747 0L18 16.262v2.324z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Briefcase(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 7c0-1.654-1.346-3-3-3H9C7.346 4 6 5.346 6 7c-1.654 0-3 1.346-3 3v7c0 1.654 1.346 3 3 3h12c1.654 0 3-1.346 3-3v-7c0-1.654-1.346-3-3-3M9 6h6c.551 0 1 .449 1 1H8c0-.551.449-1 1-1m10 11c0 .551-.449 1-1 1H6c-.551 0-1-.449-1-1v-1h14zM5 15v-5c0-.551.449-1 1-1h12c.551 0 1 .449 1 1v5zm8-3h-2c-.55 0-1 .45-1 1s.45 1 1 1h2c.55 0 1-.45 1-1s-.45-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Brush(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.177 3.823a1.997 1.997 0 0 0-2.159-.442c-4.449 1.787-7.792 4.76-10.517 9.357a1.962 1.962 0 0 0-.209.542c-1.38.215-2.6.903-3.442 1.993c-.916 1.185-1.295 2.695-1.066 4.254L3 21l1.473.217c.293.043.589.064.88.064c2.743 0 4.949-1.909 5.367-4.564a2 2 0 0 0 .544-.218c4.598-2.728 7.571-6.069 9.355-10.517a2.003 2.003 0 0 0-.442-2.159M5.353 19.281c-.192 0-.389-.016-.59-.044c-.309-2.104 1.055-3.81 3-4.021l1.021 1.021c-.192 1.76-1.605 3.044-3.431 3.044m4.89-4.502l-1.021-1.021c.38-.641.774-1.233 1.178-1.804c.027.041 1.639 1.653 1.639 1.653c-.568.401-1.158.794-1.796 1.172m2.608-1.773s-1.821-1.801-1.879-1.824c2.147-2.784 4.651-4.685 7.791-5.943c-1.255 3.127-3.144 5.623-5.912 7.767"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BusinessCard(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 20H4c-1.654 0-3-1.346-3-3V7c0-1.654 1.346-3 3-3h16c1.654 0 3 1.346 3 3v10c0 1.654-1.346 3-3 3M4 6c-.551 0-1 .449-1 1v10c0 .551.449 1 1 1h16c.551 0 1-.449 1-1V7c0-.551-.449-1-1-1zm6 9H6a1 1 0 1 1 0-2h4a1 1 0 1 1 0 2m0-4H6a1 1 0 1 1 0-2h4a1 1 0 1 1 0 2"/><circle cx="16" cy="10.5" r="2" fill="currentColor"/><path fill="currentColor" d="M16 13.356c-1.562 0-2.5.715-2.5 1.429c0 .357.938.715 2.5.715c1.466 0 2.5-.357 2.5-.715c0-.714-.98-1.429-2.5-1.429"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calculator(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 21H9c-1.7 0-3-1.3-3-3V6c0-1.7 1.3-3 3-3h8c1.7 0 3 1.3 3 3v12c0 1.7-1.3 3-3 3M9 5c-.6 0-1 .4-1 1v12c0 .6.4 1 1 1h8c.6 0 1-.4 1-1V6c0-.6-.4-1-1-1z"/><circle cx="10" cy="11" r="1" fill="currentColor"/><circle cx="13" cy="11" r="1" fill="currentColor"/><circle cx="16" cy="11" r="1" fill="currentColor"/><circle cx="10" cy="14" r="1" fill="currentColor"/><circle cx="13" cy="14" r="1" fill="currentColor"/><circle cx="16" cy="14" r="1" fill="currentColor"/><circle cx="10" cy="17" r="1" fill="currentColor"/><circle cx="13" cy="17" r="1" fill="currentColor"/><circle cx="16" cy="17" r="1" fill="currentColor"/><path fill="currentColor" d="M16 7v1h-6V7zm1-1H9v3h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 6.184V6a3 3 0 1 0-6 0h-2a3 3 0 1 0-6 0v.184A2.997 2.997 0 0 0 3 9v9c0 1.654 1.346 3 3 3h12c1.654 0 3-1.346 3-3V9a2.997 2.997 0 0 0-2-2.816M15 6a1 1 0 1 1 2 0v2a1 1 0 1 1-2 0zM7 6a1 1 0 1 1 2 0v2a1 1 0 1 1-2 0zm12 12c0 .551-.448 1-1 1H6c-.552 0-1-.449-1-1v-6h14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 6.184V6a3 3 0 1 0-6 0h-2a3 3 0 1 0-6 0v.184A2.997 2.997 0 0 0 3 9v9c0 1.654 1.346 3 3 3h12c1.654 0 3-1.346 3-3V9a2.997 2.997 0 0 0-2-2.816M15 6a1 1 0 1 1 2 0v2a1 1 0 1 1-2 0zM7 6a1 1 0 1 1 2 0v2a1 1 0 1 1-2 0zm12 12c0 .551-.448 1-1 1H6c-.552 0-1-.449-1-1v-6h14zm0-7H5V9c0-.551.448-1 1-1a2 2 0 0 0 4 0h4a2 2 0 0 0 4 0c.552 0 1 .449 1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Camera(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 6h-1.586l-1-1c-.579-.579-1.595-1-2.414-1h-4c-.819 0-1.835.421-2.414 1l-1 1H5C3.346 6 2 7.346 2 9v8c0 1.654 1.346 3 3 3h14c1.654 0 3-1.346 3-3V9c0-1.654-1.346-3-3-3m-7 10a3.5 3.5 0 1 1 .001-7.001A3.5 3.5 0 0 1 12 16m6-4.701a1.3 1.3 0 1 1 0-2.6a1.3 1.3 0 0 1 0 2.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 20H5c-1.654 0-3-1.346-3-3V9c0-1.654 1.346-3 3-3h1.586l1-1C8.165 4.421 9.182 4 10 4h4c.818 0 1.835.421 2.414 1l1 1H19c1.654 0 3 1.346 3 3v8c0 1.654-1.346 3-3 3M5 8a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V9a1 1 0 0 0-1-1h-2a.996.996 0 0 1-.707-.293L15 6.414C14.799 6.213 14.285 6 14 6h-4c-.285 0-.799.213-1 .414L7.707 7.707A.996.996 0 0 1 7 8zm7 2c1.379 0 2.5 1.121 2.5 2.5S13.379 15 12 15s-2.5-1.121-2.5-2.5S10.621 10 12 10m0-1a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7m6-.301a1.3 1.3 0 1 0 0 2.6a1.3 1.3 0 0 0 0-2.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cancel(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 4c-4.411 0-8 3.589-8 8s3.589 8 8 8s8-3.589 8-8s-3.589-8-8-8m-5 8c0-.832.224-1.604.584-2.295l6.711 6.711A4.943 4.943 0 0 1 12 17c-2.757 0-5-2.243-5-5m9.416 2.295L9.705 7.584A4.943 4.943 0 0 1 12 7c2.757 0 5 2.243 5 5c0 .832-.224 1.604-.584 2.295"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CancelOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 20.5c-4.688 0-8.5-3.812-8.5-8.5s3.812-8.5 8.497-8.5c4.69 0 8.503 3.812 8.503 8.5s-3.812 8.5-8.5 8.5m0-15c-3.586 0-6.5 2.916-6.5 6.5s2.916 6.5 6.5 6.5s6.5-2.916 6.5-6.5s-2.916-6.5-6.5-6.5m.003 3A3.502 3.502 0 0 1 15.5 12c0 .206-.02.412-.057.615l-4.057-4.059c.203-.036.408-.056.614-.056m.003-1a4.48 4.48 0 0 0-2.39.697l6.188 6.188a4.448 4.448 0 0 0 .699-2.387A4.5 4.5 0 0 0 12.003 7.5m-3.446 3.884l4.059 4.06A3.504 3.504 0 0 1 8.5 11.998c0-.206.02-.412.057-.614m-.358-1.773a4.48 4.48 0 0 0-.699 2.387A4.502 4.502 0 0 0 12 16.5a4.48 4.48 0 0 0 2.387-.699z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartArea(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 6a2 2 0 0 0-3.562-1.25l-2.789 3.486L11.2 6.4a2 2 0 0 0-2.762.351l-4 5A1.998 1.998 0 0 0 4 13v3h16zm0 13H4a1 1 0 1 0 0 2h16a1 1 0 1 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartAreaOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 17H4a1 1 0 0 1-1-1v-3c0-.68.234-1.346.658-1.874l4-5c.98-1.226 2.885-1.469 4.143-.524l1.674 1.254l2.185-2.729A2.974 2.974 0 0 1 18.001 3A3 3 0 0 1 21 6v10a1 1 0 0 1-1 1M5 15h14V6a.988.988 0 0 0-.375-.779a.996.996 0 0 0-1.406.155L14.43 8.861a1 1 0 0 1-1.381.176L10.6 7.2a1.02 1.02 0 0 0-1.381.176l-4 5A.993.993 0 0 0 5 13zm15 6H4a1 1 0 1 1 0-2h16a1 1 0 1 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartBar(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 4a2 2 0 0 0-4 0v12h4zm5 4a2 2 0 0 0-4 0v8h4zM9 11a2 2 0 0 0-4 0v5h4zm10 8H5a1 1 0 1 0 0 2h14a1 1 0 1 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartBarOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 5c-.771 0-1.468.301-2 .779V4c0-1.654-1.346-3-3-3S9 2.346 9 4v4.779A2.985 2.985 0 0 0 7 8c-1.654 0-3 1.346-3 3v6h16V8c0-1.654-1.346-3-3-3m-5-2c.551 0 1 .448 1 1v11h-2V4c0-.552.449-1 1-1M8 15H6v-4a1.001 1.001 0 0 1 2 0zm10 0h-2V8a1.001 1.001 0 0 1 2 0zm1 6H5a1 1 0 1 1 0-2h14a1 1 0 1 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartLine(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.75 15.561a2 2 0 0 0 2.811-.313l2.789-3.486L12.8 13.6a2 2 0 0 0 2.762-.352l4-5a2 2 0 0 0-3.124-2.499l-2.789 3.486L11.2 7.4a2 2 0 0 0-2.762.35l-4 5a2.001 2.001 0 0 0 .312 2.811M5 21h14a1 1 0 1 0 0-2H5a1 1 0 1 0 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartLineOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.999 17a3.014 3.014 0 0 1-1.873-.658a2.978 2.978 0 0 1-1.107-2.011a2.979 2.979 0 0 1 .639-2.206l4-5c.978-1.225 2.883-1.471 4.143-.523l1.674 1.254l2.184-2.729a3 3 0 1 1 4.682 3.747l-4 5c-.977 1.226-2.882 1.471-4.143.526l-1.674-1.256l-2.184 2.729A2.977 2.977 0 0 1 5.999 17M10 8a.997.997 0 0 0-.781.374l-4 5.001a.99.99 0 0 0-.213.734c.03.266.161.504.369.67a.996.996 0 0 0 1.406-.155l3.395-4.244L13.4 12.8c.42.316 1.056.231 1.381-.176l4-5.001a.992.992 0 0 0 .213-.734a.994.994 0 0 0-.369-.67a.996.996 0 0 0-1.406.156l-3.395 4.242L10.6 8.2A.986.986 0 0 0 10 8m9 13H5a1 1 0 1 1 0-2h14a1 1 0 1 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartPie(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m11.614 13.98l4.908 4.922c.39.391.99.36 1.286-.106a8.99 8.99 0 0 0 1.393-4.815a9.005 9.005 0 0 0-1.972-5.631zM9 14.396V7.041a7.008 7.008 0 0 0-6 6.939C3 17.856 6.134 21 10 21a6.946 6.946 0 0 0 4.186-1.403zm7.331-8.183c.39-.391.365-.999-.089-1.313a10.925 10.925 0 0 0-4.251-1.765c-.544-.1-.991.312-.991.865v7.56z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartPieOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m18.227 7.609l.557-.559a2.002 2.002 0 0 0-.135-2.947a13.513 13.513 0 0 0-7.469-3.097L11 1a2.01 2.01 0 0 0-1.35.523C9.236 1.902 9 2.438 9 3v2.229c-3.657.865-6.333 4.188-6.333 8.006c0 4.547 3.688 8.244 8.224 8.244c1.594 0 3.11-.479 4.441-1.345c.277.142.583.226.9.226l.109-.004a1.996 1.996 0 0 0 1.453-.75a10.138 10.138 0 0 0 2.204-6.297a10.099 10.099 0 0 0-1.771-5.7m-7.336 11.87c-3.438 0-6.224-2.793-6.224-6.244A6.229 6.229 0 0 1 10 7.071v6.408l4.609 4.754a6.169 6.169 0 0 1-3.718 1.246M11 12.025V3a11.535 11.535 0 0 1 6.366 2.641zm.214 1.269l5.019-5.028a8.075 8.075 0 0 1 1.769 5.043a8.066 8.066 0 0 1-1.769 5.051z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeft(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.414 5.586a2 2 0 0 0-2.828 0L5.171 12l6.415 6.414c.39.391.902.586 1.414.586s1.024-.195 1.414-.586a2 2 0 0 0 0-2.828L10.829 12l3.585-3.586a2 2 0 0 0 0-2.828"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeftOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 20a2.98 2.98 0 0 1-2.122-.879L3.757 12l7.122-7.121c1.133-1.133 3.11-1.133 4.243 0C15.688 5.445 16 6.199 16 7s-.312 1.555-.879 2.122L12.243 12l2.878 2.879c.567.566.879 1.32.879 2.121s-.312 1.555-.879 2.122A2.98 2.98 0 0 1 13 20m-6.415-8l5.708 5.707a1.024 1.024 0 0 0 1.414 0c.189-.189.293-.439.293-.707s-.104-.518-.293-.707L9.415 12l4.292-4.293c.189-.189.293-.44.293-.707s-.104-.518-.293-.707a1.023 1.023 0 0 0-1.414-.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRight(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.586 5.586a2 2 0 0 0 0 2.828L12.171 12l-3.585 3.586a2 2 0 1 0 2.828 2.828L17.829 12l-6.415-6.414a2 2 0 0 0-2.828 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRightOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 20a2.98 2.98 0 0 1-2.122-.879C7.312 18.555 7 17.801 7 17s.312-1.555.879-2.122L10.757 12L7.879 9.121C7.312 8.555 7 7.801 7 7s.312-1.555.879-2.122c1.133-1.132 3.109-1.133 4.243.001L19.243 12l-7.122 7.121A2.976 2.976 0 0 1 10 20m0-14a.995.995 0 0 0-1 1c0 .267.104.518.293.707L13.585 12l-4.292 4.293C9.104 16.482 9 16.732 9 17s.104.518.293.707a1.023 1.023 0 0 0 1.414.001L16.415 12l-5.708-5.707A.991.991 0 0 0 10 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clipboard(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 3H7C5.346 3 4 4.346 4 6v12c0 1.654 1.346 3 3 3h10c1.654 0 3-1.346 3-3V6c0-1.654-1.346-3-3-3M9 5h6v1c0 .551-.449 1-1 1h-4c-.551 0-1-.449-1-1zm9 13c0 .551-.449 1-1 1H7c-.551 0-1-.449-1-1V6c0-.551.449-1 1-1h1v1c0 1.1.9 2 2 2h4c1.1 0 2-.9 2-2V5h1c.551 0 1 .449 1 1zm-2-1H8a.5.5 0 0 1 0-1h8a.5.5 0 0 1 0 1m0-3H8a.5.5 0 0 1 0-1h8a.5.5 0 0 1 0 1m0-3H8a.5.5 0 0 1 0-1h8a.5.5 0 0 1 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudStorage(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m17 9l-.351.015A5.967 5.967 0 0 0 11 5c-3.309 0-6 2.691-6 6l.001.126A4.007 4.007 0 0 0 2 15c0 2.206 1.794 4 4 4h5v-4.586l-1.293 1.293a.997.997 0 0 1-1.414 0a.999.999 0 0 1 0-1.414l2.999-2.999a1 1 0 0 1 1.416 0l2.999 2.999a.999.999 0 1 1-1.414 1.414L13 14.414V19h4c2.757 0 5-2.243 5-5s-2.243-5-5-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudStorageOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16.5 8.999l-.352.015c-.824-2.375-3.312-4.015-5.898-4.015c-3.309 0-6.25 2.69-6.25 6v.126C3 11.57 1.25 13.139 1.25 15c0 2.206 2.044 4 4.25 4h11c2.757 0 5-2.244 5-5s-2.243-5.001-5-5.001m0 8.001H12v-3.794l2.146 2.146a.502.502 0 0 0 .708 0a.5.5 0 0 0 0-.707l-2.998-3l-.164-.107a.499.499 0 0 0-.383 0l-.162.107l-3 3a.502.502 0 0 0 0 .707a.504.504 0 0 0 .708 0L11 13.206V17H5.5c-1.104 0-2-.896-2-2s.896-2 1.908-2.005l1.422.015l-.248-1.201A4.004 4.004 0 0 1 10.5 7a3.98 3.98 0 0 1 3.93 3.334l.187 1.102l1.073-.306c.312-.089.569-.13.812-.13c1.653 0 3 1.346 3 3s-1.348 3-3.002 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Code(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.171 18a1.992 1.992 0 0 1-1.414-.586L2.343 13l4.414-4.414a2 2 0 1 1 2.828 2.828L8 13l1.585 1.586A2 2 0 0 1 8.171 18m7.658 0a2 2 0 0 1-1.414-3.414L16 13l-1.585-1.586a2 2 0 1 1 2.828-2.828L21.657 13l-4.414 4.414a1.99 1.99 0 0 1-1.414.586"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.828 19a2.978 2.978 0 0 1-2.121-.879L.586 13l5.121-5.121C6.84 6.745 8.819 6.745 9.95 7.88a3.002 3.002 0 0 1-.001 4.241L9.071 13l.878.879A3.002 3.002 0 0 1 7.828 19m-4.414-6l3.707 3.707c.38.379 1.039.377 1.413.001a1.001 1.001 0 0 0 .001-1.415L6.243 13l2.292-2.293a1 1 0 0 0 0-1.414a1.023 1.023 0 0 0-1.414 0zm12.758 6a3.002 3.002 0 0 1-2.121-5.121l.878-.879l-.878-.879a3.002 3.002 0 0 1 0-4.242c1.129-1.133 3.109-1.134 4.242 0L23.414 13l-5.121 5.121a2.978 2.978 0 0 1-2.121.879m-.001-10a1.001 1.001 0 0 0-.706 1.707L17.757 13l-2.292 2.293a1 1 0 0 0 0 1.414a1.022 1.022 0 0 0 1.414 0L20.586 13l-3.707-3.707A.997.997 0 0 0 16.171 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Coffee(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 19H5a1 1 0 1 1 0-2h12a1 1 0 1 1 0 2m.5-14H5v9c0 1.1.9 2 2 2h8c1.1 0 2-.9 2-2v-2h.5c1.93 0 3.5-1.57 3.5-3.5S19.43 5 17.5 5M15 14H7V7h8zm2.5-4H16V7h1.5c.827 0 1.5.673 1.5 1.5s-.673 1.5-1.5 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cog(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9.387 17.548l.371 1.482c.133.533.692.97 1.242.97h1c.55 0 1.109-.437 1.242-.97l.371-1.482a.961.961 0 0 1 1.203-.694l1.467.42c.529.151 1.188-.114 1.462-.591l.5-.867c.274-.477.177-1.179-.219-1.562l-1.098-1.061a.96.96 0 0 1 .001-1.39l1.096-1.061c.396-.382.494-1.084.22-1.561l-.501-.867c-.275-.477-.933-.742-1.461-.591l-1.467.42a.963.963 0 0 1-1.204-.694l-.37-1.48C13.109 5.437 12.55 5 12 5h-1c-.55 0-1.109.437-1.242.97l-.37 1.48a.964.964 0 0 1-1.204.695l-1.467-.42c-.529-.152-1.188.114-1.462.59l-.5.867c-.274.477-.177 1.179.22 1.562l1.096 1.059a.965.965 0 0 1 0 1.391l-1.098 1.061c-.395.383-.494 1.085-.219 1.562l.501.867c.274.477.933.742 1.462.591l1.467-.42a.96.96 0 0 1 1.203.693M11.5 10.5a2 2 0 1 1 0 4a2 2 0 0 1 0-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CogOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13 5l.855 3.42l3.389-.971l1.501 2.6l-2.535 2.449l2.535 2.451l-1.5 2.6l-3.39-.971L13 20h-3l-.855-3.422l-3.39.971l-1.501-2.6l2.535-2.451l-2.534-2.449l1.5-2.6l3.39.971L10 5zm0-2h-3c-.918 0-1.718.625-1.939 1.516l-.354 1.412l-1.4-.4a2 2 0 0 0-2.283.922l-1.5 2.6a2 2 0 0 0 .342 2.438l1.047 1.011l-1.048 1.015a2 2 0 0 0-.343 2.438l1.502 2.6a1.997 1.997 0 0 0 2.283.924l1.399-.401l.354 1.415A2 2 0 0 0 10 22h3c.918 0 1.718-.625 1.939-1.516l.354-1.414l1.399.4a2 2 0 0 0 2.283-.923l1.5-2.6c.459-.796.317-1.8-.342-2.438l-1.047-1.013l1.047-1.013a2 2 0 0 0 .342-2.438l-1.5-2.6a2 2 0 0 0-2.283-.924l-1.4.401l-.354-1.413A1.997 1.997 0 0 0 13 3m-1.5 7.5a2 2 0 1 1 0 4a2 2 0 0 1 0-4m0-1c-1.654 0-3 1.346-3 3s1.346 3 3 3s3-1.346 3-3s-1.346-3-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compass(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 5c3.859.001 7 3.142 7 7.001c0 3.858-3.141 6.998-7 6.999c-3.859 0-7-3.14-7-6.999s3.141-7 7-7.001m0-2a9 9 0 0 0 0 18a9 9 0 0 0 0-18m4.182 4.819a.498.498 0 0 0-.491-.127L9.74 9.398a.498.498 0 0 0-.342.343l-1.707 5.951a.496.496 0 0 0 .481.637l.138-.02l5.95-1.708a.498.498 0 0 0 .342-.343l1.707-5.949a.498.498 0 0 0-.127-.49M8.9 15.101l1.383-4.817l3.434 3.435z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Contacts(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 3H8C6.346 3 5 4.346 5 6v1H4a1 1 0 1 0 0 2h1v2H4a1 1 0 1 0 0 2h1v2H4a1 1 0 1 0 0 2h1v1c0 1.654 1.346 3 3 3h11c1.654 0 3-1.346 3-3V6c0-1.654-1.346-3-3-3M7 6c0-.551.449-1 1-1v2H7zm0 3h1v2H7zm0 4h1v2H7zm0 5v-1h1v2c-.551 0-1-.449-1-1m13 0c0 .551-.449 1-1 1H9V5h10c.551 0 1 .449 1 1z"/><circle cx="14" cy="10.5" r="2" fill="currentColor"/><path fill="currentColor" d="M14 13.356c-1.562 0-2.5.715-2.5 1.429c0 .357.938.715 2.5.715c1.466 0 2.5-.357 2.5-.715c0-.714-.98-1.429-2.5-1.429"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCard(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 7H6c-1.654 0-3 1.346-3 3v7c0 1.654 1.346 3 3 3h11c1.654 0 3-1.346 3-3v-7c0-1.654-1.346-3-3-3m1 10a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1v-4h13zm0-6H5v-1a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1zm-4 5h2a.5.5 0 0 0 0-1h-2a.5.5 0 0 0 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CssThree(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.7 3.4l-.6 3.2h12.3L17 8.7H4.7l-.6 3.2h12.3l-.7 3.6l-5 1.7l-4.3-1.7l.3-1.6h-3L3 17.7l7.1 2.9l8.2-2.9l1.1-5.8l.2-1.2L21 3.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Database(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.983 8.791C18.807 5.087 15.747 2.125 12 2.125S5.193 5.087 5.017 8.791L5 8.875v6.25c0 3.86 3.141 7 7 7s7-3.14 7-7v-6.25zM12 17.625a5.502 5.502 0 0 1-5-3.222v-.388a6.979 6.979 0 0 0 10 0v.388a5.502 5.502 0 0 1-5 3.222m0-13.5c2.757 0 5 2.243 5 5s-2.243 5-5 5s-5-2.243-5-5s2.243-5 5-5m0 16c-2.271 0-4.172-1.532-4.778-3.609C8.41 17.809 10.11 18.625 12 18.625s3.59-.816 4.778-2.109c-.606 2.077-2.507 3.609-4.778 3.609"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Delete(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 4c-4.419 0-8 3.582-8 8s3.581 8 8 8s8-3.582 8-8s-3.581-8-8-8m3.707 10.293a.999.999 0 1 1-1.414 1.414L12 13.414l-2.293 2.293a.997.997 0 0 1-1.414 0a.999.999 0 0 1 0-1.414L10.586 12L8.293 9.707a.999.999 0 1 1 1.414-1.414L12 10.586l2.293-2.293a.999.999 0 1 1 1.414 1.414L13.414 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeleteOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 3c-4.963 0-9 4.038-9 9s4.037 9 9 9s9-4.038 9-9s-4.037-9-9-9m0 16c-3.859 0-7-3.14-7-7s3.141-7 7-7s7 3.14 7 7s-3.141 7-7 7m.707-7l2.646-2.646a.502.502 0 0 0 0-.707a.502.502 0 0 0-.707 0L12 11.293L9.354 8.646a.5.5 0 0 0-.707.707L11.293 12l-2.646 2.646a.5.5 0 0 0 .707.708L12 12.707l2.646 2.646a.5.5 0 1 0 .708-.706z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeviceDesktop(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 1H3C1.346 1 0 2.346 0 4v11c0 1.654 1.346 3 3 3h6v2H6a1 1 0 0 0 0 2h12a1 1 0 0 0 0-2h-3v-2h6c1.654 0 3-1.346 3-3V4c0-1.654-1.346-3-3-3m-7 19h-4v-2h4zm8-5c0 .551-.449 1-1 1H3c-.551 0-1-.449-1-1V4c0-.551.449-1 1-1h18c.551 0 1 .449 1 1zM20 4H4c-.55 0-1 .45-1 1v8c0 .55.45 1 1 1h16c.55 0 1-.45 1-1V5c0-.55-.45-1-1-1m0 9H4V5h16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeviceLaptop(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.989 16.049c.009-.315.011-.657.011-1.049V6c0-1.654-1.346-3-3-3H5C3.346 3 2 4.346 2 6v9c0 .385.002.73.012 1.049A2.504 2.504 0 0 0 0 18.5C0 19.878 1.122 21 2.5 21h19c1.378 0 2.5-1.122 2.5-2.5a2.504 2.504 0 0 0-2.011-2.451M4 6c0-.551.449-1 1-1h14c.551 0 1 .449 1 1v9c0 .388-.005.726-.014 1H19V7c0-.55-.45-1-1-1H6c-.55 0-1 .45-1 1v9h-.98c-.012-.264-.02-.599-.02-1zm14 10H6V7h12zm3.5 3h-19c-.271 0-.5-.229-.5-.5s.229-.5.5-.5h19c.271 0 .5.229.5.5s-.229.5-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DevicePhone(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 3H8C6.346 3 5 4.346 5 6v12c0 1.654 1.346 3 3 3h7c1.654 0 3-1.346 3-3V6c0-1.654-1.346-3-3-3m1 15c0 .551-.449 1-1 1H8c-.551 0-1-.449-1-1V6c0-.551.449-1 1-1h7c.551 0 1 .449 1 1zM14 6H9c-.55 0-1 .45-1 1v9c0 .55.45 1 1 1h1.5a1 1 0 1 0 2 0H14c.55 0 1-.45 1-1V7c0-.55-.45-1-1-1m0 10H9V7h5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeviceTablet(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 4H8c-.55 0-1 .45-1 1v12c0 .55.45 1 1 1h3.5a1 1 0 1 0 2 0H17c.55 0 1-.45 1-1V5c0-.55-.45-1-1-1m0 13H8V5h9zm1-16H7C5.346 1 4 2.346 4 4v15c0 1.654 1.346 3 3 3h11c1.654 0 3-1.346 3-3V4c0-1.654-1.346-3-3-3m1 18c0 .551-.449 1-1 1H7c-.551 0-1-.449-1-1V4c0-.551.449-1 1-1h11c.551 0 1 .449 1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Directions(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.908 9.5l-2.754-2.607A3.417 3.417 0 0 0 15.917 6H13v-.5a1.5 1.5 0 0 0-3 0V6H6.5C4.57 6 3 7.57 3 9.5c0 1.396.828 2.596 2.016 3.157L3.172 14.5l2.561 2.561c.57.57 1.46.939 2.268.939h2.2l.8 4h1l.8-4h2.7c1.931 0 3.5-1.57 3.5-3.5c0-.902-.353-1.718-.915-2.339l.072-.056zM15.5 16H8c-.275 0-.658-.158-.854-.354L6 14.5l1.146-1.146C7.341 13.159 7.723 13 8 13h7.5a1.5 1.5 0 0 1 0 3m1.279-5.344c-.199.19-.586.344-.862.344H6.5a1.5 1.5 0 0 1 0-3h9.417c.276 0 .663.154.862.344L18 9.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Divide(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="12" cy="6" r="2.25" fill="currentColor"/><circle cx="12" cy="18" r="2.25" fill="currentColor"/><path fill="currentColor" d="M18 10H6a2 2 0 0 0 0 4h12a2 2 0 0 0 0-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DivideOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 8.5c-1.654 0-3-1.346-3-3s1.346-3 3-3s3 1.346 3 3s-1.346 3-3 3m0-4a1.001 1.001 0 0 0 0 2a1.001 1.001 0 0 0 0-2m0 17c-1.654 0-3-1.346-3-3s1.346-3 3-3s3 1.346 3 3s-1.346 3-3 3m0-4a1.001 1.001 0 0 0 0 2a1.001 1.001 0 0 0 0-2m6-2.5H6c-1.654 0-3-1.346-3-3s1.346-3 3-3h12c1.654 0 3 1.346 3 3s-1.346 3-3 3M6 11a1.001 1.001 0 0 0 0 2h12a1.001 1.001 0 0 0 0-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Document(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m19.707 7.293l-4-4A.996.996 0 0 0 15 3H7C5.346 3 4 4.346 4 6v12c0 1.654 1.346 3 3 3h10c1.654 0 3-1.346 3-3V8a.996.996 0 0 0-.293-.707M17.586 8H16.5c-.827 0-1.5-.673-1.5-1.5V5.414zM17 19H7a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h7v1.5C14 7.879 15.121 9 16.5 9H18v9a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentAdd(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 12h-2v-2a1 1 0 1 0-2 0v2H9a1 1 0 1 0 0 2h2v2a1 1 0 1 0 2 0v-2h2a1 1 0 1 0 0-2m4.707-4.707l-4-4A.996.996 0 0 0 15 3H7C5.346 3 4 4.346 4 6v12c0 1.654 1.346 3 3 3h10c1.654 0 3-1.346 3-3V8a.996.996 0 0 0-.293-.707M17.586 8H16.5c-.827 0-1.5-.673-1.5-1.5V5.414zM17 19H7a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h7v1.5C14 7.879 15.121 9 16.5 9H18v9a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentDelete(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m19.707 7.293l-4-4A.996.996 0 0 0 15 3H7C5.346 3 4 4.346 4 6v12c0 1.654 1.346 3 3 3h10c1.654 0 3-1.346 3-3V8a.996.996 0 0 0-.293-.707M17.586 8H16.5c-.827 0-1.5-.673-1.5-1.5V5.414zM17 19H7a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h7v1.5C14 7.879 15.121 9 16.5 9H18v9a1 1 0 0 1-1 1m-2-5H9a1 1 0 1 1 0-2h6a1 1 0 1 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentText(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 21H7c-1.654 0-3-1.346-3-3V6c0-1.654 1.346-3 3-3h10c1.654 0 3 1.346 3 3v12c0 1.654-1.346 3-3 3M7 5c-.551 0-1 .449-1 1v12c0 .551.449 1 1 1h10c.551 0 1-.449 1-1V6c0-.551-.449-1-1-1zm9 6H8a.5.5 0 0 1 0-1h8a.5.5 0 0 1 0 1m0-3H8a.5.5 0 0 1 0-1h8a.5.5 0 0 1 0 1m0 6H8a.5.5 0 0 1 0-1h8a.5.5 0 0 1 0 1m0 3H8a.5.5 0 0 1 0-1h8a.5.5 0 0 1 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Download(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.707 7.404c-.189-.188-.448-.283-.707-.283s-.518.095-.707.283L13 9.697V3a1 1 0 0 0-2 0v6.697L8.707 7.404a.997.997 0 0 0-1.414 0a1 1 0 0 0 0 1.414L12 13.5l4.709-4.684a1 1 0 0 0-.002-1.412M20.987 16a.98.98 0 0 0-.039-.316l-2-6A.998.998 0 0 0 18 9h-.219c-.094.188-.21.368-.367.525L15.932 11h1.348l1.667 5H5.054l1.667-5h1.348L6.586 9.525A1.96 1.96 0 0 1 6.219 9H6a.998.998 0 0 0-.948.684l-2 6a.98.98 0 0 0-.039.316C3 16 3 21 3 21a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1s0-5-.013-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownloadOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.986 17c0-.105-.004-.211-.038-.316l-2-6A.997.997 0 0 0 18 10h-.561l.682-.678a3.002 3.002 0 0 0 0-4.242c-.81-.812-2.068-1.078-3.121-.709V3c0-1.654-1.346-3-3-3S9 1.346 9 3v1.371c-1.052-.369-2.311-.103-3.121.709a3.003 3.003 0 0 0 .002 4.244l.68.676H6a.997.997 0 0 0-.948.684l-2 6a1.007 1.007 0 0 0-.038.316C3 17 3 22 3 22a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1s0-5-.014-5M7.293 6.494a1 1 0 0 1 1.414 0L11 8.787V3a1 1 0 0 1 2 0v5.787l2.293-2.293a1.025 1.025 0 0 1 1.414 0a.998.998 0 0 1 .002 1.412L12 12.59L7.293 7.908a.999.999 0 0 1 0-1.414M6.721 12h1.852l3.429 3.41L15.43 12h1.852l1.667 5H5.055zM19 21H5v-3h14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dropbox(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m3 12.9l5.3 3.5l3.7-3.1L6.7 10zm5.3-9.3L3 7.1L6.7 10L12 6.7zM21 7.1l-5.3-3.5L12 6.7l5.3 3.3zm-9 6.2l3.7 3.1l5.3-3.5l-3.7-2.9zm0 1.2l-3.7 3.1l-1.6-1.1v1.2l5.3 3.2l5.3-3.2v-1.2l-1.6 1.1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Edit(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.561 5.318l-2.879-2.879A1.495 1.495 0 0 0 17.621 2c-.385 0-.768.146-1.061.439L13 6H4a1 1 0 0 0-1 1v13a1 1 0 0 0 1 1h13a1 1 0 0 0 1-1v-9l3.561-3.561c.293-.293.439-.677.439-1.061s-.146-.767-.439-1.06M11.5 14.672L9.328 12.5l6.293-6.293l2.172 2.172zm-2.561-1.339l1.756 1.728L9 15zM16 19H5V8h6l-3.18 3.18c-.293.293-.478.812-.629 1.289c-.16.5-.191 1.056-.191 1.47V17h3.061c.414 0 1.108-.1 1.571-.29c.464-.19.896-.347 1.188-.64L16 13zm2.5-11.328L16.328 5.5l1.293-1.293l2.171 2.172z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eject(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 17.5A4.505 4.505 0 0 1 8.5 13a1 1 0 1 0-2 0c0 3.584 2.916 6.5 6.5 6.5s6.5-2.916 6.5-6.5s-2.916-6.5-6.5-6.5a1 1 0 1 0 0 2c2.481 0 4.5 2.019 4.5 4.5s-2.019 4.5-4.5 4.5M10.656 4a1 1 0 0 1 0 2H7.413l1.708 1.707l4.093 4.092a1.001 1.001 0 0 1-1.414 1.416L7.707 9.121L6 7.413v3.243a1 1 0 0 1-2 0V4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EjectOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 8.551V7a2 2 0 0 0-2-2H4v10a2 2 0 0 0 2 2h1.066c.893 2.887 3.646 5 6.809 5c3.859 0 7.062-3.016 7.062-6.875c.001-3.161-2.068-5.74-4.937-6.574m-8 1.862v3.243a1 1 0 0 1-2 0V7h6.656a1 1 0 0 1 0 2H9.413l4.801 4.799a1 1 0 0 1-.707 1.707a.996.996 0 0 1-.707-.291zm6 9.618c-2.757 0-5-2.26-5-5.016c0-.705-.004-1.371.21-1.979l2.883 2.884c.39.39.901.584 1.414.584s1.022-.194 1.414-.584a2.002 2.002 0 0 0 0-2.83l-2.567-2.567c.517-.226 1.065-.398 1.646-.398c2.757 0 5 2.182 5 4.938c0 2.757-2.243 4.968-5 4.968"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Equals(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 7H6a2 2 0 0 0 0 4h12a2 2 0 0 0 0-4m0 7H6a2 2 0 0 0 0 4h12a2 2 0 0 0 0-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EqualsOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 12H6c-1.654 0-3-1.346-3-3s1.346-3 3-3h12c1.654 0 3 1.346 3 3s-1.346 3-3 3M6 8a1.001 1.001 0 0 0 0 2h12a1.001 1.001 0 0 0 0-2zm12 11H6c-1.654 0-3-1.346-3-3s1.346-3 3-3h12c1.654 0 3 1.346 3 3s-1.346 3-3 3M6 15a1.001 1.001 0 0 0 0 2h12a1.001 1.001 0 0 0 0-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Export(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 16.5v.5c1.691-2.578 3.6-3.953 6-4v3c0 .551.511 1 1.143 1c.364 0 .675-.158.883-.391C17.959 14.58 22 10.5 22 10.5s-4.041-4.082-5.975-6.137A1.262 1.262 0 0 0 15.143 4C14.511 4 14 4.447 14 5v3c-4.66 0-6 4.871-6 8.5M5 21h14a1 1 0 0 0 1-1v-6.046c-.664.676-1.364 1.393-2 2.047V19H6V7h7V5H5a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExportOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.711 9.796c-.041-.041-4.055-4.096-5.982-6.146A2.277 2.277 0 0 0 15.143 3C13.961 3 13 3.896 13 5H5a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1v-6.045a835.426 835.426 0 0 1 2.711-2.751a.999.999 0 0 0 0-1.408m-7.432 6.145l-.136.059l-.144-.04V12h-1c-1.771.034-3.336.68-4.753 1.958c.43-2.215 1.6-4.958 4.753-4.958h1V5.042L15.143 5l.154.05c1.436 1.525 4.051 4.187 5.297 5.45c-.253.257-4.342 4.422-5.315 5.441M6 19V7h8v1c-4.66 0-6 4.871-6 8.5v.5c1.691-2.578 3.6-3.953 6-4v3c0 .551.512 1 1.143 1c.364 0 .676-.158.883-.391c.539-.565 1.242-1.291 1.976-2.043V19z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eye(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.821 12.43c-.083-.119-2.062-2.944-4.793-4.875C15.612 6.552 13.826 6 12 6c-1.825 0-3.611.552-5.03 1.555c-2.731 1.931-4.708 4.756-4.791 4.875a1 1 0 0 0 0 1.141c.083.119 2.06 2.944 4.791 4.875C8.389 19.448 10.175 20 12 20c1.826 0 3.612-.552 5.028-1.555c2.731-1.931 4.71-4.756 4.793-4.875a.996.996 0 0 0 0-1.14M12 16.5c-1.934 0-3.5-1.57-3.5-3.5c0-1.934 1.566-3.5 3.5-3.5c1.93 0 3.5 1.566 3.5 3.5c0 1.93-1.57 3.5-3.5 3.5m2-3.5c0 1.102-.898 2-2 2a2 2 0 1 1 2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 9c1.211 0 2.381.355 3.297 1.004c1.301.92 2.43 2.124 3.165 2.996c-.735.872-1.864 2.077-3.166 2.996C14.381 16.645 13.211 17 12 17s-2.382-.355-3.299-1.004C7.4 15.076 6.271 13.872 5.537 13c.734-.872 1.863-2.076 3.164-2.995C9.618 9.355 10.789 9 12 9m0-2c-1.691 0-3.242.516-4.453 1.371C4.928 10.223 3 13 3 13s1.928 2.777 4.547 4.629C8.758 18.484 10.309 19 12 19s3.242-.516 4.451-1.371C19.07 15.777 21 13 21 13s-1.93-2.777-4.549-4.629C15.242 7.516 13.691 7 12 7m0 5a1 1 0 1 0-.002 1.998A1 1 0 0 0 12 12m0 4c-1.654 0-3-1.346-3-3s1.346-3 3-3s3 1.346 3 3s-1.346 3-3 3m0-5a2 2 0 1 0 .001 4.001A2 2 0 0 0 12 11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Feather(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.68 1.017L11.5.983l-.18.033a10.15 10.15 0 0 0-.82 19.779V22a1 1 0 1 0 2 0v-1.205a10.147 10.147 0 0 0-.82-19.778M12.5 18.7V7a1 1 0 1 0-2 0v11.7a8.139 8.139 0 0 1-5.49-7.483l3.137 3.137a.498.498 0 0 0 .707 0a.5.5 0 0 0 0-.707L5.085 9.878a8.116 8.116 0 0 1 .877-2.709l2.184 2.185a.502.502 0 0 0 .708 0a.5.5 0 0 0 0-.707L6.501 6.294A8.132 8.132 0 0 1 11.5 3.019a8.14 8.14 0 0 1 4.999 3.275l-2.353 2.353a.5.5 0 0 0 .708.707l2.184-2.185c.444.832.744 1.745.877 2.709l-3.769 3.769a.5.5 0 0 0 .708.707l3.137-3.137A8.141 8.141 0 0 1 12.5 18.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Film(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 8v7h8V8zm7 6H9V9h6zm2-12h-3v2h-4V2H7C5.346 2 4 3.346 4 5v13c0 1.654 1.346 3 3 3h3v-2h4v2h3c1.654 0 3-1.346 3-3V5c0-1.654-1.346-3-3-3m1 4a1 1 0 1 0 0 2v1a1 1 0 1 0 0 2v1a1 1 0 1 0 0 2v1a1 1 0 1 0 0 2v1c0 .551-.448 1-1 1h-1v-2H8v2H7c-.552 0-1-.449-1-1v-1a1 1 0 1 0 0-2v-1a1 1 0 1 0 0-2v-1a1 1 0 1 0 0-2V8a1 1 0 1 0 0-2V5c0-.551.448-1 1-1h1v2h8V4h1c.552 0 1 .449 1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Filter(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 6H5c-1.1 0-1.4.6-.6 1.4l4.2 4.2c.8.8 1.4 2.3 1.4 3.4v5l4-2v-3.5c0-.8.6-2.1 1.4-2.9l4.2-4.2c.8-.8.5-1.4-.6-1.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flag(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.383 4.318a1 1 0 0 0-1.09.217a3.248 3.248 0 0 1-4.586 0a5.25 5.25 0 0 0-7.414 0A.997.997 0 0 0 5 5.242v13a1 1 0 1 0 2 0v-4.553a3.248 3.248 0 0 1 4.293.26a5.25 5.25 0 0 0 7.414 0a1 1 0 0 0 .293-.707v-8a1 1 0 0 0-.617-.924"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlagOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.383 4.318a1 1 0 0 0-1.09.217a3.248 3.248 0 0 1-4.586 0a5.25 5.25 0 0 0-7.414 0A.997.997 0 0 0 5 5.242v13a1 1 0 0 0 2 0v-4.553a3.248 3.248 0 0 1 4.293.26a5.25 5.25 0 0 0 7.414 0a1 1 0 0 0 .293-.707v-8a1 1 0 0 0-.617-.924m-7.09 1.631A5.254 5.254 0 0 0 17 7.087v2.311a3.746 3.746 0 0 1-4.646-.51C10.906 7.441 8.756 7.145 7 7.961V5.689a3.247 3.247 0 0 1 4.293.26m1.414 6.585A5.23 5.23 0 0 0 9 11.002c-.681 0-1.361.131-2 .394V9.085a3.746 3.746 0 0 1 4.646.51A4.73 4.73 0 0 0 15 10.981c.687 0 1.366-.164 2-.459v2.272a3.248 3.248 0 0 1-4.293-.26"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flash(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m17.502 12.033l-4.241-2.458l2.138-5.131A1.003 1.003 0 0 0 14.505 3a1.004 1.004 0 0 0-.622.214l-.07.06l-7.5 7.1a1.002 1.002 0 0 0 .185 1.592l4.242 2.46l-2.163 5.19a.999.999 0 0 0 1.611 1.11l7.5-7.102a1.002 1.002 0 0 0-.186-1.591"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlashOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 4h.005M14.5 4L12 10l5 2.898L9.5 20l2.5-6l-5-2.9zm0-2a2.024 2.024 0 0 0-1.379.551L5.624 9.646a1.998 1.998 0 0 0-.61 1.686c.072.626.437 1.182.982 1.498l3.482 2.021l-1.826 4.381a2.003 2.003 0 0 0 1.847 2.77c.498 0 .993-.186 1.375-.548l7.5-7.103a1.995 1.995 0 0 0 .61-1.685a1.999 1.999 0 0 0-.982-1.498L14.52 9.15l1.789-4.293A2 2 0 0 0 14.5 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlowChildren(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 16a2.99 2.99 0 0 0-2.816 2H11c-1.654 0-3-1.346-3-3v-3.025A4.954 4.954 0 0 0 11 13h3.184A2.99 2.99 0 0 0 17 15a3 3 0 1 0 0-6a2.99 2.99 0 0 0-2.816 2H11c-1.654 0-3-1.346-3-3v-.184A2.99 2.99 0 0 0 10 5a3 3 0 1 0-6 0a2.99 2.99 0 0 0 2 2.816V15c0 2.757 2.243 5 5 5h3.184A2.99 2.99 0 0 0 17 22a3 3 0 1 0 0-6m0-5a1.001 1.001 0 1 1-1 1c0-.551.448-1 1-1M7 4a1.001 1.001 0 1 1-1 1c0-.551.448-1 1-1m10 16a1.001 1.001 0 0 1 0-2a1.001 1.001 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlowMerge(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 16.184v-1.851c0-1.93-1.57-3.5-3.5-3.5c-.827 0-1.5-.673-1.5-1.5V7.816A2.997 2.997 0 0 0 15 5c0-1.654-1.346-3-3-3S9 3.346 9 5c0 1.302.839 2.401 2 2.815v1.518c0 .827-.673 1.5-1.5 1.5c-1.93 0-3.5 1.57-3.5 3.5v1.851A2.997 2.997 0 0 0 4 19c0 1.654 1.346 3 3 3s3-1.346 3-3a2.997 2.997 0 0 0-2-2.816v-1.851c0-.827.673-1.5 1.5-1.5c.979 0 1.864-.407 2.5-1.058a3.487 3.487 0 0 0 2.5 1.058c.827 0 1.5.673 1.5 1.5v1.851A2.997 2.997 0 0 0 14 19c0 1.654 1.346 3 3 3s3-1.346 3-3a2.997 2.997 0 0 0-2-2.816M7 20a1.001 1.001 0 0 1 0-2a1.001 1.001 0 0 1 0 2m5-16a1.001 1.001 0 1 1-1 1c0-.551.448-1 1-1m5 16a1.001 1.001 0 0 1 0-2a1.001 1.001 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlowParallel(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 16.184V7.816A2.997 2.997 0 0 0 20 5c0-1.654-1.346-3-3-3s-3 1.346-3 3c0 1.302.839 2.401 2 2.815v8.369A2.997 2.997 0 0 0 14 19c0 1.654 1.346 3 3 3s3-1.346 3-3a2.997 2.997 0 0 0-2-2.816M17 4a1.001 1.001 0 1 1-1 1c0-.551.448-1 1-1m0 16a1.001 1.001 0 0 1 0-2a1.001 1.001 0 0 1 0 2M10 5c0-1.654-1.346-3-3-3S4 3.346 4 5c0 1.302.839 2.401 2 2.815v8.369A2.997 2.997 0 0 0 4 19c0 1.654 1.346 3 3 3s3-1.346 3-3a2.997 2.997 0 0 0-2-2.816V7.816A2.997 2.997 0 0 0 10 5M7 4a1.001 1.001 0 1 1-1 1c0-.551.448-1 1-1m0 16a1.001 1.001 0 0 1 0-2a1.001 1.001 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlowSwitch(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 16.184V15.5c0-.848.512-1.595 1.287-2.047a7.008 7.008 0 0 1-1.822-1.131C6.561 13.136 6 14.26 6 15.5v.684A2.997 2.997 0 0 0 4 19c0 1.654 1.346 3 3 3s3-1.346 3-3a2.997 2.997 0 0 0-2-2.816M7 20a1.001 1.001 0 0 1 0-2a1.001 1.001 0 0 1 0 2m9-12.185v.351c0 .985-.535 1.852-1.345 2.36a7.016 7.016 0 0 1 1.823 1.1C17.414 10.748 18 9.524 18 8.167v-.351A2.997 2.997 0 0 0 20 5c0-1.654-1.346-3-3-3s-3 1.346-3 3c0 1.302.839 2.401 2 2.815M17 4a1.001 1.001 0 1 1-1 1c0-.551.448-1 1-1m.935 12.164C17.525 13.251 15.024 11 12 11a4.004 4.004 0 0 1-3.92-3.209A3 3 0 0 0 10 5c0-1.654-1.346-3-3-3S4 3.346 4 5c0 1.326.87 2.44 2.065 2.836C6.475 10.749 8.976 13 12 13a4.004 4.004 0 0 1 3.92 3.209A3 3 0 0 0 14 19c0 1.654 1.346 3 3 3s3-1.346 3-3c0-1.326-.87-2.44-2.065-2.836M7 4a1.001 1.001 0 1 1-1 1c0-.551.448-1 1-1m10 16a1.001 1.001 0 0 1 0-2a1.001 1.001 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folder(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 6h-6a2 2 0 0 0-2-2H6C4.346 4 3 5.346 3 7v10c0 1.654 1.346 3 3 3h12c1.654 0 3-1.346 3-3V9c0-1.654-1.346-3-3-3M6 6h4a2 2 0 0 0 2 2h6a1 1 0 0 1 1 1H5V7a1 1 0 0 1 1-1m12 12H6a1 1 0 0 1-1-1v-7h14v7a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderAdd(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 6h-6a2 2 0 0 0-2-2H6C4.346 4 3 5.346 3 7v10c0 1.654 1.346 3 3 3h12c1.654 0 3-1.346 3-3V9c0-1.654-1.346-3-3-3m0 12H6a1 1 0 0 1-1-1v-7h4c.275 0 .5-.225.5-.5S9.275 9 9 9H5V7a1 1 0 0 1 1-1h4a2 2 0 0 0 2 2h6a1 1 0 0 1 1 1h-4c-.275 0-.5.225-.5.5s.225.5.5.5h4v7a1 1 0 0 1-1 1m-3-6h-2v-2a1 1 0 1 0-2 0v2H9a1 1 0 1 0 0 2h2v2a1 1 0 1 0 2 0v-2h2a1 1 0 1 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderDelete(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 6h-6a2 2 0 0 0-2-2H6C4.346 4 3 5.346 3 7v10c0 1.654 1.346 3 3 3h12c1.654 0 3-1.346 3-3V9c0-1.654-1.346-3-3-3M6 6h4a2 2 0 0 0 2 2h6a1 1 0 0 1 1 1H5V7a1 1 0 0 1 1-1m12 12H6a1 1 0 0 1-1-1v-7h14v7a1 1 0 0 1-1 1m-3-4H9a1 1 0 1 1 0-2h6a1 1 0 1 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOpen(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.3 8h-2.4c-.4-1.2-1.5-2-2.8-2h-6c0-1.1-.9-2-2-2H5C3.3 4 2 5.3 2 7v10c0 1.7 1.3 3 3 3h12c1.7 0 3.4-1.3 3.8-3L23 9c.1-.6-.2-1-.7-1M4 9V7c0-.6.4-1 1-1h4c0 1.1.9 2 2 2h6c.6 0 1 .4 1 1H6.9c-.6 0-1.1.4-1.3 1L4 16.3zm14.9 7.5c-.2.8-1.1 1.5-1.9 1.5H5s-.4-.2-.2-.8l1.9-7c0-.1.2-.2.3-.2h13.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gift(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 8h-2.352c.219-.457.352-.961.352-1.5C17 4.57 15.43 3 13.5 3c-.979 0-1.864.407-2.5 1.058A3.487 3.487 0 0 0 8.5 3C6.57 3 5 4.57 5 6.5c0 .539.133 1.043.352 1.5H3a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1v5c0 1.654 1.346 3 3 3h10c1.654 0 3-1.346 3-3v-5a1 1 0 0 0 1-1V9a1 1 0 0 0-1-1m-1 4h-5v-2h5zm-8-5h2v1h-2zm2 3v2h-2v-2zm1.5-5c.827 0 1.5.673 1.5 1.5S14.327 8 13.5 8a1.43 1.43 0 0 1-.5-.097V7a.99.99 0 0 0-.913-.982A1.497 1.497 0 0 1 13.5 5M7 6.5C7 5.673 7.673 5 8.5 5c.657 0 1.211.428 1.413 1.018A.99.99 0 0 0 9 7v.903A1.43 1.43 0 0 1 8.5 8C7.673 8 7 7.327 7 6.5M9 10v2H4v-2zM6 20c-.551 0-1-.449-1-1v-6h4v7zm4 0v-7h2v7zm6 0h-3v-7h4v6c0 .551-.449 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Globe(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 20H7a1 1 0 1 0 0 2h10a1 1 0 1 0 0-2h-4v-1.23a8.925 8.925 0 0 0 4.363-2.406A8.942 8.942 0 0 0 20 10a8.93 8.93 0 0 0-1.968-5.619l.675-.673a1.001 1.001 0 0 0-1.414-1.416l-2.052 2.049l.708.708C17.271 6.371 18 8.13 18 10s-.729 3.627-2.051 4.949S12.87 17 11 17s-3.627-.729-4.949-2.051a.999.999 0 1 0-1.414 1.414A8.938 8.938 0 0 0 11 19zm0-16c1.657 0 3.157.672 4.243 1.757A5.985 5.985 0 0 1 17 10a5.985 5.985 0 0 1-1.757 4.242A5.982 5.982 0 0 1 11 16a5.98 5.98 0 0 1-4.242-1.757C5.673 13.157 5.002 11.657 5.002 10s.671-3.157 1.756-4.243A5.98 5.98 0 0 1 11 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlobeOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 6c2.206 0 4 1.794 4 4s-1.794 4-4 4c-2.204 0-3.998-1.794-3.998-4S8.796 6 11 6m0-2a5.998 5.998 0 0 0-5.998 6A5.998 5.998 0 1 0 17 10a6 6 0 0 0-6-6m6 16h-4v-1.23a8.92 8.92 0 0 0 4.363-2.406A8.944 8.944 0 0 0 20 10.001a8.927 8.927 0 0 0-1.968-5.619l.675-.673a1.001 1.001 0 0 0-1.414-1.416l-2.052 2.049l.708.708C17.271 6.371 18 8.13 18 10s-.729 3.627-2.051 4.949S12.87 17 11 17s-3.627-.729-4.949-2.051a.999.999 0 1 0-1.414 1.414A8.942 8.942 0 0 0 11 19v1H7a1 1 0 1 0 0 2h10a1 1 0 1 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GroupIcon(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 14c1.381 0 2.631-.56 3.536-1.465C16.44 11.631 17 10.381 17 9s-.56-2.631-1.464-3.535C14.631 4.56 13.381 4 12 4s-2.631.56-3.536 1.465C7.56 6.369 7 7.619 7 9s.56 2.631 1.464 3.535A4.985 4.985 0 0 0 12 14m8 1a2.495 2.495 0 0 0 2.5-2.5c0-.69-.279-1.315-.732-1.768A2.492 2.492 0 0 0 20 10a2.495 2.495 0 0 0-2.5 2.5A2.496 2.496 0 0 0 20 15m0 .59c-1.331 0-2.332.406-2.917.968C15.968 15.641 14.205 15 12 15c-2.266 0-3.995.648-5.092 1.564C6.312 15.999 5.3 15.59 4 15.59c-2.188 0-3.5 1.09-3.5 2.182c0 .545 1.312 1.092 3.5 1.092c.604 0 1.146-.051 1.623-.133l-.04.27c0 1 2.406 2 6.417 2c3.762 0 6.417-1 6.417-2l-.02-.255c.463.073.995.118 1.603.118c2.051 0 3.5-.547 3.5-1.092c0-1.092-1.373-2.182-3.5-2.182M4 15c.69 0 1.315-.279 1.768-.732A2.492 2.492 0 0 0 6.5 12.5A2.495 2.495 0 0 0 4 10a2.496 2.496 0 0 0-2.5 2.5A2.495 2.495 0 0 0 4 15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GroupOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 14c2.764 0 5-2.238 5-5s-2.236-5-5-5s-5 2.238-5 5s2.236 5 5 5m0-8c1.654 0 3 1.346 3 3s-1.346 3-3 3s-3-1.346-3-3s1.346-3 3-3m8 9a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5m0-4c.827 0 1.5.673 1.5 1.5S20.827 14 20 14s-1.5-.673-1.5-1.5s.673-1.5 1.5-1.5m0 4.59c-1.331 0-2.332.406-2.917.969C15.968 15.641 14.205 15 12 15c-2.266 0-3.995.648-5.092 1.564c-.596-.565-1.608-.975-2.908-.975c-2.188 0-3.5 1.091-3.5 2.183c0 .545 1.312 1.092 3.5 1.092c.604 0 1.146-.051 1.623-.133l-.04.27c0 1 2.405 2 6.417 2c3.762 0 6.417-1 6.417-2l-.021-.255c.463.073.996.118 1.604.118c2.051 0 3.5-.547 3.5-1.092c0-1.092-1.373-2.182-3.5-2.182M4 17.863c-1.309 0-2.068-.207-2.417-.354c.239-.405 1.003-.92 2.417-.92c1.107 0 1.837.351 2.208.706l-.235.344c-.452.119-1.108.224-1.973.224M12 19c-2.163 0-3.501-.312-4.184-.561C8.337 17.761 9.734 17 12 17c2.169 0 3.59.761 4.148 1.425c-.755.27-2.162.575-4.148.575m8-1.137c-.914 0-1.546-.103-1.973-.213a3.42 3.42 0 0 0-.248-.375c.356-.345 1.071-.685 2.221-.685c1.324 0 2.141.501 2.404.911c-.39.163-1.205.362-2.404.362M4 15a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5m0-4c.827 0 1.5.673 1.5 1.5S4.827 14 4 14s-1.5-.673-1.5-1.5S3.173 11 4 11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Headphones(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 13c0-4.963-4.037-9-9-9s-9 4.037-9 9v2.6l.023.113c-.013.243-.023.5-.023.787v2C3 20.43 4.57 22 6.5 22s3.5-1.57 3.5-3.5v-2c0-1.511-.83-2.79-1.982-3.278C8.113 11.1 9.855 9.399 12 9.399s3.887 1.7 3.982 3.822C14.83 13.711 14 14.989 14 16.5v2c0 1.93 1.57 3.5 3.5 3.5s3.5-1.57 3.5-3.5v-2c0-.287-.01-.544-.023-.787L21 15.6zM5 13c0-1.586.538-3.046 1.432-4.221l.89.889A5.959 5.959 0 0 0 6.02 13C6 13 6 14 6 14H5zm3 5.5c0 .827-.673 1.5-1.5 1.5S5 19.327 5 18.5v-2c0-.666.057-1.176.114-1.5H7c.473 0 1 .616 1 1.5zm7.77-9.338l-.354.354C14.504 8.603 13.291 8.1 12 8.1s-2.504.503-3.417 1.416l-.354-.354l-1.141-1.141l-.627-.626A7.78 7.78 0 0 1 12 5.1c2.093 0 4.06.815 5.539 2.295l-.627.626zM19 18.5c0 .827-.673 1.5-1.5 1.5s-1.5-.673-1.5-1.5v-2c0-.884.527-1.5 1-1.5h1.886c.057.324.114.834.114 1.5zm0-4.5h-1v-1a6.129 6.129 0 0 0-1.322-3.332l.891-.889A6.952 6.952 0 0 1 19 13z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heart(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 10.375a4.375 4.375 0 0 0-8.75 0c0 1.127.159 2.784 1.75 4.375L12 20s5.409-3.659 7-5.25s1.75-3.248 1.75-4.375a4.375 4.375 0 0 0-8.75 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartFullOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.2 9.4c0 1.3.2 3.3 2 5.1c1.6 1.6 6.9 5.2 7.1 5.4c.2.1.4.2.6.2s.4-.1.6-.2c.2-.2 5.5-3.7 7.1-5.4c1.8-1.8 2-3.8 2-5.1c0-3-2.4-5.4-5.4-5.4c-1.6 0-3.2.9-4.2 2.3C11 4.9 9.4 4 7.6 4C4.7 4 2.2 6.4 2.2 9.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartHalfOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.2 9.4c0 1.3.2 3.3 2 5.1c1.6 1.6 6.9 5.2 7.1 5.4c.2.1.4.2.6.2s.4-.1.6-.2c.2-.2 5.5-3.7 7.1-5.4c1.8-1.8 2-3.8 2-5.1c0-3-2.4-5.4-5.4-5.4c-1.6 0-3.2.9-4.2 2.3C11 4.9 9.4 4 7.6 4C4.7 4 2.2 6.4 2.2 9.4m9.8 1c.6 0 1-.4 1-1C13 7.5 14.5 6 16.4 6s3.4 1.5 3.4 3.4c0 1.1-.2 2.4-1.5 3.7c-1.2 1.2-4.9 3.8-6.3 4.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 20a1 1 0 0 1-.561-.172c-.225-.151-5.508-3.73-7.146-5.371C2.462 12.626 2.25 10.68 2.25 9.375A5.38 5.38 0 0 1 7.625 4c1.802 0 3.398.891 4.375 2.256A5.373 5.373 0 0 1 16.375 4a5.38 5.38 0 0 1 5.375 5.375c0 1.305-.212 3.251-2.043 5.082c-1.641 1.641-6.923 5.22-7.146 5.371A1 1 0 0 1 12 20M7.625 6A3.379 3.379 0 0 0 4.25 9.375c0 1.093.173 2.384 1.457 3.668c1.212 1.212 4.883 3.775 6.293 4.746c1.41-.971 5.081-3.534 6.293-4.746c1.284-1.284 1.457-2.575 1.457-3.668C19.75 7.514 18.236 6 16.375 6S13 7.514 13 9.375a1 1 0 1 1-2 0A3.379 3.379 0 0 0 7.625 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Home(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 3s-6.186 5.34-9.643 8.232A1.041 1.041 0 0 0 2 12a1 1 0 0 0 1 1h2v7a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1v-4h4v4a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1v-7h2a1 1 0 0 0 1-1a.98.98 0 0 0-.383-.768C18.184 8.34 12 3 12 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.262 10.468c-3.39-2.854-9.546-8.171-9.607-8.225L12 1.68l-.652.563c-.062.053-6.221 5.368-9.66 8.248A2.042 2.042 0 0 0 1 12a2 2 0 0 0 2 2h1v6a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-6h1a2 2 0 0 0 2-2c0-.598-.275-1.161-.738-1.532M14 20h-4v-5h4zm4-8l.002 8H15v-6H9v6H6v-8H2.999C5.764 9.688 10.314 5.773 12 4.32c1.686 1.453 6.234 5.367 9 7.681z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HtmlFive(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.1 3.5l.7 1.1l.7-1.1V5h1V2h-1l-.7 1.1l-.6-1.1h-1.1v3h1zM18.4 5V4H17V2h-1v3zM9.8 5h1V3h.9V2H8.9v1h.9zM6.6 4h.9v1h1V2h-1v1h-.9V2h-1v3h1zM5 6l1.2 14.4L12 22l5.8-1.6L19 6zm11.3 4.6H9.5l.2 1.8h6.4l-.5 5.5l-3.6 1l-3.6-1l-.3-2.9h1.8l.1 1.5l2 .5l2-.5l.2-2.3H8l-.5-5.3h9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Image(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="8.5" cy="8.5" r="2.5" fill="currentColor"/><path fill="currentColor" d="M16 10c-2 0-3 3-4.5 3s-1.499-1-3.5-1c-2 0-3.001 4-3.001 4H19s-1-6-3-6m4-7H4c-1.103 0-2 .897-2 2v12c0 1.103.897 2 2 2h16c1.103 0 2-.897 2-2V5c0-1.103-.897-2-2-2m0 14H4V5h16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 7.999c.827 0 1.5.673 1.5 1.5s-.673 1.5-1.5 1.5s-1.5-.673-1.5-1.5s.673-1.5 1.5-1.5m0-1a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5m7.5 5c.45.051 1.27 1.804 1.779 4.001H6.387c.434-1.034 1.055-2.001 1.612-2.001c.806 0 1.125.185 1.53.42c.447.258 1.006.58 1.97.58c1.138 0 1.942-.885 2.653-1.666c.627-.687 1.218-1.334 1.848-1.334m0-1c-2 0-3 3-4.5 3s-1.499-1-3.5-1c-2 0-3.001 4-3.001 4H19s-1-6-3-6M22 6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2zm-2 12H4V6h16.003z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Infinity(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.433 8.596a4.282 4.282 0 0 0-3.036 1.246l-1.396 1.34l-1.375-1.32a4.287 4.287 0 0 0-3.055-1.266a4.289 4.289 0 0 0-3.053 1.266a4.288 4.288 0 0 0-1.267 3.055c0 1.152.449 2.238 1.266 3.053a4.285 4.285 0 0 0 3.054 1.266a4.272 4.272 0 0 0 3.036-1.248l1.395-1.338l1.376 1.32c.815.816 1.901 1.266 3.055 1.266s2.238-.449 3.053-1.266c.817-.814 1.267-1.9 1.267-3.055s-.449-2.238-1.266-3.055a4.296 4.296 0 0 0-3.054-1.264m-7.576 5.605c-.687.688-1.884.688-2.572 0a1.807 1.807 0 0 1-.533-1.285c0-.486.189-.941.535-1.287c.342-.344.799-.533 1.284-.533s.942.189 1.305.551l1.321 1.27zm8.861 0c-.687.689-1.866.705-2.59-.018l-1.321-1.27l1.339-1.285c.688-.688 1.886-.688 2.573-.002c.344.346.533.801.533 1.287s-.19.944-.534 1.288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfinityOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.434 8.596c1.152 0 2.237.449 3.053 1.264c.815.816 1.266 1.9 1.266 3.056s-.45 2.239-1.268 3.056c-.813.815-1.898 1.266-3.053 1.266s-2.238-.449-3.055-1.266l-1.376-1.32l-1.395 1.338a4.27 4.27 0 0 1-3.036 1.248a4.281 4.281 0 0 1-3.054-1.267c-.815-.813-1.265-1.899-1.265-3.053s.45-2.237 1.267-3.056a4.295 4.295 0 0 1 3.053-1.266c1.154 0 2.239.449 3.055 1.266l1.375 1.32l1.396-1.34a4.279 4.279 0 0 1 3.037-1.246m0-2c-1.679 0-3.25.645-4.433 1.813a6.252 6.252 0 0 0-4.43-1.813a6.266 6.266 0 0 0-4.467 1.853c-1.194 1.192-1.852 2.78-1.852 4.469s.658 3.274 1.852 4.468a6.278 6.278 0 0 0 4.468 1.852a6.25 6.25 0 0 0 4.431-1.814a6.249 6.249 0 0 0 4.431 1.814a6.27 6.27 0 0 0 4.469-1.854a6.263 6.263 0 0 0 1.852-4.467a6.285 6.285 0 0 0-1.852-4.47a6.283 6.283 0 0 0-4.469-1.851m-8.863 5.5c.225 0 .426.088.612.271l.57.548l-.603.579a.814.814 0 0 1-.578.223a.815.815 0 0 1-.58-.223a.812.812 0 0 1-.24-.578c0-.221.084-.422.243-.581a.803.803 0 0 1 .576-.239m0-1c-.486 0-.942.189-1.285.533a1.809 1.809 0 0 0-.535 1.287c0 .484.189.941.533 1.285c.344.344.815.516 1.287.516a1.81 1.81 0 0 0 1.285-.516l1.339-1.285l-1.321-1.27a1.819 1.819 0 0 0-1.303-.55m8.863 1.017a.81.81 0 0 1 .576.219c.158.159.242.359.242.582s-.083.422-.243.581a.802.802 0 0 1-.571.228a.858.858 0 0 1-.617-.261l-.571-.548l.603-.578a.814.814 0 0 1 .581-.223m0-1c-.472 0-.943.172-1.287.516l-1.34 1.285l1.322 1.27a1.85 1.85 0 0 0 1.311.539a1.798 1.798 0 0 0 1.813-1.808a1.81 1.81 0 0 0-.532-1.287a1.821 1.821 0 0 0-1.287-.515"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Info(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.17 15.4l-5.91-9.85c-.78-1.3-1.96-2.04-3.26-2.04s-2.48.74-3.26 2.03L2.83 15.4c-.44.73-.66 1.49-.66 2.21c0 .57.14 1.13.42 1.62C3.23 20.35 4.47 21 6 21h12c1.53 0 2.77-.65 3.41-1.77c.28-.49.42-1.02.42-1.58c.01-.74-.21-1.51-.66-2.25M12 8.45c.85 0 1.55.7 1.55 1.55c0 .85-.69 1.55-1.55 1.55c-.85 0-1.55-.7-1.55-1.55c0-.86.69-1.55 1.55-1.55m1.69 8.46c-.03.04-.8.92-2.07.92h-.15c-.51-.03-.93-.25-1.18-.63c-.31-.47-.36-1.11-.12-1.82l.41-1.22c.23-.68.01-.79-.11-.85l-.14-.02c-.25 0-.6.15-.71.21c-.1.05-.23.03-.31-.07c-.07-.1-.07-.23.01-.32c.03-.04.87-.99 2.22-.91c.51.03.93.25 1.18.63c.32.47.36 1.11.12 1.83l-.41 1.22c-.23.68-.01.79.11.85l.14.02c.25 0 .6-.15.71-.2c.11-.06.23-.03.31.07c.07.07.07.2-.01.29"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoLarge(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.839 17.525c-.006.002-.559.186-1.039.186c-.265 0-.372-.055-.406-.079c-.168-.117-.48-.336.054-1.4l1-1.994c.593-1.184.681-2.329.245-3.225c-.356-.733-1.039-1.236-1.92-1.416a4.776 4.776 0 0 0-.958-.097c-1.849 0-3.094 1.08-3.146 1.126a.5.5 0 0 0 .493.848c.005-.002.559-.187 1.039-.187c.263 0 .369.055.402.078c.169.118.482.34-.051 1.402l-1 1.995c-.594 1.185-.681 2.33-.245 3.225c.356.733 1.038 1.236 1.921 1.416c.314.063.636.097.954.097c1.85 0 3.096-1.08 3.148-1.126a.5.5 0 0 0-.491-.849"/><circle cx="13" cy="6.001" r="2.5" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoLargeOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14.234 16.014l.554-1.104c.808-1.61.897-3.228.253-4.552a3.93 3.93 0 0 0-.443-.693A4.005 4.005 0 0 0 17 6.001c0-2.206-1.794-4-4-4s-4 1.794-4 4c0 .783.234 1.508.624 2.126c-1.696.33-2.806 1.248-2.947 1.375a1.997 1.997 0 0 0 1.089 3.484l-.554 1.104c-.808 1.61-.897 3.228-.254 4.552c.565 1.164 1.621 1.955 2.972 2.229c.413.084.836.127 1.254.127c2.368 0 3.965-1.347 4.14-1.501a1.996 1.996 0 0 0-1.09-3.483M13 4.001a2 2 0 1 1-.001 4.001A2 2 0 0 1 13 4.001M11.184 19c-.271 0-.559-.025-.854-.087c-1.642-.334-2.328-1.933-1.328-3.927l1-1.995c.5-.996.47-1.63-.108-2.035c-.181-.125-.431-.169-.689-.169c-.577 0-1.201.214-1.201.214S9.137 10 10.816 10c.271 0 .56.025.856.087c1.64.334 2.328 1.933 1.328 3.927l-1 1.993c-.5.998-.472 1.632.106 2.035c.181.126.433.169.692.169c.577 0 1.2-.212 1.2-.212S12.865 19 11.184 19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 5.51c.56 0 1.12.35 1.54 1.06l5.91 9.85c.86 1.42.2 2.58-1.45 2.58H6c-1.65 0-2.31-1.16-1.46-2.57l5.91-9.85c.43-.72.99-1.07 1.55-1.07m0-2c-1.3 0-2.48.74-3.26 2.03L2.83 15.4c-.44.74-.66 1.5-.66 2.23c0 .56.14 1.11.42 1.6C3.23 20.35 4.47 21 6 21h12c1.53 0 2.77-.65 3.41-1.77c.29-.51.43-1.07.42-1.65c-.01-.71-.23-1.46-.66-2.18l-5.91-9.85c-.78-1.3-1.96-2.04-3.26-2.04m1.5 13.24s-.71.36-1.07.18c-.36-.18-.43-.54-.23-1.15l.41-1.22c.4-1.22-.12-2.08-1.08-2.13c-1.26-.07-2.02.83-2.02.83s.71-.36 1.07-.18c.36.18.43.54.23 1.15l-.41 1.22c-.4 1.22.12 2.07 1.08 2.13c1.26.06 2.02-.83 2.02-.83"/><circle cx="12" cy="10" r="1.3" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InputChecked(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 19H8c-1.654 0-3-1.346-3-3V8c0-1.654 1.346-3 3-3h5a1 1 0 1 1 0 2H8c-.552 0-1 .449-1 1v8c0 .551.448 1 1 1h8c.552 0 1-.449 1-1v-3a1 1 0 1 1 2 0v3c0 1.654-1.346 3-3 3m-2.834-4.167c-.35 0-.689-.139-.941-.391l-2.668-2.667a1.334 1.334 0 0 1 1.887-1.885l1.416 1.417l3.475-5.455a1.334 1.334 0 1 1 2.332 1.294l-4.334 7a1.332 1.332 0 0 1-.98.673z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InputCheckedOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.885 6.177A3 3 0 0 0 17.001 4c-.862 0-1.661.381-2.219 1H8C6.346 5 5 6.346 5 8v8c0 1.654 1.346 3 3 3h8c1.654 0 3-1.346 3-3V9.546l.622-1.088c.39-.7.482-1.51.263-2.281m-3.758.338a.994.994 0 0 1 .459-.416c.301-.113.623-.127.9.027c.232.13.402.343.476.6a1.007 1.007 0 0 1-.062.685l-.021.065l-4.006 7.011a1.01 1.01 0 0 1-.742.506l-.132.009a.993.993 0 0 1-.707-.293l-3-3a1 1 0 0 1 1.414-1.414l1.125 1.125l.92.92l.652-1.125zM16 17H8c-.552 0-1-.449-1-1V8c0-.551.448-1 1-1h6.689l-2.15 3.712l-1.125-1.125c-.391-.391-.902-.586-1.414-.586s-1.023.195-1.414.586a2 2 0 0 0 0 2.828l3 3a2 2 0 0 0 1.414.586l.277-.02a2.001 2.001 0 0 0 1.471-1.01L17 11.031V16c0 .551-.448 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Key(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 11c0 .732.166 1.424.449 2.051L5 17v1.5S5.896 20 7 20h2v-2h2v-2h2.5c2.762 0 5-2.238 5-5s-2.238-5-5-5s-5 2.238-5 5m5 2a2 2 0 1 1 .001-4.001A2 2 0 0 1 13.5 13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 21H4v-4.414l3.783-3.783A5.927 5.927 0 0 1 7.5 11c0-3.309 2.691-6 6-6s6 2.691 6 6s-2.691 6-6 6H12v2h-2zm-4-2h2v-2h2v-2h3.5c2.206 0 4-1.794 4-4s-1.794-4-4-4s-4 1.794-4 4c0 .559.121 1.109.359 1.639l.285.631L6 17.414zm7.5-9.002a1.002 1.002 0 0 1 0 2.002a1.001 1.001 0 0 1 0-2.002m0-1A2.001 2.001 0 1 0 13.502 13a2.001 2.001 0 0 0-.002-4.002"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Keyboard(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 13h7v2H8zm-3 0h2v2H5zm0-4h2v1H5zm3 3v-1H5v1h2zm0-3h1v1H8zm1 2h1v1H9zm1-2h1v1h-1zm1 2h1v1h-1zm1-2h1v1h-1zm1 2h1v1h-1zm1-2h1v1h-1zm1 2h1v1h-1zm1-2h1v1h-1zm1 3h2V9h-1v2h-1zm1 1h-1v1h-1v1h3v-1h-1zm2-7H4c-1.1 0-2 .9-2 2v8c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V8c0-1.1-.9-2-2-2m0 10H4V8h16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Leaf(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 11c0-4.9-3.499-9.1-8.32-9.983L11.5.983l-.18.033a10.15 10.15 0 0 0-.82 19.779V22a1 1 0 1 0 2 0v-1.205A10.147 10.147 0 0 0 20 11m-7.5 7.7v-2.993l4.354-4.354a.5.5 0 0 0-.707-.707L12.5 14.293v-3.586l2.354-2.354a.5.5 0 0 0-.707-.707L12.5 9.293V6a1 1 0 1 0-2 0v3.293L8.854 7.646a.5.5 0 0 0-.707.707l2.354 2.354v3.586l-3.646-3.646a.5.5 0 0 0-.707.707l4.354 4.354V18.7A8.144 8.144 0 0 1 11.5 3.019a8.145 8.145 0 0 1 1 15.681"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lightbulb(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 5.5a.5.5 0 0 0 0 1c1.083 0 1.964.881 1.964 1.964a.5.5 0 0 0 1 0A2.968 2.968 0 0 0 12.5 5.5m0-4.5C8.364 1 5 4.364 5 8.5c0 1.486.44 2.922 1.274 4.165l.08.135C8.179 15.406 8.5 16.23 8.5 17v3a1 1 0 0 0 1 1h2c0 .26.11.52.29.71c.19.18.45.29.71.29c.26 0 .52-.11.71-.29c.18-.19.29-.45.29-.71h2a1 1 0 0 0 1-1v-3c0-.782.319-1.61 2.132-4.199A7.453 7.453 0 0 0 20 8.5C20 4.364 16.636 1 12.5 1m2 18h-4v-1h4zm2.495-7.347c-1.466 2.093-2.143 3.289-2.385 4.347H13.5v-2a1 1 0 0 0-2 0v2h-1.113c-.24-1.03-.898-2.2-2.306-4.22l-.077-.129A5.454 5.454 0 0 1 7 8.5C7 5.467 9.467 3 12.5 3S18 5.467 18 8.5a5.463 5.463 0 0 1-1.005 3.153"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Link(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.277 6.321a1.1 1.1 0 0 0-1.556 0L15 8.043l-.308-.308c-1.168-1.168-3.216-1.168-4.384 0l-4.172 4.172c-.584.584-.906 1.363-.906 2.192s.322 1.608.906 2.192l.308.308l-1.722 1.722a1.1 1.1 0 1 0 1.556 1.556L8 18.155l.308.308c.584.584 1.362.906 2.192.906s1.608-.322 2.192-.906l4.172-4.172c.584-.584.906-1.362.906-2.192s-.322-1.608-.906-2.192l-.308-.308l1.722-1.722a1.1 1.1 0 0 0-.001-1.556m-2.969 6.414l-4.172 4.172c-.168.168-.402.253-.636.253s-.468-.084-.636-.253l-.308-.308l.722-.722a1.1 1.1 0 1 0-1.556-1.556L8 15.043l-.308-.308c-.168-.168-.261-.395-.261-.636s.093-.468.261-.636l4.172-4.172c.168-.168.394-.261.636-.261s.468.093.636.261l.308.308l-.722.722a1.1 1.1 0 1 0 1.556 1.556l.722-.722l.308.308c.168.168.261.395.261.636s-.093.468-.261.636"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.5 5.999a1.1 1.1 0 0 1 .777 1.878L16.557 9.6l.307.308c.585.584.906 1.362.906 2.192s-.321 1.607-.906 2.191l-4.172 4.172c-.584.584-1.361.906-2.191.906s-1.607-.322-2.191-.906L8 18.154l-1.723 1.723c-.215.215-.495.322-.777.322s-.562-.107-.777-.322a1.1 1.1 0 0 1 0-1.557l1.72-1.72l-.308-.309c-.583-.584-.905-1.361-.905-2.191s.321-1.608.905-2.192l4.173-4.173c.584-.584 1.387-.875 2.191-.875s1.607.291 2.191.875l.31.308l1.723-1.723c.215-.215.495-.321.777-.321m0-2c-.828 0-1.605.321-2.191.908l-.492.491a5.223 5.223 0 0 0-2.316-.539c-1.363 0-2.677.533-3.605 1.461l-4.172 4.173A5.057 5.057 0 0 0 3.23 14.1c0 .822.192 1.616.558 2.327l-.479.48A3.078 3.078 0 0 0 2.4 19.1c0 .827.322 1.605.908 2.191c.584.586 1.363.908 2.191.908s1.605-.322 2.191-.908l.48-.48a5.082 5.082 0 0 0 2.328.559a5.06 5.06 0 0 0 3.605-1.492l4.172-4.172a5.06 5.06 0 0 0 1.492-3.605c0-.824-.192-1.617-.558-2.328l.479-.48A3.076 3.076 0 0 0 20.6 7.1c0-.828-.322-1.606-.908-2.192a3.073 3.073 0 0 0-2.192-.909m-6.1 7.169c.017.535.233 1.036.613 1.416c.381.38.881.598 1.416.614l-1.832 1.831a2.089 2.089 0 0 0-.613-1.415A2.088 2.088 0 0 0 9.568 13zm1.1-2.139a.898.898 0 0 0-.637.262l-4.172 4.172a.894.894 0 0 0-.26.637c0 .24.092.467.26.635l.309.308l.723-.723c.215-.215.495-.321.777-.321s.562.106.777.321a1.1 1.1 0 0 1 0 1.557l-.72.723l.308.308c.168.168.401.253.636.253s.468-.084.637-.253l4.172-4.173a.893.893 0 0 0 .26-.635a.894.894 0 0 0-.26-.637l-.31-.309l-.723.723c-.215.215-.495.322-.777.322s-.562-.107-.777-.322a1.1 1.1 0 0 1 0-1.557l.72-.72l-.307-.309a.897.897 0 0 0-.636-.262"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Location(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.657 5.304c-3.124-3.073-8.189-3.073-11.313 0a7.78 7.78 0 0 0 0 11.13L12 21.999l5.657-5.565a7.78 7.78 0 0 0 0-11.13M12 13.499c-.668 0-1.295-.26-1.768-.732a2.503 2.503 0 0 1 0-3.536c.472-.472 1.1-.732 1.768-.732s1.296.26 1.768.732a2.503 2.503 0 0 1 0 3.536c-.472.472-1.1.732-1.768.732"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationArrow(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.368 19.102c.349 1.049 1.011 1.086 1.478.086l5.309-11.375c.467-1.002.034-1.434-.967-.967L4.812 12.154c-1.001.467-.963 1.129.085 1.479L9 15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationArrowOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.087 20.914c-.353 0-1.219-.146-1.668-1.496L8.21 15.791l-3.628-1.209c-1.244-.415-1.469-1.172-1.493-1.587s.114-1.193 1.302-1.747l11.375-5.309c1.031-.479 1.922-.309 2.348.362c.224.351.396.97-.053 1.933l-5.309 11.375c-.529 1.135-1.272 1.305-1.665 1.305m-5.39-8.068l4.094 1.363l1.365 4.093l4.775-10.233z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 5c1.609 0 3.12.614 4.254 1.73C17.38 7.837 18 9.309 18 10.87s-.62 3.03-1.745 4.139L12 19.193l-4.254-4.186c-1.125-1.107-1.745-2.576-1.745-4.139s.62-3.032 1.745-4.141A6.04 6.04 0 0 1 12 5m0-2a8.04 8.04 0 0 0-5.657 2.305a7.782 7.782 0 0 0 0 11.131L12 21.999l5.657-5.565a7.78 7.78 0 0 0 0-11.129A8.039 8.039 0 0 0 12 3m0 5.499c.668 0 1.296.26 1.768.731a2.502 2.502 0 0 1 0 3.537c-.473.472-1.1.731-1.768.731s-1.295-.26-1.768-.731a2.502 2.502 0 0 1 0-3.537A2.49 2.49 0 0 1 12 8.499m0-1a3.501 3.501 0 1 0 2.475 5.975a3.503 3.503 0 0 0 0-4.951A3.489 3.489 0 0 0 12 7.499"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockClosed(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 10h-1V8c0-2.205-1.794-4-4-4S8 5.795 8 8v2H7c-1.103 0-2 .896-2 2v7c0 1.104.897 2 2 2h10c1.103 0 2-.896 2-2v-7c0-1.104-.897-2-2-2m-5 8.299a1.3 1.3 0 1 1 0-2.6a1.3 1.3 0 0 1 0 2.6M14 11h-4V8c0-1.104.897-2 2-2s2 .896 2 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockClosedOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="12" cy="17" r="1.3" fill="currentColor"/><path fill="currentColor" d="M17 10h-1V8c0-2.206-1.794-4-4-4S8 5.794 8 8v2H7a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-7a2 2 0 0 0-2-2m-7-2a2 2 0 0 1 4 0v3h-4zm7 11H7v-7h10.003z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockOpen(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 4c-2.206 0-4 1.795-4 4v3h-4v-1H7c-1.103 0-2 .896-2 2v7c0 1.104.897 2 2 2h10c1.103 0 2-.896 2-2v-7c0-1.104-.897-2-2-2h-1V8c0-1.104.897-2 2-2s2 .896 2 2v3a1 1 0 1 0 2 0V8c0-2.205-1.794-4-4-4m-6 14.299a1.3 1.3 0 1 1 0-2.6a1.3 1.3 0 0 1 0 2.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockOpenOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="12" cy="17" r="1.3" fill="currentColor"/><path fill="currentColor" d="M18 4c-2.206 0-4 1.794-4 4v3h-4v-1H7a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-7a2 2 0 0 0-2-2h-1V8a2 2 0 0 1 4 0v3a1 1 0 0 0 2 0V8c0-2.206-1.794-4-4-4m-1 15H7v-7h10.003z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mail(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 7H5a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2m-9.684 7.316l1.602 1.4c.305.266.691.398 1.082.398s.777-.133 1.082-.398l1.602-1.4l-.037.037l3.646 3.646H5.707l3.646-3.646zM5 17.293V10.54l3.602 3.151zm10.398-3.602L19 10.54v6.75zM19 9v.21l-6.576 5.754a.68.68 0 0 1-.848 0L5 9.21V9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Map(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.383 3.076a1 1 0 0 0-1.09.217L15.426 7.16l-4.301-3.441a1.002 1.002 0 0 0-1.332.074l-4.5 4.5A.996.996 0 0 0 5 9v10a.999.999 0 0 0 1.707.707l3.867-3.867l4.301 3.441c.396.316.971.285 1.332-.074l4.5-4.5A.996.996 0 0 0 21 14V4a.999.999 0 0 0-.617-.924M7 16.586V9.414l3-3v7.24c-.07.043-3 2.932-3 2.932m4.125-2.867L11 13.651V6.182s3.959 3.143 4 3.166v7.473zM19 13.586l-3 3V9.35c.07-.043 3-2.936 3-2.936z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaEject(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 16H7a2 2 0 0 0 0 4h10a2 2 0 0 0 0-4m1.433-5.396L12 4l-6.433 6.604A2 2 0 0 0 7 14h10a2 2 0 0 0 1.433-3.396"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaEjectOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 21H8c-1.654 0-3-1.346-3-3s1.346-3 3-3h8c1.654 0 3 1.346 3 3s-1.346 3-3 3m-8-4a1.001 1.001 0 0 0 0 2h8a1.001 1.001 0 0 0 0-2zm4-10.134l4.964 5.096L17 12l-10 .004l.08-.087zM12 4l-6.433 6.604A2 2 0 0 0 7 14h10a2 2 0 0 0 2-2c0-.543-.218-1.033-.568-1.393z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaFastForward(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.053 6.912A1.797 1.797 0 0 0 12 8.201v9a1.8 1.8 0 0 0 3.053 1.287C17.434 16.174 21 12.701 21 12.701s-3.566-3.474-5.947-5.789m-9 0A1.797 1.797 0 0 0 3 8.201v9a1.8 1.8 0 0 0 3.053 1.287C8.434 16.174 12 12.701 12 12.701S8.434 9.227 6.053 6.912"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaFastForwardOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14 8.676l4.133 4.025L14 16.727zM13.8 6.4c-.994 0-1.8.807-1.8 1.801v9a1.8 1.8 0 0 0 3.053 1.287C17.434 16.174 21 12.701 21 12.701s-3.566-3.475-5.944-5.789A1.81 1.81 0 0 0 13.8 6.4M5 8.676l4.133 4.025L5 16.727zM4.8 6.4c-.994 0-1.8.807-1.8 1.801v9a1.8 1.8 0 0 0 3.053 1.287C8.434 16.174 12 12.701 12 12.701L6.056 6.912A1.81 1.81 0 0 0 4.8 6.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaPause(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 6a2 2 0 0 0-2 2v8a2 2 0 0 0 4 0V8a2 2 0 0 0-2-2m7 0a2 2 0 0 0-2 2v8a2 2 0 0 0 4 0V8a2 2 0 0 0-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaPauseOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 20c-1.654 0-3-1.346-3-3V8c0-1.654 1.346-3 3-3s3 1.346 3 3v9c0 1.654-1.346 3-3 3M8 7c-.552 0-1 .449-1 1v9a1.001 1.001 0 0 0 2 0V8c0-.551-.448-1-1-1m7 13c-1.654 0-3-1.346-3-3V8c0-1.654 1.346-3 3-3s3 1.346 3 3v9c0 1.654-1.346 3-3 3m0-13c-.552 0-1 .449-1 1v9a1.001 1.001 0 0 0 2 0V8c0-.551-.448-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaPlay(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.396 18.433L17 12l-6.604-6.433A2 2 0 0 0 7 7v10a2 2 0 0 0 3.396 1.433"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaPlayOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8.998 7.002l.085.078L14.134 12l-5.096 4.964L9 17zM9 5a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2c.543 0 1.033-.218 1.393-.568L17 12l-6.604-6.433A2.008 2.008 0 0 0 9 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaPlayReverse(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 19c1.1 0 2-.9 2-2V7c0-1.1-.9-2-2-2c-.5 0-1 .2-1.4.6C10 8.1 6 12 6 12s4 3.9 6.6 6.4c.4.4.9.6 1.4.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaPlayReverseOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 7v10l-5.1-5zm-1.4-1.4C10 8.1 6 12 6 12s4 3.9 6.6 6.4c.4.4.9.6 1.4.6c1.1 0 2-.9 2-2V7c0-1.1-.9-2-2-2c-.5 0-1 .2-1.4.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaRecord(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 12a5.985 5.985 0 0 0-1.757-4.243A5.985 5.985 0 0 0 12 6a5.985 5.985 0 0 0-4.242 1.757A5.982 5.982 0 0 0 6 12c0 1.656.672 3.156 1.758 4.242S10.344 18 12 18a5.982 5.982 0 0 0 4.243-1.758A5.985 5.985 0 0 0 18 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaRecordOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 8c2.205 0 4 1.794 4 4s-1.795 4-4 4s-4-1.794-4-4s1.795-4 4-4m0-2a6 6 0 1 0 0 12a6 6 0 0 0 0-12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaRewind(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.2 6.4a1.79 1.79 0 0 0-1.253.512C6.566 9.227 3 12.701 3 12.701l5.944 5.789A1.802 1.802 0 0 0 12 17.201v-9c0-.994-.806-1.801-1.8-1.801m9 0a1.79 1.79 0 0 0-1.253.512C15.566 9.227 12 12.701 12 12.701l5.944 5.789A1.802 1.802 0 0 0 21 17.201v-9c0-.994-.806-1.801-1.8-1.801"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaRewindOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 8.676v8.05l-4.133-4.025zm.2-2.276a1.79 1.79 0 0 0-1.253.512C6.566 9.227 3 12.701 3 12.701l5.944 5.789A1.802 1.802 0 0 0 12 17.201v-9c0-.994-.806-1.801-1.8-1.801M19 8.676v8.051l-4.133-4.025zm.2-2.276a1.79 1.79 0 0 0-1.253.512C15.566 9.227 12 12.701 12 12.701l5.944 5.789A1.802 1.802 0 0 0 21 17.201v-9c0-.994-.806-1.801-1.8-1.801"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaStop(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 6H8c-1.1 0-2 .9-2 2v8c0 1.1.9 2 2 2h8c1.1 0 2-.9 2-2V8c0-1.1-.9-2-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaStopOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 8v8H8V8zm0-2H8c-1.1 0-2 .9-2 2v8c0 1.1.9 2 2 2h8c1.1 0 2-.9 2-2V8c0-1.1-.9-2-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Message(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 7c.542 0 1 .458 1 1v7c0 .542-.458 1-1 1H9.171L9 16.171V16H6c-.542 0-1-.458-1-1V8c0-.542.458-1 1-1zm0-2H6C4.35 5 3 6.35 3 8v7c0 1.65 1.35 3 3 3h1v3l3-3h8c1.65 0 3-1.35 3-3V8c0-1.65-1.35-3-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageTyping(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 6H5C3.35 6 2 7.35 2 9v7c0 1.65 1.35 3 3 3h1v3l3-3h9c1.65 0 3-1.35 3-3V9c0-1.65-1.35-3-3-3m1 10c0 .542-.458 1-1 1H5c-.542 0-1-.458-1-1V9c0-.542.458-1 1-1h13c.542 0 1 .458 1 1zM7 14.5a2 2 0 1 1 .001-4.001A2 2 0 0 1 7 14.5m0-3a1 1 0 1 0 0 2a1 1 0 0 0 0-2m4.5 3a2 2 0 1 1 .001-4.001A2 2 0 0 1 11.5 14.5m0-3a1 1 0 1 0 0 2a1 1 0 0 0 0-2m4.5 3a2 2 0 1 1 .001-4.001A2 2 0 0 1 16 14.5m0-3a1 1 0 1 0 0 2a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Messages(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 7h-3c0-1.65-1.35-3-3-3H3C1.35 4 0 5.35 0 7v7c0 1.65 1.35 3 3 3v3l3-3c0 1.65 1.35 3 3 3h8l3 3v-3h1c1.65 0 3-1.35 3-3v-7c0-1.65-1.35-3-3-3M3 15c-.542 0-1-.458-1-1V7c0-.542.458-1 1-1h12c.542 0 1 .458 1 1v1H9.5A2.502 2.502 0 0 0 7 10.5V15zm19 2c0 .542-.458 1-1 1H9c-.542 0-1-.458-1-1v-6.5C8 9.673 8.673 9 9.5 9H21c.542 0 1 .458 1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Microphone(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 16c2.206 0 4-1.795 4-4V6c0-2.206-1.794-4-4-4S8 3.794 8 6v6c0 2.205 1.794 4 4 4m7-4v-2a1 1 0 1 0-2 0v2c0 2.757-2.243 5-5 5s-5-2.243-5-5v-2a1 1 0 1 0-2 0v2c0 3.52 2.613 6.432 6 6.92V20H8a1 1 0 1 0 0 2h8a1 1 0 1 0 0-2h-3v-1.08c3.387-.488 6-3.4 6-6.92"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicrophoneOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 16c-2.206 0-4-1.795-4-4V6c0-2.205 1.794-4 4-4s4 1.795 4 4v6c0 2.205-1.794 4-4 4m0-12c-1.103 0-2 .896-2 2v6c0 1.104.897 2 2 2s2-.896 2-2V6c0-1.104-.897-2-2-2m7 8v-2a1 1 0 1 0-2 0v2c0 2.757-2.243 5-5 5s-5-2.243-5-5v-2a1 1 0 1 0-2 0v2c0 3.52 2.613 6.432 6 6.92V20H8a1 1 0 1 0 0 2h8a1 1 0 1 0 0-2h-3v-1.08c3.387-.488 6-3.4 6-6.92"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Minus(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 11H6a2 2 0 0 0 0 4h12a2 2 0 0 0 0-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 16H6c-1.654 0-3-1.346-3-3s1.346-3 3-3h12c1.654 0 3 1.346 3 3s-1.346 3-3 3M6 12c-.551 0-1 .449-1 1s.449 1 1 1h12c.551 0 1-.449 1-1s-.449-1-1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MortarBoard(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.5 7.9S14.7 5 13.6 4.4s-1.7-.7-3 0C9.3 5 3.9 7.7 3.9 7.7c-.8.4-1.5 1.2-1.5 2s.2 1.2.2 1.2s-.1.3-.3 1.5C2 13.6 2 15.1 2 15.7c0 1.5 1.3 2.6 2.2 2.7c.9.1 1.6-.1 1.6-.1c1.4 1.3 3.7 2.1 6.4 2.1c4.4 0 7.8-2.2 7.8-5c0-1.1-.4-2.7-.4-2.7l1.1-.6c.9-.5 1.3-1.4 1.3-2.3c-.1-.8-.6-1.5-1.5-1.9m-8.2 10.4c-3.2 0-5.8-1.3-5.8-3l.5-2.8l4.2 2.1c.6.3 1.5.3 2.2 0l4.3-2.1l.4 2.8c0 1.6-2.5 3-5.8 3m7.3-8.1L13 13.6c-.4.2-1 .2-1.4 0l-5.7-2.9c-.2.5-.3 1.2-.3 2c0 1.4.2 2.4.2 2.9s-.3.8-.7.8H5c-.4 0-.8-.3-.8-.8s0-1.6.3-3.1c.2-.9.4-1.7.6-2.2l-.2-.1c-.4-.2-.4-.5 0-.7l6.7-3.4c.4-.2.9-.2 1.3 0s6.7 3.4 6.7 3.4c.4.2.4.5 0 .7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func News(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 4H3a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h18a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2M3 6h8v12H3zm18 12h-9V6h9.003zm-1-4.5c0-.275-.225-.5-.5-.5h-1c-.275 0-.5.225-.5.5v3c0 .275.225.5.5.5h1c.275 0 .5-.225.5-.5zm-3-6c0-.275-.225-.5-.5-.5h-3c-.275 0-.5.225-.5.5v5c0 .275.225.5.5.5h3c.275 0 .5-.225.5-.5zm1.5 2.5h1c.275 0 .5-.225.5-.5s-.225-.5-.5-.5h-1c-.275 0-.5.225-.5.5s.225.5.5.5m0 2h1c.275 0 .5-.225.5-.5s-.225-.5-.5-.5h-1c-.275 0-.5.225-.5.5s.225.5.5.5m-5 3h3c.275 0 .5-.225.5-.5s-.225-.5-.5-.5h-3c-.275 0-.5.225-.5.5s.225.5.5.5m3 1h-3c-.275 0-.5.225-.5.5s.225.5.5.5h3c.275 0 .5-.225.5-.5s-.225-.5-.5-.5m2-8h1c.275 0 .5-.225.5-.5s-.225-.5-.5-.5h-1c-.275 0-.5.225-.5.5s.225.5.5.5M10 7.5c0-.275-.225-.5-.5-.5h-5c-.275 0-.5.225-.5.5v3c0 .275.225.5.5.5h5c.275 0 .5-.225.5-.5zM9.501 14h-5c-.274 0-.5.225-.5.5s.226.5.5.5h5c.274 0 .499-.225.499-.5s-.225-.5-.499-.5m0-2h-5c-.274 0-.5.225-.5.5s.226.5.5.5h5c.274 0 .499-.225.499-.5s-.225-.5-.499-.5m0 4h-5c-.274 0-.5.225-.5.5s.226.5.5.5h5c.274 0 .499-.225.499-.5s-.225-.5-.499-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Notes(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.831 4.059a.49.49 0 0 0-.394-.121l-11 1.25A.5.5 0 0 0 7 5.684V16c-1.654 0-3 1.122-3 2.5S5.346 21 7 21s3-1.122 3-2.5v-7.609l6-.625V14c-1.654 0-3 1.122-3 2.5s1.346 2.5 3 2.5s3-1.122 3-2.5V4.434a.501.501 0 0 0-.169-.375"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotesOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.324 4.367c-.368-.324-.84-.5-1.324-.5l-.248.016l-9 1.25A1.999 1.999 0 0 0 6 7.117v6.111c-1.746.551-3 2.034-3 3.772c0 2.205 2.019 4 4.5 4c1.695 0 3.169-.842 3.937-2.078A4.788 4.788 0 0 0 14.5 20c2.481 0 4.5-1.795 4.5-4V5.867c0-.574-.246-1.119-.676-1.5M11 16v-4.256l3-.45v1.737c-1.693.208-3 1.46-3 2.969m6 0c0 1.104-1.119 2-2.5 2s-2.5-.896-2.5-2s1.119-2 2.5-2c.172 0 .338.014.5.041v-3.908l-5 .75V17c0 1.104-1.119 2-2.5 2S5 18.104 5 17s1.119-2 2.5-2c.172 0 .338.014.5.041V7.117l9-1.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pen(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.329 7.207c0-1.212-.472-2.352-1.329-3.207s-1.996-1.329-3.207-1.329a4.5 4.5 0 0 0-3.18 1.304c-.027.025-7.967 7.964-7.967 7.964c-.373.373-.717.91-.967 1.514c-.195.473-1.979 5.863-2.336 6.939a1 1 0 0 0 1.263 1.263c1.076-.355 6.465-2.141 6.938-2.336c.603-.248 1.14-.592 1.515-.967l2.157-2.156l.076.01c.64 0 1.28-.244 1.769-.732l4.5-4.5a2.487 2.487 0 0 0 .588-2.572c.107-.386.18-.783.18-1.195M9.465 17.586c-.406.143-1.145.393-2.038.691l-1.704-1.704c.301-.894.551-1.634.691-2.038zm-4.097.047l.999.999l-1.498.499zm7.698-3.113l-2.42 2.42l-.235.18l-3.53-3.529l.18-.234l7.131-7.131l2.731 2.731l-3.69 3.69c-.513.512-.549 1.289-.167 1.873m6.08-4.959l-4.5 4.5a.502.502 0 0 1-.708 0a.5.5 0 0 1 0-.707l4.5-4.5a.502.502 0 0 1 .707 0a.499.499 0 0 1 .001.707m.107-1.764a1.489 1.489 0 0 0-1.521.35l-.102.102l-2.731-2.731l.078-.078c.984-.98 2.652-.981 3.608-.023a2.52 2.52 0 0 1 .743 1.793c.001.199-.03.394-.075.587"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pencil(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 6.879L17.121 3A1.497 1.497 0 0 0 15 3L4.061 13.939c-.293.293-.558.727-.75 1.188C3.119 15.59 3 16.086 3 16.5V21h4.5c.414 0 .908-.119 1.371-.311c.463-.192.896-.457 1.189-.75L21 9a1.497 1.497 0 0 0 0-2.121M5.768 15.061l8.293-8.293L15.293 8L7 16.293zM7.5 19H6l-1-1v-1.5c0-.077.033-.305.158-.605c.01-.02 2.967 2.938 2.967 2.938c-.322.134-.548.167-.625.167m1.439-.768L7.707 17L16 8.707l1.232 1.232zm9-9L14.767 6.06l1.293-1.293l3.17 3.172z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Phone(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.374 7.083l3.711-3.71l-.438-.434c-.566-.566-1.555-.566-2.121 0L12.94 4.525c-.284.284-.44.661-.44 1.061s.156.777.438 1.06zm-6.728 5.856c-.566-.566-1.555-.566-2.121 0l-1.586 1.586c-.283.284-.439.661-.439 1.061s.156.777.441 1.062l.437.432l3.703-3.703zm11.791-8.21l-.354-.354l-3.708 3.708l.65.649a.5.5 0 0 1 0 .708l-5.586 5.586a.513.513 0 0 1-.707 0l-.65-.65l-3.702 3.71l.354.354c.26.26 1.246 1.105 3.056 1.105c1.616 0 4.256-.712 7.65-4.105c6.773-6.775 3.158-10.55 2.997-10.711"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m19.502 3.672l-1.795-1.793C17.141 1.312 16.387 1 15.586 1s-1.555.312-2.121.879l-1.586 1.586a3.002 3.002 0 0 0 0 4.242l1.379 1.379l-4.172 4.172l-1.379-1.379C7.141 11.312 6.387 11 5.586 11s-1.555.312-2.121.879l-1.586 1.586a3.002 3.002 0 0 0 0 4.242L3.673 19.5c.465.465 1.796 1.545 4.116 1.545c2.764 0 5.694-1.529 8.711-4.545c6.245-6.246 4.825-11.002 3.002-12.828m-6.209 1.207l1.586-1.586a.997.997 0 0 1 1.414 0l1.083 1.082l-3.001 3l-1.082-1.082a.999.999 0 0 1 0-1.414m-10 11.414a.999.999 0 0 1 0-1.414l1.586-1.586a.997.997 0 0 1 1.414 0l1.082 1.082l-2.999 3zm11.793-1.207c-3.083 3.082-5.551 3.959-7.297 3.959c-1.349 0-2.267-.523-2.702-.959c-.004-.004 2.995-3.004 2.995-3.004l.297.297a.997.997 0 0 0 1.414 0l5.586-5.586a.999.999 0 0 0 0-1.414l-.297-.297l3.001-3c1.003 1.004 2.467 4.539-2.997 10.004"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pi(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.707 8.535a.999.999 0 0 0-1.414 0a3.247 3.247 0 0 1-4.586 0a5.25 5.25 0 0 0-7.414 0a.999.999 0 1 0 1.414 1.414c.374-.374.82-.624 1.293-.776V17a1 1 0 1 0 2 0V9.174a3.19 3.19 0 0 1 1.293.775A5.222 5.222 0 0 0 14 11.386V17a1 1 0 1 0 2 0v-5.614a5.215 5.215 0 0 0 2.707-1.437a.999.999 0 0 0 0-1.414"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PiOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.121 7.121c-.566-.567-1.32-.879-2.121-.879s-1.555.312-2.121.879c-.233.233-.546.362-.879.362s-.646-.129-.879-.362C12.755 5.755 10.936 5.003 9 5.003s-3.755.752-5.121 2.118C3.312 7.688 3 8.441 3 9.242s.312 1.555.879 2.121c.566.567 1.32.879 2.121.879V17c0 1.654 1.346 3 3 3s3-1.346 3-3c0 1.654 1.346 3 3 3s3-1.346 3-3v-4.166a7.21 7.21 0 0 0 2.12-1.47c.568-.567.88-1.321.88-2.122s-.312-1.554-.879-2.121m-1.414 2.828A5.222 5.222 0 0 1 16 11.386V17a1 1 0 1 1-2 0v-5.614a5.215 5.215 0 0 1-2.707-1.437A3.19 3.19 0 0 0 10 9.174V17a1 1 0 1 1-2 0V9.173a3.186 3.186 0 0 0-1.293.776a.997.997 0 0 1-1.414 0a.999.999 0 0 1 0-1.414C6.314 7.514 7.657 7.003 9 7.003s2.685.511 3.707 1.532A3.235 3.235 0 0 0 15 9.483c.831 0 1.661-.316 2.293-.948a.997.997 0 0 1 1.414 0a.999.999 0 0 1 0 1.414"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pin(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.729 4.271a1 1 0 0 0-1.414-.004a1.004 1.004 0 0 0-.225.355c-.832 1.736-1.748 2.715-2.904 3.293C10.889 8.555 9.4 9 7 9a1.006 1.006 0 0 0-.923.617a1.001 1.001 0 0 0 .217 1.09l3.243 3.243L5 20l6.05-4.537l3.242 3.242a.975.975 0 0 0 .326.217c.122.051.252.078.382.078s.26-.027.382-.078A.996.996 0 0 0 16 18c0-2.4.444-3.889 1.083-5.166c.577-1.156 1.556-2.072 3.293-2.904a.983.983 0 0 0 .354-.225a1 1 0 0 0-.004-1.414z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.436 7.586l-3.998-4.02c-.752-.756-2.063-.764-2.83-.006c-.196.196-.35.436-.418.629c-.653 1.362-1.354 2.215-2.254 2.727l-.217.105C10.751 7.506 9.434 8 7 8c-.266 0-.521.052-.766.152a2.022 2.022 0 0 0-1.082 1.084a2.022 2.022 0 0 0 0 1.525c.104.249.25.471.435.651l3.235 3.235L5 20l5.352-3.822l3.227 3.227c.186.189.406.339.656.441a1.974 1.974 0 0 0 1.531 0a1.963 1.963 0 0 0 1.08-1.078c.103-.242.155-.507.155-.768c0-2.436.494-3.752.978-4.721c.496-.992 1.369-1.748 2.754-2.414c.271-.104.51-.256.711-.457a2.005 2.005 0 0 0-.008-2.822m-5.248 4.801c-.819 1.643-1.188 3.37-1.195 5.604L7 10c2.139 0 3.814-.335 5.396-1.084l.235-.105c1.399-.699 2.468-1.893 3.388-3.834l3.924 4.051c-1.863.893-3.056 1.96-3.755 3.359"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pipette(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.384 7.331c.073-1.199-.354-2.388-1.146-3.179C19.506 3.421 18.445 3 17.326 3c-1.176 0-2.206.453-2.825 1.243c-.692.883-1.392 2.625-1.769 3.647l-1.616-1.617a.999.999 0 0 0-1.414 0a.997.997 0 0 0 0 1.414l.293.293l-5.231 5.232c-.375.375-.719.912-.968 1.516c-.019.043-1.726 4.328-.093 5.959c.527.526 1.33.707 2.178.707c1.778-.002 3.753-.787 3.783-.801c.602-.248 1.141-.592 1.514-.967l5.232-5.232l.293.293a.997.997 0 0 0 1.414 0a.999.999 0 0 0 0-1.414L16.5 11.657c1.023-.376 2.766-1.075 3.648-1.769c.721-.562 1.17-1.493 1.236-2.557M5.119 19.275c-.247-.295-.105-1.508.154-2.58l2.422 2.423c-1.071.261-2.283.403-2.576.157m4.645-1.061c-.188.188-.511.388-.865.533l-.116.042l-3.181-3.18l.043-.117c.146-.354.346-.678.533-.864l5.232-5.231l3.586 3.586z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plane(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.996 13.507L14 10.081V4.125c0-.827-.673-1.5-1.5-1.5s-1.5.673-1.5 1.5v5.956l-5.996 3.426a1 1 0 0 0 .77 1.829L11 13.844v4.451l-1.625 1.3a1 1 0 0 0 .996 1.709l2.129-.852l2.129.852a.998.998 0 0 0 1.235-.426a1.001 1.001 0 0 0-.239-1.284L14 18.294v-4.451l5.226 1.493l.274.039a1 1 0 0 0 .496-1.868M12.5 4.375a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlaneOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.988 12.396L16 9.545V4.75c0-1.93-1.57-3.5-3.5-3.5S9 2.82 9 4.75v4.795l-4.988 2.851a3 3 0 0 0 2.313 5.489L9 17.12v.838l-.874.699a3 3 0 0 0 3.216 5.025c.004-.001.5-.183 1.158-.183l1.158.183a2.998 2.998 0 0 0 3.216-5.025L16 17.958v-.838l2.676.765a2.999 2.999 0 0 0 2.312-5.489m-.566 2.992a1 1 0 0 1-1.196.573L14 14.469v4.451l1.625 1.3a1 1 0 0 1-1.072 1.675c-.008-.004-.824-.395-2.053-.395s-2.045.391-2.053.395a1 1 0 0 1-1.072-1.675L11 18.92v-4.451l-5.226 1.493a1 1 0 0 1-.77-1.829L11 10.706V4.75c0-.827.673-1.5 1.5-1.5s1.5.673 1.5 1.5v5.956l5.996 3.426a1 1 0 0 1 .426 1.256"/><circle cx="12.5" cy="4.5" r=".5" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plug(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 6h-1V3c0-.6-.4-1-1-1h-2c-.6 0-1 .4-1 1v3h-2V3c0-.6-.4-1-1-1H8c-.6 0-1 .4-1 1v3H6c-.6 0-1 .4-1 1v4.2c.2 2.5 1.8 4.6 4 5.6V20c0 1.1.9 2 2 2h2c1.1 0 2-.9 2-2v-3.2c2.2-1 3.7-3.1 4-5.6V7c0-.6-.4-1-1-1m-4-3h2v3h-2zM8 3h2v3H8zm5 17h-2v-2h2zm-1-4.5c-2.2 0-4.1-1.5-4.7-3.5h9.5c-.7 2-2.6 3.5-4.8 3.5m5-5c0 .2 0 .3-.1.5H7.1c-.1-.2-.1-.3-.1-.5V8h10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plus(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 10h-4V6a2 2 0 0 0-4 0l.071 4H6a2 2 0 0 0 0 4l4.071-.071L10 18a2 2 0 0 0 4 0v-4.071L18 14a2 2 0 0 0 0-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21c-1.654 0-3-1.346-3-3l.053-3.053L6.018 15C4.346 15 3 13.654 3 12s1.346-3 3-3l3.053-.054L9 6.018C9 4.346 10.346 3 12 3s3 1.346 3 3l.055 2.946L18.018 9C19.654 9 21 10.346 21 12s-1.346 3-3 3l-2.945-.053L15 18.018C15 19.654 13.654 21 12 21m-1-8v5.018c0 .533.449.982 1 .982s1-.449 1-1v-5h5.018c.533 0 .982-.449.982-1s-.449-1-1-1h-5V6c0-.569-.449-1-1-1s-1 .449-1 1v5H6c-.569 0-1 .449-1 1s.449 1 1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PointOfInterest(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.5 11c1.93 0 3.5-1.57 3.5-3.5S18.43 4 16.5 4S13 5.57 13 7.5V9h-2V7.5C11 5.57 9.43 4 7.5 4S4 5.57 4 7.5S5.57 11 7.5 11H9v2H7.5C5.57 13 4 14.57 4 16.5S5.57 20 7.5 20s3.5-1.57 3.5-3.5V15h2v1.5c0 1.93 1.57 3.5 3.5 3.5s3.5-1.57 3.5-3.5s-1.57-3.5-3.5-3.5H15v-2zM15 7.5a1.501 1.501 0 0 1 3 0c0 .826-.673 1.5-1.5 1.5H15zm-6 9c0 .826-.673 1.5-1.5 1.5S6 17.326 6 16.5c0-.828.673-1.5 1.5-1.5H9zM9 9H7.5A1.501 1.501 0 1 1 9 7.5zm4 4h-2v-2h2zm3.5 2a1.501 1.501 0 0 1 0 3c-.827 0-1.5-.674-1.5-1.5V15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PointOfInterestOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.5 4C18.43 4 20 5.57 20 7.5S18.43 11 16.5 11H15v2h1.5c1.93 0 3.5 1.57 3.5 3.5S18.43 20 16.5 20S13 18.43 13 16.5V15h-2v1.5c0 1.93-1.57 3.5-3.5 3.5S4 18.43 4 16.5S5.57 13 7.5 13H9v-2H7.5C5.57 11 4 9.43 4 7.5S5.57 4 7.5 4S11 5.57 11 7.5V9h2V7.5C13 5.57 14.57 4 16.5 4M15 9h1.5A1.501 1.501 0 1 0 15 7.5zM7.5 9H9V7.5a1.501 1.501 0 0 0-3 0C6 8.326 6.673 9 7.5 9m9 9a1.501 1.501 0 0 0 0-3H15v1.5c0 .826.673 1.5 1.5 1.5m-9 0c.827 0 1.5-.674 1.5-1.5V15H7.5a1.501 1.501 0 0 0 0 3m9-16A5.498 5.498 0 0 0 12 4.341A5.498 5.498 0 0 0 7.5 2A5.506 5.506 0 0 0 2 7.5c0 1.857.926 3.504 2.341 4.5A5.498 5.498 0 0 0 2 16.5C2 19.532 4.467 22 7.5 22a5.498 5.498 0 0 0 4.5-2.341A5.498 5.498 0 0 0 16.5 22c3.033 0 5.5-2.468 5.5-5.5a5.498 5.498 0 0 0-2.341-4.5A5.498 5.498 0 0 0 22 7.5C22 4.468 19.533 2 16.5 2M13 11h-2v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Power(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.5 18.573a6.46 6.46 0 0 1-4.596-1.903C5.677 15.442 5 13.81 5 12.073s.677-3.369 1.904-4.597A.999.999 0 1 1 8.318 8.89C7.468 9.741 7 10.871 7 12.073s.468 2.333 1.318 3.183c.85.85 1.979 1.317 3.182 1.317s2.332-.468 3.182-1.317c.851-.85 1.318-1.98 1.318-3.183s-.468-2.333-1.318-3.183a.999.999 0 1 1 1.414-1.414C17.323 8.705 18 10.337 18 12.073s-.677 3.369-1.904 4.597a6.46 6.46 0 0 1-4.596 1.903m0-7.573a1 1 0 0 1-1-1V5a1 1 0 1 1 2 0v5a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PowerOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.51 6.062a2.984 2.984 0 0 0-3.01-.729V5a3.001 3.001 0 0 0-6 0v.332a2.99 2.99 0 0 0-3.01.73A8.45 8.45 0 0 0 3 12.073a8.45 8.45 0 0 0 2.49 6.011c1.604 1.605 3.739 2.489 6.01 2.489s4.405-.884 6.01-2.489c1.605-1.605 2.49-3.74 2.49-6.011s-.885-4.405-2.49-6.011M10.5 5a1 1 0 1 1 2 0v5a1 1 0 1 1-2 0zm-1 3.803V10a2 2 0 0 0 4 0V8.818c.095.284.248.554.475.78c.661.661 1.025 1.54 1.025 2.475s-.364 1.814-1.025 2.476c-1.322 1.321-3.627 1.321-4.949 0C8.364 13.887 8 13.008 8 12.073s.364-1.814 1.025-2.476c.231-.23.383-.504.475-.794m6.596 7.867c-1.228 1.228-2.859 1.903-4.596 1.903s-3.368-.676-4.596-1.903C5.677 15.442 5 13.81 5 12.073s.677-3.369 1.904-4.597A.999.999 0 1 1 8.318 8.89C7.468 9.741 7 10.871 7 12.073s.468 2.333 1.318 3.183c.85.85 1.979 1.317 3.182 1.317s2.332-.468 3.182-1.317c.851-.85 1.318-1.98 1.318-3.183s-.468-2.333-1.318-3.183a.999.999 0 1 1 1.414-1.414C17.323 8.705 18 10.337 18 12.073s-.677 3.369-1.904 4.597"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Printer(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 5V3a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1v2C4.346 5 3 6.346 3 8v10c0 1.654 1.346 3 3 3h11c1.654 0 3-1.346 3-3V8c0-1.654-1.346-3-3-3M8 4h7v5H8zM6 7v3a1 1 0 0 0 1 1h9a1 1 0 0 0 1-1V7c.551 0 1 .449 1 1v2.5c0 .827-.673 1.5-1.5 1.5h-10c-.827 0-1.5-.673-1.5-1.5V8c0-.551.449-1 1-1m11 12H6c-.551 0-1-.449-1-1v-5.513c.419.318.935.513 1.5.513h10c.565 0 1.081-.195 1.5-.513V18c0 .551-.449 1-1 1M13.5 7h-4a.5.5 0 0 0 0 1h4a.5.5 0 0 0 0-1m1.5 9H8a.5.5 0 0 0 0 1h7a.5.5 0 0 0 0-1M13.5 5h-4a.5.5 0 0 0 0 1h4a.5.5 0 0 0 0-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Puzzle(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.25 11.25c.364 0 .704.145.984.391c.549.332.766-.034.766-.391V9.5c0-.825-.675-1.5-1.5-1.5h-2.75c-.356 0-.724-.216-.391-.766c.246-.28.391-.619.391-.984c0-.967-1.007-1.75-2.25-1.75s-2.25.783-2.25 1.75c0 .3.095.576.255.823c.507.673.136.927-.255.927H6.5C5.675 8 5 8.675 5 9.5v1.75c0 .391.254.762.928.244c.246-.149.522-.244.822-.244c.966 0 1.75 1.008 1.75 2.25s-.784 2.25-1.75 2.25c-.364 0-.704-.145-.984-.391c-.549-.332-.766.034-.766.391v2.75c0 .825.675 1.5 1.5 1.5h2.75c.391 0 .762-.254.243-.927a1.593 1.593 0 0 1-.243-.823c0-.967 1.007-1.75 2.25-1.75s2.25.783 2.25 1.75c0 .365-.145.704-.391.984c-.333.55.035.766.391.766h2.75c.825 0 1.5-.675 1.5-1.5v-2.75c0-.391-.254-.762-.928-.244a1.579 1.579 0 0 1-.822.244c-.966 0-1.75-1.008-1.75-2.25s.784-2.25 1.75-2.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PuzzleOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 11.25V9.5C20 7.57 18.43 6 16.5 6h-.759C15.6 4.018 13.788 2.5 11.5 2.5S7.4 4.018 7.259 6H6.5C4.57 6 3 7.57 3 9.5v1.75c0 1.012.514 1.847 1.295 2.246a2.356 2.356 0 0 0-.894.825A2.664 2.664 0 0 0 3 15.75v2.75C3 20.43 4.57 22 6.5 22h2.75c.976 0 1.831-.497 2.242-1.299l.036.066c.435.772 1.266 1.233 2.222 1.233h2.75c1.93 0 3.5-1.57 3.5-3.5v-2.75c0-1.013-.515-1.849-1.297-2.247c.776-.411 1.297-1.256 1.297-2.253m-2 7.25c0 .825-.675 1.5-1.5 1.5h-2.75c-.356 0-.724-.216-.391-.766c.246-.28.391-.619.391-.984c0-.967-1.007-1.75-2.25-1.75s-2.25.783-2.25 1.75c0 .3.095.576.255.823c.507.673.136.927-.255.927H6.5c-.825 0-1.5-.675-1.5-1.5v-2.75c0-.258.113-.521.384-.521c.104 0 .229.039.382.13c.28.246.62.391.984.391c.966 0 1.75-1.008 1.75-2.25s-.784-2.25-1.75-2.25c-.3 0-.576.095-.822.255c-.237.171-.422.243-.562.243c-.26 0-.366-.245-.366-.498V9.5C5 8.675 5.675 8 6.5 8h2.75c.391 0 .762-.254.243-.927a1.593 1.593 0 0 1-.243-.823c0-.967 1.007-1.75 2.25-1.75s2.25.783 2.25 1.75c0 .365-.145.704-.391.984c-.333.55.035.766.391.766h2.75c.825 0 1.5.675 1.5 1.5v1.75c0 .258-.113.521-.384.521c-.104 0-.229-.039-.382-.13a1.488 1.488 0 0 0-.984-.391c-.966 0-1.75 1.008-1.75 2.25s.784 2.25 1.75 2.25c.3 0 .576-.095.822-.255c.237-.171.422-.244.562-.243c.259 0 .365.245.365.498v2.75zM5 12.694c.116.032.236.054.365.054c.342 0 .683-.119 1.038-.364l.069-.041a.507.507 0 0 1 .278-.093c.354 0 .75.535.75 1.25s-.396 1.25-.75 1.25a.496.496 0 0 1-.324-.142l-.143-.104a1.73 1.73 0 0 0-.899-.275c-.134 0-.261.023-.384.059zm12.635 1.558c-.342 0-.683.119-1.038.364l-.069.041a.506.506 0 0 1-.277.093c-.354 0-.75-.535-.75-1.25s.396-1.25.75-1.25c.108 0 .217.048.324.142l.143.104c.302.183.604.275.899.275c.136 0 .262-.025.384-.062v1.597a1.372 1.372 0 0 0-.366-.054M10.692 20c.101-.346.093-.816-.305-1.396l-.044-.074a.522.522 0 0 1-.094-.279c0-.354.534-.75 1.25-.75s1.25.396 1.25.75a.497.497 0 0 1-.143.325l-.104.142c-.325.537-.311.979-.22 1.284h-1.59z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Radar(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 20c3.86 0 7-3.141 7-7s-3.14-7-7.003-7C8.139 6 5 9.141 5 13s3.14 7 7 7M11 8.102V10a1 1 0 1 0 2 0V8.102A5.013 5.013 0 0 1 16.899 12H15a1 1 0 1 0 0 2h1.899A5.013 5.013 0 0 1 13 17.898V16a1 1 0 1 0-2 0v1.898A5.013 5.013 0 0 1 7.101 14H9a1 1 0 1 0 0-2H7.101A5.012 5.012 0 0 1 11 8.102"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RadarOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.997 4.5C7.312 4.5 3.5 8.312 3.5 13s3.813 8.5 8.5 8.5c4.688 0 8.5-3.812 8.5-8.5s-3.812-8.5-8.503-8.5m.003 15c-3.584 0-6.5-2.916-6.5-6.5S8.414 6.5 12 6.5c3.584 0 6.5 2.916 6.5 6.5s-2.916 6.5-6.5 6.5m3.348-7.469L15.5 12h.879A4.497 4.497 0 0 0 13 8.622V9.5c0 .551-.449 1-1 1a.989.989 0 0 1-.969-.846L11 9.5v-.88A4.5 4.5 0 0 0 7.62 12h.88l.153.031c.476.076.847.472.847.969s-.371.893-.846.969L8.5 14h-.878A4.506 4.506 0 0 0 11 17.379V16.5l.031-.154c.077-.476.472-.846.969-.846s.893.371.969.848L13 16.5v.879A4.503 4.503 0 0 0 16.379 14H15.5l-.152-.031c-.477-.076-.848-.472-.848-.969s.371-.893.848-.969m-.446 2.867a3.51 3.51 0 0 1-1.004 1.002c-.256-.81-1.004-1.401-1.897-1.401s-1.642.592-1.898 1.401c-.4-.262-.74-.603-1.003-1.002A1.995 1.995 0 0 0 10.501 13c0-.895-.592-1.643-1.402-1.898c.263-.399.603-.74 1.004-1.002a1.994 1.994 0 0 0 1.898 1.401c.894 0 1.644-.593 1.899-1.403c.399.264.74.604 1.002 1.004A1.995 1.995 0 0 0 13.501 13a1.992 1.992 0 0 0 1.401 1.898"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Refresh(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.872 13.191H18V8.064c-.008-1.135-.671-1.408-1.473-.605l-1.154 1.158a5.756 5.756 0 0 0-3.566-1.23c-1.55 0-3.009.604-4.104 1.701A5.748 5.748 0 0 0 6 13.191c0 1.553.604 3.012 1.701 4.107A5.77 5.77 0 0 0 11.807 19c1.55 0 3.009-.605 4.106-1.703c.296-.297.558-.621.78-.965a1.16 1.16 0 1 0-1.954-1.255c-.133.207-.292.4-.468.58a3.465 3.465 0 0 1-2.464 1.02a3.46 3.46 0 0 1-2.464-1.02a3.466 3.466 0 0 1-1.02-2.465c0-.93.362-1.805 1.02-2.461a3.466 3.466 0 0 1 2.464-1.021c.688 0 1.346.201 1.909.572l-1.448 1.451c-.803.802-.53 1.458.604 1.458"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RefreshOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.368 4.998c-.488 0-1.2.145-1.956.773A7.966 7.966 0 0 0 12 5a8 8 0 0 0 0 16c4.312 0 8-3.316 8-8V8.064c-.016-2.111-1.375-3.066-2.632-3.066M18 13h-5.128c-1.134 0-1.407-.561-.604-1.363l1.448-1.402a3.431 3.431 0 0 0-1.909-.549c-.93 0-1.805.375-2.464 1.033a3.47 3.47 0 0 0-1.02 2.467a3.47 3.47 0 0 0 1.02 2.469c.659.658 1.534 1.021 2.464 1.021s1.805-.36 2.465-1.019c.177-.18.334-.372.468-.579a1.161 1.161 0 0 1 1.955 1.256a5.962 5.962 0 0 1-.78.965a5.764 5.764 0 0 1-4.106 1.703a5.762 5.762 0 0 1-4.104-1.701a5.767 5.767 0 0 1-1.701-4.106c0-1.551.604-3.012 1.702-4.104a5.759 5.759 0 0 1 4.104-1.7c1.311 0 2.551.436 3.566 1.229l1.154-1.158c.311-.312.602-.461.841-.461c.377 0 .627.372.632 1.065V13zm-7.08.05c.162.392.63.95 1.952.95h1.299s-.21.504-.614.907s-1.086.745-1.75.745a2.459 2.459 0 0 1-2.485-2.467c0-.664.258-1.139.726-1.604c.472-.472 1.097-.581 1.759-.581l-.246.123c-.935.934-.803 1.536-.641 1.927"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rss(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.002 15.999a2 2 0 1 0-.004 4a2 2 0 0 0 .004-4M6 4a2 2 0 0 0 0 4c5.514 0 10 4.486 10 10a2 2 0 0 0 4 0c0-7.72-6.28-14-14-14m0 6a2 2 0 0 0 0 4c2.205 0 4 1.794 4 4a2 2 0 0 0 4 0c0-4.411-3.589-8-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RssOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 4.999a3.01 3.01 0 0 0-3.011 3l.005 9c0 2.209 1.793 4 4.002 4l9.003.001c1.655 0 3-1.346 3-3.001c.001-7.179-5.819-13-12.999-13m1.001 14A1.998 1.998 0 0 1 7 17a1.999 1.999 0 0 1 2.001-2.001A1.996 1.996 0 0 1 11 17a1.995 1.995 0 0 1-1.999 1.999m4.499 0a1.5 1.5 0 0 1-1.5-1.5c0-1.931-1.57-3.5-3.5-3.5a1.5 1.5 0 1 1 0-3c3.584 0 6.5 2.916 6.5 6.5a1.5 1.5 0 0 1-1.5 1.5m4 0a1.5 1.5 0 0 1-1.5-1.5c0-4.136-3.364-7.5-7.5-7.5a1.5 1.5 0 1 1 0-3c5.79 0 10.5 4.71 10.5 10.5a1.5 1.5 0 0 1-1.5 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Scissors(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.625 5.515c-1-1.522-2.915-1.67-4.397-.824l-.186.107l-.076-.003c-1.042 0-2.01.511-2.604 1.369l-.034.045c-.43.645-.723 1.236-1.005 1.809c-.255.516-.5 1.01-.824 1.483c-.325-.475-.57-.97-.826-1.486c-.283-.571-.575-1.162-1.004-1.806l-.033-.044a3.159 3.159 0 0 0-2.603-1.37a3.17 3.17 0 0 0-3.167 3.166a3.171 3.171 0 0 0 3.167 3.168c.775 0 1.515-.287 2.087-.791l.652 1.198c-1.621 1.876-2.979 4.054-3.019 4.121c-1.236 1.702.705 4.42.789 4.534a.498.498 0 0 0 .405.207a.511.511 0 0 0 .439-.261l3.112-5.717l3.113 5.717a.5.5 0 0 0 .407.26a.478.478 0 0 0 .437-.206c.083-.114 2.024-2.832.809-4.504l-.323-.521c-1.076-1.736-1.187-1.916-2.715-3.634l.651-1.195a3.16 3.16 0 0 0 2.088.791a3.171 3.171 0 0 0 3.167-3.168a3.13 3.13 0 0 0-.547-1.768c.472-.27.997-.123 1.456.095c.466.191.897-.377.584-.772M7 9a1 1 0 1 1 0-2a1 1 0 0 1 0 2m4.5 3.395a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1M16 9a1 1 0 1 1 0-2a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScissorsOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m18.124 5.27l.25.013c.163.022.319.02.468.077c.146.05.292.084.42.146c.254.131.471.262.631.422c.174.137.248.279.321.371l.108.146l-.166-.074c-.105-.043-.258-.139-.442-.16a1.648 1.648 0 0 0-.442-.047l-.15.002l-.318.061c-.103.01-.201.078-.295.113c-.186.088-.35.203-.461.34l-.057.076c.451.477.732 1.115.732 1.824a2.663 2.663 0 0 1-4.855 1.518l-1.657 2.625c1.712 1.92 3.22 4.348 3.22 4.348c1.037 1.429-.789 3.947-.789 3.947l-3.552-6.522l-3.551 6.522s-1.826-2.52-.789-3.947c0 0 1.507-2.428 3.22-4.348l-1.656-2.625a2.665 2.665 0 0 1-2.188 1.15A2.67 2.67 0 0 1 3.459 8.58a2.664 2.664 0 0 1 4.885-1.473c.955 1.428 1.739 3.156 2.748 4.334c1.008-1.178 1.792-2.906 2.746-4.336a2.67 2.67 0 0 1 2.221-1.191l.186.018l.326-.188l.383-.211c.132-.072.297-.107.449-.158c.224-.07.475-.105.721-.105m-2.069 4.376a1.067 1.067 0 1 0-.752-1.82l-.24.377a1.063 1.063 0 0 0 .992 1.443m-9.931 0a1.058 1.058 0 0 0 .989-1.443l-.237-.377a1.065 1.065 0 1 0-.752 1.82m4.965 3.762a.397.397 0 0 0 .396-.395a.394.394 0 1 0-.79 0c0 .217.176.395.394.395M18.124 3.27c-.448 0-.901.066-1.312.191l-.117.036c-.168.051-.426.126-.707.28l-.271.148a4.657 4.657 0 0 0-3.467 1.961l-.079.108c-.39.584-.741 1.192-1.082 1.784a30.161 30.161 0 0 0-1.081-1.784l-.078-.107a4.656 4.656 0 0 0-3.807-1.974a4.671 4.671 0 0 0-4.667 4.666a4.674 4.674 0 0 0 5.575 4.578a41.116 41.116 0 0 0-1.948 2.803c-1.714 2.467.392 5.619.835 6.229c.377.521.98.826 1.619.826l.129-.004a2 2 0 0 0 1.628-1.039l1.795-3.298l1.795 3.298c.328.604.943.995 1.628 1.039l.129.004c.639 0 1.241-.306 1.619-.826c.443-.61 2.549-3.764.834-6.229a41.484 41.484 0 0 0-1.947-2.803a4.675 4.675 0 0 0 5.576-4.578l-.004-.179a2 2 0 0 0 1.168-3.203l-.049-.067l-.039-.053a3.262 3.262 0 0 0-.572-.644a4.592 4.592 0 0 0-1.045-.705c-.275-.136-.488-.201-.628-.244l-.062-.02a3.112 3.112 0 0 0-.784-.16l-.054-.006l-.086-.01z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBag(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 4H7C5.346 4 4 5.346 4 7v11c0 1.654 1.346 3 3 3h10c1.654 0 3-1.346 3-3V7c0-1.654-1.346-3-3-3m1 14c0 .551-.448 1-1 1H7c-.552 0-1-.449-1-1v-7.28c.296.174.635.28 1 .28h1.5c0 1.93 1.57 3.5 3.5 3.5s3.5-1.57 3.5-3.5H17c.365 0 .704-.106 1-.279zm-8.5-7h5c0 1.378-1.121 2.5-2.5 2.5S9.5 12.378 9.5 11M18 9c0 .551-.448 1-1 1H7c-.552 0-1-.449-1-1V7c0-.551.448-1 1-1h10c.552 0 1 .449 1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingCart(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.756 5.345A1.003 1.003 0 0 0 20 5H6.181l-.195-1.164A1 1 0 0 0 5 3H2.75a1 1 0 1 0 0 2h1.403l1.86 11.164l.045.124l.054.151l.12.179l.095.112l.193.13l.112.065a.97.97 0 0 0 .367.075H18a1 1 0 1 0 0-2H7.847l-.166-1H19a1 1 0 0 0 .99-.858l1-7a1.002 1.002 0 0 0-.234-.797M18.847 7l-.285 2H15V7zM14 7v2h-3V7zm0 3v2h-3v-2zm-4-3v2H7l-.148.03L6.514 7zm-2.986 3H10v2H7.347zM15 12v-2h3.418l-.285 2z"/><circle cx="8.5" cy="19.5" r="1.5" fill="currentColor"/><circle cx="17.5" cy="19.5" r="1.5" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialAtCircular(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.844 7.5c-2.481 0-4.438 2.019-4.438 4.5s2.05 4.5 4.531 4.5c.908 0 1.799-.27 2.547-.778a.5.5 0 0 0 .139-.694a.495.495 0 0 0-.691-.132a3.466 3.466 0 0 1-1.965.604c-1.93 0-3.499-1.57-3.499-3.5s1.446-3.5 3.376-3.5s3.375 1.57 3.375 3.5v.25a.75.75 0 0 1-1.5 0V10.5c0-.276-.099-.5-.375-.5c-.205 0-.318.124-.396.301a1.864 1.864 0 0 0-1.01-.301a1.987 1.987 0 0 0-1.984 2a2.008 2.008 0 0 0 3.446 1.391c.319.369.664.609 1.192.609c.965 0 1.627-.785 1.627-1.75V12c0-2.481-1.894-4.5-4.375-4.5m.125 5.5c-.551 0-1-.449-1-1s.449-1 1-1s1 .449 1 1s-.449 1-1 1M12 21c-4.963 0-9-4.037-9-9s4.037-9 9-9s9 4.037 9 9s-4.037 9-9 9m0-16c-3.859 0-7 3.141-7 7s3.141 7 7 7s7-3.141 7-7s-3.141-7-7-7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialDribbble(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 3c-4.962 0-9 4.037-9 9s4.038 9 9 9s9-4.037 9-9s-4.038-9-9-9m6.962 8.275a12.343 12.343 0 0 0-5.205.262c-.18-.436-.383-.859-.59-1.283a12.355 12.355 0 0 0 3.713-3.262a6.978 6.978 0 0 1 2.082 4.283M16.13 6.361a11.398 11.398 0 0 1-3.401 3.009a23.175 23.175 0 0 0-2.807-4.056a6.957 6.957 0 0 1 6.208 1.047m-7.183-.65a22.227 22.227 0 0 1 2.892 4.117a11.394 11.394 0 0 1-6.717.899a7.011 7.011 0 0 1 3.825-5.016M5 12l.015-.294c.676.111 1.353.188 2.024.188a12.38 12.38 0 0 0 5.237-1.187c.182.373.365.744.525 1.127A12.382 12.382 0 0 0 6.7 16.56A6.968 6.968 0 0 1 5 12m2.393 5.257a11.61 11.61 0 0 1 5.764-4.487a22.154 22.154 0 0 1 1.292 5.779A6.94 6.94 0 0 1 12 19a6.96 6.96 0 0 1-4.607-1.743m8.014.852a23.207 23.207 0 0 0-1.293-5.632a11.402 11.402 0 0 1 4.871-.19a6.997 6.997 0 0 1-3.578 5.822"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialDribbbleCircular(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21c-4.962 0-9-4.037-9-9s4.038-9 9-9s9 4.037 9 9s-4.038 9-9 9m0-16c-3.86 0-7 3.141-7 7s3.14 7 7 7s7-3.141 7-7s-3.14-7-7-7m0 1.5c-3.033 0-5.5 2.468-5.5 5.5s2.467 5.5 5.5 5.5s5.5-2.468 5.5-5.5s-2.467-5.5-5.5-5.5m4.49 5.402a7.933 7.933 0 0 0-3.103.042a15.172 15.172 0 0 0-.417-.965a7.915 7.915 0 0 0 2.284-2.07a4.47 4.47 0 0 1 1.236 2.993m-2-3.646a6.918 6.918 0 0 1-1.975 1.832a14.639 14.639 0 0 0-1.692-2.412A4.448 4.448 0 0 1 12 7.5c.921 0 1.776.28 2.49.756m-4.641-.184c.687.758 1.278 1.59 1.776 2.473a6.928 6.928 0 0 1-3.998.437a4.493 4.493 0 0 1 2.222-2.91M7.5 12c.468.064.936.121 1.399.121a7.912 7.912 0 0 0 3.185-.683c.123.261.238.524.344.793A7.91 7.91 0 0 0 8.7 15.036A4.47 4.47 0 0 1 7.5 12m1.948 3.699a6.918 6.918 0 0 1 3.318-2.527a13.72 13.72 0 0 1 .596 3.095A4.458 4.458 0 0 1 12 16.5a4.466 4.466 0 0 1-2.552-.801m4.872.137c-.099-1-.296-1.983-.593-2.937a6.94 6.94 0 0 1 2.683.001a4.505 4.505 0 0 1-2.09 2.936"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialFacebook(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 10h3v3h-3v7h-3v-7H7v-3h3V8.745c0-1.189.374-2.691 1.118-3.512C11.862 4.41 12.791 4 13.904 4H16v3h-2.1c-.498 0-.9.402-.9.899z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialFacebookCircular(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.354 5.624C16.604 3.883 14.466 3 12 3c-2.489 0-4.633.884-6.373 2.625C3.884 7.366 3 9.512 3 12c0 2.465.883 4.603 2.624 6.354C7.365 20.11 9.51 21 12 21c2.467 0 4.605-.89 6.356-2.643C20.111 16.604 21 14.465 21 12c0-2.488-.89-4.634-2.646-6.376m-1.412 11.319c-1.137 1.139-2.436 1.788-3.942 1.985V14h2v-2h-2v-1.4a.6.6 0 0 1 .601-.6H15V8h-1.397c-.742 0-1.361.273-1.857.822c-.496.547-.746 1.215-.746 2.008V12H9v2h2v4.93c-1.522-.195-2.826-.845-3.957-1.984C5.668 15.562 5 13.944 5 12c0-1.966.667-3.588 2.042-4.96C8.412 5.667 10.034 5 12 5c1.945 0 3.562.668 4.945 2.043C18.328 8.415 19 10.037 19 12c0 1.941-.673 3.559-2.058 4.943"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialFlickr(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 16c-2.206 0-4-1.794-4-4s1.794-4 4-4s4 1.794 4 4s-1.794 4-4 4m0-6c-1.103 0-2 .897-2 2s.897 2 2 2s2-.897 2-2s-.897-2-2-2m9-2c-2.206 0-4 1.794-4 4s1.794 4 4 4s4-1.794 4-4s-1.794-4-4-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialFlickrCircular(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21c-2.489 0-4.635-.89-6.376-2.646C3.883 16.603 3 14.465 3 12c0-2.488.884-4.634 2.627-6.375C7.367 3.884 9.512 3 12 3c2.466 0 4.604.883 6.354 2.624C20.109 7.366 21 9.512 21 12c0 2.465-.89 4.604-2.644 6.357C16.605 20.11 14.467 21 12 21m0-16c-1.966 0-3.588.667-4.958 2.04C5.668 8.412 5 10.034 5 12c0 1.944.668 3.562 2.043 4.945C8.415 18.328 10.036 19 12 19c1.943 0 3.56-.673 4.941-2.057C18.327 15.559 19 13.941 19 12c0-1.963-.672-3.585-2.055-4.957C15.562 5.668 13.945 5 12 5m-3 9.5c-1.379 0-2.5-1.121-2.5-2.5S7.621 9.5 9 9.5s2.5 1.121 2.5 2.5s-1.121 2.5-2.5 2.5m0-4c-.827 0-1.5.673-1.5 1.5s.673 1.5 1.5 1.5s1.5-.673 1.5-1.5s-.673-1.5-1.5-1.5m6 4c-1.379 0-2.5-1.121-2.5-2.5s1.121-2.5 2.5-2.5s2.5 1.121 2.5 2.5s-1.121 2.5-2.5 2.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialGithub(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.435 12.973c.269 0 .492.133.686.396c.192.265.294.588.294.975c0 .385-.102.711-.294.973c-.193.265-.417.396-.686.396c-.278 0-.522-.131-.715-.396c-.192-.262-.294-.588-.294-.973c0-.387.102-.71.294-.975c.192-.264.436-.396.715-.396m3.44-3.559c.746.811 1.125 1.795 1.125 2.953c0 .748-.086 1.423-.259 2.023c-.175.597-.394 1.084-.654 1.459a4.268 4.268 0 0 1-.974.989a4.94 4.94 0 0 1-1.065.623a5.465 5.465 0 0 1-1.111.306a9 9 0 0 1-.943.123l-.685.014l-.547.015a17.567 17.567 0 0 1-1.524 0l-.547-.015l-.685-.014a8.966 8.966 0 0 1-.943-.123a5.28 5.28 0 0 1-1.111-.306a4.888 4.888 0 0 1-1.064-.623a4.253 4.253 0 0 1-.975-.989c-.261-.375-.479-.862-.654-1.459c-.173-.6-.259-1.275-.259-2.023c0-1.158.379-2.143 1.125-2.953c-.082-.041-.085-.447-.008-1.217a7.071 7.071 0 0 1 .495-2.132c.934.099 2.09.629 3.471 1.581c.466-.119 1.101-.183 1.917-.183c.852 0 1.491.064 1.918.184c.629-.425 1.23-.771 1.805-1.034c.584-.261 1.005-.416 1.269-.457l.396-.09c.27.649.434 1.36.496 2.132c.076.769.073 1.175-.009 1.216m-5.845 7.82c1.688 0 2.954-.202 3.821-.607c.855-.404 1.292-1.238 1.292-2.496c0-.73-.273-1.34-.822-1.828a1.845 1.845 0 0 0-.989-.486c-.375-.061-.949-.061-1.72 0c-.769.062-1.298.09-1.582.09c-.385 0-.8-.018-1.319-.059c-.52-.04-.928-.065-1.223-.078a3.727 3.727 0 0 0-.958.108a1.913 1.913 0 0 0-.853.425c-.521.469-.79 1.077-.79 1.828c0 1.258.426 2.092 1.28 2.496c.85.405 2.113.607 3.802.607zm-2.434-4.261c.268 0 .492.133.685.396c.192.265.294.588.294.975c0 .385-.102.711-.294.973c-.192.265-.417.396-.685.396c-.279 0-.522-.131-.716-.396c-.192-.262-.294-.588-.294-.973c0-.387.102-.71.294-.975c.193-.264.436-.396.716-.396"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialGithubCircular(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21c-4.963 0-9-4.038-9-9s4.037-9 9-9s9 4.038 9 9s-4.037 9-9 9m0-16c-3.859 0-7 3.14-7 7s3.141 7 7 7s7-3.14 7-7s-3.141-7-7-7m1.565 7.626c.171 0 .316.084.441.255c.124.169.187.378.187.625c0 .248-.062.457-.187.626c-.125.169-.271.254-.441.254c-.181 0-.337-.084-.461-.254c-.124-.169-.187-.378-.187-.626s.062-.456.187-.625c.125-.171.281-.255.461-.255m2.21-2.289c.482.522.725 1.155.725 1.898c0 .482-.057.915-.166 1.301a3.196 3.196 0 0 1-.42.939a2.717 2.717 0 0 1-.627.635a3.26 3.26 0 0 1-.685.401c-.208.085-.446.15-.716.196a5.221 5.221 0 0 1-.606.079l-.44.009l-.352.01l-.488.011l-.488-.011l-.352-.01l-.44-.009a5.168 5.168 0 0 1-.606-.079a3.272 3.272 0 0 1-.716-.196a3.189 3.189 0 0 1-.684-.401a2.74 2.74 0 0 1-.628-.635a3.196 3.196 0 0 1-.42-.939a4.78 4.78 0 0 1-.166-1.301c0-.743.242-1.376.725-1.898c-.053-.026-.056-.286-.008-.782a4.65 4.65 0 0 1 .319-1.37c.602.064 1.343.404 2.23 1.017c.3-.078.71-.118 1.233-.118c.549 0 .959.04 1.234.118a8.291 8.291 0 0 1 1.16-.666c.374-.168.644-.267.814-.293l.254-.058c.172.417.277.875.32 1.37c.05.496.047.756-.006.782m-3.754 5.027c1.083 0 1.899-.129 2.454-.39c.553-.26.833-.796.833-1.605c0-.469-.176-.861-.529-1.174a1.192 1.192 0 0 0-.638-.313c-.238-.039-.607-.039-1.104 0c-.495.04-.834.058-1.016.058c-.248 0-.517-.013-.851-.039l-.783-.049a2.408 2.408 0 0 0-.616.069a1.235 1.235 0 0 0-.55.273c-.336.3-.507.691-.507 1.174c0 .809.274 1.345.821 1.605c.547.261 1.361.39 2.444.39m-1.524-2.737c.17 0 .316.084.44.255c.124.169.187.378.187.625c0 .248-.062.457-.187.626c-.124.169-.271.254-.44.254c-.182 0-.337-.084-.462-.254c-.124-.169-.187-.378-.187-.626s.062-.456.187-.625c.125-.171.28-.255.462-.255"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialGooglePlus(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12.9 13.5l-.7-.5c-.2-.2-.5-.4-.5-.8s.3-.7.6-1c.8-.6 1.7-1.3 1.7-2.8c0-1.5-.9-2.3-1.4-2.7h1.2L15 5h-4.1c-1 0-2.4.2-3.5 1.1c-.8.7-1.2 1.7-1.2 2.6c0 1.5 1.2 3.1 3.3 3.1h.6c-.1.2-.2.4-.2.7c0 .6.3 1 .6 1.3c-.9.1-2.5.2-3.8.9c-1.2.7-1.5 1.7-1.5 2.4c0 1.5 1.4 2.8 4.2 2.8c3.4 0 5.2-1.9 5.2-3.7c0-1.3-.8-1.9-1.7-2.7m-2.5-2.2c-1.7 0-2.5-2.2-2.5-3.5c0-.5.1-1 .4-1.5c.3-.4.9-.7 1.4-.7c1.6 0 2.5 2.2 2.5 3.6c0 .4 0 1-.5 1.4c-.3.4-.9.7-1.3.7m0 7.9c-2.1 0-3.5-1-3.5-2.4s1.3-1.9 1.7-2c.8-.3 1.9-.3 2.1-.3h.5c1.5 1.1 2.1 1.6 2.1 2.6c0 1.2-1 2.1-2.9 2.1M17 12h-2v-1h2V9.1l1-.1v2h2v1h-2v2h-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialGooglePlusCircular(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12.8 13.1l-.4-.3c-.1-.1-.3-.2-.3-.5l.3-.6c.5-.4 1-.8 1-1.7s-.6-1.4-.8-1.6h.7L14 8h-2.4c-.6 0-1.4.1-2.1.6c-.5.4-.8 1-.8 1.6c0 .9.7 1.9 2 1.9h.4c-.1.1-.1.2-.1.4c0 .4.2.6.4.8c-.5 0-1.5.1-2.3.5c-.7.4-.9 1-.9 1.4c0 .9.8 1.7 2.5 1.7c2 0 3.1-1.1 3.1-2.2c0-.7-.5-1.1-1-1.6m-1.6-1.3c-1 0-1.5-1.3-1.5-2.1c.1-.4.1-.7.3-.9s.5-.4.8-.4c1 0 1.5 1.3 1.5 2.2c0 .2 0 .6-.3.9c-.1.1-.5.3-.8.3m.1 4.7c-1.3 0-2.1-.6-2.1-1.4c0-.8.8-1.1 1-1.2c.5-.2 1.1-.2 1.2-.2h.3c.9.6 1.3 1 1.3 1.6c0 .7-.6 1.2-1.7 1.2M15 12h-1v-1h1v-1h1v1h1v1h-1v1h-1zm-3 9c-5 0-9-4-9-9s4-9 9-9s9 4 9 9s-4 9-9 9m0-16c-3.9 0-7 3.1-7 7s3.1 7 7 7s7-3.1 7-7s-3.1-7-7-7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialInstagram(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 3H6C4.3 3 3 4.3 3 6v12c0 1.7 1.3 3 3 3h12c1.7 0 3-1.3 3-3V6c0-1.7-1.3-3-3-3m-6 6c1.7 0 3 1.3 3 3s-1.3 3-3 3s-3-1.3-3-3s1.3-3 3-3m3.8-2c0-.7.6-1.2 1.2-1.2s1.2.6 1.2 1.2s-.5 1.2-1.2 1.2s-1.2-.5-1.2-1.2M18 19H6c-.6 0-1-.4-1-1v-6h2c0 2.8 2.2 5 5 5s5-2.2 5-5h2v6c0 .6-.4 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialInstagramCircular(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 3c-5 0-9 4-9 9s4 9 9 9s9-4 9-9s-4-9-9-9m0 7c1.1 0 2 .9 2 2s-.9 2-2 2s-2-.9-2-2s.9-2 2-2m2.8-2c0-.7.6-1.2 1.2-1.2s1.2.6 1.2 1.2s-.5 1.2-1.2 1.2s-1.2-.5-1.2-1.2M12 19c-3.9 0-7-3.1-7-7h3c0 2.2 1.8 4 4 4s4-1.8 4-4h3c0 3.9-3.1 7-7 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialLastFm(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.942 16.182c2.374 0 3.558-.791 3.558-2.373c0-1.304-.749-2.132-2.254-2.49l-1.119-.235c-.637-.159-.951-.495-.951-1.009c0-.594.396-.889 1.186-.889c.869 0 1.323.334 1.363 1.006l1.717-.178c-.114-1.463-1.109-2.195-2.962-2.195c-2.019 0-3.026.832-3.026 2.495c0 1.182.654 1.935 1.958 2.251l1.188.236c.79.196 1.186.555 1.186 1.068c0 .631-.614.949-1.842.949c-1.498 0-2.489-.732-2.962-2.195l-.597-1.721c-.354-1.145-.796-1.947-1.334-2.401c-.53-.45-1.367-.683-2.519-.683c-1.069 0-2.007.396-2.815 1.188c-.811.791-1.217 1.838-1.217 3.142c0 1.223.383 2.203 1.156 2.935a3.864 3.864 0 0 0 2.756 1.099c1.069 0 1.918-.256 2.55-.77l-.53-1.485c-.554.556-1.211.833-1.96.833c-.63 0-1.175-.248-1.628-.744c-.455-.492-.686-1.137-.686-1.927c0-.989.247-1.708.743-2.163c.497-.455 1.056-.681 1.689-.681c.674 0 1.155.177 1.457.53c.296.357.56.912.797 1.662l.537 1.721c.632 2.014 2.154 3.024 4.561 3.024"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialLastFmCircular(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21c-2.489 0-4.635-.89-6.376-2.646C3.883 16.603 3 14.464 3 12c0-2.489.884-4.633 2.627-6.375C7.367 3.884 9.512 3 12 3c2.466 0 4.604.883 6.354 2.624C20.109 7.366 21 9.511 21 12c0 2.464-.89 4.604-2.644 6.357C16.605 20.111 14.467 21 12 21m0-16c-1.966 0-3.588.667-4.958 2.04C5.668 8.412 5 10.034 5 12c0 1.944.668 3.562 2.043 4.945C8.415 18.328 10.036 19 12 19c1.943 0 3.56-.673 4.941-2.056C18.327 15.559 19 13.942 19 12c0-1.963-.672-3.585-2.055-4.957C15.562 5.668 13.945 5 12 5m2.199 9.333c1.335 0 2-.444 2-1.333c0-.733-.422-1.199-1.267-1.4l-.632-.133c-.354-.089-.534-.277-.534-.566c0-.334.224-.5.666-.5c.49 0 .746.188.767.565l.967-.1c-.063-.822-.622-1.233-1.666-1.233c-1.134 0-1.699.467-1.699 1.401c0 .665.365 1.088 1.099 1.267l.668.133c.443.11.666.312.666.601c0 .354-.345.532-1.034.532c-.844 0-1.398-.411-1.666-1.233l-.334-.967c-.199-.644-.449-1.095-.75-1.35c-.3-.255-.771-.384-1.416-.384c-.601 0-1.128.223-1.584.667c-.456.445-.683 1.033-.683 1.767c0 .688.216 1.239.649 1.649c.435.413.95.617 1.55.617c.602 0 1.078-.144 1.434-.433l-.299-.834a1.505 1.505 0 0 1-1.101.468c-.354 0-.662-.14-.916-.417c-.257-.277-.385-.64-.385-1.084c0-.556.139-.961.417-1.217s.594-.383.951-.383c.379 0 .648.1.816.299c.167.201.315.512.45.935l.3.967c.356 1.133 1.212 1.699 2.566 1.699"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialLinkedin(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 19H5V9h3zm11 0h-3v-5.342c0-1.392-.496-2.085-1.479-2.085c-.779 0-1.273.388-1.521 1.165V19h-3s.04-9 0-10h2.368l.183 2h.062c.615-1 1.598-1.678 2.946-1.678c1.025 0 1.854.285 2.487 1.001c.637.717.954 1.679.954 3.03z"/><ellipse cx="6.5" cy="6.5" fill="currentColor" rx="1.55" ry="1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialLinkedinCircular(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.033 15.3h-1.6v-5.199h1.6zm-.8-5.866c-.577 0-.866-.267-.866-.8a.74.74 0 0 1 .25-.567a.868.868 0 0 1 .616-.233c.577 0 .866.268.866.801s-.288.799-.866.799m6.734 5.866h-1.633v-2.9c0-.755-.268-1.133-.801-1.133c-.422 0-.699.211-.834.633c-.043.067-.066.201-.066.4v3H11v-3.533c0-.8-.012-1.355-.033-1.666h1.4l.1.699c.367-.556.9-.833 1.633-.833c.557 0 1.006.194 1.35.583c.346.389.518.95.518 1.684V15.3zM12 21c-4.963 0-9-4.037-9-9s4.037-9 9-9s9 4.037 9 9s-4.037 9-9 9m0-16c-3.859 0-7 3.141-7 7s3.141 7 7 7s7-3.141 7-7s-3.141-7-7-7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialPinterest(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.486 4.771c-4.23 0-6.363 3.033-6.363 5.562c0 1.533.581 2.894 1.823 3.401c.205.084.387.004.446-.221l.182-.717c.061-.221.037-.3-.127-.495c-.359-.422-.588-.972-.588-1.747c0-2.25 1.683-4.264 4.384-4.264c2.392 0 3.706 1.463 3.706 3.412c0 2.568-1.137 4.734-2.824 4.734c-.932 0-1.629-.77-1.405-1.715c.268-1.13.786-2.347.786-3.16c0-.729-.392-1.336-1.2-1.336c-.952 0-1.718.984-1.718 2.304c0 .841.286 1.409.286 1.409L8.728 16.79c-.34 1.44-.051 3.206-.025 3.385c.013.104.149.131.21.051c.088-.115 1.223-1.517 1.607-2.915c.111-.396.627-2.445.627-2.445c.311.589 1.213 1.108 2.175 1.108c2.863 0 4.804-2.608 4.804-6.103c-.003-2.64-2.24-5.1-5.64-5.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialPinterestCircular(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21c-4.963 0-9-4.037-9-9s4.037-9 9-9s9 4.037 9 9s-4.037 9-9 9m0-16c-3.859 0-7 3.141-7 7s3.141 7 7 7s7-3.141 7-7s-3.141-7-7-7m.335 3c-2.468 0-3.712 1.77-3.712 3.244c0 .895.338 1.688 1.063 1.984c.119.049.226.002.261-.129l.105-.418c.035-.129.021-.175-.074-.289c-.209-.246-.344-.566-.344-1.019c0-1.312.982-2.487 2.558-2.487c1.396 0 2.161.853 2.161 1.99c0 1.498-.662 2.762-1.646 2.762c-.543 0-.95-.449-.82-1.001c.156-.658.459-1.368.459-1.843c0-.426-.229-.779-.7-.779c-.556 0-1.002.574-1.002 1.344c0 .49.166.822.166.822l-.669 2.83c-.198.84-.029 1.87-.015 1.974c.008.062.087.077.123.03c.052-.067.713-.885.938-1.7c.064-.23.366-1.427.366-1.427c.18.344.707.646 1.268.646c1.67 0 2.803-1.521 2.803-3.56C15.623 9.436 14.318 8 12.335 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialSkype(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.668 13.312a9.16 9.16 0 0 0 .089-1.251c0-2.437-.846-4.544-2.513-6.263c-1.685-1.735-3.755-2.615-6.152-2.615a8.21 8.21 0 0 0-.801.038a5.716 5.716 0 0 0-2.426-.526c-1.609 0-3.053.61-4.174 1.765C3.586 5.595 3 7.045 3 8.652c0 .832.156 1.619.466 2.348a8.15 8.15 0 0 0-.07 1.062c0 2.438.853 4.547 2.532 6.267c1.693 1.732 3.768 2.61 6.164 2.61c.254 0 .547-.02.896-.06c.69.283 1.409.426 2.146.426c1.588 0 3.025-.614 4.157-1.777A5.912 5.912 0 0 0 21 15.317c0-.677-.111-1.349-.332-2.005M15.5 15.348c-.284.427-.729.781-1.339 1.065c-.609.243-1.31.365-2.1.365c-.954 0-1.756-.173-2.404-.519a3.068 3.068 0 0 1-1.096-1.003c-.284-.447-.428-.862-.428-1.248c0-.243.092-.457.274-.639a.945.945 0 0 1 .7-.274c.203 0 .406.072.609.213c.162.162.283.366.364.609c.021.02.153.253.396.7c.102.141.284.283.547.425c.245.122.568.183.975.183c.548 0 1.005-.121 1.37-.364c.324-.224.486-.507.486-.853c0-.284-.092-.498-.274-.639a1.856 1.856 0 0 0-.699-.427l-.319-.075l-.457-.123a3.733 3.733 0 0 0-.441-.106c-.689-.141-1.277-.324-1.766-.548c-.426-.162-.811-.445-1.156-.851c-.283-.386-.426-.843-.426-1.37c0-.528.152-.994.457-1.4c.304-.406.74-.711 1.308-.913c.569-.224 1.219-.334 1.949-.334c.548 0 1.066.07 1.552.213c.468.162.843.355 1.127.579c.263.243.477.486.639.729c.142.324.214.589.214.791a.904.904 0 0 1-.275.669a.956.956 0 0 1-.699.306a.929.929 0 0 1-.578-.184c-.102-.101-.233-.283-.396-.547a2.072 2.072 0 0 0-.577-.762c-.243-.182-.619-.273-1.127-.273c-.507 0-.892.102-1.156.305c-.284.184-.426.396-.426.639c0 .162.05.296.152.396c.102.143.232.255.396.336c.102.06.274.132.518.212l.41.106c.213.053.391.087.533.107c.283.061.771.192 1.461.396c.406.141.791.323 1.156.548c.346.242.599.517.762.82c.162.365.242.771.242 1.217a2.756 2.756 0 0 1-.458 1.523"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialSkypeOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.865 5c.751 0 1.44.202 2.069.609c.324-.08.711-.121 1.157-.121c1.846 0 3.418.67 4.717 2.008c1.299 1.339 1.948 2.962 1.948 4.87c0 .466-.051.953-.152 1.461c.264.589.396 1.187.396 1.794c0 1.097-.38 2.036-1.142 2.815c-.761.781-1.668 1.173-2.725 1.173a3.714 3.714 0 0 1-1.825-.488c-.527.081-.933.122-1.217.122c-1.847 0-3.425-.67-4.733-2.009s-1.963-2.962-1.963-4.868c0-.447.051-.902.152-1.37A3.897 3.897 0 0 1 5 8.957c0-1.096.376-2.029 1.126-2.799C6.876 5.386 7.79 5 8.865 5M12 15.53c-.406 0-.729-.061-.975-.183c-.263-.142-.445-.284-.547-.425c-.243-.447-.376-.681-.396-.7a1.53 1.53 0 0 0-.364-.609a1.072 1.072 0 0 0-.61-.213a.947.947 0 0 0-.7.274a.874.874 0 0 0-.274.639c0 .386.144.801.428 1.248c.263.405.629.741 1.096 1.003c.648.346 1.45.519 2.404.519c.79 0 1.49-.122 2.1-.365c.609-.284 1.055-.639 1.339-1.065c.304-.467.456-.975.456-1.522c0-.445-.08-.852-.242-1.217c-.163-.304-.416-.578-.762-.82a5.854 5.854 0 0 0-1.156-.548a22.257 22.257 0 0 0-1.461-.396a5.869 5.869 0 0 1-.533-.107l-.41-.106a2.897 2.897 0 0 1-.518-.212a1.053 1.053 0 0 1-.396-.336a.527.527 0 0 1-.152-.396c0-.243.142-.455.426-.639c.265-.203.649-.305 1.156-.305c.508 0 .884.092 1.127.273c.263.225.456.478.577.762c.163.264.295.446.396.547a.929.929 0 0 0 .578.184a.954.954 0 0 0 .699-.306a.9.9 0 0 0 .275-.669c0-.202-.072-.467-.214-.791a4.128 4.128 0 0 0-.639-.729c-.284-.224-.659-.417-1.127-.579a5.47 5.47 0 0 0-1.552-.213c-.73 0-1.38.11-1.948.335c-.566.2-1.003.505-1.307.911a2.267 2.267 0 0 0-.457 1.4c0 .527.143.984.426 1.37c.346.405.73.688 1.156.851c.488.224 1.076.407 1.766.548c.121.021.27.056.441.106l.457.123l.319.075c.284.103.517.243.699.427c.183.141.274.354.274.639c0 .346-.162.629-.486.853c-.364.243-.821.364-1.369.364M8.865 3c-1.609 0-3.053.61-4.174 1.765C3.586 5.899 3 7.35 3 8.957c0 .832.156 1.619.466 2.348a8.15 8.15 0 0 0-.07 1.062c0 2.438.853 4.547 2.532 6.267c1.693 1.732 3.768 2.61 6.164 2.61c.254 0 .547-.02.896-.06c.69.283 1.409.426 2.146.426c1.588 0 3.025-.614 4.157-1.777A5.912 5.912 0 0 0 21 15.622c0-.677-.111-1.349-.332-2.004a9.16 9.16 0 0 0 .089-1.251c0-2.437-.846-4.544-2.513-6.263c-1.685-1.735-3.755-2.615-6.152-2.615a8.21 8.21 0 0 0-.801.038A5.717 5.717 0 0 0 8.865 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialTumbler(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.527 17.921v-2.066c-.669.448-1.32.67-1.952.67c-.298 0-.631-.094-1.004-.277c-.223-.151-.354-.317-.393-.503c-.11-.224-.178-.708-.178-1.454V11h3V9h-3V5.644h-1.772c-.149.782-.298 1.338-.448 1.673c-.184.41-.482.782-.891 1.116a4 4 0 0 1-1.285.725V11H9v4.521c0 .52.073.964.223 1.337c.111.298.334.595.671.893c.259.262.631.484 1.115.67c.595.15 1.114.223 1.562.223c.52 0 1.004-.056 1.45-.167a5.75 5.75 0 0 0 1.506-.556"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialTumblerCircular(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.377 15.59v-1.234c-.399.268-.788.4-1.166.4c-.178 0-.377-.057-.6-.166c-.134-.09-.211-.189-.234-.301c-.066-.133-.1-.422-.1-.867v-1.966h1.834v-1.233h-1.834V8.256h-1.066c-.089.467-.178.8-.267 1c-.11.244-.288.467-.533.666c-.245.201-.5.345-.767.434v1.101h.833v2.7c0 .311.044.576.134.799c.066.178.199.355.4.533c.154.156.377.289.666.4c.355.09.666.133.934.133c.311 0 .6-.033.866-.1c.312-.067.612-.178.9-.332"/><path fill="currentColor" d="M12 21c-4.963 0-9-4.037-9-9s4.037-9 9-9s9 4.037 9 9s-4.037 9-9 9m0-16c-3.859 0-7 3.141-7 7s3.141 7 7 7s7-3.141 7-7s-3.141-7-7-7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialTwitter(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.89 7.012c.808-.496 1.343-1.173 1.605-2.034a8.68 8.68 0 0 1-2.351.861c-.703-.756-1.593-1.14-2.66-1.14c-1.043 0-1.924.366-2.643 1.078a3.56 3.56 0 0 0-1.076 2.605c0 .309.039.585.117.819C8.806 9.096 6.26 7.82 4.254 5.364c-.34.601-.51 1.213-.51 1.846c0 1.301.549 2.332 1.645 3.089c-.625-.053-1.176-.211-1.645-.47c0 .929.273 1.705.82 2.388a3.623 3.623 0 0 0 2.115 1.291c-.312.08-.641.118-.979.118c-.312 0-.533-.026-.664-.083c.23.757.664 1.371 1.291 1.841a3.652 3.652 0 0 0 2.152.743c-1.332 1.045-2.855 1.562-4.578 1.562c-.422 0-.721-.006-.902-.038C4.696 18.753 6.585 19.3 8.675 19.3c2.139 0 4.029-.542 5.674-1.626c1.645-1.078 2.859-2.408 3.639-3.974a10.77 10.77 0 0 0 1.172-4.892V8.34A7.788 7.788 0 0 0 21 6.419a8.142 8.142 0 0 1-2.11.593"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialTwitterCircular(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.279 10.283c.358-.221.597-.521.713-.904a3.854 3.854 0 0 1-1.045.383a1.55 1.55 0 0 0-1.182-.507c-.464 0-.855.163-1.175.479a1.585 1.585 0 0 0-.478 1.158c0 .137.017.26.052.364c-1.368-.048-2.499-.614-3.391-1.706a1.644 1.644 0 0 0-.227.82c0 .578.244 1.036.73 1.373a1.81 1.81 0 0 1-.73-.209c0 .413.121.758.365 1.062c.243.3.557.492.939.573a1.734 1.734 0 0 1-.435.053a.826.826 0 0 1-.296-.037c.104.337.296.609.574.818c.277.21.597.32.957.33a3.196 3.196 0 0 1-2.035.694c-.188 0-.32-.002-.4-.017a4.534 4.534 0 0 0 2.521.733c.951 0 1.792-.241 2.522-.723c.73-.479 1.271-1.07 1.617-1.767a4.793 4.793 0 0 0 .521-2.174v-.209c.336-.253.609-.538.818-.854a3.452 3.452 0 0 1-.935.267M12 21c-2.49 0-4.635-.89-6.376-2.646C3.883 16.603 3 14.465 3 12c0-2.488.884-4.634 2.627-6.375C7.367 3.884 9.511 3 12 3c2.466 0 4.604.883 6.354 2.624C20.11 7.366 21 9.512 21 12c0 2.465-.889 4.604-2.644 6.357C16.605 20.11 14.467 21 12 21m0-16c-1.966 0-3.588.667-4.958 2.04C5.667 8.412 5 10.034 5 12c0 1.944.668 3.562 2.043 4.945C8.415 18.328 10.036 19 12 19c1.943 0 3.56-.673 4.942-2.057C18.327 15.559 19 13.941 19 12c0-1.963-.672-3.585-2.055-4.957C15.562 5.668 13.945 5 12 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialVimeo(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.92 8.776c-.329 1.929-1.211 3.758-2.649 5.48c-1.436 1.725-2.71 2.957-3.819 3.695c-.699.331-1.293.34-1.786.034c-.493-.31-.883-.751-1.169-1.325c-.165-.328-.565-1.569-1.202-3.728c-.636-2.155-1.017-3.315-1.139-3.479c-.083-.163-.288-.184-.616-.061c-.33.122-.555.226-.678.309a2.083 2.083 0 0 0-.308.245L5 9.209l.616-.74a13.325 13.325 0 0 1 1.724-1.54c.7-.534 1.314-.862 1.848-.987c.371-.08.679-.028.924.156c.247.184.452.484.616.894c.165.409.289.811.369 1.199c.083.392.165.854.248 1.387c.081.534.164.945.246 1.232c.451 1.93.821 2.896 1.109 2.896c.41 0 1.067-.863 1.971-2.59c.41-.779.432-1.426.062-1.941c-.369-.512-.943-.522-1.724-.029a3.608 3.608 0 0 1 1.046-2.031c1.026-1.109 2.157-1.521 3.388-1.234c1.273.247 1.765 1.213 1.477 2.895"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialVimeoCircular(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.463 9.141c-.512 0-.988.238-1.43.715c-.311.312-.5.678-.566 1.101c.207-.131.387-.196.541-.196c.159 0 .29.07.393.212c.199.278.188.629-.033 1.051c-.489.934-.846 1.4-1.066 1.4c-.156 0-.356-.522-.602-1.567a6.23 6.23 0 0 1-.133-.667a14.497 14.497 0 0 0-.133-.75c-.045-.211-.111-.428-.2-.649s-.2-.384-.333-.483a.536.536 0 0 0-.327-.104l-.173.021c-.289.067-.623.245-1 .534a7.082 7.082 0 0 0-.934.833l-.334.4l.301.399l.166-.133c.066-.045.189-.101.367-.167l.191-.043c.069 0 .116.025.143.076c.066.089.271.717.615 1.884c.346 1.167.562 1.839.65 2.017c.156.311.367.55.633.717c.13.08.271.121.427.121c.165 0 .346-.047.54-.139c.601-.399 1.289-1.066 2.067-2c.778-.933 1.255-1.922 1.433-2.966c.156-.911-.11-1.434-.799-1.567a1.63 1.63 0 0 0-.404-.05M12 21c-4.963 0-9-4.037-9-9s4.037-9 9-9s9 4.037 9 9s-4.037 9-9 9m0-16c-3.859 0-7 3.141-7 7s3.141 7 7 7s7-3.141 7-7s-3.141-7-7-7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialYoutube(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.8 8.6c-.2-1.5-.4-2.6-1-3C21.2 5.1 16 5 12 5s-9.2.1-9.8.6c-.6.4-.8 1.5-1 3S1 11 1 12s0 1.9.2 3.4s.4 2.6 1 3c.6.5 5.8.6 9.8.6c4 0 9.2-.1 9.8-.6c.6-.4.8-1.5 1-3S23 13 23 12s0-1.9-.2-3.4m-12.8 7V8.4l6 3.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialYoutubeCircular(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.48 13.14h.73v3.61h.7v-3.61h.73v-.61H8.48zM12.17 16c-.12.14-.53.42-.53.02v-2.39h-.62v2.61c0 .79.79.58 1.16.17v.34h.62v-3.12h-.62V16zm2.31-2.39c-.36 0-.59.27-.59.27v-1.36h-.63v4.23h.63v-.24s.21.28.59.28c.33 0 .58-.29.58-.69v-1.73c0-.47-.22-.76-.58-.76m-.07 2.41c0 .23-.16.34-.37.25a.479.479 0 0 1-.15-.11v-1.94c.04-.04.09-.07.13-.1c.22-.11.39.06.39.29zm2.31-.16c0 .24-.13.4-.28.41c-.16.01-.32-.12-.32-.41v-.59h1.19v-.8a.91.91 0 0 0-.26-.66a.919.919 0 0 0-.63-.24c-.22 0-.45.07-.63.21c-.19.15-.31.38-.31.69v1.4c0 .28.09.5.23.66c.17.18.4.28.64.29c.29.01.6-.11.78-.36c.11-.15.18-.35.18-.59v-.16h-.59zm-.6-1.39c0-.17.1-.37.29-.37s.31.18.31.37v.32h-.6z"/><path fill="currentColor" d="M12.97 3a9 9 0 1 0 .001 18.001A9 9 0 0 0 12.97 3m1.58 3.37h.8v2.68c0 .17.08.17.11.17c.12 0 .3-.13.39-.22V6.36h.8V9.9h-.8v-.31c-.11.1-.22.18-.33.24c-.15.08-.29.11-.43.11a.52.52 0 0 1-.41-.17c-.09-.11-.13-.28-.13-.49zM12 7.3c0-.55.45-1 1-1s1 .45 1 1V9c0 .55-.45 1-1 1s-1-.45-1-1zM9.92 5.15l.48 1.76l.49-1.76h.91l-.94 2.78V9.9h-.89V7.93l-.96-2.78zm7.9 12.54c-.51.5-4.83.51-4.83.51s-4.31-.01-4.83-.51c-.51-.5-.51-2.99-.51-3.01c0-.01 0-2.5.51-3.01c.51-.5 4.83-.51 4.83-.51s4.31.01 4.83.51c.51.5.52 2.99.52 3.01c0 0 0 2.5-.52 3.01"/><path fill="currentColor" d="M12.98 9.35c.17 0 .25-.1.26-.26v-1.9c0-.13-.13-.24-.25-.24s-.25.1-.25.24v1.9c0 .15.07.25.24.26"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortAlphabetically(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.895 16.553l-4-8c-.339-.678-1.45-.678-1.789 0l-4 8a1 1 0 0 0 1.789.895L3.618 16h4.764l.724 1.447a.998.998 0 0 0 1.341.448c.494-.248.695-.848.448-1.342M4.618 14L6 11.236L7.382 14zM22 18h-6a1.001 1.001 0 0 1-.8-1.6L20 10h-4a1 1 0 0 1 0-2h6a1.001 1.001 0 0 1 .8 1.6L18 16h4a1 1 0 0 1 0 2m-8-4h-2a1 1 0 0 1 0-2h2a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortAlphabeticallyOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.618 14h2.764L7 11.236zM21 14l2.4-3.2a2.986 2.986 0 0 0 .284-3.142A2.984 2.984 0 0 0 21 6h-6c-1.654 0-3 1.346-3 3c0 .77.295 1.469.774 2H12.5c-.368 0-.708.107-1.005.281L9.684 7.658C9.186 6.663 8.157 6.044 7 6.044s-2.186.619-2.684 1.614l-4 8c-.358.717-.416 1.53-.163 2.291s.788 1.376 1.504 1.735c.414.207.879.316 1.342.316a2.988 2.988 0 0 0 2.684-1.657L5.854 18h2.291l.171.342A2.984 2.984 0 0 0 11 20c.464 0 .928-.109 1.342-.316a3.01 3.01 0 0 0 .652-.458c.54.488 1.246.774 2.006.774h6c1.654 0 3-1.346 3-3s-1.346-3-3-3m-9.553 3.895a1.003 1.003 0 0 1-1.342-.448L9.382 16H4.618l-.724 1.447a1 1 0 1 1-1.788-.895l4-8c.169-.338.532-.508.894-.508s.725.169.895.508l4 8a1.002 1.002 0 0 1-.448 1.343M12.5 14a1 1 0 0 1 0-2h1a1 1 0 0 1 0 2zm8.5 4h-6a1.001 1.001 0 0 1-.8-1.6L19 10h-4a1 1 0 0 1 0-2h6a1.001 1.001 0 0 1 .8 1.6L17 16h4a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortNumerically(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 18a1 1 0 0 1-1-1v-6.382l-.553.276a1 1 0 1 1-.894-1.789l2-1A.998.998 0 0 1 5 9v8a1 1 0 0 1-1 1m9 0H8a1.002 1.002 0 0 1-.707-1.707l4-4c.254-.254.394-.591.394-.95c0-.358-.14-.695-.394-.949c-.508-.508-1.39-.508-1.9.001a1.33 1.33 0 0 0-.393.948a1 1 0 0 1-2 0c0-.894.348-1.733.98-2.364c1.265-1.263 3.464-1.263 4.727.001c.632.631.979 1.471.979 2.363c0 .893-.348 1.733-.979 2.364L10.414 16H13a1 1 0 0 1 0 2m7.955-5.623a2.725 2.725 0 0 0 .545-1.627A2.753 2.753 0 0 0 18.75 8a2.739 2.739 0 0 0-2.44 1.484a1 1 0 1 0 1.776.92a.75.75 0 1 1 .664 1.096a1 1 0 0 0 0 2c.689 0 1.25.561 1.25 1.25S19.439 16 18.75 16s-1.25-.561-1.25-1.25a1 1 0 0 0-2 0c0 1.792 1.458 3.25 3.25 3.25S22 16.542 22 14.75a3.23 3.23 0 0 0-1.045-2.373"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortNumericallyOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23.292 12.134c.138-.445.208-.91.208-1.384A4.756 4.756 0 0 0 18.75 6c-1.396 0-2.685.61-3.573 1.632l-.056-.067c-.973-.974-2.349-1.533-3.776-1.533c-1.422 0-2.794.556-3.77 1.525a3.398 3.398 0 0 0-1.122-1.108A3.054 3.054 0 0 0 4.84 6c-.482 0-.955.109-1.369.316l-1.406.747C.623 7.784.014 9.589.752 11.065c.272.543.714.982 1.248 1.272V17c0 1.654 1.346 3 3 3c.766 0 1.458-.297 1.989-.771c.54.487 1.25.771 2.011.771h5c.778 0 1.479-.305 2.01-.795c.796.5 1.731.795 2.74.795A5.256 5.256 0 0 0 24 14.75a5.23 5.23 0 0 0-.708-2.616M6 17a1 1 0 0 1-2 0v-6.382a1.052 1.052 0 0 1-.471.106c-.401 0-.813-.203-.988-.553a1.007 1.007 0 0 1 .463-1.342l1.361-.724c.141-.07.307-.105.475-.105c.199 0 .4.05.561.149c.294.183.599.504.599.851zm8 1H9a1.002 1.002 0 0 1-.707-1.707l4-4c.254-.254.394-.591.394-.95c0-.358-.14-.695-.394-.949s-.601-.381-.949-.381s-.696.127-.952.382a1.333 1.333 0 0 0-.392.948a1 1 0 0 1-2 0c0-.894.348-1.733.98-2.364c.632-.631 1.498-.947 2.364-.947s1.731.316 2.363.948c.632.631.979 1.471.979 2.363c0 .893-.348 1.733-.979 2.364L11.414 16H14a1 1 0 0 1 0 2m4.75 0a3.254 3.254 0 0 1-3.25-3.25a1 1 0 0 1 2 0c0 .689.561 1.25 1.25 1.25S20 15.439 20 14.75s-.561-1.25-1.25-1.25a1 1 0 0 1 0-2a.75.75 0 1 0-.665-1.096a1 1 0 0 1-1.776-.92A2.741 2.741 0 0 1 18.75 8a2.753 2.753 0 0 1 2.75 2.75c0 .611-.207 1.17-.545 1.627A3.23 3.23 0 0 1 22 14.75A3.254 3.254 0 0 1 18.75 18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spanner(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.285 7.119a.505.505 0 0 0-.354-.344a.493.493 0 0 0-.477.126l-2.616 2.557l-1.914-.383l-.381-1.907l2.645-2.585a.5.5 0 0 0-.199-.835A4.956 4.956 0 0 0 15.5 3.5c-2.757 0-5 2.243-5 5c0 .323.038.65.118 1.01c-.562.463-1.096.862-1.701 1.314c-.865.646-1.845 1.377-3.182 2.506A3.557 3.557 0 0 0 4.5 16c0 1.93 1.57 3.5 3.5 3.5a3.483 3.483 0 0 0 2.662-1.25a58.432 58.432 0 0 0 2.544-3.209c.442-.591.832-1.111 1.283-1.66c.36.081.688.119 1.011.119c2.757 0 5-2.243 5-5a4.85 4.85 0 0 0-.215-1.381M8 17a1 1 0 1 1 0-2a1 1 0 1 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpannerOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="8" cy="16" r="1" fill="currentColor"/><path fill="currentColor" d="M20.733 4.657a.997.997 0 0 0-1.399.009a.997.997 0 0 0 .01-1.4A6.47 6.47 0 0 0 15.5 2A6.5 6.5 0 0 0 9 8.5l.031.379c-.337.239-2.893 2.147-4.258 3.301C3.638 13.17 3 14.555 3 16c0 2.757 2.243 5 5 5c1.465 0 2.854-.65 3.811-1.784c1.173-1.375 3.08-3.923 3.317-4.229L15.5 15A6.5 6.5 0 0 0 22 8.5c0-1.44-.474-2.766-1.267-3.843M8 19a3 3 0 0 1-3-3c0-.92.423-1.732 1.064-2.292c2.368-2.002 3.617-2.748 5.115-4.015A4.475 4.475 0 0 1 11 8.5A4.5 4.5 0 0 1 15.5 4c.47 0 .914.092 1.339.226L14 7l.5 2.5l2.5.5l2.805-2.741c.115.396.195.807.195 1.241a4.5 4.5 0 0 1-4.5 4.5c-.416 0-.811-.074-1.193-.18c-1.267 1.498-2.013 2.748-4.024 5.105A2.98 2.98 0 0 1 8 19M19.384 6.271l-2.705 2.645l-1.329-.266l-.263-1.314l2.726-2.663c.651.393 1.19.939 1.571 1.598"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spiral(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 11.8c1-.4.7-1.8 0-2.4c-1.1-.9-2.7-.4-3.4.8c-1.5 2.4.9 5 3.4 4.9c2.7-.2 4.3-2.9 3.7-5.4c-.7-3-3.9-4.5-6.7-3.6c-2.6.8-4.2 3.5-4 6.2c.3 3 2.6 5.4 5.5 5.9c2.8.5 5.7-.8 7.2-3.2c.7-1.1 1.2-2.4 1.2-3.8c0-.5.5-1 1.1-.9c.8 0 1 .8.9 1.4c-.4 4.7-4.5 8.6-9.3 8.6c-5.9 0-10.5-6.2-8-11.8C6.1 3.1 13.9 2 16.9 7.3c1.5 2.5 1.2 5.8-.9 7.9c-2 2-5.3 2.4-7.7.7c-2.2-1.6-2.9-4.9-1.1-7.2c1.7-2.3 5.5-2.4 7 .2c1.1 1.9 0 5.2-2.5 4.9c-1.6 0-3-1.7-2.1-3.2c.6-.9 1.9-.6 2.3.1c.2.8.1 1.1.1 1.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Star(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9.362 9.158l-5.268.584c-.19.023-.358.15-.421.343s0 .394.14.521c1.566 1.429 3.919 3.569 3.919 3.569c-.002 0-.646 3.113-1.074 5.19a.496.496 0 0 0 .734.534c1.844-1.048 4.606-2.624 4.606-2.624l4.604 2.625c.168.092.378.09.541-.029a.5.5 0 0 0 .195-.505l-1.071-5.191l3.919-3.566a.499.499 0 0 0-.28-.865c-2.108-.236-5.269-.586-5.269-.586l-2.183-4.83a.499.499 0 0 0-.909 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarFullOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m3.1 11.3l3.6 3.3l-1 4.6c-.1.6.1 1.2.6 1.5c.2.2.5.3.8.3c.2 0 .4 0 .6-.1c0 0 .1 0 .1-.1l4.1-2.3l4.1 2.3s.1 0 .1.1c.5.2 1.1.2 1.5-.1c.5-.3.7-.9.6-1.5l-1-4.6c.4-.3 1-.9 1.6-1.5l1.9-1.7l.1-.1c.4-.4.5-1 .3-1.5s-.6-.9-1.2-1h-.1l-4.7-.5l-1.9-4.3s0-.1-.1-.1c-.1-.7-.6-1-1.1-1c-.5 0-1 .3-1.3.8c0 0 0 .1-.1.1L8.7 8.2L4 8.7h-.1c-.5.1-1 .5-1.2 1c-.1.6 0 1.2.4 1.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarHalf(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.5 4.3c-.9 1.9-2.2 4.8-2.2 4.8s-3.1.4-5.2.6c-.2 0-.4.2-.4.3c-.1.2 0 .4.1.5c1.6 1.4 3.9 3.6 3.9 3.6s-.6 3.1-1.1 5.2c0 .2 0 .4.2.5c.2.2.4.2.6.1c1.8-1 4.6-2.6 4.6-2.6V4c-.2 0-.4.2-.5.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarHalfOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m3.1 11.3l3.6 3.3l-1 4.6c-.1.6.1 1.2.6 1.5c.2.2.5.3.8.3c.2 0 .4 0 .6-.1c0 0 .1 0 .1-.1l4.1-2.3l4.1 2.3s.1 0 .1.1c.5.2 1.1.2 1.5-.1c.5-.3.7-.9.6-1.5l-1-4.6c.4-.3 1-.9 1.6-1.5l1.9-1.7l.1-.1c.4-.4.5-1 .3-1.5s-.6-.9-1.2-1h-.1l-4.7-.5l-1.9-4.3s0-.1-.1-.1c-.1-.7-.6-1-1.1-1c-.5 0-1 .3-1.3.8c0 0 0 .1-.1.1L8.7 8.2L4 8.7h-.1c-.5.1-1 .5-1.2 1c-.1.6 0 1.2.4 1.6m8.9 5V5.8l1.7 3.8c.1.3.5.5.8.6l4.2.5l-3.1 2.8c-.3.2-.4.6-.3 1c0 .2.5 2.2.8 4.1l-3.6-2.1c-.2-.2-.3-.2-.5-.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.855 20.966c-.224 0-.443-.05-.646-.146l-.104-.051l-4.107-2.343l-4.107 2.344l-.106.053a1.524 1.524 0 0 1-1.521-.143a1.505 1.505 0 0 1-.586-1.509l.957-4.642l-1.602-1.457l-1.895-1.725l-.078-.082a1.503 1.503 0 0 1-.34-1.492c.173-.524.62-.912 1.16-1.009l.102-.018l4.701-.521l1.946-4.31l.06-.11a1.5 1.5 0 0 1 1.309-.771c.543 0 1.044.298 1.309.77l.06.112l1.948 4.312l4.701.521l.104.017c.539.1.986.486 1.158 1.012c.17.521.035 1.098-.34 1.494l-.078.078l-3.498 3.184l.957 4.632a1.514 1.514 0 0 1-.59 1.519a1.488 1.488 0 0 1-.874.281m-8.149-6.564c-.039.182-.466 2.246-.845 4.082l3.643-2.077a1 1 0 0 1 .99 0l3.643 2.075l-.849-4.104a.998.998 0 0 1 .308-.942l3.1-2.822l-4.168-.461a1 1 0 0 1-.801-.584l-1.728-3.821l-1.726 3.821c-.146.322-.45.543-.801.584l-4.168.461l3.1 2.822a.995.995 0 0 1 .302.966"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Starburst(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m19.064 10.109l1.179-2.387a.5.5 0 0 0-.416-.72l-2.656-.172l-.172-2.656a.5.5 0 0 0-.721-.416l-2.385 1.18l-1.477-2.215c-.186-.278-.646-.278-.832 0l-1.477 2.215l-2.385-1.18a.5.5 0 0 0-.721.416L6.83 6.83l-2.657.171a.5.5 0 0 0-.416.721l1.179 2.386l-2.214 1.477a.501.501 0 0 0 0 .832l2.215 1.477l-1.18 2.386a.498.498 0 0 0 .416.72l2.656.171L7 19.828a.5.5 0 0 0 .721.416l2.386-1.179l1.477 2.214a.501.501 0 0 0 .832 0l1.477-2.214l2.386 1.179a.5.5 0 0 0 .721-.416l.171-2.656L19.827 17a.5.5 0 0 0 .416-.721l-1.179-2.385l2.214-1.478a.501.501 0 0 0 0-.832z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarburstOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.556 11.169l-1.849-1.232l.984-1.993a1 1 0 0 0-.832-1.441l-2.218-.143l-.144-2.218a.997.997 0 0 0-1.44-.832l-1.991.985l-1.233-1.849c-.371-.557-1.293-.557-1.664 0L9.935 4.294L7.943 3.31a.998.998 0 0 0-1.441.832l-.143 2.217l-2.218.143a1.002 1.002 0 0 0-.832 1.441l.984 1.993l-1.849 1.233a1.001 1.001 0 0 0 0 1.664l1.85 1.233l-.985 1.992a1 1 0 0 0 .832 1.441l2.218.143l.143 2.218a1 1 0 0 0 1.441.832l1.992-.985l1.233 1.849a1.001 1.001 0 0 0 1.664 0l1.233-1.849l1.991.985a.998.998 0 0 0 1.441-.832l.143-2.217l2.219-.144a.998.998 0 0 0 .832-1.441l-.984-1.992l1.849-1.233a1.001 1.001 0 0 0 0-1.664m-4.032 2.997l.71 1.435l-1.6.104a1.001 1.001 0 0 0-.934.934l-.103 1.598l-1.435-.709a.998.998 0 0 0-1.275.342L12 19.2l-.889-1.333a1 1 0 0 0-1.275-.341l-1.436.709l-.103-1.598a1.001 1.001 0 0 0-.934-.934L5.767 15.6l.71-1.435a1 1 0 0 0-.342-1.275l-1.333-.889l1.332-.888c.418-.279.564-.825.342-1.275l-.71-1.436l1.6-.103c.502-.033.901-.432.934-.934l.103-1.598l1.435.709c.448.221.996.076 1.275-.342L12 4.803l.889 1.333c.279.418.826.563 1.275.342l1.436-.71l.104 1.599c.033.501.433.9.934.933l1.598.103l-.709 1.437a1 1 0 0 0 .342 1.275l1.332.888l-1.333.889a.997.997 0 0 0-.344 1.274"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stopwatch(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.414 8.902a.993.993 0 0 0 .293-.195l.5-.5a.999.999 0 1 0-1.414-1.414l-.5.5l-.115.173a8.969 8.969 0 0 0-5.189-2.41L13 5V4h1c.55 0 1-.45 1-1s-.45-1-1-1h-4c-.55 0-1 .45-1 1s.45 1 1 1h1v1l.012.057A9 9 0 0 0 12 23a9 9 0 0 0 7.414-14.098M12 21c-3.859 0-7-3.14-7-7s3.141-7 7-7s7 3.14 7 7s-3.141 7-7 7m1-8v-2c0-.55-.45-1-1-1s-1 .45-1 1v3c0 .55.45 1 1 1h3c.55 0 1-.45 1-1s-.45-1-1-1zm-1-5c-3.312 0-6 2.688-6 6s2.688 6 6 6s6-2.688 6-6s-2.688-6-6-6m0 11c-2.757 0-5-2.243-5-5s2.243-5 5-5s5 2.243 5 5s-2.243 5-5 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Support(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 3.5c-4.688 0-8.5 3.812-8.5 8.5s3.812 8.5 8.5 8.5s8.5-3.812 8.5-8.5s-3.812-8.5-8.5-8.5m6.5 8.5a6.485 6.485 0 0 1-.718 2.956l-1.931-1.931c.088-.332.147-.674.147-1.025c0-.355-.062-.693-.147-1.021l1.932-1.932c.455.889.717 1.891.717 2.953m-13 0c0-1.064.264-2.066.718-2.956l1.933 1.933a4.03 4.03 0 0 0-.147 1.022c0 .353.062.69.147 1.021l-1.934 1.934A6.47 6.47 0 0 1 5.5 12m3.068-2.02L6.793 8.205l1.414-1.414l1.775 1.775A4.015 4.015 0 0 0 8.568 9.98m-1.777 5.813l1.773-1.773c.17.289.362.564.605.809s.52.438.807.607l-1.771 1.771zm3.795-2.379A2 2 0 0 1 12 10a2 2 0 0 1 1.416 3.413c-.755.757-2.073.757-2.83.001m6.623-5.207l-1.775 1.775a4.007 4.007 0 0 0-1.412-1.416l1.773-1.773zm-2.378 6.619c.241-.242.435-.518.604-.803l1.771 1.771l-1.414 1.414l-1.772-1.772c.291-.17.567-.366.811-.61m.125-8.608L13.023 8.15C12.695 8.062 12.355 8 12 8s-.693.062-1.021.148L9.047 6.216A6.462 6.462 0 0 1 12 5.499a6.499 6.499 0 0 1 2.956.719M9.044 17.782l1.933-1.933c.332.088.672.149 1.023.149s.691-.062 1.021-.147l1.932 1.932A6.462 6.462 0 0 1 12 18.5a6.485 6.485 0 0 1-2.956-.718"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TabsOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 4H8a2 2 0 0 0-2 2v2H5a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h9a2 2 0 0 0 2-2v-1h2a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2M5 19v-9h8.5c.275 0 .5.225.5.5V19zm13-3h-3v-5.5c0-.827-.673-1.5-1.5-1.5H8V6h10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tag(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 4c1.279 0 2.559.488 3.535 1.465L16 9l5 5l-7 7l-5.498-5.498c-.037.033-3.037-2.967-3.037-2.967A4.998 4.998 0 0 1 9 4m0-2c-1.87 0-3.628.729-4.949 2.051C2.729 5.371 2 7.129 2 8.999s.729 3.628 2.051 4.95l3 3c.107.107.227.201.35.279l5.187 5.186c.391.391.9.586 1.413.586s1.022-.195 1.414-.586l7-7c.78-.781.78-2.047 0-2.828l-5-5l-3.45-3.521A6.971 6.971 0 0 0 9 2m0 5.498a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3m0-1A2.503 2.503 0 0 0 6.5 9c0 1.377 1.121 2.498 2.5 2.498S11.5 10.377 11.5 9A2.503 2.503 0 0 0 9 6.498"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tags(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.422 9.594l-6.465-6.535C13.628 1.729 11.87 1 10 1s-3.628.729-4.95 2.051a6.998 6.998 0 0 0-2.027 5.314A6.95 6.95 0 0 0 2 11.999c0 1.87.729 3.628 2.051 4.95l3.053 2.984l3.482 3.48c.391.392.902.587 1.414.587s1.023-.195 1.414-.586l7-7a1.999 1.999 0 0 0 .008-2.82l-.093-.094l1.085-1.086a1.999 1.999 0 0 0 .008-2.82M12 22l-3.498-3.497l-3.037-2.968A4.998 4.998 0 0 1 9 7c1.279 0 2.559.488 3.535 1.465L19 15zm1.957-14.941A6.958 6.958 0 0 0 9 5a6.954 6.954 0 0 0-3.565.982a4.933 4.933 0 0 1 1.03-1.518C7.441 3.488 8.721 3 10 3s2.559.488 3.535 1.465L20 11l-1.078 1.078zM9 10.499a1.5 1.5 0 1 1 .001 2.999A1.5 1.5 0 0 1 9 10.499m0-1A2.504 2.504 0 0 0 6.5 12c0 1.378 1.122 2.499 2.5 2.499s2.5-1.121 2.5-2.499A2.504 2.504 0 0 0 9 9.499"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThLarge(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 3H6a2.99 2.99 0 0 0-2.119.881A2.99 2.99 0 0 0 3 6v2c0 .825.337 1.575.881 2.119A2.99 2.99 0 0 0 6 11h2a2.99 2.99 0 0 0 2.119-.881A2.99 2.99 0 0 0 11 8V6a2.99 2.99 0 0 0-.881-2.119A2.99 2.99 0 0 0 8 3m10 0h-2a2.99 2.99 0 0 0-2.119.881A2.99 2.99 0 0 0 13 6v2c0 .825.337 1.575.881 2.119A2.99 2.99 0 0 0 16 11h2a2.99 2.99 0 0 0 2.119-.881A2.99 2.99 0 0 0 21 8V6a2.99 2.99 0 0 0-.881-2.119A2.99 2.99 0 0 0 18 3M8 13H6a2.99 2.99 0 0 0-2.119.881A2.99 2.99 0 0 0 3 16v2c0 .825.337 1.575.881 2.119A2.99 2.99 0 0 0 6 21h2a2.99 2.99 0 0 0 2.119-.881A2.99 2.99 0 0 0 11 18v-2a2.99 2.99 0 0 0-.881-2.119A2.99 2.99 0 0 0 8 13m10 0h-2a2.99 2.99 0 0 0-2.119.881A2.99 2.99 0 0 0 13 16v2c0 .825.337 1.575.881 2.119A2.99 2.99 0 0 0 16 21h2a2.99 2.99 0 0 0 2.119-.881A2.99 2.99 0 0 0 21 18v-2a2.99 2.99 0 0 0-.881-2.119A2.99 2.99 0 0 0 18 13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThLargeOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 2H4c-1.103 0-2 .896-2 2v5c0 1.104.897 2 2 2h5c1.103 0 2-.896 2-2V4c0-1.104-.897-2-2-2m0 7H4V4h5zm11-7h-5a2 2 0 0 0-2 2v5a2 2 0 0 0 2 2h5a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2m0 7h-5V4h5zM9 13H4c-1.103 0-2 .896-2 2v5c0 1.104.897 2 2 2h5c1.103 0 2-.896 2-2v-5c0-1.104-.897-2-2-2m0 7H4v-5h5zm11-7h-5a2 2 0 0 0-2 2v5a2 2 0 0 0 2 2h5a2 2 0 0 0 2-2v-5a2 2 0 0 0-2-2m0 7h-5v-5h5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThList(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 17h-7c-1.103 0-2 .897-2 2s.897 2 2 2h7c1.103 0 2-.897 2-2s-.897-2-2-2m0-7h-7c-1.103 0-2 .897-2 2s.897 2 2 2h7c1.103 0 2-.897 2-2s-.897-2-2-2m0-7h-7c-1.103 0-2 .897-2 2s.897 2 2 2h7c1.103 0 2-.897 2-2s-.897-2-2-2"/><circle cx="5" cy="19" r="2.5" fill="currentColor"/><circle cx="5" cy="12" r="2.5" fill="currentColor"/><circle cx="5" cy="5" r="2.5" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThListOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 18c.55 0 1 .45 1 1s-.45 1-1 1h-7c-.55 0-1-.45-1-1s.45-1 1-1zm0-2h-7c-1.654 0-3 1.346-3 3s1.346 3 3 3h7c1.654 0 3-1.346 3-3s-1.346-3-3-3m0-5c.55 0 1 .45 1 1s-.45 1-1 1h-7c-.55 0-1-.45-1-1s.45-1 1-1zm0-2h-7c-1.654 0-3 1.346-3 3s1.346 3 3 3h7c1.654 0 3-1.346 3-3s-1.346-3-3-3m0-5c.55 0 1 .45 1 1s-.45 1-1 1h-7c-.55 0-1-.45-1-1s.45-1 1-1zm0-2h-7c-1.654 0-3 1.346-3 3s1.346 3 3 3h7c1.654 0 3-1.346 3-3s-1.346-3-3-3M6 16H4a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2m0 4H4v-2h2zM6 9H4a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2m0 4H4v-2h2zM6 2H4a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2m0 4H4V4h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThMenu(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 17H5c-1.103 0-2 .897-2 2s.897 2 2 2h14c1.103 0 2-.897 2-2s-.897-2-2-2m0-7H5c-1.103 0-2 .897-2 2s.897 2 2 2h14c1.103 0 2-.897 2-2s-.897-2-2-2m0-7H5c-1.103 0-2 .897-2 2s.897 2 2 2h14c1.103 0 2-.897 2-2s-.897-2-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThMenuOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 18c.55 0 1 .45 1 1s-.45 1-1 1H5c-.55 0-1-.45-1-1s.45-1 1-1zm0-2H5c-1.654 0-3 1.346-3 3s1.346 3 3 3h14c1.654 0 3-1.346 3-3s-1.346-3-3-3m0-5c.55 0 1 .45 1 1s-.45 1-1 1H5c-.55 0-1-.45-1-1s.45-1 1-1zm0-2H5c-1.654 0-3 1.346-3 3s1.346 3 3 3h14c1.654 0 3-1.346 3-3s-1.346-3-3-3m0-5c.55 0 1 .45 1 1s-.45 1-1 1H5c-.55 0-1-.45-1-1s.45-1 1-1zm0-2H5C3.346 2 2 3.346 2 5s1.346 3 3 3h14c1.654 0 3-1.346 3-3s-1.346-3-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThSmall(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="5" cy="19" r="2.5" fill="currentColor"/><circle cx="5" cy="12" r="2.5" fill="currentColor"/><circle cx="5" cy="5" r="2.5" fill="currentColor"/><circle cx="12" cy="19" r="2.5" fill="currentColor"/><circle cx="12" cy="12" r="2.5" fill="currentColor"/><circle cx="12" cy="5" r="2.5" fill="currentColor"/><circle cx="19" cy="19" r="2.5" fill="currentColor"/><circle cx="19" cy="12" r="2.5" fill="currentColor"/><circle cx="19" cy="5" r="2.5" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThSmallOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 16H4a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2m0 4H4v-2h2zM6 9H4a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2m0 4H4v-2h2zM6 2H4a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2m0 4H4V4h2zm7 10h-2a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2m0 4h-2v-2h2zm0-11h-2a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2m0 4h-2v-2h2zm0-11h-2a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2m0 4h-2V4h2zm7 10h-2a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2m0 4h-2v-2h2zm0-11h-2a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2m0 4h-2v-2h2zm0-11h-2a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2m0 4h-2V4h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Thermometer(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 15.071V9.5c0-.275-.225-.5-.5-.5s-.5.225-.5.5v5.571c-.86.224-1.5 1-1.5 1.929c0 1.103.896 2 2 2s2-.897 2-2c0-.929-.64-1.705-1.5-1.929m3-1.612V5.5C16 3.57 14.43 2 12.5 2S9 3.57 9 5.5v7.959A4.937 4.937 0 0 0 7.5 17c0 2.757 2.243 5 5 5s5-2.243 5-5c0-1.39-.578-2.639-1.5-3.541M12.5 20c-1.654 0-3-1.346-3-3a2.99 2.99 0 0 1 1.5-2.583V5.5c0-.827.673-1.5 1.5-1.5s1.5.673 1.5 1.5v8.917A2.99 2.99 0 0 1 15.5 17c0 1.654-1.346 3-3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsDown(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 5c-.755 0-1.438.289-1.965.751l-.188-.192C13.887 4.822 11.182 4 9 4c-1.879 0-2.607.293-3.252.552l-.316.124c-.834.305-1.578 1.229-1.738 2.2l-.664 5.972c-.174 1.039.441 2.127 1.4 2.478c.394.144 2.512.405 3.883.56a20.912 20.912 0 0 0-.312 3.616c0 1.379 1.121 2.5 2.5 2.5s2.5-1.121 2.5-2.5c0-1.875.667-2.737 1.616-3.699a2.985 2.985 0 0 0 2.384 1.199a3.004 3.004 0 0 0 2.999-3v-6A3.005 3.005 0 0 0 17 5m-6 14.5c0 .275-.225.5-.5.5s-.5-.225-.5-.5c0-1.805.256-3.241.479-4.293l.297-1.398l-1.321.188c-.605-.05-3.934-.447-4.335-.552c-.058-.028-.132-.18-.108-.321l.663-5.976c.037-.223.291-.539.443-.594l.377-.146C7.039 6.189 7.51 6 9.001 6c1.914 0 4.118.753 4.633 1.146c.156.12.366.564.366.854v4.977c-.001.026-.04.649-.707 1.316C12.38 15.206 11 16.586 11 19.5m7-5.5a1 1 0 0 1-2 0V8a1 1 0 0 1 2 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsOk(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.7 12.6c-.1-.3-.2-.6-.4-.9c.3-.6.5-1 .3-1.9c-.2-.8-.8-1.7-1.6-2.5l-1.7-.9l.8-1.8c.4-.8.4-1.8-.1-2.5c-.5-.8-1.4-1.3-2.3-1.3c-1 0-1.8.6-2.3 1.5c0 0-.8-.2-1.8.1l-1.3 1.2s-1.7.5-2 2s-.4 3.4-.8 4.1S5 13 3.8 14.4c-.9 1.1-.7 2.3.1 3.2l5 5c.8.8 2.7 1.1 4-.3c1.4-1.4.5-3.3.5-3.3c.4-.3 1.5-1.2 2.6-1.5s2.8-.9 3.7-2.1c.8-1 1.2-1.8 1-2.8m-9.3 8.3c-.4.4-1 .4-1.4 0l-4.2-4.2c-.2-.2-.3-.4-.3-.7s.1-.5.3-.7l.7-.7l4.9 4.9c.2.2.3.4.3.7c0 .3-.1.5-.3.7M18 11c-.5 0-.7-.5-1.1-.9s-1.3-.5-2-.2s-1 1-1 1.7c0 1 .7 1.9 1.7 1.9c.7 0 .9-.1 1.3-.3c.6-.4.8-.8 1.2-.8s.7.2.7.8s-.6 1.3-1.2 1.7s-1.1.5-1.7.6c-.6.1-1 .1-1.8.6c-.7.4-1.4.9-1.8 1.3l-4.6-4.6c.9-1.1 1.7-2.3 1.8-2.9l.7-3.9c.1-.4.4-.5.6-.5c.3 0 .6.2.7.4l.4-1.3c.1-.3.4-.5.6-.5c.4 0 .8.3.7.8l-.5 2.4c.6-1.2 1.5-2.7 2.1-3.8c.2-.3.4-.6.9-.6s.9.6.7 1.1c-.2.4-2 3.5-2.8 4.8c-.1.1 0 .2.2.1c.3-.2 1-.7 1.7-.7c1.2-.1 1.8.4 2.1.6c.4.3.8.8 1.1 1.3c.2.5-.2.9-.7.9m-.2.7c-.4 0-.7.1-.9.4s-.6.7-1.1.7c-.7 0-1.1-.5-1.1-1.1s.4-1 1.1-1c.5 0 .9.4 1.1.7s.5.3.9.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsUp(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.57 8.676c-.391-.144-2.512-.406-3.883-.56C15.902 6.861 16 5.711 16 4.5C16 3.121 14.878 2 13.5 2S11 3.121 11 4.5c0 1.875-.666 2.738-1.616 3.699A2.99 2.99 0 0 0 7 7c-1.654 0-3 1.346-3 3v6c0 1.654 1.346 3 3 3c.755 0 1.438-.29 1.965-.752l.188.193c.96.736 3.667 1.559 5.848 1.559c1.879 0 2.608-.293 3.253-.553l.316-.123c.834-.305 1.576-1.227 1.736-2.2l.666-5.974c.173-1.037-.443-2.125-1.402-2.474M7 17c-.551 0-1-.448-1-1v-6a1.001 1.001 0 0 1 2 0v6c0 .552-.449 1-1 1m11.327-.15c-.037.224-.292.541-.443.596l-.376.146c-.545.219-1.016.408-2.508.408c-1.914 0-4.118-.753-4.632-1.146C10.21 16.734 10 16.29 10 16v-4.98c.003-.047.051-.656.707-1.312C11.62 8.794 13 7.414 13 4.5c0-.275.225-.5.5-.5s.5.225.5.5c0 1.407-.146 2.73-.479 4.293l-.297 1.396l1.321-.188c.603.05 3.933.447 4.334.55c.058.03.132.183.111.323z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tick(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.972 6.251a1.999 1.999 0 0 0-2.72.777l-3.713 6.682l-2.125-2.125a2 2 0 1 0-2.828 2.828l4 4c.378.379.888.587 1.414.587l.277-.02a2 2 0 0 0 1.471-1.009l5-9a2 2 0 0 0-.776-2.72"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TickOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 20a2.978 2.978 0 0 1-2.121-.879l-4-4C4.312 14.555 4 13.801 4 13s.312-1.555.879-2.122c1.133-1.133 3.109-1.133 4.242 0l1.188 1.188l3.069-5.523a2.999 2.999 0 0 1 5.507.632a2.975 2.975 0 0 1-.263 2.282l-5 9A3.015 3.015 0 0 1 11 20m-4-8c-.268 0-.518.104-.707.293S6 12.732 6 13s.104.518.293.707l4 4a1.002 1.002 0 0 0 1.581-.221l5-9a.993.993 0 0 0 .088-.76a.992.992 0 0 0-.478-.6a1.015 1.015 0 0 0-1.357.388l-4.357 7.841l-3.062-3.062A.996.996 0 0 0 7 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ticket(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.485 8.071l-5.364-5.364c-1.128-1.128-3.111-1.136-4.248-.018l-9.148 9.002a2.967 2.967 0 0 0-.891 2.115a2.959 2.959 0 0 0 .873 2.121l5.365 5.365c.567.567 1.325.88 2.133.88c.799 0 1.551-.307 2.115-.862l9.147-9.003a2.963 2.963 0 0 0 .891-2.115a2.962 2.962 0 0 0-.873-2.121m-1.421 2.811l-9.146 9.003c-.381.373-1.056.37-1.432-.006l-1.275-1.275a1.999 1.999 0 0 0-.062-2.752a1.997 1.997 0 0 0-2.752-.063l-1.275-1.274c-.186-.187-.288-.435-.287-.699s.105-.513.293-.697l9.148-9.002c.189-.186.441-.288.713-.288c.273 0 .529.104.719.294l1.275 1.275a1.995 1.995 0 0 0 .062 2.751a1.997 1.997 0 0 0 2.752.063l1.274 1.274c.187.187.288.435.287.699s-.105.512-.294.697m-8.463 6.16l-4.657-4.656l5.649-5.429l4.657 4.656zm-3.23-4.643l3.243 3.242L15.82 11.6l-3.241-3.242z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Time(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 13c0-.55-.45-1-1-1h-3c-.55 0-1 .45-1 1s.45 1 1 1h3c.55 0 1-.45 1-1m-4-7c3.859 0 7 3.141 7 7s-3.141 7-7 7s-7-3.141-7-7s3.141-7 7-7m0-2c-4.971 0-9 4.029-9 9s4.029 9 9 9s9-4.029 9-9s-4.029-9-9-9m1 6c0-.55-.45-1-1-1s-1 .45-1 1v3c0 .55.45 1 1 1s1-.45 1-1zm-1-2c2.757 0 5 2.243 5 5s-2.243 5-5 5s-5-2.243-5-5s2.243-5 5-5m0-1a6 6 0 0 0 0 12a6 6 0 0 0 0-12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Times(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.414 6.586a2 2 0 0 0-2.828 0L12 9.172L9.414 6.586a2 2 0 1 0-2.828 2.828L9.171 12l-2.585 2.586a2 2 0 1 0 2.828 2.828L12 14.828l2.586 2.586c.39.391.902.586 1.414.586s1.024-.195 1.414-.586a2 2 0 0 0 0-2.828L14.829 12l2.585-2.586a2 2 0 0 0 0-2.828"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimesOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 19a2.98 2.98 0 0 1-2.122-.879L12 16.242l-1.879 1.879c-1.133 1.133-3.109 1.133-4.243 0C5.312 17.555 5 16.801 5 16s.312-1.555.879-2.122L7.757 12l-1.878-1.879C5.312 9.555 5 8.801 5 8s.312-1.555.879-2.122c1.133-1.132 3.109-1.133 4.243.001L12 7.758l1.879-1.879c1.133-1.133 3.109-1.133 4.243 0C18.688 6.445 19 7.199 19 8s-.312 1.555-.879 2.122L16.243 12l1.878 1.879c.567.566.879 1.32.879 2.121s-.312 1.555-.879 2.122A2.98 2.98 0 0 1 16 19m-4-5.586l3.293 3.293a1.024 1.024 0 0 0 1.414 0c.189-.189.293-.439.293-.707s-.104-.518-.293-.707L13.415 12l3.292-3.293c.189-.189.293-.44.293-.707s-.104-.518-.293-.707a1.023 1.023 0 0 0-1.414-.001L12 10.586L8.707 7.293a1.024 1.024 0 0 0-1.414 0C7.104 7.482 7 7.733 7 8s.104.518.293.707L10.585 12l-3.292 3.293C7.104 15.482 7 15.732 7 16s.104.518.293.707a1.023 1.023 0 0 0 1.414.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trash(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 7h-1V6a2 2 0 0 0-2-2H8a2 2 0 0 0-2 2v1H5a1 1 0 0 0 0 2v8c0 2.206 1.794 4 4 4h5c2.206 0 4-1.794 4-4V9a1 1 0 0 0 0-2M8 6h7v1H8zm8 11a2 2 0 0 1-2 2H9a2 2 0 0 1-2-2V9h9zm-7.5-6.5c-.275 0-.5.225-.5.5v6c0 .275.225.5.5.5s.5-.225.5-.5v-6c0-.275-.225-.5-.5-.5m2 0c-.275 0-.5.225-.5.5v6c0 .275.225.5.5.5s.5-.225.5-.5v-6c0-.275-.225-.5-.5-.5m2 0c-.275 0-.5.225-.5.5v6c0 .275.225.5.5.5s.5-.225.5-.5v-6c0-.275-.225-.5-.5-.5m2 0c-.275 0-.5.225-.5.5v6c0 .275.225.5.5.5s.5-.225.5-.5v-6c0-.275-.225-.5-.5-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tree(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.781 17.375L18.081 14H19a.999.999 0 0 0 .819-1.573l-7-10a1.001 1.001 0 0 0-1.393-.246a.968.968 0 0 0-.221.231c-.025.015-7.025 10.015-7.025 10.015A1 1 0 0 0 5 14h.919l-2.7 3.375c-.24.301-.287.712-.121 1.059c.167.345.518.566.902.566h7v3a1 1 0 1 0 2 0v-3h7a1.001 1.001 0 0 0 .781-1.625M13 17v-5a1 1 0 1 0-2 0v5H6.081l2.7-3.375c.24-.301.287-.712.121-1.059A1.004 1.004 0 0 0 8 12H6.92L12 4.744L17.08 12H16a1.001 1.001 0 0 0-.78 1.625L17.92 17z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Upload(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.987 16a.98.98 0 0 0-.039-.316l-2-6A.998.998 0 0 0 18 9h-4v2h3.279l1.667 5H5.054l1.667-5H10V9H6a.998.998 0 0 0-.948.684l-2 6a.98.98 0 0 0-.039.316C3 16 3 21 3 21a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1s0-5-.013-5M16 7.904c.259 0 .518-.095.707-.283a1 1 0 0 0 0-1.414L12 1.5L7.293 6.207a1 1 0 0 0 0 1.414c.189.189.448.283.707.283s.518-.094.707-.283L11 5.328V12a1 1 0 0 0 2 0V5.328l2.293 2.293a.997.997 0 0 0 .707.283"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UploadOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.986 17c0-.105-.004-.211-.038-.316l-2-6a.996.996 0 0 0-.56-.594a2.995 2.995 0 0 0-.269-3.914L12 .055L5.879 6.176a2.998 2.998 0 0 0-.27 3.914a.987.987 0 0 0-.559.594l-2 6a1.007 1.007 0 0 0-.038.316C3 17 3 22 3 22a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1s0-5-.014-5M7.293 7.59L12 2.883l4.707 4.707a.999.999 0 0 1 0 1.414a1.025 1.025 0 0 1-1.414 0L13 6.711V12.5a1 1 0 0 1-2 0V6.711L8.707 9.004a1.025 1.025 0 0 1-1.414 0a.999.999 0 0 1 0-1.414M6.721 12H9v.5c0 1.654 1.346 3 3 3s3-1.346 3-3V12h2.279l1.666 5H5.053zM5 21v-3h14v3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func User(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 9c0-1.381-.56-2.631-1.464-3.535C14.631 4.56 13.381 4 12 4s-2.631.56-3.536 1.465C7.56 6.369 7 7.619 7 9s.56 2.631 1.464 3.535C9.369 13.44 10.619 14 12 14s2.631-.56 3.536-1.465A4.984 4.984 0 0 0 17 9M6 19c0 1 2.25 2 6 2c3.518 0 6-1 6-2c0-2-2.354-4-6-4c-3.75 0-6 2-6 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserAdd(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 14c1.381 0 2.631-.56 3.536-1.465C13.44 11.631 14 10.381 14 9s-.56-2.631-1.464-3.535C11.631 4.56 10.381 4 9 4s-2.631.56-3.536 1.465C4.56 6.369 4 7.619 4 9s.56 2.631 1.464 3.535A4.985 4.985 0 0 0 9 14m0 7c3.518 0 6-1 6-2c0-2-2.354-4-6-4c-3.75 0-6 2-6 4c0 1 2.25 2 6 2m12-9h-2v-2a1 1 0 1 0-2 0v2h-2a1 1 0 1 0 0 2h2v2a1 1 0 1 0 2 0v-2h2a1 1 0 1 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserAddOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 14h-6a1 1 0 1 1 0-2h6a1 1 0 1 1 0 2m-3 3a1 1 0 0 1-1-1v-6a1 1 0 1 1 2 0v6a1 1 0 0 1-1 1M9 6c1.654 0 3 1.346 3 3s-1.346 3-3 3s-3-1.346-3-3s1.346-3 3-3m0-2C6.236 4 4 6.238 4 9s2.236 5 5 5s5-2.238 5-5s-2.236-5-5-5m0 13c2.021 0 3.301.771 3.783 1.445C12.1 18.705 10.814 19 9 19c-1.984 0-3.206-.305-3.818-.542C5.641 17.743 6.959 17 9 17m0-2c-3.75 0-6 2-6 4c0 1 2.25 2 6 2c3.518 0 6-1 6-2c0-2-2.354-4-6-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserDelete(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 14h-6a1 1 0 1 1 0-2h6a1 1 0 1 1 0 2m-7-5c0 1.381-.56 2.631-1.464 3.535C11.631 13.44 10.381 14 9 14s-2.631-.56-3.536-1.465C4.56 11.631 4 10.381 4 9s.56-2.631 1.464-3.535C6.369 4.56 7.619 4 9 4s2.631.56 3.536 1.465A4.984 4.984 0 0 1 14 9m-5 6c-3.75 0-6 2-6 4c0 1 2.25 2 6 2c3.518 0 6-1 6-2c0-2-2.354-4-6-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserDeleteOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 14h-6a1 1 0 1 1 0-2h6a1 1 0 1 1 0 2M9 6c1.654 0 3 1.346 3 3s-1.346 3-3 3s-3-1.346-3-3s1.346-3 3-3m0-2C6.236 4 4 6.238 4 9s2.236 5 5 5s5-2.238 5-5s-2.236-5-5-5m0 13c2.021 0 3.301.771 3.783 1.445C12.1 18.705 10.814 19 9 19c-1.984 0-3.206-.305-3.818-.542C5.641 17.743 6.959 17 9 17m0-2c-3.75 0-6 2-6 4c0 1 2.25 2 6 2c3.518 0 6-1 6-2c0-2-2.354-4-6-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 6c1.654 0 3 1.346 3 3s-1.346 3-3 3s-3-1.346-3-3s1.346-3 3-3m0-2C9.236 4 7 6.238 7 9s2.236 5 5 5s5-2.238 5-5s-2.236-5-5-5m0 13c2.021 0 3.301.771 3.783 1.445c-.683.26-1.969.555-3.783.555c-1.984 0-3.206-.305-3.818-.542C8.641 17.743 9.959 17 12 17m0-2c-3.75 0-6 2-6 4c0 1 2.25 2 6 2c3.518 0 6-1 6-2c0-2-2.354-4-6-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VendorAndroid(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.1 8h9.8c.1 0 .1 0 .1-.1c-.1-.5-.2-.9-.4-1.3c-.4-.9-1.1-1.5-1.9-2l-.4-.2s0-.1.1-.1c.3-.4.6-.8.9-1.3c.1-.1.1-.2 0-.3c-.1-.1-.2 0-.3.1l-.9 1.3s-.1.1-.1 0c-.8-.3-1.6-.4-2.4-.4c-.6 0-1.1.2-1.6.4h-.1L9 2.9s0-.1-.1-.1c-.1-.1-.2-.1-.2 0c-.1 0-.1.1 0 .2c0 0 0 .1.1.1c.2.4.5.8.8 1.2l.1.1h-.1c-.6.3-1.1.7-1.5 1.2c-.6.6-1 1.4-1.1 2.3c0 .1 0 .1.1.1m7.1-2.8c.4 0 .8.3.8.8c0 .4-.3.8-.7.8c-.4 0-.8-.3-.8-.8c-.1-.4.3-.8.7-.8m-4.3 0c.4 0 .8.3.8.8c0 .4-.3.8-.7.8c-.5 0-.9-.4-.9-.8s.4-.8.8-.8M5 9c-.5 0-1 .5-1 1v5c0 .5.5 1 1 1s1-.5 1-1v-5c0-.5-.5-1-1-1m14 0c-.5 0-1 .5-1 1v5c0 .5.5 1 1 1s1-.5 1-1v-5c0-.5-.5-1-1-1M7 17c0 .5.5 1 1 1h1v3c0 .5.5 1 1 1s1-.5 1-1v-3h2v3c0 .5.5 1 1 1s1-.5 1-1v-3h1c.5 0 1-.5 1-1V9H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VendorApple(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.9 6.6s-.1-1.6.9-3L15.6 2s.1 1.6-.9 3zm5.4 5.6c0-1.5.8-2.8 2-3.6l-.9-.9c-.5-.3-1.1-.7-2.4-.7c-1.4 0-2.4.9-3.7.9c-1.3 0-2.2-.8-3.1-.9c-.7 0-1.4 0-2.1.3c-.5.2-1.2.7-1.6 1.2c-.6.6-1.2 1.9-1.3 3.1c-.1 1.2-.1 2.1.2 3.2c.2.9.6 1.8 1 2.6c.3.6.6 1.2 1 1.8c.3.4.7.8 1.1 1.1c.3.2.6.4 1 .6c.2 0 .5.1.8.1c.6-.1 1.6-.9 2.4-1.1c.4-.1.8-.1 1.3 0c.7.1 1.4.9 2.2 1c.6.1 1.2 0 1.7-.3c.4-.2.7-.5 1-.9c.4-.4.7-.9 1-1.3c.4-.7.9-1.5 1.1-2.3c-1.6-.6-2.7-2.1-2.7-3.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VendorMicrosoft(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 12.5c0-.3-.2-.5-.5-.5h-6c-.3 0-.5.2-.5.5v5c0 .3.2.5.5.6l6 .7c.3 0 .5-.2.5-.4zm1.5-.5c-.3 0-.5.2-.5.5v5.9c0 .3.2.5.5.6l9 1c.3 0 .5-.2.5-.4v-7c0-.3-.2-.5-.5-.5zM10 4.7c0-.3-.2-.5-.5-.4l-6 .7c-.3 0-.5.2-.5.5v5c0 .3.2.5.5.5h6c.3 0 .5-.2.5-.5zm1.5-.6c-.3 0-.5.3-.5.6v5.9c0 .3.2.5.5.5h9c.3 0 .5-.2.5-.5v-7c0-.3-.2-.5-.5-.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Video(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.525 7.149a.997.997 0 0 0-.972-.044L19 8.382V8c0-1.654-1.346-3-3-3H5C3.346 5 2 6.346 2 8v8c0 1.654 1.346 3 3 3h11c1.654 0 3-1.346 3-3v-.382l2.553 1.276a.992.992 0 0 0 .972-.043c.295-.183.475-.504.475-.851V8c0-.347-.18-.668-.475-.851M7 13.5a1.5 1.5 0 1 1-.001-2.999A1.5 1.5 0 0 1 7 13.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="7" cy="11" r="1" fill="currentColor"/><path fill="currentColor" d="M21.585 7.188a.995.995 0 0 0-.901-.137l-1.707.568A3.004 3.004 0 0 0 16 4.999H5c-1.654 0-3 1.347-3 3v6c0 1.653 1.346 3 3 3h3v1.111l.008.09c.066.738.381 1.423.887 1.928c.562.562 1.311.872 2.104.872s1.542-.31 2.104-.87c.574-.577.898-1.346.896-2.113V17h2a3.01 3.01 0 0 0 2.978-2.62l1.707.568c.303.104.64.051.9-.138a.997.997 0 0 0 .415-.812v-6a1 1 0 0 0-.414-.81M12 18.023a.995.995 0 0 1-.309.689a.976.976 0 0 1-1.382 0a1.115 1.115 0 0 1-.309-.691V16h2zM17 14a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1zm3-1.389s-1.895-.605-2-.605V9.994c.105 0 2-.605 2-.605z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Volume(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.138 5.824c-.449 0-.905.152-1.356.453L13.11 8.058C12.357 8.561 10.904 9 10 9c-1.654 0-3 1.346-3 3v2c0 1.654 1.346 3 3 3c.904 0 2.357.439 3.109.941l2.672 1.781c.451.301.907.453 1.356.453c.898.001 1.863-.68 1.863-2.175V8c0-1.495-.965-2.176-1.862-2.176M14 16.146C12.907 15.495 11.211 15 10 15a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1c1.211 0 2.907-.495 4-1.146zM17 18l-.006.12l-.104-.062l-1.89-1.26V9.202l1.891-1.261l.104-.062L17 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeDown(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.138 5.824c-.449 0-.905.152-1.356.453L11.11 8.058C10.357 8.561 8.904 9 8 9c-1.654 0-3 1.346-3 3v2c0 1.654 1.346 3 3 3c.904 0 2.357.439 3.109.941l2.672 1.781c.451.301.907.453 1.356.453c.898.001 1.863-.68 1.863-2.175V8c0-1.495-.965-2.176-1.862-2.176M8 15a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1c1.211 0 2.907-.495 4-1.146v6.293C10.907 15.495 9.211 15 8 15m7 3l-.006.12l-.104-.062l-1.89-1.26V9.202l1.891-1.261l.104-.062L15 8zm3.292-7.706a.999.999 0 0 0 .002 1.414c.345.345.535.803.535 1.291c0 .489-.19.948-.536 1.294a.999.999 0 1 0 1.414 1.414c.724-.723 1.122-1.685 1.122-2.708s-.398-1.984-1.123-2.707a.999.999 0 0 0-1.414.002"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeMute(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.707 5.293a.999.999 0 0 0-1.414 0l-1.551 1.551c-.345-.688-.987-1.02-1.604-1.02c-.449 0-.905.152-1.356.453L11.11 8.058C10.357 8.561 8.904 9 8 9c-1.654 0-3 1.346-3 3v2c0 1.237.754 2.302 1.826 2.76l-1.533 1.533a.999.999 0 1 0 1.414 1.414l2.527-2.527c.697.174 1.416.455 1.875.762l2.672 1.781c.451.301.907.453 1.356.453c.898 0 1.863-.681 1.863-2.176V9.414l2.707-2.707a.999.999 0 0 0 0-1.414m-4.816 2.648l.104-.062L15 8v1.293l-2 2V9.202zM7 12a1 1 0 0 1 1-1c1.211 0 2.907-.495 4-1.146v2.439l-2.83 2.83A6.535 6.535 0 0 0 8 15a1 1 0 0 1-1-1zm3.301 3.406L12 13.707v2.439a8.267 8.267 0 0 0-1.699-.74m4.693 2.714l-.104-.062l-1.89-1.26v-4.091l2-2V18z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeUp(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.706 10.292a.999.999 0 1 0-1.412 1.416c.345.345.535.803.535 1.291c0 .489-.19.948-.536 1.294a.999.999 0 1 0 1.414 1.414c.724-.723 1.122-1.685 1.122-2.708s-.398-1.984-1.123-2.707m2-2a1 1 0 1 0-1.412 1.416a4.616 4.616 0 0 1 1.364 3.287a4.628 4.628 0 0 1-1.365 3.298a.999.999 0 1 0 1.414 1.414a6.617 6.617 0 0 0 1.951-4.713a6.603 6.603 0 0 0-1.952-4.702m2-2a1 1 0 1 0-1.412 1.416a7.42 7.42 0 0 1 2.192 5.284a7.437 7.437 0 0 1-2.193 5.301a.999.999 0 1 0 1.414 1.414a9.427 9.427 0 0 0 2.779-6.717a9.402 9.402 0 0 0-2.78-6.698m-8.568-.468c-.449 0-.905.152-1.356.453L8.109 8.059C7.357 8.561 5.904 9 5 9c-1.654 0-3 1.346-3 3v2c0 1.654 1.346 3 3 3c.904 0 2.357.439 3.109.941l2.672 1.781c.451.301.907.453 1.356.453c.898.001 1.863-.68 1.863-2.175V8c0-1.495-.965-2.176-1.862-2.176M5 15a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1c1.211 0 2.907-.495 4-1.146v6.293C7.907 15.495 6.211 15 5 15m7 3l-.006.12l-.104-.062l-1.89-1.26V9.202l1.891-1.261l.104-.062L12 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Warning(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.171 15.398l-5.912-9.854C14.483 4.251 13.296 3.511 12 3.511s-2.483.74-3.259 2.031l-5.912 9.856c-.786 1.309-.872 2.705-.235 3.83C3.23 20.354 4.472 21 6 21h12c1.528 0 2.77-.646 3.406-1.771c.637-1.125.551-2.521-.235-3.831M12 17.549c-.854 0-1.55-.695-1.55-1.549c0-.855.695-1.551 1.55-1.551s1.55.696 1.55 1.551c0 .854-.696 1.549-1.55 1.549m1.633-7.424c-.011.031-1.401 3.468-1.401 3.468c-.038.094-.13.156-.231.156s-.193-.062-.231-.156l-1.391-3.438a1.776 1.776 0 0 1-.129-.655c0-.965.785-1.75 1.75-1.75a1.752 1.752 0 0 1 1.633 2.375"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WarningOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 5.511c.561 0 1.119.354 1.544 1.062l5.912 9.854C20.307 17.842 19.65 19 18 19H6c-1.65 0-2.307-1.159-1.456-2.573l5.912-9.854c.425-.708.983-1.062 1.544-1.062m0-2c-1.296 0-2.482.74-3.259 2.031l-5.912 9.856c-.786 1.309-.872 2.705-.235 3.83S4.473 21 6 21h12c1.527 0 2.77-.646 3.406-1.771s.551-2.521-.235-3.83l-5.912-9.854C14.482 4.251 13.296 3.511 12 3.511"/><circle cx="12" cy="16" r="1.3" fill="currentColor"/><path fill="currentColor" d="M13.5 10c0-.83-.671-1.5-1.5-1.5a1.499 1.499 0 0 0-1.389 2.062C11.165 11.938 12 14 12 14l1.391-3.438c.068-.173.109-.363.109-.562"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Watch(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 13h2c.55 0 1-.45 1-1s-.45-1-1-1h-1v-1c0-.55-.45-1-1-1s-1 .45-1 1v2c0 .55.45 1 1 1m5-5.895V5c0-1.654-1.346-3-3-3h-4C8.346 2 7 3.346 7 5v2.105C5.764 8.368 5 10.094 5 12s.764 3.632 2 4.895V19c0 1.654 1.346 3 3 3h4c1.654 0 3-1.346 3-3v-2.105c1.236-1.262 2-2.988 2-4.895s-.764-3.632-2-4.895M9 5c0-.551.449-1 1-1h4c.551 0 1 .449 1 1v1.809a5.96 5.96 0 0 0-6 0zm6 14c0 .551-.449 1-1 1h-4c-.551 0-1-.449-1-1v-1.811a5.952 5.952 0 0 0 6 0zm-3-2c-2.757 0-5-2.243-5-5s2.243-5 5-5s5 2.243 5 5s-2.243 5-5 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Waves(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 19a5.23 5.23 0 0 1-3.707-1.532a3.246 3.246 0 0 0-4.586 0a.999.999 0 1 1-1.414-1.414a5.25 5.25 0 0 1 7.414 0a3.248 3.248 0 0 0 4.586 0a.999.999 0 1 1 1.414 1.414A5.227 5.227 0 0 1 15 19m0-4a5.23 5.23 0 0 1-3.707-1.532a3.246 3.246 0 0 0-4.586 0a.999.999 0 1 1-1.414-1.414a5.25 5.25 0 0 1 7.414 0a3.248 3.248 0 0 0 4.586 0a.999.999 0 1 1 1.414 1.414A5.227 5.227 0 0 1 15 15m0-4a5.23 5.23 0 0 1-3.707-1.532a3.246 3.246 0 0 0-4.586 0a.999.999 0 1 1-1.414-1.414a5.25 5.25 0 0 1 7.414 0a3.248 3.248 0 0 0 4.586 0a.999.999 0 1 1 1.414 1.414A5.227 5.227 0 0 1 15 11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WavesOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.221 10.761a2.978 2.978 0 0 0-.1-4.121c-.566-.567-1.32-.879-2.121-.879s-1.555.312-2.121.879a1.239 1.239 0 0 1-.879.361c-.333 0-.646-.129-.879-.362C12.755 5.273 10.936 4.521 9 4.521s-3.755.752-5.121 2.118a2.982 2.982 0 0 0-.1 4.122c-.498.551-.779 1.252-.779 2s.281 1.448.779 2a2.978 2.978 0 0 0 .1 4.121c.566.566 1.32.879 2.121.879s1.555-.312 2.121-.879c.234-.233.545-.362.878-.362c.333 0 .646.129.88.363A7.199 7.199 0 0 0 15 21a7.192 7.192 0 0 0 5.121-2.118a2.98 2.98 0 0 0 .1-4.121c.498-.552.779-1.252.779-2s-.281-1.449-.779-2m-1.514 6.707A5.227 5.227 0 0 1 15 19a5.23 5.23 0 0 1-3.707-1.532A3.235 3.235 0 0 0 9 16.52c-.831 0-1.661.316-2.292.948a1.001 1.001 0 0 1-1.415-1.414a5.226 5.226 0 0 1 3.706-1.532a5.23 5.23 0 0 1 3.708 1.532a3.233 3.233 0 0 0 2.293.947a3.23 3.23 0 0 0 2.293-.947a.997.997 0 0 1 1.414 0a.999.999 0 0 1 0 1.414M5.293 8.054A5.225 5.225 0 0 1 9 6.521c1.343 0 2.685.511 3.707 1.532A3.23 3.23 0 0 0 15 9.001a3.23 3.23 0 0 0 2.293-.947a.997.997 0 0 1 1.414 0a.999.999 0 0 1 0 1.414A5.227 5.227 0 0 1 15 11a5.23 5.23 0 0 1-3.707-1.532A3.235 3.235 0 0 0 9 8.52c-.831 0-1.661.316-2.292.948a1.001 1.001 0 0 1-1.415-1.414m13.414 5.414A5.227 5.227 0 0 1 15 15a5.23 5.23 0 0 1-3.707-1.532A3.235 3.235 0 0 0 9 12.52c-.831 0-1.661.316-2.292.948a1.001 1.001 0 0 1-1.415-1.414a5.226 5.226 0 0 1 3.706-1.532a5.23 5.23 0 0 1 3.708 1.532a3.233 3.233 0 0 0 2.293.947a3.23 3.23 0 0 0 2.293-.947a.997.997 0 0 1 1.414 0a.999.999 0 0 1 0 1.414"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WeatherCloudy(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 19H6c-2.206 0-4-1.794-4-4a4.007 4.007 0 0 1 3.001-3.874L5 11c0-3.309 2.691-6 6-6a5.969 5.969 0 0 1 5.65 4.015C19.592 8.769 22 11.128 22 14c0 2.757-2.243 5-5 5M5.905 12.994C4.897 13 4 13.897 4 15s.897 2 2 2h11c1.654 0 3-1.346 3-3s-1.346-3-3-3c-.243 0-.5.041-.81.13l-1.075.307l-.186-1.103A3.978 3.978 0 0 0 11 7a4.004 4.004 0 0 0-3.918 4.811l.244 1.199z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WeatherDownpour(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 22a1 1 0 0 1-1-1v-6a1 1 0 1 1 2 0v6a1 1 0 0 1-1 1m-6 0a1 1 0 0 1-1-1v-6a1 1 0 1 1 2 0v6a1 1 0 0 1-1 1m3 2a1 1 0 0 1-1-1v-6a1 1 0 1 1 2 0v6a1 1 0 0 1-1 1m-6-6c-2.206 0-4-1.794-4-4a4.007 4.007 0 0 1 3.001-3.874L5 10c0-3.309 2.691-6 6-6a5.97 5.97 0 0 1 5.65 4.015C19.586 7.771 22 10.128 22 13a5.011 5.011 0 0 1-3.666 4.819a1 1 0 1 1-.532-1.927A3.008 3.008 0 0 0 20 13c0-1.654-1.346-3-3-3c-.242 0-.499.041-.811.13l-1.074.306l-.185-1.102A3.98 3.98 0 0 0 11 6a4.004 4.004 0 0 0-3.918 4.808l.248 1.202l-1.422-.016C4.897 12 4 12.897 4 14s.897 2 2 2a1 1 0 1 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WeatherNight(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5 20c-.861 0-1.71-.151-2.523-.451l-1.317-.485l.89-1.087c1.275-1.56 1.95-3.454 1.95-5.477s-.675-3.917-1.951-5.477l-.89-1.087l1.317-.485A7.275 7.275 0 0 1 10.5 5c4.136 0 7.5 3.364 7.5 7.5S14.636 20 10.5 20m-.509-2.024A5.507 5.507 0 0 0 16 12.5a5.507 5.507 0 0 0-6.009-5.476c.991 1.645 1.509 3.511 1.509 5.476s-.518 3.831-1.509 5.476"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WeatherPartlySunny(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14.5 3l-1 3l-1-3a1.053 1.053 0 1 1 2 0m5.364 3.05l-2.829 1.415l1.415-2.829a1.055 1.055 0 0 1 1.886.943a1.054 1.054 0 0 1-.472.471M21.5 12l-3-1l3-1a1.054 1.054 0 1 1 0 2M8.55 4.636l1.415 2.829L7.136 6.05A1.055 1.055 0 1 1 8.55 4.636m9.226 7.706c.139-.424.224-.871.224-1.342c0-2.481-2.019-4.5-4.5-4.5c-1.34 0-2.537.594-3.357 1.528L10 8a5.98 5.98 0 0 0-4.469 2.011a1.054 1.054 0 0 0-1.365 1.322c.057.167.158.299.277.411A5.984 5.984 0 0 0 4 14l.002.126A4.007 4.007 0 0 0 1 18c0 2.206 1.795 4 4 4h11c2.757 0 5-2.243 5-5c0-2.129-1.344-3.939-3.224-4.658M13.5 8.5c1.379 0 2.5 1.121 2.5 2.5c0 .366-.096.706-.238 1.019c-.354.021-.72.074-1.118.188a4.993 4.993 0 0 0-2.967-2.905A2.485 2.485 0 0 1 13.5 8.5M16 20H5c-1.104 0-2-.897-2-2s.896-2 1.908-2.006l1.422.016l-.248-1.202A4.005 4.005 0 0 1 10 10l.069-.014a4.012 4.012 0 0 1 3.847 3.27l.038.186a.985.985 0 0 0 .946.712l.289-.023c.312-.09.569-.131.811-.131c1.654 0 3 1.346 3 3s-1.346 3-3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WeatherShower(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 18a1 1 0 1 1 0-2c1.654 0 3-1.346 3-3s-1.346-3-3-3c-.243 0-.5.041-.81.13l-1.075.307l-.185-1.103A3.98 3.98 0 0 0 11 6a4.004 4.004 0 0 0-3.918 4.811l.244 1.199l-1.42-.016C4.897 12 4 12.897 4 14s.897 2 2 2a1 1 0 1 1 0 2c-2.206 0-4-1.794-4-4a4.007 4.007 0 0 1 3.001-3.874L5 10c0-3.309 2.691-6 6-6a5.969 5.969 0 0 1 5.65 4.015C19.589 7.771 22 10.128 22 13c0 2.757-2.243 5-5 5m-6.5 0l1-3l1 3a1.053 1.053 0 1 1-2 0m3 2l1-3l1 3a1.053 1.053 0 1 1-2 0m-6 0l1-3l1 3a1.053 1.053 0 1 1-2 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WeatherSnow(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.5 15.134l-2.457-.503l1.483-.396a1 1 0 1 0-.518-1.932l-1.508.403c.037-.231.071-.464.071-.706s-.034-.476-.071-.707l1.51.404l.26.034a1 1 0 0 0 .258-1.966l-1.483-.397l2.455-.502a1.054 1.054 0 1 0-1-1.732l-1.663 1.874l.398-1.479a1 1 0 0 0-1.932-.52l-.407 1.517a4.467 4.467 0 0 0-1.219-.717l1.102-1.102a.999.999 0 1 0-1.414-1.414l-1.086 1.086L14.072 4a1.053 1.053 0 1 0-2 0l.793 2.379l-1.086-1.086a.999.999 0 1 0-1.414 1.414l1.102 1.102a4.508 4.508 0 0 0-1.219.717l-.405-1.515a1 1 0 1 0-1.932.516l.396 1.485l-1.662-1.877a1.057 1.057 0 0 0-1.822.487c-.117.57.251 1.128.822 1.245l2.454.503l-1.48.396a1 1 0 0 0 .258 1.966l.26-.034l1.508-.404a4.224 4.224 0 0 0-.073.706c0 .242.034.475.071.707l-1.508-.404a1 1 0 1 0-.518 1.932l1.483.397l-2.455.502a1.052 1.052 0 0 0-.488 1.821a1.053 1.053 0 0 0 1.488-.089l1.665-1.878l-.398 1.484a1 1 0 0 0 .707 1.225l.26.034a1 1 0 0 0 .965-.741l.406-1.515c.366.298.771.544 1.22.716l-1.104 1.102a.999.999 0 1 0 1.414 1.414l.706-.707h.252l-.666 1.999a1.053 1.053 0 1 0 2 0L13.406 18h.252l.707.707a.999.999 0 1 0 1.414-1.414l-1.102-1.103a4.448 4.448 0 0 0 1.22-.717l.406 1.517a1 1 0 0 0 .965.741l.26-.034a1 1 0 0 0 .707-1.225l-.397-1.48l1.662 1.874a1.054 1.054 0 1 0 1-1.732m-7.428-.634c-1.379 0-2.5-1.121-2.5-2.5s1.121-2.5 2.5-2.5s2.5 1.121 2.5 2.5s-1.121 2.5-2.5 2.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WeatherStormy(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 18a1 1 0 1 1 0-2c1.654 0 3-1.346 3-3s-1.346-3-3-3c-.238 0-.496.042-.813.131l-1.071.301l-.186-1.098A3.98 3.98 0 0 0 11 6a4.005 4.005 0 0 0-3.918 4.806l.26 1.24l-1.436-.052C4.896 12 4 12.897 4 14s.896 2 2 2a1 1 0 1 1 0 2c-2.205 0-4-1.794-4-4a4.007 4.007 0 0 1 3.002-3.874L5 10c0-3.309 2.691-6 6-6a5.967 5.967 0 0 1 5.649 4.015C19.574 7.774 22 10.127 22 13c0 2.757-2.243 5-5 5m-4.361-4l-4.5 4.051l3 1.449l-1.5 3.5l4.5-4.05l-3-1.45z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WeatherSunny(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13 4l-1 2.934L11 4a1.056 1.056 0 1 1 2 0m-9 7l2.934 1L4 13a1.056 1.056 0 1 1 0-2m7 9l1-2.934L13 20a1.056 1.056 0 1 1-2 0m9-7.002l-2.934-1l2.934-1a1.056 1.056 0 1 1 0 2M7.05 5.636l1.367 2.781L5.636 7.05A1.057 1.057 0 1 1 7.05 5.636M5.636 16.949l2.781-1.367l-1.367 2.781a1.057 1.057 0 1 1-1.414-1.414m11.313 1.414l-1.367-2.781l2.781 1.367a1.057 1.057 0 1 1-1.414 1.414m1.413-11.315L15.58 8.416l1.368-2.782a1.057 1.057 0 1 1 1.414 1.414M12 16.5c-2.481 0-4.5-2.019-4.5-4.5S9.519 7.5 12 7.5s4.5 2.019 4.5 4.5s-2.019 4.5-4.5 4.5m0-7c-1.379 0-2.5 1.121-2.5 2.5s1.121 2.5 2.5 2.5s2.5-1.121 2.5-2.5s-1.121-2.5-2.5-2.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WeatherWindy(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 5a1 1 0 1 0 0 2a1 1 0 0 1 0 2H8a1 1 0 1 0 0 2h6a1 1 0 0 1 0 2H7.6c-1.654 0-3 1.346-3 3s1.346 3 3 3a1 1 0 1 0 0-2a1 1 0 0 1 0-2H14c1.654 0 3-1.346 3-3c0-.353-.072-.686-.185-1H19c1.654 0 3-1.346 3-3s-1.346-3-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WeatherWindyCloudy(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.798 15.75a.978.978 0 0 1-.4-.084A4.004 4.004 0 0 1 2 12a4.007 4.007 0 0 1 3.001-3.874L5 8c0-3.309 2.691-6 6-6a5.974 5.974 0 0 1 5.902 5.001a1 1 0 0 1-.82 1.152a1.007 1.007 0 0 1-1.152-.82A3.98 3.98 0 0 0 11 4a4.004 4.004 0 0 0-3.919 4.812l.259 1.27l-1.431-.088C4.897 10 4 10.897 4 12c0 .795.471 1.515 1.2 1.834a1 1 0 0 1-.402 1.916M19 7a1 1 0 1 0 0 2a1 1 0 0 1 0 2H9.4a1 1 0 1 0 0 2H14a1 1 0 0 1 0 2H9c-1.654 0-3 1.346-3 3s1.346 3 3 3a1 1 0 1 0 0-2a1 1 0 0 1 0-2h5c1.654 0 3-1.346 3-3c0-.353-.072-.686-.185-1H19c1.654 0 3-1.346 3-3s-1.346-3-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WiFi(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.414 19.412a1.994 1.994 0 0 0 0-2.826a1.994 1.994 0 0 0-2.828-.002a2 2 0 1 0 2.828 2.828m7.071-7.897a1.99 1.99 0 0 1-1.414-.586c-3.899-3.899-10.243-3.898-14.143 0A2 2 0 0 1 2.099 8.1c5.459-5.458 14.341-5.458 19.799 0a2 2 0 0 1-1.413 3.415M7.757 15.757a2 2 0 0 1-1.414-3.414c3.118-3.119 8.194-3.119 11.313 0a2 2 0 0 1-2.829 2.829a4.005 4.005 0 0 0-5.657 0a1.99 1.99 0 0 1-1.413.585"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WiFiOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.157 10.764c0-.785-.269-1.464-.706-2.048c-.045-.094-.131-.149-.21-.226a3.485 3.485 0 0 0-.536-.48c-4.45-3.739-10.965-3.735-15.414.006c-.193.142-.742.738-.742.738a3.48 3.48 0 0 0-.706 2.09c0 .816.362 1.758.759 2.155l5.775 5.796a3.477 3.477 0 0 0 2.622 1.204c.996 0 1.709-.167 2.526-1c.004 0 5.565-5.646 5.565-5.646c.706-.703 1.067-1.699 1.067-2.589m-9.156 7.234a1.5 1.5 0 1 1-.002-3a1.5 1.5 0 0 1 .002 3m3.888-3.268a1.494 1.494 0 0 1-1.061.438c-.385 0-.768-.146-1.061-.438a2.501 2.501 0 0 0-3.536 0a1.502 1.502 0 0 1-2.122-2.123a5.508 5.508 0 0 1 7.779 0a1.5 1.5 0 0 1 .001 2.123m2.829-2.828c-.293.293-.677.438-1.061.438s-.769-.146-1.062-.438a6.508 6.508 0 0 0-9.192 0a1.504 1.504 0 0 1-2.122 0a1.503 1.503 0 0 1 0-2.123c3.704-3.701 9.729-3.701 13.435 0c.587.588.587 1.537.002 2.123"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wine(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.568 9.432c0-2.55-.906-5.592-.944-5.72A1.001 1.001 0 0 0 15.666 3H8.334c-.441 0-.83.289-.958.712c-.038.128-.944 3.17-.944 5.72c0 2.735 1.984 5.011 4.587 5.477L11 15v4h-1a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2h-1v-4l-.019-.092c2.603-.466 4.587-2.741 4.587-5.476M12 13a3.565 3.565 0 0 1-3.511-3h7.021A3.563 3.563 0 0 1 12 13M8.445 9c.062-1.468.422-3.093.653-4h5.803c.231.907.591 2.532.653 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func World(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2c-4.971 0-9 4.029-9 9s4.029 9 9 9s9-4.029 9-9s-4.029-9-9-9m2 2c0 1-.5 2-1.5 2S11 7 11 8v3s1 0 1-3a1 1 0 1 1 2 0v3a1 1 0 1 0 1 1h1v-2l1 1l-1 1c0 3 0 3-2 4c0-1-1-1-3-1v-2l-2-2V9c-1 0-1 1-1 1l-.561-.561l-2.39-2.39c.11-.192.225-.382.35-.564l.523-.678A7.977 7.977 0 0 1 12 3a8.05 8.05 0 0 1 2 .262z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WorldOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2c-4.971 0-9 4.029-9 9s4.029 9 9 9s9-4.029 9-9s-4.029-9-9-9m0 6a1 1 0 1 1 2 0v3a1 1 0 1 0 1 1h1v-2l1 1l-1 1c0 3 0 3-2 4c0-1-1-1-3-1v-2l-2-2V9c-1 0-1 1-1 1l-.561-.561l-1.652-1.651A7.004 7.004 0 0 1 12 4c.688 0 1.353.104 1.981.29C13.895 5.185 13.402 6 12.5 6c-1 0-1.5 1-1.5 2v3s1 0 1-3m0 10c-3.859 0-7-3.14-7-7c0-.776.133-1.521.367-2.219l1.926 1.926l1 1L10 13.414V15a1 1 0 0 0 1 1c.779 0 1.651 0 2.006.091a1.028 1.028 0 0 0 1.02.902a.946.946 0 0 0 .422-.098c2.348-1.174 2.539-1.644 2.552-4.479l.708-.708a.999.999 0 0 0 0-1.414l-1-1A1.001 1.001 0 0 0 15 10V8a2 2 0 0 0-.883-1.658c.421-.411.712-.995.826-1.685C17.335 5.772 19 8.192 19 11c0 3.86-3.141 7-7 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Zoom(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 4c-3.859 0-7 3.141-7 7c0 .763.127 1.495.354 2.183l-.749.75l-.511.512l-1.008 1.045a3.076 3.076 0 0 0-.891 2.185a3.134 3.134 0 0 0 3.13 3.131c.757 0 1.504-.278 2.104-.784l.064-.055l.061-.061l1.512-1.51l.75-.749A6.983 6.983 0 0 0 13 18c3.859 0 7-3.141 7-7s-3.141-7-7-7m0 12c-2.757 0-5-2.243-5-5s2.243-5 5-5s5 2.243 5 5s-2.243 5-5 5m0-9c-2.205 0-4 1.794-4 4s1.795 4 4 4s4-1.794 4-4s-1.795-4-4-4m0 7a3.001 3.001 0 0 1 0-6a3.001 3.001 0 0 1 0 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomIn(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 11h-2V9a.5.5 0 0 0-1 0v2H8a.5.5 0 0 0 0 1h2v2a.5.5 0 0 0 1 0v-2h2a.5.5 0 0 0 0-1m5.432 3.97l-.536-.537l-.749-.75c.227-.688.354-1.42.354-2.183c0-3.859-3.141-7-7-7s-7 3.141-7 7s3.141 7 7 7c.763 0 1.496-.127 2.184-.354l.75.749l1.512 1.51l.061.061l.064.055a3.28 3.28 0 0 0 2.104.784a3.134 3.134 0 0 0 3.13-3.131c0-.84-.328-1.628-.924-2.218zM5.5 11.5c0-2.757 2.243-5 5-5s5 2.243 5 5s-2.243 5-5 5s-5-2.243-5-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomInOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 11h-2V9c0-.275-.225-.5-.5-.5s-.5.225-.5.5v2H8c-.275 0-.5.225-.5.5s.225.5.5.5h2v2c0 .275.225.5.5.5s.5-.225.5-.5v-2h2c.275 0 .5-.225.5-.5s-.225-.5-.5-.5m6.381 4.956l-2.244-2.283A6.896 6.896 0 0 0 17.5 11.5c0-3.859-3.141-7-7-7s-7 3.141-7 7s3.141 7 7 7c.762 0 1.488-.137 2.173-.364l2.397 2.386a3.279 3.279 0 0 0 2.104.783a3.134 3.134 0 0 0 3.131-3.131c0-.84-.328-1.628-.924-2.218m-3.901-1.11l2.492 2.531c.205.203.332.486.332.797c0 .625-.507 1.131-1.131 1.131c-.312 0-.594-.127-.816-.313l-2.512-2.511a6.056 6.056 0 0 0 1.635-1.635M5.5 11.5c0-2.757 2.243-5 5-5s5 2.243 5 5s-2.243 5-5 5s-5-2.243-5-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomOut(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 11H8a.5.5 0 0 0 0 1h5a.5.5 0 0 0 0-1m6.381 4.956l-.949-.986l-.537-.537l-.749-.75c.227-.688.354-1.42.354-2.183c0-3.859-3.14-7-7-7s-7 3.141-7 7s3.14 7 7 7c.763 0 1.496-.127 2.184-.354l.75.749l1.512 1.51l.06.061l.065.055a3.28 3.28 0 0 0 2.104.784a3.134 3.134 0 0 0 3.13-3.131c0-.84-.328-1.628-.924-2.218M5.5 11.5c0-2.757 2.243-5 5-5s5 2.243 5 5s-2.243 5-5 5s-5-2.243-5-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomOutOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 12H8c-.275 0-.5-.225-.5-.5s.225-.5.5-.5h5c.275 0 .5.225.5.5s-.225.5-.5.5m6.381 3.956l-2.245-2.283A6.868 6.868 0 0 0 17.5 11.5c0-3.859-3.141-7-7-7s-7 3.141-7 7s3.141 7 7 7c.761 0 1.488-.137 2.173-.364l2.397 2.386a3.279 3.279 0 0 0 2.104.783a3.134 3.134 0 0 0 3.131-3.131c0-.84-.328-1.628-.924-2.218m-3.901-1.11l2.492 2.531c.205.203.332.486.332.797c0 .625-.507 1.131-1.131 1.131c-.312 0-.594-.127-.816-.313l-2.512-2.511a6.056 6.056 0 0 0 1.635-1.635M5.5 11.5c0-2.757 2.243-5 5-5s5 2.243 5 5s-2.243 5-5 5s-5-2.243-5-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomOutline(children ...ElementRenderer) *TypcnIcon {
	return &TypcnIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 8a3.001 3.001 0 0 1 0 6a3.001 3.001 0 0 1 0-6m0-1c-2.206 0-4 1.794-4 4s1.794 4 4 4s4-1.794 4-4s-1.794-4-4-4M4.195 17.674a3.134 3.134 0 0 0 3.131 3.131c.756 0 1.503-.277 2.104-.783l2.397-2.386A6.887 6.887 0 0 0 14 18c3.86 0 7-3.141 7-7s-3.14-7-7-7c-3.859 0-7 3.141-7 7c0 .761.136 1.486.364 2.173l-2.245 2.283a3.099 3.099 0 0 0-.924 2.218m6.459-1.694l-2.512 2.511a1.272 1.272 0 0 1-.816.313a1.131 1.131 0 0 1-1.131-1.131c0-.311.127-.594.332-.797l2.492-2.531c.435.645.99 1.2 1.635 1.635M14 16c-2.757 0-5-2.243-5-5s2.243-5 5-5s5 2.243 5 5s-2.243 5-5 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
