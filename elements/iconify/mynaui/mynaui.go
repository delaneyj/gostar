package mynaui

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = ""
	hAttr          = "1em"
	viewbox        = "0 0 24 24"
	fill           = "currentColor"
)

type MynauiIcon struct {
	*SVGSVGElement
}

type MynauiIconFn func(children ...ElementRenderer) *MynauiIcon

var IconLookup = map[string]MynauiIconFn{
	"academicHat":            AcademicHat,
	"accessibility":          Accessibility,
	"activity":               Activity,
	"activitySquare":         ActivitySquare,
	"addQueue":               AddQueue,
	"aeroplane":              Aeroplane,
	"airplay":                Airplay,
	"airpods":                Airpods,
	"alarm":                  Alarm,
	"alarmCheck":             AlarmCheck,
	"alarmMinus":             AlarmMinus,
	"alarmPlus":              AlarmPlus,
	"alarmSnooze":            AlarmSnooze,
	"alarmX":                 AlarmX,
	"alignBottom":            AlignBottom,
	"alignHorizontal":        AlignHorizontal,
	"alignLeft":              AlignLeft,
	"alignRight":             AlignRight,
	"alignTop":               AlignTop,
	"alignVertical":          AlignVertical,
	"alt":                    Alt,
	"anchor":                 Anchor,
	"aperture":               Aperture,
	"api":                    Api,
	"ar":                     Ar,
	"archive":                Archive,
	"arrowDiagonalOne":       ArrowDiagonalOne,
	"arrowDiagonalTwo":       ArrowDiagonalTwo,
	"arrowDown":              ArrowDown,
	"arrowDownCircle":        ArrowDownCircle,
	"arrowDownLeft":          ArrowDownLeft,
	"arrowDownLeftCircle":    ArrowDownLeftCircle,
	"arrowDownLeftSquare":    ArrowDownLeftSquare,
	"arrowDownLeftWaves":     ArrowDownLeftWaves,
	"arrowDownRight":         ArrowDownRight,
	"arrowDownRightCircle":   ArrowDownRightCircle,
	"arrowDownRightSquare":   ArrowDownRightSquare,
	"arrowDownRightWaves":    ArrowDownRightWaves,
	"arrowDownSquare":        ArrowDownSquare,
	"arrowDownWaves":         ArrowDownWaves,
	"arrowLeft":              ArrowLeft,
	"arrowLeftCircle":        ArrowLeftCircle,
	"arrowLeftRight":         ArrowLeftRight,
	"arrowLeftSquare":        ArrowLeftSquare,
	"arrowLeftWaves":         ArrowLeftWaves,
	"arrowLongDown":          ArrowLongDown,
	"arrowLongDownLeft":      ArrowLongDownLeft,
	"arrowLongDownRight":     ArrowLongDownRight,
	"arrowLongLeft":          ArrowLongLeft,
	"arrowLongRight":         ArrowLongRight,
	"arrowLongUp":            ArrowLongUp,
	"arrowLongUpLeft":        ArrowLongUpLeft,
	"arrowLongUpRight":       ArrowLongUpRight,
	"arrowRight":             ArrowRight,
	"arrowRightCircle":       ArrowRightCircle,
	"arrowRightSquare":       ArrowRightSquare,
	"arrowRightWaves":        ArrowRightWaves,
	"arrowUp":                ArrowUp,
	"arrowUpCircle":          ArrowUpCircle,
	"arrowUpDown":            ArrowUpDown,
	"arrowUpLeft":            ArrowUpLeft,
	"arrowUpLeftCircle":      ArrowUpLeftCircle,
	"arrowUpLeftSquare":      ArrowUpLeftSquare,
	"arrowUpLeftWaves":       ArrowUpLeftWaves,
	"arrowUpRight":           ArrowUpRight,
	"arrowUpRightCircle":     ArrowUpRightCircle,
	"arrowUpRightSquare":     ArrowUpRightSquare,
	"arrowUpRightWaves":      ArrowUpRightWaves,
	"arrowUpSquare":          ArrowUpSquare,
	"arrowUpWaves":           ArrowUpWaves,
	"asteriskCircle":         AsteriskCircle,
	"asteriskHexagon":        AsteriskHexagon,
	"asteriskOctagon":        AsteriskOctagon,
	"asteriskSquare":         AsteriskSquare,
	"asteriskWaves":          AsteriskWaves,
	"at":                     At,
	"atom":                   Atom,
	"bank":                   Bank,
	"baseball":               Baseball,
	"batteryCharging":        BatteryCharging,
	"batteryChargingFour":    BatteryChargingFour,
	"batteryChargingOne":     BatteryChargingOne,
	"batteryChargingThree":   BatteryChargingThree,
	"batteryChargingTwo":     BatteryChargingTwo,
	"batteryCheck":           BatteryCheck,
	"batteryEmpty":           BatteryEmpty,
	"batteryFull":            BatteryFull,
	"batteryMinus":           BatteryMinus,
	"batteryPlus":            BatteryPlus,
	"batteryX":               BatteryX,
	"bell":                   Bell,
	"bellCheck":              BellCheck,
	"bellMinus":              BellMinus,
	"bellOn":                 BellOn,
	"bellPlus":               BellPlus,
	"bellSlash":              BellSlash,
	"bellSnooze":             BellSnooze,
	"bellX":                  BellX,
	"bitcoin":                Bitcoin,
	"bitcoinCircle":          BitcoinCircle,
	"bitcoinHexagon":         BitcoinHexagon,
	"bitcoinOctagon":         BitcoinOctagon,
	"bitcoinSquare":          BitcoinSquare,
	"bitcoinWaves":           BitcoinWaves,
	"bluetooth":              Bluetooth,
	"boat":                   Boat,
	"book":                   Book,
	"bookOpen":               BookOpen,
	"bookmark":               Bookmark,
	"bookmarkCheck":          BookmarkCheck,
	"bookmarkMinus":          BookmarkMinus,
	"bookmarkPlus":           BookmarkPlus,
	"bookmarkX":              BookmarkX,
	"boundingBox":            BoundingBox,
	"bowl":                   Bowl,
	"box":                    Box,
	"brandChrome":            BrandChrome,
	"brandDribbble":          BrandDribbble,
	"brandFacebook":          BrandFacebook,
	"brandFigma":             BrandFigma,
	"brandFramer":            BrandFramer,
	"brandGithub":            BrandGithub,
	"brandGitlab":            BrandGitlab,
	"brandGoogle":            BrandGoogle,
	"brandInstagram":         BrandInstagram,
	"brandLinkedin":          BrandLinkedin,
	"brandPinterest":         BrandPinterest,
	"brandPocket":            BrandPocket,
	"brandSlack":             BrandSlack,
	"brandSpotify":           BrandSpotify,
	"brandThreads":           BrandThreads,
	"brandTrello":            BrandTrello,
	"brandTwitch":            BrandTwitch,
	"brandTwitter":           BrandTwitter,
	"brandX":                 BrandX,
	"brandYoutube":           BrandYoutube,
	"briefcase":              Briefcase,
	"brightnessHigh":         BrightnessHigh,
	"brightnessLow":          BrightnessLow,
	"calendar":               Calendar,
	"calendarCheck":          CalendarCheck,
	"calendarDown":           CalendarDown,
	"calendarMinus":          CalendarMinus,
	"calendarPlus":           CalendarPlus,
	"calendarSlash":          CalendarSlash,
	"calendarUp":             CalendarUp,
	"calendarX":              CalendarX,
	"camera":                 Camera,
	"cameraSlash":            CameraSlash,
	"campfire":               Campfire,
	"cart":                   Cart,
	"cartCheck":              CartCheck,
	"cartMinus":              CartMinus,
	"cartPlus":               CartPlus,
	"cartX":                  CartX,
	"castScreen":             CastScreen,
	"centerFocus":            CenterFocus,
	"chartBar":               ChartBar,
	"chartBarOne":            ChartBarOne,
	"chartBarTwo":            ChartBarTwo,
	"chartBubble":            ChartBubble,
	"chartGraph":             ChartGraph,
	"chartLine":              ChartLine,
	"chartPie":               ChartPie,
	"chartPieOne":            ChartPieOne,
	"chartPieTwo":            ChartPieTwo,
	"chat":                   Chat,
	"chatCheck":              ChatCheck,
	"chatDots":               ChatDots,
	"chatMessages":           ChatMessages,
	"chatMinus":              ChatMinus,
	"chatPlus":               ChatPlus,
	"chatX":                  ChatX,
	"check":                  Check,
	"checkCircle":            CheckCircle,
	"checkCircleOne":         CheckCircleOne,
	"checkHexagon":           CheckHexagon,
	"checkOctagon":           CheckOctagon,
	"checkSquare":            CheckSquare,
	"checkSquareOne":         CheckSquareOne,
	"checkWaves":             CheckWaves,
	"chevronDoubleDown":      ChevronDoubleDown,
	"chevronDoubleDownLeft":  ChevronDoubleDownLeft,
	"chevronDoubleDownRight": ChevronDoubleDownRight,
	"chevronDoubleLeft":      ChevronDoubleLeft,
	"chevronDoubleRight":     ChevronDoubleRight,
	"chevronDoubleUp":        ChevronDoubleUp,
	"chevronDoubleUpLeft":    ChevronDoubleUpLeft,
	"chevronDoubleUpRight":   ChevronDoubleUpRight,
	"chevronDown":            ChevronDown,
	"chevronDownCircle":      ChevronDownCircle,
	"chevronDownLeft":        ChevronDownLeft,
	"chevronDownLeftCircle":  ChevronDownLeftCircle,
	"chevronDownLeftSquare":  ChevronDownLeftSquare,
	"chevronDownLeftWaves":   ChevronDownLeftWaves,
	"chevronDownRight":       ChevronDownRight,
	"chevronDownRightCircle": ChevronDownRightCircle,
	"chevronDownRightSquare": ChevronDownRightSquare,
	"chevronDownRightWaves":  ChevronDownRightWaves,
	"chevronDownSquare":      ChevronDownSquare,
	"chevronDownWaves":       ChevronDownWaves,
	"chevronLeft":            ChevronLeft,
	"chevronLeftCircle":      ChevronLeftCircle,
	"chevronLeftSquare":      ChevronLeftSquare,
	"chevronLeftWaves":       ChevronLeftWaves,
	"chevronRight":           ChevronRight,
	"chevronRightCircle":     ChevronRightCircle,
	"chevronRightSquare":     ChevronRightSquare,
	"chevronRightWaves":      ChevronRightWaves,
	"chevronUp":              ChevronUp,
	"chevronUpCircle":        ChevronUpCircle,
	"chevronUpDown":          ChevronUpDown,
	"chevronUpLeft":          ChevronUpLeft,
	"chevronUpLeftCircle":    ChevronUpLeftCircle,
	"chevronUpLeftSquare":    ChevronUpLeftSquare,
	"chevronUpLeftWaves":     ChevronUpLeftWaves,
	"chevronUpRight":         ChevronUpRight,
	"chevronUpRightCircle":   ChevronUpRightCircle,
	"chevronUpRightSquare":   ChevronUpRightSquare,
	"chevronUpRightWaves":    ChevronUpRightWaves,
	"chevronUpSquare":        ChevronUpSquare,
	"chevronUpWaves":         ChevronUpWaves,
	"chip":                   Chip,
	"circle":                 Circle,
	"circleDashed":           CircleDashed,
	"circleHalf":             CircleHalf,
	"circleHalfCircle":       CircleHalfCircle,
	"circleNotch":            CircleNotch,
	"click":                  Click,
	"clipboard":              Clipboard,
	"clockCircle":            ClockCircle,
	"clockEight":             ClockEight,
	"clockEleven":            ClockEleven,
	"clockFive":              ClockFive,
	"clockFour":              ClockFour,
	"clockHand":              ClockHand,
	"clockHexagon":           ClockHexagon,
	"clockNine":              ClockNine,
	"clockOctagon":           ClockOctagon,
	"clockOne":               ClockOne,
	"clockSeven":             ClockSeven,
	"clockSix":               ClockSix,
	"clockSquare":            ClockSquare,
	"clockTen":               ClockTen,
	"clockThree":             ClockThree,
	"clockTwelve":            ClockTwelve,
	"clockTwo":               ClockTwo,
	"clockWaves":             ClockWaves,
	"cloud":                  Cloud,
	"cloudDownload":          CloudDownload,
	"cloudLightning":         CloudLightning,
	"cloudRain":              CloudRain,
	"cloudSlash":             CloudSlash,
	"cloudSnow":              CloudSnow,
	"cloudUpload":            CloudUpload,
	"cocktail":               Cocktail,
	"code":                   Code,
	"codeCircle":             CodeCircle,
	"codeHexagon":            CodeHexagon,
	"codeOctagon":            CodeOctagon,
	"codeSquare":             CodeSquare,
	"codeWaves":              CodeWaves,
	"coffee":                 Coffee,
	"cog":                    Cog,
	"cogOne":                 CogOne,
	"cogThree":               CogThree,
	"cogTwo":                 CogTwo,
	"columns":                Columns,
	"command":                Command,
	"compass":                Compass,
	"components":             Components,
	"confetti":               Confetti,
	"config":                 Config,
	"configVertical":         ConfigVertical,
	"contactless":            Contactless,
	"contactlessCircle":      ContactlessCircle,
	"controller":             Controller,
	"cookie":                 Cookie,
	"copy":                   Copy,
	"copyleft":               Copyleft,
	"copyright":              Copyright,
	"copyrightSlash":         CopyrightSlash,
	"cornerDownLeft":         CornerDownLeft,
	"cornerDownRight":        CornerDownRight,
	"cornerLeftDown":         CornerLeftDown,
	"cornerLeftUp":           CornerLeftUp,
	"cornerRightDown":        CornerRightDown,
	"cornerRightUp":          CornerRightUp,
	"cornerUpLeft":           CornerUpLeft,
	"cornerUpRight":          CornerUpRight,
	"creditCard":             CreditCard,
	"creditCardCheck":        CreditCardCheck,
	"creditCardMinus":        CreditCardMinus,
	"creditCardPlus":         CreditCardPlus,
	"creditCardX":            CreditCardX,
	"croissant":              Croissant,
	"crop":                   Crop,
	"crosshair":              Crosshair,
	"cupcake":                Cupcake,
	"danger":                 Danger,
	"dangerCircle":           DangerCircle,
	"dangerHexagon":          DangerHexagon,
	"dangerOctagon":          DangerOctagon,
	"dangerSquare":           DangerSquare,
	"dangerTriangle":         DangerTriangle,
	"dangerWaves":            DangerWaves,
	"database":               Database,
	"dazeCircle":             DazeCircle,
	"dazeGhost":              DazeGhost,
	"dazeSquare":             DazeSquare,
	"delete":                 Delete,
	"desktop":                Desktop,
	"diamond":                Diamond,
	"diceFive":               DiceFive,
	"diceFour":               DiceFour,
	"diceOne":                DiceOne,
	"diceSix":                DiceSix,
	"diceThree":              DiceThree,
	"diceTwo":                DiceTwo,
	"dislike":                Dislike,
	"divide":                 Divide,
	"dollar":                 Dollar,
	"dollarCircle":           DollarCircle,
	"dollarHexagon":          DollarHexagon,
	"dollarOctagon":          DollarOctagon,
	"dollarSquare":           DollarSquare,
	"dollarWaves":            DollarWaves,
	"dots":                   Dots,
	"dotsCircle":             DotsCircle,
	"dotsHexagon":            DotsHexagon,
	"dotsOctagon":            DotsOctagon,
	"dotsSquare":             DotsSquare,
	"dotsVertical":           DotsVertical,
	"dotsVerticalCircle":     DotsVerticalCircle,
	"dotsVerticalHexagon":    DotsVerticalHexagon,
	"dotsVerticalOctagon":    DotsVerticalOctagon,
	"dotsVerticalSquare":     DotsVerticalSquare,
	"dotsVerticalWaves":      DotsVerticalWaves,
	"dotsWaves":              DotsWaves,
	"download":               Download,
	"drop":                   Drop,
	"earth":                  Earth,
	"edit":                   Edit,
	"editOne":                EditOne,
	"egg":                    Egg,
	"eight":                  Eight,
	"eightCircle":            EightCircle,
	"eightDiamond":           EightDiamond,
	"eightHexagon":           EightHexagon,
	"eightOctagon":           EightOctagon,
	"eightSquare":            EightSquare,
	"eightWaves":             EightWaves,
	"elevator":               Elevator,
	"envelope":               Envelope,
	"euro":                   Euro,
	"euroCircle":             EuroCircle,
	"euroHexagon":            EuroHexagon,
	"euroOctagon":            EuroOctagon,
	"euroSquare":             EuroSquare,
	"euroWaves":              EuroWaves,
	"exclude":                Exclude,
	"externalLink":           ExternalLink,
	"eye":                    Eye,
	"eyeSlash":               EyeSlash,
	"faceId":                 FaceId,
	"fatArrowDown":           FatArrowDown,
	"fatArrowDownLeft":       FatArrowDownLeft,
	"fatArrowDownRight":      FatArrowDownRight,
	"fatArrowLeft":           FatArrowLeft,
	"fatArrowRight":          FatArrowRight,
	"fatArrowUp":             FatArrowUp,
	"fatArrowUpLeft":         FatArrowUpLeft,
	"fatArrowUpRight":        FatArrowUpRight,
	"fatCornerDownLeft":      FatCornerDownLeft,
	"fatCornerDownRight":     FatCornerDownRight,
	"fatCornerLeftDown":      FatCornerLeftDown,
	"fatCornerLeftUp":        FatCornerLeftUp,
	"fatCornerRightDown":     FatCornerRightDown,
	"fatCornerRightUp":       FatCornerRightUp,
	"fatCornerUpLeft":        FatCornerUpLeft,
	"fatCornerUpRight":       FatCornerUpRight,
	"female":                 Female,
	"file":                   File,
	"fileCheck":              FileCheck,
	"fileMinus":              FileMinus,
	"filePlus":               FilePlus,
	"fileText":               FileText,
	"fileX":                  FileX,
	"film":                   Film,
	"filter":                 Filter,
	"filterOne":              FilterOne,
	"fineTune":               FineTune,
	"fire":                   Fire,
	"five":                   Five,
	"fiveCircle":             FiveCircle,
	"fiveDiamond":            FiveDiamond,
	"fiveHexagon":            FiveHexagon,
	"fiveOctagon":            FiveOctagon,
	"fiveSquare":             FiveSquare,
	"fiveWaves":              FiveWaves,
	"flag":                   Flag,
	"flagOne":                FlagOne,
	"flask":                  Flask,
	"folder":                 Folder,
	"folderCheck":            FolderCheck,
	"folderHeart":            FolderHeart,
	"folderMinus":            FolderMinus,
	"folderOne":              FolderOne,
	"folderPlus":             FolderPlus,
	"folderSlash":            FolderSlash,
	"folderTwo":              FolderTwo,
	"folderX":                FolderX,
	"forward":                Forward,
	"forwardCircle":          ForwardCircle,
	"forwardHexagon":         ForwardHexagon,
	"forwardOctagon":         ForwardOctagon,
	"forwardSquare":          ForwardSquare,
	"forwardWaves":           ForwardWaves,
	"four":                   Four,
	"fourCircle":             FourCircle,
	"fourDiamond":            FourDiamond,
	"fourHexagon":            FourHexagon,
	"fourOctagon":            FourOctagon,
	"fourSquare":             FourSquare,
	"fourWaves":              FourWaves,
	"frame":                  Frame,
	"funnyCircle":            FunnyCircle,
	"funnyGhost":             FunnyGhost,
	"funnySquare":            FunnySquare,
	"ghostDaze":              GhostDaze,
	"ghostFunny":             GhostFunny,
	"ghostIndifferent":       GhostIndifferent,
	"ghostSad":               GhostSad,
	"ghostSmile":             GhostSmile,
	"gift":                   Gift,
	"gitBranch":              GitBranch,
	"gitCircle":              GitCircle,
	"gitCommit":              GitCommit,
	"gitDiff":                GitDiff,
	"gitHexagon":             GitHexagon,
	"gitMerge":               GitMerge,
	"gitOctagon":             GitOctagon,
	"gitPullRequest":         GitPullRequest,
	"gitSquare":              GitSquare,
	"gitWaves":               GitWaves,
	"globe":                  Globe,
	"grid":                   Grid,
	"gridOne":                GridOne,
	"hardDrive":              HardDrive,
	"hash":                   Hash,
	"hashCircle":             HashCircle,
	"hashHexagon":            HashHexagon,
	"hashOctagon":            HashOctagon,
	"hashSquare":             HashSquare,
	"hashWaves":              HashWaves,
	"heading":                Heading,
	"headingOne":             HeadingOne,
	"headingThree":           HeadingThree,
	"headingTwo":             HeadingTwo,
	"headphones":             Headphones,
	"heart":                  Heart,
	"heartBroken":            HeartBroken,
	"heartCheck":             HeartCheck,
	"heartCircle":            HeartCircle,
	"heartHexagon":           HeartHexagon,
	"heartMinus":             HeartMinus,
	"heartOctagon":           HeartOctagon,
	"heartPlus":              HeartPlus,
	"heartSquare":            HeartSquare,
	"heartWaves":             HeartWaves,
	"heartX":                 HeartX,
	"hexagon":                Hexagon,
	"home":                   Home,
	"homeCheck":              HomeCheck,
	"homeMinus":              HomeMinus,
	"homePlus":               HomePlus,
	"homeSmile":              HomeSmile,
	"homeX":                  HomeX,
	"image":                  Image,
	"imageCircle":            ImageCircle,
	"imageRectangle":         ImageRectangle,
	"inbox":                  Inbox,
	"inboxArchive":           InboxArchive,
	"inboxCheck":             InboxCheck,
	"inboxDown":              InboxDown,
	"inboxMinus":             InboxMinus,
	"inboxPlus":              InboxPlus,
	"inboxUp":                InboxUp,
	"inboxX":                 InboxX,
	"incognito":              Incognito,
	"indifferentCircle":      IndifferentCircle,
	"indifferentGhost":       IndifferentGhost,
	"indifferentSquare":      IndifferentSquare,
	"infinity":               Infinity,
	"info":                   Info,
	"infoCircle":             InfoCircle,
	"infoHexagon":            InfoHexagon,
	"infoOctagon":            InfoOctagon,
	"infoSquare":             InfoSquare,
	"infoTriangle":           InfoTriangle,
	"infoWaves":              InfoWaves,
	"intersect":              Intersect,
	"key":                    Key,
	"keyboard":               Keyboard,
	"keyboardBrightnessHigh": KeyboardBrightnessHigh,
	"keyboardBrightnessLow":  KeyboardBrightnessLow,
	"label":                  Label,
	"lamp":                   Lamp,
	"layersOne":              LayersOne,
	"layersThree":            LayersThree,
	"layersTwo":              LayersTwo,
	"layout":                 Layout,
	"leaf":                   Leaf,
	"leaves":                 Leaves,
	"letterA":                LetterA,
	"letterAcircle":          LetterAcircle,
	"letterAdiamond":         LetterAdiamond,
	"letterAhexagon":         LetterAhexagon,
	"letterAoctagon":         LetterAoctagon,
	"letterAsquare":          LetterAsquare,
	"letterAwaves":           LetterAwaves,
	"letterB":                LetterB,
	"letterBcircle":          LetterBcircle,
	"letterBdiamond":         LetterBdiamond,
	"letterBhexagon":         LetterBhexagon,
	"letterBoctagon":         LetterBoctagon,
	"letterBsquare":          LetterBsquare,
	"letterBwaves":           LetterBwaves,
	"letterC":                LetterC,
	"letterCcircle":          LetterCcircle,
	"letterCdiamond":         LetterCdiamond,
	"letterChexagon":         LetterChexagon,
	"letterCoctagon":         LetterCoctagon,
	"letterCsquare":          LetterCsquare,
	"letterCwaves":           LetterCwaves,
	"letterD":                LetterD,
	"letterDcircle":          LetterDcircle,
	"letterDdiamond":         LetterDdiamond,
	"letterDhexagon":         LetterDhexagon,
	"letterDoctagon":         LetterDoctagon,
	"letterDsquare":          LetterDsquare,
	"letterDwaves":           LetterDwaves,
	"letterE":                LetterE,
	"letterEcircle":          LetterEcircle,
	"letterEdiamond":         LetterEdiamond,
	"letterEhexagon":         LetterEhexagon,
	"letterEoctagon":         LetterEoctagon,
	"letterEsquare":          LetterEsquare,
	"letterEwaves":           LetterEwaves,
	"letterF":                LetterF,
	"letterFcircle":          LetterFcircle,
	"letterFdiamond":         LetterFdiamond,
	"letterFhexagon":         LetterFhexagon,
	"letterFoctagon":         LetterFoctagon,
	"letterFsquare":          LetterFsquare,
	"letterFwaves":           LetterFwaves,
	"letterG":                LetterG,
	"letterGcircle":          LetterGcircle,
	"letterGdiamond":         LetterGdiamond,
	"letterGhexagon":         LetterGhexagon,
	"letterGoctagon":         LetterGoctagon,
	"letterGsquare":          LetterGsquare,
	"letterGwaves":           LetterGwaves,
	"letterH":                LetterH,
	"letterHcircle":          LetterHcircle,
	"letterHdiamond":         LetterHdiamond,
	"letterHhexagon":         LetterHhexagon,
	"letterHoctagon":         LetterHoctagon,
	"letterHsquare":          LetterHsquare,
	"letterHwaves":           LetterHwaves,
	"letterI":                LetterI,
	"letterIcircle":          LetterIcircle,
	"letterIdiamond":         LetterIdiamond,
	"letterIhexagon":         LetterIhexagon,
	"letterIoctagon":         LetterIoctagon,
	"letterIsquare":          LetterIsquare,
	"letterIwaves":           LetterIwaves,
	"letterJ":                LetterJ,
	"letterJcircle":          LetterJcircle,
	"letterJdiamond":         LetterJdiamond,
	"letterJhexagon":         LetterJhexagon,
	"letterJoctagon":         LetterJoctagon,
	"letterJsquare":          LetterJsquare,
	"letterJwaves":           LetterJwaves,
	"letterK":                LetterK,
	"letterKcircle":          LetterKcircle,
	"letterKdiamond":         LetterKdiamond,
	"letterKhexagon":         LetterKhexagon,
	"letterKoctagon":         LetterKoctagon,
	"letterKsquare":          LetterKsquare,
	"letterKwaves":           LetterKwaves,
	"letterL":                LetterL,
	"letterLcircle":          LetterLcircle,
	"letterLdiamond":         LetterLdiamond,
	"letterLhexagon":         LetterLhexagon,
	"letterLoctagon":         LetterLoctagon,
	"letterLsquare":          LetterLsquare,
	"letterLwaves":           LetterLwaves,
	"letterM":                LetterM,
	"letterMcircle":          LetterMcircle,
	"letterMdiamond":         LetterMdiamond,
	"letterMhexagon":         LetterMhexagon,
	"letterMoctagon":         LetterMoctagon,
	"letterMsquare":          LetterMsquare,
	"letterMwaves":           LetterMwaves,
	"letterN":                LetterN,
	"letterNcircle":          LetterNcircle,
	"letterNdiamond":         LetterNdiamond,
	"letterNhexagon":         LetterNhexagon,
	"letterNoctagon":         LetterNoctagon,
	"letterNsquare":          LetterNsquare,
	"letterNwaves":           LetterNwaves,
	"letterO":                LetterO,
	"letterOcircle":          LetterOcircle,
	"letterOdiamond":         LetterOdiamond,
	"letterOhexagon":         LetterOhexagon,
	"letterOoctagon":         LetterOoctagon,
	"letterOsquare":          LetterOsquare,
	"letterOwaves":           LetterOwaves,
	"letterP":                LetterP,
	"letterPcircle":          LetterPcircle,
	"letterPdiamond":         LetterPdiamond,
	"letterPhexagon":         LetterPhexagon,
	"letterPoctagon":         LetterPoctagon,
	"letterPsquare":          LetterPsquare,
	"letterPwaves":           LetterPwaves,
	"letterQ":                LetterQ,
	"letterQcircle":          LetterQcircle,
	"letterQdiamond":         LetterQdiamond,
	"letterQhexagon":         LetterQhexagon,
	"letterQoctagon":         LetterQoctagon,
	"letterQsquare":          LetterQsquare,
	"letterQwaves":           LetterQwaves,
	"letterR":                LetterR,
	"letterRcircle":          LetterRcircle,
	"letterRdiamond":         LetterRdiamond,
	"letterRhexagon":         LetterRhexagon,
	"letterRoctagon":         LetterRoctagon,
	"letterRsquare":          LetterRsquare,
	"letterRwaves":           LetterRwaves,
	"letterS":                LetterS,
	"letterScircle":          LetterScircle,
	"letterSdiamond":         LetterSdiamond,
	"letterShexagon":         LetterShexagon,
	"letterSoctagon":         LetterSoctagon,
	"letterSsquare":          LetterSsquare,
	"letterSwaves":           LetterSwaves,
	"letterT":                LetterT,
	"letterTcircle":          LetterTcircle,
	"letterTdiamond":         LetterTdiamond,
	"letterThexagon":         LetterThexagon,
	"letterToctagon":         LetterToctagon,
	"letterTsquare":          LetterTsquare,
	"letterTwaves":           LetterTwaves,
	"letterU":                LetterU,
	"letterUcircle":          LetterUcircle,
	"letterUdiamond":         LetterUdiamond,
	"letterUhexagon":         LetterUhexagon,
	"letterUoctagon":         LetterUoctagon,
	"letterUsquare":          LetterUsquare,
	"letterUwaves":           LetterUwaves,
	"letterV":                LetterV,
	"letterVcircle":          LetterVcircle,
	"letterVdiamond":         LetterVdiamond,
	"letterVhexagon":         LetterVhexagon,
	"letterVoctagon":         LetterVoctagon,
	"letterVsquare":          LetterVsquare,
	"letterVwaves":           LetterVwaves,
	"letterW":                LetterW,
	"letterWcircle":          LetterWcircle,
	"letterWdiamond":         LetterWdiamond,
	"letterWhexagon":         LetterWhexagon,
	"letterWoctagon":         LetterWoctagon,
	"letterWsquare":          LetterWsquare,
	"letterWwaves":           LetterWwaves,
	"letterX":                LetterX,
	"letterXcircle":          LetterXcircle,
	"letterXdiamond":         LetterXdiamond,
	"letterXhexagon":         LetterXhexagon,
	"letterXoctagon":         LetterXoctagon,
	"letterXsquare":          LetterXsquare,
	"letterXwaves":           LetterXwaves,
	"letterY":                LetterY,
	"letterYcircle":          LetterYcircle,
	"letterYdiamond":         LetterYdiamond,
	"letterYhexagon":         LetterYhexagon,
	"letterYoctagon":         LetterYoctagon,
	"letterYsquare":          LetterYsquare,
	"letterYwaves":           LetterYwaves,
	"letterZ":                LetterZ,
	"letterZcircle":          LetterZcircle,
	"letterZdiamond":         LetterZdiamond,
	"letterZhexagon":         LetterZhexagon,
	"letterZoctagon":         LetterZoctagon,
	"letterZsquare":          LetterZsquare,
	"letterZwaves":           LetterZwaves,
	"lightning":              Lightning,
	"lightningSlash":         LightningSlash,
	"like":                   Like,
	"lineChartCircle":        LineChartCircle,
	"lineChartHexagon":       LineChartHexagon,
	"lineChartOctagon":       LineChartOctagon,
	"lineChartSquare":        LineChartSquare,
	"lineChartWaves":         LineChartWaves,
	"link":                   Link,
	"linkOne":                LinkOne,
	"linkTwo":                LinkTwo,
	"list":                   List,
	"listCheck":              ListCheck,
	"listNumber":             ListNumber,
	"location":               Location,
	"locationCheck":          LocationCheck,
	"locationMinus":          LocationMinus,
	"locationPlus":           LocationPlus,
	"locationX":              LocationX,
	"lock":                   Lock,
	"lockCircle":             LockCircle,
	"lockHexagon":            LockHexagon,
	"lockKeyhole":            LockKeyhole,
	"lockOctagon":            LockOctagon,
	"lockOpen":               LockOpen,
	"lockOpenKeyhole":        LockOpenKeyhole,
	"lockOpenPassword":       LockOpenPassword,
	"lockPassword":           LockPassword,
	"lockSquare":             LockSquare,
	"lockWaves":              LockWaves,
	"login":                  Login,
	"logout":                 Logout,
	"magnet":                 Magnet,
	"male":                   Male,
	"map":                    Map,
	"mask":                   Mask,
	"math":                   Math,
	"mathSquare":             MathSquare,
	"maximize":               Maximize,
	"maximizeOne":            MaximizeOne,
	"menu":                   Menu,
	"message":                Message,
	"messageCheck":           MessageCheck,
	"messageDots":            MessageDots,
	"messageMinus":           MessageMinus,
	"messagePlus":            MessagePlus,
	"messageX":               MessageX,
	"microphone":             Microphone,
	"microphoneSlash":        MicrophoneSlash,
	"minimize":               Minimize,
	"minimizeOne":            MinimizeOne,
	"minus":                  Minus,
	"minusCircle":            MinusCircle,
	"minusHexagon":           MinusHexagon,
	"minusOctagon":           MinusOctagon,
	"minusSquare":            MinusSquare,
	"minusWaves":             MinusWaves,
	"mobile":                 Mobile,
	"mobileSignalFive":       MobileSignalFive,
	"mobileSignalFour":       MobileSignalFour,
	"mobileSignalOne":        MobileSignalOne,
	"mobileSignalThree":      MobileSignalThree,
	"mobileSignalTwo":        MobileSignalTwo,
	"moon":                   Moon,
	"mountain":               Mountain,
	"mousePointer":           MousePointer,
	"move":                   Move,
	"moveDiagonal":           MoveDiagonal,
	"moveDiagonalOne":        MoveDiagonalOne,
	"moveHorizontal":         MoveHorizontal,
	"moveVertical":           MoveVertical,
	"music":                  Music,
	"musicCircle":            MusicCircle,
	"musicHexagon":           MusicHexagon,
	"musicOctagon":           MusicOctagon,
	"musicSquare":            MusicSquare,
	"musicWaves":             MusicWaves,
	"myna":                   Myna,
	"navigation":             Navigation,
	"navigationOne":          NavigationOne,
	"nine":                   Nine,
	"nineCircle":             NineCircle,
	"nineDiamond":            NineDiamond,
	"nineHexagon":            NineHexagon,
	"nineOctagon":            NineOctagon,
	"nineSquare":             NineSquare,
	"nineWaves":              NineWaves,
	"notification":           Notification,
	"octagon":                Octagon,
	"octagonCheck":           OctagonCheck,
	"octagonDanger":          OctagonDanger,
	"octagonInfo":            OctagonInfo,
	"octagonMinus":           OctagonMinus,
	"octagonPlus":            OctagonPlus,
	"octagonSlash":           OctagonSlash,
	"octagonX":               OctagonX,
	"one":                    One,
	"oneCircle":              OneCircle,
	"oneDiamond":             OneDiamond,
	"oneHexagon":             OneHexagon,
	"oneOctagon":             OneOctagon,
	"oneSquare":              OneSquare,
	"oneWaves":               OneWaves,
	"option":                 Option,
	"package":                Package,
	"paint":                  Paint,
	"panelBottom":            PanelBottom,
	"panelBottomClose":       PanelBottomClose,
	"panelBottomInactive":    PanelBottomInactive,
	"panelBottomOpen":        PanelBottomOpen,
	"panelLeft":              PanelLeft,
	"panelLeftClose":         PanelLeftClose,
	"panelLeftInactive":      PanelLeftInactive,
	"panelLeftOpen":          PanelLeftOpen,
	"panelRight":             PanelRight,
	"panelRightClose":        PanelRightClose,
	"panelRightInactive":     PanelRightInactive,
	"panelRightOpen":         PanelRightOpen,
	"panelTop":               PanelTop,
	"panelTopClose":          PanelTopClose,
	"panelTopInactive":       PanelTopInactive,
	"panelTopOpen":           PanelTopOpen,
	"paperclip":              Paperclip,
	"parking":                Parking,
	"password":               Password,
	"path":                   Path,
	"pause":                  Pause,
	"pauseCircle":            PauseCircle,
	"pauseHexagon":           PauseHexagon,
	"pauseOctagon":           PauseOctagon,
	"pauseSquare":            PauseSquare,
	"pauseWaves":             PauseWaves,
	"pen":                    Pen,
	"pencil":                 Pencil,
	"percentage":             Percentage,
	"percentageCircle":       PercentageCircle,
	"percentageHexagon":      PercentageHexagon,
	"percentageOctagon":      PercentageOctagon,
	"percentageSquare":       PercentageSquare,
	"percentageWaves":        PercentageWaves,
	"pin":                    Pin,
	"pizza":                  Pizza,
	"planet":                 Planet,
	"play":                   Play,
	"playCircle":             PlayCircle,
	"playHexagon":            PlayHexagon,
	"playOctagon":            PlayOctagon,
	"playSquare":             PlaySquare,
	"playWaves":              PlayWaves,
	"plus":                   Plus,
	"plusCircle":             PlusCircle,
	"plusHexagon":            PlusHexagon,
	"plusOctagon":            PlusOctagon,
	"plusSquare":             PlusSquare,
	"plusWaves":              PlusWaves,
	"pokeball":               Pokeball,
	"power":                  Power,
	"presentation":           Presentation,
	"printer":                Printer,
	"puzzle":                 Puzzle,
	"question":               Question,
	"questionCircle":         QuestionCircle,
	"questionHexagon":        QuestionHexagon,
	"questionOctagon":        QuestionOctagon,
	"questionSquare":         QuestionSquare,
	"questionWaves":          QuestionWaves,
	"radio":                  Radio,
	"rainbow":                Rainbow,
	"receptionBell":          ReceptionBell,
	"record":                 Record,
	"rectangle":              Rectangle,
	"rectangleVertical":      RectangleVertical,
	"redo":                   Redo,
	"refresh":                Refresh,
	"refreshAlt":             RefreshAlt,
	"repeat":                 Repeat,
	"rewind":                 Rewind,
	"rewindCircle":           RewindCircle,
	"rewindHexagon":          RewindHexagon,
	"rewindOctagon":          RewindOctagon,
	"rewindSquare":           RewindSquare,
	"rewindWaves":            RewindWaves,
	"rhombus":                Rhombus,
	"ribbon":                 Ribbon,
	"roomService":            RoomService,
	"rows":                   Rows,
	"rss":                    Rss,
	"ruler":                  Ruler,
	"rupee":                  Rupee,
	"rupeeCircle":            RupeeCircle,
	"rupeeHexagon":           RupeeHexagon,
	"rupeeOctagon":           RupeeOctagon,
	"rupeeSquare":            RupeeSquare,
	"rupeeWaves":             RupeeWaves,
	"sadCircle":              SadCircle,
	"sadGhost":               SadGhost,
	"sadSquare":              SadSquare,
	"save":                   Save,
	"scan":                   Scan,
	"scissors":               Scissors,
	"seaWaves":               SeaWaves,
	"search":                 Search,
	"searchCheck":            SearchCheck,
	"searchCircle":           SearchCircle,
	"searchHexagon":          SearchHexagon,
	"searchMinus":            SearchMinus,
	"searchOctagon":          SearchOctagon,
	"searchPlus":             SearchPlus,
	"searchSlash":            SearchSlash,
	"searchSquare":           SearchSquare,
	"searchWaves":            SearchWaves,
	"searchX":                SearchX,
	"selectMultiple":         SelectMultiple,
	"send":                   Send,
	"servers":                Servers,
	"seven":                  Seven,
	"sevenCircle":            SevenCircle,
	"sevenDiamond":           SevenDiamond,
	"sevenHexagon":           SevenHexagon,
	"sevenOctagon":           SevenOctagon,
	"sevenSquare":            SevenSquare,
	"sevenWaves":             SevenWaves,
	"share":                  Share,
	"shield":                 Shield,
	"shieldCheck":            ShieldCheck,
	"shieldCrossed":          ShieldCrossed,
	"shieldMinus":            ShieldMinus,
	"shieldOne":              ShieldOne,
	"shieldPlus":             ShieldPlus,
	"shieldSlash":            ShieldSlash,
	"shieldTwo":              ShieldTwo,
	"shieldX":                ShieldX,
	"shootingStar":           ShootingStar,
	"shoppingBag":            ShoppingBag,
	"shuffle":                Shuffle,
	"shuffleAlt":             ShuffleAlt,
	"sidebar":                Sidebar,
	"sidebarAlt":             SidebarAlt,
	"signal":                 Signal,
	"signalCircle":           SignalCircle,
	"signalHexagon":          SignalHexagon,
	"signalOctagon":          SignalOctagon,
	"signalSquare":           SignalSquare,
	"signalWaves":            SignalWaves,
	"six":                    Six,
	"sixCircle":              SixCircle,
	"sixDiamond":             SixDiamond,
	"sixHexagon":             SixHexagon,
	"sixOctagon":             SixOctagon,
	"sixSquare":              SixSquare,
	"sixWaves":               SixWaves,
	"skipBack":               SkipBack,
	"skipForward":            SkipForward,
	"slash":                  Slash,
	"slashCircle":            SlashCircle,
	"slashHexagon":           SlashHexagon,
	"slashOctagon":           SlashOctagon,
	"slashSquare":            SlashSquare,
	"slashWaves":             SlashWaves,
	"smileCircle":            SmileCircle,
	"smileGhost":             SmileGhost,
	"smileSquare":            SmileSquare,
	"snow":                   Snow,
	"sofa":                   Sofa,
	"sort":                   Sort,
	"sparkles":               Sparkles,
	"speaker":                Speaker,
	"spinner":                Spinner,
	"spinnerOne":             SpinnerOne,
	"square":                 Square,
	"squareDashed":           SquareDashed,
	"squareHalf":             SquareHalf,
	"star":                   Star,
	"stop":                   Stop,
	"stopCircle":             StopCircle,
	"stopHexagon":            StopHexagon,
	"stopOctagon":            StopOctagon,
	"stopSquare":             StopSquare,
	"stopWaves":              StopWaves,
	"store":                  Store,
	"subtract":               Subtract,
	"sun":                    Sun,
	"sunrise":                Sunrise,
	"sunset":                 Sunset,
	"support":                Support,
	"swatches":               Swatches,
	"table":                  Table,
	"tablet":                 Tablet,
	"tag":                    Tag,
	"tagPlus":                TagPlus,
	"tallyFive":              TallyFive,
	"tallyFour":              TallyFour,
	"tallyOne":               TallyOne,
	"tallyThree":             TallyThree,
	"tallyTwo":               TallyTwo,
	"target":                 Target,
	"telephone":              Telephone,
	"telephoneCall":          TelephoneCall,
	"telephoneForward":       TelephoneForward,
	"telephoneIn":            TelephoneIn,
	"telephoneMissed":        TelephoneMissed,
	"telephoneOut":           TelephoneOut,
	"telephoneSlash":         TelephoneSlash,
	"terminal":               Terminal,
	"textAlignCenter":        TextAlignCenter,
	"textAlignLeft":          TextAlignLeft,
	"textAlignRight":         TextAlignRight,
	"textJustify":            TextJustify,
	"thermometer":            Thermometer,
	"three":                  Three,
	"threeCircle":            ThreeCircle,
	"threeDiamond":           ThreeDiamond,
	"threeHexagon":           ThreeHexagon,
	"threeOctagon":           ThreeOctagon,
	"threeSquare":            ThreeSquare,
	"threeWaves":             ThreeWaves,
	"ticket":                 Ticket,
	"ticketSlash":            TicketSlash,
	"toggleLeft":             ToggleLeft,
	"toggleRight":            ToggleRight,
	"tool":                   Tool,
	"train":                  Train,
	"trash":                  Trash,
	"tree":                   Tree,
	"trendingDown":           TrendingDown,
	"trendingUp":             TrendingUp,
	"triangle":               Triangle,
	"triangleDanger":         TriangleDanger,
	"triangleInfo":           TriangleInfo,
	"truck":                  Truck,
	"tv":                     Tv,
	"two":                    Two,
	"twoCircle":              TwoCircle,
	"twoDiamond":             TwoDiamond,
	"twoHexagon":             TwoHexagon,
	"twoOctagon":             TwoOctagon,
	"twoSquare":              TwoSquare,
	"twoWaves":               TwoWaves,
	"typeBold":               TypeBold,
	"typeItalic":             TypeItalic,
	"typeText":               TypeText,
	"typeUnderline":          TypeUnderline,
	"umbrella":               Umbrella,
	"undo":                   Undo,
	"union":                  Union,
	"unlink":                 Unlink,
	"upload":                 Upload,
	"user":                   User,
	"userCheck":              UserCheck,
	"userCircle":             UserCircle,
	"userHexagon":            UserHexagon,
	"userMinus":              UserMinus,
	"userOctagon":            UserOctagon,
	"userPlus":               UserPlus,
	"userSettings":           UserSettings,
	"userSquare":             UserSquare,
	"userWaves":              UserWaves,
	"userX":                  UserX,
	"users":                  Users,
	"usersGroup":             UsersGroup,
	"video":                  Video,
	"videoSlash":             VideoSlash,
	"volumeCheck":            VolumeCheck,
	"volumeHigh":             VolumeHigh,
	"volumeLow":              VolumeLow,
	"volumeMinus":            VolumeMinus,
	"volumeNone":             VolumeNone,
	"volumePlus":             VolumePlus,
	"volumeSlash":            VolumeSlash,
	"volumeX":                VolumeX,
	"watch":                  Watch,
	"waves":                  Waves,
	"webcam":                 Webcam,
	"wheel":                  Wheel,
	"wifi":                   Wifi,
	"wifiCheck":              WifiCheck,
	"wifiLow":                WifiLow,
	"wifiMedium":             WifiMedium,
	"wifiMinus":              WifiMinus,
	"wifiPlus":               WifiPlus,
	"wifiSlash":              WifiSlash,
	"wifiX":                  WifiX,
	"winds":                  Winds,
	"wine":                   Wine,
	"winkCircle":             WinkCircle,
	"winkGhost":              WinkGhost,
	"winkSquare":             WinkSquare,
	"wrench":                 Wrench,
	"x":                      X,
	"xcircle":                Xcircle,
	"xhexagon":               Xhexagon,
	"xoctagon":               Xoctagon,
	"xsquare":                Xsquare,
	"xtriangle":              Xtriangle,
	"xwaves":                 Xwaves,
	"yen":                    Yen,
	"yenCircle":              YenCircle,
	"yenHexagon":             YenHexagon,
	"yenOctagon":             YenOctagon,
	"yenSquare":              YenSquare,
	"yenWaves":               YenWaves,
	"zero":                   Zero,
	"zeroCircle":             ZeroCircle,
	"zeroDiamond":            ZeroDiamond,
	"zeroHexagon":            ZeroHexagon,
	"zeroOctagon":            ZeroOctagon,
	"zeroSquare":             ZeroSquare,
	"zeroWaves":              ZeroWaves,
}

func AcademicHat(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.217 3.5a5.17 5.17 0 0 0-4.434 0L3.092 6.637c-1.456.682-1.456 3.044 0 3.726l6.69 3.137a5.17 5.17 0 0 0 4.435 0l6.691-3.137c1.456-.682 1.456-3.044 0-3.726zM22 8.5v5"/><path d="M5 11.5v5.125C5 19.543 9.694 21 12 21s7-1.457 7-4.375V11.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Accessibility(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="m8 9.5l4 .778m0 0l4-.778m-4 .778v2.333m0 0L10.4 16.5m1.6-3.889l1.6 3.889M12 7.25v.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Activity(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 12h-2.962c-.21 0-.316 0-.405.042a.51.51 0 0 0-.201.173c-.061.088-.092.205-.155.44l-1.817 6.846c-.233.875-.349 1.313-.524 1.426a.435.435 0 0 1-.485-.002c-.175-.115-.288-.554-.514-1.43l-3.873-14.99c-.227-.876-.34-1.315-.515-1.43a.435.435 0 0 0-.485-.002c-.175.113-.291.55-.524 1.426l-1.817 6.845c-.063.236-.094.353-.154.44a.51.51 0 0 1-.202.174C6.278 12 6.173 12 5.962 12H3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ActivitySquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17 12h-2l-2 5l-2-10l-2 5H7"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AddQueue(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 9v10.4c0 .56 0 .84.109 1.054a1 1 0 0 0 .437.437C3.76 21 4.04 21 4.598 21H15m-1-8v-3m0 0V7m0 3h-3m3 0h3M7 13.8V6.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C8.52 3 9.08 3 10.2 3h7.6c1.12 0 1.68 0 2.108.218a2 2 0 0 1 .874.874C21 4.52 21 5.08 21 6.2v7.6c0 1.12 0 1.68-.218 2.108a2.001 2.001 0 0 1-.874.874c-.428.218-.986.218-2.104.218h-7.607c-1.118 0-1.678 0-2.105-.218a2 2 0 0 1-.874-.874C7 15.48 7 14.92 7 13.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Aeroplane(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m14.134 6.992l3.396-3.397a2.033 2.033 0 0 1 2.875 2.875l-3.397 3.396l1.838 9.145c.235 1.17-2.034 2.8-2.638 1.25l-2.69-6.904l-3.493 3.493c.17 2.041.207 2.72-1.224 4.15l-2.175-3.626L3 15.199c1.43-1.431 2.109-1.395 4.15-1.224l3.493-3.492l-6.904-2.691c-1.55-.604.08-2.874 1.25-2.638z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Airplay(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12.632 16.511c-.216-.284-.324-.426-.455-.477a.482.482 0 0 0-.354 0c-.13.051-.239.193-.455.477l-2.363 3.106c-.336.443-.505.664-.505.85c0 .163.07.316.19.417c.139.116.408.116.947.116h4.726c.539 0 .808 0 .947-.116a.543.543 0 0 0 .19-.416c0-.187-.169-.408-.505-.85z"/><path d="M17.4 18h.6a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3H6a3 3 0 0 0-3 3v9a3 3 0 0 0 3 3h.6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Airpods(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 18V8a3.5 3.5 0 1 0-3.5 3.5c.274-.006.5.214.5.488V18a1.5 1.5 0 0 0 3 0m4 0V8a3.5 3.5 0 1 1 3.5 3.5a.489.489 0 0 0-.5.488V18a1.5 1.5 0 0 1-3 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Alarm(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 5.231L6.15 3M21 5.231L17.85 3"/><circle cx="12" cy="13" r="8"/><path d="M12 8.5v5l3 2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlarmCheck(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 5.231L6.15 3M21 5.231L17.85 3"/><circle cx="12" cy="13" r="8"/><path d="m10 13.323l1.379 1.575a.299.299 0 0 0 .466-.022l2.8-3.876"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlarmMinus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 5.231L6.15 3M21 5.231L17.85 3"/><circle cx="12" cy="13" r="8"/><path d="M9.5 13h5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlarmPlus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 5.231L6.15 3M21 5.231L17.85 3"/><circle cx="12" cy="13" r="8"/><path d="M9.5 13h5M12 10.5v5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlarmSnooze(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 5.231L6.15 3M21 5.231L17.85 3"/><circle cx="12" cy="13" r="8"/><path d="M10.75 11h2.5l-2.5 4h2.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlarmX(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 5.231L6.15 3M21 5.231L17.85 3"/><circle cx="12" cy="13" r="8"/><path d="m10 11l4 4m0-4l-4 4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignBottom(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 14V6c0-.932 0-1.398-.152-1.765a2 2 0 0 0-1.083-1.083C8.398 3 7.932 3 7 3c-.932 0-1.398 0-1.765.152a2 2 0 0 0-1.083 1.083C4 4.602 4 5.068 4 6v8c0 .932 0 1.398.152 1.765a2 2 0 0 0 1.083 1.083C5.602 17 6.068 17 7 17c.932 0 1.398 0 1.765-.152a2 2 0 0 0 1.083-1.083C10 15.398 10 14.932 10 14m10 0v-4c0-.932 0-1.398-.152-1.765a2 2 0 0 0-1.083-1.083C18.398 7 17.932 7 17 7c-.932 0-1.398 0-1.765.152a2 2 0 0 0-1.083 1.083C14 8.602 14 9.068 14 10v4c0 .932 0 1.398.152 1.765a2 2 0 0 0 1.083 1.083C15.602 17 16.068 17 17 17c.932 0 1.398 0 1.765-.152a2 2 0 0 0 1.083-1.083C20 15.398 20 14.932 20 14m1 7H3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignHorizontal(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 21v-1m0-17v1m0 10H8c-.932 0-1.398 0-1.765.152a2 2 0 0 0-1.083 1.083C5 15.602 5 16.068 5 17c0 .932 0 1.398.152 1.765a2 2 0 0 0 1.083 1.083C6.602 20 7.068 20 8 20h4m0-6h4c.932 0 1.398 0 1.765.152a2 2 0 0 1 1.083 1.083C19 15.602 19 16.068 19 17c0 .932 0 1.398-.152 1.765a2 2 0 0 1-1.083 1.083C17.398 20 16.932 20 16 20h-4m0-6v-4m0 0h2c.932 0 1.398 0 1.765-.152a2 2 0 0 0 1.083-1.083C17 8.398 17 7.932 17 7c0-.932 0-1.398-.152-1.765a2 2 0 0 0-1.083-1.083C15.398 4 14.932 4 14 4h-2m0 6h-2c-.932 0-1.398 0-1.765-.152a2 2 0 0 1-1.083-1.083C7 8.398 7 7.932 7 7c0-.932 0-1.398.152-1.765a2 2 0 0 1 1.083-1.083C8.602 4 9.068 4 10 4h2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignLeft(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 14h-8c-.932 0-1.398 0-1.765.152a2 2 0 0 0-1.083 1.083C7 15.602 7 16.068 7 17c0 .932 0 1.398.152 1.765a2 2 0 0 0 1.083 1.083C8.602 20 9.068 20 10 20h8c.932 0 1.398 0 1.765-.152a2 2 0 0 0 1.083-1.083C21 18.398 21 17.932 21 17c0-.932 0-1.398-.152-1.765a2 2 0 0 0-1.083-1.083C19.398 14 18.932 14 18 14M14 4h-4c-.932 0-1.398 0-1.765.152a2 2 0 0 0-1.083 1.083C7 5.602 7 6.068 7 7c0 .932 0 1.398.152 1.765a2 2 0 0 0 1.083 1.083C8.602 10 9.068 10 10 10h4c.932 0 1.398 0 1.765-.152a2 2 0 0 0 1.083-1.083C17 8.398 17 7.932 17 7c0-.932 0-1.398-.152-1.765a2 2 0 0 0-1.083-1.083C15.398 4 14.932 4 14 4M3 21V3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignRight(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14 14H6c-.932 0-1.398 0-1.765.152a2 2 0 0 0-1.083 1.083C3 15.602 3 16.068 3 17c0 .932 0 1.398.152 1.765a2 2 0 0 0 1.083 1.083C4.602 20 5.068 20 6 20h8c.932 0 1.398 0 1.765-.152a2 2 0 0 0 1.083-1.083C17 18.398 17 17.932 17 17c0-.932 0-1.398-.152-1.765a2 2 0 0 0-1.083-1.083C15.398 14 14.932 14 14 14m0-10h-4c-.932 0-1.398 0-1.765.152a2 2 0 0 0-1.083 1.083C7 5.602 7 6.068 7 7c0 .932 0 1.398.152 1.765a2 2 0 0 0 1.083 1.083C8.602 10 9.068 10 10 10h4c.932 0 1.398 0 1.765-.152a2 2 0 0 0 1.083-1.083C17 8.398 17 7.932 17 7c0-.932 0-1.398-.152-1.765a2 2 0 0 0-1.083-1.083C15.398 4 14.932 4 14 4m7 17V3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignTop(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 18v-8c0-.932 0-1.398-.152-1.765a2 2 0 0 0-1.083-1.083C8.398 7 7.932 7 7 7c-.932 0-1.398 0-1.765.152a2 2 0 0 0-1.083 1.083C4 8.602 4 9.068 4 10v8c0 .932 0 1.398.152 1.765a2 2 0 0 0 1.083 1.083C5.602 21 6.068 21 7 21c.932 0 1.398 0 1.765-.152a2 2 0 0 0 1.083-1.083C10 19.398 10 18.932 10 18m10-4v-4c0-.932 0-1.398-.152-1.765a2 2 0 0 0-1.083-1.083C18.398 7 17.932 7 17 7c-.932 0-1.398 0-1.765.152a2 2 0 0 0-1.083 1.083C14 8.602 14 9.068 14 10v4c0 .932 0 1.398.152 1.765a2 2 0 0 0 1.083 1.083C15.602 17 16.068 17 17 17c.932 0 1.398 0 1.765-.152a2 2 0 0 0 1.083-1.083C20 15.398 20 14.932 20 14m1-11H3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignVertical(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 12h1m17 0h-1m-10 0V8c0-.932 0-1.398-.152-1.765a2 2 0 0 0-1.083-1.083C8.398 5 7.932 5 7 5c-.932 0-1.398 0-1.765.152a2 2 0 0 0-1.083 1.083C4 6.602 4 7.068 4 8v4m6 0v4c0 .932 0 1.398-.152 1.765a2 2 0 0 1-1.083 1.083C8.398 19 7.932 19 7 19c-.932 0-1.398 0-1.765-.152a2 2 0 0 1-1.083-1.083C4 17.398 4 16.932 4 16v-4m6 0h4m0 0v2c0 .932 0 1.398.152 1.765a2 2 0 0 0 1.083 1.083C15.602 17 16.068 17 17 17c.932 0 1.398 0 1.765-.152a2 2 0 0 0 1.083-1.083C20 15.398 20 14.932 20 14v-2m-6 0v-2c0-.932 0-1.398.152-1.765a2 2 0 0 1 1.083-1.083C15.602 7 16.068 7 17 7c.932 0 1.398 0 1.765.152a2 2 0 0 1 1.083 1.083C20 8.602 20 9.068 20 10v2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Alt(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 5.25h5.625l6.75 13.5H21m-6.75-13.5H21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Anchor(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 12H3a9 9 0 1 0 18 0h-2m-7-5v14m0-14a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Aperture(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m14.079 8.4l5.166 8.946M9.921 8.4h10.332M7.842 12l5.166-8.946M9.921 15.6L4.755 6.654m9.324 8.946H3.747M16.158 12l-5.166 8.946"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Api(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.5 13L7 11.5l5.5 5.5l-1.5 1.5c-.75.75-3.5 2-5.5 0s-.75-4.75 0-5.5M3 21l2.5-2.5m13-7.5L17 12.5L11.5 7L13 5.5c.75-.75 3.5-2 5.5 0s.75 4.75 0 5.5m-6-3l-2 2M21 3l-2.5 2.5m-2.5 6l-2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ar(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.5 20.886c-1.463-.144-2.447-.47-3.182-1.204c-.735-.735-1.06-1.72-1.204-3.182M7.5 3.114c-1.463.144-2.447.47-3.182 1.204c-.735.735-1.06 1.72-1.204 3.182M16.5 3.114c1.463.144 2.447.47 3.182 1.204c.735.735 1.06 1.72 1.204 3.182M16.5 20.886c1.463-.144 2.447-.47 3.182-1.204c.735-.735 1.06-1.72 1.204-3.182M8 10l4 2m-4-2v4l4 2m-4-6l4-2l4 2m-4 2l4-2m-4 2v4m4-6v4l-4 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Archive(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.5 11.5h3M20 8v11a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V8m17 0V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDiagonalOne(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4.929 12.707l7.778-7.778m0 0v4.95m0-4.95h-4.95m11.314 6.364l-7.778 7.778m0 0h4.95m-4.95 0v-4.95"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDiagonalTwo(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m11.293 4.929l7.778 7.778m0 0h-4.95m4.95 0v-4.95m-6.364 11.314L4.93 11.293m0 0v4.95m0-4.95h4.95"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDown(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 4.5v15m0 0l-6-5.625m6 5.625l6-5.625"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 7.5v9m3.5-3.5L12 16.5L8.5 13"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownLeft(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m17.5 6.5l-11 11m0 0h9m-9 0v-9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownLeftCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m15.182 8.818l-6.364 6.364m4.95 0h-4.95v-4.95"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownLeftSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m15.182 8.818l-6.364 6.364m4.95 0h-4.95v-4.95"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownLeftWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m15.182 8.818l-6.364 6.364m4.95 0h-4.95v-4.95"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownRight(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m6.5 6.5l11 11m0 0h-9m9 0v-9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownRightCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8.818 8.818l6.364 6.364m0-4.95v4.95h-4.95"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownRightSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8.818 8.818l6.364 6.364m0-4.95v4.95h-4.95"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownRightWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8.818 8.818l6.364 6.364m0-4.95v4.95h-4.95"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 7.5v9m3.5-3.5L12 16.5L8.5 13"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 7.5v9m3.5-3.5L12 16.5L8.5 13"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeft(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.5 12h-15m0 0l5.625-6M4.5 12l5.625 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16.5 12h-9m3.5 3.5L7.5 12L11 8.5"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftRight(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.5 7.5h11m0 0L14 11m3.5-3.5L14 4m3.5 12.5h-11m0 0L10 20m-3.5-3.5L10 13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16.5 12h-9m3.5 3.5L7.5 12L11 8.5"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16.5 12h-9m3.5 3.5L7.5 12L11 8.5"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongDown(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8 18l4 4m0 0l4-4m-4 4V2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongDownLeft(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11 19H5m0 0v-6m0 6L19 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongDownRight(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 13v6m0 0h-6m6 0L5 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongLeft(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m6 8l-4 4m0 0l4 4m-4-4h20"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongRight(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m18 8l4 4m0 0l-4 4m4-4H2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongUp(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8 6l4-4m0 0l4 4m-4-4v20"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongUpLeft(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 11V5m0 0h6M5 5l14 14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLongUpRight(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13 5h6m0 0v6m0-6L5 19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRight(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.5 12h15m0 0l-5.625-6m5.625 6l-5.625 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.5 12h9M13 8.5l3.5 3.5l-3.5 3.5"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.5 12h9M13 8.5l3.5 3.5l-3.5 3.5"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.5 12h9M13 8.5l3.5 3.5l-3.5 3.5"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUp(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 19.5v-15m0 0l-6 5.625M12 4.5l6 5.625"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 16.5v-9M8.5 11L12 7.5l3.5 3.5"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpDown(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.5 17.5v-11m0 0L11 10M7.5 6.5L4 10m12.5-3.5v11m0 0L20 14m-3.5 3.5L13 14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpLeft(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m17.5 17.5l-11-11m0 0h9m-9 0v9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpLeftCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15.182 15.182L8.818 8.818m0 4.95v-4.95h4.95"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpLeftSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15.182 15.182L8.818 8.818m0 4.95v-4.95h4.95"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpLeftWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15.182 15.182L8.818 8.818m0 4.95v-4.95h4.95"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpRight(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m6.5 17.5l11-11m0 0h-9m9 0v9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpRightCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8.818 15.182l6.364-6.364m-4.95 0h4.95v4.95"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpRightSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8.818 15.182l6.364-6.364m-4.95 0h4.95v4.95"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpRightWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8.818 15.182l6.364-6.364m-4.95 0h4.95v4.95"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 16.5v-9M8.5 11L12 7.5l3.5 3.5"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 16.5v-9M8.5 11L12 7.5l3.5 3.5"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AsteriskCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M12 8.5v7m-3-1.75l6-3.5m-6 0l6 3.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AsteriskHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8M12 8.5v7m-3-1.75l6-3.5m-6 0l6 3.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AsteriskOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132zM12 8.5v7m-3-1.75l6-3.5m-6 0l6 3.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AsteriskSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12m9-3.5v7m-3-1.75l6-3.5m-6 0l6 3.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AsteriskWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83M12 8.5v7m-3-1.75l6-3.5m-6 0l6 3.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func At(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 11.996V7.998m0 3.998c0-5.157-8-5.157-8 0c0 5.167 8 5.11 8 0m0 0c0 5 5 5 5 0C21 7.027 16.97 3 12 3s-9 4.027-9 8.996c0 4.968 4.03 8.995 9 8.995c1.675.084 3.938-.421 5.776-1.831"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Atom(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15.407 8.593c4.6 4.6 6.802 9.853 4.92 11.735c-1.88 1.881-7.135-.322-11.734-4.921c-4.6-4.6-6.802-9.853-4.92-11.735c1.88-1.881 7.134.322 11.734 4.921"/><path d="M8.594 8.593c-4.6 4.6-6.803 9.853-4.921 11.735c1.881 1.881 7.135-.322 11.734-4.921c4.6-4.6 6.803-9.853 4.921-11.735c-1.881-1.881-7.135.322-11.734 4.921M11.75 12h.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bank(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 8.94V18m5-9.06V18M7 8.94V18m5.447-14.894l7.764 3.908c.944.475.608 1.907-.447 1.907H4.236c-1.055 0-1.391-1.432-.447-1.907l7.764-3.908a.994.994 0 0 1 .894 0M19.5 21h-15a1.5 1.5 0 0 1 0-3h15a1.5 1.5 0 0 1 0 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Baseball(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M3.804 9.804c5.022.94 7.697 5.573 6 10.392m10.392-6c-5.022-.94-7.697-5.573-6-10.392"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryCharging(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.5 7H18a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1h-3.5m-7-10H4a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h3.5M21 11v2m-9.999-6L8.5 12h5.002L11 17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryChargingFour(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 7H4a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V8a1 1 0 0 0-1-1m3 4v2M6.5 10v4m3-4v4m3-4v4m3-2v2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryChargingOne(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 7H4a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V8a1 1 0 0 0-1-1m3 4v2M6.5 10v4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryChargingThree(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 7H4a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V8a1 1 0 0 0-1-1m3 4v2M6.5 10v4m3-4v4m3-4v4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryChargingTwo(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 7H4a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V8a1 1 0 0 0-1-1m3 4v2M6.5 10v4m3-4v4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryCheck(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18 7H4a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V8a1 1 0 0 0-1-1m3 4v2"/><path d="m9.26 12.242l1.034 1.181c.095.109.266.1.35-.016l2.1-2.907"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryEmpty(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 7H4a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V8a1 1 0 0 0-1-1m3 4v2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryFull(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 7H4a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V8a1 1 0 0 0-1-1m3 4v2M6.5 10v4m3-4v4m3-4v4m3-4v4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryMinus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 7H4a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V8a1 1 0 0 0-1-1m-9 5h4m8-1v2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryPlus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 7H4a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V8a1 1 0 0 0-1-1m-9 5h4m-2-2v4m10-3v2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryX(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 7H4a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V8a1 1 0 0 0-1-1m-8.5 3.5l3 3m0-3l-3 3M21 11v2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bell(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18.831 10.256c0-3.701-2.46-7.256-6.813-7.256s-6.813 3.555-6.813 7.256c0 1.502-.988 2.654-1.818 3.859c-3.73 5.971 20.807 5.703 17.262 0c-.83-1.205-1.818-2.357-1.818-3.859"/><path d="M8.643 18.368C9.272 19.92 10.07 21 12 21c1.929 0 2.728-1.08 3.357-2.632"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellCheck(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18.831 10.256c0-3.701-2.46-7.256-6.813-7.256s-6.813 3.555-6.813 7.256c0 1.502-.988 2.654-1.818 3.859c-3.73 5.971 20.807 5.703 17.262 0c-.83-1.205-1.818-2.357-1.818-3.859"/><path d="M8.643 18.368C9.272 19.92 10.07 21 12 21c1.929 0 2.728-1.08 3.357-2.632M9.6 10.323l1.379 1.575a.299.299 0 0 0 .466-.022L14.245 8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellMinus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18.831 10.256c0-3.701-2.46-7.256-6.813-7.256s-6.813 3.555-6.813 7.256c0 1.502-.988 2.654-1.818 3.859c-3.73 5.971 20.807 5.703 17.262 0c-.83-1.205-1.818-2.357-1.818-3.859"/><path d="M8.643 18.368C9.272 19.92 10.07 21 12 21c1.929 0 2.728-1.08 3.357-2.632M9.5 10h5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellOn(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18.831 10.256c0-3.701-2.46-7.256-6.813-7.256s-6.813 3.555-6.813 7.256c0 1.502-.988 2.654-1.818 3.859c-3.73 5.971 20.807 5.703 17.262 0c-.83-1.205-1.818-2.357-1.818-3.859"/><path d="M12 12a2 2 0 1 0 0-4a2 2 0 0 0 0 4m-3.357 6.368C9.272 19.92 10.07 21 12 21c1.929 0 2.728-1.08 3.357-2.632"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellPlus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18.831 10.256c0-3.701-2.46-7.256-6.813-7.256s-6.813 3.555-6.813 7.256c0 1.502-.988 2.654-1.818 3.859c-3.73 5.971 20.807 5.703 17.262 0c-.83-1.205-1.818-2.357-1.818-3.859"/><path d="M8.643 18.368C9.272 19.92 10.07 21 12 21c1.929 0 2.728-1.08 3.357-2.632M9.5 10h5M12 7.5v5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellSlash(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8.643 18.368C9.272 19.92 10.07 21 12 21c1.929 0 2.728-1.08 3.357-2.632"/><path d="M17.766 6.234L21 3m-3.234 3.234c.699 1.189 1.065 2.595 1.065 4.022c0 1.502.988 2.654 1.818 3.859c2.478 3.986-8.763 5.317-14.482 3.718M17.766 6.234l-11.6 11.6M3 21l3.167-3.167M14.66 3.5a6.91 6.91 0 0 0-2.642-.5c-4.353 0-6.813 3.555-6.813 7.256c0 1.502-.988 2.654-1.818 3.859c-.473.758-.491 1.415-.17 1.973"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellSnooze(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18.831 10.256c0-3.701-2.46-7.256-6.813-7.256s-6.813 3.555-6.813 7.256c0 1.502-.988 2.654-1.818 3.859c-3.73 5.971 20.807 5.703 17.262 0c-.83-1.205-1.818-2.357-1.818-3.859"/><path d="M8.643 18.368C9.272 19.92 10.07 21 12 21c1.929 0 2.728-1.08 3.357-2.632M10.75 8h2.5l-2.5 4h2.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellX(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18.831 10.256c0-3.701-2.46-7.256-6.813-7.256s-6.813 3.555-6.813 7.256c0 1.502-.988 2.654-1.818 3.859c-3.73 5.971 20.807 5.703 17.262 0c-.83-1.205-1.818-2.357-1.818-3.859"/><path d="M8.643 18.368C9.272 19.92 10.07 21 12 21c1.929 0 2.728-1.08 3.357-2.632M10 8l4 4m0-4l-4 4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bitcoin(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.273 5.25V3m0 2.25H8.182m1.09 0h4.364m0 0V3m0 2.25c1.808 0 3.273 1.511 3.273 3.375S15.444 12 13.636 12m-4.363 9v-2.25m0 0H8.182m1.09 0h4.364m0 2.25v-2.25m0 0h1.091c1.808 0 3.273-1.511 3.273-3.375S16.535 12 14.727 12h-1.09M6 5.25h2.182m0 0V12m5.454 0H8.182M6 18.75h2.182m0 0V12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BitcoinCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 8V6m0 12v-2m-3-4h4a2 2 0 1 0 0-4H9zm0 0h5a2 2 0 1 1 0 4H9z"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BitcoinHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/><path d="M12 8V6m0 12v-2m-3-4h4a2 2 0 1 0 0-4H9zm0 0h5a2 2 0 1 1 0 4H9z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BitcoinOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/><path d="M12 8V6m0 12v-2m-3-4h4a2 2 0 1 0 0-4H9zm0 0h5a2 2 0 1 1 0 4H9z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BitcoinSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/><path d="M12 8V6m0 12v-2m-3-4h4a2 2 0 1 0 0-4H9zm0 0h5a2 2 0 1 1 0 4H9z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BitcoinWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/><path d="M12 8V6m0 12v-2m-3-4h4a2 2 0 1 0 0-4H9zm0 0h5a2 2 0 1 1 0 4H9z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bluetooth(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m6.5 7.678l10.54 8.558c.242.196.363.294.408.41c.04.104.042.216.004.32c-.042.117-.16.218-.396.42l-3.57 3.05c-.432.37-.649.556-.833.564a.526.526 0 0 1-.416-.174c-.118-.133-.118-.408-.118-.958V4.132c0-.55 0-.825.118-.958A.526.526 0 0 1 12.653 3c.184.008.4.193.833.563l3.603 3.08c.226.192.339.289.382.402c.037.1.039.21.004.31c-.04.115-.15.214-.371.412L6.5 17.283"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Boat(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4.8 17.194l-.597-3.343a2 2 0 0 1 1.326-2.246l.171-.058m0 0l5.658-1.92a2 2 0 0 1 1.284 0l5.658 1.92m-12.6 0v-3.8a2 2 0 0 1 2-2h2.5m8.1 5.8l.171.058a2 2 0 0 1 1.326 2.246l-.597 3.342m-.9-5.646v-3.8a2 2 0 0 0-2-2h-2.5m-3.6 0h3.6m-3.6 0V4.8a1.8 1.8 0 0 1 3.6 0v.947M3 20.399c.647.657 2.222.843 3.725.216c1.032-.43 2.336-.441 3.391-.071a5.83 5.83 0 0 0 3.768 0c1.055-.37 2.359-.36 3.39.07c1.504.628 3.08.442 3.726-.216"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Book(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 20.25c0 .414.336.75.75.75h10.652C17.565 21 18 20.635 18 19.4v-1.445M5 20.25A2.25 2.25 0 0 1 7.25 18h10.152c.226 0 .425-.014.598-.045M5 20.25V6.2c0-1.136-.072-2.389 1.092-2.982C6.52 3 7.08 3 8.2 3h9.2c1.236 0 1.6.437 1.6 1.6v11.8c0 .995-.282 1.425-1 1.555"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookOpen(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 9.8V20m0-10.2c0-1.704.107-3.584-1.638-4.473C9.72 5 8.88 5 7.2 5H4.6C3.364 5 3 5.437 3 6.6v8.8c0 .568-.036 1.195.546 1.491c.214.109.493.109 1.052.109h2.833c2.377 0 3.26 1.036 4.569 3m0-10.2c0-1.704-.108-3.584 1.638-4.473C14.279 5 15.12 5 16.8 5h2.6c1.235 0 1.6.436 1.6 1.6v8.8c0 .567.035 1.195-.546 1.491c-.213.109-.493.109-1.052.109h-2.833c-2.377 0-3.26 1.036-4.569 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookmark(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.527 20.841C6.861 21.274 6 20.772 6 19.952V3.942c0-.52.336-.942.75-.942h10.5c.414 0 .75.422.75.942v16.01c0 .82-.861 1.322-1.527.89l-3.946-2.562a.962.962 0 0 0-1.054 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkCheck(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.527 20.841C6.861 21.274 6 20.772 6 19.952V3.942c0-.52.336-.942.75-.942h10.5c.414 0 .75.422.75.942v16.01c0 .82-.861 1.322-1.527.89l-3.946-2.562a.962.962 0 0 0-1.054 0z"/><path d="m9.7 9.822l1.379 1.576a.299.299 0 0 0 .466-.022l2.8-3.876"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkMinus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.527 20.841C6.861 21.274 6 20.772 6 19.952V3.942c0-.52.336-.942.75-.942h10.5c.414 0 .75.422.75.942v16.01c0 .82-.861 1.322-1.527.89l-3.946-2.562a.962.962 0 0 0-1.054 0zM9.5 9.5h5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkPlus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.527 20.841C6.861 21.274 6 20.772 6 19.952V3.942c0-.52.336-.942.75-.942h10.5c.414 0 .75.422.75.942v16.01c0 .82-.861 1.322-1.527.89l-3.946-2.562a.962.962 0 0 0-1.054 0zM9.5 9.5h5M12 7v5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkX(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.527 20.841C6.861 21.274 6 20.772 6 19.952V3.942c0-.52.336-.942.75-.942h10.5c.414 0 .75.422.75.942v16.01c0 .82-.861 1.322-1.527.89l-3.946-2.562a.962.962 0 0 0-1.054 0zM10 7.5l4 4m0-4l-4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoundingBox(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 17V7m0 10a2 2 0 1 0 2 2m-2-2a2 2 0 0 1 2 2M5 7a2 2 0 0 0 2-2M5 7a2 2 0 1 1 2-2m0 0h10m0 0a2 2 0 0 0 2 2m-2-2a2 2 0 1 1 2 2m0 0v10m0 0a2 2 0 0 0-2 2m2-2a2 2 0 1 1-2 2M7 19h10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bowl(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.017 21c.555 0 1.005-.448 1.005-1v-.45c0-.307.164-.563.433-.715a9.079 9.079 0 0 0 1.944-1.471a8.954 8.954 0 0 0 2.595-5.366c.061-.549-.395-.998-.95-.998H3.956c-.555 0-1.011.45-.95.998A8.953 8.953 0 0 0 5.6 17.364a9.081 9.081 0 0 0 1.833 1.408c.33.193.55.537.548.918v.307A1.003 1.003 0 0 0 8.986 21zM6 5v2m12-2v2m-6-4v4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Box(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m12 21l8.131-4.208c.316-.164.474-.245.589-.366a1 1 0 0 0 .226-.373c.054-.159.054-.336.054-.692V7.533M12 21l-8.131-4.208c-.316-.164-.474-.245-.589-.366a1.009 1.009 0 0 1-.226-.373C3 15.894 3 15.716 3 15.359V7.533M12 21v-9.063m9-4.404l-9 4.404m9-4.404l-8.27-4.28c-.267-.138-.4-.208-.541-.235a.994.994 0 0 0-.378 0c-.14.027-.274.097-.542.235L3 7.533m0 0l9 4.404"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrandChrome(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 21a9 9 0 1 0 0-18a9 9 0 0 0 0 18"/><path d="M12 15.6a3.6 3.6 0 1 0 0-7.2a3.6 3.6 0 0 0 0 7.2m8.253-7.2H12M4.755 6.654L8.886 13.8m2.106 7.146l4.122-7.146"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrandDribbble(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 12c-1.313 0-4.936-.495-8.178.928c-3.522 1.547-6.072 3.946-7.184 5.438"/><path d="M8.625 3.654c1.409 1.3 4.482 4.61 5.625 7.896c1.143 3.286 1.566 7.326 1.827 8.476"/><path d="M3.07 10.875c1.7.102 6.2.195 9.08-1.035s5.358-3.492 6.208-4.21"/><path d="M12 21a9 9 0 1 0 0-18a9 9 0 0 0 0 18"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrandFacebook(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.5 10v4h3v7h4v-7h3l1-4h-4V8c0-.545.455-1 1-1h3V3h-3c-2.723 0-5 2.277-5 5v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrandFigma(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6 6a3 3 0 0 1 3-3h3v6H9a3 3 0 0 1-3-3m6-3h3a3 3 0 1 1 0 6h-3z"/><path d="M12 12a3 3 0 1 1 6 0a3 3 0 0 1-6 0m-6 6a3 3 0 0 1 3-3h3v3a3 3 0 0 1-6 0m0-6a3 3 0 0 1 3-3h3v6H9a3 3 0 0 1-3-3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrandFramer(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="1.5" d="m5.5 3l13 12.6h-13V9.3h13V3zM12 15.6V21l-6.5-5.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrandGithub(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15.172 15.299c1.202-.25 2.293-.682 3.14-1.316c1.448-1.084 2.188-2.758 2.188-4.411c0-1.16-.44-2.243-1.204-3.16c-.425-.511.819-3.872-.286-3.359c-1.105.514-2.725 1.198-3.574.947c-.909-.268-1.9-.416-2.936-.416c-.9 0-1.766.111-2.574.317c-1.174.298-2.296-.363-3.426-.848c-1.13-.484-.513 3.008-.849 3.422C4.921 7.38 4.5 8.44 4.5 9.572c0 1.653.895 3.327 2.343 4.41c.965.722 2.174 1.183 3.527 1.41"/><path d="M10.37 15.391c-.58.637-.869 1.24-.869 1.813V21m5.671-5.701c.549.719.823 1.364.823 1.936V21M3.5 15.668c.45.054.783.26 1 .618c.326.537 1.537 2.526 2.913 2.526H9.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrandGitlab(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.981 21L21 13.708L18.498 3l-3.015 6.497H8.997L5.507 3L3 13.708z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrandGoogle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="1.5" d="M20.839 10.38h-8.656v3.33h5.065c-.092.81-.645 2.07-1.842 2.88c-.737.54-1.842.9-3.223.9c-3.079 0-5.525-2.572-5.525-5.58c0-2.923 2.585-5.49 5.525-5.49c1.75 0 2.855.72 3.591 1.35l2.579-2.52C16.787 3.9 14.669 3 12.183 3C8.592 3 5.461 4.98 3.987 7.95a8.799 8.799 0 0 0 0 8.1C5.461 19.02 8.592 21 12.183 21c2.486 0 4.604-.81 6.078-2.16c2.4-2.1 3.095-5.427 2.578-8.46"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrandInstagram(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16.5 3h-9A4.5 4.5 0 0 0 3 7.5v9A4.5 4.5 0 0 0 7.5 21h9a4.5 4.5 0 0 0 4.5-4.5v-9A4.5 4.5 0 0 0 16.5 3"/><path d="M15.462 11.487a3.5 3.5 0 1 1-6.925 1.026a3.5 3.5 0 0 1 6.925-1.026M17 6.5h.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrandLinkedin(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16.5 3h-9A4.5 4.5 0 0 0 3 7.5v9A4.5 4.5 0 0 0 7.5 21h9a4.5 4.5 0 0 0 4.5-4.5v-9A4.5 4.5 0 0 0 16.5 3"/><path d="M8 16.375V10.75m4 5.625V13.5m0 0v-2.75m0 2.75c0-1.288 1.222-2 2.4-2c1.6 0 1.6 1.375 1.6 2.875v2m-8-8.75v.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrandPinterest(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.452 13.18c-1.108-2.262.4-6.668 5.472-5.948c5.587.794 4.581 9.478-.077 9.138c-1.474-.107-2.031-1.328-2.177-2.576m0 0c-.11-.946.017-1.907.16-2.41c.244-.857.649-.74.353.393c-.144.552-.32 1.245-.513 2.017m0 0a652.28 652.28 0 0 0-1.63 6.708"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrandPocket(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4.8 3h14.4c.477 0 .935.199 1.273.553c.337.354.527.835.527 1.336v6.667c0 2.504-.948 4.907-2.636 6.678C16.676 20.005 14.387 21 12 21a8.634 8.634 0 0 1-3.444-.719a8.984 8.984 0 0 1-2.92-2.047C3.948 16.463 3 14.06 3 11.556V4.889c0-.501.19-.982.527-1.336A1.758 1.758 0 0 1 4.8 3"/><path d="m9 10.25l3 3l3-3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrandSlack(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="1.5" d="M14.25 3.27c-.9 0-1.71.72-1.71 1.71v4.14c0 .9.72 1.71 1.71 1.71c.9 0 1.71-.72 1.71-1.71V4.89c0-.9-.72-1.62-1.71-1.62m3.69 7.47h1.44c.81 0 1.44-.63 1.44-1.44c0-.81-.63-1.44-1.44-1.44c-.81 0-1.44.63-1.44 1.44zM3 9.48c0 .9.72 1.71 1.71 1.71h4.14c.9 0 1.71-.72 1.71-1.71c0-.9-.72-1.71-1.71-1.71H4.71C3.72 7.86 3 8.58 3 9.48m7.56-5.04c0-.81-.63-1.44-1.44-1.44c-.81 0-1.44.63-1.44 1.44c0 .81.63 1.44 1.44 1.44h1.44zm-.9 16.38c.9 0 1.71-.72 1.71-1.71v-4.14c0-.9-.72-1.71-1.71-1.71c-.9 0-1.71.72-1.71 1.71v4.14c.09.9.81 1.71 1.71 1.71m-3.6-7.56H4.62c-.81 0-1.44.63-1.44 1.44c0 .81.63 1.44 1.44 1.44c.81 0 1.44-.63 1.44-1.44zM21 14.52c0-.9-.72-1.71-1.71-1.71h-4.14c-.9 0-1.71.72-1.71 1.71c0 .9.72 1.71 1.71 1.71h4.14c.9 0 1.71-.81 1.71-1.71m-7.56 3.6v1.44c0 .81.63 1.44 1.44 1.44c.81 0 1.44-.63 1.44-1.44c0-.81-.63-1.44-1.44-1.44z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrandSpotify(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path stroke-miterlimit="10" d="M6.5 9.284c3.633-1.4 7.77-.9 11 1.3m-9.688 1.8c2.725-1 5.752-.7 8.073 1m-7.165 2c2.018-.8 4.239-.5 6.055.7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrandThreads(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.77 8.515c2.23-1.812 5.444-.845 5.823 2.135c.403 3.163-.4 5.67-3.52 5.67c-2.895 0-2.806-2.52-2.806-2.52c0-2.7 4.589-3.06 7.262-1.71c4.9 3.15 1.336 8.91-4.01 8.91C8.09 21 4.5 18.75 4.5 12s3.59-9 8.02-9c3.125 0 5.944 1.626 6.98 4.879"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrandTrello(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10 7H7v9h3zm7 0h-3v5h3z"/><path d="M16.5 3h-9A4.5 4.5 0 0 0 3 7.5v9A4.5 4.5 0 0 0 7.5 21h9a4.5 4.5 0 0 0 4.5-4.5v-9A4.5 4.5 0 0 0 16.5 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrandTwitch(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.556 11.1V7.5M20 3H4v14.4h4.444V21L12 17.4h4.444L20 13.8zm-8.889 8.1V7.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrandTwitter(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 4.85c-.9.443-1.782.623-2.7.896c-1.009-1.145-2.505-1.208-3.942-.667C12.427 5.806 12 7.529 12 9.364c-2.92.075-5.521-1.262-7.2-3.618c0 0-3.764 6.723 3.6 9.95c-1.685 1.127-3.365 1.888-5.4 1.808c2.977 1.631 6.222 2.192 9.03 1.372c4.63-1.351 7.36-5.722 7.334-10.397c0-.225 1.359-2.506 1.636-3.629"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrandX(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m19 4l-5.93 6.93M5 20l5.93-6.93m0 0l5.795 6.587c.19.216.483.343.794.343h1.474c.836 0 1.307-.85.793-1.435L13.07 10.93m-2.14 2.14L4.214 5.435C3.7 4.85 4.17 4 5.007 4h1.474c.31 0 .604.127.794.343l5.795 6.587"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrandYoutube(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.5 9.908v4.184a.41.41 0 0 0 .412.408a.42.42 0 0 0 .228-.068l3.175-2.074a.405.405 0 0 0 .003-.678l-3.175-2.11a.415.415 0 0 0-.573.11a.404.404 0 0 0-.07.228"/><path d="M2 12c0-3.3 0-4.95 1.464-5.975C4.93 5 7.286 5 12 5c4.714 0 7.071 0 8.535 1.025C22 7.05 22 8.7 22 12s0 4.95-1.465 5.975C19.072 19 16.714 19 12 19s-7.071 0-8.536-1.025C2 16.95 2 15.3 2 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Briefcase(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path stroke-miterlimit="10" d="M8.308 21h7.384c3.71 0 4.375-1.45 4.569-3.213l.692-7.2c.25-2.196-.397-3.987-4.338-3.987h-9.23c-3.941 0-4.587 1.791-4.338 3.987l.692 7.2C3.933 19.55 4.598 21 8.308 21m0-14.4v-.72c0-1.593 0-2.88 2.954-2.88h1.476c2.954 0 2.954 1.287 2.954 2.88v.72"/><path stroke-miterlimit="10" d="M9.812 13.331A15.26 15.26 0 0 1 3.234 11m11 2.331A15.26 15.26 0 0 0 20.812 11"/><circle cx="12" cy="13.5" r="2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrightnessHigh(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 12h3M5 5l2.121 2.121M19 5l-2.121 2.121M5 19l2.121-2.121M19 19l-2.121-2.121M12 3v3m0 15v-3m6-6h3m-6 0a3 3 0 1 1-6 0a3 3 0 0 1 6 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrightnessLow(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15 12a3 3 0 1 1-6 0a3 3 0 0 1 6 0M3 12h1m1-7l.707.707M19 5l-.707.707M5 19l.707-.707M19 19l-.707-.707M12 3v1m0 17v-1m8-8h1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 12c0-4.243 0-5.364 1.318-6.682C5.636 4 7.758 4 12 4c4.243 0 6.364 0 7.682 1.318C21 6.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12m13.5-7V3m-9 2V3M3.25 8h17.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarCheck(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 12c0-4.243 0-5.364 1.318-6.682C5.636 4 7.758 4 12 4c4.243 0 6.364 0 7.682 1.318C21 6.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12m13.5-7V3m-9 2V3M3.25 8h17.5"/><path d="m10.258 14.242l1.034 1.181c.095.109.266.1.35-.016l2.1-2.907"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarDown(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 12c0-4.243 0-5.364 1.318-6.682C5.636 4 7.758 4 12 4c4.243 0 6.364 0 7.682 1.318C21 6.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12m13.5-7V3m-9 2V3M3.25 8h17.5m-8.747 3.5L12 16.78"/><path d="M14.5 14.847L12 17.5l-2.5-2.653"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarMinus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 12c0-4.243 0-5.364 1.318-6.682C5.636 4 7.758 4 12 4c4.243 0 6.364 0 7.682 1.318C21 6.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12m13.5-7V3m-9 2V3M3.25 8h17.5M10 14h4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarPlus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 12c0-4.243 0-5.364 1.318-6.682C5.636 4 7.758 4 12 4c4.243 0 6.364 0 7.682 1.318C21 6.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12m13.5-7V3m-9 2V3M3.25 8h17.5M10 14h4m-2-2v4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarSlash(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.318 19.682C5.636 21 7.758 21 12 21c4.243 0 6.364 0 7.682-1.318C21 18.364 21 16.242 21 12c0-4.243 0-5.364-1.318-6.682a3.346 3.346 0 0 0-.553-.447M4.318 19.682L3 21m1.318-1.318L16 8M7.5 4.114c-1.463.144-2.447.47-3.182 1.204C3 6.636 3 7.758 3 12c0 1.374 0 2.526.045 3.5M7.5 4.114V5m0-.886V3m0 1.114c2.23-.22 4.507-.137 6.746-.104m4.883.861L21 3m-1.871 1.871L16 8m0 0h4.75M3.25 8h7.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarUp(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 12c0-4.243 0-5.364 1.318-6.682C5.636 4 7.758 4 12 4c4.243 0 6.364 0 7.682 1.318C21 6.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12m13.5-7V3m-9 2V3M3.25 8h17.5m-8.747 9.5L12 12.22"/><path d="M14.5 14.154L12 11.5l-2.5 2.653"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarX(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 12c0-4.243 0-5.364 1.318-6.682C5.636 4 7.758 4 12 4c4.243 0 6.364 0 7.682 1.318C21 6.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12m13.5-7V3m-9 2V3M3.25 8h17.5M10.5 12.5l3 3m0-3l-3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Camera(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20.52 20.414c.308-.375.48-.884.48-1.414V7c0-.53-.172-1.04-.48-1.414C20.215 5.21 19.799 5 19.365 5h-8.981C8.659 5 8.325 3.269 6.827 3.026C6.563 2.983 6.289 3 6.022 3c-.953 0-1.429 0-1.804.159a2 2 0 0 0-1.059 1.06C3 4.592 3 5.068 3 6.021V19c0 .53.172 1.04.48 1.414c.306.375.722.586 1.156.586h14.728c.434 0 .85-.21 1.157-.586M16 3h3"/><path d="M13.5 17a4 4 0 1 0 0-8a4 4 0 0 0 0 8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraSlash(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 3h1.5M3 21l.528-.528M21 3l-2 2M3.528 20.472c.302.339.697.528 1.108.528h14.728c.434 0 .85-.21 1.157-.586c.307-.375.479-.884.479-1.414V7c0-.53-.172-1.04-.48-1.414C20.215 5.21 19.799 5 19.365 5H19M3.528 20.472l6.184-6.184m0 0A4.002 4.002 0 0 0 17.5 13a4.002 4.002 0 0 0-2.712-3.788m-5.076 5.076l5.076-5.076m0 0L19 5m-4.308 0h-4.309C8.659 5 8.325 3.269 6.827 3.026C6.563 2.983 6.289 3 6.022 3c-.953 0-1.429 0-1.804.159a2 2 0 0 0-1.059 1.06C3 4.592 3 5.068 3 6.021v9.734"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Campfire(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m5.422 20.706l13.156-4.789m0 4.789L5.422 15.917m10.963-4.349A4.751 4.751 0 0 1 12 14.5c-2.623 0-4.75-2.134-4.75-4.767c0-2.632.998-3.709 2.558-6.233c2.923 1.283 2.923 5.133 2.923 5.133s.96-1.856 2.923-2.75c.63 1.86 1.478 3.89.731 5.685"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cart(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16.5 21a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m-8 0a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3M3.71 5.4h15.214c1.378 0 2.373 1.27 1.995 2.548l-1.654 5.6C19.01 14.408 18.196 15 17.27 15H8.112c-.927 0-1.742-.593-1.996-1.452zm0 0L3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CartCheck(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m11 10.242l1.034 1.181c.095.109.266.1.35-.016l2.1-2.907M16.5 21a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m-8 0a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3"/><path d="M3.71 5.4h15.214c1.378 0 2.373 1.27 1.995 2.548l-1.654 5.6C19.01 14.408 18.196 15 17.27 15H8.112c-.927 0-1.742-.593-1.996-1.452zm0 0L3 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CartMinus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16.5 21a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m-8 0a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3M3.71 5.4h15.214c1.378 0 2.373 1.27 1.995 2.548l-1.654 5.6C19.01 14.408 18.196 15 17.27 15H8.112c-.927 0-1.742-.593-1.996-1.452zm0 0L3 3m7.5 7h4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CartPlus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.5 10h4m-2-2v4m4 9a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m-8 0a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3M3.71 5.4h15.214c1.378 0 2.373 1.27 1.995 2.548l-1.654 5.6C19.01 14.408 18.196 15 17.27 15H8.112c-.927 0-1.742-.593-1.996-1.452zm0 0L3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CartX(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m11 8.5l3 3m0-3l-3 3m5.5 9.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m-8 0a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3M3.71 5.4h15.214c1.378 0 2.373 1.27 1.995 2.548l-1.654 5.6C19.01 14.408 18.196 15 17.27 15H8.112c-.927 0-1.742-.593-1.996-1.452zm0 0L3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CastScreen(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 8.028c0-.956 0-1.434.108-1.826a3 3 0 0 1 2.094-2.094C5.594 4 6.072 4 7.028 4H16.2c1.68 0 2.52 0 3.162.327a3 3 0 0 1 1.311 1.311C21 6.28 21 7.12 21 8.8v6.4c0 1.68 0 2.52-.327 3.162a3 3 0 0 1-1.311 1.311C18.72 20 17.88 20 16.2 20h-2.053"/><path d="M11 20a8 8 0 0 0-8-8m0 8.004l.353-.354M7 20a4 4 0 0 0-4-4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CenterFocus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.5 20.886c-1.463-.144-2.447-.47-3.182-1.204c-.735-.735-1.06-1.72-1.204-3.182M7.5 3.114c-1.463.144-2.447.47-3.182 1.204c-.735.735-1.06 1.72-1.204 3.182M16.5 3.114c1.463.144 2.447.47 3.182 1.204c.735.735 1.06 1.72 1.204 3.182M16.5 20.886c1.463-.144 2.447-.47 3.182-1.204c.735-.735 1.06-1.72 1.204-3.182"/><path stroke-miterlimit="1" d="M15 12a3 3 0 1 0-6 0a3 3 0 0 0 6 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartBar(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 21V3m-5 18V9M7 21v-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartBarOne(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 8v13M12 3v18m-6-9v9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartBarTwo(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15 9.429V5a2 2 0 0 0-2-2h-2a2 2 0 0 0-2 2v8.286m6-3.857V21m0-11.571h4a2 2 0 0 1 2 2V19a2 2 0 0 1-2 2h-4m0 0H9m0 0v-7.714M9 21H5a2 2 0 0 1-2-2v-3.714a2 2 0 0 1 2-2h4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartBubble(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 8.863a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0M13 8A5 5 0 1 1 3 8a5 5 0 0 1 10 0m5.969 9.5a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartGraph(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 3v18h18"/><path d="m19.8 7.8l-6 6l-3-3L6.6 15"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartLine(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 16.5L9 10l4 6l8-9.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartPie(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 21a9 9 0 0 0 9-9h-9V3a9 9 0 0 0 0 18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartPieOne(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 12a9 9 0 0 1-9 9m9-9a9 9 0 0 0-9-9m9 9h-9m0 9a9 9 0 1 1 0-18m0 18v-9m0-9v9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartPieTwo(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 3a9 9 0 1 0 6.364 15.364M12 3a9 9 0 0 1 6.364 15.364M12 3v9l6.364 6.364"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chat(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 21a9 9 0 1 0-9-9c0 1.44.338 2.8.94 4.007c.453.911-.177 2.14-.417 3.037a1.17 1.17 0 0 0 1.433 1.433c.897-.24 2.126-.87 3.037-.416A8.964 8.964 0 0 0 12 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatCheck(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m9.5 12.323l1.379 1.575a.299.299 0 0 0 .466-.022l2.8-3.876"/><path d="M12 21a9 9 0 1 0-9-9c0 1.44.338 2.8.94 4.007c.453.911-.177 2.14-.417 3.037a1.17 1.17 0 0 0 1.433 1.433c.897-.24 2.126-.87 3.037-.416A8.964 8.964 0 0 0 12 21"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatDots(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 12.75v-.5m4 .5v-.5m-8 .5v-.5M12 21a9 9 0 1 0-9-9c0 1.44.338 2.8.94 4.007c.453.911-.177 2.14-.417 3.037a1.17 1.17 0 0 0 1.433 1.433c.897-.24 2.126-.87 3.037-.416A8.964 8.964 0 0 0 12 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatMessages(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 12a5 5 0 0 1 7 7m-7-7a4.993 4.993 0 0 0-2 4a5 5 0 0 0 .224 1.483c.272.88.076 1.86-.099 2.784a.468.468 0 0 0 .592.539c.848-.232 1.691-.43 2.557-.112A4.99 4.99 0 0 0 8 21a4.991 4.991 0 0 0 4-2m-7-7c0-4.685 2.875-9 8-9a8 8 0 0 1 7.532 10.7c-.476 1.326.037 3.102.337 4.568a.451.451 0 0 1-.584.526c-1.312-.41-2.852-.986-4.085-.466c-1.334.562-2.736.672-4.2.672"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatMinus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.5 12h5M12 21a9 9 0 1 0-9-9c0 1.44.338 2.8.94 4.007c.453.911-.177 2.14-.417 3.037a1.17 1.17 0 0 0 1.433 1.433c.897-.24 2.126-.87 3.037-.416A8.964 8.964 0 0 0 12 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatPlus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.5 12h5M12 9.5v5m0 6.5a9 9 0 1 0-9-9c0 1.44.338 2.8.94 4.007c.453.911-.177 2.14-.417 3.037a1.17 1.17 0 0 0 1.433 1.433c.897-.24 2.126-.87 3.037-.416A8.964 8.964 0 0 0 12 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatX(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m10 10l4 4m0-4l-4 4m2 7a9 9 0 1 0-9-9c0 1.44.338 2.8.94 4.007c.453.911-.177 2.14-.417 3.037a1.17 1.17 0 0 0 1.433 1.433c.897-.24 2.126-.87 3.037-.416A8.964 8.964 0 0 0 12 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Check(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m6 13.626l1.606 1.722c.886.95 1.329 1.424 1.825 1.574c.436.131.9.096 1.315-.1c.473-.224.852-.761 1.612-1.836L18 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8.667 12.633l1.505 1.721a1 1 0 0 0 1.564-.073L15.333 9.3"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckCircleOne(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m9 8.806l3.562 3.94a.788.788 0 0 0 1.206-.055L21 3"/><path d="M21 12a9 9 0 1 1-9-9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8.667 12.633l1.505 1.721a1 1 0 0 0 1.564-.073L15.333 9.3"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/><path d="m8.667 12.633l1.505 1.721a1 1 0 0 0 1.564-.073L15.333 9.3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/><path d="m8.667 12.633l1.505 1.721a1 1 0 0 0 1.564-.073L15.333 9.3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckSquareOne(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m9 8.806l3.562 3.94a.788.788 0 0 0 1.206-.055L21 3"/><path d="M21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8.667 12.633l1.505 1.721a1 1 0 0 0 1.564-.073L15.333 9.3"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleDown(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m6 6l6 6l6-6M6 12l6 6l6-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleDownLeft(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6.257 9.257v8.486h8.486"/><path d="M10.257 5.257v8.486h8.486"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleDownRight(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.257 17.743h8.486V9.257"/><path d="M5.257 13.743h8.486V5.257"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleLeft(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m12 6l-6 6l6 6m6-12l-6 6l6 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleRight(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m12 18l6-6l-6-6M6 18l6-6l-6-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleUp(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m18 12l-6-6l-6 6m12 6l-6-6l-6 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleUpLeft(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.743 6.257H6.257v8.486"/><path d="M18.743 10.257h-8.486v8.486"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDoubleUpRight(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17.743 14.743V6.257H9.257"/><path d="M13.743 18.743v-8.486H5.257"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDown(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m6 9l6 6l6-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDownCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m15.5 10.75l-3.5 3.5l-3.5-3.5"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDownLeft(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.879 7.636v8.485h8.485"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDownLeftCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.95 13.95H10V9"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDownLeftSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.95 13.95H10V9"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDownLeftWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.95 13.95H10V9"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDownRight(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16.121 7.636v8.485H7.636"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDownRightCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13.95 9v4.95H9"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDownRightSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13.95 9v4.95H9"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDownRightWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13.95 9v4.95H9"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDownSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m15.5 10.75l-3.5 3.5l-3.5-3.5"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDownWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m15.5 10.75l-3.5 3.5l-3.5-3.5"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeft(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m15 6l-6 6l6 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeftCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13.25 15.5L9.75 12l3.5-3.5"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeftSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13.25 15.5L9.75 12l3.5-3.5"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeftWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13.25 15.5L9.75 12l3.5-3.5"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRight(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9 18l6-6l-6-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRightCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m10.75 8.5l3.5 3.5l-3.5 3.5"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRightSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m10.75 8.5l3.5 3.5l-3.5 3.5"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRightWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m10.75 8.5l3.5 3.5l-3.5 3.5"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUp(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m18 15l-6-6l-6 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUpCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8.5 13.25l3.5-3.5l3.5 3.5"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUpDown(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8 16l4 4l4-4M8 8l4-4l4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUpLeft(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.879 16.364V7.879h8.485"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUpLeftCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10 14.95V10h4.95"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUpLeftSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10 14.95V10h4.95"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUpLeftWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10 14.95V10h4.95"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUpRight(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16.121 16.364V7.879H7.636"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUpRightCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 10h4.95v4.95"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUpRightSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 10h4.95v4.95"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUpRightWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 10h4.95v4.95"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUpSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8.5 13.25l3.5-3.5l3.5 3.5"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUpWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8.5 13.25l3.5-3.5l3.5 3.5"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chip(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 3H7a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1M6 6H3m3 4H3m3 4H3m3 4H3M21 6h-3m3 4h-3m3 4h-3m3 4h-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Circle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="12" cy="12" r="9" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleDashed(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-dasharray="3.5 3.5" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleHalf(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 21a9 9 0 1 0 0-18m0 18a9 9 0 1 1 0-18m0 18V3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleHalfCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<defs><path id="mynauiCircleHalfCircle0" d="M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0"/></defs><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><use href="#mynauiCircleHalfCircle0"/><use href="#mynauiCircleHalfCircle0"/><path d="M12 17a5 5 0 0 0 0-10z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleNotch(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.717 4A9.049 9.049 0 0 0 3 11.956C3 16.951 7.03 21 12 21s9-4.05 9-9.044A9.049 9.049 0 0 0 16.283 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Click(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.92 6.956L5.45 5.473m9.309 1.483l1.47-1.483m-10.78 10.88l1.47-1.484m3.92-9.89V3m-5.88 7.913H3m13.875 5.923l3.814-1.506a.496.496 0 0 0 0-.921l-9.165-3.615a.492.492 0 0 0-.635.64l3.582 9.251c.162.42.75.42.912 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clipboard(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 5a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v2H9z"/><path d="M9 5H7a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2h-2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 6v6l4 2"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockEight(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 6v6l-4 2"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockEleven(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 6v6L9.5 8"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockFive(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 6v6l2.5 4"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockFour(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 6v6l4 2"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockHand(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.147 3v9l6 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/><path d="M12 6v6l4 2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockNine(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 6v6H7.5"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/><path d="M12 6v6l4 2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockOne(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 6v6l2.5-4"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockSeven(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 6v6l-2.5 4"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockSix(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 6v10.5"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/><path d="M12 6v6l4 2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockTen(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 6v6l-4-2"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockThree(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 6v6h4.5"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockTwelve(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 6v6"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockTwo(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 6v6l4-2"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/><path d="M12 6v6l4 2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloud(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20.278 17.497c3.678-3.154-.214-7.384-4.256-7.384C13.175-.969-3.526 8.197 3.875 16.55"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudDownload(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m11.966 11.136l-.004 8M19.825 17c4.495-3.16.475-7.73-3.706-7.73C13.296-1.732-3.265 7.368 4.074 15.662m11.07 1.156L11.962 20L8.78 16.818"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudLightning(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19.825 17c4.495-3.16.475-7.73-3.706-7.73C13.296-1.732-3.265 7.368 4.074 15.662"/><path d="M11.501 11L9 16h5.002L11.5 21"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudRain(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.004 19L12 14m4.004 7L16 16m-7.996 1L8 12m11.825 5c4.495-3.16.475-7.73-3.706-7.73C13.296-1.732-3.265 7.368 4.074 15.662"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudSlash(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m2.875 21l2.793-2.793M20.875 3l-5.532 5.532m0 0c.274.536.503 1.145.68 1.831c4.041 0 7.933 4.23 4.255 7.384c-.91.78-2.245 1.003-3.406 1.003H8.026c-.827 0-1.632-.19-2.358-.543m9.675-9.675l-9.675 9.675M3.24 16C-1.307 9.542 7.728 3 13 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudSnow(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.004 17.5L12 17m4.004-1.5L16 15m-7.996.5L8 15m4.004 6L12 20.5m4.004-1.5L16 18.5m-7.996.5L8 18.5M19.825 17c4.495-3.16.475-7.73-3.706-7.73C13.296-1.732-3.265 7.368 4.074 15.662"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudUpload(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m11.966 20l-.004-8m7.863 5c4.495-3.16.475-7.73-3.706-7.73C13.296-1.732-3.265 7.368 4.074 15.662"/><path d="m15.144 14.318l-3.182-3.182l-3.182 3.182"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cocktail(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m12 14.25l7.777-9.625C20.306 3.97 19.835 3 18.988 3H5.012c-.847 0-1.318.97-.789 1.625zm0 0V21M6.546 7.5h10.908M7.329 21h9.342"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Code(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m17.09 7.974l.23.23c1.789 1.79 2.683 2.684 2.683 3.796c0 1.112-.894 2.007-2.683 3.796l-.23.23M13.876 5l-3.751 14M6.91 7.974l-.23.23C4.892 9.994 3.997 10.888 3.997 12c0 1.112.895 2.007 2.685 3.796l.23.23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m14.908 9.7l.132.131c1.022 1.023 1.534 1.534 1.534 2.169s-.512 1.147-1.534 2.17l-.132.13M13.072 8l-2.143 8M9.092 9.7l-.132.131C7.938 10.854 7.427 11.365 7.427 12s.51 1.147 1.533 2.17l.132.13"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m14.908 9.7l.132.131c1.022 1.023 1.534 1.534 1.534 2.169s-.512 1.147-1.534 2.17l-.132.13M13.072 8l-2.143 8M9.092 9.7l-.132.131C7.938 10.854 7.427 11.365 7.427 12s.51 1.147 1.533 2.17l.132.13"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/><path d="m14.908 9.7l.132.131c1.022 1.023 1.534 1.534 1.534 2.169s-.512 1.147-1.534 2.17l-.132.13M13.072 8l-2.143 8M9.092 9.7l-.132.131C7.938 10.854 7.427 11.365 7.427 12s.51 1.147 1.533 2.17l.132.13"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m14.908 9.7l.132.131c1.022 1.023 1.534 1.534 1.534 2.169s-.512 1.147-1.534 2.17l-.132.13M13.072 8l-2.143 8M9.092 9.7l-.132.131C7.938 10.854 7.427 11.365 7.427 12s.51 1.147 1.533 2.17l.132.13"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CodeWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m14.908 9.7l.132.131c1.022 1.023 1.534 1.534 1.534 2.169s-.512 1.147-1.534 2.17l-.132.13M13.072 8l-2.143 8M9.092 9.7l-.132.131C7.938 10.854 7.427 11.365 7.427 12s.51 1.147 1.533 2.17l.132.13"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Coffee(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 10h1.5a2.5 2.5 0 0 1 0 5H18m0-5c0-.53-.36-1-.923-1H4.923A.923.923 0 0 0 4 9.923V17a4 4 0 0 0 4 4h6a4 4 0 0 0 4-4v-2m0-5v5M16 3l-2 3m-1-3l-2 3m-1-3L8 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cog(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="1" stroke-width="1.5"><path d="M10.11 3.9a1 1 0 0 1 .995-.9h1.79a1 1 0 0 1 .995.9l.033.333a7.953 7.953 0 0 1 2.209.915l.259-.212a1 1 0 0 1 1.34.067l1.266 1.266a1 1 0 0 1 .067 1.34l-.212.26c.409.676.72 1.419.915 2.208l.332.033a1 1 0 0 1 .901.995v1.79a1 1 0 0 1-.9.995l-.333.033a7.951 7.951 0 0 1-.915 2.209l.212.259a1 1 0 0 1-.067 1.34l-1.266 1.266a1 1 0 0 1-1.34.067l-.26-.212a7.947 7.947 0 0 1-2.208.915l-.033.332a1 1 0 0 1-.995.901h-1.79a1 1 0 0 1-.995-.9l-.033-.333a7.95 7.95 0 0 1-2.209-.915l-.259.212a1 1 0 0 1-1.34-.067L5.003 17.73a1 1 0 0 1-.067-1.34l.212-.26a7.953 7.953 0 0 1-.915-2.208L3.9 13.89a1 1 0 0 1-.9-.995v-1.79a1 1 0 0 1 .9-.995l.333-.033a7.953 7.953 0 0 1 .915-2.209l-.212-.259a1 1 0 0 1 .067-1.34L6.27 5.003a1 1 0 0 1 1.34-.067l.26.212a7.947 7.947 0 0 1 2.208-.915z"/><circle cx="2.5" cy="2.5" r="2.5" transform="matrix(1 0 0 -1 9.5 14.5)"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CogOne(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="2.5" cy="2.5" r="2.5" stroke-miterlimit="1" transform="matrix(1 0 0 -1 9.5 14.5)"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CogThree(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8.4 18.235a7.2 7.2 0 1 0 7.2-12.47m-7.2 12.47A7.2 7.2 0 0 1 5.765 8.4M8.4 18.235l-.9 1.56m8.1-14.03A7.2 7.2 0 0 0 5.765 8.4M15.6 5.765l.9-1.56M5.765 8.4l-1.56-.9m10.295 6l5.294 3M12 21v-1.8M12 9V3m4.5 16.794l-.899-1.558m-8.1-14.03l.898 1.558M20.999 12h-1.798m-16.2 0h1.798m14.995-4.5l-1.558.899M9.5 13.5l-5.294 3"/><circle cx="2.5" cy="2.5" r="2.5" stroke-miterlimit="1" transform="matrix(1 0 0 -1 9.5 14.5)"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CogTwo(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="2.5" cy="2.5" r="2.5" stroke-miterlimit="1" transform="matrix(1 0 0 -1 9.5 14.5)"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Columns(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17.5 3a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1h-3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1zm-8 0a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1h-3a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Command(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.012 5.977v12.046c0 2.645 3.316 3.953 5.14 2.13c1.825-1.825.516-5.141-2.13-5.141H5.978c-2.645 0-3.953 3.316-2.13 5.14c1.825 1.824 5.142.516 5.142-2.13V5.978c0-2.645-3.317-3.953-5.141-2.13c-1.824 1.825-.516 5.142 2.13 5.142h12.045c2.645 0 3.954-3.317 2.13-5.141c-1.825-1.824-5.141-.516-5.141 2.13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compass(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="m9 15l1.5-4.5L15 9l-1.5 4.5z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Components(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.85 6.15L12 3l3.15 3.15L12 9.3zm5.85 6.3l3.15-3.15L21 12.45l-3.15 3.15zm-5.85 5.4L12 14.7l3.15 3.15L12 21zM3 12l3.15-3.15L9.3 12l-3.15 3.15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Confetti(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.675 20.971a.508.508 0 0 1-.65-.637L5.615 9.21c.12-.374.6-.475.862-.183l7.756 7.544a.509.509 0 0 1-.212.82zm7.219-11.695L13.3 6.66c1.283-1.395 1.444-2.615.481-3.661M8 5.25v-.5m12-.25V4m0 9.5V13m-2 5.5V18m-4.219-5.586l2.406-2.615c1.605-1.744 3.209-1.744 4.813 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Config(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 5h-3m-4.25-2v4M13 5H3m4 7H3m7.75-2v4M21 12H11m10 7h-3m-4.25-2v4M13 19H3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ConfigVertical(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 3v3m-2 4.25h4M5 11v10m7-4v4m-2-7.75h4M12 3v10m7-10v3m-2 4.25h4M19 11v10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Contactless(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17.46 21c1.32-2.76 2.04-5.76 2.04-9s-.72-6.36-2.04-9m-4.32 15.96C14.1 16.8 14.7 14.4 14.7 12c0-2.4-.6-4.92-1.56-7.08m-4.32 12C9.54 15.36 9.9 13.68 9.9 12c0-1.68-.36-3.48-1.08-4.92M4.5 14.76c.36-.84.6-1.8.6-2.76c0-.96-.24-2.04-.6-2.88"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ContactlessCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8 9.99a9.03 9.03 0 0 1 0 4.02m2.975-5.59a13.01 13.01 0 0 1 .5 3.58a13.01 13.01 0 0 1-.5 3.58m3.25-8.72a17.01 17.01 0 0 1 .79 5.14a17.01 17.01 0 0 1-.79 5.14"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Controller(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.5 11.5v3M6 13h3m3-4.653c2.005 0 3.7-1.888 5.786-1.212c2.264.733 3.82 3.413 3.708 9.492c-.022 1.224-.336 2.578-1.546 3.106c-2.797 1.221-4.397-2.328-7-2.328h-1.897c-2.605 0-4.213 3.545-6.998 2.328c-1.21-.528-1.525-1.882-1.547-3.107c-.113-6.078 1.444-8.758 3.708-9.491C8.299 6.459 9.994 8.347 12 8.347m0-4.565v4.342M14.874 13h3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cookie(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.003 21a9.003 9.003 0 0 0 8.996-8.658c.006-.153-.16-.25-.298-.181c-2.476 1.247-4.006-.077-3.757-1.854a.229.229 0 0 0-.252-.257c-2.171.303-3.086-1.014-2.744-2.804a.225.225 0 0 0-.201-.261c-2.043-.182-2.212-2.54-1.861-3.69c.043-.142-.059-.3-.207-.295a9.003 9.003 0 0 0 .324 18M15 16.354l.354-.354M10 17.354l.354-.354M8 8.354L8.354 8M7 13.354L7.354 13M12 12.354l.354-.354"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copy(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20.829 12.861c.171-.413.171-.938.171-1.986s0-1.573-.171-1.986a2.25 2.25 0 0 0-1.218-1.218c-.413-.171-.938-.171-1.986-.171H11.1c-1.26 0-1.89 0-2.371.245a2.25 2.25 0 0 0-.984.984C7.5 9.209 7.5 9.839 7.5 11.1v6.525c0 1.048 0 1.573.171 1.986c.229.551.667.99 1.218 1.218c.413.171.938.171 1.986.171s1.573 0 1.986-.171m7.968-7.968a2.25 2.25 0 0 1-1.218 1.218c-.413.171-.938.171-1.986.171s-1.573 0-1.986.171a2.25 2.25 0 0 0-1.218 1.218c-.171.413-.171.938-.171 1.986s0 1.573-.171 1.986a2.25 2.25 0 0 1-1.218 1.218m7.968-7.968a11.678 11.678 0 0 1-7.75 7.9l-.218.068M16.5 7.5v-.9c0-1.26 0-1.89-.245-2.371a2.25 2.25 0 0 0-.983-.984C14.79 3 14.16 3 12.9 3H6.6c-1.26 0-1.89 0-2.371.245a2.25 2.25 0 0 0-.984.984C3 4.709 3 5.339 3 6.6v6.3c0 1.26 0 1.89.245 2.371c.216.424.56.768.984.984c.48.245 1.111.245 2.372.245H7.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copyleft(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M9.75 9c.48-.6 1.07-1 2-1c4.172 0 4.172 8 0 8c-.93 0-1.52-.4-2-1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copyright(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M14 9c-.48-.6-1.07-1-2-1c-4.171 0-4.171 8 0 8c.93 0 1.52-.4 2-1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CopyrightSlash(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m3 21l2.636-2.636M21 3l-2.636 2.636m0 0A9 9 0 1 1 5.636 18.364M18.364 5.636L5.636 18.364m-2.02-3.087A9 9 0 0 1 15.277 3.615"/><path d="M14 9c-.48-.6-1.07-1-2-1c-4.171 0-4.171 8 0 8c.93 0 1.52-.4 2-1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerDownLeft(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 13.973h10a4 4 0 0 0 4-4V5M5 13.973l4.78-5.027M5 13.973L9.78 19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerDownRight(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 13.973H9a4 4 0 0 1-4-4V5m14 8.973l-4.78-5.027M19 13.973L14.22 19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerLeftDown(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.027 19V9a4 4 0 0 1 4-4H19m-8.973 14l5.028-4.78M10.027 19L5 14.22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerLeftUp(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.027 5v10a4 4 0 0 0 4 4H19M10.027 5l5.028 4.78M10.027 5L5 9.78"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerRightDown(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13.973 19V9a4 4 0 0 0-4-4H5m8.973 14l-5.027-4.78M13.973 19L19 14.22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerRightUp(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13.973 5v10a4 4 0 0 1-4 4H5m8.973-14L8.946 9.78M13.973 5L19 9.78"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerUpLeft(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 10.027h10a4 4 0 0 1 4 4V19M5 10.027l4.78 5.028M5 10.027L9.78 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerUpRight(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 10.027H9a4 4 0 0 0-4 4V19m14-8.973l-4.78 5.028M19 10.027L14.22 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCard(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 9v-.8c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C4.52 5 5.08 5 6.2 5h11.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105V9M3 9h18M3 9v6.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h11.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.428.218-.986.218-2.104V9M6 15h4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCardCheck(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 9v-.8c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C4.52 5 5.08 5 6.2 5h6.3M3 9h9.5M3 9v6.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h11.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.428.218-.986.218-2.104V11m-5-3.678l1.484 1.576a.337.337 0 0 0 .502-.022L21 5M6 15h4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCardMinus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 9v6.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h11.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.428.218-.986.218-2.104V11M3 9v-.8c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C4.52 5 5.08 5 6.2 5h6.3M3 9h9.5M6 15h4m5.5-7.5h5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCardPlus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 9v6.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h11.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.428.218-.986.218-2.104V11M3 9v-.8c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C4.52 5 5.08 5 6.2 5h6.3M3 9h9.5M6 15h4m5.5-7.5h5M18 5v5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCardX(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 9v-.8c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C4.52 5 5.08 5 6.2 5h6.3M3 9h9.5M3 9v6.8c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.427.218.987.218 2.105.218h11.606c1.118 0 1.677 0 2.104-.218c.377-.192.683-.498.875-.874c.218-.428.218-.986.218-2.104V11M6 15h4m6.5-10l4 4m0-4l-4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Croissant(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.803 9.997L9.75 7.69C9.48 7.032 9.12 6.6 8.4 6.6m0 0H5.25C3.711 6.6 3 7.05 3 8.85c.075 1.614.827 3.266 1.915 4.462M8.4 6.6c0-1.395.216-3.6-1.8-3.6c-1.8 0-2.25 2.029-2.25 3.676M14 13.21l2.311 1.04c.657.27 1.089.63 1.089 1.35m0 0v3.15c0 1.539-.45 2.25-2.25 2.25c-1.614-.076-3.264-.824-4.459-1.912M17.4 15.6c1.395 0 3.6-.216 3.6 1.8c0 1.8-2.029 2.25-3.676 2.25M5.34 13l5.211-2.89c1.701-.945 4.311 1.602 3.339 3.34l-2.898 5.228C9.12 22.044 1.911 14.908 5.34 13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crop(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6.6 3v9.6c0 1.68 0 2.52.327 3.162a3 3 0 0 0 1.311 1.311c.642.327 1.482.327 3.162.327H21"/><path d="M9.2 6.6h3.4c1.68 0 2.52 0 3.162.327a3 3 0 0 1 1.311 1.311c.327.642.327 1.482.327 3.162v3.4M6.6 6.6H3M17.4 21v-3.6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crosshair(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 12h-3M6 12H3m9-6V3m0 18v-3"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cupcake(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m5 14l.804 5.626C5.948 20.636 6.308 21 7.385 21H10m-5-7h4m-4 0c-1.303-.604-2-2.236-2-3.666c0-1.536 1.03-2.85 2.49-3.397A.787.787 0 0 0 6 6.208c0-1.265 1.12-2.291 2.5-2.291c.668 0 1.31.322 1.941-.066A5.833 5.833 0 0 1 13.5 3C16.538 3 19 5.257 19 8.042c0 1.256 2 1.594 2 3.208c0 1.277-.712 2.44-2 2.75m0 0h-4m4 0l-.804 5.628C18.044 20.693 17.635 21 16.615 21H14m1-7H9m6 0l-1 7m-5-7l1 7m0 0h4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Danger(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 19v-.5M12 5v10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DangerCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 7.627v5.5m0 3.246v-.5"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DangerHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 7.627v5.5m0 3.246v-.5m8.5-.073V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DangerOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132zM12 7.627v5.5m0 3.246v-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DangerSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 7.627v5.5m0 3.246v-.5M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DangerTriangle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 8.5V14m0 3.247v-.5m-6.02-5.985C8.608 5.587 9.92 3 12 3c2.08 0 3.393 2.587 6.02 7.762l.327.644c2.182 4.3 3.274 6.45 2.287 8.022C19.648 21 17.208 21 12.327 21h-.654c-4.88 0-7.321 0-8.307-1.572c-.987-1.572.105-3.722 2.287-8.022z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DangerWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 7.627v5.5m0 3.246v-.5M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Database(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 3C7.582 3 4 4.29 4 5.88c0 4.16 16 4.16 16 0C20 4.29 16.418 3 12 3m8 8.75c0 4.667-16 4.667-16 0"/><path d="M4 6v12.165c0 3.78 16 3.78 16 0V6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DazeCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8 12l2-1.5L8 9m8 3l-2-1.5L16 9m0 7.25l-1.333-1l-1.334 1l-1.333-1l-1.333 1l-1.334-1l-1.333 1"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DazeGhost(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8 12l2-1.5L8 9m8 3l-2-1.5L16 9m0 7.25l-1.333-1l-1.334 1l-1.333-1l-1.333 1l-1.334-1l-1.333 1"/><path d="M3 18.562v-6.518C3 7.05 7.03 3 12 3s9 4.05 9 9.044v6.517c0 1.162-.967 2.519-2 2c-.835-.42-2.223-.52-3 0c-.874.585-2.126.585-3 0c-.885-.593-1.649-.57-2.5 0c-.874.585-2.126.585-3 0c-.777-.52-1.665-.42-2.5 0c-1.033.519-2-.838-2-2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DazeSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8 12l2-1.5L8 9m8 3l-2-1.5L16 9m0 7.25l-1.333-1l-1.334 1l-1.333-1l-1.333 1l-1.334-1l-1.333 1"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Delete(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m13.5 10l4 4m0-4l-4 4m6.095 4.5H9.298a2 2 0 0 1-1.396-.568l-5.35-5.216a1 1 0 0 1 0-1.432l5.35-5.216A2 2 0 0 1 9.298 5.5h10.297c.95 0 2.223.541 2.223 1.625v9.75c0 1.084-1.273 1.625-2.223 1.625"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Desktop(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 20h3m3 0h-3m0 0v-3m0 0h7a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Diamond(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.707 10.295a2.41 2.41 0 0 0 0 3.41l7.588 7.588a2.409 2.409 0 0 0 3.41 0l7.588-7.588a2.409 2.409 0 0 0 0-3.41l-7.588-7.588a2.41 2.41 0 0 0-3.41 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceFive(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12m5.25-4h-.5m.5 8h-.5m8.5-8h-.5m.5 8h-.5m-3.5-4h-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceFour(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12m5.25-4h-.5m.5 8h-.5m8.5-8h-.5m.5 8h-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceOne(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12m8.75 0h.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceSix(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12m5.25 0h-.5m.5-4h-.5m.5 8h-.5m8.5-4h-.5m.5-4h-.5m.5 8h-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceThree(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12m5.25-4h-.5m4.5 4h-.5m4.5 4h-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceTwo(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12m5.25-4h-.5m8.5 8h-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dislike(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.46 13.895H4.927C2.381 13.895 5.691 3 7.515 3h12.521c.532 0 .964.424.964.947v9.385a.945.945 0 0 1-.502.832c-2.062 1.106-4.481 2.012-5.678 4.129l-1.28 2.266a.874.874 0 0 1-.762.441c-3.18 0-2.237-4.63-1.805-6.47a.52.52 0 0 0-.513-.635"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Divide(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 12h18m-9-6.256V5m0 14v-.744"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dollar(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 5h-5m0 0H9.5a3.5 3.5 0 1 0 0 7H12m0-7V3m0 2v7m0 0h2.5a3.5 3.5 0 1 1 0 7H12m0-7v7m0 0H6m6 0v2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DollarCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M15.333 7.722H12m0 0h-1.667C9.045 7.722 8 8.68 8 9.862C8 11.041 9.045 12 10.333 12H12m0-4.278V6.5m0 1.222V12m0 0h1.667c1.288 0 2.333.958 2.333 2.139c0 1.181-1.045 2.139-2.333 2.139H12M12 12v4.278m0 0H8m4 0V17.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DollarHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/><path d="M15.333 7.722H12m0 0h-1.667C9.045 7.722 8 8.68 8 9.862C8 11.041 9.045 12 10.333 12H12m0-4.278V6.5m0 1.222V12m0 0h1.667c1.288 0 2.333.958 2.333 2.139c0 1.181-1.045 2.139-2.333 2.139H12M12 12v4.278m0 0H8m4 0V17.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DollarOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/><path d="M15.333 7.722H12m0 0h-1.667C9.045 7.722 8 8.68 8 9.862C8 11.041 9.045 12 10.333 12H12m0-4.278V6.5m0 1.222V12m0 0h1.667c1.288 0 2.333.958 2.333 2.139c0 1.181-1.045 2.139-2.333 2.139H12M12 12v4.278m0 0H8m4 0V17.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DollarSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/><path d="M15.333 7.722H12m0 0h-1.667C9.045 7.722 8 8.68 8 9.862C8 11.041 9.045 12 10.333 12H12m0-4.278V6.5m0 1.222V12m0 0h1.667c1.288 0 2.333.958 2.333 2.139c0 1.181-1.045 2.139-2.333 2.139H12M12 12v4.278m0 0H8m4 0V17.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DollarWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/><path d="M15.333 7.722H12m0 0h-1.667C9.045 7.722 8 8.68 8 9.862C8 11.041 9.045 12 10.333 12H12m0-4.278V6.5m0 1.222V12m0 0h1.667c1.288 0 2.333.958 2.333 2.139c0 1.181-1.045 2.139-2.333 2.139H12M12 12v4.278m0 0H8m4 0V17.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dots(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 12.25v-.5m4 .5v-.5m-8 .5v-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotsCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 12.25v-.5m4 .5v-.5m-8 .5v-.5"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotsHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 12.25v-.5m4 .5v-.5m-8 .5v-.5m12.5 4.05V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotsOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132zM12 12.25v-.5m4 .5v-.5m-8 .5v-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotsSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 12.25v-.5m4 .5v-.5m-8 .5v-.5M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotsVertical(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.25 12h-.5m.5-4h-.5m.5 8h-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotsVerticalCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12.25 12h-.5m.5-4h-.5m.5 8h-.5"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotsVerticalHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.25 12h-.5m.5-4h-.5m.5 8h-.5m8.75-.2V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotsVerticalOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132zM12.25 12h-.5m.5-4h-.5m.5 8h-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotsVerticalSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.25 12h-.5m.5-4h-.5m.5 8h-.5M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotsVerticalWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.25 12h-.5m.5-4h-.5m.5 8h-.5M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotsWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 12.25v-.5m4 .5v-.5m-8 .5v-.5m1.713-8.11c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Download(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 16.004V17a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-1M12 4.5v11m3.5-3.5L12 15.5L8.5 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Drop(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.495 3c3.58 3.56 9.345 7.602 6.932 13.397C18.275 19.163 15.492 21 12.5 21c-2.992 0-5.775-1.837-6.927-4.603C3.161 10.607 8.919 6.561 12.495 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Earth(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13 3.048a5 5 0 0 0 .982 8.3c2.018 1.013 2.789-.352 3.881.384c.71.478.897 1.44.42 2.149c-.501.742-1.283 1.119-1.148 2.336c.077.687.498 1.278 1.045 1.783M4 9.28a4.979 4.979 0 0 1 2.806 1.846a4.981 4.981 0 0 1 .992 3.424c-.052.626.356 1.258.881 1.603A2.71 2.71 0 0 1 9 20.44"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Edit(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 21h16M5.666 13.187A2.278 2.278 0 0 0 5 14.797V18h3.223c.604 0 1.183-.24 1.61-.668l9.5-9.505a2.278 2.278 0 0 0 0-3.22l-.938-.94a2.277 2.277 0 0 0-3.222.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditOne(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.533 11.15A1.823 1.823 0 0 0 9 12.438V15h2.578c.483 0 .947-.192 1.289-.534l7.6-7.604a1.822 1.822 0 0 0 0-2.577l-.751-.751a1.822 1.822 0 0 0-2.578 0z"/><path d="M21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Egg(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 21c4.004 0 7.25-3.224 7.25-7.2S16.004 3 12 3S4.75 9.824 4.75 13.8c0 3.976 3.246 7.2 7.25 7.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eight(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.875 12S15 12.625 15 14.5S13.6 17 11.875 17c-1.726 0-3.125-.625-3.125-2.5s3.125-2.5 3.125-2.5m0 0S15 11.375 15 9.5S13.6 7 11.875 7c-1.726 0-3.125.625-3.125 2.5s3.125 2.5 3.125 2.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EightCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M12 12s2.5.5 2.5 2s-1.12 2-2.5 2s-2.5-.5-2.5-2s2.5-2 2.5-2m0 0s2.5-.5 2.5-2s-1.12-2-2.5-2s-2.5.5-2.5 2s2.5 2 2.5 2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EightDiamond(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 12s2.5.5 2.5 2s-1.12 2-2.5 2s-2.5-.5-2.5-2s2.5-2 2.5-2m0 0s2.5-.5 2.5-2s-1.12-2-2.5-2s-2.5.5-2.5 2s2.5 2 2.5 2"/><path d="M2.707 10.295a2.41 2.41 0 0 0 0 3.41l7.588 7.588a2.409 2.409 0 0 0 3.41 0l7.588-7.588a2.409 2.409 0 0 0 0-3.41l-7.588-7.588a2.41 2.41 0 0 0-3.41 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EightHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 12s2.5.5 2.5 2s-1.12 2-2.5 2s-2.5-.5-2.5-2s2.5-2 2.5-2m0 0s2.5-.5 2.5-2s-1.12-2-2.5-2s-2.5.5-2.5 2s2.5 2 2.5 2"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EightOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 12s2.5.5 2.5 2s-1.12 2-2.5 2s-2.5-.5-2.5-2s2.5-2 2.5-2m0 0s2.5-.5 2.5-2s-1.12-2-2.5-2s-2.5.5-2.5 2s2.5 2 2.5 2"/><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EightSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 12s2.5.5 2.5 2s-1.12 2-2.5 2s-2.5-.5-2.5-2s2.5-2 2.5-2m0 0s2.5-.5 2.5-2s-1.12-2-2.5-2s-2.5.5-2.5 2s2.5 2 2.5 2"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EightWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 12s2.5.5 2.5 2s-1.12 2-2.5 2s-2.5-.5-2.5-2s2.5-2 2.5-2m0 0s2.5-.5 2.5-2s-1.12-2-2.5-2s-2.5.5-2.5 2s2.5 2 2.5 2"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Elevator(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 8.294h4.5a2 2 0 0 1 2 2V19a2 2 0 0 1-2 2H12m0-12.706H7.5a2 2 0 0 0-2 2V19a2 2 0 0 0 2 2H12m0-12.706V21M7.125 4.588L8.75 3l1.625 1.588M13.625 3l1.625 1.588L16.875 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Envelope(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2 12c0-3.771 0-5.657 1.464-6.828C4.93 4 7.286 4 12 4c4.714 0 7.071 0 8.535 1.172C22 6.343 22 8.229 22 12c0 3.771 0 5.657-1.465 6.828C19.072 20 16.714 20 12 20s-7.071 0-8.536-1.172C2 17.657 2 15.771 2 12"/><path d="M20.667 5.31L15.84 9.8c-1.836 1.53-2.755 2.295-3.841 2.295c-1.086 0-2.004-.765-3.841-2.296L3.333 5.31"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Euro(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14 10H7.12m0 0H4m3.12 0c.55-4.254 3.01-6 7.38-6c2.418 0 4.251.535 5.5 1.733M7.12 10c-.08.614-.12 1.28-.12 2s.04 1.386.12 2M14 14H7.12m0 0H4m3.12 0c.55 4.254 3.01 6 7.38 6c2.418 0 4.251-.535 5.5-1.733"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EuroCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M12 12H8.5m0 0H7m1.5 0c0 2 1.565 4 3.75 4c1.209 0 2.126-.267 2.75-.867M8.5 12c0-2 1.25-4 3.75-4c1.209 0 2.126.267 2.75.866"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EuroHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/><path d="M12 12H8.5m0 0H7m1.5 0c0 2 1.565 4 3.75 4c1.209 0 2.126-.267 2.75-.867M8.5 12c0-2 1.25-4 3.75-4c1.209 0 2.126.267 2.75.866"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EuroOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/><path d="M12 12H8.5m0 0H7m1.5 0c0 2 1.565 4 3.75 4c1.209 0 2.126-.267 2.75-.867M8.5 12c0-2 1.25-4 3.75-4c1.209 0 2.126.267 2.75.866"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EuroSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/><path d="M12 12H8.5m0 0H7m1.5 0c0 2 1.565 4 3.75 4c1.209 0 2.126-.267 2.75-.867M8.5 12c0-2 1.25-4 3.75-4c1.209 0 2.126.267 2.75.866"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EuroWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/><path d="M12 12H8.5m0 0H7m1.5 0c0 2 1.565 4 3.75 4c1.209 0 2.126-.267 2.75-.867M8.5 12c0-2 1.25-4 3.75-4c1.209 0 2.126.267 2.75.866"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Exclude(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.4 15.6H6a3 3 0 0 1-3-3V6a3 3 0 0 1 3-3h6.6a3 3 0 0 1 3 3v2.4m-7.2 7.2V18a3 3 0 0 0 3 3H18a3 3 0 0 0 3-3v-6.6a3 3 0 0 0-3-3h-2.4m-7.2 7.2v-1.8m0 1.8h1.8m5.4-7.2h-1.8m1.8 0v1.8m-7.2 0a1.8 1.8 0 0 1 1.8-1.8m5.4 5.4a1.8 1.8 0 0 1-1.8 1.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExternalLink(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 3.007c-2.946.032-4.59.219-5.682 1.311C3 5.636 3 7.758 3 12c0 4.243 0 6.364 1.318 7.682C5.636 21 7.758 21 12 21c4.243 0 6.364 0 7.682-1.318c1.061-1.061 1.268-2.643 1.308-5.434M21 3h-6.75M21 3v6.75M21 3l-8.25 8.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eye(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.55 13.406c-.272-.373-.408-.56-.502-.92a2.46 2.46 0 0 1 0-.971c.094-.361.23-.548.502-.92C4.039 8.55 7.303 5 12 5c4.697 0 7.961 3.55 9.45 5.594c.272.373.408.56.502.92a2.46 2.46 0 0 1 0 .971c-.094.361-.23.548-.502.92C19.961 15.45 16.697 19 12 19c-4.697 0-7.961-3.55-9.45-5.594"/><path d="M12 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeSlash(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17.357 6.643L21 3m-3.643 3.643c1.878 1.2 3.262 2.81 4.093 3.951c.272.373.408.56.502.92a2.46 2.46 0 0 1 0 .971c-.094.361-.23.548-.502.92C19.961 15.45 16.697 19 12 19c-2.075 0-3.87-.693-5.357-1.643M17.357 6.643l-3.943 3.943M3 21l3.643-3.643m0 0l3.943-3.943M3.86 15a17.603 17.603 0 0 1-1.311-1.594c-.272-.373-.408-.56-.502-.92a2.46 2.46 0 0 1 0-.971c.094-.361.23-.548.502-.92C4.039 8.55 7.303 5 12 5c.59 0 1.158.056 1.703.158m-.289 5.428a2 2 0 0 1-2.828 2.828m2.828-2.828l-2.828 2.828"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceId(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11 13.75h1v-4m4-.25V8m-7 8.5c1.5 1.5 4.5 1.5 6 0m-7-7V8m-.5 12.886c-1.463-.144-2.447-.47-3.182-1.204c-.735-.735-1.06-1.72-1.204-3.182M7.5 3.114c-1.463.144-2.447.47-3.182 1.204c-.735.735-1.06 1.72-1.204 3.182M16.5 3.114c1.463.144 2.447.47 3.182 1.204c.735.735 1.06 1.72 1.204 3.182M16.5 20.886c1.463-.144 2.447-.47 3.182-1.204c.735-.735 1.06-1.72 1.204-3.182"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FatArrowDown(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.483 13.105c-.43 0-.645.545-.34.863l7.516 6.884a.466.466 0 0 0 .682 0l7.517-6.884c.304-.318.088-.863-.341-.863H15.68a.495.495 0 0 1-.483-.506V3.506A.494.494 0 0 0 14.716 3H9.284a.494.494 0 0 0-.482.506v9.093c0 .28-.216.506-.483.506z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FatArrowDownLeft(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.932 18.963c.315.315-.062 1.026-.519 1.037H5.66c-.58 0-.871 0-1.093-.113a1.037 1.037 0 0 1-.454-.454C4 19.213 4 18.922 4 18.34V8.587c.01-.457.722-.834 1.038-.519l2.714 2.78c.195.195.52.187.725-.018l6.671-6.671a.513.513 0 0 1 .725-.017l3.985 3.985a.513.513 0 0 1-.017.725l-6.671 6.671a.513.513 0 0 0-.017.725z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FatArrowDownRight(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.068 18.963c-.315.315.062 1.026.519 1.037h9.753c.58 0 .872 0 1.093-.113c.196-.1.354-.258.454-.454c.113-.221.113-.512.113-1.093V8.587c-.01-.457-.722-.834-1.038-.519l-2.714 2.78a.513.513 0 0 1-.725-.018L8.852 4.16a.513.513 0 0 0-.725-.017L4.142 8.127a.513.513 0 0 0 .017.725l6.671 6.671c.205.205.213.53.017.725z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FatArrowLeft(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.895 4.483c0-.43-.545-.645-.863-.34l-6.884 7.516a.467.467 0 0 0 0 .682l6.884 7.517c.318.304.863.088.863-.341V15.68c0-.267.227-.483.506-.483h9.093c.28 0 .506-.216.506-.482V9.284a.494.494 0 0 0-.506-.482h-9.093a.495.495 0 0 1-.506-.483z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FatArrowRight(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13.105 4.483c0-.43.545-.645.863-.34l6.884 7.516a.466.466 0 0 1 0 .682l-6.884 7.517c-.318.304-.863.088-.863-.341V15.68a.495.495 0 0 0-.506-.483H3.506A.494.494 0 0 1 3 14.716V9.284c0-.266.226-.482.506-.482h9.093c.28 0 .506-.216.506-.483z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FatArrowUp(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.483 10.895c-.43 0-.645-.545-.34-.863l7.516-6.884a.467.467 0 0 1 .682 0l7.517 6.884c.304.318.088.863-.341.863H15.68a.495.495 0 0 0-.483.506v9.093c0 .28-.216.506-.482.506H9.284a.494.494 0 0 1-.482-.506v-9.093a.495.495 0 0 0-.483-.506z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FatArrowUpLeft(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.932 5.038c.315-.316-.062-1.027-.519-1.038H5.66c-.58 0-.871 0-1.093.113a1.04 1.04 0 0 0-.454.453C4 4.789 4 5.08 4 5.66v9.753c.01.457.722.834 1.038.519l2.714-2.78a.513.513 0 0 1 .725.018l6.671 6.671c.205.205.53.213.725.017l3.985-3.985a.513.513 0 0 0-.017-.725L13.17 8.477a.513.513 0 0 1-.017-.725z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FatArrowUpRight(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.068 5.038C7.753 4.722 8.13 4.01 8.587 4h9.753c.58 0 .872 0 1.093.113c.196.1.354.258.454.453C20 4.789 20 5.08 20 5.66v9.753c-.01.457-.722.834-1.038.519l-2.714-2.78a.513.513 0 0 0-.725.018l-6.671 6.67a.513.513 0 0 1-.725.017l-3.985-3.985a.513.513 0 0 1 .017-.725l6.671-6.671a.513.513 0 0 0 .017-.725z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FatCornerDownLeft(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 4.077c-2.202 2.81-4.157 4.406-5.866 4.785c-1.709.38-3.336.436-4.88.172V4L3 12.214L10.253 20v-4.784C13.11 15.192 15.54 14.12 17.54 12S20.693 7.239 21 4.077"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FatCornerDownRight(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 4.077c2.202 2.81 4.157 4.406 5.866 4.785c1.709.38 3.336.436 4.88.172V4L21 12.214L13.747 20v-4.784C10.89 15.192 8.46 14.12 6.46 12S3.307 7.239 3 4.077"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FatCornerLeftDown(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.923 3c-2.81 2.202-4.406 4.157-4.785 5.866c-.38 1.709-.436 3.336-.172 4.88H20L11.786 21L4 13.747h4.784C8.808 10.89 9.88 8.46 12 6.46S16.761 3.307 19.923 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FatCornerLeftUp(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.923 21c-2.81-2.202-4.406-4.157-4.785-5.866c-.38-1.709-.436-3.336-.172-4.88H20L11.786 3L4 10.253h4.784C8.808 13.11 9.88 15.54 12 17.54s4.761 3.154 7.923 3.461"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FatCornerRightDown(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.077 3c2.81 2.202 4.406 4.157 4.785 5.866c.38 1.709.436 3.336.172 4.88H4L12.214 21L20 13.747h-4.784C15.192 10.89 14.12 8.46 12 6.46S7.239 3.307 4.077 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FatCornerRightUp(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.077 21c2.81-2.202 4.406-4.157 4.785-5.866c.38-1.709.436-3.336.172-4.88H4L12.214 3L20 10.253h-4.784C15.192 13.11 14.12 15.54 12 17.54S7.239 20.693 4.077 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FatCornerUpLeft(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 19.923c-2.202-2.81-4.157-4.406-5.866-4.785c-1.709-.38-3.336-.436-4.88-.172V20L3 11.786L10.253 4v4.784C13.11 8.808 15.54 9.88 17.54 12s3.154 4.761 3.461 7.923"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FatCornerUpRight(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 19.923c2.202-2.81 4.157-4.406 5.866-4.785c1.709-.38 3.336-.436 4.88-.172V20L21 11.786L13.747 4v4.784C10.89 8.808 8.46 9.88 6.46 12S3.307 16.761 3 19.923"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Female(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 15a6 6 0 1 0 0-12a6 6 0 0 0 0 12m0 0v6m-2-2h4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func File(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.728 3H7.5a2.25 2.25 0 0 0-2.25 2.25v13.5A2.25 2.25 0 0 0 7.5 21h9a2.25 2.25 0 0 0 2.25-2.25V12M9.728 3C10.971 3 12 4.007 12 5.25V7.5a2.25 2.25 0 0 0 2.25 2.25h2.25A2.25 2.25 0 0 1 18.75 12M9.728 3c3.69 0 9.022 5.36 9.022 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileCheck(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.728 3H7.5a2.25 2.25 0 0 0-2.25 2.25v13.5A2.25 2.25 0 0 0 7.5 21h9a2.25 2.25 0 0 0 2.25-2.25V12M9.728 3C10.971 3 12 4.007 12 5.25V7.5a2.25 2.25 0 0 0 2.25 2.25h2.25A2.25 2.25 0 0 1 18.75 12M9.728 3c3.69 0 9.022 5.36 9.022 9"/><path d="m10 16.242l1.039 1.181c.095.109.267.1.351-.016L13.5 14.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileMinus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 16h4M9.728 3H7.5a2.25 2.25 0 0 0-2.25 2.25v13.5A2.25 2.25 0 0 0 7.5 21h9a2.25 2.25 0 0 0 2.25-2.25V12M9.728 3C10.971 3 12 4.007 12 5.25V7.5a2.25 2.25 0 0 0 2.25 2.25h2.25A2.25 2.25 0 0 1 18.75 12M9.728 3c3.69 0 9.022 5.36 9.022 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilePlus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 15.5h4m-2-2v4M9.728 3H7.5a2.25 2.25 0 0 0-2.25 2.25v13.5A2.25 2.25 0 0 0 7.5 21h9a2.25 2.25 0 0 0 2.25-2.25V12M9.728 3C10.971 3 12 4.007 12 5.25V7.5a2.25 2.25 0 0 0 2.25 2.25h2.25A2.25 2.25 0 0 1 18.75 12M9.728 3c3.69 0 9.022 5.36 9.022 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileText(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.478 3H7.25A2.25 2.25 0 0 0 5 5.25v13.5A2.25 2.25 0 0 0 7.25 21h9a2.25 2.25 0 0 0 2.25-2.25V12M9.478 3c1.243 0 2.272 1.007 2.272 2.25V7.5A2.25 2.25 0 0 0 14 9.75h2.25A2.25 2.25 0 0 1 18.5 12M9.478 3c3.69 0 9.022 5.36 9.022 9M9 16.5h6m-6-3h4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileX(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m10.5 14.5l2.828 2.828m0-2.828L10.5 17.328M9.728 3H7.5a2.25 2.25 0 0 0-2.25 2.25v13.5A2.25 2.25 0 0 0 7.5 21h9a2.25 2.25 0 0 0 2.25-2.25V12M9.728 3C10.971 3 12 4.007 12 5.25V7.5a2.25 2.25 0 0 0 2.25 2.25h2.25A2.25 2.25 0 0 1 18.75 12M9.728 3c3.69 0 9.022 5.36 9.022 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Film(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.4 3.468v17.064m9-17.064v17.064M7.401 7.473H3.486M7.401 12H3.027m4.374 4.473H3.432m17.469-9h-3.915M20.901 12h-4.374m-.054 0h-9.9m14.328 4.473h-3.969M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Filter(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.5 7h15M7 12h10m-7 5h4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterOne(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.046 17.676v-3.918c0-.554 0-.832-.05-1.1a2.93 2.93 0 0 0-.219-.686c-.114-.247-.274-.474-.595-.926L5.935 6.467c-.566-.797-.849-1.196-.836-1.529a.977.977 0 0 1 .38-.735C5.743 4 6.232 4 7.21 4h9.581c.978 0 1.467 0 1.73.203a.97.97 0 0 1 .38.735c.014.333-.27.732-.835 1.53l-3.247 4.578c-.32.452-.481.679-.595.926a2.93 2.93 0 0 0-.22.687c-.05.267-.05.544-.05 1.1v5.871m-3.907-1.954c1.654-.732 3.908-.296 3.908 1.954m-3.908-1.954c-.033 2.235 2.262 2.792 3.908 1.954"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FineTune(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 8.5h11m-18 0a2 2 0 1 0 4 0a2 2 0 0 0-4 0m0 7h11m3 0a2 2 0 1 0 4 0a2 2 0 0 0-4 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fire(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19.27 16.202A7.812 7.812 0 0 1 12.06 21c-4.313 0-7.81-3.492-7.81-7.8S5.89 7.13 8.455 3c4.806 2.1 4.806 8.4 4.806 8.4s1.579-3.038 4.807-4.5c1.034 3.042 2.43 6.365 1.202 9.302"/><path d="M12 18a5 5 0 0 1-5-5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Five(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.375 7c-2.5.625-5.625 0-5.625 0v3.566h3.473c1.534 0 2.777 1.064 2.777 2.377v1.399c0 3.522-6.25 3.566-6.25 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FiveCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M14 8c-2 .5-4.5 0-4.5 0v2.852h2.778c1.227 0 2.222.852 2.222 1.902v1.12c0 2.818-5 2.852-5 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FiveDiamond(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14 8c-2 .5-4.5 0-4.5 0v2.852h2.778c1.227 0 2.222.852 2.222 1.902v1.12c0 2.818-5 2.852-5 0"/><path d="M2.707 10.295a2.41 2.41 0 0 0 0 3.41l7.588 7.588a2.409 2.409 0 0 0 3.41 0l7.588-7.588a2.409 2.409 0 0 0 0-3.41l-7.588-7.588a2.41 2.41 0 0 0-3.41 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FiveHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14 8c-2 .5-4.5 0-4.5 0v2.852h2.778c1.227 0 2.222.852 2.222 1.902v1.12c0 2.818-5 2.852-5 0"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FiveOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14 8c-2 .5-4.5 0-4.5 0v2.852h2.778c1.227 0 2.222.852 2.222 1.902v1.12c0 2.818-5 2.852-5 0"/><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FiveSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14 8c-2 .5-4.5 0-4.5 0v2.852h2.778c1.227 0 2.222.852 2.222 1.902v1.12c0 2.818-5 2.852-5 0"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FiveWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14 8c-2 .5-4.5 0-4.5 0v2.852h2.778c1.227 0 2.222.852 2.222 1.902v1.12c0 2.818-5 2.852-5 0"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flag(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 21v-5.313m0 0c5.818-4.55 10.182 4.55 16 0V4.313c-5.818 4.55-10.182-4.55-16 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlagOne(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4.75 14l13.78-4.04c.96-.282.96-1.638 0-1.92L4.75 4m0 10V4m0 10v7m0-17V3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flask(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.894 3.5v8l-3.66 3.965c-.932 1.265-2.795 3.276-.948 4.622c.568.413 1.615.413 3.71.413h6.009c2.094 0 3.141 0 3.709-.413c1.847-1.346-.016-3.357-.949-4.622l-3.66-3.965v-8m-4.21 0h4.21m-4.21 0H8.84m5.265 0h1.053"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folder(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 6a2 2 0 0 1 2-2h1.745a2 2 0 0 1 1.322.5l2.272 2a2 2 0 0 0 1.322.5H19a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderCheck(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 6a2 2 0 0 1 2-2h1.745a2 2 0 0 1 1.322.5l2.272 2a2 2 0 0 0 1.322.5H19a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/><path d="m10.258 13.242l1.034 1.181c.095.109.266.1.35-.016l2.1-2.907"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderHeart(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 6a2 2 0 0 1 2-2h1.745a2 2 0 0 1 1.322.5l2.272 2a2 2 0 0 0 1.322.5H19a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/><path d="M8.62 12.024c0 1.972 3.38 4.226 3.38 4.226s3.38-2.254 3.38-4.226c0-1.88-2.55-2.454-3.38-.798c-.814-1.625-3.38-1.053-3.38.798"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderMinus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 6a2 2 0 0 1 2-2h1.745a2 2 0 0 1 1.322.5l2.272 2a2 2 0 0 0 1.322.5H19a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2zm7 7h4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOne(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.661 7H19a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V7m8.661 0a2 2 0 0 1-1.322-.5l-2.272-2A2 2 0 0 0 6.745 4H5a2 2 0 0 0-2 2v1m8.661 0h-8.66"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderPlus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 6a2 2 0 0 1 2-2h1.745a2 2 0 0 1 1.322.5l2.272 2a2 2 0 0 0 1.322.5H19a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2zm7 7.25h4m-2-2v4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderSlash(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3 21l1.177-1.177M21 3l-4 4m0 0h2a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2H5c-.293 0-.572-.063-.823-.177M17 7L4.177 19.823M3 16.5V6a2 2 0 0 1 2-2h1.745a2 2 0 0 1 1.322.5l2.272 2a2 2 0 0 0 1.322.5H13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderTwo(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 7h-7.339a2 2 0 0 1-1.322-.5l-2.272-2M19 7a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h1.745a2 2 0 0 1 1.322.5M19 7a2.5 2.5 0 0 0-2.5-2.5H8.066"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderX(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 6a2 2 0 0 1 2-2h1.745a2 2 0 0 1 1.322.5l2.272 2a2 2 0 0 0 1.322.5H19a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2zm7.5 6l2.828 2.828m0-2.828L10.5 14.828"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Forward(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.985 9.419C11.328 10.515 12 11.063 12 12s-.672 1.485-2.015 2.582c-.371.302-.74.587-1.077.824a18.14 18.14 0 0 1-.98.635c-1.341.816-2.011 1.223-2.612.772c-.602-.451-.656-1.396-.766-3.285A27.07 27.07 0 0 1 4.5 12c0-.47.02-.993.05-1.528c.11-1.89.164-2.834.766-3.285c.6-.451 1.27-.044 2.611.771c.348.212.684.427.98.636c.339.237.707.522 1.078.825m7.5 0C18.828 10.515 19.5 11.063 19.5 12s-.672 1.485-2.015 2.582c-.371.302-.74.587-1.077.824c-.297.209-.633.424-.98.635c-1.341.816-2.011 1.223-2.613.772c-.6-.451-.655-1.396-.764-3.285A27.07 27.07 0 0 1 12 12c0-.47.02-.993.05-1.528c.11-1.89.164-2.834.765-3.285c.602-.451 1.272-.044 2.612.771c.348.212.684.427.98.636c.339.237.707.522 1.078.825"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ForwardCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.242 10.71c.672.547 1.008.821 1.008 1.29c0 .469-.336.743-1.008 1.29c-.185.152-.37.295-.538.413a9.098 9.098 0 0 1-.49.318c-.67.407-1.006.611-1.306.385c-.3-.225-.328-.697-.383-1.642A13.532 13.532 0 0 1 7.5 12c0-.235.01-.497.025-.764c.055-.945.082-1.417.383-1.642c.3-.226.635-.022 1.306.385c.174.106.342.214.49.318c.169.118.353.261.538.412m5.75.001c.672.546 1.008.82 1.008 1.29c0 .469-.336.743-1.008 1.29c-.185.152-.37.295-.538.413a9.093 9.093 0 0 1-.49.318c-.67.407-1.006.611-1.306.385c-.3-.225-.328-.697-.383-1.642A13.577 13.577 0 0 1 13.25 12c0-.235.01-.497.025-.764c.055-.945.082-1.417.383-1.642c.3-.226.636-.022 1.306.385c.174.106.341.214.49.318c.169.118.353.261.538.412"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ForwardHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.242 10.71c.672.547 1.008.821 1.008 1.29c0 .469-.336.743-1.008 1.29c-.185.152-.37.295-.538.413a9.098 9.098 0 0 1-.49.318c-.67.407-1.006.611-1.306.385c-.3-.225-.328-.697-.383-1.642A13.532 13.532 0 0 1 7.5 12c0-.235.01-.497.025-.764c.055-.945.082-1.417.383-1.642c.3-.226.635-.022 1.306.385c.174.106.342.214.49.318c.169.118.353.261.538.412m5.75.001c.672.546 1.008.82 1.008 1.29c0 .469-.336.743-1.008 1.29c-.185.152-.37.295-.538.413a9.093 9.093 0 0 1-.49.318c-.67.407-1.006.611-1.306.385c-.3-.225-.328-.697-.383-1.642A13.577 13.577 0 0 1 13.25 12c0-.235.01-.497.025-.764c.055-.945.082-1.417.383-1.642c.3-.226.636-.022 1.306.385c.174.106.341.214.49.318c.169.118.353.261.538.412"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ForwardOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/><path d="M10.242 10.71c.672.547 1.008.821 1.008 1.29c0 .469-.336.743-1.008 1.29c-.185.152-.37.295-.538.413a9.098 9.098 0 0 1-.49.318c-.67.407-1.006.611-1.306.385c-.3-.225-.328-.697-.383-1.642A13.532 13.532 0 0 1 7.5 12c0-.235.01-.497.025-.764c.055-.945.082-1.417.383-1.642c.3-.226.635-.022 1.306.385c.174.106.342.214.49.318c.169.118.353.261.538.412m5.75.001c.672.546 1.008.82 1.008 1.29c0 .469-.336.743-1.008 1.29c-.185.152-.37.295-.538.413a9.093 9.093 0 0 1-.49.318c-.67.407-1.006.611-1.306.385c-.3-.225-.328-.697-.383-1.642A13.577 13.577 0 0 1 13.25 12c0-.235.01-.497.025-.764c.055-.945.082-1.417.383-1.642c.3-.226.636-.022 1.306.385c.174.106.341.214.49.318c.169.118.353.261.538.412"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ForwardSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.242 10.71c.672.547 1.008.821 1.008 1.29c0 .469-.336.743-1.008 1.29c-.185.152-.37.295-.538.413a9.098 9.098 0 0 1-.49.318c-.67.407-1.006.611-1.306.385c-.3-.225-.328-.697-.383-1.642A13.532 13.532 0 0 1 7.5 12c0-.235.01-.497.025-.764c.055-.945.082-1.417.383-1.642c.3-.226.635-.022 1.306.385c.174.106.342.214.49.318c.169.118.353.261.538.412m5.75.001c.672.546 1.008.82 1.008 1.29c0 .469-.336.743-1.008 1.29c-.185.152-.37.295-.538.413a9.093 9.093 0 0 1-.49.318c-.67.407-1.006.611-1.306.385c-.3-.225-.328-.697-.383-1.642A13.577 13.577 0 0 1 13.25 12c0-.235.01-.497.025-.764c.055-.945.082-1.417.383-1.642c.3-.226.636-.022 1.306.385c.174.106.341.214.49.318c.169.118.353.261.538.412"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ForwardWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.242 10.71c.672.547 1.008.821 1.008 1.29c0 .469-.336.743-1.008 1.29c-.185.152-.37.295-.538.413a9.098 9.098 0 0 1-.49.318c-.67.407-1.006.611-1.306.385c-.3-.225-.328-.697-.383-1.642A13.532 13.532 0 0 1 7.5 12c0-.235.01-.497.025-.764c.055-.945.082-1.417.383-1.642c.3-.226.635-.022 1.306.385c.174.106.342.214.49.318c.169.118.353.261.538.412m5.75.001c.672.546 1.008.82 1.008 1.29c0 .469-.336.743-1.008 1.29c-.185.152-.37.295-.538.413a9.093 9.093 0 0 1-.49.318c-.67.407-1.006.611-1.306.385c-.3-.225-.328-.697-.383-1.642A13.577 13.577 0 0 1 13.25 12c0-.235.01-.497.025-.764c.055-.945.082-1.417.383-1.642c.3-.226.636-.022 1.306.385c.174.106.341.214.49.318c.169.118.353.261.538.412"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Four(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.917 7c-1.042 3.75-4.167 6.875-4.167 6.875H15M13.438 17v-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FourCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M12.583 8c-.833 3-3.333 5.5-3.333 5.5h5M13 16v-4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FourDiamond(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12.583 8c-.833 3-3.333 5.5-3.333 5.5h5M13 16v-4"/><path d="M2.707 10.295a2.41 2.41 0 0 0 0 3.41l7.588 7.588a2.409 2.409 0 0 0 3.41 0l7.588-7.588a2.409 2.409 0 0 0 0-3.41l-7.588-7.588a2.41 2.41 0 0 0-3.41 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FourHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12.583 8c-.833 3-3.333 5.5-3.333 5.5h5M13 16v-4"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FourOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12.583 8c-.833 3-3.333 5.5-3.333 5.5h5M13 16v-4"/><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FourSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/><path d="M12.583 8c-.833 3-3.333 5.5-3.333 5.5h5M13 16v-4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FourWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/><path d="M12.583 8c-.833 3-3.333 5.5-3.333 5.5h5M13 16v-4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Frame(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 6.6H3m18 10.8H3M6.6 3v18M17.4 3v18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FunnyCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8.699 10.646l-.104-.49m4.995-.55l-.104-.49M8.5 15.57c1.258.316 2.686.316 4.123-.069c1.436-.385 2.672-1.099 3.604-2.001"/><path d="m12.587 15.637l.478.974a1.5 1.5 0 1 0 2.693-1.322l-.46-.935"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FunnyGhost(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8.699 10.646l-.104-.49m4.995-.55l-.104-.49M8.5 15.57c1.258.316 2.686.316 4.123-.069c1.436-.385 2.672-1.099 3.604-2.001"/><path d="m12.587 15.637l.478.974a1.5 1.5 0 1 0 2.693-1.322l-.46-.935"/><path d="M3 18.562v-6.518C3 7.05 7.03 3 12 3s9 4.05 9 9.044v6.517c0 1.162-.967 2.519-2 2c-.835-.42-2.223-.52-3 0c-.874.585-2.126.585-3 0c-.885-.593-1.649-.57-2.5 0c-.874.585-2.126.585-3 0c-.777-.52-1.665-.42-2.5 0c-1.033.519-2-.838-2-2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FunnySquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8.699 10.646l-.104-.49m4.995-.55l-.104-.49M8.5 15.57c1.258.316 2.686.316 4.123-.069c1.436-.385 2.672-1.099 3.604-2.001"/><path d="m12.587 15.637l.478.974a1.5 1.5 0 1 0 2.693-1.322l-.46-.935"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GhostDaze(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8 12l2-1.5L8 9m8 3l-2-1.5L16 9m0 7.25l-1.333-1l-1.334 1l-1.333-1l-1.333 1l-1.334-1l-1.333 1"/><path d="M3 18.561v-6.517C3 7.05 7.03 3 12 3s9 4.05 9 9.044v6.517c0 1.162-.967 2.519-2 2c-.835-.42-2.223-.52-3 0c-.874.585-2.126.585-3 0c-.885-.593-1.649-.57-2.5 0c-.874.585-2.126.585-3 0c-.777-.52-1.665-.42-2.5 0c-1.033.519-2-.838-2-2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GhostFunny(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8.699 10.646l-.104-.49m4.995-.55l-.104-.49M8.5 15.57c1.258.316 2.686.316 4.123-.069c1.436-.385 2.672-1.099 3.604-2.001"/><path d="m12.587 15.637l.478.974a1.5 1.5 0 1 0 2.693-1.322l-.46-.935"/><path d="M3 18.561v-6.517C3 7.05 7.03 3 12 3s9 4.05 9 9.044v6.517c0 1.162-.967 2.519-2 2c-.835-.42-2.223-.52-3 0c-.874.585-2.126.585-3 0c-.885-.593-1.649-.57-2.5 0c-.874.585-2.126.585-3 0c-.777-.52-1.665-.42-2.5 0c-1.033.519-2-.838-2-2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GhostIndifferent(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 15.5h6m-5.5-5V10m5 .5V10M3 18.561v-6.517C3 7.05 7.03 3 12 3s9 4.05 9 9.044v6.517c0 1.162-.967 2.519-2 2c-.835-.42-2.223-.52-3 0c-.874.585-2.126.585-3 0c-.885-.593-1.649-.57-2.5 0c-.874.585-2.126.585-3 0c-.777-.52-1.665-.42-2.5 0c-1.033.519-2-.838-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GhostSad(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 16c.85-.63 1.885-1 3-1s2.15.37 3 1m-5.5-5.5V10m5 .5V10"/><path d="M3 18.561v-6.517C3 7.05 7.03 3 12 3s9 4.05 9 9.044v6.517c0 1.162-.967 2.519-2 2c-.835-.42-2.223-.52-3 0c-.874.585-2.126.585-3 0c-.885-.593-1.649-.57-2.5 0c-.874.585-2.126.585-3 0c-.777-.52-1.665-.42-2.5 0c-1.033.519-2-.838-2-2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GhostSmile(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 15c.85.63 1.885 1 3 1s2.15-.37 3-1m-5.5-4.5V10m5 .5V10"/><path d="M3 18.561v-6.517C3 7.05 7.03 3 12 3s9 4.05 9 9.044v6.517c0 1.162-.967 2.519-2 2c-.835-.42-2.223-.52-3 0c-.874.585-2.126.585-3 0c-.885-.593-1.649-.57-2.5 0c-.874.585-2.126.585-3 0c-.777-.52-1.665-.42-2.5 0c-1.033.519-2-.838-2-2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gift(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 21v-9m0-5H7.95c-2.77 0-2.94-4 0-4C11.1 3 12 7 12 7m0 0h4.05c2.896 0 2.896-4 0-4C12.9 3 12 7 12 7"/><path d="M20 12v7a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-7m17 0V9a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v3z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitBranch(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6 15a3 3 0 1 1 0 6a3 3 0 0 1 0-6"/><path d="M18 9a3 3 0 1 1 0-6a3 3 0 0 1 0 6m0 0a9 9 0 0 1-9 9m-3-3V3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M14 12a2 2 0 1 1-4 0m4 0a2 2 0 1 0-4 0m4 0h3m-7 0H7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitCommit(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6m-9-3h6m6 0h6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitDiff(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18 21a3 3 0 1 0 0-6a3 3 0 0 0 0 6m0-6V7.5a2 2 0 0 0-2-2h-2.5"/><path d="M14.5 8L12 5.5L14.5 3M6 3a3 3 0 1 0 0 6a3 3 0 0 0 0-6m0 6v7.5a2 2 0 0 0 2 2h2.5"/><path d="m9.5 16l2.5 2.5L9.5 21"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/><path d="M14 12a2 2 0 1 1-4 0m4 0a2 2 0 1 0-4 0m4 0h3m-7 0H7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitMerge(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 9a3 3 0 1 0 0-6a3 3 0 0 0 0 6m9 9a3 3 0 1 0 6 0a3 3 0 0 0-6 0m0 0a9 9 0 0 1-9-9m0 0v12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/><path d="M14 12a2 2 0 1 1-4 0m4 0a2 2 0 1 0-4 0m4 0h3m-7 0H7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitPullRequest(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6 9a3 3 0 1 0 0-6a3 3 0 0 0 0 6m12 12a3 3 0 1 0 0-6a3 3 0 0 0 0 6M6 9v12m12-6V7.5a2 2 0 0 0-2-2h-2.5"/><path d="M14.5 8L12 5.5L14.5 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/><path d="M14 12a2 2 0 1 1-4 0m4 0a2 2 0 1 0-4 0m4 0h3m-7 0H7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GitWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/><path d="M14 12a2 2 0 1 1-4 0m4 0a2 2 0 1 0-4 0m4 0h3m-7 0H7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Globe(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 21a9 9 0 1 0 0-18m0 18a9 9 0 1 1 0-18m0 18c2.761 0 3.941-5.163 3.941-9c0-3.837-1.18-9-3.941-9m0 18c-2.761 0-3.941-5.163-3.941-9c0-3.837 1.18-9 3.941-9M3.5 9h17m-17 6h17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Grid(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 6.75c0-1.768 0-2.652.55-3.2C4.097 3 4.981 3 6.75 3c1.768 0 2.652 0 3.2.55c.55.548.55 1.432.55 3.2c0 1.768 0 2.652-.55 3.2c-.548.55-1.432.55-3.2.55c-1.768 0-2.652 0-3.2-.55C3 9.403 3 8.519 3 6.75m0 10.507c0-1.768 0-2.652.55-3.2c.548-.55 1.432-.55 3.2-.55c1.768 0 2.652 0 3.2.55c.55.548.55 1.432.55 3.2c0 1.768 0 2.652-.55 3.2c-.548.55-1.432.55-3.2.55c-1.768 0-2.652 0-3.2-.55C3 19.91 3 19.026 3 17.258M13.5 6.75c0-1.768 0-2.652.55-3.2c.548-.55 1.432-.55 3.2-.55c1.768 0 2.652 0 3.2.55c.55.548.55 1.432.55 3.2c0 1.768 0 2.652-.55 3.2c-.548.55-1.432.55-3.2.55c-1.768 0-2.652 0-3.2-.55c-.55-.548-.55-1.432-.55-3.2m0 10.507c0-1.768 0-2.652.55-3.2c.548-.55 1.432-.55 3.2-.55c1.768 0 2.652 0 3.2.55c.55.548.55 1.432.55 3.2c0 1.768 0 2.652-.55 3.2c-.548.55-1.432.55-3.2.55c-1.768 0-2.652 0-3.2-.55c-.55-.548-.55-1.432-.55-3.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GridOne(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.906 3.029V20.97m6.188-17.94v17.94M3.029 15.094H20.97M3.03 8.906h17.94M12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HardDrive(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.25 13h17.5M5.45 4.11l-2.162 7.847A8 8 0 0 0 3 14.082V19a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-4.918a8 8 0 0 0-.288-2.125L18.55 4.11A2 2 0 0 0 16.76 3H7.24a2 2 0 0 0-1.79 1.11M17 17h1m-5 0h1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hash(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.697 4L6.678 21M17.054 4l-3.019 17M21 8.781H3m18 7.438H3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HashCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m10.905 8l-1.437 8m4.937-8l-1.437 8m3.314-5.75H7.718m8.564 3.5H7.718"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HashHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m10.905 8l-1.437 8m4.937-8l-1.437 8m3.314-5.75H7.718m8.564 3.5H7.718M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HashOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132zM10.905 8l-1.437 8m4.937-8l-1.437 8m3.314-5.75H7.718m8.564 3.5H7.718"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HashSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m10.905 8l-1.437 8m4.937-8l-1.437 8m3.314-5.75H7.718m8.564 3.5H7.718M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HashWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m10.905 8l-1.437 8m4.937-8l-1.437 8m3.314-5.75H7.718m8.564 3.5H7.718M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heading(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.25 4.5v15m9.5-15v15M7.25 12h9.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadingOne(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 4.5v15m9.5-15v15M3.75 12h9.5m3.988-.508L19.625 9v10.492"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadingThree(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 4.5v15m9.5-15v15M3.75 12h9.5M16 9.5h5L17.5 14c2 0 3.5 1 3.5 3c0 2.74-3.408 3.2-5 1.45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadingTwo(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 11.92c0-3.226 5-3.226 5 0c0 2.85-5 4.966-5 7.5h5M3.75 4.5v15m9.5-15v15M3.75 12h9.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Headphones(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 17v-5a9 9 0 1 0-18 0v5"/><path d="M16 14.958c0-.511 0-.767.059-.97c.135-.468.49-.824.93-.934c1.272-.318 1.53.864 2.443 1.232l.069.028c.992.417 1.497 1.478 1.495 2.554v.357c-.003.95-.51 1.835-1.353 2.272c-.939.485-1.252 1.752-2.615 1.46c-.437-.094-.797-.429-.95-.883C16 19.843 16 19.54 16 18.938zm-8 4.084c0 .511 0 .766-.059.97c-.135.468-.49.824-.93.934c-1.272.318-1.53-.865-2.443-1.232a5.682 5.682 0 0 1-.062-.025c-.998-.418-1.504-1.48-1.502-2.557v-.364c.003-.946.509-1.828 1.353-2.265c.939-.485 1.252-1.752 2.615-1.46c.437.094.797.429.95.883c.078.231.078.533.078 1.136z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heart(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.75 3.5C5.127 3.5 3 5.76 3 8.547C3 14.125 12 20.5 12 20.5s9-6.375 9-11.953C21 5.094 18.873 3.5 16.25 3.5c-1.86 0-3.47 1.136-4.25 2.79c-.78-1.654-2.39-2.79-4.25-2.79"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartBroken(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.75 3.5C5.127 3.5 3 5.76 3 8.547C3 14.125 12 20.5 12 20.5s9-6.375 9-11.953C21 5.094 18.873 3.5 16.25 3.5c-1.86 0-3.47 1.136-4.25 2.79c-.78-1.654-2.39-2.79-4.25-2.79"/><path d="m12.15 6l-2 4l4 1l-2 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartCheck(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.75 3.5C5.127 3.5 3 5.76 3 8.547C3 14.125 12 20.5 12 20.5s9-6.375 9-11.953C21 5.094 18.873 3.5 16.25 3.5c-1.86 0-3.47 1.136-4.25 2.79c-.78-1.654-2.39-2.79-4.25-2.79"/><path d="m10.25 12.492l1.039 1.181c.095.109.267.1.351-.016l2.11-2.907"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.291 9.5a1.78 1.78 0 0 0-1.781 1.781c0 1.969 3.375 4.219 3.375 4.219s3.375-2.25 3.375-4.219c0-1.219-.797-1.781-1.781-1.781c-.698 0-1.302.4-1.594.985a1.781 1.781 0 0 0-1.594-.985"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.291 9.5a1.78 1.78 0 0 0-1.781 1.781c0 1.969 3.375 4.219 3.375 4.219s3.375-2.25 3.375-4.219c0-1.219-.797-1.781-1.781-1.781c-.698 0-1.302.4-1.594.985a1.781 1.781 0 0 0-1.594-.985"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartMinus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.75 3.5C5.127 3.5 3 5.76 3 8.547C3 14.125 12 20.5 12 20.5s9-6.375 9-11.953C21 5.094 18.873 3.5 16.25 3.5c-1.86 0-3.47 1.136-4.25 2.79c-.78-1.654-2.39-2.79-4.25-2.79M10 12h4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/><path d="M10.291 9.5a1.78 1.78 0 0 0-1.781 1.781c0 1.969 3.375 4.219 3.375 4.219s3.375-2.25 3.375-4.219c0-1.219-.797-1.781-1.781-1.781c-.698 0-1.302.4-1.594.985a1.781 1.781 0 0 0-1.594-.985"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartPlus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.75 3.5C5.127 3.5 3 5.76 3 8.547C3 14.125 12 20.5 12 20.5s9-6.375 9-11.953C21 5.094 18.873 3.5 16.25 3.5c-1.86 0-3.47 1.136-4.25 2.79c-.78-1.654-2.39-2.79-4.25-2.79M10 12h4m-2-2v4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.291 9.5a1.78 1.78 0 0 0-1.781 1.781c0 1.969 3.375 4.219 3.375 4.219s3.375-2.25 3.375-4.219c0-1.219-.797-1.781-1.781-1.781c-.698 0-1.302.4-1.594.985a1.781 1.781 0 0 0-1.594-.985"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.291 9.5a1.78 1.78 0 0 0-1.781 1.781c0 1.969 3.375 4.219 3.375 4.219s3.375-2.25 3.375-4.219c0-1.219-.797-1.781-1.781-1.781c-.698 0-1.302.4-1.594.985a1.781 1.781 0 0 0-1.594-.985"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartX(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.75 3.5C5.127 3.5 3 5.76 3 8.547C3 14.125 12 20.5 12 20.5s9-6.375 9-11.953C21 5.094 18.873 3.5 16.25 3.5c-1.86 0-3.47 1.136-4.25 2.79c-.78-1.654-2.39-2.79-4.25-2.79M10.5 11l2.828 2.828m0-2.828L10.5 13.828"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Home(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6.133 21C4.955 21 4 20.02 4 18.81v-8.802c0-.665.295-1.295.8-1.71l5.867-4.818a2.09 2.09 0 0 1 2.666 0l5.866 4.818c.506.415.801 1.045.801 1.71v8.802c0 1.21-.955 2.19-2.133 2.19z"/><path d="M9.5 21v-5.5a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2V21"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeCheck(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m9.68 13.323l1.379 1.575a.299.299 0 0 0 .466-.022l2.8-3.876"/><path d="M6.133 21C4.955 21 4 20.02 4 18.81v-8.802c0-.665.295-1.295.8-1.71l5.867-4.818a2.09 2.09 0 0 1 2.666 0l5.866 4.818c.506.415.801 1.045.801 1.71v8.802c0 1.21-.955 2.19-2.133 2.19z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeMinus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.5 13.5h5M6.133 21C4.955 21 4 20.02 4 18.81v-8.802c0-.665.295-1.295.8-1.71l5.867-4.818a2.09 2.09 0 0 1 2.666 0l5.866 4.818c.506.415.801 1.045.801 1.71v8.802c0 1.21-.955 2.19-2.133 2.19z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomePlus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.5 13.5h5M12 11v5m-5.867 5C4.955 21 4 20.02 4 18.81v-8.802c0-.665.295-1.295.8-1.71l5.867-4.818a2.09 2.09 0 0 1 2.666 0l5.866 4.818c.506.415.801 1.045.801 1.71v8.802c0 1.21-.955 2.19-2.133 2.19z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeSmile(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6.133 21C4.955 21 4 20.02 4 18.81v-8.802c0-.665.295-1.295.8-1.71l5.867-4.818a2.09 2.09 0 0 1 2.666 0l5.866 4.818c.506.415.801 1.045.801 1.71v8.802c0 1.21-.955 2.19-2.133 2.19z"/><path d="M9 16c.85.63 1.885 1 3 1s2.15-.37 3-1m-5.5-4.5V11m5 .5V11"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeX(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m10 11.5l4 4m0-4l-4 4M6.133 21C4.955 21 4 20.02 4 18.81v-8.802c0-.665.295-1.295.8-1.71l5.867-4.818a2.09 2.09 0 0 1 2.666 0l5.866 4.818c.506.415.801 1.045.801 1.71v8.802c0 1.21-.955 2.19-2.133 2.19z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Image(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6.5 8a2 2 0 1 0 4 0a2 2 0 0 0-4 0m14.427 1.99c-6.61-.908-12.31 4-11.927 10.51"/><path d="M3 13.066c2.78-.385 5.275.958 6.624 3.1"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 9a2 2 0 1 0 4 0a2 2 0 0 0-4 0m13.718 1.08c-6.38-.75-11.85 3.906-11.716 10.144"/><path d="M3.201 13.04c2.698-.294 5.106 1.036 6.423 3.126"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageRectangle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6 9a2 2 0 1 0 4 0a2 2 0 0 0-4 0m15.927-.01c-6.61-.908-12.31 4-11.927 10.51"/><path d="M2 13.066c2.78-.385 6.851 1.293 8.2 3.434"/><path d="M2 12c0-3.771 0-5.657 1.464-6.828C4.93 4 7.286 4 12 4c4.714 0 7.071 0 8.535 1.172C22 6.343 22 8.229 22 12c0 3.771 0 5.657-1.465 6.828C19.072 20 16.714 20 12 20s-7.071 0-8.536-1.172C2 17.657 2 15.771 2 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Inbox(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3.25 13h3.68a2 2 0 0 1 1.664.89l.812 1.22a2 2 0 0 0 1.664.89h1.86a2 2 0 0 0 1.664-.89l.812-1.22A2 2 0 0 1 17.07 13h3.68"/><path d="m5.45 4.11l-2.162 7.847A8 8 0 0 0 3 14.082V19a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-4.918a8 8 0 0 0-.288-2.125L18.55 4.11A2 2 0 0 0 16.76 3H7.24a2 2 0 0 0-1.79 1.11"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxArchive(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3.25 13h3.68a2 2 0 0 1 1.664.89l.812 1.22a2 2 0 0 0 1.664.89h1.86a2 2 0 0 0 1.664-.89l.812-1.22A2 2 0 0 1 17.07 13h3.68"/><path d="m5.45 4.11l-2.162 7.847A8 8 0 0 0 3 14.082V19a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-4.918a8 8 0 0 0-.288-2.125L18.55 4.11A2 2 0 0 0 16.76 3H7.24a2 2 0 0 0-1.79 1.11M9 9.5h6m-4.5-3h3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxCheck(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3.25 13h3.68a2 2 0 0 1 1.664.89l.812 1.22a2 2 0 0 0 1.664.89h1.86a2 2 0 0 0 1.664-.89l.812-1.22A2 2 0 0 1 17.07 13h3.68"/><path d="m5.45 4.11l-2.162 7.847A8 8 0 0 0 3 14.082V19a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-4.918a8 8 0 0 0-.288-2.125L18.55 4.11A2 2 0 0 0 16.76 3H7.24a2 2 0 0 0-1.79 1.11"/><path d="m10.3 8.742l1.034 1.182c.095.108.266.1.35-.017L13.784 7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxDown(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3.25 13h3.68a2 2 0 0 1 1.664.89l.812 1.22a2 2 0 0 0 1.664.89h1.86a2 2 0 0 0 1.664-.89l.812-1.22A2 2 0 0 1 17.07 13h3.68"/><path d="m5.45 4.11l-2.162 7.847A8 8 0 0 0 3 14.082V19a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-4.918a8 8 0 0 0-.288-2.125L18.55 4.11A2 2 0 0 0 16.76 3H7.24a2 2 0 0 0-1.79 1.11M12.003 6L12 11.28"/><path d="M14.5 9.347L12 12L9.5 9.347"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxMinus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3.25 13h3.68a2 2 0 0 1 1.664.89l.812 1.22a2 2 0 0 0 1.664.89h1.86a2 2 0 0 0 1.664-.89l.812-1.22A2 2 0 0 1 17.07 13h3.68"/><path d="m5.45 4.11l-2.162 7.847A8 8 0 0 0 3 14.082V19a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-4.918a8 8 0 0 0-.288-2.125L18.55 4.11A2 2 0 0 0 16.76 3H7.24a2 2 0 0 0-1.79 1.11M10 8.5h4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxPlus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3.25 13h3.68a2 2 0 0 1 1.664.89l.812 1.22a2 2 0 0 0 1.664.89h1.86a2 2 0 0 0 1.664-.89l.812-1.22A2 2 0 0 1 17.07 13h3.68"/><path d="m5.45 4.11l-2.162 7.847A8 8 0 0 0 3 14.082V19a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-4.918a8 8 0 0 0-.288-2.125L18.55 4.11A2 2 0 0 0 16.76 3H7.24a2 2 0 0 0-1.79 1.11M10 8.5h4m-2-2v4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxUp(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3.25 13h3.68a2 2 0 0 1 1.664.89l.812 1.22a2 2 0 0 0 1.664.89h1.86a2 2 0 0 0 1.664-.89l.812-1.22A2 2 0 0 1 17.07 13h3.68"/><path d="m5.45 4.11l-2.162 7.847A8 8 0 0 0 3 14.082V19a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-4.918a8 8 0 0 0-.288-2.125L18.55 4.11A2 2 0 0 0 16.76 3H7.24a2 2 0 0 0-1.79 1.11M12.003 12L12 6.72"/><path d="M14.5 8.653L12 6L9.5 8.653"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxX(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3.25 13h3.68a2 2 0 0 1 1.664.89l.812 1.22a2 2 0 0 0 1.664.89h1.86a2 2 0 0 0 1.664-.89l.812-1.22A2 2 0 0 1 17.07 13h3.68"/><path d="m5.45 4.11l-2.162 7.847A8 8 0 0 0 3 14.082V19a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-4.918a8 8 0 0 0-.288-2.125L18.55 4.11A2 2 0 0 0 16.76 3H7.24a2 2 0 0 0-1.79 1.11M10.5 7l3 3m0-3l-3 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Incognito(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 17.5a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0M3 10.412h18m-16-.189l.614-2.6c.545-2.31.818-3.466 1.632-4.139c.545-.45.81.073 1.601.468c1.004.502 2.177.103 3.186-.39c1.008-.494 2.973-.755 4.036-.393c.976.334.983.762 1.391 1.71c.733 1.703 1.114 3.54 1.54 5.344M10 17.5a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0m0-1.206l.658-.349a2.85 2.85 0 0 1 2.684 0l.658.349"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IndifferentCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 15.5h6m-5.5-5V10m5 .5V10"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IndifferentGhost(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 15.5h6m-5.5-5V10m5 .5V10M3 18.562v-6.518C3 7.05 7.03 3 12 3s9 4.05 9 9.044v6.517c0 1.162-.967 2.519-2 2c-.835-.42-2.223-.52-3 0c-.874.585-2.126.585-3 0c-.885-.593-1.649-.57-2.5 0c-.874.585-2.126.585-3 0c-.777-.52-1.665-.42-2.5 0c-1.033.519-2-.838-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IndifferentSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 15.5h6m-5.5-5V10m5 .5V10M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Infinity(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14 9.417C14.838 8.575 15.793 8 17 8a4 4 0 0 1 0 8c-4.5 0-5.5-8-10-8a4 4 0 1 0 0 8c1.207 0 2.162-.575 3-1.417"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Info(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 19.5v-10h-.5m0 10h1m-.5-14V5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 16v-5h-.5m0 5h1M12 8.5V8"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 16v-5h-.5m0 5h1M12 8.5V8"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/><path d="M12 16v-5h-.5m0 5h1M12 8.5V8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/><path d="M12 16v-5h-.5m0 5h1M12 8.5V8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoTriangle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 17v-5h-.5m0 5h1M12 9.5V9"/><path d="M5.98 10.762C8.608 5.587 9.92 3 12 3c2.08 0 3.393 2.587 6.02 7.762l.327.644c2.182 4.3 3.274 6.45 2.287 8.022C19.648 21 17.208 21 12.327 21h-.654c-4.88 0-7.321 0-8.307-1.572c-.987-1.572.105-3.722 2.287-8.022z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfoWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 16v-5h-.5m0 5h1M12 8.5V8"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Intersect(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.4 15.6h2.4c1.68 0 2.52 0 3.162-.327a3 3 0 0 0 1.311-1.311c.327-.642.327-1.482.327-3.162V8.4m-7.2 7.2v.9m0-.9v-2.4c0-1.68 0-2.52.327-3.162a3 3 0 0 1 1.311-1.311C10.68 8.4 11.52 8.4 13.2 8.4h2.4m-7.2 7.2h-.9m8.1-7.2h.9m-.9 0v-.9M8.4 19.65V21h1.8m5.4 0h-1.8m5.4 0H21v-1.8m0-5.4v1.8m0-5.4V8.4h-1.35M15.6 4.35V3h-1.8M8.4 3h1.8M4.8 3H3v1.8m0 5.4V8.4m0 5.4v1.8h1.35"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Key(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m15.362 9.065l1.32 1.32c.995.995 1.345-.84 2.734-1.07c.466-.078.877-.236 1.053-.752c.156-.456-.021-.885-.574-1.438L18.5 5.731M7.5 21a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9m3.5-8L21 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Keyboard(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2 11c0-2.828 0-4.243.879-5.121C3.757 5 5.172 5 8 5h8c2.828 0 4.243 0 5.121.879C22 6.757 22 8.172 22 11v2c0 2.828 0 4.243-.879 5.121C20.243 19 18.828 19 16 19H8c-2.828 0-4.243 0-5.121-.879C2 17.243 2 15.828 2 13zm5 5h10M5 9h3m3 0h3m3 0h2M5 12h2m3 0h3m3 0h3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyboardBrightnessHigh(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 17h4M3 17h3m6-9v3m6 6h3m-4.879-4.28L18.58 11M7.457 12.72L5 11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyboardBrightnessLow(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 17h8M3 17h1m8-8v1m6 2.207l.707-.707m-13 .707L5 11.5M20 17h1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Label(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m10.98 20.194l-7.298-7.298c-.37-.37-.58-.87-.587-1.392L3 4.015A1.002 1.002 0 0 1 4.015 3l7.489.095a2.005 2.005 0 0 1 1.392.587l7.298 7.298c.674.673 1.192 1.959.424 2.727l-6.91 6.91c-.769.769-2.055.25-2.728-.423M8.019 7.552l-.707-.707"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lamp(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 21v-9m0 9H9m3 0h3m-3-9h6l-2.513-7.702A2 2 0 0 0 13.614 3h-3.228a2 2 0 0 0-1.873 1.298L6 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayersOne(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m12 16l9-4l-9-4l-9 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayersThree(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m21 12l-9 4l-9-4m18 4l-9 4l-9-4m18-8l-9 4l-9-4l9-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayersTwo(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m21 14l-9 4l-9-4m18-4l-9 4l-9-4l9-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Layout(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 10v10M20.5 9.5h-17M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Leaf(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.973 18.028c7.625 4.576 13.726-1.525 12.963-12.964C7.498 4.302 1.398 10.403 5.973 18.028m0 0L4 20m1.973-1.972L10.1 13.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Leaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 12.07V20m0-8.07c0-4.353 3.538-7.887 7.919-7.93c.053.37.081.748.081 1.132c0 4.353-3.538 7.886-7.919 7.93A7.92 7.92 0 0 1 12 11.928m0 0C12 7.576 8.462 4.042 4.081 4A7.98 7.98 0 0 0 4 5.132c0 4.353 3.538 7.886 7.919 7.93A7.92 7.92 0 0 0 12 11.928"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterA(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.25 13.667L11.75 7l-2.5 6.667m5 0L15.5 17m-1.25-3.333h-5M8 17l1.25-3.333"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterAcircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M14 12.833L12 7.5l-2 5.333m4 0l1 2.667m-1-2.667h-4M9 15.5l1-2.667"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterAdiamond(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14 12.833L12 7.5l-2 5.333m4 0l1 2.667m-1-2.667h-4M9 15.5l1-2.667"/><path d="M2.707 10.295a2.41 2.41 0 0 0 0 3.41l7.588 7.588a2.409 2.409 0 0 0 3.41 0l7.588-7.588a2.409 2.409 0 0 0 0-3.41l-7.588-7.588a2.41 2.41 0 0 0-3.41 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterAhexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14 12.833L12 7.5l-2 5.333m4 0l1 2.667m-1-2.667h-4M9 15.5l1-2.667"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterAoctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14 12.833L12 7.5l-2 5.333m4 0l1 2.667m-1-2.667h-4M9 15.5l1-2.667"/><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterAsquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14 12.833L12 7.5l-2 5.333m4 0l1 2.667m-1-2.667h-4M9 15.5l1-2.667"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterAwaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14 12.833L12 7.5l-2 5.333m4 0l1 2.667m-1-2.667h-4M9 15.5l1-2.667"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterB(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12V7h4.589c2.74 0 3.124 4.072.57 5M9 12v5h4.589c2.74 0 3.124-4.072.57-5M9 12h5.158"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterBcircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M9.5 12V8h3.671c2.192 0 2.5 3.258.456 4M9.5 12v4h3.671c2.192 0 2.5-3.258.456-4M9.5 12h4.127"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterBdiamond(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.5 12V8h3.671c2.192 0 2.5 3.258.456 4M9.5 12v4h3.671c2.192 0 2.5-3.258.456-4M9.5 12h4.127"/><path d="M2.707 10.295a2.41 2.41 0 0 0 0 3.41l7.588 7.588a2.409 2.409 0 0 0 3.41 0l7.588-7.588a2.409 2.409 0 0 0 0-3.41l-7.588-7.588a2.41 2.41 0 0 0-3.41 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterBhexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.5 12V8h3.671c2.192 0 2.5 3.258.456 4M9.5 12v4h3.671c2.192 0 2.5-3.258.456-4M9.5 12h4.127"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterBoctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.5 12V8h3.671c2.192 0 2.5 3.258.456 4M9.5 12v4h3.671c2.192 0 2.5-3.258.456-4M9.5 12h4.127"/><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterBsquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.5 12V8h3.671c2.192 0 2.5 3.258.456 4M9.5 12v4h3.671c2.192 0 2.5-3.258.456-4M9.5 12h4.127"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterBwaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.5 12V8h3.671c2.192 0 2.5 3.258.456 4M9.5 12v4h3.671c2.192 0 2.5-3.258.456-4M9.5 12h4.127"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterC(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.5 9.5v-.312C15.5 7.979 14.52 7 13.313 7H10.5A2.5 2.5 0 0 0 8 9.5v5a2.5 2.5 0 0 0 2.5 2.5h2.813c1.208 0 2.187-.98 2.187-2.187V14.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterCcircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M15 10v-.25A1.75 1.75 0 0 0 13.25 8H11a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h2.25A1.75 1.75 0 0 0 15 14.25V14"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterCdiamond(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15 10v-.25A1.75 1.75 0 0 0 13.25 8H11a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h2.25A1.75 1.75 0 0 0 15 14.25V14"/><path d="M2.707 10.295a2.41 2.41 0 0 0 0 3.41l7.588 7.588a2.409 2.409 0 0 0 3.41 0l7.588-7.588a2.409 2.409 0 0 0 0-3.41l-7.588-7.588a2.41 2.41 0 0 0-3.41 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterChexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15 10v-.25A1.75 1.75 0 0 0 13.25 8H11a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h2.25A1.75 1.75 0 0 0 15 14.25V14"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterCoctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15 10v-.25A1.75 1.75 0 0 0 13.25 8H11a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h2.25A1.75 1.75 0 0 0 15 14.25V14"/><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterCsquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15 10v-.25A1.75 1.75 0 0 0 13.25 8H11a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h2.25A1.75 1.75 0 0 0 15 14.25V14"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterCwaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15 10v-.25A1.75 1.75 0 0 0 13.25 8H11a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h2.25A1.75 1.75 0 0 0 15 14.25V14"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterD(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 16.375v-8.75C8 7.28 8.28 7 8.625 7h2.5a4.375 4.375 0 0 1 4.375 4.375v1.25A4.375 4.375 0 0 1 11.125 17h-2.5A.625.625 0 0 1 8 16.375"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterDcircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M9.5 15.5v-7A.5.5 0 0 1 10 8h2a3.5 3.5 0 0 1 3.5 3.5v1A3.5 3.5 0 0 1 12 16h-2a.5.5 0 0 1-.5-.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterDdiamond(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.5 15.5v-7A.5.5 0 0 1 10 8h2a3.5 3.5 0 0 1 3.5 3.5v1A3.5 3.5 0 0 1 12 16h-2a.5.5 0 0 1-.5-.5"/><path d="M2.707 10.295a2.41 2.41 0 0 0 0 3.41l7.588 7.588a2.409 2.409 0 0 0 3.41 0l7.588-7.588a2.409 2.409 0 0 0 0-3.41l-7.588-7.588a2.41 2.41 0 0 0-3.41 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterDhexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.5 15.5v-7A.5.5 0 0 1 10 8h2a3.5 3.5 0 0 1 3.5 3.5v1A3.5 3.5 0 0 1 12 16h-2a.5.5 0 0 1-.5-.5"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterDoctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.5 15.5v-7A.5.5 0 0 1 10 8h2a3.5 3.5 0 0 1 3.5 3.5v1A3.5 3.5 0 0 1 12 16h-2a.5.5 0 0 1-.5-.5"/><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterDsquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.5 15.5v-7A.5.5 0 0 1 10 8h2a3.5 3.5 0 0 1 3.5 3.5v1A3.5 3.5 0 0 1 12 16h-2a.5.5 0 0 1-.5-.5"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterDwaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.5 15.5v-7A.5.5 0 0 1 10 8h2a3.5 3.5 0 0 1 3.5 3.5v1A3.5 3.5 0 0 1 12 16h-2a.5.5 0 0 1-.5-.5"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterE(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.25 7H9v5m6.25 5H9v-5m0 0h5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterEcircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M14.5 8h-5v4m5 4h-5v-4m0 0h4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterEdiamond(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.5 8h-5v4m5 4h-5v-4m0 0h4"/><path d="M2.707 10.295a2.41 2.41 0 0 0 0 3.41l7.588 7.588a2.409 2.409 0 0 0 3.41 0l7.588-7.588a2.409 2.409 0 0 0 0-3.41l-7.588-7.588a2.41 2.41 0 0 0-3.41 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterEhexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.5 8h-5v4m5 4h-5v-4m0 0h4"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterEoctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.5 8h-5v4m5 4h-5v-4m0 0h4"/><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterEsquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.5 8h-5v4m5 4h-5v-4m0 0h4"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterEwaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.5 8h-5v4m5 4h-5v-4m0 0h4"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterF(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.5 7H9.25v5m0 0v5m0-5h5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterFcircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M14.75 8h-5v4m0 0v4m0-4h4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterFdiamond(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.75 8h-5v4m0 0v4m0-4h4"/><path d="M2.707 10.295a2.41 2.41 0 0 0 0 3.41l7.588 7.588a2.409 2.409 0 0 0 3.41 0l7.588-7.588a2.409 2.409 0 0 0 0-3.41l-7.588-7.588a2.41 2.41 0 0 0-3.41 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterFhexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.75 8h-5v4m0 0v4m0-4h4"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterFoctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.75 8h-5v4m0 0v4m0-4h4"/><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterFsquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.75 8h-5v4m0 0v4m0-4h4"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterFwaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.75 8h-5v4m0 0v4m0-4h4"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterG(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13.167 12h1.458c.345 0 .625.28.625.625v3.75c0 .345-.28.625-.625.625h-3.75A1.875 1.875 0 0 1 9 15.125v-6.25C9 7.839 9.84 7 10.875 7h2.5c1.036 0 1.875.84 1.875 1.875"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterGcircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M12.833 12H14a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-.5.5h-3a1.5 1.5 0 0 1-1.5-1.5v-5A1.5 1.5 0 0 1 11 8h2a1.5 1.5 0 0 1 1.5 1.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterGdiamond(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12.833 12H14a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-.5.5h-3a1.5 1.5 0 0 1-1.5-1.5v-5A1.5 1.5 0 0 1 11 8h2a1.5 1.5 0 0 1 1.5 1.5"/><path d="M2.707 10.295a2.41 2.41 0 0 0 0 3.41l7.588 7.588a2.409 2.409 0 0 0 3.41 0l7.588-7.588a2.409 2.409 0 0 0 0-3.41l-7.588-7.588a2.41 2.41 0 0 0-3.41 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterGhexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12.833 12H14a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-.5.5h-3a1.5 1.5 0 0 1-1.5-1.5v-5A1.5 1.5 0 0 1 11 8h2a1.5 1.5 0 0 1 1.5 1.5"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterGoctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12.833 12H14a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-.5.5h-3a1.5 1.5 0 0 1-1.5-1.5v-5A1.5 1.5 0 0 1 11 8h2a1.5 1.5 0 0 1 1.5 1.5"/><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterGsquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12.833 12H14a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-.5.5h-3a1.5 1.5 0 0 1-1.5-1.5v-5A1.5 1.5 0 0 1 11 8h2a1.5 1.5 0 0 1 1.5 1.5"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterGwaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12.833 12H14a.5.5 0 0 1 .5.5v3a.5.5 0 0 1-.5.5h-3a1.5 1.5 0 0 1-1.5-1.5v-5A1.5 1.5 0 0 1 11 8h2a1.5 1.5 0 0 1 1.5 1.5"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterH(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 7v5m0 0v5m0-5h6.25m0-5v5m0 0v5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterHcircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M9.5 8v4m0 0v4m0-4h5m0-4v4m0 0v4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterHdiamond(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.5 8v4m0 0v4m0-4h5m0-4v4m0 0v4M2.707 10.295a2.41 2.41 0 0 0 0 3.41l7.588 7.588a2.409 2.409 0 0 0 3.41 0l7.588-7.588a2.409 2.409 0 0 0 0-3.41l-7.588-7.588a2.41 2.41 0 0 0-3.41 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterHhexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.5 8v4m0 0v4m0-4h5m0-4v4m0 0v4m6-.2V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterHoctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.5 8v4m0 0v4m0-4h5m0-4v4m0 0v4M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterHsquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.5 8v4m0 0v4m0-4h5m0-4v4m0 0v4M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterHwaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.5 8v4m0 0v4m0-4h5m0-4v4m0 0v4M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterI(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 7h3.125m0 0h3.125m-3.125 0v10m3.125 0h-3.125m0 0H9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterIcircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M9.5 8H12m0 0h2.5M12 8v8m2.5 0H12m0 0H9.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterIdiamond(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.5 8H12m0 0h2.5M12 8v8m2.5 0H12m0 0H9.5m-6.793-5.705a2.41 2.41 0 0 0 0 3.41l7.588 7.588a2.409 2.409 0 0 0 3.41 0l7.588-7.588a2.409 2.409 0 0 0 0-3.41l-7.588-7.588a2.41 2.41 0 0 0-3.41 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterIhexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.5 8H12m0 0h2.5M12 8v8m2.5 0H12m0 0H9.5m11-.2V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterIoctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.5 8H12m0 0h2.5M12 8v8m2.5 0H12m0 0H9.5M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterIsquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.5 8H12m0 0h2.5M12 8v8m2.5 0H12m0 0H9.5M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterIwaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.5 8H12m0 0h2.5M12 8v8m2.5 0H12m0 0H9.5m.213-12.36c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterJ(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.25 7v7.5a2.5 2.5 0 0 1-2.5 2.5H11.5A2.5 2.5 0 0 1 9 14.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterJcircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M14.25 8v6a2 2 0 0 1-2 2h-1a2 2 0 0 1-2-2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterJdiamond(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.25 8v6a2 2 0 0 1-2 2h-1a2 2 0 0 1-2-2"/><path d="M2.707 10.295a2.41 2.41 0 0 0 0 3.41l7.588 7.588a2.409 2.409 0 0 0 3.41 0l7.588-7.588a2.409 2.409 0 0 0 0-3.41l-7.588-7.588a2.41 2.41 0 0 0-3.41 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterJhexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.25 8v6a2 2 0 0 1-2 2h-1a2 2 0 0 1-2-2"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterJoctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.25 8v6a2 2 0 0 1-2 2h-1a2 2 0 0 1-2-2"/><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterJsquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.25 8v6a2 2 0 0 1-2 2h-1a2 2 0 0 1-2-2"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterJwaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.25 8v6a2 2 0 0 1-2 2h-1a2 2 0 0 1-2-2"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterK(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 7v5m0 0v5m0-5h.625m0 0l5.625 5m-5.625-5l5.625-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterKcircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M9.75 8v4m0 0v4m0-4h.5m0 0l4.5 4m-4.5-4l4.5-4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterKdiamond(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.75 8v4m0 0v4m0-4h.5m0 0l4.5 4m-4.5-4l4.5-4M2.707 10.295a2.41 2.41 0 0 0 0 3.41l7.588 7.588a2.409 2.409 0 0 0 3.41 0l7.588-7.588a2.409 2.409 0 0 0 0-3.41l-7.588-7.588a2.41 2.41 0 0 0-3.41 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterKhexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.75 8v4m0 0v4m0-4h.5m0 0l4.5 4m-4.5-4l4.5-4m5.75 7.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterKoctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.75 8v4m0 0v4m0-4h.5m0 0l4.5 4m-4.5-4l4.5-4M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterKsquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.75 8v4m0 0v4m0-4h.5m0 0l4.5 4m-4.5-4l4.5-4M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterKwaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.75 8v4m0 0v4m0-4h.5m0 0l4.5 4m-4.5-4l4.5-4M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterL(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 7v9.5h7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterLcircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M9.5 8v7.5H15"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterLdiamond(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.5 8v7.5H15"/><path d="M2.707 10.295a2.41 2.41 0 0 0 0 3.41l7.588 7.588a2.409 2.409 0 0 0 3.41 0l7.588-7.588a2.409 2.409 0 0 0 0-3.41l-7.588-7.588a2.41 2.41 0 0 0-3.41 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterLhexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.5 8v7.5H15"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterLoctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.5 8v7.5H15"/><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterLsquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.5 8v7.5H15"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterLwaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.5 8v7.5H15"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterM(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 17V7l3.75 5l3.75-5v10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterMcircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M9 16V8l3 4l3-4v8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterMdiamond(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 16V8l3 4l3-4v8"/><path d="M2.707 10.295a2.41 2.41 0 0 0 0 3.41l7.588 7.588a2.409 2.409 0 0 0 3.41 0l7.588-7.588a2.409 2.409 0 0 0 0-3.41l-7.588-7.588a2.41 2.41 0 0 0-3.41 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterMhexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 16V8l3 4l3-4v8"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterMoctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 16V8l3 4l3-4v8"/><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterMsquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 16V8l3 4l3-4v8"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterMwaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 16V8l3 4l3-4v8"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterN(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 17V7l7.5 10V7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterNcircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M9 16V8l6 8V8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterNdiamond(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 16V8l6 8V8"/><path d="M2.707 10.295a2.41 2.41 0 0 0 0 3.41l7.588 7.588a2.409 2.409 0 0 0 3.41 0l7.588-7.588a2.409 2.409 0 0 0 0-3.41l-7.588-7.588a2.41 2.41 0 0 0-3.41 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterNhexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 16V8l6 8V8"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterNoctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 16V8l6 8V8"/><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterNsquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 16V8l6 8V8"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterNwaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 16V8l6 8V8"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterO(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 14.5v-5A2.5 2.5 0 0 1 10.5 7H13a2.5 2.5 0 0 1 2.5 2.5v5A2.5 2.5 0 0 1 13 17h-2.5A2.5 2.5 0 0 1 8 14.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterOcircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M9 14v-4a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-2a2 2 0 0 1-2-2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterOdiamond(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 14v-4a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-2a2 2 0 0 1-2-2"/><path d="M2.707 10.295a2.41 2.41 0 0 0 0 3.41l7.588 7.588a2.409 2.409 0 0 0 3.41 0l7.588-7.588a2.409 2.409 0 0 0 0-3.41l-7.588-7.588a2.41 2.41 0 0 0-3.41 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterOhexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 14v-4a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-2a2 2 0 0 1-2-2"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterOoctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 14v-4a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-2a2 2 0 0 1-2-2"/><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterOsquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 14v-4a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-2a2 2 0 0 1-2-2"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterOwaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 14v-4a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-2a2 2 0 0 1-2-2"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterP(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12V7.625C9 7.28 9.28 7 9.625 7h3.75c1.036 0 1.875.84 1.875 1.875v1.25c0 1.036-.84 1.875-1.875 1.875zm0 0v5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterPcircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M9.75 12V8.5a.5.5 0 0 1 .5-.5h3a1.5 1.5 0 0 1 1.5 1.5v1a1.5 1.5 0 0 1-1.5 1.5zm0 0v4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterPdiamond(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.75 12V8.5a.5.5 0 0 1 .5-.5h3a1.5 1.5 0 0 1 1.5 1.5v1a1.5 1.5 0 0 1-1.5 1.5zm0 0v4"/><path d="M2.707 10.295a2.41 2.41 0 0 0 0 3.41l7.588 7.588a2.409 2.409 0 0 0 3.41 0l7.588-7.588a2.409 2.409 0 0 0 0-3.41l-7.588-7.588a2.41 2.41 0 0 0-3.41 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterPhexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.75 12V8.5a.5.5 0 0 1 .5-.5h3a1.5 1.5 0 0 1 1.5 1.5v1a1.5 1.5 0 0 1-1.5 1.5zm0 0v4"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterPoctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.75 12V8.5a.5.5 0 0 1 .5-.5h3a1.5 1.5 0 0 1 1.5 1.5v1a1.5 1.5 0 0 1-1.5 1.5zm0 0v4"/><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterPsquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.75 12V8.5a.5.5 0 0 1 .5-.5h3a1.5 1.5 0 0 1 1.5 1.5v1a1.5 1.5 0 0 1-1.5 1.5zm0 0v4"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterPwaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.75 12V8.5a.5.5 0 0 1 .5-.5h3a1.5 1.5 0 0 1 1.5 1.5v1a1.5 1.5 0 0 1-1.5 1.5zm0 0v4"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterQ(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.385 16.882V9.471A2.466 2.466 0 0 0 12.923 7h-2.462A2.466 2.466 0 0 0 8 9.47v4.942a2.466 2.466 0 0 0 2.461 2.47zm0 0l-2.462-2.47m2.462 2.47L16 17.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterQcircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M15 16v-6a2 2 0 0 0-2-2h-2a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2zm0 0l-2-2m2 2l.5.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterQdiamond(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15 16v-6a2 2 0 0 0-2-2h-2a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2zm0 0l-2-2m2 2l.5.5"/><path d="M2.707 10.295a2.41 2.41 0 0 0 0 3.41l7.588 7.588a2.409 2.409 0 0 0 3.41 0l7.588-7.588a2.409 2.409 0 0 0 0-3.41l-7.588-7.588a2.41 2.41 0 0 0-3.41 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterQhexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15 16v-6a2 2 0 0 0-2-2h-2a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2zm0 0l-2-2m2 2l.5.5"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterQoctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15 16v-6a2 2 0 0 0-2-2h-2a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2zm0 0l-2-2m2 2l.5.5"/><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterQsquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15 16v-6a2 2 0 0 0-2-2h-2a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2zm0 0l-2-2m2 2l.5.5"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterQwaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15 16v-6a2 2 0 0 0-2-2h-2a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2zm0 0l-2-2m2 2l.5.5"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterR(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12V7.625C9 7.28 9.28 7 9.625 7h3.75c1.036 0 1.875.84 1.875 1.875v1.25c0 1.036-.84 1.875-1.875 1.875h-2.187M9 12v5m0-5h2.188m0 0l3.75 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterRcircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M9.75 12V8.5a.5.5 0 0 1 .5-.5h3a1.5 1.5 0 0 1 1.5 1.5v1a1.5 1.5 0 0 1-1.5 1.5H11.5m-1.75 0v4m0-4h1.75m0 0l3 4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterRdiamond(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.75 12V8.5a.5.5 0 0 1 .5-.5h3a1.5 1.5 0 0 1 1.5 1.5v1a1.5 1.5 0 0 1-1.5 1.5H11.5m-1.75 0v4m0-4h1.75m0 0l3 4"/><path d="M2.707 10.295a2.41 2.41 0 0 0 0 3.41l7.588 7.588a2.409 2.409 0 0 0 3.41 0l7.588-7.588a2.409 2.409 0 0 0 0-3.41l-7.588-7.588a2.41 2.41 0 0 0-3.41 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterRhexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.75 12V8.5a.5.5 0 0 1 .5-.5h3a1.5 1.5 0 0 1 1.5 1.5v1a1.5 1.5 0 0 1-1.5 1.5H11.5m-1.75 0v4m0-4h1.75m0 0l3 4"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterRoctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.75 12V8.5a.5.5 0 0 1 .5-.5h3a1.5 1.5 0 0 1 1.5 1.5v1a1.5 1.5 0 0 1-1.5 1.5H11.5m-1.75 0v4m0-4h1.75m0 0l3 4"/><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterRsquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.75 12V8.5a.5.5 0 0 1 .5-.5h3a1.5 1.5 0 0 1 1.5 1.5v1a1.5 1.5 0 0 1-1.5 1.5H11.5m-1.75 0v4m0-4h1.75m0 0l3 4"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterRwaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.75 12V8.5a.5.5 0 0 1 .5-.5h3a1.5 1.5 0 0 1 1.5 1.5v1a1.5 1.5 0 0 1-1.5 1.5H11.5m-1.75 0v4m0-4h1.75m0 0l3 4"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterS(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.25 8.563V8.25C15.25 7.56 14.69 7 14 7h-3.75C9.56 7 9 7.56 9 8.25v1.23c0 .767.467 1.457 1.179 1.742l3.892 1.556a1.875 1.875 0 0 1 1.179 1.741v1.231c0 .69-.56 1.25-1.25 1.25h-3.75C9.56 17 9 16.44 9 15.75v-.312"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterScircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M14.5 9.25V9a1 1 0 0 0-1-1h-3a1 1 0 0 0-1 1v.984a1.5 1.5 0 0 0 .943 1.393l3.114 1.246c.57.228.943.78.943 1.393V15a1 1 0 0 1-1 1h-3a1 1 0 0 1-1-1v-.25"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterSdiamond(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.5 9.25V9a1 1 0 0 0-1-1h-3a1 1 0 0 0-1 1v.984a1.5 1.5 0 0 0 .943 1.393l3.114 1.246c.57.228.943.78.943 1.393V15a1 1 0 0 1-1 1h-3a1 1 0 0 1-1-1v-.25"/><path d="M2.707 10.295a2.41 2.41 0 0 0 0 3.41l7.588 7.588a2.409 2.409 0 0 0 3.41 0l7.588-7.588a2.409 2.409 0 0 0 0-3.41l-7.588-7.588a2.41 2.41 0 0 0-3.41 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterShexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.5 9.25V9a1 1 0 0 0-1-1h-3a1 1 0 0 0-1 1v.984a1.5 1.5 0 0 0 .943 1.393l3.114 1.246c.57.228.943.78.943 1.393V15a1 1 0 0 1-1 1h-3a1 1 0 0 1-1-1v-.25"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterSoctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.5 9.25V9a1 1 0 0 0-1-1h-3a1 1 0 0 0-1 1v.984a1.5 1.5 0 0 0 .943 1.393l3.114 1.246c.57.228.943.78.943 1.393V15a1 1 0 0 1-1 1h-3a1 1 0 0 1-1-1v-.25"/><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterSsquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.5 9.25V9a1 1 0 0 0-1-1h-3a1 1 0 0 0-1 1v.984a1.5 1.5 0 0 0 .943 1.393l3.114 1.246c.57.228.943.78.943 1.393V15a1 1 0 0 1-1 1h-3a1 1 0 0 1-1-1v-.25"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterSwaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.5 9.25V9a1 1 0 0 0-1-1h-3a1 1 0 0 0-1 1v.984a1.5 1.5 0 0 0 .943 1.393l3.114 1.246c.57.228.943.78.943 1.393V15a1 1 0 0 1-1 1h-3a1 1 0 0 1-1-1v-.25"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterT(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 7h3.75m0 0h3.75m-3.75 0v10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterTcircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M9 8.25h3m0 0h3m-3 0v8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterTdiamond(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 8.25h3m0 0h3m-3 0v8m-9.293-5.955a2.41 2.41 0 0 0 0 3.41l7.588 7.588a2.409 2.409 0 0 0 3.41 0l7.588-7.588a2.409 2.409 0 0 0 0-3.41l-7.588-7.588a2.41 2.41 0 0 0-3.41 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterThexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 8.25h3m0 0h3m-3 0v8m8.5-.45V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterToctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 8.25h3m0 0h3m-3 0v8M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterTsquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 8.25h3m0 0h3m-3 0v8M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterTwaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 8.25h3m0 0h3m-3 0v8M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterU(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 7v7.5a2.5 2.5 0 0 0 2.5 2.5H13a2.5 2.5 0 0 0 2.5-2.5V7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterUcircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M9 8v6a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2V8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterUdiamond(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 8v6a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2V8"/><path d="M2.707 10.295a2.41 2.41 0 0 0 0 3.41l7.588 7.588a2.409 2.409 0 0 0 3.41 0l7.588-7.588a2.409 2.409 0 0 0 0-3.41l-7.588-7.588a2.41 2.41 0 0 0-3.41 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterUhexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 8v6a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2V8"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterUoctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 8v6a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2V8"/><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterUsquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 8v6a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2V8"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterUwaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 8v6a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2V8"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterV(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8 7l3.75 10L15.5 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterVcircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="m9 8.25l3 8l3-8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterVdiamond(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m9 8.25l3 8l3-8"/><path d="M2.707 10.295a2.41 2.41 0 0 0 0 3.41l7.588 7.588a2.409 2.409 0 0 0 3.41 0l7.588-7.588a2.409 2.409 0 0 0 0-3.41l-7.588-7.588a2.41 2.41 0 0 0-3.41 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterVhexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m9 8.25l3 8l3-8"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterVoctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m9 8.25l3 8l3-8"/><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterVsquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m9 8.25l3 8l3-8"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterVwaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m9 8.25l3 8l3-8"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterW(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 7v10l3.75-5l3.75 5V7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterWcircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M9 8v8l3-4l3 4V8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterWdiamond(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 8v8l3-4l3 4V8"/><path d="M2.707 10.295a2.41 2.41 0 0 0 0 3.41l7.588 7.588a2.409 2.409 0 0 0 3.41 0l7.588-7.588a2.409 2.409 0 0 0 0-3.41l-7.588-7.588a2.41 2.41 0 0 0-3.41 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterWhexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 8v8l3-4l3 4V8"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterWoctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 8v8l3-4l3 4V8"/><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterWsquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 8v8l3-4l3 4V8"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterWwaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 8v8l3-4l3 4V8"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterX(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8 7l7.5 10M8 17l7.5-10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterXcircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="m9 8l6 8m-6 0l6-8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterXdiamond(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9 8l6 8m-6 0l6-8M2.707 10.295a2.41 2.41 0 0 0 0 3.41l7.588 7.588a2.409 2.409 0 0 0 3.41 0l7.588-7.588a2.409 2.409 0 0 0 0-3.41l-7.588-7.588a2.41 2.41 0 0 0-3.41 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterXhexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9 8l6 8m-6 0l6-8m5.5 7.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterXoctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9 8l6 8m-6 0l6-8M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterXsquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9 8l6 8m-6 0l6-8M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterXwaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9 8l6 8m-6 0l6-8M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterY(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8 7l3.75 5m3.75-5l-3.75 5m0 0v5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterYcircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="m9 8.25l3 4m3-4l-3 4m0 0v4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterYdiamond(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9 8.25l3 4m3-4l-3 4m0 0v4m-9.293-5.955a2.41 2.41 0 0 0 0 3.41l7.588 7.588a2.409 2.409 0 0 0 3.41 0l7.588-7.588a2.409 2.409 0 0 0 0-3.41l-7.588-7.588a2.41 2.41 0 0 0-3.41 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterYhexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9 8.25l3 4m3-4l-3 4m0 0v4m8.5-.45V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterYoctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9 8.25l3 4m3-4l-3 4m0 0v4M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterYsquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9 8.25l3 4m3-4l-3 4m0 0v4M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterYwaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9 8.25l3 4m3-4l-3 4m0 0v4M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterZ(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 7h7.5L8 17h7.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterZcircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M9 8h6l-6 8h6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterZdiamond(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 8h6l-6 8h6"/><path d="M2.707 10.295a2.41 2.41 0 0 0 0 3.41l7.588 7.588a2.409 2.409 0 0 0 3.41 0l7.588-7.588a2.409 2.409 0 0 0 0-3.41l-7.588-7.588a2.41 2.41 0 0 0-3.41 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterZhexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 8h6l-6 8h6"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterZoctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 8h6l-6 8h6"/><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterZsquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 8h6l-6 8h6"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LetterZwaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 8h6l-6 8h6"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lightning(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.748 3.572c.059-.503-.532-.777-.835-.388L4.11 13.197c-.258.33-.038.832.364.832h6.988c.285 0 .506.267.47.57l-.68 5.83c-.06.502.53.776.834.387l7.802-10.013c.258-.33.038-.832-.364-.832h-6.988c-.285 0-.506-.267-.47-.57z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightningSlash(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3 21l6.971-6.971M21 3l-6.971 6.971m-4.058 4.058h1.492c.285 0 .506.267.47.57l-.68 5.83c-.06.502.53.776.834.387l7.802-10.013c.258-.33.038-.832-.364-.832h-5.496m-4.058 4.058l4.058-4.058M5.5 14.029H4.475c-.402 0-.622-.502-.364-.832l7.802-10.013c.303-.389.894-.115.835.388l-.46 3.928"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Like(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.54 10.105h5.533c2.546 0-.764 10.895-2.588 10.895H4.964A.956.956 0 0 1 4 20.053v-9.385c0-.347.193-.666.502-.832C6.564 8.73 8.983 7.824 10.18 5.707l1.28-2.266A.874.874 0 0 1 12.222 3c3.18 0 2.237 4.63 1.805 6.47a.52.52 0 0 0 .513.635"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LineChartCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="m16 10l-3.5 3.5l-2-2L8 14"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LineChartHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/><path d="m16 10l-3.5 3.5l-2-2L8 14"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LineChartOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/><path d="m16 10l-3.5 3.5l-2-2L8 14"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LineChartSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/><path d="m16 10l-3.5 3.5l-2-2L8 14"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LineChartWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/><path d="m16 10l-3.5 3.5l-2-2L8 14"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Link(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m15.988 13l3.902-3.902c1.437-1.437 1.485-3.718.107-5.095c-1.377-1.378-3.658-1.33-5.095.107L11 8.012m2 7.95l-3.892 3.88c-1.432 1.43-3.64 1.615-5.082.107c-1.442-1.507-1.326-3.639.107-5.068L8.025 11M9 15l6-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkOne(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m10.936 8.324l4.38-4.385c1.276-1.276 3.372-1.248 4.683.063c1.31 1.312 1.338 3.41.062 4.686l-3.803 3.807m-8.516-.99L3.94 15.312c-1.277 1.276-1.25 3.374.06 4.686c1.31 1.311 3.407 1.34 4.683.063l4.38-4.385m-2.065-2.666c-1.311-1.311-1.34-3.41-.063-4.686m2.128 2.603c1.312 1.311 1.34 3.41.063 4.686"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkTwo(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14 15.5h3.4c1.988 0 3.6-1.567 3.6-3.5s-1.612-3.5-3.6-3.5H14m-4 7l-3.397-.007c-1.987-.003-3.647-1.426-3.602-3.502c.045-2.075 1.606-3.494 3.593-3.491l3.397.007M7.757 12h8.486"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func List(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 6.5h12M8 12h12M8 17.5h12M4 6.5h1M4 12h1m-1 5.5h1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListCheck(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.5 16.039L4.688 19.5L3 18.346m4.5-8.077l-2.812 3.462L3 12.577M7.5 4.5L4.688 7.962L3 6.808M11 17.5h10M11 12h10M11 6.5h10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListNumber(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11 17.5h10M11 12h10M11 6.5h10M3.5 15.455v-.174c0-.777.672-1.406 1.5-1.406h.04c.807 0 1.46.613 1.46 1.368c0 .33-.114.65-.324.912L3.5 19.5h3m-3-14.063l2-.937v5.625"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Location(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12.56 20.82a.964.964 0 0 1-1.12 0C6.611 17.378 1.486 10.298 6.667 5.182A7.592 7.592 0 0 1 12 3c2 0 3.919.785 5.333 2.181c5.181 5.116.056 12.196-4.773 15.64"/><path d="M12 12a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationCheck(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12.56 20.82a.964.964 0 0 1-1.12 0C6.611 17.378 1.486 10.298 6.667 5.182A7.592 7.592 0 0 1 12 3c2 0 3.919.785 5.333 2.181c5.181 5.116.056 12.196-4.773 15.64"/><path d="m9.6 10.323l1.379 1.575a.299.299 0 0 0 .466-.022L14.245 8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationMinus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.56 20.82a.964.964 0 0 1-1.12 0C6.611 17.378 1.486 10.298 6.667 5.182A7.592 7.592 0 0 1 12 3c2 0 3.919.785 5.333 2.181c5.181 5.116.056 12.196-4.773 15.64M9.5 10h5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationPlus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.56 20.82a.964.964 0 0 1-1.12 0C6.611 17.378 1.486 10.298 6.667 5.182A7.592 7.592 0 0 1 12 3c2 0 3.919.785 5.333 2.181c5.181 5.116.056 12.196-4.773 15.64M9.5 10h5M12 7.5v5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationX(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.56 20.82a.964.964 0 0 1-1.12 0C6.611 17.378 1.486 10.298 6.667 5.182A7.592 7.592 0 0 1 12 3c2 0 3.919.785 5.333 2.181c5.181 5.116.056 12.196-4.773 15.64M10 8l4 4m0-4l-4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lock(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 10V8c0-2.761 1.239-5 4-5s4 2.239 4 5v2M3.5 17.8v-4.6c0-1.12 0-1.68.218-2.107a2 2 0 0 1 .874-.875c.428-.217.988-.217 2.108-.217h10.6c1.12 0 1.68 0 2.108.217a2 2 0 0 1 .874.874c.218.428.218.988.218 2.108v4.6c0 1.12 0 1.68-.218 2.108a2 2 0 0 1-.874.874C18.98 21 18.42 21 17.3 21H6.7c-1.12 0-1.68 0-2.108-.218a2 2 0 0 1-.874-.874C3.5 19.481 3.5 18.921 3.5 17.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M14.004 10.947V8.925c0-2.641-4.008-2.491-4.008 0v2.021m-.994 0h5.996c.553 0 1.002.453 1.002 1.011v3.032c0 .558-.449 1.011-1.002 1.011H9.002A1.006 1.006 0 0 1 8 14.99v-3.033c0-.558.449-1.01 1.002-1.01"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/><path d="M14.004 10.947V8.925c0-2.641-4.008-2.491-4.008 0v2.021m-.994 0h5.996c.553 0 1.002.453 1.002 1.011v3.032c0 .558-.449 1.011-1.002 1.011H9.002A1.006 1.006 0 0 1 8 14.99v-3.033c0-.558.449-1.01 1.002-1.01"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockKeyhole(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 10V8c0-2.761 1.239-5 4-5s4 2.239 4 5v2M3.5 17.8v-4.6c0-1.12 0-1.68.218-2.107a2 2 0 0 1 .874-.875c.428-.217.988-.217 2.108-.217h10.6c1.12 0 1.68 0 2.108.217a2 2 0 0 1 .874.874c.218.428.218.988.218 2.108v4.6c0 1.12 0 1.68-.218 2.108a2 2 0 0 1-.874.874C18.98 21 18.42 21 17.3 21H6.7c-1.12 0-1.68 0-2.108-.218a2 2 0 0 1-.874-.874C3.5 19.481 3.5 18.921 3.5 17.8M16 14v3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/><path d="M14.004 10.947V8.925c0-2.641-4.008-2.491-4.008 0v2.021m-.994 0h5.996c.553 0 1.002.453 1.002 1.011v3.032c0 .558-.449 1.011-1.002 1.011H9.002A1.006 1.006 0 0 1 8 14.99v-3.033c0-.558.449-1.01 1.002-1.01"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockOpen(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 10V8c0-2.761 1.239-5 4-5c2.094 0 3.313 1.288 3.78 3.114M3.5 17.8v-4.6c0-1.12 0-1.68.218-2.107a2 2 0 0 1 .874-.875c.428-.217.988-.217 2.108-.217h10.6c1.12 0 1.68 0 2.108.217a2 2 0 0 1 .874.874c.218.428.218.988.218 2.108v4.6c0 1.12 0 1.68-.218 2.108a2 2 0 0 1-.874.874C18.98 21 18.42 21 17.3 21H6.7c-1.12 0-1.68 0-2.108-.218a2 2 0 0 1-.874-.874C3.5 19.481 3.5 18.921 3.5 17.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockOpenKeyhole(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 10V8c0-2.761 1.239-5 4-5c2.094 0 3.313 1.288 3.78 3.114M3.5 17.8v-4.6c0-1.12 0-1.68.218-2.107a2 2 0 0 1 .874-.875c.428-.217.988-.217 2.108-.217h10.6c1.12 0 1.68 0 2.108.217a2 2 0 0 1 .874.874c.218.428.218.988.218 2.108v4.6c0 1.12 0 1.68-.218 2.108a2 2 0 0 1-.874.874C18.98 21 18.42 21 17.3 21H6.7c-1.12 0-1.68 0-2.108-.218a2 2 0 0 1-.874-.874C3.5 19.481 3.5 18.921 3.5 17.8M16 14v3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockOpenPassword(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 10V8c0-2.761 1.239-5 4-5c2.094 0 3.313 1.288 3.78 3.114M3.5 17.8v-4.6c0-1.12 0-1.68.218-2.107a2 2 0 0 1 .874-.875c.428-.217.988-.217 2.108-.217h10.6c1.12 0 1.68 0 2.108.217a2 2 0 0 1 .874.874c.218.428.218.988.218 2.108v4.6c0 1.12 0 1.68-.218 2.108a2 2 0 0 1-.874.874C18.98 21 18.42 21 17.3 21H6.7c-1.12 0-1.68 0-2.108-.218a2 2 0 0 1-.874-.874C3.5 19.481 3.5 18.921 3.5 17.8m8.5-2.05v-.5m4 .5v-.5m-8 .5v-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockPassword(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 10V8c0-2.761 1.239-5 4-5s4 2.239 4 5v2M3.5 17.8v-4.6c0-1.12 0-1.68.218-2.107a2 2 0 0 1 .874-.875c.428-.217.988-.217 2.108-.217h10.6c1.12 0 1.68 0 2.108.217a2 2 0 0 1 .874.874c.218.428.218.988.218 2.108v4.6c0 1.12 0 1.68-.218 2.108a2 2 0 0 1-.874.874C18.98 21 18.42 21 17.3 21H6.7c-1.12 0-1.68 0-2.108-.218a2 2 0 0 1-.874-.874C3.5 19.481 3.5 18.921 3.5 17.8m8.5-2.05v-.5m4 .5v-.5m-8 .5v-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/><path d="M14.004 10.947V8.925c0-2.641-4.008-2.491-4.008 0v2.021m-.994 0h5.996c.553 0 1.002.453 1.002 1.011v3.032c0 .558-.449 1.011-1.002 1.011H9.002A1.006 1.006 0 0 1 8 14.99v-3.033c0-.558.449-1.01 1.002-1.01"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/><path d="M14.004 10.947V8.925c0-2.641-4.008-2.491-4.008 0v2.021m-.994 0h5.996c.553 0 1.002.453 1.002 1.011v3.032c0 .558-.449 1.011-1.002 1.011H9.002A1.006 1.006 0 0 1 8 14.99v-3.033c0-.558.449-1.01 1.002-1.01"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Login(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13.496 21H6.5c-1.105 0-2-1.151-2-2.571V5.57c0-1.419.895-2.57 2-2.57h7"/><path d="M13 15.5L9.5 12L13 8.5m6.5 3.496h-10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Logout(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13.496 21H6.5c-1.105 0-2-1.151-2-2.571V5.57c0-1.419.895-2.57 2-2.57h7M16 15.5l3.5-3.5L16 8.5m-6.5 3.496h10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Magnet(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 7.5V12a9 9 0 1 0 18 0V7.5m-18 0V5a2 2 0 0 1 2-2h1.625a2 2 0 0 1 2 2v2.5M3 7.5h5.625m0 0V12a3.375 3.375 0 1 0 6.75 0V7.5m0 0V5a2 2 0 0 1 2-2H19a2 2 0 0 1 2 2v2.5m-5.625 0H21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Male(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.232 9.747a6 6 0 1 0-8.465 8.506a6 6 0 0 0 8.465-8.506m0 0L20 4m0 0h-4m4 0v4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Map(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 8.593c0-1.527 0-2.29.393-2.735c.139-.159.308-.285.497-.372c1.416-.653 3.272 1.066 4.77 1.013c.197-.007.394-.035.587-.082c2.184-.535 3.552-3.08 5.798-3.39c1.287-.18 2.7.598 3.904 1.014c.99.342 1.485.513 1.768.92C21 5.368 21 5.91 21 6.99v8.422c0 1.526 0 2.29-.393 2.735a1.493 1.493 0 0 1-.497.371c-1.416.653-3.272-1.065-4.77-1.012a2.903 2.903 0 0 0-.587.081c-2.184.535-3.552 3.08-5.798 3.391c-1.281.178-4.847-.75-5.672-1.935C3 18.636 3 18.096 3 17.014zm6-2.052v14.255m6-17.615v14.255"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mask(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m5.7 8.77l5.175-1.517a4 4 0 0 1 2.25 0L18.3 8.769m-12.6 0V7.35a1.35 1.35 0 0 0-2.7 0v2.188a2 2 0 0 0 2 2h.7m0-2.769v2.77m12.6-2.77v2.77m0-2.77V7.35a1.35 1.35 0 1 1 2.7 0v2.188a2 2 0 0 1-2 2h-.7m0 0v.162a6.3 6.3 0 1 1-12.6 0v-.161"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Math(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.429 18.366h6M15 8.707h6m-6-4.39h6M3 6.512h3.429m0 0h3.428m-3.428 0V3m0 3.512v3.512M15.6 21l2.425-2.484m0 0l2.424-2.483m-2.424 2.483L15.6 16.033m2.425 2.483L20.449 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MathSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12m4 3.625h3.5m3.25-5.5h3.5m-3.5-2.5h3.5m-10.5 1.25h2m0 0h2m-2 0v-2m0 2v2m5.35 6.25l1.414-1.414m0 0l1.415-1.414m-1.415 1.414L14.1 14.296m1.414 1.415l1.415 1.414"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Maximize(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.01 9.75c.04-2.79.247-4.371 1.308-5.432C5.379 3.258 6.96 3.05 9.75 3.01M3.01 14.26c.04 2.79.247 4.37 1.308 5.432C5.379 20.752 6.96 20.96 9.75 21M21 9.75c-.04-2.79-.247-4.371-1.308-5.432c-1.061-1.06-2.643-1.268-5.432-1.308M21 14.26c-.04 2.79-.247 4.37-1.308 5.432c-1.061 1.06-2.643 1.268-5.432 1.308"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MaximizeOne(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14 4h6v6M10 20H4v-6M20 4l-6 6M4 20l6-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Menu(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.5 6.5h15M4.5 12h15m-15 5.5h15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Message(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.464 16.828C2 15.657 2 14.771 2 11c0-3.771 0-5.657 1.464-6.828C4.93 3 7.286 3 12 3c4.714 0 7.071 0 8.535 1.172C22 5.343 22 7.229 22 11c0 3.771 0 4.657-1.465 5.828C19.072 18 16.714 18 12 18c-2.51 0-3.8 1.738-6 3v-3.212c-1.094-.163-1.899-.45-2.536-.96"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageCheck(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m9.7 10.823l1.379 1.575a.299.299 0 0 0 .466-.022l2.8-3.876"/><path d="M3.464 16.828C2 15.657 2 14.771 2 11c0-3.771 0-5.657 1.464-6.828C4.93 3 7.286 3 12 3c4.714 0 7.071 0 8.535 1.172C22 5.343 22 7.229 22 11c0 3.771 0 4.657-1.465 5.828C19.072 18 16.714 18 12 18c-2.51 0-3.8 1.738-6 3v-3.212c-1.094-.163-1.899-.45-2.536-.96"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageDots(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 11v-.5m4 .5v-.5M8 11v-.5m-4.536 6.328C2 15.657 2 14.771 2 11c0-3.771 0-5.657 1.464-6.828C4.93 3 7.286 3 12 3c4.714 0 7.071 0 8.535 1.172C22 5.343 22 7.229 22 11c0 3.771 0 4.657-1.465 5.828C19.072 18 16.714 18 12 18c-2.51 0-3.8 1.738-6 3v-3.212c-1.094-.163-1.899-.45-2.536-.96"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageMinus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.464 16.828C2 15.657 2 14.771 2 11c0-3.771 0-5.657 1.464-6.828C4.93 3 7.286 3 12 3c4.714 0 7.071 0 8.535 1.172C22 5.343 22 7.229 22 11c0 3.771 0 4.657-1.465 5.828C19.072 18 16.714 18 12 18c-2.51 0-3.8 1.738-6 3v-3.212c-1.094-.163-1.899-.45-2.536-.96M9.5 10.5h5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessagePlus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.5 10.5h5M12 8v5m-8.536 3.828C2 15.657 2 14.771 2 11c0-3.771 0-5.657 1.464-6.828C4.93 3 7.286 3 12 3c4.714 0 7.071 0 8.535 1.172C22 5.343 22 7.229 22 11c0 3.771 0 4.657-1.465 5.828C19.072 18 16.714 18 12 18c-2.51 0-3.8 1.738-6 3v-3.212c-1.094-.163-1.899-.45-2.536-.96"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageX(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m10 8.5l4 4m0-4l-4 4m-6.536 4.328C2 15.657 2 14.771 2 11c0-3.771 0-5.657 1.464-6.828C4.93 3 7.286 3 12 3c4.714 0 7.071 0 8.535 1.172C22 5.343 22 7.229 22 11c0 3.771 0 4.657-1.465 5.828C19.072 18 16.714 18 12 18c-2.51 0-3.8 1.738-6 3v-3.212c-1.094-.163-1.899-.45-2.536-.96"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Microphone(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16 6.429C16 4.535 14.21 3 12 3S8 4.535 8 6.429v5.142C8 13.465 9.79 15 12 15s4-1.535 4-3.429z"/><path d="M5 11a7 7 0 1 0 14 0m-7 7v3m-4 0h8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicrophoneSlash(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 18v3m-4 0h8m-8-9.429V6.43C8 4.535 9.79 3 12 3c1.309 0 2.47.539 3.2 1.371M5 11c0 1.046.23 2.039.641 2.93M19 11a7 7 0 0 1-11.425 5.425M3 21l4.575-4.575M21 3l-5 5m0 0v3.571C16 13.465 14.21 15 12 15a4.444 4.444 0 0 1-2.348-.652M16 8l-6.348 6.348m0 0l-2.077 2.077"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Minimize(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.75 3.01c-.04 2.79-.247 4.37-1.308 5.432C7.38 9.502 5.799 9.71 3.01 9.75M9.75 21c-.04-2.79-.247-4.371-1.308-5.432S5.799 14.3 3.01 14.26M14.26 3.01c.04 2.79.247 4.37 1.308 5.432C16.629 9.502 18.211 9.71 21 9.75M14.26 21c.04-2.79.247-4.371 1.308-5.432S18.211 14.3 21 14.26"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinimizeOne(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 10h-6V4M4 14h6v6M20 4l-6 6m-4 4l-6 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Minus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 12h12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8.5 12h7"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.5 12h7m5 3.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132zM8.5 12h7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12m5.5 0h7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.5 12h7M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mobile(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16.286 3H7.714C6.768 3 6 3.806 6 4.8v14.4c0 .994.768 1.8 1.714 1.8h8.572c.947 0 1.714-.806 1.714-1.8V4.8c0-.994-.767-1.8-1.714-1.8M10.5 6h3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MobileSignalFive(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 21v-1m18 1V3m-9 18v-9m4.5 9V8m-9 13v-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MobileSignalFour(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 21v-1m9 1v-9m4.5 9V8m-9 13v-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MobileSignalOne(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 21v-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MobileSignalThree(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 21v-1m9 1v-9m-4.5 9v-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MobileSignalTwo(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 21v-1m4.5 1v-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Moon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 12.808c-.5 5.347-5.849 9.14-11.107 7.983C-.077 18.6 1.15 3.909 11.112 3C6.394 9.296 14.618 17.462 21 12.808"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mountain(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m13.006 15.265l2.117-3.14m0 0l.251-.37a1.637 1.637 0 0 1 2.861.234l2.596 5.541c.536 1.142-.27 2.47-1.497 2.47H4.666c-1.224 0-2.03-1.32-1.501-2.462l5.808-12.56a1.641 1.641 0 0 1 3.015.03z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MousePointer(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m12.866 12.866l5.628-2.412c.942-.404.886-1.758-.086-2.082L5.469 4.059c-.871-.29-1.7.539-1.41 1.41l4.313 12.939c.324.972 1.678 1.028 2.082.086zm0 0L20 20"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Move(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.7 9.3L3 12m0 0l2.7 2.7M3 12h18M9.3 5.7L12 3m0 0l2.7 2.7M12 3v18m2.7-2.7L12 21m0 0l-2.7-2.7m9-9L21 12m0 0l-2.7 2.7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoveDiagonal(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13 5h6m0 0v6m0-6L5 19m6 0H5m0 0v-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoveDiagonalOne(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 11V5m0 0h6M5 5l14 14m0-6v6m0 0h-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoveHorizontal(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m18 8l4 4m0 0l-4 4m4-4H2m4-4l-4 4m0 0l4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoveVertical(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8 18l4 4m0 0l4-4m-4 4V2M8 6l4-4m0 0l4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Music(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.625 17.65c0 1.574-1.26 2.85-2.812 2.85C4.259 20.5 3 19.224 3 17.65c0-1.573 1.26-2.849 2.813-2.849c1.553 0 2.812 1.276 2.812 2.85m0 0V5.462c0-.52.394-.954.909-1.001l10.375-.956A1.002 1.002 0 0 1 21 4.506V16.51m0 0c0 1.573-1.26 2.85-2.812 2.85c-1.554 0-2.813-1.277-2.813-2.85c0-1.574 1.26-2.85 2.813-2.85c1.553 0 2.812 1.276 2.812 2.85"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M12.5 14.5V8.6a.6.6 0 0 1 .6-.6h1.4m-2 6.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/><path d="M12.5 14.5V8.6a.6.6 0 0 1 .6-.6h1.4m-2 6.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/><path d="M12.5 14.5V8.6a.6.6 0 0 1 .6-.6h1.4m-2 6.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/><path d="M12.5 14.5V8.6a.6.6 0 0 1 .6-.6h1.4m-2 6.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/><path d="M12.5 14.5V8.6a.6.6 0 0 1 .6-.6h1.4m-2 6.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Myna(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m3.25 10.347l.705-.66l1.41-1.318c2.22-2.292 5.803-2.292 8.004 0l4.228 3.952a.288.288 0 0 1 0 .42l-8.67 8.105"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Navigation(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.54 3.333a.485.485 0 0 1 .92 0l6.5 16.916c.178.464-.272.917-.685.69l-6.05-3.314a.464.464 0 0 0-.45 0l-6.05 3.315c-.413.226-.863-.227-.685-.69z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NavigationOne(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.414 12.086a.546.546 0 0 1-.101-1.024l16.905-8.007a.546.546 0 0 1 .727.727l-8.007 16.905a.546.546 0 0 1-1.024-.1l-1.62-6.483a.546.546 0 0 0-.398-.397z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Nine(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 17h1.25A3.75 3.75 0 0 0 15 13.25v-3.125m0 0a3.125 3.125 0 1 0-6.25 0c0 1.726 1.4 2.5 3.125 2.5c1.726 0 3.125-.774 3.125-2.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NineCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M10.5 16h1a3 3 0 0 0 3-3v-2.5m0 0a2.5 2.5 0 0 0-5 0c0 1.38 1.12 2 2.5 2s2.5-.62 2.5-2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NineDiamond(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.5 16h1a3 3 0 0 0 3-3v-2.5m0 0a2.5 2.5 0 0 0-5 0c0 1.38 1.12 2 2.5 2s2.5-.62 2.5-2"/><path d="M2.707 10.295a2.41 2.41 0 0 0 0 3.41l7.588 7.588a2.409 2.409 0 0 0 3.41 0l7.588-7.588a2.409 2.409 0 0 0 0-3.41l-7.588-7.588a2.41 2.41 0 0 0-3.41 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NineHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.5 16h1a3 3 0 0 0 3-3v-2.5m0 0a2.5 2.5 0 0 0-5 0c0 1.38 1.12 2 2.5 2s2.5-.62 2.5-2"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NineOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.5 16h1a3 3 0 0 0 3-3v-2.5m0 0a2.5 2.5 0 0 0-5 0c0 1.38 1.12 2 2.5 2s2.5-.62 2.5-2"/><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NineSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.5 16h1a3 3 0 0 0 3-3v-2.5m0 0a2.5 2.5 0 0 0-5 0c0 1.38 1.12 2 2.5 2s2.5-.62 2.5-2"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NineWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.5 16h1a3 3 0 0 0 3-3v-2.5m0 0a2.5 2.5 0 0 0-5 0c0 1.38 1.12 2 2.5 2s2.5-.62 2.5-2"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Notification(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3"/><circle cx="18.25" cy="5.75" r="2.75"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Octagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OctagonCheck(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132L7.805 3.47Z"/><path d="m8.667 12.833l1.505 1.721a1 1 0 0 0 1.564-.073L15.333 9.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OctagonDanger(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132L7.805 3.47ZM12 7.627v5.5m0 3.246v-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OctagonInfo(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132L7.805 3.47Z"/><path d="M12 16v-5h-.5m0 5h1M12 8.5V8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OctagonMinus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132L7.805 3.47ZM8.5 12h7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OctagonPlus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132L7.805 3.47ZM15.5 12H12m0 0H8.5m3.5 0V8.5m0 3.5v3.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OctagonSlash(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132L7.805 3.47ZM9 15l6-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OctagonX(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132L7.805 3.47ZM15 9l-6 6m0-6l6 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func One(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.438 17V7L9 9.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OneCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M12.75 16V8L10 10"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OneDiamond(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12.75 16V8L10 10"/><path d="M2.707 10.295a2.41 2.41 0 0 0 0 3.41l7.588 7.588a2.409 2.409 0 0 0 3.41 0l7.588-7.588a2.409 2.409 0 0 0 0-3.41l-7.588-7.588a2.41 2.41 0 0 0-3.41 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OneHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12.75 16V8L10 10"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OneOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12.75 16V8L10 10"/><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OneSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12.75 16V8L10 10"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OneWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12.75 16V8L10 10"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Option(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 7h5.094c.33 0 .495 0 .643.047c.132.042.253.111.357.202c.117.103.202.245.372.528l5.068 8.446c.17.284.255.425.372.528c.103.09.224.16.356.202c.148.047.314.047.644.047H21M15 7h6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Package(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m12 21l8.131-4.208c.316-.164.474-.245.589-.366a1 1 0 0 0 .226-.373c.054-.159.054-.336.054-.692V7.533M12 21l-8.131-4.208c-.316-.164-.474-.245-.589-.366a1.009 1.009 0 0 1-.226-.373C3 15.894 3 15.716 3 15.359V7.533M12 21v-9.063m9-4.404l-9 4.404m9-4.404l-4.5-2.33M3 7.534l8.269-4.28c.268-.138.401-.208.542-.235a.994.994 0 0 1 .378 0c.14.027.274.097.541.235l3.77 1.95M3 7.534l4.5 2.202m4.5 2.202L7.5 9.735m0 0l9-4.531m-9 4.531v2.202"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paint(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16.5 5.5h1c1.105 0 2 .395 2 1.5v2a2 2 0 0 1-2 2h-7a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h1m5-15.5V4a1 1 0 0 0-1-1h-10a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PanelBottom(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.5 15h17M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PanelBottomClose(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3.5 15h17M15 8l-3 3l-3-3"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PanelBottomInactive(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14 15h1m4 0h1.5m-17 0H5m4 0h1m-7-3c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PanelBottomOpen(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3.5 15h17M9 10l3-3l3 3"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PanelLeft(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 3.5v17M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PanelLeftClose(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 3.5v17m7-5.5l-3-3l3-3"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PanelLeftInactive(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 14v1m0 4v1.5m0-17V5m0 4v1m-6 2c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PanelLeftOpen(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 3.5v17M14 9l3 3l-3 3"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PanelRight(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15 3.5v17M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PanelRightClose(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15 3.5v17M8 9l3 3l-3 3"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PanelRightInactive(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15 14v1m0 4v1.5m0-17V5m0 4v1M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PanelRightOpen(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15 3.5v17M10 15l-3-3l3-3"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PanelTop(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.5 9h17M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PanelTopClose(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3.5 9h17M9 16l3-3l3 3"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PanelTopInactive(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14 9h1m4 0h1.5m-17 0H5m4 0h1m-7 3c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PanelTopOpen(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3.5 9h17M15 14l-3 3l-3-3"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paperclip(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3.845 10.224l5.739-5.53c5.429-5.237 14.246 2.843 8.596 8.294l-7.112 6.862c-3.684 3.554-9.667-1.929-5.833-5.628l7.01-6.763c1.939-1.87 5.087 1.015 3.07 2.962L9.492 16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Parking(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.5 15v-2.4m0 0h2.276c2.299 0 2.299-3.6 0-3.6H10.5z"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Password(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m13.5 9l-3 6m0-6l3 6m-3.75-3h4.5m-7.5-3l-3 6m0-6l3 6M3 12h4.5m12.75-3l-3 6m0-6l3 6m-3.75-3H21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Path(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.121 15.879a3 3 0 1 0-4.243 4.243a3 3 0 0 0 4.243-4.243m0 0L15.88 8.12m0 0a3 3 0 1 0 4.243-4.243A3 3 0 0 0 15.88 8.12m0 0l.004-.004"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pause(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 6.5H8a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-10a1 1 0 0 0-1-1m6.5 0h-1a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-10a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10 9v6m4-6v6"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 9v6m4-6v6m6.5.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132zM10 9v6m4-6v6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 9v6m4-6v6M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 9v6m4-6v6M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pen(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m7.63 7.63l3.006 3.006M3.042 4.34l2.616 9.157c.778 2.723 2.933 4.6 5.764 4.6a5.61 5.61 0 0 0 1.264-.148c.392-.09.812-.001 1.097.283l2.46 2.46c.41.41 1.075.41 1.485 0l2.964-2.965a1.05 1.05 0 0 0 0-1.486l-2.455-2.455c-.284-.284-.373-.703-.284-1.094c.094-.415.145-.84.145-1.27c0-2.832-1.877-4.987-4.6-5.765L4.341 3.042a1.05 1.05 0 0 0-1.3 1.299"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pencil(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.8 20.199A2.733 2.733 0 0 1 6.869 21H3v-3.844c0-.724.288-1.419.8-1.931m5 4.974l-5-4.974m5 4.974l9.974-9.978M3.8 15.225l9.984-9.995m0 0l1.426-1.428a2.733 2.733 0 0 1 3.867-.001l1.126 1.127a2.733 2.733 0 0 1 0 3.865l-1.428 1.428M13.783 5.23l4.991 4.991"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Percentage(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 5L5 19m12.5 1c1.667 0 2.5-.857 2.5-3s-.833-3-2.5-3s-2.5.857-2.5 3s.833 3 2.5 3m-11-10C8.167 10 9 9.143 9 7s-.833-3-2.5-3S4 4.857 4 7s.833 3 2.5 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PercentageCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.854 9.854L9.5 9.5m5.004 5.004l-.354-.354m-4.65.35l5-5"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PercentageHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.854 9.854L9.5 9.5m5.004 5.004l-.354-.354m-4.65.35l5-5m6 6.3V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PercentageOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132zm2.049 6.385L9.5 9.5m5.004 5.004l-.354-.354m-4.65.35l5-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PercentageSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.854 9.854L9.5 9.5m5.004 5.004l-.354-.354m-4.65.35l5-5M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PercentageWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.854 9.854L9.5 9.5m5.004 5.004l-.354-.354m-4.65.35l5-5M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pin(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3 21l4.63-4.631m.005-.005l-2.78-2.78c-.954-.953.006-2.996 1.31-3.078c1.178-.075 3.905.352 4.812-.555l2.49-2.49c.617-.618.225-2 .185-2.762c-.058-1.016 1.558-2.271 2.415-1.414l4.647 4.648c.86.858-.4 2.469-1.413 2.415c-.762-.04-2.145-.432-2.763.185l-2.49 2.49c-.906.907-.48 3.633-.554 4.811c-.082 1.305-2.125 2.265-3.08 1.31z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pizza(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.234 16.925a15.98 15.98 0 0 1 10.69-10.691M7.582 7.58a18.756 18.756 0 0 1 7.33-4.53c.536-.18 1.103.136 1.265.678l4.779 15.928a1.042 1.042 0 0 1-1.298 1.298L3.73 16.176c-.542-.162-.858-.729-.679-1.266a18.756 18.756 0 0 1 4.53-7.33M11 15l.354.354M15 11l.354.354M16 16l.354.354"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Planet(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.722 17.777a7.303 7.303 0 0 0 4.716 1.723c4.085 0 7.396-3.358 7.396-7.5a7.59 7.59 0 0 0-.16-1.556M7.722 17.777A7.527 7.527 0 0 1 5.042 12c0-4.142 3.311-7.5 7.395-7.5c3.559 0 6.53 2.549 7.236 5.944M7.722 17.777c1.807-.42 3.958-1.293 6.127-2.563c2.524-1.478 4.577-3.202 5.825-4.77M7.722 17.777c-2.246.52-3.963.34-4.528-.654c-.583-1.024.182-2.688 1.849-4.458m14.631-2.22c1.157-1.454 1.623-2.772 1.132-3.635c-.498-.875-1.888-1.12-3.746-.81"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Play(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.581 9.402C16.194 10.718 17 11.375 17 12.5c0 1.125-.806 1.783-2.419 3.098a23.1 23.1 0 0 1-1.292.99c-.356.25-.759.508-1.176.762c-1.609.978-2.413 1.467-3.134.926c-.722-.542-.787-1.675-.918-3.943A32.48 32.48 0 0 1 8 12.5c0-.563.023-1.192.06-1.833c.132-2.267.197-3.401.919-3.943c.721-.541 1.525-.052 3.134.926c.417.254.82.512 1.176.762a23.1 23.1 0 0 1 1.292.99"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13.242 10.71c.672.547 1.008.821 1.008 1.29c0 .469-.336.743-1.008 1.29c-.185.152-.37.295-.538.413a9.093 9.093 0 0 1-.49.318c-.67.407-1.006.611-1.306.385c-.3-.225-.328-.697-.383-1.642A13.577 13.577 0 0 1 10.5 12c0-.235.01-.497.025-.764c.055-.945.082-1.417.383-1.642c.3-.226.636-.022 1.306.385c.174.106.341.214.49.318c.169.118.353.261.538.412"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13.242 10.71c.672.547 1.008.821 1.008 1.29c0 .469-.336.743-1.008 1.29c-.185.152-.37.295-.538.413a9.093 9.093 0 0 1-.49.318c-.67.407-1.006.611-1.306.385c-.3-.225-.328-.697-.383-1.642A13.577 13.577 0 0 1 10.5 12c0-.235.01-.497.025-.764c.055-.945.082-1.417.383-1.642c.3-.226.636-.022 1.306.385c.174.106.341.214.49.318c.169.118.353.261.538.412"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/><path d="M13.242 10.71c.672.547 1.008.821 1.008 1.29c0 .469-.336.743-1.008 1.29c-.185.152-.37.295-.538.413a9.093 9.093 0 0 1-.49.318c-.67.407-1.006.611-1.306.385c-.3-.225-.328-.697-.383-1.642A13.577 13.577 0 0 1 10.5 12c0-.235.01-.497.025-.764c.055-.945.082-1.417.383-1.642c.3-.226.636-.022 1.306.385c.174.106.341.214.49.318c.169.118.353.261.538.412"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlaySquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13.242 10.71c.672.547 1.008.821 1.008 1.29c0 .469-.336.743-1.008 1.29c-.185.152-.37.295-.538.413a9.093 9.093 0 0 1-.49.318c-.67.407-1.006.611-1.306.385c-.3-.225-.328-.697-.383-1.642A13.577 13.577 0 0 1 10.5 12c0-.235.01-.497.025-.764c.055-.945.082-1.417.383-1.642c.3-.226.636-.022 1.306.385c.174.106.341.214.49.318c.169.118.353.261.538.412"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13.242 10.71c.672.547 1.008.821 1.008 1.29c0 .469-.336.743-1.008 1.29c-.185.152-.37.295-.538.413a9.093 9.093 0 0 1-.49.318c-.67.407-1.006.611-1.306.385c-.3-.225-.328-.697-.383-1.642A13.577 13.577 0 0 1 10.5 12c0-.235.01-.497.025-.764c.055-.945.082-1.417.383-1.642c.3-.226.636-.022 1.306.385c.174.106.341.214.49.318c.169.118.353.261.538.412"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 12h-6m0 0H6m6 0V6m0 6v6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15.5 12H12m0 0H8.5m3.5 0V8.5m0 3.5v3.5"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.5 12H12m0 0H8.5m3.5 0V8.5m0 3.5v3.5m8.5.3V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132zM15.5 12H12m0 0H8.5m3.5 0V8.5m0 3.5v3.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.5 12H12m0 0H8.5m3.5 0V8.5m0 3.5v3.5M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.5 12H12m0 0H8.5m3.5 0V8.5m0 3.5v3.5M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pokeball(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4m2-2h7M3 12h7"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Power(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17.953 5.25a9 9 0 1 1-11.906 0M12 3v9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Presentation(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 4h18m-1.9 0v9.778c0 .471-.19.923-.527 1.257c-.338.333-.796.52-1.273.52H6.7c-.477 0-.935-.187-1.273-.52a1.767 1.767 0 0 1-.527-1.257V4m2.6 16l4.5-4.444L16.5 20"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Printer(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19 10V5a1 1 0 0 0-1-1H6a1 1 0 0 0-1 1v5m15 0H4a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1v-8a1 1 0 0 0-1-1"/><path d="M17.5 20v-3a1 1 0 0 0-1-1H11a1 1 0 0 0-1 1v3m-4-7h2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Puzzle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.4 12c.56 0 .84 0 1.054-.109a1 1 0 0 0 .437-.437C21 11.24 21 10.96 21 10.4V8.2c0-.56 0-.84-.109-1.054a1 1 0 0 0-.437-.437C20.24 6.6 19.96 6.6 19.4 6.6h-1.55a.9.9 0 0 1-.9-.9a2.7 2.7 0 1 0-5.4 0a.9.9 0 0 1-.9.9H9.1c-.56 0-.84 0-1.054.109a1 1 0 0 0-.437.437C7.5 7.36 7.5 7.64 7.5 8.2v2.2c0 .56 0 .84-.109 1.054a1 1 0 0 1-.437.437C6.74 12 6.46 12 5.9 12h-.2a2.7 2.7 0 1 0 0 5.4h.2c.56 0 .84 0 1.054.109a1 1 0 0 1 .437.437c.109.214.109.494.109 1.054v.4c0 .56 0 .84.109 1.054a1 1 0 0 0 .437.437C8.26 21 8.54 21 9.1 21h10.3c.56 0 .84 0 1.054-.109a1 1 0 0 0 .437-.437C21 20.24 21 19.96 21 19.4V19c0-.56 0-.84-.109-1.054a1 1 0 0 0-.437-.437c-.214-.109-.494-.109-1.054-.109h-.2a2.7 2.7 0 1 1 0-5.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Question(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 21v-.5m0-3c0-5.1 5-3.825 5-8.924c0-6.768-10-6.768-10 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 13.496c0-2.003 2-1.503 2-3.506c0-2.659-4-2.659-4 0m2 6.007v-.5"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 13.496c0-2.003 2-1.503 2-3.506c0-2.659-4-2.659-4 0m2 6.007v-.5"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/><path d="M12 13.496c0-2.003 2-1.503 2-3.506c0-2.659-4-2.659-4 0m2 6.007v-.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 13.496c0-2.003 2-1.503 2-3.506c0-2.659-4-2.659-4 0m2 6.007v-.5"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 13.496c0-2.003 2-1.503 2-3.506c0-2.659-4-2.659-4 0m2 6.007v-.5"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Radio(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 13.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m-3.449 2.205a5.234 5.234 0 0 1-1.148-1.7a5.192 5.192 0 0 1 1.148-5.71M18.07 5.5a9.148 9.148 0 0 1 2.719 6.5a9.148 9.148 0 0 1-2.72 6.5M15.24 8.295a5.232 5.232 0 0 1 1.148 1.7a5.191 5.191 0 0 1-1.148 5.71M5.512 18.5A9.148 9.148 0 0 1 2.793 12c0-2.438.978-4.776 2.72-6.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rainbow(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 17.5v-2a9 9 0 1 1 18 0v2m-15 0v-2a6 6 0 0 1 12 0v2m-9 0v-2a3 3 0 1 1 6 0v2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReceptionBell(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 19h18M12 8V5m0 3h-2a5 5 0 0 0-5 5v3h14v-3a5 5 0 0 0-5-5zm0-3h-2m2 0h2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Record(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 12.5a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0m-11 0a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0M6.5 16h11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rectangle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2 12c0-3.771 0-5.657 1.464-6.828C4.93 4 7.286 4 12 4c4.714 0 7.071 0 8.535 1.172C22 6.343 22 8.229 22 12c0 3.771 0 5.657-1.465 6.828C19.072 20 16.714 20 12 20s-7.071 0-8.536-1.172C2 17.657 2 15.771 2 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RectangleVertical(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 2c3.771 0 5.657 0 6.828 1.464C20 4.93 20 7.286 20 12c0 4.714 0 7.071-1.172 8.535C17.657 22 15.771 22 12 22c-3.771 0-5.657 0-6.828-1.465C4 19.072 4 16.714 4 12s0-7.071 1.172-8.536C6.343 2 8.229 2 12 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Redo(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18.364 18.364A9 9 0 1 1 12 3c4.058 0 6.518 2.705 9 5.5"/><path d="M21 4.5v4h-4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Refresh(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20.5 8c-1.392-3.179-4.823-5-8.522-5C7.299 3 3.453 6.552 3 11.1"/><path d="M16.489 8.4h3.97a.54.54 0 0 0 .54-.54V3.9M3.5 16c1.392 3.179 4.823 5 8.522 5c4.679 0 8.525-3.552 8.978-8.1"/><path d="M7.511 15.6h-3.97a.54.54 0 0 0-.541.54v3.96"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RefreshAlt(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3.5 8c1.392-3.179 4.823-5 8.522-5c4.679 0 8.525 3.552 8.978 8.1"/><path d="M7.511 8.4h-3.97a.54.54 0 0 1-.54-.54V3.9M20.5 16c-1.392 3.179-4.822 5-8.522 5C7.299 21 3.453 17.448 3 12.9"/><path d="M16.489 15.6h3.97a.54.54 0 0 1 .541.54v3.96"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Repeat(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m17 3l4 3l-4 3"/><path d="M3 12v-2a4 4 0 0 1 4-4h14M7 21l-4-3l4-3"/><path d="M21 12v2a4 4 0 0 1-4 4H3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rewind(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6.515 9.419C5.172 10.515 4.5 11.063 4.5 12s.672 1.485 2.015 2.582c.371.302.74.587 1.077.824c.297.209.633.424.98.635c1.341.816 2.011 1.223 2.613.772c.6-.451.655-1.396.765-3.285c.03-.535.05-1.058.05-1.528s-.02-.993-.05-1.528c-.11-1.89-.164-2.834-.765-3.285c-.602-.451-1.272-.044-2.612.771a17.63 17.63 0 0 0-.98.636c-.339.237-.707.522-1.078.825"/><path d="M14.016 9.419C12.672 10.515 12 11.063 12 12s.672 1.485 2.015 2.582c.371.302.74.587 1.077.824c.297.209.633.424.98.635c1.341.816 2.011 1.223 2.613.772c.6-.451.655-1.396.765-3.285c.03-.535.05-1.058.05-1.528s-.02-.993-.05-1.528c-.11-1.89-.164-2.834-.765-3.285c-.602-.451-1.272-.044-2.612.771a17.63 17.63 0 0 0-.98.636a19.26 19.26 0 0 0-1.078.825"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RewindCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8.008 10.71C7.336 11.256 7 11.53 7 12c0 .469.336.743 1.008 1.29c.185.152.37.295.538.413c.149.104.316.212.49.318c.67.407 1.006.611 1.306.385c.3-.225.328-.697.383-1.642c.015-.267.025-.53.025-.764c0-.235-.01-.497-.025-.764c-.055-.945-.082-1.417-.383-1.642c-.3-.226-.635-.022-1.306.385a9.078 9.078 0 0 0-.49.318a9.647 9.647 0 0 0-.538.412m5.75.001c-.672.547-1.008.821-1.008 1.29c0 .469.336.743 1.008 1.29c.185.152.37.295.538.413c.149.104.316.212.49.318c.67.407 1.006.611 1.306.385c.3-.225.328-.697.383-1.642c.015-.267.025-.53.025-.764c0-.235-.01-.497-.025-.764c-.055-.945-.082-1.417-.383-1.642c-.3-.226-.635-.022-1.306.385a9.07 9.07 0 0 0-.49.318a9.632 9.632 0 0 0-.538.412"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RewindHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8.008 10.71C7.336 11.256 7 11.53 7 12c0 .469.336.743 1.008 1.29c.185.152.37.295.538.413c.149.104.316.212.49.318c.67.407 1.006.611 1.306.385c.3-.225.328-.697.383-1.642c.015-.267.025-.53.025-.764c0-.235-.01-.497-.025-.764c-.055-.945-.082-1.417-.383-1.642c-.3-.226-.635-.022-1.306.385a9.078 9.078 0 0 0-.49.318a9.647 9.647 0 0 0-.538.412m5.75.001c-.672.547-1.008.821-1.008 1.29c0 .469.336.743 1.008 1.29c.185.152.37.295.538.413c.149.104.316.212.49.318c.67.407 1.006.611 1.306.385c.3-.225.328-.697.383-1.642c.015-.267.025-.53.025-.764c0-.235-.01-.497-.025-.764c-.055-.945-.082-1.417-.383-1.642c-.3-.226-.635-.022-1.306.385a9.07 9.07 0 0 0-.49.318a9.632 9.632 0 0 0-.538.412"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RewindOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/><path d="M8.008 10.71C7.336 11.256 7 11.53 7 12c0 .469.336.743 1.008 1.29c.185.152.37.295.538.413c.149.104.316.212.49.318c.67.407 1.006.611 1.306.385c.3-.225.328-.697.383-1.642c.015-.267.025-.53.025-.764c0-.235-.01-.497-.025-.764c-.055-.945-.082-1.417-.383-1.642c-.3-.226-.635-.022-1.306.385a9.078 9.078 0 0 0-.49.318a9.647 9.647 0 0 0-.538.412m5.75.001c-.672.547-1.008.821-1.008 1.29c0 .469.336.743 1.008 1.29c.185.152.37.295.538.413c.149.104.316.212.49.318c.67.407 1.006.611 1.306.385c.3-.225.328-.697.383-1.642c.015-.267.025-.53.025-.764c0-.235-.01-.497-.025-.764c-.055-.945-.082-1.417-.383-1.642c-.3-.226-.635-.022-1.306.385a9.07 9.07 0 0 0-.49.318a9.632 9.632 0 0 0-.538.412"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RewindSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8.008 10.71C7.336 11.256 7 11.53 7 12c0 .469.336.743 1.008 1.29c.185.152.37.295.538.413c.149.104.316.212.49.318c.67.407 1.006.611 1.306.385c.3-.225.328-.697.383-1.642c.015-.267.025-.53.025-.764c0-.235-.01-.497-.025-.764c-.055-.945-.082-1.417-.383-1.642c-.3-.226-.635-.022-1.306.385a9.078 9.078 0 0 0-.49.318a9.647 9.647 0 0 0-.538.412m5.75.001c-.672.547-1.008.821-1.008 1.29c0 .469.336.743 1.008 1.29c.185.152.37.295.538.413c.149.104.316.212.49.318c.67.407 1.006.611 1.306.385c.3-.225.328-.697.383-1.642c.015-.267.025-.53.025-.764c0-.235-.01-.497-.025-.764c-.055-.945-.082-1.417-.383-1.642c-.3-.226-.635-.022-1.306.385a9.07 9.07 0 0 0-.49.318a9.632 9.632 0 0 0-.538.412"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RewindWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8.008 10.71C7.336 11.256 7 11.53 7 12c0 .469.336.743 1.008 1.29c.185.152.37.295.538.413c.149.104.316.212.49.318c.67.407 1.006.611 1.306.385c.3-.225.328-.697.383-1.642c.015-.267.025-.53.025-.764c0-.235-.01-.497-.025-.764c-.055-.945-.082-1.417-.383-1.642c-.3-.226-.635-.022-1.306.385a9.078 9.078 0 0 0-.49.318a9.647 9.647 0 0 0-.538.412m5.75.001c-.672.547-1.008.821-1.008 1.29c0 .469.336.743 1.008 1.29c.185.152.37.295.538.413c.149.104.316.212.49.318c.67.407 1.006.611 1.306.385c.3-.225.328-.697.383-1.642c.015-.267.025-.53.025-.764c0-.235-.01-.497-.025-.764c-.055-.945-.082-1.417-.383-1.642c-.3-.226-.635-.022-1.306.385a9.07 9.07 0 0 0-.49.318a9.632 9.632 0 0 0-.538.412"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rhombus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m5.219 11.34l5.96-7.925a1.019 1.019 0 0 1 1.642 0l5.96 7.925c.292.388.292.932 0 1.32l-5.96 7.925a1.019 1.019 0 0 1-1.642 0L5.22 12.66a1.104 1.104 0 0 1 0-1.32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ribbon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18 9A6 6 0 1 1 6 9a6 6 0 0 1 12 0"/><path d="m8 13.472l-1 6.44c0 .81 1.782 1.336 2.447.974l2.106-1.147a.927.927 0 0 1 .894 0l2.106 1.147c.665.362 2.447-.165 2.447-.975l-1-6.439"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RoomService(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 6.167V4m0 2.167a8 8 0 0 1 8 8V17H4v-2.833a8 8 0 0 1 8-8M3 20h18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rows(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 6.5a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1zm0 8a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rss(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 11.25A6.75 6.75 0 0 1 12.75 18M6 6a12 12 0 0 1 12 12m-11.5-.146l.354-.354"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ruler(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m5.914 12l2.282 2.283m.76-5.326l2.283 2.282M12 5.914l2.283 2.282m-9.74 5.174l8.826-8.826c.853-.852 1.28-1.278 1.77-1.438c.433-.14.898-.14 1.33 0c.491.16.917.586 1.768 1.437l1.22 1.219c.852.852 1.278 1.28 1.438 1.77c.14.433.14.897 0 1.33c-.16.49-.586.917-1.438 1.77l-8.826 8.826c-.853.852-1.28 1.278-1.77 1.438c-.433.14-.897.14-1.33 0c-.49-.16-.918-.586-1.77-1.439l-1.22-1.218c-.85-.852-1.276-1.277-1.436-1.769a2.152 2.152 0 0 1 0-1.33c.16-.49.586-.917 1.438-1.77"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rupee(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.5 3.5h11m-11 4.722h11M14.292 20.5L6.5 12.944h2.75c6.111 0 6.111-9.444 0-9.444"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RupeeCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M9 7.5h6m-6 2.778h6M13.25 17.5L9 13.056h1.5c3.334 0 3.334-5.556 0-5.556"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RupeeHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/><path d="M9 7.5h6m-6 2.778h6M13.25 17.5L9 13.056h1.5c3.334 0 3.334-5.556 0-5.556"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RupeeOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/><path d="M9 7.5h6m-6 2.778h6M13.25 17.5L9 13.056h1.5c3.334 0 3.334-5.556 0-5.556"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RupeeSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/><path d="M9 7.5h6m-6 2.778h6M13.25 17.5L9 13.056h1.5c3.334 0 3.334-5.556 0-5.556"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RupeeWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/><path d="M9 7.5h6m-6 2.778h6M13.25 17.5L9 13.056h1.5c3.334 0 3.334-5.556 0-5.556"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SadCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 16c.85-.63 1.885-1 3-1s2.15.37 3 1m-5.5-5.5V10m5 .5V10"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SadGhost(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 16c.85-.63 1.885-1 3-1s2.15.37 3 1m-5.5-5.5V10m5 .5V10"/><path d="M3 18.562v-6.518C3 7.05 7.03 3 12 3s9 4.05 9 9.044v6.517c0 1.162-.967 2.519-2 2c-.835-.42-2.223-.52-3 0c-.874.585-2.126.585-3 0c-.885-.593-1.649-.57-2.5 0c-.874.585-2.126.585-3 0c-.777-.52-1.665-.42-2.5 0c-1.033.519-2-.838-2-2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SadSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 16c.85-.63 1.885-1 3-1s2.15.37 3 1m-5.5-5.5V10m5 .5V10"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Save(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16.25 21v-4.765a1.59 1.59 0 0 0-1.594-1.588H9.344a1.59 1.59 0 0 0-1.594 1.588V21m8.5-17.715v2.362a1.59 1.59 0 0 1-1.594 1.588H9.344A1.59 1.59 0 0 1 7.75 5.647V3m8.5.285A3.196 3.196 0 0 0 14.93 3H7.75m8.5.285c.344.156.661.374.934.645l2.382 2.375A3.17 3.17 0 0 1 20.5 8.55v9.272A3.182 3.182 0 0 1 17.313 21H6.688A3.182 3.182 0 0 1 3.5 17.823V6.176A3.182 3.182 0 0 1 6.688 3H7.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Scan(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.114 7.5c.144-1.463.47-2.447 1.204-3.182C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318c.735.735 1.06 1.72 1.204 3.182m0 9c-.144 1.463-.47 2.447-1.204 3.182C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318c-.735-.735-1.06-1.72-1.204-3.182M3 12h.5m8.25 0h.5m-5 0h.5m8.5 0h.5m3.75 0h.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Scissors(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.79 12L7 8.855m14-2.903L7 15.145M8 6.92c0 1.337-1.12 2.42-2.5 2.42S3 8.256 3 6.919C3 5.583 4.12 4.5 5.5 4.5S8 5.583 8 6.92m0 10.162c0-1.336-1.12-2.42-2.5-2.42S3 15.745 3 17.081S4.12 19.5 5.5 19.5S8 18.417 8 17.08m13 .968l-6.066-3.983"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SeaWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3 5.918l1.764-.887a4.973 4.973 0 0 1 4.472 0l.528.266a4.973 4.973 0 0 0 4.472 0l.528-.266a4.973 4.973 0 0 1 4.472 0L21 5.918M3 10.446l1.764-.888a4.973 4.973 0 0 1 4.472 0l.528.266a4.973 4.973 0 0 0 4.472 0l.528-.266a4.973 4.973 0 0 1 4.472 0l1.764.888M3 14.973l1.764-.888a4.973 4.973 0 0 1 4.472 0l.528.266a4.973 4.973 0 0 0 4.472 0l.528-.266a4.973 4.973 0 0 1 4.472 0l1.764.888M3 19.5l1.764-.887a4.973 4.973 0 0 1 4.472 0l.528.265a4.973 4.973 0 0 0 4.472 0l.528-.265a4.973 4.973 0 0 1 4.472 0L21 19.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Search(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m16.893 16.92l3.08 3.08m-.889-8.419c0 4.187-3.383 7.581-7.555 7.581c-4.173 0-7.556-3.394-7.556-7.58C3.973 7.393 7.356 4 11.528 4c4.173 0 7.556 3.394 7.556 7.581"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchCheck(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m16.893 16.92l3.08 3.08m-.889-8.419c0 4.187-3.383 7.581-7.555 7.581c-4.173 0-7.556-3.394-7.556-7.58C3.973 7.393 7.356 4 11.528 4c4.173 0 7.556 3.394 7.556 7.581"/><path d="m9.8 11.992l1.034 1.181c.095.109.266.1.35-.016l2.1-2.907"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M13.856 13.85a3.429 3.429 0 1 0-4.855-4.842a3.429 3.429 0 0 0 4.855 4.842m0 0L16 16"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13.856 13.85a3.429 3.429 0 1 0-4.855-4.842a3.429 3.429 0 0 0 4.855 4.842m0 0L16 16"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchMinus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m16.893 16.92l3.08 3.08m-.889-8.419c0 4.187-3.383 7.581-7.555 7.581c-4.173 0-7.556-3.394-7.556-7.58C3.973 7.393 7.356 4 11.528 4c4.173 0 7.556 3.394 7.556 7.581M9.5 11.75h4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/><path d="M13.856 13.85a3.429 3.429 0 1 0-4.855-4.842a3.429 3.429 0 0 0 4.855 4.842m0 0L16 16"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchPlus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m16.893 16.92l3.08 3.08m-.889-8.419c0 4.187-3.383 7.581-7.555 7.581c-4.173 0-7.556-3.394-7.556-7.58C3.973 7.393 7.356 4 11.528 4c4.173 0 7.556 3.394 7.556 7.581M9.5 11.5h4m-2-2v4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchSlash(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16.893 16.92a7.57 7.57 0 0 0 2.191-5.339a7.566 7.566 0 0 0-1.779-4.886m-.412 10.225a7.52 7.52 0 0 1-5.365 2.242a7.51 7.51 0 0 1-4.889-1.801m10.254-.441l3.08 3.08M17.305 6.695L21 3m-3.695 3.695L6.64 17.36M3 21l3.639-3.639m-1.904-2.458a7.572 7.572 0 0 1-.762-3.322C3.973 7.394 7.356 4 11.528 4a7.51 7.51 0 0 1 3.293.756"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/><path d="M13.856 13.85a3.429 3.429 0 1 0-4.855-4.842a3.429 3.429 0 0 0 4.855 4.842m0 0L16 16"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13.856 13.85a3.429 3.429 0 1 0-4.855-4.842a3.429 3.429 0 0 0 4.855 4.842m0 0L16 16"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchX(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m16.893 16.92l3.08 3.08m-.889-8.419c0 4.187-3.383 7.581-7.555 7.581c-4.173 0-7.556-3.394-7.556-7.58C3.973 7.393 7.356 4 11.528 4c4.173 0 7.556 3.394 7.556 7.581M10.25 10.25l2.829 2.828m-.001-2.828l-2.828 2.828"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SelectMultiple(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 9v10.4c0 .56 0 .84.109 1.054a1 1 0 0 0 .437.437C3.76 21 4.04 21 4.598 21H15m-8-7.2V6.2c0-1.12 0-1.68.218-2.108c.192-.377.497-.682.874-.874C8.52 3 9.08 3 10.2 3h7.6c1.12 0 1.68 0 2.108.218a2 2 0 0 1 .874.874C21 4.52 21 5.08 21 6.2v7.6c0 1.12 0 1.68-.218 2.108a2.001 2.001 0 0 1-.874.874c-.428.218-.986.218-2.104.218h-7.607c-1.118 0-1.678 0-2.105-.218a2 2 0 0 1-.874-.874C7 15.48 7 14.92 7 13.8"/><path d="m11.6 10.323l1.379 1.575a.299.299 0 0 0 .466-.022L16.245 8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Send(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m14 10l-3 3m9.288-9.969a.535.535 0 0 1 .68.681l-5.924 16.93a.535.535 0 0 1-.994.04l-3.219-7.242a.535.535 0 0 0-.271-.271l-7.242-3.22a.535.535 0 0 1 .04-.993z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Servers(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 5.7c0-.663.448-1.2 1-1.2h16c.552 0 1 .537 1 1.2v3.6c0 .663-.448 1.2-1 1.2H4c-.552 0-1-.537-1-1.2zm3 1.8h2m-2 9h2m-5-1.8c0-.663.448-1.2 1-1.2h16c.552 0 1 .537 1 1.2v3.6c0 .663-.448 1.2-1 1.2H4c-.552 0-1-.537-1-1.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Seven(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.5 17c0-3.75 5-10 5-10s-3.75.625-6.25 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SevenCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M10.75 16c0-3 4-8 4-8s-3 .5-5 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SevenDiamond(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.75 16c0-3 4-8 4-8s-3 .5-5 0"/><path d="M2.707 10.295a2.41 2.41 0 0 0 0 3.41l7.588 7.588a2.409 2.409 0 0 0 3.41 0l7.588-7.588a2.409 2.409 0 0 0 0-3.41l-7.588-7.588a2.41 2.41 0 0 0-3.41 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SevenHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.75 16c0-3 4-8 4-8s-3 .5-5 0"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SevenOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.75 16c0-3 4-8 4-8s-3 .5-5 0"/><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SevenSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/><path d="M10.75 16c0-3 4-8 4-8s-3 .5-5 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SevenWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/><path d="M10.75 16c0-3 4-8 4-8s-3 .5-5 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Share(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.5 12a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0m5-5.5l-5 3.5m5 7.5l-5-3.5m10 4.5a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0m0-13a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shield(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.467 20.82a.88.88 0 0 0 1.066 0C14.168 19.593 19 15.586 19 11.016v-4.93a.514.514 0 0 0-.457-.515a12.048 12.048 0 0 1-5.582-2.046l-.61-.417a.62.62 0 0 0-.702 0l-.61.417a12.048 12.048 0 0 1-5.582 2.046a.514.514 0 0 0-.457.515v4.93c0 4.57 4.832 8.577 6.467 9.802"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldCheck(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m10.258 11.242l1.034 1.181c.095.109.266.1.35-.016l2.1-2.907"/><path d="M11.467 20.82a.88.88 0 0 0 1.066 0C14.168 19.593 19 15.586 19 11.016v-4.93a.514.514 0 0 0-.457-.515a12.048 12.048 0 0 1-5.582-2.046l-.61-.417a.62.62 0 0 0-.702 0l-.61.417a12.048 12.048 0 0 1-5.582 2.046a.514.514 0 0 0-.457.515v4.93c0 4.57 4.832 8.577 6.467 9.802"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldCrossed(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 21s7-4.6 7-10m-7 10s-7-4.6-7-10m7 10V3m7 8V6.16a.509.509 0 0 0-.457-.506c-1.998-.2-3.915-.89-5.582-2.009L12 3m7 8H5m7-8l-.961.645a12.188 12.188 0 0 1-5.582 2.01A.508.508 0 0 0 5 6.16V11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldMinus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 11h4m-2.533 9.82a.88.88 0 0 0 1.066 0C14.168 19.593 19 15.586 19 11.016v-4.93a.514.514 0 0 0-.457-.515a12.048 12.048 0 0 1-5.582-2.046l-.61-.417a.62.62 0 0 0-.702 0l-.61.417a12.048 12.048 0 0 1-5.582 2.046a.514.514 0 0 0-.457.515v4.93c0 4.57 4.832 8.577 6.467 9.802"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldOne(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.183 15.092C5.478 13.836 5 12.455 5 11.018v-4.93c0-.267.198-.489.457-.515a12.048 12.048 0 0 0 5.582-2.047l.61-.417a.62.62 0 0 1 .702 0l.61.417a12.048 12.048 0 0 0 5.582 2.047c.26.026.457.248.457.514v4.93c0 1.438-.478 2.819-1.183 4.075m-11.634 0c1.538 2.74 4.16 4.887 5.282 5.727a.882.882 0 0 0 1.07 0c1.122-.84 3.744-2.988 5.282-5.727m-11.634 0l5.25-3.677a.985.985 0 0 1 1.134 0l5.25 3.677"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldPlus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 11h4m-2-2v4m-.533 7.82a.88.88 0 0 0 1.066 0C14.168 19.593 19 15.586 19 11.016v-4.93a.514.514 0 0 0-.457-.515a12.048 12.048 0 0 1-5.582-2.046l-.61-.417a.62.62 0 0 0-.702 0l-.61.417a12.048 12.048 0 0 1-5.582 2.046a.514.514 0 0 0-.457.515v4.93c0 4.57 4.832 8.577 6.467 9.802"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldSlash(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3 21l4.27-4.27M21 3l-2.561 2.561m0 0l.104.011c.26.027.457.25.457.515v4.93c0 4.57-4.832 8.577-6.467 9.802a.877.877 0 0 1-1.066 0c-.886-.663-2.711-2.144-4.197-4.09M18.44 5.562L7.27 16.73M5.646 14c-.4-.949-.646-1.951-.646-2.983v-4.93c0-.266.198-.488.457-.515a12.048 12.048 0 0 0 5.582-2.046l.61-.417a.62.62 0 0 1 .702 0l.61.417c.646.442 1.329.817 2.039 1.124"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldTwo(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 21s7-4.6 7-10V6.16a.509.509 0 0 0-.457-.506c-1.998-.2-3.915-.89-5.582-2.009L12 3m0 18s-7-4.6-7-10V6.16c0-.261.198-.48.457-.506c1.998-.2 3.915-.89 5.582-2.009L12 3m0 18V3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldX(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m10.5 9.5l2.829 2.828M13.328 9.5L10.5 12.328m.967 8.492a.88.88 0 0 0 1.066 0C14.168 19.593 19 15.586 19 11.016v-4.93a.514.514 0 0 0-.457-.515a12.048 12.048 0 0 1-5.582-2.046l-.61-.417a.62.62 0 0 0-.702 0l-.61.417a12.048 12.048 0 0 1-5.582 2.046a.514.514 0 0 0-.457.515v4.93c0 4.57 4.832 8.577 6.467 9.802"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShootingStar(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 5.647L9.353 3m.53 6.882L4.058 4.06M5.647 12L3 9.353m5.294 5.294l4.235-2.118l2.118-4.235l2.118 4.235L21 14.647l-4.235 2.118L14.647 21l-2.118-4.235z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBag(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.935 7H7.773c-1.072 0-1.962.81-2.038 1.858l-.73 10C4.921 20.015 5.858 21 7.043 21h9.914c1.185 0 2.122-.985 2.038-2.142l-.73-10C18.189 7.81 17.299 7 16.227 7h-1.162m-6.13 0V5c0-1.105.915-2 2.043-2h2.044c1.128 0 2.043.895 2.043 2v2m-6.13 0h6.13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shuffle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.5 4h4.412v4.444M4 20L19 5m1 10.5V20h-4.5M14 14l5.294 5.333M4 4l6 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShuffleAlt(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m19 16.765l2 2.117L19 21m0-18l2 2.118l-2 2.117"/><path d="M21 5.118h-3.15C14.62 5.118 12 8.199 12 12c0 3.801 2.62 6.882 5.85 6.882H21m-18 0h3.15C9.38 18.882 12 15.801 12 12c0-3.801-2.62-6.882-5.85-6.882H3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sidebar(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 3.5v17M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SidebarAlt(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15 3.5v17M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Signal(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20.5 10v4m-4-8v12M12 3v18M7.5 6v12m-4-8v4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignalCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 13.5v-3m3 4.5V9m3 4.5v-3"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignalHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 13.5v-3m3 4.5V9m3 4.5v-3m5.5 5.3V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignalOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132zM9 13.5v-3m3 4.5V9m3 4.5v-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignalSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 13.5v-3m3 4.5V9m3 4.5v-3M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignalWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 13.5v-3m3 4.5V9m3 4.5v-3M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Six(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13.75 7H12.5a3.75 3.75 0 0 0-3.75 3.75v3.125m0 0a3.125 3.125 0 1 0 6.25 0c0-1.726-1.4-2.5-3.125-2.5c-1.726 0-3.125.774-3.125 2.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SixCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M13.5 8h-1a3 3 0 0 0-3 3v2.5m0 0a2.5 2.5 0 0 0 5 0c0-1.38-1.12-2-2.5-2s-2.5.62-2.5 2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SixDiamond(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13.5 8h-1a3 3 0 0 0-3 3v2.5m0 0a2.5 2.5 0 0 0 5 0c0-1.38-1.12-2-2.5-2s-2.5.62-2.5 2"/><path d="M2.707 10.295a2.41 2.41 0 0 0 0 3.41l7.588 7.588a2.409 2.409 0 0 0 3.41 0l7.588-7.588a2.409 2.409 0 0 0 0-3.41l-7.588-7.588a2.41 2.41 0 0 0-3.41 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SixHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13.5 8h-1a3 3 0 0 0-3 3v2.5m0 0a2.5 2.5 0 0 0 5 0c0-1.38-1.12-2-2.5-2s-2.5.62-2.5 2"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SixOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13.5 8h-1a3 3 0 0 0-3 3v2.5m0 0a2.5 2.5 0 0 0 5 0c0-1.38-1.12-2-2.5-2s-2.5.62-2.5 2"/><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SixSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13.5 8h-1a3 3 0 0 0-3 3v2.5m0 0a2.5 2.5 0 0 0 5 0c0-1.38-1.12-2-2.5-2s-2.5.62-2.5 2"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SixWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13.5 8h-1a3 3 0 0 0-3 3v2.5m0 0a2.5 2.5 0 0 0 5 0c0-1.38-1.12-2-2.5-2s-2.5.62-2.5 2"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SkipBack(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.5 3v18m4.726-8.22l8.65 6.92a1 1 0 0 0 1.624-.78V5.08a1 1 0 0 0-1.625-.78l-8.649 6.92a1 1 0 0 0 0 1.56"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SkipForward(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.5 3v18m-4.726-8.22l-8.65 6.92a1 1 0 0 1-1.624-.78V5.08a1 1 0 0 1 1.625-.78l8.649 6.92a1 1 0 0 1 0 1.56"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Slash(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m9 15l6-6"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SlashCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="m9 15l6-6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SlashHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9 15l6-6m5.5 6.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SlashOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132zM9 15l6-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SlashSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12m6 3l6-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SlashWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9 15l6-6M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmileCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 15c.85.63 1.885 1 3 1s2.15-.37 3-1m-5.5-4.5V10m5 .5V10"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmileGhost(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 15c.85.63 1.885 1 3 1s2.15-.37 3-1m-5.5-4.5V10m5 .5V10"/><path d="M3 18.562v-6.518C3 7.05 7.03 3 12 3s9 4.05 9 9.044v6.517c0 1.162-.967 2.519-2 2c-.835-.42-2.223-.52-3 0c-.874.585-2.126.585-3 0c-.885-.593-1.649-.57-2.5 0c-.874.585-2.126.585-3 0c-.777-.52-1.665-.42-2.5 0c-1.033.519-2-.838-2-2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SmileSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.5 10.5V10m5 .5V10M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/><path d="M9 15c.85.63 1.885 1 3 1s2.15-.37 3-1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Snow(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 19V5M9.953 3L12 5l2.047-2M9.953 21L12 19l2.047 2m-8.251-5.5l12.408-7m.749-2.732L18.204 8.5L21 9.232M3 14.768l2.796.732l-.75 2.732m.75-9.732l12.408 7M21 14.768l-2.796.732l.75 2.732M5.047 5.768L5.796 8.5L3 9.232"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sofa(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.5 18v2m9-2v2M4.8 9.844C3.9 9.864 3 10.271 3 11v4c0 1.657 1.209 3 2.7 3h12.6c1.491 0 2.7-1.343 2.7-3v-4c0-.812-.9-1.176-1.8-1.156m-14.4 0c.9-.02 1.8.344 1.8 1.156c0 1.363.225 3 2 3h6.8c1.775 0 2-1.637 2-3c0-.73.9-1.135 1.8-1.156m-14.4 0V7c0-1.657 1.209-3 2.7-3h9c1.491 0 2.7 1.343 2.7 3v2.844"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sort(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.5 7h15m-15 5h10m-10 5h4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sparkles(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.77 21c1.123-4.649 2.486-6.099 6.23-7.364c-3.934-1.328-5.16-2.94-6.23-7.363c-1.124 4.649-2.488 6.098-6.232 7.363c3.93 1.327 5.163 2.95 6.231 7.364m-8.308-9.818c.512-2.42 1.502-3.512 3.461-4.091C7.963 6.512 6.973 5.42 6.462 3C5.972 5.315 5.047 6.485 3 7.09c1.959.58 2.95 1.672 3.462 4.092"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Speaker(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 14a3 3 0 1 0 6 0a3 3 0 0 0-6 0m2.5-7h1"/><path d="M12 3c3.418 0 5.127 0 6.188 1.318C19.25 5.636 19.25 7.758 19.25 12c0 4.243 0 6.364-1.062 7.682C17.127 21 15.418 21 12 21c-3.418 0-5.127 0-6.188-1.318C4.75 18.364 4.75 16.242 4.75 12c0-4.243 0-6.364 1.062-7.682C6.873 3 8.582 3 12 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spinner(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 3v3m6.366-.366l-2.12 2.12M21 12h-3m.366 6.366l-2.12-2.12M12 21v-3m-6.366.366l2.12-2.12M3 12h3m-.366-6.366l2.12 2.12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpinnerOne(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 3v3m0 15v-3m-7.794-1.5L6.804 15M21 12h-3m-1.5 7.794L15 17.196M3 12h3m1.5-7.794L9 6.804m-1.5 12.99L9 17.196m10.794-.696L17.196 15M4.206 7.5L6.804 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Square(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareDashed(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-dasharray="3.5 3.5" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareHalf(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 3.5v17M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Star(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.854 3.5a.979.979 0 0 0-1.708 0a26.978 26.978 0 0 0-2.057 4.762c-.139.431-.551.73-1.023.743a29.398 29.398 0 0 0-4.267.425c-.774.136-1.065 1.018-.515 1.556a31.484 31.484 0 0 0 3.41 2.892c.367.269.518.73.378 1.152a26.807 26.807 0 0 0-1.14 4.927c-.1.755.708 1.288 1.41.928a28.593 28.593 0 0 0 3.98-2.472a1.148 1.148 0 0 1 1.356 0a28.505 28.505 0 0 0 3.98 2.472c.701.36 1.51-.173 1.41-.928a26.81 26.81 0 0 0-1.14-4.928c-.14-.42.01-.882.378-1.151a31.497 31.497 0 0 0 3.41-2.892c.55-.538.26-1.42-.515-1.556a29.046 29.046 0 0 0-4.267-.425a1.097 1.097 0 0 1-1.023-.743a26.982 26.982 0 0 0-2.057-4.761"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stop(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 12.5c0-2.828 0-4.243.879-5.121C7.757 6.5 9.172 6.5 12 6.5c2.828 0 4.243 0 5.121.879C18 8.257 18 9.672 18 12.5c0 2.828 0 4.243-.879 5.121c-.878.879-2.293.879-5.121.879c-2.828 0-4.243 0-5.121-.879C6 16.743 6 15.328 6 12.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.5 12c0-1.178 0-1.768.366-2.134c.366-.366.956-.366 2.134-.366s1.768 0 2.134.366c.366.366.366.956.366 2.134s0 1.768-.366 2.134c-.366.366-.956.366-2.134.366s-1.768 0-2.134-.366C9.5 13.768 9.5 13.178 9.5 12"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.5 12c0-1.178 0-1.768.366-2.134c.366-.366.956-.366 2.134-.366s1.768 0 2.134.366c.366.366.366.956.366 2.134s0 1.768-.366 2.134c-.366.366-.956.366-2.134.366s-1.768 0-2.134-.366C9.5 13.768 9.5 13.178 9.5 12"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/><path d="M9.5 12c0-1.178 0-1.768.366-2.134c.366-.366.956-.366 2.134-.366s1.768 0 2.134.366c.366.366.366.956.366 2.134s0 1.768-.366 2.134c-.366.366-.956.366-2.134.366s-1.768 0-2.134-.366C9.5 13.768 9.5 13.178 9.5 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.5 12c0-1.178 0-1.768.366-2.134c.366-.366.956-.366 2.134-.366s1.768 0 2.134.366c.366.366.366.956.366 2.134s0 1.768-.366 2.134c-.366.366-.956.366-2.134.366s-1.768 0-2.134-.366C9.5 13.768 9.5 13.178 9.5 12"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.5 12c0-1.178 0-1.768.366-2.134c.366-.366.956-.366 2.134-.366s1.768 0 2.134.366c.366.366.366.956.366 2.134s0 1.768-.366 2.134c-.366.366-.956.366-2.134.366s-1.768 0-2.134-.366C9.5 13.768 9.5 13.178 9.5 12"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Store(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.996 10.621V19a2 2 0 0 1-2 2H6.004a2 2 0 0 1-1.999-2v-8.379M7.502 8.75l.5-5.75m-.5 5.75c0 2.902 4.498 2.902 4.498 0m-4.498 0c0 3.176-5.155 2.52-4.433-.248l1.045-4.007A2 2 0 0 1 6.048 3h11.904a2 2 0 0 1 1.934 1.495l1.045 4.007c.722 2.769-4.433 3.424-4.433.248M12 8.75V3m0 5.75c0 2.902 4.498 2.902 4.498 0m0 0l-.5-5.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Subtract(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.6 4.35V3h-1.8M8.4 3h1.8M4.8 3H3v1.8m0 5.4V8.4m0 5.4v1.8h1.35m5.85 0H8.4v.6c0 1.68 0 2.52.327 3.162a3 3 0 0 0 1.311 1.311C10.68 21 11.52 21 13.2 21h3c1.68 0 2.52 0 3.162-.327a3 3 0 0 0 1.311-1.311C21 18.72 21 17.88 21 16.2v-3c0-1.68 0-2.52-.327-3.162a3 3 0 0 0-1.311-1.311C18.72 8.4 17.88 8.4 16.2 8.4h-.6v1.8m-1.8 5.4h1.8v-1.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sun(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="4"/><path d="M12 3v2m0 14.004v2M5 12H3m18 0h-2m0-7l-2 2M5 5l2 2m0 10l-2 2m14 0l-2-2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sunrise(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.5 17.5a3.5 3.5 0 1 0-7 0M11.9 3v7m-6.002 1.398l1.278 1.278M3 17.4h1.8m14.2 0h1.8m-4.176-4.724l1.278-1.278M21 21H3M8.3 6.6L11.9 3l3.6 3.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sunset(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.5 17.5a3.5 3.5 0 1 0-7 0M11.9 3v7m-6.002 1.398l1.278 1.278M3 17.4h1.8m14.2 0h1.8m-4.176-4.724l1.278-1.278M21 21H3M8.3 7l3.6 3.6L15.5 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Support(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="m18 6l-3.525 3.525M6 18l3.525-3.525M6 6l3.525 3.525M18 18l-3.525-3.525m-4.95 0c-1.348-1.348-1.348-3.601 0-4.95m0 4.95c1.348 1.348 3.601 1.348 4.95 0m0 0c1.348-1.348 1.348-3.601 0-4.95m0 0c-1.348-1.348-3.601-1.348-4.95 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Swatches(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.5 21H18c1.657 0 3-1.398 3-3.123c0-1.308.13-2.63-1.297-3.253M7.98 20.664l10.287-4.67a3.037 3.037 0 0 0 1.436-1.37m-7.613-3.787l3.013-1.718c1.553-.886 3.5-.186 4.198 1.509l.525 1.273a3.232 3.232 0 0 1-.123 2.723m-9.821 3.718c-.5 1.912-2.42 3.047-4.287 2.535c-1.867-.513-2.975-2.478-2.475-4.39L6.18 5.27c.452-1.661 2.114-2.624 3.71-2.15l1.17.348c1.615.48 2.536 2.246 2.045 3.918zM6.5 17l-.242.854"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Table(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 3.5v17m11.5-11h-17M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tablet(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17.714 3H6.286C5.023 3 4 3.806 4 4.8v14.4c0 .994 1.023 1.8 2.286 1.8h11.428C18.977 21 20 20.194 20 19.2V4.8c0-.994-1.023-1.8-2.286-1.8M10.5 6h3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tag(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.472 5.5H14.77a2 2 0 0 1 1.396.568l5.35 5.216a1 1 0 0 1 0 1.432l-5.35 5.216a2 2 0 0 1-1.396.568H4.472c-.95 0-2.222-.541-2.222-1.625v-9.75C2.25 6.041 3.523 5.5 4.472 5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagPlus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.472 5.5H14.77a2 2 0 0 1 1.396.568l5.35 5.216a1 1 0 0 1 0 1.432l-5.35 5.216a2 2 0 0 1-1.396.568H4.472c-.95 0-2.222-.541-2.222-1.625v-9.75C2.25 6.041 3.523 5.5 4.472 5.5M6 12h5M8.5 9.5v5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TallyFive(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22 4L2 20M4 4v16M9.333 4v16m5.334-16v16M20 4v16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TallyFour(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 4v16M9.333 4v16m5.334-16v16M20 4v16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TallyOne(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 4v16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TallyThree(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 4v16M9.333 4v16m5.334-16v16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TallyTwo(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 4v16M9.333 4v16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Target(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 17.25a5.25 5.25 0 1 0 0-10.5a5.25 5.25 0 0 0 0 10.5"/><path d="M12 13.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Telephone(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.6 14.522c-2.395 2.52-8.504-3.534-6.1-6.064c1.468-1.545-.19-3.31-1.108-4.609c-1.723-2.435-5.504.927-5.39 3.066c.363 6.746 7.66 14.74 14.726 14.042c2.21-.218 4.75-4.21 2.214-5.669c-1.267-.73-3.008-2.17-4.342-.767"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TelephoneCall(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.6 14.522c-2.395 2.52-8.504-3.534-6.1-6.064c1.468-1.545-.19-3.31-1.108-4.609c-1.723-2.435-5.504.927-5.39 3.066c.363 6.746 7.66 14.74 14.726 14.042c2.21-.218 4.75-4.21 2.214-5.669c-1.267-.73-3.008-2.17-4.342-.767M14 3a7 7 0 0 1 7 7m-7-3a3 3 0 0 1 3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TelephoneForward(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.6 14.522c-2.395 2.52-8.504-3.534-6.1-6.064c1.468-1.545-.19-3.31-1.108-4.609c-1.723-2.435-5.504.927-5.39 3.066c.363 6.746 7.66 14.74 14.726 14.042c2.21-.218 4.75-4.21 2.214-5.669c-1.267-.73-3.008-2.17-4.342-.767M15 5.5h6.429M19 3l2.5 2.5L19 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TelephoneIn(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.6 14.522c-2.395 2.52-8.504-3.534-6.1-6.064c1.468-1.545-.19-3.31-1.108-4.609c-1.723-2.435-5.504.927-5.39 3.066c.363 6.746 7.66 14.74 14.726 14.042c2.21-.218 4.75-4.21 2.214-5.669c-1.267-.73-3.008-2.17-4.342-.767M20.5 3L16 7.5m3.5 0H16V4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TelephoneMissed(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.6 14.522c-2.395 2.52-8.504-3.534-6.1-6.064c1.468-1.545-.19-3.31-1.108-4.609c-1.723-2.435-5.504.927-5.39 3.066c.363 6.746 7.66 14.74 14.726 14.042c2.21-.218 4.75-4.21 2.214-5.669c-1.267-.73-3.008-2.17-4.342-.767M17 3l4 4m0-4l-4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TelephoneOut(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.6 14.522c-2.395 2.52-8.504-3.534-6.1-6.064c1.468-1.545-.19-3.31-1.108-4.609c-1.723-2.435-5.504.927-5.39 3.066c.363 6.746 7.66 14.74 14.726 14.042c2.21-.218 4.75-4.21 2.214-5.669c-1.267-.73-3.008-2.17-4.342-.767M15 7.5L19.5 3M16 3h3.5v3.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TelephoneSlash(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.684 13.316L21 3M10.684 13.316c1.53 1.556 3.701 2.484 4.916 1.205c1.334-1.404 3.075.038 4.342.767c2.536 1.46-.004 5.451-2.214 5.67c-3.657.36-7.376-1.606-10.163-4.523m3.119-3.119l-3.119 3.119M3 21l4.565-4.565M5.5 13.845c-1.447-2.193-2.374-4.634-2.497-6.93c-.115-2.139 3.666-5.501 5.389-3.066c.918 1.298 2.576 3.064 1.108 4.609c-.493.519-.628 1.186-.51 1.897"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Terminal(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8 9l3 3l-3 3m5 0h3"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextAlignCenter(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.5 6h15M7 10h10M4.5 14h15M7 18h10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextAlignLeft(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.5 6h15m-15 4h10m-10 4h15m-15 4h10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextAlignRight(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.5 6h15m-10 4h10m-15 4h15m-10 4h10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextJustify(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.5 6h15m-15 4h15m-15 4h15m-15 4h15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Thermometer(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="1.5" d="M14.155 13.86a.326.326 0 0 1-.114-.116a.312.312 0 0 1-.041-.155v-8.66c0-.512-.21-1.002-.586-1.364A2.038 2.038 0 0 0 12 3c-.53 0-1.04.203-1.414.565A1.894 1.894 0 0 0 10 4.929v8.66a.313.313 0 0 1-.041.155a.327.327 0 0 1-.114.116a3.971 3.971 0 0 0-1.396 1.493a3.803 3.803 0 0 0-.445 1.965a3.8 3.8 0 0 0 1.266 2.644a4.085 4.085 0 0 0 2.82 1.037a4.071 4.071 0 0 0 2.77-1.16A3.787 3.787 0 0 0 16 17.145c0-.652-.168-1.294-.49-1.867a3.976 3.976 0 0 0-1.355-1.417"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Three(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.75 9.493c0-3.324 6.25-3.324 6.25 0c0 0 0 2.507-3.125 2.507C15 12 15 14.507 15 14.507c0 3.324-6.25 3.324-6.25 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThreeCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M9.5 9.994c0-2.659 5-2.659 5 0c0 0 0 2.006-2.5 2.006c2.5 0 2.5 2.006 2.5 2.006c0 2.659-5 2.659-5 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThreeDiamond(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.5 9.994c0-2.659 5-2.659 5 0c0 0 0 2.006-2.5 2.006c2.5 0 2.5 2.006 2.5 2.006c0 2.659-5 2.659-5 0"/><path d="M2.707 10.295a2.41 2.41 0 0 0 0 3.41l7.588 7.588a2.409 2.409 0 0 0 3.41 0l7.588-7.588a2.409 2.409 0 0 0 0-3.41l-7.588-7.588a2.41 2.41 0 0 0-3.41 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThreeHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.5 9.994c0-2.659 5-2.659 5 0c0 0 0 2.006-2.5 2.006c2.5 0 2.5 2.006 2.5 2.006c0 2.659-5 2.659-5 0"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThreeOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.5 9.994c0-2.659 5-2.659 5 0c0 0 0 2.006-2.5 2.006c2.5 0 2.5 2.006 2.5 2.006c0 2.659-5 2.659-5 0"/><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThreeSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.5 9.994c0-2.659 5-2.659 5 0c0 0 0 2.006-2.5 2.006c2.5 0 2.5 2.006 2.5 2.006c0 2.659-5 2.659-5 0"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThreeWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.5 9.994c0-2.659 5-2.659 5 0c0 0 0 2.006-2.5 2.006c2.5 0 2.5 2.006 2.5 2.006c0 2.659-5 2.659-5 0"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ticket(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 19a2 2 0 0 0 2-2v-3a2 2 0 1 1 0-4V7a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v3a2 2 0 1 1 0 4v3a2 2 0 0 0 2 2zm-7-7.25v.5M12 8v.5m0 7v.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TicketSlash(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 8v.5m0 7v.5m-9 1v-3a2 2 0 1 0 0-4V7a2 2 0 0 1 2-2h10M5 19h14a2 2 0 0 0 2-2v-3a2 2 0 1 1 0-4V7a2 2 0 0 0-2-2M3 21L21 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToggleLeft(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 12a5 5 0 0 0-5-5H8a5 5 0 0 0 0 10h8a5 5 0 0 0 5-5"/><path d="M5.5 12a2.5 2.5 0 1 0 5 0a2.5 2.5 0 0 0-5 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToggleRight(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 12a5 5 0 0 1 5-5h8a5 5 0 0 1 0 10H8a5 5 0 0 1-5-5"/><path d="M18.5 12a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tool(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 7.86c0-.43-.056-.849-.161-1.246c-.092-.349-.522-.432-.776-.177L18.34 8.16a1.767 1.767 0 1 1-2.5-2.5l1.723-1.722c.255-.255.172-.685-.177-.777a4.86 4.86 0 0 0-5.828 6.326c.071.2.031.424-.118.573L3.3 18.2c-.4.4-.4 1.049 0 1.448L4.352 20.7c.4.4 1.047.4 1.447 0l8.14-8.14c.15-.15.374-.19.573-.119A4.86 4.86 0 0 0 21 7.86"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Train(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 18H8m8 0l2 3m-2-3a3 3 0 0 0 3-3v-4M8 18l-2 3m2-3a3 3 0 0 1-3-3v-4m3 3v1m8-1v1M5 11h7m-7 0V6m7 5h7m-7 0V6m7 5V6m-7 0H5m7 0h7M5 6a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trash(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.286 8.571L7.429 20h9.142l1.143-11.429M13.5 15.5v-5m-3 5v-5M4.571 6.286h4.572m0 0l.382-1.529a1 1 0 0 1 .97-.757h3.01a1 1 0 0 1 .97.757l.382 1.529m-5.714 0h5.714m0 0h4.572"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tree(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 17h7l-4.5-6.5h3L12 3l-5.5 7.5h3L5 17zm0 0v4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrendingDown(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15 18h6v-6"/><path d="m3 7l4.443 5.223c.31.365.466.547.658.64a1 1 0 0 0 .546.09c.212-.024.418-.146.83-.39l2.826-1.674c.385-.229.578-.343.778-.37a1 1 0 0 1 .52.068c.187.077.344.237.658.556L21 18"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrendingUp(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15 7h6v6"/><path d="m3 18l4.443-5.223c.31-.365.466-.547.658-.64a1 1 0 0 1 .546-.09c.212.024.418.146.83.39l2.826 1.674c.385.229.578.343.778.37a1 1 0 0 0 .52-.068c.187-.077.344-.237.658-.556L21 7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Triangle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.98 10.762C8.608 5.587 9.92 3 12 3c2.08 0 3.393 2.587 6.02 7.762l.327.644c2.182 4.3 3.274 6.45 2.287 8.022C19.648 21 17.208 21 12.327 21h-.654c-4.88 0-7.321 0-8.307-1.572c-.987-1.572.105-3.722 2.287-8.022z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleDanger(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 8.5V14m0 3.247v-.5m-6.02-5.986C8.608 5.587 9.92 3 12 3c2.08 0 3.393 2.587 6.02 7.761l.327.645c2.182 4.3 3.274 6.45 2.287 8.022C19.648 21 17.208 21 12.327 21h-.654c-4.88 0-7.321 0-8.307-1.572c-.987-1.572.105-3.722 2.287-8.022z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleInfo(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 17v-5h-.5m0 5h1M12 9.5V9"/><path d="M5.98 10.761C8.608 5.587 9.92 3 12 3c2.08 0 3.393 2.587 6.02 7.761l.327.645c2.182 4.3 3.274 6.45 2.287 8.022C19.648 21 17.208 21 12.327 21h-.654c-4.88 0-7.321 0-8.307-1.572c-.987-1.572.105-3.722 2.287-8.022z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Truck(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.207 16.455C9.207 17.86 8.095 19 6.724 19s-2.483-1.14-2.483-2.546m4.966 0c0-1.405-1.112-2.545-2.483-2.545s-2.483 1.14-2.483 2.545m4.966 0h5.586m-10.552 0H3V6a1 1 0 0 1 1-1h9.793a1 1 0 0 1 1 1v2.182m5.586 8.272c0 1.406-1.111 2.546-2.482 2.546c-1.372 0-2.483-1.14-2.483-2.546m4.965 0c0-1.405-1.111-2.545-2.482-2.545c-1.372 0-2.483 1.14-2.483 2.545m4.965 0H21v-5.09l-2.515-2.579a2 2 0 0 0-1.431-.603h-2.26m.62 8.272h-.62m0 0V8.182"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tv(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m12 7l6.632-4m-7.106 4L5.368 3M3 9.154C3 7.964 3.895 7 5 7h14c1.105 0 2 .964 2 2.154v9.692c0 1.19-.895 2.154-2 2.154H5c-1.105 0-2-.964-2-2.154z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Two(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.75 9.92c0-3.894 5.77-3.894 5.77 0c0 2.94-3.77 5.476-5.77 7.08c0 0 3.75-.625 6.25 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwoCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M9.5 10.337c0-3.116 4.615-3.116 4.615 0c0 2.352-3.015 4.38-4.615 5.663c0 0 3-.5 5 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwoDiamond(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.5 10.337c0-3.116 4.615-3.116 4.615 0c0 2.352-3.015 4.38-4.615 5.663c0 0 3-.5 5 0"/><path d="M2.707 10.295a2.41 2.41 0 0 0 0 3.41l7.588 7.588a2.409 2.409 0 0 0 3.41 0l7.588-7.588a2.409 2.409 0 0 0 0-3.41l-7.588-7.588a2.41 2.41 0 0 0-3.41 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwoHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.5 10.337c0-3.116 4.615-3.116 4.615 0c0 2.352-3.015 4.38-4.615 5.663c0 0 3-.5 5 0"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwoOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.5 10.337c0-3.116 4.615-3.116 4.615 0c0 2.352-3.015 4.38-4.615 5.663c0 0 3-.5 5 0"/><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwoSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.5 10.337c0-3.116 4.615-3.116 4.615 0c0 2.352-3.015 4.38-4.615 5.663c0 0 3-.5 5 0"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwoWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.5 10.337c0-3.116 4.615-3.116 4.615 0c0 2.352-3.015 4.38-4.615 5.663c0 0 3-.5 5 0"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TypeBold(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 3h8c1.06 0 2.078.474 2.828 1.318C16.578 5.162 17 6.307 17 7.5c0 1.193-.421 2.338-1.172 3.182C15.078 11.526 14.061 12 13 12H5zm0 9h10.039a4.44 4.44 0 0 1 3.154 1.318A4.52 4.52 0 0 1 19.5 16.5a4.52 4.52 0 0 1-1.307 3.182A4.442 4.442 0 0 1 15.038 21H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TypeItalic(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 3H9m6 18H5m9.5-18L10 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TypeText(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.53 20L7.764 4L3 20m1.596-5.03h6.337m4.244-3.03C16.765 10.933 21 9.925 21 13.451V20m0-5.541c-1.588.504-6.353.504-6.353 3.526c0 3.023 4.765 2.015 6.353-.504"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TypeUnderline(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 3v7.539c0 1.713.632 3.357 1.757 4.569C8.883 16.319 10.41 17 12 17c1.591 0 3.117-.68 4.243-1.892C17.368 13.896 18 12.252 18 10.538V3M4 21h16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Umbrella(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15.636 13c-2.424-2.424-4.848-2.424-7.272 0C5.878 10.87 4.486 10.87 2 13C2 7.477 6.477 3 12 3s10 4.477 10 10c-2.486-2.13-3.878-2.13-6.364 0"/><path d="M12 11.5v7.273c0 3.519-5.5 3.519-5.5 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Undo(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M5.636 18.364A9 9 0 1 0 12 3C7.942 3 5.482 5.705 3 8.5"/><path d="M3 4.5v4h4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Union(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 16.2v-2.1c0-1.68 0-2.52-.327-3.162a3 3 0 0 0-1.311-1.311C18.72 9.3 17.88 9.3 16.2 9.3h-1.5V7.8c0-1.68 0-2.52-.327-3.162a3 3 0 0 0-1.311-1.311C12.42 3 11.58 3 9.9 3H7.8c-1.68 0-2.52 0-3.162.327a3 3 0 0 0-1.311 1.311C3 5.28 3 6.12 3 7.8v2.1c0 1.68 0 2.52.327 3.162a3 3 0 0 0 1.311 1.311c.642.327 1.482.327 3.162.327h1.5v1.5c0 1.68 0 2.52.327 3.162a3 3 0 0 0 1.311 1.311C11.58 21 12.42 21 14.1 21h2.1c1.68 0 2.52 0 3.162-.327a3 3 0 0 0 1.311-1.311C21 18.72 21 17.88 21 16.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unlink(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m15.988 13l3.902-3.902c1.437-1.437 1.485-3.718.107-5.095c-1.377-1.378-3.658-1.33-5.095.107L11 8.012M3 9h1.5M9 4.5V3m12 12h-1.5M15 19.5V21m-2-5.038l-3.892 3.88c-1.432 1.43-3.64 1.615-5.082.107c-1.442-1.507-1.326-3.639.107-5.068L8.025 11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Upload(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 16.004V17a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-1m-8-.5v-11M15.5 8L12 4.5L8.5 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func User(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="7.5" r="3"/><path d="M19.5 20.5c-.475-9.333-14.525-9.333-15 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserCheck(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="7.5" r="3"/><path d="M19.5 20.5c-.475-9.333-14.525-9.333-15 0"/><path d="m10.258 18.992l1.034 1.181c.095.109.266.1.35-.016l2.1-2.907"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="9.1" r="2.5"/><circle cx="12" cy="12" r="9"/><path d="M17 19.2c-.317-6.187-9.683-6.187-10 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="9.1" r="2.5"/><path d="M7.046 18.705c.756-5.527 9.152-5.527 9.908 0"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserMinus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="7.5" r="3"/><path d="M19.5 20.5c-.475-9.333-14.525-9.333-15 0m5.5-2h4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="9.1" r="2.5"/><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/><path d="M7 19.525V19.2c.317-6.187 9.683-6.187 10 0v.325"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserPlus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="7.5" r="3"/><path d="M19.5 20.5c-.475-9.333-14.525-9.333-15 0m5.5-2h4m-2-2v4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserSettings(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="7.5" r="3"/><path d="M19.5 20.5c-.475-9.333-14.525-9.333-15 0"/><path d="M11.192 17.565c.394-.21.591-.315.808-.315c.217 0 .414.105.808.315l.134.072c.394.21.591.315.7.488c.108.173.108.383.108.804v.142c0 .42 0 .63-.108.804c-.109.173-.306.278-.7.488l-.134.072c-.394.21-.591.315-.808.315c-.217 0-.414-.105-.808-.315l-.134-.072c-.394-.21-.591-.315-.7-.488c-.108-.173-.108-.383-.108-.804v-.142c0-.42 0-.63.108-.804c.109-.173.306-.278.7-.488z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="9.1" r="2.5"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/><path d="M7 20.5v-1.3c.317-6.187 9.683-6.187 10 0v1.3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="9.1" r="2.5"/><path d="M17 19.2c-.317-6.187-9.683-6.187-10 0"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserX(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="7.5" r="3"/><path d="M19.5 20.5c-.475-9.333-14.525-9.333-15 0m6.086-3l2.828 2.828m0-2.828l-2.828 2.828"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Users(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21 19.75c0-2.09-1.67-5.068-4-5.727m-2 5.727c0-2.651-2.686-6-6-6s-6 3.349-6 6"/><circle cx="9" cy="7.25" r="3"/><path d="M15 10.25a3 3 0 1 0 0-6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UsersGroup(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 19.5c0-1.657-2.239-3-5-3s-5 1.343-5 3m14-3c0-1.23-1.234-2.287-3-2.75M3 16.5c0-1.23 1.234-2.287 3-2.75m12-4.014a3 3 0 1 0-4-4.472M6 9.736a3 3 0 0 1 4-4.472m2 8.236a3 3 0 1 1 0-6a3 3 0 0 1 0 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Video(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 15.75v-7.5a2 2 0 0 1 2-2h8.5a2 2 0 0 1 2 2v7.5a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2m17.168-8.759l-4 3.563a.5.5 0 0 0-.168.373v1.778a.5.5 0 0 0 .168.373l4 3.563a.5.5 0 0 0 .832-.374V7.365a.5.5 0 0 0-.832-.374"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoSlash(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m20.168 6.991l-4 3.563a.5.5 0 0 0-.168.373v1.778a.5.5 0 0 0 .168.373l4 3.563a.5.5 0 0 0 .832-.374V7.365a.5.5 0 0 0-.832-.374M3 15.75v-7.5a2 2 0 0 1 2-2h8.5M3 21l3.25-3.25M21 3l-5.5 5.5m0 0v7.25a2 2 0 0 1-2 2H6.25M15.5 8.5l-9.25 9.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeCheck(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13 3v18c-3.5 0-6.1-4.58-6.1-4.58s-3.9.586-3.9-1V8.505c0-1.586 3.9-1 3.9-1S9.5 3 13 3m3.36 9.323l1.379 1.575a.299.299 0 0 0 .466-.022l2.8-3.876"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeHigh(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13 3v18c-3.5 0-6.1-4.58-6.1-4.58s-3.9.586-3.9-1V8.505c0-1.586 3.9-1 3.9-1S9.5 3 13 3m5.1 16c3.866-3.866 3.866-10.134 0-14"/><path d="M16 16a5.657 5.657 0 0 0 0-8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeLow(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13 3v18c-3.5 0-6.1-4.58-6.1-4.58s-3.9.586-3.9-1V8.505c0-1.586 3.9-1 3.9-1S9.5 3 13 3m3 13a5.657 5.657 0 0 0 0-8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeMinus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13 3v18c-3.5 0-6.1-4.58-6.1-4.58s-3.9.586-3.9-1V8.505c0-1.586 3.9-1 3.9-1S9.5 3 13 3m3 9h5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeNone(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13 3v18c-3.5 0-6.1-4.58-6.1-4.58s-3.9.586-3.9-1V8.505c0-1.586 3.9-1 3.9-1S9.5 3 13 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumePlus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13 3v18c-3.5 0-6.1-4.58-6.1-4.58s-3.9.586-3.9-1V8.505c0-1.586 3.9-1 3.9-1S9.5 3 13 3m3 9h5m-2.5-2.5v5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeSlash(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18.1 19c3.715-3.715 3.86-9.648.436-13.536M16 16a5.657 5.657 0 0 0 0-8M3 21l4.16-4.16M21 3l-2.464 2.464m0 0L13 11m0 0v10c-2.846 0-5.098-3.029-5.84-4.16M13 11l-5.84 5.84M13 7V3C9.5 3 6.9 7.505 6.9 7.505S3 6.92 3 8.505v6.914c0 .39.236.65.59.818"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeX(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13 3v18c-3.5 0-6.1-4.58-6.1-4.58s-3.9.586-3.9-1V8.505c0-1.586 3.9-1 3.9-1S9.5 3 13 3m4 7l4 4m0-4l-4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Watch(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15.667 7c-.316-1.377-.418-4-2.348-4H10.68C8.751 3 8.65 5.623 8.333 7m0 10c.316 1.377.418 4 2.348 4h2.638c1.93 0 2.032-2.623 2.348-4m1.833-7v4c0 1.657-1.231 3-2.75 3h-5.5c-1.519 0-2.75-1.343-2.75-3v-4c0-1.657 1.231-3 2.75-3h5.5c1.519 0 2.75 1.343 2.75 3"/><path d="M12 10v2.5l1.604 1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Waves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Webcam(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 17.4A7.2 7.2 0 1 0 12 3a7.2 7.2 0 0 0 0 14.4m0 0V21m-4.5 0h9m-1.8-10.8a2.7 2.7 0 1 1-5.4 0a2.7 2.7 0 0 1 5.4 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wheel(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.794 16.5A9 9 0 0 1 7.5 19.794M19.794 16.5A9 9 0 0 0 16.5 4.206M19.794 16.5L13.732 13M7.5 19.794A9 9 0 0 1 4.206 7.5M7.5 19.794l3.5-6.062M4.206 7.5A9 9 0 0 1 16.5 4.206M4.206 7.5l6.062 3.5M16.5 4.206L13 10.268M13.732 13a2 2 0 0 1-2.732.732M13.732 13A2 2 0 0 0 13 10.268m-2 3.464A2 2 0 0 1 10.268 11m0 0A2 2 0 0 1 13 10.268m.932 1.214l6.761-1.811m-8.175 4.26l1.811 6.762m-4.26-8.175l-6.762 1.811m8.175-4.26L9.671 3.306"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wifi(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2 9.483c5.603-5.31 14.397-5.31 20 0"/><path d="M19 12.9c-3.866-3.867-10.134-3.867-14 0m11 3.257a5.657 5.657 0 0 0-8 0m4 3.093v-.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiCheck(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2 9.483A14.498 14.498 0 0 1 12 5.5m3 .822l1.379 1.576a.299.299 0 0 0 .466-.022L19.645 4M19 12.9c-3.866-3.867-10.134-3.867-14 0m11 3.257a5.657 5.657 0 0 0-8 0m4 3.093v-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiLow(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 16.157a5.657 5.657 0 0 0-8 0m4 3.093v-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiMedium(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 12.9c-3.866-3.867-10.134-3.867-14 0m11 3.257a5.657 5.657 0 0 0-8 0m4 3.093v-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiMinus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2 9.483A14.498 14.498 0 0 1 12 5.5m7 7.4c-3.866-3.867-10.134-3.867-14 0M15.5 6h5M16 16.157a5.657 5.657 0 0 0-8 0m4 3.093v-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiPlus(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2 9.483A14.498 14.498 0 0 1 12 5.5m7 7.4c-3.866-3.867-10.134-3.867-14 0M15.5 6h5M18 3.5v5m-2 7.657a5.657 5.657 0 0 0-8 0m4 3.093v-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiSlash(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M22 9.483a14.485 14.485 0 0 0-4.5-2.907M19 12.9a9.854 9.854 0 0 0-5-2.697M5 12.9a9.856 9.856 0 0 1 4-2.437m3 8.787v-.5M3 21L21 3M2 9.483A14.498 14.498 0 0 1 12 5.5m4 10.657a5.657 5.657 0 0 0-8 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiX(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2 9.483A14.498 14.498 0 0 1 12 5.5M16 4l4 4m0-4l-4 4m3 4.9c-3.866-3.867-10.134-3.867-14 0m11 3.257a5.657 5.657 0 0 0-8 0m4 3.093v-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Winds(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18.325 7c3.485 0 3.485 5 0 5H3m15.411 9c3.452 0 3.452-5 0-5H3m7.872-13c3.506 0 3.506 5 0 5H3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wine(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.4 21h7.2M12 21v-6.75m-.75 0h1.5a5 5 0 0 0 5-5V3.5a.5.5 0 0 0-.5-.5H6.75a.5.5 0 0 0-.5.5v5.75a5 5 0 0 0 5 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WinkCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 15c.85.63 1.885 1 3 1s2.15-.37 3-1m-.5-4.5V10m-5.5.5h1"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WinkGhost(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 15c.85.63 1.885 1 3 1s2.15-.37 3-1m-.5-4.5V10"/><path d="M3 18.562v-6.518C3 7.05 7.03 3 12 3s9 4.05 9 9.044v6.517c0 1.162-.967 2.519-2 2c-.835-.42-2.223-.52-3 0c-.874.585-2.126.585-3 0c-.885-.593-1.649-.57-2.5 0c-.874.585-2.126.585-3 0c-.777-.52-1.665-.42-2.5 0c-1.033.519-2-.838-2-2M9 10.5h1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WinkSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9 15c.85.63 1.885 1 3 1s2.15-.37 3-1m-.5-4.5V10"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12m6-1.5h1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wrench(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.917 7.076a.947.947 0 0 1 0 1.326L8.402 9.917a.947.947 0 0 1-1.326 0L4.528 7.37c-.495-.494-1.327-.333-1.446.356a5.682 5.682 0 0 0 6.626 6.554c.82-.15 1.707-.022 2.296.566l5.566 5.567a2.01 2.01 0 1 0 2.842-2.842l-5.567-5.566c-.588-.589-.716-1.477-.566-2.296a5.684 5.684 0 0 0-6.554-6.626c-.689.12-.85.951-.356 1.446z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func X(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 6L6 18M6 6l12 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Xcircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m15 9l-6 6m0-6l6 6"/><circle cx="12" cy="12" r="9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Xhexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m15 9l-6 6m0-6l6 6m5.5.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Xoctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132zM15 9l-6 6m0-6l6 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Xsquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m15 9l-6 6m0-6l6 6M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Xtriangle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.98 10.762C8.608 5.587 9.92 3 12 3c2.08 0 3.393 2.587 6.02 7.762l.327.644c2.182 4.3 3.274 6.45 2.287 8.022C19.648 21 17.208 21 12.327 21h-.654c-4.88 0-7.321 0-8.307-1.572c-.987-1.572.105-3.722 2.287-8.022zM14 11.5l-4 4m0-4l4 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Xwaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m15 9l-6 6m0-6l6 6M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Yen(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17.455 12H12m0 0H6.545M12 12L6 3.5m6 8.5l6-8.5M12 12v4.25m5.454 0H12m0 0H6.545m5.455 0v4.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YenCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M15.182 12.5H12m0 0H8.818m3.182 0l-3.5-5m3.5 5l3.5-5m-3.5 5V15m3.182 0H12m0 0H8.818M12 15v2.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YenHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8m-5.318-3.3H12m0 0H8.818m3.182 0l-3.5-5m3.5 5l3.5-5m-3.5 5V15m3.182 0H12m0 0H8.818M12 15v2.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YenOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132zm7.377 9.031H12m0 0H8.818m3.182 0l-3.5-5m3.5 5l3.5-5m-3.5 5V15m3.182 0H12m0 0H8.818M12 15v2.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YenSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12m12.182.5H12m0 0H8.818m3.182 0l-3.5-5m3.5 5l3.5-5m-3.5 5V15m3.182 0H12m0 0H8.818M12 15v2.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YenWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83m5.469 8.859H12m0 0H8.818m3.182 0l-3.5-5m3.5 5l3.5-5m-3.5 5V15m3.182 0H12m0 0H8.818M12 15v2.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Zero(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.75 13.875v-3.75a3.125 3.125 0 1 1 6.25 0v3.75a3.125 3.125 0 1 1-6.25 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZeroCircle(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M9.5 13.5v-3a2.5 2.5 0 0 1 5 0v3a2.5 2.5 0 0 1-5 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZeroDiamond(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.5 13.5v-3a2.5 2.5 0 0 1 5 0v3a2.5 2.5 0 0 1-5 0"/><path d="M2.707 10.295a2.41 2.41 0 0 0 0 3.41l7.588 7.588a2.409 2.409 0 0 0 3.41 0l7.588-7.588a2.409 2.409 0 0 0 0-3.41l-7.588-7.588a2.41 2.41 0 0 0-3.41 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZeroHexagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.5 13.5v-3a2.5 2.5 0 0 1 5 0v3a2.5 2.5 0 0 1-5 0"/><path d="M20.5 15.8V8.2a1.91 1.91 0 0 0-.944-1.645l-6.612-3.8a1.88 1.88 0 0 0-1.888 0l-6.612 3.8A1.895 1.895 0 0 0 3.5 8.2v7.602a1.91 1.91 0 0 0 .944 1.644l6.612 3.8a1.88 1.88 0 0 0 1.888 0l6.612-3.8A1.895 1.895 0 0 0 20.5 15.8"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZeroOctagon(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.5 13.5v-3a2.5 2.5 0 0 1 5 0v3a2.5 2.5 0 0 1-5 0"/><path d="M7.805 3.469C8.16 3.115 8.451 3 8.937 3h6.126c.486 0 .778.115 1.132.469l4.336 4.336c.354.354.469.646.469 1.132v6.126c0 .5-.125.788-.469 1.132l-4.336 4.336c-.354.354-.646.469-1.132.469H8.937c-.5 0-.788-.125-1.132-.469L3.47 16.195c-.355-.355-.47-.646-.47-1.132V8.937c0-.5.125-.788.469-1.132z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZeroSquare(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.5 13.5v-3a2.5 2.5 0 0 1 5 0v3a2.5 2.5 0 0 1-5 0"/><path d="M3 12c0-4.243 0-6.364 1.318-7.682C5.636 3 7.758 3 12 3c4.243 0 6.364 0 7.682 1.318C21 5.636 21 7.758 21 12c0 4.243 0 6.364-1.318 7.682C18.364 21 16.242 21 12 21c-4.243 0-6.364 0-7.682-1.318C3 18.364 3 16.242 3 12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZeroWaves(children ...ElementRenderer) *MynauiIcon {
	return &MynauiIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.5 13.5v-3a2.5 2.5 0 0 1 5 0v3a2.5 2.5 0 0 1-5 0"/><path d="M9.713 3.64c.581-.495.872-.743 1.176-.888a2.577 2.577 0 0 1 2.222 0c.304.145.595.393 1.176.888c.599.51 1.207.768 2.007.831c.761.061 1.142.092 1.46.204c.734.26 1.312.837 1.571 1.572c.112.317.143.698.204 1.46c.063.8.32 1.407.83 2.006c.496.581.744.872.889 1.176c.336.703.336 1.52 0 2.222c-.145.304-.393.595-.888 1.176a3.306 3.306 0 0 0-.831 2.007c-.061.761-.092 1.142-.204 1.46a2.577 2.577 0 0 1-1.572 1.571c-.317.112-.698.143-1.46.204c-.8.063-1.407.32-2.006.83c-.581.496-.872.744-1.176.889a2.577 2.577 0 0 1-2.222 0c-.304-.145-.595-.393-1.176-.888a3.306 3.306 0 0 0-2.007-.831c-.761-.061-1.142-.092-1.46-.204a2.577 2.577 0 0 1-1.571-1.572c-.112-.317-.143-.698-.204-1.46a3.305 3.305 0 0 0-.83-2.006c-.496-.581-.744-.872-.89-1.176a2.577 2.577 0 0 1 .001-2.222c.145-.304.393-.595.888-1.176c.52-.611.769-1.223.831-2.007c.061-.761.092-1.142.204-1.46a2.577 2.577 0 0 1 1.572-1.571c.317-.112.698-.143 1.46-.204a3.305 3.305 0 0 0 2.006-.83"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
