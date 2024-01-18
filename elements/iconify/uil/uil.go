package uil

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = "4.0.8"
	hAttr          = "1em"
	viewbox        = "0 0 24 24"
	fill           = "currentColor"
)

type UilIcon struct {
	*SVGSVGElement
}

type UilIconFn func(children ...ElementRenderer) *UilIcon

var IconLookup = map[string]UilIconFn{
	"abacus":                       Abacus,
	"accessibleIconAlt":            AccessibleIconAlt,
	"adjust":                       Adjust,
	"adjustAlt":                    AdjustAlt,
	"adjustCircle":                 AdjustCircle,
	"adjustHalf":                   AdjustHalf,
	"adobe":                        Adobe,
	"adobeAlt":                     AdobeAlt,
	"airplay":                      Airplay,
	"align":                        Align,
	"alignAlt":                     AlignAlt,
	"alignCenter":                  AlignCenter,
	"alignCenterAlt":               AlignCenterAlt,
	"alignCenterH":                 AlignCenterH,
	"alignCenterJustify":           AlignCenterJustify,
	"alignCenterV":                 AlignCenterV,
	"alignJustify":                 AlignJustify,
	"alignLeft":                    AlignLeft,
	"alignLeftJustify":             AlignLeftJustify,
	"alignLetterRight":             AlignLetterRight,
	"alignRight":                   AlignRight,
	"alignRightJustify":            AlignRightJustify,
	"amazon":                       Amazon,
	"ambulance":                    Ambulance,
	"analysis":                     Analysis,
	"analytics":                    Analytics,
	"anchor":                       Anchor,
	"android":                      Android,
	"androidAlt":                   AndroidAlt,
	"androidPhoneSlash":            AndroidPhoneSlash,
	"angleDoubleDown":              AngleDoubleDown,
	"angleDoubleLeft":              AngleDoubleLeft,
	"angleDoubleRight":             AngleDoubleRight,
	"angleDoubleUp":                AngleDoubleUp,
	"angleDown":                    AngleDown,
	"angleLeft":                    AngleLeft,
	"angleLeftB":                   AngleLeftB,
	"angleRight":                   AngleRight,
	"angleRightB":                  AngleRightB,
	"angleUp":                      AngleUp,
	"angry":                        Angry,
	"ankh":                         Ankh,
	"annoyed":                      Annoyed,
	"annoyedAlt":                   AnnoyedAlt,
	"apple":                        Apple,
	"appleAlt":                     AppleAlt,
	"apps":                         Apps,
	"archive":                      Archive,
	"archiveAlt":                   ArchiveAlt,
	"archway":                      Archway,
	"arrow":                        Arrow,
	"arrowBreak":                   ArrowBreak,
	"arrowCircleDown":              ArrowCircleDown,
	"arrowCircleLeft":              ArrowCircleLeft,
	"arrowCircleRight":             ArrowCircleRight,
	"arrowCircleUp":                ArrowCircleUp,
	"arrowCompressH":               ArrowCompressH,
	"arrowDown":                    ArrowDown,
	"arrowDownLeft":                ArrowDownLeft,
	"arrowDownRight":               ArrowDownRight,
	"arrowFromRight":               ArrowFromRight,
	"arrowFromTop":                 ArrowFromTop,
	"arrowGrowth":                  ArrowGrowth,
	"arrowLeft":                    ArrowLeft,
	"arrowRandom":                  ArrowRandom,
	"arrowResizeDiagonal":          ArrowResizeDiagonal,
	"arrowRight":                   ArrowRight,
	"arrowToBottom":                ArrowToBottom,
	"arrowToRight":                 ArrowToRight,
	"arrowUp":                      ArrowUp,
	"arrowUpLeft":                  ArrowUpLeft,
	"arrowUpRight":                 ArrowUpRight,
	"arrowsH":                      ArrowsH,
	"arrowsHalt":                   ArrowsHalt,
	"arrowsLeftDown":               ArrowsLeftDown,
	"arrowsMaximize":               ArrowsMaximize,
	"arrowsMerge":                  ArrowsMerge,
	"arrowsResize":                 ArrowsResize,
	"arrowsResizeH":                ArrowsResizeH,
	"arrowsResizeV":                ArrowsResizeV,
	"arrowsRightDown":              ArrowsRightDown,
	"arrowsShrinkH":                ArrowsShrinkH,
	"arrowsShrinkV":                ArrowsShrinkV,
	"arrowsUpRight":                ArrowsUpRight,
	"arrowsV":                      ArrowsV,
	"arrowsValt":                   ArrowsValt,
	"assistiveListeningSystems":    AssistiveListeningSystems,
	"asterisk":                     Asterisk,
	"at":                           At,
	"atom":                         Atom,
	"autoFlash":                    AutoFlash,
	"award":                        Award,
	"awardAlt":                     AwardAlt,
	"babyCarriage":                 BabyCarriage,
	"backpack":                     Backpack,
	"backspace":                    Backspace,
	"backward":                     Backward,
	"bag":                          Bag,
	"bagAlt":                       BagAlt,
	"bagSlash":                     BagSlash,
	"balanceScale":                 BalanceScale,
	"ban":                          Ban,
	"bandAid":                      BandAid,
	"bars":                         Bars,
	"baseballBall":                 BaseballBall,
	"basketball":                   Basketball,
	"basketballHoop":               BasketballHoop,
	"bath":                         Bath,
	"batteryBolt":                  BatteryBolt,
	"batteryEmpty":                 BatteryEmpty,
	"bed":                          Bed,
	"bedDouble":                    BedDouble,
	"behance":                      Behance,
	"behanceAlt":                   BehanceAlt,
	"bell":                         Bell,
	"bellSchool":                   BellSchool,
	"bellSlash":                    BellSlash,
	"bill":                         Bill,
	"bing":                         Bing,
	"bitcoin":                      Bitcoin,
	"bitcoinAlt":                   BitcoinAlt,
	"bitcoinCircle":                BitcoinCircle,
	"bitcoinSign":                  BitcoinSign,
	"blackBerry":                   BlackBerry,
	"blogger":                      Blogger,
	"bloggerAlt":                   BloggerAlt,
	"bluetoothB":                   BluetoothB,
	"bold":                         Bold,
	"bolt":                         Bolt,
	"boltAlt":                      BoltAlt,
	"boltSlash":                    BoltSlash,
	"book":                         Book,
	"bookAlt":                      BookAlt,
	"bookMedical":                  BookMedical,
	"bookOpen":                     BookOpen,
	"bookReader":                   BookReader,
	"bookmark":                     Bookmark,
	"bookmarkFull":                 BookmarkFull,
	"books":                        Books,
	"boombox":                      Boombox,
	"borderAlt":                    BorderAlt,
	"borderBottom":                 BorderBottom,
	"borderClear":                  BorderClear,
	"borderHorizontal":             BorderHorizontal,
	"borderInner":                  BorderInner,
	"borderLeft":                   BorderLeft,
	"borderOut":                    BorderOut,
	"borderRight":                  BorderRight,
	"borderTop":                    BorderTop,
	"borderVertical":               BorderVertical,
	"bowlingBall":                  BowlingBall,
	"box":                          Box,
	"bracketsCurly":                BracketsCurly,
	"brain":                        Brain,
	"briefcase":                    Briefcase,
	"briefcaseAlt":                 BriefcaseAlt,
	"bright":                       Bright,
	"brightness":                   Brightness,
	"brightnessEmpty":              BrightnessEmpty,
	"brightnessHalf":               BrightnessHalf,
	"brightnessLow":                BrightnessLow,
	"brightnessMinus":              BrightnessMinus,
	"brightnessPlus":               BrightnessPlus,
	"bringBottom":                  BringBottom,
	"bringFront":                   BringFront,
	"browser":                      Browser,
	"brushAlt":                     BrushAlt,
	"bug":                          Bug,
	"building":                     Building,
	"bullseye":                     Bullseye,
	"bus":                          Bus,
	"busAlt":                       BusAlt,
	"busSchool":                    BusSchool,
	"calculator":                   Calculator,
	"calculatorAlt":                CalculatorAlt,
	"calendarAlt":                  CalendarAlt,
	"calendarSlash":                CalendarSlash,
	"calender":                     Calender,
	"calling":                      Calling,
	"camera":                       Camera,
	"cameraChange":                 CameraChange,
	"cameraPlus":                   CameraPlus,
	"cameraSlash":                  CameraSlash,
	"cancel":                       Cancel,
	"capsule":                      Capsule,
	"capture":                      Capture,
	"car":                          Car,
	"carSideview":                  CarSideview,
	"carSlash":                     CarSlash,
	"carWash":                      CarWash,
	"cardAtm":                      CardAtm,
	"caretRight":                   CaretRight,
	"cell":                         Cell,
	"celsius":                      Celsius,
	"channel":                      Channel,
	"channelAdd":                   ChannelAdd,
	"chart":                        Chart,
	"chartBar":                     ChartBar,
	"chartBarAlt":                  ChartBarAlt,
	"chartDown":                    ChartDown,
	"chartGrowth":                  ChartGrowth,
	"chartGrowthAlt":               ChartGrowthAlt,
	"chartLine":                    ChartLine,
	"chartPie":                     ChartPie,
	"chartPieAlt":                  ChartPieAlt,
	"chat":                         Chat,
	"chatBubbleUser":               ChatBubbleUser,
	"chatInfo":                     ChatInfo,
	"check":                        Check,
	"checkCircle":                  CheckCircle,
	"checkSquare":                  CheckSquare,
	"circle":                       Circle,
	"circleLayer":                  CircleLayer,
	"circuit":                      Circuit,
	"clapperBoard":                 ClapperBoard,
	"clinicMedical":                ClinicMedical,
	"clipboard":                    Clipboard,
	"clipboardAlt":                 ClipboardAlt,
	"clipboardBlank":               ClipboardBlank,
	"clipboardNotes":               ClipboardNotes,
	"clock":                        Clock,
	"clockEight":                   ClockEight,
	"clockFive":                    ClockFive,
	"clockNine":                    ClockNine,
	"clockSeven":                   ClockSeven,
	"clockTen":                     ClockTen,
	"clockThree":                   ClockThree,
	"clockTwo":                     ClockTwo,
	"closedCaptioning":             ClosedCaptioning,
	"closedCaptioningSlash":        ClosedCaptioningSlash,
	"cloud":                        Cloud,
	"cloudBlock":                   CloudBlock,
	"cloudBookmark":                CloudBookmark,
	"cloudCheck":                   CloudCheck,
	"cloudComputing":               CloudComputing,
	"cloudDataConnection":          CloudDataConnection,
	"cloudDatabaseTree":            CloudDatabaseTree,
	"cloudDownload":                CloudDownload,
	"cloudDrizzle":                 CloudDrizzle,
	"cloudExclamation":             CloudExclamation,
	"cloudHail":                    CloudHail,
	"cloudHeart":                   CloudHeart,
	"cloudInfo":                    CloudInfo,
	"cloudLock":                    CloudLock,
	"cloudMeatball":                CloudMeatball,
	"cloudMoon":                    CloudMoon,
	"cloudMoonHail":                CloudMoonHail,
	"cloudMoonMeatball":            CloudMoonMeatball,
	"cloudMoonRain":                CloudMoonRain,
	"cloudMoonShowers":             CloudMoonShowers,
	"cloudQuestion":                CloudQuestion,
	"cloudRain":                    CloudRain,
	"cloudRainSun":                 CloudRainSun,
	"cloudRedo":                    CloudRedo,
	"cloudShare":                   CloudShare,
	"cloudShield":                  CloudShield,
	"cloudShowers":                 CloudShowers,
	"cloudShowersAlt":              CloudShowersAlt,
	"cloudShowersHeavy":            CloudShowersHeavy,
	"cloudSlash":                   CloudSlash,
	"cloudSun":                     CloudSun,
	"cloudSunHail":                 CloudSunHail,
	"cloudSunMeatball":             CloudSunMeatball,
	"cloudSunRain":                 CloudSunRain,
	"cloudSunRainAlt":              CloudSunRainAlt,
	"cloudSunTear":                 CloudSunTear,
	"cloudTimes":                   CloudTimes,
	"cloudUnlock":                  CloudUnlock,
	"cloudUpload":                  CloudUpload,
	"cloudWifi":                    CloudWifi,
	"cloudWind":                    CloudWind,
	"clouds":                       Clouds,
	"club":                         Club,
	"codeBranch":                   CodeBranch,
	"coffee":                       Coffee,
	"cog":                          Cog,
	"coins":                        Coins,
	"columns":                      Columns,
	"comment":                      Comment,
	"commentAdd":                   CommentAdd,
	"commentAlt":                   CommentAlt,
	"commentAltBlock":              CommentAltBlock,
	"commentAltChartLines":         CommentAltChartLines,
	"commentAltCheck":              CommentAltCheck,
	"commentAltDots":               CommentAltDots,
	"commentAltDownload":           CommentAltDownload,
	"commentAltEdit":               CommentAltEdit,
	"commentAltExclamation":        CommentAltExclamation,
	"commentAltHeart":              CommentAltHeart,
	"commentAltImage":              CommentAltImage,
	"commentAltInfo":               CommentAltInfo,
	"commentAltLines":              CommentAltLines,
	"commentAltLock":               CommentAltLock,
	"commentAltMedical":            CommentAltMedical,
	"commentAltMessage":            CommentAltMessage,
	"commentAltNotes":              CommentAltNotes,
	"commentAltPlus":               CommentAltPlus,
	"commentAltQuestion":           CommentAltQuestion,
	"commentAltRedo":               CommentAltRedo,
	"commentAltSearch":             CommentAltSearch,
	"commentAltShare":              CommentAltShare,
	"commentAltShield":             CommentAltShield,
	"commentAltSlash":              CommentAltSlash,
	"commentAltUpload":             CommentAltUpload,
	"commentAltVerify":             CommentAltVerify,
	"commentBlock":                 CommentBlock,
	"commentChartLine":             CommentChartLine,
	"commentCheck":                 CommentCheck,
	"commentDots":                  CommentDots,
	"commentDownload":              CommentDownload,
	"commentEdit":                  CommentEdit,
	"commentExclamation":           CommentExclamation,
	"commentHeart":                 CommentHeart,
	"commentImage":                 CommentImage,
	"commentInfo":                  CommentInfo,
	"commentInfoAlt":               CommentInfoAlt,
	"commentLines":                 CommentLines,
	"commentLock":                  CommentLock,
	"commentMedical":               CommentMedical,
	"commentMessage":               CommentMessage,
	"commentNotes":                 CommentNotes,
	"commentPlus":                  CommentPlus,
	"commentQuestion":              CommentQuestion,
	"commentRedo":                  CommentRedo,
	"commentSearch":                CommentSearch,
	"commentShare":                 CommentShare,
	"commentShield":                CommentShield,
	"commentSlash":                 CommentSlash,
	"commentUpload":                CommentUpload,
	"commentVerify":                CommentVerify,
	"comments":                     Comments,
	"commentsAlt":                  CommentsAlt,
	"compactDisc":                  CompactDisc,
	"comparison":                   Comparison,
	"compass":                      Compass,
	"compress":                     Compress,
	"compressAlt":                  CompressAlt,
	"compressAltLeft":              CompressAltLeft,
	"compressArrows":               CompressArrows,
	"compressLines":                CompressLines,
	"compressPoint":                CompressPoint,
	"compressV":                    CompressV,
	"confused":                     Confused,
	"constructor":                  Constructor,
	"copy":                         Copy,
	"copyAlt":                      CopyAlt,
	"copyLandscape":                CopyLandscape,
	"copyright":                    Copyright,
	"cornerDownLeft":               CornerDownLeft,
	"cornerDownRight":              CornerDownRight,
	"cornerDownRightAlt":           CornerDownRightAlt,
	"cornerLeftDown":               CornerLeftDown,
	"cornerRightDown":              CornerRightDown,
	"cornerUpLeft":                 CornerUpLeft,
	"cornerUpLeftAlt":              CornerUpLeftAlt,
	"cornerUpRight":                CornerUpRight,
	"cornerUpRightAlt":             CornerUpRightAlt,
	"coronavirus":                  Coronavirus,
	"createDashboard":              CreateDashboard,
	"creativeCommonsPd":            CreativeCommonsPd,
	"creditCard":                   CreditCard,
	"creditCardSearch":             CreditCardSearch,
	"crockery":                     Crockery,
	"cropAlt":                      CropAlt,
	"cropAltRotateLeft":            CropAltRotateLeft,
	"cropAltRotateRight":           CropAltRotateRight,
	"crosshair":                    Crosshair,
	"crosshairAlt":                 CrosshairAlt,
	"crosshairs":                   Crosshairs,
	"cssThreeSimple":               CssThreeSimple,
	"cube":                         Cube,
	"dashboard":                    Dashboard,
	"dataSharing":                  DataSharing,
	"database":                     Database,
	"databaseAlt":                  DatabaseAlt,
	"desert":                       Desert,
	"desktop":                      Desktop,
	"desktopAlt":                   DesktopAlt,
	"desktopAltSlash":              DesktopAltSlash,
	"desktopCloudAlt":              DesktopCloudAlt,
	"desktopSlash":                 DesktopSlash,
	"dialpad":                      Dialpad,
	"dialpadAlt":                   DialpadAlt,
	"diamond":                      Diamond,
	"diary":                        Diary,
	"diaryAlt":                     DiaryAlt,
	"diceFive":                     DiceFive,
	"diceFour":                     DiceFour,
	"diceOne":                      DiceOne,
	"diceSix":                      DiceSix,
	"diceThree":                    DiceThree,
	"diceTwo":                      DiceTwo,
	"direction":                    Direction,
	"directions":                   Directions,
	"discord":                      Discord,
	"dizzyMeh":                     DizzyMeh,
	"dna":                          Dna,
	"docker":                       Docker,
	"documentInfo":                 DocumentInfo,
	"documentLayoutCenter":         DocumentLayoutCenter,
	"documentLayoutLeft":           DocumentLayoutLeft,
	"documentLayoutRight":          DocumentLayoutRight,
	"dollarAlt":                    DollarAlt,
	"dollarSign":                   DollarSign,
	"dollarSignAlt":                DollarSignAlt,
	"downloadAlt":                  DownloadAlt,
	"draggabledots":                Draggabledots,
	"dribbble":                     Dribbble,
	"drill":                        Drill,
	"dropbox":                      Dropbox,
	"dumbbell":                     Dumbbell,
	"ear":                          Ear,
	"edit":                         Edit,
	"editAlt":                      EditAlt,
	"eighteenPlus":                 EighteenPlus,
	"elipsisDoubleValt":            ElipsisDoubleValt,
	"ellipsisH":                    EllipsisH,
	"ellipsisV":                    EllipsisV,
	"emoji":                        Emoji,
	"englishToChinese":             EnglishToChinese,
	"enter":                        Enter,
	"envelope":                     Envelope,
	"envelopeAdd":                  EnvelopeAdd,
	"envelopeAlt":                  EnvelopeAlt,
	"envelopeBlock":                EnvelopeBlock,
	"envelopeBookmark":             EnvelopeBookmark,
	"envelopeCheck":                EnvelopeCheck,
	"envelopeDownload":             EnvelopeDownload,
	"envelopeDownloadAlt":          EnvelopeDownloadAlt,
	"envelopeEdit":                 EnvelopeEdit,
	"envelopeExclamation":          EnvelopeExclamation,
	"envelopeHeart":                EnvelopeHeart,
	"envelopeInfo":                 EnvelopeInfo,
	"envelopeLock":                 EnvelopeLock,
	"envelopeMinus":                EnvelopeMinus,
	"envelopeOpen":                 EnvelopeOpen,
	"envelopeQuestion":             EnvelopeQuestion,
	"envelopeReceive":              EnvelopeReceive,
	"envelopeRedo":                 EnvelopeRedo,
	"envelopeSearch":               EnvelopeSearch,
	"envelopeSend":                 EnvelopeSend,
	"envelopeShare":                EnvelopeShare,
	"envelopeShield":               EnvelopeShield,
	"envelopeStar":                 EnvelopeStar,
	"envelopeTimes":                EnvelopeTimes,
	"envelopeUpload":               EnvelopeUpload,
	"envelopeUploadAlt":            EnvelopeUploadAlt,
	"envelopes":                    Envelopes,
	"equalCircle":                  EqualCircle,
	"estate":                       Estate,
	"euro":                         Euro,
	"euroCircle":                   EuroCircle,
	"exchange":                     Exchange,
	"exchangeAlt":                  ExchangeAlt,
	"exclamation":                  Exclamation,
	"exclamationCircle":            ExclamationCircle,
	"exclamationOctagon":           ExclamationOctagon,
	"exclamationTriangle":          ExclamationTriangle,
	"exclude":                      Exclude,
	"exit":                         Exit,
	"expandAlt":                    ExpandAlt,
	"expandArrows":                 ExpandArrows,
	"expandArrowsAlt":              ExpandArrowsAlt,
	"expandFromCorner":             ExpandFromCorner,
	"expandLeft":                   ExpandLeft,
	"expandRight":                  ExpandRight,
	"export":                       Export,
	"exposureAlt":                  ExposureAlt,
	"exposureIncrease":             ExposureIncrease,
	"externalLinkAlt":              ExternalLinkAlt,
	"eye":                          Eye,
	"eyeSlash":                     EyeSlash,
	"facebook":                     Facebook,
	"facebookF":                    FacebookF,
	"facebookMessenger":            FacebookMessenger,
	"facebookMessengerAlt":         FacebookMessengerAlt,
	"fahrenheit":                   Fahrenheit,
	"fastMail":                     FastMail,
	"fastMailAlt":                  FastMailAlt,
	"favorite":                     Favorite,
	"feedback":                     Feedback,
	"fidgetSpinner":                FidgetSpinner,
	"file":                         File,
	"fileAlt":                      FileAlt,
	"fileBlank":                    FileBlank,
	"fileBlockAlt":                 FileBlockAlt,
	"fileBookmarkAlt":              FileBookmarkAlt,
	"fileCheck":                    FileCheck,
	"fileCheckAlt":                 FileCheckAlt,
	"fileContract":                 FileContract,
	"fileContractDollar":           FileContractDollar,
	"fileCopyAlt":                  FileCopyAlt,
	"fileDownload":                 FileDownload,
	"fileDownloadAlt":              FileDownloadAlt,
	"fileEditAlt":                  FileEditAlt,
	"fileExclamation":              FileExclamation,
	"fileExclamationAlt":           FileExclamationAlt,
	"fileExport":                   FileExport,
	"fileGraph":                    FileGraph,
	"fileHeart":                    FileHeart,
	"fileImport":                   FileImport,
	"fileInfoAlt":                  FileInfoAlt,
	"fileLandscape":                FileLandscape,
	"fileLandscapeAlt":             FileLandscapeAlt,
	"fileLanscapeSlash":            FileLanscapeSlash,
	"fileLockAlt":                  FileLockAlt,
	"fileMedical":                  FileMedical,
	"fileMedicalAlt":               FileMedicalAlt,
	"fileMinus":                    FileMinus,
	"fileMinusAlt":                 FileMinusAlt,
	"fileNetwork":                  FileNetwork,
	"filePlus":                     FilePlus,
	"filePlusAlt":                  FilePlusAlt,
	"fileQuestion":                 FileQuestion,
	"fileQuestionAlt":              FileQuestionAlt,
	"fileRedoAlt":                  FileRedoAlt,
	"fileSearchAlt":                FileSearchAlt,
	"fileShareAlt":                 FileShareAlt,
	"fileShieldAlt":                FileShieldAlt,
	"fileSlash":                    FileSlash,
	"fileTimes":                    FileTimes,
	"fileTimesAlt":                 FileTimesAlt,
	"fileUpload":                   FileUpload,
	"fileUploadAlt":                FileUploadAlt,
	"filesLandscapes":              FilesLandscapes,
	"filesLandscapesAlt":           FilesLandscapesAlt,
	"film":                         Film,
	"filter":                       Filter,
	"filterSlash":                  FilterSlash,
	"fire":                         Fire,
	"fiveHundredPx":                FiveHundredPx,
	"flask":                        Flask,
	"flaskPotion":                  FlaskPotion,
	"flipH":                        FlipH,
	"flipHalt":                     FlipHalt,
	"flipV":                        FlipV,
	"flipValt":                     FlipValt,
	"flower":                       Flower,
	"focus":                        Focus,
	"focusAdd":                     FocusAdd,
	"focusTarget":                  FocusTarget,
	"folder":                       Folder,
	"folderCheck":                  FolderCheck,
	"folderDownload":               FolderDownload,
	"folderExclamation":            FolderExclamation,
	"folderHeart":                  FolderHeart,
	"folderInfo":                   FolderInfo,
	"folderLock":                   FolderLock,
	"folderMedical":                FolderMedical,
	"folderMinus":                  FolderMinus,
	"folderNetwork":                FolderNetwork,
	"folderOpen":                   FolderOpen,
	"folderPlus":                   FolderPlus,
	"folderQuestion":               FolderQuestion,
	"folderSlash":                  FolderSlash,
	"folderTimes":                  FolderTimes,
	"folderUpload":                 FolderUpload,
	"font":                         Font,
	"football":                     Football,
	"footballAmerican":             FootballAmerican,
	"footballBall":                 FootballBall,
	"forecastcloudMoonTear":        ForecastcloudMoonTear,
	"forwadedCall":                 ForwadedCall,
	"forward":                      Forward,
	"frown":                        Frown,
	"gameStructure":                GameStructure,
	"gift":                         Gift,
	"github":                       Github,
	"githubAlt":                    GithubAlt,
	"gitlab":                       Gitlab,
	"glass":                        Glass,
	"glassMartini":                 GlassMartini,
	"glassMartiniAlt":              GlassMartiniAlt,
	"glassMartiniAltSlash":         GlassMartiniAltSlash,
	"glassTea":                     GlassTea,
	"globe":                        Globe,
	"gold":                         Gold,
	"golfBall":                     GolfBall,
	"google":                       Google,
	"googleDrive":                  GoogleDrive,
	"googleDriveAlt":               GoogleDriveAlt,
	"googleHangouts":               GoogleHangouts,
	"googleHangoutsAlt":            GoogleHangoutsAlt,
	"googlePlay":                   GooglePlay,
	"graduationCap":                GraduationCap,
	"graphBar":                     GraphBar,
	"grid":                         Grid,
	"grids":                        Grids,
	"grin":                         Grin,
	"grinTongueWink":               GrinTongueWink,
	"grinTongueWinkAlt":            GrinTongueWinkAlt,
	"gripHorizontalLine":           GripHorizontalLine,
	"hardHat":                      HardHat,
	"hdd":                          Hdd,
	"headSide":                     HeadSide,
	"headSideCough":                HeadSideCough,
	"headSideMask":                 HeadSideMask,
	"headphoneSlash":               HeadphoneSlash,
	"headphones":                   Headphones,
	"headphonesAlt":                HeadphonesAlt,
	"heart":                        Heart,
	"heartAlt":                     HeartAlt,
	"heartBreak":                   HeartBreak,
	"heartMedical":                 HeartMedical,
	"heartRate":                    HeartRate,
	"heartSign":                    HeartSign,
	"heartbeat":                    Heartbeat,
	"hindiToChinese":               HindiToChinese,
	"hipchat":                      Hipchat,
	"history":                      History,
	"historyAlt":                   HistoryAlt,
	"home":                         Home,
	"homeAlt":                      HomeAlt,
	"horizontalAlignCenter":        HorizontalAlignCenter,
	"horizontalAlignLeft":          HorizontalAlignLeft,
	"horizontalAlignRight":         HorizontalAlignRight,
	"horizontalDistributionCenter": HorizontalDistributionCenter,
	"horizontalDistributionLeft":   HorizontalDistributionLeft,
	"horizontalDistributionRight":  HorizontalDistributionRight,
	"hospital":                     Hospital,
	"hospitalSquareSign":           HospitalSquareSign,
	"hospitalSymbol":               HospitalSymbol,
	"hourglass":                    Hourglass,
	"houseUser":                    HouseUser,
	"htmlFive":                     HtmlFive,
	"htmlFiveAlt":                  HtmlFiveAlt,
	"htmlThree":                    HtmlThree,
	"htmlThreeAlt":                 HtmlThreeAlt,
	"hunting":                      Hunting,
	"icons":                        Icons,
	"illustration":                 Illustration,
	"image":                        Image,
	"imageAltSlash":                ImageAltSlash,
	"imageBlock":                   ImageBlock,
	"imageBroken":                  ImageBroken,
	"imageCheck":                   ImageCheck,
	"imageDownload":                ImageDownload,
	"imageEdit":                    ImageEdit,
	"imageLock":                    ImageLock,
	"imageMinus":                   ImageMinus,
	"imagePlus":                    ImagePlus,
	"imageQuestion":                ImageQuestion,
	"imageRedo":                    ImageRedo,
	"imageResizeLandscape":         ImageResizeLandscape,
	"imageResizeSquare":            ImageResizeSquare,
	"imageSearch":                  ImageSearch,
	"imageShare":                   ImageShare,
	"imageShield":                  ImageShield,
	"imageSlash":                   ImageSlash,
	"imageTimes":                   ImageTimes,
	"imageUpload":                  ImageUpload,
	"imageV":                       ImageV,
	"images":                       Images,
	"import":                       Import,
	"inbox":                        Inbox,
	"incomingCall":                 IncomingCall,
	"info":                         Info,
	"infoCircle":                   InfoCircle,
	"instagram":                    Instagram,
	"instagramAlt":                 InstagramAlt,
	"intercom":                     Intercom,
	"intercomAlt":                  IntercomAlt,
	"invoice":                      Invoice,
	"italic":                       Italic,
	"jackhammer":                   Jackhammer,
	"javaScript":                   JavaScript,
	"kayak":                        Kayak,
	"keySkeleton":                  KeySkeleton,
	"keySkeletonAlt":               KeySkeletonAlt,
	"keyboard":                     Keyboard,
	"keyboardAlt":                  KeyboardAlt,
	"keyboardHide":                 KeyboardHide,
	"keyboardShow":                 KeyboardShow,
	"keyholeCircle":                KeyholeCircle,
	"keyholeSquare":                KeyholeSquare,
	"keyholeSquareFull":            KeyholeSquareFull,
	"kid":                          Kid,
	"label":                        Label,
	"labelAlt":                     LabelAlt,
	"lamp":                         Lamp,
	"language":                     Language,
	"laptop":                       Laptop,
	"laptopCloud":                  LaptopCloud,
	"laptopConnection":             LaptopConnection,
	"laughing":                     Laughing,
	"layerGroup":                   LayerGroup,
	"layerGroupSlash":              LayerGroupSlash,
	"layers":                       Layers,
	"layersAlt":                    LayersAlt,
	"layersSlash":                  LayersSlash,
	"leftArrowFromLeft":            LeftArrowFromLeft,
	"leftArrowToLeft":              LeftArrowToLeft,
	"leftIndent":                   LeftIndent,
	"leftIndentAlt":                LeftIndentAlt,
	"leftToRightTextDirection":     LeftToRightTextDirection,
	"letterChineseA":               LetterChineseA,
	"letterEnglishA":               LetterEnglishA,
	"letterHindiA":                 LetterHindiA,
	"letterJapaneseA":              LetterJapaneseA,
	"lifeRing":                     LifeRing,
	"lightbulb":                    Lightbulb,
	"lightbulbAlt":                 LightbulbAlt,
	"line":                         Line,
	"lineAlt":                      LineAlt,
	"lineSpacing":                  LineSpacing,
	"link":                         Link,
	"linkAdd":                      LinkAdd,
	"linkAlt":                      LinkAlt,
	"linkBroken":                   LinkBroken,
	"linkH":                        LinkH,
	"linkedin":                     Linkedin,
	"linkedinAlt":                  LinkedinAlt,
	"linux":                        Linux,
	"liraSign":                     LiraSign,
	"listOl":                       ListOl,
	"listOlAlt":                    ListOlAlt,
	"listUiAlt":                    ListUiAlt,
	"listUl":                       ListUl,
	"locationArrow":                LocationArrow,
	"locationArrowAlt":             LocationArrowAlt,
	"locationPinAlt":               LocationPinAlt,
	"locationPoint":                LocationPoint,
	"lock":                         Lock,
	"lockAccess":                   LockAccess,
	"lockAlt":                      LockAlt,
	"lockOpenAlt":                  LockOpenAlt,
	"lockSlash":                    LockSlash,
	"lottiefiles":                  Lottiefiles,
	"lottiefilesAlt":               LottiefilesAlt,
	"luggageCart":                  LuggageCart,
	"mailbox":                      Mailbox,
	"mailboxAlt":                   MailboxAlt,
	"map":                          Map,
	"mapMarker":                    MapMarker,
	"mapMarkerAlt":                 MapMarkerAlt,
	"mapMarkerEdit":                MapMarkerEdit,
	"mapMarkerInfo":                MapMarkerInfo,
	"mapMarkerMinus":               MapMarkerMinus,
	"mapMarkerPlus":                MapMarkerPlus,
	"mapMarkerQuestion":            MapMarkerQuestion,
	"mapMarkerShield":              MapMarkerShield,
	"mapMarkerSlash":               MapMarkerSlash,
	"mapPin":                       MapPin,
	"mapPinAlt":                    MapPinAlt,
	"mars":                         Mars,
	"masterCard":                   MasterCard,
	"maximizeLeft":                 MaximizeLeft,
	"medal":                        Medal,
	"medicalDrip":                  MedicalDrip,
	"medicalSquare":                MedicalSquare,
	"medicalSquareFull":            MedicalSquareFull,
	"mediumM":                      MediumM,
	"medkit":                       Medkit,
	"meetingBoard":                 MeetingBoard,
	"megaphone":                    Megaphone,
	"meh":                          Meh,
	"mehAlt":                       MehAlt,
	"mehClosedEye":                 MehClosedEye,
	"message":                      Message,
	"metro":                        Metro,
	"microphone":                   Microphone,
	"microphoneSlash":              MicrophoneSlash,
	"microscope":                   Microscope,
	"microsoft":                    Microsoft,
	"minus":                        Minus,
	"minusCircle":                  MinusCircle,
	"minusPath":                    MinusPath,
	"minusSquare":                  MinusSquare,
	"minusSquareFull":              MinusSquareFull,
	"missedCall":                   MissedCall,
	"mobileAndroid":                MobileAndroid,
	"mobileAndroidAlt":             MobileAndroidAlt,
	"mobileVibrate":                MobileVibrate,
	"modem":                        Modem,
	"moneyBill":                    MoneyBill,
	"moneyBillSlash":               MoneyBillSlash,
	"moneyBillStack":               MoneyBillStack,
	"moneyInsert":                  MoneyInsert,
	"moneyStack":                   MoneyStack,
	"moneyWithdraw":                MoneyWithdraw,
	"moneyWithdrawal":              MoneyWithdrawal,
	"moneybag":                     Moneybag,
	"moneybagAlt":                  MoneybagAlt,
	"monitor":                      Monitor,
	"monitorHeartRate":             MonitorHeartRate,
	"moon":                         Moon,
	"moonEclipse":                  MoonEclipse,
	"moonset":                      Moonset,
	"mountains":                    Mountains,
	"mountainsSun":                 MountainsSun,
	"mouse":                        Mouse,
	"mouseAlt":                     MouseAlt,
	"mouseAltTwo":                  MouseAltTwo,
	"multiply":                     Multiply,
	"music":                        Music,
	"musicNote":                    MusicNote,
	"musicTuneSlash":               MusicTuneSlash,
	"na":                           Na,
	"navigator":                    Navigator,
	"nerd":                         Nerd,
	"newspaper":                    Newspaper,
	"ninja":                        Ninja,
	"noEntry":                      NoEntry,
	"notebooks":                    Notebooks,
	"notes":                        Notes,
	"objectGroup":                  ObjectGroup,
	"objectUngroup":                ObjectUngroup,
	"octagon":                      Octagon,
	"okta":                         Okta,
	"opera":                        Opera,
	"operaAlt":                     OperaAlt,
	"outgoingCall":                 OutgoingCall,
	"package":                      Package,
	"padlock":                      Padlock,
	"pagelines":                    Pagelines,
	"pagerduty":                    Pagerduty,
	"paintTool":                    PaintTool,
	"palette":                      Palette,
	"panelAdd":                     PanelAdd,
	"panoramaH":                    PanoramaH,
	"panoramaHalt":                 PanoramaHalt,
	"panoramaV":                    PanoramaV,
	"paperclip":                    Paperclip,
	"paragraph":                    Paragraph,
	"parcel":                       Parcel,
	"parkingCircle":                ParkingCircle,
	"parkingSquare":                ParkingSquare,
	"pathfinder":                   Pathfinder,
	"pathfinderUnite":              PathfinderUnite,
	"pause":                        Pause,
	"pauseCircle":                  PauseCircle,
	"paypal":                       Paypal,
	"pen":                          Pen,
	"pentagon":                     Pentagon,
	"percentage":                   Percentage,
	"phone":                        Phone,
	"phoneAlt":                     PhoneAlt,
	"phonePause":                   PhonePause,
	"phoneSlash":                   PhoneSlash,
	"phoneTimes":                   PhoneTimes,
	"phoneVolume":                  PhoneVolume,
	"picture":                      Picture,
	"pizzaSlice":                   PizzaSlice,
	"plane":                        Plane,
	"planeArrival":                 PlaneArrival,
	"planeDeparture":               PlaneDeparture,
	"planeFly":                     PlaneFly,
	"play":                         Play,
	"playCircle":                   PlayCircle,
	"plug":                         Plug,
	"plus":                         Plus,
	"plusCircle":                   PlusCircle,
	"plusSquare":                   PlusSquare,
	"podium":                       Podium,
	"polygon":                      Polygon,
	"postStamp":                    PostStamp,
	"postcard":                     Postcard,
	"pound":                        Pound,
	"poundCircle":                  PoundCircle,
	"power":                        Power,
	"prescriptionBottle":           PrescriptionBottle,
	"presentation":                 Presentation,
	"presentationCheck":            PresentationCheck,
	"presentationEdit":             PresentationEdit,
	"presentationLine":             PresentationLine,
	"presentationLinesAlt":         PresentationLinesAlt,
	"presentationMinus":            PresentationMinus,
	"presentationPlay":             PresentationPlay,
	"presentationPlus":             PresentationPlus,
	"presentationTimes":            PresentationTimes,
	"previous":                     Previous,
	"pricetagAlt":                  PricetagAlt,
	"print":                        Print,
	"printSlash":                   PrintSlash,
	"process":                      Process,
	"processor":                    Processor,
	"programmingLanguage":          ProgrammingLanguage,
	"pump":                         Pump,
	"puzzlePiece":                  PuzzlePiece,
	"qrcodeScan":                   QrcodeScan,
	"question":                     Question,
	"questionCircle":               QuestionCircle,
	"rainbow":                      Rainbow,
	"raindrops":                    Raindrops,
	"raindropsAlt":                 RaindropsAlt,
	"react":                        React,
	"receipt":                      Receipt,
	"receiptAlt":                   ReceiptAlt,
	"recordAudio":                  RecordAudio,
	"redditAlienAlt":               RedditAlienAlt,
	"redo":                         Redo,
	"refresh":                      Refresh,
	"registered":                   Registered,
	"repeat":                       Repeat,
	"restaurant":                   Restaurant,
	"rightIndentAlt":               RightIndentAlt,
	"rightToLeftTextDirection":     RightToLeftTextDirection,
	"robot":                        Robot,
	"rocket":                       Rocket,
	"ropeWay":                      RopeWay,
	"rotateThreeHundredSixty":      RotateThreeHundredSixty,
	"rss":                          Rss,
	"rssAlt":                       RssAlt,
	"rssInterface":                 RssInterface,
	"ruler":                        Ruler,
	"rulerCombined":                RulerCombined,
	"rupeeSign":                    RupeeSign,
	"sad":                          Sad,
	"sadCry":                       SadCry,
	"sadCrying":                    SadCrying,
	"sadDizzy":                     SadDizzy,
	"sadSquint":                    SadSquint,
	"sanitizer":                    Sanitizer,
	"sanitizerAlt":                 SanitizerAlt,
	"save":                         Save,
	"scalingLeft":                  ScalingLeft,
	"scalingRight":                 ScalingRight,
	"scenery":                      Scenery,
	"schedule":                     Schedule,
	"screw":                        Screw,
	"scroll":                       Scroll,
	"scrollH":                      ScrollH,
	"search":                       Search,
	"searchAlt":                    SearchAlt,
	"searchMinus":                  SearchMinus,
	"searchPlus":                   SearchPlus,
	"selfie":                       Selfie,
	"server":                       Server,
	"serverAlt":                    ServerAlt,
	"serverConnection":             ServerConnection,
	"serverNetwork":                ServerNetwork,
	"serverNetworkAlt":             ServerNetworkAlt,
	"servers":                      Servers,
	"servicemark":                  Servicemark,
	"setting":                      Setting,
	"seventeenPlus":                SeventeenPlus,
	"share":                        Share,
	"shareAlt":                     ShareAlt,
	"shield":                       Shield,
	"shieldCheck":                  ShieldCheck,
	"shieldExclamation":            ShieldExclamation,
	"shieldPlus":                   ShieldPlus,
	"shieldQuestion":               ShieldQuestion,
	"shieldSlash":                  ShieldSlash,
	"ship":                         Ship,
	"shop":                         Shop,
	"shoppingBag":                  ShoppingBag,
	"shoppingBasket":               ShoppingBasket,
	"shoppingCart":                 ShoppingCart,
	"shoppingCartAlt":              ShoppingCartAlt,
	"shovel":                       Shovel,
	"shrink":                       Shrink,
	"shuffle":                      Shuffle,
	"shutter":                      Shutter,
	"shutterAlt":                   ShutterAlt,
	"sick":                         Sick,
	"sigma":                        Sigma,
	"signAlt":                      SignAlt,
	"signInAlt":                    SignInAlt,
	"signLeft":                     SignLeft,
	"signOutAlt":                   SignOutAlt,
	"signRight":                    SignRight,
	"signal":                       Signal,
	"signalAlt":                    SignalAlt,
	"signalAltThree":               SignalAltThree,
	"signin":                       Signin,
	"signout":                      Signout,
	"silence":                      Silence,
	"silentSquint":                 SilentSquint,
	"simCard":                      SimCard,
	"sitemap":                      Sitemap,
	"sixPlus":                      SixPlus,
	"sixteenPlus":                  SixteenPlus,
	"skipForward":                  SkipForward,
	"skipForwardAlt":               SkipForwardAlt,
	"skipForwardCircle":            SkipForwardCircle,
	"skype":                        Skype,
	"skypeAlt":                     SkypeAlt,
	"slack":                        Slack,
	"slackAlt":                     SlackAlt,
	"sliderH":                      SliderH,
	"sliderHrange":                 SliderHrange,
	"slidersV":                     SlidersV,
	"slidersValt":                  SlidersValt,
	"smile":                        Smile,
	"smileBeam":                    SmileBeam,
	"smileDizzy":                   SmileDizzy,
	"smileSquintWink":              SmileSquintWink,
	"smileSquintWinkAlt":           SmileSquintWinkAlt,
	"smileWink":                    SmileWink,
	"smileWinkAlt":                 SmileWinkAlt,
	"snapchatAlt":                  SnapchatAlt,
	"snapchatGhost":                SnapchatGhost,
	"snapchatSquare":               SnapchatSquare,
	"snowFlake":                    SnowFlake,
	"snowflake":                    Snowflake,
	"snowflakeAlt":                 SnowflakeAlt,
	"socialDistancing":             SocialDistancing,
	"sort":                         Sort,
	"sortAmountDown":               SortAmountDown,
	"sortAmountUp":                 SortAmountUp,
	"sorting":                      Sorting,
	"spaceKey":                     SpaceKey,
	"spade":                        Spade,
	"sperms":                       Sperms,
	"spin":                         Spin,
	"spinner":                      Spinner,
	"spinnerAlt":                   SpinnerAlt,
	"square":                       Square,
	"squareFull":                   SquareFull,
	"squareShape":                  SquareShape,
	"squint":                       Squint,
	"star":                         Star,
	"starHalfAlt":                  StarHalfAlt,
	"stepBackward":                 StepBackward,
	"stepBackwardAlt":              StepBackwardAlt,
	"stepBackwardCircle":           StepBackwardCircle,
	"stepForward":                  StepForward,
	"stethoscope":                  Stethoscope,
	"stethoscopeAlt":               StethoscopeAlt,
	"stopCircle":                   StopCircle,
	"stopwatch":                    Stopwatch,
	"stopwatchSlash":               StopwatchSlash,
	"store":                        Store,
	"storeAlt":                     StoreAlt,
	"storeSlash":                   StoreSlash,
	"streering":                    Streering,
	"stretcher":                    Stretcher,
	"subject":                      Subject,
	"subway":                       Subway,
	"subwayAlt":                    SubwayAlt,
	"suitcase":                     Suitcase,
	"suitcaseAlt":                  SuitcaseAlt,
	"sun":                          Sun,
	"sunset":                       Sunset,
	"surprise":                     Surprise,
	"swatchbook":                   Swatchbook,
	"swiggy":                       Swiggy,
	"swimmer":                      Swimmer,
	"sync":                         Sync,
	"syncExclamation":              SyncExclamation,
	"syncSlash":                    SyncSlash,
	"syringe":                      Syringe,
	"table":                        Table,
	"tableTennis":                  TableTennis,
	"tablet":                       Tablet,
	"tablets":                      Tablets,
	"tachometerFast":               TachometerFast,
	"tachometerFastAlt":            TachometerFastAlt,
	"tag":                          Tag,
	"tagAlt":                       TagAlt,
	"tape":                         Tape,
	"taxi":                         Taxi,
	"tear":                         Tear,
	"telegram":                     Telegram,
	"telegramAlt":                  TelegramAlt,
	"telescope":                    Telescope,
	"temperature":                  Temperature,
	"temperatureEmpty":             TemperatureEmpty,
	"temperatureHalf":              TemperatureHalf,
	"temperatureMinus":             TemperatureMinus,
	"temperaturePlus":              TemperaturePlus,
	"temperatureQuarter":           TemperatureQuarter,
	"temperatureThreeQuarter":      TemperatureThreeQuarter,
	"tenPlus":                      TenPlus,
	"tennisBall":                   TennisBall,
	"textFields":                   TextFields,
	"textIcon":                     TextIcon,
	"textSize":                     TextSize,
	"textStrikeThrough":            TextStrikeThrough,
	"th":                           Th,
	"thLarge":                      ThLarge,
	"thSlash":                      ThSlash,
	"thermometer":                  Thermometer,
	"thirteenPlus":                 ThirteenPlus,
	"threePlus":                    ThreePlus,
	"thumbsDown":                   ThumbsDown,
	"thumbsUp":                     ThumbsUp,
	"thunderstorm":                 Thunderstorm,
	"thunderstormMoon":             ThunderstormMoon,
	"thunderstormSun":              ThunderstormSun,
	"ticket":                       Ticket,
	"times":                        Times,
	"timesCircle":                  TimesCircle,
	"timesSquare":                  TimesSquare,
	"toggleOff":                    ToggleOff,
	"toggleOn":                     ToggleOn,
	"toiletPaper":                  ToiletPaper,
	"topArrowFromTop":              TopArrowFromTop,
	"topArrowToTop":                TopArrowToTop,
	"tornado":                      Tornado,
	"trademark":                    Trademark,
	"trademarkCircle":              TrademarkCircle,
	"trafficBarrier":               TrafficBarrier,
	"trafficLight":                 TrafficLight,
	"transaction":                  Transaction,
	"trash":                        Trash,
	"trashAlt":                     TrashAlt,
	"trees":                        Trees,
	"triangle":                     Triangle,
	"trophy":                       Trophy,
	"trowel":                       Trowel,
	"truck":                        Truck,
	"truckLoading":                 TruckLoading,
	"tumblr":                       Tumblr,
	"tumblrAlt":                    TumblrAlt,
	"tumblrSquare":                 TumblrSquare,
	"tvRetro":                      TvRetro,
	"tvRetroSlash":                 TvRetroSlash,
	"twelvePlus":                   TwelvePlus,
	"twentyOnePlus":                TwentyOnePlus,
	"twitter":                      Twitter,
	"twitterAlt":                   TwitterAlt,
	"umbrella":                     Umbrella,
	"unamused":                     Unamused,
	"underline":                    Underline,
	"university":                   University,
	"unlock":                       Unlock,
	"unlockAlt":                    UnlockAlt,
	"upload":                       Upload,
	"uploadAlt":                    UploadAlt,
	"usdCircle":                    UsdCircle,
	"usdSquare":                    UsdSquare,
	"user":                         User,
	"userArrows":                   UserArrows,
	"userCheck":                    UserCheck,
	"userCircle":                   UserCircle,
	"userExclamation":              UserExclamation,
	"userLocation":                 UserLocation,
	"userMd":                       UserMd,
	"userMinus":                    UserMinus,
	"userNurse":                    UserNurse,
	"userPlus":                     UserPlus,
	"userSquare":                   UserSquare,
	"userTimes":                    UserTimes,
	"usersAlt":                     UsersAlt,
	"utensils":                     Utensils,
	"utensilsAlt":                  UtensilsAlt,
	"vectorSquare":                 VectorSquare,
	"vectorSquareAlt":              VectorSquareAlt,
	"venus":                        Venus,
	"verticalAlignBottom":          VerticalAlignBottom,
	"verticalAlignCenter":          VerticalAlignCenter,
	"verticalAlignTop":             VerticalAlignTop,
	"verticalDistributeBottom":     VerticalDistributeBottom,
	"verticalDistributionCenter":   VerticalDistributionCenter,
	"verticalDistributionTop":      VerticalDistributionTop,
	"video":                        Video,
	"videoQuestion":                VideoQuestion,
	"videoSlash":                   VideoSlash,
	"virusSlash":                   VirusSlash,
	"visualStudio":                 VisualStudio,
	"vk":                           Vk,
	"vkAlt":                        VkAlt,
	"voicemail":                    Voicemail,
	"voicemailRectangle":           VoicemailRectangle,
	"volleyball":                   Volleyball,
	"volume":                       Volume,
	"volumeDown":                   VolumeDown,
	"volumeMute":                   VolumeMute,
	"volumeOff":                    VolumeOff,
	"volumeUp":                     VolumeUp,
	"vuejs":                        Vuejs,
	"vuejsAlt":                     VuejsAlt,
	"wall":                         Wall,
	"wallet":                       Wallet,
	"watch":                        Watch,
	"watchAlt":                     WatchAlt,
	"water":                        Water,
	"waterDropSlash":               WaterDropSlash,
	"waterGlass":                   WaterGlass,
	"webGrid":                      WebGrid,
	"webGridAlt":                   WebGridAlt,
	"webSection":                   WebSection,
	"webSectionAlt":                WebSectionAlt,
	"webcam":                       Webcam,
	"weight":                       Weight,
	"whatsapp":                     Whatsapp,
	"whatsappAlt":                  WhatsappAlt,
	"wheelBarrow":                  WheelBarrow,
	"wheelchair":                   Wheelchair,
	"wheelchairAlt":                WheelchairAlt,
	"wifi":                         Wifi,
	"wifiRouter":                   WifiRouter,
	"wifiSlash":                    WifiSlash,
	"wind":                         Wind,
	"windMoon":                     WindMoon,
	"windSun":                      WindSun,
	"window":                       Window,
	"windowGrid":                   WindowGrid,
	"windowMaximize":               WindowMaximize,
	"windowSection":                WindowSection,
	"windows":                      Windows,
	"windsock":                     Windsock,
	"windy":                        Windy,
	"wordpress":                    Wordpress,
	"wordpressSimple":              WordpressSimple,
	"wrapText":                     WrapText,
	"wrench":                       Wrench,
	"x":                            X,
	"xadd":                         Xadd,
	"yen":                          Yen,
	"yenCircle":                    YenCircle,
	"yinYang":                      YinYang,
	"youtube":                      Youtube,
	"youtubeAlt":                   YoutubeAlt,
	"zeroPlus":                     ZeroPlus,
}

func Abacus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2a1 1 0 0 0-1 1v3h-4V5a1 1 0 0 0-2 0v1h-2V5a1 1 0 0 0-2 0v1H8V5a1 1 0 0 0-2 0v1H4V3a1 1 0 0 0-2 0v16a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V3a1 1 0 0 0-1-1m-1 17a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-3h2v1a1 1 0 0 0 2 0v-1h2v1a1 1 0 0 0 2 0v-1h4v1a1 1 0 0 0 2 0v-1h2Zm0-5h-2v-1a1 1 0 0 0-2 0v1h-4v-1a1 1 0 0 0-2 0v1H8v-1a1 1 0 0 0-2 0v1H4V8h2v1a1 1 0 0 0 2 0V8h2v1a1 1 0 0 0 2 0V8h2v1a1 1 0 0 0 2 0V8h4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AccessibleIconAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 7a2 2 0 1 0-2-2a2 2 0 0 0 2 2m-3.3 11.4A4 4 0 1 1 9.05 12a1 1 0 1 0-.22-2a6 6 0 0 0 .67 12a6 6 0 0 0 4.8-2.4a1 1 0 0 0-1.6-1.2m6.8 1.6h-1v-5a1 1 0 0 0-1-1h-4.57l1.69-4.66v-.31a1.1 1.1 0 0 0 0-.18a1.06 1.06 0 0 0 0-.19a1.4 1.4 0 0 0-.09-.17a.72.72 0 0 0-.11-.15a.64.64 0 0 0-.15-.13s0-.06-.08-.08L9.71 5.55h-.12a1.06 1.06 0 0 0-.19-.06H9a.8.8 0 0 0-.2.07h-.11L6 7.13A1 1 0 0 0 6.48 9A1 1 0 0 0 7 8.87l2.23-1.3l3.24 1.88l-1.89 5.21a.88.88 0 0 0 0 .16a.58.58 0 0 0 0 .18a3 3 0 0 0 .08.38l.11.15a.57.57 0 0 0 .11.16a.67.67 0 0 0 .14.09a1.22 1.22 0 0 0 .19.12a1 1 0 0 0 .34.06h5v5a1 1 0 0 0 1 1h2a1 1 0 0 0 0-2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Adjust(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m-1 17.93a8 8 0 0 1 0-15.86Zm2 0V4.07a8 8 0 0 1 0 15.86"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdjustAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.83 7.32a.2.2 0 0 0 0-.08a10 10 0 0 0-3.38-3.65A9.89 9.89 0 0 0 12 2a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h.28a10 10 0 0 0 8.55-14.68M13 4.06a8 8 0 0 1 2.49.74L13 9.12Zm0 9.06l4.17-7.22a7.89 7.89 0 0 1 1.58 1.83L13 17.69Zm1.16 6.57L19.75 10a8.36 8.36 0 0 1 .25 2a7.94 7.94 0 0 1-5.84 7.69"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdjustCircle(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 7A7.74 7.74 0 1 0 7 17A7.74 7.74 0 1 0 17 7m-1.6 1.6a6.12 6.12 0 0 1 .11 1.14a5.92 5.92 0 0 1-.16 1.34l-2.44-2.42a5.92 5.92 0 0 1 1.34-.16a6.12 6.12 0 0 1 1.14.11Zm-.94 4.4A6 6 0 0 1 13 14.46L9.54 11A6 6 0 0 1 11 9.54Zm-8 1.46a5.75 5.75 0 1 1 8-8h-.25a7.76 7.76 0 0 0-7.71 7.79c0 .08.01.17.01.25Zm2.1.89a6.12 6.12 0 0 1-.11-1.14a5.92 5.92 0 0 1 .16-1.34l2.43 2.43a5.92 5.92 0 0 1-1.34.16a6.12 6.12 0 0 1-1.09-.07ZM14.25 20a5.77 5.77 0 0 1-4.75-2.51h.25a7.76 7.76 0 0 0 7.75-7.74V9.5A5.75 5.75 0 0 1 14.25 20"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdjustHalf(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.83 7.32a10.11 10.11 0 0 0-3.44-3.73A10 10 0 1 0 12 22h.29a10 10 0 0 0 8.54-14.68M11 19.93a8 8 0 0 1 0-15.86Zm2-15.86a8.07 8.07 0 0 1 2.49.74L13 9.12Zm0 9l4.17-7.17a8.14 8.14 0 0 1 1.58 1.83L13 17.69Zm1.15 6.58L19.74 10a8.16 8.16 0 0 1 .26 2a8 8 0 0 1-5.85 7.7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Adobe(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.07 17.28h2.78l1.75 3.44h2.54L12 9.87ZM2 3v18L9.42 3Zm12.48 0L22 20.81V3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdobeAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.247 2.403a.999.999 0 0 0-.83-.444H1.992a1 1 0 0 0-1 1v18.082a1 1 0 0 0 1.925.38L10.342 3.34a1 1 0 0 0-.095-.936m-7.255 13.57V3.96h4.933Zm19-14.014h-7.518a1 1 0 0 0-.921 1.388l7.517 17.85a.999.999 0 0 0 .921.612a1.049 1.049 0 0 0 .198-.02a1.001 1.001 0 0 0 .803-.98V2.959a1 1 0 0 0-1-1m-1 13.899L15.98 3.959h5.012ZM12.949 9.52a1 1 0 0 0-.926-.646h-.01a1 1 0 0 0-.928.627l-3.059 7.631a1 1 0 0 0 .929 1.372h2.254l1.522 2.99a.999.999 0 0 0 .892.547h2.612a1 1 0 0 0 .936-1.353Zm1.287 10.521l-1.523-2.99a.999.999 0 0 0-.89-.547h-1.39l1.553-3.875l2.802 7.412Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Airplay(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.83 13.45a1 1 0 0 0-1.66 0l-4 6a1 1 0 0 0 0 1A1 1 0 0 0 8 21h8a1 1 0 0 0 .88-.53a1 1 0 0 0-.05-1ZM9.87 19L12 15.8l2.13 3.2ZM19 3H5a3 3 0 0 0-3 3v9a3 3 0 0 0 3 3h.85a1 1 0 1 0 0-2H5a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1h-.85a1 1 0 0 0 0 2H19a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Align(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 15h10a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2m0-4h10a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2m0-4h18a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2m18 10H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2m-1.36-7.43a1 1 0 1 0-1.28 1.53l1.08.9l-1.08.9a1 1 0 0 0-.13 1.41a1 1 0 0 0 .77.36a1 1 0 0 0 .64-.24l2-1.66a1 1 0 0 0 0-1.54Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 11H3a1 1 0 0 0 0 2h7a1 1 0 0 0 0-2m0 8H3a1 1 0 0 0 0 2h7a1 1 0 0 0 0-2m4-14h7a1 1 0 0 0 0-2h-7a1 1 0 0 0 0 2m-4 2H3a1 1 0 0 0 0 2h7a1 1 0 0 0 0-2m0 8H3a1 1 0 0 0 0 2h7a1 1 0 0 0 0-2m0-12H7a1 1 0 0 0 0 2h3a1 1 0 0 0 0-2m11 4h-7a1 1 0 0 0 0 2h7a1 1 0 0 0 0-2m0 4h-7a1 1 0 0 0 0 2h7a1 1 0 0 0 0-2m-4 8h-3a1 1 0 0 0 0 2h3a1 1 0 0 0 0-2m4-4h-7a1 1 0 0 0 0 2h7a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignCenter(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 7h18a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2m4 2a1 1 0 0 0 0 2h10a1 1 0 0 0 0-2Zm14 4H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2m-4 4H7a1 1 0 0 0 0 2h10a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignCenterAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 8h14a1 1 0 0 0 0-2H5a1 1 0 0 0 0 2m16 3H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2m-2 5H5a1 1 0 0 0 0 2h14a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignCenterH(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.92 11.62a1 1 0 0 0-.21-.33l-2.5-2.5a1 1 0 0 0-1.42 1.42l.8.79H16a1 1 0 0 0 0 2h2.59l-.8.79a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l2.5-2.5a1 1 0 0 0 .21-.33a1 1 0 0 0 0-.76M8 11H5.41l.8-.79a1 1 0 0 0-1.42-1.42l-2.5 2.5a1 1 0 0 0-.21.33a1 1 0 0 0 0 .76a1 1 0 0 0 .21.33l2.5 2.5a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-.8-.79H8a1 1 0 0 0 0-2m4-4a1 1 0 0 0-1 1v8a1 1 0 0 0 2 0V8a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignCenterJustify(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 15H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2M3 5h18a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2m14 14H7a1 1 0 0 0 0 2h10a1 1 0 0 0 0-2m4-12H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2m0 4H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignCenterV(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.21 6.21l.79-.8V8a1 1 0 0 0 2 0V5.41l.79.8a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-2.5-2.5a1 1 0 0 0-.33-.21a1 1 0 0 0-.76 0a1 1 0 0 0-.33.21l-2.5 2.5a1 1 0 0 0 1.42 1.42M16 11H8a1 1 0 0 0 0 2h8a1 1 0 0 0 0-2m-2.21 6.79l-.79.8V16a1 1 0 0 0-2 0v2.59l-.79-.8a1 1 0 0 0-1.42 1.42l2.5 2.5a1 1 0 0 0 .33.21a.94.94 0 0 0 .76 0a1 1 0 0 0 .33-.21l2.5-2.5a1 1 0 0 0-1.42-1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignJustify(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 7h18a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2m18 10H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2m0-4H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2m0-4H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignLeft(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 7h18a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2m0 4h14a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2m18 2H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2m-4 4H3a1 1 0 0 0 0 2h14a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignLeftJustify(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 5h18a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2m12 14H3a1 1 0 0 0 0 2h12a1 1 0 0 0 0-2m6-8H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2m0-4H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2m0 8H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignLetterRight(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 4h11a1 1 0 0 0 0-2H10a1 1 0 0 0 0 2m11 16H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2m0-14H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2m0 10H11a1 1 0 0 0 0 2h10a1 1 0 0 0 0-2m0-6H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignRight(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 7h18a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2m18 10H7a1 1 0 0 0 0 2h14a1 1 0 0 0 0-2m0-8H7a1 1 0 0 0 0 2h14a1 1 0 0 0 0-2m0 4H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignRightJustify(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 5h18a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2m18 14H11a1 1 0 0 0 0 2h10a1 1 0 0 0 0-2m0-8H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2m0 4H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2m0-8H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Amazon(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.04 17.52q.1-.16.32-.02a21.308 21.308 0 0 0 10.88 2.9a21.524 21.524 0 0 0 7.74-1.46q.1-.04.29-.12t.27-.12a.356.356 0 0 1 .47.12q.17.24-.11.44q-.36.26-.92.6a14.99 14.99 0 0 1-3.84 1.58A16.175 16.175 0 0 1 12 22a16.017 16.017 0 0 1-5.9-1.09a16.246 16.246 0 0 1-4.98-3.07a.273.273 0 0 1-.12-.2a.215.215 0 0 1 .04-.12m6.02-5.7a4.036 4.036 0 0 1 .68-2.36A4.197 4.197 0 0 1 9.6 7.98a10.063 10.063 0 0 1 2.66-.66q.54-.06 1.76-.16v-.34a3.562 3.562 0 0 0-.28-1.72a1.5 1.5 0 0 0-1.32-.6h-.16a2.189 2.189 0 0 0-1.14.42a1.64 1.64 0 0 0-.62 1a.508.508 0 0 1-.4.46L7.8 6.1q-.34-.08-.34-.36a.587.587 0 0 1 .02-.14a3.834 3.834 0 0 1 1.67-2.64A6.268 6.268 0 0 1 12.26 2h.5a5.054 5.054 0 0 1 3.56 1.18a3.81 3.81 0 0 1 .37.43a3.875 3.875 0 0 1 .27.41a2.098 2.098 0 0 1 .18.52q.08.34.12.47a2.856 2.856 0 0 1 .06.56q.02.43.02.51v4.84a2.868 2.868 0 0 0 .15.95a2.475 2.475 0 0 0 .29.62q.14.19.46.61a.599.599 0 0 1 .12.32a.346.346 0 0 1-.16.28q-1.66 1.44-1.8 1.56a.557.557 0 0 1-.58.04q-.28-.24-.49-.46t-.3-.32a4.466 4.466 0 0 1-.29-.39q-.2-.29-.28-.39a4.91 4.91 0 0 1-2.2 1.52a6.038 6.038 0 0 1-1.68.2a3.505 3.505 0 0 1-2.53-.95a3.553 3.553 0 0 1-.99-2.69m3.44-.4a1.895 1.895 0 0 0 .39 1.25a1.294 1.294 0 0 0 1.05.47a1.022 1.022 0 0 0 .17-.02a1.022 1.022 0 0 1 .15-.02a2.033 2.033 0 0 0 1.3-1.08a3.13 3.13 0 0 0 .33-.83a3.8 3.8 0 0 0 .12-.73q.01-.28.01-.92v-.5a7.287 7.287 0 0 0-1.76.16a2.144 2.144 0 0 0-1.76 2.22m8.4 6.44a.626.626 0 0 1 .12-.16a3.14 3.14 0 0 1 .96-.46a6.52 6.52 0 0 1 1.48-.22a1.195 1.195 0 0 1 .38.02q.9.08 1.08.3a.655.655 0 0 1 .08.36v.14a4.56 4.56 0 0 1-.38 1.65a3.84 3.84 0 0 1-1.06 1.53a.302.302 0 0 1-.18.08a.177.177 0 0 1-.08-.02q-.12-.06-.06-.22a7.632 7.632 0 0 0 .74-2.42a.513.513 0 0 0-.08-.32q-.2-.24-1.12-.24q-.34 0-.8.04q-.5.06-.92.12a.232.232 0 0 1-.16-.04a.065.065 0 0 1-.02-.08a.153.153 0 0 1 .02-.06"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ambulance(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5 9.5h-1v-1a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2h1v1a1 1 0 0 0 2 0v-1h1a1 1 0 0 0 0-2m12.43 2.15l-.06-.11a.61.61 0 0 0-.07-.14l-2.4-3.2A3 3 0 0 0 18 7h-2V6a3 3 0 0 0-3-3H4a3 3 0 0 0-3 3v11a1 1 0 0 0 1 1h1a3 3 0 0 0 6 0h6a3 3 0 0 0 6 0h1a1 1 0 0 0 1-1v-5a1 1 0 0 0-.07-.35M6 19a1 1 0 1 1 1-1a1 1 0 0 1-1 1m8-3H8.22a3 3 0 0 0-4.44 0H3V6a1 1 0 0 1 1-1h9a1 1 0 0 1 1 1Zm2-7h2a1 1 0 0 1 .8.4L20 11h-4Zm2 10a1 1 0 1 1 1-1a1 1 0 0 1-1 1m3-3h-.78a3 3 0 0 0-4.22-.22V13h5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Analysis(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.71 7.29a1 1 0 0 0-1.42 0L14 13.59l-4.29-4.3a1 1 0 0 0-1.42 0l-6 6a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0L9 11.41l4.29 4.3a1 1 0 0 0 1.42 0l7-7a1 1 0 0 0 0-1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Analytics(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 12a1 1 0 0 0-1 1v8a1 1 0 0 0 2 0v-8a1 1 0 0 0-1-1m5-10a1 1 0 0 0-1 1v18a1 1 0 0 0 2 0V3a1 1 0 0 0-1-1m10 14a1 1 0 0 0-1 1v4a1 1 0 0 0 2 0v-4a1 1 0 0 0-1-1m-5-8a1 1 0 0 0-1 1v12a1 1 0 0 0 2 0V9a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Anchor(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 13h-2a1 1 0 0 0 0 2h.91A6 6 0 0 1 13 19.91V11h1a1 1 0 0 0 0-2h-1V7.82a3 3 0 1 0-2 0V9h-1a1 1 0 0 0 0 2h1v8.91A6 6 0 0 1 6.09 15H7a1 1 0 0 0 0-2H5a1 1 0 0 0-1 1a8 8 0 0 0 16 0a1 1 0 0 0-1-1m-7-7a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Android(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14.975 3.019l.96-1.732a.193.193 0 0 0-.338-.187l-.97 1.75a6.541 6.541 0 0 0-5.253 0l-.97-1.75a.193.193 0 0 0-.34.187l.96 1.732a5.546 5.546 0 0 0-3.092 4.876h12.137a5.546 5.546 0 0 0-3.094-4.876M9.2 5.674a.507.507 0 1 1 .507-.506a.507.507 0 0 1-.507.506m5.602 0a.507.507 0 1 1 .507-.506a.507.507 0 0 1-.507.506M5.93 17.171A1.467 1.467 0 0 0 7.4 18.64h.973v3a1.36 1.36 0 1 0 2.721 0v-3h1.814v3a1.36 1.36 0 1 0 2.72 0v-3h.974a1.467 1.467 0 0 0 1.468-1.468V8.375H5.93Zm-1.867-9.03a1.362 1.362 0 0 0-1.36 1.361v5.669a1.36 1.36 0 1 0 2.72 0V9.502a1.362 1.362 0 0 0-1.36-1.36m15.872 0a1.362 1.362 0 0 0-1.36 1.361v5.669a1.36 1.36 0 1 0 2.72 0v-5.67a1.362 1.362 0 0 0-1.36-1.36"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AndroidAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.992 9a1 1 0 0 0-1 1v4a1 1 0 0 0 2 0v-4a1 1 0 0 0-1-1m18 0a1 1 0 0 0-1 1v4a1 1 0 0 0 2 0v-4a1 1 0 0 0-1-1m-4.135-5.5a1 1 0 1 0-1.731-1l-.614 1.063a5.928 5.928 0 0 0-5.04 0L8.857 2.5a1 1 0 1 0-1.731 1l.692 1.198A5.979 5.979 0 0 0 5.992 9v8a1 1 0 0 0 1 1h2v3a1 1 0 0 0 2 0v-3h2v3a1 1 0 0 0 2 0v-3h2a1 1 0 0 0 1-1V9a5.979 5.979 0 0 0-1.827-4.302ZM15.992 16h-8v-5h8Zm-8-7a4 4 0 0 1 8 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AndroidPhoneSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.29 17.71a1 1 0 0 0 1.42 0a1.15 1.15 0 0 0 .21-.33a.94.94 0 0 0 0-.76a1.15 1.15 0 0 0-.21-.33a1 1 0 0 0-1.42 0a1.15 1.15 0 0 0-.21.33a.94.94 0 0 0 0 .76a1.15 1.15 0 0 0 .21.33M8.66 4H16a1 1 0 0 1 1 1v7.34a1 1 0 0 0 2 0V5a3 3 0 0 0-3-3H8.66a1 1 0 0 0 0 2m13.05 16.29l-18-18a1 1 0 0 0-1.42 1.42L5 6.41V19a3 3 0 0 0 3 3h8a3 3 0 0 0 2.76-1.83l1.53 1.54a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42M17 19a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1V8.41l10 10Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleDoubleDown(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.29 11.46a1 1 0 0 0 1.42 0l3-3A1 1 0 1 0 14.29 7L12 9.34L9.71 7a1 1 0 1 0-1.42 1.46Zm3 1.08L12 14.84l-2.29-2.3A1 1 0 0 0 8.29 14l3 3a1 1 0 0 0 1.42 0l3-3a1 1 0 0 0-1.42-1.42Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleDoubleLeft(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.46 8.29a1 1 0 0 0-1.42 0l-3 3a1 1 0 0 0 0 1.42l3 3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L9.16 12l2.3-2.29a1 1 0 0 0 0-1.42m3.2 3.71L17 9.71a1 1 0 0 0-1.42-1.42l-3 3a1 1 0 0 0 0 1.42l3 3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleDoubleRight(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.46 8.29A1 1 0 1 0 7 9.71L9.34 12L7 14.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l3-3a1 1 0 0 0 0-1.42Zm8.5 3l-3-3a1 1 0 0 0-1.42 1.42l2.3 2.29l-2.3 2.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l3-3a1 1 0 0 0 .04-1.42Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleDoubleUp(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.71 12.54a1 1 0 0 0-1.42 0l-3 3A1 1 0 0 0 9.71 17L12 14.66L14.29 17a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Zm-3-1.08L12 9.16l2.29 2.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-3-3a1 1 0 0 0-1.42 0l-3 3a1 1 0 0 0 1.42 1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleDown(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 9.17a1 1 0 0 0-1.41 0L12 12.71L8.46 9.17a1 1 0 0 0-1.41 0a1 1 0 0 0 0 1.42l4.24 4.24a1 1 0 0 0 1.42 0L17 10.59a1 1 0 0 0 0-1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleLeft(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m11.29 12l3.54-3.54a1 1 0 0 0 0-1.41a1 1 0 0 0-1.42 0l-4.24 4.24a1 1 0 0 0 0 1.42L13.41 17a1 1 0 0 0 .71.29a1 1 0 0 0 .71-.29a1 1 0 0 0 0-1.41Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleLeftB(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8.5 12.8l5.7 5.6c.4.4 1 .4 1.4 0c.4-.4.4-1 0-1.4l-4.9-5l4.9-5c.4-.4.4-1 0-1.4c-.2-.2-.4-.3-.7-.3c-.3 0-.5.1-.7.3l-5.7 5.6c-.4.5-.4 1.1 0 1.6c0-.1 0-.1 0 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleRight(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14.83 11.29l-4.24-4.24a1 1 0 0 0-1.42 0a1 1 0 0 0 0 1.41L12.71 12l-3.54 3.54a1 1 0 0 0 0 1.41a1 1 0 0 0 .71.29a1 1 0 0 0 .71-.29l4.24-4.24a1 1 0 0 0 0-1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleRightB(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.54 11.29L9.88 5.64a1 1 0 0 0-1.42 0a1 1 0 0 0 0 1.41l4.95 5L8.46 17a1 1 0 0 0 0 1.41a1 1 0 0 0 .71.3a1 1 0 0 0 .71-.3l5.66-5.65a1 1 0 0 0 0-1.47"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleUp(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m17 13.41l-4.29-4.24a1 1 0 0 0-1.42 0l-4.24 4.24a1 1 0 0 0 0 1.42a1 1 0 0 0 1.41 0L12 11.29l3.54 3.54a1 1 0 0 0 .7.29a1 1 0 0 0 .71-.29a1 1 0 0 0 .05-1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Angry(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 11a1 1 0 0 0 .89-.55a1 1 0 0 0-.44-1.34l-2-1a1 1 0 1 0-.9 1.78l2 1A.93.93 0 0 0 10 11m2-9a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8m-3.64-4.67a1 1 0 0 0-.13 1.4a1 1 0 0 0 1.41.13a3.76 3.76 0 0 1 4.72 0a1 1 0 0 0 .64.23a1 1 0 0 0 .64-1.76a5.81 5.81 0 0 0-7.28 0m7.19-7.22l-2 1a1 1 0 0 0-.44 1.34A1 1 0 0 0 14 11a.93.93 0 0 0 .45-.11l2-1a1 1 0 0 0-.9-1.78"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ankh(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 12h-3.09A7.23 7.23 0 0 0 17 7A5 5 0 0 0 7 7a7.23 7.23 0 0 0 2.09 5H6a1 1 0 0 0 0 2h5v7a1 1 0 0 0 2 0v-7h5a1 1 0 0 0 0-2m-6-.16c-.93-.62-3-2.26-3-4.84a3 3 0 0 1 6 0c0 2.58-2.07 4.23-3 4.84"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Annoyed(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 9a1 1 0 1 0 1 1a1 1 0 0 0-1-1m6 4a5 5 0 0 0-4.37 2.57a1 1 0 0 0 .37 1.36a1 1 0 0 0 .49.13a1 1 0 0 0 .87-.51A3 3 0 0 1 15 15a1 1 0 0 0 0-2m0-4a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-3-7a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AnnoyedAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14.66 13.56l-4.19 1.5A1 1 0 0 0 10.8 17a1 1 0 0 0 .34-.06l4.2-1.5a1 1 0 1 0-.68-1.88m-4-5a1 1 0 0 0-1.41 0a1 1 0 0 1-1.42 0a1 1 0 0 0-1.41 0a1 1 0 0 0 0 1.42a3 3 0 0 0 4.24 0a1 1 0 0 0-.04-1.44Zm7 0a1 1 0 0 0-1.41 0a1 1 0 0 1-1.42 0a1 1 0 0 0-1.41 0a1 1 0 0 0 0 1.42a3 3 0 0 0 4.24 0a1 1 0 0 0-.04-1.44ZM12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Apple(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.94 5.19A4.38 4.38 0 0 0 16 2a4.44 4.44 0 0 0-3 1.52a4.17 4.17 0 0 0-1 3.09a3.69 3.69 0 0 0 2.94-1.42m2.52 7.44a4.51 4.51 0 0 1 2.16-3.81a4.66 4.66 0 0 0-3.66-2c-1.56-.16-3 .91-3.83.91s-2-.89-3.3-.87a4.92 4.92 0 0 0-4.14 2.53C2.93 12.45 4.24 17 6 19.47c.8 1.21 1.8 2.58 3.12 2.53s1.75-.82 3.28-.82s2 .82 3.3.79s2.22-1.24 3.06-2.45a11 11 0 0 0 1.38-2.85a4.41 4.41 0 0 1-2.68-4.04"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppleAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.504 15.744a3.386 3.386 0 0 1-2.054-3.127a3.497 3.497 0 0 1 1.687-2.947a1 1 0 0 0 .3-1.415a5.574 5.574 0 0 0-4.016-2.385a6.422 6.422 0 0 0 .965-3.977a.984.984 0 0 0-1.033-.892a6.69 6.69 0 0 0-4.444 2.261a6.842 6.842 0 0 0-1.48 2.9a5.264 5.264 0 0 0-1.62-.278a5.925 5.925 0 0 0-4.99 3.008C1.93 12.169 3.086 17.08 5.13 20.037c.82 1.184 2.051 2.962 3.908 2.962c.037 0 .074 0 .11-.002a4.74 4.74 0 0 0 1.793-.463a3.274 3.274 0 0 1 1.446-.354a3.045 3.045 0 0 1 1.357.338a4.752 4.752 0 0 0 1.962.456c1.913-.036 3.043-1.687 3.868-2.893a12.042 12.042 0 0 0 1.513-3.108a1.003 1.003 0 0 0-.584-1.23m-7.08-11.176a4.941 4.941 0 0 1 1.919-1.322a4.855 4.855 0 0 1-1.037 2.15a4.39 4.39 0 0 1-1.264 1.041l-.003.001l-.085.034c-.186.073-.417.164-.595.223a4.825 4.825 0 0 1 1.064-2.127m4.5 14.385c-.803 1.173-1.435 2.006-2.254 2.021a2.367 2.367 0 0 1-1.13-.292a4.97 4.97 0 0 0-2.152-.502a5.214 5.214 0 0 0-2.229.513a2.885 2.885 0 0 1-1.086.306h-.03c-.813 0-1.687-1.262-2.268-2.101c-1.452-2.1-2.792-6.283-1.223-9.008a3.917 3.917 0 0 1 3.29-2.006h.034a4.758 4.758 0 0 1 1.642.444c.13.051.26.103.389.15a.858.858 0 0 0 .1.038a3.315 3.315 0 0 0 1.118.237a4.436 4.436 0 0 0 1.564-.42a4.956 4.956 0 0 1 2.222-.498a3.787 3.787 0 0 1 2.25.854a5.31 5.31 0 0 0-1.711 3.945v.001a5.3 5.3 0 0 0 2.45 4.503a10.198 10.198 0 0 1-.975 1.815"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Apps(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 13H3a1 1 0 0 0-1 1v7a1 1 0 0 0 1 1h7a1 1 0 0 0 1-1v-7a1 1 0 0 0-1-1m-1 7H4v-5h5ZM21 2h-7a1 1 0 0 0-1 1v7a1 1 0 0 0 1 1h7a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1m-1 7h-5V4h5Zm1 4h-7a1 1 0 0 0-1 1v7a1 1 0 0 0 1 1h7a1 1 0 0 0 1-1v-7a1 1 0 0 0-1-1m-1 7h-5v-5h5ZM10 2H3a1 1 0 0 0-1 1v7a1 1 0 0 0 1 1h7a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1M9 9H4V4h5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Archive(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 14h4a1 1 0 0 0 0-2h-4a1 1 0 0 0 0 2m9-11H5a3 3 0 0 0-3 3v3a1 1 0 0 0 1 1h1v8a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-8h1a1 1 0 0 0 1-1V6a3 3 0 0 0-3-3m-1 15a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1v-8h12Zm2-10H4V6a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArchiveAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 13h4a1 1 0 0 0 0-2h-4a1 1 0 0 0 0 2m9-10H5a3 3 0 0 0-1 5.82V18a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V8.82A3 3 0 0 0 19 3m-1 15a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V9h12Zm1-11H5a1 1 0 0 1 0-2h14a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Archway(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 20h-1V8h1a1 1 0 0 0 0-2h-1V3a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v3H2a1 1 0 0 0 0 2h1v12H2a1 1 0 0 0 0 2h20a1 1 0 0 0 0-2m-7 0H9v-3.53a6.21 6.21 0 0 1 3-5.33a6.21 6.21 0 0 1 3 5.33Zm4 0h-2v-3.53a8.17 8.17 0 0 0-4.55-7.36a1 1 0 0 0-.9 0A8.17 8.17 0 0 0 7 16.47V20H5V8h14Zm0-14H5V4h14Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Arrow(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.71 6.29a1 1 0 0 0-1.42 0l-5 5a1 1 0 0 0 0 1.42l5 5a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L5.41 12l4.3-4.29a1 1 0 0 0 0-1.42m11 5l-5-5a1 1 0 0 0-1.42 1.42l4.3 4.29l-4.3 4.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l5-5a1 1 0 0 0 0-1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowBreak(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 11h10a1 1 0 0 0 0-2h-4V5.41l.79.8a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-2.5-2.5a1 1 0 0 0-.33-.21a1 1 0 0 0-.76 0a1 1 0 0 0-.33.21l-2.5 2.5a1 1 0 0 0 1.42 1.42l.79-.8V9H7a1 1 0 0 0 0 2m10 2H7a1 1 0 0 0 0 2h4v3.59l-.79-.8a1 1 0 0 0-1.42 1.42l2.5 2.5a1 1 0 0 0 .33.21a.94.94 0 0 0 .76 0a1 1 0 0 0 .33-.21l2.5-2.5a1 1 0 0 0-1.42-1.42l-.79.8V15h4a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleDown(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.29 15.71a1 1 0 0 0 .33.21a1 1 0 0 0 .76 0a1 1 0 0 0 .33-.21l3-3a1 1 0 0 0-1.42-1.42L13 12.59V9a1 1 0 0 0-2 0v3.59l-1.29-1.3a1 1 0 0 0-1.42 0a1 1 0 0 0 0 1.42ZM12 22A10 10 0 1 0 2 12a10 10 0 0 0 10 10m0-18a8 8 0 1 1-8 8a8 8 0 0 1 8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleLeft(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.29 11.29a1 1 0 0 0-.21.33a1 1 0 0 0 0 .76a1 1 0 0 0 .21.33l3 3a1 1 0 0 0 1.42-1.42L11.41 13H15a1 1 0 0 0 0-2h-3.59l1.3-1.29a1 1 0 0 0 0-1.42a1 1 0 0 0-1.42 0ZM2 12A10 10 0 1 0 12 2A10 10 0 0 0 2 12m18 0a8 8 0 1 1-8-8a8 8 0 0 1 8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleRight(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.71 12.71a1 1 0 0 0 .21-.33a1 1 0 0 0 0-.76a1 1 0 0 0-.21-.33l-3-3a1 1 0 0 0-1.42 1.42l1.3 1.29H9a1 1 0 0 0 0 2h3.59l-1.3 1.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0ZM22 12a10 10 0 1 0-10 10a10 10 0 0 0 10-10M4 12a8 8 0 1 1 8 8a8 8 0 0 1-8-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleUp(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.71 8.29a1 1 0 0 0-.33-.21a1 1 0 0 0-.76 0a1 1 0 0 0-.33.21l-3 3a1 1 0 0 0 1.42 1.42l1.29-1.3V15a1 1 0 0 0 2 0v-3.59l1.29 1.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCompressH(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 5a1 1 0 0 0-1 1v12a1 1 0 0 0 2 0V6a1 1 0 0 0-1-1m-1.29 6.29l-2.5-2.5a1 1 0 1 0-1.42 1.42l.8.79H3a1 1 0 0 0 0 2h4.59l-.8.79a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l2.5-2.5a1 1 0 0 0 .21-.33a.94.94 0 0 0 0-.76a1 1 0 0 0-.21-.33M21 11h-4.59l.8-.79a1 1 0 0 0-1.42-1.42l-2.5 2.5a1 1 0 0 0 0 1.42l2.5 2.5a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-.8-.79H21a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDown(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.71 11.29a1 1 0 0 0-1.42 0L13 14.59V7a1 1 0 0 0-2 0v7.59l-3.29-3.3a1 1 0 0 0-1.42 1.42l5 5a1 1 0 0 0 .33.21a.94.94 0 0 0 .76 0a1 1 0 0 0 .33-.21l5-5a1 1 0 0 0 0-1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownLeft(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 16H9.41l8.3-8.29a1 1 0 1 0-1.42-1.42L8 14.59V7a1 1 0 0 0-2 0v10a1 1 0 0 0 .08.38a1 1 0 0 0 .54.54A1 1 0 0 0 7 18h10a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownRight(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 6a1 1 0 0 0-1 1v7.59l-8.29-8.3a1 1 0 0 0-1.42 1.42l8.3 8.29H7a1 1 0 0 0 0 2h10a1 1 0 0 0 .38-.08a1 1 0 0 0 .54-.54A1 1 0 0 0 18 17V7a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowFromRight(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 4a1 1 0 0 0-1 1v14a1 1 0 0 0 2 0V5a1 1 0 0 0-1-1m18.92 7.62a1 1 0 0 0-.21-.33l-4-4a1 1 0 1 0-1.42 1.42l2.3 2.29H7a1 1 0 0 0 0 2h11.59l-2.3 2.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l4-4a1 1 0 0 0 .21-.33a1 1 0 0 0 0-.76"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowFromTop(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.29 16.29L13 18.59V7a1 1 0 0 0-2 0v11.59l-2.29-2.3a1 1 0 1 0-1.42 1.42l4 4a1 1 0 0 0 .33.21a.94.94 0 0 0 .76 0a1 1 0 0 0 .33-.21l4-4a1 1 0 0 0-1.42-1.42M19 2H5a1 1 0 0 0 0 2h14a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowGrowth(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.92 6.62a1 1 0 0 0-.54-.54A1 1 0 0 0 21 6h-5a1 1 0 0 0 0 2h2.59L13 13.59l-3.29-3.3a1 1 0 0 0-1.42 0l-6 6a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0L9 12.41l3.29 3.3a1 1 0 0 0 1.42 0L20 9.41V12a1 1 0 0 0 2 0V7a1 1 0 0 0-.08-.38"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeft(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 11H9.41l3.3-3.29a1 1 0 1 0-1.42-1.42l-5 5a1 1 0 0 0-.21.33a1 1 0 0 0 0 .76a1 1 0 0 0 .21.33l5 5a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L9.41 13H17a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRandom(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.7 10a1 1 0 0 0 1.41 0a1 1 0 0 0 0-1.41l-6.27-6.3a1 1 0 0 0-1.42 1.42ZM21 14a1 1 0 0 0-1 1v3.59L15.44 14A1 1 0 0 0 14 15.44L18.59 20H15a1 1 0 0 0 0 2h6a1 1 0 0 0 .38-.08a1 1 0 0 0 .54-.54A1 1 0 0 0 22 21v-6a1 1 0 0 0-1-1m.92-11.38a1 1 0 0 0-.54-.54A1 1 0 0 0 21 2h-6a1 1 0 0 0 0 2h3.59L2.29 20.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0L20 5.41V9a1 1 0 0 0 2 0V3a1 1 0 0 0-.08-.38"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowResizeDiagonal(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.92 2.62a1 1 0 0 0-.54-.54A1 1 0 0 0 21 2h-6a1 1 0 0 0 0 2h3.59L4 18.59V15a1 1 0 0 0-2 0v6a1 1 0 0 0 .08.38a1 1 0 0 0 .54.54A1 1 0 0 0 3 22h6a1 1 0 0 0 0-2H5.41L20 5.41V9a1 1 0 0 0 2 0V3a1 1 0 0 0-.08-.38"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRight(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.92 11.62a1 1 0 0 0-.21-.33l-5-5a1 1 0 0 0-1.42 1.42l3.3 3.29H7a1 1 0 0 0 0 2h7.59l-3.3 3.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l5-5a1 1 0 0 0 .21-.33a1 1 0 0 0 0-.76"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowToBottom(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 20H5a1 1 0 0 0 0 2h14a1 1 0 0 0 0-2m-7.71-2.29a1 1 0 0 0 .33.21a.94.94 0 0 0 .76 0a1 1 0 0 0 .33-.21l4-4a1 1 0 0 0-1.42-1.42L13 14.59V3a1 1 0 0 0-2 0v11.59l-2.29-2.3a1 1 0 1 0-1.42 1.42Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowToRight(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m17.71 11.29l-4-4a1 1 0 1 0-1.42 1.42l2.3 2.29H3a1 1 0 0 0 0 2h11.59l-2.3 2.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l4-4a1 1 0 0 0 .21-.33a1 1 0 0 0 0-.76a1 1 0 0 0-.21-.33M21 4a1 1 0 0 0-1 1v14a1 1 0 0 0 2 0V5a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUp(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m17.71 11.29l-5-5a1 1 0 0 0-.33-.21a1 1 0 0 0-.76 0a1 1 0 0 0-.33.21l-5 5a1 1 0 0 0 1.42 1.42L11 9.41V17a1 1 0 0 0 2 0V9.41l3.29 3.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpLeft(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.41 8H17a1 1 0 0 0 0-2H7a1 1 0 0 0-.38.08a1 1 0 0 0-.54.54A1 1 0 0 0 6 7v10a1 1 0 0 0 2 0V9.41l8.29 8.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpRight(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.92 6.62a1 1 0 0 0-.54-.54A1 1 0 0 0 17 6H7a1 1 0 0 0 0 2h7.59l-8.3 8.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0L16 9.41V17a1 1 0 0 0 2 0V7a1 1 0 0 0-.08-.38"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsH(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.92 11.62a1 1 0 0 0-.21-.33l-4-4a1 1 0 1 0-1.42 1.42l2.3 2.29H5.41l2.3-2.29a1 1 0 1 0-1.42-1.42l-4 4a1 1 0 0 0-.21.33a1 1 0 0 0 0 .76a1 1 0 0 0 .21.33l4 4a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L5.41 13h13.18l-2.3 2.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l4-4a1 1 0 0 0 .21-.33a1 1 0 0 0 0-.76"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsHalt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.92 11.62a1 1 0 0 0-.21-.33l-2.5-2.5a1 1 0 0 0-1.42 1.42l.8.79H14a1 1 0 0 0 0 2h4.59l-.8.79a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l2.5-2.5a1 1 0 0 0 .21-.33a1 1 0 0 0 0-.76M10 11H5.41l.8-.79a1 1 0 0 0-1.42-1.42l-2.5 2.5a1 1 0 0 0-.21.33a1 1 0 0 0 0 .76a1 1 0 0 0 .21.33l2.5 2.5a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-.8-.79H10a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsLeftDown(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.71 16.29a1 1 0 0 0-1.42 0L18 18.59V9a3 3 0 0 0-3-3H5.41l2.3-2.29a1 1 0 0 0-1.42-1.42l-4 4a1 1 0 0 0-.21.33a1 1 0 0 0 0 .76a1 1 0 0 0 .21.33l4 4a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L5.41 8H15a1 1 0 0 1 1 1v9.59l-2.29-2.3a1 1 0 0 0-1.42 1.42l4 4a1 1 0 0 0 .33.21a.94.94 0 0 0 .76 0a1 1 0 0 0 .33-.21l4-4a1 1 0 0 0 0-1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsMaximize(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 14a1 1 0 0 0-1 1v3.59L5.41 4H9a1 1 0 0 0 0-2H3a1 1 0 0 0-.38.08a1 1 0 0 0-.54.54A1 1 0 0 0 2 3v6a1 1 0 0 0 2 0V5.41L18.59 20H15a1 1 0 0 0 0 2h6a1 1 0 0 0 .38-.08a1 1 0 0 0 .54-.54A1 1 0 0 0 22 21v-6a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsMerge(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.71 11.29l-2.5-2.5a1 1 0 1 0-1.42 1.42l.8.79H4V7a1 1 0 0 0-2 0v10a1 1 0 0 0 2 0v-4h3.59l-.8.79a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l2.5-2.5a1 1 0 0 0 .21-.33a1 1 0 0 0 0-.76a1 1 0 0 0-.21-.33M21 6a1 1 0 0 0-1 1v4h-3.59l.8-.79a1 1 0 0 0-1.42-1.42l-2.5 2.5a1 1 0 0 0-.21.33a1 1 0 0 0 0 .76a1 1 0 0 0 .21.33l2.5 2.5a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-.8-.79H20v4a1 1 0 0 0 2 0V7a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsResize(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.92 11.62a1 1 0 0 0-.21-.33l-2.5-2.5a1 1 0 0 0-1.42 1.42l.8.79H13V8a1 1 0 0 0-2 0v3H5.41l.8-.79a1 1 0 0 0-1.42-1.42l-2.5 2.5a1 1 0 0 0-.21.33a1 1 0 0 0 0 .76a1 1 0 0 0 .21.33l2.5 2.5a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-.8-.79H11v3a1 1 0 0 0 2 0v-3h5.59l-.8.79a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l2.5-2.5a1 1 0 0 0 .21-.33a1 1 0 0 0 0-.76"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsResizeH(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 6a1 1 0 0 0-1 1v4H5.41l.8-.79a1 1 0 0 0-1.42-1.42l-2.5 2.5a1 1 0 0 0-.21.33a1 1 0 0 0 0 .76a1 1 0 0 0 .21.33l2.5 2.5a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-.8-.79H9v4a1 1 0 0 0 2 0V7a1 1 0 0 0-1-1m11.92 5.62a1 1 0 0 0-.21-.33l-2.5-2.5a1 1 0 0 0-1.42 1.42l.8.79H15V7a1 1 0 0 0-2 0v10a1 1 0 0 0 2 0v-4h3.59l-.8.79a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l2.5-2.5a1 1 0 0 0 .21-.33a1 1 0 0 0 0-.76"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsResizeV(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 11h-3V5.41l.79.8a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-2.5-2.5a1 1 0 0 0-.33-.21a1 1 0 0 0-.76 0a1 1 0 0 0-.33.21l-2.5 2.5a1 1 0 0 0 1.42 1.42l.79-.8V11H8a1 1 0 0 0 0 2h3v5.59l-.79-.8a1 1 0 0 0-1.42 1.42l2.5 2.5a1 1 0 0 0 .33.21a.94.94 0 0 0 .76 0a1 1 0 0 0 .33-.21l2.5-2.5a1 1 0 0 0-1.42-1.42l-.79.8V13h3a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsRightDown(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.92 6.62a1 1 0 0 0-.21-.33l-3-3a1 1 0 0 0-1.42 1.42L17.59 6H9a3 3 0 0 0-3 3v8.59l-1.29-1.3a1 1 0 0 0-1.42 1.42l3 3a1 1 0 0 0 .33.21a.94.94 0 0 0 .76 0a1 1 0 0 0 .33-.21l3-3a1 1 0 0 0-1.42-1.42L8 17.59V9a1 1 0 0 1 1-1h8.59l-1.3 1.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l3-3a1 1 0 0 0 .21-.33a1 1 0 0 0 0-.76"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsShrinkH(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m17.71 11.29l-2.5-2.5a1 1 0 0 0-1.42 1.42l.8.79H9.41l.8-.79a1 1 0 0 0-1.42-1.42l-2.5 2.5a1 1 0 0 0-.21.33a1 1 0 0 0 0 .76a1 1 0 0 0 .21.33l2.5 2.5a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-.8-.79h5.18l-.8.79a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l2.5-2.5a1 1 0 0 0 .21-.33a1 1 0 0 0 0-.76a1 1 0 0 0-.21-.33M3 6a1 1 0 0 0-1 1v10a1 1 0 0 0 2 0V7a1 1 0 0 0-1-1m18 0a1 1 0 0 0-1 1v10a1 1 0 0 0 2 0V7a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsShrinkV(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.79 10.21a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-2.5-2.5a1 1 0 0 0-.33-.21a1 1 0 0 0-.76 0a1 1 0 0 0-.33.21l-2.5 2.5a1 1 0 0 0 1.42 1.42l.79-.8v5.18l-.79-.8a1 1 0 0 0-1.42 1.42l2.5 2.5a1 1 0 0 0 .33.21a.94.94 0 0 0 .76 0a1 1 0 0 0 .33-.21l2.5-2.5a1 1 0 0 0-1.42-1.42l-.79.8V9.41ZM7 4h10a1 1 0 0 0 0-2H7a1 1 0 0 0 0 2m10 16H7a1 1 0 0 0 0 2h10a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsUpRight(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.92 16.62a1 1 0 0 0-.21-.33l-3-3a1 1 0 0 0-1.42 1.42l1.3 1.29H9a1 1 0 0 1-1-1V6.41l1.29 1.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-3-3a1 1 0 0 0-.33-.21a1 1 0 0 0-.76 0a1 1 0 0 0-.33.21l-3 3a1 1 0 0 0 1.42 1.42L6 6.41V15a3 3 0 0 0 3 3h8.59l-1.3 1.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l3-3a1 1 0 0 0 .21-.33a1 1 0 0 0 0-.76"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsV(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.29 16.29L13 18.59V5.41l2.29 2.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-4-4a1 1 0 0 0-.33-.21a1 1 0 0 0-.76 0a1 1 0 0 0-.33.21l-4 4a1 1 0 1 0 1.42 1.42L11 5.41v13.18l-2.29-2.3a1 1 0 1 0-1.42 1.42l4 4a1 1 0 0 0 .33.21a.94.94 0 0 0 .76 0a1 1 0 0 0 .33-.21l4-4a1 1 0 0 0-1.42-1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsValt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.21 6.21l.79-.8V10a1 1 0 0 0 2 0V5.41l.79.8a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-2.5-2.5a1 1 0 0 0-.33-.21a1 1 0 0 0-.76 0a1 1 0 0 0-.33.21l-2.5 2.5a1 1 0 0 0 1.42 1.42m3.58 11.58l-.79.8V14a1 1 0 0 0-2 0v4.59l-.79-.8a1 1 0 0 0-1.42 1.42l2.5 2.5a1 1 0 0 0 .33.21a.94.94 0 0 0 .76 0a1 1 0 0 0 .33-.21l2.5-2.5a1 1 0 0 0-1.42-1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AssistiveListeningSystems(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 3a7 7 0 0 0-7 7a1 1 0 0 0 2 0a5 5 0 0 1 10 0a5.07 5.07 0 0 1-.71 2.57L11.73 20A2 2 0 0 1 10 21a2 2 0 0 1-2-2a1 1 0 0 0-2 0a4 4 0 0 0 4 4a4 4 0 0 0 3.5-2.07l3.56-7.43A6.93 6.93 0 0 0 18 10a7 7 0 0 0-7-7M4 15a1 1 0 1 0 1 1a1 1 0 0 0-1-1M17.59 1.2a1 1 0 1 0-1.2 1.6A9 9 0 0 1 20 10a1 1 0 0 0 2 0a11.06 11.06 0 0 0-4.41-8.8M11 9a1 1 0 0 1 1 1a1 1 0 0 0 2 0a3 3 0 0 0-6 0a3 3 0 0 0 .51 1.68a3.5 3.5 0 0 0 .47.54l.2.22a1 1 0 0 1 0 1.11a1 1 0 0 0 .25 1.39a1 1 0 0 0 .57.18a1 1 0 0 0 .82-.43a3 3 0 0 0 0-3.39a3.39 3.39 0 0 0-.35-.42l-.14-.14a1.37 1.37 0 0 1-.16-.18A1 1 0 0 1 10 10a1 1 0 0 1 1-1m-4 6a1 1 0 1 0-1-1a1 1 0 0 0 1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Asterisk(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.562 14.634L14 12l4.562-2.634a1 1 0 0 0-1-1.732L13 10.268V5a1 1 0 0 0-2 0v5.268L6.438 7.634a1 1 0 0 0-1 1.732L10 12l-4.562 2.634a1 1 0 0 0 1 1.732L11 13.732V19a1 1 0 0 0 2 0v-5.268l4.562 2.634a1 1 0 0 0 1-1.732"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func At(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a10 10 0 1 0 5 18.66a1 1 0 1 0-1-1.73A8 8 0 1 1 20 12v.75a1.75 1.75 0 0 1-3.5 0V8.5a1 1 0 0 0-1-1a1 1 0 0 0-1 .79A4.45 4.45 0 0 0 12 7.5a4.5 4.5 0 1 0 3.3 7.5a3.74 3.74 0 0 0 6.7-2.25V12A10 10 0 0 0 12 2m0 12.5a2.5 2.5 0 1 1 2.5-2.5a2.5 2.5 0 0 1-2.5 2.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Atom(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 11a1 1 0 1 0 1 1a1 1 0 0 0-1-1m7.62 1l.11-.14C21.08 10 21.4 8.29 20.66 7S18.26 5.14 16 5.37h-.18C14.91 3.3 13.56 2 12 2S9.09 3.3 8.19 5.4L8 5.37C5.74 5.14 4.08 5.71 3.34 7s-.42 3 .93 4.86l.11.14l-.11.14C2.92 14 2.6 15.71 3.34 17C4 18.1 5.27 18.68 7 18.68c.31 0 .63 0 1-.05h.18C9.09 20.7 10.44 22 12 22s2.91-1.3 3.81-3.4h.18c.34 0 .66.05 1 .05c1.77 0 3.07-.58 3.7-1.68c.74-1.29.42-3-.93-4.86ZM5.07 8c.25-.44 1-.68 2-.68h.49a14.78 14.78 0 0 0-.35 1.87a15 15 0 0 0-1.45 1.25C5 9.44 4.78 8.5 5.07 8m0 8c-.29-.5 0-1.44.67-2.47a15 15 0 0 0 1.45 1.25a14.94 14.94 0 0 0 .35 1.88c-1.24.08-2.18-.16-2.47-.66M12 4c.56 0 1.23.66 1.8 1.83a17.6 17.6 0 0 0-1.8.63a17.6 17.6 0 0 0-1.8-.63C10.77 4.66 11.44 4 12 4m0 16c-.56 0-1.23-.66-1.8-1.83a17.6 17.6 0 0 0 1.8-.63a17.6 17.6 0 0 0 1.8.63C13.23 19.34 12.56 20 12 20m2.93-6.31c-.46.32-.93.62-1.43.91s-1 .55-1.5.78q-.75-.35-1.5-.78c-.5-.29-1-.59-1.43-.91C9 13.15 9 12.59 9 12s0-1.15.07-1.69c.46-.32.93-.62 1.43-.91s1-.55 1.5-.78q.75.35 1.5.78c.5.29 1 .59 1.43.91c0 .54.07 1.1.07 1.69s0 1.15-.07 1.69m4 2.31c-.29.5-1.23.75-2.47.66a14.94 14.94 0 0 0 .35-1.88a15 15 0 0 0 1.45-1.25c.74 1.03.96 1.97.67 2.47m-.67-5.53a15 15 0 0 0-1.45-1.25a14.78 14.78 0 0 0-.35-1.87h.49c1 0 1.73.24 2 .68s.05 1.41-.69 2.44"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AutoFlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.87 8.6A1 1 0 0 0 17 8h-4.58l1.27-4.74a1 1 0 0 0-.17-.87a1 1 0 0 0-.8-.39h-7a1 1 0 0 0-1 .74l-2.68 10a1 1 0 0 0 1 1.26h3.85l-1.81 6.74a1 1 0 0 0 1.71.93l10.9-12a1 1 0 0 0 .18-1.07m-9.79 8.68l1.08-4a1 1 0 0 0-.16-.89a1 1 0 0 0-.81-.39H4.35l2.14-8h4.93l-1.27 4.74a1 1 0 0 0 1 1.26h3.57ZM19 13h-1a3 3 0 0 0-3 3v5a1 1 0 0 0 2 0v-2h3v2a1 1 0 0 0 2 0v-5a3 3 0 0 0-3-3m1 4h-3v-1a1 1 0 0 1 1-1h1a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Award(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.87 17.25l-2.71-4.68A6.9 6.9 0 0 0 19 9.25a7 7 0 0 0-14 0a6.9 6.9 0 0 0 .84 3.32l-2.71 4.68a1 1 0 0 0 .87 1.5h2.87l1.46 2.46a1 1 0 0 0 .18.22a1 1 0 0 0 .69.28h.14a1 1 0 0 0 .73-.49L12 17.9l1.93 3.35a1 1 0 0 0 .73.48h.14a1 1 0 0 0 .7-.28a.87.87 0 0 0 .17-.21l1.46-2.46H20a1 1 0 0 0 .87-.5a1 1 0 0 0 0-1.03M9.19 18.78l-.89-1.49a1 1 0 0 0-.85-.49H5.72l1.43-2.48a7 7 0 0 0 3.57 1.84ZM12 14.25a5 5 0 1 1 5-5a5 5 0 0 1-5 5m4.55 2.55a1 1 0 0 0-.85.49l-.89 1.49l-1.52-2.65a7.06 7.06 0 0 0 3.56-1.84l1.43 2.48Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AwardAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1a7 7 0 0 0-5 11.89V22a1 1 0 0 0 1.45.89L12 21.12l3.55 1.77A1 1 0 0 0 16 23a1 1 0 0 0 .53-.15A1 1 0 0 0 17 22v-9.11A7 7 0 0 0 12 1m3 19.38l-2.55-1.27a1 1 0 0 0-.9 0L9 20.38v-6.06a7 7 0 0 0 2 .6V16a1 1 0 0 0 2 0v-1.08a7 7 0 0 0 2-.6ZM12 13a5 5 0 1 1 5-5a5 5 0 0 1-5 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BabyCarriage(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 16a3 3 0 1 0 3 3a3 3 0 0 0-3-3m0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1M9 16a3 3 0 1 0 3 3a3 3 0 0 0-3-3m0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1M22 8.5A6.51 6.51 0 0 0 15.5 2H15a1 1 0 0 0-1 1v5H7.52L6.27 4.65A1 1 0 0 0 5.33 4H3a1 1 0 0 0 0 2h1.64l1.25 3.37l.51 1.37v.09A6.44 6.44 0 0 0 12.5 15h3A6.49 6.49 0 0 0 22 8.5m-3.32 3.18A4.47 4.47 0 0 1 15.5 13h-3a4.47 4.47 0 0 1-4.16-2.8a.14.14 0 0 1 0-.06L8.26 10h11.48a4.32 4.32 0 0 1-1.06 1.68M16 8V4a4.49 4.49 0 0 1 4 4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Backpack(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 10h-2a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2m5 0a4 4 0 0 0-3-3.86V5a3 3 0 0 0-6 0v1.14A4 4 0 0 0 6 10a4 4 0 0 0-4 4v3a3 3 0 0 0 3 3h1.18A3 3 0 0 0 9 22h6a3 3 0 0 0 2.82-2H19a3 3 0 0 0 3-3v-3a4 4 0 0 0-4-4M6 18H5a1 1 0 0 1-1-1v-3a2 2 0 0 1 2-2Zm5-13a1 1 0 0 1 2 0v1h-2Zm5 14a1 1 0 0 1-1 1H9a1 1 0 0 1-1-1v-1a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2Zm0-4.44a3.91 3.91 0 0 0-2-.56h-4a3.91 3.91 0 0 0-2 .56V10a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2ZM20 17a1 1 0 0 1-1 1h-1v-6a2 2 0 0 1 2 2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Backspace(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 5H9.83a3 3 0 0 0-2.12.88l-5.42 5.41a1 1 0 0 0 0 1.42l5.42 5.41a3 3 0 0 0 2.12.88H19a3 3 0 0 0 3-3V8a3 3 0 0 0-3-3m1 11a1 1 0 0 1-1 1H9.83a1.05 1.05 0 0 1-.71-.29L4.41 12l4.71-4.71A1.05 1.05 0 0 1 9.83 7H19a1 1 0 0 1 1 1Zm-3.29-6.71a1 1 0 0 0-1.42 0L14 10.59l-1.29-1.3a1 1 0 0 0-1.42 1.42l1.3 1.29l-1.3 1.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l1.29-1.3l1.29 1.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L15.41 12l1.3-1.29a1 1 0 0 0 0-1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Backward(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.67 5.85a2.63 2.63 0 0 0-2.67 0l-4 2.3a2.67 2.67 0 0 0-4-2.31L3.33 9.69a2.67 2.67 0 0 0 0 4.62L10 18.16a2.66 2.66 0 0 0 2.67 0A2.65 2.65 0 0 0 14 15.85l4 2.31a2.69 2.69 0 0 0 1.33.36a2.61 2.61 0 0 0 1.34-.37a2.63 2.63 0 0 0 1.33-2.3v-7.7a2.63 2.63 0 0 0-1.33-2.3m-8.67 10a.66.66 0 0 1-.33.58a.69.69 0 0 1-.67 0l-6.67-3.85a.67.67 0 0 1 0-1.16L11 7.57a.67.67 0 0 1 .67 0a.66.66 0 0 1 .33.58Zm8 0a.67.67 0 0 1-1 .57l-5-2.88v-3.08l5-2.88a.67.67 0 0 1 1 .57Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bag(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 6h-3V5a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v1H5a3 3 0 0 0-3 3v9a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3m-9-1h4v1h-4Zm10 13a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-5.61L8.68 14A1.19 1.19 0 0 0 9 14h6a1.19 1.19 0 0 0 .32-.05L20 12.39Zm0-7.72L14.84 12H9.16L4 10.28V9a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BagAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 6.5h-3v-1a3 3 0 0 0-3-3h-2a3 3 0 0 0-3 3v1H5a3 3 0 0 0-3 3v9a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-9a3 3 0 0 0-3-3m-9-1a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v1h-4Zm10 13a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V13a21.71 21.71 0 0 0 8 1.53A21.75 21.75 0 0 0 20 13Zm0-7.69a19.89 19.89 0 0 1-16 0V9.5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BagSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 6h-3V5a2 2 0 0 0-2-2h-4a1.86 1.86 0 0 0-.61.1a1 1 0 0 0-.64 1.27A1 1 0 0 0 10 5h4v1h-1.34a1 1 0 0 0 0 2H19a1 1 0 0 1 1 1v1.28l-2.57.86a1 1 0 0 0-.63 1.27a1 1 0 0 0 .95.68a1.19 1.19 0 0 0 .32-.05l1.93-.65v2.95a1 1 0 1 0 2 0V9a3 3 0 0 0-3-3M3.71 2.29a1 1 0 0 0-1.42 1.42L4.62 6A3 3 0 0 0 2 9v9a3 3 0 0 0 3 3h14a3.07 3.07 0 0 0 .53-.06l.76.77a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM4 9a1 1 0 0 1 1-1h1.59l4 4H9.16L4 10.28Zm1 10a1 1 0 0 1-1-1v-5.61L8.68 14A1.19 1.19 0 0 0 9 14h3.59l5 5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BalanceScale(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.964 13.823a.948.948 0 0 0-.028-.175l-2.305-6.137A2.996 2.996 0 0 0 22 5a1 1 0 0 0-2 0a1 1 0 0 1-1.882.473A2.893 2.893 0 0 0 15.54 4H13V3a1 1 0 0 0-2 0v1H8.46a2.893 2.893 0 0 0-2.578 1.473A1 1 0 0 1 4 5a1 1 0 0 0-2 0a2.996 2.996 0 0 0 1.369 2.511l-2.305 6.137a.948.948 0 0 0-.028.175A.949.949 0 0 0 1 14c0 .01.003.018.003.027c0 .013.003.025.004.039a3.994 3.994 0 0 0 7.986 0c.001-.014.004-.026.004-.039c0-.01.003-.018.003-.027a.949.949 0 0 0-.036-.177a.948.948 0 0 0-.028-.175L6.629 7.504A2.99 2.99 0 0 0 7.643 6.42A.917.917 0 0 1 8.46 6H11v14H8a1 1 0 0 0 0 2h8a1 1 0 0 0 0-2h-3V6h2.54a.917.917 0 0 1 .817.42a2.99 2.99 0 0 0 1.014 1.084l-2.307 6.144a.948.948 0 0 0-.028.175A.949.949 0 0 0 15 14c0 .01.003.018.003.027c0 .013.003.025.004.039a3.994 3.994 0 0 0 7.986 0c.001-.014.004-.026.004-.039c0-.01.003-.018.003-.027a.949.949 0 0 0-.036-.177M5 8.856L6.556 13H3.444ZM6.723 15A2.023 2.023 0 0 1 5 16a2 2 0 0 1-1.731-1ZM19 8.856L20.556 13h-3.112ZM19 16a2 2 0 0 1-1.731-1h3.454A2.023 2.023 0 0 1 19 16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ban(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 0 1-8-8a7.92 7.92 0 0 1 1.69-4.9L16.9 18.31A7.92 7.92 0 0 1 12 20m6.31-3.1L7.1 5.69A7.92 7.92 0 0 1 12 4a8 8 0 0 1 8 8a7.92 7.92 0 0 1-1.69 4.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BandAid(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.82 11.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.41 0a1 1 0 0 0 0-1.42a1 1 0 0 0-1.41 0m2.47 2.48a1 1 0 0 0 0 1.41a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.41a1 1 0 0 0-1.42 0m9-10A6 6 0 0 0 12 3.55a6 6 0 0 0-8.24.2A6 6 0 0 0 3.57 12A6 6 0 0 0 8 22a5.92 5.92 0 0 0 4-1.55a6 6 0 0 0 8.25-.2a6 6 0 0 0 .18-8.25a6 6 0 0 0-.18-8.25Zm-1.46 1.4a4 4 0 0 1 .17 5.39L13.44 5a4.07 4.07 0 0 1 5.39.17M5.17 18.83A4 4 0 0 1 5 13.44l5.6 5.6a4.08 4.08 0 0 1-5.43-.21m13.66 0a4.08 4.08 0 0 1-5.64 0l-8-8a4 4 0 0 1 0-5.64a4 4 0 0 1 5.64 0l8 8a4 4 0 0 1 0 5.64m-5.06-7.54a1 1 0 0 0 0 1.42a1 1 0 0 0 1.41 0a1 1 0 0 0 0-1.42a1 1 0 0 0-1.41 0m-2.48-2.47a1 1 0 0 0 0 1.41a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.41a1 1 0 0 0-1.42 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bars(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 8h18a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2m18 8H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2m0-5H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BaseballBall(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.922 6.44a1 1 0 0 0-1.703 1.05q.265.428.565.843a1 1 0 0 0 1.62-1.174q-.257-.353-.482-.719m-2.73 9.222a1 1 0 1 0-1.609 1.188c.17.23.327.466.476.709a1 1 0 1 0 1.705-1.047c-.178-.29-.368-.574-.572-.85m-3.707-3.429a1 1 0 0 0-1.045 1.705q.364.223.713.479A1 1 0 1 0 8.331 12.8q-.414-.302-.846-.567m10.06-2.167q-.364-.224-.715-.48a1 1 0 0 0-1.178 1.618q.415.302.847.567a1 1 0 1 0 1.047-1.705m1.447-5.065a9.9 9.9 0 1 0 0 14.001a9.913 9.913 0 0 0 0-14.001M17.72 17.729a8.03 8.03 0 0 1-4.516 2.273a.97.97 0 0 0-1.746.074a8.062 8.062 0 0 1-7.535-7.532a.975.975 0 0 0 .073-1.747a8.04 8.04 0 0 1 6.784-6.792a.997.997 0 0 0 .857.498a1.028 1.028 0 0 0 .23-.026a.982.982 0 0 0 .658-.546a8.054 8.054 0 0 1 7.538 7.538a.972.972 0 0 0-.074 1.741a8.046 8.046 0 0 1-2.27 4.519"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Basketball(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a10 10 0 0 0-8.4 4.59a10 10 0 0 0 14.66 13.2A10 10 0 0 0 12 2m2 2.26A8 8 0 0 1 19.74 10A9.78 9.78 0 0 0 15 11.38a15.7 15.7 0 0 0-2.4-2.21A10.06 10.06 0 0 0 14 4.26M12 4a7.9 7.9 0 0 1-1.14 4.07c-.15-.08-.29-.17-.44-.24a15.52 15.52 0 0 0-4.09-1.47A8 8 0 0 1 12 4M5 8.16A13.75 13.75 0 0 1 9.49 9.6l.13.08A7.93 7.93 0 0 1 4 12a8 8 0 0 1 1-3.84m5 11.58A8 8 0 0 1 4.26 14a9.9 9.9 0 0 0 7.08-3.21a14 14 0 0 1 2 1.8A10 10 0 0 0 10 19.74m2 .26a8 8 0 0 1 2.56-5.85c.06.08.12.15.17.23a14 14 0 0 1 1.84 4.18A7.93 7.93 0 0 1 12 20m6.21-3a16 16 0 0 0-1.8-3.75L16.2 13a7.93 7.93 0 0 1 3.8-1a8 8 0 0 1-1.79 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BasketballHoop(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 10H6a1 1 0 0 0 0 2h1v9a1 1 0 0 0 2 0v-1h2v1a1 1 0 0 0 2 0v-1h2v1a1 1 0 0 0 2 0v-9h1a1 1 0 0 0 0-2m-7 8H9v-2h2Zm0-4H9v-2h2Zm4 4h-2v-2h2Zm0-4h-2v-2h2Zm5-12H4a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3a1 1 0 0 0 0-2a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1a1 1 0 0 0 0 2a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3m-6 6a1 1 0 0 0 0-2h-4a1 1 0 0 0 0 2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bath(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 12H5V6.41a1.975 1.975 0 0 1 1.04-1.759a1.995 1.995 0 0 1 1.148-.23a3.491 3.491 0 0 0 .837 3.554l1.06 1.06a1 1 0 0 0 1.415 0L14.035 5.5a1 1 0 0 0 0-1.414l-1.06-1.06a3.494 3.494 0 0 0-4.53-.343A3.992 3.992 0 0 0 3 6.41V12H2a1 1 0 0 0 0 2h1v3a2.995 2.995 0 0 0 2 2.816V21a1 1 0 0 0 2 0v-1h10v1a1 1 0 0 0 2 0v-1.184A2.995 2.995 0 0 0 21 17v-3h1a1 1 0 0 0 0-2M9.44 4.44a1.502 1.502 0 0 1 2.12 0l.354.353l-2.121 2.121l-.354-.353a1.501 1.501 0 0 1 0-2.122ZM19 17a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1v-3h14Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryBolt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.69 15H4V9h2.31a1 1 0 0 0 0-2H4a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h2.69a1 1 0 1 0 0-2m7.2-2.56a1.27 1.27 0 0 0 .06-.18a1.42 1.42 0 0 0 0-.2v-.18a.65.65 0 0 0-.05-.2a.89.89 0 0 0-.08-.17a.86.86 0 0 0-.1-.16l-.16-.13l-.09-.09h-.06l-.18-.06h-3.5l1.45-2.5a1 1 0 1 0-1.74-1l-2.31 4v.06a1.27 1.27 0 0 0-.06.18a1.42 1.42 0 0 0 0 .2S7 12 7 12v.12a.65.65 0 0 0 .05.2a.89.89 0 0 0 .08.17a.86.86 0 0 0 .1.16l.16.13a.76.76 0 0 0 .09.09h.16A1 1 0 0 0 8 13h3.27l-1.45 2.5a1 1 0 0 0 1.74 1l2.31-4s.01-.04.02-.06M21 10a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0v-2a1 1 0 0 0-1-1m-4-3h-2.69a1 1 0 0 0 0 2H17v6h-2.31a1 1 0 1 0 0 2H17a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryEmpty(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 7H4a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h13a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2m0 8H4V9h13Zm4-5a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0v-2a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bed(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 12.5a3 3 0 1 0-3-3a3 3 0 0 0 3 3m0-4a1 1 0 1 1-1 1a1 1 0 0 1 1-1m13-2h-8a1 1 0 0 0-1 1v6H3v-8a1 1 0 0 0-2 0v13a1 1 0 0 0 2 0v-3h18v3a1 1 0 0 0 2 0v-9a3 3 0 0 0-3-3m1 7h-8v-5h7a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BedDouble(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 3.5H4a3 3 0 0 0-3 3v13a1 1 0 0 0 1 1h4a1 1 0 0 0 .83-.45l1.71-2.55h6.92l1.71 2.55a1 1 0 0 0 .83.45h4a1 1 0 0 0 1-1v-13a3 3 0 0 0-3-3m1 15h-2.46L16.83 16a1 1 0 0 0-.83-.5H8a1 1 0 0 0-.83.45L5.46 18.5H3v-5h18Zm-14-7v-1a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v1Zm6 0v-1a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v1Zm8 0h-2v-1a3 3 0 0 0-3-3h-2a3 3 0 0 0-2 .78a3 3 0 0 0-2-.78H8a3 3 0 0 0-3 3v1H3v-5a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Behance(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.07 6.35H15v1.41h5.09ZM19 16.05a2.23 2.23 0 0 1-1.3.37a2.23 2.23 0 0 1-1.7-.54a2.49 2.49 0 0 1-.62-1.76H22a6.47 6.47 0 0 0-.17-2a5.08 5.08 0 0 0-.8-1.73a4.17 4.17 0 0 0-1.42-1.21a4.37 4.37 0 0 0-2-.45a4.88 4.88 0 0 0-1.9.37a4.51 4.51 0 0 0-1.47 1a4.4 4.4 0 0 0-.95 1.52a5.4 5.4 0 0 0-.33 1.91a5.52 5.52 0 0 0 .32 1.94a4.46 4.46 0 0 0 .88 1.53a4 4 0 0 0 1.46 1a5.2 5.2 0 0 0 1.94.34a4.77 4.77 0 0 0 2.64-.7a4.21 4.21 0 0 0 1.63-2.35h-2.21a1.54 1.54 0 0 1-.62.76m-3.43-4.12a1.87 1.87 0 0 1 1-1.14a2.28 2.28 0 0 1 1-.2a1.73 1.73 0 0 1 1.36.49a2.91 2.91 0 0 1 .63 1.45h-4.15a3 3 0 0 1 .11-.6Zm-5.29-.48a3.06 3.06 0 0 0 1.28-1a2.72 2.72 0 0 0 .43-1.58a3.28 3.28 0 0 0-.29-1.48a2.4 2.4 0 0 0-.82-1a3.24 3.24 0 0 0-1.27-.52a7.54 7.54 0 0 0-1.64-.16H2v12.58h6.1a6.55 6.55 0 0 0 1.65-.21a4.55 4.55 0 0 0 1.43-.65a3.13 3.13 0 0 0 1-1.14a3.41 3.41 0 0 0 .37-1.65a3.47 3.47 0 0 0-.57-2a3 3 0 0 0-1.75-1.19ZM4.77 7.86h2.59a4.17 4.17 0 0 1 .71.06a1.64 1.64 0 0 1 .61.22a1.05 1.05 0 0 1 .42.44a1.42 1.42 0 0 1 .16.72a1.36 1.36 0 0 1-.47 1.15a2 2 0 0 1-1.22.35h-2.8Zm4.84 7.44a1.28 1.28 0 0 1-.45.5a2 2 0 0 1-.65.26a3.33 3.33 0 0 1-.78.08h-3v-3.45h3a2.4 2.4 0 0 1 1.45.41a1.65 1.65 0 0 1 .54 1.39a1.77 1.77 0 0 1-.11.81"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BehanceAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.108 15.032a.997.997 0 0 0-1.215.722A2.998 2.998 0 0 1 14.992 15v-1h7a1 1 0 0 0 1-1a5 5 0 1 0-10 0v2a4.998 4.998 0 0 0 9.839 1.247a.999.999 0 0 0-.723-1.215M17.992 10a3.011 3.011 0 0 1 2.118.873A3.044 3.044 0 0 1 20.809 12h-5.633a2.995 2.995 0 0 1 2.816-2m-2-3h4a1 1 0 0 0 0-2h-4a1 1 0 0 0 0 2M9.91 11.718A3.987 3.987 0 0 0 6.992 5h-5a1 1 0 0 0-1 1v13a1 1 0 0 0 1 1h5.5a4.492 4.492 0 0 0 2.418-8.282M2.992 7h4a2 2 0 1 1 0 4h-4Zm4.5 11h-4.5v-5h4.5a2.5 2.5 0 0 1 0 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bell(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 13.18V10a6 6 0 0 0-5-5.91V3a1 1 0 0 0-2 0v1.09A6 6 0 0 0 6 10v3.18A3 3 0 0 0 4 16v2a1 1 0 0 0 1 1h3.14a4 4 0 0 0 7.72 0H19a1 1 0 0 0 1-1v-2a3 3 0 0 0-2-2.82M8 10a4 4 0 0 1 8 0v3H8Zm4 10a2 2 0 0 1-1.72-1h3.44A2 2 0 0 1 12 20m6-3H6v-1a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellSchool(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 8a2 2 0 1 0-2.27 2a4.49 4.49 0 0 1-3 5.85a3 3 0 0 0-1.3-1.43a7 7 0 1 0-10.9 0A3 3 0 0 0 2 17v1a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3v-.19a6.47 6.47 0 0 0 4.58-8.59A2 2 0 0 0 21 8m-7 10a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-1a1 1 0 0 1 1-1h.41a6.94 6.94 0 0 0 7.18 0H13a1 1 0 0 1 1 1Zm-5-3a5 5 0 1 1 5-5a5 5 0 0 1-5 5m0-6a1 1 0 1 0 1 1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.07 6.12A3.48 3.48 0 0 1 12 6a4 4 0 0 1 4 4v1.34a1 1 0 0 0 2 0V10a6 6 0 0 0-5-5.91V3a1 1 0 0 0-2 0v1.1l-.45.08a1 1 0 0 0 .52 1.94m10.64 14.17l-18-18a1 1 0 0 0-1.42 1.42l4.12 4.11A6 6 0 0 0 6 10v3.18A3 3 0 0 0 4 16v2a1 1 0 0 0 1 1h3.14a4 4 0 0 0 7.72 0h1.73l2.7 2.71a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42M8 10a3.31 3.31 0 0 1 0-.55L11.59 13H8Zm4 10a2 2 0 0 1-1.72-1h3.44A2 2 0 0 1 12 20m-6-3v-1a1 1 0 0 1 1-1h6.59l2 2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bill(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.5 10.5H12a1 1 0 0 0 0-2h-1V8a1 1 0 0 0-2 0v.55a2.5 2.5 0 0 0 .5 4.95h1a.5.5 0 0 1 0 1H8a1 1 0 0 0 0 2h1v.5a1 1 0 0 0 2 0v-.55a2.5 2.5 0 0 0-.5-4.95h-1a.5.5 0 0 1 0-1M21 12h-3V3a1 1 0 0 0-.5-.87a1 1 0 0 0-1 0l-3 1.72l-3-1.72a1 1 0 0 0-1 0l-3 1.72l-3-1.72a1 1 0 0 0-1 0A1 1 0 0 0 2 3v16a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-6a1 1 0 0 0-1-1M5 20a1 1 0 0 1-1-1V4.73l2 1.14a1.08 1.08 0 0 0 1 0l3-1.72l3 1.72a1.08 1.08 0 0 0 1 0l2-1.14V19a3 3 0 0 0 .18 1Zm15-1a1 1 0 0 1-2 0v-5h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bing(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.1 8.6l1.7 4.3l2.8 1.3L9 17.5V3.4L5 2v17.8L9 22l10-5.8v-4.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bitcoin(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m11.136 12.117l-.596 2.415c.736.185 3.004.921 3.34-.441c.35-1.421-2.009-1.789-2.744-1.974m.813-3.297l-.54 2.191c.612.154 2.5.784 2.806-.455c.318-1.293-1.654-1.581-2.266-1.736M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m4.358 8.575a1.743 1.743 0 0 1-1.385 1.611a1.933 1.933 0 0 1 .997 2.661c-.586 1.692-1.977 1.835-3.827 1.481l-.449 1.82l-1.085-.274l.443-1.795a35.27 35.27 0 0 1-.864-.227l-.445 1.804l-1.084-.273l.45-1.824c-.254-.065-.511-.135-.774-.201l-1.412-.356l.539-1.256s.8.215.788.199a.394.394 0 0 0 .498-.26l1.217-4.939a.583.583 0 0 0-.505-.638c.016-.011-.789-.198-.789-.198l.29-1.172l1.495.378l-.001.006c.225.056.457.11.693.164l.444-1.801l1.085.273l-.436 1.766c.291.068.584.135.87.207l.432-1.755l1.085.274l-.445 1.802c1.37.477 2.372 1.193 2.175 2.523"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BitcoinAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.304 11.241a3.998 3.998 0 0 0-3.312-6.239v-1a1 1 0 1 0-2 0v1h-2v-1a1 1 0 1 0-2 0v1h-1a1 1 0 0 0 0 2h1v10h-1a1 1 0 0 0 0 2h1v1a1 1 0 0 0 2 0v-1h2v1a1 1 0 0 0 2 0v-1h2a3.99 3.99 0 0 0 1.312-7.76m-7.312-4.24h4a2 2 0 0 1 0 4h-4Zm6 10h-6v-4h6a2 2 0 1 1 0 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BitcoinCircle(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1a11 11 0 1 0 11 11A11 11 0 0 0 12 1m0 20a9 9 0 1 1 9-9a9 9 0 0 1-9 9m2-14V6a1 1 0 0 0-2 0v1h-1V6a1 1 0 0 0-2 0v1H8a1 1 0 0 0 0 2h1v6H8a1 1 0 0 0 0 2h1v1a1 1 0 0 0 2 0v-1h1v1a1 1 0 0 0 2 0v-1a3 3 0 0 0 3-3a3 3 0 0 0-.77-2a3 3 0 0 0 .77-2a3 3 0 0 0-3-3m0 8h-3v-2h3a1 1 0 0 1 0 2m0-4h-3V9h3a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BitcoinSign(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 9a4 4 0 0 0-4-4V3a1 1 0 0 0-2 0v2h-2V3a1 1 0 0 0-2 0v2H6a1 1 0 0 0 0 2h1v10H6a1 1 0 0 0 0 2h3v2a1 1 0 0 0 2 0v-2h2v2a1 1 0 0 0 2 0v-2a4 4 0 0 0 2.62-7A4 4 0 0 0 19 9m-4 8H9v-4h6a2 2 0 0 1 0 4m0-6H9V7h6a2 2 0 0 1 0 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BlackBerry(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2M7.19 13.35H5.27l.52-2.41h1.84c1.2 0 1.48.58 1.48 1.05c0 .65-.43 1.36-1.92 1.36m.66-3.64H5.94l.52-2.41H8.3c1.19 0 1.47.59 1.47 1c0 .7-.42 1.41-1.92 1.41m3.75 7.43H9.68l.53-2.42h1.84c1.19 0 1.47.59 1.47 1.06c0 .65-.42 1.36-1.92 1.36m.71-3.79H10.4l.52-2.41h1.84c1.19 0 1.47.58 1.47 1.05c0 .65-.42 1.36-1.92 1.36M13 9.71h-1.94l.52-2.41h1.84c1.2 0 1.48.59 1.48 1c0 .7-.43 1.41-1.9 1.41m3.74 5.61h-1.93l.52-2.42h1.84c1.19 0 1.48.59 1.48 1.06c0 .65-.43 1.36-1.93 1.36Zm.72-3.44h-1.94L16 9.46h1.84c1.2 0 1.48.59 1.48 1c.04.71-.39 1.42-1.88 1.42Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Blogger(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.71 10.46h2.239a.77.77 0 1 0 0-1.542H9.71a.77.77 0 1 0 0 1.542m9.718-8.458H4.555a2.573 2.573 0 0 0-2.563 2.563v14.873a2.573 2.573 0 0 0 2.563 2.564h14.873a2.573 2.573 0 0 0 2.564-2.564V4.565a2.573 2.573 0 0 0-2.564-2.563m-1.072 9.678l-.02 2.675a4.052 4.052 0 0 1-4.038 4.022h-4.64a4.05 4.05 0 0 1-4.041-4.025V9.657a4.053 4.053 0 0 1 4.042-4.03h2.823a4.46 4.46 0 0 1 3.12 2.2a3.156 3.156 0 0 1 .362 1.367c.068.506.103.881.33 1.09c.32.289 1.508.094 1.743.278l.178.14l.11.221l.035.178Zm-4.097 1.863H9.71a.75.75 0 1 0 0 1.5h4.549a.75.75 0 1 0 0-1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BloggerAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.991 13h-4a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2m-4-2h1a1 1 0 0 0 0-2h-1a1 1 0 0 0 0 2m6-1V9a4.004 4.004 0 0 0-4-4h-2a5.006 5.006 0 0 0-5 5v4a5.006 5.006 0 0 0 5 5h4a5.006 5.006 0 0 0 5-5v-1a3.005 3.005 0 0 0-3-3m0 0h-1Zm1 4a3.003 3.003 0 0 1-3 3h-4a3.003 3.003 0 0 1-3-3v-4a3.003 3.003 0 0 1 3-3h2a2.003 2.003 0 0 1 2 2v1a2.003 2.003 0 0 0 2 2a1 1 0 0 1 1 1Zm3-13h-16a3.003 3.003 0 0 0-3 3v16a3.003 3.003 0 0 0 3 3h16a3.003 3.003 0 0 0 3-3V4a3.003 3.003 0 0 0-3-3m1 19a1 1 0 0 1-1 1h-16a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BluetoothB(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.41 12l3.8-3.79a1 1 0 0 0 0-1.42l-4.5-4.5a1 1 0 0 0-.33-.21a1 1 0 0 0-.76 0a1 1 0 0 0-.54.54A1 1 0 0 0 11 3v6.59l-2.79-2.8a1 1 0 1 0-1.42 1.42l3.8 3.79l-3.8 3.79a1 1 0 1 0 1.42 1.42l2.79-2.8V21a1 1 0 0 0 .08.38a1 1 0 0 0 .54.54a.94.94 0 0 0 .76 0a1 1 0 0 0 .33-.21l4.5-4.5a1 1 0 0 0 0-1.42ZM13 5.41l2.09 2.09L13 9.59Zm0 13.18v-4.18l2.09 2.09Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bold(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.5 9.5A3.5 3.5 0 0 0 13 6H8.5a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1H13a3.49 3.49 0 0 0 2.44-6a3.5 3.5 0 0 0 1.06-2.5M13 16H9.5v-3H13a1.5 1.5 0 0 1 0 3m0-5H9.5V8H13a1.5 1.5 0 0 1 0 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bolt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.87 8.6A1 1 0 0 0 19 8h-4.58l1.27-4.74a1 1 0 0 0-.17-.87a1 1 0 0 0-.79-.39h-7a1 1 0 0 0-1 .74l-2.68 10a1 1 0 0 0 .17.87a1 1 0 0 0 .8.39h3.87l-1.81 6.74a1 1 0 0 0 1.71.93l10.9-12a1 1 0 0 0 .18-1.07m-9.79 8.68l1.07-4a1 1 0 0 0-.17-.87a1 1 0 0 0-.79-.39H6.35L8.49 4h4.93l-1.27 4.74a1 1 0 0 0 1 1.26h3.57Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoltAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.89 9.55A1 1 0 0 0 19 9h-5V3a1 1 0 0 0-.69-1a1 1 0 0 0-1.12.36l-8 11a1 1 0 0 0-.08 1A1 1 0 0 0 5 15h5v6a1 1 0 0 0 .69.95A1.12 1.12 0 0 0 11 22a1 1 0 0 0 .81-.41l8-11a1 1 0 0 0 .08-1.04M12 17.92V14a1 1 0 0 0-1-1H7l5-6.92V10a1 1 0 0 0 1 1h4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoltSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.67 4.23A1 1 0 0 0 9.6 4h5.1l-1.27 4.74a1 1 0 0 0 .17.87a1 1 0 0 0 .79.39H18l-1.13 1.24a1 1 0 0 0 .07 1.41a1 1 0 0 0 .67.26a1 1 0 0 0 .74-.33L21 9.67A1 1 0 0 0 20.23 8h-4.54L17 3.26a1 1 0 0 0-.18-.87A1 1 0 0 0 16 2H9a1 1 0 0 0-1 .74V3a1 1 0 0 0 .67 1.23m13 16.06l-18-18a1 1 0 0 0-1.38 1.42L6.61 8l-1.26 4.74a1 1 0 0 0 .18.87a1 1 0 0 0 .79.39h3.84l-1.81 6.74a1 1 0 0 0 .49 1.14a1 1 0 0 0 .48.12a1 1 0 0 0 .74-.33l4.85-5.34l5.38 5.38a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM7.62 12l.63-2.34L10.59 12Zm3.73 5.28l1-3.56l1.2 1.19Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Book(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 6H9a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1m-1 4h-4V8h4Zm3-8H5a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h12a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3m1 17a1 1 0 0 1-1 1H6V4h11a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 2H8a4 4 0 0 0-4 4v12a4 4 0 0 0 4 4h10a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2M6 6a2 2 0 0 1 2-2h10v10H8a3.91 3.91 0 0 0-2 .56Zm2 14a2 2 0 0 1 0-4h10v4Zm2-12h4a1 1 0 0 0 0-2h-4a1 1 0 0 0 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookMedical(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 9h-2V7a1 1 0 0 0-2 0v2H7a1 1 0 0 0 0 2h2v2a1 1 0 0 0 2 0v-2h2a1 1 0 0 0 0-2m5 6V5a3 3 0 0 0-3-3H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3M4 15V5a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1m17-9a1 1 0 0 0-1 1v10a3 3 0 0 1-3 3H7a1 1 0 0 0 0 2h10a5 5 0 0 0 5-5V7a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookOpen(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.17 2.06A13.1 13.1 0 0 0 19 1.87a12.94 12.94 0 0 0-7 2.05a12.94 12.94 0 0 0-7-2a13.1 13.1 0 0 0-2.17.19a1 1 0 0 0-.83 1v12a1 1 0 0 0 1.17 1a10.9 10.9 0 0 1 8.25 1.91l.12.07h.11a.91.91 0 0 0 .7 0h.11l.12-.07A10.9 10.9 0 0 1 20.83 16A1 1 0 0 0 22 15V3a1 1 0 0 0-.83-.94M11 15.35a12.87 12.87 0 0 0-6-1.48H4v-10a8.69 8.69 0 0 1 1 0a10.86 10.86 0 0 1 6 1.8Zm9-1.44h-1a12.87 12.87 0 0 0-6 1.48V5.67a10.86 10.86 0 0 1 6-1.8a8.69 8.69 0 0 1 1 0Zm1.17 4.15a13.1 13.1 0 0 0-2.17-.19a12.94 12.94 0 0 0-7 2.05a12.94 12.94 0 0 0-7-2.05a13.1 13.1 0 0 0-2.17.19A1 1 0 0 0 2 19.21a1 1 0 0 0 1.17.79a10.9 10.9 0 0 1 8.25 1.91a1 1 0 0 0 1.16 0A10.9 10.9 0 0 1 20.83 20a1 1 0 0 0 1.17-.79a1 1 0 0 0-.83-1.15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookReader(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.18 10.19A11.9 11.9 0 0 0 18 10c-.42 0-.83 0-1.24.08a5.91 5.91 0 0 0-1.91-1.65a3.81 3.81 0 0 0 1-2.57a3.86 3.86 0 0 0-7.72 0a3.81 3.81 0 0 0 1 2.57a6.11 6.11 0 0 0-1.91 1.64C6.83 10 6.42 10 6 10a11.9 11.9 0 0 0-2.18.21a1 1 0 0 0-.82 1v8.25a1 1 0 0 0 .36.77a1 1 0 0 0 .82.22A9.75 9.75 0 0 1 6 20.23a9.89 9.89 0 0 1 5.45 1.63l.13.05A1.09 1.09 0 0 0 12 22a.87.87 0 0 0 .28-.05h.07l.13-.05A9.89 9.89 0 0 1 18 20.23a9.75 9.75 0 0 1 1.82.18a1 1 0 0 0 .82-.22a1 1 0 0 0 .36-.77v-8.25a1 1 0 0 0-.82-.98M12 4a1.86 1.86 0 0 1 0 3.71A1.86 1.86 0 0 1 12 4m-1 15.33a11.92 11.92 0 0 0-5-1.1c-.33 0-.66 0-1 .05V12a9.63 9.63 0 0 1 2.52.05h.11A10 10 0 0 1 11 13.33Zm1-7.73a11.77 11.77 0 0 0-1.38-.68h-.06c-.33-.13-.66-.26-1-.36A4 4 0 0 1 12 9.69a4 4 0 0 1 2.44.85A12.43 12.43 0 0 0 12 11.6m7 6.68a11.6 11.6 0 0 0-6 1v-6a9.76 9.76 0 0 1 3.37-1.22h.2A9.39 9.39 0 0 1 19 12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookmark(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 2H8a3 3 0 0 0-3 3v16a1 1 0 0 0 .5.87a1 1 0 0 0 1 0l5.5-3.18l5.5 3.18a1 1 0 0 0 .5.13a1 1 0 0 0 .5-.13A1 1 0 0 0 19 21V5a3 3 0 0 0-3-3m1 17.27l-4.5-2.6a1 1 0 0 0-1 0L7 19.27V5a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkFull(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 2H6a1 1 0 0 0-1 1v18a1 1 0 0 0 1.65.76L12 17.27l5.29 4.44A1 1 0 0 0 18 22a.84.84 0 0 0 .38-.08A1 1 0 0 0 19 21V3a1 1 0 0 0-1-1m-1 16.86l-4.36-3.66a1 1 0 0 0-1.28 0L7 18.86V4h10Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Books(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m22.47 18.82l-1-3.86l-3.15-11.59a1 1 0 0 0-1.22-.71l-3.87 1a1 1 0 0 0-.73-.33h-10a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1v-8l2.2 8.22a1 1 0 0 0 1 .74a1.15 1.15 0 0 0 .26 0L21.79 20a1 1 0 0 0 .61-.47a1.05 1.05 0 0 0 .07-.71m-16 .55h-3v-2h3Zm0-4h-3v-6h3Zm0-8h-3v-2h3Zm5 12h-3v-2h3Zm0-4h-3v-6h3Zm0-8h-3v-2h3Zm2.25-1.74l2.9-.78l.52 1.93l-2.9.78Zm2.59 9.66l-1.55-5.8l2.9-.78l1.55 5.8Zm1 3.86l-.52-1.93l2.9-.78l.52 1.93Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Boombox(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2H7a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3m1 17a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1Zm-3.62-8.2A3 3 0 0 0 15 9a3 3 0 0 0-6 0a3 3 0 0 0 .62 1.8a4 4 0 1 0 4.76 0M12 8a1 1 0 1 1-1 1a1 1 0 0 1 1-1m0 8a2 2 0 1 1 2-2a2 2 0 0 1-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.5 18.5a1 1 0 1 0 1 1a1 1 0 0 0-1-1m9-15a1 1 0 0 0-1-1h-16a1 1 0 0 0-1 1v16a1 1 0 0 0 2 0v-15h15a1 1 0 0 0 1-1m-5 15a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-8 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1m12-12a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0 4a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0 4a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0 4a1 1 0 1 0 1 1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderBottom(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 13.5a1 1 0 1 0-1-1a1 1 0 0 0 1 1m0 4a1 1 0 1 0-1-1a1 1 0 0 0 1 1m0-8a1 1 0 1 0-1-1a1 1 0 0 0 1 1m-4-4a1 1 0 1 0-1-1a1 1 0 0 0 1 1m0 8a1 1 0 1 0-1-1a1 1 0 0 0 1 1m12-8a1 1 0 1 0-1-1a1 1 0 0 0 1 1m-4 8a1 1 0 1 0-1-1a1 1 0 0 0 1 1m-4-8a1 1 0 1 0-1-1a1 1 0 0 0 1 1m4 0a1 1 0 1 0-1-1a1 1 0 0 0 1 1m4 10a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0-4a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0 8H4a1 1 0 0 0 0 2h16a1 1 0 0 0 0-2m0-12a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-16 6a1 1 0 1 0-1-1a1 1 0 0 0 1 1m0-4a1 1 0 1 0-1-1a1 1 0 0 0 1 1m0 8a1 1 0 1 0-1-1a1 1 0 0 0 1 1m0-12a1 1 0 1 0-1-1a1 1 0 0 0 1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderClear(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 11a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0 8a1 1 0 1 0 1 1a1 1 0 0 0-1-1m4 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1M4 3a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0 8a1 1 0 1 0 1 1a1 1 0 0 0-1-1m16-6a1 1 0 1 0-1-1a1 1 0 0 0 1 1M4 19a1 1 0 1 0 1 1a1 1 0 0 0-1-1m8-4a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-8 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0-8a1 1 0 1 0 1 1a1 1 0 0 0-1-1m4-4a1 1 0 1 0 1 1a1 1 0 0 0-1-1m12 12a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0 4a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0-12a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-8 4a1 1 0 1 0 1 1a1 1 0 0 0-1-1m8 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-4-8a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-4 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0 4a1 1 0 1 0 1 1a1 1 0 0 0-1-1m4 12a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0-8a1 1 0 1 0 1 1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderHorizontal(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 5a1 1 0 1 0-1-1a1 1 0 0 0 1 1M4 9a1 1 0 1 0-1-1a1 1 0 0 0 1 1m8 0a1 1 0 1 0-1-1a1 1 0 0 0 1 1M4 5a1 1 0 1 0-1-1a1 1 0 0 0 1 1m12 0a1 1 0 1 0-1-1a1 1 0 0 0 1 1m-4 0a1 1 0 1 0-1-1a1 1 0 0 0 1 1m8 0a1 1 0 1 0-1-1a1 1 0 0 0 1 1M4 19a1 1 0 1 0 1 1a1 1 0 0 0-1-1m16 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0-4a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0-4H4a1 1 0 0 0 0 2h16a1 1 0 0 0 0-2m-4 8a1 1 0 1 0 1 1a1 1 0 0 0-1-1M4 15a1 1 0 1 0 1 1a1 1 0 0 0-1-1m16-8a1 1 0 1 0 1 1a1 1 0 0 0-1-1M8 19a1 1 0 1 0 1 1a1 1 0 0 0-1-1m4 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0-4a1 1 0 1 0 1 1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderInner(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 5a1 1 0 1 0-1-1a1 1 0 0 0 1 1M4 9a1 1 0 1 0-1-1a1 1 0 0 0 1 1m0-4a1 1 0 1 0-1-1a1 1 0 0 0 1 1m0 14a1 1 0 1 0 1 1a1 1 0 0 0-1-1M20 5a1 1 0 1 0-1-1a1 1 0 0 0 1 1m0 4a1 1 0 1 0-1-1a1 1 0 0 0 1 1m-4-4a1 1 0 1 0-1-1a1 1 0 0 0 1 1m4 14a1 1 0 1 0 1 1a1 1 0 0 0-1-1M4 15a1 1 0 1 0 1 1a1 1 0 0 0-1-1m16 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-4 4a1 1 0 1 0 1 1a1 1 0 0 0-1-1m5-7a1 1 0 0 0-1-1h-7V4a1 1 0 0 0-2 0v7H4a1 1 0 0 0 0 2h7v7a1 1 0 0 0 2 0v-7h7a1 1 0 0 0 1-1M8 19a1 1 0 1 0 1 1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderLeft(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.5 19a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0-4a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0-4a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-8-8a1 1 0 0 0-1 1v16a1 1 0 0 0 2 0V4a1 1 0 0 0-1-1m16 2a1 1 0 1 0-1-1a1 1 0 0 0 1 1m-8 2a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-4 4a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0 8a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0-16a1 1 0 1 0 1 1a1 1 0 0 0-1-1m12 8a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0 4a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-8-12a1 1 0 1 0 1 1a1 1 0 0 0-1-1m8 16a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0-12a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-4-4a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0 16a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0-8a1 1 0 1 0 1 1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderOut(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 9a1 1 0 1 0-1-1a1 1 0 0 0 1 1m0 4a1 1 0 1 0-1-1a1 1 0 0 0 1 1m4 0a1 1 0 1 0-1-1a1 1 0 0 0 1 1m-4 4a1 1 0 1 0-1-1a1 1 0 0 0 1 1m8-14H4a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1m-1 16H5V5h14ZM8 13a1 1 0 1 0-1-1a1 1 0 0 0 1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderRight(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 7a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0-4a1 1 0 1 0 1 1a1 1 0 0 0-1-1m4 16a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-4 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1m4-8a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0-8a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-4 8a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0 4a1 1 0 1 0 1 1a1 1 0 0 0-1-1m12-12a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0 8a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-4 8a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0-16a1 1 0 1 0 1 1a1 1 0 0 0-1-1m8 0a1 1 0 0 0-1 1v16a1 1 0 0 0 2 0V4a1 1 0 0 0-1-1m-4 16a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-4-12a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0 8a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0-4a1 1 0 1 0 1 1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderTop(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 18.5a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0-8a1 1 0 1 0 1 1a1 1 0 0 0-1-1m4 8a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-8-12a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0 12a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0-14h16a1 1 0 0 0 0-2H4a1 1 0 0 0 0 2m0 10a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0-4a1 1 0 1 0 1 1a1 1 0 0 0-1-1m8-4a1 1 0 1 0 1 1a1 1 0 0 0-1-1m8 8a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0 4a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-8-4a1 1 0 1 0 1 1a1 1 0 0 0-1-1m8-8a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0 4a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-8 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1m4 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0 8a1 1 0 1 0 1 1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BorderVertical(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 3a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0 4a1 1 0 1 0 1 1a1 1 0 0 0-1-1m16-2a1 1 0 1 0-1-1a1 1 0 0 0 1 1M7 11a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0 8a1 1 0 1 0 1 1a1 1 0 0 0-1-1M7 3a1 1 0 1 0 1 1a1 1 0 0 0-1-1M3 15a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0-4a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0 8a1 1 0 1 0 1 1a1 1 0 0 0-1-1m16-4a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0-4a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0 8a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0-12a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-4-4a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-4 0a1 1 0 0 0-1 1v16a1 1 0 0 0 2 0V4a1 1 0 0 0-1-1m4 16a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0-8a1 1 0 1 0 1 1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BowlingBall(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.992 8.002a1 1 0 1 0 1 1a1 1 0 0 0-1-1m4-6a10 10 0 1 0 10 10a10.012 10.012 0 0 0-10-10m0 18a8 8 0 1 1 8-8a8.01 8.01 0 0 1-8 8m-1-10a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0-4a1 1 0 1 0 1 1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Box(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.49 7.52a.19.19 0 0 1 0-.08a.17.17 0 0 1 0-.07v-.09l-.06-.15a.48.48 0 0 0-.09-.11l-.09-.08h-.05l-3.94-2.49l-3.72-2.3a.85.85 0 0 0-.29-.15h-.08a.82.82 0 0 0-.27 0h-.1a1.13 1.13 0 0 0-.33.13L4 6.78l-.09.07l-.09.08l-.1.07l-.05.06l-.06.15v.15a.69.69 0 0 0 0 .2v8.73a1 1 0 0 0 .47.85l7.5 4.64l.15.06h.08a.86.86 0 0 0 .52 0h.08l.15-.06L20 17.21a1 1 0 0 0 .47-.85V7.63s.02-.07.02-.11M12 4.17l1.78 1.1l-5.59 3.46l-1.79-1.1Zm-1 15l-5.5-3.36V9.42l5.5 3.4Zm1-8.11l-1.91-1.15l5.59-3.47l1.92 1.19Zm6.5 4.72L13 19.2v-6.38l5.5-3.4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BracketsCurly(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 6a2 2 0 0 1 2-2a1 1 0 0 0 0-2a4 4 0 0 0-4 4v3a2 2 0 0 1-2 2a1 1 0 0 0 0 2a2 2 0 0 1 2 2v3a4 4 0 0 0 4 4a1 1 0 0 0 0-2a2 2 0 0 1-2-2v-3a4 4 0 0 0-1.38-3A4 4 0 0 0 6 9Zm16 5a2 2 0 0 1-2-2V6a4 4 0 0 0-4-4a1 1 0 0 0 0 2a2 2 0 0 1 2 2v3a4 4 0 0 0 1.38 3A4 4 0 0 0 18 15v3a2 2 0 0 1-2 2a1 1 0 0 0 0 2a4 4 0 0 0 4-4v-3a2 2 0 0 1 2-2a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Brain(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 11a4 4 0 0 0-2-3.48A3 3 0 0 0 20 7a3 3 0 0 0-3-3h-.18A3 3 0 0 0 12 2.78A3 3 0 0 0 7.18 4H7a3 3 0 0 0-3 3a3 3 0 0 0 0 .52a4 4 0 0 0-.55 6.59A4 4 0 0 0 7 20h.18A3 3 0 0 0 12 21.22A3 3 0 0 0 16.82 20H17a4 4 0 0 0 3.5-5.89A4 4 0 0 0 22 11M11 8.55a4.72 4.72 0 0 0-.68-.32a1 1 0 0 0-.64 1.9A2 2 0 0 1 11 12v1.55a4.72 4.72 0 0 0-.68-.32a1 1 0 0 0-.64 1.9A2 2 0 0 1 11 17v2a1 1 0 0 1-1 1a1 1 0 0 1-.91-.6a4.07 4.07 0 0 0 .48-.33a1 1 0 1 0-1.28-1.54A2 2 0 0 1 7 18a2 2 0 0 1-2-2a2 2 0 0 1 .32-1.06A3.82 3.82 0 0 0 6 15a1 1 0 0 0 0-2a1.84 1.84 0 0 1-.69-.13A2 2 0 0 1 5 9.25a3.1 3.1 0 0 0 .46.35a1 1 0 1 0 1-1.74a.9.9 0 0 1-.34-.33A.92.92 0 0 1 6 7a1 1 0 0 1 1-1a.76.76 0 0 1 .21 0a3.85 3.85 0 0 0 .19.47a1 1 0 0 0 1.37.37a1 1 0 0 0 .36-1.34A1.06 1.06 0 0 1 9 5a1 1 0 0 1 2 0Zm7.69 4.32A1.84 1.84 0 0 1 18 13a1 1 0 0 0 0 2a3.82 3.82 0 0 0 .68-.06A2 2 0 0 1 19 16a2 2 0 0 1-2 2a2 2 0 0 1-1.29-.47a1 1 0 0 0-1.28 1.54a4.07 4.07 0 0 0 .48.33a1 1 0 0 1-.91.6a1 1 0 0 1-1-1v-2a2 2 0 0 1 1.32-1.87a1 1 0 0 0-.64-1.9a4.72 4.72 0 0 0-.68.32V12a2 2 0 0 1 1.32-1.87a1 1 0 0 0-.64-1.9a4.72 4.72 0 0 0-.68.32V5a1 1 0 0 1 2 0a1.06 1.06 0 0 1-.13.5a1 1 0 0 0 .36 1.37a1 1 0 0 0 1.37-.37a3.85 3.85 0 0 0 .19-.5a.76.76 0 0 1 .21 0a1 1 0 0 1 1 1a1 1 0 0 1-.17.55a.9.9 0 0 1-.33.31a1 1 0 0 0 1 1.74a2.66 2.66 0 0 0 .5-.35a2 2 0 0 1-.27 3.62Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Briefcase(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 6h-4V5a3 3 0 0 0-3-3h-4a3 3 0 0 0-3 3v1H3a1 1 0 0 0-1 1v4a3 3 0 0 0 1 2.22V19a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-5.78A3 3 0 0 0 22 11V7a1 1 0 0 0-1-1M9 5a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v1H9Zm10 14a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1v-5h2v1a1 1 0 0 0 2 0v-1h6v1a1 1 0 0 0 2 0v-1h2Zm1-8a1 1 0 0 1-1 1h-2v-1a1 1 0 0 0-2 0v1H9v-1a1 1 0 0 0-2 0v1H5a1 1 0 0 1-1-1V8h16Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BriefcaseAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 6.5h-3v-1a3 3 0 0 0-3-3h-2a3 3 0 0 0-3 3v1H5a3 3 0 0 0-3 3v9a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-9a3 3 0 0 0-3-3m-9-1a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v1h-4Zm10 13a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V13a21.27 21.27 0 0 0 3 .94v.59a1 1 0 0 0 2 0v-.21a23 23 0 0 0 3 .21a23 23 0 0 0 3-.21v.21a1 1 0 0 0 2 0v-.59a21.27 21.27 0 0 0 3-.94Zm0-7.69a20.39 20.39 0 0 1-3 1v-.31a1 1 0 0 0-2 0v.74a20.11 20.11 0 0 1-6 0v-.74a1 1 0 0 0-2 0v.33a20.39 20.39 0 0 1-3-1V9.5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bright(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 8a4 4 0 1 0 4 4a4 4 0 0 0-4-4m0 6a2 2 0 1 1 2-2a2 2 0 0 1-2 2m9.71-2.71L19.36 9V5.64a1 1 0 0 0-1-1h-3.31l-2.34-2.35a1 1 0 0 0-1.42 0L9 4.64H5.64a1 1 0 0 0-1 1V9l-2.35 2.29a1 1 0 0 0 0 1.42l2.35 2.34v3.31a1 1 0 0 0 1 1H9l2.34 2.35a1 1 0 0 0 1.42 0l2.34-2.35h3.31a1 1 0 0 0 1-1v-3.31l2.35-2.34a1 1 0 0 0-.05-1.42m-4.05 2.64a1 1 0 0 0-.3.71v2.72h-2.72a1 1 0 0 0-.71.3L12 19.59l-1.93-1.93a1 1 0 0 0-.71-.3H6.64v-2.72a1 1 0 0 0-.3-.71L4.41 12l1.93-1.93a1 1 0 0 0 .3-.71V6.64h2.72a1 1 0 0 0 .71-.3L12 4.41l1.93 1.93a1 1 0 0 0 .71.3h2.72v2.72a1 1 0 0 0 .3.71L19.59 12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Brightness(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 12a1 1 0 0 0-1-1H3a1 1 0 0 0 0 2h1a1 1 0 0 0 1-1m.64 5l-.71.71a1 1 0 0 0 0 1.41a1 1 0 0 0 1.41 0l.71-.71A1 1 0 0 0 5.64 17M12 5a1 1 0 0 0 1-1V3a1 1 0 0 0-2 0v1a1 1 0 0 0 1 1m5.66 2.34a1 1 0 0 0 .7-.29l.71-.71a1 1 0 1 0-1.41-1.41l-.66.71a1 1 0 0 0 0 1.41a1 1 0 0 0 .66.29m-12-.29a1 1 0 0 0 .7.29a1 1 0 0 0 .71-.29a1 1 0 0 0 0-1.41l-.71-.71a1 1 0 0 0-1.43 1.41ZM21 11h-1a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2m-2.64 6A1 1 0 0 0 17 18.36l.71.71a1 1 0 0 0 1.41 0a1 1 0 0 0 0-1.41ZM12 6.5a5.5 5.5 0 1 0 5.5 5.5A5.51 5.51 0 0 0 12 6.5m0 9a3.5 3.5 0 1 1 3.5-3.5a3.5 3.5 0 0 1-3.5 3.5m0 3.5a1 1 0 0 0-1 1v1a1 1 0 0 0 2 0v-1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrightnessEmpty(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.71 11.29L19.36 9V5.64a1 1 0 0 0-1-1h-3.31l-2.34-2.35a1 1 0 0 0-1.42 0L9 4.64H5.64a1 1 0 0 0-1 1V9l-2.35 2.29a1 1 0 0 0 0 1.42l2.35 2.34v3.31a1 1 0 0 0 1 1H9l2.34 2.35a1 1 0 0 0 1.42 0l2.34-2.35h3.31a1 1 0 0 0 1-1v-3.31l2.35-2.34a1 1 0 0 0-.05-1.42m-4.05 2.64a1 1 0 0 0-.3.71v2.72h-2.72a1 1 0 0 0-.71.3L12 19.59l-1.93-1.93a1 1 0 0 0-.71-.3H6.64v-2.72a1 1 0 0 0-.3-.71L4.41 12l1.93-1.93a1 1 0 0 0 .3-.71V6.64h2.72a1 1 0 0 0 .71-.3L12 4.41l1.93 1.93a1 1 0 0 0 .71.3h2.72v2.72a1 1 0 0 0 .3.71L19.59 12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrightnessHalf(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 8a1 1 0 0 0 0 2a2 2 0 0 1 0 4a1 1 0 0 0 0 2a4 4 0 0 0 0-8m9.71 3.29L19.36 9V5.64a1 1 0 0 0-1-1h-3.31l-2.34-2.35a1 1 0 0 0-1.42 0L9 4.64H5.64a1 1 0 0 0-1 1V9l-2.35 2.29a1 1 0 0 0 0 1.42l2.35 2.34v3.31a1 1 0 0 0 1 1H9l2.34 2.35a1 1 0 0 0 1.42 0l2.34-2.35h3.31a1 1 0 0 0 1-1v-3.31l2.35-2.34a1 1 0 0 0-.05-1.42m-4.05 2.64a1 1 0 0 0-.3.71v2.72h-2.72a1 1 0 0 0-.71.3L12 19.59l-1.93-1.93a1 1 0 0 0-.71-.3H6.64v-2.72a1 1 0 0 0-.3-.71L4.41 12l1.93-1.93a1 1 0 0 0 .3-.71V6.64h2.72a1 1 0 0 0 .71-.3L12 4.41l1.93 1.93a1 1 0 0 0 .71.3h2.72v2.72a1 1 0 0 0 .3.71L19.59 12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrightnessLow(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 11a1 1 0 1 0 1 1a1 1 0 0 0-1-1m1.93 6.66a1 1 0 1 0 1.41 0a1 1 0 0 0-1.41 0M6.34 6.34a1 1 0 1 0-1.41 0a1 1 0 0 0 1.41 0M12 4a1 1 0 1 0-1-1a1 1 0 0 0 1 1m5.66 13.66a1 1 0 1 0 1.41 0a1 1 0 0 0-1.41 0M21 11a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-3.34-6.07a1 1 0 1 0 1.41 0a1 1 0 0 0-1.41 0M12 20a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0-14a6 6 0 1 0 6 6a6 6 0 0 0-6-6m0 10a4 4 0 1 1 4-4a4 4 0 0 1-4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrightnessMinus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 11h-4a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2m7.71.29L19.36 9V5.64a1 1 0 0 0-1-1h-3.31l-2.34-2.35a1 1 0 0 0-1.42 0L9 4.64H5.64a1 1 0 0 0-1 1V9l-2.35 2.29a1 1 0 0 0 0 1.42l2.35 2.34v3.31a1 1 0 0 0 1 1H9l2.34 2.35a1 1 0 0 0 1.42 0l2.34-2.35h3.31a1 1 0 0 0 1-1v-3.31l2.35-2.34a1 1 0 0 0-.05-1.42m-4.05 2.64a1 1 0 0 0-.3.71v2.72h-2.72a1 1 0 0 0-.71.3L12 19.59l-1.93-1.93a1 1 0 0 0-.71-.3H6.64v-2.72a1 1 0 0 0-.3-.71L4.41 12l1.93-1.93a1 1 0 0 0 .3-.71V6.64h2.72a1 1 0 0 0 .71-.3L12 4.41l1.93 1.93a1 1 0 0 0 .71.3h2.72v2.72a1 1 0 0 0 .3.71L19.59 12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrightnessPlus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 11h-1v-1a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2h1v1a1 1 0 0 0 2 0v-1h1a1 1 0 0 0 0-2m7.71.29L19.36 9V5.64a1 1 0 0 0-1-1h-3.31l-2.34-2.35a1 1 0 0 0-1.42 0L9 4.64H5.64a1 1 0 0 0-1 1V9l-2.35 2.29a1 1 0 0 0 0 1.42l2.35 2.34v3.31a1 1 0 0 0 1 1H9l2.34 2.35a1 1 0 0 0 1.42 0l2.34-2.35h3.31a1 1 0 0 0 1-1v-3.31l2.35-2.34a1 1 0 0 0-.05-1.42m-4.05 2.64a1 1 0 0 0-.3.71v2.72h-2.72a1 1 0 0 0-.71.3L12 19.59l-1.93-1.93a1 1 0 0 0-.71-.3H6.64v-2.72a1 1 0 0 0-.3-.71L4.41 12l1.93-1.93a1 1 0 0 0 .3-.71V6.64h2.72a1 1 0 0 0 .71-.3L12 4.41l1.93 1.93a1 1 0 0 0 .71.3h2.72v2.72a1 1 0 0 0 .3.71L19.59 12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BringBottom(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 14h-2a1 1 0 1 0 0 2h1v4H4v-4h7a1 1 0 0 0 0-2h-1V9a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1M4 10h4v4H4Zm8.71-3.3L14 5.41V17a1 1 0 1 0 2 0V5.41l1.29 1.29A1 1 0 0 0 18 7a1 1 0 0 0 .71-.29a1 1 0 0 0 0-1.41l-3-3a1 1 0 0 0-1.42 0l-3 3a1 1 0 0 0 0 1.41a1 1 0 0 0 1.42-.01"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BringFront(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.29 17.3L10 18.59V7a1 1 0 0 0-2 0v11.6l-1.29-1.3a1 1 0 0 0-1.42 0a1 1 0 0 0 0 1.41l3 3a1 1 0 0 0 1.42 0l3-3a1 1 0 0 0 0-1.41a1 1 0 0 0-1.42 0M22 3a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h2a1 1 0 0 0 0-2H4V4h16v4h-7a1 1 0 0 0 0 2h1v5a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V9Zm-2 11h-4v-4h4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Browser(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 2H9a3 3 0 0 0-3 3v1H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-1h1a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3m-3 17a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-7h12Zm0-9H4V9a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1Zm4 5a1 1 0 0 1-1 1h-1V9a3 3 0 0 0-.18-1H20Zm0-9H8V5a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrushAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 18a1 1 0 1 0 1 1a1 1 0 0 0-1-1m6-17H6a1 1 0 0 0-1 1v9a3 3 0 0 0 3 3h1v2.37a4 4 0 1 0 6 0V14h1a3 3 0 0 0 3-3V2a1 1 0 0 0-1-1m-6 20a2 2 0 0 1-1.33-3.48a1 1 0 0 0 .33-.74V14h2v2.78a1 1 0 0 0 .33.74A2 2 0 0 1 12 21m5-10a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1v-1h10Zm0-3H7V3h10Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bug(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 14h2a1 1 0 0 0 0-2h-2v-1a5.15 5.15 0 0 0-.21-1.36A5 5 0 0 0 22 5a1 1 0 0 0-2 0a3 3 0 0 1-2.14 2.87A5 5 0 0 0 16 6.4a2.58 2.58 0 0 0 0-.4a4 4 0 0 0-8 0a2.58 2.58 0 0 0 0 .4a5 5 0 0 0-1.9 1.47A3 3 0 0 1 4 5a1 1 0 0 0-2 0a5 5 0 0 0 3.21 4.64A5.15 5.15 0 0 0 5 11v1H3a1 1 0 0 0 0 2h2v1a7 7 0 0 0 .14 1.38A5 5 0 0 0 2 21a1 1 0 0 0 2 0a3 3 0 0 1 1.81-2.74a7 7 0 0 0 12.38 0A3 3 0 0 1 20 21a1 1 0 0 0 2 0a5 5 0 0 0-3.14-4.62A7 7 0 0 0 19 15Zm-8 5.9A5 5 0 0 1 7 15v-4a3 3 0 0 1 3-3h1ZM10 6a2 2 0 0 1 4 0Zm7 9a5 5 0 0 1-4 4.9V8h1a3 3 0 0 1 3 3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Building(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 8h1a1 1 0 0 0 0-2h-1a1 1 0 0 0 0 2m0 4h1a1 1 0 0 0 0-2h-1a1 1 0 0 0 0 2M9 8h1a1 1 0 0 0 0-2H9a1 1 0 0 0 0 2m0 4h1a1 1 0 0 0 0-2H9a1 1 0 0 0 0 2m12 8h-1V3a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v17H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2m-8 0h-2v-4h2Zm5 0h-3v-5a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1v5H6V4h12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bullseye(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 5a7 7 0 1 0 7 7a7 7 0 0 0-7-7m0 12a5 5 0 1 1 5-5a5 5 0 0 1-5 5m0-8a3 3 0 1 0 3 3a3 3 0 0 0-3-3m0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1m0-12a11 11 0 1 0 11 11A11 11 0 0 0 12 1m0 20a9 9 0 1 1 9-9a9 9 0 0 1-9 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 17a1 1 0 0 0 1-1a1.36 1.36 0 0 0 0-.2a.64.64 0 0 0-.06-.18a.76.76 0 0 0-.09-.18l-.12-.15a1 1 0 0 0-.33-.21a1 1 0 0 0-.6-.08l-.18.06l-.18.09a1.58 1.58 0 0 0-.15.12l-.12.15a.76.76 0 0 0-.09.18a.64.64 0 0 0-.06.18a1.36 1.36 0 0 0 0 .2a1 1 0 0 0 1 1Zm8 0a1 1 0 0 0 1-1a1.36 1.36 0 0 0 0-.2a.64.64 0 0 0-.06-.18a.76.76 0 0 0-.09-.18l-.12-.15a1.15 1.15 0 0 0-.33-.21a1 1 0 0 0-.76 0a1.15 1.15 0 0 0-.33.21l-.12.15a.76.76 0 0 0-.09.18a.64.64 0 0 0-.06.18a1.36 1.36 0 0 0 0 .2a1 1 0 0 0 .29.7a1 1 0 0 0 .67.3m-3-12h-2a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2m5-3h-12a3 3 0 0 0-3 3v12a3 3 0 0 0 2 2.82V21a1 1 0 0 0 2 0v-1h10v1a1 1 0 0 0 2 0v-1.18a3 3 0 0 0 2-2.82V5a3 3 0 0 0-3-3m1 15a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-3h14Zm0-5h-14V5a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BusAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.5 2h-12a3 3 0 0 0-3 3v12a3 3 0 0 0 2 2.82V21a1 1 0 0 0 2 0v-1h10v1a1 1 0 0 0 2 0v-1.18a3 3 0 0 0 2-2.82V5a3 3 0 0 0-3-3m-13 6h6v4h-6Zm14 9a1 1 0 0 1-1 1h-12a1 1 0 0 1-1-1v-3h14Zm0-5h-6V8h6Zm0-6h-14V5a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1Zm-3.38 10.92a1 1 0 0 0 .38.08a1 1 0 0 0 1-1a1.36 1.36 0 0 0 0-.2a.64.64 0 0 0-.06-.18a.76.76 0 0 0-.09-.18a1.58 1.58 0 0 0-.12-.15l-.15-.12a.76.76 0 0 0-.18-.09a.64.64 0 0 0-.2-.08a1 1 0 0 0-.91.27a1.58 1.58 0 0 0-.12.15a.76.76 0 0 0-.09.18a.64.64 0 0 0-.06.18a1.36 1.36 0 0 0 0 .2a1 1 0 0 0 .29.7a1 1 0 0 0 .31.24m-8 0a1 1 0 0 0 .38.08a1 1 0 0 0 1-1a1.36 1.36 0 0 0 0-.2a.64.64 0 0 0-.06-.18a.76.76 0 0 0-.09-.18a1.58 1.58 0 0 0-.12-.15l-.15-.12l-.18-.09l-.2-.08a1 1 0 0 0-.91.27a1.58 1.58 0 0 0-.12.15a.76.76 0 0 0-.09.18a.64.64 0 0 0-.06.18a1.36 1.36 0 0 0 0 .2a1 1 0 0 0 .29.7a1 1 0 0 0 .31.24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BusSchool(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 12.5v4a1 1 0 0 0 1 1h1a3 3 0 0 0 6 0h6a3 3 0 0 0 6 0h1a1 1 0 0 0 1-1v-10a3 3 0 0 0-3-3H8.44A3 3 0 0 0 5.6 5.55L4.16 9.86l-2.71 1.81a1 1 0 0 0-.45.83m20-3h-2v-4h1a1 1 0 0 1 1 1Zm-4 8a1 1 0 1 1 1 1a1 1 0 0 1-1-1m-2-6h6v4h-.78a3 3 0 0 0-4.44 0H15Zm0-6h2v4h-2Zm-4 6h2v4h-2Zm0-6h2v4h-2Zm-2 4H6.39l1.1-3.32a1 1 0 0 1 .95-.68H9Zm-4 8a1 1 0 1 1 1 1a1 1 0 0 1-1-1M3 13l2.3-1.5H9v4h-.78a3 3 0 0 0-4.44 0H3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calculator(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.71 17.29a1 1 0 0 0-.16-.12a.56.56 0 0 0-.17-.09a.6.6 0 0 0-.19-.06a.93.93 0 0 0-.57.06a.9.9 0 0 0-.54.54a.84.84 0 0 0-.08.38a1 1 0 0 0 .07.38a1.46 1.46 0 0 0 .22.33A1 1 0 0 0 12 19a.84.84 0 0 0 .38-.08a1.15 1.15 0 0 0 .33-.21A1 1 0 0 0 13 18a1 1 0 0 0-.08-.38a1 1 0 0 0-.21-.33m-4.16-4.12a.56.56 0 0 0-.17-.09a.6.6 0 0 0-.19-.08a.86.86 0 0 0-.39 0l-.18.06l-.18.09l-.15.12A1.05 1.05 0 0 0 7 14a1 1 0 0 0 .29.71a1.15 1.15 0 0 0 .33.21A1 1 0 0 0 9 14a1.05 1.05 0 0 0-.29-.71Zm.16 4.12a1 1 0 0 0-.33-.21A1 1 0 0 0 7.8 17l-.18.06a.76.76 0 0 0-.18.09a1.58 1.58 0 0 0-.15.12a1 1 0 0 0-.21.33a.94.94 0 0 0 0 .76a1.15 1.15 0 0 0 .21.33A1 1 0 0 0 8 19a.84.84 0 0 0 .38-.08a1.15 1.15 0 0 0 .33-.21a1.15 1.15 0 0 0 .21-.33a.94.94 0 0 0 0-.76a1 1 0 0 0-.21-.33m2.91-4.21a1 1 0 0 0-.33.21A1.05 1.05 0 0 0 11 14a1 1 0 0 0 1.38.92a1.15 1.15 0 0 0 .33-.21A1 1 0 0 0 13 14a1.05 1.05 0 0 0-.29-.71a1 1 0 0 0-1.09-.21m5.09 4.21a1.15 1.15 0 0 0-.33-.21a1 1 0 0 0-1.09.21a1 1 0 0 0-.21.33a.94.94 0 0 0 0 .76a1.15 1.15 0 0 0 .21.33A1 1 0 0 0 16 19a.84.84 0 0 0 .38-.08a1.15 1.15 0 0 0 .33-.21a1 1 0 0 0 .21-1.09a1 1 0 0 0-.21-.33M16 5H8a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1m-1 4H9V7h6Zm3-8H6a3 3 0 0 0-3 3v16a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V4a3 3 0 0 0-3-3m1 19a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1Zm-2.45-6.83a.56.56 0 0 0-.17-.09a.6.6 0 0 0-.19-.06a.86.86 0 0 0-.39 0l-.18.06l-.18.09l-.15.12A1.05 1.05 0 0 0 15 14a1 1 0 0 0 1.38.92a1.15 1.15 0 0 0 .33-.21A1 1 0 0 0 17 14a1.05 1.05 0 0 0-.29-.71Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalculatorAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.5 8H6v.5a1 1 0 0 0 2 0V8h.5a1 1 0 0 0 0-2H8v-.5a1 1 0 0 0-2 0V6h-.5a1 1 0 0 0 0 2m-.62 11.12a1 1 0 0 0 1.41 0l.71-.71l.71.71a1 1 0 0 0 1.41 0a1 1 0 0 0 0-1.41L8.41 17l.71-.71a1 1 0 0 0-1.41-1.41l-.71.71l-.71-.71a1 1 0 0 0-1.41 1.41l.71.71l-.71.71a1 1 0 0 0 0 1.41M20 1H4a3 3 0 0 0-3 3v16a3 3 0 0 0 3 3h16a3 3 0 0 0 3-3V4a3 3 0 0 0-3-3m-9 20H4a1 1 0 0 1-1-1v-7h8Zm0-10H3V4a1 1 0 0 1 1-1h7Zm10 9a1 1 0 0 1-1 1h-7v-8h8Zm0-9h-8V3h7a1 1 0 0 1 1 1Zm-5.5 5.5h3a1 1 0 0 0 0-2h-3a1 1 0 0 0 0 2m3-10.5h-3a1 1 0 0 0 0 2h3a1 1 0 0 0 0-2m-3 13.5h3a1 1 0 0 0 0-2h-3a1 1 0 0 0 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 19a1 1 0 1 0-1-1a1 1 0 0 0 1 1m5 0a1 1 0 1 0-1-1a1 1 0 0 0 1 1m0-4a1 1 0 1 0-1-1a1 1 0 0 0 1 1m-5 0a1 1 0 1 0-1-1a1 1 0 0 0 1 1m7-12h-1V2a1 1 0 0 0-2 0v1H8V2a1 1 0 0 0-2 0v1H5a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3m1 17a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-9h16Zm0-11H4V6a1 1 0 0 1 1-1h1v1a1 1 0 0 0 2 0V5h8v1a1 1 0 0 0 2 0V5h1a1 1 0 0 1 1 1ZM7 15a1 1 0 1 0-1-1a1 1 0 0 0 1 1m0 4a1 1 0 1 0-1-1a1 1 0 0 0 1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.66 7H15v1a1 1 0 0 0 2 0V7h1a1 1 0 0 1 1 1v3h-1.34a1 1 0 0 0 0 2H19v1.34a1 1 0 1 0 2 0V8a3 3 0 0 0-3-3h-1V4a1 1 0 0 0-2 0v1h-3.34a1 1 0 0 0 0 2m10.05 13.29l-1.6-1.6l-16.4-16.4a1 1 0 0 0-1.42 1.42l1.91 1.9A3 3 0 0 0 3 8v10a3 3 0 0 0 3 3h12a3 3 0 0 0 1.29-.3l1 1a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.41M5 8a1 1 0 0 1 .66-.93L9.59 11H5Zm1 11a1 1 0 0 1-1-1v-5h6.59l6 6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calender(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 4h-2V3a1 1 0 0 0-2 0v1H9V3a1 1 0 0 0-2 0v1H5a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3m1 15a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-7h16Zm0-9H4V7a1 1 0 0 1 1-1h2v1a1 1 0 0 0 2 0V6h6v1a1 1 0 0 0 2 0V6h2a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calling(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.47 9a1 1 0 1 0-1-1a1 1 0 0 0 1 1m-3 0a1 1 0 1 0-1-1a1 1 0 0 0 1 1m-3 0a1 1 0 1 0-1-1a1 1 0 0 0 1 1m4.44 4c-.22 0-.45-.07-.67-.12a9.44 9.44 0 0 1-1.31-.39a2 2 0 0 0-2.48 1l-.22.45a12.48 12.48 0 0 1-2.67-2a12.83 12.83 0 0 1-2-2.66L10 9a2 2 0 0 0 1-2.48a10.33 10.33 0 0 1-.39-1.31c-.05-.23-.09-.45-.12-.68a3 3 0 0 0-3-2.49h-3a3 3 0 0 0-2.24 1a3 3 0 0 0-.73 2.39A19 19 0 0 0 18 21.91a2.56 2.56 0 0 0 .39 0a3 3 0 0 0 3-3v-3A3 3 0 0 0 18.91 13m.49 6a1 1 0 0 1-1.15 1a17.12 17.12 0 0 1-9.87-4.85a17.14 17.14 0 0 1-4.84-9.93a1 1 0 0 1 .25-.82a1 1 0 0 1 .74-.34h3a1 1 0 0 1 1 .79q.06.41.15.81a11.12 11.12 0 0 0 .46 1.55l-1.4.65a1 1 0 0 0-.49 1.33a14.49 14.49 0 0 0 7 7a1 1 0 0 0 .76 0a1 1 0 0 0 .56-.52l.63-1.4a13.69 13.69 0 0 0 1.58.46c.26.06.54.11.81.15a1 1 0 0 1 .78 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Camera(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 6.5h-1.28l-.32-1a3 3 0 0 0-2.84-2H9.44A3 3 0 0 0 6.6 5.55l-.32 1H5a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-8a3 3 0 0 0-3-3.05m1 11a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-8a1 1 0 0 1 1-1h2a1 1 0 0 0 1-.68l.54-1.64a1 1 0 0 1 .95-.68h5.12a1 1 0 0 1 .95.68l.54 1.64a1 1 0 0 0 .9.68h2a1 1 0 0 1 1 1Zm-8-9a4 4 0 1 0 4 4a4 4 0 0 0-4-4m0 6a2 2 0 1 1 2-2a2 2 0 0 1-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraChange(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12.29 5.21l1.5 1.5a1 1 0 0 0 1.42 0a1 1 0 0 0 .13-1.21H19a1 1 0 0 0 0-2h-3.66a1 1 0 0 0-.13-1.21a1 1 0 0 0-1.42 0l-1.5 1.5a1 1 0 0 0-.21.33a1 1 0 0 0 0 .76a1 1 0 0 0 .21.33m10.63 3.91a1 1 0 0 0-.21-.33l-1.5-1.5a1 1 0 0 0-1.42 0a1 1 0 0 0-.13 1.21H16a1 1 0 0 0 0 2h3.66a1 1 0 0 0 .13 1.21a1 1 0 0 0 1.42 0l1.5-1.5a1 1 0 0 0 .21-.33a1 1 0 0 0 0-.76M11 10a4 4 0 1 0 4 4a4 4 0 0 0-4-4m0 6a2 2 0 1 1 2-2a2 2 0 0 1-2 2m9-3a1 1 0 0 0-1 1v5a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1v-8a1 1 0 0 1 1-1h2a1 1 0 0 0 1-.69l.54-1.62a1 1 0 0 1 .9-.69H10a1 1 0 0 0 0-2H8.44a3 3 0 0 0-2.85 2.06L5.28 8H4a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-5a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraPlus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 10.5a1 1 0 0 0-1 1v7a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1v-8a1 1 0 0 1 1-1h2a1 1 0 0 0 1-.68l.54-1.64a1 1 0 0 1 .95-.68H14a1 1 0 0 0 0-2H8.44A3 3 0 0 0 5.6 6.55l-.32 1H4a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-7a1 1 0 0 0-1-1.05m-9-1a4 4 0 1 0 4 4a4 4 0 0 0-4-4m0 6a2 2 0 1 1 2-2a2 2 0 0 1-2 2m11-11h-1v-1a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2h1v1a1 1 0 0 0 2 0v-1h1a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.71 2.29a1 1 0 0 0-1.42 1.42l2.8 2.79H5a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h14.08l1.21 1.22a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Zm6.49 9.33l2.68 2.68a2 2 0 0 1-.88.2a2 2 0 0 1-2-2a2 2 0 0 1 .2-.88M5 18.5a1 1 0 0 1-1-1v-8a1 1 0 0 1 1-1h2.07l1.7 1.69A3.92 3.92 0 0 0 8 12.5a4 4 0 0 0 4 4a3.92 3.92 0 0 0 2.32-.77l2.77 2.77Zm14-12h-1.28l-.31-1a3 3 0 0 0-2.85-2h-4.4a1 1 0 0 0 0 2h4.4a1 1 0 0 1 .95.68l.54 1.63a1 1 0 0 0 .95.69h2a1 1 0 0 1 1 1v5.84a1 1 0 1 0 2 0V9.5a3 3 0 0 0-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cancel(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.77 11.36l-5-6A1 1 0 0 0 16 5H5a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h11a1 1 0 0 0 .77-.36l5-6a1 1 0 0 0 0-1.28M15.53 17H5a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1h10.53l4.17 5Zm-2.82-7.71a1 1 0 0 0-1.42 0L10 10.59l-1.29-1.3a1 1 0 1 0-1.42 1.42L8.59 12l-1.3 1.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l1.29-1.3l1.29 1.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L11.41 12l1.3-1.29a1 1 0 0 0 0-1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Capsule(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.5 4.5a5.12 5.12 0 0 0-7.24 0L4.5 12.26a5.12 5.12 0 1 0 7.24 7.24l7.76-7.76a5.12 5.12 0 0 0 0-7.24m-9.18 13.59a3.21 3.21 0 0 1-4.41 0a3.13 3.13 0 0 1 0-4.41l3.18-3.18l4.41 4.41Zm7.77-7.77l-3.18 3.18l-4.41-4.41l3.18-3.18a3.12 3.12 0 0 1 4.41 4.41"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Capture(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 9a1 1 0 0 0 1-1V5a1 1 0 0 1 1-1h3a1 1 0 0 0 0-2H5a3 3 0 0 0-3 3v3a1 1 0 0 0 1 1m5 11H5a1 1 0 0 1-1-1v-3a1 1 0 0 0-2 0v3a3 3 0 0 0 3 3h3a1 1 0 0 0 0-2m4-12a4 4 0 1 0 4 4a4 4 0 0 0-4-4m0 6a2 2 0 1 1 2-2a2 2 0 0 1-2 2m7-12h-3a1 1 0 0 0 0 2h3a1 1 0 0 1 1 1v3a1 1 0 0 0 2 0V5a3 3 0 0 0-3-3m2 13a1 1 0 0 0-1 1v3a1 1 0 0 1-1 1h-3a1 1 0 0 0 0 2h3a3 3 0 0 0 3-3v-3a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Car(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.62 13.08a.9.9 0 0 0-.54.54a1 1 0 0 0 1.3 1.3a1.15 1.15 0 0 0 .33-.21a1.15 1.15 0 0 0 .21-.33A.84.84 0 0 0 8 14a1.05 1.05 0 0 0-.29-.71a1 1 0 0 0-1.09-.21m13.14-4L18.4 5.05a3 3 0 0 0-2.84-2H8.44a3 3 0 0 0-2.84 2L4.24 9.11A3 3 0 0 0 2 12v4a3 3 0 0 0 2 2.82V20a1 1 0 0 0 2 0v-1h12v1a1 1 0 0 0 2 0v-1.18A3 3 0 0 0 22 16v-4a3 3 0 0 0-2.24-2.89ZM7.49 5.68A1 1 0 0 1 8.44 5h7.12a1 1 0 0 1 1 .68L17.61 9H6.39ZM20 16a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Zm-3.38-2.92a.9.9 0 0 0-.54.54a1 1 0 0 0 1.3 1.3a.9.9 0 0 0 .54-.54A.84.84 0 0 0 18 14a1.05 1.05 0 0 0-.29-.71a1 1 0 0 0-1.09-.21M13 13h-2a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CarSideview(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 9.5h-.32l-1.25-3.12a3 3 0 0 0-2.78-1.88h-6A3 3 0 0 0 5.7 6.91L5.18 9.5H5a3 3 0 0 0-3 3v3a1 1 0 0 0 1 1h1a3 3 0 0 0 6 0h4a3 3 0 0 0 6 0h1a1 1 0 0 0 1-1v-3a3 3 0 0 0-3-3m-6-3h1.65a1 1 0 0 1 .92.63l.95 2.37H13Zm-5.34.8a1 1 0 0 1 1-.8H11v3H7.22ZM7 17.5a1 1 0 1 1 1-1a1 1 0 0 1-1 1m10 0a1 1 0 1 1 1-1a1 1 0 0 1-1 1m3-3h-.78a3 3 0 0 0-4.44 0H9.22a3 3 0 0 0-4.44 0H4v-2a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CarSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 13a1 1 0 1 0 1 1a1 1 0 0 0-1-1m12.76-3.89l-1.35-4.06a3 3 0 0 0-2.85-2h-5.9a1 1 0 0 0 0 2h5.9a1 1 0 0 1 1 .69L17.61 9h-1.95a1 1 0 0 0 0 2H19a1 1 0 0 1 1 1v3.34a1 1 0 1 0 2 0V12a3 3 0 0 0-2.24-2.89m-16-6.82a1 1 0 0 0-1.47 1.42l2.82 2.81l-.87 2.59A3 3 0 0 0 2 12v4a3 3 0 0 0 2 2.82V20a1 1 0 0 0 2 0v-1h11.59l.41.41V20a1 1 0 0 0 1 1a.91.91 0 0 0 .46-.13l.83.84a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Zm3 5.81l.9.9H6.39ZM5 17a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1h4.59l2 2H11a1 1 0 0 0 0 2h2a.91.91 0 0 0 .46-.13L15.59 17Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CarWash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 4a1 1 0 0 0 .71-.29l1-1a1 1 0 0 0-1.42-1.42l-1 1a1 1 0 0 0 0 1.42A1 1 0 0 0 7.5 4m4 0a1 1 0 0 0 .71-.29l1-1a1 1 0 1 0-1.42-1.42l-1 1a1 1 0 0 0 0 1.42a1 1 0 0 0 .71.29m4 0a1 1 0 0 0 .71-.29l1-1a1 1 0 1 0-1.42-1.42l-1 1a1 1 0 0 0 0 1.42a1 1 0 0 0 .71.29m2.42 11.62a.76.76 0 0 0-.09-.18l-.12-.15l-.15-.12a.76.76 0 0 0-.18-.09a.6.6 0 0 0-.19-.06a1 1 0 0 0-.9.27a1.58 1.58 0 0 0-.12.15a.76.76 0 0 0-.09.18a.64.64 0 0 0-.06.18a1.36 1.36 0 0 0 0 .2a1 1 0 0 0 .29.7a.91.91 0 0 0 .33.22A1 1 0 0 0 17 17a1 1 0 0 0 1-1a1.36 1.36 0 0 0 0-.2a.64.64 0 0 0-.08-.18m1.84-4.51L18.4 7.05a3 3 0 0 0-2.84-2H8.44a3 3 0 0 0-2.84 2l-1.36 4.06A3 3 0 0 0 2 14v4a3 3 0 0 0 2 2.82V22a1 1 0 0 0 2 0v-1h12v1a1 1 0 0 0 2 0v-1.18A3 3 0 0 0 22 18v-4a3 3 0 0 0-2.24-2.89M7.49 7.68A1 1 0 0 1 8.44 7h7.12a1 1 0 0 1 1 .68L17.61 11H6.39ZM20 18a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Zm-7-3h-2a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2m-5.08.62a.76.76 0 0 0-.09-.18l-.12-.15a1 1 0 0 0-1.09-.21a1 1 0 0 0-.33.21a1.58 1.58 0 0 0-.12.15a.76.76 0 0 0-.09.18a.64.64 0 0 0-.08.18a1.36 1.36 0 0 0 0 .2a1 1 0 0 0 1.71.7A1 1 0 0 0 8 16a1.36 1.36 0 0 0 0-.2a.64.64 0 0 0-.08-.18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CardAtm(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 4.5H5a3 3 0 0 0-3 3v9a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-9a3 3 0 0 0-3-3m1 12a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-9a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Zm-4-6a3 3 0 0 0-1.51.42a3 3 0 1 0 0 5.16A3 3 0 1 0 16 10.5m-2.83 4a1 1 0 0 1-.17 0a1 1 0 0 1 0-2a1 1 0 0 1 .17 0a2.8 2.8 0 0 0 0 1.92Zm2.83 0a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretRight(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m19.5 11.13l-14-8.08a1 1 0 0 0-1 0a1 1 0 0 0-.5.87v16.16a1 1 0 0 0 .5.87a1 1 0 0 0 1 0l14-8.08a1 1 0 0 0 0-1.74M6 18.35V5.65L17 12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cell(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.49 4.73L17 2.17a1 1 0 0 0-1 0l-4 2.28l-4-2.28a1 1 0 0 0-1 0L2.51 4.73A1 1 0 0 0 2 5.6v5.12a1 1 0 0 0 .51.87l4 2.27v4.54a1 1 0 0 0 .51.87l4.5 2.56a1 1 0 0 0 1 0L17 19.27a1 1 0 0 0 .51-.87v-4.54l4-2.27a1 1 0 0 0 .51-.87V5.6a1 1 0 0 0-.53-.87M4 10.14v-4l3.5-2l3.5 2v4l-3.5 2Zm11.5 7.68l-3.5 2l-3.5-2v-4l3.5-2l3.5 2Zm4.5-7.68l-3.5 2l-3.5-2v-4l3.5-2l3.5 2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Celsius(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 19h-6a3 3 0 0 1-3-3V8a3 3 0 0 1 3-3h6a1 1 0 0 0 0-2h-6a5 5 0 0 0-5 5v8a5 5 0 0 0 5 5h6a1 1 0 0 0 0-2M5 3a3 3 0 1 0 3 3a3 3 0 0 0-3-3m0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Channel(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 16a3 3 0 0 0-1.73.56l-2.45-1.45A3.74 3.74 0 0 0 16 14a4 4 0 0 0-3-3.86V7.82a3 3 0 1 0-2 0v2.32A4 4 0 0 0 8 14a3.74 3.74 0 0 0 .18 1.11l-2.45 1.45A3 3 0 0 0 4 16a3 3 0 1 0 3 3a3 3 0 0 0-.12-.8l2.3-1.37a4 4 0 0 0 5.64 0l2.3 1.37A3 3 0 1 0 20 16M4 20a1 1 0 1 1 1-1a1 1 0 0 1-1 1m8-16a1 1 0 1 1-1 1a1 1 0 0 1 1-1m0 12a2 2 0 1 1 2-2a2 2 0 0 1-2 2m8 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChannelAdd(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 7h1v1a1 1 0 0 0 2 0V7h1a1 1 0 0 0 0-2h-1V4a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2m2 9a3 3 0 0 0-1.73.56l-2.45-1.45A3.74 3.74 0 0 0 16 14a4 4 0 0 0-3-3.86V7.82a3 3 0 1 0-2 0v2.32A4 4 0 0 0 8 14a3.74 3.74 0 0 0 .18 1.11l-2.45 1.45A3 3 0 0 0 4 16a3 3 0 1 0 3 3a3 3 0 0 0-.12-.8l2.3-1.37a4 4 0 0 0 5.64 0l2.3 1.37A3 3 0 1 0 20 16M4 20a1 1 0 1 1 1-1a1 1 0 0 1-1 1m8-16a1 1 0 1 1-1 1a1 1 0 0 1 1-1m0 12a2 2 0 1 1 2-2a2 2 0 0 1-2 2m8 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chart(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 6a1 1 0 0 0-1 1v10a1 1 0 0 0 2 0V7a1 1 0 0 0-1-1m-5 6a1 1 0 0 0-1 1v4a1 1 0 0 0 2 0v-4a1 1 0 0 0-1-1m10-2a1 1 0 0 0-1 1v6a1 1 0 0 0 2 0v-6a1 1 0 0 0-1-1m2-8H5a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3m1 17a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartBar(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 20h-1V5a1 1 0 0 0-2 0v15h-2v-7a1 1 0 0 0-2 0v7h-2V9a1 1 0 0 0-2 0v11H8v-3a1 1 0 0 0-2 0v3H4V3a1 1 0 0 0-2 0v18a1 1 0 0 0 1 1h18a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartBarAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 20H4v-2h3a1 1 0 0 0 0-2H4v-2h11a1 1 0 0 0 0-2H4v-2h7a1 1 0 0 0 0-2H4V6h15a1 1 0 0 0 0-2H4V3a1 1 0 0 0-2 0v18a1 1 0 0 0 1 1h18a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartDown(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 11a1 1 0 0 0-1 1v2.59l-6.29-6.3a1 1 0 0 0-1.42 0L9 11.59l-5.29-5.3a1 1 0 0 0-1.42 1.42l6 6a1 1 0 0 0 1.42 0l3.29-3.3L18.59 16H16a1 1 0 0 0 0 2h5a1 1 0 0 0 .38-.08a1 1 0 0 0 .54-.54A1 1 0 0 0 22 17v-5a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartGrowth(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 20h-1V5a1 1 0 0 0-2 0v15h-2V9a1 1 0 0 0-2 0v11h-2v-7a1 1 0 0 0-2 0v7H8v-3a1 1 0 0 0-2 0v3H4V3a1 1 0 0 0-2 0v18a1 1 0 0 0 1 1h18a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartGrowthAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 20H4v-2h3a1 1 0 0 0 0-2H4v-2h7a1 1 0 0 0 0-2H4v-2h11a1 1 0 0 0 0-2H4V6h15a1 1 0 0 0 0-2H4V3a1 1 0 0 0-2 0v18a1 1 0 0 0 1 1h18a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartLine(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 16a1.5 1.5 0 0 0 1.5-1.5a.77.77 0 0 0 0-.15l2.79-2.79h.46l1.61 1.61v.08a1.5 1.5 0 1 0 3 0v-.08L20 9.5A1.5 1.5 0 1 0 18.5 8a.77.77 0 0 0 0 .15l-3.61 3.61h-.16L13 10a1.49 1.49 0 0 0-3 0l-3 3a1.5 1.5 0 0 0 0 3m13.5 4h-17V3a1 1 0 0 0-2 0v18a1 1 0 0 0 1 1h18a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartPie(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a10 10 0 1 0 4.93 18.69h.12A10 10 0 0 0 12 2m1 2.07A8 8 0 0 1 19.93 11H13ZM12 20a8 8 0 0 1-1-15.93V12a1.09 1.09 0 0 0 .07.35v.15l4 6.87A7.81 7.81 0 0 1 12 20m4.83-1.64L13.73 13h6.2a8 8 0 0 1-3.1 5.36"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartPieAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 12h-7V5c0-.6-.4-1-1-1c-5 0-9 4-9 9s4 9 9 9s9-4 9-9c0-.6-.4-1-1-1m-7 7.9c-3.8.6-7.4-2.1-7.9-5.9c-.6-3.8 2.1-7.4 5.9-7.9V13c0 .6.4 1 1 1h6.9c-.4 3.1-2.8 5.5-5.9 5.9M15 2c-.6 0-1 .4-1 1v6c0 .6.4 1 1 1h6c.6 0 1-.4 1-1c0-3.9-3.1-7-7-7m1 6V4.1c2 .4 3.5 1.9 3.9 3.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chat(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 2H6a3 3 0 0 0-3 3v11a3 3 0 0 0 3 3h2.59l2.7 2.71A1 1 0 0 0 12 22a1 1 0 0 0 .65-.24L15.87 19H18a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3m1 14a1 1 0 0 1-1 1h-2.5a1 1 0 0 0-.65.24l-2.8 2.4l-2.34-2.35A1 1 0 0 0 9 17H6a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatBubbleUser(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.58 11.3a3.24 3.24 0 0 0 .71-2a3.29 3.29 0 0 0-6.58 0a3.24 3.24 0 0 0 .71 2a5 5 0 0 0-2 2.31a1 1 0 1 0 1.84.78A3 3 0 0 1 12 12.57a3 3 0 0 1 2.75 1.82a1 1 0 0 0 .92.61a1.09 1.09 0 0 0 .39-.08a1 1 0 0 0 .53-1.31a5 5 0 0 0-2.01-2.31M12 10.57a1.29 1.29 0 1 1 1.29-1.28A1.29 1.29 0 0 1 12 10.57M18 2H6a3 3 0 0 0-3 3v11a3 3 0 0 0 3 3h2.59l2.7 2.71A1 1 0 0 0 12 22a1 1 0 0 0 .65-.24L15.87 19H18a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3m1 14a1 1 0 0 1-1 1h-2.5a1 1 0 0 0-.65.24l-2.8 2.4l-2.34-2.35A1 1 0 0 0 9 17H6a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatInfo(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.29 3.71a1 1 0 0 0 1.42 0a1.15 1.15 0 0 0 .21-.33A1 1 0 0 0 21 3a1 1 0 0 0-.29-.71l-.15-.12a.76.76 0 0 0-.18-.09a1 1 0 0 0-1.09.21A1 1 0 0 0 19 3a1 1 0 0 0 .08.38a1.15 1.15 0 0 0 .21.33M20 5a1 1 0 0 0-1 1v4a1 1 0 0 0 2 0V6a1 1 0 0 0-1-1m.06 8a1 1 0 0 0-1.11.87A7 7 0 0 1 12 20H6.41l.64-.63a1 1 0 0 0 0-1.41A7 7 0 0 1 12 6a6.91 6.91 0 0 1 3.49.94a1 1 0 0 0 1-1.72A8.84 8.84 0 0 0 12 4a9 9 0 0 0-7 14.62l-1.71 1.67a1 1 0 0 0-.21 1.09A1 1 0 0 0 4 22h8a9 9 0 0 0 8.93-7.88a1 1 0 0 0-.87-1.12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Check(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.71 7.21a1 1 0 0 0-1.42 0l-7.45 7.46l-3.13-3.14A1 1 0 1 0 5.29 13l3.84 3.84a1 1 0 0 0 1.42 0l8.16-8.16a1 1 0 0 0 0-1.47"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckCircle(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14.72 8.79l-4.29 4.3l-1.65-1.65a1 1 0 1 0-1.41 1.41l2.35 2.36a1 1 0 0 0 .71.29a1 1 0 0 0 .7-.29l5-5a1 1 0 0 0 0-1.42a1 1 0 0 0-1.41 0M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckSquare(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.21 14.75a1 1 0 0 0 1.42 0l4.08-4.08a1 1 0 0 0-1.42-1.42l-3.37 3.38l-1.21-1.22a1 1 0 0 0-1.42 1.42ZM21 2H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1m-1 18H4V4h16Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Circle(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleLayer(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 2a7 7 0 0 0-6.88 5.74A6 6 0 0 0 5 12.41A5 5 0 1 0 11.59 19a6 6 0 0 0 4.67-3.09A7 7 0 0 0 15 2m-5 15.43c0 .1 0 .2-.07.31a3 3 0 1 1-3.64-3.64l.28-.1A2.94 2.94 0 0 1 10 17.43m5-3.67a1.8 1.8 0 0 1-.05.19a3.74 3.74 0 0 1-.17.54a4 4 0 0 1-2.7 2.4c0-.18 0-.35-.06-.53s0-.27 0-.4s-.12-.38-.18-.57s-.07-.24-.12-.36a4.21 4.21 0 0 0-.3-.55c0-.09-.09-.19-.15-.28a5.3 5.3 0 0 0-.6-.73l-.2-.17a5.52 5.52 0 0 0-.53-.43a2.9 2.9 0 0 0-.34-.19a4 4 0 0 0-.5-.27a2.34 2.34 0 0 0-.4-.13a3.13 3.13 0 0 0-.52-.16c-.14 0-.29 0-.44-.06L7.13 12a4 4 0 0 1 2.39-2.7a3.27 3.27 0 0 1 .53-.17l.2-.05A3.74 3.74 0 0 1 11 9a4 4 0 0 1 4 4a3.84 3.84 0 0 1-.08.76Zm2-.16V13a6 6 0 0 0-6-6h-.6a5 5 0 1 1 6.6 6.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Circuit(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 13a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0-4a1 1 0 1 0 1 1a1 1 0 0 0-1-1m4 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1m7 4a1 1 0 0 0 0-2h-2V9h2a1 1 0 0 0 0-2h-2.18A3 3 0 0 0 17 5.18V3a1 1 0 0 0-2 0v2h-2V3a1 1 0 0 0-2 0v2H9V3a1 1 0 0 0-2 0v2.18A3 3 0 0 0 5.18 7H3a1 1 0 0 0 0 2h2v2H3a1 1 0 0 0 0 2h2v2H3a1 1 0 0 0 0 2h2.18A3 3 0 0 0 7 18.82V21a1 1 0 0 0 2 0v-2h2v2a1 1 0 0 0 2 0v-2h2v2a1 1 0 0 0 2 0v-2.18A3 3 0 0 0 18.82 17H21a1 1 0 0 0 0-2h-2v-2Zm-4 3a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1Zm-3-3a1 1 0 1 0 1 1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClapperBoard(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 2H5a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3m-2.91 2l-4 4H7.91l4-4ZM4 5a1 1 0 0 1 1-1h4.09l-4 4H4Zm16 14a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-9h16Zm0-11h-5.09l4-4H19a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClinicMedical(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 12v1h-1a1 1 0 0 0 0 2h1v1a1 1 0 0 0 2 0v-1h1a1 1 0 0 0 0-2h-1v-1a1 1 0 0 0-2 0m10.664-1.748l-9-8a.999.999 0 0 0-1.328 0l-9 8a1 1 0 0 0 1.328 1.496L4 11.449V21a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1v-9.551l.336.299a1 1 0 0 0 1.328-1.496M18 20H6V9.671l6-5.333l6 5.333Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clipboard(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 4h-1.18A3 3 0 0 0 13 2h-2a3 3 0 0 0-2.82 2H7a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3m-7 1a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v1h-4Zm8 14a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1h1v1a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V6h1a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 14H9a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2m0-4h-4a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2m2-6h-1.18A3 3 0 0 0 13 2h-2a3 3 0 0 0-2.82 2H7a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3m-7 1a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v1h-4Zm8 14a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1h1v1a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V6h1a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardBlank(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 4h-1a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2H7a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3m-7 0h4v2h-4zm8 15a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1h1a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2h1a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardNotes(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 14H9a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2m4-10h-1.18A3 3 0 0 0 13 2h-2a3 3 0 0 0-2.82 2H7a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3m-7 1a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v1h-4Zm8 14a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1h1v1a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V6h1a1 1 0 0 1 1 1Zm-3-9H9a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clock(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.098 12.634L13 11.423V7a1 1 0 0 0-2 0v5a1 1 0 0 0 .5.866l2.598 1.5a1 1 0 1 0 1-1.732M12 2a10 10 0 1 0 10 10A10.011 10.011 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8.01 8.01 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockEight(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 6a1 1 0 0 0-1 1v4.384l-2.43 1.223a1 1 0 1 0 .898 1.786l2.981-1.5A.999.999 0 0 0 13 12V7a1 1 0 0 0-1-1m0-4a10 10 0 1 0 10 10A10.011 10.011 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8.01 8.01 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockFive(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10.011 10.011 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8.01 8.01 0 0 1-8 8m1-8.251V7a1 1 0 0 0-2 0v5a1.006 1.006 0 0 0 .118.472l1.5 2.799a1 1 0 0 0 1.764-.944Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockNine(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 6a1 1 0 0 0-1 1v4H9a1 1 0 0 0 0 2h3a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1m0-4a10 10 0 1 0 10 10A10.011 10.011 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8.01 8.01 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockSeven(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10.011 10.011 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8.01 8.01 0 0 1-8 8m0-14a1 1 0 0 0-1 1v4.749l-1.382 2.578a1 1 0 0 0 1.764.944l1.5-2.799A1.006 1.006 0 0 0 13 12V7a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockTen(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 6a1 1 0 0 0-1 1v3.268l-1.098-.634a1 1 0 0 0-1 1.732l2.598 1.5A1 1 0 0 0 13 12V7a1 1 0 0 0-1-1m0-4a10 10 0 1 0 10 10A10.011 10.011 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8.01 8.01 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockThree(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 11h-2V7a1 1 0 0 0-2 0v5a1 1 0 0 0 1 1h3a1 1 0 0 0 0-2m-3-9a10 10 0 1 0 10 10A10.011 10.011 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8.01 8.01 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockTwo(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10.011 10.011 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8.01 8.01 0 0 1-8 8m2.098-10.366L13 10.268V7a1 1 0 0 0-2 0v5a1 1 0 0 0 1.5.866l2.598-1.5a1 1 0 1 0-1-1.732"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClosedCaptioning(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.24 13.14a1 1 0 0 0-1.37.36a1 1 0 0 1-1.58.19A.93.93 0 0 1 8 13v-2a1 1 0 0 1 1.88-.48a1 1 0 0 0 1.37.34a1 1 0 0 0 .34-1.38a3.08 3.08 0 0 0-.46-.59A3 3 0 0 0 9 8a3 3 0 0 0-3 3v2a3 3 0 0 0 5.59 1.5a1 1 0 0 0-.35-1.36m6 0a1 1 0 0 0-1.37.36a1 1 0 0 1-1.58.19A.93.93 0 0 1 14 13v-2a1 1 0 0 1 1.88-.48a1 1 0 0 0 1.37.34a1 1 0 0 0 .34-1.38a3.08 3.08 0 0 0-.46-.59A3 3 0 0 0 15 8a3 3 0 0 0-3 3v2a3 3 0 0 0 5.59 1.5a1 1 0 0 0-.35-1.36M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClosedCaptioningSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 0 1-8-8a7.92 7.92 0 0 1 1.69-4.9L7.2 8.61A3 3 0 0 0 6 11v2a3 3 0 0 0 5.59 1.5a1 1 0 1 0-1.72-1a1 1 0 0 1-1.58.19A.93.93 0 0 1 8 13v-2a1 1 0 0 1 .67-.92L12 13.46A3 3 0 0 0 14.54 16l2.36 2.36A7.92 7.92 0 0 1 12 20m6.31-3.1l-1.52-1.52a2.94 2.94 0 0 0 .8-.88a1 1 0 1 0-1.72-1a1 1 0 0 1-.55.41L14 12.59V11a1 1 0 0 1 1.88-.48a1 1 0 0 0 1.37.34a1 1 0 0 0 .34-1.38a3.08 3.08 0 0 0-.46-.59A3 3 0 0 0 12 10.62l-.35-.35a1 1 0 0 0-.1-.79a3.08 3.08 0 0 0-.46-.59a2.94 2.94 0 0 0-1.67-.84L7.1 5.69A7.92 7.92 0 0 1 12 4a8 8 0 0 1 8 8a7.92 7.92 0 0 1-1.69 4.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloud(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.42 9.21a7 7 0 0 0-13.36 1.9A4 4 0 0 0 6 19h11a5 5 0 0 0 1.42-9.79M17 17H6a2 2 0 0 1 0-4a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.61a1 1 0 0 0 .78.66A3 3 0 0 1 17 17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudBlock(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.42 7.72A7 7 0 0 0 5.06 9.61a4 4 0 0 0-.38 7.66a1.13 1.13 0 0 0 .32.05a1 1 0 0 0 .32-2A2 2 0 0 1 4 13.5a2 2 0 0 1 2-2a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.61a1 1 0 0 0 .78.67a3 3 0 0 1 1 5.53a1 1 0 1 0 1 1.74A5 5 0 0 0 22 12.5a5 5 0 0 0-3.58-4.78m-9.25 6a4 4 0 1 0 5.66 0a4.1 4.1 0 0 0-5.66-.05ZM10 16.5a2 2 0 0 1 2-2a2.09 2.09 0 0 1 .51.07L10.07 17a2.09 2.09 0 0 1-.07-.5m3.41 1.41a2 2 0 0 1-1.91.5L13.93 16a2.09 2.09 0 0 1 .07.51a2 2 0 0 1-.59 1.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudBookmark(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 11h-5a1 1 0 0 0-1 1v8a1 1 0 0 0 1.56.83l1.94-1.3l1.89 1.26A1 1 0 0 0 15 21a1 1 0 0 0 .44-.1A1 1 0 0 0 16 20v-8a1 1 0 0 0-1-1m-1 7.12l-.94-.63a1 1 0 0 0-1.12 0l-.94.64V13h3Zm4.42-10.9A7 7 0 0 0 5.06 9.11A4 4 0 0 0 6 17a1 1 0 0 0 0-2a2 2 0 0 1 0-4a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.61a1 1 0 0 0 .78.67a3 3 0 0 1 1 5.53a1 1 0 1 0 1 1.74A5 5 0 0 0 22 12a5 5 0 0 0-3.58-4.78"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudCheck(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.29 14.19L11 17.48l-1.29-1.29a1 1 0 0 0-1.42 0a1 1 0 0 0 0 1.41l2 2a1 1 0 0 0 1.42 0l4-4a1 1 0 0 0 0-1.41a1 1 0 0 0-1.42 0m4.13-5.87a7 7 0 0 0-13.36 1.9a4 4 0 0 0-.38 7.65A1 1 0 1 0 5.32 16A2 2 0 0 1 4 14.1a2 2 0 0 1 2-2a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.6a1 1 0 0 0 .78.66a3 3 0 0 1 .24 5.84a1 1 0 0 0 .25 2h.25a5 5 0 0 0 .17-9.62Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudComputing(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 20h-3a1 1 0 0 1-1-1v-3a5 5 0 0 0 1.42-9.78A7 7 0 0 0 5.06 8.11A4 4 0 0 0 6 16h1v3a1 1 0 0 1-1 1H3a1 1 0 0 0 0 2h3a3 3 0 0 0 3-3v-3h2v5a1 1 0 0 0 2 0v-5h2v3a3 3 0 0 0 3 3h3a1 1 0 0 0 0-2M6 14a2 2 0 0 1 0-4a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.61a1 1 0 0 0 .78.67A3 3 0 0 1 20 11a3 3 0 0 1-3 3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudDataConnection(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 18.5h-6.18A3 3 0 0 0 13 16.68V13.5h3.17a4.33 4.33 0 0 0 1.3-8.5A6 6 0 0 0 6.06 6.63A3.5 3.5 0 0 0 7 13.5h4v3.18a3 3 0 0 0-1.82 1.82H3a1 1 0 0 0 0 2h6.18a3 3 0 0 0 5.64 0H21a1 1 0 0 0 0-2m-14-7a1.5 1.5 0 0 1 0-3a1 1 0 0 0 1-1a4 4 0 0 1 7.79-1.29a1 1 0 0 0 .78.67a2.31 2.31 0 0 1 1.93 2.29a2.34 2.34 0 0 1-2.33 2.33Zm5 9a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudDatabaseTree(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.5 14.5a2 2 0 0 0-2-2h-6a2 2 0 0 0-2 2h-3v-3h2.33A3.66 3.66 0 0 0 13 4.37a5 5 0 0 0-9.43 1.28a3 3 0 0 0 .93 5.85h3v8a1 1 0 0 0 1 1h4a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2v-2a2 2 0 0 0-.28-1a2 2 0 0 0 .28-1Zm-18-5a1 1 0 0 1 0-2a1 1 0 0 0 1-1a3 3 0 0 1 5.84-1a1 1 0 0 0 .78.66a1.65 1.65 0 0 1 1.38 1.67a1.67 1.67 0 0 1-1.67 1.67Zm8 9h-3v-2h3a2 2 0 0 0 .28 1a2 2 0 0 0-.28 1m2 2v-2h6v2Zm0-4v-2h6v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudDownload(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.29 17.29L13 18.59V13a1 1 0 0 0-2 0v5.59l-1.29-1.3a1 1 0 0 0-1.42 1.42l3 3a1 1 0 0 0 .33.21a.94.94 0 0 0 .76 0a1 1 0 0 0 .33-.21l3-3a1 1 0 0 0-1.42-1.42m4.13-11.07A7 7 0 0 0 5.06 8.11A4 4 0 0 0 6 16a1 1 0 0 0 0-2a2 2 0 0 1 0-4a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.61a1 1 0 0 0 .78.67a3 3 0 0 1 .24 5.84a1 1 0 1 0 .5 1.94a5 5 0 0 0 .17-9.62"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudDrizzle(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 19a1 1 0 0 0-1 1v1a1 1 0 0 0 2 0v-1a1 1 0 0 0-1-1m-4-8a1 1 0 0 0-1 1v1a1 1 0 0 0 2 0v-1a1 1 0 0 0-1-1m0 5a1 1 0 0 0-1 1v1a1 1 0 0 0 2 0v-1a1 1 0 0 0-1-1m4-2a1 1 0 0 0-1 1v1a1 1 0 0 0 2 0v-1a1 1 0 0 0-1-1m6.42-7.78A7 7 0 0 0 5.06 8.11A4 4 0 0 0 2 12a4 4 0 0 0 1.34 3a1 1 0 1 0 1.32-1.5A2 2 0 0 1 4 12a2 2 0 0 1 2-2a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.61a1 1 0 0 0 .78.67A3 3 0 0 1 20 11a2.91 2.91 0 0 1-.74 2a1 1 0 0 0 1.48 1.34a5 5 0 0 0-2.32-8.08ZM16 11a1 1 0 0 0-1 1v1a1 1 0 0 0 2 0v-1a1 1 0 0 0-1-1m0 5a1 1 0 0 0-1 1v1a1 1 0 0 0 2 0v-1a1 1 0 0 0-1-1m-4-7a1 1 0 0 0-1 1v1a1 1 0 0 0 2 0v-1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudExclamation(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.71 18.79a1.15 1.15 0 0 0-.33-.21a1 1 0 0 0-1.09.21a1.05 1.05 0 0 0-.29.71a1 1 0 0 0 .29.71a1.15 1.15 0 0 0 .33.21a1 1 0 0 0 1.3-1.3a1 1 0 0 0-.21-.33M12 12.5a1 1 0 0 0-1 1v3a1 1 0 0 0 2 0v-3a1 1 0 0 0-1-1m6.42-4.79a7 7 0 0 0-13.36 1.9A4 4 0 0 0 6 17.5h2a1 1 0 0 0 0-2H6a2 2 0 0 1 0-4a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.61a1 1 0 0 0 .78.66A3 3 0 0 1 17 15.5h-1a1 1 0 0 0 0 2h1a5 5 0 0 0 1.42-9.79"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudHail(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 17a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-4-6a1 1 0 1 0 1 1a1 1 0 0 0-1-1m4 2a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-4 6a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0-4a1 1 0 1 0 1 1a1 1 0 0 0-1-1m8 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-4-6a1 1 0 1 0 1 1a1 1 0 0 0-1-1m4 2a1 1 0 1 0 1 1a1 1 0 0 0-1-1m2.42-4.78A7 7 0 0 0 5.06 8.11A4 4 0 0 0 2 12a4 4 0 0 0 1.34 3a1 1 0 1 0 1.32-1.5A2 2 0 0 1 4 12a2 2 0 0 1 2-2a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.61a1 1 0 0 0 .78.67A3 3 0 0 1 20 11a2.91 2.91 0 0 1-.74 2a1 1 0 0 0 1.48 1.34a5 5 0 0 0-2.32-8.08ZM16 19a1 1 0 1 0 1 1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudHeart(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.42 7.72A7 7 0 0 0 5.06 9.61a4 4 0 0 0-.38 7.66a1.13 1.13 0 0 0 .32.05a1 1 0 0 0 .32-2A2 2 0 0 1 4 13.5a2 2 0 0 1 2-2a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.61a1 1 0 0 0 .78.67a3 3 0 0 1 1 5.53a1 1 0 1 0 1 1.74A5 5 0 0 0 22 12.5a5 5 0 0 0-3.58-4.78M12 12.83a2.94 2.94 0 0 0-3.43.53a2.93 2.93 0 0 0 0 4.13l2.72 2.72a1 1 0 0 0 1.42 0l2.72-2.72a2.93 2.93 0 0 0 0-4.13a2.94 2.94 0 0 0-3.43-.53m2 3.24l-2 2l-2-2a.88.88 0 0 1-.27-.65a.89.89 0 0 1 .27-.65a.92.92 0 0 1 1.3 0a1 1 0 0 0 1.42 0a.94.94 0 0 1 1.3 0a.89.89 0 0 1 .27.65a.88.88 0 0 1-.29.65"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudInfo(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.29 12.29A1 1 0 0 0 12 14h.19a.6.6 0 0 0 .19-.06a.56.56 0 0 0 .17-.09l.15-.12a1 1 0 0 0 0-1.42a1 1 0 0 0-1.41-.02M12 15a1 1 0 0 0-1 1v3a1 1 0 0 0 2 0v-3a1 1 0 0 0-1-1m6.42-6.78a7 7 0 0 0-13.36 1.89A4 4 0 0 0 6 18h2a1 1 0 0 0 0-2H6a2 2 0 0 1 0-4a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.61a1 1 0 0 0 .78.67A3 3 0 0 1 20 13a3 3 0 0 1-3 3h-1a1 1 0 0 0 0 2h1a5 5 0 0 0 1.42-9.78"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudLock(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.42 6.72A7 7 0 0 0 5.06 8.61a4 4 0 0 0-.38 7.66a1.13 1.13 0 0 0 .32.05a1 1 0 0 0 .32-2A2 2 0 0 1 4 12.5a2 2 0 0 1 2-2a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.61a1 1 0 0 0 .78.67a3 3 0 0 1 1 5.53a1 1 0 1 0 1 1.74A5 5 0 0 0 22 11.5a5 5 0 0 0-3.58-4.78m-3.42 9V14.5a3 3 0 0 0-6 0v1.18a3 3 0 0 0 1 5.82h4a3 3 0 0 0 1-5.82Zm-4-1.22a1 1 0 0 1 2 0v1h-2Zm3 5h-4a1 1 0 0 1 0-2h4a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudMeatball(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 15.92h-.77l.39-.67a1 1 0 0 0-1.74-1l-.38.67l-.38-.67a1 1 0 0 0-1.74 1l.39.67H9.5a1 1 0 0 0 0 2h.77l-.39.66a1 1 0 0 0 1.74 1l.38-.66l.38.66a1 1 0 0 0 1.74-1l-.39-.66h.77a1 1 0 0 0 0-2m3.92-7.79A7 7 0 0 0 5.06 10A4 4 0 0 0 2 13.92a4 4 0 0 0 3.34 3.93h.16a1 1 0 0 0 .16-2a2 2 0 0 1-1.66-2a2 2 0 0 1 2-2a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.61a1 1 0 0 0 .78.66a3 3 0 0 1 .62 5.72a1 1 0 1 0 .74 1.85a5 5 0 0 0-.45-9.41Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudMoon(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.7 9.07a1 1 0 0 0-.93-.26a3.13 3.13 0 0 1-.66.08a3 3 0 0 1-3-3a3.13 3.13 0 0 1 .08-.66a1 1 0 0 0-.26-.93A1 1 0 0 0 16 4a4.93 4.93 0 0 0-3.83 4.21A6.24 6.24 0 0 0 10.5 8a6 6 0 0 0-5.94 5.13A3.5 3.5 0 0 0 5.5 20h9.17A4.33 4.33 0 0 0 19 15.67a4.23 4.23 0 0 0-.49-2A4.94 4.94 0 0 0 22 10a1 1 0 0 0-.3-.93m-7 8.93H5.5a1.5 1.5 0 0 1 0-3a1 1 0 0 0 1-1a4 4 0 0 1 6.18-3.34a3.94 3.94 0 0 1 1.57 2a1 1 0 0 0 .78.67a2.33 2.33 0 0 1-.36 4.67Zm2.44-6.11a2.61 2.61 0 0 1-.42 0a4.6 4.6 0 0 0-.72-.31a6.06 6.06 0 0 0-1.82-2.28v-.37a3 3 0 0 1 1.05-2.28a5 5 0 0 0 4.23 4.23a3 3 0 0 1-2.35 1.01Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudMoonHail(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.21 16.29a1 1 0 0 0-.33-.21a1 1 0 0 0-1.09.21a1.15 1.15 0 0 0-.21.33a.84.84 0 0 0-.08.38a1 1 0 0 0 1.38.92a1.15 1.15 0 0 0 .33-.21a1 1 0 0 0 .21-1.09a1.15 1.15 0 0 0-.21-.33m0 4a1.15 1.15 0 0 0-.33-.21a1 1 0 0 0-1.09.21a1.15 1.15 0 0 0-.21.33a1 1 0 1 0 1.84 0a1.15 1.15 0 0 0-.21-.33m3.85-6.12l-.18-.09l-.18-.08a1 1 0 0 0-.58.06a.93.93 0 0 0-.33.21a1 1 0 0 0-.29.71a1 1 0 0 0 .29.71a1.15 1.15 0 0 0 .33.21a.94.94 0 0 0 .76 0a1.15 1.15 0 0 0 .33-.21a1 1 0 0 0 .29-.69a1.05 1.05 0 0 0-.29-.71Zm.15 4.12a1 1 0 0 0-1.09-.21a.9.9 0 0 0-.54.54a1 1 0 0 0 .21 1.09a1.15 1.15 0 0 0 .33.21a.94.94 0 0 0 .76 0a1.15 1.15 0 0 0 .33-.21a1 1 0 0 0 .21-1.09a1.15 1.15 0 0 0-.21-.33M21.7 7.07a1 1 0 0 0-.94-.26a3 3 0 0 1-.65.08a3 3 0 0 1-3-3a3 3 0 0 1 .08-.65A1 1 0 0 0 16 2a4.93 4.93 0 0 0-3.83 4.21A6.24 6.24 0 0 0 10.5 6a6 6 0 0 0-5.94 5.13a3.5 3.5 0 0 0-.46 6.58a1.14 1.14 0 0 0 .4.08a1 1 0 0 0 .4-1.92A1.48 1.48 0 0 1 4 14.5A1.5 1.5 0 0 1 5.5 13a1 1 0 0 0 1-1a4 4 0 0 1 4-4a4.06 4.06 0 0 1 2.19.66a4 4 0 0 1 1.58 2a1 1 0 0 0 .78.67a2.32 2.32 0 0 1 .95 4.28a1 1 0 0 0 1.1 1.68a4.34 4.34 0 0 0 1.9-3.62a4.23 4.23 0 0 0-.49-2A4.94 4.94 0 0 0 22 8a1 1 0 0 0-.3-.93m-4.59 2.82a2.72 2.72 0 0 1-.42 0a4.6 4.6 0 0 0-.69-.35a6.06 6.06 0 0 0-1.82-2.28v-.37a3 3 0 0 1 1.05-2.28a5 5 0 0 0 4.23 4.23a3 3 0 0 1-2.35 1.05"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudMoonMeatball(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.7 8a1 1 0 0 0-.94-.26a3 3 0 0 1-.65.08a3 3 0 0 1-3-3a3.05 3.05 0 0 1 .08-.66a1 1 0 0 0-.26-.94a1 1 0 0 0-.93-.28a5 5 0 0 0-3.83 4.22a5.86 5.86 0 0 0-1.67-.24A6 6 0 0 0 4.56 12A3.52 3.52 0 0 0 2 15.42a3.47 3.47 0 0 0 1.41 2.79a1 1 0 1 0 1.18-1.61A1.46 1.46 0 0 1 4 15.42a1.5 1.5 0 0 1 1.5-1.5a1 1 0 0 0 1-1a4 4 0 0 1 6.17-3.35a3.9 3.9 0 0 1 1.57 2a1 1 0 0 0 .78.66a2.33 2.33 0 0 1 .54 4.44a1 1 0 0 0-.52 1.32a1 1 0 0 0 .92.6a1 1 0 0 0 .4-.09a4.33 4.33 0 0 0 2.6-4a4.23 4.23 0 0 0-.49-2A4.94 4.94 0 0 0 22 8.92a1 1 0 0 0-.3-.92m-4.59 2.82a2.72 2.72 0 0 1-.42 0a4.6 4.6 0 0 0-.72-.31a5.91 5.91 0 0 0-1.82-2.28v-.37a3 3 0 0 1 1.05-2.28a5 5 0 0 0 4.23 4.22a3 3 0 0 1-2.32 1.01Zm-4.61 6.1h-.77l.39-.67a1 1 0 0 0-1.74-1l-.38.67l-.38-.67a1 1 0 0 0-1.74 1l.39.67H7.5a1 1 0 0 0 0 2h.77l-.39.66a1 1 0 0 0 1.74 1l.38-.66l.38.66a1 1 0 1 0 1.74-1l-.39-.66h.77a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudMoonRain(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 14a1 1 0 0 0-1 1v1a1 1 0 0 0 2 0v-1a1 1 0 0 0-1-1m0 5a1 1 0 0 0-1 1v1a1 1 0 0 0 2 0v-1a1 1 0 0 0-1-1M21.7 7.07a1 1 0 0 0-.94-.26a3 3 0 0 1-.65.08a3 3 0 0 1-3-3a3.13 3.13 0 0 1 .08-.66a1 1 0 0 0-.26-.93A1 1 0 0 0 16 2a4.93 4.93 0 0 0-3.83 4.21A6.24 6.24 0 0 0 10.5 6a6 6 0 0 0-5.94 5.13a3.5 3.5 0 0 0-.46 6.58a1.14 1.14 0 0 0 .4.08a1 1 0 0 0 .4-1.92A1.48 1.48 0 0 1 4 14.5A1.5 1.5 0 0 1 5.5 13a1 1 0 0 0 1-1a4 4 0 0 1 4-4a3.91 3.91 0 0 1 2.17.66a3.94 3.94 0 0 1 1.57 2a1 1 0 0 0 .78.67a2.32 2.32 0 0 1 .98 4.28a1 1 0 0 0 1.1 1.68a4.32 4.32 0 0 0 1.9-3.62a4.23 4.23 0 0 0-.49-2A4.94 4.94 0 0 0 22 8a1 1 0 0 0-.3-.93m-4.59 2.82a2.72 2.72 0 0 1-.42 0a4.6 4.6 0 0 0-.69-.35a6.06 6.06 0 0 0-1.82-2.28v-.37a3 3 0 0 1 1.05-2.28a5 5 0 0 0 4.23 4.23a3 3 0 0 1-2.35 1.05M12.5 13a1 1 0 0 0-1 1v1a1 1 0 0 0 2 0v-1a1 1 0 0 0-1-1m0 5a1 1 0 0 0-1 1v1a1 1 0 0 0 2 0v-1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudMoonShowers(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 14a1 1 0 0 0-1 1v6a1 1 0 0 0 2 0v-6a1 1 0 0 0-1-1m4-1a1 1 0 0 0-1 1v6a1 1 0 0 0 2 0v-6a1 1 0 0 0-1-1m9.2-5.93a1 1 0 0 0-.94-.26a3 3 0 0 1-.65.08a3 3 0 0 1-3-3a3.13 3.13 0 0 1 .08-.66a1 1 0 0 0-.26-.93A1 1 0 0 0 16 2a4.93 4.93 0 0 0-3.83 4.21A6.24 6.24 0 0 0 10.5 6a6 6 0 0 0-5.94 5.13a3.5 3.5 0 0 0-.46 6.58a1.14 1.14 0 0 0 .4.08a1 1 0 0 0 .4-1.92A1.48 1.48 0 0 1 4 14.5A1.5 1.5 0 0 1 5.5 13a1 1 0 0 0 1-1a4 4 0 0 1 4-4a3.91 3.91 0 0 1 2.17.66a3.94 3.94 0 0 1 1.57 2a1 1 0 0 0 .78.67a2.32 2.32 0 0 1 .98 4.28a1 1 0 0 0 1.1 1.68a4.32 4.32 0 0 0 1.9-3.62a4.23 4.23 0 0 0-.49-2A4.94 4.94 0 0 0 22 8a1 1 0 0 0-.3-.93m-4.59 2.82a2.72 2.72 0 0 1-.42 0a4.6 4.6 0 0 0-.69-.35a6.06 6.06 0 0 0-1.82-2.28v-.37a3 3 0 0 1 1.05-2.28a5 5 0 0 0 4.23 4.23a3 3 0 0 1-2.35 1.05"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudQuestion(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.42 8.22a7 7 0 0 0-13.36 1.89A4 4 0 0 0 6 18h2a1 1 0 0 0 0-2H6a2 2 0 0 1 0-4a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.61a1 1 0 0 0 .78.67A3 3 0 0 1 20 13a3 3 0 0 1-3 3a1 1 0 0 0 0 2a5 5 0 0 0 1.42-9.78m-5.5 10.4a.56.56 0 0 0-.09-.17l-.12-.16a1.15 1.15 0 0 0-.33-.21a1 1 0 0 0-1.09.21l-.12.16a.56.56 0 0 0-.09.17a.64.64 0 0 0-.06.18a1.5 1.5 0 0 0 0 .2a1.23 1.23 0 0 0 0 .19a.6.6 0 0 0 .06.19a.56.56 0 0 0 .09.17l.12.16A1.05 1.05 0 0 0 12 20a1 1 0 0 0 .38-.08a1.15 1.15 0 0 0 .33-.21l.12-.16a.56.56 0 0 0 .09-.17a.6.6 0 0 0 .06-.19A1.23 1.23 0 0 0 13 19a1.5 1.5 0 0 0 0-.2a.64.64 0 0 0-.08-.18M12 11a3 3 0 0 0-2.6 1.5a1 1 0 0 0 .37 1.37a1 1 0 0 0 1.36-.37A1 1 0 0 1 12 13a1 1 0 0 1 0 2a1 1 0 0 0 0 2a3 3 0 0 0 0-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudRain(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.56 12.67a1 1 0 0 0-1.12 0c-.11.08-2.69 1.86-2.69 4.58a3.25 3.25 0 0 0 6.5 0c0-2.75-2.58-4.51-2.69-4.58M12 18.5a1.25 1.25 0 0 1-1.25-1.25A3.66 3.66 0 0 1 12 14.8a3.61 3.61 0 0 1 1.25 2.45A1.25 1.25 0 0 1 12 18.5m6.42-10.78A7 7 0 0 0 5.06 9.61a4 4 0 0 0 .61 7.87h.08a1 1 0 0 0 1-.92a1 1 0 0 0-.92-1.08A2 2 0 0 1 4 13.5a2 2 0 0 1 2-2a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.61a1 1 0 0 0 .78.67a3 3 0 0 1 .43 5.79a1 1 0 0 0 .62 1.9a5 5 0 0 0-.14-9.53"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudRainSun(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 7h-.8a4.25 4.25 0 0 0-.52-1.27l.56-.56a1 1 0 0 0-1.41-1.41l-.56.56A4.25 4.25 0 0 0 17 3.8V3a1 1 0 0 0-2 0v.8a4.1 4.1 0 0 0-1.26.52l-.57-.56a1 1 0 0 0-1.41 1.41l.56.57c-.09.15-.16.32-.24.48A5.85 5.85 0 0 0 10.5 6a6 6 0 0 0-5.94 5.13a3.5 3.5 0 0 0-.46 6.58a1.14 1.14 0 0 0 .4.08a1 1 0 0 0 .4-1.92A1.48 1.48 0 0 1 4 14.5A1.5 1.5 0 0 1 5.5 13a1 1 0 0 0 1-1a4 4 0 0 1 7.78-1.29a1 1 0 0 0 .78.67a2.32 2.32 0 0 1 .94 4.23a1 1 0 0 0 1.1 1.68a4.34 4.34 0 0 0 1.9-3.62a4.19 4.19 0 0 0-.3-1.55l.13.12a1 1 0 0 0 .7.29a1 1 0 0 0 .71-.29a1 1 0 0 0 0-1.41l-.56-.56A4.25 4.25 0 0 0 20.2 9h.8a1 1 0 0 0 0-2m-3.34 2.65a2.09 2.09 0 0 1-.6.42A4.17 4.17 0 0 0 16 9.54a6.12 6.12 0 0 0-2.09-2.49a2.42 2.42 0 0 1 .46-.7a2.43 2.43 0 0 1 3.3 0a2.37 2.37 0 0 1 0 3.3ZM8.5 14a1 1 0 0 0-1 1v6a1 1 0 0 0 2 0v-6a1 1 0 0 0-1-1m4-1a1 1 0 0 0-1 1v6a1 1 0 0 0 2 0v-6a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudRedo(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 12.5a1 1 0 0 0-.91.6a4 4 0 1 0 .55 6.4a1 1 0 1 0-1.32-1.5a2 2 0 0 1-1.32.5a2 2 0 1 1 1-3.75h-.22a1 1 0 0 0 0 2H15a1 1 0 0 0 1-1V13.5a1 1 0 0 0-1-1m3.42-4.78A7 7 0 0 0 5.06 9.61a4 4 0 0 0-.38 7.66a1.13 1.13 0 0 0 .32.05a1 1 0 0 0 .32-2A2 2 0 0 1 4 13.5a2 2 0 0 1 2-2a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.61a1 1 0 0 0 .78.67a3 3 0 0 1 1 5.53a1 1 0 1 0 1 1.74A5 5 0 0 0 22 12.5a5 5 0 0 0-3.58-4.78"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudShare(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 15a2 2 0 1 0-2-2l-1.9.87a2 2 0 0 0-1.1-.33a2 2 0 0 0 0 4a1.88 1.88 0 0 0 .92-.24l2.1 1a2 2 0 1 0 .8-1.84l-1.75-.8l1.91-.88a2 2 0 0 0 1.02.22m3.92-7.78A7 7 0 0 0 5.06 9.11a4 4 0 0 0-.38 7.66a1.13 1.13 0 0 0 .32.05a1 1 0 0 0 .32-2A2 2 0 0 1 4 13a2 2 0 0 1 2-2a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.61a1 1 0 0 0 .78.67a3 3 0 0 1 1 5.53a1 1 0 1 0 1 1.74A5 5 0 0 0 22 12a5 5 0 0 0-3.58-4.78"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudShield(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.42 7.22A7 7 0 0 0 5.06 9.11a4 4 0 0 0-.38 7.66a1.13 1.13 0 0 0 .32.05a1 1 0 0 0 .32-2A2 2 0 0 1 4 13a2 2 0 0 1 2-2a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.61a1 1 0 0 0 .78.67a3 3 0 0 1 1 5.53a1 1 0 1 0 1 1.74A5 5 0 0 0 22 12a5 5 0 0 0-3.58-4.78m-3.63 4.44a2.73 2.73 0 0 1-2.2-.47a1 1 0 0 0-1.18 0a2.72 2.72 0 0 1-2.2.47a1 1 0 0 0-.84.2a1 1 0 0 0-.37.77V16a4.63 4.63 0 0 0 1.84 3.7l1.57 1.15a1 1 0 0 0 1.18 0l1.57-1.16A4.6 4.6 0 0 0 16 16v-3.37a1 1 0 0 0-.37-.77a1 1 0 0 0-.84-.2M14 16a2.62 2.62 0 0 1-1 2l-1 .72l-1-.72a2.62 2.62 0 0 1-1-2v-2.28a4.68 4.68 0 0 0 2-.55a4.68 4.68 0 0 0 2 .55Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudShowers(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 11a1 1 0 0 0-1 1v3a1 1 0 0 0 2 0v-3a1 1 0 0 0-1-1m4 6a1 1 0 0 0-1 1v3a1 1 0 0 0 2 0v-3a1 1 0 0 0-1-1m-4 0a1 1 0 0 0-1 1v3a1 1 0 0 0 2 0v-3a1 1 0 0 0-1-1M18.42 6.22A7 7 0 0 0 5.06 8.11A4 4 0 0 0 2 12a4 4 0 0 0 1.34 3a1 1 0 1 0 1.32-1.5A2 2 0 0 1 4 12a2 2 0 0 1 2-2a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.61a1 1 0 0 0 .78.67A3 3 0 0 1 20 11a2.91 2.91 0 0 1-.74 2a1 1 0 0 0 1.48 1.34a5 5 0 0 0-2.32-8.08ZM12 11a1 1 0 0 0-1 1v3a1 1 0 0 0 2 0v-3a1 1 0 0 0-1-1m4 0a1 1 0 0 0-1 1v3a1 1 0 0 0 2 0v-3a1 1 0 0 0-1-1m0 6a1 1 0 0 0-1 1v3a1 1 0 0 0 2 0v-3a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudShowersAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 13a1 1 0 0 0-1 1v5a1 1 0 0 0 2 0v-5a1 1 0 0 0-1-1m4-2a1 1 0 0 0-1 1v9a1 1 0 0 0 2 0v-9a1 1 0 0 0-1-1m6.42-4.78A7 7 0 0 0 5.06 8.11A4 4 0 0 0 2 12a4 4 0 0 0 1.34 3a1 1 0 1 0 1.32-1.5A2 2 0 0 1 4 12a2 2 0 0 1 2-2a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.61a1 1 0 0 0 .78.67A3 3 0 0 1 20 11a2.91 2.91 0 0 1-.74 2a1 1 0 0 0 1.48 1.34a5 5 0 0 0-2.32-8.08ZM16 13a1 1 0 0 0-1 1v5a1 1 0 0 0 2 0v-5a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudShowersHeavy(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.89 18.06a1 1 0 0 0-1.28.6l-.73 2a1 1 0 0 0 .6 1.28a1 1 0 0 0 .34.06a1 1 0 0 0 .94-.66l.73-2a1 1 0 0 0-.6-1.28m-4 0a1 1 0 0 0-1.28.6l-.73 2a1 1 0 0 0 .6 1.28a1 1 0 0 0 .34.06a1 1 0 0 0 .94-.66l.73-2a1 1 0 0 0-.6-1.28m0-7a1 1 0 0 0-1.28.6l-1.1 3A1 1 0 0 0 7.45 16a1 1 0 0 0 .94-.66l1.1-3a1 1 0 0 0-.6-1.28m4 0a1 1 0 0 0-1.28.6l-1.1 3a1 1 0 0 0 .94 1.34a1 1 0 0 0 .94-.66l1.1-3a1 1 0 0 0-.6-1.28m5.53-4.84A7 7 0 0 0 5.06 8.11A4 4 0 0 0 2 12a4 4 0 0 0 1.34 3a1 1 0 1 0 1.32-1.5A2 2 0 0 1 4 12a2 2 0 0 1 2-2a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.61a1 1 0 0 0 .78.67A3 3 0 0 1 20 11a2.91 2.91 0 0 1-.74 2a1 1 0 0 0 1.48 1.34a5 5 0 0 0-2.32-8.08Zm-1.53 11.84a1 1 0 0 0-1.28.6l-.73 2a1 1 0 0 0 .6 1.28a1 1 0 0 0 .34.06a1 1 0 0 0 .94-.66l.73-2a1 1 0 0 0-.6-1.28m0-7a1 1 0 0 0-1.28.6l-1.1 3a1 1 0 0 0 .94 1.34a1 1 0 0 0 .94-.66l1.1-3a1 1 0 0 0-.6-1.28"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16.71 16.29l-13-13a1 1 0 0 0-1.42 1.42l3.36 3.35a7 7 0 0 0-.59 2A4 4 0 0 0 6 18h9.59l2.7 2.71a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM6 16a2 2 0 0 1 0-4a1 1 0 0 0 1-1a5 5 0 0 1 .2-1.39L13.59 16Zm12.42-7.78A7 7 0 0 0 12 4a6.74 6.74 0 0 0-2.32.4a1 1 0 0 0 .66 1.88A4.91 4.91 0 0 1 12 6a5 5 0 0 1 4.73 3.39a1 1 0 0 0 .78.67a3 3 0 0 1 1.85 4.79a1 1 0 0 0 .16 1.4a1 1 0 0 0 .62.22a1 1 0 0 0 .78-.38a5 5 0 0 0-2.5-7.87"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudSun(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 9h-.8a4.25 4.25 0 0 0-.52-1.27l.56-.56a1 1 0 0 0-1.41-1.41l-.56.56A4.25 4.25 0 0 0 17 5.8V5a1 1 0 0 0-2 0v.8a4.1 4.1 0 0 0-1.26.52l-.57-.56a1 1 0 0 0-1.41 1.41l.56.57c-.09.15-.16.32-.24.48A5.85 5.85 0 0 0 10.5 8a6 6 0 0 0-5.94 5.13A3.5 3.5 0 0 0 5.5 20h9.17A4.33 4.33 0 0 0 19 15.67a4.19 4.19 0 0 0-.3-1.55l.13.12a1 1 0 0 0 .7.29a1 1 0 0 0 .71-.29a1 1 0 0 0 0-1.41l-.56-.56A4.25 4.25 0 0 0 20.2 11h.8a1 1 0 0 0 0-2m-6.33 9H5.5a1.5 1.5 0 0 1 0-3a1 1 0 0 0 1-1a4 4 0 0 1 7.78-1.29a1 1 0 0 0 .78.67a2.33 2.33 0 0 1-.39 4.62m3-6.35a2.17 2.17 0 0 1-.6.4a4.49 4.49 0 0 0-1.07-.51a6.12 6.12 0 0 0-2.09-2.49a2.25 2.25 0 0 1 .46-.69a2.42 2.42 0 0 1 3.29 0a2.37 2.37 0 0 1 0 3.3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudSunHail(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.21 20.29a1 1 0 0 0-1.09-.21a.93.93 0 0 0-.33.21a1.15 1.15 0 0 0-.21.33a.94.94 0 0 0 0 .76a.9.9 0 0 0 .54.54a.94.94 0 0 0 .76 0a1.15 1.15 0 0 0 .33-.21A1 1 0 0 0 9.5 21a.84.84 0 0 0-.08-.38a1.15 1.15 0 0 0-.21-.33m4-2a.93.93 0 0 0-.33-.21a1 1 0 0 0-1.09.21a1.15 1.15 0 0 0-.21.33a1 1 0 0 0 1.3 1.3a1.15 1.15 0 0 0 .33-.21a1 1 0 0 0 .21-1.09a1.15 1.15 0 0 0-.21-.33m-4-2a1 1 0 0 0-1.09-.21a1 1 0 0 0-.33.21a1.15 1.15 0 0 0-.21.33a1 1 0 0 0 .21 1.09a1.15 1.15 0 0 0 .33.21a.94.94 0 0 0 .76 0a1.15 1.15 0 0 0 .33-.21A1 1 0 0 0 9.5 17a.84.84 0 0 0-.08-.38a1.15 1.15 0 0 0-.21-.33m2.91-2.21a1 1 0 0 0-.33.21a1.05 1.05 0 0 0-.29.71a1 1 0 0 0 .29.71a1.15 1.15 0 0 0 .33.21a.84.84 0 0 0 .38.08a1 1 0 0 0 1-1a1.05 1.05 0 0 0-.29-.71a1 1 0 0 0-1.09-.21M21 7h-.8a4.25 4.25 0 0 0-.52-1.27l.56-.56a1 1 0 0 0-1.41-1.41l-.56.56A4.25 4.25 0 0 0 17 3.8V3a1 1 0 0 0-2 0v.8a4.1 4.1 0 0 0-1.26.52l-.57-.56a1 1 0 0 0-1.41 1.41l.56.57c-.09.15-.16.32-.24.48A5.85 5.85 0 0 0 10.5 6a6 6 0 0 0-5.94 5.13a3.5 3.5 0 0 0-.46 6.58a1.14 1.14 0 0 0 .4.08a1 1 0 0 0 .4-1.92A1.48 1.48 0 0 1 4 14.5A1.5 1.5 0 0 1 5.5 13a1 1 0 0 0 1-1a4 4 0 0 1 7.78-1.29a1 1 0 0 0 .78.67a2.32 2.32 0 0 1 .94 4.23a1 1 0 0 0 1.1 1.68a4.34 4.34 0 0 0 1.9-3.62a4.19 4.19 0 0 0-.3-1.55l.13.12a1 1 0 0 0 .7.29a1 1 0 0 0 .71-.29a1 1 0 0 0 0-1.41l-.56-.56A4.25 4.25 0 0 0 20.2 9h.8a1 1 0 0 0 0-2m-3.34 2.65a2.09 2.09 0 0 1-.6.42A4.17 4.17 0 0 0 16 9.54a6.12 6.12 0 0 0-2.09-2.49a2.42 2.42 0 0 1 .46-.7a2.43 2.43 0 0 1 3.3 0a2.37 2.37 0 0 1 0 3.3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudSunMeatball(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 16.92h-.77l.39-.67a1 1 0 0 0-1.74-1l-.38.67l-.38-.67a1 1 0 0 0-1.74 1l.39.67H7.5a1 1 0 0 0 0 2h.77l-.39.66a1 1 0 0 0 1.74 1l.38-.66l.38.66a1 1 0 1 0 1.74-1l-.39-.66h.77a1 1 0 0 0 0-2m8.5-9h-.8a4.25 4.25 0 0 0-.52-1.27l.56-.56a1 1 0 0 0-1.41-1.41l-.56.56A4.25 4.25 0 0 0 17 4.72v-.8a1 1 0 0 0-2 0v.8a4.1 4.1 0 0 0-1.26.52l-.57-.56a1 1 0 0 0-1.41 1.41l.56.56c-.09.16-.16.33-.24.49a5.85 5.85 0 0 0-1.58-.22A6 6 0 0 0 4.56 12A3.52 3.52 0 0 0 2 15.42a3.47 3.47 0 0 0 1.41 2.79a1 1 0 1 0 1.18-1.61A1.46 1.46 0 0 1 4 15.42a1.5 1.5 0 0 1 1.5-1.5a1 1 0 0 0 1-1a4 4 0 0 1 7.78-1.29a1 1 0 0 0 .78.66a2.33 2.33 0 0 1 .54 4.44a1 1 0 0 0-.52 1.32a1 1 0 0 0 .92.6a.93.93 0 0 0 .4-.09a4.33 4.33 0 0 0 2.6-4a4.29 4.29 0 0 0-.3-1.56l.13.13a1 1 0 0 0 .7.29a1 1 0 0 0 .71-.29a1 1 0 0 0 0-1.42l-.56-.56a4.2 4.2 0 0 0 .52-1.26h.8a1 1 0 0 0 0-2Zm-3.34 2.64a1.89 1.89 0 0 1-.6.41a4.15 4.15 0 0 0-1.06-.51A6 6 0 0 0 13.88 8a2.18 2.18 0 0 1 .46-.7a2.42 2.42 0 0 1 3.3 0a2.34 2.34 0 0 1 0 3.29Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudSunRain(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 15a1 1 0 0 0-1 1v1a1 1 0 0 0 2 0v-1a1 1 0 0 0-1-1M22 7h-1.1a5.22 5.22 0 0 0-.73-1.76l.83-.77a1 1 0 1 0-1.42-1.42l-.77.78A5 5 0 0 0 17 3.1V2a1 1 0 0 0-2 0v1.1a5.22 5.22 0 0 0-1.76.73l-.77-.78a1 1 0 0 0-1.42 1.42l.78.77a5.06 5.06 0 0 0-.77 2A5.76 5.76 0 0 0 9.5 7a6 6 0 0 0-5.94 5.13a3.5 3.5 0 0 0-.46 6.58a1.14 1.14 0 0 0 .4.08a1 1 0 0 0 .4-1.92A1.48 1.48 0 0 1 3 15.5A1.5 1.5 0 0 1 4.5 14a1 1 0 0 0 1-1a4 4 0 0 1 7.78-1.29a1 1 0 0 0 .78.67a2.32 2.32 0 0 1 .94 4.23a1 1 0 0 0 1.1 1.68a4.34 4.34 0 0 0 1.9-3.62a4.41 4.41 0 0 0-.45-1.92a5.17 5.17 0 0 0 1.21-.58l.77.78A1 1 0 0 0 21 13a1 1 0 0 0 0-1.42l-.78-.77A5 5 0 0 0 20.9 9H22a1 1 0 0 0 0-2m-3.87 3.12A3 3 0 0 1 16 11h-.06a4.12 4.12 0 0 0-1-.46a5.93 5.93 0 0 0-2-2.38V8a3 3 0 0 1 .87-2.12a3.1 3.1 0 0 1 4.25 0a3 3 0 0 1 0 4.25ZM7.5 20a1 1 0 0 0-1 1v1a1 1 0 0 0 2 0v-1a1 1 0 0 0-1-1m4-1a1 1 0 0 0-1 1v1a1 1 0 0 0 2 0v-1a1 1 0 0 0-1-1m0-5a1 1 0 0 0-1 1v1a1 1 0 0 0 2 0v-1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudSunRainAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 19a1 1 0 0 0-1 1v1a1 1 0 0 0 2 0v-1a1 1 0 0 0-1-1m0-5a1 1 0 0 0-1 1v1a1 1 0 0 0 2 0v-1a1 1 0 0 0-1-1M21 7h-.8a4.25 4.25 0 0 0-.52-1.27l.56-.56a1 1 0 0 0-1.41-1.41l-.56.56A4.25 4.25 0 0 0 17 3.8V3a1 1 0 0 0-2 0v.8a4.1 4.1 0 0 0-1.26.52l-.57-.56a1 1 0 0 0-1.41 1.41l.56.57c-.09.15-.16.32-.24.48A5.85 5.85 0 0 0 10.5 6a6 6 0 0 0-5.94 5.13a3.5 3.5 0 0 0-.46 6.58a1.14 1.14 0 0 0 .4.08a1 1 0 0 0 .4-1.92A1.48 1.48 0 0 1 4 14.5A1.5 1.5 0 0 1 5.5 13a1 1 0 0 0 1-1a4 4 0 0 1 7.78-1.29a1 1 0 0 0 .78.67a2.32 2.32 0 0 1 .94 4.23a1 1 0 0 0 1.1 1.68a4.34 4.34 0 0 0 1.9-3.62a4.19 4.19 0 0 0-.3-1.55l.13.12a1 1 0 0 0 .7.29a1 1 0 0 0 .71-.29a1 1 0 0 0 0-1.41l-.56-.56A4.25 4.25 0 0 0 20.2 9h.8a1 1 0 0 0 0-2m-3.34 2.65a2.09 2.09 0 0 1-.6.42A4.17 4.17 0 0 0 16 9.54a6.12 6.12 0 0 0-2.09-2.49a2.42 2.42 0 0 1 .46-.7a2.43 2.43 0 0 1 3.3 0a2.37 2.37 0 0 1 0 3.3ZM12.5 18a1 1 0 0 0-1 1v1a1 1 0 0 0 2 0v-1a1 1 0 0 0-1-1m0-5a1 1 0 0 0-1 1v1a1 1 0 0 0 2 0v-1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudSunTear(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.05 14.33a1 1 0 0 0-1.11 0c-.1.08-2.44 1.67-2.44 4.17a3 3 0 0 0 6 0c0-2.5-2.34-4.1-2.45-4.17m-.55 5.17a1 1 0 0 1-1-1a3 3 0 0 1 1-2a3 3 0 0 1 1 2a1 1 0 0 1-1 1M21 7.5h-.8a4.25 4.25 0 0 0-.52-1.27l.56-.56a1 1 0 0 0-1.41-1.41l-.56.56A4.25 4.25 0 0 0 17 4.3v-.8a1 1 0 0 0-2 0v.8a4.1 4.1 0 0 0-1.26.52l-.57-.56a1 1 0 0 0-1.41 1.41l.56.57c-.09.15-.16.32-.24.48a5.85 5.85 0 0 0-1.58-.22a6 6 0 0 0-5.94 5.13a3.5 3.5 0 0 0-.46 6.58a1.14 1.14 0 0 0 .4.08a1 1 0 0 0 .4-1.92A1.48 1.48 0 0 1 4 15a1.5 1.5 0 0 1 1.5-1.5a1 1 0 0 0 1-1a4 4 0 0 1 7.78-1.29a1 1 0 0 0 .78.67a2.32 2.32 0 0 1 .94 4.23a1 1 0 0 0 1.1 1.68a4.3 4.3 0 0 0 1.65-5.18l.13.13a1 1 0 0 0 .7.29a1 1 0 0 0 .71-.29a1 1 0 0 0 0-1.41l-.56-.56a4.25 4.25 0 0 0 .47-1.27h.8a1 1 0 0 0 0-2m-3.34 2.65a2.45 2.45 0 0 1-.6.41A4.17 4.17 0 0 0 16 10a6.12 6.12 0 0 0-2.09-2.49a2.25 2.25 0 0 1 .46-.69a2.42 2.42 0 0 1 3.29 0a2.37 2.37 0 0 1 0 3.3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudTimes(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.42 8.22a7 7 0 0 0-13.36 1.89A4 4 0 0 0 6 18a1 1 0 0 0 0-2a2 2 0 0 1 0-4a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.61a1 1 0 0 0 .78.67a3 3 0 0 1 .24 5.84a1 1 0 0 0 .5 1.94a5 5 0 0 0 .17-9.62m-3.71 6.07a1 1 0 0 0-1.42 0L12 15.59l-1.29-1.3a1 1 0 0 0-1.42 1.42l1.3 1.29l-1.3 1.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l1.29-1.3l1.29 1.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L13.41 17l1.3-1.29a1 1 0 0 0 0-1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudUnlock(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 15.5h-3v-1a1 1 0 0 1 1.88-.5a1 1 0 0 0 1.37.34a1 1 0 0 0 .34-1.34a3.08 3.08 0 0 0-.46-.59A3 3 0 0 0 12 11.5a3 3 0 0 0-3 3v1.18a3 3 0 0 0 1 5.82h4a3 3 0 0 0 0-6m0 4h-4a1 1 0 0 1 0-2h4a1 1 0 0 1 0 2m4.42-12.78A7 7 0 0 0 5.06 8.61a4 4 0 0 0-.38 7.66a1.13 1.13 0 0 0 .32.05a1 1 0 0 0 .32-2A2 2 0 0 1 4 12.5a2 2 0 0 1 2-2a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.61a1 1 0 0 0 .78.67a3 3 0 0 1 1 5.53a1 1 0 1 0 1 1.74A5 5 0 0 0 22 11.5a5 5 0 0 0-3.58-4.78"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudUpload(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.42 8.22a7 7 0 0 0-13.36 1.89A4 4 0 0 0 6 18a1 1 0 0 0 0-2a2 2 0 0 1 0-4a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.61a1 1 0 0 0 .78.67a3 3 0 0 1 .24 5.84a1 1 0 0 0 .5 1.94a5 5 0 0 0 .17-9.62m-5.71 2.07a1 1 0 0 0-.33-.21a1 1 0 0 0-.76 0a1 1 0 0 0-.33.21l-3 3a1 1 0 0 0 1.42 1.42l1.29-1.3V19a1 1 0 0 0 2 0v-5.59l1.29 1.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudWifi(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 5.94a7.1 7.1 0 0 1 7 0a1 1 0 0 0 1.37-.37a1 1 0 0 0-.37-1.36a9.14 9.14 0 0 0-9 0a1 1 0 0 0-.37 1.36a1 1 0 0 0 1.37.37m9.92 5.27a5.91 5.91 0 0 0-.36-.71a1 1 0 0 0-1.38-.33a1 1 0 0 0-.33 1.37a4.58 4.58 0 0 1 .38.84a1 1 0 0 0 .78.67A3 3 0 0 1 20 16a3 3 0 0 1-3 3H5.66A2 2 0 0 1 4 17.4A2 2 0 0 1 6 15a1 1 0 0 0 1-1a4.92 4.92 0 0 1 .67-2.49a1 1 0 0 0-.37-1.37a1 1 0 0 0-1.36.37a6.75 6.75 0 0 0-.88 2.6a4 4 0 0 0-2.13 1.33A4 4 0 0 0 5.46 21H17a5 5 0 0 0 1.42-9.79M14.87 9v-.06a.92.92 0 0 0 .13-.2a1 1 0 0 0-.57-1.29a6.36 6.36 0 0 0-1.74-.38h-.3a5.47 5.47 0 0 0-.81 0a3 3 0 0 0-.31 0a6.36 6.36 0 0 0-1.74.38A1 1 0 0 0 9 8.74a1.22 1.22 0 0 0 .12.19a.61.61 0 0 0 0 .11a1 1 0 0 0 1.37.36a3.08 3.08 0 0 1 3 0a1 1 0 0 0 1.38-.4M11 12a1 1 0 1 0 1-1a1 1 0 0 0-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudWind(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 18a1 1 0 1 0 1 1a1 1 0 0 0-1-1m14-4h-7a1 1 0 0 0 0 2h7a1 1 0 0 1 0 2a1 1 0 0 0 0 2a3 3 0 0 0 0-6m-5-3a1 1 0 0 0 1 1h4a3 3 0 0 0 3-3a1 1 0 0 0-2 0a1 1 0 0 1-1 1h-4a1 1 0 0 0-1 1m-4 4a1 1 0 0 0-1-1H6a2 2 0 0 1 0-4a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.61a1 1 0 1 0 1.9-.64A7 7 0 0 0 5.06 8.11A4 4 0 0 0 6 16h3a1 1 0 0 0 1-1m0-4a1 1 0 1 0 1-1a1 1 0 0 0-1 1m4 7H9a1 1 0 0 0 0 2h5a1 1 0 0 1 1 1a1 1 0 0 0 2 0a3 3 0 0 0-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clouds(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.47 10.54A6 6 0 0 0 14 7a5.82 5.82 0 0 0-1.39.18a5 5 0 0 0-9 2A3 3 0 0 0 4.5 15h1a4 4 0 0 0 0 .5A3.5 3.5 0 0 0 9 19h9.17a4.33 4.33 0 0 0 1.3-8.46M4.5 13a1 1 0 0 1 0-2a1 1 0 0 0 1-1a3 3 0 0 1 3-3a3 3 0 0 1 2.22 1a6 6 0 0 0-2.66 4.13a3.49 3.49 0 0 0-1.5.87Zm13.67 4H9a1.5 1.5 0 0 1 0-3a1 1 0 0 0 1-1a4 4 0 0 1 7.78-1.29a1 1 0 0 0 .78.67a2.33 2.33 0 0 1-.39 4.62"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Club(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 12.5a5.52 5.52 0 0 0-3.51-5.12a5.49 5.49 0 0 0-11 0A5.5 5.5 0 0 0 8.5 18h.2a6.91 6.91 0 0 1-1.24 2.39A1 1 0 0 0 8.24 22h7.52a1 1 0 0 0 .78-1.63A6.91 6.91 0 0 1 15.3 18h.2a5.51 5.51 0 0 0 5.5-5.5M10.06 20a8.89 8.89 0 0 0 .81-2.56a5.47 5.47 0 0 0 1.13-.7a5.47 5.47 0 0 0 1.13.7a8.89 8.89 0 0 0 .81 2.56Zm2.72-5.3l-.08-.08l-.08-.08h-.07l-.18-.09l-.18-.06h-.38l-.18.06a.56.56 0 0 0-.17.09h-.08l-.08.08l-.08.08a3.5 3.5 0 1 1-3.47-5.62A1.11 1.11 0 0 0 7.91 9a1.42 1.42 0 0 0 .18-.08a.83.83 0 0 0 .14-.13a.62.62 0 0 0 .21-.31a.61.61 0 0 0 .07-.17a.69.69 0 0 0 0-.2a1 1 0 0 0 0-.17a3.4 3.4 0 0 1 0-.45a3.5 3.5 0 0 1 7 0a3.4 3.4 0 0 1 0 .45a.81.81 0 0 0 0 .16a.74.74 0 0 0 0 .21a.61.61 0 0 0 .07.17a.62.62 0 0 0 .21.31a.83.83 0 0 0 .14.13a1.42 1.42 0 0 0 .18.08a1.11 1.11 0 0 0 .16.07a3.5 3.5 0 1 1-3.47 5.62Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeBranch(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 6.06a3 3 0 0 0-1.15 5.77A2 2 0 0 1 14 13.06h-4a3.91 3.91 0 0 0-2 .56V7.88a3 3 0 1 0-2 0v8.36a3 3 0 1 0 2.16.05A2 2 0 0 1 10 15.06h4a4 4 0 0 0 3.91-3.16A3 3 0 0 0 17 6.06m-10-2a1 1 0 1 1-1 1a1 1 0 0 1 1-1m0 16a1 1 0 1 1 1-1a1 1 0 0 1-1 1m10-10a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Coffee(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 17h4a5 5 0 0 0 5-5v-1h1a3 3 0 0 0 0-6h-1V4a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v8a5 5 0 0 0 5 5m9-10h1a1 1 0 0 1 0 2h-1ZM6 5h10v7a3 3 0 0 1-3 3H9a3 3 0 0 1-3-3Zm15 14H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cog(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.32 9.55l-1.89-.63l.89-1.78A1 1 0 0 0 20.13 6L18 3.87a1 1 0 0 0-1.15-.19l-1.78.89l-.63-1.89A1 1 0 0 0 13.5 2h-3a1 1 0 0 0-.95.68l-.63 1.89l-1.78-.89A1 1 0 0 0 6 3.87L3.87 6a1 1 0 0 0-.19 1.15l.89 1.78l-1.89.63a1 1 0 0 0-.68.94v3a1 1 0 0 0 .68.95l1.89.63l-.89 1.78A1 1 0 0 0 3.87 18L6 20.13a1 1 0 0 0 1.15.19l1.78-.89l.63 1.89a1 1 0 0 0 .95.68h3a1 1 0 0 0 .95-.68l.63-1.89l1.78.89a1 1 0 0 0 1.13-.19L20.13 18a1 1 0 0 0 .19-1.15l-.89-1.78l1.89-.63a1 1 0 0 0 .68-.94v-3a1 1 0 0 0-.68-.95M20 12.78l-1.2.4A2 2 0 0 0 17.64 16l.57 1.14l-1.1 1.1l-1.11-.6a2 2 0 0 0-2.79 1.16l-.4 1.2h-1.59l-.4-1.2A2 2 0 0 0 8 17.64l-1.14.57l-1.1-1.1l.6-1.11a2 2 0 0 0-1.16-2.82l-1.2-.4v-1.56l1.2-.4A2 2 0 0 0 6.36 8l-.57-1.11l1.1-1.1L8 6.36a2 2 0 0 0 2.82-1.16l.4-1.2h1.56l.4 1.2A2 2 0 0 0 16 6.36l1.14-.57l1.1 1.1l-.6 1.11a2 2 0 0 0 1.16 2.79l1.2.4ZM12 8a4 4 0 1 0 4 4a4 4 0 0 0-4-4m0 6a2 2 0 1 1 2-2a2 2 0 0 1-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Coins(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.22 12a3 3 0 0 0 .78-2a3 3 0 0 0-3-3h-5.18A3 3 0 0 0 11 3H5a3 3 0 0 0-3 3a3 3 0 0 0 .78 2a3 3 0 0 0 0 4a3 3 0 0 0 0 4A3 3 0 0 0 2 18a3 3 0 0 0 3 3h14a3 3 0 0 0 2.22-5a3 3 0 0 0 0-4M11 19H5a1 1 0 0 1 0-2h6a1 1 0 0 1 0 2m0-4H5a1 1 0 0 1 0-2h6a1 1 0 0 1 0 2m0-4H5a1 1 0 0 1 0-2h6a1 1 0 0 1 0 2m0-4H5a1 1 0 0 1 0-2h6a1 1 0 0 1 0 2m8.69 11.71A.93.93 0 0 1 19 19h-5.18a2.87 2.87 0 0 0 0-2H19a1 1 0 0 1 1 1a1 1 0 0 1-.31.71m0-4A.93.93 0 0 1 19 15h-5.18a2.87 2.87 0 0 0 0-2H19a1 1 0 0 1 1 1a1 1 0 0 1-.31.71m0-4A.93.93 0 0 1 19 11h-5.18a2.87 2.87 0 0 0 0-2H19a1 1 0 0 1 1 1a1 1 0 0 1-.31.71"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Columns(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1M11 20H4V4h7Zm9 0h-7V4h7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Comment(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2A10 10 0 0 0 2 12a9.89 9.89 0 0 0 2.26 6.33l-2 2a1 1 0 0 0-.21 1.09A1 1 0 0 0 3 22h9a10 10 0 0 0 0-20m0 18H5.41l.93-.93a1 1 0 0 0 0-1.41A8 8 0 1 1 12 20"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentAdd(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 8a1 1 0 0 0 2 0V7h1a1 1 0 0 0 0-2H7V4a1 1 0 0 0-2 0v1H4a1 1 0 0 0 0 2h1Zm13-3h-6a1 1 0 0 0 0 2h6a1 1 0 0 1 1 1v9.72l-1.57-1.45a1 1 0 0 0-.68-.27H8a1 1 0 0 1-1-1v-3a1 1 0 0 0-2 0v3a3 3 0 0 0 3 3h8.36l3 2.73A1 1 0 0 0 20 21a1.1 1.1 0 0 0 .4-.08A1 1 0 0 0 21 20V8a3 3 0 0 0-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 2H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h11.59l3.7 3.71A1 1 0 0 0 21 22a.84.84 0 0 0 .38-.08A1 1 0 0 0 22 21V5a3 3 0 0 0-3-3m1 16.59l-2.29-2.3A1 1 0 0 0 17 16H5a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentAltBlock(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 6h-4a1 1 0 0 0 0 2h4a1 1 0 0 1 1 1v9.72l-1.57-1.45a1 1 0 0 0-.68-.27H9a1 1 0 0 1-1-1v-1a1 1 0 0 0-2 0v1a3 3 0 0 0 3 3h8.36l3 2.73A1 1 0 0 0 21 22a1.1 1.1 0 0 0 .4-.08A1 1 0 0 0 22 21V9a3 3 0 0 0-3-3m-8.46 4.54A5 5 0 1 0 7 12a5 5 0 0 0 3.54-1.46M4 7a3 3 0 0 1 3-3a3 3 0 0 1 1.28.3l-4 4A3 3 0 0 1 4 7m5.7-1.29A3 3 0 0 1 10 7a3 3 0 0 1-4.28 2.7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentAltChartLines(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 8a1 1 0 0 0-1 1v4a1 1 0 0 0 2 0V9a1 1 0 0 0-1-1m-4 3a1 1 0 0 0-1 1v1a1 1 0 0 0 2 0v-1a1 1 0 0 0-1-1m11-9H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h11.59l3.7 3.71A1 1 0 0 0 21 22a.84.84 0 0 0 .38-.08A1 1 0 0 0 22 21V5a3 3 0 0 0-3-3m1 16.59l-2.29-2.3A1 1 0 0 0 17 16H5a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1ZM16 6a1 1 0 0 0-1 1v6a1 1 0 0 0 2 0V7a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentAltCheck(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m6.77 9.15l5.44-5.44a1 1 0 1 0-1.42-1.42L6.06 7L4.21 5.17a1 1 0 0 0-1.42 1.42l2.56 2.56a1 1 0 0 0 1.42 0M18.5 6H13a1 1 0 0 0 0 2h5.5a1 1 0 0 1 1 1v9.72l-1.57-1.45a1 1 0 0 0-.68-.27H8.5a1 1 0 0 1-1-1v-3.5a1 1 0 0 0-2 0V16a3 3 0 0 0 3 3h8.36l3 2.73a1 1 0 0 0 .68.27a1.1 1.1 0 0 0 .4-.08a1 1 0 0 0 .6-.92V9a3 3 0 0 0-3.04-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentAltDots(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 9a1 1 0 1 0 1 1a1 1 0 0 0-1-1m7-7H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h11.59l3.7 3.71A1 1 0 0 0 21 22a.84.84 0 0 0 .38-.08A1 1 0 0 0 22 21V5a3 3 0 0 0-3-3m1 16.59l-2.29-2.3A1 1 0 0 0 17 16H5a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1ZM8 9a1 1 0 1 0 1 1a1 1 0 0 0-1-1m8 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentAltDownload(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.12 11.92a.94.94 0 0 0 .76 0a1 1 0 0 0 .33-.21l3-3a1 1 0 1 0-1.42-1.42L7.5 8.59V3a1 1 0 0 0-2 0v5.59l-1.29-1.3a1 1 0 1 0-1.42 1.42l3 3a1 1 0 0 0 .33.21M18.5 6h-5a1 1 0 0 0 0 2h5a1 1 0 0 1 1 1v9.72l-1.57-1.45a1 1 0 0 0-.68-.27H8.5a1 1 0 0 1-1-1v-1a1 1 0 0 0-2 0v1a3 3 0 0 0 3 3h8.36l3 2.73a1 1 0 0 0 .68.27a1.1 1.1 0 0 0 .4-.08a1 1 0 0 0 .6-.92V9a3 3 0 0 0-3.04-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentAltEdit(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.5 5.5h-4a1 1 0 0 0 0 2h4a1 1 0 0 1 1 1v9.72l-1.57-1.45a1 1 0 0 0-.68-.27H8.5a1 1 0 0 1-1-1v-1a1 1 0 0 0-2 0v1a3 3 0 0 0 3 3h8.36l3 2.73a1 1 0 0 0 .68.27a1.1 1.1 0 0 0 .4-.08a1 1 0 0 0 .6-.92v-12a3 3 0 0 0-3.04-3m-9.42 7h2.42a1 1 0 0 0 1-1V9.08a1 1 0 0 0-.29-.71L6.63 2.79a1 1 0 0 0-1.41 0L2.79 5.22a1 1 0 0 0 0 1.41l5.58 5.58a1 1 0 0 0 .71.29M5.92 4.91l4.58 4.58v1h-1L4.91 5.92Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentAltExclamation(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 2H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h11.59l3.7 3.71A1 1 0 0 0 21 22a.84.84 0 0 0 .38-.08A1 1 0 0 0 22 21V5a3 3 0 0 0-3-3m1 16.59l-2.29-2.3A1 1 0 0 0 17 16H5a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1ZM12 12a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0-6a1 1 0 0 0-1 1v3a1 1 0 0 0 2 0V7a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentAltHeart(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 6.35a3.07 3.07 0 0 0-3.54.53a3 3 0 0 0 0 4.24L11.29 14a1 1 0 0 0 1.42 0l2.83-2.83a3 3 0 0 0 0-4.24A3.07 3.07 0 0 0 12 6.35m2.12 3.36L12 11.83L9.88 9.71a1 1 0 0 1 0-1.42a1 1 0 0 1 1.41 0a1 1 0 0 0 1.42 0a1 1 0 0 1 1.41 0a1 1 0 0 1 0 1.42M19 2H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h11.59l3.7 3.71A1 1 0 0 0 21 22a.84.84 0 0 0 .38-.08A1 1 0 0 0 22 21V5a3 3 0 0 0-3-3m1 16.59l-2.29-2.3A1 1 0 0 0 17 16H5a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentAltImage(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 2H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h11.59l3.7 3.71A1 1 0 0 0 21 22a.84.84 0 0 0 .38-.08A1 1 0 0 0 22 21V5a3 3 0 0 0-3-3M5.77 16H5a1 1 0 0 1-1-1v-.42l3.3-3.29a1 1 0 0 1 1.41 0l.87.87ZM20 18.59l-2.29-2.3A1 1 0 0 0 17 16H8.59l6.71-6.71a1 1 0 0 1 1.4 0l3.3 3.29Zm0-8.83l-1.88-1.87a3 3 0 0 0-4.24 0L11 10.76l-.88-.87a3.06 3.06 0 0 0-4.24 0L4 11.76V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentAltInfo(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 11a1 1 0 0 0 1-1V6a1 1 0 0 0-2 0v4a1 1 0 0 0 1 1m-.71-7.29a1 1 0 0 0 1.09.21a1.15 1.15 0 0 0 .33-.21a1.15 1.15 0 0 0 .21-.33A.84.84 0 0 0 6 3a1 1 0 0 0-.29-.71a1 1 0 0 0-1.09-.21a1 1 0 0 0-.33.21A1 1 0 0 0 4 3a1 1 0 0 0 .08.38a1.15 1.15 0 0 0 .21.33M17 6H9a1 1 0 0 0 0 2h8a1 1 0 0 1 1 1v9.72l-1.57-1.45a1 1 0 0 0-.68-.27H7a1 1 0 0 1-1-1v-2a1 1 0 0 0-2 0v2a3 3 0 0 0 3 3h8.36l3 2.73A1 1 0 0 0 19 22a1.1 1.1 0 0 0 .4-.08A1 1 0 0 0 20 21V9a3 3 0 0 0-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentAltLines(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 11H7a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2m4-4H7a1 1 0 0 0 0 2h10a1 1 0 0 0 0-2m2-5H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h11.59l3.7 3.71A1 1 0 0 0 21 22a.84.84 0 0 0 .38-.08A1 1 0 0 0 22 21V5a3 3 0 0 0-3-3m1 16.59l-2.29-2.3A1 1 0 0 0 17 16H5a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentAltLock(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 10V8a3 3 0 0 0-2-2.82V4a3 3 0 0 0-6 0v1.18A3 3 0 0 0 2 8v2a3 3 0 0 0 3 3h4a3 3 0 0 0 3-3M6 4a1 1 0 0 1 2 0v1H6Zm-2 6V8a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1m15-3h-4a1 1 0 0 0 0 2h4a1 1 0 0 1 1 1v9.72l-1.57-1.45a1 1 0 0 0-.68-.27H9a1 1 0 0 1-1-1v-1a1 1 0 0 0-2 0v1a3 3 0 0 0 3 3h8.36l3 2.73A1 1 0 0 0 21 23a1.1 1.1 0 0 0 .4-.08A1 1 0 0 0 22 22V10a3 3 0 0 0-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentAltMedical(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 9h-1V8a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2h1v1a1 1 0 0 0 2 0v-1h1a1 1 0 0 0 0-2m5-7H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h11.59l3.7 3.71A1 1 0 0 0 21 22a.84.84 0 0 0 .38-.08A1 1 0 0 0 22 21V5a3 3 0 0 0-3-3m1 16.59l-2.29-2.3A1 1 0 0 0 17 16H5a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentAltMessage(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 7H7a1 1 0 0 0 0 2h10a1 1 0 0 0 0-2m0 4H7a1 1 0 0 0 0 2h10a1 1 0 0 0 0-2m2-9H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h11.59l3.7 3.71A1 1 0 0 0 21 22a.84.84 0 0 0 .38-.08A1 1 0 0 0 22 21V5a3 3 0 0 0-3-3m1 16.59l-2.29-2.3A1 1 0 0 0 17 16H5a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentAltNotes(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 7a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0 4a1 1 0 1 0 1 1a1 1 0 0 0-1-1m10 0h-6a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2m0-4h-6a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2m2-5H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h11.59l3.7 3.71A1 1 0 0 0 21 22a.84.84 0 0 0 .38-.08A1 1 0 0 0 22 21V5a3 3 0 0 0-3-3m1 16.59l-2.29-2.3A1 1 0 0 0 17 16H5a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentAltPlus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 9h-2V7a1 1 0 0 0-2 0v2H9a1 1 0 0 0 0 2h2v2a1 1 0 0 0 2 0v-2h2a1 1 0 0 0 0-2m4-7H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h11.59l3.7 3.71A1 1 0 0 0 21 22a.84.84 0 0 0 .38-.08A1 1 0 0 0 22 21V5a3 3 0 0 0-3-3m1 16.59l-2.29-2.3A1 1 0 0 0 17 16H5a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentAltQuestion(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.29 10.3a1 1 0 0 0 1.09 1.63a1.19 1.19 0 0 0 .33-.22a1 1 0 0 0 .21-.32A.85.85 0 0 0 8 11a1 1 0 0 0-.29-.7a1 1 0 0 0-1.42 0M7 5a1 1 0 0 1 0 2a1 1 0 0 0 0 2a3 3 0 1 0-2.6-4.5a1 1 0 0 0 .37 1.37a1 1 0 0 0 1.36-.37A1 1 0 0 1 7 5m12 1h-6a1 1 0 0 0 0 2h6a1 1 0 0 1 1 1v9.72l-1.57-1.45a1 1 0 0 0-.68-.27H9a1 1 0 0 1-1-1v-1a1 1 0 0 0-2 0v1a3 3 0 0 0 3 3h8.36l3 2.73A1 1 0 0 0 21 22a1.1 1.1 0 0 0 .4-.08A1 1 0 0 0 22 21V9a3 3 0 0 0-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentAltRedo(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.3 10.75A1 1 0 1 0 9 9.25A3 3 0 1 1 7 4a3 3 0 0 1 2.23 1H8a1 1 0 0 0 0 2h3a1 1 0 0 0 1-1V3a1 1 0 0 0-2 0a5 5 0 1 0 .3 7.75M19 6h-4a1 1 0 0 0 0 2h4a1 1 0 0 1 1 1v9.72l-1.57-1.45a1 1 0 0 0-.68-.27H9a1 1 0 0 1-1-1v-1a1 1 0 0 0-2 0v1a3 3 0 0 0 3 3h8.36l3 2.73A1 1 0 0 0 21 22a1.1 1.1 0 0 0 .4-.08A1 1 0 0 0 22 21V9a3 3 0 0 0-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentAltSearch(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 6h-5a1 1 0 0 0 0 2h5a1 1 0 0 1 1 1v9.72l-1.57-1.45a1 1 0 0 0-.68-.27H9a1 1 0 0 1-1-1v-2a1 1 0 0 0-2 0v2a3 3 0 0 0 3 3h8.36l3 2.73A1 1 0 0 0 21 22a1.1 1.1 0 0 0 .4-.08A1 1 0 0 0 22 21V9a3 3 0 0 0-3-3M8.57 10l1.72 1.73a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L10 8.57a4.37 4.37 0 0 0 .65-2.26a4.32 4.32 0 1 0-8.65 0a4.32 4.32 0 0 0 4.31 4.32A4.35 4.35 0 0 0 8.57 10M4 6.31a2.29 2.29 0 0 1 .68-1.63a2.32 2.32 0 0 1 3.32 0A2.31 2.31 0 0 1 8 8a2.32 2.32 0 0 1-4-1.69"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentAltShare(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.5 6.5h-5a1 1 0 0 0 0 2h5a1 1 0 0 1 1 1v9.72l-1.57-1.45a1 1 0 0 0-.68-.27H9.5a1 1 0 0 1-1-1v-2a1 1 0 0 0-2 0v2a3 3 0 0 0 3 3h8.36l3 2.73a1 1 0 0 0 .68.27a1.1 1.1 0 0 0 .4-.08a1 1 0 0 0 .6-.92v-12a3 3 0 0 0-3.04-3m-10 5a2 2 0 1 0-1.18-3.61l-1.75-.8l1.91-.88a2 2 0 0 0 1 .29a2 2 0 1 0-2-2l-1.9.87A2 2 0 0 0 4.5 5a2 2 0 0 0 0 4a2 2 0 0 0 .93-.24l2.09 1A2 2 0 0 0 9.5 11.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentAltShield(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.57 12a1 1 0 0 0 .58-.19l1.62-1.16A4.56 4.56 0 0 0 10.68 7V3.63a1 1 0 0 0-.37-.77a1 1 0 0 0-.84-.2a3 3 0 0 1-2.33-.48a1 1 0 0 0-1.14 0a3 3 0 0 1-2.33.48a1 1 0 0 0-1.2 1V7a4.55 4.55 0 0 0 1.9 3.7L6 11.81a.94.94 0 0 0 .57.19M4.46 7V4.72a5.16 5.16 0 0 0 2.11-.55a5.12 5.12 0 0 0 2.11.55V7a2.57 2.57 0 0 1-1.07 2l-1 .74L5.53 9a2.57 2.57 0 0 1-1.07-2m14.08-1h-5a1 1 0 0 0 0 2h5a1 1 0 0 1 1 1v9.72L18 17.27a.94.94 0 0 0-.68-.27H8.54a1 1 0 0 1-1-1v-1a1 1 0 0 0-2 0v1a3 3 0 0 0 3 3h8.36l3 2.73a1 1 0 0 0 .67.27a1.15 1.15 0 0 0 .41-.08a1 1 0 0 0 .59-.92V9a3 3 0 0 0-3.03-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentAltSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.66 6H17a1 1 0 0 1 1 1v6.34a1 1 0 1 0 2 0V7a3 3 0 0 0-3-3h-6.34a1 1 0 0 0 0 2m11.05 14.29l-18-18a1 1 0 0 0-1.42 1.42l2 2A3 3 0 0 0 4 7v12a1 1 0 0 0 .62.92A.84.84 0 0 0 5 20a1 1 0 0 0 .71-.29L8.41 17h7.18l4.7 4.71a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42M8 15a1 1 0 0 0-.71.29L6 16.59V7.41L13.59 15Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentAltUpload(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.5 11a1 1 0 0 0 2 0V5.41l1.29 1.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-3-3a1 1 0 0 0-.33-.21a1 1 0 0 0-.76 0a1 1 0 0 0-.33.21l-3 3a1 1 0 0 0 1.42 1.42l1.29-1.3Zm13-5h-5a1 1 0 0 0 0 2h5a1 1 0 0 1 1 1v9.72l-1.57-1.45a1 1 0 0 0-.68-.27H8.5a1 1 0 0 1-1-1v-1a1 1 0 0 0-2 0v1a3 3 0 0 0 3 3h8.36l3 2.73a1 1 0 0 0 .68.27a1.1 1.1 0 0 0 .4-.08a1 1 0 0 0 .6-.92V9a3 3 0 0 0-3.04-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentAltVerify(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.71 6.29l-5.3 5.3l-2.12-2.12a1 1 0 1 0-1.41 1.41l2.83 2.83a1 1 0 0 0 .7.29a1 1 0 0 0 .71-.29l6-6a1 1 0 0 0 0-1.42a1 1 0 0 0-1.41 0M19 2H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h11.59l3.7 3.71A1 1 0 0 0 21 22a.84.84 0 0 0 .38-.08A1 1 0 0 0 22 21V5a3 3 0 0 0-3-3m1 16.59l-2.29-2.3A1 1 0 0 0 17 16H5a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentBlock(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 14a1 1 0 0 0-1.22.72A7 7 0 0 1 11 20H5.41l.64-.63a1 1 0 0 0 0-1.41a7 7 0 0 1 3.2-11.74a1 1 0 0 0-.5-1.94A9 9 0 0 0 4 18.62l-1.71 1.67a1 1 0 0 0-.21 1.09A1 1 0 0 0 3 22h8a9 9 0 0 0 8.72-6.75A1 1 0 0 0 19 14m1.54-10.54A5 5 0 1 0 22 7a5 5 0 0 0-1.46-3.54M14 7a3 3 0 0 1 3-3a3 3 0 0 1 1.29.3l-4 4A3 3 0 0 1 14 7m5.12 2.12a3.08 3.08 0 0 1-3.4.57l4-4A3 3 0 0 1 20 7a3 3 0 0 1-.88 2.12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentChartLine(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 10a1 1 0 0 0-1 1v4a1 1 0 0 0 2 0v-4a1 1 0 0 0-1-1m-4 3a1 1 0 0 0-1 1v1a1 1 0 0 0 2 0v-1a1 1 0 0 0-1-1m4-11A10 10 0 0 0 2 12a9.89 9.89 0 0 0 2.26 6.33l-2 2a1 1 0 0 0-.21 1.09A1 1 0 0 0 3 22h9a10 10 0 0 0 0-20m0 18H5.41l.93-.93a1 1 0 0 0 .3-.71a1 1 0 0 0-.3-.7A8 8 0 1 1 12 20m4-12a1 1 0 0 0-1 1v6a1 1 0 0 0 2 0V9a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentCheck(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.29 7.13a1 1 0 0 0 0 1.42l1.92 1.92a1 1 0 0 0 1.42 0l4.08-4.08A1 1 0 1 0 19.29 5l-3.37 3.35l-1.21-1.22a1 1 0 0 0-1.42 0m6.62 3.51a1 1 0 0 0-.91 1.08a2.62 2.62 0 0 1 0 .28a7 7 0 0 1-7 7H6.41l.64-.63a1 1 0 0 0 0-1.41a7 7 0 0 1 3.48-11.81a7.14 7.14 0 0 1 2.8 0a1 1 0 1 0 .4-2a9.15 9.15 0 0 0-3.61 0A9.05 9.05 0 0 0 3 12a9 9 0 0 0 2 5.62l-1.71 1.67a1 1 0 0 0-.21 1.09A1 1 0 0 0 4 21h8a9 9 0 0 0 9-9v-.44a1 1 0 0 0-1.09-.92"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentDots(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 11a1 1 0 1 0 1 1a1 1 0 0 0-1-1m4 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1m4 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-4-9A10 10 0 0 0 2 12a9.89 9.89 0 0 0 2.26 6.33l-2 2a1 1 0 0 0-.21 1.09A1 1 0 0 0 3 22h9a10 10 0 0 0 0-20m0 18H5.41l.93-.93a1 1 0 0 0 .3-.71a1 1 0 0 0-.3-.7A8 8 0 1 1 12 20"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentDownload(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 14a1 1 0 0 0-1.22.72A7 7 0 0 1 11 20H5.41l.64-.63a1 1 0 0 0 0-1.41A7 7 0 0 1 11 6a8.49 8.49 0 0 1 .88 0a1 1 0 1 0 .24-2A8.32 8.32 0 0 0 11 4a9 9 0 0 0-7 14.62l-1.71 1.67a1 1 0 0 0-.21 1.09A1 1 0 0 0 3 22h8a9 9 0 0 0 8.72-6.75A1 1 0 0 0 19 14m2.71-6.74a1 1 0 0 0-1.42 0L19 8.59V3a1 1 0 0 0-2 0v5.59l-1.29-1.3a1 1 0 1 0-1.42 1.42l3 3a1 1 0 0 0 .33.21a.94.94 0 0 0 .76 0a1 1 0 0 0 .33-.21l3-3a1 1 0 0 0 0-1.42Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentEdit(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.71 4.72l-2.43-2.43a1 1 0 0 0-1.41 0l-5.58 5.58a1 1 0 0 0-.29.71V11a1 1 0 0 0 1 1h2.42a1 1 0 0 0 .71-.29l5.58-5.58a1 1 0 0 0 0-1.41M15 10h-1V9l4.58-4.58l1 1Zm4 2a1 1 0 0 0-1 1a7 7 0 0 1-7 7H5.41l.64-.63a1 1 0 0 0 0-1.42A7 7 0 0 1 11 6a1 1 0 0 0 0-2a9 9 0 0 0-7 14.62l-1.71 1.67a1 1 0 0 0-.21 1.09A1 1 0 0 0 3 22h8a9 9 0 0 0 9-9a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentExclamation(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 14a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0-12A10 10 0 0 0 2 12a9.89 9.89 0 0 0 2.26 6.33l-2 2a1 1 0 0 0-.21 1.09A1 1 0 0 0 3 22h9a10 10 0 0 0 0-20m0 18H5.41l.93-.93a1 1 0 0 0 .3-.71a1 1 0 0 0-.3-.7A8 8 0 1 1 12 20m0-12a1 1 0 0 0-1 1v3a1 1 0 0 0 2 0V9a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentHeart(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 8.35a3.07 3.07 0 0 0-3.54.53a3 3 0 0 0 0 4.24L11.29 16a1 1 0 0 0 1.42 0l2.83-2.83a3 3 0 0 0 0-4.24A3.07 3.07 0 0 0 12 8.35m2.12 3.36L12 13.83l-2.12-2.12a1 1 0 0 1 0-1.42a1 1 0 0 1 1.41 0a1 1 0 0 0 1.42 0a1 1 0 0 1 1.41 0a1 1 0 0 1 0 1.42M12 2A10 10 0 0 0 2 12a9.89 9.89 0 0 0 2.26 6.33l-2 2a1 1 0 0 0-.21 1.09A1 1 0 0 0 3 22h9a10 10 0 0 0 0-20m0 18H5.41l.93-.93a1 1 0 0 0 0-1.41A8 8 0 1 1 12 20"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentImage(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a10 10 0 0 0-9.56 12.91a.29.29 0 0 0 0 .1a9.83 9.83 0 0 0 1.79 3.32l-2 2a1 1 0 0 0-.21 1.09A1 1 0 0 0 3 22h9a10 10 0 0 0 9.46-6.78v-.08A9.89 9.89 0 0 0 22 12A10 10 0 0 0 12 2m0 18H5.41l.3-.29l8.41-8.42a1 1 0 0 1 1.4 0l3.62 3.6l.23.22A8 8 0 0 1 12 20m-7.46-5.13l1.58-1.58a1 1 0 0 1 1.41 0l.87.87l-2.72 2.74a7.67 7.67 0 0 1-1.14-2.03m15.41-2l-3-3a3 3 0 0 0-4.24 0l-2.89 2.89l-.88-.87a3 3 0 0 0-4.23 0l-.71.67A5.25 5.25 0 0 1 4 12a8 8 0 0 1 16 0a8.27 8.27 0 0 1 0 .86Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentInfo(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 11a1 1 0 0 0-1 1v3a1 1 0 0 0 2 0v-3a1 1 0 0 0-1-1m0-3a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0-6A10 10 0 0 0 2 12a9.89 9.89 0 0 0 2.26 6.33l-2 2a1 1 0 0 0-.21 1.09A1 1 0 0 0 3 22h9a10 10 0 0 0 0-20m0 18H5.41l.93-.93a1 1 0 0 0 .3-.71a1 1 0 0 0-.3-.7A8 8 0 1 1 12 20"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentInfoAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 9a1 1 0 0 0-1 1v3a1 1 0 0 0 2 0v-3a1 1 0 0 0-1-1m7-7H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h11.59l3.7 3.71A1 1 0 0 0 21 22a.84.84 0 0 0 .38-.08A1 1 0 0 0 22 21V5a3 3 0 0 0-3-3m1 16.59l-2.29-2.3A1 1 0 0 0 17 16H5a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1ZM12 6a1 1 0 1 0 1 1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentLines(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 9H7a1 1 0 0 0 0 2h10a1 1 0 0 0 0-2m-4 4H7a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2M12 2A10 10 0 0 0 2 12a9.89 9.89 0 0 0 2.26 6.33l-2 2a1 1 0 0 0-.21 1.09A1 1 0 0 0 3 22h9a10 10 0 0 0 0-20m0 18H5.41l.93-.93a1 1 0 0 0 0-1.41A8 8 0 1 1 12 20"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentLock(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.92 16.13a1 1 0 0 0-1.37.37A7 7 0 0 1 11.5 20H5.91l.64-.63a1 1 0 0 0 0-1.41a7 7 0 0 1 4.07-11.91a1 1 0 1 0-.24-2a9 9 0 0 0-5.91 14.57l-1.68 1.67a1 1 0 0 0-.21 1.09a1 1 0 0 0 .92.62h8a9 9 0 0 0 7.79-4.5a1 1 0 0 0-.37-1.37m.58-9.95V5a3 3 0 0 0-6 0v1.18a3 3 0 0 0-2 2.82v2a3 3 0 0 0 3 3h4a3 3 0 0 0 3-3V9a3 3 0 0 0-2-2.82M15.5 5a1 1 0 0 1 2 0v1h-2Zm4 6a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentMedical(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2A10 10 0 0 0 2 12a9.89 9.89 0 0 0 2.26 6.33l-2 2a1 1 0 0 0-.21 1.09A1 1 0 0 0 3 22h9a10 10 0 0 0 0-20m0 18H5.41l.93-.93a1 1 0 0 0 0-1.41A8 8 0 1 1 12 20m2-9h-1v-1a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2h1v1a1 1 0 0 0 2 0v-1h1a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentMessage(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2A10 10 0 0 0 2 12a9.89 9.89 0 0 0 2.26 6.33l-2 2a1 1 0 0 0-.21 1.09A1 1 0 0 0 3 22h9a10 10 0 0 0 0-20m0 18H5.41l.93-.93a1 1 0 0 0 0-1.41A8 8 0 1 1 12 20m5-9H7a1 1 0 0 0 0 2h10a1 1 0 0 0 0-2m-2 4H9a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2M9 9h6a1 1 0 0 0 0-2H9a1 1 0 0 0 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentNotes(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 11h6a1 1 0 0 0 0-2h-6a1 1 0 0 0 0 2M7 9a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0 4a1 1 0 1 0 1 1a1 1 0 0 0-1-1m5-11A10 10 0 0 0 2 12a9.89 9.89 0 0 0 2.26 6.33l-2 2a1 1 0 0 0-.21 1.09A1 1 0 0 0 3 22h9a10 10 0 0 0 0-20m0 18H5.41l.93-.93a1 1 0 0 0 .3-.71a1 1 0 0 0-.3-.7A8 8 0 1 1 12 20m5-7h-6a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentPlus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2A10 10 0 0 0 2 12a9.89 9.89 0 0 0 2.26 6.33l-2 2a1 1 0 0 0-.21 1.09A1 1 0 0 0 3 22h9a10 10 0 0 0 0-20m0 18H5.41l.93-.93a1 1 0 0 0 0-1.41A8 8 0 1 1 12 20m3-9h-2V9a1 1 0 0 0-2 0v2H9a1 1 0 0 0 0 2h2v2a1 1 0 0 0 2 0v-2h2a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentQuestion(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.77 5.87a1 1 0 0 0 1.36-.37A1 1 0 0 1 18 6a1 1 0 0 1-1 1a1 1 0 0 0 0 2a3 3 0 1 0-2.6-4.5a1 1 0 0 0 .37 1.37m4.3 7.13a1 1 0 0 0-1.12.86A7 7 0 0 1 11 20H5.41l.65-.65a1 1 0 0 0 0-1.41A7 7 0 0 1 11 6a1 1 0 0 0 0-2a9 9 0 0 0-7 14.61l-1.71 1.68a1 1 0 0 0-.21 1.09A1 1 0 0 0 3 22h8a9 9 0 0 0 8.93-7.87a1 1 0 0 0-.86-1.13m-1.69-2.93a1 1 0 0 0-.58-.07l-.18.06l-.18.09l-.15.13a1 1 0 0 0-.21.32a.84.84 0 0 0-.08.4a1 1 0 0 0 .07.39a1 1 0 0 0 .22.32A1 1 0 0 0 17 12a1 1 0 0 0 1-1a.84.84 0 0 0-.08-.38a1.07 1.07 0 0 0-.54-.54Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentRedo(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 14a1 1 0 0 0-1.22.72A7 7 0 0 1 11 20H5.41l.64-.63a1 1 0 0 0 0-1.41a7 7 0 0 1 3.2-11.74a1 1 0 0 0-.5-1.94A9 9 0 0 0 4 18.62l-1.71 1.67a1 1 0 0 0-.21 1.09A1 1 0 0 0 3 22h8a9 9 0 0 0 8.72-6.75A1 1 0 0 0 19 14m2-12a1 1 0 0 0-1 1a5 5 0 1 0 .3 7.75a1 1 0 1 0-1.3-1.5A3 3 0 1 1 17 4a3 3 0 0 1 2.23 1H18a1 1 0 0 0 0 2h3a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentSearch(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.21 10.29l-1.73-1.72a4.37 4.37 0 0 0 .65-2.26a4.31 4.31 0 1 0-4.32 4.32a4.35 4.35 0 0 0 2.26-.63l1.72 1.73a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.44M17.45 8a2.32 2.32 0 0 1-3.95-1.69a2.29 2.29 0 0 1 .68-1.63a2.32 2.32 0 0 1 3.27 0a2.31 2.31 0 0 1 0 3.27Zm2.05 6a1 1 0 0 0-1.22.72A7 7 0 0 1 11.5 20H5.91l.64-.63a1 1 0 0 0 0-1.41a7 7 0 0 1-2-5a7 7 0 0 1 4.32-6.44a1 1 0 1 0-.74-1.86a9 9 0 0 0-3.66 14l-1.68 1.63a1 1 0 0 0-.21 1.09a1 1 0 0 0 .92.62h8a9 9 0 0 0 8.72-6.75A1 1 0 0 0 19.5 14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentShare(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.06 13.51a1 1 0 0 0-1.11.87A7 7 0 0 1 11 20.5H5.41l.64-.63a1 1 0 0 0 0-1.41a7 7 0 0 1 3.2-11.74a1 1 0 0 0-.5-1.94A9 9 0 0 0 4 19.12l-1.71 1.67a1 1 0 0 0-.21 1.09a1 1 0 0 0 .92.62h8a9 9 0 0 0 8.93-7.88a1 1 0 0 0-.87-1.11M19 7.5a2 2 0 0 0-1.18.39l-1.75-.8L18 6.21a2 2 0 0 0 1 .29a2 2 0 1 0-2-2l-1.9.87A2 2 0 0 0 14 5a2 2 0 0 0 0 4a2 2 0 0 0 .93-.24l2.09 1A2 2 0 1 0 19 7.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentShield(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.7 2.86a1 1 0 0 0-.84-.2a3 3 0 0 1-2.33-.48a1 1 0 0 0-1.15 0a3 3 0 0 1-2.33.48a1 1 0 0 0-.84.2a1 1 0 0 0-.37.77V7a4.56 4.56 0 0 0 1.91 3.7l1.62 1.16a1 1 0 0 0 1.17 0l1.62-1.16A4.56 4.56 0 0 0 22.07 7V3.63a1 1 0 0 0-.37-.77M20.07 7A2.57 2.57 0 0 1 19 9l-1 .74L16.91 9a2.57 2.57 0 0 1-1.07-2V4.72A5.17 5.17 0 0 0 18 4.17a5.12 5.12 0 0 0 2.11.55Zm-1.14 7a1 1 0 0 0-1.21.72A7 7 0 0 1 10.93 20H5.35l.65-.63A1 1 0 0 0 6 18a7 7 0 0 1 4.93-12a1 1 0 0 0 0-2a9 9 0 0 0-7 14.62l-1.7 1.67a1 1 0 0 0 .7 1.71h8a9 9 0 0 0 8.72-6.75a1 1 0 0 0-.72-1.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.19 5.23A7.12 7.12 0 0 1 12 5a7 7 0 0 1 7 7a7.12 7.12 0 0 1-.23 1.81a1 1 0 0 0 .7 1.23a1.15 1.15 0 0 0 .26 0a1 1 0 0 0 1-.74A8.91 8.91 0 0 0 21 12a9 9 0 0 0-9-9a8.91 8.91 0 0 0-2.33.3A1 1 0 0 0 9 4.53a1 1 0 0 0 1.19.7m11.52 15.06l-18-18a1 1 0 0 0-1.42 1.42L5 6.38a9 9 0 0 0 0 11.24l-1.71 1.67a1 1 0 0 0-.21 1.09A1 1 0 0 0 4 21h8a9 9 0 0 0 5.62-2l2.67 2.68a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.39M12 19H6.41l.64-.63a1 1 0 0 0 0-1.41a7 7 0 0 1-.65-9.15l9.79 9.79A7 7 0 0 1 12 19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentUpload(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 14a1 1 0 0 0-1.22.72A7 7 0 0 1 11 20H5.41l.64-.63a1 1 0 0 0 0-1.41A7 7 0 0 1 11 6a1 1 0 0 0 0-2a9 9 0 0 0-7 14.62l-1.71 1.67a1 1 0 0 0-.21 1.09A1 1 0 0 0 3 22h8a9 9 0 0 0 8.72-6.75A1 1 0 0 0 19 14m2.71-8.74l-3-3a1 1 0 0 0-.33-.21a1 1 0 0 0-.76 0a1 1 0 0 0-.33.21l-3 3a1 1 0 0 0 1.42 1.42L17 5.41V11a1 1 0 0 0 2 0V5.41l1.29 1.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentVerify(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.29 8.85l-4.73 4.74l-1.85-1.86a1 1 0 0 0-1.42 1.42l2.56 2.56a1 1 0 0 0 1.42 0l5.44-5.44a1 1 0 1 0-1.42-1.42M12 2A10 10 0 0 0 2 12a9.89 9.89 0 0 0 2.26 6.33l-2 2a1 1 0 0 0-.21 1.09A1 1 0 0 0 3 22h9a10 10 0 0 0 0-20m0 18H5.41l.93-.93a1 1 0 0 0 0-1.41A8 8 0 1 1 12 20"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Comments(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.61 19.19a7 7 0 0 0-2.74-10.57a8 8 0 1 0-14.19 6.29l-1.39 1.38a1 1 0 0 0-.21 1.09A1 1 0 0 0 3 18h5.69A7 7 0 0 0 15 22h6a1 1 0 0 0 .92-.62a1 1 0 0 0-.21-1.09ZM8 15a6.63 6.63 0 0 0 .08 1H5.41l.35-.34a1 1 0 0 0 0-1.42A5.93 5.93 0 0 1 4 10a6 6 0 0 1 6-6a5.94 5.94 0 0 1 5.65 4H15a7 7 0 0 0-7 7m10.54 5l.05.05H15a5 5 0 1 1 3.54-1.46a1 1 0 0 0-.3.7a1 1 0 0 0 .3.71"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CommentsAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 8h-1V5a3 3 0 0 0-3-3H5a3 3 0 0 0-3 3v12a1 1 0 0 0 .62.92A.84.84 0 0 0 3 18a1 1 0 0 0 .71-.29l2.81-2.82H8v1.44a3 3 0 0 0 3 3h6.92l2.37 2.38A1 1 0 0 0 21 22a.84.84 0 0 0 .38-.08A1 1 0 0 0 22 21V11a3 3 0 0 0-3-3M8 11v1.89H6.11a1 1 0 0 0-.71.29L4 14.59V5a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v3h-5a3 3 0 0 0-3 3m12 7.59l-1-1a1 1 0 0 0-.71-.3H11a1 1 0 0 1-1-1V11a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompactDisc(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8m0-11a3 3 0 1 0 3 3a3 3 0 0 0-3-3m0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Comparison(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 20H4v-.54l5-5l3.8 3.8a1 1 0 0 0 1.41 0l7.5-7.5a1 1 0 0 0 0-1.42a1 1 0 0 0-1.41 0l-6.8 6.8l-3.79-3.8a1 1 0 0 0-1.41 0L4 16.63v-5.17l5-5l2.8 2.8a1 1 0 0 0 1.41 0L18 4.47l2.19 2.19a1 1 0 0 0 1.41-1.42l-2.91-2.89a1 1 0 0 0-1.41 0l-4.8 4.8l-2.79-2.8a1 1 0 0 0-1.41 0L4 8.63V3a1 1 0 0 0-2 0v18a1 1 0 0 0 1 1h18a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compass(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m1 17.93V19a1 1 0 0 0-2 0v.93A8 8 0 0 1 4.07 13H5a1 1 0 0 0 0-2h-.93A8 8 0 0 1 11 4.07V5a1 1 0 0 0 2 0v-.93A8 8 0 0 1 19.93 11H19a1 1 0 0 0 0 2h.93A8 8 0 0 1 13 19.93m2.14-12.38l-5 2.12a1 1 0 0 0-.52.52l-2.12 5a1 1 0 0 0 .21 1.1a1 1 0 0 0 .7.3a.93.93 0 0 0 .4-.09l5-2.12a1 1 0 0 0 .52-.52l2.12-5a1 1 0 0 0-1.31-1.31m-2.49 5.1l-2.28 1l1-2.28l2.28-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compress(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 9h5a1 1 0 0 0 0-2h-4V3a1 1 0 0 0-2 0v5a1 1 0 0 0 1 1m-8 6H3a1 1 0 0 0 0 2h4v4a1 1 0 0 0 2 0v-5a1 1 0 0 0-1-1M8 2a1 1 0 0 0-1 1v4H3a1 1 0 0 0 0 2h5a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1m13 13h-5a1 1 0 0 0-1 1v5a1 1 0 0 0 2 0v-4h4a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompressAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.71 2.29a1 1 0 0 0-1.42 0l-5.79 5.8V6.5a1 1 0 0 0-2 0v4a1 1 0 0 0 .08.38a1 1 0 0 0 .54.54a1 1 0 0 0 .38.08h4a1 1 0 0 0 0-2h-1.59l5.8-5.79a1 1 0 0 0 0-1.42M10.88 12.58a1 1 0 0 0-.38-.08h-4a1 1 0 0 0 0 2h1.59l-5.8 5.79a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l5.79-5.8v1.59a1 1 0 0 0 2 0v-4a1 1 0 0 0-.08-.38a1 1 0 0 0-.54-.54"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompressAltLeft(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5 5.5a1 1 0 0 0-1 1v1.59l-5.79-5.8a1 1 0 0 0-1.42 1.42l5.8 5.79H6.5a1 1 0 0 0 0 2h4a1 1 0 0 0 .38-.08a1 1 0 0 0 .54-.54a1 1 0 0 0 .08-.38v-4a1 1 0 0 0-1-1m11.21 14.79l-5.8-5.79h1.59a1 1 0 0 0 0-2h-4a1 1 0 0 0-.38.08a1 1 0 0 0-.54.54a1 1 0 0 0-.08.38v4a1 1 0 0 0 2 0v-1.59l5.79 5.8a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompressArrows(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.38 13.08A1 1 0 0 0 10 13H6a1 1 0 0 0 0 2h1.59l-5.3 5.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0L9 16.41V18a1 1 0 0 0 2 0v-4a1 1 0 0 0-.08-.38a1 1 0 0 0-.54-.54M10 5a1 1 0 0 0-1 1v1.59l-5.29-5.3a1 1 0 0 0-1.42 1.42L7.59 9H6a1 1 0 0 0 0 2h4a1 1 0 0 0 .38-.08a1 1 0 0 0 .54-.54A1 1 0 0 0 11 10V6a1 1 0 0 0-1-1m3.62 5.92A1 1 0 0 0 14 11h4a1 1 0 0 0 0-2h-1.59l5.3-5.29a1 1 0 1 0-1.42-1.42L15 7.59V6a1 1 0 0 0-2 0v4a1 1 0 0 0 .08.38a1 1 0 0 0 .54.54M16.41 15H18a1 1 0 0 0 0-2h-4a1 1 0 0 0-.38.08a1 1 0 0 0-.54.54A1 1 0 0 0 13 14v4a1 1 0 0 0 2 0v-1.59l5.29 5.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompressLines(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 20h-4v-3.59l.79.8a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-2.5-2.5a1 1 0 0 0-.33-.21a1 1 0 0 0-.76 0a1 1 0 0 0-.33.21l-2.5 2.5a1 1 0 0 0 1.42 1.42l.79-.8V20H7a1 1 0 0 0 0 2h10a1 1 0 0 0 0-2M7 4h4v3.59l-.79-.8a1 1 0 1 0-1.42 1.42l2.5 2.5a1 1 0 0 0 .33.21a.94.94 0 0 0 .76 0a1 1 0 0 0 .33-.21l2.5-2.5a1 1 0 1 0-1.42-1.42l-.79.8V4h4a1 1 0 0 0 0-2H7a1 1 0 0 0 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompressPoint(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.71 20.29L15.41 14H17a1 1 0 0 0 0-2h-3.59l5.66-5.66a1 1 0 1 0-1.41-1.41L12 10.59V7a1 1 0 0 0-2 0v1.59l-6.29-6.3a1 1 0 0 0-1.42 1.42L8.59 10H7a1 1 0 0 0 0 2h3.59l-5.66 5.66a1 1 0 0 0 0 1.41a1 1 0 0 0 1.41 0L12 13.41V17a1 1 0 0 0 2 0v-1.59l6.29 6.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompressV(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.71 13.29a1 1 0 0 0-.33-.21a1 1 0 0 0-.76 0a1 1 0 0 0-.33.21l-2.5 2.5a1 1 0 0 0 1.42 1.42l.79-.8V21a1 1 0 0 0 2 0v-4.59l.79.8a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Zm-1.42-2.58a1 1 0 0 0 .33.21a.94.94 0 0 0 .76 0a1 1 0 0 0 .33-.21l2.5-2.5a1 1 0 1 0-1.42-1.42l-.79.8V3a1 1 0 0 0-2 0v4.59l-.79-.8a1 1 0 1 0-1.42 1.42Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Confused(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 9a1 1 0 1 0 1 1a1 1 0 0 0-1-1m5.66 4.56l-4.19 1.5A1 1 0 0 0 10.8 17a1 1 0 0 0 .34-.06l4.2-1.5a1 1 0 1 0-.68-1.88M15 9a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-3-7a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Constructor(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 9.67V9.5a7.95 7.95 0 0 0-5.59-7.62h-.06a8.32 8.32 0 0 0-2.59-.36A8.21 8.21 0 0 0 4 9.67a3 3 0 0 0 0 5.66a8 8 0 0 0 8 7.17h.23a8.13 8.13 0 0 0 7.68-7.16A3 3 0 0 0 20 9.67M12.18 20.5a6 6 0 0 1-6.09-5h11.77a6.09 6.09 0 0 1-5.68 5m6.82-7H5a1 1 0 0 1 0-2h2a1 1 0 0 0 0-2H6a6.4 6.4 0 0 1 3-5.15V7.5a1 1 0 0 0 2 0V3.59a7.34 7.34 0 0 1 .82-.09H12a6.64 6.64 0 0 1 1 .09V7.5a1 1 0 0 0 2 0V4.32a6.65 6.65 0 0 1 1.18.87A6 6 0 0 1 18 9.5h-1a1 1 0 0 0 0 2h2a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copy(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 8.94a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19a.32.32 0 0 0-.09 0a.88.88 0 0 0-.33-.11H10a3 3 0 0 0-3 3v1H6a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3v-1h1a3 3 0 0 0 3-3V9zm-6-3.53L17.59 8H16a1 1 0 0 1-1-1ZM15 19a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1h1v7a3 3 0 0 0 3 3h5Zm4-4a1 1 0 0 1-1 1h-8a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h3v3a3 3 0 0 0 3 3h3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CopyAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 20H8a3 3 0 0 1-3-3V7a1 1 0 0 0-2 0v10a5 5 0 0 0 5 5h8a1 1 0 0 0 0-2m5-11.06a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19h-.09L14.06 2H10a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V9zm-6-3.53L17.59 8H16a1 1 0 0 1-1-1ZM19 15a1 1 0 0 1-1 1h-8a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h3v3a3 3 0 0 0 3 3h3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CopyLandscape(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 19H6a3 3 0 0 1-3-3V8a1 1 0 0 0-2 0v8a5 5 0 0 0 5 5h12a1 1 0 0 0 0-2m5-9.06a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19h-.09L16.06 3H8a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-4zm-6-3.53L19.59 9H18a1 1 0 0 1-1-1ZM21 14a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h7v3a3 3 0 0 0 3 3h3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copyright(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 9h2a1 1 0 0 1 1 1a1 1 0 0 0 2 0a3 3 0 0 0-3-3h-2a3 3 0 0 0-3 3v4a3 3 0 0 0 3 3h2a3 3 0 0 0 3-3a1 1 0 0 0-2 0a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1m1-7a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerDownLeft(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.69 2a1 1 0 0 0-1 1v10.37a2 2 0 0 1-2 2h-8l2.92-2.92A1 1 0 0 0 9.24 11l-4.63 4.66a1.19 1.19 0 0 0-.22.33a1 1 0 0 0 0 .76a1 1 0 0 0 .22.33l4.63 4.63a1 1 0 0 0 .7.29a1 1 0 0 0 .71-.29a1 1 0 0 0 0-1.42l-2.92-2.92h8a4 4 0 0 0 4-4V3a1 1 0 0 0-1.04-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerDownRight(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.61 16a1.19 1.19 0 0 0-.22-.33L14.76 11a1 1 0 0 0-1.41 1.41l2.92 2.92h-7a3 3 0 0 1-3-3V3a1 1 0 1 0-2 0v9.37a5 5 0 0 0 5 5h7l-2.92 2.92a1 1 0 0 0 0 1.42a1 1 0 0 0 .71.29a1 1 0 0 0 .7-.29l4.63-4.63a1 1 0 0 0 .22-.33a1 1 0 0 0 0-.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerDownRightAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.92 12.62a1 1 0 0 0-.21-.33l-3-3a1 1 0 0 0-1.42 1.42l1.3 1.29H8a1 1 0 0 1-1-1V7a1 1 0 0 0-2 0v4a3 3 0 0 0 3 3h9.59l-1.3 1.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l3-3a1 1 0 0 0 .21-.33a1 1 0 0 0 0-.76"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerLeftDown(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 4.31h-9.37a5 5 0 0 0-5 5v7l-2.92-2.96a1 1 0 0 0-1.42 0a1 1 0 0 0 0 1.41l4.63 4.63a1 1 0 0 0 .33.22a.94.94 0 0 0 .76 0a1.19 1.19 0 0 0 .33-.22L13 14.76a1 1 0 0 0-1.41-1.41l-2.96 2.92v-7a3 3 0 0 1 3-3H21a1 1 0 0 0 0-2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerRightDown(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.71 13.35a1 1 0 0 0-1.42 0l-2.92 2.92v-8a4 4 0 0 0-4-4H3a1 1 0 1 0 0 2h10.37a2 2 0 0 1 2 2v8l-2.92-2.92A1 1 0 0 0 11 14.76l4.62 4.63a1.19 1.19 0 0 0 .33.22a.94.94 0 0 0 .76 0a1 1 0 0 0 .33-.22l4.63-4.63a1 1 0 0 0 .04-1.41"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerUpLeft(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.69 6.63h-7l2.92-2.92a1 1 0 0 0 0-1.42a1 1 0 0 0-1.41 0L4.61 6.92a1 1 0 0 0-.22.33a1 1 0 0 0 0 .76a1.19 1.19 0 0 0 .22.33L9.24 13a1 1 0 0 0 .7.3a1 1 0 0 0 .71-1.71L7.73 8.63h7a3 3 0 0 1 3 3V21a1 1 0 0 0 2 0v-9.37a5 5 0 0 0-5.04-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerUpLeftAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 9.5H7.41l1.3-1.29a1 1 0 0 0-1.42-1.42l-3 3a1 1 0 0 0-.21.33a1 1 0 0 0 0 .76a1 1 0 0 0 .21.33l3 3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-1.3-1.29H17a1 1 0 0 1 1 1v4a1 1 0 0 0 2 0v-4a3 3 0 0 0-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerUpRight(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.61 7.25a1 1 0 0 0-.22-.33l-4.63-4.63a1 1 0 0 0-1.41 0a1 1 0 0 0 0 1.42l2.92 2.92h-8a4 4 0 0 0-4 4V21a1 1 0 0 0 2 0V10.63a2 2 0 0 1 2-2h8l-2.92 2.92a1 1 0 0 0 .71 1.71a1 1 0 0 0 .7-.3l4.63-4.62a1.19 1.19 0 0 0 .22-.34a1 1 0 0 0 0-.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerUpRightAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.92 10.12a1 1 0 0 0-.21-.33l-3-3a1 1 0 1 0-1.42 1.42l1.3 1.29H7a3 3 0 0 0-3 3v4a1 1 0 0 0 2 0v-4a1 1 0 0 1 1-1h9.59l-1.3 1.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l3-3a1 1 0 0 0 .21-.33a1 1 0 0 0 0-.76"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Coronavirus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.5 9a1.5 1.5 0 1 0 1.5 1.5A1.5 1.5 0 0 0 9.5 9M9 14a1 1 0 1 0 1 1a1 1 0 0 0-1-1m6-6a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-.5 4a1.5 1.5 0 1 0 1.5 1.5a1.5 1.5 0 0 0-1.5-1.5m7.5-1h-1.059a8.952 8.952 0 0 0-1.916-4.61l.753-.754a1 1 0 0 0-1.414-1.414l-.753.753A8.952 8.952 0 0 0 13 3.059V2a1 1 0 0 0-2 0v1.059a8.952 8.952 0 0 0-4.61 1.916l-.754-.753a1 1 0 0 0-1.414 1.414l.753.753A8.952 8.952 0 0 0 3.059 11H2a1 1 0 0 0 0 2h1.059a8.952 8.952 0 0 0 1.916 4.61l-.753.754a1 1 0 1 0 1.414 1.414l.753-.753A8.952 8.952 0 0 0 11 20.941V22a1 1 0 0 0 2 0v-1.059a8.952 8.952 0 0 0 4.61-1.916l.754.753a1 1 0 0 0 1.414-1.414l-.753-.753A8.952 8.952 0 0 0 20.941 13H22a1 1 0 0 0 0-2m-4 2h.92A7.004 7.004 0 0 1 13 18.92V17a1 1 0 0 0-2 0v1.92A7.004 7.004 0 0 1 5.08 13H6a1 1 0 0 0 0-2h-.92A7.004 7.004 0 0 1 11 5.08V7a1 1 0 0 0 2 0V5.08A7.004 7.004 0 0 1 18.92 11H18a1 1 0 0 0 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreateDashboard(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 13H4a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1m-1 6H5v-4h4ZM20 3h-6a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1m-1 6h-4V5h4Zm1 7h-2v-2a1 1 0 0 0-2 0v2h-2a1 1 0 0 0 0 2h2v2a1 1 0 0 0 2 0v-2h2a1 1 0 0 0 0-2M10 3H4a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1M9 9H5V5h4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreativeCommonsPd(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 0 1-8-8a7.92 7.92 0 0 1 1.69-4.9l2.36 2.37A3.07 3.07 0 0 0 8 10v4a3 3 0 0 0 3 3h2a3 3 0 0 0 1.89-.69l2 2A7.92 7.92 0 0 1 12 20m-2-6v-2.59l3.46 3.46A.91.91 0 0 1 13 15h-2a1 1 0 0 1-1-1m8.31 2.9L16 14.53a3.07 3.07 0 0 0 0-.53a1 1 0 0 0-1-1a.91.91 0 0 0-.46.13l-4-4A.91.91 0 0 1 11 9h2a1 1 0 0 1 1 1a1 1 0 0 0 2 0a3 3 0 0 0-3-3h-2a3 3 0 0 0-1.89.69l-2-2A7.92 7.92 0 0 1 12 4a8 8 0 0 1 8 8a7.92 7.92 0 0 1-1.69 4.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCard(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 15h3a1 1 0 0 0 0-2H7a1 1 0 0 0 0 2M19 5H5a3 3 0 0 0-3 3v9a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V8a3 3 0 0 0-3-3m1 12a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-6h16Zm0-8H4V8a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCardSearch(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 17.57a4.3 4.3 0 1 0-3.67 2.06a4.37 4.37 0 0 0 2.24-.63l1.72 1.73a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM18 17a2.37 2.37 0 0 1-3.27 0a2.32 2.32 0 0 1 0-3.27a2.31 2.31 0 0 1 3.27 0A2.32 2.32 0 0 1 18 17m1-14H5a3 3 0 0 0-3 3v9a3 3 0 0 0 3 3h4a1 1 0 0 0 0-2H5a1 1 0 0 1-1-1V9h16v1a1 1 0 0 0 2 0V6a3 3 0 0 0-3-3m1 4H4V6a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Zm-10 4H7a1 1 0 0 0 0 2h3a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crockery(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 12.15V3a1 1 0 0 0-2 0v9.15a4.16 4.16 0 0 0-3 4c0 2.05 1.52 5.8 4 5.8s4-3.75 4-5.8a4.16 4.16 0 0 0-3-4M17 20c-.8 0-2-2.27-2-3.8a2.11 2.11 0 0 1 2-2.2a2.11 2.11 0 0 1 2 2.2c0 1.53-1.2 3.8-2 3.8M10 2a1 1 0 0 0-1 1v5.46l-1 .67V3a1 1 0 0 0-2 0v6.13l-1-.67V3a1 1 0 0 0-2 0v6a1 1 0 0 0 .45.83L6 11.54V21a1 1 0 0 0 2 0v-9.46l2.55-1.71A1 1 0 0 0 11 9V3a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CropAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 16h-1V7a1 1 0 0 0-1-1H8V5a1 1 0 0 0-2 0v1H5a1 1 0 0 0 0 2h1v9a1 1 0 0 0 1 1h9v1a1 1 0 0 0 2 0v-1h1a1 1 0 0 0 0-2m-3 0H8V8h8Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CropAltRotateLeft(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.71 16.29a1 1 0 0 0-1.42 1.42l.3.29H8a3 3 0 0 1-3-3v-2a1 1 0 0 0-2 0v2a5 5 0 0 0 5 5h.59l-.3.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l2-2a1 1 0 0 0 0-1.42ZM20 12h-1V5a1 1 0 0 0-1-1h-7V3a1 1 0 0 0-2 0v1H8a1 1 0 0 0 0 2h1v7a1 1 0 0 0 1 1h7v1a1 1 0 0 0 2 0v-1h1a1 1 0 0 0 0-2m-3 0h-6V6h6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CropAltRotateRight(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 4h-.59l.3-.29a1 1 0 1 0-1.42-1.42l-2 2a1 1 0 0 0 0 1.42l2 2a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-.3-.29H16a3 3 0 0 1 3 3v2a1 1 0 0 0 2 0V9a5 5 0 0 0-5-5m0 14h-1v-7a1 1 0 0 0-1-1H7V9a1 1 0 0 0-2 0v1H4a1 1 0 0 0 0 2h1v7a1 1 0 0 0 1 1h7v1a1 1 0 0 0 2 0v-1h1a1 1 0 0 0 0-2m-3 0H7v-6h6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crosshair(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 11h-1.07A8 8 0 0 0 13 4.07V3a1 1 0 0 0-2 0v1.07A8 8 0 0 0 4.07 11H3a1 1 0 0 0 0 2h1.07A8 8 0 0 0 11 19.93V21a1 1 0 0 0 2 0v-1.07A8 8 0 0 0 19.93 13H21a1 1 0 0 0 0-2m-4 2h.91A6 6 0 0 1 13 17.91V17a1 1 0 0 0-2 0v.91A6 6 0 0 1 6.09 13H7a1 1 0 0 0 0-2h-.91A6 6 0 0 1 11 6.09V7a1 1 0 0 0 2 0v-.91A6 6 0 0 1 17.91 11H17a1 1 0 0 0 0 2m-5-2a1 1 0 1 0 1 1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CrosshairAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 11a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0-9a10 10 0 1 0 10 10A10 10 0 0 0 12 2m1 17.93V17a1 1 0 0 0-2 0v2.93A8 8 0 0 1 4.07 13H7a1 1 0 0 0 0-2H4.07A8 8 0 0 1 11 4.07V7a1 1 0 0 0 2 0V4.07A8 8 0 0 1 19.93 11H17a1 1 0 0 0 0 2h2.93A8 8 0 0 1 13 19.93"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crosshairs(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 11h-1.07A8 8 0 0 0 13 4.07V3a1 1 0 0 0-2 0v1.07A8 8 0 0 0 4.07 11H3a1 1 0 0 0 0 2h1.07A8 8 0 0 0 11 19.93V21a1 1 0 0 0 2 0v-1.07A8 8 0 0 0 19.93 13H21a1 1 0 0 0 0-2m-9 7a6 6 0 1 1 6-6a6 6 0 0 1-6 6m0-9a3 3 0 1 0 3 3a3 3 0 0 0-3-3m0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CssThreeSimple(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.772 3.364A1 1 0 0 0 20 3H6a1 1 0 0 0 0 2h12.786l-.78 4H5a1 1 0 0 0 0 2h12.615l-1.163 5.955l-6.323 1.997l-5.41-1.7l.203-1.064a1 1 0 0 0-1.964-.375l-.37 1.94a1 1 0 0 0 .682 1.141l6.56 2.06a1.002 1.002 0 0 0 .601 0l7.19-2.27a1 1 0 0 0 .68-.763l2.68-13.73a1 1 0 0 0-.209-.827"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cube(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.47 7.37v-.08l-.06-.15a.71.71 0 0 0-.07-.09a.94.94 0 0 0-.09-.12l-.09-.07l-.16-.08l-7.5-4.63a1 1 0 0 0-1.06 0L4 6.78l-.09.08l-.09.07a.94.94 0 0 0-.09.12a.71.71 0 0 0-.07.09l-.06.15v.08a1.15 1.15 0 0 0 0 .26v8.74a1 1 0 0 0 .47.85l7.5 4.63a.47.47 0 0 0 .15.06s.05 0 .08 0a.86.86 0 0 0 .52 0h.08a.47.47 0 0 0 .15-.06L20 17.22a1 1 0 0 0 .47-.85V7.63a1.15 1.15 0 0 0 0-.26M11 19.21l-5.5-3.4V9.43l5.5 3.39Zm1-8.12L6.4 7.63L12 4.18l5.6 3.45Zm6.5 4.72l-5.5 3.4v-6.39l5.5-3.39Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dashboard(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.088 4.955c-.007-.008-.01-.019-.017-.026s-.018-.01-.026-.018a9.979 9.979 0 0 0-14.09 0c-.008.008-.018.01-.026.018s-.01.018-.017.026a10 10 0 1 0 14.176 0M12 20a7.985 7.985 0 0 1-6.235-3H9.78a2.964 2.964 0 0 0 4.44 0h4.015A7.985 7.985 0 0 1 12 20m-1-5a1 1 0 1 1 1 1a1.001 1.001 0 0 1-1-1m8.41.002L19.4 15H15a2.995 2.995 0 0 0-2-2.816V9a1 1 0 0 0-2 0v3.184A2.995 2.995 0 0 0 9 15H4.6l-.01.002A7.93 7.93 0 0 1 4.07 13H5a1 1 0 0 0 0-2h-.93a7.951 7.951 0 0 1 1.619-3.898l.654.655a1 1 0 1 0 1.414-1.414l-.654-.655A7.952 7.952 0 0 1 11 4.07V5a1 1 0 0 0 2 0v-.93a7.952 7.952 0 0 1 3.897 1.618l-.654.655a1 1 0 1 0 1.414 1.414l.654-.655A7.951 7.951 0 0 1 19.931 11H19a1 1 0 0 0 0 2h.93a7.93 7.93 0 0 1-.52 2.002"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DataSharing(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.52 3.87a5 5 0 0 0-9.08.13H7a3 3 0 0 0-3 3v4a1 1 0 0 0 2 0V7a1 1 0 0 1 1-1h2.78A3 3 0 0 0 9 8a3 3 0 0 0 3 3h7.33a3.66 3.66 0 0 0 1.19-7.13M19.33 9H12a1 1 0 0 1 0-2a1 1 0 0 0 1-1a3 3 0 0 1 5.84-1a1 1 0 0 0 .78.67A1.65 1.65 0 0 1 21 7.33A1.67 1.67 0 0 1 19.33 9M19 13a1 1 0 0 0-1 1v3a1 1 0 0 1-1 1h-2.26a3.66 3.66 0 0 0-2.22-2.13a5 5 0 0 0-9.45 1.28A3 3 0 0 0 4 23h7.33a3.66 3.66 0 0 0 3.6-3H17a3 3 0 0 0 3-3v-3a1 1 0 0 0-1-1m-7.67 8H4a1 1 0 0 1 0-2a1 1 0 0 0 1-1a3 3 0 0 1 5.84-1a1 1 0 0 0 .78.67A1.65 1.65 0 0 1 13 19.33A1.67 1.67 0 0 1 11.33 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Database(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 16.5a1 1 0 1 0 1 1a1 1 0 0 0-1-1M12 2C8 2 4 3.37 4 6v12c0 2.63 4 4 8 4s8-1.37 8-4V6c0-2.63-4-4-8-4m6 16c0 .71-2.28 2-6 2s-6-1.29-6-2v-3.27A13.16 13.16 0 0 0 12 16a13.16 13.16 0 0 0 6-1.27Zm0-6c0 .71-2.28 2-6 2s-6-1.29-6-2V8.73A13.16 13.16 0 0 0 12 10a13.16 13.16 0 0 0 6-1.27Zm-6-4C8.28 8 6 6.71 6 6s2.28-2 6-2s6 1.29 6 2s-2.28 2-6 2m-4 2.5a1 1 0 1 0 1 1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabaseAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 17a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0-6a1 1 0 1 0 1 1a1 1 0 0 0-1-1m8-9H8a4 4 0 0 0-4 4v12a4 4 0 0 0 4 4h8a4 4 0 0 0 4-4V6a4 4 0 0 0-4-4m2 16a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2v-2.56A3.91 3.91 0 0 0 8 16h8a3.91 3.91 0 0 0 2-.56Zm0-6a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2V9.44A3.91 3.91 0 0 0 8 10h8a3.91 3.91 0 0 0 2-.56Zm-2-4H8a2 2 0 0 1 0-4h8a2 2 0 0 1 0 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Desert(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 2a4 4 0 1 0 4 4a4 4 0 0 0-4-4m0 6a2 2 0 1 1 2-2a2 2 0 0 1-2 2m-3 12h-2v-2a1 1 0 0 0-2 0v2H9v-3.38l3.45-1.73A1 1 0 0 0 13 14v-4a1 1 0 0 0-2 0v3.38l-2 1V8a1 1 0 0 0-2 0v8.38l-2-1V13a1 1 0 0 0-2 0v3a1 1 0 0 0 .55.89L7 18.62V20H3a1 1 0 0 0 0 2h12a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Desktop(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 3H5a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h6v2H7a1 1 0 0 0 0 2h10a1 1 0 0 0 0-2h-4v-2h6a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3m1 11a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DesktopAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 2H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h3v2H5a1 1 0 0 0 0 2h14a1 1 0 0 0 0-2h-3v-2h3a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3m-5 18h-4v-2h4Zm6-5a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-1h16Zm0-3H4V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DesktopAltSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.71 20.29l-18-18a1 1 0 0 0-1.42 1.42l1 1A3 3 0 0 0 3 6v8a3 3 0 0 0 3 3h3v2H6a1 1 0 0 0 0 2h12a1 1 0 0 0 .93-.66l1.36 1.37a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42M5 6.41L9.59 11H5ZM13 19h-2v-2h2Zm-7-4a1 1 0 0 1-1-1v-1h6.59l2 2Zm9 4v-2h.59l2 2ZM9.66 5H18a1 1 0 0 1 1 1v5h-1.34a1 1 0 0 0 0 2H19v1a.37.37 0 0 1 0 .11a1 1 0 0 0 .78 1.18h.2a1 1 0 0 0 1-.8A2.84 2.84 0 0 0 21 14V6a3 3 0 0 0-3-3H9.66a1 1 0 0 0 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DesktopCloudAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 10H10a3 3 0 0 0 1.07-5.8a4 4 0 0 0-7.48 1A2.5 2.5 0 0 0 4.5 10m0-3a1 1 0 0 0 1-1a2 2 0 0 1 3.89-.64a1 1 0 0 0 .78.66A1 1 0 0 1 11 7a1 1 0 0 1-1 1H4.5a.5.5 0 0 1 0-1M19 2h-5a1 1 0 0 0 0 2h5a1 1 0 0 1 1 1v7H3a1 1 0 0 0-1 1v2a3 3 0 0 0 3 3h3v2H5a1 1 0 0 0 0 2h14a1 1 0 0 0 0-2h-3v-2h3a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3m-5 18h-4v-2h4Zm6-5a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-1h16Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DesktopSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.66 6H18a1 1 0 0 1 1 1v6a.94.94 0 0 1-.14.5a1 1 0 0 0 .31 1.38a.94.94 0 0 0 .53.16a1 1 0 0 0 .84-.46A2.94 2.94 0 0 0 21 13V7a3 3 0 0 0-3-3h-7.34a1 1 0 0 0 0 2m11.05 14.29L5.86 4.45L3.71 2.29a1 1 0 0 0-1.42 1.42l1.4 1.39A3 3 0 0 0 3 7v6a3 3 0 0 0 3 3h5v2H7a1 1 0 0 0 0 2h10a1 1 0 0 0 .93-.66l2.36 2.37a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42M6 14a1 1 0 0 1-1-1V7a1 1 0 0 1 .12-.46L12.59 14Zm7 4v-2h1.59l2 2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dialpad(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 2.25H3a.76.76 0 0 0-.75.75v4a.76.76 0 0 0 .75.75h4A.76.76 0 0 0 7.75 7V3A.76.76 0 0 0 7 2.25m-.75 4h-2.5v-2.5h2.5Zm14.75-4h-4a.76.76 0 0 0-.75.75v4a.76.76 0 0 0 .75.75h4a.76.76 0 0 0 .75-.75V3a.76.76 0 0 0-.75-.75m-.75 4h-2.5v-2.5h2.5Zm-6.25-4h-4a.76.76 0 0 0-.75.75v4a.76.76 0 0 0 .75.75h4a.76.76 0 0 0 .75-.75V3a.76.76 0 0 0-.75-.75m-.75 4h-2.5v-2.5h2.5ZM7 9.25H3a.76.76 0 0 0-.75.75v4a.76.76 0 0 0 .75.75h4a.76.76 0 0 0 .75-.75v-4A.76.76 0 0 0 7 9.25m-.75 4h-2.5v-2.5h2.5Zm7.75-4h-4a.76.76 0 0 0-.75.75v4a.76.76 0 0 0 .75.75h4a.76.76 0 0 0 .75-.75v-4a.76.76 0 0 0-.75-.75m-.75 4h-2.5v-2.5h2.5Zm7.75-4h-4a.76.76 0 0 0-.75.75v4a.76.76 0 0 0 .75.75h4a.76.76 0 0 0 .75-.75v-4a.76.76 0 0 0-.75-.75m-.75 4h-2.5v-2.5h2.5Zm-6.25 3h-4a.76.76 0 0 0-.75.75v4a.76.76 0 0 0 .75.75h4a.76.76 0 0 0 .75-.75v-4a.76.76 0 0 0-.75-.75m-.75 4h-2.5v-2.5h2.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DialpadAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 9.25A2.75 2.75 0 1 0 7.75 12A2.75 2.75 0 0 0 5 9.25m0 4A1.25 1.25 0 1 1 6.25 12A1.25 1.25 0 0 1 5 13.25m7-4A2.75 2.75 0 1 0 14.75 12A2.75 2.75 0 0 0 12 9.25m0 4A1.25 1.25 0 1 1 13.25 12A1.25 1.25 0 0 1 12 13.25m7-5.5A2.75 2.75 0 1 0 16.25 5A2.75 2.75 0 0 0 19 7.75m0-4A1.25 1.25 0 1 1 17.75 5A1.25 1.25 0 0 1 19 3.75m0 5.5A2.75 2.75 0 1 0 21.75 12A2.75 2.75 0 0 0 19 9.25m0 4A1.25 1.25 0 1 1 20.25 12A1.25 1.25 0 0 1 19 13.25m-14-11A2.75 2.75 0 1 0 7.75 5A2.75 2.75 0 0 0 5 2.25m0 4A1.25 1.25 0 1 1 6.25 5A1.25 1.25 0 0 1 5 6.25m7 10A2.75 2.75 0 1 0 14.75 19A2.75 2.75 0 0 0 12 16.25m0 4A1.25 1.25 0 1 1 13.25 19A1.25 1.25 0 0 1 12 20.25m0-18A2.75 2.75 0 1 0 14.75 5A2.75 2.75 0 0 0 12 2.25m0 4A1.25 1.25 0 1 1 13.25 5A1.25 1.25 0 0 1 12 6.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Diamond(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20 10.35l-5.78-7.41A3.06 3.06 0 0 0 9.75 3L4 10.35A3.05 3.05 0 0 0 3.51 12A3.09 3.09 0 0 0 4 13.58l.06.07l5.74 7.41A3 3 0 0 0 12 22a3.06 3.06 0 0 0 2.26-1L20 13.65a3 3 0 0 0-.06-3.3Zm-1.57 2.14l-5.67 7.22a1.11 1.11 0 0 1-1.42.07l-5.69-7.31a1 1 0 0 1-.14-.47a1.11 1.11 0 0 1 .1-.45l5.67-7.22a1.11 1.11 0 0 1 1.42-.07l5.63 7.28a1 1 0 0 1 .16.54a1.11 1.11 0 0 1-.1.41Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Diary(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2H5a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h1v1a1 1 0 0 0 1 1a1 1 0 0 0 1-1v-1h9a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3m-3 16H6V4h8Zm4-1a1 1 0 0 1-1 1h-1V4h1a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiaryAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2H5a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h12a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3M8 20H6V4h2Zm10-1a1 1 0 0 1-1 1h-7V4h7a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceFive(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2H7a5 5 0 0 0-5 5v10a5 5 0 0 0 5 5h10a5 5 0 0 0 5-5V7a5 5 0 0 0-5-5m3 15a3 3 0 0 1-3 3H7a3 3 0 0 1-3-3V7a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3ZM8 15a1 1 0 1 0 1 1a1 1 0 0 0-1-1m4-4a1 1 0 1 0 1 1a1 1 0 0 0-1-1M8 7a1 1 0 1 0 1 1a1 1 0 0 0-1-1m8 8a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0-8a1 1 0 1 0 1 1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceFour(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 14a1 1 0 1 0 1 1a1 1 0 0 0-1-1M9 8a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0 6a1 1 0 1 0 1 1a1 1 0 0 0-1-1m6-6a1 1 0 1 0 1 1a1 1 0 0 0-1-1m2-6H7a5 5 0 0 0-5 5v10a5 5 0 0 0 5 5h10a5 5 0 0 0 5-5V7a5 5 0 0 0-5-5m3 15a3 3 0 0 1-3 3H7a3 3 0 0 1-3-3V7a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceOne(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2H7a5 5 0 0 0-5 5v10a5 5 0 0 0 5 5h10a5 5 0 0 0 5-5V7a5 5 0 0 0-5-5m3 15a3 3 0 0 1-3 3H7a3 3 0 0 1-3-3V7a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3Zm-8-6a1 1 0 1 0 1 1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceSix(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 11a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0 4a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0-8a1 1 0 1 0 1 1a1 1 0 0 0-1-1m6 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1m2-5H7a5 5 0 0 0-5 5v10a5 5 0 0 0 5 5h10a5 5 0 0 0 5-5V7a5 5 0 0 0-5-5m3 15a3 3 0 0 1-3 3H7a3 3 0 0 1-3-3V7a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3Zm-5-6a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0 4a1 1 0 1 0 1 1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceThree(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 11a1 1 0 1 0 1 1a1 1 0 0 0-1-1M8 7a1 1 0 1 0 1 1a1 1 0 0 0-1-1m9-5H7a5 5 0 0 0-5 5v10a5 5 0 0 0 5 5h10a5 5 0 0 0 5-5V7a5 5 0 0 0-5-5m3 15a3 3 0 0 1-3 3H7a3 3 0 0 1-3-3V7a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3Zm-4-2a1 1 0 1 0 1 1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceTwo(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 11a1 1 0 1 0 1 1a1 1 0 0 0-1-1m8-9H7a5 5 0 0 0-5 5v10a5 5 0 0 0 5 5h10a5 5 0 0 0 5-5V7a5 5 0 0 0-5-5m3 15a3 3 0 0 1-3 3H7a3 3 0 0 1-3-3V7a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3Zm-5-6a1 1 0 1 0 1 1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Direction(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.71 10.21L12 7.91l2.29 2.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-3-3a1 1 0 0 0-1.42 0l-3 3a1 1 0 0 0 1.42 1.42m4.58 4.58L12 17.09l-2.29-2.3a1 1 0 0 0-1.42 1.42l3 3a1 1 0 0 0 1.42 0l3-3a1 1 0 0 0-1.42-1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Directions(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.12 9.88l-7-7a3.08 3.08 0 0 0-4.24 0l-7 7a3 3 0 0 0 0 4.24l7 7a3 3 0 0 0 4.24 0l7-7a3 3 0 0 0 0-4.24m-1.41 2.83l-7 7a1 1 0 0 1-1.42 0l-7-7a1 1 0 0 1 0-1.42l7-7a1 1 0 0 1 1.42 0l7 7a1 1 0 0 1 0 1.42m-5.5-3.42a1 1 0 0 0-1.42 1.42l.3.29H9.5a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0v-1h2.59l-.3.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l2-2a1 1 0 0 0 0-1.42Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Discord(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.893 10.63a1.146 1.146 0 1 0 1.049 1.141a1.096 1.096 0 0 0-1.05-1.141M18.89 2H5.11A2.114 2.114 0 0 0 3 4.12v13.906a2.114 2.114 0 0 0 2.109 2.12h11.664l-.546-1.904l1.317 1.224l1.245 1.152L21 22.572V4.119A2.114 2.114 0 0 0 18.891 2m-3.97 13.433s-.37-.442-.679-.833a3.246 3.246 0 0 0 1.862-1.224a5.878 5.878 0 0 1-1.183.607a6.77 6.77 0 0 1-1.491.442a7.206 7.206 0 0 1-2.664-.01a8.645 8.645 0 0 1-1.512-.442a6.037 6.037 0 0 1-.751-.35c-.031-.02-.062-.03-.093-.051a.142.142 0 0 1-.04-.031c-.186-.103-.289-.175-.289-.175a3.2 3.2 0 0 0 1.8 1.214c-.308.39-.689.853-.689.853a3.729 3.729 0 0 1-3.137-1.563a13.775 13.775 0 0 1 1.481-5.997a5.086 5.086 0 0 1 2.89-1.08l.103.124a6.938 6.938 0 0 0-2.705 1.347s.226-.123.607-.298a7.722 7.722 0 0 1 2.335-.648a1.005 1.005 0 0 1 .175-.02a8.702 8.702 0 0 1 2.077-.021a8.384 8.384 0 0 1 3.096.987a6.846 6.846 0 0 0-2.56-1.306l.143-.165a5.086 5.086 0 0 1 2.89 1.08a13.774 13.774 0 0 1 1.482 5.997a3.76 3.76 0 0 1-3.148 1.563"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DizzyMeh(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9 11.71l.29-.3l.29.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-.3-.29l.3-.29a1 1 0 0 0-1.46-1.42l-.29.3l-.25-.3a1 1 0 1 0-1.46 1.42l.3.29l-.3.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0ZM15 14H9a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8m5-11.71a1 1 0 0 0-1.42 0l-.29.3l-.29-.3a1 1 0 0 0-1.42 1.42l.3.29l-.3.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l.29-.3l.29.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-.3-.29l.3-.29a1 1 0 0 0 0-1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dna(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.2 13.73a1 1 0 0 0-1.41-.05A11.18 11.18 0 0 0 4 22a1 1 0 0 0 2 0a9.15 9.15 0 0 1 3.15-6.86a1 1 0 0 0 .05-1.41m10.17 4.64a10.86 10.86 0 0 0-1.6-3A14.31 14.31 0 0 0 14.06 12C16.3 10.57 20 7.4 20 2a1 1 0 0 0-2 0c0 5.4-4.59 8.17-6 8.89A13.42 13.42 0 0 1 9.31 9H12a1 1 0 0 0 0-2H7.55a9.39 9.39 0 0 1-1-2H15a1 1 0 0 0 0-2H6.06A8.14 8.14 0 0 1 6 2a1 1 0 0 0-2 0c0 7.57 7.3 10.79 7.61 10.92A12.93 12.93 0 0 1 14.7 15H12a1 1 0 0 0 0 2h4.43a9.07 9.07 0 0 1 1 2H9a1 1 0 0 0 0 2h8.94a8.26 8.26 0 0 1 .06 1a1 1 0 0 0 2 0a10.5 10.5 0 0 0-.22-2.19a9.23 9.23 0 0 0-.41-1.44"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Docker(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.8 8.8h1.8c.1 0 .2-.1.2-.2V7.1c0-.1-.1-.2-.2-.2H8.8c-.1 0-.2.1-.2.2v1.6s.1.1.2.1m2.4 2.3H13c.1 0 .2-.1.2-.2V9.3c0-.1-.1-.2-.2-.2h-1.8c-.1 0-.2.1-.2.2v1.6c0 .1.1.2.2.2m0-2.3H13c.1 0 .2-.1.2-.2V7.1l-.2-.2h-1.8c-.1 0-.2.1-.2.2v1.6s.1.1.2.1m2.5 2.3h1.8c.1 0 .2-.1.2-.2V9.3c0-.1-.1-.2-.2-.2h-1.8c-.1 0-.2.1-.2.2v1.6c0 .1.1.2.2.2m-2.5-4.6H13c.1 0 .2-.1.2-.2V4.8c0-.1-.1-.2-.2-.2h-1.8c-.1 0-.2.1-.2.2v1.6c0 .1.1.1.2.1m-7.4 4.6h1.8c.1 0 .2-.1.2-.2V9.3c0-.1-.1-.2-.2-.2H3.8c-.1 0-.2.1-.2.2v1.6zm18-1c-.5-.3-1.1-.5-1.6-.4c-.3 0-.6 0-.8.1c-.2-.9-.7-1.7-1.4-2.1l-.3-.2l-.2.3c-.3.2-.5.6-.6 1.1c-.2.8-.1 1.6.3 2.2c-.5.2-1 .3-1.5.4H2.6c-.3 0-.6.3-.6.6c0 1.2.2 2.3.6 3.4c.4 1.1 1.1 2 2 2.6c1.4.7 2.9 1 4.4.9c.8 0 1.6-.1 2.4-.2c1.1-.2 2.2-.6 3.2-1.2c.8-.5 1.5-1.1 2.2-1.8c.9-1.1 1.6-2.3 2.1-3.7h.2c.8 0 1.6-.3 2.2-.8c.3-.2.5-.5.6-.9l.1-.2zm-15.5 1H8c.1 0 .2-.1.2-.2V9.3c0-.1-.1-.2-.2-.2H6.3c-.1 0-.2.1-.2.2v1.6c0 .1.1.2.2.2m0-2.3H8c.1 0 .2-.1.2-.2V7.1c0-.1-.1-.2-.2-.2H6.3c-.1 0-.2.1-.2.2v1.6s.1.1.2.1m2.5 2.3h1.8c.1 0 .2-.1.2-.2V9.3c0-.1-.1-.2-.2-.2H8.8c-.1 0-.2.1-.2.2v1.6c0 .1.1.2.2.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentInfo(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 14a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0v-2a1 1 0 0 0-1-1m.38-2.92a1 1 0 0 0-.58-.08l-.18.06l-.18.09l-.15.12A1 1 0 0 0 11 12a1 1 0 0 0 .29.71a1 1 0 0 0 .33.21a.84.84 0 0 0 .38.08a1 1 0 0 0 .71-.29A1 1 0 0 0 13 12a1 1 0 0 0-.29-.71a1.15 1.15 0 0 0-.33-.21M20 8.94a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19h-.1a1.1 1.1 0 0 0-.31-.11H7a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V9zm-6-3.53L16.59 8H15a1 1 0 0 1-1-1ZM18 19a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentLayoutCenter(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 12h6a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1H9a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1m1-6h4v4h-4Zm11 4h-2a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2m-2-2h2a1 1 0 0 0 0-2h-2a1 1 0 0 0 0 2M3 8h2a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2m0 4h2a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2m18 2H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2m-8 4H3a1 1 0 0 0 0 2h10a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentLayoutLeft(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 12h6a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1m1-6h4v4H4Zm9 2h8a1 1 0 0 0 0-2h-8a1 1 0 0 0 0 2m0 10H3a1 1 0 0 0 0 2h10a1 1 0 0 0 0-2m8-4H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2m0-4h-8a1 1 0 0 0 0 2h8a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentLayoutRight(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 18H3a1 1 0 0 0 0 2h10a1 1 0 0 0 0-2M3 8h8a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2m0 4h8a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2m18 2H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2m0-10h-6a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1m-1 6h-4V6h4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DollarAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 11h-1V7h2a1 1 0 0 1 1 1a1 1 0 0 0 2 0a3 3 0 0 0-3-3h-2V3a1 1 0 0 0-2 0v2h-1a4 4 0 0 0 0 8h1v4H9a1 1 0 0 1-1-1a1 1 0 0 0-2 0a3 3 0 0 0 3 3h2v2a1 1 0 0 0 2 0v-2h1a4 4 0 0 0 0-8m-3 0h-1a2 2 0 0 1 0-4h1Zm3 6h-1v-4h1a2 2 0 0 1 0 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DollarSign(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 11h-4a2 2 0 0 1 0-4h5a1 1 0 0 1 1 1a1 1 0 0 0 2 0a3 3 0 0 0-3-3h-2V3a1 1 0 0 0-2 0v2h-1a4 4 0 0 0 0 8h4a2 2 0 0 1 0 4H9a1 1 0 0 1-1-1a1 1 0 0 0-2 0a3 3 0 0 0 3 3h2v2a1 1 0 0 0 2 0v-2h1a4 4 0 0 0 0-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DollarSignAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 11h-4a2 2 0 0 1 0-4h5a1 1 0 0 1 1 1a1 1 0 0 0 2 0a3 3 0 0 0-3-3h-2V3a1 1 0 0 0-2 0v2h-1a4 4 0 0 0 0 8h4a2 2 0 0 1 0 4H9a1 1 0 0 1-1-1a1 1 0 0 0-2 0a3 3 0 0 0 3 3h2v2a1 1 0 0 0 2 0v-2h1a4 4 0 0 0 0-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownloadAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.29 13.29a1 1 0 0 0 0 1.42l3 3a1 1 0 0 0 1.42 0l3-3a1 1 0 0 0-1.42-1.42L13 14.59V3a1 1 0 0 0-2 0v11.59l-1.29-1.3a1 1 0 0 0-1.42 0M18 9h-2a1 1 0 0 0 0 2h2a1 1 0 0 1 1 1v7a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1v-7a1 1 0 0 1 1-1h2a1 1 0 0 0 0-2H6a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-7a3 3 0 0 0-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Draggabledots(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 10a2 2 0 1 0 2 2a2 2 0 0 0-2-2m0 7a2 2 0 1 0 2 2a2 2 0 0 0-2-2m7-10a2 2 0 1 0-2-2a2 2 0 0 0 2 2m-7-4a2 2 0 1 0 2 2a2 2 0 0 0-2-2m7 14a2 2 0 1 0 2 2a2 2 0 0 0-2-2m0-7a2 2 0 1 0 2 2a2 2 0 0 0-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dribbble(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 2a7.94 7.94 0 0 1 5.51 2.22a24.93 24.93 0 0 1-4.83 2.18a29.07 29.07 0 0 0-2.87-4.09A7.94 7.94 0 0 1 12 4M7.89 5.15A27.16 27.16 0 0 1 10.7 9a25.11 25.11 0 0 1-6 .74h-.36a8 8 0 0 1 3.55-4.59M6 17.31A7.9 7.9 0 0 1 4 12v-.29h.68a26.67 26.67 0 0 0 7-1c.32.61.62 1.24.91 1.89a14.3 14.3 0 0 0-4.29 2.41l-.3.24a21 21 0 0 0-2 2.06M12 20a7.92 7.92 0 0 1-4.47-1.37a17.92 17.92 0 0 1 1.56-1.58l.32-.27a12.63 12.63 0 0 1 4-2.27a32 32 0 0 1 1.4 5A8.08 8.08 0 0 1 12 20m4.63-1.49a34.87 34.87 0 0 0-1.28-4.46h.34a.25.25 0 0 1 .12 0h.69a9.43 9.43 0 0 1 3.09.53a7.94 7.94 0 0 1-2.96 3.93M16.5 12h-.62a1.56 1.56 0 0 0-.39 0a6.64 6.64 0 0 0-.81.1h-.1c-.29-.67-.59-1.32-.92-2a26.57 26.57 0 0 0 5.13-2.4A8 8 0 0 1 20 12v.51a11.48 11.48 0 0 0-3.5-.51"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Drill(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 4H9a1 1 0 0 0-1 1v2H3a1 1 0 0 0 0 2h5v4a1 1 0 0 0 2 0v-1h4v7a1 1 0 0 0 1 1h2a3 3 0 0 0 3-3v-5.18A3 3 0 0 0 22 9V7a3 3 0 0 0-3-3m-1 13a1 1 0 0 1-1 1h-1v-6h2Zm2-8a1 1 0 0 1-1 1h-9V6h6v1a1 1 0 0 0 2 0V6h1a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dropbox(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.61 11.81l-3.25-2.53l3.26-2.56a1 1 0 0 0 .38-.86a1.06 1.06 0 0 0-.5-.8l-5.14-2.93a1 1 0 0 0-1.15.12L12 5.05l-3.21-2.8a1 1 0 0 0-1.15-.12L2.5 5.06a1.06 1.06 0 0 0-.5.8a1 1 0 0 0 .38.86l3.26 2.56l-3.25 2.49a1 1 0 0 0-.39.86a1 1 0 0 0 .5.8l3.41 2v2.47a1 1 0 0 0 .48.85l5.09 3.13a1 1 0 0 0 1 0l5.09-3.13a1 1 0 0 0 .48-.85v-2.48l3.41-1.95a1 1 0 0 0 .5-.8a1 1 0 0 0-.35-.86M16 4.22l3.23 1.84l-2.55 2l-3-1.84Zm-1.09 5.12l-2.91 2l-2.91-2L12 7.55ZM4.79 6.06L8 4.22l2.31 2l-3 1.84Zm0 6.39l2.5-1.92l3 2.08l-2.22 1.73Zm11.29 4.86L12 19.83l-4.09-2.52v-.8a1 1 0 0 0 .85-.18l3.24-2.5l3.24 2.5a1 1 0 0 0 .61.21a1 1 0 0 0 .24 0Zm-.15-3l-2.23-1.7l3-2.08l2.51 1.94Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dumbbell(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m17.48 6.55l-2.84-2.84a1 1 0 0 0-1.42 0a1 1 0 0 0 0 1.41l2.12 2.12l-8.1 8.1l-2.12-2.12a1 1 0 0 0-1.41 0a1 1 0 0 0 0 1.42l2.81 2.81l2.81 2.81a1 1 0 0 0 .71.3a1 1 0 0 0 .71-1.71l-2.09-2.09l8.1-8.1l2.12 2.12a1 1 0 1 0 1.41-1.42ZM3.71 17.46a1 1 0 0 0-1.42 1.42l2.83 2.83a1 1 0 0 0 .71.29a1 1 0 0 0 .71-.29a1 1 0 0 0 0-1.42Zm18-12.34l-2.83-2.83a1 1 0 0 0-1.42 1.42l2.83 2.83a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ear(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 8a1 1 0 0 1 1 1a1 1 0 0 0 2 0a3 3 0 1 0-5.018 2.22c.01.01.162.17.194.216a.988.988 0 0 1 0 1.119a1 1 0 1 0 1.648 1.13a2.983 2.983 0 0 0-.005-3.388a7.124 7.124 0 0 0-.49-.557a1.055 1.055 0 0 1-.16-.181A1 1 0 0 1 12 8m0-6a7 7 0 0 0-6.762 8.812a1 1 0 0 0 1.932-.518A5 5 0 1 1 17 9a5.114 5.114 0 0 1-.705 2.567L12.73 19A2.005 2.005 0 0 1 11 20a2.027 2.027 0 0 1-2-2a1.992 1.992 0 0 1 .269-.999a1 1 0 0 0-1.733-1.002a3.999 3.999 0 1 0 6.963 3.934l3.563-7.433A7 7 0 0 0 12 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Edit(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 12a1 1 0 0 0-1 1v6a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h6a1 1 0 0 0 0-2H5a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-6a1 1 0 0 0-1-1m-15 .76V17a1 1 0 0 0 1 1h4.24a1 1 0 0 0 .71-.29l6.92-6.93L21.71 8a1 1 0 0 0 0-1.42l-4.24-4.29a1 1 0 0 0-1.42 0l-2.82 2.83l-6.94 6.93a1 1 0 0 0-.29.71m10.76-8.35l2.83 2.83l-1.42 1.42l-2.83-2.83ZM8 13.17l5.93-5.93l2.83 2.83L10.83 16H8Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 18h4.24a1 1 0 0 0 .71-.29l6.92-6.93L19.71 8a1 1 0 0 0 0-1.42l-4.24-4.29a1 1 0 0 0-1.42 0l-2.82 2.83l-6.94 6.93a1 1 0 0 0-.29.71V17a1 1 0 0 0 1 1m9.76-13.59l2.83 2.83l-1.42 1.42l-2.83-2.83ZM6 13.17l5.93-5.93l2.83 2.83L8.83 16H6ZM21 20H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EighteenPlus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 5h1v1a1 1 0 0 0 2 0V5h1a1 1 0 0 0 0-2h-1V2a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2M7 7v10a1 1 0 0 0 2 0V7a1 1 0 0 0-2 0m14.6 2a1 1 0 0 0-.78 1.18a9 9 0 1 1-7-7a1 1 0 1 0 .4-2A10.8 10.8 0 0 0 12 1a11 11 0 1 0 11 11a10.8 10.8 0 0 0-.22-2.2A1 1 0 0 0 21.6 9M11 9v1a3 3 0 0 0 .78 2a3 3 0 0 0-.78 2v1a3 3 0 0 0 3 3h1a3 3 0 0 0 3-3v-1a3 3 0 0 0-.78-2a3 3 0 0 0 .78-2V9a3 3 0 0 0-3-3h-1a3 3 0 0 0-3 3m5 6a1 1 0 0 1-1 1h-1a1 1 0 0 1-1-1v-1a1 1 0 0 1 1-1h1a1 1 0 0 1 1 1Zm0-6v1a1 1 0 0 1-1 1h-1a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1h1a1 1 0 0 1 1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ElipsisDoubleValt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 17c-1.1 0-2 .9-2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2m7-10c1.1 0 2-.9 2-2s-.9-2-2-2s-2 .9-2 2s.9 2 2 2m-7 3c-1.1 0-2 .9-2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2m7 0c-1.1 0-2 .9-2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2m0 7c-1.1 0-2 .9-2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2m-7-14c-1.1 0-2 .9-2 2s.9 2 2 2s2-.9 2-2s-.9-2-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EllipsisH(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 10a2 2 0 1 0 2 2a2 2 0 0 0-2-2m-7 0a2 2 0 1 0 2 2a2 2 0 0 0-2-2m14 0a2 2 0 1 0 2 2a2 2 0 0 0-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EllipsisV(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 7a2 2 0 1 0-2-2a2 2 0 0 0 2 2m0 10a2 2 0 1 0 2 2a2 2 0 0 0-2-2m0-7a2 2 0 1 0 2 2a2 2 0 0 0-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Emoji(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 11a1 1 0 1 0-1-1a1 1 0 0 0 1 1m-6 0a1 1 0 1 0-1-1a1 1 0 0 0 1 1m3-9a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8m4.28-7.12a1 1 0 0 0-1.4-.16A2.78 2.78 0 0 0 14 14h-3.65a2.81 2.81 0 0 0-.88-1.31a1 1 0 0 0-1.36.2a1 1 0 0 0 .14 1.39a1 1 0 0 1 .25.72a1.09 1.09 0 0 1-.25.72a1 1 0 1 0 1.25 1.56a2.89 2.89 0 0 0 .84-1.28H14a2.72 2.72 0 0 0 .89 1.31a1 1 0 0 0 .57.18a1 1 0 0 0 .78-.38a1 1 0 0 0-.14-1.39a1 1 0 0 1-.25-.72a1.09 1.09 0 0 1 .25-.72a1 1 0 0 0 .18-1.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnglishToChinese(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.022 7h1a1.001 1.001 0 0 1 1 1v1a1 1 0 0 0 2 0V8a3.003 3.003 0 0 0-3-3h-1a1 1 0 0 0 0 2m-4 9h-1a1.001 1.001 0 0 1-1-1v-1a1 1 0 0 0-2 0v1a3.003 3.003 0 0 0 3 3h1a1 1 0 0 0 0-2m11-1a1 1 0 0 0 0-2h-3v-.5a1 1 0 0 0-2 0v.5h-3a1 1 0 0 0 0 2h5.184a6.728 6.728 0 0 1-1.225 2.527a6.668 6.668 0 0 1-.63-.983a1 1 0 1 0-1.779.912a8.678 8.678 0 0 0 .96 1.468a6.618 6.618 0 0 1-2.426 1.099a1 1 0 0 0 .427 1.954a8.635 8.635 0 0 0 3.445-1.622a8.724 8.724 0 0 0 3.469 1.622a1 1 0 1 0 .43-1.954a6.725 6.725 0 0 1-2.446-1.09A8.736 8.736 0 0 0 20.244 15Zm-11.97-3.757a1 1 0 0 0 1.94-.486l-1.757-7.03a2.281 2.281 0 0 0-4.426 0l-1.757 7.03a1 1 0 0 0 1.94.486L5.552 9h2.94ZM6.052 7l.698-2.787a.291.291 0 0 1 .544 0L7.991 7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Enter(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 6a1 1 0 0 0-1 1v4a1 1 0 0 1-1 1H7.41l1.3-1.29a1 1 0 0 0-1.42-1.42l-3 3a1 1 0 0 0-.21.33a1 1 0 0 0 0 .76a1 1 0 0 0 .21.33l3 3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L7.41 14H17a3 3 0 0 0 3-3V7a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Envelope(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 4H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3m-.41 2l-5.88 5.88a1 1 0 0 1-1.42 0L5.41 6ZM20 17a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V7.41l5.88 5.88a3 3 0 0 0 4.24 0L20 7.41Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeAdd(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 7h1v1a1 1 0 0 0 2 0V7h1a1 1 0 0 0 0-2h-1V4a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2m4 4a1 1 0 0 0-1 1v6a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V8.41l5.88 5.89a3 3 0 0 0 4.24 0l2.47-2.47a1 1 0 0 0-1.42-1.42l-2.47 2.47a1 1 0 0 1-1.4 0L5.41 7H13a1 1 0 0 0 0-2H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-6a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 4H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3M5 6h14a1 1 0 0 1 1 1l-8 4.88L4 7a1 1 0 0 1 1-1m15 11a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V9.28l7.48 4.57a1 1 0 0 0 1 0L20 9.28Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeBlock(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.5 13.26a5 5 0 1 0-3.5-1.47a5 5 0 0 0 3.5 1.47m2.12-2.88a3 3 0 0 1-3.4.58l4-4a3 3 0 0 1-.6 3.42m-4.24-4.24a3 3 0 0 1 2.12-.88a3 3 0 0 1 1.28.3l-4 4a3 3 0 0 1 .6-3.42m5.12 8.12a1 1 0 0 0-1 1v4a1 1 0 0 1-1 1h-14a1 1 0 0 1-1-1V9.67l5.88 5.88a3 3 0 0 0 2.11.88a3 3 0 0 0 2.16-.91a1 1 0 0 0 0-1.39a1 1 0 0 0-1.43 0a1 1 0 0 1-1.4 0L4.91 8.26H9.5a1 1 0 0 0 0-2h-5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-4a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeBookmark(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 15.26a1 1 0 0 0-1 1v3a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V9.67l5.88 5.89a3 3 0 0 0 2.1.87a3 3 0 0 0 1.43-.36a1 1 0 0 0 .4-1.36a1 1 0 0 0-1.36-.4a1 1 0 0 1-1.15-.17L5.41 8.26H12a1 1 0 0 0 0-2H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-3a1 1 0 0 0-1-1m0-12h-5a1 1 0 0 0-1 1v8a1 1 0 0 0 1.57.82l1.93-1.29l1.91 1.28a1.06 1.06 0 0 0 .59.19a1 1 0 0 0 .41-.09a1 1 0 0 0 .59-.91v-8a1 1 0 0 0-1-1m-1 7.12l-.94-.63a1 1 0 0 0-1.12 0l-.94.63V5.26h3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeCheck(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.29 8.71a1 1 0 0 0 1.42 0l4-4a1 1 0 1 0-1.42-1.42L17 6.59l-1.29-1.3a1 1 0 0 0-1.42 1.42ZM21 8a1 1 0 0 0-1 1v9a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V8.41l5.88 5.89a3 3 0 0 0 2.11.87a3.08 3.08 0 0 0 2.16-.9l1.72-1.72a1 1 0 1 0-1.42-1.42l-1.75 1.75a1 1 0 0 1-1.4 0L5.41 7H11a1 1 0 0 0 0-2H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V9a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeDownload(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.5 14a1 1 0 0 0-1 1v4a1 1 0 0 1-1 1h-14a1 1 0 0 1-1-1V9.41l5.88 5.89a3 3 0 0 0 4.24 0l1.64-1.64a1 1 0 1 0-1.42-1.42l-1.64 1.64a1 1 0 0 1-1.4 0L4.91 8h6.59a1 1 0 0 0 0-2h-7a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-4a1 1 0 0 0-1-1m1.71-6.71a1 1 0 0 0-1.42 0l-1.29 1.3V3a1 1 0 0 0-2 0v5.59l-1.29-1.3a1 1 0 1 0-1.42 1.42l3 3a1 1 0 0 0 .33.21a.94.94 0 0 0 .76 0a1 1 0 0 0 .33-.21l3-3a1 1 0 0 0 0-1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeDownloadAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.29 8.79a1 1 0 0 0 1.42 0l2-2a1 1 0 0 0 0-1.41a1 1 0 0 0-1.42 0l-.29.29V2.92a1 1 0 0 0-2 0v2.75l-.29-.29a1 1 0 0 0-1.42 0a1 1 0 0 0 0 1.41ZM16 11.08H8a3 3 0 0 0-3 3v5a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3v-5a3 3 0 0 0-3-3m-.42 2L12.7 16a1 1 0 0 1-1.4 0l-2.88-2.92Zm1.42 6a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1v-4.59l2.88 2.88a3 3 0 0 0 4.24 0L17 14.49Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeEdit(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 11.51h2.42a1 1 0 0 0 .71-.29l4.58-4.58a1 1 0 0 0 0-1.42L19.29 2.8a1 1 0 0 0-1.42 0l-4.58 4.58a1.05 1.05 0 0 0-.29.71v2.42a1 1 0 0 0 1 1m1-3l3.58-3.58l1 1L16 9.51h-1Zm6 2a1 1 0 0 0-1 1v7a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V8.9l5.88 5.89a3 3 0 0 0 4.27 0a1 1 0 0 0 0-1.4a1 1 0 0 0-1.43 0a1 1 0 0 1-1.4 0l-5.91-5.9H10a1 1 0 0 0 0-2H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-7a1 1 0 0 0-1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeExclamation(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 13.5a1 1 0 0 0-1 1v4a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V8.91l5.88 5.88a3 3 0 0 0 4.24 0l3.59-3.58a1 1 0 0 0-1.42-1.42l-3.58 3.59a1 1 0 0 1-1.42 0L5.41 7.5H17a1 1 0 0 0 0-2H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-4a1 1 0 0 0-1-1m0-11a1 1 0 0 0-1 1v4a1 1 0 0 0 2 0v-4a1 1 0 0 0-1-1m-.2 7a.64.64 0 0 0-.18.06a.76.76 0 0 0-.18.09l-.15.12a1.05 1.05 0 0 0-.29.71a1.23 1.23 0 0 0 0 .19a.6.6 0 0 0 .06.19a.76.76 0 0 0 .09.18a1.58 1.58 0 0 0 .12.15a1 1 0 0 0 1.42 0l.12-.15a.76.76 0 0 0 .09-.18a.6.6 0 0 0 .06-.19a1.23 1.23 0 0 0 0-.19a1 1 0 0 0-1.2-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeHeart(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.5 13a1 1 0 0 0-1 1v4a1 1 0 0 1-1 1h-14a1 1 0 0 1-1-1V8.41l5.88 5.89a3 3 0 0 0 2.11.87a3 3 0 0 0 2.15-.9l.89-.88a1 1 0 0 0-1.4-1.39l-.93.91a1 1 0 0 1-1.4 0L4.91 7H9.5a1 1 0 0 0 0-2h-5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-4a1 1 0 0 0-1-1m1.05-9a3.33 3.33 0 0 0-3.88-.54a3.25 3.25 0 0 0-3.88 5.13L17 11.71a1.05 1.05 0 0 0 .71.29a1 1 0 0 0 .71-.29l3.17-3.17A3.26 3.26 0 0 0 21.55 4m-1.41 3.12l-2.47 2.47l-2.46-2.47A1.24 1.24 0 0 1 17 5.36a1 1 0 0 0 1.42 0a1.28 1.28 0 0 1 1.76 0a1.26 1.26 0 0 1-.04 1.76"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeInfo(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 13.5a1 1 0 0 0-1 1v4a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V8.91l5.88 5.88a3 3 0 0 0 4.24 0l3.59-3.58a1 1 0 0 0-1.42-1.42l-3.58 3.59a1 1 0 0 1-1.42 0L5.41 7.5H17a1 1 0 0 0 0-2H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-4a1 1 0 0 0-1-1m0-9a1.05 1.05 0 0 0 .71-.29l.12-.16a.56.56 0 0 0 .09-.17a.64.64 0 0 0 .08-.18a1.36 1.36 0 0 0 0-.2a1.05 1.05 0 0 0-.29-.71a1 1 0 0 0-1.42 0a1 1 0 0 0 0 1.42a1.05 1.05 0 0 0 .71.29m0 1a1 1 0 0 0-1 1v4a1 1 0 0 0 2 0v-4a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeLock(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.5 7.44V6.26a3 3 0 1 0-6 0v1.18a3 3 0 0 0-2 2.82v2a3 3 0 0 0 3 3h4a3 3 0 0 0 3-3v-2a3 3 0 0 0-2-2.82m-4-1.18a1 1 0 1 1 2 0v1h-2Zm4 6a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1Zm0 5a1 1 0 0 0-1 1v1a1 1 0 0 1-1 1h-14a1 1 0 0 1-1-1V9.67l5.88 5.88a1 1 0 0 0 1.42-1.41L4.91 8.26h5.59a1 1 0 0 0 0-2h-6a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeMinus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 8a1 1 0 0 0-1 1v8a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V7.41l5.88 5.88a3 3 0 0 0 4.24 0l3.59-3.58a1 1 0 0 0-1.42-1.42l-3.58 3.59a1 1 0 0 1-1.42 0L5.41 6H13a1 1 0 0 0 0-2H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V9a1 1 0 0 0-1-1m-4-2h4a1 1 0 0 0 0-2h-4a1 1 0 0 0 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeOpen(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.21 8.82L14 2.78a2.83 2.83 0 0 0-3.9 0l-6.21 6A2.6 2.6 0 0 0 3 10.71v8.58A2.75 2.75 0 0 0 5.78 22h12.44A2.75 2.75 0 0 0 21 19.29v-8.58a2.67 2.67 0 0 0-.79-1.89m-8.77-4.6a.83.83 0 0 1 1.12 0L18 9.5l-5.47 5.28a.83.83 0 0 1-1.12 0L6 9.5ZM19 19.29a.76.76 0 0 1-.78.71H5.78a.76.76 0 0 1-.78-.71v-7.94l4.05 3.9l-1.66 1.6a1 1 0 0 0 0 1.41a1 1 0 0 0 .72.31a1 1 0 0 0 .69-.28l1.77-1.7a2.8 2.8 0 0 0 2.92 0l1.77 1.7a1 1 0 0 0 .69.28a1 1 0 0 0 .72-.31a1 1 0 0 0 0-1.41L15 15.25l4-3.9Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeQuestion(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.77 5.37A1 1 0 0 0 18.13 5a1 1 0 0 1 .87-.5a1 1 0 0 1 0 2a1 1 0 0 0 0 2A3 3 0 1 0 16.4 4a1 1 0 0 0 .37 1.37M21 13.5a1 1 0 0 0-1 1v4a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V8.91l5.88 5.89a3 3 0 0 0 4.24 0l1.64-1.64a1 1 0 1 0-1.42-1.42l-1.64 1.64a1 1 0 0 1-1.4 0L5.41 7.5H13a1 1 0 0 0 0-2H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-4a1 1 0 0 0-1-1m-2.71-3.71a1 1 0 0 0 0 1.42l.15.12a.76.76 0 0 0 .18.09a.64.64 0 0 0 .18.06h.2a1 1 0 0 0 .71-1.71a1 1 0 0 0-1.42.02"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeReceive(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8.29 6.21l2 2a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-.3-.29H15a1 1 0 0 0 0-2h-3.59l.3-.29a1 1 0 1 0-1.42-1.42l-2 2a1 1 0 0 0-.21.33a1 1 0 0 0 0 .76a1 1 0 0 0 .21.33M16 10.5H8a3 3 0 0 0-3 3v5a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3v-5a3 3 0 0 0-3-3m-.42 2l-2.88 2.88a1 1 0 0 1-1.4 0L8.42 12.5Zm1.42 6a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1v-4.59l2.88 2.87a2.94 2.94 0 0 0 2.12.89a3 3 0 0 0 2.12-.88L17 13.91Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeRedo(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.5 14.26a1 1 0 0 0-1 1v4a1 1 0 0 1-1 1h-14a1 1 0 0 1-1-1V9.67l5.88 5.88a3 3 0 0 0 2.11.88a3 3 0 0 0 2.16-.91a1 1 0 0 0 0-1.39a1 1 0 0 0-1.43 0a1 1 0 0 1-1.4 0L4.91 8.26H9.5a1 1 0 0 0 0-2h-5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-4a1 1 0 0 0-1-1m1-11a1 1 0 0 0-1 1a5 5 0 1 0-3 9A5 5 0 0 0 20.8 12a1 1 0 0 0-1.32-1.51a3 3 0 1 1 .25-4.24H18.5a1 1 0 0 0 0 2h3a1 1 0 0 0 1-1v-3a1 1 0 0 0-1-.99"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeSearch(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m22.21 10.29l-1.73-1.72a4.37 4.37 0 0 0 .65-2.26a4.31 4.31 0 1 0-4.32 4.32a4.37 4.37 0 0 0 2.26-.63l1.72 1.73a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.44M18.45 8a2.37 2.37 0 0 1-3.27 0a2.3 2.3 0 0 1-.68-1.64A2.32 2.32 0 0 1 16.81 4a2.3 2.3 0 0 1 1.64.68a2.28 2.28 0 0 1 .68 1.63A2.3 2.3 0 0 1 18.45 8m2.05 6a1 1 0 0 0-1 1v4a1 1 0 0 1-1 1h-14a1 1 0 0 1-1-1V9.41l5.88 5.89a3 3 0 0 0 4.24 0L15 13.88a1 1 0 0 0-1.42-1.42l-1.38 1.42a1 1 0 0 1-1.4 0L4.91 8H9.5a1 1 0 0 0 0-2h-5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-4a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeSend(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 10.5H8a3 3 0 0 0-3 3v5a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3v-5a3 3 0 0 0-3-3m-.42 2l-2.88 2.88a1 1 0 0 1-1.4 0L8.42 12.5Zm1.42 6a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1v-4.59l2.88 2.87a2.94 2.94 0 0 0 2.12.89a3 3 0 0 0 2.12-.88L17 13.91Zm-8-12h3.59l-.3.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l2-2a1 1 0 0 0 .21-.33a1 1 0 0 0 0-.76a1 1 0 0 0-.21-.33l-2-2a1 1 0 0 0-1.42 1.42l.3.29H9a1 1 0 0 0 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeShare(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 14a1 1 0 0 0-1 1v4a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V9.41l5.88 5.89a3 3 0 0 0 2.11.87a3.08 3.08 0 0 0 2.16-.9l1.72-1.72a1 1 0 1 0-1.42-1.42l-1.75 1.75a1 1 0 0 1-1.4 0L4.41 8H10a1 1 0 0 0 0-2H4a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-4a1 1 0 0 0-1-1m0-6a2 2 0 0 0-1.18.39l-1.75-.8L19 6.71A2 2 0 0 0 20 7a2 2 0 1 0-2-2l-1.9.87A2 2 0 0 0 15 5.5a2 2 0 0 0 0 4a1.88 1.88 0 0 0 .92-.24l2.1 1A2 2 0 1 0 20 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeShield(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.63 5.12a1 1 0 0 0-.84-.2a2.74 2.74 0 0 1-2.2-.46a1 1 0 0 0-1.18 0a2.74 2.74 0 0 1-2.2.46A1 1 0 0 0 14 5.9v3.31a4.62 4.62 0 0 0 1.84 3.7l1.57 1.16a1 1 0 0 0 1.18 0l1.57-1.16A4.62 4.62 0 0 0 22 9.21V5.9a1 1 0 0 0-.37-.78M20 9.21a2.61 2.61 0 0 1-1 2.09l-1 .7l-1-.72a2.61 2.61 0 0 1-1-2.09V7a4.67 4.67 0 0 0 2-.54A4.67 4.67 0 0 0 20 7Zm1 6a1 1 0 0 0-1 1v3a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V9.67l5.88 5.88a3 3 0 0 0 2.11.88a3 3 0 0 0 2.15-.9l-.7-.71l-.74-.68a1 1 0 0 1-1.4 0L5.41 8.26H11a1 1 0 0 0 0-2H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-3a1 1 0 0 0-1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeStar(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.64 9.74l-.29 1.73A1.55 1.55 0 0 0 14 13a1.46 1.46 0 0 0 1.58.09l1.42-.81l1.44.79A1.46 1.46 0 0 0 20 13a1.55 1.55 0 0 0 .63-1.51l-.29-1.73l1.2-1.22a1.54 1.54 0 0 0-.85-2.6l-1.62-.24l-.73-1.55a1.5 1.5 0 0 0-2.72 0l-.73 1.55l-1.62.24a1.54 1.54 0 0 0-.85 2.6Zm1.83-2.13a1.51 1.51 0 0 0 1.14-.85l.39-.83l.39.83a1.55 1.55 0 0 0 1.14.86l1 .14l-.73.74a1.57 1.57 0 0 0-.42 1.34l.16 1l-.79-.43a1.48 1.48 0 0 0-1.44 0l-.79.43l.16-1a1.54 1.54 0 0 0-.42-1.33l-.73-.75ZM21 15.26a1 1 0 0 0-1 1v3a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V9.67l5.88 5.88a2.94 2.94 0 0 0 2.1.88h.27a1 1 0 0 0 .91-1.08a1 1 0 0 0-1.09-.91a.94.94 0 0 1-.77-.28l-5.89-5.9H9a1 1 0 0 0 0-2H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-3a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeTimes(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 11a1 1 0 0 0-1 1v6a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V8.41l5.88 5.89a3 3 0 0 0 2.11.87a3.08 3.08 0 0 0 2.16-.9l1.72-1.72a1 1 0 1 0-1.42-1.42l-1.75 1.75a1 1 0 0 1-1.4 0L5.41 7H13a1 1 0 0 0 0-2H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-6a1 1 0 0 0-1-1m-.59-5l1.3-1.29a1 1 0 1 0-1.42-1.42L19 4.59l-1.29-1.3a1 1 0 0 0-1.42 1.42L17.59 6l-1.3 1.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0L19 7.41l1.29 1.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeUpload(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.5 14a1 1 0 0 0-1 1v4a1 1 0 0 1-1 1h-14a1 1 0 0 1-1-1V9.41l5.88 5.89a3 3 0 0 0 4.24 0l1.64-1.64a1 1 0 1 0-1.42-1.42l-1.64 1.64a1 1 0 0 1-1.4 0L4.91 8h6.59a1 1 0 0 0 0-2h-7a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-4a1 1 0 0 0-1-1m1.71-8.71l-3-3a1 1 0 0 0-.33-.21a1 1 0 0 0-.76 0a1 1 0 0 0-.33.21l-3 3a1 1 0 0 0 1.42 1.42l1.29-1.3V11a1 1 0 0 0 2 0V5.41l1.29 1.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EnvelopeUploadAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 11.08H8a3 3 0 0 0-3 3v5a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3v-5a3 3 0 0 0-3-3m-.42 2L12.7 16a1 1 0 0 1-1.4 0l-2.88-2.92Zm1.42 6a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1v-4.59l2.88 2.88a3 3 0 0 0 4.24 0L17 14.49ZM10.71 5.62l.29-.29v2.75a1 1 0 0 0 2 0V5.33l.29.29a1 1 0 1 0 1.42-1.41l-2-2a1 1 0 0 0-1.42 0l-2 2a1 1 0 0 0 0 1.41a1 1 0 0 0 1.42 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Envelopes(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 21.63H5a3 3 0 0 1-3-3v-8a1 1 0 0 0-2 0v8a5 5 0 0 0 5 5h12a1 1 0 0 0 0-2m4-18H7a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-10a3 3 0 0 0-3-3m-.41 2l-5.88 5.88a1 1 0 0 1-1.42 0L7.41 5.63Zm1.41 11a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V7l5.88 5.88a3 3 0 0 0 4.24 0L22 7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EqualCircle(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 13H9a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2m0-4H9a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2m-3-7a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Estate(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20 8l-6-5.26a3 3 0 0 0-4 0L4 8a3 3 0 0 0-1 2.26V19a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-8.75A3 3 0 0 0 20 8m-6 12h-4v-5a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1Zm5-1a1 1 0 0 1-1 1h-2v-5a3 3 0 0 0-3-3h-2a3 3 0 0 0-3 3v5H6a1 1 0 0 1-1-1v-8.75a1 1 0 0 1 .34-.75l6-5.25a1 1 0 0 1 1.32 0l6 5.25a1 1 0 0 1 .34.75Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Euro(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.2 17.41A6 6 0 0 1 14.46 20c-2.68 0-5-2-6-5H14a1 1 0 0 0 0-2H8.05c0-.33-.05-.67-.05-1s0-.67.05-1H14a1 1 0 0 0 0-2H8.47c1-3 3.31-5 6-5a6 6 0 0 1 4.73 2.59a1 1 0 1 0 1.6-1.18A7.92 7.92 0 0 0 14.46 2c-3.76 0-7 2.84-8.07 7H4a1 1 0 0 0 0 2h2.05v2H4a1 1 0 0 0 0 2h2.39c1.09 4.16 4.31 7 8.07 7a7.92 7.92 0 0 0 6.34-3.41a1 1 0 0 0-1.6-1.18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EuroCircle(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1a11 11 0 1 0 11 11A11 11 0 0 0 12 1m0 20a9 9 0 1 1 9-9a9 9 0 0 1-9 9m.59-13.33a3.34 3.34 0 0 1 2.62 1.38a1 1 0 0 0 1.4.19a1 1 0 0 0 .18-1.41a5.32 5.32 0 0 0-4.2-2.16A5.57 5.57 0 0 0 7.46 9.5H6a1 1 0 0 0 0 2h1v1H6a1 1 0 0 0 0 2h1.46a5.57 5.57 0 0 0 5.13 3.83a5.32 5.32 0 0 0 4.2-2.16A1 1 0 1 0 15.21 15a3.34 3.34 0 0 1-2.62 1.38a3.42 3.42 0 0 1-2.92-1.88H12a1 1 0 0 0 0-2H9.05A4.23 4.23 0 0 1 9 12a4.23 4.23 0 0 1 .05-.5H12a1 1 0 0 0 0-2H9.67a3.42 3.42 0 0 1 2.92-1.83"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Exchange(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 10a1 1 0 0 0-1-1H5.41l2.3-2.29a1 1 0 0 0-1.42-1.42l-4 4a1 1 0 0 0-.21 1.09A1 1 0 0 0 3 11h14a1 1 0 0 0 1-1m3.92 3.62A1 1 0 0 0 21 13H7a1 1 0 0 0 0 2h11.59l-2.3 2.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l4-4a1 1 0 0 0 .21-1.09"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExchangeAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.71 9.29l-4-4a1 1 0 0 0-1.42 1.42L18.59 9H7a1 1 0 0 0 0 2h14a1 1 0 0 0 .92-.62a1 1 0 0 0-.21-1.09M17 13H3a1 1 0 0 0-.92.62a1 1 0 0 0 .21 1.09l4 4a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L5.41 15H17a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Exclamation(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 14a1 1 0 0 0 1-1V7a1 1 0 0 0-2 0v6a1 1 0 0 0 1 1m0 4a1.25 1.25 0 1 0-1.25-1.25A1.25 1.25 0 0 0 12 18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationCircle(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 14a1.25 1.25 0 1 0 1.25 1.25A1.25 1.25 0 0 0 12 14m0-1.5a1 1 0 0 0 1-1v-3a1 1 0 0 0-2 0v3a1 1 0 0 0 1 1M12 2a10 10 0 1 0 10 10A10.011 10.011 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8.01 8.01 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationOctagon(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 7a1 1 0 0 0-1 1v4a1 1 0 0 0 2 0V8a1 1 0 0 0-1-1m0 8a1 1 0 1 0 1 1a1 1 0 0 0-1-1m9.71-7.44l-5.27-5.27a1.05 1.05 0 0 0-.71-.29H8.27a1.05 1.05 0 0 0-.71.29L2.29 7.56a1.05 1.05 0 0 0-.29.71v7.46a1.05 1.05 0 0 0 .29.71l5.27 5.27a1.05 1.05 0 0 0 .71.29h7.46a1.05 1.05 0 0 0 .71-.29l5.27-5.27a1.05 1.05 0 0 0 .29-.71V8.27a1.05 1.05 0 0 0-.29-.71M20 15.31L15.31 20H8.69L4 15.31V8.69L8.69 4h6.62L20 8.69Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationTriangle(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 16a1 1 0 1 0 1 1a1 1 0 0 0-1-1m10.67 1.47l-8.05-14a3 3 0 0 0-5.24 0l-8 14A3 3 0 0 0 3.94 22h16.12a3 3 0 0 0 2.61-4.53m-1.73 2a1 1 0 0 1-.88.51H3.94a1 1 0 0 1-.88-.51a1 1 0 0 1 0-1l8-14a1 1 0 0 1 1.78 0l8.05 14a1 1 0 0 1 .05 1.02ZM12 8a1 1 0 0 0-1 1v4a1 1 0 0 0 2 0V9a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Exclude(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.54 7.54h-1a1 1 0 0 0-1 1v1a1 1 0 0 0 2 0a1 1 0 1 0 0-2m5.92 5.92a1 1 0 0 0-1 1a1 1 0 0 0 0 2h1a1 1 0 0 0 1-1v-1a1 1 0 0 0-1-1M21 7.54h-4.54V3a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v12.46a1 1 0 0 0 1 1h4.54V21a1 1 0 0 0 1 1H21a1 1 0 0 0 1-1V8.54a1 1 0 0 0-1-1M20 20H9.54v-3.54a1 1 0 0 0 0-2a1 1 0 0 0-2 0H4V4h10.46v3.54a1 1 0 0 0 0 2a1 1 0 0 0 2 0H20Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Exit(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 12a1 1 0 0 0 1 1h7.59l-2.3 2.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l4-4a1 1 0 0 0 .21-.33a1 1 0 0 0 0-.76a1 1 0 0 0-.21-.33l-4-4a1 1 0 1 0-1.42 1.42l2.3 2.29H5a1 1 0 0 0-1 1M17 2H7a3 3 0 0 0-3 3v3a1 1 0 0 0 2 0V5a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1v-3a1 1 0 0 0-2 0v3a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExpandAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.79 12.79L4 18.59V17a1 1 0 0 0-2 0v4a1 1 0 0 0 .08.38a1 1 0 0 0 .54.54A1 1 0 0 0 3 22h4a1 1 0 0 0 0-2H5.41l5.8-5.79a1 1 0 0 0-1.42-1.42M21.92 2.62a1 1 0 0 0-.54-.54A1 1 0 0 0 21 2h-4a1 1 0 0 0 0 2h1.59l-5.8 5.79a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0L20 5.41V7a1 1 0 0 0 2 0V3a1 1 0 0 0-.08-.38"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExpandArrows(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 8a1 1 0 0 0 1-1V3a1 1 0 0 0-.08-.38a1 1 0 0 0-.54-.54A1 1 0 0 0 21 2h-4a1 1 0 0 0 0 2h1.59L12 10.59L5.41 4H7a1 1 0 0 0 0-2H3a1 1 0 0 0-.38.08a1 1 0 0 0-.54.54A1 1 0 0 0 2 3v4a1 1 0 0 0 2 0V5.41L10.59 12L4 18.59V17a1 1 0 0 0-2 0v4a1 1 0 0 0 .08.38a1 1 0 0 0 .54.54A1 1 0 0 0 3 22h4a1 1 0 0 0 0-2H5.41L12 13.41L18.59 20H17a1 1 0 0 0 0 2h4a1 1 0 0 0 .38-.08a1 1 0 0 0 .54-.54A1 1 0 0 0 22 21v-4a1 1 0 0 0-2 0v1.59L13.41 12L20 5.41V7a1 1 0 0 0 1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExpandArrowsAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.29 13.29L4 18.59V17a1 1 0 0 0-2 0v4a1 1 0 0 0 .08.38a1 1 0 0 0 .54.54A1 1 0 0 0 3 22h4a1 1 0 0 0 0-2H5.41l5.3-5.29a1 1 0 0 0-1.42-1.42M5.41 4H7a1 1 0 0 0 0-2H3a1 1 0 0 0-.38.08a1 1 0 0 0-.54.54A1 1 0 0 0 2 3v4a1 1 0 0 0 2 0V5.41l5.29 5.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM21 16a1 1 0 0 0-1 1v1.59l-5.29-5.3a1 1 0 0 0-1.42 1.42l5.3 5.29H17a1 1 0 0 0 0 2h4a1 1 0 0 0 .38-.08a1 1 0 0 0 .54-.54A1 1 0 0 0 22 21v-4a1 1 0 0 0-1-1m.92-13.38a1 1 0 0 0-.54-.54A1 1 0 0 0 21 2h-4a1 1 0 0 0 0 2h1.59l-5.3 5.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0L20 5.41V7a1 1 0 0 0 2 0V3a1 1 0 0 0-.08-.38"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExpandFromCorner(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 12H3a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1m-1 8H4v-6h6ZM21.92 2.62a1 1 0 0 0-.54-.54A1 1 0 0 0 21 2h-6a1 1 0 0 0 0 2h3.59l-5.3 5.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0L20 5.41V9a1 1 0 0 0 2 0V3a1 1 0 0 0-.08-.38"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExpandLeft(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.17 10.17a1 1 0 0 0 .71-.29a1 1 0 0 0 0-1.42L5.41 4H7a1 1 0 0 0 0-2H3a1 1 0 0 0-.38.08a1 1 0 0 0-.54.54A1 1 0 0 0 2 3v4a1 1 0 0 0 2 0V5.41l4.46 4.47a1 1 0 0 0 .71.29m6.37-1.71a1 1 0 0 0-1.42 0l-5.66 5.66a1 1 0 0 0 0 1.42a1 1 0 0 0 .71.29a1 1 0 0 0 .71-.29l5.66-5.66a1 1 0 0 0 0-1.42M21 16a1 1 0 0 0-1 1v1.59l-4.46-4.47a1 1 0 1 0-1.42 1.42L18.59 20H17a1 1 0 0 0 0 2h4a1 1 0 0 0 .38-.08a1 1 0 0 0 .54-.54A1 1 0 0 0 22 21v-4a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExpandRight(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 14.1L4 18.6V17c0-.6-.4-1-1-1s-1 .4-1 1v4c0 .1 0 .3.1.4c.1.2.3.4.5.5c.1.1.3.1.4.1h4c.6 0 1-.4 1-1s-.4-1-1-1H5.4l4.5-4.5c.4-.4.4-1 0-1.4c-.4-.4-1-.4-1.4 0M21.7 2.3c-.2-.2-.5-.3-.7-.3h-4c-.6 0-1 .4-1 1s.4 1 1 1h1.6l-4.5 4.5c-.4.4-.4 1 0 1.4c.2.2.4.3.7.3c.3 0 .5-.1.7-.3L20 5.4V7c0 .6.4 1 1 1s1-.4 1-1V3c0-.2-.1-.5-.3-.7m-6.2 11.8L9.9 8.5c-.4-.4-1-.4-1.4 0c-.4.4-.4 1 0 1.4l5.7 5.7c.2.2.4.3.7.3c.3 0 .5-.1.7-.3c.3-.4.3-1.1-.1-1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Export(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.71 7.71L11 5.41V15a1 1 0 0 0 2 0V5.41l2.29 2.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-4-4a1 1 0 0 0-.33-.21a1 1 0 0 0-.76 0a1 1 0 0 0-.33.21l-4 4a1 1 0 1 0 1.42 1.42M21 14a1 1 0 0 0-1 1v4a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-4a1 1 0 0 0-2 0v4a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-4a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExposureAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 6H7a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2m8-4H5a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3M4 18.59V5a1 1 0 0 1 1-1h13.59ZM20 19a1 1 0 0 1-1 1H5.41L20 5.41Zm-7-2h1v1a1 1 0 0 0 2 0v-1h1a1 1 0 0 0 0-2h-1v-1a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExposureIncrease(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 7h-1V6a1 1 0 0 0-2 0v1H7a1 1 0 0 0 0 2h1v1a1 1 0 0 0 2 0V9h1a1 1 0 0 0 0-2m2 11h4a1 1 0 0 0 0-2h-4a1 1 0 0 0 0 2m6-16H5a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3M4 18.59V5a1 1 0 0 1 1-1h13.59ZM20 19a1 1 0 0 1-1 1H5.41L20 5.41Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExternalLinkAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 10.82a1 1 0 0 0-1 1V19a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1h7.18a1 1 0 0 0 0-2H5a3 3 0 0 0-3 3v11a3 3 0 0 0 3 3h11a3 3 0 0 0 3-3v-7.18a1 1 0 0 0-1-1m3.92-8.2a1 1 0 0 0-.54-.54A1 1 0 0 0 21 2h-6a1 1 0 0 0 0 2h3.59L8.29 14.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0L20 5.41V9a1 1 0 0 0 2 0V3a1 1 0 0 0-.08-.38"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eye(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.92 11.6C19.9 6.91 16.1 4 12 4s-7.9 2.91-9.92 7.6a1 1 0 0 0 0 .8C4.1 17.09 7.9 20 12 20s7.9-2.91 9.92-7.6a1 1 0 0 0 0-.8M12 18c-3.17 0-6.17-2.29-7.9-6C5.83 8.29 8.83 6 12 6s6.17 2.29 7.9 6c-1.73 3.71-4.73 6-7.9 6m0-10a4 4 0 1 0 4 4a4 4 0 0 0-4-4m0 6a2 2 0 1 1 2-2a2 2 0 0 1-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.94 6.08A6.93 6.93 0 0 1 12 6c3.18 0 6.17 2.29 7.91 6a15.23 15.23 0 0 1-.9 1.64a1 1 0 0 0-.16.55a1 1 0 0 0 1.86.5a15.77 15.77 0 0 0 1.21-2.3a1 1 0 0 0 0-.79C19.9 6.91 16.1 4 12 4a7.77 7.77 0 0 0-1.4.12a1 1 0 1 0 .34 2ZM3.71 2.29a1 1 0 0 0-1.42 1.42l3.1 3.09a14.62 14.62 0 0 0-3.31 4.8a1 1 0 0 0 0 .8C4.1 17.09 7.9 20 12 20a9.26 9.26 0 0 0 5.05-1.54l3.24 3.25a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Zm6.36 9.19l2.45 2.45A1.81 1.81 0 0 1 12 14a2 2 0 0 1-2-2a1.81 1.81 0 0 1 .07-.52M12 18c-3.18 0-6.17-2.29-7.9-6a12.09 12.09 0 0 1 2.7-3.79L8.57 10A4 4 0 0 0 14 15.43L15.59 17A7.24 7.24 0 0 1 12 18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Facebook(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.9 2H3.1A1.1 1.1 0 0 0 2 3.1v17.8A1.1 1.1 0 0 0 3.1 22h9.58v-7.75h-2.6v-3h2.6V9a3.64 3.64 0 0 1 3.88-4a20.26 20.26 0 0 1 2.33.12v2.7H17.3c-1.26 0-1.5.6-1.5 1.47v1.93h3l-.39 3H15.8V22h5.1a1.1 1.1 0 0 0 1.1-1.1V3.1A1.1 1.1 0 0 0 20.9 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FacebookF(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.12 5.32H17V2.14A26.11 26.11 0 0 0 14.26 2c-2.72 0-4.58 1.66-4.58 4.7v2.62H6.61v3.56h3.07V22h3.68v-9.12h3.06l.46-3.56h-3.52V7.05c0-1.05.28-1.73 1.76-1.73"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FacebookMessenger(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a9.65 9.65 0 0 0-10 9.7a9.51 9.51 0 0 0 3.14 7.18a.81.81 0 0 1 .27.56v1.78a.81.81 0 0 0 1.13.71l2-.87a.75.75 0 0 1 .53 0a11 11 0 0 0 2.9.38A9.7 9.7 0 1 0 12 2m6 7.46l-2.93 4.66a1.5 1.5 0 0 1-2.17.4l-2.34-1.75a.6.6 0 0 0-.72 0l-3.16 2.4a.47.47 0 0 1-.68-.63l2.93-4.66a1.5 1.5 0 0 1 2.17-.4l2.34 1.75a.6.6 0 0 0 .72 0l3.16-2.4a.47.47 0 0 1 .68.63"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FacebookMessengerAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.991 1a10.614 10.614 0 0 0-11 10.7a10.461 10.461 0 0 0 3.414 7.866l.052 1.69A1.8 1.8 0 0 0 6.256 23a1.82 1.82 0 0 0 .726-.152L8.903 22a11.895 11.895 0 0 0 3.088.4a10.615 10.615 0 0 0 11.001-10.7a10.615 10.615 0 0 0-11-10.7m0 19.4a9.862 9.862 0 0 1-2.635-.35a1.799 1.799 0 0 0-1.196.092l-1.714.756l-.045-1.493A1.81 1.81 0 0 0 5.8 18.13a8.488 8.488 0 0 1-2.81-6.43a8.66 8.66 0 0 1 9-8.7a8.705 8.705 0 1 1 0 17.4m3.735-11.815l-2.313 2.755l-3.347-2.056a.998.998 0 0 0-1.289.21l-3.05 3.636a1 1 0 1 0 1.53 1.285l2.499-2.975l3.347 2.056a.998.998 0 0 0 1.289-.21l2.866-3.415a1 1 0 1 0-1.531-1.286"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fahrenheit(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.5 3h-7a3 3 0 0 0-3 3v14a1 1 0 0 0 2 0v-7h6a1 1 0 0 0 0-2h-6V6a1 1 0 0 1 1-1h7a1 1 0 0 0 0-2m-15 0a3 3 0 1 0 3 3a3 3 0 0 0-3-3m0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastMail(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.79 10.5h-2a1 1 0 1 0 0 2h2a1 1 0 0 0 0-2m16.78-2.84V7.6a3 3 0 0 0-2.37-1.1h-7.93a3 3 0 0 0-2 .74A2.93 2.93 0 0 0 8.31 9l-.88 5a3 3 0 0 0 .66 2.45a3 3 0 0 0 2.29 1.07h7.94a3 3 0 0 0 3-2.48l.88-5a3 3 0 0 0-.63-2.38m-2.74.84l-3.4 2.76a1 1 0 0 1-1.38-.12L11.72 8.5Zm.48 6.17a1 1 0 0 1-1 .83h-7.93a1 1 0 0 1-.76-.36a1 1 0 0 1-.22-.81l.8-4.53l2.35 2.66a3 3 0 0 0 4.14.35L20.13 10ZM5.79 6.5h-3a1 1 0 1 0 0 2h3a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastMailAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.69 9a2.93 2.93 0 0 0-1-1.71a3 3 0 0 0-2-.74H4.8a3 3 0 0 0-2.3 1.02v.06A3 3 0 0 0 1.84 10l.88 5a3 3 0 0 0 3 2.48h7.94a3 3 0 0 0 2.29-1.07a3 3 0 0 0 .62-2.41Zm-3.41-.5l-2.34 2.64a1 1 0 0 1-1.38.11L5.17 8.5Zm2.1 6.64a1 1 0 0 1-.76.36H5.68a1 1 0 0 1-1-.83L3.87 10l3.43 2.8a3 3 0 0 0 4.14-.34L13.8 9.8l.8 4.53a1 1 0 0 1-.22.81m6.83-4.64h-2a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2m0-2a1 1 0 0 0 0-2h-3a1 1 0 0 0 0 2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Favorite(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 9.67a1 1 0 0 0-.86-.67l-5.69-.83L12.9 3a1 1 0 0 0-1.8 0L8.55 8.16L2.86 9a1 1 0 0 0-.81.68a1 1 0 0 0 .25 1l4.13 4l-1 5.68a1 1 0 0 0 1.47 1.08l5.1-2.67l5.1 2.67a.93.93 0 0 0 .46.12a1 1 0 0 0 .59-.19a1 1 0 0 0 .4-1l-1-5.68l4.13-4A1 1 0 0 0 22 9.67m-6.15 4a1 1 0 0 0-.29.88l.72 4.2l-3.76-2a1.06 1.06 0 0 0-.94 0l-3.76 2l.72-4.2a1 1 0 0 0-.29-.88l-3-3l4.21-.61a1 1 0 0 0 .76-.55L12 5.7l1.88 3.82a1 1 0 0 0 .76.55l4.21.61Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Feedback(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 1h-7a2.44 2.44 0 0 0-2.41 2l-.92 5.05a2.44 2.44 0 0 0 .53 2a2.47 2.47 0 0 0 1.88.88H17l-.25.66a3.26 3.26 0 0 0 3 4.41a1 1 0 0 0 .92-.59l2.24-5.06A1 1 0 0 0 23 10V2a1 1 0 0 0-1-1m-1 8.73l-1.83 4.13a1.33 1.33 0 0 1-.45-.4a1.23 1.23 0 0 1-.14-1.16l.38-1a1.68 1.68 0 0 0-.2-1.58A1.7 1.7 0 0 0 17.35 9h-3.29a.46.46 0 0 1-.35-.16a.5.5 0 0 1-.09-.37l.92-5A.44.44 0 0 1 15 3h6ZM9.94 13.05H7.05l.25-.66A3.26 3.26 0 0 0 4.25 8a1 1 0 0 0-.92.59l-2.24 5.06a1 1 0 0 0-.09.4v8a1 1 0 0 0 1 1h7a2.44 2.44 0 0 0 2.41-2l.92-5a2.44 2.44 0 0 0-.53-2a2.47 2.47 0 0 0-1.86-1m-.48 7.58A.44.44 0 0 1 9 21H3v-6.73l1.83-4.13a1.33 1.33 0 0 1 .45.4a1.23 1.23 0 0 1 .14 1.16l-.38 1a1.68 1.68 0 0 0 .2 1.58a1.7 1.7 0 0 0 1.41.74h3.29a.46.46 0 0 1 .35.16a.5.5 0 0 1 .09.37Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FidgetSpinner(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 7a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-5.696 9.134a1 1 0 1 0 1.366.366a1 1 0 0 0-1.366-.366m11.392 0a1 1 0 1 0 .366 1.366a1 1 0 0 0-.366-1.366m2.914-2.791a4.918 4.918 0 0 0-4.526-1.197l-.419-.763a4.989 4.989 0 0 0-2.503-8.251a5.035 5.035 0 0 0-4.279.958A4.978 4.978 0 0 0 7 8a4.929 4.929 0 0 0 1.352 3.392l-.419.75a4.989 4.989 0 0 0-5.926 6.286a5.03 5.03 0 0 0 2.97 3.226a4.97 4.97 0 0 0 6.588-3.19l.867.014a4.981 4.981 0 0 0 4.76 3.524a5.017 5.017 0 0 0 4.8-3.573a4.95 4.95 0 0 0-1.382-5.086m-.528 4.495a3.006 3.006 0 0 1-4.386 1.76a2.965 2.965 0 0 1-1.352-1.705a1.994 1.994 0 0 0-1.91-1.43h-.869a1.995 1.995 0 0 0-1.91 1.43a2.98 2.98 0 0 1-3.948 1.899a2.993 2.993 0 0 1 1.767-5.704a1.967 1.967 0 0 0 2.173-.942l.436-.754a1.995 1.995 0 0 0-.281-2.369a2.98 2.98 0 0 1 .329-4.37a2.993 2.993 0 0 1 4.069 4.369a2 2 0 0 0-.283 2.37l.435.753a1.974 1.974 0 0 0 2.174.943a2.988 2.988 0 0 1 3.556 3.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func File(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 8.94a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19h-.09L13.06 2H7a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V9zm-6-3.53L16.59 8H14ZM18 19a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v5a1 1 0 0 0 1 1h5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 10h1a1 1 0 0 0 0-2H9a1 1 0 0 0 0 2m0 2a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2Zm11-3.06a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19a.32.32 0 0 0-.09 0a.88.88 0 0 0-.33-.11H7a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V9zm-6-3.53L16.59 8H15a1 1 0 0 1-1-1ZM18 19a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h3Zm-3-3H9a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileBlank(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 8.94a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19h-.09a.88.88 0 0 0-.33-.11H7a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V9zm-6-3.53L16.59 8H15a1 1 0 0 1-1-1ZM18 19a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileBlockAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 8a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2Zm5 12H5a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h3v1a1 1 0 0 0 2 0V8.94a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19a.32.32 0 0 0-.09 0L11.06 2H5a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h7a1 1 0 0 0 0-2m0-14.59L14.59 8H13a1 1 0 0 1-1-1Zm8.83 9.76a4.1 4.1 0 0 0-5.66 0a4 4 0 1 0 5.66 0M16 18a2 2 0 0 1 2-2a2.09 2.09 0 0 1 .51.07l-2.44 2.44A2.09 2.09 0 0 1 16 18m3.41 1.41a2 2 0 0 1-1.91.5l2.43-2.42A2.09 2.09 0 0 1 20 18a2 2 0 0 1-.59 1.41M11 18a1 1 0 0 0 0-2H7a1 1 0 0 0 0 2Zm2-6H7a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileBookmarkAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 10h1a1 1 0 0 0 0-2h-1a1 1 0 0 0 0 2m4 6h-4a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2m0-4h-4a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2m6.92-2.62a1 1 0 0 0-.21-1.09l-6-6a1.07 1.07 0 0 0-.28-.19a.32.32 0 0 0-.09 0l-.28-.1H5.5a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h6a1 1 0 0 0 0-2h-6a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h4a1 1 0 0 0 .92-.62M13.5 8a1 1 0 0 1-1-1V5.41L15.09 8Zm7 4h-5a1 1 0 0 0-1 1v8a1 1 0 0 0 .53.88a1 1 0 0 0 1-.05l1.97-1.3l2 1.3a1 1 0 0 0 1.5-.83v-8a1 1 0 0 0-1-1m-1 7.12l-.94-.63a1 1 0 0 0-1.12 0l-.94.63V14h3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileCheck(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 8.94a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19a.32.32 0 0 0-.09 0a.88.88 0 0 0-.33-.11H7a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V9zm-6-3.53L16.59 8H15a1 1 0 0 1-1-1ZM18 19a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h3Zm-3.71-6.71L11 15.59l-1.29-1.3a1 1 0 0 0-1.42 1.42l2 2a1 1 0 0 0 1.42 0l4-4a1 1 0 0 0-1.42-1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileCheckAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.5 20h-6a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h3v5a1 1 0 0 0 2 0V8.94a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19a.29.29 0 0 0-.1 0a1.1 1.1 0 0 0-.31-.11H5.5a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h6a1 1 0 0 0 0-2m1-14.59L15.09 8H13.5a1 1 0 0 1-1-1ZM7.5 14h6a1 1 0 0 0 0-2h-6a1 1 0 0 0 0 2m4 2h-4a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2m-4-6h1a1 1 0 0 0 0-2h-1a1 1 0 0 0 0 2m13.71 6.29a1 1 0 0 0-1.42 0l-3.29 3.3l-1.29-1.3a1 1 0 0 0-1.42 1.42l2 2a1 1 0 0 0 1.42 0l4-4a1 1 0 0 0 0-1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileContract(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.2 15c-.2.1-.4.1-.6.2c-.1-.2-.2-.4-.3-.5c-.8-.7-2-.8-3-.2C8.5 15 8 16 8 17c0 .5.5 1 1 .9c.5 0 1-.5 1-1c0-.3.1-.5.3-.7c.1 0 .2-.1.3-.1c-.3.6-.1 1.3.5 1.7c.2.1.3.1.5.1c.4 0 .8-.2 1.1-.5c.1-.1.3-.2.5-.3c.1.4.5.8 1 .8h.8c.6 0 1-.4 1-1c0-.5-.4-.9-.9-1c-.1-.2-.1-.3-.3-.5c-.3-.3-1-.5-1.6-.4M20 8.9c0-.1 0-.2-.1-.3v-.1c0-.1-.1-.2-.2-.3l-6-6c-.1-.1-.2-.1-.3-.2h-.1c-.1 0-.2 0-.3-.1H7C5.3 2 4 3.3 4 5v14c0 1.7 1.3 3 3 3h10c1.7 0 3-1.3 3-3zc0 .1 0 .1 0 0m-6-3.5L16.6 8H15c-.6 0-1-.4-1-1zM18 19c0 .6-.4 1-1 1H7c-.6 0-1-.4-1-1V5c0-.6.4-1 1-1h5v3c0 1.7 1.3 3 3 3h3zm-9-9h1c.6 0 1-.4 1-1s-.4-1-1-1H9c-.6 0-1 .4-1 1s.4 1 1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileContractDollar(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.71 20.29L20 18.56v-.31a2.75 2.75 0 0 0-2.75-2.75h-.34l-1.44-1.44a.67.67 0 0 1 .28-.06H19a1 1 0 0 0 0-2h-1.5v-1a1 1 0 0 0-2 0v1a2.74 2.74 0 0 0-1.47.59l-1.32-1.33a1 1 0 0 0-1.42 1.42L13 14.44v.31a2.75 2.75 0 0 0 2.75 2.75h.34l1.44 1.44a.67.67 0 0 1-.28.06H14a1 1 0 0 0 0 2h1.5v1a1 1 0 0 0 2 0v-1a2.74 2.74 0 0 0 1.5-.62l1.32 1.33a1 1 0 0 0 1.42 0a1 1 0 0 0-.03-1.42M10 19H5a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h6v4a1 1 0 0 0 1 1h5a1 1 0 0 0 .92-.62a1 1 0 0 0-.21-1.09l-5-5a1.07 1.07 0 0 0-.28-.19h-.09a1.31 1.31 0 0 0-.28-.1H5a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h5a1 1 0 0 0 0-2m3-14.59L14.59 6H13Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileCopyAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 20H8a3 3 0 0 1-3-3V7a1 1 0 0 0-2 0v10a5 5 0 0 0 5 5h8a1 1 0 0 0 0-2m-6-7a1 1 0 0 0 1 1h5a1 1 0 0 0 0-2h-5a1 1 0 0 0-1 1m11-4.06a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19a.32.32 0 0 0-.09 0a.88.88 0 0 0-.33-.11H10a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V9zm-6-3.53L17.59 8H16a1 1 0 0 1-1-1ZM19 15a1 1 0 0 1-1 1h-8a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h3v3a3 3 0 0 0 .18 1H11a1 1 0 0 0 0 2h8Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileDownload(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 8.94a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19h-.1a1.1 1.1 0 0 0-.31-.11H7a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V9zm-6-3.53L16.59 8H15a1 1 0 0 1-1-1ZM18 19a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h3Zm-4.71-4.71l-.29.3V12a1 1 0 0 0-2 0v2.59l-.29-.3a1 1 0 0 0-1.42 1.42l2 2a1 1 0 0 0 .33.21a.94.94 0 0 0 .76 0a1 1 0 0 0 .33-.21l2-2a1 1 0 0 0-1.42-1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileDownloadAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 8a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2Zm5 12H6a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h3v2a1 1 0 0 0 2 0V8.94a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19a.29.29 0 0 0-.1 0a1.1 1.1 0 0 0-.31-.11H6a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h7a1 1 0 0 0 0-2m0-14.59L15.59 8H14a1 1 0 0 1-1-1ZM14 12H8a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2m6.71 6.29a1 1 0 0 0-1.42 0l-.29.3V16a1 1 0 0 0-2 0v2.59l-.29-.3a1 1 0 0 0-1.42 1.42l2 2a1 1 0 0 0 .33.21a.94.94 0 0 0 .76 0a1 1 0 0 0 .33-.21l2-2a1 1 0 0 0 0-1.42M12 18a1 1 0 0 0 0-2H8a1 1 0 0 0 0 2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileEditAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.71 16.71l-2.42-2.42a1 1 0 0 0-1.42 0l-3.58 3.58a1 1 0 0 0-.29.71V21a1 1 0 0 0 1 1h2.42a1 1 0 0 0 .71-.29l3.58-3.58a1 1 0 0 0 0-1.42M16 20h-1v-1l2.58-2.58l1 1Zm-6 0H6a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h3v1a1 1 0 0 0 2 0V8.94a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19a.32.32 0 0 0-.09 0L12.06 2H6a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h4a1 1 0 0 0 0-2m3-14.59L15.59 8H14a1 1 0 0 1-1-1ZM8 14h6a1 1 0 0 0 0-2H8a1 1 0 0 0 0 2m0-4h1a1 1 0 0 0 0-2H8a1 1 0 0 0 0 2m2 6H8a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileExclamation(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.92 16.62a.76.76 0 0 0-.09-.18a1.58 1.58 0 0 0-.12-.15a1 1 0 0 0-1.09-.21a1 1 0 0 0-.33.21a1.58 1.58 0 0 0-.12.15a.76.76 0 0 0-.09.18a.64.64 0 0 0-.06.18a1.5 1.5 0 0 0 0 .2a.84.84 0 0 0 .08.38a1.15 1.15 0 0 0 .21.33A1 1 0 0 0 12 18a.84.84 0 0 0 .38-.08a.9.9 0 0 0 .54-.54A.84.84 0 0 0 13 17a1.5 1.5 0 0 0 0-.2a.64.64 0 0 0-.08-.18M20 8.94a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19h-.09a.88.88 0 0 0-.33-.11H7a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V9zm-6-3.53L16.59 8H15a1 1 0 0 1-1-1ZM18 19a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h3Zm-6-8a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0v-2a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileExclamationAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 10h1a1 1 0 0 0 0-2H9a1 1 0 0 0 0 2m6 2H9a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2m0 4H9a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2m4.71 4.29a1.58 1.58 0 0 0-.15-.12a.76.76 0 0 0-.18-.09L19.2 20a1 1 0 0 0-.58.06a.9.9 0 0 0-.54.54a.84.84 0 0 0-.08.4a1 1 0 1 0 1.92-.38a1.15 1.15 0 0 0-.21-.33M20 8.94a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19a.32.32 0 0 0-.09 0a.88.88 0 0 0-.33-.11H7a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h8a1 1 0 0 0 0-2H7a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h3v2a1 1 0 0 0 2 0V9zM15 8a1 1 0 0 1-1-1V5.41L16.59 8Zm4 7a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0v-2a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileExport(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.92 15.62a1.15 1.15 0 0 0-.21-.33l-3-3a1 1 0 0 0-1.42 1.42l1.3 1.29H12a1 1 0 0 0 0 2h5.59l-1.3 1.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l3-3a.93.93 0 0 0 .21-.33a1 1 0 0 0 0-.76M14 20H6a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h4a1 1 0 0 0 .92-.62a1 1 0 0 0-.21-1.09l-6-6a1.07 1.07 0 0 0-.28-.19h-.09l-.28-.1H6a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h8a1 1 0 0 0 0-2M13 5.41L15.59 8H14a1 1 0 0 1-1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileGraph(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 10h1a1 1 0 0 0 0-2H9a1 1 0 0 0 0 2m5.21 2.386l-1.673 2.152l-.868-.781a1 1 0 0 0-1.45.118l-2 2.5a1 1 0 1 0 1.562 1.25l1.338-1.673l.879.791a1 1 0 0 0 1.458-.13l2.334-3a1 1 0 0 0-1.58-1.227m5.778-3.448a1.009 1.009 0 0 0-.28-.643l-.001-.002l-6-6l-.001-.001a.99.99 0 0 0-.287-.193c-.03-.014-.06-.022-.092-.033a.983.983 0 0 0-.267-.054C13.04 2.011 13.021 2 13 2H7a3.003 3.003 0 0 0-3 3v14a3.003 3.003 0 0 0 3 3h10a3.003 3.003 0 0 0 3-3V9c0-.022-.011-.04-.012-.062M14 5.414L16.586 8H15a1.001 1.001 0 0 1-1-1ZM18 19a1.001 1.001 0 0 1-1 1H7a1.001 1.001 0 0 1-1-1V5a1.001 1.001 0 0 1 1-1h5v3a3.003 3.003 0 0 0 3 3h3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileHeart(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 8.94a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19h-.1a1.1 1.1 0 0 0-.31-.11H7a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V9zm-6-3.53L16.59 8H15a1 1 0 0 1-1-1ZM18 19a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h3Zm-6-7.66A2.92 2.92 0 0 0 8.57 16l2.72 2.72a1 1 0 0 0 1.42 0L15.43 16A2.92 2.92 0 0 0 12 11.34m2 1.93a.92.92 0 0 1 0 1.3l-2 2l-2-2a.92.92 0 0 1 0-1.3a.92.92 0 0 1 1.3 0a1 1 0 0 0 1.42 0a.92.92 0 0 1 1.28 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileImport(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 20H6a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h3v2a1 1 0 0 0 2 0V8.94a1.31 1.31 0 0 0-.06-.27v-.09a1.32 1.32 0 0 0-.19-.29l-6-6a1.32 1.32 0 0 0-.29-.19a.32.32 0 0 0-.09 0l-.31-.1H6a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h5a1 1 0 0 0 0-2m2-14.59L15.59 8H14a1 1 0 0 1-1-1ZM19 15h-5.59l1.3-1.29a1 1 0 0 0-1.42-1.42l-3 3a1.15 1.15 0 0 0-.21.33a1 1 0 0 0 0 .76a.93.93 0 0 0 .21.33l3 3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L13.41 17H19a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileInfoAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 16H9a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2m-6-6h1a1 1 0 0 0 0-2H9a1 1 0 0 0 0 2m6 2H9a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2m4.71 3.29a1 1 0 0 0-.33-.21a.92.92 0 0 0-.76 0a1 1 0 0 0-.33.21a1.15 1.15 0 0 0-.21.33a1 1 0 0 0 .21 1.09A1 1 0 0 0 19 17a1 1 0 0 0 .38-.08a1.15 1.15 0 0 0 .33-.21a1 1 0 0 0 .21-1.09a1.15 1.15 0 0 0-.21-.33M20 8.94a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19a.32.32 0 0 0-.09 0a.88.88 0 0 0-.33-.11H7a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h8a1 1 0 0 0 0-2H7a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h3v2a1 1 0 0 0 2 0V9zM15 8a1 1 0 0 1-1-1V5.41L16.59 8Zm4 10a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0v-2a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileLandscape(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 10.94a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19h-.1a1.1 1.1 0 0 0-.31-.11H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-6zm-6-3.53L18.59 10H17a1 1 0 0 1-1-1ZM20 17a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1h9v3a3 3 0 0 0 3 3h3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileLandscapeAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 10.94a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19h-.1a1.1 1.1 0 0 0-.31-.11H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-6zm-6-3.53L18.59 10H17a1 1 0 0 1-1-1ZM20 17a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1h9v3a3 3 0 0 0 3 3h3ZM7 12h4a1 1 0 0 0 0-2H7a1 1 0 0 0 0 2m0 2a1 1 0 0 0 0 2h10a1 1 0 0 0 0-2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileLanscapeSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 10.94a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19h-.09a.88.88 0 0 0-.33-.11h-4.39a1 1 0 0 0 0 2H14v3a3 3 0 0 0 3 3h3v3.34a1 1 0 1 0 2 0V11zM17 10a1 1 0 0 1-1-1V7.41L18.59 10ZM3.71 2.29a1 1 0 0 0-1.42 1.42l.91.9A3 3 0 0 0 2 7v10a3 3 0 0 0 3 3h13.59l1.7 1.71a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM5 18a1 1 0 0 1-1-1V7a1 1 0 0 1 .66-.93L16.59 18Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileLockAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 20H5a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h4a1 1 0 0 0 .92-.62a1 1 0 0 0-.21-1.09l-6-6a1.07 1.07 0 0 0-.28-.19a.32.32 0 0 0-.09 0l-.28-.1H5a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h5a1 1 0 0 0 0-2m2-14.59L14.59 8H13a1 1 0 0 1-1-1ZM13 13a1 1 0 0 0-1-1H7a1 1 0 0 0 0 2h5a1 1 0 0 0 1-1m-6-3h1a1 1 0 0 0 0-2H7a1 1 0 0 0 0 2m0 6a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2Zm13-.82V15a3 3 0 0 0-6 0v.18A3 3 0 0 0 12 18v1a3 3 0 0 0 3 3h4a3 3 0 0 0 3-3v-1a3 3 0 0 0-2-2.82M17 14a1 1 0 0 1 1 1h-2a1 1 0 0 1 1-1m3 5a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-1a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileMedical(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 14h-1v-1a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2h1v1a1 1 0 0 0 2 0v-1h1a1 1 0 0 0 0-2m6-5.06a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19h-.09a.88.88 0 0 0-.33-.11H7a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V9zm-6-3.53L16.59 8H15a1 1 0 0 1-1-1ZM18 19a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileMedicalAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 18a1 1 0 0 0-1 1a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h3v1a1 1 0 0 0 2 0V8.94a1.31 1.31 0 0 0-.06-.27a.32.32 0 0 0 0-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19h-.1a1.14 1.14 0 0 0-.3-.11H6a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3a1 1 0 0 0-1-1M13 5.41L15.59 8H14a1 1 0 0 1-1-1ZM20 14h-2.5a1 1 0 0 0-.71.29l-1.24 1.25l-2.8-3.2a1 1 0 0 0-1.46-.05L9.59 14H8a1 1 0 0 0 0 2h2a1 1 0 0 0 .71-.29L12 14.46l2.8 3.2a1 1 0 0 0 .72.34a1 1 0 0 0 .71-.29L17.91 16H20a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileMinus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 14h-4a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2m6-5.06a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19h-.09a.88.88 0 0 0-.33-.11H7a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V9zm-6-3.53L16.59 8H15a1 1 0 0 1-1-1ZM18 19a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileMinusAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 8a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2Zm5 12H6a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h3v5a1 1 0 0 0 2 0V8.94a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19a.29.29 0 0 0-.1 0a1.1 1.1 0 0 0-.31-.11H6a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h7a1 1 0 0 0 0-2m0-14.59L15.59 8H14a1 1 0 0 1-1-1ZM20 18h-4a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2m-7-2H8a1 1 0 0 0 0 2h5a1 1 0 0 0 0-2m1-4H8a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileNetwork(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 19h-6.18A3 3 0 0 0 13 17.18V15h3a3 3 0 0 0 3-3V7.94a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19a.32.32 0 0 0-.09 0a.88.88 0 0 0-.33-.11H8a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h3v2.18A3 3 0 0 0 9.18 19H3a1 1 0 0 0 0 2h6.18a3 3 0 0 0 5.64 0H21a1 1 0 0 0 0-2M13 4.41L15.59 7H14a1 1 0 0 1-1-1ZM8 13a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h3v3a3 3 0 0 0 3 3h3v3a1 1 0 0 1-1 1Zm4 8a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilePlus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 8.94a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19a.32.32 0 0 0-.09 0a.88.88 0 0 0-.33-.11H7a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V9zm-6-3.53L16.59 8H15a1 1 0 0 1-1-1ZM18 19a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h3Zm-4-5h-1v-1a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2h1v1a1 1 0 0 0 2 0v-1h1a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilePlusAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 18h-1v-1a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2h1v1a1 1 0 0 0 2 0v-1h1a1 1 0 0 0 0-2m-7 2H6a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h3v3a1 1 0 0 0 2 0V8.94a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19a.29.29 0 0 0-.1 0a1.1 1.1 0 0 0-.31-.11H6a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h7a1 1 0 0 0 0-2m0-14.59L15.59 8H14a1 1 0 0 1-1-1ZM8 8a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2Zm5 8H8a1 1 0 0 0 0 2h5a1 1 0 0 0 0-2m1-4H8a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileQuestion(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.57 17.29a1 1 0 0 0-1.41 0a1.06 1.06 0 0 0-.22.33a1.07 1.07 0 0 0 0 .76a1.19 1.19 0 0 0 .22.33a1 1 0 0 0 .32.21a1 1 0 0 0 .39.08a1 1 0 0 0 .92-1.38a.91.91 0 0 0-.22-.33M20 8.94a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19h-.09a.88.88 0 0 0-.33-.11H7a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V9zm-6-3.53L16.59 8H15a1 1 0 0 1-1-1ZM18 19a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h3Zm-6.13-9a3 3 0 0 0-2.6 1.5a1 1 0 1 0 1.73 1a1 1 0 0 1 .87-.5a1 1 0 0 1 0 2a1 1 0 1 0 0 2a3 3 0 0 0 0-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileQuestionAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.07 12h-5a1 1 0 0 0 0 2h5a1 1 0 0 0 0-2m1 8h-8a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h3v1a1 1 0 0 0 2 0V9a.14.14 0 0 0 0-.06a.86.86 0 0 0-.07-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1 1 0 0 0-.29-.19s-.05 0-.08 0a.88.88 0 0 0-.32-.11h-6a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h8a1 1 0 0 0 0-2Zm-1-14.59L15.65 8h-1.58a1 1 0 0 1-1-1Zm5.57 14.88a1.58 1.58 0 0 0-.15-.12a1.08 1.08 0 0 0-.36-.15a1 1 0 0 0-.9.27a1 1 0 0 0 0 1.42a1 1 0 0 0 .7.29a1 1 0 0 0 .93-1.38a1.19 1.19 0 0 0-.22-.33M13.07 16h-5a1 1 0 0 0 0 2h5a1 1 0 0 0 0-2m4.86-3a3 3 0 0 0-2.6 1.5a1 1 0 1 0 1.74 1a1 1 0 1 1 .86 1.5a1 1 0 0 0 0 2a3 3 0 0 0 0-6m-9.86-3h1a1 1 0 0 0 0-2h-1a1 1 0 0 0 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileRedoAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 13a1 1 0 0 0-1-1H8a1 1 0 0 0 0 2h5a1 1 0 0 0 1-1m-3 7H6a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h3v1a1 1 0 0 0 2 0V8.94a1.18 1.18 0 0 0-.06-.27v-.09a.92.92 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19h-.09a.86.86 0 0 0-.32-.11H6a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h5a1 1 0 0 0 0-2m2-14.59L15.59 8H14a1 1 0 0 1-1-1ZM20 14a1 1 0 0 0-.91.6A4.07 4.07 0 0 0 17 14a4 4 0 1 0 2.64 7a1 1 0 0 0-1.32-1.51A2 2 0 0 1 17 20a2 2 0 1 1 1-3.75h-.22a1 1 0 0 0 0 2H20a1 1 0 0 0 1-1V15a1 1 0 0 0-1-1M8 10h1a1 1 0 0 0 0-2H8a1 1 0 0 0 0 2m0 6a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileSearchAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 20H5a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h3v1a1 1 0 0 0 2 0V8.94a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19a.32.32 0 0 0-.09 0a.88.88 0 0 0-.33-.11H5a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h7a1 1 0 0 0 0-2m0-14.59L14.59 8H13a1 1 0 0 1-1-1ZM7 8a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2Zm14.71 12.29l-1.17-1.16A3.44 3.44 0 0 0 20 15a3.49 3.49 0 0 0-6 2.49a3.46 3.46 0 0 0 5.13 3.05l1.16 1.17a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42m-3.17-1.75a1.54 1.54 0 0 1-2.11 0a1.5 1.5 0 0 1-.43-1.05a1.46 1.46 0 0 1 .44-1.06a1.48 1.48 0 0 1 1-.43A1.47 1.47 0 0 1 19 17.49a1.5 1.5 0 0 1-.46 1.05M13 12H7a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2m-2 6a1 1 0 0 0 0-2H7a1 1 0 0 0 0 2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileShareAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 12.5a1 1 0 0 0-1-1h-5a1 1 0 0 0 0 2h5a1 1 0 0 0 1-1m5 5a2 2 0 0 0-1.18.39l-1.75-.8l1.91-.88a2 2 0 0 0 1 .29a2 2 0 1 0-2-2l-1.89.87A2 2 0 1 0 13.5 19a1.88 1.88 0 0 0 .92-.24l2.1 1a2 2 0 1 0 2-2.23Zm-8 2h-5a1 1 0 0 1-1-1v-14a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h3a1 1 0 0 0 2 0V8.44a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19a.29.29 0 0 0-.1 0a1.1 1.1 0 0 0-.26-.06H5.5a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h5a1 1 0 0 0 0-2Zm2-14.59l2.59 2.59H13.5a1 1 0 0 1-1-1Zm-5 10.59a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2Zm0-6h1a1 1 0 0 0 0-2h-1a1 1 0 0 0 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileShieldAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5 20h-5a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h3a1 1 0 0 0 2 0V8.94a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19a.32.32 0 0 0-.09 0L11.56 2H5.5a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h5a1 1 0 0 0 0-2m2-14.59L15.09 8H13.5a1 1 0 0 1-1-1ZM7.5 8a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2Zm13.63 5.86a1 1 0 0 0-.84-.2a2.77 2.77 0 0 1-2.2-.47a1 1 0 0 0-1.18 0a2.78 2.78 0 0 1-2.2.47a1 1 0 0 0-1.21 1V17a4.6 4.6 0 0 0 1.84 3.69l1.56 1.11a1 1 0 0 0 1.2 0l1.56-1.16A4.6 4.6 0 0 0 21.5 17v-2.37a1 1 0 0 0-.37-.77M19.5 17a2.62 2.62 0 0 1-1 2.09l-1 .72l-1-.72a2.62 2.62 0 0 1-1-2.09v-1.28a4.68 4.68 0 0 0 2-.55a4.68 4.68 0 0 0 2 .55Zm-9-.95h-3a1 1 0 0 0 0 2h3a1 1 0 0 0 0-2m1-4h-4a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.71 20.29l-18-18a1 1 0 0 0-1.42 1.42L4 5.41V19a3 3 0 0 0 3 3h10a3 3 0 0 0 2.39-1.2l.9.91a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42M17 20H7a1 1 0 0 1-1-1V7.41l11.93 11.93A1 1 0 0 1 17 20M8.66 4H12v3a3 3 0 0 0 3 3h3v3.34a1 1 0 1 0 2 0v-4.4a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19h-.09L13.06 2h-4.4a1 1 0 0 0 0 2M14 5.41L16.59 8H15a1 1 0 0 1-1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileTimes(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.71 12.29a1 1 0 0 0-1.42 0L12 13.59l-1.29-1.3a1 1 0 0 0-1.42 1.42l1.3 1.29l-1.3 1.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l1.29-1.3l1.29 1.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L13.41 15l1.3-1.29a1 1 0 0 0 0-1.42M20 8.94a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19h-.1a1.1 1.1 0 0 0-.31-.11H7a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V9zm-6-3.53L16.59 8H15a1 1 0 0 1-1-1ZM18 19a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileTimesAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 12H8a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2m5.41 7l1.3-1.29a1 1 0 0 0-1.42-1.42L18 17.59l-1.29-1.3a1 1 0 0 0-1.42 1.42l1.3 1.29l-1.3 1.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l1.29-1.3l1.29 1.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM12 20H6a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h3v3a1 1 0 0 0 2 0V8.94a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19a.29.29 0 0 0-.1 0a1.1 1.1 0 0 0-.31-.11H6a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h6a1 1 0 0 0 0-2m1-14.59L15.59 8H14a1 1 0 0 1-1-1ZM8 8a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2Zm4 8H8a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileUpload(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.71 11.29a1 1 0 0 0-.33-.21a1 1 0 0 0-.76 0a1 1 0 0 0-.33.21l-2 2a1 1 0 0 0 1.42 1.42l.29-.3V17a1 1 0 0 0 2 0v-2.59l.29.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM20 8.94a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19h-.1a1.1 1.1 0 0 0-.31-.11H7a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V9zm-6-3.53L16.59 8H15a1 1 0 0 1-1-1ZM18 19a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileUploadAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 20H6a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h5v3a3 3 0 0 0 3 3h3v2a1 1 0 0 0 2 0V8.94a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19a.32.32 0 0 0-.09 0a.88.88 0 0 0-.33-.11H6a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h7a1 1 0 0 0 0-2m0-14.59L15.59 8H14a1 1 0 0 1-1-1ZM8 8a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2Zm6 4H8a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2m6.71 5.29l-2-2a1 1 0 0 0-.33-.21a1 1 0 0 0-.76 0a1 1 0 0 0-.33.21l-2 2a1 1 0 0 0 1.42 1.42l.29-.3V21a1 1 0 0 0 2 0v-2.59l.29.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42M12 18a1 1 0 0 0 0-2H8a1 1 0 0 0 0 2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilesLandscapes(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23 9.94a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19h-.09a.88.88 0 0 0-.33-.11H8a3 3 0 0 0-3 3v1H4a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-1h1a3 3 0 0 0 3-3v-4zm-6-3.53L19.59 9H18a1 1 0 0 1-1-1ZM17 18a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1v-8a1 1 0 0 1 1-1h1v5a3 3 0 0 0 3 3h9Zm4-4a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h7v3a3 3 0 0 0 3 3h3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilesLandscapesAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 19H6a3 3 0 0 1-3-3V8a1 1 0 0 0-2 0v8a5 5 0 0 0 5 5h12a1 1 0 0 0 0-2m-4-8h-4a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2m9-1.06a1.31 1.31 0 0 0-.06-.27v-.09a1.07 1.07 0 0 0-.19-.28l-6-6a1.07 1.07 0 0 0-.28-.19h-.09a.88.88 0 0 0-.33-.11H8a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-4zm-6-3.53L19.59 9H18a1 1 0 0 1-1-1ZM21 14a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h7v3a3 3 0 0 0 3 3h3ZM10 9h2a1 1 0 0 0 0-2h-2a1 1 0 0 0 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Film(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2a1 1 0 0 0-1 1v2h-2V3a1 1 0 0 0-2 0v1H8V3a1 1 0 0 0-2 0v2H4V3a1 1 0 0 0-2 0v18a1 1 0 0 0 2 0v-2h2v2a1 1 0 0 0 2 0v-1h8v1a1 1 0 0 0 2 0v-2h2v2a1 1 0 0 0 2 0V3a1 1 0 0 0-1-1M6 17H4v-2h2Zm0-4H4v-2h2Zm0-4H4V7h2Zm10 9H8v-5h8Zm0-7H8V6h8Zm4 6h-2v-2h2Zm0-4h-2v-2h2Zm0-4h-2V7h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Filter(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 2H5a3 3 0 0 0-3 3v1.17a3 3 0 0 0 .25 1.2v.06a2.81 2.81 0 0 0 .59.86L9 14.41V21a1 1 0 0 0 .47.85A1 1 0 0 0 10 22a1 1 0 0 0 .45-.11l4-2A1 1 0 0 0 15 19v-4.59l6.12-6.12a2.81 2.81 0 0 0 .59-.86v-.06a3 3 0 0 0 .29-1.2V5a3 3 0 0 0-3-3m-5.71 11.29A1 1 0 0 0 13 14v4.38l-2 1V14a1 1 0 0 0-.29-.71L5.41 8h13.18ZM20 6H4V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.22 6h-6.56a1 1 0 0 0 0 2h6.56a.78.78 0 0 1 .78.78v.78h-3.78a1 1 0 1 0 0 2h2.37l-.7.69a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l1.88-1.88a2.51 2.51 0 0 0 .54-.8v-.1A2.59 2.59 0 0 0 22 9.82v-1A2.79 2.79 0 0 0 19.22 6M3.71 2.29a1 1 0 0 0-1.42 1.42l2.85 2.84A2.73 2.73 0 0 0 4 8.78v1a2.65 2.65 0 0 0 .24 1.1v.06a2.61 2.61 0 0 0 .54.81l5.41 5.4V21a1 1 0 0 0 .47.85a1 1 0 0 0 .53.15a1 1 0 0 0 .45-.11l3.56-1.78a1 1 0 0 0 .55-.89v-2l4.51 4.52a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM6 8.78a.76.76 0 0 1 .5-.72L6.59 8l1.56 1.56H6Zm8.07 7.29a1 1 0 0 0-.29.71v1.82l-1.56.78v-2.6a1 1 0 0 0-.29-.71l-4.52-4.51h2.74l4.22 4.22Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fire(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8.468 8.395l-.002.001l-.003.002Zm9.954-.187a1.237 1.237 0 0 0-.23-.175a1 1 0 0 0-1.4.411a5.782 5.782 0 0 1-1.398 1.778a8.664 8.664 0 0 0 .134-1.51a8.714 8.714 0 0 0-4.4-7.582a1 1 0 0 0-1.492.806a7.017 7.017 0 0 1-2.471 4.942l-.23.187a8.513 8.513 0 0 0-1.988 1.863a8.983 8.983 0 0 0 3.656 13.908a1 1 0 0 0 1.377-.926a1.05 1.05 0 0 0-.05-.312a6.977 6.977 0 0 1-.19-2.581a9.004 9.004 0 0 0 4.313 4.016a.997.997 0 0 0 .715.038a8.995 8.995 0 0 0 3.654-14.863m-3.905 12.831a6.964 6.964 0 0 1-3.577-4.402a8.908 8.908 0 0 1-.18-.964a1 1 0 0 0-.799-.845a.982.982 0 0 0-.191-.018a1 1 0 0 0-.867.5a8.959 8.959 0 0 0-1.205 4.718a6.985 6.985 0 0 1-1.176-9.868a6.555 6.555 0 0 1 1.562-1.458a.745.745 0 0 0 .075-.055s.296-.245.306-.25a8.968 8.968 0 0 0 2.9-4.633a6.736 6.736 0 0 1 1.385 8.088a1 1 0 0 0 1.184 1.418a7.856 7.856 0 0 0 3.862-2.688a7 7 0 0 1-3.279 10.457"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FiveHundredPx(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 15.14a1.41 1.41 0 0 0 .17.44a6.28 6.28 0 0 0 1.39 2.08a6.67 6.67 0 0 0 2.09 1.4a6.21 6.21 0 0 0 2.54.52a6.29 6.29 0 0 0 2.55-.52a6.63 6.63 0 0 0 2.08-1.4a6.39 6.39 0 0 0 1.41-2.08a6.55 6.55 0 0 0 .46-2.58a6.76 6.76 0 0 0-.51-2.56a6.35 6.35 0 0 0-1.41-2.07a6.46 6.46 0 0 0-4.63-1.92a6.63 6.63 0 0 0-2.58.55a7.09 7.09 0 0 0-1.2.68a7.14 7.14 0 0 0-1.14.94V3.23h9.05c.22 0 .34-.21.34-.62S17.48 2 17.26 2H7.47a.37.37 0 0 0-.3.13a.4.4 0 0 0-.12.29V10a.39.39 0 0 0 .17.3a1.09 1.09 0 0 0 .41.18a.73.73 0 0 0 .43 0a.92.92 0 0 0 .24-.11a1 1 0 0 0 .14-.17a6.9 6.9 0 0 1 .86-1a5.15 5.15 0 0 1 3.79-1.56a5.15 5.15 0 0 1 3.81 1.61A5.17 5.17 0 0 1 18.48 13a5.31 5.31 0 0 1-.41 2a5.36 5.36 0 0 1-2.9 3a5.4 5.4 0 0 1-2.06.4a5.09 5.09 0 0 1-2.7-.75V13a2.66 2.66 0 0 1 .71-1.79a2.53 2.53 0 0 1 2-.89a2.65 2.65 0 0 1 2 .79a2.55 2.55 0 0 1 .75 1.89a2.73 2.73 0 0 1-2.77 2.74h-.34l-.37-.07h-.15c-.21-.06-.37.1-.49.48s-.07.6.14.68a4.34 4.34 0 0 0 1.25.18a3.84 3.84 0 0 0 2.8-1.16A3.81 3.81 0 0 0 17.1 13a3.73 3.73 0 0 0-1.16-2.78a3.8 3.8 0 0 0-2.8-1.15a3.86 3.86 0 0 0-2.82 1.15a3.57 3.57 0 0 0-1.14 2.59v3.8a5.63 5.63 0 0 1-1.08-1.86c-.08-.21-.3-.25-.67-.13s-.55.29-.48.49Zm6.25-11a8.16 8.16 0 0 0-3.34.64a.25.25 0 0 0-.23.22a1.26 1.26 0 0 0 .09.43c.14.35.3.48.5.41a8.09 8.09 0 0 1 2.93-.55a7.54 7.54 0 0 1 3.08.63a8.67 8.67 0 0 1 2.31 1.48a.25.25 0 0 0 .18.08c.09 0 .24-.1.43-.29l.19-.19a.37.37 0 0 0 .06-.21a.3.3 0 0 0-.1-.2a8.62 8.62 0 0 0-2.62-1.69a9 9 0 0 0-3.53-.76Zm-1.61 9.91a.46.46 0 0 0 .2.33a.52.52 0 0 0 .35.18a.24.24 0 0 0 .17-.06l.73-.73l.7.68a.28.28 0 0 0 .21.11a.58.58 0 0 0 .36-.19c.2-.21.23-.39.07-.55l-.7-.7l.74-.74c.12-.12.08-.29-.13-.49s-.4-.27-.53-.14l-.73.72l-.72-.74a.3.3 0 0 0-.15-.05a.56.56 0 0 0-.34.2c-.23.22-.28.38-.16.5l.74.74l-.74.72a.35.35 0 0 0-.12.21Zm7.93 4.57a1.24 1.24 0 0 0-.37-.25a.28.28 0 0 0-.28.07l-.07.07a8 8 0 0 1-2.51 1.69a7.94 7.94 0 0 1-8.68-1.69A7.48 7.48 0 0 1 5.91 16a9.33 9.33 0 0 1-.51-1.77c0-.21-.25-.29-.63-.23s-.56.2-.53.4a8.52 8.52 0 0 0 .6 2.11a8.84 8.84 0 0 0 2 2.88a9 9 0 0 0 2.89 2a9.13 9.13 0 0 0 3.54.71a9.28 9.28 0 0 0 3.54-.71a9 9 0 0 0 2.89-2l.06-.06c.09-.22.02-.44-.24-.71Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flask(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.11 17.49L15 8.73V4h1a1 1 0 0 0 0-2H8a1 1 0 0 0 0 2h1v4.73l-5.11 8.76A3 3 0 0 0 6.48 22h11a3 3 0 0 0 2.59-4.51Zm-9.25-8A1 1 0 0 0 11 9V4h2v5a1 1 0 0 0 .14.5L14 11h-4Zm7.52 10a1 1 0 0 1-.86.5h-11a1 1 0 0 1-.86-.5a1 1 0 0 1 0-1L8.83 13h6.35l3.2 5.5a1 1 0 0 1 0 1ZM10 15a1 1 0 1 0 1 1a1 1 0 0 0-1-1m4 1a1 1 0 1 0 1 1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlaskPotion(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.93 11.67a.42.42 0 0 0 0-.1A7.4 7.4 0 0 0 15 7.62V4h1a1 1 0 0 0 0-2H8a1 1 0 0 0 0 2h1v3.62a7.4 7.4 0 0 0-3.89 4a.42.42 0 0 0 0 .1a7.5 7.5 0 1 0 13.86 0Zm-8.62-2.41a1 1 0 0 0 .69-.95V4h2v4.31a1 1 0 0 0 .69.95A5.43 5.43 0 0 1 16.23 11H7.77a5.43 5.43 0 0 1 2.54-1.74M12 20a5.51 5.51 0 0 1-5.5-5.5a5.34 5.34 0 0 1 .22-1.5h10.56a5.34 5.34 0 0 1 .22 1.5A5.51 5.51 0 0 1 12 20m2-4a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-4-1a1 1 0 1 0 1 1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlipH(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 11H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2M10.93 9h1.5a1 1 0 0 0 0-2h-1.5a1 1 0 0 0 0 2m4.5-1a1 1 0 0 0 1 1H17a1 1 0 0 0 .92-.62a1 1 0 0 0-.21-1.09l-.66-.65a1 1 0 0 0-1.41 0a1 1 0 0 0-.19 1.15a1.49 1.49 0 0 0-.02.21m-3.78-3.23l.35-.36l.81.81a1 1 0 0 0 1.41 0a1 1 0 0 0 0-1.41l-1.06-1.06a.91.91 0 0 0-.26-.19a1 1 0 0 0-1.61-.27l-1.06 1.06a1 1 0 0 0 1.42 1.42M17 15H7a1 1 0 0 0-.92.62a1 1 0 0 0 .21 1.09l5 5a1 1 0 0 0 1.42 0l5-5a1 1 0 0 0 .21-1.09A1 1 0 0 0 17 15m-5 4.59L9.41 17h5.18ZM7.05 9a1 1 0 0 0 .71-.29L8.82 7.6A1 1 0 0 0 7.4 6.18L6.34 7.24a1 1 0 0 0 0 1.42a1 1 0 0 0 .71.34"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlipHalt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.21 5.7a1 1 0 0 0 .24 0l1.94-.49A1 1 0 0 0 16.12 4a1 1 0 0 0-1.21-.73L13 3.73a1 1 0 0 0 .24 2ZM9.51 9h-2a1 1 0 1 0 0 2h2a1 1 0 0 0 0-2m4 0a1 1 0 1 0 0 2h2a1 1 0 0 0 0-2ZM7.39 7.15h.24l1.94-.48a1 1 0 0 0-.48-1.97l-1.94.48a1 1 0 0 0 .24 2ZM4 10.51a1 1 0 0 0 1-1v-2a1 1 0 0 0-2 0v2a1 1 0 0 0 1 1m16.62-8.3a1 1 0 0 0-.86-.21l-1 .24a1 1 0 0 0-.73 1.21a1 1 0 0 0 1 .76A1 1 0 0 0 21 4V3a1 1 0 0 0-.38-.79M20 7a1 1 0 0 0-1 1v1.14a1 1 0 0 0 .51 1.86H20a1 1 0 0 0 1-1V8a1 1 0 0 0-1-1m0 6H4a1 1 0 0 0-1 1v3a1 1 0 0 0 .76 1l16 4a1 1 0 0 0 .24 0a1 1 0 0 0 1-1v-7a1 1 0 0 0-1-1m-1 6.72l-14-3.5V15h14Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlipV(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.6 15.18a1 1 0 0 0-1.42 1.42l1.06 1.06a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Zm-3.79-5.4l-1.06 1.06a.91.91 0 0 0-.19.26a1 1 0 0 0-.27 1.61l1.06 1.06a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L4.41 12l.81-.81a1 1 0 0 0-1.41-1.41m4.19.79a1 1 0 0 0-1 1v1.5a1 1 0 0 0 2 0v-1.5a1 1 0 0 0-1-1m13.71.72l-5-5a1 1 0 0 0-1.09-.21A1 1 0 0 0 15 7v10a1 1 0 0 0 .62.92a.84.84 0 0 0 .38.08a1 1 0 0 0 .71-.29l5-5a1 1 0 0 0 0-1.42M17 14.59V9.41L19.59 12ZM12 2a1 1 0 0 0-1 1v18a1 1 0 0 0 2 0V3a1 1 0 0 0-1-1M8.38 6.08a1 1 0 0 0-1.09.21L6.64 7a1 1 0 0 0 0 1.41a1 1 0 0 0 .7.3a1 1 0 0 0 .45-.11A1 1 0 0 0 9 7.57V7a1 1 0 0 0-.62-.92"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlipValt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.94 12.24a1 1 0 0 0-1.21.76l-.49 1.94A1 1 0 0 0 4 16.12a1 1 0 0 0 1.21-.73l.49-1.94a1 1 0 0 0-.76-1.21m17 7.52l-4-16A1 1 0 0 0 17 3h-3a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h7a1 1 0 0 0 .79-.38a1 1 0 0 0 .21-.86ZM15 19V5h1.22l3.5 14ZM6.4 6.42a1 1 0 0 0-1.22.73L4.7 9.09a1 1 0 0 0 .73 1.21h.24a1 1 0 0 0 1-.76l.48-1.94a1 1 0 0 0-.75-1.18M7.51 5h2a1 1 0 0 0 0-2h-2a1 1 0 0 0 0 2M4.24 19a1 1 0 0 0-2-.24l-.24 1a1 1 0 0 0 .18.86A1 1 0 0 0 3 21h1a1 1 0 0 0 .24-2M10 6.51a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0v-2a1 1 0 0 0-1-1m0 6a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0v-2a1 1 0 0 0-1-1m0 6a1 1 0 0 0-.86.49H8a1 1 0 0 0 0 2h2a1 1 0 0 0 1-1v-.49a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flower(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.68 6.88a4.4 4.4 0 0 0-3.31-.32a4.37 4.37 0 0 0-8.73 0a4.48 4.48 0 0 0-3.31.29a4.37 4.37 0 0 0 .61 8a4.4 4.4 0 0 0-.8 2.5a5 5 0 0 0 .07.75a4.34 4.34 0 0 0 4.29 3.63a4.68 4.68 0 0 0 .64 0A4.42 4.42 0 0 0 12 20a4.42 4.42 0 0 0 2.86 1.69a4.68 4.68 0 0 0 .64 0a4.36 4.36 0 0 0 3.56-6.87a4.36 4.36 0 0 0 .62-8Zm-9.34-1.94a2.4 2.4 0 0 1 3.32 0a2.43 2.43 0 0 1 .52 2.66l-.26.59l-.66.58A4.07 4.07 0 0 0 12 8.55a4 4 0 0 0-1.61.34L9.83 7.6a2.39 2.39 0 0 1 .51-2.66m-6.1 6.84A2.37 2.37 0 0 1 7.94 9l.49.43l.35.8A3.92 3.92 0 0 0 8 12.55A2.85 2.85 0 0 0 8 13h-.55l-.84.08a2.37 2.37 0 0 1-2.37-1.3m6.6 6.08a2.38 2.38 0 0 1-4.66-.08a3.07 3.07 0 0 1 0-.42a2.33 2.33 0 0 1 1.17-2l.51-.36l.91-.1a4 4 0 0 0 2.38 1.57ZM12 14.55a2 2 0 1 1 2-2a2 2 0 0 1-2 2m5.82 3.22a2.36 2.36 0 0 1-2.68 1.94a2.39 2.39 0 0 1-2-1.85l-.14-.6l.21-.92a4 4 0 0 0 2.2-1.76l.5.3H16l.66.39a2.38 2.38 0 0 1 1.16 2.5m1.94-6a2.39 2.39 0 0 1-2.13 1.33h-.24l-.64-.1l-.75-.41a4 4 0 0 0-1-2.64l.43-.37l.63-.58a2.37 2.37 0 0 1 3.7 2.82Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Focus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 2H3a1 1 0 0 0-1 1v5a1 1 0 0 0 2 0V4h4a1 1 0 0 0 0-2m0 18H4v-4a1 1 0 0 0-2 0v5a1 1 0 0 0 1 1h5a1 1 0 0 0 0-2M21 2h-5a1 1 0 0 0 0 2h4v4a1 1 0 0 0 2 0V3a1 1 0 0 0-1-1m0 13a1 1 0 0 0-1 1v4h-4a1 1 0 0 0 0 2h5a1 1 0 0 0 1-1v-5a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FocusAdd(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 20H5a1 1 0 0 1-1-1v-3a1 1 0 0 0-2 0v3a3 3 0 0 0 3 3h3a1 1 0 0 0 0-2M3 9a1 1 0 0 0 1-1V5a1 1 0 0 1 1-1h3a1 1 0 0 0 0-2H5a3 3 0 0 0-3 3v3a1 1 0 0 0 1 1m16-7h-3a1 1 0 0 0 0 2h3a1 1 0 0 1 1 1v3a1 1 0 0 0 2 0V5a3 3 0 0 0-3-3m-3 10a1 1 0 0 0-1-1h-2V9a1 1 0 0 0-2 0v2H9a1 1 0 0 0 0 2h2v2a1 1 0 0 0 2 0v-2h2a1 1 0 0 0 1-1m5 3a1 1 0 0 0-1 1v3a1 1 0 0 1-1 1h-3a1 1 0 0 0 0 2h3a3 3 0 0 0 3-3v-3a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FocusTarget(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 9a1 1 0 0 0 1-1V5a1 1 0 0 1 1-1h3a1 1 0 0 0 0-2H5a3 3 0 0 0-3 3v3a1 1 0 0 0 1 1m5 11H5a1 1 0 0 1-1-1v-3a1 1 0 0 0-2 0v3a3 3 0 0 0 3 3h3a1 1 0 0 0 0-2m9-7a1 1 0 0 0 0-2h-1.14A4 4 0 0 0 13 8.14V7a1 1 0 0 0-2 0v1.14A4 4 0 0 0 8.14 11H7a1 1 0 0 0 0 2h1.14A4 4 0 0 0 11 15.86V17a1 1 0 0 0 2 0v-1.14A4 4 0 0 0 15.86 13Zm-5 1a2 2 0 1 1 2-2a2 2 0 0 1-2 2m9 1a1 1 0 0 0-1 1v3a1 1 0 0 1-1 1h-3a1 1 0 0 0 0 2h3a3 3 0 0 0 3-3v-3a1 1 0 0 0-1-1M19 2h-3a1 1 0 0 0 0 2h3a1 1 0 0 1 1 1v3a1 1 0 0 0 2 0V5a3 3 0 0 0-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folder(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 5.5h-6.28l-.32-1a3 3 0 0 0-2.84-2H5a3 3 0 0 0-3 3v13a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-10a3 3 0 0 0-3-3m1 13a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-13a1 1 0 0 1 1-1h4.56a1 1 0 0 1 .95.68l.54 1.64a1 1 0 0 0 .95.68h7a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderCheck(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.29 10.79L11 14.09l-1.29-1.3a1 1 0 0 0-1.42 1.42l2 2a1 1 0 0 0 1.42 0l4-4a1 1 0 0 0-1.42-1.42M19 5.5h-6.28l-.32-1a3 3 0 0 0-2.84-2H5a3 3 0 0 0-3 3v13a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-10a3 3 0 0 0-3-3m1 13a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-13a1 1 0 0 1 1-1h4.56a1 1 0 0 1 .95.68l.54 1.64a1 1 0 0 0 .95.68h7a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderDownload(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.29 13.79l-.29.3V11.5a1 1 0 0 0-2 0v2.59l-.29-.3a1 1 0 0 0-1.42 1.42l2 2a1 1 0 0 0 .33.21a.94.94 0 0 0 .76 0a1 1 0 0 0 .33-.21l2-2a1 1 0 0 0-1.42-1.42M19 5.5h-6.28l-.32-1a3 3 0 0 0-2.84-2H5a3 3 0 0 0-3 3v13a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-10a3 3 0 0 0-3-3m1 13a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-13a1 1 0 0 1 1-1h4.56a1 1 0 0 1 .95.68l.54 1.64a1 1 0 0 0 .95.68h7a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderExclamation(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.92 16.12a.76.76 0 0 0-.09-.18a1.58 1.58 0 0 0-.12-.15l-.15-.12l-.18-.09a.6.6 0 0 0-.19-.06a1 1 0 0 0-.9.27l-.12.15a.76.76 0 0 0-.09.18a.64.64 0 0 0-.06.18a1.36 1.36 0 0 0 0 .2a1 1 0 0 0 .08.38a1.15 1.15 0 0 0 .21.33a1 1 0 0 0 1.09.22a1.46 1.46 0 0 0 .33-.22a1.46 1.46 0 0 0 .22-.33a1 1 0 0 0 .05-.38a1.36 1.36 0 0 0 0-.2a.64.64 0 0 0-.08-.18M12 10.5a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0v-2a1 1 0 0 0-1-1m7-5h-6.28l-.32-1a3 3 0 0 0-2.84-2H5a3 3 0 0 0-3 3v13a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-10a3 3 0 0 0-3-3m1 13a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-13a1 1 0 0 1 1-1h4.56a1 1 0 0 1 .95.68l.54 1.64a1 1 0 0 0 .95.68h7a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderHeart(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 9.84a2.92 2.92 0 0 0-3.43 4.65l2.72 2.72a1 1 0 0 0 1.42 0l2.72-2.72A2.92 2.92 0 0 0 12 9.84m2 3.23l-2 2l-2-2a.92.92 0 0 1 0-1.3a.92.92 0 0 1 1.3 0a1 1 0 0 0 1.42 0a.92.92 0 0 1 1.3 1.3Zm5-7.57h-6.28l-.32-1a3 3 0 0 0-2.84-2H5a3 3 0 0 0-3 3v13a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-10a3 3 0 0 0-3-3m1 13a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-13a1 1 0 0 1 1-1h4.56a1 1 0 0 1 .95.68l.54 1.64a1 1 0 0 0 .95.68h7a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderInfo(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 13.5a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0v-2a1 1 0 0 0-1-1m7-8h-6.28l-.32-1a3 3 0 0 0-2.84-2H5a3 3 0 0 0-3 3v13a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-10a3 3 0 0 0-3-3m1 13a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-13a1 1 0 0 1 1-1h4.56a1 1 0 0 1 .95.68l.54 1.64a1 1 0 0 0 .95.68h7a1 1 0 0 1 1 1Zm-7.29-7.71a1 1 0 0 0-1.09-.21a.93.93 0 0 0-.33.21a1.05 1.05 0 0 0-.29.71a1 1 0 0 0 .29.7a1 1 0 0 0 1.42 0a1 1 0 0 0 .29-.7a1 1 0 0 0-.08-.38a.93.93 0 0 0-.21-.33"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderLock(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 4h4.56a1 1 0 0 1 .95.68l.54 1.64A1 1 0 0 0 11 7h7a1 1 0 0 1 1 1a1 1 0 0 0 2 0a3 3 0 0 0-3-3h-6.28l-.32-1a3 3 0 0 0-2.84-2H4a3 3 0 0 0-3 3v13a3 3 0 0 0 3 3h6a1 1 0 0 0 0-2H4a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1m17 11.18V14a3 3 0 0 0-6 0v1.18A3 3 0 0 0 13 18v2a3 3 0 0 0 3 3h4a3 3 0 0 0 3-3v-2a3 3 0 0 0-2-2.82M17 14a1 1 0 0 1 2 0v1h-2Zm4 6a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderMedical(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 12.5h-1v-1a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2h1v1a1 1 0 0 0 2 0v-1h1a1 1 0 0 0 0-2m5-7h-6.28l-.32-1a3 3 0 0 0-2.84-2H5a3 3 0 0 0-3 3v13a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-10a3 3 0 0 0-3-3m1 13a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-13a1 1 0 0 1 1-1h4.56a1 1 0 0 1 .95.68l.54 1.64a1 1 0 0 0 .95.68h7a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderMinus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 12.5h-4a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2m5-7h-6.28l-.32-1a3 3 0 0 0-2.84-2H5a3 3 0 0 0-3 3v13a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-10a3 3 0 0 0-3-3m1 13a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-13a1 1 0 0 1 1-1h4.56a1 1 0 0 1 .95.68l.54 1.64a1 1 0 0 0 .95.68h7a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderNetwork(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 18h-6.18A3 3 0 0 0 13 16.18V14h3.67A2.34 2.34 0 0 0 19 11.67V6.33A2.34 2.34 0 0 0 16.67 4h-4l-.13-.41A2.34 2.34 0 0 0 10.37 2h-3A2.34 2.34 0 0 0 5 4.33v7.34A2.34 2.34 0 0 0 7.33 14H11v2.18A3 3 0 0 0 9.18 18H3a1 1 0 0 0 0 2h6.18a3 3 0 0 0 5.64 0H21a1 1 0 0 0 0-2M7.33 12a.33.33 0 0 1-.33-.33V4.33A.33.33 0 0 1 7.33 4h3a.33.33 0 0 1 .32.23l.36 1.09A1 1 0 0 0 12 6h4.67a.33.33 0 0 1 .33.33v5.34a.33.33 0 0 1-.33.33ZM12 20a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOpen(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.78 10.37A1 1 0 0 0 22 10h-2V9a3 3 0 0 0-3-3h-6.28l-.32-1a3 3 0 0 0-2.84-2H4a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h14.4a3 3 0 0 0 2.92-2.35L23 11.22a1 1 0 0 0-.22-.85M5.37 18.22a1 1 0 0 1-1 .78H4a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h3.56a1 1 0 0 1 1 .68l.54 1.64A1 1 0 0 0 10 8h7a1 1 0 0 1 1 1v1H8a1 1 0 0 0-1 .78Zm14 0a1 1 0 0 1-1 .78H7.21a1.42 1.42 0 0 0 .11-.35L8.8 12h12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderPlus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 12.5h-1v-1a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2h1v1a1 1 0 0 0 2 0v-1h1a1 1 0 0 0 0-2m5-7h-6.28l-.32-1a3 3 0 0 0-2.84-2H5a3 3 0 0 0-3 3v13a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-10a3 3 0 0 0-3-3m1 13a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-13a1 1 0 0 1 1-1h4.56a1 1 0 0 1 .95.68l.54 1.64a1 1 0 0 0 .95.68h7a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderQuestion(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.57 16.3a.64.64 0 0 0-.15-.13l-.17-.09l-.19-.08a1 1 0 0 0-.9.28a1 1 0 0 0-.22.32a1 1 0 0 0-.07.39a1 1 0 0 0 .29.7a1 1 0 0 0 .32.22a1 1 0 0 0 .39.07a1 1 0 0 0 .38-.07a1 1 0 0 0 .32-.22a1 1 0 0 0 .3-.7a1 1 0 0 0-.08-.39a.87.87 0 0 0-.22-.3m-.7-7.3a3 3 0 0 0-2.6 1.5a1 1 0 1 0 1.73 1a1 1 0 0 1 1.87.5a1 1 0 0 1-1 1a1 1 0 1 0 0 2a3 3 0 0 0 0-6M19 5.5h-6.28l-.32-1a3 3 0 0 0-2.84-2H5a3 3 0 0 0-3 3v13a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-10a3 3 0 0 0-3-3m1 13a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-13a1 1 0 0 1 1-1h4.56a1 1 0 0 1 .95.68l.54 1.64a1 1 0 0 0 .95.68h7a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.71 20.29l-1.6-1.6l-16.4-16.4a1 1 0 0 0-1.42 1.42l1.4 1.39A3 3 0 0 0 3 7v11a3 3 0 0 0 3 3h12a3 3 0 0 0 1.29-.3l1 1a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.41M6 19a1 1 0 0 1-1-1V7a1 1 0 0 1 .12-.46L17.59 19Zm4.62-13a1 1 0 0 1 .89.67l.54 1.64A1 1 0 0 0 13 9h5a1 1 0 0 1 1 1v4.34a1 1 0 1 0 2 0V10a3 3 0 0 0-3-3h-4.28l-.32-1a3 3 0 0 0-2.68-2a1 1 0 0 0-.1 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderTimes(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.71 10.79a1 1 0 0 0-1.42 0L12 12.09l-1.29-1.3a1 1 0 0 0-1.42 1.42l1.3 1.29l-1.3 1.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l1.29-1.3l1.29 1.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-1.3-1.29l1.3-1.29a1 1 0 0 0 0-1.42M19 5.5h-6.28l-.32-1a3 3 0 0 0-2.84-2H5a3 3 0 0 0-3 3v13a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-10a3 3 0 0 0-3-3m1 13a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-13a1 1 0 0 1 1-1h4.56a1 1 0 0 1 .95.68l.54 1.64a1 1 0 0 0 .95.68h7a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderUpload(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.71 10.79a1 1 0 0 0-.33-.21a1 1 0 0 0-.76 0a1 1 0 0 0-.33.21l-2 2a1 1 0 0 0 1.42 1.42l.29-.3v2.59a1 1 0 0 0 2 0v-2.59l.29.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM19 5.5h-6.28l-.32-1a3 3 0 0 0-2.84-2H5a3 3 0 0 0-3 3v13a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-10a3 3 0 0 0-3-3m1 13a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-13a1 1 0 0 1 1-1h4.56a1 1 0 0 1 .95.68l.54 1.64a1 1 0 0 0 .95.68h7a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Font(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 20h-1V3a1 1 0 0 0-1-1h-3.5a1 1 0 0 0-.86.5L4.43 20H3a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2h-.26l3.5-6H18v6h-1a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2m-3-8h-6.59l4.66-8H18Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Football(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.07 6.11a9.85 9.85 0 0 0-4.3-3.36A10 10 0 0 0 2 12v.56A9.94 9.94 0 0 0 3.33 17a10 10 0 0 0 5.89 4.65A10.11 10.11 0 0 0 12 22a9.45 9.45 0 0 0 1.88-.18a10 10 0 0 0 8-8.41A9.46 9.46 0 0 0 22 12a9.83 9.83 0 0 0-1.93-5.89m-2 .77L17 9.74l-1.62.44L13 8.49V6.64l2.49-1.81a7.81 7.81 0 0 1 2.62 2.05ZM14 11.67L13.22 14h-2.45L10 11.67l2-1.43ZM12 4a8 8 0 0 1 1.11.09l-1.11.8l-1.11-.8A8 8 0 0 1 12 4M4.88 8.37l.4 1.32l-1.13.79a7.88 7.88 0 0 1 .73-2.11m1.37 9.17l1.38.05l.37 1.33a8.32 8.32 0 0 1-1.75-1.38M8 15.6l-3.15-.11A7.83 7.83 0 0 1 4.07 13l2.49-1.74l1.44.62l.89 2.76Zm.86-5.53l-1.56-.7l-.91-3A7.93 7.93 0 0 1 8.5 4.83L11 6.64v1.85ZM13 19.93a8.08 8.08 0 0 1-2.63-.12l-.83-2.92l.83-.89h3.07l.67 1Zm2.41-.7l.46-1.23l1.36.07a7.83 7.83 0 0 1-1.85 1.16Zm3.46-3.12L15.76 16l-.71-1.1l.89-2.76l1.51-.41l2.36 2a7.84 7.84 0 0 1-.97 2.38Zm.05-5.83L19.4 9a7.4 7.4 0 0 1 .53 2.13Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FootballAmerican(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.84 5.56a4.08 4.08 0 0 0-1.14-2.25a4.08 4.08 0 0 0-2.25-1.14a13.65 13.65 0 0 0-5.29.24a1.17 1.17 0 0 0-.2.06a14.44 14.44 0 0 0-6.69 3.8A14.59 14.59 0 0 0 2.45 13c0 .06 0 .12-.05.19a13.7 13.7 0 0 0-.24 5.3a4.08 4.08 0 0 0 1.14 2.2a4.08 4.08 0 0 0 2.25 1.14a13.12 13.12 0 0 0 2.08.17a13.8 13.8 0 0 0 3.26-.41h.14a14.54 14.54 0 0 0 10.52-10.5c0-.06 0-.12.05-.19a13.7 13.7 0 0 0 .24-5.34M16.37 4a10.44 10.44 0 0 1 1.76.14a1.68 1.68 0 0 1 .24.07L17 5.59l-1.54-1.54c.3-.05.61-.05.91-.05m-8.7 3.67a12.72 12.72 0 0 1 5.4-3.19L15.59 7l-2.15 2.15l-.73-.73a1 1 0 0 0-1.42 0a1 1 0 0 0 0 1.41l.73.73L10.56 12l-.73-.73a1 1 0 0 0-1.41 0a1 1 0 0 0 0 1.42l.73.73L7 15.59l-2.51-2.51a12.67 12.67 0 0 1 3.18-5.41m-3.46 10.7a1.68 1.68 0 0 1-.07-.24A11.38 11.38 0 0 1 4 15.46L5.59 17Zm1.66 1.49a1.68 1.68 0 0 1-.24-.07L7 18.41L8.54 20a11.38 11.38 0 0 1-2.67-.14m10.46-3.53a12.67 12.67 0 0 1-5.41 3.18L8.41 17l2.15-2.15l.73.73a1 1 0 1 0 1.42-1.41l-.71-.73L13.44 12l.73.73a1 1 0 0 0 .71.29a1 1 0 0 0 .7-1.71l-.73-.73L17 8.41l2.51 2.51a12.67 12.67 0 0 1-3.18 5.41M20 8.54L18.41 7l1.38-1.37a1.68 1.68 0 0 1 .07.24A11.38 11.38 0 0 1 20 8.54"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FootballBall(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.84 5.56a4.08 4.08 0 0 0-1.14-2.25a4.08 4.08 0 0 0-2.25-1.14A14.45 14.45 0 0 0 2.16 18.44a4.08 4.08 0 0 0 1.14 2.25a4.08 4.08 0 0 0 2.25 1.14a13.12 13.12 0 0 0 2.08.17a14.37 14.37 0 0 0 10.11-4.26a14.23 14.23 0 0 0 4.1-12.18M4.21 18.37a1.68 1.68 0 0 1-.07-.24A12.21 12.21 0 0 1 7.67 7.67A12.39 12.39 0 0 1 16.37 4a10.44 10.44 0 0 1 1.76.14a1.68 1.68 0 0 1 .24.07l-4.93 4.94l-.73-.73a1 1 0 0 0-1.42 0a1 1 0 0 0 0 1.41l.73.73L10.56 12l-.73-.73a1 1 0 0 0-1.41 0a1 1 0 0 0 0 1.42l.73.73Zm12.12-2a12.24 12.24 0 0 1-10.46 3.49a1.68 1.68 0 0 1-.24-.07l4.93-4.94l.73.73a1 1 0 1 0 1.42-1.41l-.71-.73L13.44 12l.73.73a1 1 0 0 0 .71.29a1 1 0 0 0 .7-1.71l-.73-.73l4.94-4.93a1.68 1.68 0 0 1 .07.24a12.21 12.21 0 0 1-3.53 10.44Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ForecastcloudMoonTear(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.7 7.57a1 1 0 0 0-.93-.26a3.2 3.2 0 0 1-.66.08a3 3 0 0 1-3-3a3 3 0 0 1 .08-.65A1 1 0 0 0 16 2.53a4.93 4.93 0 0 0-3.83 4.21a6.24 6.24 0 0 0-1.67-.24a6 6 0 0 0-5.94 5.13a3.5 3.5 0 0 0-.46 6.58a1.14 1.14 0 0 0 .4.08a1 1 0 0 0 .4-1.92A1.48 1.48 0 0 1 4 15a1.5 1.5 0 0 1 1.5-1.5a1 1 0 0 0 1-1a4 4 0 0 1 4-4a3.92 3.92 0 0 1 2.18.66a4 4 0 0 1 1.57 2a1 1 0 0 0 .78.67a2.32 2.32 0 0 1 .97 4.28a1 1 0 0 0 1.1 1.68a4.32 4.32 0 0 0 1.9-3.62a4.23 4.23 0 0 0-.49-2A4.94 4.94 0 0 0 22 8.5a1 1 0 0 0-.3-.93m-4.59 2.82a2.61 2.61 0 0 1-.42 0A4.6 4.6 0 0 0 16 10a6 6 0 0 0-1.82-2.28v-.37a3 3 0 0 1 1.05-2.28a5 5 0 0 0 4.23 4.23a3 3 0 0 1-2.35 1.09m-6 3.94a1 1 0 0 0-1.12 0C9.84 14.41 7.5 16 7.5 18.5a3 3 0 0 0 6 0c0-2.5-2.35-4.1-2.44-4.17Zm-.61 5.17a1 1 0 0 1-1-1a3 3 0 0 1 1-2a3 3 0 0 1 1 2a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ForwadedCall(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.94 6.56h3.58l-.79.8a1 1 0 0 0 0 1.41a1 1 0 0 0 1.41 0l2.5-2.5a1 1 0 0 0 0-1.41l-2.5-2.5a1 1 0 0 0-1.41 1.41l.79.79h-3.58a1 1 0 0 0 0 2m4.5 6.44c-.22 0-.45-.07-.67-.12a9.44 9.44 0 0 1-1.31-.39a2 2 0 0 0-2.48 1l-.22.45a12.18 12.18 0 0 1-2.66-2a12.18 12.18 0 0 1-2-2.66l.42-.28a2 2 0 0 0 1-2.48a10.33 10.33 0 0 1-.39-1.31c-.05-.22-.09-.45-.12-.68a3 3 0 0 0-3-2.49h-3a3 3 0 0 0-3 3.41a19 19 0 0 0 16.52 16.46h.38a3 3 0 0 0 2-.76a3 3 0 0 0 1-2.25v-3a3 3 0 0 0-2.47-2.9m.5 6a1 1 0 0 1-.34.75a1.06 1.06 0 0 1-.82.25A17 17 0 0 1 4.07 5.22a1.09 1.09 0 0 1 .25-.82a1 1 0 0 1 .75-.34h3a1 1 0 0 1 1 .79q.06.41.15.81a11.12 11.12 0 0 0 .46 1.55l-1.4.65a1 1 0 0 0-.49 1.33a14.49 14.49 0 0 0 7 7a1 1 0 0 0 .76 0a1 1 0 0 0 .57-.52l.62-1.4a13.69 13.69 0 0 0 1.58.46q.4.09.81.15a1 1 0 0 1 .79 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Forward(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.67 9.69L14 5.84a2.67 2.67 0 0 0-4 2.31L6 5.84a2.67 2.67 0 0 0-4 2.31v7.7a2.63 2.63 0 0 0 1.33 2.3a2.61 2.61 0 0 0 1.34.37A2.69 2.69 0 0 0 6 18.16l4-2.31a2.65 2.65 0 0 0 1.33 2.31a2.66 2.66 0 0 0 2.67 0l6.67-3.85a2.67 2.67 0 0 0 0-4.62M10 13.54l-5 2.88a.67.67 0 0 1-1-.57v-7.7a.67.67 0 0 1 1-.57l5 2.88Zm9.67-1L13 16.43a.69.69 0 0 1-.67 0a.66.66 0 0 1-.33-.58v-7.7a.66.66 0 0 1 .33-.58a.78.78 0 0 1 .34-.09a.63.63 0 0 1 .33.09l6.67 3.85a.67.67 0 0 1 0 1.16Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Frown(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.36 15.33a1 1 0 0 0-.13 1.4a1 1 0 0 0 1.41.13a3.76 3.76 0 0 1 4.72 0a1 1 0 0 0 1.41-.13a1 1 0 0 0-.13-1.4a5.81 5.81 0 0 0-7.28 0M9 11a1 1 0 1 0-1-1a1 1 0 0 0 1 1m3-9a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8m3-11a1 1 0 1 0 1 1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GameStructure(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 18h-2v-3a1 1 0 0 0-1-1h-5v-2.71l1.13.59a1 1 0 0 0 1.45-1.05l-.4-2.37l1.72-1.69a1 1 0 0 0 .26-1a1 1 0 0 0-.81-.68L14 4.72l-1.1-2.16a1 1 0 0 0-1.8 0L10 4.72l-2.39.35a1 1 0 0 0-.81.68a1 1 0 0 0 .26 1l1.76 1.71l-.4 2.37a1 1 0 0 0 1.45 1.05l1.13-.59V14H6a1 1 0 0 0-1 1v3H3a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0v-1h4v1a1 1 0 0 0 2 0v-2a1 1 0 0 0-1-1H7v-2h10v2h-2a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0v-1h4v1a1 1 0 0 0 2 0v-2a1 1 0 0 0-1-1m-9-9.37a1 1 0 0 0-.47.12l-.8.42l.15-.9a1 1 0 0 0-.29-.88l-.65-.64l.9-.13a1 1 0 0 0 .76-.54l.4-.82l.4.82a1 1 0 0 0 .76.54l.9.13l-.65.64a1 1 0 0 0-.29.88l.15.9l-.8-.42a1 1 0 0 0-.47-.12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gift(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 7h-.35A3.45 3.45 0 0 0 18 5.5a3.49 3.49 0 0 0-6-2.44A3.49 3.49 0 0 0 6 5.5A3.45 3.45 0 0 0 6.35 7H6a3 3 0 0 0-3 3v2a1 1 0 0 0 1 1h1v6a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3v-6h1a1 1 0 0 0 1-1v-2a3 3 0 0 0-3-3m-7 13H8a1 1 0 0 1-1-1v-6h4Zm0-9H5v-1a1 1 0 0 1 1-1h5Zm0-4H9.5A1.5 1.5 0 1 1 11 5.5Zm2-1.5A1.5 1.5 0 1 1 14.5 7H13ZM17 19a1 1 0 0 1-1 1h-3v-7h4Zm2-8h-6V9h5a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Github(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2.247a10 10 0 0 0-3.162 19.487c.5.088.687-.212.687-.475c0-.237-.012-1.025-.012-1.862c-2.513.462-3.163-.613-3.363-1.175a3.636 3.636 0 0 0-1.025-1.413c-.35-.187-.85-.65-.013-.662a2.001 2.001 0 0 1 1.538 1.025a2.137 2.137 0 0 0 2.912.825a2.104 2.104 0 0 1 .638-1.338c-2.225-.25-4.55-1.112-4.55-4.937a3.892 3.892 0 0 1 1.025-2.688a3.594 3.594 0 0 1 .1-2.65s.837-.262 2.75 1.025a9.427 9.427 0 0 1 5 0c1.912-1.3 2.75-1.025 2.75-1.025a3.593 3.593 0 0 1 .1 2.65a3.869 3.869 0 0 1 1.025 2.688c0 3.837-2.338 4.687-4.562 4.937a2.368 2.368 0 0 1 .674 1.85c0 1.338-.012 2.413-.012 2.75c0 .263.187.575.687.475A10.005 10.005 0 0 0 12 2.247"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GithubAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.07 20.503a1 1 0 0 0-1.18-.983c-1.31.24-2.963.276-3.402-.958a5.708 5.708 0 0 0-1.837-2.415a1.2 1.2 0 0 1-.167-.11a1 1 0 0 0-.93-.645h-.005a1 1 0 0 0-1 .995c-.004.815.81 1.338 1.141 1.514a4.44 4.44 0 0 1 .924 1.36c.365 1.023 1.423 2.576 4.466 2.376l.003.098l.004.268a1 1 0 0 0 2 0l-.005-.318c-.005-.19-.012-.464-.012-1.182M20.737 5.377a5.39 5.39 0 0 0 .09-.42a6.278 6.278 0 0 0-.408-3.293a1.002 1.002 0 0 0-.615-.58c-.356-.12-1.67-.357-4.184 1.25a13.87 13.87 0 0 0-6.354 0C6.762.75 5.455.966 5.102 1.079a.997.997 0 0 0-.631.584a6.3 6.3 0 0 0-.404 3.357c.025.127.051.246.079.354a6.27 6.27 0 0 0-1.256 3.83a8.422 8.422 0 0 0 .043.921c.334 4.603 3.334 5.984 5.424 6.459a4.591 4.591 0 0 0-.118.4a1 1 0 0 0 1.942.479a1.678 1.678 0 0 1 .468-.878a1 1 0 0 0-.546-1.745c-3.454-.395-4.954-1.802-5.18-4.899a6.61 6.61 0 0 1-.033-.738a4.258 4.258 0 0 1 .92-2.713a3.022 3.022 0 0 1 .195-.231a1 1 0 0 0 .188-1.025a3.388 3.388 0 0 1-.155-.555a4.094 4.094 0 0 1 .079-1.616a7.543 7.543 0 0 1 2.415 1.18a1.009 1.009 0 0 0 .827.133a11.777 11.777 0 0 1 6.173.001a1.005 1.005 0 0 0 .83-.138a7.572 7.572 0 0 1 2.406-1.19a4.04 4.04 0 0 1 .087 1.578a3.205 3.205 0 0 1-.169.607a1 1 0 0 0 .188 1.025c.078.087.155.18.224.268A4.122 4.122 0 0 1 20 9.203a7.039 7.039 0 0 1-.038.777c-.22 3.056-1.725 4.464-5.195 4.86a1 1 0 0 0-.546 1.746a1.63 1.63 0 0 1 .466.908a3.06 3.06 0 0 1 .093.82v2.333c-.01.648-.01 1.133-.01 1.356a1 1 0 1 0 2 0c0-.217 0-.692.01-1.34v-2.35a4.882 4.882 0 0 0-.155-1.311a4.256 4.256 0 0 0-.116-.416a6.513 6.513 0 0 0 5.445-6.424A8.697 8.697 0 0 0 22 9.203a6.13 6.13 0 0 0-1.263-3.826"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gitlab(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.94 12.865l-1.066-3.28l.001.005v-.005l-2.115-6.51a.833.833 0 0 0-.799-.57a.822.822 0 0 0-.788.576l-2.007 6.178H8.834L6.824 3.08a.822.822 0 0 0-.788-.575H6.03a.839.839 0 0 0-.796.575L3.127 9.584l-.002.006l.001-.005l-1.069 3.28a1.195 1.195 0 0 0 .435 1.34l9.229 6.705l.004.003l.012.008l-.011-.008l.002.001l.001.001a.466.466 0 0 0 .045.028l.006.004l.004.002l.003.001h.002l.005.003l.025.01l.023.01h.001l.004.002l.005.002h.002l.006.002h.003c.011.004.022.006.034.009l.013.003h.002l.005.002l.007.001h.007a.467.467 0 0 0 .066.006h.001a.469.469 0 0 0 .067-.005h.007l.007-.002l.004-.001h.002l.014-.004l.034-.008h.002l.006-.003h.002l.005-.002l.004-.001h.001l.025-.011l.023-.01l.005-.002h.002l.003-.002l.004-.002l.007-.004a.468.468 0 0 0 .044-.027l.004-.003l.005-.003l9.23-6.706a1.195 1.195 0 0 0 .434-1.339Zm-3.973-9.18l1.81 5.574h-3.62Zm-11.937 0L7.843 9.26h-3.62Zm-2.984 9.757a.255.255 0 0 1-.092-.285l.794-2.438l5.822 7.464Zm1.494-3.24h3.61l2.573 7.927Zm7.165 10.696l-.006-.005l-.011-.01l-.02-.018l.002.001l.002.002a.473.473 0 0 0 .043.037l.002.002Zm.293-1.894l-1.514-4.665l-1.344-4.138h5.72Zm.31 1.88l-.01.008l-.002.001l-.005.005l-.012.009l.002-.002a.455.455 0 0 0 .043-.036l.001-.002l.002-.002ZM15.851 10.2h3.61l-.74.947l-5.447 6.98Zm5.1 3.241l-6.523 4.74l5.824-7.463l.791 2.437a.255.255 0 0 1-.092.286"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Glass(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.75 2.33A1 1 0 0 0 19 2H5a1 1 0 0 0-.75.33a1 1 0 0 0-.25.78l1.8 16.22a3 3 0 0 0 3 2.67h6.42a3 3 0 0 0 3-2.67L20 3.11a1 1 0 0 0-.25-.78M16.2 19.11a1 1 0 0 1-1 .89H8.79a1 1 0 0 1-1-.89L6.78 10h10.44ZM17.44 8H6.56l-.44-4h11.76Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlassMartini(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 20h-5v-5.06A9 9 0 0 0 21 6a8.72 8.72 0 0 0-.67-3.39a1 1 0 0 0-.22-.32L20 2.21a.92.92 0 0 0-.21-.13a.94.94 0 0 0-.28-.08H4.5a.94.94 0 0 0-.29.06A2.12 2.12 0 0 0 4 2.2l-.12.09a1 1 0 0 0-.22.32A8.72 8.72 0 0 0 3 6a9 9 0 0 0 8 8.94V20H6a1 1 0 0 0 0 2h12a1 1 0 0 0 0-2M5 6a6.91 6.91 0 0 1 .29-2h13.42A6.91 6.91 0 0 1 19 6A7 7 0 0 1 5 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlassMartiniAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.78 3.62a1 1 0 0 0 .12-1.05A1 1 0 0 0 21 2H3a1 1 0 0 0-.9.57a1 1 0 0 0 .12 1.05L11 14.6V20H5.25a1 1 0 0 0 0 2h13.5a1 1 0 0 0 0-2H13v-5.4ZM5.08 4h13.84l-2.4 3h-9ZM12 12.65L9.08 9h5.84Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlassMartiniAltSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m19.71 16.29l-14-14a1 1 0 0 0-1.42 1.42L6.59 6H5a1 1 0 0 0-.9.57a1 1 0 0 0 .12 1L11 16.1V20H6.75a1 1 0 0 0 0 2h10.5a1 1 0 0 0 0-2H13v-3.9l1.64-2l3.65 3.65a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.46M7.08 8h1.51l1.89 1.89H8.59ZM12 14.15l-1.81-2.26h2.29l.74.74ZM14.66 8h2.26l-.63.79a1 1 0 0 0 .15 1.4a1 1 0 0 0 .63.22a1 1 0 0 0 .78-.37l1.93-2.42a1 1 0 0 0 .12-1A1 1 0 0 0 19 6h-4.34a1 1 0 0 0 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlassTea(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 3a3 3 0 0 0-2.23-1H7.23a3 3 0 0 0-3 3.33l1.56 14a3 3 0 0 0 3 2.67h6.42a3 3 0 0 0 3-2.67l1.56-14A3 3 0 0 0 19 3m-2.8 16.11a1 1 0 0 1-1 .89H8.79a1 1 0 0 1-1-.89L6.78 10h10.44ZM17.44 8H6.56l-.32-2.89a1 1 0 0 1 .25-.78A1 1 0 0 1 7.23 4h9.54a1 1 0 0 1 .74.33a1 1 0 0 1 .25.78ZM14 18a1 1 0 0 0 1-1v-4a1 1 0 0 0-2 0v4a1 1 0 0 0 1 1m-4 0a1 1 0 0 0 1-1v-4a1 1 0 0 0-2 0v4a1 1 0 0 0 1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Globe(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.41 8.64v-.05a10 10 0 0 0-18.78 0s0 0 0 .05a9.86 9.86 0 0 0 0 6.72v.05a10 10 0 0 0 18.78 0s0 0 0-.05a9.86 9.86 0 0 0 0-6.72M4.26 14a7.82 7.82 0 0 1 0-4h1.86a16.73 16.73 0 0 0 0 4Zm.82 2h1.4a12.15 12.15 0 0 0 1 2.57A8 8 0 0 1 5.08 16m1.4-8h-1.4a8 8 0 0 1 2.37-2.57A12.15 12.15 0 0 0 6.48 8M11 19.7A6.34 6.34 0 0 1 8.57 16H11Zm0-5.7H8.14a14.36 14.36 0 0 1 0-4H11Zm0-6H8.57A6.34 6.34 0 0 1 11 4.3Zm7.92 0h-1.4a12.15 12.15 0 0 0-1-2.57A8 8 0 0 1 18.92 8M13 4.3A6.34 6.34 0 0 1 15.43 8H13Zm0 15.4V16h2.43A6.34 6.34 0 0 1 13 19.7m2.86-5.7H13v-4h2.86a14.36 14.36 0 0 1 0 4m.69 4.57a12.15 12.15 0 0 0 1-2.57h1.4a8 8 0 0 1-2.4 2.57M19.74 14h-1.86a16.16 16.16 0 0 0 .12-2a16.28 16.28 0 0 0-.12-2h1.86a7.82 7.82 0 0 1 0 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gold(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 11h8a1 1 0 0 0 .77-.37A1 1 0 0 0 17 9.8l-1-5a1 1 0 0 0-1-.8H9a1 1 0 0 0-1 .8l-1 5a1 1 0 0 0 .21.83A1 1 0 0 0 8 11m1.82-5h4.36l.6 3H9.22ZM22 13.8a1 1 0 0 0-1-.8h-6a1 1 0 0 0-1 .8l-1 5a1 1 0 0 0 .21.83A1 1 0 0 0 14 20h8a1 1 0 0 0 .77-.37a1 1 0 0 0 .23-.83ZM15.22 18l.6-3h4.36l.6 3ZM9 13H3a1 1 0 0 0-1 .8l-1 5a1 1 0 0 0 .21.83A1 1 0 0 0 2 20h8a1 1 0 0 0 .77-.37a1 1 0 0 0 .23-.83l-1-5a1 1 0 0 0-1-.8m-5.78 5l.6-3h4.36l.6 3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GolfBall(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 9a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0-3a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-2-4a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8m5-12a1 1 0 1 0 1 1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Google(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.602 10.004a1 1 0 0 0-.984-.822H12.2a1 1 0 0 0-1 1v3.868a1 1 0 0 0 1 1h3.962a3.652 3.652 0 0 1-1.131 1.187a5.06 5.06 0 0 1-2.831.785a4.935 4.935 0 0 1-4.646-3.437v-.002a4.904 4.904 0 0 1 0-3.167v-.002A4.936 4.936 0 0 1 12.2 6.978a4.378 4.378 0 0 1 3.131 1.217a1 1 0 0 0 1.399-.015l2.868-2.868a1 1 0 0 0-.025-1.439A10.623 10.623 0 0 0 12.2 1a10.949 10.949 0 0 0-9.829 6.059l-.001.002A10.922 10.922 0 0 0 1.2 12a11.079 11.079 0 0 0 1.17 4.94l.001.001A10.949 10.949 0 0 0 12.2 23a10.526 10.526 0 0 0 7.295-2.687l.001-.001a10.786 10.786 0 0 0 3.304-8.084a12.515 12.515 0 0 0-.198-2.224M12.2 3a8.682 8.682 0 0 1 5.209 1.673l-1.454 1.453A6.463 6.463 0 0 0 12.2 4.978a6.886 6.886 0 0 0-5.99 3.55L5.142 7.7l-.585-.454A8.953 8.953 0 0 1 12.2 3M3.68 14.903a9.03 9.03 0 0 1 0-5.806l1.782 1.382a6.854 6.854 0 0 0 0 3.042ZM12.2 21a8.953 8.953 0 0 1-7.644-4.246l.379-.294l1.276-.99a6.885 6.885 0 0 0 5.989 3.552a7.277 7.277 0 0 0 3.306-.75l1.691 1.313A8.89 8.89 0 0 1 12.2 21m6.526-2.76l-.183-.143l-1.377-1.069a5.606 5.606 0 0 0 1.4-2.796a1 1 0 0 0-.984-1.182H13.2v-1.868h7.549c.034.345.051.695.051 1.046a9.052 9.052 0 0 1-2.074 6.011"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleDrive(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.423 13.888l-6.09-10.55H8.668l6.09 10.55ZM8.09 4.338L2 14.888l3.334 5.774l6.089-10.55Zm1.733 10.55L6.49 20.662h12.178L22 14.887Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleDriveAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 14.44a.62.62 0 0 0 0-.13a.61.61 0 0 1 0-.12l-.05-.12l-6-10.29a1 1 0 0 0-.95-.49H9a1 1 0 0 0-.5.13l-.11.08a.73.73 0 0 0-.09.08a.58.58 0 0 0-.1.12s0 0-.06.08l-6 10.33a1 1 0 0 0 0 1l3 5.08a.83.83 0 0 0 .11.15v.06a1.1 1.1 0 0 0 .44.26a.83.83 0 0 0 .22 0H18a1 1 0 0 0 .86-.49l3-5.14l.05-.12a.61.61 0 0 1 0-.12a.53.53 0 0 0 0-.13a.51.51 0 0 0 0-.13a.59.59 0 0 0 .09-.09M6 17.73l-1.79-3.1L9 6.27l.87 1.5l1 1.66L7 15.91Zm6-6.32l1.26 2.16h-2.54Zm5.43 7.3H7.7l1.84-3.14h9.72Zm-1.86-5.14l-4.83-8.28h3.69l4.83 8.28Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleHangouts(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.444 9.055a8.505 8.505 0 0 0-2.173-4.334a8.403 8.403 0 0 0-1.622-1.357a8.489 8.489 0 0 0-3.64-1.31A4.153 4.153 0 0 1 12.528 2h-1.135c-.012.029-.038.018-.06.02a8.835 8.835 0 0 0-.87.114a8.453 8.453 0 0 0-5.177 3.104a8.359 8.359 0 0 0-1.84 4.709a8.59 8.59 0 0 0 .186 2.529a8.148 8.148 0 0 0 .624 1.79l.073.146a8.601 8.601 0 0 0 1.784 2.396a8.53 8.53 0 0 0 5.763 2.333c.1.001.127.026.126.128c-.004.874-.002 1.747-.002 2.621c0 .034.003.068.005.11l.063-.026a17.989 17.989 0 0 0 4.49-2.966q.4-.365.769-.762c.105-.114.21-.23.315-.342c.163-.175.304-.37.458-.553a9.09 9.09 0 0 0 .494-.663a11.033 11.033 0 0 0 .918-1.575c.132-.279.258-.56.37-.847a.096.096 0 0 0 .01-.016a10.277 10.277 0 0 0 .6-2.44c.037-.29.06-.584.075-.877a8.14 8.14 0 0 0-.123-1.878m-9.19 2.82a2.505 2.505 0 0 1-.84 1.877l-.084.08a2.684 2.684 0 0 1-.933.481a3.8 3.8 0 0 1-.448.085a.21.21 0 0 1-.235-.152l.001-.968c.006-.095-.032-.197.044-.28a.266.266 0 0 1 .138-.085a1.774 1.774 0 0 0 .494-.169a1.132 1.132 0 0 0 .538-.731l.036-.151l-2.042-.003a.653.653 0 0 1-.174-.018a.413.413 0 0 1-.307-.386q-.002-1.543.002-3.087a.467.467 0 0 1 .138-.436a.387.387 0 0 1 .257-.102h3.001a.416.416 0 0 1 .4.527a.388.388 0 0 1 .016.153c0 1.121.007 2.243-.002 3.364m5.303.004a2.557 2.557 0 0 1-1.556 2.328c-.024.01-.046.023-.069.035l-.038.01l-.021.013a4.592 4.592 0 0 1-.59.129a.216.216 0 0 1-.269-.217l-.003-.322l.006-.578c-.001-.042-.003-.084-.003-.127c0-.147.057-.21.241-.252a1.319 1.319 0 0 0 .634-.297a1.264 1.264 0 0 0 .377-.74l-1.885-.003h-.206a.419.419 0 0 1-.432-.44q0-1.525.002-3.049l-.001-.03a.434.434 0 0 1 .203-.45a.39.39 0 0 1 .213-.059h2.982a.416.416 0 0 1 .399.527a.392.392 0 0 1 .017.153z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleHangoutsAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.992 7.69a2 2 0 0 0 0 4l.015-.001v.501a.501.501 0 0 1-.5.5a1 1 0 0 0 0 2a2.502 2.502 0 0 0 2.5-2.5v-2.5a.941.941 0 0 0-.03-.15a1.994 1.994 0 0 0-1.985-1.85m3-6.688a9.787 9.787 0 0 0-1 19.523v1.477a1 1 0 0 0 1.238.97a12.535 12.535 0 0 0 9.467-10.974a9.734 9.734 0 0 0-9.706-10.996m7.717 10.78a10.548 10.548 0 0 1-6.718 8.86v-1.066a1 1 0 0 0-1-1a7.787 7.787 0 1 1 7.788-7.787a7.945 7.945 0 0 1-.07.993M14.992 7.69a2 2 0 1 0 0 4l.015-.001v.501a.501.501 0 0 1-.5.5a1 1 0 0 0 0 2a2.502 2.502 0 0 0 2.5-2.5v-2.5a.941.941 0 0 0-.03-.15a1.994 1.994 0 0 0-1.985-1.85"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GooglePlay(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.919 10.653a503.26 503.26 0 0 1-4.023-2.323l-.003-.002L4.64 1.253a1.679 1.679 0 0 0-1.408-.16a.953.953 0 0 0-.076.03a1.418 1.418 0 0 0-.173.07a1.519 1.519 0 0 0-.738 1.364v18.986a1.435 1.435 0 0 0 .692 1.27a1.308 1.308 0 0 0 .155.064a.977.977 0 0 0 .087.035a1.379 1.379 0 0 0 .446.083a1.673 1.673 0 0 0 .831-.232l12.438-7.182l4.021-2.322a1.525 1.525 0 0 0 .842-1.334a1.49 1.49 0 0 0-.837-1.272M4.244 19.839V4.102l7.94 7.859Zm5.018-2.162l4.344-4.31l1.15 1.139zm4.342-7.125L9.206 6.2l5.554 3.207Zm2.947 2.917l-1.526-1.51l1.528-1.516c.72.419 1.843 1.07 2.616 1.515Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GraduationCap(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.49 10.19l-1-.55l-9-5h-.11a1.06 1.06 0 0 0-.19-.06h-.37a1.17 1.17 0 0 0-.2.06h-.11l-9 5a1 1 0 0 0 0 1.74L4 12.76v4.74a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3v-4.74l2-1.12v2.86a1 1 0 0 0 2 0v-3.44a1 1 0 0 0-.51-.87M16 17.5a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1v-3.63l4.51 2.5l.15.06h.09a1 1 0 0 0 .25 0a1 1 0 0 0 .25 0h.09a.47.47 0 0 0 .15-.06l4.51-2.5Zm-5-3.14L4.06 10.5L11 6.64l6.94 3.86Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GraphBar(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 13H2a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1m-1 8H3v-6h2ZM22 9h-4a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1V10a1 1 0 0 0-1-1m-1 12h-2V11h2ZM14 1h-4a1 1 0 0 0-1 1v20a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1m-1 20h-2V3h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Grid(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1M11 20H4v-4h7Zm0-6H4v-4h7Zm9 6h-7v-4h7Zm0-6h-7v-4h7Zm0-6H4V4h16Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Grids(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1M8 20H4V4h4Zm6 0h-4V4h4Zm6 0h-4V4h4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Grin(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 11a1 1 0 1 0-1-1a1 1 0 0 0 1 1m3-9a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8m3-7H9a1 1 0 0 0-1 1a4 4 0 0 0 8 0a1 1 0 0 0-1-1m-3 3a2 2 0 0 1-1.73-1h3.46A2 2 0 0 1 12 16m3-7a1 1 0 1 0 1 1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GrinTongueWink(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8m5.62-10.87a3.08 3.08 0 0 0-4.24 0a1 1 0 0 0 1.41 1.41a1 1 0 0 1 1.42 0a1 1 0 0 0 1.41 0a1 1 0 0 0 0-1.41M9 11a1 1 0 1 0-1-1a1 1 0 0 0 1 1m6 2H9a1 1 0 0 0 0 2a3 3 0 0 0 6 0a1 1 0 0 0 0-2m-3 3a1 1 0 0 1-1-1h2a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GrinTongueWinkAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.21 10.54a1 1 0 0 0 1.41 0a1 1 0 0 0 0-1.41a3.08 3.08 0 0 0-4.24 0a1 1 0 1 0 1.41 1.41a1 1 0 0 1 1.42 0M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8m3-11a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0 4H9a1 1 0 0 0 0 2a3 3 0 0 0 6 0a1 1 0 0 0 0-2m-3 3a1 1 0 0 1-1-1h2a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GripHorizontalLine(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 11h18a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2m18 2H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HardHat(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 13.17V13a8 8 0 0 0-2.42-5.74a7.84 7.84 0 0 0-3.18-1.88h-.05A8.24 8.24 0 0 0 11.76 5A8.21 8.21 0 0 0 4 13.17A3 3 0 0 0 5 19h14a3 3 0 0 0 1-5.83M19 17H5a1 1 0 0 1 0-2h2a1 1 0 0 0 0-2H6a6.41 6.41 0 0 1 3-5.15V11a1 1 0 0 0 2 0V7.09a7.34 7.34 0 0 1 .82-.09H12a5.56 5.56 0 0 1 1 .1V11a1 1 0 0 0 2 0V7.82a6.65 6.65 0 0 1 1.18.87A6 6 0 0 1 18 13h-1a1 1 0 0 0 0 2h2a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hdd(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 14.86v-.05a2.61 2.61 0 0 0-.1-.57l-1.64-9.73a3 3 0 0 0-3-2.51H6.69a3 3 0 0 0-2.95 2.51l-1.62 9.71a2.61 2.61 0 0 0-.1.57v.05C2 14.91 2 15 2 15v4a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-4zM5.71 4.83a1 1 0 0 1 1-.83h10.6a1 1 0 0 1 1 .83l1.2 7.22A2.63 2.63 0 0 0 19 12H5a2.63 2.63 0 0 0-.49.05ZM20 19a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-3.92l.08-.46A1 1 0 0 1 5 14h14a1 1 0 0 1 .92.62l.08.46Zm-3-3a1 1 0 1 0 1 1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadSide(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.23 2.003a7.372 7.372 0 0 0-5.453 2.114A7.44 7.44 0 0 0 5.5 9.5v.03l-1.904 4.044A1 1 0 0 0 4.5 15h1v2a2.002 2.002 0 0 0 2 2h1v2a1 1 0 0 0 2 0v-2a1 1 0 0 0 0-2h-3v-3a1 1 0 0 0-1-1h-.424l1.34-2.844a.99.99 0 0 0 .095-.465L7.5 9.5a5.455 5.455 0 0 1 1.67-3.947a5.527 5.527 0 0 1 4-1.55a5.685 5.685 0 0 1 5.33 5.77l-1.967 7.504a1.01 1.01 0 0 0 .006.534l1 3.466A1.001 1.001 0 0 0 18.5 22a1.018 1.018 0 0 0 .277-.04a1 1 0 0 0 .684-1.237l-.924-3.2l1.93-7.267A1.031 1.031 0 0 0 20.5 10v-.228a7.698 7.698 0 0 0-7.27-7.769"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadSideCough(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.293 20.293a1 1 0 1 0 1.414 0a1 1 0 0 0-1.414 0m-3-3a1 1 0 1 0 1.414 0a1 1 0 0 0-1.414 0m4-1a1 1 0 1 0 1.414 0a1 1 0 0 0-1.414 0M16.15 2.002a7.067 7.067 0 0 0-7.284 7.063v.016l-1.77 3.758A1 1 0 0 0 8 14.267h.866v1.8A1.935 1.935 0 0 0 10.8 18h.867v1.8a1 1 0 0 0 2 0v-1.814A.994.994 0 0 0 13.6 16l-.833.02a.94.94 0 0 0-.1-.02a.94.94 0 0 0-.128.026l-1.673.04v-2.8a1 1 0 0 0-1-1h-.289l1.205-2.558a.99.99 0 0 0 .095-.468l-.01-.174a5.025 5.025 0 0 1 1.537-3.635a5.092 5.092 0 0 1 3.686-1.429A5.239 5.239 0 0 1 21 9.322l-1.833 6.987a1.008 1.008 0 0 0 .006.533l.932 3.235a1 1 0 0 0 .961.723a1.017 1.017 0 0 0 .278-.04a1 1 0 0 0 .683-1.237l-.856-2.97l1.796-6.763A1.031 1.031 0 0 0 23 9.533v-.212a7.252 7.252 0 0 0-6.85-7.318Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadSideMask(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.23 2.003a7.367 7.367 0 0 0-5.453 2.114A7.441 7.441 0 0 0 5.5 9.465l-1.844 2.998a.995.995 0 0 0-.156.52v.04a1 1 0 0 0 .07.347l1.44 3.873a.87.87 0 0 0 .043.099A2.984 2.984 0 0 0 7.736 19H8.5v2a1 1 0 0 0 2 0v-2h2a1 1 0 0 0 .321-.053l3.7-1.256a.999.999 0 0 0 .018.12l1 3.466A1.001 1.001 0 0 0 18.5 22a1.018 1.018 0 0 0 .277-.04a1 1 0 0 0 .684-1.237l-.924-3.2l1.93-7.267A1.031 1.031 0 0 0 20.5 10v-.228a7.698 7.698 0 0 0-7.27-7.769M11.5 17H7.736a.995.995 0 0 1-.874-.513L5.938 14H11.5Zm5.523-1.591L13.5 16.605V13.72l4.345-1.448Zm1.412-5.389a.973.973 0 0 0-.251.031L12.337 12H6.29l1.074-1.747a1 1 0 0 0 .148-.562L7.5 9.5a5.455 5.455 0 0 1 1.67-3.947a5.522 5.522 0 0 1 4-1.55a5.685 5.685 0 0 1 5.33 5.77Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadphoneSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.07 5.936a9.945 9.945 0 0 0-9.39-2.671a1 1 0 0 0 .44 1.95A8.02 8.02 0 0 1 20 13h-.34a1 1 0 0 0 0 2H20v.34a1 1 0 1 0 2 0V13a9.888 9.888 0 0 0-2.93-7.064M3.706 2.293a1 1 0 0 0-1.414 1.414l2.435 2.435A9.962 9.962 0 0 0 2 13v7a1 1 0 0 0 1 1h3a3.003 3.003 0 0 0 3-3v-2a3.003 3.003 0 0 0-3-3H4a7.963 7.963 0 0 1 2.145-5.441L15 16.414V18a3.003 3.003 0 0 0 3 3h1.586l.707.707a1 1 0 0 0 1.414-1.414ZM6 15a1.001 1.001 0 0 1 1 1v2a1.001 1.001 0 0 1-1 1H4v-4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Headphones(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 3A10 10 0 0 0 2 13v7a1 1 0 0 0 1 1h3a3 3 0 0 0 3-3v-2a3 3 0 0 0-3-3H4a8 8 0 0 1 16 0h-2a3 3 0 0 0-3 3v2a3 3 0 0 0 3 3h3a1 1 0 0 0 1-1v-7A10 10 0 0 0 12 3M6 15a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1H4v-4Zm14 4h-2a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadphonesAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 13.18V11a8 8 0 0 0-16 0v2.18A3 3 0 0 0 2 16v2a3 3 0 0 0 3 3h3a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1H6v-2a6 6 0 0 1 12 0v2h-2a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h3a3 3 0 0 0 3-3v-2a3 3 0 0 0-2-2.82M7 15v4H5a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1Zm13 3a1 1 0 0 1-1 1h-2v-4h2a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heart(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.16 5A6.29 6.29 0 0 0 12 4.36a6.27 6.27 0 0 0-8.16 9.48l6.21 6.22a2.78 2.78 0 0 0 3.9 0l6.21-6.22a6.27 6.27 0 0 0 0-8.84m-1.41 7.46l-6.21 6.21a.76.76 0 0 1-1.08 0l-6.21-6.24a4.29 4.29 0 0 1 0-6a4.27 4.27 0 0 1 6 0a1 1 0 0 0 1.42 0a4.27 4.27 0 0 1 6 0a4.29 4.29 0 0 1 .08 6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.16 4.61A6.27 6.27 0 0 0 12 4a6.27 6.27 0 0 0-8.16 9.48l7.45 7.45a1 1 0 0 0 1.42 0l7.45-7.45a6.27 6.27 0 0 0 0-8.87m-1.41 7.46L12 18.81l-6.75-6.74a4.28 4.28 0 0 1 3-7.3a4.25 4.25 0 0 1 3 1.25a1 1 0 0 0 1.42 0a4.27 4.27 0 0 1 6 6.05Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartBreak(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.16 4.61A6.27 6.27 0 0 0 12 4a6.27 6.27 0 0 0-8.16 9.48l7.45 7.46a1 1 0 0 0 1.42 0l7.45-7.46a6.27 6.27 0 0 0 0-8.87m-1.41 7.45L12 18.81l-6.75-6.75a4.26 4.26 0 0 1 5.54-6.45l-1.71 4a1 1 0 0 0 0 .83a1 1 0 0 0 .65.53l2.77.7l-1.4 2.89a1 1 0 0 0 .46 1.34a1 1 0 0 0 .44.1a1 1 0 0 0 .9-.56l2-4a1 1 0 0 0 0-.86a1.05 1.05 0 0 0-.67-.55l-2.83-.71l1.45-3.39a4.26 4.26 0 0 1 5.92 6.13Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartMedical(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 11h-1v-1a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2h1v1a1 1 0 0 0 2 0v-1h1a1 1 0 0 0 0-2m6.16-6A6.29 6.29 0 0 0 12 4.41a6.27 6.27 0 0 0-8.16 9.48l6 6.05a3 3 0 0 0 4.24 0l6-6.05A6.27 6.27 0 0 0 20.16 5m-1.41 7.46l-6 6a1 1 0 0 1-1.42 0l-6-6a4.29 4.29 0 0 1 0-6a4.27 4.27 0 0 1 6 0a1 1 0 0 0 1.42 0a4.27 4.27 0 0 1 6 0a4.29 4.29 0 0 1 0 6.02Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartRate(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 11h-3.94a.78.78 0 0 0-.21 0h-.17a1.3 1.3 0 0 0-.15.1a1.67 1.67 0 0 0-.16.12a1 1 0 0 0-.09.13a1.32 1.32 0 0 0-.12.2l-1.6 4.41l-4.17-11.3a1 1 0 0 0-1.88 0L6.2 11H3a1 1 0 0 0 0 2h4.3a.86.86 0 0 0 .16-.1a1.67 1.67 0 0 0 .16-.12l.09-.13a1 1 0 0 0 .12-.2l1.62-4.53l4.16 11.42a1 1 0 0 0 .94.66a1 1 0 0 0 .94-.66l2.3-6.34H21a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartSign(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.16 5A6.29 6.29 0 0 0 12 4.41a6.27 6.27 0 0 0-8.16 9.48l6 6.05a3 3 0 0 0 4.24 0l6-6.05A6.27 6.27 0 0 0 20.16 5m-1.41 7.46l-6 6a1 1 0 0 1-1.42 0l-6-6a4.29 4.29 0 0 1 0-6a4.27 4.27 0 0 1 6 0a1 1 0 0 0 1.42 0a4.27 4.27 0 0 1 6 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heartbeat(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 10.41h-2.5a1 1 0 0 0-.71.3L16.55 12l-2.8-3.19a1 1 0 0 0-1.46 0l-1.7 1.7H9a1 1 0 0 0 0 2h2a1 1 0 0 0 .71-.29L13 10.88l2.8 3.19a1 1 0 0 0 .72.34a1 1 0 0 0 .71-.29l1.7-1.71H21a1 1 0 0 0 0-2m-7.39 5.3l-1.9 1.9a1 1 0 0 1-1.42 0L5.08 12.4a3.69 3.69 0 0 1 0-5.22a3.69 3.69 0 0 1 5.21 0a1 1 0 0 0 1.42 0a3.78 3.78 0 0 1 5.21 0a3.94 3.94 0 0 1 .58.75a1 1 0 0 0 1.72-1a6 6 0 0 0-.88-1.13A5.68 5.68 0 0 0 11 5.17a5.68 5.68 0 0 0-9 4.62a5.62 5.62 0 0 0 1.67 4L8.88 19a3 3 0 0 0 4.24 0L15 17.12a1 1 0 0 0 0-1.41a1 1 0 0 0-1.39 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HindiToChinese(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.022 2h-2a1 1 0 0 0 0 2v2H7.838a2.965 2.965 0 0 0 .184-1a3 3 0 0 0-5.598-1.5a1 1 0 1 0 1.732 1a1.002 1.002 0 0 1 .866-.5a1 1 0 0 1 0 2a1 1 0 0 0 0 2a1 1 0 0 1 0 2a1.002 1.002 0 0 1-.866-.5a1 1 0 1 0-1.732 1A3 3 0 0 0 8.022 9a2.965 2.965 0 0 0-.184-1h1.184v3a1 1 0 0 0 2 0V4a1 1 0 0 0 0-2m3 5h1a1.001 1.001 0 0 1 1 1v1a1 1 0 0 0 2 0V8a3.003 3.003 0 0 0-3-3h-1a1 1 0 0 0 0 2m-4 9h-1a1.001 1.001 0 0 1-1-1v-1a1 1 0 0 0-2 0v1a3.003 3.003 0 0 0 3 3h1a1 1 0 0 0 0-2m11-1a1 1 0 0 0 0-2h-3v-.5a1 1 0 0 0-2 0v.5h-3a1 1 0 0 0 0 2h5.184a6.728 6.728 0 0 1-1.225 2.527a6.668 6.668 0 0 1-.63-.983a1 1 0 1 0-1.779.912a8.678 8.678 0 0 0 .96 1.468a6.618 6.618 0 0 1-2.426 1.099a1 1 0 0 0 .427 1.954a8.635 8.635 0 0 0 3.445-1.622a8.724 8.724 0 0 0 3.469 1.622a1 1 0 1 0 .43-1.954a6.725 6.725 0 0 1-2.446-1.09A8.736 8.736 0 0 0 20.244 15Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hipchat(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.5 13.9c0-.2-.2-.4-.4-.4c-.1 0-.2 0-.3.1c-1.3 1.1-3 1.7-4.8 1.7c-1.7 0-3.4-.6-4.8-1.7c-.1-.1-.2-.1-.3-.1c-.2 0-.4.2-.4.4s.1.3.2.5C8 16 10 16.9 12 16.8c2 .1 3.9-.8 5.2-2.4c.2-.2.2-.3.3-.5m2.7 6.9c-.8-.4-1.4-1.2-1.7-2c-.1-.2 0-.3.1-.4c2.1-1.5 3.3-3.9 3.3-6.4c0-4.8-4.5-8.6-10-8.6s-10 3.9-10 8.6c0 4.8 4.5 8.6 10 8.6c.7 0 1.4-.1 2.1-.2c.1 0 .3 0 .4.1c1.6.9 3.4 1.5 5.2 1.5c.4.1.7-.2.8-.6v-.1c.1-.2 0-.4-.2-.5m-2.7-.8v.1c0 .1-.1.1-.1.1c-1.1-.3-2-.8-2.9-1.5c-.1-.1-.3-.1-.5-.1c-.7.2-1.4.2-2.1.2c-4.6 0-8.3-3.1-8.3-6.9C3.6 8.1 7.3 5 11.9 5c4.6 0 8.3 3.1 8.3 6.9c-.1 2.3-1.4 4.4-3.4 5.5c-.2.1-.3.3-.3.4c.3.9.6 1.6 1 2.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func History(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a10 10 0 0 0-6.88 2.77V3a1 1 0 0 0-2 0v4.5a1 1 0 0 0 1 1h4.5a1 1 0 0 0 0-2h-2.4A8 8 0 1 1 4 12a1 1 0 0 0-2 0A10 10 0 1 0 12 2m0 6a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h2a1 1 0 0 0 0-2h-1V9a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HistoryAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.44 2a10 10 0 0 0-6.88 2.77V3a1 1 0 0 0-2 0v4.5a1 1 0 0 0 1 1h4.5a1 1 0 0 0 0-2h-2.4A8 8 0 1 1 11.44 20a1 1 0 1 0 0 2a10 10 0 1 0 0-20m0 6a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h2a1 1 0 0 0 0-2h-1V9a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Home(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.66 10.25l-9-8a1 1 0 0 0-1.32 0l-9 8a1 1 0 0 0-.27 1.11A1 1 0 0 0 3 12h1v9a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1v-9h1a1 1 0 0 0 .93-.64a1 1 0 0 0-.27-1.11M13 20h-2v-3a1 1 0 0 1 2 0Zm5 0h-3v-3a3 3 0 0 0-6 0v3H6v-8h12ZM5.63 10L12 4.34L18.37 10Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20 8l-6-5.26a3 3 0 0 0-4 0L4 8a3 3 0 0 0-1 2.26V19a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-8.75A3 3 0 0 0 20 8m-6 12h-4v-5a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1Zm5-1a1 1 0 0 1-1 1h-2v-5a3 3 0 0 0-3-3h-2a3 3 0 0 0-3 3v5H6a1 1 0 0 1-1-1v-8.75a1 1 0 0 1 .34-.75l6-5.25a1 1 0 0 1 1.32 0l6 5.25a1 1 0 0 1 .34.75Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HorizontalAlignCenter(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 10h-2V7a1 1 0 0 0-1-1h-5V3a1 1 0 0 0-2 0v3H6a1 1 0 0 0-1 1v3H3a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h8v3a1 1 0 0 0 2 0v-3h8a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1M7 8h10v2H7Zm13 8H4v-4h16Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HorizontalAlignLeft(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 10h-5V7a1 1 0 0 0-1-1H4V3a1 1 0 0 0-2 0v18a1 1 0 0 0 2 0v-3h17a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1M4 8h10v2H4Zm16 8H4v-4h16Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HorizontalAlignRight(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2a1 1 0 0 0-1 1v3H9a1 1 0 0 0-1 1v3H3a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h17v3a1 1 0 0 0 2 0V3a1 1 0 0 0-1-1m-1 14H4v-4h16Zm0-6H10V8h10Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HorizontalDistributionCenter(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 5h-1V3a1 1 0 0 0-2 0v2h-1a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h1v2a1 1 0 0 0 2 0v-2h1a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1m-1 12h-2V7h2ZM11 4H9V3a1 1 0 0 0-2 0v1H5a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h2v1a1 1 0 0 0 2 0v-1h2a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1m-1 14H6V6h4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HorizontalDistributionLeft(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 4H6V3a1 1 0 0 0-2 0v18a1 1 0 0 0 2 0v-1h5a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1m-1 14H6V6h4Zm9-13h-3V3a1 1 0 0 0-2 0v18a1 1 0 0 0 2 0v-2h3a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1m-1 12h-2V7h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HorizontalDistributionRight(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 2a1 1 0 0 0-1 1v1h-5a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h5v1a1 1 0 0 0 2 0V3a1 1 0 0 0-1-1m-1 16h-4V6h4ZM9 2a1 1 0 0 0-1 1v2H5a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h3v2a1 1 0 0 0 2 0V3a1 1 0 0 0-1-1M8 17H6V7h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hospital(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 16.5h-1a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2m0-4h-1a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2m-5 4h-1a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2m0-4h-1a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2m14-6h-3v-4a1 1 0 0 0-1-1h-11a1 1 0 0 0-1 1v4h-3a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h19a1 1 0 0 0 1-1v-14a1 1 0 0 0-1-1m-1 14h-17v-12h3a1 1 0 0 0 1-1v-4h9v4a1 1 0 0 0 1 1h3Zm-4-8a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2Zm-3-5.5H13v-.5a1 1 0 0 0-2 0V7h-.5a1 1 0 0 0 0 2h.5v.5a1 1 0 0 0 2 0V9h.5a1 1 0 0 0 0-2m4 9.5h-1a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HospitalSquareSign(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 6a1 1 0 0 0-1 1v4h-4V7a1 1 0 0 0-2 0v10a1 1 0 0 0 2 0v-4h4v4a1 1 0 0 0 2 0V7a1 1 0 0 0-1-1m4-4H5a3.003 3.003 0 0 0-3 3v14a3.003 3.003 0 0 0 3 3h14a3.003 3.003 0 0 0 3-3V5a3.003 3.003 0 0 0-3-3m1 17a1.001 1.001 0 0 1-1 1H5a1.001 1.001 0 0 1-1-1V5a1.001 1.001 0 0 1 1-1h14a1.001 1.001 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HospitalSymbol(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 7a1 1 0 0 0-1 1v3h-4V8a1 1 0 0 0-2 0v8a1 1 0 0 0 2 0v-3h4v3a1 1 0 0 0 2 0V8a1 1 0 0 0-1-1m-3-5a10 10 0 1 0 10 10A10.011 10.011 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8.01 8.01 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hourglass(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.992 21.002h-1v-1.667a5 5 0 0 0-.3-1.678a.969.969 0 0 0-.036-.084a4.977 4.977 0 0 0-.664-1.237l-1.4-1.867a3.02 3.02 0 0 1-.6-1.801v-1.01a3.021 3.021 0 0 1 .878-2.12l.657-.658a4.946 4.946 0 0 0 1.397-2.839c0-.013.008-.025.008-.04l-.003-.013a5.018 5.018 0 0 0 .063-.643V3.002h1a1 1 0 0 0 0-2h-14a1 1 0 0 0 0 2h1v2.343a5.018 5.018 0 0 0 .063.643l-.003.014c0 .014.007.026.008.04A4.946 4.946 0 0 0 7.456 8.88l.657.657a3.021 3.021 0 0 1 .879 2.121v1.01a3.022 3.022 0 0 1-.6 1.8l-1.4 1.868a4.982 4.982 0 0 0-.665 1.237a.976.976 0 0 0-.036.084a5.003 5.003 0 0 0-.3 1.678v1.667h-1a1 1 0 0 0 0 2h14a1 1 0 0 0 0-2m-11-16v-2h8v2Zm.878 2.464a2.97 2.97 0 0 1-.377-.464h6.997a2.97 2.97 0 0 1-.377.464l-.657.657a4.96 4.96 0 0 0-1.422 2.879H10.95a4.96 4.96 0 0 0-1.422-2.879Zm1.122 8.202a5.037 5.037 0 0 0 .988-2.666h2.023a5.033 5.033 0 0 0 .989 2.666l1 1.334h-6Zm6 5.334h-8v-1.667a2.954 2.954 0 0 1 .027-.333h7.945a2.954 2.954 0 0 1 .028.333Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HouseUser(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.664 10.252l-9-8a.999.999 0 0 0-1.328 0l-9 8a1 1 0 0 0 1.328 1.496L4 11.449V21a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1v-9.551l.336.299a1 1 0 0 0 1.328-1.496M9.184 20a2.982 2.982 0 0 1 5.632 0Zm1.316-5.5A1.5 1.5 0 1 1 12 16a1.502 1.502 0 0 1-1.5-1.5M18 20h-1.101a5 5 0 0 0-2.259-3.228a3.468 3.468 0 0 0 .86-2.272a3.5 3.5 0 0 0-7 0a3.468 3.468 0 0 0 .86 2.272A5 5 0 0 0 7.1 20H6V9.671l6-5.333l6 5.333Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HtmlFive(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m3.183 2l1.604 18l7.202 2l7.222-2.002L20.818 2Zm14.142 5.887H8.877l.202 2.261h8.045l-.606 6.778L12 18.178l-.01.004l-4.522-1.256l-.31-3.466h2.216l.157 1.76l2.46.664h.001l2.463-.665l.256-2.863H7.06L6.464 5.68h11.059Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HtmlFiveAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.468 2.325A1 1 0 0 0 20.73 2H3.27a1 1 0 0 0-.996 1.089l1.52 17a1 1 0 0 0 .728.874l7.2 2a.996.996 0 0 0 .268.037a1.012 1.012 0 0 0 .267-.036l7.22-2a1 1 0 0 0 .729-.875l1.52-17a1 1 0 0 0-.258-.764m-3.193 16.896l-6.284 1.741l-6.266-1.74L4.363 4h15.274ZM7.82 13h6.895l-.327 3.271l-2.568.917l-3.164-1.13a1 1 0 1 0-.673 1.884l3.5 1.25a1.003 1.003 0 0 0 .673 0l3.5-1.25a1 1 0 0 0 .659-.842l.5-5a1 1 0 0 0-.995-1.1H8.725l-.3-3h7.895a1 1 0 0 0 0-2h-9a1 1 0 0 0-.995 1.1l.5 5a1 1 0 0 0 .995.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HtmlThree(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m3.195 2l1.602 17.994L11.99 22l7.212-2.013L20.805 2Zm14.281 4.123l-.534 5.994l.002.033l-.002.074l-.38 4.19l-.041.373L12 18.037l-.004.004l-4.512-1.258l-.306-3.465H9.39l.157 1.762l2.453.665l2.461-.674l.26-2.869H9.577l-.044-.485l-.101-1.136l-.052-.61h5.538l.202-2.232H6.682l-.044-.485l-.1-1.137l-.053-.61h11.044Zm0 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HtmlThreeAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.952 8h6.986l-.177 2h-4.77a1 1 0 0 0 0 2h4.593l-.264 2.977l-2.329.528l-2.328-.529l-.096-1.064a1 1 0 1 0-1.992.177l.16 1.79a1.001 1.001 0 0 0 .775.887l3.26.74a1.019 1.019 0 0 0 .443 0l3.26-.74a1.001 1.001 0 0 0 .775-.888l.432-4.868l.002-.01l-.001-.004l.346-3.908A1.001 1.001 0 0 0 16.031 6H7.952a1 1 0 0 0 0 2m12.702-5.674A1.002 1.002 0 0 0 19.916 2H4.066a1 1 0 0 0-.996 1.09l1.444 16.194a.999.999 0 0 0 .727.874l6.473 1.805a.99.99 0 0 0 .537 0l6.49-1.812a.999.999 0 0 0 .728-.874l1.443-16.188a1.002 1.002 0 0 0-.258-.763M17.538 18.41l-5.556 1.551l-5.538-1.545L5.16 4h13.664Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hunting(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 9a3 3 0 1 0 3 3a3 3 0 0 0-3-3m0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1m9-2h-1.07A8 8 0 0 0 13 4.07V3a1 1 0 0 0-2 0v1.07A8 8 0 0 0 4.07 11H3a1 1 0 0 0 0 2h1.07A8 8 0 0 0 11 19.93V21a1 1 0 0 0 2 0v-1.07A8 8 0 0 0 19.93 13H21a1 1 0 0 0 0-2m-9 7a6 6 0 1 1 6-6a6 6 0 0 1-6 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Icons(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.6 10.9c.1.1.2.1.4.1h7c.6 0 1-.4 1-1c0-.2 0-.3-.1-.4l-3.5-7c-.3-.5-.9-.7-1.4-.4c-.1.1-.3.2-.4.4l-3.5 7c-.2.4 0 1 .5 1.3m3.9-5.7L19.4 9h-3.8zM6.5 2C4 2 2 4 2 6.5S4 11 6.5 11S11 9 11 6.5S9 2 6.5 2m0 7C5.1 9 4 7.9 4 6.5S5.1 4 6.5 4S9 5.1 9 6.5S7.9 9 6.5 9m4.2 4.3c-.4-.4-1-.4-1.4 0l-2.8 2.8l-2.8-2.8c-.4-.4-1-.4-1.4 0c-.4.4-.4 1 0 1.4l2.8 2.8l-2.8 2.8c-.4.4-.4 1 0 1.4s1 .4 1.4 0l2.8-2.8l2.8 2.8c.4.4 1 .4 1.4 0c.4-.4.4-1 0-1.4l-2.8-2.8l2.8-2.8c.4-.4.4-1 0-1.4M21 13h-7c-.6 0-1 .4-1 1v7c0 .6.4 1 1 1h7c.6 0 1-.4 1-1v-7c0-.6-.4-1-1-1m-1 7h-5v-5h5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Illustration(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.9 9.6c-.2-.5-.8-.7-1.3-.5l-2.9 1.4l-2.1-2.1l-2.1-2.1l1.4-2.9c.2-.5 0-1.1-.5-1.3c-.5-.2-1.1 0-1.3.5l-1.5 3.1l-6.4 1c-.4.1-.7.4-.8.8L2 19.1c-.1.3 0 .7.3.9L4 21.7c.2.2.5.3.7.3h.2l11.6-2.4c.4-.1.7-.4.8-.8l1-6.4l3.1-1.5c.5-.2.7-.8.5-1.3m-6.5 8.2l-9.8 2l3.7-3.7c1.5.7 3.3.1 4-1.4s.1-3.3-1.4-4c-1.1-.5-2.5-.3-3.4.6c-.9.9-1.1 2.3-.6 3.4l-3.7 3.7l2-9.8l5.8-1l2.2 2.2l2.2 2.2zm-5.8-4.4c0-.3.1-.5.3-.7c.4-.4 1-.4 1.4 0c.4.4.4 1 0 1.4c-.4.4-1 .4-1.4 0c-.2-.2-.3-.5-.3-.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Image(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 4H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3M5 18a1 1 0 0 1-1-1v-2.42l3.3-3.29a1 1 0 0 1 1.4 0L15.41 18Zm15-1a1 1 0 0 1-1 1h-.77l-3.81-3.83l.88-.88a1 1 0 0 1 1.4 0l3.3 3.29Zm0-3.24l-1.88-1.87a3.06 3.06 0 0 0-4.24 0l-.88.88l-2.88-2.88a3.06 3.06 0 0 0-4.24 0L4 11.76V7a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageAltSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.71 20.29l-.93-.93l-.09-.1l-.06-.07l-.5-.5l-.13-.07l-5.18-5.2l-.09-.08l-3.2-3.21l-.1-.13l-7.72-7.71a1 1 0 0 0-1.42 1.42l1 1A3 3 0 0 0 3 6v12a3 3 0 0 0 3 3h12a2.9 2.9 0 0 0 1.27-.31s0 0 .05 0l.95 1a1 1 0 0 0 1.42 0a1 1 0 0 0 .02-1.4M5 6.41l3.24 3.24a2.84 2.84 0 0 0-.67.48L5 12.71ZM6 19a1 1 0 0 1-1-1v-2.46l4-4a.81.81 0 0 1 1.1 0L17.59 19ZM9.66 5H18a1 1 0 0 1 1 1v5.94a1 1 0 1 0-1.42 1.42l1.74 1.74a1 1 0 0 0 1.42 0a1 1 0 0 0 .29-.72V6a3 3 0 0 0-3-3H9.66a1 1 0 0 0 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageBlock(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.54 2.46A5 5 0 1 0 22 6a5 5 0 0 0-1.46-3.54M14 6a3 3 0 0 1 3-3a3 3 0 0 1 1.29.3l-4 4A3 3 0 0 1 14 6m5.12 2.12a3.08 3.08 0 0 1-3.4.57l4-4A3 3 0 0 1 20 6a3 3 0 0 1-.88 2.12M19 13a1 1 0 0 0-1 1v.39l-1.48-1.49a2.87 2.87 0 0 0-3.93 0l-.7.71l-2.48-2.49a2.87 2.87 0 0 0-3.93 0L4 12.61V7a1 1 0 0 1 1-1h4a1 1 0 0 0 0-2H5a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 .95-.17h.09A3 3 0 0 0 20 19.44a1.43 1.43 0 0 0 0-.22V14a1 1 0 0 0-1-1M5 20a1 1 0 0 1-1-1v-3.57l2.9-2.89a.79.79 0 0 1 1.09 0l3.19 3.18L15.46 20Zm13-1a1 1 0 0 1-.18.54L13.3 15l.71-.7a.79.79 0 0 1 1.09 0l2.9 2.91Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageBroken(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.71 14.54L19.21 12a1 1 0 0 0-1.42 0L15 14.84L12.21 12a1 1 0 0 0-1.42 0L8.5 14.34L6.21 12a1 1 0 0 0-1.42 0l-2.5 2.5a1 1 0 0 0-.21.33a1 1 0 0 0-.08.38V19a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-3.75a1 1 0 0 0-.08-.38a1 1 0 0 0-.21-.33M20 19a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-3.34l1.5-1.5l2.29 2.3a1 1 0 0 0 1.42 0l2.29-2.3L14.29 17a1 1 0 0 0 1.42 0l2.79-2.8l1.5 1.5ZM19 2H5a3 3 0 0 0-3 3v5.26a1.17 1.17 0 0 0 0 .27v.1a1 1 0 0 0 1.66.31L5.5 9.16l2.29 2.3a1 1 0 0 0 1.42 0l2.29-2.3L14.29 12a1 1 0 0 0 1.42 0l2.79-2.8l1.77 1.78a1 1 0 0 0 1.66-.31a.28.28 0 0 0 0-.09a.88.88 0 0 0 .06-.28V5A3 3 0 0 0 19 2m1 5.84L19.21 7a1 1 0 0 0-1.42 0L15 9.84L12.21 7a1 1 0 0 0-1.42 0L8.5 9.34L6.21 7a1 1 0 0 0-1.42 0L4 7.84V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageCheck(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.71 2.3a1 1 0 0 0-1.41 0l-3.38 3.3l-1.22-1.2a1 1 0 0 0-1.4 1.43l1.92 1.88a1 1 0 0 0 1.4 0l4.08-4a1 1 0 0 0 .01-1.41M19 9a1 1 0 0 0-1 1v4.39l-1.48-1.48a2.77 2.77 0 0 0-3.93 0l-.7.7l-2.48-2.49a2.79 2.79 0 0 0-3.93 0L4 12.6V7a1 1 0 0 1 1-1h6a1 1 0 0 0 0-2H5a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-9a1 1 0 0 0-1-1M5 20a1 1 0 0 1-1-1v-3.57l2.9-2.9a.79.79 0 0 1 1.09 0l3.17 3.17l4.29 4.3Zm13-1a.89.89 0 0 1-.18.53L13.31 15l.7-.7a.78.78 0 0 1 1.1 0L18 17.22Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageDownload(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.71 6.29a1 1 0 0 0-1.42 0L20 7.59V2a1 1 0 0 0-2 0v5.59l-1.29-1.3a1 1 0 0 0-1.42 1.42l3 3a1 1 0 0 0 .33.21a.94.94 0 0 0 .76 0a1 1 0 0 0 .33-.21l3-3a1 1 0 0 0 0-1.42M19 13a1 1 0 0 0-1 1v.38l-1.48-1.48a2.79 2.79 0 0 0-3.93 0l-.7.7l-2.48-2.48a2.85 2.85 0 0 0-3.93 0L4 12.6V7a1 1 0 0 1 1-1h8a1 1 0 0 0 0-2H5a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-5a1 1 0 0 0-1-1M5 20a1 1 0 0 1-1-1v-3.57l2.9-2.9a.79.79 0 0 1 1.09 0l3.17 3.17l4.3 4.3Zm13-1a.89.89 0 0 1-.18.53L13.31 15l.7-.7a.77.77 0 0 1 1.1 0L18 17.21Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageEdit(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.737 3.751l-2.42-2.42a1 1 0 0 0-1.414 0l-4.58 4.58a1 1 0 0 0-.293.707v2.42a1 1 0 0 0 1 1h2.42a1 1 0 0 0 .707-.293l4.58-4.58a1 1 0 0 0 0-1.414m-5.7 4.287H15.03V7.032l3.58-3.58l1.006 1.006ZM19 11a1 1 0 0 0-1 1v2.392l-1.48-1.48a2.78 2.78 0 0 0-3.929 0l-.698.697l-2.486-2.486a2.777 2.777 0 0 0-3.924 0L4 12.606V7a1.001 1.001 0 0 1 1-1h6a1 1 0 0 0 0-2H5a3.003 3.003 0 0 0-3 3v12a3.003 3.003 0 0 0 3 3h12a3.003 3.003 0 0 0 3-3v-7a1 1 0 0 0-1-1M5 20a1.001 1.001 0 0 1-1-1v-3.566l2.897-2.897a.8.8 0 0 1 1.096 0l3.168 3.167c.009.01.012.022.02.03L15.448 20Zm13-1a.971.971 0 0 1-.179.537l-4.514-4.514l.698-.698a.78.78 0 0 1 1.1 0L18 17.22Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageLock(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 5.18V4a3 3 0 0 0-6 0v1.18A3 3 0 0 0 12 8v2a3 3 0 0 0 3 3h4a3 3 0 0 0 3-3V8a3 3 0 0 0-2-2.82M16 4a1 1 0 0 1 2 0v1h-2Zm4 6a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1Zm-1 5a1 1 0 0 0-1 1v3a.89.89 0 0 1-.18.53l-8.41-8.41a2.86 2.86 0 0 0-3.93 0L4 12.6V7a1 1 0 0 1 1-1h5a1 1 0 0 0 0-2H5a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-3a1 1 0 0 0-1-1M5 20a1 1 0 0 1-1-1v-3.57l2.9-2.9a.79.79 0 0 1 1.09 0L15.46 20Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageMinus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 4.008h-4a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2M19 8a1 1 0 0 0-1 1v5.392l-1.48-1.48a2.78 2.78 0 0 0-3.929 0l-.698.697l-2.486-2.486a2.777 2.777 0 0 0-3.924 0L4 12.606V7a1.001 1.001 0 0 1 1-1h8a1 1 0 0 0 0-2H5a3.003 3.003 0 0 0-3 3v12a3.003 3.003 0 0 0 3 3h12a3.003 3.003 0 0 0 3-3V9a1 1 0 0 0-1-1M5 20a1.001 1.001 0 0 1-1-1v-3.566l2.897-2.897a.8.8 0 0 1 1.096 0l3.168 3.167c.009.01.012.022.02.03L15.448 20Zm13-1a.971.971 0 0 1-.179.537l-4.514-4.514l.698-.698a.78.78 0 0 1 1.1 0L18 17.22Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImagePlus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 10a1 1 0 0 0-1 1v3.38l-1.48-1.48a2.79 2.79 0 0 0-3.93 0l-.7.71l-2.48-2.49a2.79 2.79 0 0 0-3.93 0L4 12.61V7a1 1 0 0 1 1-1h8a1 1 0 0 0 0-2H5a3 3 0 0 0-3 3v12.22A2.79 2.79 0 0 0 4.78 22h12.44a2.88 2.88 0 0 0 .8-.12a2.74 2.74 0 0 0 2-2.65V11A1 1 0 0 0 19 10M5 20a1 1 0 0 1-1-1v-3.57l2.89-2.89a.78.78 0 0 1 1.1 0L15.46 20Zm13-1a1 1 0 0 1-.18.54L13.3 15l.71-.7a.77.77 0 0 1 1.1 0L18 17.21Zm3-15h-1V3a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2h1v1a1 1 0 0 0 2 0V6h1a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageQuestion(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 13a1 1 0 0 0-1 1v.39l-1.48-1.48a2.77 2.77 0 0 0-3.93 0l-.7.7l-2.48-2.49a2.86 2.86 0 0 0-3.93 0L4 12.6V7a1 1 0 0 1 1-1h8a1 1 0 0 0 0-2H5a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-5a1 1 0 0 0-1-1M5 20a1 1 0 0 1-1-1v-3.57l2.9-2.9a.79.79 0 0 1 1.09 0l3.17 3.17l4.29 4.3Zm13-1a1 1 0 0 1-.18.53L13.31 15l.7-.7a.78.78 0 0 1 1.1 0L18 17.22Zm1-17a3 3 0 0 0-2.6 1.5a1 1 0 0 0 .37 1.37a1 1 0 0 0 1.36-.37A1 1 0 0 1 20 5a1 1 0 0 1-1 1a1 1 0 0 0 0 2a3 3 0 0 0 0-6m.38 7.08A1 1 0 0 0 18.8 9l-.18.06l-.18.09l-.15.13A1 1 0 0 0 18 10a1 1 0 0 0 .29.71a1 1 0 0 0 1.42 0A1 1 0 0 0 20 10a1 1 0 0 0-.62-.92"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageRedo(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.5 1.5a1 1 0 0 0-1 1a5 5 0 1 0 .3 7.75a1 1 0 0 0-1.32-1.51a3 3 0 1 1 .25-4.25H18.5a1 1 0 0 0 0 2h3a1 1 0 0 0 1-1v-3a1 1 0 0 0-1-.99m-3 12a1 1 0 0 0-1 1v.39L16 13.41a2.77 2.77 0 0 0-3.93 0l-.7.7l-2.46-2.49a2.79 2.79 0 0 0-3.93 0L3.5 13.1V7.5a1 1 0 0 1 1-1h5a1 1 0 0 0 0-2h-5a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-5a1 1 0 0 0-1-1m-14 7a1 1 0 0 1-1-1v-3.57L6.4 13a.79.79 0 0 1 1.09 0l3.17 3.17L15 20.5Zm13-1a1 1 0 0 1-.18.53l-4.51-4.51l.7-.7a.78.78 0 0 1 1.1 0l2.89 2.9Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageResizeLandscape(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 11H2a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-9a1 1 0 0 0-1-1m-5.56 9l1.93-1.93a.3.3 0 0 1 .5 0L11.79 20ZM12 17.38l-.72-.71a2.41 2.41 0 0 0-3.33 0L4.61 20H3v-7h9ZM2 4.11a1 1 0 0 0 .86-.49A1.05 1.05 0 0 0 3.05 3A1 1 0 0 0 2 2a1 1 0 0 0-1 1v.1a1 1 0 0 0 1 1.01M9.91 4h.19a1 1 0 0 0 0-2h-.19a1 1 0 0 0 0 2M2 8.78a1 1 0 0 0 1-1v-.22a1 1 0 1 0-2 0v.22a1 1 0 0 0 1 1M14.09 2h-.19a1 1 0 0 0 0 2h.19a1 1 0 0 0 0-2M5.91 4h.19a1 1 0 0 0 0-2h-.19a1 1 0 0 0 0 2M22 6.4a1 1 0 0 0-1 1v.21a1 1 0 0 0 2 0V7.4a1 1 0 0 0-1-1M17.12 20h-.24a1 1 0 1 0 0 2h.24a1 1 0 0 0 0-2M21.9 2a1 1 0 0 0-.9 1a1 1 0 0 0 .1.42a1 1 0 0 0 1.9-.31V3a1.09 1.09 0 0 0-1.1-1m.1 8.9a1 1 0 0 0-1 1v.22a1 1 0 0 0 2 0v-.22a1 1 0 0 0-1-1M18.09 2h-.19a1 1 0 0 0 0 2h.19a1 1 0 0 0 0-2M22 20a.93.93 0 0 0-.44.11A1 1 0 0 0 21 21a1 1 0 0 0 1 1a1.09 1.09 0 0 0 1-1.1a1 1 0 0 0-1-.9m0-4.56a1 1 0 0 0-1 1v.22a1 1 0 1 0 2 0v-.26a1 1 0 0 0-1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageResizeSquare(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 8.1a1 1 0 0 0 1-1v-.19a1 1 0 0 0-2 0v.19a1 1 0 0 0 1 1m0-4a1 1 0 0 0 .42-.1a1 1 0 0 0-.32-2H3a1.09 1.09 0 0 0-1 1.1a1 1 0 0 0 1 .95Zm17.39-.19A1 1 0 0 0 22 3a1 1 0 0 0-1-1h-.1a1 1 0 0 0-.51 1.86Zm-8.5.09h.22a1 1 0 0 0 0-2h-.22a1 1 0 0 0 0 2m-4.5 0h.21a1 1 0 0 0 0-2h-.21a1 1 0 0 0 0 2M21 20a1 1 0 0 0-.42.1a1 1 0 0 0 .32 1.9h.1a1.09 1.09 0 0 0 1-1.1a1 1 0 0 0-1-.9m-7-9a1 1 0 0 0-1-1H3.27A1.08 1.08 0 0 0 3 10a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h10.1a1 1 0 0 0 .9-1.42Zm-2 9H5.52l3.91-3.9a.33.33 0 0 1 .47 0l2.1 2.09Zm0-4.63l-.68-.69a2.35 2.35 0 0 0-3.31 0l-4 4V12h8Zm9 0a1 1 0 0 0-1 1v.21a1 1 0 0 0 2 0v-.18a1 1 0 0 0-1-1Zm0-9a1 1 0 0 0-1 1v.23a1 1 0 1 0 2 0v-.21a1 1 0 0 0-1-1Zm0 4.5a1 1 0 0 0-1 1v.22a1 1 0 0 0 2 0v-.22a1 1 0 0 0-1-.98ZM17.1 20h-.2a1 1 0 1 0 0 2h.2a1 1 0 0 0 0-2m-.49-16a1 1 0 0 0 0-2h-.21a1 1 0 1 0 0 2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageSearch(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 13a1 1 0 0 0-1 1v.39l-1.48-1.48a2.79 2.79 0 0 0-3.93 0l-.7.7l-2.48-2.49a2.87 2.87 0 0 0-3.93 0L4 12.61V7a1 1 0 0 1 1-1h4a1 1 0 0 0 0-2H5a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-5a1 1 0 0 0-1-1M5 20a1 1 0 0 1-1-1v-3.57l2.9-2.89a.79.79 0 0 1 1.09 0l3.17 3.17L15.45 20Zm13-1a1 1 0 0 1-.18.54L13.31 15l.7-.69a.77.77 0 0 1 1.1 0L18 17.22Zm3.71-8.71L20 8.57a4.31 4.31 0 1 0-6.72.79a4.27 4.27 0 0 0 3 1.26a4.34 4.34 0 0 0 2.29-.62l1.72 1.73a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.44M18 8a2.32 2.32 0 1 1 0-3.27A2.32 2.32 0 0 1 18 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageShare(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 7a2 2 0 0 0-1.18.39l-1.75-.8L19 5.71A2 2 0 0 0 20 6a2 2 0 1 0-2-2l-1.89.87A2 2 0 1 0 15 8.5a1.88 1.88 0 0 0 .92-.24l2.1 1A2 2 0 1 0 20 7m-1 6a1 1 0 0 0-1 1v.39l-1.48-1.48a2.77 2.77 0 0 0-3.93 0l-.7.7l-2.48-2.49a2.79 2.79 0 0 0-3.93 0L4 12.6V7a1 1 0 0 1 1-1h5a1 1 0 0 0 0-2H5a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-5a1 1 0 0 0-1-1M5 20a1 1 0 0 1-1-1v-3.57l2.9-2.9a.79.79 0 0 1 1.09 0l3.17 3.17l4.29 4.3Zm13-1a.89.89 0 0 1-.18.53L13.31 15l.7-.7a.78.78 0 0 1 1.1 0L18 17.22Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageShield(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.13 2.36a1 1 0 0 0-.84-.2a2.7 2.7 0 0 1-2.2-.47a1 1 0 0 0-1.18 0a2.7 2.7 0 0 1-2.2.47a1 1 0 0 0-.84.2a1 1 0 0 0-.37.78v3.31a4.63 4.63 0 0 0 1.84 3.7l1.57 1.16a1 1 0 0 0 1.18 0l1.57-1.16a4.63 4.63 0 0 0 1.84-3.7V3.14a1 1 0 0 0-.37-.78M20.5 6.45a2.62 2.62 0 0 1-1 2.09l-1 .72l-1-.72a2.62 2.62 0 0 1-1-2.09V4.22a4.81 4.81 0 0 0 2-.54a4.81 4.81 0 0 0 2 .54Zm-2 7.05a1 1 0 0 0-1 1v.39L16 13.41a2.77 2.77 0 0 0-3.93 0l-.7.7l-2.46-2.49a2.85 2.85 0 0 0-3.93 0L3.5 13.1V7.5a1 1 0 0 1 1-1h7a1 1 0 0 0 0-2h-7a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-5a1 1 0 0 0-1-1m-14 7a1 1 0 0 1-1-1v-3.57L6.4 13a.79.79 0 0 1 1.09 0l3.17 3.17L15 20.5Zm13-1a1 1 0 0 1-.18.53l-4.51-4.51l.7-.7a.78.78 0 0 1 1.1 0l2.89 2.9Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.5 4H10a1 1 0 0 0 0 2h9.5a1 1 0 0 1 1 1v6.76l-1.88-1.88a3 3 0 0 0-1.14-.71a1 1 0 1 0-.64 1.9a.82.82 0 0 1 .36.23l3.31 3.29a.66.66 0 0 0 0 .15a.83.83 0 0 0 0 .15a1.18 1.18 0 0 0 .13.18a.48.48 0 0 0 .09.11a.9.9 0 0 0 .2.14a.6.6 0 0 0 .11.06a.91.91 0 0 0 .37.08a1 1 0 0 0 1-1V7a3 3 0 0 0-2.91-3M3.21 2.29a1 1 0 0 0-1.42 1.42L3.18 5.1A3 3 0 0 0 2.5 7v10a3 3 0 0 0 3 3h12.59l1.7 1.71a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM4.5 7a1 1 0 0 1 .12-.46l2.72 2.71a3 3 0 0 0-1 .63L4.5 11.76Zm1 11a1 1 0 0 1-1-1v-2.42l3.3-3.29a1 1 0 0 1 1.4 0L15.91 18Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageTimes(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 10a1 1 0 0 0-1 1v3.39l-1.48-1.48a2.77 2.77 0 0 0-3.93 0l-.7.7l-2.48-2.49a2.79 2.79 0 0 0-3.93 0L4 12.6V7a1 1 0 0 1 1-1h8a1 1 0 0 0 0-2H5a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-8a1 1 0 0 0-1-1M5 20a1 1 0 0 1-1-1v-3.57l2.9-2.9a.79.79 0 0 1 1.09 0l3.17 3.17l4.29 4.3Zm13-1a.89.89 0 0 1-.18.53L13.31 15l.7-.7a.78.78 0 0 1 1.1 0L18 17.22Zm2.41-14l1.3-1.29a1 1 0 0 0 0-1.41a1 1 0 0 0-1.42 0L19 3.59L17.71 2.3a1 1 0 0 0-1.42 0a1 1 0 0 0 0 1.41L17.59 5l-1.3 1.3a1 1 0 0 0 0 1.41a1 1 0 0 0 1.42 0L19 6.42l1.29 1.29a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.41Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageUpload(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 13a1 1 0 0 0-1 1v.38l-1.48-1.48a2.79 2.79 0 0 0-3.93 0l-.7.7l-2.48-2.48a2.85 2.85 0 0 0-3.93 0L4 12.6V7a1 1 0 0 1 1-1h7a1 1 0 0 0 0-2H5a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-5a1 1 0 0 0-1-1M5 20a1 1 0 0 1-1-1v-3.57l2.9-2.9a.79.79 0 0 1 1.09 0l3.17 3.17l4.3 4.3Zm13-1a.89.89 0 0 1-.18.53L13.31 15l.7-.7a.77.77 0 0 1 1.1 0L18 17.21Zm4.71-14.71l-3-3a1 1 0 0 0-.33-.21a1 1 0 0 0-.76 0a1 1 0 0 0-.33.21l-3 3a1 1 0 0 0 1.42 1.42L18 4.41V10a1 1 0 0 0 2 0V4.41l1.29 1.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageV(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 2H5a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h14a2.81 2.81 0 0 0 .49-.05l.3-.07h.12l.37-.14l.13-.07c.1-.06.21-.11.31-.18a3.79 3.79 0 0 0 .38-.32l.07-.09a2.69 2.69 0 0 0 .27-.32l.09-.13a2.31 2.31 0 0 0 .18-.35a1 1 0 0 0 .07-.15c.05-.12.08-.25.12-.38v-.15a2.6 2.6 0 0 0 .1-.6V5a3 3 0 0 0-3-3M5 20a1 1 0 0 1-1-1v-4.31l3.29-3.3a1 1 0 0 1 1.42 0l8.6 8.61Zm15-1a1 1 0 0 1-.07.36a1 1 0 0 1-.08.14a.94.94 0 0 1-.09.12l-5.35-5.35l.88-.88a1 1 0 0 1 1.42 0l3.29 3.3Zm0-5.14L18.12 12a3.08 3.08 0 0 0-4.24 0l-.88.88L10.12 10a3.08 3.08 0 0 0-4.24 0L4 11.86V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Images(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 15V5a3 3 0 0 0-3-3H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3M4 5a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v4.36l-1.08-1.09a2.56 2.56 0 0 0-1.81-.75a2.58 2.58 0 0 0-1.81.75l-.91.91l-.81-.81a2.93 2.93 0 0 0-4.11 0L4 9.85Zm.12 10.45A.94.94 0 0 1 4 15v-2.33l2.88-2.88a.91.91 0 0 1 1.29 0l.83.81Zm8.6-5.76a.52.52 0 0 1 .39-.17a.52.52 0 0 1 .39.17l2.5 2.49V15a1 1 0 0 1-1 1H6.4ZM21 6a1 1 0 0 0-1 1v10a3 3 0 0 1-3 3H7a1 1 0 0 0 0 2h10a5 5 0 0 0 5-5V7a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Import(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 14a1 1 0 0 0-1 1v4a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-4a1 1 0 0 0-2 0v4a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-4a1 1 0 0 0-1-1m-9.71 1.71a1 1 0 0 0 .33.21a.94.94 0 0 0 .76 0a1 1 0 0 0 .33-.21l4-4a1 1 0 0 0-1.42-1.42L13 12.59V3a1 1 0 0 0-2 0v9.59l-2.29-2.3a1 1 0 1 0-1.42 1.42Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Inbox(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.056 2h-14a3.003 3.003 0 0 0-3 3v14a3.003 3.003 0 0 0 3 3h14a3.003 3.003 0 0 0 3-3V5a3.003 3.003 0 0 0-3-3m-14 2h14a1.001 1.001 0 0 1 1 1v8H17.59a1.997 1.997 0 0 0-1.664.89L14.52 16H9.59l-1.406-2.11A1.997 1.997 0 0 0 6.52 13H4.056V5a1.001 1.001 0 0 1 1-1m14 16h-14a1.001 1.001 0 0 1-1-1v-4H6.52l1.406 2.11A1.997 1.997 0 0 0 9.59 18h4.93a1.997 1.997 0 0 0 1.664-.89L17.59 15h2.465v4a1.001 1.001 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IncomingCall(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.55 9a1.07 1.07 0 0 0 .39.07h4a1 1 0 0 0 0-2h-1.59l3.29-3.29a1 1 0 1 0-1.41-1.41l-3.29 3.28V4.06a1 1 0 1 0-2 0v4a1.07 1.07 0 0 0 .07.39a1 1 0 0 0 .54.55m3.89 4c-.22 0-.45-.07-.67-.12a9.44 9.44 0 0 1-1.31-.39a2 2 0 0 0-2.48 1l-.22.45a12.18 12.18 0 0 1-2.66-2a12.18 12.18 0 0 1-2-2.66l.42-.28a2 2 0 0 0 1-2.48a10.33 10.33 0 0 1-.39-1.31c-.05-.22-.09-.45-.12-.68a3 3 0 0 0-3-2.49h-3a3 3 0 0 0-3 3.41a19 19 0 0 0 16.52 16.46h.38a3 3 0 0 0 2-.76a3 3 0 0 0 1-2.25v-3a3 3 0 0 0-2.47-2.9m.5 6a1 1 0 0 1-.34.75a1.06 1.06 0 0 1-.82.25A17 17 0 0 1 4.07 5.22a1.09 1.09 0 0 1 .25-.82a1 1 0 0 1 .75-.34h3a1 1 0 0 1 1 .79q.06.41.15.81a11.12 11.12 0 0 0 .46 1.55l-1.4.65a1 1 0 0 0-.49 1.33a14.49 14.49 0 0 0 7 7a1 1 0 0 0 .76 0a1 1 0 0 0 .57-.52l.62-1.4a13.69 13.69 0 0 0 1.58.46q.4.09.81.15a1 1 0 0 1 .79 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Info(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 10a1 1 0 0 0-1 1v6a1 1 0 0 0 2 0v-6a1 1 0 0 0-1-1m0-4a1.25 1.25 0 1 0 1.25 1.25A1.25 1.25 0 0 0 12 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoCircle(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10.011 10.011 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8.01 8.01 0 0 1-8 8m0-8.5a1 1 0 0 0-1 1v3a1 1 0 0 0 2 0v-3a1 1 0 0 0-1-1m0-4a1.25 1.25 0 1 0 1.25 1.25A1.25 1.25 0 0 0 12 7.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Instagram(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.34 5.46a1.2 1.2 0 1 0 1.2 1.2a1.2 1.2 0 0 0-1.2-1.2m4.6 2.42a7.59 7.59 0 0 0-.46-2.43a4.94 4.94 0 0 0-1.16-1.77a4.7 4.7 0 0 0-1.77-1.15a7.3 7.3 0 0 0-2.43-.47C15.06 2 14.72 2 12 2s-3.06 0-4.12.06a7.3 7.3 0 0 0-2.43.47a4.78 4.78 0 0 0-1.77 1.15a4.7 4.7 0 0 0-1.15 1.77a7.3 7.3 0 0 0-.47 2.43C2 8.94 2 9.28 2 12s0 3.06.06 4.12a7.3 7.3 0 0 0 .47 2.43a4.7 4.7 0 0 0 1.15 1.77a4.78 4.78 0 0 0 1.77 1.15a7.3 7.3 0 0 0 2.43.47C8.94 22 9.28 22 12 22s3.06 0 4.12-.06a7.3 7.3 0 0 0 2.43-.47a4.7 4.7 0 0 0 1.77-1.15a4.85 4.85 0 0 0 1.16-1.77a7.59 7.59 0 0 0 .46-2.43c0-1.06.06-1.4.06-4.12s0-3.06-.06-4.12M20.14 16a5.61 5.61 0 0 1-.34 1.86a3.06 3.06 0 0 1-.75 1.15a3.19 3.19 0 0 1-1.15.75a5.61 5.61 0 0 1-1.86.34c-1 .05-1.37.06-4 .06s-3 0-4-.06a5.73 5.73 0 0 1-1.94-.3a3.27 3.27 0 0 1-1.1-.75a3 3 0 0 1-.74-1.15a5.54 5.54 0 0 1-.4-1.9c0-1-.06-1.37-.06-4s0-3 .06-4a5.54 5.54 0 0 1 .35-1.9A3 3 0 0 1 5 5a3.14 3.14 0 0 1 1.1-.8A5.73 5.73 0 0 1 8 3.86c1 0 1.37-.06 4-.06s3 0 4 .06a5.61 5.61 0 0 1 1.86.34a3.06 3.06 0 0 1 1.19.8a3.06 3.06 0 0 1 .75 1.1a5.61 5.61 0 0 1 .34 1.9c.05 1 .06 1.37.06 4s-.01 3-.06 4M12 6.87A5.13 5.13 0 1 0 17.14 12A5.12 5.12 0 0 0 12 6.87m0 8.46A3.33 3.33 0 1 1 15.33 12A3.33 3.33 0 0 1 12 15.33"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InstagramAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 9.52A2.48 2.48 0 1 0 14.48 12A2.48 2.48 0 0 0 12 9.52m9.93-2.45a6.53 6.53 0 0 0-.42-2.26a4 4 0 0 0-2.32-2.32a6.53 6.53 0 0 0-2.26-.42C15.64 2 15.26 2 12 2s-3.64 0-4.93.07a6.53 6.53 0 0 0-2.26.42a4 4 0 0 0-2.32 2.32a6.53 6.53 0 0 0-.42 2.26C2 8.36 2 8.74 2 12s0 3.64.07 4.93a6.86 6.86 0 0 0 .42 2.27a3.94 3.94 0 0 0 .91 1.4a3.89 3.89 0 0 0 1.41.91a6.53 6.53 0 0 0 2.26.42C8.36 22 8.74 22 12 22s3.64 0 4.93-.07a6.53 6.53 0 0 0 2.26-.42a3.89 3.89 0 0 0 1.41-.91a3.94 3.94 0 0 0 .91-1.4a6.6 6.6 0 0 0 .42-2.27C22 15.64 22 15.26 22 12s0-3.64-.07-4.93m-2.54 8a5.73 5.73 0 0 1-.39 1.8A3.86 3.86 0 0 1 16.87 19a5.73 5.73 0 0 1-1.81.35H8.94A5.73 5.73 0 0 1 7.13 19a3.51 3.51 0 0 1-1.31-.86A3.51 3.51 0 0 1 5 16.87a5.49 5.49 0 0 1-.34-1.81V8.94A5.49 5.49 0 0 1 5 7.13a3.51 3.51 0 0 1 .86-1.31A3.59 3.59 0 0 1 7.13 5a5.73 5.73 0 0 1 1.81-.35h6.12a5.73 5.73 0 0 1 1.81.35a3.51 3.51 0 0 1 1.31.86A3.51 3.51 0 0 1 19 7.13a5.73 5.73 0 0 1 .35 1.81V12c0 2.06.07 2.27.04 3.06Zm-1.6-7.44a2.38 2.38 0 0 0-1.41-1.41A4 4 0 0 0 15 6H9a4 4 0 0 0-1.38.26a2.38 2.38 0 0 0-1.41 1.36A4.27 4.27 0 0 0 6 9v6a4.27 4.27 0 0 0 .26 1.38a2.38 2.38 0 0 0 1.41 1.41a4.27 4.27 0 0 0 1.33.26h6a4 4 0 0 0 1.38-.26a2.38 2.38 0 0 0 1.41-1.41a4 4 0 0 0 .26-1.38V9a3.78 3.78 0 0 0-.26-1.38ZM12 15.82A3.81 3.81 0 0 1 8.19 12A3.82 3.82 0 1 1 12 15.82m4-6.89a.9.9 0 0 1 0-1.79a.9.9 0 0 1 0 1.79"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Intercom(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.5 2h-15A2.5 2.5 0 0 0 2 4.5v15A2.5 2.5 0 0 0 4.5 22h15a2.5 2.5 0 0 0 2.5-2.5v-15A2.5 2.5 0 0 0 19.5 2m-4.83 3.67a.66.66 0 0 1 .67-.67a.68.68 0 0 1 .66.66v8.9a.67.67 0 0 1-1.33 0Zm-3.34-.34a.67.67 0 0 1 .67-.67a.67.67 0 0 1 .67.67V15a.67.67 0 0 1-1.34 0ZM8 5.67a.67.67 0 0 1 1.33 0v8.9a.66.66 0 0 1-.67.66a.68.68 0 0 1-.66-.66ZM4.67 7A.67.67 0 0 1 6 7v6a.67.67 0 0 1-.67.66a.67.67 0 0 1-.66-.66ZM19.1 17.17a11.3 11.3 0 0 1-7.1 2.16a11.3 11.3 0 0 1-7.1-2.16a.67.67 0 0 1 .87-1A10.2 10.2 0 0 0 12 18a10.15 10.15 0 0 0 6.23-1.84a.67.67 0 0 1 .87 1Zm.23-4.17A.67.67 0 0 1 18 13V7a.67.67 0 0 1 .67-.66a.66.66 0 0 1 .66.66Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IntercomAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.991 15a1 1 0 0 0 1-1V6a1 1 0 1 0-2 0v8a1 1 0 0 0 1 1m4 0a1 1 0 0 0 1-1V6a1 1 0 1 0-2 0v8a1 1 0 0 0 1 1m-8-2a1 1 0 0 0 1-1V8a1 1 0 0 0-2 0v4a1 1 0 0 0 1 1m14-12h-16a3.003 3.003 0 0 0-3 3v16a3.003 3.003 0 0 0 3 3h16a3.003 3.003 0 0 0 3-3V4a3.003 3.003 0 0 0-3-3m1 19a1 1 0 0 1-1 1h-16a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1Zm-3.643-4.591A8.82 8.82 0 0 1 11.99 17a8.987 8.987 0 0 1-5.356-1.591a1 1 0 1 0-1.287 1.53A10.8 10.8 0 0 0 11.99 19a10.8 10.8 0 0 0 6.644-2.06a1 1 0 0 0-1.287-1.531M17.99 7a1 1 0 0 0-1 1v4a1 1 0 1 0 2 0V8a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Invoice(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 16H7a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2m-4-6h2a1 1 0 0 0 0-2H9a1 1 0 0 0 0 2m12 2h-3V3a1 1 0 0 0-.5-.87a1 1 0 0 0-1 0l-3 1.72l-3-1.72a1 1 0 0 0-1 0l-3 1.72l-3-1.72a1 1 0 0 0-1 0A1 1 0 0 0 2 3v16a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-6a1 1 0 0 0-1-1M5 20a1 1 0 0 1-1-1V4.73l2 1.14a1.08 1.08 0 0 0 1 0l3-1.72l3 1.72a1.08 1.08 0 0 0 1 0l2-1.14V19a3 3 0 0 0 .18 1Zm15-1a1 1 0 0 1-2 0v-5h2Zm-7-7H7a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Italic(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 6h-6a1 1 0 0 0 0 2h1.52l-3.2 8H7a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2h-1.52l3.2-8H17a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Jackhammer(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.85 15.69a1 1 0 0 0-1.41 0l-2.06 2.06a1 1 0 0 0 .45 1.67l.26.07l-.8.8a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l2-2a1 1 0 0 0 .26-1a1 1 0 0 0-.71-.71l-.26-.06l.83-.84a1 1 0 0 0 .02-1.41m-10.94 3.8l.26-.07a1 1 0 0 0 .45-1.67l-2.06-2.06a1 1 0 0 0-1.41 1.41l.83.84l-.26.06a1 1 0 0 0-.71.71a1 1 0 0 0 .26 1l2 2a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM21 4a1 1 0 0 0-1 1h-3a3 3 0 0 0-3-3h-4a3 3 0 0 0-3 3H4a1 1 0 0 0-2 0v2a1 1 0 0 0 2 0h3v2a3 3 0 0 0 2 2.83V13a2 2 0 0 0 2 2v6a1 1 0 0 0 2 0v-6a2 2 0 0 0 2-2v-1.17A3 3 0 0 0 17 9V7h3a1 1 0 0 0 2 0V5a1 1 0 0 0-1-1m-6 5a1 1 0 0 1-1 1a1 1 0 0 0-1 1v2h-2v-2a1 1 0 0 0-1-1a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1Zm-3-3a1 1 0 0 0-1 1v1a1 1 0 0 0 2 0V7a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func JavaScript(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.478 14.883a4.061 4.061 0 0 1-2.187-.398a1.439 1.439 0 0 1-.536-1.01a.222.222 0 0 0-.226-.22a46.826 46.826 0 0 0-.95 0a.211.211 0 0 0-.231.186a2.339 2.339 0 0 0 .753 1.844a3.991 3.991 0 0 0 2.228.839a8.062 8.062 0 0 0 2.533-.108a3.126 3.126 0 0 0 1.678-.904a2.338 2.338 0 0 0 .396-2.231a1.869 1.869 0 0 0-1.23-1.095c-1.28-.45-2.664-.415-3.97-.757c-.227-.07-.504-.148-.605-.388a.855.855 0 0 1 .284-.955a2.558 2.558 0 0 1 1.35-.336a4.07 4.07 0 0 1 1.883.27a1.436 1.436 0 0 1 .687.992a.243.243 0 0 0 .228.236c.314.006.628.001.943.002a.228.228 0 0 0 .247-.168a2.434 2.434 0 0 0-1.187-2.106a5.88 5.88 0 0 0-3.218-.493a3.505 3.505 0 0 0-2.176.875a2.175 2.175 0 0 0-.434 2.262a1.93 1.93 0 0 0 1.218 1.062c1.277.461 2.676.313 3.964.721c.252.085.544.216.621.495a.99.99 0 0 1-.27.946a2.97 2.97 0 0 1-1.793.439m5.819-8.445q-3.738-2.114-7.479-4.225a1.677 1.677 0 0 0-1.637 0L3.73 6.421a1.542 1.542 0 0 0-.805 1.342v8.475a1.553 1.553 0 0 0 .836 1.355c.713.388 1.406.816 2.133 1.179a3.064 3.064 0 0 0 2.738.075a2.128 2.128 0 0 0 .995-1.921c.005-2.797 0-5.594.002-8.39a.22.22 0 0 0-.207-.255a41.555 41.555 0 0 0-.953 0a.21.21 0 0 0-.228.213c-.004 2.779.001 5.558-.002 8.338a.94.94 0 0 1-.61.883a1.532 1.532 0 0 1-1.24-.166l-1.982-1.12a.237.237 0 0 1-.135-.235V7.807a.259.259 0 0 1 .157-.26q3.713-2.092 7.425-4.187a.258.258 0 0 1 .292 0l7.426 4.186a.262.262 0 0 1 .156.26v8.388a.242.242 0 0 1-.134.238q-3.656 2.068-7.317 4.13c-.116.064-.254.169-.39.09c-.64-.362-1.27-.738-1.908-1.103a.206.206 0 0 0-.23-.014a5.218 5.218 0 0 1-.882.412c-.138.056-.308.072-.403.2a1.316 1.316 0 0 0 .432.31l2.236 1.293a1.63 1.63 0 0 0 1.655.046q3.726-2.1 7.452-4.204a1.556 1.556 0 0 0 .836-1.354V7.763a1.54 1.54 0 0 0-.778-1.325"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Kayak(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.12 16.88a3 3 0 0 0-3.4-.58l-1.15-1.14a24 24 0 0 0 5.21-8.86a3.19 3.19 0 0 0-4.08-4.08a24 24 0 0 0-8.86 5.21L7.7 6.28a3 3 0 1 0-4.82.84A3 3 0 0 0 5 8a3 3 0 0 0 1.28-.3l1.15 1.14a24 24 0 0 0-5.21 8.86A3.24 3.24 0 0 0 3 21a3.17 3.17 0 0 0 2.22 1a3.74 3.74 0 0 0 1.08-.17a24 24 0 0 0 8.86-5.21l1.14 1.15a3 3 0 1 0 4.82-.84ZM5.71 5.7a1 1 0 0 1-1.42-1.41A1 1 0 0 1 5.71 5.7m12.6-1.57a1.6 1.6 0 0 1 .47-.08a1.16 1.16 0 0 1 .83.34a1.23 1.23 0 0 1 .26 1.3a22.09 22.09 0 0 1-2.13 4.64l-4.07-4.07a22.09 22.09 0 0 1 4.64-2.13M5.69 19.87a1.2 1.2 0 0 1-1.56-1.56a22.09 22.09 0 0 1 2.13-4.64l4.07 4.07a22.09 22.09 0 0 1-4.64 2.13M12 16.59L7.41 12a21.29 21.29 0 0 1 1.43-1.74l4.91 4.91A21.29 21.29 0 0 1 12 16.59m3.15-2.84l-4.9-4.91A21.29 21.29 0 0 1 12 7.41l4.6 4.6a21.29 21.29 0 0 1-1.44 1.74Zm4.55 6a1 1 0 1 1-1.42-1.41a1 1 0 0 1 1.41 0a1 1 0 0 1 .02 1.37Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeySkeleton(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21 4.41l.71-.7a1 1 0 1 0-1.42-1.42l-1.4 1.41l-2.83 2.83l-6.31 6.3a5 5 0 1 0 1.42 1.42l5.59-5.6l2.12 2.13a1 1 0 1 0 1.41-1.42l-2.12-2.12l1.42-1.41l.7.7a1 1 0 1 0 1.42-1.41ZM7 20a3 3 0 1 1 3-3a3 3 0 0 1-3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeySkeletonAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.71 6.53l-1.42-1.41l1.42-1.41a1 1 0 1 0-1.42-1.42L9.75 12.83a5 5 0 1 0 1.42 1.42l4.88-4.89l1.41 1.42a1 1 0 0 0 .71.29a1 1 0 0 0 .71-.29a1 1 0 0 0 0-1.42L17.46 8l1.42-1.42L20.29 8a1 1 0 0 0 .71.29a1 1 0 0 0 .71-.29a1 1 0 0 0 0-1.47M7 20a3 3 0 1 1 3-3a3 3 0 0 1-3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Keyboard(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.21 13.29a.93.93 0 0 0-.33-.21a1 1 0 0 0-.76 0a.9.9 0 0 0-.54.54a1 1 0 1 0 1.84 0a1 1 0 0 0-.21-.33M13.5 11h1a1 1 0 0 0 0-2h-1a1 1 0 0 0 0 2m-4 0h1a1 1 0 0 0 0-2h-1a1 1 0 0 0 0 2m-3-2h-1a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2M20 5H4a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h16a3 3 0 0 0 3-3V8a3 3 0 0 0-3-3m1 11a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1Zm-6-3H9a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2m3.5-4h-1a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2m.71 4.29a1 1 0 0 0-.33-.21a1 1 0 0 0-.76 0a.93.93 0 0 0-.33.21a1 1 0 0 0-.21.33a1 1 0 1 0 1.92.38a.84.84 0 0 0-.08-.38a1 1 0 0 0-.21-.33"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyboardAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.71 9.29a1 1 0 0 0-1.42 0a1 1 0 0 0-.21.33a1 1 0 0 0 .21 1.09a1.15 1.15 0 0 0 .33.21A.84.84 0 0 0 6 11a1 1 0 0 0 .92-1.38a1 1 0 0 0-.21-.33M10 11a1 1 0 0 0 .92-1.38a1 1 0 0 0-.21-.33a1 1 0 0 0-.9-.29a.6.6 0 0 0-.19.06l-.18.09l-.15.12A1.05 1.05 0 0 0 9 10a1 1 0 0 0 1 1m-3.62 2.08a1 1 0 0 0-.76 0A1 1 0 0 0 5 14a1 1 0 0 0 1.38.92a1.15 1.15 0 0 0 .33-.21A1 1 0 0 0 7 14a1 1 0 0 0-.29-.71a.93.93 0 0 0-.33-.21M14 13h-4a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2m.71-3.71a1 1 0 0 0-1.42 0a1 1 0 0 0-.21.33A1 1 0 1 0 15 10a.84.84 0 0 0-.08-.38a1 1 0 0 0-.21-.33m3.85 3.88a.76.76 0 0 0-.18-.09a1 1 0 0 0-.76 0a1.15 1.15 0 0 0-.33.21A1.05 1.05 0 0 0 17 14a1 1 0 1 0 2 0a1.05 1.05 0 0 0-.29-.71ZM20 5H4a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h16a3 3 0 0 0 3-3V8a3 3 0 0 0-3-3m1 11a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1Zm-2.29-6.71A1 1 0 0 0 17 10a1 1 0 1 0 1.92-.38a1 1 0 0 0-.21-.33"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyboardHide(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.71 10.29a1.15 1.15 0 0 0-.33-.21a1 1 0 0 0-1.09.21a1 1 0 0 0-.21.33a1 1 0 0 0 .21 1.09a1.15 1.15 0 0 0 .33.21a1 1 0 0 0 1.3-1.3a1 1 0 0 0-.21-.33m2.58-2.58A1 1 0 0 0 10 8a.84.84 0 0 0 .38-.08a1.15 1.15 0 0 0 .33-.21a1.15 1.15 0 0 0 .21-.33a.94.94 0 0 0 0-.76a1.15 1.15 0 0 0-.21-.33a1 1 0 0 0-1.42 0a1.15 1.15 0 0 0-.21.33a.94.94 0 0 0 0 .76a1.15 1.15 0 0 0 .21.33M6.71 6.29A1 1 0 0 0 5 7a1 1 0 0 0 .08.38a1.15 1.15 0 0 0 .21.33A1 1 0 0 0 6 8a.84.84 0 0 0 .38-.08a1.15 1.15 0 0 0 .33-.21a1.15 1.15 0 0 0 .21-.33a.94.94 0 0 0 0-.76a1.15 1.15 0 0 0-.21-.33m6.58 12L12 19.59l-1.29-1.3a1 1 0 0 0-1.42 1.42l2 2a1 1 0 0 0 1.42 0l2-2a1 1 0 0 0-1.42-1.42m5.42-12A1 1 0 0 0 17 7a.84.84 0 0 0 .08.38a1.15 1.15 0 0 0 .21.33a1 1 0 0 0 1.42 0a1.15 1.15 0 0 0 .21-.33A.84.84 0 0 0 19 7a1 1 0 0 0-.08-.38a1.15 1.15 0 0 0-.21-.33M14 10h-4a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2m6-8H4a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h16a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3m1 11a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1Zm-2.29-2.71a1 1 0 0 0-.33-.21a.92.92 0 0 0-.76 0a1.15 1.15 0 0 0-.33.21A1.05 1.05 0 0 0 17 11a1 1 0 1 0 1.92-.38a1 1 0 0 0-.21-.33m-5.09-4.21a1.15 1.15 0 0 0-.33.21A1.05 1.05 0 0 0 13 7a1 1 0 0 0 .08.38a1.15 1.15 0 0 0 .21.33A1 1 0 0 0 14 8a.84.84 0 0 0 .38-.08a1.15 1.15 0 0 0 .33-.21a1.15 1.15 0 0 0 .21-.33A.84.84 0 0 0 15 7a1.05 1.05 0 0 0-.29-.71a1 1 0 0 0-1.09-.21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyboardShow(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.71 10.29A1 1 0 0 0 5 11a1 1 0 1 0 1.92-.38a1 1 0 0 0-.21-.33m2.58-2.58A1 1 0 0 0 10 8a1 1 0 0 0 .71-.29a1.15 1.15 0 0 0 .21-.33A1 1 0 0 0 11 7a1.05 1.05 0 0 0-.29-.71l-.15-.12l-.18-.09a.6.6 0 0 0-.19-.08a1 1 0 0 0-.9.27a1 1 0 0 0-.21.33a.94.94 0 0 0 0 .76a1.15 1.15 0 0 0 .21.35M6.56 6.17a.76.76 0 0 0-.18-.09L6.2 6a1 1 0 0 0-.91.27a1 1 0 0 0-.21.33a.94.94 0 0 0 0 .76a1.15 1.15 0 0 0 .21.33a1.15 1.15 0 0 0 .33.21A.84.84 0 0 0 6 8a1 1 0 0 0 .71-.29a1.15 1.15 0 0 0 .21-.33A1 1 0 0 0 7 7a1.05 1.05 0 0 0-.29-.71Zm6.15 12.12a1 1 0 0 0-1.42 0l-2 2a1 1 0 0 0 1.42 1.42l1.29-1.3l1.29 1.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Zm6-8a1 1 0 0 0-1.42 0a1 1 0 0 0-.21.33a1 1 0 0 0 1.3 1.3a1.15 1.15 0 0 0 .33-.21A1 1 0 0 0 19 11a.84.84 0 0 0-.08-.38a1 1 0 0 0-.21-.33M14 10h-4a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2m6-8H4a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h16a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3m1 11a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1Zm-3.38-6.92a.93.93 0 0 0-.33.21A1.05 1.05 0 0 0 17 7a1 1 0 0 0 .08.38a1.15 1.15 0 0 0 .21.33A1 1 0 0 0 18 8a1 1 0 0 0 .71-.29a1.15 1.15 0 0 0 .21-.33A1 1 0 0 0 19 7a1.05 1.05 0 0 0-.29-.71a1 1 0 0 0-1.09-.21m-3.06.09l-.18-.09L14.2 6a1 1 0 0 0-.58.06a.93.93 0 0 0-.33.21a1 1 0 0 0-.21.33a.94.94 0 0 0 0 .76a1.15 1.15 0 0 0 .21.33A1 1 0 0 0 14 8a1 1 0 0 0 .71-.29a1.15 1.15 0 0 0 .21-.33A1 1 0 0 0 15 7a1.05 1.05 0 0 0-.29-.71Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyholeCircle(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 8a2 2 0 0 0-2 2a2 2 0 0 0 1 1.72V15a1 1 0 0 0 2 0v-3.28A2 2 0 0 0 14 10a2 2 0 0 0-2-2m0-6a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyholeSquare(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 2H5a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3m1 17a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1ZM12 8a2 2 0 0 0-2 2a2 2 0 0 0 1 1.72V15a1 1 0 0 0 2 0v-3.28A2 2 0 0 0 14 10a2 2 0 0 0-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyholeSquareFull(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 11.72V15a1 1 0 0 0 2 0v-3.28A2 2 0 0 0 14 10a2 2 0 0 0-4 0a2 2 0 0 0 1 1.72M21 2H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1m-1 18H4V4h16Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Kid(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 10a1 1 0 1 0-1 1a1 1 0 0 0 1-1m4.5 4.06a1 1 0 0 0-1.37.36a1.3 1.3 0 0 1-2.26 0a1 1 0 0 0-1.37-.36a1 1 0 0 0-.37 1.36a3.31 3.31 0 0 0 5.74 0a1 1 0 0 0-.37-1.36M15 9a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-3-7a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18A8 8 0 0 1 9 4.57A3 3 0 0 0 9 5a3 3 0 0 0 3 3a1 1 0 0 0 0-2a1 1 0 0 1 0-2a8 8 0 0 1 0 16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Label(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.71 11.29l-5-5A1 1 0 0 0 16 6H5a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3h11a1 1 0 0 0 .71-.29l5-5a1 1 0 0 0 0-1.42M15.59 16H5a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1h10.59l4 4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LabelAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 12a1 1 0 1 0 1-1a1 1 0 0 0-1 1m6.71-.71l-5-5A1 1 0 0 0 16 6H5a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3h11a1 1 0 0 0 .71-.29l5-5a1 1 0 0 0 0-1.42M15.59 16H5a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1h10.59l4 4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lamp(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 2.78A1 1 0 0 0 17 2H7a1 1 0 0 0-1 .78l-2 9a1 1 0 0 0 .2.85A1 1 0 0 0 5 13h3.94A8.26 8.26 0 0 1 9 14a8.92 8.92 0 0 1-2.57 6.3a1 1 0 0 0 .71 1.7h9.72a1 1 0 0 0 .71-1.7A8.92 8.92 0 0 1 15 14a8.26 8.26 0 0 1 .06-1H16v2a1 1 0 0 0 2 0v-2h1a1 1 0 0 0 .78-.37a1 1 0 0 0 .2-.85ZM9.22 20A10.9 10.9 0 0 0 11 14c0-.33 0-.67-.05-1h2.1c0 .33-.05.67-.05 1a10.9 10.9 0 0 0 1.78 6Zm-3-9L7.8 4h8.4l1.55 7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Language(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.056 12h-2a1 1 0 0 0 0 2v2H17.87a2.965 2.965 0 0 0 .185-1a3 3 0 0 0-5.598-1.5a1 1 0 1 0 1.732 1a1 1 0 0 1 .866-.5a1 1 0 0 1 0 2a1 1 0 0 0 0 2a1 1 0 1 1 0 2a1 1 0 0 1-.866-.5a1 1 0 1 0-1.732 1a3 3 0 0 0 5.598-1.5a2.965 2.965 0 0 0-.185-1h1.185v3a1 1 0 0 0 2 0v-7a1 1 0 1 0 0-2m-11.97-.757a1 1 0 1 0 1.94-.486l-1.757-7.03a2.28 2.28 0 0 0-4.425 0l-1.758 7.03a1 1 0 1 0 1.94.486L5.585 9h2.94ZM6.086 7l.697-2.787a.292.292 0 0 1 .546 0L8.026 7Zm7.97 0h1a1.001 1.001 0 0 1 1 1v1a1 1 0 0 0 2 0V8a3.003 3.003 0 0 0-3-3h-1a1 1 0 0 0 0 2m-4 9h-1a1.001 1.001 0 0 1-1-1v-1a1 1 0 0 0-2 0v1a3.003 3.003 0 0 0 3 3h1a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Laptop(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 14h-1V7a3 3 0 0 0-3-3H7a3 3 0 0 0-3 3v7H3a1 1 0 0 0-1 1v2a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-2a1 1 0 0 0-1-1M6 7a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v7H6Zm14 10a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-1h16Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LaptopCloud(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 10H10a3 3 0 0 0 1.07-5.8a4 4 0 0 0-7.48 1A2.5 2.5 0 0 0 4.5 10m0-3a1 1 0 0 0 1-1a2 2 0 0 1 3.89-.64a1 1 0 0 0 .78.66A1 1 0 0 1 11 7a1 1 0 0 1-1 1H4.5a.5.5 0 0 1 0-1M21 16h-1V9a3 3 0 0 0-3-3h-1a1 1 0 0 0 0 2h1a1 1 0 0 1 1 1v7H6v-3a1 1 0 0 0-2 0v3H3a1 1 0 0 0-1 1v2a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-2a1 1 0 0 0-1-1m-1 3a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-1h16Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LaptopConnection(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 18h-6.18A3 3 0 0 0 13 16.18V13h7a1 1 0 0 0 0-2h-1V5a3 3 0 0 0-3-3H8a3 3 0 0 0-3 3v6H4a1 1 0 0 0 0 2h7v3.18A3 3 0 0 0 9.18 18H3a1 1 0 0 0 0 2h6.18a3 3 0 0 0 5.64 0H21a1 1 0 0 0 0-2M7 11V5a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1v6Zm5 9a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Laughing(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.16 12.21a1 1 0 0 0 .71.29a1 1 0 0 0 .71-.29a1 1 0 0 0 0-1.42l-.79-.79l.79-.79a1 1 0 1 0-1.42-1.42l-1.5 1.5a1 1 0 0 0 0 1.42Zm-5.08 0l1.5-1.5a1 1 0 0 0 0-1.42l-1.5-1.5a1 1 0 0 0-1.42 1.42l.8.79l-.8.79a1 1 0 0 0 0 1.42a1 1 0 0 0 .71.29a1 1 0 0 0 .71-.29m5.28 2a3.76 3.76 0 0 1-4.72 0a1 1 0 0 0-1.28 1.54a5.68 5.68 0 0 0 7.28 0a1 1 0 1 0-1.28-1.54M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayerGroup(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m2.5 8.86l9 5.2a1 1 0 0 0 1 0l9-5.2A1 1 0 0 0 22 8a1 1 0 0 0-.5-.87l-9-5.19a1 1 0 0 0-1 0l-9 5.19A1 1 0 0 0 2 8a1 1 0 0 0 .5.86M12 4l7 4l-7 4l-7-4Zm8.5 7.17L12 16l-8.5-4.87a1 1 0 0 0-1.37.37a1 1 0 0 0 .37 1.36l9 5.2a1 1 0 0 0 1 0l9-5.2a1 1 0 0 0 .37-1.36a1 1 0 0 0-1.37-.37Zm0 4L12 20l-8.5-4.87a1 1 0 0 0-1.37.37a1 1 0 0 0 .37 1.36l9 5.2a1 1 0 0 0 1 0l9-5.2a1 1 0 0 0 .37-1.36a1 1 0 0 0-1.37-.37Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayerGroupSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.26 5L12 4l7 4l-3.15 1.83a1 1 0 0 0-.37 1.36a1 1 0 0 0 1.37.37l4.65-2.69a1 1 0 0 0 0-1.74l-9-5.2a1 1 0 0 0-1 0l-2.24 1.3a1 1 0 0 0-.37 1.37a1 1 0 0 0 1.37.4M3.71 2.29a1 1 0 0 0-1.42 1.42L4.54 6l-2 1.17a1 1 0 0 0 0 1.74l9 5.2a1 1 0 0 0 1 0l.1-.06l1.07 1.07l-1.67 1l-8.54-4.99a1 1 0 1 0-1 1.74l9 5.2a1 1 0 0 0 .5.13a1 1 0 0 0 .5-.13l2.63-1.52l1.07 1.07L12 20l-8.5-4.87a1 1 0 0 0-1 1.74l9 5.2a1 1 0 0 0 1 0l5.17-3l2.62 2.63a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM5 8l1-.58l2.75 2.75Zm15.5 3.13l-2.12 1.22a1 1 0 0 0 1 1.74l2.12-1.22a1 1 0 1 0-1-1.74"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Layers(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m2.5 10.56l9 5.2a1 1 0 0 0 1 0l9-5.2a1 1 0 0 0 0-1.73l-9-5.2a1 1 0 0 0-1 0l-9 5.2a1 1 0 0 0 0 1.73M12 5.65l7 4l-7 4.05l-7-4.01Zm8.5 7.79L12 18.35l-8.5-4.91a1 1 0 0 0-1.37.36a1 1 0 0 0 .37 1.37l9 5.2a1 1 0 0 0 1 0l9-5.2a1 1 0 0 0 .37-1.37a1 1 0 0 0-1.37-.36"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayersAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2H9a1 1 0 0 0-1 1v4H6a1 1 0 0 0-1 1v4H3a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1v-2h4a1 1 0 0 0 1-1v-2h4a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1M10 20H4v-6h6Zm5-3h-3v-4a1 1 0 0 0-1-1H7V9h8Zm5-3h-3V8a1 1 0 0 0-1-1h-6V4h10Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayersSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.49 13.94l-.34.2a1 1 0 0 0-.35 1.37a1 1 0 0 0 .86.49a1 1 0 0 0 .51-.14l.34-.2a1 1 0 0 0-1-1.72Zm-8.84-7.58l.35-.21l7 4l-1.76 1a1 1 0 0 0 .5 1.87a1 1 0 0 0 .5-.13L21.5 11a1 1 0 0 0 0-1.74l-9-5.19a1 1 0 0 0-1 0l-.85.49a1 1 0 0 0 1 1.74ZM3.71 2.29a1 1 0 0 0-1.42 1.42l3.64 3.63l-3.43 2a1 1 0 0 0 0 1.74l9 5.2a1.09 1.09 0 0 0 .5.13a1.13 1.13 0 0 0 .5-.13l1.5-.88l1.45 1.46l-3.44 2l-8.51-4.93a1 1 0 0 0-1 1.74l9 5.2a1 1 0 0 0 1 0l4.41-2.55l3.38 3.39a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Zm8.29 12l-7-4.1l2.4-1.38l5.12 5.13Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftArrowFromLeft(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 11H5.41l2.3-2.29a1 1 0 1 0-1.42-1.42l-4 4a1 1 0 0 0-.21.33a1 1 0 0 0 0 .76a1 1 0 0 0 .21.33l4 4a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L5.41 13H17a1 1 0 0 0 0-2m4-7a1 1 0 0 0-1 1v14a1 1 0 0 0 2 0V5a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftArrowToLeft(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 11H9.41l2.3-2.29a1 1 0 1 0-1.42-1.42l-4 4a1 1 0 0 0-.21.33a1 1 0 0 0 0 .76a1 1 0 0 0 .21.33l4 4a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L9.41 13H21a1 1 0 0 0 0-2M3 3a1 1 0 0 0-1 1v16a1 1 0 0 0 2 0V4a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftIndent(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 7h18a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2m0 4h10a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2m18.77-1.31a1 1 0 0 0-1.41-.12l-2 1.66a1 1 0 0 0 0 1.54l2 1.66a1 1 0 0 0 .64.24a1 1 0 0 0 .77-.36a1 1 0 0 0-.13-1.41l-1.08-.9l1.08-.9a1 1 0 0 0 .13-1.41M21 17H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2M3 15h10a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftIndentAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 5a1 1 0 0 0-1 1v12a1 1 0 0 0 2 0V6a1 1 0 0 0-1-1m4 2h8a1 1 0 0 0 0-2h-8a1 1 0 0 0 0 2M5.77 9.69a1 1 0 0 0-1.41-.13l-2 1.67a1 1 0 0 0 0 1.54l2 1.67a1 1 0 0 0 1.41-.13a1 1 0 0 0-.13-1.41L4.56 12l1.08-.9a1 1 0 0 0 .13-1.41M21 9h-8a1 1 0 0 0 0 2h8a1 1 0 0 0 0-2m0 4h-8a1 1 0 0 0 0 2h8a1 1 0 0 0 0-2m0 4h-8a1 1 0 0 0 0 2h8a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftToRightTextDirection(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.42 17.62a1 1 0 0 0-.21-.33l-3-3a1 1 0 0 0-1.42 1.42l1.3 1.29H3.5a1 1 0 0 0 0 2h14.59l-1.3 1.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l3-3a1 1 0 0 0 .21-.33a1 1 0 0 0 0-.76M8.5 10v4a1 1 0 0 0 2 0V4h2v10a1 1 0 0 0 2 0V4h1a1 1 0 0 0 0-2h-7a4 4 0 0 0 0 8m0-6v4a2 2 0 0 1 0-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterChineseA(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 5h-7V4a1 1 0 0 0-2 0v1H4a1 1 0 0 0 0 2h11.882a14.493 14.493 0 0 1-3.94 7.952A14.426 14.426 0 0 1 8.663 9.67a1 1 0 0 0-1.889.66a16.414 16.414 0 0 0 3.68 5.958a14.299 14.299 0 0 1-5.768 2.735A1 1 0 0 0 4.899 21a1.018 1.018 0 0 0 .215-.023a16.297 16.297 0 0 0 6.831-3.319a16.387 16.387 0 0 0 6.842 3.319a1 1 0 0 0 .426-1.954a14.382 14.382 0 0 1-5.79-2.733A16.48 16.48 0 0 0 17.893 7H20a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterEnglishA(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.97 19.757L15.35 5.272A2.996 2.996 0 0 0 12.437 3h-.877a2.996 2.996 0 0 0-2.91 2.272L5.03 19.757a1 1 0 0 0 1.94.486L8.28 15h7.44l1.31 5.243a1 1 0 0 0 1.94-.486M8.78 13l1.811-7.242a.998.998 0 0 1 .97-.758h.878a.998.998 0 0 1 .97.758L15.219 13Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterHindiA(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.022 3h-5a1 1 0 0 0 0 2h1.5v6h-4.95a4.951 4.951 0 0 0 1.026-3a5 5 0 0 0-9.33-2.5a1 1 0 1 0 1.731 1A3 3 0 1 1 7.598 11a1 1 0 0 0 0 2a3 3 0 1 1-2.599 4.5a1 1 0 0 0-1.731 1a5 5 0 0 0 9.33-2.5a4.951 4.951 0 0 0-1.026-3h4.95v7a1 1 0 0 0 2 0V5h1.5a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterJapaneseA(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.472 11.997a4.865 4.865 0 0 0-4-2.204a5.592 5.592 0 0 0-.131-1.024a1 1 0 1 0-1.945.462a3.553 3.553 0 0 1 .082.601a13.577 13.577 0 0 0-3.257.822c.023-1.204.077-2.407.197-3.607l.038-.382A33.435 33.435 0 0 0 14.938 6l.12-.03a1 1 0 1 0-.48-1.94l-.122.03c-.922.23-1.856.404-2.794.55l.151-1.51a1 1 0 0 0-1.99-.2l-.196 1.96c-.934.083-1.87.14-2.809.14a1 1 0 0 0 0 2c.87 0 1.74-.046 2.607-.114a46.66 46.66 0 0 0-.203 4.698c-.134.073-.27.142-.403.22a16.407 16.407 0 0 0-1.949 1.31l-.022.018a13.74 13.74 0 0 0-2.649 2.7a3.004 3.004 0 0 0 2.947 4.72a9.74 9.74 0 0 0 2.837-1.014a.996.996 0 0 0 1.822-.703c-.025-.145-.036-.291-.058-.437a13.838 13.838 0 0 0 1.314-1.155a13.167 13.167 0 0 0 2.101-2.73c.023-.039.042-.078.065-.118c.118-.21.23-.422.332-.635c.053-.111.102-.222.151-.333a10.4 10.4 0 0 0 .329-.838c.032-.096.06-.191.09-.287c.05-.169.101-.337.141-.504l.005-.018A3.015 3.015 0 0 1 18.741 13c1.019 1.767-.963 4.977-4.417 7.154a1 1 0 1 0 1.067 1.692c4.499-2.836 6.683-7.069 5.08-9.849M6.796 18.583a1.005 1.005 0 0 1-.98-1.574a11.893 11.893 0 0 1 2.291-2.323l.027-.021a13.97 13.97 0 0 1 1.144-.793c.06 1.195.173 2.386.326 3.574a8.185 8.185 0 0 1-2.808 1.137M14.126 12a8.166 8.166 0 0 1-.316.78c-.056.12-.118.239-.18.358q-.145.278-.311.554c-.085.14-.172.279-.265.418a11.48 11.48 0 0 1-1.408 1.719c-.07.07-.143.133-.214.2q-.163-1.597-.211-3.202a12.513 12.513 0 0 1 2.94-.933z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LifeRing(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 11.05v-.33l-.09-.6l-.09-.39c0-.17-.08-.34-.13-.51s-.08-.27-.13-.4a2.17 2.17 0 0 1-.07-.24s0 0 0-.05a10.1 10.1 0 0 0-5.9-5.9s0 0-.05 0l-.23-.07l-.42-.13c-.15 0-.31-.08-.46-.12l-.46-.1l-.46-.07c-.16 0-.31 0-.48-.06s-.35 0-.52 0L12 2h-.91c-.17 0-.32 0-.48.06l-.46.07l-.46.1c-.15 0-.31.07-.46.12l-.42.13l-.23.07h-.05a10.1 10.1 0 0 0-5.9 5.9s0 0 0 .05a2.17 2.17 0 0 1-.07.24c0 .13-.09.26-.13.4s-.09.34-.13.51l-.09.39l-.09.6v2.56l.09.6l.09.39c0 .17.08.34.13.51s.08.27.13.4a2.17 2.17 0 0 1 .07.24a.43.43 0 0 1 0 .07a10 10 0 0 0 5.89 5.88s0 0 .05 0l.24.07l.4.13l.51.13l.39.09l.6.09h.33c.31 0 .63.05.95.05s.64 0 .95-.05h.33l.6-.09l.39-.09l.51-.13l.4-.13l.24-.07h.05a10 10 0 0 0 5.89-5.88a.43.43 0 0 1 0-.07c0-.08.05-.16.07-.24s.09-.26.13-.4s.09-.34.13-.51l.09-.39l.09-.6v-.33c0-.31.05-.63.05-.95s.09-.56.09-.87m-6.3-6.16a8 8 0 0 1 3.46 3.46l-2.86 1a5.14 5.14 0 0 0-1.64-1.64Zm-5.36-.7c.21-.05.41-.08.61-.11h.24a8.24 8.24 0 0 1 1.72 0h.24c.2 0 .4.06.61.11h.06l-1 2.86A4.49 4.49 0 0 0 12 7a4.4 4.4 0 0 0-.73.06l-1-2.86Zm-1.94.7l1 2.86a5.14 5.14 0 0 0-1.65 1.64l-2.86-1a8 8 0 0 1 3.46-3.5Zm-4.21 8.82a4.17 4.17 0 0 1-.1-.6v-.25a7.42 7.42 0 0 1 0-1.72v-.25a4.17 4.17 0 0 1 .1-.6s0 0 0-.06l2.86 1a4.47 4.47 0 0 0 0 1.46l-2.86 1zm4.16 5.4a8 8 0 0 1-3.46-3.46l2.86-1a5.14 5.14 0 0 0 1.64 1.64Zm5.36.7c-.21.05-.41.08-.61.11h-.24a8.24 8.24 0 0 1-1.72 0h-.24c-.2 0-.4-.06-.61-.11h-.06l1-2.86a4.47 4.47 0 0 0 1.46 0l1 2.86Zm-.67-5c-.17.06-.34.1-.5.14a2.73 2.73 0 0 1-1 0c-.16 0-.33-.08-.5-.14A3 3 0 0 1 9.2 13a3.23 3.23 0 0 1-.14-.51a2.63 2.63 0 0 1 0-1a3.23 3.23 0 0 1 .13-.49A3 3 0 0 1 11 9.2c.17-.06.34-.1.5-.14a2.73 2.73 0 0 1 1 0c.16 0 .33.08.5.14a3 3 0 0 1 1.8 1.8a3.23 3.23 0 0 1 .14.51a2.63 2.63 0 0 1 0 1a3.23 3.23 0 0 1-.14.51A3 3 0 0 1 13 14.8Zm2.61 4.31l-1-2.86a5.14 5.14 0 0 0 1.64-1.64l2.86 1a8 8 0 0 1-3.5 3.49ZM20 12.86v.25a4.17 4.17 0 0 1-.1.6s0 0 0 .06l-2.86-1a4.47 4.47 0 0 0 0-1.46l2.86-1v.06a4.17 4.17 0 0 1 .1.6v.25a7.42 7.42 0 0 1 0 1.72Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lightbulb(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.45 12.46a7 7 0 0 0-1-9.83a7.09 7.09 0 0 0-5.92-1.4a7 7 0 0 0-4 11.15a4.76 4.76 0 0 1 1.08 2.86v.29a2 2 0 0 0-.61 1.4v2a2 2 0 0 0 2 2v1a1 1 0 0 0 2 0v-1h2v1a1 1 0 0 0 2 0v-1a2 2 0 0 0 2-2v-2a2 2 0 0 0-.57-1.4v-.43a4.26 4.26 0 0 1 1.02-2.64M9 18.93v-2h6v2Zm6.89-7.72a6.18 6.18 0 0 0-1.46 3.72H9.56a6.67 6.67 0 0 0-1.5-3.78a5 5 0 0 1 2.84-8A5 5 0 0 1 17 8.07a4.92 4.92 0 0 1-1.11 3.14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightbulbAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.09 2.82a8 8 0 0 0-6.68-1.66a8 8 0 0 0-6.27 6.32a8.07 8.07 0 0 0 1.72 6.65A4.54 4.54 0 0 1 7 17v3a3 3 0 0 0 3 3h4a3 3 0 0 0 3-3v-2.81A5.17 5.17 0 0 1 18.22 14a8 8 0 0 0-1.13-11.2ZM15 20a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-1h6Zm1.67-7.24A7.13 7.13 0 0 0 15 17h-2v-3a1 1 0 0 0-2 0v3H9a6.5 6.5 0 0 0-1.6-4.16a6 6 0 0 1 3.39-9.72A6 6 0 0 1 18 9a5.89 5.89 0 0 1-1.33 3.76"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Line(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.78 9.46a.38.38 0 0 0-.38.38v1.67L12 9.65a.4.4 0 0 0-.33-.19a.38.38 0 0 0-.38.38v2.84a.38.38 0 0 0 .38.38a.38.38 0 0 0 .38-.38V11l1.39 1.91a.27.27 0 0 0 .15.11a.32.32 0 0 0 .14 0A.33.33 0 0 0 14 13l.1-.07a.39.39 0 0 0 .11-.27V9.84a.38.38 0 0 0-.43-.38M9.2 12.27H8.14V9.84a.38.38 0 0 0-.38-.38a.38.38 0 0 0-.38.38v2.84a.38.38 0 0 0 .38.38H9.2a.39.39 0 0 0 .39-.38a.39.39 0 0 0-.39-.41m1.11-2.81a.39.39 0 0 0-.39.38v2.84a.39.39 0 0 0 .39.38a.38.38 0 0 0 .38-.38V9.84a.38.38 0 0 0-.38-.38M17.91 2H6.09A4.1 4.1 0 0 0 2 6.09v11.82A4.1 4.1 0 0 0 6.09 22h11.82A4.1 4.1 0 0 0 22 17.91V6.09A4.1 4.1 0 0 0 17.91 2m.31 12.28a1.55 1.55 0 0 1-.13.17a5.5 5.5 0 0 1-.8.8c-2 1.87-5.36 4.11-5.81 3.76s.64-1.76-.53-2a1 1 0 0 1-.25 0c-3.44-.48-6-2.89-6-5.78c0-3.25 3.29-5.88 7.34-5.88s7.34 2.63 7.34 5.88a5 5 0 0 1-1.16 3.05m-1.71-4.81H15a.38.38 0 0 0-.38.38v2.84a.38.38 0 0 0 .38.38h1.48a.38.38 0 0 0 .38-.38a.38.38 0 0 0-.38-.38h-1.03v-.6h1.06a.39.39 0 0 0 .38-.39a.38.38 0 0 0-.38-.38h-1.06v-.61h1.06a.38.38 0 0 0 .38-.38a.38.38 0 0 0-.38-.48"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LineAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.71 3.29a1 1 0 0 0-1.42 0l-18 18a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l18-18a1 1 0 0 0 0-1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LineSpacing(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.29 9.71a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-2-2a1 1 0 0 0-.33-.21a1 1 0 0 0-.76 0a1 1 0 0 0-.33.21l-2 2a1 1 0 0 0 1.42 1.42l.29-.3v5.18l-.29-.3a1 1 0 0 0-1.42 1.42l2 2a1 1 0 0 0 .33.21a.94.94 0 0 0 .76 0a1 1 0 0 0 .33-.21l2-2a1 1 0 0 0-1.42-1.42l-.29.3V9.41ZM11 8h10a1 1 0 0 0 0-2H11a1 1 0 0 0 0 2m10 3H11a1 1 0 0 0 0 2h10a1 1 0 0 0 0-2m0 5H11a1 1 0 0 0 0 2h10a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Link(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10 17.55l-1.77 1.72a2.47 2.47 0 0 1-3.5-3.5l4.54-4.55a2.46 2.46 0 0 1 3.39-.09l.12.1a1 1 0 0 0 1.4-1.43a2.75 2.75 0 0 0-.18-.21a4.46 4.46 0 0 0-6.09.22l-4.6 4.55a4.48 4.48 0 0 0 6.33 6.33L11.37 19A1 1 0 0 0 10 17.55M20.69 3.31a4.49 4.49 0 0 0-6.33 0L12.63 5A1 1 0 0 0 14 6.45l1.73-1.72a2.47 2.47 0 0 1 3.5 3.5l-4.54 4.55a2.46 2.46 0 0 1-3.39.09l-.12-.1a1 1 0 0 0-1.4 1.43a2.75 2.75 0 0 0 .23.21a4.47 4.47 0 0 0 6.09-.22l4.55-4.55a4.49 4.49 0 0 0 .04-6.33"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkAdd(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12.11 15.39l-3.88 3.88a2.47 2.47 0 0 1-3.5 0a2.46 2.46 0 0 1 0-3.5l3.88-3.88a1 1 0 1 0-1.42-1.42l-3.88 3.89a4.48 4.48 0 0 0 6.33 6.33l3.89-3.88a1 1 0 0 0-1.42-1.42m-3.28-.22a1 1 0 0 0 .71.29a1 1 0 0 0 .71-.29l4.92-4.92a1 1 0 1 0-1.42-1.42l-4.92 4.92a1 1 0 0 0 0 1.42M21 18h-1v-1a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2h1v1a1 1 0 0 0 2 0v-1h1a1 1 0 0 0 0-2m-4.19-4.47l3.88-3.89a4.48 4.48 0 0 0-6.33-6.33l-3.89 3.88a1 1 0 1 0 1.42 1.42l3.88-3.88a2.47 2.47 0 0 1 3.5 0a2.46 2.46 0 0 1 0 3.5l-3.88 3.88a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12.11 15.39l-3.88 3.88a2.52 2.52 0 0 1-3.5 0a2.47 2.47 0 0 1 0-3.5l3.88-3.88a1 1 0 1 0-1.42-1.42l-3.88 3.89a4.48 4.48 0 0 0 6.33 6.33l3.89-3.88a1 1 0 0 0-1.42-1.42m8.58-12.08a4.49 4.49 0 0 0-6.33 0l-3.89 3.88a1 1 0 1 0 1.42 1.42l3.88-3.88a2.52 2.52 0 0 1 3.5 0a2.47 2.47 0 0 1 0 3.5l-3.88 3.88a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l3.88-3.89a4.49 4.49 0 0 0 0-6.33M8.83 15.17a1 1 0 0 0 .71.29a1 1 0 0 0 .71-.29l4.92-4.92a1 1 0 1 0-1.42-1.42l-4.92 4.92a1 1 0 0 0 0 1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkBroken(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.76 10.59a1 1 0 0 0 .26-2l-1.76-.44a1 1 0 1 0-.52 1.93l1.76.47a.78.78 0 0 0 .26.04M8.62 5a1 1 0 0 0 1 .74a.82.82 0 0 0 .26 0a1 1 0 0 0 .7-1.22l-.47-1.76a1 1 0 1 0-1.93.52Zm4.83 10A1 1 0 0 0 12 15l-3.5 3.56a2.21 2.21 0 0 1-3.06 0a2.15 2.15 0 0 1 0-3.06L9 12a1 1 0 1 0-1.41-1.41L4 14.08A4.17 4.17 0 1 0 9.92 20l3.53-3.53a1 1 0 0 0 0-1.47M5.18 6.59a1 1 0 0 0 .7.29a1 1 0 0 0 .71-.29a1 1 0 0 0 0-1.41L5.3 3.89A1 1 0 0 0 3.89 5.3Zm16.08 7.33l-1.76-.47a1 1 0 1 0-.5 1.93l1.76.47h.26a1 1 0 0 0 .26-2ZM15.38 19a1 1 0 0 0-1.23-.7a1 1 0 0 0-.7 1.22l.47 1.76a1 1 0 0 0 1 .74a1.15 1.15 0 0 0 .26 0a1 1 0 0 0 .71-1.23Zm3.44-1.57a1 1 0 0 0-1.41 1.41l1.29 1.29a1 1 0 0 0 1.41 0a1 1 0 0 0 0-1.41ZM21.2 7a4.16 4.16 0 0 0-7.12-3l-3.53 3.56A1 1 0 1 0 12 9l3.5-3.56a2.21 2.21 0 0 1 3.06 0a2.15 2.15 0 0 1 0 3.06L15 12a1 1 0 0 0 0 1.41a1 1 0 0 0 1.41 0L20 9.92A4.19 4.19 0 0 0 21.2 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkH(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 12a1 1 0 0 0 1 1h6a1 1 0 0 0 0-2H9a1 1 0 0 0-1 1m2 3H7a3 3 0 0 1 0-6h3a1 1 0 0 0 0-2H7a5 5 0 0 0 0 10h3a1 1 0 0 0 0-2m7-8h-3a1 1 0 0 0 0 2h3a3 3 0 0 1 0 6h-3a1 1 0 0 0 0 2h3a5 5 0 0 0 0-10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Linkedin(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.47 2H3.53a1.45 1.45 0 0 0-1.47 1.43v17.14A1.45 1.45 0 0 0 3.53 22h16.94a1.45 1.45 0 0 0 1.47-1.43V3.43A1.45 1.45 0 0 0 20.47 2M8.09 18.74h-3v-9h3ZM6.59 8.48a1.56 1.56 0 1 1 0-3.12a1.57 1.57 0 1 1 0 3.12m12.32 10.26h-3v-4.83c0-1.21-.43-2-1.52-2A1.65 1.65 0 0 0 12.85 13a2 2 0 0 0-.1.73v5h-3v-9h3V11a3 3 0 0 1 2.71-1.5c2 0 3.45 1.29 3.45 4.06Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkedinAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.5 8.999a5.419 5.419 0 0 0-2.565.645A1 1 0 0 0 14 8.999h-4a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-5.5a1 1 0 1 1 2 0v5.5a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-7.5a5.507 5.507 0 0 0-5.5-5.5m3.5 12h-2v-4.5a3 3 0 1 0-6 0v4.5h-2v-10h2v.703a1 1 0 0 0 1.781.625A3.483 3.483 0 0 1 21 14.5Zm-14-12H3a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-12a1 1 0 0 0-1-1m-1 12H4v-10h2ZM5.015 1.542a3.233 3.233 0 1 0-.057 6.457h.028a3.233 3.233 0 1 0 .029-6.457m-.029 4.457h-.028a1.222 1.222 0 0 1-1.37-1.228c0-.747.56-1.229 1.427-1.229A1.234 1.234 0 0 1 6.41 4.771c0 .746-.56 1.228-1.425 1.228"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Linux(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.7 17.6c-.1-.2-.2-.4-.2-.6c0-.4-.2-.7-.5-1c-.1-.1-.3-.2-.4-.2c.6-1.8-.3-3.6-1.3-4.9c-.8-1.2-2-2.1-1.9-3.7c0-1.9.2-5.4-3.3-5.1c-3.6.2-2.6 3.9-2.7 5.2c0 1.1-.5 2.2-1.3 3.1c-.2.2-.4.5-.5.7c-1 1.2-1.5 2.8-1.5 4.3c-.2.2-.4.4-.5.6c-.1.1-.2.2-.2.3c-.1.1-.3.2-.5.3c-.4.1-.7.3-.9.7c-.1.3-.2.7-.1 1.1c.1.2.1.4 0 .7c-.2.4-.2.9 0 1.4c.3.4.8.5 1.5.6c.5 0 1.1.2 1.6.4c.5.3 1.1.5 1.7.5c.3 0 .7-.1 1-.2c.3-.2.5-.4.6-.7c.4 0 1-.2 1.7-.2c.6 0 1.2.2 2 .1c0 .1 0 .2.1.3c.2.5.7.9 1.3 1h.2c.8-.1 1.6-.5 2.1-1.1c.4-.4.9-.7 1.4-.9c.6-.3 1-.5 1.1-1c.1-.7-.1-1.1-.5-1.7M12.8 4.8c.6.1 1.1.6 1 1.2c0 .3-.1.6-.3.9h-.1c-.2-.1-.3-.1-.4-.2c.1-.1.1-.3.2-.5c0-.4-.2-.7-.4-.7c-.3 0-.5.3-.5.7v.1c-.1-.1-.3-.1-.4-.2V6c-.1-.5.3-1.1.9-1.2m-.3 2c.1.1.3.2.4.2c.1 0 .3.1.4.2c.2.1.4.2.4.5s-.3.6-.9.8c-.2.1-.3.1-.4.2c-.3.2-.6.3-1 .3c-.3 0-.6-.2-.8-.4c-.1-.1-.2-.2-.4-.3c-.1-.1-.3-.3-.4-.6c0-.1.1-.2.2-.3c.3-.2.4-.3.5-.4l.1-.1c.2-.3.6-.5 1-.5c.3.1.6.2.9.4M10.4 5c.4 0 .7.4.8 1.1v.2c-.1 0-.3.1-.4.2v-.2c0-.3-.2-.6-.4-.5c-.2 0-.3.3-.3.6c0 .2.1.3.2.4c0 0-.1.1-.2.1c-.2-.2-.4-.5-.4-.8c0-.6.3-1.1.7-1.1m-1 16.1c-.7.3-1.6.2-2.2-.2c-.6-.3-1.1-.4-1.8-.4c-.5-.1-1-.1-1.1-.3c-.1-.2-.1-.5.1-1c.1-.3.1-.6 0-.9c-.1-.3-.1-.5 0-.8c.1-.3.3-.4.6-.5c.3-.1.5-.2.7-.4c.1-.1.2-.2.3-.4c.3-.4.5-.6.8-.6c.6.1 1.1 1 1.5 1.9c.2.3.4.7.7 1c.4.5.9 1.2.9 1.6c0 .5-.2.8-.5 1m4.9-2.2c0 .1 0 .1-.1.2c-1.2.9-2.8 1-4.1.3l-.6-.9c.9-.1.7-1.3-1.2-2.5c-2-1.3-.6-3.7.1-4.8c.1-.1.1 0-.3.8c-.3.6-.9 2.1-.1 3.2c0-.8.2-1.6.5-2.4c.7-1.3 1.2-2.8 1.5-4.3c.1.1.1.1.2.1c.1.1.2.2.3.2c.2.3.6.4.9.4h.1c.4 0 .8-.1 1.1-.4c.1-.1.2-.2.4-.2c.3-.1.6-.3.9-.6c.4 1.3.8 2.5 1.4 3.6c.4.8.7 1.6.9 2.5c.3 0 .7.1 1 .3c.8.4 1.1.7 1 1.2H18c0-.3-.2-.6-.9-.9c-.7-.3-1.3-.3-1.5.4c-.1 0-.2.1-.3.1c-.8.4-.8 1.5-.9 2.6c.1.4 0 .7-.1 1.1m4.6.6c-.6.2-1.1.6-1.5 1.1c-.4.6-1.1 1-1.9.9c-.4 0-.8-.3-.9-.7c-.1-.6-.1-1.2.2-1.8c.1-.4.2-.7.3-1.1c.1-1.2.1-1.9.6-2.2c0 .5.3.8.7 1c.5 0 1-.1 1.4-.5h.2c.3 0 .5 0 .7.2c.2.2.3.5.3.7c0 .3.2.6.3.9c.5.5.5.8.5.9c-.1.2-.5.4-.9.6m-9-12c-.1 0-.1 0-.1.1c0 0 0 .1.1.1s.1.1.1.1c.3.4.8.6 1.4.7c.5-.1 1-.2 1.5-.6l.6-.3c.1 0 .1-.1.1-.1c0-.1 0-.1-.1-.1c-.2.1-.5.2-.7.3c-.4.3-.9.5-1.4.5c-.5 0-.9-.3-1.2-.6c-.1 0-.2-.1-.3-.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LiraSign(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 12a1 1 0 0 0-1 1a7.008 7.008 0 0 1-7 7v-8.865l5.217-1.159a1 1 0 0 0-.434-1.952L10 9.087V7.135l5.217-1.159a1 1 0 1 0-.434-1.952L10 5.087V3a1 1 0 0 0-2 0v2.531l-2.217.493a1 1 0 1 0 .434 1.952L8 7.58v1.95l-2.217.493a1 1 0 1 0 .434 1.952L8 11.58V21a1 1 0 0 0 1 1h1a9.01 9.01 0 0 0 9-9a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListOl(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 20H4v-.1c0-.5.4-.9.9-.9c1.4 0 2.6-.9 3-2.2c.4-1.6-.5-3.3-2.1-3.7c-1.3-.4-2.7.2-3.4 1.4c-.3.5-.1 1.1.4 1.4c.5.3 1.1.1 1.4-.4c.3-.5.9-.6 1.4-.3c.1.1.2.1.2.2c.2.3.2.6.2.9c-.2.4-.6.7-1 .7c-1.7 0-3 1.3-3 2.9V21c0 .6.4 1 1 1h4c.6 0 1-.4 1-1s-.4-1-1-1m4-13h10c.6 0 1-.4 1-1s-.4-1-1-1H11c-.6 0-1 .4-1 1s.4 1 1 1M7 9H6V3c0-.6-.4-1-1-1s-1 .4-1 1v1H3c-.6 0-1 .4-1 1s.4 1 1 1h1v3H3c-.6 0-1 .4-1 1s.4 1 1 1h4c.6 0 1-.4 1-1s-.4-1-1-1m14 7H11c-.6 0-1 .4-1 1s.4 1 1 1h10c.6 0 1-.4 1-1s-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListOlAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 20H4v-.1c0-.5.4-.9.9-.9c1.4 0 2.6-.9 3-2.2c.4-1.6-.5-3.3-2.1-3.7c-1.3-.4-2.7.2-3.4 1.4c-.3.5-.1 1.1.4 1.4c.5.3 1.1.1 1.4-.4c.3-.5.9-.6 1.4-.3c.1.1.2.1.2.2c.2.3.2.6.2.9c-.2.4-.6.7-1 .7c-1.7 0-3 1.3-3 2.9V21c0 .6.4 1 1 1h4c.6 0 1-.4 1-1s-.4-1-1-1M7 9H6V3c0-.6-.4-1-1-1s-1 .4-1 1v1H3c-.6 0-1 .4-1 1s.4 1 1 1h1v3H3c-.6 0-1 .4-1 1s.4 1 1 1h4c.6 0 1-.4 1-1s-.4-1-1-1m4-3h10c.6 0 1-.4 1-1s-.4-1-1-1H11c-.6 0-1 .4-1 1s.4 1 1 1m10 14H11c-.6 0-1 .4-1 1s.4 1 1 1h10c.6 0 1-.4 1-1s-.4-1-1-1m0-11H11c-.6 0-1 .4-1 1s.4 1 1 1h10c.6 0 1-.4 1-1s-.4-1-1-1m0 6H11c-.6 0-1 .4-1 1s.4 1 1 1h10c.6 0 1-.4 1-1s-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListUiAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.5 6a1 1 0 1 0 1 1a1 1 0 0 0-1-1m4 2h14a1 1 0 0 0 0-2h-14a1 1 0 0 0 0 2m0 3a1 1 0 1 0 1 1a1 1 0 0 0-1-1m4 5a1 1 0 1 0 1 1a1 1 0 0 0-1-1m10-5h-10a1 1 0 0 0 0 2h10a1 1 0 0 0 0-2m0 5h-6a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListUl(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.71 16.29a1 1 0 0 0-.33-.21a1 1 0 0 0-.76 0a1 1 0 0 0-.33.21a1 1 0 0 0-.21.33a1 1 0 0 0 .21 1.09a1.15 1.15 0 0 0 .33.21a.94.94 0 0 0 .76 0a1.15 1.15 0 0 0 .33-.21a1 1 0 0 0 .21-1.09a1 1 0 0 0-.21-.33M7 8h14a1 1 0 0 0 0-2H7a1 1 0 0 0 0 2m-3.29 3.29a1 1 0 0 0-1.09-.21a1.15 1.15 0 0 0-.33.21a1 1 0 0 0-.21.33a.94.94 0 0 0 0 .76a1.15 1.15 0 0 0 .21.33a1.15 1.15 0 0 0 .33.21a.94.94 0 0 0 .76 0a1.15 1.15 0 0 0 .33-.21a1.15 1.15 0 0 0 .21-.33a.94.94 0 0 0 0-.76a1 1 0 0 0-.21-.33M21 11H7a1 1 0 0 0 0 2h14a1 1 0 0 0 0-2M3.71 6.29a1 1 0 0 0-.33-.21a1 1 0 0 0-1.09.21a1.15 1.15 0 0 0-.21.33a.94.94 0 0 0 0 .76a1.15 1.15 0 0 0 .21.33a1.15 1.15 0 0 0 .33.21a1 1 0 0 0 1.09-.21a1.15 1.15 0 0 0 .21-.33a.94.94 0 0 0 0-.76a1.15 1.15 0 0 0-.21-.33M21 16H7a1 1 0 0 0 0 2h14a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationArrow(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.15 2.86a2.89 2.89 0 0 0-3-.71L4 6.88a2.9 2.9 0 0 0-.12 5.47l5.24 2a.93.93 0 0 1 .53.52l2 5.25A2.87 2.87 0 0 0 14.36 22h.07a2.88 2.88 0 0 0 2.69-2l4.73-14.17a2.89 2.89 0 0 0-.7-2.97M20 5.2l-4.78 14.18a.88.88 0 0 1-.84.62a.92.92 0 0 1-.87-.58l-2-5.25a2.91 2.91 0 0 0-1.67-1.68l-5.25-2A.9.9 0 0 1 4 9.62a.88.88 0 0 1 .62-.84L18.8 4.05A.91.91 0 0 1 20 5.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationArrowAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.68 17.65l-7-14a3 3 0 0 0-5.36 0l-7 14a3 3 0 0 0 3.9 4.08l5.37-2.4a1.06 1.06 0 0 1 .82 0l5.37 2.4a3 3 0 0 0 3.9-4.08m-2 2a1 1 0 0 1-1.13.22l-5.37-2.39a3 3 0 0 0-2.44 0L5.41 19.9a1 1 0 0 1-1.3-1.35l7-14a1 1 0 0 1 1.78 0l7 14a1 1 0 0 1-.17 1.13Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationPinAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 10.8a2 2 0 1 0-2-2a2 2 0 0 0 2 2m-.71 6.91a1 1 0 0 0 1.42 0l4.09-4.1a6.79 6.79 0 1 0-9.6 0ZM7.23 8.34a4.81 4.81 0 0 1 2.13-3.55a4.81 4.81 0 0 1 5.28 0a4.82 4.82 0 0 1 .75 7.41L12 15.59L8.61 12.2a4.77 4.77 0 0 1-1.38-3.86M19 20H5a1 1 0 0 0 0 2h14a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationPoint(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 4.48a8.45 8.45 0 0 0-12 12l5.27 5.28a1 1 0 0 0 1.42 0L18 16.43a8.45 8.45 0 0 0 0-11.95M16.57 15L12 19.59L7.43 15a6.46 6.46 0 1 1 9.14 0M9 7.41a4.32 4.32 0 0 0 0 6.1a4.31 4.31 0 0 0 7.36-3a4.24 4.24 0 0 0-1.26-3.05A4.3 4.3 0 0 0 9 7.41m4.69 4.68a2.33 2.33 0 1 1 .67-1.63a2.33 2.33 0 0 1-.72 1.63Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lock(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 9V7A5 5 0 0 0 7 7v2a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-7a3 3 0 0 0-3-3M9 7a3 3 0 0 1 6 0v2H9Zm9 12a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1v-7a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockAccess(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2h-6a1 1 0 0 0 0 2h5v5a1 1 0 0 0 2 0V3a1 1 0 0 0-1-1m0 12a1 1 0 0 0-1 1v5h-5a1 1 0 0 0 0 2h6a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1m-9-8a3 3 0 0 0-3 3v1a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2V9a3 3 0 0 0-3-3m-1 3a1 1 0 0 1 2 0v1h-2Zm4 7H9v-4h6ZM3 10a1 1 0 0 0 1-1V4h5a1 1 0 0 0 0-2H3a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1m6 10H4v-5a1 1 0 0 0-2 0v6a1 1 0 0 0 1 1h6a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 13a1 1 0 0 0-1 1v3a1 1 0 0 0 2 0v-3a1 1 0 0 0-1-1m5-4V7A5 5 0 0 0 7 7v2a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-7a3 3 0 0 0-3-3M9 7a3 3 0 0 1 6 0v2H9Zm9 12a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1v-7a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockOpenAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 13a1.49 1.49 0 0 0-1 2.61V17a1 1 0 0 0 2 0v-1.39A1.49 1.49 0 0 0 12 13m5-4H9V7a3 3 0 0 1 5.12-2.13a3.08 3.08 0 0 1 .78 1.38a1 1 0 1 0 1.94-.5a5.09 5.09 0 0 0-1.31-2.29A5 5 0 0 0 7 7v2a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-7a3 3 0 0 0-3-3m1 10a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1v-7a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.84 5.38a2 2 0 0 1 2.57.21A2 2 0 0 1 14 7v3a1 1 0 0 0 1 1h1a1 1 0 0 1 1 1v.34a1 1 0 0 0 2 0V12a3 3 0 0 0-3-3V7a4 4 0 0 0-1.17-2.83a4.06 4.06 0 0 0-5.19-.39a1 1 0 1 0 1.2 1.6m10.87 14.91l-18-18a1 1 0 0 0-1.42 1.42L7.62 9A3 3 0 0 0 5 12v6a3 3 0 0 0 3 3h8a3 3 0 0 0 2.39-1.2l1.9 1.91a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42M16 19H8a1 1 0 0 1-1-1v-6a1 1 0 0 1 1-1h1.59l2.07 2.07A1 1 0 0 0 11 14v2a1 1 0 0 0 2 0v-1.59l3.93 3.93A1 1 0 0 1 16 19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lottiefiles(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 6c-2.8 0-4.4 2.8-5.9 5.5C9.9 13.8 8.7 16 7 16c-.6 0-1 .4-1 1s.4 1 1 1c2.8 0 4.4-2.8 5.9-5.5C14.1 10.2 15.3 8 17 8c.6 0 1-.4 1-1s-.4-1-1-1m2-4H5C3.3 2 2 3.3 2 5v14c0 1.7 1.3 3 3 3h14c1.7 0 3-1.3 3-3V5c0-1.7-1.3-3-3-3m1 17c0 .6-.4 1-1 1H5c-.6 0-1-.4-1-1V5c0-.6.4-1 1-1h14c.6 0 1 .4 1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LottiefilesAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 2H5a3.003 3.003 0 0 0-3 3v14a3.003 3.003 0 0 0 3 3h14a3.003 3.003 0 0 0 3-3V5a3.003 3.003 0 0 0-3-3m-2 6c-1.66 0-2.856 2.177-4.124 4.482C11.384 15.195 9.841 18 7 18a1 1 0 0 1 0-2c1.66 0 2.856-2.177 4.124-4.482C12.616 8.805 14.159 6 17 6a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LuggageCart(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 13.5v2a1 1 0 0 0 1 1h10a3 3 0 0 0 6 0h2a1 1 0 0 0 1-1v-8a3 3 0 0 0-3-3H9a3 3 0 0 0-3 3v7H4v-1a1 1 0 0 0-2 0m13 3a1 1 0 1 1 1 1a1 1 0 0 1-1-1m-7-6h12v4h-1.78a3 3 0 0 0-4.44 0H8Zm0-3a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v1H8Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mailbox(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 12h2a1 1 0 0 0 0-2H8a1 1 0 0 0 0 2m9-6h-5V4h1a1 1 0 0 0 0-2h-2a1 1 0 0 0-1 1v3H7a4 4 0 0 0-4 4v6a1 1 0 0 0 1 1h6v4a1 1 0 0 0 2 0v-4h8a1 1 0 0 0 1-1v-6a4 4 0 0 0-4-4m-4 4v5H5v-5a2 2 0 0 1 2-2h6.56a3.91 3.91 0 0 0-.56 2m6 5h-4v-5a2 2 0 0 1 4 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailboxAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 13h2a1 1 0 0 0 0-2h-2a1 1 0 0 0 0 2m8 7h-1V9h1a1 1 0 0 0 0-2h-1.09A6 6 0 0 0 6.09 7H5a1 1 0 0 0 0 2h1v11H5a1 1 0 0 0 0 2h14a1 1 0 0 0 0-2M12 4a4 4 0 0 1 3.86 3H8.14A4 4 0 0 1 12 4m4 16H8v-2h8Zm0-4H8V9h8Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Map(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.32 5.05l-6-2h-.07a.7.7 0 0 0-.14 0h-.43L9 5L3.32 3.05a1 1 0 0 0-.9.14A1 1 0 0 0 2 4v14a1 1 0 0 0 .68.95l6 2a1 1 0 0 0 .62 0l5.7-1.9L20.68 21a1.19 1.19 0 0 0 .32 0a.94.94 0 0 0 .58-.19A1 1 0 0 0 22 20V6a1 1 0 0 0-.68-.95M8 18.61l-4-1.33V5.39l4 1.33Zm6-1.33l-4 1.33V6.72l4-1.33Zm6 1.33l-4-1.33V5.39l4 1.33Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarker(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a8 8 0 0 0-8 8c0 5.4 7.05 11.5 7.35 11.76a1 1 0 0 0 1.3 0C13 21.5 20 15.4 20 10a8 8 0 0 0-8-8m0 17.65c-2.13-2-6-6.31-6-9.65a6 6 0 0 1 12 0c0 3.34-3.87 7.66-6 9.65M12 6a4 4 0 1 0 4 4a4 4 0 0 0-4-4m0 6a2 2 0 1 1 2-2a2 2 0 0 1-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarkerAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.46 9.63A8.5 8.5 0 0 0 7.3 3.36a8.56 8.56 0 0 0-3.76 6.27A8.46 8.46 0 0 0 6 16.46l5.3 5.31a1 1 0 0 0 1.42 0L18 16.46a8.46 8.46 0 0 0 2.46-6.83m-3.86 5.42l-4.6 4.6l-4.6-4.6a6.49 6.49 0 0 1-1.87-5.22A6.57 6.57 0 0 1 8.42 5a6.47 6.47 0 0 1 7.16 0a6.57 6.57 0 0 1 2.89 4.81a6.49 6.49 0 0 1-1.87 5.24M12 6a4.5 4.5 0 1 0 4.5 4.5A4.51 4.51 0 0 0 12 6m0 7a2.5 2.5 0 1 1 2.5-2.5A2.5 2.5 0 0 1 12 13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarkerEdit(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.46 9.63A8.5 8.5 0 1 0 6 16.46l5.3 5.31a1 1 0 0 0 1.42 0L18 16.46a8.46 8.46 0 0 0 2.46-6.83m-3.86 5.42l-4.6 4.6l-4.6-4.6a6.49 6.49 0 0 1-1.87-5.22A6.57 6.57 0 0 1 8.42 5a6.47 6.47 0 0 1 7.16 0a6.57 6.57 0 0 1 2.89 4.81a6.49 6.49 0 0 1-1.87 5.24m-2.81-8.8a1 1 0 0 0-1.42 0L8.79 9.83a1 1 0 0 0-.29.7V13a1 1 0 0 0 1 1h2.42a1 1 0 0 0 .71-.29l3.58-3.58a1 1 0 0 0 0-1.41ZM11.51 12h-1v-1l2.58-2.58l1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarkerInfo(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 10a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0v-2a1 1 0 0 0-1-1m8.46-.32A8.5 8.5 0 0 0 7.3 3.36a8.56 8.56 0 0 0-3.76 6.27A8.46 8.46 0 0 0 6 16.46l5.3 5.31a1 1 0 0 0 1.42 0L18 16.46a8.46 8.46 0 0 0 2.46-6.83Zm-3.86 5.37l-4.6 4.6l-4.6-4.6a6.49 6.49 0 0 1-1.87-5.22A6.57 6.57 0 0 1 8.42 5a6.47 6.47 0 0 1 7.16 0a6.57 6.57 0 0 1 2.89 4.81a6.49 6.49 0 0 1-1.87 5.24m-3.68-7.48a.56.56 0 0 0-.09-.17l-.12-.15A1 1 0 0 0 11.8 7h-.18l-.18.09l-.15.13l-.12.15a.56.56 0 0 0-.09.17a.6.6 0 0 0-.06.19A1.23 1.23 0 0 0 11 8a.88.88 0 0 0 .08.39a1.11 1.11 0 0 0 .21.32a1.06 1.06 0 0 0 .33.22a1.07 1.07 0 0 0 .76 0a1.19 1.19 0 0 0 .33-.22a1.11 1.11 0 0 0 .21-.32A1 1 0 0 0 13 8a1.23 1.23 0 0 0 0-.19a.6.6 0 0 0-.08-.24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarkerMinus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 9.45h-4a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2m6.46.18A8.5 8.5 0 1 0 6 16.46l5.3 5.31a1 1 0 0 0 1.42 0L18 16.46a8.46 8.46 0 0 0 2.46-6.83m-3.86 5.42l-4.6 4.6l-4.6-4.6a6.49 6.49 0 0 1-1.87-5.22A6.57 6.57 0 0 1 8.42 5a6.47 6.47 0 0 1 7.16 0a6.57 6.57 0 0 1 2.89 4.81a6.49 6.49 0 0 1-1.87 5.24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarkerPlus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 9.45h-1v-1a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2h1v1a1 1 0 0 0 2 0v-1h1a1 1 0 0 0 0-2m6.46.18A8.5 8.5 0 1 0 6 16.46l5.3 5.31a1 1 0 0 0 1.42 0L18 16.46a8.46 8.46 0 0 0 2.46-6.83m-3.86 5.42l-4.6 4.6l-4.6-4.6a6.49 6.49 0 0 1-1.87-5.22A6.57 6.57 0 0 1 8.42 5a6.47 6.47 0 0 1 7.16 0a6.57 6.57 0 0 1 2.89 4.81a6.49 6.49 0 0 1-1.87 5.24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarkerQuestion(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12.44 13.11l-.17-.11a1 1 0 0 0-1.09.22a.87.87 0 0 0-.22.32a1 1 0 0 0-.08.39a1 1 0 0 0 .08.38a1.07 1.07 0 0 0 .54.54a1 1 0 0 0 .38.08a1.09 1.09 0 0 0 .39-.08a1 1 0 0 0 .32-.22a1 1 0 0 0 0-1.41ZM11.88 6A2.75 2.75 0 0 0 9.5 7.32a1 1 0 1 0 1.73 1a.77.77 0 0 1 .65-.32a.75.75 0 1 1 0 1.5a1 1 0 1 0 0 2a2.75 2.75 0 1 0 0-5.5m8.58 3.68A8.5 8.5 0 0 0 7.3 3.36a8.56 8.56 0 0 0-3.76 6.27A8.46 8.46 0 0 0 6 16.46l5.3 5.31a1 1 0 0 0 1.42 0L18 16.46a8.46 8.46 0 0 0 2.46-6.83Zm-3.86 5.37l-4.6 4.6l-4.6-4.6a6.49 6.49 0 0 1-1.87-5.22A6.57 6.57 0 0 1 8.42 5a6.47 6.47 0 0 1 7.16 0a6.57 6.57 0 0 1 2.89 4.81a6.49 6.49 0 0 1-1.87 5.24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarkerShield(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.41 6.53a2.24 2.24 0 0 1-1.82-.39a1 1 0 0 0-1.18 0a2.24 2.24 0 0 1-1.82.39a1 1 0 0 0-.84.2a1 1 0 0 0-.37.78v2.9A4.14 4.14 0 0 0 10 13.74l1.37 1a1 1 0 0 0 1.18 0l1.37-1a4.14 4.14 0 0 0 1.66-3.33v-2.9a1 1 0 0 0-.37-.78a1 1 0 0 0-.8-.2m-.79 3.88a2.15 2.15 0 0 1-.85 1.73l-.77.57l-.77-.57a2.15 2.15 0 0 1-.85-1.73V8.57A4.22 4.22 0 0 0 12 8.12a4.22 4.22 0 0 0 1.62.45Zm6.84-.78A8.5 8.5 0 0 0 7.3 3.36a8.56 8.56 0 0 0-3.76 6.27A8.46 8.46 0 0 0 6 16.46l5.3 5.31a1 1 0 0 0 1.42 0L18 16.46a8.46 8.46 0 0 0 2.46-6.83m-3.86 5.42l-4.6 4.6l-4.6-4.6a6.49 6.49 0 0 1-1.87-5.22A6.57 6.57 0 0 1 8.42 5a6.47 6.47 0 0 1 7.16 0a6.57 6.57 0 0 1 2.89 4.81a6.49 6.49 0 0 1-1.87 5.24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarkerSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.63 5.49a6 6 0 0 1 7.21 7.2a1 1 0 0 0 .74 1.21a.9.9 0 0 0 .23 0a1 1 0 0 0 1-.76a8 8 0 0 0-9.61-9.62a1 1 0 0 0 .46 2Zm11.08 14.58l-4.27-4.27L3.71 2.07a1 1 0 0 0-1.42 0a1 1 0 0 0 0 1.41L5.5 6.69A8 8 0 0 0 6.34 17l4.95 4.95a1 1 0 0 0 1.42 0l4-4l3.56 3.56a1 1 0 0 0 1.42-1.41Zm-9.59-6.76a2 2 0 0 1-1.53-.57a2 2 0 0 1-.59-1.53Zm-.12 6.5l-4.24-4.24a6 6 0 0 1-.82-7.44L8.41 9.6a4 4 0 0 0 .76 4.55A4 4 0 0 0 12 15.33a3.93 3.93 0 0 0 1.73-.41l1.58 1.58Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapPin(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.37 12.79a1 1 0 0 0-.74 1.86C17.09 15.23 18 16.13 18 17c0 1.42-2.46 3-6 3s-6-1.58-6-3c0-.87.91-1.77 2.37-2.35a1 1 0 0 0-.74-1.86C5.36 13.69 4 15.26 4 17c0 2.8 3.51 5 8 5s8-2.2 8-5c0-1.74-1.36-3.31-3.63-4.21M11 9.86V17a1 1 0 0 0 2 0V9.86a4 4 0 1 0-2 0M12 4a2 2 0 1 1-2 2a2 2 0 0 1 2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapPinAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 11.9V17a1 1 0 0 0 2 0v-5.1a5 5 0 1 0-2 0M12 4a3 3 0 1 1-3 3a3 3 0 0 1 3-3m4.21 10.42a1 1 0 1 0-.42 2C18.06 16.87 19 17.68 19 18c0 .58-2.45 2-7 2s-7-1.42-7-2c0-.32.94-1.13 3.21-1.62a1 1 0 1 0-.42-2C4.75 15.08 3 16.39 3 18c0 2.63 4.53 4 9 4s9-1.37 9-4c0-1.61-1.75-2.92-4.79-3.58"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mars(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.42 4.12a1 1 0 0 0-.54-.54a1 1 0 0 0-.38-.08h-4a1 1 0 0 0 0 2h1.59l-2.4 2.4a7 7 0 1 0 1.41 1.41l2.4-2.4V8.5a1 1 0 0 0 2 0v-4a1 1 0 0 0-.08-.38M14 17a5 5 0 1 1 0-7a5 5 0 0 1 0 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MasterCard(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.265 5.274a6.681 6.681 0 0 0-3.273.855a6.728 6.728 0 1 0 0 11.745a6.726 6.726 0 1 0 3.273-12.6m-5.028 11.183a4.667 4.667 0 0 1-1.518.273a4.728 4.728 0 0 1 0-9.456a4.667 4.667 0 0 1 1.518.273a6.687 6.687 0 0 0 0 8.91m1.755-1.057a4.695 4.695 0 0 1 0-6.796a4.695 4.695 0 0 1 0 6.796m3.273 1.33a4.667 4.667 0 0 1-1.519-.273a6.687 6.687 0 0 0 0-8.91a4.667 4.667 0 0 1 1.519-.273a4.728 4.728 0 0 1 0 9.456"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MaximizeLeft(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.41 4H7a1 1 0 0 0 0-2H3a1 1 0 0 0-.38.08a1 1 0 0 0-.54.54A1 1 0 0 0 2 3v4a1 1 0 0 0 2 0V5.41l5.79 5.8a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM21 16a1 1 0 0 0-1 1v1.59l-5.79-5.8a1 1 0 0 0-1.42 1.42l5.8 5.79H17a1 1 0 0 0 0 2h4a1 1 0 0 0 .38-.08a1 1 0 0 0 .54-.54A1 1 0 0 0 22 21v-4a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Medal(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.38 5.76a1 1 0 0 0-.47-.61l-5.2-3a1 1 0 0 0-1.37.36L12 6.57L9.66 2.51a1 1 0 0 0-1.37-.36l-5.2 3a1 1 0 0 0-.47.61a1 1 0 0 0 .1.75l4 6.83A5.91 5.91 0 0 0 6 16a6 6 0 1 0 11.34-2.72l3.9-6.76a1 1 0 0 0 .14-.76M5 6.38l3.46-2L11.68 10A5.94 5.94 0 0 0 8 11.58ZM12 20a4 4 0 0 1-4-4a4 4 0 0 1 4-4a4 4 0 1 1 0 8m4-8.45a5.9 5.9 0 0 0-1.86-1.15l-.98-1.83l2.42-4.19l3.46 2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MedicalDrip(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 6h-2V4h2a1 1 0 0 0 0-2H9a5 5 0 0 0-5 5v14a1 1 0 0 0 2 0V7a3 3 0 0 1 3-3h4v2h-2a3 3 0 0 0-3 3v4.93a3 3 0 0 0 1.34 2.5L11 17.54V18a2 2 0 0 0 2 2v1a1 1 0 0 0 2 0v-1a2 2 0 0 0 2-2v-.46l1.66-1.11a3 3 0 0 0 1.34-2.5V9a3 3 0 0 0-3-3m-1 5h2v1h-1a1 1 0 0 0 0 2h1a1 1 0 0 1-.44.76l-2.1 1.41A1 1 0 0 0 15 17v1h-2v-1a1 1 0 0 0-.45-.83l-2.1-1.41a1 1 0 0 1-.45-.83V9a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1h-2a1 1 0 0 0 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MedicalSquare(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 2H5a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3m1 17a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1ZM17 9h-2V7a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1v2H7a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h2v2a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-2h2a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1m-1 4h-2a1 1 0 0 0-1 1v2h-2v-2a1 1 0 0 0-1-1H8v-2h2a1 1 0 0 0 1-1V8h2v2a1 1 0 0 0 1 1h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MedicalSquareFull(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1m-1 18H4V4h16ZM7 14.79h2v2a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-2h2a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1h-2v-2a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1v2H7a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1m1-4h2a1 1 0 0 0 1-1v-2h2v2a1 1 0 0 0 1 1h2v2h-2a1 1 0 0 0-1 1v2h-2v-2a1 1 0 0 0-1-1H8Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediumM(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.21 6.417H22V4.083h-7.48l-2.486 9.167h-.068L9.503 4.083H2v2.334h.768a.896.896 0 0 1 .732.694v9.83a.84.84 0 0 1-.732.642H2v2.334h6v-2.334H6.5V7.25h.088l3.457 12.667h2.712L16.26 7.25h.073v10.333h-1.5v2.334H22v-2.334h-.791a.84.84 0 0 1-.709-.641v-9.83a.898.898 0 0 1 .71-.695"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Medkit(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 17h1v1a1 1 0 0 0 2 0v-1h1a1 1 0 0 0 0-2h-1v-1a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2m9-11h-2V5a3 3 0 0 0-3-3h-4a3 3 0 0 0-3 3v1H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3M9 5a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v1H9Zm11 14a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-7h16Zm0-9H4V9a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MeetingBoard(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 10h2a1 1 0 0 0 0-2H7a1 1 0 0 0 0 2m14-6h-8V3a1 1 0 0 0-2 0v1H3a1 1 0 0 0-1 1v10a3 3 0 0 0 3 3h4.59l-2.3 2.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l2.29-2.3V21a1 1 0 0 0 2 0v-1.59l2.29 2.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L14.41 18H19a3 3 0 0 0 3-3V5a1 1 0 0 0-1-1m-1 11a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V6h16ZM7 14h6a1 1 0 0 0 0-2H7a1 1 0 0 0 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Megaphone(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.991 2.002a1 1 0 0 0-1 1v.637a9.036 9.036 0 0 1-7 3.363h-6a3.003 3.003 0 0 0-3 3v2a3.003 3.003 0 0 0 3 3h.484l-2.403 5.606a1 1 0 0 0 .92 1.394h4a.999.999 0 0 0 .918-.606l2.724-6.356a9.028 9.028 0 0 1 6.357 3.325v.637a1 1 0 0 0 2 0v-16a1 1 0 0 0-1-1m-14 11a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1h1v4Zm2.341 7H6.508l2.142-5h1.825Zm10.66-4.478a11.052 11.052 0 0 0-7-2.522h-3v-4h3a11.053 11.053 0 0 0 7-2.522Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Meh(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 11a1 1 0 1 0-1-1a1 1 0 0 0 1 1m6 3H9a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2m0-5a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-3-7a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MehAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 11h1a1 1 0 0 0 0-2H9a1 1 0 0 0 0 2m6 3H9a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2m0-5h-1a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2m-3-7a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MehClosedEye(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.21 10.54a1 1 0 0 0 1.41 0a1 1 0 0 0 0-1.41a3.08 3.08 0 0 0-4.24 0a1 1 0 1 0 1.41 1.41a1 1 0 0 1 1.42 0M15 14H9a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2m2.62-4.87a3.08 3.08 0 0 0-4.24 0a1 1 0 0 0 1.41 1.41a1 1 0 0 1 1.42 0a1 1 0 0 0 1.41 0a1 1 0 0 0 0-1.41M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Message(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.34 9.32l-14-7a3 3 0 0 0-4.08 3.9l2.4 5.37a1.06 1.06 0 0 1 0 .82l-2.4 5.37A3 3 0 0 0 5 22a3.14 3.14 0 0 0 1.35-.32l14-7a3 3 0 0 0 0-5.36Zm-.89 3.57l-14 7a1 1 0 0 1-1.35-1.3l2.39-5.37a2 2 0 0 0 .08-.22h6.89a1 1 0 0 0 0-2H6.57a2 2 0 0 0-.08-.22L4.1 5.41a1 1 0 0 1 1.35-1.3l14 7a1 1 0 0 1 0 1.78"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Metro(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.71 14.29a1.002 1.002 0 0 0-1.09-.21a.9.9 0 0 0-.54.54a1 1 0 1 0 1.84 0a1.147 1.147 0 0 0-.21-.33m8 0a1.047 1.047 0 0 0-1.42 0a1.147 1.147 0 0 0-.21.33a.99.99 0 0 0 .21 1.09a1.147 1.147 0 0 0 .33.21a.941.941 0 0 0 .76 0a1.16 1.16 0 0 0 .33-.21a.99.99 0 0 0 .21-1.09a1.147 1.147 0 0 0-.21-.33m2.6 4.605a4.97 4.97 0 0 0 1.784-4.817l-1.5-8A5 5 0 0 0 14.68 2H9.319a5 5 0 0 0-4.913 4.078l-1.5 8a4.97 4.97 0 0 0 1.785 4.817l-1.398 1.398a1 1 0 1 0 1.414 1.414l1.87-1.87A5.006 5.006 0 0 0 7.818 20h8.362a5.006 5.006 0 0 0 1.243-.162l1.869 1.869a1 1 0 0 0 1.414-1.414ZM6.37 6.447A3.002 3.002 0 0 1 9.32 4h5.362a3.002 3.002 0 0 1 2.948 2.447l.347 1.85a7.955 7.955 0 0 1-11.952 0Zm12.117 10.469A2.99 2.99 0 0 1 16.181 18H7.819a3 3 0 0 1-2.948-3.553l.711-3.792a9.954 9.954 0 0 0 12.836 0l.71 3.792a2.99 2.99 0 0 1-.64 2.469"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Microphone(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 15a4 4 0 0 0 4-4V5a4 4 0 0 0-8 0v6a4 4 0 0 0 4 4M10 5a2 2 0 0 1 4 0v6a2 2 0 0 1-4 0Zm10 6a1 1 0 0 0-2 0a6 6 0 0 1-12 0a1 1 0 0 0-2 0a8 8 0 0 0 7 7.93V21H9a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2h-2v-2.07A8 8 0 0 0 20 11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicrophoneSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5 3.73a2 2 0 0 1 2.95-.14A2 2 0 0 1 14 5v3.41a1 1 0 0 0 2 0V5a4 4 0 0 0-7-2.53a1 1 0 1 0 1.5 1.26m8.22 9.54h.2a1 1 0 0 0 1-.81A7.91 7.91 0 0 0 20 11a1 1 0 0 0-2 0a5.54 5.54 0 0 1-.11 1.1a1 1 0 0 0 .83 1.17m3 6.06l-18-18a1 1 0 0 0-1.42 0a1 1 0 0 0 0 1.41L8 8.48V11a4 4 0 0 0 6 3.46l1.46 1.46A6 6 0 0 1 6 11a1 1 0 0 0-2 0a8 8 0 0 0 7 7.93V21H9a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2h-2v-2.07a7.87 7.87 0 0 0 3.85-1.59l3.4 3.4a1 1 0 0 0 1.42-1.41ZM12 13a2 2 0 0 1-2-2v-.52l2.45 2.46A1.74 1.74 0 0 1 12 13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Microscope(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 21.005h-5.184a2.964 2.964 0 0 0 .143-.591A8.044 8.044 0 0 0 20 13.005a7.945 7.945 0 0 0-2.127-5.422l.637-.638a.991.991 0 0 0 .241-.39l.708-2.122a1 1 0 0 0-.241-1.024l-2.122-2.121a.999.999 0 0 0-1.024-.242l-2.12.707a.997.997 0 0 0-.391.242L7.198 8.358a1 1 0 0 0 0 1.414l-1.416 1.415a1 1 0 0 0 0 1.415l2.122 2.12a1 1 0 0 0 1.414 0l1.414-1.413l.002.002a1 1 0 0 0 1.414 0l4.31-4.312A5.955 5.955 0 0 1 18 13.005a6.048 6.048 0 0 1-3.455 5.431a2.976 2.976 0 0 0-5.124.063a6.822 6.822 0 0 1-1.12-.554a.989.989 0 0 0 .699-.94a1 1 0 0 0-1-1H4a1 1 0 0 0 0 2h1.331a8.814 8.814 0 0 0 3.717 2.473a2.955 2.955 0 0 0 .136.527H4a1 1 0 0 0 0 2h16a1 1 0 0 0 0-2M8.611 12.602l-.708-.707l.708-.708l.707.708Zm8.318-6.904L11.44 11.19L9.32 9.065l5.489-5.489l1.311-.437l1.247 1.247ZM12 21.005a1 1 0 1 1 1-1a1.001 1.001 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Microsoft(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 22h9.5v-9.5H2zm0-10.5h9.5V2H2zM12.5 2v9.5H22V2zm0 20H22v-9.5h-9.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Minus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 11H5a1 1 0 0 0 0 2h14a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusCircle(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8m4-9H8a1 1 0 0 0 0 2h8a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusPath(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.18 4h2.1a1 1 0 0 0 0-2h-2.1a1 1 0 0 0 0 2M3 11.28a1 1 0 0 0 1-1v-2.1a1 1 0 0 0-2 0v2.1a1 1 0 0 0 1 1M14.46 4a1 1 0 0 0 2 0V3a1 1 0 0 0-1-1h-1a1 1 0 0 0 0 2M21 7.54h-4.54a1 1 0 1 0-2 0H8.54a1 1 0 0 0-1 1v5.92a1 1 0 1 0 0 2V21a1 1 0 0 0 1 1H21a1 1 0 0 0 1-1V8.54a1 1 0 0 0-1-1M20 20H9.54V9.54H20ZM4 2H3a1 1 0 0 0-1 1v1a1 1 0 0 0 2 0a1 1 0 0 0 0-2m0 12.46a1 1 0 0 0-2 0v1a1 1 0 0 0 1 1h1a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusSquare(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 2H5a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3m1 17a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Zm-4-8H8a1 1 0 0 0 0 2h8a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusSquareFull(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 13h6a1 1 0 0 0 0-2H9a1 1 0 0 0 0 2M21 2H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1m-1 18H4V4h16Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MissedCall(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 7.49a1 1 0 0 0 1-1V5.9l2.88 2.88a3 3 0 0 0 4.24 0l4.59-4.59a1 1 0 0 0 0-1.41a1 1 0 0 0-1.42 0l-4.58 4.58a1 1 0 0 1-1.42 0L8.41 4.49H9a1 1 0 0 0 0-2H6a1 1 0 0 0-.92.61a1.09 1.09 0 0 0-.08.39v3a1 1 0 0 0 1 1m15.94 7.36a16.27 16.27 0 0 0-19.88 0a2.69 2.69 0 0 0-1 2a2.66 2.66 0 0 0 .78 2.07l1.76 1.8a2.68 2.68 0 0 0 3.46.28l.47-.32a8.13 8.13 0 0 1 1-.55a1.85 1.85 0 0 0 1-2.3l-.09-.24a10.49 10.49 0 0 1 5.22 0l-.09.24a1.85 1.85 0 0 0 1 2.3a8.13 8.13 0 0 1 1 .55l.47.32a2.58 2.58 0 0 0 1.54.5a2.72 2.72 0 0 0 1.92-.79l1.81-1.82a2.66 2.66 0 0 0 .69-2.06a2.69 2.69 0 0 0-1.06-1.98m-1.14 2.64L19 19.3a.68.68 0 0 1-.86.1c-.19-.14-.38-.27-.59-.4a11.65 11.65 0 0 0-1.09-.61l.4-1.09a1 1 0 0 0-.6-1.28a12.42 12.42 0 0 0-8.5 0a1 1 0 0 0-.6 1.28l.4 1.1a9.8 9.8 0 0 0-1.1.6l-.58.4a.66.66 0 0 1-.88-.1l-1.8-1.81A.67.67 0 0 1 3 17a.76.76 0 0 1 .28-.53a14.29 14.29 0 0 1 17.44 0A.76.76 0 0 1 21 17a.67.67 0 0 1-.2.49"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MobileAndroid(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12.71 16.29l-.15-.12a.76.76 0 0 0-.18-.09L12.2 16a1 1 0 0 0-.91.27a1.15 1.15 0 0 0-.21.33a1 1 0 0 0 1.3 1.31a1.46 1.46 0 0 0 .33-.22a1 1 0 0 0 .21-1.09a1 1 0 0 0-.21-.31M16 2H8a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3m1 17a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MobileAndroidAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 2H8a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3m1 17a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1v-1h10Zm0-3H7V5a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MobileVibrate(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.88 14.12L9.73 14l-.18-.1l-.18-.05a1 1 0 0 0-.9.27a.83.83 0 0 0-.22.33a.94.94 0 0 0 0 .76a1.07 1.07 0 0 0 .54.54a1 1 0 0 0 .38.08a1.09 1.09 0 0 0 .39-.08a.87.87 0 0 0 .32-.22a1 1 0 0 0 .22-.32a1 1 0 0 0 .07-.38a.84.84 0 0 0-.08-.38a.93.93 0 0 0-.21-.33M3.51 8.76a1 1 0 0 0 .71-.3l4.24-4.24a1 1 0 0 0 0-1.41a1 1 0 0 0-1.41 0L2.81 7.05a1 1 0 0 0 0 1.41a1 1 0 0 0 .7.3m17.68 6.78a1 1 0 0 0-1.41 0l-4.24 4.24a1 1 0 0 0 .7 1.71a1 1 0 0 0 .71-.3L21.19 17a1 1 0 0 0 0-1.46m.17-5.66a3 3 0 0 0-.87-2.12l-4.25-4.25a3.08 3.08 0 0 0-4.24 0L3.51 12a3 3 0 0 0 0 4.24l4.25 4.25a3 3 0 0 0 4.24 0L20.49 12a3 3 0 0 0 .87-2.12m-2.29.71l-8.48 8.48a1 1 0 0 1-1.42 0l-4.24-4.24a1 1 0 0 1 0-1.42l8.48-8.48a1 1 0 0 1 1.42 0l4.24 4.24a1 1 0 0 1 0 1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Modem(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.48 13.13a.65.65 0 0 0-.05-.2a.89.89 0 0 0-.08-.17a.86.86 0 0 0-.1-.16l-.16-.13l-.09-.09l-14.72-8.5a1 1 0 0 0-1 1.74l11.49 6.63H3.5a1 1 0 0 0-1 1v4a3 3 0 0 0 3 3h13a3 3 0 0 0 3-3v-4s-.02-.08-.02-.12m-2 4.12a1 1 0 0 1-1 1H5.5a1 1 0 0 1-1-1v-3h15Zm-3 0a1 1 0 1 0-1-1a1 1 0 0 0 1.02 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoneyBill(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 11a1 1 0 1 0 1 1a1 1 0 0 0-1-1m12 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1m2-6H4a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h16a3 3 0 0 0 3-3V8a3 3 0 0 0-3-3m1 11a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1Zm-9-7a3 3 0 1 0 3 3a3 3 0 0 0-3-3m0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoneyBillSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 11a1 1 0 1 0 1 1a1 1 0 0 0-1-1m5.86-1.55L4.71 2.29a1 1 0 0 0-1.42 1.42L4.59 5H4a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h14.59l2.7 2.71a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Zm-.74 2.09l1.34 1.34A1 1 0 0 1 12 13a1 1 0 0 1-1-1a1 1 0 0 1 .12-.46M4 17a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1h2.59l3.1 3.1A3 3 0 0 0 9 12a3 3 0 0 0 3 3a3 3 0 0 0 1.9-.69L16.59 17ZM20 5h-7.34a1 1 0 0 0 0 2H20a1 1 0 0 1 1 1v7.34a1 1 0 1 0 2 0V8a3 3 0 0 0-3-3m-1 7a1 1 0 1 0-1 1a1 1 0 0 0 1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoneyBillStack(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 1H4a3 3 0 0 0-3 3v16a3 3 0 0 0 3 3h16a3 3 0 0 0 3-3V4a3 3 0 0 0-3-3M8 21H4a1 1 0 0 1-1-1v-1.18A3 3 0 0 0 4 19h4Zm0-4H4a1 1 0 0 1-1-1v-1.18A3 3 0 0 0 4 15h4Zm0-4H4a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h4Zm6 8h-4v-6h4Zm0-8h-4V3h4Zm7 7a1 1 0 0 1-1 1h-4v-2h4a3 3 0 0 0 1-.18Zm0-4a1 1 0 0 1-1 1h-4v-2h4a3 3 0 0 0 1-.18Zm0-4a1 1 0 0 1-1 1h-4V3h4a1 1 0 0 1 1 1Zm-3-5a1 1 0 1 0 1 1a1 1 0 0 0-1-1M6 9a1 1 0 1 0-1-1a1 1 0 0 0 1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoneyInsert(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.46 6l.54-.59V9a1 1 0 0 0 2 0V5.41l.54.55A1 1 0 0 0 15 6a1 1 0 0 0 0-1.42l-2.29-2.29a1 1 0 0 0-.33-.21a1 1 0 0 0-.76 0a1 1 0 0 0-.33.21L9 4.54A1 1 0 0 0 10.46 6M12 12a3 3 0 1 0 3 3a3 3 0 0 0-3-3m0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1m-7-1a1 1 0 1 0 1-1a1 1 0 0 0-1 1m14 0a1 1 0 1 0-1 1a1 1 0 0 0 1-1m1-7h-4a1 1 0 0 0 0 2h4a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1v-8a1 1 0 0 1 1-1h4a1 1 0 0 0 0-2H4a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h16a3 3 0 0 0 3-3v-8a3 3 0 0 0-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoneyStack(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 17H2a1 1 0 0 0 0 2h20a1 1 0 0 0 0-2m0 4H2a1 1 0 0 0 0 2h20a1 1 0 0 0 0-2M6 7a1 1 0 1 0 1 1a1 1 0 0 0-1-1m14-6H4a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h16a3 3 0 0 0 3-3V4a3 3 0 0 0-3-3m1 11a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1Zm-9-7a3 3 0 1 0 3 3a3 3 0 0 0-3-3m0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1m6-2a1 1 0 1 0 1 1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoneyWithdraw(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 12a3 3 0 1 0 3 3a3 3 0 0 0-3-3m0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1m-.71-6.29a1 1 0 0 0 .33.21a.94.94 0 0 0 .76 0a1 1 0 0 0 .33-.21L15 7.46A1 1 0 1 0 13.54 6l-.54.59V3a1 1 0 0 0-2 0v3.59L10.46 6A1 1 0 0 0 9 7.46ZM19 15a1 1 0 1 0-1 1a1 1 0 0 0 1-1m1-7h-3a1 1 0 0 0 0 2h3a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1v-8a1 1 0 0 1 1-1h3a1 1 0 0 0 0-2H4a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h16a3 3 0 0 0 3-3v-8a3 3 0 0 0-3-3M5 15a1 1 0 1 0 1-1a1 1 0 0 0-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoneyWithdrawal(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 2H2a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h3v9a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-9h3a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1M7 20v-2a2 2 0 0 1 2 2Zm10 0h-2a2 2 0 0 1 2-2Zm0-4a4 4 0 0 0-4 4h-2a4 4 0 0 0-4-4V8h10Zm4-6h-2V7a1 1 0 0 0-1-1H6a1 1 0 0 0-1 1v3H3V4h18Zm-9 5a3 3 0 1 0-3-3a3 3 0 0 0 3 3m0-4a1 1 0 1 1-1 1a1 1 0 0 1 1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Moneybag(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 11a3 3 0 1 0 3 3a3 3 0 0 0-3-3m0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1m7-9h-3V5a3 3 0 0 0-3-3h-2a3 3 0 0 0-3 3v1H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3m-9-1a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v1h-4ZM4 9a1 1 0 0 1 1-1h1a2 2 0 0 1-2 2Zm1 11a1 1 0 0 1-1-1v-1a2 2 0 0 1 2 2Zm15-1a1 1 0 0 1-1 1h-1a2 2 0 0 1 2-2Zm0-3a4 4 0 0 0-4 4H8a4 4 0 0 0-4-4v-4a4 4 0 0 0 4-4h8a4 4 0 0 0 4 4Zm0-6a2 2 0 0 1-2-2h1a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoneybagAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 7h-3V6a3 3 0 0 0-3-3h-2a3 3 0 0 0-3 3v1H5a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-8a3 3 0 0 0-3-3m-9-1a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v1h-4Zm-6 4a1 1 0 0 1 1-1h1a2 2 0 0 1-2 2Zm1 9a1 1 0 0 1-1-1v-1a2 2 0 0 1 2 2Zm15-1a1 1 0 0 1-1 1h-1a2 2 0 0 1 2-2Zm0-3a4 4 0 0 0-4 4H8a4 4 0 0 0-4-4v-2a4 4 0 0 0 4-4h8a4 4 0 0 0 4 4Zm0-4a2 2 0 0 1-2-2h1a1 1 0 0 1 1 1Zm-8 0a3 3 0 1 0 3 3a3 3 0 0 0-3-3m0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Monitor(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 2H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h2.64l-.58 1a2 2 0 0 0 0 2a2 2 0 0 0 1.75 1h6.46A2 2 0 0 0 17 21a2 2 0 0 0 0-2l-.59-1H19a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3M8.77 20L10 18h4l1.2 2ZM20 15a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-1h16Zm0-3H4V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MonitorHeartRate(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 19a1 1 0 0 0 .38-.08a1.15 1.15 0 0 0 .33-.21a1 1 0 0 0 .12-.16a.56.56 0 0 0 .09-.17a.64.64 0 0 0 .08-.18a1.36 1.36 0 0 0 0-.2a1 1 0 0 0-.08-.38a.9.9 0 0 0-.54-.54A1 1 0 0 0 8.8 17l-.18.06a.56.56 0 0 0-.17.09a1 1 0 0 0-.16.12a1 1 0 0 0-.21.33A1 1 0 0 0 8 18a1 1 0 0 0 1 1m-3.71-.29a1.15 1.15 0 0 0 .33.21A1 1 0 0 0 6 19h.19a.6.6 0 0 0 .19-.06a.76.76 0 0 0 .18-.09l.15-.12a1.15 1.15 0 0 0 .21-.33A.84.84 0 0 0 7 18a1.36 1.36 0 0 0 0-.2a.64.64 0 0 0-.06-.18a.56.56 0 0 0-.09-.17a1 1 0 0 0-.12-.16a1 1 0 0 0-1.09-.21a1 1 0 0 0-.33.21a1 1 0 0 0-.12.16a.56.56 0 0 0-.09.17a.64.64 0 0 0-.1.18a1.36 1.36 0 0 0 0 .2a1 1 0 0 0 .08.38a1.15 1.15 0 0 0 .21.33M19 2H5a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3m1 17a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-3h16Zm0-5H4v-4h4a1 1 0 0 0 .71-.29L10 8.46l2.8 3.2a1 1 0 0 0 .72.34a1 1 0 0 0 .71-.29L15.91 10H20Zm0-6h-4.5a1 1 0 0 0-.71.29l-1.24 1.25l-2.8-3.2a1 1 0 0 0-1.46 0L7.59 8H4V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Moon(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.64 13a1 1 0 0 0-1.05-.14a8.05 8.05 0 0 1-3.37.73a8.15 8.15 0 0 1-8.14-8.1a8.59 8.59 0 0 1 .25-2A1 1 0 0 0 8 2.36a10.14 10.14 0 1 0 14 11.69a1 1 0 0 0-.36-1.05m-9.5 6.69A8.14 8.14 0 0 1 7.08 5.22v.27a10.15 10.15 0 0 0 10.14 10.14a9.79 9.79 0 0 0 2.1-.22a8.11 8.11 0 0 1-7.18 4.32Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonEclipse(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a10 10 0 0 0-2.14.24h-.12a10 10 0 0 0-.1 19.44h.14A9.57 9.57 0 0 0 12 22a10 10 0 0 0 0-20m-2 17.74a8 8 0 0 1 0-15.48a8 8 0 0 1 0 15.48m4.53-.16a10 10 0 0 0 0-15.16a8 8 0 0 1 0 15.16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Moonset(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 19H8a1 1 0 0 0 0 2h3a1 1 0 0 0 0-2m9-4h-1.16A8.18 8.18 0 0 0 20 12.05a1 1 0 0 0-.34-.93a1 1 0 0 0-1-.19a6 6 0 0 1-1.92.32a6.06 6.06 0 0 1-6.06-6a6.93 6.93 0 0 1 .1-1a1 1 0 0 0-.35-.92a1 1 0 0 0-1-.18A8.06 8.06 0 0 0 4 10.68A8 8 0 0 0 5.27 15H4a1 1 0 0 0 0 2h16a1 1 0 0 0 0-2m-3.72 0H7.83a6 6 0 0 1 .88-9.36a8.06 8.06 0 0 0 8.05 7.61a7 7 0 0 0 .79 0A6.08 6.08 0 0 1 16.28 15M16 19h-1a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mountains(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.85 17.47l-5-8a1 1 0 0 0-1.7 0l-1 1.63l-3.29-5.6a1 1 0 0 0-1.72 0l-7 12A1 1 0 0 0 3 19h18a1 1 0 0 0 .85-1.53M10.45 17H4.74L10 8l2.93 5Zm2.35 0l2.2-3.43l1-1.68L19.2 17Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MountainsSun(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 10a4 4 0 1 0-4-4a4 4 0 0 0 4 4m0-6a2 2 0 1 1-2 2a2 2 0 0 1 2-2m-1.15 8.47a1 1 0 0 0-1.7 0l-1 1.63l-3.29-5.6a1 1 0 0 0-1.72 0l-7 12A1 1 0 0 0 3 22h18a1 1 0 0 0 .85-1.53ZM10.45 20H4.74L10 11l2.94 5l-1.25 2Zm2.35 0l1.49-2.37l.71-1.06l1-1.68L19.2 20Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mouse(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a7 7 0 0 0-7 7v6a7 7 0 0 0 14 0V9a7 7 0 0 0-7-7M7 9a5 5 0 0 1 4-4.9V10H7Zm10 6a5 5 0 0 1-10 0v-3h10Zm0-5h-4V4.1A5 5 0 0 1 17 9Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MouseAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 6a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0V7a1 1 0 0 0-1-1m0-4a7 7 0 0 0-7 7v6a7 7 0 0 0 14 0V9a7 7 0 0 0-7-7m5 13a5 5 0 0 1-10 0V9a5 5 0 0 1 10 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MouseAltTwo(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a7 7 0 0 0-7 7v6a7 7 0 0 0 14 0V9a7 7 0 0 0-7-7m5 13a5 5 0 0 1-10 0V9a5 5 0 0 1 4-4.9V12a1 1 0 0 0 2 0V4.1A5 5 0 0 1 17 9Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Multiply(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.41 12l6.3-6.29a1 1 0 1 0-1.42-1.42L12 10.59l-6.29-6.3a1 1 0 0 0-1.42 1.42l6.3 6.29l-6.3 6.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l6.29-6.3l6.29 6.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Music(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.65 2.24a1 1 0 0 0-.8-.23l-13 2A1 1 0 0 0 7 5v10.35A3.45 3.45 0 0 0 5.5 15A3.5 3.5 0 1 0 9 18.5v-7.64l11-1.69v4.18a3.45 3.45 0 0 0-1.5-.35a3.5 3.5 0 1 0 3.5 3.5V3a1 1 0 0 0-.35-.76M5.5 20A1.5 1.5 0 1 1 7 18.5A1.5 1.5 0 0 1 5.5 20m13-2a1.5 1.5 0 1 1 1.5-1.5a1.5 1.5 0 0 1-1.5 1.5M20 7.14L9 8.83v-3l11-1.66Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicNote(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.12 2.21a1 1 0 0 0-.86-.21l-8 2a1 1 0 0 0-.76 1v10.35A3.45 3.45 0 0 0 8 15a3.5 3.5 0 1 0 3.5 3.5v-7.72L18.74 9h.07l.19-.15l.15-.1a.93.93 0 0 0 .13-.15a.78.78 0 0 0 .1-.15a.55.55 0 0 0 .06-.18a.58.58 0 0 0 0-.19a.24.24 0 0 0 0-.08V3a1 1 0 0 0-.32-.79M8 20a1.5 1.5 0 1 1 1.5-1.5A1.5 1.5 0 0 1 8 20m9.5-12.78l-6 1.5V5.78l6-1.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicTuneSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 7.33a1 1 0 0 0 1-1v-.55l6-1.5v2.94L14.7 8.3a1 1 0 0 0 .24 2h.24L20.24 9h.07l.19-.09l.15-.1a.93.93 0 0 0 .13-.15a.78.78 0 0 0 .1-.15a.55.55 0 0 0 .06-.18a.65.65 0 0 0 0-.19A.24.24 0 0 0 21 8V3a1 1 0 0 0-1.24-1l-8 2A1 1 0 0 0 11 5v1.33a1 1 0 0 0 1 1m9.71 13l-9-9l-9-9a1 1 0 0 0-1.42 1.38l8.71 8.7v2.94A3.45 3.45 0 0 0 9.5 15a3.5 3.5 0 1 0 3.5 3.5v-4.09l7.29 7.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM9.5 20a1.5 1.5 0 1 1 1.5-1.5A1.5 1.5 0 0 1 9.5 20"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Na(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 6h-1a3 3 0 0 0-3 3v8a1 1 0 0 0 2 0v-4h3v4a1 1 0 0 0 2 0V9a3 3 0 0 0-3-3m1 5h-3V9a1 1 0 0 1 1-1h1a1 1 0 0 1 1 1ZM8 6a1 1 0 0 0-1 1v5.76L3.89 6.55A1 1 0 0 0 2 7v10a1 1 0 0 0 2 0v-5.76l3.11 6.21A1 1 0 0 0 8 18a.91.91 0 0 0 .23 0A1 1 0 0 0 9 17V7a1 1 0 0 0-1-1m4-2a1 1 0 0 0-1 1v14a1 1 0 0 0 2 0V5a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Navigator(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.17 9.23l-14-5.78a3 3 0 0 0-4 3.7L3.71 12l-1.58 4.85A3 3 0 0 0 2.94 20a3 3 0 0 0 2 .8a3 3 0 0 0 1.15-.23l14.05-5.78a3 3 0 0 0 0-5.54ZM5.36 18.7a1 1 0 0 1-1.06-.19a1 1 0 0 1-.27-1L5.49 13h13.73Zm.13-7.7L4 6.53a1 1 0 0 1 .27-1A1 1 0 0 1 5 5.22a1 1 0 0 1 .39.08L19.22 11Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Nerd(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.41 8.65v-.06a10 10 0 0 0-18.78-.06a.85.85 0 0 0-.08.24A9.87 9.87 0 0 0 2 12a10 10 0 1 0 19.41-3.35M12 4a8 8 0 0 1 6.92 4h-1.2a3 3 0 0 0-4.62.22A3.17 3.17 0 0 0 12 8a3.17 3.17 0 0 0-1.1.22A3 3 0 0 0 6.28 8h-1.2A8 8 0 0 1 12 4m4.5 6a1 1 0 1 1-1-1a1 1 0 0 1 1 1m-7 0a1 1 0 1 1-1-1a1 1 0 0 1 1 1M12 20a8 8 0 0 1-8-8a8.24 8.24 0 0 1 .26-2H5.5a3 3 0 0 0 6 .18a1 1 0 0 1 1 0a3 3 0 0 0 6-.18h1.24a8.24 8.24 0 0 1 .26 2a8 8 0 0 1-8 8m2.36-5.77a3.76 3.76 0 0 1-4.72 0a1 1 0 0 0-1.28 1.54a5.68 5.68 0 0 0 7.28 0a1 1 0 1 0-1.28-1.54"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Newspaper(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 11h-1a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2m0 4h-1a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2m-6-6h6a1 1 0 0 0 0-2h-6a1 1 0 0 0 0 2m10-6H7a1 1 0 0 0-1 1v3H3a1 1 0 0 0-1 1v10a3 3 0 0 0 3 3h13a4 4 0 0 0 4-4V4a1 1 0 0 0-1-1M6 18a1 1 0 0 1-2 0V9h2Zm14-1a2 2 0 0 1-2 2H7.82A3 3 0 0 0 8 18V5h12Zm-9-4h1a1 1 0 0 0 0-2h-1a1 1 0 0 0 0 2m0 4h1a1 1 0 0 0 0-2h-1a1 1 0 0 0 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ninja(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m19.44 5.34l-.06-.07a10 10 0 0 0-14.76 0l-.06.07A10 10 0 1 0 22 12a9.93 9.93 0 0 0-2.56-6.66M12 4a7.87 7.87 0 0 1 3.86 1H8.14A7.87 7.87 0 0 1 12 4M5.76 7h12.48a8 8 0 0 1 1.69 4H4.07a8 8 0 0 1 1.69-4M12 20a8 8 0 0 1-7.93-7h15.86A8 8 0 0 1 12 20M8 8a1 1 0 1 0 1 1a1 1 0 0 0-1-1m6 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-3 9.93a1 1 0 0 0 .49.13a1 1 0 0 0 .87-.51A3 3 0 0 1 15 16a1 1 0 0 0 0-2a5 5 0 0 0-4.37 2.57a1 1 0 0 0 .37 1.36"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoEntry(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 20h-1.24L15.37 4.2A3 3 0 0 0 12.48 2h-1a3 3 0 0 0-2.85 2.2L4.24 20H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2M10.56 4.73a1 1 0 0 1 1-.73h1a1 1 0 0 1 1 .73L14.35 8h-4.7ZM9.09 10h5.82L16 14H8ZM6.32 20l1.11-4h9.14l1.11 4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Notebooks(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 6a1 1 0 0 0-1 1v10a3 3 0 0 1-3 3H7a1 1 0 0 0 0 2h10a5 5 0 0 0 5-5V7a1 1 0 0 0-1-1m-3 9V5a3 3 0 0 0-3-3H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3M10 4h2v4.86l-.36-.3a1 1 0 0 0-1.28 0l-.36.3ZM4 15V5a1 1 0 0 1 1-1h3v7a1 1 0 0 0 1.65.76L11 10.63l1.35 1.13A1 1 0 0 0 13 12a1.06 1.06 0 0 0 .42-.09A1 1 0 0 0 14 11V4h1a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Notes(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 14H8a1 1 0 0 0 0 2h8a1 1 0 0 0 0-2m0-4h-6a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2m4-6h-3V3a1 1 0 0 0-2 0v1h-2V3a1 1 0 0 0-2 0v1H9V3a1 1 0 0 0-2 0v1H4a1 1 0 0 0-1 1v14a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V5a1 1 0 0 0-1-1m-1 15a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V6h2v1a1 1 0 0 0 2 0V6h2v1a1 1 0 0 0 2 0V6h2v1a1 1 0 0 0 2 0V6h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ObjectGroup(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 10h-2V8a1 1 0 0 0-1-1H8a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h2v2a1 1 0 0 0 1 1h5a1 1 0 0 0 1-1v-5a1 1 0 0 0-1-1m-6 1v1H9V9h3v1h-1a1 1 0 0 0-1 1m5 4h-3v-3h3Zm6 3.28V5.72A2 2 0 1 0 18.28 3H5.72A2 2 0 1 0 3 5.72v12.56A2 2 0 1 0 5.72 21h12.56A2 2 0 1 0 21 18.28m-2 0a1.91 1.91 0 0 0-.72.72H5.72a1.91 1.91 0 0 0-.72-.72V5.72A1.91 1.91 0 0 0 5.72 5h12.56a1.91 1.91 0 0 0 .72.72Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ObjectUngroup(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 18.28v-6.56A2 2 0 1 0 18.28 9H15V5.72A2 2 0 1 0 12.28 3H5.72A2 2 0 1 0 3 5.72v6.56A2 2 0 1 0 5.72 15H9v3.28A2 2 0 1 0 11.72 21h6.56A2 2 0 1 0 21 18.28M8 10a2 2 0 0 0 1 1.72V13H5.72a1.91 1.91 0 0 0-.72-.72V5.72A1.91 1.91 0 0 0 5.72 5h6.56a1.91 1.91 0 0 0 .72.72V9h-1.28A2 2 0 0 0 8 10m5 1v1.28a1.91 1.91 0 0 0-.72.72H11v-1.28a1.91 1.91 0 0 0 .72-.72Zm6 7.28a1.91 1.91 0 0 0-.72.72h-6.56a1.91 1.91 0 0 0-.72-.72V15h1.28A2 2 0 1 0 15 12.28V11h3.28a1.91 1.91 0 0 0 .72.72Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Octagon(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.71 7.57l-5.27-5.28a1.05 1.05 0 0 0-.71-.29H8.27a1.05 1.05 0 0 0-.71.29L2.29 7.57a1 1 0 0 0-.29.7v7.46a1 1 0 0 0 .29.7l5.27 5.28a1.05 1.05 0 0 0 .71.29h7.46a1.05 1.05 0 0 0 .71-.29l5.27-5.28a1 1 0 0 0 .29-.7V8.27a1 1 0 0 0-.29-.7M20 15.31L15.31 20H8.69L4 15.31V8.69L8.69 4h6.62L20 8.69Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Okta(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10s10-4.5 10-10S17.5 2 12 2m5 10c0 2.8-2.2 5-5 5s-5-2.2-5-5s2.2-5 5-5s5 2.2 5 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Opera(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.996 2c-5.462 0-9.278 3.958-9.278 9.899C2.718 17.189 6.43 22 12.004 22c5.567 0 9.278-4.819 9.278-10.101c0-5.94-3.824-9.899-9.286-9.899m0 18.384c-3.397 0-3.77-5.013-3.77-8.71V11.6c0-3.996.598-8.23 3.748-8.23s3.786 4.361 3.786 8.357c0 3.696-.367 8.657-3.764 8.657"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OperaAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.97 5c-3.643 0-3.643 5.123-3.643 6.858c0 1.953 0 7.142 3.66 7.142c3.655 0 3.655-5.162 3.655-7.105C15.642 7.32 14.406 5 11.97 5m.017 12c-1.101 0-1.66-1.73-1.66-5.195c0-2.191.285-4.805 1.644-4.805c1.454 0 1.67 3.067 1.67 4.895c0 3.388-.556 5.105-1.654 5.105m.001-16C5.937 1 1.71 5.482 1.71 11.9c0 5.456 3.847 11.1 10.285 11.1c6.434 0 10.278-5.644 10.278-11.101C22.273 5.482 18.044 1 11.988 1m.007 20C6.81 21 3.71 16.373 3.71 11.9c0-5.323 3.327-8.9 8.278-8.9c4.956 0 8.285 3.577 8.285 8.899c0 4.474-3.096 9.101-8.278 9.101"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OutgoingCall(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.44 13c-.22 0-.45-.07-.67-.12a9.44 9.44 0 0 1-1.31-.39a2 2 0 0 0-2.48 1l-.22.45a12.18 12.18 0 0 1-2.66-2a12.18 12.18 0 0 1-2-2.66l.42-.28a2 2 0 0 0 1-2.48a10.33 10.33 0 0 1-.39-1.31c-.05-.22-.09-.45-.12-.68a3 3 0 0 0-3-2.49h-3a3 3 0 0 0-3 3.41a19 19 0 0 0 16.52 16.46h.38a3 3 0 0 0 2-.76a3 3 0 0 0 1-2.25v-3a3 3 0 0 0-2.47-2.9m.5 6a1 1 0 0 1-.34.75a1.06 1.06 0 0 1-.82.25A17 17 0 0 1 4.07 5.22a1.09 1.09 0 0 1 .25-.82a1 1 0 0 1 .75-.34h3a1 1 0 0 1 1 .79q.06.41.15.81a11.12 11.12 0 0 0 .46 1.55l-1.4.65a1 1 0 0 0-.49 1.33a14.49 14.49 0 0 0 7 7a1 1 0 0 0 .76 0a1 1 0 0 0 .57-.52l.62-1.4a13.69 13.69 0 0 0 1.58.46q.4.09.81.15a1 1 0 0 1 .79 1Zm1.92-16.32a1 1 0 0 0-.54-.54a1 1 0 0 0-.38-.08h-4a1 1 0 1 0 0 2h1.58l-3.29 3.3a1 1 0 0 0 0 1.41a1 1 0 0 0 1.41 0l3.3-3.29v1.58a1 1 0 0 0 2 0v-4a1 1 0 0 0-.08-.38"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Package(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 2H5a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3m-9 2h4v3.13l-1.45-1a1 1 0 0 0-1.1 0l-1.45 1Zm10 15a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h3v5a1 1 0 0 0 .53.88a1 1 0 0 0 1-.05L12 8.2l2.45 1.63A1 1 0 0 0 16 9V4h3a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Padlock(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 13a1.49 1.49 0 0 0-1 2.61V17a1 1 0 0 0 2 0v-1.39A1.49 1.49 0 0 0 12 13m5-4V7A5 5 0 0 0 7 7v2a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-7a3 3 0 0 0-3-3M9 7a3 3 0 0 1 6 0v2H9Zm9 12a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1v-7a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pagelines(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.885 13.993c1.75-.901 2.282-3.35 2.282-3.35s-2.305-.99-4.055-.085a4.206 4.206 0 0 0-1.698 1.822a8.965 8.965 0 0 0 .06-.99a6.993 6.993 0 0 0 1.582-4.726C16.712 3.901 13.71 2 13.71 2s-2.442 2.583-2.095 5.35a6.819 6.819 0 0 0 2.518 4.03a9.322 9.322 0 0 1-.076 1.01a4.396 4.396 0 0 0-1.9-2.058c-1.774-.853-4.049.203-4.049.203s.603 2.432 2.376 3.284a4.72 4.72 0 0 0 3.258.076a9.433 9.433 0 0 1-1.458 2.9a4.393 4.393 0 0 0-2.012-1.98c-1.813-.763-4.028.404-4.028.404s.72 2.402 2.536 3.162a3.744 3.744 0 0 0 1.735.243a9.419 9.419 0 0 1-5.845 2.032a.672.672 0 0 0 0 1.344a10.786 10.786 0 0 0 7.968-3.527a4.954 4.954 0 0 0 3.336 1.194c1.96-.207 3.34-2.299 3.34-2.299s-1.792-1.753-3.75-1.543a3.54 3.54 0 0 0-1.36.456a10.744 10.744 0 0 0 .895-2.275a5.153 5.153 0 0 0 3.786-.013"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pagerduty(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 3c-1.6-.8-2.7-1-5.2-1H6v12.1h5.8c2.3 0 4-.1 5.5-1.1c1.6-1.1 2.6-3 2.5-5c.1-2.1-1-4-2.8-5m-4.6 8.6H8.9v-7h3.3c3 0 4.5 1 4.5 3.4c.1 2.6-1.8 3.6-4.3 3.6M6 22h2.9v-5.3H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaintTool(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 1h-8a3 3 0 0 0-3 3H6a3 3 0 0 0-3 3v3a3 3 0 0 0 3 3h6a1 1 0 0 1 1 1v1a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2v-1a3 3 0 0 0-3-3H6a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1h1a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V4a3 3 0 0 0-3-3m-3 16v4h-2v-4Zm4-11a1 1 0 0 1-1 1h-8a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Palette(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.42 15.54a1 1 0 0 0 0 1.41a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.41a1 1 0 0 0-1.42 0m0-8.49a1 1 0 0 0 0 1.41a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.41a1 1 0 0 0-1.42 0m4.95 10a1 1 0 1 0 1 1a1 1 0 0 0-1-1.05Zm-6-6a1 1 0 1 0 1 1a1 1 0 0 0-1-1.05Zm6-6a1 1 0 1 0 1 1a1 1 0 0 0-1-1.05Zm3.54 2.05a1 1 0 1 0 1.41 0a1 1 0 0 0-1.41-.05Zm6.3 0a11 11 0 1 0-7.85 15.74a3.87 3.87 0 0 0 2.5-1.65a4.2 4.2 0 0 0 .61-3.19a5.65 5.65 0 0 1-.1-1a5 5 0 0 1 3-4.56a3.84 3.84 0 0 0 2.06-2.25a4 4 0 0 0-.22-3.11Zm-1.7 2.44a1.9 1.9 0 0 1-1 1.09A7 7 0 0 0 15.37 17a7.3 7.3 0 0 0 .14 1.4a2.16 2.16 0 0 1-.31 1.65a1.79 1.79 0 0 1-1.21.8a8.72 8.72 0 0 1-1.62.15a9 9 0 0 1-9-9.28A9.05 9.05 0 0 1 11.85 3h.51a9 9 0 0 1 8.06 5a2 2 0 0 1 .09 1.52ZM12.37 11a1 1 0 1 0 1 1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PanelAdd(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 10h-4V3a1 1 0 0 0-1-1H8a1 1 0 0 0-1 1v5H3a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h15a1 1 0 0 0 1-1V11a1 1 0 0 0-1-1M7 20H4V10h3Zm5 0H9V4h3Zm5 0h-3v-8h3Zm4-16h-1V3a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2h1v1a1 1 0 0 0 2 0V6h1a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PanoramaH(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.54 5.16a1 1 0 0 0-1-.07A21.27 21.27 0 0 1 12 6.73a21.27 21.27 0 0 1-8.58-1.64a1 1 0 0 0-1 .07A1 1 0 0 0 2 6v12a1 1 0 0 0 .46.84a1 1 0 0 0 1 .07A21.27 21.27 0 0 1 12 17.27a21.27 21.27 0 0 1 8.58 1.64A1.06 1.06 0 0 0 21 19a1 1 0 0 0 .54-.16A1 1 0 0 0 22 18V6a1 1 0 0 0-.46-.84M20 16.52a24.77 24.77 0 0 0-8-1.25a24.77 24.77 0 0 0-8 1.25v-9a24.77 24.77 0 0 0 8 1.25a24.77 24.77 0 0 0 8-1.25Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PanoramaHalt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.46 5.83A1 1 0 0 0 20.7 5h-.11A37.49 37.49 0 0 0 3.41 5H3.3a1 1 0 0 0-.76.8a35.52 35.52 0 0 0 0 12.34a1 1 0 0 0 .76.8h.11A37.62 37.62 0 0 0 12 20a37.62 37.62 0 0 0 8.59-1h.11a1 1 0 0 0 .76-.8a35.52 35.52 0 0 0 0-12.37M19.6 17.17a35.42 35.42 0 0 1-15.2 0a33.2 33.2 0 0 1 0-10.34a35.42 35.42 0 0 1 15.2 0a33.2 33.2 0 0 1 0 10.34"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PanoramaV(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.27 12a21.11 21.11 0 0 1 1.64-8.58a1 1 0 0 0-.07-1A1 1 0 0 0 18 2H6a1 1 0 0 0-.84.46a1 1 0 0 0-.07 1A21.11 21.11 0 0 1 6.73 12a21.11 21.11 0 0 1-1.64 8.58a1 1 0 0 0 .07 1A1 1 0 0 0 6 22h12a1 1 0 0 0 .84-.46a1 1 0 0 0 .07-1A21.11 21.11 0 0 1 17.27 12m-.75 8h-9a24.77 24.77 0 0 0 1.25-8a24.77 24.77 0 0 0-1.29-8h9a24.77 24.77 0 0 0-1.25 8a24.77 24.77 0 0 0 1.29 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paperclip(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m18.08 12.42l-6.18 6.19a4.25 4.25 0 0 1-6-6l8-8a2.57 2.57 0 0 1 3.54 0a2.52 2.52 0 0 1 0 3.54l-6.9 6.89A.75.75 0 1 1 9.42 14l5.13-5.12a1 1 0 0 0-1.42-1.42L8 12.6a2.74 2.74 0 0 0 0 3.89a2.82 2.82 0 0 0 3.89 0l6.89-6.9a4.5 4.5 0 0 0-6.36-6.36l-8 8A6.25 6.25 0 0 0 13.31 20l6.19-6.18a1 1 0 1 0-1.42-1.42Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paragraph(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 13.5H3a1 1 0 0 0 0 2h10a1 1 0 0 0 0-2m8-5H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Parcel(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 14h2a1 1 0 0 0 0-2H7a1 1 0 0 0 0 2m6 2H7a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2m6-14H5a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3m-5 2v3.29l-1.51-.84a1 1 0 0 0-1 0L10 7.29V4Zm6 15a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h3v5a1 1 0 0 0 .5.86a1 1 0 0 0 1 0L12 8.47l2.51 1.4A1 1 0 0 0 15 10a1 1 0 0 0 1-1V4h3a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ParkingCircle(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 7h-3a1 1 0 0 0-1 1v8a1 1 0 0 0 2 0v-2h2a3 3 0 0 0 3-3v-1a3 3 0 0 0-3-3m1 4a1 1 0 0 1-1 1h-2V9h2a1 1 0 0 1 1 1Zm-2-9a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ParkingSquare(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 6H9a1 1 0 0 0-1 1v10a1 1 0 0 0 2 0v-3h2a4 4 0 0 0 0-8m0 6h-2V8h2a2 2 0 0 1 0 4m7-10H5a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3m1 17a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pathfinder(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 14.46a1 1 0 0 0-2 0v1a1 1 0 0 0 1 1h1a1 1 0 0 0 0-2M8.18 4h2.1a1 1 0 0 0 0-2h-2.1a1 1 0 0 0 0 2m6.28 0a1 1 0 0 0 2 0V3a1 1 0 0 0-1-1h-1a1 1 0 0 0 0 2M4 2H3a1 1 0 0 0-1 1v1a1 1 0 0 0 2 0a1 1 0 0 0 0-2m-1 9.28a1 1 0 0 0 1-1v-2.1a1 1 0 0 0-2 0v2.1a1 1 0 0 0 1 1M15.82 20h-2.1a1 1 0 1 0 0 2h2.1a1 1 0 0 0 0-2M21 7.54h-1a1 1 0 0 0 0 2a1 1 0 0 0 2 0v-1a1 1 0 0 0-1-1m0 5.18a1 1 0 0 0-1 1v2.1a1 1 0 0 0 2 0v-2.1a1 1 0 0 0-1-1m-4.54-5.18a1 1 0 1 0-2 0H8.54a1 1 0 0 0-1 1v5.92a1 1 0 1 0 0 2a1 1 0 0 0 2 0h5.92a1 1 0 0 0 1-1V9.54a1 1 0 1 0 0-2m-2 6.92H9.54V9.54h4.92ZM21 19a1 1 0 0 0-1 1a1 1 0 0 0 0 2h1a1 1 0 0 0 1-1v-1a1 1 0 0 0-1-1M9.54 20a1 1 0 0 0-2 0v1a1 1 0 0 0 1 1h1a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PathfinderUnite(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 7.54h-4.54V3a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v12.46a1 1 0 0 0 1 1h4.54V21a1 1 0 0 0 1 1H21a1 1 0 0 0 1-1V8.54a1 1 0 0 0-1-1M20 20H9.54v-4.54a1 1 0 0 0-1-1H4V4h10.46v4.54a1 1 0 0 0 1 1H20Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pause(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 2a3 3 0 0 0-3 3v14a3 3 0 0 0 6 0V5a3 3 0 0 0-3-3m1 17a1 1 0 0 1-2 0V5a1 1 0 0 1 2 0ZM8 2a3 3 0 0 0-3 3v14a3 3 0 0 0 6 0V5a3 3 0 0 0-3-3m1 17a1 1 0 0 1-2 0V5a1 1 0 0 1 2 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseCircle(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 7a1 1 0 0 0-1 1v8a1 1 0 0 0 2 0V8a1 1 0 0 0-1-1m2-5a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8m2-13a1 1 0 0 0-1 1v8a1 1 0 0 0 2 0V8a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paypal(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.437 7.104a3.823 3.823 0 0 0-.573-.523a4.725 4.725 0 0 0-1.157-3.74C17.623 1.619 15.775 1 13.214 1H7.001a1.892 1.892 0 0 0-1.864 1.592l-2.59 16.406a1.533 1.533 0 0 0 1.516 1.785h2.664l-.082.52A1.467 1.467 0 0 0 8.093 23h3.235a1.761 1.761 0 0 0 1.75-1.47l.641-4.031l.011-.055h.299c4.032 0 6.55-1.993 7.285-5.762a5.149 5.149 0 0 0-.877-4.578m-12.595 6.6l-.714 4.535l-.086.544H4.606L7.097 3h6.117c1.936 0 3.318.404 3.993 1.164a2.967 2.967 0 0 1 .607 2.733l-.018.113c-.012.076-.023.15-.044.246a5.846 5.846 0 0 1-2.005 3.67a6.677 6.677 0 0 1-4.217 1.183H9.707a1.88 1.88 0 0 0-1.865 1.595m11.51-2.405c-.552 2.828-2.243 4.145-5.323 4.145h-.484a1.761 1.761 0 0 0-1.75 1.473l-.65 4.074L8.717 21l.478-3.034l.612-3.853h1.719c.157 0 .295-.023.448-.029c.359-.012.717-.026 1.053-.068c.205-.025.393-.072.59-.108c.273-.05.545-.1.801-.171c.19-.053.368-.122.55-.186c.238-.085.474-.174.697-.279a7 7 0 0 0 .486-.257a6.771 6.771 0 0 0 .613-.392c.142-.102.282-.208.415-.32a6.564 6.564 0 0 0 .537-.52c.113-.12.228-.237.333-.367a7.09 7.09 0 0 0 .48-.693c.076-.122.161-.235.232-.363a8.332 8.332 0 0 0 .52-1.154c.009-.024.021-.044.03-.068c.004-.01.01-.02.014-.032a4.299 4.299 0 0 1 .026 2.193"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pen(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 7.24a1 1 0 0 0-.29-.71l-4.24-4.24a1 1 0 0 0-.71-.29a1 1 0 0 0-.71.29l-2.83 2.83L2.29 16.05a1 1 0 0 0-.29.71V21a1 1 0 0 0 1 1h4.24a1 1 0 0 0 .76-.29l10.87-10.93L21.71 8a1.19 1.19 0 0 0 .22-.33a1 1 0 0 0 0-.24a.7.7 0 0 0 0-.14ZM6.83 20H4v-2.83l9.93-9.93l2.83 2.83ZM18.17 8.66l-2.83-2.83l1.42-1.41l2.82 2.82Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pentagon(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.59 9.17l-9-6.54a1 1 0 0 0-1.18 0l-9 6.54a1 1 0 0 0-.36 1.12l3.44 10.58a1 1 0 0 0 1 .69h11.07a1 1 0 0 0 1-.69L22 10.29a1 1 0 0 0-.41-1.12m-4.75 10.39H7.16l-3-9.2L12 4.68l7.82 5.68Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Percentage(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.758 10.758a3 3 0 1 0-3-3a3.003 3.003 0 0 0 3 3m0-4a1 1 0 1 1-1 1a1 1 0 0 1 1-1m8.484 6.484a3 3 0 1 0 3 3a3.003 3.003 0 0 0-3-3m0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1m3.465-12.949a1 1 0 0 0-1.414 0l-14 14a1 1 0 1 0 1.414 1.414l14-14a1 1 0 0 0 0-1.414"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Phone(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.44 13c-.22 0-.45-.07-.67-.12a9.44 9.44 0 0 1-1.31-.39a2 2 0 0 0-2.48 1l-.22.45a12.18 12.18 0 0 1-2.66-2a12.18 12.18 0 0 1-2-2.66l.42-.28a2 2 0 0 0 1-2.48a10.33 10.33 0 0 1-.39-1.31c-.05-.22-.09-.45-.12-.68a3 3 0 0 0-3-2.49h-3a3 3 0 0 0-3 3.41a19 19 0 0 0 16.52 16.46h.38a3 3 0 0 0 2-.76a3 3 0 0 0 1-2.25v-3a3 3 0 0 0-2.47-2.9m.5 6a1 1 0 0 1-.34.75a1.05 1.05 0 0 1-.82.25A17 17 0 0 1 4.07 5.22a1.09 1.09 0 0 1 .25-.82a1 1 0 0 1 .75-.34h3a1 1 0 0 1 1 .79q.06.41.15.81a11.12 11.12 0 0 0 .46 1.55l-1.4.65a1 1 0 0 0-.49 1.33a14.49 14.49 0 0 0 7 7a1 1 0 0 0 .76 0a1 1 0 0 0 .57-.52l.62-1.4a13.69 13.69 0 0 0 1.58.46q.4.09.81.15a1 1 0 0 1 .79 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.44 13c-.22 0-.45-.07-.67-.12a9.44 9.44 0 0 1-1.31-.39a2 2 0 0 0-2.48 1l-.22.45a12.18 12.18 0 0 1-2.66-2a12.18 12.18 0 0 1-2-2.66l.42-.28a2 2 0 0 0 1-2.48a10.33 10.33 0 0 1-.39-1.31c-.05-.22-.09-.45-.12-.68a3 3 0 0 0-3-2.49h-3a3 3 0 0 0-3 3.41a19 19 0 0 0 16.52 16.46h.38a3 3 0 0 0 2-.76a3 3 0 0 0 1-2.25v-3a3 3 0 0 0-2.47-2.9m.5 6a1 1 0 0 1-.34.75a1.06 1.06 0 0 1-.82.25A17 17 0 0 1 4.07 5.22a1.09 1.09 0 0 1 .25-.82a1 1 0 0 1 .75-.34h3a1 1 0 0 1 1 .79q.06.41.15.81a11.12 11.12 0 0 0 .46 1.55l-1.4.65a1 1 0 0 0-.49 1.33a14.49 14.49 0 0 0 7 7a1 1 0 0 0 .76 0a1 1 0 0 0 .57-.52l.62-1.4a13.69 13.69 0 0 0 1.58.46q.4.09.81.15a1 1 0 0 1 .79 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhonePause(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.44 13c-.22 0-.45-.07-.67-.12a9.44 9.44 0 0 1-1.31-.39a2 2 0 0 0-2.48 1l-.22.45a12.18 12.18 0 0 1-2.66-2a12.18 12.18 0 0 1-2-2.66l.42-.28a2 2 0 0 0 1-2.48a10.33 10.33 0 0 1-.39-1.31c-.05-.22-.09-.45-.12-.68a3 3 0 0 0-3-2.49h-3a3 3 0 0 0-3 3.41a19 19 0 0 0 16.52 16.46h.38a3 3 0 0 0 2-.76a3 3 0 0 0 1-2.25v-3a3 3 0 0 0-2.47-2.9m.5 6a1 1 0 0 1-.34.75a1 1 0 0 1-.82.25A17 17 0 0 1 4.07 5.22a1.09 1.09 0 0 1 .25-.82a1 1 0 0 1 .75-.34h3a1 1 0 0 1 1 .79q.06.41.15.81a11.12 11.12 0 0 0 .46 1.55l-1.4.65a1 1 0 0 0-.49 1.33a14.49 14.49 0 0 0 7 7a1 1 0 0 0 .76 0a1 1 0 0 0 .57-.52l.62-1.4a13.69 13.69 0 0 0 1.58.46q.4.09.81.15a1 1 0 0 1 .79 1ZM19 10a1 1 0 0 0 1-1V5a1 1 0 0 0-2 0v4a1 1 0 0 0 1 1m-4 0a1 1 0 0 0 1-1V5a1 1 0 0 0-2 0v4a1 1 0 0 0 1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.91 13.1a1 1 0 0 0 .85.47A1 1 0 0 0 6.61 12a17 17 0 0 1-2.47-6.85a1 1 0 0 1 .24-.81A1 1 0 0 1 5.13 4h3a1 1 0 0 1 1 .8c0 .23.08.44.13.67v.13a10.33 10.33 0 0 0 .47 1.54l-1.39.66a1 1 0 0 0-.52.57a1 1 0 0 0 0 .77c.1.21.2.42.32.64a1 1 0 0 0 1.37.37a1 1 0 0 0 .5-.94l.57-.21a2 2 0 0 0 1.05-2.48a9.3 9.3 0 0 1-.39-1.3v-.1c0-.2-.08-.4-.11-.58A3 3 0 0 0 8.16 2h-3a3 3 0 0 0-2.28 1a3 3 0 0 0-.72 2.39a19.05 19.05 0 0 0 2.75 7.71m14.61-.21l-.6-.11h-.08a9.31 9.31 0 0 1-1.33-.39a2 2 0 0 0-2.47 1l-.21.46a12.39 12.39 0 0 1-1.92-1.37l8.8-8.79a1 1 0 1 0-1.42-1.42l-18 18a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l4.59-4.6a19.09 19.09 0 0 0 10.29 4.73a2.69 2.69 0 0 0 .4 0a3 3 0 0 0 2-.75a3 3 0 0 0 1-2.25v-3a3 3 0 0 0-2.47-2.93m.48 6a1 1 0 0 1-.34.75a1 1 0 0 1-.81.24a17.07 17.07 0 0 1-9.14-4.18l1.77-1.77a14.69 14.69 0 0 0 3.38 2.21a1 1 0 0 0 .77 0a1 1 0 0 0 .57-.52l.62-1.41a12 12 0 0 0 1.6.47h.11l.69.13a1 1 0 0 1 .78 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneTimes(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m19.85 5.56l1.79-1.79a1 1 0 1 0-1.41-1.41l-1.79 1.79l-1.8-1.79a1 1 0 0 0-1.41 1.41L17 5.56l-1.79 1.8a1 1 0 0 0 0 1.41a1 1 0 0 0 1.41 0L18.44 7l1.79 1.79a1 1 0 0 0 1.41 0a1 1 0 0 0 0-1.41ZM19.44 13c-.22 0-.45-.07-.67-.12a9.44 9.44 0 0 1-1.31-.39a2 2 0 0 0-2.48 1l-.22.45a12.18 12.18 0 0 1-2.66-2a12.18 12.18 0 0 1-2-2.66l.42-.28a2 2 0 0 0 1-2.48a10.33 10.33 0 0 1-.39-1.31c-.05-.22-.09-.45-.12-.68a3 3 0 0 0-3-2.49h-3a3 3 0 0 0-3 3.41a19 19 0 0 0 16.52 16.46h.38a3 3 0 0 0 2-.76a3 3 0 0 0 1-2.25v-3a3 3 0 0 0-2.47-2.9m.5 6a1 1 0 0 1-.34.75a1.06 1.06 0 0 1-.82.25A17 17 0 0 1 4.07 5.22a1.09 1.09 0 0 1 .25-.82a1 1 0 0 1 .75-.34h3a1 1 0 0 1 1 .79q.06.41.15.81a11.12 11.12 0 0 0 .46 1.55l-1.4.65a1 1 0 0 0-.49 1.33a14.49 14.49 0 0 0 7 7a1 1 0 0 0 .76 0a1 1 0 0 0 .57-.52l.62-1.4a13.69 13.69 0 0 0 1.58.46q.4.09.81.15a1 1 0 0 1 .79 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneVolume(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.41 13c-.22 0-.45-.07-.67-.12a9.86 9.86 0 0 1-1.31-.39a2 2 0 0 0-2.48 1l-.22.46a13.17 13.17 0 0 1-2.67-2a13.17 13.17 0 0 1-2-2.67l.46-.21a2 2 0 0 0 1-2.48a10.47 10.47 0 0 1-.39-1.32c-.05-.22-.09-.45-.12-.67a3 3 0 0 0-3-2.49H5a3 3 0 0 0-2.24 1a3 3 0 0 0-.73 2.4a19.07 19.07 0 0 0 5.41 11a19.07 19.07 0 0 0 11 5.41a2.56 2.56 0 0 0 .39 0a3 3 0 0 0 2-.76a3 3 0 0 0 1-2.24v-3A3 3 0 0 0 19.41 13m.49 6a1 1 0 0 1-.33.74a1 1 0 0 1-.82.25a17.16 17.16 0 0 1-9.87-4.84A17.16 17.16 0 0 1 4 5.25a1 1 0 0 1 .25-.82A1 1 0 0 1 5 4.1h3a1 1 0 0 1 1 .78c0 .27.09.55.15.82a11 11 0 0 0 .46 1.54l-1.4.66a1 1 0 0 0-.52.56a1 1 0 0 0 0 .76a14.49 14.49 0 0 0 7 7a1 1 0 0 0 .76 0a1 1 0 0 0 .56-.52l.63-1.4a12.41 12.41 0 0 0 1.58.46c.26.06.54.11.81.15a1 1 0 0 1 .78 1ZM14 2h-.7a1 1 0 0 0 .17 2H14a6 6 0 0 1 6 6v.53a1 1 0 0 0 .91 1.08h.08a1 1 0 0 0 1-.91V10A8 8 0 0 0 14 2m2 8a1 1 0 0 0 2 0a4 4 0 0 0-4-4a1 1 0 0 0 0 2a2 2 0 0 1 2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Picture(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 4a1 1 0 0 0 .71-1.7a1 1 0 0 0-1.42 0a1 1 0 0 0-.21.32A.84.84 0 0 0 19 3a1 1 0 0 0 1 1m0 9a1 1 0 0 0-1 1v.39l-1.48-1.48a2.77 2.77 0 0 0-3.93 0l-.7.7l-2.48-2.49a2.86 2.86 0 0 0-3.93 0L5 12.6V7a1 1 0 0 1 1-1h10a1 1 0 0 0 0-2H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-5a1 1 0 0 0-1-1M6 20a1 1 0 0 1-1-1v-3.57l2.9-2.9a.79.79 0 0 1 1.09 0l3.17 3.17l4.29 4.3Zm13-1a1 1 0 0 1-.18.53L14.31 15l.7-.7a.78.78 0 0 1 1.1 0L19 17.22Zm1-14a1 1 0 0 0-1 1v4a1 1 0 0 0 2 0V6a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PizzaSlice(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.51 12.48a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0-4.48a1 1 0 1 0 1 1a1 1 0 0 0-1-1M12 10a1 1 0 0 0-1 1v1a1 1 0 0 0 2 0v-1a1 1 0 0 0-1-1m8.5-1.43a3 3 0 0 0-2.3-.29a2.9 2.9 0 0 0-1.09.56L5.51 2.13a1 1 0 0 0-1 0A1 1 0 0 0 4 3v13.17A2.94 2.94 0 0 0 2 19a3 3 0 0 0 2.92 3h.58a18.57 18.57 0 0 0 16.11-9.41a3 3 0 0 0-1.1-4.02ZM6 4.73l9.89 5.71A12.57 12.57 0 0 1 6 16Zm13.87 6.88A16.58 16.58 0 0 1 5 20a1 1 0 0 1-1-1a1 1 0 0 1 .3-.72A1 1 0 0 1 5 18h.51a14.5 14.5 0 0 0 12.62-7.37a.9.9 0 0 1 .58-.44a1 1 0 0 1 .75.09a1 1 0 0 1 .42 1.33Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plane(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.75 12a1 1 0 0 0-.55-.89l-6.12-3.06v-4a3.08 3.08 0 1 0-6.16 0v4L2.8 11.11a1 1 0 0 0-.55.89v3.33a1 1 0 0 0 .43.83a1 1 0 0 0 .92.11l5.32-2V18l-1.82.6a1 1 0 0 0-.68.95V22a1 1 0 0 0 .3.71a1 1 0 0 0 .7.29h9.17a1 1 0 0 0 1-1v-2.5a1 1 0 0 0-.68-.95L15.08 18v-3.72l5.32 2a1 1 0 0 0 .92-.11a1 1 0 0 0 .43-.83Zm-7.31-.1a1 1 0 0 0-.93.11a1 1 0 0 0-.43.82v5.84a1 1 0 0 0 .69.95l1.81.6V21H8.41v-.78l1.81-.6a1 1 0 0 0 .69-.95v-5.84a1 1 0 0 0-.43-.82a1 1 0 0 0-.93-.11l-5.31 2v-1.28l6.11-3.06a1 1 0 0 0 .56-.89V4.08a1.08 1.08 0 1 1 2.16 0v4.59a1 1 0 0 0 .56.89l6.11 3.06v1.27Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlaneArrival(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.12 16.23a5 5 0 0 0-2.3-3L16.71 12l-.48-5.47a1 1 0 0 0-.49-.78l-3-1.72a1 1 0 0 0-1 0a1 1 0 0 0-.52.84l-.15 3.9l-1.75-1l-2.86-3.85a1 1 0 0 0-1.78.41l-.87 4.61a3 3 0 0 0 1.39 3.29l14.06 8.11a1 1 0 0 0 1.36-.34a4.91 4.91 0 0 0 .5-3.77M19.24 18L6.2 10.5a1 1 0 0 1-.44-1.13l.46-2.44l1.66 2.2a1 1 0 0 0 .3.27l3.35 1.94a1 1 0 0 0 1.5-.83l.16-3.9l1.09.63l.48 5.47a1 1 0 0 0 .5.78L17.82 15a2.91 2.91 0 0 1 1.36 1.78a2.74 2.74 0 0 1 .06 1.22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlaneDeparture(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 5.08a3.08 3.08 0 0 0-5.26-2.18l-2.81 2.81l-6.49-2.16a1 1 0 0 0-1 .24L4.06 6.15a1 1 0 0 0 .29 1.61l5.18 2.35l-2.6 2.6l-1.71-.86a1 1 0 0 0-1.16.15l-1.77 1.81a1 1 0 0 0 0 1.41l6.49 6.49a1 1 0 0 0 1.41 0L12 19.94a1 1 0 0 0 .19-1.16l-.86-1.71l2.6-2.6l2.35 5.18a1 1 0 0 0 1.61.29l2.36-2.36a1 1 0 0 0 .24-1l-2.16-6.49l2.77-2.83a3.05 3.05 0 0 0 .9-2.18m-2.32.77l-3.24 3.24a1 1 0 0 0-.24 1l2.16 6.48l-.9.9l-2.35-5.17a1 1 0 0 0-.73-.57a1 1 0 0 0-.89.28l-4.12 4.16a1 1 0 0 0-.19 1.15L10 19l-.56.56l-5.03-5.04L5 14l1.71.86a1 1 0 0 0 1.15-.19L12 10.51a1 1 0 0 0-.29-1.62L6.5 6.54l.9-.9l6.48 2.16a1 1 0 0 0 1-.24l3.24-3.24a1.07 1.07 0 0 1 1.53 0a1 1 0 0 1 .32.76a1.06 1.06 0 0 1-.29.77"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlaneFly(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23 9.32a1.06 1.06 0 0 0-.1-.76a4.93 4.93 0 0 0-6.75-1.8L14 8L9 5.65a1 1 0 0 0-.92 0l-3 1.73a1 1 0 0 0-.5.84a1 1 0 0 0 .46.87l3.3 2.08l-1.74 1l-4.78.58a1 1 0 0 0-.53 1.75l3.54 3.06a3 3 0 0 0 3.55.44L22.5 9.93a1 1 0 0 0 .5-.61m-15.53 7a1 1 0 0 1-1.2-.18l-1.9-1.63l2.73-.33a1 1 0 0 0 .38-.13l3.36-1.93a1 1 0 0 0 .5-.85a1 1 0 0 0-.47-.86l-3.3-2.09l1.1-.63l5 2.32a1 1 0 0 0 .92 0l2.56-1.48a3 3 0 0 1 3.36.29Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Play(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.54 9L8.88 3.46a3.42 3.42 0 0 0-5.13 3v11.12A3.42 3.42 0 0 0 7.17 21a3.43 3.43 0 0 0 1.71-.46L18.54 15a3.42 3.42 0 0 0 0-5.92Zm-1 4.19l-9.66 5.62a1.44 1.44 0 0 1-1.42 0a1.42 1.42 0 0 1-.71-1.23V6.42a1.42 1.42 0 0 1 .71-1.23A1.51 1.51 0 0 1 7.17 5a1.54 1.54 0 0 1 .71.19l9.66 5.58a1.42 1.42 0 0 1 0 2.46Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayCircle(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16 10.27l-5-2.89a2 2 0 0 0-3 1.73v5.78a2 2 0 0 0 1 1.73a2 2 0 0 0 2 0l5-2.89a2 2 0 0 0 0-3.46M15 12l-5 2.89V9.11zM12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plug(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 6h-3V3a1 1 0 0 0-2 0v3h-4V3a1 1 0 0 0-2 0v3H5a1 1 0 0 0 0 2h1v5a1 1 0 0 0 .29.71L9 16.41V21a1 1 0 0 0 2 0v-4h2v4a1 1 0 0 0 2 0v-4.59l2.71-2.7A1 1 0 0 0 18 13V8h1a1 1 0 0 0 0-2m-3 6.59L13.59 15h-3.18L8 12.59V8h8ZM11 13h2a1 1 0 0 0 0-2h-2a1 1 0 0 0 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 11h-6V5a1 1 0 0 0-2 0v6H5a1 1 0 0 0 0 2h6v6a1 1 0 0 0 2 0v-6h6a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusCircle(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8m4-9h-3V8a1 1 0 0 0-2 0v3H8a1 1 0 0 0 0 2h3v3a1 1 0 0 0 2 0v-3h3a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusSquare(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 13h2v2a1 1 0 0 0 2 0v-2h2a1 1 0 0 0 0-2h-2V9a1 1 0 0 0-2 0v2H9a1 1 0 0 0 0 2M21 2H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1m-1 18H4V4h16Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Podium(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.79 7.13a1 1 0 0 0-.79-.38H8v-.5a2 2 0 0 1 1.46-1.92a1.5 1.5 0 0 0 1 .42h1a1.5 1.5 0 0 0 0-3h-1a1.49 1.49 0 0 0-1.17.57A4 4 0 0 0 6 6.25v.5H5a1 1 0 0 0-.79.38A1 1 0 0 0 4 8l.62 2.49a3 3 0 0 0 2.48 2.22l.78 7H7a1 1 0 0 0 0 2h10a1 1 0 0 0 0-2h-.88l.78-7a3 3 0 0 0 2.45-2.23L20 8a1 1 0 0 0-.21-.87M14.1 19.75H9.9l-.78-7h5.76ZM17.41 10a1 1 0 0 1-1 .76H7.56a1 1 0 0 1-1-.76l-.28-1.25h11.44Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Polygon(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.87 11.5l-4.5-7.79a1 1 0 0 0-.87-.5h-9a1 1 0 0 0-.87.5l-4.5 7.79a1 1 0 0 0 0 1l4.5 7.79a1 1 0 0 0 .87.5h9a1 1 0 0 0 .87-.5l4.5-7.79a1 1 0 0 0 0-1m-6 7.29H8.08L4.15 12l3.93-6.79h7.84L19.85 12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PostStamp(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.5 5.5a1 1 0 0 0 1-1v-3a1 1 0 0 0-1-1h-3a1 1 0 0 0-1 1a.5.5 0 0 1-1 0a1 1 0 0 0-1-1h-3a1 1 0 0 0-1 1a.5.5 0 0 1-1 0a1 1 0 0 0-1-1h-3a1 1 0 0 0-1 1a.5.5 0 0 1-1 0a1 1 0 0 0-1-1h-3a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1a.5.5 0 0 1 0 1a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1a.5.5 0 0 1 0 1a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1a.5.5 0 0 1 0 1a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1a.5.5 0 0 1 1 0a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1a.5.5 0 0 1 1 0a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1a.5.5 0 0 1 1 0a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1v-3a1 1 0 0 0-1-1a.5.5 0 0 1 0-1a1 1 0 0 0 1-1v-3a1 1 0 0 0-1-1a.5.5 0 0 1 0-1a1 1 0 0 0 1-1v-3a1 1 0 0 0-1-1a.5.5 0 0 1 0-1m-1-1.79a2.5 2.5 0 0 0 0 4.58v1.42a2.5 2.5 0 0 0 0 4.58v1.42a2.5 2.5 0 0 0 0 4.58v1.21h-1.21a2.5 2.5 0 0 0-4.58 0h-1.42a2.5 2.5 0 0 0-4.58 0H8.29a2.5 2.5 0 0 0-4.58 0H2.5v-1.21a2.5 2.5 0 0 0 0-4.58v-1.42a2.5 2.5 0 0 0 0-4.58V8.29a2.5 2.5 0 0 0 0-4.58V2.5h1.21a2.5 2.5 0 0 0 4.58 0h1.42a2.5 2.5 0 0 0 4.58 0h1.42a2.5 2.5 0 0 0 4.58 0h1.21ZM12 5a7 7 0 0 0 0 14a6.93 6.93 0 0 0 3.5-.94a1 1 0 1 0-1-1.73A5 5 0 1 1 17 12v.5a.83.83 0 0 1-.83.83a.84.84 0 0 1-.84-.83V9.67a1 1 0 0 0-1.78-.6a3.25 3.25 0 0 0-1.55-.4a3.33 3.33 0 0 0 0 6.66a3.28 3.28 0 0 0 2.17-.82a2.84 2.84 0 0 0 4.83-2V12a7 7 0 0 0-7-7m0 8.33A1.33 1.33 0 1 1 13.33 12A1.32 1.32 0 0 1 12 13.33"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Postcard(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 11h1a1 1 0 0 0 1-1V9a1 1 0 0 0-1-1h-1a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1M6 12h5a1 1 0 0 0 0-2H6a1 1 0 0 0 0 2m16-8H2a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h20a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1m-1 14H3V6h18ZM6 16h5a1 1 0 0 0 0-2H6a1 1 0 0 0 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pound(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 20H8a4.92 4.92 0 0 0 1-3v-3h6a1 1 0 0 0 0-2H9V8.89a4.89 4.89 0 0 1 9.13-2.44a1 1 0 0 0 1.37.36a1 1 0 0 0 .37-1.36A6.9 6.9 0 0 0 7 8.89V12H4a1 1 0 0 0 0 2h3v3a3 3 0 0 1-3 3a1 1 0 0 0 0 2h16a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PoundCircle(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 15h-5.18a3 3 0 0 0 .18-1v-1h2.5a1 1 0 0 0 0-2H11v-1a1.95 1.95 0 0 1 3.63-1a1 1 0 0 0 1.74-1A4 4 0 0 0 9 10v1H8a1 1 0 0 0 0 2h1v1a1 1 0 0 1-1 1a1 1 0 0 0 0 2h8a1 1 0 0 0 0-2M12 1a11 11 0 1 0 11 11A11 11 0 0 0 12 1m0 20a9 9 0 1 1 9-9a9 9 0 0 1-9 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Power(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.21 6.21l.79-.8V10a1 1 0 0 0 2 0V5.41l.79.8a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-2.5-2.5a1 1 0 0 0-.33-.21a1 1 0 0 0-.76 0a1 1 0 0 0-.33.21l-2.5 2.5a1 1 0 0 0 1.42 1.42M18 7.56A1 1 0 1 0 16.56 9a6.45 6.45 0 1 1-9.12 0A1 1 0 1 0 6 7.56a8.46 8.46 0 1 0 12 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PrescriptionBottle(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 2H5a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2v11a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V8a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2m-2 14h-6v-4h6Zm0-6h-7a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h7v1a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1V8h10ZM5 6V4h14v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Presentation(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 14h-1V4h1a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2h1v10H3a1 1 0 0 0 0 2h8v1.15l-4.55 3A1 1 0 0 0 7 22a.94.94 0 0 0 .55-.17L11 19.55V21a1 1 0 0 0 2 0v-1.45l3.45 2.28A.94.94 0 0 0 17 22a1 1 0 0 0 .55-1.83l-4.55-3V16h8a1 1 0 0 0 0-2m-3 0H6V4h12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PresentationCheck(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.29 11.71a1 1 0 0 0 1.42 0l4-4a1 1 0 1 0-1.42-1.42L11 9.59l-1.29-1.3a1 1 0 0 0-1.42 1.42ZM21 14h-1V4h1a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2h1v10H3a1 1 0 0 0 0 2h8v1.15l-4.55 3A1 1 0 0 0 7 22a.94.94 0 0 0 .55-.17L11 19.55V21a1 1 0 0 0 2 0v-1.45l3.45 2.28A.94.94 0 0 0 17 22a1 1 0 0 0 .55-1.83l-4.55-3V16h8a1 1 0 0 0 0-2m-3 0H6V4h12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PresentationEdit(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.41 13h2.42a1 1 0 0 0 .71-.29l3.58-3.58a1 1 0 0 0 0-1.41l-2.42-2.4a1 1 0 0 0-1.41 0L8.71 8.9a1 1 0 0 0-.3.7V12a1 1 0 0 0 1 1m1-3L13 7.44l1 1L11.42 11h-1ZM21 14h-1V4h1a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2h1v10H3a1 1 0 0 0 0 2h8v1.15l-4.55 3A1 1 0 0 0 7 22a.94.94 0 0 0 .55-.17L11 19.55V21a1 1 0 0 0 2 0v-1.45l3.45 2.28A.94.94 0 0 0 17 22a1 1 0 0 0 .55-1.83l-4.55-3V16h8a1 1 0 0 0 0-2m-3 0H6V4h12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PresentationLine(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 14h-1V3a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v11H3a1 1 0 0 0 0 2h8v1.15l-4.55 3A1 1 0 0 0 7 22a.94.94 0 0 0 .55-.17L11 19.55V21a1 1 0 0 0 2 0v-1.45l3.45 2.28A.94.94 0 0 0 17 22a1 1 0 0 0 .55-1.83l-4.55-3V16h8a1 1 0 0 0 0-2m-3 0H6V4h12Zm-9-2a1 1 0 0 0 .83-.45l1.33-2l1.13 1.14a1 1 0 0 0 .81.29a1 1 0 0 0 .73-.45l2-3a1 1 0 0 0-1.66-1.1l-1.33 2l-1.13-1.14A1 1 0 0 0 10.9 7a1 1 0 0 0-.73.45l-2 3a1 1 0 0 0 .28 1.38A.94.94 0 0 0 9 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PresentationLinesAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.17 10.55a1 1 0 0 0 .73.45a1 1 0 0 0 .81-.29l1.13-1.14l1.33 2A1 1 0 0 0 15 12a.94.94 0 0 0 .55-.17a1 1 0 0 0 .28-1.38l-2-3A1 1 0 0 0 13.1 7a1 1 0 0 0-.81.29l-1.13 1.14l-1.33-2a1 1 0 1 0-1.66 1.1ZM21 14h-1V4h1a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2h1v10H3a1 1 0 0 0 0 2h8v1.15l-4.55 3A1 1 0 0 0 7 22a.94.94 0 0 0 .55-.17L11 19.55V21a1 1 0 0 0 2 0v-1.45l3.45 2.28A.94.94 0 0 0 17 22a1 1 0 0 0 .55-1.83l-4.55-3V16h8a1 1 0 0 0 0-2m-3 0H6V4h12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PresentationMinus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 10h4a1 1 0 0 0 0-2h-4a1 1 0 0 0 0 2m11 4h-1V4h1a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2h1v10H3a1 1 0 0 0 0 2h8v1.15l-4.55 3A1 1 0 0 0 7 22a.94.94 0 0 0 .55-.17L11 19.55V21a1 1 0 0 0 2 0v-1.45l3.45 2.28A.94.94 0 0 0 17 22a1 1 0 0 0 .55-1.83l-4.55-3V16h8a1 1 0 0 0 0-2m-3 0H6V4h12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PresentationPlay(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 14h-1V4h1a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2h1v10H3a1 1 0 0 0 0 2h8v1.15l-4.55 3A1 1 0 0 0 7 22a.94.94 0 0 0 .55-.17L11 19.55V21a1 1 0 0 0 2 0v-1.45l3.45 2.28A.94.94 0 0 0 17 22a1 1 0 0 0 .55-1.83l-4.55-3V16h8a1 1 0 0 0 0-2m-3 0H6V4h12Zm-8.39-1.74a1.73 1.73 0 0 0 1.76 0l3-1.74a1.76 1.76 0 0 0 0-3l-3-1.74a1.73 1.73 0 0 0-1.76 0a1.71 1.71 0 0 0-.87 1.52v3.48a1.71 1.71 0 0 0 .87 1.48m1.13-4.58L13 9l-2.28 1.32Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PresentationPlus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 14h-1V4h1a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2h1v10H3a1 1 0 0 0 0 2h8v1.15l-4.55 3A1 1 0 0 0 7 22a.94.94 0 0 0 .55-.17L11 19.55V21a1 1 0 0 0 2 0v-1.45l3.45 2.28A.94.94 0 0 0 17 22a1 1 0 0 0 .55-1.83l-4.55-3V16h8a1 1 0 0 0 0-2m-3 0H6V4h12Zm-8-4h1v1a1 1 0 0 0 2 0v-1h1a1 1 0 0 0 0-2h-1V7a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PresentationTimes(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.29 11.71a1 1 0 0 0 1.42 0l1.29-1.3l1.29 1.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L13.41 9l1.3-1.29a1 1 0 1 0-1.42-1.42L12 7.59l-1.29-1.3a1 1 0 0 0-1.42 1.42L10.59 9l-1.3 1.29a1 1 0 0 0 0 1.42M21 14h-1V4h1a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2h1v10H3a1 1 0 0 0 0 2h8v1.15l-4.55 3A1 1 0 0 0 7 22a.94.94 0 0 0 .55-.17L11 19.55V21a1 1 0 0 0 2 0v-1.45l3.45 2.28A.94.94 0 0 0 17 22a1 1 0 0 0 .55-1.83l-4.55-3V16h8a1 1 0 0 0 0-2m-3 0H6V4h12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Previous(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.41 12l3.3-3.29a1 1 0 1 0-1.42-1.42l-4 4a1 1 0 0 0 0 1.42l4 4a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM8 7a1 1 0 0 0-1 1v8a1 1 0 0 0 2 0V8a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PricetagAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 6a1 1 0 1 0 1 1a1 1 0 0 0-1-1m14.71 5.78l-9.48-9.46A1 1 0 0 0 11.5 2h-6a1 1 0 0 0-.71.29l-2.5 2.49a1 1 0 0 0-.29.71v6a1.05 1.05 0 0 0 .29.71l9.49 9.5a1.05 1.05 0 0 0 .71.29a1 1 0 0 0 .71-.29l8.51-8.51a1 1 0 0 0 .29-.71a1.05 1.05 0 0 0-.29-.7m-9.22 7.81L4 11.09V5.9L5.9 4h5.18l8.5 8.49Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Print(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 10a1 1 0 1 0 1 1a1 1 0 0 0-1-1m12-4h-1V3a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1v3H5a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3h1v3a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1v-3h1a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3M8 4h8v2H8Zm8 16H8v-4h8Zm4-5a1 1 0 0 1-1 1h-1v-1a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1v1H5a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PrintSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 10a1 1 0 1 0 1 1a1 1 0 0 0-1-1M3.71 2.29a1 1 0 0 0-1.42 1.42L4.62 6A3 3 0 0 0 2 9v6a3 3 0 0 0 3 3h1v3a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1v-1.59l2.29 2.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM6 15v1H5a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1h1.59l6 6H7a1 1 0 0 0-1 1m10 5H8v-4h6.59L16 17.41Zm3-14h-1V3a1 1 0 0 0-1-1H8.66a1 1 0 0 0 0 2H16v2h-3.34a1 1 0 0 0 0 2H19a1 1 0 0 1 1 1v6a.37.37 0 0 1 0 .11a1 1 0 0 0 .78 1.18h.2a1 1 0 0 0 1-.8A2.84 2.84 0 0 0 22 15V9a3 3 0 0 0-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Process(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.992 14.502a1 1 0 0 0-1 1v1.782a7.972 7.972 0 0 1-2-5.282a7.29 7.29 0 0 1 .052-.88a1 1 0 1 0-1.985-.24a9.173 9.173 0 0 0-.067 1.12a9.964 9.964 0 0 0 2.417 6.5H2.992a1 1 0 1 0 0 2h4a.982.982 0 0 0 .794-.422c.011-.015.026-.027.037-.043c.007-.01.007-.022.013-.032a.966.966 0 0 0 .106-.258a.952.952 0 0 0 .032-.156c.003-.03.018-.057.018-.089v-4a1 1 0 0 0-1-1m1.5-8.5H6.709a7.974 7.974 0 0 1 5.283-2a7.075 7.075 0 0 1 .88.053a1 1 0 0 0 .24-1.987a9.227 9.227 0 0 0-1.12-.066a9.964 9.964 0 0 0-6.5 2.417V3.002a1 1 0 0 0-2 0v4a.954.954 0 0 0 .039.195a.969.969 0 0 0 .141.346l.012.017a.973.973 0 0 0 .244.246c.011.008.017.02.028.028c.014.01.03.013.045.021a.958.958 0 0 0 .18.084a.988.988 0 0 0 .261.053c.018 0 .032.01.05.01h4a1 1 0 0 0 0-2m11.96 10.804a.967.967 0 0 0-.141-.345l-.011-.017a.973.973 0 0 0-.245-.246c-.011-.008-.016-.02-.028-.028c-.01-.007-.023-.007-.034-.014a1.154 1.154 0 0 0-.41-.136c-.032-.003-.059-.018-.091-.018h-4a1 1 0 0 0 0 2h1.782a7.973 7.973 0 0 1-5.282 2a7.074 7.074 0 0 1-.88-.054a1 1 0 0 0-.24 1.987a9.365 9.365 0 0 0 1.12.067a9.964 9.964 0 0 0 6.5-2.417v1.417a1 1 0 0 0 2 0v-4a.953.953 0 0 0-.04-.195Zm.54-11.304a1 1 0 0 0 0-2h-4a.952.952 0 0 0-.192.039l-.007.001a.968.968 0 0 0-.34.14l-.02.013a.974.974 0 0 0-.245.244c-.008.01-.02.016-.028.027c-.007.01-.007.023-.014.034a1.146 1.146 0 0 0-.136.413c-.003.03-.018.057-.018.089v4a1 1 0 1 0 2 0V6.719a7.975 7.975 0 0 1 2 5.283a7.289 7.289 0 0 1-.053.88a1.001 1.001 0 0 0 .872 1.113a1.03 1.03 0 0 0 .122.007a1 1 0 0 0 .991-.88a9.174 9.174 0 0 0 .068-1.12a9.964 9.964 0 0 0-2.417-6.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Processor(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 9h-4a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1m-1 4h-2v-2h2Zm8 0a1 1 0 0 0 0-2h-2V9h2a1 1 0 0 0 0-2h-2.18A3 3 0 0 0 17 5.18V3a1 1 0 0 0-2 0v2h-2V3a1 1 0 0 0-2 0v2H9V3a1 1 0 0 0-2 0v2.18A3 3 0 0 0 5.18 7H3a1 1 0 0 0 0 2h2v2H3a1 1 0 0 0 0 2h2v2H3a1 1 0 0 0 0 2h2.18A3 3 0 0 0 7 18.82V21a1 1 0 0 0 2 0v-2h2v2a1 1 0 0 0 2 0v-2h2v2a1 1 0 0 0 2 0v-2.18A3 3 0 0 0 18.82 17H21a1 1 0 0 0 0-2h-2v-2Zm-4 3a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ProgrammingLanguage(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m22 3l-3 15.1l-9.1 3l-7.9-3l.8-4.1h3.4l-.4 1.7l4.8 1.8l5.5-1.8l.8-3.8H3.2l.7-3.4h13.6l.5-2.1H4.3L5 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pump(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.54 6.29L19 4.75l-1.41-1.41a1 1 0 0 0-1.42 1.42l1 1l-.83 2.49a3 3 0 0 0 .73 3.07l2.95 3V19a1 1 0 0 1-2 0v-2a3 3 0 0 0-3-3H14V5a3 3 0 0 0-3-3H5a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h6a3 3 0 0 0 3-3v-3h1a1 1 0 0 1 1 1v2a3 3 0 0 0 6 0V9.83a5 5 0 0 0-1.46-3.54M12 19a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-7h8Zm0-9H4V5a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1Zm8 1.42l-1.54-1.54a1 1 0 0 1-.24-1l.51-1.54l.39.4A3 3 0 0 1 20 9.83Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PuzzlePiece(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 22H5a3 3 0 0 1-3-3V9a3 3 0 0 1 3-3h1a4 4 0 0 1 7.3-2.18A3.86 3.86 0 0 1 14 6h3a1 1 0 0 1 1 1v3a4 4 0 0 1 2.18 7.3A3.86 3.86 0 0 1 18 18v3a1 1 0 0 1-1 1M5 8a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h11v-3.18a1 1 0 0 1 .42-.82a1 1 0 0 1 .91-.13a1.77 1.77 0 0 0 1.74-.23a2 2 0 0 0 .93-1.37a2 2 0 0 0-.48-1.59a1.89 1.89 0 0 0-2.17-.55a1 1 0 0 1-.91-.13a1 1 0 0 1-.42-.82V8h-3.2a1 1 0 0 1-1-1.33a1.77 1.77 0 0 0-.23-1.74a1.94 1.94 0 0 0-3-.43A2 2 0 0 0 8 6a1.89 1.89 0 0 0 .13.67A1 1 0 0 1 7.18 8Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QrcodeScan(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 21H4a1 1 0 0 1-1-1v-4a1 1 0 0 0-2 0v4a3 3 0 0 0 3 3h4a1 1 0 0 0 0-2m14-6a1 1 0 0 0-1 1v4a1 1 0 0 1-1 1h-4a1 1 0 0 0 0 2h4a3 3 0 0 0 3-3v-4a1 1 0 0 0-1-1M20 1h-4a1 1 0 0 0 0 2h4a1 1 0 0 1 1 1v4a1 1 0 0 0 2 0V4a3 3 0 0 0-3-3M2 9a1 1 0 0 0 1-1V4a1 1 0 0 1 1-1h4a1 1 0 0 0 0-2H4a3 3 0 0 0-3 3v4a1 1 0 0 0 1 1m8-4H6a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1M9 9H7V7h2Zm5 2h4a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1m1-4h2v2h-2Zm-5 6H6a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1m-1 4H7v-2h2Zm5-1a1 1 0 0 0 1-1a1 1 0 0 0 0-2h-1a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1m4-3a1 1 0 0 0-1 1v3a1 1 0 0 0 0 2h1a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1m-4 4a1 1 0 1 0 1 1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Question(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.333 9.5A3.5 3.5 0 0 0 8.8 7.75a1 1 0 0 0 1.733 1a1.503 1.503 0 0 1 1.3-.75a1.5 1.5 0 1 1 0 3h-.003a.95.95 0 0 0-.19.039a1.032 1.032 0 0 0-.198.04a.983.983 0 0 0-.155.105a1.008 1.008 0 0 0-.162.11a1.005 1.005 0 0 0-.117.174a.978.978 0 0 0-.097.144a1.023 1.023 0 0 0-.043.212a.948.948 0 0 0-.035.176v1l.002.011v.491a1 1 0 0 0 1 .998h.003a1 1 0 0 0 .998-1.002l-.002-.662A3.494 3.494 0 0 0 15.333 9.5m-4.203 6.79a1 1 0 0 0 .7 1.71a1.036 1.036 0 0 0 .71-.29a1.015 1.015 0 0 0 0-1.42a1.034 1.034 0 0 0-1.41 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionCircle(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.29 15.29a1.58 1.58 0 0 0-.12.15a.76.76 0 0 0-.09.18a.64.64 0 0 0-.06.18a1.36 1.36 0 0 0 0 .2a.84.84 0 0 0 .08.38a.9.9 0 0 0 .54.54a.94.94 0 0 0 .76 0a.9.9 0 0 0 .54-.54A1 1 0 0 0 13 16a1 1 0 0 0-.29-.71a1 1 0 0 0-1.42 0M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8m0-13a3 3 0 0 0-2.6 1.5a1 1 0 1 0 1.73 1A1 1 0 0 1 12 9a1 1 0 0 1 0 2a1 1 0 0 0-1 1v1a1 1 0 0 0 2 0v-.18A3 3 0 0 0 12 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rainbow(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 12a1 1 0 0 0 0 2a5 5 0 0 1 5 5a1 1 0 0 0 2 0a7 7 0 0 0-7-7m0-8a1 1 0 0 0 0 2a13 13 0 0 1 13 13a1 1 0 0 0 2 0A15 15 0 0 0 5 4m0 4a1 1 0 0 0 0 2a9 9 0 0 1 9 9a1 1 0 0 0 2 0A11 11 0 0 0 5 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Raindrops(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5 8c0-3.49-3.3-5.74-3.44-5.83a1 1 0 0 0-1.12 0C5.8 2.27 2.5 4.55 2.5 8a4 4 0 0 0 8 0m-4 2a2 2 0 0 1-2-2a5.44 5.44 0 0 1 2-3.72A5.39 5.39 0 0 1 8.5 8a2 2 0 0 1-2 2m11.56-7.83a1 1 0 0 0-1.12 0c-.14.1-3.44 2.38-3.44 5.83a4 4 0 0 0 8 0c0-3.49-3.3-5.74-3.44-5.83M17.5 10a2 2 0 0 1-2-2a5.44 5.44 0 0 1 2-3.72a5.39 5.39 0 0 1 2 3.72a2 2 0 0 1-2 2m-4.44 2.17a1 1 0 0 0-1.12 0c-.14.1-3.44 2.38-3.44 5.83a4 4 0 0 0 8 0c0-3.49-3.3-5.74-3.44-5.83M12.5 20a2 2 0 0 1-2-2a5.44 5.44 0 0 1 2-3.72a5.39 5.39 0 0 1 2 3.72a2 2 0 0 1-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RaindropsAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 7.75C9 5 6.42 3.24 6.31 3.17a1 1 0 0 0-1.12 0C5.08 3.25 2.5 5 2.5 7.75a3.25 3.25 0 0 0 6.5 0M5.75 9A1.25 1.25 0 0 1 4.5 7.75A3.66 3.66 0 0 1 5.75 5.3A3.61 3.61 0 0 1 7 7.75A1.25 1.25 0 0 1 5.75 9m6.06 1.17a1 1 0 0 0-1.12 0c-.17.12-4.19 2.9-4.19 7.08a4.75 4.75 0 0 0 9.5 0c0-4.25-4-6.97-4.19-7.08M11.25 20a2.75 2.75 0 0 1-2.75-2.75c0-2.31 1.81-4.17 2.76-5c.94.79 2.74 2.63 2.74 5A2.75 2.75 0 0 1 11.25 20m6.81-17.83a1 1 0 0 0-1.12 0c-.14.1-3.44 2.38-3.44 5.83a4 4 0 0 0 8 0c0-3.49-3.3-5.74-3.44-5.83M17.5 10a2 2 0 0 1-2-2a5.44 5.44 0 0 1 2-3.72a5.39 5.39 0 0 1 2 3.72a2 2 0 0 1-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func React(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.103 10.438a1.786 1.786 0 1 0 2.44.654a1.786 1.786 0 0 0-2.44-.654m8.005 1.938q-.176-.201-.371-.403q.136-.144.264-.287c1.605-1.804 2.283-3.614 1.655-4.701c-.602-1.043-2.393-1.354-4.636-.918q-.331.065-.659.146q-.063-.216-.133-.43C14.467 3.49 13.238 1.999 11.982 2c-1.204 0-2.368 1.397-3.111 3.558q-.11.32-.203.644q-.219-.054-.44-.1c-2.366-.485-4.271-.165-4.898.924c-.601 1.043.027 2.75 1.528 4.472q.224.255.46.5c-.186.19-.361.381-.525.571c-1.465 1.698-2.057 3.376-1.457 4.415c.62 1.074 2.498 1.425 4.785.975q.278-.055.553-.124q.1.351.221.697C9.635 20.649 10.792 22 11.992 22c1.24 0 2.482-1.453 3.235-3.659c.06-.174.115-.355.169-.541q.355.088.715.156c2.203.417 3.952.09 4.551-.95c.619-1.075-.02-2.877-1.554-4.63M4.07 7.452c.386-.67 1.943-.932 3.986-.512q.196.04.399.09a20.464 20.464 0 0 0-.422 2.678A20.887 20.887 0 0 0 5.93 11.4q-.219-.227-.427-.465C4.216 9.461 3.708 8.081 4.07 7.452m3.887 5.728c-.51-.387-.985-.783-1.416-1.181c.43-.396.905-.79 1.415-1.176q-.028.589-.027 1.179q0 .59.028 1.178m0 3.94a7.237 7.237 0 0 1-2.64.094a1.766 1.766 0 0 1-1.241-.657c-.365-.63.111-1.978 1.364-3.43q.236-.273.488-.532a20.49 20.49 0 0 0 2.107 1.7a20.802 20.802 0 0 0 .426 2.712q-.25.063-.505.114m7.1-8.039q-.503-.317-1.018-.613q-.508-.292-1.027-.563a18.7 18.7 0 0 1 1.739-.635a18.218 18.218 0 0 1 .306 1.811M9.68 5.835c.636-1.85 1.578-2.98 2.304-2.98c.773-.001 1.777 1.218 2.434 3.197q.064.194.12.39a20.478 20.478 0 0 0-2.526.97a20.061 20.061 0 0 0-2.52-.981q.087-.3.188-.596m-.4 1.424a18.307 18.307 0 0 1 1.73.642q-1.052.542-2.048 1.181c.08-.638.187-1.249.318-1.823m-.317 7.66q.497.319 1.009.613q.522.3 1.057.576a18.196 18.196 0 0 1-1.744.665a19.144 19.144 0 0 1-.322-1.853m5.456 3.146a7.236 7.236 0 0 1-1.238 2.333a1.766 1.766 0 0 1-1.188.748c-.729 0-1.658-1.085-2.29-2.896q-.112-.321-.206-.648a20.109 20.109 0 0 0 2.53-1.01a20.8 20.8 0 0 0 2.547.979q-.072.249-.155.494m.362-1.324a19.267 19.267 0 0 1-1.762-.646q.509-.267 1.025-.565q.53-.306 1.032-.627a18.152 18.152 0 0 1-.295 1.838m.447-4.743q0 .911-.057 1.82c-.493.334-1.013.66-1.554.972c-.54.311-1.073.597-1.597.856q-.827-.396-1.622-.854q-.79-.455-1.544-.969q-.07-.91-.07-1.822q0-.911.068-1.821a24.168 24.168 0 0 1 3.158-1.823q.816.397 1.603.851q.79.454 1.55.959q.065.914.065 1.831m.956-5.093c1.922-.373 3.37-.122 3.733.507c.387.67-.167 2.148-1.554 3.706q-.115.129-.238.259a20.061 20.061 0 0 0-2.144-1.688a20.04 20.04 0 0 0-.405-2.649q.31-.076.608-.135m-.13 3.885a18.164 18.164 0 0 1 1.462 1.188a18.12 18.12 0 0 1-1.457 1.208q.023-.594.023-1.188q0-.604-.028-1.208m3.869 5.789c-.364.631-1.768.894-3.653.538q-.324-.061-.664-.146a20.069 20.069 0 0 0 .387-2.682a19.94 19.94 0 0 0 2.137-1.715q.177.183.336.364a7.234 7.234 0 0 1 1.403 2.238a1.766 1.766 0 0 1 .054 1.403"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Receipt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 12H7a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2m-1-2h4a1 1 0 0 0 0-2H8a1 1 0 0 0 0 2m1 6H7a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2m12-4h-3V3a1 1 0 0 0-.5-.87a1 1 0 0 0-1 0l-3 1.72l-3-1.72a1 1 0 0 0-1 0l-3 1.72l-3-1.72a1 1 0 0 0-1 0A1 1 0 0 0 2 3v16a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-6a1 1 0 0 0-1-1M5 20a1 1 0 0 1-1-1V4.73l2 1.14a1.08 1.08 0 0 0 1 0l3-1.72l3 1.72a1.08 1.08 0 0 0 1 0l2-1.14V19a3 3 0 0 0 .18 1Zm15-1a1 1 0 0 1-2 0v-5h2Zm-6.44-2.83a.76.76 0 0 0-.18-.09a.6.6 0 0 0-.19-.06a1 1 0 0 0-.9.27A1.05 1.05 0 0 0 12 17a1 1 0 0 0 .07.38a1.19 1.19 0 0 0 .22.33a1.15 1.15 0 0 0 .33.21a.94.94 0 0 0 .76 0a1.15 1.15 0 0 0 .33-.21A1 1 0 0 0 14 17a1.05 1.05 0 0 0-.29-.71a1.58 1.58 0 0 0-.15-.12m.14-3.88a1 1 0 0 0-1.62.33A1 1 0 0 0 13 14a1 1 0 0 0 1-1a1 1 0 0 0-.08-.38a.91.91 0 0 0-.22-.33"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReceiptAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 8h6a1 1 0 0 0 0-2h-6a1 1 0 0 0 0 2m-2 4h8a1 1 0 0 0 0-2H8a1 1 0 0 0 0 2m0 4h8a1 1 0 0 0 0-2H8a1 1 0 0 0 0 2M20 2H4a1 1 0 0 0-1 1v18a1 1 0 0 0 1.6.8l2.07-1.55l2.06 1.55a1 1 0 0 0 1.2 0L12 20.25l2.07 1.55a1 1 0 0 0 1.2 0l2.06-1.55l2.07 1.55a1 1 0 0 0 1.05.09A1 1 0 0 0 21 21V3a1 1 0 0 0-1-1m-1 17l-1.07-.8a1 1 0 0 0-1.2 0l-2.06 1.55l-2.07-1.55a1 1 0 0 0-1.2 0l-2.07 1.55l-2.06-1.55a1 1 0 0 0-1.2 0L5 19V4h14Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RecordAudio(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8m0-14a6 6 0 1 0 6 6a6 6 0 0 0-6-6m0 10a4 4 0 1 1 4-4a4 4 0 0 1-4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RedditAlienAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.41 16.867a3.375 3.375 0 0 1-2.368.633a3.368 3.368 0 0 1-2.365-.632a1 1 0 0 0-1.416 1.414a5.11 5.11 0 0 0 3.781 1.218a5.12 5.12 0 0 0 3.782-1.217a1 1 0 1 0-1.414-1.416M9.2 15.002a1 1 0 1 0-1-1a1 1 0 0 0 1 1m6-2a1 1 0 1 0 1 1a1.001 1.001 0 0 0-1-1M23 11.78a3.772 3.772 0 0 0-6.794-2.264a16.505 16.505 0 0 0-3.05-.479l.856-5.705l2.087.71a2.997 2.997 0 0 0 5.994-.064v-.023a3.029 3.029 0 0 0-3-2.955a2.977 2.977 0 0 0-2.332 1.155l-3.239-1.101a.999.999 0 0 0-1.311.798l-1.077 7.175a16.664 16.664 0 0 0-3.34.489a3.768 3.768 0 0 0-6.225 4.234A4.862 4.862 0 0 0 1 16c0 3.925 4.832 7 11 7s11-3.075 11-7a4.862 4.862 0 0 0-.569-2.249a3.783 3.783 0 0 0 .569-1.97M19.093 3a1 1 0 1 1-1 1a1.017 1.017 0 0 1 1-1M4.78 10a1.76 1.76 0 0 1 .882.25a9.979 9.979 0 0 0-2.648 1.673C3.01 11.876 3 11.828 3 11.78A1.783 1.783 0 0 1 4.78 10M12 21c-4.879 0-9-2.29-9-5s4.121-5 9-5s9 2.29 9 5s-4.121 5-9 5m8.986-9.077a9.978 9.978 0 0 0-2.648-1.673a1.76 1.76 0 0 1 .882-.25A1.783 1.783 0 0 1 21 11.78c0 .048-.01.096-.014.143"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Redo(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 11a1 1 0 0 0-1 1a8.05 8.05 0 1 1-2.22-5.5h-2.4a1 1 0 0 0 0 2h4.53a1 1 0 0 0 1-1V3a1 1 0 0 0-2 0v1.77A10 10 0 1 0 22 12a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Refresh(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.91 15.51h-4.53a1 1 0 0 0 0 2h2.4A8 8 0 0 1 4 12a1 1 0 0 0-2 0a10 10 0 0 0 16.88 7.23V21a1 1 0 0 0 2 0v-4.5a1 1 0 0 0-.97-.99M15 12a3 3 0 1 0-3 3a3 3 0 0 0 3-3m-4 0a1 1 0 1 1 1 1a1 1 0 0 1-1-1m1-10a10 10 0 0 0-6.88 2.77V3a1 1 0 0 0-2 0v4.5a1 1 0 0 0 1 1h4.5a1 1 0 0 0 0-2h-2.4A8 8 0 0 1 20 12a1 1 0 0 0 2 0A10 10 0 0 0 12 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Registered(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8m.5-13h-3a1 1 0 0 0-1 1v8a1 1 0 0 0 2 0v-2h2a1 1 0 0 1 1 1v1a1 1 0 0 0 2 0v-1a3 3 0 0 0-.78-2a3 3 0 0 0 .78-2v-1a3 3 0 0 0-3-3m1 4a1 1 0 0 1-1 1h-2V9h2a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Repeat(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.5 17.5H4v-11h7.8l-.8.79a1 1 0 0 0 1.41 1.42l2.5-2.5a1 1 0 0 0 0-1.42l-2.5-2.5a1 1 0 0 0-1.41 0a1 1 0 0 0 0 1.42l.79.79H3a1 1 0 0 0-1 1v13a1 1 0 0 0 1 1h2.5a1 1 0 0 0 0-2M21 4.5h-2.5a1 1 0 0 0 0 2H20v11h-8.37l.79-.79a1 1 0 0 0 0-1.42a1 1 0 0 0-1.41 0l-2.5 2.5a1 1 0 0 0 0 1.42l2.5 2.5a1 1 0 0 0 1.41-1.42l-.79-.79H21a1 1 0 0 0 1-1v-13a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Restaurant(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.84 11.63a3 3 0 0 0 2.16-.88l2.83-2.83a1 1 0 0 0 0-1.41a1 1 0 0 0-1.42 0l-2.86 2.82a1 1 0 0 1-1.42 0l3.54-3.53a1 1 0 1 0-1.42-1.42l-3.53 3.54a1 1 0 0 1 0-1.41l2.83-2.83a1 1 0 1 0-1.42-1.42L13.3 5.09a3 3 0 0 0 0 4.24L12 10.62l-8.27-8.3l-.1-.06a.71.71 0 0 0-.17-.11l-.18-.07L3.16 2h-.27a.57.57 0 0 0-.18 0a.7.7 0 0 0-.17.09l-.16.1h-.07l-.06.1a1 1 0 0 0-.11.17a1.07 1.07 0 0 0-.07.19s0 .07 0 .11a11 11 0 0 0 3.11 9.34l2.64 2.63l-5.41 5.4a1 1 0 0 0 0 1.42a1 1 0 0 0 .71.29a1 1 0 0 0 .71-.29l6.07-5.98l2.83-2.83l2-2a3 3 0 0 0 2.11.89m-7.65 1.82l-2.63-2.64A9.06 9.06 0 0 1 4 5.4l6.61 6.6Zm6.24.57A1 1 0 0 0 14 15.44l6.3 6.3a1 1 0 0 0 .7.26a1 1 0 0 0 .71-.29a1 1 0 0 0 0-1.42Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightIndentAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.64 9.56a1 1 0 1 0-1.28 1.54l1.08.9l-1.08.9a1 1 0 0 0-.13 1.41a1 1 0 0 0 1.41.13l2-1.67a1 1 0 0 0 0-1.54ZM9 5a1 1 0 0 0-1 1v12a1 1 0 0 0 2 0V6a1 1 0 0 0-1-1m4 2h8a1 1 0 0 0 0-2h-8a1 1 0 0 0 0 2m8 10h-8a1 1 0 0 0 0 2h8a1 1 0 0 0 0-2m0-8h-8a1 1 0 0 0 0 2h8a1 1 0 0 0 0-2m0 4h-8a1 1 0 0 0 0 2h8a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightToLeftTextDirection(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.5 17H5.91l1.3-1.29a1 1 0 0 0-1.42-1.42l-3 3a1 1 0 0 0-.21.33a1 1 0 0 0 0 .76a1 1 0 0 0 .21.33l3 3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L5.91 19H20.5a1 1 0 0 0 0-2m-9-7v4a1 1 0 0 0 2 0V4h2v10a1 1 0 0 0 2 0V4h1a1 1 0 0 0 0-2h-7a4 4 0 0 0 0 8m0-6v4a2 2 0 0 1 0-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Robot(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 15a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-7-1a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0v-2a1 1 0 0 0-1-1m20 0a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0v-2a1 1 0 0 0-1-1m-5-7h-4V5.72A2 2 0 0 0 14 4a2 2 0 0 0-4 0a2 2 0 0 0 1 1.72V7H7a3 3 0 0 0-3 3v9a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-9a3 3 0 0 0-3-3m-3.28 2l-.5 2h-2.44l-.5-2ZM18 19a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1v-9a1 1 0 0 1 1-1h1.22L9 12.24a1 1 0 0 0 1 .76h4a1 1 0 0 0 1-.76L15.78 9H17a1 1 0 0 1 1 1Zm-3-4a1 1 0 1 0 1 1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rocket(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.601 2.062a1 1 0 0 0-.713-.713A11.252 11.252 0 0 0 10.47 4.972L9.354 6.296L6.75 5.668a2.777 2.777 0 0 0-3.387 1.357l-2.2 3.9a1 1 0 0 0 .661 1.469l3.073.659a13.42 13.42 0 0 0-.555 2.434a1 1 0 0 0 .284.836l3.1 3.1a1 1 0 0 0 .708.293a.838.838 0 0 0 .086-.004a12.169 12.169 0 0 0 2.492-.49l.644 3.004a1 1 0 0 0 1.469.661l3.905-2.202a3.035 3.035 0 0 0 1.375-3.304l-.668-2.76l1.237-1.137A11.204 11.204 0 0 0 22.6 2.062ZM3.572 10.723l1.556-2.76a.826.826 0 0 1 1.07-.375l1.718.416l-.65.772a13.095 13.095 0 0 0-1.59 2.398Zm12.47 8.222l-2.715 1.532l-.43-2.005a11.34 11.34 0 0 0 2.414-1.62l.743-.683l.404 1.664a1.041 1.041 0 0 1-.416 1.112m1.615-6.965l-3.685 3.386a9.773 9.773 0 0 1-5.17 2.304l-2.405-2.404a10.932 10.932 0 0 1 2.401-5.206l1.679-1.993a.964.964 0 0 0 .078-.092L11.99 6.27a9.278 9.278 0 0 1 8.81-3.12a9.218 9.218 0 0 1-3.143 8.829Zm-.923-6.164a1.5 1.5 0 1 0 1.5 1.5a1.5 1.5 0 0 0-1.5-1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RopeWay(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 6.5h-6V4h4.62a1 1 0 0 0 0-2H6.38a1 1 0 1 0 0 2H11v2.5H5a3 3 0 0 0-3 3V19a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V9.5a3 3 0 0 0-3-3M11 20H5a1 1 0 0 1-1-1v-3.75h7a.5.5 0 0 0 0 .13v4.5a.53.53 0 0 0 0 .12m9-1a1 1 0 0 1-1 1h-6a.53.53 0 0 0 0-.12v-4.5a.5.5 0 0 0 0-.13h7Zm0-5.75H4V9.5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RotateThreeHundredSixty(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 6C6.3 6 2 8.15 2 11c0 2.45 3.19 4.38 7.71 4.88l-.42.41a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l2-2a1 1 0 0 0 .21-.33a1 1 0 0 0 0-.76a1 1 0 0 0-.21-.33l-2-2a1 1 0 0 0-1.42 1.42l.12.11C6 13.34 4 12 4 11c0-1.22 3.12-3 8-3s8 1.78 8 3c0 .83-1.45 2-4.21 2.6a1 1 0 0 0-.79 1.19a1 1 0 0 0 1.19.77c3.65-.8 5.81-2.5 5.81-4.56c0-2.85-4.3-5-10-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rss(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.88 16.88a3 3 0 0 0 0 4.24a3 3 0 0 0 4.24 0a3 3 0 0 0-4.24-4.24m2.83 2.83a1 1 0 0 1-1.42-1.42a1 1 0 0 1 1.42 0a1 1 0 0 1 0 1.42M5 12a1 1 0 0 0 0 2a5 5 0 0 1 5 5a1 1 0 0 0 2 0a7 7 0 0 0-7-7m0-4a1 1 0 0 0 0 2a9 9 0 0 1 9 9a1 1 0 0 0 2 0a11.08 11.08 0 0 0-3.22-7.78A11.08 11.08 0 0 0 5 8m10.61.39A15.11 15.11 0 0 0 5 4a1 1 0 0 0 0 2a13 13 0 0 1 13 13a1 1 0 0 0 2 0a15.11 15.11 0 0 0-4.39-10.61"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RssAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.88 16.88a3 3 0 1 0 4.24 4.24a3 3 0 0 0 0-4.24a3.08 3.08 0 0 0-4.24 0m2.83 2.83a1 1 0 0 1-1.42 0a1 1 0 0 1 0-1.42a1 1 0 0 1 1.42 0a1 1 0 0 1 0 1.42M5 12a1 1 0 0 0 0 2a5 5 0 0 1 5 5a1 1 0 0 0 2 0a7 7 0 0 0-7-7m0-4a1 1 0 0 0 0 2a9 9 0 0 1 9 9a1 1 0 0 0 2 0A11 11 0 0 0 5 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RssInterface(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 14a1 1 0 0 0 0 2a3 3 0 0 1 3 3a1 1 0 0 0 2 0a5 5 0 0 0-5-5m-.71 4.29a1 1 0 1 0 1.42 0a1 1 0 0 0-1.42 0M19 4H5a3 3 0 0 0-3 3a1 1 0 0 0 2 0a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1h-4a1 1 0 0 0 0 2h4a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3M3 10a1 1 0 0 0 0 2a7 7 0 0 1 7 7a1 1 0 0 0 2 0a9 9 0 0 0-9-9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ruler(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.61 7.05L17 1.39a1 1 0 0 0-.71-.29a1 1 0 0 0-.7.29L1.39 15.54a1 1 0 0 0 0 1.41l5.66 5.66a1 1 0 0 0 .71.29a1 1 0 0 0 .7-.29l2.83-2.83l8.49-8.49l2.83-2.83a1 1 0 0 0 0-1.41m-3.54 2.12l-.71-.71a1 1 0 0 0-1.41 0a1 1 0 0 0 0 1.42l.71.71L16.24 12l-2.12-2.12a1 1 0 0 0-1.41 1.41l2.12 2.12l-1.42 1.42l-.7-.71a1 1 0 1 0-1.42 1.42l.71.7l-1.41 1.42l-2.13-2.12a1 1 0 0 0-1.41 0a1 1 0 0 0 0 1.41l2.12 2.12l-1.41 1.42l-4.25-4.25L16.24 3.51l4.25 4.25Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RulerCombined(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V10h11a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1m-1 6h-2V7a1 1 0 0 0-2 0v1h-2V7a1 1 0 0 0-2 0v1h-2V7a1 1 0 0 0-2 0v1H7a1 1 0 0 0 0 2h1v2H7a1 1 0 0 0 0 2h1v2H7a1 1 0 0 0 0 2h1v2H4V4h16Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RupeeSign(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 7h-2.21a5.49 5.49 0 0 0-1-2H18a1 1 0 0 0 0-2H7a1 1 0 0 0 0 2h3.5a3.5 3.5 0 0 1 3.15 2H7a1 1 0 0 0 0 2h7a3.5 3.5 0 0 1-3.45 3H7a.7.7 0 0 0-.14 0a.65.65 0 0 0-.2 0a.69.69 0 0 0-.19.1l-.12.07a.75.75 0 0 0-.14.17a1.1 1.1 0 0 0-.09.14a.61.61 0 0 0 0 .18A.65.65 0 0 0 6 13a.7.7 0 0 0 0 .14a.65.65 0 0 0 0 .2a.69.69 0 0 0 .1.19s0 .08.07.12l6 7a1 1 0 0 0 1.52-1.3L9.18 14h1.32A5.5 5.5 0 0 0 16 9h2a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sad(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.36 15.33a1 1 0 0 0-.13 1.4a1 1 0 0 0 1.41.13a3.76 3.76 0 0 1 4.72 0a1 1 0 0 0 .64.23a1 1 0 0 0 .64-1.76a5.81 5.81 0 0 0-7.28 0m.85-4.79a1 1 0 0 0 1.41 0a1 1 0 0 0 0-1.41a3.08 3.08 0 0 0-4.24 0a1 1 0 1 0 1.41 1.41a1 1 0 0 1 1.42 0M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8m5.62-10.87a3.08 3.08 0 0 0-4.24 0a1 1 0 0 0 1.41 1.41a1 1 0 0 1 1.42 0a1 1 0 0 0 1.41 0a1 1 0 0 0 0-1.41"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SadCry(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a10 10 0 0 0-3.41 19.39h.06a9.81 9.81 0 0 0 6.7 0h.06A10 10 0 0 0 12 2m2 17.74a7.82 7.82 0 0 1-4 0V16h4Zm2-.82V11a1 1 0 0 0 0-2h-2a1 1 0 0 0 0 2v3h-4v-3a1 1 0 0 0 0-2H8a1 1 0 0 0 0 2v7.92a8 8 0 1 1 8 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SadCrying(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a10 10 0 0 0-3.41 19.39h.06a9.81 9.81 0 0 0 6.7 0h.06A10 10 0 0 0 12 2m2 17.74a7.82 7.82 0 0 1-4 0v-3.13a3.79 3.79 0 0 1 4 0Zm2-.82V11a1 1 0 0 0 0-2h-2a1 1 0 0 0 0 2v3.39a6 6 0 0 0-4 0V11a1 1 0 0 0 0-2H8a1 1 0 0 0 0 2v7.92a8 8 0 1 1 8 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SadDizzy(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9 11.71l.29-.3l.29.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-.3-.29l.3-.29a1 1 0 0 0-1.46-1.42l-.29.3l-.25-.3a1 1 0 1 0-1.46 1.42l.3.29l-.3.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0Zm-.6 3.62a1 1 0 0 0-.13 1.4a1 1 0 0 0 1.41.13a3.76 3.76 0 0 1 4.72 0a1 1 0 0 0 1.41-.13a1 1 0 0 0-.13-1.4a5.81 5.81 0 0 0-7.32 0ZM12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8m5-11.71a1 1 0 0 0-1.42 0l-.29.3l-.29-.3a1 1 0 0 0-1.42 1.42l.3.29l-.3.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l.29-.3l.29.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-.3-.29l.3-.29a1 1 0 0 0 0-1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SadSquint(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9.08 12.21l1.5-1.5a1 1 0 0 0 0-1.42l-1.5-1.5a1 1 0 0 0-1.42 1.42l.8.79l-.8.79a1 1 0 0 0 0 1.42a1 1 0 0 0 .71.29a1 1 0 0 0 .71-.29m-.72 3.12a1 1 0 0 0-.13 1.4a1 1 0 0 0 1.41.13a3.76 3.76 0 0 1 4.72 0a1 1 0 0 0 1.41-.13a1 1 0 0 0-.13-1.4a5.81 5.81 0 0 0-7.28 0m8.22-7.54a1 1 0 0 0-1.42 0l-1.5 1.5a1 1 0 0 0 0 1.42l1.5 1.5a1 1 0 0 0 .71.29a1 1 0 0 0 .71-.29a1 1 0 0 0 0-1.42l-.79-.79l.79-.79a1 1 0 0 0 0-1.42M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sanitizer(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 12a3 3 0 1 0 3 3a3.003 3.003 0 0 0-3-3m0 4a1 1 0 1 1 1-1a1.001 1.001 0 0 1-1 1m5.8-8.4L16 5.5V3h1a1 1 0 0 0 0-2H8.657a4.967 4.967 0 0 0-3.535 1.464l-.829.829a1 1 0 1 0 1.414 1.414l.829-.829A3.022 3.022 0 0 1 8.656 3H10v2.5L7.2 7.6A3.016 3.016 0 0 0 6 10v12a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V10a3.015 3.015 0 0 0-1.2-2.4M12 3h2v2h-2Zm6 18H8V10a1.006 1.006 0 0 1 .4-.8L11.334 7h3.333L17.6 9.2a1.005 1.005 0 0 1 .4.8Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SanitizerAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 15h-2a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2m2-7V5a1 1 0 0 0-1-1h-1V3h1a1 1 0 0 0 0-2h-4.764a4.593 4.593 0 0 0-4.13 2.553a1 1 0 0 0 1.789.894A2.603 2.603 0 0 1 10.235 3H12v1h-1a1 1 0 0 0-1 1v3a3.003 3.003 0 0 0-3 3v9a3.003 3.003 0 0 0 3 3h6a3.003 3.003 0 0 0 3-3v-9a3.003 3.003 0 0 0-3-3m-4-2h2v2h-2Zm5 14a1.001 1.001 0 0 1-1 1h-6a1.001 1.001 0 0 1-1-1v-9a1.001 1.001 0 0 1 1-1h6a1.001 1.001 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Save(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.71 9.29l-6-6a1 1 0 0 0-.32-.21A1.09 1.09 0 0 0 14 3H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-8a1 1 0 0 0-.29-.71M9 5h4v2H9Zm6 14H9v-3a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1Zm4-1a1 1 0 0 1-1 1h-1v-3a3 3 0 0 0-3-3h-4a3 3 0 0 0-3 3v3H6a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h1v3a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V6.41l4 4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScalingLeft(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 16a1 1 0 0 0-1 1v1.59L13.41 12l2.13-2.12a1 1 0 0 0-1.42-1.42L12 10.59L5.41 4H7a1 1 0 0 0 0-2H3a1 1 0 0 0-.38.08a1 1 0 0 0-.54.54A1 1 0 0 0 2 3v4a1 1 0 0 0 2 0V5.41L10.59 12l-2.13 2.12a1 1 0 0 0 0 1.42a1 1 0 0 0 .71.29a1 1 0 0 0 .71-.29L12 13.41L18.59 20H17a1 1 0 0 0 0 2h4a1 1 0 0 0 .38-.08a1 1 0 0 0 .54-.54A1 1 0 0 0 22 21v-4a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScalingRight(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.92 2.62a1 1 0 0 0-.54-.54A1 1 0 0 0 21 2h-4a1 1 0 0 0 0 2h1.59L12 10.59L9.88 8.46a1 1 0 0 0-1.42 1.42L10.59 12L4 18.59V17a1 1 0 0 0-2 0v4a1 1 0 0 0 .08.38a1 1 0 0 0 .54.54A1 1 0 0 0 3 22h4a1 1 0 0 0 0-2H5.41L12 13.41l2.12 2.13a1 1 0 0 0 .71.29a1 1 0 0 0 .71-.29a1 1 0 0 0 0-1.42L13.41 12L20 5.41V7a1 1 0 0 0 2 0V3a1 1 0 0 0-.08-.38"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Scenery(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 2H5a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h14a2.81 2.81 0 0 0 .49-.05l.3-.07h.12l.37-.14l.13-.07c.1-.06.21-.11.31-.18a3.79 3.79 0 0 0 .38-.32l.07-.09a2.69 2.69 0 0 0 .27-.32l.09-.13a2.31 2.31 0 0 0 .18-.35a1 1 0 0 0 .07-.15c.05-.12.08-.25.12-.38v-.15a2.6 2.6 0 0 0 .1-.6V5a3 3 0 0 0-3-3M5 20a1 1 0 0 1-1-1v-4.31l3.29-3.3a1 1 0 0 1 1.42 0l8.6 8.61Zm15-1a1 1 0 0 1-.07.36a1 1 0 0 1-.08.14a.94.94 0 0 1-.09.12l-5.35-5.35l.88-.88a1 1 0 0 1 1.42 0l3.29 3.3Zm0-5.14L18.12 12a3.08 3.08 0 0 0-4.24 0l-.88.88L10.12 10a3.08 3.08 0 0 0-4.24 0L4 11.86V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1ZM13.5 6A1.5 1.5 0 1 0 15 7.5A1.5 1.5 0 0 0 13.5 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Schedule(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 14a1 1 0 1 0-1-1a1 1 0 0 0 1 1m5 0a1 1 0 1 0-1-1a1 1 0 0 0 1 1m-5 4a1 1 0 1 0-1-1a1 1 0 0 0 1 1m5 0a1 1 0 1 0-1-1a1 1 0 0 0 1 1M7 14a1 1 0 1 0-1-1a1 1 0 0 0 1 1M19 4h-1V3a1 1 0 0 0-2 0v1H8V3a1 1 0 0 0-2 0v1H5a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3m1 15a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-9h16Zm0-11H4V7a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1ZM7 18a1 1 0 1 0-1-1a1 1 0 0 0 1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Screw(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.71 10.12l-7.83-7.83a1 1 0 0 0-1.7.57L11.45 8l-2 2l-.33-.19A1 1 0 0 0 8 11.44l-1.15 1.17l-.33-.19a1 1 0 0 0-1.11 1.63l-1.17 1.16l-.32-.21a1 1 0 0 0-1.37.37a1 1 0 0 0 .25 1.26l-.51.51a.93.93 0 0 0-.21.33a1 1 0 0 0-.08.38V21a1 1 0 0 0 1 1h3.13a1 1 0 0 0 .38-.08a.93.93 0 0 0 .33-.21L8.54 20l.33.19a1 1 0 0 0 1.37-.36a1 1 0 0 0-.24-1.27l1.17-1.16l.33.19a1 1 0 0 0 .49.13a1 1 0 0 0 .6-1.72l1.17-1.16l.33.19a1 1 0 0 0 .49.13a1 1 0 0 0 .62-1.77l.79-.79l5.15-.73a1 1 0 0 0 .81-.68a1 1 0 0 0-.24-1.07M5.72 20H4v-1.72l.57-.57L6.75 19Zm2.49-2.5L6 16.25l1.14-1.14l2.17 1.25Zm2.61-2.6l-2.18-1.26l1.15-1.14L12 13.75Zm2.61-2.61L11.25 11l1.14-1.14l1.72 1.72Zm2.45-1.74l-2.43-2.43l.43-3l5 5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Scroll(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.29 9.71a1 1 0 0 0 1.42 0l5-5a1 1 0 1 0-1.42-1.42L12 7.59l-4.29-4.3a1 1 0 0 0-1.42 1.42Zm1.42 4.58a1 1 0 0 0-1.42 0l-5 5a1 1 0 0 0 1.42 1.42l4.29-4.3l4.29 4.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScrollH(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.71 6.29a1 1 0 0 0-1.42 1.42L7.59 12l-4.3 4.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l5-5a1 1 0 0 0 0-1.42ZM16.41 12l4.3-4.29a1 1 0 1 0-1.42-1.42l-5 5a1 1 0 0 0 0 1.42l5 5a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Search(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.71 20.29L18 16.61A9 9 0 1 0 16.61 18l3.68 3.68a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.39M11 18a7 7 0 1 1 7-7a7 7 0 0 1-7 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.07 16.83L19 14.71a3.08 3.08 0 0 0-3.4-.57l-.9-.9a7 7 0 1 0-1.41 1.41l.89.89a3 3 0 0 0 .53 3.46l2.12 2.12a3 3 0 0 0 4.24 0a3 3 0 0 0 0-4.29m-8.48-4.24a5 5 0 1 1 0-7.08a5 5 0 0 1 0 7.08m7.07 7.07a1 1 0 0 1-1.42 0l-2.12-2.12a1 1 0 0 1 0-1.42a1 1 0 0 1 1.42 0l2.12 2.12a1 1 0 0 1 0 1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchMinus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.71 20.29L18 16.61A9 9 0 1 0 16.61 18l3.68 3.68a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.39M11 18a7 7 0 1 1 7-7a7 7 0 0 1-7 7m4-8H7a1 1 0 0 0 0 2h8a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchPlus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 10h-3V7a1 1 0 0 0-2 0v3H7a1 1 0 0 0 0 2h3v3a1 1 0 0 0 2 0v-3h3a1 1 0 0 0 0-2m6.71 10.29L18 16.61A9 9 0 1 0 16.61 18l3.68 3.68a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.39M11 18a7 7 0 1 1 7-7a7 7 0 0 1-7 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Selfie(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 2H8a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3m1 17a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1v-1h10Zm-5-5a3 3 0 0 1 2.82 2H9.18A3 3 0 0 1 12 14m-1-3a1 1 0 1 1 1 1a1 1 0 0 1-1-1m6 5h-.1a5 5 0 0 0-2.42-3.32A3 3 0 0 0 15 11a3 3 0 0 0-6 0a3 3 0 0 0 .52 1.68A5 5 0 0 0 7.1 16H7V5a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1ZM12 5a1 1 0 1 0 1 1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Server(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 17a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-6 0H6a1 1 0 0 0 0 2h3a1 1 0 0 0 0-2m9 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-3-6a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-6 0H6a1 1 0 0 0 0 2h3a1 1 0 0 0 0-2m9-6a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0 6a1 1 0 1 0 1 1a1 1 0 0 0-1-1m4-6a3 3 0 0 0-3-3H5a3 3 0 0 0-3 3v2a3 3 0 0 0 .78 2A3 3 0 0 0 2 11v2a3 3 0 0 0 .78 2A3 3 0 0 0 2 17v2a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-2a3 3 0 0 0-.78-2a3 3 0 0 0 .78-2v-2a3 3 0 0 0-.78-2A3 3 0 0 0 22 7Zm-2 14a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Zm0-6a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Zm0-6a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Zm-5-2a1 1 0 1 0 1 1a1 1 0 0 0-1-1M9 5H6a1 1 0 0 0 0 2h3a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ServerAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 17a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0-6a1 1 0 1 0 1 1a1 1 0 0 0-1-1m0-6a1 1 0 1 0 1 1a1 1 0 0 0-1-1m12 0a3 3 0 0 0-3-3H7a3 3 0 0 0-3 3v2a3 3 0 0 0 .78 2A3 3 0 0 0 4 11v2a3 3 0 0 0 .78 2A3 3 0 0 0 4 17v2a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-2a3 3 0 0 0-.78-2a3 3 0 0 0 .78-2v-2a3 3 0 0 0-.78-2A3 3 0 0 0 20 7Zm-2 14a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1Zm0-6a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1Zm0-6a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ServerConnection(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 13a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-4 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1m15-9a3 3 0 0 0-3-3H5a3 3 0 0 0-3 3v4a3 3 0 0 0 .78 2A3 3 0 0 0 2 12v4a3 3 0 0 0 3 3h6v2H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2h-8v-2h6a3 3 0 0 0 3-3v-4a3 3 0 0 0-.78-2A3 3 0 0 0 22 8Zm-2 12a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Zm0-8a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Zm-9-3a1 1 0 1 0 1 1a1 1 0 0 0-1-1M7 5a1 1 0 1 0 1 1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ServerNetwork(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 6a1 1 0 1 0-1-1a1 1 0 0 0 1 1m13 13h-6.18A3 3 0 0 0 13 17.18V15h4a3 3 0 0 0 3-3v-2a3 3 0 0 0-.78-2A3 3 0 0 0 20 6V4a3 3 0 0 0-3-3H7a3 3 0 0 0-3 3v2a3 3 0 0 0 .78 2A3 3 0 0 0 4 10v2a3 3 0 0 0 3 3h4v2.18A3 3 0 0 0 9.18 19H3a1 1 0 0 0 0 2h6.18a3 3 0 0 0 5.64 0H21a1 1 0 0 0 0-2M6 4a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1Zm1 9a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1Zm5 8a1 1 0 1 1 1-1a1 1 0 0 1-1 1M8 10a1 1 0 1 0 1 1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ServerNetworkAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 6h3a1 1 0 0 0 0-2H8a1 1 0 0 0 0 2m8 0a1 1 0 1 0-1-1a1 1 0 0 0 1 1m0 4a1 1 0 1 0 1 1a1 1 0 0 0-1-1m5 9h-6.18A3 3 0 0 0 13 17.18V15h4a3 3 0 0 0 3-3v-2a3 3 0 0 0-.78-2A3 3 0 0 0 20 6V4a3 3 0 0 0-3-3H7a3 3 0 0 0-3 3v2a3 3 0 0 0 .78 2A3 3 0 0 0 4 10v2a3 3 0 0 0 3 3h4v2.18A3 3 0 0 0 9.18 19H3a1 1 0 0 0 0 2h6.18a3 3 0 0 0 5.64 0H21a1 1 0 0 0 0-2M6 4a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1Zm1 9a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1Zm5 8a1 1 0 1 1 1-1a1 1 0 0 1-1 1m-1-11H8a1 1 0 0 0 0 2h3a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Servers(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 17a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-4 0H9a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2m0-6H9a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2m4 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1m8-3a3 3 0 0 0-3-3h-1a3 3 0 0 0-3-3H8a3 3 0 0 0-3 3H4a3 3 0 0 0-3 3v2a3 3 0 0 0 .78 2A3 3 0 0 0 1 14v2a3 3 0 0 0 3 3h1a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3h1a3 3 0 0 0 3-3v-2a3 3 0 0 0-.78-2a3 3 0 0 0 .78-2ZM5 17H4a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1h1a3 3 0 0 0 .78 2A3 3 0 0 0 5 17m0-6H4a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1h1a3 3 0 0 0 .78 2A3 3 0 0 0 5 11m12 8a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1Zm0-6a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1Zm0-6a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1Zm4 9a1 1 0 0 1-1 1h-1a3 3 0 0 0-.78-2a3 3 0 0 0 .78-2h1a1 1 0 0 1 1 1Zm0-6a1 1 0 0 1-1 1h-1a3 3 0 0 0-.78-2A3 3 0 0 0 19 7h1a1 1 0 0 1 1 1Zm-6-5a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-4 0H9a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Servicemark(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.5 9h4a1 1 0 0 0 0-2h-4a3 3 0 0 0 0 6h2a1 1 0 0 1 0 2h-4a1 1 0 0 0 0 2h4a3 3 0 0 0 0-6h-2a1 1 0 0 1 0-2m15.92-1.38a1 1 0 0 0-.54-.54a1 1 0 0 0-.76 0a1 1 0 0 0-.33.21L17 10.09l-2.79-2.8a1 1 0 0 0-.33-.21a1 1 0 0 0-.76 0a1 1 0 0 0-.54.54a1 1 0 0 0-.08.38v8a1 1 0 0 0 2 0v-5.59l1.79 1.8a1 1 0 0 0 1.42 0l1.79-1.8V16a1 1 0 0 0 2 0V8a1 1 0 0 0-.08-.38"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Setting(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.9 12.66a1 1 0 0 1 0-1.32l1.28-1.44a1 1 0 0 0 .12-1.17l-2-3.46a1 1 0 0 0-1.07-.48l-1.88.38a1 1 0 0 1-1.15-.66l-.61-1.83a1 1 0 0 0-.95-.68h-4a1 1 0 0 0-1 .68l-.56 1.83a1 1 0 0 1-1.15.66L5 4.79a1 1 0 0 0-1 .48L2 8.73a1 1 0 0 0 .1 1.17l1.27 1.44a1 1 0 0 1 0 1.32L2.1 14.1a1 1 0 0 0-.1 1.17l2 3.46a1 1 0 0 0 1.07.48l1.88-.38a1 1 0 0 1 1.15.66l.61 1.83a1 1 0 0 0 1 .68h4a1 1 0 0 0 .95-.68l.61-1.83a1 1 0 0 1 1.15-.66l1.88.38a1 1 0 0 0 1.07-.48l2-3.46a1 1 0 0 0-.12-1.17ZM18.41 14l.8.9l-1.28 2.22l-1.18-.24a3 3 0 0 0-3.45 2L12.92 20h-2.56L10 18.86a3 3 0 0 0-3.45-2l-1.18.24l-1.3-2.21l.8-.9a3 3 0 0 0 0-4l-.8-.9l1.28-2.2l1.18.24a3 3 0 0 0 3.45-2L10.36 4h2.56l.38 1.14a3 3 0 0 0 3.45 2l1.18-.24l1.28 2.22l-.8.9a3 3 0 0 0 0 3.98m-6.77-6a4 4 0 1 0 4 4a4 4 0 0 0-4-4m0 6a2 2 0 1 1 2-2a2 2 0 0 1-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SeventeenPlus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 7v2a1 1 0 0 0 2 0V8h2.78L14 16.8a1 1 0 0 0 .8 1.2h.2a1 1 0 0 0 1-.8l2-10a1 1 0 0 0-.21-.83A1 1 0 0 0 17 6h-5a1 1 0 0 0-1 1m7-2h1v1a1 1 0 0 0 2 0V5h1a1 1 0 0 0 0-2h-1V2a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2M7 7v10a1 1 0 0 0 2 0V7a1 1 0 0 0-2 0m14.6 2a1 1 0 0 0-.78 1.18a9 9 0 1 1-7-7a1 1 0 1 0 .4-2A10.8 10.8 0 0 0 12 1a11 11 0 1 0 11 11a10.8 10.8 0 0 0-.22-2.2A1 1 0 0 0 21.6 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Share(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.707 11.293l-8-8A1 1 0 0 0 12 4v3.545A11.015 11.015 0 0 0 2 18.5V20a1 1 0 0 0 1.784.62a11.456 11.456 0 0 1 7.887-4.049c.05-.006.175-.016.329-.026V20a1 1 0 0 0 1.707.707l8-8a1 1 0 0 0 0-1.414M14 17.586V15.5a1 1 0 0 0-1-1c-.255 0-1.296.05-1.562.085a14.005 14.005 0 0 0-7.386 2.948A9.013 9.013 0 0 1 13 9.5a1 1 0 0 0 1-1V6.414L19.586 12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 14a4 4 0 0 0-3.08 1.48l-5.1-2.35a3.64 3.64 0 0 0 0-2.26l5.1-2.35A4 4 0 1 0 14 6a4.17 4.17 0 0 0 .07.71L8.79 9.14a4 4 0 1 0 0 5.72l5.28 2.43A4.17 4.17 0 0 0 14 18a4 4 0 1 0 4-4m0-10a2 2 0 1 1-2 2a2 2 0 0 1 2-2M6 14a2 2 0 1 1 2-2a2 2 0 0 1-2 2m12 6a2 2 0 1 1 2-2a2 2 0 0 1-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shield(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.63 3.65a1 1 0 0 0-.84-.2a8 8 0 0 1-6.22-1.27a1 1 0 0 0-1.14 0a8 8 0 0 1-6.22 1.27a1 1 0 0 0-.84.2a1 1 0 0 0-.37.78v7.45a9 9 0 0 0 3.77 7.33l3.65 2.6a1 1 0 0 0 1.16 0l3.65-2.6A9 9 0 0 0 20 11.88V4.43a1 1 0 0 0-.37-.78M18 11.88a7 7 0 0 1-2.93 5.7L12 19.77l-3.07-2.19A7 7 0 0 1 6 11.88v-6.3a10 10 0 0 0 6-1.39a10 10 0 0 0 6 1.39Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldCheck(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.63 3.65a1 1 0 0 0-.84-.2a8 8 0 0 1-6.22-1.27a1 1 0 0 0-1.14 0a8 8 0 0 1-6.22 1.27a1 1 0 0 0-.84.2a1 1 0 0 0-.37.78v7.45a9 9 0 0 0 3.77 7.33l3.65 2.6a1 1 0 0 0 1.16 0l3.65-2.6A9 9 0 0 0 20 11.88V4.43a1 1 0 0 0-.37-.78M18 11.88a7 7 0 0 1-2.93 5.7L12 19.77l-3.07-2.19A7 7 0 0 1 6 11.88v-6.3a10 10 0 0 0 6-1.39a10 10 0 0 0 6 1.39Zm-4.46-2.29l-2.69 2.7l-.89-.9a1 1 0 0 0-1.42 1.42l1.6 1.6a1 1 0 0 0 1.42 0L15 11a1 1 0 0 0-1.42-1.42Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldExclamation(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.92 15a.56.56 0 0 0-.09-.17l-.12-.15a1 1 0 0 0-1.42 0a.61.61 0 0 0-.12.15a.56.56 0 0 0-.09.17a.7.7 0 0 0-.06.19a1.23 1.23 0 0 0 0 .19a.88.88 0 0 0 .08.39a1 1 0 0 0 1.3.54a1.19 1.19 0 0 0 .33-.22a1 1 0 0 0 .21-.32a1 1 0 0 0 .08-.39a1.23 1.23 0 0 0 0-.19a.7.7 0 0 0-.1-.19M12 7.36a1 1 0 0 0-1 1v3a1 1 0 0 0 2 0v-3a1 1 0 0 0-1-1m7.63-3.71a1 1 0 0 0-.84-.2a8 8 0 0 1-6.22-1.27a1 1 0 0 0-1.14 0a8 8 0 0 1-6.22 1.27a1 1 0 0 0-.84.2a1 1 0 0 0-.37.78v7.45a9 9 0 0 0 3.77 7.33l3.65 2.6a1 1 0 0 0 1.16 0l3.65-2.6A9 9 0 0 0 20 11.88V4.43a1 1 0 0 0-.37-.78M18 11.88a7 7 0 0 1-2.93 5.7L12 19.77l-3.07-2.19A7 7 0 0 1 6 11.88v-6.3a10 10 0 0 0 6-1.39a10 10 0 0 0 6 1.39Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldPlus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 11h-1v-1a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2h1v1a1 1 0 0 0 2 0v-1h1a1 1 0 0 0 0-2m5.63-7.35a1.007 1.007 0 0 0-.835-.203a7.98 7.98 0 0 1-6.223-1.267a.999.999 0 0 0-1.144 0a7.976 7.976 0 0 1-6.223 1.267A1 1 0 0 0 4 4.427v7.456a9.019 9.019 0 0 0 3.769 7.324l3.65 2.607a1 1 0 0 0 1.162 0l3.65-2.607A9.017 9.017 0 0 0 20 11.883V4.426a1.001 1.001 0 0 0-.37-.776M18 11.883a7.016 7.016 0 0 1-2.93 5.696L12 19.771L8.93 17.58A7.017 7.017 0 0 1 6 11.883v-6.3a9.955 9.955 0 0 0 6-1.391a9.955 9.955 0 0 0 6 1.391Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldQuestion(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.29 14.66a1 1 0 0 0-.29.7a1 1 0 0 0 .08.39a1 1 0 0 0 1.92-.39a1 1 0 0 0-.29-.7a1 1 0 0 0-1.42 0m8.34-11a1 1 0 0 0-.84-.2a8 8 0 0 1-6.22-1.27a1 1 0 0 0-1.14 0a8 8 0 0 1-6.22 1.26a1 1 0 0 0-.84.2a1 1 0 0 0-.37.78v7.45a9 9 0 0 0 3.77 7.33l3.65 2.6a1 1 0 0 0 1.16 0l3.65-2.6A9 9 0 0 0 20 11.88V4.43a1 1 0 0 0-.37-.78ZM18 11.88a7 7 0 0 1-2.93 5.7L12 19.77l-3.07-2.19A7 7 0 0 1 6 11.88v-6.3a10 10 0 0 0 6-1.39a10 10 0 0 0 6 1.39Zm-6-4.52a3 3 0 0 0-2.6 1.5a1 1 0 0 0 1.73 1a1 1 0 1 1 .87 1.5a1 1 0 0 0 0 2a3 3 0 1 0 0-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.71 20.29l-16-16l-2-2a1 1 0 0 0-1.42 1.42L4 5.41v6.47a9 9 0 0 0 3.77 7.32l3.65 2.61a1 1 0 0 0 1.16 0l3.65-2.61a8.21 8.21 0 0 0 .86-.7l3.2 3.21a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42m-6.64-2.71L12 19.77l-3.07-2.19A7 7 0 0 1 6 11.88V7.41l9.67 9.68c-.19.17-.39.33-.6.49m-5-12.51A10.15 10.15 0 0 0 12 4.19a9.82 9.82 0 0 0 6 1.39v6.3a6.88 6.88 0 0 1-.1 1.18a1 1 0 0 0 .83 1.15h.16a1 1 0 0 0 1-.84a9.77 9.77 0 0 0 .12-1.5V4.43a1 1 0 0 0-.37-.77a1 1 0 0 0-.83-.21a7.89 7.89 0 0 1-6.23-1.27a1 1 0 0 0-1.14 0a8 8 0 0 1-2 1a1 1 0 1 0 .64 1.9Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ship(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.62 17.28a1 1 0 0 0 1.86-.74l-1.12-2.82L11 12.25V17a1 1 0 0 0 2 0v-4.75l6.64 1.47l-1.12 2.82a1 1 0 0 0 .56 1.3a1 1 0 0 0 .37.07a1 1 0 0 0 .93-.63l1.55-3.91a1 1 0 0 0-.05-.84a1 1 0 0 0-.66-.51L18 11.31V7a1 1 0 0 0-1-1h-2V3a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1v3H7a1 1 0 0 0-1 1v4.31L2.78 12a1 1 0 0 0-.66.51a1 1 0 0 0-.05.84ZM11 4h2v2h-2ZM8 8h8v2.86L12.22 10h-.44L8 10.86Zm12.71 11.28a4.38 4.38 0 0 0-1 .45a2.08 2.08 0 0 1-2.1 0a4.62 4.62 0 0 0-4.54 0a2.14 2.14 0 0 1-2.12 0a4.64 4.64 0 0 0-4.54 0a2.08 2.08 0 0 1-2.1 0a4.38 4.38 0 0 0-1-.45A1 1 0 0 0 2 20a1 1 0 0 0 .67 1.24a2.1 2.1 0 0 1 .57.25a4 4 0 0 0 2 .55a4.14 4.14 0 0 0 2.08-.56a2.65 2.65 0 0 1 2.56 0a4.15 4.15 0 0 0 4.12 0a2.65 2.65 0 0 1 2.56 0a4 4 0 0 0 4.1 0a2.1 2.1 0 0 1 .57-.25A1 1 0 0 0 22 20a1 1 0 0 0-1.29-.72"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shop(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 5H2a1 1 0 0 0-1 1v4a3 3 0 0 0 2 2.82V22a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1v-9.18A3 3 0 0 0 23 10V6a1 1 0 0 0-1-1m-7 2h2v3a1 1 0 0 1-2 0Zm-4 0h2v3a1 1 0 0 1-2 0ZM7 7h2v3a1 1 0 0 1-2 0Zm-3 4a1 1 0 0 1-1-1V7h2v3a1 1 0 0 1-1 1m10 10h-4v-2a2 2 0 0 1 4 0Zm5 0h-3v-2a4 4 0 0 0-8 0v2H5v-8.18a3.17 3.17 0 0 0 1-.6a3 3 0 0 0 4 0a3 3 0 0 0 4 0a3 3 0 0 0 4 0a3.17 3.17 0 0 0 1 .6Zm2-11a1 1 0 0 1-2 0V7h2ZM4.3 3H20a1 1 0 0 0 0-2H4.3a1 1 0 0 0 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBag(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 7h-3V6a4 4 0 0 0-8 0v1H5a1 1 0 0 0-1 1v11a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V8a1 1 0 0 0-1-1m-9-1a2 2 0 0 1 4 0v1h-4Zm8 13a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V9h2v1a1 1 0 0 0 2 0V9h4v1a1 1 0 0 0 2 0V9h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBasket(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 18a1 1 0 0 0 1-1v-2a1 1 0 0 0-2 0v2a1 1 0 0 0 1 1m-4 0a1 1 0 0 0 1-1v-2a1 1 0 0 0-2 0v2a1 1 0 0 0 1 1m9-12h-1.38l-1.73-3.45a1 1 0 1 0-1.78.9L15.38 6H8.62l1.27-2.55a1 1 0 0 0-1.78-.9L6.38 6H5a3 3 0 0 0-.92 5.84l.74 7.46a3 3 0 0 0 3 2.7h8.38a3 3 0 0 0 3-2.7l.74-7.46A3 3 0 0 0 19 6m-1.81 13.1a1 1 0 0 1-1 .9H7.81a1 1 0 0 1-1-.9L6.1 12h11.8ZM19 10H5a1 1 0 0 1 0-2h14a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingCart(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 19a1.5 1.5 0 1 0 1.5 1.5A1.5 1.5 0 0 0 8.5 19M19 16H7a1 1 0 0 1 0-2h8.491a3.013 3.013 0 0 0 2.885-2.176l1.585-5.55A1 1 0 0 0 19 5H6.74a3.007 3.007 0 0 0-2.82-2H3a1 1 0 0 0 0 2h.921a1.005 1.005 0 0 1 .962.725l.155.545v.005l1.641 5.742A3 3 0 0 0 7 18h12a1 1 0 0 0 0-2m-1.326-9l-1.22 4.274a1.005 1.005 0 0 1-.963.726H8.754l-.255-.892L7.326 7ZM16.5 19a1.5 1.5 0 1 0 1.5 1.5a1.5 1.5 0 0 0-1.5-1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingCartAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.5 15a3 3 0 0 0-1.9-2.78l1.87-7a1 1 0 0 0-.18-.87A1 1 0 0 0 20.5 4H6.8l-.33-1.26A1 1 0 0 0 5.5 2h-2v2h1.23l2.48 9.26a1 1 0 0 0 1 .74H18.5a1 1 0 0 1 0 2h-13a1 1 0 0 0 0 2h1.18a3 3 0 1 0 5.64 0h2.36a3 3 0 1 0 5.82 1a2.94 2.94 0 0 0-.4-1.47A3 3 0 0 0 21.5 15m-3.91-3H9L7.34 6H19.2ZM9.5 20a1 1 0 1 1 1-1a1 1 0 0 1-1 1m8 0a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shovel(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.71 7.38l-5.09-5.09a1 1 0 0 0-1.41 0a1 1 0 0 0 0 1.42L17 5.54L11.58 11l-1-1a3 3 0 0 0-4.25 0l-3.45 3.42A3 3 0 0 0 2 15.55V19a3 3 0 0 0 3 3h3.45a3 3 0 0 0 2.13-.88L14 17.69a3 3 0 0 0 0-4.25l-1-1L18.46 7l1.83 1.83a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.45m-9.11 8.89l-3.44 3.44a1 1 0 0 1-.71.29H5a1 1 0 0 1-1-1v-3.45a1 1 0 0 1 .29-.71l3.44-3.44a1 1 0 0 1 1.41 0l1 1l-.89.9a1 1 0 0 0 0 1.41A1 1 0 0 0 10 15a1 1 0 0 0 .7-.29l.9-.89l1 1a1 1 0 0 1 0 1.45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shrink(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.71 11.29l-2.5-2.5a1 1 0 1 0-1.42 1.42l.8.79H3a1 1 0 0 0 0 2h4.59l-.8.79a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l2.5-2.5a1 1 0 0 0 .21-.33a1 1 0 0 0 0-.76a1 1 0 0 0-.21-.33M21 11h-4.59l.8-.79a1 1 0 0 0-1.42-1.42l-2.5 2.5a1 1 0 0 0-.21.33a1 1 0 0 0 0 .76a1 1 0 0 0 .21.33l2.5 2.5a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-.8-.79H21a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shuffle(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 10a1 1 0 0 0 1-1V5.41L8.56 10A1 1 0 0 0 10 10a1 1 0 0 0 0-1.41L5.41 4H9a1 1 0 0 0 0-2H3a1 1 0 0 0-.38.08a1 1 0 0 0-.54.54A1 1 0 0 0 2 3v6a1 1 0 0 0 1 1m12.3 4a1 1 0 0 0-1.41 1.41l6.27 6.28a1 1 0 0 0 .71.29a1 1 0 0 0 .71-.29a1 1 0 0 0 0-1.42ZM9 20H5.41l16.3-16.29a1 1 0 1 0-1.42-1.42L4 18.59V15a1 1 0 0 0-2 0v6a1 1 0 0 0 .08.38a1 1 0 0 0 .54.54A1 1 0 0 0 3 22h6a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shutter(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.07 4.93A10 10 0 1 0 4.93 19.07A10 10 0 1 0 19.07 4.93M18.23 7h-5.47l2.35-2.35A8.14 8.14 0 0 1 18.23 7M9 4.6a8.15 8.15 0 0 1 3.87-.54L9 7.93ZM7 5.77v5.47L5.19 9.43l-.54-.54A8.14 8.14 0 0 1 7 5.77M4.6 15a8.12 8.12 0 0 1-.54-3.87L7.93 15Zm1.17 2h5.47l-2.35 2.35A8.14 8.14 0 0 1 5.77 17M15 19.4a8.13 8.13 0 0 1-3.87.54L15 16.07Zm0-6.16L13.24 15h-2.49L9 13.24v-2.48L10.76 9h2.48L15 10.76Zm2 5v-5.48l2.35 2.35A8.14 8.14 0 0 1 17 18.23ZM16.07 9h3.33a8.13 8.13 0 0 1 .54 3.87Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShutterAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 11a1 1 0 0 1 0-.16c0-.28-.08-.56-.13-.84A9.54 9.54 0 0 0 21 7.62a10 10 0 0 0-7-5.41l-.84-.13h-.53L12 2h-1.16l-.84.2A10 10 0 0 0 2.21 10c-.05.28-.09.56-.13.84a1 1 0 0 1 0 .16v2a1 1 0 0 1 0 .16c0 .28.08.56.13.84A9.54 9.54 0 0 0 3 16.38a10 10 0 0 0 7 5.41l.84.13l.16.08h1.01c.34 0 .68 0 1-.05h.16l.83-.15a10 10 0 0 0 7.79-7.8c.05-.28.09-.56.13-.84A1 1 0 0 1 22 13c0-.33.05-.67.05-1s-.05-.68-.05-1m-8.84-6.9l.5.07A8 8 0 0 1 18.24 7h-6.82Zm-2.74.08l.4-.06L7.38 10L5.7 7.08a8 8 0 0 1 4.67-2.91ZM4.59 15a8 8 0 0 1-.42-1.37c0-.22-.08-.45-.1-.68a.5.5 0 0 1 0-.12a8.22 8.22 0 0 1 0-1.62a.5.5 0 0 1 0-.12c0-.23.06-.46.1-.68a7.76 7.76 0 0 1 .38-1.31L8 15Zm6.3 4.91l-.5-.07A8 8 0 0 1 5.76 17h6.82ZM10.27 15l-1.73-3l1.73-3h3.46l1.73 3l-1.73 3Zm3.36 4.83l-.4.06L16.62 14l1.68 2.92a8 8 0 0 1-4.67 2.91m6.33-7a.5.5 0 0 1 0 .12c0 .23-.06.46-.1.68a7.76 7.76 0 0 1-.38 1.27L16 9h3.37a8 8 0 0 1 .42 1.37c0 .22.08.45.1.68a.5.5 0 0 1 0 .12a8.22 8.22 0 0 1 0 1.62Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sick(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 11a1 1 0 1 0-1-1a1 1 0 0 0 1 1m6 0a1 1 0 1 0-1-1a1 1 0 0 0 1 1m.27 3.2a1 1 0 0 0-1.2 0l-.74.55l-.73-.55a1 1 0 0 0-1.2 0l-.73.55l-.74-.55a1 1 0 0 0-1.2 0l-1.33 1a1 1 0 1 0 1.2 1.6l.73-.55l.74.55l.12.06a.83.83 0 0 0 .22.08h.12a1 1 0 0 0 .25 0h.1a1.06 1.06 0 0 0 .34-.16l.73-.55l.73.55a1 1 0 0 0 1 .11l.1-.05a.39.39 0 0 0 .11-.06l.74-.55l.73.55a1 1 0 0 0 .6.2a1 1 0 0 0 .8-.4a1 1 0 0 0-.2-1.4ZM12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sigma(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 16h-5.59l3.3-3.29a1 1 0 0 0 0-1.42L10.41 8H16a1 1 0 0 0 0-2H8a1 1 0 0 0-.92.62a1 1 0 0 0 .21 1.09l4.3 4.29l-4.3 4.29a1 1 0 0 0-.21 1.09A1 1 0 0 0 8 18h8a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.78 11.88l-2-2.5A1 1 0 0 0 19 9h-6V3a1 1 0 0 0-2 0v1H5a1 1 0 0 0-.78.38l-2 2.5a1 1 0 0 0 0 1.24l2 2.5A1 1 0 0 0 5 11h6v9H9a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2h-2v-4h6a1 1 0 0 0 .78-.38l2-2.5a1 1 0 0 0 0-1.24M11 9H5.48l-1.2-1.5L5.48 6H11Zm7.52 5H13v-3h5.52l1.2 1.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignInAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.5 15.1a1 1 0 0 0-1.34.45A8 8 0 1 1 12 4a7.93 7.93 0 0 1 7.16 4.45a1 1 0 0 0 1.8-.9a10 10 0 1 0 0 8.9a1 1 0 0 0-.46-1.35M21 11h-9.59l2.3-2.29a1 1 0 1 0-1.42-1.42l-4 4a1 1 0 0 0-.21.33a1 1 0 0 0 0 .76a1 1 0 0 0 .21.33l4 4a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L11.41 13H21a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignLeft(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 5h-3V3a1 1 0 0 0-2 0v2H6a1 1 0 0 0-.78.38l-2 2.5a1 1 0 0 0 0 1.24l2 2.5A1 1 0 0 0 6 12h5v8H9a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2h-2v-8h3a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1m-1 5H6.48l-1.2-1.5L6.48 7H15Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignOutAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12.59 13l-2.3 2.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l4-4a1 1 0 0 0 .21-.33a1 1 0 0 0 0-.76a1 1 0 0 0-.21-.33l-4-4a1 1 0 1 0-1.42 1.42l2.3 2.29H3a1 1 0 0 0 0 2ZM12 2a10 10 0 0 0-9 5.55a1 1 0 0 0 1.8.9A8 8 0 1 1 12 20a7.93 7.93 0 0 1-7.16-4.45a1 1 0 0 0-1.8.9A10 10 0 1 0 12 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignRight(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 6v5a1 1 0 0 0 1 1h3v8H9a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2h-2v-8h5a1 1 0 0 0 .78-.37l2-2.5a1 1 0 0 0 0-1.25l-2-2.5A1 1 0 0 0 18 5h-5V3a1 1 0 0 0-2 0v2H8a1 1 0 0 0-1 1m2 1h8.52l1.2 1.5l-1.2 1.5H9Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Signal(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 15a1 1 0 0 0-1 1v3a1 1 0 0 0 2 0v-3a1 1 0 0 0-1-1m4-3a1 1 0 0 0-1 1v6a1 1 0 0 0 2 0v-6a1 1 0 0 0-1-1m8-8a1 1 0 0 0-1 1v14a1 1 0 0 0 2 0V5a1 1 0 0 0-1-1m-4 4a1 1 0 0 0-1 1v10a1 1 0 0 0 2 0V9a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignalAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 14a1 1 0 0 0-1 1v6a1 1 0 0 0 2 0v-6a1 1 0 0 0-1-1m-5 4a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0v-2a1 1 0 0 0-1-1M20 2a1 1 0 0 0-1 1v18a1 1 0 0 0 2 0V3a1 1 0 0 0-1-1m-5 7a1 1 0 0 0-1 1v11a1 1 0 0 0 2 0V10a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignalAltThree(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 9h-4a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1V10a1 1 0 0 0-1-1m-1 12h-2V11h2Zm9-20h-4a1 1 0 0 0-1 1v20a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1m-1 20h-2V3h2ZM6 15H2a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1m-1 6H3v-4h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Signin(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 12a1 1 0 0 0-1-1h-7.59l2.3-2.29a1 1 0 1 0-1.42-1.42l-4 4a1 1 0 0 0-.21.33a1 1 0 0 0 0 .76a1 1 0 0 0 .21.33l4 4a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L11.41 13H19a1 1 0 0 0 1-1M17 2H7a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-3a1 1 0 0 0-2 0v3a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v3a1 1 0 0 0 2 0V5a3 3 0 0 0-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Signout(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 12a1 1 0 0 0 1 1h7.59l-2.3 2.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l4-4a1 1 0 0 0 .21-.33a1 1 0 0 0 0-.76a1 1 0 0 0-.21-.33l-4-4a1 1 0 1 0-1.42 1.42l2.3 2.29H5a1 1 0 0 0-1 1M17 2H7a3 3 0 0 0-3 3v3a1 1 0 0 0 2 0V5a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1v-3a1 1 0 0 0-2 0v3a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Silence(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 11a1 1 0 1 0-1-1a1 1 0 0 0 1 1m6 0a1 1 0 1 0-1-1a1 1 0 0 0 1 1m1 3a1 1 0 0 0-2 0h-1a1 1 0 0 0-2 0h-1a1 1 0 0 0-2 0a1 1 0 0 0 0 2a1 1 0 0 0 2 0h1a1 1 0 0 0 2 0h1a1 1 0 0 0 2 0a1 1 0 0 0 0-2M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SilentSquint(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.66 12.21a1 1 0 0 0 .71.29a1 1 0 0 0 .71-.29l1.5-1.5a1 1 0 0 0 0-1.42l-1.5-1.5a1 1 0 0 0-1.42 1.42l.8.79l-.8.79a1 1 0 0 0 0 1.42m7.5 0a1 1 0 0 0 .71.29a1 1 0 0 0 .71-.29a1 1 0 0 0 0-1.42l-.79-.79l.79-.79a1 1 0 1 0-1.42-1.42l-1.5 1.5a1 1 0 0 0 0 1.42Zm.11 2a1 1 0 0 0-1.2 0l-.74.55l-.73-.55a1 1 0 0 0-1.2 0l-.73.55l-.74-.55a1 1 0 0 0-1.2 0l-1.33 1a1 1 0 1 0 1.2 1.6l.73-.55l.74.55a.67.67 0 0 0 .12.06a.83.83 0 0 0 .22.08h.48a1.12 1.12 0 0 0 .33-.16l.73-.55l.73.55a1 1 0 0 0 1 .11l.1-.05a.39.39 0 0 0 .11-.06l.74-.55l.73.55a1 1 0 0 0 .6.2a1 1 0 0 0 .8-.4a1 1 0 0 0-.2-1.4ZM12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SimCard(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 3.5H7A3.5 3.5 0 0 0 3.5 7v10A3.5 3.5 0 0 0 7 20.5h10a3.5 3.5 0 0 0 3.5-3.5V7A3.5 3.5 0 0 0 17 3.5m-6.5 2h3v3h-3Zm-2 13H7A1.5 1.5 0 0 1 5.5 17v-1.5h3Zm5 0h-3v-3h3Zm5-1.5a1.5 1.5 0 0 1-1.5 1.5h-1.5v-3h3Zm0-3.5h-13V7A1.5 1.5 0 0 1 7 5.5h1.5v4a1 1 0 0 0 1 1h9Zm0-5h-3v-3H17A1.5 1.5 0 0 1 18.5 7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sitemap(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 15h-2v-3a1 1 0 0 0-1-1h-6V9h2a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1H9a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h2v2H5a1 1 0 0 0-1 1v3H2a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1H6v-2h12v2h-2a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1v-6a1 1 0 0 0-1-1M7 17v4H3v-4Zm3-10V3h4v4Zm11 14h-4v-4h4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SixPlus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 5h1v1a1 1 0 0 0 2 0V5h1a1 1 0 0 0 0-2h-1V2a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2m-6.5 3h1a1 1 0 0 1 1 1a1 1 0 0 0 2 0a3 3 0 0 0-3-3h-1a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3h1a3 3 0 0 0 3-3v-1a3 3 0 0 0-3-3h-2V9a1 1 0 0 1 1-1m1 5a1 1 0 0 1 1 1v1a1 1 0 0 1-1 1h-1a1 1 0 0 1-1-1v-2Zm9.1-4a1 1 0 0 0-.78 1.18a9 9 0 1 1-7-7a1 1 0 1 0 .4-2A10.8 10.8 0 0 0 12 1a11 11 0 1 0 11 11a10.8 10.8 0 0 0-.22-2.2A1 1 0 0 0 21.6 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SixteenPlus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 5h1v1a1 1 0 0 0 2 0V5h1a1 1 0 0 0 0-2h-1V2a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2m3.6 4a1 1 0 0 0-.78 1.18a9 9 0 1 1-7-7a1 1 0 1 0 .4-2A10.8 10.8 0 0 0 12 1a11 11 0 1 0 11 11a10.8 10.8 0 0 0-.22-2.2A1 1 0 0 0 21.6 9M7 7v10a1 1 0 0 0 2 0V7a1 1 0 0 0-2 0m4 2v6a3 3 0 0 0 3 3h1a3 3 0 0 0 3-3v-1a3 3 0 0 0-3-3h-2V9a1 1 0 0 1 1-1h1a1 1 0 0 1 1 1a1 1 0 0 0 2 0a3 3 0 0 0-3-3h-1a3 3 0 0 0-3 3m4 4a1 1 0 0 1 1 1v1a1 1 0 0 1-1 1h-1a1 1 0 0 1-1-1v-2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SkipForward(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 3a3 3 0 0 0-3 3v2.84L7 3.47a3.21 3.21 0 0 0-3.29 0A3.38 3.38 0 0 0 2 6.42v11.16a3.38 3.38 0 0 0 1.72 3a3.24 3.24 0 0 0 1.61.42A3.28 3.28 0 0 0 7 20.53l9-5.37V18a3 3 0 0 0 6 0V6a3 3 0 0 0-3-3m-3.68 10.23L6 18.81a1.23 1.23 0 0 1-1.28 0A1.4 1.4 0 0 1 4 17.58V6.42a1.4 1.4 0 0 1 .71-1.25A1.29 1.29 0 0 1 5.33 5a1.23 1.23 0 0 1 .67.19l9.33 5.58a1.45 1.45 0 0 1 0 2.46ZM20 18a1 1 0 0 1-2 0V6a1 1 0 0 1 2 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SkipForwardAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 3a3 3 0 0 0-3 3v12a3 3 0 0 0 6 0V6a3 3 0 0 0-3-3m1 15a1 1 0 0 1-2 0V6a1 1 0 0 1 2 0Zm14.68-8.35L14 5.66a2.6 2.6 0 0 0-2.64 0A2.74 2.74 0 0 0 10 8v8a2.74 2.74 0 0 0 1.37 2.38a2.57 2.57 0 0 0 2.64 0l6.67-4a2.75 2.75 0 0 0 0-4.7Zm-1 3l-6.66 4a.61.61 0 0 1-.63 0A.72.72 0 0 1 12 16V8a.72.72 0 0 1 .36-.64a.64.64 0 0 1 .31-.08a.63.63 0 0 1 .32.09l6.66 4a.76.76 0 0 1 0 1.28Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SkipForwardCircle(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.5 7a1 1 0 0 0-1 1v1.69l-4-2.31a2 2 0 0 0-3 1.73v5.78a2 2 0 0 0 1 1.73a2 2 0 0 0 2 0l4-2.31V16a1 1 0 0 0 2 0V8a1 1 0 0 0-1-1m-1 5l-5 2.89V9.11l5 2.88ZM12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Skype(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.435 14.156a9.586 9.586 0 0 0 .211-2.027a9.477 9.477 0 0 0-9.54-9.422a9.114 9.114 0 0 0-1.625.141A5.536 5.536 0 0 0 2 7.466a5.429 5.429 0 0 0 .753 2.756a10.02 10.02 0 0 0-.189 1.884a9.339 9.339 0 0 0 9.54 9.258a8.567 8.567 0 0 0 1.743-.166a5.58 5.58 0 0 0 2.616.802a5.433 5.433 0 0 0 4.97-7.844m-4.995 1.837a3.631 3.631 0 0 1-1.625 1.225a6.34 6.34 0 0 1-2.52.447a6.217 6.217 0 0 1-2.898-.612a3.733 3.733 0 0 1-1.32-1.178a2.574 2.574 0 0 1-.494-1.413a.88.88 0 0 1 .307-.684a1.09 1.09 0 0 1 .776-.282a.944.944 0 0 1 .637.212a1.793 1.793 0 0 1 .447.659a3.398 3.398 0 0 0 .495.873a1.79 1.79 0 0 0 .73.564a3.014 3.014 0 0 0 1.249.236a2.922 2.922 0 0 0 1.72-.447a1.332 1.332 0 0 0 .66-1.131a1.135 1.135 0 0 0-.354-.871a2.185 2.185 0 0 0-.92-.52c-.376-.117-.895-.235-1.53-.376a13.99 13.99 0 0 1-2.144-.636a3.348 3.348 0 0 1-1.366-1.013a2.474 2.474 0 0 1-.495-1.578a2.63 2.63 0 0 1 .542-1.602a3.412 3.412 0 0 1 1.53-1.084a6.652 6.652 0 0 1 2.38-.376a6.403 6.403 0 0 1 1.885.258a4.072 4.072 0 0 1 1.318.66a2.916 2.916 0 0 1 .778.872a1.803 1.803 0 0 1 .236.87a.962.962 0 0 1-.307.708a.991.991 0 0 1-.753.306a.974.974 0 0 1-.636-.189a2.382 2.382 0 0 1-.471-.611a2.937 2.937 0 0 0-.778-.967a2.376 2.376 0 0 0-1.46-.353a2.703 2.703 0 0 0-1.508.377a1.076 1.076 0 0 0-.565.896a.958.958 0 0 0 .188.565a1.419 1.419 0 0 0 .542.4a2.693 2.693 0 0 0 .683.26c.236.07.613.164 1.154.282c.66.142 1.273.306 1.815.471a5.43 5.43 0 0 1 1.389.636a2.857 2.857 0 0 1 .895.942a2.828 2.828 0 0 1 .33 1.39a2.89 2.89 0 0 1-.542 1.814"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SkypeAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.14 11.813a5.076 5.076 0 0 0-1.292-.593c-.28-.085-.59-.168-.91-.247c-.28-.078-.612-.158-1.022-.248a9.315 9.315 0 0 1-1.436-.424a1.496 1.496 0 0 1-.616-.447a.843.843 0 0 1-.16-.566a.967.967 0 0 1 .205-.597a1.598 1.598 0 0 1 .7-.475A4.012 4.012 0 0 1 12.031 8a3.787 3.787 0 0 1 1.106.146a2.083 2.083 0 0 1 .663.322a1.235 1.235 0 0 1 .325.343a1 1 0 1 0 1.761-.948a3.147 3.147 0 0 0-.837-.958a4.006 4.006 0 0 0-1.319-.669A5.768 5.768 0 0 0 12.032 6a5.963 5.963 0 0 0-2.145.35A3.552 3.552 0 0 0 8.31 7.492a2.977 2.977 0 0 0-.604 1.797a2.839 2.839 0 0 0 .58 1.792a3.5 3.5 0 0 0 1.438 1.072a10.582 10.582 0 0 0 1.307.408c.008.003.014.01.022.012c.192.058.498.135.94.23c.173.038.335.079.497.12c.016.004.039.009.054.014l.018.002c.248.064.487.129.706.196a3.023 3.023 0 0 1 .763.344a1.127 1.127 0 0 1 .363.368a1.201 1.201 0 0 1 .118.585a1.152 1.152 0 0 1-.214.732a1.763 1.763 0 0 1-.802.585a3.787 3.787 0 0 1-1.487.252a3.689 3.689 0 0 1-1.703-.344a1.756 1.756 0 0 1-.616-.547a1.016 1.016 0 0 1-.202-.503a1 1 0 0 0-2 0a2.94 2.94 0 0 0 .556 1.64a3.774 3.774 0 0 0 1.342 1.187a5.621 5.621 0 0 0 2.623.567a5.708 5.708 0 0 0 2.254-.405a3.71 3.71 0 0 0 1.665-1.273a3.146 3.146 0 0 0 .584-1.926a3.09 3.09 0 0 0-.375-1.53a3.165 3.165 0 0 0-.997-1.053m7.221 1.878A10.491 10.491 0 0 0 10.31 1.64a6.499 6.499 0 0 0-8.67 8.67a10.491 10.491 0 0 0 12.05 12.05a6.499 6.499 0 0 0 8.67-8.67M16.5 21a4.506 4.506 0 0 1-2.17-.558a1.004 1.004 0 0 0-.677-.106A8.492 8.492 0 0 1 3.5 12a8.583 8.583 0 0 1 .164-1.654a1.003 1.003 0 0 0-.106-.677A4.5 4.5 0 0 1 9.67 3.558a1 1 0 0 0 .678.106A8.492 8.492 0 0 1 20.5 12a8.583 8.583 0 0 1-.164 1.654a1.003 1.003 0 0 0 .106.677A4.499 4.499 0 0 1 16.5 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Slack(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 14.67a2 2 0 1 0 4 0v-2H4a2 2 0 0 0-2 2m12.64-3.34a2 2 0 0 0 2-2V4a2 2 0 1 0-4 0v5.33a2 2 0 0 0 2.02 2Zm7.32-2a2 2 0 1 0-4 0v2h2a2 2 0 0 0 2.04-2ZM9.34 12.67a2 2 0 0 0-2 2V20a2 2 0 1 0 4 0v-5.33a2 2 0 0 0-2-2M14.66 18h-2v2a2 2 0 1 0 2-2M20 12.67h-5.34a2 2 0 0 0 0 4H20a2 2 0 0 0 0-4M9.34 7.33H4a2 2 0 1 0 0 4h5.34a2 2 0 0 0 0-4m0-5.33a2 2 0 0 0 0 4h2V4a2 2 0 0 0-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SlackAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23 9.5A3.504 3.504 0 0 0 19.5 6a3.46 3.46 0 0 0-1.5.351V4.5a3.495 3.495 0 0 0-6-2.442A3.487 3.487 0 0 0 6.351 6H4.5a3.495 3.495 0 0 0-2.442 6A3.487 3.487 0 0 0 6 17.649V19.5a3.495 3.495 0 0 0 6 2.442A3.487 3.487 0 0 0 17.649 18H19.5a3.495 3.495 0 0 0 2.442-6A3.486 3.486 0 0 0 23 9.5m-10-5a1.5 1.5 0 0 1 3 0v5a1.5 1.5 0 0 1-3 0Zm-7 10A1.5 1.5 0 1 1 4.5 13H6Zm5 5a1.5 1.5 0 0 1-3 0v-5a1.5 1.5 0 0 1 3 0ZM9.5 11h-5a1.5 1.5 0 0 1 0-3h5a1.5 1.5 0 0 1 0 3M11 6H9.5A1.5 1.5 0 1 1 11 4.5Zm1 6.058a3.725 3.725 0 0 0-.058-.058l.058-.058l.058.058zM14.5 21a1.502 1.502 0 0 1-1.5-1.5V18h1.5a1.5 1.5 0 0 1 0 3m5-5h-5a1.5 1.5 0 0 1 0-3h5a1.5 1.5 0 0 1 0 3m0-5H18V9.5a1.5 1.5 0 1 1 1.5 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SliderH(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 11h-3.184a2.982 2.982 0 0 0-5.632 0H3a1 1 0 0 0 0 2h9.184a2.982 2.982 0 0 0 5.632 0H21a1 1 0 0 0 0-2m-6 2a1 1 0 1 1 1-1a1.001 1.001 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SliderHrange(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 11h-1.184a2.982 2.982 0 0 0-5.632 0H9.816a2.982 2.982 0 0 0-5.632 0H3a1 1 0 0 0 0 2h1.184a2.982 2.982 0 0 0 5.632 0h4.368a2.982 2.982 0 0 0 5.632 0H21a1 1 0 0 0 0-2M7 13a1 1 0 1 1 1-1a1.001 1.001 0 0 1-1 1m10 0a1 1 0 1 1 1-1a1.001 1.001 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SlidersV(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 6H6V3a1 1 0 0 0-2 0v3H3a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2m-2 4a1 1 0 0 0-1 1v10a1 1 0 0 0 2 0V11a1 1 0 0 0-1-1m7 8a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0v-2a1 1 0 0 0-1-1m9-8h-1V3a1 1 0 0 0-2 0v7h-1a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2m-2 4a1 1 0 0 0-1 1v6a1 1 0 0 0 2 0v-6a1 1 0 0 0-1-1m-5 0h-1V3a1 1 0 0 0-2 0v11h-1a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SlidersValt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 8.18V3a1 1 0 0 0-2 0v5.18a3 3 0 0 0 0 5.64V21a1 1 0 0 0 2 0v-7.18a3 3 0 0 0 0-5.64M19 12a1 1 0 1 1 1-1a1 1 0 0 1-1 1m-6 2.18V3a1 1 0 0 0-2 0v11.18a3 3 0 0 0 0 5.64V21a1 1 0 0 0 2 0v-1.18a3 3 0 0 0 0-5.64M12 18a1 1 0 1 1 1-1a1 1 0 0 1-1 1M6 6.18V3a1 1 0 0 0-2 0v3.18a3 3 0 0 0 0 5.64V21a1 1 0 0 0 2 0v-9.18a3 3 0 0 0 0-5.64M5 10a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Smile(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.36 14.23a3.76 3.76 0 0 1-4.72 0a1 1 0 0 0-1.28 1.54a5.68 5.68 0 0 0 7.28 0a1 1 0 1 0-1.28-1.54M9 11a1 1 0 1 0-1-1a1 1 0 0 0 1 1m6-2a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-3-7a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmileBeam(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.36 14.23a3.76 3.76 0 0 1-4.72 0a1 1 0 0 0-1.28 1.54a5.68 5.68 0 0 0 7.28 0a1 1 0 1 0-1.28-1.54m-5.15-3.69a1 1 0 0 0 1.41 0a1 1 0 0 0 0-1.41a3.08 3.08 0 0 0-4.24 0a1 1 0 1 0 1.41 1.41a1 1 0 0 1 1.42 0m8.41-1.41a3.08 3.08 0 0 0-4.24 0a1 1 0 0 0 1.41 1.41a1 1 0 0 1 1.42 0a1 1 0 0 0 1.41 0a1 1 0 0 0 0-1.41M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmileDizzy(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.36 14.23a3.76 3.76 0 0 1-4.72 0a1 1 0 0 0-1.28 1.54a5.68 5.68 0 0 0 7.28 0a1 1 0 1 0-1.28-1.54M9 11.71l.29-.3l.29.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-.3-.29l.3-.29a1 1 0 0 0-1.46-1.42l-.29.3l-.25-.3a1 1 0 1 0-1.46 1.42l.3.29l-.3.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0ZM12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8m5-11.71a1 1 0 0 0-1.42 0l-.29.3l-.29-.3a1 1 0 0 0-1.42 1.42l.3.29l-.3.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l.29-.3l.29.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-.3-.29l.3-.29a1 1 0 0 0 0-1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmileSquintWink(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.42 12.21a1 1 0 0 0 .71.29a1 1 0 0 0 .71-.29l1.5-1.5a1 1 0 0 0 0-1.42l-1.5-1.5a1 1 0 1 0-1.42 1.42l.79.79l-.79.79a1 1 0 0 0 0 1.42m5.94 2a3.76 3.76 0 0 1-4.72 0a1 1 0 0 0-1.28 1.54a5.68 5.68 0 0 0 7.28 0a1 1 0 1 0-1.28-1.54M15 9a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-3-7a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmileSquintWinkAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 11a1 1 0 1 0-1-1a1 1 0 0 0 1 1m5.16 1.21a1 1 0 0 0 .71.29a1 1 0 0 0 .71-.29a1 1 0 0 0 0-1.42l-.79-.79l.79-.79a1 1 0 1 0-1.42-1.42l-1.5 1.5a1 1 0 0 0 0 1.42ZM12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8m2.36-5.77a3.76 3.76 0 0 1-4.72 0a1 1 0 0 0-1.28 1.54a5.68 5.68 0 0 0 7.28 0a1 1 0 1 0-1.28-1.54"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmileWink(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.36 14.23a3.76 3.76 0 0 1-4.72 0a1 1 0 0 0-1.28 1.54a5.68 5.68 0 0 0 7.28 0a1 1 0 1 0-1.28-1.54M10.5 10A1.5 1.5 0 1 0 9 11.5a1.5 1.5 0 0 0 1.5-1.5M15 9h-1a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2m-3-7a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmileWinkAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 11a1 1 0 1 0-1-1a1 1 0 0 0 1 1m5.36 3.23a3.76 3.76 0 0 1-4.72 0a1 1 0 0 0-1.28 1.54a5.68 5.68 0 0 0 7.28 0a1 1 0 1 0-1.28-1.54M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8m5.62-10.87a3.08 3.08 0 0 0-4.24 0a1 1 0 0 0 1.41 1.41a1 1 0 0 1 1.42 0a1 1 0 0 0 1.41 0a1 1 0 0 0 0-1.41"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SnapchatAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.951 15.614a4.724 4.724 0 0 1-2.981-2.173a1 1 0 1 0-1.657 1.121a7.688 7.688 0 0 0 2.403 2.335c-.135.025-.281.05-.442.075a1.367 1.367 0 0 0-1.076 1.207a6.062 6.062 0 0 0-2.014-.004a4.64 4.64 0 0 0-1.958.956a3.484 3.484 0 0 1-2.104.87h-.26a3.485 3.485 0 0 1-2.106-.872a4.627 4.627 0 0 0-1.929-.95a6.39 6.39 0 0 0-2.04.005a1.368 1.368 0 0 0-1.062-1.21a12.74 12.74 0 0 1-.435-.075A6.857 6.857 0 0 0 6.085 15.4a6.714 6.714 0 0 0 .635-.868a1 1 0 0 0-1.696-1.06a4.907 4.907 0 0 1-.448.616a4.252 4.252 0 0 1-2.553 1.528a1.224 1.224 0 0 0-1.032 1.236a1.28 1.28 0 0 0 .115.533c.316.716 1.156 1.168 2.785 1.474l.03.13c.03.118.058.239.093.348a1.289 1.289 0 0 0 1.278.945a2.59 2.59 0 0 0 .603-.087a4.67 4.67 0 0 1 1.588-.046a3.008 3.008 0 0 1 1.123.618a5.413 5.413 0 0 0 3.255 1.235h.261a5.445 5.445 0 0 0 3.274-1.248a2.915 2.915 0 0 1 1.131-.61a3.485 3.485 0 0 1 .624-.052a4.27 4.27 0 0 1 .955.106a3.788 3.788 0 0 0 .616.064a1.256 1.256 0 0 0 1.245-.923c.038-.122.067-.24.094-.355l.032-.124c1.624-.305 2.467-.754 2.767-1.44a1.17 1.17 0 0 0 .127-.48a1.244 1.244 0 0 0-1.036-1.326m-18.15-4.466a1.005 1.005 0 0 0 1.057-.484a3.138 3.138 0 0 0 1.275.338a1.794 1.794 0 0 0 1.265-.499a1 1 0 0 0 .317-.79l-.036-.602a9.902 9.902 0 0 1 .156-3.561a4.26 4.26 0 0 1 3.966-2.544l.388-.004a4.264 4.264 0 0 1 3.96 2.547a9.917 9.917 0 0 1 .156 3.564l-.01.163l-.027.444a1.028 1.028 0 0 0 .312.778a1.795 1.795 0 0 0 1.254.503a3.271 3.271 0 0 0 1.241-.365a1 1 0 0 0 .916.598h.018a1 1 0 0 0 .982-1.017a1.642 1.642 0 0 0-1.185-1.451a1.914 1.914 0 0 0-1.477.01a9.471 9.471 0 0 0-.354-4.042a6.236 6.236 0 0 0-5.796-3.732l-.393.004a6.222 6.222 0 0 0-5.777 3.73a9.586 9.586 0 0 0-.352 4.094l-.111-.052a1.909 1.909 0 0 0-2.537 1.199a1.002 1.002 0 0 0 .793 1.171"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SnapchatGhost(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.798 16.987c-2.867-.472-4.151-3.401-4.204-3.526l-.006-.011a1.07 1.07 0 0 1-.102-.898c.192-.454.83-.656 1.251-.79c.106-.033.205-.065.283-.096c.763-.3.918-.613.914-.822a.662.662 0 0 0-.5-.543l-.007-.002a.946.946 0 0 0-.356-.069a.755.755 0 0 0-.313.063a2.54 2.54 0 0 1-.955.266a.821.821 0 0 1-.53-.178l.032-.53l.004-.065a10.102 10.102 0 0 0-.24-4.035a5.248 5.248 0 0 0-4.874-3.14l-.402.005a5.24 5.24 0 0 0-4.864 3.137a10.09 10.09 0 0 0-.242 4.031q.02.299.036.598a.848.848 0 0 1-.584.178a2.453 2.453 0 0 1-1.014-.268a.575.575 0 0 0-.245-.049a.834.834 0 0 0-.81.533c-.082.43.532.743.906.89c.08.032.178.063.283.096c.422.134 1.06.336 1.252.79a1.072 1.072 0 0 1-.102.898l-.006.011a7.028 7.028 0 0 1-1.069 1.663A5.215 5.215 0 0 1 2.2 16.987a.24.24 0 0 0-.2.25a.38.38 0 0 0 .03.13c.177.411 1.059.75 2.553.981c.14.022.198.25.28.622c.032.15.066.305.113.465a.293.293 0 0 0 .32.228a2.485 2.485 0 0 0 .424-.06a5.53 5.53 0 0 1 1.12-.127a4.954 4.954 0 0 1 .809.068a3.877 3.877 0 0 1 1.535.784a4.443 4.443 0 0 0 2.69 1.06c.033 0 .067-.001.1-.004c.04.002.095.004.151.004a4.448 4.448 0 0 0 2.692-1.06a3.873 3.873 0 0 1 1.533-.784a4.973 4.973 0 0 1 .808-.068a5.593 5.593 0 0 1 1.12.119a2.391 2.391 0 0 0 .425.053h.024a.279.279 0 0 0 .295-.22a6.52 6.52 0 0 0 .114-.462c.081-.371.14-.598.28-.62c1.494-.23 2.377-.57 2.551-.978a.385.385 0 0 0 .032-.13a.24.24 0 0 0-.2-.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SnapchatSquare(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.973 6.68a6.124 6.124 0 0 0-.098-1.073a4.372 4.372 0 0 0-.406-1.246a4.324 4.324 0 0 0-.832-1.11a4.125 4.125 0 0 0-1.816-1.036a7.36 7.36 0 0 0-1.92-.205L16.898 2H7.099v.01a10.488 10.488 0 0 0-1.101.05a5.243 5.243 0 0 0-1.176.264A4.262 4.262 0 0 0 2.219 5.17a7.338 7.338 0 0 0-.205 1.905l-.006 9.838a9.445 9.445 0 0 0 .09 1.333a4.616 4.616 0 0 0 .41 1.345a4.305 4.305 0 0 0 1.201 1.454a3.903 3.903 0 0 0 1.203.651a6.516 6.516 0 0 0 1.976.29c.42.003.839.014 1.258.012c3.047-.013 6.094.022 9.14-.019a7.19 7.19 0 0 0 1.2-.127a4.06 4.06 0 0 0 2.007-.977a4.162 4.162 0 0 0 1.326-2.212a8.062 8.062 0 0 0 .173-1.75v-.117c0-.046-.017-9.984-.019-10.115m-2.676 9.25c-.128.3-.774.548-1.868.717c-.102.016-.146.182-.205.454c-.024.112-.05.222-.083.337a.204.204 0 0 1-.216.162h-.018a1.746 1.746 0 0 1-.31-.04a4.097 4.097 0 0 0-.821-.086a3.637 3.637 0 0 0-.592.05a2.836 2.836 0 0 0-1.123.574a3.257 3.257 0 0 1-1.97.776a2.33 2.33 0 0 1-.112-.003a.854.854 0 0 1-.073.003a3.253 3.253 0 0 1-1.97-.776a2.84 2.84 0 0 0-1.123-.574a3.633 3.633 0 0 0-.592-.05a4.047 4.047 0 0 0-.82.093a1.82 1.82 0 0 1-.311.044a.214.214 0 0 1-.234-.167a4.084 4.084 0 0 1-.083-.34c-.06-.273-.103-.44-.205-.456c-1.094-.168-1.74-.417-1.869-.718a.278.278 0 0 1-.023-.095a.176.176 0 0 1 .147-.183a3.818 3.818 0 0 0 2.296-1.365A5.145 5.145 0 0 0 7.9 13.07l.004-.008a.785.785 0 0 0 .075-.658c-.14-.332-.607-.48-.916-.578a3.27 3.27 0 0 1-.207-.07c-.274-.108-.724-.337-.664-.652a.61.61 0 0 1 .593-.39a.42.42 0 0 1 .18.036a1.796 1.796 0 0 0 .742.196a.621.621 0 0 0 .428-.131q-.012-.219-.027-.438a7.388 7.388 0 0 1 .177-2.951a3.837 3.837 0 0 1 3.562-2.298l.295-.002a3.843 3.843 0 0 1 3.568 2.298a7.398 7.398 0 0 1 .176 2.955l-.003.047l-.023.389a.602.602 0 0 0 .388.13a1.86 1.86 0 0 0 .7-.195a.552.552 0 0 1 .228-.046a.694.694 0 0 1 .261.05l.004.002a.485.485 0 0 1 .367.398c.003.153-.11.381-.669.602c-.057.022-.13.046-.207.07c-.31.098-.776.246-.916.578a.784.784 0 0 0 .074.657l.004.009a4.522 4.522 0 0 0 3.079 2.582a.176.176 0 0 1 .146.183a.28.28 0 0 1-.023.096"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SnowFlake(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 11h-1.59l1.3-1.29a1 1 0 0 0-1.42-1.42L16.59 11h-2.18l2.3-2.29a1 1 0 1 0-1.42-1.42L13 9.59V7.41l2.71-2.7a1 1 0 1 0-1.42-1.42L13 4.59V3a1 1 0 0 0-2 0v1.59l-1.29-1.3a1 1 0 0 0-1.42 1.42L11 7.41v2.18l-2.29-2.3a1 1 0 1 0-1.42 1.42L9.59 11H7.41l-2.7-2.71a1 1 0 0 0-1.42 1.42L4.59 11H3a1 1 0 0 0 0 2h1.59l-1.3 1.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0L7.41 13h2.18l-2.3 2.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l2.29-2.3v2.18l-2.71 2.7a1 1 0 0 0 1.42 1.42l1.29-1.3V21a1 1 0 0 0 2 0v-1.59l1.29 1.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L13 16.59v-2.18l2.29 2.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L14.41 13h2.18l2.7 2.71a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L19.41 13H21a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Snowflake(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.16 16.13l-2-1.15l.89-.24a1 1 0 1 0-.52-1.93l-2.82.76L14 12l2.71-1.57l2.82.76h.26a1 1 0 0 0 .26-2L19.16 9l2-1.15a1 1 0 0 0-1-1.74L18 7.37l.3-1.11a1 1 0 1 0-1.93-.52l-.82 3L13 10.27V7.14l2.07-2.07a1 1 0 0 0 0-1.41a1 1 0 0 0-1.42 0l-.65.65V2a1 1 0 0 0-2 0v2.47l-.81-.81a1 1 0 0 0-1.42 0a1 1 0 0 0 0 1.41L11 7.3v3L8.43 8.78l-.82-3a1 1 0 1 0-1.93.52L6 7.37L3.84 6.13a1 1 0 0 0-1 1.74l2 1.13l-.84.26a1 1 0 0 0 .26 2h.26l2.82-.76L10 12l-2.71 1.57l-2.82-.76A1 1 0 1 0 4 14.74l.89.24l-2 1.15a1 1 0 0 0 1 1.74L6 16.63l-.3 1.11A1 1 0 0 0 6.39 19a1.15 1.15 0 0 0 .26 0a1 1 0 0 0 1-.74l.82-3L11 13.73v3.13l-2.07 2.07a1 1 0 0 0 0 1.41a1 1 0 0 0 .71.3a1 1 0 0 0 .71-.3l.65-.65V22a1 1 0 0 0 2 0v-2.47l.81.81a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.41L13 16.7v-3l2.57 1.49l.82 3a1 1 0 0 0 1 .74a1.15 1.15 0 0 0 .26 0a1 1 0 0 0 .71-1.23L18 16.63l2.14 1.24a1 1 0 1 0 1-1.74Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SnowflakeAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.93 17.66a1 1 0 1 0 1.41 0a1 1 0 0 0-1.41 0M19.07 6.34a1 1 0 1 0-1.41 0a1 1 0 0 0 1.41 0m-12.73 0a1 1 0 1 0-1.41 0a1 1 0 0 0 1.41 0m11.32 11.32a1 1 0 1 0 1.41 0a1 1 0 0 0-1.41 0M21 11h-1.59l1.3-1.29a1 1 0 0 0-1.42-1.42L16.59 11h-2.18l2.3-2.29a1 1 0 1 0-1.42-1.42L13 9.59V7.41l2.71-2.7a1 1 0 1 0-1.42-1.42L13 4.59V3a1 1 0 0 0-2 0v1.59l-1.29-1.3a1 1 0 0 0-1.42 1.42L11 7.41v2.18l-2.29-2.3a1 1 0 1 0-1.42 1.42L9.59 11H7.41l-2.7-2.71a1 1 0 0 0-1.42 1.42L4.59 11H3a1 1 0 0 0 0 2h1.59l-1.3 1.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0L7.41 13h2.18l-2.3 2.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l2.29-2.3v2.18l-2.71 2.7a1 1 0 0 0 1.42 1.42l1.29-1.3V21a1 1 0 0 0 2 0v-1.59l1.29 1.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L13 16.59v-2.18l2.29 2.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L14.41 13h2.18l2.7 2.71a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L19.41 13H21a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SocialDistancing(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 17H5.914l.293-.293a1 1 0 0 0-1.414-1.414l-2 2a1.004 1.004 0 0 0 0 1.414l2 2a1 1 0 0 0 1.414-1.414L5.914 19H8.5a1 1 0 0 0 0-2m12.707.293l-2-2a1 1 0 0 0-1.414 1.414l.293.293H15.5a1 1 0 0 0 0 2h2.586l-.293.293a1 1 0 1 0 1.414 1.414l2-2a1.004 1.004 0 0 0 0-1.414m-.567-7.521A3.468 3.468 0 0 0 21.5 7.5a3.5 3.5 0 0 0-7 0a3.468 3.468 0 0 0 .86 2.272A4.988 4.988 0 0 0 13 14a1 1 0 0 0 2 0a3 3 0 0 1 6 0a1 1 0 0 0 2 0a4.988 4.988 0 0 0-2.36-4.228M18 9a1.5 1.5 0 1 1 1.5-1.5A1.502 1.502 0 0 1 18 9m-9.36.772A3.468 3.468 0 0 0 9.5 7.5a3.5 3.5 0 0 0-7 0a3.468 3.468 0 0 0 .86 2.272A4.988 4.988 0 0 0 1 14a1 1 0 0 0 2 0a3 3 0 0 1 6 0a1 1 0 0 0 2 0a4.988 4.988 0 0 0-2.36-4.228M6 9a1.5 1.5 0 1 1 1.5-1.5A1.502 1.502 0 0 1 6 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sort(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.29 14.29L12 18.59l-4.29-4.3a1 1 0 0 0-1.42 1.42l5 5a1 1 0 0 0 1.42 0l5-5a1 1 0 0 0-1.42-1.42M7.71 9.71L12 5.41l4.29 4.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-5-5a1 1 0 0 0-1.42 0l-5 5a1 1 0 0 0 1.42 1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortAmountDown(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m6.29 14.29l-.29.3V7a1 1 0 0 0-2 0v7.59l-.29-.3a1 1 0 0 0-1.42 1.42l2 2a1 1 0 0 0 .33.21a.94.94 0 0 0 .76 0a1 1 0 0 0 .33-.21l2-2a1 1 0 0 0-1.42-1.42M11 8h10a1 1 0 0 0 0-2H11a1 1 0 0 0 0 2m10 3H11a1 1 0 0 0 0 2h10a1 1 0 0 0 0-2m0 5H11a1 1 0 0 0 0 2h10a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SortAmountUp(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.71 6.29a1 1 0 0 0-.33-.21a1 1 0 0 0-.76 0a1 1 0 0 0-.33.21l-2 2a1 1 0 0 0 1.42 1.42l.29-.3V17a1 1 0 0 0 2 0V9.41l.29.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM11 8h10a1 1 0 0 0 0-2H11a1 1 0 0 0 0 2m10 8H11a1 1 0 0 0 0 2h10a1 1 0 0 0 0-2m0-5H11a1 1 0 0 0 0 2h10a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sorting(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.29 10.21a1 1 0 0 0 1.42 0l3-3a1 1 0 1 0-1.42-1.42L12 8.09l-2.29-2.3a1 1 0 0 0-1.42 1.42Zm1.42 4.58a1 1 0 0 0-1.42 0l-3 3a1 1 0 0 0 1.42 1.42l2.29-2.3l2.29 2.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpaceKey(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 9a1 1 0 0 0-1 1v3H4v-3a1 1 0 0 0-2 0v4a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spade(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.06 12.28a5.69 5.69 0 0 0-1.67-4L14 2.81a2.83 2.83 0 0 0-4 0L4.61 8.23a5.69 5.69 0 0 0-1.67 4A5.7 5.7 0 0 0 8.66 18a6.88 6.88 0 0 1-1.23 2.37A1 1 0 0 0 8.24 22h7.52a1 1 0 0 0 .78-1.63A6.84 6.84 0 0 1 15.31 18a5.69 5.69 0 0 0 5.75-5.71ZM18 14.91a3.78 3.78 0 0 1-3.66.95a1 1 0 0 0-.17 0h-.2a1.41 1.41 0 0 0-.22.06h-.15a.69.69 0 0 0-.13.11a.75.75 0 0 0-.17.14a.6.6 0 0 1-.06.11a2.53 2.53 0 0 0-.12.23a1.1 1.1 0 0 0 0 .18v.18a8.84 8.84 0 0 0 .82 3.13h-3.88a8.62 8.62 0 0 0 .88-3v-.19a1.1 1.1 0 0 0 0-.18a1.12 1.12 0 0 0-.13-.24a.53.53 0 0 0-.06-.1a.54.54 0 0 0-.16-.14a1.27 1.27 0 0 0-.13-.11h-.14a.88.88 0 0 0-.23-.07H9.9a1.1 1.1 0 0 0-.18 0A3.81 3.81 0 0 1 6 14.91a3.75 3.75 0 0 1-1.09-2.63A3.69 3.69 0 0 1 6 9.65l5.4-5.42a.81.81 0 0 1 1.13 0L18 9.65a3.69 3.69 0 0 1 1.09 2.63A3.78 3.78 0 0 1 18 14.91"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sperms(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.06 20.1a2 2 0 0 1-.65-.48a2.47 2.47 0 0 1-.67-1.14a4.19 4.19 0 0 0-1.31-2.06a3.57 3.57 0 0 0-1-3.28a3.28 3.28 0 0 0-4.59-.35a3.29 3.29 0 0 0 .35 4.6a3.87 3.87 0 0 0 2.21 1.12h.36a2.85 2.85 0 0 0 1.53-.44a2.05 2.05 0 0 1 .51.93A4.46 4.46 0 0 0 9 21a4.27 4.27 0 0 0 1.2.88a1 1 0 0 0 .44.1a1 1 0 0 0 .44-1.9Zm-4.74-3.77a.86.86 0 0 1-.7.19A1.8 1.8 0 0 1 4.56 16c-.57-.56-.73-1.39-.36-1.77a.82.82 0 0 1 .56-.2a1.8 1.8 0 0 1 1.24.53a1.8 1.8 0 0 1 .55 1.06a.86.86 0 0 1-.23.71m15-6.12a2.17 2.17 0 0 1-.91-.59a2.47 2.47 0 0 1-.67-1.14a4.19 4.19 0 0 0-1.31-2.06a3.57 3.57 0 0 0-1.05-3.28a3.28 3.28 0 0 0-4.59-.35a3.29 3.29 0 0 0 .35 4.6a3.87 3.87 0 0 0 2.26 1.12h.36a2.85 2.85 0 0 0 1.53-.44a2.05 2.05 0 0 1 .51.93a4.54 4.54 0 0 0 1.2 2a4.36 4.36 0 0 0 1.7 1.08a1.25 1.25 0 0 0 .32.05a1 1 0 0 0 .95-.68a1 1 0 0 0-.65-1.24m-5-3.88a.86.86 0 0 1-.7.19A1.8 1.8 0 0 1 14.56 6c-.57-.56-.73-1.39-.36-1.77a.82.82 0 0 1 .56-.2a1.8 1.8 0 0 1 1.24.53a1.8 1.8 0 0 1 .55 1.06a.86.86 0 0 1-.23.71m5.39 14.12a4.45 4.45 0 0 0-3.23-3.23a3.18 3.18 0 0 1-1.39-.82a2.93 2.93 0 0 1-.8-1.38A4.65 4.65 0 0 0 13 11.77a2.92 2.92 0 0 1-1.38-.8a3 3 0 0 1-.81-1.39a4.68 4.68 0 0 0-1-1.86a3.94 3.94 0 0 0 .19-1.5a4.75 4.75 0 0 0-1.4-2.8A4.78 4.78 0 0 0 5.78 2A3.43 3.43 0 0 0 3 3a3.43 3.43 0 0 0-1 2.78a4.75 4.75 0 0 0 1.4 2.8A4.78 4.78 0 0 0 6.22 10h.43a3.39 3.39 0 0 0 2-.6a2.64 2.64 0 0 1 .31.71a5.12 5.12 0 0 0 3.6 3.59a2.61 2.61 0 0 1 1.83 1.84a5.11 5.11 0 0 0 3.6 3.6a2.57 2.57 0 0 1 1.8 1.86a2.8 2.8 0 0 0 .16.46a1 1 0 0 0 .88.54a1.21 1.21 0 0 0 .44-.1a1 1 0 0 0 .48-1.33s-.03-.1-.04-.12M7.64 7.64c-.64.63-1.92.41-2.81-.47A2.78 2.78 0 0 1 4 5.56a1.47 1.47 0 0 1 .35-1.2a1.39 1.39 0 0 1 1-.36a2.71 2.71 0 0 1 1.83.83A2.78 2.78 0 0 1 8 6.44a1.47 1.47 0 0 1-.36 1.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spin(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 3a7 7 0 0 0 0 14a5 5 0 0 0 0-10a3 3 0 0 0 0 6a1 1 0 0 0 0-2a1 1 0 0 1 0-2a3 3 0 0 1 0 6a5 5 0 0 1 0-10a7 7 0 0 1 0 14a9 9 0 0 1-9-9a1 1 0 0 0-2 0a11 11 0 0 0 11 11a9 9 0 0 0 0-18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spinner(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.1 16c-.3-.5-.9-.6-1.4-.4c-.5.3-.6.9-.4 1.4c.3.5.9.6 1.4.4c.5-.3.6-.9.4-1.4m-.4-9.4c-.5-.2-1.1-.1-1.4.4c-.2.5-.1 1.1.4 1.4c.5.2 1.1.1 1.4-.4c.2-.5.1-1.1-.4-1.4m15.6 1.8c.5-.3.6-.9.4-1.4c-.3-.5-.9-.6-1.4-.4c-.5.3-.6.9-.4 1.4c.3.5.9.6 1.4.4M4 12c0-.6-.4-1-1-1s-1 .4-1 1s.4 1 1 1s1-.4 1-1m3.2 6.8c-.5.1-.9.7-.7 1.2c.1.5.7.9 1.2.7c.5-.1.9-.7.7-1.2c-.1-.5-.6-.8-1.2-.7M21 11c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m-.7 4.6c-.5-.3-1.1-.1-1.4.4c-.3.5-.1 1.1.4 1.4c.5.3 1.1.1 1.4-.4c.2-.5.1-1.1-.4-1.4M17 3.3c-.5-.3-1.1-.1-1.4.4c-.3.5-.1 1.1.4 1.4c.5.3 1.1.1 1.4-.4c.2-.5.1-1.1-.4-1.4m-.2 15.5c-.5-.1-1.1.2-1.2.7c-.1.5.2 1.1.7 1.2c.5.1 1.1-.2 1.2-.7c.1-.5-.2-1-.7-1.2M12 20c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1m0-18c-.6 0-1 .4-1 1s.4 1 1 1s1-.4 1-1s-.4-1-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpinnerAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.804 15a1 1 0 0 0-1.366-.366l-1.732 1a1 1 0 0 0 1 1.732l1.732-1A1 1 0 0 0 6.804 15M3.706 8.366l1.732 1a1 1 0 1 0 1-1.732l-1.732-1a1 1 0 0 0-1 1.732M6 12a1 1 0 0 0-1-1H3a1 1 0 0 0 0 2h2a1 1 0 0 0 1-1m11.196-3a1 1 0 0 0 1.366.366l1.732-1a1 1 0 1 0-1-1.732l-1.732 1A1 1 0 0 0 17.196 9M15 6.804a1 1 0 0 0 1.366-.366l1-1.732a1 1 0 1 0-1.732-1l-1 1.732A1 1 0 0 0 15 6.804m5.294 8.83l-1.732-1a1 1 0 1 0-1 1.732l1.732 1a1 1 0 0 0 1-1.732m-3.928 1.928a1 1 0 1 0-1.732 1l1 1.732a1 1 0 1 0 1.732-1ZM21 11h-2a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2m-9 7a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0v-2a1 1 0 0 0-1-1m-3-.804a1 1 0 0 0-1.366.366l-1 1.732a1 1 0 0 0 1.732 1l1-1.732A1 1 0 0 0 9 17.196M12 2a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0V3a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Square(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1m-1 18H4V4h16Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareFull(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1m-1 18H4V4h16Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareShape(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2H7a5 5 0 0 0-5 5v10a5 5 0 0 0 5 5h10a5 5 0 0 0 5-5V7a5 5 0 0 0-5-5m3 15a3 3 0 0 1-3 3H7a3 3 0 0 1-3-3V7a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Squint(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 14H9a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2m-5.92-1.79l1.5-1.5a1 1 0 0 0 0-1.42l-1.5-1.5a1 1 0 0 0-1.42 1.42l.8.79l-.8.79a1 1 0 0 0 0 1.42a1 1 0 0 0 .71.29a1 1 0 0 0 .71-.29M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8m4.58-12.21a1 1 0 0 0-1.42 0l-1.5 1.5a1 1 0 0 0 0 1.42l1.5 1.5a1 1 0 0 0 .71.29a1 1 0 0 0 .71-.29a1 1 0 0 0 0-1.42l-.79-.79l.79-.79a1 1 0 0 0 0-1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Star(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 9.67a1 1 0 0 0-.86-.67l-5.69-.83L12.9 3a1 1 0 0 0-1.8 0L8.55 8.16L2.86 9a1 1 0 0 0-.81.68a1 1 0 0 0 .25 1l4.13 4l-1 5.68a1 1 0 0 0 .4 1a1 1 0 0 0 1.05.07L12 18.76l5.1 2.68a.93.93 0 0 0 .46.12a1 1 0 0 0 .59-.19a1 1 0 0 0 .4-1l-1-5.68l4.13-4A1 1 0 0 0 22 9.67m-6.15 4a1 1 0 0 0-.29.89l.72 4.19l-3.76-2a1 1 0 0 0-.94 0l-3.76 2l.72-4.19a1 1 0 0 0-.29-.89l-3-3l4.21-.61a1 1 0 0 0 .76-.55L12 5.7l1.88 3.82a1 1 0 0 0 .76.55l4.21.61Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarHalfAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 9.67a1 1 0 0 0-.86-.67l-5.69-.83L12.9 3a1 1 0 0 0-1.8 0L8.55 8.16L2.86 9a1 1 0 0 0-.81.68a1 1 0 0 0 .25 1l4.13 4l-1 5.68a1 1 0 0 0 1.47 1.08l5.1-2.67l5.1 2.67a.93.93 0 0 0 .46.12a1 1 0 0 0 .59-.19a1 1 0 0 0 .4-1l-1-5.68l4.13-4A1 1 0 0 0 22 9.67M11 17l-3.23 1.7l.72-4.2a1 1 0 0 0-.29-.88l-3-3l4.21-.61a1 1 0 0 0 .76-.55L11 7.73Zm4.8-3.38a1 1 0 0 0-.29.88l.72 4.2L13 17V7.73l.88 1.79a1 1 0 0 0 .76.55l4.21.61Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StepBackward(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.28 3.43a3.23 3.23 0 0 0-3.29 0L8 8.84V6a3 3 0 0 0-6 0v12a3 3 0 0 0 6 0v-2.84l9 5.37a3.28 3.28 0 0 0 1.68.47a3.24 3.24 0 0 0 1.61-.43a3.38 3.38 0 0 0 1.72-3V6.42a3.38 3.38 0 0 0-1.73-2.99M6 18a1 1 0 0 1-2 0V6a1 1 0 0 1 2 0Zm14-.42a1.4 1.4 0 0 1-.71 1.25a1.23 1.23 0 0 1-1.28 0l-9.33-5.6a1.45 1.45 0 0 1 0-2.46L18 5.19a1.23 1.23 0 0 1 .67-.19a1.29 1.29 0 0 1 .62.17A1.4 1.4 0 0 1 20 6.42Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StepBackwardAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 3a3 3 0 0 0-3 3v12a3 3 0 0 0 6 0V6a3 3 0 0 0-3-3m1 15a1 1 0 0 1-2 0V6a1 1 0 0 1 2 0ZM12.63 5.63a2.6 2.6 0 0 0-2.64 0l-6.67 4a2.75 2.75 0 0 0 0 4.7l6.67 4a2.57 2.57 0 0 0 2.64 0A2.74 2.74 0 0 0 14 16V8a2.74 2.74 0 0 0-1.37-2.37M12 16a.72.72 0 0 1-.36.64a.61.61 0 0 1-.63 0l-6.66-4a.76.76 0 0 1 0-1.28l6.66-4a.63.63 0 0 1 .32-.09a.64.64 0 0 1 .31.08A.72.72 0 0 1 12 8Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StepBackwardCircle(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.5 7.38a2 2 0 0 0-2 0l-4 2.31V8a1 1 0 0 0-2 0v8a1 1 0 0 0 2 0v-1.69l4 2.31a2 2 0 0 0 2 0a2 2 0 0 0 1-1.73V9.11a2 2 0 0 0-1-1.73m-1 7.51L9.5 12l5-2.89ZM12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StepForward(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.71 7.29a1 1 0 1 0-1.42 1.42l3.3 3.29l-3.3 3.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l4-4a1 1 0 0 0 0-1.42ZM16 7a1 1 0 0 0-1 1v8a1 1 0 0 0 2 0V8a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stethoscope(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 8a2.993 2.993 0 0 0-1 5.816V15.5a4.5 4.5 0 0 1-9 0v-.59A6.004 6.004 0 0 0 14 9V3a1 1 0 0 0-1-1h-2a1 1 0 0 0 0 2h1v5a4 4 0 0 1-8 0V4h1a1 1 0 0 0 0-2H3a1 1 0 0 0-1 1v6a6.004 6.004 0 0 0 5 5.91v.59a6.5 6.5 0 0 0 13 0v-1.684A2.993 2.993 0 0 0 19 8m0 4a1 1 0 1 1 1-1a1.001 1.001 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StethoscopeAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 8a2.993 2.993 0 0 0-1 5.816V15.5a4.5 4.5 0 0 1-9 0v-1.02l3.124-2.498A4.976 4.976 0 0 0 14 8.078V3a1 1 0 0 0-1-1h-2a1 1 0 0 0 0 2h1v4.078a2.986 2.986 0 0 1-1.125 2.342L8 12.72l-2.874-2.3A2.985 2.985 0 0 1 4 8.078V4h1a1 1 0 0 0 0-2H3a1 1 0 0 0-1 1v5.078a4.975 4.975 0 0 0 1.876 3.904L7 14.48v1.02a6.5 6.5 0 0 0 13 0v-1.684A2.993 2.993 0 0 0 19 8m0 4a1 1 0 1 1 1-1a1.001 1.001 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopCircle(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 8H9a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V9a1 1 0 0 0-1-1m-1 6h-4v-4h4ZM12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stopwatch(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m18.3 8.59l.91-.9a1 1 0 0 0-1.42-1.42l-.9.91a8 8 0 0 0-9.79 0l-.91-.92a1 1 0 0 0-1.42 1.43l.92.91A7.92 7.92 0 0 0 4 13.5a8 8 0 1 0 14.3-4.91M12 19.5a6 6 0 1 1 6-6a6 6 0 0 1-6 6m-2-15h4a1 1 0 0 0 0-2h-4a1 1 0 0 0 0 2m3 6a1 1 0 0 0-2 0v1.89a1.5 1.5 0 1 0 2 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopwatchSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.6 5.63a1 1 0 0 0 .36 2a6.18 6.18 0 0 1 1-.09a6 6 0 0 1 6 6a6.18 6.18 0 0 1-.09 1a1 1 0 0 0 .8 1.16h.18a1 1 0 0 0 1-.82A7.45 7.45 0 0 0 20 13.5a8 8 0 0 0-1.7-4.91l.91-.9a1 1 0 0 0-1.42-1.42l-.9.91A8 8 0 0 0 12 5.5a7.45 7.45 0 0 0-1.4.13M10 4.5h4a1 1 0 0 0 0-2h-4a1 1 0 0 0 0 2m3.49 9.08v-.16l1.34-1.33a1 1 0 1 0-1.42-1.42L12.08 12h-.16L5.71 5.79a1 1 0 0 0-1.42 1.42l.48.48l.91.91A8 8 0 0 0 16.9 19.82l1.39 1.39a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM12 19.5A6 6 0 0 1 7.11 10l3.4 3.39v.08A1.5 1.5 0 0 0 12 15h.08l3.39 3.4A6 6 0 0 1 12 19.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Store(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 7.82a1.25 1.25 0 0 0 0-.19l-2-5A1 1 0 0 0 19 2H5a1 1 0 0 0-.93.63l-2 5a1.25 1.25 0 0 0 0 .19A.58.58 0 0 0 2 8a4 4 0 0 0 2 3.4V21a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1v-9.56A4 4 0 0 0 22 8a.58.58 0 0 0 0-.18M13 20h-2v-4h2Zm5 0h-3v-5a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1v5H6v-8a4 4 0 0 0 3-1.38a4 4 0 0 0 6 0A4 4 0 0 0 18 12Zm0-10a2 2 0 0 1-2-2a1 1 0 0 0-2 0a2 2 0 0 1-4 0a1 1 0 0 0-2 0a2 2 0 0 1-4 .15L5.68 4h12.64L20 8.15A2 2 0 0 1 18 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StoreAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 2H2a1 1 0 0 0-1 1v4a3 3 0 0 0 2 2.82V21a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V9.82A3 3 0 0 0 23 7V3a1 1 0 0 0-1-1m-7 2h2v3a1 1 0 0 1-2 0Zm-4 0h2v3a1 1 0 0 1-2 0ZM7 4h2v3a1 1 0 0 1-2 0ZM4 8a1 1 0 0 1-1-1V4h2v3a1 1 0 0 1-1 1m10 12h-4v-4a2 2 0 0 1 4 0Zm5 0h-3v-4a4 4 0 0 0-8 0v4H5V9.82a3.17 3.17 0 0 0 1-.6a3 3 0 0 0 4 0a3 3 0 0 0 4 0a3 3 0 0 0 4 0a3.17 3.17 0 0 0 1 .6Zm2-13a1 1 0 0 1-2 0V4h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StoreSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.7 4h9.623l1.661 4.154A1.995 1.995 0 0 1 16 8a1 1 0 0 0-2 0a1.826 1.826 0 0 1-.134.709a1 1 0 0 0 .543 1.305a.947.947 0 0 0 .174.035A3.988 3.988 0 0 0 18 12v1.3a1 1 0 0 0 2 0v-1.856a3.985 3.985 0 0 0 1.996-3.405c0-.01.003-.018.003-.028L22 8a.949.949 0 0 0-.035-.171a.952.952 0 0 0-.036-.2l-2-5A1 1 0 0 0 19 2H8.7a1 1 0 1 0 0 2m14.007 17.293l-2.933-2.933a.943.943 0 0 0-.154-.154l-9.95-9.95a.973.973 0 0 0-.206-.206L5.182 3.768a.963.963 0 0 0-.128-.128L2.707 1.293a1 1 0 0 0-1.414 1.414l1.964 1.964l-1.178 2.94l-.007.017a.953.953 0 0 0-.035.189A.948.948 0 0 0 2 8l.001.012l.003.024A3.986 3.986 0 0 0 4 11.444V21a1 1 0 0 0 1 1h14a.993.993 0 0 0 .93-.656l1.363 1.363a1 1 0 0 0 1.414-1.414M4.016 8.153l.778-1.945L7.67 9.084A1.972 1.972 0 0 1 6 10a1.996 1.996 0 0 1-1.984-1.846ZM9 15v5H6v-8a3.964 3.964 0 0 0 3.102-1.484L12.586 14H10a1 1 0 0 0-1 1m4 5h-2v-4h2Zm5 0h-3v-3.586l3 3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Streering(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 12a1 1 0 1 0 1 1a1 1 0 0 0-1-1m9.71-2.36a10 10 0 0 0-19.4 0a9.75 9.75 0 0 0 0 4.72a10 10 0 0 0 7.3 7.34a9.67 9.67 0 0 0 4.7 0a10 10 0 0 0 7.31-7.31a9.75 9.75 0 0 0 0-4.72ZM12 4a8 8 0 0 1 7.41 5H4.59A8 8 0 0 1 12 4m-8 8a8.26 8.26 0 0 1 .07-1H6v2H4.07A8.26 8.26 0 0 1 4 12m5 7.41A8 8 0 0 1 4.59 15H7a2 2 0 0 1 2 2Zm4 .52a8.26 8.26 0 0 1-1 .07a8.26 8.26 0 0 1-1-.07V18h2Zm.14-3.93h-2.28A4 4 0 0 0 8 13.14V11h8v2.14A4 4 0 0 0 13.14 16M15 19.41V17a2 2 0 0 1 2-2h2.41A8 8 0 0 1 15 19.41M19.93 13H18v-2h1.93a8.26 8.26 0 0 1 .07 1a8.26 8.26 0 0 1-.07 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stretcher(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 4.5h3a1 1 0 0 0 0-2h-3a1 1 0 0 0 0 2m3 2H3a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h1.55l5.11 2.56l-2.58 1.29A3 3 0 0 0 5 15.5a3 3 0 1 0 3 3a2.2 2.2 0 0 0 0-.36l3.94-2l4.06 2.1a2.3 2.3 0 0 0 0 .26a3 3 0 1 0 3-3a3 3 0 0 0-2.15.92l-2.72-1.36l5.11-2.56H21a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1m-16 13a1 1 0 1 1 1-1a1 1 0 0 1-1 1m14-2a1 1 0 1 1-1 1a1 1 0 0 1 1-1m-7.1-3.56L9 12.5h5.75ZM20 10.5H4v-2h16Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Subject(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 16H3a1 1 0 0 0 0 2h10a1 1 0 0 0 0-2M3 8h18a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2m18 3H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Subway(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 17a1 1 0 0 0 1-1a1.36 1.36 0 0 0 0-.2a.64.64 0 0 0-.06-.18a.76.76 0 0 0-.09-.18l-.12-.15a1.15 1.15 0 0 0-.33-.21a1 1 0 0 0-.76 0a1.15 1.15 0 0 0-.33.21l-.12.15a.76.76 0 0 0-.09.18a.64.64 0 0 0-.06.18a1.36 1.36 0 0 0 0 .2a1 1 0 0 0 1 1Zm2-15H6a3 3 0 0 0-3 3v12a3 3 0 0 0 1.2 2.39l-.91.9a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0L6.41 20h11.18l1.7 1.71a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-.91-.9A3 3 0 0 0 21 17V5a3 3 0 0 0-3-3M5 8h6v4H5Zm14 9a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1v-3h14Zm0-5h-6V8h6Zm0-6H5V5a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1ZM8 17a1 1 0 0 0 1-1a1.36 1.36 0 0 0 0-.2a.64.64 0 0 0-.06-.18a.76.76 0 0 0-.09-.18l-.12-.15a1.15 1.15 0 0 0-.33-.21a1 1 0 0 0-.76 0a1.15 1.15 0 0 0-.33.21l-.12.15a.76.76 0 0 0-.09.18a.64.64 0 0 0-.1.18a1.36 1.36 0 0 0 0 .2a1 1 0 0 0 .29.7A1 1 0 0 0 8 17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SubwayAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 17h1a1 1 0 0 0 0-2h-1a1 1 0 0 0 0 2m6 0V9a3 3 0 0 0-3-3h-5V4h4a1 1 0 0 0 0-2H7a1 1 0 0 0 0 2h4v2H6a3 3 0 0 0-3 3v8a3 3 0 0 0 1.2 2.39l-.91.9a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0L6.41 20h11.18l1.7 1.71a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-.91-.9A3 3 0 0 0 21 17M5 9a1 1 0 0 1 1-1h5v4H5Zm14 8a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1v-3h14Zm0-5h-6V8h5a1 1 0 0 1 1 1ZM8 17h1a1 1 0 0 0 0-2H8a1 1 0 0 0 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Suitcase(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 6.5h-3v-1a3 3 0 0 0-3-3h-2a3 3 0 0 0-3 3v1H5a3 3 0 0 0-3 3v9a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-9a3 3 0 0 0-3-3m-9-1a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v1h-4Zm10 13a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-5.05h3v1.05a1 1 0 0 0 2 0v-1.05h6v1.05a1 1 0 0 0 2 0v-1.05h3Zm0-7H4v-2a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SuitcaseAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 6h-3V5a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v1H5a3 3 0 0 0-3 3v9a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3m-9-1h4v1h-4Zm10 13a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-5h4v1a1 1 0 0 0 2 0v-1h4v1a1 1 0 0 0 2 0v-1h4Zm0-7H4V9a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sun(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.64 17l-.71.71a1 1 0 0 0 0 1.41a1 1 0 0 0 1.41 0l.71-.71A1 1 0 0 0 5.64 17M5 12a1 1 0 0 0-1-1H3a1 1 0 0 0 0 2h1a1 1 0 0 0 1-1m7-7a1 1 0 0 0 1-1V3a1 1 0 0 0-2 0v1a1 1 0 0 0 1 1M5.64 7.05a1 1 0 0 0 .7.29a1 1 0 0 0 .71-.29a1 1 0 0 0 0-1.41l-.71-.71a1 1 0 0 0-1.41 1.41Zm12 .29a1 1 0 0 0 .7-.29l.71-.71a1 1 0 1 0-1.41-1.41l-.64.71a1 1 0 0 0 0 1.41a1 1 0 0 0 .66.29ZM21 11h-1a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2m-9 8a1 1 0 0 0-1 1v1a1 1 0 0 0 2 0v-1a1 1 0 0 0-1-1m6.36-2A1 1 0 0 0 17 18.36l.71.71a1 1 0 0 0 1.41 0a1 1 0 0 0 0-1.41ZM12 6.5a5.5 5.5 0 1 0 5.5 5.5A5.51 5.51 0 0 0 12 6.5m0 9a3.5 3.5 0 1 1 3.5-3.5a3.5 3.5 0 0 1-3.5 3.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sunset(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.66 8.34a1 1 0 0 0 .7-.29l.71-.71a1 1 0 1 0-1.41-1.41l-.66.71a1 1 0 0 0 0 1.41a1 1 0 0 0 .66.29M12 6a1 1 0 0 0 1-1V4a1 1 0 0 0-2 0v1a1 1 0 0 0 1 1m-8 6H3a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2m1.64-3.95a1 1 0 0 0 .7.29a1 1 0 0 0 .71-.29a1 1 0 0 0 0-1.41l-.71-.71a1 1 0 0 0-1.41 1.41ZM21 12h-1a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2m-10 7H8a1 1 0 0 0 0 2h3a1 1 0 0 0 0-2m7-4h-.88a5.39 5.39 0 0 0 .38-2a5.5 5.5 0 0 0-11 0a5.39 5.39 0 0 0 .38 2H6a1 1 0 0 0 0 2h12a1 1 0 0 0 0-2m-3.15 0h-5.7a3.44 3.44 0 0 1-.65-2a3.5 3.5 0 0 1 7 0a3.44 3.44 0 0 1-.65 2M16 19h-1a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Surprise(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8M10 9a1 1 0 1 0-1 1a1 1 0 0 0 1-1m5-1a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-3 3a3.39 3.39 0 0 0-3.25 3.5A3.39 3.39 0 0 0 12 18a3.39 3.39 0 0 0 3.25-3.5A3.39 3.39 0 0 0 12 11m0 5a1.39 1.39 0 0 1-1.25-1.5A1.39 1.39 0 0 1 12 13a1.39 1.39 0 0 1 1.25 1.5A1.39 1.39 0 0 1 12 16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Swatchbook(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 16a1 1 0 1 0 1 1a1 1 0 0 0-1-1m12.06-4l1.23-1.23a3 3 0 0 0 0-4.24l-2.83-2.82a3 3 0 0 0-4.24 0L12 4.94A3 3 0 0 0 9 2H5a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-4a3 3 0 0 0-2.94-3M10 19a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1Zm2-11.24l2.64-2.64a1 1 0 0 1 1.41 0L18.88 8a1 1 0 0 1 0 1.41L16 12.29l-4 3.95ZM20 19a1 1 0 0 1-1 1h-7.18a3.12 3.12 0 0 0 .17-.92L17.07 14H19a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Swiggy(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.878 5.928a.297.297 0 0 1 .314.335l.003.276c.001.593-.002 1.185.002 1.777c.004.434.08.517.5.57a14.832 14.832 0 0 0 3.148-.046a4.948 4.948 0 0 0 1.573-.374a.507.507 0 0 0 .344-.603a6.831 6.831 0 0 0-5.343-5.71a6.608 6.608 0 0 0-3.648.242a6.718 6.718 0 0 0-2.694 1.752A6.272 6.272 0 0 0 5.23 8.432a11.188 11.188 0 0 0 1.526 5.517a1.342 1.342 0 0 0 1.33.748c.653-.016 1.307-.004 1.96-.004v-.003h2.084c.25 0 .447.06.445.372c-.005.726 0 1.45-.003 2.176c-.002.22-.064.432-.326.434c-.264.002-.327-.209-.33-.43c-.005-.347 0-.694 0-1.041c0-.451-.073-.557-.523-.64a7.702 7.702 0 0 0-2.348-.02a2.596 2.596 0 0 0-.8.196c-.212.1-.296.245-.193.467c.107.231.206.467.333.688a43.875 43.875 0 0 0 3.426 4.956c.154.199.273.203.431.01c.349-.426.715-.839 1.06-1.268a34.226 34.226 0 0 0 3.577-5.26a14.171 14.171 0 0 0 1.494-3.871a1.203 1.203 0 0 0-1.012-1.635a6.943 6.943 0 0 0-1.67-.162c-.909-.006-1.818-.003-2.727-.002c-.245 0-.427-.07-.427-.358c0-1.01-.002-2.022.002-3.033c0-.216.104-.349.339-.34"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Swimmer(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 9.28a2 2 0 1 0-2-2a2 2 0 0 0 2 2M2.71 13.92a3 3 0 0 1 .57.25a4.06 4.06 0 0 0 4.1 0a2.62 2.62 0 0 1 2.56 0l.21.1a4.14 4.14 0 0 0 3.87-.13a2.62 2.62 0 0 1 2.56 0a4.25 4.25 0 0 0 2.08.56a4 4 0 0 0 2-.56a3 3 0 0 1 .57-.25a1 1 0 1 0-.52-1.89a4.82 4.82 0 0 0-1 .44a2.1 2.1 0 0 1-2.1 0a4.62 4.62 0 0 0-4.54 0a2.52 2.52 0 0 1-.29.12L14.34 11a.75.75 0 0 0 .09-.15a1 1 0 0 0 .12-.16a1.29 1.29 0 0 0 0-.19a1.06 1.06 0 0 0 0-.19a1.13 1.13 0 0 0 0-.18a1.06 1.06 0 0 0 0-.19a1.51 1.51 0 0 0-.1-.17a.7.7 0 0 0-.1-.16l-3.07-3.33a1.1 1.1 0 0 0-.28-.22a.8.8 0 0 0-.21 0a.53.53 0 0 0-.17 0A.89.89 0 0 0 10.2 6L6.66 7.32a1 1 0 0 0-.6 1.28a1 1 0 0 0 1.28.6l2.93-1.07l2 2.12l-1.9 1.9a4.62 4.62 0 0 0-3.94.28a2.1 2.1 0 0 1-2.1 0a4.82 4.82 0 0 0-1-.44a1 1 0 1 0-.58 1.91Zm18 3.09a4.82 4.82 0 0 0-1 .44a2.1 2.1 0 0 1-2.1 0a4.62 4.62 0 0 0-4.54 0a2.11 2.11 0 0 1-2.12 0a4.62 4.62 0 0 0-4.54 0a2.1 2.1 0 0 1-2.1 0a4.82 4.82 0 0 0-1-.44a1 1 0 1 0-.58 1.91a3 3 0 0 1 .57.25a4.06 4.06 0 0 0 4.1 0a2.62 2.62 0 0 1 2.56 0a4.15 4.15 0 0 0 4.12 0a2.62 2.62 0 0 1 2.56 0a4.25 4.25 0 0 0 2.08.56a4 4 0 0 0 2-.56a3 3 0 0 1 .57-.25a1 1 0 1 0-.58-1.92Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sync(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.91 15.51h-4.53a1 1 0 0 0 0 2h2.4A8 8 0 0 1 4 12a1 1 0 0 0-2 0a10 10 0 0 0 16.88 7.23V21a1 1 0 0 0 2 0v-4.5a1 1 0 0 0-.97-.99M12 2a10 10 0 0 0-6.88 2.77V3a1 1 0 0 0-2 0v4.5a1 1 0 0 0 1 1h4.5a1 1 0 0 0 0-2h-2.4A8 8 0 0 1 20 12a1 1 0 0 0 2 0A10 10 0 0 0 12 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SyncExclamation(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.29 15.71A1 1 0 0 0 13 15a1.05 1.05 0 0 0-.29-.71a1 1 0 0 0-1.09-.21a1 1 0 0 0-.33.21A1.05 1.05 0 0 0 11 15a1 1 0 0 0 .29.71m8.62-.2h-4.53a1 1 0 0 0 0 2h2.4A8 8 0 0 1 4 12a1 1 0 0 0-2 0a10 10 0 0 0 16.88 7.23V21a1 1 0 0 0 2 0v-4.5a1 1 0 0 0-.97-.99M12 2a10 10 0 0 0-6.88 2.77V3a1 1 0 0 0-2 0v4.5a1 1 0 0 0 1 1h4.5a1 1 0 0 0 0-2h-2.4A8 8 0 0 1 20 12a1 1 0 0 0 2 0A10 10 0 0 0 12 2m0 11a1 1 0 0 0 1-1V9a1 1 0 0 0-2 0v3a1 1 0 0 0 1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SyncSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.88 15.5h-4.5a1 1 0 0 0 0 2h2.4A8 8 0 0 1 12 20a8.08 8.08 0 0 1-3.12-.63l-1.49 1.49A9.83 9.83 0 0 0 12 22a10 10 0 0 0 6.88-2.77V21a1 1 0 0 0 2 0v-4.5a1 1 0 0 0-1-1m-1.57-8.4l1.43-1.43l2-2a1 1 0 1 0-1.42-1.42l-2 2A9.89 9.89 0 0 0 12 2a10 10 0 0 0-6.88 2.77V3a1 1 0 0 0-2 0v4.5a1 1 0 0 0 1 1h4.5a1 1 0 0 0 0-2h-2.4a7.93 7.93 0 0 1 10.67-.81l-11.2 11.2A7.93 7.93 0 0 1 4 12a1 1 0 0 0-2 0a9.89 9.89 0 0 0 2.27 6.32l-2 2a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l2-2l1.43-1.43Zm1.06 1.78A8.08 8.08 0 0 1 20 12a1 1 0 0 0 2 0a9.83 9.83 0 0 0-1.14-4.61Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Syringe(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.71 2.29a1 1 0 0 0-1.42 0l-2.12 2.12l-.71-.7a1 1 0 0 0-1.41 0l-1.41 1.41l-.71-.71a1 1 0 0 0-1.42 0l-7.77 7.78l-.74-.7a1 1 0 0 0-1.38 1.41l3.53 3.54l-1.73 1.74l-.71-.72a1 1 0 0 0-1.42 1.42l2.83 2.83a1 1 0 0 0 .71.29a1 1 0 0 0 .71-.29a1 1 0 0 0 0-1.42l-.71-.7l1.74-1.74l3.53 3.53a1 1 0 0 0 .71.3a1 1 0 0 0 .7-1.71l-.7-.71l7.78-7.77a1 1 0 0 0 0-1.42l-.71-.71L20.29 8a1 1 0 0 0 0-1.41l-.7-.71l2.12-2.12a1 1 0 0 0 0-1.47M7.57 15l-1.42-1.39l1.41-1.42L9 13.61Zm2.82 2.83L9 16.44L10.39 15l1.42 1.42ZM13.22 15L9 10.78l4.24-4.24l.71.7l3.53 3.54Zm4.24-7l-1.41-1.46l.71-.71l.7.7l.7.7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Table(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1M8 20H4v-4h4Zm0-6H4v-4h4Zm0-6H4V4h4Zm6 12h-4v-4h4Zm0-6h-4v-4h4Zm0-6h-4V4h4Zm6 12h-4v-4h4Zm0-6h-4v-4h4Zm0-6h-4V4h4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableTennis(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.71 16.1l-1.64-1.64a.94.94 0 0 1-.22-1.07a5.78 5.78 0 0 0 .54-2.39a.36.36 0 0 0 0-.1a5.74 5.74 0 0 0-1.06-3.34a14.17 14.17 0 0 0-5.17-4.42a7 7 0 0 0-8 1.31l-.67.67a7 7 0 0 0-1.31 8.05l.1.17a3 3 0 0 0-.84 2.06A3 3 0 0 0 7 17.94c.18.14.34.29.52.42a5.55 5.55 0 0 0 1.22.64h.09c.18.07.37.13.57.19h.15a5.08 5.08 0 0 0 .95.15h.62c.21 0 .41 0 .62-.06h.17a5.46 5.46 0 0 0 1.42-.45a1 1 0 0 1 1.07.21l1.46 1.46a3.4 3.4 0 0 0 2.39 1a2.85 2.85 0 0 0 2-.85l.38-.38a3 3 0 0 0 .08-4.17m-15.3.32a1 1 0 1 1 1-1a1 1 0 0 1-1 1m3.1.14l-.26-.2a2.94 2.94 0 0 0 .16-.94a3 3 0 0 0-3-3h-.38l-.09-.16a5 5 0 0 1 .93-5.75l.67-.67A5 5 0 0 1 12.28 5a12 12 0 0 1 4.26 3.57Zm10.78 2.37l-.37.38c-.42.42-1.07.34-1.61-.2l-1.46-1.45a3 3 0 0 0-3.34-.61a3.39 3.39 0 0 1-1 .31a2.84 2.84 0 0 1-.58.05h-.44l6.87-6.87a3.8 3.8 0 0 1-.34 2a3 3 0 0 0 .61 3.34l1.64 1.65a1 1 0 0 1 .02 1.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tablet(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2H7a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3m1 17a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1Zm-5.29-2.71a1 1 0 0 0-.91-.29l-.18.06a.76.76 0 0 0-.18.09l-.15.12a1 1 0 0 0-.21.33a1 1 0 0 0 .21 1.09a1.46 1.46 0 0 0 .33.22a1 1 0 0 0 1.09-.22A1 1 0 0 0 13 17a.84.84 0 0 0-.08-.38a1.15 1.15 0 0 0-.21-.33"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tablets(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.54 3.46a5 5 0 1 0 0 7.08a5 5 0 0 0 0-7.08m-5.66 1.42A3 3 0 0 1 17 4a3 3 0 0 1 1.28.3l-4 4a3 3 0 0 1 .6-3.42m4.24 4.24a3 3 0 0 1-3.4.58l4-4a3 3 0 0 1-.6 3.42M8 10a5.93 5.93 0 0 0-4.21 1.73A6 6 0 0 0 8 22a6 6 0 0 0 4.14-1.66l.12-.08a1.05 1.05 0 0 1 .07-.11A6 6 0 0 0 8 10m-2.83 8.83A4 4 0 0 1 4.56 14L10 19.44a4 4 0 0 1-4.83-.61m6.27-.83L6 12.56A4 4 0 0 1 8 12a4 4 0 0 1 3.44 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TachometerFast(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.29 10.29l-2.78 2.78A2.09 2.09 0 0 0 12 13a2 2 0 0 0-2 2a2.09 2.09 0 0 0 .07.51l-.78.78a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l.78-.78A2.09 2.09 0 0 0 12 17a2 2 0 0 0 2-2a2.09 2.09 0 0 0-.07-.51l2.78-2.78a1 1 0 0 0-1.42-1.42M12 4A10 10 0 0 0 2 14a9.91 9.91 0 0 0 1.69 5.56a1 1 0 0 0 1.66-1.12a8 8 0 1 1 13.3 0a1 1 0 0 0 .27 1.39a1 1 0 0 0 .56.17a1 1 0 0 0 .83-.44A9.91 9.91 0 0 0 22 14A10 10 0 0 0 12 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TachometerFastAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 5a10 10 0 0 0-8.66 15a1 1 0 0 0 1.74-1A7.92 7.92 0 0 1 4 15a8 8 0 1 1 14.93 4a1 1 0 0 0 .37 1.37a1 1 0 0 0 1.36-.37A10 10 0 0 0 12 5m2.84 5.76l-1.55 1.54A2.91 2.91 0 0 0 12 12a3 3 0 1 0 3 3a2.9 2.9 0 0 0-.3-1.28l1.55-1.54a1 1 0 0 0 0-1.42a1 1 0 0 0-1.41 0M12 16a1 1 0 0 1 0-2a1 1 0 0 1 .7.28a1 1 0 0 1 .3.72a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tag(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.12 10.71l-8.41-8.42A1 1 0 0 0 12 2H3a1 1 0 0 0-1 1v9a1 1 0 0 0 .29.71l8.42 8.41a3 3 0 0 0 4.24 0L21.12 15a3 3 0 0 0 0-4.24Zm-1.41 2.82l-6.18 6.17a1 1 0 0 1-1.41 0L4 11.59V4h7.59l8.12 8.12a1 1 0 0 1 .29.71a1 1 0 0 1-.29.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.5 6A1.5 1.5 0 1 0 9 7.5A1.5 1.5 0 0 0 7.5 6m13.62 4.71l-8.41-8.42A1 1 0 0 0 12 2H3a1 1 0 0 0-1 1v9a1 1 0 0 0 .29.71l8.42 8.41a3 3 0 0 0 4.24 0L21.12 15a3 3 0 0 0 0-4.24Zm-1.41 2.82l-6.18 6.17a1 1 0 0 1-1.41 0L4 11.59V4h7.59l8.12 8.12a1 1 0 0 1 .29.71a1 1 0 0 1-.29.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tape(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 7a4 4 0 1 0 4 4a4 4 0 0 0-4-4m0 6a2 2 0 1 1 2-2a2 2 0 0 1-2 2m1-10h-2a7 7 0 0 0-7 7v3h-1a1 1 0 0 0 0 2h1v2h-4a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0v-1h12a7 7 0 0 0 7-7v-2a7 7 0 0 0-7-7m5 9a5 5 0 0 1-5 5h-7v-7a5 5 0 0 1 5-5h2a5 5 0 0 1 5 5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Taxi(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.38 14.58a1 1 0 0 0-.58-.06a.64.64 0 0 0-.18.06a.76.76 0 0 0-.18.09l-.15.12a1.05 1.05 0 0 0-.29.71a1 1 0 0 0 .29.71a1.19 1.19 0 0 0 .33.22a1 1 0 0 0 1.09-.22a1.15 1.15 0 0 0 .21-.33a.84.84 0 0 0 .08-.38a1.05 1.05 0 0 0-.29-.71a.93.93 0 0 0-.33-.21m2.62-3.9V7.5a3 3 0 0 0-3-3h-.78l-.77-2.32a1 1 0 0 0-.95-.68h-5a1 1 0 0 0-.95.68L7.78 4.5H7a3 3 0 0 0-3 3v3.18a3 3 0 0 0-2 2.82v6a1 1 0 0 0 1 1h1v1a1 1 0 0 0 2 0v-1h12v1a1 1 0 0 0 2 0v-1h1a1 1 0 0 0 1-1v-6a3 3 0 0 0-2-2.82M10.22 3.5h3.56l.33 1H9.89ZM6 7.5a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v3H6Zm14 11H4v-5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1ZM6.62 14.58a1 1 0 0 0-.33.21a1.05 1.05 0 0 0-.29.71a1 1 0 0 0 1.38.93a1.19 1.19 0 0 0 .33-.22A1 1 0 0 0 8 15.5a1.05 1.05 0 0 0-.29-.71a1 1 0 0 0-1.09-.21M13 14.5h-2a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tear(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.56 2.17a1 1 0 0 0-1.12 0c-.3.2-7.19 5-7.19 12.08a7.75 7.75 0 0 0 15.5 0c0-7.2-6.9-11.89-7.19-12.08M12 20a5.76 5.76 0 0 1-5.75-5.75c0-5 4.21-8.77 5.75-10c1.55 1.21 5.75 5 5.75 10A5.76 5.76 0 0 1 12 20"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Telegram(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.994 2a10 10 0 1 0 10 10a10 10 0 0 0-10-10m3.18 15.152a.705.705 0 0 1-1.002.352l-2.715-2.11l-1.742 1.608a.3.3 0 0 1-.285.039l.334-2.989l.01.009l.007-.059s4.885-4.448 5.084-4.637c.202-.189.135-.23.135-.23c.012-.23-.361 0-.361 0l-6.473 4.164l-2.695-.918s-.414-.148-.453-.475c-.041-.324.466-.5.466-.5l10.717-4.258s.881-.392.881.258Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TelegramAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.265 2.428a2.048 2.048 0 0 0-2.078-.324L2.266 9.339a2.043 2.043 0 0 0 .104 3.818l3.625 1.261l2.02 6.682a.998.998 0 0 0 .119.252c.008.012.019.02.027.033a.988.988 0 0 0 .211.215a.972.972 0 0 0 .07.05a.986.986 0 0 0 .31.136l.013.001l.006.003a1.022 1.022 0 0 0 .203.02l.018-.003a.993.993 0 0 0 .301-.052c.023-.008.042-.02.064-.03a.993.993 0 0 0 .205-.114a250.76 250.76 0 0 1 .152-.129l2.702-2.983l4.03 3.122a2.023 2.023 0 0 0 1.241.427a2.054 2.054 0 0 0 2.008-1.633l3.263-16.017a2.03 2.03 0 0 0-.693-1.97M9.37 14.736a.994.994 0 0 0-.272.506l-.31 1.504l-.784-2.593l4.065-2.117Zm8.302 5.304l-4.763-3.69a1.001 1.001 0 0 0-1.353.12l-.866.955l.306-1.487l7.083-7.083a1 1 0 0 0-1.169-1.593L6.745 12.554L3.02 11.191L20.999 4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Telescope(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.9 7.59l-1-3.87a3 3 0 0 0-3.71-2.12l-1.93.52a1 1 0 0 0-.71 1.23l.26 1L4.19 7.16a1 1 0 0 0-.71 1.22l.26 1l-1 .26a1 1 0 0 0 .25 2a1.09 1.09 0 0 0 .26 0l1-.27l.26 1a1 1 0 0 0 .46.6a1 1 0 0 0 .5.14a.75.75 0 0 0 .26 0L9 12.08v.42a2.9 2.9 0 0 0 .3 1.28l-5 5a1 1 0 0 0 1.41 1.42l5-5l.28.11v6.19a1 1 0 0 0 2 0v-6.18a2.52 2.52 0 0 0 .29-.12l5 5a1 1 0 1 0 1.41-1.42l-5-5A3.09 3.09 0 0 0 15 12.5v-2l1.35-.36l.25 1a1 1 0 0 0 1 .74h.26l1.93-.52a3 3 0 0 0 2.11-3.77M13 12.5a1 1 0 0 1-.28.69a1 1 0 0 1-.69.28a1 1 0 0 1-.7-.29a1 1 0 0 1-.29-.7v-1L13 11Zm-6.81-1.74l-.52-1.93l9.66-2.59l.26 1l.26 1Zm13.68-1.9a1 1 0 0 1-.61.47l-1 .26l-.78-2.9L17 4.76l-.26-1l1-.26a1 1 0 0 1 1.23.71l1 3.87a1 1 0 0 1-.1.78"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Temperature(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 15.28V5.5a1 1 0 0 0-2 0v9.78A2 2 0 0 0 10 17a2 2 0 0 0 4 0a2 2 0 0 0-1-1.72M16.5 13V5.5a4.5 4.5 0 0 0-9 0V13a6 6 0 0 0 3.21 9.83A7 7 0 0 0 12 23a6 6 0 0 0 4.5-10m-2 7.07a4 4 0 0 1-6.42-2.2a4 4 0 0 1 1.1-3.76a1 1 0 0 0 .3-.71V5.5a2.5 2.5 0 0 1 5 0v7.94a1 1 0 0 0 .3.71a4 4 0 0 1-.28 6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TemperatureEmpty(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 15a2 2 0 1 0 2 2a2 2 0 0 0-2-2m4.5-2V5.5a4.5 4.5 0 0 0-9 0V13a6 6 0 0 0 3.21 9.83A7 7 0 0 0 12 23a6 6 0 0 0 4.5-10m-2 7.07a4 4 0 0 1-5.32-6a1 1 0 0 0 .3-.71V5.5a2.5 2.5 0 0 1 5 0v7.94a1 1 0 0 0 .3.71a4 4 0 0 1-.28 6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TemperatureHalf(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 15.28V10.5a1 1 0 0 0-2 0v4.78A2 2 0 0 0 10 17a2 2 0 0 0 4 0a2 2 0 0 0-1-1.72M16.5 13V5.5a4.5 4.5 0 0 0-9 0V13a6 6 0 0 0 3.21 9.83A7 7 0 0 0 12 23a6 6 0 0 0 4.5-10m-2 7.07a4 4 0 0 1-5.32-6a1 1 0 0 0 .3-.71V5.5a2.5 2.5 0 0 1 5 0v7.94a1 1 0 0 0 .3.71a4 4 0 0 1-.28 6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TemperatureMinus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 5.5a4.5 4.5 0 0 0-9 0V13a6 6 0 0 0 3.21 9.83a7 7 0 0 0 1.28.17A6 6 0 0 0 14 13Zm-2 14.61a4 4 0 0 1-5.32-6a1 1 0 0 0 .3-.71V5.5a2.5 2.5 0 0 1 5 0v7.94a1 1 0 0 0 .3.71a4 4 0 0 1-.28 6Zm-1.5-4.83V5.5a1 1 0 0 0-2 0v9.78a2 2 0 0 0-1 1.72a2 2 0 0 0 4 0a2 2 0 0 0-1-1.72m9-12.78h-3a1 1 0 0 0 0 2h3a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TemperaturePlus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5 15.28V5.5a1 1 0 0 0-2 0v9.78a2 2 0 0 0-1 1.72a2 2 0 0 0 4 0a2 2 0 0 0-1-1.72m9-12.78H19V2a1 1 0 0 0-2 0v.5h-.5a1 1 0 0 0 0 2h.5V5a1 1 0 0 0 2 0v-.5h.5a1 1 0 0 0 0-2m-5.5 3a4.5 4.5 0 0 0-9 0V13a6 6 0 0 0 3.21 9.83a7 7 0 0 0 1.28.17A6 6 0 0 0 14 13Zm-2 14.61a4 4 0 0 1-6.42-2.2a4 4 0 0 1 1.1-3.76a1 1 0 0 0 .3-.71V5.5a2.5 2.5 0 0 1 5 0v7.94a1 1 0 0 0 .3.71a4 4 0 0 1-.28 6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TemperatureQuarter(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.5 13V5.5a4.5 4.5 0 0 0-9 0V13a6 6 0 0 0 3.21 9.83A7 7 0 0 0 12 23a6 6 0 0 0 4.5-10m-2 7.07a4 4 0 0 1-5.32-6a1 1 0 0 0 .3-.71V5.5a2.5 2.5 0 0 1 5 0v7.94a1 1 0 0 0 .3.71a4 4 0 0 1-.28 6ZM13 15.28V12.5a1 1 0 0 0-2 0v2.78A2 2 0 0 0 10 17a2 2 0 0 0 4 0a2 2 0 0 0-1-1.72"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TemperatureThreeQuarter(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 15.28V8.5a1 1 0 0 0-2 0v6.78A2 2 0 0 0 10 17a2 2 0 0 0 4 0a2 2 0 0 0-1-1.72M16.5 13V5.5a4.5 4.5 0 0 0-9 0V13a6 6 0 0 0 3.21 9.83A7 7 0 0 0 12 23a6 6 0 0 0 4.5-10m-2 7.07a4 4 0 0 1-5.32-6a1 1 0 0 0 .3-.71V5.5a2.5 2.5 0 0 1 5 0v7.94a1 1 0 0 0 .3.71a4 4 0 0 1-.28 6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TenPlus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 5h1v1a1 1 0 0 0 2 0V5h1a1 1 0 0 0 0-2h-1V2a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2m-7 4v6a3 3 0 0 0 3 3h1a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3h-1a3 3 0 0 0-3 3m5 0v6a1 1 0 0 1-1 1h-1a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1h1a1 1 0 0 1 1 1m5.6 0a1 1 0 0 0-.78 1.18a9 9 0 1 1-7-7a1 1 0 1 0 .4-2A10.8 10.8 0 0 0 12 1a11 11 0 1 0 11 11a10.8 10.8 0 0 0-.22-2.2A1 1 0 0 0 21.6 9M7 7v10a1 1 0 0 0 2 0V7a1 1 0 0 0-2 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TennisBall(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a10 10 0 0 0-7.35 16.76l.09.1a10 10 0 0 0 14.52 0l.09-.1A10 10 0 0 0 12 2M5.61 16.79a7.93 7.93 0 0 1 0-9.58a6 6 0 0 1 0 9.58M12 20a7.91 7.91 0 0 1-5-1.77A8 8 0 0 0 7 5.77a7.95 7.95 0 0 1 10 0a8 8 0 0 0 0 12.46A7.91 7.91 0 0 1 12 20m6.39-3.21a6 6 0 0 1 0-9.58a7.93 7.93 0 0 1 0 9.58"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextFields(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 6a1 1 0 0 0-1-1H3a1 1 0 0 0 0 2h5v11a1 1 0 0 0 2 0V7h5a1 1 0 0 0 1-1m5 5h-6a1 1 0 0 0 0 2h2v5a1 1 0 0 0 2 0v-5h2a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextIcon(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 6H7a1 1 0 0 0 0 2h4v9a1 1 0 0 0 2 0V8h4a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextSize(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 11H3a1 1 0 0 0 0 2h2v5a1 1 0 0 0 2 0v-5h2a1 1 0 0 0 0-2m12-6H9a1 1 0 0 0 0 2h5v11a1 1 0 0 0 2 0V7h5a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextStrikeThrough(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 13H9a1 1 0 0 0 0 2h2v2a1 1 0 0 0 2 0v-2h2a1 1 0 0 0 0-2m2-7H7a1 1 0 0 0 0 2h4v2a1 1 0 0 0 2 0V8h4a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Th(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1M11 20H4v-7h7Zm0-9H4V4h7Zm9 9h-7v-7h7Zm0-9h-7V4h7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThLarge(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 3H4a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1m-9 16H5v-6h6Zm0-8H5V5h6Zm8 8h-6v-6h6Zm0-8h-6V5h6Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m22.71 21.29l-1-1l-6-6l-6-6l-6-6l-1-1a1 1 0 0 0-1.42 1.42l.71.7V21a1 1 0 0 0 1 1h17.59l.7.71a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42M4 5.41L6.59 8H4ZM8 20H4v-4h4Zm0-6H4v-4h4Zm2-2.59L12.59 14H10ZM14 20h-4v-4h4Zm2 0v-2.59L18.59 20ZM8.67 4H14v5a1 1 0 0 0 1 1h5v5.33a1 1 0 1 0 2 0V3a1 1 0 0 0-1-1H8.67a1 1 0 0 0 0 2M16 4h4v4h-4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Thermometer(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16.29 6.29l-7 7a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l7-7a1 1 0 1 0-1.42-1.42m4.25-2.83a5 5 0 0 0-7.08 0l-8.17 8.23a1 1 0 0 0-.29.7v5.19l-2.71 2.71a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0L6.42 19h5.19a1 1 0 0 0 .7-.29l8.23-8.17a5 5 0 0 0 0-7.08m-1.42 5.66L11.2 17H7v-4.2l7.88-7.92a3 3 0 0 1 4.24 4.24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThirteenPlus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 5h1v1a1 1 0 0 0 2 0V5h1a1 1 0 0 0 0-2h-1V2a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2m-7 4a1 1 0 0 0 2 0a1 1 0 0 1 1-1h1a1 1 0 0 1 1 1v1a1 1 0 0 1-1 1h-.5a1 1 0 0 0 0 2h.5a1 1 0 0 1 1 1v1a1 1 0 0 1-1 1h-1a1 1 0 0 1-1-1a1 1 0 0 0-2 0a3 3 0 0 0 3 3h1a3 3 0 0 0 3-3v-1a3 3 0 0 0-.78-2a3 3 0 0 0 .78-2V9a3 3 0 0 0-3-3h-1a3 3 0 0 0-3 3M7 7v10a1 1 0 0 0 2 0V7a1 1 0 0 0-2 0m14.6 2a1 1 0 0 0-.78 1.18a9 9 0 1 1-7-7a1 1 0 1 0 .4-2A10.8 10.8 0 0 0 12 1a11 11 0 1 0 11 11a10.8 10.8 0 0 0-.22-2.2A1 1 0 0 0 21.6 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThreePlus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 5h1v1a1 1 0 0 0 2 0V5h1a1 1 0 0 0 0-2h-1V2a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2m-2.5 4a3 3 0 0 0-3-3h-1a3 3 0 0 0-3 3a1 1 0 0 0 2 0a1 1 0 0 1 1-1h1a1 1 0 0 1 1 1v1a1 1 0 0 1-1 1H12a1 1 0 0 0 0 2h.5a1 1 0 0 1 1 1v1a1 1 0 0 1-1 1h-1a1 1 0 0 1-1-1a1 1 0 0 0-2 0a3 3 0 0 0 3 3h1a3 3 0 0 0 3-3v-1a3 3 0 0 0-.78-2a3 3 0 0 0 .78-2Zm6.1 0a1 1 0 0 0-.78 1.18a9 9 0 1 1-7-7a1 1 0 1 0 .4-2A10.8 10.8 0 0 0 12 1a11 11 0 1 0 11 11a10.8 10.8 0 0 0-.22-2.2A1 1 0 0 0 21.6 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsDown(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 2H6.27a3 3 0 0 0-2.95 2.46l-1.27 7A3 3 0 0 0 5 15h4.56L9 16.43A4.13 4.13 0 0 0 12.89 22a1 1 0 0 0 .91-.59L16.65 15H19a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3m-4 11.79l-2.72 6.12a2.13 2.13 0 0 1-1.38-2.78l.53-1.43A2 2 0 0 0 9.56 13H5a1 1 0 0 1-.77-.36a1 1 0 0 1-.23-.82l1.27-7a1 1 0 0 1 1-.82H15ZM20 12a1 1 0 0 1-1 1h-2V4h2a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsUp(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.3 10.08A3 3 0 0 0 19 9h-4.56L15 7.57A4.13 4.13 0 0 0 11.11 2a1 1 0 0 0-.91.59L7.35 9H5a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h12.73a3 3 0 0 0 2.95-2.46l1.27-7a3 3 0 0 0-.65-2.46M7 20H5a1 1 0 0 1-1-1v-7a1 1 0 0 1 1-1h2Zm13-7.82l-1.27 7a1 1 0 0 1-1 .82H9v-9.79l2.72-6.12a2.11 2.11 0 0 1 1.38 2.78l-.53 1.43a2 2 0 0 0 1.87 2.7H19a1 1 0 0 1 .77.36a1 1 0 0 1 .23.82"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Thunderstorm(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.5 15h-2.27l1.45-2.5a1 1 0 1 0-1.74-1l-2.31 4a1 1 0 0 0 0 1a1 1 0 0 0 .87.5h2.27l-1.45 2.5a1 1 0 0 0 1.74 1l2.31-4a1 1 0 0 0 0-1a1 1 0 0 0-.87-.5m4.92-7.79a7 7 0 0 0-13.36 1.9A4 4 0 0 0 6 17a1 1 0 0 0 0-2a2 2 0 0 1 0-4a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.61a1 1 0 0 0 .78.66A3 3 0 0 1 17 15a1 1 0 0 0 0 2a5 5 0 0 0 1.42-9.79"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThunderstormMoon(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.58 16.5h-1.26l.86-1.5a1 1 0 1 0-1.73-1l-1.73 3a1 1 0 0 0 .86 1.5h1.27L9 20a1 1 0 0 0 1.74 1l1.73-3a1 1 0 0 0 0-1a1 1 0 0 0-.89-.5M21.7 7.57a1 1 0 0 0-.93-.26a3.2 3.2 0 0 1-.66.08a3 3 0 0 1-3-3a3 3 0 0 1 .08-.65A1 1 0 0 0 16 2.53a4.93 4.93 0 0 0-3.83 4.21a6.24 6.24 0 0 0-1.67-.24a6 6 0 0 0-5.94 5.13a3.49 3.49 0 0 0-.34 6.62a1 1 0 1 0 .72-1.86a1.5 1.5 0 0 1 .56-2.89a1 1 0 0 0 1-1a4 4 0 0 1 4-4a3.92 3.92 0 0 1 2.18.66a4 4 0 0 1 1.57 2a1 1 0 0 0 .78.67a2.33 2.33 0 0 1 .25 4.53a1 1 0 0 0 .27 2a.84.84 0 0 0 .27 0A4.33 4.33 0 0 0 19 14.17a4.23 4.23 0 0 0-.49-2A4.94 4.94 0 0 0 22 8.5a1 1 0 0 0-.3-.93m-4.59 2.82a2.61 2.61 0 0 1-.42 0A4.6 4.6 0 0 0 16 10a6 6 0 0 0-1.82-2.28v-.37a3 3 0 0 1 1.05-2.28a5 5 0 0 0 4.23 4.23a3 3 0 0 1-2.35 1.09"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThunderstormSun(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.58 16.5h-1.26l.86-1.5a1 1 0 1 0-1.73-1l-1.73 3a1 1 0 0 0 .86 1.5h1.27L9 20a1 1 0 0 0 1.74 1l1.73-3a1 1 0 0 0 0-1a1 1 0 0 0-.89-.5m9.42-9h-.8a4.25 4.25 0 0 0-.52-1.27l.56-.56a1 1 0 0 0-1.41-1.41l-.56.56A4.25 4.25 0 0 0 17 4.3v-.8a1 1 0 0 0-2 0v.8a4.25 4.25 0 0 0-1.27.52l-.56-.56a1 1 0 0 0-1.41 1.41l.56.57c-.09.15-.16.32-.24.48a5.85 5.85 0 0 0-1.58-.22a6 6 0 0 0-5.94 5.13a3.49 3.49 0 0 0-.34 6.62a1 1 0 1 0 .72-1.86a1.5 1.5 0 0 1 .56-2.89a1 1 0 0 0 1-1a4 4 0 0 1 7.78-1.29a1 1 0 0 0 .78.67a2.33 2.33 0 0 1 .25 4.53a1 1 0 0 0 .27 2a.84.84 0 0 0 .27 0a4.3 4.3 0 0 0 2.85-5.72l.13.13a1 1 0 0 0 .7.29a1 1 0 0 0 .71-.29a1 1 0 0 0 0-1.41l-.56-.56a4.25 4.25 0 0 0 .52-1.35h.8a1 1 0 0 0 0-2m-3.34 2.65a2.28 2.28 0 0 1-.6.41A4.17 4.17 0 0 0 16 10a6.12 6.12 0 0 0-2.09-2.49a2.42 2.42 0 0 1 .46-.7a2.43 2.43 0 0 1 3.3 0a2.37 2.37 0 0 1 0 3.3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ticket(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 10a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0v-2a1 1 0 0 0-1-1m12 1a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1a1 1 0 0 1 0 2a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1a1 1 0 0 1 0-2m-1-1.82a3 3 0 0 0 0 5.64V17H10a1 1 0 0 0-2 0H4v-2.18a3 3 0 0 0 0-5.64V7h4a1 1 0 0 0 2 0h10Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Times(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.41 12l4.3-4.29a1 1 0 1 0-1.42-1.42L12 10.59l-4.29-4.3a1 1 0 0 0-1.42 1.42l4.3 4.29l-4.3 4.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l4.29-4.3l4.29 4.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimesCircle(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.71 8.29a1 1 0 0 0-1.42 0L12 10.59l-2.29-2.3a1 1 0 0 0-1.42 1.42l2.3 2.29l-2.3 2.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l2.29-2.3l2.29 2.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L13.41 12l2.3-2.29a1 1 0 0 0 0-1.42m3.36-3.36A10 10 0 1 0 4.93 19.07A10 10 0 1 0 19.07 4.93m-1.41 12.73A8 8 0 1 1 20 12a7.95 7.95 0 0 1-2.34 5.66"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimesSquare(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.71 8.29a1 1 0 0 0-1.42 0L12 10.59l-2.29-2.3a1 1 0 0 0-1.42 1.42l2.3 2.29l-2.3 2.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l2.29-2.3l2.29 2.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L13.41 12l2.3-2.29a1 1 0 0 0 0-1.42M19 2H5a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3m1 17a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToggleOff(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 8.5a3.5 3.5 0 1 0 3.5 3.5A3.5 3.5 0 0 0 8 8.5m0 5A1.5 1.5 0 1 1 9.5 12A1.5 1.5 0 0 1 8 13.5M16 5H8a7 7 0 0 0 0 14h8a7 7 0 0 0 0-14m0 12H8A5 5 0 0 1 8 7h8a5 5 0 0 1 0 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToggleOn(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 8.5a3.5 3.5 0 1 0 3.5 3.5A3.5 3.5 0 0 0 16 8.5m0 5a1.5 1.5 0 1 1 1.5-1.5a1.5 1.5 0 0 1-1.5 1.5M16 5H8a7 7 0 0 0 0 14h8a7 7 0 0 0 0-14m0 12H8A5 5 0 0 1 8 7h8a5 5 0 0 1 0 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToiletPaper(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.26 20.357a7.523 7.523 0 0 1-1.76-4.905v-7.46c0-3.308-2.243-6-5-6h-9c-2.757 0-5 2.692-5 6s2.243 6 5 6h3v1.46a9.527 9.527 0 0 0 2.24 6.206a1.001 1.001 0 0 0 .76.35h9a1 1 0 0 0 .76-1.651M6.5 11.993c-1.654 0-3-1.795-3-4s1.346-4 3-4s3 1.794 3 4s-1.346 4-3 4m6.479 8.014a7.58 7.58 0 0 1-1.479-4.555v-7c0-.028-.014-.052-.016-.08c.007-.126.016-.251.016-.38a6.68 6.68 0 0 0-1.284-4H15.5c1.654 0 3 1.795 3 4v7.46a9.71 9.71 0 0 0 1.118 4.555ZM6.5 6.743a1.146 1.146 0 0 0-1 1.25a1.146 1.146 0 0 0 1 1.25a1.146 1.146 0 0 0 1-1.25a1.146 1.146 0 0 0-1-1.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TopArrowFromTop(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 20H5a1 1 0 0 0 0 2h14a1 1 0 0 0 0-2M8.71 7.71L11 5.41V17a1 1 0 0 0 2 0V5.41l2.29 2.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-4-4a1 1 0 0 0-.33-.21a1 1 0 0 0-.76 0a1 1 0 0 0-.33.21l-4 4a1 1 0 1 0 1.42 1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TopArrowToTop(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.71 6.29a1 1 0 0 0-.33-.21a1 1 0 0 0-.76 0a1 1 0 0 0-.33.21l-4 4a1 1 0 1 0 1.42 1.42L11 9.41V21a1 1 0 0 0 2 0V9.41l2.29 2.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM19 2H5a1 1 0 0 0 0 2h14a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tornado(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 21H8a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2m1-4H6a1 1 0 0 0 0 2h5a1 1 0 0 0 0-2m7-15a1 1 0 0 0-1-1H3a1 1 0 0 0 0 2h14a1 1 0 0 0 1-1m3 3H6a1 1 0 0 0 0 2h15a1 1 0 0 0 0-2m-2 4h-9a1 1 0 0 0 0 2h9a1 1 0 0 0 0-2m-5 4H8a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trademark(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.5 7h-6a1 1 0 0 0 0 2h2v7a1 1 0 0 0 2 0V9h2a1 1 0 0 0 0-2m11.92.62a1 1 0 0 0-.54-.54a1 1 0 0 0-.76 0a1 1 0 0 0-.33.21L17 10.09l-2.79-2.8a1 1 0 0 0-.33-.21a1 1 0 0 0-.76 0a1 1 0 0 0-.54.54a1 1 0 0 0-.08.38v8a1 1 0 0 0 2 0v-5.59l1.79 1.8a1 1 0 0 0 1.42 0l1.79-1.8V16a1 1 0 0 0 2 0V8a1 1 0 0 0-.08-.38"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrademarkCircle(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 9H7a1 1 0 0 0 0 2h.5v3a1 1 0 0 0 2 0v-3h.5a1 1 0 0 0 0-2m7.38.08a1 1 0 0 0-1.09.21L15 10.59l-1.29-1.3a1 1 0 0 0-1.09-.21A1 1 0 0 0 12 10v4a1 1 0 0 0 2 0v-1.59l.29.3a1 1 0 0 0 1.42 0l.29-.3V14a1 1 0 0 0 2 0v-4a1 1 0 0 0-.62-.92M12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrafficBarrier(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 5h-1V4a1 1 0 0 0-2 0v1H6V4a1 1 0 0 0-2 0v1H3a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h1v7a1 1 0 0 0 2 0v-7h12v7a1 1 0 0 0 2 0v-7h1a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1M4 9.52V7h2.52ZM5.34 11l4-4h3.33l-4 4Zm6.15 0l4-4h3.18l-4 4ZM20 11h-2.51L20 8.49Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrafficLight(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 5.5A1.5 1.5 0 1 0 13.5 7A1.5 1.5 0 0 0 12 5.5m10 3h-.54a6 6 0 0 0 1.54-4a1 1 0 0 0-1-1h-4.18a3 3 0 0 0-2.82-2H9a3 3 0 0 0-2.82 2H2a1 1 0 0 0-1 1a6 6 0 0 0 1.54 4H2a1 1 0 0 0-1 1a6 6 0 0 0 1.54 4H2a1 1 0 0 0-1 1a6 6 0 0 0 5.16 5.93A3 3 0 0 0 9 22.5h6a3 3 0 0 0 2.84-2.07A6 6 0 0 0 23 14.5a1 1 0 0 0-1-1h-.54a6 6 0 0 0 1.54-4a1 1 0 0 0-1-1M6 18.37a4 4 0 0 1-2.87-2.87H6Zm0-5a4 4 0 0 1-2.87-2.87H6Zm0-5a4.09 4.09 0 0 1-1.83-1a4.09 4.09 0 0 1-1-1.83H6ZM16 19.5a1 1 0 0 1-1 1H9a1 1 0 0 1-1-1v-15a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1Zm3.83-2.17a4.09 4.09 0 0 1-1.83 1V15.5h2.87a4.09 4.09 0 0 1-1.04 1.83m0-5a4.09 4.09 0 0 1-1.83 1V10.5h2.87a4.09 4.09 0 0 1-1.04 1.83m0-5a4.09 4.09 0 0 1-1.83 1V5.5h2.87a4.09 4.09 0 0 1-1.04 1.83M12 15.5a1.5 1.5 0 1 0 1.5 1.5a1.5 1.5 0 0 0-1.5-1.5m0-5a1.5 1.5 0 1 0 1.5 1.5a1.5 1.5 0 0 0-1.5-1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Transaction(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 2H10a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3m1 10a1 1 0 0 1-1 1H10a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1Zm-3.5-4a1.49 1.49 0 0 0-1 .39a1.5 1.5 0 1 0 0 2.22a1.5 1.5 0 1 0 1-2.61M16 17a1 1 0 0 0-1 1v1a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1v-4h1a1 1 0 0 0 0-2H3v-1a1 1 0 0 1 1-1a1 1 0 0 0 0-2a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-1a1 1 0 0 0-1-1M6 18h1a1 1 0 0 0 0-2H6a1 1 0 0 0 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 6h-4V5a3 3 0 0 0-3-3h-2a3 3 0 0 0-3 3v1H4a1 1 0 0 0 0 2h1v11a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V8h1a1 1 0 0 0 0-2M10 5a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v1h-4Zm7 14a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1V8h10Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrashAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 18a1 1 0 0 0 1-1v-6a1 1 0 0 0-2 0v6a1 1 0 0 0 1 1M20 6h-4V5a3 3 0 0 0-3-3h-2a3 3 0 0 0-3 3v1H4a1 1 0 0 0 0 2h1v11a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V8h1a1 1 0 0 0 0-2M10 5a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v1h-4Zm7 14a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1V8h10Zm-3-1a1 1 0 0 0 1-1v-6a1 1 0 0 0-2 0v6a1 1 0 0 0 1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trees(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.5 5a4.86 4.86 0 0 0-2.5.69A6 6 0 0 0 2.5 8v4a6 6 0 0 0 5 5.91V21a1 1 0 0 0 2 0v-3.09a6.08 6.08 0 0 0 2.78-1.26a5 5 0 0 0 3.22 2.25V21a1 1 0 0 0 2 0v-2.1a5 5 0 0 0 4-4.9v-4a5 5 0 0 0-5-5m-5 5v4a5.23 5.23 0 0 0 .06.57a4 4 0 0 1-2.06 1.3V13a1 1 0 0 0-2 0v2.86a4 4 0 0 1-3-3.86V8a4 4 0 0 1 6.83-2.84a3.94 3.94 0 0 1 1.06 2A5 5 0 0 0 11.5 10m8 4a3 3 0 0 1-2 2.82V13a1 1 0 0 0-2 0v3.82a3 3 0 0 1-2-2.82v-4a3 3 0 0 1 6 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Triangle(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.87 19.29l-9-15.58a1 1 0 0 0-1.74 0l-9 15.58a1 1 0 0 0 0 1a1 1 0 0 0 .87.5h18a1 1 0 0 0 .87-.5a1 1 0 0 0 0-1m-17.14-.5L12 6.21l7.27 12.58Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trophy(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 4h-3V3a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1v1H3a1 1 0 0 0-1 1v3a4 4 0 0 0 4 4h1.54A6 6 0 0 0 11 13.91V16h-1a3 3 0 0 0-3 3v2a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1v-2a3 3 0 0 0-3-3h-1v-2.09A6 6 0 0 0 16.46 12H18a4 4 0 0 0 4-4V5a1 1 0 0 0-1-1M6 10a2 2 0 0 1-2-2V6h2v2a6 6 0 0 0 .35 2Zm8 8a1 1 0 0 1 1 1v1H9v-1a1 1 0 0 1 1-1Zm2-10a4 4 0 0 1-8 0V4h8Zm4 0a2 2 0 0 1-2 2h-.35A6 6 0 0 0 18 8V6h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trowel(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.12 2.88a3.08 3.08 0 0 0-4.24 0l-2.42 2.41a3 3 0 0 0-.57 3.41l-2.15 2.15l-2-2a3 3 0 0 0-5 1.17l-2.66 8a3 3 0 0 0 .72 3.07A3 3 0 0 0 5 22a2.87 2.87 0 0 0 1-.16l8-2.66a3 3 0 0 0 1.17-5l-2-2l2.15-2.15a3 3 0 0 0 3.41-.57l2.41-2.42a3 3 0 0 0-.02-4.16M13.7 15.63a1 1 0 0 1-.4 1.65L5.32 20a1 1 0 0 1-1-.25a1 1 0 0 1-.25-1l2.67-8a1 1 0 0 1 .7-.75a1.07 1.07 0 0 1 .23 0a1 1 0 0 1 .7.29l2 2l-.87.86a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l.86-.87Zm6-9.92l-2.41 2.41a1 1 0 0 1-1.7-.71a1 1 0 0 1 .29-.7l2.41-2.42a1 1 0 0 1 1.42 0a1 1 0 0 1 0 1.42Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Truck(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 12.5v5a1 1 0 0 0 1 1h1a3 3 0 0 0 6 0h6a3 3 0 0 0 6 0h1a1 1 0 0 0 1-1v-12a3 3 0 0 0-3-3h-9a3 3 0 0 0-3 3v2H6a3 3 0 0 0-2.4 1.2l-2.4 3.2a.61.61 0 0 0-.07.14l-.06.11a1 1 0 0 0-.07.35m16 6a1 1 0 1 1 1 1a1 1 0 0 1-1-1m-7-13a1 1 0 0 1 1-1h9a1 1 0 0 1 1 1v11h-.78a3 3 0 0 0-4.44 0H10Zm-2 6H4l1.2-1.6a1 1 0 0 1 .8-.4h2Zm-3 7a1 1 0 1 1 1 1a1 1 0 0 1-1-1m-2-5h5v2.78a3 3 0 0 0-4.22.22H3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TruckLoading(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 16h-2.18a3 3 0 0 0 .18-1v-5a3 3 0 0 0-3-3h-6a3 3 0 0 0-3 3v5a3 3 0 0 0 .18 1H7a1 1 0 0 1-1-1V5a3 3 0 0 0-3-3H2a1 1 0 0 0 0 2h1a1 1 0 0 1 1 1v10a3 3 0 0 0 2.22 2.88a3 3 0 1 0 5.6.12h3.36a3 3 0 1 0 5.64 0H22a1 1 0 0 0 0-2M9 20a1 1 0 1 1 1-1a1 1 0 0 1-1 1m2-4a1 1 0 0 1-1-1v-5a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v5a1 1 0 0 1-1 1Zm7 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tumblr(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.79 18a4.65 4.65 0 0 1-1.62.35a1.75 1.75 0 0 1-1.92-2v-6.23h4v-3h-4V2h-2.92a.15.15 0 0 0-.14.15a6.11 6.11 0 0 1-3.94 5.39v2.58h2v6.54c0 2.23 1.65 5.41 6 5.34c1.47 0 3.11-.64 3.47-1.17Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TumblrAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m17.434 20.508l-.962-2.853a1 1 0 0 0-1.38-.583a3.763 3.763 0 0 1-1.208.249a.937.937 0 0 1-.66-.165a1.2 1.2 0 0 1-.239-.808V12h3.017a1 1 0 0 0 1-1V7.095a1 1 0 0 0-1-1H13V2a1 1 0 0 0-1-1H9.07a1.148 1.148 0 0 0-1.137 1.04a5.093 5.093 0 0 1-3.28 4.558a1 1 0 0 0-.662.94v3.585a1 1 0 0 0 1 1h1.025v4.535a6.411 6.411 0 0 0 1.886 4.478A6.905 6.905 0 0 0 12.877 23h.163c1.546-.026 3.618-.648 4.273-1.608a.998.998 0 0 0 .12-.883m-4.427.49a4.87 4.87 0 0 1-3.702-1.288a4.37 4.37 0 0 1-1.29-3.052v-5.535a1 1 0 0 0-1-1H5.992V8.206A6.954 6.954 0 0 0 9.81 3H11v4.095a1 1 0 0 0 1 1h3.002V10h-3.017a1 1 0 0 0-1 1v5.365a3.077 3.077 0 0 0 .857 2.235a2.767 2.767 0 0 0 2.096.72a5.908 5.908 0 0 0 .947-.113l.425 1.26a5.909 5.909 0 0 1-2.303.531"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TumblrSquare(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 2v20h20V2Zm11.57 15.85a3.4 3.4 0 0 1-3.75-3.33v-4.08H8.56V8.83A3.83 3.83 0 0 0 11 5.47a.09.09 0 0 1 .09-.09h1.82v3.17h2.5v1.89h-2.49v3.89a1.11 1.11 0 0 0 1.2 1.23a3 3 0 0 0 1-.22l.6 1.78a3.34 3.34 0 0 1-2.15.73"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TvRetro(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 6h-3.59l2.3-2.29a1 1 0 1 0-1.42-1.42L12 5.54l-1.17-2a1 1 0 1 0-1.74 1L10 6H6a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3v1a1 1 0 0 0 2 0v-1h8v1a1 1 0 0 0 2 0v-1a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3m1 11a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TvRetroSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.62 7.92A1 1 0 0 0 12 8h6a1 1 0 0 1 1 1v5.34a1 1 0 1 0 2 0V9a3 3 0 0 0-3-3h-3.59l2.3-2.29a1 1 0 1 0-1.42-1.42l-4 4a1 1 0 0 0-.21.33a1 1 0 0 0 0 .76a1 1 0 0 0 .54.54m10.09 12.37l-18-18a1 1 0 0 0-1.42 1.42l2.54 2.53A3 3 0 0 0 3 9v8a3 3 0 0 0 3 3v1a1 1 0 0 0 2 0v-1h8v1a1 1 0 0 0 2 0v-1a3.07 3.07 0 0 0 .53-.06l1.76 1.77a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42M6 18a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1h.59l10 10Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwelvePlus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 9a1 1 0 0 0 2 0a1 1 0 0 1 1-1h1a1 1 0 0 1 1 1v1a1 1 0 0 1-1 1h-1a3 3 0 0 0-3 3v3a1 1 0 0 0 1 1h5a1 1 0 0 0 0-2h-4v-2a1 1 0 0 1 1-1h1a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3h-1a3 3 0 0 0-3 3m7-4h1v1a1 1 0 0 0 2 0V5h1a1 1 0 0 0 0-2h-1V2a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2M7 7v10a1 1 0 0 0 2 0V7a1 1 0 0 0-2 0m14.6 2a1 1 0 0 0-.78 1.18a9 9 0 1 1-7-7a1 1 0 1 0 .4-2A10.8 10.8 0 0 0 12 1a11 11 0 1 0 11 11a10.8 10.8 0 0 0-.22-2.2A1 1 0 0 0 21.6 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwentyOnePlus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 18a1 1 0 0 0 1-1V7a1 1 0 0 0-2 0v10a1 1 0 0 0 1 1m1-13h1v1a1 1 0 0 0 2 0V5h1a1 1 0 0 0 0-2h-1V2a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2m3.6 4a1 1 0 0 0-.78 1.18a9 9 0 1 1-7-7a1 1 0 1 0 .4-2A10.8 10.8 0 0 0 12 1a11 11 0 1 0 11 11a10.8 10.8 0 0 0-.22-2.2A1 1 0 0 0 21.6 9M10 13h1a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3h-1a3 3 0 0 0-3 3a1 1 0 0 0 2 0a1 1 0 0 1 1-1h1a1 1 0 0 1 1 1v1a1 1 0 0 1-1 1h-1a3 3 0 0 0-3 3v3a1 1 0 0 0 1 1h5a1 1 0 0 0 0-2H9v-2a1 1 0 0 1 1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Twitter(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 5.8a8.49 8.49 0 0 1-2.36.64a4.13 4.13 0 0 0 1.81-2.27a8.21 8.21 0 0 1-2.61 1a4.1 4.1 0 0 0-7 3.74a11.64 11.64 0 0 1-8.45-4.29a4.16 4.16 0 0 0-.55 2.07a4.09 4.09 0 0 0 1.82 3.41a4.05 4.05 0 0 1-1.86-.51v.05a4.1 4.1 0 0 0 3.3 4a3.93 3.93 0 0 1-1.1.17a4.9 4.9 0 0 1-.77-.07a4.11 4.11 0 0 0 3.83 2.84A8.22 8.22 0 0 1 3 18.34a7.93 7.93 0 0 1-1-.06a11.57 11.57 0 0 0 6.29 1.85A11.59 11.59 0 0 0 20 8.45v-.53a8.43 8.43 0 0 0 2-2.12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwitterAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.991 3.95a1 1 0 0 0-1.51-.86a7.48 7.48 0 0 1-1.874.794a5.152 5.152 0 0 0-3.374-1.242a5.232 5.232 0 0 0-5.223 5.063a11.032 11.032 0 0 1-6.814-3.924a1.012 1.012 0 0 0-.857-.365a.999.999 0 0 0-.785.5a5.276 5.276 0 0 0-.242 4.769l-.002.001a1.041 1.041 0 0 0-.496.89a3.042 3.042 0 0 0 .027.439a5.185 5.185 0 0 0 1.568 3.312a.998.998 0 0 0-.066.77a5.204 5.204 0 0 0 2.362 2.922a7.465 7.465 0 0 1-3.59.448A1 1 0 0 0 1.45 19.3a12.942 12.942 0 0 0 7.01 2.061a12.788 12.788 0 0 0 12.465-9.363a12.822 12.822 0 0 0 .535-3.646l-.001-.2a5.77 5.77 0 0 0 1.532-4.202m-3.306 3.212a.995.995 0 0 0-.234.702c.01.165.009.331.009.488a10.824 10.824 0 0 1-.454 3.08a10.685 10.685 0 0 1-10.546 7.93a10.938 10.938 0 0 1-2.55-.301a9.48 9.48 0 0 0 2.942-1.564a1 1 0 0 0-.602-1.786a3.208 3.208 0 0 1-2.214-.935q.224-.042.445-.105a1 1 0 0 0-.08-1.943a3.198 3.198 0 0 1-2.25-1.726a5.3 5.3 0 0 0 .545.046a1.02 1.02 0 0 0 .984-.696a1 1 0 0 0-.4-1.137a3.196 3.196 0 0 1-1.425-2.673c0-.066.002-.133.006-.198a13.014 13.014 0 0 0 8.21 3.48a1.02 1.02 0 0 0 .817-.36a1 1 0 0 0 .206-.867a3.157 3.157 0 0 1-.087-.729a3.23 3.23 0 0 1 3.226-3.226a3.184 3.184 0 0 1 2.345 1.02a.993.993 0 0 0 .921.298a9.27 9.27 0 0 0 1.212-.322a6.681 6.681 0 0 1-1.026 1.524"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Umbrella(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 11.24a.22.22 0 0 1 0-.08v-.19c0-.23-.06-.46-.1-.69a.75.75 0 0 1 0-.16c-.05-.25-.12-.49-.19-.73a8.91 8.91 0 0 0-5.86-5.87h-.08c-.22-.07-.45-.13-.68-.18h-.18a5.21 5.21 0 0 0-.55-.08h-.24V3a1 1 0 0 0-2 0v.06a8.7 8.7 0 0 0-1 .18a4.71 4.71 0 0 0-.62.16h-.13c-.25.08-.48.17-.72.26a8.93 8.93 0 0 0-5.23 5.62a.31.31 0 0 0 0 .08a6.38 6.38 0 0 0-.19.72v.16q-.08.36-.12.75s0 .07 0 .11v.9a1 1 0 0 0 1 1h7v6a1 1 0 0 1-2 0a1 1 0 0 0-2 0a3 3 0 0 0 6 0v-6h7a1 1 0 0 0 1-1c-.11-.26-.11-.51-.11-.76M8 11H5.08v-.11c0-.04 0-.28.08-.41s0-.13 0-.19s.08-.32.13-.48v-.08A7 7 0 0 1 9.1 5.64A16.09 16.09 0 0 0 8 11m2 0c.19-3.91 1.44-6 2-6s1.79 2.09 2 6Zm6 0a16.09 16.09 0 0 0-1.1-5.36a7 7 0 0 1 3.73 4.12a4.61 4.61 0 0 1 .15.53a.83.83 0 0 0 0 .15c0 .14.06.29.08.43s0 .07 0 .11Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unamused(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14.66 13.56l-4.19 1.5A1 1 0 0 0 10.8 17a1 1 0 0 0 .34-.06l4.2-1.5a1 1 0 1 0-.68-1.88m-4-3a1 1 0 0 0 0-1.41a3.08 3.08 0 0 0-4.24 0a1 1 0 1 0 1.41 1.41a1 1 0 0 1 1.42 0a1 1 0 0 0 1.41 0m7-1.41a3.08 3.08 0 0 0-4.24 0a1 1 0 0 0 1.41 1.41a1 1 0 0 1 1.42 0a1 1 0 0 0 1.41 0a1 1 0 0 0-.04-1.43ZM12 2a10 10 0 1 0 10 10A10 10 0 0 0 12 2m0 18a8 8 0 1 1 8-8a8 8 0 0 1-8 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Underline(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 15.5a5 5 0 0 0 5-5v-5a1 1 0 0 0-2 0v5a3 3 0 0 1-6 0v-5a1 1 0 0 0-2 0v5a5 5 0 0 0 5 5m5 2H7a1 1 0 0 0 0 2h10a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func University(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 10a1 1 0 0 0 1-1V6a.999.999 0 0 0-.684-.948l-9-3a1.002 1.002 0 0 0-.632 0l-9 3A.999.999 0 0 0 2 6v3a1 1 0 0 0 1 1h1v7.184A2.995 2.995 0 0 0 2 20v2a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1v-2a2.995 2.995 0 0 0-2-2.816V10Zm-1 11H4v-1a1.001 1.001 0 0 1 1-1h14a1.001 1.001 0 0 1 1 1ZM6 17v-7h2v7Zm4 0v-7h4v7Zm6 0v-7h2v7ZM4 8V6.72l8-2.666l8 2.667V8Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unlock(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 9H9V7a3 3 0 0 1 5.12-2.13a3.08 3.08 0 0 1 .78 1.38a1 1 0 1 0 1.94-.5a5.09 5.09 0 0 0-1.31-2.29A5 5 0 0 0 7 7v2a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-7a3 3 0 0 0-3-3m1 10a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1v-7a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnlockAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 13a1 1 0 0 0-1 1v3a1 1 0 0 0 2 0v-3a1 1 0 0 0-1-1m5-4H9V7a3 3 0 0 1 5.12-2.13a3.08 3.08 0 0 1 .78 1.38a1 1 0 1 0 1.94-.5a5.09 5.09 0 0 0-1.31-2.29A5 5 0 0 0 7 7v2a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-7a3 3 0 0 0-3-3m1 10a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1v-7a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Upload(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.71 7.71L11 5.41V15a1 1 0 0 0 2 0V5.41l2.29 2.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-4-4a1 1 0 0 0-.33-.21a1 1 0 0 0-.76 0a1 1 0 0 0-.33.21l-4 4a1 1 0 1 0 1.42 1.42M21 12a1 1 0 0 0-1 1v6a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-6a1 1 0 0 0-2 0v6a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-6a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UploadAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.71 6.71L11 5.41V17a1 1 0 0 0 2 0V5.41l1.29 1.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-3-3a1 1 0 0 0-1.42 0l-3 3a1 1 0 0 0 1.42 1.42M18 9h-2a1 1 0 0 0 0 2h2a1 1 0 0 1 1 1v7a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1v-7a1 1 0 0 1 1-1h2a1 1 0 0 0 0-2H6a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-7a3 3 0 0 0-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UsdCircle(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 9h4a1 1 0 0 0 0-2h-2V6a1 1 0 0 0-2 0v1a3 3 0 0 0 0 6h2a1 1 0 0 1 0 2H9a1 1 0 0 0 0 2h2v1a1 1 0 0 0 2 0v-1a3 3 0 0 0 0-6h-2a1 1 0 0 1 0-2m1-8a11 11 0 1 0 11 11A11 11 0 0 0 12 1m0 20a9 9 0 1 1 9-9a9 9 0 0 1-9 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UsdSquare(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 9h4a1 1 0 0 0 0-2h-2V6a1 1 0 0 0-2 0v1a3 3 0 0 0 0 6h2a1 1 0 0 1 0 2H9a1 1 0 0 0 0 2h2v1a1 1 0 0 0 2 0v-1a3 3 0 0 0 0-6h-2a1 1 0 0 1 0-2m8-7H5a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3m1 17a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func User(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.71 12.71a6 6 0 1 0-7.42 0a10 10 0 0 0-6.22 8.18a1 1 0 0 0 2 .22a8 8 0 0 1 15.9 0a1 1 0 0 0 1 .89h.11a1 1 0 0 0 .88-1.1a10 10 0 0 0-6.25-8.19M12 12a4 4 0 1 1 4-4a4 4 0 0 1-4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserArrows(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.64 15.772a3.468 3.468 0 0 0 .86-2.272a3.5 3.5 0 0 0-7 0a3.468 3.468 0 0 0 .86 2.272A4.988 4.988 0 0 0 13 20a1 1 0 0 0 2 0a3 3 0 0 1 6 0a1 1 0 0 0 2 0a4.988 4.988 0 0 0-2.36-4.228M18 15a1.5 1.5 0 1 1 1.5-1.5A1.502 1.502 0 0 1 18 15M6.793 7.707l2 2a1 1 0 0 0 1.414-1.414L9.914 8h4.172l-.293.293a1 1 0 1 0 1.414 1.414l2-2a1.004 1.004 0 0 0 0-1.414l-2-2a1 1 0 0 0-1.414 1.414l.293.293H9.914l.293-.293a1 1 0 0 0-1.414-1.414l-2 2a1.004 1.004 0 0 0 0 1.414m1.847 8.065A3.468 3.468 0 0 0 9.5 13.5a3.5 3.5 0 0 0-7 0a3.468 3.468 0 0 0 .86 2.272A4.988 4.988 0 0 0 1 20a1 1 0 0 0 2 0a3 3 0 0 1 6 0a1 1 0 0 0 2 0a4.988 4.988 0 0 0-2.36-4.228M6 15a1.5 1.5 0 1 1 1.5-1.5A1.502 1.502 0 0 1 6 15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserCheck(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.3 12.22A4.92 4.92 0 0 0 15 8.5a5 5 0 0 0-10 0a4.92 4.92 0 0 0 1.7 3.72A8 8 0 0 0 2 19.5a1 1 0 0 0 2 0a6 6 0 0 1 12 0a1 1 0 0 0 2 0a8 8 0 0 0-4.7-7.28M10 11.5a3 3 0 1 1 3-3a3 3 0 0 1-3 3m11.71-2.37a1 1 0 0 0-1.42 0l-2 2l-.62-.63a1 1 0 0 0-1.42 0a1 1 0 0 0 0 1.41l1.34 1.34a1 1 0 0 0 1.41 0l2.67-2.67a1 1 0 0 0 .04-1.45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserCircle(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a10 10 0 0 0-7.35 16.76a10 10 0 0 0 14.7 0A10 10 0 0 0 12 2m0 18a8 8 0 0 1-5.55-2.25a6 6 0 0 1 11.1 0A8 8 0 0 1 12 20m-2-10a2 2 0 1 1 2 2a2 2 0 0 1-2-2m8.91 6A8 8 0 0 0 15 12.62a4 4 0 1 0-6 0A8 8 0 0 0 5.09 16A7.92 7.92 0 0 1 4 12a8 8 0 0 1 16 0a7.92 7.92 0 0 1-1.09 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserExclamation(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.55 12.22a4.92 4.92 0 0 0 1.7-3.72a5 5 0 0 0-10 0A4.92 4.92 0 0 0 8 12.22a8 8 0 0 0-4.7 7.28a1 1 0 0 0 2 0a6 6 0 0 1 12 0a1 1 0 0 0 2 0a8 8 0 0 0-4.75-7.28m-3.3-.72a3 3 0 1 1 3-3a3 3 0 0 1-3 3m8.5-5a1 1 0 0 0-1 1v2a1 1 0 0 0 2 0v-2a1 1 0 0 0-1-1M19 11.79a1.05 1.05 0 0 0-.29.71a1 1 0 0 0 .29.71a1.15 1.15 0 0 0 .33.21a.94.94 0 0 0 .76 0a.9.9 0 0 0 .54-.54a.84.84 0 0 0 .08-.38a1 1 0 0 0-1.71-.71"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserLocation(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.49 10.86a3.09 3.09 0 1 0-5 0a4.67 4.67 0 0 0-1.12 1A1 1 0 1 0 10 13.12a2.62 2.62 0 0 1 2.05-1a2.62 2.62 0 0 1 2.05 1a1 1 0 0 0 .78.37a1 1 0 0 0 .78-1.62a4.67 4.67 0 0 0-1.17-1.01M12 10.13A1.09 1.09 0 1 1 13.09 9A1.09 1.09 0 0 1 12 10.13m8.46-.5A8.5 8.5 0 0 0 7.3 3.36a8.56 8.56 0 0 0-3.76 6.27A8.46 8.46 0 0 0 6 16.46l5.3 5.31a1 1 0 0 0 1.42 0L18 16.46a8.46 8.46 0 0 0 2.46-6.83m-3.86 5.42l-4.6 4.6l-4.6-4.6a6.49 6.49 0 0 1-1.87-5.22A6.57 6.57 0 0 1 8.42 5a6.47 6.47 0 0 1 7.16 0a6.57 6.57 0 0 1 2.89 4.81a6.49 6.49 0 0 1-1.87 5.24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserMd(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.899 13.229l-.005-.002c-.063-.027-.124-.058-.188-.083A5.988 5.988 0 0 0 18 8.434a5.29 5.29 0 0 0-.045-.63a.946.946 0 0 0 .038-.122l.281-2.397a3.006 3.006 0 0 0-2.442-3.302l-.79-.143a16.931 16.931 0 0 0-6.083 0l-.791.143a3.006 3.006 0 0 0-2.442 3.302l.28 2.397a.946.946 0 0 0 .039.122a5.29 5.29 0 0 0-.045.63a5.988 5.988 0 0 0 2.294 4.71c-.064.025-.125.056-.188.083l-.005.002a9.948 9.948 0 0 0-6.035 8.097a1 1 0 0 0 1.988.217a7.948 7.948 0 0 1 4.216-6.185l3.023 3.023a1 1 0 0 0 1.414 0l3.023-3.023a7.948 7.948 0 0 1 4.216 6.185a1 1 0 0 0 .992.892a1.048 1.048 0 0 0 .11-.006a1 1 0 0 0 .886-1.103a9.948 9.948 0 0 0-6.036-8.097ZM7.712 5.051a1.002 1.002 0 0 1 .814-1.1l.79-.143a14.93 14.93 0 0 1 5.368 0l.79.143a1.002 1.002 0 0 1 .814 1.1l-.178 1.514H7.89ZM12 16.261l-1.65-1.651a7.85 7.85 0 0 1 3.3 0Zm0-3.826a4.005 4.005 0 0 1-3.998-3.87h7.996A4.005 4.005 0 0 1 12 12.435"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserMinus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 10.5h-4a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2m-7.7 1.72A4.92 4.92 0 0 0 15 8.5a5 5 0 0 0-10 0a4.92 4.92 0 0 0 1.7 3.72A8 8 0 0 0 2 19.5a1 1 0 0 0 2 0a6 6 0 0 1 12 0a1 1 0 0 0 2 0a8 8 0 0 0-4.7-7.28M10 11.5a3 3 0 1 1 3-3a3 3 0 0 1-3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserNurse(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.174 12.916c-.018-.008-.035-.017-.053-.024c-.138-.062-.274-.127-.415-.183a6 6 0 1 0-7.412 0c-.14.056-.277.121-.415.183l-.054.024a9.946 9.946 0 0 0-5.76 7.976a1 1 0 1 0 1.99.216A7.945 7.945 0 0 1 8.04 15.05l3.252 3.251a1 1 0 0 0 1.414 0l3.252-3.251a7.945 7.945 0 0 1 3.987 6.058a1 1 0 0 0 .992.892a1.048 1.048 0 0 0 .11-.006a1 1 0 0 0 .886-1.102a9.946 9.946 0 0 0-5.76-7.976ZM8.041 7.594a3.977 3.977 0 0 1 7.918 0ZM12 16.18l-1.937-1.937a7.834 7.834 0 0 1 3.874 0ZM12 12a4.003 4.003 0 0 1-3.664-2.406h7.328A4.003 4.003 0 0 1 12 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserPlus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 10.5h-1v-1a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2h1v1a1 1 0 0 0 2 0v-1h1a1 1 0 0 0 0-2m-7.7 1.72A4.92 4.92 0 0 0 15 8.5a5 5 0 0 0-10 0a4.92 4.92 0 0 0 1.7 3.72A8 8 0 0 0 2 19.5a1 1 0 0 0 2 0a6 6 0 0 1 12 0a1 1 0 0 0 2 0a8 8 0 0 0-4.7-7.28M10 11.5a3 3 0 1 1 3-3a3 3 0 0 1-3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserSquare(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.81 12.28a3.73 3.73 0 0 0 1-2.5a3.78 3.78 0 0 0-7.56 0a3.73 3.73 0 0 0 1 2.5A5.94 5.94 0 0 0 6 16.89a1 1 0 0 0 2 .22a4 4 0 0 1 7.94 0A1 1 0 0 0 17 18h.11a1 1 0 0 0 .88-1.1a5.94 5.94 0 0 0-3.18-4.62M12 11.56a1.78 1.78 0 1 1 1.78-1.78A1.78 1.78 0 0 1 12 11.56M19 2H5a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V5a3 3 0 0 0-3-3m1 17a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserTimes(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.3 12.22A4.92 4.92 0 0 0 15 8.5a5 5 0 0 0-10 0a4.92 4.92 0 0 0 1.7 3.72A8 8 0 0 0 2 19.5a1 1 0 0 0 2 0a6 6 0 0 1 12 0a1 1 0 0 0 2 0a8 8 0 0 0-4.7-7.28M10 11.5a3 3 0 1 1 3-3a3 3 0 0 1-3 3m10.91.5l.8-.79a1 1 0 0 0-1.42-1.42l-.79.8l-.79-.8a1 1 0 0 0-1.42 1.42l.8.79l-.8.79a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l.79-.8l.79.8a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UsersAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.3 12.22A4.92 4.92 0 0 0 14 8.5a5 5 0 0 0-10 0a4.92 4.92 0 0 0 1.7 3.72A8 8 0 0 0 1 19.5a1 1 0 0 0 2 0a6 6 0 0 1 12 0a1 1 0 0 0 2 0a8 8 0 0 0-4.7-7.28M9 11.5a3 3 0 1 1 3-3a3 3 0 0 1-3 3m9.74.32A5 5 0 0 0 15 3.5a1 1 0 0 0 0 2a3 3 0 0 1 3 3a3 3 0 0 1-1.5 2.59a1 1 0 0 0-.5.84a1 1 0 0 0 .45.86l.39.26l.13.07a7 7 0 0 1 4 6.38a1 1 0 0 0 2 0a9 9 0 0 0-4.23-7.68"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Utensils(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 2a1 1 0 0 0-1 1v5.46l-1 .67V3a1 1 0 0 0-2 0v6.13l-1-.67V3a1 1 0 0 0-2 0v6a1 1 0 0 0 .45.83L15 11.54V21a1 1 0 0 0 2 0v-9.46l2.55-1.71A1 1 0 0 0 20 9V3a1 1 0 0 0-1-1M9 2a5 5 0 0 0-5 5v6a1 1 0 0 0 1 1h3v7a1 1 0 0 0 2 0V3a1 1 0 0 0-1-1M8 12H6V7a3 3 0 0 1 2-2.83Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UtensilsAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.53 14.13a1 1 0 0 0-1.42 0a1 1 0 0 0 0 1.41l6.18 6.18a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.41Zm1.23-2.49a3 3 0 0 0 2.12-.88l2.83-2.83a1 1 0 0 0 0-1.41a1 1 0 0 0-1.42 0l-2.83 2.83a1 1 0 0 1-1.41 0l3.54-3.54a1 1 0 0 0 0-1.41a1 1 0 0 0-1.42 0l-3.53 3.53a1 1 0 0 1 0-1.41l2.82-2.83a1 1 0 1 0-1.41-1.41l-2.83 2.83a3 3 0 0 0 0 4.24L12 10.59l-1.56-1.54a4.16 4.16 0 0 0-.74-5C8.26 2.61 4.53 1 2.77 2.79S2.6 8.27 4 9.72A4.36 4.36 0 0 0 6.94 11h.14A3.88 3.88 0 0 0 9 10.46L10.57 12l-8.28 8.28a1 1 0 1 0 1.42 1.41l9-9l1.92-1.92a3 3 0 0 0 2.13.87m-8.33-3.2A1.93 1.93 0 0 1 7 9a2.26 2.26 0 0 1-1.54-.7C4.38 7.22 3.62 4.77 4.19 4.2a1 1 0 0 1 .66-.2a5.87 5.87 0 0 1 3.44 1.47a2.12 2.12 0 0 1 .14 2.97"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VectorSquare(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 16.18V7.82A3 3 0 1 0 16.18 4H7.82A3 3 0 1 0 4 7.82v8.36A3 3 0 1 0 7.82 20h8.36A3 3 0 1 0 20 16.18M19 4a1 1 0 1 1-1 1a1 1 0 0 1 1-1M5 4a1 1 0 1 1-1 1a1 1 0 0 1 1-1m0 16a1 1 0 1 1 1-1a1 1 0 0 1-1 1m11.18-2H7.82A3 3 0 0 0 6 16.18V7.82A3 3 0 0 0 7.82 6h8.36A3 3 0 0 0 18 7.82v8.36A3 3 0 0 0 16.18 18M19 20a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VectorSquareAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 7H8a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1V8a1 1 0 0 0-1-1m-1 8H9V9h6Zm6 3.28V5.72A2 2 0 1 0 18.28 3H5.72A2 2 0 1 0 3 5.72v12.56A2 2 0 1 0 5.72 21h12.56A2 2 0 1 0 21 18.28m-2 0a1.91 1.91 0 0 0-.72.72H5.72a1.91 1.91 0 0 0-.72-.72V5.72A1.91 1.91 0 0 0 5.72 5h12.56a1.91 1.91 0 0 0 .72.72Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Venus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 9a7 7 0 1 0-8 6.92V18h-1a1 1 0 0 0 0 2h1v1a1 1 0 0 0 2 0v-1h1a1 1 0 0 0 0-2h-1v-2.08A7 7 0 0 0 19 9m-7 5a5 5 0 1 1 5-5a5 5 0 0 1-5 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VerticalAlignBottom(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 20h-3V9a1 1 0 0 0-1-1h-3V3a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1v17H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2m-9 0H8V4h4Zm4 0h-2V10h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VerticalAlignCenter(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 11h-3V6a1 1 0 0 0-1-1h-3V3a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1v8H3a1 1 0 0 0 0 2h3v8a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1v-2h3a1 1 0 0 0 1-1v-5h3a1 1 0 0 0 0-2m-9 9H8V4h4Zm4-3h-2V7h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VerticalAlignTop(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2H3a1 1 0 0 0 0 2h3v17a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1v-5h3a1 1 0 0 0 1-1V4h3a1 1 0 0 0 0-2m-9 18H8V4h4Zm4-6h-2V4h2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VerticalDistributeBottom(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 18h-1v-5a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v5H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2m-3 0H6v-4h12ZM3 10h18a1 1 0 0 0 0-2h-2V5a1 1 0 0 0-1-1H6a1 1 0 0 0-1 1v3H3a1 1 0 0 0 0 2m4-4h10v2H7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VerticalDistributionCenter(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 15h-1v-2a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v2H3a1 1 0 0 0 0 2h1v2a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1v-2h1a1 1 0 0 0 0-2m-3 3H6v-4h12ZM3 8h2v1a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V8h2a1 1 0 0 0 0-2h-2V5a1 1 0 0 0-1-1H6a1 1 0 0 0-1 1v1H3a1 1 0 0 0 0 2m4-2h10v2H7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VerticalDistributionTop(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 6h1v5a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V6h1a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2m3 0h12v4H6Zm15 8H3a1 1 0 0 0 0 2h2v3a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-3h2a1 1 0 0 0 0-2m-4 4H7v-2h10Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Video(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.53 7.15a1 1 0 0 0-1 0L17 8.89A3 3 0 0 0 14 6H5a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3h9a3 3 0 0 0 3-2.89l3.56 1.78A1 1 0 0 0 21 17a1 1 0 0 0 .53-.15A1 1 0 0 0 22 16V8a1 1 0 0 0-.47-.85M15 15a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1h9a1 1 0 0 1 1 1Zm5-.62l-3-1.5v-1.76l3-1.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoQuestion(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.21 17.29a1.047 1.047 0 0 0-1.42 0a1.027 1.027 0 0 0-.21.33a.942.942 0 0 0 0 .76a1.154 1.154 0 0 0 .21.33A1 1 0 0 0 10.5 18a1 1 0 0 0-.08-.38a1.152 1.152 0 0 0-.21-.33M9.5 9a3.01 3.01 0 0 0-2.598 1.5a1 1 0 1 0 1.73 1A1 1 0 1 1 9.5 13a1 1 0 0 0-1 1v1a1 1 0 0 0 2 0v-.184A2.993 2.993 0 0 0 9.5 9m12.025-2.85a.999.999 0 0 0-.972-.045l-3.564 1.782A2.998 2.998 0 0 0 14 5H5a3.003 3.003 0 0 0-3 3v6a3.003 3.003 0 0 0 3 3h1a1 1 0 0 0 0-2H5a1.001 1.001 0 0 1-1-1V8a1.001 1.001 0 0 1 1-1h9a1.001 1.001 0 0 1 1 1v6a1.001 1.001 0 0 1-1 1h-.5a1 1 0 0 0 0 2h.5a2.998 2.998 0 0 0 2.989-2.888l3.564 1.782A1 1 0 0 0 22 15V7a1 1 0 0 0-.475-.85M20 13.381l-3-1.5v-1.764l3-1.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7.71 6.29l-4-4a1 1 0 0 0-1.42 1.42L4.62 6A3 3 0 0 0 2 9v6a3 3 0 0 0 3 3h9a3 3 0 0 0 1.9-.69l4.39 4.4a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM14 16H5a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1h1.59l7.87 7.88A1 1 0 0 1 14 16m7.53-8.85a1 1 0 0 0-1 0L17 8.89A3 3 0 0 0 14 6h-1.34a1 1 0 0 0 0 2H14a1 1 0 0 1 1 1v1.5a1.62 1.62 0 0 0 0 .19a.65.65 0 0 0 .05.2s.05.06.07.1a1 1 0 0 0 .15.21s.1.06.15.1l.17.11a.85.85 0 0 0 .23 0a.7.7 0 0 0 .14 0a1.62 1.62 0 0 0 .19 0a.65.65 0 0 0 .2-.05L20 9.62v5.72a1 1 0 1 0 2 0V8a1 1 0 0 0-.47-.85"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VirusSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.17 5.236a7.514 7.514 0 0 1 .83-.158v1.27a1 1 0 1 0 2 0V5.071a6.946 6.946 0 0 1 3.95 1.98v.001l.001.001a6.987 6.987 0 0 1 1.97 3.949H18a1 1 0 0 0 0 2h.922a7.65 7.65 0 0 1-.16.827a1 1 0 0 0 .718 1.217a.983.983 0 0 0 .25.032a1 1 0 0 0 .968-.75a9.594 9.594 0 0 0 .236-1.326H22a1 1 0 0 0 0-2h-1.06a8.933 8.933 0 0 0-1.912-4.614l.75-.75a1 1 0 0 0-1.414-1.414l-.753.753A8.911 8.911 0 0 0 13 3.067V1.999a1 1 0 1 0-2 0v1.066a9.473 9.473 0 0 0-1.33.236a1 1 0 0 0 .5 1.936ZM14 9.002a1 1 0 1 0 1-1a1 1 0 0 0-1 1m4.377 7.963l-.007-.011l-.012-.008L2.707 1.295a1 1 0 0 0-1.414 1.414l3.679 3.679a8.932 8.932 0 0 0-1.913 4.614H2a1 1 0 0 0 0 2h1.06a8.948 8.948 0 0 0 1.911 4.615l-.75.75a1 1 0 1 0 1.415 1.413l.75-.75A8.946 8.946 0 0 0 11 20.94v1.063a1 1 0 0 0 2 0V20.94a8.946 8.946 0 0 0 4.614-1.909l3.679 3.679a1 1 0 0 0 1.414-1.414ZM13 18.92v-1.917a1 1 0 0 0-2 0v1.917a6.986 6.986 0 0 1-3.945-1.96l-.005-.007l-.007-.005a6.994 6.994 0 0 1-1.963-3.945H6a1 1 0 0 0 0-2h-.921a6.943 6.943 0 0 1 1.32-3.187L8.253 9.67a1.5 1.5 0 0 0 2.08 2.08l5.853 5.853A6.957 6.957 0 0 1 13 18.919Zm-4-4.917a1 1 0 1 0 1 1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VisualStudio(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.85 2L9.09 9.77l-4.9-3.86L2.05 7v10l2.15 1.09l4.93-3.85L16.87 22L22 19.93V4ZM4.37 14.3V9.65l2.44 2.43Zm12.33 1.29L12.05 12l4.65-3.59Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vk(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.073 2H8.938C3.332 2 2 3.333 2 8.927v6.136C2 20.667 3.323 22 8.927 22h6.136C20.667 22 22 20.677 22 15.073V8.938C22 3.332 20.677 2 15.073 2m3.073 14.27h-1.459c-.552 0-.718-.447-1.708-1.437c-.864-.833-1.229-.937-1.448-.937c-.302 0-.385.083-.385.5v1.312c0 .355-.115.563-1.042.563a5.692 5.692 0 0 1-4.448-2.667a11.626 11.626 0 0 1-2.302-4.833c0-.219.083-.417.5-.417h1.459c.375 0 .51.167.656.552c.708 2.084 1.916 3.896 2.406 3.896c.188 0 .27-.083.27-.552v-2.146c-.062-.979-.582-1.062-.582-1.416a.36.36 0 0 1 .374-.334h2.292c.313 0 .417.156.417.531v2.896c0 .313.135.417.229.417c.188 0 .333-.104.677-.448a11.999 11.999 0 0 0 1.792-2.98a.628.628 0 0 1 .635-.416h1.459c.437 0 .53.219.437.531a18.205 18.205 0 0 1-1.958 3.365c-.157.24-.22.365 0 .646c.145.219.656.646 1 1.052a6.486 6.486 0 0 1 1.229 1.708c.125.406-.084.615-.5.615"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VkAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.744 13.53a5.519 5.519 0 0 0-.978-.837a7.873 7.873 0 0 0 2.859-4.471a1 1 0 1 0-1.95-.444a5.86 5.86 0 0 1-3.02 3.903V8.002a.958.958 0 0 0-.044-.216a1.024 1.024 0 0 0-.035-.171a.991.991 0 0 0-.167-.25c-.018-.02-.027-.048-.046-.068a1.001 1.001 0 0 0-.246-.167c-.026-.014-.045-.037-.072-.048a1.02 1.02 0 0 0-.236-.049a.945.945 0 0 0-.152-.031L11.003 7H11a.995.995 0 0 0-.35 1.929v4.89a11.307 11.307 0 0 1-3.01-5.984a1 1 0 1 0-1.972.334a13.334 13.334 0 0 0 5.4 8.644a.986.986 0 0 0 .128.064a.94.94 0 0 0 .108.054a.994.994 0 0 0 .35.071a.983.983 0 0 0 .424-.102c.01-.005.021-.002.031-.007a.993.993 0 0 0 .24-.176c.015-.014.024-.03.038-.045a.988.988 0 0 0 .16-.237a.952.952 0 0 0 .037-.087a.988.988 0 0 0 .07-.346v-2.126a3.551 3.551 0 0 1 1.616 1.005l1.647 1.797a1 1 0 1 0 1.474-1.352ZM15.073 1H8.938C2.78 1 1 2.778 1 8.927v6.136C1 21.22 2.778 23 8.927 23h6.136C21.22 23 23 21.222 23 15.073V8.938C23 2.78 21.222 1 15.073 1M21 15.073c0 5.04-.888 5.927-5.937 5.927H8.927C3.887 21 3 20.112 3 15.063V8.927C3 3.887 3.888 3 8.938 3h6.135C20.113 3 21 3.888 21 8.938Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Voicemail(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 8a4 4 0 0 0-4 4a3.91 3.91 0 0 0 .56 2H9.44a3.91 3.91 0 0 0 .56-2a4 4 0 1 0-4 4h12a4 4 0 0 0 0-8M6 14a2 2 0 1 1 2-2a2 2 0 0 1-2 2m12 0a2 2 0 1 1 2-2a2 2 0 0 1-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VoicemailRectangle(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 4H4a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h16a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3m1 13a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1Zm-5-8a3 3 0 0 0-2.82 4h-2.36A3 3 0 1 0 8 15h8a3 3 0 0 0 0-6m-8 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1m8 0a1 1 0 1 1 1-1a1 1 0 0 1-1 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Volleyball(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.59 6.9a.86.86 0 0 0-.07-.1A10 10 0 0 0 7.6 3h-.07a10 10 0 0 0-1 17.19l.33.2l.1.07A9.93 9.93 0 0 0 12 22h.21a10 10 0 0 0 8.38-15.1M19 8.06a7.64 7.64 0 0 1 .65 1.46a10 10 0 0 0-3-.49a.81.81 0 0 0-.31 0a9.78 9.78 0 0 0-3.58.73a7.85 7.85 0 0 1-1.84-1.6a8.16 8.16 0 0 1 8.08-.1M12 4a7.86 7.86 0 0 1 4 1.07A7.77 7.77 0 0 0 15 5a10 10 0 0 0-5.2 1.47a8 8 0 0 1-.64-1.94A7.92 7.92 0 0 1 12 4M6 6.71a8.26 8.26 0 0 1 1.33-1.19A9.9 9.9 0 0 0 12 11.61a7.89 7.89 0 0 1-.77 2.88A8 8 0 0 1 6 7zM4 12a8.1 8.1 0 0 1 .36-2.37a10 10 0 0 0 5.7 6.56a7.84 7.84 0 0 1-2.93 2.14A8 8 0 0 1 4 12m7.86 8a7.8 7.8 0 0 1-2.61-.49a9.94 9.94 0 0 0 3.23-3.22A10 10 0 0 0 14 11.41a7.71 7.71 0 0 1 1.78-.36A8 8 0 0 1 11.86 20m4.22-1.12A9.94 9.94 0 0 0 18 13a10.69 10.69 0 0 0-.18-1.88a8.34 8.34 0 0 1 2.17.7V12a8 8 0 0 1-3.91 6.87Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Volume(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.43 4.1a1 1 0 0 0-1 .12L6.65 8H3a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h3.65l4.73 3.78A1 1 0 0 0 12 20a.91.91 0 0 0 .43-.1A1 1 0 0 0 13 19V5a1 1 0 0 0-.57-.9M11 16.92l-3.38-2.7A1 1 0 0 0 7 14H4v-4h3a1 1 0 0 0 .62-.22L11 7.08Zm8.66-10.58a1 1 0 0 0-1.42 1.42a6 6 0 0 1-.38 8.84a1 1 0 0 0 .64 1.76a1 1 0 0 0 .64-.23a8 8 0 0 0 .52-11.79m-2.83 2.83a1 1 0 1 0-1.42 1.42A2 2 0 0 1 16 12a2 2 0 0 1-.71 1.53a1 1 0 0 0-.13 1.41a1 1 0 0 0 1.41.12A4 4 0 0 0 18 12a4.06 4.06 0 0 0-1.17-2.83"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeDown(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.83 9.17a1 1 0 1 0-1.42 1.42A2 2 0 0 1 18 12a2 2 0 0 1-.71 1.53a1 1 0 0 0-.13 1.41a1 1 0 0 0 1.41.12A4 4 0 0 0 20 12a4.06 4.06 0 0 0-1.17-2.83m-4.4-5.07a1 1 0 0 0-1 .12L8.65 8H5a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h3.65l4.73 3.78A1 1 0 0 0 14 20a.91.91 0 0 0 .43-.1A1 1 0 0 0 15 19V5a1 1 0 0 0-.57-.9M13 16.92l-3.38-2.7A1 1 0 0 0 9 14H6v-4h3a1 1 0 0 0 .62-.22L13 7.08Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeMute(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.43 4.1a1 1 0 0 0-1 .12L6.65 8H3a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h3.65l4.73 3.78A1 1 0 0 0 12 20a.91.91 0 0 0 .43-.1A1 1 0 0 0 13 19V5a1 1 0 0 0-.57-.9M11 16.92l-3.38-2.7A1 1 0 0 0 7 14H4v-4h3a1 1 0 0 0 .62-.22L11 7.08ZM19.91 12l1.8-1.79a1 1 0 0 0-1.42-1.42l-1.79 1.8l-1.79-1.8a1 1 0 0 0-1.42 1.42l1.8 1.79l-1.8 1.79a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l1.79-1.8l1.79 1.8a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeOff(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.93 4.1a1 1 0 0 0-1 .12L11.15 8H7.5a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h3.65l4.73 3.78a1 1 0 0 0 .62.22a.91.91 0 0 0 .43-.1a1 1 0 0 0 .57-.9V5a1 1 0 0 0-.57-.9M15.5 16.92l-3.38-2.7a1 1 0 0 0-.62-.22h-3v-4h3a1 1 0 0 0 .62-.22l3.38-2.7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeUp(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.43 4.1a1 1 0 0 0-1 .12L6.65 8H3a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h3.65l4.73 3.78A1 1 0 0 0 12 20a.91.91 0 0 0 .43-.1A1 1 0 0 0 13 19V5a1 1 0 0 0-.57-.9M11 16.92l-3.38-2.7A1 1 0 0 0 7 14H4v-4h3a1 1 0 0 0 .62-.22L11 7.08Zm4.14-12.83a1 1 0 1 0-.28 2a6 6 0 0 1 0 11.86a1 1 0 0 0 .14 2h.14a8 8 0 0 0 0-15.82Zm-.46 9.78a1 1 0 0 0 .32 2a1.13 1.13 0 0 0 .32-.05a4 4 0 0 0 0-7.54a1 1 0 0 0-.64 1.9a2 2 0 0 1 0 3.74Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vuejs(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.03 2.443h-3.642L12.013 6.4L9.63 2.444l-2.646-.001H.831L12.03 21.558L23.168 2.443Zm-6.005 15.15L4.322 4.443h2.824l4.885 8.406l4.847-8.407h2.81Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VuejsAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.867 3.374a1 1 0 0 0-.866-.502l-4.97-.004h-3.652a1.002 1.002 0 0 0-.817.425l-.559.796l-.564-.798a.998.998 0 0 0-.816-.423H6.968l-4.973.026A1 1 0 0 0 1.137 4.4l10.02 17.106a1 1 0 0 0 .863.494a1 1 0 0 0 .864-.496l9.98-17.128a1.002 1.002 0 0 0 .003-1.002M10.105 4.868l1.085 1.533a.999.999 0 0 0 .816.423h.001a1.003 1.003 0 0 0 .817-.425L13.9 4.87l1.363-.002l-3.244 5.454l-3.275-5.454Zm1.912 14.15L3.74 4.885l2.67-.015l4.754 7.918a1 1 0 0 0 .857.485h.006a1 1 0 0 0 .857-.489l4.708-7.916l2.67.003Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wall(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 9a1 1 0 0 0 0-2h-1V5h1a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2h3v2H3a1 1 0 0 0 0 2h1v2H3a1 1 0 0 0 0 2h3v2H3a1 1 0 0 0 0 2h1v2H3a1 1 0 0 0 0 2h18a1 1 0 0 0 0-2h-3v-2h3a1 1 0 0 0 0-2h-1v-2h1a1 1 0 0 0 0-2h-3V9ZM8 5h4v2H8Zm8 4v2h-4V9ZM6 9h4v2H6Zm6 4v2H8v-2Zm-2 6H6v-2h4Zm6 0h-4v-2h4Zm2-4h-4v-2h4Zm-4-8V5h4v2Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wallet(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 7h-1V6a3 3 0 0 0-3-3H5a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-8a3 3 0 0 0-3-3M5 5h10a1 1 0 0 1 1 1v1H5a1 1 0 0 1 0-2m15 10h-1a1 1 0 0 1 0-2h1Zm0-4h-1a3 3 0 0 0 0 6h1v1a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V8.83A3 3 0 0 0 5 9h14a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Watch(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 6.78V3a1 1 0 0 0-1-1H8a1 1 0 0 0-1 1v3.78A3 3 0 0 0 6 9v6a3 3 0 0 0 1 2.22V21a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1v-3.78A3 3 0 0 0 18 15V9a3 3 0 0 0-1-2.22M9 4h6v2H9Zm6 16H9v-2h6Zm1-5a1 1 0 0 1-1 1H9a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WatchAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m17 8.61l-1-5.77A1 1 0 0 0 15 2H9a1 1 0 0 0-1 .84L7 8.61a6 6 0 0 0 0 6.78l1 5.77A1 1 0 0 0 9 22h6a1 1 0 0 0 1-.84l1-5.77a6 6 0 0 0 0-6.78M9.85 4h4.3l.44 2.59a6 6 0 0 0-5.18 0Zm4.3 16h-4.3l-.44-2.59a6 6 0 0 0 5.18 0ZM12 16a4 4 0 1 1 4-4a4 4 0 0 1-4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Water(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.72 7.65a2.56 2.56 0 0 1 .56.24a4 4 0 0 0 4.1 0a2.6 2.6 0 0 1 2.56 0a4.15 4.15 0 0 0 4.12 0a2.6 2.6 0 0 1 2.56 0a4.25 4.25 0 0 0 2.08.56a3.88 3.88 0 0 0 2-.56a2.56 2.56 0 0 1 .56-.24a1 1 0 0 0-.56-1.92a4.45 4.45 0 0 0-1 .45a2.08 2.08 0 0 1-2.1 0a4.64 4.64 0 0 0-4.54 0a2.11 2.11 0 0 1-2.12 0a4.64 4.64 0 0 0-4.54 0a2.08 2.08 0 0 1-2.1 0a4.45 4.45 0 0 0-1-.45a1 1 0 1 0-.56 1.92Zm18 8.08a4.45 4.45 0 0 0-1 .45a2.08 2.08 0 0 1-2.1 0a4.64 4.64 0 0 0-4.54 0a2.11 2.11 0 0 1-2.12 0a4.64 4.64 0 0 0-4.54 0a2.08 2.08 0 0 1-2.1 0a4.45 4.45 0 0 0-1-.45a1 1 0 1 0-.56 1.92a2.56 2.56 0 0 1 .56.24a4 4 0 0 0 4.1 0a2.6 2.6 0 0 1 2.56 0a4.15 4.15 0 0 0 4.12 0a2.6 2.6 0 0 1 2.56 0a4.25 4.25 0 0 0 2.08.56a3.88 3.88 0 0 0 2-.56a2.56 2.56 0 0 1 .56-.24a1 1 0 0 0-.56-1.92Zm0-5a4.45 4.45 0 0 0-1 .45a2.08 2.08 0 0 1-2.1 0a4.64 4.64 0 0 0-4.54 0a2.11 2.11 0 0 1-2.12 0a4.64 4.64 0 0 0-4.54 0a2.08 2.08 0 0 1-2.1 0a4.45 4.45 0 0 0-1-.45a1 1 0 0 0-1.32.68a1 1 0 0 0 .68 1.24a2.56 2.56 0 0 1 .56.24a4 4 0 0 0 4.1 0a2.6 2.6 0 0 1 2.56 0a4.15 4.15 0 0 0 4.12 0a2.6 2.6 0 0 1 2.56 0a4.25 4.25 0 0 0 2.08.56a3.88 3.88 0 0 0 2-.56a2.56 2.56 0 0 1 .56-.24a1 1 0 0 0-.56-1.92Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WaterDropSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.71 20.29l-18-18a1 1 0 0 0-1.42 1.42l4 4a12.46 12.46 0 0 0-2 6.57A7.76 7.76 0 0 0 12 22a7.64 7.64 0 0 0 5.87-2.71l2.42 2.42a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42M12 20a5.76 5.76 0 0 1-5.75-5.75a10.3 10.3 0 0 1 1.47-5.11l8.74 8.73A5.67 5.67 0 0 1 12 20M10.85 5.24c.45-.42.85-.75 1.15-1c1.43 1.12 5.13 4.43 5.68 8.88a1 1 0 0 0 1 .88h.12a1 1 0 0 0 .87-1.11c-.78-6.43-6.85-10.55-7.1-10.72a1 1 0 0 0-1.12 0a18.73 18.73 0 0 0-1.96 1.61a1 1 0 0 0 1.36 1.46"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WaterGlass(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.08 7a1 1 0 0 0-1.08.92l-.35 4.55a2.67 2.67 0 0 1-1.2-.77a1 1 0 0 0-1.5 0a2.65 2.65 0 0 1-3.9 0a1 1 0 0 0-1.5 0a2.7 2.7 0 0 1-1.2.77L7 7.92A1 1 0 0 0 5.92 7A1 1 0 0 0 5 8.08l.86 11.15a3 3 0 0 0 3 2.77h6.3a3 3 0 0 0 3-2.77L19 8.08A1 1 0 0 0 18.08 7m-1.94 12.08a1 1 0 0 1-1 .92H8.85a1 1 0 0 1-1-.92L7.5 14.5a4.77 4.77 0 0 0 1.8-.79a4.66 4.66 0 0 0 5.4 0a4.77 4.77 0 0 0 1.8.79ZM12 10a3.26 3.26 0 0 0 3.25-3.25c0-2.75-2.58-4.51-2.69-4.58a1 1 0 0 0-1.12 0c-.11.08-2.69 1.83-2.69 4.58A3.26 3.26 0 0 0 12 10m0-5.7a3.61 3.61 0 0 1 1.25 2.45a1.25 1.25 0 0 1-2.5 0A3.66 3.66 0 0 1 12 4.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WebGrid(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1m-7 18H4v-7h10Zm0-9H4V4h10Zm6 9h-4V4h4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WebGridAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1M11 20H4V10h7Zm9 0h-7V10h7Zm0-12H4V4h16Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WebSection(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1m-7 18H4V4h10Zm6 0h-4V4h4Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WebSectionAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1M8 20H4V4h4Zm12 0H10V4h10Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Webcam(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 13a3 3 0 1 0-3-3a3 3 0 0 0 3 3m0-4a1 1 0 1 1-1 1a1 1 0 0 1 1-1m9.59 9.16A10 10 0 0 0 19 13.89a8 8 0 1 0-14 0a9.9 9.9 0 0 0-2.6 4.27a3 3 0 0 0 .47 2.63A3 3 0 0 0 5.3 22h13.4a3 3 0 0 0 2.42-1.21a3 3 0 0 0 .47-2.63M12 4a6 6 0 1 1-6 6a6 6 0 0 1 6-6m7.52 15.59a1 1 0 0 1-.82.41H5.3a1 1 0 0 1-.82-.41a1 1 0 0 1-.15-.87a7.85 7.85 0 0 1 1.88-3.22a8 8 0 0 0 11.58 0a7.85 7.85 0 0 1 1.88 3.22a1 1 0 0 1-.15.87"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Weight(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19 4h-1.45A3.08 3.08 0 0 0 17 3a3 3 0 0 0-2.25-1H9.27A3 3 0 0 0 7 3a3.08 3.08 0 0 0-.57 1H5a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3m-10.48.34A1 1 0 0 1 9.27 4h5.46a1 1 0 0 1 .75.34a1 1 0 0 1 .25.78l-.5 4a1 1 0 0 1-1 .88h-1.64l1.14-2.4a1 1 0 0 0-1.8-.86L10.37 10h-.6a1 1 0 0 1-1-.88l-.5-4a1 1 0 0 1 .25-.78M20 19a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1h1.37l.42 3.37a3 3 0 0 0 3 2.63h4.46a3 3 0 0 0 3-2.63L17.63 6H19a1 1 0 0 1 1 1Zm-6-3h-4a1 1 0 0 0 0 2h4a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Whatsapp(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.6 14c-.2-.1-1.5-.7-1.7-.8c-.2-.1-.4-.1-.6.1c-.2.2-.6.8-.8 1c-.1.2-.3.2-.5.1c-.7-.3-1.4-.7-2-1.2c-.5-.5-1-1.1-1.4-1.7c-.1-.2 0-.4.1-.5c.1-.1.2-.3.4-.4c.1-.1.2-.3.2-.4c.1-.1.1-.3 0-.4c-.1-.1-.6-1.3-.8-1.8c-.1-.7-.3-.7-.5-.7h-.5c-.2 0-.5.2-.6.3c-.6.6-.9 1.3-.9 2.1c.1.9.4 1.8 1 2.6c1.1 1.6 2.5 2.9 4.2 3.7c.5.2.9.4 1.4.5c.5.2 1 .2 1.6.1c.7-.1 1.3-.6 1.7-1.2c.2-.4.2-.8.1-1.2zm2.5-9.1C15.2 1 8.9 1 5 4.9c-3.2 3.2-3.8 8.1-1.6 12L2 22l5.3-1.4c1.5.8 3.1 1.2 4.7 1.2c5.5 0 9.9-4.4 9.9-9.9c.1-2.6-1-5.1-2.8-7m-2.7 14c-1.3.8-2.8 1.3-4.4 1.3c-1.5 0-2.9-.4-4.2-1.1l-.3-.2l-3.1.8l.8-3l-.2-.3c-2.4-4-1.2-9 2.7-11.5S16.6 3.7 19 7.5c2.4 3.9 1.3 9-2.6 11.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WhatsappAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 6.55a12.61 12.61 0 0 0-.1-1.29a4.29 4.29 0 0 0-.37-1.08a3.66 3.66 0 0 0-.71-1a3.91 3.91 0 0 0-1-.71a4.28 4.28 0 0 0-1.08-.36A10.21 10.21 0 0 0 17.46 2H6.55a12.61 12.61 0 0 0-1.29.1a4.29 4.29 0 0 0-1.08.37a3.66 3.66 0 0 0-1 .71a3.91 3.91 0 0 0-.71 1a4.28 4.28 0 0 0-.36 1.08A10.21 10.21 0 0 0 2 6.54v10.91a12.61 12.61 0 0 0 .1 1.29a4.29 4.29 0 0 0 .37 1.08a3.66 3.66 0 0 0 .71 1a3.91 3.91 0 0 0 1 .71a4.28 4.28 0 0 0 1.08.36a10.21 10.21 0 0 0 1.28.11h10.91a12.61 12.61 0 0 0 1.29-.1a4.29 4.29 0 0 0 1.08-.37a3.66 3.66 0 0 0 1-.71a3.91 3.91 0 0 0 .71-1a4.28 4.28 0 0 0 .36-1.08a10.21 10.21 0 0 0 .11-1.28V7.08zM12.23 19a7.12 7.12 0 0 1-3.43-.9l-3.8 1l1-3.72a7.11 7.11 0 0 1-1-3.58a7.18 7.18 0 1 1 7.23 7.2m0-13.13A6 6 0 0 0 7.18 15l.14.23l-.6 2.19L9 16.8l.22.13a6 6 0 0 0 3 .83a6 6 0 0 0 6-6a6 6 0 0 0-6-6Zm3.5 8.52a1.82 1.82 0 0 1-1.21.85a2.33 2.33 0 0 1-1.12-.07a8.9 8.9 0 0 1-1-.38a8 8 0 0 1-3.06-2.7a3.48 3.48 0 0 1-.73-1.85a2 2 0 0 1 .63-1.5a.65.65 0 0 1 .48-.22H10c.11 0 .26 0 .4.31s.51 1.24.56 1.33a.34.34 0 0 1 0 .31a1.14 1.14 0 0 1-.18.3c-.09.11-.19.24-.27.32s-.18.18-.08.36a5.56 5.56 0 0 0 1 1.24a5 5 0 0 0 1.44.89c.18.09.29.08.39 0s.45-.52.57-.7s.24-.15.4-.09s1.05.49 1.23.58s.29.13.34.21a1.56 1.56 0 0 1-.07.78Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WheelBarrow(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2h-2.3l-.16.07l-.17.11a.8.8 0 0 0-.13.13a.86.86 0 0 0-.1.16a.71.71 0 0 0-.08.18v.09L17.38 6h-1.14l-3.12-3.11a3.06 3.06 0 0 0-4.24 0L5.76 6H3a1 1 0 0 0-.87.5a1 1 0 0 0 0 1l4 7l-.77 1.5A2.2 2.2 0 0 0 5 16a3 3 0 1 0 3 3a3 3 0 0 0-.85-2.08l1-2l1.38-.14l3.94 5.91A2.93 2.93 0 0 0 16 22a3.18 3.18 0 0 0 1.13-.21a3 3 0 0 0 1.87-3.3L18 13l1.79-9H21a1 1 0 0 0 0-2M5 20a1 1 0 1 1 1-1a1 1 0 0 1-1 1m5.3-15.71a1 1 0 0 1 1.4 0L13.41 6H8.59Zm-2.75 8.65L4.72 8H17l-.82 4.08Zm8.81 7a1 1 0 0 1-1.2-.38l-3.34-5l4.37-.43l.81 4.69a1 1 0 0 1-.64 1.11Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wheelchair(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 6.5a2 2 0 1 0-2-2a2 2 0 0 0 2 2m7.5 14h-1v-5a1 1 0 0 0-1-1h-5v-2h5a1 1 0 0 0 0-2h-5v-2a1 1 0 0 0-2 0v7a1 1 0 0 0 1 1h5v5a1 1 0 0 0 1 1h2a1 1 0 0 0 0-2m-6.8-1.6a4 4 0 0 1-7.2-2.4a4 4 0 0 1 2.4-3.66A1 1 0 1 0 7.1 11a6 6 0 1 0 7.2 9.1a1 1 0 0 0-1.6-1.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WheelchairAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 16.18V16a3 3 0 0 0-2-2.82V7h1a1 1 0 0 0 0-2H7a3 3 0 0 0-3-3H3a1 1 0 0 0 0 2h1a1 1 0 0 1 1 1v7.42A5 5 0 1 0 12 17a4.94 4.94 0 0 0-.42-2H17a1 1 0 0 1 1 1v.18a3 3 0 1 0 2 0M7 20a3 3 0 1 1 3-3a3 3 0 0 1-3 3m9-7h-6a4.93 4.93 0 0 0-3-1v-1h9Zm0-4H7V7h9Zm3 11a1 1 0 1 1 1-1a1 1 0 0 1-1 1M7 16a1 1 0 1 0 1 1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wifi(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 15a3 3 0 1 0 3 3a3 3 0 0 0-3-3m0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1m0-8a7.06 7.06 0 0 0-4.95 2.05a1 1 0 0 0 0 1.41a1 1 0 0 0 1.41 0a5 5 0 0 1 7.08 0a1 1 0 0 0 .7.3a1 1 0 0 0 .76-1.71A7.06 7.06 0 0 0 12 11m0-4a11.08 11.08 0 0 0-7.78 3.22a1 1 0 0 0 1.42 1.42a9 9 0 0 1 12.72 0a1 1 0 0 0 .71.29a1 1 0 0 0 .71-.29a1 1 0 0 0 0-1.42A11.08 11.08 0 0 0 12 7m10.61.39a15 15 0 0 0-21.22 0a1 1 0 0 0 1.42 1.42a13 13 0 0 1 18.38 0a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiRouter(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.9 17a1 1 0 1 0 1 1a1 1 0 0 0-1-1m-3 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1m6 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1M15 8.5a1 1 0 0 1 1.73 0a1 1 0 0 0 1.36.37a1 1 0 0 0 .41-1.37a3 3 0 0 0-5.2 0a1 1 0 0 0 1.7 1m7-3a7 7 0 0 0-12.12 0a1 1 0 0 0 .37 1.37a1 1 0 0 0 .45.13a1 1 0 0 0 .87-.5a5 5 0 0 1 8.66 0a1 1 0 0 0 1.37.37A1 1 0 0 0 22 5.5M17.9 14h-1v-3a1 1 0 1 0-2 0v3h-10a3 3 0 0 0-3 3v2a3 3 0 0 0 3 3h13a3 3 0 0 0 3-3v-2a3 3 0 0 0-3-3m1 5a1 1 0 0 1-1 1h-13a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1h13a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiSlash(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.92 5.51L3.71 2.29a1 1 0 0 0-1.42 1.42L4.56 6A15.21 15.21 0 0 0 1.4 8.39a1 1 0 0 0 0 1.41a1 1 0 0 0 .71.3a1 1 0 0 0 .7-.29a13.07 13.07 0 0 1 3.24-2.35L7.54 9a10.78 10.78 0 0 0-3.32 2.27a1 1 0 1 0 1.42 1.4a8.8 8.8 0 0 1 3.45-2.12l1.62 1.61a7.07 7.07 0 0 0-3.66 1.94a1 1 0 1 0 1.42 1.4A5 5 0 0 1 12 14a4.13 4.13 0 0 1 .63.05l7.66 7.66a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM12 16a3 3 0 1 0 3 3a3 3 0 0 0-3-3m0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1M22.61 8.39A15 15 0 0 0 10.29 4.1a1 1 0 1 0 .22 2A13.07 13.07 0 0 1 21.2 9.81a1 1 0 0 0 1.41-1.42m-4.25 4.24a1 1 0 0 0 1.42-1.4a10.75 10.75 0 0 0-4.84-2.82a1 1 0 1 0-.52 1.92a8.94 8.94 0 0 1 3.94 2.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wind(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.5 9a1 1 0 1 0-1-1a1 1 0 0 0 1 1m4 0h7a3 3 0 0 0 0-6a3.06 3.06 0 0 0-1.5.4a1 1 0 0 0-.37 1.37a1 1 0 0 0 1.37.36a1.09 1.09 0 0 1 .5-.13a1 1 0 0 1 0 2h-7a1 1 0 0 0 0 2m-4 4h7a1 1 0 0 0 0-2h-7a1 1 0 0 0 0 2m17-4a1 1 0 1 0-1-1a1 1 0 0 0 1 1m-2 2h-4a1 1 0 0 0 0 2h4a1 1 0 0 1 0 2a1 1 0 0 0 0 2a3 3 0 0 0 0-6m-6 4h-4a1 1 0 0 0 0 2h4a1 1 0 0 1 0 2a1.09 1.09 0 0 1-.5-.13a1 1 0 1 0-1 1.73a3.06 3.06 0 0 0 1.5.4a3 3 0 0 0 0-6m-8 0h-1a1 1 0 0 0 0 2h1a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindMoon(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 19H7a1 1 0 0 0 0 2h5a1.013 1.013 0 0 1 1 1a1 1 0 0 0 2 0a3.003 3.003 0 0 0-3-3m2-10a1 1 0 1 0-1-1a1 1 0 0 0 1 1m4 0h2a3.003 3.003 0 0 0 3-3a1 1 0 0 0-2 0a1.013 1.013 0 0 1-1 1h-2a1 1 0 0 0 0 2m2 6h-2.161a8.043 8.043 0 0 0 1.146-2.95a1 1 0 0 0-1.305-1.117a5.97 5.97 0 0 1-1.92.317A6.062 6.062 0 0 1 9.7 5.2a7.155 7.155 0 0 1 .098-1.049A1 1 0 0 0 8.49 3.053A8.032 8.032 0 0 0 4.266 15H3a1 1 0 0 0 0 2h5a1 1 0 0 0 0-2H6.83a6.028 6.028 0 0 1 .882-9.36a8.065 8.065 0 0 0 8.048 7.61a7.878 7.878 0 0 0 .789-.04A6.027 6.027 0 0 1 15.277 15H12a1 1 0 0 0 0 2h8a1 1 0 0 1 0 2a1 1 0 0 0 0 2a3 3 0 0 0 0-6M2.62 19.08a1.147 1.147 0 0 0-.33.21A1.028 1.028 0 0 0 2 20a.99.99 0 0 0 1.38.92a1.16 1.16 0 0 0 .33-.21A.993.993 0 0 0 4 20a1.052 1.052 0 0 0-.29-.71a1.002 1.002 0 0 0-1.09-.21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindSun(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 4.5a1 1 0 0 0 1-1v-1a1 1 0 0 0-2 0v1a1 1 0 0 0 1 1m-7 7a1 1 0 0 0-1-1H2a1 1 0 0 0 0 2h1a1 1 0 0 0 1-1m.636 4.95l-.707.707a1 1 0 1 0 1.414 1.414l.707-.707a1 1 0 1 0-1.414-1.414m0-9.9A1 1 0 0 0 6.05 5.136l-.707-.707a1 1 0 0 0-1.414 1.414Zm12.021.293a.997.997 0 0 0 .707-.293l.707-.707a1 1 0 1 0-1.414-1.414l-.707.707a1 1 0 0 0 .707 1.707M13 12a1 1 0 0 0 1 1h5a3.003 3.003 0 0 0 3-3a1 1 0 0 0-2 0a1.013 1.013 0 0 1-1 1h-5a1 1 0 0 0-1 1m7 3h-5a1 1 0 0 0 0 2h5a1 1 0 0 1 0 2a1 1 0 0 0 0 2a3 3 0 0 0 0-6M9 19a1 1 0 1 0 1 1a1 1 0 0 0-1-1m6 0h-2a1 1 0 0 0 0 2h2a1.013 1.013 0 0 1 1 1a1 1 0 0 0 2 0a3.003 3.003 0 0 0-3-3m-4-4a3.5 3.5 0 0 1 0-7a3.415 3.415 0 0 1 2.188.774a1 1 0 1 0 1.265-1.548A5.393 5.393 0 0 0 11 6a5.5 5.5 0 0 0 0 11a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Window(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 5a1 1 0 1 0 1 1a1 1 0 0 0-1-1M6 5a1 1 0 1 0 1 1a1 1 0 0 0-1-1m8 0a1 1 0 1 0 1 1a1 1 0 0 0-1-1m6-4H4a3 3 0 0 0-3 3v16a3 3 0 0 0 3 3h16a3 3 0 0 0 3-3V4a3 3 0 0 0-3-3m1 19a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1v-9h18Zm0-11H3V4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowGrid(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1M8 20H4V4h4Zm12 0H10v-7h10Zm0-9H10V4h10Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowMaximize(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1m-1 18H4V10h16Zm0-12H4V4h16Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindowSection(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 2H3a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1M8 20H4V10h4Zm6 0h-4V10h4Zm6 0h-4V10h4Zm0-12H4V4h16Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Windows(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 2L11.2 3.6v8l10.8-.1zM10.2 12.5L2 12.4v6.8l8.1 1.1zM2 4.8v6.8h8.1V3.7zm9.1 7.7v7.9L22 22v-9.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Windsock(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.08 5L10 4.33l-3-.25V3a1 1 0 0 0-2 0v18a1 1 0 0 0 2 0v-7.08l3-.25l8.08-.67a1 1 0 0 0 .92-1V6a1 1 0 0 0-.92-1M9 11.75l-2 .16V6.09l2 .16Zm4-.34l-2 .17V6.42l2 .17Zm4-.33l-2 .17v-4.5l2 .17Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Windy(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 15a1.73 1.73 0 0 1-.86-.23a3.11 3.11 0 0 0-3.27 0a1.73 1.73 0 0 1-1.73 0a3.11 3.11 0 0 0-3.27 0A1.74 1.74 0 0 1 7 15a1 1 0 0 0 0 2a3.72 3.72 0 0 0 1.9-.52a1.13 1.13 0 0 1 1.2 0a3.75 3.75 0 0 0 3.8 0a1.13 1.13 0 0 1 1.2 0A3.72 3.72 0 0 0 17 17a1 1 0 0 0 0-2m0 4a1.73 1.73 0 0 1-.86-.23a3.11 3.11 0 0 0-3.27 0a1.73 1.73 0 0 1-1.73 0a3.11 3.11 0 0 0-3.27 0A1.74 1.74 0 0 1 7 19a1 1 0 0 0 0 2a3.72 3.72 0 0 0 1.9-.52a1.13 1.13 0 0 1 1.2 0a3.75 3.75 0 0 0 3.8 0a1.13 1.13 0 0 1 1.2 0A3.72 3.72 0 0 0 17 21a1 1 0 0 0 0-2m1.42-11.79a7 7 0 0 0-13.36 1.9A4 4 0 0 0 2 13a4 4 0 0 0 1.34 3a1 1 0 0 0 .66.25a1 1 0 0 0 .75-.35a1 1 0 0 0-.09-1.41A1.93 1.93 0 0 1 4 13a2 2 0 0 1 2-2a1 1 0 0 0 1-1a5 5 0 0 1 9.73-1.61a1 1 0 0 0 .78.66A3 3 0 0 1 20 12a2.93 2.93 0 0 1-.74 2a1 1 0 1 0 1.48 1.33A4.91 4.91 0 0 0 22 12a5 5 0 0 0-3.58-4.79"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wordpress(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.667 12a8.334 8.334 0 0 0 4.697 7.5L4.388 8.607A8.3 8.3 0 0 0 3.667 12m8.48.729l-2.501 7.265a8.337 8.337 0 0 0 5.121-.133a.746.746 0 0 1-.06-.115zm5.479-1.15a4.389 4.389 0 0 0-.687-2.298a3.903 3.903 0 0 1-.819-1.954a1.443 1.443 0 0 1 1.4-1.48c.037 0 .072.005.107.008a8.331 8.331 0 0 0-12.59 1.568c.196.006.38.01.537.01c.871 0 2.22-.106 2.22-.106a.345.345 0 0 1 .054.687s-.452.052-.954.079l3.035 9.026l1.824-5.47l-1.299-3.556c-.449-.027-.874-.08-.874-.08a.345.345 0 0 1 .053-.686s1.376.106 2.195.106c.871 0 2.221-.106 2.221-.106a.344.344 0 0 1 .053.687s-.452.052-.953.079l3.011 8.958l.86-2.725c.336-.88.54-1.806.606-2.746m1.743-2.72a7.866 7.866 0 0 1-.634 2.985l-2.545 7.359a8.334 8.334 0 0 0 3.123-11.2c.038.283.056.57.056.856M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2m3.659 18.662a9.388 9.388 0 0 1-8.914-.867a9.432 9.432 0 0 1-3.407-4.136a9.386 9.386 0 0 1 .867-8.914a9.427 9.427 0 0 1 4.136-3.406a9.388 9.388 0 0 1 8.914.866a9.432 9.432 0 0 1 3.407 4.136a9.386 9.386 0 0 1-.867 8.914a9.427 9.427 0 0 1-4.136 3.407"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WordpressSimple(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10.011 10.011 0 0 0 12 2M3.01 12a8.955 8.955 0 0 1 .778-3.66l4.289 11.751A8.991 8.991 0 0 1 3.009 12M12 20.99a8.987 8.987 0 0 1-2.54-.366l2.698-7.839l2.763 7.572a.844.844 0 0 0 .065.123a8.971 8.971 0 0 1-2.986.51m1.239-13.207c.541-.028 1.03-.085 1.03-.085a.372.372 0 0 0-.058-.741s-1.457.114-2.397.114c-.883 0-2.368-.114-2.368-.114a.372.372 0 0 0-.057.741s.459.057.943.085l1.401 3.838l-1.968 5.901l-3.274-9.739c.542-.028 1.03-.085 1.03-.085a.372.372 0 0 0-.058-.741s-1.456.114-2.396.114c-.169 0-.368-.004-.579-.01A8.988 8.988 0 0 1 18.071 5.37c-.039-.003-.076-.008-.116-.008a1.557 1.557 0 0 0-1.51 1.596a4.21 4.21 0 0 0 .883 2.109a4.736 4.736 0 0 1 .741 2.48a10.883 10.883 0 0 1-.684 2.906l-.897 2.996ZM16.52 19.77l2.746-7.94a8.489 8.489 0 0 0 .684-3.22a6.91 6.91 0 0 0-.06-.925a8.992 8.992 0 0 1-3.37 12.085"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WrapText(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 7h18a1 1 0 0 0 0-2H3a1 1 0 0 0 0 2m6 8H3a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2m9.5-5H3a1 1 0 0 0 0 2h15.5a1.5 1.5 0 0 1 0 3h-3.09l.3-.29a1 1 0 0 0-1.42-1.42l-2 2a1 1 0 0 0-.21.33a1 1 0 0 0 0 .76a1 1 0 0 0 .21.33l2 2a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-.3-.29h3.09a3.5 3.5 0 0 0 0-7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wrench(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.71 15.58l-4.52-4.51a6.85 6.85 0 0 0 .14-1.4A7.67 7.67 0 0 0 6.42 2.72a1 1 0 0 0-.57.74a1 1 0 0 0 .28.88l4.35 4.34l-1.8 1.8l-4.34-4.35a1 1 0 0 0-.88-.27a1 1 0 0 0-.74.56a7.67 7.67 0 0 0 7 10.91a6.85 6.85 0 0 0 1.4-.14l4.51 4.52a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42l-4.9-4.9a1 1 0 0 0-.95-.26a5.88 5.88 0 0 1-1.48.2A5.67 5.67 0 0 1 4 9.67a6 6 0 0 1 .08-1L8 12.6a1 1 0 0 0 1.42 0l3.18-3.21a1 1 0 0 0 0-1.39L8.71 4.08a6.12 6.12 0 0 1 1-.08a5.67 5.67 0 0 1 5.66 5.67a5.88 5.88 0 0 1-.2 1.48a1 1 0 0 0 .26.95l4.9 4.9a1 1 0 0 0 1.42-1.42Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func X(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8 19a3 3 0 0 1-3-3V8a3 3 0 0 1 3-3a1 1 0 0 0 0-2a5 5 0 0 0-5 5v8a5 5 0 0 0 5 5a1 1 0 0 0 0-2m7.71-3.29a1 1 0 0 0 0-1.42L13.41 12l2.3-2.29a1 1 0 0 0-1.42-1.42L12 10.59l-2.29-2.3a1 1 0 0 0-1.42 1.42l2.3 2.29l-2.3 2.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l2.29-2.3l2.29 2.3a1 1 0 0 0 1.42 0M16 3a1 1 0 0 0 0 2a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3a1 1 0 0 0 0 2a5 5 0 0 0 5-5V8a5 5 0 0 0-5-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Xadd(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.71 7.29a1 1 0 0 0-1.42 0L11 9.59l-2.29-2.3a1 1 0 1 0-1.42 1.42L9.59 11l-2.3 2.29a1 1 0 0 0 0 1.42a1 1 0 0 0 1.42 0l2.29-2.3l2.29 2.3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L12.41 11l2.3-2.29a1 1 0 0 0 0-1.42M7 18a3 3 0 0 1-3-3V7a3 3 0 0 1 3-3a1 1 0 0 0 0-2a5 5 0 0 0-5 5v8a5 5 0 0 0 5 5a1 1 0 0 0 0-2M18 7v6a1 1 0 0 0 2 0V7a5 5 0 0 0-5-5a1 1 0 0 0 0 2a3 3 0 0 1 3 3m3 11h-1v-1a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2h1v1a1 1 0 0 0 2 0v-1h1a1 1 0 0 0 0-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Yen(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.55 2.17a1 1 0 0 0-1.38.28L12 10.2L6.83 2.45a1 1 0 0 0-1.66 1.1l5 7.45H7a1 1 0 0 0 0 2h4v2H7a1 1 0 0 0 0 2h4v4a1 1 0 0 0 2 0v-4h4a1 1 0 0 0 0-2h-4v-2h4a1 1 0 0 0 0-2h-3.13l5-7.45a1 1 0 0 0-.32-1.38"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YenCircle(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1a11 11 0 1 0 11 11A11 11 0 0 0 12 1m0 20a9 9 0 1 1 9-9a9 9 0 0 1-9 9m2.83-14.55L12 10.7L9.17 6.45a1 1 0 0 0-1.39-.28a1 1 0 0 0-.28 1.38l2.11 3.17H9a1 1 0 0 0 0 2h2l.05.08v.92H9a1 1 0 0 0 0 2h2V18a1 1 0 0 0 2 0v-2.28h2a1 1 0 0 0 0-2h-2v-.92l.05-.08H15a1 1 0 0 0 0-2h-.61l2.11-3.17a1 1 0 0 0-.28-1.38a1 1 0 0 0-1.39.28"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YinYang(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 15.5a1 1 0 1 0 1 1a1 1 0 0 0-1-1M12 2a5.545 5.545 0 0 0-.562.029A9.993 9.993 0 0 0 12 22a5.545 5.545 0 0 0 .562-.028A9.993 9.993 0 0 0 12 2m0 18A7.989 7.989 0 0 1 6.71 6.015A5.484 5.484 0 0 0 12 13a3.5 3.5 0 0 1 0 7m5.29-2.015A5.484 5.484 0 0 0 12 11a3.5 3.5 0 0 1 0-7a7.989 7.989 0 0 1 5.29 13.985M12 6.5a1 1 0 1 0 1 1a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Youtube(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23 9.71a8.5 8.5 0 0 0-.91-4.13a2.92 2.92 0 0 0-1.72-1A78.36 78.36 0 0 0 12 4.27a78.45 78.45 0 0 0-8.34.3a2.87 2.87 0 0 0-1.46.74c-.9.83-1 2.25-1.1 3.45a48.29 48.29 0 0 0 0 6.48a9.55 9.55 0 0 0 .3 2a3.14 3.14 0 0 0 .71 1.36a2.86 2.86 0 0 0 1.49.78a45.18 45.18 0 0 0 6.5.33c3.5.05 6.57 0 10.2-.28a2.88 2.88 0 0 0 1.53-.78a2.49 2.49 0 0 0 .61-1a10.58 10.58 0 0 0 .52-3.4c.04-.56.04-3.94.04-4.54M9.74 14.85V8.66l5.92 3.11c-1.66.92-3.85 1.96-5.92 3.08"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YoutubeAlt(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path d="M14.89 10.48l-3-1.74a1.73 1.73 0 0 0-1.76 0a1.71 1.71 0 0 0-.87 1.52v3.48a1.71 1.71 0 0 0 .87 1.52a1.73 1.73 0 0 0 1.76 0l3-1.74a1.76 1.76 0 0 0 0-3zm-3.65 2.84v-2.64L13.52 12zM19 4H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3zm1 13a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1z" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZeroPlus(children ...ElementRenderer) *UilIcon {
	return &UilIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 5h1v1a1 1 0 0 0 2 0V5h1a1 1 0 0 0 0-2h-1V2a1 1 0 0 0-2 0v1h-1a1 1 0 0 0 0 2m-5.5 1h-1a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3h1a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3m1 9a1 1 0 0 1-1 1h-1a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1h1a1 1 0 0 1 1 1Zm8.1-6a1 1 0 0 0-.78 1.18a9 9 0 1 1-7-7a1 1 0 1 0 .4-2A10.8 10.8 0 0 0 12 1a11 11 0 1 0 11 11a10.8 10.8 0 0 0-.22-2.2A1 1 0 0 0 21.6 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
