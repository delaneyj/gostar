package mage

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = ""
	hAttr          = "1em"
	viewbox        = "0 0 24 24"
	fill           = "currentColor"
)

type MageIcon struct {
	*SVGSVGElement
}

type MageIconFn func(children ...ElementRenderer) *MageIcon

var IconLookup = map[string]MageIconFn{
	"adobe":                            Adobe,
	"aeroplane":                        Aeroplane,
	"aeroplaneFill":                    AeroplaneFill,
	"afterEffects":                     AfterEffects,
	"alarmClock":                       AlarmClock,
	"alarmClockFill":                   AlarmClockFill,
	"alignCenter":                      AlignCenter,
	"alignLeft":                        AlignLeft,
	"alignRight":                       AlignRight,
	"amazon":                           Amazon,
	"apple":                            Apple,
	"archive":                          Archive,
	"archiveDrawer":                    ArchiveDrawer,
	"archiveDrawerFill":                ArchiveDrawerFill,
	"archiveFill":                      ArchiveFill,
	"arrowDown":                        ArrowDown,
	"arrowDownCircle":                  ArrowDownCircle,
	"arrowDownCircleFill":              ArrowDownCircleFill,
	"arrowDownLeft":                    ArrowDownLeft,
	"arrowDownLeftCircle":              ArrowDownLeftCircle,
	"arrowDownLeftCircleFill":          ArrowDownLeftCircleFill,
	"arrowDownLeftSquare":              ArrowDownLeftSquare,
	"arrowDownLeftSquareFill":          ArrowDownLeftSquareFill,
	"arrowDownRight":                   ArrowDownRight,
	"arrowDownRightCircle":             ArrowDownRightCircle,
	"arrowDownRightCircleFill":         ArrowDownRightCircleFill,
	"arrowDownRightSquare":             ArrowDownRightSquare,
	"arrowDownRightSquareFill":         ArrowDownRightSquareFill,
	"arrowDownSquare":                  ArrowDownSquare,
	"arrowDownSquareFill":              ArrowDownSquareFill,
	"arrowLeft":                        ArrowLeft,
	"arrowLeftCircle":                  ArrowLeftCircle,
	"arrowLeftCircleFill":              ArrowLeftCircleFill,
	"arrowLeftSquare":                  ArrowLeftSquare,
	"arrowLeftSquareFill":              ArrowLeftSquareFill,
	"arrowRight":                       ArrowRight,
	"arrowRightCircle":                 ArrowRightCircle,
	"arrowRightCircleFill":             ArrowRightCircleFill,
	"arrowRightSquare":                 ArrowRightSquare,
	"arrowRightSquareFill":             ArrowRightSquareFill,
	"arrowUp":                          ArrowUp,
	"arrowUpCircle":                    ArrowUpCircle,
	"arrowUpCircleFill":                ArrowUpCircleFill,
	"arrowUpLeft":                      ArrowUpLeft,
	"arrowUpLeftCircle":                ArrowUpLeftCircle,
	"arrowUpLeftCircleFill":            ArrowUpLeftCircleFill,
	"arrowUpLeftSquare":                ArrowUpLeftSquare,
	"arrowUpLeftSquareFill":            ArrowUpLeftSquareFill,
	"arrowUpRight":                     ArrowUpRight,
	"arrowUpRightCircle":               ArrowUpRightCircle,
	"arrowUpRightCircleFill":           ArrowUpRightCircleFill,
	"arrowUpRightSquare":               ArrowUpRightSquare,
	"arrowUpRightSquareFill":           ArrowUpRightSquareFill,
	"arrowUpSquare":                    ArrowUpSquare,
	"arrowUpSquareFill":                ArrowUpSquareFill,
	"arrowlist":                        Arrowlist,
	"arrowsAllDirection":               ArrowsAllDirection,
	"arrowsAllDirectionTwo":            ArrowsAllDirectionTwo,
	"attachment":                       Attachment,
	"bagA":                             BagA,
	"bagAfill":                         BagAfill,
	"bagB":                             BagB,
	"bagBfill":                         BagBfill,
	"barCodeScan":                      BarCodeScan,
	"basket":                           Basket,
	"basketFill":                       BasketFill,
	"batteryCharging":                  BatteryCharging,
	"batteryChargingFill":              BatteryChargingFill,
	"batteryDead":                      BatteryDead,
	"batteryDeadFill":                  BatteryDeadFill,
	"batteryEmpty":                     BatteryEmpty,
	"batteryEmptyFill":                 BatteryEmptyFill,
	"batteryFull":                      BatteryFull,
	"batteryFullFill":                  BatteryFullFill,
	"batteryHalf":                      BatteryHalf,
	"batteryHalfFill":                  BatteryHalfFill,
	"batteryLow":                       BatteryLow,
	"batteryLowFill":                   BatteryLowFill,
	"behance":                          Behance,
	"bellNotificationSquare":           BellNotificationSquare,
	"bellNotificationSquareFill":       BellNotificationSquareFill,
	"bolt":                             Bolt,
	"boltFill":                         BoltFill,
	"book":                             Book,
	"bookFill":                         BookFill,
	"bookText":                         BookText,
	"bookTextFill":                     BookTextFill,
	"bookmark":                         Bookmark,
	"bookmarkCheck":                    BookmarkCheck,
	"bookmarkCheckFill":                BookmarkCheckFill,
	"bookmarkCross":                    BookmarkCross,
	"bookmarkCrossFill":                BookmarkCrossFill,
	"bookmarkDownload":                 BookmarkDownload,
	"bookmarkDownloadFill":             BookmarkDownloadFill,
	"bookmarkFill":                     BookmarkFill,
	"bookmarkMinus":                    BookmarkMinus,
	"bookmarkMinusFill":                BookmarkMinusFill,
	"bookmarkPlus":                     BookmarkPlus,
	"bookmarkPlusFill":                 BookmarkPlusFill,
	"bookmarkQuestionMark":             BookmarkQuestionMark,
	"bookmarkQuestionMarkFill":         BookmarkQuestionMarkFill,
	"bookmarkUpload":                   BookmarkUpload,
	"bookmarkUploadFill":               BookmarkUploadFill,
	"box":                              Box,
	"boxCheck":                         BoxCheck,
	"boxCheckFill":                     BoxCheckFill,
	"boxCross":                         BoxCross,
	"boxCrossFill":                     BoxCrossFill,
	"boxFill":                          BoxFill,
	"boxMinus":                         BoxMinus,
	"boxMinusFill":                     BoxMinusFill,
	"boxPlus":                          BoxPlus,
	"boxPlusFill":                      BoxPlusFill,
	"boxQuestionMark":                  BoxQuestionMark,
	"boxQuestionMarkFill":              BoxQuestionMarkFill,
	"boxThreeD":                        BoxThreeD,
	"boxThreeDcheck":                   BoxThreeDcheck,
	"boxThreeDcheckFill":               BoxThreeDcheckFill,
	"boxThreeDcross":                   BoxThreeDcross,
	"boxThreeDcrossFill":               BoxThreeDcrossFill,
	"boxThreeDdownload":                BoxThreeDdownload,
	"boxThreeDdownloadFill":            BoxThreeDdownloadFill,
	"boxThreeDfill":                    BoxThreeDfill,
	"boxThreeDminus":                   BoxThreeDminus,
	"boxThreeDminusFill":               BoxThreeDminusFill,
	"boxThreeDnotification":            BoxThreeDnotification,
	"boxThreeDnotificationFill":        BoxThreeDnotificationFill,
	"boxThreeDplus":                    BoxThreeDplus,
	"boxThreeDplusFill":                BoxThreeDplusFill,
	"boxThreeDquestionMark":            BoxThreeDquestionMark,
	"boxThreeDquestionMarkFill":        BoxThreeDquestionMarkFill,
	"boxThreeDscan":                    BoxThreeDscan,
	"boxThreeDscanFill":                BoxThreeDscanFill,
	"boxThreeDupload":                  BoxThreeDupload,
	"boxThreeDuploadFill":              BoxThreeDuploadFill,
	"boxUpload":                        BoxUpload,
	"boxUploadFill":                    BoxUploadFill,
	"briefcase":                        Briefcase,
	"briefcaseFill":                    BriefcaseFill,
	"broadcast":                        Broadcast,
	"broadcastFill":                    BroadcastFill,
	"buildingA":                        BuildingA,
	"buildingAfill":                    BuildingAfill,
	"buildingB":                        BuildingB,
	"buildingBfill":                    BuildingBfill,
	"buildingTree":                     BuildingTree,
	"buildingTreeFill":                 BuildingTreeFill,
	"calculator":                       Calculator,
	"calculatorFill":                   CalculatorFill,
	"calendar":                         Calendar,
	"calendarCheck":                    CalendarCheck,
	"calendarCheckFill":                CalendarCheckFill,
	"calendarCross":                    CalendarCross,
	"calendarCrossFill":                CalendarCrossFill,
	"calendarDownload":                 CalendarDownload,
	"calendarDownloadFill":             CalendarDownloadFill,
	"calendarFill":                     CalendarFill,
	"calendarMinus":                    CalendarMinus,
	"calendarMinusFill":                CalendarMinusFill,
	"calendarPlus":                     CalendarPlus,
	"calendarPlusFill":                 CalendarPlusFill,
	"calendarQuestionMark":             CalendarQuestionMark,
	"calendarQuestionMarkFill":         CalendarQuestionMarkFill,
	"calendarThree":                    CalendarThree,
	"calendarThreeFill":                CalendarThreeFill,
	"calendarTwo":                      CalendarTwo,
	"calendarTwoFill":                  CalendarTwoFill,
	"calendarUpload":                   CalendarUpload,
	"calendarUploadFill":               CalendarUploadFill,
	"camera":                           Camera,
	"cameraFill":                       CameraFill,
	"cameraTwo":                        CameraTwo,
	"cameraTwoFill":                    CameraTwoFill,
	"cancel":                           Cancel,
	"cancelFill":                       CancelFill,
	"caretDown":                        CaretDown,
	"caretDownFill":                    CaretDownFill,
	"caretLeft":                        CaretLeft,
	"caretLeftFill":                    CaretLeftFill,
	"caretRight":                       CaretRight,
	"caretRightFill":                   CaretRightFill,
	"caretUp":                          CaretUp,
	"caretUpFill":                      CaretUpFill,
	"chart":                            Chart,
	"chartB":                           ChartB,
	"chartBfill":                       ChartBfill,
	"chartDown":                        ChartDown,
	"chartDownB":                       ChartDownB,
	"chartDownBfill":                   ChartDownBfill,
	"chartDownFill":                    ChartDownFill,
	"chartFifteen":                     ChartFifteen,
	"chartFifteenFill":                 ChartFifteenFill,
	"chartFifty":                       ChartFifty,
	"chartFiftyFill":                   ChartFiftyFill,
	"chartFill":                        ChartFill,
	"chartTwentyFive":                  ChartTwentyFive,
	"chartTwentyFiveFill":              ChartTwentyFiveFill,
	"chartUp":                          ChartUp,
	"chartUpB":                         ChartUpB,
	"chartUpBfill":                     ChartUpBfill,
	"chartUpFill":                      ChartUpFill,
	"chartVertical":                    ChartVertical,
	"chartVerticalFill":                ChartVerticalFill,
	"check":                            Check,
	"checkCircle":                      CheckCircle,
	"checkCircleFill":                  CheckCircleFill,
	"checkSquare":                      CheckSquare,
	"checkSquareFill":                  CheckSquareFill,
	"checklist":                        Checklist,
	"checklistNote":                    ChecklistNote,
	"checklistNoteFill":                ChecklistNoteFill,
	"chevronDown":                      ChevronDown,
	"chevronDownCircle":                ChevronDownCircle,
	"chevronDownCircleFill":            ChevronDownCircleFill,
	"chevronDownSquare":                ChevronDownSquare,
	"chevronDownSquareFill":            ChevronDownSquareFill,
	"chevronLeft":                      ChevronLeft,
	"chevronLeftCircle":                ChevronLeftCircle,
	"chevronLeftCircleFill":            ChevronLeftCircleFill,
	"chevronLeftSquare":                ChevronLeftSquare,
	"chevronLeftSquareFill":            ChevronLeftSquareFill,
	"chevronRight":                     ChevronRight,
	"chevronRightCircle":               ChevronRightCircle,
	"chevronRightCircleFill":           ChevronRightCircleFill,
	"chevronRightSquare":               ChevronRightSquare,
	"chevronRightSquareFill":           ChevronRightSquareFill,
	"chevronUp":                        ChevronUp,
	"chevronUpCircle":                  ChevronUpCircle,
	"chevronUpCircleFill":              ChevronUpCircleFill,
	"chevronUpSquare":                  ChevronUpSquare,
	"chevronUpSquareFill":              ChevronUpSquareFill,
	"chip":                             Chip,
	"chipFill":                         ChipFill,
	"clipboard":                        Clipboard,
	"clipboardFill":                    ClipboardFill,
	"clipboardThree":                   ClipboardThree,
	"clipboardTwo":                     ClipboardTwo,
	"clipboardTwoFill":                 ClipboardTwoFill,
	"clock":                            Clock,
	"clockFill":                        ClockFill,
	"coinA":                            CoinA,
	"coinAfill":                        CoinAfill,
	"coinB":                            CoinB,
	"coinBfill":                        CoinBfill,
	"colorPicker":                      ColorPicker,
	"colorPickerFill":                  ColorPickerFill,
	"colorSwatch":                      ColorSwatch,
	"colorSwatchFill":                  ColorSwatchFill,
	"compactDisk":                      CompactDisk,
	"compactDiskFill":                  CompactDiskFill,
	"compass":                          Compass,
	"compassFill":                      CompassFill,
	"contactBook":                      ContactBook,
	"contactBookFill":                  ContactBookFill,
	"copy":                             Copy,
	"copyFill":                         CopyFill,
	"creditCard":                       CreditCard,
	"creditCardFill":                   CreditCardFill,
	"crownA":                           CrownA,
	"crownAfill":                       CrownAfill,
	"crownB":                           CrownB,
	"crownBfill":                       CrownBfill,
	"cupHot":                           CupHot,
	"cupHotFill":                       CupHotFill,
	"dashMenu":                         DashMenu,
	"dashboard":                        Dashboard,
	"dashboardBar":                     DashboardBar,
	"dashboardBarFill":                 DashboardBarFill,
	"dashboardBarNotification":         DashboardBarNotification,
	"dashboardChart":                   DashboardChart,
	"dashboardChartArrow":              DashboardChartArrow,
	"dashboardChartArrowFill":          DashboardChartArrowFill,
	"dashboardChartFill":               DashboardChartFill,
	"dashboardChartNotification":       DashboardChartNotification,
	"dashboardChartStar":               DashboardChartStar,
	"dashboardCheck":                   DashboardCheck,
	"dashboardCheckFill":               DashboardCheckFill,
	"dashboardCircleBar":               DashboardCircleBar,
	"dashboardCircleBarFill":           DashboardCircleBarFill,
	"dashboardCircleChart":             DashboardCircleChart,
	"dashboardCircleChartFill":         DashboardCircleChartFill,
	"dashboardCross":                   DashboardCross,
	"dashboardCrossFill":               DashboardCrossFill,
	"dashboardFill":                    DashboardFill,
	"dashboardFour":                    DashboardFour,
	"dashboardFourFill":                DashboardFourFill,
	"dashboardMinus":                   DashboardMinus,
	"dashboardMinusFill":               DashboardMinusFill,
	"dashboardPlus":                    DashboardPlus,
	"dashboardPlusFill":                DashboardPlusFill,
	"dashboardThree":                   DashboardThree,
	"dashboardThreeFill":               DashboardThreeFill,
	"dashboardTwo":                     DashboardTwo,
	"dashboardTwoFill":                 DashboardTwoFill,
	"database":                         Database,
	"databaseFill":                     DatabaseFill,
	"databaseTwo":                      DatabaseTwo,
	"databaseTwoFill":                  DatabaseTwoFill,
	"deliveryTruck":                    DeliveryTruck,
	"deliveryTruckFill":                DeliveryTruckFill,
	"dialerKeypad":                     DialerKeypad,
	"dialerKeypadFill":                 DialerKeypadFill,
	"directionRight":                   DirectionRight,
	"directionRightFill":               DirectionRightFill,
	"directionRightTwo":                DirectionRightTwo,
	"directionRightTwoFill":            DirectionRightTwoFill,
	"directionUp":                      DirectionUp,
	"directionUpFill":                  DirectionUpFill,
	"directionUpRight":                 DirectionUpRight,
	"directionUpRightFill":             DirectionUpRightFill,
	"directionUpRightTwo":              DirectionUpRightTwo,
	"directionUpRightTwoFill":          DirectionUpRightTwoFill,
	"directionUpTwo":                   DirectionUpTwo,
	"directionUpTwoFill":               DirectionUpTwoFill,
	"discord":                          Discord,
	"divide":                           Divide,
	"divideCircle":                     DivideCircle,
	"divideCircleFill":                 DivideCircleFill,
	"divideSquare":                     DivideSquare,
	"divideSquareFill":                 DivideSquareFill,
	"dollar":                           Dollar,
	"dollarFill":                       DollarFill,
	"dots":                             Dots,
	"dotsCircle":                       DotsCircle,
	"dotsCircleFill":                   DotsCircleFill,
	"dotsHorizontal":                   DotsHorizontal,
	"dotsHorizontalCircle":             DotsHorizontalCircle,
	"dotsHorizontalCircleFill":         DotsHorizontalCircleFill,
	"dotsHorizontalSquare":             DotsHorizontalSquare,
	"dotsHorizontalSquareFill":         DotsHorizontalSquareFill,
	"dotsMenu":                         DotsMenu,
	"dotsSquare":                       DotsSquare,
	"dotsSquareFill":                   DotsSquareFill,
	"doubleArrowCircle":                DoubleArrowCircle,
	"doubleArrowCircleFill":            DoubleArrowCircleFill,
	"doubleArrowDown":                  DoubleArrowDown,
	"doubleArrowLeft":                  DoubleArrowLeft,
	"doubleArrowRight":                 DoubleArrowRight,
	"doubleArrowUp":                    DoubleArrowUp,
	"doubleCircle":                     DoubleCircle,
	"doubleCircleFill":                 DoubleCircleFill,
	"doubleDiamond":                    DoubleDiamond,
	"doubleDiamondFill":                DoubleDiamondFill,
	"doubleSquare":                     DoubleSquare,
	"doubleSquareFill":                 DoubleSquareFill,
	"download":                         Download,
	"dribbble":                         Dribbble,
	"earth":                            Earth,
	"earthFill":                        EarthFill,
	"edit":                             Edit,
	"editFill":                         EditFill,
	"editPen":                          EditPen,
	"editPenFill":                      EditPenFill,
	"electricity":                      Electricity,
	"electricityDanger":                ElectricityDanger,
	"electricityDangerFill":            ElectricityDangerFill,
	"electricityFill":                  ElectricityFill,
	"email":                            Email,
	"emailFill":                        EmailFill,
	"emailNotification":                EmailNotification,
	"emailOpened":                      EmailOpened,
	"emailOpenedFill":                  EmailOpenedFill,
	"exchangeA":                        ExchangeA,
	"exchangeB":                        ExchangeB,
	"exclamationCircle":                ExclamationCircle,
	"exclamationCircleFill":            ExclamationCircleFill,
	"exclamationHexagon":               ExclamationHexagon,
	"exclamationHexagonFill":           ExclamationHexagonFill,
	"exclamationSquare":                ExclamationSquare,
	"exclamationSquareFill":            ExclamationSquareFill,
	"exclamationTriangle":              ExclamationTriangle,
	"exclamationTriangleFill":          ExclamationTriangleFill,
	"externalLink":                     ExternalLink,
	"eye":                              Eye,
	"eyeClosed":                        EyeClosed,
	"eyeFill":                          EyeFill,
	"eyeOff":                           EyeOff,
	"eyeOffFill":                       EyeOffFill,
	"facebook":                         Facebook,
	"facebookCircle":                   FacebookCircle,
	"facebookSquare":                   FacebookSquare,
	"fastForward":                      FastForward,
	"fastForwardBack":                  FastForwardBack,
	"fastForwardBackFill":              FastForwardBackFill,
	"fastForwardFill":                  FastForwardFill,
	"female":                           Female,
	"figma":                            Figma,
	"file":                             File,
	"fileCheck":                        FileCheck,
	"fileCheckFill":                    FileCheckFill,
	"fileCross":                        FileCross,
	"fileCrossFill":                    FileCrossFill,
	"fileDownload":                     FileDownload,
	"fileDownloadFill":                 FileDownloadFill,
	"fileFill":                         FileFill,
	"fileMinus":                        FileMinus,
	"fileMinusFill":                    FileMinusFill,
	"filePlus":                         FilePlus,
	"filePlusFill":                     FilePlusFill,
	"fileQuestionMark":                 FileQuestionMark,
	"fileQuestionMarkFill":             FileQuestionMarkFill,
	"fileRecords":                      FileRecords,
	"fileRecordsFill":                  FileRecordsFill,
	"fileThree":                        FileThree,
	"fileTwo":                          FileTwo,
	"fileTwoFill":                      FileTwoFill,
	"fileUpload":                       FileUpload,
	"fileUploadFill":                   FileUploadFill,
	"filter":                           Filter,
	"filterFill":                       FilterFill,
	"filterSquare":                     FilterSquare,
	"filterSquareFill":                 FilterSquareFill,
	"filterTwo":                        FilterTwo,
	"filterTwoFill":                    FilterTwoFill,
	"fingerprint":                      Fingerprint,
	"fingerprintMinimal":               FingerprintMinimal,
	"fireA":                            FireA,
	"fireAfill":                        FireAfill,
	"fireB":                            FireB,
	"fireBfill":                        FireBfill,
	"firstAidKit":                      FirstAidKit,
	"firstAidKitFill":                  FirstAidKitFill,
	"flag":                             Flag,
	"flagFill":                         FlagFill,
	"focus":                            Focus,
	"focusFill":                        FocusFill,
	"folder":                           Folder,
	"folderCheck":                      FolderCheck,
	"folderCheckFill":                  FolderCheckFill,
	"folderCross":                      FolderCross,
	"folderCrossFill":                  FolderCrossFill,
	"folderFill":                       FolderFill,
	"folderMinus":                      FolderMinus,
	"folderMinusFill":                  FolderMinusFill,
	"folderOpen":                       FolderOpen,
	"folderOpenFill":                   FolderOpenFill,
	"folderPlus":                       FolderPlus,
	"folderPlusFill":                   FolderPlusFill,
	"folderTwo":                        FolderTwo,
	"folderTwoFill":                    FolderTwoFill,
	"gameboy":                          Gameboy,
	"gameboyFill":                      GameboyFill,
	"gemA":                             GemA,
	"gemAfill":                         GemAfill,
	"gemB":                             GemB,
	"gemBfill":                         GemBfill,
	"gemC":                             GemC,
	"gemCfill":                         GemCfill,
	"gemStone":                         GemStone,
	"gemStoneFill":                     GemStoneFill,
	"gif":                              Gif,
	"gifFill":                          GifFill,
	"gift":                             Gift,
	"giftFill":                         GiftFill,
	"github":                           Github,
	"globe":                            Globe,
	"globeFill":                        GlobeFill,
	"goals":                            Goals,
	"goalsFill":                        GoalsFill,
	"google":                           Google,
	"handicapped":                      Handicapped,
	"handicappedFill":                  HandicappedFill,
	"hash":                             Hash,
	"headphoneMute":                    HeadphoneMute,
	"headphoneMuteFill":                HeadphoneMuteFill,
	"healthCircle":                     HealthCircle,
	"healthCircleFill":                 HealthCircleFill,
	"healthSquare":                     HealthSquare,
	"healthSquareFill":                 HealthSquareFill,
	"heaphone":                         Heaphone,
	"heaphoneFill":                     HeaphoneFill,
	"heart":                            Heart,
	"heartFill":                        HeartFill,
	"heartHealth":                      HeartHealth,
	"heartHealthFill":                  HeartHealthFill,
	"heartSquare":                      HeartSquare,
	"heartSquareFill":                  HeartSquareFill,
	"home":                             Home,
	"homeCheck":                        HomeCheck,
	"homeCheckFill":                    HomeCheckFill,
	"homeCross":                        HomeCross,
	"homeCrossFill":                    HomeCrossFill,
	"homeFill":                         HomeFill,
	"homeFour":                         HomeFour,
	"homeFourFill":                     HomeFourFill,
	"homeHeart":                        HomeHeart,
	"homeHeartFill":                    HomeHeartFill,
	"homePlus":                         HomePlus,
	"homePlusFill":                     HomePlusFill,
	"homeSecurityLock":                 HomeSecurityLock,
	"homeSecurityLockFill":             HomeSecurityLockFill,
	"homeThree":                        HomeThree,
	"homeThreeFill":                    HomeThreeFill,
	"homeTwo":                          HomeTwo,
	"homeTwoFill":                      HomeTwoFill,
	"hospitalCircle":                   HospitalCircle,
	"hospitalCircleFill":               HospitalCircleFill,
	"hospitalPlus":                     HospitalPlus,
	"hospitalPlusFill":                 HospitalPlusFill,
	"hospitalShield":                   HospitalShield,
	"hospitalShieldFill":               HospitalShieldFill,
	"hospitalSquare":                   HospitalSquare,
	"hospitalSquareFill":               HospitalSquareFill,
	"hourGlass":                        HourGlass,
	"hourGlassFill":                    HourGlassFill,
	"idCard":                           IdCard,
	"idCardFill":                       IdCardFill,
	"illustrator":                      Illustrator,
	"image":                            Image,
	"imageCheck":                       ImageCheck,
	"imageCross":                       ImageCross,
	"imageDownload":                    ImageDownload,
	"imageFill":                        ImageFill,
	"imageMinus":                       ImageMinus,
	"imagePlus":                        ImagePlus,
	"imageQuestionMark":                ImageQuestionMark,
	"imageUpload":                      ImageUpload,
	"inbox":                            Inbox,
	"inboxCheck":                       InboxCheck,
	"inboxCheckFill":                   InboxCheckFill,
	"inboxCross":                       InboxCross,
	"inboxCrossFill":                   InboxCrossFill,
	"inboxDownload":                    InboxDownload,
	"inboxDownloadFill":                InboxDownloadFill,
	"inboxFill":                        InboxFill,
	"inboxMinus":                       InboxMinus,
	"inboxMinusFill":                   InboxMinusFill,
	"inboxNotification":                InboxNotification,
	"inboxNotificationFill":            InboxNotificationFill,
	"inboxPlus":                        InboxPlus,
	"inboxPlusFill":                    InboxPlusFill,
	"inboxQuestionMark":                InboxQuestionMark,
	"inboxQuestionMarkFill":            InboxQuestionMarkFill,
	"inboxStar":                        InboxStar,
	"inboxStarFill":                    InboxStarFill,
	"inboxUpload":                      InboxUpload,
	"inboxUploadFill":                  InboxUploadFill,
	"informationCircle":                InformationCircle,
	"informationCircleFill":            InformationCircleFill,
	"informationSquare":                InformationSquare,
	"informationSquareFill":            InformationSquareFill,
	"instagramCircle":                  InstagramCircle,
	"instagramSquare":                  InstagramSquare,
	"key":                              Key,
	"keyFill":                          KeyFill,
	"keyboard":                         Keyboard,
	"keyboardFill":                     KeyboardFill,
	"larrowDownLeft":                   LarrowDownLeft,
	"larrowDownRight":                  LarrowDownRight,
	"larrowLeftDown":                   LarrowLeftDown,
	"larrowLeftUp":                     LarrowLeftUp,
	"larrowRightDown":                  LarrowRightDown,
	"larrowRightUp":                    LarrowRightUp,
	"larrowUpLeft":                     LarrowUpLeft,
	"larrowUpRight":                    LarrowUpRight,
	"laptop":                           Laptop,
	"laptopFill":                       LaptopFill,
	"layoutCenter":                     LayoutCenter,
	"layoutCenterFill":                 LayoutCenterFill,
	"layoutDown":                       LayoutDown,
	"layoutDownFill":                   LayoutDownFill,
	"layoutGrid":                       LayoutGrid,
	"layoutGridFill":                   LayoutGridFill,
	"layoutLeft":                       LayoutLeft,
	"layoutLeftFill":                   LayoutLeftFill,
	"layoutRight":                      LayoutRight,
	"layoutRightFill":                  LayoutRightFill,
	"layoutUp":                         LayoutUp,
	"layoutUpFill":                     LayoutUpFill,
	"layoutUpLeft":                     LayoutUpLeft,
	"layoutUpLeftFill":                 LayoutUpLeftFill,
	"layoutUpRight":                    LayoutUpRight,
	"layoutUpRightFill":                LayoutUpRightFill,
	"lens":                             Lens,
	"lensFill":                         LensFill,
	"lightBulb":                        LightBulb,
	"lightBulbElectricity":             LightBulbElectricity,
	"lightBulbElectricityFill":         LightBulbElectricityFill,
	"lightBulbFill":                    LightBulbFill,
	"lightBulbOff":                     LightBulbOff,
	"lightBulbOffFill":                 LightBulbOffFill,
	"line":                             Line,
	"link":                             Link,
	"linkedin":                         Linkedin,
	"location":                         Location,
	"locationFill":                     LocationFill,
	"locationPin":                      LocationPin,
	"lock":                             Lock,
	"lockFill":                         LockFill,
	"lockSquare":                       LockSquare,
	"lockSquareFill":                   LockSquareFill,
	"login":                            Login,
	"logout":                           Logout,
	"magnetDown":                       MagnetDown,
	"magnetDownFill":                   MagnetDownFill,
	"magnetLeft":                       MagnetLeft,
	"magnetLeftFill":                   MagnetLeftFill,
	"magnetRight":                      MagnetRight,
	"magnetRightFill":                  MagnetRightFill,
	"magnetUp":                         MagnetUp,
	"magnetUpFill":                     MagnetUpFill,
	"male":                             Male,
	"mapMarker":                        MapMarker,
	"mapMarkerFill":                    MapMarkerFill,
	"maximize":                         Maximize,
	"mediaReelH":                       MediaReelH,
	"mediaReelHfill":                   MediaReelHfill,
	"mediaReelV":                       MediaReelV,
	"mediaReelVfill":                   MediaReelVfill,
	"medium":                           Medium,
	"megaphoneA":                       MegaphoneA,
	"megaphoneAfill":                   MegaphoneAfill,
	"megaphoneB":                       MegaphoneB,
	"megaphoneBfill":                   MegaphoneBfill,
	"memoryCard":                       MemoryCard,
	"memoryCardFill":                   MemoryCardFill,
	"message":                          Message,
	"messageCheck":                     MessageCheck,
	"messageCheckFill":                 MessageCheckFill,
	"messageCheckRound":                MessageCheckRound,
	"messageCheckRoundFill":            MessageCheckRoundFill,
	"messageConversation":              MessageConversation,
	"messageConversationFill":          MessageConversationFill,
	"messageDots":                      MessageDots,
	"messageDotsCheck":                 MessageDotsCheck,
	"messageDotsCross":                 MessageDotsCross,
	"messageDotsDownload":              MessageDotsDownload,
	"messageDotsFill":                  MessageDotsFill,
	"messageDotsMinus":                 MessageDotsMinus,
	"messageDotsPlus":                  MessageDotsPlus,
	"messageDotsQuestionMark":          MessageDotsQuestionMark,
	"messageDotsRound":                 MessageDotsRound,
	"messageDotsRoundCheck":            MessageDotsRoundCheck,
	"messageDotsRoundCross":            MessageDotsRoundCross,
	"messageDotsRoundDownload":         MessageDotsRoundDownload,
	"messageDotsRoundFill":             MessageDotsRoundFill,
	"messageDotsRoundMinus":            MessageDotsRoundMinus,
	"messageDotsRoundPlus":             MessageDotsRoundPlus,
	"messageDotsRoundQuestionMark":     MessageDotsRoundQuestionMark,
	"messageDotsRoundUpload":           MessageDotsRoundUpload,
	"messageDotsUpload":                MessageDotsUpload,
	"messageFill":                      MessageFill,
	"messageInfoRound":                 MessageInfoRound,
	"messageInfoRoundFill":             MessageInfoRoundFill,
	"messageInformation":               MessageInformation,
	"messageInformationFill":           MessageInformationFill,
	"messageMinus":                     MessageMinus,
	"messageMinusFill":                 MessageMinusFill,
	"messageMinusRound":                MessageMinusRound,
	"messageMinusRoundFill":            MessageMinusRoundFill,
	"messagePlus":                      MessagePlus,
	"messagePlusFill":                  MessagePlusFill,
	"messagePlusRound":                 MessagePlusRound,
	"messagePlusRoundFill":             MessagePlusRoundFill,
	"messageQuestionMark":              MessageQuestionMark,
	"messageQuestionMarkFill":          MessageQuestionMarkFill,
	"messageQuestionMarkRound":         MessageQuestionMarkRound,
	"messageQuestionMarkRoundFill":     MessageQuestionMarkRoundFill,
	"messageRound":                     MessageRound,
	"messageRoundFill":                 MessageRoundFill,
	"messageRoundSquare":               MessageRoundSquare,
	"messageRoundSquareFill":           MessageRoundSquareFill,
	"messageSquare":                    MessageSquare,
	"messageSquareFill":                MessageSquareFill,
	"messenger":                        Messenger,
	"meta":                             Meta,
	"microphone":                       Microphone,
	"microphoneFill":                   MicrophoneFill,
	"microphoneMute":                   MicrophoneMute,
	"microphoneMuteFill":               MicrophoneMuteFill,
	"microsoftWindows":                 MicrosoftWindows,
	"minimize":                         Minimize,
	"minus":                            Minus,
	"minusCircle":                      MinusCircle,
	"minusCircleFill":                  MinusCircleFill,
	"minusSquare":                      MinusSquare,
	"minusSquareFill":                  MinusSquareFill,
	"mobilePhone":                      MobilePhone,
	"mobilePhoneFill":                  MobilePhoneFill,
	"moneyExchange":                    MoneyExchange,
	"moon":                             Moon,
	"moonFill":                         MoonFill,
	"mouse":                            Mouse,
	"mouseFill":                        MouseFill,
	"mousePointer":                     MousePointer,
	"mousePointerFill":                 MousePointerFill,
	"mouseTwo":                         MouseTwo,
	"mouseTwoFill":                     MouseTwoFill,
	"multiply":                         Multiply,
	"multiplyCircle":                   MultiplyCircle,
	"multiplyCircleFill":               MultiplyCircleFill,
	"multiplySquare":                   MultiplySquare,
	"multiplySquareFill":               MultiplySquareFill,
	"music":                            Music,
	"musicAlternate":                   MusicAlternate,
	"musicAlternateFill":               MusicAlternateFill,
	"musicCircle":                      MusicCircle,
	"musicCircleFill":                  MusicCircleFill,
	"musicFill":                        MusicFill,
	"musicSquare":                      MusicSquare,
	"musicSquareFill":                  MusicSquareFill,
	"netflix":                          Netflix,
	"next":                             Next,
	"nextFill":                         NextFill,
	"note":                             Note,
	"noteFill":                         NoteFill,
	"noteText":                         NoteText,
	"noteTextFill":                     NoteTextFill,
	"noteWithText":                     NoteWithText,
	"noteWithTextFill":                 NoteWithTextFill,
	"notificationBell":                 NotificationBell,
	"notificationBellCheck":            NotificationBellCheck,
	"notificationBellCheckFill":        NotificationBellCheckFill,
	"notificationBellCross":            NotificationBellCross,
	"notificationBellCrossFill":        NotificationBellCrossFill,
	"notificationBellDownload":         NotificationBellDownload,
	"notificationBellDownloadFill":     NotificationBellDownloadFill,
	"notificationBellFill":             NotificationBellFill,
	"notificationBellMinus":            NotificationBellMinus,
	"notificationBellMinusFill":        NotificationBellMinusFill,
	"notificationBellMuted":            NotificationBellMuted,
	"notificationBellMutedFill":        NotificationBellMutedFill,
	"notificationBellPending":          NotificationBellPending,
	"notificationBellPendingFill":      NotificationBellPendingFill,
	"notificationBellPlus":             NotificationBellPlus,
	"notificationBellPlusFill":         NotificationBellPlusFill,
	"notificationBellQuestionMark":     NotificationBellQuestionMark,
	"notificationBellQuestionMarkFill": NotificationBellQuestionMarkFill,
	"notificationBellSnooze":           NotificationBellSnooze,
	"notificationBellSnoozeTwo":        NotificationBellSnoozeTwo,
	"notificationBellSnoozeTwoFill":    NotificationBellSnoozeTwoFill,
	"notificationBellUpload":           NotificationBellUpload,
	"notificationBellUploadFill":       NotificationBellUploadFill,
	"notion":                           Notion,
	"packageBox":                       PackageBox,
	"packageBoxFill":                   PackageBoxFill,
	"pause":                            Pause,
	"pauseFill":                        PauseFill,
	"pauseSquare":                      PauseSquare,
	"pauseSquareFill":                  PauseSquareFill,
	"paypal":                           Paypal,
	"pen":                              Pen,
	"penFill":                          PenFill,
	"phone":                            Phone,
	"phoneCall":                        PhoneCall,
	"phoneCallFill":                    PhoneCallFill,
	"phoneCancel":                      PhoneCancel,
	"phoneCancelFill":                  PhoneCancelFill,
	"phoneCross":                       PhoneCross,
	"phoneCrossFill":                   PhoneCrossFill,
	"phoneFill":                        PhoneFill,
	"phoneIncoming":                    PhoneIncoming,
	"phoneIncomingFill":                PhoneIncomingFill,
	"phoneMinus":                       PhoneMinus,
	"phoneMinusFill":                   PhoneMinusFill,
	"phoneMissedCall":                  PhoneMissedCall,
	"phoneMissedCallFill":              PhoneMissedCallFill,
	"phoneOutgoing":                    PhoneOutgoing,
	"phoneOutgoingFill":                PhoneOutgoingFill,
	"phonePlus":                        PhonePlus,
	"phonePlusFill":                    PhonePlusFill,
	"phoneRingingLoud":                 PhoneRingingLoud,
	"phoneRingingLoudFill":             PhoneRingingLoudFill,
	"photoshop":                        Photoshop,
	"pin":                              Pin,
	"pinFill":                          PinFill,
	"pinterest":                        Pinterest,
	"play":                             Play,
	"playCircle":                       PlayCircle,
	"playCircleFill":                   PlayCircleFill,
	"playFill":                         PlayFill,
	"playSquare":                       PlaySquare,
	"playSquareFill":                   PlaySquareFill,
	"playlist":                         Playlist,
	"playlistAdd":                      PlaylistAdd,
	"playlistAlternate":                PlaylistAlternate,
	"playlistAlternateFill":            PlaylistAlternateFill,
	"playlistFill":                     PlaylistFill,
	"playstore":                        Playstore,
	"plus":                             Plus,
	"plusCircle":                       PlusCircle,
	"plusCircleFill":                   PlusCircleFill,
	"plusSquare":                       PlusSquare,
	"plusSquareFill":                   PlusSquareFill,
	"premierPro":                       PremierPro,
	"preview":                          Preview,
	"previewCircle":                    PreviewCircle,
	"previewCircleFill":                PreviewCircleFill,
	"previewFill":                      PreviewFill,
	"previous":                         Previous,
	"previousFill":                     PreviousFill,
	"printer":                          Printer,
	"printerFill":                      PrinterFill,
	"qrCode":                           QrCode,
	"qrCodeFill":                       QrCodeFill,
	"questionMarkCircle":               QuestionMarkCircle,
	"questionMarkCircleFill":           QuestionMarkCircleFill,
	"questionMarkSquare":               QuestionMarkSquare,
	"questionMarkSquareFill":           QuestionMarkSquareFill,
	"reddit":                           Reddit,
	"refresh":                          Refresh,
	"refreshReverse":                   RefreshReverse,
	"reload":                           Reload,
	"reloadReverse":                    ReloadReverse,
	"ribbon":                           Ribbon,
	"ribbonFill":                       RibbonFill,
	"robot":                            Robot,
	"robotAppreciate":                  RobotAppreciate,
	"robotAppreciateFill":              RobotAppreciateFill,
	"robotDead":                        RobotDead,
	"robotDeadFill":                    RobotDeadFill,
	"robotFill":                        RobotFill,
	"robotHappy":                       RobotHappy,
	"robotHappyFill":                   RobotHappyFill,
	"robotSad":                         RobotSad,
	"robotSadFill":                     RobotSadFill,
	"robotScreen":                      RobotScreen,
	"robotScreenFill":                  RobotScreenFill,
	"robotUwu":                         RobotUwu,
	"robotUwuFill":                     RobotUwuFill,
	"robotWink":                        RobotWink,
	"robotWinkFill":                    RobotWinkFill,
	"rocket":                           Rocket,
	"rocketFill":                       RocketFill,
	"roundSticker":                     RoundSticker,
	"roundStickerFill":                 RoundStickerFill,
	"saveFloppy":                       SaveFloppy,
	"saveFloppyFill":                   SaveFloppyFill,
	"scaleDown":                        ScaleDown,
	"scaleUp":                          ScaleUp,
	"scan":                             Scan,
	"scanUser":                         ScanUser,
	"scanUserFill":                     ScanUserFill,
	"screencast":                       Screencast,
	"screencastFill":                   ScreencastFill,
	"search":                           Search,
	"searchFill":                       SearchFill,
	"searchSquare":                     SearchSquare,
	"searchSquareFill":                 SearchSquareFill,
	"securityShield":                   SecurityShield,
	"securityShieldFill":               SecurityShieldFill,
	"selectBox":                        SelectBox,
	"server":                           Server,
	"serverFill":                       ServerFill,
	"serverTwo":                        ServerTwo,
	"serverTwoFill":                    ServerTwoFill,
	"settings":                         Settings,
	"settingsFill":                     SettingsFill,
	"share":                            Share,
	"shareFill":                        ShareFill,
	"shieldCheck":                      ShieldCheck,
	"shieldCheckFill":                  ShieldCheckFill,
	"shieldCross":                      ShieldCross,
	"shieldCrossFill":                  ShieldCrossFill,
	"shieldPlus":                       ShieldPlus,
	"shieldPlusFill":                   ShieldPlusFill,
	"shieldQuestionMark":               ShieldQuestionMark,
	"shieldQuestionMarkFill":           ShieldQuestionMarkFill,
	"shiledMinus":                      ShiledMinus,
	"shiledMinusFill":                  ShiledMinusFill,
	"shop":                             Shop,
	"shopFill":                         ShopFill,
	"shoppingBag":                      ShoppingBag,
	"shoppingBagFill":                  ShoppingBagFill,
	"shoppingCart":                     ShoppingCart,
	"shoppingCartFill":                 ShoppingCartFill,
	"shutDown":                         ShutDown,
	"shutDownFill":                     ShutDownFill,
	"simCard":                          SimCard,
	"simCardFill":                      SimCardFill,
	"slack":                            Slack,
	"snapchat":                         Snapchat,
	"soundWaves":                       SoundWaves,
	"spotify":                          Spotify,
	"stack":                            Stack,
	"stackFill":                        StackFill,
	"star":                             Star,
	"starCircle":                       StarCircle,
	"starCircleFill":                   StarCircleFill,
	"starFill":                         StarFill,
	"starMoving":                       StarMoving,
	"starMovingFill":                   StarMovingFill,
	"starSquare":                       StarSquare,
	"starSquareFill":                   StarSquareFill,
	"starsA":                           StarsA,
	"starsAfill":                       StarsAfill,
	"starsB":                           StarsB,
	"starsBfill":                       StarsBfill,
	"starsC":                           StarsC,
	"starsCfill":                       StarsCfill,
	"steam":                            Steam,
	"stop":                             Stop,
	"stopCircle":                       StopCircle,
	"stopCircleFill":                   StopCircleFill,
	"stopFill":                         StopFill,
	"stopSquare":                       StopSquare,
	"stopSquareFill":                   StopSquareFill,
	"stripe":                           Stripe,
	"sun":                              Sun,
	"sunFill":                          SunFill,
	"swimRingFill":                     SwimRingFill,
	"tablet":                           Tablet,
	"tabletFill":                       TabletFill,
	"tag":                              Tag,
	"tagCheck":                         TagCheck,
	"tagCheckFill":                     TagCheckFill,
	"tagCross":                         TagCross,
	"tagCrossFill":                     TagCrossFill,
	"tagFill":                          TagFill,
	"tagMinus":                         TagMinus,
	"tagMinusFill":                     TagMinusFill,
	"tagPlus":                          TagPlus,
	"tagPlusFill":                      TagPlusFill,
	"tagQuestionMark":                  TagQuestionMark,
	"tagQuestionMarkFill":              TagQuestionMarkFill,
	"tagTwo":                           TagTwo,
	"tagTwoFill":                       TagTwoFill,
	"telegram":                         Telegram,
	"television":                       Television,
	"televisionCheck":                  TelevisionCheck,
	"televisionCheckFill":              TelevisionCheckFill,
	"televisionCross":                  TelevisionCross,
	"televisionCrossFill":              TelevisionCrossFill,
	"televisionDownload":               TelevisionDownload,
	"televisionDownloadFill":           TelevisionDownloadFill,
	"televisionFill":                   TelevisionFill,
	"televisionMinus":                  TelevisionMinus,
	"televisionMinusFill":              TelevisionMinusFill,
	"televisionPlus":                   TelevisionPlus,
	"televisionPlusFill":               TelevisionPlusFill,
	"televisionQuestionMark":           TelevisionQuestionMark,
	"televisionQuestionMarkFill":       TelevisionQuestionMarkFill,
	"televisionUpload":                 TelevisionUpload,
	"televisionUploadFill":             TelevisionUploadFill,
	"threads":                          Threads,
	"threadsSquare":                    ThreadsSquare,
	"threeDboxSquare":                  ThreeDboxSquare,
	"threeDboxSquareFill":              ThreeDboxSquareFill,
	"thumbsDown":                       ThumbsDown,
	"thumbsDownFill":                   ThumbsDownFill,
	"thumbsUp":                         ThumbsUp,
	"thumbsUpFill":                     ThumbsUpFill,
	"tiktok":                           Tiktok,
	"tiktokCircle":                     TiktokCircle,
	"trash":                            Trash,
	"trashFill":                        TrashFill,
	"trashSquare":                      TrashSquare,
	"trashSquareFill":                  TrashSquareFill,
	"trashThree":                       TrashThree,
	"trashThreeFill":                   TrashThreeFill,
	"trashTwo":                         TrashTwo,
	"trashTwoFill":                     TrashTwoFill,
	"trophy":                           Trophy,
	"trophyCircle":                     TrophyCircle,
	"trophyCircleFill":                 TrophyCircleFill,
	"trophyDown":                       TrophyDown,
	"trophyDownFill":                   TrophyDownFill,
	"trophyFill":                       TrophyFill,
	"trophyStar":                       TrophyStar,
	"trophyStarFill":                   TrophyStarFill,
	"trophyUp":                         TrophyUp,
	"trophyUpFill":                     TrophyUpFill,
	"tube":                             Tube,
	"tubeFill":                         TubeFill,
	"twitter":                          Twitter,
	"twtich":                           Twtich,
	"unlocked":                         Unlocked,
	"unlockedFill":                     UnlockedFill,
	"upload":                           Upload,
	"user":                             User,
	"userCheck":                        UserCheck,
	"userCheckFill":                    UserCheckFill,
	"userCircle":                       UserCircle,
	"userCircleFill":                   UserCircleFill,
	"userCross":                        UserCross,
	"userCrossFill":                    UserCrossFill,
	"userFill":                         UserFill,
	"userMinus":                        UserMinus,
	"userMinusFill":                    UserMinusFill,
	"userPlus":                         UserPlus,
	"userPlusFill":                     UserPlusFill,
	"userQuestionMark":                 UserQuestionMark,
	"userQuestionMarkFill":             UserQuestionMarkFill,
	"userSquare":                       UserSquare,
	"userSquareFill":                   UserSquareFill,
	"users":                            Users,
	"usersFill":                        UsersFill,
	"verifiedCheck":                    VerifiedCheck,
	"verifiedCheckFill":                VerifiedCheckFill,
	"video":                            Video,
	"videoCheck":                       VideoCheck,
	"videoCheckFill":                   VideoCheckFill,
	"videoCross":                       VideoCross,
	"videoCrossFill":                   VideoCrossFill,
	"videoDownload":                    VideoDownload,
	"videoDownloadFill":                VideoDownloadFill,
	"videoFill":                        VideoFill,
	"videoMinus":                       VideoMinus,
	"videoMinusFill":                   VideoMinusFill,
	"videoPlayer":                      VideoPlayer,
	"videoPlayerFill":                  VideoPlayerFill,
	"videoPlus":                        VideoPlus,
	"videoPlusFill":                    VideoPlusFill,
	"videoQuestionMark":                VideoQuestionMark,
	"videoQuestionMarkFill":            VideoQuestionMarkFill,
	"videoUpload":                      VideoUpload,
	"videoUploadFill":                  VideoUploadFill,
	"visa":                             Visa,
	"visaSquare":                       VisaSquare,
	"voicemail":                        Voicemail,
	"voicemailFill":                    VoicemailFill,
	"volumeDown":                       VolumeDown,
	"volumeDownFill":                   VolumeDownFill,
	"volumeMute":                       VolumeMute,
	"volumeMuteFill":                   VolumeMuteFill,
	"volumeUp":                         VolumeUp,
	"volumeUpFill":                     VolumeUpFill,
	"volumeZero":                       VolumeZero,
	"volumeZeroFill":                   VolumeZeroFill,
	"waterDrop":                        WaterDrop,
	"waterDropFill":                    WaterDropFill,
	"waterGlass":                       WaterGlass,
	"waterGlassFill":                   WaterGlassFill,
	"weChat":                           WeChat,
	"whatsapp":                         Whatsapp,
	"whatsappFilled":                   WhatsappFilled,
	"wifi":                             Wifi,
	"wrench":                           Wrench,
	"wrenchFill":                       WrenchFill,
	"x":                                X,
	"xsquare":                          Xsquare,
	"youtube":                          Youtube,
	"zap":                              Zap,
	"zapCircle":                        ZapCircle,
	"zapCircleFill":                    ZapCircleFill,
	"zapFill":                          ZapFill,
	"zapSquare":                        ZapSquare,
	"zapSquareFill":                    ZapSquareFill,
	"zoomIn":                           ZoomIn,
	"zoomInFill":                       ZoomInFill,
	"zoomOut":                          ZoomOut,
	"zoomOutFill":                      ZoomOutFill,
}

func Adobe(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.443 3.104L2 20.897V3.104zm2.564 6.566l4.715 11.227h-3.084l-1.398-3.55H8.744zM22 3.104v17.793L14.653 3.104z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Aeroplane(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.3" d="M13.235 3.257a1.747 1.747 0 0 1 .508 1.24v5.388l5.917 2.541a.468.468 0 0 1 .274.295l.519 1.545a.508.508 0 0 1-.6.65l-6.1-1.534l-.294 5.245l1.169.875v1.748a338.247 338.247 0 0 1-2.633-.708s-1.073.295-2.623.708v-1.748l1.17-.875l-.296-5.245l-6.099 1.535a.509.509 0 0 1-.6-.65l.519-1.546a.468.468 0 0 1 .274-.295l5.917-2.541V4.497a1.749 1.749 0 0 1 .508-1.24a1.759 1.759 0 0 1 2.47 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AeroplaneFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.72 14.621a.799.799 0 0 1-.19.35a.72.72 0 0 1-.33.21a.78.78 0 0 1-.4 0l-5.81-1.46l-.27 4.8l1.06.8a.219.219 0 0 1 .1.2v1.75a.269.269 0 0 1-.09.2a.3.3 0 0 1-.16 0h-.06l-2.57-.69l-2.56.69a.25.25 0 0 1-.31-.24v-1.75a.25.25 0 0 1 .1-.2l1.06-.8l-.27-4.8l-5.81 1.46a.731.731 0 0 1-.39 0a.75.75 0 0 1-.34-.21a.73.73 0 0 1-.18-.35a.65.65 0 0 1 0-.39l.52-1.55a.61.61 0 0 1 .15-.27a.801.801 0 0 1 .27-.18l5.75-2.47v-5.23a2.001 2.001 0 0 1 .58-1.41a2.06 2.06 0 0 1 2.83 0a1.999 1.999 0 0 1 .58 1.42v5.22l5.76 2.47a.739.739 0 0 1 .42.45l.52 1.54a.913.913 0 0 1 .04.44"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AfterEffects(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.84 12.4c0 .09 0 .09-.07.09H8c.33-1 .68-2.06 1-3.1c0 .08 0 .16.07.24c.07.08.09.33.13.49l.18.56c0 .19.11.37.17.56l.19.57z"/><path fill="currentColor" d="M17.08 11.77a1 1 0 0 0-.76-.67a1.697 1.697 0 0 0-.55 0a.929.929 0 0 0-.64.47a1.67 1.67 0 0 0-.2.58H17c.1 0 .13 0 .12-.14a1.098 1.098 0 0 0-.04-.24m0 0a1 1 0 0 0-.76-.67a1.697 1.697 0 0 0-.55 0a.929.929 0 0 0-.64.47a1.67 1.67 0 0 0-.2.58H17c.1 0 .13 0 .12-.14a1.098 1.098 0 0 0-.04-.24m0 0a1 1 0 0 0-.76-.67a1.697 1.697 0 0 0-.55 0a.929.929 0 0 0-.64.47a1.67 1.67 0 0 0-.2.58H17c.1 0 .13 0 .12-.14a1.098 1.098 0 0 0-.04-.24m0 0a1 1 0 0 0-.76-.67a1.697 1.697 0 0 0-.55 0a.929.929 0 0 0-.64.47a1.67 1.67 0 0 0-.2.58H17c.1 0 .13 0 .12-.14a1.098 1.098 0 0 0-.04-.24M17 2H7a5 5 0 0 0-5 5v10a5 5 0 0 0 5 5h10a5 5 0 0 0 5-5V7a5 5 0 0 0-5-5m-4.27 14h-1.62a.141.141 0 0 1-.142-.068a.137.137 0 0 1-.018-.052c-.08-.25-.17-.49-.25-.74c-.08-.25-.2-.57-.29-.86c0-.16-.13-.18-.27-.18H7.62c-.11 0-.11 0-.15.12c-.04.12-.1.33-.16.5l-.16.52l-.21.61a.24.24 0 0 1 0 .08c0 .13-.16.11-.26.11H5.39c-.15 0-.2-.07-.14-.21c.06-.14.11-.3.16-.44l.21-.64l.39-1.1c.11-.3.21-.61.31-.92c.1-.31.19-.53.29-.8l.21-.63l.37-1l.3-.88c.1-.28.2-.56.29-.85a1.84 1.84 0 0 0 0-.41c0-.21 0-.21.21-.21H9.8c.16 0 .15 0 .2.15c.05.15.18.49.26.73c.08.24.15.42.22.62l.24.65l.24.7c.1.27.2.53.29.8l.26.72c.08.24.17.48.26.73s.12.37.19.55c.07.18.14.4.22.6c.08.2.13.38.2.57l.29.81l.12.35c.04-.01.01.07-.06.07m6-3c.005.09.005.18 0 .27c0 .09-.15.12-.24.12h-3.53c.036.328.18.634.41.87a3 3 0 0 0 2.8.18c.08 0 .12 0 .12.08v1.13c0 .18-.07.2-.22.26l-.36.11a2.171 2.171 0 0 1-.35.07c-.25 0-.5.07-.75.08c-.25.01-.34 0-.51 0H16c-.2 0-.41 0-.6-.07a2.69 2.69 0 0 1-1.44-.8a2.64 2.64 0 0 1-.56-.88a2.77 2.77 0 0 1-.2-.75a8.61 8.61 0 0 1 0-.88c.015-.361.083-.718.2-1.06a3.18 3.18 0 0 1 .41-.8a2.53 2.53 0 0 1 .9-.82c.156-.09.32-.168.49-.23a5 5 0 0 1 .56-.1c.17-.01.34-.01.51 0c.25.004.498.037.74.1c.22.062.431.149.63.26a2 2 0 0 1 .74.74c.1.18.18.36.26.54c.043.127.076.258.1.39a3 3 0 0 1 .03 1.19zm-2.45-1.85a1.697 1.697 0 0 0-.55 0a.929.929 0 0 0-.64.47a1.67 1.67 0 0 0-.2.58H17c.1 0 .13 0 .12-.14a1.1 1.1 0 0 0-.06-.26a1 1 0 0 0-.74-.7zm.76.67a1 1 0 0 0-.76-.67a1.697 1.697 0 0 0-.55 0a.929.929 0 0 0-.64.47a1.67 1.67 0 0 0-.2.58H17c.1 0 .13 0 .12-.14c0-.098-.014-.196-.04-.29zm0 0a1 1 0 0 0-.76-.67a1.697 1.697 0 0 0-.55 0a.929.929 0 0 0-.64.47a1.67 1.67 0 0 0-.2.58H17c.1 0 .13 0 .12-.14c0-.098-.014-.196-.04-.29zm0 0a1 1 0 0 0-.76-.67a1.697 1.697 0 0 0-.55 0a.929.929 0 0 0-.64.47a1.67 1.67 0 0 0-.2.58H17c.1 0 .13 0 .12-.14c0-.098-.014-.196-.04-.29zm0 0a1 1 0 0 0-.76-.67a1.697 1.697 0 0 0-.55 0a.929.929 0 0 0-.64.47a1.67 1.67 0 0 0-.2.58H17c.1 0 .13 0 .12-.14c0-.098-.014-.196-.04-.29z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlarmClock(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10" stroke-width="1.5"><path d="M14.319 15.203L12 12.883V8.247M20.116 21l-2.319-2.319m-11.594 0L3.883 21M18.377 3l2.319 2.319M5.623 3L3.304 5.319"/><path d="M12 21a8.116 8.116 0 1 0 0-16.233A8.116 8.116 0 0 0 12 21Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlarmClockFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.765 18.579a8.87 8.87 0 1 0-13.54 0l-1.88 1.88a.75.75 0 0 0 .53 1.28a.71.71 0 0 0 .53-.22l1.88-1.88a8.81 8.81 0 0 0 11.42 0l1.88 1.88a.71.71 0 0 0 .53.22a.741.741 0 0 0 .53-.22a.75.75 0 0 0 0-1.06zm-3.75-2.68a1 1 0 0 1-1.41 0l-2.61-2.61v-5.06a1 1 0 1 1 2 0v4.23l2 2a1 1 0 0 1 .02 1.44m5.67-9.8a.79.79 0 0 1-.53-.22l-2.31-2.32a.74.74 0 0 1 0-1.06a.75.75 0 0 1 1.06 0l2.36 2.28a.75.75 0 0 1 0 1.06a.79.79 0 0 1-.58.26m-17.42 0a.79.79 0 0 1-.53-.22a.75.75 0 0 1 0-1.06l2.32-2.32a.75.75 0 1 1 1.06 1.06l-2.32 2.32a.75.75 0 0 1-.53.22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignCenter(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10" stroke-width="1.5" d="M6.286 12h11.428m-8 5.714h4.572M4 6.286h16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignLeft(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10" stroke-width="1.5" d="M4.5 12h8m-8 6.25h15m-15-12.5h15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignRight(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10" stroke-width="1.5" d="M19.5 12h-8m8-6.25h-15m15 12.5h-15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Amazon(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.048 14.693a6.089 6.089 0 0 1-2.051 1.338a5.484 5.484 0 0 1-2.194.186a2.994 2.994 0 0 1-2.194-1.162a3.982 3.982 0 0 1 .066-4.531c1.448-1.821 3.85-1.722 5.836-2.04v-1.02a1.448 1.448 0 0 0-1.854-1.34c-.933.154-1.24.626-1.525 1.47c-.088.21-.263.231-.472.187c-.603 0-1.206-.098-1.81-.175c-.285 0-.647-.165-.526-.527c.124-.542.364-1.05.702-1.492a4.564 4.564 0 0 1 3.379-1.722a5.617 5.617 0 0 1 4.3.943a3.828 3.828 0 0 1 1.02 3.204v4.059a2.622 2.622 0 0 0 .592 1.667c.156.16.276.35.352.56a.384.384 0 0 1-.11.384c-.549.46-1.097.932-1.646 1.382a.504.504 0 0 1-.757 0a4.914 4.914 0 0 1-1.108-1.371m-.493-4.388c-1.35.076-3.182.263-3.291 1.974c-.06.475.056.956.329 1.35a1.481 1.481 0 0 0 2.194-.11c.888-.878.779-2.15.724-3.27zm-1.536 9.928A15.04 15.04 0 0 1 2.146 16.6c-.154-.098-.253-.493.055-.384a20.141 20.141 0 0 0 14.92 1.91c.997-.187 1.919-.703 2.895-.878c.867.395-1.24 1.228-1.525 1.448a15.883 15.883 0 0 1-6.472 1.536"/><path fill="currentColor" d="M19.994 16.58c-.592 0-1.042.054-1.59.12c0 0-.121 0-.121-.055c.208-.811 2.808-.998 3.488-.636c.384.208.176.746.143 1.097a3.994 3.994 0 0 1-1.207 2.194c-.197.22-.406 0-.263-.208a6.318 6.318 0 0 0 .592-2.063c.044-.493-.735-.395-1.042-.45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Apple(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.89 8.748a4.217 4.217 0 0 0-2.227 4.086a4.463 4.463 0 0 0 2.744 3.828a.362.362 0 0 1 0 .094a11.778 11.778 0 0 1-2.45 4.11a3.121 3.121 0 0 1-1.66 1.06a2.885 2.885 0 0 1-1.448-.118c-.495-.153-.978-.353-1.484-.495a4.228 4.228 0 0 0-2.661.236c-.429.157-.865.29-1.308.4a1.99 1.99 0 0 1-1.566-.294a5.887 5.887 0 0 1-1.413-1.284a12.377 12.377 0 0 1-2.673-5.994a7.537 7.537 0 0 1 .435-4.44a5.064 5.064 0 0 1 3.734-3.062a4.263 4.263 0 0 1 2.272.189l1.555.53c.273.105.575.105.848 0a18.84 18.84 0 0 1 2.014-.648a4.864 4.864 0 0 1 5.11 1.59z"/><path fill="currentColor" d="M16.191 2a3.921 3.921 0 0 1-.235 1.814a4.934 4.934 0 0 1-2.143 2.496a3.097 3.097 0 0 1-1.614.436c-.153 0-.188 0-.2-.2a4.18 4.18 0 0 1 .907-2.709a4.71 4.71 0 0 1 3.05-1.813z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Archive(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M4.382 8.813v8.5c0 .845.344 1.656.957 2.253a3.305 3.305 0 0 0 2.308.934h8.706c.866 0 1.696-.336 2.308-.934a3.15 3.15 0 0 0 .957-2.253v-8.5m0-5.313H4.382c-.901 0-1.632.714-1.632 1.594v2.125c0 .88.73 1.593 1.632 1.593h15.236c.901 0 1.632-.713 1.632-1.593V5.094c0-.88-.73-1.594-1.632-1.594"/><path stroke-miterlimit="10" d="M12 17v-5"/><path stroke-linejoin="round" d="m9.707 14.895l1.967 1.967a.458.458 0 0 0 .652 0l1.967-1.967"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArchiveDrawer(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19.194 9.079V7.821c0-.538-.216-1.054-.602-1.434a2.07 2.07 0 0 0-1.453-.594H6.86a2.07 2.07 0 0 0-1.453.594c-.386.38-.602.896-.602 1.434V9.08"/><path d="M6.861 5.793V4.779c0-.538.217-1.054.602-1.435a2.07 2.07 0 0 1 1.454-.594h6.166a2.07 2.07 0 0 1 1.454.594c.385.38.602.897.602 1.435v1.014m.781 3.043H6.08c-1.84 0-3.33 1.47-3.33 3.286v5.842c0 1.815 1.49 3.286 3.33 3.286h11.84c1.84 0 3.33-1.471 3.33-3.286v-5.842c0-1.815-1.49-3.286-3.33-3.286"/><path d="M7.889 12.893v1.014c0 .538.216 1.054.602 1.434c.385.38.908.594 1.453.594h4.112a2.07 2.07 0 0 0 1.453-.594c.386-.38.602-.896.602-1.434v-1.014"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArchiveDrawerFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.97 8.685v-.8a2.74 2.74 0 0 0-.83-2a2.88 2.88 0 0 0-1.23-.71v-.37a2.768 2.768 0 0 0-.82-2a2.82 2.82 0 0 0-2-.8H8.92a2.8 2.8 0 0 0-2.81 2.77v.37a2.88 2.88 0 0 0-1.23.71a2.74 2.74 0 0 0-.82 2v.8a4 4 0 0 0-2.06 3.5v5.84a4.07 4.07 0 0 0 4.08 4h11.84a4.07 4.07 0 0 0 4.08-4v-5.84a4 4 0 0 0-2.03-3.47m-3.09 5.28a2.781 2.781 0 0 1-2.8 2.78H9.97a2.81 2.81 0 0 1-2.62-1.725a2.76 2.76 0 0 1-.21-1.085v-1a.75.75 0 0 1 1.5 0v1a1.27 1.27 0 0 0 .38.9c.248.244.582.381.93.38h4.11a1.32 1.32 0 0 0 .92-.38a1.22 1.22 0 0 0 .38-.9v-1a.75.75 0 1 1 1.5 0zm1.59-5.78a3.08 3.08 0 0 0-.53 0H6.1a3 3 0 0 0-.52 0v-.3a1.23 1.23 0 0 1 .38-.9a1.32 1.32 0 0 1 .92-.38h10.28c.348 0 .682.136.93.38a1.27 1.27 0 0 1 .38.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArchiveFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.62 2.725H4.39A2.37 2.37 0 0 0 2 5.065v2.12a2.38 2.38 0 0 0 1.64 2.29v7.86a3.9 3.9 0 0 0 1.18 2.79a4 4 0 0 0 2.83 1.15h8.71a4 4 0 0 0 2.83-1.15a3.9 3.9 0 0 0 1.18-2.79v-7.86A2.38 2.38 0 0 0 22 7.235v-2.12a2.361 2.361 0 0 0-2.38-2.39m-4.79 12.67l-2 2a1.27 1.27 0 0 1-.39.26a1.091 1.091 0 0 1-.31.08h-.3a1.381 1.381 0 0 1-.29-.08a1.159 1.159 0 0 1-.39-.26l-2-2a.75.75 0 0 1 1.06-1.06l1 1v-3.36a.75.75 0 1 1 1.5 0v3.37l1-1a.75.75 0 0 1 1.06 1.06zm5.67-8.21a.861.861 0 0 1-.88.85H4.39a.87.87 0 0 1-.89-.85v-2.12a.86.86 0 0 1 .89-.84h15.23a.85.85 0 0 1 .88.84z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDown(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="M12 20V4"/><path stroke-linejoin="round" d="m4.34 12.968l6.572 6.572a1.531 1.531 0 0 0 2.176 0l6.573-6.572"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownCircle(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="M12 17.542V6.458"/><path stroke-linejoin="round" d="m6.722 12.697l4.529 4.58a1.057 1.057 0 0 0 1.499 0l4.528-4.58"/><path stroke-linejoin="round" d="M12 21.5a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownCircleFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.75A10.25 10.25 0 1 0 22.25 12A10.26 10.26 0 0 0 12 1.75m6 11.65l-4.53 4.58a1.93 1.93 0 0 1-.67.45a1.89 1.89 0 0 1-.79.16a2 2 0 0 1-.79-.16a1.85 1.85 0 0 1-.66-.45L6.03 13.4a1 1 0 1 1 1.42-1.41l3.56 3.6V6.46a1 1 0 1 1 2 0v9.08l3.57-3.55A1.001 1.001 0 0 1 18 13.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownLeft(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="M6.343 17.657L17.657 6.343"/><path stroke-linejoin="round" d="M5.899 7.267v9.296A1.53 1.53 0 0 0 7.437 18.1h9.296"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownLeftCircle(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="m8.232 15.768l7.832-7.842"/><path stroke-linejoin="round" d="M7.926 8.612v6.396a1.055 1.055 0 0 0 1.055 1.056h6.397"/><path stroke-linejoin="round" d="M12 21.5a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownLeftCircleFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.75A10.25 10.25 0 1 0 22.25 12A10.26 10.26 0 0 0 12 1.75m4.77 6.88l-6.38 6.39h5a1 1 0 0 1 0 2h-6.4a2 2 0 0 1-1.26-.44a.9.9 0 0 1-.19-.15a1.07 1.07 0 0 1-.14-.16a2 2 0 0 1-.46-1.29V8.61a1 1 0 0 1 2 0v5l6.43-6.44a1 1 0 1 1 1.41 1.41z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownLeftSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-miterlimit="10" d="m8.187 15.813l7.626-7.626"/><path stroke-linecap="round" stroke-linejoin="round" d="M7.889 8.845v6.238a1.028 1.028 0 0 0 1.027 1.028h6.24"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownLeftSquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 2h-6.5A6.76 6.76 0 0 0 2 8.75v6.5A6.76 6.76 0 0 0 8.75 22h6.5A6.76 6.76 0 0 0 22 15.25v-6.5A6.76 6.76 0 0 0 15.25 2m1.27 6.89l-6.2 6.2h4.84a1 1 0 0 1 0 2H8.92a2 2 0 0 1-1.26-.45a.73.73 0 0 1-.16-.14a.65.65 0 0 1-.13-.16a2 2 0 0 1-.46-1.28V8.82a1 1 0 0 1 2 0v4.86l6.22-6.22a1.001 1.001 0 1 1 1.41 1.41z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownRight(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="M17.657 17.657L6.343 6.343"/><path stroke-linejoin="round" d="M7.267 18.101h9.296a1.53 1.53 0 0 0 1.538-1.538V7.267"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownRightCircle(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="M15.768 15.768L7.926 7.926"/><path stroke-linejoin="round" d="M8.612 16.075h6.396a1.056 1.056 0 0 0 1.056-1.056V8.622"/><path stroke-linejoin="round" d="M12 21.5a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownRightCircleFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.75A10.25 10.25 0 1 0 22.25 12A10.26 10.26 0 0 0 12 1.75m5 13.27a2 2 0 0 1-.42 1.23a.82.82 0 0 1-.16.22a.996.996 0 0 1-.2.15a2 2 0 0 1-1.22.49H8.6a1 1 0 0 1 0-2h5L7.16 8.67a1 1 0 0 1 1.41-1.41l6.39 6.38v-5a1 1 0 1 1 2 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownRightSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-miterlimit="10" d="M15.813 15.813L8.187 8.187"/><path stroke-linecap="round" stroke-linejoin="round" d="M8.845 16.111h6.238a1.028 1.028 0 0 0 1.028-1.028V8.845"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownRightSquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 2h-6.5A6.76 6.76 0 0 0 2 8.75v6.5A6.76 6.76 0 0 0 8.75 22h6.5A6.76 6.76 0 0 0 22 15.25v-6.5A6.76 6.76 0 0 0 15.25 2m1.86 13.08a2 2 0 0 1-.46 1.28a1.002 1.002 0 0 1-.29.29a2 2 0 0 1-1.28.46H8.85a1 1 0 0 1 0-2h4.85L7.5 8.89a1 1 0 1 1 1.41-1.41l6.2 6.2V8.84a1 1 0 0 1 2 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-miterlimit="10" d="M12 17.19V6.398"/><path stroke-linecap="round" stroke-linejoin="round" d="m6.861 12.462l4.41 4.42a1.03 1.03 0 0 0 1.459 0l4.41-4.42"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownSquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 2h-6.5A6.76 6.76 0 0 0 2 8.75v6.5A6.76 6.76 0 0 0 8.75 22h6.5A6.76 6.76 0 0 0 22 15.25v-6.5A6.76 6.76 0 0 0 15.25 2m2.6 11.17l-4.41 4.42c-.188.19-.413.339-.66.44a2 2 0 0 1-1.56 0a1.998 1.998 0 0 1-.66-.44l-4.41-4.42a1 1 0 1 1 1.42-1.41L11 15.2V6.4a1 1 0 0 1 2 0v8.77l3.43-3.41a1.001 1.001 0 0 1 1.42 1.41"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeft(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="M4 12h16"/><path stroke-linejoin="round" d="M11.033 4.34L4.46 10.911a1.53 1.53 0 0 0 0 2.176l6.573 6.573"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftCircle(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="M6.458 12h11.084"/><path stroke-linejoin="round" d="m11.303 6.722l-4.528 4.529a1.056 1.056 0 0 0 0 1.499l4.528 4.528"/><path stroke-linejoin="round" d="M12 21.5a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftCircleFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.75A10.25 10.25 0 1 0 22.25 12A10.26 10.26 0 0 0 12 1.75M17.54 13h-9l3.51 3.57a1 1 0 0 1 0 1.41a1 1 0 0 1-1.41 0l-4.53-4.52a2.19 2.19 0 0 1-.45-.67a2 2 0 0 1 0-1.58a2.08 2.08 0 0 1 .45-.67l4.53-4.53a1 1 0 1 1 1.41 1.42L8.48 11h9.1a1 1 0 0 1 0 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-miterlimit="10" d="M6.81 12H17.6"/><path stroke-linecap="round" stroke-linejoin="round" d="m11.537 6.861l-4.419 4.41a1.028 1.028 0 0 0 0 1.459l4.42 4.409"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftSquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 2h-6.5A6.76 6.76 0 0 0 2 8.75v6.5A6.76 6.76 0 0 0 8.75 22h6.5A6.76 6.76 0 0 0 22 15.25v-6.5A6.76 6.76 0 0 0 15.25 2m2.35 11H8.83l3.41 3.43a1.002 1.002 0 0 1-.326 1.636a.999.999 0 0 1-.384.074a.999.999 0 0 1-.71-.29L6.4 13.44a2.12 2.12 0 0 1-.44-.66a2 2 0 0 1 0-1.56a2.12 2.12 0 0 1 .44-.66l4.42-4.41a1 1 0 1 1 1.41 1.42L8.79 11h8.8a1 1 0 0 1 0 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRight(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="M20 12H4"/><path stroke-linejoin="round" d="m12.968 19.66l6.572-6.572a1.53 1.53 0 0 0 0-2.176L12.968 4.34"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightCircle(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="M17.542 12H6.458"/><path stroke-linejoin="round" d="m12.697 17.278l4.58-4.528a1.058 1.058 0 0 0 0-1.5l-4.58-4.528"/><path stroke-linejoin="round" d="M12 21.5a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightCircleFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.75A10.25 10.25 0 1 0 22.25 12A10.26 10.26 0 0 0 12 1.75m6.43 11a2.19 2.19 0 0 1-.45.67l-4.58 4.53a1 1 0 0 1-1.42-.01a1 1 0 0 1 0-1.41l3.61-3.57H6.44a1 1 0 0 1 0-2h9.09l-3.55-3.57a1 1 0 1 1 1.42-1.41l4.58 4.53c.19.194.343.42.45.67a2 2 0 0 1 0 1.58z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-miterlimit="10" d="M17.396 12H6.604"/><path stroke-linecap="round" stroke-linejoin="round" d="m12.678 17.139l4.46-4.41a1.032 1.032 0 0 0 .226-1.124a1.03 1.03 0 0 0-.225-.335l-4.46-4.409"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightSquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 2h-6.5A6.76 6.76 0 0 0 2 8.75v6.5A6.76 6.76 0 0 0 8.75 22h6.5A6.76 6.76 0 0 0 22 15.25v-6.5A6.76 6.76 0 0 0 15.25 2m3 10.78a2.16 2.16 0 0 1-.45.66l-4.46 4.41a1 1 0 0 1-.7.29a1 1 0 0 1-.71-.3a1 1 0 0 1 0-1.41L15.5 13H6.6a1 1 0 0 1 0-2h8.82l-3.44-3.43a1 1 0 1 1 1.4-1.42l4.46 4.41c.19.19.343.414.45.66c.2.5.2 1.06 0 1.56z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUp(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="M12 4v16"/><path stroke-linejoin="round" d="M19.66 11.033L13.089 4.46a1.53 1.53 0 0 0-2.176 0L4.34 11.033"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpCircle(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="M12 6.458v11.084"/><path stroke-linejoin="round" d="m17.278 11.303l-4.529-4.528a1.056 1.056 0 0 0-1.498 0l-4.529 4.528"/><path stroke-linejoin="round" d="M12 21.5a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpCircleFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.75A10.25 10.25 0 1 0 22.25 12A10.26 10.26 0 0 0 12 1.75m6 10.26a1 1 0 0 1-1.42 0l-3.57-3.57v9.1a1 1 0 0 1-2 0v-9l-3.57 3.51a1 1 0 1 1-1.41-1.41l4.55-4.53a2.07 2.07 0 0 1 1.46-.61a2 2 0 0 1 .79.16c.253.1.482.254.67.45l4.53 4.53a1 1 0 0 1-.04 1.37z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpLeft(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="m6.343 6.343l11.314 11.314"/><path stroke-linejoin="round" d="M16.733 5.899H7.437A1.53 1.53 0 0 0 5.9 7.437v9.296"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpLeftCircle(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="m8.232 8.232l7.832 7.832"/><path stroke-linejoin="round" d="M15.388 7.926H8.992A1.056 1.056 0 0 0 7.936 8.98v6.397"/><path stroke-linejoin="round" d="M12 21.5a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpLeftCircleFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.75A10.25 10.25 0 1 0 22.25 12A10.26 10.26 0 0 0 12 1.75m4.77 15a1 1 0 0 1-1.41 0l-6.38-6.37v5a1 1 0 1 1-2 0v-6.4a2 2 0 0 1 .4-1.21a1.14 1.14 0 0 1 .18-.25a.83.83 0 0 1 .21-.16a2.05 2.05 0 0 1 1.26-.44h6.4a1 1 0 0 1 0 2h-5l6.43 6.44a1 1 0 0 1-.09 1.41z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpLeftSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-miterlimit="10" d="m8.187 8.187l7.626 7.626"/><path stroke-linecap="round" stroke-linejoin="round" d="M15.155 7.889H8.916A1.028 1.028 0 0 0 7.89 8.917v6.238"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpLeftSquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 2h-6.5A6.76 6.76 0 0 0 2 8.75v6.5A6.76 6.76 0 0 0 8.75 22h6.5A6.76 6.76 0 0 0 22 15.25v-6.5A6.76 6.76 0 0 0 15.25 2m1.25 14.5a1 1 0 0 1-1.41 0l-6.2-6.2v4.84a1 1 0 0 1-2 0V8.9a2 2 0 0 1 .46-1.28a1 1 0 0 1 .29-.29a2 2 0 0 1 1.28-.46h6.24a1 1 0 0 1 0 2H10.3l6.22 6.22a1 1 0 0 1-.02 1.41"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpRight(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="M17.657 6.343L6.343 17.657"/><path stroke-linejoin="round" d="M18.101 16.733V7.437A1.53 1.53 0 0 0 16.563 5.9H7.267"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpRightCircle(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="m15.768 8.232l-7.842 7.832"/><path stroke-linejoin="round" d="M16.074 15.388V8.992a1.055 1.055 0 0 0-1.055-1.056H8.612"/><path stroke-linejoin="round" d="M12 21.5a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpRightCircleFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.75A10.25 10.25 0 1 0 22.25 12A10.26 10.26 0 0 0 12 1.75m5.08 13.64a1 1 0 0 1-2 0v-5l-6.44 6.43a1.001 1.001 0 1 1-1.42-1.41l6.39-6.38h-5a1 1 0 0 1 0-2h6.4a2 2 0 0 1 1.25.42c.076.045.146.099.21.16a.31.31 0 0 1 .09.11a.36.36 0 0 1 .07.1a2 2 0 0 1 .44 1.26z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpRightSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-miterlimit="10" d="m15.813 8.187l-7.626 7.626"/><path stroke-linecap="round" stroke-linejoin="round" d="M16.111 15.155V8.917a1.028 1.028 0 0 0-1.028-1.028H8.845"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpRightSquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 2h-6.5A6.76 6.76 0 0 0 2 8.75v6.5A6.76 6.76 0 0 0 8.75 22h6.5A6.76 6.76 0 0 0 22 15.25v-6.5A6.76 6.76 0 0 0 15.25 2m1.86 13.16a1 1 0 1 1-2 0V10.3l-6.22 6.22a1 1 0 0 1-1.41 0a1 1 0 0 1 0-1.41l6.2-6.2H8.84a1 1 0 0 1 0-2h6.24a2 2 0 0 1 1.28.46a1 1 0 0 1 .29.29a2 2 0 0 1 .46 1.28z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-miterlimit="10" d="M12 6.81V17.6"/><path stroke-linecap="round" stroke-linejoin="round" d="m17.14 11.537l-4.41-4.419a1.028 1.028 0 0 0-1.46 0l-4.409 4.42"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpSquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 2h-6.5A6.76 6.76 0 0 0 2 8.75v6.5A6.76 6.76 0 0 0 8.75 22h6.5A6.76 6.76 0 0 0 22 15.25v-6.5A6.76 6.76 0 0 0 15.25 2m2.59 10.25a1 1 0 0 1-1.41 0L13 8.81v8.8a1 1 0 1 1-2 0V8.84l-3.43 3.41a1 1 0 1 1-1.41-1.42l4.41-4.42a2 2 0 0 1 2.88 0l4.41 4.42a1 1 0 0 1-.02 1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Arrowlist(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.104 5.552H21M8.104 12H21M8.104 18.448H21m-18-8.06L4.612 12L3 13.612M3 3.94l1.612 1.612L3 7.164m0 9.672l1.612 1.612L3 20.06"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsAllDirection(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21.5 12h-19m15.833 3.167L21.5 12l-3.167-3.167M5.667 15.167L2.5 12l3.167-3.167m3.166 9.5L12 21.5l3.167-3.167M8.833 5.667L12 2.5l3.167 3.167M12 21.5v-19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowsAllDirectionTwo(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20.375 20.375L3.625 3.625m11.163 16.75h5.587v-5.587M3.625 9.212V3.625h5.587M3.625 14.788v5.587h5.587m5.576-16.75h5.587v5.587M3.625 20.375l16.75-16.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Attachment(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.948 7.308L7.42 12.835a2.607 2.607 0 1 0 3.689 3.688l5.982-5.982a4.548 4.548 0 0 0 0-6.452a4.549 4.549 0 0 0-6.4.065l-6.034 5.918a6.517 6.517 0 0 0 9.215 9.214l7.377-7.312"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BagA(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6.45 9.23h11.1a1.85 1.85 0 0 1 1.85 1.85v6.472a3.698 3.698 0 0 1-3.7 3.698H8.3a3.7 3.7 0 0 1-3.7-3.698V11.08a1.849 1.849 0 0 1 1.85-1.85"/><path d="M16.625 11.553V7.381a4.62 4.62 0 0 0-2.852-4.278a4.627 4.627 0 0 0-6.398 4.278v4.172"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BagAfill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.41 9.238a2.62 2.62 0 0 0-1.84-.77h-.18v-1.06a5.318 5.318 0 0 0-1.57-3.81a5.55 5.55 0 0 0-1.76-1.19a5.44 5.44 0 0 0-4.12 0a5.6 5.6 0 0 0-1.75 1.17a5.43 5.43 0 0 0-1.57 3.8v1.1h-.17a2.61 2.61 0 0 0-2.6 2.6v6.48a4.44 4.44 0 0 0 4.45 4.44h7.4a4.441 4.441 0 0 0 3.14-1.3a4.391 4.391 0 0 0 1.31-3.14v-6.48a2.568 2.568 0 0 0-.74-1.84m-11.27 2.31a.75.75 0 1 1-1.5 0v-1.58h1.5zm0-3.08v-1.06a4 4 0 0 1 1.13-2.74a4.09 4.09 0 0 1 1.26-.84a3.92 3.92 0 0 1 4.26.84a4 4 0 0 1 .84 1.26c.195.468.294.972.29 1.48v1.1zm9.25 3.08a.75.75 0 1 1-1.5 0v-1.58h1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BagB(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6.45 9.23h11.1a1.85 1.85 0 0 1 1.85 1.85v6.472a3.698 3.698 0 0 1-3.7 3.698H8.3a3.7 3.7 0 0 1-3.7-3.698V11.08a1.849 1.849 0 0 1 1.85-1.85"/><path d="M7.375 9.23V7.381A4.621 4.621 0 0 1 12 2.75a4.627 4.627 0 0 1 4.625 4.631v1.85"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BagBfill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.41 9.238a2.62 2.62 0 0 0-1.84-.77h-.18v-1.06a5.318 5.318 0 0 0-1.57-3.81a5.55 5.55 0 0 0-1.76-1.19a5.44 5.44 0 0 0-4.12 0a5.6 5.6 0 0 0-1.75 1.17a5.43 5.43 0 0 0-1.57 3.8v1.1h-.17a2.61 2.61 0 0 0-2.6 2.6v6.48a4.44 4.44 0 0 0 4.45 4.44h7.4a4.441 4.441 0 0 0 3.14-1.3a4.391 4.391 0 0 0 1.31-3.14v-6.48a2.568 2.568 0 0 0-.74-1.84m-3.52-.77H8.14v-1.06a4 4 0 0 1 1.13-2.74a4.09 4.09 0 0 1 1.26-.84a3.92 3.92 0 0 1 4.26.84a4 4 0 0 1 .84 1.26c.195.468.294.972.29 1.48z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarCodeScan(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21.25 8.917V5.833a3.083 3.083 0 0 0-3.083-3.083h-3.084m0 18.5h3.084a3.083 3.083 0 0 0 3.083-3.083v-3.084m-18.5 0v3.084a3.083 3.083 0 0 0 3.083 3.083h3.084m0-18.5H5.833A3.083 3.083 0 0 0 2.75 5.833v3.084m3.189-.835v7.836M8.97 8.082v7.836M12 8.082v7.836m3.03-7.836v7.836m3.031-7.836v7.836"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Basket(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.298 9.566H4.702a1.964 1.964 0 0 0-1.535.744a1.944 1.944 0 0 0-.363 1.66l1.565 6.408a3.894 3.894 0 0 0 1.4 2.072c.682.519 1.517.8 2.376.8h7.708c.859 0 1.694-.281 2.376-.8a3.894 3.894 0 0 0 1.4-2.072l1.565-6.407a1.94 1.94 0 0 0-1.044-2.208a1.964 1.964 0 0 0-.854-.197M8.087 13.46v3.895M12 13.46v3.895m3.913-3.895v3.895m2.935-7.789a6.8 6.8 0 0 0-2.006-4.82A6.864 6.864 0 0 0 12 2.75a6.864 6.864 0 0 0-4.842 1.996a6.8 6.8 0 0 0-2.005 4.82"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BasketFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.947 10.941a2.82 2.82 0 0 0-.52-1.09a2.77 2.77 0 0 0-.94-.76a2.47 2.47 0 0 0-.92-.25a7.46 7.46 0 0 0-2.19-4.62a7.6 7.6 0 0 0-10.74 0a7.46 7.46 0 0 0-2.19 4.62a2.47 2.47 0 0 0-.92.25a2.68 2.68 0 0 0-.94.76a2.74 2.74 0 0 0-.52 2.3l1.57 6.43a4.65 4.65 0 0 0 4.5 3.42h7.71a4.67 4.67 0 0 0 4.51-3.44l1.56-6.41c.1-.396.11-.81.03-1.21m-13.1 6.42a.75.75 0 0 1-1.5 0v-3.94a.75.75 0 1 1 1.5 0zm3.91 0a.75.75 0 1 1-1.5 0v-3.94a.75.75 0 0 1 1.5 0zm3.91 0a.75.75 0 1 1-1.5 0v-3.94a.75.75 0 0 1 1.5 0zm-10.71-8.54a6 6 0 0 1 1.74-3.54a6.11 6.11 0 0 1 8.62 0a6 6 0 0 1 1.74 3.54z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryCharging(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.333 5.87H5.556A2.806 2.806 0 0 0 2.75 8.674v6.65a2.806 2.806 0 0 0 2.806 2.806h8.777a2.806 2.806 0 0 0 2.806-2.806v-6.65a2.806 2.806 0 0 0-2.806-2.806m5.355 9.221h.534a1.028 1.028 0 0 0 1.028-1.028V9.95a1.028 1.028 0 0 0-1.028-1.028h-.534"/><path d="M10.972 8.917L7.89 12H12l-3.083 3.083"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryChargingFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.31 5.12H5.53a3.55 3.55 0 0 0-3.55 3.55v6.65a3.56 3.56 0 0 0 3.55 3.56h8.78a3.57 3.57 0 0 0 3.56-3.56V8.68a3.56 3.56 0 0 0-3.56-3.56m-1.6 7.6L9.63 15.8a1 1 0 0 1-1.41 0a1 1 0 0 1 0-1.41l1.37-1.38H7.9a1 1 0 0 1-.71-1.71l3.08-3.08a1 1 0 1 1 1.42 1.41l-1.38 1.38h1.7a1 1 0 0 1 .7 1.71m7.49 3.13h-.49a.75.75 0 1 1 0-1.5h.53a.27.27 0 0 0 .28-.28V9.96a.3.3 0 0 0-.08-.2a.32.32 0 0 0-.2-.08h-.53a.75.75 0 1 1 0-1.5h.53a1.78 1.78 0 0 1 1.78 1.78v4.16a1.78 1.78 0 0 1-1.78 1.78z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryDead(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M14.333 5.87H5.556A2.806 2.806 0 0 0 2.75 8.674v6.65a2.806 2.806 0 0 0 2.806 2.806h8.777a2.806 2.806 0 0 0 2.806-2.806v-6.65a2.806 2.806 0 0 0-2.806-2.806m5.355 9.221h.534a1.028 1.028 0 0 0 1.028-1.028V9.95a1.028 1.028 0 0 0-1.028-1.028h-.534"/><path stroke-miterlimit="10" d="M12.013 9.931L7.876 14.07m0-4.139l4.137 4.138"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryDeadFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.33 5.12H5.55A3.55 3.55 0 0 0 2 8.67v6.65a3.56 3.56 0 0 0 3.55 3.56h8.78a3.56 3.56 0 0 0 3.55-3.56V8.68a3.549 3.549 0 0 0-3.55-3.56m-1.61 8.24a1 1 0 0 1-.71 1.71a1 1 0 0 1-.71-.29l-1.36-1.36l-1.36 1.36a1 1 0 0 1-.71.29a1 1 0 0 1-.71-1.71L8.52 12l-1.36-1.36a1.003 1.003 0 1 1 1.42-1.42l1.36 1.36l1.36-1.36a1.004 1.004 0 0 1 1.42 1.42L11.36 12zm7.5 2.49h-.54a.75.75 0 1 1 0-1.5h.54a.27.27 0 0 0 .28-.28V9.96a.251.251 0 0 0-.09-.2a.27.27 0 0 0-.19-.08h-.54a.75.75 0 0 1 0-1.5h.54A1.79 1.79 0 0 1 22 9.96v4.16a1.78 1.78 0 0 1-1.78 1.78z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryEmpty(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.333 5.87H5.556A2.806 2.806 0 0 0 2.75 8.674v6.65a2.806 2.806 0 0 0 2.806 2.806h8.777a2.806 2.806 0 0 0 2.806-2.806v-6.65a2.806 2.806 0 0 0-2.806-2.806m5.355 9.221h.534a1.028 1.028 0 0 0 1.028-1.028V9.95a1.028 1.028 0 0 0-1.028-1.028h-.534"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryEmptyFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.89 8.67v6.65a3.56 3.56 0 0 1-3.55 3.56H5.56A3.57 3.57 0 0 1 2 15.32V8.67a3.56 3.56 0 0 1 3.56-3.55h8.78a3.55 3.55 0 0 1 3.55 3.55m2.34 7.17h-.54a.75.75 0 1 1 0-1.5h.54a.27.27 0 0 0 .19-.08a.32.32 0 0 0 .08-.2V9.95a.3.3 0 0 0-.08-.2a.27.27 0 0 0-.19-.08h-.54a.75.75 0 1 1 0-1.5h.54A1.78 1.78 0 0 1 22 9.95v4.16a1.768 1.768 0 0 1-.52 1.26a1.79 1.79 0 0 1-1.25.47"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryFull(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.333 5.87H5.556A2.806 2.806 0 0 0 2.75 8.674v6.65a2.806 2.806 0 0 0 2.806 2.806h8.777a2.806 2.806 0 0 0 2.806-2.806v-6.65a2.806 2.806 0 0 0-2.806-2.806m5.355 9.221h.534a1.028 1.028 0 0 0 1.028-1.028V9.95a1.028 1.028 0 0 0-1.028-1.028h-.534M6.861 9.95v4.11m3.083-4.11v4.11m3.084-4.11v4.11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryFullFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.32 5.12H5.55a3.56 3.56 0 0 0-3.56 3.55v6.65a3.57 3.57 0 0 0 3.56 3.56h8.77a3.57 3.57 0 0 0 3.56-3.56V8.68a3.56 3.56 0 0 0-3.56-3.56m-6.47 8.94a1 1 0 0 1-2 0v-4.1a1 1 0 1 1 2 0zm3.08 0a1 1 0 1 1-2 0v-4.1a1 1 0 0 1 2 0zm3.09 0a1 1 0 1 1-2 0v-4.1a1 1 0 0 1 2 0zm6.21 1.79h-.53a.75.75 0 1 1 0-1.5h.53a.28.28 0 0 0 .2-.08a.32.32 0 0 0 .08-.2V9.96a.27.27 0 0 0-.28-.28h-.53a.75.75 0 1 1 0-1.5h.53a1.8 1.8 0 0 1 1.26.52a1.749 1.749 0 0 1 .52 1.26v4.16a1.77 1.77 0 0 1-.52 1.26a1.8 1.8 0 0 1-1.26.47"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryHalf(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.333 5.87H5.556A2.806 2.806 0 0 0 2.75 8.674v6.65a2.806 2.806 0 0 0 2.806 2.806h8.777a2.806 2.806 0 0 0 2.806-2.806v-6.65a2.806 2.806 0 0 0-2.806-2.806m5.355 9.221h.534a1.028 1.028 0 0 0 1.028-1.028V9.95a1.028 1.028 0 0 0-1.028-1.028h-.534M6.861 9.95v4.11m3.083-4.11v4.11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryHalfFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.33 5.12H5.55A3.55 3.55 0 0 0 2 8.67v6.65a3.56 3.56 0 0 0 3.55 3.56h8.78a3.57 3.57 0 0 0 3.56-3.56V8.68a3.56 3.56 0 0 0-3.56-3.56m-6.47 8.94a1 1 0 1 1-2 0v-4.1a1 1 0 0 1 2 0zm3.08 0a1 1 0 1 1-2 0v-4.1a1 1 0 0 1 2 0zm9.28 1.79h-.53a.75.75 0 1 1 0-1.5h.53a.3.3 0 0 0 .2-.08a.32.32 0 0 0 .08-.2V9.96a.281.281 0 0 0-.08-.2a.3.3 0 0 0-.2-.08h-.53a.75.75 0 1 1 0-1.5h.53A1.78 1.78 0 0 1 22 9.96v4.16a1.78 1.78 0 0 1-1.78 1.78z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryLow(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.333 5.87H5.556A2.806 2.806 0 0 0 2.75 8.674v6.65a2.806 2.806 0 0 0 2.806 2.806h8.777a2.806 2.806 0 0 0 2.806-2.806v-6.65a2.806 2.806 0 0 0-2.806-2.806m5.355 9.221h.534a1.028 1.028 0 0 0 1.028-1.028V9.95a1.028 1.028 0 0 0-1.028-1.028h-.534M6.861 9.95v4.11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryLowFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.325 5.12h-8.78a3.55 3.55 0 0 0-3.55 3.55v6.65a3.56 3.56 0 0 0 3.55 3.56h8.78a3.57 3.57 0 0 0 3.56-3.56V8.68a3.56 3.56 0 0 0-3.56-3.56m-6.47 8.94a1 1 0 0 1-2 0v-4.1a1 1 0 1 1 2 0zm12.36 1.79h-.53a.75.75 0 0 1 0-1.5h.54a.27.27 0 0 0 .19-.08a.281.281 0 0 0 .09-.2V9.96a.25.25 0 0 0-.09-.2a.27.27 0 0 0-.19-.08h-.54a.75.75 0 0 1 0-1.5h.54a1.78 1.78 0 0 1 1.78 1.78v4.16a1.78 1.78 0 0 1-1.78 1.78z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Behance(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.354 11.107V9.502c0-.095 0-.133.124-.133c1.063 0 1.927-.063 2.356.19c.43.252.448.6.455.704c.008.104.004.253-.196.616s-1.095.37-1.741.37h-.855c-.105 0-.143-.038-.143-.142m2.946 1.682a1.121 1.121 0 0 1-.324 1.662a1.71 1.71 0 0 1-.646.133H7.487c-.095 0-.133 0-.133-.123v-1.9c0-.095 0-.133.124-.133h1.225c.76 0 1.235-.067 1.597.36"/><path fill="currentColor" d="M16.989 11.909a.869.869 0 0 0-.566-.498a1.842 1.842 0 0 0-1.492.235a1.426 1.426 0 0 0-.4.809c0 .107 0 .127.127.127h2.37c.117 0 .117 0 .107-.147a1.307 1.307 0 0 0-.147-.526m0 0a.869.869 0 0 0-.566-.498a1.842 1.842 0 0 0-1.492.235a1.426 1.426 0 0 0-.4.809c0 .107 0 .127.127.127h2.37c.117 0 .117 0 .107-.147a1.307 1.307 0 0 0-.147-.526M18.1 2.5H6.4a3.9 3.9 0 0 0-3.9 3.9v11.7A3.9 3.9 0 0 0 6.4 22h11.7a3.9 3.9 0 0 0 3.9-3.9V6.4a3.9 3.9 0 0 0-3.9-3.9m-3.822 6.21c0-.077 0-.107.088-.107h3.159v.82s.003.058-.049.058h-3.061c-.137 0-.137 0-.137-.136zm-2.925 7.167c-1.2.653-3.217.39-5.606.39V8.184h3.812c.975 0 2.135.244 2.447 1.277a1.95 1.95 0 0 1-.975 2.409a2.213 2.213 0 0 1 .312 4.007zm7.264-2.301h-3.9c-.107 0-.205.058-.186.195c.117.575.38 1.18.976 1.268a1.473 1.473 0 0 0 1.686-.663s.049-.059.078-.059h1.268s.078 0 .058.068a3.45 3.45 0 0 1-.585 1.043a3.051 3.051 0 0 1-4.787-.75a2.465 2.465 0 0 1-.253-.878a3.023 3.023 0 0 1 2.203-3.597c2.213-.478 3.617 1.043 3.617 3.188a.165.165 0 0 1-.185.185zm-2.184-2.165a1.843 1.843 0 0 0-.771-.042c-.257.04-.487.089-.721.277c-.234.187-.358.503-.4.809c0 .107 0 .127.127.127h2.36c.117 0 .117 0 .107-.147a1.307 1.307 0 0 0-.147-.526a.869.869 0 0 0-.565-.498z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellNotificationSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.4" d="M11.974 16.208h4.786a1.095 1.095 0 0 0 1.063-1.53c-.252-.764-1.262-1.682-1.262-2.584c0-2.003 0-2.53-.986-3.708a3.532 3.532 0 0 0-1.162-.903l-.55-.267a.765.765 0 0 1-.36-.513a1.422 1.422 0 0 0-1.53-1.2a1.423 1.423 0 0 0-1.49 1.2a.765.765 0 0 1-.398.513l-.55.267a3.533 3.533 0 0 0-1.163.903c-.986 1.177-.986 1.705-.986 3.708c0 .902-.964 1.728-1.216 2.539c-.153.489-.237 1.575 1.04 1.575z"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.4" d="M14.267 16.208a2.244 2.244 0 0 1-1.408 2.132a2.244 2.244 0 0 1-.885.162a2.245 2.245 0 0 1-2.294-2.294"/><rect width="18.5" height="18.5" x="2.75" y="2.75" stroke-width="1.5" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellNotificationSquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 2.02h-6.5A6.76 6.76 0 0 0 2 8.77v6.46a6.76 6.76 0 0 0 6.75 6.75h6.5A6.76 6.76 0 0 0 22 15.23v-6.5a6.76 6.76 0 0 0-6.75-6.71m2 13.12a.49.49 0 0 1-.08.23a.26.26 0 0 1-.15.12a.492.492 0 0 1-.24 0H7.23c-.18 0-.33 0-.37-.09a.8.8 0 0 1 0-.57a3.8 3.8 0 0 1 .47-.83a3.51 3.51 0 0 0 .78-1.92c0-2 0-2.28.82-3.26a3 3 0 0 1 .93-.72l.57-.28a1.24 1.24 0 0 0 .48-.4c.15-.178.25-.391.29-.62a.77.77 0 0 1 .26-.46a.74.74 0 0 1 .45-.15h.15a.86.86 0 0 1 .28 0a.76.76 0 0 1 .23.13a.68.68 0 0 1 .27.48c.044.2.126.39.24.56c.129.187.3.342.5.45l.55.27c.358.17.676.416.93.72c.82 1 .82 1.29.82 3.26c.078.735.37 1.43.84 2c.196.267.36.556.49.86a.34.34 0 0 1 .02.22zm-3.53 2.23c-.216.222-.48.392-.77.5a2.28 2.28 0 0 1-.94.17a2.998 2.998 0 0 1-.89-.16a2.32 2.32 0 0 1-1.21-1.16h4.25a2 2 0 0 1-.46.65z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bolt(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20.567 14.54V9.46a4.164 4.164 0 0 0-2.082-3.602l-4.403-2.55a4.163 4.163 0 0 0-4.164 0l-4.403 2.55A4.164 4.164 0 0 0 3.433 9.46v5.08a4.164 4.164 0 0 0 2.082 3.602l4.403 2.55a4.163 4.163 0 0 0 4.164 0l4.403-2.55a4.164 4.164 0 0 0 2.082-3.602"/><path d="M12.906 8.346h-1.812a2.082 2.082 0 0 0-1.81 1.041l-.907 1.572a2.082 2.082 0 0 0 0 2.082l.906 1.572a2.082 2.082 0 0 0 1.811 1.04h1.812a2.083 2.083 0 0 0 1.81-1.04l.907-1.572a2.082 2.082 0 0 0 0-2.082l-.906-1.572a2.081 2.081 0 0 0-1.811-1.04"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoltFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.662 6.998a4.914 4.914 0 0 0-1.791-1.801l-4.414-2.552a5.004 5.004 0 0 0-4.914 0L5.139 5.197A4.924 4.924 0 0 0 2.677 9.45v5.085a4.914 4.914 0 0 0 2.452 4.253l4.414 2.552a4.904 4.904 0 0 0 4.914 0l4.414-2.552a4.884 4.884 0 0 0 1.791-1.791c.432-.75.66-1.598.66-2.463V9.45c0-.86-.228-1.707-.66-2.452m-4.854 6.255l-.81 1.402a2.492 2.492 0 0 1-.921.93c-.384.22-.819.338-1.261.34h-1.632a2.562 2.562 0 0 1-1.26-.34a2.492 2.492 0 0 1-.922-.93l-.81-1.402a2.492 2.492 0 0 1 0-2.512l.81-1.4c.216-.39.535-.712.921-.932c.384-.22.819-.337 1.261-.34h1.632c.442.003.877.12 1.26.34c.385.223.703.544.921.931l.811 1.401a2.491 2.491 0 0 1 0 2.512"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Book(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 16.401a1.15 1.15 0 0 0 1.16 1.15a16.705 16.705 0 0 1 3.535.333c1.64.204 3.204.81 4.555 1.761V6.442A10.238 10.238 0 0 0 7.445 4.68a16.588 16.588 0 0 0-3.6-.322a1.15 1.15 0 0 0-1.074 1.15zm18.5 0a1.15 1.15 0 0 1-1.16 1.15a16.705 16.705 0 0 0-3.535.333c-1.64.204-3.204.81-4.555 1.761V6.442a10.238 10.238 0 0 1 4.555-1.762a16.588 16.588 0 0 1 3.6-.322a1.15 1.15 0 0 1 1.073 1.15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 16.653c0 .25-.05.5-.15.73a2 2 0 0 1-.41.62c-.181.171-.391.31-.62.41a1.939 1.939 0 0 1-.74.14a15.22 15.22 0 0 0-3.37.32a9.312 9.312 0 0 0-3.71 1.27V5.233c1.091-.52 2.26-.858 3.46-1a17.43 17.43 0 0 1 3.71-.33a1.92 1.92 0 0 1 1.3.61c.33.352.513.817.51 1.3zM11 5.233v14.91a9.248 9.248 0 0 0-3.65-1.27a16.174 16.174 0 0 0-3.43-.32a1.87 1.87 0 0 1-.74-.14a2.16 2.16 0 0 1-.62-.41a1.81 1.81 0 0 1-.41-.62a1.83 1.83 0 0 1-.15-.73v-10.9a1.9 1.9 0 0 1 1.78-1.89a16.999 16.999 0 0 1 3.79.33A10.68 10.68 0 0 1 11 5.233"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookText(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.756 16.358a1.09 1.09 0 0 0 1.154 1.198a16.576 16.576 0 0 1 3.54.338c1.635.2 3.197.794 4.552 1.731V6.448A10.16 10.16 0 0 0 7.45 4.694a16.597 16.597 0 0 0-3.605-.316a1.09 1.09 0 0 0-1.09 1.09zm18.492 0a1.089 1.089 0 0 1-1.154 1.154a16.576 16.576 0 0 0-3.54.338a10.16 10.16 0 0 0-4.552 1.775V6.448a10.16 10.16 0 0 1 4.552-1.754a16.597 16.597 0 0 1 3.605-.316a1.089 1.089 0 0 1 1.089 1.155zM5.621 8.234h1.252m-1.252 6.011h1.834M5.78 11.24h3.35"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookTextFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 5.797v10.76a1.75 1.75 0 0 1-.12.72a1.86 1.86 0 0 1-1.83 1.19a15.391 15.391 0 0 0-3.35.32a9.66 9.66 0 0 0-3.7 1.29V5.218a10.82 10.82 0 0 1 3.46-1a16.84 16.84 0 0 1 3.72-.32a1.62 1.62 0 0 1 .71.15c.23.098.437.24.61.42a2 2 0 0 1 .39.64c.08.221.118.455.11.69M7.6 4.187a17 17 0 0 0-3.76-.33a1.81 1.81 0 0 0-1.83 1.83v10.87a1.79 1.79 0 0 0 .51 1.46a1.82 1.82 0 0 0 1.4.56a15.6 15.6 0 0 1 3.44.34a9.265 9.265 0 0 1 3.64 1.23V5.217a10.44 10.44 0 0 0-3.4-1.03m-2 3.52h1.25a.75.75 0 1 1 0 1.5H5.6a.75.75 0 1 1 0-1.5m1.84 7.51H5.6a.75.75 0 1 1 0-1.5h1.84a.75.75 0 1 1 0 1.5m1.67-3H5.76a.75.75 0 1 1 0-1.5h3.35a.75.75 0 1 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookmark(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m10.94 18.339l-3.43 2.548a1.71 1.71 0 0 1-2.76-1.23V6.35a3.735 3.735 0 0 1 3.87-3.597h6.76a3.742 3.742 0 0 1 3.87 3.597v13.309a1.708 1.708 0 0 1-2.76 1.229l-3.43-2.548a1.801 1.801 0 0 0-2.12 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkCheck(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m10.94 18.339l-3.43 2.548a1.71 1.71 0 0 1-2.76-1.23V6.35a3.735 3.735 0 0 1 3.87-3.597h6.76a3.742 3.742 0 0 1 3.87 3.597v13.309a1.708 1.708 0 0 1-2.76 1.229l-3.43-2.548a1.801 1.801 0 0 0-2.12 0"/><path d="m9 9.918l1.689 1.689a.639.639 0 0 0 .908 0L15 8.204"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkCheckFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.59 3.204a4.41 4.41 0 0 0-3.19-1.2H8.67a4.49 4.49 0 0 0-4.65 4.34v13.36c.03.447.178.879.43 1.25c.254.36.6.646 1 .83c.408.18.857.249 1.3.2a2.47 2.47 0 0 0 1.2-.5l3.43-2.54a1.07 1.07 0 0 1 1.23 0l3.41 2.53a2.46 2.46 0 0 0 1.22.51h.3a2.498 2.498 0 0 0 1-.22c.403-.18.749-.467 1-.83a2.47 2.47 0 0 0 .44-1.3V6.324a4.41 4.41 0 0 0-1.39-3.12m-3 5.53l-3.4 3.4a1.439 1.439 0 0 1-.45.31a1.61 1.61 0 0 1-.53.1a1.398 1.398 0 0 1-.54-.1a1.42 1.42 0 0 1-.45-.32l-1.69-1.68a.75.75 0 0 1 1.06-1.06l1.62 1.62l3.32-3.33a.75.75 0 0 1 1.06 0a.74.74 0 0 1 0 1.06"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkCross(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m10.94 18.339l-3.43 2.548a1.71 1.71 0 0 1-2.76-1.23V6.35a3.735 3.735 0 0 1 3.87-3.597h6.76a3.742 3.742 0 0 1 3.87 3.597v13.309a1.708 1.708 0 0 1-2.76 1.229l-3.43-2.548a1.801 1.801 0 0 0-2.12 0"/><path stroke-miterlimit="10" d="m14.25 7.75l-4.5 4.49m0-4.48l4.5 4.49"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkCrossFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.59 3.204a4.41 4.41 0 0 0-3.19-1.2H8.67a4.49 4.49 0 0 0-4.65 4.34v13.36c.03.447.178.879.43 1.25c.254.36.6.646 1 .83c.408.18.857.249 1.3.2a2.47 2.47 0 0 0 1.2-.5l3.43-2.54a1.07 1.07 0 0 1 1.23 0l3.41 2.53a2.46 2.46 0 0 0 1.22.51h.3a2.5 2.5 0 0 0 1-.22a2.41 2.41 0 0 0 1-.83a2.47 2.47 0 0 0 .44-1.3V6.324a4.41 4.41 0 0 0-1.39-3.12m-3.79 8.51a.742.742 0 0 1 0 1.06a.72.72 0 0 1-.53.23a.79.79 0 0 1-.53-.22l-1.73-1.72l-1.71 1.71a.753.753 0 1 1-1.06-1.07l1.71-1.7l-1.71-1.71a.75.75 0 0 1 1.06-1.06l1.71 1.71l1.73-1.72a.75.75 0 0 1 1.06 0a.741.741 0 0 1 0 1.06l-1.72 1.72z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkDownload(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m10.94 18.339l-3.43 2.548a1.71 1.71 0 0 1-2.76-1.23V6.35a3.735 3.735 0 0 1 3.87-3.597h6.76a3.742 3.742 0 0 1 3.87 3.597v13.309a1.708 1.708 0 0 1-2.76 1.229l-3.43-2.548a1.801 1.801 0 0 0-2.12 0"/><path stroke-miterlimit="10" d="M12 13V7"/><path stroke-linejoin="round" d="m9.249 10.474l2.36 2.36a.552.552 0 0 0 .782 0l2.36-2.36"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkDownloadFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.59 3.204a4.42 4.42 0 0 0-3.19-1.2H8.67a4.44 4.44 0 0 0-3.22 1.2a4.5 4.5 0 0 0-1.43 3.14v13.36c.03.448.178.879.43 1.25c.254.36.6.646 1 .83c.408.18.857.25 1.3.2a2.465 2.465 0 0 0 1.2-.5l3.43-2.54a1.07 1.07 0 0 1 1.23 0l3.41 2.53c.352.278.774.455 1.22.51h.3a2.5 2.5 0 0 0 1-.22a2.41 2.41 0 0 0 1-.83a2.47 2.47 0 0 0 .44-1.3V6.324a4.41 4.41 0 0 0-1.39-3.12m-3.28 7.8l-2.37 2.36a1.192 1.192 0 0 1-.42.28a.999.999 0 0 1-.34.09h-.3a.999.999 0 0 1-.34-.09a1.29 1.29 0 0 1-.42-.28l-2.36-2.36a.75.75 0 0 1 1.06-1.06l1.47 1.47v-4.41a.75.75 0 1 1 1.5 0v4.41l1.48-1.47a.75.75 0 0 1 1.06 1.06z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.97 6.321v13.33a2.47 2.47 0 0 1-1.45 2.13a2.53 2.53 0 0 1-1.3.2a2.46 2.46 0 0 1-1.22-.51l-3.41-2.53a1.07 1.07 0 0 0-1.23 0l-3.43 2.56a2.465 2.465 0 0 1-1.2.5h-.3a2.43 2.43 0 0 1-1-.22a2.5 2.5 0 0 1-1-.83a2.53 2.53 0 0 1-.43-1.25V6.342a4.49 4.49 0 0 1 4.65-4.34h6.73A4.49 4.49 0 0 1 20 6.321z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkMinus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m10.94 18.339l-3.43 2.548a1.71 1.71 0 0 1-2.76-1.23V6.35a3.735 3.735 0 0 1 3.87-3.597h6.76a3.742 3.742 0 0 1 3.87 3.597v13.309a1.708 1.708 0 0 1-2.76 1.229l-3.43-2.548a1.801 1.801 0 0 0-2.12 0"/><path stroke-miterlimit="10" d="M9 10.007h6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkMinusFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.585 3.194a4.41 4.41 0 0 0-3.19-1.2h-6.73a4.49 4.49 0 0 0-4.65 4.34v13.36c.03.447.178.879.43 1.25c.254.36.6.646 1 .83a2.43 2.43 0 0 0 1 .22h.3a2.47 2.47 0 0 0 1.2-.5l3.43-2.54a1.07 1.07 0 0 1 1.23 0l3.41 2.53c.352.278.774.454 1.22.51c.443.049.892-.02 1.3-.2a2.47 2.47 0 0 0 1.45-2.13V6.314a4.409 4.409 0 0 0-1.4-3.12m-3.57 7.55h-6a.75.75 0 1 1 0-1.5h6a.75.75 0 1 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkPlus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m10.94 18.339l-3.43 2.548a1.71 1.71 0 0 1-2.76-1.23V6.35a3.735 3.735 0 0 1 3.87-3.597h6.76a3.742 3.742 0 0 1 3.87 3.597v13.309a1.708 1.708 0 0 1-2.76 1.229l-3.43-2.548a1.801 1.801 0 0 0-2.12 0"/><path stroke-miterlimit="10" d="M11.993 7v6M9 10.007h6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkPlusFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.585 3.194a4.41 4.41 0 0 0-3.19-1.2h-6.73a4.49 4.49 0 0 0-4.65 4.34v13.36c.03.447.178.879.43 1.25c.254.36.6.646 1 .83a2.43 2.43 0 0 0 1 .22h.3a2.47 2.47 0 0 0 1.2-.5l3.43-2.54a1.07 1.07 0 0 1 1.23 0l3.41 2.53c.352.278.774.454 1.22.51c.443.049.892-.02 1.3-.2a2.47 2.47 0 0 0 1.45-2.13V6.314a4.409 4.409 0 0 0-1.4-3.12m-3.57 7.55h-2.25v2.25a.75.75 0 1 1-1.5 0v-2.25h-2.25a.75.75 0 1 1 0-1.5h2.25v-2.25a.75.75 0 1 1 1.5 0v2.25h2.25a.75.75 0 1 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkQuestionMark(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m10.94 18.339l-3.43 2.548a1.71 1.71 0 0 1-2.76-1.23V6.35a3.735 3.735 0 0 1 3.87-3.597h6.76a3.742 3.742 0 0 1 3.87 3.597v13.309a1.708 1.708 0 0 1-2.76 1.229l-3.43-2.548a1.801 1.801 0 0 0-2.12 0"/><path stroke-miterlimit="10" d="M10.274 7.628a1.834 1.834 0 0 1 1.999-1.04a1.78 1.78 0 0 1 1.304.93a1.545 1.545 0 0 1-.9 2.124a1.14 1.14 0 0 0-.734 1.03v.424"/><path stroke-linejoin="round" d="M11.91 13.442h.005"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkQuestionMarkFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.59 3.204a4.42 4.42 0 0 0-3.19-1.2H8.67a4.49 4.49 0 0 0-4.65 4.34v13.36c.03.447.178.879.43 1.25c.254.36.6.646 1 .83c.408.18.857.249 1.3.2a2.47 2.47 0 0 0 1.2-.5l3.43-2.54a1.07 1.07 0 0 1 1.23 0l3.41 2.53a2.46 2.46 0 0 0 1.22.51h.3a2.501 2.501 0 0 0 1-.22a2.41 2.41 0 0 0 1-.83a2.47 2.47 0 0 0 .44-1.3V6.324a4.41 4.41 0 0 0-1.39-3.12m-6.65 11a.75.75 0 0 1-.75-.75a.74.74 0 0 1 .74-.75a.75.75 0 1 1 0 1.5zm2.38-5.12a2.26 2.26 0 0 1-1.38 1.28a.3.3 0 0 0-.15.13a.31.31 0 0 0-.07.21v.4a.75.75 0 1 1-1.5 0v-.42a1.91 1.91 0 0 1 .34-1.06a1.93 1.93 0 0 1 .88-.67a1.15 1.15 0 0 0 .31-.18a.92.92 0 0 0 .19-.28a.87.87 0 0 0 .06-.32a.69.69 0 0 0-.08-.32a1.05 1.05 0 0 0-.3-.34a1.14 1.14 0 0 0-.44-.18a1.17 1.17 0 0 0-.72.1c-.21.118-.379.3-.48.52a.75.75 0 0 1-1.36-.63a2.55 2.55 0 0 1 2.81-1.46a2.38 2.38 0 0 1 1.06.43c.325.23.59.535.77.89c.158.31.24.652.24 1a2.33 2.33 0 0 1-.15.89z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkUpload(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m10.94 18.339l-3.43 2.548a1.71 1.71 0 0 1-2.76-1.23V6.35a3.735 3.735 0 0 1 3.87-3.597h6.76a3.742 3.742 0 0 1 3.87 3.597v13.309a1.708 1.708 0 0 1-2.76 1.229l-3.43-2.548a1.801 1.801 0 0 0-2.12 0"/><path stroke-miterlimit="10" d="M12 7v6"/><path stroke-linejoin="round" d="m14.752 9.526l-2.361-2.36a.55.55 0 0 0-.782 0l-2.36 2.36"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkUploadFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.59 3.204a4.41 4.41 0 0 0-3.19-1.2H8.67a4.49 4.49 0 0 0-4.65 4.34v13.36c.03.447.178.879.43 1.25c.254.36.6.646 1 .83c.408.18.857.249 1.3.2a2.47 2.47 0 0 0 1.2-.5l3.43-2.54a1.07 1.07 0 0 1 1.23 0l3.41 2.53a2.46 2.46 0 0 0 1.22.51h.3a2.5 2.5 0 0 0 1-.22a2.41 2.41 0 0 0 1-.83a2.47 2.47 0 0 0 .44-1.3V6.324a4.41 4.41 0 0 0-1.39-3.12m-3.28 6.85a.738.738 0 0 1-.53.22a.708.708 0 0 1-.53-.22l-1.48-1.47v4.42a.75.75 0 1 1-1.5 0v-4.42l-1.47 1.47a.75.75 0 0 1-1.06-1.06l2.36-2.36c.124-.122.27-.22.43-.29a1.14 1.14 0 0 1 .33-.08a.86.86 0 0 1 .32 0c.121.01.24.04.35.09a1.4 1.4 0 0 1 .42.28l2.36 2.36a.75.75 0 0 1 .03 1.06z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Box(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.382 8.813v8.5c0 .845.344 1.656.957 2.253a3.305 3.305 0 0 0 2.308.934h8.706c.866 0 1.696-.336 2.308-.934a3.15 3.15 0 0 0 .957-2.253v-8.5m0-5.313H4.382c-.901 0-1.632.714-1.632 1.594v2.125c0 .88.73 1.593 1.632 1.593h15.236c.901 0 1.632-.713 1.632-1.593V5.094c0-.88-.73-1.594-1.632-1.594"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxCheck(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4.382 8.813v8.5c0 .845.344 1.656.957 2.253a3.305 3.305 0 0 0 2.308.934h8.706c.866 0 1.696-.336 2.308-.934a3.15 3.15 0 0 0 .957-2.253v-8.5m0-5.313H4.382c-.901 0-1.632.714-1.632 1.594v2.125c0 .88.73 1.593 1.632 1.593h15.236c.901 0 1.632-.713 1.632-1.593V5.094c0-.88-.73-1.594-1.632-1.594"/><path d="m9 14.574l1.689 1.689a.637.637 0 0 0 .908 0L15 12.86"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxCheckFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.62 2.725H4.39A2.37 2.37 0 0 0 2 5.065v2.12a2.38 2.38 0 0 0 1.64 2.29v7.86a3.9 3.9 0 0 0 1.18 2.79a4 4 0 0 0 2.83 1.15h8.71a4 4 0 0 0 2.83-1.15a3.9 3.9 0 0 0 1.18-2.79v-7.86A2.38 2.38 0 0 0 22 7.235v-2.12a2.36 2.36 0 0 0-2.38-2.39m-4.08 10.64l-3.4 3.4a1.422 1.422 0 0 1-.985.41a1.36 1.36 0 0 1-.535-.11a1.25 1.25 0 0 1-.45-.31l-1.69-1.68a.75.75 0 1 1 1.06-1.06l1.62 1.62l3.32-3.33a.752.752 0 0 1 1.079-.019a.748.748 0 0 1-.019 1.079m5-6.18a.861.861 0 0 1-.88.85H4.39a.87.87 0 0 1-.89-.85v-2.12a.86.86 0 0 1 .89-.84h15.23a.85.85 0 0 1 .88.84z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxCross(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M4.382 8.813v8.5c0 .845.344 1.656.957 2.253a3.305 3.305 0 0 0 2.308.934h8.706c.866 0 1.696-.336 2.308-.934a3.15 3.15 0 0 0 .957-2.253v-8.5m0-5.313H4.382c-.901 0-1.632.714-1.632 1.594v2.125c0 .88.73 1.593 1.632 1.593h15.236c.901 0 1.632-.713 1.632-1.593V5.094c0-.88-.73-1.594-1.632-1.594"/><path stroke-miterlimit="10" d="m14 12.656l-4 3.992m0-3.983l4 3.991"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxCrossFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.62 2.725H4.39A2.37 2.37 0 0 0 2 5.065v2.12a2.38 2.38 0 0 0 1.64 2.29v7.86a3.9 3.9 0 0 0 1.18 2.79a4 4 0 0 0 2.83 1.15h8.71a4 4 0 0 0 2.83-1.15a3.9 3.9 0 0 0 1.18-2.79v-7.86A2.38 2.38 0 0 0 22 7.235v-2.12a2.36 2.36 0 0 0-2.38-2.39m-5.08 13.37a.742.742 0 0 1 0 1.06a.71.71 0 0 1-.53.22a.741.741 0 0 1-.53-.22L12 15.685l-1.46 1.46a.71.71 0 0 1-.53.22a.739.739 0 0 1-.53-.22a.75.75 0 0 1 0-1.06l1.46-1.46l-1.46-1.46a.75.75 0 0 1 1.06-1.06l1.46 1.46l1.48-1.47a.75.75 0 0 1 1.06 0a.742.742 0 0 1 0 1.06l-1.47 1.47zm6-8.91a.861.861 0 0 1-.88.85H4.39a.87.87 0 0 1-.89-.85v-2.12a.86.86 0 0 1 .89-.84h15.23a.85.85 0 0 1 .88.84z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.62 2.725H4.39A2.37 2.37 0 0 0 2 5.065v2.12a2.38 2.38 0 0 0 1.64 2.29v7.86a3.9 3.9 0 0 0 1.18 2.79a4 4 0 0 0 2.83 1.15h8.71a4 4 0 0 0 3.699-2.427a3.94 3.94 0 0 0 .311-1.513v-7.86A2.38 2.38 0 0 0 22 7.235v-2.12a2.361 2.361 0 0 0-2.38-2.39m.88 4.46a.861.861 0 0 1-.88.85H4.39a.87.87 0 0 1-.89-.85v-2.12a.86.86 0 0 1 .89-.84h15.23a.85.85 0 0 1 .88.84z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxMinus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.382 8.813v8.5c0 .845.344 1.656.957 2.253a3.305 3.305 0 0 0 2.308.934h8.706c.866 0 1.696-.336 2.308-.934a3.15 3.15 0 0 0 .957-2.253v-8.5m0-5.313H4.382c-.901 0-1.632.714-1.632 1.594v2.125c0 .88.73 1.593 1.632 1.593h15.236c.901 0 1.632-.713 1.632-1.593V5.094c0-.88-.73-1.594-1.632-1.594M8.735 15.188h6.53"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxMinusFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.62 2.725H4.39A2.37 2.37 0 0 0 2 5.065v2.12a2.38 2.38 0 0 0 1.64 2.29v7.86a3.9 3.9 0 0 0 1.18 2.79a4 4 0 0 0 2.83 1.15h8.71a4 4 0 0 0 3.699-2.427c.202-.479.308-.993.311-1.513v-7.86A2.38 2.38 0 0 0 22 7.235v-2.12a2.361 2.361 0 0 0-2.38-2.39m-4.35 13.18H8.74a.75.75 0 0 1 0-1.5h6.53a.75.75 0 1 1 0 1.5m5.23-8.72a.861.861 0 0 1-.88.85H4.39a.87.87 0 0 1-.89-.85v-2.12a.86.86 0 0 1 .89-.84h15.23a.85.85 0 0 1 .88.84z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxPlus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M4.382 8.813v8.5c0 .845.344 1.656.957 2.253a3.305 3.305 0 0 0 2.308.934h8.706c.866 0 1.696-.336 2.308-.934a3.15 3.15 0 0 0 .957-2.253v-8.5m0-5.313H4.382c-.901 0-1.632.714-1.632 1.594v2.125c0 .88.73 1.593 1.632 1.593h15.236c.901 0 1.632-.713 1.632-1.593V5.094c0-.88-.73-1.594-1.632-1.594"/><path stroke-miterlimit="10" d="M11.995 12.156v5M9.5 14.662h5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxPlusFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.62 2.725H4.39A2.37 2.37 0 0 0 2 5.065v2.12a2.38 2.38 0 0 0 1.64 2.29v7.86a3.9 3.9 0 0 0 1.18 2.79a4 4 0 0 0 2.83 1.15h8.71a4 4 0 0 0 3.699-2.427c.202-.479.308-.993.311-1.513v-7.86A2.38 2.38 0 0 0 22 7.235v-2.12a2.361 2.361 0 0 0-2.38-2.39m-5.08 12.66h-1.75v1.74a.75.75 0 0 1-1.5 0v-1.74H9.54a.75.75 0 1 1 0-1.5h1.75v-1.76a.75.75 0 1 1 1.5 0v1.76h1.75a.75.75 0 0 1 0 1.5m6-8.2a.861.861 0 0 1-.88.85H4.39a.87.87 0 0 1-.89-.85v-2.12a.86.86 0 0 1 .89-.84h15.23a.85.85 0 0 1 .88.84z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxQuestionMark(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M4.382 8.813v8.5c0 .845.344 1.656.957 2.253a3.305 3.305 0 0 0 2.308.934h8.706c.866 0 1.696-.336 2.308-.934a3.15 3.15 0 0 0 .957-2.253v-8.5m0-5.313H4.382c-.901 0-1.632.714-1.632 1.594v2.125c0 .88.73 1.593 1.632 1.593h15.236c.901 0 1.632-.713 1.632-1.593V5.094c0-.88-.73-1.594-1.632-1.594"/><path stroke-miterlimit="10" d="M10.496 12.589a1.598 1.598 0 0 1 1.742-.906a1.55 1.55 0 0 1 1.137.81a1.345 1.345 0 0 1-.784 1.851a.994.994 0 0 0-.64.897v.37"/><path stroke-linejoin="round" d="M11.922 17.656h.004"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxQuestionMarkFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.62 2.725H4.39A2.37 2.37 0 0 0 2 5.065v2.12a2.38 2.38 0 0 0 1.64 2.29v7.86a3.9 3.9 0 0 0 1.18 2.79a4 4 0 0 0 2.83 1.15h8.71a4 4 0 0 0 2.83-1.15a3.9 3.9 0 0 0 1.18-2.79v-7.86A2.38 2.38 0 0 0 22 7.235v-2.12a2.36 2.36 0 0 0-2.38-2.39m-7.69 15.65a.75.75 0 1 1 0-1.5a.75.75 0 0 1 0 1.5m2.16-4.52a2.05 2.05 0 0 1-.51.72a2 2 0 0 1-.75.45l-.08.08a.21.21 0 0 0 0 .13v.35a.75.75 0 1 1-1.5 0v-.37c.005-.36.12-.708.33-1c.197-.284.476-.5.8-.62a.687.687 0 0 0 .16-.12a.56.56 0 0 0 .14-.2a.7.7 0 0 0 .05-.25a.5.5 0 0 0-.06-.24a.76.76 0 0 0-.23-.25a.9.9 0 0 0-.9-.06a.82.82 0 0 0-.38.4a.75.75 0 1 1-1.36-.63a2.36 2.36 0 0 1 2.56-1.33a2.32 2.32 0 0 1 1.66 1.2a2.2 2.2 0 0 1 .22.88a2.1 2.1 0 0 1-.15.86m6.41-6.67a.861.861 0 0 1-.88.85H4.39a.87.87 0 0 1-.89-.85v-2.12a.86.86 0 0 1 .89-.84h15.23a.85.85 0 0 1 .88.84z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxThreeD(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M10.55 2.876L4.595 6.182a2.978 2.978 0 0 0-1.529 2.611v6.414a2.978 2.978 0 0 0 1.529 2.61l5.957 3.307a2.978 2.978 0 0 0 2.898 0l5.957-3.306a2.978 2.978 0 0 0 1.529-2.611V8.793a2.978 2.978 0 0 0-1.529-2.61L13.45 2.876a2.978 2.978 0 0 0-2.898 0Z"/><path d="M20.33 6.996L12 12L3.67 6.996M12 21.49V12"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxThreeDcheck(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20.935 11.009V8.793a2.978 2.978 0 0 0-1.529-2.61l-5.957-3.307a2.978 2.978 0 0 0-2.898 0L4.594 6.182a2.978 2.978 0 0 0-1.529 2.611v6.414a2.978 2.978 0 0 0 1.529 2.61l5.957 3.307a2.978 2.978 0 0 0 2.898 0l2.522-1.4"/><path d="M20.33 6.996L12 12L3.67 6.996M12 21.49V12m5.466 3.996l1.408 1.407a.531.531 0 0 0 .757 0l2.835-2.836"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxThreeDcheckFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.954 8.804v2.21a.75.75 0 1 1-1.5 0v-2.21a1.998 1.998 0 0 0-.13-.76l-7.3 4.38v8.19c.115-.034.226-.081.33-.14l2.53-1.4a.75.75 0 0 1 1 .29a.75.75 0 0 1-.3 1l-2.52 1.4a3.72 3.72 0 0 1-3.62 0l-6-3.3a3.79 3.79 0 0 1-1.92-3.27v-6.39c0-.669.18-1.325.52-1.9c.058-.103.125-.2.2-.29a1 1 0 0 1 .12-.15a3.45 3.45 0 0 1 1.08-.93l6-3.31a3.81 3.81 0 0 1 3.62 0l6 3.31c.421.231.789.548 1.08.93l.12.15c.076.09.143.187.2.29a3.64 3.64 0 0 1 .49 1.9"/><path fill="currentColor" d="M18.524 18.264a1.331 1.331 0 0 1-.49-.09a1.37 1.37 0 0 1-.42-.29l-1.4-1.4a.737.737 0 0 1 0-1.06a.75.75 0 0 1 1.06 0l1.25 1.25l2.69-2.68a.75.75 0 0 1 1.06 1.06l-2.84 2.84a1.222 1.222 0 0 1-.41.27a1.242 1.242 0 0 1-.5.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxThreeDcross(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M20.935 11.009V8.793a2.978 2.978 0 0 0-1.529-2.61l-5.957-3.307a2.978 2.978 0 0 0-2.898 0L4.594 6.182a2.978 2.978 0 0 0-1.529 2.611v6.414a2.978 2.978 0 0 0 1.529 2.61l5.957 3.307a2.978 2.978 0 0 0 2.898 0l2.522-1.4"/><path stroke-linejoin="round" d="M20.33 6.996L12 12L3.67 6.996M12 21.49V12"/><path stroke-miterlimit="10" d="m22.935 14.745l-4 3.992m0-3.983l4 3.991"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxThreeDcrossFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.695 8.804v2.21a.75.75 0 0 1-1.5 0v-2.21a2 2 0 0 0-.13-.76l-7.3 4.38v8.19c.115-.034.226-.081.33-.14l2.53-1.4a.75.75 0 0 1 1 .29a.75.75 0 0 1-.3 1l-2.52 1.4a3.72 3.72 0 0 1-3.62 0l-6-3.3a3.79 3.79 0 0 1-1.92-3.27v-6.39c0-.669.18-1.325.52-1.9c.058-.103.125-.2.2-.29a1 1 0 0 1 .12-.15a3.45 3.45 0 0 1 1.08-.93l6-3.31a3.81 3.81 0 0 1 3.62 0l6 3.31c.421.231.789.548 1.08.93l.12.15c.076.09.143.187.2.29a3.64 3.64 0 0 1 .49 1.9"/><path fill="currentColor" d="M22.445 18.174a.75.75 0 1 1-1.06 1.06l-1.45-1.43l-1.49 1.42a.708.708 0 0 1-.53.22a.74.74 0 0 1-.53-.22a.75.75 0 0 1 0-1.06l1.46-1.46l-1.4-1.46a.75.75 0 0 1 0-1.06a.742.742 0 0 1 1.06 0l1.46 1.46l1.48-1.47a.76.76 0 0 1 1.07 0a.75.75 0 0 1 0 1.06l-1.47 1.47z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxThreeDdownload(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M20.935 11.009V8.793a2.978 2.978 0 0 0-1.529-2.61l-5.957-3.307a2.978 2.978 0 0 0-2.898 0L4.594 6.182a2.978 2.978 0 0 0-1.529 2.611v6.414a2.978 2.978 0 0 0 1.529 2.61l5.957 3.307a2.978 2.978 0 0 0 2.898 0l2.522-1.4"/><path stroke-linejoin="round" d="M20.33 6.996L12 12L3.67 6.996M12 21.49V12"/><path stroke-miterlimit="10" d="M19.97 19.245v-5"/><path stroke-linejoin="round" d="m17.676 17.14l1.968 1.968a.46.46 0 0 0 .65 0l1.968-1.968"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxThreeDdownloadFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.086 8.804v2.21a.75.75 0 1 1-1.5 0v-2.21a1.998 1.998 0 0 0-.13-.76l-7.3 4.38v8.19c.115-.034.225-.081.33-.14l2.53-1.4a.75.75 0 0 1 1 .29a.75.75 0 0 1-.3 1l-2.52 1.4a3.72 3.72 0 0 1-3.62 0l-6-3.3a3.79 3.79 0 0 1-1.92-3.27v-6.39c0-.669.18-1.325.52-1.9c.057-.103.124-.2.2-.29l.12-.15a3.45 3.45 0 0 1 1.08-.93l6-3.31a3.81 3.81 0 0 1 3.62 0l6 3.31c.42.231.788.548 1.08.93a.996.996 0 0 1 .12.15c.075.09.142.187.2.29a3.64 3.64 0 0 1 .49 1.9"/><path fill="currentColor" d="m22.196 17.624l-2 2a1.18 1.18 0 0 1-.39.26a1.06 1.06 0 0 1-.46.1c-.158 0-.314-.03-.46-.09a1.28 1.28 0 0 1-.4-.27l-2-2a.74.74 0 0 1 0-1.06a.75.75 0 0 1 1.06 0l1 1v-3.36a.75.75 0 0 1 1.5 0v3.38l1-1a.748.748 0 0 1 1.079-.02a.752.752 0 0 1-.02 1.08z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxThreeDfill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.22 6.894a3.68 3.68 0 0 0-1.4-1.37l-6-3.31a3.83 3.83 0 0 0-3.63 0l-6 3.31a3.68 3.68 0 0 0-1.4 1.37a3.74 3.74 0 0 0-.52 1.9v6.41a3.79 3.79 0 0 0 1.92 3.27l6 3.3a3.74 3.74 0 0 0 3.63 0l6-3.31a3.72 3.72 0 0 0 1.91-3.26v-6.36a3.64 3.64 0 0 0-.51-1.95m-1 8.31a2.2 2.2 0 0 1-1.14 1.95l-6 3.31a1.62 1.62 0 0 1-.33.14v-8.18l7.3-4.39c.092.242.136.5.13.76z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxThreeDminus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M20.935 11.009V8.793a2.978 2.978 0 0 0-1.529-2.61l-5.957-3.307a2.978 2.978 0 0 0-2.898 0L4.594 6.182a2.978 2.978 0 0 0-1.529 2.611v6.414a2.978 2.978 0 0 0 1.529 2.61l5.957 3.307a2.978 2.978 0 0 0 2.898 0l2.522-1.4"/><path stroke-linejoin="round" d="M20.33 6.996L12 12L3.67 6.996M12 21.49V12"/><path stroke-miterlimit="10" d="M17.409 16.47h5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxThreeDminusFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.945 8.804v2.21a.75.75 0 1 1-1.5 0v-2.21a2 2 0 0 0-.13-.76l-7.3 4.38v8.19a1.56 1.56 0 0 0 .33-.14l2.53-1.4a.75.75 0 0 1 1 .29a.75.75 0 0 1-.3 1l-2.52 1.4a3.7 3.7 0 0 1-1.81.47a3.77 3.77 0 0 1-1.82-.47l-5.95-3.3a3.79 3.79 0 0 1-1.92-3.27v-6.39c0-.669.18-1.325.52-1.9c.057-.103.124-.2.2-.29l.12-.15a3.45 3.45 0 0 1 1.08-.93l6-3.31a3.81 3.81 0 0 1 3.62 0l6 3.31c.42.231.789.548 1.08.93a.963.963 0 0 1 .12.15c.076.09.143.187.2.29c.32.58.476 1.237.45 1.9"/><path fill="currentColor" d="M21.695 17.174h-5a.75.75 0 1 1 0-1.5h5a.75.75 0 1 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxThreeDnotification(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M20.935 11.009V8.793a2.978 2.978 0 0 0-1.529-2.61l-5.957-3.307a2.978 2.978 0 0 0-2.898 0L4.594 6.182a2.978 2.978 0 0 0-1.529 2.611v6.414a2.978 2.978 0 0 0 1.529 2.61l5.957 3.307a2.978 2.978 0 0 0 2.898 0l2.522-1.4"/><path stroke-linecap="round" stroke-linejoin="round" d="M20.33 6.996L12 12L3.67 6.996M12 21.49V12"/><circle cx="20.329" cy="16.501" r="2.376"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxThreeDnotificationFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.84 8.804v2.21a.75.75 0 1 1-1.5 0v-2.21a2 2 0 0 0-.13-.76l-7.3 4.38v8.19a1.56 1.56 0 0 0 .33-.14l2.53-1.4a.76.76 0 0 1 1 .29a.75.75 0 0 1-.3 1l-2.52 1.4a3.72 3.72 0 0 1-3.62 0l-6-3.3a3.79 3.79 0 0 1-1.92-3.27v-6.39c0-.669.18-1.325.52-1.9c.057-.103.124-.2.2-.29a1 1 0 0 1 .12-.15a3.45 3.45 0 0 1 1.08-.93l6-3.31a3.81 3.81 0 0 1 3.62 0l6 3.31c.42.231.789.548 1.08.93l.12.15c.076.09.143.187.2.29a3.64 3.64 0 0 1 .49 1.9"/><path fill="currentColor" d="M22.59 16.454a3.13 3.13 0 1 1-6.26.02a3.13 3.13 0 0 1 6.26-.02"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxThreeDplus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M20.935 11.009V8.793a2.978 2.978 0 0 0-1.529-2.61l-5.957-3.307a2.978 2.978 0 0 0-2.898 0L4.594 6.182a2.978 2.978 0 0 0-1.529 2.611v6.414a2.978 2.978 0 0 0 1.529 2.61l5.957 3.307a2.978 2.978 0 0 0 2.898 0l2.522-1.4"/><path stroke-linejoin="round" d="M20.33 6.996L12 12L3.67 6.996M12 21.49V12"/><path stroke-miterlimit="10" d="M19.903 13.965v5m-2.494-2.495h5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxThreeDplusFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.955 8.804v2.21a.75.75 0 1 1-1.5 0v-2.21a2 2 0 0 0-.13-.76l-7.3 4.38v8.19a1.56 1.56 0 0 0 .33-.14l2.53-1.4a.76.76 0 0 1 1 .29a.75.75 0 0 1-.3 1l-2.52 1.4a3.7 3.7 0 0 1-1.81.47a3.77 3.77 0 0 1-1.82-.47l-5.95-3.3a3.79 3.79 0 0 1-1.92-3.27v-6.39c0-.669.18-1.325.52-1.9c.057-.103.124-.2.2-.29a3.56 3.56 0 0 1 1.2-1.08l6-3.31a3.81 3.81 0 0 1 3.62 0l6 3.31c.42.231.789.548 1.08.93a.963.963 0 0 1 .12.15c.076.09.143.187.2.29c.32.58.476 1.237.45 1.9"/><path fill="currentColor" d="M22.435 16.424a.75.75 0 0 1-.75.75h-1.76v1.75a.75.75 0 1 1-1.5 0v-1.75h-1.72a.75.75 0 1 1 0-1.5h1.74v-1.75a.75.75 0 1 1 1.5 0v1.75h1.76a.76.76 0 0 1 .73.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxThreeDquestionMark(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M20.935 11.009V8.793a2.978 2.978 0 0 0-1.529-2.61l-5.957-3.307a2.978 2.978 0 0 0-2.898 0L4.594 6.182a2.978 2.978 0 0 0-1.529 2.611v6.414a2.978 2.978 0 0 0 1.529 2.61l5.957 3.307a2.978 2.978 0 0 0 2.898 0l2.522-1.4"/><path stroke-linejoin="round" d="M20.33 6.996L12 12L3.67 6.996M12 21.49V12"/><path stroke-miterlimit="10" d="M18.15 14.714a1.332 1.332 0 0 1 1.452-.755a1.291 1.291 0 0 1 .948.675a1.121 1.121 0 0 1-.654 1.542a.828.828 0 0 0-.533.748v.309"/><path stroke-linejoin="round" d="M19.34 18.937h.002"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxThreeDquestionMarkFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.715 8.804v2.21a.75.75 0 1 1-1.5 0v-2.21a2 2 0 0 0-.13-.76l-7.3 4.38v8.19a1.56 1.56 0 0 0 .33-.14l2.53-1.4a.75.75 0 0 1 1 .29a.75.75 0 0 1-.3 1l-2.52 1.4a3.72 3.72 0 0 1-3.62 0l-6-3.3a3.79 3.79 0 0 1-1.92-3.27v-6.39c0-.669.18-1.325.52-1.9c.057-.103.124-.2.2-.29a1 1 0 0 1 .12-.15a3.45 3.45 0 0 1 1.08-.93l6-3.31a3.81 3.81 0 0 1 3.62 0l6 3.31c.42.231.788.548 1.08.93l.12.15c.075.09.142.187.2.29a3.64 3.64 0 0 1 .49 1.9"/><path fill="currentColor" d="M19.395 17.944a.76.76 0 0 1-.75-.75v-.31c.004-.32.105-.63.29-.89a1.6 1.6 0 0 1 .73-.56l.16-.09a.378.378 0 0 0 .09-.13c.01-.05.01-.1 0-.15a.468.468 0 0 0 0-.16a.518.518 0 0 0-.15-.15a.6.6 0 0 0-.89.23a.75.75 0 1 1-1.36-.63a2.11 2.11 0 0 1 2.27-1.18a2 2 0 0 1 .85.36a1.89 1.89 0 0 1 .82 1.49c.002.267-.049.532-.15.78a1.798 1.798 0 0 1-.45.64c-.192.18-.421.316-.67.4v.35a.748.748 0 0 1-.79.75m-.02 1.7a.75.75 0 1 1 0-1.5a.75.75 0 0 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxThreeDscan(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M11.008 5.758L6.933 8.02a2.038 2.038 0 0 0-1.046 1.786v4.388a2.038 2.038 0 0 0 1.046 1.786l4.075 2.262a2.038 2.038 0 0 0 1.984 0l4.075-2.262a2.04 2.04 0 0 0 1.046-1.786V9.806a2.038 2.038 0 0 0-1.046-1.786l-4.075-2.262a2.038 2.038 0 0 0-1.984 0Z"/><path d="M17.699 8.577L12 12L6.301 8.577M12 18.494V12"/><path stroke-linecap="round" stroke-linejoin="round" d="M21.25 7.889V5.833a3.083 3.083 0 0 0-3.083-3.083h-3.084m0 18.5h3.084a3.083 3.083 0 0 0 3.083-3.083V16.11m-18.5.001v2.056a3.083 3.083 0 0 0 3.083 3.083h3.084m0-18.5H5.833A3.083 3.083 0 0 0 2.75 5.833V7.89"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxThreeDscanFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.48 8.38a2.783 2.783 0 0 0-1.05-1.002l-4.07-2.275a2.845 2.845 0 0 0-2.72 0L6.57 7.378a2.783 2.783 0 0 0-1.05 1.003c-.25.432-.381.924-.38 1.423v4.391c-.001.5.13.991.38 1.424c.253.423.616.77 1.05 1.002l4.07 2.266a2.815 2.815 0 0 0 2.72 0l4.07-2.266c.434-.233.797-.58 1.05-1.002c.249-.43.38-.917.38-1.414v-4.36a2.833 2.833 0 0 0-.38-1.464m-1.12 5.825c0 .229-.058.453-.17.652c-.119.2-.288.367-.49.481l-3.95 2.196V12.43l4.6-2.777c.005.05.005.1 0 .15zm3.89-5.575a.749.749 0 0 1-.75-.751V5.875a2.322 2.322 0 0 0-.68-1.655a2.347 2.347 0 0 0-1.65-.691h-3.09a.749.749 0 0 1-.75-.752a.753.753 0 0 1 .75-.752h3.09a3.845 3.845 0 0 1 2.71 1.133A3.81 3.81 0 0 1 22 5.875V7.88a.753.753 0 0 1-.75.751m-3.08 13.394h-3.09a.749.749 0 0 1-.75-.752a.753.753 0 0 1 .75-.752h3.09a2.316 2.316 0 0 0 2.155-1.44c.117-.284.176-.589.175-.896V16.12a.753.753 0 0 1 .75-.752a.749.749 0 0 1 .75.752v2.065a3.838 3.838 0 0 1-1.119 2.718a3.82 3.82 0 0 1-2.711 1.122m-9.25 0H5.83a3.81 3.81 0 0 1-2.711-1.122A3.83 3.83 0 0 1 2 18.185V16.12a.753.753 0 0 1 .75-.752a.75.75 0 0 1 .75.752v2.065a2.33 2.33 0 0 0 1.437 2.161a2.3 2.3 0 0 0 .893.175h3.09a.75.75 0 0 1 .75.752a.753.753 0 0 1-.75.752M2.75 8.631A.749.749 0 0 1 2 7.88V5.875a3.818 3.818 0 0 1 1.12-2.717a3.845 3.845 0 0 1 2.71-1.133h3.09a.749.749 0 0 1 .75.752a.753.753 0 0 1-.75.752H5.83a2.347 2.347 0 0 0-2.156 1.45a2.32 2.32 0 0 0-.174.896V7.88a.753.753 0 0 1-.75.751"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxThreeDupload(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M20.935 11.009V8.793a2.978 2.978 0 0 0-1.529-2.61l-5.957-3.307a2.978 2.978 0 0 0-2.898 0L4.594 6.182a2.978 2.978 0 0 0-1.529 2.611v6.414a2.978 2.978 0 0 0 1.529 2.61l5.957 3.307a2.978 2.978 0 0 0 2.898 0l2.522-1.4"/><path stroke-linejoin="round" d="M20.33 6.996L12 12L3.67 6.996M12 21.49V12"/><path stroke-miterlimit="10" d="M19.97 14.245v5"/><path stroke-linejoin="round" d="m22.262 16.35l-1.967-1.967a.46.46 0 0 0-.652 0l-1.967 1.967"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxThreeDuploadFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.026 8.804v2.21a.75.75 0 1 1-1.5 0v-2.21a2 2 0 0 0-.13-.76l-7.3 4.38v8.19c.115-.034.226-.081.33-.14l2.53-1.4a.75.75 0 0 1 1 .29a.75.75 0 0 1-.3 1l-2.52 1.4a3.72 3.72 0 0 1-3.62 0l-6-3.3a3.79 3.79 0 0 1-1.92-3.27v-6.39c0-.669.18-1.325.52-1.9c.057-.103.124-.2.2-.29a3.56 3.56 0 0 1 1.2-1.08l6-3.31a3.81 3.81 0 0 1 3.62 0l6 3.31c.42.231.789.548 1.08.93a.963.963 0 0 1 .12.15c.076.09.143.187.2.29a3.64 3.64 0 0 1 .49 1.9"/><path fill="currentColor" d="M22.136 16.804a.738.738 0 0 1-1.06 0l-1-1v3.38a.75.75 0 0 1-1.5 0v-3.38l-1 1a.739.739 0 0 1-.53.22a.71.71 0 0 1-.53-.22a.74.74 0 0 1 0-1.06l2-2a1.36 1.36 0 0 1 .39-.27c.1-.041.203-.068.31-.08h.32c.107.014.211.045.31.09a1 1 0 0 1 .38.26l2 2a.748.748 0 0 1-.09 1.06"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxUpload(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M4.382 8.813v8.5c0 .845.344 1.656.957 2.253a3.305 3.305 0 0 0 2.308.934h8.706c.866 0 1.696-.336 2.308-.934a3.15 3.15 0 0 0 .957-2.253v-8.5m0-5.313H4.382c-.901 0-1.632.714-1.632 1.594v2.125c0 .88.73 1.593 1.632 1.593h15.236c.901 0 1.632-.713 1.632-1.593V5.094c0-.88-.73-1.594-1.632-1.594"/><path stroke-miterlimit="10" d="M12 12v5"/><path stroke-linejoin="round" d="m14.293 14.105l-1.967-1.967a.458.458 0 0 0-.652 0l-1.967 1.967"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxUploadFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.62 2.725H4.39A2.37 2.37 0 0 0 2 5.065v2.12a2.38 2.38 0 0 0 1.64 2.29v7.86a3.9 3.9 0 0 0 1.18 2.79a4 4 0 0 0 2.83 1.15h8.71a4 4 0 0 0 2.83-1.15a3.9 3.9 0 0 0 1.18-2.79v-7.86A2.38 2.38 0 0 0 22 7.235v-2.12a2.361 2.361 0 0 0-2.38-2.39m-4.79 11.88a.741.741 0 0 1-.53.22a.712.712 0 0 1-.53-.22l-1-1v3.38a.75.75 0 1 1-1.5 0v-3.38l-1 1a.75.75 0 0 1-1.06-1.06l2-2c.11-.112.244-.2.39-.26a1.11 1.11 0 0 1 .3-.08a.86.86 0 0 1 .32 0c.107.011.211.038.31.08c.145.062.277.15.39.26l2 2a.75.75 0 0 1-.09 1.06m5.67-7.42a.861.861 0 0 1-.88.85H4.39a.87.87 0 0 1-.89-.85v-2.12a.86.86 0 0 1 .89-.84h15.23a.85.85 0 0 1 .88.84z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Briefcase(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.25 8.912v1.544a1.03 1.03 0 0 1-1.028 1.03l-6.166 1.543v-.514a1.03 1.03 0 0 0-1.028-1.03h-2.056a1.027 1.027 0 0 0-1.028 1.03v.514l-6.166-1.544a1.027 1.027 0 0 1-1.028-1.03V8.913a2.576 2.576 0 0 1 2.57-2.574h13.36a2.575 2.575 0 0 1 2.57 2.574"/><path d="m3.778 11.485l.36 7.288a2.204 2.204 0 0 0 2.178 1.977h11.368a2.197 2.197 0 0 0 2.178-1.977l.36-7.288"/><path d="M13.028 11.485h-2.056a1.03 1.03 0 0 0-1.028 1.03v1.03c0 .568.46 1.028 1.028 1.028h2.056c.567 0 1.028-.46 1.028-1.029v-1.03c0-.568-.46-1.029-1.028-1.029m2.055-5.147V4.28a1.03 1.03 0 0 0-1.027-1.029H9.944a1.027 1.027 0 0 0-1.027 1.03v2.058"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BriefcaseFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.286 12.548v.997a.279.279 0 0 1-.269.28h-2.054a.29.29 0 0 1-.279-.28v-.997a.354.354 0 0 1 0-.11a.218.218 0 0 1 .07-.09a.27.27 0 0 1 .19-.079h2.083a.268.268 0 0 1 .26.28"/><path fill="currentColor" d="M21.003 6.616a3.31 3.31 0 0 0-2.343-.997h-2.84V4.313a1.725 1.725 0 0 0-.519-1.246a1.756 1.756 0 0 0-1.256-.529H9.946a1.785 1.785 0 0 0-1.774 1.775v1.306H5.34a3.33 3.33 0 0 0-2.343.997A3.26 3.26 0 0 0 2 8.96v1.495a1.745 1.745 0 0 0 .519 1.257c.153.154.336.276.538.359l.329 6.75a2.92 2.92 0 0 0 2.911 2.641h11.336c.737.01 1.45-.26 1.994-.757a2.99 2.99 0 0 0 .997-1.925l.33-6.71c.2-.085.382-.206.537-.358A1.8 1.8 0 0 0 22 10.455V8.96a3.299 3.299 0 0 0-.997-2.343M9.667 4.333a.28.28 0 0 1 .09-.19a.25.25 0 0 1 .19-.09h4.077a.25.25 0 0 1 .2.09c.05.05.079.119.08.19v1.306H9.646zm5.125 9.242a.17.17 0 0 1 0 .07a1.755 1.755 0 0 1-1.755 1.705h-2.054a1.765 1.765 0 0 1-1.765-1.705a.17.17 0 0 1 0-.07v-.997c0-.145.02-.29.06-.429a1.595 1.595 0 0 1 .23-.548a1.34 1.34 0 0 1 .209-.26c.247-.242.558-.408.897-.478c.12-.014.24-.014.36 0h2.053c.12-.014.24-.014.359 0c.336.07.644.237.887.479c.08.077.15.164.21.26c.06.09.11.187.15.288c.091.218.139.452.139.688z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Broadcast(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.937 18.938A9.15 9.15 0 0 1 2.75 12a9.15 9.15 0 0 1 3.187-6.939M8.4 7.617a4.824 4.824 0 0 0-.815.809A5.617 5.617 0 0 0 6.299 12c0 1.303.454 2.566 1.286 3.575c.238.3.512.572.814.808M14.203 12c-.01.405-.131.8-.352 1.14a2.206 2.206 0 0 1-3.758 0a2.166 2.166 0 0 1 0-2.28a2.206 2.206 0 0 1 3.758 0c.221.341.343.735.352 1.14m1.397 4.383a4.83 4.83 0 0 0 .815-.808A5.617 5.617 0 0 0 17.701 12a5.617 5.617 0 0 0-1.286-3.574a4.821 4.821 0 0 0-.814-.809m2.462-2.554A9.15 9.15 0 0 1 21.25 12a9.15 9.15 0 0 1-3.187 6.937"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BroadcastFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.935 19.687a.74.74 0 0 1-.49-.19a9.252 9.252 0 0 1-1-1a9.87 9.87 0 0 1 0-13a9.248 9.248 0 0 1 1-1a.75.75 0 0 1 1.06.07a.76.76 0 0 1-.07 1.06a8.631 8.631 0 0 0-.86.85a8.41 8.41 0 0 0 0 11c.27.297.556.578.86.84a.75.75 0 0 1-.5 1.32z"/><path fill="currentColor" d="M8.405 17.127a.76.76 0 0 1-.47-.16a5.122 5.122 0 0 1-.93-.93a6.37 6.37 0 0 1 0-8.09a5.13 5.13 0 0 1 .94-.92a.74.74 0 0 1 1.05.13a.75.75 0 0 1-.13 1a4 4 0 0 0-.69.68a4.88 4.88 0 0 0 0 6.21c.205.257.44.488.7.69a.76.76 0 0 1 .13 1.06a.77.77 0 0 1-.6.33"/><path fill="currentColor" d="M14.955 11.997a2.9 2.9 0 0 1-.47 1.53a3 3 0 0 1-1 .95a2.9 2.9 0 0 1-1.54.44a3 3 0 0 1-1.54-.43a3.1 3.1 0 0 1-1-1a2.93 2.93 0 0 1 0-3.07a3 3 0 0 1 4.06-1a3.1 3.1 0 0 1 1 1a2.84 2.84 0 0 1 .49 1.58"/><path fill="currentColor" d="M15.605 17.127a.77.77 0 0 1-.59-.28a.75.75 0 0 1 .13-1.06a4 4 0 0 0 .68-.68a4.85 4.85 0 0 0 0-6.21a4.501 4.501 0 0 0-.69-.69a.75.75 0 1 1 .92-1.18c.352.268.668.58.94.93a6.35 6.35 0 0 1 0 8.09a5.073 5.073 0 0 1-.93.92a.74.74 0 0 1-.46.16"/><path fill="currentColor" d="M18.065 19.687a.75.75 0 0 1-.49-1.32c.303-.263.587-.547.85-.85a8.38 8.38 0 0 0 0-11a9.626 9.626 0 0 0-.85-.84a.76.76 0 0 1-.08-1.06a.75.75 0 0 1 1.06-.07a9.35 9.35 0 0 1 1 1a9.87 9.87 0 0 1 0 13a9.348 9.348 0 0 1-1 1a.7.7 0 0 1-.49.14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuildingA(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8.531 3.212h6.938a2.775 2.775 0 0 1 2.775 2.775v14.8H5.756v-14.8a2.775 2.775 0 0 1 2.775-2.775M2.75 20.788h18.5"/><path d="M11.075 14.313h1.85a1.387 1.387 0 0 1 1.387 1.387v5.088H9.689V15.7a1.387 1.387 0 0 1 1.387-1.387m-1.851-7.4h5.55m-5.55 3.7h5.55"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuildingAfill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.2 20.035h-2.21V5.985a3.52 3.52 0 0 0-3.52-3.52H8.53A3.53 3.53 0 0 0 5 5.985v14.05H2.75a.75.75 0 1 0 0 1.5h18.5a.75.75 0 0 0 0-1.5zm-7.69 0h-3.12v-4.34a.64.64 0 0 1 .18-.45a.63.63 0 0 1 .45-.18h1.85a.63.63 0 0 1 .64.63zm1.21-8.67H9.17a.75.75 0 0 1 0-1.5h5.55a.75.75 0 0 1 0 1.5m0-3.7H9.17a.75.75 0 0 1 0-1.5h5.55a.75.75 0 0 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuildingB(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M5.978 3.212h6.938a2.775 2.775 0 0 1 2.775 2.775v14.8H3.203v-14.8a2.775 2.775 0 0 1 2.775-2.775M2.75 20.788h18.5"/><path d="M8.531 14.313h1.85A1.388 1.388 0 0 1 11.77 15.7v5.088H7.144V15.7a1.387 1.387 0 0 1 1.387-1.387m-1.859-7.4h5.55m-5.55 3.7h5.55m3.468-1.388h1.85A2.775 2.775 0 0 1 20.317 12v8.788"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuildingBfill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.27 19.995h-.21v-8a3.52 3.52 0 0 0-3.52-3.52h-1.1v-2.49a3.52 3.52 0 0 0-3.53-3.52H5.98a3.52 3.52 0 0 0-3.53 3.52v14.12a.73.73 0 0 0-.45.68a.74.74 0 0 0 .75.75h18.5a.75.75 0 1 0 0-1.5zm-15.33-9.42a.75.75 0 0 1 .75-.75h5.55a.75.75 0 1 1 0 1.5H6.69a.76.76 0 0 1-.75-.75m.75-4.45h5.55a.75.75 0 1 1 0 1.5H6.69a.75.75 0 1 1 0-1.5m1.22 13.87v-4.34a.628.628 0 0 1 .64-.63h1.85a.63.63 0 0 1 .64.63v4.34zm11.67 0h-3.12V9.935h1.1a2 2 0 0 1 2 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuildingTree(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.206 3.41h5.263a2.713 2.713 0 0 1 2.712 2.713V20.59H6.493V6.123A2.713 2.713 0 0 1 9.206 3.41m-6.04 17.18H21.25"/><path d="M11.693 14.26h.325a1.356 1.356 0 0 1 1.356 1.357v4.973h-3.038v-4.973a1.356 1.356 0 0 1 1.357-1.357M9.884 7.027h3.906m-3.906 3.617h3.906m3.391-1.357h.687A2.713 2.713 0 0 1 20.581 12v8.59m-16.475 0v-6.782m1.357-2.712a1.356 1.356 0 1 0-2.713 0v1.356a1.356 1.356 0 1 0 2.713 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BuildingTreeFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.345 19.884V12.02a3.476 3.476 0 0 0-3.406-3.466V6.139a3.476 3.476 0 0 0-3.466-3.476h-5.28A3.447 3.447 0 0 0 6.75 3.665a3.416 3.416 0 0 0-1.002 2.454v3.647A2.104 2.104 0 0 0 2 11.098v1.343a2.074 2.074 0 0 0 1.353 1.953v5.44h-.19a.752.752 0 0 0 0 1.503h18.122a.751.751 0 0 0 .08-1.493zm-16.65-7.443a.61.61 0 0 1-1.212 0v-1.363a.612.612 0 1 1 1.212 0zm1.002 7.393h-.892v-5.39c.351-.126.66-.348.892-.64zm8.395 0H9.544v-3.186a2.124 2.124 0 0 1 2.114-2.113h.33a2.105 2.105 0 0 1 1.493.62c.391.4.61.935.611 1.493zm-.33-8.455H9.843a.752.752 0 0 1 0-1.503h3.918a.752.752 0 1 1 0 1.503m0-3.627H9.843a.752.752 0 1 1 0-1.502h3.918a.752.752 0 1 1 0 1.502m6.01 12.082H17.87v-9.808a1.9 1.9 0 0 1 1.322.572c.366.368.574.863.582 1.382z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calculator(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16 2.75H8a4 4 0 0 0-4 4v10.5a4 4 0 0 0 4 4h8a4 4 0 0 0 4-4V6.75a4 4 0 0 0-4-4"/><path d="M16.083 6H7.917C7.411 6 7 6.448 7 7v2c0 .552.41 1 .917 1h8.166c.506 0 .917-.448.917-1V7c0-.552-.41-1-.917-1M7.5 13h1m3 0h1m3 0h1m-9 2.5h1m3 0h1m3 0h1m-9 2.5h1m3 0h1m3 0h1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalculatorFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 2H8a4.75 4.75 0 0 0-4.75 4.75v10.5A4.75 4.75 0 0 0 8 22h8a4.76 4.76 0 0 0 4.75-4.75V6.75A4.76 4.76 0 0 0 16 2M8.5 18.75h-1a.75.75 0 1 1 0-1.5h1a.75.75 0 1 1 0 1.5m0-2.5h-1a.75.75 0 1 1 0-1.5h1a.75.75 0 1 1 0 1.5m0-2.5h-1a.75.75 0 1 1 0-1.5h1a.75.75 0 1 1 0 1.5m4 5h-1a.75.75 0 1 1 0-1.5h1a.75.75 0 1 1 0 1.5m0-2.5h-1a.75.75 0 1 1 0-1.5h1a.75.75 0 1 1 0 1.5m0-2.5h-1a.75.75 0 1 1 0-1.5h1a.75.75 0 1 1 0 1.5m4 5h-1a.75.75 0 1 1 0-1.5h1a.75.75 0 1 1 0 1.5m0-2.5h-1a.75.75 0 1 1 0-1.5h1a.75.75 0 1 1 0 1.5m0-2.5h-1a.75.75 0 1 1 0-1.5h1a.75.75 0 1 1 0 1.5m1.25-5.36a1.711 1.711 0 0 1-1.67 1.75H7.91a1.7 1.7 0 0 1-1.66-1.75v-2a1.71 1.71 0 0 1 1.66-1.75h8.17a1.72 1.72 0 0 1 1.67 1.75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 4.625H7a4 4 0 0 0-4 4v8.75a4 4 0 0 0 4 4h10a4 4 0 0 0 4-4v-8.75a4 4 0 0 0-4-4m-14 6h18m-4-8v4m-10-4v4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarCheck(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17 4.625H7a4 4 0 0 0-4 4v8.75a4 4 0 0 0 4 4h10a4 4 0 0 0 4-4v-8.75a4 4 0 0 0-4-4m-14 5h18m-4-7v4m-10-4v4"/><path d="m9 14.714l1.689 1.689a.639.639 0 0 0 .908 0L15 13"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarCheckFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.75 3.935v-1.31a.75.75 0 1 0-1.5 0v1.25h-8.5v-1.25a.75.75 0 0 0-1.5 0v1.31a4.76 4.76 0 0 0-4 4.69v8.75A4.75 4.75 0 0 0 7 22.125h10a4.75 4.75 0 0 0 4.75-4.75v-8.75a4.76 4.76 0 0 0-4-4.69m-2.22 9.6l-3.4 3.4a1.37 1.37 0 0 1-.45.3a1.24 1.24 0 0 1-.54.11a1.23 1.23 0 0 1-.53-.11a1.31 1.31 0 0 1-.46-.3l-1.68-1.69a.75.75 0 1 1 1.06-1.06l1.62 1.62l3.32-3.33a.75.75 0 0 1 1.06 1.06m4.72-4.66H3.75v-.25a3.24 3.24 0 0 1 2.5-3.15v1.15a.75.75 0 1 0 1.5 0v-1.25h8.5v1.25a.75.75 0 1 0 1.5 0v-1.15a3.24 3.24 0 0 1 2.5 3.15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarCross(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M17 4.625H7a4 4 0 0 0-4 4v8.75a4 4 0 0 0 4 4h10a4 4 0 0 0 4-4v-8.75a4 4 0 0 0-4-4m-14 5h18m-4-7v4m-10-4v4"/><path stroke-miterlimit="10" d="m14 13l-4 3.991m0-3.982L14 17"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarCrossFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.75 3.935v-1.31a.75.75 0 1 0-1.5 0v1.25h-8.5v-1.25a.75.75 0 0 0-1.5 0v1.31a4.76 4.76 0 0 0-4 4.69v8.75A4.75 4.75 0 0 0 7 22.125h10a4.75 4.75 0 0 0 4.75-4.75v-8.75a4.76 4.76 0 0 0-4-4.69m-3.22 12.54a.75.75 0 0 1-1.06 1.06l-1.48-1.47l-1.46 1.46a.75.75 0 0 1-1.06-1.06l1.46-1.46l-1.46-1.46a.75.75 0 0 1 0-1.06a.74.74 0 0 1 1.06 0l1.46 1.46l1.48-1.47a.75.75 0 0 1 1.06 1.06l-1.47 1.47zm5.72-7.6H3.75v-.25a3.24 3.24 0 0 1 2.5-3.15v1.15a.75.75 0 1 0 1.5 0v-1.25h8.5v1.25a.75.75 0 1 0 1.5 0v-1.15a3.24 3.24 0 0 1 2.5 3.15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarDownload(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M17 4.625H7a4 4 0 0 0-4 4v8.75a4 4 0 0 0 4 4h10a4 4 0 0 0 4-4v-8.75a4 4 0 0 0-4-4m-14 5h18m-4-7v4m-10-4v4"/><path stroke-miterlimit="10" d="M12 17.468v-5"/><path stroke-linejoin="round" d="m9.707 15.363l1.967 1.967a.458.458 0 0 0 .652 0l1.967-1.967"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarDownloadFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.5 4.06V2.75a.75.75 0 1 0-1.5 0V4H5.5V2.75a.75.75 0 0 0-1.5 0v1.31a4.76 4.76 0 0 0-4 4.69v8.75a4.75 4.75 0 0 0 4.75 4.75h10a4.75 4.75 0 0 0 4.75-4.75V8.75a4.76 4.76 0 0 0-4-4.69m-2.93 12l-2 2a1.399 1.399 0 0 1-.39.26a1.673 1.673 0 0 1-.3.08h-.3a1.382 1.382 0 0 1-.29-.08a1.16 1.16 0 0 1-.39-.26l-2-2a.74.74 0 0 1 0-1.06a.75.75 0 0 1 1.06 0l1 1v-3.37a.75.75 0 1 1 1.5 0V16l1-1a.75.75 0 0 1 1.06 0a.74.74 0 0 1 .06 1.02zm5.43-7H1.5v-.25A3.24 3.24 0 0 1 4 5.66v1.15a.75.75 0 1 0 1.5 0V5.56H14v1.25a.75.75 0 1 0 1.5 0V5.66A3.24 3.24 0 0 1 18 8.81z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.75 3.935v-1.31a.75.75 0 1 0-1.5 0v1.25h-8.5v-1.25a.75.75 0 1 0-1.5 0v1.31a4.76 4.76 0 0 0-4 4.69v8.75A4.75 4.75 0 0 0 7 22.125h10a4.75 4.75 0 0 0 4.75-4.75v-8.75a4.76 4.76 0 0 0-4-4.69m2.5 5.94H3.75v-1.25a3.24 3.24 0 0 1 2.5-3.15v1.15a.75.75 0 0 0 1.5 0v-1.25h8.5v1.25a.75.75 0 1 0 1.5 0v-1.15a3.24 3.24 0 0 1 2.5 3.15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarMinus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M17 4.625H7a4 4 0 0 0-4 4v8.75a4 4 0 0 0 4 4h10a4 4 0 0 0 4-4v-8.75a4 4 0 0 0-4-4m-14 5h18m-4-7v4m-10-4v4"/><path stroke-miterlimit="10" d="M9.5 14.989h5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarMinusFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.75 3.935v-1.31a.75.75 0 1 0-1.5 0v1.25h-8.5v-1.25a.75.75 0 1 0-1.5 0v1.31a4.76 4.76 0 0 0-4 4.69v8.75A4.75 4.75 0 0 0 7 22.125h10a4.75 4.75 0 0 0 4.75-4.75v-8.75a4.76 4.76 0 0 0-4-4.69m-3.25 11.78h-5a.75.75 0 1 1 0-1.5h5a.75.75 0 1 1 0 1.5m5.75-6.87H3.75v-.25a3.24 3.24 0 0 1 2.5-3.15v1.15a.75.75 0 1 0 1.5 0v-1.25h8.5v1.25a.75.75 0 1 0 1.5 0v-1.15a3.24 3.24 0 0 1 2.5 3.15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarPlus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M17 4.625H7a4 4 0 0 0-4 4v8.75a4 4 0 0 0 4 4h10a4 4 0 0 0 4-4v-8.75a4 4 0 0 0-4-4m-14 5h18m-4-7v4m-10-4v4"/><path stroke-miterlimit="10" d="M11.995 12.483v5M9.5 14.989h5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarPlusFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.75 3.935v-1.31a.75.75 0 1 0-1.5 0v1.25h-8.5v-1.25a.75.75 0 1 0-1.5 0v1.31a4.76 4.76 0 0 0-4 4.69v8.75A4.75 4.75 0 0 0 7 22.125h10a4.75 4.75 0 0 0 4.75-4.75v-8.75a4.76 4.76 0 0 0-4-4.69m-3.25 11.78h-1.76v1.74a.75.75 0 1 1-1.5 0v-1.74H9.5a.75.75 0 1 1 0-1.5h1.74v-1.76a.75.75 0 1 1 1.5 0v1.76h1.76a.75.75 0 1 1 0 1.5m5.75-6.87H3.75v-.25a3.24 3.24 0 0 1 2.5-3.15v1.15a.75.75 0 1 0 1.5 0v-1.25h8.5v1.25a.75.75 0 1 0 1.5 0v-1.15a3.24 3.24 0 0 1 2.5 3.15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarQuestionMark(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M17 4.625H7a4 4 0 0 0-4 4v8.75a4 4 0 0 0 4 4h10a4 4 0 0 0 4-4v-8.75a4 4 0 0 0-4-4m-14 5h18m-4-7v4m-10-4v4"/><path stroke-miterlimit="10" d="M10.496 12.932a1.599 1.599 0 0 1 1.742-.906a1.55 1.55 0 0 1 1.137.81a1.345 1.345 0 0 1-.784 1.852a.993.993 0 0 0-.64.897v.37"/><path stroke-linejoin="round" d="M11.922 18h.004"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarQuestionMarkFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.75 3.935v-1.31a.75.75 0 1 0-1.5 0v1.25h-8.5v-1.25a.75.75 0 0 0-1.5 0v1.31a4.76 4.76 0 0 0-4 4.69v8.75A4.75 4.75 0 0 0 7 22.125h10a4.75 4.75 0 0 0 4.75-4.75v-8.75a4.76 4.76 0 0 0-4-4.69m-5.82 14.78a.75.75 0 1 1 0-1.5a.75.75 0 0 1 0 1.5m2.16-4.52a2.39 2.39 0 0 1-.5.72a2.35 2.35 0 0 1-.76.45a.17.17 0 0 0-.08.07a.24.24 0 0 0-.05.14v.34a.75.75 0 1 1-1.5 0v-.37c.005-.36.12-.709.33-1a1.68 1.68 0 0 1 .8-.61a.8.8 0 0 0 .24-.14c.052-.063.1-.13.14-.2a.55.55 0 0 0 0-.25a.642.642 0 0 0-.05-.24a.77.77 0 0 0-.23-.25a.72.72 0 0 0-.35-.14a.88.88 0 0 0-.56.07a.87.87 0 0 0-.38.4a.752.752 0 1 1-1.37-.62a2.4 2.4 0 0 1 1-1.11a2.47 2.47 0 0 1 2.48.17c.292.213.531.49.7.81a2 2 0 0 1 .21.88a2.09 2.09 0 0 1-.07.92zm6.16-5.36H3.75v-.25a3.24 3.24 0 0 1 2.5-3.15v1.15a.75.75 0 1 0 1.5 0v-1.25h8.5v1.25a.75.75 0 1 0 1.5 0v-1.15a3.24 3.24 0 0 1 2.5 3.15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarThree(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17 4.625H7a4 4 0 0 0-4 4v8.75a4 4 0 0 0 4 4h10a4 4 0 0 0 4-4v-8.75a4 4 0 0 0-4-4m0-2v4m-10-4v4m-3.88 9.737h17.76"/><path d="M12 10.515a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3m0 3a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarThreeFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.75 3.935v-1.31a.75.75 0 1 0-1.5 0v1.25h-8.5v-1.25a.75.75 0 1 0-1.5 0v1.31a4.76 4.76 0 0 0-4 4.69v8.75A4.75 4.75 0 0 0 7 22.125h10a4.75 4.75 0 0 0 4.75-4.75v-8.75a4.76 4.76 0 0 0-4-4.69m-8 5.08a2.25 2.25 0 0 1 4.5 0a2.23 2.23 0 0 1-.58 1.5c.373.41.58.945.58 1.5a2.25 2.25 0 0 1-4.5 0c0-.555.207-1.09.58-1.5a2.23 2.23 0 0 1-.58-1.5m10.5 8.36a3.25 3.25 0 0 1-3.25 3.25H7a3.25 3.25 0 0 1-3.25-3.25v-.26h16.5z"/><path fill="currentColor" d="M12 9.765a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5m0 3a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarTwo(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 4.625H7a4 4 0 0 0-4 4v8.75a4 4 0 0 0 4 4h10a4 4 0 0 0 4-4v-8.75a4 4 0 0 0-4-4m-14 6h18m-4-8v4m-10-4v4m.375 7.515h1.028m7.194 0h1.028m-5.139 0h1.028m-5.139 3.084h1.028m7.194 0h1.028m-5.139 0h1.028"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarTwoFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.75 3.935v-1.31a.75.75 0 1 0-1.5 0v1.25h-8.5v-1.25a.75.75 0 1 0-1.5 0v1.31a4.76 4.76 0 0 0-4 4.69v8.75A4.75 4.75 0 0 0 7 22.125h10a4.75 4.75 0 0 0 4.75-4.75v-8.75a4.76 4.76 0 0 0-4-4.69m-9.35 14h-1a.75.75 0 1 1 0-1.5h1a.75.75 0 1 1 0 1.5m0-3.08h-1a.75.75 0 1 1 0-1.5h1a.75.75 0 1 1 0 1.5m4.11 3.08h-1a.75.75 0 1 1 0-1.5h1a.75.75 0 1 1 0 1.5m0-3.08h-1a.75.75 0 1 1 0-1.5h1a.75.75 0 1 1 0 1.5m4.11 3.08h-1a.75.75 0 1 1 0-1.5h1a.75.75 0 1 1 0 1.5m0-3.08h-1a.75.75 0 1 1 0-1.5h1a.75.75 0 1 1 0 1.5m3.63-5H3.75v-1.25a3.24 3.24 0 0 1 2.5-3.15v1.15a.75.75 0 1 0 1.5 0v-1.25h8.5v1.25a.75.75 0 1 0 1.5 0v-1.15a3.24 3.24 0 0 1 2.5 3.15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarUpload(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M17 4.625H7a4 4 0 0 0-4 4v8.75a4 4 0 0 0 4 4h10a4 4 0 0 0 4-4v-8.75a4 4 0 0 0-4-4m-14 5h18m-4-7v4m-10-4v4"/><path stroke-miterlimit="10" d="M12 12.468v5"/><path stroke-linejoin="round" d="m14.293 14.573l-1.967-1.968a.46.46 0 0 0-.652 0l-1.967 1.968"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarUploadFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.75 3.935v-1.31a.75.75 0 1 0-1.5 0v1.25h-8.5v-1.25a.75.75 0 0 0-1.5 0v1.31a4.76 4.76 0 0 0-4 4.69v8.75A4.75 4.75 0 0 0 7 22.125h10a4.75 4.75 0 0 0 4.75-4.75v-8.75a4.76 4.76 0 0 0-4-4.69m-2.92 11.17a.731.731 0 0 1-.53.22a.74.74 0 0 1-.53-.22l-1-1v3.38a.75.75 0 1 1-1.5 0v-3.38l-1 1a.75.75 0 1 1-1.06-1.06l2-2c.116-.11.252-.198.4-.26a1.11 1.11 0 0 1 .3-.08c.106-.01.213-.01.32 0c.103.012.204.04.3.08c.148.062.284.15.4.26l2 2a.74.74 0 0 1-.1 1.06m5.43-6.23H3.76v-.25a3.24 3.24 0 0 1 2.5-3.15v1.15a.75.75 0 1 0 1.5 0v-1.25h8.5v1.25a.75.75 0 1 0 1.5 0v-1.15a3.24 3.24 0 0 1 2.5 3.15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Camera(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M5.833 19.708h12.334a3.083 3.083 0 0 0 3.083-3.083V9.431a3.083 3.083 0 0 0-3.083-3.084h-1.419c-.408 0-.8-.163-1.09-.452l-1.15-1.151a1.542 1.542 0 0 0-1.09-.452h-2.836c-.41 0-.8.163-1.09.452l-1.15 1.151c-.29.29-.682.452-1.09.452H5.833A3.083 3.083 0 0 0 2.75 9.431v7.194a3.083 3.083 0 0 0 3.083 3.083"/><path d="M12 16.625a4.111 4.111 0 1 0 0-8.222a4.111 4.111 0 0 0 0 8.222"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.88 6.715a3.88 3.88 0 0 0-2.71-1.12h-1.42a.79.79 0 0 1-.56-.23l-1.15-1.15a2.3 2.3 0 0 0-1.62-.67h-2.84a2.3 2.3 0 0 0-1.62.67l-1.15 1.15a.79.79 0 0 1-.56.23H5.83A3.82 3.82 0 0 0 2 9.435v7.19a3.82 3.82 0 0 0 3.83 3.83h12.34a3.86 3.86 0 0 0 2.71-1.12a3.82 3.82 0 0 0 1.12-2.71v-7.19a3.838 3.838 0 0 0-1.12-2.72m-8.92 9.48a3.68 3.68 0 1 1 3.68-3.68a3.69 3.69 0 0 1-3.68 3.68"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraTwo(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><rect width="18.5" height="16" x="2.75" y="4" rx="4"/><circle cx="12" cy="12.5" r="3.75"/><path stroke-linecap="round" d="M17 7.003h.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraTwoFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.25 3H6.75A4.75 4.75 0 0 0 2 7.75v8a4.75 4.75 0 0 0 4.75 4.75h10.5A4.75 4.75 0 0 0 22 15.75v-8A4.75 4.75 0 0 0 17.25 3M12 16.18a3.93 3.93 0 1 1 0-7.86a3.93 3.93 0 0 1 0 7.86m5.5-8.68H17A.75.75 0 1 1 17 6h.5a.75.75 0 1 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cancel(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 21.5a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19m6.713-2.787L5.287 5.287"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CancelFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.75A10.25 10.25 0 1 0 22.25 12A10.18 10.18 0 0 0 12 1.75m6.69 15.89L6.37 5.31a8.75 8.75 0 0 1 12.32 12.33"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretDown(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m12.777 15.628l6.094-6.769A.518.518 0 0 0 18.488 8H5.512a.519.519 0 0 0-.383.86l6.094 6.81a1.037 1.037 0 0 0 1.554-.042"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretDownFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.696 8.72a1.22 1.22 0 0 1-.3.64l-6.09 6.76a1.85 1.85 0 0 1-.58.46a1.7 1.7 0 0 1-1.42.03a1.75 1.75 0 0 1-.62-.42l-6.1-6.83a1.28 1.28 0 0 1-.31-.64a1.31 1.31 0 0 1 .56-1.26a1.36 1.36 0 0 1 .68-.21h13a1.293 1.293 0 0 1 1.15.76c.081.228.092.476.03.71"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretLeft(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8.372 12.777l6.769 6.094a.518.518 0 0 0 .859-.383V5.512a.518.518 0 0 0-.86-.383l-6.81 6.094a1.036 1.036 0 0 0 .042 1.554"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretLeftFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.75 5.484v13a1.26 1.26 0 0 1-.76 1.16a1.25 1.25 0 0 1-.51.1h-.19a1.2 1.2 0 0 1-.65-.31l-6.76-6.08a2 2 0 0 1-.46-.59a1.78 1.78 0 0 1-.17-.73c-.008-.251.04-.5.14-.73c.095-.237.242-.449.43-.62l6.82-6.11a1.26 1.26 0 0 1 .65-.3a1.23 1.23 0 0 1 1.25.56c.131.192.204.417.21.65"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretRight(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.628 11.223L8.86 5.129a.517.517 0 0 0-.86.383v12.976a.519.519 0 0 0 .86.383l6.81-6.094a1.037 1.037 0 0 0-.042-1.554"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretRightFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.75 11.989a1.82 1.82 0 0 1-.57 1.36l-6.82 6.1a1.27 1.27 0 0 1-.65.31h-.19a1.3 1.3 0 0 1-.52-.1a1.23 1.23 0 0 1-.54-.47a1.19 1.19 0 0 1-.21-.68v-13a1.2 1.2 0 0 1 .21-.69a1.23 1.23 0 0 1 1.25-.56c.24.039.464.143.65.3l6.76 6.09c.19.162.344.363.45.59c.114.234.175.49.18.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretUp(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.223 8.372L5.129 15.14a.517.517 0 0 0 .383.859h12.976a.519.519 0 0 0 .383-.86l-6.094-6.81a1.036 1.036 0 0 0-1.554.042"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretUpFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.66 16.01a1.35 1.35 0 0 1-.47.54c-.203.13-.439.2-.68.2h-13a1.29 1.29 0 0 1-.69-.2a1.28 1.28 0 0 1-.56-1.25a1.27 1.27 0 0 1 .31-.65l6.09-6.77a1.71 1.71 0 0 1 .58-.45a1.83 1.83 0 0 1 .73-.18c.253.003.503.05.74.14c.23.101.438.247.61.43l6.11 6.83c.163.182.267.408.3.65a1.24 1.24 0 0 1-.07.71"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chart(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.3" d="M6.209 12.324H4.401c-.579 0-1.048.47-1.048 1.048v6.83c0 .578.47 1.048 1.048 1.048H6.21c.58 0 1.049-.47 1.049-1.049v-6.829a1.05 1.05 0 0 0-1.049-1.049m6.694-9.573h-1.808c-.58 0-1.049.47-1.049 1.049V20.2c0 .58.47 1.049 1.05 1.049h1.807c.58 0 1.049-.47 1.049-1.049V3.8c0-.58-.47-1.049-1.05-1.049m6.696 5.176H17.79c-.58 0-1.049.47-1.049 1.05V20.2c0 .58.47 1.049 1.049 1.049h1.808a1.05 1.05 0 0 0 1.049-1.049V8.976c0-.58-.47-1.049-1.05-1.049"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartB(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 9h4v12H5a2 2 0 0 1-2-2v-8a2 2 0 0 1 2-2m6-6h2a2 2 0 0 1 2 2v16H9V5a2 2 0 0 1 2-2m4 4h4a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartBfill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.94 7.05a2.72 2.72 0 0 0-1.94-.8h-3.25V5A2.75 2.75 0 0 0 13 2.25h-2A2.73 2.73 0 0 0 8.25 5v3.25H5A2.73 2.73 0 0 0 2.25 11v8A2.75 2.75 0 0 0 5 21.75h14A2.77 2.77 0 0 0 21.75 19V9a2.738 2.738 0 0 0-.81-1.95m-6.69 13.2h-4.5V4.97A1.23 1.23 0 0 1 11 3.72h2c.329.002.644.13.88.36a1.3 1.3 0 0 1 .37.89z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartDown(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.3" d="M18 12h1.874c.6 0 1.087.487 1.087 1.087v7.076c0 .6-.487 1.087-1.087 1.087h-1.873c-.6 0-1.087-.487-1.087-1.087v-7.076c0-.6.487-1.087 1.087-1.087m-6.938-4.625h1.874c.6 0 1.086.487 1.086 1.087v11.701c0 .6-.486 1.087-1.086 1.087h-1.874c-.6 0-1.086-.487-1.086-1.087V8.462c0-.6.486-1.087 1.086-1.087M4.126 2.75h1.873c.6 0 1.087.487 1.087 1.087v16.326c0 .6-.487 1.087-1.087 1.087H4.126c-.6 0-1.087-.487-1.087-1.087V3.837c0-.6.487-1.087 1.087-1.087"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartDownB(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 21h-4v-9h4a2 2 0 0 1 2 2v5a2 2 0 0 1-2 2m-4 0H9V7h4a2 2 0 0 1 2 2zm-6 0H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartDownBfill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.94 12.05a2.723 2.723 0 0 0-1.94-.8h-3.25V9A2.75 2.75 0 0 0 13 6.25H9.75V5A2.75 2.75 0 0 0 7 2.25H5A2.73 2.73 0 0 0 2.25 5v14A2.75 2.75 0 0 0 5 21.75h14A2.77 2.77 0 0 0 21.75 19v-5a2.738 2.738 0 0 0-.81-1.95m-6.69 8.2h-4.5V7.75H13c.329.002.644.13.88.36a1.3 1.3 0 0 1 .37.89z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartDownFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.46 13.06v7.08a1.59 1.59 0 0 1-1.59 1.58h-1.86a1.59 1.59 0 0 1-1.59-1.58v-7.08a1.59 1.59 0 0 1 1.59-1.59h1.87a1.59 1.59 0 0 1 1.58 1.59m-6.94-4.59v11.7a1.59 1.59 0 0 1-1.59 1.58h-1.87a1.59 1.59 0 0 1-1.59-1.58V8.47a1.59 1.59 0 0 1 1.59-1.59h1.87a1.59 1.59 0 0 1 1.59 1.59M7.58 3.84v16.33A1.58 1.58 0 0 1 6 21.75H4.12a1.58 1.58 0 0 1-1.58-1.58V3.84a1.59 1.59 0 0 1 1.58-1.59h1.89a1.59 1.59 0 0 1 1.57 1.59"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartFifteen(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 21.5a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19"/><path d="M12 2.5V12l6.721-6.721"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartFifteenFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m19.37 4.82l-.22-.22a10.29 10.29 0 1 0 .25.25zm-6.61 5.3v-6.9a8.65 8.65 0 0 1 4.88 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartFifty(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 21.5a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19m0-19v19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartFiftyFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.75A10.25 10.25 0 1 0 22.25 12A10.26 10.26 0 0 0 12 1.75m.75 19V3.29a8.74 8.74 0 0 1 0 17.42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.755 13.38v6.83a1.54 1.54 0 0 1-1.54 1.54h-1.81a1.54 1.54 0 0 1-1.55-1.54v-6.83a1.54 1.54 0 0 1 1.55-1.55h1.81a1.538 1.538 0 0 1 1.54 1.55m6.7-9.58v16.41a1.54 1.54 0 0 1-1.55 1.54h-1.81a1.55 1.55 0 0 1-1.55-1.54V3.8a1.56 1.56 0 0 1 1.55-1.55h1.81a1.55 1.55 0 0 1 1.55 1.55m6.69 5.18v11.23a1.54 1.54 0 0 1-1.54 1.54h-1.81a1.538 1.538 0 0 1-1.55-1.54V8.98a1.55 1.55 0 0 1 1.55-1.55h1.85a1.55 1.55 0 0 1 1.5 1.55"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartTwentyFive(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 21.5a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19"/><path d="M12 2.5V12h9.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartTwentyFiveFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.75A10.25 10.25 0 1 0 22.25 12A10.26 10.26 0 0 0 12 1.75m.75 9.5v-8a8.74 8.74 0 0 1 8 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartUp(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.3" d="M6 12H4.125c-.6 0-1.087.487-1.087 1.087v7.076c0 .6.487 1.087 1.087 1.087h1.873c.6 0 1.087-.487 1.087-1.087v-7.076c0-.6-.487-1.087-1.087-1.087m6.939-4.625h-1.873c-.6 0-1.087.487-1.087 1.087v11.701c0 .6.486 1.087 1.086 1.087h1.874c.6 0 1.087-.487 1.087-1.087V8.462c0-.6-.487-1.087-1.087-1.087m6.937-4.625h-1.873c-.6 0-1.087.487-1.087 1.087v16.326c0 .6.487 1.087 1.087 1.087h1.873c.6 0 1.087-.487 1.087-1.087V3.837c0-.6-.487-1.087-1.087-1.087"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartUpB(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 12h4v9H5a2 2 0 0 1-2-2v-5a2 2 0 0 1 2-2m6-5h4v14H9V9a2 2 0 0 1 2-2m6-4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2h-4V5a2 2 0 0 1 2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartUpBfill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.94 3.05a2.72 2.72 0 0 0-1.94-.8h-2A2.73 2.73 0 0 0 14.25 5v1.25H11A2.73 2.73 0 0 0 8.25 9v2.25H5A2.732 2.732 0 0 0 2.25 14v5A2.75 2.75 0 0 0 5 21.75h14A2.77 2.77 0 0 0 21.75 19V5a2.738 2.738 0 0 0-.81-1.95m-6.69 17.2h-4.5V8.97A1.23 1.23 0 0 1 11 7.72h3.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartUpFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.58 13.06v7.08A1.58 1.58 0 0 1 6 21.72H4.12a1.58 1.58 0 0 1-1.58-1.58v-7.08a1.59 1.59 0 0 1 1.58-1.59H6a1.59 1.59 0 0 1 1.58 1.59m6.94-4.59v11.7a1.59 1.59 0 0 1-1.59 1.58h-1.87a1.59 1.59 0 0 1-1.59-1.58V8.47a1.59 1.59 0 0 1 1.59-1.59h1.87a1.59 1.59 0 0 1 1.59 1.59m6.94-4.63v16.33a1.59 1.59 0 0 1-1.59 1.58H18a1.59 1.59 0 0 1-1.59-1.58V3.84A1.59 1.59 0 0 1 18 2.25h1.87a1.59 1.59 0 0 1 1.59 1.59"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartVertical(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.3" d="M13.632 6.352V4.589c0-.565-.458-1.023-1.023-1.023H7.103c-.565 0-1.023.458-1.023 1.023v1.763c0 .565.458 1.023 1.023 1.023h5.506c.565 0 1.023-.458 1.023-1.023m7.618 6.53v-1.764c0-.564-.458-1.022-1.023-1.022H7.103c-.565 0-1.023.458-1.023 1.022v1.764c0 .564.458 1.022 1.023 1.022h13.124c.565 0 1.023-.458 1.023-1.022m-3.33 6.528v-1.762c0-.565-.458-1.023-1.023-1.023H7.114c-.565 0-1.023.458-1.023 1.023v1.763c0 .565.458 1.023 1.023 1.023h9.783c.565 0 1.023-.458 1.023-1.023M2.75 3.294v17.412"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartVerticalFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.685 3.065h-5.51c-.84 0-1.52.68-1.52 1.52v1.77c0 .84.68 1.52 1.52 1.52h5.51c.84 0 1.52-.68 1.52-1.52v-1.77c0-.84-.68-1.52-1.52-1.52m9.14 8.05v1.77a1.53 1.53 0 0 1-1.53 1.52H7.175a1.52 1.52 0 0 1-1.52-1.52v-1.77a1.52 1.52 0 0 1 1.52-1.52h13.12a1.53 1.53 0 0 1 1.53 1.52m-3.32 6.53v1.77a1.53 1.53 0 0 1-1.53 1.52h-9.78a1.52 1.52 0 0 1-1.52-1.52v-1.77a1.52 1.52 0 0 1 1.52-1.52h9.78a1.53 1.53 0 0 1 1.53 1.52m-15.68 3.71a.65.65 0 0 1-.65-.65V3.295a.65.65 0 1 1 1.3 0v17.41a.65.65 0 0 1-.65.65"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Check(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4.5 11.795l4.221 4.221a1.596 1.596 0 0 0 2.272 0L19.5 7.51"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckCircle(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m6.9 12.087l2.664 2.663a1.009 1.009 0 0 0 1.433 0l5.367-5.368"/><path d="M12 21.5a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckCircleFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.75A10.25 10.25 0 1 0 22.25 12A10.26 10.26 0 0 0 12 1.75m5.07 8.34l-5.37 5.37a1.83 1.83 0 0 1-.65.44c-.497.2-1.053.2-1.55 0a2 2 0 0 1-.65-.44L6.19 12.8a1.001 1.001 0 1 1 1.41-1.42l2.67 2.67l5.38-5.37a1 1 0 0 1 1.42 0a1 1 0 0 1 0 1.38z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m7.393 12.084l2.593 2.593a.983.983 0 0 0 1.395 0l5.227-5.226"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckSquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 2h-6.5A6.76 6.76 0 0 0 2 8.75v6.5A6.76 6.76 0 0 0 8.75 22h6.5A6.76 6.76 0 0 0 22 15.25v-6.5A6.76 6.76 0 0 0 15.25 2m2.06 8.16l-5.22 5.22a2 2 0 0 1-1.41.59a2 2 0 0 1-.76-.15a1.999 1.999 0 0 1-.64-.44l-2.59-2.59a1 1 0 0 1 1.41-1.41l2.59 2.59l5.21-5.23a1.002 1.002 0 0 1 1.41 1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Checklist(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3 6l1 1l2-2m-3 7l1 1l2-2m-3 7l1 1l2-2M9 6h12M9 12h12M9 18h12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChecklistNote(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M11.692 7.889h4.52M11.692 12h4.52m-4.52 4.111h4.52M8.066 8.506a.617.617 0 1 0 0-1.234a.617.617 0 0 0 0 1.234m0 4.111a.617.617 0 1 0 0-1.234a.617.617 0 0 0 0 1.234m0 4.111a.617.617 0 1 0 0-1.234a.617.617 0 0 0 0 1.234"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChecklistNoteFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 2h-6.5A6.76 6.76 0 0 0 2 8.75v6.5A6.75 6.75 0 0 0 8.75 22h6.5A6.75 6.75 0 0 0 22 15.25v-6.5A6.76 6.76 0 0 0 15.25 2M8.04 17.48a1.37 1.37 0 1 1 1.37-1.37a1.36 1.36 0 0 1-1.37 1.37m0-4.11A1.37 1.37 0 1 1 9.41 12a1.36 1.36 0 0 1-1.37 1.42zm0-4.11a1.37 1.37 0 1 1 1.37-1.37a1.36 1.36 0 0 1-1.37 1.37m8.15 7.6h-4.52a.75.75 0 1 1 0-1.5h4.52a.75.75 0 1 1 0 1.5m0-4.11h-4.52a.75.75 0 1 1 0-1.5h4.52a.75.75 0 1 1 0 1.5m0-4.11h-4.52a.75.75 0 0 1 0-1.5h4.52a.75.75 0 1 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDown(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4 8.417l6.587 6.587a2.013 2.013 0 0 0 2.826 0L20 8.417"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDownCircle(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m7.467 10.333l3.777 3.778a1.135 1.135 0 0 0 1.606 0l3.683-3.778"/><path d="M12 21.5a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDownCircleFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.98 1.75A10.25 10.25 0 1 0 22.25 12A10.258 10.258 0 0 0 11.98 1.75m5.27 9.31l-3.68 3.77a2.19 2.19 0 0 1-.7.48a2.3 2.3 0 0 1-.82.16a2.25 2.25 0 0 1-.81-.16a2.21 2.21 0 0 1-.7-.47l-3.8-3.78a1 1 0 0 1 0-1.41a1 1 0 0 1 1.41 0l3.87 3.82l3.77-3.81a1.01 1.01 0 1 1 1.46 1.4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDownSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m7.586 10.178l3.678 3.678a1.102 1.102 0 0 0 1.564 0l3.586-3.678"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronDownSquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 2h-6.5A6.76 6.76 0 0 0 2 8.75v6.5A6.76 6.76 0 0 0 8.75 22h6.5A6.76 6.76 0 0 0 22 15.25v-6.5A6.76 6.76 0 0 0 15.25 2m1.88 8.88l-3.59 3.67a2.071 2.071 0 0 1-1.49.63a2.051 2.051 0 0 1-.81-.16a2 2 0 0 1-.68-.46l-3.68-3.68a1 1 0 0 1 1.41-1.41l3.68 3.68h.11l3.62-3.69a1.001 1.001 0 0 1 1.43 1.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeft(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m15.583 4l-6.587 6.587a2.013 2.013 0 0 0 0 2.826L15.583 20"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeftCircle(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m13.667 7.467l-3.778 3.777a1.134 1.134 0 0 0 0 1.606l3.778 3.683"/><path d="M12 21.5a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeftCircleFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.75A10.25 10.25 0 1 0 22.25 12A10.26 10.26 0 0 0 12 1.75m2.37 14.07a1 1 0 1 1-1.4 1.43L9.2 13.57a2.08 2.08 0 0 1-.48-.7a2.3 2.3 0 0 1-.16-.82a2.25 2.25 0 0 1 .16-.81a2.1 2.1 0 0 1 .47-.7l3.77-3.78a1.004 1.004 0 0 1 1.42 1.42l-3.82 3.88z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeftSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m13.822 7.586l-3.678 3.678a1.102 1.102 0 0 0 0 1.564l3.678 3.586"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronLeftSquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 2h-6.5A6.76 6.76 0 0 0 2 8.75v6.5A6.76 6.76 0 0 0 8.75 22h6.5A6.76 6.76 0 0 0 22 15.25v-6.5A6.76 6.76 0 0 0 15.25 2m-.75 13.7a1 1 0 1 1-1.42 1.41L9.5 13.5a2 2 0 0 1-.47-.69a2.11 2.11 0 0 1 0-1.61a2 2 0 0 1 .46-.68l3.68-3.68a1 1 0 1 1 1.41 1.41l-3.68 3.68v.12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRight(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m8.417 20l6.587-6.587a2.013 2.013 0 0 0 0-2.826L8.417 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRightCircle(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m10.333 16.533l3.778-3.778a1.133 1.133 0 0 0 0-1.605l-3.778-3.683"/><path d="M12 21.5a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRightCircleFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.75A10.25 10.25 0 1 0 22.25 12A10.26 10.26 0 0 0 12 1.75m3.28 11a2.09 2.09 0 0 1-.46.69l-3.78 3.78a1 1 0 0 1-1.41-1.42l3.82-3.87l-3.81-3.77a1.004 1.004 0 1 1 1.4-1.44l3.77 3.69c.204.2.367.438.48.7c.104.261.158.54.16.82a2.11 2.11 0 0 1-.17.85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRightSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m10.178 16.414l3.678-3.678a1.102 1.102 0 0 0 0-1.564l-3.678-3.586"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronRightSquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 2h-6.5A6.76 6.76 0 0 0 2 8.75v6.5A6.76 6.76 0 0 0 8.75 22h6.5A6.76 6.76 0 0 0 22 15.25v-6.5A6.76 6.76 0 0 0 15.25 2m-.23 10.77a2.109 2.109 0 0 1-.46.67l-3.68 3.68a1 1 0 0 1-1.41 0a1 1 0 0 1 0-1.41l3.68-3.68v-.12L9.5 8.3a1 1 0 1 1 1.4-1.43l3.67 3.59a2.069 2.069 0 0 1 .63 1.49a2.07 2.07 0 0 1-.18.82"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUp(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m20 15.583l-6.587-6.587a2.013 2.013 0 0 0-2.826 0L4 15.583"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUpCircle(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m16.533 13.667l-3.777-3.778a1.135 1.135 0 0 0-1.606 0l-3.683 3.778"/><path d="M12 21.5a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUpCircleFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.75A10.25 10.25 0 1 0 22.25 12A10.26 10.26 0 0 0 12 1.75m5.24 12.63a1 1 0 0 1-1.41 0l-3.87-3.82l-3.77 3.81a1.004 1.004 0 0 1-1.44-1.4l3.69-3.77a2.08 2.08 0 0 1 .7-.48a2.19 2.19 0 0 1 1.63 0c.263.108.5.268.7.47l3.77 3.77a1.002 1.002 0 0 1 0 1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUpSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m16.414 13.822l-3.678-3.678a1.105 1.105 0 0 0-1.564 0l-3.586 3.678"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChevronUpSquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 2h-6.5A6.76 6.76 0 0 0 2 8.75v6.5A6.76 6.76 0 0 0 8.75 22h6.5A6.76 6.76 0 0 0 22 15.25v-6.5A6.76 6.76 0 0 0 15.25 2m1.87 12.5a1 1 0 0 1-1.41 0l-3.68-3.68h-.12L8.3 14.5a1 1 0 1 1-1.43-1.4l3.63-3.6a2 2 0 0 1 .69-.47a2.14 2.14 0 0 1 1.61 0a2 2 0 0 1 .68.46l3.68 3.68a1 1 0 0 1-.04 1.33"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chip(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M14.423 4.71H9.728a4.705 4.705 0 0 0-4.705 4.706v4.695a4.705 4.705 0 0 0 4.705 4.705h4.695a4.705 4.705 0 0 0 4.705-4.705V9.416a4.705 4.705 0 0 0-4.705-4.706Z"/><path d="M13.314 8.044h-2.476a2.481 2.481 0 0 0-2.482 2.481v2.476a2.482 2.482 0 0 0 2.482 2.482h2.476a2.48 2.48 0 0 0 2.48-2.482v-2.476a2.48 2.48 0 0 0-2.48-2.481ZM22 8.74h-2.922m-14.005 0H2m20 6.51h-3.013m-13.823 0H2M15.26 22v-3.254m0-13.966V2M8.74 22v-3.285m0-13.904V2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChipFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.992 14.506h-2.145c.01-.13.01-.26 0-.39V9.494h2.125a.747.747 0 0 0 .748-.75a.752.752 0 0 0-.748-.75h-2.335A5.435 5.435 0 0 0 15.976 4.2V2a.752.752 0 0 0-.749-.751a.747.747 0 0 0-.748.75V4H9.491V2a.751.751 0 0 0-.748-.751a.747.747 0 0 0-.749.75V4.25a5.426 5.426 0 0 0-3.512 3.742H2.018a.747.747 0 0 0-.748.75a.751.751 0 0 0 .748.75h2.265v4.613c-.01.13-.01.26 0 .39H1.998a.747.747 0 0 0-.748.75a.751.751 0 0 0 .748.75h2.674a5.434 5.434 0 0 0 3.362 3.282V22a.751.751 0 0 0 .748.75a.747.747 0 0 0 .749-.75v-2.442h4.988v2.432a.751.751 0 0 0 .748.75a.747.747 0 0 0 .749-.75v-2.672a5.374 5.374 0 0 0 3.512-3.331h2.474a.747.747 0 0 0 .748-.75a.751.751 0 0 0-.748-.75zm-5.448-1.5c0 .857-.339 1.68-.943 2.288a3.228 3.228 0 0 1-2.28.953h-2.473a3.229 3.229 0 0 1-2.254-.967a3.246 3.246 0 0 1-.929-2.274v-2.472c0-.857.34-1.679.944-2.285a3.218 3.218 0 0 1 2.279-.946h2.474c.854 0 1.674.34 2.278.946c.605.606.944 1.428.944 2.285z"/><path fill="currentColor" d="M15.048 10.534v2.471a1.734 1.734 0 0 1-1.726 1.741h-2.474a1.734 1.734 0 0 1-1.221-.513a1.743 1.743 0 0 1-.505-1.228v-2.47a1.724 1.724 0 0 1 1.726-1.73h2.474c.458-.001.897.18 1.22.505c.324.325.506.765.506 1.224"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clipboard(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16.94 4.697H17c.796 0 1.559.308 2.121.856A2.88 2.88 0 0 1 20 7.618v9.737a3.844 3.844 0 0 1-1.172 2.754A4.055 4.055 0 0 1 16 21.25H8c-1.06 0-2.078-.41-2.828-1.14A3.844 3.844 0 0 1 4 17.355V7.618c0-.764.308-1.499.857-2.045a3.04 3.04 0 0 1 2.083-.876"/><path d="M15.94 2.75h-8c-.552 0-1 .436-1 .974V5.67c0 .538.448.974 1 .974h8c.552 0 1-.436 1-.974V3.724a.987.987 0 0 0-1-.974"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.69 5.015a3.68 3.68 0 0 0-2-1v-.28a1.73 1.73 0 0 0-1.75-1.72h-8a1.74 1.74 0 0 0-1.75 1.72v.3a3.8 3.8 0 0 0-1.86 1a3.58 3.58 0 0 0-1.08 2.58v9.72a4.6 4.6 0 0 0 1.4 3.29A4.73 4.73 0 0 0 8 21.985h8a4.73 4.73 0 0 0 3.35-1.36a4.6 4.6 0 0 0 1.4-3.29v-9.73a3.66 3.66 0 0 0-1.06-2.59m-3.45.66a.24.24 0 0 1-.25.22h-8a.24.24 0 0 1-.25-.22v-2a.24.24 0 0 1 .25-.22h8a.24.24 0 0 1 .25.22z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardThree(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16.94 4.697H17c.796 0 1.559.308 2.121.856A2.88 2.88 0 0 1 20 7.618v9.737a3.844 3.844 0 0 1-1.172 2.754A4.055 4.055 0 0 1 16 21.25H8c-1.06 0-2.078-.41-2.828-1.14A3.844 3.844 0 0 1 4 17.355V7.618c0-.764.308-1.499.857-2.045a3.04 3.04 0 0 1 2.083-.876"/><path d="M15.94 2.75h-8c-.552 0-1 .436-1 .974V5.67c0 .538.448.974 1 .974h8c.552 0 1-.436 1-.974V3.724a.987.987 0 0 0-1-.974m-7.787 8.71h7.694m-7.694 4.398h7.694"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardTwo(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16.94 4.697H17c.796 0 1.559.308 2.121.856A2.88 2.88 0 0 1 20 7.618v9.737a3.843 3.843 0 0 1-1.172 2.754A4.055 4.055 0 0 1 16 21.25H8c-1.06 0-2.078-.41-2.828-1.14A3.843 3.843 0 0 1 4 17.354V7.618c0-.764.308-1.499.857-2.045a3.04 3.04 0 0 1 2.083-.876"/><path d="M15.94 2.75h-8c-.552 0-1 .436-1 .974V5.67c0 .538.448.974 1 .974h8c.552 0 1-.436 1-.974V3.724a.987.987 0 0 0-1-.974m-7.787 8.71h7.694m-7.694 4.398h7.694"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardTwoFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.69 5.015a3.68 3.68 0 0 0-2-1v-.28a1.73 1.73 0 0 0-1.75-1.72h-8a1.74 1.74 0 0 0-1.75 1.72v.3a3.8 3.8 0 0 0-1.86 1a3.58 3.58 0 0 0-1.08 2.58v9.72a4.6 4.6 0 0 0 1.4 3.29A4.73 4.73 0 0 0 8 21.985h8a4.73 4.73 0 0 0 3.35-1.36a4.598 4.598 0 0 0 1.4-3.29v-9.73a3.66 3.66 0 0 0-1.06-2.59m-12-1.29a.24.24 0 0 1 .25-.22h8a.24.24 0 0 1 .25.22v2a.24.24 0 0 1-.25.22h-8a.24.24 0 0 1-.25-.22zm8.16 13.13h-7.7a1 1 0 0 1 0-2h7.7a1 1 0 1 1 0 2m0-4.39h-7.7a1 1 0 0 1 0-2h7.7a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clock(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="m15.172 15.172l-3.167-3.167V5.672"/><path stroke-linejoin="round" d="M12 21.5a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClockFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.75A10.25 10.25 0 1 0 22.25 12A10.26 10.26 0 0 0 12 1.75m3.88 14.13a1 1 0 0 1-.71.3a1 1 0 0 1-.7-.3l-3.46-3.46V5.68a1 1 0 1 1 2 0v5.92l2.87 2.87a1 1 0 0 1 0 1.41"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoinA(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.25 9.714c-.02.855-.296 1.67-.777 2.286c-1.415 2.011-4.68 3.429-8.473 3.429c-3.793 0-7.058-1.418-8.473-3.429c-.48-.616-.758-1.43-.777-2.286C2.75 6.56 6.894 4 12 4s9.25 2.56 9.25 5.714"/><path d="M21.25 9.714v4.572C21.25 17.44 17.106 20 12 20s-9.25-2.56-9.25-5.714V9.714"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoinAfill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 3.25c-5.59 0-10 2.82-10 6.46v4.58c0 3.62 4.39 6.46 10 6.46s10-2.84 10-6.46V9.72c0-3.63-4.43-6.47-10-6.47m8.5 11c0 2.69-3.89 5-8.5 5c-4.61 0-8.5-2.27-8.5-5V13.1c1.78 1.87 5 3 8.5 3s6.73-1.18 8.5-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoinB(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.25 9.714c-.02.855-.296 1.67-.777 2.286c-1.415 2.011-4.68 3.429-8.473 3.429c-3.793 0-7.058-1.418-8.473-3.429c-.48-.616-.758-1.43-.777-2.286C2.75 6.56 6.894 4 12 4s9.25 2.56 9.25 5.714"/><path d="M17.722 9.714a1.703 1.703 0 0 1-.48 1.139c-.876 1.002-2.896 1.708-5.242 1.708s-4.366-.706-5.242-1.708a1.703 1.703 0 0 1-.48-1.139c0-1.571 2.563-2.846 5.722-2.846c3.159 0 5.722 1.275 5.722 2.846"/><path d="M21.25 9.714v4.572C21.25 17.44 17.106 20 12 20s-9.25-2.56-9.25-5.714V9.714"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoinBfill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 3.25c-5.59 0-10 2.82-10 6.46v4.58c0 3.62 4.39 6.46 10 6.46s10-2.84 10-6.46V9.72c0-3.63-4.43-6.47-10-6.47m0 2.87c3.69 0 6.47 1.55 6.47 3.6a2.46 2.46 0 0 1-.69 1.64c-1 1.18-3.26 1.93-5.78 1.93c-2.52 0-4.75-.75-5.81-2a2.39 2.39 0 0 1-.66-1.61C5.49 7.67 8.27 6.12 12 6.12m8.5 8.17c0 2.69-3.89 5-8.5 5c-4.61 0-8.5-2.27-8.5-5v-1.15c1.78 1.87 5 3 8.5 3s6.73-1.18 8.5-3z"/><path fill="currentColor" d="M16.93 9.71a1 1 0 0 1-.29.65c-.64.73-2.33 1.45-4.68 1.45c-2.35 0-4-.72-4.71-1.48a.91.91 0 0 1-.26-.61c0-1 2.12-2.1 5-2.1s4.93 1.1 4.94 2.09"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ColorPicker(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m16.527 11.46l-7.994 7.964a1.87 1.87 0 0 1-2.303.262l-1.235 1.229a1.114 1.114 0 0 1-1.591 0l-.319-.347a1.116 1.116 0 0 1 0-1.595l1.217-1.22a1.88 1.88 0 0 1 .262-2.316l7.947-7.964z"/><path d="M18.792 11.797a1.412 1.412 0 0 1-1.985 0l-4.68-4.643a1.418 1.418 0 0 1 .455-2.296a1.393 1.393 0 0 1 1.53.308l.57.562l1.845-2.223a2.805 2.805 0 0 1 3.9.07a2.818 2.818 0 0 1 .068 3.907L18.23 9.237l.562.572a1.397 1.397 0 0 1 0 1.988"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ColorPickerFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m17.042 10.894l-4-4.006a.9.9 0 0 0-.2-.14l4.37 4.356a.691.691 0 0 0-.17-.21"/><path fill="currentColor" d="M20.962 3.002A3.558 3.558 0 0 0 18.502 2a3.637 3.637 0 0 0-2.55 1.002l-1.32 1.592a2.37 2.37 0 0 0-.68-.45a2.167 2.167 0 0 0-1.66 0a2.21 2.21 0 0 0-.7.47a2.185 2.185 0 0 0-.16 2.864l-7.4 7.421a2.636 2.636 0 0 0-.63 2.694l-.85.852a1.873 1.873 0 0 0-.41.61c-.09.227-.137.468-.14.712c.002.25.05.498.14.73c.095.216.227.412.39.582l.33.36c.176.176.383.318.61.42c.232.093.48.14.73.141c.247-.002.49-.05.72-.14a1.94 1.94 0 0 0 .6-.41l.87-.862c.36.121.743.16 1.12.11a2.629 2.629 0 0 0 1.55-.741l7.45-7.441c.371.28.825.432 1.29.43c.568 0 1.114-.222 1.52-.62a2.135 2.135 0 0 0 0-3.005l1.6-1.242l.09-.08a3.568 3.568 0 0 0-.08-5.007zm-1 3.886l-2.21 1.722a.772.772 0 0 0-.29.54a.792.792 0 0 0 .21.582l.57.57a.74.74 0 0 1 .14.22a.53.53 0 0 1 .05.251a.56.56 0 0 1-.05.25a.603.603 0 0 1-.14.211a.69.69 0 0 1-.93 0l-.11-.1l-4.37-4.357l-.18-.17a.662.662 0 0 1 0-.931a.8.8 0 0 1 .21-.15a.7.7 0 0 1 .25 0a.739.739 0 0 1 .25 0c.076.04.147.09.21.15l.58.57a.73.73 0 0 0 .56.21a.77.77 0 0 0 .54-.27l1.78-2.153a1.94 1.94 0 0 1 1.44-.56a2.058 2.058 0 0 1 1.85 1.236a2.066 2.066 0 0 1-.35 2.199z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ColorSwatch(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8.88 21.25h9.87a2.5 2.5 0 0 0 2.5-2.5v-3.63a2.5 2.5 0 0 0-2.5-2.48h-1.27m-6.1 6.09l6.1-6.11l1.87-1.87a2.49 2.49 0 0 0 0-3.53l-2.57-2.57a2.49 2.49 0 0 0-3.53 0l-1.87 1.87"/><path d="M8.88 2.75H5.25a2.5 2.5 0 0 0-2.5 2.5v13.5a2.5 2.5 0 0 0 2.5 2.5h3.63a2.5 2.5 0 0 0 2.5-2.5V5.25a2.5 2.5 0 0 0-2.5-2.5"/><path d="M7.065 18.594a1.594 1.594 0 1 0 0-3.188a1.594 1.594 0 0 0 0 3.188"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ColorSwatchFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.893 12.84a3.23 3.23 0 0 0-1.796-.91l.639-.64c.3-.304.537-.664.698-1.06a3.207 3.207 0 0 0 0-2.48a3.16 3.16 0 0 0-.698-1.06l-2.564-2.56a2.993 2.993 0 0 0-.997-.71a3.244 3.244 0 0 0-2.484 0a3.113 3.113 0 0 0-.998.7l-.638.64a3.242 3.242 0 0 0-1.086-1.973A3.227 3.227 0 0 0 8.863 2H5.242a3.248 3.248 0 0 0-2.29.955A3.264 3.264 0 0 0 2 5.25v13.5c0 .862.342 1.689.95 2.298c.608.61 1.432.952 2.292.952h13.466a3.254 3.254 0 0 0 2.295-1A3.239 3.239 0 0 0 22 18.7v-3.58a3.246 3.246 0 0 0-1.107-2.28M6.928 19.35a2.34 2.34 0 0 1-2.166-1.45a2.356 2.356 0 0 1 .508-2.562A2.341 2.341 0 0 1 9.272 17a2.344 2.344 0 0 1-2.344 2.35m5.057-12.52l1.646-1.65c.162-.163.356-.293.569-.38c.426-.17.9-.17 1.326 0c.21.093.402.221.569.38l2.563 2.57a2 2 0 0 1 .38.57a1.788 1.788 0 0 1 0 1.34c-.09.21-.219.4-.38.56l-6.673 6.7z"/><path fill="currentColor" d="M7.795 17a.852.852 0 0 1-1.007.845a.847.847 0 0 1-.671-.665a.852.852 0 0 1 .83-1.02a.847.847 0 0 1 .848.84"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompactDisk(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M12.25 2.75a9.5 9.5 0 1 1 0 19a9.5 9.5 0 0 1 0-19Z"/><path d="M14.62 12.25a2.37 2.37 0 1 0-4.74 0a2.37 2.37 0 0 0 4.74 0Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12.25 18.93a6.691 6.691 0 0 1-5.24-2.53m7.24-10.53a6.69 6.69 0 0 1 4.4 4.48"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompactDiskFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.24 4.758a10.24 10.24 0 1 0 3 7.25a10.19 10.19 0 0 0-3-7.25M12 19.438a7.44 7.44 0 0 1-5.83-2.81a.75.75 0 0 1 1.17-.93a6 6 0 0 0 2.07 1.65a6 6 0 0 0 2.59.59a.75.75 0 1 1 0 1.5m0-4.31a3.12 3.12 0 1 1 0-6.24a3.12 3.12 0 0 1 0 6.24m6.61-4.3a.78.78 0 0 1-.22 0a.74.74 0 0 1-.71-.53a5.93 5.93 0 0 0-3.91-4a.75.75 0 0 1 .45-1.43a7.35 7.35 0 0 1 3.07 1.88a7.45 7.45 0 0 1 1.82 3.1a.74.74 0 0 1-.5 1.01z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compass(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 21.5a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19"/><path d="m7.778 16.222l1.942-5.837a1.056 1.056 0 0 1 .675-.665l5.827-1.942l-1.942 5.837a1.055 1.055 0 0 1-.665.665z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompassFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.727a10.25 10.25 0 1 0 10.25 10.25A10.26 10.26 0 0 0 12 1.727m-2.15 8.38a.4.4 0 0 1 .1-.17a.46.46 0 0 1 .18-.11l6-2l-2 6.06a.489.489 0 0 1-.1.15a.37.37 0 0 1-.17.1l-6 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ContactBook(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15.854 2.75H9.687A5.139 5.139 0 0 0 4.55 7.889v8.222c0 2.838 2.3 5.139 5.138 5.139h6.167c2.838 0 5.139-2.3 5.139-5.139V7.89c0-2.838-2.3-5.139-5.139-5.139M3.007 8.403H6.09m-3.083 7.194H6.09"/><path d="M16.67 16.745c0-1.653-1.843-2.996-3.495-2.996c-1.653 0-3.495 1.343-3.495 2.995m3.495-5.376a2.348 2.348 0 1 0 0-4.696a2.348 2.348 0 0 0 0 4.696"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ContactBookFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.833 2H9.676A5.885 5.885 0 0 0 3.8 7.66H3a.75.75 0 0 0 0 1.5h.79v5.69h-.76a.75.75 0 0 0 0 1.5h.8A5.881 5.881 0 0 0 9.706 22h6.157a5.886 5.886 0 0 0 5.887-5.89V7.89A5.902 5.902 0 0 0 15.833 2m-2.679 3.93a3.097 3.097 0 0 1 3.038 3.708a3.1 3.1 0 0 1-4.23 2.253a3.099 3.099 0 0 1-.995-5.057a3.097 3.097 0 0 1 2.187-.904m3.499 11.57a.75.75 0 0 1-.75-.75c0-1.21-1.51-2.25-2.749-2.25s-2.738 1-2.738 2.25a.75.75 0 0 1-1.5 0c0-2.15 2.24-3.75 4.238-3.75c2 0 4.248 1.6 4.248 3.75a.76.76 0 0 1-.75.75"/><path fill="currentColor" d="M14.754 9.02a1.6 1.6 0 1 1-3.2.02a1.6 1.6 0 0 1 3.2-.02"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copy(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18.327 7.286h-8.044a1.932 1.932 0 0 0-1.925 1.938v10.088c0 1.07.862 1.938 1.925 1.938h8.044a1.932 1.932 0 0 0 1.925-1.938V9.224c0-1.07-.862-1.938-1.925-1.938"/><path d="M15.642 7.286V4.688c0-.514-.203-1.007-.564-1.37a1.918 1.918 0 0 0-1.361-.568H5.673c-.51 0-1 .204-1.36.568a1.945 1.945 0 0 0-.565 1.37v10.088c0 .514.203 1.007.564 1.37c.361.364.85.568 1.361.568h2.685"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CopyFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.355 6.54h-1.94V4.69a2.69 2.69 0 0 0-1.646-2.484A2.66 2.66 0 0 0 13.745 2h-8.05a2.68 2.68 0 0 0-2.67 2.69v10.09a2.68 2.68 0 0 0 2.67 2.69h1.94v1.85a2.68 2.68 0 0 0 2.67 2.68h8a2.68 2.68 0 0 0 2.67-2.68V9.23a2.69 2.69 0 0 0-2.62-2.69M7.635 9.23v6.74h-1.94a1.18 1.18 0 0 1-1.17-1.19V4.69a1.18 1.18 0 0 1 1.17-1.19h8.05a1.18 1.18 0 0 1 1.17 1.19v1.85h-4.61a2.69 2.69 0 0 0-2.67 2.69"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCard(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17.139 3.778H6.86a4.111 4.111 0 0 0-4.111 4.11v8.223a4.11 4.11 0 0 0 4.111 4.111h10.28a4.11 4.11 0 0 0 4.111-4.11V7.888a4.11 4.11 0 0 0-4.111-4.111m4.11 5.14H2.75m3.468 7.194h5.139"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCardFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.14 3.025H6.86A4.87 4.87 0 0 0 2 7.885v8.23a4.87 4.87 0 0 0 4.86 4.86h10.28a4.87 4.87 0 0 0 4.86-4.86v-8.23a4.87 4.87 0 0 0-4.86-4.86m-5.78 14.09H6.22a1 1 0 0 1 0-2h5.14a1 1 0 1 1 0 2m9.14-8.95h-17v-.28a3.37 3.37 0 0 1 3.36-3.36h10.28a3.37 3.37 0 0 1 3.36 3.36z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CrownA(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.349 5.18L16.11 8.317l-3.196-4.382a1.027 1.027 0 0 0-1.83 0L7.89 8.318L4.65 5.18a1.028 1.028 0 0 0-1.901.504v9.802a5.14 5.14 0 0 0 5.139 5.139h8.222a5.14 5.14 0 0 0 5.139-5.14V5.684a1.028 1.028 0 0 0-1.901-.504M8 16.513h8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CrownAfill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.56 4.607a1.81 1.81 0 0 0-2-.54a1.75 1.75 0 0 0-.79.63l-2.56 2.48l-2.66-3.64a1.769 1.769 0 0 0-2.48-.65a1.76 1.76 0 0 0-.62.65l-2.66 3.64l-2.56-2.48a1.77 1.77 0 0 0-2.84-.09A1.76 1.76 0 0 0 2 5.687v9.8a5.89 5.89 0 0 0 5.89 5.89h8.22a5.91 5.91 0 0 0 5.89-5.89v-9.83a1.83 1.83 0 0 0-.44-1.05m-5.57 12.91h-8a1 1 0 0 1 0-2h8a1 1 0 1 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CrownB(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m19.349 5.255l-3.238 3.11l-3.196-4.354a1.027 1.027 0 0 0-1.83 0L7.89 8.366L4.65 5.255a1.028 1.028 0 0 0-1.901.503l1.593 11.203a4.111 4.111 0 0 0 4.111 3.587h7.195a4.111 4.111 0 0 0 4.11-3.587L21.25 5.758a1.028 1.028 0 0 0-1.901-.503M8 16.447h8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CrownBfill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.626 4.682a1.811 1.811 0 0 0-.93-.61a1.791 1.791 0 0 0-1.912.69l-2.561 2.471l-2.662-3.632a1.902 1.902 0 0 0-.62-.64a1.782 1.782 0 0 0-2.482.64L7.797 7.233L5.236 4.762a1.81 1.81 0 0 0-1.911-.69a1.791 1.791 0 0 0-1.321 1.79l1.59 11.198a4.853 4.853 0 0 0 4.814 4.242h7.274a4.854 4.854 0 0 0 4.823-4.242l1.491-11.207a.492.492 0 0 0 0-.12a1.801 1.801 0 0 0-.37-1.051M16.053 17.42H8.048a1 1 0 1 1 0-2.001h8.005a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CupHot(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3.5 9.256h12v7.996a3.997 3.997 0 0 1-4 3.998h-4a4.001 4.001 0 0 1-4-3.998z"/><path d="M15.5 10.256h3a2 2 0 0 1 2 1.999v3.998a1.998 1.998 0 0 1-2 1.999h-3.13M5.89 6.748a1.41 1.41 0 0 0 0-2a1.41 1.41 0 0 1 0-1.998m3.25 3.998a1.419 1.419 0 0 0 0-2a1.41 1.41 0 0 1 0-1.998m3.49 3.998a1.41 1.41 0 0 0 0-2a1.41 1.41 0 0 1 0-1.998"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CupHotFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.406 10.377a2.754 2.754 0 0 0-1.928-.805H16.24v-.249a.756.756 0 0 0-.745-.745H3.564a.746.746 0 0 0-.746.745v7.954A4.723 4.723 0 0 0 7.541 22h3.977a4.662 4.662 0 0 0 1.8-.368a4.593 4.593 0 0 0 2.555-2.545v-.07h2.575a2.754 2.754 0 0 0 2.734-2.734v-3.977a2.714 2.714 0 0 0-.776-1.929m-.686 5.906a1.252 1.252 0 0 1-.368.875a1.213 1.213 0 0 1-.875.368H16.23a2.104 2.104 0 0 0 0-.249v-6.214h2.238a1.253 1.253 0 0 1 1.242 1.243zM5.94 7.573a.746.746 0 0 1-.527-1.272a.597.597 0 0 0 .14-.209a.557.557 0 0 0 0-.249a.577.577 0 0 0 0-.258a.596.596 0 0 0-.14-.209a2.217 2.217 0 0 1-.636-1.521a2.237 2.237 0 0 1 .636-1.531a.746.746 0 1 1 .994 1.064a.597.597 0 0 0-.139.208a.746.746 0 0 0 0 .259a.736.736 0 0 0 0 .248a.78.78 0 0 0 .15.22a2.108 2.108 0 0 1 .626 1.52a2.068 2.068 0 0 1-.637 1.522a.726.726 0 0 1-.467.208m3.231 0a.756.756 0 0 1-.527-.208a.746.746 0 0 1 0-1.054a.656.656 0 0 0 0-.945a2.068 2.068 0 0 1-.457-.696a2.078 2.078 0 0 1 0-1.64c.11-.263.268-.503.467-.706a.746.746 0 1 1 .995 1.064a.596.596 0 0 0-.14.208a.746.746 0 0 0 0 .259a.736.736 0 0 0 0 .248a.802.802 0 0 0 .15.22a2.147 2.147 0 0 1 0 3.032a.736.736 0 0 1-.488.218m3.47 0a.746.746 0 0 1-.527-1.272a.597.597 0 0 0 .14-.209a.558.558 0 0 0 0-.249a.579.579 0 0 0 0-.258a.597.597 0 0 0-.14-.209a2.178 2.178 0 0 1-.636-1.521a2.108 2.108 0 0 1 .636-1.531a.746.746 0 0 1 1.054 1.054a.477.477 0 0 0-.149.218a.658.658 0 0 0 0 .249a.74.74 0 0 0 0 .249a.616.616 0 0 0 .15.218a2.068 2.068 0 0 1 .625 1.521c.002.283-.052.564-.159.826a2.377 2.377 0 0 1-.467.696a.765.765 0 0 1-.527.218"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DashMenu(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10" stroke-width="1.5" d="M4.5 12h15m-15 5.77h15M4.5 6.23h15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dashboard(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.557 2.75H4.682A1.932 1.932 0 0 0 2.75 4.682v3.875a1.942 1.942 0 0 0 1.932 1.942h3.875a1.942 1.942 0 0 0 1.942-1.942V4.682A1.942 1.942 0 0 0 8.557 2.75m10.761 0h-3.875a1.942 1.942 0 0 0-1.942 1.932v3.875a1.943 1.943 0 0 0 1.942 1.942h3.875a1.942 1.942 0 0 0 1.932-1.942V4.682a1.932 1.932 0 0 0-1.932-1.932m0 10.75h-3.875a1.942 1.942 0 0 0-1.942 1.933v3.875a1.942 1.942 0 0 0 1.942 1.942h3.875a1.942 1.942 0 0 0 1.932-1.942v-3.875a1.932 1.932 0 0 0-1.932-1.932M8.557 13.5H4.682a1.943 1.943 0 0 0-1.932 1.943v3.875a1.932 1.932 0 0 0 1.932 1.932h3.875a1.942 1.942 0 0 0 1.942-1.932v-3.875a1.942 1.942 0 0 0-1.942-1.942"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DashboardBar(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor"><rect width="18.5" height="18.5" x="2.75" y="2.75" stroke-width="1.5" rx="6"/><path stroke-linecap="round" stroke-width="1.6" d="M7.672 16.222v-5.099m4.451 5.099V7.778m4.205 8.444V9.82"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DashboardBarFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.21 2H8.75A6.76 6.76 0 0 0 2 8.75v6.5A6.76 6.76 0 0 0 8.75 22h6.5A6.76 6.76 0 0 0 22 15.25v-6.5A6.76 6.76 0 0 0 15.21 2M8.43 16.23a.8.8 0 1 1-1.6 0v-5.1a.8.8 0 0 1 1.6 0zm4.45 0a.8.8 0 1 1-1.6 0V7.78a.8.8 0 0 1 1.6 0zm4.21 0a.8.8 0 1 1-1.6 0V9.82a.8.8 0 0 1 1.6 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DashboardBarNotification(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor"><circle cx="19" cy="5" r="2.5" stroke-width="1.5"/><path stroke-linecap="round" stroke-width="1.5" d="M21.25 10v5.25a6 6 0 0 1-6 6h-6.5a6 6 0 0 1-6-6v-6.5a6 6 0 0 1 6-6H14"/><path stroke-linecap="round" stroke-width="1.6" d="M8.276 16.036v-4.388m3.83 4.388V8.769m3.618 7.267v-5.51"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DashboardChart(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/><path stroke-linecap="round" stroke-linejoin="round" d="m7 15l2.45-3.26a1 1 0 0 1 1.33-.25L13.17 13a1 1 0 0 0 1.37-.29L17 9"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DashboardChartArrow(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/><path stroke-linecap="round" stroke-linejoin="round" d="m7 14.844l2.322-3.09a.947.947 0 0 1 .588-.36a.948.948 0 0 1 .673.123l2.265 1.43c.21.132.46.176.702.124a.948.948 0 0 0 .597-.398l2.332-3.517"/><path stroke-linecap="round" stroke-linejoin="round" d="m13.683 9.678l2.796-.522l.521 2.797"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DashboardChartArrowFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.21 2H8.75A6.76 6.76 0 0 0 2 8.75v6.5A6.76 6.76 0 0 0 8.75 22h6.5A6.76 6.76 0 0 0 22 15.25v-6.5A6.76 6.76 0 0 0 15.21 2m1.89 10.69h-.14a.76.76 0 0 1-.74-.62l-.18-1l-1.31 2a1.71 1.71 0 0 1-2.32.5l-2.27-1.44a.18.18 0 0 0-.13 0a.2.2 0 0 0-.13.08L7.56 15.3a.77.77 0 0 1-.6.3a.74.74 0 0 1-.45-.15a.75.75 0 0 1-.15-1l2.32-3.09a1.71 1.71 0 0 1 2.25-.43l2.28 1.44a.23.23 0 0 0 .28-.06l1.34-2l-1.08.15a.753.753 0 0 1-.28-1.48l2.76-.51h.36a.12.12 0 0 1 .08 0l.15.08l.15.12c.036.043.066.09.09.14a.47.47 0 0 1 .06.15l.52 2.8a.75.75 0 0 1-.54.93"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DashboardChartFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.21 2H8.75A6.76 6.76 0 0 0 2 8.75v6.5A6.76 6.76 0 0 0 8.75 22h6.5A6.76 6.76 0 0 0 22 15.25v-6.5A6.76 6.76 0 0 0 15.21 2m2.37 7.42l-2.45 3.71a1.77 1.77 0 0 1-2.4.51l-2.39-1.51a.2.2 0 0 0-.17 0a.24.24 0 0 0-.16.09l-2.45 3.24a.75.75 0 0 1-1 .15a.75.75 0 0 1-.15-1l2.45-3.26a1.75 1.75 0 0 1 2.33-.43l2.39 1.51a.21.21 0 0 0 .19 0a.23.23 0 0 0 .15-.1l2.46-3.71a.75.75 0 0 1 1.25.83z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DashboardChartNotification(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="19" cy="5" r="2.5"/><path stroke-linecap="round" d="M21.25 10v5.25a6 6 0 0 1-6 6h-6.5a6 6 0 0 1-6-6v-6.5a6 6 0 0 1 6-6H14"/><path stroke-linecap="round" stroke-linejoin="round" d="m7.27 15.045l2.205-2.934a.9.9 0 0 1 1.197-.225l2.151 1.359a.9.9 0 0 0 1.233-.261l2.214-3.34"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DashboardChartStar(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M14 21H8.75a6 6 0 0 1-6-6V8.5a6 6 0 0 1 6-6h6.5a6 6 0 0 1 6 6v5.25"/><path stroke-linejoin="round" d="m7.5 14.7l2.205-2.934a.9.9 0 0 1 1.197-.225l2.151 1.359a.9.9 0 0 0 1.233-.261L16.5 9.3m2.433 10.38l1.083.569a.276.276 0 0 0 .4-.293l-.207-1.205a.277.277 0 0 1 .08-.245l.876-.854a.276.276 0 0 0-.152-.473l-1.21-.174a.277.277 0 0 1-.207-.152l-.553-1.105a.276.276 0 0 0-.497 0l-.553 1.105a.275.275 0 0 1-.207.152l-1.194.174a.275.275 0 0 0-.23.34a.27.27 0 0 0 .076.133l.876.854a.277.277 0 0 1 .08.245l-.207 1.205a.276.276 0 0 0 .4.293l1.083-.57a.277.277 0 0 1 .263 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DashboardCheck(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.557 2.75H4.682A1.932 1.932 0 0 0 2.75 4.682v3.875a1.942 1.942 0 0 0 1.932 1.942h3.875a1.942 1.942 0 0 0 1.942-1.942V4.682A1.942 1.942 0 0 0 8.558 2.75m10.76 0h-3.875a1.942 1.942 0 0 0-1.942 1.932v3.875a1.943 1.943 0 0 0 1.942 1.942h3.875a1.942 1.942 0 0 0 1.932-1.942V4.682a1.932 1.932 0 0 0-1.932-1.932M8.557 13.5H4.682a1.943 1.943 0 0 0-1.932 1.943v3.875a1.932 1.932 0 0 0 1.932 1.932h3.875a1.942 1.942 0 0 0 1.942-1.932v-3.875a1.943 1.943 0 0 0-1.942-1.942m5.892 3.794l1.647 1.647a.622.622 0 0 0 .886 0l3.32-3.319"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DashboardCheckFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 4.68v3.88A2.44 2.44 0 0 1 8.56 11H4.68a2.379 2.379 0 0 1-1.72-.72a2.41 2.41 0 0 1-.71-1.72V4.69a2.44 2.44 0 0 1 2.43-2.44h3.88A2.46 2.46 0 0 1 11 4.68m10.75.01v3.87a2.45 2.45 0 0 1-1.498 2.252a2.42 2.42 0 0 1-.932.188h-3.88A2.419 2.419 0 0 1 13 8.56V4.69a2.393 2.393 0 0 1 .72-1.72a2.42 2.42 0 0 1 1.72-.72h3.88a2.44 2.44 0 0 1 2.43 2.44M11 15.45v3.87a2.44 2.44 0 0 1-2.44 2.43H4.68a2.42 2.42 0 0 1-2.43-2.43v-3.87A2.46 2.46 0 0 1 4.68 13h3.88A2.46 2.46 0 0 1 11 15.45m5.54 4.43a1.348 1.348 0 0 1-.53-.1a1.39 1.39 0 0 1-.45-.31l-1.64-1.64a.738.738 0 0 1 0-1.06a.75.75 0 0 1 1.06 0l1.56 1.57l3.23-3.24a.74.74 0 0 1 1.06 0a.75.75 0 0 1 0 1.06l-3.32 3.32a1.42 1.42 0 0 1-.45.3a1.3 1.3 0 0 1-.52.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DashboardCircleBar(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor"><rect width="19" height="19" x="2.5" y="2.5" stroke-width="1.5" rx="9.5"/><path stroke-linecap="round" stroke-width="1.6" d="M7.55 14.621V9.38m4.5 6.961V7.66m4.5 6.961V9.38"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DashboardCircleBarFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.75A10.25 10.25 0 1 0 22.25 12A10.26 10.26 0 0 0 12 1.75M8.35 14.63a.8.8 0 0 1-1.6 0V9.38a.8.8 0 1 1 1.6 0zm4.5 1.71a.8.8 0 1 1-1.6 0V7.66a.8.8 0 0 1 1.6 0zm4.5-1.71a.8.8 0 0 1-1.6 0V9.38a.8.8 0 0 1 1.6 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DashboardCircleChart(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><rect width="19" height="19" x="2.5" y="2.5" rx="9.5"/><path stroke-linecap="round" stroke-linejoin="round" d="m7.465 14.72l2.222-2.956a.907.907 0 0 1 1.207-.226l2.167 1.369a.907.907 0 0 0 1.243-.263l2.23-3.365"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DashboardCircleChartFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.75A10.25 10.25 0 1 0 22.25 12A10.26 10.26 0 0 0 12 1.75m5.16 8l-2.23 3.36a1.65 1.65 0 0 1-1 .71a1.622 1.622 0 0 1-1.23-.23l-2.17-1.37a.202.202 0 0 0-.11 0a.11.11 0 0 0-.09.06l-2.23 3a.73.73 0 0 1-.6.3a.78.78 0 0 1-.45-.15a.75.75 0 0 1-.14-1.05l2.22-2.95c.243-.33.6-.557 1-.64a1.69 1.69 0 0 1 1.18.22l2.17 1.38a.17.17 0 0 0 .12 0a.14.14 0 0 0 .1-.07l2.23-3.36a.75.75 0 0 1 1.25.83z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DashboardCross(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M8.557 2.75H4.682A1.932 1.932 0 0 0 2.75 4.682v3.875a1.942 1.942 0 0 0 1.932 1.942h3.875a1.942 1.942 0 0 0 1.942-1.942V4.682A1.942 1.942 0 0 0 8.558 2.75m10.76 0h-3.875a1.942 1.942 0 0 0-1.942 1.932v3.875a1.943 1.943 0 0 0 1.942 1.942h3.875a1.942 1.942 0 0 0 1.932-1.942V4.682a1.932 1.932 0 0 0-1.932-1.932M8.557 13.5H4.682a1.943 1.943 0 0 0-1.932 1.943v3.875a1.932 1.932 0 0 0 1.932 1.932h3.875a1.942 1.942 0 0 0 1.942-1.932v-3.875a1.943 1.943 0 0 0-1.942-1.942"/><path stroke-miterlimit="10" d="m19.891 14.854l-5.032 5.032m.001-5.021l5.031 5.032"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DashboardCrossFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 4.68v3.88a2.45 2.45 0 0 1-1.514 2.259A2.42 2.42 0 0 1 8.55 11H4.68a2.44 2.44 0 0 1-2.43-2.44V4.69a2.44 2.44 0 0 1 2.43-2.44h3.87A2.44 2.44 0 0 1 11 4.68m10.75.01v3.87a2.45 2.45 0 0 1-.71 1.72a2.378 2.378 0 0 1-1.72.72h-3.88a2.5 2.5 0 0 1-1.73-.71A2.441 2.441 0 0 1 13 8.56V4.69a2.39 2.39 0 0 1 .72-1.72a2.42 2.42 0 0 1 1.72-.72h3.87a2.46 2.46 0 0 1 2.44 2.44M11 15.45v3.87a2.44 2.44 0 0 1-2.45 2.43H4.68a2.45 2.45 0 0 1-1.72-.71a2.41 2.41 0 0 1-.71-1.72v-3.87a2.41 2.41 0 0 1 .71-1.72A2.47 2.47 0 0 1 4.68 13h3.87A2.46 2.46 0 0 1 11 15.45m9.42 3.92a.75.75 0 0 1-1.06 1.06l-2-2l-2 2a.739.739 0 0 1-.53.22a.708.708 0 0 1-.53-.22a.737.737 0 0 1 0-1.06l2-2l-2-2a.737.737 0 0 1 0-1.06a.75.75 0 0 1 1.06 0l2 2l2-2a.75.75 0 0 1 1.06 1.06l-2 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DashboardFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.995 4.68v3.88A2.44 2.44 0 0 1 8.545 11h-3.86a2.38 2.38 0 0 1-1.72-.72a2.41 2.41 0 0 1-.71-1.72V4.69a2.44 2.44 0 0 1 2.43-2.44h3.87a2.42 2.42 0 0 1 1.72.72a2.39 2.39 0 0 1 .72 1.71m10.75.01v3.87a2.46 2.46 0 0 1-2.43 2.44h-3.88a2.5 2.5 0 0 1-1.73-.71a2.44 2.44 0 0 1-.71-1.73V4.69a2.39 2.39 0 0 1 .72-1.72a2.42 2.42 0 0 1 1.72-.72h3.87a2.46 2.46 0 0 1 2.44 2.44m0 10.75v3.87a2.46 2.46 0 0 1-2.43 2.44h-3.88a2.5 2.5 0 0 1-1.75-.69a2.42 2.42 0 0 1-.71-1.73v-3.87a2.391 2.391 0 0 1 .72-1.72a2.421 2.421 0 0 1 1.72-.72h3.87a2.46 2.46 0 0 1 2.44 2.44zm-10.75.01v3.87a2.46 2.46 0 0 1-2.45 2.43h-3.86a2.42 2.42 0 0 1-2.43-2.43v-3.87A2.46 2.46 0 0 1 4.685 13h3.87a2.49 2.49 0 0 1 1.73.72a2.45 2.45 0 0 1 .71 1.73"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DashboardFour(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.625 2.75a3.875 3.875 0 1 0 0 7.75a3.875 3.875 0 0 0 0-7.75m10.75 0a3.875 3.875 0 1 0 0 7.75a3.875 3.875 0 0 0 0-7.75M6.625 13.5a3.875 3.875 0 1 0 0 7.75a3.875 3.875 0 0 0 0-7.75m10.75 0a3.875 3.875 0 1 0 0 7.75a3.875 3.875 0 0 0 0-7.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DashboardFourFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.01 6.625a4.16 4.16 0 0 1-.34 1.67a4.34 4.34 0 0 1-4 2.7a4.37 4.37 0 0 1-3.09-7.46a4.37 4.37 0 0 1 6.19 0a4.47 4.47 0 0 1 1.28 3.09zm10.75 0a4.33 4.33 0 0 1-.33 1.67a4.599 4.599 0 0 1-1 1.42a4.38 4.38 0 0 1-6.19 0a4.37 4.37 0 0 1 0-6.18a4.35 4.35 0 0 1 3.09-1.29a4.4 4.4 0 0 1 3.1 1.29a4.348 4.348 0 0 1 1.33 3.09m-10.75 10.75a4.16 4.16 0 0 1-.36 1.68a4.24 4.24 0 0 1-.94 1.42a4.35 4.35 0 0 1-3.1 1.28a4.37 4.37 0 0 1-3.09-7.46a4.371 4.371 0 0 1 6.19 0a4.47 4.47 0 0 1 1.28 3.09zm10.75 0a4.37 4.37 0 1 1-4.38-4.38a4.42 4.42 0 0 1 3.1 1.29a4.35 4.35 0 0 1 1.28 3.09"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DashboardMinus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M8.557 2.75H4.682A1.932 1.932 0 0 0 2.75 4.682v3.875a1.942 1.942 0 0 0 1.932 1.942h3.875a1.942 1.942 0 0 0 1.942-1.942V4.682A1.942 1.942 0 0 0 8.557 2.75m10.761 0h-3.875a1.942 1.942 0 0 0-1.942 1.932v3.875a1.943 1.943 0 0 0 1.942 1.942h3.875a1.942 1.942 0 0 0 1.932-1.942V4.682a1.932 1.932 0 0 0-1.932-1.932M8.557 13.5H4.682a1.943 1.943 0 0 0-1.932 1.943v3.875a1.932 1.932 0 0 0 1.932 1.932h3.875a1.942 1.942 0 0 0 1.942-1.932v-3.875a1.942 1.942 0 0 0-1.942-1.942"/><path stroke-miterlimit="10" d="M14.317 17.375h6.116"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DashboardMinusFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 4.68v3.88A2.44 2.44 0 0 1 8.55 11H4.68a2.38 2.38 0 0 1-1.72-.72a2.41 2.41 0 0 1-.71-1.72V4.69a2.44 2.44 0 0 1 2.43-2.44h3.87a2.46 2.46 0 0 1 1.73.72A2.42 2.42 0 0 1 11 4.68m10.75.01v3.87A2.44 2.44 0 0 1 19.32 11h-3.88a2.5 2.5 0 0 1-1.73-.71A2.44 2.44 0 0 1 13 8.56V4.69a2.39 2.39 0 0 1 .72-1.72a2.42 2.42 0 0 1 1.72-.72h3.83a2.46 2.46 0 0 1 2.44 2.44zM11 15.45v3.87a2.439 2.439 0 0 1-2.45 2.43H4.68a2.42 2.42 0 0 1-2.43-2.43v-3.87A2.46 2.46 0 0 1 4.68 13h3.87a2.47 2.47 0 0 1 1.73.72c.46.458.719 1.08.72 1.73m9.43 2.68h-6.16a.75.75 0 0 1 0-1.5h6.12a.75.75 0 1 1 0 1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DashboardPlus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M8.557 2.75H4.682A1.932 1.932 0 0 0 2.75 4.682v3.875a1.942 1.942 0 0 0 1.932 1.942h3.875a1.942 1.942 0 0 0 1.942-1.942V4.682A1.942 1.942 0 0 0 8.557 2.75m10.761 0h-3.875a1.942 1.942 0 0 0-1.942 1.932v3.875a1.943 1.943 0 0 0 1.942 1.942h3.875a1.942 1.942 0 0 0 1.932-1.942V4.682a1.932 1.932 0 0 0-1.932-1.932M8.557 13.5H4.682a1.943 1.943 0 0 0-1.932 1.943v3.875a1.932 1.932 0 0 0 1.932 1.932h3.875a1.942 1.942 0 0 0 1.942-1.932v-3.875a1.942 1.942 0 0 0-1.942-1.942"/><path stroke-miterlimit="10" d="M17.368 13.817v7.116m-3.551-3.55h7.116"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DashboardPlusFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.98 4.68v3.88a2.45 2.45 0 0 1-1.514 2.259A2.42 2.42 0 0 1 8.53 11H4.66a2.44 2.44 0 0 1-2.43-2.44V4.69a2.44 2.44 0 0 1 2.43-2.44h3.87a2.5 2.5 0 0 1 1.73.71a2.46 2.46 0 0 1 .72 1.72m10.75.01v3.87a2.45 2.45 0 0 1-.71 1.72a2.379 2.379 0 0 1-1.72.72h-3.84a2.5 2.5 0 0 1-1.73-.71a2.441 2.441 0 0 1-.71-1.73V4.69a2.391 2.391 0 0 1 .72-1.72a2.42 2.42 0 0 1 1.72-.72h3.87a2.46 2.46 0 0 1 2.44 2.44zM10.98 15.45v3.87a2.439 2.439 0 0 1-2.45 2.43H4.66a2.45 2.45 0 0 1-1.72-.71a2.41 2.41 0 0 1-.71-1.72v-3.87a2.41 2.41 0 0 1 .71-1.72A2.46 2.46 0 0 1 4.66 13h3.87a2.46 2.46 0 0 1 2.45 2.45m10.68 1.94a.76.76 0 0 1-.75.75h-2.82v2.8a.75.75 0 0 1-1.5 0v-2.8h-2.8a.75.75 0 1 1 0-1.5h2.8v-2.82a.75.75 0 1 1 1.5 0v2.82h2.82a.75.75 0 0 1 .75.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DashboardThree(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.318 2.75H4.682A1.932 1.932 0 0 0 2.75 4.682v3.885A1.932 1.932 0 0 0 4.682 10.5h14.636a1.932 1.932 0 0 0 1.932-1.932V4.682a1.932 1.932 0 0 0-1.932-1.932M8.567 13.5H4.682a1.932 1.932 0 0 0-1.932 1.933v3.885a1.932 1.932 0 0 0 1.932 1.932h3.885a1.932 1.932 0 0 0 1.932-1.932v-3.885A1.942 1.942 0 0 0 8.567 13.5m10.751 0h-3.885a1.942 1.942 0 0 0-1.932 1.933v3.885a1.932 1.932 0 0 0 1.932 1.932h3.885a1.933 1.933 0 0 0 1.932-1.932v-3.885a1.932 1.932 0 0 0-1.932-1.932"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DashboardThreeFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.75 4.69v3.88a2.41 2.41 0 0 1-.71 1.72a2.47 2.47 0 0 1-1.72.71H4.68a2.42 2.42 0 0 1-2.43-2.43V4.69a2.44 2.44 0 0 1 2.43-2.44h14.64a2.44 2.44 0 0 1 1.72.72a2.408 2.408 0 0 1 .71 1.72M11 15.43v3.89a2.42 2.42 0 0 1-2.43 2.43H4.68a2.42 2.42 0 0 1-2.43-2.43v-3.88A2.44 2.44 0 0 1 4.68 13h3.89a2.42 2.42 0 0 1 1.71.72a2.46 2.46 0 0 1 .72 1.71m10.75.01v3.88a2.41 2.41 0 0 1-.71 1.72a2.47 2.47 0 0 1-1.72.71h-3.89a2.45 2.45 0 0 1-1.72-.71a2.409 2.409 0 0 1-.71-1.72v-3.88A2.46 2.46 0 0 1 15.43 13h3.89a2.44 2.44 0 0 1 1.72.72a2.407 2.407 0 0 1 .71 1.72"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DashboardTwo(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.557 2.75H4.682A1.932 1.932 0 0 0 2.75 4.682v3.875a1.942 1.942 0 0 0 1.932 1.942h3.875a1.942 1.942 0 0 0 1.942-1.942V4.682A1.942 1.942 0 0 0 8.557 2.75m10.761 0h-3.875a1.942 1.942 0 0 0-1.942 1.932v3.875a1.943 1.943 0 0 0 1.942 1.942h3.875a1.942 1.942 0 0 0 1.932-1.942V4.682a1.932 1.932 0 0 0-1.932-1.932M8.557 13.5H4.682a1.943 1.943 0 0 0-1.932 1.943v3.875a1.932 1.932 0 0 0 1.932 1.932h3.875a1.942 1.942 0 0 0 1.942-1.932v-3.875a1.942 1.942 0 0 0-1.942-1.942m8.818-.001a3.875 3.875 0 1 0 0 7.75a3.875 3.875 0 0 0 0-7.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DashboardTwoFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 4.68v3.88a2.45 2.45 0 0 1-1.509 2.258A2.409 2.409 0 0 1 8.56 11H4.68a2.44 2.44 0 0 1-2.43-2.44V4.69a2.44 2.44 0 0 1 2.43-2.44h3.88A2.44 2.44 0 0 1 11 4.68m10.75.01v3.87a2.41 2.41 0 0 1-.71 1.72a2.378 2.378 0 0 1-1.72.72h-3.88a2.45 2.45 0 0 1-2.256-1.502A2.4 2.4 0 0 1 13 8.56V4.69a2.391 2.391 0 0 1 .72-1.72a2.42 2.42 0 0 1 1.72-.72h3.88a2.44 2.44 0 0 1 2.43 2.44M11 15.45v3.87a2.44 2.44 0 0 1-2.44 2.43H4.68a2.45 2.45 0 0 1-1.72-.71a2.41 2.41 0 0 1-.71-1.72v-3.87a2.41 2.41 0 0 1 .71-1.72A2.47 2.47 0 0 1 4.68 13h3.88A2.46 2.46 0 0 1 11 15.45m10.75 1.93A4.37 4.37 0 1 1 17.37 13a4.4 4.4 0 0 1 4.049 2.707c.22.53.332 1.099.331 1.673"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Database(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.922 12.005v5.281a2.641 2.641 0 0 1-1.32 2.271A12.715 12.715 0 0 1 12 21.247a12.715 12.715 0 0 1-6.601-1.703a2.64 2.64 0 0 1-1.32-2.258v-5.28A11.566 11.566 0 0 0 12 14.645a11.566 11.566 0 0 0 7.922-2.64m0-6.601A11.566 11.566 0 0 1 12 8.044a11.566 11.566 0 0 1-7.922-2.64A11.566 11.566 0 0 1 12 2.764a11.566 11.566 0 0 1 7.922 2.64m0 0v6.601m-15.844 0V5.404"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabaseFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.435 4.898a.47.47 0 0 0-.1-.1h-.06a14 14 0 0 0-16.69 0a.24.24 0 0 0-.1.1a.66.66 0 0 0-.18.47v11.89a3.42 3.42 0 0 0 .47 1.67c.295.51.72.935 1.23 1.23a13.521 13.521 0 0 0 7 1.8h.28a13.41 13.41 0 0 0 6.71-1.79c.514-.295.94-.719 1.24-1.23a3.54 3.54 0 0 0 .46-1.69V5.378a.7.7 0 0 0-.26-.48m-1.33 6.7a12.52 12.52 0 0 1-14.35 0v-4.7a12.13 12.13 0 0 0 7.14 1.92h.59a12.27 12.27 0 0 0 6.65-1.92z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabaseTwo(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18.938 14.315v3.467a2.309 2.309 0 0 1-1.157 1.987A11.14 11.14 0 0 1 12 21.248a11.141 11.141 0 0 1-5.781-1.49a2.312 2.312 0 0 1-1.157-1.976v-3.467A10.133 10.133 0 0 0 12 16.626c2.52.12 4.994-.704 6.938-2.31m0-.001V9.694A10.134 10.134 0 0 1 12 12.004c-2.52.12-4.994-.704-6.937-2.31v4.621m13.875-9.243A10.134 10.134 0 0 0 12 2.762c-2.52-.12-4.994.704-6.937 2.31"/><path d="M18.938 9.69V5.067A10.134 10.134 0 0 1 12 7.378c-2.52.12-4.994-.704-6.937-2.31v4.621"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabaseTwoFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.449 4.895v-.09l-.06-.13l-.09-.1l-.06-.06a12.31 12.31 0 0 0-14.77 0l-.09.09a.67.67 0 0 0-.06.12a.27.27 0 0 0 0 .12a.47.47 0 0 0 0 .17v12.72a3.05 3.05 0 0 0 1.53 2.62a11.79 11.79 0 0 0 6.14 1.59h.24a11.89 11.89 0 0 0 5.93-1.58a3.06 3.06 0 0 0 1.12-1.11c.271-.464.413-.993.41-1.53V5.035a.88.88 0 0 0-.24-.14m-7.66-1.42a9.37 9.37 0 0 1 5.63 1.56a11.08 11.08 0 0 1-11.34 0a9.5 9.5 0 0 1 5.71-1.56m6.18 10.42a10.85 10.85 0 0 1-12.37 0v-2.79a10.77 10.77 0 0 0 6.15 1.6a10.94 10.94 0 0 0 6.22-1.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeliveryTruck(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16 6.25v9.51c-.12.149-.217.314-.29.49H8.29a2.5 2.5 0 0 0-4.58 0H3a1 1 0 0 1-1-1v-9a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2m6 7.11v2.89h-1.71a2.49 2.49 0 0 0-4.29-.49V7.25h2.43a1 1 0 0 1 .86.49l.91 1.51l1.23 2.05a4 4 0 0 1 .57 2.06"/><path d="M8.5 17.25a2.5 2.5 0 1 1-4.79-1a2.5 2.5 0 0 1 4.79 1m12 0a2.5 2.5 0 1 1-4.79-1c.073-.176.17-.341.29-.49a2.49 2.49 0 0 1 4.29.49c.14.315.212.656.21 1m-9.5-6H6m5-3H6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeliveryTruckFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m22.031 10.875l-2.136-3.543a1.764 1.764 0 0 0-1.497-.846h-1.677v-.249a2.73 2.73 0 0 0-.804-1.935a2.748 2.748 0 0 0-1.94-.802H3.994a2.732 2.732 0 0 0-2.541 1.687a2.71 2.71 0 0 0-.204 1.05v8.958a1.739 1.739 0 0 0 1.507 1.722c-.005.09-.005.18 0 .269a3.147 3.147 0 0 0 .948 2.279A3.242 3.242 0 0 0 6 20.46a3.279 3.279 0 0 0 2.285-.956a3.26 3.26 0 0 0 .96-2.279a2.074 2.074 0 0 0 0-.248h5.509a2.065 2.065 0 0 0 0 .248a3.146 3.146 0 0 0 .948 2.28A3.241 3.241 0 0 0 18 20.5a3.278 3.278 0 0 0 2.285-.956a3.26 3.26 0 0 0 .959-2.279a2.065 2.065 0 0 0 0-.249H22a.76.76 0 0 0 .749-.746v-2.876c0-.89-.25-1.762-.719-2.519m-14.293 6.31a1.688 1.688 0 0 1-.519 1.225a1.79 1.79 0 0 1-2.466 0a1.732 1.732 0 0 1-.508-1.234a1.608 1.608 0 0 1 .14-.687c.132-.313.359-.577.648-.757a1.74 1.74 0 0 1 .998-.288c.338 0 .668.1.948.288c.287.183.513.446.65.757c.098.215.15.45.149.687zm3.244-4.976h-4.99a1 1 0 0 1-.999-.995a.994.994 0 0 1 .998-.996h4.991a1 1 0 0 1 .998.996a.994.994 0 0 1-.998.995m0-3.424h-4.99a1 1 0 0 1-.999-.995a.994.994 0 0 1 .998-.995h4.991a1 1 0 0 1 .998.995a.994.994 0 0 1-.998.995m8.734 8.4a1.687 1.687 0 0 1-.52 1.225a1.79 1.79 0 0 1-2.465 0a1.732 1.732 0 0 1-.509-1.234a1.64 1.64 0 0 1 .33-1.006c.246-.327.599-.56.998-.657h.25a.32.32 0 0 1 .139 0h.2c.303.035.592.148.838.329c.247.181.44.425.559.707c.099.215.15.45.15.686z"/><path fill="currentColor" d="M17.96 15.434a.318.318 0 0 0-.14 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DialerKeypad(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 4.926a1.088 1.088 0 1 0 0-2.176a1.088 1.088 0 0 0 0 2.176m5.441 0a1.088 1.088 0 1 0 0-2.176a1.088 1.088 0 0 0 0 2.176m-10.882 0a1.088 1.088 0 1 0 0-2.176a1.088 1.088 0 0 0 0 2.176M12 10.368a1.088 1.088 0 1 0 0-2.177a1.088 1.088 0 0 0 0 2.177m5.441 0a1.088 1.088 0 1 0 0-2.177a1.088 1.088 0 0 0 0 2.177m-10.882 0a1.088 1.088 0 1 0 0-2.177a1.088 1.088 0 0 0 0 2.177M12 15.809a1.088 1.088 0 1 0 0-2.177a1.088 1.088 0 0 0 0 2.177m0 5.441a1.088 1.088 0 1 0 0-2.176a1.088 1.088 0 0 0 0 2.176m5.441-5.441a1.088 1.088 0 1 0 0-2.177a1.088 1.088 0 0 0 0 2.177m-10.882 0a1.088 1.088 0 1 0 0-2.177a1.088 1.088 0 0 0 0 2.177"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DialerKeypadFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.84 3.84a1.838 1.838 0 0 1-3.141 1.3a1.838 1.838 0 1 1 2.602-2.601c.345.345.54.812.54 1.3m3.599 1.839a1.84 1.84 0 1 0 0-3.679a1.84 1.84 0 0 0 0 3.68M8.4 3.84a1.838 1.838 0 0 1-3.141 1.3A1.839 1.839 0 1 1 7.86 2.539c.345.345.539.812.539 1.3m5.441 5.437a1.839 1.839 0 1 1-3.677-.041a1.839 1.839 0 0 1 3.677.041m3.6 1.839a1.84 1.84 0 1 0 0-3.679a1.84 1.84 0 0 0 0 3.68M8.4 9.276a1.839 1.839 0 0 1-3.141 1.3a1.839 1.839 0 1 1 3.141-1.3m5.44 5.438a1.838 1.838 0 0 1-2.193 1.815a1.84 1.84 0 1 1 2.193-1.815m0 5.446a1.838 1.838 0 0 1-2.199 1.805a1.84 1.84 0 0 1-.663-3.333a1.84 1.84 0 0 1 2.862 1.529m5.44-5.447a1.838 1.838 0 0 1-2.193 1.815a1.84 1.84 0 0 1-1.348-2.507a1.84 1.84 0 0 1 3.541.692m-10.88 0a1.838 1.838 0 0 1-3.137 1.314A1.84 1.84 0 1 1 8.4 14.714"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DirectionRight(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.477 10.826L3.275 6.288a2.154 2.154 0 0 1 2.771-3.016l13.959 6.663a2.155 2.155 0 0 1 0 3.906l-14.16 6.865a2.154 2.154 0 0 1-2.771-3.073l3.403-4.74a2.155 2.155 0 0 0 0-2.067"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DirectionRightFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.001 11.86a2.85 2.85 0 0 1-.46 1.56a2.81 2.81 0 0 1-1.22 1.07l-14.15 6.86a2.82 2.82 0 0 1-1.27.3c-.185 0-.369-.02-.55-.06a2.86 2.86 0 0 1-1.6-.9a2.91 2.91 0 0 1-.32-3.48l3.44-4.78c.08-.19.12-.394.12-.6a1.43 1.43 0 0 0-.15-.64l-3.18-4.5a2.87 2.87 0 0 1-.39-1.82a2.89 2.89 0 0 1 2.3-2.51a3 3 0 0 1 1.79.2l14 6.67a2.93 2.93 0 0 1 1.64 2.63"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DirectionRightTwo(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.392 10.884L3.19 6.28a2.164 2.164 0 0 1 2.785-3.03L20 9.945a2.165 2.165 0 0 1 0 3.924L5.772 20.767a2.164 2.164 0 0 1-2.8-2.943l3.42-4.762a2.093 2.093 0 0 0 0-2.178m7.013 1.168H6.68"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DirectionRightTwoFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.002 11.884a2.91 2.91 0 0 1-1.69 2.64l-14.21 6.9a3.18 3.18 0 0 1-1.19.25c-.195 0-.39-.02-.58-.06a2.85 2.85 0 0 1-1.53-.85a2.88 2.88 0 0 1-.77-1.57a2.84 2.84 0 0 1 .27-1.72l.06-.11l2.23-3.1a1.66 1.66 0 0 1 1.25-.7l1.16-.08l6.4-.45a1 1 0 0 0 0-2l-6.3-.44l-1.09-.07a1.7 1.7 0 0 1-1.27-.72l-2.17-3.12a2.91 2.91 0 0 1 1.92-4.34a2.82 2.82 0 0 1 1.79.2l14 6.7a2.92 2.92 0 0 1 1.72 2.64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DirectionUp(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.832 17.493L6.32 20.678a2.142 2.142 0 0 1-3-2.756L9.947 4.039a2.142 2.142 0 0 1 3.885 0l6.827 14.083a2.142 2.142 0 0 1-3.057 2.756l-4.713-3.385a2.144 2.144 0 0 0-2.057 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DirectionUpFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.883 1.976a2.85 2.85 0 0 1 1.56.46a2.81 2.81 0 0 1 1.07 1.22l6.86 14.15c.198.394.3.829.3 1.27c0 .185-.02.37-.06.55a2.86 2.86 0 0 1-.9 1.6a2.91 2.91 0 0 1-3.48.32l-4.78-3.44c-.19-.08-.394-.12-.6-.12a1.43 1.43 0 0 0-.64.15l-4.5 3.18a2.87 2.87 0 0 1-1.82.39a2.89 2.89 0 0 1-2.51-2.3a3 3 0 0 1 .2-1.79l6.67-14a2.93 2.93 0 0 1 2.63-1.64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DirectionUpRight(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9.697 12.695l-5.271-.91a2.074 2.074 0 0 1-.167-3.941l14.045-4.968a2.075 2.075 0 0 1 2.66 2.66l-4.968 14.318a2.075 2.075 0 0 1-3.981-.205l-.91-5.546a2.075 2.075 0 0 0-1.408-1.408"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DirectionUpRightFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.334 2.46c.397.388.67.886.78 1.43a2.73 2.73 0 0 1-.11 1.62l-5.15 14.86a2.67 2.67 0 0 1-.69 1.11a2.058 2.058 0 0 1-.43.35a2.91 2.91 0 0 1-4.45-1.74l-.95-5.81a1.5 1.5 0 0 0-.34-.51a1.36 1.36 0 0 0-.56-.35l-5.43-.93a2.89 2.89 0 0 1-1.56-1a2.94 2.94 0 0 1-.63-1.68a2.9 2.9 0 0 1 .48-1.72a3 3 0 0 1 1.41-1.13l14.59-5.16a2.94 2.94 0 0 1 1.61-.1a2.88 2.88 0 0 1 1.43.76"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DirectionUpRightTwo(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9.706 12.781l-5.316-.953a2.085 2.085 0 0 1-.167-3.96l14.11-4.992a2.084 2.084 0 0 1 2.673 2.673l-4.992 14.386a2.084 2.084 0 0 1-3.91-.098l-.914-5.572a2.015 2.015 0 0 0-1.484-1.484m5.571-3.979l-4.579 4.579"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DirectionUpRightTwoFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.38 2.5a3 3 0 0 1 .78 1.44a2.84 2.84 0 0 1-.12 1.63l-5.17 14.92a3.002 3.002 0 0 1-.67 1a2.86 2.86 0 0 1-2.13.85a2.9 2.9 0 0 1-2.68-2v-.12l-.62-3.77a1.71 1.71 0 0 1 .4-1.39l.76-.87l4.21-4.84a1.005 1.005 0 0 0-.71-1.715a1.005 1.005 0 0 0-.71.294l-4.76 4.15l-.83.72a1.69 1.69 0 0 1-1.4.39l-3.74-.67a3 3 0 0 1-1.56-1a2.92 2.92 0 0 1-.15-3.42a2.83 2.83 0 0 1 1.41-1.12l14.66-5.19a2.93 2.93 0 0 1 3.06.67z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DirectionUpTwo(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.884 17.608L6.28 20.81a2.166 2.166 0 0 1-3.03-2.785L9.946 4.001a2.164 2.164 0 0 1 3.924 0l6.897 14.227a2.164 2.164 0 0 1-2.943 2.8l-4.762-3.42a2.092 2.092 0 0 0-2.178 0m1.168-7.013v6.724"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DirectionUpTwoFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.907 1.975a2.91 2.91 0 0 1 2.64 1.69l6.9 14.21c.16.377.244.78.25 1.19c0 .195-.02.39-.06.58a2.85 2.85 0 0 1-.85 1.53a2.88 2.88 0 0 1-1.57.77a2.841 2.841 0 0 1-1.72-.27l-.11-.06l-3.1-2.23a1.66 1.66 0 0 1-.7-1.25l-.08-1.16l-.45-6.4a1 1 0 0 0-2 0l-.44 6.3l-.07 1.09a1.702 1.702 0 0 1-.72 1.27l-3.12 2.17a2.91 2.91 0 0 1-4.34-1.92a2.82 2.82 0 0 1 .2-1.79l6.7-14a2.92 2.92 0 0 1 2.64-1.72"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Discord(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.888 7.335a5.134 5.134 0 0 0-2.893-2.418a9.144 9.144 0 0 0-2.275-.508a9.963 9.963 0 0 0-.508 1.038a15.039 15.039 0 0 0-4.56 0a11.372 11.372 0 0 0-.519-1.038c-.752.082-1.493.249-2.208.497a5.123 5.123 0 0 0-2.904 2.44a16.176 16.176 0 0 0-1.91 9.717a16.562 16.562 0 0 0 4.98 2.528a4.339 4.339 0 0 0 1.104-1.777c-.54-.202-1.06-.45-1.557-.74c-.089-.122.254-.32.364-.354a11.826 11.826 0 0 0 10.037 0c.1 0 .453.232.364.354c-.441.342-1.424.585-1.59.828a7.4 7.4 0 0 0 1.105 1.69a16.628 16.628 0 0 0 4.99-2.53a16.232 16.232 0 0 0-2.02-9.727M8.669 14.7a1.943 1.943 0 0 1-1.92-1.955a1.943 1.943 0 0 1 1.92-1.91a1.942 1.942 0 0 1 1.933 1.965a1.943 1.943 0 0 1-1.933 1.9m6.625 0a1.943 1.943 0 0 1-1.932-1.944a1.932 1.932 0 1 1 3.865.034a1.932 1.932 0 0 1-1.933 1.899z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Divide(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-width="1.5" d="M19.5 12.001h-15"/><path stroke-width="2.5" d="M11.994 17.314h.012m-.012-10.628h.012"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DivideCircle(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-width="1.5" d="M17.014 12H6.986"/><path stroke-width="2.5" d="M11.995 15.894h.01m-.01-7.788h.01"/><path stroke-width="1.5" d="M12 21.5a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DivideCircleFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.75A10.25 10.25 0 1 0 22.25 12A10.26 10.26 0 0 0 12 1.75m0 5.11a1.25 1.25 0 1 1 0 2.5a1.25 1.25 0 0 1 0-2.5m0 10.29a1.25 1.25 0 1 1 0-2.5a1.25 1.25 0 0 1 0 2.5M17 13H7a1 1 0 0 1 0-2h10a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DivideSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16.882 12H7.118"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M11.995 15.792h.01m-.01-7.583h.01"/><rect width="18.5" height="18.5" x="2.75" y="2.75" stroke-width="1.5" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DivideSquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 2h-6.5A6.76 6.76 0 0 0 2 8.75v6.5A6.76 6.76 0 0 0 8.75 22h6.5A6.76 6.76 0 0 0 22 15.25v-6.5A6.76 6.76 0 0 0 15.25 2M12 7.21a1 1 0 1 1 0 2a1 1 0 0 1 0-2m0 9.58a1 1 0 1 1 0-2a1 1 0 0 1 0 2M16.88 13H7.12a1 1 0 0 1 0-2h9.76a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dollar(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 21.5a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19"/><path d="M9 14.433a2.82 2.82 0 0 0 3 2.57c2.42 0 3-1.39 3-2.57s-1-2.43-3-2.43s-3-.79-3-2.4a2.75 2.75 0 0 1 3-2.6a2.89 2.89 0 0 1 3 2.6M12 18.5v-1.3m0-11.7v1.499"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DollarFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.75A10.25 10.25 0 1 0 22.25 12A10.26 10.26 0 0 0 12 1.75m.75 16v.79a.75.75 0 0 1-1.5 0v-.79a3.47 3.47 0 0 1-1.76-.81a3.51 3.51 0 0 1-1.24-2.4a.75.75 0 0 1 1.23-.638a.76.76 0 0 1 .27.508a2 2 0 0 0 .71 1.39a2.09 2.09 0 0 0 1.49.5c1.56 0 2.3-.6 2.3-1.82c0-.81-.71-1.68-2.25-1.68c-3.26 0-3.75-2-3.75-3.15a3.5 3.5 0 0 1 2.42-3.2c.19-.059.383-.102.58-.13v-.78a.75.75 0 0 1 1.5 0v.81a3.61 3.61 0 0 1 3 3.22a.75.75 0 0 1-1.49.15a2.14 2.14 0 0 0-2.22-1.92h-.18a2.5 2.5 0 0 0-.72.1a2 2 0 0 0-1.18 1.03a1.87 1.87 0 0 0-.2.77c0 .82.27 1.62 2.25 1.62c2.46 0 3.75 1.6 3.75 3.18a3.11 3.11 0 0 1-3.05 3.21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dots(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 5.92A.96.96 0 1 0 12 4a.96.96 0 0 0 0 1.92m0 7.04a.96.96 0 1 0 0-1.92a.96.96 0 0 0 0 1.92M12 20a.96.96 0 1 0 0-1.92a.96.96 0 0 0 0 1.92"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotsCircle(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 21.5a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19"/><path d="M12 7.736a.673.673 0 1 0 0-1.346a.673.673 0 0 0 0 1.346m0 4.937a.673.673 0 1 0 0-1.346a.673.673 0 0 0 0 1.346m0 4.937a.673.673 0 1 0 0-1.346a.673.673 0 0 0 0 1.346"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotsCircleFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.75A10.25 10.25 0 1 0 22.25 12A10.26 10.26 0 0 0 12 1.75m0 16.86a1.67 1.67 0 1 1 0-3.34a1.67 1.67 0 0 1 0 3.34m0-4.93a1.68 1.68 0 1 1-.02-3.36a1.68 1.68 0 0 1 .02 3.36m0-4.94a1.68 1.68 0 1 1 1.67-1.67A1.67 1.67 0 0 1 12 8.74"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotsHorizontal(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18.08 12A.96.96 0 1 0 20 12a.96.96 0 0 0-1.92 0m-7.04 0a.96.96 0 1 0 1.92 0a.96.96 0 0 0-1.92 0M4 12a.96.96 0 1 0 1.92 0A.96.96 0 0 0 4 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotsHorizontalCircle(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.5 12a9.5 9.5 0 1 0 19 0a9.5 9.5 0 0 0-19 0"/><path d="M16.264 12a.673.673 0 1 0 1.346 0a.673.673 0 0 0-1.346 0m-4.937 0a.673.673 0 1 0 1.346 0a.673.673 0 0 0-1.346 0M6.39 12a.673.673 0 1 0 1.346 0a.673.673 0 0 0-1.346 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotsHorizontalCircleFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.75A10.25 10.25 0 1 0 22.25 12A10.26 10.26 0 0 0 12 1.75M7.06 13.68a1.68 1.68 0 1 1 0-3.36a1.68 1.68 0 0 1 0 3.36m4.94 0a1.68 1.68 0 1 1-.02-3.36a1.68 1.68 0 0 1 .02 3.36m4.94 0a1.68 1.68 0 1 1-.02-3.36a1.68 1.68 0 0 1 .02 3.36"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotsHorizontalSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M16.151 12a.656.656 0 1 0 1.311 0a.656.656 0 0 0-1.31 0m-4.808 0a.656.656 0 1 0 1.311 0a.656.656 0 0 0-1.31 0m-4.807 0a.656.656 0 1 0 1.31 0a.656.656 0 0 0-1.31 0"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotsHorizontalSquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 2h-6.5A6.76 6.76 0 0 0 2 8.75v6.5A6.76 6.76 0 0 0 8.75 22h6.5A6.76 6.76 0 0 0 22 15.25v-6.5A6.76 6.76 0 0 0 15.25 2M7.19 13.66a1.66 1.66 0 1 1 0-3.319a1.66 1.66 0 0 1 0 3.319m4.81 0A1.659 1.659 0 1 1 13.66 12A1.67 1.67 0 0 1 12 13.66m4.81 0a1.66 1.66 0 1 1 .06-3.32a1.66 1.66 0 0 1-.06 3.32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotsMenu(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 6.826A.913.913 0 1 0 12 5a.913.913 0 0 0 0 1.826m6.088 0a.913.913 0 1 0 0-1.826a.913.913 0 0 0 0 1.826m-12.176 0a.913.913 0 1 0 0-1.826a.913.913 0 0 0 0 1.826M12 12.913a.913.913 0 1 0 0-1.826a.913.913 0 0 0 0 1.826m6.088 0a.913.913 0 1 0 0-1.826a.913.913 0 0 0 0 1.826m-12.176 0a.913.913 0 1 0 0-1.826a.913.913 0 0 0 0 1.826M12 19a.913.913 0 1 0 0-1.826A.913.913 0 0 0 12 19m6.088 0a.913.913 0 1 0 0-1.826a.913.913 0 0 0 0 1.826M5.912 19a.913.913 0 1 0 0-1.826a.913.913 0 0 0 0 1.826"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotsSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M12 7.849a.656.656 0 1 0 0-1.311a.656.656 0 0 0 0 1.31m0 4.808a.656.656 0 1 0 0-1.312a.656.656 0 0 0 0 1.312m0 4.806a.656.656 0 1 0 0-1.31a.656.656 0 0 0 0 1.31"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DotsSquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 2h-6.5A6.76 6.76 0 0 0 2 8.75v6.5A6.76 6.76 0 0 0 8.75 22h6.5A6.76 6.76 0 0 0 22 15.25v-6.5A6.76 6.76 0 0 0 15.25 2M12 18.5a1.66 1.66 0 1 1-.02-3.318A1.66 1.66 0 0 1 12 18.5m0-4.8a1.66 1.66 0 1 1 1.65-1.66A1.67 1.67 0 0 1 12 13.66zm0-4.81a1.66 1.66 0 1 1 1.65-1.66A1.66 1.66 0 0 1 12 8.85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleArrowCircle(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M12 2.5a9.5 9.5 0 1 1 0 19a9.5 9.5 0 0 1 0-19Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M6 12h12M9 9l-3 3l3 3m6 0l3-3l-3-3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleArrowCircleFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.25 4.754a10.25 10.25 0 1 0 3 7.25a10.192 10.192 0 0 0-3-7.25m-.33 7.62c-.027.067-.06.13-.1.19l-.12.14l-3 3a1 1 0 0 1-.7.3a1 1 0 0 1-.71-1.71l1.29-1.29H8.41l1.29 1.29a1 1 0 0 1 0 1.41a1 1 0 0 1-.7.3a1.001 1.001 0 0 1-.71-.3l-3-3a1 1 0 0 1-.25-1a.33.33 0 0 1 0-.14a1 1 0 0 1 .21-.32l3-3a1 1 0 1 1 1.41 1.41l-1.29 1.3h7.17l-1.29-1.3a1 1 0 1 1 1.41-1.41l3 3l.13.15c.034.058.064.118.09.18a1.001 1.001 0 0 1 .04.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleArrowDown(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m19 5.36l-5.763 5.763a1.738 1.738 0 0 1-2.474 0L5 5.36m14 7l-5.763 5.763a1.738 1.738 0 0 1-2.474 0L5 12.36"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleArrowLeft(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m18.64 19l-5.763-5.763a1.737 1.737 0 0 1 0-2.474L18.64 5m-7 14l-5.763-5.763a1.739 1.739 0 0 1 0-2.474L11.64 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleArrowRight(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m5.36 19l5.763-5.763a1.738 1.738 0 0 0 0-2.474L5.36 5m7 14l5.763-5.763a1.738 1.738 0 0 0 0-2.474L12.36 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleArrowUp(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m5 18.64l5.763-5.763a1.738 1.738 0 0 1 2.474 0L19 18.64m-14-7l5.763-5.763a1.739 1.739 0 0 1 2.474 0L19 11.64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleCircle(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><rect width="19" height="19" x="2.5" y="2.5" rx="9.5"/><rect width="8.216" height="8.216" x="7.892" y="7.892" rx="4.108"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleCircleFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.75a10.25 10.25 0 1 0 0 20.5a10.25 10.25 0 0 0 0-20.5m0 15.119a4.869 4.869 0 1 1 0-9.738a4.869 4.869 0 0 1 0 9.738"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleDiamond(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g stroke="currentColor" stroke-width="1.5" clip-path="url(#mageDoubleDiamond0)"><path d="M8.754 3.225L3.225 8.754a4.59 4.59 0 0 0 0 6.492l5.529 5.529a4.59 4.59 0 0 0 6.492 0l5.529-5.529a4.59 4.59 0 0 0 0-6.492l-5.529-5.529a4.59 4.59 0 0 0-6.492 0Z"/><path d="M10.576 8.15L8.15 10.577a2.014 2.014 0 0 0 0 2.848l2.425 2.425a2.014 2.014 0 0 0 2.848 0l2.425-2.425a2.014 2.014 0 0 0 0-2.848L13.424 8.15a2.014 2.014 0 0 0-2.848 0Z"/></g><defs><clipPath id="mageDoubleDiamond0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleDiamondFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.993 8.351L15.65 3.007a5.17 5.17 0 0 0-7.297 0L3.007 8.35a5.17 5.17 0 0 0 0 7.297l5.345 5.345a5.17 5.17 0 0 0 7.297 0l5.344-5.344a5.17 5.17 0 0 0 0-7.298m-5.045 5.413l-2.184 2.175a2.465 2.465 0 0 1-1.76.734a2.502 2.502 0 0 1-1.768-.734l-2.174-2.175a2.513 2.513 0 0 1 0-3.528l2.174-2.174a2.561 2.561 0 0 1 3.528 0l2.184 2.184a2.513 2.513 0 0 1 0 3.528z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/><rect width="9" height="9" x="7.5" y="7.5" rx="3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleSquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 2h-6.5A6.76 6.76 0 0 0 2 8.75v6.5A6.75 6.75 0 0 0 8.75 22h6.5A6.75 6.75 0 0 0 22 15.25v-6.5A6.76 6.76 0 0 0 15.25 2m1.54 11.37a3.42 3.42 0 0 1-3.42 3.42h-2.74a3.42 3.42 0 0 1-3.42-3.42v-2.74a3.42 3.42 0 0 1 3.42-3.42h2.74a3.42 3.42 0 0 1 3.42 3.42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Download(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="M12 15.238V3.213"/><path stroke-linejoin="round" d="m7.375 10.994l3.966 3.966a.937.937 0 0 0 1.318 0l3.966-3.966"/><path stroke-linejoin="round" d="M2.75 13.85v4.625a2.313 2.313 0 0 0 2.313 2.313h13.874a2.313 2.313 0 0 0 2.313-2.313V13.85"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dribbble(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 12a10 10 0 1 1 10 10A10.011 10.011 0 0 1 2 12m11.395 1.125a11.023 11.023 0 0 0-6.48 5.41l.113.09a8.256 8.256 0 0 0 4.106 1.598a8.379 8.379 0 0 0 3.633-.349c.135 0 .18-.101.158-.247a29.975 29.975 0 0 0-1.058-4.984c-.146-.506-.315-.99-.472-1.518M12.27 10.48a25.117 25.117 0 0 1-8.616 1.125a8.313 8.313 0 0 0 2.07 5.883a12.745 12.745 0 0 1 7.086-5.759zm-8.358-.56c2.57.153 5.146-.14 7.616-.866a20.025 20.025 0 0 0-3.206-4.387a.146.146 0 0 0-.214 0a8.211 8.211 0 0 0-3.16 2.925a8.065 8.065 0 0 0-1.013 2.362zm16.356 2.97a35.184 35.184 0 0 0-5.107 0c.259 1.012.562 2.002.787 3.003c.225 1.001.383 2.036.574 3.06a8.402 8.402 0 0 0 3.768-6.03zm-2.947-7.3a8.313 8.313 0 0 0-7.368-1.654a20.81 20.81 0 0 1 3.273 4.5a10.518 10.518 0 0 0 4.095-2.847m-3.375 4.308c.214.528.416 1.034.641 1.54c0 0 .113.068.169.068c.45 0 .9-.079 1.35-.09h4.173a8.055 8.055 0 0 0-1.946-4.78a10.27 10.27 0 0 1-4.376 3.262z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Earth(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 21.5a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19"/><path d="M5.37 5.19c4.24 1 2.25 4.72 2.19 8.5c0 4.1 3.36 1.62 3.21 3.56a12.93 12.93 0 0 0 .83 4.24m3.14-18.58c0 1.38 0 2.57-.74 2.78c-1.11.31-.28 4.23-1 6.09c-.72 1.86 2.53 2.5 4 1c1.47-1.5.39-2.78 2-2.78c.57 0 1.44.1 2.32.13"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EarthFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.95 10.105v-.15a10.27 10.27 0 0 0-10-8.19a10.14 10.14 0 0 0-7 2.78l-.21.19a.1.1 0 0 0 0 .05a10.23 10.23 0 0 0 6.86 17.45h.4a10.26 10.26 0 0 0 10.25-10.25a10.05 10.05 0 0 0-.3-1.88m-9.94 10.66a12.147 12.147 0 0 1-.61-3.44c.11-1.52-1.21-1.66-1.78-1.72c-.86-.09-1.43-.15-1.43-1.88c.029-.898.119-1.794.27-2.68c.33-2.3.7-4.86-1.72-6.11a8.72 8.72 0 0 1 5.14-1.67a8.59 8.59 0 0 1 2 .23a3.601 3.601 0 0 1-.18 1.49c-1.16.33-1.18 1.85-1.2 3.62c.043.983-.058 1.967-.3 2.92a1.9 1.9 0 0 0 .76 2.38c.545.32 1.168.482 1.8.47a3.72 3.72 0 0 0 2.67-1.05a4 4 0 0 0 1.12-2.19c.03-.108.05-.218.06-.33a.66.66 0 0 1 .29 0c.28 0 .65 0 1 .06h.62a8.721 8.721 0 0 1-8.54 9.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Edit(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19.09 14.441v4.44a2.37 2.37 0 0 1-2.369 2.369H5.12a2.37 2.37 0 0 1-2.369-2.383V7.279a2.356 2.356 0 0 1 2.37-2.37H9.56"/><path d="M6.835 15.803v-2.165c.002-.357.144-.7.395-.953l9.532-9.532a1.362 1.362 0 0 1 1.934 0l2.151 2.151a1.36 1.36 0 0 1 0 1.934l-9.532 9.532a1.361 1.361 0 0 1-.953.395H8.197a1.362 1.362 0 0 1-1.362-1.362M19.09 8.995l-4.085-4.086"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.698 21.996h-11.6a3.06 3.06 0 0 1-2.2-.92a3.09 3.09 0 0 1-.9-2.21V7.276a3 3 0 0 1 .91-2.19a3 3 0 0 1 1-.67a3.06 3.06 0 0 1 1.2-.24h4.44a.75.75 0 0 1 0 1.5h-4.44a2 2 0 0 0-.63.12a1.62 1.62 0 0 0-.99 1.5v11.59a1.62 1.62 0 0 0 .47 1.16a1.62 1.62 0 0 0 1.15.47h11.6c.213 0 .423-.04.62-.12a1.54 1.54 0 0 0 .52-.35a1.49 1.49 0 0 0 .35-.52a1.51 1.51 0 0 0 .13-.63v-4.44a.75.75 0 1 1 1.5 0v4.47a3.06 3.06 0 0 1-.92 2.2a3.16 3.16 0 0 1-1 .68c-.387.14-.798.205-1.21.19"/><path fill="currentColor" d="M21.808 5.456a1.86 1.86 0 0 0-.46-.68l-2.15-2.15a1.86 1.86 0 0 0-.68-.46a2.1 2.1 0 0 0-2.31.46l-1.71 1.71v.05l-7.74 7.73a2.11 2.11 0 0 0-.61 1.48v2.17a2.12 2.12 0 0 0 2.11 2.11h2.17a2.07 2.07 0 0 0 1.48-.62l7.74-7.74l1.72-1.72c.202-.19.36-.422.46-.68a2 2 0 0 0 0-1.63zm-1.38 1.05a.56.56 0 0 1-.14.2l-1.22 1.22l-3-3l1.23-1.23a.64.64 0 0 1 .44-.18a.59.59 0 0 1 .23.05c.076.032.145.08.2.14l2.16 2.15a.69.69 0 0 1 .13.2a.59.59 0 0 1 0 .23a.6.6 0 0 1-.03.22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditPen(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4.144 16.735l.493-3.425a.973.973 0 0 1 .293-.587l9.665-9.664a1.032 1.032 0 0 1 .973-.281a5.114 5.114 0 0 1 2.346 1.372a5.079 5.079 0 0 1 1.384 2.346a1.068 1.068 0 0 1-.282.973l-9.664 9.664a1.173 1.173 0 0 1-.598.294l-3.437.492a1.044 1.044 0 0 1-1.173-1.184m8.633-11.846l4.41 4.398M3.79 21.25h16.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EditPenFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.03 6.313a6.15 6.15 0 0 0-4.28-4.26a1.87 1.87 0 0 0-.91 0a1.89 1.89 0 0 0-.77.46l-1.82 1.83l-7.83 7.83a1.65 1.65 0 0 0-.52 1l-.49 3.42a1.74 1.74 0 0 0 .07.82c.09.262.237.502.43.7c.192.2.428.35.69.44c.188.06.383.09.58.09c.083.01.167.01.25 0l3.45-.49a2 2 0 0 0 1-.5l9.67-9.67a1.93 1.93 0 0 0 .45-.77a1.9 1.9 0 0 0 .03-.9m-1.47.5a.27.27 0 0 1-.07.13l-1.3 1.3l-3.35-3.31l1.29-1.3a.35.35 0 0 1 .13-.08a.24.24 0 0 1 .12 0a4.62 4.62 0 0 1 3.19 3.15a.26.26 0 0 1-.01.11m1.66 15.19H3.79a.75.75 0 1 1 0-1.5h16.42a.75.75 0 1 1 0 1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Electricity(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.217 2.75h7.246a.525.525 0 0 1 .483.735l-2.835 6.617h4.779a.21.21 0 0 1 .157.347L9.301 21.16a.2.2 0 0 1-.358-.168l.967-7.74H5.436a.526.526 0 0 1-.494-.725l3.78-9.452a.525.525 0 0 1 .495-.326"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ElectricityDanger(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m13.242 9.8l-3.725 3.725h4.966l-3.725 3.725"/><path d="M10.367 4.462L2.752 17.655a1.885 1.885 0 0 0 1.634 2.827h15.228a1.885 1.885 0 0 0 1.633-2.827L13.633 4.462a1.885 1.885 0 0 0-3.266 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ElectricityDangerFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.855 17.249l-7.61-13.18a2.64 2.64 0 0 0-4.57 0l-7.61 13.19a2.66 2.66 0 0 0 0 2.64a2.6 2.6 0 0 0 1 1c.402.23.857.35 1.32.35h15.22c.463 0 .918-.12 1.32-.35a2.6 2.6 0 0 0 1-1a2.6 2.6 0 0 0 0-2.64zm-6.88-3.18l-3.73 3.72a.71.71 0 0 1-.53.22a.74.74 0 0 1-.53-.22a.75.75 0 0 1 0-1.06l2.45-2.44h-3.16a.75.75 0 0 1-.69-.47a.74.74 0 0 1 .16-.81l3.73-3.73a.75.75 0 0 1 1.06 1.06l-2.45 2.45h3.16a.74.74 0 0 1 .69.46a.76.76 0 0 1-.16.82"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ElectricityFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.836 10.486a.91.91 0 0 1-.21.47l-9.75 10.71a.94.94 0 0 1-.49.33c-.083.01-.167.01-.25 0a1 1 0 0 1-.41-.09a.92.92 0 0 1-.45-.46a.91.91 0 0 1-.07-.58l.86-6.86h-3.63a1.661 1.661 0 0 1-.6-.15a1.29 1.29 0 0 1-.68-.99a1.29 1.29 0 0 1 .09-.62l3.78-9.45c.1-.239.266-.444.48-.59a1.28 1.28 0 0 1 .72-.21h7.24c.209.004.414.055.6.15c.188.105.349.253.47.43c.112.179.18.38.2.59a1.18 1.18 0 0 1-.1.61l-2.39 5.57h3.65a1 1 0 0 1 .51.16a1 1 0 0 1 .43 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Email(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><rect width="18.5" height="17" x="2.682" y="3.5" rx="4"/><path stroke-linecap="round" stroke-linejoin="round" d="m2.729 7.59l7.205 4.13a3.956 3.956 0 0 0 3.975 0l7.225-4.13"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmailFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.25 2.75H6.75A4.75 4.75 0 0 0 2 7.5v9a4.75 4.75 0 0 0 4.75 4.75h10.5A4.76 4.76 0 0 0 22 16.5v-9a4.76 4.76 0 0 0-4.75-4.75m-3.65 8.32a3.26 3.26 0 0 1-3.23 0L3.52 7.14a3.25 3.25 0 0 1 3.23-2.89h10.5a3.26 3.26 0 0 1 3.23 2.89z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmailNotification(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M21.25 10.745V16.5a4 4 0 0 1-4 4H6.75a4 4 0 0 1-4-4v-9a4 4 0 0 1 4-4h7.151"/><path stroke-linecap="round" stroke-linejoin="round" d="M2.75 7.59L10 11.72a3.94 3.94 0 0 0 4 0l2.219-1.257"/><circle cx="19" cy="5" r="2.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmailOpened(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="m4.85 6.095l5.25-2.84a4 4 0 0 1 3.8 0l5.25 2.84a4 4 0 0 1 2.1 3.51v7.62a4 4 0 0 1-4 4H6.75a4 4 0 0 1-4-4v-7.62a4 4 0 0 1 2.1-3.51Z"/><path stroke-linecap="round" stroke-linejoin="round" d="m2.91 8.495l7.09 4.1a4 4 0 0 0 4 0l7.14-4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EmailOpenedFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.84 8.561l-.09-.36a5.115 5.115 0 0 0-.47-1a4.75 4.75 0 0 0-1.82-1.74l-5.25-2.84a4.8 4.8 0 0 0-4.51 0l-5.25 2.84a4.75 4.75 0 0 0-1.82 1.74a4.33 4.33 0 0 0-.46 1v.14A4.62 4.62 0 0 0 2 9.571v7.62a4.76 4.76 0 0 0 4.75 4.75h10.5a4.75 4.75 0 0 0 4.75-4.75v-7.63a4.788 4.788 0 0 0-.16-1m-8.26 3.35a3.3 3.3 0 0 1-3.25 0L3.8 8.131c.03-.076.067-.15.11-.22a3.25 3.25 0 0 1 1.25-1.19l5.25-2.84a3.2 3.2 0 0 1 1.54-.39a3.24 3.24 0 0 1 1.55.39l5.25 2.84a3.22 3.22 0 0 1 1.4 1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExchangeA(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3.53 11.47v2.118a4.235 4.235 0 0 0 4.235 4.236H20.47M3.53 6.176h12.705a4.235 4.235 0 0 1 4.236 4.236v2.117"/><path d="m17.294 14.647l3.177 3.176L17.294 21M6.706 9.353L3.529 6.176L6.706 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExchangeB(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19.75 6.75h-12a4 4 0 0 0-4 4v2m16-1v2a4 4 0 0 1-4 4h-12"/><path d="m16.75 9.75l3-3l-3-3m-10 11l-3 3l3 3"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationCircle(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-width="1.5" d="M12 2.5A9.5 9.5 0 0 0 2.5 12a9.5 9.5 0 1 0 19 0A9.5 9.5 0 0 0 12 2.5m-.005 4.222v6.334"/><path stroke-width="2" d="M12.044 16.557h-.01"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationCircleFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.75A10.25 10.25 0 1 0 22.25 12A10.26 10.26 0 0 0 12 1.75m-1 5a1 1 0 0 1 2 0v6.33a1 1 0 0 1-2 0zm1 11.08a1.25 1.25 0 1 1 1.25-1.25a1.24 1.24 0 0 1-1.21 1.23z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationHexagon(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-width="1.5" d="M3.701 15.734V8.266a1.79 1.79 0 0 1 .89-1.542l6.52-3.734a1.766 1.766 0 0 1 1.778 0l6.473 3.734a1.79 1.79 0 0 1 .937 1.542v7.468a1.79 1.79 0 0 1-.89 1.542l-6.52 3.734a1.766 1.766 0 0 1-1.778 0l-6.473-3.735a1.79 1.79 0 0 1-.937-1.54m8.294-8.995v6.319"/><path stroke-width="2" d="M12.044 16.553h-.01"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationHexagonFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.68 6.967a2.46 2.46 0 0 0-.94-.9l-6.47-3.73a2.56 2.56 0 0 0-2.53 0l-6.53 3.74a2.49 2.49 0 0 0-.92.93a2.53 2.53 0 0 0-.34 1.26v7.48c.012.451.14.892.37 1.28c.235.372.558.682.94.9l6.48 3.73a2.43 2.43 0 0 0 1.26.34a2.41 2.41 0 0 0 1.26-.34l6.53-3.73a2.64 2.64 0 0 0 .92-.93a2.5 2.5 0 0 0 .34-1.26v-7.48a2.64 2.64 0 0 0-.37-1.29M11 6.727a1 1 0 0 1 2 0v6.32a1 1 0 1 1-2 0zm1 11.07a1.25 1.25 0 1 1 .04 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.958 7.563v6.166"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 17.139h-.009"/><rect width="18.5" height="18.5" x="2.75" y="2.75" stroke-width="1.5" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationSquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 2h-6.5A6.76 6.76 0 0 0 2 8.75v6.5A6.76 6.76 0 0 0 8.75 22h6.5A6.76 6.76 0 0 0 22 15.25v-6.5A6.76 6.76 0 0 0 15.25 2m-4.29 5.56a1 1 0 0 1 2 0v6.17a1 1 0 1 1-2 0zm1 10.58a1 1 0 1 1 .03 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationTriangle(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-width="1.5" d="M10.367 4.462L2.752 17.655a1.885 1.885 0 0 0 1.634 2.827h15.228a1.885 1.885 0 0 0 1.633-2.827L13.633 4.462a1.885 1.885 0 0 0-3.266 0m1.628 3.116v6.277"/><path stroke-width="2" d="M12.043 17.326h-.009"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExclamationTriangleFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.86 17.249l-7.61-13.19a2.65 2.65 0 0 0-4.57 0l-7.61 13.19a2.6 2.6 0 0 0 0 2.64a2.6 2.6 0 0 0 1 1a2.67 2.67 0 0 0 1.32.36h15.23a2.67 2.67 0 0 0 1.32-.35a2.6 2.6 0 0 0 1-1c.23-.402.35-.857.35-1.32c0-.463-.12-.814-.35-1.216zm-10.9-9.7a1 1 0 1 1 2 0v6.28a1 1 0 0 1-2 0zm1.05 11a1.25 1.25 0 1 1 .471-.092a1.25 1.25 0 0 1-.481.092z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ExternalLink(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18.607 11.035v7.929a2.272 2.272 0 0 1-2.3 2.286H5.05a2.272 2.272 0 0 1-2.299-2.3V7.693a2.273 2.273 0 0 1 2.3-2.3h7.928M21.25 2.75L10.679 13.321M15.964 2.75h5.286v5.286"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eye(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.899 12.735a1.872 1.872 0 0 1 0-1.47c.808-1.92 2.1-3.535 3.716-4.647C8.232 5.506 10.103 4.945 12 5.004c1.897-.059 3.768.502 5.385 1.614C19 7.73 20.293 9.345 21.1 11.265a1.872 1.872 0 0 1 0 1.47c-.808 1.92-2.1 3.535-3.716 4.647c-1.617 1.112-3.488 1.673-5.385 1.614c-1.897.059-3.768-.502-5.385-1.614C5 16.27 3.707 14.655 2.9 12.735"/><path d="M12 15.5a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeClosed(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20.645 7c-.731 2.05-1.958 3.813-3.534 5.082c-1.493 1.212-3.286 1.835-5.111 1.774c-1.825.06-3.618-.562-5.111-1.774C5.313 10.813 4.086 9.05 3.355 7M12 13.857V17m5.7-1.095l-1.919-2.947m-7.562 0L6.3 15.905m15.2-4.381L19.315 9.64m-14.63 0L2.5 11.523"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.79 10.993a11.36 11.36 0 0 0-3.992-5.002a10.53 10.53 0 0 0-11.596 0a11.29 11.29 0 0 0-3.992 5.002a2.556 2.556 0 0 0-.21 1c.004.343.072.683.2 1a11.3 11.3 0 0 0 3.992 5.003a9.574 9.574 0 0 0 5.768 1.75h.3a9.67 9.67 0 0 0 5.538-1.73a11.28 11.28 0 0 0 3.992-5.002a2.55 2.55 0 0 0 .21-1a2.705 2.705 0 0 0-.21-1.021m-9.77 4.782a3.746 3.746 0 0 1-3.474-2.315a3.77 3.77 0 0 1 .807-4.103a3.75 3.75 0 0 1 6.41 2.656a3.76 3.76 0 0 1-1.107 2.674a3.74 3.74 0 0 1-2.676 1.088z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeOff(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M5.45 16.92a10.78 10.78 0 0 1-2.55-3.71a1.85 1.85 0 0 1 0-1.46A10.59 10.59 0 0 1 6.62 7.1A9 9 0 0 1 12 5.48a8.81 8.81 0 0 1 4 .85m2.56 1.72a10.85 10.85 0 0 1 2.54 3.7a1.85 1.85 0 0 1 0 1.46a10.59 10.59 0 0 1-3.72 4.65A9 9 0 0 1 12 19.48a8.81 8.81 0 0 1-4-.85"/><path d="M8.71 13.65a3.28 3.28 0 0 1-.21-1.17a3.5 3.5 0 0 1 3.5-3.5c.4-.002.796.07 1.17.21m2.12 2.12c.14.374.212.77.21 1.17a3.5 3.5 0 0 1-3.5 3.5a3.28 3.28 0 0 1-1.17-.21M3 20L19 4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeOffFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.01 4.256a9.64 9.64 0 0 0-5.813 1.74a11.284 11.284 0 0 0-3.988 5a2.556 2.556 0 0 0-.209 1c.004.343.072.682.2 1a11.561 11.561 0 0 0 2.641 3.84l3.59-3.69a3.76 3.76 0 0 1 1.36-4.22a3.733 3.733 0 0 1 3.365-.52l3.23-3.24a9.328 9.328 0 0 0-4.376-.91m9.78 6.72a11.275 11.275 0 0 0-3.639-4.73l-3.34 3.31a3.66 3.66 0 0 1 .918 2.44c0 .996-.394 1.95-1.095 2.656a3.744 3.744 0 0 1-2.644 1.105c-.9.005-1.77-.322-2.443-.92l-3.23 3.24a9.511 9.511 0 0 0 5.643 1.67h.3a9.67 9.67 0 0 0 5.543-1.75a11.274 11.274 0 0 0 3.988-5c.136-.316.207-.656.209-1a2.707 2.707 0 0 0-.21-1.02"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Facebook(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.602 21.026v-7.274H6.818a.545.545 0 0 1-.545-.545V10.33a.545.545 0 0 1 .545-.545h2.773V7a4.547 4.547 0 0 1 4.86-4.989h2.32a.556.556 0 0 1 .557.546v2.436a.557.557 0 0 1-.557.545h-1.45c-1.566 0-1.867.742-1.867 1.833v2.413h3.723a.533.533 0 0 1 .546.603l-.337 2.888a.545.545 0 0 1-.545.476h-3.364v7.274a.962.962 0 0 1-.975.974h-1.937a.961.961 0 0 1-.963-.974"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FacebookCircle(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 12.08a10 10 0 0 1-8.91 9.893V14.8h2.35a.423.423 0 0 0 .422-.37l.254-2.202a.402.402 0 0 0-.402-.465H13.09v-1.8c0-.836.233-1.407 1.428-1.407h1.112a.423.423 0 0 0 .412-.424V6.238a.423.423 0 0 0-.423-.413H13.82a3.482 3.482 0 0 0-3.714 3.81v2.116H8.339a.413.413 0 0 0-.413.424v2.2a.413.413 0 0 0 .413.413h1.767v7.037A10 10 0 0 1 2 12.08a10 10 0 1 1 20 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FacebookSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.234 2.25H6.744a4.51 4.51 0 0 0-3.198 1.33A4.47 4.47 0 0 0 2.25 6.778v10.444a4.455 4.455 0 0 0 1.296 3.198a4.494 4.494 0 0 0 3.197 1.33h4.763v-6.966h-1.82a.428.428 0 0 1-.427-.425v-2.236a.435.435 0 0 1 .438-.436h1.809v-2.18a3.54 3.54 0 0 1 .996-2.826a3.573 3.573 0 0 1 2.811-1.065h1.854a.428.428 0 0 1 .427.436v1.89a.424.424 0 0 1-.427.424h-1.123c-1.236 0-1.472.582-1.472 1.431v1.879h2.696a.428.428 0 0 1 .427.48l-.27 2.237a.424.424 0 0 1-.427.38h-2.415v6.966h2.674a4.509 4.509 0 0 0 3.197-1.33a4.471 4.471 0 0 0 1.296-3.199V6.778a4.453 4.453 0 0 0-1.304-3.206a4.494 4.494 0 0 0-3.212-1.322"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastForward(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.75 6.506v10.988a1.43 1.43 0 0 0 2.346 1.073l6.41-5.48a1.431 1.431 0 0 0 0-2.174l-6.41-5.48A1.43 1.43 0 0 0 2.75 6.506"/><path d="M12.007 6.506v10.988a1.43 1.43 0 0 0 2.347 1.073l6.395-5.48a1.431 1.431 0 0 0 0-2.174l-6.395-5.48a1.43 1.43 0 0 0-2.347 1.073"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastForwardBack(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.25 17.494V6.506a1.43 1.43 0 0 0-2.346-1.073l-6.41 5.48a1.431 1.431 0 0 0 0 2.174l6.41 5.48a1.43 1.43 0 0 0 2.346-1.073"/><path d="M11.993 17.494V6.506a1.43 1.43 0 0 0-2.347-1.073l-6.395 5.48a1.43 1.43 0 0 0 0 2.174l6.395 5.48a1.43 1.43 0 0 0 2.347-1.073"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastForwardBackFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 6.477v11a2.2 2.2 0 0 1-.36 1.16a2.15 2.15 0 0 1-.91.78a2.08 2.08 0 0 1-.91.2a1.42 1.42 0 0 1-.29 0a2.18 2.18 0 0 1-1.11-.48l-5.68-4.86v3.22a2.28 2.28 0 0 1-.35 1.16a2.24 2.24 0 0 1-.92.78a2.08 2.08 0 0 1-.91.2c-.097.01-.194.01-.29 0a2.11 2.11 0 0 1-1.1-.48l-6.41-5.49a2.19 2.19 0 0 1-.56-2.57a2 2 0 0 1 .56-.74l6.4-5.48a2.17 2.17 0 0 1 3.58 1.63v3.23l5.68-4.86c.315-.267.7-.44 1.11-.5a2.25 2.25 0 0 1 1.2.18c.371.172.686.446.91.79c.223.335.348.727.36 1.13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FastForwardFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 12.02a2.17 2.17 0 0 1-.76 1.66l-6.4 5.48a2.14 2.14 0 0 1-1.11.49c-.096.01-.194.01-.29 0a2.08 2.08 0 0 1-.91-.2a2.24 2.24 0 0 1-.92-.78a2.28 2.28 0 0 1-.35-1.16v-3.22l-5.68 4.85a2.11 2.11 0 0 1-1.11.49c-.096.01-.194.01-.29 0a2.08 2.08 0 0 1-.91-.2a2.15 2.15 0 0 1-.91-.78A2.2 2.2 0 0 1 2 17.49v-11a2.21 2.21 0 0 1 1.27-1.95a2.18 2.18 0 0 1 2.31.31l5.68 4.86V6.5a2.179 2.179 0 0 1 2.47-2.13a2.21 2.21 0 0 1 1.1.49l6.41 5.49c.239.204.43.456.56.74c.133.293.201.61.2.93"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Female(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 14.313A5.781 5.781 0 1 0 12 2.75a5.781 5.781 0 0 0 0 11.563m0 0v6.937m-3.469-3.469h6.938"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Figma(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#mageFigma0)"><path fill="currentColor" d="M17.393 8.937A3.87 3.87 0 0 0 15.04 2H8.948a3.88 3.88 0 0 0-2.373 6.937a3.88 3.88 0 0 0 0 6.136a3.87 3.87 0 1 0 6.221 3.068v-2.929a3.806 3.806 0 0 0 2.224.716a3.87 3.87 0 0 0 2.351-6.937zm-6.179 9.204a2.266 2.266 0 1 1-2.266-2.266h2.266zm0-3.87H8.948a2.266 2.266 0 1 1 0-4.532h2.266zm0-6.135H8.948a2.266 2.266 0 0 1 0-4.533h2.266zm3.827 6.136a2.245 2.245 0 0 1-2.223-1.882a1.754 1.754 0 0 1 0-.77a2.267 2.267 0 1 1 2.223 2.652m0-6.136h-2.223V3.603h2.223a2.266 2.266 0 1 1 0 4.533"/></g><defs><clipPath id="mageFigma0"><path fill="#fff" d="M5.079 2h13.843v20H5.079z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func File(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.186 2.753v3.596c0 .487.194.955.54 1.3a1.85 1.85 0 0 0 1.306.539h4.125"/><path d="M20.25 8.568v8.568a4.251 4.251 0 0 1-1.362 2.97a4.283 4.283 0 0 1-3.072 1.14h-7.59a4.294 4.294 0 0 1-3.1-1.124a4.265 4.265 0 0 1-1.376-2.986V6.862a4.25 4.25 0 0 1 1.362-2.97a4.283 4.283 0 0 1 3.072-1.14h5.714a3.503 3.503 0 0 1 2.361.905l2.96 2.722a2.971 2.971 0 0 1 1.031 2.189"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileCheck(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.186 2.753v3.596c0 .487.194.955.54 1.3a1.85 1.85 0 0 0 1.306.539h4.125"/><path d="M20.25 8.568v8.568a4.251 4.251 0 0 1-1.362 2.97a4.283 4.283 0 0 1-3.072 1.14h-7.59a4.294 4.294 0 0 1-3.1-1.124a4.265 4.265 0 0 1-1.376-2.986V6.862a4.25 4.25 0 0 1 1.362-2.97a4.283 4.283 0 0 1 3.072-1.14h5.714a3.503 3.503 0 0 1 2.361.905l2.96 2.722a2.971 2.971 0 0 1 1.031 2.189"/><path d="m8.36 13.682l1.879 1.88a.71.71 0 0 0 1.01 0l3.787-3.787"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileCheckFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.83 7.424a3.01 3.01 0 0 0-.15-.38a3.85 3.85 0 0 0-.92-1.22l-3-2.72a4.22 4.22 0 0 0-2.85-1.1h-5.7A5 5 0 0 0 3 6.864v10.3a5 5 0 0 0 3.31 4.53a4.74 4.74 0 0 0 1.92.3h7.56a5 5 0 0 0 5.21-4.86v-8.57a3.63 3.63 0 0 0-.17-1.14m-5.06 5.06l-3.78 3.79a2 2 0 0 1-.56.37a1.73 1.73 0 0 1-1.88-.38l-1.87-1.87a1.004 1.004 0 0 1 1.42-1.42l1.67 1.68l3.59-3.59a1 1 0 0 1 1.41 1.42m.29-5.3a.82.82 0 0 1-.84-.83v-2.53A45.668 45.668 0 0 1 19 7.184z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileCross(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M14.186 2.753v3.596c0 .487.194.955.54 1.3a1.85 1.85 0 0 0 1.306.539h4.125"/><path stroke-linejoin="round" d="M20.25 8.568v8.568a4.251 4.251 0 0 1-1.362 2.97a4.283 4.283 0 0 1-3.072 1.14h-7.59a4.294 4.294 0 0 1-3.1-1.124a4.265 4.265 0 0 1-1.376-2.986V6.862a4.25 4.25 0 0 1 1.362-2.97a4.283 4.283 0 0 1 3.072-1.14h5.714a3.503 3.503 0 0 1 2.361.905l2.96 2.722a2.971 2.971 0 0 1 1.031 2.189"/><path stroke-miterlimit="10" d="m14.51 11.513l-5.03 5.032m-.001-5.021l5.032 5.032"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileCrossFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.65 7.048c-.4-.998-3.13-3.133-3.89-3.932a4.154 4.154 0 0 0-2.39-1.067c.09-.11-6 0-6.16 0a5.01 5.01 0 0 0-3.622 1.336A4.989 4.989 0 0 0 3 6.898v10.277a4.98 4.98 0 0 0 1.605 3.501a5.001 5.001 0 0 0 3.625 1.319h7.56a5.01 5.01 0 0 0 3.623-1.337A4.99 4.99 0 0 0 21 17.145v-8.55a3.784 3.784 0 0 0-.35-1.547m-5.43 8.79c.93.908-.51 2.325-1.42 1.407l-1.81-1.806c-.82.619-2.11 2.874-3.22 1.796c-1.11-1.078 1.22-2.424 1.8-3.213c-.62-.818-2.89-2.115-1.8-3.213c1.09-1.097 2.4 1.178 3.22 1.796c.83-.628 2.13-2.893 3.23-1.806c1.1 1.088-1.19 2.395-1.81 3.223zm.81-8.401c-1.67 0-.94-2.624-1.09-3.702c.3.106.577.272.81.489c.31.33 3.26 2.864 3.4 3.213z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileDownload(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M14.186 2.753v3.596c0 .487.194.955.54 1.3a1.85 1.85 0 0 0 1.306.539h4.125"/><path stroke-linejoin="round" d="M20.25 8.568v8.568a4.251 4.251 0 0 1-1.362 2.97a4.283 4.283 0 0 1-3.072 1.14h-7.59a4.294 4.294 0 0 1-3.1-1.124a4.265 4.265 0 0 1-1.376-2.986V6.862a4.25 4.25 0 0 1 1.362-2.97a4.283 4.283 0 0 1 3.072-1.14h5.714a3.503 3.503 0 0 1 2.361.905l2.96 2.722a2.971 2.971 0 0 1 1.031 2.189"/><path stroke-miterlimit="10" d="M12 17.273v-6.774"/><path stroke-linejoin="round" d="m8.894 14.42l2.665 2.666a.622.622 0 0 0 .882 0l2.665-2.665"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileDownloadFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.68 7.014a3.85 3.85 0 0 0-.92-1.22l-3-2.72a4.15 4.15 0 0 0-2.39-1.07H8.21A5 5 0 0 0 3 6.864v10.3a5 5 0 0 0 3.31 4.53a4.74 4.74 0 0 0 1.92.3h7.56a5 5 0 0 0 5.21-4.86v-8.57a3.75 3.75 0 0 0-.32-1.55m-4.84 8.08l-2.66 2.67a1.69 1.69 0 0 1-.53.35c-.133.06-.275.097-.42.11a.9.9 0 0 1-.4 0a1.308 1.308 0 0 1-.42-.11a1.69 1.69 0 0 1-.53-.35l-2.66-2.67a1 1 0 0 1 1.41-1.41l1.4 1.4v-4.61a1 1 0 1 1 2 0v4.61l1.4-1.4a1 1 0 0 1 1.41 1.41m.22-7.69a1.08 1.08 0 0 1-1.09-1.08v-2.65c.28.152.551.322.81.51l3 2.73c.166.14.307.305.42.49z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.68 7.03a3.85 3.85 0 0 0-.92-1.22l-3-2.72a4.13 4.13 0 0 0-2.42-1.079a.41.41 0 0 0-.19 0H8.21A5.002 5.002 0 0 0 3 6.87v10.297a4.997 4.997 0 0 0 3.31 4.528a4.741 4.741 0 0 0 1.92.3h7.56A5.002 5.002 0 0 0 21 17.136V8.569a3.748 3.748 0 0 0-.32-1.54m-4.62.39a1.08 1.08 0 0 1-1.09-1.08V3.69c.66.16 3.23 2.8 3.79 3.24c.166.139.307.304.42.49z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileMinus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M14.186 2.753v3.596c0 .487.194.955.54 1.3a1.85 1.85 0 0 0 1.306.539h4.125"/><path stroke-linejoin="round" d="M20.25 8.568v8.568a4.251 4.251 0 0 1-1.362 2.97a4.283 4.283 0 0 1-3.072 1.14h-7.59a4.294 4.294 0 0 1-3.1-1.124a4.265 4.265 0 0 1-1.376-2.986V6.862a4.25 4.25 0 0 1 1.362-2.97a4.283 4.283 0 0 1 3.072-1.14h5.714a3.503 3.503 0 0 1 2.361.905l2.96 2.722a2.971 2.971 0 0 1 1.031 2.189"/><path stroke-miterlimit="10" d="M8.442 14.103h7.116"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileMinusFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.68 7.014a3.85 3.85 0 0 0-.92-1.22l-3-2.72a4.15 4.15 0 0 0-2.39-1.07H8.21A5 5 0 0 0 3 6.864v10.3a5 5 0 0 0 3.31 4.53a4.74 4.74 0 0 0 1.92.3h7.56a5 5 0 0 0 5.21-4.86v-8.57a3.75 3.75 0 0 0-.32-1.55m-5.09 8.06H8.47a1 1 0 0 1 0-2h7.12a1 1 0 0 1 0 2m.47-7.67a1.08 1.08 0 0 1-1.09-1.08v-2.65c.66.16 3.23 2.8 3.79 3.24c.166.14.307.305.42.49z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilePlus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M14.186 2.753v3.596c0 .487.194.955.54 1.3a1.85 1.85 0 0 0 1.306.539h4.125"/><path stroke-linejoin="round" d="M20.25 8.568v8.568a4.251 4.251 0 0 1-1.362 2.97a4.283 4.283 0 0 1-3.072 1.14h-7.59a4.294 4.294 0 0 1-3.1-1.124a4.265 4.265 0 0 1-1.376-2.986V6.862a4.25 4.25 0 0 1 1.362-2.97a4.283 4.283 0 0 1 3.072-1.14h5.714a3.503 3.503 0 0 1 2.361.905l2.96 2.722a2.971 2.971 0 0 1 1.031 2.189"/><path stroke-miterlimit="10" d="M11.57 10.424v7.116m-3.55-3.55h7.117"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilePlusFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.65 7.03a3.849 3.849 0 0 0-.92-1.22c-1.66-1.36-3.41-3.738-5.39-3.798a.39.39 0 0 0-.19 0H8.21A5.001 5.001 0 0 0 3 6.87v10.296a4.997 4.997 0 0 0 3.31 4.528a4.741 4.741 0 0 0 1.92.3h7.56A5.002 5.002 0 0 0 21 17.136V8.569a3.746 3.746 0 0 0-.35-1.54m-5.51 7.997h-2.57v2.549a1 1 0 1 1-2 0v-2.55H8a1 1 0 1 1 0-1.999h2.55v-2.569a1 1 0 0 1 2 0v2.57h2.57a1 1 0 0 1 .02 1.948zm.89-7.557a1.08 1.08 0 0 1-1.09-1.08V3.761c.67.13 3.24 2.789 3.79 3.219c.166.14.308.305.42.49z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileQuestionMark(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M14.186 2.753v3.596c0 .487.194.955.54 1.3a1.85 1.85 0 0 0 1.306.539h4.125"/><path stroke-linejoin="round" d="M20.25 8.568v8.568a4.251 4.251 0 0 1-1.362 2.97a4.283 4.283 0 0 1-3.072 1.14h-7.59a4.294 4.294 0 0 1-3.1-1.124a4.265 4.265 0 0 1-1.376-2.986V6.862a4.25 4.25 0 0 1 1.362-2.97a4.283 4.283 0 0 1 3.072-1.14h5.714a3.503 3.503 0 0 1 2.361.905l2.96 2.722a2.971 2.971 0 0 1 1.031 2.189"/><path stroke-miterlimit="10" d="M9.862 11.48a1.834 1.834 0 0 1 2-1.04a1.78 1.78 0 0 1 1.304.93a1.544 1.544 0 0 1-.9 2.124a1.14 1.14 0 0 0-.734 1.03v.425"/><path stroke-linejoin="round" d="M11.499 17.295h.004"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileQuestionMarkFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.65 7.041c-.4-.999-3.12-3.136-3.89-3.936a4.132 4.132 0 0 0-2.42-1.078c0-.06-5.95 0-6.13 0a5.005 5.005 0 0 0-4.77 2.937A4.99 4.99 0 0 0 3 6.88v10.29a4.99 4.99 0 0 0 3.285 4.528a5.005 5.005 0 0 0 1.945.297h7.56a5.005 5.005 0 0 0 4.77-2.937A4.99 4.99 0 0 0 21 17.14V8.58a3.79 3.79 0 0 0-.35-1.539M11.5 18.32a1.001 1.001 0 0 1-1-1a.999.999 0 0 1 1.707-.706a.999.999 0 0 1-.707 1.706m2.61-5.264a2.537 2.537 0 0 1-1.52 1.418c0 .68-.22 1.518-1.06 1.508c-1.81-.18-1-3.046.38-3.386a.55.55 0 0 0 .35-.77a.87.87 0 0 0-1.49.11a1.002 1.002 0 0 1-1.82-.838c1.38-3.157 6.48-1.379 5.16 1.918zm1.92-5.584c-1.67 0-.94-2.627-1.09-3.726a25.06 25.06 0 0 1 4.21 3.726z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileRecords(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8.593 3.217H4.698A1.948 1.948 0 0 0 2.75 5.164v13.633c0 1.075.872 1.947 1.948 1.947h3.895a1.948 1.948 0 0 0 1.947-1.947V5.164a1.948 1.948 0 0 0-1.947-1.947"/><path d="M6.645 17.379a1.503 1.503 0 1 0 0-3.007a1.503 1.503 0 0 0 0 3.007M10.54 7.93l3.116 11.685a1.948 1.948 0 0 0 2.386 1.373l3.768-.974a1.947 1.947 0 0 0 1.373-2.386L17.658 4.385a1.947 1.947 0 0 0-2.386-1.373l-3.758 1.003c-.406.111-.764.35-1.023.682"/><path d="M16.665 17.241a1.502 1.502 0 1 0 0-3.004a1.502 1.502 0 0 0 0 3.004"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileRecordsFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.861 17.43l-3.52-13.24a2.73 2.73 0 0 0-1.26-1.64a2.65 2.65 0 0 0-2-.26l-3.77 1a2.06 2.06 0 0 0-.57.25a2.66 2.66 0 0 0-2.15-1.07h-3.9a2.69 2.69 0 0 0-2.69 2.69v13.63a2.69 2.69 0 0 0 2.69 2.7h3.9a2.7 2.7 0 0 0 2.7-2.7v-5.12l1.64 6.14a2.63 2.63 0 0 0 1.26 1.63c.407.236.87.36 1.34.36c.236.003.471-.024.7-.08l3.77-1a2.669 2.669 0 0 0 1.9-3.3zm-15.26.26a1.81 1.81 0 1 1 1.8-1.81a1.81 1.81 0 0 1-1.8 1.78zm10-.11a1.84 1.84 0 1 1 1.84-1.84a1.84 1.84 0 0 1-1.82 1.81z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileThree(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.186 2.753v3.596c0 .487.194.955.54 1.3a1.85 1.85 0 0 0 1.306.539h4.125"/><path d="M20.25 8.568v8.568a4.251 4.251 0 0 1-1.362 2.97a4.283 4.283 0 0 1-3.072 1.14h-7.59a4.294 4.294 0 0 1-3.1-1.124a4.265 4.265 0 0 1-1.376-2.986V6.862a4.25 4.25 0 0 1 1.362-2.97a4.283 4.283 0 0 1 3.072-1.14h5.714a3.503 3.503 0 0 1 2.361.905l2.96 2.722a2.971 2.971 0 0 1 1.031 2.189M7.647 7.647h3.265M7.647 12h8.706m-8.706 4.353h8.706"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileTwo(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.186 2.753v3.596c0 .487.195.955.54 1.3a1.85 1.85 0 0 0 1.306.539h4.125"/><path d="M20.25 8.568v8.568a4.251 4.251 0 0 1-1.362 2.97a4.283 4.283 0 0 1-3.072 1.14h-7.59a4.294 4.294 0 0 1-3.1-1.124a4.265 4.265 0 0 1-1.376-2.986V6.862a4.25 4.25 0 0 1 1.362-2.97a4.283 4.283 0 0 1 3.072-1.14h5.714a3.503 3.503 0 0 1 2.361.905l2.96 2.722a2.971 2.971 0 0 1 1.031 2.189M7.647 7.647h3.265M7.647 12h8.706m-8.706 4.353h8.706"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileTwoFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.68 7.014a3.85 3.85 0 0 0-.92-1.22l-3-2.72a4.15 4.15 0 0 0-2.39-1.07H8.21A5 5 0 0 0 3 6.864v10.3a5 5 0 0 0 3.31 4.53a4.738 4.738 0 0 0 1.92.3h7.56a4.998 4.998 0 0 0 5.21-4.86v-8.57a3.75 3.75 0 0 0-.32-1.55m-13-.4h3.26a1 1 0 0 1 0 2H7.68a1 1 0 1 1 0-2m8.7 10.71h-8.7a1 1 0 1 1 0-2h8.7a1 1 0 0 1 0 1.98zm0-4.35h-8.7a1 1 0 1 1 0-2h8.7a1 1 0 1 1 0 2m-.32-5.57a1.08 1.08 0 0 1-1.09-1.08v-2.65c.66.16 3.23 2.8 3.79 3.24a2 2 0 0 1 .42.49z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileUpload(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M14.186 2.753v3.596c0 .487.194.955.54 1.3a1.85 1.85 0 0 0 1.306.539h4.125"/><path stroke-linejoin="round" d="M20.25 8.568v8.568a4.251 4.251 0 0 1-1.362 2.97a4.283 4.283 0 0 1-3.072 1.14h-7.59a4.294 4.294 0 0 1-3.1-1.124a4.265 4.265 0 0 1-1.376-2.986V6.862a4.25 4.25 0 0 1 1.362-2.97a4.283 4.283 0 0 1 3.072-1.14h5.714a3.503 3.503 0 0 1 2.361.905l2.96 2.722a2.971 2.971 0 0 1 1.031 2.189"/><path stroke-miterlimit="10" d="M12 10.499v6.774"/><path stroke-linejoin="round" d="m15.106 13.35l-2.665-2.665a.62.62 0 0 0-.882 0l-2.665 2.666"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileUploadFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.65 7.048c-.4-.998-3.12-3.133-3.89-3.932a4.154 4.154 0 0 0-2.39-1.067c.11-.11-6 0-6.16 0a5.01 5.01 0 0 0-3.622 1.336A4.989 4.989 0 0 0 3 6.898v10.277a4.98 4.98 0 0 0 1.605 3.501a5.001 5.001 0 0 0 3.625 1.319h7.56a5.01 5.01 0 0 0 3.623-1.337A4.99 4.99 0 0 0 21 17.145v-8.55a3.784 3.784 0 0 0-.35-1.547m-4.84 6.984c-1 1.058-2.15-.878-2.81-1.397v4.61a.997.997 0 0 1-1 .998a1.002 1.002 0 0 1-1-.998v-4.61l-1.4 1.397c-.9.928-2.34-.509-1.41-1.407l2.66-2.654a1.642 1.642 0 0 1 2.3 0l2.66 2.654a.998.998 0 0 1 0 1.417zm.22-6.605c-1.67 0-.94-2.744-1.09-3.702a23.762 23.762 0 0 1 4.21 3.702z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Filter(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10" stroke-width="1.5" d="M21.25 12H8.895m-4.361 0H2.75m18.5 6.607h-5.748m-4.361 0H2.75m18.5-13.214h-3.105m-4.361 0H2.75m13.214 2.18a2.18 2.18 0 1 0 0-4.36a2.18 2.18 0 0 0 0 4.36Zm-9.25 6.607a2.18 2.18 0 1 0 0-4.36a2.18 2.18 0 0 0 0 4.36Zm6.607 6.608a2.18 2.18 0 1 0 0-4.361a2.18 2.18 0 0 0 0 4.36Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 18.605a.75.75 0 0 1-.75.75h-5.1a2.93 2.93 0 0 1-5.66 0H2.75a.75.75 0 1 1 0-1.5h7.74a2.93 2.93 0 0 1 5.66 0h5.1a.75.75 0 0 1 .75.75m0-13.21a.75.75 0 0 1-.75.75H18.8a2.93 2.93 0 0 1-5.66 0H2.75a.75.75 0 1 1 0-1.5h10.39a2.93 2.93 0 0 1 5.66 0h2.45a.74.74 0 0 1 .75.75m0 6.6a.74.74 0 0 1-.75.75H9.55a2.93 2.93 0 0 1-5.66 0H2.75a.75.75 0 1 1 0-1.5h1.14a2.93 2.93 0 0 1 5.66 0h11.7a.75.75 0 0 1 .75.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/><path stroke-linecap="round" stroke-linejoin="round" d="M17.75 8.805h-1.437m-3.594 0H6.25m11.5 6.39h-6.469m-3.593 0H6.25m3.234 1.797a1.797 1.797 0 1 0 0-3.594a1.797 1.797 0 0 0 0 3.594m5.032-6.39a1.797 1.797 0 1 0 0-3.594a1.797 1.797 0 0 0 0 3.594"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterSquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 2h-6.5A6.76 6.76 0 0 0 2 8.75v6.5A6.75 6.75 0 0 0 8.75 22h6.5A6.75 6.75 0 0 0 22 15.25v-6.5A6.76 6.76 0 0 0 15.25 2m2.5 13.95h-5.84a2.54 2.54 0 0 1-4.86 0h-.8a.75.75 0 1 1 0-1.5h.8a2.54 2.54 0 0 1 4.86 0h5.84a.75.75 0 1 1 0 1.5m0-6.39h-.8a2.54 2.54 0 0 1-4.86 0H6.25a.75.75 0 0 1 0-1.5h5.84a2.54 2.54 0 0 1 4.86 0h.8a.75.75 0 0 1 0 1.5"/><path fill="currentColor" d="M10.53 15.2a1 1 0 1 1-1.997-.1a1 1 0 0 1 1.997.1m5.03-6.39a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterTwo(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10" stroke-width="1.5" d="m19.795 4.413l-5.074 7.852a1.79 1.79 0 0 0-.287.987v4.788a1.229 1.229 0 0 1-.678 1.09l-3.662 1.826a.356.356 0 0 1-.528-.322v-7.382a1.803 1.803 0 0 0-.287-.987L4.205 4.413A.976.976 0 0 1 5.112 3h13.776a.975.975 0 0 1 .907 1.412Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilterTwoFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.62 3.89a1.74 1.74 0 0 1-.19.86v.07l-5.07 7.86a1.06 1.06 0 0 0-.17.57v4.79a2 2 0 0 1-1.09 1.76l-3.66 1.83a.999.999 0 0 1-.51.12a1.1 1.1 0 0 1-.55-.16a1.05 1.05 0 0 1-.4-.41a1.11 1.11 0 0 1-.13-.57v-7.35c0-.209-.06-.413-.17-.59L3.6 4.82a1.84 1.84 0 0 1-.22-.93c.016-.3.113-.59.28-.84a1.75 1.75 0 0 1 .64-.6a1.85 1.85 0 0 1 .87-.2h13.75c.29-.006.575.062.83.2c.268.136.493.344.65.6c.146.256.221.545.22.84"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fingerprint(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.278 14.611a2.86 2.86 0 0 0-.308-1.476c-1.231-2.482 1.712-5.187 4.213-4.132c3.51 1.439 2.931 8.016 1.391 10.82m.666-16.268C8.587.354.631 7.278 3.267 14.165"/><path d="M7.061 7.662A6.847 6.847 0 0 1 10.13 5.94a6.811 6.811 0 0 1 3.513.033a6.849 6.849 0 0 1 3.036 1.78a6.923 6.923 0 0 1 1.763 3.062m-6.639.893c.615 1.659.802 3.448.542 5.199a10.9 10.9 0 0 1-2.192 4.343M5.916 10.305A5.811 5.811 0 0 0 6.1 13.88a3.818 3.818 0 0 1-.1 2.565a3.787 3.787 0 0 1-1.698 1.915M17.95 5.119c1.723.893 2.857 3.35 3.3 4.963m-2.39 3.4a14.587 14.587 0 0 1-.27 5.162m-9.521-1.352a5.4 5.4 0 0 1-1.983 2.394"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FingerprintMinimal(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12.036 12.158s-.986 4.736 4.492 6.993"/><path d="M13.752 20.952c-2.871-1.621-5.086-4.726-4.725-8.127a4.047 4.047 0 0 1 1.737-3.306a2.935 2.935 0 0 1 3.9 1.261c.942 1.388.275 3.179.964 4.641c1.06 1.801 3.75 1.346 4.767-.297c1.526-2.352.715-5.943-.843-8.105c-2.447-3.337-6.043-5.107-11.14-3.338c-4.789 1.632-7.162 8.381-4.64 12.863"/><path d="M7.982 19.002a9 9 0 0 1-1.864-3.484a8.902 8.902 0 0 1-.183-3.933a6.443 6.443 0 0 1 .91-2.45A6.562 6.562 0 0 1 8.662 7.23a6.279 6.279 0 0 1 3.578-.869c.811.025 1.61.208 2.349.537a6.16 6.16 0 0 1 1.96 1.38a8.283 8.283 0 0 1 1.692 5.51"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FireA(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4.21 14.434c0 6.328 5.843 6.816 7.79 6.816s7.79-.488 7.79-6.816c0-2.869-2.819-3.772-3.895-7.79c-6.816 7.79-5.842-3.894-5.842-3.894S4.21 8.592 4.21 14.434"/><path d="M8.02 13.694c-.422 2.17 1.345 3.862 3.024 4.189"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FireAfill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.78 10.451a10.594 10.594 0 0 1-2.176-4.009a.74.74 0 0 0-1.258-.293c-2.293 2.624-3.248 2.498-3.492 2.39c-.976-.439-1.249-3.687-1.083-5.745a.713.713 0 0 0-.42-.722a.702.702 0 0 0-.819.147c-.253.243-6.067 6.135-6.067 12.221C3.465 21.455 9.991 22 12 22c2.01 0 8.535-.546 8.535-7.56a6.233 6.233 0 0 0-1.756-3.989m-7.736 8.418h-.195a4.711 4.711 0 0 1-3.794-5.345a.993.993 0 1 1 1.95.38a2.75 2.75 0 0 0 2.263 3.053a.98.98 0 1 1-.185 1.95z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FireB(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20.17 13.344c0-4.368-3.953-4.23-3.953-10.594C9.763 4.341 9.23 7.365 9.23 12.988c-1.463.149-2.797-2.273-3.637-3.597c-3.874 5.07-1.235 11.859 6.67 11.859a7.906 7.906 0 0 0 7.907-7.906"/><path d="M16.938 12.988a5.11 5.11 0 0 1-5.93 4.942"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FireBfill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.046 8.744a8.375 8.375 0 0 1-2.096-5.99a.738.738 0 0 0-.28-.589a.72.72 0 0 0-.648-.14c-6.609 1.628-7.457 4.872-7.537 9.983A10.75 10.75 0 0 1 6.49 9.432l-.27-.419a.729.729 0 0 0-.599-.35a.689.689 0 0 0-.629.29a8.205 8.205 0 0 0-1.168 8.585C5.231 20.373 8.295 22 12.248 22a8.705 8.705 0 0 0 6.11-2.535a8.565 8.565 0 0 0 2.535-6.11a6.908 6.908 0 0 0-1.847-4.611m-1.727 6.818a6.07 6.07 0 0 1-3.883 3.244a6.29 6.29 0 0 1-1.607.21a5.501 5.501 0 0 1-.998-.08a1.012 1.012 0 0 1 .33-1.997c.578.09 1.17.06 1.736-.09a4.184 4.184 0 0 0 1.547-.808a3.993 3.993 0 0 0 1.068-1.378c.257-.53.393-1.109.4-1.697a.998.998 0 0 1 1.996 0a5.992 5.992 0 0 1-.609 2.596z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FirstAidKit(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3.765 11.48a1 1 0 0 1-1-1V8.9a2.56 2.56 0 0 1 2.54-2.57h13.36a2.56 2.56 0 0 1 2.57 2.57v1.55a1 1 0 0 1-1 1m-16.47.03l.36 7.28a2.2 2.2 0 0 0 2.18 2h11.36a2.2 2.2 0 0 0 2.18-2l.36-7.28"/><path d="M15.065 6.33V4.27a1.002 1.002 0 0 0-.3-.73a1 1 0 0 0-.72-.3h-4.12a1 1 0 0 0-.72.3a1 1 0 0 0-.3.73v2.06m7.52 6.016v1.967a.492.492 0 0 1-.492.492h-2.458v2.458a.492.492 0 0 1-.492.491h-1.966a.492.492 0 0 1-.492-.491v-2.458H8.067a.492.492 0 0 1-.492-.492v-1.967a.492.492 0 0 1 .492-.491h2.458V9.397a.492.492 0 0 1 .492-.492h1.966a.492.492 0 0 1 .492.492v2.458h2.458a.492.492 0 0 1 .492.491"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FirstAidKitFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.73 7.636a3.302 3.302 0 0 0-3.051-2.048h-2.872V4.29a1.655 1.655 0 0 0-.13-.69a1.628 1.628 0 0 0-.38-.569a1.74 1.74 0 0 0-.58-.39a1.713 1.713 0 0 0-.67-.14H9.923a1.433 1.433 0 0 0-.69.14c-.21.09-.4.22-.56.38a1.698 1.698 0 0 0-.39.58a1.656 1.656 0 0 0-.14.679v1.308H5.291c-.875.01-1.71.364-2.327.985A3.324 3.324 0 0 0 2 8.904v1.579a1.745 1.745 0 0 0 .52 1.238c.146.149.32.268.51.35l.34 6.752c.063.732.398 1.413.941 1.908a3.01 3.01 0 0 0 2.001.769h11.376a2.934 2.934 0 0 0 2.921-2.707l.33-6.732a1.722 1.722 0 0 0 .93-.938c.088-.213.132-.44.131-.67V8.914a3.311 3.311 0 0 0-.27-1.278M9.644 4.27a.36.36 0 0 1 0-.1l.07-.09l.08-.06h4.352l.09.07a.219.219 0 0 1 0 .08a.2.2 0 0 1 0 .11v1.308H9.574zm7.524 10.048a1.24 1.24 0 0 1-.36.87a1.22 1.22 0 0 1-.881.369h-1.71v1.708a1.243 1.243 0 0 1-.361.869a1.22 1.22 0 0 1-.88.37h-2.001a1.251 1.251 0 0 1-1.251-1.24v-1.707H8.033a1.252 1.252 0 0 1-1.154-.762a1.227 1.227 0 0 1-.097-.477V12.32a1.226 1.226 0 0 1 .37-.879c.236-.229.552-.358.881-.36h1.7V9.425a1.227 1.227 0 0 1 1.252-1.239h2c.33 0 .646.129.88.36c.23.236.358.55.361.879v1.708h1.711c.33 0 .646.129.88.36c.23.235.36.55.36.878z"/><path fill="currentColor" d="M15.667 12.61v1.449h-2.201a.75.75 0 0 0-.75.749v2.197h-1.451v-2.197a.737.737 0 0 0-.75-.75H8.312V12.61h2.201a.75.75 0 0 0 .75-.749V9.664h1.451v2.197a.759.759 0 0 0 .75.75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flag(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.382 14.72s1.089-1.088 4.353-1.088c3.265 0 3.265 2.177 6.53 2.177a11.26 11.26 0 0 0 4.353-1.088V3.838a11.253 11.253 0 0 1-4.353 1.088C12 4.926 12 2.75 8.735 2.75c-3.264 0-4.353 1.088-4.353 1.088m0 17.412V3.838"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlagFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.372 3.86v10.88a.75.75 0 0 1-.43.68a11.851 11.851 0 0 1-4.64 1.16a5.9 5.9 0 0 1-3.71-1.21a4.45 4.45 0 0 0-2.85-1a7.16 7.16 0 0 0-3.61.73v6.13a.75.75 0 1 1-1.5 0V3.86a.769.769 0 0 1 0-.15a.76.76 0 0 1 .31-.47c.38-.32 1.73-1.22 4.78-1.22a5.87 5.87 0 0 1 3.68 1.22a4.46 4.46 0 0 0 2.85 1a10.26 10.26 0 0 0 4-1a.74.74 0 0 1 .72.05a.73.73 0 0 1 .4.57"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Focus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M20.325 8.3V6.45a2.775 2.775 0 0 0-2.775-2.775h-2.775m0 16.65h2.775a2.775 2.775 0 0 0 2.775-2.775V15.7m-16.65 0v1.85a2.775 2.775 0 0 0 2.775 2.775h2.775m0-16.65H6.45A2.775 2.775 0 0 0 3.675 6.45V8.3"/><path d="M12 8a4 4 0 1 1 0 8a4 4 0 0 1 0-8Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FocusFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.228 9.06a.743.743 0 0 1-.743-.742V6.489a1.976 1.976 0 0 0-1.98-1.977h-2.752a.743.743 0 0 1-.743-.741a.74.74 0 0 1 .742-.741h2.753a3.518 3.518 0 0 1 2.466 1.027A3.506 3.506 0 0 1 21 6.52v1.829a.751.751 0 0 1-.772.711m-2.723 11.89h-2.752a.743.743 0 0 1-.526-1.265a.743.743 0 0 1 .525-.217h2.753a1.982 1.982 0 0 0 1.98-1.977v-1.809a.741.741 0 0 1 1.268-.524a.74.74 0 0 1 .217.524v1.829a3.506 3.506 0 0 1-1.03 2.461A3.518 3.518 0 0 1 17.476 21zm-8.267 0H6.485a3.498 3.498 0 0 1-2.465-1.025A3.486 3.486 0 0 1 3 17.461v-1.779a.74.74 0 0 1 1.268-.524a.74.74 0 0 1 .217.524v1.829a1.975 1.975 0 0 0 1.98 1.977h2.753a.743.743 0 0 1 .742.741a.74.74 0 0 1-.742.741zM3.743 9.06A.743.743 0 0 1 3 8.317V6.489c0-.923.367-1.81 1.02-2.464A3.498 3.498 0 0 1 6.485 3h2.753a.743.743 0 0 1 .742.741a.74.74 0 0 1-.742.742H6.485a1.982 1.982 0 0 0-1.98 1.977v1.828a.75.75 0 0 1-.762.771m12.94 2.916a4.69 4.69 0 0 1-1.069 2.978a4.704 4.704 0 0 1-5.853 1.16a4.688 4.688 0 0 1-.392-8.045a4.709 4.709 0 0 1 6.96 2.109c.236.57.357 1.181.354 1.798"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folder(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21.25 9.883v7.698a3.083 3.083 0 0 1-3.083 3.083H5.833a3.083 3.083 0 0 1-3.083-3.083V6.419a3.083 3.083 0 0 1 3.083-3.083h3.084a3.083 3.083 0 0 1 2.57 1.377l.873 1.326a1.748 1.748 0 0 0 1.449.77h4.358a3.084 3.084 0 0 1 3.083 3.074"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderCheck(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8.447 13.334l2 2a.757.757 0 0 0 1.076 0l4.03-4.03"/><path d="M21.25 9.883v7.698a3.083 3.083 0 0 1-3.083 3.083H5.833a3.083 3.083 0 0 1-3.083-3.083V6.419a3.083 3.083 0 0 1 3.083-3.083h3.084a3.083 3.083 0 0 1 2.57 1.377l.873 1.326a1.748 1.748 0 0 0 1.449.77h4.358a3.084 3.084 0 0 1 3.083 3.074"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderCheckFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.87 7.175a3.821 3.821 0 0 0-2.71-1.12h-4.35a1 1 0 0 1-.47-.11a1.001 1.001 0 0 1-.36-.32l-.87-1.33a3.801 3.801 0 0 0-3.2-1.71H5.83A3.82 3.82 0 0 0 2 6.415v11.16a3.84 3.84 0 0 0 3.83 3.84h12.33a3.86 3.86 0 0 0 3.84-3.84v-7.65a3.85 3.85 0 0 0-1.13-2.75m-4.61 4.83l-4 4a1.702 1.702 0 0 1-1.25.52a1.701 1.701 0 0 1-.67-.13a1.651 1.651 0 0 1-.58-.39l-2-2a1 1 0 0 1 1.41-1.41l1.83 1.83l3.86-3.86a.998.998 0 0 1 1.42 0a1.002 1.002 0 0 1-.02 1.44"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderCross(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="m14.678 11.067l-5.356 5.355m0-5.355l5.356 5.355"/><path stroke-linejoin="round" d="M21.25 9.883v7.698a3.083 3.083 0 0 1-3.083 3.083H5.833a3.083 3.083 0 0 1-3.083-3.083V6.419a3.083 3.083 0 0 1 3.083-3.083h3.084a3.083 3.083 0 0 1 2.57 1.377l.873 1.326a1.748 1.748 0 0 0 1.449.77h4.358a3.084 3.084 0 0 1 3.083 3.074"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderCrossFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.87 7.175a3.81 3.81 0 0 0-2.71-1.11h-4.35a.94.94 0 0 1-.47-.12a1 1 0 0 1-.36-.32l-.87-1.33a3.87 3.87 0 0 0-3.2-1.71H5.83A3.82 3.82 0 0 0 2 6.425v11.16a3.82 3.82 0 0 0 3.83 3.83h12.33a3.84 3.84 0 0 0 3.84-3.83v-7.7a3.89 3.89 0 0 0-1.13-2.71m-5.49 8.54a1 1 0 0 1-1.41 1.42l-2-2l-2 2a1 1 0 0 1-1.41-1.42l2-2l-2-2a1 1 0 0 1 0-1.41a1 1 0 0 1 1.41 0l2 2l2-2a1 1 0 0 1 1.41 0a1 1 0 0 1 0 1.41l-2 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 9.885v7.7a3.85 3.85 0 0 1-2.373 3.542a3.8 3.8 0 0 1-1.467.288H5.83A3.82 3.82 0 0 1 2 17.585V6.425a3.82 3.82 0 0 1 3.83-3.84h3.08a3.87 3.87 0 0 1 3.2 1.71l.87 1.33a1 1 0 0 0 .36.32a.94.94 0 0 0 .47.12h4.35a3.79 3.79 0 0 1 2.71 1.11A3.85 3.85 0 0 1 22 9.885"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderMinus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="M8.213 13.447h7.574"/><path stroke-linejoin="round" d="M21.25 9.883v7.698a3.083 3.083 0 0 1-3.083 3.083H5.833a3.083 3.083 0 0 1-3.083-3.083V6.419a3.083 3.083 0 0 1 3.083-3.083h3.084a3.083 3.083 0 0 1 2.57 1.377l.873 1.326a1.748 1.748 0 0 0 1.449.77h4.358a3.084 3.084 0 0 1 3.083 3.074"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderMinusFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.87 7.175a3.79 3.79 0 0 0-2.7-1.11h-4.36a.94.94 0 0 1-.47-.12a1.111 1.111 0 0 1-.36-.32l-.87-1.33a3.87 3.87 0 0 0-3.19-1.71H5.83A3.82 3.82 0 0 0 2 6.425v11.16a3.88 3.88 0 0 0 1.12 2.71a3.84 3.84 0 0 0 2.71 1.12h12.34a3.88 3.88 0 0 0 2.71-1.12a3.84 3.84 0 0 0 1.12-2.71v-7.7a3.81 3.81 0 0 0-1.13-2.71m-5.08 7.27H8.21a1 1 0 0 1 0-2h7.58a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOpen(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.375 16.047h9.25m4.625-6.164v7.698a3.083 3.083 0 0 1-3.083 3.083H5.833a3.083 3.083 0 0 1-3.083-3.083V6.419a3.083 3.083 0 0 1 3.083-3.083h3.084a3.083 3.083 0 0 1 2.57 1.377l.873 1.326a1.748 1.748 0 0 0 1.449.77h4.358a3.084 3.084 0 0 1 3.083 3.074"/><path d="m11.517 4.723l6.927-1.11a1.531 1.531 0 0 1 1.778 1.521v2.457"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOpenFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.99 7.325v-2.15a2.29 2.29 0 0 0-1.66-2.21a2.31 2.31 0 0 0-1-.05l-6.51 1a3.67 3.67 0 0 0-1.09-.87a3.8 3.8 0 0 0-1.81-.46H5.84A3.82 3.82 0 0 0 2 6.425v11.16a3.82 3.82 0 0 0 3.84 3.83h12.33a3.82 3.82 0 0 0 2.71-1.12a3.88 3.88 0 0 0 1.12-2.71v-7.7a3.73 3.73 0 0 0-1.01-2.56m-4.34 9.76H7.4a1 1 0 0 1 0-2h9.25a1 1 0 0 1 0 2m2.88-10.71a3.78 3.78 0 0 0-1.3-.23h-4.4a1 1 0 0 1-.47-.12a1 1 0 0 1-.35-.32l-.23-.34l5.8-.93a.77.77 0 0 1 .34 0a.73.73 0 0 1 .3.16a.6.6 0 0 1 .2.27a.71.71 0 0 1 .07.33z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderPlus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="M11.993 10.307v6.874m-3.43-3.437h6.874"/><path stroke-linejoin="round" d="M21.25 9.883v7.698a3.083 3.083 0 0 1-3.083 3.083H5.833a3.083 3.083 0 0 1-3.083-3.083V6.419a3.083 3.083 0 0 1 3.083-3.083h3.084a3.083 3.083 0 0 1 2.57 1.377l.873 1.326a1.748 1.748 0 0 0 1.449.77h4.358a3.084 3.084 0 0 1 3.083 3.074"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderPlusFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.88 7.175a3.83 3.83 0 0 0-2.71-1.11h-4.36a1 1 0 0 1-.47-.12a1.07 1.07 0 0 1-.35-.32l-.87-1.33a3.9 3.9 0 0 0-3.2-1.71H5.84a3.86 3.86 0 0 0-3.548 2.366A3.81 3.81 0 0 0 2 6.425v11.16a3.81 3.81 0 0 0 1.13 2.71a3.86 3.86 0 0 0 2.71 1.12h12.33a3.82 3.82 0 0 0 3.83-3.83v-7.7a3.8 3.8 0 0 0-1.12-2.71m-5.44 7.57H13v2.44a1 1 0 0 1-2 0v-2.44H8.57a1 1 0 0 1 0-2H11v-2.44a1 1 0 1 1 2 0v2.44h2.44a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderTwo(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.375 16.047h9.25m4.625-6.164v7.698a3.083 3.083 0 0 1-3.083 3.083H5.833a3.083 3.083 0 0 1-3.083-3.083V6.419a3.083 3.083 0 0 1 3.083-3.083h3.084a3.083 3.083 0 0 1 2.57 1.377l.873 1.326a1.748 1.748 0 0 0 1.449.77h4.358a3.084 3.084 0 0 1 3.083 3.074"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderTwoFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.88 7.175a3.84 3.84 0 0 0-2.71-1.12h-4.36a1 1 0 0 1-.46-.11a1 1 0 0 1-.36-.32l-.87-1.33a3.84 3.84 0 0 0-3.2-1.71H5.84A3.84 3.84 0 0 0 2 6.415v11.16a3.86 3.86 0 0 0 3.84 3.84h12.33a3.84 3.84 0 0 0 3.83-3.84v-7.65a3.8 3.8 0 0 0-1.12-2.75m-4.25 9.87H7.38a1 1 0 0 1 0-2h9.25a1 1 0 1 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gameboy(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17 2.75H7a3 3 0 0 0-3 3v12.5a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V5.75a3 3 0 0 0-3-3"/><path d="M16.5 5.56h-9c-.552 0-1 .383-1 .857v4.286c0 .473.448.857 1 .857h9c.552 0 1-.384 1-.857V6.417c0-.474-.448-.857-1-.857m-7.967 8.552v3.473m-1.737-1.737h3.46M12 19h2m1.136-2.904v.4m1.764-2.384v.4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GameboyFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 2H7a3.75 3.75 0 0 0-3.75 3.75v12.5A3.75 3.75 0 0 0 7 22h10a3.76 3.76 0 0 0 3.75-3.75V5.75A3.76 3.76 0 0 0 17 2m-6.75 14.6h-1v1a.75.75 0 1 1-1.5 0v-1h-1a.75.75 0 1 1 0-1.5h1v-1a.75.75 0 1 1 1.5 0v1h1a.75.75 0 1 1 0 1.5M14 19.75h-2a.75.75 0 1 1 0-1.5h2a.75.75 0 1 1 0 1.5m1.88-3.25a.75.75 0 1 1-1.5 0v-.4a.75.75 0 1 1 1.5 0zm1.77-2a.75.75 0 1 1-1.5 0v-.4a.75.75 0 1 1 1.5 0zm.6-4.47a1.69 1.69 0 0 1-1.75 1.6h-9a1.68 1.68 0 0 1-1.75-1.6V5.74A1.68 1.68 0 0 1 7.5 4.13h9a1.69 1.69 0 0 1 1.75 1.61z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GemA(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m18.612 10.82l-5.737-7.879a1.055 1.055 0 0 0-.382-.322a1.1 1.1 0 0 0-.986 0a1.055 1.055 0 0 0-.381.322l-5.738 7.88A1.999 1.999 0 0 0 5 12c0 .422.135.834.388 1.18l5.738 7.879c.098.135.229.246.38.322a1.101 1.101 0 0 0 .987 0c.152-.076.283-.187.381-.322l5.738-7.88a1.992 1.992 0 0 0 0-2.359"/><path d="M5.015 12.195L12 15.078l6.985-2.883M12 2.5v19"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GemAfill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m19.294 10.377l-5.74-7.896a1.913 1.913 0 0 0-.652-.55a1.603 1.603 0 0 0-.681-.181h-.301c-.238.01-.47.072-.681.18c-.26.13-.487.319-.662.551l-5.74 7.896A2.735 2.735 0 0 0 4.304 12c-.005.06-.005.12 0 .18a.39.39 0 0 0 0 .15a.33.33 0 0 0 0 .14c.057.414.216.806.461 1.143l5.751 7.896c.17.235.394.427.652.56c.214.1.445.162.68.181h.301a1.833 1.833 0 0 0 1.333-.741l5.741-7.886c.249-.347.41-.75.471-1.172a.14.14 0 0 0 0-.07c.01-.06.01-.12 0-.18a.68.68 0 0 0 0-.201a2.745 2.745 0 0 0-.4-1.623m-1.212 2.364l-.701 1.002l-4.56 1.884v4.379l-.49.67a.2.2 0 0 1-.11.091a.341.341 0 0 1-.311 0a.24.24 0 0 1-.11-.09l-.481-.661v-4.389l-4.57-1.884l-.7-1.002A1.232 1.232 0 0 1 5.807 12a1.002 1.002 0 0 1 0-.3l5.47 2.264v-9.94l.472-.65a.29.29 0 0 1 .12-.101a.41.41 0 0 1 .16 0a.4.4 0 0 1 .15 0a.22.22 0 0 1 .11.1l.492.671v9.92l5.46-2.265a.761.761 0 0 1 .05.301c.01.263-.063.523-.21.741"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GemB(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m18.612 10.819l-5.737-7.88a1.055 1.055 0 0 0-.382-.323a1.1 1.1 0 0 0-.986 0a1.055 1.055 0 0 0-.381.323l-5.738 7.88A2 2 0 0 0 5 11.999c0 .422.135.834.388 1.18l5.738 7.88c.098.135.229.246.38.322a1.1 1.1 0 0 0 .987 0c.152-.076.283-.187.381-.322l5.738-7.88c.253-.346.388-.758.388-1.18a2 2 0 0 0-.388-1.18"/><path d="M5.015 12.194L12 15.077l6.985-2.883M12 21.5v-6.423"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GemBfill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m19.276 10.394l-5.728-7.877a1.849 1.849 0 0 0-3 0l-5.727 7.877a2.71 2.71 0 0 0-.53 1.62a1.09 1.09 0 0 0 0 .18c0 .102.02.204.06.3c.057.411.215.803.46 1.139l5.728 7.877a2 2 0 0 0 .66.56a2 2 0 0 0 .68.18h.3a1.83 1.83 0 0 0 1.329-.74l5.728-7.877a2.62 2.62 0 0 0 .47-1.16a.14.14 0 0 0 0-.07a.603.603 0 0 0 0-.18a.683.683 0 0 0 0-.2a2.72 2.72 0 0 0-.43-1.629M6.76 13.713l4.558 1.88v4.378zm6.058 6.248v-4.369l4.548-1.879z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GemC(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m18.612 10.82l-5.737-7.881a1.055 1.055 0 0 0-.382-.323a1.1 1.1 0 0 0-.986 0a1.055 1.055 0 0 0-.381.323l-5.738 7.88A2 2 0 0 0 5 12a2 2 0 0 0 .388 1.18l5.738 7.881c.098.136.229.246.38.323a1.1 1.1 0 0 0 .987 0c.152-.077.283-.187.381-.323l5.738-7.88A2 2 0 0 0 19 12a2 2 0 0 0-.388-1.18"/><path d="M5.015 12.195L12 15.078l6.985-2.883"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GemCfill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m19.25 10.377l-5.732-7.882a1.77 1.77 0 0 0-.66-.55a1.85 1.85 0 0 0-2.31.55l-5.732 7.882a2.73 2.73 0 0 0-.53 1.62c-.005.06-.005.12 0 .18c0 .103.02.206.06.3c.058.413.216.804.46 1.14l5.732 7.883a1.81 1.81 0 0 0 1.49.75a1.87 1.87 0 0 0 .83-.19a1.93 1.93 0 0 0 .65-.55l5.732-7.882c.248-.347.41-.748.47-1.17a.141.141 0 0 0 0-.07a.544.544 0 0 0 0-.18a.684.684 0 0 0 0-.2a2.74 2.74 0 0 0-.46-1.631m-7.002 10.242a.27.27 0 0 1-.11.1a.42.42 0 0 1-.31 0a.36.36 0 0 1-.11-.1l-5.001-6.921l5.001 2.07a.72.72 0 0 0 .57 0l5.001-2.07z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GemStone(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.238 9.153a1.102 1.102 0 0 1-.231.54l-8.14 10.476c-.151.19-.36.326-.596.386a.98.98 0 0 1-.473 0a1.101 1.101 0 0 1-.672-.397L2.986 9.693a1.102 1.102 0 0 1-.067-1.267l2.886-4.494l.088-.11a1.102 1.102 0 0 1 .838-.397h10.53a.937.937 0 0 1 .254 0a1.1 1.1 0 0 1 .672.474l2.886 4.494a1.1 1.1 0 0 1 .165.76"/><path d="m14.2 3.458l1.574 5.695l-3.503 11.39m-.473.012L7.777 9.153l2.016-5.695M2.754 9.153h18.484"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GemStoneFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9.367 19.08l-7.025-9.04h3.836zm5.987-9.04l-3.36 10.946L8.12 10.039zm6.28 0l-7.209 9.272l2.859-9.273zM22 8.206h-4.704l-1.43-5.192h2.652a.39.39 0 0 1 .257.183zm-6.597 0h-7.27l1.833-5.18h3.995zM8.011 3.039l-1.82 5.167H2L5.128 3.32l.097-.11a.501.501 0 0 1 .16-.123a.391.391 0 0 1 .17 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gif(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><path fill="currentColor" d="M10.988 11.798v.781a2.811 2.811 0 0 1-.351 1.45a2.399 2.399 0 0 1-.984.931c-.44.224-.93.336-1.423.325a3.108 3.108 0 0 1-1.581-.395a2.688 2.688 0 0 1-1.054-1.133A3.751 3.751 0 0 1 5.208 12c-.006-.473.07-.943.228-1.388a2.9 2.9 0 0 1 .633-1.028c.269-.283.595-.504.957-.65c.374-.15.775-.225 1.178-.22c.343-.002.684.051 1.01.159c.297.1.576.248.825.439c.24.19.443.422.598.685c.155.27.256.57.298.878H9.557a1.406 1.406 0 0 0-.175-.404a1.045 1.045 0 0 0-.29-.298a1.23 1.23 0 0 0-.387-.194a1.696 1.696 0 0 0-.483-.035c-.31-.01-.615.073-.878.237a1.599 1.599 0 0 0-.571.712c-.15.358-.223.745-.211 1.133c-.008.388.06.773.202 1.133c.123.287.324.533.58.712c.26.17.567.256.878.246c.27.008.539-.05.782-.167a1.15 1.15 0 0 0 .518-.492c.108-.206.165-.435.167-.668H8.283v-.992zm2.462-2.882v6.211a.088.088 0 0 1-.087.088h-1.177a.08.08 0 0 1-.065-.023a.08.08 0 0 1-.023-.065v-6.21a.08.08 0 0 1 .053-.085a.08.08 0 0 1 .035-.003h1.177a.088.088 0 0 1 .088.087m1.108 6.211v-6.21a.088.088 0 0 1 .088-.088h4.146v1.115h-2.758a.088.088 0 0 0-.088.088v1.344a.097.097 0 0 0 .088.088h2.512v1.115h-2.512a.088.088 0 0 0-.088.088v2.46a.08.08 0 0 1-.088.088h-1.177a.089.089 0 0 1-.087-.009a.088.088 0 0 1-.036-.079"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16.625 3.75h-9.25c-2.554 0-4.625 2.052-4.625 4.583v7.334c0 2.531 2.07 4.583 4.625 4.583h9.25c2.554 0 4.625-2.052 4.625-4.583V8.333c0-2.531-2.07-4.583-4.625-4.583"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GifFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.62 3H7.37A5.36 5.36 0 0 0 2 8.33v7.33A5.36 5.36 0 0 0 7.37 21h9.25A5.37 5.37 0 0 0 22 15.66V8.33A5.36 5.36 0 0 0 16.62 3m-5.64 9.58a2.761 2.761 0 0 1-.35 1.45a2.44 2.44 0 0 1-1 .93a3 3 0 0 1-1.42.32a3.09 3.09 0 0 1-1.58-.39a2.65 2.65 0 0 1-1.06-1.14A3.8 3.8 0 0 1 5.18 12c0-.473.078-.942.23-1.39a2.92 2.92 0 0 1 .64-1a2.73 2.73 0 0 1 1-.65a3.27 3.27 0 0 1 2.19-.06c.299.1.58.249.83.44a2.39 2.39 0 0 1 .89 1.56H9.58a1.2 1.2 0 0 0-.85-.89a1.562 1.562 0 0 0-.48 0a1.48 1.48 0 0 0-1.45.95a2.74 2.74 0 0 0-.21 1.13a3 3 0 0 0 .2 1.14c.123.287.324.532.58.71c.262.168.569.252.88.24c.269.008.536-.047.78-.16a1.2 1.2 0 0 0 .52-.5a1.37 1.37 0 0 0 .16-.66h-1.4v-1h2.7zm2.47 2.54a.092.092 0 0 1 0 .07h-1.27V8.98h1.27a.09.09 0 0 1 0 .06zm5.34-5.18h-2.77a.07.07 0 0 0 0 .06v1.34a.16.16 0 0 0 0 .06h2.57v1.12h-2.57a.09.09 0 0 0 0 .06v2.5h-1.26V8.82a.07.07 0 0 1 .08-.08h3.95z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gift(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.806 12v7.194A2.056 2.056 0 0 0 6.86 21.25h10.28a2.056 2.056 0 0 0 2.055-2.056V12m.513-5.139H4.292c-.852 0-1.542.69-1.542 1.542v2.055c0 .852.69 1.542 1.542 1.542h15.416c.852 0 1.542-.69 1.542-1.542V8.403c0-.852-.69-1.542-1.542-1.542m-12.785 0C6.018 5.71 5.833 2.75 8.917 2.75c3.494 0 3.032 4.111 3.083 4.111c.051 0-.36-4.111 3.083-4.111c3.084 0 2.878 2.96 1.974 4.111M12 21.25V6.861"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GiftFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.69 6.094h-1.47a3.56 3.56 0 0 0-.16-2.5a3 3 0 0 0-3-1.61A3.38 3.38 0 0 0 12 3.574a3.477 3.477 0 0 0-.47-.59a3.54 3.54 0 0 0-2.63-1a3.07 3.07 0 0 0-3 1.62a3.58 3.58 0 0 0-.16 2.49H4.29A2.3 2.3 0 0 0 2 8.384v2a2.31 2.31 0 0 0 2.06 2.29v6.55a2.82 2.82 0 0 0 2.8 2.81h10.28a2.82 2.82 0 0 0 2.8-2.81v-6.5a2.31 2.31 0 0 0 2.06-2.29v-2a2.3 2.3 0 0 0-2.31-2.34m-4.63-2.61a1.67 1.67 0 0 1 1.63.74a2.269 2.269 0 0 1-.1 1.84h-3.86c.22-2.58 1.78-2.58 2.31-2.58zm-7.83.78a1.67 1.67 0 0 1 1.67-.78a2.09 2.09 0 0 1 1.56.55a3.22 3.22 0 0 1 .75 2.06H7.32a2.29 2.29 0 0 1-.11-1.87zm-3.75 6.17v-2a.79.79 0 0 1 .79-.79h7v3.58h-7a.79.79 0 0 1-.81-.79zm17 0a.789.789 0 0 1-.79.8h-7v-3.64h7a.79.79 0 0 1 .79.79z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Github(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.963 2.382C.554 2.621-1.82 17.93 8.852 21.602c.498.093.684-.219.684-.478v-1.68c-2.79.601-3.38-1.317-3.38-1.317a2.603 2.603 0 0 0-1.121-1.442c-.902-.612.072-.602.072-.602a2.074 2.074 0 0 1 1.536 1.038a2.167 2.167 0 0 0 2.924.819c.052-.5.275-.965.633-1.317c-2.23-.25-4.564-1.1-4.564-4.875a3.755 3.755 0 0 1 1.038-2.645a3.464 3.464 0 0 1 .103-2.634s.84-.26 2.76 1.037a9.584 9.584 0 0 1 5.02 0c1.908-1.276 2.748-1.038 2.748-1.038c.365.828.398 1.763.093 2.614a3.754 3.754 0 0 1 1.037 2.645c0 3.786-2.344 4.626-4.574 4.865c1.038.55.602 4.086.664 4.522c0 .259.176.57.695.477c10.642-3.64 8.152-18.97-3.257-19.209"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Globe(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 21.5a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19"/><path d="M12 21.5c2.332 0 4.222-4.253 4.222-9.5S14.332 2.5 12 2.5c-2.332 0-4.222 4.253-4.222 9.5s1.89 9.5 4.222 9.5M2.5 12h19"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlobeFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.478 21.439a10.103 10.103 0 0 1-6.504-8.573h4.32a17.54 17.54 0 0 0 2.184 8.573m7.505-8.573c-.182 5.465-2.174 9.208-3.973 9.208c-1.8 0-3.791-3.743-3.974-9.208zm6.043 0a10.055 10.055 0 0 1-6.514 8.573a17.434 17.434 0 0 0 2.194-8.573zm0-1.732h-4.32c.02-3-.735-5.952-2.194-8.573a10.054 10.054 0 0 1 6.514 8.573m-6.043 0H8.036c.183-5.475 2.174-9.208 3.974-9.208c1.799 0 3.781 3.733 3.973 9.208M8.478 2.61a17.483 17.483 0 0 0-2.184 8.572h-4.32A10.074 10.074 0 0 1 8.478 2.61"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Goals(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.66 10.66A1.9 1.9 0 0 0 10.1 12a1.9 1.9 0 0 0 1.9 1.9a1.9 1.9 0 0 0 1.34-.56"/><path d="M12 6.3a5.7 5.7 0 1 0 5.7 5.7"/><path d="M12 2.5a9.5 9.5 0 1 0 9.5 9.5m-5.975-3.524L12.95 11.05"/><path d="M20.94 5.844L17.7 6.3l.456-3.24a.19.19 0 0 0-.313-.161l-2.148 2.137a1.9 1.9 0 0 0-.513 1.72l.342 1.72l1.72.341a1.9 1.9 0 0 0 1.72-.513L21.1 6.157a.19.19 0 0 0-.162-.313"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoalsFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.992 14.65a2.65 2.65 0 0 1-1.87-4.53a.753.753 0 1 1 1.06 1.07a1.13 1.13 0 0 0-.34.81a1.14 1.14 0 0 0 2 .81a.76.76 0 0 1 1.07 0a.77.77 0 0 1 0 1.07a2.67 2.67 0 0 1-1.92.77"/><path fill="currentColor" d="M22.242 12a10.25 10.25 0 0 1-10.25 10.25a10.25 10.25 0 0 1 0-20.5a.75.75 0 0 1 .75.75v3.8a.76.76 0 0 1-.75.75a4.86 4.86 0 0 0-2.75.83a4.93 4.93 0 0 0-2.11 5.08a5 5 0 0 0 9.81-1a.741.741 0 0 1 .75-.75h3.8a.75.75 0 0 1 .75.79"/><path fill="currentColor" d="M21.882 6.13a1.06 1.06 0 0 1-.23.53l-2.17 2.17a2.57 2.57 0 0 1-1.11.66a2.45 2.45 0 0 1-.76.11a3.171 3.171 0 0 1-.53 0l-1.33-.26l-2.28 2.29a.79.79 0 0 1-.53.22a.75.75 0 0 1-.53-.22a.74.74 0 0 1 0-1.06l2.29-2.3l-.27-1.32a2.76 2.76 0 0 1 .06-1.28a2.57 2.57 0 0 1 .66-1.11l2.15-2.15a1 1 0 0 1 .53-.25a.89.89 0 0 1 .54.09a1 1 0 0 1 .4.37a1 1 0 0 1 .12.58l-.32 2.26l2.25-.31a.94.94 0 0 1 .93.51a1 1 0 0 1 .13.47"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Google(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.999 12.044a9.811 9.811 0 0 1-1.323 5.002a9.784 9.784 0 0 1-3.682 3.628a9.97 9.97 0 0 1-4.951 1.348a10.004 10.004 0 0 1-4.972-1.36a10.032 10.032 0 0 1-3.64-3.656A9.995 9.995 0 0 1 2 12.044c0-1.758.465-3.485 1.345-5.006a10.032 10.032 0 0 1 3.64-3.656a10.004 10.004 0 0 1 4.971-1.36c2.407.01 4.73.886 6.545 2.47a.657.657 0 0 1 0 .96l-1.852 1.856a.656.656 0 0 1-.872 0a6.041 6.041 0 0 0-3.734-1.413a6.039 6.039 0 0 0-5.264 3.053a6.137 6.137 0 0 0-.872 3.096a6.105 6.105 0 0 0 .829 3.042a6.038 6.038 0 0 0 5.306 3.01a6.06 6.06 0 0 0 3.036-.82a6.252 6.252 0 0 0 2.605-3.075h-4.564a.645.645 0 0 1-.646-.658v-2.578a.636.636 0 0 1 .646-.647h8.148a.645.645 0 0 1 .646.539c.064.392.093.79.086 1.187"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Handicapped(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.572 12.51h2.125a2.362 2.362 0 0 1 1.527.5c.432.345.731.829.848 1.37l.898 4.89c.04.174.04.355 0 .53a1.11 1.11 0 0 1-.24.47a1.095 1.095 0 0 1-.418.32c-.163.076-.34.114-.52.11a1.175 1.175 0 0 1-.758-.25a1.12 1.12 0 0 1-.429-.67l-.698-4.08H15.13"/><path d="M8.135 10.5c-6.127 1.55-5.269 10.88 1.407 11c4.62.06 6.986-5.41 4.6-8.81"/><path d="M14.332 8.71V13a5.54 5.54 0 0 0-3.183-2.49v4.56a1.194 1.194 0 0 1-1.188 1.2a1.196 1.196 0 0 1-1.197-1.2l-.599-4.38a2.004 2.004 0 0 1 .828-1.88l1.597-1.13c1.247-.96 3.742-.82 3.742 1.03"/><path d="M13.195 7.06a2.272 2.272 0 0 0 2.102-1.408a2.284 2.284 0 0 0-.494-2.484A2.274 2.274 0 0 0 10.92 4.78c0 .604.24 1.184.667 1.612a2.273 2.273 0 0 0 1.608.668"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandicappedFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.726 19.115l-.909-4.912a3.186 3.186 0 0 0-1.118-1.797a3.125 3.125 0 0 0-1.997-.66h-1.637V8.703a2.297 2.297 0 0 0-.39-1.308a.81.81 0 0 0 .18-.1a2.996 2.996 0 0 0 1.288-3.105a3.085 3.085 0 0 0-.829-1.548a3.045 3.045 0 0 0-1.547-.828a2.995 2.995 0 0 0-3.605 2.995c-.005.69.228 1.36.66 1.897c-.243.1-.47.231-.68.39L8.555 8.221A2.716 2.716 0 0 0 7.447 9.95a6.34 6.34 0 0 0 2.076 12.3h.08a6.06 6.06 0 0 0 5.422-3.325a6.479 6.479 0 0 0 .748-2.466h.48l.599 3.485c.092.438.34.829.699 1.098c.326.26.73.404 1.148.41c.29.014.577-.044.838-.17c.269-.115.505-.293.69-.52a1.86 1.86 0 0 0 .399-.788c.092-.277.126-.569.1-.859m-6.99-.918a4.543 4.543 0 0 1-4.093 2.525h-.07a4.852 4.852 0 0 1-1.996-9.235l.489 3.624c.02.495.222.965.569 1.318a1.936 1.936 0 0 0 1.388.57a1.906 1.906 0 0 0 1.358-.57c.36-.368.564-.862.569-1.377v-3.425a4.783 4.783 0 0 1 1.777 1.737a4.992 4.992 0 0 1 .02 4.833z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hash(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.886 21L9.443 3m5.114 18l2.557-18M3 8.855h18M3 15.247h18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadphoneMute(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3.518 20.537a2.626 2.626 0 0 1-.768-1.866v-2.604a2.644 2.644 0 0 1 5.239 0m7.982 2.604a2.645 2.645 0 0 0 5.279 0v-2.604a2.644 2.644 0 0 0-5.279 0z"/><path d="M2.75 18.671v-6.595A9.23 9.23 0 0 1 18.506 5.55m1.746 2.305a9.34 9.34 0 0 1 .998 4.22v6.596M3 21.056L20.96 3.095"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadphoneMuteFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.488 2.556a.75.75 0 0 0-1.06 0l-1.95 1.95a10 10 0 0 0-10.32-1.66a10 10 0 0 0-6.16 9.22v6.6a3.5 3.5 0 0 0 .51 1.81a.74.74 0 0 0 0 1.06a.709.709 0 0 0 .53.22a.74.74 0 0 0 .53-.22l.52-.52l4.41-4.41l.09-.09l10.49-10.44l2.46-2.46a.75.75 0 0 0-.05-1.06m-13.2 12.14a3.43 3.43 0 0 0-.69-.83a3.47 3.47 0 0 0-4.1-.26v-1.53a8.48 8.48 0 0 1 13.92-6.5zm13.71-2.62v6.6c.005.05.005.1 0 .15a3.38 3.38 0 0 1-3.37 3.08a3.39 3.39 0 0 1-3.39-3.19v-2.64a3.4 3.4 0 0 1 1.06-2.31a3.49 3.49 0 0 1 4.22-.33v-1.36a8.72 8.72 0 0 0-.92-3.88a.75.75 0 0 1 .33-1a.77.77 0 0 1 1 .33a10.19 10.19 0 0 1 1.07 4.55"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HealthCircle(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.5 12.89H7l3-5l4 9l3-5h4.43"/><path d="M12 21.5a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HealthCircleFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.24 11.33a10.197 10.197 0 0 0-3.343-6.93a10.28 10.28 0 0 0-7.247-2.644a10.274 10.274 0 0 0-7.047 3.138a10.191 10.191 0 0 0-2.853 7.143v.23c.011.484.058.967.14 1.444a10.205 10.205 0 0 0 3.72 6.308a10.283 10.283 0 0 0 13.691-.804a10.194 10.194 0 0 0 2.949-6.7v-.468c0-.16 0-.488-.01-.717m-1.53 1.594h-3.142l-2.702 4.504a.998.998 0 0 1-.86.478h-.06a1.004 1.004 0 0 1-.86-.588l-3.212-7.214l-2 3.338a1 1 0 0 1-.861.479H3.47a8.357 8.357 0 0 1-.21-1.884v-.11h3.182l2.71-4.493a.999.999 0 0 1 .921-.478a1.003 1.003 0 0 1 .85.588l3.222 7.214l2.001-3.338a.999.999 0 0 1 .86-.489h3.682c.05.367.073.737.07 1.106a6.59 6.59 0 0 1-.05.887"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HealthSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/><path stroke-linecap="round" stroke-linejoin="round" d="M2.75 12.5H7l3-5l4 9l3-5h4.25"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HealthSquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 2h-6.5A6.76 6.76 0 0 0 2 8.75v6.5A6.75 6.75 0 0 0 8.75 22h6.5A6.75 6.75 0 0 0 22 15.25v-6.5A6.76 6.76 0 0 0 15.25 2m5.25 10.5h-2.94l-2.7 4.52a1 1 0 0 1-.86.48h-.06a1 1 0 0 1-.86-.59L9.87 9.67l-2 3.35a1 1 0 0 1-.86.48h-3.5v-2h2.93l2.71-4.51a1 1 0 0 1 1.77.11l3.22 7.24l2-3.35a1 1 0 0 1 .86-.49h3.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heaphone(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8.036 15.964a2.643 2.643 0 1 0-5.286 0v2.643a2.643 2.643 0 0 0 5.286 0zm7.928 2.643a2.643 2.643 0 1 0 5.286 0v-2.643a2.643 2.643 0 0 0-5.286 0z"/><path d="M21.25 18.607V12a9.25 9.25 0 1 0-18.5 0v6.607"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeaphoneFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.998 11.908v6.6a3.4 3.4 0 1 1-6.79 0v-2.64a3.39 3.39 0 0 1 5.29-2.81v-1.15a8.5 8.5 0 1 0-17 0v1.15a3.33 3.33 0 0 1 1.89-.58a3.39 3.39 0 0 1 3.39 3.39v2.64a3.39 3.39 0 1 1-6.78 0v-6.6a10 10 0 0 1 20 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heart(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 7.23c-1.733-3.924-5.764-4.273-7.641-2.562c-1.529 1.373-2.263 4.665-.867 7.695C5.9 17.573 12 20.309 12 20.309s6.101-2.736 8.508-7.946c1.396-3.03.662-6.322-.867-7.695C17.764 2.957 13.733 3.306 12 7.229"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.19 12.683c-2.5 5.41-8.62 8.2-8.88 8.32a.848.848 0 0 1-.62 0c-.25-.12-6.38-2.91-8.88-8.32c-1.55-3.37-.69-7 1-8.56a4.93 4.93 0 0 1 4.36-1.05a6.16 6.16 0 0 1 3.78 2.62a6.15 6.15 0 0 1 3.79-2.62a4.93 4.93 0 0 1 4.36 1.05c1.78 1.56 2.65 5.19 1.09 8.56"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartHealth(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 7.194c-1.73-3.92-5.76-4.23-7.64-2.56c-1.53 1.33-2.26 4.66-.87 7.69c2.41 5.21 8.51 8 8.51 8s6.1-2.74 8.51-7.95c1.39-3 .66-6.32-.87-7.69c-1.88-1.72-5.91-1.41-7.64 2.51"/><path d="M3.34 11.964H8l3 3l3-6l2 3h4.66"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartHealthFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.13 4.155a5 5 0 0 0-4.39-1.07A6 6 0 0 0 12 5.665a6 6 0 0 0-3.72-2.58a5.09 5.09 0 0 0-4.4 1c-1.58 1.38-2.45 4.44-1.46 7.54c.112.342.246.676.4 1c.04.075.077.152.11.23c2.57 5.24 8.51 8 8.77 8.13a.672.672 0 0 0 .31.07a.702.702 0 0 0 .31-.07c.25-.11 6.25-2.85 8.8-8.15l.08-.17c.158-.34.295-.691.41-1.05c.94-3 .08-6.06-1.48-7.46m-.31 7.93c-.14.314-.3.618-.48.91h-3.31a1 1 0 0 1-.83-.45l-1.05-1.56l-2.23 4.46a1 1 0 0 1-.73.54h-.16a1 1 0 0 1-.71-.3l-2.71-2.7H4.7a10.595 10.595 0 0 1-.5-1a6.336 6.336 0 0 1-.38-1h4.21a.999.999 0 0 1 .71.29l2 2l2.38-4.76a1 1 0 0 1 .84-.55a1 1 0 0 1 .89.44l1.7 2.56h3.7a6.572 6.572 0 0 1-.43 1.12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M12 9.402c-1.11-2.513-3.69-2.736-4.893-1.64c-.978.879-1.448 2.987-.555 4.927C8.094 16.025 12 17.777 12 17.777s3.907-1.752 5.448-5.088c.894-1.94.424-4.048-.555-4.928C15.691 6.666 13.11 6.89 12 9.401"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartSquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 2h-6.5A6.76 6.76 0 0 0 2 8.75v6.5A6.76 6.76 0 0 0 8.75 22h6.5A6.76 6.76 0 0 0 22 15.25v-6.5A6.76 6.76 0 0 0 15.25 2m1.65 11a10.12 10.12 0 0 1-4.58 4.29a.77.77 0 0 1-.64 0A10.12 10.12 0 0 1 7.1 13a4.16 4.16 0 0 1 .62-4.69a2.8 2.8 0 0 1 2.49-.61a3.25 3.25 0 0 1 1.79 1a3.28 3.28 0 0 1 1.79-1a2.83 2.83 0 0 1 2.5.61a4.18 4.18 0 0 1 .61 4.65z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Home(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m19.633 7.11l-6.474-4.02a2.228 2.228 0 0 0-2.362 0L4.324 7.133A2.228 2.228 0 0 0 3.31 9.362l1.67 10.027a2.228 2.228 0 0 0 2.228 1.86h9.582a2.229 2.229 0 0 0 2.229-1.86l1.67-10.027a2.228 2.228 0 0 0-1.058-2.251"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeCheck(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m19.633 7.11l-6.474-4.02a2.228 2.228 0 0 0-2.362 0L4.324 7.133A2.228 2.228 0 0 0 3.31 9.362l1.67 10.027a2.228 2.228 0 0 0 2.228 1.86h9.582a2.229 2.229 0 0 0 2.229-1.86l1.67-10.027a2.228 2.228 0 0 0-1.058-2.251"/><path d="m8.175 12.908l2.153 2.153a.814.814 0 0 0 1.158 0l4.34-4.339"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeCheckFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.179 7.76a3.06 3.06 0 0 0-1.16-1.28l-6.47-4a3 3 0 0 0-3.16 0l-6.47 4a3 3 0 0 0-1.36 3l1.67 10a3 3 0 0 0 2.93 2.49h9.61a3.24 3.24 0 0 0 2-.7a2.94 2.94 0 0 0 1-1.79l1.68-10a3 3 0 0 0-.27-1.72m-4.65 3.68l-4.34 4.33a1.81 1.81 0 0 1-2 .4a1.81 1.81 0 0 1-.59-.4l-2.15-2.15a1.001 1.001 0 1 1 1.41-1.41l2 2l4.21-4.21a1.004 1.004 0 0 1 1.42 1.42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeCross(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m19.633 7.11l-6.474-4.02a2.228 2.228 0 0 0-2.362 0L4.324 7.133A2.228 2.228 0 0 0 3.31 9.362l1.67 10.027a2.228 2.228 0 0 0 2.228 1.86h9.582a2.229 2.229 0 0 0 2.229-1.86l1.67-10.027a2.228 2.228 0 0 0-1.058-2.251"/><path stroke-miterlimit="10" d="m15.238 10.186l-6.475 6.475m0-6.461l6.475 6.475"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeCrossFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.187 7.763a2.93 2.93 0 0 0-1.16-1.28l-6.47-4a3 3 0 0 0-3.15 0l-6.48 4a3 3 0 0 0-1.12 1.29a2.92 2.92 0 0 0-.23 1.7l1.67 10a2.9 2.9 0 0 0 1 1.79a3 3 0 0 0 1.9.7h9.61a3.19 3.19 0 0 0 2-.7a3 3 0 0 0 1-1.79l1.67-10a3 3 0 0 0-.24-1.71m-5.24 8.21a1.002 1.002 0 0 1-.701 1.71a.999.999 0 0 1-.709-.29l-2.54-2.54l-2.52 2.52a1.001 1.001 0 1 1-1.42-1.41l2.53-2.52l-2.53-2.53a1 1 0 1 1 1.42-1.41l2.52 2.52l2.54-2.54a1 1 0 1 1 1.41 1.42l-2.55 2.54z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.444 9.473l-1.67 10a3 3 0 0 1-1 1.79a3.21 3.21 0 0 1-2 .7h-9.61a3 3 0 0 1-2.93-2.49l-1.67-10a3 3 0 0 1 .23-1.7a3 3 0 0 1 1.12-1.29l6.48-4a3 3 0 0 1 3.15 0l6.47 4c.507.3.911.747 1.16 1.28a3 3 0 0 1 .27 1.71"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeFour(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m19.632 7.106l-6.466-4.018a2.226 2.226 0 0 0-2.359 0L4.341 7.106A2.226 2.226 0 0 0 3.33 9.33l1.669 10.016a2.225 2.225 0 0 0 2.226 1.858H9.76v-3.65a1.558 1.558 0 0 1 1.558-1.558h1.335a1.558 1.558 0 0 1 1.558 1.558v3.695h2.56a2.226 2.226 0 0 0 2.226-1.859l1.669-10.015a2.226 2.226 0 0 0-1.035-2.27M7.535 10.878h8.903"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeFourFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.188 7.768a3 3 0 0 0-1.15-1.31l-6.47-4a3 3 0 0 0-3.15 0l-6.47 4a3 3 0 0 0-1.12 1.28a3 3 0 0 0-.23 1.7l1.67 10a3 3 0 0 0 1 1.79a2.92 2.92 0 0 0 1.93.69h2.54a.74.74 0 0 0 .75-.75v-3.65a.82.82 0 0 1 .81-.81h1.33a.84.84 0 0 1 .58.24a.81.81 0 0 1 .23.57v3.72a.76.76 0 0 0 .75.75h2.59a3 3 0 0 0 2.94-2.49l1.67-10a2.94 2.94 0 0 0-.2-1.73m-4.74 4.1h-8.86a1 1 0 1 1 0-2h8.9a1 1 0 0 1 0 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeHeart(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m19.633 7.11l-6.474-4.02a2.228 2.228 0 0 0-2.362 0L4.324 7.133A2.228 2.228 0 0 0 3.31 9.362l1.67 10.027a2.228 2.228 0 0 0 2.228 1.86h9.582a2.229 2.229 0 0 0 2.229-1.86l1.67-10.027a2.228 2.228 0 0 0-1.058-2.251"/><path d="M12 10.818c-.894-2.024-2.973-2.203-3.94-1.32c-.789.707-1.168 2.405-.448 3.968C8.854 16.153 12 17.564 12 17.564s3.146-1.411 4.388-4.098c.72-1.563.341-3.26-.447-3.969c-.968-.882-3.047-.703-3.941 1.321"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeHeartFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.184 7.765a2.93 2.93 0 0 0-1.16-1.28l-6.46-4a3 3 0 0 0-3.16 0l-6.48 4a3 3 0 0 0-1.12 1.29a3 3 0 0 0-.23 1.7l1.67 10a3 3 0 0 0 2.93 2.49h9.63a3.17 3.17 0 0 0 1.95-.7a3 3 0 0 0 1-1.79l1.67-10a2.91 2.91 0 0 0-.24-1.71m-4.73 6a9.18 9.18 0 0 1-4.18 3.91a.54.54 0 0 1-.27.07a.52.52 0 0 1-.27-.07a9.13 9.13 0 0 1-4.18-3.91a3.79 3.79 0 0 1 .55-4.25a2.52 2.52 0 0 1 2.25-.55a3 3 0 0 1 1.65 1a3 3 0 0 1 1.65-1a2.52 2.52 0 0 1 2.25.55a3.82 3.82 0 0 1 .55 4.22z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomePlus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m19.633 7.11l-6.474-4.02a2.228 2.228 0 0 0-2.362 0L4.324 7.133A2.228 2.228 0 0 0 3.31 9.362l1.67 10.027a2.228 2.228 0 0 0 2.228 1.86h9.582a2.229 2.229 0 0 0 2.229-1.86l1.67-10.027a2.228 2.228 0 0 0-1.058-2.251"/><path stroke-miterlimit="10" d="M11.992 9.14v7.666m-3.825-3.825h7.667"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomePlusFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.201 7.76a3 3 0 0 0-1.17-1.28l-6.46-4a3 3 0 0 0-3.16 0l-6.47 4a3 3 0 0 0-1.36 3l1.67 10a3 3 0 0 0 2.93 2.49h9.61a3.1 3.1 0 0 0 1.95-.7a3 3 0 0 0 1-1.79l1.67-10a3 3 0 0 0-.21-1.72m-5.36 6.23h-2.84v2.82a1 1 0 0 1-2 0v-2.82h-2.8a1 1 0 0 1 0-2h2.82V9.15a1 1 0 0 1 2 0v2.84h2.84a1 1 0 0 1 0 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeSecurityLock(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m19.633 7.11l-6.474-4.02a2.228 2.228 0 0 0-2.362 0L4.324 7.133A2.228 2.228 0 0 0 3.31 9.362l1.67 10.027a2.228 2.228 0 0 0 2.228 1.86h9.582a2.229 2.229 0 0 0 2.229-1.86l1.67-10.027a2.228 2.228 0 0 0-1.058-2.251"/><path d="M14.611 11.538H9.39c-.721 0-1.306.54-1.306 1.208v3.623c0 .667.585 1.208 1.306 1.208h5.222c.721 0 1.306-.541 1.306-1.208v-3.623c0-.667-.585-1.208-1.306-1.208"/><path d="M9.585 11.538v-1.207a2.415 2.415 0 1 1 4.83 0v1.207"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeSecurityLockFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.183 7.765a2.93 2.93 0 0 0-1.16-1.28l-6.47-4a3 3 0 0 0-3.16 0l-6.47 4a3 3 0 0 0-1.12 1.29a2.92 2.92 0 0 0-.24 1.7l1.68 10a2.94 2.94 0 0 0 1 1.79a3 3 0 0 0 1.9.7h9.62a3.12 3.12 0 0 0 2-.7a2.94 2.94 0 0 0 1-1.79l1.68-10a3 3 0 0 0-.26-1.71m-4.73 8.49a1.92 1.92 0 0 1-2 1.87h-5a1.921 1.921 0 0 1-2-1.87v-3.5a1.89 1.89 0 0 1 1.43-1.8v-.5a3 3 0 0 1 .89-2.14a3.07 3.07 0 0 1 4.27 0a3 3 0 0 1 .89 2.14v.5a1.89 1.89 0 0 1 1.43 1.8z"/><path fill="currentColor" d="M13.613 10.455v.43h-3.19v-.43a1.59 1.59 0 0 1 .47-1.13a1.63 1.63 0 0 1 2.25 0a1.59 1.59 0 0 1 .47 1.13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeThree(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19.643 7.116L13.162 3.09a2.231 2.231 0 0 0-2.365 0L4.316 7.116a2.23 2.23 0 0 0-1.015 2.231l1.673 10.04a2.23 2.23 0 0 0 2.231 1.863H16.8a2.231 2.231 0 0 0 2.23-1.863l1.674-10.04a2.231 2.231 0 0 0-1.06-2.23"/><path d="M13.653 10.14h-3.347c-.924 0-1.673.748-1.673 1.672v3.347c0 .924.75 1.673 1.673 1.673h3.347c.924 0 1.673-.749 1.673-1.673v-3.347c0-.924-.749-1.673-1.673-1.673"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeThreeFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.176 7.755a2.94 2.94 0 0 0-1.15-1.27l-6.47-4a3 3 0 0 0-3.16 0l-6.49 4a2.93 2.93 0 0 0-1.12 1.29a3 3 0 0 0-.23 1.7l1.67 10a3 3 0 0 0 2.93 2.49h9.63a3.19 3.19 0 0 0 2-.7a3 3 0 0 0 1-1.79l1.67-10.06a2.9 2.9 0 0 0-.28-1.66m-6 7.08a1.92 1.92 0 0 1-1.92 1.92h-2.67a1.92 1.92 0 0 1-1.92-1.92v-2.67a1.92 1.92 0 0 1 1.92-1.92h2.67a1.92 1.92 0 0 1 1.92 1.92z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeTwo(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m19.633 7.11l-6.474-4.02a2.228 2.228 0 0 0-2.362 0L4.324 7.133A2.228 2.228 0 0 0 3.31 9.362l1.67 10.027a2.228 2.228 0 0 0 2.228 1.86h9.582a2.229 2.229 0 0 0 2.229-1.86l1.67-10.027a2.228 2.228 0 0 0-1.058-2.251M8.636 16.459h6.685"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomeTwoFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.205 7.765a2.93 2.93 0 0 0-1.16-1.28l-6.47-4a3 3 0 0 0-3.16 0l-6.47 4a3 3 0 0 0-1.12 1.29a2.92 2.92 0 0 0-.24 1.7l1.68 10a2.94 2.94 0 0 0 1 1.79a3 3 0 0 0 1.9.7h9.62a3 3 0 0 0 1.94-.7a2.9 2.9 0 0 0 1-1.79l1.68-10a3 3 0 0 0-.2-1.71m-5.86 9.7h-6.69a1 1 0 0 1 0-2h6.69a1 1 0 1 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HospitalCircle(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 21.5a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19"/><path d="M16.763 10.942v2.116a.53.53 0 0 1-.53.53h-2.645v2.645a.529.529 0 0 1-.53.53h-2.116a.53.53 0 0 1-.53-.53v-2.645H7.768a.529.529 0 0 1-.53-.53v-2.116a.529.529 0 0 1 .53-.53h2.645V7.768a.529.529 0 0 1 .53-.53h2.116a.529.529 0 0 1 .53.53v2.645h2.645a.53.53 0 0 1 .53.53"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HospitalCircleFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.75A10.25 10.25 0 1 0 22.25 12A10.26 10.26 0 0 0 12 1.75m5.51 11.31a1.28 1.28 0 0 1-1.28 1.28h-1.89v1.9a1.3 1.3 0 0 1-1.28 1.28h-2.12a1.27 1.27 0 0 1-.9-.38a1.29 1.29 0 0 1-.38-.9v-1.9H7.77a1.271 1.271 0 0 1-.91-.37a1.29 1.29 0 0 1-.37-.91v-2.11a1.29 1.29 0 0 1 .37-.91a1.27 1.27 0 0 1 .91-.37h1.89v-1.9a1.32 1.32 0 0 1 .37-.9a1.27 1.27 0 0 1 .91-.38H13a1.26 1.26 0 0 1 .9.37c.24.242.377.569.38.91v1.9h1.89c.34 0 .667.132.91.37c.239.243.372.57.37.91z"/><path fill="currentColor" d="M16 11.165v1.67h-2.42a.75.75 0 0 0-.75.75v2.43h-1.68v-2.43a.76.76 0 0 0-.75-.75H8v-1.67h2.42a.76.76 0 0 0 .75-.75v-2.43h1.68v2.43a.75.75 0 0 0 .75.75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HospitalPlus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21.25 9.944v4.112a1.028 1.028 0 0 1-1.028 1.027h-5.139v5.14a1.028 1.028 0 0 1-1.027 1.027H9.944a1.028 1.028 0 0 1-1.027-1.028v-5.139h-5.14a1.028 1.028 0 0 1-1.027-1.027V9.944a1.028 1.028 0 0 1 1.028-1.027h5.139v-5.14A1.028 1.028 0 0 1 9.944 2.75h4.112a1.028 1.028 0 0 1 1.027 1.028v5.139h5.14a1.028 1.028 0 0 1 1.027 1.027"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HospitalPlusFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 9.95v4.11a1.78 1.78 0 0 1-1.78 1.78h-4.39v4.39a1.731 1.731 0 0 1-.52 1.25a1.8 1.8 0 0 1-1.26.52H9.94a1.8 1.8 0 0 1-1.26-.52a1.81 1.81 0 0 1-.52-1.25v-4.39H3.78A1.78 1.78 0 0 1 2 14.06V9.95a1.78 1.78 0 0 1 1.78-1.78h4.38V3.78a1.8 1.8 0 0 1 1.103-1.646A1.75 1.75 0 0 1 9.94 2H14a1.8 1.8 0 0 1 1.26.52a1.768 1.768 0 0 1 .52 1.26v4.39h4.39c.472.003.924.19 1.26.52A1.78 1.78 0 0 1 22 9.95"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HospitalShield(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16.045 10.45v1.797a.45.45 0 0 1-.45.45h-2.247v2.247a.45.45 0 0 1-.45.45h-1.797a.45.45 0 0 1-.45-.45v-2.247H8.406a.45.45 0 0 1-.45-.45V10.45a.45.45 0 0 1 .45-.45h2.247V7.753a.45.45 0 0 1 .45-.449H12.9a.45.45 0 0 1 .45.45V10h2.246a.45.45 0 0 1 .45.45"/><path d="M20.672 11.89V6.61a1.928 1.928 0 0 0-1.32-1.831L14.438 3.14a7.805 7.805 0 0 0-4.876 0L4.648 4.778a1.927 1.927 0 0 0-1.32 1.83v5.28a7.709 7.709 0 0 0 3.603 6.524l4.048 2.544a1.927 1.927 0 0 0 2.042 0l4.047-2.544a7.708 7.708 0 0 0 3.604-6.523"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HospitalShieldFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.945 5.069a2.75 2.75 0 0 0-1.33-1l-4.92-1.64a8.58 8.58 0 0 0-5.35 0l-4.91 1.64a2.67 2.67 0 0 0-1.83 2.54v5.28a8.44 8.44 0 0 0 4 7.16l4 2.55a2.72 2.72 0 0 0 2.84 0l4-2.55a8.54 8.54 0 0 0 2.9-3.07a8.439 8.439 0 0 0 1.05-4.09v-5.28a2.66 2.66 0 0 0-.45-1.54m-4.13 7.2a1.2 1.2 0 0 1-1.2 1.2h-1.49v1.5a1.2 1.2 0 0 1-1.2 1.2h-1.8a1.191 1.191 0 0 1-1.2-1.2v-1.53h-1.5a1.25 1.25 0 0 1-.85-.35a1.21 1.21 0 0 1-.35-.85v-1.8c0-.316.126-.618.35-.84a1.21 1.21 0 0 1 .85-.36h1.5v-1.49a1.2 1.2 0 0 1 1.2-1.2h1.8a1.15 1.15 0 0 1 .84.35a1.2 1.2 0 0 1 .36.85v1.49h1.49a1.18 1.18 0 0 1 .85.36a1.15 1.15 0 0 1 .35.84z"/><path fill="currentColor" d="M15.315 10.769v1.2h-1.94a.75.75 0 0 0-.75.75v1.95h-1.2v-1.95a.76.76 0 0 0-.75-.75h-1.95v-1.2h1.95a.75.75 0 0 0 .75-.75v-1.94h1.2v1.94a.74.74 0 0 0 .75.75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HospitalSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/><path stroke-linecap="round" stroke-linejoin="round" d="M16.763 10.942v2.116a.53.53 0 0 1-.53.53h-2.645v2.645a.529.529 0 0 1-.53.53h-2.116a.53.53 0 0 1-.53-.53v-2.645H7.768a.529.529 0 0 1-.53-.53v-2.116a.529.529 0 0 1 .53-.53h2.645V7.768a.529.529 0 0 1 .53-.53h2.116a.529.529 0 0 1 .53.53v2.645h2.645a.529.529 0 0 1 .53.53"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HospitalSquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 2h-6.5A6.76 6.76 0 0 0 2 8.75v6.5A6.75 6.75 0 0 0 8.75 22h6.5A6.75 6.75 0 0 0 22 15.25v-6.5A6.76 6.76 0 0 0 15.25 2m2.26 11.06a1.28 1.28 0 0 1-1.28 1.28h-1.89v1.9a1.3 1.3 0 0 1-1.28 1.28h-2.12a1.271 1.271 0 0 1-.9-.38a1.29 1.29 0 0 1-.38-.9v-1.9H7.77a1.271 1.271 0 0 1-.91-.37a1.29 1.29 0 0 1-.37-.91v-2.11a1.29 1.29 0 0 1 .37-.91a1.27 1.27 0 0 1 .91-.37h1.89v-1.9a1.32 1.32 0 0 1 .37-.9a1.27 1.27 0 0 1 .91-.38h2.1a1.26 1.26 0 0 1 .9.37c.24.242.377.569.38.91v1.9h1.89c.34 0 .667.132.91.37c.238.243.371.57.37.91z"/><path fill="currentColor" d="M16.04 11.17v1.67h-2.42a.75.75 0 0 0-.75.75v2.43h-1.68v-2.43a.76.76 0 0 0-.75-.75h-2.4v-1.67h2.42a.76.76 0 0 0 .75-.75V7.99h1.68v2.43a.75.75 0 0 0 .75.75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HourGlass(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m18.77 19.071l-.113 1.696a.524.524 0 0 1-.473.483H5.852a.524.524 0 0 1-.514-.483l-.113-1.696a2.62 2.62 0 0 1 .442-1.655L9.336 12a.822.822 0 0 0 0-.946L5.943 6.553a2.713 2.713 0 0 1-.535-1.614V3.274a.524.524 0 0 1 .524-.524h12.17a.524.524 0 0 1 .523.524V4.94c0 .581-.188 1.147-.534 1.614l-3.371 4.42a.822.822 0 0 0 0 .945l3.628 5.457c.33.5.479 1.1.421 1.696m-9.393-.905h5.283"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HourGlassFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m18.968 16.99l-3.65-5.54l3.37-4.42a3.52 3.52 0 0 0 .69-2.07V3.27a1.27 1.27 0 0 0-.38-.9a1.251 1.251 0 0 0-.9-.37H5.928a1.26 1.26 0 0 0-1.27 1.27v1.67A3.45 3.45 0 0 0 5.348 7l3.36 4.58l-3.67 5.41a3.36 3.36 0 0 0-.56 2.13l.11 1.7A1.282 1.282 0 0 0 5.848 22h12.41a1.25 1.25 0 0 0 .78-.37a1.3 1.3 0 0 0 .37-.81l.11-1.68a3.32 3.32 0 0 0-.55-2.15m-4.31 2.2h-5.23a1 1 0 1 1 0-2h5.28a1 1 0 0 1 0 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IdCard(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.5 16.052V7.948a4.135 4.135 0 0 0-1.236-2.945a4.25 4.25 0 0 0-2.985-1.22H6.72a4.25 4.25 0 0 0-2.985 1.22A4.135 4.135 0 0 0 2.5 7.948v8.104c0 1.105.445 2.164 1.236 2.945a4.25 4.25 0 0 0 2.985 1.22H17.28c1.12 0 2.193-.44 2.985-1.22a4.136 4.136 0 0 0 1.236-2.945"/><path d="M8.552 12.14a2.054 2.054 0 1 0 0-4.108a2.054 2.054 0 0 0 0 4.108m3.081 3.828c0-.812-.324-1.59-.902-2.165a3.091 3.091 0 0 0-4.358 0a3.052 3.052 0 0 0-.902 2.165m9.097-7.049h3.594M14.568 12h1.54m-1.54 3.081h3.594"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IdCardFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.82 4.465a5.06 5.06 0 0 0-3.51-1.43H6.75a4.93 4.93 0 0 0-5 4.91v8.11a4.89 4.89 0 0 0 1.46 3.47a5 5 0 0 0 3.51 1.44h10.56a5 5 0 0 0 3.51-1.44a4.83 4.83 0 0 0 1.46-3.47v-8.11a4.869 4.869 0 0 0-1.43-3.48m-9.16 12.25a.75.75 0 0 1-.75-.75a2.29 2.29 0 0 0-.68-1.63a2.43 2.43 0 0 0-3.3 0a2.28 2.28 0 0 0-.68 1.63a.75.75 0 0 1-1.5 0a3.8 3.8 0 0 1 1.12-2.7a3.93 3.93 0 0 1 1.2-.85a2.77 2.77 0 0 1-1.32-2.37a2.8 2.8 0 1 1 5.6 0a2.77 2.77 0 0 1-1.28 2.37c.457.19.875.465 1.23.81a3.88 3.88 0 0 1 1.12 2.7a.74.74 0 0 1-.76.79m6.53-.88H14.6a.75.75 0 0 1 0-1.5h3.59a.75.75 0 1 1 0 1.5m-4.34-3.84a.75.75 0 0 1 .75-.75h1.54a.75.75 0 1 1 0 1.5H14.6a.74.74 0 0 1-.75-.75m4.34-2.33H14.6a.75.75 0 0 1 0-1.5h3.59a.75.75 0 0 1 0 1.5"/><path fill="currentColor" d="M8.58 11.385a1.3 1.3 0 1 0 0-2.6a1.3 1.3 0 0 0 0 2.6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Illustrator(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.41 9.48c.35 1.14.71 2.28 1.07 3.44H9.36c.36-1.16.7-2.3 1.05-3.44"/><path fill="currentColor" d="M17 2H7a5 5 0 0 0-5 5v10a5 5 0 0 0 5 5h10a5 5 0 0 0 5-5V7a5 5 0 0 0-5-5m-2.51 14.78h-1.56a.24.24 0 0 1-.27-.2c-.19-.56-.39-1.12-.58-1.69a.27.27 0 0 0-.32-.22H9.1a.27.27 0 0 0-.33.23c-.17.56-.36 1.1-.53 1.66a.26.26 0 0 1-.31.22H6.56c-.24 0-.26 0-.18-.25l2.16-6.23c.2-.57.4-1.14.58-1.71a1.88 1.88 0 0 0 .08-.52c0-.21.08-.26.27-.25h1.91c.14 0 .21 0 .26.19l2.46 6.93c.19.54.38 1.07.56 1.6c.08.22.07.24-.17.24m3.06-.17c0 .17-.07.22-.23.22h-1.44c-.2 0-.25-.08-.25-.26v-6.28c0-.2.06-.28.27-.27h1.41c.16 0 .24 0 .24.22c-.007 2.12-.007 4.243 0 6.37m-1-7.3a1 1 0 0 1-1.06-1.08a1.071 1.071 0 0 1 1.1-1.06a1 1 0 0 1 1 1.08a1 1 0 0 1-1 1.06z"/><path fill="currentColor" d="M11.48 12.92H9.36c.36-1.16.7-2.3 1-3.44c.4 1.14.76 2.28 1.12 3.44"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Image(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16.24 3.5h-8.5a5 5 0 0 0-5 5v7a5 5 0 0 0 5 5h8.5a5 5 0 0 0 5-5v-7a5 5 0 0 0-5-5"/><path d="m2.99 17l2.75-3.2a2.2 2.2 0 0 1 2.77-.27a2.2 2.2 0 0 0 2.77-.27l2.33-2.33a4 4 0 0 1 5.16-.43l2.49 1.93M7.99 10.17a1.66 1.66 0 1 0 0-3.32a1.66 1.66 0 0 0 0 3.32"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageCheck(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.25 13V8.5a5 5 0 0 0-5-5h-8.5a5 5 0 0 0-5 5v7a5 5 0 0 0 5 5h6.26"/><path d="m3.01 17l2.74-3.2a2.2 2.2 0 0 1 2.77-.27a2.2 2.2 0 0 0 2.77-.27l2.33-2.33a4 4 0 0 1 5.16-.43l2.47 1.91M8.01 10.17a1.66 1.66 0 1 0-.02-3.32a1.66 1.66 0 0 0 .02 3.32m8.49 7.759l1.407 1.407a.531.531 0 0 0 .757 0L21.5 16.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageCross(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M21.25 13V8.5a5 5 0 0 0-5-5h-8.5a5 5 0 0 0-5 5v7a5 5 0 0 0 5 5h6.26"/><path stroke-linejoin="round" d="m3.01 17l2.74-3.2a2.2 2.2 0 0 1 2.77-.27a2.2 2.2 0 0 0 2.77-.27l2.33-2.33a4 4 0 0 1 5.16-.43l2.47 1.91M8.01 10.17a1.66 1.66 0 1 0-.02-3.32a1.66 1.66 0 0 0 .02 3.32"/><path stroke-miterlimit="10" d="m21.5 16.5l-4 3.991m0-3.982l4 3.991"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageDownload(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M21.25 13V8.5a5 5 0 0 0-5-5h-8.5a5 5 0 0 0-5 5v7a5 5 0 0 0 5 5h6.26"/><path stroke-linejoin="round" d="m3.01 17l2.74-3.2a2.2 2.2 0 0 1 2.77-.27a2.2 2.2 0 0 0 2.77-.27l2.33-2.33a4 4 0 0 1 5.16-.43l2.47 1.91M8.01 10.17a1.66 1.66 0 1 0-.02-3.32a1.66 1.66 0 0 0 .02 3.32"/><path stroke-miterlimit="10" d="M18.707 20v-5"/><path stroke-linejoin="round" d="m16.414 17.895l1.967 1.967a.459.459 0 0 0 .652 0L21 17.895"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.25 2.75h-8.5A5.76 5.76 0 0 0 2 8.5v7a5.76 5.76 0 0 0 5.75 5.75h8.5A5.76 5.76 0 0 0 22 15.5v-7a5.76 5.76 0 0 0-5.75-5.75M8 6.1a2.41 2.41 0 1 1-.922 4.635A2.41 2.41 0 0 1 8.01 6.1zm12.5 6.68l-2.18-1.69a3.26 3.26 0 0 0-4.17.37l-2.33 2.33a3 3 0 0 1-3.72.36a1.48 1.48 0 0 0-.94-.24a1.46 1.46 0 0 0-.88.42l-2.43 2.84a4.25 4.25 0 0 1-.35-1.91l1.68-1.95a3 3 0 0 1 3.76-.41a1.43 1.43 0 0 0 1.82-.18l2.33-2.32a4.77 4.77 0 0 1 6.13-.51l1.28 1z"/><path fill="currentColor" d="M8.91 8.51a.91.91 0 1 1-1.82 0a.91.91 0 0 1 1.82 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageMinus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M21.25 13V8.5a5 5 0 0 0-5-5h-8.5a5 5 0 0 0-5 5v7a5 5 0 0 0 5 5h6.26"/><path stroke-linejoin="round" d="m3.01 17l2.74-3.2a2.2 2.2 0 0 1 2.77-.27a2.2 2.2 0 0 0 2.77-.27l2.33-2.33a4 4 0 0 1 5.16-.43l2.47 1.91M8.01 10.17a1.66 1.66 0 1 0-.02-3.32a1.66 1.66 0 0 0 .02 3.32"/><path stroke-miterlimit="10" d="M16.5 18.005h5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImagePlus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M21.25 13V8.5a5 5 0 0 0-5-5h-8.5a5 5 0 0 0-5 5v7a5 5 0 0 0 5 5h6.26"/><path stroke-linejoin="round" d="m3.01 17l2.74-3.2a2.2 2.2 0 0 1 2.77-.27a2.2 2.2 0 0 0 2.77-.27l2.33-2.33a4 4 0 0 1 5.16-.43l2.47 1.91M8.01 10.17a1.66 1.66 0 1 0-.02-3.32a1.66 1.66 0 0 0 .02 3.32"/><path stroke-miterlimit="10" d="M18.994 15.5v5M16.5 18.005h5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageQuestionMark(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M21.25 13V8.5a5 5 0 0 0-5-5h-8.5a5 5 0 0 0-5 5v7a5 5 0 0 0 5 5h6.26"/><path stroke-linejoin="round" d="m3.01 17l2.74-3.2a2.2 2.2 0 0 1 2.77-.27a2.2 2.2 0 0 0 2.77-.27l2.33-2.33a4 4 0 0 1 5.16-.43l2.47 1.91M8.01 10.17a1.66 1.66 0 1 0-.02-3.32a1.66 1.66 0 0 0 .02 3.32"/><path stroke-miterlimit="10" d="M16.92 15.89a1.598 1.598 0 0 1 1.743-.906a1.55 1.55 0 0 1 1.137.81a1.347 1.347 0 0 1-.784 1.851a.994.994 0 0 0-.64.898v.37"/><path stroke-linejoin="round" d="M18.347 20.958h.002"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ImageUpload(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M21.25 13V8.5a5 5 0 0 0-5-5h-8.5a5 5 0 0 0-5 5v7a5 5 0 0 0 5 5h6.26"/><path stroke-linejoin="round" d="m3.01 17l2.74-3.2a2.2 2.2 0 0 1 2.77-.27a2.2 2.2 0 0 0 2.77-.27l2.33-2.33a4 4 0 0 1 5.16-.43l2.47 1.91M8.01 10.17a1.66 1.66 0 1 0-.02-3.32a1.66 1.66 0 0 0 .02 3.32"/><path stroke-miterlimit="10" d="M18.707 15v5"/><path stroke-linejoin="round" d="m21 17.105l-1.967-1.967a.458.458 0 0 0-.652 0l-1.967 1.967"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Inbox(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M2.75 12H6a2 2 0 0 1 2 2a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2a2 2 0 0 1 2-2h3.25"/><path d="M15.25 2.75h-6.5a6 6 0 0 0-6 6v6.5a6 6 0 0 0 6 6h6.5a6 6 0 0 0 6-6v-6.5a6 6 0 0 0-6-6Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxCheck(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M2.75 12H6a2 2 0 0 1 2 2a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2a2 2 0 0 1 2-2h3.25"/><path d="M21.25 10.375v4.875a6 6 0 0 1-6 6h-6.5a6 6 0 0 1-6-6v-6.5a6 6 0 0 1 6-6h4.875"/><path stroke-linejoin="round" d="m16.25 5.582l1.407 1.407a.532.532 0 0 0 .757 0l2.836-2.836"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxCheckFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 10.37v4.88A6.76 6.76 0 0 1 15.25 22h-6.5A6.76 6.76 0 0 1 2 15.25v-6.5A6.76 6.76 0 0 1 8.75 2h4.87a.75.75 0 1 1 0 1.5H8.75A5.25 5.25 0 0 0 3.5 8.75v2.5H6a2.718 2.718 0 0 1 1.94.8c.512.52.803 1.22.81 1.95A1.25 1.25 0 0 0 10 15.25h4a1.219 1.219 0 0 0 .88-.37a1.26 1.26 0 0 0 .37-.88A2.73 2.73 0 0 1 18 11.25h2.5v-.88a.75.75 0 1 1 1.5 0"/><path fill="currentColor" d="M18.03 7.9a1.38 1.38 0 0 1-.49-.1a1.62 1.62 0 0 1-.42-.28l-1.4-1.41a.75.75 0 1 1 1.06-1.06l1.25 1.26l2.69-2.69a.75.75 0 0 1 1.06 1.06l-2.84 2.84a1.3 1.3 0 0 1-.41.28a1.39 1.39 0 0 1-.5.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxCross(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M2.75 12H6a2 2 0 0 1 2 2a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2a2 2 0 0 1 2-2h3.25"/><path d="M21.25 10.375v4.875a6 6 0 0 1-6 6h-6.5a6 6 0 0 1-6-6v-6.5a6 6 0 0 1 6-6h4.875"/><path stroke-miterlimit="10" d="m20.75 2.75l-3.5 3.492m0-3.484l3.5 3.492"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxCrossFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 10.37v4.88A6.76 6.76 0 0 1 15.25 22h-6.5A6.76 6.76 0 0 1 2 15.25v-6.5A6.76 6.76 0 0 1 8.75 2h4.87a.75.75 0 1 1 0 1.5H8.75A5.25 5.25 0 0 0 3.5 8.75v2.5H6a2.718 2.718 0 0 1 1.94.8c.512.52.803 1.22.81 1.95A1.25 1.25 0 0 0 10 15.25h4a1.219 1.219 0 0 0 .88-.37a1.26 1.26 0 0 0 .37-.88A2.728 2.728 0 0 1 18 11.25h2.5v-.88a.75.75 0 1 1 1.5 0"/><path fill="currentColor" d="M21.28 5.72a.75.75 0 0 1-1.06 1.06l-1.23-1.22l-1.16 1.21a.74.74 0 0 1-.53.22a.73.73 0 0 1-.53-.22a.75.75 0 0 1 0-1.06l1.21-1.21l-1.21-1.21a.75.75 0 0 1 0-1.06a.739.739 0 0 1 1.06 0l1.21 1.21l1.23-1.22a.75.75 0 0 1 1.06 1.06L20.11 4.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxDownload(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M2.75 12H6a2 2 0 0 1 2 2a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2a2 2 0 0 1 2-2h3.25"/><path d="M21.25 10.375v4.875a6 6 0 0 1-6 6h-6.5a6 6 0 0 1-6-6v-6.5a6 6 0 0 1 6-6h4.875"/><path stroke-miterlimit="10" d="M18.957 7.75v-5"/><path stroke-linejoin="round" d="m16.664 5.645l1.967 1.967a.459.459 0 0 0 .652 0l1.967-1.967"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxDownloadFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.918 10.37v4.88a6.775 6.775 0 0 1-1.972 4.77a6.718 6.718 0 0 1-4.75 1.98H8.721a6.718 6.718 0 0 1-4.75-1.98A6.774 6.774 0 0 1 2 15.25v-6.5a6.774 6.774 0 0 1 1.972-4.77A6.718 6.718 0 0 1 8.722 2h4.85a.746.746 0 0 1 .747.75a.751.751 0 0 1-.747.75h-4.85a5.218 5.218 0 0 0-3.697 1.538A5.261 5.261 0 0 0 3.494 8.75v2.5h2.49a2.697 2.697 0 0 1 1.932.8c.51.52.799 1.22.806 1.95a1.25 1.25 0 0 0 1.245 1.25h3.984a1.21 1.21 0 0 0 .876-.37c.233-.234.365-.55.368-.88a2.74 2.74 0 0 1 .797-1.95a2.719 2.719 0 0 1 1.942-.8h2.49v-.88a.752.752 0 0 1 .747-.75a.745.745 0 0 1 .747.75"/><path fill="currentColor" d="M18.88 8.5a.745.745 0 0 1-.737-.75v-5a.742.742 0 0 1 1.026-.696a.746.746 0 0 1 .468.696v5a.761.761 0 0 1-.757.75"/><path fill="currentColor" d="M21.748 5.11a.745.745 0 0 0-1.055 0l-.996 1l-.747.75l-1.763-1.77a.745.745 0 0 0-1.256.535a.752.752 0 0 0 .2.525l1.992 2c.12.112.258.204.409.27a1.161 1.161 0 0 0 .916-.01c.146-.06.278-.148.388-.26l1.992-2a.752.752 0 0 0-.08-1.04"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 13v2.25A6.76 6.76 0 0 1 15.25 22h-6.5A6.76 6.76 0 0 1 2 15.25V13h4a1 1 0 0 1 .71.29c.183.192.286.445.29.71a3 3 0 0 0 3 3h4a3 3 0 0 0 3-3a1 1 0 0 1 1-1z"/><path fill="currentColor" d="M22 8.75V11h-4a3 3 0 0 0-3 3a1 1 0 0 1-1 1h-4a1 1 0 0 1-.71-.29A1 1 0 0 1 9 14a3 3 0 0 0-3-3H2V8.75A6.76 6.76 0 0 1 8.75 2h6.5A6.76 6.76 0 0 1 22 8.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxMinus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M2.75 12H6a2 2 0 0 1 2 2a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2a2 2 0 0 1 2-2h3.25"/><path d="M21.25 10.375v4.875a6 6 0 0 1-6 6h-6.5a6 6 0 0 1-6-6v-6.5a6 6 0 0 1 6-6h4.875"/><path stroke-miterlimit="10" d="M17.25 5.286h4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxMinusFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 10.37v4.88A6.76 6.76 0 0 1 15.25 22h-6.5A6.76 6.76 0 0 1 2 15.25v-6.5A6.76 6.76 0 0 1 8.75 2h4.87a.75.75 0 1 1 0 1.5H8.75A5.25 5.25 0 0 0 3.5 8.75v2.5H6a2.718 2.718 0 0 1 1.94.8c.512.52.803 1.22.81 1.95A1.25 1.25 0 0 0 10 15.25h4a1.219 1.219 0 0 0 .88-.37a1.26 1.26 0 0 0 .37-.88A2.73 2.73 0 0 1 18 11.25h2.5v-.88a.75.75 0 1 1 1.5 0"/><path fill="currentColor" d="M21.25 6.03h-4a.75.75 0 1 1 0-1.5h4a.75.75 0 1 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxNotification(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M2.75 12H6a2 2 0 0 1 2 2a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2a2 2 0 0 1 2-2h3.25"/><path stroke-linecap="round" d="M21.25 10.375v4.875a6 6 0 0 1-6 6h-6.5a6 6 0 0 1-6-6v-6.5a6 6 0 0 1 6-6h4.875"/><circle cx="19" cy="5" r="2.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxNotificationFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g fill="currentColor" clip-path="url(#mageInboxNotificationFill0)"><path d="M21.875 10.495v4.88a6.76 6.76 0 0 1-6.75 6.75h-6.5a6.76 6.76 0 0 1-6.75-6.75v-6.5a6.76 6.76 0 0 1 6.75-6.75h4.87a.75.75 0 1 1 0 1.5h-4.87a5.25 5.25 0 0 0-5.25 5.25v2.5h2.5a2.75 2.75 0 0 1 2.75 2.75a1.25 1.25 0 0 0 1.25 1.25h4a1.26 1.26 0 0 0 1.25-1.25a2.73 2.73 0 0 1 2.75-2.75h2.5v-.88a.75.75 0 1 1 1.5 0"/><path d="M18.875 8.375a3.25 3.25 0 1 0 0-6.5a3.25 3.25 0 0 0 0 6.5"/></g><defs><clipPath id="mageInboxNotificationFill0"><path fill="#fff" d="M1.875 1.875h20.25v20.25H1.875z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxPlus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M2.75 12H6a2 2 0 0 1 2 2a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2a2 2 0 0 1 2-2h3.25"/><path d="M21.25 10.375v4.875a6 6 0 0 1-6 6h-6.5a6 6 0 0 1-6-6v-6.5a6 6 0 0 1 6-6h4.875"/><path stroke-miterlimit="10" d="M18.745 2.75v5M16.25 5.255h5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxPlusFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 10.37v4.88A6.76 6.76 0 0 1 15.25 22h-6.5A6.76 6.76 0 0 1 2 15.25v-6.5A6.76 6.76 0 0 1 8.75 2h4.87a.75.75 0 1 1 0 1.5H8.75A5.25 5.25 0 0 0 3.5 8.75v2.5H6a2.718 2.718 0 0 1 1.94.8c.512.52.803 1.22.81 1.95A1.25 1.25 0 0 0 10 15.25h4a1.219 1.219 0 0 0 .88-.37a1.26 1.26 0 0 0 .37-.88A2.73 2.73 0 0 1 18 11.25h2.5v-.88a.75.75 0 1 1 1.5 0"/><path fill="currentColor" d="M22 5.25a.75.75 0 0 1-.75.75h-1.76v1.75a.75.75 0 1 1-1.5 0V6h-1.74a.75.75 0 1 1 0-1.5h1.74V2.75a.75.75 0 1 1 1.5 0V4.5h1.76a.76.76 0 0 1 .75.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxQuestionMark(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M2.75 12H6a2 2 0 0 1 2 2a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2a2 2 0 0 1 2-2h3.25"/><path d="M21.25 10.375v4.875a6 6 0 0 1-6 6h-6.5a6 6 0 0 1-6-6v-6.5a6 6 0 0 1 6-6h4.875"/><path stroke-miterlimit="10" d="M17.614 3.527a1.332 1.332 0 0 1 1.451-.755a1.292 1.292 0 0 1 .948.675a1.12 1.12 0 0 1-.653 1.543a.828.828 0 0 0-.534.748v.308"/><path stroke-linejoin="round" d="M18.802 7.75h.003"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxQuestionMarkFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 10.412v4.862a6.724 6.724 0 0 1-1.98 4.753A6.772 6.772 0 0 1 15.25 22h-6.5a6.772 6.772 0 0 1-4.77-1.973A6.723 6.723 0 0 1 2 15.274V8.797a6.724 6.724 0 0 1 1.98-4.753a6.772 6.772 0 0 1 4.77-1.972h4.87a.751.751 0 0 1 .75.747a.746.746 0 0 1-.75.747H8.75a5.26 5.26 0 0 0-3.712 1.532a5.222 5.222 0 0 0-1.538 3.7v2.49H6a2.728 2.728 0 0 1 1.94.798c.512.518.803 1.215.81 1.943A1.251 1.251 0 0 0 10 15.274h4a1.224 1.224 0 0 0 .88-.368c.234-.233.367-.548.37-.877a2.71 2.71 0 0 1 .8-1.944A2.73 2.73 0 0 1 18 11.29h2.5v-.877a.746.746 0 0 1 .75-.748a.751.751 0 0 1 .75.748"/><path fill="currentColor" d="M18.83 6.844a.751.751 0 0 1-.75-.747v-.299a1.6 1.6 0 0 1 1-1.445a.57.57 0 0 0 .16-.09a.379.379 0 0 0 .09-.129c.01-.05.01-.1 0-.15a.259.259 0 0 0 0-.149a.44.44 0 0 0-.15-.16a.521.521 0 0 0-.23-.099a.602.602 0 0 0-.4 0a.638.638 0 0 0-.26.28a.746.746 0 0 1-.716.459a.752.752 0 0 1-.686-.503a.745.745 0 0 1 .042-.585c.19-.426.517-.776.93-.996a2.157 2.157 0 0 1 1.34-.2c.308.053.599.175.85.36a1.844 1.844 0 0 1 .22 2.899a1.902 1.902 0 0 1-.67.399v.338a.734.734 0 0 1-.458.764a.742.742 0 0 1-.312.053m0 1.704a.751.751 0 0 1-.75-.747a.746.746 0 0 1 .75-.747a.751.751 0 0 1 .75.747a.746.746 0 0 1-.75.747"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxStar(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M2.75 12H6a2 2 0 0 1 2 2a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2a2 2 0 0 1 2-2h3.25"/><path d="M21.25 10.375v4.875a6 6 0 0 1-6 6h-6.5a6 6 0 0 1-6-6v-6.5a6 6 0 0 1 6-6h4.875"/><path stroke-linejoin="round" d="m19.107 1.864l-.31.924a2.358 2.358 0 0 1-1.505 1.505l-.934.31a.157.157 0 0 0 0 .3l.934.31a2.356 2.356 0 0 1 1.493 1.493l.31.936a.157.157 0 0 0 .298 0l.322-.924a2.358 2.358 0 0 1 1.492-1.493l.935-.31a.157.157 0 0 0 0-.3l-.923-.322a2.356 2.356 0 0 1-1.504-1.505l-.31-.935a.157.157 0 0 0-.298.011"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxStarFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.5 10.86v4.88a6.76 6.76 0 0 1-6.75 6.75h-6.5a6.76 6.76 0 0 1-6.75-6.75v-6.5a6.76 6.76 0 0 1 6.75-6.75h4.87a.75.75 0 1 1 0 1.5H8.25A5.25 5.25 0 0 0 3 9.24v2.5h2.5a2.75 2.75 0 0 1 2.75 2.75a1.25 1.25 0 0 0 1.25 1.25h4a1.26 1.26 0 0 0 1.25-1.25a2.73 2.73 0 0 1 2.75-2.75H20v-.88a.75.75 0 0 1 1.5 0"/><path fill="currentColor" d="M22.5 5.25c-.01.188-.07.37-.17.53a1 1 0 0 1-.46.34l-.93.31a1.53 1.53 0 0 0-1 1l-.33.94c-.07.17-.18.32-.32.44a.94.94 0 0 1-.54.17a1.11 1.11 0 0 1-.53-.17a.94.94 0 0 1-.33-.46l-.31-.93a1.66 1.66 0 0 0-.39-.63a1.64 1.64 0 0 0-.63-.39l-.93-.31a1.06 1.06 0 0 1-.45-.33a.9.9 0 0 1-.17-.53c.01-.193.071-.38.18-.54a.86.86 0 0 1 .44-.32l.93-.31a1.59 1.59 0 0 0 1-1l.31-.93a1.45 1.45 0 0 1 .3-.43a1.25 1.25 0 0 1 .53-.19a1 1 0 0 1 .54.15a1 1 0 0 1 .35.45l.32 1a1.6 1.6 0 0 0 1 1l.94.33c.17.07.32.18.44.32c.116.14.19.31.21.49"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxUpload(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M2.75 12H6a2 2 0 0 1 2 2a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2a2 2 0 0 1 2-2h3.25"/><path d="M21.25 10.375v4.875a6 6 0 0 1-6 6h-6.5a6 6 0 0 1-6-6v-6.5a6 6 0 0 1 6-6h4.875"/><path stroke-miterlimit="10" d="M18.957 2.75v5"/><path stroke-linejoin="round" d="m21.25 4.855l-1.967-1.967a.458.458 0 0 0-.652 0l-1.967 1.967"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxUploadFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.97 10.393v4.87A6.745 6.745 0 0 1 15.23 22H8.74A6.751 6.751 0 0 1 2 15.264V8.777A6.745 6.745 0 0 1 8.74 2.04h4.862a.749.749 0 1 1 0 1.497H8.74a5.243 5.243 0 0 0-5.242 5.24v2.495h2.496a2.717 2.717 0 0 1 1.937.798a2.82 2.82 0 0 1 .809 1.946a1.247 1.247 0 0 0 1.248 1.248h3.994a1.22 1.22 0 0 0 .878-.37c.234-.233.367-.548.37-.878a2.722 2.722 0 0 1 2.746-2.745h2.496v-.878a.749.749 0 0 1 1.497 0"/><path fill="currentColor" d="M18.924 8.527a.749.749 0 0 1-.739-.748v-4.99a.74.74 0 0 1 .74-.749a.75.75 0 0 1 .758.749v4.99a.758.758 0 0 1-.759.748"/><path fill="currentColor" d="M21.8 4.355L19.803 2.36a1.049 1.049 0 0 0-.39-.27a1.219 1.219 0 0 0-.459-.09c-.15.002-.3.032-.44.09a1.358 1.358 0 0 0-.389.27l-1.997 1.995a.748.748 0 0 0 0 1.058a.74.74 0 0 0 1.059 0l.998-.998l.75-.748l1.757 1.756a.728.728 0 0 0 .529.22a.74.74 0 0 0 .529-.22a.748.748 0 0 0 .05-1.068"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InformationCircle(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-width="1.5" d="M12 21.5a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19m.005-4.222v-6.333"/><path stroke-width="2" d="M11.956 7.443h.01"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InformationCircleFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.75A10.25 10.25 0 1 0 22.25 12A10.26 10.26 0 0 0 12 1.75m0 4.69a1 1 0 1 1-.03 0zm1 10.83a1 1 0 1 1-2 0v-6.33a1 1 0 0 1 2 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InformationSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 17.139v-6.167"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.958 7.563h.008"/><rect width="18.5" height="18.5" x="2.75" y="2.75" stroke-width="1.5" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InformationSquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 2h-6.5A6.76 6.76 0 0 0 2 8.75v6.5A6.76 6.76 0 0 0 8.75 22h6.5A6.76 6.76 0 0 0 22 15.25v-6.5A6.76 6.76 0 0 0 15.25 2m-3.29 4.56a1 1 0 1 1-.376.073a1 1 0 0 1 .386-.073zm1 10.58a1 1 0 1 1-2 0v-6.17a1 1 0 1 1 2 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InstagramCircle(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.61 12.243a1.6 1.6 0 1 1-1.56-1.63a1.62 1.62 0 0 1 1.56 1.63"/><path fill="currentColor" d="M14.763 7.233H9.338a2.024 2.024 0 0 0-2.024 2.024v5.547a2.024 2.024 0 0 0 2.024 2.024h5.425a2.024 2.024 0 0 0 2.024-2.024V9.267a2.026 2.026 0 0 0-2.024-2.034m-2.713 7.723a2.703 2.703 0 1 1 2.642-2.703a2.672 2.672 0 0 1-2.642 2.703m2.936-5.405a.496.496 0 0 1-.496-.506a.506.506 0 1 1 1.012 0a.496.496 0 0 1-.557.506z"/><path fill="currentColor" d="M12.05 2a10 10 0 1 0-.1 20a10 10 0 0 0 .1-20m6.073 12.702a3.39 3.39 0 0 1-3.41 3.411H9.389a3.392 3.392 0 0 1-3.411-3.41V9.378a3.39 3.39 0 0 1 3.41-3.411h5.325a3.39 3.39 0 0 1 3.41 3.41z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InstagramSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.823 12.234c-.016.35-.13.688-.331.975a1.697 1.697 0 0 1-.829.643a1.765 1.765 0 0 1-1.053.088a1.803 1.803 0 0 1-.926-.516a1.892 1.892 0 0 1-.468-.976a1.755 1.755 0 0 1 .127-1.043c.144-.327.38-.606.682-.8c.307-.19.662-.291 1.024-.292c.477.026.926.232 1.258.575a1.853 1.853 0 0 1 .516 1.346"/><path fill="currentColor" d="M17.265 8.002a2.263 2.263 0 0 0-1.248-1.248a2.564 2.564 0 0 0-.887-.175H8.968A2.31 2.31 0 0 0 6.667 8.88v6.279a2.292 2.292 0 0 0 .682 1.628a2.32 2.32 0 0 0 1.619.673h6.162a2.32 2.32 0 0 0 2.123-1.419a2.3 2.3 0 0 0 .178-.882v-6.27a2.553 2.553 0 0 0-.166-.887m-2.437 5.441a2.926 2.926 0 0 1-.644.975c-.28.283-.611.51-.975.673a3.129 3.129 0 0 1-2.486-.028a3.08 3.08 0 0 1-1.765-3.365a3.22 3.22 0 0 1 .829-1.59a3.11 3.11 0 0 1 3.354-.692c.567.23 1.05.628 1.384 1.141a3.03 3.03 0 0 1 .527 1.677c.014.415-.063.827-.224 1.209M15.9 8.626a.555.555 0 1 1-1.102 0a.557.557 0 1 1 1.102 0"/><path fill="currentColor" d="M16.875 2.25h-9.75A4.875 4.875 0 0 0 2.25 7.125v9.75a4.875 4.875 0 0 0 4.875 4.875h9.75a4.875 4.875 0 0 0 4.875-4.875v-9.75a4.875 4.875 0 0 0-4.875-4.875m2.067 12.812c.01.51-.087 1.019-.283 1.491a3.9 3.9 0 0 1-2.096 2.096c-.473.196-.98.292-1.492.283H9.075a3.754 3.754 0 0 1-1.492-.282a4.007 4.007 0 0 1-1.258-.839a3.9 3.9 0 0 1-.838-1.258a3.725 3.725 0 0 1-.312-1.492V9.018a3.754 3.754 0 0 1 .283-1.492A3.9 3.9 0 0 1 7.535 5.41a3.9 3.9 0 0 1 1.54-.263h6.045a3.802 3.802 0 0 1 2.73 1.121c.357.362.641.79.838 1.258c.195.473.292.98.283 1.492z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Key(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19.543 4.471A6.501 6.501 0 0 0 8.618 7.255a6.52 6.52 0 0 0 .34 4.447L2.5 18.179V21.5h3.364l-.318-1.99l1.965.33l1.67-1.661l-.318-1.99l1.976.318l1.47-1.472a6.497 6.497 0 0 0 8.06-2.261a6.52 6.52 0 0 0-.838-8.338z"/><path d="M16.99 10.23a2.258 2.258 0 0 1-3.476-2.853a2.26 2.26 0 0 1 3.477-.339a2.275 2.275 0 0 1 0 3.192"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.953 8.463a7.107 7.107 0 0 0-2.106-4.372l-.108-.088A7.078 7.078 0 0 0 8.086 6.854a7.09 7.09 0 0 0 .078 4.684l-5.95 5.973a.732.732 0 0 0-.214.517v3.24a.723.723 0 0 0 .731.732H5.98a.72.72 0 0 0 .722-.84l-.147-.975l.907.146a.75.75 0 0 0 .644-.195l1.629-1.63a.722.722 0 0 0 .204-.634l-.156-.976l.976.146a.731.731 0 0 0 .634-.205l1.082-1.093a7.064 7.064 0 0 0 6.787-1.042a7.076 7.076 0 0 0 2.702-6.317zm-4.525 2.372a2.925 2.925 0 0 1-1.54.83a3.05 3.05 0 0 1-.586.058a3.1 3.1 0 0 1-1.15-.224a2.995 2.995 0 0 1-1.346-1.113a2.93 2.93 0 0 1-.508-1.669a2.998 2.998 0 0 1 1.853-2.772a3.03 3.03 0 0 1 3.277.654a3.017 3.017 0 0 1 0 4.236"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Keyboard(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.981 6.583H4.02c-.701 0-1.269.647-1.269 1.444v7.946c0 .797.568 1.444 1.269 1.444h15.96c.701 0 1.269-.647 1.269-1.444V8.027c0-.797-.568-1.444-1.269-1.444M9.357 13.806h5.286m2.748 0h1.058m-11.84 0H5.55m.001-3.612H6.61m2.885 0h1.057m2.896 0h1.057m2.886 0h1.058"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyboardFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 5.81H4A2.11 2.11 0 0 0 2 8v8a2.11 2.11 0 0 0 2 2.19h16A2.11 2.11 0 0 0 22 16V8a2.11 2.11 0 0 0-2-2.19m-6.53 3.36h1.05a1 1 0 0 1 0 2h-1.05a1 1 0 0 1 0-2m-4 0h1.06a1 1 0 1 1 0 2H9.47a1 1 0 0 1 0-2m-2.88 5.61H5.53a1 1 0 0 1 0-2h1.06a1 1 0 1 1 0 2m0-3.61H5.53a1 1 0 0 1 0-2h1.06a1 1 0 1 1 0 2m8 3.61H9.3a1 1 0 0 1 0-2h5.29a1 1 0 1 1 0 2m3.81 0h-1.06a1 1 0 0 1 0-2h1.06a1 1 0 1 1 0 2m0-3.61h-1.06a1 1 0 0 1 0-2h1.06a1 1 0 1 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LarrowDownLeft(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m9.32 20.5l-5.21-5.21a1.214 1.214 0 0 1 0-1.724l5.21-5.209"/><path d="M20.249 3.5v7.286a3.643 3.643 0 0 1-3.643 3.643H3.759"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LarrowDownRight(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m14.68 20.5l5.21-5.21a1.212 1.212 0 0 0 0-1.724l-5.21-5.209"/><path d="M3.751 3.5v7.286a3.643 3.643 0 0 0 3.643 3.643h12.847"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LarrowLeftDown(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m3.5 14.68l5.21 5.21a1.212 1.212 0 0 0 1.724 0l5.209-5.21"/><path d="M20.5 3.751h-7.286a3.643 3.643 0 0 0-3.643 3.643v12.847"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LarrowLeftUp(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m3.5 9.32l5.21-5.21a1.214 1.214 0 0 1 1.724 0l5.209 5.21"/><path d="M20.5 20.249h-7.286a3.643 3.643 0 0 1-3.643-3.643V3.759"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LarrowRightDown(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m20.5 14.68l-5.21 5.21a1.213 1.213 0 0 1-1.724 0l-5.209-5.21"/><path d="M3.5 3.751h7.286a3.643 3.643 0 0 1 3.643 3.643v12.847"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LarrowRightUp(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m20.5 9.32l-5.21-5.21a1.215 1.215 0 0 0-1.724 0L8.357 9.32"/><path d="M3.5 20.249h7.286a3.643 3.643 0 0 0 3.643-3.643V3.759"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LarrowUpLeft(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.32 3.5L4.11 8.71a1.214 1.214 0 0 0 0 1.724l5.21 5.209"/><path d="M20.249 20.5v-7.286a3.643 3.643 0 0 0-3.643-3.643H3.759"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LarrowUpRight(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m14.68 3.5l5.21 5.21a1.212 1.212 0 0 1 0 1.724l-5.21 5.209"/><path d="M3.751 20.5v-7.286a3.643 3.643 0 0 1 3.643-3.643h12.847"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Laptop(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.75 16.163V6.212a1.106 1.106 0 0 0-1.105-1.106H5.377a1.106 1.106 0 0 0-1.105 1.106v9.95m-1.772.001h19v1.365c0 .363-.167.71-.464.966c-.297.256-.7.4-1.12.4H4.084a1.72 1.72 0 0 1-1.12-.4a1.278 1.278 0 0 1-.463-.966z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LaptopFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.47 15.39h-1V6.22a1.87 1.87 0 0 0-.54-1.32a1.85 1.85 0 0 0-1.31-.54H5.35a1.91 1.91 0 0 0-.71.14a2 2 0 0 0-1 1a1.87 1.87 0 0 0-.14.71v9.2h-1a.75.75 0 0 0-.75.75v1.36a2 2 0 0 0 .73 1.54a2.45 2.45 0 0 0 1.6.58h15.84a2.5 2.5 0 0 0 1.61-.58a2 2 0 0 0 .72-1.54v-1.36a.749.749 0 0 0-.78-.77m-.85 2.11a.93.93 0 0 1-.2.58a.802.802 0 0 1-.62.31H4.14a.8.8 0 0 1-.62-.31a.93.93 0 0 1-.2-.58v-.88h17.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutCenter(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M2.75 8.917h18.5M12 21.25V8.917"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutCenterFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 9.92V22H8.75A6.76 6.76 0 0 1 2 15.25V9.92zm11 0v5.33A6.76 6.76 0 0 1 15.25 22H13V9.92zm-.06-2H2.06A6.75 6.75 0 0 1 8.75 2h6.5a6.75 6.75 0 0 1 6.69 5.92"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutDown(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M21.25 15.084H2.75"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutDownFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.94 16.08A6.75 6.75 0 0 1 15.25 22h-6.5a6.75 6.75 0 0 1-6.69-5.92zM22 8.75v5.33H2V8.75A6.76 6.76 0 0 1 8.75 2h6.5A6.76 6.76 0 0 1 22 8.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutGrid(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M12 2.75v18.5M21.25 12H2.75"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutGridFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11 13v9H8.75A6.76 6.76 0 0 1 2 15.25V13zm11 0v2.25A6.76 6.76 0 0 1 15.25 22H13v-9zm0-4.25V11h-9V2h2.25A6.76 6.76 0 0 1 22 8.75M11 2v9H2V8.75A6.76 6.76 0 0 1 8.75 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutLeft(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M8.917 21.25V2.75"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutLeftFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.92 2.06v19.88A6.75 6.75 0 0 1 2 15.25v-6.5a6.75 6.75 0 0 1 5.92-6.69M22 8.75v6.5A6.76 6.76 0 0 1 15.25 22H9.92V2h5.33A6.76 6.76 0 0 1 22 8.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutRight(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M15.084 2.75v18.5"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutRightFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 8.75v6.5a6.75 6.75 0 0 1-5.92 6.69V2.06A6.75 6.75 0 0 1 22 8.75M14.08 2v20H8.75A6.76 6.76 0 0 1 2 15.25v-6.5A6.76 6.76 0 0 1 8.75 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutUp(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M2.75 8.917h18.5"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutUpFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.94 7.92H2.06A6.75 6.75 0 0 1 8.75 2h6.5a6.75 6.75 0 0 1 6.69 5.92m.06 2v5.33A6.76 6.76 0 0 1 15.25 22h-6.5A6.76 6.76 0 0 1 2 15.25V9.92z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutUpLeft(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M2.75 9.167h18.5"/><path stroke-linecap="square" stroke-linejoin="round" d="M9.167 21.25V9.167"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutUpLeftFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.97 8.17H2.03A6.76 6.76 0 0 1 8.75 2h6.5a6.76 6.76 0 0 1 6.72 6.17m-13.8 2v11.8A6.76 6.76 0 0 1 2 15.25v-5.08zm13.83 0v5.08A6.76 6.76 0 0 1 15.25 22h-5.08V10.17z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutUpRight(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M2.75 9.167h18.5M15.334 21.25V9.167"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutUpRightFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.97 8.17H2.03A6.76 6.76 0 0 1 8.75 2h6.5a6.76 6.76 0 0 1 6.72 6.17m-7.64 2V22H8.75A6.76 6.76 0 0 1 2 15.25v-5.08zm7.67 0v5.08a6.74 6.74 0 0 1-5.67 6.65V10.17z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lens(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 21.5a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19"/><path d="M16.113 14.375v-4.75L12 7.25L7.886 9.625v4.75L12 16.75zm1.255-10.213L12.01 7.25m9.462 5.425l-5.359-3.05m0 10.935v-6.185m-9.49 5.453l5.368-3.078m-9.462-5.443l5.358 3.068M7.886 3.44v6.185"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LensFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.272 17.438l-4.37 2.504a10.064 10.064 0 0 1-2.728-3.207A10.029 10.029 0 0 1 2 12.697l4.856 2.781h.06zm5.581-1.238v5.038a10.101 10.101 0 0 1-8.312-.277l4.787-2.732h.09zm6.087-2.938a10.012 10.012 0 0 1-1.419 3.998a10.05 10.05 0 0 1-2.95 3.057v-9.54zm.06-1.94L17.144 8.57l-.209-.09l-3.207-1.87l4.37-2.503a10.004 10.004 0 0 1 2.727 3.19A9.97 9.97 0 0 1 22 11.322m-5.571-8.273l-4.786 2.77h-.08L8.147 7.8V2.762a10.13 10.13 0 0 1 8.282.257zm-10 .663v9.54l-4.37-2.494a10.022 10.022 0 0 1 1.423-3.993a10.06 10.06 0 0 1 2.947-3.053m9.424 6.096v4.423L12 16.448l-3.853-2.217V9.808L12 7.59z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightBulb(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17.252 12.49c-.284 2.365-1.833 3.31-2.502 3.996c-.67.688-.55.825-.505 1.834a.916.916 0 0 1-.916.971h-2.658a.918.918 0 0 1-.917-.971c0-.99.092-1.22-.504-1.834c-.76-.76-2.548-1.833-2.548-4.784a5.307 5.307 0 1 1 10.55.788"/><path d="M10.46 19.236v1.512c0 .413.23.752.513.752h2.053c.285 0 .514-.34.514-.752v-1.512m-2.32-10.54a2.227 2.227 0 0 0-2.226 2.227m10.338.981h1.834m-3.68-6.012l1.301-1.301M18.486 17l1.301 1.3M12 2.377V3.86m-6.76.73l1.292 1.302M4.24 18.3L5.532 17m-.864-5.096H2.835"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightBulbElectricity(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18.156 16.314c-2.583 2.085-5.177 1.491-6.469 1.491c-1.292 0-1.373.262-2.321 1.27a1.29 1.29 0 0 1-.424.302a1.223 1.223 0 0 1-1.01 0a1.293 1.293 0 0 1-.433-.303l-2.594-2.568a1.26 1.26 0 0 1-.303-.454a1.297 1.297 0 0 1-.101-.514a1.228 1.228 0 0 1 .404-.937c1.009-.947 1.292-1.128 1.292-2.327c0-1.51-.717-4.372 2.24-7.334a7.43 7.43 0 0 1 2.523-1.662a7.542 7.542 0 0 1 3.028-.524a7.47 7.47 0 0 1 5.45 2.61a7.52 7.52 0 0 1 .8 8.662a7.529 7.529 0 0 1-2.051 2.288zm-13.554-.01l-1.514 1.511a1.007 1.007 0 0 0-.202 1.26l2.11 2.014c.282.292.847.172 1.291-.242l1.423-1.52"/><path d="m15.248 5.873l-5.063 4.273a.33.33 0 0 0-.094.12a.414.414 0 0 0-.036.155a.33.33 0 0 0 .067.136a.3.3 0 0 0 .134.031l2.913.68l-1.26 3.382a.138.138 0 0 0-.014.063l.052.012l.062.014l5.063-4.273a.224.224 0 0 0 .212-.114a.276.276 0 0 0 .037-.155a.33.33 0 0 0-.067-.136a.3.3 0 0 0-.135-.031l-2.902-.677l1.172-3.338a.074.074 0 0 0 .015-.062l-.063-.014z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightBulbElectricityFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.666 7.817a8.314 8.314 0 0 0-4.32-5.011a8.304 8.304 0 0 0-6.63-.22a8.059 8.059 0 0 0-2.78 1.833a8.371 8.371 0 0 0-2.49 7.177v.702c0 .801 0 .841-.87 1.613l-.17.16a2.004 2.004 0 0 0-.65 1.494c0 .169.02.337.06.501l-1.22 1.223a1.775 1.775 0 0 0-.57 1.002c-.068.4.007.812.21 1.163a.54.54 0 0 0 .14.16l2.08 2.005a1.31 1.31 0 0 0 .94.38a2.158 2.158 0 0 0 1.44-.63l1.11-1.193c.433.11.89.07 1.3-.11c.257-.107.487-.271.67-.482l.29-.31c.65-.702.66-.722 1.48-.722h.67a8.347 8.347 0 0 0 6.15-1.594a.573.573 0 0 0 .13-.08a8.289 8.289 0 0 0 2.945-4.072a8.307 8.307 0 0 0 .055-5.029zM5.806 20.306a.8.8 0 0 1-.32.2l-1.95-1.864a.331.331 0 0 1 0-.1a.32.32 0 0 1 .12-.19l1-1.003l2 2.005zm11.62-9.632a1.001 1.001 0 0 1-.34.32l-.12.05l-4.29 3.63a.67.67 0 0 1-.42.16h-.15a.79.79 0 0 1-.6-.662c0-.097.02-.192.06-.28l.84-2.266l-1.77-.41a.798.798 0 0 1-.3-.091a.56.56 0 0 1-.21-.16a.833.833 0 0 1-.2-.411a.884.884 0 0 1 0-.16a.975.975 0 0 1 .1-.382a.74.74 0 0 1 .26-.34l4.4-3.729a.659.659 0 0 1 .8 0a.77.77 0 0 1 .38.471c.014.106.014.214 0 .32a.461.461 0 0 1 0 .091l-.76 2.205l1.77.411a1 1 0 0 1 .3.09c.09.037.166.1.22.18a.844.844 0 0 1 .19.392a.944.944 0 0 1-.16.571"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightBulbFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.754 2.305a.75.75 0 0 0-1.5 0v1.48a.75.75 0 0 0 1.5 0zm5.111 7.99a5.879 5.879 0 0 0-1.11-2.22a6 6 0 0 0-1.91-1.59a6.17 6.17 0 0 0-2.38-.69a6 6 0 0 0-2.46.33a6 6 0 0 0-2.13 1.29a6.18 6.18 0 0 0-1.43 2a6 6 0 0 0-.49 2.43a6.09 6.09 0 0 0 2.41 5l.35.33c.3.31.3.31.29 1v.32a1.58 1.58 0 0 0 .1.65c.07.222.194.425.36.59c.076.078.16.148.25.21v1a1.38 1.38 0 0 0 1.26 1.5h2a1.39 1.39 0 0 0 1.27-1.5v-1a1.64 1.64 0 0 0 .25-.21c.157-.166.277-.364.35-.58a1.66 1.66 0 0 0 .1-.66v-.37c0-.55 0-.55.31-.9l.38-.35a6.17 6.17 0 0 0 2.33-4.07a5.882 5.882 0 0 0-.1-2.51m-5.07 10.63h-1.58v-.63h1.58zm-.79-10.56a1.23 1.23 0 0 0-1.23 1.23a1 1 0 1 1-2 0a3.21 3.21 0 0 1 3.23-3.23a1 1 0 0 1 0 2m9.16 2.5h-1.83a.75.75 0 0 1 0-1.5h1.83a.75.75 0 1 1 0 1.5m-3.68-6.01a.74.74 0 0 1-.53-.22a.75.75 0 0 1 0-1.06l1.3-1.3a.75.75 0 0 1 1.06 1.06l-1.3 1.3a.73.73 0 0 1-.53.22m2.3 12.44a.79.79 0 0 1-.53-.22l-1.3-1.3a.75.75 0 0 1 .242-1.226a.74.74 0 0 1 .818.166l1.3 1.3a.74.74 0 0 1 0 1.06a.75.75 0 0 1-.53.22M6.535 6.855a.75.75 0 0 1-.53-.22l-1.3-1.3a.753.753 0 1 1 1.07-1.06l1.29 1.3a.75.75 0 0 1-.53 1.28m-2.29 12.44a.71.71 0 0 1-.53-.22a.75.75 0 0 1 0-1.06l1.3-1.3a.75.75 0 0 1 1.06 1.06l-1.29 1.31a.79.79 0 0 1-.54.21m.42-6.43h-1.83a.75.75 0 1 1 0-1.5h1.83a.75.75 0 1 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightBulbOff(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18.131 16.26c-2.565 2.07-5.13 1.478-6.413 1.478c-1.283 0-1.36.27-2.296 1.287a1.284 1.284 0 0 1-1.436.296a1.284 1.284 0 0 1-.424-.296l-2.565-2.573a1.287 1.287 0 0 1 0-1.865c1.026-.94 1.282-1.12 1.282-2.303c0-1.505-.705-4.348 2.22-7.28a7.425 7.425 0 0 1 5.535-2.2A7.41 7.41 0 0 1 19.4 5.39a7.455 7.455 0 0 1 1.75 5.707a7.463 7.463 0 0 1-2.993 5.162zm-13.429-.014l-1.5 1.505c-.411.412-.527.978-.245 1.287l2.091 2.006c.282.283.847.168 1.283-.244l1.41-1.505"/><path d="M15.912 6.508a3.103 3.103 0 0 0-3.39-.678a3.112 3.112 0 0 0-1.01.678"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightBulbOffFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.657 7.843a8.167 8.167 0 0 0-4.293-5.023a8.167 8.167 0 0 0-6.585-.242a8.144 8.144 0 0 0-2.815 1.86a8.349 8.349 0 0 0-2.473 7.143v.703c0 .804 0 .834-.915 1.638l-.14.13a2.068 2.068 0 0 0-.643 1.487c-.002.156.019.312.06.462l-1.237 1.246a1.698 1.698 0 0 0-.242 2.371l2.091 2.01A1.307 1.307 0 0 0 5.41 22a2.132 2.132 0 0 0 1.428-.633l1.136-1.206c.161.04.327.06.493.06a2.011 2.011 0 0 0 1.488-.643l.301-.331c.634-.704.644-.714 1.448-.714h.673a8.318 8.318 0 0 0 6.143-1.597l.11-.07a8.191 8.191 0 0 0 3.308-5.707a8.033 8.033 0 0 0-.281-3.316M5.802 20.312a.733.733 0 0 1-.311.19l-2.01-1.878a.653.653 0 0 1 .18-.292l1.005-1.004l2.011 2.01zM16.66 7.25a1.006 1.006 0 0 1-1.417 0a2.222 2.222 0 0 0-.684-.462a2.173 2.173 0 0 0-1.629 0c-.254.11-.486.267-.683.462a1.006 1.006 0 0 1-1.428-1.416c.38-.389.835-.697 1.337-.905a4.214 4.214 0 0 1 3.167 0c.505.209.963.516 1.347.905a1.005 1.005 0 0 1-.01 1.376z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Line(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g fill="currentColor" clip-path="url(#mageLine0)"><path d="M9.33 12.22v.52c0 .09-.05.13-.14.13H7.12c-.12 0-.16 0-.16-.16V9.5c0-.13 0-.17.17-.17h.46c.12 0 .16.05.16.16v2.56h1.47a.13.13 0 0 1 .115.113a.13.13 0 0 1-.005.057m1.28-2.73v3.22c0 .12 0 .16-.17.16c-.16.01-.32.01-.48 0a.14.14 0 0 1-.16-.16V9.49a.15.15 0 0 1 .2-.17h.47a.15.15 0 0 1 .14.17m3.63-.01v3.23c0 .11 0 .15-.16.15h-.51a.14.14 0 0 1-.12-.06l-.75-1l-.7-.95v-.07v1.94c0 .14 0 .17-.17.17h-.46a.14.14 0 0 1-.155-.095a.14.14 0 0 1-.005-.065V9.5a.14.14 0 0 1 .15-.15h.55a.13.13 0 0 1 .09.06L13.1 11l.29.39v.05V9.5c0-.13.05-.18.18-.18h.45c.16 0 .22.05.22.16m2.91.02v.5c0 .11 0 .16-.15.16h-1.35c-.07 0-.07 0-.07.07v.42c0 .08 0 .08.07.08H17c.13 0 .17.05.17.18v.46a.14.14 0 0 1-.16.16h-1.36c-.07 0-.07 0-.07.08V12c0 .08 0 .08.07.08H17c.12 0 .16 0 .16.17v.46c0 .12 0 .16-.16.16h-2.05c-.18 0-.17 0-.17-.16V9.5c0-.12.05-.17.17-.17h2c.16 0 .2.05.2.17"/><path d="M18 2H6a4 4 0 0 0-4 4v12a4 4 0 0 0 4 4h12a4 4 0 0 0 4-4V6a4 4 0 0 0-4-4m1.1 11a6.358 6.358 0 0 1-1 1.58a13.928 13.928 0 0 1-2.18 2c-.92.723-1.882 1.39-2.88 2c-.31.19-.63.36-.95.52a.94.94 0 0 1-.46.13a.23.23 0 0 1-.27-.26a2.863 2.863 0 0 1 0-.41a4.93 4.93 0 0 0 .11-.94a.48.48 0 0 0-.28-.45a1.7 1.7 0 0 0-.53-.16a8.82 8.82 0 0 1-2.76-.88a6.87 6.87 0 0 1-2.13-1.69a5.47 5.47 0 0 1-1.21-2.46c0-.21-.06-.42-.08-.64c-.02-.22 0-.41 0-.62a5 5 0 0 1 .36-1.61A6.08 6.08 0 0 1 6 7.25a7.52 7.52 0 0 1 2.51-1.73A9.15 9.15 0 0 1 10.36 5l.76-.1h1.69a8.61 8.61 0 0 1 3.28.95c.71.365 1.356.844 1.91 1.42a5.49 5.49 0 0 1 1.27 2.14c.11.323.181.66.21 1v.65A4.72 4.72 0 0 1 19.1 13"/><path d="M14.24 9.48v3.23c0 .11 0 .15-.16.15h-.51a.14.14 0 0 1-.12-.06l-.75-1l-.7-.95v-.07v1.94c0 .14 0 .17-.17.17h-.46a.14.14 0 0 1-.155-.095a.14.14 0 0 1-.005-.065V9.5a.14.14 0 0 1 .15-.15h.55a.13.13 0 0 1 .09.06L13.1 11l.29.39v.05V9.5c0-.13.05-.18.18-.18h.45c.16 0 .22.05.22.16m1.34.72v.42c0 .08 0 .08.07.08H17c.13 0 .17.05.17.18v.46a.14.14 0 0 1-.16.16h-1.36c-.07 0-.07 0-.07.08V12c0 .08 0 .08.07.08H17c.12 0 .16 0 .16.17v.46c0 .12 0 .16-.16.16h-2.05c-.18 0-.17 0-.17-.16V9.5c0-.12.05-.17.17-.17h2c.12 0 .16.05.16.17v.5c0 .11 0 .16-.15.16h-1.31c-.07-.03-.07-.03-.07.04m-6.25 2.02v.52c0 .09-.05.13-.14.13H7.12c-.12 0-.16 0-.16-.16V9.5c0-.13 0-.17.17-.17h.46c.12 0 .16.05.16.16v2.56h1.47a.13.13 0 0 1 .115.113a.13.13 0 0 1-.005.057m1.28-2.73v3.22c0 .12 0 .16-.17.16c-.16.01-.32.01-.48 0a.14.14 0 0 1-.16-.16V9.49a.15.15 0 0 1 .2-.17h.47a.15.15 0 0 1 .14.17"/></g><defs><clipPath id="mageLine0"><path fill="#fff" d="M2 2h20v20H2z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Link(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.522 13.48a4.094 4.094 0 0 0 2.892 1.199a4.74 4.74 0 0 0 1.063-.136a4.176 4.176 0 0 0 1.828-1.063l.969-.968l2.878-2.888a4.085 4.085 0 0 0-2.922-6.873a4.095 4.095 0 0 0-2.862 1.096L11.49 6.736"/><path d="m12.445 17.336l-2.892 2.888a4.094 4.094 0 0 1-6.801-2.944a4.085 4.085 0 0 1 1.031-2.833l2.892-2.888l.969-.968A4.175 4.175 0 0 1 9.47 9.53a4.096 4.096 0 0 1 3.956 1.062"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Linkedin(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.303 2.25H6.697A4.447 4.447 0 0 0 2.25 6.697v10.606a4.447 4.447 0 0 0 4.447 4.447h10.606a4.447 4.447 0 0 0 4.447-4.447V6.697a4.447 4.447 0 0 0-4.447-4.447m-8.46 15.742a.4.4 0 0 1-.4.423h-1.78a.411.411 0 0 1-.4-.412V10.6a.4.4 0 0 1 .4-.411h1.78a.399.399 0 0 1 .4.411zM7.52 8.632a1.467 1.467 0 1 1 .022-2.935A1.467 1.467 0 0 1 7.52 8.63m10.817 9.35a.389.389 0 0 1-.378.388H16.08a.39.39 0 0 1-.378-.389v-3.424c0-.511.156-2.223-1.356-2.223c-1.179 0-1.412 1.2-1.457 1.734v3.991a.39.39 0 0 1-.378.39h-1.823a.389.389 0 0 1-.389-.39v-7.493a.389.389 0 0 1 .39-.378h1.822a.389.389 0 0 1 .39.378v.645a2.59 2.59 0 0 1 2.434-1.112c3.035 0 3.024 2.835 3.024 4.447z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Location(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10" stroke-width="1.5"><path d="M21.5 12h-2.111M12 2.5v2.111M2.5 12h2.111M12 21.5v-2.111m0 0A7.389 7.389 0 1 0 12 4.61a7.389 7.389 0 0 0 0 14.778Z"/><path d="M12 16.222a4.222 4.222 0 1 0 0-8.444a4.222 4.222 0 0 0 0 8.444Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.45 11.227h-1.39a8.18 8.18 0 0 0-7.36-7.36v-1.39a.75.75 0 0 0-1.5 0v1.39a8.17 8.17 0 0 0-7.31 7.36H2.5a.75.75 0 1 0 0 1.5h1.39a8.18 8.18 0 0 0 7.36 7.36v1.39a.75.75 0 0 0 1.5 0v-1.39a8.19 8.19 0 0 0 7.36-7.36h1.39a.75.75 0 1 0 0-1.5zm-9.5 7.39a6.64 6.64 0 1 1 6.64-6.64a6.65 6.65 0 0 1-6.64 6.65z"/><path fill="currentColor" d="M16.48 11.987a4.54 4.54 0 1 1-4.53-4.54a4.53 4.53 0 0 1 4.53 4.54"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LocationPin(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 13.632A5.441 5.441 0 1 0 12 2.75a5.441 5.441 0 0 0 0 10.882m0 0v7.618"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lock(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 9.688H7c-1.38 0-2.5 1.035-2.5 2.312v6.938c0 1.277 1.12 2.312 2.5 2.312h10c1.38 0 2.5-1.035 2.5-2.312V12c0-1.277-1.12-2.312-2.5-2.312m-9.625 0V7.374a4.625 4.625 0 0 1 9.25 0v2.313m-8.094 8.094h6.938"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.38 9.1V7.51a5.378 5.378 0 0 0-7.487-5.18A5.38 5.38 0 0 0 6.63 7.51V9.1a3.12 3.12 0 0 0-2.88 3v6.94A3.16 3.16 0 0 0 7 22.1h10a3.16 3.16 0 0 0 3.25-3.06V12.1a3.12 3.12 0 0 0-2.87-3M8.13 7.51a3.84 3.84 0 0 1 1.13-2.74A3.91 3.91 0 0 1 12 3.63a3.89 3.89 0 0 1 3.88 3.88v1.56H8.13zm7.34 11.4h-6.9a1 1 0 0 1 0-2h6.94a1 1 0 1 1 0 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M15.13 10.552H8.87c-.865 0-1.565.648-1.565 1.448v4.343c0 .8.7 1.448 1.565 1.448h6.26c.865 0 1.566-.648 1.566-1.448V12c0-.8-.7-1.448-1.565-1.448m-6.027 0V9.104a2.896 2.896 0 1 1 5.792 0v1.448"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockSquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 2h-6.5A6.76 6.76 0 0 0 2 8.75v6.5A6.76 6.76 0 0 0 8.75 22h6.5A6.76 6.76 0 0 0 22 15.25v-6.5A6.76 6.76 0 0 0 15.25 2m1.26 13.44a2 2 0 0 1-2 1.94h-5a2 2 0 0 1-2-1.94V12a2 2 0 0 1 1.42-1.85v-.44a3.08 3.08 0 0 1 .91-2.18a3.17 3.17 0 0 1 4.36 0a3.059 3.059 0 0 1 .91 2.18v.44A2 2 0 0 1 16.53 12z"/><path fill="currentColor" d="M13.5 9.71v.35h-3v-.35a1.5 1.5 0 0 1 3 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Login(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M10.973 2.755h5.136a4.918 4.918 0 0 1 5.136 4.623v9.244a4.918 4.918 0 0 1-5.136 4.623h-5.136"/><path stroke-miterlimit="10" d="M16.109 12H2.755"/><path stroke-linejoin="round" d="m11.397 17.136l4.404-4.404a1.04 1.04 0 0 0 0-1.464l-4.405-4.404"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Logout(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M13.477 21.245H8.34a4.918 4.918 0 0 1-5.136-4.623V7.378A4.918 4.918 0 0 1 8.34 2.755h5.136"/><path stroke-miterlimit="10" d="M20.795 12H7.442"/><path stroke-linejoin="round" d="m16.083 17.136l4.404-4.404a1.04 1.04 0 0 0 0-1.464l-4.404-4.404"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagnetDown(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M3.304 20.163v-8.404a8.893 8.893 0 0 1 8.12-8.99a8.697 8.697 0 0 1 9.273 8.697v8.697a1.087 1.087 0 0 1-1.087 1.087h-3.26a1.087 1.087 0 0 1-1.087-1.087v-8.697a3.261 3.261 0 1 0-6.523 0v8.697a1.087 1.087 0 0 1-1.087 1.087H4.391a1.087 1.087 0 0 1-1.087-1.087Zm11.958-5.436h5.435m-17.393 0h5.435"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagnetDownFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.646 7.675a9.471 9.471 0 0 0-2.21-3.14a9.32 9.32 0 0 0-3.28-2a9.55 9.55 0 0 0-12.63 9.2v8.41a1.78 1.78 0 0 0 .54 1.29a1.81 1.81 0 0 0 1.3.54h3.28a1.831 1.831 0 0 0 1.83-1.83v-8.7a2.54 2.54 0 0 1 .74-1.78a2.57 2.57 0 0 1 3.55 0a2.5 2.5 0 0 1 .74 1.78v8.7a1.83 1.83 0 0 0 1.83 1.83h3.31a1.831 1.831 0 0 0 1.83-1.83v-8.7a9.422 9.422 0 0 0-.83-3.77m-12.69 12.47a.32.32 0 0 1-.2.305a.331.331 0 0 1-.13.025h-3.26a.31.31 0 0 1-.24-.1a.29.29 0 0 1-.1-.23v-4.69h3.93zm12 0a.34.34 0 0 1-.33.33h-3.28a.34.34 0 0 1-.33-.33v-4.69h3.93z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagnetLeft(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M3.837 3.304h8.404a8.893 8.893 0 0 1 8.99 8.12a8.698 8.698 0 0 1-8.697 9.273H3.837A1.087 1.087 0 0 1 2.75 19.61v-3.26a1.087 1.087 0 0 1 1.087-1.087h8.697a3.261 3.261 0 1 0 0-6.523H3.837A1.087 1.087 0 0 1 2.75 7.652V4.391a1.087 1.087 0 0 1 1.087-1.087Zm5.436 11.958v5.435m0-17.393v5.435"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagnetLeftFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.973 11.355a9.63 9.63 0 0 0-9.73-8.8h-8.41a1.83 1.83 0 0 0-1.83 1.83v3.26a1.83 1.83 0 0 0 1.83 1.84h8.7a2.51 2.51 0 1 1 0 5.02h-8.7a1.88 1.88 0 0 0-1.3.54a1.86 1.86 0 0 0-.53 1.3v3.26c0 .486.19.953.53 1.3a1.88 1.88 0 0 0 1.3.54h8.72a9.41 9.41 0 0 0 8.89-6.27a9.301 9.301 0 0 0 .53-3.82m-13.45-3.37h-4.69a.29.29 0 0 1-.23-.1a.3.3 0 0 1-.1-.24v-3.26a.32.32 0 0 1 .2-.304a.33.33 0 0 1 .13-.026h4.69zm0 12h-4.69a.33.33 0 0 1-.23-.1a.309.309 0 0 1-.1-.24v-3.26a.36.36 0 0 1 .14-.29a.37.37 0 0 1 .24-.1h4.69z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagnetRight(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M20.163 20.696h-8.404a8.893 8.893 0 0 1-8.99-8.12a8.697 8.697 0 0 1 8.697-9.273h8.697A1.087 1.087 0 0 1 21.25 4.39v3.26a1.087 1.087 0 0 1-1.087 1.087h-8.697a3.261 3.261 0 1 0 0 6.523h8.697a1.087 1.087 0 0 1 1.087 1.087v3.261a1.087 1.087 0 0 1-1.087 1.087ZM14.727 8.738V3.303m0 17.393v-5.435"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagnetRightFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.461 9.485h8.7a1.82 1.82 0 0 0 1.3-.54a1.87 1.87 0 0 0 .54-1.3v-3.26a1.88 1.88 0 0 0-.54-1.3a1.86 1.86 0 0 0-1.3-.53h-8.72a9.42 9.42 0 0 0-8.89 6.26a9.35 9.35 0 0 0-.53 3.82a9.64 9.64 0 0 0 9.59 8.81h8.55a1.82 1.82 0 0 0 1.3-.54a1.87 1.87 0 0 0 .54-1.3v-3.26a1.87 1.87 0 0 0-.54-1.3a1.82 1.82 0 0 0-1.3-.54h-8.7a2.51 2.51 0 0 1-2.51-2.51a2.53 2.53 0 0 1 2.51-2.51m4-5.43h4.69a.36.36 0 0 1 .24.09a.37.37 0 0 1 .1.24v3.26a.35.35 0 0 1-.21.315a.33.33 0 0 1-.13.025h-4.69zm0 11.95h4.69a.34.34 0 0 1 .34.34v3.26a.35.35 0 0 1-.1.24a.37.37 0 0 1-.24.1h-4.69z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagnetUp(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M20.696 3.837v8.404a8.893 8.893 0 0 1-8.12 8.99a8.698 8.698 0 0 1-9.273-8.697V3.837A1.087 1.087 0 0 1 4.39 2.75h3.26a1.087 1.087 0 0 1 1.087 1.087v8.697a3.261 3.261 0 1 0 6.523 0V3.837a1.087 1.087 0 0 1 1.087-1.087h3.261a1.087 1.087 0 0 1 1.087 1.087ZM8.738 9.273H3.303m17.393 0h-5.435"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagnetUpFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.895 2.54a1.79 1.79 0 0 0-1.3-.54h-3.26a1.82 1.82 0 0 0-1.3.54a1.86 1.86 0 0 0-.53 1.3v8.7a2.5 2.5 0 0 1-.74 1.77a2.57 2.57 0 0 1-3.55 0a2.46 2.46 0 0 1-.74-1.77v-8.7c0-.486-.19-.952-.53-1.3a1.82 1.82 0 0 0-1.3-.54h-3.26a1.85 1.85 0 0 0-1.701 1.136a1.82 1.82 0 0 0-.14.704v8.7a9.44 9.44 0 0 0 3 6.91a9.468 9.468 0 0 0 6.46 2.55h.65a9.63 9.63 0 0 0 8.8-9.74v-8.4a1.82 1.82 0 0 0-.56-1.32m-12.92 6h-3.93V3.85a.31.31 0 0 1 .1-.24a.35.35 0 0 1 .24-.1h3.25a.33.33 0 0 1 .23.1a.31.31 0 0 1 .1.24zm12 0h-3.93V3.85a.34.34 0 0 1 .09-.24a.37.37 0 0 1 .24-.1h3.26a.35.35 0 0 1 .316.21a.311.311 0 0 1 .023.13z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Male(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.98 20.75a6.73 6.73 0 1 0 0-13.461a6.73 6.73 0 0 0 0 13.461m10.77-17.5l-6.004 6.004m6.004-.619V3.25h-5.385"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarker(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10" stroke-width="1.5"><path d="M12 12.8a3.35 3.35 0 1 0 0-6.7a3.35 3.35 0 0 0 0 6.7Z"/><path d="M12 2.75c-6.7 0-7.817 5.583-6.7 9.815c.983 3.708 3.93 6.242 5.874 8.32a1.117 1.117 0 0 0 1.652 0c1.943-2.078 4.891-4.612 5.874-8.32c1.117-4.232 0-9.815-6.7-9.815Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarkerFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.296 4.637a8.71 8.71 0 0 0-12.57 0c-1.53 2-2 5-1.14 8.08c.88 3.33 3.23 5.74 5.12 7.67l.92 1a2 2 0 0 0 .63.45a1.86 1.86 0 0 0 1.51 0a2 2 0 0 0 .62-.44l.93-1c1.89-1.93 4.24-4.34 5.12-7.67c.81-3.09.39-6.09-1.14-8.09m-6.29 8.25a3.48 3.48 0 1 1 3.215-2.15a3.471 3.471 0 0 1-3.215 2.14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Maximize(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 15.469v3.469a2.313 2.313 0 0 0 2.313 2.312H8.53m12.72-5.781v3.469a2.312 2.312 0 0 1-2.312 2.312h-3.47M2.75 8.531V5.062A2.312 2.312 0 0 1 5.063 2.75H8.53m6.939 0h3.469a2.313 2.313 0 0 1 2.312 2.313V8.53"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaReelH(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M2.946 8h18.108M2.946 16h18.108M12 2.75v18.5m5-.324V16m0-8V3.184M7 20.926V16m0-8V3.184"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaReelHfill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.235 17.589v3.659a6.165 6.165 0 0 1-2.976-3.66zm16.506 0a6.165 6.165 0 0 1-2.976 3.659v-3.66zm-10.623 0V22H8.14a5.542 5.542 0 0 1-1.176-.118v-4.294zm5.882 0v4.294c-.387.08-.781.12-1.177.117h-2.976v-4.41zm5-9.412h-9.118v7.647H22zm-10.882 0H2v7.647h9.118zM5.235 2.753v3.66H2.26a6.165 6.165 0 0 1 2.976-3.66m16.505 3.659H18.73V2.753a6.164 6.164 0 0 1 3.012 3.66M11.118 2v4.412H6.965V2.118c.387-.08.78-.12 1.176-.117zM17 2.118v4.294h-4.153V2h2.941c.407-.005.814.035 1.212.118"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaReelV(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M15.875 3.04v17.92M8.125 3.04v17.92M20.96 12H3.04m0 4.843h5.085m7.75 0h5.085M3.04 7.157h5.085m7.75 0h5.085"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MediaReelVfill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.342 18.577a6.118 6.118 0 0 1-3.906 3.2v-3.2zM22 12.882v2.941a4.5 4.5 0 0 1-.094.989h-4.47v-3.93zm0-4.706v2.942h-4.564v-3.93h4.47c.068.325.1.657.095.988m-.659-2.752h-3.906v-3.2a6.118 6.118 0 0 1 3.906 3.2M15.671 2H8.318v9.118h7.353zm-9.118.224v3.2H2.66a6.07 6.07 0 0 1 3.894-3.2m-.001 4.964v3.93H2V8.176a4.542 4.542 0 0 1 .095-.988zm0 5.694v3.93H2.095A4.542 4.542 0 0 1 2 15.823v-2.94zm0 5.695v3.2a6.071 6.071 0 0 1-3.894-3.2zm9.118-5.695H8.318V22h7.353z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Medium(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.57 17.576a5.57 5.57 0 1 0 0-11.141a5.57 5.57 0 0 0 0 11.141m8.944-.433c1.512 0 2.738-2.3 2.738-5.138c0-2.837-1.226-5.137-2.738-5.137s-2.738 2.3-2.738 5.137c0 2.838 1.226 5.138 2.738 5.138m4.408-.539c.59 0 1.067-2.06 1.067-4.599c0-2.54-.478-4.598-1.067-4.598c-.59 0-1.068 2.059-1.068 4.598c0 2.54.478 4.599 1.068 4.599"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MegaphoneA(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m5.549 10.819l-1.826 1.615a1.414 1.414 0 0 0-.288 1.77l1.653 2.9a1.404 1.404 0 0 0 1.662.629l2.297-.783z"/><path d="M9.258 4.59a26.669 26.669 0 0 1-1.71 4.072a7.178 7.178 0 0 1-2 2.157l3.499 6.112a7.294 7.294 0 0 1 2.882-.668c1.464.066 2.92.25 4.353.552"/><path d="m9.253 4.591l1.215-.706a1.395 1.395 0 0 1 1.917.517l5.607 9.774a1.42 1.42 0 0 1-.519 1.92l-1.215.707zM3.56 14.416l-.606.358a1.396 1.396 0 0 0-.658.86a1.411 1.411 0 0 0 .149 1.074a1.4 1.4 0 0 0 .854.662a1.384 1.384 0 0 0 1.068-.149l.567-.358m4.804-.203l1.701 2.97a1.44 1.44 0 0 1-.509 1.933a1.404 1.404 0 0 1-1.922-.522l-1.922-3.414m12.55-10.735l-2.498 1.45m4.612 3.531h-2.883M16.225 2.25l-1.442 2.515"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MegaphoneAfill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m18.44 13.791l-5.499-9.585a2.098 2.098 0 0 0-1.274-.98a2.137 2.137 0 0 0-1.607.205l-1.196.696a.236.236 0 0 0-.078.06l-.069.087a.627.627 0 0 0-.147.275a25.426 25.426 0 0 1-1.588 3.783a6.106 6.106 0 0 1-1.754 1.892l-1.794 1.588a2.136 2.136 0 0 0-.686 1.244a2.27 2.27 0 0 0 .059.98c-.246.136-.46.324-.627.55a2.107 2.107 0 0 0-.363.744a2.166 2.166 0 0 0 0 .833a2.137 2.137 0 0 0 .806 1.397c.22.168.473.29.743.358c.175.048.357.07.539.068c.097.01.196.01.294 0c.272-.033.533-.127.764-.274c.216.243.489.428.794.54c.237.093.49.14.745.136c.139.003.277-.013.412-.049l1.597 2.853a2.166 2.166 0 0 0 1.843 1.058a2.137 2.137 0 0 0 2.039-1.578a2.186 2.186 0 0 0-.206-1.607l-1.196-2.078c.322-.072.65-.111.98-.118c1.395.059 2.782.236 4.146.53h.343l.157-.07l.147-.088l1.049-.607c.482-.279.835-.737.98-1.274a2.137 2.137 0 0 0-.353-1.569M6.551 16.948a.647.647 0 0 1-.431 0a.598.598 0 0 1-.324-.275l-1.49-2.607l-.097-.167a.657.657 0 0 1 .137-.823l1.107-.98l.412.725l2.127 3.725zm10.674-1.96a.696.696 0 0 1-.304.401l-.549.314l-6.136-10.694l.559-.323a.598.598 0 0 1 .48-.059a.628.628 0 0 1 .392.294l5.499 9.576a.657.657 0 0 1 .059.549zm-.255-5.823a.716.716 0 0 1-.637-.362a.735.735 0 0 1 .265-.98l2.45-1.422a.725.725 0 0 1 .98.265a.735.735 0 0 1-.265.98l-2.45 1.421a.725.725 0 0 1-.343.098m4.518 3.47h-2.832a.735.735 0 1 1 0-1.47h2.832a.735.735 0 0 1 0 1.47m-6.831-6.969a.677.677 0 0 1-.363-.098a.726.726 0 0 1-.274-.98l1.411-2.47a.736.736 0 0 1 1.274.735l-1.411 2.46a.725.725 0 0 1-.637.353"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MegaphoneB(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.824 2.945a13.018 13.018 0 0 1-5.2 2.327L4.318 6.333A1.947 1.947 0 0 0 2.75 8.241v4.596a1.947 1.947 0 0 0 1.568 1.947l5.306 1.062c1.88.37 3.656 1.152 5.2 2.288a.974.974 0 0 0 1.558-.779V3.724a.974.974 0 0 0-1.558-.78"/><path d="m8.407 15.563l1.636 4.372a.973.973 0 0 1-.905 1.315h-1.14a1.947 1.947 0 0 1-1.83-1.266L4.22 14.707m8.267-7.089v5.842m6.037-7.059l2.531-1.46m-2.531 9.736l2.531 1.46m-2.726-5.597h2.921"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MegaphoneBfill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.87 2.822a1.7 1.7 0 0 0-.69-.64a1.85 1.85 0 0 0-.93-.18a1.77 1.77 0 0 0-.88.34a12.27 12.27 0 0 1-4.9 2.19l-5.3 1.06a2.79 2.79 0 0 0-1.56.94A2.71 2.71 0 0 0 2 8.242v4.58a2.73 2.73 0 0 0 .59 1.74c.282.354.65.63 1.07.8l1.8 4.88a2.74 2.74 0 0 0 1 1.28c.446.31.977.478 1.52.48h1.16a1.77 1.77 0 0 0 1.41-.75c.153-.228.252-.488.29-.76a1.75 1.75 0 0 0-.1-.82l-1.14-3.06a11.999 11.999 0 0 1 4.77 2.12c.257.191.562.309.88.34h.16c.265-.002.525-.067.76-.19a1.71 1.71 0 0 0 1-1.54V3.722a1.66 1.66 0 0 0-.3-.9m-3.39 10.64a1 1 0 0 1-2 0v-5.84a1 1 0 1 1 2 0zm5.04-6.31a.76.76 0 0 1-.65-.38a.74.74 0 0 1 .28-1l2.53-1.46a.74.74 0 0 1 1 .27a.75.75 0 0 1-.27 1l-2.53 1.46a.75.75 0 0 1-.36.11m2.56 9.73a.69.69 0 0 1-.37-.1l-2.53-1.46a.74.74 0 0 1-.28-1a.76.76 0 0 1 1-.28l2.53 1.47a.74.74 0 0 1 .27 1a.75.75 0 0 1-.62.37m.17-5.59h-2.92a.75.75 0 1 1 0-1.5h2.92a.75.75 0 0 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MemoryCard(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.85 12h-5.7a1.9 1.9 0 0 0-1.9 1.9v2.85c0 1.05.85 1.9 1.9 1.9h5.7a1.9 1.9 0 0 0 1.9-1.9V13.9a1.9 1.9 0 0 0-1.9-1.9M7.668 4.999v1.89m2.85-1.89v1.89m2.85-1.89v1.89"/><path d="M18.222 6.633v3.135h1.615V17.7a3.866 3.866 0 0 1-3.923 3.8H8.086a3.866 3.866 0 0 1-3.923-3.8V6.3a3.866 3.866 0 0 1 3.923-3.8h7.828a3.866 3.866 0 0 1 3.923 3.8v.333z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MemoryCardFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.835 7.38a.75.75 0 0 0 .75-.75v-.35a4.61 4.61 0 0 0-4.67-4.53h-7.89a4.61 4.61 0 0 0-4.61 4.55v11.41a4.61 4.61 0 0 0 4.6 4.54h7.89a4.68 4.68 0 0 0 3.28-1.3a4.621 4.621 0 0 0 1.4-3.25V9.76a.75.75 0 0 0-.75-.75h-.86V7.38zM12.625 5a.75.75 0 1 1 1.5 0v1.89a.75.75 0 0 1-1.5 0zm-2.86 0a.75.75 0 1 1 1.5 0v1.89a.75.75 0 0 1-1.5 0zm-2.85 0a.75.75 0 0 1 1.5 0v1.89a.75.75 0 1 1-1.5 0zm10.59 11.75a2.65 2.65 0 0 1-2.65 2.65h-5.7a2.65 2.65 0 0 1-2.65-2.65V13.9a2.65 2.65 0 0 1 2.65-2.65h5.7a2.65 2.65 0 0 1 2.65 2.65z"/><path fill="currentColor" d="M16.005 13.9v2.85a1.16 1.16 0 0 1-1.15 1.15h-5.7a1.15 1.15 0 0 1-1.15-1.15V13.9a1.14 1.14 0 0 1 1.15-1.15h5.7a1.15 1.15 0 0 1 1.15 1.15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Message(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21.25 7.506v6.607a3.963 3.963 0 0 1-3.964 3.965h-2.643l-2.18 2.18a.636.636 0 0 1-.925 0l-2.18-2.18H6.713a3.964 3.964 0 0 1-3.964-3.965V7.506a3.964 3.964 0 0 1 3.964-3.964h10.572a3.964 3.964 0 0 1 3.964 3.964"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageCheck(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.25 7.506v6.607a3.963 3.963 0 0 1-3.964 3.965h-2.643l-2.18 2.18a.636.636 0 0 1-.925 0l-2.18-2.18H6.713a3.964 3.964 0 0 1-3.964-3.965V7.506a3.964 3.964 0 0 1 3.964-3.964h10.572a3.964 3.964 0 0 1 3.964 3.964"/><path d="m8.447 10.685l2 2a.756.756 0 0 0 1.076 0l4.03-4.03"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageCheckFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.6 4.17a4.67 4.67 0 0 0-3.33-1.38H6.7a4.71 4.71 0 0 0-4.72 4.72v6.6a4.71 4.71 0 0 0 4.72 4.72h2.33l1.95 1.94c.127.143.284.255.46.33c.171.07.355.108.54.11a1.41 1.41 0 0 0 .56-.12a1.37 1.37 0 0 0 .44-.3l2-2h2.33a4.72 4.72 0 0 0 4.71-4.72v-6.6a4.7 4.7 0 0 0-1.42-3.3m-4.36 5.19l-4 4a1.7 1.7 0 0 1-1.24.52a1.79 1.79 0 0 1-.68-.13a1.92 1.92 0 0 1-.57-.39l-2-2a1 1 0 1 1 1.42-1.41L11 11.78l3.86-3.86a1 1 0 0 1 1.41 1.41z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageCheckRound(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.25 12a9.226 9.226 0 0 1-2.705 6.54A9.251 9.251 0 0 1 12 21.25a9.189 9.189 0 0 1-3.795-.81l-3.867.572a1.195 1.195 0 0 1-1.361-1.43l.537-3.923A8.943 8.943 0 0 1 2.75 12a9.228 9.228 0 0 1 2.705-6.54A9.25 9.25 0 0 1 12 2.75a9.26 9.26 0 0 1 6.545 2.71A9.236 9.236 0 0 1 21.25 12"/><path d="m7.88 11.887l2.319 2.32a.875.875 0 0 0 1.248 0l4.673-4.674"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageCheckRoundFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.265 8.18a9.659 9.659 0 0 0-2.17-3.25a10 10 0 0 0-10.9-2.17a10.18 10.18 0 0 0-3.25 2.17a10 10 0 0 0-2.16 3.25a9.83 9.83 0 0 0-.76 3.82a9.56 9.56 0 0 0 .74 3.77l-.5 3.65a2 2 0 0 0 0 .94c.093.304.25.583.46.82a2 2 0 0 0 .79.5a2 2 0 0 0 .93.07l3.65-.54a9.719 9.719 0 0 0 3.88.79a9.88 9.88 0 0 0 3.83-.76a10 10 0 0 0 3.24-2.17a9.72 9.72 0 0 0 2.17-3.24a10 10 0 0 0 0-7.65zm-4.41 2.06l-4.68 4.67c-.17.18-.38.32-.61.41a1.792 1.792 0 0 1-1.44 0a1.812 1.812 0 0 1-.62-.41l-2.31-2.31a1.004 1.004 0 0 1 1.42-1.42l2.24 2.24l4.58-4.59a1 1 0 1 1 1.42 1.41"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageConversation(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12.686 18.222c.401.078.809.119 1.218.123a7.276 7.276 0 0 0 2.992-.64l3.095.455a1.03 1.03 0 0 0 1.032-1.135l-.372-3.126a7.222 7.222 0 0 0 .599-2.9a7.304 7.304 0 0 0-7.336-7.284a7.295 7.295 0 0 0-7.109 5.654"/><path d="M13.904 14.745a5.664 5.664 0 0 1-1.218 3.477a5.614 5.614 0 0 1-4.375 2.063a5.665 5.665 0 0 1-2.29-.495l-2.311.413a.754.754 0 0 1-.826-.877l.32-2.363a5.5 5.5 0 0 1-.454-2.219A5.582 5.582 0 0 1 6.805 9.37a5.231 5.231 0 0 1 1.517-.217a5.592 5.592 0 0 1 5.582 5.593"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageConversationFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.98 10.995a8.11 8.11 0 0 0-.63-3.08a7.89 7.89 0 0 0-1.75-2.61a8 8 0 0 0-5.67-2.34a8 8 0 0 0-7.73 5.82a6.25 6.25 0 0 0-2.9 2.14a6.33 6.33 0 0 0-1.28 3.81a6.12 6.12 0 0 0 .43 2.32l-.29 2.14a1.51 1.51 0 0 0 1 1.67c.236.086.49.11.74.07l2.09-.37a6.429 6.429 0 0 0 2.38.47c.94 0 1.87-.208 2.72-.61a6.2 6.2 0 0 0 1.92-1.41c.308.046.618.073.93.08c1.06 0 2.11-.21 3.09-.62l2.88.42h.11a1.79 1.79 0 0 0 1.32-.58a1.73 1.73 0 0 0 .37-.64c.08-.241.11-.497.09-.75l-.35-2.93a7.94 7.94 0 0 0 .53-3m-9.88 6.69a4.849 4.849 0 0 1-1.69 1.32a5.49 5.49 0 0 1-2.1.47a4.998 4.998 0 0 1-2-.43a.63.63 0 0 0-.3-.07h-.13l-2.27.38l.33-2.36a.89.89 0 0 0-.06-.4a4.73 4.73 0 0 1-.39-1.93a4.82 4.82 0 0 1 3.56-4.58a5.49 5.49 0 0 1 1.28-.18a4.85 4.85 0 0 1 3.42 1.43a4.792 4.792 0 0 1 1.41 3.41a4.9 4.9 0 0 1-1.06 2.94"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageDots(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.25 7.506v6.607a3.963 3.963 0 0 1-3.964 3.965h-2.643l-2.18 2.18a.636.636 0 0 1-.925 0l-2.18-2.18H6.713a3.964 3.964 0 0 1-3.964-3.965V7.506a3.964 3.964 0 0 1 3.964-3.964h10.572a3.964 3.964 0 0 1 3.964 3.964"/><path d="M12 11.338a.66.66 0 1 0 0-1.321a.66.66 0 0 0 0 1.321m4.625 0a.66.66 0 1 0 0-1.321a.66.66 0 0 0 0 1.321m-9.25 0a.66.66 0 1 0 0-1.321a.66.66 0 0 0 0 1.321"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageDotsCheck(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 11.338a.66.66 0 1 0 0-1.321a.66.66 0 0 0 0 1.321m4.625 0a.66.66 0 1 0 0-1.321a.66.66 0 0 0 0 1.321m-9.25 0a.66.66 0 1 0 0-1.321a.66.66 0 0 0 0 1.321"/><path d="M21.25 10.94v3.17a4 4 0 0 1-1.16 2.81a4.16 4.16 0 0 1-1.29.86a4.11 4.11 0 0 1-1.51.3h-2.65l-2.18 2.18a.8.8 0 0 1-.21.15a.65.65 0 0 1-.5 0a.8.8 0 0 1-.21-.15l-2.18-2.18H6.71a4 4 0 0 1-2.8-1.16a4 4 0 0 1-1.16-2.81v-6.6a4 4 0 0 1 4-4H13m3.25 1.929l1.266 1.266a.479.479 0 0 0 .682 0l2.552-2.552"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageDotsCross(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M12 11.338a.66.66 0 1 0 0-1.321a.66.66 0 0 0 0 1.321m4.625 0a.66.66 0 1 0 0-1.321a.66.66 0 0 0 0 1.321m-9.25 0a.66.66 0 1 0 0-1.321a.66.66 0 0 0 0 1.321"/><path stroke-linejoin="round" d="M21.25 10.94v3.17a4 4 0 0 1-1.16 2.81a4.16 4.16 0 0 1-1.29.86a4.11 4.11 0 0 1-1.51.3h-2.65l-2.18 2.18a.8.8 0 0 1-.21.15a.65.65 0 0 1-.5 0a.8.8 0 0 1-.21-.15l-2.18-2.18H6.71a4 4 0 0 1-2.8-1.16a4 4 0 0 1-1.16-2.81v-6.6a4 4 0 0 1 4-4H13"/><path stroke-miterlimit="10" d="m20.25 3.75l-3.5 3.492m0-3.484l3.5 3.492"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageDotsDownload(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M12 11.338a.66.66 0 1 0 0-1.321a.66.66 0 0 0 0 1.321m4.625 0a.66.66 0 1 0 0-1.321a.66.66 0 0 0 0 1.321m-9.25 0a.66.66 0 1 0 0-1.321a.66.66 0 0 0 0 1.321"/><path stroke-linejoin="round" d="M21.25 10.94v3.17a4 4 0 0 1-1.16 2.81a4.16 4.16 0 0 1-1.29.86a4.11 4.11 0 0 1-1.51.3h-2.65l-2.18 2.18a.8.8 0 0 1-.21.15a.65.65 0 0 1-.5 0a.8.8 0 0 1-.21-.15l-2.18-2.18H6.71a4 4 0 0 1-2.8-1.16a4 4 0 0 1-1.16-2.81v-6.6a4 4 0 0 1 4-4H13"/><path stroke-miterlimit="10" d="M18.477 8.25v-4.5"/><path stroke-linejoin="round" d="m16.414 6.356l1.77 1.77a.413.413 0 0 0 .586 0l1.77-1.77"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageDotsFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.605 4.17a4.67 4.67 0 0 0-3.33-1.38H6.705a4.71 4.71 0 0 0-4.71 4.72v6.6a4.71 4.71 0 0 0 4.71 4.72h2.33l1.95 1.94c.127.143.284.255.46.33c.175.072.361.11.55.11c.189-.002.375-.04.55-.11a1.58 1.58 0 0 0 .44-.31l2-2h2.33a4.69 4.69 0 0 0 3.33-1.38a4.8 4.8 0 0 0 1-1.53c.234-.575.357-1.19.36-1.81v-6.6a4.67 4.67 0 0 0-1.4-3.3m-13.24 8.17a1.66 1.66 0 1 1 1.66-1.66a1.67 1.67 0 0 1-1.66 1.66m4.63 0a1.66 1.66 0 1 1 0-3.32a1.66 1.66 0 0 1 0 3.32m4.62 0a1.66 1.66 0 1 1 1.66-1.66a1.67 1.67 0 0 1-1.66 1.66"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageDotsMinus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M12 11.338a.66.66 0 1 0 0-1.321a.66.66 0 0 0 0 1.321m4.625 0a.66.66 0 1 0 0-1.321a.66.66 0 0 0 0 1.321m-9.25 0a.66.66 0 1 0 0-1.321a.66.66 0 0 0 0 1.321"/><path stroke-linejoin="round" d="M21.25 10.94v3.17a4 4 0 0 1-1.16 2.81a4.16 4.16 0 0 1-1.29.86a4.11 4.11 0 0 1-1.51.3h-2.65l-2.18 2.18a.8.8 0 0 1-.21.15a.65.65 0 0 1-.5 0a.8.8 0 0 1-.21-.15l-2.18-2.18H6.71a4 4 0 0 1-2.8-1.16a4 4 0 0 1-1.16-2.81v-6.6a4 4 0 0 1 4-4H13"/><path stroke-miterlimit="10" d="M17.024 5.87h4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageDotsPlus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M12 11.338a.66.66 0 1 0 0-1.321a.66.66 0 0 0 0 1.321m4.625 0a.66.66 0 1 0 0-1.321a.66.66 0 0 0 0 1.321m-9.25 0a.66.66 0 1 0 0-1.321a.66.66 0 0 0 0 1.321"/><path stroke-linejoin="round" d="M21.25 10.94v3.17a4 4 0 0 1-1.16 2.81a4.16 4.16 0 0 1-1.29.86a4.11 4.11 0 0 1-1.51.3h-2.65l-2.18 2.18a.8.8 0 0 1-.21.15a.65.65 0 0 1-.5 0a.8.8 0 0 1-.21-.15l-2.18-2.18H6.71a4 4 0 0 1-2.8-1.16a4 4 0 0 1-1.16-2.81v-6.6a4 4 0 0 1 4-4H13"/><path stroke-miterlimit="10" d="M18.519 2.75v5m-2.495-2.495h5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageDotsQuestionMark(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M12 11.338a.66.66 0 1 0 0-1.321a.66.66 0 0 0 0 1.321m4.625 0a.66.66 0 1 0 0-1.321a.66.66 0 0 0 0 1.321m-9.25 0a.66.66 0 1 0 0-1.321a.66.66 0 0 0 0 1.321"/><path stroke-linejoin="round" d="M21.25 10.94v3.17a4 4 0 0 1-1.16 2.81a4.16 4.16 0 0 1-1.29.86a4.11 4.11 0 0 1-1.51.3h-2.65l-2.18 2.18a.8.8 0 0 1-.21.15a.65.65 0 0 1-.5 0a.8.8 0 0 1-.21-.15l-2.18-2.18H6.71a4 4 0 0 1-2.8-1.16a4 4 0 0 1-1.16-2.81v-6.6a4 4 0 0 1 4-4H13"/><path stroke-miterlimit="10" d="M17.614 3.45a1.2 1.2 0 0 1 1.306-.68a1.163 1.163 0 0 1 .853.607a1.01 1.01 0 0 1-.588 1.389a.745.745 0 0 0-.48.673v.277"/><path stroke-linejoin="round" d="M18.74 7.896h.002"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageDotsRound(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.25 12a9.226 9.226 0 0 1-2.705 6.54A9.251 9.251 0 0 1 12 21.25a9.189 9.189 0 0 1-3.795-.81l-3.867.572a1.195 1.195 0 0 1-1.361-1.43l.537-3.923A8.943 8.943 0 0 1 2.75 12a9.228 9.228 0 0 1 2.705-6.54A9.25 9.25 0 0 1 12 2.75a9.26 9.26 0 0 1 6.545 2.71A9.236 9.236 0 0 1 21.25 12"/><path d="M12 12.61a.61.61 0 1 0 0-1.221a.61.61 0 0 0 0 1.221m4.279 0a.61.61 0 1 0 0-1.221a.61.61 0 0 0 0 1.221m-8.558 0a.61.61 0 1 0 .001-1.221a.61.61 0 0 0 0 1.221"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageDotsRoundCheck(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.25 12a9.35 9.35 0 0 1-.7 3.54a9.31 9.31 0 0 1-5 5a9.31 9.31 0 0 1-7.34-.11L4.34 21c-.19.045-.39.045-.58 0a1.28 1.28 0 0 1-.48-.31A1.15 1.15 0 0 1 3 19.58l.53-3.92A8.86 8.86 0 0 1 2.75 12a9.35 9.35 0 0 1 .7-3.54a9.31 9.31 0 0 1 5-5A9.31 9.31 0 0 1 12 2.75"/><path d="M12 12.61a.61.61 0 1 0 0-1.221a.61.61 0 0 0 0 1.221m4.279 0a.61.61 0 1 0 0-1.221a.61.61 0 0 0 0 1.221m-8.558 0a.61.61 0 1 0 .001-1.221a.61.61 0 0 0 0 1.221m7.528-6.742l1.689 1.688a.638.638 0 0 0 .908 0l3.403-3.403"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageDotsRoundCross(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M21.25 12a9.35 9.35 0 0 1-.7 3.54a9.31 9.31 0 0 1-5 5a9.31 9.31 0 0 1-7.34-.11L4.34 21c-.19.045-.39.045-.58 0a1.28 1.28 0 0 1-.48-.31A1.15 1.15 0 0 1 3 19.58l.53-3.92A8.86 8.86 0 0 1 2.75 12a9.35 9.35 0 0 1 .7-3.54a9.31 9.31 0 0 1 5-5A9.31 9.31 0 0 1 12 2.75"/><path stroke-linejoin="round" d="M12 12.61a.61.61 0 1 0 0-1.221a.61.61 0 0 0 0 1.221m4.279 0a.61.61 0 1 0 0-1.221a.61.61 0 0 0 0 1.221m-8.558 0a.61.61 0 1 0 .001-1.221a.61.61 0 0 0 0 1.221"/><path stroke-miterlimit="10" d="m20.25 2.75l-4.5 4.49m0-4.48l4.5 4.49"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageDotsRoundDownload(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M21.25 12a9.35 9.35 0 0 1-.7 3.54a9.31 9.31 0 0 1-5 5a9.31 9.31 0 0 1-7.34-.11L4.34 21c-.19.045-.39.045-.58 0a1.28 1.28 0 0 1-.48-.31A1.15 1.15 0 0 1 3 19.58l.53-3.92A8.86 8.86 0 0 1 2.75 12a9.35 9.35 0 0 1 .7-3.54a9.31 9.31 0 0 1 5-5A9.31 9.31 0 0 1 12 2.75"/><path stroke-miterlimit="10" d="M17.936 8.25v-5.5"/><path stroke-linejoin="round" d="m15.414 5.935l2.164 2.164a.505.505 0 0 0 .716 0l2.164-2.164M12 12.61a.61.61 0 1 0 0-1.221a.61.61 0 0 0 0 1.221m4.279 0a.61.61 0 1 0 0-1.221a.61.61 0 0 0 0 1.221m-8.558 0a.61.61 0 1 0 .001-1.221a.61.61 0 0 0 0 1.221"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageDotsRoundFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.25 8.18a9.78 9.78 0 0 0-2.16-3.25a10 10 0 0 0-14.15 0a9.76 9.76 0 0 0-2.17 3.25A10 10 0 0 0 2.01 12a9.74 9.74 0 0 0 .74 3.77l-.5 3.65a1.95 1.95 0 0 0 1.29 2.26c.297.098.613.122.92.07l3.65-.54a9.758 9.758 0 0 0 3.88.79a10 10 0 0 0 9.24-13.82zM7.73 13.61a1.61 1.61 0 1 1 .001-3.22a1.61 1.61 0 0 1 0 3.22m4.28 0a1.61 1.61 0 1 1 .001-3.22a1.61 1.61 0 0 1 0 3.22m4.28 0a1.61 1.61 0 1 1 .001-3.22a1.61 1.61 0 0 1 0 3.22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageDotsRoundMinus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M21.25 12a9.35 9.35 0 0 1-.7 3.54a9.31 9.31 0 0 1-5 5a9.31 9.31 0 0 1-7.34-.11L4.34 21c-.19.045-.39.045-.58 0a1.28 1.28 0 0 1-.48-.31A1.15 1.15 0 0 1 3 19.58l.53-3.92A8.86 8.86 0 0 1 2.75 12a9.35 9.35 0 0 1 .7-3.54a9.31 9.31 0 0 1 5-5A9.31 9.31 0 0 1 12 2.75"/><path stroke-miterlimit="10" d="M15.024 5.87h6.226"/><path stroke-linejoin="round" d="M12 12.61a.61.61 0 1 0 0-1.221a.61.61 0 0 0 0 1.221m4.279 0a.61.61 0 1 0 0-1.221a.61.61 0 0 0 0 1.221m-8.558 0a.61.61 0 1 0 .001-1.221a.61.61 0 0 0 0 1.221"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageDotsRoundPlus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M21.25 12a9.35 9.35 0 0 1-.7 3.54a9.31 9.31 0 0 1-5 5a9.31 9.31 0 0 1-7.34-.11L4.34 21c-.19.045-.39.045-.58 0a1.28 1.28 0 0 1-.48-.31A1.15 1.15 0 0 1 3 19.58l.53-3.92A8.86 8.86 0 0 1 2.75 12a9.35 9.35 0 0 1 .7-3.54a9.31 9.31 0 0 1 5-5A9.31 9.31 0 0 1 12 2.75"/><path stroke-miterlimit="10" d="M17.768 2.75v5.5m-2.744-2.744h5.5"/><path stroke-linejoin="round" d="M12 12.61a.61.61 0 1 0 0-1.221a.61.61 0 0 0 0 1.221m4.279 0a.61.61 0 1 0 0-1.221a.61.61 0 0 0 0 1.221m-8.558 0a.61.61 0 1 0 .001-1.221a.61.61 0 0 0 0 1.221"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageDotsRoundQuestionMark(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M21.25 12a9.35 9.35 0 0 1-.7 3.54a9.31 9.31 0 0 1-5 5a9.31 9.31 0 0 1-7.34-.11L4.34 21c-.19.045-.39.045-.58 0a1.28 1.28 0 0 1-.48-.31A1.15 1.15 0 0 1 3 19.58l.53-3.92A8.86 8.86 0 0 1 2.75 12a9.35 9.35 0 0 1 .7-3.54a9.31 9.31 0 0 1 5-5A9.31 9.31 0 0 1 12 2.75"/><path stroke-linejoin="round" d="M12 12.61a.61.61 0 1 0 0-1.221a.61.61 0 0 0 0 1.221m4.279 0a.61.61 0 1 0 0-1.221a.61.61 0 0 0 0 1.221m-8.558 0a.61.61 0 1 0 .001-1.221a.61.61 0 0 0 0 1.221"/><path stroke-miterlimit="10" d="M16.614 3.605a1.465 1.465 0 0 1 1.597-.83a1.422 1.422 0 0 1 1.042.742a1.234 1.234 0 0 1-.719 1.697a.91.91 0 0 0-.587.822v.34"/><path stroke-linejoin="round" d="M17.921 8.25h.003"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageDotsRoundUpload(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M21.25 12a9.35 9.35 0 0 1-.7 3.54a9.31 9.31 0 0 1-5 5a9.31 9.31 0 0 1-7.34-.11L4.34 21c-.19.045-.39.045-.58 0a1.28 1.28 0 0 1-.48-.31A1.15 1.15 0 0 1 3 19.58l.53-3.92A8.86 8.86 0 0 1 2.75 12a9.35 9.35 0 0 1 .7-3.54a9.31 9.31 0 0 1 5-5A9.31 9.31 0 0 1 12 2.75"/><path stroke-miterlimit="10" d="M18.394 2.75v5.5"/><path stroke-linejoin="round" d="m20.917 5.065l-2.165-2.164a.503.503 0 0 0-.716 0l-2.164 2.164M12 12.61a.61.61 0 1 0 0-1.221a.61.61 0 0 0 0 1.221m4.279 0a.61.61 0 1 0 0-1.221a.61.61 0 0 0 0 1.221m-8.558 0a.61.61 0 1 0 .001-1.221a.61.61 0 0 0 0 1.221"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageDotsUpload(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M12 11.338a.66.66 0 1 0 0-1.321a.66.66 0 0 0 0 1.321m4.625 0a.66.66 0 1 0 0-1.321a.66.66 0 0 0 0 1.321m-9.25 0a.66.66 0 1 0 0-1.321a.66.66 0 0 0 0 1.321"/><path stroke-linejoin="round" d="M21.25 10.94v3.17a4 4 0 0 1-1.16 2.81a4.16 4.16 0 0 1-1.29.86a4.11 4.11 0 0 1-1.51.3h-2.65l-2.18 2.18a.8.8 0 0 1-.21.15a.65.65 0 0 1-.5 0a.8.8 0 0 1-.21-.15l-2.18-2.18H6.71a4 4 0 0 1-2.8-1.16a4 4 0 0 1-1.16-2.81v-6.6a4 4 0 0 1 4-4H13"/><path stroke-miterlimit="10" d="M18.853 3.75v4.5"/><path stroke-linejoin="round" d="m20.917 5.644l-1.771-1.77a.412.412 0 0 0-.586 0l-1.77 1.77"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m22.045 7.77l-.05 6.3a4.679 4.679 0 0 1-.36 1.81a4.688 4.688 0 0 1-2.55 2.55a4.68 4.68 0 0 1-1.81.36h-2.33l-2 2a1.27 1.27 0 0 1-.44.3a1.33 1.33 0 0 1-.55.12a1.44 1.44 0 0 1-.55-.11a1.39 1.39 0 0 1-.46-.33l-1.95-1.94h-2.33a4.71 4.71 0 0 1-4.71-4.72v-6.6a4.71 4.71 0 0 1 1.38-3.34a4.78 4.78 0 0 1 3.33-1.38h10.66a4.71 4.71 0 0 1 4.72 4.72z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageInfoRound(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-width="1.5" d="M21.25 12a9.226 9.226 0 0 1-2.705 6.54A9.251 9.251 0 0 1 12 21.25a9.189 9.189 0 0 1-3.795-.81l-3.867.572a1.195 1.195 0 0 1-1.361-1.43l.537-3.923A8.943 8.943 0 0 1 2.75 12a9.228 9.228 0 0 1 2.705-6.54A9.25 9.25 0 0 1 12 2.75a9.26 9.26 0 0 1 6.545 2.71A9.236 9.236 0 0 1 21.25 12m-9.191 4.543v-5.559"/><path stroke-width="2" d="M11.941 7.457h.006"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageInfoRoundFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.265 8.18a10 10 0 0 0-2.16-3.25a10 10 0 0 0-14.15 0a9.66 9.66 0 0 0-2.17 3.25a9.83 9.83 0 0 0-.76 3.82a9.56 9.56 0 0 0 .74 3.77l-.5 3.65c-.07.31-.07.631 0 .94a2 2 0 0 0 .46.81a1.92 1.92 0 0 0 1.72.58l3.65-.54a9.758 9.758 0 0 0 3.88.79a10 10 0 0 0 9.24-13.82zm-8.18 8.36a1 1 0 1 1-2 0v-5.55a1 1 0 0 1 2 0zm-1.11-7.85a1.25 1.25 0 1 1 0-2.5a1.25 1.25 0 0 1 0 2.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageInformation(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-width="1.5" d="M21.25 7.506v6.607a3.963 3.963 0 0 1-3.964 3.965h-2.643l-2.18 2.18a.636.636 0 0 1-.925 0l-2.18-2.18H6.713a3.964 3.964 0 0 1-3.964-3.965V7.506a3.964 3.964 0 0 1 3.964-3.964h10.572a3.964 3.964 0 0 1 3.964 3.964m-9.244 7.743v-5.05"/><path stroke-width="2" d="M11.898 6.994h.006"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageInformationFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.575 4.17a4.67 4.67 0 0 0-3.33-1.38H6.715a4.73 4.73 0 0 0-4.72 4.72v6.6a4.73 4.73 0 0 0 4.72 4.72h2.33l1.95 1.94c.13.14.286.253.46.33c.171.07.355.108.54.11a1.44 1.44 0 0 0 1-.42l2-2h2.33a4.73 4.73 0 0 0 3.33-1.38a4.79 4.79 0 0 0 1-1.53c.231-.575.35-1.19.35-1.81v-6.6a4.71 4.71 0 0 0-1.43-3.3m-7.61 11.08a1 1 0 0 1-2 0V10.2a1 1 0 1 1 2 0zm-1.1-7a1.25 1.25 0 1 1 0-2.5a1.25 1.25 0 0 1 0 2.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageMinus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M21.25 7.506v6.607a3.963 3.963 0 0 1-3.964 3.965h-2.643l-2.18 2.18a.636.636 0 0 1-.925 0l-2.18-2.18H6.713a3.964 3.964 0 0 1-3.964-3.965V7.506a3.964 3.964 0 0 1 3.964-3.964h10.572a3.964 3.964 0 0 1 3.964 3.964"/><path stroke-miterlimit="10" d="M8.459 10.87h7.082"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageMinusFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.6 4.17a4.67 4.67 0 0 0-3.29-1.38H6.7a4.71 4.71 0 0 0-4.72 4.72v6.6a4.71 4.71 0 0 0 4.72 4.72h2.33l1.95 1.94c.127.143.284.255.46.33c.171.07.355.108.54.11a1.41 1.41 0 0 0 .56-.12a1.37 1.37 0 0 0 .44-.3l2-2h2.33a4.72 4.72 0 0 0 4.71-4.72v-6.6a4.71 4.71 0 0 0-1.42-3.3m-5.08 7.7H8.44a1 1 0 0 1 0-2h7.08a1 1 0 1 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageMinusRound(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M21.25 12a9.226 9.226 0 0 1-2.705 6.54A9.251 9.251 0 0 1 12 21.25a9.189 9.189 0 0 1-3.795-.81l-3.867.572a1.195 1.195 0 0 1-1.361-1.43l.537-3.923A8.943 8.943 0 0 1 2.75 12a9.228 9.228 0 0 1 2.705-6.54A9.25 9.25 0 0 1 12 2.75a9.26 9.26 0 0 1 6.545 2.71A9.236 9.236 0 0 1 21.25 12"/><path stroke-miterlimit="10" d="M8.098 12h7.805"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageMinusRoundFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.265 8.18a9.659 9.659 0 0 0-2.17-3.25a10 10 0 0 0-10.9-2.17a10.18 10.18 0 0 0-3.25 2.17a10 10 0 0 0-2.16 3.25a9.83 9.83 0 0 0-.76 3.82a9.56 9.56 0 0 0 .74 3.77l-.5 3.65a2 2 0 0 0 0 .94c.093.304.25.583.46.82a2 2 0 0 0 .79.5a2 2 0 0 0 .93.07l3.65-.54a9.719 9.719 0 0 0 3.88.79a9.88 9.88 0 0 0 3.83-.76a10.001 10.001 0 0 0 3.24-2.17a9.72 9.72 0 0 0 2.17-3.24a10 10 0 0 0 0-7.65zM15.925 13h-7.8a1 1 0 1 1 0-2h7.8a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessagePlus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M21.25 7.506v6.607a3.963 3.963 0 0 1-3.964 3.965h-2.643l-2.18 2.18a.636.636 0 0 1-.925 0l-2.18-2.18H6.713a3.964 3.964 0 0 1-3.964-3.965V7.506a3.964 3.964 0 0 1 3.964-3.964h10.572a3.964 3.964 0 0 1 3.964 3.964"/><path stroke-miterlimit="10" d="M11.992 7.322v7.081M8.459 10.87h7.082"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessagePlusFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.595 4.17a4.78 4.78 0 0 0-3.33-1.38H6.695a4.71 4.71 0 0 0-4.72 4.72v6.6a4.71 4.71 0 0 0 4.72 4.72h2.36l1.94 1.94c.133.14.293.253.47.33a1.409 1.409 0 0 0 1.09 0a1.5 1.5 0 0 0 .45-.31l2-2h2.33a4.73 4.73 0 0 0 3.33-1.38a4.77 4.77 0 0 0 1-1.53a4.68 4.68 0 0 0 .36-1.81v-6.6a4.71 4.71 0 0 0-1.43-3.3m-5.08 7.7h-2.55v2.53a1 1 0 1 1-2 0v-2.53h-2.53a1 1 0 1 1 0-2h2.53V7.32a1 1 0 0 1 2 0v2.55h2.55a1 1 0 1 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessagePlusRound(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M21.25 12a9.226 9.226 0 0 1-2.705 6.54A9.251 9.251 0 0 1 12 21.25a9.189 9.189 0 0 1-3.795-.81l-3.867.572a1.195 1.195 0 0 1-1.361-1.43l.537-3.923A8.943 8.943 0 0 1 2.75 12a9.228 9.228 0 0 1 2.705-6.54A9.25 9.25 0 0 1 12 2.75a9.26 9.26 0 0 1 6.545 2.71A9.236 9.236 0 0 1 21.25 12"/><path stroke-miterlimit="10" d="M11.693 7.816v7.795M7.8 11.722h7.804"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessagePlusRoundFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.265 8.18a9.758 9.758 0 0 0-2.17-3.25a10 10 0 0 0-10.9-2.17a10.18 10.18 0 0 0-3.25 2.17A9.94 9.94 0 0 0 2.025 12a9.56 9.56 0 0 0 .74 3.77l-.5 3.65a1.87 1.87 0 0 0 0 .94a2 2 0 0 0 .46.82a2 2 0 0 0 .79.5c.296.098.612.122.92.07l3.66-.54a9.719 9.719 0 0 0 3.88.79a10 10 0 0 0 7.07-2.93a9.72 9.72 0 0 0 2.17-3.24a10 10 0 0 0 0-7.65zm-5.6 4.51h-2.91v2.89a1 1 0 1 1-2 0v-2.89h-2.89a1 1 0 1 1 0-2h2.89v-2.9a1 1 0 0 1 2 0v2.9h2.91a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageQuestionMark(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round"><path stroke-linejoin="round" stroke-width="1.5" d="M21.25 7.506v6.607a3.963 3.963 0 0 1-3.964 3.965h-2.643l-2.18 2.18a.636.636 0 0 1-.925 0l-2.18-2.18H6.713a3.964 3.964 0 0 1-3.964-3.965V7.506a3.964 3.964 0 0 1 3.964-3.964h10.572a3.964 3.964 0 0 1 3.964 3.964"/><path stroke-miterlimit="10" stroke-width="1.5" d="M9.539 8.185a2.615 2.615 0 0 1 2.85-1.482a2.537 2.537 0 0 1 1.86 1.325a2.201 2.201 0 0 1-1.283 3.029a1.625 1.625 0 0 0-1.047 1.468v.606"/><path stroke-linejoin="round" stroke-width="2" d="M11.898 15.766h.006m-.006 0h.006"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageQuestionMarkFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.615 4.17a4.69 4.69 0 0 0-3.33-1.38H6.705a4.71 4.71 0 0 0-4.71 4.72v6.6a4.71 4.71 0 0 0 4.71 4.72h2.34l1.94 1.94a1.34 1.34 0 0 0 1.55.32a1.11 1.11 0 0 0 .44-.3l2-2h2.34a4.73 4.73 0 0 0 3.33-1.38a4.77 4.77 0 0 0 1-1.53a4.68 4.68 0 0 0 .36-1.81v-6.6a4.71 4.71 0 0 0-1.39-3.3m-8.72 12.85a1.25 1.25 0 1 1 0-2.499a1.25 1.25 0 0 1 0 2.499m3.31-6.8a3.21 3.21 0 0 1-1.92 1.79a.65.65 0 0 0-.26.2a.61.61 0 0 0-.11.35v.57a1 1 0 1 1-2 0v-.6a2.64 2.64 0 0 1 1.69-2.41a1.22 1.22 0 0 0 .76-.68a1.16 1.16 0 0 0 .09-.49c0-.17-.038-.337-.11-.49a1.629 1.629 0 0 0-.45-.5a1.65 1.65 0 0 0-.66-.27a1.69 1.69 0 0 0-1.07.15a1.65 1.65 0 0 0-.72.76a1 1 0 1 1-1.82-.83a3.61 3.61 0 0 1 1.62-1.7a3.66 3.66 0 0 1 2.32-.35a3.51 3.51 0 0 1 2.56 1.84a3.2 3.2 0 0 1 .08 2.66"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageQuestionMarkRound(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round"><path stroke-linejoin="round" stroke-width="1.5" d="M21.25 12a9.226 9.226 0 0 1-2.705 6.54A9.251 9.251 0 0 1 12 21.25a9.189 9.189 0 0 1-3.795-.81l-3.867.572a1.195 1.195 0 0 1-1.361-1.43l.537-3.923A8.943 8.943 0 0 1 2.75 12a9.228 9.228 0 0 1 2.705-6.54A9.25 9.25 0 0 1 12 2.75a9.26 9.26 0 0 1 6.545 2.71A9.236 9.236 0 0 1 21.25 12"/><path stroke-miterlimit="10" stroke-width="1.5" d="M9.287 8.667a2.88 2.88 0 0 1 3.142-1.631a2.797 2.797 0 0 1 2.05 1.459a2.422 2.422 0 0 1-1.414 3.334a1.791 1.791 0 0 0-1.154 1.615v.667"/><path stroke-linejoin="round" stroke-width="2" d="M11.888 17.012h.006m-.006 0h.006"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageQuestionMarkRoundFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.27 8.181a9.76 9.76 0 0 0-2.17-3.25a10.001 10.001 0 0 0-14.15 0a10.18 10.18 0 0 0-2.17 3.25a9.82 9.82 0 0 0-.75 3.82a9.56 9.56 0 0 0 .74 3.77l-.5 3.65a1.87 1.87 0 0 0 0 .94c.08.306.239.585.46.81a2 2 0 0 0 .79.51c.296.098.612.122.92.07l3.66-.54a9.71 9.71 0 0 0 3.87.79a9.91 9.91 0 0 0 3.83-.76a10 10 0 0 0 3.24-2.17a9.72 9.72 0 0 0 2.17-3.24a10 10 0 0 0 0-7.65zm-9.35 10.08a1.25 1.25 0 1 1 0-2.5a1.25 1.25 0 0 1 0 2.5m3.55-7.39a3.339 3.339 0 0 1-.83 1.17a3.43 3.43 0 0 1-1.23.74a.79.79 0 0 0-.33.26a.77.77 0 0 0-.14.44v.63a1 1 0 0 1-2 0v-.66a2.75 2.75 0 0 1 .51-1.57a2.87 2.87 0 0 1 1.29-1a1.425 1.425 0 0 0 .89-.8a1.65 1.65 0 0 0 .11-.59a1.501 1.501 0 0 0-.14-.58a1.8 1.8 0 0 0-.53-.58a1.74 1.74 0 0 0-.77-.31a2 2 0 0 0-1.24.17a1.91 1.91 0 0 0-.84.89a1 1 0 0 1-1.64.263A1 1 0 0 1 8.4 8.24a3.89 3.89 0 0 1 1.74-1.82a3.93 3.93 0 0 1 2.5-.38a3.84 3.84 0 0 1 2.75 2a3.409 3.409 0 0 1 .08 2.84z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageRound(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21.25 12a9.226 9.226 0 0 1-2.705 6.54A9.251 9.251 0 0 1 12 21.25a9.189 9.189 0 0 1-3.795-.81l-3.867.572a1.195 1.195 0 0 1-1.361-1.43l.537-3.923A8.943 8.943 0 0 1 2.75 12a9.228 9.228 0 0 1 2.705-6.54A9.25 9.25 0 0 1 12 2.75a9.26 9.26 0 0 1 6.545 2.71A9.236 9.236 0 0 1 21.25 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageRoundFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.99 12a9.88 9.88 0 0 1-.76 3.83a9.721 9.721 0 0 1-2.17 3.24a10 10 0 0 1-3.24 2.17a9.88 9.88 0 0 1-3.83.76a9.72 9.72 0 0 1-3.88-.79l-3.65.54a1.939 1.939 0 0 1-2.21-2.33l.5-3.65A9.56 9.56 0 0 1 2.01 12a9.83 9.83 0 0 1 .76-3.82a10 10 0 0 1 2.16-3.25a10 10 0 0 1 14.15 0a9.66 9.66 0 0 1 2.17 3.25a9.83 9.83 0 0 1 .74 3.82"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageRoundSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M17.328 12a5.316 5.316 0 0 1-3.287 4.923a5.335 5.335 0 0 1-4.227-.061l-2.228.33a.688.688 0 0 1-.783-.825l.309-2.259A5.151 5.151 0 0 1 6.672 12A5.315 5.315 0 0 1 9.96 7.078a5.334 5.334 0 0 1 6.965 2.884c.267.646.404 1.339.403 2.038"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageRoundSquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 2h-6.5A6.76 6.76 0 0 0 2 8.75v6.5A6.76 6.76 0 0 0 8.75 22h6.5A6.76 6.76 0 0 0 22 15.25v-6.5A6.76 6.76 0 0 0 15.25 2m1.61 12A5.26 5.26 0 0 1 12 17.25a5.27 5.27 0 0 1-1.91-.37l-1.62.24a1.21 1.21 0 0 1-.65-.06a1.25 1.25 0 0 1-.57-.35a1.41 1.41 0 0 1-.36-1.27l.21-1.59a5.38 5.38 0 0 1-.34-1.87a5.23 5.23 0 0 1 1.54-3.7A5.26 5.26 0 0 1 16.88 14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M17.482 9.829v3.916a2.35 2.35 0 0 1-2.35 2.35h-1.566l-1.292 1.291a.376.376 0 0 1-.548 0l-1.292-1.292H8.868a2.35 2.35 0 0 1-2.35-2.35V9.83a2.35 2.35 0 0 1 2.35-2.35h6.265a2.35 2.35 0 0 1 2.35 2.35"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MessageSquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 2h-6.5A6.76 6.76 0 0 0 2 8.75v6.5A6.76 6.76 0 0 0 8.75 22h6.5A6.76 6.76 0 0 0 22 15.25v-6.5A6.76 6.76 0 0 0 15.25 2m2.51 11.61a3 3 0 0 1-.22 1.14a3 3 0 0 1-.65 1a2.81 2.81 0 0 1-1 .65a3 3 0 0 1-1.14.22h-1l-.89.89a1.189 1.189 0 0 1-.38.27a1.4 1.4 0 0 1-1 0a1.289 1.289 0 0 1-.42-.3l-.87-.87h-1a2.919 2.919 0 0 1-2.11-.87a3 3 0 0 1-.87-2.1v-3.49a3 3 0 0 1 3-3h5.56a3 3 0 0 1 3 3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Messenger(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.772 2.013A9.695 9.695 0 0 0 4.774 4.89a9.624 9.624 0 0 0-2.77 7.019a9.379 9.379 0 0 0 3.12 6.96a.835.835 0 0 1 .291.65v1.45c0 .487.21 1.032.827 1.032a.864.864 0 0 0 .384-.104l1.863-.824a.9.9 0 0 1 .652 0c5.204 1.415 11.049-1.31 12.469-6.67a9.652 9.652 0 0 0-1.708-8.726a9.716 9.716 0 0 0-3.618-2.816a9.753 9.753 0 0 0-4.512-.847m4.494 9.187c-2.55 4.49-2.561 3.48-5.763 1.346c-.373 0-.559.278-1.805 1.16c-1.246.882-1.455 1.253-1.828.858c-.372-.394.78-1.705 2.329-4.234c1.35-2.227 3.213.859 4.4 1.044a11.341 11.341 0 0 0 2.096-1.508c.816-.626 1.048-.87 1.374-.614c.326.255.047.614-.803 1.948"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Meta(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.035 13.42c.26-6.058 5.405-11.598 10.229-5.27c4.326-6.307 9.45-.643 9.679 4.959a6.981 6.981 0 0 1-.706 4.066a3.33 3.33 0 0 1-4.42 1.038c-1.43-.84-2.25-2.386-3.111-3.756c-.57-.944-1.141-1.877-1.701-2.832c-1.256 2.003-2.822 5.727-4.97 6.733a3.485 3.485 0 0 1-4.211-1.183a6.224 6.224 0 0 1-.789-3.755m18.02.623c-.083-2.563-.747-6.422-3.756-6.961c-1.41 0-2.272 1.317-2.998 2.344v.166c.789 1.162 1.473 2.396 2.21 3.59c.602.87 1.587 2.81 2.531 3.184c1.577.498 2.064-1.058 2.013-2.303zm-15.883-.25v.54a2.48 2.48 0 0 0 .56 1.701a1.463 1.463 0 0 0 1.691.27a5.52 5.52 0 0 0 1.608-1.712c1.038-1.473 1.92-3.008 2.915-4.502c0-.104-.176-.28-.217-.384c-.851-1.14-2.075-2.635-3.694-2.075c-2.074 1.245-2.79 3.932-2.863 6.163"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Microphone(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15.64 5.836c0-1.704-1.63-3.086-3.64-3.086c-2.01 0-3.64 1.382-3.64 3.086v5.575c0 1.704 1.63 3.086 3.64 3.086c2.01 0 3.64-1.382 3.64-3.086z"/><path d="M5.328 10.616a6.672 6.672 0 1 0 13.344 0M12 21.25v-3.962M8.36 21.25h7.28"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicrophoneFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.43 5.838v5.57a4.43 4.43 0 0 1-8.78 0v-5.57a4.43 4.43 0 0 1 8.78 0"/><path fill="currentColor" d="M12.79 17.998v2.5h2.89a.75.75 0 0 1 0 1.5H8.4a.75.75 0 1 1 0-1.5h2.85v-2.5a7.4 7.4 0 0 1-6.67-7.38a.75.75 0 0 1 1.5 0a5.92 5.92 0 1 0 11.84 0a.75.75 0 0 1 1.5 0a7.4 7.4 0 0 1-6.67 7.38z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicrophoneMute(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15.3 12.705a3.17 3.17 0 0 1-1.3 1.3m-4-.001a3 3 0 0 1-1.65-2.579V5.848A3.398 3.398 0 0 1 12 2.76a3.401 3.401 0 0 1 3.64 3.089v2.518M5.33 10.626a6.643 6.643 0 0 0 2 4.707c.204.242.428.466.67.67m2.82 1.179a6.673 6.673 0 0 0 7.221-3.735a6.662 6.662 0 0 0 .629-2.821M12 21.25v-3.998M8.36 21.25h7.28"/><path d="m20 4.01l-4.36 4.357l-5.63 5.627l-1.98 1.979L4 20"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicrophoneMuteFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.986 14.776a.731.731 0 0 1-.66-.4a.74.74 0 0 1 .3-1a2.452 2.452 0 0 0 1-1.001a.739.739 0 0 1 1.359.127a.75.75 0 0 1-.048.563a3.852 3.852 0 0 1-1.611 1.61a.73.73 0 0 1-.34.1"/><path fill="currentColor" d="M19.409 10.644a7.545 7.545 0 0 1-.7 3.141a7.475 7.475 0 0 1-4.823 4.003a6.439 6.439 0 0 1-1.151.2v2.511h2.892a.75.75 0 0 1 0 1.501H8.342a.75.75 0 0 1 0-1.5h2.892v-2.512a5.523 5.523 0 0 1-.57-.07a.752.752 0 0 1 .27-1.481c.294.052.592.078.89.08h.33a5.499 5.499 0 0 0 1.341-.19a5.924 5.924 0 0 0 2.282-1.18a5.903 5.903 0 0 0 2.13-4.543a.75.75 0 0 1 1.502 0zM20.52 4.56l-4.363 4.353l-5.623 5.623l-6.004 6.003a.79.79 0 0 1-.53.22a.75.75 0 0 1-.53-.22a.741.741 0 0 1 0-1.06l3.462-3.462l-.18-.2a7.293 7.293 0 0 1-2.172-5.183a.751.751 0 0 1 .74-.76a.76.76 0 0 1 .751.75a5.923 5.923 0 0 0 1.771 4.172l.14.16l.85-.84a4.004 4.004 0 0 1-.65-.77a3.662 3.662 0 0 1-.59-1.892V5.851a4.132 4.132 0 0 1 1.45-2.872a4.923 4.923 0 0 1 5.894 0a4.132 4.132 0 0 1 1.441 2.802v.79l3.092-3.091a.74.74 0 0 1 1.06 0a.751.751 0 0 1-.01 1.08"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicrosoftWindows(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.75 7.189V2.865c0-.102 0-.115.115-.115h8.622c.128 0 .14 0 .14.128V11.5c0 .128 0 .128-.14.128H2.865c-.102 0-.115 0-.115-.116zM7.189 21.25H2.865c-.102 0-.115 0-.115-.116V12.59c0-.128 0-.128.128-.128h8.635c.102 0 .115 0 .115.115v8.57c0 .09 0 .103-.116.103zM21.25 7.189v4.31c0 .116 0 .116-.116.116h-8.557c-.102 0-.128 0-.128-.115V2.865c0-.09 0-.102.115-.102h8.48c.206 0 .206 0 .206.205zm-8.763 9.661v-4.273c0-.09 0-.115.103-.09h8.621c.026 0 0 .09 0 .142v8.518a.064.064 0 0 1-.017.06a.064.064 0 0 1-.06.017H12.54s-.09 0-.077-.09V16.85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Minimize(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.531 21.25v-3.469A2.312 2.312 0 0 0 6.22 15.47H2.75m12.719 5.78v-3.469a2.312 2.312 0 0 1 2.312-2.312h3.469M8.531 2.75v3.469A2.312 2.312 0 0 1 6.22 8.53H2.75m18.5.001h-3.469A2.312 2.312 0 0 1 15.47 6.22V2.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Minus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.5 12h-15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusCircle(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17.014 12H6.986M12 21.5a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusCircleFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.75A10.25 10.25 0 1 0 22.25 12A10.26 10.26 0 0 0 12 1.75M17 13H7a1 1 0 1 1 0-2h10a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M16.882 12H7.118"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusSquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 2h-6.5A6.76 6.76 0 0 0 2 8.75v6.5A6.76 6.76 0 0 0 8.75 22h6.5A6.76 6.76 0 0 0 22 15.25v-6.5A6.76 6.76 0 0 0 15.25 2m1.63 11H7.12a1 1 0 0 1 0-2h9.76a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MobilePhone(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16.184 2.75H7.816c-.998 0-1.808.833-1.808 1.86v14.78c0 1.027.81 1.86 1.808 1.86h8.368c.998 0 1.808-.833 1.808-1.86V4.61c0-1.027-.81-1.86-1.808-1.86"/><path d="M12 18.773a.52.52 0 1 0 0-1.038a.52.52 0 0 0 0 1.038M10.003 5.272h3.994"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MobilePhoneFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.165 2h-8.35a2.59 2.59 0 0 0-2.56 2.61v14.78A2.59 2.59 0 0 0 7.815 22h8.37a2.59 2.59 0 0 0 2.56-2.61V4.61A2.588 2.588 0 0 0 16.165 2m-4.18 17.77a1.52 1.52 0 1 1-.02-3.04a1.52 1.52 0 0 1 .02 3.04m2-13.5h-4a1 1 0 1 1 0-2h4a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoneyExchange(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m5.795 14.306l-1.772-1.775l-1.773 1.775m15.955-4.579l1.772 1.776l1.773-1.776"/><path d="M19.977 11.503c0-2.12-.84-4.151-2.336-5.65A7.971 7.971 0 0 0 12 3.513a7.869 7.869 0 0 0-2.97.577a7.977 7.977 0 0 0-4.555 4.75m-.452 3.69a7.997 7.997 0 0 0 1.827 5.082a7.967 7.967 0 0 0 9.966 1.927a7.985 7.985 0 0 0 3.585-4.034"/><path d="M9.58 13.978A2.279 2.279 0 0 0 12 16.054c1.952 0 2.42-1.123 2.42-2.076c0-.953-.807-1.963-2.42-1.963s-2.42-.638-2.42-1.938a2.224 2.224 0 0 1 1.537-2.003c.285-.092.585-.125.883-.097a2.329 2.329 0 0 1 2.42 2.1M12 17.264v-1.051m0-9.45v1.21"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Moon(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21.182 14.86A9.474 9.474 0 1 1 9.139 2.819a1.053 1.053 0 0 1 1.38 1.295a7.705 7.705 0 0 0-.085 4.642a6.999 6.999 0 0 0 4.81 4.811c1.52.45 3.14.42 4.643-.084a1.053 1.053 0 0 1 1.295 1.379"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.88 15.147a10.23 10.23 0 0 1-19.63-5.64a10.28 10.28 0 0 1 6.63-7.37a1.77 1.77 0 0 1 1-.07a1.8 1.8 0 0 1 .89.45a1.81 1.81 0 0 1 .48 1.84a7 7 0 0 0-.08 4.21a6.272 6.272 0 0 0 4.3 4.31a6.92 6.92 0 0 0 4.2-.08a1.83 1.83 0 0 1 1 0a1.8 1.8 0 0 1 1.3 1.39a1.82 1.82 0 0 1-.09.96"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mouse(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.5 9.286C19.5 5.538 16.142 2.5 12 2.5c-4.142 0-7.5 3.038-7.5 6.786v5.428c0 3.748 3.358 6.786 7.5 6.786c4.142 0 7.5-3.038 7.5-6.786zM12 10.643V2.5m-7.5 8.143h15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MouseFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.745 11.306v3.81C19.745 18.935 16.303 22 12 22s-7.745-3.098-7.745-6.885v-3.81zM11.14 2v7.585H4.254v-.7A7.263 7.263 0 0 1 11.139 2m8.606 6.885v.7h-6.884V2a7.275 7.275 0 0 1 6.884 6.885"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MousePointer(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.4" d="m6.244 3.114l12.298 8.66A.693.693 0 0 1 18.346 13l-4.62.877a.565.565 0 0 0-.334.82l2.31 4.377a.693.693 0 0 1-.22.981l-1.663.866a.693.693 0 0 1-.935-.289l-2.31-4.387a.577.577 0 0 0-.866-.232L6.325 19.27a.692.692 0 0 1-1.155-.554V3.703a.693.693 0 0 1 1.074-.589"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MousePointerFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.512 12.566a1.41 1.41 0 0 1-.32.69a1.49 1.49 0 0 1-.64.42l-4.48.85l2.24 4.23c.088.137.15.29.18.45c.035.185.035.375 0 .56a1.309 1.309 0 0 1-.23.51c-.112.157-.255.29-.42.39l-1.69.88a1.36 1.36 0 0 1-.65.16c-.139 0-.277-.02-.41-.06a1.42 1.42 0 0 1-.82-.68l-2.25-4.27l-3.21 3.08a1.4 1.4 0 0 1-1.55.24a1.43 1.43 0 0 1-.79-1.33V3.696a1.51 1.51 0 0 1 .19-.71a1.35 1.35 0 0 1 .54-.52a1.385 1.385 0 0 1 1.43.06l12.3 8.66c.21.153.376.36.48.6c.107.245.141.515.1.78"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MouseTwo(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19.5 9.286C19.5 5.538 16.142 2.5 12 2.5c-4.142 0-7.5 3.038-7.5 6.786v5.428c0 3.748 3.358 6.786 7.5 6.786c4.142 0 7.5-3.038 7.5-6.786z"/><path d="M13.743 10.421c0-.986-.784-1.785-1.75-1.785c-.967 0-1.75.8-1.75 1.785v1.429c0 .986.783 1.786 1.75 1.786c.966 0 1.75-.8 1.75-1.786zM12 8.636V2.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MouseTwoFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.065 9.331v5.308C20.065 18.696 16.438 22 12 22s-8.065-3.304-8.065-7.36V9.33c0-3.841 3.226-6.999 7.332-7.331v6.08a2.464 2.464 0 0 0-1.72 2.356v1.398a2.444 2.444 0 1 0 4.887 0v-1.398a2.464 2.464 0 0 0-1.7-2.346V2c4.095.303 7.33 3.49 7.33 7.331"/><path fill="currentColor" d="M13 10.289v1.43a1 1 0 1 1-2 0v-1.43a1 1 0 0 1 2 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Multiply(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 5L5 19m14 0L5 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MultiplyCircle(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m15.958 8.042l-7.916 7.916m7.916 0L8.042 8.042M12 21.5a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MultiplyCircleFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.75A10.25 10.25 0 1 0 22.25 12A10.26 10.26 0 0 0 12 1.75m4.67 13.5a1 1 0 0 1 0 1.42a1.001 1.001 0 0 1-1.42 0L12 13.42l-3.25 3.25a1 1 0 0 1-1.41-1.42L10.59 12L7.34 8.75a1 1 0 1 1 1.41-1.41L12 10.59l3.25-3.25a1 1 0 1 1 1.42 1.41L13.42 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MultiplySquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m15.854 8.146l-7.708 7.708m7.708 0L8.146 8.146"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MultiplySquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 2h-6.5A6.76 6.76 0 0 0 2 8.75v6.5A6.76 6.76 0 0 0 8.75 22h6.5A6.76 6.76 0 0 0 22 15.25v-6.5A6.76 6.76 0 0 0 15.25 2m1.31 13.15a1 1 0 0 1 0 1.41a1 1 0 0 1-1.41 0L12 13.41l-3.15 3.15a1 1 0 0 1-1.41 0a1 1 0 0 1 0-1.41L10.59 12L7.44 8.85a1 1 0 1 1 1.41-1.41L12 10.59l3.15-3.15a1 1 0 1 1 1.41 1.41L13.41 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Music(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M7.97 21.5a4.03 4.03 0 1 0 0-8.06a4.03 4.03 0 0 0 0 8.06ZM12 17.47V2.5m0 0l6.39 1.82a2.303 2.303 0 0 1 1.67 2.21v.968a2.303 2.303 0 0 1-2.878 2.222L12 8.258"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicAlternate(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M17.985 19.703a3.265 3.265 0 1 0 0-6.53a3.265 3.265 0 0 0 0 6.53ZM6.015 20.79a3.265 3.265 0 1 0 0-6.528a3.265 3.265 0 0 0 0 6.529Zm15.235-4.352v-8.01M9.28 17.526v-6.921"/><path stroke-linejoin="round" d="M21.25 8.429L9.28 10.605v-3.57a2.176 2.176 0 0 1 1.784-2.176l8.902-1.632a1.088 1.088 0 0 1 1.284 1.088z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicAlternateFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.984 4.345v12.12a4 4 0 1 1-1.5-3.13v-4l-10.47 1.9v6.3a4 4 0 1 1-4-4a4 4 0 0 1 2.52.9v-7.41a2.93 2.93 0 0 1 .66-1.9a3 3 0 0 1 1.74-1l8.9-1.63a1.83 1.83 0 0 1 2 1.05c.111.251.163.525.15.8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicCircle(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><rect width="19" height="19" x="2.5" y="2.5" rx="9.5"/><path stroke-linecap="round" d="M9.667 17.5a2.333 2.333 0 1 0 0-4.667a2.333 2.333 0 0 0 0 4.667ZM12 15.167V6.5m0 0l3.7 1.053a1.333 1.333 0 0 1 .967 1.28v.56A1.334 1.334 0 0 1 15 10.68l-3-.847"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicCircleFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.75A10.25 10.25 0 1 0 22.25 12A10.26 10.26 0 0 0 12 1.75m5.42 7.64a2.22 2.22 0 0 1-.22.92a2.12 2.12 0 0 1-.6.73a2.06 2.06 0 0 1-.85.39c-.14.015-.28.015-.42 0c-.175 0-.35-.02-.52-.06l-2.06-.58v4.34a3.09 3.09 0 1 1-3.08-3.09a2.94 2.94 0 0 1 1.58.45v-6a.639.639 0 0 1 0-.2v-.08a.72.72 0 0 1 .54-.44h.44l3.61 1a2.09 2.09 0 0 1 1.09.75c.27.362.42.799.43 1.25z"/><path fill="currentColor" d="M11.25 15.17a1.59 1.59 0 1 1-1.58-1.59a1.58 1.58 0 0 1 1.58 1.59m4.67-6.34v.56a.77.77 0 0 1-.07.26a.45.45 0 0 1-.16.2a.74.74 0 0 1-.24.11a.53.53 0 0 1-.25 0l-2.45-.69V7.49l2.74.78a.66.66 0 0 1 .31.21c.068.105.11.225.12.35"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.82 6.49v1a3 3 0 0 1-.32 1.34a3 3 0 0 1-.87 1.07a3.09 3.09 0 0 1-1.25.57a3.306 3.306 0 0 1-1.38-.03l-4.24-1.2v8.22a4.79 4.79 0 1 1-1.5-3.47V2.46a.639.639 0 0 1 0-.2v-.08a.76.76 0 0 1 .54-.44h.44l6.3 1.79a3 3 0 0 1 2.22 2.93z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M9.667 17.5a2.333 2.333 0 1 0 0-4.667a2.333 2.333 0 0 0 0 4.667ZM12 15.167V6.5m0 0l3.7 1.053a1.333 1.333 0 0 1 .967 1.28v.56A1.334 1.334 0 0 1 15 10.68l-3-.847"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicSquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 2h-6.5A6.76 6.76 0 0 0 2 8.75v6.5A6.76 6.76 0 0 0 8.75 22h6.5A6.76 6.76 0 0 0 22 15.25v-6.5A6.76 6.76 0 0 0 15.25 2m2.17 7.39a2.22 2.22 0 0 1-.22.92a2.12 2.12 0 0 1-.6.73a2.06 2.06 0 0 1-.85.39c-.14.015-.28.015-.42 0c-.175 0-.35-.02-.52-.06l-2.06-.58v4.34a3.09 3.09 0 1 1-3.08-3.09a2.94 2.94 0 0 1 1.58.45v-6a.639.639 0 0 1 0-.2v-.08a.72.72 0 0 1 .54-.44a.34.34 0 0 1 .14 0h.3l3.61 1a2.09 2.09 0 0 1 1.09.75c.27.362.42.799.43 1.25z"/><path fill="currentColor" d="M11.25 15.17a1.59 1.59 0 1 1-1.58-1.59a1.58 1.58 0 0 1 1.58 1.59m4.67-6.34v.56a.77.77 0 0 1-.07.26a.45.45 0 0 1-.16.2a.74.74 0 0 1-.24.11a.53.53 0 0 1-.25 0l-2.45-.69V7.49l2.74.78a.66.66 0 0 1 .31.21c.068.105.11.225.12.35"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Netflix(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.518 13.187V21.5a4.982 4.982 0 0 0-1.783-.192a6.529 6.529 0 0 0-1.907.192v-19zM17.172 2.5v18.863l-3.389-9.603V2.5z"/><path fill="currentColor" d="M6.828 2.5h3.69l6.654 18.89a18.67 18.67 0 0 0-3.786 0C12.836 19.936 6.828 2.5 6.828 2.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Next(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20.062 20.25V3.75M3.938 5.416V18.58a1.447 1.447 0 0 0 2.329 1.143l9.113-7.088a1.447 1.447 0 0 0-.087-2.344L6.18 4.216a1.447 1.447 0 0 0-2.242 1.2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NextFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.095 21a.75.75 0 0 1-.75-.75V3.75a.75.75 0 0 1 1.5 0v16.5a.74.74 0 0 1-.75.75m-3.4-9.589a2.25 2.25 0 0 1-.85 1.82l-9.11 7.09c-.326.247-.713.4-1.12.44h-.23a2.142 2.142 0 0 1-1-.22a2.201 2.201 0 0 1-.9-.81a2.17 2.17 0 0 1-.33-1.16V5.421a2.22 2.22 0 0 1 .31-1.12a2.25 2.25 0 0 1 .85-.8a2.18 2.18 0 0 1 2.24.1l9.12 6.08c.29.191.53.448.7.75a2.3 2.3 0 0 1 .32.98"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Note(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.25 6.861v6.342a2.057 2.057 0 0 1-.606 1.459l-5.982 5.982a2.055 2.055 0 0 1-1.46.606h-6.34a4.111 4.111 0 0 1-4.112-4.111V6.86a4.111 4.111 0 0 1 4.111-4.11H17.14a4.111 4.111 0 0 1 4.111 4.111"/><path d="m14.056 21.075l-.514-4.11a3.082 3.082 0 0 1 3.443-3.444l4.11.514"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoteFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.57 3.42A4.82 4.82 0 0 0 17.13 2H6.89a4.87 4.87 0 0 0-4.86 4.86v10.27A4.88 4.88 0 0 0 6.89 22h6.35c.37-.003.737-.078 1.08-.22l.16-.08l.08-.05a2.44 2.44 0 0 0 .67-.48l6-6c.203-.2.372-.434.5-.69l.08-.17a.56.56 0 0 0 0-.12c.107-.295.16-.606.16-.92V6.93a4.82 4.82 0 0 0-1.4-3.51m-6 16.19l-.34-2.75a2.29 2.29 0 0 1 .11-1a2.22 2.22 0 0 1 .55-.89a2.28 2.28 0 0 1 .9-.56a2.321 2.321 0 0 1 1.05-.11l2.75.34z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoteText(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20.221 17.084v-8.11a4.166 4.166 0 0 0-4.166-4.197h-8.11A4.166 4.166 0 0 0 3.78 8.944v8.11a4.166 4.166 0 0 0 4.166 4.196h8.11a4.166 4.166 0 0 0 4.166-4.166M16.055 6.805V2.75m-8.11 4.055V2.75m-.507 8.11h9.124m-9.124 5.068h9.124"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoteTextFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.62 7.06a4.91 4.91 0 0 0-1.07-1.58a4.81 4.81 0 0 0-1.6-1.08a5.091 5.091 0 0 0-1.14-.32V2.75a.75.75 0 1 0-1.5 0v1.27H8.7V2.75a.75.75 0 1 0-1.5 0v1.33a4.84 4.84 0 0 0-2.73 1.38a4.89 4.89 0 0 0-1.44 3.48v8.1a5 5 0 0 0 .37 1.9a4.76 4.76 0 0 0 1.06 1.6c.454.464.998.832 1.6 1.08c.598.251 1.24.38 1.89.38h8.11a4.93 4.93 0 0 0 4.91-4.92V8.97a4.742 4.742 0 0 0-.35-1.91m-4 9.59H7.45a.75.75 0 1 1 0-1.5h9.1a.75.75 0 1 1 0 1.5zm0-5.07H7.45a.75.75 0 0 1 0-1.5h9.1a.75.75 0 1 1 0 1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoteWithText(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.25 6.861v6.342a2.057 2.057 0 0 1-.606 1.459l-5.982 5.982a2.055 2.055 0 0 1-1.46.606h-6.34a4.111 4.111 0 0 1-4.112-4.111V6.86a4.111 4.111 0 0 1 4.111-4.11H17.14a4.111 4.111 0 0 1 4.111 4.111"/><path d="m14.056 21.075l-.514-4.11a3.082 3.082 0 0 1 3.443-3.444l4.11.514M6.861 7.889h9.25M6.861 12h4.111m-4.111 4.111h3.083"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoteWithTextFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.57 3.42A4.82 4.82 0 0 0 17.13 2H6.89a4.87 4.87 0 0 0-4.86 4.86v10.27A4.88 4.88 0 0 0 6.89 22h6.35c.37-.003.737-.078 1.08-.22l.16-.08l.08-.05a2.44 2.44 0 0 0 .67-.48l6-6c.203-.2.372-.434.5-.69l.08-.17c.004-.04.063-.12.059-.16c.107-.296.102-.566.1-.88V6.93a4.82 4.82 0 0 0-1.4-3.51M9.94 16.86H6.85a.75.75 0 1 1 0-1.5h3.09a.75.75 0 0 1 0 1.5m1-4.11H6.82a.75.75 0 0 1 0-1.5h4.12a.75.75 0 0 1 0 1.5M6.82 8.63a.75.75 0 1 1 0-1.5h9.25a.75.75 0 0 1 0 1.5zm7.77 11l-.34-2.75a2.265 2.265 0 0 1 .11-1a2.22 2.22 0 0 1 .55-.89a2.28 2.28 0 0 1 .9-.56a2.32 2.32 0 0 1 1.05-.11l2.75.34z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationBell(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M11.962 17.986h6.81a1.555 1.555 0 0 0 1.512-2.175c-.36-1.088-1.795-2.393-1.795-3.677c0-2.85 0-3.6-1.404-5.276a5.025 5.025 0 0 0-1.653-1.283l-.783-.38a1.089 1.089 0 0 1-.511-.73a2.023 2.023 0 0 0-2.176-1.707a2.023 2.023 0 0 0-2.12 1.707a1.089 1.089 0 0 1-.567.73l-.783.38A5.025 5.025 0 0 0 6.84 6.858c-1.403 1.676-1.403 2.426-1.403 5.276c0 1.284-1.37 2.458-1.73 3.611c-.217.697-.337 2.241 1.48 2.241z"/><path d="M15.225 17.986a3.198 3.198 0 0 1-3.263 3.263A3.195 3.195 0 0 1 8.7 17.986"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationBellCheck(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M11.962 17.986h6.81a1.555 1.555 0 0 0 1.512-2.175c-.36-1.088-1.795-2.393-1.795-3.677c0-2.85 0-3.6-1.403-5.276a5.026 5.026 0 0 0-1.654-1.283l-.783-.38a1.088 1.088 0 0 1-.511-.73a2.023 2.023 0 0 0-2.176-1.707a2.024 2.024 0 0 0-2.12 1.707a1.089 1.089 0 0 1-.566.73l-.784.38A5.025 5.025 0 0 0 6.84 6.858c-1.403 1.676-1.403 2.426-1.403 5.276c0 1.284-1.37 2.458-1.73 3.611c-.217.697-.337 2.241 1.48 2.241z"/><path d="M15.226 17.986a3.196 3.196 0 0 1-3.264 3.263A3.195 3.195 0 0 1 8.7 17.986m.8-6.788l1.407 1.407a.534.534 0 0 0 .757 0L14.5 9.769"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationBellCheckFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.003 15.578a7.056 7.056 0 0 0-.87-1.57a4.09 4.09 0 0 1-.89-1.88c0-2.89 0-3.87-1.58-5.76a5.79 5.79 0 0 0-1.9-1.47l-.73-.35a.32.32 0 0 1-.1-.1a.23.23 0 0 1-.05-.1a2.77 2.77 0 0 0-2.93-2.34a2.77 2.77 0 0 0-2.84 2.29a.31.31 0 0 1-.07.14a.34.34 0 0 1-.09.08l-.78.38a5.63 5.63 0 0 0-1.91 1.48c-1.57 1.88-1.57 2.86-1.57 5.75a3.84 3.84 0 0 1-.82 1.77a6.568 6.568 0 0 0-.88 1.62a2.79 2.79 0 0 0 .26 2.37a2.24 2.24 0 0 0 1.94.85h2.82c.043.268.117.53.22.78c.198.497.498.947.88 1.32c.37.38.816.677 1.31.87c.46.188.953.287 1.45.29h.16a4 4 0 0 0 2.79-1.16a4 4 0 0 0 .87-1.31c.101-.255.178-.52.23-.79h2.94a2.372 2.372 0 0 0 1-.23a2.4 2.4 0 0 0 .88-.76c.226-.322.364-.698.4-1.09a2.2 2.2 0 0 0-.14-1.08m-6-5.28l-2.81 2.83c-.113.124-.253.22-.41.28a1.192 1.192 0 0 1-.49.1a1.26 1.26 0 0 1-.91-.38l-1.41-1.4a.75.75 0 0 1 0-1.06a.74.74 0 0 1 1.06 0l1.26 1.25l2.68-2.68a.75.75 0 1 1 1.06 1.06z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationBellCross(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M11.962 17.986h6.81a1.555 1.555 0 0 0 1.512-2.175c-.36-1.088-1.795-2.393-1.795-3.677c0-2.85 0-3.6-1.403-5.276a5.026 5.026 0 0 0-1.654-1.283l-.783-.38a1.088 1.088 0 0 1-.511-.73a2.023 2.023 0 0 0-2.176-1.707a2.024 2.024 0 0 0-2.12 1.707a1.089 1.089 0 0 1-.566.73l-.784.38A5.025 5.025 0 0 0 6.84 6.858c-1.403 1.676-1.403 2.426-1.403 5.276c0 1.284-1.37 2.458-1.73 3.611c-.217.697-.337 2.241 1.48 2.241z"/><path stroke-linejoin="round" d="M15.226 17.986a3.196 3.196 0 0 1-3.264 3.263A3.195 3.195 0 0 1 8.7 17.986"/><path stroke-miterlimit="10" d="m13.75 9.497l-3.5 3.492m0-3.485l3.5 3.493"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationBellCrossFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.003 15.578a7.056 7.056 0 0 0-.87-1.57a4.09 4.09 0 0 1-.89-1.88c0-2.89 0-3.87-1.58-5.76a5.79 5.79 0 0 0-1.9-1.47l-.73-.35a.32.32 0 0 1-.1-.1a.23.23 0 0 1-.05-.1a2.77 2.77 0 0 0-2.93-2.34a2.77 2.77 0 0 0-2.84 2.29a.31.31 0 0 1-.07.14a.34.34 0 0 1-.09.08l-.78.38a5.63 5.63 0 0 0-1.91 1.48c-1.57 1.88-1.57 2.86-1.57 5.75a3.84 3.84 0 0 1-.82 1.77a6.569 6.569 0 0 0-.88 1.62a2.79 2.79 0 0 0 .26 2.37a2.24 2.24 0 0 0 1.94.85h2.82c.043.268.117.53.22.78c.198.497.498.947.88 1.32c.37.38.816.677 1.31.87c.46.188.953.287 1.45.29h.16a4 4 0 0 0 2.79-1.16a4 4 0 0 0 .87-1.31c.101-.255.178-.52.23-.79h2.94a2.372 2.372 0 0 0 1-.23a2.4 2.4 0 0 0 .88-.76a2.26 2.26 0 0 0 .4-1.09a2.2 2.2 0 0 0-.14-1.08m-6.72-3.1a.75.75 0 0 1-1.06 1.06l-1.23-1.22l-1.21 1.16a.75.75 0 0 1-1.06-1.06l1.21-1.21l-1.21-1.21a.77.77 0 0 1 0-1.07a.75.75 0 0 1 1.06 0l1.22 1.22l1.22-1.22a.74.74 0 0 1 1.06 0a.75.75 0 0 1 0 1.06l-1.22 1.22z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationBellDownload(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M11.962 17.986h6.81a1.555 1.555 0 0 0 1.512-2.175c-.36-1.088-1.795-2.393-1.795-3.677c0-2.85 0-3.6-1.403-5.276a5.026 5.026 0 0 0-1.654-1.283l-.783-.38a1.088 1.088 0 0 1-.511-.73a2.023 2.023 0 0 0-2.176-1.707a2.024 2.024 0 0 0-2.12 1.707a1.089 1.089 0 0 1-.566.73l-.784.38A5.025 5.025 0 0 0 6.84 6.858c-1.403 1.676-1.403 2.426-1.403 5.276c0 1.284-1.37 2.458-1.73 3.611c-.217.697-.337 2.241 1.48 2.241z"/><path stroke-linejoin="round" d="M15.226 17.986a3.196 3.196 0 0 1-3.264 3.263A3.195 3.195 0 0 1 8.7 17.986"/><path stroke-miterlimit="10" d="M12 13.861v-5"/><path stroke-linejoin="round" d="m9.707 11.756l1.967 1.968a.457.457 0 0 0 .652 0l1.967-1.968"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationBellDownloadFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.003 15.578a7.056 7.056 0 0 0-.87-1.57a4.09 4.09 0 0 1-.89-1.88c0-2.89 0-3.87-1.58-5.76a5.79 5.79 0 0 0-1.9-1.47l-.73-.35a.32.32 0 0 1-.1-.1a.23.23 0 0 1-.05-.1a2.77 2.77 0 0 0-2.93-2.34a2.77 2.77 0 0 0-2.84 2.29a.31.31 0 0 1-.07.14a.34.34 0 0 1-.09.08l-.78.38a5.63 5.63 0 0 0-1.91 1.48c-1.57 1.88-1.57 2.86-1.57 5.75a3.84 3.84 0 0 1-.82 1.77a6.568 6.568 0 0 0-.88 1.62a2.79 2.79 0 0 0 .26 2.37a2.24 2.24 0 0 0 1.94.85h2.82c.043.268.117.53.22.78c.198.497.498.947.88 1.32c.37.38.816.677 1.31.87c.46.188.953.287 1.45.29h.16a3.89 3.89 0 0 0 1.47-.29a3.75 3.75 0 0 0 1.32-.87a4 4 0 0 0 .87-1.31c.101-.255.178-.52.23-.79h2.94a2.372 2.372 0 0 0 1-.23a2.4 2.4 0 0 0 .88-.76c.226-.322.364-.698.4-1.09a2.2 2.2 0 0 0-.14-1.08m-6.18-3.29l-2 2a1.249 1.249 0 0 1-.39.27c-.1.04-.203.067-.31.08h-.32a1.178 1.178 0 0 1-.7-.35l-2-2a.75.75 0 1 1 1.06-1.06l1 1v-3.38a.75.75 0 1 1 1.5 0v3.38l1-1a.75.75 0 1 1 1.06 1.06z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationBellFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.161 16.656a2.26 2.26 0 0 1-.41 1.088a2.27 2.27 0 0 1-1.89 1h-2.94a4.44 4.44 0 0 1-.23.788a3.996 3.996 0 0 1-2.18 2.178c-.495.2-1.026.298-1.56.29h-.08a3.862 3.862 0 0 1-1.44-.29a3.751 3.751 0 0 1-1.32-.87a3.846 3.846 0 0 1-.87-1.308a4.44 4.44 0 0 1-.23-.789h-2.82a2.242 2.242 0 0 1-1.94-.849a2.784 2.784 0 0 1-.26-2.367a6.72 6.72 0 0 1 .88-1.618a3.833 3.833 0 0 0 .82-1.768c0-2.886 0-3.865 1.58-5.743a5.719 5.719 0 0 1 1.9-1.478l.78-.38a.41.41 0 0 0 .1-.09a.31.31 0 0 0 .06-.13a2.995 2.995 0 0 1 1.905-2.142a3.003 3.003 0 0 1 2.835.434a2.716 2.716 0 0 1 1 1.758v.1a.35.35 0 0 0 .11.1l.72.35c.73.35 1.378.85 1.9 1.468c1.58 1.888 1.58 2.867 1.58 5.753c.134.69.44 1.336.89 1.878c.36.481.652 1.009.87 1.568c.164.332.247.698.24 1.069"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationBellMinus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M11.962 17.986h6.81a1.555 1.555 0 0 0 1.512-2.175c-.36-1.088-1.795-2.393-1.795-3.677c0-2.85 0-3.6-1.403-5.276a5.026 5.026 0 0 0-1.654-1.283l-.783-.38a1.088 1.088 0 0 1-.511-.73a2.023 2.023 0 0 0-2.176-1.707a2.024 2.024 0 0 0-2.12 1.707a1.089 1.089 0 0 1-.566.73l-.784.38A5.025 5.025 0 0 0 6.84 6.858c-1.403 1.676-1.403 2.426-1.403 5.276c0 1.284-1.37 2.458-1.73 3.611c-.217.697-.337 2.241 1.48 2.241z"/><path stroke-linejoin="round" d="M15.226 17.986a3.196 3.196 0 0 1-3.264 3.263A3.195 3.195 0 0 1 8.7 17.986"/><path stroke-miterlimit="10" d="M9.5 11.496h5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationBellMinusFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.003 15.578a7.056 7.056 0 0 0-.87-1.57a4.09 4.09 0 0 1-.89-1.88c0-2.89 0-3.87-1.58-5.76a5.79 5.79 0 0 0-1.9-1.47l-.73-.35a.32.32 0 0 1-.1-.1a.23.23 0 0 1-.05-.1a2.77 2.77 0 0 0-2.93-2.34a2.77 2.77 0 0 0-2.84 2.29a.31.31 0 0 1-.06.13a.41.41 0 0 1-.1.09l-.78.38a5.63 5.63 0 0 0-1.91 1.48c-1.57 1.88-1.57 2.86-1.57 5.75a3.84 3.84 0 0 1-.82 1.77a6.568 6.568 0 0 0-.88 1.62a2.79 2.79 0 0 0 .26 2.37a2.24 2.24 0 0 0 1.94.85h2.82c.043.268.117.53.22.78a4 4 0 0 0 .87 1.32c.375.378.825.675 1.32.87c.46.188.953.287 1.45.29h.08a3.92 3.92 0 0 0 1.55-.29a3.75 3.75 0 0 0 1.32-.87a3.8 3.8 0 0 0 .87-1.32c.101-.252.178-.513.23-.78h2.94a2.33 2.33 0 0 0 1-.23a2.4 2.4 0 0 0 .88-.76a2.26 2.26 0 0 0 .4-1.09a2.2 2.2 0 0 0-.14-1.08m-6.5-3.33h-5a.75.75 0 1 1 0-1.5h5a.75.75 0 0 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationBellMuted(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M5.443 17.368c-1.153-.293-1.048-1.509-.87-2.096c.346-1.11 1.687-2.242 1.687-3.469c0-2.724 0-3.479 1.352-5.093a4.988 4.988 0 0 1 1.583-1.237l.744-.366a1.048 1.048 0 0 0 .513-.703a1.928 1.928 0 0 1 2.096-1.645a1.939 1.939 0 0 1 2.044 1.645a1.048 1.048 0 0 0 .555.703l.744.366c.319.159.62.349.902.566m1.135 2.257c.67 1.048.67 2.012.67 4.192c0 1.226 1.394 2.431 1.74 3.532a1.511 1.511 0 0 1-.72 1.94a1.507 1.507 0 0 1-.737.156H8.14"/><path stroke-linejoin="round" d="M15.476 18.105a3.083 3.083 0 0 1-3.145 3.144a3.083 3.083 0 0 1-3.143-3.144"/><path stroke-miterlimit="10" d="M3.542 20.115L19.493 3.33"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationBellMutedFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.96 5.953a.8.8 0 0 1-.27.58L5.5 17.883a.74.74 0 0 1-.89.14c-3.11-1.47 0-4.37.44-6.21c-.25-3.71 1-6 4.1-7.37a3 3 0 0 1 5.87 0c.62.4 1.84.62 1.94 1.51m3.78 10.9c-.17 2.48-3.34 2-5 2c-.61 4.11-7 4.11-7.65 0c-.76.18-1.42-.51-.8-1.17l9.69-9.93c2.14-.66 1.86 2.55 2 4.72c.05 1.45 1.98 2.74 1.76 4.38"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationBellPending(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M10.568 2.975a2.06 2.06 0 0 0-.73 1.27a1 1 0 0 1-.2.42a1 1 0 0 1-.36.3l-.79.38a5.06 5.06 0 0 0-1.65 1.29c-1.4 1.67-1.4 2.42-1.4 5.27c0 1.29-1.37 2.46-1.73 3.62c-.22.69-.34 2.25 1.48 2.25h13.58a1.57 1.57 0 0 0 .77-.16a1.64 1.64 0 0 0 .6-.51c.148-.218.24-.469.27-.73a1.59 1.59 0 0 0-.13-.78c-.36-1.09-1.79-2.39-1.79-3.68v-2.13"/><path stroke-linecap="round" stroke-linejoin="round" d="M15.228 17.775c.003.427-.075.851-.23 1.25a3.348 3.348 0 0 1-.71 1.06a3.19 3.19 0 0 1-2.33.94a3.2 3.2 0 0 1-1.26-.25a3.29 3.29 0 0 1-1.77-1.77a3.2 3.2 0 0 1-.23-1.23"/><circle cx="15.228" cy="5.475" r="2.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationBellPendingFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.37 16.547c-.039.398-.18.779-.41 1.105a2.374 2.374 0 0 1-1.954 1.023H16.01c-.05.27-.126.533-.225.788a4.196 4.196 0 0 1-.89 1.35a4.093 4.093 0 0 1-1.35.9A4.02 4.02 0 0 1 12.03 22h-.112a3.907 3.907 0 0 1-1.545-.317a4.093 4.093 0 0 1-2.22-2.22a3.631 3.631 0 0 1-.225-.788H5.032a2.301 2.301 0 0 1-1.985-.87a2.804 2.804 0 0 1-.256-2.434c.221-.592.524-1.15.9-1.658a3.897 3.897 0 0 0 .83-1.81c0-2.967 0-3.96 1.616-5.883A5.851 5.851 0 0 1 8.08 4.506l.809-.388a.205.205 0 0 0 .081-.082a.246.246 0 0 0 .072-.133a2.865 2.865 0 0 1 1.023-1.73a.767.767 0 0 1 .92 0a.788.788 0 0 1 .297.88s-1.186 3.704 1.023 5.944c2.21 2.24 5.903.133 5.944.113a.747.747 0 0 1 .757 0a.757.757 0 0 1 .379.655v2.179c.138.71.451 1.374.91 1.933c.366.49.665 1.027.89 1.596c.15.337.213.707.184 1.074"/><path fill="currentColor" d="M15.303 8.107a2.782 2.782 0 1 0 0-5.564a2.782 2.782 0 0 0 0 5.564"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationBellPlus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M11.962 17.986h6.81a1.555 1.555 0 0 0 1.512-2.175c-.36-1.088-1.795-2.393-1.795-3.677c0-2.85 0-3.6-1.403-5.276a5.026 5.026 0 0 0-1.654-1.283l-.783-.38a1.088 1.088 0 0 1-.511-.73a2.023 2.023 0 0 0-2.176-1.707a2.024 2.024 0 0 0-2.12 1.707a1.089 1.089 0 0 1-.566.73l-.784.38A5.025 5.025 0 0 0 6.84 6.858c-1.403 1.676-1.403 2.426-1.403 5.276c0 1.284-1.37 2.458-1.73 3.611c-.217.697-.337 2.241 1.48 2.241z"/><path stroke-linejoin="round" d="M15.226 17.986a3.196 3.196 0 0 1-3.264 3.263A3.195 3.195 0 0 1 8.7 17.986"/><path stroke-miterlimit="10" d="M11.995 8.99v5M9.5 11.496h5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationBellPlusFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.005 15.606a7.079 7.079 0 0 0-.87-1.574a4.103 4.103 0 0 1-.89-1.884c0-2.896 0-3.878-1.58-5.772a5.792 5.792 0 0 0-1.9-1.473l-.73-.35a.32.32 0 0 1-.1-.101a.232.232 0 0 1 0-.1a2.777 2.777 0 0 0-1.009-1.745a2.766 2.766 0 0 0-1.921-.6a2.686 2.686 0 0 0-1.85.611a2.737 2.737 0 0 0-1 1.684a.311.311 0 0 1-.06.13a.41.41 0 0 1-.1.09l-.78.38a5.632 5.632 0 0 0-1.91 1.484c-1.57 1.884-1.57 2.866-1.57 5.762a3.852 3.852 0 0 1-.82 1.774a6.586 6.586 0 0 0-.92 1.583a2.802 2.802 0 0 0 .26 2.376a2.24 2.24 0 0 0 1.94.851h2.82c.043.269.117.531.22.782c.2.495.495.944.87 1.323c.375.38.824.676 1.32.872c.46.188.953.287 1.45.29h.08a3.92 3.92 0 0 0 1.55-.29c.497-.194.947-.49 1.32-.872a3.81 3.81 0 0 0 .87-1.323c.1-.253.178-.514.23-.782h2.94c.346-.001.688-.08 1-.23c.351-.178.653-.44.88-.762c.226-.323.364-.7.4-1.092a2.209 2.209 0 0 0-.14-1.042m-6.5-3.337h-1.76v1.743a.752.752 0 0 1-.75.752a.749.749 0 0 1-.75-.752V12.27h-1.74a.75.75 0 0 1-.75-.752a.752.752 0 0 1 .75-.752h1.74V9.002a.752.752 0 0 1 .75-.752a.749.749 0 0 1 .75.752v1.763h1.76a.749.749 0 0 1 .75.752a.752.752 0 0 1-.75.752"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationBellQuestionMark(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M11.962 17.986h6.81a1.555 1.555 0 0 0 1.512-2.175c-.36-1.088-1.795-2.393-1.795-3.677c0-2.85 0-3.6-1.403-5.276a5.026 5.026 0 0 0-1.654-1.283l-.783-.38a1.088 1.088 0 0 1-.511-.73a2.023 2.023 0 0 0-2.176-1.707a2.023 2.023 0 0 0-2.12 1.707a1.089 1.089 0 0 1-.566.73l-.784.38A5.025 5.025 0 0 0 6.84 6.858c-1.403 1.676-1.403 2.426-1.403 5.276c0 1.284-1.37 2.458-1.73 3.611c-.217.697-.337 2.241 1.48 2.241z"/><path stroke-linejoin="round" d="M15.226 17.986a3.196 3.196 0 0 1-3.264 3.263A3.196 3.196 0 0 1 8.7 17.986"/><path stroke-miterlimit="10" d="M10.496 9.932a1.599 1.599 0 0 1 1.742-.906a1.55 1.55 0 0 1 1.137.81a1.345 1.345 0 0 1-.784 1.852a.993.993 0 0 0-.64.897v.37"/><path stroke-linejoin="round" d="M11.922 15h.004"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationBellQuestionMarkFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.003 15.578a7.056 7.056 0 0 0-.87-1.57a4.09 4.09 0 0 1-.89-1.88c0-2.89 0-3.87-1.58-5.76a5.79 5.79 0 0 0-1.9-1.47l-.73-.35a.32.32 0 0 1-.1-.1a.23.23 0 0 1-.05-.1a2.77 2.77 0 0 0-2.93-2.34a2.77 2.77 0 0 0-2.84 2.29a.31.31 0 0 1-.07.14a.34.34 0 0 1-.09.08l-.78.38a5.63 5.63 0 0 0-1.91 1.48c-1.57 1.88-1.57 2.86-1.57 5.75a3.84 3.84 0 0 1-.82 1.77a6.569 6.569 0 0 0-.88 1.62a2.79 2.79 0 0 0 .26 2.37a2.24 2.24 0 0 0 1.94.85h2.82c.043.268.117.53.22.78c.198.497.498.947.88 1.32c.37.38.816.677 1.31.87c.46.188.953.287 1.45.29h.16a4 4 0 0 0 2.79-1.16a4 4 0 0 0 .87-1.31c.101-.255.178-.52.23-.79h2.94a2.372 2.372 0 0 0 1-.23a2.4 2.4 0 0 0 .88-.76a2.26 2.26 0 0 0 .4-1.09a2.2 2.2 0 0 0-.14-1.08m-9.07.17a.75.75 0 1 1 0-1.5a.75.75 0 0 1 0 1.5m2.16-4.52a2.16 2.16 0 0 1-.5.72c-.222.197-.48.35-.76.45a.2.2 0 0 0-.08.08a.19.19 0 0 0-.05.13v.35a.75.75 0 0 1-1.5 0v-.38a1.82 1.82 0 0 1 .32-1a1.75 1.75 0 0 1 .81-.61a.799.799 0 0 0 .24-.14a.9.9 0 0 0 .14-.2a.53.53 0 0 0 0-.25a.741.741 0 0 0-.05-.25a1 1 0 0 0-.23-.24a.72.72 0 0 0-.35-.14a.82.82 0 0 0-.56.08a.8.8 0 0 0-.38.39a.752.752 0 1 1-1.37-.62a2.38 2.38 0 0 1 1.05-1.1a2.34 2.34 0 0 1 1.51-.23c.358.055.7.188 1 .39c.292.214.531.49.7.81a2 2 0 0 1 .21.88c.014.301-.037.601-.15.88"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationBellSnooze(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.568 2.975a2.06 2.06 0 0 0-.73 1.27a1 1 0 0 1-.2.42a1 1 0 0 1-.36.3l-.79.38a5.06 5.06 0 0 0-1.65 1.29c-1.4 1.67-1.4 2.42-1.4 5.27c0 1.29-1.37 2.46-1.73 3.62c-.22.69-.34 2.25 1.48 2.25h13.58a1.57 1.57 0 0 0 .77-.16a1.64 1.64 0 0 0 .6-.51a1.62 1.62 0 0 0 .27-.73a1.589 1.589 0 0 0-.13-.78c-.36-1.09-1.79-2.39-1.79-3.68v-2.13"/><path d="M15.228 17.775c.003.427-.075.851-.23 1.25a3.348 3.348 0 0 1-.71 1.06a3.19 3.19 0 0 1-2.33.94a3.2 3.2 0 0 1-1.26-.25a3.29 3.29 0 0 1-1.77-1.77a3.2 3.2 0 0 1-.23-1.23m1.45-8.85h4.24a.19.19 0 0 1 .14.32l-4.06 4.06a.19.19 0 0 0 .035.289a.19.19 0 0 0 .105.03h4.24m-.75-10.459h2.69a.1.1 0 0 1 .096.119a.1.1 0 0 1-.026.05l-2.59 2.59a.1.1 0 0 0 .015.153a.1.1 0 0 0 .055.018h2.66"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationBellSnoozeTwo(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.65 9.65h4.24a.19.19 0 0 1 .14.32l-4.06 4.06a.19.19 0 0 0 .14.32h4.24"/><path d="M11.962 17.986h6.81a1.555 1.555 0 0 0 1.512-2.175c-.36-1.088-1.795-2.393-1.795-3.677c0-2.85 0-3.6-1.403-5.276a5.026 5.026 0 0 0-1.654-1.283l-.783-.38a1.088 1.088 0 0 1-.511-.73a2.023 2.023 0 0 0-2.176-1.707a2.024 2.024 0 0 0-2.12 1.707a1.089 1.089 0 0 1-.566.73l-.784.38A5.025 5.025 0 0 0 6.84 6.858c-1.403 1.676-1.403 2.426-1.403 5.276c0 1.284-1.37 2.458-1.73 3.611c-.217.697-.337 2.241 1.48 2.241z"/><path d="M15.226 17.986a3.196 3.196 0 0 1-3.264 3.263A3.195 3.195 0 0 1 8.7 17.986"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationBellSnoozeTwoFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.003 15.578a7.056 7.056 0 0 0-.87-1.57a4.09 4.09 0 0 1-.89-1.88c0-2.89 0-3.87-1.58-5.76a5.79 5.79 0 0 0-1.9-1.47l-.73-.35a.32.32 0 0 1-.1-.1a.23.23 0 0 1-.05-.1a2.77 2.77 0 0 0-2.93-2.34a2.77 2.77 0 0 0-2.84 2.29a.31.31 0 0 1-.07.14a.34.34 0 0 1-.09.08l-.78.38a5.63 5.63 0 0 0-1.91 1.48c-1.57 1.88-1.57 2.86-1.57 5.75a3.84 3.84 0 0 1-.82 1.77a6.569 6.569 0 0 0-.88 1.62a2.79 2.79 0 0 0 .26 2.37a2.24 2.24 0 0 0 1.94.85h2.82c.043.268.117.53.22.78c.198.497.498.947.88 1.32c.37.38.816.677 1.31.87c.46.188.953.287 1.45.29h.16a3.89 3.89 0 0 0 1.47-.29a3.75 3.75 0 0 0 1.32-.87a4 4 0 0 0 .87-1.31c.101-.255.178-.52.23-.79h2.94a2.372 2.372 0 0 0 1-.23a2.4 2.4 0 0 0 .88-.76c.226-.322.364-.698.4-1.09a2.2 2.2 0 0 0-.14-1.08m-11.35-6.68h4.24a.87.87 0 0 1 .52.16a.89.89 0 0 1 .34.4a1 1 0 0 1 .07.54a1.089 1.089 0 0 1-.24.48l-3.12 3.12h2.89a.75.75 0 1 1 0 1.5h-4.24a.87.87 0 0 1-.52-.16a.84.84 0 0 1-.34-.4a.91.91 0 0 1-.07-.53a1.07 1.07 0 0 1 .24-.49l3.12-3.12h-2.89a.75.75 0 0 1 0-1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationBellUpload(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M11.962 17.986h6.81a1.555 1.555 0 0 0 1.512-2.175c-.36-1.088-1.795-2.393-1.795-3.677c0-2.85 0-3.6-1.403-5.276a5.026 5.026 0 0 0-1.654-1.283l-.783-.38a1.088 1.088 0 0 1-.511-.73a2.023 2.023 0 0 0-2.176-1.707a2.024 2.024 0 0 0-2.12 1.707a1.089 1.089 0 0 1-.566.73l-.784.38A5.025 5.025 0 0 0 6.84 6.858c-1.403 1.676-1.403 2.426-1.403 5.276c0 1.284-1.37 2.458-1.73 3.611c-.217.697-.337 2.241 1.48 2.241z"/><path stroke-linejoin="round" d="M15.226 17.986a3.196 3.196 0 0 1-3.264 3.263A3.195 3.195 0 0 1 8.7 17.986"/><path stroke-miterlimit="10" d="M12 8.861v5"/><path stroke-linejoin="round" d="M14.293 10.966L12.326 9a.459.459 0 0 0-.652 0l-1.967 1.967"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NotificationBellUploadFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.003 15.578a7.056 7.056 0 0 0-.87-1.57a4.09 4.09 0 0 1-.89-1.88c0-2.89 0-3.87-1.58-5.76a5.79 5.79 0 0 0-1.9-1.47l-.73-.35a.32.32 0 0 1-.1-.1a.23.23 0 0 1-.05-.1a2.77 2.77 0 0 0-2.93-2.34a2.77 2.77 0 0 0-2.84 2.29a.31.31 0 0 1-.07.14a.34.34 0 0 1-.09.08l-.78.38a5.63 5.63 0 0 0-1.91 1.48c-1.57 1.88-1.57 2.86-1.57 5.75a3.84 3.84 0 0 1-.82 1.77a6.568 6.568 0 0 0-.88 1.62a2.79 2.79 0 0 0 .26 2.37a2.24 2.24 0 0 0 1.94.85h2.82c.043.268.117.53.22.78c.198.497.498.947.88 1.32c.37.38.816.677 1.31.87c.46.188.953.287 1.45.29h.16a4 4 0 0 0 2.79-1.16a4 4 0 0 0 .87-1.31c.101-.255.178-.52.23-.79h2.94a2.372 2.372 0 0 0 1-.23a2.4 2.4 0 0 0 .88-.76a2.26 2.26 0 0 0 .4-1.09a2.2 2.2 0 0 0-.14-1.08m-6.18-4.1a.75.75 0 0 1-.53.22a.79.79 0 0 1-.53-.22l-1-1v3.38a.75.75 0 0 1-1.5 0v-3.38l-1 1a.75.75 0 0 1-1.06 0a.74.74 0 0 1 0-1.06l2-2a1.31 1.31 0 0 1 .4-.26a1 1 0 0 1 .44-.1a1.2 1.2 0 0 1 .45.09c.147.07.282.16.4.27l2 2a.74.74 0 0 1-.07 1.06"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Notion(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.432 12.904v5.994c-.017.697-.355 1.082-1.053 1.14c-.801.066-1.606.108-2.41.156c-1.62.1-3.241.198-4.861.294c-1.435.085-2.87.16-4.303.257c-.634.041-1.133-.159-1.484-.703c-.104-.161-.237-.306-.355-.46c-.6-.774-1.195-1.55-1.798-2.32a2.708 2.708 0 0 1-.607-1.763c.02-3.41.008-6.818.01-10.227c0-.129.01-.257.03-.384c.092-.5.402-.81.908-.858c.828-.078 1.669-.132 2.504-.195l3.99-.294c1.278-.093 2.556-.182 3.834-.283a2.578 2.578 0 0 1 1.786.53c1.126.817 2.27 1.607 3.412 2.404a.88.88 0 0 1 .403.797c-.011 1.971-.006 3.943-.006 5.915m-1.162.077V7.877c0-.492-.173-.662-.65-.634l-4.15.236l-5.09.298c-.598.035-1.196.07-1.794.11c-.317.022-.461.162-.508.473c-.012.09-.017.18-.015.27v9.86c0 .767.204.959.974.914c.77-.045 1.539-.095 2.317-.14l8.239-.473c.442-.025.643-.22.678-.659c.006-.077 0-.154 0-.236zm-1.247-6.916l.027-.051a1.635 1.635 0 0 0-.2-.205c-.546-.399-1.098-.79-1.643-1.183c-.357-.273-.8-.41-1.249-.385c-.894.054-1.786.125-2.68.19l-5.744.417c-.412.03-.829.062-1.231.11a.298.298 0 0 0-.21.161c-.017.049.064.159.127.21c.45.363.908.72 1.362 1.078c.245.204.558.31.876.297c.155-.007.31 0 .464-.014l4.168-.256l5.538-.326c.132-.015.263-.03.395-.043"/><path fill="currentColor" d="M15.527 14.688v-4.66c-.252-.03-.493-.065-.736-.083c-.133-.01-.164-.07-.15-.187a.563.563 0 0 1 .492-.5c.803-.06 1.606-.101 2.405-.149c.119.397 0 .592-.38.667c-.419.087-.419.087-.419.512v6.845c0 .159-.042.256-.192.303a8.236 8.236 0 0 1-.807.255c-.439.092-.828-.03-1.089-.409a60.377 60.377 0 0 1-1.065-1.588c-.804-1.244-1.598-2.493-2.397-3.74c-.016-.026-.038-.051-.083-.113c-.01.086-.019.135-.019.186v4.66c0 .137.044.193.176.218c.296.058.591.129.88.194a.6.6 0 0 1-.54.576a7.14 7.14 0 0 1-.731.06c-.592.036-1.184.064-1.775.106c-.173.012-.208-.06-.178-.201a.406.406 0 0 1 .315-.355c.186-.045.37-.104.554-.156v-6.69l-.829-.074a.592.592 0 0 1 .533-.772c.725-.058 1.455-.08 2.178-.162c.309-.034.453.08.608.319c1.023 1.576 2.06 3.146 3.093 4.717z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PackageBox(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m4.988 17.452l5.75 3.448a2.45 2.45 0 0 0 2.524 0l5.75-3.448c.366-.219.67-.53.88-.901c.205-.37.315-.786.318-1.21V8.278a2.462 2.462 0 0 0-1.198-2.122l-5.75-3.065a2.514 2.514 0 0 0-2.524 0l-5.75 3.065A2.461 2.461 0 0 0 3.79 8.277v7.065c.003.423.113.839.318 1.209c.21.371.514.682.88.901M19.881 7.078L12 11.81"/><path d="M4.119 7.078L12 11.81v9.44m4.38-8.316V9.179L8.066 4.522"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PackageBoxFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.15 12.335v9.18a.552.552 0 0 1-.15-.08l-6.51-3.91a1.88 1.88 0 0 1-.7-.7a1.91 1.91 0 0 1-.25-1v-8.07zm9.31-4.58v8.1a2.09 2.09 0 0 1-.27.95a1.74 1.74 0 0 1-.69.71l-6.51 3.91l-.14.07v-9.17l3.26-2v2.77a.85.85 0 1 0 1.7 0v-3.74zm-5.18 1.15l-3.28 2l-7.66-4.6l.11-.07l3.06-1.63zm4.37-2.62l-2.71 1.62l-7.64-4.28l1.66-.87a2 2 0 0 1 1-.27a2.08 2.08 0 0 1 1 .28l6.47 3.46a.49.49 0 0 1 .22.06"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pause(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.5 4.5h-3a1 1 0 0 0-1 1v13a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1v-13a1 1 0 0 0-1-1m10 0h-3a1 1 0 0 0-1 1v13a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1v-13a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.25 5.5v13a1.75 1.75 0 0 1-1.75 1.75h-3a1.75 1.75 0 0 1-1.75-1.75v-13A1.76 1.76 0 0 1 5.5 3.75h3a1.75 1.75 0 0 1 1.75 1.75m10 0v13a1.75 1.75 0 0 1-1.75 1.75h-3a1.75 1.75 0 0 1-1.75-1.75v-13a1.76 1.76 0 0 1 1.75-1.75h3a1.75 1.75 0 0 1 1.75 1.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M10.074 7.723H8.223a.5.5 0 0 0-.5.5v7.554a.5.5 0 0 0 .5.5h1.851a.5.5 0 0 0 .5-.5V8.223a.5.5 0 0 0-.5-.5m5.703 0h-1.851a.5.5 0 0 0-.5.5v7.554a.5.5 0 0 0 .5.5h1.851a.5.5 0 0 0 .5-.5V8.223a.5.5 0 0 0-.5-.5"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PauseSquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 2h-6.5A6.76 6.76 0 0 0 2 8.75v6.5A6.76 6.76 0 0 0 8.75 22h6.5A6.76 6.76 0 0 0 22 15.25v-6.5A6.76 6.76 0 0 0 15.25 2m-3.93 13.09a1.24 1.24 0 0 1-1.23 1.23H8.58a1.24 1.24 0 0 1-1.23-1.23V8.91a1.24 1.24 0 0 1 1.23-1.23h1.51a1.24 1.24 0 0 1 1.23 1.23zm5.32 0a1.24 1.24 0 0 1-1.23 1.23H13.9a1.24 1.24 0 0 1-1.23-1.23V8.91a1.24 1.24 0 0 1 1.23-1.23h1.51a1.24 1.24 0 0 1 1.23 1.23z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paypal(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.026 6.117c-.213 5.113-3.887 6.804-8.485 6.425c-.758 0-.894.778-.972 1.37c-.31 1.556-.398 3.305-.816 4.773h-3.45a.476.476 0 0 1-.496-.467c.632-4.14 1.312-8.29 1.944-12.44c.145-.924.3-1.847.437-2.77a.826.826 0 0 1 .408-.613a31.305 31.305 0 0 1 4.277-.117c2.819-.116 7.25-.048 7.153 3.84"/><path fill="currentColor" d="M18.678 8.061c2.497 1.41 1.526 5.356-.175 7.105a6.356 6.356 0 0 1-4.724 1.439a.784.784 0 0 0-.768.826l-.583 3.625a.728.728 0 0 1-.33.554a18.641 18.641 0 0 1-3.383.107a.417.417 0 0 1-.428-.486c.078-.534.166-1.03.253-1.506a43.817 43.817 0 0 1 1.03-5.832a.836.836 0 0 1 .652-.292c.972 0 1.876 0 2.809-.058a5.987 5.987 0 0 0 5.598-5.472z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pen(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m2.777 19.783l.607-4.162c.037-.272.161-.525.355-.72L15.5 3.124a1.265 1.265 0 0 1 1.19-.341a6.21 6.21 0 0 1 2.832 1.694a6.21 6.21 0 0 1 1.682 2.846a1.265 1.265 0 0 1-.341 1.19L9.089 20.275a1.265 1.265 0 0 1-.721.354l-4.161.607a1.264 1.264 0 0 1-1.43-1.454M13.275 5.364l5.363 5.363"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PenFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.911 7.142a7.45 7.45 0 0 0-5.07-5.07a2 2 0 0 0-1 0a2 2 0 0 0-.87.52l-2.23 2.24l-9.54 9.53a2.09 2.09 0 0 0-.56 1.14l-.61 4.16a2.09 2.09 0 0 0 .07.93c.102.3.27.573.49.8c.224.222.494.392.79.5c.207.068.423.101.64.1h.29l4.16-.61a2 2 0 0 0 1.15-.56l9.55-9.55l2.22-2.21a1.999 1.999 0 0 0 .54-1.92zm-1.46.63a.52.52 0 0 1-.14.22l-1.69 1.69l-4.32-4.31l1.71-1.71a.581.581 0 0 1 .23-.13h.23a5.94 5.94 0 0 1 4 4a.57.57 0 0 1-.02.24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Phone(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10" stroke-width="1.5" d="M12.735 20.191a15.2 15.2 0 0 1-.92-.447a19.012 19.012 0 0 1-4.1-3.12A18.002 18.002 0 0 1 3.88 11.42a11.259 11.259 0 0 1-1.022-3.325a5.925 5.925 0 0 1 .37-3.465c.289-.47.637-.9 1.035-1.279a1.801 1.801 0 0 1 1.278-.601c.505.076.962.34 1.278.742c.69.767 1.43 1.457 2.159 2.186c.287.246.466.595.498.972c-.012.317-.134.62-.345.857c-.242.307-.536.588-.817.882a1.535 1.535 0 0 0-.46 1.279a3.67 3.67 0 0 0 .881 1.457c.486.665.971 1.28 1.52 1.931a13.573 13.573 0 0 0 3.463 2.865a1.277 1.277 0 0 0 1.278.153a4.05 4.05 0 0 0 1.137-.946c.275-.335.669-.55 1.099-.601c.383.02.744.184 1.01.46c.344.294.638.64.958.959c.319.32.575.55.843.844c.321.283.624.586.907.908c.22.284.324.64.294.997a2.1 2.1 0 0 1-.703 1.087a4.78 4.78 0 0 1-3.756 1.458a10.673 10.673 0 0 1-4.05-1.049Z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneCall(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10" stroke-width="1.5" d="M12.531 20.217c-.3-.137-.599-.274-.898-.437a18.564 18.564 0 0 1-4.005-3.044a17.568 17.568 0 0 1-3.743-5.078a10.979 10.979 0 0 1-.998-3.243a5.776 5.776 0 0 1 .362-3.381a6.238 6.238 0 0 1 1.01-1.248A1.759 1.759 0 0 1 5.508 3.2c.493.074.94.332 1.248.723c.674.749 1.397 1.423 2.108 2.134c.28.24.455.58.487.948c-.012.31-.131.605-.337.836c-.237.3-.524.574-.798.86a1.497 1.497 0 0 0-.45 1.248c.172.535.467 1.023.861 1.423a34.17 34.17 0 0 0 1.485 1.883a13.248 13.248 0 0 0 3.38 2.795a1.25 1.25 0 0 0 1.249.15c.423-.237.8-.55 1.11-.923a1.635 1.635 0 0 1 1.073-.587c.374.02.726.18.986.45c.336.286.623.623.935.935c.312.312.562.537.824.823a10 10 0 0 1 .885.886c.214.277.317.625.287.973a2.045 2.045 0 0 1-.686 1.06a4.666 4.666 0 0 1-3.668 1.423a10.432 10.432 0 0 1-3.955-1.023Zm5.22-9.019a5.059 5.059 0 0 0-5.059-5.059m8.527 4.705a8.094 8.094 0 0 0-8.095-8.094"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneCallFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.33 19.035a2.57 2.57 0 0 1-.884 1.432a5.251 5.251 0 0 1-3.738 1.564h-.325a10.973 10.973 0 0 1-4.205-1.087h-.01c-.305-.142-.62-.284-.925-.457a19.127 19.127 0 0 1-4.185-3.18a18.193 18.193 0 0 1-3.9-5.292A11.692 11.692 0 0 1 2.14 8.572a6.38 6.38 0 0 1 .407-3.708a6.827 6.827 0 0 1 1.148-1.432A2.194 2.194 0 0 1 5.29 2.69a2.51 2.51 0 0 1 1.687.935c.457.497 1.015 1.015 1.473 1.493l.63.62c.37.328.599.786.64 1.28c0 .453-.167.89-.468 1.229a9.141 9.141 0 0 1-.62.68l-.203.213c-.118.11-.208.246-.264.397c-.05.147-.07.302-.06.457c.161.431.414.823.74 1.148c.509.69 1.017 1.29 1.535 1.94a12.9 12.9 0 0 0 3.29 2.733c.127.093.273.155.428.182c.134.01.27-.01.396-.06c.355-.209.67-.477.934-.793a2.174 2.174 0 0 1 1.422-.782a2.032 2.032 0 0 1 1.423.61c.203.172.426.406.64.63l.304.314l.315.305l.539.548c.321.285.623.59.904.915c.282.39.409.872.355 1.35m-3.646-6.958a.772.772 0 0 1-.762-.762a4.37 4.37 0 0 0-4.378-4.378a.762.762 0 0 1 0-1.524a5.893 5.893 0 0 1 5.902 5.902a.762.762 0 0 1-.762.762"/><path fill="currentColor" d="M21.209 11.72a.772.772 0 0 1-.762-.761a7.455 7.455 0 0 0-7.456-7.467a.762.762 0 1 1 0-1.523a8.98 8.98 0 0 1 8.98 8.99a.762.762 0 0 1-.762.762"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneCancel(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10" stroke-width="1.5" d="M7.254 16.205a17.147 17.147 0 0 1-3.424-4.74a11.244 11.244 0 0 1-.971-3.322a5.916 5.916 0 0 1 .37-3.463a6.389 6.389 0 0 1 1.035-1.278a1.802 1.802 0 0 1 1.278-.6c.505.075.962.34 1.278.741C7.51 4.31 8.25 5 8.979 5.728c.287.246.466.594.498.97c-.012.318-.134.62-.345.857c-.242.307-.536.588-.817.882a1.533 1.533 0 0 0-.46 1.277c.175.548.477 1.047.881 1.457c.486.664.971 1.278 1.52 1.93l.09.089m1.866 1.648c.472.387.976.733 1.507 1.035a1.277 1.277 0 0 0 1.278.153c.434-.243.82-.563 1.137-.945c.275-.335.669-.55 1.099-.601c.383.02.744.184 1.01.46c.344.294.638.639.958.958c.319.32.575.55.843.843c.321.283.624.586.907.908c.22.283.324.639.294.996c-.114.427-.36.807-.703 1.086a4.778 4.778 0 0 1-3.756 1.457a10.682 10.682 0 0 1-4.012-.958c-.307-.14-.614-.281-.92-.448a16.813 16.813 0 0 1-2.696-1.852m-5.264 1.622L18.012 5.72"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneCancelFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.16 12.267l-.74.73l-2.478 2.416l-.73.72A17.999 17.999 0 0 1 3.13 11.68A12.222 12.222 0 0 1 2.1 8.104a6.568 6.568 0 0 1 .421-3.834a6.866 6.866 0 0 1 1.182-1.48a2.2 2.2 0 0 1 .73-.535c.287-.136.598-.213.915-.226a2.6 2.6 0 0 1 1.737.956c.473.535 1.028 1.028 1.532 1.573l.647.637a1.758 1.758 0 0 1 .175 2.58a7.605 7.605 0 0 1-.627.689l-.226.226c-.121.12-.216.263-.278.422a1.1 1.1 0 0 0 0 .483c.166.447.425.854.761 1.192c.319.524.689.997 1.09 1.48"/><path fill="currentColor" d="M18.681 6.08L4.126 20.294a.812.812 0 0 1-.545.226a.76.76 0 0 1-.545-.236a.77.77 0 0 1 0-1.09l3.084-3.083l.73-.72l2.477-2.415l.74-.73l7.442-7.288a.77.77 0 0 1 1.09 0a.773.773 0 0 1 .082 1.12m3.331 12.799a2.682 2.682 0 0 1-.915 1.47a5.21 5.21 0 0 1-1.902 1.254a5.324 5.324 0 0 1-1.963.37h-.35a11.605 11.605 0 0 1-4.306-1.028c-.32-.154-.648-.298-1.028-.483a16.873 16.873 0 0 1-2.056-1.337a15.091 15.091 0 0 1-.792-.627a.503.503 0 0 1-.082-.688l.103-.113l.195-.196l2.426-2.456l.411-.422l.082-.082a.524.524 0 0 1 .638 0c.28.233.574.446.884.637c.19.14.39.267.596.38c.13.1.281.167.442.196c.14.011.28-.01.411-.062a3.55 3.55 0 0 0 1.028-.832a2.242 2.242 0 0 1 1.46-.802a2.159 2.159 0 0 1 1.47.627c.216.185.442.421.668.658l.308.329l.319.308a10 10 0 0 1 .565.576c.337.285.65.598.936.935c.321.388.483.884.452 1.388"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneCross(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10" stroke-width="1.5" d="M12.744 20.255c-.313-.145-.615-.29-.906-.447a19.33 19.33 0 0 1-4.113-3.132a18.008 18.008 0 0 1-3.89-5.201c-.5-1.06-.83-2.193-.973-3.356a5.91 5.91 0 0 1 .37-3.467c.287-.48.637-.92 1.039-1.31a1.755 1.755 0 0 1 1.263-.592a1.888 1.888 0 0 1 1.285.738c.705.76 1.431 1.466 2.169 2.237c.283.25.463.597.503.973c-.029.299-.15.58-.347.806c-.257.313-.547.593-.827.895a1.532 1.532 0 0 0-.458 1.23c.178.546.48 1.044.883 1.454c.492.66.984 1.32 1.52 1.935a13.828 13.828 0 0 0 3.476 2.864a1.297 1.297 0 0 0 1.241.156a4.004 4.004 0 0 0 1.118-.95a1.72 1.72 0 0 1 1.118-.604c.39.014.758.178 1.028.458c.336.302.638.649.962.973c.324.325.57.548.838.84c.319.288.62.594.906.916a1.332 1.332 0 0 1 .29.996c-.1.44-.344.833-.693 1.118a4.793 4.793 0 0 1-3.767 1.454a10.702 10.702 0 0 1-4.035-.984ZM17.182 5.68l-3.633 3.636m0-3.636l3.633 3.636"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneCrossFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.024 18.905a2.634 2.634 0 0 1-.906 1.482a5.32 5.32 0 0 1-1.914 1.265a5.524 5.524 0 0 1-1.955.36h-.36a11.414 11.414 0 0 1-4.343-1.028c-.329-.155-.648-.31-.957-.474a20.377 20.377 0 0 1-4.353-3.313a18.894 18.894 0 0 1-4.116-5.485A11.804 11.804 0 0 1 2.091 8.09a6.452 6.452 0 0 1 .412-3.828a7.49 7.49 0 0 1 1.183-1.523a2.49 2.49 0 0 1 .73-.536c.288-.133.6-.206.917-.216a2.47 2.47 0 0 1 1.029.31c.286.17.537.394.74.658c.381.411.814.843 1.246 1.286l.947.967c.372.34.606.806.658 1.307v.103A1.945 1.945 0 0 1 9.47 7.73c-.175.217-.37.422-.576.638l-.278.299a1.1 1.1 0 0 0-.35.813c.164.463.427.885.772 1.235c.546.72 1.03 1.4 1.585 2.058a13.655 13.655 0 0 0 3.447 2.84a.854.854 0 0 0 .834.123c.36-.215.677-.494.936-.823c.374-.452.91-.74 1.492-.803a2.059 2.059 0 0 1 1.482.628c.206.175.422.402.638.628c.216.226.227.247.35.36l.32.32c.195.184.37.36.555.565c.332.295.645.611.936.947c.166.2.291.431.37.679c.047.22.06.445.042.669"/><path fill="currentColor" d="M17.877 8.717a.763.763 0 0 1 0 1.091a.792.792 0 0 1-.546.226a.813.813 0 0 1-.545-.226l-1.317-1.327l-1.328 1.327a.772.772 0 0 1-1.09-1.09l1.327-1.328l-1.328-1.328a.772.772 0 0 1 1.091-1.09l1.328 1.327l1.317-1.327a.772.772 0 0 1 1.09 0a.763.763 0 0 1 0 1.09L16.55 7.39z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.963 18.855a2.736 2.736 0 0 1-.898 1.47a5.36 5.36 0 0 1-3.848 1.602h-.358a11.363 11.363 0 0 1-4.287-1.082c-.326-.153-.643-.296-1.02-.47A19.792 19.792 0 0 1 7.253 17.1a18.58 18.58 0 0 1-4.012-5.451A11.902 11.902 0 0 1 2.15 8.106a6.523 6.523 0 0 1 .418-3.808a6.88 6.88 0 0 1 1.174-1.48a2.304 2.304 0 0 1 1.634-.745a2.542 2.542 0 0 1 1.725.95c.47.52 1.02 1.02 1.52 1.55l.644.634c.38.333.615.802.653 1.306c.001.464-.17.911-.48 1.256a9.287 9.287 0 0 1-.622.694l-.215.225a1.147 1.147 0 0 0-.286.418c-.052.154-.07.318-.05.48c.164.444.421.848.755 1.184c.52.704 1.02 1.317 1.582 2.042a13.27 13.27 0 0 0 3.4 2.807c.123.1.27.167.428.194c.14.021.281 0 .408-.062a3.52 3.52 0 0 0 1.021-.826c.36-.444.882-.726 1.45-.787a2.042 2.042 0 0 1 1.46.623c.233.201.455.416.663.643l.306.327l.317.306c.193.194.377.368.56.572c.333.286.644.596.93.929c.293.374.441.842.418 1.317"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneIncoming(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="M12.744 20.255a17.03 17.03 0 0 1-.906-.447a19.324 19.324 0 0 1-4.113-3.132a18.007 18.007 0 0 1-3.89-5.201c-.5-1.06-.83-2.193-.973-3.356a5.91 5.91 0 0 1 .37-3.467c.287-.48.637-.92 1.039-1.31a1.755 1.755 0 0 1 1.263-.592a1.889 1.889 0 0 1 1.285.738c.704.76 1.431 1.466 2.169 2.237c.283.25.463.598.503.973c-.029.299-.15.58-.347.806c-.257.313-.547.593-.827.895a1.532 1.532 0 0 0-.458 1.23c.178.546.48 1.044.883 1.454c.492.66.984 1.32 1.52 1.935a13.816 13.816 0 0 0 3.477 2.864a1.296 1.296 0 0 0 1.24.156a4.003 4.003 0 0 0 1.118-.95a1.72 1.72 0 0 1 1.118-.604c.39.014.758.178 1.028.458c.336.302.638.649.962.973c.324.325.57.548.838.84c.319.288.62.594.906.916a1.33 1.33 0 0 1 .29.996c-.1.44-.344.833-.693 1.118a4.798 4.798 0 0 1-3.767 1.454a10.703 10.703 0 0 1-4.035-.984Zm.38-9.395l5.589-5.593"/><path stroke-linejoin="round" d="M13.124 7.493v3.366h3.353"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneIncomingFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.99 18.89a2.623 2.623 0 0 1-.893 1.478c-.541.55-1.192.98-1.91 1.263a5.577 5.577 0 0 1-1.961.359h-.36a11.295 11.295 0 0 1-4.322-1.027a20.596 20.596 0 0 1-.955-.472a19.824 19.824 0 0 1-4.343-3.306a18.974 18.974 0 0 1-4.106-5.472a11.777 11.777 0 0 1-1.027-3.614a6.52 6.52 0 0 1 .42-3.82A6.756 6.756 0 0 1 3.716 2.77a2.29 2.29 0 0 1 .718-.544c.288-.131.598-.204.914-.216c.362.026.713.135 1.027.318c.286.162.534.383.729.647c.39.411.821.842 1.252 1.284c.431.441.627.636.935.965c.371.34.605.803.657 1.304v.102a2.054 2.054 0 0 1-.473 1.11a8.42 8.42 0 0 1-.585.636l-.277.298a1.16 1.16 0 0 0-.277.41a1.027 1.027 0 0 0-.062.4c.16.462.419.883.76 1.233c.534.708 1.027 1.396 1.581 2.053a13.626 13.626 0 0 0 3.44 2.834a.83.83 0 0 0 .43.174a.76.76 0 0 0 .401-.051c.36-.217.68-.495.945-.821c.368-.451.9-.739 1.478-.801a2.053 2.053 0 0 1 1.489.626c.221.196.43.405.626.626l.35.36l.318.318c.195.185.37.36.564.565c.195.205.627.605.934.944c.165.2.288.43.36.678c.054.218.068.444.041.667"/><path fill="currentColor" d="m19.434 5.655l-4.426 4.425h1.582a.77.77 0 1 1 0 1.54h-3.47a.739.739 0 0 1-.299-.061a.811.811 0 0 1-.472-.709V7.39a.77.77 0 1 1 1.54 0v1.602l4.425-4.425a.77.77 0 0 1 1.089 0a.77.77 0 0 1 .03 1.088"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneMinus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10" stroke-width="1.5" d="M12.744 20.255c-.313-.145-.615-.29-.906-.447a19.327 19.327 0 0 1-4.113-3.132a18.007 18.007 0 0 1-3.89-5.201c-.5-1.06-.83-2.193-.973-3.356a5.91 5.91 0 0 1 .37-3.467c.287-.48.637-.92 1.039-1.31a1.755 1.755 0 0 1 1.263-.592a1.889 1.889 0 0 1 1.285.738c.705.76 1.431 1.466 2.169 2.237c.283.25.463.598.503.973c-.029.299-.15.58-.347.806c-.257.313-.547.593-.827.895a1.532 1.532 0 0 0-.458 1.23c.178.546.48 1.044.883 1.454c.492.66.984 1.32 1.52 1.935a13.816 13.816 0 0 0 3.477 2.864a1.296 1.296 0 0 0 1.24.156a4.003 4.003 0 0 0 1.118-.95a1.72 1.72 0 0 1 1.118-.604c.39.014.758.178 1.028.458c.336.302.638.649.962.973c.324.325.57.548.838.84c.319.288.62.594.906.916a1.33 1.33 0 0 1 .29.996c-.1.44-.344.833-.693 1.118a4.798 4.798 0 0 1-3.767 1.454a10.703 10.703 0 0 1-4.035-.984ZM12.8 7.504h5.13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneMinusFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.958 18.858a2.617 2.617 0 0 1-.9 1.471a5.375 5.375 0 0 1-1.9 1.257a5.508 5.508 0 0 1-1.952.358h-.347a11.308 11.308 0 0 1-4.313-1.022a20.607 20.607 0 0 1-.95-.47a20.239 20.239 0 0 1-4.323-3.29a18.855 18.855 0 0 1-4.088-5.447a11.445 11.445 0 0 1-1.021-3.597a6.418 6.418 0 0 1 .408-3.802a7.03 7.03 0 0 1 1.175-1.502a1.983 1.983 0 0 1 2.637-.44c.286.164.536.383.736.644c.378.409.807.838 1.236 1.278l.94.96c.372.352.598.83.634 1.339v.102c-.038.41-.207.796-.48 1.104a10.91 10.91 0 0 1-.573.633l-.286.297c-.12.113-.211.253-.265.408a1.093 1.093 0 0 0-.072.399c.159.462.42.881.767 1.226l.071.092c.48.644 1.022 1.308 1.502 1.911a13.562 13.562 0 0 0 3.424 2.82a.79.79 0 0 0 .419.174c.135.02.272.002.398-.05a3.32 3.32 0 0 0 .94-.818c.371-.45.903-.735 1.482-.797a2.046 2.046 0 0 1 1.472.623c.204.174.419.399.633.624c.215.224.225.245.348.357l.316.317c.195.184.368.358.552.562c.328.295.638.61.93.94c.164.2.29.429.368.675c.06.216.087.44.082.664M18.83 7.423a.766.766 0 0 1-.766.766h-5.242a.766.766 0 1 1 0-1.533h5.283a.777.777 0 0 1 .726.767"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneMissedCall(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="m6.889 12.731l.777-.269c1.332-.37 2.71-.557 4.092-.557a14.46 14.46 0 0 1 5.139.746a8.903 8.903 0 0 1 2.45 1.394a4.73 4.73 0 0 1 1.752 2.171c.104.434.155.879.15 1.325a1.456 1.456 0 0 1-.38 1.055a1.6 1.6 0 0 1-1.154.309h-2.46a1.235 1.235 0 0 1-.836-.26a1.116 1.116 0 0 1-.28-.676v-.996a1.225 1.225 0 0 0-.437-.996a2.986 2.986 0 0 0-1.325-.318c-.647-.1-1.294-.2-1.991-.24a10.87 10.87 0 0 0-3.585.35a.997.997 0 0 0-.787.617a3.35 3.35 0 0 0-.11 1.175a1.305 1.305 0 0 1-.278.995a1.156 1.156 0 0 1-.837.319H4.708a8.845 8.845 0 0 1-.996 0a1.135 1.135 0 0 1-.736-.398a1.703 1.703 0 0 1-.22-.996a3.853 3.853 0 0 1 1.305-2.988a8.495 8.495 0 0 1 2.828-1.762Z"/><path stroke-linejoin="round" d="m16.986 5.084l-3.923 3.923a1.504 1.504 0 0 1-2.111 0L7.028 5.084"/><path stroke-linejoin="round" d="M7.028 9.067V5.083h3.984"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneMissedCallFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.981 17.81c.017.255-.022.511-.113.75c-.09.249-.23.476-.411.669a2.232 2.232 0 0 1-1.368.504h-2.726a1.81 1.81 0 0 1-1.183-.38a1.718 1.718 0 0 1-.483-1.03v-1.028a.647.647 0 0 0-.062-.35a.545.545 0 0 0-.164-.236a2.397 2.397 0 0 0-1.029-.247c-.75-.113-1.388-.206-2.057-.247a10.727 10.727 0 0 0-3.528.34a.617.617 0 0 0-.298.123a.545.545 0 0 0-.165.206c-.077.32-.101.65-.072.977c.03.242.01.487-.061.72a1.79 1.79 0 0 1-.34.669a1.708 1.708 0 0 1-1.275.524H4.517c-.343.02-.686.02-1.028 0a1.728 1.728 0 0 1-1.122-.597a2.242 2.242 0 0 1-.35-1.398a4.412 4.412 0 0 1 .37-1.852a4.515 4.515 0 0 1 1.143-1.563a9.165 9.165 0 0 1 3.085-1.913l.802-.278a16.28 16.28 0 0 1 4.372-.596a15.17 15.17 0 0 1 5.45.792a9.69 9.69 0 0 1 2.675 1.522a5.225 5.225 0 0 1 1.964 2.437c.095.488.13.986.103 1.482M6.296 4.439l.113-.082a.452.452 0 0 0-.113.092zm.699-.216a1.026 1.026 0 0 0-.308 0z"/><path fill="currentColor" d="M17.63 5.54L13.6 9.57c-.437.43-1.024.67-1.636.669a2.283 2.283 0 0 1-1.625-.669L7.613 6.856v2.232a.771.771 0 0 1-1.543 0V4.974a.782.782 0 0 1 .34-.638a.679.679 0 0 1 .277-.123h4.259a.771.771 0 0 1 0 1.542H8.714l2.715 2.716a.792.792 0 0 0 1.028 0l4.032-4.032a.771.771 0 0 1 1.09 0a.792.792 0 0 1 .052 1.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneOutgoing(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="M12.744 20.255c-.313-.145-.615-.29-.906-.447a19.328 19.328 0 0 1-4.114-3.132a18.008 18.008 0 0 1-3.891-5.201c-.5-1.06-.829-2.193-.973-3.356a5.908 5.908 0 0 1 .37-3.467c.288-.48.637-.92 1.04-1.31a1.756 1.756 0 0 1 1.262-.592a1.89 1.89 0 0 1 1.286.738c.705.76 1.431 1.466 2.17 2.237c.283.25.462.597.502.973c-.028.299-.15.58-.346.806c-.257.313-.548.593-.828.895a1.533 1.533 0 0 0-.458 1.23c.178.546.48 1.044.883 1.454c.492.66.984 1.32 1.52 1.935a13.82 13.82 0 0 0 3.478 2.864a1.296 1.296 0 0 0 1.241.156a4.003 4.003 0 0 0 1.118-.95a1.721 1.721 0 0 1 1.118-.604c.39.014.758.178 1.029.458c.335.302.637.649.961.973c.325.325.57.548.839.84c.318.288.62.594.905.916a1.33 1.33 0 0 1 .291.996c-.1.44-.344.833-.693 1.118a4.799 4.799 0 0 1-3.768 1.454a10.707 10.707 0 0 1-4.036-.984Zm5.97-14.988l-5.59 5.592"/><path stroke-linejoin="round" d="M18.714 8.611V5.255H15.36"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneOutgoingFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.935 18.844a2.58 2.58 0 0 1-.897 1.468a5.278 5.278 0 0 1-1.897 1.255a5.496 5.496 0 0 1-1.948.357h-.357a11.31 11.31 0 0 1-4.304-1.02a20.428 20.428 0 0 1-.949-.47a20.19 20.19 0 0 1-4.314-3.283a18.9 18.9 0 0 1-4.08-5.437a11.699 11.699 0 0 1-1.02-3.59a6.487 6.487 0 0 1 .419-3.794A7.018 7.018 0 0 1 3.76 2.831c.2-.226.442-.41.714-.54c.289-.131.6-.204.918-.215c.358.026.706.13 1.02.306c.283.17.532.392.734.653c.377.408.805.836 1.234 1.275l.938.959c.364.34.595.8.653 1.295v.102a1.928 1.928 0 0 1-.48 1.101c-.173.215-.377.419-.57.633l-.286.296c-.12.114-.214.253-.276.408a1.295 1.295 0 0 0-.06.397c.157.461.42.88.764 1.224l.061.092c.49.643 1.02 1.305 1.51 1.907a13.536 13.536 0 0 0 3.417 2.815c.12.097.264.157.418.174a.74.74 0 0 0 .397-.051c.36-.214.678-.49.939-.817c.37-.448.9-.733 1.479-.795a2.039 2.039 0 0 1 1.468.622c.204.174.418.398.633.622c.214.225.224.245.346.357l.317.316c.193.184.367.357.55.561c.33.293.64.606.928.939c.164.199.289.427.367.673c.05.23.064.468.041.704"/><path fill="currentColor" d="M19.63 5.136v3.427a.765.765 0 0 1-1.53 0v-1.56l-4.395 4.395a.747.747 0 0 1-.54.225a.755.755 0 0 1-.54-.225a.765.765 0 0 1 0-1.08L17.03 5.91h-1.58a.765.765 0 0 1 0-1.53h3.579c.1.012.196.05.275.113l.113.102a.816.816 0 0 1 .224.52z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhonePlus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10" stroke-width="1.5" d="M12.744 20.255a17.03 17.03 0 0 1-.906-.447a19.324 19.324 0 0 1-4.113-3.132a18.008 18.008 0 0 1-3.89-5.201c-.5-1.06-.83-2.193-.973-3.356a5.91 5.91 0 0 1 .37-3.467c.287-.48.637-.92 1.039-1.31a1.755 1.755 0 0 1 1.263-.592a1.888 1.888 0 0 1 1.285.738c.705.76 1.431 1.466 2.169 2.237c.283.25.463.597.503.973c-.029.299-.15.58-.347.806c-.257.313-.547.593-.827.895a1.533 1.533 0 0 0-.458 1.23c.178.546.48 1.044.883 1.454c.492.66.984 1.32 1.52 1.935a13.816 13.816 0 0 0 3.477 2.864a1.296 1.296 0 0 0 1.24.156a4.003 4.003 0 0 0 1.118-.95a1.72 1.72 0 0 1 1.118-.604c.39.014.758.178 1.028.458c.336.302.638.649.962.973c.324.325.57.548.838.84c.319.288.62.594.906.916a1.33 1.33 0 0 1 .29.996c-.1.44-.344.833-.693 1.118a4.798 4.798 0 0 1-3.767 1.454a10.703 10.703 0 0 1-4.035-.984ZM15.36 4.931v5.134M12.8 7.504h5.13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhonePlusFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.047 18.925a2.655 2.655 0 0 1-.899 1.497c-.544.55-1.199.98-1.92 1.26a5.61 5.61 0 0 1-1.973.362h-.362a11.36 11.36 0 0 1-4.348-1.033c-.34-.155-.66-.31-.96-.475a19.944 19.944 0 0 1-4.369-3.326a18.816 18.816 0 0 1-4.13-5.504A11.867 11.867 0 0 1 2.051 8.08a6.517 6.517 0 0 1 .424-3.853A6.796 6.796 0 0 1 3.663 2.71c.201-.227.448-.41.723-.537c.289-.134.602-.207.92-.217c.363.025.716.131 1.032.31c.293.16.547.382.744.65l1.25 1.292c.32.32.63.64.95.98c.377.335.613.8.66 1.302v.103a1.983 1.983 0 0 1-.485 1.126a8.095 8.095 0 0 1-.578.64l-.279.29a1.27 1.27 0 0 0-.279.413a1.115 1.115 0 0 0-.062.413c.162.464.423.887.765 1.24c.547.722 1.032 1.404 1.59 2.013a13.702 13.702 0 0 0 3.46 2.85a.97.97 0 0 0 .434.187c.136.018.275 0 .403-.052c.356-.223.673-.502.94-.826a2.334 2.334 0 0 1 1.497-.816c.562.006 1.1.232 1.497.63c.197.186.414.402.63.63l.351.372l.331.32c.196.186.372.361.547.557c.335.302.652.623.95.96c.168.2.292.432.362.683c.057.23.067.468.031.702M18.949 7.378a.774.774 0 0 1-.775.775h-1.88v1.87a.775.775 0 1 1-1.549 0v-1.87h-1.89a.774.774 0 1 1 0-1.55h1.87V4.725a.774.774 0 1 1 1.549 0v1.88h1.88a.774.774 0 0 1 .795.774"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneRingingLoud(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10" stroke-width="1.5" d="M6.89 13.042c.254-.1.513-.186.776-.259a14.937 14.937 0 0 1 4.092-.557a14.227 14.227 0 0 1 5.138.736c.88.307 1.707.75 2.45 1.315a4.72 4.72 0 0 1 1.752 2.17c.104.434.154.879.15 1.324a1.426 1.426 0 0 1-.379 1.056a1.563 1.563 0 0 1-1.155.308h-2.46a1.195 1.195 0 0 1-.836-.268a1.095 1.095 0 0 1-.278-.677v-.996a1.206 1.206 0 0 0-.439-.996a2.838 2.838 0 0 0-1.324-.328a20.5 20.5 0 0 0-1.991-.23a11.138 11.138 0 0 0-3.585.34a1.046 1.046 0 0 0-.786.616c-.107.383-.144.78-.11 1.175a1.314 1.314 0 0 1-.279.996c-.23.206-.527.32-.836.319H4.709a8.843 8.843 0 0 1-.996 0a1.085 1.085 0 0 1-.736-.398a1.662 1.662 0 0 1-.22-.996a3.853 3.853 0 0 1 1.285-2.918a8.495 8.495 0 0 1 2.848-1.732Zm8.642-3.744a4.978 4.978 0 0 0-7.05 0m9.161-2.111a7.966 7.966 0 0 0-11.271 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneRingingLoudFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.988 18.053c.012.254-.026.509-.113.749a1.907 1.907 0 0 1-.41.677c-.382.314-.86.487-1.354.492h-2.729a1.488 1.488 0 0 1-1.661-1.457V17.49a.667.667 0 0 0-.062-.36a.75.75 0 0 0-.164-.235a2.391 2.391 0 0 0-1.026-.246a19.175 19.175 0 0 0-2.051-.236c-1.183-.082-2.37.03-3.518.328a.565.565 0 0 0-.298.123a.595.595 0 0 0-.174.215a3.508 3.508 0 0 0-.072.954a1.76 1.76 0 0 1-.051.729a2.051 2.051 0 0 1-.339.666c-.338.329-.79.512-1.261.513H4.51a8.78 8.78 0 0 1-1.026 0a1.537 1.537 0 0 1-.636-.184a1.754 1.754 0 0 1-.492-.41a2.297 2.297 0 0 1-.339-1.396a4.472 4.472 0 0 1 1.488-3.323a9.047 9.047 0 0 1 3.077-1.877c.274-.11.555-.203.84-.277a15.82 15.82 0 0 1 4.34-.595c1.844-.074 3.686.19 5.436.78c.952.339 1.847.82 2.656 1.425a5.374 5.374 0 0 1 1.96 2.452c.122.496.18 1.006.174 1.518m-6.38-7.93a.76.76 0 0 1-.544-.225a4.506 4.506 0 0 0-1.405-.944a4.43 4.43 0 0 0-3.323 0a4.564 4.564 0 0 0-1.416.944a.76.76 0 0 1-1.087 0a.77.77 0 0 1 0-1.088a5.888 5.888 0 0 1 8.318 0a.76.76 0 0 1 0 1.088a.75.75 0 0 1-.543.225"/><path fill="currentColor" d="M17.772 7.96a.76.76 0 0 1-.544-.226a7.416 7.416 0 0 0-8.062-1.6a7.334 7.334 0 0 0-2.41 1.6a.759.759 0 0 1-1.087 0a.77.77 0 0 1 0-1.087a8.934 8.934 0 0 1 12.688.03a.76.76 0 0 1 0 1.088a.75.75 0 0 1-.585.195"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Photoshop(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.45 10.48a1.3 1.3 0 0 1-1 1.33a4.83 4.83 0 0 1-1.61.19c-.14 0-.15-.09-.15-.19V9.26c0-.07.07-.19.11-.19a4.91 4.91 0 0 1 1.71.11a1.28 1.28 0 0 1 .94 1.3"/><path fill="currentColor" d="M17 2H7a5 5 0 0 0-5 5v10a5 5 0 0 0 5 5h10a5 5 0 0 0 5-5V7a5 5 0 0 0-5-5m-6 11a3.48 3.48 0 0 1-1.68.54H7.98c-.28 0-.28 0-.28.27v2.32c0 .2-.06.26-.26.26H6.12c-.19 0-.24-.07-.24-.25V7.71c0-.2 0-.26.25-.26h3.13a3.34 3.34 0 0 1 1.62.47a2.75 2.75 0 0 1 1.4 2.39A2.83 2.83 0 0 1 11 13m5.92 3.3a4.62 4.62 0 0 1-2.73.19c-.272-.06-.54-.14-.8-.24a.29.29 0 0 1-.16-.2v-1.51c.32.12.62.26.93.36a3.66 3.66 0 0 0 1.61.14c.107-.02.212-.053.31-.1a.37.37 0 0 0 .08-.63a4 4 0 0 0-.73-.4c-.41-.2-.83-.36-1.22-.59a1.82 1.82 0 0 1-1-1.93a2 2 0 0 1 1.36-1.63a4.21 4.21 0 0 1 2-.17c.32 0 .63.12.95.18c.17 0 .23.14.22.31v1.17c0 .22-.05.24-.25.16a3.87 3.87 0 0 0-2-.34a.9.9 0 0 0-.28.08a.37.37 0 0 0-.1.63c.289.192.593.36.91.5c.38.19.78.34 1.15.55a1.73 1.73 0 0 1 1 1.79a1.92 1.92 0 0 1-1.26 1.64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pin(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m17.942 6.076l2.442 2.442a1.22 1.22 0 0 1-.147 1.855l-1.757.232a1.697 1.697 0 0 0-.94.452c-.72.696-1.453 1.428-2.674 2.637c-.21.212-.358.478-.427.769l-.94 3.772a1.22 1.22 0 0 1-1.978.379l-3.04-3.052l-3.052-3.04a1.221 1.221 0 0 1 .379-1.978l3.747-.964a1.8 1.8 0 0 0 .77-.44c1.379-1.355 1.88-1.855 2.66-2.698c.233-.25.383-.565.428-.903l.232-1.783a1.221 1.221 0 0 1 1.856-.146zm-9.51 9.498L3.256 20.75"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.447 9.559a1.83 1.83 0 0 1-.25.82a2 2 0 0 1-.56.63a.71.71 0 0 1-.34.13l-1.76.23a1 1 0 0 0-.52.26c-.53.51-1.07 1-1.81 1.78l-.85.84a.93.93 0 0 0-.23.41l-.94 3.78a.56.56 0 0 1 0 .12a2 2 0 0 1-1.44 1.15h-.36a2.25 2.25 0 0 1-.58-.08a1.94 1.94 0 0 1-.81-.49l-2.54-2.53l-4.67 4.67a.75.75 0 0 1-1.06-1.06l4.67-4.67l-2.5-2.5a2 2 0 0 1-.48-.82a1.8 1.8 0 0 1-.05-.95a1.94 1.94 0 0 1 .39-.85a2 2 0 0 1 .75-.58h.12l3.74-1a1 1 0 0 0 .44-.25c1.39-1.37 1.87-1.85 2.63-2.67a.86.86 0 0 0 .23-.5l.24-1.77a.71.71 0 0 1 .13-.35a2 2 0 0 1 2.28-.69a2 2 0 0 1 .72.46l4.88 4.9a2 2 0 0 1 .57 1.55z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pinterest(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.535 14.938a10.1 10.1 0 0 1-4.973 5.901a10.287 10.287 0 0 1-7.734.718a12.94 12.94 0 0 0 1.93-4.935l.441.341a3.345 3.345 0 0 0 2.888.586a4.847 4.847 0 0 0 1.785-.748a4.789 4.789 0 0 0 1.341-1.384a6.58 6.58 0 0 0 .98-5.49a5.54 5.54 0 0 0-1.617-2.732a5.64 5.64 0 0 0-2.866-1.413a7.224 7.224 0 0 0-5.496.831a5.673 5.673 0 0 0-1.879 1.812a5.605 5.605 0 0 0-.859 2.452a5.105 5.105 0 0 0 .916 3.784a2.68 2.68 0 0 0 1.078.906c.259.128.388-.086.474-.352c.184-.522.324-.789.108-1.215a4.232 4.232 0 0 1-.3-4.1a4.273 4.273 0 0 1 1.352-1.646a4.338 4.338 0 0 1 1.988-.799a4.79 4.79 0 0 1 2.931.427a3.747 3.747 0 0 1 1.484 1.316c.366.563.572 1.213.596 1.881a5.589 5.589 0 0 1-.851 3.667a2.743 2.743 0 0 1-.912.932a2.779 2.779 0 0 1-1.244.41a1.469 1.469 0 0 1-1.279-.475a1.434 1.434 0 0 1-.316-1.315c.205-.842.485-1.673.7-2.526c.093-.347.126-.708.098-1.066a1.279 1.279 0 0 0-.664-1.053a1.316 1.316 0 0 0-1.255-.013c-.306.152-.566.38-.756.663c-.189.282-.3.609-.322.947a3.797 3.797 0 0 0 .14 1.961a.823.823 0 0 1 0 .458c-.377 1.62-.808 3.198-1.077 4.85c-.09.771-.116 1.549-.076 2.324v.5a10.115 10.115 0 0 1-5.339-5.049a9.95 9.95 0 0 1-.471-7.29a10.074 10.074 0 0 1 4.645-5.682a10.262 10.262 0 0 1 7.309-1.068c1.326.342 2.569.947 3.652 1.778a10.04 10.04 0 0 1 2.648 3.058a9.903 9.903 0 0 1 .802 7.848"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Play(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.644 5.025V18.97a1.533 1.533 0 0 0 2.467 1.21l9.656-7.509a1.532 1.532 0 0 0-.092-2.483L8.019 3.753a1.533 1.533 0 0 0-2.375 1.272"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayCircle(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 21.5a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19"/><path d="M8.909 7.864v8.27a.909.909 0 0 0 1.463.717l5.725-4.452a.909.909 0 0 0-.055-1.473L10.317 7.11a.909.909 0 0 0-1.408.754"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayCircleFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.75A10.25 10.25 0 1 0 22.25 12A10.26 10.26 0 0 0 12 1.75m4.3 10.57a1.39 1.39 0 0 1-.4.5l-4.76 3.7a1.3 1.3 0 0 1-.7.28h-.14a1.42 1.42 0 0 1-.6-.14a1.47 1.47 0 0 1-.57-.51a1.43 1.43 0 0 1-.2-.73V8.55a1.34 1.34 0 0 1 .19-.7a1.3 1.3 0 0 1 .54-.5a1.25 1.25 0 0 1 .7-.16c.25.007.492.083.7.22l4.76 3.18a1.331 1.331 0 0 1 .61 1.09c.007.22-.037.44-.13.64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.105 11.446a2.34 2.34 0 0 1-.21 1c-.15.332-.38.62-.67.84l-9.65 7.51a2.3 2.3 0 0 1-1.17.46h-.23a2.17 2.17 0 0 1-1-.24a2.29 2.29 0 0 1-1.28-2v-14a2.2 2.2 0 0 1 .33-1.17a2.27 2.27 0 0 1 2.05-1.1c.412.02.812.148 1.16.37l9.66 6.44c.294.204.54.47.72.78c.19.34.29.721.29 1.11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlaySquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M8.99 7.973v8.052a.885.885 0 0 0 1.424.699l5.575-4.336a.884.884 0 0 0-.053-1.433L10.36 7.238a.884.884 0 0 0-1.371.735"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlaySquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 2h-6.5A6.76 6.76 0 0 0 2 8.75v6.5A6.76 6.76 0 0 0 8.75 22h6.5A6.76 6.76 0 0 0 22 15.25v-6.5A6.76 6.76 0 0 0 15.25 2m1 10.39a1.5 1.5 0 0 1-.44.56l-4.51 3.5a1.46 1.46 0 0 1-.78.32h-.16a1.56 1.56 0 0 1-.66-.16a1.4 1.4 0 0 1-.62-.56a1.53 1.53 0 0 1-.24-.8V8.73a1.64 1.64 0 0 1 .22-.78a1.53 1.53 0 0 1 .58-.55a1.6 1.6 0 0 1 .8-.19c.273.022.536.108.77.25l4.51 3c.197.139.361.32.48.53a1.45 1.45 0 0 1 0 1.38z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Playlist(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 5.5h16m-16 6h16m-16 6h8m3.513-2.462v4.922a.54.54 0 0 0 .871.428l3.408-2.65a.54.54 0 0 0-.033-.877l-3.407-2.272a.54.54 0 0 0-.839.45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlaylistAdd(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M4 4.75h16m-16 6h16m-16 6h8"/><path stroke-miterlimit="10" d="M17.745 14.5V19M15.5 16.755H20"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlaylistAlternate(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 6.5h8m-16 6h16m-16 6h16M4 4.038V8.96a.54.54 0 0 0 .87.428l3.409-2.65a.54.54 0 0 0-.033-.877L4.838 3.589A.54.54 0 0 0 4 4.04"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlaylistAlternateFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 8.266h-8a.75.75 0 1 1 0-1.5h8a.75.75 0 1 1 0 1.5m0 6H4a.75.75 0 1 1 0-1.5h16a.75.75 0 1 1 0 1.5m0 6H4a.75.75 0 1 1 0-1.5h16a.75.75 0 1 1 0 1.5M9.23 7.286a1.23 1.23 0 0 1-.12.59a1.29 1.29 0 0 1-.37.47l-3.41 2.67c-.19.15-.42.24-.66.26h-.13a1.249 1.249 0 0 1-.57-.13a1.31 1.31 0 0 1-.52-.47a1.34 1.34 0 0 1-.2-.69v-4.97a1.31 1.31 0 0 1 .68-1.13a1.27 1.27 0 0 1 .67-.15c.231.013.455.085.65.21l3.41 2.28c.17.11.312.262.41.44c.108.188.163.403.16.62"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlaylistFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20 5.245H4a.75.75 0 0 1 0-1.5h16a.75.75 0 1 1 0 1.5m0 6H4a.75.75 0 1 1 0-1.5h16a.75.75 0 1 1 0 1.5m-8 6H4a.75.75 0 1 1 0-1.5h8a.75.75 0 1 1 0 1.5m8.72-1c.002.21-.039.417-.12.61a1.43 1.43 0 0 1-.38.47l-3.38 2.67c-.19.15-.42.24-.66.26h-.13a1.28 1.28 0 0 1-.57-.13a1.38 1.38 0 0 1-.52-.47a1.28 1.28 0 0 1-.2-.68v-4.98a1.35 1.35 0 0 1 .68-1.13a1.32 1.32 0 0 1 .67-.15c.232.008.457.081.65.21l3.41 2.28a1.28 1.28 0 0 1 .58 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Playstore(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m3.637 3.434l8.74 8.571l-8.675 8.65a2.123 2.123 0 0 1-.326-.613a2.45 2.45 0 0 1 0-.755V4.567c-.026-.395.065-.79.26-1.133m12.506 4.833l-2.853 2.826L4.653 2.6c.28-.097.58-.124.873-.078c.46.126.899.32 1.302.573l7.816 4.325c.508.273 1.003.56 1.498.847M13.29 12.93l2.839 2.788l-2.058 1.146l-6.279 3.49c-.52.287-1.042.561-1.55.874a1.798 1.798 0 0 1-1.472.195zm7.36-.925a1.915 1.915 0 0 1-.99 1.72l-2.346 1.302l-3.087-3.022l3.1-3.074c.795.443 1.577.886 2.358 1.303a1.888 1.888 0 0 1 .964 1.771"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 4.5v15m7.5-7.5h-15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusCircle(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 6.722v10.556M17.278 12H6.722M12 21.5a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusCircleFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.96 1.75A10.25 10.25 0 1 0 22.25 12A10.261 10.261 0 0 0 11.96 1.75M17.25 13h-4.28v4.27a1 1 0 0 1-2 0V13H6.69a1 1 0 1 1 0-2h4.28V6.68a1 1 0 0 1 2 0v4.28h4.28a1 1 0 0 1 0 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M12 6.861V17.14M17.14 12H6.86"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusSquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 2h-6.5A6.76 6.76 0 0 0 2 8.75v6.5A6.76 6.76 0 0 0 8.75 22h6.5A6.76 6.76 0 0 0 22 15.25v-6.5A6.76 6.76 0 0 0 15.25 2m1.89 11H13v4.14a1 1 0 0 1-2 0V13H6.86a1 1 0 0 1 0-2H11V6.86a1 1 0 0 1 2 0V11h4.14a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PremierPro(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.63 11.11a1.16 1.16 0 0 1-.81.77A4 4 0 0 1 8.2 12c-.11 0-.15 0-.15-.16V9.4c0-.09 0-.17.12-.17c.474-.007.949.017 1.42.07a1.3 1.3 0 0 1 1.04 1.81"/><path fill="currentColor" d="M17 2H7a5 5 0 0 0-5 5v10a5 5 0 0 0 5 5h10a5 5 0 0 0 5-5V7a5 5 0 0 0-5-5m-5.14 10.57a3 3 0 0 1-2 1c-.5.06-1 0-1.5.08H8v2.57c0 .2 0 .23-.21.23H6.37c-.15 0-.21-.05-.21-.21V7.79c0-.13 0-.2.18-.2H9.1A4 4 0 0 1 11 8a2.71 2.71 0 0 1 1.55 2.43a2.88 2.88 0 0 1-.69 2.14m5.94-1.29c0 .14-.08.17-.2.17a3.621 3.621 0 0 0-1.45.26l-.13.06a.58.58 0 0 0-.4.62v3.82c0 .23 0 .24-.23.24h-1.45c-.21 0-.23 0-.23-.23v-5.19c0-.38 0-.76-.05-1.14c0-.13 0-.19.17-.18h1.33a.23.23 0 0 1 .27.22c0 .2.07.41.1.58c.3-.21.59-.44.91-.63a2.35 2.35 0 0 1 1.19-.31c.12 0 .18 0 .18.17z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Preview(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M10.584 21.25h-1.25a6 6 0 0 1-6-6v-6.5a6 6 0 0 1 6-6h6.5a6 6 0 0 1 6 6V10"/><path stroke-linejoin="round" d="m16.55 16.864l-.51 2.959a1.164 1.164 0 0 1-2.212.093l-2.789-7.882a1.164 1.164 0 0 1 1.493-1.493l8.036 2.788a1.164 1.164 0 0 1-.115 2.234l-3.113.51a1.166 1.166 0 0 0-.79.791"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PreviewCircle(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m16.219 17.114l-.51 2.959a1.164 1.164 0 0 1-2.213.093l-2.788-7.882A1.164 1.164 0 0 1 12.2 10.79l8.036 2.788a1.164 1.164 0 0 1-.116 2.234l-3.112.51a1.165 1.165 0 0 0-.79.791"/><path d="M21.502 9.314A9.726 9.726 0 1 0 9.297 21.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PreviewCircleFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.773 14.768c-.029.414-.186.81-.45 1.13a1.892 1.892 0 0 1-.998.63l-3.157.521l-.09.09a.381.381 0 0 0-.09.15l-.5 2.902a1.922 1.922 0 0 1-1.778 1.471h-.09c-.374 0-.74-.111-1.05-.32a1.911 1.911 0 0 1-.739-.92l-2.787-7.906a1.904 1.904 0 0 1 .45-2.001c.253-.263.58-.44.939-.51a1.866 1.866 0 0 1 1.069.07l7.992 2.781c.404.135.754.394 1 .74c.215.351.313.761.28 1.172"/><path fill="currentColor" d="M9.305 22.243a.834.834 0 0 1-.22 0a10.469 10.469 0 0 1-4.5-2.83a10.491 10.491 0 0 1-2.448-10A10.49 10.49 0 0 1 4.82 4.819a10.467 10.467 0 0 1 9.902-2.765a10.466 10.466 0 0 1 4.669 2.54a10.49 10.49 0 0 1 2.822 4.51a.743.743 0 0 1-1.059.886a.76.76 0 0 1-.37-.436a9.008 9.008 0 0 0-2.41-3.894a8.988 8.988 0 0 0-8.585-2.143a8.989 8.989 0 0 0-3.953 2.306a9.01 9.01 0 0 0-2.377 8.536a8.99 8.99 0 0 0 6.075 6.443a.77.77 0 0 1 .49 1a.75.75 0 0 1-.72.44"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PreviewFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 22H8.75A6.76 6.76 0 0 1 2 15.25v-6.5A6.76 6.76 0 0 1 8.75 2h6.5A6.76 6.76 0 0 1 22 8.75V10a.75.75 0 1 1-1.5 0V8.75a5.26 5.26 0 0 0-5.25-5.25h-6.5A5.26 5.26 0 0 0 3.5 8.75v6.5a5.26 5.26 0 0 0 5.25 5.25H10a.75.75 0 1 1 0 1.5"/><path fill="currentColor" d="M21.52 14.53a1.89 1.89 0 0 1-1.48 1.76l-3.16.52l-.09.09a.418.418 0 0 0-.09.14l-.5 2.91a1.94 1.94 0 0 1-.68 1.06a1.91 1.91 0 0 1-1.1.41h-.09a1.94 1.94 0 0 1-1.79-1.24l-2.79-7.9a1.9 1.9 0 0 1 .45-2a1.82 1.82 0 0 1 .94-.51a1.87 1.87 0 0 1 1.07.07l8 2.78a1.93 1.93 0 0 1 1.29 1.91z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Previous(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.889 3.7v16.6m16.222-1.676V5.38a1.455 1.455 0 0 0-2.343-1.15L8.6 11.363a1.456 1.456 0 0 0 .087 2.357l9.169 6.113a1.456 1.456 0 0 0 2.255-1.208"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PreviousFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.885 21.06a.76.76 0 0 1-.75-.75V3.69a.75.75 0 0 1 1.5 0v16.6a.75.75 0 0 1-.75.77m16.98-15.713v13.25a2.35 2.35 0 0 1-.32 1.13a2.2 2.2 0 0 1-1.89 1.07h-.1a2.089 2.089 0 0 1-1.11-.36l-9.13-6.12a2.25 2.25 0 0 1-.71-.76a2.29 2.29 0 0 1-.27-1a2.18 2.18 0 0 1 .2-1a2.22 2.22 0 0 1 .64-.81l9.14-7.09a2.22 2.22 0 0 1 1.13-.44a2.2 2.2 0 0 1 2.09 1.02c.204.335.318.718.33 1.11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Printer(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16.625 14.313h-9.25c-.639 0-1.156.517-1.156 1.156v4.625c0 .638.517 1.156 1.156 1.156h9.25c.639 0 1.156-.518 1.156-1.156v-4.625c0-.639-.517-1.156-1.156-1.156"/><path d="M17.781 17.781h2.313a1.156 1.156 0 0 0 1.156-1.156v-5.781a2.312 2.312 0 0 0-2.312-2.313H5.063a2.312 2.312 0 0 0-2.313 2.313v5.781a1.156 1.156 0 0 0 1.156 1.156H6.22"/><path d="M16.625 8.531V3.906a1.156 1.156 0 0 0-1.156-1.156H8.53a1.156 1.156 0 0 0-1.156 1.156v4.625"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PrinterFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.1 8.67a3 3 0 0 0-2.16-.89h-1.56V3.9A1.89 1.89 0 0 0 15.47 2H8.53a1.87 1.87 0 0 0-1.34.55a1.9 1.9 0 0 0-.56 1.35v3.88H5.06A3.06 3.06 0 0 0 2 10.84v5.78a1.91 1.91 0 0 0 1.91 1.91h1.56v1.56A1.92 1.92 0 0 0 7.38 22h9.25a1.91 1.91 0 0 0 1.9-1.91v-1.56h1.57a1.91 1.91 0 0 0 1.9-1.91v-5.78a3 3 0 0 0-.9-2.17m-.6 7.95a.398.398 0 0 1-.12.29a.42.42 0 0 1-.28.12h-1.57v-1.57a1.9 1.9 0 0 0-1.9-1.9H7.38a1.91 1.91 0 0 0-1.91 1.9v1.57H3.91a.45.45 0 0 1-.29-.12a.4.4 0 0 1-.12-.29v-5.78a1.61 1.61 0 0 1 .46-1.11a1.56 1.56 0 0 1 1.1-.45h13.88a1.55 1.55 0 0 1 1.56 1.56z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QrCode(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.433 5.833h-2.7a.9.9 0 0 0-.9.9v2.7a.9.9 0 0 0 .9.9h2.7a.9.9 0 0 0 .9-.9v-2.7a.9.9 0 0 0-.9-.9m0 7.815h-2.7a.9.9 0 0 0-.9.9v2.7a.9.9 0 0 0 .9.9h2.7a.9.9 0 0 0 .9-.9v-2.7a.9.9 0 0 0-.9-.9m7.834-7.815h-2.7a.9.9 0 0 0-.9.9v2.7a.9.9 0 0 0 .9.9h2.7a.9.9 0 0 0 .9-.9v-2.7a.9.9 0 0 0-.9-.9m0 7.834h-2.7a.9.9 0 0 0-.9.9v2.7a.9.9 0 0 0 .9.9h2.7a.9.9 0 0 0 .9-.9v-2.7a.9.9 0 0 0-.9-.9m3.983-4.75V5.833a3.083 3.083 0 0 0-3.083-3.083h-3.084m0 18.5h3.084a3.083 3.083 0 0 0 3.083-3.083v-3.084m-18.5 0v3.084a3.083 3.083 0 0 0 3.083 3.083h3.084m0-18.5H5.833A3.083 3.083 0 0 0 2.75 5.833v3.084"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QrCodeFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.12 6.96v2.595a1.574 1.574 0 0 1-.968 1.484a1.584 1.584 0 0 1-.612.123h-2.6a1.591 1.591 0 0 1-1.59-1.577V6.99a1.585 1.585 0 0 1 1.59-1.587h2.6a1.592 1.592 0 0 1 1.58 1.557m0 7.495v2.595a1.585 1.585 0 0 1-1.58 1.587h-2.6a1.592 1.592 0 0 1-1.59-1.587v-2.595a1.585 1.585 0 0 1 1.59-1.577h2.6a1.582 1.582 0 0 1 1.58 1.577m7.54-7.495v2.595a1.585 1.585 0 0 1-1.54 1.607h-2.59a1.591 1.591 0 0 1-1.59-1.577V6.99a1.585 1.585 0 0 1 1.59-1.587h2.59a1.592 1.592 0 0 1 1.54 1.557m0 7.515v2.595a1.585 1.585 0 0 1-1.54 1.587h-2.59a1.592 1.592 0 0 1-1.59-1.587v-2.595a1.586 1.586 0 0 1 1.59-1.577h2.59a1.591 1.591 0 0 1 1.54 1.577"/><path fill="currentColor" d="M21.25 9.695a.76.76 0 0 1-.75-.749V5.862a2.31 2.31 0 0 0-.68-1.686a2.352 2.352 0 0 0-1.65-.679h-3.05a.75.75 0 0 1-.75-.748a.748.748 0 0 1 .75-.749h3.08a3.864 3.864 0 0 1 2.68 1.178A3.79 3.79 0 0 1 22 5.882v3.084a.748.748 0 0 1-.75.729M18.17 22h-3.05a.75.75 0 0 1-.75-.748a.748.748 0 0 1 .75-.749h3.08a2.324 2.324 0 0 0 2.137-1.462a2.31 2.31 0 0 0 .163-.893v-3.054a.748.748 0 0 1 .75-.749a.75.75 0 0 1 .75.749v3.054a3.806 3.806 0 0 1-2.363 3.534c-.465.191-.964.29-1.467.288zm-9.25 0H5.84a3.848 3.848 0 0 1-2.722-1.13A3.832 3.832 0 0 1 2 18.149v-3.054a.748.748 0 0 1 .75-.749a.751.751 0 0 1 .75.749v3.054a2.334 2.334 0 0 0 2.34 2.325h3.08a.75.75 0 0 1 .75.749a.747.747 0 0 1-.75.748zM2.75 9.695A.76.76 0 0 1 2 8.946V5.862a3.799 3.799 0 0 1 1.12-2.684A3.864 3.864 0 0 1 5.83 2.06h3.08a.75.75 0 0 1 .75.749a.748.748 0 0 1-.75.748H5.83a2.333 2.333 0 0 0-2.34 2.325v3.084a.748.748 0 0 1-.74.729"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionMarkCircle(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round"><path stroke-miterlimit="10" stroke-width="1.5" d="M9.008 8.84a3.185 3.185 0 0 1 3.471-1.806a3.09 3.09 0 0 1 2.265 1.614a2.682 2.682 0 0 1-1.562 3.689a1.98 1.98 0 0 0-1.276 1.787v.738"/><path stroke-linejoin="round" stroke-width="2" d="M11.881 17.424h.008"/><path stroke-linejoin="round" stroke-width="1.5" d="M12 21.5a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionMarkCircleFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.75A10.25 10.25 0 1 0 22.25 12A10.26 10.26 0 0 0 12 1.75m-.11 16.8a1 1 0 1 1 0-2a1 1 0 0 1 0 2m3.82-7.32a3.73 3.73 0 0 1-2.21 2.05a1 1 0 0 0-.41.33a.9.9 0 0 0-.18.54v.71a1 1 0 1 1-2 0v-.74a2.88 2.88 0 0 1 .55-1.67a2.92 2.92 0 0 1 1.37-1a1.59 1.59 0 0 0 .63-.38A1.63 1.63 0 0 0 14 9.81a1.66 1.66 0 0 0-.16-.69a2 2 0 0 0-.62-.69a2 2 0 0 0-.89-.36a2.27 2.27 0 0 0-1.44.2a2.2 2.2 0 0 0-1 1a1 1 0 0 1-1.82-.83a4.17 4.17 0 0 1 4.56-2.37a4 4 0 0 1 1.72.7A4.1 4.1 0 0 1 15.6 8.2a3.7 3.7 0 0 1 .08 3.05z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionMarkSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-miterlimit="10" stroke-width="1.5" d="M8.88 8.54a3.29 3.29 0 0 1 3.576-1.86a3.201 3.201 0 0 1 2.334 1.663a2.76 2.76 0 0 1-1.61 3.8a2.048 2.048 0 0 0-1.314 1.842v.76"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.96 17.162h.009"/><rect width="18.5" height="18.5" x="2.75" y="2.75" stroke-width="1.5" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuestionMarkSquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 2h-6.5A6.76 6.76 0 0 0 2 8.75v6.5A6.76 6.76 0 0 0 8.75 22h6.5A6.76 6.76 0 0 0 22 15.25v-6.5A6.76 6.76 0 0 0 15.25 2m-3.28 16.58a1 1 0 1 1 0-2a1 1 0 0 1 0 2m3.79-7.59a3.69 3.69 0 0 1-.91 1.29a3.59 3.59 0 0 1-1.35.81a1 1 0 0 0-.63.93v.73a1 1 0 0 1-2 0v-.77a3.06 3.06 0 0 1 .56-1.7a3.13 3.13 0 0 1 1.39-1.07a1.77 1.77 0 0 0 .67-.39c.186-.17.332-.378.43-.61a1.77 1.77 0 0 0 .13-.72a1.8 1.8 0 0 0-.16-.72a2.32 2.32 0 0 0-.65-.72c-.279-.2-.6-.33-.94-.38a2.29 2.29 0 0 0-2.51 1.29a1.002 1.002 0 0 1-1.82-.84a4.35 4.35 0 0 1 1.92-2a4.27 4.27 0 0 1 2.74-.41a4.08 4.08 0 0 1 1.77.72c.53.379.964.875 1.27 1.45c.24.493.37 1.032.38 1.58a3.551 3.551 0 0 1-.29 1.53"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reddit(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a10 10 0 1 0 0 20a10 10 0 0 0 0-20m6.235 11.43c-.072 0-.155.123-.145.196a3.58 3.58 0 0 1-1.676 3.26a6.688 6.688 0 0 1-3.025 1.184a8.231 8.231 0 0 1-4.115-.36a5.412 5.412 0 0 1-2.799-2.058a3.086 3.086 0 0 1-.514-1.996a.267.267 0 0 0-.154-.298a1.534 1.534 0 1 1 1.687-2.49a.195.195 0 0 0 .288 0a7.202 7.202 0 0 1 3.56-1.162c.4 0 .4 0 .483-.422c.226-1.029.453-2.058.669-3.086c.062-.289.185-.37.463-.31l2.253.474a.216.216 0 0 0 .257-.092a1.029 1.029 0 0 1 1.204-.433a1.101 1.101 0 0 1 .782 1.03a1.132 1.132 0 0 1-.659 1.028a1.111 1.111 0 0 1-1.45-.607c0-.113-.062-.226-.103-.35l-2.212-.463c-.226 1.03-.453 2.14-.69 3.22a7.66 7.66 0 0 1 4.054 1.297a1.523 1.523 0 0 1 1.163-.412a1.492 1.492 0 0 1 1.378 1.142a1.554 1.554 0 0 1-.69 1.708z"/><path fill="currentColor" d="M14.53 16.105a3.951 3.951 0 0 1-2.345.751a4.023 4.023 0 0 1-2.736-.792a.236.236 0 0 1 0-.36a.259.259 0 0 1 .303-.067a.259.259 0 0 1 .088.067a4.116 4.116 0 0 0 3.827.35c.19-.084.37-.192.535-.32c.298-.298.648.103.329.37m-3.837-2.921a1.08 1.08 0 0 1-1.348 1.029h-.093a1.029 1.029 0 0 1-.7-.977a1.029 1.029 0 0 1 1.03-1.091a1.101 1.101 0 0 1 1.11 1.04m4.774-.001a1.028 1.028 0 1 1-2.05-.164a1.028 1.028 0 0 1 2.05.164"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Refresh(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="M17.605 7.705A7.885 7.885 0 0 0 12 5.382a7.929 7.929 0 0 0-7.929 7.929A7.94 7.94 0 0 0 12 21.25a7.94 7.94 0 0 0 7.929-7.94"/><path stroke-linejoin="round" d="m16.88 2.75l.95 3.858a1.332 1.332 0 0 1-.97 1.609l-3.869.948"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RefreshReverse(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="M6.395 7.705A7.885 7.885 0 0 1 12 5.382a7.929 7.929 0 0 1 7.929 7.929A7.94 7.94 0 0 1 12 21.25a7.939 7.939 0 0 1-7.929-7.94"/><path stroke-linejoin="round" d="m7.12 2.75l-.95 3.858a1.332 1.332 0 0 0 .97 1.609l3.869.948"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reload(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="M18.024 7.043A8.374 8.374 0 0 0 3.74 12.955"/><path stroke-linejoin="round" d="m17.35 2.75l.832 3.372a1.123 1.123 0 0 1-.854 1.382l-3.372.843"/><path stroke-miterlimit="10" d="M5.976 16.957a8.374 8.374 0 0 0 14.285-5.912"/><path stroke-linejoin="round" d="m6.65 21.25l-.832-3.372a1.124 1.124 0 0 1 .855-1.382l3.371-.843"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReloadReverse(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="M18.024 16.957A8.373 8.373 0 0 1 3.74 11.045"/><path stroke-linejoin="round" d="m17.35 21.25l.832-3.372a1.123 1.123 0 0 0-.854-1.382l-3.372-.843"/><path stroke-miterlimit="10" d="M5.976 7.043a8.373 8.373 0 0 1 14.285 5.912"/><path stroke-linejoin="round" d="m6.65 2.75l-.832 3.372a1.124 1.124 0 0 0 .855 1.382l3.371.843"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ribbon(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M14.432 10.606L17.14 12.3a2.844 2.844 0 0 0 3.964-.97c.259-.44.395-.94.396-1.449V6.853a2.863 2.863 0 0 0-1.47-2.496a2.844 2.844 0 0 0-2.89.077L14.432 6.13m-4.864 0L6.861 4.434a2.845 2.845 0 0 0-3.965.97c-.259.44-.395.94-.396 1.45v3.028a2.863 2.863 0 0 0 .942 2.14a2.85 2.85 0 0 0 2.222.717a2.702 2.702 0 0 0 1.197-.419l2.707-1.695"/><path d="M14.85 7.377v1.981c.001.42-.136.828-.39 1.162l-.066.086a1.9 1.9 0 0 1-1.444.657h-1.9a1.896 1.896 0 0 1-1.444-.657a1.31 1.31 0 0 1-.142-.2a1.908 1.908 0 0 1-.314-1.048v-1.98a1.907 1.907 0 0 1 1.9-1.906h1.9a1.9 1.9 0 0 1 1.9 1.905m3.8 5.363l1.51 5.637a.191.191 0 0 1-.177.239a.19.19 0 0 1-.108-.03l-3.087-1.79l-1.748 3.105a.19.19 0 0 1-.294.05a.192.192 0 0 1-.058-.098L12.4 11.282"/><path d="M9.426 10.406c.04.07.088.138.142.2L6.861 12.3c-.36.233-.771.376-1.197.42l-1.51 5.656a.19.19 0 0 0 .275.21l3.097-1.79L9.312 19.9a.19.19 0 0 0 .192.097a.19.19 0 0 0 .159-.145l2.29-8.571m4.797-2.913h-1.862m-5.776 0H7.25"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RibbonFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.778 5.024a3.667 3.667 0 0 0-1.364-1.321a3.506 3.506 0 0 0-1.845-.45a3.625 3.625 0 0 0-1.815.54l-2.277 1.421a2.59 2.59 0 0 0-1.524-.49h-1.876c-.553 0-1.093.17-1.544.49L7.266 3.793a3.625 3.625 0 0 0-1.815-.54a3.525 3.525 0 0 0-1.845.45A3.617 3.617 0 0 0 1.75 6.856v3.003c-.007.51.096 1.015.301 1.481c.204.468.51.885.893 1.222c.376.342.82.6 1.304.76c.146.051.297.088.45.11l-1.273 4.775a.86.86 0 0 0 0 .51c.056.176.164.33.311.441c.145.11.32.177.502.19a.923.923 0 0 0 .541-.12l2.457-1.411L8.64 20.25a.882.882 0 0 0 .411.4c.132.063.276.098.422.1h.12a.92.92 0 0 0 .782-.7l1.796-6.708l1.785 6.667a.94.94 0 0 0 .27.49c.149.135.334.222.532.251h.12c.154-.005.304-.042.442-.11a.993.993 0 0 0 .38-.37l1.385-2.453l2.417 1.401a.994.994 0 0 0 .521.14a.87.87 0 0 0 .522-.17a.891.891 0 0 0 .33-.44a.91.91 0 0 0 0-.57l-1.293-4.826c.28-.078.549-.19.802-.33A3.608 3.608 0 0 0 22.25 9.87V6.866a3.638 3.638 0 0 0-.472-1.842m-6.167 2.593h1.153a.753.753 0 0 1 .752.75a.75.75 0 0 1-.752.751H15.61zm-8.375 0H8.39v1.471H7.236a.753.753 0 0 1-.752-.75a.75.75 0 0 1 .752-.751zm2.006 10.75l-1.112-1.93a.752.752 0 0 0-1.003-.28l-1.936 1.12l1.003-3.883c.357-.091.695-.24 1.003-.441l2.257-1.411c.13.091.267.172.41.24c.318.14.658.22 1.004.24zm1.765-7.847a1.125 1.125 0 0 1-.481-.1a1.143 1.143 0 0 1-.391-.3l-.08-.121a1.12 1.12 0 0 1-.191-.64V7.355a1.16 1.16 0 0 1 1.153-1.15h1.906a1.145 1.145 0 0 1 .812.34c.215.215.337.506.341.81v2.002c-.003.249-.054.49-.202.691l-.029.06a1.004 1.004 0 0 1-.39.3a1.125 1.125 0 0 1-.482.1zm6.138 5.635a.753.753 0 0 0-.572-.07a.743.743 0 0 0-.461.35l-1.073 1.913l-1.705-6.357a2.678 2.678 0 0 0 1.123-.46l2.267 1.411c.4.25.847.418 1.313.49l1.003 3.834z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Robot(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M14.706 4.313H9.294a4.981 4.981 0 0 0-4.982 4.981v5.412a4.982 4.982 0 0 0 4.982 4.982h5.412a4.982 4.982 0 0 0 4.982-4.982V9.294a4.982 4.982 0 0 0-4.982-4.982Z"/><path d="M19.606 15.588h1.619a1.025 1.025 0 0 0 1.025-1.025V9.438a1.025 1.025 0 0 0-1.025-1.025h-1.62m-15.21 7.175h-1.62a1.025 1.025 0 0 1-1.025-1.025V9.438a1.025 1.025 0 0 1 1.025-1.025h1.62"/><path stroke-linecap="round" stroke-linejoin="round" d="M2.765 8.413v-4.1m18.46 4.1l-.01-4.1m-12.3 4.45v2.698m6.15-2.698v2.698M9.94 15.588h4.1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RobotAppreciate(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M14.706 4.313H9.294a4.981 4.981 0 0 0-4.982 4.981v5.412a4.982 4.982 0 0 0 4.982 4.982h5.412a4.982 4.982 0 0 0 4.982-4.982V9.294a4.982 4.982 0 0 0-4.982-4.982Z"/><path d="M19.606 15.588h1.619a1.025 1.025 0 0 0 1.025-1.025V9.438a1.025 1.025 0 0 0-1.025-1.025h-1.62m-15.21 7.175h-1.62a1.025 1.025 0 0 1-1.025-1.025V9.438a1.025 1.025 0 0 1 1.025-1.025h1.62"/><path stroke-linecap="round" stroke-linejoin="round" d="M2.765 8.413v-4.1m18.46 4.1l-.01-4.1M9.94 15.588h4.1m-6.16-4.613L8.903 9.95l1.025 1.025m4.102 0l1.025-1.025l1.024 1.025"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RobotAppreciateFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.078 8.347a1.368 1.368 0 0 0-.488-.325V4.647a.717.717 0 0 0-.717-.717a.755.755 0 0 0-.717.717V7.85h-.21a5.479 5.479 0 0 0-5.25-3.92H9.427a5.48 5.48 0 0 0-5.25 3.92H3.9V4.647a.717.717 0 1 0-1.434 0v3.385a1.501 1.501 0 0 0-.469.315A1.722 1.722 0 0 0 1.5 9.552v4.896a1.702 1.702 0 0 0 1.702 1.702h.956a5.48 5.48 0 0 0 5.25 3.92h5.183a5.48 5.48 0 0 0 5.25-3.92h.955a1.702 1.702 0 0 0 1.702-1.702V9.552c.02-.44-.131-.872-.42-1.205M3.996 14.716H3.24a.267.267 0 0 1-.191-.077a.287.287 0 0 1-.076-.191V9.552a.258.258 0 0 1 .248-.268h.775a.545.545 0 0 0 0 .125v5.182a.545.545 0 0 0 0 .125m4.695-3.118a.832.832 0 0 1-1.148 0a.822.822 0 0 1 0-1.157l.956-.956a.813.813 0 0 1 1.148 0l.956.956a.822.822 0 0 1 0 1.157a.813.813 0 0 1-1.147 0l-.412-.41zm5.316 4.648h-3.92a.813.813 0 0 1 0-1.626h3.92a.813.813 0 1 1 0 1.626m2.525-4.648a.822.822 0 0 1-.574.23a.813.813 0 0 1-.574-.23l-.411-.41l-.402.41a.813.813 0 0 1-1.385-.578c0-.217.086-.425.238-.579l.956-.956a.822.822 0 0 1 1.157 0l.956.956a.822.822 0 0 1 .039 1.157m4.608 2.869a.287.287 0 0 1-.076.191a.267.267 0 0 1-.191.077h-.756a.55.55 0 0 0 0-.125V9.37a.55.55 0 0 0 0-.124h.765a.248.248 0 0 1 .182.077c.048.052.076.12.076.19z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RobotDead(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M14.706 4.313H9.294a4.982 4.982 0 0 0-4.982 4.981v5.412a4.982 4.982 0 0 0 4.982 4.982h5.412a4.982 4.982 0 0 0 4.982-4.982V9.294a4.982 4.982 0 0 0-4.982-4.982Z"/><path d="M19.606 15.588h1.619a1.025 1.025 0 0 0 1.025-1.025V9.438a1.025 1.025 0 0 0-1.025-1.025h-1.62m-15.21 7.175h-1.62a1.025 1.025 0 0 1-1.025-1.025V9.438a1.025 1.025 0 0 1 1.025-1.025h1.62"/><path stroke-linecap="round" stroke-linejoin="round" d="M2.765 8.413v-4.1m18.46 4.1l-.01-4.1M9.897 9.372l-2.102 2.102m0-2.102l2.102 2.102m4.206-2.102l2.102 2.102m0-2.102l-2.102 2.102M9.899 16.1h4.1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RobotDeadFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.078 8.347a1.368 1.368 0 0 0-.488-.325V4.647a.717.717 0 1 0-1.434 0V7.85h-.21a5.479 5.479 0 0 0-5.25-3.92H9.427a5.48 5.48 0 0 0-5.25 3.92H3.9V4.647a.717.717 0 1 0-1.434 0v3.385a1.501 1.501 0 0 0-.469.315A1.722 1.722 0 0 0 1.5 9.552v4.896a1.702 1.702 0 0 0 1.702 1.702h.956a5.48 5.48 0 0 0 5.25 3.92h5.183a5.48 5.48 0 0 0 5.25-3.92h.955a1.702 1.702 0 0 0 1.702-1.702V9.552c.02-.44-.131-.872-.42-1.205M3.996 14.716H3.24a.267.267 0 0 1-.191-.077a.287.287 0 0 1-.076-.191V9.552a.258.258 0 0 1 .248-.268h.775a.545.545 0 0 0 0 .125v5.182a.545.545 0 0 0 0 .125m4.618-2.64a.841.841 0 0 1-1.157 0a.822.822 0 0 1 0-1.157l.43-.43l-.43-.43a.813.813 0 0 1 .579-1.385c.216 0 .424.085.578.237l.43.43l.43-.43a.812.812 0 1 1 1.148 1.148l-.43.43l.43.43a.812.812 0 0 1-.573 1.387a.813.813 0 0 1-.574-.23l-.43-.43zm5.355 4.657h-3.92a.813.813 0 1 1 0-1.625h3.92a.813.813 0 1 1 0 1.625m2.687-5.814a.822.822 0 0 1 0 1.157a.841.841 0 0 1-1.157 0l-.43-.43l-.43.43a.812.812 0 0 1-.574.23a.813.813 0 0 1-.574-1.386l.43-.43l-.43-.431a.812.812 0 1 1 1.147-1.148l.43.43l.43-.43a.822.822 0 0 1 1.158 0a.812.812 0 0 1 0 1.148l-.43.43zm4.485 3.529a.287.287 0 0 1-.077.191a.265.265 0 0 1-.191.077h-.756a.55.55 0 0 0 0-.125v-5.22a.55.55 0 0 0 0-.125h.765a.248.248 0 0 1 .182.077c.048.052.076.12.076.19z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RobotFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.078 8.347a1.368 1.368 0 0 0-.488-.325V4.647a.717.717 0 0 0-.717-.717a.727.727 0 0 0-.717.717V7.85h-.21a5.479 5.479 0 0 0-5.25-3.92H9.427a5.48 5.48 0 0 0-5.25 3.92H3.9V4.647a.717.717 0 1 0-1.434 0v3.385a1.501 1.501 0 0 0-.469.315A1.722 1.722 0 0 0 1.5 9.552v4.896a1.702 1.702 0 0 0 1.702 1.702h.956a5.48 5.48 0 0 0 5.25 3.92h5.183a5.48 5.48 0 0 0 5.25-3.92h.955a1.702 1.702 0 0 0 1.702-1.702V9.552c.02-.44-.131-.872-.42-1.205M3.996 14.716H3.24a.267.267 0 0 1-.191-.077a.287.287 0 0 1-.076-.191V9.552a.258.258 0 0 1 .248-.268h.775a.545.545 0 0 0 0 .125v5.182a.545.545 0 0 0 0 .125m4.303-3.232V8.902a.813.813 0 1 1 1.616 0v2.582a.813.813 0 1 1-1.616 0m5.737 4.78h-3.92a.812.812 0 1 1 0-1.625h3.92a.813.813 0 0 1 0 1.626m1.788-4.78a.813.813 0 1 1-1.626 0V8.902a.813.813 0 0 1 1.626 0zm5.345 2.964c0 .07-.028.14-.076.191a.265.265 0 0 1-.192.077h-.755a.55.55 0 0 0 0-.125v-5.22a.55.55 0 0 0 0-.125h.765a.249.249 0 0 1 .182.077c.048.052.075.12.076.19z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RobotHappy(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M14.706 4.313H9.294a4.982 4.982 0 0 0-4.982 4.981v5.412a4.982 4.982 0 0 0 4.982 4.982h5.412a4.982 4.982 0 0 0 4.982-4.982V9.294a4.982 4.982 0 0 0-4.982-4.982Z"/><path d="M19.606 15.588h1.619a1.025 1.025 0 0 0 1.025-1.025V9.438a1.025 1.025 0 0 0-1.025-1.025h-1.62m-15.21 7.175h-1.62a1.025 1.025 0 0 1-1.025-1.025V9.438a1.025 1.025 0 0 1 1.025-1.025h1.62"/><path stroke-linecap="round" stroke-linejoin="round" d="M2.765 8.413v-4.1m18.46 4.1l-.01-4.1M9.95 15.237a2.91 2.91 0 0 0 4.1 0m-6.17-4.262L8.903 9.95l1.025 1.025m4.102 0l1.025-1.025l1.024 1.025"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RobotHappyFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.078 8.347a1.368 1.368 0 0 0-.488-.325V4.647a.717.717 0 1 0-1.434 0V7.85h-.21a5.479 5.479 0 0 0-5.25-3.92H9.427a5.48 5.48 0 0 0-5.25 3.92H3.9V4.647a.717.717 0 1 0-1.434 0v3.385a1.501 1.501 0 0 0-.469.315A1.722 1.722 0 0 0 1.5 9.552v4.896a1.702 1.702 0 0 0 1.702 1.702h.956a5.48 5.48 0 0 0 5.25 3.92h5.183a5.48 5.48 0 0 0 5.25-3.92h.955a1.702 1.702 0 0 0 1.702-1.702V9.552c.02-.44-.131-.872-.42-1.205M3.996 14.716H3.24a.267.267 0 0 1-.191-.077a.287.287 0 0 1-.076-.191V9.552a.258.258 0 0 1 .248-.268h.775a.545.545 0 0 0 0 .125v5.182a.545.545 0 0 0 0 .125m4.695-3.118a.813.813 0 0 1-1.386-.578c0-.217.086-.425.238-.579l.956-.956a.813.813 0 0 1 1.148 0l.956.956a.812.812 0 0 1-.574 1.387a.813.813 0 0 1-.573-.23l-.412-.41zm5.9 4.074a3.605 3.605 0 0 1-5.068 0a.813.813 0 0 1 .885-1.326a.804.804 0 0 1 .262.178a2.017 2.017 0 0 0 2.773 0a.804.804 0 0 1 1.148 0a.813.813 0 0 1 0 1.148m1.912-4.074a.813.813 0 0 1-1.148 0l-.41-.41l-.402.41a.822.822 0 0 1-.574.23a.812.812 0 0 1-.574-.23a.822.822 0 0 1 0-1.157l.957-.956a.813.813 0 0 1 1.147 0l.956.956a.822.822 0 0 1 .077 1.157zm4.609 2.869a.287.287 0 0 1-.077.191a.268.268 0 0 1-.191.077h-.755a.55.55 0 0 0 0-.125V9.37a.55.55 0 0 0 0-.124h.765a.248.248 0 0 1 .181.077c.049.052.076.12.077.19z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RobotSad(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M14.706 4.313H9.294a4.982 4.982 0 0 0-4.982 4.981v5.412a4.982 4.982 0 0 0 4.982 4.982h5.412a4.982 4.982 0 0 0 4.982-4.982V9.294a4.982 4.982 0 0 0-4.982-4.982Z"/><path d="M19.606 15.588h1.619a1.025 1.025 0 0 0 1.025-1.025V9.438a1.025 1.025 0 0 0-1.025-1.025h-1.62m-15.21 7.175h-1.62a1.025 1.025 0 0 1-1.025-1.025V9.438a1.025 1.025 0 0 1 1.025-1.025h1.62"/><path stroke-linecap="round" stroke-linejoin="round" d="M2.765 8.413v-4.1m18.46 4.1l-.01-4.1m-12.3 4.45v2.698m6.15-2.698v2.698M9.95 15.588a2.89 2.89 0 0 1 4.1 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RobotSadFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.078 8.347a1.368 1.368 0 0 0-.488-.325V4.647a.717.717 0 1 0-1.434 0V7.85h-.21a5.479 5.479 0 0 0-5.25-3.92H9.427a5.48 5.48 0 0 0-5.25 3.92H3.9V4.647a.717.717 0 1 0-1.434 0v3.385a1.501 1.501 0 0 0-.469.315A1.722 1.722 0 0 0 1.5 9.552v4.896a1.702 1.702 0 0 0 1.702 1.702h.956a5.48 5.48 0 0 0 5.25 3.92h5.183a5.48 5.48 0 0 0 5.25-3.92h.955a1.702 1.702 0 0 0 1.702-1.702V9.552c.02-.44-.131-.872-.42-1.205M3.996 14.716H3.24a.267.267 0 0 1-.191-.077a.287.287 0 0 1-.076-.191V9.552a.258.258 0 0 1 .248-.268h.775a.545.545 0 0 0 0 .125v5.182a.545.545 0 0 0 0 .125m4.303-3.232V8.902a.813.813 0 1 1 1.616 0v2.582a.813.813 0 1 1-1.616 0m6.292 4.522a.794.794 0 0 1-.574.24a.823.823 0 0 1-.574-.24a1.912 1.912 0 0 0-.64-.43a1.913 1.913 0 0 0-2.133.43a.813.813 0 1 1-1.147-1.147c.331-.334.724-.6 1.157-.784a3.576 3.576 0 0 1 3.91.784a.813.813 0 0 1 0 1.147m1.204-4.522a.813.813 0 0 1-1.625 0V8.902a.813.813 0 0 1 1.625 0zm5.345 2.964a.287.287 0 0 1-.076.191a.265.265 0 0 1-.191.077h-.756a.55.55 0 0 0 0-.125v-5.22a.55.55 0 0 0 0-.125h.765a.248.248 0 0 1 .182.077c.048.052.076.12.076.19z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RobotScreen(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M14.706 4.313H9.294a4.982 4.982 0 0 0-4.982 4.981v5.412a4.982 4.982 0 0 0 4.982 4.982h5.412a4.982 4.982 0 0 0 4.982-4.982V9.294a4.982 4.982 0 0 0-4.982-4.982Z"/><path d="M19.606 15.588h1.619a1.025 1.025 0 0 0 1.025-1.025V9.438a1.025 1.025 0 0 0-1.025-1.025h-1.62m-15.21 7.175h-1.62a1.025 1.025 0 0 1-1.025-1.025V9.438a1.025 1.025 0 0 1 1.025-1.025h1.62"/><path stroke-linecap="round" stroke-linejoin="round" d="M2.765 8.413v-4.1m18.46 4.1l-.01-4.1"/><path d="M13.507 7.704h-3.014a2.79 2.79 0 0 0-2.789 2.79v3.013a2.79 2.79 0 0 0 2.79 2.789h3.013a2.79 2.79 0 0 0 2.789-2.79v-3.013a2.79 2.79 0 0 0-2.79-2.789Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RobotScreenFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.078 8.347a1.368 1.368 0 0 0-.488-.325V4.647a.717.717 0 1 0-1.434 0V7.85h-.21a5.479 5.479 0 0 0-5.25-3.92H9.427a5.48 5.48 0 0 0-5.25 3.92H3.9V4.647a.717.717 0 1 0-1.434 0v3.385a1.501 1.501 0 0 0-.469.315A1.722 1.722 0 0 0 1.5 9.552v4.896a1.702 1.702 0 0 0 1.702 1.702h.956a5.48 5.48 0 0 0 5.25 3.92h5.183a5.48 5.48 0 0 0 5.25-3.92h.955a1.702 1.702 0 0 0 1.702-1.702V9.552c.02-.44-.131-.872-.42-1.205M3.996 14.716H3.24a.267.267 0 0 1-.191-.077a.287.287 0 0 1-.076-.191V9.552a.258.258 0 0 1 .248-.268h.775a.545.545 0 0 0 0 .125v5.182a.545.545 0 0 0 0 .125m12.89-1.272a3.394 3.394 0 0 1-3.386 3.385h-2.868a3.394 3.394 0 0 1-3.385-3.385v-2.869a3.395 3.395 0 0 1 3.385-3.385H13.5a3.395 3.395 0 0 1 3.385 3.385zm4.255.956a.287.287 0 0 1-.077.191a.265.265 0 0 1-.191.077h-.756a.55.55 0 0 0 0-.124V9.37a.55.55 0 0 0 0-.124h.765a.248.248 0 0 1 .182.077c.048.052.076.12.076.19z"/><path fill="currentColor" d="M13.5 8.605h-2.887a1.95 1.95 0 0 0-1.951 1.951v2.888a1.95 1.95 0 0 0 1.95 1.95H13.5a1.95 1.95 0 0 0 1.951-1.95v-2.888a1.95 1.95 0 0 0-1.95-1.95"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RobotUwu(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M14.706 4.313H9.294a4.981 4.981 0 0 0-4.982 4.981v5.412a4.982 4.982 0 0 0 4.982 4.982h5.412a4.982 4.982 0 0 0 4.982-4.982V9.294a4.982 4.982 0 0 0-4.982-4.982Z"/><path d="M19.606 15.588h1.619a1.025 1.025 0 0 0 1.025-1.025V9.438a1.025 1.025 0 0 0-1.025-1.025h-1.62m-15.21 7.175h-1.62a1.025 1.025 0 0 1-1.025-1.025V9.438a1.025 1.025 0 0 1 1.025-1.025h1.62"/><path stroke-linecap="round" stroke-linejoin="round" d="M2.765 8.413v-4.1m18.46 4.1l-.01-4.1M14.05 9.704c0 2.296 2.87 2.296 2.87 0m-10.045 0c0 2.296 2.87 2.296 2.87 0m-.328 5.494c0 1.722 2.583 1.302 2.583.39c0 .912 2.583 1.332 2.583-.39"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RobotUwuFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.078 8.347a1.368 1.368 0 0 0-.488-.325V4.647a.717.717 0 1 0-1.434 0V7.85h-.21a5.479 5.479 0 0 0-5.25-3.92H9.427a5.48 5.48 0 0 0-5.25 3.92H3.9V4.647a.717.717 0 1 0-1.434 0v3.385a1.501 1.501 0 0 0-.469.315A1.722 1.722 0 0 0 1.5 9.552v4.896a1.702 1.702 0 0 0 1.702 1.702h.956a5.48 5.48 0 0 0 5.25 3.92h5.183a5.48 5.48 0 0 0 5.25-3.92h.955a1.702 1.702 0 0 0 1.702-1.702V9.552c.02-.44-.131-.872-.42-1.205M3.996 14.716H3.24a.267.267 0 0 1-.191-.077a.287.287 0 0 1-.076-.191V9.552a.258.258 0 0 1 .248-.268h.775a.545.545 0 0 0 0 .125v5.182a.545.545 0 0 0 0 .125m3.07-3.184A2.448 2.448 0 0 1 6.443 9.8a.717.717 0 0 1 .707-.717a.727.727 0 0 1 .727.717c0 .688.344.956.65.956c.306 0 .66-.249.66-.956a.717.717 0 0 1 1.434 0a2.162 2.162 0 0 1-2.094 2.371a2.008 2.008 0 0 1-1.463-.64m7.525 4.962a2.085 2.085 0 0 1-1.291.411a2.486 2.486 0 0 1-1.244-.325a2.409 2.409 0 0 1-2.533-.086a1.779 1.779 0 0 1-.65-1.434a.717.717 0 1 1 1.434 0c0 .239.076.315.22.363a1.031 1.031 0 0 0 .822-.095a.717.717 0 0 1 1.415 0a1.034 1.034 0 0 0 .822.095c.144-.048.22-.124.22-.363a.717.717 0 0 1 1.435 0a1.78 1.78 0 0 1-.65 1.434m.803-4.322a2.047 2.047 0 0 1-1.473-.64A2.448 2.448 0 0 1 13.3 9.8a.717.717 0 1 1 1.434 0c0 .688.354.956.66.956c.306 0 .65-.249.65-.956a.717.717 0 1 1 1.434 0a2.161 2.161 0 0 1-2.084 2.371m5.737 2.276a.287.287 0 0 1-.076.191a.267.267 0 0 1-.192.077h-.755a.538.538 0 0 0 0-.125V9.37a.538.538 0 0 0 0-.125h.765a.249.249 0 0 1 .181.077c.049.052.076.12.077.19z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RobotWink(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M14.706 4.313H9.294a4.981 4.981 0 0 0-4.982 4.981v5.412a4.982 4.982 0 0 0 4.982 4.982h5.412a4.982 4.982 0 0 0 4.982-4.982V9.294a4.982 4.982 0 0 0-4.982-4.982Z"/><path d="M19.606 15.588h1.619a1.025 1.025 0 0 0 1.025-1.025V9.438a1.025 1.025 0 0 0-1.025-1.025h-1.62m-15.21 7.175h-1.62a1.025 1.025 0 0 1-1.025-1.025V9.438a1.025 1.025 0 0 1 1.025-1.025h1.62"/><path stroke-linecap="round" stroke-linejoin="round" d="M2.765 8.413v-4.1m18.46 4.1l-.01-4.1m-7.185 6.662h2.049M8.925 9.796v2.05m1.025 3.229a2.911 2.911 0 0 0 4.1 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RobotWinkFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.078 8.347a1.368 1.368 0 0 0-.488-.325V4.647a.717.717 0 1 0-1.434 0V7.85h-.21a5.479 5.479 0 0 0-5.25-3.92H9.427a5.48 5.48 0 0 0-5.25 3.92H3.9V4.647a.717.717 0 1 0-1.434 0v3.385a1.501 1.501 0 0 0-.469.315A1.722 1.722 0 0 0 1.5 9.552v4.896a1.702 1.702 0 0 0 1.702 1.702h.956a5.48 5.48 0 0 0 5.25 3.92h5.183a5.48 5.48 0 0 0 5.25-3.92h.955a1.702 1.702 0 0 0 1.702-1.702V9.552c.02-.44-.131-.872-.42-1.205M3.996 14.716H3.24a.267.267 0 0 1-.191-.077a.287.287 0 0 1-.076-.191V9.552a.258.258 0 0 1 .248-.268h.775a.545.545 0 0 0 0 .125v5.182a.545.545 0 0 0 0 .125m4.312-2.869v-1.96a.813.813 0 1 1 1.616 0v1.96a.813.813 0 1 1-1.616 0m6.283 3.662a3.605 3.605 0 0 1-5.068 0a.813.813 0 0 1 .885-1.326a.804.804 0 0 1 .262.179a2.017 2.017 0 0 0 2.773 0a.804.804 0 0 1 1.148 0a.813.813 0 0 1 0 1.157zm1.367-3.69h-1.913a.812.812 0 0 1-.574-1.385a.803.803 0 0 1 .574-.232h1.913a.805.805 0 0 1 .754 1.117a.812.812 0 0 1-.754.509zm5.182 2.62a.287.287 0 0 1-.076.19a.267.267 0 0 1-.191.077h-.756a.55.55 0 0 0 0-.124V9.37a.55.55 0 0 0 0-.124h.765a.248.248 0 0 1 .182.077c.048.052.076.12.076.19z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rocket(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.25 4.527a1.774 1.774 0 0 0-1.81-1.777a6.943 6.943 0 0 0-2.888.61a5.549 5.549 0 0 0-1.471.938l-7.695 7.706a3.27 3.27 0 0 0 0 4.621a3.27 3.27 0 0 0 4.62 0l7.707-7.706a5.55 5.55 0 0 0 .937-1.504a6.92 6.92 0 0 0 .6-2.888"/><path d="M14.775 4.603a17.298 17.298 0 0 0 4.665 4.622m-9.744.468l-2.43-.49a2.758 2.758 0 0 0-2.464.752L2.917 11.84a.545.545 0 0 0 .381.937h3.27m7.75 1.538l.48 2.42a2.714 2.714 0 0 1-.742 2.463l-1.896 1.896a.545.545 0 0 1-.927-.392v-3.27m-6.169 1.504l6.943-6.932"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RocketFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.818 3.55a2.593 2.593 0 0 0-1.394-1.37A2.534 2.534 0 0 0 19.49 2h-.08a7.801 7.801 0 0 0-3.158.68c-.624.268-1.2.636-1.705 1.09l-.27.28l-.07.06l-4.774 4.768l-2.005-.41a3.48 3.48 0 0 0-1.695.09a3.672 3.672 0 0 0-1.454.87l-1.875 1.869a1.289 1.289 0 0 0-.37.66c-.062.253-.04.519.06.76c.099.238.266.443.48.589c.213.144.465.221.723.22h2.476a3.99 3.99 0 0 0 .612 3.049l-1.845 1.83a.749.749 0 0 0 .532 1.279a.734.734 0 0 0 .53-.22l1.846-1.83a3.94 3.94 0 0 0 2.266.7c.266 0 .531-.027.792-.08v2.44a1.316 1.316 0 0 0 .792 1.199c.162.069.336.102.512.1c.08.01.16.01.24 0a1.36 1.36 0 0 0 .672-.35l1.895-1.9c.412-.4.712-.899.873-1.449a3.448 3.448 0 0 0 .08-1.69l-.401-1.999l4.733-4.718l.21-.2l.21-.21a6.216 6.216 0 0 0 1.003-1.7a7.588 7.588 0 0 0 .672-3.188a2.523 2.523 0 0 0-.18-1.04M7.9 16.094a3.048 3.048 0 0 1-.461-.61l4.01-3.998a.753.753 0 0 1 1.284.53c0 .198-.08.389-.22.53l-4.011 3.998a2.435 2.435 0 0 1-.602-.45M20.514 4.539a6.28 6.28 0 0 1-.541 2.59a4.552 4.552 0 0 1-.652 1.089a16.662 16.662 0 0 1-3.56-3.509c.336-.27.707-.492 1.103-.66a6.234 6.234 0 0 1 2.567-.55c.137.007.272.034.401.08c.128.05.244.125.341.22c.1.092.179.205.23.33c.052.127.08.263.081.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RoundSticker(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.5 12L12 21.5a9.5 9.5 0 1 1 9.5-9.5"/><path d="M12 21.5a7.905 7.905 0 0 1-.56-3.272a6.787 6.787 0 0 1 6.788-6.756a8.053 8.053 0 0 1 3.272.56"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RoundStickerFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.555 6.305A10.25 10.25 0 1 0 4.093 18.483a10.26 10.26 0 0 0 7.882 3.772c.097 0 .193-.02.28-.06a.718.718 0 0 0 .25-.16l9.5-9.5a.79.79 0 0 0 .21-.39a.608.608 0 0 0 0-.14a10.21 10.21 0 0 0-1.66-5.7m-8.17 13.78a7.134 7.134 0 0 1-.17-1.85a6.001 6.001 0 0 1 6.07-6a7.232 7.232 0 0 1 1.8.16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SaveFloppy(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.25 9.16v7.987a4.1 4.1 0 0 1-1.204 2.901a4.113 4.113 0 0 1-2.906 1.202H6.86a4.113 4.113 0 0 1-2.906-1.202a4.1 4.1 0 0 1-1.204-2.901V6.853a4.1 4.1 0 0 1 1.204-2.901A4.113 4.113 0 0 1 6.86 2.75h8.35a3.004 3.004 0 0 1 2.25.998l3 3.415c.501.545.783 1.256.79 1.997"/><path d="M7 21.22v-5.241a1.995 1.995 0 0 1 2-1.997h6a2.002 2.002 0 0 1 2 1.997v5.241M15.8 2.81v4.183a1.526 1.526 0 0 1-1.52 1.528H9.72A1.531 1.531 0 0 1 8.2 6.993V2.75m1.946 15.108h3.708"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SaveFloppyFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.01 6.62l-3-3.42a3.77 3.77 0 0 0-1.27-.92a3.38 3.38 0 0 0-.75-.22a.41.41 0 0 0-.19 0c-.2 0-.39-.05-.59-.06H6.86a4.82 4.82 0 0 0-3.44 1.42A4.89 4.89 0 0 0 2 6.85v10.29A4.85 4.85 0 0 0 6.86 22h10.28A4.87 4.87 0 0 0 22 17.14v-8a3.77 3.77 0 0 0-.99-2.52M8.94 3.42h6.1v3.49a.76.76 0 0 1-.23.55a.81.81 0 0 1-.54.24H9.71a.82.82 0 0 1-.55-.24a.8.8 0 0 1-.22-.55zm8.8 16.94c-.2.04-.405.06-.61.06h-.89V15.9a1.26 1.26 0 0 0-1.25-1.25h-6a1.25 1.25 0 0 0-1.25 1.25v4.52h-.89a4.059 4.059 0 0 1-.61 0v-4.47a2.75 2.75 0 0 1 2.75-2.75h6a2.77 2.77 0 0 1 2.75 2.75zm-8.36-2.59a.76.76 0 0 1 .75-.75h3.71a.76.76 0 0 1 .696 1.039a.741.741 0 0 1-.696.461h-3.71a.74.74 0 0 1-.75-.71z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScaleDown(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.2 19.404V13.8H4.596m14.808-3.6H13.8V4.596M21 3l-7.2 7.2m-3.6 3.6L3 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScaleUp(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 15.396V21h5.604m6.792-18H21v5.604M21 3l-7.2 7.2m-3.6 3.6L3 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Scan(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20.325 8.3V6.45a2.775 2.775 0 0 0-2.775-2.775h-2.775m0 16.65h2.775a2.775 2.775 0 0 0 2.775-2.775V15.7m-16.65 0v1.85a2.775 2.775 0 0 0 2.775 2.775h2.775m0-16.65H6.45A2.775 2.775 0 0 0 3.675 6.45V8.3M2.75 12h18.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScanUser(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.25 8.814V6.758a3.083 3.083 0 0 0-3.083-3.083h-3.084m0 18.5h3.084a3.083 3.083 0 0 0 3.083-3.083v-2.056m-18.5 0v2.056a3.083 3.083 0 0 0 3.083 3.083h3.084m0-18.5H5.833A3.083 3.083 0 0 0 2.75 6.758v2.056"/><path d="M18.177 22.175c0-2.92-3.256-5.294-6.177-5.294c-2.92 0-6.176 2.373-6.176 5.294M12 14.234a3.53 3.53 0 1 0 0-7.06a3.53 3.53 0 0 0 0 7.06"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScanUserFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.221 8.623a.75.75 0 0 1-.749-.748V5.83a2.333 2.333 0 0 0-2.326-2.334H15.1A.75.75 0 1 1 15.1 2h3.076a3.827 3.827 0 0 1 2.706 1.127A3.75 3.75 0 0 1 22 5.83v2.045a.738.738 0 0 1-.779.748m-18.472 0A.75.75 0 0 1 2 7.875V5.83a3.757 3.757 0 0 1 1.118-2.703A3.824 3.824 0 0 1 5.824 2H8.9a.75.75 0 1 1 0 1.496H5.824A2.358 2.358 0 0 0 3.488 5.83v2.045a.738.738 0 0 1-.739.748m19.231 7.452v2.105a3.798 3.798 0 0 1-1.128 2.703A3.876 3.876 0 0 1 18.166 22H5.814a3.876 3.876 0 0 1-2.696-1.167A3.828 3.828 0 0 1 2.01 18.18v-2.105a.748.748 0 0 1 1.498 0v2.105a2.313 2.313 0 0 0 1.607 2.174c.55-2.992 3.905-5.187 6.84-5.187c2.936 0 6.28 2.205 6.83 5.187a2.28 2.28 0 0 0 .998-.578a2.326 2.326 0 0 0 .69-1.646v-2.055a.748.748 0 0 1 1.497 0zm-6.1-6.315a3.888 3.888 0 0 1-2.4 3.603a3.897 3.897 0 0 1-5.315-2.836a3.887 3.887 0 0 1 1.663-3.995a3.897 3.897 0 0 1 4.91.488a3.888 3.888 0 0 1 1.141 2.74"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Screencast(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4.6 20.325a1.85 1.85 0 0 0-1.85-1.85m4.815 1.85A4.815 4.815 0 0 0 2.75 15.51m7.949 4.815a7.949 7.949 0 0 0-7.949-7.949"/><path d="M13.925 20.325h3.625a3.7 3.7 0 0 0 3.7-3.7v-9.25a3.7 3.7 0 0 0-3.7-3.7H6.45a3.7 3.7 0 0 0-3.7 3.7V9.15"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScreencastFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<defs><path id="mageScreencastFill0" fill="currentColor" d="M10.691 21.07a.74.74 0 0 1-.749-.748a7.193 7.193 0 0 0-7.193-7.193a.75.75 0 1 1 0-1.499a8.701 8.701 0 0 1 8.692 8.692a.75.75 0 0 1-.75.749"/></defs><path fill="currentColor" d="M4.597 21.071a.74.74 0 0 1-.749-.75a1.109 1.109 0 0 0-1.099-1.098a.75.75 0 1 1 0-1.499a2.607 2.607 0 0 1 2.598 2.598a.75.75 0 0 1-.75.749"/><path fill="currentColor" d="M7.554 21.07a.75.75 0 0 1-.749-.748a4.066 4.066 0 0 0-4.056-4.056a.75.75 0 1 1 0-1.499c.73 0 1.453.142 2.128.42a5.634 5.634 0 0 1 2.997 2.997c.278.675.42 1.398.42 2.128a.74.74 0 0 1-.74.759"/><use href="#mageScreencastFill0"/><path fill="currentColor" d="M21.98 7.375v9.24a4.455 4.455 0 0 1-4.445 4.446h-3.627a.739.739 0 0 1-.74-.84a9.99 9.99 0 0 0-2.996-7.672a8.371 8.371 0 0 0-7.253-2.667a.73.73 0 0 1-.62-.15a.74.74 0 0 1-.28-.58V7.376A4.456 4.456 0 0 1 6.466 2.93h11.089A4.465 4.465 0 0 1 22 7.375z"/><use href="#mageScreencastFill0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Search(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.783 18.828a8.046 8.046 0 0 0 7.439-4.955a8.034 8.034 0 0 0-1.737-8.765a8.045 8.045 0 0 0-13.735 5.68c0 2.131.846 4.174 2.352 5.681a8.046 8.046 0 0 0 5.68 2.359m5.706-2.337l4.762 4.759"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.78 21.79a.82.82 0 0 1-.53.21a.8.8 0 0 1-.53-.21l-4.273-4.28a7.253 7.253 0 0 1-.78.59a8.84 8.84 0 0 1-4.883 1.48a8.71 8.71 0 0 1-3.372-.67A8.525 8.525 0 0 1 4.57 17A8.788 8.788 0 0 1 7.04 2.842a8.8 8.8 0 0 1 5.475-.672a8.788 8.788 0 0 1 6.05 4.482a8.775 8.775 0 0 1 .354 7.518a8.482 8.482 0 0 1-1.401 2.29l4.272 4.28a.75.75 0 0 1-.01 1.05"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M10.82 15.857a5.038 5.038 0 1 0 0-10.076a5.038 5.038 0 0 0 0 10.076m3.575-1.462l3.197 3.197"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SearchSquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 2h-6.5A6.76 6.76 0 0 0 2 8.75v6.5A6.76 6.76 0 0 0 8.75 22h6.5A6.76 6.76 0 0 0 22 15.25v-6.5A6.76 6.76 0 0 0 15.25 2m2 14.94a.78.78 0 0 1-.57.23a.8.8 0 0 1-.56-.23l-2-2a4.81 4.81 0 1 1 1.13-1.13l2 2a.79.79 0 0 1 .01 1.13z"/><path fill="currentColor" d="M14.5 11a3.21 3.21 0 1 1-6.418 0a3.21 3.21 0 0 1 6.418 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SecurityShield(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 11.543A2.17 2.17 0 1 0 12 7.2a2.17 2.17 0 0 0 0 4.342m0 .001v3.256"/><path d="M20.672 11.89V6.61a1.928 1.928 0 0 0-1.32-1.831L14.438 3.14a7.805 7.805 0 0 0-4.876 0L4.648 4.778a1.927 1.927 0 0 0-1.32 1.83v5.28a7.709 7.709 0 0 0 3.603 6.524l4.048 2.544a1.927 1.927 0 0 0 2.042 0l4.047-2.544a7.708 7.708 0 0 0 3.604-6.523"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SecurityShieldFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.872 5.05a2.824 2.824 0 0 0-1.402-1.026l-4.801-1.6a8.68 8.68 0 0 0-5.357 0l-4.8 1.6A2.824 2.824 0 0 0 3.117 5.05a2.824 2.824 0 0 0-.536 1.657v5.158a8.472 8.472 0 0 0 3.972 7.22l3.954 2.486a2.823 2.823 0 0 0 2.993 0l3.954-2.485a8.473 8.473 0 0 0 3.963-7.174V6.754a2.824 2.824 0 0 0-.546-1.704m-7.935 7.286v2.269a.941.941 0 0 1-1.883 0v-2.269a2.984 2.984 0 1 1 1.883 0"/><path fill="currentColor" d="M13.101 9.502a1.101 1.101 0 1 1-2.202 0a1.101 1.101 0 0 1 2.202 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SelectBox(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M2.75 6.861c0-3.402.71-4.111 4.111-4.111M2.75 14.056V9.944M6.861 21.25c-3.402 0-4.111-.71-4.111-4.111m11.306 4.111H9.944m11.306-4.111c0 3.402-.71 4.111-4.111 4.111M21.25 9.944v4.112M17.139 2.75c3.402 0 4.111.71 4.111 4.111M9.944 2.75h4.112"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Server(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3.002 5.127l.011 5.715a1.152 1.152 0 0 0 1.155 1.15l15.68-.03a1.152 1.152 0 0 0 1.15-1.155l-.011-5.714a1.152 1.152 0 0 0-1.154-1.15l-15.68.03a1.152 1.152 0 0 0-1.15 1.154m14.548 2.835h-3.018m-10.37 4.032c-.636 0-1.152.516-1.152 1.152v5.761c0 .636.516 1.152 1.152 1.152h15.68c.637 0 1.153-.516 1.153-1.152v-5.76c0-.637-.516-1.153-1.152-1.153m-2.293 4.033h-3.018"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ServerFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.735 10.82V5.11a1.889 1.889 0 0 0-1.9-1.9H4.155a1.86 1.86 0 0 0-1.34.56a1.89 1.89 0 0 0-.56 1.34v5.72c.001.417.142.822.4 1.15a1.81 1.81 0 0 0-.4 1.15v5.76a1.9 1.9 0 0 0 1.9 1.9h15.68a1.91 1.91 0 0 0 1.91-1.9v-5.76a1.87 1.87 0 0 1 0-2.34zm-4.2 6.22h-3a1 1 0 1 1 0-2h3a1 1 0 1 1 0 2m0-8.06h-3a1 1 0 0 1 0-2h3a1 1 0 1 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ServerTwo(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4.159 9.134c-.637 0-1.152.516-1.152 1.152v3.456c0 .636.515 1.152 1.152 1.152h15.678c.636 0 1.152-.515 1.152-1.152v-3.456c0-.636-.516-1.152-1.152-1.152m-2.293 2.88h-3.018"/><path d="m3.004 4.554l.01 3.456c.002.636.52 1.15 1.156 1.149l15.678-.05a1.152 1.152 0 0 0 1.148-1.155l-.01-3.456A1.152 1.152 0 0 0 19.83 3.35l-15.678.049c-.636.002-1.15.52-1.148 1.155m14.54 1.689h-3.018M4.159 14.894c-.637 0-1.152.516-1.152 1.152v3.456c0 .636.515 1.152 1.152 1.152h15.678c.636 0 1.152-.516 1.152-1.152v-3.456c0-.636-.516-1.152-1.152-1.152m-2.293 2.88h-3.018"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ServerTwoFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.53 13.04h-3a1 1 0 1 1 0-2h3a1 1 0 0 1 0 2"/><path fill="currentColor" d="M21.73 7.98V4.52a1.86 1.86 0 0 0-.56-1.34a1.89 1.89 0 0 0-1.34-.56H4.15a1.88 1.88 0 0 0-1.34.56a1.86 1.86 0 0 0-.55 1.31v3.5a1.87 1.87 0 0 1 0 2.27v3.46a1.85 1.85 0 0 1 0 2.3v3.46a1.91 1.91 0 0 0 1.9 1.9h15.68a1.91 1.91 0 0 0 1.9-1.9v-3.46a1.85 1.85 0 0 1 0-2.3v-3.41a1.82 1.82 0 0 0-.41-1.16a1.84 1.84 0 0 0 .4-1.17m-7.22-2.71h3a1 1 0 0 1 0 2h-3a1 1 0 1 1 0-2m3 13.53h-3a1 1 0 0 1 0-2h3a1 1 0 0 1 0 2m2.69-5a.4.4 0 0 1-.4.4H4.14a.4.4 0 0 1-.4-.4v-3.49a.39.39 0 0 1 .34-.39h.07l15.49-.05h.18a.4.4 0 0 1 .4.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Settings(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12.132 15.404a3.364 3.364 0 1 0 0-6.728a3.364 3.364 0 0 0 0 6.728"/><path d="M20.983 15.094a9.43 9.43 0 0 1-1.802 3.1l-2.124-.482a7.245 7.245 0 0 1-2.801 1.56l-.574 2.079a9.462 9.462 0 0 1-1.63.149a9.117 9.117 0 0 1-2.032-.23l-.609-2.146a7.475 7.475 0 0 1-2.457-1.493l-2.1.54a9.357 9.357 0 0 1-1.837-3.33l1.55-1.722a7.186 7.186 0 0 1 .069-2.652L3.107 8.872a9.356 9.356 0 0 1 2.067-3.353l2.17.54A7.68 7.68 0 0 1 9.319 4.91l.574-2.124a8.886 8.886 0 0 1 2.17-.287c.585 0 1.17.054 1.745.16l.551 2.113c.83.269 1.608.68 2.296 1.217l2.182-.563a9.368 9.368 0 0 1 2.043 3.1l-1.48 1.607a7.405 7.405 0 0 1 .068 3.364z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SettingsFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.51 14.59l-1.25-1.32a7.878 7.878 0 0 0-.06-2.9l1.22-1.32a.76.76 0 0 0 .14-.79a10.257 10.257 0 0 0-2.2-3.35a.74.74 0 0 0-.72-.19l-1.84.47a8.48 8.48 0 0 0-1.83-1l-.45-1.72a.73.73 0 0 0-.59-.55a9.92 9.92 0 0 0-1.89-.17a9.36 9.36 0 0 0-2.35.31a.73.73 0 0 0-.53.53l-.48 1.77a8.23 8.23 0 0 0-1.52.88l-1.82-.45a.73.73 0 0 0-.72.21a10 10 0 0 0-2.23 3.62a.76.76 0 0 0 .16.77l1.26 1.31a8.85 8.85 0 0 0-.1 1.27c0 .3 0 .6.05.9l-1.31 1.46a.75.75 0 0 0-.16.73a10 10 0 0 0 2 3.59a.75.75 0 0 0 .76.24l1.72-.44a7.918 7.918 0 0 0 2 1.23l.5 1.79a.77.77 0 0 0 .56.53c.721.163 1.459.247 2.2.25c.59-.006 1.178-.063 1.76-.17a.75.75 0 0 0 .59-.53l.47-1.69a8.109 8.109 0 0 0 2.38-1.34l1.76.4a.74.74 0 0 0 .73-.24a10.118 10.118 0 0 0 2-3.34a.76.76 0 0 0-.21-.75m-9.39 1.27a3.81 3.81 0 1 1-.021-7.619a3.81 3.81 0 0 1 .02 7.62"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Share(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.015 15.809a3.265 3.265 0 1 0 0-6.53a3.265 3.265 0 0 0 0 6.53m11.97-6.529a3.265 3.265 0 1 0 0-6.53a3.265 3.265 0 0 0 0 6.53m0 11.97a3.265 3.265 0 1 0 0-6.53a3.265 3.265 0 0 0 0 6.53m-2.971-4.614l-6.028-2.742m6.126-6.312l-6.224 3.395"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.96 17.977a4 4 0 1 1-7.94-.7l-4.9-2.22a4.001 4.001 0 0 1-3.11 1.49a4 4 0 1 1 0-8a4 4 0 0 1 2.92 1.28l5.16-2.8a3.582 3.582 0 0 1-.13-1a4 4 0 1 1 1.09 2.74l-5.16 2.81c.09.326.134.662.13 1a4 4 0 0 1-.06.7l4.9 2.22a4 4 0 0 1 7.13 2.52z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldCheck(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8.67 10.909l1.875 1.874a.708.708 0 0 0 1.008 0l3.777-3.777"/><path d="M20.672 11.89V6.61a1.928 1.928 0 0 0-1.32-1.831L14.438 3.14a7.805 7.805 0 0 0-4.876 0L4.648 4.778a1.927 1.927 0 0 0-1.32 1.83v5.28a7.709 7.709 0 0 0 3.603 6.524l4.048 2.544a1.927 1.927 0 0 0 2.042 0l4.047-2.544a7.708 7.708 0 0 0 3.604-6.523"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldCheckFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.92 5.037a2.629 2.629 0 0 0-1.314-.966l-4.939-1.643a8.563 8.563 0 0 0-5.344 0L4.393 4.07a2.658 2.658 0 0 0-1.816 2.522v5.297a8.525 8.525 0 0 0 1.053 4.088a8.399 8.399 0 0 0 2.9 3.064l4.059 2.561a2.707 2.707 0 0 0 2.822 0l4.06-2.561a8.438 8.438 0 0 0 3.952-7.152V6.593a2.652 2.652 0 0 0-.502-1.556m-5.025 4.755l-3.673 3.644a1.809 1.809 0 0 1-.531.367c-.204.08-.42.123-.638.126a1.787 1.787 0 0 1-.638-.126a1.73 1.73 0 0 1-.541-.367L8.066 11.63a.966.966 0 1 1 1.373-1.363l1.614 1.614l3.45-3.46a.967.967 0 0 1 1.363 1.372z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldCross(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="m14.51 8.485l-5.02 5.019m0-5.008l5.02 5.019"/><path stroke-linejoin="round" d="M20.672 11.89V6.61a1.928 1.928 0 0 0-1.32-1.831L14.438 3.14a7.805 7.805 0 0 0-4.876 0L4.648 4.778a1.927 1.927 0 0 0-1.32 1.83v5.28a7.709 7.709 0 0 0 3.603 6.524l4.048 2.544a1.927 1.927 0 0 0 2.042 0l4.047-2.544a7.708 7.708 0 0 0 3.604-6.523"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldCrossFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.92 5.037a2.629 2.629 0 0 0-1.314-.966l-4.939-1.643a8.563 8.563 0 0 0-5.344 0L4.393 4.07a2.658 2.658 0 0 0-1.816 2.522v5.297a8.525 8.525 0 0 0 1.053 4.088a8.399 8.399 0 0 0 2.9 3.064l4.059 2.561a2.707 2.707 0 0 0 2.822 0l4.06-2.561a8.438 8.438 0 0 0 3.952-7.152V6.593a2.652 2.652 0 0 0-.502-1.556m-5.798 7.732a.967.967 0 1 1-1.363 1.373l-1.75-1.75l-1.73 1.74a.966.966 0 0 1-.686.28a.967.967 0 0 1-.686-1.653l1.74-1.73l-1.74-1.74a.97.97 0 0 1 1.373-1.372l1.73 1.74l1.749-1.75a.966.966 0 1 1 1.363 1.373l-1.74 1.75z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldPlus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="M11.992 7.451v7.098m-3.541-3.541h7.098"/><path stroke-linejoin="round" d="M20.672 11.89V6.61a1.928 1.928 0 0 0-1.32-1.831L14.438 3.14a7.805 7.805 0 0 0-4.876 0L4.648 4.778a1.927 1.927 0 0 0-1.32 1.83v5.28a7.709 7.709 0 0 0 3.603 6.524l4.048 2.544a1.927 1.927 0 0 0 2.042 0l4.047-2.544a7.708 7.708 0 0 0 3.604-6.523"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldPlusFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.92 5.037a2.629 2.629 0 0 0-1.314-.966l-4.939-1.643a8.563 8.563 0 0 0-5.344 0L4.393 4.07a2.658 2.658 0 0 0-1.816 2.522v5.297a8.437 8.437 0 0 0 3.953 7.152l4.059 2.561a2.707 2.707 0 0 0 2.822 0l4.06-2.561a8.438 8.438 0 0 0 3.952-7.152V6.593a2.652 2.652 0 0 0-.502-1.556m-5.489 6.969h-2.474v2.464a.966.966 0 1 1-1.933 0v-2.455H8.569a.967.967 0 0 1 0-1.933h2.455V7.608a.967.967 0 0 1 1.933 0v2.465h2.474a.966.966 0 0 1 0 1.933"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldQuestionMark(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="M10.278 8.634a1.83 1.83 0 0 1 1.994-1.037a1.774 1.774 0 0 1 1.301.927a1.542 1.542 0 0 1-.897 2.119a1.137 1.137 0 0 0-.733 1.027v.423"/><path stroke-linejoin="round" d="M11.91 14.433h.005m8.757-2.543V6.61a1.928 1.928 0 0 0-1.32-1.831L14.438 3.14a7.805 7.805 0 0 0-4.876 0L4.648 4.778a1.927 1.927 0 0 0-1.32 1.83v5.28a7.709 7.709 0 0 0 3.603 6.524l4.048 2.544a1.927 1.927 0 0 0 2.042 0l4.047-2.544a7.708 7.708 0 0 0 3.604-6.523"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldQuestionMarkFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.92 5.037a2.629 2.629 0 0 0-1.314-.966l-4.939-1.643a8.563 8.563 0 0 0-5.344 0L4.393 4.07a2.658 2.658 0 0 0-1.816 2.522v5.297a8.437 8.437 0 0 0 3.953 7.152l4.059 2.561a2.707 2.707 0 0 0 2.822 0l4.06-2.561a8.438 8.438 0 0 0 3.952-7.152V6.593a2.652 2.652 0 0 0-.502-1.556m-9.007 10.448a.967.967 0 1 1 .01 0zm2.523-5.248a2.453 2.453 0 0 1-.59.84a2.53 2.53 0 0 1-.89.532v.484a.967.967 0 1 1-1.932 0v-.416a2.088 2.088 0 0 1 1.334-1.933a.454.454 0 0 0 .222-.126a.57.57 0 0 0 .126-.174a.551.551 0 0 0 0-.212a.522.522 0 0 0 0-.213a.657.657 0 0 0-.213-.222a.667.667 0 0 0-.32-.135a.802.802 0 0 0-.908.454a.967.967 0 0 1-1.582.252a.966.966 0 0 1-.167-1.064a2.716 2.716 0 0 1 2.977-1.546c.405.065.788.224 1.121.464c.344.255.625.585.822.966c.155.327.24.682.25 1.044c-.007.348-.09.691-.24 1.005z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShiledMinus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="M8.451 11h7.098"/><path stroke-linejoin="round" d="M20.672 11.89V6.61a1.928 1.928 0 0 0-1.32-1.831L14.438 3.14a7.805 7.805 0 0 0-4.876 0L4.648 4.778a1.927 1.927 0 0 0-1.32 1.83v5.28a7.709 7.709 0 0 0 3.603 6.524l4.048 2.544a1.927 1.927 0 0 0 2.042 0l4.047-2.544a7.708 7.708 0 0 0 3.604-6.523"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShiledMinusFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.92 5.037a2.629 2.629 0 0 0-1.314-.966l-4.939-1.643a8.563 8.563 0 0 0-5.344 0L4.393 4.07a2.658 2.658 0 0 0-1.816 2.522v5.297a8.437 8.437 0 0 0 3.953 7.152l4.059 2.561a2.707 2.707 0 0 0 2.822 0l4.06-2.561a8.438 8.438 0 0 0 3.952-7.152V6.593a2.652 2.652 0 0 0-.502-1.556m-5.489 6.959H8.57a.966.966 0 1 1 0-1.933h6.862a.967.967 0 0 1 0 1.933"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shop(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.25 9.944a3.083 3.083 0 0 1-2.056 2.899a2.847 2.847 0 0 1-1.027.185a3.084 3.084 0 0 1-2.899-2.056a2.848 2.848 0 0 1-.185-1.028c.003.351-.06.7-.185 1.028A3.083 3.083 0 0 1 12 13.028a3.083 3.083 0 0 1-2.898-2.056a2.846 2.846 0 0 1-.185-1.028c.002.351-.06.7-.185 1.028a3.083 3.083 0 0 1-2.899 2.056c-.35.002-.7-.06-1.027-.185A3.084 3.084 0 0 1 2.75 9.944l.462-1.623l1.11-3.166a2.056 2.056 0 0 1 1.943-1.377h11.47a2.056 2.056 0 0 1 1.942 1.377l1.11 3.166z"/><path d="M19.194 12.843v5.324a2.056 2.056 0 0 1-2.055 2.055H6.86a2.055 2.055 0 0 1-2.056-2.055v-5.324m4.113 4.296h6.166"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShopFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.914 9.73l-.48-1.66l-1.11-3.17a2.78 2.78 0 0 0-1-1.36a2.74 2.74 0 0 0-1.62-.52H6.234a2.8 2.8 0 0 0-2.65 1.88l-1.13 3.21l-.46 1.62a.76.76 0 0 0 0 .21a3.85 3.85 0 0 0 2.06 3.39v4.83a2.8 2.8 0 0 0 .82 2a2.84 2.84 0 0 0 2 .82h10.28a2.8 2.8 0 0 0 2.81-2.81v-4.83a3.74 3.74 0 0 0 1.35-1.18a3.79 3.79 0 0 0 .7-2.21a1.482 1.482 0 0 0-.1-.22m-6.89 8.4h-6.17a1 1 0 1 1 0-2h6.17a1 1 0 0 1 0 2m5-6.85c-.282.399-.68.7-1.14.86a2.3 2.3 0 0 1-2.08-.31a2.34 2.34 0 0 1-.99-1.86a.75.75 0 1 0-1.5 0v.05a2.42 2.42 0 0 1-.14.74a2.38 2.38 0 0 1-.86 1.12a2.27 2.27 0 0 1-1.33.43a2.32 2.32 0 0 1-2.2-1.57a2 2 0 0 1-.14-.73a.75.75 0 0 0-1.5 0a2.36 2.36 0 0 1-.99 1.87a2.33 2.33 0 0 1-1.35.43a2.548 2.548 0 0 1-.77-.14a2.28 2.28 0 0 1-1.13-.85a2.33 2.33 0 0 1-.42-1.24l.41-1.48l1.11-3.16a1.31 1.31 0 0 1 1.24-.88h11.47c.27.004.535.088.76.24c.219.16.383.383.47.64l1.1 3.12l.43 1.52a2.35 2.35 0 0 1-.47 1.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBag(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4.484 6.219v12.14a2.89 2.89 0 0 0 2.891 2.891h9.25a2.89 2.89 0 0 0 2.89-2.89V6.218"/><path d="m19.516 6.219l-3.134-3.134a1.157 1.157 0 0 0-.82-.335H8.438a1.156 1.156 0 0 0-.821.335L4.484 6.22zM8.531 9.688l.359.705A2.89 2.89 0 0 0 11.48 12h1.04a2.89 2.89 0 0 0 2.59-1.607l.359-.705"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBagFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.165 6.041a.29.29 0 0 0 0-.1a.357.357 0 0 0 0-.09a.657.657 0 0 0-.12-.17l-3.14-3.13a1.81 1.81 0 0 0-.62-.41a1.49 1.49 0 0 0-.73-.14h-7.12a1.86 1.86 0 0 0-.73.14a1.81 1.81 0 0 0-.62.41l-3.13 3.13a.78.78 0 0 0-.22.52v12.15a3.64 3.64 0 0 0 3.64 3.65h9.25a3.65 3.65 0 0 0 3.64-3.65V6.192a.768.768 0 0 0-.1-.15m-3.89 4.08l-.36.7a3.91 3.91 0 0 1-3.48 2.16h-1a3.91 3.91 0 0 1-2-.59a3.83 3.83 0 0 1-1.43-1.58l-.36-.69a1 1 0 1 1 1.78-.91l.36.71a1.91 1.91 0 0 0 1.7 1.06h1a1.91 1.91 0 0 0 1.7-1.06l.36-.71a1.001 1.001 0 1 1 1.78.91zm-10.06-4.68l1.85-1.85a.41.41 0 0 1 .13-.08a.651.651 0 0 1 .16 0h7.28a.41.41 0 0 1 .13.08l1.86 1.85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingCart(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-width="2.5" d="M7.537 21.25h.01m10.391 0h.011"/><path stroke-width="1.5" d="M4.628 6.526h13.097a4.001 4.001 0 0 1 1.543.237A1.734 1.734 0 0 1 20.29 7.88a3.716 3.716 0 0 1-.213 1.686c-.131.57-.25 1.188-.369 1.71a94.616 94.616 0 0 0-.736 3.562a2.92 2.92 0 0 1-.629 1.567a2.243 2.243 0 0 1-1.686.582H9.805a8.643 8.643 0 0 1-1.662 0a1.637 1.637 0 0 1-1.33-.985a16.623 16.623 0 0 1-.463-2.019c-.095-.534-.214-1.068-.332-1.603c-.416-1.983-.89-3.942-1.39-5.853m0 0C4.32 5.267 4 4.009 3.69 2.75M19.553 12H5.934"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingCartFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.548 22.253a1.25 1.25 0 1 1 0-2.5a1.25 1.25 0 0 1 0 2.5m10.4 0a1.25 1.25 0 1 1 0-2.5a1.25 1.25 0 0 1 0 2.5m3.08-14.7a.682.682 0 0 0 0-.14a2.48 2.48 0 0 0-1.49-1.61a4.68 4.68 0 0 0-1.8-.28H5.228l-.79-3.2a.75.75 0 1 0-1.46.35l.94 3.78c.52 2 .94 3.76 1.3 5.43c0 .14 0 .27.08.4c.12.52.24 1.05.33 1.56c.12.74.29 1.472.51 2.19c.163.396.429.741.77 1a2.48 2.48 0 0 0 1.18.46h.91c.273.015.547.015.82 0h6.73a3.06 3.06 0 0 0 2.26-.78l.08-.08a3.67 3.67 0 0 0 .78-1.9c.16-.93.37-1.88.56-2.77a.43.43 0 0 0 0-.11c.06-.22.1-.44.15-.65c.05-.21.11-.52.17-.79c.06-.27.13-.62.17-.83a4.48 4.48 0 0 0 .31-2.03m-2.81 7a2.06 2.06 0 0 1-.43 1.11a1.44 1.44 0 0 1-.48.28a1.67 1.67 0 0 1-.66.07h-6.92a8.462 8.462 0 0 1-1.51 0a1.001 1.001 0 0 1-.43-.17a.91.91 0 0 1-.26-.29a15.438 15.438 0 0 1-.44-1.93c-.07-.37-.15-.74-.22-1.11h11.75c-.14.63-.28 1.31-.4 2.01z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShutDown(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.743 7.42a5.078 5.078 0 1 0 4.514 0M12 6.362v4.514"/><path d="M12 21.5a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShutDownFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.75A10.25 10.25 0 1 0 22.25 12A10.26 10.26 0 0 0 12 1.75m-.75 4.62a.75.75 0 0 1 1.5 0v4.51a.75.75 0 0 1-1.5 0zm6.42 6.93a5.82 5.82 0 1 1-8.26-6.55a.751.751 0 0 1 .66 1.35a4.3 4.3 0 0 0-2.28 4.86a4.3 4.3 0 0 0 1.52 2.4a4.44 4.44 0 0 0 5.38 0a4.33 4.33 0 0 0-.77-7.26a.753.753 0 1 1 .67-1.35a5.88 5.88 0 0 1 2.68 2.74a5.82 5.82 0 0 1 .36 3.81z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SimCard(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20.444 10.586v6.692a4.222 4.222 0 0 1-4.222 4.222H7.778a4.222 4.222 0 0 1-4.222-4.222V6.722A4.222 4.222 0 0 1 7.778 2.5h4.58a4.223 4.223 0 0 1 2.988 1.235l3.863 3.863a4.222 4.222 0 0 1 1.235 2.988"/><path d="M15.167 9.889H8.833A2.111 2.111 0 0 0 6.723 12v4.222a2.11 2.11 0 0 0 2.11 2.111h6.334a2.111 2.111 0 0 0 2.11-2.11V12a2.111 2.111 0 0 0-2.11-2.111m-3.167 0v8.444m-5.278-4.222h10.556"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SimCardFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.794 8.663a4.987 4.987 0 0 0-1.068-1.616l-3.86-3.85a5.047 5.047 0 0 0-1.606-1.078a4.988 4.988 0 0 0-1.896-.369H7.785a4.988 4.988 0 0 0-4.988 4.988v10.524a4.988 4.988 0 0 0 4.988 4.988h8.43a4.988 4.988 0 0 0 4.988-4.988V10.56a4.798 4.798 0 0 0-.41-1.896m-9.537 8.878H8.843a1.357 1.357 0 0 1-1.357-1.356v-1.357h3.77zm0-4.21h-3.77v-1.356a1.357 1.357 0 0 1 1.356-1.357h2.414zm5.257 2.854a1.357 1.357 0 0 1-1.357 1.357h-2.404v-2.714h3.761zm0-2.853h-3.76v-2.714h2.403a1.357 1.357 0 0 1 1.357 1.357z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Slack(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.923 16.52h-2.39a1.984 1.984 0 0 1-1.973-1.195a2.006 2.006 0 0 1 .47-2.263a1.99 1.99 0 0 1 1.502-.53h4.858a1.978 1.978 0 0 1 1.969 1.63a1.951 1.951 0 0 1-1.147 2.173a2.21 2.21 0 0 1-.876.174c-.8.022-1.601.01-2.413.01m-9.435.501v-2.477a2.003 2.003 0 0 1 .56-1.402a1.987 1.987 0 0 1 1.377-.608a1.942 1.942 0 0 1 1.393.522c.377.352.6.84.62 1.357c.043 1.738.043 3.477 0 5.215A1.94 1.94 0 0 1 10.805 21a1.922 1.922 0 0 1-1.423.495a1.954 1.954 0 0 1-1.359-.614a1.97 1.97 0 0 1-.535-1.395c-.01-.815 0-1.64 0-2.466m8.938-9.963v2.434a1.996 1.996 0 0 1-.524 1.5a1.98 1.98 0 0 1-2.242.469a1.981 1.981 0 0 1-1.078-1.165a1.996 1.996 0 0 1-.106-.804V4.46a1.963 1.963 0 0 1 .605-1.386a1.947 1.947 0 0 1 1.408-.537a1.962 1.962 0 0 1 1.383.602a1.979 1.979 0 0 1 .553 1.408c.011.836 0 1.673 0 2.51M6.97 11.511H4.545a1.962 1.962 0 0 1-1.393-.579a1.978 1.978 0 0 1-.427-2.155a1.978 1.978 0 0 1 1.066-1.07a1.97 1.97 0 0 1 .754-.15h4.923a1.962 1.962 0 0 1 1.392.579a1.98 1.98 0 0 1-1.392 3.375zm4.478-6.171v.902c0 .18-.06.261-.216.261H9.165A1.916 1.916 0 0 1 7.9 5.787a1.929 1.929 0 0 1-.4-1.402c.022-.492.227-.958.574-1.306a1.965 1.965 0 0 1 3.342 1.12c.032.38.032.487.032.832v.214zm-5.009 7.204c.06.813.06 1.63 0 2.444a1.902 1.902 0 0 1-.754 1.18a1.887 1.887 0 0 1-1.356.34a1.988 1.988 0 0 1-1.293-.627a2.003 2.003 0 0 1-.536-1.338a1.96 1.96 0 0 1 .497-1.346c.33-.369.786-.599 1.278-.643c.736-.065 1.471-.01 2.164-.01M17.443 11.5V9.329c.052-.509.299-.977.689-1.305c.39-.329.891-.492 1.399-.455c.522 0 1.023.208 1.392.579a1.981 1.981 0 0 1 0 2.796c-.37.371-.87.58-1.392.58c-.671 0-1.363-.022-2.088-.022m-4.967 6.072c.8-.055 1.603-.055 2.402 0c.488.09.92.367 1.208.773c.286.406.405.908.329 1.4a1.99 1.99 0 0 1-.67 1.264a1.98 1.98 0 0 1-1.343.485a1.922 1.922 0 0 1-1.314-.528a1.937 1.937 0 0 1-.6-1.287c-.044-.695-.012-1.401-.012-2.107"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Snapchat(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.74 5.457a3.378 3.378 0 0 0 0-.55A3.089 3.089 0 0 0 18.6 2.25H5.42a3.96 3.96 0 0 0-.906.128a3.081 3.081 0 0 0-2.264 3.06v13.075c-.004.268.023.535.079.797a3.1 3.1 0 0 0 3.14 2.44H18.53a3.43 3.43 0 0 0 1.142-.187a3.032 3.032 0 0 0 2.067-2.951c.01-4.418 0-8.786 0-13.155m-3.326 10.914c-.314.18-.536.24-.877.361a.464.464 0 0 0-.315.276a1.074 1.074 0 0 1-.984.708l-.817.05a2.835 2.835 0 0 0-1.24.491a3.84 3.84 0 0 1-3.672.305a3.08 3.08 0 0 1-.64-.344a2.954 2.954 0 0 0-1.693-.492a2.516 2.516 0 0 1-.542-.05a.985.985 0 0 1-.777-.648a.442.442 0 0 0-.345-.305a2.628 2.628 0 0 1-1.142-.581a.876.876 0 0 1 .197-1.505c.355-.207.709-.404 1.044-.64c.41-.297.747-.684.984-1.131c.148-.295.138-.325-.128-.502a7.699 7.699 0 0 1-.807-.59a1.17 1.17 0 0 1-.364-1.28a1.22 1.22 0 0 1 1.437-.855h.187a4.11 4.11 0 0 1 1.26-3.542a4.195 4.195 0 0 1 2.815-.984a4.166 4.166 0 0 1 3.18 1.318a4.298 4.298 0 0 1 .984 3.267c.223-.02.447-.02.67 0a1.142 1.142 0 0 1 .984.866a1.2 1.2 0 0 1-.374 1.25a8.99 8.99 0 0 1-.847.629c-.226.147-.236.177-.118.423a3.876 3.876 0 0 0 1.585 1.574l.433.236a.905.905 0 0 1 .03 1.633z"/><path fill="currentColor" d="M17.832 15.316c0 .068-.098.167-.177.196c-.315.128-.63.256-.984.355a.542.542 0 0 0-.463.521a.296.296 0 0 1-.315.295l-.984.069c-.503.084-.98.279-1.398.57a2.896 2.896 0 0 1-3.258-.058a3.585 3.585 0 0 0-1.969-.561h-.05c-.383 0-.432-.06-.491-.433a.383.383 0 0 0-.256-.335c-.168-.068-.335-.108-.502-.177c-.168-.068-.433-.157-.64-.255c-.207-.099-.187-.119-.187-.197c0-.079.098-.148.177-.187a5.176 5.176 0 0 0 1.723-1.309c.3-.36.52-.779.65-1.23a.423.423 0 0 0-.168-.53c-.266-.178-.541-.355-.797-.542a2.391 2.391 0 0 1-.374-.325a.403.403 0 0 1 0-.462a.354.354 0 0 1 .393-.187c.203.042.404.095.6.157l.13.03l.133.056c.276.099.27-.036.26-.322a13.846 13.846 0 0 1 0-1.918a2.675 2.675 0 0 1 1.88-2.214a3.447 3.447 0 0 1 2.598.069a2.765 2.765 0 0 1 1.782 2.656v1.575c0 .255.05.285.276.206l.669-.226a.743.743 0 0 1 .256 0a.386.386 0 0 1 .374.197a.364.364 0 0 1 0 .442a2.33 2.33 0 0 1-.364.335c-.227.157-.463.295-.7.453c-.235.157-.373.344-.245.728a3.59 3.59 0 0 0 1.27 1.77c.315.246.506.347.862.59c.151.102.3.129.29.198"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SoundWaves(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 9.938v4.124m3.7-7.217v10.31m3.7-13.918v17.526m3.7-13.631v9.736m3.7-7.302v4.868m3.7-3.465v2.062"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spotify(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.053 2A10 10 0 1 0 22 12.053A10.01 10.01 0 0 0 12.053 2m3.892 14.57a.72.72 0 0 1-1.018.18a5.303 5.303 0 0 0-2.715-.741a5.123 5.123 0 0 0-2.704.742a.72.72 0 0 1-1.06-.18a.732.732 0 0 1 .18-1.06a6.542 6.542 0 0 1 3.552-1.061a6.68 6.68 0 0 1 3.563 1.06a.731.731 0 0 1 .149 1.06zm1.283-2.937a.806.806 0 0 1-.636.34a.732.732 0 0 1-.456-.149a7.974 7.974 0 0 0-8.166 0a.788.788 0 0 1-.901-1.294a9.448 9.448 0 0 1 9.968 0a.784.784 0 0 1 .138 1.103zm1.495-2.937a.859.859 0 0 1-.679.329a.902.902 0 0 1-.509-.18a9.756 9.756 0 0 0-10.965 0a.858.858 0 0 1-1.06-1.348a11.463 11.463 0 0 1 13.043 0a.87.87 0 0 1 .117 1.199z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stack(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m20.27 7.796l-3.61 1.876l-4.392 2.236a.548.548 0 0 1-.536 0L7.341 9.672L3.73 7.796a.594.594 0 0 1 0-1.06l8.014-3.925a.571.571 0 0 1 .512 0l8.014 3.925a.594.594 0 0 1 0 1.06"/><path d="m7.34 9.672l-3.61 1.723a.594.594 0 0 0 0 1.06l3.61 1.876l4.392 2.236a.547.547 0 0 0 .536 0l4.391-2.236l3.611-1.875a.594.594 0 0 0 0-1.014l-3.61-1.77"/><path d="m7.34 14.33l-3.61 1.725a.594.594 0 0 0 0 1.06l8.002 4.065a.547.547 0 0 0 .536 0l8.002-4.065a.594.594 0 0 0 0-1.014l-3.61-1.77"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StackFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.125 12.62a1.35 1.35 0 0 0 0-1.31a1.519 1.519 0 0 0-.54-.53l-2.28-1.12l2.3-1.18a1.33 1.33 0 0 0 .53-.49c.13-.214.2-.46.2-.71c0-.248-.07-.49-.2-.7a1.27 1.27 0 0 0-.55-.5l-8-3.92a1.32 1.32 0 0 0-1.18 0l-8 3.93a1.37 1.37 0 0 0-.54.5a1.33 1.33 0 0 0 0 1.4c.124.214.308.388.53.5l2.28 1.17l-2.28 1.09a1.34 1.34 0 0 0-.53.5a1.33 1.33 0 0 0 0 1.4c.124.214.308.388.53.5l2.28 1.17l-2.28 1.09a1.34 1.34 0 0 0-.53.5a1.33 1.33 0 0 0 0 1.4c.127.212.31.385.53.5l8 4c.191.11.408.169.63.17a1.18 1.18 0 0 0 .61-.16l8.06-4.09a1.35 1.35 0 0 0 .47-1.8a1.522 1.522 0 0 0-.54-.53l-2.29-1.12l2.36-1.21a1.38 1.38 0 0 0 .43-.45m-9.14 3.26l-4.26-2.17l-.19-.09l-3.22-1.66l3-1.45l4 2.07c.192.107.41.162.63.16a1.17 1.17 0 0 0 .61-.15l4-2.07l3 1.45z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Star(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m12.495 18.587l4.092 2.15a1.044 1.044 0 0 0 1.514-1.106l-.783-4.552a1.045 1.045 0 0 1 .303-.929l3.31-3.226a1.043 1.043 0 0 0-.575-1.785l-4.572-.657A1.044 1.044 0 0 1 15 7.907l-2.088-4.175a1.044 1.044 0 0 0-1.88 0L8.947 7.907a1.044 1.044 0 0 1-.783.575l-4.51.657a1.044 1.044 0 0 0-.584 1.785l3.309 3.226a1.044 1.044 0 0 1 .303.93l-.783 4.55a1.044 1.044 0 0 0 1.513 1.107l4.093-2.15a1.043 1.043 0 0 1 .991 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarCircle(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="m11.737 6.71l-.569 1.693a4.323 4.323 0 0 1-2.757 2.76l-1.713.569a.287.287 0 0 0 0 .548l1.713.569a4.317 4.317 0 0 1 2.736 2.738l.569 1.715a.288.288 0 0 0 .547 0l.59-1.694a4.323 4.323 0 0 1 2.736-2.738l1.713-.569a.288.288 0 0 0 0-.547l-1.692-.591a4.32 4.32 0 0 1-2.757-2.76l-.569-1.715a.288.288 0 0 0-.547.022"/><rect width="19" height="19" x="2.5" y="2.5" rx="9.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarCircleFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.75a10.25 10.25 0 1 0 0 20.5a10.25 10.25 0 0 0 0-20.5m5.69 10.84a1 1 0 0 1-.5.35l-1.61.54a3.46 3.46 0 0 0-1.31.81a3.36 3.36 0 0 0-.81 1.31l-.55 1.61a1.11 1.11 0 0 1-.35.46a1 1 0 0 1-.57.19a.91.91 0 0 1-.57-.19a1 1 0 0 1-.36-.48l-.53-1.61a3.34 3.34 0 0 0-2.13-2.12l-1.6-.53a1 1 0 0 1-.48-.36a.93.93 0 0 1-.19-.57a1 1 0 0 1 .19-.58a.9.9 0 0 1 .47-.34l1.61-.54a3.56 3.56 0 0 0 1.33-.81c.37-.374.647-.83.81-1.33l.54-1.58a1 1 0 0 1 .33-.48a.89.89 0 0 1 .57-.21a1 1 0 0 1 .56.16a.92.92 0 0 1 .38.47l.54 1.63c.168.503.448.961.82 1.34a3.59 3.59 0 0 0 1.32.81l1.61.56a.9.9 0 0 1 .47.35a1 1 0 0 1 .18.57c-.001.202-.06.4-.17.57"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#mageStarFill0)"><path fill="currentColor" d="M21.95 10.605a1.75 1.75 0 0 1-.5.86l-3.3 3.22a.42.42 0 0 0-.08.12a.34.34 0 0 0 0 .14l.78 4.56c.065.336.03.684-.1 1a1.65 1.65 0 0 1-.61.77a1.83 1.83 0 0 1-.92.35h-.13a1.77 1.77 0 0 1-.84-.21l-4.1-2.14a.28.28 0 0 0-.28 0l-4.1 2.15a1.88 1.88 0 0 1-1 .21a1.83 1.83 0 0 1-.93-.35a1.75 1.75 0 0 1-.61-.78a1.82 1.82 0 0 1-.1-1l.78-4.55a.23.23 0 0 0 0-.14a.37.37 0 0 0-.07-.12l-3.3-3.24a1.78 1.78 0 0 1-.49-.85a1.75 1.75 0 0 1 0-1a1.81 1.81 0 0 1 1.49-1.21l4.5-.66a.18.18 0 0 0 .12-.05a.31.31 0 0 0 .09-.11l2.1-4.18c.143-.3.369-.553.65-.73a1.79 1.79 0 0 1 2.57.74l2.08 4.16a.4.4 0 0 0 .1.12a.21.21 0 0 0 .13.05l4.57.66c.332.05.644.192.9.41c.251.217.441.496.55.81c.106.32.124.662.05.99"/></g><defs><clipPath id="mageStarFill0"><path fill="#fff" d="M2 2.395h20v19.21H2z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarMoving(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m14.524 17.649l3.513 1.84a.867.867 0 0 0 .941-.063a.92.92 0 0 0 .307-.392a.962.962 0 0 0 .053-.486l-.677-3.904a.995.995 0 0 1 0-.434a.92.92 0 0 1 .233-.37l2.835-2.762a.93.93 0 0 0 .233-.92a.899.899 0 0 0-.72-.614l-3.925-.56a.92.92 0 0 1-.677-.498L14.884 4.91a.889.889 0 0 0-.783-.508a.836.836 0 0 0-.476.138a.804.804 0 0 0-.328.37l-1.799 3.576a.931.931 0 0 1-.666.498l-3.872.56a.931.931 0 0 0-.455.201a.868.868 0 0 0-.275.413a.952.952 0 0 0 .253.92L9.32 13.84c.102.106.182.232.233.37a.995.995 0 0 1 0 .434l-.677 3.904a.857.857 0 0 0 0 .486a.92.92 0 0 0 .306.392a.867.867 0 0 0 .942.063l3.513-1.84a.898.898 0 0 1 .846 0zM8 5.4H2m3 12.07H2m1.5-5.773H2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarMovingFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.833 8.142l1.8-3.57a1.64 1.64 0 0 1 1.49-.92c.306 0 .606.09.86.26c.251.166.452.398.58.67l1.76 3.57l.11.08l3.92.57c.302.04.586.165.82.36c.234.205.41.467.51.76a1.66 1.66 0 0 1 0 .91a1.57 1.57 0 0 1-.44.77l-2.84 2.77a.11.11 0 0 0 0 .11l.68 3.93c.047.297.016.6-.09.88a1.7 1.7 0 0 1-1.4 1.05a1.59 1.59 0 0 1-.91-.2l-3.38-1.77l-.17-.07h-.14l-3.52 1.84a1.61 1.61 0 0 1-.76.19h-.17a1.7 1.7 0 0 1-.84-.32a1.54 1.54 0 0 1-.55-.71a1.61 1.61 0 0 1 0-1l.66-3.81a.491.491 0 0 0 0-.11h-.05l-2.82-2.74a1.69 1.69 0 0 1-.46-.8a1.62 1.62 0 0 1 .53-1.65a1.59 1.59 0 0 1 .83-.36l3.87-.57zm-2.83-2h-6a.75.75 0 0 1 0-1.5h6a.75.75 0 1 1 0 1.5m-3 12.07h-3a.75.75 0 1 1 0-1.5h3a.75.75 0 1 1 0 1.5m-1.46-5.77h-1.5a.75.75 0 1 1 0-1.5h1.5a.75.75 0 1 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/><path stroke-linecap="round" stroke-linejoin="round" d="m11.713 6.229l-.62 1.847a4.717 4.717 0 0 1-3.008 3.01l-1.869.622a.314.314 0 0 0 0 .597l1.87.621a4.711 4.711 0 0 1 2.983 2.987l.62 1.87a.315.315 0 0 0 .598 0l.644-1.847a4.717 4.717 0 0 1 2.984-2.986l1.869-.621a.314.314 0 0 0 0-.598l-1.845-.644a4.712 4.712 0 0 1-3.008-3.011l-.62-1.87a.314.314 0 0 0-.598.023"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarSquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#mageStarSquareFill0)"><path fill="currentColor" d="M15.25 2h-6.5A6.76 6.76 0 0 0 2 8.75v6.53a6.76 6.76 0 0 0 6.75 6.75h6.5A6.75 6.75 0 0 0 22 15.28v-6.5A6.748 6.748 0 0 0 15.25 2m2.65 10.59a1 1 0 0 1-.48.34l-1.68.56a3.709 3.709 0 0 0-1.4.87a3.5 3.5 0 0 0-.86 1.4l-.58 1.67a.9.9 0 0 1-.35.47a1 1 0 0 1-.56.18a1 1 0 0 1-.56-.18a.93.93 0 0 1-.35-.48l-.56-1.69a3.59 3.59 0 0 0-2.26-2.26l-1.69-.55a1 1 0 0 1-.47-.36a.94.94 0 0 1 0-1.12a.92.92 0 0 1 .47-.34l1.68-.56a3.58 3.58 0 0 0 1.41-.88c.399-.383.7-.857.88-1.38l.56-1.66a1 1 0 0 1 .32-.48a1 1 0 0 1 .56-.2a1 1 0 0 1 .56.16c.17.109.3.27.37.46l.57 1.71a3.38 3.38 0 0 0 .87 1.41c.395.4.877.701 1.41.88l1.67.58a1 1 0 0 1 .48.35c.105.165.164.355.17.55a1 1 0 0 1-.18.55"/></g><defs><clipPath id="mageStarSquareFill0"><path fill="#fff" d="M2 2h20v20H2z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarsA(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m14.92 3.582l-.624 1.86a4.75 4.75 0 0 1-3.029 3.031l-1.882.625a.316.316 0 0 0 0 .602l1.882.625a4.743 4.743 0 0 1 3.005 3.007l.625 1.884a.317.317 0 0 0 .6 0l.649-1.86a4.75 4.75 0 0 1 3.005-3.007l1.881-.625a.316.316 0 0 0 0-.602l-1.858-.649a4.743 4.743 0 0 1-3.028-3.031l-.625-1.884a.317.317 0 0 0-.6.024m-8.062 8.601l-.446 1.328a3.393 3.393 0 0 1-2.163 2.165l-1.345.447a.225.225 0 0 0-.112.347a.226.226 0 0 0 .112.083l1.345.446a3.388 3.388 0 0 1 2.146 2.148l.446 1.346a.226.226 0 0 0 .43 0l.462-1.329a3.392 3.392 0 0 1 2.146-2.148l1.345-.446a.226.226 0 0 0 0-.43l-1.328-.464a3.389 3.389 0 0 1-2.163-2.165l-.446-1.345a.226.226 0 0 0-.43.017"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarsAfill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.97 9.47a1.07 1.07 0 0 1-.73 1.01l-1.88.62a4 4 0 0 0-2.53 2.54l-.65 1.87a1.09 1.09 0 0 1-.39.52a1.06 1.06 0 0 1-1.63-.54l-.63-1.88a3.931 3.931 0 0 0-1-1.56a4.062 4.062 0 0 0-1.57-1l-1.88-.63a1.05 1.05 0 0 1-.531-.38a1.08 1.08 0 0 1 0-1.25a1.05 1.05 0 0 1 .54-.39l1.87-.63a4.001 4.001 0 0 0 1.59-1c.45-.443.793-.984 1-1.58l.62-1.85a1 1 0 0 1 .36-.53a1 1 0 0 1 .62-.22a1.1 1.1 0 0 1 .63.18a1 1 0 0 1 .41.51l.63 1.91c.207.596.55 1.137 1 1.58a4 4 0 0 0 1.58 1l1.87.66a1 1 0 0 1 .52.38c.13.194.194.425.18.66M12.1 16.4a1 1 0 0 1-.18.57a1 1 0 0 1-.48.35l-1.35.45a2.63 2.63 0 0 0-1 .64a2.711 2.711 0 0 0-.64 1l-.47 1.34a1 1 0 0 1-.35.48a1 1 0 0 1-1.15 0a1 1 0 0 1-.35-.48l-.44-1.34a2.67 2.67 0 0 0-.641-1a2.71 2.71 0 0 0-1-.64l-1.35-.44a1 1 0 0 1-.48-.36a1 1 0 0 1-.19-.57a1 1 0 0 1 .68-.93l1.34-.44a2.78 2.78 0 0 0 1.64-1.64l.45-1.32a1 1 0 0 1 .33-.48a1 1 0 0 1 1.14-.05c.171.118.303.285.38.48l.45 1.37a2.78 2.78 0 0 0 1.64 1.64l1.34.47a.89.89 0 0 1 .47.35a.94.94 0 0 1 .21.55"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarsB(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m14.92 8.797l-.624 1.86a4.75 4.75 0 0 1-3.029 3.03l-1.882.626a.316.316 0 0 0 0 .601l1.882.626a4.744 4.744 0 0 1 3.005 3.007l.625 1.883a.317.317 0 0 0 .6 0l.649-1.86a4.749 4.749 0 0 1 3.005-3.007l1.881-.625a.316.316 0 0 0 0-.601l-1.858-.65a4.744 4.744 0 0 1-3.028-3.03l-.625-1.884a.317.317 0 0 0-.6.024M6.859 3.516l-.446 1.329A3.392 3.392 0 0 1 4.25 7.01l-1.345.446a.226.226 0 0 0 0 .43l1.345.447a3.388 3.388 0 0 1 2.146 2.148l.446 1.345a.226.226 0 0 0 .43 0l.462-1.328A3.392 3.392 0 0 1 9.88 8.35l1.345-.447a.226.226 0 0 0 0-.43L9.897 7.01a3.388 3.388 0 0 1-2.163-2.165l-.446-1.346a.226.226 0 0 0-.43.017"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarsBfill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.97 14.607a1.07 1.07 0 0 1-.73 1l-1.88.62a3.929 3.929 0 0 0-1.56 1a4.06 4.06 0 0 0-1 1.57l-.65 1.87a1.14 1.14 0 0 1-.38.52a1.1 1.1 0 0 1-.63.2a1 1 0 0 1-.62-.2a1.07 1.07 0 0 1-.39-.53l-.63-1.88a4 4 0 0 0-2.53-2.54l-1.88-.62a1.13 1.13 0 0 1-.53-.39a1.06 1.06 0 0 1 .54-1.64l1.87-.62a4 4 0 0 0 2.56-2.55l.62-1.86a1 1 0 0 1 .36-.52a1 1 0 0 1 .61-.23a1 1 0 0 1 .64.18a1 1 0 0 1 .41.52l.63 1.9a4 4 0 0 0 2.55 2.56l1.87.65a1 1 0 0 1 .52.38a1.1 1.1 0 0 1 .23.61M12.1 7.656a1 1 0 0 1-.67.93l-1.34.44a2.63 2.63 0 0 0-1 .64a2.67 2.67 0 0 0-.64 1l-.47 1.34a1 1 0 0 1-.34.47a1.05 1.05 0 0 1-.58.19a1 1 0 0 1-.93-.68l-.44-1.34a2.63 2.63 0 0 0-.64-1a2.71 2.71 0 0 0-1-.64l-1.35-.45a.92.92 0 0 1-.48-.36a.93.93 0 0 1-.19-.57a1 1 0 0 1 .19-.58a1 1 0 0 1 .49-.34l1.34-.45a2.67 2.67 0 0 0 1-.64c.29-.277.509-.62.64-1l.45-1.32a1 1 0 0 1 .33-.48a.93.93 0 0 1 .56-.2a.87.87 0 0 1 .58.16a1 1 0 0 1 .38.47l.45 1.37c.135.378.354.72.64 1a2.67 2.67 0 0 0 1 .64l1.35.47a1 1 0 0 1 .65.92z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarsC(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m15.238 10.81l-.569 1.694a4.325 4.325 0 0 1-2.757 2.76l-1.713.569a.288.288 0 0 0 0 .548l1.713.569a4.318 4.318 0 0 1 2.736 2.738l.568 1.715a.287.287 0 0 0 .548 0l.59-1.694a4.323 4.323 0 0 1 2.735-2.738l1.714-.569a.288.288 0 0 0 0-.548l-1.692-.59a4.318 4.318 0 0 1-2.757-2.76l-.569-1.715a.289.289 0 0 0-.448-.126a.288.288 0 0 0-.099.148m-8.43-4.914l-.413 1.231a3.145 3.145 0 0 1-2.006 2.007l-1.246.414a.21.21 0 0 0 0 .398l1.246.415a3.14 3.14 0 0 1 1.99 1.99l.413 1.248a.21.21 0 0 0 .398 0l.43-1.232a3.145 3.145 0 0 1 1.99-1.99l1.245-.415a.21.21 0 0 0 0-.398l-1.23-.43A3.141 3.141 0 0 1 7.62 7.128l-.414-1.247a.21.21 0 0 0-.398.016m7.849-3.422l-.207.616a1.572 1.572 0 0 1-1.002 1.004l-.623.207a.104.104 0 0 0-.052.16a.104.104 0 0 0 .052.039l.623.207a1.57 1.57 0 0 1 .995.995l.206.624a.105.105 0 0 0 .2 0l.214-.616a1.573 1.573 0 0 1 .995-.995l.623-.207a.105.105 0 0 0 0-.2l-.615-.214a1.571 1.571 0 0 1-1.003-1.004l-.207-.624a.105.105 0 0 0-.199.008"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarsCfill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.738 16.13a1 1 0 0 1-.19.61a1 1 0 0 1-.52.38l-1.71.57a3.57 3.57 0 0 0-1.4.86a3.5 3.5 0 0 0-.86 1.4l-.6 1.7a1 1 0 0 1-.36.51a1.08 1.08 0 0 1-.62.19a1 1 0 0 1-1-.71l-.57-1.71a3.5 3.5 0 0 0-.86-1.4a3.789 3.789 0 0 0-1.4-.87l-1.71-.56a1.11 1.11 0 0 1-.51-.37a1.08 1.08 0 0 1-.21-.62a1 1 0 0 1 .71-1l1.72-.57a3.54 3.54 0 0 0 2.28-2.28l.57-1.69a1 1 0 0 1 .95-.73c.215 0 .426.059.61.17c.182.125.322.303.4.51l.58 1.74a3.54 3.54 0 0 0 2.28 2.28l1.7.6a1 1 0 0 1 .51.38a1 1 0 0 1 .21.61m-9.999-6.36a1 1 0 0 1-.17.55a1 1 0 0 1-.47.35l-1.26.42c-.353.122-.673.32-.94.58a2.48 2.48 0 0 0-.58.94l-.43 1.24a.89.89 0 0 1-.35.47a1 1 0 0 1-.56.18a1 1 0 0 1-.57-.19a1 1 0 0 1-.34-.47l-.41-1.25a2.44 2.44 0 0 0-.58-.93a2.22 2.22 0 0 0-.93-.58l-1.25-.42a.93.93 0 0 1-.48-.35a1 1 0 0 1 .48-1.47l1.25-.41a2.49 2.49 0 0 0 1.53-1.53l.41-1.23a1 1 0 0 1 .32-.47a1 1 0 0 1 .55-.2a1 1 0 0 1 .57.16a1 1 0 0 1 .37.46l.42 1.28a2.49 2.49 0 0 0 1.53 1.53l1.25.43a.92.92 0 0 1 .46.35a.94.94 0 0 1 .18.56m5.789-5.36a1 1 0 0 1-.17.51a.82.82 0 0 1-.42.3l-.62.21a.84.84 0 0 0-.52.52l-.22.63a.929.929 0 0 1-.29.39a.82.82 0 0 1-.52.18a1.08 1.08 0 0 1-.49-.15a.92.92 0 0 1-.32-.44l-.21-.62a.719.719 0 0 0-.2-.32a.76.76 0 0 0-.32-.2l-.62-.2a1 1 0 0 1-.42-.31a.88.88 0 0 1-.16-.51a.94.94 0 0 1 .17-.51a.88.88 0 0 1 .42-.3l.61-.2a.91.91 0 0 0 .33-.2a.939.939 0 0 0 .2-.33l.21-.62c.06-.155.155-.292.28-.4a1 1 0 0 1 .49-.19a.94.94 0 0 1 .53.16a1 1 0 0 1 .32.41l.21.64a.942.942 0 0 0 .2.33a1 1 0 0 0 .32.2l.63.21a1 1 0 0 1 .41.3a.87.87 0 0 1 .17.51"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Steam(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.057 13.698a2.078 2.078 0 0 0-.74.14l.983.382c.214.089.423.19.625.305a1.54 1.54 0 0 1 .639 1.272a1.56 1.56 0 0 1-.465 1.098c-.294.29-.692.454-1.106.454a2.087 2.087 0 0 1-.817-.23l-1.111-.444a2.165 2.165 0 0 0 1.966 1.272c.566 0 1.108-.224 1.508-.622a2.12 2.12 0 0 0 0-3.005a2.137 2.137 0 0 0-1.508-.622zM11.994 2.25a9.779 9.779 0 0 0-6.646 2.59a9.703 9.703 0 0 0-3.098 6.403l5.274 2.239c.46-.284.993-.425 1.533-.407h.14l2.325-3.282a1.166 1.166 0 0 1 0-.216c0-.755.224-1.493.645-2.12a3.828 3.828 0 0 1 1.72-1.406a3.846 3.846 0 0 1 4.175.827a3.811 3.811 0 0 1 .83 4.16a3.82 3.82 0 0 1-1.41 1.712c-.63.42-1.371.643-2.129.643H15.2l-3.359 2.404v.102a2.768 2.768 0 0 1-.734 1.834a2.788 2.788 0 0 1-3.687.367a2.772 2.772 0 0 1-1.083-1.654L2.62 14.88a9.74 9.74 0 0 0 2.896 4.45a9.818 9.818 0 0 0 10.099 1.717a9.778 9.778 0 0 0 4.213-3.241a9.71 9.71 0 0 0 .886-10.166a9.758 9.758 0 0 0-3.59-3.915a9.812 9.812 0 0 0-5.117-1.463z"/><path fill="currentColor" d="M15.302 7.033a2.561 2.561 0 0 0-2.36 1.57a2.535 2.535 0 0 0 .554 2.773a2.559 2.559 0 0 0 3.93-.386a2.537 2.537 0 0 0-.318-3.212a2.56 2.56 0 0 0-1.806-.745m0 4.44a1.934 1.934 0 0 1-1.79-1.178a1.913 1.913 0 0 1 .409-2.097a1.931 1.931 0 0 1 3.297 1.353a1.92 1.92 0 0 1-1.916 1.921"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stop(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17.25 5H6.75A1.75 1.75 0 0 0 5 6.75v10.5c0 .966.784 1.75 1.75 1.75h10.5A1.75 1.75 0 0 0 19 17.25V6.75A1.75 1.75 0 0 0 17.25 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopCircle(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 21.5a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19"/><path d="M15.089 7.882H8.91a1.03 1.03 0 0 0-1.03 1.03v6.177c0 .568.462 1.03 1.03 1.03h6.178a1.03 1.03 0 0 0 1.03-1.03V8.91a1.03 1.03 0 0 0-1.03-1.03"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopCircleFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.75A10.25 10.25 0 1 0 22.25 12A10.26 10.26 0 0 0 12 1.75m4.28 13a1.56 1.56 0 0 1-1.57 1.56H9.29a1.56 1.56 0 0 1-1.57-1.56V9.28a1.56 1.56 0 0 1 1.57-1.56h5.42a1.56 1.56 0 0 1 1.57 1.56z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.75 6.75v10.5a2.51 2.51 0 0 1-2.5 2.5H6.75a2.51 2.51 0 0 1-2.5-2.5V6.75a2.5 2.5 0 0 1 2.5-2.5h10.5a2.5 2.5 0 0 1 2.5 2.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M15.008 7.99H8.993c-.554 0-1.003.449-1.003 1.002v6.015c0 .554.45 1.003 1.003 1.003h6.015c.553 0 1.002-.449 1.002-1.003V8.993c0-.553-.449-1.002-1.002-1.002"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StopSquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 2h-6.5A6.76 6.76 0 0 0 2 8.75v6.5A6.76 6.76 0 0 0 8.75 22h6.5A6.76 6.76 0 0 0 22 15.25v-6.5A6.76 6.76 0 0 0 15.25 2m1 12.52a1.69 1.69 0 0 1-1.75 1.68h-5a1.69 1.69 0 0 1-1.7-1.7v-5a1.69 1.69 0 0 1 1.7-1.7h5a1.69 1.69 0 0 1 1.7 1.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stripe(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.374 9.668h1.217l.07.331l.251-.18a1.619 1.619 0 0 1 1.88.14c.435.38.717.906.794 1.478a3.66 3.66 0 0 1-.19 2.01a1.74 1.74 0 0 1-1.348 1.137a1.508 1.508 0 0 1-1.116-.221l-.15-.09v1.557l-1.388.292zm1.387 2.433v.945a.201.201 0 0 0 .07.16a.814.814 0 0 0 1.217-.17a1.77 1.77 0 0 0 .242-1.006a1.577 1.577 0 0 0-.262-.895a.783.783 0 0 0-1.206-.16a.222.222 0 0 0-.06.14c-.01.352 0 .674 0 .986m4.524.512a.795.795 0 0 0 .382.674c.289.155.622.208.945.15c.366-.02.725-.112 1.055-.27v.974a.17.17 0 0 1-.11.181a3.286 3.286 0 0 1-1.75.282a2.121 2.121 0 0 1-1.287-.613a2.202 2.202 0 0 1-.552-1.086a3.327 3.327 0 0 1 .24-2.202a1.92 1.92 0 0 1 3.6.231c.118.323.183.662.191 1.006v.683h-2.714zm1.407-1.005a.845.845 0 0 0-.342-.774a.623.623 0 0 0-.714 0a.876.876 0 0 0-.372.784zM2.134 12.985l.252.12c.34.166.708.268 1.086.303c.153.014.308.014.462 0a.301.301 0 0 0 .11-.533c-.1-.07-.209-.131-.321-.181c-.322-.131-.644-.242-1.005-.392a1.377 1.377 0 0 1-.252-2.242a2.01 2.01 0 0 1 1.257-.453c.489-.03.978.041 1.438.211a.12.12 0 0 1 .09.131v1.176c-.221-.07-.432-.17-.653-.22c-.222-.051-.533-.071-.795-.101a.542.542 0 0 0-.22 0a.251.251 0 0 0-.071.452c.107.072.222.133.342.181c.341.113.677.244 1.005.392a1.267 1.267 0 0 1 .684 1.428a1.428 1.428 0 0 1-1.197 1.236c-.726.171-1.489.1-2.171-.2a.1.1 0 0 1-.06-.081v-1.207a.08.08 0 0 1 .02-.02m3.719-4.222l1.347-.292v1.197h1.006v1.176H7.2v2.01a.523.523 0 0 0 .654.543c.132-.013.263-.037.392-.07v1.015c0 .07 0 .111-.08.131a2.08 2.08 0 0 1-1.458 0a1.317 1.317 0 0 1-.865-1.156a1.337 1.337 0 0 1 0-.201V8.763zm2.846.915h1.095c.08 0 .1 0 .11.09c.018.102.041.203.071.302l.04-.1a.925.925 0 0 1 1.006-.312v1.246a3.498 3.498 0 0 0-.463 0a.915.915 0 0 0-.332.1a.442.442 0 0 0-.261.433v3.066H8.699zm2.774 0h1.377v4.835h-1.377zm-.01-.423V8.25c0-.06 0-.09.07-.1l1.287-.272v1.086z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sun(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" clip-path="url(#mageSun0)"><path d="M12 17.885a5.885 5.885 0 1 0 0-11.77a5.885 5.885 0 0 0 0 11.77m-9.281-5.879H1.5m21 0h-1.207m-9.287-9.287V1.5m0 21v-1.207M5.435 5.435l-.859-.859m14.848 14.848l-.86-.86m.001-13.129l.858-.859M4.576 19.424l.86-.86"/></g><defs><clipPath id="mageSun0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SunFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.891 12a5.944 5.944 0 0 1-3.68 5.499a5.968 5.968 0 0 1-6.496-1.295A5.948 5.948 0 0 1 11.943 6.05a5.956 5.956 0 0 1 4.21 1.743A5.94 5.94 0 0 1 17.89 12M3.203 13.048H2.05A1.05 1.05 0 0 1 1 12a1.047 1.047 0 0 1 1.05-1.048h1.153A1.05 1.05 0 0 1 4.253 12a1.047 1.047 0 0 1-1.05 1.048m18.747 0h-1.143A1.05 1.05 0 0 1 19.758 12a1.047 1.047 0 0 1 1.05-1.048h1.143A1.05 1.05 0 0 1 23 12a1.047 1.047 0 0 1-1.05 1.048m-9.965-8.8a1.05 1.05 0 0 1-1.05-1.048V2.048A1.047 1.047 0 0 1 11.986 1a1.05 1.05 0 0 1 1.049 1.048V3.2a1.047 1.047 0 0 1-1.05 1.048m0 18.752a1.05 1.05 0 0 1-1.05-1.047V20.8a1.047 1.047 0 0 1 1.05-1.048a1.05 1.05 0 0 1 1.049 1.048v1.152A1.047 1.047 0 0 1 11.984 23M5.753 6.825a1.05 1.05 0 0 1-.745-.314l-.819-.807a1.051 1.051 0 0 1 .745-1.796c.28 0 .548.111.745.308l.819.817a1.047 1.047 0 0 1 0 1.478a1.05 1.05 0 0 1-.745.314m13.271 13.221a1.05 1.05 0 0 1-.735-.304l-.818-.817a1.047 1.047 0 0 1 1.14-1.739c.13.063.245.152.34.262l.818.817a1.047 1.047 0 0 1 0 1.477a1.05 1.05 0 0 1-.745.304m-.808-13.221a1.05 1.05 0 0 1-1.034-1.254c.04-.204.142-.391.29-.538l.818-.817a1.05 1.05 0 0 1 1.48 1.488l-.82.807a1.05 1.05 0 0 1-.734.314M4.934 20.046a1.051 1.051 0 0 1-.745-.304a1.046 1.046 0 0 1 0-1.477l.819-.817a1.048 1.048 0 0 1 1.49 0a1.047 1.047 0 0 1 0 1.477l-.819.817a1.049 1.049 0 0 1-.745.304"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SwimRingFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.785 8.9a10.18 10.18 0 0 0-9.76-7.15c-1.058 0-2.11.158-3.12.47a10.28 10.28 0 0 0-6.69 6.69a10.11 10.11 0 0 0-.46 3.09a10.22 10.22 0 0 0 10.22 10.25a9.884 9.884 0 0 0 3.12-.48a10.28 10.28 0 0 0 6.69-6.69a10.09 10.09 0 0 0 .46-3.08v-.09a10.22 10.22 0 0 0-.46-3.01M12.005 3a10.39 10.39 0 0 1 2.42.28l-.65 3.23a6.73 6.73 0 0 0-3.59 0l-.65-3.23a10.13 10.13 0 0 1 2.47-.28m-5.56 10.68l-3.08.61a10.05 10.05 0 0 1 0-4.67l3.09.61a6.32 6.32 0 0 0 0 3.45zm5.55 7.3a9.866 9.866 0 0 1-2.46-.3l.65-3.25c.59.16 1.199.24 1.81.24a6.21 6.21 0 0 0 1.82-.26l.65 3.26a11.44 11.44 0 0 1-2.47.35zm4.09-7.3a4.273 4.273 0 0 1-2.36 2.36a4.399 4.399 0 0 1-3.45 0a4.317 4.317 0 0 1-2.37-2.36a4.46 4.46 0 0 1-.35-1.73a4.5 4.5 0 0 1 2.72-4.09a4.44 4.44 0 0 1 1.72-.35a4.5 4.5 0 0 1 4.1 2.72c.23.547.35 1.136.35 1.73a4.28 4.28 0 0 1-.36 1.72m1.48 0c.151-.561.229-1.14.23-1.72a5.94 5.94 0 0 0-.25-1.73l3.11-.61c.36 1.536.36 3.134 0 4.67z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tablet(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m10.196 5.342l-4.845 4.846a5.992 5.992 0 1 0 8.473 8.473l4.845-4.845a5.992 5.992 0 0 0-8.473-8.474"/><path d="M16.846 15.639a12.922 12.922 0 0 1-5.053-3.421a12.923 12.923 0 0 1-3.42-5.053m3.633 0a2.559 2.559 0 0 1 3.633 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TabletFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.205 4.793a6.74 6.74 0 0 0-9.54 0l-1.75 1.76a.34.34 0 0 0-.12.12l-3 3a6.742 6.742 0 0 0 9.53 9.54l4.85-4.85a6.74 6.74 0 0 0 .03-9.57m-7.91 1.65a3.56 3.56 0 0 1 3.9-.78c.432.18.823.445 1.15.78a1 1 0 0 1 0 1.41a1 1 0 0 1-1.41 0a1.57 1.57 0 0 0-1.71-.34a1.43 1.43 0 0 0-.51.34a1.001 1.001 0 0 1-1.42-1.41m4.24 9.43a13.68 13.68 0 0 1-4.3-3.17a13.49 13.49 0 0 1-3.12-4.24l1.14-1.14a12.16 12.16 0 0 0 3 4.32a12.29 12.29 0 0 0 4.38 3.09z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tag(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m17.524 17.524l-2.722 2.723a2.567 2.567 0 0 1-3.634 0L4.13 13.209A3.852 3.852 0 0 1 3 10.487V5.568A2.568 2.568 0 0 1 5.568 3h4.919c1.021 0 2 .407 2.722 1.13l7.038 7.038a2.567 2.567 0 0 1 0 3.634z"/><path d="M9.126 11.694a2.568 2.568 0 1 0 0-5.137a2.568 2.568 0 0 0 0 5.137"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagCheck(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.25 12v3.38a3.38 3.38 0 0 1-3.38 3.382H8.661a3.38 3.38 0 0 1-2.389-.992L3.22 13.6a2.964 2.964 0 0 1 0-3.2l3.054-4.17a3.38 3.38 0 0 1 2.39-.992h9.206a3.38 3.38 0 0 1 3.38 3.382z"/><path d="m9.61 11.905l1.946 1.946a.735.735 0 0 0 1.047 0l3.922-3.921"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagCheckFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.79 5.72a4.11 4.11 0 0 0-2.92-1.21H8.66a4.12 4.12 0 0 0-2.92 1.22a.27.27 0 0 0-.07.08l-3.08 4.21a3.66 3.66 0 0 0-.59 2a3.74 3.74 0 0 0 .61 2l3.06 4.17l.07.09a4.11 4.11 0 0 0 2.92 1.21h9.21A4.13 4.13 0 0 0 22 15.36V8.59a4.11 4.11 0 0 0-1.21-2.87m-3.56 4.94l-3.92 3.92c-.16.161-.35.29-.56.38a1.67 1.67 0 0 1-1.34 0a1.78 1.78 0 0 1-.57-.38L8.9 12.64a1.004 1.004 0 0 1 1.42-1.42l1.76 1.77l3.74-3.74a1 1 0 1 1 1.41 1.41"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagCross(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="m16.216 9.389l-5.21 5.21M11.005 9.4l5.211 5.211"/><path stroke-linejoin="round" d="M21.25 12v3.38a3.38 3.38 0 0 1-3.38 3.382H8.661a3.38 3.38 0 0 1-2.389-.992L3.22 13.6a2.964 2.964 0 0 1 0-3.2l3.054-4.17a3.38 3.38 0 0 1 2.39-.992h9.206a3.38 3.38 0 0 1 3.38 3.382z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagCrossFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.818 5.695a4.17 4.17 0 0 0-2.93-1.21h-9.18a4.1 4.1 0 0 0-2.92 1.22l-.08.08l-3.08 4.21a3.74 3.74 0 0 0 0 4.05l3 4.17l.08.09a4.091 4.091 0 0 0 2.92 1.21h9.2a4.17 4.17 0 0 0 2.93-1.21a4.109 4.109 0 0 0 1.21-2.92v-6.77a4.11 4.11 0 0 0-1.15-2.92m-3.87 8.21a1 1 0 1 1-1.42 1.41l-1.9-1.9l-1.89 1.89a1 1 0 0 1-1.42-1.41l1.9-1.89l-1.9-1.9a1 1 0 1 1 1.42-1.41l1.89 1.89l1.9-1.9a1 1 0 1 1 1.42 1.41l-1.9 1.91z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.52 11.615a3.3 3.3 0 0 0-.76-1l-7-7a4.56 4.56 0 0 0-3.25-1.35H5.59a3.31 3.31 0 0 0-3.32 3.31v4.92a4.58 4.58 0 0 0 1.35 3.26l7 7a3.3 3.3 0 0 0 1.08.72c.401.171.833.26 1.27.26a3.33 3.33 0 0 0 2.34-1l2.73-2.72l2.72-2.72a3.3 3.3 0 0 0 .72-1.08a3.35 3.35 0 0 0 0-2.54zm-12.37.28a2.87 2.87 0 1 1 2.87-2.87a2.881 2.881 0 0 1-2.87 2.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagMinus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="M9.62 12h7.37"/><path stroke-linejoin="round" d="M21.25 12v3.38a3.38 3.38 0 0 1-3.38 3.382H8.661a3.38 3.38 0 0 1-2.389-.992L3.22 13.6a2.964 2.964 0 0 1 0-3.2l3.054-4.17a3.38 3.38 0 0 1 2.39-.992h9.206a3.38 3.38 0 0 1 3.38 3.382z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagMinusFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.837 5.695a4.11 4.11 0 0 0-2.92-1.21h-9.21a4.12 4.12 0 0 0-2.92 1.22l-.07.08l-3.08 4.21a3.69 3.69 0 0 0 0 4.05l3 4.17l.07.09a4.11 4.11 0 0 0 2.92 1.21h9.21a4.13 4.13 0 0 0 4.13-4.13v-6.77a4.11 4.11 0 0 0-1.13-2.92m-3.8 7.31h-7.37a1 1 0 0 1 0-2h7.37a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagPlus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="M13.211 8.315v7.37m-3.676-3.677h7.369"/><path stroke-linejoin="round" d="M21.25 12v3.38a3.38 3.38 0 0 1-3.38 3.382H8.661a3.38 3.38 0 0 1-2.389-.992L3.22 13.6a2.964 2.964 0 0 1 0-3.2l3.054-4.17a3.38 3.38 0 0 1 2.39-.992h9.206a3.38 3.38 0 0 1 3.38 3.382z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagPlusFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.82 5.725a4.15 4.15 0 0 0-2.92-1.21H8.74a4.12 4.12 0 0 0-3 1.21a.27.27 0 0 0-.07.08l-3.08 4.21a3.66 3.66 0 0 0-.59 2a3.74 3.74 0 0 0 .61 2l3.06 4.17l.07.09a4.109 4.109 0 0 0 2.92 1.21h9.21a4.15 4.15 0 0 0 2.92-1.21a4.11 4.11 0 0 0 1.21-2.92v-6.77a4.11 4.11 0 0 0-1.18-2.86m-3.89 7.31h-2.69v2.69a1 1 0 1 1-2 0v-2.68H9.56a1 1 0 0 1 0-2h2.68v-2.69a1 1 0 1 1 2 0v2.69h2.69a1 1 0 1 1 0 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagQuestionMark(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M21.25 12v3.38a3.38 3.38 0 0 1-3.38 3.382H8.661a3.38 3.38 0 0 1-2.389-.992L3.22 13.6a2.964 2.964 0 0 1 0-3.2l3.054-4.17a3.38 3.38 0 0 1 2.39-.992h9.206a3.38 3.38 0 0 1 3.38 3.382z"/><path stroke-miterlimit="10" d="M10.911 9.543a1.9 1.9 0 0 1 2.07-1.076a1.842 1.842 0 0 1 1.351.962a1.598 1.598 0 0 1-.931 2.2a1.18 1.18 0 0 0-.761 1.066v.44"/><path stroke-linejoin="round" d="M12.606 15.565h.004"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagQuestionMarkFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.837 5.695a4.111 4.111 0 0 0-2.93-1.21h-9.2a4.1 4.1 0 0 0-2.92 1.22l-.08.08l-3.08 4.21a3.74 3.74 0 0 0 0 4.05l3 4.17l.08.09a4.089 4.089 0 0 0 2.92 1.21h9.2a4.13 4.13 0 0 0 4.14-4.13v-6.77a4.11 4.11 0 0 0-1.13-2.92m-8.18 11a1 1 0 1 1 0-2a1 1 0 0 1 0 2m2.67-5.52a2.61 2.61 0 0 1-.63.9a2.73 2.73 0 0 1-.93.55v.05a.19.19 0 0 0 0 .1v.41a1 1 0 1 1-2 0v-.49a2.161 2.161 0 0 1 1.4-2a.66.66 0 0 0 .25-.14a.6.6 0 0 0 .14-.21a.49.49 0 0 0 0-.24a.54.54 0 0 0-.06-.25a.76.76 0 0 0-.23-.25a.82.82 0 0 0-.36-.15a1 1 0 0 0-.61.08a.9.9 0 0 0-.41.43a1 1 0 0 1-1.841-.073a1 1 0 0 1 .031-.767a2.91 2.91 0 0 1 1.3-1.36a2.86 2.86 0 0 1 1.86-.28a2.802 2.802 0 0 1 2.05 1.49a2.611 2.611 0 0 1 .07 2.16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagTwo(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m17.524 17.524l-2.722 2.723a2.567 2.567 0 0 1-3.634 0L4.13 13.209A3.852 3.852 0 0 1 3 10.487V5.568A2.568 2.568 0 0 1 5.568 3h4.919c1.021 0 2 .407 2.722 1.13l7.038 7.038a2.567 2.567 0 0 1 0 3.634z"/><path d="M9.126 11.694a2.568 2.568 0 1 0 0-5.137a2.568 2.568 0 0 0 0 5.137m3.326 4.392l3.634-3.634"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagTwoFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.48 11.695a3.412 3.412 0 0 0-.72-1.08l-7-7a4.56 4.56 0 0 0-3.25-1.35H5.59a3.31 3.31 0 0 0-3.32 3.31v4.92a4.58 4.58 0 0 0 1.35 3.26l7 7a3.33 3.33 0 0 0 2.35.98a3.28 3.28 0 0 0 1.27-.26c.4-.169.763-.413 1.07-.72l2.73-2.72l2.72-2.72c.307-.31.551-.677.72-1.08a3.35 3.35 0 0 0 0-2.54m-12.37.46a3.05 3.05 0 1 1 3-3.05a3.06 3.06 0 0 1-3 3.08zm7.66 1l-3.63 3.64a1 1 0 1 1-1.41-1.42l3.63-3.63a1 1 0 1 1 1.41 1.42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Telegram(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.077 2.001a9.99 9.99 0 0 1 7.077 3.014A10.01 10.01 0 0 1 22 12.167c-.678 13.192-19.556 13.08-20-.155a9.953 9.953 0 0 1 2.936-7.127a9.932 9.932 0 0 1 7.141-2.884m1.818 8.376s.016.01-.107.166a5.43 5.43 0 0 1-.489.512l-2.544 2.47c-.544.544-.533.878.1 1.334c.811.578 1.633 1.112 2.467 1.68c.833.567 1.855 1.4 2.278.177a6.6 6.6 0 0 0 .244-1c.178-.98.356-1.958.511-2.948c.211-1.39.411-2.78.589-4.182c.089-.69-.278-1.001-.967-.834a5.552 5.552 0 0 0-.833.266l-7.256 3.06c-.688.288-1.377.6-2.055.934c-.167.088-.378.333-.367.5c.011.167.245.356.434.434c.466.189.966.322 1.455.478a2.375 2.375 0 0 0 2.222-.367a78.078 78.078 0 0 1 2.811-1.913c.445-.3.756-.467 1.19-.768c.222-.11.272-.162.317-.095z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Television(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18.25 3.33H5.75a3 3 0 0 0-3 3v7.87a3 3 0 0 0 3 3h12.5a3 3 0 0 0 3-3V6.33a3 3 0 0 0-3-3M12 17.2v3.47m-3.47 0h6.94"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TelevisionCheck(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18.25 3.33H5.75a3 3 0 0 0-3 3v7.87a3 3 0 0 0 3 3h12.5a3 3 0 0 0 3-3V6.33a3 3 0 0 0-3-3M12 17.2v3.47m-3.47 0h6.94"/><path d="m9 10.183l1.689 1.689a.639.639 0 0 0 .908 0L15 8.469"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TelevisionCheckFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.9 3.665a3.7 3.7 0 0 0-2.65-1.1H5.75A3.76 3.76 0 0 0 2 6.315v7.87a3.76 3.76 0 0 0 3.75 3.75h5.5v2H8.53a.75.75 0 0 0 0 1.5h6.94a.75.75 0 1 0 0-1.5h-2.72v-2h5.5a3.719 3.719 0 0 0 2.65-1.1a3.76 3.76 0 0 0 1.1-2.65v-7.87a3.721 3.721 0 0 0-1.1-2.65m-5.37 5.32l-3.4 3.4a1.409 1.409 0 0 1-.99.41a1.449 1.449 0 0 1-.54-.11a1.347 1.347 0 0 1-.45-.31l-1.68-1.68a.75.75 0 1 1 1.06-1.06l1.62 1.62l3.32-3.33a.75.75 0 1 1 1.06 1.06"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TelevisionCross(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M18.25 3.33H5.75a3 3 0 0 0-3 3v7.87a3 3 0 0 0 3 3h12.5a3 3 0 0 0 3-3V6.33a3 3 0 0 0-3-3M12 17.2v3.47m-3.47 0h6.94"/><path stroke-miterlimit="10" d="m14.25 8.015l-4.5 4.49m0-4.48l4.5 4.49"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TelevisionCrossFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.9 3.665a3.7 3.7 0 0 0-2.65-1.1H5.75A3.76 3.76 0 0 0 2 6.315v7.87a3.76 3.76 0 0 0 3.75 3.75h5.5v2H8.53a.75.75 0 0 0 0 1.5h6.94a.75.75 0 1 0 0-1.5h-2.72v-2h5.5a3.719 3.719 0 0 0 2.65-1.1a3.76 3.76 0 0 0 1.1-2.65v-7.87a3.721 3.721 0 0 0-1.1-2.65m-6.12 8.3a.75.75 0 0 1-.53 1.28a.77.77 0 0 1-.53-.21l-1.75-1.69l-1.71 1.71a.77.77 0 0 1-.817.154a.75.75 0 0 1-.243-1.224l1.71-1.7l-1.71-1.71a.77.77 0 0 1 0-1.07a.75.75 0 0 1 1.06 0l1.71 1.71l1.73-1.72a.75.75 0 0 1 1.06 0a.77.77 0 0 1 0 1.07l-1.72 1.72z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TelevisionDownload(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M18.25 3.33H5.75a3 3 0 0 0-3 3v7.87a3 3 0 0 0 3 3h12.5a3 3 0 0 0 3-3V6.33a3 3 0 0 0-3-3M12 17.2v3.47m-3.47 0h6.94"/><path stroke-miterlimit="10" d="M12 13.265v-6"/><path stroke-linejoin="round" d="m9.249 10.739l2.36 2.36a.552.552 0 0 0 .782 0l2.36-2.36"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TelevisionDownloadFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.9 3.665a3.7 3.7 0 0 0-2.65-1.1H5.75A3.76 3.76 0 0 0 2 6.315v7.87a3.76 3.76 0 0 0 3.75 3.75h5.5v2H8.53a.75.75 0 0 0 0 1.5h6.94a.75.75 0 1 0 0-1.5h-2.72v-2h5.5a3.719 3.719 0 0 0 2.65-1.1a3.76 3.76 0 0 0 1.1-2.65v-7.87a3.721 3.721 0 0 0-1.1-2.65m-5.62 6.64a.73.73 0 0 1-.53.22a.74.74 0 0 1-.53-.22l-1.47-1.47v4.41a.75.75 0 1 1-1.5 0v-4.38l-1.47 1.47a.75.75 0 0 1-1.06-1.06l2.36-2.36a1.34 1.34 0 0 1 .42-.28c.11-.05.229-.08.35-.09a.76.76 0 0 1 .3 0c.12.01.24.04.35.09c.156.066.299.16.42.28l2.36 2.36a.741.741 0 0 1 0 1.06z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TelevisionFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 6.315v7.87a3.76 3.76 0 0 1-2.315 3.466a3.72 3.72 0 0 1-1.435.284h-5.5v2h2.72a.75.75 0 1 1 0 1.5H8.53a.75.75 0 1 1 0-1.5h2.72v-2h-5.5a3.72 3.72 0 0 1-2.65-1.1a3.76 3.76 0 0 1-1.1-2.65v-7.87a3.76 3.76 0 0 1 3.75-3.75h12.5A3.76 3.76 0 0 1 22 6.315"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TelevisionMinus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M18.25 3.33H5.75a3 3 0 0 0-3 3v7.87a3 3 0 0 0 3 3h12.5a3 3 0 0 0 3-3V6.33a3 3 0 0 0-3-3M12 17.2v3.47m-3.47 0h6.94"/><path stroke-miterlimit="10" d="M9 10.272h6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TelevisionMinusFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.9 3.665a3.72 3.72 0 0 0-2.65-1.1H5.75A3.76 3.76 0 0 0 2 6.315v7.87a3.76 3.76 0 0 0 2.315 3.466a3.72 3.72 0 0 0 1.435.284h5.5v2H8.53a.75.75 0 1 0 0 1.5h6.94a.75.75 0 1 0 0-1.5h-2.72v-2h5.5a3.72 3.72 0 0 0 2.65-1.1a3.76 3.76 0 0 0 1.1-2.65v-7.87a3.721 3.721 0 0 0-1.1-2.65m-5.9 7.34H9a.75.75 0 0 1 0-1.5h6a.75.75 0 1 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TelevisionPlus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M18.25 3.33H5.75a3 3 0 0 0-3 3v7.87a3 3 0 0 0 3 3h12.5a3 3 0 0 0 3-3V6.33a3 3 0 0 0-3-3M12 17.2v3.47m-3.47 0h6.94"/><path stroke-miterlimit="10" d="M11.993 7.265v6M9 10.272h6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TelevisionPlusFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.9 3.665a3.72 3.72 0 0 0-2.65-1.1H5.75A3.76 3.76 0 0 0 2 6.315v7.87a3.76 3.76 0 0 0 2.315 3.466a3.72 3.72 0 0 0 1.435.284h5.5v2H8.53a.75.75 0 1 0 0 1.5h6.94a.75.75 0 1 0 0-1.5h-2.72v-2h5.5a3.72 3.72 0 0 0 2.65-1.1a3.76 3.76 0 0 0 1.1-2.65v-7.87a3.721 3.721 0 0 0-1.1-2.65m-5.9 7.34h-2.26v2.24a.75.75 0 1 1-1.5 0v-2.24H8.97a.75.75 0 0 1 0-1.5h2.24v-2.26a.75.75 0 1 1 1.5 0v2.26h2.26a.75.75 0 1 1 0 1.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TelevisionQuestionMark(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M18.25 3.33H5.75a3 3 0 0 0-3 3v7.87a3 3 0 0 0 3 3h12.5a3 3 0 0 0 3-3V6.33a3 3 0 0 0-3-3M12 17.2v3.47m-3.47 0h6.94"/><path stroke-miterlimit="10" d="M10.274 7.893a1.834 1.834 0 0 1 1.999-1.04a1.78 1.78 0 0 1 1.304.93a1.545 1.545 0 0 1-.9 2.124a1.14 1.14 0 0 0-.734 1.03v.424"/><path stroke-linejoin="round" d="M11.91 13.707h.005"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TelevisionQuestionMarkFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.9 3.665a3.7 3.7 0 0 0-2.65-1.1H5.75A3.76 3.76 0 0 0 2 6.315v7.87a3.76 3.76 0 0 0 3.75 3.75h5.5v2H8.53a.75.75 0 1 0 0 1.5h6.94a.75.75 0 1 0 0-1.5h-2.72v-2h5.5a3.72 3.72 0 0 0 2.65-1.1a3.76 3.76 0 0 0 1.1-2.65v-7.87a3.72 3.72 0 0 0-1.1-2.65m-9.74 10a.75.75 0 1 1 0 .06zm3.08-4.208c-.124.3-.262.407-.5.628a2.361 2.361 0 0 1-.82.49a.26.26 0 0 0-.15.13a.39.39 0 0 0-.08.21v.4a.75.75 0 1 1-1.5 0v-.42a2.06 2.06 0 0 1 .35-1.07a2 2 0 0 1 .87-.67a.83.83 0 0 0 .31-.17a.67.67 0 0 0 .19-.28a.72.72 0 0 0-.01-.64a1.05 1.05 0 0 0-.3-.34a1.14 1.14 0 0 0-.44-.18a1.1 1.1 0 0 0-.72.1a1.18 1.18 0 0 0-.48.51a.75.75 0 1 1-1.37-.62a2.59 2.59 0 0 1 2.82-1.47a2.52 2.52 0 0 1 1.06.44a2.6 2.6 0 0 1 .77.88a2.37 2.37 0 0 1 .05 1.97z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TelevisionUpload(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M18.25 3.33H5.75a3 3 0 0 0-3 3v7.87a3 3 0 0 0 3 3h12.5a3 3 0 0 0 3-3V6.33a3 3 0 0 0-3-3M12 17.2v3.47m-3.47 0h6.94"/><path stroke-miterlimit="10" d="M12 7.265v6"/><path stroke-linejoin="round" d="m14.752 9.791l-2.361-2.36a.55.55 0 0 0-.782 0L9.25 9.79"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TelevisionUploadFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.9 3.665a3.7 3.7 0 0 0-2.65-1.1H5.75A3.76 3.76 0 0 0 2 6.315v7.87a3.76 3.76 0 0 0 2.315 3.466a3.72 3.72 0 0 0 1.435.284h5.5v2H8.53a.75.75 0 1 0 0 1.5h6.94a.75.75 0 1 0 0-1.5h-2.72v-2h5.5a3.72 3.72 0 0 0 2.65-1.1a3.76 3.76 0 0 0 1.1-2.65v-7.87a3.72 3.72 0 0 0-1.1-2.65m-5.62 7.59l-2.31 2.41a1.19 1.19 0 0 1-.42.28a1.3 1.3 0 0 1-1 0a1.21 1.21 0 0 1-.43-.29l-2.35-2.35a.75.75 0 1 1 1.06-1.06l1.47 1.47v-4.42a.75.75 0 1 1 1.5 0v4.37l1.47-1.47a.75.75 0 0 1 1.226.255a.74.74 0 0 1-.216.835z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Threads(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.742 8.181l-1.63.433a.5.5 0 0 1 0-.11a7.88 7.88 0 0 0-.795-1.872a5.665 5.665 0 0 0-3.079-2.465a7.92 7.92 0 0 0-2.112-.392a8.382 8.382 0 0 0-1.59 0a6.912 6.912 0 0 0-2.365.664a5.463 5.463 0 0 0-2.324 2.213a7.576 7.576 0 0 0-.805 2.093c-.147.584-.245 1.18-.291 1.78c-.036.58-.036 1.162 0 1.741a11.58 11.58 0 0 0 .462 3.018a7.224 7.224 0 0 0 1.007 2.083a5.373 5.373 0 0 0 2.505 1.922c.62.238 1.27.387 1.932.443a9.648 9.648 0 0 0 1.941 0a6.037 6.037 0 0 0 2.174-.584a4.759 4.759 0 0 0 1.921-1.65a3.36 3.36 0 0 0 .463-2.958a2.766 2.766 0 0 0-1.107-1.52l-.352-.25c0 .13 0 .24-.05.351c-.1.622-.309 1.22-.614 1.771a3.654 3.654 0 0 1-2.656 1.902a4.576 4.576 0 0 1-2.213-.111a3.48 3.48 0 0 1-1.54-.895a2.898 2.898 0 0 1-.835-1.781a3.019 3.019 0 0 1 1.389-2.948a4.567 4.567 0 0 1 1.63-.644a9.468 9.468 0 0 1 1.77-.12c.455.01.909.05 1.359.12h.08v-.06a2.94 2.94 0 0 0-.382-1.077a1.902 1.902 0 0 0-1.258-.885a3.21 3.21 0 0 0-1.871 0a2.163 2.163 0 0 0-1.097.815l-.05.14l-1.379-.955a.614.614 0 0 0 .07-.101a3.914 3.914 0 0 1 2.456-1.56a4.86 4.86 0 0 1 2.545.111a3.512 3.512 0 0 1 2.153 1.831c.253.499.42 1.036.493 1.59c0 .211.05.423.07.634v.08l.413.202a5.32 5.32 0 0 1 1.69 1.338a4.589 4.589 0 0 1 1.007 2.012c.082.44.109.89.08 1.338a5.212 5.212 0 0 1-1.288 3.13a6.812 6.812 0 0 1-3.732 2.243c-1.229.273-2.495.33-3.743.17a8.522 8.522 0 0 1-2.757-.804a7.042 7.042 0 0 1-2.807-2.485a8.945 8.945 0 0 1-1.137-2.576a12.839 12.839 0 0 1-.392-2.012A16.702 16.702 0 0 1 4 11.764c0-.574 0-1.158.1-1.741c.077-.659.198-1.31.363-1.952c.207-.821.526-1.61.946-2.345a7.043 7.043 0 0 1 4.024-3.25a9.105 9.105 0 0 1 1.851-.412a11.208 11.208 0 0 1 2.395 0c.94.102 1.857.353 2.716.745a7.264 7.264 0 0 1 3.11 2.636c.54.826.948 1.733 1.207 2.686c.005.02.015.037.03.05m-5.665 4.025h-.07A8.338 8.338 0 0 0 14 12.065a9.392 9.392 0 0 0-1.067 0a4.617 4.617 0 0 0-1.177.16a2.013 2.013 0 0 0-.925.544a1.31 1.31 0 0 0 .11 1.952c.226.179.486.309.765.382a3.22 3.22 0 0 0 1.278.06c.291-.036.574-.124.835-.26c.389-.235.692-.587.865-1.007c.225-.54.357-1.116.392-1.7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThreadsSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<defs><path id="mageThreadsSquare0" fill="currentColor" d="M14 12.31c-.022.42-.117.833-.28 1.22a1.601 1.601 0 0 1-.63.71c-.186.1-.39.165-.6.19a2.37 2.37 0 0 1-.92 0a1.57 1.57 0 0 1-.55-.27a.999.999 0 0 1-.08-1.42a1.49 1.49 0 0 1 .67-.38c.272-.085.555-.125.84-.12c.26-.015.52-.015.78 0c.242.019.482.052.72.1z"/></defs><use href="#mageThreadsSquare0"/><path fill="currentColor" d="M17 2H7a5 5 0 0 0-5 5v10a5 5 0 0 0 5 5h10a5 5 0 0 0 5-5V7a5 5 0 0 0-5-5M7.52 14.53A5 5 0 0 0 8.24 16a4 4 0 0 0 1.81 1.39c.447.166.915.274 1.39.32c.466.045.934.045 1.4 0a4.56 4.56 0 0 0 1.57-.41a3.58 3.58 0 0 0 1.39-1.2a2.42 2.42 0 0 0 .33-2.1a2 2 0 0 0-.8-1.09l-.2-.14c0 .09 0 .17-.05.25c-.07.45-.22.882-.44 1.28a2.628 2.628 0 0 1-1.92 1.34a3.3 3.3 0 0 1-1.59-.08A2.55 2.55 0 0 1 10 14.9a2.17 2.17 0 0 1-.61-1.29a2.2 2.2 0 0 1 1-2.12a3.289 3.289 0 0 1 1.2-.49c.423-.07.851-.1 1.28-.09a7.92 7.92 0 0 1 1 .09h.06a2.41 2.41 0 0 0-.27-.78a1.382 1.382 0 0 0-.89-.64a2.3 2.3 0 0 0-1.35 0a1.66 1.66 0 0 0-.79.59v.07l-1-.69v-.07a2.84 2.84 0 0 1 1.77-1.17a3.63 3.63 0 0 1 1.85.08a2.55 2.55 0 0 1 1.55 1.33c.176.359.295.744.35 1.14a3.606 3.606 0 0 1 .05.52l.3.14a4 4 0 0 1 1.22 1c.346.427.573.937.66 1.48c.071.328.095.665.07 1a3.75 3.75 0 0 1-.93 2.25a4.93 4.93 0 0 1-2.7 1.63a8.226 8.226 0 0 1-1.41.17a8 8 0 0 1-1.29-.05a6.319 6.319 0 0 1-2-.58a5.2 5.2 0 0 1-2-1.79a6.75 6.75 0 0 1-.83-1.86c-.134-.495-.231-1-.29-1.51V12c0-.42 0-.84.07-1.26a9.49 9.49 0 0 1 .23-1.41A6.31 6.31 0 0 1 7 7.67a5.09 5.09 0 0 1 2.86-2.35A7.43 7.43 0 0 1 11.2 5a7.61 7.61 0 0 1 1.72 0a6.35 6.35 0 0 1 2 .52a5.17 5.17 0 0 1 2.24 1.9A6.64 6.64 0 0 1 18 9.38l-1.18.32v-.08a5.562 5.562 0 0 0-.58-1.35A4.08 4.08 0 0 0 14 6.52a5.6 5.6 0 0 0-1.52-.29a7.33 7.33 0 0 0-1.15 0a5 5 0 0 0-1.7.48A3.93 3.93 0 0 0 8 8.31a5.76 5.76 0 0 0-.57 1.49a7.89 7.89 0 0 0-.21 1.29a10.38 10.38 0 0 0 0 1.25a8.55 8.55 0 0 0 .3 2.19"/><use href="#mageThreadsSquare0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThreeDboxSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.4" d="m7.104 15.81l4.013 2.407a1.742 1.742 0 0 0 1.766 0l4.012-2.407a1.732 1.732 0 0 0 .837-1.48V9.402a1.731 1.731 0 0 0-.837-1.479l-4.012-2.14a1.743 1.743 0 0 0-1.766 0l-4.013 2.14a1.731 1.731 0 0 0-.836 1.48v4.929a1.731 1.731 0 0 0 .836 1.479m10.399-7.247L12 11.866"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.4" d="M6.497 8.564L12 11.866v6.592"/><rect width="18.5" height="18.5" x="2.75" y="2.75" stroke-width="1.5" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThreeDboxSquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.25 2.02h-6.5A6.76 6.76 0 0 0 2 8.77v6.46a6.76 6.76 0 0 0 6.75 6.75h6.5A6.76 6.76 0 0 0 22 15.23v-6.5a6.76 6.76 0 0 0-6.75-6.71m1.64 12.83a.94.94 0 0 1-.36.38l-3.84 2.29v-5.23l4.33-2.61v4.67a1 1 0 0 1-.13.5m-9.79 0a1 1 0 0 1-.14-.51V9.67l4.34 2.6v5.26l-3.84-2.31a1 1 0 0 1-.36-.36zm4.91-8.57a1 1 0 0 1 .54.16l3.81 2L12 11.06L7.63 8.44l3.84-2a1 1 0 0 1 .54-.15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsDown(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M5.75 2.75H4.568c-.98 0-1.775.795-1.775 1.776v8.284c0 .98.795 1.775 1.775 1.775h1.184c.98 0 1.775-.794 1.775-1.775V4.526c0-.98-.795-1.776-1.775-1.776"/><path d="m21.16 11.757l-1.42-7.101a2.368 2.368 0 0 0-2.367-1.906h-7.48a2.367 2.367 0 0 0-2.367 2.367v7.101a3.231 3.231 0 0 0 1.184 2.367l.982 5.918a.887.887 0 0 0 1.278.65l1.1-.543a3.551 3.551 0 0 0 1.87-4.048l-.496-1.965h5.396a2.368 2.368 0 0 0 2.32-2.84"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsDownFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.89 11.6l-1.42-7.09a3.12 3.12 0 0 0-1.1-1.83a3.19 3.19 0 0 0-2-.68H9.89a3.11 3.11 0 0 0-2.21.91A2.48 2.48 0 0 0 5.75 2H4.56a2.52 2.52 0 0 0-2.52 2.52v8.29a2.52 2.52 0 0 0 2.52 2.52h1.2a2.48 2.48 0 0 0 1.85-.82c.12.167.257.321.41.46l.94 5.65a1.65 1.65 0 0 0 .87 1.2c.23.114.483.175.74.18a1.68 1.68 0 0 0 .74-.18l1.1-.54a4.29 4.29 0 0 0 2.26-4.91l-.26-1h4.44c.46.004.915-.1 1.33-.3a3.13 3.13 0 0 0 1.72-3.44zM6.76 12.81a1 1 0 0 1-1 1H4.57a1 1 0 0 1-1-1V4.52a1 1 0 0 1 1-1h1.19a1 1 0 0 1 1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsUp(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M5.75 9.415H4.568c-.98 0-1.775.794-1.775 1.775v8.284c0 .98.795 1.776 1.775 1.776h1.184c.98 0 1.775-.795 1.775-1.776V11.19c0-.98-.795-1.775-1.775-1.775"/><path d="m21.16 12.243l-1.42 7.101a2.367 2.367 0 0 1-2.367 1.906h-7.48a2.367 2.367 0 0 1-2.367-2.367v-7.101A3.231 3.231 0 0 1 8.71 9.415l.982-5.918a.888.888 0 0 1 1.278-.65l1.1.544a3.55 3.55 0 0 1 1.87 4.047l-.496 1.965h5.396a2.367 2.367 0 0 1 2.32 2.84"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThumbsUpFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.863 11.026c-.11-.45-.319-.87-.61-1.23a3 3 0 0 0-1.08-.84a3.09 3.09 0 0 0-1.34-.3h-4.43l.26-1a4.3 4.3 0 0 0-.26-2.85a4.37 4.37 0 0 0-2-2.05l-1.1-.54a1.67 1.67 0 0 0-1.48 0a1.65 1.65 0 0 0-.58.5a1.54 1.54 0 0 0-.29.7l-.94 5.65a3 3 0 0 0-.41.46a2.48 2.48 0 0 0-1.87-.9h-1.18a2.52 2.52 0 0 0-2.52 2.52v8.29a2.52 2.52 0 0 0 2.52 2.52h1.18a2.5 2.5 0 0 0 1.94-.91a3.13 3.13 0 0 0 2.21.91h7.52a3.12 3.12 0 0 0 2-.69a3.06 3.06 0 0 0 1.1-1.82l1.42-7.09a3.16 3.16 0 0 0-.06-1.33m-15.13 8.45a1 1 0 0 1-1 1h-1.18a1 1 0 0 1-1-1v-8.29a1 1 0 0 1 1-1h1.18a1 1 0 0 1 1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tiktok(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.357 7.75a.537.537 0 0 0-.495-.516a4.723 4.723 0 0 1-2.415-.938a4.85 4.85 0 0 1-1.887-3.3a.538.538 0 0 0-.517-.496h-2.108a.517.517 0 0 0-.517.527v12.59a2.794 2.794 0 0 1-2.974 2.762a2.815 2.815 0 0 1-2.51-3.711A2.836 2.836 0 0 1 9.93 12.78a.506.506 0 0 0 .558-.506V9.807s-.896-.063-1.202-.063a5.271 5.271 0 0 0-4.101 1.93a5.789 5.789 0 0 0-1.37 2.52a5.862 5.862 0 0 0 2.14 6.072A5.926 5.926 0 0 0 9.591 21.5a5.946 5.946 0 0 0 4.207-1.719a5.841 5.841 0 0 0 1.75-4.133V8.71a7.844 7.844 0 0 0 4.218 1.613a.517.517 0 0 0 .548-.527V7.751z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TiktokCircle(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2a10 10 0 1 0 10 10A10.01 10.01 0 0 0 12 2m5.939 7.713v.646a.37.37 0 0 1-.38.37a5.364 5.364 0 0 1-2.903-1.108v4.728a3.938 3.938 0 0 1-1.18 2.81a4.011 4.011 0 0 1-2.87 1.17a4.103 4.103 0 0 1-2.862-1.17a3.98 3.98 0 0 1-1.026-3.805c.159-.642.48-1.232.933-1.713a3.58 3.58 0 0 1 2.79-1.313h.82v1.703a.348.348 0 0 1-.39.348a1.918 1.918 0 0 0-1.23 3.631c.27.155.572.246.882.267c.24.01.48-.02.708-.092a1.928 1.928 0 0 0 1.313-1.816V5.754a.359.359 0 0 1 .359-.36h1.415a.359.359 0 0 1 .359.34a3.303 3.303 0 0 0 1.282 2.245a3.25 3.25 0 0 0 1.641.636a.37.37 0 0 1 .338.35z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trash(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M5.47 6.015v12.514a2.72 2.72 0 0 0 2.721 2.721h7.618a2.72 2.72 0 0 0 2.72-2.72V6.014m-15.235.001h17.412"/><path d="M8.735 6.015V4.382a1.632 1.632 0 0 1 1.633-1.632h3.264a1.632 1.632 0 0 1 1.633 1.632v1.633M9.824 16.992v-5.439m4.353 5.439v-5.439"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrashFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.725 5.275h-4.69v-.89a2.41 2.41 0 0 0-.7-1.68a2.38 2.38 0 0 0-1.69-.7h-3.26a2.38 2.38 0 0 0-1.69.7a2.4 2.4 0 0 0-.69 1.68v.89h-4.69a.75.75 0 1 0 0 1.5h1.42v11.76a3.45 3.45 0 0 0 1 2.46a3.5 3.5 0 0 0 2.45 1h7.62a3.5 3.5 0 0 0 2.45-1a3.452 3.452 0 0 0 1-2.46V6.775h1.43a.75.75 0 0 0 0-1.5zm-11.2-.89a.87.87 0 0 1 .26-.62a.9.9 0 0 1 .62-.26h3.26a.88.88 0 0 1 .63.26a.91.91 0 0 1 .26.62v.89h-5zm1.33 12.61a1 1 0 1 1-2 0v-5.43a1 1 0 0 1 2 0zm4.36 0a1 1 0 0 1-2 0v-5.43a1 1 0 0 1 2 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrashSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.4" d="M7.874 8.218v7.908a1.719 1.719 0 0 0 1.72 1.719h4.813a1.72 1.72 0 0 0 1.719-1.72V8.219m-9.627-.001h11.002m-7.564 0V7.187a1.032 1.032 0 0 1 1.032-1.032h2.062a1.031 1.031 0 0 1 1.032 1.032v1.031"/><rect width="18.5" height="18.5" x="2.75" y="2.75" stroke-width="1.5" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrashSquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.59 7.47v.56h-3.12v-.56a.57.57 0 0 1 .11-.28a.56.56 0 0 1 .27-.1h2.55a.31.31 0 0 1 .11.09a.47.47 0 0 1 .09.12a.32.32 0 0 1-.01.17"/><path fill="currentColor" d="M15.25 2h-6.5A6.76 6.76 0 0 0 2 8.75v6.5A6.76 6.76 0 0 0 8.75 22h6.5A6.75 6.75 0 0 0 22 15.25v-6.5A6.75 6.75 0 0 0 15.25 2m2.25 6.88a.51.51 0 0 1-.15.21a.64.64 0 0 1-.22.14a.57.57 0 0 1-.26 0h-.55v6.35a2.09 2.09 0 0 1-.63 1.51a2.07 2.07 0 0 1-1.51.63H9.93a2.15 2.15 0 0 1-2.14-2.14V9.21h-.54a.59.59 0 0 1-.23 0a.6.6 0 0 1-.2-.14a.66.66 0 0 1-.14-.2a.69.69 0 0 1 0-.24a.63.63 0 0 1 .19-.43a.59.59 0 0 1 .43-.19h2.42v-.29c.01-.407.17-.795.45-1.09a1.56 1.56 0 0 1 1.08-.44h1.83a1.54 1.54 0 0 1 1.08.45c.148.14.264.311.34.5c.081.182.119.38.11.58v.29h2.42a.61.61 0 0 1 .44.19a.58.58 0 0 1 .18.43a.57.57 0 0 1-.12.25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrashThree(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M5.47 6.015v12.514a2.72 2.72 0 0 0 2.721 2.721h7.618a2.72 2.72 0 0 0 2.72-2.72V6.014m-15.235.001h17.412"/><path d="M8.735 6.015V4.382a1.632 1.632 0 0 1 1.633-1.632h3.264a1.632 1.632 0 0 1 1.633 1.632v1.633m-5.984 6.081h5.438m-5.438 4.353h5.438"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrashThreeFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.71 5.275h-4.69v-.89a2.37 2.37 0 0 0-.7-1.68a2.39 2.39 0 0 0-1.68-.7h-3.26a2.38 2.38 0 0 0-1.69.7a2.41 2.41 0 0 0-.7 1.68v.89H3.3a.75.75 0 1 0 0 1.5h1.43v11.76a3.48 3.48 0 0 0 1 2.46a3.54 3.54 0 0 0 2.46 1h7.62a3.47 3.47 0 0 0 3.47-3.47V6.775h1.42a.75.75 0 0 0 0-1.5zm-11.22-.89a.9.9 0 0 1 .549-.813a.88.88 0 0 1 .341-.067h3.26a.9.9 0 0 1 .812.544c.045.106.068.22.068.336v.89h-5zm5.24 13.07H9.29a1 1 0 0 1 0-2h5.44a1 1 0 0 1 0 2m0-4.35H9.29a1 1 0 0 1 0-2h5.44a1 1 0 0 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrashTwo(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M5.47 6.015v12.514a2.72 2.72 0 0 0 2.721 2.721h7.618a2.72 2.72 0 0 0 2.72-2.72V6.014m-15.235.001h17.412"/><path d="M8.735 6.015V4.382a1.632 1.632 0 0 1 1.633-1.632h3.264a1.632 1.632 0 0 1 1.633 1.632v1.633"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrashTwoFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.705 5.27h-4.7v-.89a2.4 2.4 0 0 0-.69-1.68a2.38 2.38 0 0 0-1.69-.7h-3.26a2.44 2.44 0 0 0-1.69.7a2.4 2.4 0 0 0-.69 1.68v.89h-4.69a.75.75 0 0 0 0 1.5h1.42v11.76A3.47 3.47 0 0 0 8.185 22h7.62a3.47 3.47 0 0 0 3.47-3.47V6.77h1.43a.75.75 0 0 0 0-1.5m-6.2 0h-5v-.89a.9.9 0 0 1 .25-.62a.94.94 0 0 1 .63-.26h3.26a.88.88 0 0 1 .63.26a.89.89 0 0 1 .25.62z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trophy(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M15.08 2.752H8.92a1.97 1.97 0 0 0-2.062 1.85v6.077a5.494 5.494 0 0 0 3.024 4.8a4.523 4.523 0 0 0 4.236 0a5.494 5.494 0 0 0 3.024-4.8V4.602a1.97 1.97 0 0 0-2.062-1.85Z"/><path d="M17.142 4.602h2.054a1.97 1.97 0 0 1 2.053 1.85a6.806 6.806 0 0 1-.87 3.311a7.74 7.74 0 0 1-2.423 2.608l-.814.555l-.795.481M6.858 4.602H4.804a1.97 1.97 0 0 0-2.053 1.85c.003 1.16.302 2.3.87 3.311a7.74 7.74 0 0 0 2.423 2.608l.814.555l.795.481m2.294 5.068v-2.969m4.106 2.969v-2.969M8.81 18.475h6.38a1.952 1.952 0 0 1 1.952 1.952v.333a.49.49 0 0 1-.49.49H7.348a.49.49 0 0 1-.49-.49v-.333a1.952 1.952 0 0 1 1.951-1.952Z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrophyCircle(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M15.08 2.752H8.92a1.97 1.97 0 0 0-2.062 1.85v6.077a5.494 5.494 0 0 0 3.024 4.8a4.523 4.523 0 0 0 4.236 0a5.494 5.494 0 0 0 3.024-4.8V4.602a1.97 1.97 0 0 0-2.062-1.85Z"/><path stroke-linecap="round" d="M17.142 4.602h2.054a1.97 1.97 0 0 1 2.053 1.85a6.806 6.806 0 0 1-.87 3.311a7.74 7.74 0 0 1-2.423 2.608l-.814.555l-.795.481M6.858 4.602H4.804a1.97 1.97 0 0 0-2.053 1.85c.003 1.16.302 2.3.87 3.311a7.74 7.74 0 0 0 2.423 2.608l.814.555l.795.481m2.294 5.068v-2.969m4.106 2.969v-2.969M8.81 18.475h6.38a1.952 1.952 0 0 1 1.952 1.952v.333a.49.49 0 0 1-.49.49H7.348a.49.49 0 0 1-.49-.49v-.333a1.952 1.952 0 0 1 1.951-1.952Z"/><rect width="4" height="4" x="10" y="7.004" rx="2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrophyCircleFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.083 4.548a2.64 2.64 0 0 0-1.904-.69h-1.426a2.722 2.722 0 0 0-.768-1.161a2.676 2.676 0 0 0-1.924-.69H8.979a2.702 2.702 0 0 0-1.994.69a2.623 2.623 0 0 0-.748 1.16H4.831a2.643 2.643 0 0 0-1.934.69A2.731 2.731 0 0 0 2 6.46a7.607 7.607 0 0 0 .997 3.683a8.433 8.433 0 0 0 2.642 2.862l.848.57l.638.39a6.172 6.172 0 0 0 2.083 2.002v1.761H8.82a2.683 2.683 0 0 0-1.904.791a2.703 2.703 0 0 0-.788 1.91v.331a1.235 1.235 0 0 0 .359.88c.233.232.549.362.877.361h9.272a1.223 1.223 0 0 0 1.145-.765c.062-.15.093-.313.091-.476v-.33c0-.717-.283-1.404-.788-1.91a2.687 2.687 0 0 0-1.903-.792h-.39v-1.771a6.2 6.2 0 0 0 2.094-2.002l.658-.4l.808-.55a8.504 8.504 0 0 0 2.652-2.872A7.698 7.698 0 0 0 22 6.409a2.755 2.755 0 0 0-.917-1.861M4.303 9.42a6.071 6.071 0 0 1-.768-2.902A1.224 1.224 0 0 1 4.36 5.44c.151-.051.312-.072.471-.062h1.306v5.344c.008.285.032.569.07.85a6.776 6.776 0 0 1-1.904-2.15m7.707 2.362a2.734 2.734 0 0 1-1.523-.464A2.75 2.75 0 0 1 9.32 8.494a2.755 2.755 0 0 1 .75-1.409a2.738 2.738 0 0 1 2.988-.596a2.744 2.744 0 0 1 .897 4.461c-.505.51-1.19.802-1.906.813zm7.707-2.362a7.145 7.145 0 0 1-1.904 2.152c.038-.289.061-.58.07-.87V5.377h1.335a1.153 1.153 0 0 1 .868.31c.227.2.37.48.399.781a6.232 6.232 0 0 1-.768 2.952"/><path fill="currentColor" d="M12.01 10.262c.688 0 1.246-.56 1.246-1.251c0-.69-.558-1.25-1.246-1.25s-1.246.56-1.246 1.25s.558 1.25 1.246 1.25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrophyDown(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M15.08 2.752H8.92a1.97 1.97 0 0 0-2.062 1.85v6.077a5.494 5.494 0 0 0 3.024 4.8a4.523 4.523 0 0 0 4.236 0a5.494 5.494 0 0 0 3.024-4.8V4.602a1.97 1.97 0 0 0-2.062-1.85Z"/><path d="M17.142 4.602h2.054a1.97 1.97 0 0 1 2.053 1.85a6.806 6.806 0 0 1-.87 3.311a7.741 7.741 0 0 1-2.423 2.608l-.814.555l-.795.481M6.858 4.602H4.804a1.97 1.97 0 0 0-2.053 1.85c.003 1.16.302 2.3.87 3.311a7.74 7.74 0 0 0 2.423 2.608l.814.555l.795.481m2.294 5.068v-2.969m4.106 2.969v-2.969M8.81 18.475h6.38a1.952 1.952 0 0 1 1.952 1.952v.333a.49.49 0 0 1-.49.49H7.348a.49.49 0 0 1-.49-.49v-.333a1.952 1.952 0 0 1 1.951-1.952Z"/><path stroke-miterlimit="10" d="M12 11.504v-5"/><path stroke-linejoin="round" d="m9.707 9.399l1.967 1.967a.459.459 0 0 0 .652 0l1.967-1.967"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrophyDownFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.043 4.546a2.62 2.62 0 0 0-1.904-.69h-1.426a2.723 2.723 0 0 0-.768-1.162a2.676 2.676 0 0 0-1.924-.69H8.929a2.653 2.653 0 0 0-1.934.69c-.354.31-.616.712-.758 1.161H4.832a2.643 2.643 0 0 0-1.935.691A2.711 2.711 0 0 0 2 6.457a7.608 7.608 0 0 0 .997 3.683a8.534 8.534 0 0 0 2.642 2.863l.848.57l.638.39a6.171 6.171 0 0 0 2.083 2.002v1.762H8.82a2.692 2.692 0 0 0-2.692 2.702v.33a1.235 1.235 0 0 0 .761 1.15c.15.061.312.092.475.091h9.272a1.245 1.245 0 0 0 1.236-1.241v-.33a2.703 2.703 0 0 0-.787-1.912a2.716 2.716 0 0 0-1.905-.79h-.388v-1.772a6.202 6.202 0 0 0 2.093-2.001l.658-.4l.808-.551a8.503 8.503 0 0 0 2.652-2.873A7.699 7.699 0 0 0 22 6.407a2.755 2.755 0 0 0-.957-1.861M4.263 9.42a6.073 6.073 0 0 1-.767-2.902a1.225 1.225 0 0 1 .824-1.08c.151-.051.312-.072.472-.061h1.306v5.344c.008.284.031.569.07.85A6.778 6.778 0 0 1 4.263 9.42m10.519.53l-1.994 2.002a1.157 1.157 0 0 1-.698.34a.566.566 0 0 1-.3 0a1.154 1.154 0 0 1-.318-.08a1.396 1.396 0 0 1-.39-.26L9.09 9.95a.752.752 0 0 1 0-1.06a.766.766 0 0 1 1.067 0l.997 1V6.518a.752.752 0 0 1 .747-.751a.746.746 0 0 1 .748.75V9.89l.997-1a.746.746 0 0 1 1.257.535a.752.752 0 0 1-.2.525zm4.885-.53a6.997 6.997 0 0 1-1.894 2.152c.038-.29.061-.58.07-.871V5.377h1.335c.318-.01.628.1.868.31c.227.2.37.479.399.78a6.123 6.123 0 0 1-.778 2.953"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrophyFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.083 4.585a2.62 2.62 0 0 0-1.904-.69h-1.426a2.722 2.722 0 0 0-.768-1.162a2.676 2.676 0 0 0-1.924-.69H8.979a2.703 2.703 0 0 0-1.994.69a2.623 2.623 0 0 0-.748 1.161H4.831a2.643 2.643 0 0 0-1.934.69A2.731 2.731 0 0 0 2 6.497a7.607 7.607 0 0 0 .997 3.682a8.434 8.434 0 0 0 2.642 2.862l.848.57l.638.391a6.171 6.171 0 0 0 2.083 2.002v1.76H8.82c-.714 0-1.4.285-1.904.792a2.707 2.707 0 0 0-.788 1.91v.33a1.234 1.234 0 0 0 .359.881c.233.232.549.361.877.36h9.272a1.222 1.222 0 0 0 1.145-.764c.062-.15.093-.313.091-.476v-.33a2.71 2.71 0 0 0-.788-1.91a2.69 2.69 0 0 0-1.903-.792h-.39v-1.771a6.2 6.2 0 0 0 2.094-2.002l.658-.4l.808-.55a8.505 8.505 0 0 0 2.652-2.873A7.697 7.697 0 0 0 22 6.447a2.755 2.755 0 0 0-.917-1.861M4.303 9.458a6.071 6.071 0 0 1-.768-2.902a1.224 1.224 0 0 1 .825-1.08c.151-.05.312-.072.471-.06h1.306v5.343c.008.284.032.568.07.85a6.776 6.776 0 0 1-1.904-2.151m15.414 0a7.146 7.146 0 0 1-1.904 2.152c.038-.29.061-.58.07-.871V5.415h1.335a1.153 1.153 0 0 1 .868.31c.227.2.37.48.399.781a6.232 6.232 0 0 1-.768 2.952"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrophyStar(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M15.08 2.752H8.92a1.97 1.97 0 0 0-2.062 1.85v6.077a5.494 5.494 0 0 0 3.024 4.8a4.523 4.523 0 0 0 4.236 0a5.494 5.494 0 0 0 3.024-4.8V4.602a1.97 1.97 0 0 0-2.062-1.85Z"/><path d="M17.142 4.602h2.054a1.97 1.97 0 0 1 2.053 1.85a6.806 6.806 0 0 1-.87 3.311a7.74 7.74 0 0 1-2.423 2.608l-.814.555l-.795.481M6.858 4.602H4.804a1.97 1.97 0 0 0-2.053 1.85c.003 1.16.302 2.3.87 3.311a7.74 7.74 0 0 0 2.423 2.608l.814.555l.795.481m2.294 5.068v-2.969m4.106 2.969v-2.969M8.81 18.475h6.38a1.952 1.952 0 0 1 1.952 1.952v.333a.49.49 0 0 1-.49.49H7.348a.49.49 0 0 1-.49-.49v-.333a1.952 1.952 0 0 1 1.951-1.952Z"/><path stroke-linejoin="round" d="m12.143 10.308l1.18.62a.302.302 0 0 0 .437-.319l-.226-1.313a.302.302 0 0 1 .087-.268l.955-.93a.3.3 0 0 0-.166-.515l-1.319-.19a.3.3 0 0 1-.225-.165l-.603-1.205a.302.302 0 0 0-.542 0l-.602 1.205a.301.301 0 0 1-.226.165l-1.3.19a.301.301 0 0 0-.17.515l.955.93a.301.301 0 0 1 .088.268l-.226 1.313a.3.3 0 0 0 .436.32l1.18-.62a.3.3 0 0 1 .287 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrophyStarFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.083 4.548a2.619 2.619 0 0 0-1.904-.69h-1.426a2.722 2.722 0 0 0-.768-1.162a2.666 2.666 0 0 0-1.924-.69H8.979a2.713 2.713 0 0 0-1.994.69a2.623 2.623 0 0 0-.748 1.161H4.831a2.643 2.643 0 0 0-1.934.69A2.731 2.731 0 0 0 2 6.46a7.607 7.607 0 0 0 .997 3.683a8.433 8.433 0 0 0 2.642 2.862l.848.57l.638.39a6.172 6.172 0 0 0 2.083 2.002v1.76H8.82c-.714 0-1.4.286-1.904.792a2.707 2.707 0 0 0-.788 1.91v.331a1.235 1.235 0 0 0 .359.88c.233.232.549.362.877.361h9.272a1.223 1.223 0 0 0 1.145-.765c.062-.15.093-.313.091-.476v-.33c0-.717-.283-1.404-.788-1.91a2.687 2.687 0 0 0-1.903-.792h-.39v-1.771a6.2 6.2 0 0 0 2.094-2.002l.658-.4l.808-.55a8.505 8.505 0 0 0 2.652-2.872A7.697 7.697 0 0 0 22 6.409a2.755 2.755 0 0 0-.917-1.861M4.303 9.42a6.072 6.072 0 0 1-.768-2.902a1.224 1.224 0 0 1 .825-1.08c.151-.05.312-.072.471-.06h1.306v5.343c.008.285.032.569.07.85a6.776 6.776 0 0 1-1.904-2.15m10.558-1.22a.832.832 0 0 1-.239.44l-.658.64l.16.92a.921.921 0 0 1-.349.87a.996.996 0 0 1-.469.161h-.06a.875.875 0 0 1-.418-.11l-.818-.43l-.828.43a.785.785 0 0 1-.478.1a.856.856 0 0 1-.758-.55a.914.914 0 0 1-.05-.49l.15-.921l-.658-.64a.841.841 0 0 1-.24-.43a.783.783 0 0 1 0-.501a.75.75 0 0 1 .27-.39a.846.846 0 0 1 .448-.21l.898-.13l.428-.851a.77.77 0 0 1 .32-.35a.885.885 0 0 1 .937 0a.77.77 0 0 1 .329.36l.418.84l.928.13a.916.916 0 0 1 .707.61a.753.753 0 0 1 .03.481zm4.856 1.22a7.146 7.146 0 0 1-1.904 2.152c.038-.29.061-.58.07-.871V5.378h1.335a1.153 1.153 0 0 1 .868.31c.227.2.37.48.399.781a6.233 6.233 0 0 1-.768 2.952"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrophyUp(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M15.08 2.752H8.92a1.97 1.97 0 0 0-2.062 1.85v6.077a5.494 5.494 0 0 0 3.024 4.8a4.523 4.523 0 0 0 4.236 0a5.494 5.494 0 0 0 3.024-4.8V4.602a1.97 1.97 0 0 0-2.062-1.85Z"/><path d="M17.142 4.602h2.054a1.97 1.97 0 0 1 2.053 1.85a6.806 6.806 0 0 1-.87 3.311a7.74 7.74 0 0 1-2.423 2.608l-.814.555l-.795.481M6.858 4.602H4.804a1.97 1.97 0 0 0-2.053 1.85c.003 1.16.302 2.3.87 3.311a7.74 7.74 0 0 0 2.423 2.608l.814.555l.795.481m2.294 5.068v-2.969m4.106 2.969v-2.969M8.81 18.475h6.38a1.952 1.952 0 0 1 1.952 1.952v.333a.49.49 0 0 1-.49.49H7.348a.49.49 0 0 1-.49-.49v-.333a1.952 1.952 0 0 1 1.951-1.952Z"/><path stroke-miterlimit="10" d="M12 6.504v5"/><path stroke-linejoin="round" d="M14.293 8.609L12.326 6.64a.458.458 0 0 0-.652 0L9.707 8.61"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrophyUpFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.043 4.546a2.62 2.62 0 0 0-1.904-.69h-1.426a2.723 2.723 0 0 0-.768-1.162a2.676 2.676 0 0 0-1.924-.69H8.929a2.653 2.653 0 0 0-1.934.69c-.354.31-.616.712-.758 1.161H4.832a2.643 2.643 0 0 0-1.935.691A2.711 2.711 0 0 0 2 6.457a7.608 7.608 0 0 0 .997 3.683a8.534 8.534 0 0 0 2.642 2.863l.848.57l.638.39a6.171 6.171 0 0 0 2.083 2.002v1.762H8.82a2.692 2.692 0 0 0-2.692 2.702v.33a1.235 1.235 0 0 0 .761 1.15c.15.061.312.092.475.091h9.272a1.245 1.245 0 0 0 1.236-1.241v-.33a2.703 2.703 0 0 0-.787-1.912a2.716 2.716 0 0 0-1.905-.79h-.388v-1.772a6.202 6.202 0 0 0 2.093-2.001l.658-.4l.808-.551a8.503 8.503 0 0 0 2.652-2.873A7.699 7.699 0 0 0 22 6.407a2.755 2.755 0 0 0-.957-1.861M4.263 9.42a6.073 6.073 0 0 1-.767-2.902a1.205 1.205 0 0 1 .398-.831a1.205 1.205 0 0 1 .898-.31h1.306v5.344c.008.284.031.569.07.85A6.778 6.778 0 0 1 4.263 9.42m10.519-.26a.746.746 0 0 1-1.057 0l-.997-1.001v3.383a.752.752 0 0 1-.748.75a.746.746 0 0 1-.748-.75V8.159l-.997 1a.75.75 0 0 1-1.284-.526c0-.2.077-.392.218-.534l1.993-2.002c.116-.11.252-.198.4-.26a1.226 1.226 0 0 1 .917 0c.145.06.278.148.388.26L14.861 8.1a.752.752 0 0 1-.08 1.04zm4.885.26a6.997 6.997 0 0 1-1.894 2.152c.038-.29.061-.58.07-.871V5.377h1.335c.318-.01.628.1.868.31c.227.2.37.479.399.78a6.123 6.123 0 0 1-.778 2.953"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tube(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18.713 5.278A9.5 9.5 0 1 0 5.287 18.722A9.5 9.5 0 0 0 18.713 5.278M12 16.214a4.222 4.222 0 1 1 .007-8.445A4.222 4.222 0 0 1 12 16.214m6.713-10.936l-3.736 3.737m3.736 9.69l-3.736-3.737m-5.954 0l-3.736 3.737m3.736-9.69L5.287 5.278"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TubeFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.121 12.95c.12.636.369 1.242.73 1.78l-2.7 2.7l-1.07 1.06a10.119 10.119 0 0 1-1.55-2.56a10.23 10.23 0 0 1-.34-6.9a10 10 0 0 1 1.9-3.54l3.76 3.76a3 3 0 0 0-.23.38a5 5 0 0 0-.5 3.32m11.34 6.96a10.221 10.221 0 0 1-11.33 1.13a10.12 10.12 0 0 1-1.68-1.13l3.77-3.77a5 5 0 0 0 5.48 0l2.76 2.69zm3.79-7.92a10.34 10.34 0 0 1-2.32 6.5l-3.77-3.77a5 5 0 0 0 .8-2.24a5 5 0 0 0-.8-3.24l3.77-3.77a10.27 10.27 0 0 1 2.32 6.52m-3.79-7.91l-3.72 3.75a5.18 5.18 0 0 0-1.28-.6a5 5 0 0 0-3.35.16a4.23 4.23 0 0 0-.84.45l-2.71-2.7l-1.1-1.06a7.58 7.58 0 0 1 .8-.6a10.23 10.23 0 0 1 12.19.6zm-3 8.26a3.46 3.46 0 0 1-.4 1.29a3.2 3.2 0 0 1-.6.82c-.24.237-.513.439-.81.6a3.439 3.439 0 0 1-1.64.41a3.54 3.54 0 0 1-1.65-.41a4.233 4.233 0 0 1-.56-.38a3.24 3.24 0 0 1-.84-1a3.48 3.48 0 0 1 1.73-4.88a3.42 3.42 0 0 1 2.34-.11a3.5 3.5 0 0 1 1.87 1.4c.063.088.117.181.16.28a3.29 3.29 0 0 1 .4 1.98"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Twitter(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.98 4.003a16.56 16.56 0 0 1-2.66 1.015a4.216 4.216 0 0 0-3.698-1.28a4.316 4.316 0 0 0-3.477 4.945a.393.393 0 0 0 0 .11a11.878 11.878 0 0 1-8.666-4.338a4.184 4.184 0 0 0 1.292 5.597a4.14 4.14 0 0 1-1.899-.519v.056a4.228 4.228 0 0 0 3.312 4.117c-.361.09-.732.135-1.104.133a3.744 3.744 0 0 1-.795-.066a4.228 4.228 0 0 0 3.919 2.914a8.467 8.467 0 0 1-5.2 1.788A7.562 7.562 0 0 1 2 18.42a11.734 11.734 0 0 0 6.425 1.888A11.855 11.855 0 0 0 20.457 8.374v-.54a4.549 4.549 0 0 0 1.524-3.831"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Twtich(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.899 6.926h-1.455v4.21H16.9zm-3.795 0H11.65v4.21h1.454z"/><path fill="currentColor" d="M20.398 3H7.156L3.547 6.543v12.674h4.167V23l3.695-3.74h2.778l6.266-6.276zM8.053 14.733V4.771H18.79v7.316l-2.68 2.679h-2.733l-2.187 2.11V14.69z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unlocked(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.376 9.69V7.378a4.624 4.624 0 0 1 7.895-3.272c.196.193.37.406.52.636m-7.259 13.04h6.936"/><path d="M17 9.688H7c-1.38 0-2.5 1.035-2.5 2.312v6.938c0 1.277 1.12 2.312 2.5 2.312h10c1.38 0 2.5-1.035 2.5-2.312V12c0-1.277-1.12-2.312-2.5-2.312"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnlockedFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.96 8.951H8.08v-1.56c0-.766.23-1.515.66-2.15a3.87 3.87 0 0 1 4-1.65a3.8 3.8 0 0 1 2 1.06c.159.157.3.332.42.52a.754.754 0 0 0 1.26-.83a4.937 4.937 0 0 0-.62-.75a5.26 5.26 0 0 0-2.75-1.47a5.38 5.38 0 0 0-6.43 5.27v1.59a3.12 3.12 0 0 0-2.87 3v6.94A3.16 3.16 0 0 0 7 21.981h10a3.16 3.16 0 0 0 3.25-3.06v-6.94a3.16 3.16 0 0 0-3.29-3.03m-1.53 9.84H8.49a1 1 0 0 1 0-2h6.94a1 1 0 1 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Upload(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="M12 3.212v12.026"/><path stroke-linejoin="round" d="M16.625 7.456L12.66 3.49a.937.937 0 0 0-1.318 0L7.375 7.456M2.75 13.85v4.625a2.312 2.312 0 0 0 2.313 2.313h13.875a2.312 2.312 0 0 0 2.312-2.313V13.85"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func User(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.618 21.25c0-3.602-4.016-6.53-7.618-6.53c-3.602 0-7.618 2.928-7.618 6.53M12 11.456a4.353 4.353 0 1 0 0-8.706a4.353 4.353 0 0 0 0 8.706"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserCheck(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.359 14.724c-3.6 0-7.62 2.928-7.62 6.526m7.62-9.785a4.362 4.362 0 0 0 4.035-2.683a4.355 4.355 0 0 0-3.171-5.948a4.362 4.362 0 0 0-5.215 4.274a4.356 4.356 0 0 0 4.35 4.357m.904 6.897l1.688 1.689a.637.637 0 0 0 .909 0l3.403-3.403"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserCheckFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.865 21.25a.75.75 0 0 1-.22.53a.79.79 0 0 1-.53.22H4.481a.75.75 0 0 1-.75-.75c0-4.102 4.491-7.282 8.372-7.282a.76.76 0 0 1 .75.75zm4.001-12.174a5.072 5.072 0 0 1-4.711 3.141a5.122 5.122 0 0 1-5.102-5.112a5.08 5.08 0 0 1 .87-2.84a5.081 5.081 0 0 1 5.252-2.161c.991.192 1.902.68 2.61 1.4a5.051 5.051 0 0 1 1.4 2.621a5.14 5.14 0 0 1-.32 2.951m-.389 11.534a1.38 1.38 0 0 1-.48-.09a1.3 1.3 0 0 1-.41-.28l-1.33-1.33a.75.75 0 1 1 1.06-1.061l1.17 1.17l2.52-2.53a.75.75 0 0 1 1.06 1.06l-2.7 2.69a1.09 1.09 0 0 1-.4.27a1.191 1.191 0 0 1-.49.1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserCircle(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 21.5a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19"/><path d="M6.374 19.653a6.333 6.333 0 0 1 11.252 0M12 13.056a3.399 3.399 0 1 0 0-6.798a3.399 3.399 0 0 0 0 6.798"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserCircleFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.967 1.752c-2.15.01-4.244.695-5.984 1.957a10.234 10.234 0 0 0-.126 16.493c.05.049.107.09.17.12a10.228 10.228 0 0 0 11.95 0a.831.831 0 0 0 .18-.13a10.235 10.235 0 0 0-.172-16.506a10.278 10.278 0 0 0-5.998-1.934zm0 3.76a4.162 4.162 0 0 1 3.878 2.534a4.144 4.144 0 0 1-.882 4.543A4.158 4.158 0 0 1 7.86 9.632a4.147 4.147 0 0 1 1.21-2.898a4.16 4.16 0 0 1 2.897-1.222m4.627 13.92a8.754 8.754 0 0 1-9.245 0a8.094 8.094 0 0 1-1.212-.9a7.126 7.126 0 0 1 2.144-2a7.23 7.23 0 0 1 7.382 0a7.125 7.125 0 0 1 2.143 2c-.375.337-.78.638-1.212.9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserCross(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M11.959 14.724c-3.6 0-7.62 2.928-7.62 6.526m7.62-9.785a4.362 4.362 0 0 0 4.035-2.683a4.355 4.355 0 0 0-3.17-5.948a4.362 4.362 0 0 0-5.215 4.274a4.356 4.356 0 0 0 4.35 4.357"/><path stroke-miterlimit="10" d="m19.661 15.487l-5 4.989m0-4.978l5 4.989"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserCrossFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.975 14.718v6.53a.76.76 0 0 1-.75.75h-7.62a.75.75 0 0 1-.75-.75c0-4.1 4.49-7.28 8.37-7.28a.76.76 0 0 1 .75.75m3.97-5.64a5.191 5.191 0 0 1-1.88 2.29a5.11 5.11 0 0 1-5.451.121a5.1 5.1 0 0 1-1.609-7.211a5.09 5.09 0 0 1 2.29-1.89a5.17 5.17 0 0 1 3-.28a5 5 0 0 1 2.61 1.4a5.05 5.05 0 0 1 1.4 2.62a5.14 5.14 0 0 1-.36 2.95m2.98 10.34a.75.75 0 0 1-1.06 1.06l-1.44-1.44l-1.44 1.44a.77.77 0 0 1-.817.154a.75.75 0 0 1-.242-1.224l1.43-1.43l-1.43-1.43a.75.75 0 0 1 1.06-1.06l1.43 1.43l1.45-1.44a.75.75 0 0 1 1.06 1.06l-1.44 1.44z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.37 21.25a.75.75 0 0 1-.75.75H4.38a.75.75 0 0 1-.75-.75c0-4.1 4.5-7.28 8.37-7.28c3.87 0 8.37 3.18 8.37 7.28M17.1 7.11A5.1 5.1 0 1 1 12 2a5.11 5.11 0 0 1 5.1 5.11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserMinus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M12.125 14.719c-3.6 0-7.62 2.928-7.62 6.526m7.62-9.785a4.361 4.361 0 0 0 4.035-2.683a4.355 4.355 0 0 0-3.17-5.948a4.362 4.362 0 0 0-5.215 4.274a4.356 4.356 0 0 0 4.35 4.357"/><path stroke-miterlimit="10" d="M13.495 17.988h6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserMinusFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.166 21.25a.751.751 0 0 1-.75.75h-7.64a.75.75 0 0 1-.75-.75c0-4.105 4.504-7.289 8.378-7.289a.75.75 0 0 1 .751.751zM17.14 9.076a5.005 5.005 0 0 1-1.892 2.293a5.095 5.095 0 0 1-6.446-.638A5.106 5.106 0 0 1 7.31 7.114a5.105 5.105 0 0 1 .861-2.843a5.165 5.165 0 0 1 2.302-1.882A5.115 5.115 0 0 1 16.05 3.5a5.005 5.005 0 0 1 1.392 2.623c.197.992.092 2.02-.3 2.953m2.081 9.671h-3.934a.75.75 0 0 1 0-1.502h3.934a.751.751 0 0 1 0 1.502"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserPlus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M12.125 14.719c-3.6 0-7.62 2.928-7.62 6.526m7.62-9.785a4.362 4.362 0 0 0 4.035-2.683a4.355 4.355 0 0 0-3.17-5.948a4.362 4.362 0 0 0-5.215 4.274a4.356 4.356 0 0 0 4.35 4.357"/><path stroke-miterlimit="10" d="M16.488 14.983v5.997m-2.993-2.992h6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserPlusFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.181 21.25a.751.751 0 0 1-.75.75H4.792a.75.75 0 0 1-.75-.75c0-4.105 4.504-7.289 8.379-7.289a.75.75 0 0 1 .75.751zm3.975-12.174a5.005 5.005 0 0 1-1.892 2.293a5.095 5.095 0 0 1-6.447-.638a5.106 5.106 0 0 1-1.492-3.617a5.105 5.105 0 0 1 .861-2.843a5.165 5.165 0 0 1 2.302-1.882A5.116 5.116 0 0 1 16.064 3.5a5.005 5.005 0 0 1 1.392 2.623c.197.992.092 2.02-.3 2.953m2.802 8.91a.739.739 0 0 1-.75.75h-1.362v1.352a.75.75 0 0 1-1.501 0v-1.361h-1.362a.75.75 0 0 1 0-1.502h1.402v-1.362a.751.751 0 0 1 1.501 0v1.362h1.362a.751.751 0 0 1 .71.76"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserQuestionMark(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M11.959 14.724c-3.6 0-7.62 2.928-7.62 6.526m7.62-9.785a4.362 4.362 0 0 0 4.035-2.683a4.355 4.355 0 0 0-3.17-5.948a4.362 4.362 0 0 0-5.215 4.274a4.356 4.356 0 0 0 4.35 4.357"/><path stroke-miterlimit="10" d="M15.318 15.92a1.599 1.599 0 0 1 1.742-.907a1.55 1.55 0 0 1 1.137.81a1.347 1.347 0 0 1-.784 1.851a.993.993 0 0 0-.64.898v.37"/><path stroke-linejoin="round" d="M16.745 20.987h.002"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserQuestionMarkFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.37 21.248a.75.75 0 0 1-.75.75H4.99a.75.75 0 0 1-.75-.75c0-4.1 4.49-7.28 8.37-7.28a.76.76 0 0 1 .75.75zm3.96-12.17a5.19 5.19 0 0 1-1.88 2.29a5.11 5.11 0 0 1-5.452.121A5.1 5.1 0 0 1 8.39 4.278a5.09 5.09 0 0 1 2.29-1.89a5.17 5.17 0 0 1 3-.28a5 5 0 0 1 2.61 1.4a5.05 5.05 0 0 1 1.4 2.62a5.14 5.14 0 0 1-.36 2.95m.09 10.61a.76.76 0 0 1-.75-.75v-.37a1.74 1.74 0 0 1 1.12-1.6a.61.61 0 0 0 .24-.14a.5.5 0 0 0 .15-.2a.53.53 0 0 0 0-.25a.691.691 0 0 0 0-.24a.81.81 0 0 0-.58-.39a.81.81 0 0 0-.56.08a.82.82 0 0 0-.38.4a.754.754 0 0 1-1.37-.63a2.36 2.36 0 0 1 2.56-1.33c.36.05.704.187 1 .4c.291.21.53.483.7.8a2 2 0 0 1 .21.88a2.05 2.05 0 0 1-.67 1.58a2 2 0 0 1-.75.45a.2.2 0 0 0-.08.08a.21.21 0 0 0 0 .13v.35a.76.76 0 0 1-.84.75m-.02 2.15a.78.78 0 1 0 0-1.56a.78.78 0 0 0 0 1.56"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M6.022 20.504c.284-1.394.974-2.138 2.076-3.038a6.167 6.167 0 0 1 7.805 0c1.101.9 1.882 1.644 2.165 3.038M12 13.028a3.31 3.31 0 1 0 0-6.619a3.31 3.31 0 0 0 0 6.619"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserSquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.257 2H8.753A6.765 6.765 0 0 0 2 8.75v6.5a6.727 6.727 0 0 0 3.122 5.68a5.9 5.9 0 0 0 1.06.56a6.646 6.646 0 0 0 2.561.51h6.504c.9 0 1.791-.18 2.62-.53a6.503 6.503 0 0 0 1.131-.62A6.71 6.71 0 0 0 22 15.26v-6.5A6.758 6.758 0 0 0 15.257 2m-3.252 4.58a3.143 3.143 0 0 1 3.081 3.753a3.14 3.14 0 0 1-4.283 2.288a3.142 3.142 0 0 1-1.94-2.901a3.15 3.15 0 0 1 3.142-3.14m5.002 13.63a5.006 5.006 0 0 1-1.7.29H8.803a5.264 5.264 0 0 1-3.391-1.25a6.53 6.53 0 0 1 2.1-2.56a7.176 7.176 0 0 1 9.085 0a6.871 6.871 0 0 1 2.151 2.52c-.523.45-.828.698-1.486.907z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Users(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17.928 19.634h2.138a1.165 1.165 0 0 0 1.116-1.555a6.851 6.851 0 0 0-6.117-3.95m0-2.759a3.664 3.664 0 0 0 3.665-3.664a3.664 3.664 0 0 0-3.665-3.674m-1.04 16.795a1.908 1.908 0 0 0 1.537-3.035a8.026 8.026 0 0 0-6.222-3.196a8.026 8.026 0 0 0-6.222 3.197a1.909 1.909 0 0 0 1.536 3.034zM9.34 11.485a4.16 4.16 0 0 0 4.15-4.161a4.151 4.151 0 0 0-8.302 0a4.16 4.16 0 0 0 4.151 4.16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UsersFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.987 18.73a1.999 1.999 0 0 1-.34.85a1.9 1.9 0 0 1-1.56.8h-1.651a.741.741 0 0 1-.6-.31a.758.758 0 0 1-.11-.67c.37-1.18.29-2.51-3.061-4.64a.77.77 0 0 1-.32-.85a.76.76 0 0 1 .72-.54a7.614 7.614 0 0 1 6.792 4.39a2 2 0 0 1 .13.97M19.486 7.7a4.43 4.43 0 0 1-4.421 4.42a.761.761 0 0 1-.65-1.13a6.158 6.158 0 0 0 0-6.53a.75.75 0 0 1 .61-1.18a4.292 4.292 0 0 1 3.13 1.34a4.46 4.46 0 0 1 1.291 3.12z"/><path fill="currentColor" d="M16.675 18.7a2.649 2.649 0 0 1-1.26 2.48c-.418.257-.9.392-1.39.39H4.652a2.631 2.631 0 0 1-1.39-.39A2.62 2.62 0 0 1 2.01 18.7a2.62 2.62 0 0 1 .5-1.35a8.792 8.792 0 0 1 6.812-3.51a8.775 8.775 0 0 1 6.842 3.5a2.7 2.7 0 0 1 .51 1.36M14.245 7.32a4.92 4.92 0 0 1-4.902 4.91a4.903 4.903 0 0 1-4.797-5.858a4.9 4.9 0 0 1 6.678-3.57a4.902 4.902 0 0 1 3.03 4.518z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VerifiedCheck(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m12.717 3.656l1.137-.904a1.134 1.134 0 0 1 1.808.584l.432 1.51a1.137 1.137 0 0 0 1.137.836h1.467a1.13 1.13 0 0 1 .96.449a1.149 1.149 0 0 1 .178 1.05l-.546 1.626a1.152 1.152 0 0 0 .41 1.293l1.33.973a1.143 1.143 0 0 1 .47.927a1.15 1.15 0 0 1-.47.927l-1.33.973a1.145 1.145 0 0 0-.41 1.293l.546 1.626a1.152 1.152 0 0 1-.602 1.394a1.13 1.13 0 0 1-.536.105h-1.467a1.132 1.132 0 0 0-.712.22a1.144 1.144 0 0 0-.425.616l-.432 1.51a1.147 1.147 0 0 1-.748.782a1.13 1.13 0 0 1-1.06-.198l-1.137-.904a1.132 1.132 0 0 0-1.434 0l-1.137.904a1.135 1.135 0 0 1-1.808-.584l-.432-1.51a1.145 1.145 0 0 0-.425-.617a1.132 1.132 0 0 0-.712-.219H5.302a1.13 1.13 0 0 1-.96-.449a1.148 1.148 0 0 1-.178-1.05l.546-1.626A1.152 1.152 0 0 0 4.3 13.9l-1.33-.973A1.143 1.143 0 0 1 2.5 12a1.15 1.15 0 0 1 .47-.927L4.3 10.1a1.144 1.144 0 0 0 .41-1.293l-.477-1.66a1.152 1.152 0 0 1 .602-1.394a1.13 1.13 0 0 1 .535-.105h1.467a1.145 1.145 0 0 0 1.137-.836l.432-1.51A1.145 1.145 0 0 1 9.17 2.6a1.13 1.13 0 0 1 1.011.209l1.138.904a1.132 1.132 0 0 0 1.399-.057"/><path d="m8.106 11.894l2.192 2.192a.829.829 0 0 0 1.18 0l4.417-4.418"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VerifiedCheckFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.02 11.164a1.84 1.84 0 0 0-.57-.67l-1.33-1a.35.35 0 0 1-.14-.2a.36.36 0 0 1 0-.25l.55-1.63a2 2 0 0 0 .06-.9a1.809 1.809 0 0 0-.36-.84a1.859 1.859 0 0 0-.7-.57a1.75 1.75 0 0 0-.85-.17h-1.5a.41.41 0 0 1-.39-.3l-.43-1.5a1.88 1.88 0 0 0-.46-.81a2 2 0 0 0-.78-.49a2 2 0 0 0-.92-.06a1.88 1.88 0 0 0-.83.39l-1.14.9a.35.35 0 0 1-.23.09a.36.36 0 0 1-.22-.05l-1.13-.9a1.85 1.85 0 0 0-.8-.38a1.87 1.87 0 0 0-.88 0a1.93 1.93 0 0 0-.78.43a2.08 2.08 0 0 0-.51.79l-.43 1.51a.38.38 0 0 1-.15.22a.41.41 0 0 1-.27.07H5.41a1.92 1.92 0 0 0-.89.18a1.78 1.78 0 0 0-.71.57a1.93 1.93 0 0 0-.36.83c-.05.293-.03.595.06.88L4 8.993a.41.41 0 0 1-.14.45l-1.33 1c-.242.18-.44.412-.58.68a1.93 1.93 0 0 0 0 1.71a2 2 0 0 0 .58.68l1.33 1a.41.41 0 0 1 .14.45l-.55 1.63a2 2 0 0 0-.07.91c.05.298.174.58.36.82c.183.25.428.45.71.58c.265.126.557.184.85.17h1.49a.38.38 0 0 1 .25.08a.34.34 0 0 1 .14.21l.43 1.51a2 2 0 0 0 .46.8a1.89 1.89 0 0 0 2.54.17l1.15-.91a.39.39 0 0 1 .49 0l1.13.9c.24.202.53.337.84.39c.113.01.227.01.34 0a1.9 1.9 0 0 0 .58-.09a1.871 1.871 0 0 0 1.24-1.28l.44-1.52a.34.34 0 0 1 .14-.21a.4.4 0 0 1 .27-.08h1.43a2 2 0 0 0 .89-.17a1.911 1.911 0 0 0 1.06-1.4a1.92 1.92 0 0 0-.07-.92l-.54-1.62a.36.36 0 0 1 0-.25a.35.35 0 0 1 .14-.2l1.33-1a1.87 1.87 0 0 0 .57-.68a1.82 1.82 0 0 0 .21-.86a1.881 1.881 0 0 0-.23-.78m-5.44-.76l-4.42 4.42a2.011 2.011 0 0 1-.59.4c-.222.09-.46.138-.7.14a1.711 1.711 0 0 1-.71-.15a1.863 1.863 0 0 1-.6-.4l-2.18-2.19a1 1 0 0 1 1.41-1.41l2.08 2.08l4.3-4.31a1 1 0 0 1 1.41 0a.998.998 0 0 1 0 1.46z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Video(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 5.32H6.095A3.595 3.595 0 0 0 2.5 8.923v6.162a3.595 3.595 0 0 0 3.595 3.595H12a3.595 3.595 0 0 0 3.595-3.595V8.924A3.594 3.594 0 0 0 12 5.319m9.5 4.119v5.135c0 .25-.071.496-.205.708a1.355 1.355 0 0 1-.555.493a1.27 1.27 0 0 1-.73.124a1.367 1.367 0 0 1-.677-.278l-3.225-2.588a1.376 1.376 0 0 1-.503-1.047c0-.2.045-.397.133-.575c.092-.168.218-.315.37-.432l3.225-2.567a1.36 1.36 0 0 1 .678-.278c.25-.032.504.011.729.123a1.336 1.336 0 0 1 .76 1.182"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoCheck(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m6.547 11.932l1.407 1.407a.531.531 0 0 0 .758 0l2.835-2.836"/><path d="M12 5.32H6.095A3.595 3.595 0 0 0 2.5 8.923v6.162a3.595 3.595 0 0 0 3.595 3.595H12a3.595 3.595 0 0 0 3.595-3.595V8.924A3.594 3.594 0 0 0 12 5.32m9.5 4.118v5.135c0 .25-.071.496-.205.708a1.355 1.355 0 0 1-.555.493a1.27 1.27 0 0 1-.73.124a1.366 1.366 0 0 1-.677-.278l-3.225-2.588a1.377 1.377 0 0 1-.503-1.047c0-.2.045-.396.133-.575c.092-.168.218-.315.37-.432l3.225-2.567a1.36 1.36 0 0 1 .678-.278c.25-.032.504.011.729.124a1.325 1.325 0 0 1 .76 1.181"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoCheckFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.93 8.33a2.06 2.06 0 0 0-.86-.77a2.11 2.11 0 0 0-2.2.24l-2.52 2v-.9a4.37 4.37 0 0 0-2.68-4A4.4 4.4 0 0 0 12 4.57H6.1a4.33 4.33 0 0 0-1.67.33a4.38 4.38 0 0 0-2.35 2.36a4.31 4.31 0 0 0-.33 1.66v6.17a4.34 4.34 0 0 0 4.35 4.34h5.83a4.37 4.37 0 0 0 4.017-2.677a4.32 4.32 0 0 0 .333-1.663v-.9l2.52 2c.324.256.718.41 1.13.44h.25a2 2 0 0 0 .9-.21a2.13 2.13 0 0 0 .85-.79a2.05 2.05 0 0 0 .32-1.11V9.37a2.12 2.12 0 0 0-.32-1.04m-9.85 2.68l-2.84 2.84a1.31 1.31 0 0 1-.91.38a1.38 1.38 0 0 1-.49-.1a1.42 1.42 0 0 1-.41-.28l-1.41-1.41a.739.739 0 0 1 0-1.06a.75.75 0 0 1 1.06 0l1.25 1.26l2.69-2.69a.75.75 0 0 1 1.06 1.06"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoCross(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="m11.047 10l-4 3.991m0-3.982l4 3.991"/><path stroke-linejoin="round" d="M12 5.32H6.095A3.595 3.595 0 0 0 2.5 8.923v6.162a3.595 3.595 0 0 0 3.595 3.595H12a3.595 3.595 0 0 0 3.595-3.595V8.924A3.594 3.594 0 0 0 12 5.32m9.5 4.118v5.135c0 .25-.071.496-.205.708a1.355 1.355 0 0 1-.555.493a1.27 1.27 0 0 1-.73.124a1.366 1.366 0 0 1-.677-.278l-3.225-2.588a1.377 1.377 0 0 1-.503-1.047c0-.2.045-.396.133-.575c.092-.168.218-.315.37-.432l3.225-2.567a1.36 1.36 0 0 1 .678-.278c.25-.032.504.011.729.124a1.325 1.325 0 0 1 .76 1.181"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoCrossFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.93 8.33a2.06 2.06 0 0 0-.86-.77a2.11 2.11 0 0 0-2.2.24l-2.52 2v-.9a4.37 4.37 0 0 0-2.68-4A4.4 4.4 0 0 0 12 4.57H6.1a4.33 4.33 0 0 0-1.67.33a4.38 4.38 0 0 0-2.35 2.36a4.31 4.31 0 0 0-.33 1.66v6.17a4.34 4.34 0 0 0 4.35 4.34h5.83a4.37 4.37 0 0 0 4.017-2.677a4.32 4.32 0 0 0 .333-1.663v-.9l2.52 2c.324.256.718.41 1.13.44h.25a2 2 0 0 0 .9-.21a2.13 2.13 0 0 0 .85-.79a2.05 2.05 0 0 0 .32-1.11V9.37a2.12 2.12 0 0 0-.32-1.04m-10.35 5.12a.75.75 0 0 1-1.06 1.06l-1.48-1.47l-1.46 1.46a.74.74 0 0 1-.53.22a.729.729 0 0 1-.53-.22a.739.739 0 0 1 0-1.06l1.46-1.46l-1.46-1.46a.75.75 0 0 1 1.06-1.06l1.46 1.46l1.48-1.47a.75.75 0 0 1 1.06 1.06l-1.47 1.47z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoDownload(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="M9.047 14.5v-5"/><path stroke-linejoin="round" d="m6.754 12.395l1.968 1.967a.458.458 0 0 0 .65 0l1.968-1.967"/><path stroke-linejoin="round" d="M12 5.32H6.095A3.595 3.595 0 0 0 2.5 8.923v6.162a3.595 3.595 0 0 0 3.595 3.595H12a3.595 3.595 0 0 0 3.595-3.595V8.924A3.594 3.594 0 0 0 12 5.32m9.5 4.118v5.135c0 .25-.071.496-.205.708a1.355 1.355 0 0 1-.555.493a1.27 1.27 0 0 1-.73.124a1.366 1.366 0 0 1-.677-.278l-3.225-2.588a1.377 1.377 0 0 1-.503-1.047c0-.2.045-.396.133-.575c.092-.168.218-.315.37-.432l3.225-2.567a1.36 1.36 0 0 1 .678-.278c.25-.032.504.011.729.124a1.325 1.325 0 0 1 .76 1.181"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoDownloadFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.93 8.33a2.06 2.06 0 0 0-.86-.77a2.11 2.11 0 0 0-2.2.24l-2.52 2v-.9a4.37 4.37 0 0 0-2.68-4A4.4 4.4 0 0 0 12 4.57H6.1a4.33 4.33 0 0 0-1.67.33a4.38 4.38 0 0 0-2.35 2.36a4.31 4.31 0 0 0-.33 1.66v6.17a4.34 4.34 0 0 0 4.35 4.34h5.83a4.37 4.37 0 0 0 4.017-2.677a4.32 4.32 0 0 0 .333-1.663v-.9l2.52 2c.289.23.634.378 1 .43h.25a2 2 0 0 0 .9-.21a2.13 2.13 0 0 0 .98-.78a2.05 2.05 0 0 0 .32-1.11V9.37a2.12 2.12 0 0 0-.32-1.04M11.87 12.9l-2 2a1.049 1.049 0 0 1-.39.27a1.22 1.22 0 0 1-.46.09a1.1 1.1 0 0 1-.47-.1a1.27 1.27 0 0 1-.39-.26l-2-2a.75.75 0 0 1 1.06-1.06l1 1V9.46a.75.75 0 0 1 1.5 0v3.38l1-1a.75.75 0 0 1 1.06 1.06z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.205 9.41v5.13c.01.382-.087.76-.28 1.09a2.13 2.13 0 0 1-.86.77a2 2 0 0 1-.9.21h-.25a2.07 2.07 0 0 1-1-.43l-2.53-2v.91a4.34 4.34 0 0 1-4.34 4.34h-5.91a4.37 4.37 0 0 1-3.07-1.27a4.31 4.31 0 0 1-1.27-3.07V8.92a4.298 4.298 0 0 1 .33-1.66a4.38 4.38 0 0 1 2.35-2.36a4.31 4.31 0 0 1 1.66-.33h5.79a4.4 4.4 0 0 1 1.67.33a4.38 4.38 0 0 1 2.35 2.36c.22.529.33 1.097.32 1.67v.9l2.53-2a2.09 2.09 0 0 1 3.06.53c.207.313.328.675.35 1.05"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoMinus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="M6.547 12.005h5"/><path stroke-linejoin="round" d="M12 5.32H6.095A3.595 3.595 0 0 0 2.5 8.923v6.162a3.595 3.595 0 0 0 3.595 3.595H12a3.595 3.595 0 0 0 3.595-3.595V8.924A3.594 3.594 0 0 0 12 5.32m9.5 4.118v5.135c0 .25-.071.496-.205.708a1.355 1.355 0 0 1-.555.493a1.27 1.27 0 0 1-.73.124a1.366 1.366 0 0 1-.677-.278l-3.225-2.588a1.377 1.377 0 0 1-.503-1.047c0-.2.045-.396.133-.575c.092-.168.218-.315.37-.432l3.225-2.567a1.36 1.36 0 0 1 .678-.278c.25-.032.504.011.729.124a1.325 1.325 0 0 1 .76 1.181"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoMinusFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.92 8.36a2.06 2.06 0 0 0-.86-.77a2.11 2.11 0 0 0-2.2.24l-2.53 2v-.9a4.14 4.14 0 0 0-.32-1.67a4.33 4.33 0 0 0-4-2.69H6.1a4.31 4.31 0 0 0-1.66.33a4.38 4.38 0 0 0-2.35 2.36a4.31 4.31 0 0 0-.33 1.66v6.17a4.34 4.34 0 0 0 4.34 4.34h5.82a4.34 4.34 0 0 0 4.34-4.34v-.91l2.53 2c.289.23.634.378 1 .43h.25a2 2 0 0 0 .9-.21a2.13 2.13 0 0 0 .98-.74a2.05 2.05 0 0 0 .32-1.11V9.42a2.13 2.13 0 0 0-.32-1.06m-10.38 4.4h-5a.75.75 0 1 1 0-1.5h5a.75.75 0 1 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoPlayer(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M3.196 7.873h17.608m-4.997 0V2.877M8.193 7.873V2.877m1.947 9.051v4.922c0 .101.032.2.091.286c.06.085.145.154.246.199a.663.663 0 0 0 .633-.057l3.798-2.65a.556.556 0 0 0 .176-.199a.495.495 0 0 0-.02-.492a.568.568 0 0 0-.192-.186l-3.798-2.272a.66.66 0 0 0-.616-.025a.58.58 0 0 0-.232.198a.5.5 0 0 0-.086.276"/><rect width="18.5" height="18.5" x="2.75" y="2.75" rx="6"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoPlayerFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.706 6.259h-4.328V2.212a6.14 6.14 0 0 1 4.328 4.047M6.65 2.212v4.047H2.275A6.13 6.13 0 0 1 6.65 2.212M15.613 2H8.415v4.259h7.198zM22 8.024H2.004a.671.671 0 0 0 0 .152v7.647A6.177 6.177 0 0 0 8.156 22h7.645a6.186 6.186 0 0 0 6.176-6.177V8.176A.67.67 0 0 0 22 8.024m-6.328 7.14a1.39 1.39 0 0 1-.376.448l-3.53 2.47a1.257 1.257 0 0 1-.634.224h-.118a1.27 1.27 0 0 1-1.059-.541c-.13-.195-.2-.425-.2-.66v-4.61c-.001-.23.064-.455.189-.648c.13-.177.299-.322.494-.423c.2-.09.416-.135.635-.13c.211.01.417.07.6.177L15.2 13.6a1.177 1.177 0 0 1 .589.976c.016.204-.025.407-.118.589"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoPlus(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="M9.042 9.5v5m-2.495-2.494h5"/><path stroke-linejoin="round" d="M12 5.32H6.095A3.595 3.595 0 0 0 2.5 8.923v6.162a3.595 3.595 0 0 0 3.595 3.595H12a3.595 3.595 0 0 0 3.595-3.595V8.924A3.594 3.594 0 0 0 12 5.32m9.5 4.118v5.135c0 .25-.071.496-.205.708a1.355 1.355 0 0 1-.555.493a1.27 1.27 0 0 1-.73.124a1.366 1.366 0 0 1-.677-.278l-3.225-2.588a1.377 1.377 0 0 1-.503-1.047c0-.2.045-.396.133-.575c.092-.168.218-.315.37-.432l3.225-2.567a1.36 1.36 0 0 1 .678-.278c.25-.032.504.011.729.124a1.325 1.325 0 0 1 .76 1.181"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoPlusFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.92 8.36a2.06 2.06 0 0 0-.86-.77a2.11 2.11 0 0 0-2.2.24l-2.53 2v-.9a4.142 4.142 0 0 0-.32-1.67a4.33 4.33 0 0 0-4-2.69H6.1a4.31 4.31 0 0 0-1.66.33a4.38 4.38 0 0 0-2.35 2.36a4.31 4.31 0 0 0-.33 1.66v6.17a4.34 4.34 0 0 0 4.34 4.34h5.82a4.34 4.34 0 0 0 4.34-4.34v-.91l2.53 2c.289.23.634.378 1 .43h.25a2 2 0 0 0 .9-.21a2.13 2.13 0 0 0 .98-.74a2.05 2.05 0 0 0 .32-1.11V9.42a2.13 2.13 0 0 0-.32-1.06m-10.38 4.4H9.78v1.75a.75.75 0 1 1-1.5 0v-1.75H6.54a.75.75 0 1 1 0-1.5h1.74V9.51a.75.75 0 1 1 1.5 0v1.75h1.76a.75.75 0 1 1 0 1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoQuestionMark(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M12 5.32H6.095A3.595 3.595 0 0 0 2.5 8.923v6.162a3.595 3.595 0 0 0 3.595 3.595H12a3.595 3.595 0 0 0 3.595-3.595V8.924A3.594 3.594 0 0 0 12 5.32m9.5 4.118v5.135c0 .25-.071.496-.205.708a1.355 1.355 0 0 1-.555.493a1.27 1.27 0 0 1-.73.124a1.366 1.366 0 0 1-.677-.278l-3.225-2.588a1.377 1.377 0 0 1-.503-1.047c0-.2.045-.396.133-.575c.092-.168.218-.315.37-.432l3.225-2.567a1.36 1.36 0 0 1 .678-.278c.25-.032.504.011.729.124a1.325 1.325 0 0 1 .76 1.181"/><path stroke-miterlimit="10" d="M7.543 9.948a1.599 1.599 0 0 1 1.742-.906a1.55 1.55 0 0 1 1.137.81a1.345 1.345 0 0 1-.784 1.851a.994.994 0 0 0-.64.898v.37"/><path stroke-linejoin="round" d="M8.97 15.015h.004"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoQuestionMarkFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.925 8.32a2.11 2.11 0 0 0-.86-.76a2.11 2.11 0 0 0-2.2.24l-2.52 2v-.9a4.37 4.37 0 0 0-2.68-4a4.4 4.4 0 0 0-1.67-.33h-5.9a4.33 4.33 0 0 0-1.67.33a4.38 4.38 0 0 0-2.35 2.36a4.31 4.31 0 0 0-.33 1.66v6.17a4.34 4.34 0 0 0 4.35 4.34h5.83a4.37 4.37 0 0 0 4.017-2.677a4.32 4.32 0 0 0 .333-1.663v-.9l2.52 2c.324.256.718.41 1.13.44h.25a2 2 0 0 0 .9-.21a2.09 2.09 0 0 0 1.18-1.88V9.39a2.19 2.19 0 0 0-.33-1.07m-13 7.42a.75.75 0 1 1 0-1.5a.75.75 0 0 1 0 1.5m2.17-4.52c-.116.274-.29.52-.51.72a1.86 1.86 0 0 1-.75.45a.3.3 0 0 0-.09.08a.33.33 0 0 0 0 .13v.35a.75.75 0 0 1-1.5 0v-.32a1.74 1.74 0 0 1 1.12-1.6a.76.76 0 0 0 .24-.13a.6.6 0 0 0 .14-.21a.52.52 0 0 0 .05-.24a.54.54 0 0 0-.06-.25a.67.67 0 0 0-.23-.25a.91.91 0 0 0-.91-.06a.84.84 0 0 0-.37.4a.76.76 0 0 1-1.24.198a.75.75 0 0 1-.13-.828a2.36 2.36 0 0 1 1.05-1.11a2.39 2.39 0 0 1 1.51-.23c.36.053.703.19 1 .4c.29.212.528.485.7.8c.13.276.201.575.21.88a2.09 2.09 0 0 1-.23.82"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoUpload(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="M9.047 9.5v5"/><path stroke-linejoin="round" d="M11.34 11.605L9.373 9.638a.458.458 0 0 0-.651 0l-1.968 1.967"/><path stroke-linejoin="round" d="M12 5.32H6.095A3.595 3.595 0 0 0 2.5 8.923v6.162a3.595 3.595 0 0 0 3.595 3.595H12a3.595 3.595 0 0 0 3.595-3.595V8.924A3.594 3.594 0 0 0 12 5.32m9.5 4.118v5.135c0 .25-.071.496-.205.708a1.355 1.355 0 0 1-.555.493a1.27 1.27 0 0 1-.73.124a1.366 1.366 0 0 1-.677-.278l-3.225-2.588a1.377 1.377 0 0 1-.503-1.047c0-.2.045-.396.133-.575c.092-.168.218-.315.37-.432l3.225-2.567a1.36 1.36 0 0 1 .678-.278c.25-.032.504.011.729.124a1.325 1.325 0 0 1 .76 1.181"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoUploadFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.93 8.33a2.06 2.06 0 0 0-.86-.77a2.11 2.11 0 0 0-2.2.24l-2.52 2v-.9a4.37 4.37 0 0 0-2.68-4A4.4 4.4 0 0 0 12 4.57H6.1a4.33 4.33 0 0 0-1.67.33a4.38 4.38 0 0 0-2.35 2.36a4.31 4.31 0 0 0-.33 1.66v6.17a4.34 4.34 0 0 0 4.35 4.34h5.83a4.37 4.37 0 0 0 4.017-2.677a4.32 4.32 0 0 0 .333-1.663v-.9l2.52 2c.289.23.634.378 1 .43h.25a2 2 0 0 0 .9-.21a2.13 2.13 0 0 0 .98-.78a2.05 2.05 0 0 0 .32-1.11V9.37a2.12 2.12 0 0 0-.32-1.04m-10.06 3.78a.75.75 0 0 1-1.06 0l-1-1v3.38a.75.75 0 1 1-1.5 0v-3.38l-1 1a.74.74 0 0 1-1.06 0a.75.75 0 0 1 0-1.06l2-2a1.36 1.36 0 0 1 .39-.27c.15-.06.309-.09.47-.09a1.15 1.15 0 0 1 .85.36l2 2a.75.75 0 0 1-.13 1.06z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Visa(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.689 8.87L22 15.129h-1.502c-.06-.262-.107-.536-.167-.798c-.06-.262 0-.131-.155-.131h-1.823a.143.143 0 0 0-.167.13c-.084.24-.167.478-.262.704c-.096.226-.084.107-.12.107h-1.608v-.143c.794-1.883 1.589-3.77 2.383-5.661a.691.691 0 0 1 .62-.453c.524-.024 1.001-.012 1.49-.012m-1.12 1.753l-.847 2.277h1.335zM6.112 13.126v-.143c.477-1.323.965-2.646 1.43-3.969c0-.107.096-.155.215-.155h1.585v.155l-2.467 5.96a.203.203 0 0 1-.227.154H5.266c-.12 0-.167 0-.203-.155c-.44-1.716-.894-3.432-1.335-5.16a.346.346 0 0 0-.155-.215A7.08 7.08 0 0 0 2.12 8.99s-.084-.06-.12-.083a.334.334 0 0 1 .155 0h2.384a.715.715 0 0 1 .822.68l.68 3.491a.299.299 0 0 0 0 .084zm5.554 1.764l.239-1.395l.167.096a2.991 2.991 0 0 0 1.978.274a.62.62 0 0 0 .548-.465a.584.584 0 0 0-.357-.596c-.239-.155-.5-.274-.727-.429a4.067 4.067 0 0 1-.81-.596a1.645 1.645 0 0 1 0-2.11a2.625 2.625 0 0 1 1.346-.786a4.112 4.112 0 0 1 2.384.119h.06l-.227 1.287a6.779 6.779 0 0 0-.918-.214a6.062 6.062 0 0 0-.93 0a.596.596 0 0 0-.297.155a.359.359 0 0 0 0 .548a2.8 2.8 0 0 0 .464.346c.31.19.632.345.93.548a1.715 1.715 0 0 1 .048 2.896a3.134 3.134 0 0 1-1.645.644a4.565 4.565 0 0 1-2.253-.322m-2.646.238l1.025-6.257h1.633l-1.025 6.257z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VisaSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m18.35 10.85l-.7 1.92h1.12c-.15-.66-.29-1.29-.42-1.92m0 0l-.7 1.92h1.12c-.15-.66-.29-1.29-.42-1.92M19 4H5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3M7.51 14.64H6.35c-.1 0-.14 0-.16-.13c-.38-1.44-.75-2.89-1.13-4.33a.24.24 0 0 0-.13-.18a6.09 6.09 0 0 0-1.21-.51l-.11-.08a.65.65 0 0 1 .13 0h2a.61.61 0 0 1 .69.58L7 12.87a.639.639 0 0 0 0 .07l.05-.12l1.2-3.33a.15.15 0 0 1 .17-.12h1.34c0 .05 0 .09-.05.13l-2.07 5a.18.18 0 0 1-.13.14m2 0c.28-1.76.57-3.51.85-5.25h1.37a504.513 504.513 0 0 0-.85 5.25zm5.49-.49c-.4.31-.877.5-1.38.55a4 4 0 0 1-1.92-.27c.07-.39.14-.78.2-1.17l.15.08a2.48 2.48 0 0 0 1.65.22c.22-.05.42-.14.46-.39a.49.49 0 0 0-.3-.5c-.2-.13-.41-.22-.61-.35a3.481 3.481 0 0 1-.67-.51a1.36 1.36 0 0 1 0-1.76c.3-.335.7-.566 1.14-.66a3.26 3.26 0 0 1 2 .1h.06c-.07.37-.13.74-.19 1.08a7.917 7.917 0 0 0-.78-.18a4.082 4.082 0 0 0-.78 0a.5.5 0 0 0-.25.13a.3.3 0 0 0 0 .46c.12.108.25.202.39.28c.26.16.53.3.78.47a1.43 1.43 0 0 1 .05 2.42m4.11.48c0-.23-.09-.45-.13-.67c-.04-.22-.05-.11-.13-.11h-1.52a.13.13 0 0 0-.14.1c-.07.2-.14.4-.22.6a.14.14 0 0 1-.1.08h-1.35l.05-.12l2-4.74a.57.57 0 0 1 .52-.39h1.21c.37 1.75.73 3.5 1.1 5.25zm-.8-3.79l-.7 1.92h1.12c-.13-.65-.27-1.28-.4-1.91z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Voicemail(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.644 15.894a3.894 3.894 0 1 0 0-7.788a3.894 3.894 0 0 0 0 7.788m10.712 0a3.894 3.894 0 1 0 0-7.788a3.894 3.894 0 0 0 0 7.788m-10.712 0h10.712"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VoicemailFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.985 12a4.65 4.65 0 0 1-4.64 4.65H6.655a4.65 4.65 0 1 1 3.41-1.5h3.88a4.64 4.64 0 1 1 8.06-3.15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeDown(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M4.603 14.088V9.912c0-.416.143-.814.396-1.108c.253-.293.596-.458.954-.458h1.89c.237-.001.47-.073.675-.21l3.97-2.63a1.21 1.21 0 0 1 .672-.204c.236.001.467.074.671.212c.204.137.373.334.491.57c.118.237.18.506.182.78v10.273a1.762 1.762 0 0 1-.182.778a1.477 1.477 0 0 1-.49.571c-.205.137-.436.21-.672.212a1.208 1.208 0 0 1-.672-.204l-3.97-2.631a1.233 1.233 0 0 0-.675-.209h-1.89c-.358 0-.701-.165-.954-.459a1.702 1.702 0 0 1-.396-1.107"/><path stroke-miterlimit="10" d="M17.831 15.715a5.344 5.344 0 0 0 0-7.559"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeDownFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.26 6.86v10.28c0 .385-.09.765-.26 1.11a2.21 2.21 0 0 1-.74.86a2 2 0 0 1-1.09.34a1.93 1.93 0 0 1-1.08-.33l-4-2.63a.5.5 0 0 0-.26-.08H5.94a2 2 0 0 1-1.52-.72a2.47 2.47 0 0 1-.58-1.6V9.91a2.45 2.45 0 0 1 .58-1.59a2 2 0 0 1 1.52-.72h1.93a.43.43 0 0 0 .26-.09l4-2.63a2.7 2.7 0 0 1 1.09-.33a2 2 0 0 1 1.09.34c.316.217.572.514.74.86c.153.35.224.729.21 1.11m2.61 9.715a.79.79 0 0 1-.53-.22a.75.75 0 0 1 0-1.06a4.59 4.59 0 0 0 0-6.5a.75.75 0 1 1 1.03-1.06a6.08 6.08 0 0 1 0 8.62a.79.79 0 0 1-.5.22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeMute(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-miterlimit="10" d="m20.951 9.554l-4.902 4.902m0-4.902l4.902 4.902"/><path stroke-linejoin="round" d="M3.049 14.088V9.912c0-.416.142-.814.395-1.108c.253-.293.597-.458.955-.458h1.89c.236-.001.469-.073.675-.21l3.969-2.63a1.21 1.21 0 0 1 .672-.204c.236.001.467.074.671.212c.204.137.374.334.492.57c.118.237.18.506.18.78v10.273c0 .273-.062.542-.18.778a1.476 1.476 0 0 1-.492.571c-.203.137-.435.21-.67.212a1.207 1.207 0 0 1-.673-.204l-3.97-2.631a1.233 1.233 0 0 0-.674-.209h-1.89c-.358 0-.702-.165-.955-.459a1.702 1.702 0 0 1-.395-1.107"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeMuteFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.495 13.93a.75.75 0 0 1-1.06 1.06l-1.92-1.92l-1.93 1.92a.712.712 0 0 1-.53.22a.739.739 0 0 1-.53-.22a.75.75 0 0 1 0-1.06l1.92-1.92l-1.92-1.92a.75.75 0 0 1 0-1.06a.739.739 0 0 1 1.06 0l1.93 1.92l1.92-1.92a.74.74 0 0 1 1.06 0a.75.75 0 0 1 0 1.06l-1.92 1.92zm-7.79-7.07v10.28c.002.385-.087.766-.26 1.11a2.21 2.21 0 0 1-.74.86a2 2 0 0 1-1.09.34a1.9 1.9 0 0 1-1.07-.33l-4-2.63a.559.559 0 0 0-.27-.08h-1.89a2 2 0 0 1-1.52-.72a2.47 2.47 0 0 1-.58-1.6V9.91a2.45 2.45 0 0 1 .58-1.59a2 2 0 0 1 1.52-.72h1.89a.48.48 0 0 0 .27-.09l4-2.63a2 2 0 0 1 1.08-.33a2 2 0 0 1 1.08.34c.318.217.573.514.74.86c.172.345.261.725.26 1.11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeUp(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M3 14.088V9.912c0-.416.142-.814.395-1.108c.254-.293.597-.458.955-.458h1.89c.237-.001.47-.073.675-.21l3.97-2.63c.204-.135.436-.205.672-.204c.236.001.467.074.67.212c.205.137.374.334.492.57c.118.237.18.506.181.78v10.273c0 .273-.063.542-.181.778a1.477 1.477 0 0 1-.491.571c-.204.137-.435.21-.671.212a1.208 1.208 0 0 1-.673-.204l-3.969-2.631a1.233 1.233 0 0 0-.675-.209H4.35c-.358 0-.701-.165-.955-.459A1.702 1.702 0 0 1 3 14.089"/><path stroke-miterlimit="10" d="M16.228 15.715a5.344 5.344 0 0 0 0-7.559m2.267 9.884a8.552 8.552 0 0 0 0-12.093"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeUpFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.65 6.86v10.28c.001.385-.088.765-.26 1.11a2.14 2.14 0 0 1-.74.86a2 2 0 0 1-1.09.34a1.93 1.93 0 0 1-1.08-.33l-4-2.63a.5.5 0 0 0-.26-.08H4.35a2 2 0 0 1-1.52-.72a2.45 2.45 0 0 1-.58-1.6V9.91a2.45 2.45 0 0 1 .58-1.59a2 2 0 0 1 1.52-.72h1.9a.48.48 0 0 0 .27-.09l4-2.63a2 2 0 0 1 1.08-.33a2 2 0 0 1 1.09.34c.318.217.573.514.74.86c.158.348.233.728.22 1.11m2.6 9.68a.79.79 0 0 1-.53-.22a.75.75 0 0 1 0-1.06a4.591 4.591 0 0 0 0-6.5a.75.75 0 1 1 1.06-1.06a6.08 6.08 0 0 1 0 8.62a.79.79 0 0 1-.53.22"/><path fill="currentColor" d="M18.5 18.78a.729.729 0 0 1-.53-.22a.738.738 0 0 1 0-1.06a7.81 7.81 0 0 0 0-11a.74.74 0 0 1 0-1.06a.75.75 0 0 1 1.06 0a9.31 9.31 0 0 1 0 13.15a.74.74 0 0 1-.53.19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeZero(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.05 14.088V9.912c0-.416.142-.814.395-1.108c.253-.293.597-.458.955-.458h1.89c.237-.001.47-.073.675-.21l3.97-2.63c.204-.135.436-.205.672-.204c.236.001.467.074.67.212c.205.137.374.334.492.57c.118.237.18.506.181.78v10.273c0 .273-.063.542-.181.778a1.477 1.477 0 0 1-.491.571c-.204.137-.435.21-.671.212a1.208 1.208 0 0 1-.673-.204l-3.969-2.631a1.233 1.233 0 0 0-.675-.209H8.4c-.358 0-.702-.165-.955-.459a1.702 1.702 0 0 1-.395-1.107"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeZeroFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.679 6.845v10.28a2.5 2.5 0 0 1-.26 1.11a2.15 2.15 0 0 1-.75.86a1.92 1.92 0 0 1-1.08.34a1.93 1.93 0 0 1-1.08-.33l-4-2.63a.5.5 0 0 0-.26-.08h-1.83a2 2 0 0 1-1.52-.72a2.42 2.42 0 0 1-.58-1.6v-4.18a2.4 2.4 0 0 1 .58-1.59a2 2 0 0 1 1.52-.69h1.89a.43.43 0 0 0 .26-.09l4-2.63a2.08 2.08 0 0 1 1.09-.33a2 2 0 0 1 1.09.34c.317.217.572.514.74.86c.142.342.207.71.19 1.08"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WaterDrop(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20 13.277c0-4.525-4.59-8.481-6.81-10.136a2.004 2.004 0 0 0-2.38 0C8.59 4.796 4 8.752 4 13.277c0 5.98 5 7.973 8 7.973s8-1.993 8-7.973"/><path d="M7 13.277c0 1.322.527 2.59 1.464 3.524A5.009 5.009 0 0 0 12 18.26"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WaterDropFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.631 2.523a2.822 2.822 0 0 0-3.273 0c-2.662 2.002-7.116 6.005-7.116 10.74C3.242 19.667 8.477 22 12 22c3.523 0 8.758-2.332 8.758-8.738c0-4.724-4.464-8.757-7.127-10.74M12 19.268a6.007 6.007 0 0 1-4.254-1.762a5.926 5.926 0 0 1-1.751-4.234a1 1 0 1 1 2.001 0a3.923 3.923 0 0 0 1.171 2.823A4.003 4.003 0 0 0 12 17.265a1 1 0 0 1 0 2.003"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WaterGlass(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19.206 3H4.804a1.036 1.036 0 0 0-.788.354a1.062 1.062 0 0 0-.258.832L5.4 18.194c.09.777.46 1.492 1.04 2.01a3.119 3.119 0 0 0 2.1.796h6.97a3.119 3.119 0 0 0 2.1-.796c.58-.518.95-1.233 1.04-2.01l1.591-14.008a1.07 1.07 0 0 0-.254-.828A1.048 1.048 0 0 0 19.206 3"/><path d="M4.375 9.353c2.837-1.059 5.85-2.001 8.206-.138c2.355 1.864 4.616 1.546 6.96 1.06M7.819 14.573l.24 1.705c.067.418.274.8.587 1.081a1.81 1.81 0 0 0 1.13.464l1.047.053"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WaterGlassFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.935 3.521a1.79 1.79 0 0 0-.98-1.11a1.7 1.7 0 0 0-.74-.16H4.825a1.49 1.49 0 0 0-.74.15a1.76 1.76 0 0 0-.62.46a1.79 1.79 0 0 0-.37.66a1.75 1.75 0 0 0-.07.75l1.64 14c.11.958.57 1.84 1.29 2.48a3.85 3.85 0 0 0 2.57 1h7a3.87 3.87 0 0 0 2.58-1a4 4 0 0 0 1.28-2.48l.89-7.87v-.12l.69-6a2 2 0 0 0-.03-.76m-10.1 15.35h-1a2.88 2.88 0 0 1-1.75-.72a2.82 2.82 0 0 1-.9-1.67l-.24-1.72a1.01 1.01 0 0 1 2-.28l.24 1.7a.79.79 0 0 0 .26.48a.8.8 0 0 0 .51.21h1a1 1 0 1 1 0 2zm8-9.23c-2.11.41-3.91.49-5.82-1c-2.37-1.87-5.24-1.27-8-.29l-.5-4.24a.29.29 0 0 1 0-.13a.26.26 0 0 1 .07-.11a.2.2 0 0 1 .1-.08a.22.22 0 0 1 .12 0h14.52a.39.39 0 0 1 .1.08c.03.03.05.069.06.11a.22.22 0 0 1 0 .14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WeChat(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.278 8.12a7.21 7.21 0 0 0-6.991-4.454C5.27 3.666 2 6.467 2 9.914a5.875 5.875 0 0 0 2.571 4.759L3.707 16.4a.427.427 0 0 0 .59.548l2.189-1.259a8.37 8.37 0 0 0 1.849.416c.853 0 .219-1.094.317-1.608c0-3.217 3.184-5.832 7.09-5.832a.438.438 0 0 0 .536-.547m-9.727.524a1.247 1.247 0 0 1-1.247-1.247A1.247 1.247 0 0 1 7.8 7.42a1.247 1.247 0 0 1-1.25 1.224m5 0a1.247 1.247 0 1 1 .033 0z"/><path fill="currentColor" d="M22 14.498c0-2.757-2.8-5-6.247-5s-6.258 2.188-6.258 5s2.811 5 6.258 5a7.527 7.527 0 0 0 1.914-.24l2.068 1.039a.416.416 0 0 0 .57-.537l-.559-1.411A4.703 4.703 0 0 0 22 14.498m-8.337-.832a.831.831 0 1 1 0-1.663a.843.843 0 0 1 .842.832a.842.842 0 0 1-.842.831m4.168 0a.832.832 0 1 1 0-1.663a.832.832 0 0 1 0 1.663"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Whatsapp(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.794 2.642c4.68.9 7.86 4.954 7.7 9.773c-.147 4.444-3.742 8.318-8.27 8.873a9.547 9.547 0 0 1-5.318-.86a1.374 1.374 0 0 0-.84-.082c-1.356.325-2.701.694-4.05 1.045c-.148.038-.3.064-.516.109c.405-1.482.783-2.904 1.187-4.32c.103-.358.074-.644-.091-.995c-1.575-3.345-1.438-6.627.708-9.66c2.151-3.042 5.187-4.34 8.92-3.96c.176.017.351.042.57.077m5.914 11.004c.268-1.096.3-2.207.038-3.298c-.788-3.288-2.834-5.419-6.16-6.093c-3.263-.66-6.003.44-7.936 3.138c-1.936 2.702-1.978 5.6-.38 8.503c.205.373.26.68.135 1.078c-.228.728-.405 1.472-.628 2.298c.925-.24 1.736-.47 2.558-.65c.233-.051.538-.024.742.088c4.696 2.585 10.277.2 11.63-5.064"/><path fill="currentColor" d="M9.745 8.158c.179.427.3.84.51 1.203c.3.518.209.953-.205 1.318c-.445.392-.379.725-.06 1.175c.735 1.036 1.658 1.813 2.823 2.322c.32.14.563.164.77-.157c.086-.132.206-.24.306-.364c.583-.726.4-.72 1.324-.319c.291.127.582.262.851.428c.27.165.68.335.732.57c.118.52-.048 1.048-.482 1.434c-.8.712-1.72.83-2.725.552c-2.174-.6-3.846-1.904-5.127-3.71c-.452-.637-.856-1.344-1.11-2.078c-.308-.894-.09-1.756.546-2.513c.375-.445.83-.545 1.325-.426c.199.048.338.344.522.565"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WhatsappFilled(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none"><g clip-path="url(#mageWhatsappFilled0)"><path fill="currentColor" d="m13.79 2.64l-.57-.08a9.13 9.13 0 0 0-8.92 4a9.1 9.1 0 0 0-.71 9.66a1.3 1.3 0 0 1 .1 1c-.41 1.41-.79 2.83-1.19 4.32l.5-.15c1.35-.36 2.7-.72 4.05-1.05a1.45 1.45 0 0 1 .85.08a9.45 9.45 0 1 0 5.89-17.78m2.52 13.12a2.76 2.76 0 0 1-2.72.56a9.19 9.19 0 0 1-5.13-3.71a8.51 8.51 0 0 1-1.11-2.08a2.49 2.49 0 0 1 .55-2.52a1.23 1.23 0 0 1 1.32-.42c.2.05.34.34.52.56c.146.413.317.817.51 1.21a.94.94 0 0 1-.2 1.31c-.45.4-.38.73-.06 1.18a6.71 6.71 0 0 0 2.82 2.32c.32.14.56.17.77-.16c.09-.13.21-.24.31-.36c.58-.73.4-.72 1.32-.32c.293.123.577.267.85.43c.27.16.68.33.74.57a1.45 1.45 0 0 1-.49 1.43"/></g><defs><clipPath id="mageWhatsappFilled0"><path fill="#fff" d="M2.5 2.5h19v19h-19z"/></clipPath></defs></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wifi(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13.308 17.886a1.308 1.308 0 1 1-2.616 0a1.308 1.308 0 0 1 2.616 0m-5.011-3.702a5.234 5.234 0 0 1 7.406 0M5.524 11.41a9.16 9.16 0 0 1 12.952 0"/><path d="M2.75 8.636a13.083 13.083 0 0 1 18.5 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wrench(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.252 14.407a6.823 6.823 0 0 1-6.532 1.776a2.016 2.016 0 0 0-1.95.487l-4.28 4.294a.974.974 0 0 1-1.385 0l-2.067-2.069a.976.976 0 0 1 0-1.385l4.28-4.284a1.952 1.952 0 0 0 .488-1.951a6.835 6.835 0 0 1 1.912-6.65a6.823 6.823 0 0 1 6.736-1.566l-2.925 2.927c-.75.752-1.112 2.342-.341 3.103l1.715 1.727c.76.761 2.35.41 3.11-.351l2.925-2.927a6.836 6.836 0 0 1-1.686 6.869"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WrenchFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.763 11.382a7.573 7.573 0 0 1-3.47 4.693a7.56 7.56 0 0 1-5.772.827a1.27 1.27 0 0 0-.67 0a1.178 1.178 0 0 0-.57.31l-4.266 4.29a1.88 1.88 0 0 1-.56.37a1.683 1.683 0 0 1-.669.13a1.648 1.648 0 0 1-.659-.13a1.759 1.759 0 0 1-.56-.37L2.5 19.432a1.61 1.61 0 0 1-.37-.56a1.771 1.771 0 0 1 0-1.33a1.61 1.61 0 0 1 .37-.56l4.277-4.28a1.17 1.17 0 0 0 .32-.56c.06-.209.06-.43 0-.64a7.586 7.586 0 0 1 2.117-7.42a7.522 7.522 0 0 1 3.497-1.88a7.428 7.428 0 0 1 3.997.15a.74.74 0 0 1 .31 1.24L14.1 6.522a2.131 2.131 0 0 0-.56 1.41a.91.91 0 0 0 .21.63l1.719 1.73a1.1 1.1 0 0 0 .91.18a2.128 2.128 0 0 0 1.138-.53l2.918-2.93a.78.78 0 0 1 .71-.2a.749.749 0 0 1 .539.51a7.606 7.606 0 0 1 .08 4.06"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func X(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.751 3h3.067l-6.7 7.625L22 21h-6.172l-4.833-6.293L5.464 21h-3.07l7.167-8.155L2 3h6.328l4.37 5.752zm-1.076 16.172h1.7L7.404 4.732H5.58z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Xsquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8.78 6.91l7.68 10.11h-1.19L7.51 6.91z"/><path fill="currentColor" d="M17 2H7a5 5 0 0 0-5 5v10a5 5 0 0 0 5 5h10a5 5 0 0 0 5-5V7a5 5 0 0 0-5-5m-2.32 16.3l-3.38-4.4l-3.88 4.4H5.28l5-5.71L5 5.7h4.43l3.06 4l3.51-4h2.14L13.48 11L19 18.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Youtube(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 11.939a26.45 26.45 0 0 0-.41-4.801a2.5 2.5 0 0 0-1.803-1.714a60.96 60.96 0 0 0-7.81-.41c-2.609-.03-5.217.11-7.808.42a2.522 2.522 0 0 0-1.76 1.76a26.382 26.382 0 0 0-.408 4.8c-.01 1.61.127 3.216.409 4.8a2.555 2.555 0 0 0 1.78 1.782c2.592.303 5.2.44 7.81.409a59.37 59.37 0 0 0 7.787-.41a2.521 2.521 0 0 0 1.759-1.758c.307-1.608.46-3.242.454-4.878m-7.411.582l-4.005 2.271a.53.53 0 0 1-.785-.464V9.616a.52.52 0 0 1 .785-.454l4.137 2.378a.52.52 0 0 1-.022.908z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Zap(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m13.058 2.577l-7.941 10.74a.612.612 0 0 0-.068.597a.543.543 0 0 0 .188.235a.468.468 0 0 0 .273.084h5.986l-.907 7.023a.243.243 0 0 0 .027.135a.21.21 0 0 0 .094.09c.04.02.083.024.125.013a.2.2 0 0 0 .107-.07l7.941-10.741a.612.612 0 0 0 .068-.597a.544.544 0 0 0-.188-.234a.468.468 0 0 0-.273-.085h-5.986l.907-7.023a.243.243 0 0 0-.027-.135a.21.21 0 0 0-.094-.09a.183.183 0 0 0-.125-.013a.199.199 0 0 0-.107.07"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZapCircle(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8.196 20.731a9.511 9.511 0 0 1 0-17.424m7.609 17.424a9.511 9.511 0 0 0 0-17.424"/><path d="m12.859 4.333l-6.476 8.76a.57.57 0 0 0-.098.228a.514.514 0 0 0 0 .253a.465.465 0 0 0 .155.196a.408.408 0 0 0 .22.065h4.894l-.742 5.71a.171.171 0 0 0 0 .105a.163.163 0 0 0 .073.074a.139.139 0 0 0 .098 0a.128.128 0 0 0 .09-.057l6.475-8.76a.562.562 0 0 0 .098-.236a.496.496 0 0 0-.195-.44a.408.408 0 0 0-.22-.066h-4.894l.742-5.71a.171.171 0 0 0 0-.105a.163.163 0 0 0-.074-.074a.139.139 0 0 0-.097 0a.13.13 0 0 0-.05.057"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZapCircleFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.279 21.234c-.1 0-.2-.02-.293-.058a10.023 10.023 0 0 1 0-18.35a.734.734 0 1 1 .586 1.347a8.55 8.55 0 0 0 0 15.617a.741.741 0 0 1 .38.976a.722.722 0 0 1-.673.468m7.428-.029a.722.722 0 0 1-.293-1.396a8.55 8.55 0 0 0 0-15.616a.735.735 0 0 1 .585-1.347a10.014 10.014 0 0 1 0 18.35a.78.78 0 0 1-.293.01"/><path fill="currentColor" d="M18.234 10.764c-.026.216-.106.42-.234.596l-6.354 8.54c-.065.081-.14.153-.224.215a.899.899 0 0 1-.352.126h-.146a.974.974 0 0 1-.312 0a.976.976 0 0 1-.46-.429a.365.365 0 0 1 0-.107a.81.81 0 0 1 0-.46l.616-4.704H6.864a1.143 1.143 0 0 1-1.045-.702a.478.478 0 0 1 0-.108a1.269 1.269 0 0 1 0-.605c.038-.17.112-.329.215-.468l6.325-8.56a.79.79 0 0 1 .38-.303a.976.976 0 0 1 .615 0a.976.976 0 0 1 .459.42a.588.588 0 0 1 0 .117a.78.78 0 0 1 0 .45l-.605 4.713h3.904c.214 0 .424.061.605.176c.193.129.343.312.43.527c.075.179.105.373.087.566"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZapFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.745 10.5a1.41 1.41 0 0 1-.26.66l-7.94 10.73a.94.94 0 0 1-.53.35a.827.827 0 0 1-.22 0a1.099 1.099 0 0 1-.4-.08a1 1 0 0 1-.55-1l.8-6.21h-5.13a1.41 1.41 0 0 1-.7-.22a1.33 1.33 0 0 1-.46-.56a1.45 1.45 0 0 1-.1-.69c.03-.236.12-.46.26-.65l7.94-10.71a.93.93 0 0 1 .51-.34a1 1 0 0 1 .63.06a.94.94 0 0 1 .44.42a1 1 0 0 1 .11.55l-.8 6.21h5.14a1.16 1.16 0 0 1 .7.22c.194.141.35.33.45.55c.096.223.134.467.11.71"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZapSquare(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21.25 7.889V5.833a3.083 3.083 0 0 0-3.083-3.083h-3.084m0 18.5h3.084a3.083 3.083 0 0 0 3.083-3.083V16.11m-18.5.001v2.056a3.083 3.083 0 0 0 3.083 3.083h3.084m0-18.5H5.833A3.083 3.083 0 0 0 2.75 5.833V7.89m10.109-3.557l-6.476 8.76a.57.57 0 0 0-.098.228a.514.514 0 0 0 0 .253a.465.465 0 0 0 .155.196a.408.408 0 0 0 .22.065h4.894l-.742 5.71a.171.171 0 0 0 0 .105a.163.163 0 0 0 .073.074a.139.139 0 0 0 .098 0a.128.128 0 0 0 .09-.057l6.475-8.76a.562.562 0 0 0 .098-.236a.496.496 0 0 0-.195-.44a.408.408 0 0 0-.22-.066h-4.894l.742-5.71a.171.171 0 0 0 0-.105a.163.163 0 0 0-.074-.074a.139.139 0 0 0-.097 0a.13.13 0 0 0-.05.057"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZapSquareFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.25 8.632a.738.738 0 0 1-.75-.752V5.875a2.369 2.369 0 0 0-.691-1.653a2.357 2.357 0 0 0-1.649-.693h-3.08a.75.75 0 0 1-.75-.752a.753.753 0 0 1 .75-.752h3.08a3.822 3.822 0 0 1 2.71 1.133A3.78 3.78 0 0 1 22 5.875V7.88a.753.753 0 0 1-.75.752m-3.09 13.393h-3.08a.75.75 0 0 1-.75-.752a.753.753 0 0 1 .75-.752h3.08c.619 0 1.213-.245 1.651-.683a2.349 2.349 0 0 0 .689-1.653V16.11a.753.753 0 0 1 .75-.752a.749.749 0 0 1 .75.752v2.065a3.807 3.807 0 0 1-1.13 2.717a3.855 3.855 0 0 1-2.71 1.133m-9.25 0H5.83a3.855 3.855 0 0 1-2.71-1.123A3.79 3.79 0 0 1 2 18.185V16.11a.753.753 0 0 1 .75-.752a.749.749 0 0 1 .75.752v2.065a2.33 2.33 0 0 0 .68 1.654c.44.436 1.032.681 1.65.682h3.08a.75.75 0 0 1 .75.752a.753.753 0 0 1-.75.752zM2.75 8.632A.738.738 0 0 1 2 7.88V5.875a3.778 3.778 0 0 1 1.12-2.717a3.83 3.83 0 0 1 2.71-1.133h3.08a.75.75 0 0 1 .75.752a.753.753 0 0 1-.75.752H5.83a2.337 2.337 0 0 0-1.648.69A2.349 2.349 0 0 0 3.5 5.875V7.88a.753.753 0 0 1-.75.752m15.64 2.105a1.32 1.32 0 0 1-.24.611l-6.51 8.772a1.205 1.205 0 0 1-.23.22a.92.92 0 0 1-.36.131h-.15a.999.999 0 0 1-.32 0a1.002 1.002 0 0 1-.47-.44a.378.378 0 0 1 0-.111a.836.836 0 0 1 0-.471l.63-4.833h-4a1.167 1.167 0 0 1-1.07-.721a.493.493 0 0 1 0-.11a1.306 1.306 0 0 1 0-.622c.039-.174.114-.338.22-.481l6.48-8.792a.81.81 0 0 1 .39-.311a.998.998 0 0 1 .63 0a1 1 0 0 1 .47.431c.004.04.004.08 0 .12a.804.804 0 0 1 0 .461l-.62 4.843h4c.22 0 .435.062.62.18c.198.132.351.32.44.541c.077.184.108.384.09.582"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomIn(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.244 17.738a7.494 7.494 0 1 0 0-14.988a7.494 7.494 0 0 0 0 14.988m5.318-2.176l5.688 5.688M10.244 6.245v7.998m3.999-3.999H6.245"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomInFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.857 20.437l-5.23-5.22a8.27 8.27 0 1 0-1.41 1.41l5.22 5.23a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42m-7.72-9.29h-3v3a1 1 0 1 1-2 0v-3h-3a1 1 0 1 1 0-2h3v-3a1 1 0 0 1 2 0v3h3a1 1 0 1 1 0 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomOut(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.244 17.738a7.494 7.494 0 1 0 0-14.988a7.494 7.494 0 0 0 0 14.988m5.318-2.176l5.688 5.688m-7.007-11.006H6.245"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomOutFill(children ...ElementRenderer) *MageIcon {
	return &MageIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.785 20.35l-5.22-5.22a8.18 8.18 0 1 0-1.41 1.42l5.22 5.22a1 1 0 1 0 1.41-1.42m-15.71-9.29a1 1 0 1 1 0-2h8a1 1 0 0 1 0 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
