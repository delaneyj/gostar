package circum

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = "1.0.0"
	hAttr          = "1em"
	viewbox        = "0 0 24 24"
	fill           = "currentColor"
)

type CircumIcon struct {
	*SVGSVGElement
}

type CircumIconFn func(children ...ElementRenderer) *CircumIcon

var IconLookup = map[string]CircumIconFn{
	"airportSignOne":   AirportSignOne,
	"alarmOff":         AlarmOff,
	"alarmOn":          AlarmOn,
	"alignBottom":      AlignBottom,
	"alignCenterH":     AlignCenterH,
	"alignCenterV":     AlignCenterV,
	"alignLeft":        AlignLeft,
	"alignRight":       AlignRight,
	"alignTop":         AlignTop,
	"apple":            Apple,
	"at":               At,
	"avocado":          Avocado,
	"bacon":            Bacon,
	"badgeDollar":      BadgeDollar,
	"bagOne":           BagOne,
	"bandage":          Bandage,
	"bank":             Bank,
	"barcode":          Barcode,
	"baseball":         Baseball,
	"basketball":       Basketball,
	"batteryCharging":  BatteryCharging,
	"batteryEmpty":     BatteryEmpty,
	"batteryFull":      BatteryFull,
	"beakerOne":        BeakerOne,
	"beerMugFull":      BeerMugFull,
	"bellOff":          BellOff,
	"bellOn":           BellOn,
	"bezier":           Bezier,
	"bitcoin":          Bitcoin,
	"bluetooth":        Bluetooth,
	"bookmark":         Bookmark,
	"bookmarkCheck":    BookmarkCheck,
	"bookmarkMinus":    BookmarkMinus,
	"bookmarkPlus":     BookmarkPlus,
	"bookmarkRemove":   BookmarkRemove,
	"bowlNoodles":      BowlNoodles,
	"boxList":          BoxList,
	"boxes":            Boxes,
	"brightnessDown":   BrightnessDown,
	"brightnessUp":     BrightnessUp,
	"bullhorn":         Bullhorn,
	"burger":           Burger,
	"calculatorOne":    CalculatorOne,
	"calculatorTwo":    CalculatorTwo,
	"calendar":         Calendar,
	"calendarDate":     CalendarDate,
	"camera":           Camera,
	"chatOne":          ChatOne,
	"chatTwo":          ChatTwo,
	"circleAlert":      CircleAlert,
	"circleCheck":      CircleCheck,
	"circleChevDown":   CircleChevDown,
	"circleChevLeft":   CircleChevLeft,
	"circleChevRight":  CircleChevRight,
	"circleChevUp":     CircleChevUp,
	"circleInfo":       CircleInfo,
	"circleList":       CircleList,
	"circleMinus":      CircleMinus,
	"circleMore":       CircleMore,
	"circlePlus":       CirclePlus,
	"circleQuestion":   CircleQuestion,
	"circleRemove":     CircleRemove,
	"clockOne":         ClockOne,
	"clockTwo":         ClockTwo,
	"cloud":            Cloud,
	"cloudDrizzle":     CloudDrizzle,
	"cloudMoon":        CloudMoon,
	"cloudOff":         CloudOff,
	"cloudOn":          CloudOn,
	"cloudRainbow":     CloudRainbow,
	"cloudSun":         CloudSun,
	"coffeeBean":       CoffeeBean,
	"coffeeCup":        CoffeeCup,
	"coinInsert":       CoinInsert,
	"coinsOne":         CoinsOne,
	"compassOne":       CompassOne,
	"creditCardOff":    CreditCardOff,
	"creditCardOne":    CreditCardOne,
	"creditCardTwo":    CreditCardTwo,
	"crop":             Crop,
	"dark":             Dark,
	"database":         Database,
	"deliveryTruck":    DeliveryTruck,
	"desktop":          Desktop,
	"desktopMouseOne":  DesktopMouseOne,
	"desktopMouseTwo":  DesktopMouseTwo,
	"discountOne":      DiscountOne,
	"dollar":           Dollar,
	"droplet":          Droplet,
	"dumbbell":         Dumbbell,
	"edit":             Edit,
	"eraser":           Eraser,
	"export":           Export,
	"faceFrown":        FaceFrown,
	"faceMeh":          FaceMeh,
	"faceSmile":        FaceSmile,
	"facebook":         Facebook,
	"fileOff":          FileOff,
	"fileOn":           FileOn,
	"filter":           Filter,
	"flagOne":          FlagOne,
	"floppyDisk":       FloppyDisk,
	"folderOff":        FolderOff,
	"folderOn":         FolderOn,
	"football":         Football,
	"forkKnife":        ForkKnife,
	"fries":            Fries,
	"gift":             Gift,
	"glass":            Glass,
	"globe":            Globe,
	"gps":              Gps,
	"gridFourOne":      GridFourOne,
	"gridFourTwo":      GridFourTwo,
	"gridThreeOne":     GridThreeOne,
	"gridThreeTwo":     GridThreeTwo,
	"gridTwoH":         GridTwoH,
	"gridTwoV":         GridTwoV,
	"hardDrive":        HardDrive,
	"hashtag":          Hashtag,
	"headphones":       Headphones,
	"heart":            Heart,
	"home":             Home,
	"hospitalOne":      HospitalOne,
	"hotdog":           Hotdog,
	"iceCream":         IceCream,
	"imageOff":         ImageOff,
	"imageOn":          ImageOn,
	"import":           Import,
	"inboxIn":          InboxIn,
	"inboxOut":         InboxOut,
	"indent":           Indent,
	"instagram":        Instagram,
	"keyboard":         Keyboard,
	"laptop":           Laptop,
	"lemon":            Lemon,
	"light":            Light,
	"lineHeight":       LineHeight,
	"link":             Link,
	"linkedin":         Linkedin,
	"locationArrowOne": LocationArrowOne,
	"locationOff":      LocationOff,
	"locationOn":       LocationOn,
	"lock":             Lock,
	"login":            Login,
	"logout":           Logout,
	"lollipop":         Lollipop,
	"mail":             Mail,
	"map":              Map,
	"mapPin":           MapPin,
	"maximizeOne":      MaximizeOne,
	"maximizeTwo":      MaximizeTwo,
	"medal":            Medal,
	"medicalCase":      MedicalCase,
	"medicalClipboard": MedicalClipboard,
	"medicalCross":     MedicalCross,
	"medicalMask":      MedicalMask,
	"memoPad":          MemoPad,
	"menuBurger":       MenuBurger,
	"menuFries":        MenuFries,
	"menuKebab":        MenuKebab,
	"microchip":        Microchip,
	"microphoneOff":    MicrophoneOff,
	"microphoneOn":     MicrophoneOn,
	"minimizeOne":      MinimizeOne,
	"minimizeTwo":      MinimizeTwo,
	"mobileFour":       MobileFour,
	"mobileOne":        MobileOne,
	"mobileThree":      MobileThree,
	"mobileTwo":        MobileTwo,
	"moneyBill":        MoneyBill,
	"moneyCheckOne":    MoneyCheckOne,
	"monitor":          Monitor,
	"mountainOne":      MountainOne,
	"mugOne":           MugOne,
	"musicNoteOne":     MusicNoteOne,
	"noWaitingSign":    NoWaitingSign,
	"palette":          Palette,
	"paperplane":       Paperplane,
	"parkingOne":       ParkingOne,
	"passportOne":      PassportOne,
	"pauseOne":         PauseOne,
	"pen":              Pen,
	"penpot":           Penpot,
	"percent":          Percent,
	"phone":            Phone,
	"pickerEmpty":      PickerEmpty,
	"pickerHalf":       PickerHalf,
	"pill":             Pill,
	"pillsBottleOne":   PillsBottleOne,
	"pizza":            Pizza,
	"plane":            Plane,
	"playOne":          PlayOne,
	"plugOne":          PlugOne,
	"power":            Power,
	"rainbow":          Rainbow,
	"read":             Read,
	"receipt":          Receipt,
	"redo":             Redo,
	"repeat":           Repeat,
	"rollingSuitcase":  RollingSuitcase,
	"route":            Route,
	"router":           Router,
	"ruler":            Ruler,
	"satelliteOne":     SatelliteOne,
	"saveDownOne":      SaveDownOne,
	"saveDownTwo":      SaveDownTwo,
	"saveUpOne":        SaveUpOne,
	"saveUpTwo":        SaveUpTwo,
	"search":           Search,
	"server":           Server,
	"settings":         Settings,
	"shareOne":         ShareOne,
	"shareTwo":         ShareTwo,
	"shirt":            Shirt,
	"shop":             Shop,
	"shoppingBasket":   ShoppingBasket,
	"shoppingCart":     ShoppingCart,
	"shoppingTag":      ShoppingTag,
	"shuffle":          Shuffle,
	"signpostDuoOne":   SignpostDuoOne,
	"signpostLone":     SignpostLone,
	"signpostRone":     SignpostRone,
	"sliderHorizontal": SliderHorizontal,
	"sliderVertical":   SliderVertical,
	"speaker":          Speaker,
	"squareAlert":      SquareAlert,
	"squareCheck":      SquareCheck,
	"squareChevDown":   SquareChevDown,
	"squareChevLeft":   SquareChevLeft,
	"squareChevRight":  SquareChevRight,
	"squareChevUp":     SquareChevUp,
	"squareInfo":       SquareInfo,
	"squareMinus":      SquareMinus,
	"squareMore":       SquareMore,
	"squarePlus":       SquarePlus,
	"squareQuestion":   SquareQuestion,
	"squareRemove":     SquareRemove,
	"star":             Star,
	"stethoscope":      Stethoscope,
	"stickyNote":       StickyNote,
	"stopOne":          StopOne,
	"stopSignOne":      StopSignOne,
	"stopwatch":        Stopwatch,
	"streamOff":        StreamOff,
	"streamOn":         StreamOn,
	"sun":              Sun,
	"tabletsOne":       TabletsOne,
	"tempHigh":         TempHigh,
	"textAlignCenter":  TextAlignCenter,
	"textAlignJustify": TextAlignJustify,
	"textAlignLeft":    TextAlignLeft,
	"textAlignRight":   TextAlignRight,
	"textIcon":         TextIcon,
	"timer":            Timer,
	"trash":            Trash,
	"trophy":           Trophy,
	"turnLone":         TurnLone,
	"turnRone":         TurnRone,
	"twitter":          Twitter,
	"umbrella":         Umbrella,
	"undo":             Undo,
	"unlock":           Unlock,
	"unread":           Unread,
	"usb":              Usb,
	"user":             User,
	"vault":            Vault,
	"vial":             Vial,
	"videoOff":         VideoOff,
	"videoOn":          VideoOn,
	"viewBoard":        ViewBoard,
	"viewColumn":       ViewColumn,
	"viewList":         ViewList,
	"viewTable":        ViewTable,
	"viewTimeline":     ViewTimeline,
	"virus":            Virus,
	"voicemail":        Voicemail,
	"volume":           Volume,
	"volumeHigh":       VolumeHigh,
	"volumeMute":       VolumeMute,
	"wallet":           Wallet,
	"warning":          Warning,
	"wavePulseOne":     WavePulseOne,
	"wheat":            Wheat,
	"wifiOff":          WifiOff,
	"wifiOn":           WifiOn,
	"youtube":          Youtube,
	"zoomIn":           ZoomIn,
	"zoomOut":          ZoomOut,
}

func AirportSignOne(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.105 12.609v2.279a.119.119 0 0 0 .061.105l.622.355a.49.49 0 0 1 .242.365l.049.413a.243.243 0 0 1-.307.263l-1.641-.459a.486.486 0 0 0-.262 0l-1.641.459a.244.244 0 0 1-.308-.263l.05-.413a.487.487 0 0 1 .242-.365l.621-.355a.12.12 0 0 0 .062-.105v-2.279a.122.122 0 0 0-.137-.121l-3.485.435A.242.242 0 0 1 7 12.682v-.624a.486.486 0 0 1 .316-.455l3.5-1.313a.122.122 0 0 0 .079-.114v-.741a4.756 4.756 0 0 1 .1-.981a1.015 1.015 0 0 1 1.2-.833a1.063 1.063 0 0 1 .819.9l.015.094a6.3 6.3 0 0 1 .077.976v.587a.121.121 0 0 0 .079.114l3.5 1.313a.486.486 0 0 1 .316.455v.624a.243.243 0 0 1-.274.241l-3.484-.435a.121.121 0 0 0-.138.119"/><path fill="currentColor" d="M12 21.933A9.933 9.933 0 1 1 21.933 12A9.944 9.944 0 0 1 12 21.933m0-18.866A8.933 8.933 0 1 0 20.933 12A8.943 8.943 0 0 0 12 3.067"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlarmOff(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.564 3.856a.5.5 0 0 0-.7.71l.29.29l-.5.5a2.019 2.019 0 0 0-.01 2.85l.65.67a8.273 8.273 0 0 0-.71 3.39A8.427 8.427 0 0 0 12 20.686a8.275 8.275 0 0 0 5.72-2.26c.57.57 1.14 1.15 1.71 1.71a.5.5 0 0 0 .71-.7Zm-.21 2.21l.51-.5c.32.33.65.65.98.98a6.38 6.38 0 0 0-1.06 1.4l-.43-.44a1.032 1.032 0 0 1 0-1.44M12 19.686a7.43 7.43 0 0 1-7.42-7.42a7.312 7.312 0 0 1 1.96-5.02l2.59 2.59l7.88 7.88a7.27 7.27 0 0 1-5.01 1.97m8.354-11.47a2.04 2.04 0 0 0 0-2.86l-1.46-1.45a2.01 2.01 0 0 0-2.85 0l-.68.67a8.528 8.528 0 0 0-6.38-.17c-.6.23-.34 1.19.27.97a7.419 7.419 0 0 1 9.64 9.64c-.22.6.74.86.97.26a8.506 8.506 0 0 0-.17-6.39Zm-2.4-1.9a8.068 8.068 0 0 0-1.65-1.27l.44-.43a1.026 1.026 0 0 1 1.45 0l1.45 1.45a1.014 1.014 0 0 1 0 1.44l-.43.44a8.262 8.262 0 0 0-1.26-1.63"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlarmOn(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.352 8.213a2.017 2.017 0 0 0 0-2.851L18.9 3.9a2.038 2.038 0 0 0-2.86 0l-.67.659A8.238 8.238 0 0 0 12 3.852a8.332 8.332 0 0 0-3.39.71L7.962 3.9a2.038 2.038 0 0 0-2.86 0l-1.45 1.462a2.02 2.02 0 0 0-.01 2.851l.65.67a8.419 8.419 0 1 0 16.13 3.39a8.4 8.4 0 0 0-.72-3.411ZM4.362 6.062l1.45-1.45a1.016 1.016 0 0 1 1.44 0l.44.43a8.427 8.427 0 0 0-2.91 2.9l-.42-.43a1.027 1.027 0 0 1 0-1.45M12 19.682a7.415 7.415 0 1 1 7.42-7.409A7.421 7.421 0 0 1 12 19.682m7.22-11.75a8.578 8.578 0 0 0-2.91-2.89l.44-.43a1.016 1.016 0 0 1 1.44 0l1.45 1.45a1.027 1.027 0 0 1 0 1.451Z"/><path fill="currentColor" d="M17.042 12.763H12a.455.455 0 0 1-.27-.081c-.03-.02-.05-.039-.07-.049a.442.442 0 0 1-.16-.36V7.232a.5.5 0 0 1 1 0v4.531h4.54a.5.5 0 0 1 .002 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignBottom(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.548 20.922h16.9a.5.5 0 0 0 0-1h-16.9a.5.5 0 0 0 0 1M9 18.919H6.565a2.5 2.5 0 0 1-2.5-2.5V5.578a2.5 2.5 0 0 1 2.5-2.5H9a2.5 2.5 0 0 1 2.5 2.5v10.841a2.5 2.5 0 0 1-2.5 2.5M6.565 4.078a1.5 1.5 0 0 0-1.5 1.5v10.841a1.5 1.5 0 0 0 1.5 1.5H9a1.5 1.5 0 0 0 1.5-1.5V5.578a1.5 1.5 0 0 0-1.5-1.5Zm10.872 14.841H15a2.5 2.5 0 0 1-2.5-2.5V10.55a2.5 2.5 0 0 1 2.5-2.5h2.434a2.5 2.5 0 0 1 2.5 2.5v5.869a2.5 2.5 0 0 1-2.497 2.5M15 9.05a1.5 1.5 0 0 0-1.5 1.5v5.869a1.5 1.5 0 0 0 1.5 1.5h2.434a1.5 1.5 0 0 0 1.5-1.5V10.55a1.5 1.5 0 0 0-1.5-1.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignCenterH(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.42 4.062H12.5v-.51a.5.5 0 0 0-1 0v.51H6.58a2.507 2.507 0 0 0-2.5 2.5V9a2.5 2.5 0 0 0 2.5 2.5h4.92v1H9.06a2.507 2.507 0 0 0-2.5 2.5v2.44a2.507 2.507 0 0 0 2.5 2.5h2.44v.51a.5.5 0 0 0 1 0v-.51h2.43a2.5 2.5 0 0 0 2.5-2.5V15a2.5 2.5 0 0 0-2.5-2.5H12.5v-1h4.92a2.5 2.5 0 0 0 2.5-2.5V6.562a2.507 2.507 0 0 0-2.5-2.5m-5.92 14.88H9.06a1.511 1.511 0 0 1-1.5-1.5V15a1.5 1.5 0 0 1 1.5-1.5h2.44Zm0-8.44H6.58A1.5 1.5 0 0 1 5.08 9V6.562a1.5 1.5 0 0 1 1.5-1.5h4.92Zm3.43 3a1.5 1.5 0 0 1 1.5 1.5v2.44a1.5 1.5 0 0 1-1.5 1.5H12.5V13.5ZM18.92 9a1.5 1.5 0 0 1-1.5 1.5H12.5V5.062h4.92a1.5 1.5 0 0 1 1.5 1.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignCenterV(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.446 11.5h-.51V9.07a2.5 2.5 0 0 0-2.5-2.5h-2.43a2.5 2.5 0 0 0-2.5 2.5v2.43H11.5V6.58A2.5 2.5 0 0 0 9 4.08H6.566a2.5 2.5 0 0 0-2.5 2.5v4.92h-.52a.5.5 0 0 0 0 1h.52v4.92a2.5 2.5 0 0 0 2.5 2.5H9a2.5 2.5 0 0 0 2.5-2.5V12.5h1.01v2.43a2.5 2.5 0 0 0 2.5 2.5h2.43a2.5 2.5 0 0 0 2.5-2.5V12.5h.51a.5.5 0 0 0-.004-1M10.5 17.42a1.5 1.5 0 0 1-1.5 1.5H6.566a1.5 1.5 0 0 1-1.5-1.5V12.5H10.5Zm0-5.92H5.066V6.58a1.5 1.5 0 0 1 1.5-1.5H9a1.5 1.5 0 0 1 1.5 1.5Zm8.44 3.43a1.5 1.5 0 0 1-1.5 1.5h-2.43a1.5 1.5 0 0 1-1.5-1.5V12.5h5.43Zm0-3.43h-5.43V9.07a1.5 1.5 0 0 1 1.5-1.5h2.43a1.5 1.5 0 0 1 1.5 1.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignLeft(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.078 3.548v16.9a.5.5 0 0 0 1 0v-16.9a.5.5 0 0 0-1 0M18.422 11.5H7.582a2.5 2.5 0 0 1-2.5-2.5V6.565a2.5 2.5 0 0 1 2.5-2.5h10.84a2.5 2.5 0 0 1 2.5 2.5V9a2.5 2.5 0 0 1-2.5 2.5M7.582 5.065a1.5 1.5 0 0 0-1.5 1.5V9a1.5 1.5 0 0 0 1.5 1.5h10.84a1.5 1.5 0 0 0 1.5-1.5V6.565a1.5 1.5 0 0 0-1.5-1.5Zm5.869 14.873H7.582a2.5 2.5 0 0 1-2.5-2.5V15a2.5 2.5 0 0 1 2.5-2.5h5.869a2.5 2.5 0 0 1 2.5 2.5v2.436a2.5 2.5 0 0 1-2.5 2.502M7.582 13.5a1.5 1.5 0 0 0-1.5 1.5v2.436a1.5 1.5 0 0 0 1.5 1.5h5.869a1.5 1.5 0 0 0 1.5-1.5V15a1.5 1.5 0 0 0-1.5-1.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignRight(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.922 3.548v16.9a.5.5 0 0 0 1 0v-16.9a.5.5 0 0 0-1 0M16.419 11.5H5.578a2.5 2.5 0 0 1-2.5-2.5V6.565a2.5 2.5 0 0 1 2.5-2.5h10.841a2.5 2.5 0 0 1 2.5 2.5V9a2.5 2.5 0 0 1-2.5 2.5M5.578 5.065a1.5 1.5 0 0 0-1.5 1.5V9a1.5 1.5 0 0 0 1.5 1.5h10.841a1.5 1.5 0 0 0 1.5-1.5V6.565a1.5 1.5 0 0 0-1.5-1.5Zm10.841 14.873H10.55a2.5 2.5 0 0 1-2.5-2.5V15a2.5 2.5 0 0 1 2.5-2.5h5.869a2.5 2.5 0 0 1 2.5 2.5v2.436a2.5 2.5 0 0 1-2.5 2.502M10.55 13.5a1.5 1.5 0 0 0-1.5 1.5v2.436a1.5 1.5 0 0 0 1.5 1.5h5.869a1.5 1.5 0 0 0 1.5-1.5V15a1.5 1.5 0 0 0-1.5-1.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignTop(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.548 4.078h16.9a.5.5 0 0 0 0-1h-16.9a.5.5 0 0 0 0 1M9 20.922H6.565a2.5 2.5 0 0 1-2.5-2.5V7.582a2.5 2.5 0 0 1 2.5-2.5H9a2.5 2.5 0 0 1 2.5 2.5v10.84a2.5 2.5 0 0 1-2.5 2.5M6.565 6.082a1.5 1.5 0 0 0-1.5 1.5v10.84a1.5 1.5 0 0 0 1.5 1.5H9a1.5 1.5 0 0 0 1.5-1.5V7.582a1.5 1.5 0 0 0-1.5-1.5Zm10.873 9.869H15a2.5 2.5 0 0 1-2.5-2.5V7.582a2.5 2.5 0 0 1 2.5-2.5h2.435a2.5 2.5 0 0 1 2.5 2.5v5.869a2.5 2.5 0 0 1-2.497 2.5M15 6.082a1.5 1.5 0 0 0-1.5 1.5v5.869a1.5 1.5 0 0 0 1.5 1.5h2.435a1.5 1.5 0 0 0 1.5-1.5V7.582a1.5 1.5 0 0 0-1.5-1.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Apple(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14.875 6.612l.05-.05a3.229 3.229 0 0 0 .95-2.58a.976.976 0 0 0-.9-.9a3.229 3.229 0 0 0-2.58.95a3.279 3.279 0 0 0-.85 1.46a4.661 4.661 0 0 0-2.69-1.75a.5.5 0 1 0-.22.98a3.664 3.664 0 0 1 2.59 2.2a5.577 5.577 0 0 0-1.9-.32a5.847 5.847 0 0 0-5.84 5.84c0 2.98 2.41 8.49 5.84 8.49a5.821 5.821 0 0 0 2.4-.52a.683.683 0 0 1 .56 0a5.73 5.73 0 0 0 2.38.52c3.44 0 5.85-5.51 5.85-8.49a5.838 5.838 0 0 0-5.64-5.83m-1.77-1.87a2.3 2.3 0 0 1 1.78-.68c0 .06.01.12.01.17a2.326 2.326 0 0 1-.67 1.63a2.359 2.359 0 0 1-1.79.66a2.247 2.247 0 0 1 .67-1.78m1.56 15.19a4.787 4.787 0 0 1-1.97-.43a1.718 1.718 0 0 0-.69-.15a1.649 1.649 0 0 0-.69.15a4.879 4.879 0 0 1-1.99.43c-2.58 0-4.84-4.67-4.84-7.49a4.855 4.855 0 0 1 6.83-4.42a1.56 1.56 0 0 0 .67.15h.02a1.683 1.683 0 0 0 .69-.15a4.777 4.777 0 0 1 1.97-.42a4.852 4.852 0 0 1 4.85 4.84c0 2.82-2.27 7.49-4.85 7.49"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func At(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.09 21.925a9.846 9.846 0 0 1-3.838-.747a9.673 9.673 0 0 1-5.247-5.248A10.034 10.034 0 0 1 2.244 12a10.425 10.425 0 0 1 .695-3.8a9.606 9.606 0 0 1 2-3.169A9.269 9.269 0 0 1 8.1 2.862a10.605 10.605 0 0 1 4.175-.787a10.516 10.516 0 0 1 4.334.827a8.437 8.437 0 0 1 3.031 2.217a8.622 8.622 0 0 1 1.707 3.1a9.263 9.263 0 0 1 .377 3.487a5.809 5.809 0 0 1-1.3 3.6a3.6 3.6 0 0 1-2.724 1.167a3.628 3.628 0 0 1-2.162-.609a2.82 2.82 0 0 1-1.119-1.694l.5.106a2.582 2.582 0 0 1-1.3 1.3a4.37 4.37 0 0 1-1.873.424a3.681 3.681 0 0 1-1.866-.46a3.2 3.2 0 0 1-1.237-1.271A3.843 3.843 0 0 1 8.2 12.4a3.88 3.88 0 0 1 .456-1.926a3.191 3.191 0 0 1 1.263-1.26a3.792 3.792 0 0 1 1.853-.443a4.716 4.716 0 0 1 1.767.364a2.622 2.622 0 0 1 1.383 1.3l-.5.5V9.461a.4.4 0 0 1 .4-.4h.232a.4.4 0 0 1 .4.4v3.518a2.723 2.723 0 0 0 .529 1.674a2.173 2.173 0 0 0 1.853.708a2.281 2.281 0 0 0 1.323-.41a2.938 2.938 0 0 0 .967-1.178a4.947 4.947 0 0 0 .437-1.852a9.439 9.439 0 0 0-.417-3.574A7.285 7.285 0 0 0 18.5 5.588a7.424 7.424 0 0 0-2.679-1.78a9.605 9.605 0 0 0-3.547-.622a9.041 9.041 0 0 0-3.758.741a8.252 8.252 0 0 0-2.773 2a8.8 8.8 0 0 0-1.72 2.838a9.27 9.27 0 0 0-.589 3.262a8.568 8.568 0 0 0 .682 3.408A8.951 8.951 0 0 0 6 18.24a8.707 8.707 0 0 0 2.785 1.892a8.515 8.515 0 0 0 3.389.682a9.851 9.851 0 0 0 2.679-.378a8.451 8.451 0 0 0 2-.831a.4.4 0 0 1 .553.158l.1.192a.4.4 0 0 1-.141.526a9.832 9.832 0 0 1-2.391 1.04a10.5 10.5 0 0 1-2.884.404m-.29-7.066a2.469 2.469 0 0 0 1.786-.649a2.427 2.427 0 0 0 .675-1.839a2.414 2.414 0 0 0-.7-1.886a2.532 2.532 0 0 0-1.761-.629a2.482 2.482 0 0 0-1.839.649a2.523 2.523 0 0 0-.65 1.866a2.4 2.4 0 0 0 .682 1.865a2.574 2.574 0 0 0 1.807.623"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Avocado(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.982 21.939a7.759 7.759 0 0 1-.818-.044A7.58 7.58 0 0 1 6.1 9.6a3.578 3.578 0 0 0 .684-2.271a5.128 5.128 0 0 1 3.8-5.085a5.266 5.266 0 0 1 4.6.892a5.185 5.185 0 0 1 2.039 4.14A3.6 3.6 0 0 0 17.9 9.61a7.574 7.574 0 0 1-5.918 12.329m.009-18.877a4.538 4.538 0 0 0-1.158.152a4.126 4.126 0 0 0-3.055 4.07a4.532 4.532 0 0 1-.9 2.947a6.555 6.555 0 0 0-1.366 5.231a6.643 6.643 0 0 0 5.759 5.438a6.575 6.575 0 0 0 5.851-10.662a4.453 4.453 0 0 1-.9-2.9a4.214 4.214 0 0 0-4.228-4.273Z"/><ellipse cx="11.999" cy="14.856" fill="currentColor" rx="2.5" ry="3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bacon(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.605 7.455l-3.49-3.49a.8.8 0 0 0-1.08-.04a1.833 1.833 0 0 1-.93.37a3.787 3.787 0 0 0-2.21 1.12a3.918 3.918 0 0 0-1.13 2.22a1.681 1.681 0 0 1-.53 1.1a1.753 1.753 0 0 1-1.1.53a4.026 4.026 0 0 0-3.35 3.35a1.677 1.677 0 0 1-.53 1.1a1.721 1.721 0 0 1-1.11.53a4.041 4.041 0 0 0-1.62.63a1.1 1.1 0 0 0-.14 1.66l3.5 3.5a.781.781 0 0 0 .55.23a.822.822 0 0 0 .53-.19a1.759 1.759 0 0 1 .93-.38a3.8 3.8 0 0 0 2.21-1.12a3.948 3.948 0 0 0 1.14-2.22a1.71 1.71 0 0 1 .52-1.1a1.776 1.776 0 0 1 1.11-.53a4.03 4.03 0 0 0 3.34-3.35a1.66 1.66 0 0 1 .53-1.1a1.721 1.721 0 0 1 1.11-.53a4.018 4.018 0 0 0 1.61-.62a1.091 1.091 0 0 0 .14-1.67M4.1 15.7a3.15 3.15 0 0 1 1.24-.47a2.635 2.635 0 0 0 1.63-.81a2.587 2.587 0 0 0 .8-1.61a2.852 2.852 0 0 1 .86-1.7a2.9 2.9 0 0 1 1.7-.86a2.745 2.745 0 0 0 1.62-.8a2.687 2.687 0 0 0 .8-1.62a2.9 2.9 0 0 1 .86-1.7a2.814 2.814 0 0 1 1.69-.85a2.819 2.819 0 0 0 1.24-.48l1.3 1.3a2.362 2.362 0 0 1-.98.35a3.515 3.515 0 0 0-2.95 2.95a2.136 2.136 0 0 1-.67 1.36a2.159 2.159 0 0 1-1.36.67a3.44 3.44 0 0 0-1.96.99a3.351 3.351 0 0 0-.98 1.96a2.355 2.355 0 0 1-2.03 2.03a3.242 3.242 0 0 0-1.58.66L4.1 15.835Zm15.8-7.4a3.059 3.059 0 0 1-1.23.47a2.659 2.659 0 0 0-1.63.81a2.587 2.587 0 0 0-.8 1.61a2.852 2.852 0 0 1-.86 1.7a2.883 2.883 0 0 1-1.69.86a2.812 2.812 0 0 0-2.43 2.42a2.878 2.878 0 0 1-.86 1.7a2.8 2.8 0 0 1-1.68.85a2.808 2.808 0 0 0-1.25.48l-1.3-1.29a2.423 2.423 0 0 1 .97-.35a3.377 3.377 0 0 0 1.96-.99a3.44 3.44 0 0 0 .99-1.96a2.217 2.217 0 0 1 .66-1.36a2.292 2.292 0 0 1 1.36-.67a3.317 3.317 0 0 0 1.96-.99a3.351 3.351 0 0 0 .99-1.96a2.355 2.355 0 0 1 2.03-2.03a3.479 3.479 0 0 0 1.59-.66l1.23 1.23Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BadgeDollar(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21.953c-.895 0-1.545-.743-2.118-1.4a3.671 3.671 0 0 0-1.033-.946a3.8 3.8 0 0 0-1.466-.077a3.012 3.012 0 0 1-2.421-.494a3.014 3.014 0 0 1-.494-2.421a3.82 3.82 0 0 0-.077-1.466a3.671 3.671 0 0 0-.946-1.033c-.655-.573-1.4-1.222-1.4-2.118s.743-1.545 1.4-2.118a3.66 3.66 0 0 0 .946-1.034a3.815 3.815 0 0 0 .077-1.465a3.012 3.012 0 0 1 .494-2.421a3.015 3.015 0 0 1 2.422-.5a3.794 3.794 0 0 0 1.465-.07a3.666 3.666 0 0 0 1.033-.945c.573-.655 1.223-1.4 2.118-1.4s1.545.742 2.118 1.4a3.66 3.66 0 0 0 1.034.946a3.807 3.807 0 0 0 1.464.077a3.018 3.018 0 0 1 2.422.5a3.012 3.012 0 0 1 .5 2.422a3.81 3.81 0 0 0 .077 1.464a3.66 3.66 0 0 0 .946 1.034c.655.573 1.4 1.223 1.4 2.118s-.743 1.545-1.4 2.118a3.666 3.666 0 0 0-.945 1.033a3.815 3.815 0 0 0-.077 1.465a3.012 3.012 0 0 1-.5 2.422a3.018 3.018 0 0 1-2.421.494a3.818 3.818 0 0 0-1.465.077a3.673 3.673 0 0 0-1.034.946c-.574.649-1.219 1.392-2.119 1.392M8.093 18.5a2.952 2.952 0 0 1 1.138.183a4.233 4.233 0 0 1 1.4 1.21c.454.52.924 1.057 1.365 1.057s.911-.537 1.366-1.057a4.225 4.225 0 0 1 1.4-1.21a4.365 4.365 0 0 1 1.908-.152c.672.041 1.366.085 1.653-.2s.245-.982.2-1.653a4.387 4.387 0 0 1 .152-1.909a4.241 4.241 0 0 1 1.209-1.4c.52-.454 1.057-.924 1.057-1.365s-.537-.911-1.057-1.365a4.234 4.234 0 0 1-1.209-1.4a4.381 4.381 0 0 1-.152-1.908c.041-.671.084-1.365-.2-1.653s-.982-.246-1.653-.2a4.384 4.384 0 0 1-1.908-.152a4.234 4.234 0 0 1-1.4-1.209c-.454-.52-.924-1.057-1.365-1.057s-.911.537-1.365 1.057a4.241 4.241 0 0 1-1.4 1.209a4.417 4.417 0 0 1-1.909.152c-.67-.041-1.364-.084-1.653.2s-.244.981-.2 1.652a4.37 4.37 0 0 1-.156 1.9a4.226 4.226 0 0 1-1.21 1.4c-.52.454-1.057.925-1.057 1.365s.537.911 1.057 1.366a4.238 4.238 0 0 1 1.21 1.4a4.378 4.378 0 0 1 .152 1.91c-.041.672-.084 1.366.2 1.653s.98.245 1.653.2c.259-.005.519-.024.774-.024"/><path fill="currentColor" d="M14.5 13.5a2.006 2.006 0 0 1-2 2v1.01a.5.5 0 0 1-.5.49a.492.492 0 0 1-.5-.49V15.5h-1.25a.5.5 0 0 1-.5-.5a.5.5 0 0 1 .5-.5h2.25a1 1 0 1 0 0-2h-1a2 2 0 0 1 0-4V7.453A.473.473 0 0 1 12 7a.48.48 0 0 1 .5.45V8.5h1.25a.5.5 0 0 1 .5.5a.508.508 0 0 1-.5.5H11.5a1 1 0 0 0 0 2h1a2 2 0 0 1 2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BagOne(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.37 17.65a16.777 16.777 0 0 0-2.01-5.54a17.037 17.037 0 0 0-3.74-4.55l-.1-.08a.121.121 0 0 1-.02-.15l1.49-2.59a1.12 1.12 0 0 0 0-1.12a1.092 1.092 0 0 0-.97-.55H8.98a1.1 1.1 0 0 0-.97.55a1.12 1.12 0 0 0 0 1.12l1.5 2.59a.124.124 0 0 1-.03.15l-.09.08a17.327 17.327 0 0 0-5.76 10.09a4.051 4.051 0 0 0-.04.48a2.8 2.8 0 0 0 2.8 2.8h11.23a2.782 2.782 0 0 0 2.13-.99a2.834 2.834 0 0 0 .62-2.29M8.88 4.24a.1.1 0 0 1 0-.12a.106.106 0 0 1 .1-.05h6.04a.143.143 0 0 1 .11.05a.163.163 0 0 1 0 .12l-1.59 2.8h-3.08Zm5.09 4.08a16.436 16.436 0 0 1 5.42 9.5a1.817 1.817 0 0 1-.4 1.47a1.786 1.786 0 0 1-1.37.64H6.39a1.805 1.805 0 0 1-1.8-1.8a1.628 1.628 0 0 1 .03-.31a16.286 16.286 0 0 1 5.42-9.5l.32-.28h3.28Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bandage(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.435 7.5H6.565a4.5 4.5 0 0 0 0 9h10.87a4.5 4.5 0 0 0 0-9m-9.93 8h-.94a3.5 3.5 0 0 1 0-7h.94Zm8 0h-7v-7h7Zm1.93 0h-.93v-7h.93a3.5 3.5 0 0 1 0 7"/><circle cx="10.252" cy="10.501" r=".625" fill="currentColor"/><circle cx="10.252" cy="13.501" r=".625" fill="currentColor"/><circle cx="13.752" cy="10.5" r=".625" fill="currentColor"/><circle cx="13.752" cy="13.5" r=".625" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bank(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.505 17.943v-7.581a1.491 1.491 0 0 0 1.39-1.12a1.468 1.468 0 0 0-.7-1.68l-7.45-4.3a1.521 1.521 0 0 0-1.49 0l-7.45 4.3a1.468 1.468 0 0 0-.7 1.68a1.487 1.487 0 0 0 1.45 1.12h.13v7.57h-.12a1.5 1.5 0 0 0 0 3h14.87a1.5 1.5 0 0 0 .07-2.989M4.555 9.362a.505.505 0 0 1-.25-.94l7.45-4.289a.474.474 0 0 1 .49 0L19.7 8.422a.5.5 0 0 1-.25.94Zm13.95 1v7.57H14.9v-7.57Zm-4.61 0v7.57h-3.61v-7.57Zm-4.61 0v7.57h-3.6v-7.57Zm10.15 9.57H4.565a.5.5 0 0 1-.5-.5a.5.5 0 0 1 .5-.5h14.87a.5.5 0 0 1 .5.5a.5.5 0 0 1-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Barcode(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.066 4.065H3.648a1.732 1.732 0 0 0-.963.189a1.368 1.368 0 0 0-.619 1.226v4.585a.5.5 0 0 0 1 0v-4.28a1.794 1.794 0 0 1 .014-.518c.077-.236.319-.2.514-.2h4.472a.5.5 0 0 0 0-1Zm-6.003 9.872v4.418a1.733 1.733 0 0 0 .189.963a1.369 1.369 0 0 0 1.227.619h4.584a.5.5 0 0 0 0-1h-4.28a1.831 1.831 0 0 1-.518-.014c-.236-.077-.2-.319-.2-.514v-4.472a.5.5 0 0 0-1 0Zm13.871 5.998h4.418a1.732 1.732 0 0 0 .963-.189a1.368 1.368 0 0 0 .619-1.226v-4.585a.5.5 0 0 0-1 0v4.28a1.794 1.794 0 0 1-.014.518c-.077.236-.319.2-.514.2h-4.472a.5.5 0 0 0 0 1Zm6.003-9.872V5.645a1.733 1.733 0 0 0-.189-.963a1.369 1.369 0 0 0-1.227-.619h-4.584a.5.5 0 0 0 0 1h4.28a1.831 1.831 0 0 1 .518.014c.236.077.2.319.2.514v4.472a.5.5 0 0 0 1 0Z"/><rect width="1" height="8.709" x="10.999" y="7.643" fill="currentColor" rx=".5"/><rect width="1" height="8.709" x="14.249" y="7.643" fill="currentColor" rx=".5"/><rect width="1" height="8.709" x="16.499" y="7.643" fill="currentColor" rx=".5"/><rect width="1" height="8.709" x="6.499" y="7.643" fill="currentColor" rx=".5"/><rect width="1.5" height="8.709" x="8.499" y="7.643" fill="currentColor" rx=".75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Baseball(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.02 4.976A9.927 9.927 0 1 0 15.74 21.2a9.908 9.908 0 0 0 6.19-9.2a9.856 9.856 0 0 0-2.91-7.024m-13.34.71a8.9 8.9 0 0 1 6.04-2.61a8.461 8.461 0 0 1-.34 2.26l-.34-.19a.5.5 0 0 0-.5.86l.5.29a9.227 9.227 0 0 1-1.57 2.47l-.35-.35a.5.5 0 0 0-.7 0a.5.5 0 0 0 0 .71l.34.34a8.875 8.875 0 0 1-2.47 1.58l-.29-.51a.5.5 0 0 0-.68-.19a.505.505 0 0 0-.18.69l.2.34a8.2 8.2 0 0 1-2.26.35a8.827 8.827 0 0 1 2.6-6.04M11.74 17a.5.5 0 1 0-.5.87l.49.29a10.008 10.008 0 0 0-.45 2.74a8.9 8.9 0 0 1-8.18-8.17a9.378 9.378 0 0 0 2.75-.46l.29.5a.5.5 0 0 0 .43.25a.475.475 0 0 0 .25-.07a.493.493 0 0 0 .18-.68l-.21-.36a9.461 9.461 0 0 0 2.68-1.73l.36.36a.5.5 0 0 0 .35.15a.508.508 0 0 0 .36-.15a.513.513 0 0 0 0-.71l-.36-.36A9.665 9.665 0 0 0 11.9 6.8l.37.21a.475.475 0 0 0 .25.07a.511.511 0 0 0 .44-.25a.494.494 0 0 0-.19-.68l-.51-.29a9.789 9.789 0 0 0 .46-2.76a8.924 8.924 0 0 1 8.18 8.18a10.08 10.08 0 0 0-2.74.46l-.28-.49a.505.505 0 0 0-.69-.18a.491.491 0 0 0-.18.68l.2.35a9.684 9.684 0 0 0-2.68 1.73l-.35-.35a.5.5 0 0 0-.71 0a.5.5 0 0 0 0 .7l.36.36a9.2 9.2 0 0 0-1.73 2.67Zm6.58 1.32a8.851 8.851 0 0 1-6.04 2.6a8.388 8.388 0 0 1 .34-2.25l.35.2a.451.451 0 0 0 .25.07a.5.5 0 0 0 .43-.25a.505.505 0 0 0-.18-.69l-.51-.29a8.7 8.7 0 0 1 1.57-2.47l.36.36a.5.5 0 0 0 .7-.71l-.36-.36a9.124 9.124 0 0 1 2.48-1.57l.3.52a.5.5 0 0 0 .43.25a.451.451 0 0 0 .25-.07a.505.505 0 0 0 .19-.68l-.21-.36a8.449 8.449 0 0 1 2.25-.34a8.992 8.992 0 0 1-.66 3.14a9.172 9.172 0 0 1-1.94 2.896Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Basketball(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.02 4.971a9.941 9.941 0 1 0 0 14.05a9.941 9.941 0 0 0 0-14.05m-13.34.71a8.894 8.894 0 0 1 6.05-2.6a8.812 8.812 0 0 1-2.61 6.04a8.75 8.75 0 0 1-6.04 2.61a8.875 8.875 0 0 1 2.6-6.05m-2.58 7.05a9.772 9.772 0 0 0 6.73-2.9a9.8 9.8 0 0 0 2.9-6.73a8.908 8.908 0 0 1 5.23 2.24L5.34 17.951a8.881 8.881 0 0 1-2.24-5.22m8.18 8.17a8.872 8.872 0 0 1-5.23-2.24l12.61-12.62a8.91 8.91 0 0 1 2.24 5.24a9.86 9.86 0 0 0-9.62 9.62m7.04-2.59a8.856 8.856 0 0 1-6.04 2.61a8.851 8.851 0 0 1 8.64-8.64a8.847 8.847 0 0 1-2.6 6.03"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryCharging(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.505 18.5H4.065a2 2 0 0 1-2-2v-9a2 2 0 0 1 2-2h13.44a2 2 0 0 1 2 2v1h.93a1.5 1.5 0 0 1 1.5 1.5v4a1.5 1.5 0 0 1-1.5 1.5h-.93v1a2 2 0 0 1-2 2m-13.44-12a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h13.44a1 1 0 0 0 1-1v-1.25a.752.752 0 0 1 .75-.75h1.18a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-.5-.5h-1.18a.752.752 0 0 1-.75-.75V7.5a1 1 0 0 0-1-1Z"/><path fill="currentColor" d="M13.174 11.191h-1.283a.11.11 0 0 1-.1-.15l.655-1.669a.251.251 0 0 0-.233-.342H9.274a.248.248 0 0 0-.231.157l-.751 1.853a.11.11 0 0 0 .1.151h1.437a.11.11 0 0 1 .1.144l-.776 3.53a.085.085 0 0 0 .139.081l3.947-3.561a.109.109 0 0 0-.065-.194"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryEmpty(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.505 18.5H4.065a2 2 0 0 1-2-2v-9a2 2 0 0 1 2-2h13.44a2 2 0 0 1 2 2v1h.93a1.5 1.5 0 0 1 1.5 1.5v4a1.5 1.5 0 0 1-1.5 1.5h-.93v1a2 2 0 0 1-2 2m-13.44-12a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h13.44a1 1 0 0 0 1-1v-1.25a.752.752 0 0 1 .75-.75h1.18a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-.5-.5h-1.18a.752.752 0 0 1-.75-.75V7.5a1 1 0 0 0-1-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryFull(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.505 18.5H4.065a2 2 0 0 1-2-2v-9a2 2 0 0 1 2-2h13.44a2 2 0 0 1 2 2v1h.93a1.5 1.5 0 0 1 1.5 1.5v4a1.5 1.5 0 0 1-1.5 1.5h-.93v1a2 2 0 0 1-2 2m-13.44-12a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h13.44a1 1 0 0 0 1-1v-1.25a.751.751 0 0 1 .75-.75h1.18a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-.5-.5h-1.18a.751.751 0 0 1-.75-.75V7.5a1 1 0 0 0-1-1Z"/><rect width="13.437" height="8.998" x="4.063" y="7.499" fill="currentColor" rx="1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BeakerOne(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m19.447 18.645l-.51-1.52a17.9 17.9 0 0 0-4.02-6.66a1.493 1.493 0 0 1-.42-1.04v-6.36H15a.5.5 0 0 0 0-1H9a.5.5 0 0 0 0 1h.5v6.36a1.484 1.484 0 0 1-.41 1.04a17.9 17.9 0 0 0-4.02 6.66l-.52 1.52a2.5 2.5 0 0 0 2.37 3.29h10.16a2.5 2.5 0 0 0 2.37-3.29Zm-9.64-7.49a2.477 2.477 0 0 0 .69-1.73v-6.36h3v6.36a2.486 2.486 0 0 0 .7 1.73a16.907 16.907 0 0 1 3.01 4.38H6.787a16.943 16.943 0 0 1 3.02-4.38m8.49 9.16a1.507 1.507 0 0 1-1.22.62H6.917a1.5 1.5 0 0 1-1.42-1.98l.51-1.52q.15-.45.33-.9h11.32q.18.45.33.9l.51 1.52a1.5 1.5 0 0 1-.197 1.36Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BeerMugFull(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.356 9.63h-.97V7.99a2.938 2.938 0 0 0 .5-1.65a1.77 1.77 0 0 0-.01-.23a2.905 2.905 0 0 0-1.64-2.38a2.956 2.956 0 0 0-2.4-.07a3.278 3.278 0 0 0-5.62 0a2.9 2.9 0 0 0-1.68-.17a2.866 2.866 0 0 0-2.16 1.75a2.948 2.948 0 0 0 .3 2.77v11.42a2.5 2.5 0 0 0 2.5 2.5h7.71a2.5 2.5 0 0 0 2.5-2.5v-.99l.91-.36a2.433 2.433 0 0 0 1.54-2.27V11.1a1.481 1.481 0 0 0-1.48-1.47m-1.97 9.8a1.5 1.5 0 0 1-1.5 1.5h-7.71a1.5 1.5 0 0 1-1.5-1.5v-8.09a2.858 2.858 0 0 0 1.93.74c.13 0 .25-.01.37-.02v6.34a.5.5 0 0 0 .5.5a.5.5 0 0 0 .5-.5v-6.58a.17.17 0 0 0-.01-.07a2.939 2.939 0 0 0 1.57-2.46h4.42a2.86 2.86 0 0 0 1.43-.38Zm-.01-11.77a1.949 1.949 0 0 1-1.42.63h-4.61a.8.8 0 0 0-.79.61a1.231 1.231 0 0 0-.02.2a1.975 1.975 0 0 1-1.05 1.78a1.934 1.934 0 0 1-2.8-1.72a1.808 1.808 0 0 1 .17-.77a.6.6 0 0 0-.13-.68a1.939 1.939 0 0 1-.41-2.11a1.868 1.868 0 0 1 1.4-1.13a2.531 2.531 0 0 1 .38-.03a1.909 1.909 0 0 1 .86.2a.766.766 0 0 0 .59.06A.8.8 0 0 0 9 4.32a2.273 2.273 0 0 1 4.06 0a.751.751 0 0 0 .44.38a.8.8 0 0 0 .59-.05a1.917 1.917 0 0 1 2.79 1.54a1.886 1.886 0 0 1-.504 1.47m2.46 8.15a1.428 1.428 0 0 1-.92 1.34l-.52.22v-6.74h.96a.478.478 0 0 1 .48.47Z"/><path fill="currentColor" d="M13.577 18.9a.5.5 0 0 1-.5-.5v-6.58a.5.5 0 0 1 1 0v6.58a.5.5 0 0 1-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellOff(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.624 5.2c-.571.3-.079 1.124.5.864c.509-.227 1.068-.287 1.115-.95c.028-.41.014-.81.476-.993a.776.776 0 0 1 1.035.71c.048.461.035.821.548 1.024a4.811 4.811 0 0 1 2.812 2.432a5.63 5.63 0 0 1 .414 2.467v2.02a.5.5 0 0 0 1 0c0-1.646.185-3.394-.521-4.929a5.542 5.542 0 0 0-3.019-2.808c-.034-.013-.155-.069-.227-.092v-.059a2.009 2.009 0 0 0-.257-.945a1.739 1.739 0 0 0-3.1.172a1.992 1.992 0 0 0-.153.792v.052a6.93 6.93 0 0 0-.623.243M4.57 3.86c-.46-.46-1.17.25-.71.7c1.06 1.06 2.12 2.13 3.18 3.19a5.535 5.535 0 0 0-.57 2.44v4.54a2.122 2.122 0 0 0-1.88 2.11v.53a2.121 2.121 0 0 0 2.12 2.12h3.59a1.725 1.725 0 0 0 3.4 0h3.59a2.12 2.12 0 0 0 1.15-.34l.99.99a.5.5 0 0 0 .71-.71ZM17.7 18.41a1.15 1.15 0 0 1-.41.08H6.71a1.118 1.118 0 0 1-1.12-1.12v-.53a1.118 1.118 0 0 1 1.12-1.12a.762.762 0 0 0 .76-.77v-4.76a4.375 4.375 0 0 1 .33-1.68Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellOn(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.79 15.34a2.087 2.087 0 0 0-1.26-.61v-4.54a5.5 5.5 0 0 0-1.62-3.91a5.826 5.826 0 0 0-2.15-1.33v-.06a1.8 1.8 0 0 0-1.61-1.81a1.749 1.749 0 0 0-1.91 1.75v.12a5.547 5.547 0 0 0-3.77 5.24v4.54a2.122 2.122 0 0 0-1.88 2.11v.53a2.121 2.121 0 0 0 2.12 2.12h3.59a1.725 1.725 0 0 0 3.4 0h3.59a2.121 2.121 0 0 0 2.12-2.12v-.53a2.1 2.1 0 0 0-.62-1.5m-.38 2.03a1.118 1.118 0 0 1-1.12 1.12H6.71a1.118 1.118 0 0 1-1.12-1.12v-.53a1.118 1.118 0 0 1 1.12-1.12a.762.762 0 0 0 .76-.77v-4.76a4.555 4.555 0 0 1 3.24-4.34a.729.729 0 0 0 .53-.71v-.31a.735.735 0 0 1 .25-.56a.744.744 0 0 1 .51-.2h.07a.807.807 0 0 1 .69.82v.25a.729.729 0 0 0 .53.71a4.668 4.668 0 0 1 1.91 1.14a4.468 4.468 0 0 1 1.33 3.2v4.76a.8.8 0 0 0 .22.55a.773.773 0 0 0 .54.22a1.127 1.127 0 0 1 1.12 1.12Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bezier(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.435 8.005a1.5 1.5 0 1 0-1.41-2H13.5v-.25a.749.749 0 0 0-.75-.75h-1.5a.755.755 0 0 0-.75.75v.25H4.975a1.5 1.5 0 0 0-2.91.5a1.5 1.5 0 0 0 2.91.5h3.79c-2.5 1.61-4.23 5-4.47 8.99h-.28a.749.749 0 0 0-.75.75v1.5a.749.749 0 0 0 .75.75h1.5a.755.755 0 0 0 .75-.75v-1.5a.755.755 0 0 0-.75-.75h-.21c.27-4.22 2.38-7.78 5.19-8.73a.747.747 0 0 0 .75.74h1.5a.741.741 0 0 0 .75-.74c2.81.95 4.93 4.51 5.21 8.73h-.22a.749.749 0 0 0-.75.75v1.5a.749.749 0 0 0 .75.75h1.5a.755.755 0 0 0 .75-.75v-1.5a.755.755 0 0 0-.75-.75H19.7c-.24-3.99-1.97-7.38-4.46-8.99h3.78a1.5 1.5 0 0 0 1.415 1m0-2a.508.508 0 0 1 .5.5a.5.5 0 0 1-.5.5a.5.5 0 0 1-.5-.5a.5.5 0 0 1 .5-.5m-16.87 1a.5.5 0 0 1-.5-.5a.5.5 0 0 1 .5-.5a.508.508 0 0 1 .5.5a.5.5 0 0 1-.5.5M5.265 18h-1v-1h1ZM12.5 7.005h-1v-1h1ZM18.735 17h1v1h-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bitcoin(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21.934A9.934 9.934 0 1 1 21.934 12A9.945 9.945 0 0 1 12 21.934m0-18.868A8.934 8.934 0 1 0 20.934 12A8.944 8.944 0 0 0 12 3.066"/><path fill="currentColor" d="M14.28 11.78a1.994 1.994 0 0 0-1.53-3.28h-.25V7.47A.489.489 0 0 0 12 7a.483.483 0 0 0-.5.47V8.5h-1.25a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h1.25v1.03a.483.483 0 0 0 .5.47a.489.489 0 0 0 .5-.47V15.5h.75a2.006 2.006 0 0 0 2-2a2.033 2.033 0 0 0-.97-1.72M10.25 9.5h2.5a1 1 0 0 1 0 2h-2.5Zm3 5h-3v-2h3a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bluetooth(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.649 15.938L11.829 12l6.82-3.94a.984.984 0 0 0 .5-.87a.968.968 0 0 0-.5-.861L12.029 2.5a.989.989 0 0 0-1 0a1 1 0 0 0-.5.87v7.769q-2.1-1.23-4.22-2.44c-.24-.139-.47-.279-.71-.409a.5.5 0 0 0-.51.86l4.95 2.85c-1.41.81-2.83 1.62-4.23 2.44c-.24.129-.48.27-.72.41a.5.5 0 0 0 .51.86c1.65-.951 3.28-1.891 4.93-2.85v7.769a.993.993 0 0 0 .5.871a.969.969 0 0 0 1 0l6.62-3.82a1.007 1.007 0 0 0 0-1.74Zm-7.12-12.57l6.62 3.82l-6.62 3.83Zm0 17.259v-7.639l6.62 3.82Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookmark(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.6 21.945a1.483 1.483 0 0 1-1.01-.4l-4.251-3.9a.5.5 0 0 0-.68 0l-4.25 3.9a1.5 1.5 0 0 1-2.516-1.1V4.57a2.5 2.5 0 0 1 2.5-2.5h9.214a2.5 2.5 0 0 1 2.5 2.5v15.872a1.481 1.481 0 0 1-.9 1.374a1.507 1.507 0 0 1-.607.129M12 16.51a1.5 1.5 0 0 1 1.018.395l4.251 3.9a.5.5 0 0 0 .839-.368V4.57a1.5 1.5 0 0 0-1.5-1.5H7.393a1.5 1.5 0 0 0-1.5 1.5v15.872a.5.5 0 0 0 .839.368l4.251-3.91A1.5 1.5 0 0 1 12 16.51"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkCheck(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.6 21.938a1.482 1.482 0 0 1-1.011-.4l-4.251-3.9a.5.5 0 0 0-.678 0l-4.25 3.9a1.5 1.5 0 0 1-2.517-1.1V4.563a2.5 2.5 0 0 1 2.5-2.5h9.214a2.5 2.5 0 0 1 2.5 2.5v15.872a1.483 1.483 0 0 1-.9 1.375a1.526 1.526 0 0 1-.607.128M12 16.5a1.5 1.5 0 0 1 1.018.395l4.251 3.905a.5.5 0 0 0 .838-.368V4.563a1.5 1.5 0 0 0-1.5-1.5H7.393a1.5 1.5 0 0 0-1.5 1.5v15.872a.5.5 0 0 0 .839.368l4.251-3.903A1.5 1.5 0 0 1 12 16.5"/><path fill="currentColor" d="M14.85 9.08c-.11.12-.23.23-.35.35c-.83.83-1.65 1.65-2.47 2.48a.513.513 0 0 1-.71 0c-.47-.48-.94-.95-1.42-1.42a.5.5 0 0 1 .71-.71c.35.36.7.71 1.06 1.06c.83-.82 1.65-1.65 2.48-2.47c.45-.46 1.16.25.7.71"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkMinus(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.6 21.938a1.482 1.482 0 0 1-1.011-.4l-4.251-3.9a.5.5 0 0 0-.678 0l-4.25 3.9a1.5 1.5 0 0 1-2.517-1.1V4.563a2.5 2.5 0 0 1 2.5-2.5h9.214a2.5 2.5 0 0 1 2.5 2.5v15.872a1.483 1.483 0 0 1-.9 1.375a1.526 1.526 0 0 1-.607.128M12 16.5a1.5 1.5 0 0 1 1.018.395l4.251 3.905a.5.5 0 0 0 .838-.368V4.563a1.5 1.5 0 0 0-1.5-1.5H7.393a1.5 1.5 0 0 0-1.5 1.5v15.872a.5.5 0 0 0 .839.368l4.251-3.903A1.5 1.5 0 0 1 12 16.5"/><path fill="currentColor" d="M10 10.277a.5.5 0 0 1 0-1h4a.5.5 0 0 1 0 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkPlus(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.6 21.938a1.482 1.482 0 0 1-1.011-.4l-4.251-3.9a.5.5 0 0 0-.678 0l-4.25 3.9a1.5 1.5 0 0 1-2.517-1.1V4.563a2.5 2.5 0 0 1 2.5-2.5h9.214a2.5 2.5 0 0 1 2.5 2.5v15.872a1.483 1.483 0 0 1-.9 1.375a1.526 1.526 0 0 1-.607.128M12 16.5a1.5 1.5 0 0 1 1.018.395l4.251 3.905a.5.5 0 0 0 .838-.368V4.563a1.5 1.5 0 0 0-1.5-1.5H7.393a1.5 1.5 0 0 0-1.5 1.5v15.872a.5.5 0 0 0 .839.368l4.251-3.903A1.5 1.5 0 0 1 12 16.5"/><path fill="currentColor" d="M14 10.28h-1.5v1.5a.5.5 0 0 1-1 0v-1.5H10a.5.5 0 0 1 0-1h1.5v-1.5a.5.5 0 0 1 1 0v1.5H14a.5.5 0 0 1 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkRemove(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.6 21.938a1.482 1.482 0 0 1-1.011-.4l-4.251-3.9a.5.5 0 0 0-.678 0l-4.25 3.9a1.5 1.5 0 0 1-2.517-1.1V4.563a2.5 2.5 0 0 1 2.5-2.5h9.214a2.5 2.5 0 0 1 2.5 2.5v15.872a1.483 1.483 0 0 1-.9 1.375a1.526 1.526 0 0 1-.607.128M12 16.5a1.5 1.5 0 0 1 1.018.395l4.251 3.905a.5.5 0 0 0 .838-.368V4.563a1.5 1.5 0 0 0-1.5-1.5H7.393a1.5 1.5 0 0 0-1.5 1.5v15.872a.5.5 0 0 0 .839.368l4.251-3.903A1.5 1.5 0 0 1 12 16.5"/><path fill="currentColor" d="M10.23 10.84a.5.5 0 0 0 .71.71L12 10.491l1.06 1.059a.523.523 0 0 0 .71 0a.513.513 0 0 0 0-.71l-1.061-1.061L13.77 8.72a.5.5 0 0 0-.71-.71c-.35.35-.7.71-1.06 1.06l-1.06-1.06a.5.5 0 0 0-.71 0a.524.524 0 0 0 0 .71c.35.35.71.7 1.06 1.06Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BowlNoodles(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.5 10.705a.948.948 0 0 0-.92-.67h-5.07v-1.68l4.75.17h.02a.734.734 0 0 0 .73-.73a.718.718 0 0 0-.75-.72l-4.75.17v-1.84l4.78-.67a.723.723 0 0 0 .62-.72a.487.487 0 0 0-.01-.12a.716.716 0 0 0-.87-.58l-4.6.98a1.5 1.5 0 0 0-2.92.47v.15l-1.14.24a1.494 1.494 0 0 0-2.86.61v.01l-1.2.25a.267.267 0 0 0-.2.26v.04a.257.257 0 0 0 .29.21l1.11-.15V7.5l-1.25.04a.263.263 0 0 0-.25.26a.256.256 0 0 0 .25.26l1.25.04v1.94H5.425a.963.963 0 0 0-.92.68a10.119 10.119 0 0 0 1.19 8.53l.61.92a1.233 1.233 0 0 0 1.05.57h9.3a1.228 1.228 0 0 0 1.04-.57l.61-.92a10.136 10.136 0 0 0 1.195-8.545m-7.99-5.94a.5.5 0 0 1 .5-.5a.5.5 0 0 1 .5.5v5.24h-1Zm-2 1.2l1-.14v1.53l-1 .03Zm0 2.25l1 .03v1.79h-1Zm-2-2.45a.5.5 0 0 1 .5-.5a.5.5 0 0 1 .5.5v4.24h-1Zm9.96 12.93l-.6.93a.261.261 0 0 1-.21.11h-9.3a.236.236 0 0 1-.21-.11l-.61-.93a9.229 9.229 0 0 1-1.11-7.66l13.12-.03a9.122 9.122 0 0 1-1.085 7.695Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxList(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.562 8.062h-2a1.5 1.5 0 0 1-1.5-1.5v-2a1.5 1.5 0 0 1 1.5-1.5h2a1.5 1.5 0 0 1 1.5 1.5v2a1.5 1.5 0 0 1-1.5 1.5m-2-4a.5.5 0 0 0-.5.5v2a.5.5 0 0 0 .5.5h2a.5.5 0 0 0 .5-.5v-2a.5.5 0 0 0-.5-.5Zm2 16.876h-2a1.5 1.5 0 0 1-1.5-1.5v-2a1.5 1.5 0 0 1 1.5-1.5h2a1.5 1.5 0 0 1 1.5 1.5v2a1.5 1.5 0 0 1-1.5 1.5m-2-4a.5.5 0 0 0-.5.5v2a.5.5 0 0 0 .5.5h2a.5.5 0 0 0 .5-.5v-2a.5.5 0 0 0-.5-.5Zm2-2.438h-2a1.5 1.5 0 0 1-1.5-1.5v-2a1.5 1.5 0 0 1 1.5-1.5h2a1.5 1.5 0 0 1 1.5 1.5v2a1.5 1.5 0 0 1-1.5 1.5m-2-4a.5.5 0 0 0-.5.5v2a.5.5 0 0 0 .5.5h2a.5.5 0 0 0 .5-.5v-2a.5.5 0 0 0-.5-.5Zm15.876-4.438h-9a.5.5 0 0 1 0-1h9a.5.5 0 0 1 0 1m0 6.438h-9a.5.5 0 0 1 0-1h9a.5.5 0 0 1 0 1m0 6.435h-9a.5.5 0 1 1 0-1h9a.5.5 0 0 1 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Boxes(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.435 11.5h-2.72V4.56a1.5 1.5 0 0 0-1.5-1.5h-6.43a1.5 1.5 0 0 0-1.5 1.5v6.94h-2.72a1.5 1.5 0 0 0-1.5 1.5v6.44a1.5 1.5 0 0 0 1.5 1.5H11a1.468 1.468 0 0 0 1-.39a1.487 1.487 0 0 0 1 .39h6.44a1.5 1.5 0 0 0 1.5-1.5V13a1.5 1.5 0 0 0-1.505-1.5M11.5 19.44a.5.5 0 0 1-.5.5H4.565a.5.5 0 0 1-.5-.5V13a.5.5 0 0 1 .5-.5h1.97v2a.5.5 0 0 0 .5.5h1.5a.508.508 0 0 0 .5-.5v-2H11.5ZM8.285 11.5V4.56a.5.5 0 0 1 .5-.5h1.96v2a.5.5 0 0 0 .5.5h1.5a.5.5 0 0 0 .5-.5v-2h1.97a.5.5 0 0 1 .5.5v6.94Zm11.65 7.94a.508.508 0 0 1-.5.5H13a.508.508 0 0 1-.5-.5V12.5h2.47v2a.5.5 0 0 0 .5.5h1.5a.5.5 0 0 0 .5-.5v-2h1.97a.5.5 0 0 1 .5.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrightnessDown(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 17.5a5.5 5.5 0 1 1 5.5-5.5a5.506 5.506 0 0 1-5.5 5.5m0-10a4.5 4.5 0 1 0 4.5 4.5A4.505 4.505 0 0 0 12 7.5"/><circle cx="12" cy="2.813" r=".75" fill="currentColor"/><circle cx="12" cy="21.187" r=".75" fill="currentColor"/><circle cx="21.187" cy="12" r=".75" fill="currentColor"/><circle cx="2.813" cy="12" r=".75" fill="currentColor"/><circle cx="18.496" cy="5.504" r=".75" fill="currentColor"/><circle cx="5.504" cy="18.496" r=".75" fill="currentColor"/><circle cx="18.496" cy="18.496" r=".75" fill="currentColor"/><circle cx="5.504" cy="5.504" r=".75" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrightnessUp(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 17.5a5.5 5.5 0 1 1 5.5-5.5a5.506 5.506 0 0 1-5.5 5.5m0-10a4.5 4.5 0 1 0 4.5 4.5A4.505 4.505 0 0 0 12 7.5"/><circle cx="12" cy="3.063" r="1" fill="currentColor"/><circle cx="12" cy="20.937" r="1" fill="currentColor"/><circle cx="20.937" cy="12" r="1" fill="currentColor"/><circle cx="3.063" cy="12" r="1" fill="currentColor"/><circle cx="18.319" cy="5.681" r="1" fill="currentColor"/><circle cx="5.681" cy="18.319" r="1" fill="currentColor"/><circle cx="18.319" cy="18.319" r="1" fill="currentColor"/><circle cx="5.681" cy="5.681" r="1" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bullhorn(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.5 16.474a2.5 2.5 0 0 1-1.436-2.264V9.791a2.5 2.5 0 0 1 2.5-2.5h4.343c.793 0 1.581-.132 2.33-.392c1.859-.705 3.792-1.727 5.24-2.922l.869-.718a1 1 0 0 1 1.587.808v6.717a1.24 1.24 0 0 1 0 2.433v6.718a1.001 1.001 0 0 1-1.588.807l-.868-.718c-1.446-1.195-3.364-2.214-5.226-2.891a7.07 7.07 0 0 0-2.328-.394c-.609-.029-1.265-.029-1.265-.029v2.147a2.08 2.08 0 0 1-4.158 0zm1 .236v2.147a1.079 1.079 0 1 0 2.158 0V16.71H5.564zm6-.882l.142.04a17.632 17.632 0 0 1 6.473 3.385l.818.677V4.071l-.82.677a17.605 17.605 0 0 1-6.468 3.379l-.145.041zm-2.842-.118H10.5V8.291H5.564a1.5 1.5 0 0 0-1.5 1.5v4.419a1.5 1.5 0 0 0 1.499 1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Burger(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.325 11.015a7.344 7.344 0 0 0-14.66 0a2.014 2.014 0 0 0-1.6 1.96v.16a2.016 2.016 0 0 0 1.64 1.97l.27 2.45a2.593 2.593 0 0 0 2.59 2.32h8.87a2.593 2.593 0 0 0 2.59-2.32l.27-2.45a2.016 2.016 0 0 0 1.64-1.97v-.16a2 2 0 0 0-1.61-1.96M12 5.125a6.365 6.365 0 0 1 6.34 5.85H5.665A6.362 6.362 0 0 1 12 5.125m6.04 12.32a1.6 1.6 0 0 1-1.6 1.43H7.565a1.6 1.6 0 0 1-1.6-1.43l-.26-2.31H18.3Zm1.9-4.31a1 1 0 0 1-1 1H5.065a1 1 0 0 1-1-1v-.16a1 1 0 0 1 1-1h13.87a1 1 0 0 1 1 1Z"/><circle cx="12" cy="6.622" r=".5" fill="currentColor"/><circle cx="8.323" cy="8.323" r=".5" fill="currentColor"/><circle cx="15.676" cy="8.323" r=".5" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalculatorOne(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.5 21.937h-9a2.5 2.5 0 0 1-2.5-2.5V4.563a2.5 2.5 0 0 1 2.5-2.5h9a2.5 2.5 0 0 1 2.5 2.5v14.874a2.5 2.5 0 0 1-2.5 2.5m-9-18.874a1.5 1.5 0 0 0-1.5 1.5v14.874a1.5 1.5 0 0 0 1.5 1.5h9a1.5 1.5 0 0 0 1.5-1.5V4.563a1.5 1.5 0 0 0-1.5-1.5Z"/><path fill="currentColor" d="M14.5 9.757h-5a1.5 1.5 0 0 1-1.5-1.5V6.563a1.5 1.5 0 0 1 1.5-1.5h5a1.5 1.5 0 0 1 1.5 1.5v1.694a1.5 1.5 0 0 1-1.5 1.5m-5-3.694a.5.5 0 0 0-.5.5v1.694a.5.5 0 0 0 .5.5h5a.5.5 0 0 0 .5-.5V6.563a.5.5 0 0 0-.5-.5Z"/><circle cx="12" cy="11.508" r=".75" fill="currentColor"/><circle cx="15.25" cy="11.508" r=".75" fill="currentColor"/><circle cx="8.75" cy="11.508" r=".75" fill="currentColor"/><circle cx="12" cy="14.848" r=".75" fill="currentColor"/><circle cx="15.25" cy="14.848" r=".75" fill="currentColor"/><circle cx="8.75" cy="14.848" r=".75" fill="currentColor"/><circle cx="15.25" cy="18.187" r=".75" fill="currentColor"/><path fill="currentColor" d="M12.248 18.687H8.5a.5.5 0 0 1 0-1h3.744a.5.5 0 1 1 0 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalculatorTwo(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.45 7.83h-2.8v2.81a.5.5 0 0 1-1 0V7.83h-2.81a.5.5 0 0 1 0-1h2.81V4.02a.5.5 0 0 1 1 0v2.81h2.8a.5.5 0 0 1 0 1m-16.905 0a.5.5 0 0 1 0-1h6.619a.5.5 0 0 1 0 1Zm10.291 8.22a.5.5 0 0 1 0-1h6.619a.5.5 0 0 1 0 1Zm0 4.141a.5.5 0 0 1 0-1h6.619a.5.5 0 0 1 0 1ZM9.55 19.61a.5.5 0 0 1-.71.7l-1.98-1.98c-.66.65-1.33 1.32-1.99 1.98a.5.5 0 0 1-.71-.7l1.99-1.99l-1.99-1.99a.5.5 0 0 1 .71-.7l.58.58l1.4 1.4c.67-.66 1.33-1.32 1.99-1.98a.5.5 0 0 1 .71.7l-1.99 1.99Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.438 4.954H16.5V3.546c0-.262-.23-.512-.5-.5a.509.509 0 0 0-.5.5v1.408h-7V3.546c0-.262-.23-.512-.5-.5a.509.509 0 0 0-.5.5v1.408H5.562a2.503 2.503 0 0 0-2.5 2.5v11c0 1.379 1.122 2.5 2.5 2.5h12.875c1.379 0 2.5-1.121 2.5-2.5v-11a2.502 2.502 0 0 0-2.499-2.5m-12.876 1H7.5v.592c0 .262.23.512.5.5c.271-.012.5-.22.5-.5v-.592h7v.592c0 .262.23.512.5.5c.271-.012.5-.22.5-.5v-.592h1.937c.827 0 1.5.673 1.5 1.5v1.584H4.062V7.454c0-.827.673-1.5 1.5-1.5m12.876 14H5.562c-.827 0-1.5-.673-1.5-1.5v-8.416h15.875v8.416a1.5 1.5 0 0 1-1.499 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarDate(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.435 4.955h-1.94v-1.41c0-.26-.23-.51-.5-.5c-.27.01-.5.22-.5.5v1.41h-7v-1.41c0-.26-.23-.51-.5-.5c-.27.01-.5.22-.5.5v1.41h-1.93a2.5 2.5 0 0 0-2.5 2.5v11a2.5 2.5 0 0 0 2.5 2.5h12.87a2.5 2.5 0 0 0 2.5-2.5v-11a2.5 2.5 0 0 0-2.5-2.5m1.5 13.5c0 .83-.67 1.5-1.5 1.5H5.565c-.83 0-1.5-.67-1.5-1.5v-8.42h15.87zm0-9.42H4.065v-1.58c0-.83.67-1.5 1.5-1.5h1.93v.59c0 .26.23.51.5.5c.27-.01.5-.22.5-.5v-.59h7v.59c0 .26.23.51.5.5c.27-.01.5-.22.5-.5v-.59h1.94c.83 0 1.5.67 1.5 1.5z"/><path fill="currentColor" d="M11.492 17.173v-3.46a.075.075 0 0 0-.114-.064l-.638.392a.15.15 0 0 1-.228-.128v-.651c0-.105.055-.203.146-.257l.764-.457a.3.3 0 0 1 .154-.043h.626a.3.3 0 0 1 .3.3v4.367a.3.3 0 0 1-.3.3h-.409a.298.298 0 0 1-.301-.299"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Camera(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.435 19.925H3.565a1.5 1.5 0 0 1-1.5-1.5v-9.14a1.5 1.5 0 0 1 1.5-1.5h2.658a.5.5 0 0 0 .5-.454l.166-1.8a1.49 1.49 0 0 1 1.5-1.454h7.23a1.5 1.5 0 0 1 1.5 1.5l.164 1.756a.5.5 0 0 0 .5.454h2.658a1.5 1.5 0 0 1 1.5 1.5v9.14a1.5 1.5 0 0 1-1.506 1.498M3.565 8.785a.5.5 0 0 0-.5.5v9.14a.5.5 0 0 0 .5.5h16.87a.5.5 0 0 0 .5-.5v-9.14a.5.5 0 0 0-.5-.5h-2.658a1.5 1.5 0 0 1-1.494-1.362l-.166-1.8a.515.515 0 0 0-.5-.546H8.385a.5.5 0 0 0-.5.5l-.168 1.846a1.5 1.5 0 0 1-1.494 1.362Z"/><path fill="currentColor" d="M12 17.282a4 4 0 1 1 4-4a4 4 0 0 1-4 4m0-7a3 3 0 1 0 3 3a3 3 0 0 0-3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatOne(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.316 19.937a1.251 1.251 0 0 1-1.251-1.247v-1.716L2.068 6.56a2.5 2.5 0 0 1 2.5-2.5H19.44a2.5 2.5 0 0 1 2.5 2.5v8.41a2.5 2.5 0 0 1-2.5 2.5H6.918a1.49 1.49 0 0 0-1.06.439L4.2 19.57a1.246 1.246 0 0 1-.884.367M4.568 5.062a1.5 1.5 0 0 0-1.5 1.5L3.06 16.973v1.714a.25.25 0 0 0 .427.176L5.151 17.2a2.482 2.482 0 0 1 1.767-.732H19.44a1.5 1.5 0 0 0 1.5-1.5V6.562a1.5 1.5 0 0 0-1.5-1.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatTwo(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 19.937a1.243 1.243 0 0 1-.833-.319l-1.886-1.686a1.531 1.531 0 0 0-1.08-.458h-3.64a2.5 2.5 0 0 1-2.5-2.5l.006-8.41a2.5 2.5 0 0 1 2.5-2.5h14.872a2.5 2.5 0 0 1 2.5 2.5v8.411a2.5 2.5 0 0 1-2.5 2.5H15.79a1.483 1.483 0 0 0-1.062.441l-1.895 1.7a1.243 1.243 0 0 1-.833.321M4.567 5.063a1.5 1.5 0 0 0-1.5 1.5l-.006 8.411a1.5 1.5 0 0 0 1.5 1.5H8.2a2.483 2.483 0 0 1 1.767.732l1.864 1.667a.248.248 0 0 0 .333 0l1.874-1.682a2.5 2.5 0 0 1 1.751-.716h3.649a1.5 1.5 0 0 0 1.5-1.5V6.563a1.5 1.5 0 0 0-1.5-1.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleAlert(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 9a.5.5 0 0 0-1 0v4.02a.5.5 0 0 0 1 0Z"/><circle cx="12" cy="15.001" r=".5" fill="currentColor"/><path fill="currentColor" d="M12 21.935A9.933 9.933 0 1 1 21.934 12A9.945 9.945 0 0 1 12 21.935m0-18.866A8.933 8.933 0 1 0 20.934 12A8.944 8.944 0 0 0 12 3.069"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleCheck(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.81 10.4a.5.5 0 0 0-.71-.71l-3.56 3.56l-1.73-1.73a.5.5 0 0 0-.71.71l2.08 2.08a.513.513 0 0 0 .71 0Z"/><path fill="currentColor" d="M12 21.934A9.934 9.934 0 1 1 21.933 12A9.945 9.945 0 0 1 12 21.934m0-18.867A8.934 8.934 0 1 0 20.933 12A8.944 8.944 0 0 0 12 3.067"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleChevDown(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.35 13.85a.492.492 0 0 1-.7 0l-3-3a.495.495 0 0 1 .7-.7L12 12.79l2.65-2.64a.495.495 0 0 1 .7.7Z"/><path fill="currentColor" d="M21.933 12A9.933 9.933 0 1 1 12 2.067A9.944 9.944 0 0 1 21.933 12M3.067 12A8.933 8.933 0 1 0 12 3.067A8.943 8.943 0 0 0 3.067 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleChevLeft(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.15 12.35a.492.492 0 0 1 0-.7l3-3a.495.495 0 0 1 .7.7L11.21 12l2.64 2.65a.495.495 0 0 1-.7.7Z"/><path fill="currentColor" d="M12 21.933A9.933 9.933 0 1 1 21.933 12A9.944 9.944 0 0 1 12 21.933m0-18.866A8.933 8.933 0 1 0 20.933 12A8.943 8.943 0 0 0 12 3.067"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleChevRight(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.85 11.65a.492.492 0 0 1 0 .7l-3 3a.495.495 0 0 1-.7-.7L12.79 12l-2.64-2.65a.495.495 0 0 1 .7-.7Z"/><path fill="currentColor" d="M12 2.067A9.933 9.933 0 1 1 2.067 12A9.944 9.944 0 0 1 12 2.067m0 18.866A8.933 8.933 0 1 0 3.067 12A8.943 8.943 0 0 0 12 20.933"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleChevUp(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.65 10.15a.492.492 0 0 1 .7 0l3 3a.495.495 0 0 1-.7.7L12 11.21l-2.65 2.64a.495.495 0 0 1-.7-.7Z"/><path fill="currentColor" d="M2.067 12A9.933 9.933 0 1 1 12 21.934A9.944 9.944 0 0 1 2.067 12m18.866 0A8.933 8.933 0 1 0 12 20.934A8.943 8.943 0 0 0 20.933 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleInfo(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.5 15a.5.5 0 0 0 1 0v-4.019a.5.5 0 0 0-1 0Z"/><circle cx="12" cy="8.999" r=".5" fill="currentColor"/><path fill="currentColor" d="M12 2.065A9.934 9.934 0 1 1 2.066 12A9.945 9.945 0 0 1 12 2.065m0 18.867A8.934 8.934 0 1 0 3.066 12A8.944 8.944 0 0 0 12 20.932"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleList(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.438 6.062h-9a.5.5 0 0 1 0-1h9a.5.5 0 0 1 0 1m0 6.438h-9a.5.5 0 0 1 0-1h9a.5.5 0 0 1 0 1m0 6.435h-9a.5.5 0 1 1 0-1h9a.5.5 0 0 1 0 1M5.562 8.062a2.5 2.5 0 1 1 2.5-2.5a2.5 2.5 0 0 1-2.5 2.5m0-4a1.5 1.5 0 1 0 1.5 1.5a1.5 1.5 0 0 0-1.5-1.5m0 10.438a2.5 2.5 0 1 1 2.5-2.5a2.5 2.5 0 0 1-2.5 2.5m0-4a1.5 1.5 0 1 0 1.5 1.5a1.5 1.5 0 0 0-1.5-1.5m0 10.438a2.5 2.5 0 1 1 2.5-2.5a2.5 2.5 0 0 1-2.5 2.5m0-4a1.5 1.5 0 1 0 1.5 1.5a1.5 1.5 0 0 0-1.5-1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleMinus(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 11.5a.5.5 0 0 1 0 1H9a.5.5 0 0 1 0-1Z"/><path fill="currentColor" d="M12 21.934A9.933 9.933 0 1 1 21.932 12A9.945 9.945 0 0 1 12 21.934m0-18.866A8.933 8.933 0 1 0 20.932 12A8.944 8.944 0 0 0 12 3.068"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleMore(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="12.001" cy="12" r="1" fill="currentColor"/><circle cx="16.001" cy="12" r="1" fill="currentColor"/><circle cx="8.001" cy="12" r="1" fill="currentColor"/><path fill="currentColor" d="M12 21.932A9.934 9.934 0 1 1 21.934 12A9.944 9.944 0 0 1 12 21.932m0-18.867A8.934 8.934 0 1 0 20.934 12A8.944 8.944 0 0 0 12 3.065"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CirclePlus(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15 12.5h-2.5V15a.5.5 0 0 1-1 0v-2.5H9a.5.5 0 0 1 0-1h2.5V9a.5.5 0 0 1 1 0v2.5H15a.5.5 0 0 1 0 1"/><path fill="currentColor" d="M12 21.932A9.934 9.934 0 1 1 21.932 12A9.944 9.944 0 0 1 12 21.932m0-18.867A8.934 8.934 0 1 0 20.932 12A8.944 8.944 0 0 0 12 3.065"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleQuestion(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.411 12.459a1.547 1.547 0 0 1 .341-.6a2.644 2.644 0 0 1 .535-.417a2.2 2.2 0 0 0 .363-.286a1.2 1.2 0 0 0 .256-.363a1.084 1.084 0 0 0 .094-.452a.923.923 0 0 0-.142-.517a.938.938 0 0 0-.374-.338a1.123 1.123 0 0 0-.519-.119a1.173 1.173 0 0 0-.495.107a.934.934 0 0 0-.389.335a.884.884 0 0 0-.111.224a.516.516 0 0 1-.483.359a.506.506 0 0 1-.479-.675a1.661 1.661 0 0 1 .178-.349a1.8 1.8 0 0 1 .748-.634a2.437 2.437 0 0 1 1.031-.215a2.4 2.4 0 0 1 1.082.231a1.737 1.737 0 0 1 .721.641a1.772 1.772 0 0 1 .257.96a1.841 1.841 0 0 1-.118.678a1.685 1.685 0 0 1-.334.536a2.289 2.289 0 0 1-.52.417a2.277 2.277 0 0 0-.462.369a1.113 1.113 0 0 0-.256.455a2.344 2.344 0 0 0-.045.283a.487.487 0 0 1-.483.429a.484.484 0 0 1-.483-.531a2.931 2.931 0 0 1 .087-.528"/><circle cx="11.793" cy="14.891" r=".587" fill="currentColor"/><path fill="currentColor" d="M12 21.931A9.934 9.934 0 1 1 21.934 12A9.945 9.945 0 0 1 12 21.931m0-18.867A8.934 8.934 0 1 0 20.934 12A8.943 8.943 0 0 0 12 3.064"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleRemove(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.525 13.765a.5.5 0 0 0 .71.71c.59-.59 1.175-1.18 1.765-1.76l1.765 1.76a.5.5 0 0 0 .71-.71c-.59-.58-1.18-1.175-1.76-1.765c.41-.42.82-.825 1.23-1.235c.18-.18.35-.36.53-.53a.5.5 0 0 0-.71-.71L12 11.293l-1.765-1.768a.5.5 0 0 0-.71.71L11.293 12Z"/><path fill="currentColor" d="M12 21.933A9.933 9.933 0 1 1 21.934 12A9.945 9.945 0 0 1 12 21.933m0-18.866A8.933 8.933 0 1 0 20.934 12A8.944 8.944 0 0 0 12 3.067"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockOne(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21.933A9.933 9.933 0 1 1 21.933 12A9.944 9.944 0 0 1 12 21.933m0-18.866A8.933 8.933 0 1 0 20.933 12A8.943 8.943 0 0 0 12 3.067"/><path fill="currentColor" d="M11.5 6a.5.5 0 0 1 1 0v4.8c1.13-1.13 2.26-2.27 3.39-3.4a.5.5 0 0 1 .71.71l-4.26 4.25a.463.463 0 0 1-.58.07c-.01-.02-.02-.02-.03-.02a.425.425 0 0 1-.22-.33Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockTwo(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21.933A9.933 9.933 0 1 1 21.933 12A9.944 9.944 0 0 1 12 21.933m0-18.866A8.933 8.933 0 1 0 20.933 12A8.943 8.943 0 0 0 12 3.067"/><path fill="currentColor" d="M18 12.5h-6a.429.429 0 0 1-.34-.14c-.01 0-.01-.01-.02-.02a.429.429 0 0 1-.14-.34V6a.5.5 0 0 1 1 0v5.5H18a.5.5 0 0 1 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloud(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.1 19.4H9.646a7.492 7.492 0 0 1-7.588-7.046A7.4 7.4 0 0 1 9.452 4.6a7.434 7.434 0 0 1 7.136 5.447a4.731 4.731 0 0 1 4.092 1.441a4.664 4.664 0 0 1 1.26 3.529A4.789 4.789 0 0 1 17.1 19.4M3.057 12.309A6.493 6.493 0 0 0 9.646 18.4H17.1a3.787 3.787 0 0 0 3.839-3.453a3.7 3.7 0 0 0-4.5-3.86l-.2.046l-.269-.127a.617.617 0 0 1-.273-.392A6.422 6.422 0 0 0 9.452 5.6a6.4 6.4 0 0 0-6.395 6.711Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudDrizzle(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.605 16.787v1.018a.5.5 0 0 0 1 0v-1.018a.516.516 0 0 0-.146-.354a.5.5 0 0 0-.854.354m-6.105 0v1.018a.516.516 0 0 0 .146.353a.5.5 0 0 0 .854-.353v-1.018a.521.521 0 0 0-.146-.354a.5.5 0 0 0-.854.354m3.052 3.556v1.018a.5.5 0 0 0 1 0v-1.018a.516.516 0 0 0-.146-.354a.5.5 0 0 0-.854.354m-6.106 0v1.018a.5.5 0 0 0 1 0v-1.018a.521.521 0 0 0-.146-.354a.5.5 0 0 0-.854.354m-3.053-3.556v1.018a.5.5 0 0 0 1 0v-1.018a.521.521 0 0 0-.146-.354a.5.5 0 0 0-.854.354M16.1 14.228h-5.99a6.116 6.116 0 0 1-6.194-5.754a6.044 6.044 0 0 1 6.037-6.335a6.07 6.07 0 0 1 5.8 4.366a3.919 3.919 0 0 1 3.288 1.2a3.85 3.85 0 0 1 1.038 2.908a3.946 3.946 0 0 1-3.979 3.615M4.915 8.427a5.117 5.117 0 0 0 5.194 4.8H16.1a2.944 2.944 0 0 0 2.986-2.682a2.873 2.873 0 0 0-3.494-3l-.2.046l-.25-.124a.592.592 0 0 1-.262-.377a5.061 5.061 0 0 0-4.927-3.951a5.043 5.043 0 0 0-5.038 5.288"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudMoon(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.36 11.54a.71.71 0 0 0-.73-.29a5.278 5.278 0 0 1-6.34-4.78a5.379 5.379 0 0 1 .37-2.42a.729.729 0 0 0-.15-.78a.7.7 0 0 0-.76-.16a6.317 6.317 0 0 0-3.98 5.88a5.494 5.494 0 0 0-1.22-.13a6.039 6.039 0 0 0-6.05 6.03c0 .1.01.2.01.3a6.114 6.114 0 0 0 6.19 5.75h5.99a3.941 3.941 0 0 0 3.98-3.61a3.755 3.755 0 0 0-.63-2.38a6.283 6.283 0 0 0 3.36-2.65a.682.682 0 0 0-.04-.76m-6.67 8.4H8.7a5.121 5.121 0 0 1-5.19-4.8a5.042 5.042 0 0 1 5.04-5.28a5.059 5.059 0 0 1 4.92 3.95a.548.548 0 0 0 .26.37l.25.13l.2-.05a2.873 2.873 0 0 1 3.49 3a2.931 2.931 0 0 1-2.98 2.68m2.61-5.83a3.917 3.917 0 0 0-2.95-.89a6.043 6.043 0 0 0-3.57-3.92a5.338 5.338 0 0 1 2.73-4.98a6.325 6.325 0 0 0 4.51 7.85a6.642 6.642 0 0 0 2.12.17a5.2 5.2 0 0 1-2.84 1.77"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudOff(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.8 16.577c-.41.49-1.12-.22-.7-.71a3.585 3.585 0 0 0 .37-4.04A3.266 3.266 0 0 0 16.6 10.3a.5.5 0 0 1-.56-.23a5.391 5.391 0 0 0-5.3-3.1c-.64.04-.64-.96 0-1a6.346 6.346 0 0 1 5.99 3.26a4.255 4.255 0 0 1 4.6 2.1a4.579 4.579 0 0 1-.53 5.247M4.941 4.237a.5.5 0 0 0-.7.7l2.69 2.69a6.273 6.273 0 0 0-1.94 3.78a3.342 3.342 0 0 0-2.65 4.6a3.518 3.518 0 0 0 3.48 2.05h11.53c.58.57 1.14 1.14 1.71 1.71a.5.5 0 0 0 .71-.71Zm.3 12.81a2.352 2.352 0 0 1-2.16-2.25a2.309 2.309 0 0 1 2.35-2.42a.515.515 0 0 0 .5-.5a5.377 5.377 0 0 1 1.71-3.54q4.35 4.365 8.71 8.72Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudOn(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.917 13.484a4.381 4.381 0 0 0-5.19-4.26a6.281 6.281 0 0 0-11.75 2.19a3.237 3.237 0 0 0-2.66 2a3.433 3.433 0 0 0 .82 3.74c1.12 1.03 2.54.89 3.94.89h10.15a4.514 4.514 0 0 0 4.69-4.32Zm-4.65 3.56c-1.19.01-2.38 0-3.56 0c-2.75 0-5.49.06-8.23 0a2.383 2.383 0 0 1-2.33-1.73a2.333 2.333 0 0 1 2.28-2.94a.515.515 0 0 0 .5-.5a5.3 5.3 0 0 1 10.11-1.81a.5.5 0 0 0 .56.23a3.366 3.366 0 0 1 4.33 3.32a3.489 3.489 0 0 1-3.66 3.43"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudRainbow(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.8 6.648a8.6 8.6 0 0 0-12.78.82a6.329 6.329 0 0 0-.761-.05a6.212 6.212 0 0 0-6.2 6.2c0 .1.01.2.01.3a6.277 6.277 0 0 0 6.351 5.89h6.159a4.04 4.04 0 0 0 4.081-3.7a3.916 3.916 0 0 0-1.07-2.97a3.98 3.98 0 0 0-3.37-1.23a5.582 5.582 0 0 0-.38-.97a2.617 2.617 0 0 1 3.75-.08c.46.45 1.169-.26.71-.71a3.66 3.66 0 0 0-2.77-1.05a3.594 3.594 0 0 0-2.2.96a6.746 6.746 0 0 0-1.02-1.12a5.131 5.131 0 0 1 7.031.17c.46.45 1.169-.26.71-.71a6.134 6.134 0 0 0-4.51-1.77a5.982 5.982 0 0 0-4.031 1.73a5.632 5.632 0 0 0-1.409-.65a7.615 7.615 0 0 1 10.99-.35c.455.45 1.164-.258.709-.71m-10.56 2.71a.712.712 0 0 0 .11.08a5.238 5.238 0 0 1 1.979 3.06a.6.6 0 0 0 .271.37l.25.13l.2-.05a2.977 2.977 0 0 1 3.61 3.1a3.037 3.037 0 0 1-3.081 2.76H8.416a5.27 5.27 0 0 1-5.351-4.94a5.2 5.2 0 0 1 8.171-4.51Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudSun(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.465 9.47l-1.9-1.05l.58-2.16a.968.968 0 0 0-.25-.93a.957.957 0 0 0-.93-.24l-2.09.6l-1.13-1.94a.988.988 0 0 0-.83-.47a.967.967 0 0 0-.82.48l-1.06 1.91l-2.16-.58a.951.951 0 0 0-.92.24a.962.962 0 0 0-.25.93l.6 2.1l-.77.44A6.3 6.3 0 0 0 8.1 8.63a6.039 6.039 0 0 0-6.04 6.03c0 .1.01.2.01.3a6.115 6.115 0 0 0 6.19 5.76h5.98a3.952 3.952 0 0 0 3.99-3.62a3.876 3.876 0 0 0-.35-1.88l1.1.3a1.007 1.007 0 0 0 .25.03a.907.907 0 0 0 .67-.28a.95.95 0 0 0 .25-.92l-.6-2.1l1.93-1.12a.956.956 0 0 0 .47-.83a.945.945 0 0 0-.485-.83m-7.73 4.57a2.873 2.873 0 0 1 3.49 3a2.947 2.947 0 0 1-2.99 2.68h-5.98a5.307 5.307 0 0 1-3.6-1.39a4.935 4.935 0 0 1-1.6-3.41A5.043 5.043 0 0 1 8.1 9.63a5.109 5.109 0 0 1 4.09 2.09a5.932 5.932 0 0 1 .4.65a4.974 4.974 0 0 1 .43 1.21a.64.64 0 0 0 .715.46m-.92-3.73a2.106 2.106 0 0 1 4.133-.578a2.114 2.114 0 0 1-2.033 2.688a2.241 2.241 0 0 1-2.1-2.11m4.12 3.64a3.9 3.9 0 0 0-1.08-.67a3.11 3.11 0 1 0-4.01-3.34a6.475 6.475 0 0 0-1.09-.69l.06-.03a.978.978 0 0 0 .44-1.07l-.64-2.1l2.17.58a.981.981 0 0 0 1.07-.44l1.03-1.93l1.12 1.92a.952.952 0 0 0 1.08.45l2.09-.63l-.57 2.14a.922.922 0 0 0 .44 1.09l1.92 1.04l-1.92 1.11a.941.941 0 0 0-.45 1.08l.63 2.09Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoffeeBean(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.151 4.868a6.744 6.744 0 0 0-5.96-1.69a12.009 12.009 0 0 0-6.54 3.47a11.988 11.988 0 0 0-3.48 6.55a6.744 6.744 0 0 0 1.69 5.95a6.406 6.406 0 0 0 4.63 1.78a11.511 11.511 0 0 0 7.87-3.56c3.939-3.94 4.739-9.55 1.79-12.5m-14.99 8.48a11.041 11.041 0 0 1 3.19-5.99a10.976 10.976 0 0 1 5.99-3.19a8.016 8.016 0 0 1 1.18-.09a5.412 5.412 0 0 1 3.92 1.49a.689.689 0 0 1 .11.13a6.542 6.542 0 0 1-2.12 1.23a7.666 7.666 0 0 0-2.96 1.93a7.666 7.666 0 0 0-1.93 2.96a6.589 6.589 0 0 1-1.71 2.63a6.7 6.7 0 0 1-2.63 1.71a7.478 7.478 0 0 0-2.35 1.36a6.18 6.18 0 0 1-.69-4.17m12.49 3.31c-3.55 3.55-8.52 4.35-11.08 1.79a1.538 1.538 0 0 1-.12-.13a6.677 6.677 0 0 1 2.13-1.23a7.862 7.862 0 0 0 2.96-1.93a7.738 7.738 0 0 0 1.93-2.96a6.589 6.589 0 0 1 1.71-2.63a6.589 6.589 0 0 1 2.63-1.71a7.6 7.6 0 0 0 2.34-1.37c1.64 2.712.67 7-2.5 10.17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoffeeCup(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.79 5.67a1.5 1.5 0 0 0-1.16-1.46l-.38-1.38a1.081 1.081 0 0 0-1.05-.76H7.79a1.06 1.06 0 0 0-1.04.76l-.38 1.38a1.537 1.537 0 0 0-1.16 1.55a1.476 1.476 0 0 0 1.06 1.42l.1 2.77a.75.75 0 0 0-.42.22a.768.768 0 0 0-.21.56l.24 5.76a.759.759 0 0 0 .65.72l.08 2.22a2.579 2.579 0 0 0 2.59 2.5h5.39a2.581 2.581 0 0 0 2.6-2.5l.08-2.22a.76.76 0 0 0 .64-.72l.24-5.76a.768.768 0 0 0-.21-.56a.72.72 0 0 0-.41-.22l.1-2.77a1.534 1.534 0 0 0 1.06-1.51m-11-2.6l8.49.03l.3 1.07H7.44Zm8.5 16.33a1.578 1.578 0 0 1-1.6 1.53H9.3a1.575 1.575 0 0 1-1.59-1.53l-.08-2.18h8.74ZM9.9 13.58a2.1 2.1 0 1 1 2.1 2.1a2.1 2.1 0 0 1-2.1-2.1m6.73-3.65H7.37l-.1-2.67h9.45Zm.66-3.67H6.71a.522.522 0 0 1-.5-.59a.5.5 0 0 1 .5-.5h10.58a.528.528 0 0 1 .5.59a.5.5 0 0 1-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoinInsert(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.44 17.7h-3.67a7.484 7.484 0 0 0 1.78-4.86A7.55 7.55 0 1 0 6.23 17.7H2.56a.508.508 0 0 0-.5.5a.5.5 0 0 0 .5.5h18.88a.5.5 0 0 0 .5-.5a.508.508 0 0 0-.5-.5m-5.03 0H7.62a6.546 6.546 0 1 1 8.78-.01Z"/><path fill="currentColor" d="M14 13.965a1.616 1.616 0 0 1-1.5 1.61v.65a.485.485 0 0 1-.5.48a.491.491 0 0 1-.5-.48v-.64h-.81a.5.5 0 0 1-.5-.5a.508.508 0 0 1 .5-.5h1.69a.617.617 0 0 0 .62-.62a.623.623 0 0 0-.62-.62h-.75a1.618 1.618 0 0 1-.13-3.23v-.65a.491.491 0 0 1 .5-.48a.485.485 0 0 1 .5.48v.64h.81a.5.5 0 0 1 0 1h-1.68a.62.62 0 0 0 0 1.24h.75a1.626 1.626 0 0 1 1.62 1.62"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoinsOne(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.46 10.905a3.469 3.469 0 0 0-2.47 1.04a2.3 2.3 0 0 0-.86-1.73a2.257 2.257 0 0 0 .86-1.78a2.288 2.288 0 0 0-2.28-2.29H4.35a2.284 2.284 0 0 0-1.43 4.07a2.282 2.282 0 0 0 0 3.57a2.277 2.277 0 0 0 1.43 4.06h9.36a2.29 2.29 0 0 0 2.06-1.29a3.434 3.434 0 0 0 2.69 1.3a3.475 3.475 0 1 0 0-6.95m-4.75 5.94H4.35a1.28 1.28 0 1 1 0-2.56h9.36a1.28 1.28 0 1 1 0 2.56m0-3.56H4.35a1.285 1.285 0 1 1 0-2.57h9.36a1.285 1.285 0 0 1 0 2.57m0-3.57H4.35a1.285 1.285 0 1 1 0-2.57h9.36a1.285 1.285 0 0 1 0 2.57m4.75 7.14a2.475 2.475 0 1 1 2.48-2.48a2.477 2.477 0 0 1-2.48 2.48"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompassOne(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.661 20.937a1.589 1.589 0 0 1-1.117-.48a1.534 1.534 0 0 1-.4-1.59l3.436-9.93A3.8 3.8 0 0 1 8.938 6.58l9.93-3.439a1.537 1.537 0 0 1 1.589.4a1.532 1.532 0 0 1 .4 1.588l-3.437 9.932a3.8 3.8 0 0 1-2.358 2.358l-9.93 3.439a1.442 1.442 0 0 1-.471.079M19.337 4.062a.424.424 0 0 0-.142.024L9.267 7.525a2.8 2.8 0 0 0-1.742 1.741L4.087 19.2a.6.6 0 0 0 .717.718l9.93-3.439a2.8 2.8 0 0 0 1.741-1.741L19.913 4.8a.551.551 0 0 0-.163-.553a.609.609 0 0 0-.413-.185"/><circle cx="12" cy="12" r="1.563" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCardOff(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.94 7.64v9.3a.5.5 0 0 1-.5.5a.5.5 0 0 1-.5-.5v-7.3h-7.45a.75.75 0 0 1 0-1.5h7.45v-.5a1.5 1.5 0 0 0-1.5-1.5H9.89a.5.5 0 0 1 0-1h9.55a2.5 2.5 0 0 1 2.5 2.5M8.064 14.246h-3a.5.5 0 0 1 0-1h3a.5.5 0 0 1 0 1"/><path fill="currentColor" d="M18.935 14.248h-.944a.5.5 0 0 1 0-1h.944a.5.5 0 0 1 0 1m-.175 3.802L4.01 3.3c-.46-.46-1.17.25-.71.7l1.14 1.14a2.5 2.5 0 0 0-2.38 2.5v8.72a2.5 2.5 0 0 0 2.5 2.5h13.6L20 20.7c.45.46 1.16-.25.7-.71ZM3.06 7.64a1.5 1.5 0 0 1 1.5-1.5h.88c.66.67 1.33 1.34 2 2H3.06Zm9.49 5.61h-.12a.5.5 0 0 0-.5.5a.508.508 0 0 0 .5.5h1.12l3.61 3.61H4.56a1.5 1.5 0 0 1-1.5-1.5V9.64h5.88Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCardOne(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.44 5.14H4.56a2.5 2.5 0 0 0-2.5 2.5v8.72a2.5 2.5 0 0 0 2.5 2.5h14.88a2.5 2.5 0 0 0 2.5-2.5V7.64a2.5 2.5 0 0 0-2.5-2.5M3.06 7.64a1.5 1.5 0 0 1 1.5-1.5h14.88a1.5 1.5 0 0 1 1.5 1.5v.5H3.06Zm17.88 8.72a1.5 1.5 0 0 1-1.5 1.5H4.56a1.5 1.5 0 0 1-1.5-1.5V9.64h17.88Z"/><path fill="currentColor" d="M8.063 14.247h-3a.5.5 0 1 1 0-1h3a.5.5 0 1 1 0 1m10.871.003h-6.5a.5.5 0 1 1 0-1h6.5a.5.5 0 0 1 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCardTwo(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.437 18.859H4.563a2.5 2.5 0 0 1-2.5-2.5V7.641a2.5 2.5 0 0 1 2.5-2.5h14.874a2.5 2.5 0 0 1 2.5 2.5v8.718a2.5 2.5 0 0 1-2.5 2.5M4.563 6.141a1.5 1.5 0 0 0-1.5 1.5v8.718a1.5 1.5 0 0 0 1.5 1.5h14.874a1.5 1.5 0 0 0 1.5-1.5V7.641a1.5 1.5 0 0 0-1.5-1.5Z"/><path fill="currentColor" d="M8.063 14.247h-3a.5.5 0 1 1 0-1h3a.5.5 0 1 1 0 1m10.871.002h-6.5a.5.5 0 0 1 0-1h6.5a.5.5 0 0 1 0 1"/><rect width="2" height="4" x="16.434" y="7.14" fill="currentColor" rx=".5" transform="rotate(-90 17.434 9.14)"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crop(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.624 6.623H3.549a.5.5 0 0 1 0-1h2.075V3.55a.5.5 0 0 1 1 0v2.073h9.191a2.562 2.562 0 0 1 2.561 2.561v9.193h2.075a.5.5 0 0 1 0 1h-2.075v2.073a.5.5 0 0 1-1 0v-2.073H8.185a2.562 2.562 0 0 1-2.561-2.561zm11.752 10.754V8.184c0-.862-.699-1.561-1.561-1.561H6.624v9.193c0 .862.699 1.561 1.561 1.561z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dark(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.741 20.917a9.389 9.389 0 0 1-1.395-.105a9.141 9.141 0 0 1-1.465-17.7a1.177 1.177 0 0 1 1.21.281a1.273 1.273 0 0 1 .325 1.293a8.112 8.112 0 0 0-.353 2.68a8.266 8.266 0 0 0 4.366 6.857a7.628 7.628 0 0 0 3.711.993a1.242 1.242 0 0 1 .994 1.963a9.148 9.148 0 0 1-7.393 3.738M10.261 4.05a.211.211 0 0 0-.065.011a8.137 8.137 0 1 0 9.131 12.526a.224.224 0 0 0 .013-.235a.232.232 0 0 0-.206-.136a8.619 8.619 0 0 1-4.188-1.116a9.274 9.274 0 0 1-4.883-7.7a9.123 9.123 0 0 1 .4-3.008a.286.286 0 0 0-.069-.285a.184.184 0 0 0-.133-.057"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Database(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2.06c-3.53 0-6.18 1.23-6.18 2.86v14.16c0 1.63 2.65 2.86 6.18 2.86s6.18-1.23 6.18-2.86V4.92c0-1.63-2.66-2.86-6.18-2.86m5.18 17.02c0 .78-1.97 1.86-5.18 1.86s-5.18-1.08-5.18-1.86v-3.12A9.349 9.349 0 0 0 12 17.22a9.373 9.373 0 0 0 5.18-1.26Zm0-4.72c0 .78-1.97 1.86-5.18 1.86s-5.18-1.08-5.18-1.86v-3.12A9.349 9.349 0 0 0 12 12.5a9.373 9.373 0 0 0 5.18-1.26Zm0-4.72c0 .78-1.97 1.86-5.18 1.86s-5.18-1.08-5.18-1.86V6.52A9.349 9.349 0 0 0 12 7.78a9.373 9.373 0 0 0 5.18-1.26ZM12 6.78c-3.21 0-5.18-1.08-5.18-1.86S8.79 3.06 12 3.06s5.18 1.08 5.18 1.86S15.21 6.78 12 6.78"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeliveryTruck(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.47 11.185l-1.03-1.43a2.5 2.5 0 0 0-2.03-1.05h-4.38v-2.14a2.5 2.5 0 0 0-2.5-2.5H4.56a2.507 2.507 0 0 0-2.5 2.5v9.94a1.5 1.5 0 0 0 1.5 1.5h1.22a2.242 2.242 0 0 0 4.44 0h5.56a2.242 2.242 0 0 0 4.44 0h1.22a1.5 1.5 0 0 0 1.5-1.5v-3.87a2.508 2.508 0 0 0-.47-1.45M7 18.935a1.25 1.25 0 1 1 1.25-1.25A1.25 1.25 0 0 1 7 18.935m6.03-1.93H9.15a2.257 2.257 0 0 0-4.3 0H3.56a.5.5 0 0 1-.5-.5v-9.94a1.5 1.5 0 0 1 1.5-1.5h6.97a1.5 1.5 0 0 1 1.5 1.5Zm3.97 1.93a1.25 1.25 0 1 1 1.25-1.25a1.25 1.25 0 0 1-1.25 1.25m3.94-2.43a.5.5 0 0 1-.5.5h-1.29a2.257 2.257 0 0 0-4.3 0h-.82v-7.3h4.38a1.516 1.516 0 0 1 1.22.63l1.03 1.43a1.527 1.527 0 0 1 .28.87Z"/><path fill="currentColor" d="M18.029 12.205h-2a.5.5 0 0 1 0-1h2a.5.5 0 0 1 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Desktop(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.435 3.065H5.565a2.5 2.5 0 0 0-2.5 2.5v8.87a2.5 2.5 0 0 0 2.5 2.5h2.91l-.37 3H7a.5.5 0 0 0 0 1h10.01a.5.5 0 0 0 0-1H15.9l-.37-3h2.91a2.5 2.5 0 0 0 2.5-2.5v-8.87a2.5 2.5 0 0 0-2.505-2.5m-9.33 16.87l.38-3h5.03l.37 3Zm10.83-5.5a1.5 1.5 0 0 1-1.5 1.5H5.565a1.5 1.5 0 0 1-1.5-1.5v-.5h15.87Zm0-1.5H4.065v-7.37a1.5 1.5 0 0 1 1.5-1.5h12.87a1.5 1.5 0 0 1 1.5 1.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DesktopMouseOne(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.435 2.065h-2.87a6.5 6.5 0 0 0-6.5 6.5v6.87a6.5 6.5 0 0 0 6.5 6.5h2.87a6.5 6.5 0 0 0 6.5-6.5v-6.87a6.5 6.5 0 0 0-6.5-6.5m-8.37 6.5a5.51 5.51 0 0 1 5.5-5.5h.94v6.44h-6.44Zm13.87 6.87a5.5 5.5 0 0 1-5.5 5.5h-2.87a5.5 5.5 0 0 1-5.5-5.5v-4.93h13.87Zm0-5.93h-6.43v-6.44h.93a5.5 5.5 0 0 1 5.5 5.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DesktopMouseTwo(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.437 21.938h-2.874a6.508 6.508 0 0 1-6.5-6.5V8.562a6.508 6.508 0 0 1 6.5-6.5h2.874a6.508 6.508 0 0 1 6.5 6.5v6.876a6.508 6.508 0 0 1-6.5 6.5M10.563 3.062a5.506 5.506 0 0 0-5.5 5.5v6.876a5.506 5.506 0 0 0 5.5 5.5h2.874a5.506 5.506 0 0 0 5.5-5.5V8.562a5.506 5.506 0 0 0-5.5-5.5Z"/><path fill="currentColor" d="M11.5 6.541v4a.5.5 0 0 0 1 0v-4a.5.5 0 0 0-1 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiscountOne(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.953 12c0 .591-.346 1.124-.839 1.61c-.295.29-.639.568-.942.85c-.242.225-.46.446-.562.692c-.107.257-.114.576-.105.913c.011.416.056.855.059 1.265c.006.691-.123 1.304-.526 1.708c-.404.403-1.017.532-1.708.526c-.41-.004-.849-.048-1.264-.059c-.337-.009-.657-.002-.914.105c-.246.102-.467.32-.692.562c-.282.303-.56.647-.85.941c-.486.494-1.019.84-1.61.84c-.591 0-1.124-.346-1.61-.84c-.29-.294-.568-.638-.85-.941c-.225-.242-.447-.46-.692-.562c-.257-.107-.577-.114-.913-.105c-.416.011-.855.055-1.265.059c-.691.006-1.305-.123-1.708-.526c-.404-.404-.532-1.017-.526-1.708c.003-.41.048-.849.059-1.265c.009-.337.002-.656-.105-.914c-.102-.245-.32-.466-.562-.691c-.302-.282-.646-.56-.941-.85c-.493-.486-.84-1.019-.84-1.61c0-.591.347-1.124.84-1.61c.295-.29.639-.568.941-.85c.242-.225.46-.446.562-.691c.107-.258.114-.577.105-.914c-.011-.416-.056-.855-.059-1.265c-.006-.691.122-1.304.526-1.708c.403-.403 1.017-.532 1.708-.526c.41.004.849.048 1.265.059c.336.009.656.002.913-.105c.245-.102.467-.32.692-.562c.282-.303.56-.647.85-.941c.486-.494 1.019-.84 1.61-.84c.591 0 1.124.346 1.61.84c.29.294.568.638.85.941c.225.242.446.46.692.562c.257.107.577.114.914.105c.415-.011.854-.055 1.264-.059c.691-.006 1.304.123 1.708.526c.403.404.532 1.017.526 1.708c-.003.41-.048.849-.059 1.265c-.009.337-.002.656.105.913c.102.246.32.467.562.692c.303.282.647.56.942.85c.493.486.839 1.019.839 1.61m-1 0c0-.188-.088-.355-.206-.518c-.164-.226-.388-.437-.622-.646c-.583-.521-1.205-1.04-1.439-1.604c-.242-.585-.177-1.399-.136-2.178c.017-.315.027-.622-.015-.895c-.029-.191-.08-.365-.204-.489c-.125-.125-.299-.176-.49-.205c-.273-.042-.58-.032-.895-.015c-.779.041-1.593.106-2.177-.136c-.565-.234-1.084-.855-1.605-1.439c-.209-.234-.42-.458-.646-.622c-.163-.118-.33-.206-.518-.206c-.187 0-.355.088-.518.206c-.226.164-.437.388-.646.622c-.521.584-1.04 1.205-1.605 1.439c-.584.242-1.398.177-2.177.136c-.315-.017-.622-.027-.895.015c-.192.029-.365.08-.49.205c-.125.124-.175.298-.204.489c-.042.273-.032.58-.016.895c.042.779.107 1.593-.135 2.177c-.234.565-.855 1.084-1.439 1.605c-.234.209-.458.42-.622.646c-.118.163-.206.33-.206.518s.088.355.206.518c.164.226.388.437.622.646c.584.521 1.205 1.04 1.439 1.605c.242.584.177 1.398.135 2.177c-.016.315-.026.622.016.895c.029.191.079.365.204.489c.125.125.298.176.49.205c.273.042.58.032.895.015c.779-.041 1.593-.106 2.177.136c.565.234 1.084.855 1.605 1.439c.209.234.42.458.646.622c.163.118.331.206.518.206c.188 0 .355-.088.518-.206c.226-.164.437-.388.646-.622c.521-.584 1.04-1.205 1.605-1.439c.584-.242 1.398-.177 2.177-.136c.315.017.622.027.895-.015c.191-.029.365-.08.49-.205c.124-.124.175-.298.204-.489c.042-.273.032-.58.015-.895c-.041-.779-.106-1.593.136-2.178c.234-.564.856-1.083 1.439-1.604c.234-.209.458-.42.622-.646c.118-.163.206-.33.206-.518m-10.531-1.762a1.015 1.015 0 1 1-1.435-1.436a1.015 1.015 0 0 1 1.435 1.436M14.893 8.4a.5.5 0 0 1 .707.707L9.107 15.6a.5.5 0 0 1-.707-.707zm-1.315 5.363a1.015 1.015 0 1 1 1.435 1.436a1.015 1.015 0 0 1-1.435-1.436"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dollar(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21.934A9.934 9.934 0 1 1 21.934 12A9.945 9.945 0 0 1 12 21.934m0-18.868A8.934 8.934 0 1 0 20.934 12A8.944 8.944 0 0 0 12 3.066"/><path fill="currentColor" d="M14.5 13.5a2.006 2.006 0 0 1-2 2v1.01a.5.5 0 0 1-1 0V15.5h-1.25a.5.5 0 0 1 0-1h2.25a1 1 0 0 0 0-2h-1a2 2 0 0 1 0-4V7.49a.5.5 0 0 1 1 0V8.5h1.25a.5.5 0 0 1 0 1H11.5a1 1 0 0 0 0 2h1a2.006 2.006 0 0 1 2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Droplet(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21.936A7.27 7.27 0 0 1 7.656 20.5c-2.332-1.724-3.187-5.6-1.868-8.46l4.875-9.173A1.515 1.515 0 0 1 12 2.064a1.512 1.512 0 0 1 1.337.805l4.863 9.148c1.331 2.888.475 6.762-1.856 8.485A7.274 7.274 0 0 1 12 21.936m0-18.872a.51.51 0 0 0-.456.274l-4.861 9.147c-1.1 2.4-.376 5.777 1.568 7.212a6.4 6.4 0 0 0 7.5 0c1.942-1.435 2.668-4.817 1.554-7.237l-4.85-9.122A.507.507 0 0 0 12 3.064"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dumbbell(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.435 11.5h-.38V8.12a1.626 1.626 0 0 0-1.62-1.62h-.63v-.38a1.625 1.625 0 0 0-3.25 0v5.38h-7.11V6.12a1.625 1.625 0 0 0-3.25 0v.38h-.63a1.62 1.62 0 0 0-1.62 1.62v3.38h-.38a.5.5 0 1 0 0 1h.38v3.37a1.622 1.622 0 0 0 1.62 1.63H5.2v.37a1.625 1.625 0 1 0 3.25 0V12.5h7.11v5.37a1.625 1.625 0 1 0 3.25 0v-.37h.63a1.628 1.628 0 0 0 1.62-1.63V12.5h.38a.5.5 0 1 0 0-1ZM5.2 16.5h-.63a.625.625 0 0 1-.62-.63V8.12a.623.623 0 0 1 .62-.62h.63Zm2.25 1.37a.634.634 0 0 1-.63.63a.625.625 0 0 1-.62-.63V6.12a.623.623 0 0 1 .62-.62a.632.632 0 0 1 .63.62Zm10.36 0a.625.625 0 1 1-1.25 0V6.12a.625.625 0 0 1 1.25 0Zm2.25-2a.625.625 0 0 1-.62.63h-.63v-9h.63a.623.623 0 0 1 .62.62Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Edit(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.548 20.938h16.9a.5.5 0 0 0 0-1h-16.9a.5.5 0 0 0 0 1M9.71 17.18a2.587 2.587 0 0 0 1.12-.65l9.54-9.54a1.75 1.75 0 0 0 0-2.47l-.94-.93a1.788 1.788 0 0 0-2.47 0l-9.54 9.53a2.473 2.473 0 0 0-.64 1.12L6.04 17a.737.737 0 0 0 .19.72a.767.767 0 0 0 .53.22Zm.41-1.36a1.468 1.468 0 0 1-.67.39l-.97.26l-1-1l.26-.97a1.521 1.521 0 0 1 .39-.67l.38-.37l1.99 1.99Zm1.09-1.08l-1.99-1.99l6.73-6.73l1.99 1.99Zm8.45-8.45L18.65 7.3l-1.99-1.99l1.01-1.02a.748.748 0 0 1 1.06 0l.93.94a.754.754 0 0 1 0 1.06"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eraser(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.454 19.028h-7.01l6.62-6.63a2.935 2.935 0 0 0 .87-2.09a2.844 2.844 0 0 0-.87-2.05l-3.42-3.44a2.93 2.93 0 0 0-4.13.01L3.934 13.4a2.946 2.946 0 0 0 0 4.14l1.48 1.49h-1.86a.5.5 0 0 0 0 1h16.9a.5.5 0 0 0 0-1.002m-7.24-13.5a1.956 1.956 0 0 1 2.73 0l3.42 3.44a1.868 1.868 0 0 1 .57 1.35a1.93 1.93 0 0 1-.57 1.37l-5.64 5.64l-6.15-6.16Zm-1.19 13.5h-5.2l-2.18-2.2a1.931 1.931 0 0 1 0-2.72l2.23-2.23l6.15 6.15Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Export(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.552 20.968a2.577 2.577 0 0 1-2.5-2.73c-.012-2.153 0-4.306 0-6.459a.5.5 0 0 1 1 0c0 2.2-.032 4.4 0 6.6c.016 1.107.848 1.589 1.838 1.589h12.463A1.546 1.546 0 0 0 19.825 19a3.023 3.023 0 0 0 .1-1.061v-6.16a.5.5 0 0 1 1 0c0 2.224.085 4.465 0 6.687a2.567 2.567 0 0 1-2.67 2.5Z"/><path fill="currentColor" d="M12.337 3.176a.455.455 0 0 0-.311-.138c-.015 0-.028 0-.043-.006s-.027 0-.041.006a.457.457 0 0 0-.312.138L7.961 6.845a.5.5 0 0 0 .707.707l2.816-2.815v10.742a.5.5 0 0 0 1 0V4.737L15.3 7.552a.5.5 0 0 0 .707-.707Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceFrown(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21.942A9.942 9.942 0 1 1 21.942 12A9.953 9.953 0 0 1 12 21.942m0-18.884A8.942 8.942 0 1 0 20.942 12A8.952 8.952 0 0 0 12 3.058"/><path fill="currentColor" d="M17.206 16.481a6.033 6.033 0 0 0-10.412 0a.5.5 0 0 0 .863.5a5.033 5.033 0 0 1 8.685 0a.5.5 0 0 0 .864-.5"/><circle cx="9" cy="9.011" r="1.25" fill="currentColor"/><circle cx="15" cy="9.011" r="1.25" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceMeh(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21.942A9.942 9.942 0 1 1 21.942 12A9.953 9.953 0 0 1 12 21.942m0-18.884A8.942 8.942 0 1 0 20.942 12A8.952 8.952 0 0 0 12 3.058"/><circle cx="9.001" cy="8.99" r="1.25" fill="currentColor"/><circle cx="15.001" cy="8.99" r="1.25" fill="currentColor"/><path fill="currentColor" d="M8.438 15.939h7.125a.5.5 0 0 0 0-1H8.438a.5.5 0 0 0 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FaceSmile(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21.942A9.942 9.942 0 1 1 21.942 12A9.953 9.953 0 0 1 12 21.942m0-18.884A8.942 8.942 0 1 0 20.942 12A8.952 8.952 0 0 0 12 3.058"/><path fill="currentColor" d="M16.693 13.744a5.041 5.041 0 0 1-9.387 0c-.249-.59-1.111-.081-.863.505a6.026 6.026 0 0 0 11.114 0c.247-.586-.614-1.1-.864-.505"/><circle cx="9" cy="9.011" r="1.25" fill="currentColor"/><circle cx="15" cy="9.011" r="1.25" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Facebook(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.02 4.975A9.93 9.93 0 0 0 2.07 12A9.935 9.935 0 0 0 12 21.935a9.98 9.98 0 0 0 3.8-.75a10.189 10.189 0 0 0 3.22-2.16a9.934 9.934 0 0 0 0-14.05m-.7 13.34a8.921 8.921 0 0 1-5.32 2.57v-6.56h1.88a1 1 0 0 0 0-2H13v-2.74a1 1 0 0 1 1-1h1.2a1 1 0 0 0 0-2h-1.7a2.5 2.5 0 0 0-2.5 2.5v3.24H9.13a1 1 0 1 0 0 2H11v6.56a8.919 8.919 0 1 1 9.26-5.47a9.061 9.061 0 0 1-1.94 2.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileOff(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 3.308a.5.5 0 0 0-.7.71l.76.76v14.67a2.5 2.5 0 0 0 2.5 2.5h10.88a2.476 2.476 0 0 0 2.28-1.51l.28.28c.45.45 1.16-.26.7-.71Zm14.92 16.33a1.492 1.492 0 0 1-1.48 1.31H6.56a1.5 1.5 0 0 1-1.5-1.5V5.778Zm-5.54-16.55v2.92a2.5 2.5 0 0 0 2.5 2.5h3.07l-.01 6.7a.5.5 0 0 0 1 0v-6.67a2.057 2.057 0 0 0-.75-1.47c-1.3-1.26-2.59-2.53-3.89-3.8a3.924 3.924 0 0 0-1.41-1.13a6.523 6.523 0 0 0-1.71-.06H6.81a.5.5 0 0 0 0 1Zm4.83 4.42h-2.33a1.5 1.5 0 0 1-1.5-1.5v-2.24Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileOn(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m19.485 7.35l-4.97-4.86a1.466 1.466 0 0 0-1.05-.43h-6.9a2.5 2.5 0 0 0-2.5 2.5v14.88a2.507 2.507 0 0 0 2.5 2.5h10.87a2.507 2.507 0 0 0 2.5-2.5V8.42a1.49 1.49 0 0 0-.45-1.07m-1.27.15h-2.34a1.5 1.5 0 0 1-1.5-1.5V3.75Zm.72 11.94a1.5 1.5 0 0 1-1.5 1.5H6.565a1.5 1.5 0 0 1-1.5-1.5V4.56a1.5 1.5 0 0 1 1.5-1.5h6.81V6a2.5 2.5 0 0 0 2.5 2.5h3.06Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Filter(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.037 20.937a1.015 1.015 0 0 1-.518-.145l-3.334-2a2.551 2.551 0 0 1-1.233-2.176v-4.525a1.526 1.526 0 0 0-.284-.891L4.013 4.658a1.01 1.01 0 0 1 .822-1.6h14.33a1.009 1.009 0 0 1 .822 1.6L15.332 11.2a1.527 1.527 0 0 0-.285.891v7.834a1.013 1.013 0 0 1-1.01 1.012M4.835 4.063l4.647 6.557a2.515 2.515 0 0 1 .47 1.471v4.524a1.543 1.543 0 0 0 .747 1.318l3.334 2l.014-7.843a2.516 2.516 0 0 1 .471-1.471l4.654-6.542Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlagOne(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.565 3.18a.809.809 0 0 0-.81-.02l-1.13.56c-1.63.87-3.82.83-6.5-.13a9.141 9.141 0 0 0-7.3.52l-.76.41v-.96a.5.5 0 0 0-1 0v16.88a.5.5 0 0 0 1 0V15.9a.836.836 0 0 0 .2-.08l1.03-.55a8.163 8.163 0 0 1 6.5-.46c2.95 1.06 5.41 1.08 7.3.07l1.44-.72a.759.759 0 0 0 .4-.66V3.82a.751.751 0 0 0-.37-.64m-.63 10.16l-1.31.66c-1.63.87-3.82.83-6.5-.13a9.141 9.141 0 0 0-7.3.52l-.76.4V5.65L5.3 4.99a8.122 8.122 0 0 1 6.5-.46c2.95 1.06 5.41 1.08 7.29.08l.85-.43Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FloppyDisk(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.015 7.015l-4.15-3.39a2.54 2.54 0 0 0-1.58-.56h-9.72a1.5 1.5 0 0 0-1.5 1.5v14.87a1.5 1.5 0 0 0 1.5 1.5h14.87a1.5 1.5 0 0 0 1.5-1.5V8.955a2.507 2.507 0 0 0-.92-1.94m-13.45-2.95h5.75v1.37a.5.5 0 0 1-.5.5h-4.75a.5.5 0 0 1-.5-.5Zm0 15.87v-5.93a1.5 1.5 0 0 1 1.5-1.5h7.87a1.5 1.5 0 0 1 1.5 1.5v5.93Zm13.37-.5a.5.5 0 0 1-.5.5h-1v-5.93a2.507 2.507 0 0 0-2.5-2.5h-7.87a2.5 2.5 0 0 0-2.5 2.5v5.93h-1a.5.5 0 0 1-.5-.5V4.565a.5.5 0 0 1 .5-.5h1v1.37a1.5 1.5 0 0 0 1.5 1.5h4.75a1.5 1.5 0 0 0 1.5-1.5v-1.37h.97a1.514 1.514 0 0 1 .95.34l4.14 3.38a1.483 1.483 0 0 1 .56 1.17Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOff(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.013 3.3a.5.5 0 0 0-.711.71l.25.25a2.438 2.438 0 0 0-1.49 2.24v11a2.453 2.453 0 0 0 2.451 2.44h14.72l.759.76c.461.46 1.171-.25.711-.7Zm.5 15.64a1.45 1.45 0 0 1-1.451-1.44v-11a1.444 1.444 0 0 1 1.31-1.43c1.521 1.53 3.06 3.07 4.591 4.59q4.485 4.485 8.96 8.97l.31.31Zm16.925-1.247a.5.5 0 0 1-.5-.5V9.175a1.445 1.445 0 0 0-1.445-1.444h-6.666a1.5 1.5 0 0 1-1.474-1.225l-.05-.267a1.445 1.445 0 0 0-1.42-1.178H8.8a.5.5 0 0 1 0-1h1.083a2.446 2.446 0 0 1 2.4 1.994l.05.268a.5.5 0 0 0 .491.408h6.666a2.448 2.448 0 0 1 2.445 2.444v8.018a.5.5 0 0 1-.497.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOn(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.435 19.94H4.565a2.5 2.5 0 0 1-2.5-2.5V6.56a2.5 2.5 0 0 1 2.5-2.5h5.27a2.5 2.5 0 0 1 2.457 2.04l.042.222a.5.5 0 0 0 .491.408h6.61a2.5 2.5 0 0 1 2.5 2.5v8.21a2.5 2.5 0 0 1-2.5 2.5M4.565 5.06a1.5 1.5 0 0 0-1.5 1.5v10.88a1.5 1.5 0 0 0 1.5 1.5h14.87a1.5 1.5 0 0 0 1.5-1.5V9.23a1.5 1.5 0 0 0-1.5-1.5h-6.61a1.5 1.5 0 0 1-1.474-1.225l-.042-.221A1.5 1.5 0 0 0 9.835 5.06Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Football(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.278 4.757a1.64 1.64 0 0 0-1.03-1.04a12.248 12.248 0 0 0-15.53 15.53a1.64 1.64 0 0 0 1.04 1.03a12.306 12.306 0 0 0 3.95.66a12.262 12.262 0 0 0 11.57-16.18m-15.2 14.58a.725.725 0 0 1-.42-.42a11.379 11.379 0 0 1-.58-4.26l5.26 5.26a11.352 11.352 0 0 1-4.26-.58m11.56-2.71a11.179 11.179 0 0 1-6.03 3.14l-6.38-6.38a11.083 11.083 0 0 1 3.14-6.02a11.193 11.193 0 0 1 6.03-3.15l6.38 6.38a11.245 11.245 0 0 1-3.14 6.03m3.29-7.3l-5.26-5.26c.21 0 .41-.01.62-.01a11.154 11.154 0 0 1 3.63.61a.682.682 0 0 1 .42.41a11.543 11.543 0 0 1 .59 4.25"/><path fill="currentColor" d="M10.4 15.257a.5.5 0 0 0 .35.15a.508.508 0 0 0 .36-.15a.5.5 0 0 0 0-.7l-.48-.48l1.37-1.37l.48.48a.518.518 0 0 0 .35.14a.543.543 0 0 0 .36-.14a.513.513 0 0 0 0-.71l-.48-.48l1.37-1.37l.48.48a.5.5 0 0 0 .7-.71l-1.66-1.66a.5.5 0 0 0-.71 0a.5.5 0 0 0 0 .7l.49.49L12 11.3l-.48-.48a.495.495 0 1 0-.7.7l.48.48l-1.37 1.38l-.49-.49a.5.5 0 0 0-.7 0a.5.5 0 0 0 0 .71Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ForkKnife(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.83 2.25a6.469 6.469 0 0 0-4.37 6.12v3.15a1.784 1.784 0 0 0 1.78 1.78h2.7v8.14a.5.5 0 0 0 .5.5a.508.508 0 0 0 .5-.5V2.56a.508.508 0 0 0-.5-.5a.467.467 0 0 0-.17.03ZM16.24 12.3a.781.781 0 0 1-.78-.78V8.37a5.482 5.482 0 0 1 3.48-5.1v9.03Zm-4.8-10.23a.5.5 0 0 0-.5.5v4.98H8.5V2.57a.5.5 0 0 0-.5-.5a.5.5 0 0 0-.5.5v4.98H5.06V2.57a.5.5 0 0 0-.5-.5a.5.5 0 0 0-.5.5v6.48a2.507 2.507 0 0 0 2.5 2.5h.94v9.89a.508.508 0 0 0 .5.5a.5.5 0 0 0 .5-.5v-9.89h.94a2.5 2.5 0 0 0 2.5-2.5V2.57a.5.5 0 0 0-.5-.5m-.5 6.98a1.5 1.5 0 0 1-1.5 1.5H6.56a1.511 1.511 0 0 1-1.5-1.5v-.5h5.88Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fries(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.51 9.535a1.091 1.091 0 0 0-.81-.36h-1.03a.5.5 0 0 0-.17.02v-4.38a1.5 1.5 0 0 0-1.5-1.5h-.5a1.3 1.3 0 0 0-.52.1a1.474 1.474 0 0 0-1.48-1.35H13a1.5 1.5 0 0 0-1.5 1.5v1.59a1.386 1.386 0 0 0-.5-.09h-.5a1.348 1.348 0 0 0-.5.09v-.34a1.5 1.5 0 0 0-1.5-1.5H8a1.5 1.5 0 0 0-1.5 1.5V9.2a.5.5 0 0 0-.17-.02H5.3a1.1 1.1 0 0 0-1.08 1.2l.85 8.98a2.84 2.84 0 0 0 2.84 2.58h8.18a2.84 2.84 0 0 0 2.84-2.58l.85-8.98a1.112 1.112 0 0 0-.27-.845M15 4.815a.5.5 0 0 1 .5-.5h.5a.5.5 0 0 1 .5.5v6.14h.01a2.915 2.915 0 0 1-1.51 2.06Zm-2.5 8.53v-9.78a.5.5 0 0 1 .5-.5h.5a.5.5 0 0 1 .5.5v9.76Zm-2.5-.02v-6.76a.5.5 0 0 1 .5-.5h.5a.5.5 0 0 1 .5.5v6.78Zm-2.5-8.51a.5.5 0 0 1 .5-.5h.5a.5.5 0 0 1 .5.5v8.2a2.877 2.877 0 0 1-1.5-2.06Zm11.29 5.391l-.85 9.049a1.85 1.85 0 0 1-1.85 1.68H7.91a1.84 1.84 0 0 1-1.84-1.68l-.86-9.08h1.12a.1.1 0 0 1 .09.08l.05.56a3.891 3.891 0 0 0 3.88 3.53h3.3a3.884 3.884 0 0 0 3.88-3.53l.05-.56a.106.106 0 0 1 .09-.08Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gift(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.435 5.568h-2.38a1.979 1.979 0 0 0-.51-1.92a2.022 2.022 0 0 0-2.83 0L12 5.367l-1.71-1.719a2 2 0 0 0-2.83 0a1.979 1.979 0 0 0-.51 1.92H4.565a1.5 1.5 0 0 0-1.5 1.5v1A1.487 1.487 0 0 0 4 9.448v8.99a2.507 2.507 0 0 0 2.5 2.5h11a2.5 2.5 0 0 0 2.5-2.5v-8.98a1.509 1.509 0 0 0 .94-1.39v-1a1.5 1.5 0 0 0-1.505-1.5M8.165 4.357a1 1 0 0 1 1.41 0l1.21 1.211h-2.77a.989.989 0 0 1 .15-1.211M11 19.938H6.5a1.5 1.5 0 0 1-1.5-1.5v-8.87h6Zm0-11.37H4.565a.5.5 0 0 1-.5-.5v-1a.5.5 0 0 1 .5-.5H11Zm3.43-4.211A1 1 0 0 1 16 5.568h-2.78ZM19 18.438a1.5 1.5 0 0 1-1.5 1.5H13V9.568h6Zm.94-10.37a.5.5 0 0 1-.5.5H13v-2h6.44a.5.5 0 0 1 .5.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Glass(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.279 2.54a1.475 1.475 0 0 0-1.1-.48H6.819a1.47 1.47 0 0 0-1.09.48a1.5 1.5 0 0 0-.41 1.12l1.06 15.94a2.51 2.51 0 0 0 2.49 2.34h6.26a2.519 2.519 0 0 0 2.5-2.34l1.05-15.94a1.5 1.5 0 0 0-.4-1.12m-1.65 16.99a1.508 1.508 0 0 1-1.5 1.41h-6.26a1.506 1.506 0 0 1-1.49-1.41l-.64-9.62a2.981 2.981 0 0 0 1.17-.49a1.828 1.828 0 0 1 1.18-.39a1.858 1.858 0 0 1 1.19.39a3.025 3.025 0 0 0 3.45 0a1.879 1.879 0 0 1 1.19-.39a1.828 1.828 0 0 1 1.18.39a3 3 0 0 0 1.16.49Zm.7-10.62a2.317 2.317 0 0 1-.69-.33a2.98 2.98 0 0 0-3.45 0a1.885 1.885 0 0 1-1.18.38a1.939 1.939 0 0 1-1.19-.38a2.818 2.818 0 0 0-1.73-.55a2.809 2.809 0 0 0-1.72.55a2.374 2.374 0 0 1-.7.33l-.35-5.32a.468.468 0 0 1 .14-.37a.484.484 0 0 1 .36-.16h10.36a.523.523 0 0 1 .37.16a.46.46 0 0 1 .13.37Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Globe(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.645 2.428a8.1 8.1 0 0 0-1.61-.3a9.332 9.332 0 0 0-3.6.28l-.07.02a9.928 9.928 0 0 0 .01 19.15a9.091 9.091 0 0 0 2.36.34a1.274 1.274 0 0 0 .27.02a9.65 9.65 0 0 0 2.63-.36a9.931 9.931 0 0 0 .01-19.15m-.27.96a8.943 8.943 0 0 1 5.84 5.11h-4.26a13.778 13.778 0 0 0-2.74-5.35a8.254 8.254 0 0 1 1.16.24m-2.37-.09a12.78 12.78 0 0 1 2.91 5.2h-5.84a12.545 12.545 0 0 1 2.93-5.198Zm3.16 6.2a13.193 13.193 0 0 1 0 5.01h-6.32a12.185 12.185 0 0 1-.25-2.5a12.353 12.353 0 0 1 .25-2.51Zm-5.6-6.09l.07-.02a9.152 9.152 0 0 1 1.16-.23A13.618 13.618 0 0 0 8.045 8.5H3.8a9 9 0 0 1 5.765-5.092m-6.5 8.6a8.71 8.71 0 0 1 .37-2.51h4.39a13.95 13.95 0 0 0-.23 2.51a13.757 13.757 0 0 0 .23 2.5h-4.39a8.591 8.591 0 0 1-.37-2.5m6.57 8.61a8.9 8.9 0 0 1-5.84-5.11h4.24a13.632 13.632 0 0 0 2.77 5.35a8.1 8.1 0 0 1-1.17-.24m-.56-5.11h5.84a12.638 12.638 0 0 1-2.91 5.21a12.872 12.872 0 0 1-2.93-5.21m5.3 5.11a11.551 11.551 0 0 1-1.17.24a13.8 13.8 0 0 0 2.75-5.35h4.26a8.924 8.924 0 0 1-5.84 5.11m1.8-6.11a13.611 13.611 0 0 0 0-5.01h4.39a8.379 8.379 0 0 1 .37 2.51a8.687 8.687 0 0 1-.36 2.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gps(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 14.5a2.5 2.5 0 1 1 2.5-2.5a2.5 2.5 0 0 1-2.5 2.5m0-4a1.5 1.5 0 1 0 1.5 1.5a1.5 1.5 0 0 0-1.5-1.5"/><path fill="currentColor" d="M21.435 11.505h-1.46a7.98 7.98 0 0 0-7.48-7.48v-1.46a.508.508 0 0 0-.5-.5a.515.515 0 0 0-.5.5v1.46a8 8 0 0 0-7.48 7.48h-1.45a.5.5 0 1 0 0 1h1.45a8.012 8.012 0 0 0 7.48 7.48v1.45a.508.508 0 0 0 .5.5a.5.5 0 0 0 .5-.5v-1.45a8 8 0 0 0 7.48-7.48h1.46a.5.5 0 0 0 0-1M12 19.005a7 7 0 1 1 7-7a7.021 7.021 0 0 1-7 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GridFourOne(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 11H5.563a2.5 2.5 0 0 1-2.5-2.5V5.564a2.5 2.5 0 0 1 2.5-2.5H8.5a2.5 2.5 0 0 1 2.5 2.5V8.5A2.5 2.5 0 0 1 8.5 11M5.563 4.064a1.5 1.5 0 0 0-1.5 1.5V8.5a1.5 1.5 0 0 0 1.5 1.5H8.5A1.5 1.5 0 0 0 10 8.5V5.564a1.5 1.5 0 0 0-1.5-1.5ZM18.436 11H15.5A2.5 2.5 0 0 1 13 8.5V5.564a2.5 2.5 0 0 1 2.5-2.5h2.934a2.5 2.5 0 0 1 2.5 2.5V8.5a2.5 2.5 0 0 1-2.498 2.5M15.5 4.064a1.5 1.5 0 0 0-1.5 1.5V8.5a1.5 1.5 0 0 0 1.5 1.5h2.934a1.5 1.5 0 0 0 1.5-1.5V5.564a1.5 1.5 0 0 0-1.5-1.5Zm-7 16.872H5.564a2.5 2.5 0 0 1-2.5-2.5V15.5a2.5 2.5 0 0 1 2.5-2.5H8.5a2.5 2.5 0 0 1 2.5 2.5v2.936a2.5 2.5 0 0 1-2.5 2.5M5.564 14a1.5 1.5 0 0 0-1.5 1.5v2.936a1.5 1.5 0 0 0 1.5 1.5H8.5a1.5 1.5 0 0 0 1.5-1.5V15.5A1.5 1.5 0 0 0 8.5 14Zm12.872 6.936H15.5a2.5 2.5 0 0 1-2.5-2.5V15.5a2.5 2.5 0 0 1 2.5-2.5h2.934a2.5 2.5 0 0 1 2.5 2.5v2.936a2.5 2.5 0 0 1-2.498 2.5M15.5 14a1.5 1.5 0 0 0-1.5 1.5v2.936a1.5 1.5 0 0 0 1.5 1.5h2.934a1.5 1.5 0 0 0 1.5-1.5V15.5a1.5 1.5 0 0 0-1.5-1.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GridFourTwo(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 13.933H5.563a2.5 2.5 0 0 1-2.5-2.5v-5.87a2.5 2.5 0 0 1 2.5-2.5H8.5a2.5 2.5 0 0 1 2.5 2.5v5.87a2.5 2.5 0 0 1-2.5 2.5m-2.937-9.87a1.5 1.5 0 0 0-1.5 1.5v5.87a1.5 1.5 0 0 0 1.5 1.5H8.5a1.5 1.5 0 0 0 1.5-1.5v-5.87a1.5 1.5 0 0 0-1.5-1.5ZM8.5 20.935H5.564a2.5 2.5 0 0 1 0-5H8.5a2.5 2.5 0 1 1 0 5m-2.934-4a1.5 1.5 0 0 0 0 3H8.5a1.5 1.5 0 1 0 0-3Zm12.87 4H15.5a2.5 2.5 0 0 1-2.5-2.5v-5.87a2.5 2.5 0 0 1 2.5-2.5h2.934a2.5 2.5 0 0 1 2.5 2.5v5.87a2.5 2.5 0 0 1-2.498 2.5m-2.936-9.87a1.5 1.5 0 0 0-1.5 1.5v5.87a1.5 1.5 0 0 0 1.5 1.5h2.934a1.5 1.5 0 0 0 1.5-1.5v-5.87a1.5 1.5 0 0 0-1.5-1.5Zm2.936-3.002H15.5a2.5 2.5 0 0 1 0-5h2.934a2.5 2.5 0 0 1 0 5Zm-2.934-4a1.5 1.5 0 0 0 0 3h2.934a1.5 1.5 0 0 0 0-3Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GridThreeOne(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.434 20.936H5.563a2.5 2.5 0 0 1-2.5-2.5V15.5a2.5 2.5 0 0 1 2.5-2.5h5.871a2.5 2.5 0 0 1 2.5 2.5v2.933a2.5 2.5 0 0 1-2.5 2.503M5.563 14a1.5 1.5 0 0 0-1.5 1.5v2.933a1.5 1.5 0 0 0 1.5 1.5h5.871a1.5 1.5 0 0 0 1.5-1.5V15.5a1.5 1.5 0 0 0-1.5-1.5Zm12.872 6.936a2.5 2.5 0 0 1-2.5-2.5V15.5a2.5 2.5 0 0 1 5 0v2.934a2.5 2.5 0 0 1-2.5 2.502m0-6.934a1.5 1.5 0 0 0-1.5 1.5v2.934a1.5 1.5 0 0 0 3 0V15.5a1.5 1.5 0 0 0-1.5-1.5Zm0-3.002H5.563a2.5 2.5 0 0 1-2.5-2.5V5.564a2.5 2.5 0 0 1 2.5-2.5h12.872a2.5 2.5 0 0 1 2.5 2.5V8.5a2.5 2.5 0 0 1-2.5 2.5M5.563 4.064a1.5 1.5 0 0 0-1.5 1.5V8.5a1.5 1.5 0 0 0 1.5 1.5h12.872a1.5 1.5 0 0 0 1.5-1.5V5.564a1.5 1.5 0 0 0-1.5-1.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GridThreeTwo(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.437 11h-5.871a2.5 2.5 0 0 1-2.5-2.5V5.564a2.5 2.5 0 0 1 2.5-2.5h5.871a2.5 2.5 0 0 1 2.5 2.5V8.5a2.5 2.5 0 0 1-2.5 2.5m-5.871-6.936a1.5 1.5 0 0 0-1.5 1.5V8.5a1.5 1.5 0 0 0 1.5 1.5h5.871a1.5 1.5 0 0 0 1.5-1.5V5.564a1.5 1.5 0 0 0-1.5-1.5ZM5.565 11a2.5 2.5 0 0 1-2.5-2.5V5.564a2.5 2.5 0 1 1 5 0V8.5a2.5 2.5 0 0 1-2.5 2.5m0-6.934a1.5 1.5 0 0 0-1.5 1.5V8.5a1.5 1.5 0 0 0 3 0V5.564a1.5 1.5 0 0 0-1.5-1.5Zm12.872 16.87H5.565a2.5 2.5 0 0 1-2.5-2.5V15.5a2.5 2.5 0 0 1 2.5-2.5h12.872a2.5 2.5 0 0 1 2.5 2.5v2.934a2.5 2.5 0 0 1-2.5 2.502M5.565 14a1.5 1.5 0 0 0-1.5 1.5v2.934a1.5 1.5 0 0 0 1.5 1.5h12.872a1.5 1.5 0 0 0 1.5-1.5V15.5a1.5 1.5 0 0 0-1.5-1.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GridTwoH(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.437 11H5.565a2.5 2.5 0 0 1-2.5-2.5V5.564a2.5 2.5 0 0 1 2.5-2.5h12.872a2.5 2.5 0 0 1 2.5 2.5V8.5a2.5 2.5 0 0 1-2.5 2.5M5.565 4.064a1.5 1.5 0 0 0-1.5 1.5V8.5a1.5 1.5 0 0 0 1.5 1.5h12.872a1.5 1.5 0 0 0 1.5-1.5V5.564a1.5 1.5 0 0 0-1.5-1.5Zm12.872 16.872H5.565a2.5 2.5 0 0 1-2.5-2.5V15.5a2.5 2.5 0 0 1 2.5-2.5h12.872a2.5 2.5 0 0 1 2.5 2.5v2.934a2.5 2.5 0 0 1-2.5 2.502M5.565 14a1.5 1.5 0 0 0-1.5 1.5v2.934a1.5 1.5 0 0 0 1.5 1.5h12.872a1.5 1.5 0 0 0 1.5-1.5V15.5a1.5 1.5 0 0 0-1.5-1.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GridTwoV(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.436 20.937H15.5a2.5 2.5 0 0 1-2.5-2.5V5.565a2.5 2.5 0 0 1 2.5-2.5h2.933a2.5 2.5 0 0 1 2.5 2.5v12.872a2.5 2.5 0 0 1-2.497 2.5M15.5 4.065a1.5 1.5 0 0 0-1.5 1.5v12.872a1.5 1.5 0 0 0 1.5 1.5h2.933a1.5 1.5 0 0 0 1.5-1.5V5.565a1.5 1.5 0 0 0-1.5-1.5Zm-7 16.872H5.564a2.5 2.5 0 0 1-2.5-2.5V5.565a2.5 2.5 0 0 1 2.5-2.5H8.5a2.5 2.5 0 0 1 2.5 2.5v12.872a2.5 2.5 0 0 1-2.5 2.5M5.564 4.065a1.5 1.5 0 0 0-1.5 1.5v12.872a1.5 1.5 0 0 0 1.5 1.5H8.5a1.5 1.5 0 0 0 1.5-1.5V5.565a1.5 1.5 0 0 0-1.5-1.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HardDrive(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.905 14.325l-1.83-10.04a1.507 1.507 0 0 0-1.47-1.22h-11.2A1.493 1.493 0 0 0 4.925 4.3l-1.84 10.03a2.452 2.452 0 0 0-.02.27v4.84a1.5 1.5 0 0 0 1.5 1.5h14.87a1.511 1.511 0 0 0 1.5-1.5V14.6a1.241 1.241 0 0 0-.03-.275m-15-9.85a.5.5 0 0 1 .5-.41h11.2a.511.511 0 0 1 .49.4l1.74 9.54H4.165Zm14.03 14.96a.5.5 0 0 1-.5.5H4.565a.5.5 0 0 1-.5-.5l.01-4.43h15.86Z"/><circle cx="5.561" cy="17.47" r=".5" fill="currentColor"/><circle cx="7.561" cy="17.47" r=".5" fill="currentColor"/><path fill="currentColor" d="M18.45 17.97a.5.5 0 0 0 0-1h-5a.5.5 0 0 0 0 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hashtag(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.435 15.506H16.2l.61-7h3.63a.5.5 0 0 0 .5-.5a.5.5 0 0 0-.5-.5H16.9l.34-3.87a.509.509 0 0 0-.46-.54a.5.5 0 0 0-.54.46l-.35 3.95H8.9l.34-3.87a.509.509 0 0 0-.46-.54a.491.491 0 0 0-.54.46l-.35 3.95H3.565a.5.5 0 0 0-.5.5a.5.5 0 0 0 .5.5h4.24l-.62 7h-3.62a.5.5 0 0 0-.5.5a.5.5 0 0 0 .5.5h3.54l-.34 3.86a.508.508 0 0 0 .45.54h.05a.516.516 0 0 0 .5-.46l.34-3.94h7l-.34 3.86a.508.508 0 0 0 .45.54h.05a.516.516 0 0 0 .5-.46l.34-3.94h4.33a.5.5 0 0 0 .5-.5a.5.5 0 0 0-.5-.5m-5.25 0H8.2l.61-7h7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Headphones(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.23 3.065h-.46a8.211 8.211 0 0 0-8.2 8.2v6.25a3.385 3.385 0 0 0 .89 2.3a3.423 3.423 0 0 0 2.53 1.12h.53a1.225 1.225 0 0 0 1.22-1.22v-4.4A1.225 1.225 0 0 0 7.52 14.1h-.41a3.6 3.6 0 0 0-2.54 1.05v-3.88a7.208 7.208 0 0 1 7.2-7.2h.46a7.208 7.208 0 0 1 7.2 7.2v3.88a3.6 3.6 0 0 0-2.54-1.05h-.41a1.225 1.225 0 0 0-1.22 1.22v4.4a1.225 1.225 0 0 0 1.22 1.22h.53a3.423 3.423 0 0 0 2.53-1.12a3.385 3.385 0 0 0 .89-2.3v-6.25a8.211 8.211 0 0 0-8.2-8.205m-7.65 14.21A2.511 2.511 0 0 1 7.11 15.1h.41a.222.222 0 0 1 .22.22v4.4a.222.222 0 0 1-.22.22h-.53a2.422 2.422 0 0 1-1.79-.79a2.322 2.322 0 0 1-.63-1.63a1.927 1.927 0 0 1 .01-.245m14.22 1.87a2.422 2.422 0 0 1-1.79.79h-.53a.222.222 0 0 1-.22-.22v-4.4a.222.222 0 0 1 .22-.22h.41a2.511 2.511 0 0 1 2.53 2.18a1.927 1.927 0 0 1 .01.24a2.322 2.322 0 0 1-.63 1.63"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heart(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 20.043a.977.977 0 0 1-.7-.288L4.63 13.08a5.343 5.343 0 0 1 1.423-8.567A5.266 5.266 0 0 1 12 5.371a5.272 5.272 0 0 1 5.947-.858a5.343 5.343 0 0 1 1.423 8.567l-6.676 6.675a.977.977 0 0 1-.694.288M8.355 4.963a4.015 4.015 0 0 0-1.844.437a4.4 4.4 0 0 0-2.389 3.243a4.345 4.345 0 0 0 1.215 3.73l6.675 6.675l6.651-6.675a4.345 4.345 0 0 0 1.215-3.73A4.4 4.4 0 0 0 17.489 5.4a4.338 4.338 0 0 0-4.968.852a.744.744 0 0 1-1.042 0a4.474 4.474 0 0 0-3.124-1.289"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Home(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m19.37 10.22l-6.2-7.6a1.5 1.5 0 0 0-2.33-.01l-6.21 7.61a2.5 2.5 0 0 0-.57 1.59v7.63a2.507 2.507 0 0 0 2.5 2.5h10.88a2.507 2.507 0 0 0 2.5-2.5v-7.63a2.5 2.5 0 0 0-.57-1.59M10 20.94v-5.5a1.5 1.5 0 0 1 1.5-1.5h1a1.5 1.5 0 0 1 1.5 1.5v5.5Zm8.94-1.5a1.511 1.511 0 0 1-1.5 1.5H15v-5.5a2.5 2.5 0 0 0-2.5-2.5h-1a2.5 2.5 0 0 0-2.5 2.5v5.5H6.56a1.511 1.511 0 0 1-1.5-1.5v-7.63a1.474 1.474 0 0 1 .34-.95l6.22-7.61a.474.474 0 0 1 .38-.19a.479.479 0 0 1 .39.19l6.21 7.61a1.474 1.474 0 0 1 .34.95Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HospitalOne(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.435 7.5h-1.93V5.56a2.5 2.5 0 0 0-2.5-2.5h-4a2.5 2.5 0 0 0-2.5 2.5V7.5h-1.94a2.5 2.5 0 0 0-2.5 2.5v9.44a1.511 1.511 0 0 0 1.5 1.5h14.87a1.5 1.5 0 0 0 1.5-1.5V10a2.5 2.5 0 0 0-2.5-2.5M7.505 19.94h-2.94a.508.508 0 0 1-.5-.5V10a1.5 1.5 0 0 1 1.5-1.5h1.94Zm8 0h-1.5v-2.5a2.038 2.038 0 0 0-.59-1.42a2 2 0 0 0-3.41 1.42v2.5h-1.5V5.56a1.5 1.5 0 0 1 1.5-1.5h4a1.5 1.5 0 0 1 1.5 1.5Zm4.43-.5a.5.5 0 0 1-.5.5h-2.93V8.5h1.93a1.5 1.5 0 0 1 1.5 1.5Z"/><path fill="currentColor" d="M14.505 8.56a.5.5 0 0 1-.5.5h-1.5v1.5a.5.5 0 0 1-.5.5a.5.5 0 0 1-.5-.5v-1.5h-1.5a.5.5 0 0 1-.5-.5a.5.5 0 0 1 .5-.5h1.5v-1.5a.5.5 0 0 1 .5-.5a.508.508 0 0 1 .5.5v1.5h1.5a.508.508 0 0 1 .5.5m-8.719 4.657a.5.5 0 0 1-.5-.5v-1.5a.5.5 0 0 1 1 0v1.5a.5.5 0 0 1-.5.5m0 4.5a.5.5 0 0 1-.5-.5v-1.5a.5.5 0 0 1 1 0v1.5a.5.5 0 0 1-.5.5m12.435-4.5a.5.5 0 0 1-.5-.5v-1.5a.5.5 0 0 1 1 0v1.5a.5.5 0 0 1-.5.5m0 4.5a.5.5 0 0 1-.5-.5v-1.5a.5.5 0 0 1 1 0v1.5a.5.5 0 0 1-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hotdog(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m19.1 9.349l-9.74 9.74l.01.01l9.74-9.74Z"/><path fill="currentColor" d="m20.276 9.119l-.47-.47a3.157 3.157 0 0 0-.03-4.43a3.212 3.212 0 0 0-4.42-.03l-.48-.48a2.3 2.3 0 0 0-3.18 0l-7.98 7.98a2.263 2.263 0 0 0 0 3.18l.48.48a3.145 3.145 0 0 0 .03 4.42a3.089 3.089 0 0 0 2.23.92a3.126 3.126 0 0 0 2.2-.89l.47.47a2.245 2.245 0 0 0 3.18 0l7.97-7.97a2.245 2.245 0 0 0 0-3.18m-15.85 3.27l7.97-7.97a1.243 1.243 0 0 1 1.77 0l.47.47l-9.736 9.74l-.47-.47a1.249 1.249 0 0 1-.004-1.77m3.52 6.7a2.2 2.2 0 0 1-3.02-.03a2.149 2.149 0 0 1-.03-3.01l11.16-11.16a2.163 2.163 0 0 1 1.49-.6a2.155 2.155 0 0 1 1.55 3.65Zm11.63-7.49l-7.98 7.97a1.275 1.275 0 0 1-1.76 0l-.47-.47l-.01-.01l9.74-9.74l.01.01l.47.47a1.268 1.268 0 0 1 0 1.771Z"/><path fill="currentColor" d="M6.57 17.569a.5.5 0 0 1-.354-.854a4.533 4.533 0 0 1 1.357-.967a3.491 3.491 0 0 0 1.1-.778a3.514 3.514 0 0 0 .779-1.1a5.034 5.034 0 0 1 2.324-2.324a3.517 3.517 0 0 0 1.1-.78a3.536 3.536 0 0 0 .78-1.1a4.534 4.534 0 0 1 .97-1.359a4.54 4.54 0 0 1 1.359-.97a3.53 3.53 0 0 0 1.1-.78a.5.5 0 1 1 .707.707a4.516 4.516 0 0 1-1.36.969a3.506 3.506 0 0 0-1.1.781a3.535 3.535 0 0 0-.781 1.1a4.516 4.516 0 0 1-.969 1.36a4.5 4.5 0 0 1-1.359.969a4.029 4.029 0 0 0-1.874 1.874a4.5 4.5 0 0 1-.967 1.357a4.524 4.524 0 0 1-1.358.968a3.51 3.51 0 0 0-1.1.777a.5.5 0 0 1-.354.15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IceCream(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.54 5.94a4.594 4.594 0 0 0-9.08 0a3.065 3.065 0 0 0-.76 5.85l3.92 9.25a1.5 1.5 0 0 0 2.76 0l3.92-9.26a3.058 3.058 0 0 0-.76-5.84m-4.08 14.71a.5.5 0 0 1-.92 0l-3.65-8.62h8.22Zm3.64-9.62H7.9a2.06 2.06 0 1 1 .01-4.12a.5.5 0 0 0 .5-.48a3.6 3.6 0 0 1 7.18 0a.506.506 0 0 0 .51.48a2.06 2.06 0 0 1 0 4.12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageOff(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.937 14.218V5.564a1.5 1.5 0 0 0-1.5-1.5H7.809a.5.5 0 0 1 0-1h10.628a2.5 2.5 0 0 1 2.5 2.5v10.624a.5.5 0 0 1-1 .001v-.556l-4.583-4.584c-.456-.456.251-1.163.707-.707zm-.121 6.304a2.486 2.486 0 0 1-1.379.415H5.563a2.5 2.5 0 0 1-2.5-2.5V5.564c0-.51.153-.984.414-1.38l-.263-.263c-.456-.456.251-1.163.707-.707l.263.263l16.339 16.338l.263.263c.455.456-.252 1.163-.707.707zM8.712 9.419L6.711 7.418a1.5 1.5 0 0 0 2.001 2.001M5.979 6.686l-1.77-1.77a1.51 1.51 0 0 0-.146.648v10.717l1.926-1.926a1.5 1.5 0 0 1 2.122 0l.555.554a.497.497 0 0 0 .706 0l2.415-2.415l-2.343-2.343a2.5 2.5 0 0 1-3.465-3.465M4.063 17.695v.741a1.5 1.5 0 0 0 1.5 1.5h12.874c.232 0 .451-.052.647-.145l-6.59-6.59l-2.414 2.415a1.5 1.5 0 0 1-2.122 0l-.554-.554a.5.5 0 0 0-.708 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageOn(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.435 3.06H5.565a2.5 2.5 0 0 0-2.5 2.5v12.88a2.507 2.507 0 0 0 2.5 2.5h12.87a2.507 2.507 0 0 0 2.5-2.5V5.56a2.5 2.5 0 0 0-2.5-2.5m-14.37 2.5a1.5 1.5 0 0 1 1.5-1.5h12.87a1.5 1.5 0 0 1 1.5 1.5v8.66l-3.88-3.88a1.509 1.509 0 0 0-2.12 0l-4.56 4.57a.513.513 0 0 1-.71 0l-.56-.56a1.522 1.522 0 0 0-2.12 0l-1.92 1.92Zm15.87 12.88a1.5 1.5 0 0 1-1.5 1.5H5.565a1.5 1.5 0 0 1-1.5-1.5v-.75L6.7 15.06a.5.5 0 0 1 .35-.14a.524.524 0 0 1 .36.14l.55.56a1.509 1.509 0 0 0 2.12 0l4.57-4.57a.5.5 0 0 1 .71 0l4.58 4.58Z"/><path fill="currentColor" d="M8.062 10.565a2.5 2.5 0 1 1 2.5-2.5a2.5 2.5 0 0 1-2.5 2.5m0-4a1.5 1.5 0 1 0 1.5 1.5a1.5 1.5 0 0 0-1.5-1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Import(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.552 20.968a2.577 2.577 0 0 1-2.5-2.73c-.012-2.153 0-4.306 0-6.459a.5.5 0 0 1 1 0c0 2.2-.032 4.4 0 6.6c.016 1.107.848 1.589 1.838 1.589h12.463A1.546 1.546 0 0 0 19.825 19a3.023 3.023 0 0 0 .1-1.061v-6.16a.5.5 0 0 1 1 0c0 2.224.085 4.465 0 6.687a2.567 2.567 0 0 1-2.67 2.5Z"/><path fill="currentColor" d="M11.63 15.818a.459.459 0 0 0 .312.138c.014 0 .027.005.042.006s.027 0 .041-.006a.457.457 0 0 0 .312-.138l3.669-3.669a.5.5 0 0 0-.707-.707l-2.815 2.815V3.515a.5.5 0 0 0-1 0v10.742l-2.816-2.815a.5.5 0 0 0-.707.707Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxIn(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.5 9.361v-4.83a.5.5 0 0 1 1 0v4.82l1.27-1.27a.524.524 0 0 1 .71 0a.513.513 0 0 1 0 .71l-2.13 2.12a.492.492 0 0 1-.7 0l-2.12-2.12a.5.5 0 0 1 0-.71a.511.511 0 0 1 .7 0Zm8.988 10.588H3.512a1.451 1.451 0 0 1-1.45-1.449v-5.639a1.451 1.451 0 0 1 1.45-1.449h4.1a1.444 1.444 0 0 1 1.3.8l1.373 2.726a.449.449 0 0 0 .4.247h2.629a.448.448 0 0 0 .4-.248l1.373-2.724a1.442 1.442 0 0 1 1.3-.8h4.1a1.451 1.451 0 0 1 1.45 1.449V18.5a1.451 1.451 0 0 1-1.449 1.449M3.512 12.412a.45.45 0 0 0-.45.449V18.5a.45.45 0 0 0 .45.449h16.976a.45.45 0 0 0 .45-.449v-5.639a.45.45 0 0 0-.45-.449h-4.1a.449.449 0 0 0-.4.247l-1.378 2.725a1.445 1.445 0 0 1-1.295.8h-2.629a1.442 1.442 0 0 1-1.295-.8l-1.373-2.725a.449.449 0 0 0-.4-.247Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxOut(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.488 19.944H3.512a1.452 1.452 0 0 1-1.45-1.45v-5.638a1.452 1.452 0 0 1 1.45-1.45h4.1a1.442 1.442 0 0 1 1.3.8l1.373 2.725a.449.449 0 0 0 .4.247h2.629a.448.448 0 0 0 .4-.248l1.376-2.73a1.442 1.442 0 0 1 1.3-.8h4.1a1.452 1.452 0 0 1 1.45 1.45v5.638a1.452 1.452 0 0 1-1.452 1.456M3.512 12.406a.451.451 0 0 0-.45.45v5.638a.45.45 0 0 0 .45.45h16.976a.45.45 0 0 0 .45-.45v-5.638a.451.451 0 0 0-.45-.45h-4.1a.449.449 0 0 0-.4.247l-1.378 2.725a1.445 1.445 0 0 1-1.295.8h-2.629a1.444 1.444 0 0 1-1.295-.8l-1.373-2.725a.449.449 0 0 0-.4-.247ZM12.5 5.753v4.83a.5.5 0 0 1-1 0v-4.82l-1.27 1.27a.524.524 0 0 1-.71 0a.513.513 0 0 1 0-.71l2.13-2.12a.492.492 0 0 1 .7 0l2.12 2.12a.5.5 0 0 1 0 .71a.511.511 0 0 1-.7 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Indent(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.437 4.064H3.563a.5.5 0 0 1 0-1h16.874a.5.5 0 0 1 0 1m0 5.624h-9a.5.5 0 0 1 0-1h9a.5.5 0 0 1 0 1m0 5.624h-9a.5.5 0 0 1 0-1h9a.5.5 0 0 1 0 1m0 5.624H3.563a.5.5 0 1 1 0-1h16.874a.5.5 0 1 1 0 1M7.91 11.65a.492.492 0 0 1 0 .7l-2 2a.495.495 0 0 1-.7-.7l1.15-1.15H3.54a.5.5 0 0 1 0-1h2.81c-.38-.38-.76-.76-1.14-1.15a.495.495 0 0 1 .7-.7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Instagram(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.437 20.937H5.563a2.5 2.5 0 0 1-2.5-2.5V5.563a2.5 2.5 0 0 1 2.5-2.5h12.874a2.5 2.5 0 0 1 2.5 2.5v12.874a2.5 2.5 0 0 1-2.5 2.5M5.563 4.063a1.5 1.5 0 0 0-1.5 1.5v12.874a1.5 1.5 0 0 0 1.5 1.5h12.874a1.5 1.5 0 0 0 1.5-1.5V5.563a1.5 1.5 0 0 0-1.5-1.5Z"/><path fill="currentColor" d="M12 16.594A4.595 4.595 0 1 1 16.6 12a4.6 4.6 0 0 1-4.6 4.594M12 8.4a3.595 3.595 0 1 0 3.6 3.6A3.6 3.6 0 0 0 12 8.4"/><circle cx="17.2" cy="6.83" r="1.075" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Keyboard(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.437 18.5H4.562a2.5 2.5 0 0 1-2.5-2.5V8a2.5 2.5 0 0 1 2.5-2.5h14.875a2.5 2.5 0 0 1 2.5 2.5v8a2.5 2.5 0 0 1-2.5 2.5M4.562 6.5a1.5 1.5 0 0 0-1.5 1.5v8a1.5 1.5 0 0 0 1.5 1.5h14.875a1.5 1.5 0 0 0 1.5-1.5V8a1.5 1.5 0 0 0-1.5-1.5Z"/><path fill="currentColor" d="M5.548 16.5h12.9a.5.5 0 0 0 0-1h-12.9a.5.5 0 0 0 0 1"/><circle cx="5.82" cy="9.248" r=".75" fill="currentColor"/><circle cx="9.94" cy="9.248" r=".75" fill="currentColor"/><circle cx="14.06" cy="9.248" r=".75" fill="currentColor"/><circle cx="18.18" cy="9.248" r=".75" fill="currentColor"/><circle cx="5.82" cy="12.998" r=".75" fill="currentColor"/><circle cx="9.94" cy="12.998" r=".75" fill="currentColor"/><circle cx="14.06" cy="12.998" r=".75" fill="currentColor"/><circle cx="18.18" cy="12.998" r=".75" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Laptop(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.485 16.155a.992.992 0 0 0-.77-.36h-.33v-9.23a2.5 2.5 0 0 0-2.5-2.5H6.115a2.5 2.5 0 0 0-2.5 2.5V15.8h-.34a1 1 0 0 0-.98 1.17l.3 1.73a1.5 1.5 0 0 0 1.48 1.24h15.85a1.5 1.5 0 0 0 1.48-1.24l.3-1.73a.986.986 0 0 0-.22-.815m-16.87-9.59a1.5 1.5 0 0 1 1.5-1.5h11.77a1.5 1.5 0 0 1 1.5 1.5V15.8H4.615Zm15.8 11.96a.494.494 0 0 1-.49.41H4.075a.494.494 0 0 1-.49-.41l-.31-1.73h17.44Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lemon(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6 20.924a2.172 2.172 0 0 1-1.545-.642l-.734-.733a2.207 2.207 0 0 1-.16-2.947a1.18 1.18 0 0 0 .272-1.117a9.105 9.105 0 0 1 2.372-9.277a9.1 9.1 0 0 1 9.277-2.371a1.149 1.149 0 0 0 1.062-.229l.055-.044a2.205 2.205 0 0 1 2.946.161l.734.733a2.207 2.207 0 0 1 .16 2.947a1.179 1.179 0 0 0-.272 1.116A9.11 9.11 0 0 1 17.8 17.8a9.109 9.109 0 0 1-9.282 2.37a1.14 1.14 0 0 0-1.062.229A2.324 2.324 0 0 1 6 20.924M12.812 4.4a8.427 8.427 0 0 0-5.9 2.519a8.1 8.1 0 0 0-2.133 8.246a2.149 2.149 0 0 1-.395 2.014a1.227 1.227 0 0 0 .044 1.667l.734.733a1.209 1.209 0 0 0 1.613.088a2.175 2.175 0 0 1 2.067-.438a8.1 8.1 0 0 0 8.246-2.133a8.1 8.1 0 0 0 2.133-8.246a2.144 2.144 0 0 1 .395-2.013a1.229 1.229 0 0 0-.044-1.668l-.734-.733a1.206 1.206 0 0 0-1.612-.089l-.052.042a2.148 2.148 0 0 1-2.016.4a7.213 7.213 0 0 0-2.346-.389"/><path fill="currentColor" d="M6.457 12.286a.523.523 0 0 1-.178-.032a.5.5 0 0 1-.29-.646a9.841 9.841 0 0 1 5.338-5.5a.5.5 0 1 1 .386.921a8.845 8.845 0 0 0-4.789 4.939a.5.5 0 0 1-.467.318"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Light(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 18.09A6.09 6.09 0 1 1 18.09 12A6.1 6.1 0 0 1 12 18.09m0-11.18A5.09 5.09 0 1 0 17.09 12A5.1 5.1 0 0 0 12 6.91m-.5-4.342v1.6a.5.5 0 1 0 1 0v-1.6a.5.5 0 1 0-1 0m1 18.864v-1.6a.5.5 0 0 0-1 0v1.6a.5.5 0 1 0 1 0m8.932-9.932h-1.6a.5.5 0 0 0 0 1h1.6a.5.5 0 1 0 0-1m-18.864 1h1.6a.5.5 0 1 0 0-1h-1.6a.5.5 0 1 0 0 1m15.748-7.523l-.992.992l-.141.141a.514.514 0 0 0-.146.353a.508.508 0 0 0 .146.354a.5.5 0 0 0 .354.146a.515.515 0 0 0 .353-.146l.992-.992l.141-.141a.515.515 0 0 0 .147-.354a.508.508 0 0 0-.147-.353a.5.5 0 0 0-.353-.147a.522.522 0 0 0-.354.147M5.684 19.023l.992-.992l.141-.141a.514.514 0 0 0 .146-.353a.508.508 0 0 0-.146-.354a.5.5 0 0 0-.354-.146a.515.515 0 0 0-.353.146l-.992.992l-.141.141a.515.515 0 0 0-.147.354a.508.508 0 0 0 .147.353a.5.5 0 0 0 .353.147a.522.522 0 0 0 .354-.147m13.339-.707l-.992-.992l-.141-.141a.514.514 0 0 0-.353-.146a.508.508 0 0 0-.354.146a.5.5 0 0 0-.146.354a.515.515 0 0 0 .146.353l.992.992l.141.141a.515.515 0 0 0 .354.147a.508.508 0 0 0 .353-.147a.5.5 0 0 0 .147-.353a.522.522 0 0 0-.147-.354M4.977 5.684l.992.992l.141.141a.514.514 0 0 0 .353.146a.508.508 0 0 0 .354-.146a.5.5 0 0 0 .146-.354a.515.515 0 0 0-.146-.353l-.992-.992l-.141-.141a.515.515 0 0 0-.354-.147a.508.508 0 0 0-.353.147a.5.5 0 0 0-.147.353a.522.522 0 0 0 .147.354"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LineHeight(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.439 4.062h-9a.5.5 0 1 1 0-1h9a.5.5 0 0 1 0 1m0 5.624h-9a.5.5 0 0 1 0-1h9a.5.5 0 0 1 0 1m0 5.624h-9a.5.5 0 0 1 0-1h9a.5.5 0 0 1 0 1m0 5.624h-9a.5.5 0 0 1 0-1h9a.5.5 0 0 1 0 1M3.208 18.8a.5.5 0 0 1 .71-.71l1.14 1.14V4.775l-1.14 1.14a.513.513 0 0 1-.71 0a.5.5 0 0 1 0-.71l2-2a.494.494 0 0 1 .34-.14a.549.549 0 0 1 .37.14l2 2a.524.524 0 0 1 0 .71a.5.5 0 0 1-.71 0l-1.15-1.15v14.47l1.15-1.15a.5.5 0 1 1 .71.71l-2 2a.513.513 0 0 1-.71 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Link(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.9 8a4.055 4.055 0 0 1 1.352.135a2.511 2.511 0 0 1-.7 4.863a.5.5 0 0 0 0 1a3.508 3.508 0 0 0 2.944-5.2A3.557 3.557 0 0 0 11.434 7H5.59a3.5 3.5 0 0 0-.19 7c.724.041 1.458 0 2.183 0a.5.5 0 0 0 0-1c-1.323 0-2.915.262-3.891-.843A2.522 2.522 0 0 1 5.59 8Z"/><path fill="currentColor" d="M18.41 17a3.5 3.5 0 0 0 .192-6.994c-.724-.041-1.458 0-2.183 0a.5.5 0 0 0 0 1c1.323 0 2.915-.262 3.891.843A2.522 2.522 0 0 1 18.41 16H13.1a4.055 4.055 0 0 1-1.352-.135a2.511 2.511 0 0 1 .7-4.863a.5.5 0 0 0 0-1a3.508 3.508 0 0 0-2.944 5.2A3.557 3.557 0 0 0 12.566 17Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Linkedin(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.44 3.06H5.56a2.507 2.507 0 0 0-2.5 2.5v12.88a2.507 2.507 0 0 0 2.5 2.5h12.88a2.5 2.5 0 0 0 2.5-2.5V5.56a2.5 2.5 0 0 0-2.5-2.5m1.5 15.38a1.511 1.511 0 0 1-1.5 1.5H5.56a1.511 1.511 0 0 1-1.5-1.5V5.56a1.511 1.511 0 0 1 1.5-1.5h12.88a1.511 1.511 0 0 1 1.5 1.5Z"/><path fill="currentColor" d="M6.376 10.748a1 1 0 1 1 2 0v6.5a1 1 0 0 1-2 0Z"/><circle cx="7.376" cy="6.744" r="1" fill="currentColor"/><path fill="currentColor" d="M17.62 13.37v3.88a1 1 0 1 1-2 0v-3.88a1.615 1.615 0 1 0-3.23 0v3.88a1 1 0 0 1-2 0v-6.5a1.016 1.016 0 0 1 1-1a.94.94 0 0 1 .84.47a3.609 3.609 0 0 1 5.39 3.15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationArrowOne(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.472 20.937a1.438 1.438 0 0 1-1.3-.812L10.3 14.343a1.418 1.418 0 0 0-.642-.641l-5.784-2.871a1.462 1.462 0 0 1 .186-2.695l14.952-5a1.46 1.46 0 0 1 1.849 1.847l-5 14.952a1.439 1.439 0 0 1-1.284.994a.525.525 0 0 1-.105.008m5.007-16.874a.488.488 0 0 0-.149.024l-14.952 5a.46.46 0 0 0-.058.849l5.78 2.869a2.444 2.444 0 0 1 1.1 1.095l2.87 5.782a.443.443 0 0 0 .445.255a.45.45 0 0 0 .4-.312l5-14.953a.462.462 0 0 0-.433-.607Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationOff(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.2 14.462a.5.5 0 0 1-.417-.775a6.791 6.791 0 0 0 1.048-4.627a6.909 6.909 0 0 0-6.022-5.946a6.834 6.834 0 0 0-4.557 1.061a.5.5 0 1 1-.545-.838a7.882 7.882 0 0 1 10.909 10.9a.5.5 0 0 1-.416.225M4 3.3a.5.5 0 0 0-.7.7l1.92 1.92a7.784 7.784 0 0 0-1.11 4.03a7.879 7.879 0 0 0 1.44 4.55l5.06 6.74a1.724 1.724 0 0 0 1.39.69a1.705 1.705 0 0 0 1.38-.69l3.06-4.09c.5.49.99.99 1.48 1.48c.7.69 1.39 1.38 2.08 2.07c.45.46 1.16-.25.7-.71Zm8.58 17.34a.734.734 0 0 1-.58.29a.754.754 0 0 1-.59-.29l-5.05-6.73a6.9 6.9 0 0 1-.41-7.26q1.5 1.515 3.01 3.01q3.39 3.39 6.77 6.78Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationOn(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21.933a1.715 1.715 0 0 1-1.384-.691L5.555 14.5a7.894 7.894 0 1 1 12.885-.009l-5.055 6.749a1.717 1.717 0 0 1-1.385.693m-.008-18.867a6.81 6.81 0 0 0-4.578 1.749a6.891 6.891 0 0 0-1.05 9.1l5.051 6.727a.725.725 0 0 0 .584.292a.732.732 0 0 0 .586-.292l5.044-6.734A6.874 6.874 0 0 0 12.81 3.113a7.277 7.277 0 0 0-.818-.047"/><path fill="currentColor" d="M12 12.5a2.5 2.5 0 1 1 2.5-2.5a2.5 2.5 0 0 1-2.5 2.5m0-4a1.5 1.5 0 1 0 1.5 1.5A1.5 1.5 0 0 0 12 8.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lock(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.44 9.33h-1.1V6.4a4.34 4.34 0 0 0-8.68 0v2.93h-1.1a2.5 2.5 0 0 0-2.5 2.5v7.61a2.507 2.507 0 0 0 2.5 2.5h10.88a2.507 2.507 0 0 0 2.5-2.5v-7.61a2.5 2.5 0 0 0-2.5-2.5M8.66 6.4a3.34 3.34 0 0 1 6.68 0v2.93H8.66Zm10.28 13.04a1.511 1.511 0 0 1-1.5 1.5H6.56a1.511 1.511 0 0 1-1.5-1.5v-7.61a1.5 1.5 0 0 1 1.5-1.5h10.88a1.5 1.5 0 0 1 1.5 1.5Z"/><path fill="currentColor" d="M13 14.95a.984.984 0 0 1-.5.86v1.5a.5.5 0 0 1-1 0v-1.5a.984.984 0 0 1-.5-.86a1 1 0 0 1 2 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Login(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.944 18.432a2.577 2.577 0 0 1-2.729 2.5c-2.153.012-4.307 0-6.46 0a.5.5 0 0 1 0-1c2.2 0 4.4.032 6.6 0c1.107-.016 1.589-.848 1.589-1.838V5.63a1.545 1.545 0 0 0-.969-1.471a3.027 3.027 0 0 0-1.061-.095h-6.159a.5.5 0 0 1 0-1c2.225 0 4.465-.085 6.688 0a2.566 2.566 0 0 1 2.5 2.67Z"/><path fill="currentColor" d="M15.794 12.354a.459.459 0 0 0 .138-.312a.3.3 0 0 0 .006-.042a.29.29 0 0 0-.006-.041a.455.455 0 0 0-.138-.313l-3.669-3.668a.5.5 0 0 0-.707.707l2.816 2.815H3.492a.5.5 0 0 0 0 1h10.742l-2.816 2.815a.5.5 0 0 0 .707.707Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Logout(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.968 18.448a2.577 2.577 0 0 1-2.73 2.5c-2.153.012-4.306 0-6.459 0a.5.5 0 0 1 0-1c2.2 0 4.4.032 6.6 0c1.107-.016 1.589-.848 1.589-1.838V5.647A1.546 1.546 0 0 0 19 4.175a3.023 3.023 0 0 0-1.061-.095h-6.16a.5.5 0 0 1 0-1c2.224 0 4.465-.085 6.687 0a2.567 2.567 0 0 1 2.5 2.67Z"/><path fill="currentColor" d="M3.176 11.663a.455.455 0 0 0-.138.311c0 .015 0 .028-.006.043s0 .027.006.041a.457.457 0 0 0 .138.312l3.669 3.669a.5.5 0 0 0 .707-.707l-2.815-2.816h10.742a.5.5 0 0 0 0-1H4.737L7.552 8.7a.5.5 0 0 0-.707-.707Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lollipop(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 6.565h-.19a6 6 0 0 0-11.62 0H6a1.5 1.5 0 1 0 0 3h.19a5.992 5.992 0 0 0 5.31 4.48v7.39a.5.5 0 0 0 .5.5a.5.5 0 0 0 .5-.5v-7.39a6.013 6.013 0 0 0 5.31-4.48H18a1.5 1.5 0 1 0 0-3m-6-3.5a4.991 4.991 0 0 1 4.77 3.5H7.23a4.991 4.991 0 0 1 4.77-3.5m0 10a4.991 4.991 0 0 1-4.77-3.5h9.54a4.991 4.991 0 0 1-4.77 3.5m6-4.5H6a.5.5 0 0 1-.5-.5a.5.5 0 0 1 .5-.5h12a.5.5 0 0 1 .5.5a.508.508 0 0 1-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mail(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.435 4.065H4.565a2.5 2.5 0 0 0-2.5 2.5v10.87a2.5 2.5 0 0 0 2.5 2.5h14.87a2.5 2.5 0 0 0 2.5-2.5V6.565a2.5 2.5 0 0 0-2.5-2.5m-14.87 1h14.87a1.489 1.489 0 0 1 1.49 1.39c-2.47 1.32-4.95 2.63-7.43 3.95a6.172 6.172 0 0 1-1.06.53a2.083 2.083 0 0 1-1.67-.39c-1.42-.75-2.84-1.51-4.25-2.26c-1.14-.6-2.3-1.21-3.44-1.82a1.491 1.491 0 0 1 1.49-1.4m16.37 12.37a1.5 1.5 0 0 1-1.5 1.5H4.565a1.5 1.5 0 0 1-1.5-1.5V7.6c2.36 1.24 4.71 2.5 7.07 3.75a5.622 5.622 0 0 0 1.35.6a2.872 2.872 0 0 0 2-.41c1.45-.76 2.89-1.53 4.34-2.29c1.04-.56 2.07-1.1 3.11-1.65Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Map(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.21 4.576a1.906 1.906 0 0 0-1.63-.35l-3.53.89a1.086 1.086 0 0 1-.44 0l-4.98-1.24a2.041 2.041 0 0 0-.92 0L4.5 4.936a1.893 1.893 0 0 0-1.44 1.84v11.15a1.871 1.871 0 0 0 .73 1.5a1.906 1.906 0 0 0 1.63.35l3.53-.89a1.086 1.086 0 0 1 .44 0l4.98 1.24a2.315 2.315 0 0 0 .46.05a2.4 2.4 0 0 0 .46-.05l4.21-1.06a1.893 1.893 0 0 0 1.44-1.84V6.076a1.871 1.871 0 0 0-.73-1.5M8.67 17.926l-3.49.87a.89.89 0 0 1-1.12-.87V6.776a.9.9 0 0 1 .68-.87l3.93-.99Zm5.66 1.16l-4.66-1.16V4.916l4.66 1.16Zm5.61-1.86a.9.9 0 0 1-.68.87l-3.93.99V6.076l3.49-.87a.908.908 0 0 1 .78.16a.886.886 0 0 1 .34.71Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapPin(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2.06a5.5 5.5 0 0 0-.5 10.97v8.41a.5.5 0 0 0 .5.5a.5.5 0 0 0 .5-.5v-8.41A5.5 5.5 0 0 0 12 2.06m0 10a4.5 4.5 0 1 1 4.5-4.5a4.5 4.5 0 0 1-4.5 4.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MaximizeOne(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.065 16.16a.5.5 0 0 1 1 0v3.07l.01-.01l6.07-6.07a.5.5 0 0 1 .71.71c-.29.29-.58.57-.87.86c-1.74 1.74-3.47 3.48-5.21 5.22h3.07a.5.5 0 0 1 0 1h-4.28a.429.429 0 0 1-.34-.14c-.01-.01-.02-.01-.02-.02a.384.384 0 0 1-.13-.26c-.009-.078-.01-4.36-.01-4.36m17.87-12.6v4.28a.5.5 0 0 1-1 0V4.77l-.01.01q-3.045 3.03-6.07 6.07a.5.5 0 0 1-.71-.71c.29-.29.58-.57.86-.86c1.75-1.74 3.48-3.48 5.22-5.22h-3.07a.5.5 0 0 1 0-1h4.28a.429.429 0 0 1 .34.14c.01.01.02.01.02.02a.429.429 0 0 1 .14.34"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MaximizeTwo(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.513 3.066H4.93a2.058 2.058 0 0 0-1.15.22a1.6 1.6 0 0 0-.717 1.437v5.793a.5.5 0 0 0 1 0V5.107a2.521 2.521 0 0 1 .022-.689c.115-.373.469-.352.777-.352h5.651a.5.5 0 0 0 0-1m-7.45 10.422v5.583a2.057 2.057 0 0 0 .221 1.15a1.6 1.6 0 0 0 1.436.717h5.793a.5.5 0 0 0 0-1H5.1a2.483 2.483 0 0 1-.689-.022c-.372-.115-.352-.469-.352-.777v-5.651a.5.5 0 0 0-1 0Zm10.424 7.446h5.583a2.058 2.058 0 0 0 1.15-.22a1.6 1.6 0 0 0 .717-1.437v-5.793a.5.5 0 0 0-1 0v5.409a2.521 2.521 0 0 1-.022.689c-.115.373-.469.352-.777.352h-5.651a.5.5 0 0 0 0 1m7.45-10.422V4.929a2.057 2.057 0 0 0-.221-1.15a1.6 1.6 0 0 0-1.436-.717h-5.793a.5.5 0 0 0 0 1H18.9a2.483 2.483 0 0 1 .689.022c.372.115.352.469.352.777v5.651a.5.5 0 0 0 1 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Medal(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.692 3.755a1.519 1.519 0 0 0-1.27-.69h-4.41a1.487 1.487 0 0 0-1.36.87L12 7.485l-1.66-3.55a1.487 1.487 0 0 0-1.36-.87H4.572a1.5 1.5 0 0 0-1.35 2.14l3.73 8.02a5.638 5.638 0 0 0-.46 2.21a5.5 5.5 0 0 0 11 0a5.419 5.419 0 0 0-.46-2.2l3.75-8.03a1.525 1.525 0 0 0-.09-1.45m-16.57 1.03a.527.527 0 0 1 .03-.49a.494.494 0 0 1 .42-.23h4.41a.507.507 0 0 1 .46.29l2.61 5.58h-.06a5.505 5.505 0 0 0-4.43 2.25Zm7.87 15.15a4.5 4.5 0 1 1 4.5-4.5a4.5 4.5 0 0 1-4.5 4.5m7.89-15.15l-3.46 7.4a5.454 5.454 0 0 0-3.21-2.11l-.66-1.42l2-4.3a.507.507 0 0 1 .46-.29h4.41a.482.482 0 0 1 .42.23a.505.505 0 0 1 .04.49"/><path fill="currentColor" d="m12.077 16.88l1.024.538a.174.174 0 0 0 .253-.184l-.2-1.14a.174.174 0 0 1 .051-.154l.828-.807a.175.175 0 0 0-.1-.3l-1.133-.164a.177.177 0 0 1-.132-.1l-.512-1.037a.174.174 0 0 0-.313 0l-.512 1.037a.174.174 0 0 1-.131.1l-1.145.166a.175.175 0 0 0-.1.3l.828.807a.174.174 0 0 1 .05.154l-.2 1.14a.174.174 0 0 0 .253.184l1.024-.538a.172.172 0 0 1 .167-.002"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MedicalCase(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.435 6.465h-1.43v-.9a2.5 2.5 0 0 0-2.5-2.5h-5a2.5 2.5 0 0 0-2.5 2.5v.9h-1.44a2.5 2.5 0 0 0-2.5 2.5v9.47a2.5 2.5 0 0 0 2.5 2.5h12.87a2.5 2.5 0 0 0 2.5-2.5v-9.47a2.5 2.5 0 0 0-2.5-2.5m-10.43-.9a1.5 1.5 0 0 1 1.5-1.5h5a1.5 1.5 0 0 1 1.5 1.5v.9h-8Zm11.93 12.87a1.5 1.5 0 0 1-1.5 1.5H5.565a1.5 1.5 0 0 1-1.5-1.5v-9.47a1.5 1.5 0 0 1 1.5-1.5h12.87a1.5 1.5 0 0 1 1.5 1.5Z"/><path fill="currentColor" d="M14.505 13.675a.5.5 0 0 1-.5.5h-1.5v1.5a.5.5 0 0 1-.5.5a.5.5 0 0 1-.5-.5v-1.5h-1.5a.5.5 0 0 1-.5-.5a.5.5 0 0 1 .5-.5h1.5v-1.5a.5.5 0 0 1 .5-.5a.508.508 0 0 1 .5.5v1.5h1.5a.508.508 0 0 1 .5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MedicalClipboard(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.44 3.5h-1.69v-.53a.949.949 0 0 0-1-.91h-5.5a.949.949 0 0 0-1 .91v.53H6.56A2.5 2.5 0 0 0 4.06 6v13.44a2.5 2.5 0 0 0 2.5 2.5h10.88a2.5 2.5 0 0 0 2.5-2.5V6a2.5 2.5 0 0 0-2.5-2.5m-8.19-.44l5.5.01v1.12c0 .61-.69 1.12-1.5 1.12h-2.5c-.82 0-1.5-.51-1.5-1.12Zm9.69 16.38a1.511 1.511 0 0 1-1.5 1.5H6.56a1.5 1.5 0 0 1-1.5-1.5V6a1.5 1.5 0 0 1 1.5-1.5h1.72a2.4 2.4 0 0 0 2.47 1.81h2.5a2.4 2.4 0 0 0 2.47-1.81h1.72a1.5 1.5 0 0 1 1.5 1.5Z"/><path fill="currentColor" d="M14.5 12.71a.5.5 0 0 1-.5.5h-1.5v1.5a.5.5 0 0 1-.5.5a.5.5 0 0 1-.5-.5v-1.5H10a.5.5 0 0 1-.5-.5a.5.5 0 0 1 .5-.5h1.5v-1.5a.5.5 0 0 1 .5-.5a.508.508 0 0 1 .5.5v1.5H14a.508.508 0 0 1 .5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MedicalCross(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.943 20.93h-1.886a2.388 2.388 0 0 1-2.386-2.386V15.3l-3.215.029a2.39 2.39 0 0 1-2.387-2.386v-1.886A2.39 2.39 0 0 1 5.456 8.67H8.7l-.029-3.214a2.388 2.388 0 0 1 2.386-2.386h1.886a2.388 2.388 0 0 1 2.386 2.386V8.7l3.215-.03a2.39 2.39 0 0 1 2.387 2.387v1.886a2.39 2.39 0 0 1-2.387 2.386H15.3l.028 3.215a2.388 2.388 0 0 1-2.385 2.386M5.456 9.67a1.388 1.388 0 0 0-1.387 1.387v1.886a1.388 1.388 0 0 0 1.387 1.386H8.7a.972.972 0 0 1 .972.971v3.244a1.388 1.388 0 0 0 1.386 1.386h1.886a1.388 1.388 0 0 0 1.386-1.386V15.3a.972.972 0 0 1 .972-.971h3.243a1.388 1.388 0 0 0 1.387-1.386v-1.886a1.388 1.388 0 0 0-1.388-1.387H15.3a.972.972 0 0 1-.972-.97V5.456a1.388 1.388 0 0 0-1.385-1.386h-1.886a1.388 1.388 0 0 0-1.386 1.386V8.7a.972.972 0 0 1-.972.97Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MedicalMask(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.435 12.53H19.5V9.17h1.94a.5.5 0 0 0 .5-.5a.5.5 0 0 0-.5-.5H19.5v-.94a2.5 2.5 0 0 0-2.5-2.5H7a2.507 2.507 0 0 0-2.5 2.5v.94H2.565a.5.5 0 0 0-.5.5a.5.5 0 0 0 .5.5H4.5v3.36H2.565a.5.5 0 0 0-.5.5a.5.5 0 0 0 .5.5H4.5v.33a4.283 4.283 0 0 0 2.43 3.84l1.74.82a7.79 7.79 0 0 0 6.67 0l1.73-.82h.01a4.274 4.274 0 0 0 2.42-3.84v-.33h1.94a.5.5 0 0 0 .5-.5a.5.5 0 0 0-.505-.5M18.5 13.86a3.238 3.238 0 0 1-1.85 2.93l-1.73.82a6.767 6.767 0 0 1-5.83 0l-1.73-.82a3.248 3.248 0 0 1-1.86-2.93V7.23A1.5 1.5 0 0 1 7 5.73h10a1.5 1.5 0 0 1 1.5 1.5Z"/><path fill="currentColor" d="M14.5 14.534h-5a.5.5 0 0 1 0-1h5a.5.5 0 0 1 0 1m1.5-4.366H8a.5.5 0 1 1 0-1h8a.5.5 0 0 1 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MemoPad(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.44 2.065H6.56a2.507 2.507 0 0 0-2.5 2.5v14.87a2.507 2.507 0 0 0 2.5 2.5h10.88a2.5 2.5 0 0 0 2.5-2.5V4.565a2.5 2.5 0 0 0-2.5-2.5m1.5 17.37a1.5 1.5 0 0 1-1.5 1.5H6.56a1.5 1.5 0 0 1-1.5-1.5V6.505h13.88Z"/><path fill="currentColor" d="M7.549 9.506a.5.5 0 0 1 0-1h8.909a.5.5 0 0 1 0 1Zm0 3a.5.5 0 0 1 0-1h6.5a.5.5 0 0 1 0 1Zm.017 5.868a.5.5 0 1 1 0-1h3.251a.5.5 0 0 1 0 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuBurger(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.563 4.063a.5.5 0 0 1 0-1l16.874-.001a.5.5 0 0 1 0 1zm0 8.438a.5.5 0 0 1 0-1l16.874-.002a.5.5 0 0 1 0 1zm0 8.438a.5.5 0 0 1 0-1l16.874-.002a.5.5 0 0 1 0 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuFries(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.437 19.937a.5.5 0 0 1 0 1l-16.874.002a.5.5 0 0 1 0-1zm0-8.437a.5.5 0 0 1 0 1l-10 .001a.5.5 0 0 1 0-1zm0-8.438a.5.5 0 0 1 0 1l-16.874.001a.5.5 0 0 1 0-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MenuKebab(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.5 12a2.5 2.5 0 0 1-5 0a2.5 2.5 0 0 1 5 0m-1 0a1.5 1.5 0 1 0-3.001.001A1.5 1.5 0 0 0 13.5 12m1-7.437a2.5 2.5 0 0 1-5 0a2.5 2.5 0 0 1 5 0m-1 0a1.5 1.5 0 1 0-3.001.001a1.5 1.5 0 0 0 3.001-.001m1 14.874a2.5 2.5 0 0 1-5 0a2.5 2.5 0 0 1 5 0m-1 0a1.5 1.5 0 1 0-3.001.001a1.5 1.5 0 0 0 3.001-.001"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Microchip(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 14.5h-2A1.5 1.5 0 0 1 9.5 13v-2A1.5 1.5 0 0 1 11 9.5h2a1.5 1.5 0 0 1 1.5 1.5v2a1.5 1.5 0 0 1-1.5 1.5m-2-4a.5.5 0 0 0-.5.5v2a.5.5 0 0 0 .5.5h2a.5.5 0 0 0 .5-.5v-2a.5.5 0 0 0-.5-.5Z"/><path fill="currentColor" d="M20.435 14.5h-1.93v-5h1.93a.5.5 0 0 0 0-1h-1.93V8a2.507 2.507 0 0 0-2.5-2.5h-.5V3.565a.508.508 0 0 0-.5-.5a.5.5 0 0 0-.5.5V5.5h-5V3.565a.508.508 0 0 0-.5-.5a.5.5 0 0 0-.5.5V5.5h-.5a2.5 2.5 0 0 0-2.5 2.5v.5h-1.94a.5.5 0 1 0 0 1h1.94v5h-1.94a.5.5 0 1 0 0 1h1.94v.5a2.5 2.5 0 0 0 2.5 2.5h.5v1.94a.5.5 0 0 0 1 0V18.5h5v1.94a.5.5 0 0 0 1 0V18.5h.5a2.507 2.507 0 0 0 2.5-2.5v-.5h1.93a.5.5 0 0 0 0-1m-2.93 1.5a1.5 1.5 0 0 1-1.5 1.5h-8a1.5 1.5 0 0 1-1.5-1.5V8a1.5 1.5 0 0 1 1.5-1.5h8a1.511 1.511 0 0 1 1.5 1.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicrophoneOff(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.5 7.046v4.72a.5.5 0 0 1-1 0v-1.82H14a.5.5 0 0 1 0-1h1.5v-2h-1.93a.5.5 0 0 1 0-1h1.87a3.23 3.23 0 0 0-.2-.72a3.533 3.533 0 0 0-6.16-.59c-.36.53-1.23.03-.87-.5a4.509 4.509 0 0 1 7.71.21a5.255 5.255 0 0 1 .58 2.7m3.64 12.39q-2.625-2.64-5.27-5.28q-4.185-4.17-8.36-8.356c-.65-.64-1.3-1.29-1.94-1.94a.5.5 0 0 0-.71.71Q5.69 6.381 7.5 8.206v3.92a4.591 4.591 0 0 0 3.59 4.61a4.629 4.629 0 0 0 3.9-1.04c.24.24.48.47.71.71a5.252 5.252 0 0 1-6.62.67a5.2 5.2 0 0 1-2.05-2.76a7.608 7.608 0 0 1-.24-2.33v-2.2a.5.5 0 0 0-1 0a15.463 15.463 0 0 0 .34 4.99a6.276 6.276 0 0 0 5.37 4.17v1.99H8a.5.5 0 0 0 0 1h8a.5.5 0 0 0 0-1h-3.5v-2a6.118 6.118 0 0 0 3.91-1.82l1.08 1.08c.65.65 1.3 1.3 1.95 1.94a.5.5 0 0 0 .7-.7m-11.2-5.42a3.991 3.991 0 0 1-.44-2.03v-2.78l5.77 5.77a3.521 3.521 0 0 1-5.33-.96"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicrophoneOn(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.989 2.065a4.507 4.507 0 0 0-4.5 4.5v5.76a4.5 4.5 0 0 0 9 0v-5.76a4.507 4.507 0 0 0-4.5-4.5m0 13.76a3.5 3.5 0 0 1-3.5-3.5v-5.76a3.5 3.5 0 0 1 6.94-.62h-1.87a.5.5 0 0 0-.5.5a.5.5 0 0 0 .5.5h1.93v2h-1.93a.5.5 0 0 0-.5.5a.508.508 0 0 0 .5.5h1.93v2h-1.94a.508.508 0 0 0-.5.5a.515.515 0 0 0 .5.5h1.88a3.492 3.492 0 0 1-3.44 2.88"/><path fill="currentColor" d="M12.489 18.925v2.01h3.5a.5.5 0 0 1 0 1h-8a.5.5 0 0 1 0-1h3.5v-1.99a6.055 6.055 0 0 1-2.74-.88a6.291 6.291 0 0 1-2.97-5.14c-.03-1.04 0-2.09 0-3.13a.5.5 0 0 1 1 0c0 1.04-.03 2.09 0 3.13A5.212 5.212 0 0 0 17.2 12.7c.01-.96 0-1.93 0-2.9a.5.5 0 0 1 1 0a26.322 26.322 0 0 1-.08 3.97a6.235 6.235 0 0 1-5.631 5.155"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinimizeOne(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 17.78a.5.5 0 0 1-1 0v-3.07l-6.08 6.08a.5.5 0 0 1-.71-.71c.29-.29.58-.57.87-.86C5.82 17.48 7.55 15.74 9.3 14H6.22a.5.5 0 0 1 0-1h4.28a.429.429 0 0 1 .34.14c.01.01.02.01.02.02a.384.384 0 0 1 .13.26ZM14.7 10h3.08a.5.5 0 0 1 0 1H13.5a.429.429 0 0 1-.34-.14c-.01-.01-.02-.01-.02-.02a.384.384 0 0 1-.13-.26L13 6.22a.5.5 0 0 1 1 0v3.07l.01-.01l6.07-6.07a.5.5 0 0 1 .71.71Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinimizeTwo(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.563 11.016h5.583A2.057 2.057 0 0 0 10.3 10.8a1.6 1.6 0 0 0 .717-1.436V3.566a.5.5 0 0 0-1 0v5.408a2.481 2.481 0 0 1-.022.689c-.115.373-.468.353-.777.353H3.563a.5.5 0 0 0 0 1m7.45 9.422v-5.583a2.065 2.065 0 0 0-.22-1.15a1.6 1.6 0 0 0-1.437-.717H3.563a.5.5 0 0 0 0 1h5.409a2.482 2.482 0 0 1 .689.022c.373.115.352.469.352.777v5.651a.5.5 0 0 0 1 0m9.424-7.454h-5.583a2.057 2.057 0 0 0-1.15.221a1.6 1.6 0 0 0-.717 1.436v5.793a.5.5 0 0 0 1 0v-5.408a2.481 2.481 0 0 1 .022-.689c.115-.373.468-.353.777-.353h5.651a.5.5 0 0 0 0-1m-7.45-9.422v5.583a2.065 2.065 0 0 0 .22 1.15a1.6 1.6 0 0 0 1.437.717h5.793a.5.5 0 0 0 0-1h-5.409a2.482 2.482 0 0 1-.689-.022c-.373-.115-.352-.469-.352-.777V3.562a.5.5 0 0 0-1 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MobileFour(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="12.003" cy="18.937" r="1" fill="currentColor"/><path fill="currentColor" d="M16.725 2.065h-9.45a2.386 2.386 0 0 0-2.24 2.5v14.87a2.386 2.386 0 0 0 2.24 2.5h9.45a2.379 2.379 0 0 0 2.24-2.5V4.565a2.379 2.379 0 0 0-2.24-2.5m1.24 17.37a1.384 1.384 0 0 1-1.24 1.5h-9.45a1.391 1.391 0 0 1-1.24-1.5v-2.5h11.93Zm0-3.5H6.035V4.565a1.391 1.391 0 0 1 1.24-1.5h9.45a1.384 1.384 0 0 1 1.24 1.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MobileOne(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 18.933h4a.5.5 0 0 0 0-1h-4a.5.5 0 0 0 0 1"/><path fill="currentColor" d="M16.727 21.937H7.273a2.384 2.384 0 0 1-2.239-2.5V4.563a2.384 2.384 0 0 1 2.239-2.5h9.454a2.384 2.384 0 0 1 2.239 2.5v14.874a2.384 2.384 0 0 1-2.239 2.5M7.273 3.063a1.39 1.39 0 0 0-1.239 1.5v14.874a1.39 1.39 0 0 0 1.239 1.5h9.454a1.39 1.39 0 0 0 1.239-1.5V4.563a1.39 1.39 0 0 0-1.239-1.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MobileThree(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.73 2.065H7.27a2.386 2.386 0 0 0-2.24 2.5v14.87a2.386 2.386 0 0 0 2.24 2.5h9.46a2.386 2.386 0 0 0 2.24-2.5V4.565a2.386 2.386 0 0 0-2.24-2.5m1.24 17.37a1.391 1.391 0 0 1-1.24 1.5H7.27a1.391 1.391 0 0 1-1.24-1.5V4.565a1.391 1.391 0 0 1 1.24-1.5H8.8v.51a1 1 0 0 0 1 1h4.4a1 1 0 0 0 1-1v-.51h1.53a1.391 1.391 0 0 1 1.24 1.5Z"/><path fill="currentColor" d="M10 18.934h4a.5.5 0 0 0 0-1h-4a.5.5 0 0 0 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MobileTwo(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="12" cy="17.937" r="1" fill="currentColor"/><path fill="currentColor" d="M16.727 21.937H7.273a2.384 2.384 0 0 1-2.239-2.5V4.563a2.384 2.384 0 0 1 2.239-2.5h9.454a2.384 2.384 0 0 1 2.239 2.5v14.874a2.384 2.384 0 0 1-2.239 2.5M7.273 3.063a1.39 1.39 0 0 0-1.239 1.5v14.874a1.39 1.39 0 0 0 1.239 1.5h9.454a1.39 1.39 0 0 0 1.239-1.5V4.563a1.39 1.39 0 0 0-1.239-1.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoneyBill(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.44 5.78H4.56a2.507 2.507 0 0 0-2.5 2.5v7.44a2.514 2.514 0 0 0 2.5 2.5h14.88a2.507 2.507 0 0 0 2.5-2.5V8.28a2.5 2.5 0 0 0-2.5-2.5M3.06 8.28a1.5 1.5 0 0 1 1.5-1.5h1.48a3.521 3.521 0 0 1-2.98 2.98Zm1.5 8.94a1.511 1.511 0 0 1-1.5-1.5v-1.48a3.521 3.521 0 0 1 2.98 2.98Zm16.38-1.5a1.5 1.5 0 0 1-1.5 1.5h-1.48a3.521 3.521 0 0 1 2.98-2.98Zm0-2.49a4.528 4.528 0 0 0-3.99 3.99h-9.9a4.528 4.528 0 0 0-3.99-3.99v-2.46a4.528 4.528 0 0 0 3.99-3.99h9.9a4.528 4.528 0 0 0 3.99 3.99Zm0-3.47a3.521 3.521 0 0 1-2.98-2.98h1.48a1.5 1.5 0 0 1 1.5 1.5Z"/><circle cx="12.002" cy="11.998" r="2" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoneyCheckOne(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.437 18.218H4.563a2.5 2.5 0 0 1-2.5-2.5V8.282a2.5 2.5 0 0 1 2.5-2.5h14.874a2.5 2.5 0 0 1 2.5 2.5v7.436a2.5 2.5 0 0 1-2.5 2.5M4.563 6.782a1.5 1.5 0 0 0-1.5 1.5v7.436a1.5 1.5 0 0 0 1.5 1.5h14.874a1.5 1.5 0 0 0 1.5-1.5V8.282a1.5 1.5 0 0 0-1.5-1.5Z"/><path fill="currentColor" d="M12 12.786H5.064a.5.5 0 0 1 0-1H12a.5.5 0 0 1 0 1m2 2.928H5.064a.5.5 0 1 1 0-1H14a.5.5 0 0 1 0 1"/><rect width="4" height="2" x="15.436" y="8.283" fill="currentColor" rx=".5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Monitor(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.435 3.06H5.565a2.5 2.5 0 0 0-2.5 2.5v8.88a2.507 2.507 0 0 0 2.5 2.5h2.91l-.37 3H7a.5.5 0 0 0 0 1h10.01a.5.5 0 0 0 0-1H15.9l-.37-3h2.91a2.507 2.507 0 0 0 2.5-2.5V5.56a2.5 2.5 0 0 0-2.505-2.5M14.9 19.94H9.115l.37-3h5.03Zm5.04-5.5a1.5 1.5 0 0 1-1.5 1.5H5.565a1.5 1.5 0 0 1-1.5-1.5V5.56a1.5 1.5 0 0 1 1.5-1.5h12.87a1.5 1.5 0 0 1 1.5 1.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MountainOne(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.857 19.525l-6.57-14.96a2.5 2.5 0 0 0-4.58-.01l-6.56 14.96a1 1 0 0 0 .07.96a.985.985 0 0 0 .84.46h15.89a1 1 0 0 0 .91-1.41m-10.23-14.56a1.5 1.5 0 0 1 2.75 0l2.43 5.53l-1.45 1.45a.5.5 0 0 1-.71 0l-2.04-2.03a1.5 1.5 0 0 0-1.06-.44h-1.9Zm-6.57 14.96l4.15-9.45h2.34a.491.491 0 0 1 .36.15l2.03 2.03A1.508 1.508 0 0 0 14 13.1a1.491 1.491 0 0 0 1.06-.44l1.18-1.17l3.71 8.45Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MugOne(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.115 5.91v-.94a1.2 1.2 0 0 0-1.2-1.2H4.265a1.2 1.2 0 0 0-1.2 1.2v9.14a11.321 11.321 0 0 0 .8 4.17A3.3 3.3 0 0 0 7 20.23h7.19a3.312 3.312 0 0 0 3.14-1.95a10.989 10.989 0 0 0 .74-3.13l1.04-.52a3.319 3.319 0 0 0 1.83-2.97V9.19a3.326 3.326 0 0 0-2.825-3.28m-1.73 12.01a2.3 2.3 0 0 1-2.2 1.31H7a2.312 2.312 0 0 1-2.2-1.31a10.238 10.238 0 0 1-.73-3.81V4.97a.2.2 0 0 1 .2-.2h12.65a.2.2 0 0 1 .2.2v9.14a10.238 10.238 0 0 1-.735 3.81m3.55-6.26a2.287 2.287 0 0 1-1.28 2.07l-.54.27V6.93a2.316 2.316 0 0 1 1.82 2.26Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicNoteOne(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.05 3.657a2.487 2.487 0 0 0-2.03-.56l-7.88 1.33a2.483 2.483 0 0 0-2.08 2.46v8.82a3 3 0 1 0 1 2.23v-8.55l10.88-1.83v6.22a2.936 2.936 0 0 0-2-.77a3 3 0 1 0 3 3V5.567a2.513 2.513 0 0 0-.89-1.91M6.06 19.937a2 2 0 1 1 2-2a1.993 1.993 0 0 1-2 2m11.88-1.93a2 2 0 1 1 2-2a2 2 0 0 1-2 2m2-11.46L9.06 8.377v-1.49a1.483 1.483 0 0 1 1.25-1.47l7.88-1.33a1.493 1.493 0 0 1 1.75 1.48Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoWaitingSign(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2.07a9.93 9.93 0 1 0 7.03 16.95a.374.374 0 0 0 .06-.07A9.837 9.837 0 0 0 21.935 12A9.944 9.944 0 0 0 12 2.07m0 18.86A8.945 8.945 0 0 1 3.065 12a8.84 8.84 0 0 1 2.28-5.95l12.61 12.61A8.925 8.925 0 0 1 12 20.93m6.67-2.98L6.045 5.34a8.934 8.934 0 0 1 12.62 12.61Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Palette(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.31 20.93a1.62 1.62 0 0 0 1.62-1.62v-3.38a1.62 1.62 0 0 0-1.62-1.62h-4.94l4.66-4.66a1.642 1.642 0 0 0 0-2.3l-2.39-2.39a1.636 1.636 0 0 0-2.3 0L9.69 9.62V4.56a1.5 1.5 0 0 0-1.5-1.5H4.57a1.5 1.5 0 0 0-1.5 1.5v13.88a2.507 2.507 0 0 0 2.5 2.5Zm-9.62-9.89l5.36-5.37a.628.628 0 0 1 .88 0l2.39 2.39a.628.628 0 0 1 0 .88l-8.63 8.63ZM4.07 4.56a.5.5 0 0 1 .5-.5h3.62a.5.5 0 0 1 .5.5v3.92H4.07Zm0 4.92h4.62v4.43H4.07Zm3.22 10.45l-1.72.01a1.5 1.5 0 0 1-1.5-1.5v-3.53h4.62v3.53a1.5 1.5 0 0 1-1.4 1.49m12.64-.62a.623.623 0 0 1-.62.62H9.19a2.381 2.381 0 0 0 .42-.86l3.76-3.76h5.94a.623.623 0 0 1 .62.62Z"/><circle cx="6.382" cy="17.419" r=".844" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paperplane(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.061 11.077l-17.32-6.92a.994.994 0 0 0-1.17.32a1 1 0 0 0-.01 1.22l4.49 6a.525.525 0 0 1-.01.62L2.511 18.3a1.02 1.02 0 0 0 0 1.22a1 1 0 0 0 .8.4a1.021 1.021 0 0 0 .38-.07l17.36-6.9a1.006 1.006 0 0 0 .01-1.87Zm-17.69-5.99l16.06 6.42H8.061a1.329 1.329 0 0 0-.21-.41Zm-.06 13.82l4.53-5.98a1.212 1.212 0 0 0 .22-.42h11.38Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ParkingOne(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21.933A9.933 9.933 0 1 1 21.933 12A9.944 9.944 0 0 1 12 21.933m0-18.866A8.933 8.933 0 1 0 20.933 12A8.943 8.943 0 0 0 12 3.067"/><path fill="currentColor" d="M12.569 8.5h-1.75a.749.749 0 0 0-.75.75v5.74a.5.5 0 0 0 .5.5a.5.5 0 0 0 .5-.5V13.5h1.5a2.5 2.5 0 0 0 0-5m0 4h-1.5v-3h1.5a1.5 1.5 0 0 1 0 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PassportOne(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.5 21.936h-9a2.5 2.5 0 0 1-2.5-2.5V4.564a2.5 2.5 0 0 1 2.5-2.5h9a2.5 2.5 0 0 1 2.5 2.5v14.872a2.5 2.5 0 0 1-2.5 2.5m-9-18.872a1.5 1.5 0 0 0-1.5 1.5v14.872a1.5 1.5 0 0 0 1.5 1.5h9a1.5 1.5 0 0 0 1.5-1.5V4.564a1.5 1.5 0 0 0-1.5-1.5Z"/><path fill="currentColor" d="M12 12.563a3.5 3.5 0 1 1 3.5-3.5a3.5 3.5 0 0 1-3.5 3.5m0-6a2.5 2.5 0 1 0 2.5 2.5a2.5 2.5 0 0 0-2.5-2.5m3 11.875H9a.5.5 0 0 1 0-1h6a.5.5 0 1 1 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseOne(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.25 21.937H6.564a2.5 2.5 0 0 1-2.5-2.5V4.563a2.5 2.5 0 0 1 2.5-2.5H8.25a2.5 2.5 0 0 1 2.5 2.5v14.874a2.5 2.5 0 0 1-2.5 2.5M6.564 3.063a1.5 1.5 0 0 0-1.5 1.5v14.874a1.5 1.5 0 0 0 1.5 1.5H8.25a1.5 1.5 0 0 0 1.5-1.5V4.563a1.5 1.5 0 0 0-1.5-1.5Zm10.872 18.874H15.75a2.5 2.5 0 0 1-2.5-2.5V4.563a2.5 2.5 0 0 1 2.5-2.5h1.686a2.5 2.5 0 0 1 2.5 2.5v14.874a2.5 2.5 0 0 1-2.5 2.5M15.75 3.063a1.5 1.5 0 0 0-1.5 1.5v14.874a1.5 1.5 0 0 0 1.5 1.5h1.686a1.5 1.5 0 0 0 1.5-1.5V4.563a1.5 1.5 0 0 0-1.5-1.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pen(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.235 11.284a2.3 2.3 0 0 0-3.01-.149l-1.781-5.391a2.484 2.484 0 0 0-2.1-1.7l-8.581-.93a1.5 1.5 0 0 0-1.648 1.651l.93 8.579a2.479 2.479 0 0 0 1.7 2.1l5.39 1.77a2.258 2.258 0 0 0-.51 1.43a2.257 2.257 0 0 0 2.25 2.25a2.263 2.263 0 0 0 1.591-.661l5.77-5.769a2.249 2.249 0 0 0 0-3.181Zm-14.18 3.21a1.5 1.5 0 0 1-1.02-1.26l-.9-8.39l4.01 4.01a1.188 1.188 0 0 0 .281 1.221a1.167 1.167 0 1 0 1.649-1.651a1.143 1.143 0 0 0-1.209-.269l-4.02-4.02l8.39.9a1.476 1.476 0 0 1 1.259 1.02l1.931 5.86l-4.51 4.51Zm11.709-2.51a1.25 1.25 0 0 1 2.13.891a1.237 1.237 0 0 1-.369.88l-5.771 5.77a1.277 1.277 0 0 1-1.769 0a1.253 1.253 0 0 1 0-1.76Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Penpot(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.415 7.22a.755.755 0 0 0-.35.64v9.67a1.536 1.536 0 0 0 .88 1.38l5.96 2.82a2.618 2.618 0 0 0 2.19 0l5.96-2.82a1.536 1.536 0 0 0 .88-1.38V7.86a.742.742 0 0 0-.36-.64l-2.23-1.15v-1a.98.98 0 0 0-.15-.52l-1.33-2.16a.749.749 0 0 0-1.28 0L13.3 4.5l-.66-1.07a.755.755 0 0 0-.64-.35a.791.791 0 0 0-.64.36l-.65 1.06L9.4 2.39a.749.749 0 0 0-1.28 0L6.805 4.55a.98.98 0 0 0-.15.52v1Zm12.93-.01l.7.35l-.7.33Zm-2.44-3.43h.64l.5.81h-1.64Zm1.44 1.81v2.77l-.61.29V5.59Zm-2.23 0h.62v3.53l-.62.3Zm-5.67-1.81h.64l.5.81h-1.64Zm1.44 1.81v3.82l-.61-.29V5.59Zm-2.23 0h.62v3.06l-.62-.29Zm-1 1.62v.68l-.7-.33ZM5.365 18a.52.52 0 0 1-.3-.47V8.24l6.43 3.04v9.59Zm6.13-7.83l-.61-.29V6.64h.61Zm-.32-4.53l.51-.82h.62l.51.82Zm1.94 1v3.25l-.62.29V6.64ZM12.5 20.88v-9.59l6.44-3.05v9.29a.512.512 0 0 1-.31.47Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Percent(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21.934A9.934 9.934 0 1 1 21.933 12A9.945 9.945 0 0 1 12 21.934m0-18.868A8.934 8.934 0 1 0 20.933 12A8.944 8.944 0 0 0 12 3.066"/><path fill="currentColor" d="M9 10.258a1.5 1.5 0 1 1 1.061-.439A1.5 1.5 0 0 1 9 10.258m0-2a.5.5 0 1 0 .353.146A.5.5 0 0 0 9 8.259Zm-1.242 8.485a.5.5 0 0 1-.358-.853l8.489-8.49a.5.5 0 0 1 .707.707L8.111 16.6a.5.5 0 0 1-.353.143M15 16.742a1.5 1.5 0 1 1 1.061-.438a1.493 1.493 0 0 1-1.061.438m0-2a.5.5 0 0 0-.354.147a.5.5 0 0 0-.146.352a.5.5 0 1 0 1 0a.5.5 0 0 0-.5-.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Phone(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.436 20.938A11.384 11.384 0 0 1 4.572 3.9a1.668 1.668 0 0 1 1.241-.822a1.716 1.716 0 0 1 1.454.492l3.139 3.14a1.715 1.715 0 0 1 0 2.427l-.295.3a1.937 1.937 0 0 0 0 2.736l1.72 1.721a1.983 1.983 0 0 0 2.736 0l.29-.29a1.719 1.719 0 0 1 2.428 0l3.139 3.139a1.724 1.724 0 0 1 .492 1.455a1.669 1.669 0 0 1-.822 1.239a11.327 11.327 0 0 1-5.658 1.501M6.042 4.063a.793.793 0 0 0-.1.006a.673.673 0 0 0-.5.331a10.375 10.375 0 0 0 14.152 14.167a.674.674 0 0 0 .331-.5a.734.734 0 0 0-.208-.618l-3.139-3.139a.717.717 0 0 0-1.014 0l-.29.29a3.006 3.006 0 0 1-4.15 0L9.4 12.876a2.939 2.939 0 0 1 0-4.149l.3-.3a.717.717 0 0 0 0-1.014L6.56 4.277a.729.729 0 0 0-.518-.214"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PickerEmpty(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.94 5.876a2.805 2.805 0 0 0-.84-2.01a2.856 2.856 0 0 0-3.97 0l-2.21 2.21l-.75-.75a.8.8 0 0 0-1.1 0a.785.785 0 0 0 0 1.1l.75.75l-8.77 8.76a3.248 3.248 0 0 0-.92 2.13l-.07 1.52a1.311 1.311 0 0 0 .38.97a1.332 1.332 0 0 0 .91.38h.06l1.52-.07a3.248 3.248 0 0 0 2.13-.92l8.76-8.77l.75.75a.8.8 0 0 0 1.1 0a.785.785 0 0 0 0-1.1l-.75-.75L20.1 7.9a2.828 2.828 0 0 0 .84-2.024M7.35 19.236a2.22 2.22 0 0 1-1.46.63l-1.53.07a.243.243 0 0 1-.21-.09a.3.3 0 0 1-.09-.21l.07-1.53a2.22 2.22 0 0 1 .63-1.46l8.77-8.76l2.59 2.59ZM19.4 7.2l-2.18 2.18l-2.59-2.59l2.21-2.22a1.861 1.861 0 0 1 2.56 0a1.846 1.846 0 0 1 .54 1.31a1.869 1.869 0 0 1-.54 1.32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PickerHalf(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.936 5.889a2.825 2.825 0 0 0-4.81-2.02l-2.21 2.22l-.75-.75a.771.771 0 0 0-.55-.22a.8.8 0 0 0-.55.22a.785.785 0 0 0 0 1.1l.75.75l-8.76 8.76a3.154 3.154 0 0 0-.92 2.13l-.07 1.52a1.316 1.316 0 0 0 1.28 1.35h.06l1.52-.07a3.21 3.21 0 0 0 2.13-.93l8.76-8.76l.75.75a.8.8 0 0 0 1.1 0a.785.785 0 0 0 0-1.1l-.75-.75l2.18-2.18a2.845 2.845 0 0 0 .84-2.02m-8.56 8.33H7.2l6.33-6.32l2.59 2.59ZM19.4 7.2l-2.18 2.19l-2.594-2.59l2.21-2.22a1.823 1.823 0 0 1 2.56 0a1.859 1.859 0 0 1 0 2.62Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pill(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.31 4.691a5.5 5.5 0 0 0-7.78 0l-6.84 6.84a5.5 5.5 0 0 0 3.89 9.39a5.524 5.524 0 0 0 3.89-1.61l6.84-6.84a5.5 5.5 0 0 0 0-7.78m-.71 7.07l-3.42 3.42l-6.36-6.36L12.24 5.4a4.5 4.5 0 0 1 7.68 3.17a4.429 4.429 0 0 1-1.32 3.191"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PillsBottleOne(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.435 2.06H6.565a2.5 2.5 0 0 0-2.5 2.5v2a1.492 1.492 0 0 0 1.22 1.47v11.41a2.5 2.5 0 0 0 2.5 2.5h8.43a2.5 2.5 0 0 0 2.5-2.5V8.03a1.492 1.492 0 0 0 1.22-1.47v-2a2.5 2.5 0 0 0-2.5-2.5m.28 17.38a1.5 1.5 0 0 1-1.5 1.5h-8.43a1.5 1.5 0 0 1-1.5-1.5v-.88h3.52a.491.491 0 0 0 .48-.5a.485.485 0 0 0-.48-.5h-3.52V15h2.57a.5.5 0 0 0 0-1h-2.57v-2.55h3.52a.491.491 0 0 0 .48-.5a.485.485 0 0 0-.48-.5h-3.52V8.06h11.43Zm1.22-12.88a.5.5 0 0 1-.5.5H5.565a.5.5 0 0 1-.5-.5v-2a1.5 1.5 0 0 1 1.5-1.5h10.87a1.5 1.5 0 0 1 1.5 1.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pizza(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.807 13.437l-.01-.04a19.05 19.05 0 0 0-10.23-10.21a1.574 1.574 0 0 0-2.08.93l-5.32 14.69a1.58 1.58 0 0 0 1.48 2.12a1.654 1.654 0 0 0 .54-.09l14.7-5.32a1.585 1.585 0 0 0 .91-.85a1.547 1.547 0 0 0 .01-1.23m-6.98 2.98a1 1 0 0 0 .2.16L4.847 19.9a.582.582 0 0 1-.6-.14a.556.556 0 0 1-.14-.61l2.39-6.6a1 1 0 0 0 .16.2a1.81 1.81 0 0 0 2.56-2.56a1.782 1.782 0 0 0-1.7-.47l1.09-2.98a17.346 17.346 0 0 1 6.82 5.57a2.447 2.447 0 0 0-1.6.71a2.4 2.4 0 0 0 0 3.397m6.05-2.15a.592.592 0 0 1-.33.31l-1.32.47c-.11-.23-.22-.45-.33-.67c-.12-.24-.25-.48-.38-.71c-.31-.55-.65-1.08-1-1.58a18.655 18.655 0 0 0-7.57-6.3l.48-1.33a.561.561 0 0 1 .31-.33a.456.456 0 0 1 .23-.05a.793.793 0 0 1 .25.05a18.013 18.013 0 0 1 9.67 9.68v.02a.561.561 0 0 1-.01.44"/><circle cx="7.835" cy="16.489" r="1.075" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plane(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.9 20.936h-1.05a.911.911 0 0 1-.9-1.023l.693-5.548H7.3l-.513.9a1.329 1.329 0 0 1-.992.657L5.1 16a.9.9 0 0 1-.8-.31a.912.912 0 0 1-.185-.839l.774-2.769a.318.318 0 0 0 0-.173l-.775-2.764A.909.909 0 0 1 5.1 8l.695.083a1.331 1.331 0 0 1 .992.656l.513.9h3.34l-.694-5.551a.911.911 0 0 1 .9-1.024H11.9a1.327 1.327 0 0 1 1.236.857l2.144 5.714h1.046a8.5 8.5 0 0 1 1.758.184a2.166 2.166 0 0 1 1.429.9a2.209 2.209 0 0 1 .365 1.7A2.288 2.288 0 0 1 17.95 14.2l-.16.024a10.926 10.926 0 0 1-1.721.137h-.787l-2.144 5.714a1.327 1.327 0 0 1-1.238.861m-.951-1h.951a.323.323 0 0 0 .3-.209l2.214-5.905a.71.71 0 0 1 .661-.457h.991a9.946 9.946 0 0 0 1.567-.125l.16-.025a1.3 1.3 0 0 0 1.1-.979a1.227 1.227 0 0 0-.2-.937a1.2 1.2 0 0 0-.793-.5a7.647 7.647 0 0 0-1.577-.167h-1.25a.711.711 0 0 1-.661-.456L12.2 4.273a.323.323 0 0 0-.3-.209h-.951l.722 5.778a.7.7 0 0 1-.7.793H7.127a.7.7 0 0 1-.614-.359l-.6-1.045a.32.32 0 0 0-.241-.16L5.113 9l.738 2.64a1.34 1.34 0 0 1 0 .711L5.113 15l.562-.067a.32.32 0 0 0 .241-.16l.6-1.049a.7.7 0 0 1 .612-.355h3.846a.7.7 0 0 1 .7.794Zm-3.563-5.724v.005ZM7.381 9.78"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayOne(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.562 21.94a2.5 2.5 0 0 1-2.5-2.5V4.56A2.5 2.5 0 0 1 7.978 2.5l10.877 7.439a2.5 2.5 0 0 1 0 4.12L7.977 21.5a2.5 2.5 0 0 1-1.415.44m0-18.884a1.494 1.494 0 0 0-.7.177a1.477 1.477 0 0 0-.8 1.327v14.879a1.5 1.5 0 0 0 2.35 1.235l10.877-7.44a1.5 1.5 0 0 0 0-2.471L7.413 3.326a1.491 1.491 0 0 0-.849-.27Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlugOne(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.305 5.755H15.5v-3.21a.5.5 0 0 0-1 0v3.21h-5v-3.21a.5.5 0 0 0-1 0v3.21H6.7a1.566 1.566 0 0 0-1.57 1.57v4.28a7.046 7.046 0 0 0 6.37 7.11v2.72a.5.5 0 0 0 1 0v-2.7a6.874 6.874 0 0 0 6.38-6.86v-4.55a1.573 1.573 0 0 0-1.575-1.57m.57 6.12a5.875 5.875 0 0 1-6.06 5.87a6.054 6.054 0 0 1-5.69-6.14v-4.28a.563.563 0 0 1 .57-.57h10.61a.57.57 0 0 1 .57.57Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Power(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.638 5.209a8.782 8.782 0 1 0 13.917 8.96a8.871 8.871 0 0 0-3.189-8.96c-.5-.39-1.214.312-.707.707a7.93 7.93 0 0 1 3.082 7.113a7.787 7.787 0 0 1-15.308.956a7.9 7.9 0 0 1 2.912-8.069c.507-.394-.205-1.1-.707-.707"/><path fill="currentColor" d="M12.5 12.519a.5.5 0 0 1-1 0V3.548a.5.5 0 0 1 1 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rainbow(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.089 16.71A9 9 0 0 1 8.97 8.326a8.912 8.912 0 0 1 11.941 8.384a.5.5 0 0 0 1 0a10.033 10.033 0 0 0-6.46-9.291a9.981 9.981 0 0 0-11.06 2.944a10.058 10.058 0 0 0-2.3 6.347a.5.5 0 0 0 1 0Z"/><path fill="currentColor" d="M5.985 16.71A6.078 6.078 0 0 1 12 10.7a6.078 6.078 0 0 1 6.015 6.015a.5.5 0 0 0 1 0a7.013 7.013 0 0 0-12.409-4.487a7.151 7.151 0 0 0-1.621 4.482a.5.5 0 0 0 1 0"/><path fill="currentColor" d="M8.88 16.71a3.12 3.12 0 0 1 6.24 0a.5.5 0 0 0 1 0a4.119 4.119 0 0 0-7.255-2.669a4.219 4.219 0 0 0-.985 2.669a.5.5 0 0 0 1 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Read(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 18.883a10.8 10.8 0 0 1-9.675-5.728a2.6 2.6 0 0 1 0-2.31A10.8 10.8 0 0 1 12 5.117a10.8 10.8 0 0 1 9.675 5.728a2.6 2.6 0 0 1 0 2.31A10.8 10.8 0 0 1 12 18.883m0-12.766a9.787 9.787 0 0 0-8.78 5.176a1.586 1.586 0 0 0 0 1.415A9.788 9.788 0 0 0 12 17.883a9.787 9.787 0 0 0 8.78-5.176a1.584 1.584 0 0 0 0-1.414A9.787 9.787 0 0 0 12 6.117"/><path fill="currentColor" d="M12 16.049A4.049 4.049 0 1 1 16.049 12A4.054 4.054 0 0 1 12 16.049m0-7.1A3.049 3.049 0 1 0 15.049 12A3.052 3.052 0 0 0 12 8.951Z"/><circle cx="12" cy="12" r="2.028" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Receipt(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 21.919a1.454 1.454 0 0 1-.791-.232l-1.564-1.021a.47.47 0 0 0-.439-.028l-1.776.829a1.466 1.466 0 0 1-1.4-.087a1.214 1.214 0 0 1-.581-1.02V3.641a1.217 1.217 0 0 1 .584-1.021a1.469 1.469 0 0 1 1.4-.087l1.775.829a.469.469 0 0 0 .439-.026l1.563-1.023a1.464 1.464 0 0 1 1.581 0l1.564 1.022a.469.469 0 0 0 .44.026l1.775-.829a1.461 1.461 0 0 1 1.4.087a1.217 1.217 0 0 1 .581 1.021v16.72a1.216 1.216 0 0 1-.581 1.02a1.46 1.46 0 0 1-1.4.087l-1.77-.828a.474.474 0 0 0-.441.027l-1.564 1.021a1.448 1.448 0 0 1-.795.232M9.4 19.6a1.44 1.44 0 0 1 .79.234l1.564 1.02a.464.464 0 0 0 .487 0l1.565-1.021a1.462 1.462 0 0 1 1.41-.095l1.774.828a.463.463 0 0 0 .437-.024a.221.221 0 0 0 .118-.177V3.641a.219.219 0 0 0-.118-.177a.461.461 0 0 0-.437-.025l-1.775.829a1.458 1.458 0 0 1-1.409-.095l-1.563-1.022a.467.467 0 0 0-.486 0l-1.565 1.021a1.467 1.467 0 0 1-1.41.1l-1.775-.833a.461.461 0 0 0-.437.025a.219.219 0 0 0-.118.177V20.36a.221.221 0 0 0 .118.177a.468.468 0 0 0 .437.024l1.776-.829A1.461 1.461 0 0 1 9.4 19.6"/><path fill="currentColor" d="M15.046 7.4H8.954a.5.5 0 0 1 0-1h6.092a.5.5 0 0 1 0 1m0 3.553H8.954a.5.5 0 0 1 0-1h6.092a.5.5 0 0 1 0 1M12 14.5H8.954a.5.5 0 0 1 0-1H12a.5.5 0 0 1 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Redo(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.061 13.67A7.958 7.958 0 0 0 16.2 19.74a8.061 8.061 0 0 0 3.77-6.77a.5.5 0 0 0-1 0a6.976 6.976 0 0 1-11 5.7a6.969 6.969 0 0 1 1-11.97a10.075 10.075 0 0 1 4.64-.69v1.45a.5.5 0 0 0 .81.39l2.47-1.95a.5.5 0 0 0 0-.79L14.4 3.17a.5.5 0 0 0-.8.4v1.44c-.71-.01-1.43-.03-2.13.02a7.985 7.985 0 0 0-7.41 8.64Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Repeat(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.078 17.562a.493.493 0 0 1 .131-.476l2-2a.509.509 0 0 1 .707 0c.199.183.185.522 0 .707l-1.147 1.146h15.669a1.5 1.5 0 0 0 1.5-1.5V12a.5.5 0 0 1 1 0v3.439a2.5 2.5 0 0 1-2.5 2.5H3.769l1.147 1.147a.509.509 0 0 1 0 .707c-.183.199-.522.185-.707 0l-2-2a.466.466 0 0 1-.131-.231M21.923 6.457a.499.499 0 0 1-.132.476l-2 2a.509.509 0 0 1-.707 0c-.199-.183-.185-.522 0-.707l1.147-1.147H4.562a1.5 1.5 0 0 0-1.5 1.5v3.439a.5.5 0 0 1-1 0V8.579a2.5 2.5 0 0 1 2.5-2.5h15.669l-1.146-1.146a.509.509 0 0 1 0-.707c.183-.199.522-.185.707 0l2 2c.065.063.11.143.131.231"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RollingSuitcase(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.25 7.49H15V3.56a1.5 1.5 0 0 0-1.5-1.5h-3A1.511 1.511 0 0 0 9 3.56v3.93H7.75a2.5 2.5 0 0 0-2.5 2.5v8.44a2.5 2.5 0 0 0 2.5 2.5h.5v.01a1 1 0 0 0 2 0v-.01h3.5v.01a1 1 0 0 0 2 0v-.01h.5a2.5 2.5 0 0 0 2.5-2.5V9.99a2.5 2.5 0 0 0-2.5-2.5M10 3.56a.508.508 0 0 1 .5-.5h3a.5.5 0 0 1 .5.5V7.5h-4Zm7.75 14.87a1.5 1.5 0 0 1-1.5 1.5h-8.5a1.5 1.5 0 0 1-1.5-1.5V9.99a1.511 1.511 0 0 1 1.5-1.5h8.5a1.5 1.5 0 0 1 1.5 1.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Route(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.792 17.086c-.58-.58-1.16-1.17-1.75-1.75c-.08-.08-.16-.17-.25-.25a.492.492 0 0 0-.7 0a.5.5 0 0 0 0 .71l1.14 1.14H9.282a2.22 2.22 0 0 1 0-4.44h3a3.215 3.215 0 1 0 0-6.43h-5.27a2.5 2.5 0 1 0 0 1h5.27a2.215 2.215 0 1 1 0 4.43h-3a3.22 3.22 0 1 0 0 6.44h10.96l-.9.9c-.09.08-.17.17-.25.25a.5.5 0 0 0 0 .71a.511.511 0 0 0 .7 0l1.75-1.75l.25-.25a.5.5 0 0 0 0-.71m-17.23-9.02a1.5 1.5 0 1 1 1.5-1.5a1.5 1.5 0 0 1-1.5 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Router(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.44 10.94h-1.51v-6.4a.5.5 0 0 0-1 0v6.4H7.06V7a.5.5 0 0 0-1 0v3.94h-1.5a2.507 2.507 0 0 0-2.5 2.5v4a2.514 2.514 0 0 0 2.5 2.5h14.88a2.507 2.507 0 0 0 2.5-2.5v-4a2.5 2.5 0 0 0-2.5-2.5m1.5 6.5a1.5 1.5 0 0 1-1.5 1.5H4.56a1.511 1.511 0 0 1-1.5-1.5v-4a1.5 1.5 0 0 1 1.5-1.5h14.88a1.5 1.5 0 0 1 1.5 1.5Z"/><circle cx="4.75" cy="15.436" r=".75" fill="currentColor"/><circle cx="8.25" cy="15.436" r=".75" fill="currentColor"/><path fill="currentColor" d="M18.5 16.936h-5a1.5 1.5 0 1 1 0-3h5a1.5 1.5 0 0 1 0 3m-5-2a.5.5 0 1 0 0 1h5a.5.5 0 0 0 0-1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ruler(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.479 3.5a1.5 1.5 0 0 0-2.12 0L3.5 7.35a1.516 1.516 0 0 0-.44 1.06a1.5 1.5 0 0 0 .44 1.06L14.519 20.5a1.509 1.509 0 0 0 2.13 0l3.85-3.86a1.491 1.491 0 0 0 0-2.12Zm-1.12 3.58a.5.5 0 0 0 0 .71a.524.524 0 0 0 .71 0c.55-.56 1.09-1.1 1.65-1.64l1.25 1.25l-.9.9a.483.483 0 0 0 0 .7a.5.5 0 0 0 .71 0c.29-.3.6-.6.9-.89l1.25 1.25l-1.64 1.65a.495.495 0 0 0 .7.7c.56-.55 1.1-1.09 1.65-1.64l1.25 1.25l-.9.9a.524.524 0 0 0-.14.36a.5.5 0 0 0 .14.35a.513.513 0 0 0 .71 0l.9-.9l1.26 1.26l-1.65 1.64a.5.5 0 0 0 .71.71c.55-.56 1.09-1.1 1.65-1.64l1.23 1.23a.5.5 0 0 1 0 .7l-3.86 3.86a.5.5 0 0 1-.71 0L4.209 8.77a.491.491 0 0 1-.15-.36a.485.485 0 0 1 .15-.35l3.86-3.86a.508.508 0 0 1 .7 0l1.24 1.24Z"/><path fill="currentColor" d="m18.939 12.96l-.04-.04c.01 0 .01 0 .02.01s.02.02.02.03"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SatelliteOne(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.135 12.37a5.447 5.447 0 0 0 3.42-1.2a.982.982 0 0 0 .37-.72a1.04 1.04 0 0 0-.31-.8l-2.78-2.78c.39-.39.8-.8 1.19-1.2c.08-.07.15-.14.23-.22a.511.511 0 0 0 0-.7a.5.5 0 0 0-.71 0c-.48.47-.94.94-1.42 1.41l-2.78-2.78a1.077 1.077 0 0 0-.8-.31a1 1 0 0 0-.72.37a5.454 5.454 0 0 0-1.19 3.67l-1.45 1.46l-2.33-2.33a.978.978 0 0 0-1.41 0l-3.08 3.08a1 1 0 0 0 0 1.41L5.7 13.06l-.41.4a2.65 2.65 0 0 0 0 3.74l1.51 1.51a2.632 2.632 0 0 0 3.74 0l.4-.4l2.33 2.33a1 1 0 0 0 1.41 0l3.08-3.09a1 1 0 0 0 0-1.41l-2.32-2.32l1.45-1.46a2.09 2.09 0 0 0 .245.01m-13.07-2.34l3.09-3.09l2.32 2.33L6.4 12.35Zm12.99 6.82l-3.08 3.08l-2.33-2.33l3.08-3.08Zm-5.23-8.51a5.482 5.482 0 0 0 3.84 3.83l-5.84 5.84a1.642 1.642 0 0 1-2.32 0l-1.52-1.52a1.642 1.642 0 0 1 0-2.32Zm2.12 1.71a4.417 4.417 0 0 1-.3-5.96l3.13 3.13l3.14 3.14l.02.03a4.5 4.5 0 0 1-5.99-.34"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SaveDownOne(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.617 6.383a7.944 7.944 0 0 1-1.748 12.568a8.028 8.028 0 0 1-11.586-5.043a8.028 8.028 0 0 1 2.095-7.517c.451-.46-.256-1.168-.707-.707a8.946 8.946 0 0 0 9.756 14.586a8.946 8.946 0 0 0 2.9-14.594c-.451-.461-1.158.247-.707.707Z"/><path fill="currentColor" d="m15.355 10.6l-3 3a.5.5 0 0 1-.35.15a.508.508 0 0 1-.36-.15l-3-3a.5.5 0 0 1 .71-.71l2.14 2.14V3.555a.508.508 0 0 1 .5-.5a.5.5 0 0 1 .5.5v8.49l2.15-2.16a.5.5 0 0 1 .71.71Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SaveDownTwo(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.437 20.948H5.563a2.372 2.372 0 0 1-2.5-2.21v-11a2.372 2.372 0 0 1 2.5-2.211h.462a.5.5 0 0 1 0 1h-.462a1.38 1.38 0 0 0-1.5 1.211v11a1.38 1.38 0 0 0 1.5 1.21h12.874a1.38 1.38 0 0 0 1.5-1.21v-11a1.38 1.38 0 0 0-1.5-1.211h-.462a.5.5 0 0 1 0-1h.462a2.372 2.372 0 0 1 2.5 2.211v11a2.372 2.372 0 0 1-2.5 2.21"/><path fill="currentColor" d="m15.355 10.592l-3 3a.5.5 0 0 1-.35.15a.508.508 0 0 1-.36-.15l-3-3a.5.5 0 0 1 .71-.71l2.14 2.139V3.552a.508.508 0 0 1 .5-.5a.5.5 0 0 1 .5.5v8.49l2.15-2.16a.5.5 0 0 1 .71.71"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SaveUpOne(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.617 6.374a7.946 7.946 0 0 1-1.748 12.569A8.028 8.028 0 0 1 4.283 13.9a8.029 8.029 0 0 1 2.095-7.518c.451-.46-.256-1.168-.707-.707a8.946 8.946 0 0 0 9.756 14.587a8.946 8.946 0 0 0 2.9-14.595c-.451-.46-1.158.247-.707.707Z"/><path fill="currentColor" d="m8.645 6.213l3-3a.5.5 0 0 1 .35-.15a.508.508 0 0 1 .36.15l3 3a.5.5 0 0 1-.71.71l-2.14-2.14v8.47a.508.508 0 0 1-.5.5a.5.5 0 0 1-.5-.5v-8.49l-2.15 2.16a.5.5 0 0 1-.71-.71"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SaveUpTwo(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.437 20.937H5.563a2.372 2.372 0 0 1-2.5-2.211v-11a2.372 2.372 0 0 1 2.5-2.212h.462a.5.5 0 0 1 0 1h-.462a1.381 1.381 0 0 0-1.5 1.212v11a1.38 1.38 0 0 0 1.5 1.211h12.874a1.38 1.38 0 0 0 1.5-1.211v-11a1.381 1.381 0 0 0-1.5-1.212h-.462a.5.5 0 0 1 0-1h.462a2.372 2.372 0 0 1 2.5 2.212v11a2.372 2.372 0 0 1-2.5 2.211"/><path fill="currentColor" d="m8.645 6.213l3-3a.5.5 0 0 1 .35-.15a.508.508 0 0 1 .36.15l3 3a.5.5 0 0 1-.71.71l-2.14-2.14v8.47a.508.508 0 0 1-.5.5a.5.5 0 0 1-.5-.5v-8.49l-2.15 2.16a.5.5 0 0 1-.71-.71"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Search(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.031 20.79c.46.46 1.17-.25.71-.7l-3.75-3.76a7.904 7.904 0 0 0 2.04-5.31c0-4.39-3.57-7.96-7.96-7.96s-7.96 3.57-7.96 7.96c0 4.39 3.57 7.96 7.96 7.96c1.98 0 3.81-.73 5.21-1.94zM4.11 11.02c0-3.84 3.13-6.96 6.96-6.96c3.84 0 6.96 3.12 6.96 6.96s-3.12 6.96-6.96 6.96c-3.83 0-6.96-3.12-6.96-6.96"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Server(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.437 11H3.563a1.5 1.5 0 0 1-1.5-1.5V5.565a1.5 1.5 0 0 1 1.5-1.5h16.874a1.5 1.5 0 0 1 1.5 1.5v3.93a1.5 1.5 0 0 1-1.5 1.505M3.563 5.065a.5.5 0 0 0-.5.5v3.93a.5.5 0 0 0 .5.5h16.874a.5.5 0 0 0 .5-.5v-3.93a.5.5 0 0 0-.5-.5Zm16.874 14.87H3.563a1.5 1.5 0 0 1-1.5-1.5v-3.93a1.5 1.5 0 0 1 1.5-1.5h16.874a1.5 1.5 0 0 1 1.5 1.5v3.93a1.5 1.5 0 0 1-1.5 1.5m-16.874-5.93a.5.5 0 0 0-.5.5v3.93a.5.5 0 0 0 .5.5h16.874a.5.5 0 0 0 .5-.5v-3.93a.5.5 0 0 0-.5-.5Z"/><circle cx="5.563" cy="7.53" r=".5" fill="currentColor"/><circle cx="7.563" cy="7.53" r=".5" fill="currentColor"/><path fill="currentColor" d="M13.452 8.03a.5.5 0 0 1 0-1h5a.5.5 0 0 1 0 1Z"/><circle cx="5.563" cy="16.47" r=".5" fill="currentColor"/><circle cx="7.563" cy="16.47" r=".5" fill="currentColor"/><path fill="currentColor" d="M13.452 16.97a.5.5 0 0 1 0-1h5a.5.5 0 0 1 0 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Settings(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.6 20.936h-1.3a.883.883 0 0 1-.852-.654l-.774-2.833l-2.5 1.435a.886.886 0 0 1-1.06-.138l-.925-.919a.884.884 0 0 1-.143-1.066l1.469-2.545l-.006-.016l-2.787-.747a.882.882 0 0 1-.654-.851V11.3a.882.882 0 0 1 .652-.85l2.839-.777L5.12 7.171a.885.885 0 0 1 .141-1.062l.918-.918a.885.885 0 0 1 1.061-.142l2.552 1.465h.012l.745-2.79a.881.881 0 0 1 .851-.655h1.3a.883.883 0 0 1 .852.655l.762 2.838l2.509-1.441a.885.885 0 0 1 1.059.138l.926.919a.882.882 0 0 1 .141 1.067l-1.466 2.532l.008.022l2.786.746a.883.883 0 0 1 .653.851v1.3a.883.883 0 0 1-.654.852l-2.837.774l1.439 2.505a.881.881 0 0 1-.141 1.063l-.917.917a.888.888 0 0 1-1.063.141l-2.539-1.462l-.018.014l-.745 2.785a.885.885 0 0 1-.855.651m-1.21-1h1.119l.738-2.756a.888.888 0 0 1 .528-.592l.134-.052a.873.873 0 0 1 .76.057l2.51 1.445l.789-.789l-1.423-2.478a.881.881 0 0 1-.048-.78l.052-.125a.875.875 0 0 1 .584-.51l2.8-.749v-1.12l-2.755-.737a.885.885 0 0 1-.592-.529l-.052-.132a.882.882 0 0 1 .057-.763l1.449-2.508l-.8-.79l-2.48 1.425a.878.878 0 0 1-.772.052l-.115-.047a.888.888 0 0 1-.518-.588l-.748-2.806h-1.115l-.738 2.762a.883.883 0 0 1-.539.6l-.12.045a.874.874 0 0 1-.751-.058L6.822 5.962l-.789.789l1.422 2.476a.886.886 0 0 1 .046.785l-.051.12a.876.876 0 0 1-.579.5l-2.8.758v1.121l2.757.738a.889.889 0 0 1 .591.525l.048.121a.874.874 0 0 1-.055.77l-1.454 2.516l.8.791l2.47-1.419a.878.878 0 0 1 .787-.045l.106.044a.874.874 0 0 1 .526.591Zm-1.64-2.454h.008Zm-.15-.061h.007Zm4.655-10.885"/><path fill="currentColor" d="M12 15a3 3 0 1 1 3-3a3 3 0 0 1-3 3m0-5a2 2 0 1 0 2 2a2 2 0 0 0-2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareOne(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.223 11.075a.5.5 0 0 0 .7.71l7-7v3.58a.508.508 0 0 0 .5.5a.5.5 0 0 0 .5-.5v-4.79a.5.5 0 0 0-.5-.5h-4.79a.5.5 0 0 0 0 1h3.58Z"/><path fill="currentColor" d="M17.876 20.926H6.124a3.053 3.053 0 0 1-3.05-3.05V6.124a3.053 3.053 0 0 1 3.05-3.05h6.028a.5.5 0 0 1 0 1H6.124a2.053 2.053 0 0 0-2.05 2.05v11.752a2.053 2.053 0 0 0 2.05 2.05h11.752a2.053 2.053 0 0 0 2.05-2.05v-6.027a.5.5 0 0 1 1 0v6.027a3.053 3.053 0 0 1-3.05 3.05"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareTwo(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.44 15.94a2.5 2.5 0 0 0-1.96.95l-8.51-4.25a2.356 2.356 0 0 0 0-1.29l8.5-4.25a2.5 2.5 0 1 0-.53-1.54a2.269 2.269 0 0 0 .09.65l-8.5 4.25a2.5 2.5 0 1 0 0 3.08l8.5 4.25a2.269 2.269 0 0 0-.09.65a2.5 2.5 0 1 0 2.5-2.5m0-11.88a1.5 1.5 0 1 1-1.5 1.5a1.5 1.5 0 0 1 1.5-1.5M5.56 13.5a1.5 1.5 0 1 1 1.5-1.5a1.5 1.5 0 0 1-1.5 1.5m12.88 6.44a1.5 1.5 0 1 1 1.5-1.5a1.5 1.5 0 0 1-1.5 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shirt(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.657 21.949H7.343a2.5 2.5 0 0 1-2.5-2.5v-8.227a6.468 6.468 0 0 1 .112-1.2l.269-1.432a5.572 5.572 0 0 0 .094-1.015V3.3a1.252 1.252 0 0 1 1.25-1.25h2.247a1.251 1.251 0 0 1 1.25 1.25v3.113a1.935 1.935 0 0 0 3.87 0V3.3a1.251 1.251 0 0 1 1.25-1.25h2.247a1.252 1.252 0 0 1 1.25 1.25v4.275a5.486 5.486 0 0 0 .1 1.015l.269 1.431a6.57 6.57 0 0 1 .111 1.2v8.227a2.5 2.5 0 0 1-2.505 2.501M6.568 3.051a.251.251 0 0 0-.25.25v4.274a6.543 6.543 0 0 1-.111 1.2l-.27 1.432a5.5 5.5 0 0 0-.094 1.015v8.227a1.5 1.5 0 0 0 1.5 1.5h9.314a1.5 1.5 0 0 0 1.5-1.5v-8.227a5.519 5.519 0 0 0-.094-1.016l-.269-1.43a6.453 6.453 0 0 1-.112-1.2V3.3a.251.251 0 0 0-.25-.25h-2.247a.251.251 0 0 0-.25.25v3.113a2.935 2.935 0 0 1-5.87 0V3.3a.251.251 0 0 0-.25-.25Z"/><path fill="currentColor" d="M11.986 17.333v-3.459a.075.075 0 0 0-.114-.064l-.638.392a.149.149 0 0 1-.228-.128v-.65a.3.3 0 0 1 .145-.258l.764-.457a.3.3 0 0 1 .154-.043h.631a.3.3 0 0 1 .3.3v4.367a.3.3 0 0 1-.3.3h-.409a.3.3 0 0 1-.305-.3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shop(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.6 5.26a2.512 2.512 0 0 0-2.48-2.2H5.885a2.512 2.512 0 0 0-2.48 2.19l-.3 2.47a3.411 3.411 0 0 0 1.16 2.56v8.16a2.5 2.5 0 0 0 2.5 2.5h10.47a2.5 2.5 0 0 0 2.5-2.5v-8.16A3.411 3.411 0 0 0 20.9 7.72Zm-6.59 14.68h-4v-4.08a1.5 1.5 0 0 1 1.5-1.5h1a1.5 1.5 0 0 1 1.5 1.5Zm4.73-1.5a1.5 1.5 0 0 1-1.5 1.5h-2.23v-4.08a2.5 2.5 0 0 0-2.5-2.5h-1a2.5 2.5 0 0 0-2.5 2.5v4.08H6.765a1.5 1.5 0 0 1-1.5-1.5v-7.57a3.223 3.223 0 0 0 1.24.24a3.358 3.358 0 0 0 2.58-1.19a.241.241 0 0 1 .34 0a3.358 3.358 0 0 0 2.58 1.19A3.393 3.393 0 0 0 14.6 9.92a.219.219 0 0 1 .16-.07a.238.238 0 0 1 .17.07a3.358 3.358 0 0 0 2.58 1.19a3.173 3.173 0 0 0 1.23-.24Zm-1.23-8.33a2.386 2.386 0 0 1-1.82-.83a1.2 1.2 0 0 0-.92-.43h-.01a1.194 1.194 0 0 0-.92.42a2.476 2.476 0 0 1-3.65 0a1.24 1.24 0 0 0-1.86 0A2.405 2.405 0 0 1 4.1 7.78l.3-2.4a1.517 1.517 0 0 1 1.49-1.32h12.23a1.5 1.5 0 0 1 1.49 1.32l.29 2.36a2.392 2.392 0 0 1-2.395 2.37Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBasket(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.44 7.937H17.3l-1.21-4.51a.508.508 0 0 0-.61-.35a.489.489 0 0 0-.35.61l1.14 4.25H7.74l1.14-4.25a.5.5 0 0 0-.36-.61a.513.513 0 0 0-.61.35l-1.2 4.51H4.56a1.5 1.5 0 0 0-.32 2.96l.74 7.77a2.492 2.492 0 0 0 2.49 2.27h9.06a2.492 2.492 0 0 0 2.49-2.27l.74-7.77a1.5 1.5 0 0 0-.32-2.96m-1.41 10.64a1.5 1.5 0 0 1-1.5 1.36H7.47a1.5 1.5 0 0 1-1.5-1.36l-.72-7.64h13.5Zm1.41-8.64H4.56a.508.508 0 0 1-.5-.5a.5.5 0 0 1 .5-.5h14.88a.5.5 0 0 1 .5.5a.508.508 0 0 1-.5.5"/><path fill="currentColor" d="M9.5 17.432a.5.5 0 0 1-.5-.5v-3a.5.5 0 0 1 1 0v3a.5.5 0 0 1-.5.5m5 0a.5.5 0 0 1-.5-.5v-3a.5.5 0 0 1 1 0v3a.5.5 0 0 1-.5.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingCart(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.437 19.934a1 1 0 1 1-2 0a1 1 0 0 1 2 0M6.22 15.668l-.945-10.9a.75.75 0 0 0-.749-.693H3.56a.5.5 0 0 1 0-1h.966a1.75 1.75 0 0 1 1.746 1.617l.139 1.818h13.03c.885 0 1.577.76 1.494 1.638l-.668 7.52a2.5 2.5 0 0 1-2.489 2.267H8.709a2.5 2.5 0 0 1-2.489-2.267m.274-8.158l.722 8.066a1.5 1.5 0 0 0 1.493 1.359h9.069a1.5 1.5 0 0 0 1.493-1.359l.668-7.518a.5.5 0 0 0-.498-.548zm4.454 12.424a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingTag(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.605 5.988a2.8 2.8 0 0 0-2.6-2.59l-4.56-.32a2.842 2.842 0 0 0-2.17.81L3.9 11.278a2.794 2.794 0 0 0 0 3.95l4.87 4.88a2.8 2.8 0 0 0 3.96 0l7.38-7.39a2.779 2.779 0 0 0 .81-2.17ZM12.015 19.4a1.8 1.8 0 0 1-2.54 0l-4.87-4.87a1.793 1.793 0 0 1 0-2.55l1.17-1.17l7.42 7.42Zm7.38-7.38l-5.5 5.5l-7.41-7.42l5.5-5.5a1.786 1.786 0 0 1 1.27-.53c.04 0 .08.01.12.01l4.56.32a1.8 1.8 0 0 1 1.67 1.66l.32 4.56a1.829 1.829 0 0 1-.525 1.398Z"/><circle cx="17" cy="6.999" r=".862" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shuffle(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.794 17.082a.513.513 0 0 1 0 .71c-.08.08-.17.16-.25.25c-.58.58-1.17 1.16-1.75 1.75a.5.5 0 0 1-.71-.71c.09-.08.17-.17.25-.25l.9-.9h-3.85a2.509 2.509 0 0 1-2.15-1.21L12 12.982l-2.24 3.74a2.509 2.509 0 0 1-2.15 1.21H2.564a.5.5 0 0 1-.5-.5a.5.5 0 0 1 .5-.5h5.05A1.507 1.507 0 0 0 8.9 16.2l2.51-4.2L8.9 7.792a1.507 1.507 0 0 0-1.29-.73H2.564a.5.5 0 0 1 0-1h5.05a2.518 2.518 0 0 1 2.15 1.22L12 11.032l2.24-3.75a2.489 2.489 0 0 1 2.14-1.22h3.85l-1.15-1.15a.5.5 0 1 1 .71-.7c.08.08.17.16.25.25c.58.58 1.17 1.16 1.75 1.75a.5.5 0 0 1 0 .7c-.08.09-.17.17-.25.25c-.58.59-1.17 1.17-1.75 1.75a.5.5 0 1 1-.71-.7l.25-.25l.9-.9h-3.85a1.519 1.519 0 0 0-1.29.73L12.584 12l2.51 4.2a1.519 1.519 0 0 0 1.29.73h3.84l-1.14-1.14a.5.5 0 0 1 .71-.71Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignpostDuoOne(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.437 11.5H6.536a2.628 2.628 0 0 1-1.723-.629L2.564 8.915a1.329 1.329 0 0 1 .006-2.084L4.813 4.88a2.619 2.619 0 0 1 1.723-.629h13.9a1.451 1.451 0 0 1 1.5 1.393v4.463a1.451 1.451 0 0 1-1.499 1.393m-13.9-6.25a1.64 1.64 0 0 0-1.067.384L3.215 7.6a.364.364 0 0 0-.152.281a.349.349 0 0 0 .141.27l.011.01l2.254 1.961a1.644 1.644 0 0 0 1.067.384h13.9a.463.463 0 0 0 .5-.394V5.644a.463.463 0 0 0-.5-.393Zm10.927 14.499H3.563a1.451 1.451 0 0 1-1.5-1.394v-4.463a1.451 1.451 0 0 1 1.5-1.393h13.9a2.621 2.621 0 0 1 1.724.63l2.249 1.956a1.329 1.329 0 0 1-.007 2.083l-2.242 1.951a2.625 2.625 0 0 1-1.723.63M3.563 13.5a.463.463 0 0 0-.5.393v4.463a.463.463 0 0 0 .5.394h13.9a1.644 1.644 0 0 0 1.068-.385l2.253-1.96a.362.362 0 0 0 .152-.28a.351.351 0 0 0-.141-.271l-.011-.01l-2.254-1.96a1.64 1.64 0 0 0-1.067-.384Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignpostLone(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.437 16H6.536a2.489 2.489 0 0 1-1.744-.709L2.542 13.1a1.5 1.5 0 0 1 .007-2.2l2.243-2.191A2.483 2.483 0 0 1 6.536 8h13.9a1.5 1.5 0 0 1 1.5 1.5v5a1.5 1.5 0 0 1-1.499 1.5M6.536 9a1.491 1.491 0 0 0-1.046.425l-2.255 2.2a.5.5 0 0 0-.172.375a.494.494 0 0 0 .162.369l.01.01l2.254 2.2A1.492 1.492 0 0 0 6.536 15h13.9a.5.5 0 0 0 .5-.5v-5a.5.5 0 0 0-.5-.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignpostRone(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.563 8h13.9a2.489 2.489 0 0 1 1.744.709l2.25 2.192a1.5 1.5 0 0 1-.007 2.2l-2.243 2.187a2.483 2.483 0 0 1-1.743.712H3.563a1.5 1.5 0 0 1-1.5-1.5v-5a1.5 1.5 0 0 1 1.5-1.5m13.9 7a1.491 1.491 0 0 0 1.046-.425l2.255-2.2a.5.5 0 0 0 .173-.375a.494.494 0 0 0-.162-.369l-.01-.01l-2.254-2.2A1.492 1.492 0 0 0 17.464 9H3.563a.5.5 0 0 0-.5.5v5a.5.5 0 0 0 .5.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SliderHorizontal(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.45 7.5H9.99A2 2 0 0 0 8.06 6h-1a2 2 0 0 0-1.93 1.5H2.55a.5.5 0 0 0-.5.5a.508.508 0 0 0 .5.5h2.58A2 2 0 0 0 7.06 10h1a2 2 0 0 0 1.93-1.5h11.46a.5.5 0 0 0 0-1M8.06 9h-1a1.006 1.006 0 0 1-1-.98V8a1 1 0 0 1 1-1h1a1 1 0 1 1 0 2m13.39 6.5h-2.58a2 2 0 0 0-1.93-1.5h-1a2 2 0 0 0-1.93 1.5H2.55a.5.5 0 0 0 0 1h11.46a2 2 0 0 0 1.93 1.5h1a2 2 0 0 0 1.93-1.5h2.58a.5.5 0 0 0 .5-.5a.508.508 0 0 0-.5-.5m-3.51.5a1 1 0 0 1-1 1h-1a1 1 0 1 1 0-2h1a1.006 1.006 0 0 1 1 .98Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SliderVertical(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.5 5.13V2.55a.5.5 0 0 0-.5-.5a.508.508 0 0 0-.5.5v2.58A2 2 0 0 0 6 7.06v1a2 2 0 0 0 1.5 1.93v11.46a.5.5 0 0 0 1 0V9.99A2 2 0 0 0 10 8.06v-1a2 2 0 0 0-1.5-1.93M9 8.06a1 1 0 1 1-2 0v-1a1.006 1.006 0 0 1 .98-1H8a1 1 0 0 1 1 1Zm7.5 5.95V2.55a.5.5 0 0 0-1 0v11.46a2 2 0 0 0-1.5 1.93v1a2 2 0 0 0 1.5 1.93v2.58a.5.5 0 0 0 .5.5a.508.508 0 0 0 .5-.5v-2.58a2 2 0 0 0 1.5-1.93v-1a2 2 0 0 0-1.5-1.93m.5 2.93a1.006 1.006 0 0 1-.98 1H16a1 1 0 0 1-1-1v-1a1 1 0 1 1 2 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Speaker(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.437 21.938H6.562a2.5 2.5 0 0 1-2.5-2.5V4.562a2.5 2.5 0 0 1 2.5-2.5h10.875a2.5 2.5 0 0 1 2.5 2.5v14.876a2.5 2.5 0 0 1-2.5 2.5M6.562 3.062a1.5 1.5 0 0 0-1.5 1.5v14.876a1.5 1.5 0 0 0 1.5 1.5h10.875a1.5 1.5 0 0 0 1.5-1.5V4.562a1.5 1.5 0 0 0-1.5-1.5Z"/><path fill="currentColor" d="M12 18.939a3.75 3.75 0 1 1 3.75-3.75a3.755 3.755 0 0 1-3.75 3.75m0-6.5a2.75 2.75 0 1 0 2.75 2.75a2.752 2.752 0 0 0-2.75-2.75m0-2.876a2.25 2.25 0 1 1 2.25-2.25A2.253 2.253 0 0 1 12 9.563m0-3.5a1.25 1.25 0 1 0 1.25 1.25A1.251 1.251 0 0 0 12 6.063"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareAlert(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.437 20.936H5.562a2.5 2.5 0 0 1-2.5-2.5V5.562a2.5 2.5 0 0 1 2.5-2.5h12.875a2.5 2.5 0 0 1 2.5 2.5v12.874a2.5 2.5 0 0 1-2.5 2.5M5.562 4.062a1.5 1.5 0 0 0-1.5 1.5v12.874a1.5 1.5 0 0 0 1.5 1.5h12.875a1.5 1.5 0 0 0 1.5-1.5V5.562a1.5 1.5 0 0 0-1.5-1.5Z"/><path fill="currentColor" d="M11.5 10.982a.5.5 0 0 1 1 0V15a.5.5 0 0 1-1 0Z"/><circle cx="12" cy="9" r=".5" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareCheck(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.437 20.939H5.563a2.5 2.5 0 0 1-2.5-2.5V5.566a2.5 2.5 0 0 1 2.5-2.5h12.874a2.5 2.5 0 0 1 2.5 2.5v12.873a2.5 2.5 0 0 1-2.5 2.5M5.563 4.066a1.5 1.5 0 0 0-1.5 1.5v12.873a1.5 1.5 0 0 0 1.5 1.5h12.874a1.5 1.5 0 0 0 1.5-1.5V5.566a1.5 1.5 0 0 0-1.5-1.5Z"/><path fill="currentColor" d="M15.81 10.4c.45-.46-.26-1.17-.71-.71l-3.56 3.56c-.58-.58-1.16-1.15-1.73-1.73a.5.5 0 0 0-.71.71l2.08 2.08a.513.513 0 0 0 .71 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareChevDown(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.65 10.85a.495.495 0 0 1 .7-.7L12 12.79l2.65-2.64a.495.495 0 0 1 .7.7l-3 3a.492.492 0 0 1-.7 0Z"/><path fill="currentColor" d="M3.063 18.437V5.563a2.5 2.5 0 0 1 2.5-2.5h12.874a2.5 2.5 0 0 1 2.5 2.5v12.874a2.5 2.5 0 0 1-2.5 2.5H5.563a2.5 2.5 0 0 1-2.5-2.5M19.937 5.563a1.5 1.5 0 0 0-1.5-1.5H5.563a1.5 1.5 0 0 0-1.5 1.5v12.874a1.5 1.5 0 0 0 1.5 1.5h12.874a1.5 1.5 0 0 0 1.5-1.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareChevLeft(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m11.21 12l2.64 2.65a.495.495 0 0 1-.7.7c-.13-.12-.25-.24-.38-.37c-.87-.87-1.75-1.75-2.62-2.63a.492.492 0 0 1 0-.7l3-3a.495.495 0 0 1 .7.7Z"/><path fill="currentColor" d="M18.437 20.939H5.562a2.5 2.5 0 0 1-2.5-2.5V5.566a2.5 2.5 0 0 1 2.5-2.5h12.875a2.5 2.5 0 0 1 2.5 2.5v12.873a2.5 2.5 0 0 1-2.5 2.5M5.562 4.066a1.5 1.5 0 0 0-1.5 1.5v12.873a1.5 1.5 0 0 0 1.5 1.5h12.875a1.5 1.5 0 0 0 1.5-1.5V5.566a1.5 1.5 0 0 0-1.5-1.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareChevRight(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.85 15.35a.495.495 0 0 1-.7-.7L12.79 12l-2.64-2.65a.495.495 0 0 1 .7-.7l3 3a.492.492 0 0 1 0 .7Z"/><path fill="currentColor" d="M18.437 20.937H5.563a2.5 2.5 0 0 1-2.5-2.5V5.563a2.5 2.5 0 0 1 2.5-2.5h12.874a2.5 2.5 0 0 1 2.5 2.5v12.874a2.5 2.5 0 0 1-2.5 2.5M5.563 4.063a1.5 1.5 0 0 0-1.5 1.5v12.874a1.5 1.5 0 0 0 1.5 1.5h12.874a1.5 1.5 0 0 0 1.5-1.5V5.563a1.5 1.5 0 0 0-1.5-1.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareChevUp(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.35 13.15a.495.495 0 0 1-.7.7L12 11.21l-2.65 2.64a.495.495 0 0 1-.7-.7l3-3a.492.492 0 0 1 .7 0Z"/><path fill="currentColor" d="M20.937 5.563v12.874a2.5 2.5 0 0 1-2.5 2.5H5.563a2.5 2.5 0 0 1-2.5-2.5V5.563a2.5 2.5 0 0 1 2.5-2.5h12.874a2.5 2.5 0 0 1 2.5 2.5M4.063 18.437a1.5 1.5 0 0 0 1.5 1.5h12.874a1.5 1.5 0 0 0 1.5-1.5V5.563a1.5 1.5 0 0 0-1.5-1.5H5.563a1.5 1.5 0 0 0-1.5 1.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareInfo(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.438 20.937H5.564a2.5 2.5 0 0 1-2.5-2.5V5.563a2.5 2.5 0 0 1 2.5-2.5h12.874a2.5 2.5 0 0 1 2.5 2.5v12.874a2.5 2.5 0 0 1-2.5 2.5M5.564 4.063a1.5 1.5 0 0 0-1.5 1.5v12.874a1.5 1.5 0 0 0 1.5 1.5h12.874a1.5 1.5 0 0 0 1.5-1.5V5.563a1.5 1.5 0 0 0-1.5-1.5Z"/><path fill="currentColor" d="M12.5 9a.5.5 0 0 0-1 0v4.018a.5.5 0 0 0 1 0Z"/><circle cx="12" cy="14.999" r=".5" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareMinus(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.438 20.938H5.564a2.5 2.5 0 0 1-2.5-2.5V5.564a2.5 2.5 0 0 1 2.5-2.5h12.874a2.5 2.5 0 0 1 2.5 2.5v12.874a2.5 2.5 0 0 1-2.5 2.5M5.564 4.064a1.5 1.5 0 0 0-1.5 1.5v12.874a1.5 1.5 0 0 0 1.5 1.5h12.874a1.5 1.5 0 0 0 1.5-1.5V5.564a1.5 1.5 0 0 0-1.5-1.5Z"/><path fill="currentColor" d="M9 12.5a.5.5 0 0 1 0-1h6a.5.5 0 0 1 0 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareMore(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.437 20.938H5.562a2.5 2.5 0 0 1-2.5-2.5V5.565a2.5 2.5 0 0 1 2.5-2.5h12.875a2.5 2.5 0 0 1 2.5 2.5v12.873a2.5 2.5 0 0 1-2.5 2.5M5.562 4.065a1.5 1.5 0 0 0-1.5 1.5v12.873a1.5 1.5 0 0 0 1.5 1.5h12.875a1.5 1.5 0 0 0 1.5-1.5V5.565a1.5 1.5 0 0 0-1.5-1.5Z"/><circle cx="11.999" cy="12.002" r="1" fill="currentColor"/><circle cx="15.999" cy="12.002" r="1" fill="currentColor"/><circle cx="7.999" cy="12.002" r="1" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquarePlus(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.438 20.938H5.563a2.5 2.5 0 0 1-2.5-2.5V5.564a2.5 2.5 0 0 1 2.5-2.5h12.875a2.5 2.5 0 0 1 2.5 2.5v12.874a2.5 2.5 0 0 1-2.5 2.5M5.563 4.064a1.5 1.5 0 0 0-1.5 1.5v12.874a1.5 1.5 0 0 0 1.5 1.5h12.875a1.5 1.5 0 0 0 1.5-1.5V5.564a1.5 1.5 0 0 0-1.5-1.5Z"/><path fill="currentColor" d="M15 12.5h-2.5V15a.5.5 0 0 1-1 0v-2.5H9a.5.5 0 0 1 0-1h2.5V9a.5.5 0 0 1 1 0v2.5H15a.5.5 0 0 1 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareQuestion(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.41 12.461a1.555 1.555 0 0 1 .341-.6a2.68 2.68 0 0 1 .535-.417a2.2 2.2 0 0 0 .363-.285a1.218 1.218 0 0 0 .256-.364a1.083 1.083 0 0 0 .095-.451a.927.927 0 0 0-.142-.518a.946.946 0 0 0-.374-.338a1.135 1.135 0 0 0-.519-.119a1.188 1.188 0 0 0-.5.107a.934.934 0 0 0-.389.335a.884.884 0 0 0-.111.224a.515.515 0 0 1-.483.359a.506.506 0 0 1-.479-.675a1.653 1.653 0 0 1 .178-.348a1.789 1.789 0 0 1 .748-.634a2.609 2.609 0 0 1 2.113.015a1.733 1.733 0 0 1 .721.642a1.766 1.766 0 0 1 .257.959a1.833 1.833 0 0 1-.118.678a1.674 1.674 0 0 1-.334.536a2.289 2.289 0 0 1-.52.417a2.245 2.245 0 0 0-.462.37a1.1 1.1 0 0 0-.256.454a2.344 2.344 0 0 0-.045.283a.486.486 0 0 1-.483.429a.484.484 0 0 1-.483-.53a2.928 2.928 0 0 1 .091-.529"/><circle cx="11.792" cy="14.894" r=".587" fill="currentColor"/><path fill="currentColor" d="M18.438 20.938H5.563a2.5 2.5 0 0 1-2.5-2.5V5.564a2.5 2.5 0 0 1 2.5-2.5h12.875a2.5 2.5 0 0 1 2.5 2.5v12.874a2.5 2.5 0 0 1-2.5 2.5M5.563 4.064a1.5 1.5 0 0 0-1.5 1.5v12.874a1.5 1.5 0 0 0 1.5 1.5h12.875a1.5 1.5 0 0 0 1.5-1.5V5.564a1.5 1.5 0 0 0-1.5-1.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareRemove(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.437 20.937H5.563a2.5 2.5 0 0 1-2.5-2.5V5.563a2.5 2.5 0 0 1 2.5-2.5h12.874a2.5 2.5 0 0 1 2.5 2.5v12.874a2.5 2.5 0 0 1-2.5 2.5M5.563 4.063a1.5 1.5 0 0 0-1.5 1.5v12.874a1.5 1.5 0 0 0 1.5 1.5h12.874a1.5 1.5 0 0 0 1.5-1.5V5.563a1.5 1.5 0 0 0-1.5-1.5Z"/><path fill="currentColor" d="M13.767 14.477a.5.5 0 0 0 .71-.71L12.707 12l1.77-1.77a.5.5 0 0 0-.71-.7L12 11.3l-1.77-1.77a.5.5 0 0 0-.7.7c.59.59 1.17 1.18 1.77 1.77l-1.77 1.77c-.46.45.25 1.16.7.71L12 12.707Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Star(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.6 20.463a1.5 1.5 0 0 1-.7-.174l-3.666-1.927a.5.5 0 0 0-.464 0L8.1 20.289a1.5 1.5 0 0 1-2.177-1.581l.7-4.082a.5.5 0 0 0-.143-.442l-2.964-2.891a1.5 1.5 0 0 1 .832-2.559l4.1-.6a.5.5 0 0 0 .376-.273l1.833-3.714a1.5 1.5 0 0 1 2.69 0l1.833 3.714a.5.5 0 0 0 .376.274l4.1.6a1.5 1.5 0 0 1 .832 2.559l-2.965 2.891a.5.5 0 0 0-.144.442l.7 4.082a1.5 1.5 0 0 1-1.479 1.754m-3.9-2.986l3.664 1.923a.5.5 0 0 0 .725-.527l-.7-4.082a1.5 1.5 0 0 1 .432-1.328l2.965-2.89a.5.5 0 0 0-.277-.853l-4.1-.6a1.5 1.5 0 0 1-1.13-.821l-1.83-3.705a.516.516 0 0 0-.9 0l-1.83 3.714a1.5 1.5 0 0 1-1.13.82l-4.1.6a.5.5 0 0 0-.277.853l2.968 2.887a1.5 1.5 0 0 1 .431 1.332l-.7 4.082a.5.5 0 0 0 .726.527l3.663-1.932a1.5 1.5 0 0 1 1.4 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stethoscope(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.185 9.256a2.748 2.748 0 0 0-.5 5.45v2.31a2.923 2.923 0 0 1-2.92 2.92h-2.78a2.923 2.923 0 0 1-2.92-2.92v-.98a5.5 5.5 0 0 0 5-5.47v-5.28a1.483 1.483 0 0 0-1.03-1.42l-2.31-.78a.5.5 0 0 0-.63.32a.491.491 0 0 0 .31.63l2.32.78a.486.486 0 0 1 .34.47v5.28a4.5 4.5 0 0 1-9 0v-5.28a.486.486 0 0 1 .34-.47l2.32-.78a.491.491 0 0 0 .31-.63a.5.5 0 0 0-.63-.32l-2.31.78a1.483 1.483 0 0 0-1.03 1.42v5.28a5.5 5.5 0 0 0 5 5.47v.98a3.926 3.926 0 0 0 3.92 3.92h2.78a3.926 3.926 0 0 0 3.92-3.92v-2.31a2.748 2.748 0 0 0-.5-5.45m0 4.5a1.75 1.75 0 1 1 1.75-1.75a1.758 1.758 0 0 1-1.75 1.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StickyNote(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.44 3.065H5.56a2.507 2.507 0 0 0-2.5 2.5v12.87a2.507 2.507 0 0 0 2.5 2.5h8.68A2.482 2.482 0 0 0 16 20.2l4.21-4.2a2.505 2.505 0 0 0 .73-1.77V5.565a2.5 2.5 0 0 0-2.5-2.5m-4.38 13.5v3.37h-8.5a1.5 1.5 0 0 1-1.5-1.5V5.565a1.5 1.5 0 0 1 1.5-1.5h12.88a1.5 1.5 0 0 1 1.5 1.5v8.5h-3.38a2.507 2.507 0 0 0-2.5 2.5m1 3.13v-3.13a1.5 1.5 0 0 1 1.5-1.5h3.13Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopOne(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.437 20.937H5.563a2.5 2.5 0 0 1-2.5-2.5V5.563a2.5 2.5 0 0 1 2.5-2.5h12.874a2.5 2.5 0 0 1 2.5 2.5v12.874a2.5 2.5 0 0 1-2.5 2.5M5.563 4.063a1.5 1.5 0 0 0-1.5 1.5v12.874a1.5 1.5 0 0 0 1.5 1.5h12.874a1.5 1.5 0 0 0 1.5-1.5V5.563a1.5 1.5 0 0 0-1.5-1.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopSignOne(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.531 15.688H10.14a.5.5 0 0 1 0-1h2.391a1.094 1.094 0 0 0 0-2.188h-1.063a2.094 2.094 0 0 1 0-4.188h2.391a.5.5 0 0 1 0 1h-2.391a1.094 1.094 0 0 0 0 2.188h1.063a2.094 2.094 0 0 1 0 4.188"/><path fill="currentColor" d="M15.079 21.933H8.92a2.482 2.482 0 0 1-1.767-.733L2.8 16.847a2.484 2.484 0 0 1-.732-1.768V8.921A2.486 2.486 0 0 1 2.8 7.153L7.153 2.8a2.482 2.482 0 0 1 1.767-.733h6.159a2.482 2.482 0 0 1 1.767.732L21.2 7.154a2.482 2.482 0 0 1 .732 1.767v6.158a2.491 2.491 0 0 1-.731 1.768L16.846 21.2a2.482 2.482 0 0 1-1.767.733M8.92 3.067a1.511 1.511 0 0 0-1.06.439L3.506 7.861a1.489 1.489 0 0 0-.439 1.06v6.158a1.491 1.491 0 0 0 .439 1.061l4.354 4.354a1.511 1.511 0 0 0 1.06.439h6.159a1.511 1.511 0 0 0 1.06-.439l4.355-4.354a1.494 1.494 0 0 0 .439-1.061V8.921a1.511 1.511 0 0 0-.439-1.06l-4.355-4.355a1.511 1.511 0 0 0-1.06-.439Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stopwatch(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.925 7.828a7.862 7.862 0 0 1 1.97 5.215A7.898 7.898 0 0 1 12 20.938a7.899 7.899 0 0 1-7.896-7.895c0-4.189 3.271-7.621 7.396-7.879V4.061H9.913c-.645 0-.643-1 0-1h4.174c.645 0 .644 1 0 1H12.5v1.103a7.865 7.865 0 0 1 4.718 1.956l1.135-1.134a.509.509 0 0 1 .707 0c.199.183.185.522 0 .707zm.97 5.215A6.898 6.898 0 0 0 12 6.148a6.9 6.9 0 0 0-6.896 6.895A6.898 6.898 0 0 0 12 19.938a6.898 6.898 0 0 0 6.895-6.895m-6.395.001c0 .645-1 .643-1 0V8.34c0-.644 1-.643 1 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StreamOff(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.911 4.206c-.45-.45-1.16.26-.71.71l.32.32v.01A10.646 10.646 0 0 0 5.541 19.8c.48.43 1.19-.28.71-.71a9.623 9.623 0 0 1-1.01-13.13l2.27 2.27a6.022 6.022 0 0 0 .61 8.18c.48.44 1.19-.27.71-.7a5.024 5.024 0 0 1-.61-6.77l2.61 2.61a1.13 1.13 0 0 0-.09.45a1.248 1.248 0 0 0 1.25 1.24a1.13 1.13 0 0 0 .45-.09l4.77 4.77l.86.86a3.024 3.024 0 0 1-.31.31a.355.355 0 0 0-.11.16a.406.406 0 0 0-.04.19a.381.381 0 0 0 .04.19a.386.386 0 0 0 .11.17a.5.5 0 0 0 .35.14a.585.585 0 0 0 .13-.02a.432.432 0 0 0 .22-.12c.11-.1.22-.2.32-.3c.1.09.19.19.29.29c.45.45 1.16-.26.71-.71Zm12.819.702a9.624 9.624 0 0 1 2.3 11.1c-.265.582.6 1.09.864.505A10.647 10.647 0 0 0 18.438 4.2c-.475-.433-1.185.272-.708.707Zm-2.58 3.383a5.016 5.016 0 0 1 1.6 4.572a.515.515 0 0 0 .349.615a.5.5 0 0 0 .615-.349a6.042 6.042 0 0 0-1.852-5.546c-.476-.431-1.185.274-.708.708Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StreamOn(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.26 19.089a9.625 9.625 0 0 1-.026-14.178C6.709 4.475 6 3.769 5.527 4.2a10.516 10.516 0 0 0 .026 15.6c.475.433 1.184-.273.707-.707Z"/><path fill="currentColor" d="M8.84 15.706a5.024 5.024 0 0 1-.014-7.412c.474-.437-.234-1.143-.707-.707a6.028 6.028 0 0 0 .014 8.826c.474.434 1.183-.272.707-.707"/><circle cx="12" cy="12" r="1.244" fill="currentColor"/><path fill="currentColor" d="M17.74 4.911a9.625 9.625 0 0 1 .026 14.178c-.475.436.234 1.142.707.707A10.516 10.516 0 0 0 18.447 4.2c-.475-.433-1.184.273-.707.707Z"/><path fill="currentColor" d="M15.16 8.294a5.024 5.024 0 0 1 .014 7.412c-.474.437.234 1.143.707.707a6.028 6.028 0 0 0-.014-8.826c-.474-.434-1.183.272-.707.707"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sun(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 20.937a1.074 1.074 0 0 1-.94-.542L9.61 17.9a.084.084 0 0 0-.1-.041l-2.782.741A1.087 1.087 0 0 1 5.4 17.272l.748-2.8a.088.088 0 0 0-.041-.1l-2.5-1.439a1.086 1.086 0 0 1 0-1.881L6.1 9.61a.087.087 0 0 0 .041-.1L5.4 6.728A1.087 1.087 0 0 1 6.728 5.4l2.8.748a.091.091 0 0 0 .1-.041l1.439-2.5A1.076 1.076 0 0 1 12 3.063a1.074 1.074 0 0 1 .94.542L14.39 6.1a.084.084 0 0 0 .1.041l2.782-.741A1.087 1.087 0 0 1 18.6 6.728l-.748 2.8a.087.087 0 0 0 .041.1l2.5 1.439a1.086 1.086 0 0 1 0 1.881L17.9 14.39a.089.089 0 0 0-.041.1l.748 2.784a1.087 1.087 0 0 1-1.335 1.326l-2.8-.748a.089.089 0 0 0-.1.041l-1.439 2.5a1.076 1.076 0 0 1-.94.544Zm-2.466-4.084a1.091 1.091 0 0 1 .942.541l1.448 2.5a.082.082 0 0 0 .075.043a.081.081 0 0 0 .074-.043l1.44-2.5a1.083 1.083 0 0 1 1.221-.507l2.8.747a.087.087 0 0 0 .106-.106l-.747-2.785a1.089 1.089 0 0 1 .5-1.222l2.5-1.448a.086.086 0 0 0 0-.15l-2.5-1.439a1.086 1.086 0 0 1-.507-1.221l.747-2.8a.08.08 0 0 0-.022-.083a.086.086 0 0 0-.085-.023l-2.784.747a1.088 1.088 0 0 1-1.222-.5l-1.448-2.5A.082.082 0 0 0 12 4.063a.081.081 0 0 0-.074.043l-1.44 2.5a1.087 1.087 0 0 1-1.222.507l-2.8-.747a.087.087 0 0 0-.106.106l.752 2.782a1.089 1.089 0 0 1-.5 1.222l-2.5 1.448a.082.082 0 0 0-.047.076a.081.081 0 0 0 .043.074l2.5 1.44a1.087 1.087 0 0 1 .507 1.221l-.747 2.8a.08.08 0 0 0 .022.083a.087.087 0 0 0 .085.023l2.784-.747a1.077 1.077 0 0 1 .277-.041"/><path fill="currentColor" d="M12 15.875A3.875 3.875 0 1 1 15.875 12A3.88 3.88 0 0 1 12 15.875m0-6.75A2.875 2.875 0 1 0 14.875 12A2.879 2.879 0 0 0 12 9.125"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TabletsOne(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.76 3.065a6.171 6.171 0 0 0-6.11 5.58a6.159 6.159 0 1 0 6.71 6.71a6.159 6.159 0 0 0-.6-12.29m-5.53 16.87A5.166 5.166 0 0 1 5.24 11.5l7.27 7.26a5.153 5.153 0 0 1-3.28 1.175m3.99-1.88l-7.27-7.27a5.165 5.165 0 0 1 7.27 7.27m2.15-3.71a6.12 6.12 0 0 0-5.72-5.71a5.157 5.157 0 1 1 5.72 5.71"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TempHigh(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.863 13.4V4.939a2.929 2.929 0 0 0-.84-2.03a2.859 2.859 0 0 0-2.23-.82a2.948 2.948 0 0 0-2.66 3l.01 8.28a4.755 4.755 0 0 0 1.9 8.46a5.093 5.093 0 0 0 .95.09a4.759 4.759 0 0 0 4.76-4.75a4.684 4.684 0 0 0-1.89-3.769m-.48 6.66a3.783 3.783 0 0 1-3.15.78a3.7 3.7 0 0 1-2.92-2.98a3.745 3.745 0 0 1 1.43-3.69a.962.962 0 0 0 .39-.77V5.089a1.968 1.968 0 0 1 1.73-2a.66.66 0 0 1 .14-.01a1.878 1.878 0 0 1 1.86 1.86V13.4a.962.962 0 0 0 .39.77a3.742 3.742 0 0 1 .13 5.89"/><path fill="currentColor" d="M13.893 17.169a1.89 1.89 0 0 1-3.78 0a1.858 1.858 0 0 1 1.39-1.81V5.4a.5.5 0 0 1 1 0v9.96a1.869 1.869 0 0 1 1.39 1.809"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextAlignCenter(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.437 4.063H3.563a.5.5 0 1 1 0-1h16.874a.5.5 0 1 1 0 1M16.5 8.5h-9a.5.5 0 0 1 0-1h9a.5.5 0 0 1 0 1m0 8h-9a.5.5 0 1 1 0-1h9a.5.5 0 1 1 0 1m3.937-4H3.563a.5.5 0 0 1 0-1h16.874a.5.5 0 0 1 0 1m0 8.437H3.563a.5.5 0 1 1 0-1h16.874a.5.5 0 0 1 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextAlignJustify(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.437 4.064H3.563a.5.5 0 0 1 0-1h16.874a.5.5 0 0 1 0 1m0 4.436H3.563a.5.5 0 0 1 0-1h16.874a.5.5 0 0 1 0 1m0 8H3.563a.5.5 0 1 1 0-1h16.874a.5.5 0 0 1 0 1m0-4H3.563a.5.5 0 0 1 0-1h16.874a.5.5 0 0 1 0 1m0 8.436H3.563a.5.5 0 1 1 0-1h16.874a.5.5 0 1 1 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextAlignLeft(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.438 4.063H3.563a.5.5 0 1 1 0-1h16.875a.5.5 0 1 1 0 1M12.562 8.5h-9a.5.5 0 0 1 0-1h9a.5.5 0 0 1 0 1m0 8h-9a.5.5 0 1 1 0-1h9a.5.5 0 0 1 0 1m7.874-4H3.562a.5.5 0 1 1 0-1h16.874a.5.5 0 0 1 0 1m0 8.437H3.562a.5.5 0 0 1 0-1h16.874a.5.5 0 0 1 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextAlignRight(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.437 4.064H3.563a.5.5 0 0 1 0-1h16.874a.5.5 0 0 1 0 1m0 4.436h-9a.5.5 0 1 1 0-1h9a.5.5 0 0 1 0 1m0 8h-9a.5.5 0 1 1 0-1h9a.5.5 0 0 1 0 1m0-4H3.563a.5.5 0 0 1 0-1h16.874a.5.5 0 0 1 0 1m0 8.436H3.563a.5.5 0 1 1 0-1h16.874a.5.5 0 1 1 0 1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextIcon(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.44 2.06H5.56a1.5 1.5 0 0 0-1.5 1.5v4.5a.5.5 0 0 0 1 0v-1H10v13.88H8a.5.5 0 0 0 0 1h8a.5.5 0 0 0 0-1h-2V7.06h4.94v1a.5.5 0 0 0 1 0v-4.5a1.5 1.5 0 0 0-1.5-1.5m.5 4H14a1 1 0 0 0-1 1v13.88h-2V7.06a1 1 0 0 0-1-1H5.06v-2.5a.5.5 0 0 1 .5-.5h12.88a.5.5 0 0 1 .5.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Timer(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.336 9.685a9.934 9.934 0 0 0 11.256 12.123A9.931 9.931 0 0 0 20.708 7.23A10.046 10.046 0 0 0 12 2.072a.507.507 0 0 0-.5.5v4.2a.5.5 0 0 0 1 0v-4.2l-.5.5a8.935 8.935 0 0 1 8.433 11.892a8.938 8.938 0 0 1-13.965 4.063A9.041 9.041 0 0 1 3.3 9.951c.142-.627-.822-.9-.964-.266"/><path fill="currentColor" d="M7.4 8.117a.5.5 0 0 1 .707-.707l4.243 4.242a.5.5 0 0 1-.707.707Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trash(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.45 4.06h-4.18v-.5a1.5 1.5 0 0 0-1.5-1.5h-3.54a1.5 1.5 0 0 0-1.5 1.5v.5H4.55a.5.5 0 0 0 0 1h.72l.42 14.45a2.493 2.493 0 0 0 2.5 2.43h7.62a2.493 2.493 0 0 0 2.5-2.43l.42-14.45h.72a.5.5 0 0 0 0-1m-9.72-.5a.5.5 0 0 1 .5-.5h3.54a.5.5 0 0 1 .5.5v.5H9.73Zm7.58 15.92a1.5 1.5 0 0 1-1.5 1.46H8.19a1.5 1.5 0 0 1-1.5-1.46L6.26 5.06h11.48Z"/><path fill="currentColor" d="M8.375 8a.5.5 0 0 1 1 0l.25 10a.5.5 0 0 1-1 0Zm7.25.007a.5.5 0 0 0-1 0l-.25 10a.5.5 0 0 0 1 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trophy(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.435 5.055h-.97c.01-.12.02-.24.02-.36a1.645 1.645 0 0 0-.45-1.18a1.462 1.462 0 0 0-1.05-.45h-9.96a1.484 1.484 0 0 0-1.06.45a1.6 1.6 0 0 0-.44 1.18c0 .12.01.24.02.36h-.98a1.5 1.5 0 0 0-1.5 1.5v2a4.5 4.5 0 0 0 4.27 4.49c1.07 2.3 2.53 3.79 4.17 4.04v2.85h-4a.5.5 0 1 0 0 1h9a.5.5 0 0 0 0-1h-4v-2.85c1.64-.25 3.1-1.74 4.17-4.04a4.493 4.493 0 0 0 4.26-4.49v-2a1.5 1.5 0 0 0-1.5-1.5m-15.37 3.5v-2a.5.5 0 0 1 .5-.5h1.04a22.9 22.9 0 0 0 1.28 5.93a3.5 3.5 0 0 1-2.82-3.43m7.94 7.57c-2.82 0-5.23-5.04-5.48-11.47a.573.573 0 0 1 .16-.44a.48.48 0 0 1 .34-.15h9.96a.442.442 0 0 1 .33.15a.62.62 0 0 1 .17.44c-.25 6.43-2.66 11.47-5.48 11.47m7.93-7.57a3.508 3.508 0 0 1-2.8 3.42a23.353 23.353 0 0 0 1.27-5.92h1.03a.5.5 0 0 1 .5.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TurnLone(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.939 9.509v10.93a.508.508 0 0 1-.5.5a.5.5 0 0 1-.5-.5V9.509a3.5 3.5 0 0 0-3.5-3.5h-9.9l-.01 1.44a.5.5 0 0 1-.81.4l-2.47-1.96a.5.5 0 0 1 0-.78l2.49-1.94a.5.5 0 0 1 .81.4l-.01 1.44h9.9a4.507 4.507 0 0 1 4.5 4.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TurnRone(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.061 9.509v10.93a.508.508 0 0 0 .5.5a.5.5 0 0 0 .5-.5V9.509a3.5 3.5 0 0 1 3.5-3.5h9.9l.01 1.44a.5.5 0 0 0 .81.4l2.47-1.96a.5.5 0 0 0 0-.78l-2.49-1.94a.5.5 0 0 0-.81.4l.01 1.44h-9.9a4.507 4.507 0 0 0-4.5 4.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Twitter(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.913 5.322a1.034 1.034 0 0 1 .837 1.629l-1.042 1.481c-.064 5.086-1.765 8.539-5.056 10.264A10.917 10.917 0 0 1 9.6 19.835a12.233 12.233 0 0 1-6.2-1.524a.76.76 0 0 1-.317-.8a.768.768 0 0 1 .63-.6a20.6 20.6 0 0 0 3.745-.886C2 13.5 3.19 7.824 3.71 6.081a1.028 1.028 0 0 1 1.729-.422a9.931 9.931 0 0 0 5.995 2.95A4.188 4.188 0 0 1 12.725 5.3a4.125 4.125 0 0 1 5.7.02ZM4.521 17.794c1.862.872 6.226 1.819 9.667.016c2.955-1.549 4.476-4.732 4.521-9.461a.771.771 0 0 1 .142-.436l1.081-1.538l-.041-.053c-.518-.007-1.029-.014-1.55 0a.835.835 0 0 1-.547-.221a3.13 3.13 0 0 0-4.383-.072a3.174 3.174 0 0 0-.935 2.87a.646.646 0 0 1-.154.545a.591.591 0 0 1-.516.205a10.924 10.924 0 0 1-7.084-3.295c-.67 2.078-1.52 7.094 3.869 9.065a.632.632 0 0 1 .416.538a.625.625 0 0 1-.3.6a13.178 13.178 0 0 1-4.186 1.237m15.147-9.305"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Umbrella(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 4.06v-.5a.509.509 0 0 0-.15-.35a.483.483 0 0 0-.7 0a.491.491 0 0 0-.15.35v.5a8.41 8.41 0 0 0-7.88 7.82a.976.976 0 0 0 .27.74a1.029 1.029 0 0 0 .74.32h6.87v5.22a1.653 1.653 0 0 1-.62 1.54a1.528 1.528 0 0 1-2.38-1.16a.5.5 0 0 0-1 0a2.433 2.433 0 0 0 2.43 2.4a2.45 2.45 0 0 0 2.57-2.29c.08-1.39 0-2.81 0-4.2v-1.51h6.87a1.029 1.029 0 0 0 .74-.32a.976.976 0 0 0 .27-.74a8.41 8.41 0 0 0-7.88-7.82m6.87 7.88l-14.75.01a7.4 7.4 0 0 1 14.76-.02c0 .01 0 .01-.01.01"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Undo(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.939 13.67A7.958 7.958 0 0 1 7.8 19.74a8.061 8.061 0 0 1-3.77-6.77a.5.5 0 0 1 1 0a6.976 6.976 0 0 0 11 5.7a6.969 6.969 0 0 0-1-11.97a10.075 10.075 0 0 0-4.64-.69v1.45a.5.5 0 0 1-.81.39L7.109 5.9a.5.5 0 0 1 0-.79L9.6 3.17a.5.5 0 0 1 .8.4v1.44c.71-.01 1.43-.03 2.13.02a7.985 7.985 0 0 1 7.41 8.64Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unlock(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.44 9.33h-1.1c0-.97.01-1.95 0-2.92a4.343 4.343 0 0 0-7.98-2.37c-.36.53.51 1.03.87.5a3.365 3.365 0 0 1 5.23-.39c1.04 1.11.88 2.57.88 3.96v1.22H6.56a2.5 2.5 0 0 0-2.5 2.5v7.61a2.507 2.507 0 0 0 2.5 2.5h10.88a2.507 2.507 0 0 0 2.5-2.5v-7.61a2.5 2.5 0 0 0-2.5-2.5m1.5 10.11a1.511 1.511 0 0 1-1.5 1.5H6.56a1.511 1.511 0 0 1-1.5-1.5v-7.61a1.5 1.5 0 0 1 1.5-1.5h10.88a1.5 1.5 0 0 1 1.5 1.5Z"/><path fill="currentColor" d="M13 14.95a.984.984 0 0 1-.5.86v1.5a.5.5 0 0 1-1 0v-1.5a.984.984 0 0 1-.5-.86a1 1 0 0 1 2 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unread(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.182 6.314a10.215 10.215 0 0 1 7.9 1.782a9.145 9.145 0 0 1 2.5 2.817a1.875 1.875 0 0 1 .082 2.024a9.266 9.266 0 0 1-1.485 2.008c-.446.464.26 1.172.707.707c1.1-1.144 2.533-2.86 1.9-4.554a8.845 8.845 0 0 0-2.721-3.5A11.243 11.243 0 0 0 9.916 5.35c-.633.11-.364 1.074.266.964m9.642 12.796q-3.045-3.045-6.09-6.08c-.93-.93-1.85-1.86-2.77-2.77q-2.115-2.115-4.21-4.22l-1.86-1.86c-.45-.45-1.16.26-.71.71l1.9 1.9a10.42 10.42 0 0 0-3.22 3.12a3.743 3.743 0 0 0-.8 2.28a4.581 4.581 0 0 0 .99 2.17a10.925 10.925 0 0 0 8.18 4.5A11.379 11.379 0 0 0 17 17.71l.25.25l1.86 1.86c.454.45 1.164-.26.714-.71m-10.3-8.88c.25.24.49.49.73.73A2.039 2.039 0 0 0 12 14.03a2.023 2.023 0 0 0 1.04-.28c.25.24.49.49.73.73a3.047 3.047 0 0 1-4.25-4.25Zm-3.7 5.6a9.558 9.558 0 0 1-1.81-1.84c-.53-.71-1.19-1.62-.85-2.55a8.348 8.348 0 0 1 3.65-3.92c.67.67 1.34 1.33 2 2a4.04 4.04 0 0 0 5.67 5.67c.6.59 1.19 1.19 1.78 1.78a10.4 10.4 0 0 1-10.44-1.14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Usb(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.44 8.355h-2.13v-.14a1.443 1.443 0 0 0-1.44-1.45H7.29a5.235 5.235 0 0 0 0 10.47h9.58a1.443 1.443 0 0 0 1.44-1.45v-.14h2.13a1.511 1.511 0 0 0 1.5-1.5v-4.29a1.5 1.5 0 0 0-1.5-1.5m-3.13 7.43a.446.446 0 0 1-.44.45H7.29a4.235 4.235 0 0 1 0-8.47h9.58a.446.446 0 0 1 .44.45Zm3.63-1.64a.508.508 0 0 1-.5.5h-2.13v-5.29h2.13a.5.5 0 0 1 .5.5Z"/><path fill="currentColor" d="M6.29 13.444A1.446 1.446 0 1 1 7.738 12a1.447 1.447 0 0 1-1.448 1.444m0-1.892a.446.446 0 1 0 .448.448a.446.446 0 0 0-.448-.448"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func User(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.438 21.937H6.562a2.5 2.5 0 0 1-2.5-2.5v-.827c0-3.969 3.561-7.2 7.938-7.2s7.938 3.229 7.938 7.2v.827a2.5 2.5 0 0 1-2.5 2.5M12 12.412c-3.826 0-6.938 2.78-6.938 6.2v.827a1.5 1.5 0 0 0 1.5 1.5h10.876a1.5 1.5 0 0 0 1.5-1.5v-.829c0-3.418-3.112-6.198-6.938-6.198m0-2.501a3.924 3.924 0 1 1 3.923-3.924A3.927 3.927 0 0 1 12 9.911m0-6.847a2.924 2.924 0 1 0 2.923 2.923A2.926 2.926 0 0 0 12 3.064"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vault(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.935 6.06h-7.87a2 2 0 0 0-2 2v6a1.993 1.993 0 0 0 2 2h7.87a2 2 0 0 0 2-2v-6a2.006 2.006 0 0 0-2-2m1 8a1 1 0 0 1-1 1h-7.87a.99.99 0 0 1-1-1v-6a1 1 0 0 1 1-1h7.87a1 1 0 0 1 1 1Z"/><path fill="currentColor" d="M18.435 3.06H5.565a2.507 2.507 0 0 0-2.5 2.5v11a2.5 2.5 0 0 0 2.5 2.5v.38a1.5 1.5 0 0 0 1.5 1.5h1.43a1.5 1.5 0 0 0 1.5-1.5v-.38h4v.38a1.5 1.5 0 0 0 1.5 1.5h1.44a1.5 1.5 0 0 0 1.5-1.5v-.38a2.5 2.5 0 0 0 2.5-2.5v-11a2.507 2.507 0 0 0-2.5-2.5m-9.44 16.38a.5.5 0 0 1-.5.5h-1.43a.5.5 0 0 1-.5-.5v-.38h2.43Zm8.44 0a.5.5 0 0 1-.5.5H15.5a.508.508 0 0 1-.5-.5v-.38h2.44Zm2.5-2.88a1.5 1.5 0 0 1-1.5 1.5H5.565a1.5 1.5 0 0 1-1.5-1.5v-11a1.5 1.5 0 0 1 1.5-1.5h12.87a1.5 1.5 0 0 1 1.5 1.5Z"/><path fill="currentColor" d="M14.265 10.56h-.61A1.656 1.656 0 0 0 12.5 9.4v-.61a.491.491 0 0 0-.5-.48a.5.5 0 0 0-.5.48v.61a1.656 1.656 0 0 0-1.16 1.16h-.61a.5.5 0 0 0-.48.5a.491.491 0 0 0 .48.5h.61a1.656 1.656 0 0 0 1.16 1.16v.62a.489.489 0 0 0 .5.47a.483.483 0 0 0 .5-.47v-.62a1.622 1.622 0 0 0 1.16-1.16h.61a.485.485 0 0 0 .48-.5a.491.491 0 0 0-.485-.5M12 11.81a.75.75 0 1 1 .75-.75a.749.749 0 0 1-.75.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vial(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.779 9.441l-.48-.47l-5.26-5.271l-.48-.48a.5.5 0 0 0-.7 0a.5.5 0 0 0 0 .71l.47.48l-10.17 10.171a3.694 3.694 0 0 0 0 5.22l.04.04a3.706 3.706 0 0 0 5.23 0L19.6 9.671l.47.48a.52.52 0 0 0 .71 0a.513.513 0 0 0-.001-.71m-12.06 9.69a2.7 2.7 0 0 1-3.81 0l-.04-.04a2.7 2.7 0 0 1 0-3.81l1.7-1.7h7.71Zm6.56-6.55h-7.71l7.47-7.46l3.85 3.85Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoOff(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.783 9.51v6.53a.5.5 0 0 1-1 0V9.6c0-.85.06-1.72 0-2.57c-.03-.37-.27-.5-.61-.3c-.39.22-.76.51-1.13.76c-.73.49-1.47.99-2.2 1.49c0 .71-.01 1.41-.01 2.11a.5.5 0 0 1-1 0V9.12c0-.58.01-1.16 0-1.74a1.524 1.524 0 0 0-1.56-1.5c-1.22-.01-2.43 0-3.65 0a.5.5 0 0 1 0-1h2.13c.6 0 1.22-.05 1.81.01a2.54 2.54 0 0 1 2.27 2.5c0 .13.01.26 0 .39c.77-.53 1.55-1.05 2.32-1.57a2.466 2.466 0 0 1 1.26-.6a1.364 1.364 0 0 1 1.37 1.36c.03.84 0 1.7 0 2.54m-1.01 9.57q-6.5-6.51-12.99-13c-.62-.62-1.24-1.24-1.87-1.86c-.45-.46-1.16.25-.7.71l.28.28A2.468 2.468 0 0 0 3.2 7.38v9.24a2.5 2.5 0 0 0 2.5 2.5h7.63a2.5 2.5 0 0 0 2.5-2.5v-.07l1.37 1.37c.62.62 1.24 1.24 1.87 1.86a.5.5 0 0 0 .703-.7m-4.94-2.46a1.5 1.5 0 0 1-1.5 1.5H5.7a1.5 1.5 0 0 1-1.5-1.5V7.38a1.5 1.5 0 0 1 1.04-1.42l9.59 9.59Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoOn(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.05 5.05a1.433 1.433 0 0 0-1.48.08l-3.32 2.24v-.81a2.5 2.5 0 0 0-2.5-2.5H4.69a2.5 2.5 0 0 0-2.5 2.5v10.88a2.5 2.5 0 0 0 2.5 2.5h9.06a2.5 2.5 0 0 0 2.5-2.5v-.81l3.32 2.24a1.5 1.5 0 0 0 .81.24a1.414 1.414 0 0 0 1.43-1.43V6.32a1.437 1.437 0 0 0-.76-1.27m-5.8 12.39a1.5 1.5 0 0 1-1.5 1.5H4.69a1.5 1.5 0 0 1-1.5-1.5V6.56a1.5 1.5 0 0 1 1.5-1.5h9.06a1.5 1.5 0 0 1 1.5 1.5Zm5.56.24a.415.415 0 0 1-.23.38a.425.425 0 0 1-.45-.02l-3.88-2.62V8.58l3.88-2.62a.425.425 0 0 1 .45-.02a.415.415 0 0 1 .23.38Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewBoard(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.437 20.936H5.563a2.5 2.5 0 0 1-2.5-2.5V5.562a2.5 2.5 0 0 1 2.5-2.5h12.874a2.5 2.5 0 0 1 2.5 2.5v12.874a2.5 2.5 0 0 1-2.5 2.5M5.563 4.062a1.5 1.5 0 0 0-1.5 1.5v12.874a1.5 1.5 0 0 0 1.5 1.5h12.874a1.5 1.5 0 0 0 1.5-1.5V5.562a1.5 1.5 0 0 0-1.5-1.5Z"/><path fill="currentColor" d="M12.5 14.544a.5.5 0 0 1-1 0v-8a.5.5 0 0 1 1 0Zm4.217-2.091a.5.5 0 0 1-1 0V6.544a.5.5 0 0 1 1 0ZM8.28 6.544a.5.5 0 0 0-1 0v4a.5.5 0 0 0 1 0Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewColumn(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.44 3.06H5.56a2.507 2.507 0 0 0-2.5 2.5v12.88a2.514 2.514 0 0 0 2.5 2.5h12.88a2.514 2.514 0 0 0 2.5-2.5V5.56a2.507 2.507 0 0 0-2.5-2.5M8.67 19.94H5.56a1.511 1.511 0 0 1-1.5-1.5V5.56a1.5 1.5 0 0 1 1.5-1.5h3.11Zm1-15.88h4.66v15.88H9.67Zm10.27 14.38a1.511 1.511 0 0 1-1.5 1.5h-3.11V4.06h3.11a1.5 1.5 0 0 1 1.5 1.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewList(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.436 20.937H5.562a2.5 2.5 0 0 1-2.5-2.5V5.563a2.5 2.5 0 0 1 2.5-2.5h12.874a2.5 2.5 0 0 1 2.5 2.5v12.874a2.5 2.5 0 0 1-2.5 2.5M5.562 4.063a1.5 1.5 0 0 0-1.5 1.5v12.874a1.5 1.5 0 0 0 1.5 1.5h12.874a1.5 1.5 0 0 0 1.5-1.5V5.563a1.5 1.5 0 0 0-1.5-1.5Z"/><path fill="currentColor" d="M6.544 8.283a.519.519 0 0 1-.353-.147a.5.5 0 0 1 0-.707a.512.512 0 0 1 .353-.146H7.55a.516.516 0 0 1 .353.146a.5.5 0 0 1 .147.354a.5.5 0 0 1-.5.5Zm0 4.217a.523.523 0 0 1-.353-.146a.5.5 0 0 1 0-.708a.516.516 0 0 1 .353-.146H7.55a.521.521 0 0 1 .353.146a.5.5 0 0 1 0 .708a.516.516 0 0 1-.353.146Zm0 4.22a.519.519 0 0 1-.353-.147a.5.5 0 0 1 0-.707a.516.516 0 0 1 .353-.146H7.55a.516.516 0 0 1 .353.146a.5.5 0 0 1 .147.354a.5.5 0 0 1-.5.5Zm4.01-8.439a.5.5 0 0 1 0-1h6.9a.5.5 0 0 1 0 1Zm0 4.219a.5.5 0 0 1 0-1h6.9a.5.5 0 0 1 0 1Zm0 4.218a.5.5 0 0 1 0-1h6.9a.5.5 0 0 1 0 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewTable(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.44 3.06H5.56a2.507 2.507 0 0 0-2.5 2.5v12.88a2.507 2.507 0 0 0 2.5 2.5h12.88a2.514 2.514 0 0 0 2.5-2.5V5.56a2.514 2.514 0 0 0-2.5-2.5M8.71 19.94H5.56a1.5 1.5 0 0 1-1.5-1.5v-3.11h4.65Zm0-5.61H4.06V9.67h4.65Zm0-5.66H4.06V5.56a1.5 1.5 0 0 1 1.5-1.5h3.15Zm11.23 9.77a1.511 1.511 0 0 1-1.5 1.5H9.71v-4.61h10.23Zm0-4.11H9.71V9.67h10.23Zm0-5.66H9.71V4.06h8.73a1.511 1.511 0 0 1 1.5 1.5Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ViewTimeline(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.436 20.94H5.562a2.5 2.5 0 0 1-2.5-2.5V5.567a2.5 2.5 0 0 1 2.5-2.5h12.874a2.5 2.5 0 0 1 2.5 2.5V18.44a2.5 2.5 0 0 1-2.5 2.5M5.562 4.067a1.5 1.5 0 0 0-1.5 1.5V18.44a1.5 1.5 0 0 0 1.5 1.5h12.874a1.5 1.5 0 0 0 1.5-1.5V5.567a1.5 1.5 0 0 0-1.5-1.5Z"/><path fill="currentColor" d="M6.544 8.287a.5.5 0 0 1 0-1H12a.5.5 0 0 1 0 1ZM9.271 12.5a.5.5 0 0 1 0-1h5.454a.5.5 0 0 1 0 1ZM12 16.724a.5.5 0 0 1 0-1h5.455a.5.5 0 0 1 0 1Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Virus(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.86 19.905a.485.485 0 0 0 .35.15a.469.469 0 0 0 .35-.15a.483.483 0 0 0 0-.7l-.53-.53l1.74-1.74a6.426 6.426 0 0 0 3.73 1.54v2.46h-.75a.5.5 0 0 0 0 1h2.5a.5.5 0 1 0 0-1h-.75v-2.46a6.426 6.426 0 0 0 3.73-1.54l1.74 1.74l-.53.53a.483.483 0 0 0 0 .7a.469.469 0 0 0 .35.15a.485.485 0 0 0 .35-.15l1.77-1.76a.513.513 0 0 0 0-.71a.5.5 0 0 0-.71 0l-.52.53l-1.74-1.74a6.435 6.435 0 0 0 1.54-3.73h2.46v.75a.5.5 0 0 0 .5.5a.5.5 0 0 0 .5-.5v-2.5a.508.508 0 0 0-.5-.5a.5.5 0 0 0-.5.5v.75h-2.46a6.418 6.418 0 0 0-1.55-3.72l1.75-1.74l.52.53a.508.508 0 0 0 .36.15a.5.5 0 0 0 .35-.15a.513.513 0 0 0 0-.71l-1.77-1.77a.5.5 0 0 0-.7 0a.5.5 0 0 0 0 .71l.53.53l-1.74 1.74a6.382 6.382 0 0 0-3.73-1.55v-2.45h.75a.5.5 0 0 0 .5-.5a.5.5 0 0 0-.5-.5h-2.5a.5.5 0 0 0-.5.5a.5.5 0 0 0 .5.5h.75v2.45a6.382 6.382 0 0 0-3.73 1.55l-1.74-1.74l.53-.53a.5.5 0 0 0 0-.71a.5.5 0 0 0-.7 0l-1.77 1.77a.513.513 0 0 0 0 .71a.5.5 0 0 0 .35.15a.508.508 0 0 0 .36-.15l.52-.53l1.75 1.74A6.418 6.418 0 0 0 5.52 11.5H3.06v-.75a.5.5 0 0 0-.5-.5a.508.508 0 0 0-.5.5v2.5a.5.5 0 0 0 .5.5a.5.5 0 0 0 .5-.5v-.75h2.46a6.435 6.435 0 0 0 1.54 3.73l-1.74 1.74l-.52-.53a.5.5 0 0 0-.71 0a.513.513 0 0 0 0 .71Zm10.03-4.02A5.5 5.5 0 1 1 17.5 12a5.448 5.448 0 0 1-1.61 3.885"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Voicemail(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.435 7.5a4.5 4.5 0 0 0-2.82 8h-5.23a4.494 4.494 0 1 0-2.82 1h10.87a4.5 4.5 0 0 0 0-9M3.065 12a3.5 3.5 0 1 1 3.56 3.5h-.06a3.5 3.5 0 0 1-3.5-3.5m14.37 3.5h-.06a3.53 3.53 0 1 1 .06 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Volume(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.849 20.934a1.555 1.555 0 0 1-.781-.212l-4.16-2.4a3.769 3.769 0 0 0-1.877-.5H7.214a2.631 2.631 0 0 1-2.628-2.627V8.809a2.631 2.631 0 0 1 2.628-2.627h3.817a3.747 3.747 0 0 0 1.877-.5l4.16-2.4a1.564 1.564 0 0 1 2.346 1.354v14.733a1.57 1.57 0 0 1-1.565 1.565M7.214 7.182a1.63 1.63 0 0 0-1.628 1.627v6.382a1.629 1.629 0 0 0 1.628 1.627h3.817a4.756 4.756 0 0 1 2.377.637l4.16 2.4a.543.543 0 0 0 .563 0a.553.553 0 0 0 .283-.487V4.632a.565.565 0 0 0-.846-.489l-4.16 2.4a4.753 4.753 0 0 1-2.377.637Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeHigh(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.816 19.937a1.446 1.446 0 0 1-.717-.194l-3.669-2.12a3.257 3.257 0 0 0-1.625-.434H4.439a2.379 2.379 0 0 1-2.375-2.376V9.187a2.378 2.378 0 0 1 2.375-2.375h3.366a3.257 3.257 0 0 0 1.625-.436l3.67-2.117A1.437 1.437 0 0 1 15.255 5.5v13a1.424 1.424 0 0 1-.718 1.245a1.445 1.445 0 0 1-.721.192M4.439 7.812a1.377 1.377 0 0 0-1.375 1.375v5.626a1.378 1.378 0 0 0 1.375 1.376h3.366a4.254 4.254 0 0 1 2.125.569l3.67 2.118a.439.439 0 0 0 .657-.38V5.5a.438.438 0 0 0-.657-.379L9.93 7.242a4.251 4.251 0 0 1-2.125.57Zm13.968-1.55a7.79 7.79 0 0 1 .021 11.476c-.474.437.235 1.143.707.707a8.793 8.793 0 0 0-.021-12.89c-.474-.434-1.184.272-.707.707"/><path fill="currentColor" d="M16.91 9.031a4.021 4.021 0 0 1 .012 5.938c-.474.438.234 1.143.707.707a5.025 5.025 0 0 0-.012-7.352c-.474-.434-1.183.272-.707.707"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeMute(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.817 19.936a1.424 1.424 0 0 1-.719-.2L9.43 17.624a3.254 3.254 0 0 0-1.625-.436H4.44a2.377 2.377 0 0 1-2.375-2.375V9.187A2.378 2.378 0 0 1 4.44 6.811h3.365a3.257 3.257 0 0 0 1.625-.434l3.67-2.118A1.438 1.438 0 0 1 15.255 5.5v13a1.423 1.423 0 0 1-.718 1.245a1.439 1.439 0 0 1-.72.191M4.44 7.811a1.377 1.377 0 0 0-1.375 1.376v5.626a1.377 1.377 0 0 0 1.375 1.375h3.365a4.247 4.247 0 0 1 2.125.571l3.67 2.117a.437.437 0 0 0 .439 0a.433.433 0 0 0 .218-.379V5.5a.438.438 0 0 0-.657-.379L9.93 7.242a4.25 4.25 0 0 1-2.125.569Zm13.166 6.634a.5.5 0 0 1-.7-.711c.17-.169.34-.349.52-.52l1.21-1.209c-.57-.581-1.15-1.161-1.73-1.74a.5.5 0 0 1 .7-.71l1.74 1.739l1.74-1.739a.5.5 0 0 1 .7.71l-1.73 1.74l1.73 1.729a.5.5 0 0 1-.7.711L19.343 12.7Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wallet(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.435 4.065H4.565a2.5 2.5 0 0 0-2.5 2.5v10.87a2.5 2.5 0 0 0 2.5 2.5h14.87a2.5 2.5 0 0 0 2.5-2.5V6.565a2.5 2.5 0 0 0-2.5-2.5m1.5 9.93h-6.42a2 2 0 0 1 0-4h6.42Zm-6.42-5a3 3 0 0 0 0 6h6.42v2.44a1.5 1.5 0 0 1-1.5 1.5H4.565a1.5 1.5 0 0 1-1.5-1.5V6.565a1.5 1.5 0 0 1 1.5-1.5h14.87a1.5 1.5 0 0 1 1.5 1.5v2.43Z"/><circle cx="14.519" cy="11.996" r="1" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Warning(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.5 8.752a.5.5 0 0 0-1 0v6a.5.5 0 0 0 1 0Z"/><circle cx="11.999" cy="16.736" r=".5" fill="currentColor"/><path fill="currentColor" d="M18.642 20.934H5.385a2.5 2.5 0 0 1-2.222-3.644L9.792 4.421a2.5 2.5 0 0 1 4.444 0l6.629 12.869a2.5 2.5 0 0 1-2.223 3.644M12.014 4.065a1.478 1.478 0 0 0-1.334.814L4.052 17.748a1.5 1.5 0 0 0 1.333 2.186h13.257a1.5 1.5 0 0 0 1.334-2.186L13.348 4.879a1.478 1.478 0 0 0-1.334-.814"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WavePulseOne(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.974 18a1.446 1.446 0 0 1-1.259-.972l-1.843-4.145c-.115-.26-.262-.378-.349-.378H2.562a.5.5 0 1 1 0-1h2.961a1.444 1.444 0 0 1 1.263.972l1.839 4.145c.116.261.258.378.349.378c.088 0 .229-.113.344-.368L13.7 6.956A1.423 1.423 0 0 1 14.958 6a1.449 1.449 0 0 1 1.26.975l1.839 4.151c.11.249.259.379.349.379h3.028a.5.5 0 0 1 0 1H18.41a1.444 1.444 0 0 1-1.263-.975l-1.839-4.151c-.116-.261-.259-.378-.35-.379c-.088 0-.229.114-.344.368l-4.385 9.676A1.437 1.437 0 0 1 8.974 18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wheat(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.247 13.836a3.115 3.115 0 0 0 .79-2.78a1.053 1.053 0 0 0-.8-.81a3.1 3.1 0 0 0 .8-2.79a1.061 1.061 0 0 0-.82-.82a3.211 3.211 0 0 0-2.04.25A3.09 3.09 0 0 0 14 4.816a3.1 3.1 0 0 0-1.41-2.57a1.043 1.043 0 0 0-1.16-.01A3.124 3.124 0 0 0 10 4.816a3.052 3.052 0 0 0 .83 2.07a3.154 3.154 0 0 0-2.04-.25a1.048 1.048 0 0 0-.82.82a3.1 3.1 0 0 0 .79 2.79a1.041 1.041 0 0 0-.79.81a3.11 3.11 0 0 0 .78 2.78a1.071 1.071 0 0 0-.78.82a3.031 3.031 0 0 0 3 3.7a2.436 2.436 0 0 0 .53-.05v3.15a.5.5 0 0 0 1 0v-3.15a2.469 2.469 0 0 0 .54.05a3.054 3.054 0 0 0 2.17-.88a3.124 3.124 0 0 0 .83-2.82a1.083 1.083 0 0 0-.793-.82m-3.83 3.48a2.12 2.12 0 0 1-1.92-.55a2.041 2.041 0 0 1-.51-1.96a2.558 2.558 0 0 1 .47-.04a1.984 1.984 0 0 1 1.45.59a2.011 2.011 0 0 1 .51 1.96m0-3.6a2.112 2.112 0 0 1-1.92-.55a2.022 2.022 0 0 1-.51-1.95a1.93 1.93 0 0 1 .47-.05a1.984 1.984 0 0 1 1.45.59a2.011 2.011 0 0 1 .51 1.96m0-3.6a2.112 2.112 0 0 1-1.92-.55a2.022 2.022 0 0 1-.51-1.95a2.592 2.592 0 0 1 .47-.05a1.984 1.984 0 0 1 1.45.59a2.011 2.011 0 0 1 .51 1.96m.56-3.55A2.1 2.1 0 0 1 11 4.816a2.005 2.005 0 0 1 1.04-1.74a2.1 2.1 0 0 1 .96 1.74a2.054 2.054 0 0 1-1.023 1.75m2.53 10.2a2.072 2.072 0 0 1-1.96.51a2.384 2.384 0 0 1-.05-.45v-.02a2.065 2.065 0 0 1 .59-1.46a1.99 1.99 0 0 1 1.4-.57a2.279 2.279 0 0 1 .57.07a2.14 2.14 0 0 1-.55 1.92m0-3.6a2.047 2.047 0 0 1-1.96.51a2.384 2.384 0 0 1-.05-.45v-.02a2.024 2.024 0 0 1 .59-1.45a1.957 1.957 0 0 1 1.4-.58a2.863 2.863 0 0 1 .57.07a2.14 2.14 0 0 1-.55 1.92m0-3.6a2.024 2.024 0 0 1-1.96.51a2.384 2.384 0 0 1-.05-.45v-.02a2.024 2.024 0 0 1 .59-1.45a1.955 1.955 0 0 1 1.41-.57a2.259 2.259 0 0 1 .56.07a2.12 2.12 0 0 1-.55 1.91"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiOff(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.37 6.564a12.392 12.392 0 0 1 10.71 3.93c.436.476 1.141-.233.708-.708A13.324 13.324 0 0 0 10.37 5.564c-.631.076-.638 1.077 0 1m3.537 3.719a8.641 8.641 0 0 1 4.442 2.617c.434.477 1.139-.232.707-.707a9.586 9.586 0 0 0-4.883-2.871c-.626-.146-.893.818-.266.965Z"/><circle cx="12.003" cy="16.922" r="1.12" fill="currentColor"/><path fill="currentColor" d="M19.773 19.06a.5.5 0 0 1-.71.71l-5.84-5.84A4.478 4.478 0 0 0 8.7 15.24c-.43.48-1.14-.23-.71-.7a5.47 5.47 0 0 1 4.06-1.78l-2.37-2.37a8.693 8.693 0 0 0-4.03 2.53c-.43.48-1.13-.23-.7-.71A9.439 9.439 0 0 1 8.893 9.6l-2.01-2.01a12.557 12.557 0 0 0-3.96 2.94a.5.5 0 1 1-.7-.71a13.109 13.109 0 0 1 3.91-2.98l-1.9-1.9a.5.5 0 0 1 .71-.71Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiOn(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.922 10.777a12.194 12.194 0 0 1 18.155-.034c.436.476 1.141-.233.707-.707a13.189 13.189 0 0 0-19.569.034c-.432.475.273 1.184.707.707"/><path fill="currentColor" d="M5.654 13.169a8.615 8.615 0 0 1 12.691-.024c.437.475 1.143-.234.707-.707a9.621 9.621 0 0 0-14.106.024c-.433.474.272 1.184.708.707"/><path fill="currentColor" d="M8.7 15.492a4.47 4.47 0 0 1 6.6-.013c.438.474 1.143-.235.707-.707a5.475 5.475 0 0 0-8.015.013c-.434.474.271 1.183.707.707Z"/><circle cx="11.999" cy="17.172" r="1.12" fill="currentColor"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Youtube(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.437 19.937H4.562a2.5 2.5 0 0 1-2.5-2.5V6.563a2.5 2.5 0 0 1 2.5-2.5h14.875a2.5 2.5 0 0 1 2.5 2.5v10.874a2.5 2.5 0 0 1-2.5 2.5M4.562 5.063a1.5 1.5 0 0 0-1.5 1.5v10.874a1.5 1.5 0 0 0 1.5 1.5h14.875a1.5 1.5 0 0 0 1.5-1.5V6.563a1.5 1.5 0 0 0-1.5-1.5Z"/><path fill="currentColor" d="M14.568 11.149L10.6 8.432a1.032 1.032 0 0 0-1.614.851v5.434a1.032 1.032 0 0 0 1.614.851l3.972-2.717a1.031 1.031 0 0 0-.004-1.702"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomIn(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.279 17.039a7.926 7.926 0 0 1-5.206 1.941a7.964 7.964 0 0 1-7.96-7.96a7.964 7.964 0 0 1 7.96-7.96a7.964 7.964 0 0 1 7.96 7.96a7.93 7.93 0 0 1-2.04 5.319l.165.165l3.583 3.582c.455.456-.252 1.163-.707.708zm1.754-6.019a6.964 6.964 0 0 0-6.96-6.96a6.963 6.963 0 0 0-6.96 6.96a6.963 6.963 0 0 0 6.96 6.96a6.964 6.964 0 0 0 6.96-6.96m-7.46.5h-1.5c-.645 0-.643-1 0-1h1.5v-1.5c0-.645 1-.643 1 0v1.5h1.5c.645 0 .643 1 0 1h-1.5v1.5c0 .645-1 .643-1 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomOut(children ...ElementRenderer) *CircumIcon {
	return &CircumIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.279 17.039a7.926 7.926 0 0 1-5.206 1.941a7.964 7.964 0 0 1-7.96-7.96a7.964 7.964 0 0 1 7.96-7.96a7.964 7.964 0 0 1 7.96 7.96a7.93 7.93 0 0 1-2.04 5.319l.165.165l3.583 3.582c.455.456-.252 1.163-.707.708zm1.754-6.019a6.964 6.964 0 0 0-6.96-6.96a6.963 6.963 0 0 0-6.96 6.96a6.963 6.963 0 0 0 6.96 6.96a6.964 6.964 0 0 0 6.96-6.96m-4.96-.5c.645 0 .643 1 0 1h-4c-.645 0-.643-1 0-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
