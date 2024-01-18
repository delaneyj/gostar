package si_glyph

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = "0.0.2"
	hAttr          = "1em"
	viewbox        = "0 0 17 0"
	fill           = "currentColor"
)

type SiGlyphIcon struct {
	*SVGSVGElement
}

type SiGlyphIconFn func(children ...ElementRenderer) *SiGlyphIcon

var IconLookup = map[string]SiGlyphIconFn{
	"abacus":                   Abacus,
	"adjustmentHorizon":        AdjustmentHorizon,
	"adjustmentVertical":       AdjustmentVertical,
	"airBalloon":               AirBalloon,
	"airplane":                 Airplane,
	"airplaneTwo":              AirplaneTwo,
	"alarmClock":               AlarmClock,
	"alien":                    Alien,
	"alighLeft":                AlighLeft,
	"alignCenter":              AlignCenter,
	"alignLeft":                AlignLeft,
	"alignRight":               AlignRight,
	"ambulance":                Ambulance,
	"anchor":                   Anchor,
	"android":                  Android,
	"angleOne":                 AngleOne,
	"angleTwo":                 AngleTwo,
	"antennaOne":               AntennaOne,
	"apron":                    Apron,
	"arrowBackward":            ArrowBackward,
	"arrowChange":              ArrowChange,
	"arrowCircleRycycle":       ArrowCircleRycycle,
	"arrowDown":                ArrowDown,
	"arrowForward":             ArrowForward,
	"arrowFourWay":             ArrowFourWay,
	"arrowFullscreen":          ArrowFullscreen,
	"arrowFullscreenTwo":       ArrowFullscreenTwo,
	"arrowLeft":                ArrowLeft,
	"arrowLeftRight":           ArrowLeftRight,
	"arrowMove":                ArrowMove,
	"arrowReload":              ArrowReload,
	"arrowResizeFive":          ArrowResizeFive,
	"arrowResizeFour":          ArrowResizeFour,
	"arrowResizeOne":           ArrowResizeOne,
	"arrowResizeSix":           ArrowResizeSix,
	"arrowResizeThree":         ArrowResizeThree,
	"arrowResizeTwo":           ArrowResizeTwo,
	"arrowRight":               ArrowRight,
	"arrowShuffle":             ArrowShuffle,
	"arrowThickDown":           ArrowThickDown,
	"arrowThickLeft":           ArrowThickLeft,
	"arrowThickRight":          ArrowThickRight,
	"arrowThickThinDown":       ArrowThickThinDown,
	"arrowThickThinUp":         ArrowThickThinUp,
	"arrowThickUp":             ArrowThickUp,
	"arrowThinDown":            ArrowThinDown,
	"arrowThinLeft":            ArrowThinLeft,
	"arrowThinLeftBottom":      ArrowThinLeftBottom,
	"arrowThinLeftTop":         ArrowThinLeftTop,
	"arrowThinRight":           ArrowThinRight,
	"arrowThinRightBottom":     ArrowThinRightBottom,
	"arrowThinRightTop":        ArrowThinRightTop,
	"arrowThinUp":              ArrowThinUp,
	"arrowThreeWayOne":         ArrowThreeWayOne,
	"arrowThreeWayTwo":         ArrowThreeWayTwo,
	"arrowTriangleRecycle":     ArrowTriangleRecycle,
	"arrowTwoLeftRight":        ArrowTwoLeftRight,
	"arrowTwoUp":               ArrowTwoUp,
	"arrowTwoWay":              ArrowTwoWay,
	"arrowTwoWayLeftRight":     ArrowTwoWayLeftRight,
	"arrowTwoWayRight":         ArrowTwoWayRight,
	"arrowTwoWayRightBottom":   ArrowTwoWayRightBottom,
	"arrowUp":                  ArrowUp,
	"arrowUpDown":              ArrowUpDown,
	"arrowWave":                ArrowWave,
	"artBoard":                 ArtBoard,
	"askterisk":                Askterisk,
	"atmCard":                  AtmCard,
	"axe":                      Axe,
	"baby":                     Baby,
	"babyMilkBotl":             BabyMilkBotl,
	"babyStroller":             BabyStroller,
	"babyTroller":              BabyTroller,
	"backPack":                 BackPack,
	"backwardPage":             BackwardPage,
	"badgeName":                BadgeName,
	"bag":                      Bag,
	"bagPlus":                  BagPlus,
	"bagRemove":                BagRemove,
	"balance":                  Balance,
	"balloon":                  Balloon,
	"bandage":                  Bandage,
	"bank":                     Bank,
	"barcode":                  Barcode,
	"barrier":                  Barrier,
	"baseball":                 Baseball,
	"baseballStick":            BaseballStick,
	"basket":                   Basket,
	"basketArrowDown":          BasketArrowDown,
	"basketArrowLeft":          BasketArrowLeft,
	"basketArrowRight":         BasketArrowRight,
	"basketArrowUp":            BasketArrowUp,
	"basketChecked":            BasketChecked,
	"basketError":              BasketError,
	"basketPlus":               BasketPlus,
	"basketRemove":             BasketRemove,
	"basketball":               Basketball,
	"batteryCharging":          BatteryCharging,
	"batteryEmpty":             BatteryEmpty,
	"batteryFull":              BatteryFull,
	"batteryHalf":              BatteryHalf,
	"batteryHalfTwo":           BatteryHalfTwo,
	"batteryLow":               BatteryLow,
	"bed":                      Bed,
	"bell":                     Bell,
	"bicycleMan":               BicycleMan,
	"bicycleOne":               BicycleOne,
	"bicycleThree":             BicycleThree,
	"bikini":                   Bikini,
	"billiardBall":             BilliardBall,
	"binocular":                Binocular,
	"birthdayCake":             BirthdayCake,
	"biscuit":                  Biscuit,
	"blender":                  Blender,
	"bloodBag":                 BloodBag,
	"bluetooth":                Bluetooth,
	"board":                    Board,
	"boat":                     Boat,
	"bolt":                     Bolt,
	"bombOne":                  BombOne,
	"bombTwo":                  BombTwo,
	"bone":                     Bone,
	"book":                     Book,
	"bookFour":                 BookFour,
	"bookOne":                  BookOne,
	"bookOpen":                 BookOpen,
	"bookPerson":               BookPerson,
	"bookThree":                BookThree,
	"bookcase":                 Bookcase,
	"bookmark":                 Bookmark,
	"botlMilk":                 BotlMilk,
	"botlTwo":                  BotlTwo,
	"bowTie":                   BowTie,
	"box":                      Box,
	"boxDownload":              BoxDownload,
	"boxUpload":                BoxUpload,
	"bread":                    Bread,
	"briefcase":                Briefcase,
	"briefcasePerson":          BriefcasePerson,
	"brightness":               Brightness,
	"brushAndPencil":           BrushAndPencil,
	"brushOne":                 BrushOne,
	"brushTwo":                 BrushTwo,
	"bubbleChat":               BubbleChat,
	"bubbleMessage":            BubbleMessage,
	"bubbleMessageDot":         BubbleMessageDot,
	"bubbleMessageDotTwo":      BubbleMessageDotTwo,
	"bubbleMessageHi":          BubbleMessageHi,
	"bubbleMessageQuote":       BubbleMessageQuote,
	"bubbleMessageTalk":        BubbleMessageTalk,
	"bubbleMessageText":        BubbleMessageText,
	"bucket":                   Bucket,
	"bug":                      Bug,
	"building":                 Building,
	"bulletCheckedList":        BulletCheckedList,
	"bulletList":               BulletList,
	"bulletListTwo":            BulletListTwo,
	"bus":                      Bus,
	"buttonArrowDown":          ButtonArrowDown,
	"buttonArrowLeft":          ButtonArrowLeft,
	"buttonArrowRight":         ButtonArrowRight,
	"buttonArrowUp":            ButtonArrowUp,
	"buttonBuy":                ButtonBuy,
	"buttonError":              ButtonError,
	"buttonHd":                 ButtonHd,
	"buttonPlay":               ButtonPlay,
	"buttonPlus":               ButtonPlus,
	"buttonRemove":             ButtonRemove,
	"buttonSale":               ButtonSale,
	"buttonSell":               ButtonSell,
	"buttonStarburst":          ButtonStarburst,
	"buttonTriangleUp":         ButtonTriangleUp,
	"buttonTv":                 ButtonTv,
	"cabinCable":               CabinCable,
	"cabinet":                  Cabinet,
	"calculator":               Calculator,
	"calculatorTwo":            CalculatorTwo,
	"calendarEmpty":            CalendarEmpty,
	"calendarOne":              CalendarOne,
	"calendarThree":            CalendarThree,
	"call":                     Call,
	"callForward":              CallForward,
	"callReply":                CallReply,
	"camera":                   Camera,
	"cameraProjector":          CameraProjector,
	"cameraSecurity":           CameraSecurity,
	"canWater":                 CanWater,
	"candle":                   Candle,
	"candy":                    Candy,
	"candyStick":               CandyStick,
	"car":                      Car,
	"carGarage":                CarGarage,
	"casette":                  Casette,
	"cashierMachine":           CashierMachine,
	"castle":                   Castle,
	"caterpillarMachine":       CaterpillarMachine,
	"centeJustify":             CenteJustify,
	"chairOne":                 ChairOne,
	"chairTwo":                 ChairTwo,
	"championCup":              ChampionCup,
	"chartColumn":              ChartColumn,
	"chartColumnDecrease":      ChartColumnDecrease,
	"chartColumnIncrease":      ChartColumnIncrease,
	"chartDecrease":            ChartDecrease,
	"chartPiece":               ChartPiece,
	"checked":                  Checked,
	"cheese":                   Cheese,
	"cherry":                   Cherry,
	"christmasMistletoe":       ChristmasMistletoe,
	"christmassBall":           ChristmassBall,
	"christmassEgg":            ChristmassEgg,
	"christmassHat":            ChristmassHat,
	"christmassTree":           ChristmassTree,
	"circle":                   Circle,
	"circleBackward":           CircleBackward,
	"circleControlPad":         CircleControlPad,
	"circleDrashed":            CircleDrashed,
	"circleError":              CircleError,
	"circleForward":            CircleForward,
	"circleHelp":               CircleHelp,
	"circleInfo":               CircleInfo,
	"circleLoadLeft":           CircleLoadLeft,
	"circleLoadRight":          CircleLoadRight,
	"circlePlus":               CirclePlus,
	"circleRemove":             CircleRemove,
	"circleStar":               CircleStar,
	"circleTriangleDown":       CircleTriangleDown,
	"circleTriangleLeft":       CircleTriangleLeft,
	"circleTriangleRight":      CircleTriangleRight,
	"city":                     City,
	"clamp":                    Clamp,
	"clapboard":                Clapboard,
	"clapboardPlay":            ClapboardPlay,
	"clip":                     Clip,
	"clipboard":                Clipboard,
	"clipboardChecked":         ClipboardChecked,
	"clock":                    Clock,
	"close":                    Close,
	"cloud":                    Cloud,
	"cloudCloud":               CloudCloud,
	"cloudDownload":            CloudDownload,
	"cloudHeavyRain":           CloudHeavyRain,
	"cloudPlus":                CloudPlus,
	"cloudRainHeavyRain":       CloudRainHeavyRain,
	"cloudRemove":              CloudRemove,
	"cloudSnow":                CloudSnow,
	"cloudSun":                 CloudSun,
	"cloudThunder":             CloudThunder,
	"cloudUpload":              CloudUpload,
	"clover":                   Clover,
	"cockPot":                  CockPot,
	"cocktail":                 Cocktail,
	"coconut":                  Coconut,
	"code":                     Code,
	"coffeeMachine":            CoffeeMachine,
	"colorPicker":              ColorPicker,
	"columnDecrease":           ColumnDecrease,
	"columnIncrease":           ColumnIncrease,
	"columnWave":               ColumnWave,
	"comb":                     Comb,
	"compass":                  Compass,
	"cone":                     Cone,
	"congratulationHat":        CongratulationHat,
	"connectOne":               ConnectOne,
	"connectTwo":               ConnectTwo,
	"contactBook":              ContactBook,
	"contacter":                Contacter,
	"contrast":                 Contrast,
	"controlPad":               ControlPad,
	"corkscrew":                Corkscrew,
	"cornerFlag":               CornerFlag,
	"coverFlow":                CoverFlow,
	"coverFood":                CoverFood,
	"cow":                      Cow,
	"cpu":                      Cpu,
	"cran":                     Cran,
	"crop":                     Crop,
	"crossHair":                CrossHair,
	"crossHairTwo":             CrossHairTwo,
	"crown":                    Crown,
	"cruise":                   Cruise,
	"cubic":                    Cubic,
	"cupCake":                  CupCake,
	"curtain":                  Curtain,
	"customerSupport":          CustomerSupport,
	"dashboard":                Dashboard,
	"dataArrowDown":            DataArrowDown,
	"database":                 Database,
	"databaseDownload":         DatabaseDownload,
	"databaseError":            DatabaseError,
	"databasePlus":             DatabasePlus,
	"databaseRemove":           DatabaseRemove,
	"databaseShare":            DatabaseShare,
	"databaseUpload":           DatabaseUpload,
	"delete":                   Delete,
	"deliciousCircle":          DeliciousCircle,
	"deny":                     Deny,
	"desktop":                  Desktop,
	"dialNumber":               DialNumber,
	"dialNumberOne":            DialNumberOne,
	"diamond":                  Diamond,
	"diceFive":                 DiceFive,
	"diceOne":                  DiceOne,
	"diceSix":                  DiceSix,
	"diceThree":                DiceThree,
	"diceTwo":                  DiceTwo,
	"disc":                     Disc,
	"discAdd":                  DiscAdd,
	"discDeny":                 DiscDeny,
	"discDownload":             DiscDownload,
	"discError":                DiscError,
	"discPause":                DiscPause,
	"discPlay":                 DiscPlay,
	"discPlayTwo":              DiscPlayTwo,
	"discRemove":               DiscRemove,
	"discStop":                 DiscStop,
	"discUpload":               DiscUpload,
	"dish":                     Dish,
	"diskAntenna":              DiskAntenna,
	"dna":                      Dna,
	"document":                 Document,
	"documentArrowDown":        DocumentArrowDown,
	"documentArrowLeft":        DocumentArrowLeft,
	"documentArrowRight":       DocumentArrowRight,
	"documentArrowUp":          DocumentArrowUp,
	"documentBackward":         DocumentBackward,
	"documentBulletList":       DocumentBulletList,
	"documentChecked":          DocumentChecked,
	"documentCopy":             DocumentCopy,
	"documentDoc":              DocumentDoc,
	"documentEdit":             DocumentEdit,
	"documentError":            DocumentError,
	"documentForward":          DocumentForward,
	"documentHelp":             DocumentHelp,
	"documentMusic":            DocumentMusic,
	"documentPdf":              DocumentPdf,
	"documentPlus":             DocumentPlus,
	"documentRemove":           DocumentRemove,
	"documentRss":              DocumentRss,
	"documentSearch":           DocumentSearch,
	"documentStar":             DocumentStar,
	"documentWarning":          DocumentWarning,
	"dog":                      Dog,
	"door":                     Door,
	"downstair":                Downstair,
	"downwardsArrowToBar":      DownwardsArrowToBar,
	"drill":                    Drill,
	"dropWater":                DropWater,
	"dropbox":                  Dropbox,
	"drum":                     Drum,
	"easal":                    Easal,
	"edit":                     Edit,
	"egg":                      Egg,
	"eject":                    Eject,
	"electron":                 Electron,
	"elevatorDown":             ElevatorDown,
	"elevatorUp":               ElevatorUp,
	"emoticon":                 Emoticon,
	"endPage":                  EndPage,
	"erase":                    Erase,
	"euro":                     Euro,
	"excavator":                Excavator,
	"extinguisher":             Extinguisher,
	"eyeDrop":                  EyeDrop,
	"eyeGlass":                 EyeGlass,
	"factory":                  Factory,
	"faucet":                   Faucet,
	"feather":                  Feather,
	"female":                   Female,
	"fence":                    Fence,
	"fenceTwo":                 FenceTwo,
	"fileBox":                  FileBox,
	"fileDownload":             FileDownload,
	"fileUpload":               FileUpload,
	"film":                     Film,
	"filmThirtyFiveMm":         FilmThirtyFiveMm,
	"finder":                   Finder,
	"fingerUp":                 FingerUp,
	"fire":                     Fire,
	"fireAlarm":                FireAlarm,
	"fireHydrant":              FireHydrant,
	"fireWood":                 FireWood,
	"firstAidBriefcase":        FirstAidBriefcase,
	"fish":                     Fish,
	"flag":                     Flag,
	"float":                    Float,
	"floppyDisk":               FloppyDisk,
	"flower":                   Flower,
	"flowerPot":                FlowerPot,
	"folder":                   Folder,
	"folderContact":            FolderContact,
	"folderError":              FolderError,
	"folderMusic":              FolderMusic,
	"folderOpen":               FolderOpen,
	"folderPlus":               FolderPlus,
	"folderRemove":             FolderRemove,
	"folderRss":                FolderRss,
	"folderSearch":             FolderSearch,
	"folderShare":              FolderShare,
	"folderWarning":            FolderWarning,
	"footSign":                 FootSign,
	"forcus":                   Forcus,
	"forklift":                 Forklift,
	"forwardPage":              ForwardPage,
	"fridge":                   Fridge,
	"fullscreen":               Fullscreen,
	"gameControll":             GameControll,
	"gameOne":                  GameOne,
	"gasStation":               GasStation,
	"gate":                     Gate,
	"gear":                     Gear,
	"gearOne":                  GearOne,
	"ghost":                    Ghost,
	"gift":                     Gift,
	"glassWater":               GlassWater,
	"global":                   Global,
	"globe":                    Globe,
	"glove":                    Glove,
	"golfBall":                 GolfBall,
	"golfFlag":                 GolfFlag,
	"guitar":                   Guitar,
	"hamburger":                Hamburger,
	"hammer":                   Hammer,
	"hammerAndWrench":          HammerAndWrench,
	"hand":                     Hand,
	"handLamp":                 HandLamp,
	"handSwitch":               HandSwitch,
	"handcuff":                 Handcuff,
	"hanger":                   Hanger,
	"hardware":                 Hardware,
	"hat":                      Hat,
	"hatChef":                  HatChef,
	"head":                     Head,
	"headSet":                  HeadSet,
	"heart":                    Heart,
	"heartDelete":              HeartDelete,
	"heartPlus":                HeartPlus,
	"heartRemove":              HeartRemove,
	"heartSignal":              HeartSignal,
	"helicopter":               Helicopter,
	"helicopterPad":            HelicopterPad,
	"helmWheel":                HelmWheel,
	"helmet":                   Helmet,
	"hightheel":                Hightheel,
	"history":                  History,
	"hockey":                   Hockey,
	"homePage":                 HomePage,
	"horse":                    Horse,
	"horseShoe":                HorseShoe,
	"hospital":                 Hospital,
	"house":                    House,
	"iceCream":                 IceCream,
	"idBadge":                  IdBadge,
	"image":                    Image,
	"inColumns":                InColumns,
	"inboxDownload":            InboxDownload,
	"inboxUpload":              InboxUpload,
	"infinity":                 Infinity,
	"infinityTwo":              InfinityTwo,
	"info":                     Info,
	"ipod":                     Ipod,
	"iron":                     Iron,
	"jumpBackward":             JumpBackward,
	"jumpDoublePageLeftRight":  JumpDoublePageLeftRight,
	"jumpForward":              JumpForward,
	"jumpPageLeftRight":        JumpPageLeftRight,
	"jumpPageUpDown":           JumpPageUpDown,
	"kette":                    Kette,
	"key":                      Key,
	"keyTwo":                   KeyTwo,
	"keyboard":                 Keyboard,
	"knife":                    Knife,
	"ladderPool":               LadderPool,
	"lamp":                     Lamp,
	"lampDesk":                 LampDesk,
	"lavabo":                   Lavabo,
	"lawHammer":                LawHammer,
	"layoutFour":               LayoutFour,
	"layoutOne":                LayoutOne,
	"layoutThree":              LayoutThree,
	"layoutTwo":                LayoutTwo,
	"leaf":                     Leaf,
	"leftJustify":              LeftJustify,
	"leftwardsArrowToBar":      LeftwardsArrowToBar,
	"lightAlarm":               LightAlarm,
	"lightBulb":                LightBulb,
	"lightHouse":               LightHouse,
	"like":                     Like,
	"likeUnlike":               LikeUnlike,
	"lineTwoAnglePoint":        LineTwoAnglePoint,
	"linkOne":                  LinkOne,
	"linkThree":                LinkThree,
	"linkTwo":                  LinkTwo,
	"load":                     Load,
	"lock":                     Lock,
	"lockUnlock":               LockUnlock,
	"louderSpeaker":            LouderSpeaker,
	"louderTwo":                LouderTwo,
	"magnet":                   Magnet,
	"magnifier":                Magnifier,
	"magnifierTwo":             MagnifierTwo,
	"mail":                     Mail,
	"mailBox":                  MailBox,
	"mailEmpty":                MailEmpty,
	"mailHasMail":              MailHasMail,
	"mailInbox":                MailInbox,
	"mailSend":                 MailSend,
	"male":                     Male,
	"manDoctor":                ManDoctor,
	"mapSquare":                MapSquare,
	"mapThree":                 MapThree,
	"markSnorker":              MarkSnorker,
	"maskOne":                  MaskOne,
	"maskTwo":                  MaskTwo,
	"medalStar":                MedalStar,
	"microphoneOne":            MicrophoneOne,
	"microphoneTwo":            MicrophoneTwo,
	"microscope":               Microscope,
	"mobile":                   Mobile,
	"mocule":                   Mocule,
	"money":                    Money,
	"moneyBag":                 MoneyBag,
	"moneyCoin":                MoneyCoin,
	"moneyThree":               MoneyThree,
	"moonStar":                 MoonStar,
	"motobike":                 Motobike,
	"mountain":                 Mountain,
	"moviePlay":                MoviePlay,
	"movieProjector":           MovieProjector,
	"multifunctionKnife":       MultifunctionKnife,
	"mushrooms":                Mushrooms,
	"music":                    Music,
	"musicNote":                MusicNote,
	"mustache":                 Mustache,
	"network":                  Network,
	"networkTwo":               NetworkTwo,
	"newspaper":                Newspaper,
	"noDog":                    NoDog,
	"noSmoke":                  NoSmoke,
	"note":                     Note,
	"noteTwo":                  NoteTwo,
	"oldPhone":                 OldPhone,
	"open":                     Open,
	"pallette":                 Pallette,
	"paperClip":                PaperClip,
	"paperPlane":               PaperPlane,
	"paperRoll":                PaperRoll,
	"paperShredder":            PaperShredder,
	"pause":                    Pause,
	"pawn":                     Pawn,
	"paypal":                   Paypal,
	"pen":                      Pen,
	"penNib":                   PenNib,
	"pencil":                   Pencil,
	"percent":                  Percent,
	"person":                   Person,
	"personChecked":            PersonChecked,
	"personDoorMan":            PersonDoorMan,
	"personError":              PersonError,
	"personMan":                PersonMan,
	"personPeople":             PersonPeople,
	"personPlus":               PersonPlus,
	"personPrison":             PersonPrison,
	"personPublic":             PersonPublic,
	"personRemove":             PersonRemove,
	"personTalk":               PersonTalk,
	"personTwo":                PersonTwo,
	"personWoman":              PersonWoman,
	"petCarrier":               PetCarrier,
	"phoneFax":                 PhoneFax,
	"phoneNumber":              PhoneNumber,
	"piano":                    Piano,
	"pick":                     Pick,
	"picture":                  Picture,
	"pictureCopy":              PictureCopy,
	"pictureTwo":               PictureTwo,
	"piggyBank":                PiggyBank,
	"pill":                     Pill,
	"pinLocation":              PinLocation,
	"pinLocationAdd":           PinLocationAdd,
	"pinLocationDelete":        PinLocationDelete,
	"pinLocationLove":          PinLocationLove,
	"pinLocationMap":           PinLocationMap,
	"pinLocationOne":           PinLocationOne,
	"pinLocationRemove":        PinLocationRemove,
	"pinLocationTwo":           PinLocationTwo,
	"pinMap":                   PinMap,
	"pingPongRacket":           PingPongRacket,
	"pipe":                     Pipe,
	"pizza":                    Pizza,
	"plugin":                   Plugin,
	"plus":                     Plus,
	"podium":                   Podium,
	"pokerFour":                PokerFour,
	"pokerOne":                 PokerOne,
	"pokerThree":               PokerThree,
	"pokerTwo":                 PokerTwo,
	"print":                    Print,
	"puzzle":                   Puzzle,
	"quoteClose":               QuoteClose,
	"quoteOpen":                QuoteOpen,
	"radio":                    Radio,
	"radioactive":              Radioactive,
	"record":                   Record,
	"reduce":                   Reduce,
	"reelAudio":                ReelAudio,
	"reelFilm":                 ReelFilm,
	"remove":                   Remove,
	"repeat":                   Repeat,
	"resizeInFrame":            ResizeInFrame,
	"resizeOutFrame":           ResizeOutFrame,
	"retroeexcavadora":         Retroeexcavadora,
	"ribbon":                   Ribbon,
	"rightJustify":             RightJustify,
	"rightwardsArrowToBar":     RightwardsArrowToBar,
	"ring":                     Ring,
	"road":                     Road,
	"rocket":                   Rocket,
	"roller":                   Roller,
	"rss":                      Rss,
	"rugbyBall":                RugbyBall,
	"ruler":                    Ruler,
	"safeBox":                  SafeBox,
	"safePin":                  SafePin,
	"satellite":                Satellite,
	"saw":                      Saw,
	"scissor":                  Scissor,
	"scissorLineCut":           ScissorLineCut,
	"screenFul":                ScreenFul,
	"screenScale":              ScreenScale,
	"screw":                    Screw,
	"screwDriver":              ScrewDriver,
	"sdCard":                   SdCard,
	"seriff":                   Seriff,
	"sewingMachine":            SewingMachine,
	"sewingRoll":               SewingRoll,
	"shareFive":                ShareFive,
	"shareOne":                 ShareOne,
	"shareThree":               ShareThree,
	"shareTwo":                 ShareTwo,
	"sheep":                    Sheep,
	"shield":                   Shield,
	"shieldPlus":               ShieldPlus,
	"shieldStar":               ShieldStar,
	"shieldTwo":                ShieldTwo,
	"shovel":                   Shovel,
	"shower":                   Shower,
	"signBoard":                SignBoard,
	"signFoot":                 SignFoot,
	"signIn":                   SignIn,
	"signOut":                  SignOut,
	"signP":                    SignP,
	"signRoadOne":              SignRoadOne,
	"signRoadTwo":              SignRoadTwo,
	"signalOne":                SignalOne,
	"signalThree":              SignalThree,
	"signalTwo":                SignalTwo,
	"siteMap":                  SiteMap,
	"siteMapRevert":            SiteMapRevert,
	"skateboard":               Skateboard,
	"skull":                    Skull,
	"sleep":                    Sleep,
	"slideShow":                SlideShow,
	"smartphone":               Smartphone,
	"sms":                      Sms,
	"snow":                     Snow,
	"soccerYard":               SoccerYard,
	"sock":                     Sock,
	"socket":                   Socket,
	"solarBlind":               SolarBlind,
	"sos":                      Sos,
	"sound":                    Sound,
	"soundMute":                SoundMute,
	"spaceShip":                SpaceShip,
	"spanner":                  Spanner,
	"spoonFork":                SpoonFork,
	"spray":                    Spray,
	"square":                   Square,
	"squareChecked":            SquareChecked,
	"squareDashedOne":          SquareDashedOne,
	"squareDashedTwo":          SquareDashedTwo,
	"squareDelicious":          SquareDelicious,
	"squareEightAnglePoint":    SquareEightAnglePoint,
	"squareFour":               SquareFour,
	"squareFourAnglePoint":     SquareFourAnglePoint,
	"squarePlus":               SquarePlus,
	"squareStar":               SquareStar,
	"stamps":                   Stamps,
	"starCross":                StarCross,
	"starStick":                StarStick,
	"stele":                    Stele,
	"stelescope":               Stelescope,
	"stereoJack":               StereoJack,
	"stethoscope":              Stethoscope,
	"store":                    Store,
	"stove":                    Stove,
	"strawberry":               Strawberry,
	"streetTwo":                StreetTwo,
	"strolley":                 Strolley,
	"strolleyArrowDown":        StrolleyArrowDown,
	"strolleyArrowUp":          StrolleyArrowUp,
	"strolleyPlus":             StrolleyPlus,
	"strolleyRemove":           StrolleyRemove,
	"subway":                   Subway,
	"suitcase":                 Suitcase,
	"suitcasePerson":           SuitcasePerson,
	"syringe":                  Syringe,
	"tshirt":                   Tshirt,
	"tablet":                   Tablet,
	"tag":                      Tag,
	"tagOne":                   TagOne,
	"tagOnePlus":               TagOnePlus,
	"tagPrice":                 TagPrice,
	"tank":                     Tank,
	"targer":                   Targer,
	"teaCup":                   TeaCup,
	"teeth":                    Teeth,
	"telephoneBox":             TelephoneBox,
	"telescope":                Telescope,
	"television":               Television,
	"tennisRacketBall":         TennisRacketBall,
	"tentCamp":                 TentCamp,
	"tentOne":                  TentOne,
	"testTube":                 TestTube,
	"testTubeEmpty":            TestTubeEmpty,
	"testTubeTwo":              TestTubeTwo,
	"textSearch":               TextSearch,
	"thermometer":              Thermometer,
	"threeBall":                ThreeBall,
	"ticTacToe":                TicTacToe,
	"timeGlass":                TimeGlass,
	"timeReload":               TimeReload,
	"timer":                    Timer,
	"toilet":                   Toilet,
	"toolBox":                  ToolBox,
	"trafficLight":             TrafficLight,
	"train":                    Train,
	"trainRail":                TrainRail,
	"trash":                    Trash,
	"tree":                     Tree,
	"trees":                    Trees,
	"triangleDoubleArrowDown":  TriangleDoubleArrowDown,
	"triangleDoubleArrowLeft":  TriangleDoubleArrowLeft,
	"triangleDoubleArrowRight": TriangleDoubleArrowRight,
	"triangleDoubleArrowUp":    TriangleDoubleArrowUp,
	"triangleDown":             TriangleDown,
	"triangleLeft":             TriangleLeft,
	"triangleRight":            TriangleRight,
	"triangleUp":               TriangleUp,
	"trolleyArrowDown":         TrolleyArrowDown,
	"trolleyArrowUp":           TrolleyArrowUp,
	"trolleyBriefcase":         TrolleyBriefcase,
	"trolleyError":             TrolleyError,
	"trolleyFull":              TrolleyFull,
	"trolleyPlus":              TrolleyPlus,
	"trolleyRemove":            TrolleyRemove,
	"trolleyTwo":               TrolleyTwo,
	"truck":                    Truck,
	"trumpet":                  Trumpet,
	"turnOff":                  TurnOff,
	"twoArrowDown":             TwoArrowDown,
	"twoArrowInDownUp":         TwoArrowInDownUp,
	"twoArrowInLeftRight":      TwoArrowInLeftRight,
	"twoArrowLeft":             TwoArrowLeft,
	"twoArrowRight":            TwoArrowRight,
	"typewriter":               Typewriter,
	"umberllaChair":            UmberllaChair,
	"umbrellaClose":            UmbrellaClose,
	"umbrellaOpen":             UmbrellaOpen,
	"umbrellaSea":              UmbrellaSea,
	"upstair":                  Upstair,
	"upwardsArrowToBar":        UpwardsArrowToBar,
	"upwardsArrowWithLoop":     UpwardsArrowWithLoop,
	"usb":                      Usb,
	"view":                     View,
	"wacomTablet":              WacomTablet,
	"wall":                     Wall,
	"wallet":                   Wallet,
	"washMachine":              WashMachine,
	"washMachineTwo":           WashMachineTwo,
	"watch":                    Watch,
	"webcam":                   Webcam,
	"weightDown":               WeightDown,
	"weightKilograms":          WeightKilograms,
	"weightUp":                 WeightUp,
	"wheelChair":               WheelChair,
	"wheelSteel":               WheelSteel,
	"wieght":                   Wieght,
	"wifiOne":                  WifiOne,
	"windTurbines":             WindTurbines,
	"window":                   Window,
	"woodStove":                WoodStove,
	"wrench":                   Wrench,
	"wrenchScrewdriver":        WrenchScrewdriver,
	"yen":                      Yen,
	"yingYang":                 YingYang,
	"zip":                      Zip,
	"zoomIn":                   ZoomIn,
	"zoomOut":                  ZoomOut,
}

func Abacus(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" transform="translate(0 1)"><path d="M14.194 13.958H1.765c-1.449 0-1.729-1.15-1.729-2.564V2.607C.036 1.193.315.043 1.765.043h12.429c1.448 0 1.728 1.15 1.728 2.564v8.787c0 1.414-.279 2.564-1.728 2.564M1.923 1A.937.937 0 0 0 1 1.948v10.104c0 .522.414.948.923.948h12.154a.937.937 0 0 0 .923-.948V1.948A.937.937 0 0 0 14.077 1z"/><path d="M3 0h.948v13.068H3zm5 0h.948v13.068H8zm4 0h.948v13.068H12z"/><ellipse cx="3.438" cy="4.976" rx="1.438" ry=".976"/><ellipse cx="3.438" cy="7.976" rx="1.438" ry=".976"/><ellipse cx="3.438" cy="10.976" rx="1.438" ry=".976"/><ellipse cx="8.477" cy="2.976" rx="1.477" ry=".976"/><ellipse cx="8.477" cy="5.976" rx="1.477" ry=".976"/><ellipse cx="8.477" cy="10.976" rx="1.477" ry=".976"/><ellipse cx="12.435" cy="7.977" rx="1.435" ry=".977"/><ellipse cx="12.435" cy="10.977" rx="1.435" ry=".977"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdjustmentHorizon(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 1v1.956h7.928V1zM0 1h1.975v1.975H0zm0 6v1.988h5.011V7zm11 0v1.979h4.987V7zM0 13h8.019v1.962H0zm14 0h2v1.961h-2zm-3.032 3.04c1.087 0 2.008-.822 2.008-1.88c0-1.06-.921-2.16-2.008-2.16s-1.924.961-1.924 2.02c0 1.059.837 2.02 1.924 2.02M7.906 9.993c1.053 0 2.103-1.017 2.103-2.076C10.009 6.858 8.959 6 7.906 6A1.912 1.912 0 0 0 6 7.917c0 1.059.853 2.076 1.906 2.076m-2.85-6.012c1.059 0 1.966-1.005 1.966-2.064C7.022.858 5.976 0 4.917 0A1.917 1.917 0 0 0 3 1.917c0 1.059.997 2.064 2.056 2.064"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AdjustmentVertical(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 0v5.007h1.95V0zm0 11v4.958h2.01V11zm6 3v1.976h1.966V14zM8 0v8h2V0zm6 9v6.942h1.977V9zm0-9v2.933h2.009V0zm.917 8.049c1.059 0 2.094-.994 2.094-2.081S15.976 4 14.917 4S13 4.881 13 5.968s.858 2.081 1.917 2.081m-6 4.961c1.059 0 2.04-1.051 2.04-2.104C10.958 9.853 9.977 9 8.918 9A1.912 1.912 0 0 0 7 10.906c0 1.053.858 2.103 1.917 2.103zm-6-3.021c1.059 0 2.057-1.013 2.057-2.072C4.974 6.858 3.976 6 2.917 6A1.917 1.917 0 0 0 1 7.917c0 1.059.858 2.072 1.917 2.072"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AirBalloon(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.959 15.312a.678.678 0 0 1-.678.678h-2.6a.678.678 0 0 1-.676-.678v-2.234h3.953v2.234zm3.026-10.561c0-1.769-.89-3.707-3.009-4.446c1.238 1.027 3.428 3.403-1.43 10.664L9.974.155s-.463-.123-.99-.123c-.512 0-.941.113-.941.113l.5 10.813C3.668 3.582 5.84 1.295 7.092.294c-2.138.733-3.076 2.683-3.076 4.457C4.016 7.466 8.561 12 7.786 12h2.506c-.71 0 3.693-4.48 3.693-7.249"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Airplane(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="m9.264 10.32l3.52 5.48a.798.798 0 0 0 .51.199a.744.744 0 0 0 .531-.207c.389-.581-.561-4.316-1.861-8.172z"/><path d="M15.613.42c-.584-.586-1.328-.525-1.828-.027l-4.08 4.064c-2.437-.776-6.846-2.275-8.51-2.275c-.232 0-.334.026-.369.037a.758.758 0 0 0-.014 1.012l6.107 4l-3.032 3.02s-2.339-.482-2.681-.527c-.477-.062-1.027.356-.002.879c1.197.609 2.599 1.317 2.599 1.317s.863 1.679 1.39 2.604c.671 1.119 1.065.576.985 0l-.348-2.515l2.884-2.992l2.829-2.819l4.026-4.011c.501-.498.626-1.183.044-1.767"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AirplaneTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" transform="translate(0 1)"><ellipse cx="7.98" cy="7.904" rx="1.98" ry="1.904"/><path d="M3.754 2.492c.797.888 1.325 1.735 1.182 1.895c-.143.158-.904-.431-1.701-1.317c-.795-.887-1.325-1.735-1.182-1.895c.142-.159.904.43 1.701 1.317m9.023.577c-.797.887-1.559 1.477-1.701 1.317c-.143-.16.386-1.007 1.183-1.895c.797-.887 1.558-1.477 1.7-1.317c.142.16-.387 1.009-1.182 1.895m-6.78 11.372c0 .267-.22.481-.49.481a.485.485 0 0 1-.491-.481v-.924c0-.266.219-.479.491-.479c.271 0 .49.214.49.479zm4.98.055a.477.477 0 0 1-.465.489a.477.477 0 0 1-.465-.489v-.938c0-.271.209-.488.465-.488c.257 0 .465.218.465.488z"/><path d="M15.957.496c0 .818-.57 1.482-1.275 1.482H1.304C.6 1.978.03 1.314.03.496c0 0-.195-.419 1.274-.419h13.378c1.373 0 1.275.419 1.275.419M8.035 5.072c1.611 0 2.326.843 2.326.843h2.59c-.6-1.661-2.597-2.882-4.973-2.882c-2.377 0-4.373 1.221-4.973 2.882h2.699c.001 0 .714-.843 2.331-.843m.01 5.893c-1.836 0-2.338-.934-2.338-.934H3c.607 1.71 2.609 2.968 4.989 2.968s4.384-1.258 4.99-2.968h-2.597c.001 0-.503.934-2.337.934M6 8.935s-4.955.331-4.955-.809C1.045 6.981 6 7.049 6 7.049zm4.121-1.901s4.816-.203 4.816.947c0 1.151-4.816.947-4.816.947z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlarmClock(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M14.654.659c-1.312-.951-3.05-.764-3.99.383a10.161 10.161 0 0 1 4.922 3.58c.764-1.275.379-3.008-.932-3.963M3.644 2.427a10.137 10.137 0 0 1 2.73-1.395C5.433-.104 3.694-.288 2.383.654C1.072 1.6.686 3.316 1.451 4.579a10.185 10.185 0 0 1 2.193-2.152m11.28 6.206c0-.198-.012-.393-.028-.586c-.272-3.14-2.698-5.646-5.765-5.966c-.136-.016-.273-.017-.412-.022c-.082-.004-.164-.014-.248-.014h-.004c-3.568 0-6.459 2.949-6.459 6.588c0 .029.004.059.005.088a6.606 6.606 0 0 0 2.028 4.694a1.286 1.286 0 0 0-.693 1.143c0 .709.564 1.284 1.26 1.284c.666 0 1.204-.527 1.25-1.194c.689.311 1.44.507 2.231.554c.126.008.251.02.378.02c.145 0 .285-.013.43-.022a6.321 6.321 0 0 0 2.256-.586c.029.684.576 1.229 1.253 1.229c.696 0 1.259-.575 1.259-1.284c0-.523-.305-.97-.743-1.171a6.617 6.617 0 0 0 2.002-4.755m-6.397 4.681c-2.563 0-4.647-2.116-4.647-4.716s2.084-4.714 4.647-4.714c2.562 0 4.647 2.114 4.647 4.714c0 2.6-2.085 4.716-4.647 4.716"/><path d="M8 5h1v4H8z"/><path d="M8 8h3v1H8z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Alien(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.014.143c-3.855 0-6.983 2.979-6.983 6.651C2.031 10.469 7.209 16 9.014 16C10.822 16 16 10.469 16 6.794C16 3.122 12.873.143 9.014.143M7.895 10.895c-.316.318-1.414-.271-2.448-1.321C4.411 8.528 3.829 7.42 4.145 7.1c.315-.321 1.412.269 2.447 1.317c1.033 1.048 1.619 2.155 1.303 2.478m2.219-.008c-.32-.32.27-1.426 1.321-2.473c1.049-1.047 2.158-1.636 2.48-1.314c.32.318-.271 1.424-1.32 2.47c-1.05 1.047-2.161 1.634-2.481 1.317"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlighLeft(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.043 1.938c0 .518.42.938.938.938h14.082a.938.938 0 0 0 0-1.876H1.981a.938.938 0 0 0-.938.938m0 12c0 .518.42.938.938.938h14.082a.938.938 0 0 0 0-1.876H1.981a.938.938 0 0 0-.938.938m0-6c0 .518.42.938.938.938h10.082a.938.938 0 0 0 0-1.876H1.981a.938.938 0 0 0-.938.938m0 3c0 .518.42.938.938.938h8.082a.938.938 0 0 0 0-1.876H1.981a.938.938 0 0 0-.938.938m0-6c0 .518.42.938.938.938h6.082a.938.938 0 0 0 0-1.876H1.981a.938.938 0 0 0-.938.938"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignCenter(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M17 1.938c0 .518-.42.938-.938.938H1.98A.938.938 0 0 1 1.98 1h14.082c.518 0 .938.42.938.938m0 12c0 .518-.42.938-.938.938H1.98a.938.938 0 0 1 0-1.876h14.082c.518 0 .938.42.938.938m0-6c0 .518-.42.938-.938.938H1.98A.938.938 0 0 1 1.98 7h14.082c.518 0 .938.42.938.938m-3 3c0 .518-.42.938-.938.938H4.98a.938.938 0 0 1 0-1.876h8.082c.518 0 .938.42.938.938m0-6c0 .518-.42.938-.938.938H4.98A.938.938 0 0 1 4.98 4h8.082c.518 0 .938.42.938.938"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignLeft(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.043 1.938c0 .518.42.938.938.938h14.082a.938.938 0 0 0 0-1.876H1.981a.938.938 0 0 0-.938.938m0 12c0 .518.42.938.938.938h14.082a.938.938 0 0 0 0-1.876H1.981a.938.938 0 0 0-.938.938m0-6c0 .518.42.938.938.938h10.082a.938.938 0 0 0 0-1.876H1.981a.938.938 0 0 0-.938.938m0 3c0 .518.42.938.938.938h8.082a.938.938 0 0 0 0-1.876H1.981a.938.938 0 0 0-.938.938m0-6c0 .518.42.938.938.938h6.082a.938.938 0 0 0 0-1.876H1.981a.938.938 0 0 0-.938.938"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignRight(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M17 1.938c0 .518-.42.938-.938.938H1.98A.938.938 0 0 1 1.98 1h14.082c.518 0 .938.42.938.938m0 12c0 .518-.42.938-.938.938H1.98a.938.938 0 0 1 0-1.876h14.082c.518 0 .938.42.938.938m0-6c0 .518-.42.938-.938.938H5.98A.938.938 0 0 1 5.98 7h10.082c.518 0 .938.42.938.938m0 3c0 .518-.42.938-.938.938H7.98a.938.938 0 0 1 0-1.876h8.082c.518 0 .938.42.938.938m0-6c0 .518-.42.938-.938.938H9.98A.938.938 0 0 1 9.98 4h6.082c.518 0 .938.42.938.938"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ambulance(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><g transform="translate(2 2)"><circle cx="1.433" cy="9.433" r="1.433"/><ellipse cx="9.451" cy="9.474" rx="1.451" ry="1.474"/><path d="M8.953 2.047c0-1.122.1-2.031-.944-2.031c-1.045 0-.993.909-.993 2.031z"/></g><path d="M15.338 7.735S12.729 4 11.994 4H1.392C.659 4 .064 4.599.064 5.336v5.315c0 .736.328 1.336.328 1.336h.553a2.656 2.656 0 0 1-.043-.431a2.6 2.6 0 1 1 5.2 0c0 .147-.02.29-.043.431h2.887a2.656 2.656 0 0 1-.043-.431a2.6 2.6 0 1 1 5.2 0c0 .147-.02.29-.043.431h.594c.734 0 1.329-.6 1.329-1.336v-1.33c-.002-.834-.645-1.586-.645-1.586m-7.291-.688H7.043v1.005H5.958V7.047H4.953V5.963h1.005V4.958h1.085v1.005h1.004zm1.871.995V4.904h1.367c.636 0 1.375 1.031 1.375 1.031l1.464 2.106z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Anchor(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m15.469 6.09l-2.426 1.447l1.331.348c-.009 2.278-1.944 4.213-4.416 4.625V5.917h1.984v-.834H9.958V3.637c.568-.352.96-.967.96-1.693c0-1.109-.883-1.92-1.971-1.92S7.01.871 7.01 1.981c0 .751.396 1.355.989 1.69v1.412H6.015v.834H8v6.595c-2.472-.409-4.306-2.352-4.31-4.639l1.3-.334L2.611 6.08L1.054 8.55l1.337-.344c.171 2.76 2.388 5.024 5.304 5.563c.857.248.934.845.934.845l.402 1.327l.413-1.318s.171-.648.933-.854c2.906-.539 5.118-2.797 5.296-5.548l1.299.338zM8.953 1.001a.96.96 0 0 1 .957.964a.96.96 0 0 1-.957.963a.96.96 0 0 1-.958-.963a.96.96 0 0 1 .958-.964"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Android(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.018 6.062s-.008 4.412-.008 5.562c0 1.034.117 1.256 1.021 1.35v2.333c0 .356.437.646.974.646c.538 0 .976-.29.976-.646v-2.312l2.051.004v2.308c0 .356.438.646.975.646c.537 0 .975-.29.975-.646V12.98c.926-.1.978-.371.978-1.335l.001-5.582zm9.977-1.032c-.537 0-.975.29-.975.647l.005 3.634c0 .356.438.646.975.646c.537 0 .976-.29.976-.646l-.005-3.634c0-.357-.44-.647-.976-.647m-12.01 0c-.54 0-.98.295-.98.661v3.642c0 .365.44.661.98.661c.543 0 .983-.296.983-.661V5.691c.001-.366-.44-.661-.983-.661m7.989-2.843c.011-.012.026-.018.035-.032l.929-1.634c.102-.159.041-.365-.134-.457a.384.384 0 0 0-.499.123L9.377 1.82c-.018.028-.025.059-.033.089a4.17 4.17 0 0 0-1.357-.231a4.2 4.2 0 0 0-1.372.234c-.009-.031-.015-.062-.034-.092L5.709.168a.383.383 0 0 0-.498-.123c-.176.092-.236.298-.135.457l.872 1.653c.01.016.028.023.042.037c-1.057.579-1.759 1.598-1.947 2.803h7.948c-.19-1.209-.958-2.229-2.017-2.808M7.042 4.042H5.938V2.98h1.104zm3-.021H8.98V2.938h1.062z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleOne(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" transform="translate(2 1)"><ellipse cx="13.479" cy="13.458" rx="1.479" ry="1.458"/><ellipse cx="13.479" cy="9.458" rx="1.479" ry="1.458"/><ellipse cx="9.479" cy="13.458" rx="1.479" ry="1.458"/><ellipse cx="5.479" cy="13.458" rx="1.479" ry="1.458"/><ellipse cx="9.479" cy="9.458" rx="1.479" ry="1.458"/><ellipse cx="13.479" cy="5.458" rx="1.479" ry="1.458"/><ellipse cx="13.479" cy="1.458" rx="1.479" ry="1.458"/><ellipse cx="9.479" cy="5.458" rx="1.479" ry="1.458"/><ellipse cx="5.479" cy="9.458" rx="1.479" ry="1.458"/><ellipse cx="1.479" cy="13.458" rx="1.479" ry="1.458"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M.648 15.938a.617.617 0 0 1-.436-1.052L14.946.18a.617.617 0 0 1 .871.872L1.085 15.757a.615.615 0 0 1-.437.181m5.004.072a.615.615 0 0 1-.441-.185a.623.623 0 0 1 0-.882l9.723-9.723a.62.62 0 0 1 .881 0a.62.62 0 0 1 0 .883l-9.723 9.722a.615.615 0 0 1-.44.185m4.981-.025a.606.606 0 0 1-.426-1.031l4.752-4.752a.602.602 0 0 1 .854 0a.606.606 0 0 1 0 .853l-4.753 4.752a.6.6 0 0 1-.427.178"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AntennaOne(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.987 0H3.112c-2.127 0-.266 1.982-.266 1.982l5.234 6.08v7.854h1.875V8.077l5.232-6.064c.001 0 1.955-2.013-.2-2.013M7.64 5.582L4.204 1.52S3.526.97 4.35.97c.825 0 3.668.014 3.668.014v4.531c0 .673-.378.067-.378.067m2.302-.066V.985s2.909-.014 3.752-.014c.844 0 .15.55.15.55L10.33 5.583c.001-.001-.388.605-.388-.067"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Apron(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.75 6.318c-1.443-.513-2.355 1.072-3.107 1.651c0 0-.888.224-.888.05c0-3.217 2.703-4.856.015-4.856v-.004l-.108-.438C11.208 1.06 10.457.046 9.485.046c-.962 0-1.754.996-2.209 2.631c0 0-.092.354-.111.441c-2.519 0 .021 1.759.021 4.936c0 .164-.89-.085-.89-.085c-.752-.579-1.598-2.208-3.041-1.695c-.178.062-.262.238-.188.39c.074.152.28.224.458.16c.906-.32 1.354.859 2.298 1.587c0 0 .79.614 1.291.781c-.351 3.169-1.957 6.78-1.957 6.78l8.695.024s-1.648-3.646-2.02-6.808c.498-.168 1.283-.778 1.283-.778c.944-.728 1.458-1.863 2.364-1.543c.178.063.384-.008.458-.16c.075-.15-.009-.326-.187-.389M9.48 5.055c-.792 0-1.377-.888-1.502-1.569c.334-1.713.979-2.579 1.577-2.579h-.133c.599 0 1.243.866 1.577 2.579c-.125.682-.71 1.569-1.502 1.569z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowBackward(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m1.307 5.988l5.309-4.645c.411-.41.891-.479 1.302-.068v3.132l.229-.001c5.016 0 8.738 3.563 8.738 8.41c0 1.688-.774 1.073-1.097.484c-1.522-2.78-4.197-4.677-7.681-4.677l-.19.001v3.065c-.411.41-.941.361-1.302.068L1.306 7.474a1.052 1.052 0 0 1 .001-1.486"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowChange(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.539 8.001c.828 0 1.379-.551 1.379-1.379l-.004-.435s-.205-1.233 1.44-1.233L12.015 5v1.759a.83.83 0 0 0 1.17 0l2.573-2.572a.83.83 0 0 0 0-1.17L13.185.292a.83.83 0 0 0-1.17 0v1.845c-.161-.047-6.453-.074-6.453-.074c-3.711 0-4.523 2.429-4.523 3.407v1.031a1.5 1.5 0 0 0 1.5 1.5m11.961.041c-.828 0-1.5.584-1.5 1.412l.002.957c-.045.357-.645.59-1.525.59H4.938l-.002-1.559a.83.83 0 0 0-1.17 0l-2.572 2.572a.83.83 0 0 0 0 1.17l2.572 2.572a.826.826 0 0 0 1.17 0l.002-1.851h6.539c3.711 0 4.523-2.442 4.523-3.421V9.453c0-.827-.672-1.411-1.5-1.411"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowCircleRycycle(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m16.945 9.738l-1.439-.426c.007-.133.017-.264.017-.396c0-3.53-2.419-6.593-5.752-7.281L9.51 2.996c2.709.561 4.675 3.049 4.675 5.92v.006l-1.302-.385a.586.586 0 0 0-.213.781l1.182 2.207c.15.276.486.375.754.223l2.127-1.229a.586.586 0 0 0 .212-.781M2.438 8.626c0-2.353 1.326-4.454 3.3-5.44l.612 1.658a.589.589 0 0 0 .788-.314l1.071-2.49a.63.63 0 0 0-.304-.818l-2.4-1.11a.587.587 0 0 0-.788.313l.541 1.465C2.778 3.083 1.1 5.695 1.1 8.627c0 .824.135 1.646.404 2.443l1.262-.457a6.246 6.246 0 0 1-.328-1.987m6.275 5.981c-1.046 0-2.065-.263-2.955-.811l1.127-1.188a.59.59 0 0 0-.603-.595H3.656a.607.607 0 0 0-.593.625l.018 2.728a.606.606 0 0 0 .604.614l1.116-1.176a6.923 6.923 0 0 0 3.912 1.19c2.013 0 3.946-.895 5.305-2.454l-.992-.931c-1.106 1.27-2.678 1.998-4.313 1.998"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDown(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.16 6.246c.225 0 .45.062.65.196l6.229 4.156l6.037-4.197a1.175 1.175 0 0 1 1.304 1.958l-6.688 4.63a1.174 1.174 0 0 1-1.304.002l-6.88-4.589a1.178 1.178 0 0 1 .652-2.156"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowForward(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m15.584 5.988l-5.24-4.645c-.404-.41-.879-.479-1.283-.068v3.132l-.227-.001c-4.95 0-8.75 3.563-8.75 8.41c0 1.688.766 1.073 1.083.484c1.501-2.78 4.267-4.677 7.705-4.677l.188.001v3.065c.404.41.929.361 1.283.068l5.24-4.283c.405-.41.405-1.075.001-1.486"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowFourWay(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m16.844 7.747l-2.062-2.591a.551.551 0 0 0-.779 0v1.875H9.989V2.969h1.855a.558.558 0 0 0 0-.78L9.391.188a.541.541 0 0 0-.771 0L6.188 2.189a.554.554 0 0 0 0 .78H8.01v4.062H3.985V5.172a.551.551 0 0 0-.779 0S1.012 7.576 1.012 8.07c0 .428 2.194 2.68 2.194 2.68a.55.55 0 0 0 .779 0V8.984H8.01v4.047H6.172a.54.54 0 0 0 0 .76s2.332 2.192 2.826 2.192c.43 0 2.846-2.192 2.846-2.192a.54.54 0 0 0 0-.76H9.989V8.984h4.014v1.812a.551.551 0 0 0 .779 0l2.062-2.27a.551.551 0 0 0 0-.779"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowFullscreen(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.196.083h-3.867a.69.69 0 0 0-.688.69l1.588 1.594l-1.917 1.917a.979.979 0 0 0 0 1.39a.984.984 0 0 0 1.391 0l1.914-1.915l1.579 1.585c.38 0 .687-.31.687-.69V.773c0-.38-.307-.69-.687-.69m-.004 10.562l-1.619 1.612l-1.952-1.952a.983.983 0 1 0-1.392 1.39l1.951 1.95l-1.56 1.554c0 .38.309.687.69.687h3.881c.381 0 .69-.307.69-.687v-3.866a.686.686 0 0 0-.689-.688M4.758 2.359L6.342.78a.69.69 0 0 0-.691-.687h-3.88a.688.688 0 0 0-.69.687v3.866c0 .381.31.688.69.688l1.595-1.587l1.969 1.968a.978.978 0 0 0 1.39 0a.983.983 0 0 0 0-1.389zm1.193 7.286l-2.59 2.59l-1.594-1.601a.69.69 0 0 0-.688.69v3.881c0 .381.309.69.688.69h3.867a.69.69 0 0 0 .687-.69L4.75 13.627l2.592-2.592a.982.982 0 1 0-1.391-1.39"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowFullscreenTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.988 6.979A1.01 1.01 0 0 0 16 5.97V1.008a1.01 1.01 0 0 0-1.012-1.009h-4.977a1.01 1.01 0 0 0-1.012 1.009l1.58 1.575l-2.57 2.57l-2.57-2.57l1.58-1.575A1.01 1.01 0 0 0 6.007-.001H1.03A1.01 1.01 0 0 0 .018 1.008V5.97A1.01 1.01 0 0 0 1.03 6.979l1.59-1.585l2.574 2.574l-2.596 2.597L1.028 9C.471 9 .019 9.45.019 10.006v4.946c0 .555.452 1.006 1.009 1.006H5.99c.558 0 1.009-.451 1.009-1.006l-1.582-1.577l2.592-2.592l2.592 2.592l-1.582 1.577c0 .555.451 1.006 1.009 1.006h4.962c.557 0 1.009-.451 1.009-1.006v-4.946C15.999 9.45 15.547 9 14.99 9l-1.57 1.565l-2.596-2.597l2.574-2.574z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeft(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.978 1.162c0 .225-.062.45-.196.65L6.626 8.041l4.197 6.037c.359.541.213 1.27-.328 1.629a1.174 1.174 0 0 1-1.63-.325l-4.63-6.688a1.172 1.172 0 0 1-.002-1.304L8.822.51a1.178 1.178 0 0 1 2.156.652"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftRight(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 9h8.066v2.864l3.85-3.82l-3.85-4.013L13 7H5l-.114-2.97L1.03 8.07l3.856 3.866z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowMove(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.946 9.976c1.075 0 2.044-.89 2.044-1.988C9.99 6.89 9.02 6 7.946 6C6.871 6 6 6.89 6 7.988c0 1.098.871 1.988 1.946 1.988M5.21 3.984c-.198-.229-.236-.68-.034-.914L7.55.136a.466.466 0 0 1 .729-.005l2.537 2.951c.202.229.202.67 0 .902zm5.61 8.057a.47.47 0 0 1 0 .688l-2.506 3.104a.587.587 0 0 1-.725 0l-2.365-3.104a.467.467 0 0 1 0-.686l5.597-.002zm2.11-6.773l2.934 2.306a.5.5 0 0 1 .006.76l-2.948 2.411c-.228.207-.654.211-.886 0V5.262s.661-.204.894.006m-8.917 5.521c-.23.209-.688.207-.917 0L.164 8.33a.502.502 0 0 1 .004-.76l2.908-2.357c.229-.207.707-.209.937 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowReload(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.796 8.908L15.234 7.21a.553.553 0 0 0-.776 0l-1.564 1.698a.543.543 0 0 0 0 .772h1.294a5.345 5.345 0 0 1-3.789 3.79a5.378 5.378 0 0 1-5.767-2.119l-1.091.751a6.709 6.709 0 0 0 7.196 2.643A6.665 6.665 0 0 0 15.55 9.68h1.245a.544.544 0 0 0 .001-.772M5.475 8.021a.543.543 0 0 0 0-.772H4.018a5.339 5.339 0 0 1 3.771-3.738a5.373 5.373 0 0 1 5.766 2.121l1.092-.752a6.712 6.712 0 0 0-7.199-2.645a6.67 6.67 0 0 0-4.8 5.014H1.196a.543.543 0 0 0 0 .772l1.638 1.637a.553.553 0 0 0 .776 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowResizeFive(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.988 0h-4.977a1.01 1.01 0 0 0-1.012 1.009l1.58 1.575l-7.98 7.981L1.029 9C.472 9 .02 9.45.02 10.006v4.946c0 .555.452 1.006 1.009 1.006h4.962c.558 0 1.009-.451 1.009-1.006l-1.582-1.577l7.98-7.98l1.59 1.585A1.01 1.01 0 0 0 16 5.971V1.009A1.01 1.01 0 0 0 14.988 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowResizeFour(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.12 16.003L16 14.122l-3.038-3.057L15.018 9H9v6.047l2.087-2.097zM7 .969L4.913 3.065L1.88.013L0 1.894L3.038 4.95L1.042 6.917H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowResizeOne(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.995 10.852L5.133 9.008l-3.026 2.98L.062 9.972v5.903h5.987l-2.076-2.047zM9.961.008l2.097 2.087l-3.053 3.033l1.88 1.88l3.057-3.038l1.967 1.996V.008z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowResizeSix(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.031 0h4.977A1.01 1.01 0 0 1 7.02 1.009L5.44 2.584l7.98 7.981L14.99 9c.557 0 1.009.45 1.009 1.006v4.946c0 .555-.452 1.006-1.009 1.006h-4.962a1.007 1.007 0 0 1-1.009-1.006l1.582-1.577l-7.98-7.98l-1.59 1.585A1.01 1.01 0 0 1 .019 5.971V1.009A1.011 1.011 0 0 1 1.031 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowResizeThree(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.995 1.852L14.133.008l-3.026 2.98L9.062.972v5.903h5.987l-2.076-2.047zM.961 9.008l2.097 2.087l-3.053 3.033l1.88 1.88l3.057-3.038l1.967 1.996V9.008z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowResizeTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m5.075 6.95l1.843-1.862l-2.98-3.026L5.955.018H.052v5.986l2.046-2.076zm10.928 2.966l-2.171 2.097l-3.033-3.053l-1.881 1.881l3.039 3.056l-1.996 2.084h6.042z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRight(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.077 1.162c0 .225.062.45.196.65l4.156 6.229l-4.197 6.037a1.175 1.175 0 0 0 .328 1.629a1.174 1.174 0 0 0 1.63-.325l4.63-6.688c.264-.394.266-.908.002-1.304L8.233.51a1.178 1.178 0 0 0-2.156.652"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowShuffle(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.121 5.958h1.934v.854a.55.55 0 0 0 .778 0l1.965-1.352a.552.552 0 0 0 0-.778l-1.965-1.495a.548.548 0 0 0-.778 0v.849H12.09c-.195-.008-1.936-.032-3.238 1.222c-.857.824-1.292 1.662-1.292 3.103c0 .873-.226 1.534-.669 1.964c-.697.675-1.771.742-1.818.741H1.084v1.898l4.062.002c.451 0 1.955-.09 3.113-1.194c.861-.819 1.297-1.968 1.297-3.411c0-.873.226-1.226.672-1.662c.702-.686 1.86-.737 1.893-.741m4.719 5.729l-2.027-1.52a.55.55 0 0 0-.778 0v.914h-2.154s-.653.008-1.28-.282l-.909 1.653c.964.445 1.906.48 2.163.48l.063-.001h2.117v.901a.55.55 0 0 0 .778 0l2.027-1.369a.55.55 0 0 0 0-.776M6.555 6.329l1.052-1.618c-1.188-.666-2.445-.633-2.54-.63H1v1.89l4.111-.001c.012.004.778-.015 1.444.359"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowThickDown(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m8.327 15.886l5.447-5.94a.65.65 0 0 0-.002-.849l-3.841-.005V1.068c0-.553-.437-1-.976-1H7.004a.987.987 0 0 0-.976 1v8.02l-3.95-.005a.652.652 0 0 0 .004.848l5.485 5.954a.501.501 0 0 0 .76.001"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowThickLeft(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m.133 8.367l5.94 5.346a.648.648 0 0 0 .849-.002l.005-3.728h8.024a.99.99 0 0 0 1-.982V7.035a.99.99 0 0 0-1-.982h-8.02l.005-3.81a.65.65 0 0 0-.848.003L.134 7.603a.503.503 0 0 0-.001.764"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowThickRight(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m16.818 7.646l-5.94-5.44a.64.64 0 0 0-.849.002l-.005 3.793H2a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h8.019l-.006 3.877a.642.642 0 0 0 .849-.003l5.954-5.452a.518.518 0 0 0 .002-.777"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowThickThinDown(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M2.336 8.667c.199 0 .404.045.596.141l6.07 3.034l6.069-3.034a1.332 1.332 0 1 1 1.192 2.385l-6.666 3.333a1.332 1.332 0 0 1-1.193 0l-6.666-3.333a1.332 1.332 0 0 1 .598-2.526"/><path d="M2.336 4.334c.1 0 .201.022.297.07l6.369 3.184l6.367-3.184a.669.669 0 0 1 .895.298a.668.668 0 0 1-.299.895L9.299 8.93a.665.665 0 0 1-.596 0L2.037 5.597a.668.668 0 0 1 .299-1.263"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowThickThinUp(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M15.666 8.333c-.201 0-.404-.045-.596-.141L9.002 5.158L2.931 8.192a1.332 1.332 0 1 1-1.192-2.385l6.666-3.333a1.332 1.332 0 0 1 1.193 0l6.666 3.333a1.332 1.332 0 0 1-.598 2.526"/><path d="M15.666 12.666a.66.66 0 0 1-.297-.07L9.002 9.412l-6.369 3.184a.667.667 0 1 1-.596-1.193L8.703 8.07a.665.665 0 0 1 .596 0l6.666 3.333a.668.668 0 0 1-.299 1.263"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowThickUp(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m8.65 1.158l-5.485 5.94a.648.648 0 0 0 .002.849l3.868.005v8.024c0 .553.439 1 .982 1h1.965a.99.99 0 0 0 .982-1v-8.02l3.811.005a.65.65 0 0 0-.004-.848L9.414 1.159a.506.506 0 0 0-.764-.001"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowThinDown(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.031 15.97L2.339 9.665c-.459-.459-.296-1.16-.094-1.359c.403-.402.998-.342 1.402.061l4.438 3.896L8.086.948a.94.94 0 0 1 .952-.927a.937.937 0 0 1 .95.927l.013 11.212l4.271-3.854a1.036 1.036 0 0 1 1.461 0c.405.4.304.95-.101 1.352z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowThinLeft(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m1.042 8.01l6.305-6.693c.459-.459 1.16-.296 1.359-.094c.402.403.342.998-.061 1.402L4.749 7.063l11.315.002a.94.94 0 0 1 .927.952a.937.937 0 0 1-.927.95L4.852 8.98l3.854 4.271a1.036 1.036 0 0 1 0 1.461c-.4.405-.95.304-1.352-.101z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowThinLeftBottom(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m1.124 15.922l-.017-8.721c0-.648.611-1.029.895-1.027c.57.001.948.464.949 1.034l.01 5.347L15.355.355a.938.938 0 0 1 1.327.018A.936.936 0 0 1 16.7 1.7L4.554 14.012l5.333.008c.57.002 1.033.465 1.033 1.033c.003.57-.457.887-1.027.885z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowThinLeftTop(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.885.006c.573-.002 1.075.403 1.073.974c0 .568-.507.942-1.079.944l-5.354.008l12.193 12.312a.934.934 0 0 1-.018 1.327a.948.948 0 0 1-1.334.018l-12.44-12.2l-.011 5.569c-.001.57-.381 1.033-.952 1.034c-.285.002-.898-.379-.898-1.027L1.082.022z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowThinRight(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m16.992 7.98l-6.305 6.693c-.459.459-1.16.296-1.359.094c-.402-.403-.342-.998.061-1.402l3.895-4.438L1.97 8.925a.94.94 0 0 1-.927-.952a.937.937 0 0 1 .927-.95l11.212-.013l-3.855-4.271a1.036 1.036 0 0 1 0-1.461c.4-.405.95-.304 1.352.101z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowThinRightBottom(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.148 15.938c-.573.002-1.034-.314-1.032-.885c0-.568.465-1.031 1.038-1.033l5.353-.008L1.314 1.7A.934.934 0 0 1 1.331.373A.948.948 0 0 1 2.665.355l12.441 12.2l.01-5.347c.002-.57.381-1.033.953-1.034c.285-.002.898.379.898 1.027l-.017 8.721z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowThinRightTop(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m16.908.021l.017 8.931c0 .648-.611 1.029-.895 1.027c-.57-.001-.948-.464-.949-1.034l-.01-5.557l-12.394 12.2a.94.94 0 0 1-1.328-.018c-.372-.372-.381-.966-.017-1.327L13.478 1.931l-5.333-.008c-.57-.002-1.062-.376-1.062-.944C7.08.409 7.568.003 8.14.005z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowThinUp(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m9.002.02l6.693 6.305c.459.459.297 1.16.094 1.359c-.402.402-.998.342-1.402-.061L9.949 3.728l-.002 11.314a.941.941 0 0 1-.951.928a.94.94 0 0 1-.951-.928L8.033 3.831L3.762 7.685a1.036 1.036 0 0 1-1.461 0c-.404-.4-.303-.949.102-1.352z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowThreeWayOne(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M10.816 3.989a.577.577 0 0 0 0-.817L8.433.169a.578.578 0 0 0-.818 0L5.232 3.172a.577.577 0 0 0 0 .817h1.811v11.979h1.936V3.989z"/><path d="m14.516 5.008l-2.986.35a.459.459 0 0 0-.454.462l.808.821l-1.86 1.89v7.47h1.956V9.359l1.29-1.306l.903.92a.46.46 0 0 0 .455-.462l.342-3.042a.458.458 0 0 0-.454-.461M5.002 8.842a.463.463 0 0 0-.462-.462L1.5 8.031a.462.462 0 0 0-.463.461l.35 3.04c0 .256.207.463.462.463l.884-.883l1.314 1.622v3.197l1.9.052v-4.025L4.129 9.714z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowThreeWayTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.062 4.938L8.958.059L6.03 4.94H8v1.806c0 .74.234 1.161.974 1.161c.737 0 .995-.421.995-1.161V4.939zm-1.021 9.996l5.782-.006l-2.552-5.087l-1.055 1.662l-1.525-.968c-.625-.397-1.105-.424-1.502.2c-.396.622-.178 1.065.447 1.462l1.525.968zm-4.102-.084l-5.782-.006l2.552-5.087l1.055 1.662l1.525-.968c.625-.397 1.106-.424 1.502.2c.395.622.178 1.065-.447 1.462l-1.525.968z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowTriangleRecycle(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.9 12.815L14.873 9.48A.608.608 0 0 0 14 9.281a.693.693 0 0 0-.187.926l1.398 2.357h-2.996v-1.439c-.12-.129-.309-.137-.422-.017l-2.04 1.827c-.112.121-.106.322.015.451l2.023 1.475c.122.128.311.135.423.014v-1.026h4.154a.625.625 0 0 0 .554-.35a.707.707 0 0 0-.024-.684zm-10.957-.207H2.861l1.68-2.752l.883.656a.322.322 0 0 0 .26-.37L5.55 7.316a.31.31 0 0 0-.365-.254l-2.515.832a.327.327 0 0 0-.258.372l1.07.799l-2.3 3.801a.724.724 0 0 0-.024.704a.64.64 0 0 0 .569.361H6c.357 0 .647-.31.647-.691c0-.381-.345-.632-.704-.632m4.079-5.991l2.458.842c.158.032.31-.087.342-.266l.157-2.716c.03-.179-.071-.348-.227-.379l-.9.65l-2.05-3.37c-.23-.382-.826-.383-1.058.001L6.84 4.512a.696.696 0 0 0 .187.927a.609.609 0 0 0 .873-.2l1.374-2.261l1.535 2.525l-1.013.732c-.033.18.068.35.226.382"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowTwoLeftRight(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.688 7.716a.842.842 0 0 0 1.17 0V5.992h9.641c.828 0 1.5-.66 1.5-1.472c0-.812-.672-1.472-1.5-1.472H4.858V1.249a.839.839 0 0 0-1.17 0l-2.48 2.708a.799.799 0 0 0 0 1.146zm12.07 3.114l-2.573-2.538a.834.834 0 0 0-1.171 0v1.731H2.583c-.828 0-1.5.664-1.5 1.481c0 .817.672 1.481 1.5 1.481h9.431v1.732a.838.838 0 0 0 1.171 0l2.573-2.538c.322-.318.322-1.031 0-1.349"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowTwoUp(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m8.973 5.175l5.002 3.744v-3.94L8.973 1.036L4.004 5.078v3.918l.012.011zm0 7.783l5.002 3.951v-3.938L8.973 9.005l-4.969 4.064v3.918l.012.01z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowTwoWay(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.8 10.167a.55.55 0 0 0-.779 0v.888H9.98c-.345-.073-2.035-.643-2.035-3.144c0-2.456 1.563-2.894 1.972-2.96H13v.903a.551.551 0 0 0 .779 0l2.06-1.427a.551.551 0 0 0 0-.779l-2.06-1.481a.55.55 0 0 0-.779 0v.847l-3.261.006c-1.198.131-3.314 1.176-3.711 4.022H1c-.553 0-1 .377-1 .929a1 1 0 0 0 1 1h5.044c.437 2.75 2.53 3.814 3.698 3.989l3.278.01v.863a.551.551 0 0 0 .779 0l2.039-1.405a.551.551 0 0 0 0-.779z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowTwoWayLeftRight(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.763 2.989L14.203.385c-.294-.3-.918-.3-1.211 0v1.814c-1.459.062-3.101.913-3.992 2.144c-.893-1.23-2.651-2.082-4.109-2.144V.385c-.295-.3-.918-.3-1.212 0L1.237 2.989a.776.776 0 0 0 0 1.085l2.442 2.603c.294.3.917.3 1.212 0V4.915c1.401.088 3.156 1.255 3.156 2.689v7.333h1.875V7.604c0-1.435 1.668-2.602 3.07-2.689v1.762c.293.3.917.3 1.211 0l2.56-2.603a.776.776 0 0 0 0-1.085"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowTwoWayRight(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.276 6.716a.842.842 0 0 1-1.17 0V4.992H2.464c-.827 0-1.5-.66-1.5-1.472c0-.812.673-1.472 1.5-1.472h9.642V.249a.839.839 0 0 1 1.17 0l2.48 2.708a.799.799 0 0 1 0 1.146zm2.482 4.114l-2.573-2.538a.834.834 0 0 0-1.171 0v1.731H2.583c-.828 0-1.5.664-1.5 1.481c0 .817.672 1.481 1.5 1.481h9.431v1.732a.838.838 0 0 0 1.171 0l2.573-2.538c.322-.318.322-1.031 0-1.349"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowTwoWayRightBottom(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.807 2.894L14.049.135a.666.666 0 0 0-.944.005l-.01 1.963H4.383c-.738 0-1.336.599-1.336 1.335v8.622l-1.912.012a.666.666 0 0 0 .006.943l2.787 2.786a.663.663 0 0 0 .943.006l2.758-2.758c.26-.26.258-.756-.005-1.018l-1.666.012V5.941c0-.738.272-.983 1.011-.983h6.112l-.01 1.671c.26.259.681.257.942-.005l2.789-2.787a.667.667 0 0 0 .005-.943"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUp(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.812 9.896a1.15 1.15 0 0 1-.65-.197l-6.23-4.156L2.895 9.74a1.175 1.175 0 0 1-1.629-.328a1.173 1.173 0 0 1 .326-1.629L8.28 3.152a1.175 1.175 0 0 1 1.303-.002l6.881 4.59a1.176 1.176 0 0 1-.652 2.156"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpDown(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.021 3.961V12H5.062l3.819 3.991L12.896 12H10V3.961h2.896L8.857.105L4.991 3.961z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowWave(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.031 7.125V9h-2.062l-.004-4H8.016l.027 6H5.984L5.98 7.031H2.035L2.031 9H1.004v1h1.949l.004-2h2.07l.004 4h3.922l.004-6h2.059l.004 4h3.011v1.958L17 9.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArtBoard(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.953 12.938H3.046V3.022h9.907zm-8.99-.876h8.135V3.833H3.963zM12 0h.959v1.943H12zM3 0h.959v1.984H3zm9 14h.938v1.922H12zm-9 0h.959v1.984H3zm11-2h1.875v.994H14zM0 12h1.855v.918H0zm14-9h1.98v.938H14zM0 3h1.938v.957H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Askterisk(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.991 6.153c-.351 0-1.952.454-3.948.788c1.422-1.581 2.768-2.755 2.961-3.062c.576-.917.235-2.092-.763-2.622c-.999-.53-2.274-.216-2.852.702c-.19.305-.622 1.949-1.39 3.86c-.768-1.911-1.199-3.556-1.388-3.86c-.579-.918-1.854-1.232-2.852-.702c-.998.531-1.34 1.705-.762 2.622c.19.307 1.537 1.48 2.961 3.062c-1.998-.334-3.6-.788-3.95-.788c-1.08 0-1.955.842-1.955 1.882c0 1.038.875 1.88 1.955 1.88c.351 0 1.93-.448 3.907-.782c-1.408 1.557-2.731 2.711-2.92 3.014c-.576.92-.234 2.092.764 2.624c.998.528 2.273.216 2.851-.703c.19-.304.622-1.948 1.39-3.86c.769 1.912 1.199 3.557 1.392 3.86c.575.919 1.851 1.233 2.85.703c.998-.531 1.339-1.704.763-2.622c-.189-.305-1.511-1.459-2.918-3.016c1.978.334 3.557.782 3.905.782c1.08 0 1.956-.842 1.956-1.88c-.001-1.04-.877-1.882-1.957-1.882"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AtmCard(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.666 3.07H2.333C1.597 3.07 1 3.641 1 4.347v.639h16v-.639c0-.705-.598-1.277-1.334-1.277M1 11.648c0 .731.597 1.323 1.333 1.323h13.333c.736 0 1.334-.592 1.334-1.323V7.07H1zm6-2.745h4v1.125H7zm-2.042 2.03H13v1.096H4.958z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Axe(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M7.038 8.5c-.671.672-2.108.652-4.198-1.438C.752 4.974.793 3.474 1.402 2.865c.61-.61.883.455 2.363-.471C5.244 1.466 5.564.487 6.081 1.003L8.897 3.82c.518.517-.797 1.173-1.39 2.316c-.549 1.065.202 1.694-.469 2.364m9.259 6.457c-.142.143-.461.053-.715-.201l-6.961-6.96c-.252-.253-.344-.573-.201-.716l.577-.576c.141-.143.462-.051.716.201l6.959 6.961c.254.253.344.572.201.715z"/><path d="M3.008.137h.972v1.702h-.972z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Baby(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.975 1.988A1.979 1.979 0 0 1 9 3.971a1.979 1.979 0 0 1-1.972-1.983C7.028.892 7.914.005 9 .005a1.98 1.98 0 0 1 1.975 1.983m-.121 8.044c.3.31.104.456-.005.565l-1.023 1.132c-.299.311-1.379.312-1.677.002l-.999-1.113c-.06-.061-.297-.254.005-.566zm.021-4.682a1.369 1.369 0 0 0-.862-.315h-2.02c-.328 0-.63.126-.876.328a503.43 503.43 0 0 0-2.036 2.774a.464.464 0 0 0 .087.589a.346.346 0 0 0 .277.07a.365.365 0 0 0 .239-.171L7.02 7.581v.979s-.182.397.459.397h3.11c.348 0 .41-.397.41-.397v-.951l1.336 1.028a.373.373 0 0 0 .241.172a.343.343 0 0 0 .276-.071a.458.458 0 0 0 .089-.589s-2.01-2.769-2.066-2.799M7.421 15.964c-.173 0-.524-.126-.626-.3l-1.266-2.309c-.145-.244.912-2.254.912-2.254l1.154 1.281s-.567.655-.586.84c-.02.185.854 1.793.854 1.793c.167.276.103.656-.144.844a.478.478 0 0 1-.298.105m3.11 0c.174 0 .525-.126.628-.3l1.265-2.309c.145-.244-.911-2.254-.911-2.254l-1.154 1.281s.568.655.587.84c.019.185-.855 1.793-.855 1.793c-.167.276-.103.656.143.844a.477.477 0 0 0 .297.105"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BabyMilkBotl(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.443 5H6.566c-.947 0-1.518.674-1.518 1.453v8.133c0 .776.766 1.406 1.711 1.406h4.49c.946 0 1.711-.63 1.711-1.406V6.453C12.961 5.674 12.39 5 11.443 5m.577 9.243c0 .436-.525.792-1.17.792H7.14c-.646 0-1.169-.356-1.169-.792V7.574c0-.22.141-1.214 3.059-.329c2.92.889 2.99.113 2.99.329zM10.982 2.958L9.71 1.992V.683c0-.382-.332-.691-.741-.691c-.408 0-.739.31-.739.691v1.284l-1.181.99l-.031-.004V4h3.965V2.954v.004z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BabyStroller(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M4.002 12.037a1.962 1.962 0 0 0-1.967 1.955c0 1.08.882 1.956 1.967 1.956a1.962 1.962 0 0 0 1.967-1.956a1.962 1.962 0 0 0-1.967-1.955m0 3.174a1.222 1.222 0 0 1-1.225-1.219c0-.674.549-1.219 1.225-1.219s1.227.545 1.227 1.219s-.551 1.219-1.227 1.219m7.991-3.174a1.955 1.955 0 1 0 .001 3.913a1.955 1.955 0 0 0-.001-3.913m0 3.176a1.22 1.22 0 1 1 .003-2.439a1.22 1.22 0 0 1-.003 2.439M11.121 0S8.65 3.41 8.338 4.908h7.57C15.908-.643 13.117 0 11.121 0"/><path d="M15.943 6.107H6.396L4.093 4.041H.509c-.28 0-.507.208-.507.465c0 .256.227.463.507.463h3.163l1.27 1.139H.068C.136 9.29-.68 11.92 7.658 11.92c8.719 0 8.272-2.6 8.285-5.813"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BabyTroller(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.954.007c-1.043 0-1.792.359-1.89 1.28L15 3.526c-.812-.958-2.083-1.786-3.346-2.239c-.471-.169-1.607-.461-1.607.147V5.29c0 .382-.209.755-.594.755H1.762c-.494 0-.72.548-.72 1.243c0 1.451.796 2.752 3.492 3.342l1.474 2.41H4.736a1.926 1.926 0 0 0-1.69-1.014c-1.065 0-2.046.866-2.046 1.933c0 1.068.98 2.039 2.046 2.039c1.061 0 1.92-.96 1.931-2.02h6.045c.002 1.082.88 2.022 1.962 2.022C14.066 16 15 15.057 15 13.973c0-1.083-.934-1.959-2.016-1.959c-.744 0-1.383.417-1.716 1.026h-1.221l1.467-2.403c3.253-.739 4.455-2.64 4.455-4.711c0-.138-.069-4.343-.069-4.343c0-.49.403-.606 1.054-.606zm-9.74 13.035L5.43 10.794c.802.113 1.736.174 2.827.174c.874 0 1.65-.061 2.355-.164l-1.794 2.238z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BackPack(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M9.033 6c-1.074 0-1.938.331-2.568.984c-1.451 1.507-1.443 3.436-1.44 4.996l.001.216c0 .805.281 1.254.941 1.501c.646.242 1.576.261 2.726.261h.682c2.55 0 3.666-.166 3.666-1.762l.001-.216c.004-1.561.01-3.489-1.44-4.996C10.971 6.331 10.106 6 9.033 6"/><path d="M14.945 10.107c-.095-1.677-.371-3.577-1.139-5.116c.065.009.126.025.195.025c.544 0 .98-.276.98-.617v-.757c0-.342-.437-.618-.98-.618c-.541 0-.981.276-.981.618v.134a4.663 4.663 0 0 0-2.057-1.422v-1.32c0-.543-.509-.984-1.134-.984H8.183c-.625 0-1.135.441-1.135.984v1.323a4.689 4.689 0 0 0-2.126 1.521v-.236c0-.342-.426-.618-.954-.618c-.524 0-.952.276-.952.618v.757c0 .341.428.617.952.617c.09 0 .17-.018.252-.033c-.755 1.509-1.037 3.368-1.137 5.023c-2.139.031-2.066 2.221-2.066 3.114c0 .895.944.871 2.107.871V13.8c.497 2.198 2.941 2.142 5.887 2.142c2.901 0 5.318.059 5.867-2.035c1.167 0 2.098-.009 2.098-.835c.001-.832.058-2.836-2.031-2.965M7.953 1.313c0-.187.223-.345.488-.345h1.1c.266 0 .49.158.49.345v.803a6.422 6.422 0 0 0-1.02-.086c-.379 0-.726.039-1.059.094zm6.09 11.344c0 2.425-2.387 2.425-4.693 2.425h-.738c-1.315 0-2.387-.021-3.197-.298c-1.004-.344-1.492-1.04-1.492-2.127l-.002-.211c-.002-1.645-.01-4.706 1.771-6.397c.822-.778 1.928-1.174 3.291-1.174c1.36 0 2.465.396 3.289 1.174c1.782 1.691 1.774 4.753 1.771 6.397z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BackwardPage(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="m.252 4.516l3.905-3.25c.276-.288.593-.335.871-.048c0 0-.037 2.19-.006 2.19c.679 0 1.083.113 1.7.314A6.329 6.329 0 0 1 8.93 4.99c1.276 1.123 2.093 2.402 2.093 4.213c0 1.182-.523.752-.742.339c-1.035-1.945-2.923-2.928-5.282-2.928c-.03 0 .003 2.146.003 2.146a.607.607 0 0 1-.87.047l-3.88-3.25a.752.752 0 0 1 0-1.041"/><path d="M1.954 9.079v4.983h13.118V4.937H9.871c-.275-.312-.619-.717-1.219-.919H16v10.964H1.08V8.258z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BadgeName(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.997 5.883c-.528 0-.957-.212-.957-.991V1.856c0-.781.429-.856.957-.856s.956.075.956.856v3.036c0 .779-.428.991-.956.991"/><path d="m11 4.938l.058.615c0 1.188-.922 1.497-2.059 1.497c-1.139 0-2.063-.309-2.063-1.497V5l-3.889-.062a1.986 1.986 0 0 0-2.006 2v7c0 1.132.896 1.992 2.006 1.992h11.908c1.105 0 2.003-.918 2.003-2.05V7.05c0-1.131-.897-2.05-2.003-2.05zm-4.5 3c.827 0 1.5.688 1.5 1.536c0 .846-.673 2.464-1.5 2.464c-.829 0-1.5-1.62-1.5-2.464c0-.848.672-1.536 1.5-1.536m3.483 6H3.002s-.12-1.987 1.672-1.987c.373.573.887 1.126 1.834 1.126c.949 0 1.399-.557 1.77-1.139c2.017 0 1.705 2 1.705 2M15 13.896h-4v-1h4zm0-2.004h-4v-1h4zm0-1.996h-4v-1h4z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bag(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M9.494 4.011a.49.49 0 0 1-.489-.492a2.512 2.512 0 0 0-2.503-2.514A2.511 2.511 0 0 0 4 3.519a.49.49 0 1 1-.979 0C3.021 1.59 4.583.021 6.502.021c1.92 0 3.482 1.569 3.482 3.498a.49.49 0 0 1-.49.492"/><path d="M12.492 4.011a.49.49 0 0 1-.488-.492c0-1.386-1.119-2.514-2.492-2.514a.49.49 0 0 1-.488-.492a.49.49 0 0 1 .488-.492c1.912 0 3.469 1.569 3.469 3.498a.491.491 0 0 1-.489.492m.55-.911V16h1.329c.327 0 .594-.254.594-.566V3.626c0-.314-.267-.565-.594-.565z"/><path d="M1.645 3.061c-.33 0-.601.252-.601.565v11.808c0 .312.271.566.601.566h10.306l.006-12.939z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BagPlus(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M9.494 4.011a.49.49 0 0 1-.489-.492a2.512 2.512 0 0 0-2.503-2.514A2.511 2.511 0 0 0 4 3.519a.49.49 0 1 1-.979 0C3.021 1.59 4.583.021 6.502.021c1.92 0 3.482 1.569 3.482 3.498a.49.49 0 0 1-.49.492"/><path d="M12.492 4.011a.49.49 0 0 1-.488-.492c0-1.386-1.119-2.514-2.492-2.514a.49.49 0 0 1-.488-.492a.49.49 0 0 1 .488-.492c1.912 0 3.469 1.569 3.469 3.498a.491.491 0 0 1-.489.492M11 13h5v1h-5z"/><path d="M13 11h1v5h-1zm1.371-8h-1.329v7h1.923V3.565c0-.313-.267-.565-.594-.565M10 12h1.951l.005-9l-10.312.061c-.33 0-.601.252-.601.565V15.4c0 .312.271.6.601.6H12s-.05-.416-.049-1H10z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BagRemove(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M9.494 4.011a.49.49 0 0 1-.489-.492a2.512 2.512 0 0 0-2.503-2.514A2.511 2.511 0 0 0 4 3.519a.49.49 0 1 1-.979 0C3.021 1.59 4.583.021 6.502.021c1.92 0 3.482 1.569 3.482 3.498a.49.49 0 0 1-.49.492"/><path d="M12.492 4.011a.49.49 0 0 1-.488-.492c0-1.386-1.119-2.514-2.492-2.514a.49.49 0 0 1-.488-.492a.49.49 0 0 1 .488-.492c1.912 0 3.469 1.569 3.469 3.498a.491.491 0 0 1-.489.492M11 14h5v1h-5zm3.371-10.939h-1.329v9.897h1.923V3.626c0-.313-.267-.565-.594-.565"/><path d="M11.951 13.047L11.956 3H1.644a.614.614 0 0 0-.601.6v11.773c0 .312.227.627.557.627H10v-2.953z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Balance(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.953 14H6.016v1.031H5V16h7v-.969h-1.047zM8.016 4v8.953h.953V4H13.5L8.969 2.833V1.016h-.953v1.817L3.469 4zM3.492 5.979l1.9 4.005l.566-.207l-2.225-4.609c-.053-.098-.157-.16-.298-.152c-.128.005-.239.075-.279.175L1.012 9.8l.588.162zm9.69-.784l-2.158 4.619l.592.162l1.902-3.99L15.43 10l.57-.208l-2.238-4.619c-.053-.097-.16-.159-.299-.151c-.13.003-.24.075-.281.173m2.802 5.836c0 1.061-1.112 1.922-2.484 1.922c-1.372 0-2.484-.861-2.484-1.922zm-10 0c0 1.061-1.112 1.922-2.484 1.922c-1.372 0-2.484-.861-2.484-1.922z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Balloon(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.998 13.272c-1.634 0-1.838-.415-1.863-.497c-.102-.352.514-1.103.996-1.512l-.027-.041c2.15-.667 3.807-4.533 3.807-6.773A4.446 4.446 0 0 0 7.469 0a4.446 4.446 0 0 0-4.442 4.449c0 2.35 1.818 6.486 4.121 6.849c-.393.462-.787 1.106-.611 1.702c.197.674.98.988 2.463.988c.123 0 .18.147.068.305c-.332.457-1.217.991-2.62.991V16c2.171 0 3.374-1.169 3.374-1.977a.73.73 0 0 0-.255-.572a.88.88 0 0 0-.569-.179M4.643 4.648c0 .688-.596.396-.596-.604c0-1.688 1.332-3.058 2.978-3.058c.978 0 1.269.613.595.613c-1.643.001-2.977 1.36-2.977 3.049"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bandage(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M.927 10.543C-.238 11.711-.25 13.59.902 14.742l.354.354c1.151 1.15 3.029 1.139 4.195-.027l1.536-1.534l-4.525-4.523zm14.202-9.321l-.354-.354C13.623-.283 11.742-.27 10.57.899L9.008 2.463l4.527 4.525l1.562-1.563c1.169-1.169 1.182-3.052.032-4.203M3.469 7.969L8 12.469l4.562-4.438l-4.588-4.646zm4.557-1.943H6.979V4.979h1.047zm2 2.016H8.979V6.98h1.047zm-.015 2H8.98V8.98h1.031zm-2-2H6.98V6.98h1.031zm-2 0H4.98V6.98h1.031zm2.015 2H6.979V8.98h1.047z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bank(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16 14v-1h-1.002v-1H14V6.032h-.969V12H11V6.032h-1V12H8V6.032H7V12H5V6.032h-.99L4 12H3v1H2v1H1v1h16v-1zM3.021 6h11.958V5h1V4h-1.01L9.031.094L3.031 4H2v1h1.021z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Barcode(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14 2h1.953v10.914H14zm-3 0h1.967v10.914H11zM9 2h.95v10.914H9zM5 2h1.972v10.914H5zM3 2h.973v10.914H3zM1 2h.973v10.914H1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Barrier(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.703 2a.67.67 0 0 0-.667.673v3.494c0 .211.1.39.25.514a.649.649 0 0 0 .417.159s.305.019.305.028V15h1.007v-2h11.999v2h.994V6.868c0-.009.263-.028.263-.028a.67.67 0 0 0 .667-.673V2.673A.672.672 0 0 0 16.271 2zm9.308 1H13l-2.01 3H9zM6.729 3h1.713L6.713 6H5zm-1.752 7h2.039l-.963 2H4.016l.96-2zm4 0h2.039l-.963 2H8.016l.96-2zm4 0h2.039l-.963 2h-2.037zM2.095 5.71A.624.624 0 0 1 2 5.392V3.61c0-.337.26-.61.582-.61H4L2.582 6a.572.572 0 0 1-.487-.29M15 7.023V9H3V7h11.042c-.001.01.958.017.958.024zm.357-1.038h-1.31l1.762-3c.166.117.279.29.279.49v1.872c0 .354-.326.638-.731.638"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Baseball(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M3.518 14.562c.916-.816 2.194-1.372 3.345-1.599c1.734-.341 3.031-1.979 3.019-3.812A3.159 3.159 0 0 0 8.916 6.9a3.104 3.104 0 0 0-2.14-.839l-.153.003c-1.808.081-2.85 1.008-3.376 3.007c-.318 1.22-1.033 2.495-1.836 3.4a7.997 7.997 0 0 0 2.107 2.091"/><path d="M8 .046C3.591.046.016 3.603.016 7.989a7.86 7.86 0 0 0 .764 3.373c.545-.703 1.057-1.64 1.307-2.596c.663-2.52 2.13-3.795 4.484-3.899l.201-.005c1.119 0 2.176.417 2.977 1.175a4.36 4.36 0 0 1 1.332 3.105c.018 2.4-1.695 4.548-3.985 4.996c-.876.174-1.776.544-2.464 1.045a7.975 7.975 0 0 0 3.369.749c4.409 0 7.984-3.557 7.984-7.943C15.985 3.603 12.409.046 8 .046"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BaseballStick(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m3.291 15.114l-.275-.277c3.4-4.055 6.998-5.629 9.898-8.164c2.951-2.579 3.102-2.885 3.45-3.234c.822-.822.87-2.105.109-2.867c-.763-.763-2.047-.715-2.869.107c-.349.348-.655.5-3.233 3.45c-2.535 2.901-4.457 6.485-8.185 9.881l-.276-.277c-.136-.136-.431-.06-.66.169c-.229.229-.305.523-.168.66l1.382 1.381c.136.137.431.061.66-.168c.228-.229.305-.524.167-.661"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Basket(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.244 5.084h-1.928l-1.035-3.637a.614.614 0 0 0-.838-.286l-.071.036a.657.657 0 0 0-.289.863l.741 3.023H5.17l.768-3.019a.658.658 0 0 0-.281-.865l-.07-.036a.616.616 0 0 0-.842.277L3.672 5.083h-1.89c-.4 0-.725.302-.725.672c0 0 .324 1.22.725 1.22h.636l.449 6.677s.052 1.309 2.113 1.309h7.021c2.145 0 2.115-1.336 2.115-1.336l.363-6.649h.766c.4 0 .725-1.22.725-1.22c-.001-.37-.326-.672-.726-.672M4.031 7.031H2.969V5.969h1.062zm2.031 5.711c0 .35-.251.633-.559.633H5.48c-.311 0-.561-.283-.561-.633V8.413c0-.35.25-.632.561-.632h.023c.308 0 .559.282.559.632zm2.998-.195c0 .334-.237.605-.529.605H8.51c-.293 0-.529-.271-.529-.605V8.448c0-.334.236-.604.529-.604h.021c.292 0 .529.271.529.604zm2.968.092c0 .338-.245.611-.546.611h-.022c-.302 0-.545-.273-.545-.611V8.483c0-.338.243-.611.545-.611h.022c.301 0 .546.273.546.611zm2.003-5.577h-1.094V5.968h1.094z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BasketArrowDown(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.975 12.953h1.979V11h.306l.22-4.024h.766c.4 0 .725-1.22.725-1.22c0-.37-.324-.672-.725-.672h-1.928l-1.035-3.709a.614.614 0 0 0-.838-.286l-.071.036a.657.657 0 0 0-.289.863l.741 3.096H4.172L4.94 2a.658.658 0 0 0-.281-.865l-.07-.036a.616.616 0 0 0-.842.277L2.674 5.084H.784c-.4 0-.725.302-.725.672c0 0 .324 1.22.725 1.22h.636l.449 6.677s.052 1.309 2.113 1.309h7.021c.574 0 .984-.1 1.289-.24zm.963-6.984h1.094v1.094h-1.094zM3.031 7.031H1.969V5.969h1.062zm2 5.554c0 .333-.237.603-.527.603h-.023c-.293 0-.529-.27-.529-.603V8.461c0-.333.236-.602.529-.602h.023c.29 0 .527.269.527.602zm3.03-.038c0 .334-.246.605-.543.605h-.024c-.301 0-.544-.271-.544-.605V8.448c0-.334.243-.604.544-.604h.024c.297 0 .543.271.543.604zm2.453.672h-.021c-.293 0-.529-.271-.529-.607V8.479c0-.336.236-.607.529-.607h.021c.291 0 .53.271.53.607v4.133c0 .335-.239.607-.53.607m2.517.828l1.469 1.87l1.488-1.87H15V12h-.965v2.047z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BasketArrowLeft(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m10.461 13.486l2.725-2.14l.293-5.371h.766c.4 0 .725-1.22.725-1.22c0-.37-.363-.711-.764-.711h-1.928L11.271.374a.614.614 0 0 0-.838-.286l-.071.036a.657.657 0 0 0-.289.863l.713 3.057H4.132L4.918.979a.658.658 0 0 0-.282-.865L4.567.078a.616.616 0 0 0-.842.277L2.633 4.044H.743c-.4 0-.686.341-.686.711c0 0 .324 1.22.725 1.22h.636l.449 6.677s.052 1.309 2.113 1.309h7.021c.019 0 .033-.004.052-.004zm1.477-8.517h1.094v1.094h-1.094zM3.031 6.031H1.969V4.969h1.062zm2.031 5.711c0 .35-.251.633-.559.633H4.48c-.311 0-.561-.283-.561-.633V7.413c0-.35.25-.632.561-.632h.023c.308 0 .559.282.559.632zm2.998-.195c0 .334-.237.605-.529.605H7.51c-.293 0-.529-.271-.529-.605V7.448c0-.334.236-.604.529-.604h.021c.292 0 .529.271.529.604zm1.856.092V7.483c0-.338.243-.611.545-.611h.022c.301 0 .546.273.546.611v4.156c0 .338-.245.611-.546.611h-.022c-.302 0-.545-.273-.545-.611m4.027.383l-1.87 1.469l1.87 1.488v-.988h2.047v-.965h-2.047z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BasketArrowRight(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.979 13.012c-.1.142-.323.238-.497.238h-.022c-.302 0-.545-.273-.545-.611V8.483c0-.338.243-.611.545-.611h.022c.301 0 .546.273.546.611v4.156a.645.645 0 0 1-.111.352h2.077V10.94l.257.2l.228-4.164h.766c.4 0 .725-1.22.725-1.22c0-.37-.324-.672-.725-.672h-1.928l-1.035-3.637a.614.614 0 0 0-.838-.286l-.071.036a.657.657 0 0 0-.289.863l.741 3.023H4.171l.768-3.019a.658.658 0 0 0-.281-.865l-.07-.036a.616.616 0 0 0-.842.277L2.673 5.083H.783c-.4 0-.725.302-.725.672c0 0 .324 1.22.725 1.22h.636l.449 6.677s.052 1.309 2.113 1.309h7v-1.949zm.959-7.043h1.094v1.094h-1.094zM3.031 7.031H1.969V5.969h1.062zm2.031 5.711c0 .35-.251.633-.559.633H4.48c-.311 0-.561-.283-.561-.633V8.413c0-.35.25-.632.561-.632h.023c.308 0 .559.282.559.632zm2.998-.195c0 .334-.237.605-.529.605H7.51c-.293 0-.529-.271-.529-.605V8.448c0-.334.236-.604.529-.604h.021c.292 0 .529.271.529.604zm6.01 3.412l1.868-1.448l-1.868-1.467v.974h-2.041v.951h2.041z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BasketArrowUp(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m11.018 14.956l2.184-2.915l.277-5.091h.766c.4 0 .725-1.224.725-1.224c0-.371-.324-.674-.725-.674h-1.928l-1.035-3.646a.614.614 0 0 0-.838-.287l-.071.036a.66.66 0 0 0-.289.866l.741 3.031H4.171l.768-3.026a.662.662 0 0 0-.281-.868l-.07-.036a.615.615 0 0 0-.842.278L2.673 5.052H.783c-.4 0-.725.303-.725.674c0 0 .324 1.224.725 1.224h.636l.449 6.695s.052 1.312 2.113 1.312h7.021c.004.001.009-.001.016-.001m.92-9.016h1.094v1.097h-1.094zM3.031 7.006H1.969V5.941h1.062zm2.031 5.553c0 .332-.251.601-.559.601H4.48c-.311 0-.561-.269-.561-.601V8.452c0-.332.25-.6.561-.6h.023c.308 0 .559.268.559.6zm2.999-.022c0 .335-.251.607-.559.607h-.024c-.31 0-.56-.272-.56-.607v-4.11c0-.335.25-.606.56-.606h.024c.308 0 .559.271.559.606zm2.463.601h-.022c-.301 0-.544-.27-.544-.602V8.449c0-.333.243-.602.544-.602h.022c.302 0 .546.269.546.602v4.087c0 .332-.244.602-.546.602m5.465.846L14.51 12l-1.479 1.984h.99v1.995h.963v-1.995z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BasketChecked(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" fill-rule="evenodd"><path fill="currentColor" d="m9.652 13.02l.28-.289c-.005-.031-.017-.059-.017-.092V8.483c0-.338.243-.611.545-.611h.021c.301 0 .546.273.546.611v3.113l.7-.725l1.125 1.117l.372-.386l.253-4.627h.766c.4 0 .725-1.22.725-1.22c0-.37-.324-.672-.725-.672h-1.928L11.28 1.446a.614.614 0 0 0-.838-.286l-.071.036a.657.657 0 0 0-.289.863l.741 3.023H4.169l.768-3.019a.658.658 0 0 0-.281-.865l-.07-.036a.616.616 0 0 0-.842.277L2.671 5.082H.781c-.4 0-.725.302-.725.672c0 0 .324 1.22.725 1.22h.636l.449 6.677s.052 1.309 2.113 1.309H11c.215 0 .393-.021.566-.046zm2.286-7.051h1.094v1.094h-1.094zM3.031 7.031H1.969V5.969h1.062zm2.031 5.711c0 .35-.251.633-.559.633H4.48c-.311 0-.561-.283-.561-.633V8.413c0-.35.25-.632.561-.632h.023c.308 0 .559.282.559.632zm2.999-.195c0 .334-.237.605-.529.605h-.021c-.293 0-.529-.271-.529-.605V8.448c0-.334.236-.604.529-.604h.021c.292 0 .529.271.529.604z"/><path stroke="currentColor" d="M11.092 12.972L13 14.958l2.943-3.972"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BasketError(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="m11.504 13.524l-.566-.567c-.098.172-.261.293-.455.293h-.021c-.302 0-.545-.273-.545-.611V8.483c0-.338.243-.611.545-.611h.021c.301 0 .546.273.546.611v2.541l.991-.991l1.225 1.227l.234-4.342h.766c.4 0 .725-1.161.725-1.161c0-.37-.324-.672-.725-.672h-1.928l-1.035-3.729a.614.614 0 0 0-.838-.286l-.071.036a.657.657 0 0 0-.289.863l.741 3.116H4.171l.768-3.084a.658.658 0 0 0-.281-.865l-.07-.036a.616.616 0 0 0-.842.277L2.673 5.085H.783c-.4 0-.725.302-.725.672c0 0 .324 1.161.725 1.161h.636l.449 6.735s.052 1.309 2.113 1.309h6.089zm.434-7.555h1.094v1.094h-1.094zM3.031 7.031H1.969V5.969h1.062zm2.031 5.711c0 .35-.251.633-.559.633H4.48c-.311 0-.561-.283-.561-.633V8.413c0-.35.25-.632.561-.632h.023c.308 0 .559.282.559.632zm3.04-.195c0 .334-.26.605-.578.605h-.025a.593.593 0 0 1-.58-.605V8.448c0-.334.259-.604.58-.604h.025c.318 0 .578.271.578.604z"/><path d="M15.969 15.281L14.17 13.48l1.78-1.78l-.667-.663l-1.779 1.78l-1.77-1.77l-.707.708l1.769 1.77l-1.744 1.746l.664.663l1.745-1.745l1.8 1.8z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BasketPlus(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M9.927 12.918a.648.648 0 0 1-.067-.279V8.483c0-.338.286-.576.588-.576h.021c.301 0 .594.238.594.576v3.46h.902V9.982h1.293l.164-3.006h.766c.4 0 .725-1.22.725-1.22c0-.37-.324-.672-.725-.672H12.26l-1.035-3.637a.614.614 0 0 0-.838-.286l-.072.036a.658.658 0 0 0-.288.863l.741 3.023H4.114l.768-3.019a.658.658 0 0 0-.281-.865l-.07-.036a.616.616 0 0 0-.842.277L2.616 5.083H.726c-.4 0-.725.302-.725.672c0 0 .324 1.22.725 1.22h.636l.449 6.677s.052 1.309 2.113 1.309h6.004v-2.043zM11.969 6h1.062v1.062h-1.062zM3.094 7.031H1.912V5.906h1.182zm1.912 5.711c0 .35-.251.633-.559.633h-.023c-.311 0-.561-.283-.561-.633V8.413c0-.35.25-.632.561-.632h.023c.308 0 .559.282.559.632zm2.998-.195c0 .334-.23.605-.514.605h-.021c-.284 0-.514-.271-.514-.605V8.448c0-.334.229-.604.514-.604h.021c.283 0 .514.271.514.604z"/><path d="M16 13.012h-2.008v-1.906h-.937v1.906h-2.003v.894h2.003v2.032h.937v-2.032H16z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BasketRemove(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M11 13h4.953v1H11z"/><path d="M14.244 5.084h-1.928l-1.035-3.637a.614.614 0 0 0-.838-.286l-.071.036a.657.657 0 0 0-.289.863l.741 3.023H4.17l.768-3.019a.658.658 0 0 0-.281-.865l-.07-.036a.616.616 0 0 0-.842.277L2.672 5.083H.782c-.4 0-.725.302-.725.672c0 0 .324 1.22.725 1.22h.636l.449 6.677s.052 1.309 2.113 1.309h5.989v-2.067a.65.65 0 0 1-.053-.255V8.483c0-.338.243-.611.545-.611h.021c.301 0 .546.273.546.611v3.486h2.177l.273-4.993h.766c.4 0 .725-1.22.725-1.22c0-.37-.324-.672-.725-.672M3.031 7.031H1.969V5.969h1.062zm2.031 5.711c0 .35-.251.633-.559.633H4.48c-.311 0-.561-.283-.561-.633V8.413c0-.35.25-.632.561-.632h.023c.308 0 .559.282.559.632zm2.999-.195c0 .334-.244.605-.544.605h-.022c-.302 0-.545-.271-.545-.605V8.448c0-.334.243-.604.545-.604h.022c.3 0 .544.271.544.604zm4.97-5.485h-1.094V5.968h1.094z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Basketball(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.831 2.645A7.465 7.465 0 0 0 .107 8.739c1.667-.379 3.312-.117 4.527.832l2.06-2.061zm4.376 12.238a7.465 7.465 0 0 0 6.149-1.714L7.447 8.261l-2.061 2.061c.955 1.222 1.214 2.881.821 4.561M.354 9.804a7.436 7.436 0 0 0 1.46 2.587l2.022-2.021c-.938-.681-2.193-.855-3.482-.566m4.792 4.824c.298-1.298.124-2.563-.56-3.508l-2.025 2.025a7.458 7.458 0 0 0 2.585 1.483m8.002-2.257a7.465 7.465 0 0 0 1.668-6.133c-1.663.376-3.304.113-4.514-.834l-2.06 2.061zM8.707.175a7.456 7.456 0 0 0-6.078 1.677L7.49 6.714l2.061-2.061c-.94-1.202-1.207-2.828-.844-4.478m5.846 5.001a7.411 7.411 0 0 0-1.471-2.551l-1.98 1.98c.931.674 2.173.85 3.451.571M9.775.428c-.273 1.269-.094 2.501.574 3.426l1.979-1.978A7.416 7.416 0 0 0 9.775.428"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryCharging(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M11.56 1.003H9.907V.031H7v.972H5.4c-.766 0-1.384.52-1.384 1.163v12.633c0 .643.618 1.164 1.384 1.164h6.159c.762 0 1.382-.521 1.382-1.164V2.166c0-.644-.62-1.163-1.381-1.163m.471 13.467c0 .343-.324.62-.72.62H5.632c-.395 0-.719-.277-.719-.62V2.529c0-.343.324-.621.719-.621h5.679c.396 0 .72.278.72.621z"/><path d="m10 3.006l-3.959 7.683L8.005 9.42l-.989 4.584l3.952-7.821L8.881 7.5z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryEmpty(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.306 1h-1.307V.016H8V1H6.77C5.834 1 5 1.721 5 2.646v11.679C5 15.249 5.834 16 6.77 16h5.536c.936 0 1.694-.751 1.694-1.675V2.646C14 1.721 13.241 1 12.306 1M13 14c0 .516-.484 1-1 1H7c-.515 0-1.011-.484-1.011-1V3c0-.514.496-1.051 1.011-1.051h5c.516 0 1 .536 1 1.051z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryFull(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M10.446 3H6.534C6.279 3 6 3.301 6 3.58v9.902c0 .273.278.497.535.497h3.912c.254 0 .553-.224.553-.497V3.58c0-.279-.297-.58-.554-.58"/><path d="M11.306 1.021H9.999V.005H7v1.016H5.77c-.936 0-1.694.766-1.694 1.709v11.562c0 .942.759 1.709 1.694 1.709h5.536c.936 0 1.694-.767 1.694-1.709V2.73c0-.944-.759-1.709-1.694-1.709M12 14c0 .525-.494 1-1 1H6c-.506 0-1.011-.475-1.011-1V3c0-.523.505-1 1.011-1h5c.506 0 1 .476 1 1z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryHalf(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M7 13.448c0 .281.272.511.529.511h3.912c.255 0 .559-.229.559-.511V5.069L7 8.281z"/><path d="M12.252 1H11V.016H8.023V1H6.715c-.936 0-1.694.575-1.694 1.5v11.825c0 .924.759 1.675 1.694 1.675h5.536c.936 0 1.694-.751 1.694-1.675V2.5c.001-.925-.757-1.5-1.693-1.5M13 14c0 .516-.484 1-1 1H7c-.515 0-1-.484-1-1V3c0-.514.485-1.051 1-1.051h5c.516 0 1 .536 1 1.051z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryHalfTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M7.062 13.472a.48.48 0 0 0 .471.487h3.936a.48.48 0 0 0 .47-.487V8.094l-4.876 2.065v3.313z"/><path d="M12.252 1H11V0H8.023v1H6.715c-.936 0-1.694.492-1.694 1.417v12.125c0 .924.759 1.458 1.694 1.458h5.536c.936 0 1.748-.535 1.748-1.458V2.417C14 1.492 13.188 1 12.252 1M13 14c0 .516-.484 1-1 1H7c-.515 0-1-.484-1-1V3c0-.514.485-1 1-1h5c.516 0 1 .485 1 1z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryLow(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M7.062 13.508c0 .248.21.451.467.451h3.912c.255 0 .559-.203.559-.451v-2.407l-4.938 2.056z"/><path d="M12.252 1H11V.016H8.023V1H6.715c-.936 0-1.694.658-1.694 1.583v11.742c0 .924.759 1.675 1.694 1.675h5.536c.936 0 1.694-.751 1.694-1.675V2.583C13.946 1.658 13.188 1 12.252 1M13 14c0 .516-.484 1-1 1H7c-.515 0-1.064-.484-1.064-1V3c0-.514.55-1.051 1.064-1.051h5c.516 0 1 .536 1 1.051z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bed(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" transform="translate(1 5)"><path d="M15.031 2.016v1.015H1.938V.078H.009v5.893h1.929V4.938h13.093v1.033h.938V2.016z"/><ellipse cx="3.985" cy=".97" rx=".985" ry=".97"/><path d="M13.991 1.998H6.013v-.866c0-.598.643-1.083 1.434-1.083h5.109c.793 0 1.435.485 1.435 1.083z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bell(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.141 1.318c0-.711-.559-1.286-1.248-1.286S7.646.607 7.646 1.318c0 .08.009.16.022.237C5.002 2.594 3 4.772 3 10h12c.001-5.229-2.216-7.405-4.883-8.445c.015-.077.024-.157.024-.237M16 13H2l.906-2h11.906zm-7.039 3a1.959 1.959 0 0 0 1.961-1.957H7A1.96 1.96 0 0 0 8.961 16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BicycleMan(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" transform="translate(1 2)"><path d="M8.5 12.969c-.359-.002-.495-.289-.494-.653l.015-2.784L5.153 6.44a.667.667 0 0 1 .121-.925l3.562-2.933a.644.644 0 0 1 .91.121l1.544 1.805l2.289 1.097c.335.13.504.511.377.852a.647.647 0 0 1-.839.383l-2.461-1.163a.65.65 0 0 1-.285-.215l-1.26-1.429L6.687 6.16l2.191 2.538c.086.116.09.214.09.36l.017 3.254c-.001.364-.126.658-.485.657"/><path d="M12.98 7.982c-1.645 0-2.977 1.344-2.977 3.001c0 1.657 1.332 3.001 2.977 3.001c1.645 0 2.978-1.344 2.978-3.001c0-1.657-1.333-3.001-2.978-3.001m-.005 5.064a2.066 2.066 0 0 1-2.072-2.062c0-1.139.928-2.062 2.072-2.062c1.141 0 2.069.923 2.069 2.062a2.067 2.067 0 0 1-2.069 2.062M2.991 8.023a2.97 2.97 0 1 0 0 5.942a2.97 2.97 0 1 0 0-5.942m0 5.04a2.068 2.068 0 1 1 0-4.138a2.07 2.07 0 0 1 0 4.138"/><ellipse cx="11.469" cy="1.486" rx="1.469" ry="1.486"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BicycleOne(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M2.998 8.023a2.987 2.987 0 0 0-2.982 2.982a2.987 2.987 0 0 0 2.982 2.983a2.986 2.986 0 0 0 2.983-2.983a2.986 2.986 0 0 0-2.983-2.982m0 5.024a2.043 2.043 0 0 1-2.04-2.041c0-1.125.915-2.04 2.04-2.04c1.126 0 2.041.915 2.041 2.04a2.043 2.043 0 0 1-2.041 2.041m9.987-4.982a2.946 2.946 0 0 0-2.935 2.952a2.945 2.945 0 0 0 2.935 2.951c1.616 0 2.932-1.324 2.932-2.951s-1.315-2.952-2.932-2.952m-.036 4.965a2.03 2.03 0 0 1-2.029-2.024a2.03 2.03 0 0 1 4.058 0a2.027 2.027 0 0 1-2.029 2.024"/><path d="M6.712 9.944a.535.535 0 0 1-.535-.534l-.773-3.504h-.799c-.328 0-.585-.382-.585-.868c0-.494.252-.867.585-.867c2.442.324 2.364.726 2.364.86c0 .495-.174.875-.508.875h-.03l.683 2.707l4.465-2.158l-.293-.906a.533.533 0 0 1 .12-.53l1.115-2.061h-1.916a.534.534 0 0 1-.534-.533c0-.295.239-.425.534-.425h2.85c.214 0 .407.018.491.213a.533.533 0 0 1-.102.577l-1.455 2.696l.564 2.092a.537.537 0 0 1-1.021.325l-.128-.333c.001 0-5.038 2.374-5.092 2.374"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BicycleThree(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.84.042c0 1.098-1.744 1.991-2.826 1.991S5.339 1.139 5.339.042H1.1v.916h3.111c.602.915 1.689 2.008 2.859 2.266v5.443s.392-.767.954-.767c.396 0 .812.317.94.75V3.221c1.191-.264 2.45-1.323 3.04-2.264h2.912V.041zm-2.816 10c-.529 0-.958.327-.958.731v4.537c0 .404.429.731.958.731c.529 0 .958-.327.958-.731v-4.537c0-.404-.428-.731-.958-.731m-2.571.959H3.581c-.29 0-.522.215-.522.479c0 .265.232.479.522.479h1.872c.288 0 .522-.215.522-.479c.001-.264-.234-.479-.522-.479m7.041 2.041h-1.872c-.288 0-.522.206-.522.458c0 .253.234.459.522.459h1.872c.289 0 .521-.206.521-.459c.001-.252-.232-.458-.521-.458"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bikini(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.942 10.074C10 10.074 6 10.07 4.138 10.07c-.354 0-.03.467.104.503c1.645.44 3.027 2.347 3.734 4.97c.148.551.971.592 1.114.053c.702-2.643 1.957-4.383 3.73-5.018c.192-.07.476-.504.122-.504m-.858-6.442C11.606 2.184 10.565.083 8.528.083c-2.052 0-3.095 2.132-3.566 3.581c-.535.572-.908 2.064-.908 2.988c0 2.168 2.693 1.038 3.52.101a4.137 4.137 0 0 0 1.917.017c.844.926 3.502 2.011 3.502-.117c-.001-.945-.364-2.476-.909-3.021M7.861 5.926c-.05-.935-1.214-2.094-1.968-2.469C6.324 2.281 7.132.833 8.528.833c1.396 0 2.206 1.458 2.638 2.636c-.766.394-1.944 1.552-1.98 2.475a3.331 3.331 0 0 1-.619.063c-.241 0-.477-.03-.706-.081"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BilliardBall(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M9.01.057a7.94 7.94 0 1 0 .002 15.88A7.94 7.94 0 0 0 9.01.057m-1 12.057c-2.804 0-5.076-2.283-5.076-5.099c0-2.814 2.272-5.098 5.076-5.098c2.805 0 5.078 2.283 5.078 5.098c0 2.816-2.274 5.099-5.078 5.099"/><path d="M8 4h.938v5.938H8z"/><path d="M7 4h1.969v.938H7z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Binocular(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M0 14h6v.916H0zm10 0h6v.916h-6zM2.041 3.012L.666 12.917h4.667l-.687-3.905h1.323v-6h-.985V1.979h.981v-.95H2.006v.95H3v1.033zm8.959 0h-.969v6h1.344l-.709 3.905h4.667l-1.375-9.905h-.974V1.979h.981v-.95h-3.959v.95H11z"/><path d="M8.984 2.016H7.016V3H5.969v1.969h1.066v3.062H5.969v1.94h4.062v-1.94H8.994V4.969h1.037V3H8.984z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BirthdayCake(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M2.002 13.062v1.107c0 1.305-.107 1.775 1.482 1.775h8.91c1.588 0 1.48-.471 1.48-1.775v-1.107zM4 3h.935v3.072H4zm3 0h1v2.969H7zm4 0h.99v2.851H11zM4.965.975c.1.406-.058.873-.437.915c-.354.04-.625-.417-.353-.87c.279-.46.353-.96.353-.96s.32.427.437.915m2.973.046c.105.426-.061.918-.469.962c-.381.042-.672-.439-.378-.914c.3-.485.378-1.01.378-1.01s.342.45.469.962m4.036.002c.112.426-.066.916-.498.96c-.404.042-.714-.438-.401-.913a2.847 2.847 0 0 0 .401-1.008s.363.45.498.961"/><path d="M3.146 8.437c.415 0 1.016-.394 1.193-.532c.005-.006.727-.684 1.602-.684c.871 0 1.623.672 1.654.701c.162.14.698.515 1.174.515c.562 0 1.108-.505 1.108-.505c.04-.038.787-.711 1.688-.711c.905 0 1.626.678 1.656.707c.161.138.464.341.737.442c-.031-1.301-.629-2.34-2.176-2.34H4.259c-1.391 0-2.018.837-2.154 1.949c.225.167.688.458 1.041.458m9.571 1.032c-.018-.016-.557-.519-1.151-.519c-.599 0-1.165.507-1.172.513c-.026.023-.751.7-1.624.7c-.865 0-1.641-.667-1.674-.696c-.163-.146-.682-.517-1.154-.517c-.562 0-1.104.503-1.111.507c-.115.092-.932.706-1.684.706c-.386 0-.772-.159-1.07-.325v2.107h11.889v-1.814c-.611-.127-1.177-.601-1.249-.662M0 14h15.851v1.935H0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Biscuit(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.8 6.826a.862.862 0 0 0-.217-.307l-.024-.009a1.127 1.127 0 0 0-.418-.239L9.352 4.126c.379-.436.617-.997.617-1.619a2.476 2.476 0 1 0-4.954 0c0 .633.244 1.204.635 1.642L.863 6.312c-.642.202-.984.936-.762 1.637l.025.08c.222.701.922 1.105 1.564.902l1.982-.964c.491-.045.898-.032 1.015.079c.312.295-1.819 4.369-2.424 5.561c-.438.859-.235 1.84.452 2.188l.075.039c.688.349 1.558-.088 2.036-.925c.344-.603 1.518-3.582 2.576-3.582c1.058-.001 2.476 2.977 2.782 3.582c.438.859 1.349 1.273 2.036.925l.076-.039c.688-.349.891-1.329.451-2.188c-.385-.761-2.835-5.06-2.626-5.561c.068-.164.575-.126 1.24-.021l1.948.865c.643.203 1.342-.201 1.564-.902l.025-.08a1.353 1.353 0 0 0-.098-1.082m-6.739-.847v1.052H6.907V5.979zm.001 4.083H6.906V8.906h1.156z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Blender(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m10.822 12.062l-.864.007L9.987 11H7.03l-.01 1.06l-.902-.002c-1.19.809-2.097 2.134-2.097 2.983c0 1.229 8.929 1.25 8.929 0c0-.453-.827-2.049-2.128-2.979m-1.779 2H7.958v-1.104h1.085zm2.814-12.039c.076-.512.122-.874.122-.992c0-1.312-6.917-1.291-6.917 0c0 1.311 1.682 8.901 3.031 8.901h.733c.448 0 .892-.84 1.287-2.005c3.229-.413 3.882-3.235 3.886-3.293V2.023zm-1.378 4.956c.419-1.303.769-2.87.983-4h1.558l.001 1.692c-.009.085-.238 1.891-2.542 2.308"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BloodBag(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M12.958 9.221V1.783S13.14.052 10.896.052H6.084c-2.244 0-2.062 1.731-2.062 1.731v7.438s-.146 1.748 1.062 1.748h1.959v.948h1.001V14H7.043v1h1.001v.917H9V15h.938v-1H9v-2.083h1v-.948h1.896c1.354 0 1.062-1.748 1.062-1.748m-2.308.946H10V10H7.042v.167H6.39c-1.546 0-1.478-1.356-1.478-1.356V4.954h2.057v-.901H4.912V2.954h2.041v-.901H4.965c.104-.424.423-1.136 1.424-1.136h4.261c1.35 0 1.479 1.445 1.479 1.445v6.449c.001 0 .121 1.356-1.479 1.356"/><path d="M6.023 8.137c0 .696.703.807.703.807h3.541s.704-.026.704-.807v-3.11H6.023z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bluetooth(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="m8 11.967l.969-.943L8 10zm.03-5.985l.939-.945l-.939-1.028z"/><path d="M7.506 0C5.033 0 3.027 1.546 3.027 3.453v9.095C3.027 14.454 5.033 16 7.506 16c2.472 0 4.477-1.546 4.477-3.452V3.453C11.982 1.546 9.978 0 7.506 0m2.946 11.011L6.94 14V8.815l-2.22 2.066l-.674-.748l2.09-2.07L4 6.1l.72-.666l2.22 1.912V2l3.56 3.117l-2.862 2.978z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Board(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 2H3a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2M5 10H4V7H3V6h2zm4-2H8v1H7V8H6V7h1V6h1v1h1zm3 2h-1V7h-1V6h2zm3-1h-2V8h2zm0-2h-2V6h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Boat(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m15.978 14.895l.988-1.829c-4.372 0-10.502.51-12.888-1.036C1.028 11.524 1.042 11 1.042 11s.68 3.895 3.036 3.895zm-3.961-3.905L12.014 2l4.899 8.99zm-1.002.032V-.052L6 10.253z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bolt(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.289.023L6.925 0L2.984 8H8l-4.334 7.916L14.924 4.941H10.35z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BombOne(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.004 5.016c-.814 0-1.962 1.12-1.962 2.917c0 1.796 1.071 3.029 1.976 3.029s6.561-.335 7.332-.335c.772 0 2.672-.438 2.672-2.694c-.001-2.256-1.712-2.655-2.672-2.655c-.96 0-6.532-.262-7.346-.262m-3.031.368c0 .327-.177.593-.395.593H2.394c-.218 0-.394-.266-.394-.593V3.606c0-.328.177-.593.394-.593l1.578 1.778zm.015 5.804L2.41 12.957c-.217 0-.395-.264-.395-.59v-1.769c0-.326.178-.59.395-.59h1.184c.218 0 .394.264.394.59zm-2.55-4.185c-.327 0-.406.443-.406.989c0 .547.079.988.406.988h2.531V7.002H1.438z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BombTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.427.016c-.093.952-.854 1.255-2.095 1.639c-.584.182-1.205.379-1.712.72c-.89-.473-1.511-.363-1.511-.363L9.808 3.628a6.034 6.034 0 0 1 3.104 2.271L14.2 4.296s.095-.688-.958-1.513c.378-.2.825-.344 1.287-.487c1.135-.352 2.325-.832 2.47-2.296zm-7.57 4.348a5.961 5.961 0 0 0-1.895-.312a5.964 5.964 0 1 0 4.957 2.649a5.968 5.968 0 0 0-3.062-2.337M7.299 6.077c-2.707 0-4.162 1.509-4.162 3.966H2.031c0-2.635 2.212-4.81 5.038-5.079c-.001 0 .388 1.113.23 1.113"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bone(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.321 13.61s.062 1.264.649 1.852c.684.683 1.781.694 2.452.024c.954-.954 1.084-2.273.97-3.367l5.766-5.765c1.087.109 2.404-.027 3.375-.998c.662-.662.647-1.752-.034-2.434c-.589-.588-1.786-.577-1.786-.577s-.038-1.246-.631-1.839c-.686-.688-1.789-.703-2.457-.033c-.97.97-1.085 2.322-.958 3.429L4.939 9.631c-1.091-.109-2.416.026-3.395 1.004c-.664.663-.651 1.75.025 2.427c.587.584 1.752.548 1.752.548"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Book(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M0 0h2v16H0zm11 6v3h.885L12 6z"/><path d="M3 0v16h10.82c.651 0 1.18-.453 1.18-1.01V1.01C15 .452 14.472 0 13.82 0zm10.051 9.053h-.971l-.018 1.01H7.906l.018-1.021h-.967V6.99h.958V5.948l3.042-.01v-.887H7.026v.997h-.997v3.91h1.012v1.017h4.006v1.039H6.961v-.951H5.958v-1.031H4.953V5.991h1.02V4.973h.965V3.938h4.094v1.021h.979v.99h1.041z"/><path d="M8 7h2v2H8z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookFour(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 0v16H3.016v-2H4v-2h-.984V4H4V2h-.984V0zM2 2v2h1V2zm0 10v2h1v-2zm15-1.93h-1v2.843h1zm0-4.052h-1V9h1zm0-3.917h-1V4.96h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookOne(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 3.007C4.691 1.444 2.11 2.32 0 3.61v10.325c2.105-1.298 5.248-2.364 8-.946v-9.98zm0 9.995c3.629-1.463 5.919-.353 8 .937V3.621c-2.081-1.285-4.118-2.438-8-.845z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookOpen(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M7.969 2.74C4.92 1.288 3.942 2.326 2 3.523v7.445c1.941-1.205 3.434-1.969 5.969-.651zm.058.018v7.591c3.352-1.361 4.035-.549 5.957.651V3.542c-1.922-1.194-2.537-2.268-5.957-.784"/><path d="M15.938 13H.051V3.049h.902v9.029h14.078V3.094h.907z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookPerson(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 0h.918v16H2zm12.905 0H4v16h10.905c.604 0 1.095-.453 1.095-1.01V1.01C16 .452 15.51 0 14.905 0m-4.899 2.875c1.095 0 1.985.886 1.985 1.979s-.891 3.184-1.985 3.184c-1.096 0-1.984-2.09-1.984-3.184a1.98 1.98 0 0 1 1.984-1.979M14.976 12H5.002s-.171-2.987 2.387-2.987c.537.577 1.272.931 2.625.931c1.354 0 1.999-.36 2.529-.944c2.878 0 2.433 3 2.433 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookThree(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m4.048 13.774l2.014 2.122l8.847-2.065V1.965l-.977-.581l-7.87 2.577l-1.065-1.04l7.7-2.285l-1.148-.625l-7.487 2.122z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookcase(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.667.037H3.332A1.32 1.32 0 0 0 2.016 1.36l.01 14.075c0 .376.642.58.995.58c.354 0 .938-.204.938-.58v-1.407c.009 0 .017.003.017.003h10.01c.013 0 .023-.003.036-.004v1.408c0 .376.635.602.979.602c.347 0 .958-.226.958-.602l.027-14.075A1.322 1.322 0 0 0 14.667.037m-1.688.921h1.041V3h-1.041zm-2.021 0h1v2h-1zm-1.985 0h1.048v2H8.973zm6.048 9.063H2.98V8.98h12.041zM3.958 8V6h1v2zm2.014-.007v-2h1v2zM7.961 8V5.993h1V8zm7.06-3H2.98V3.979h12.041z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookmark(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.677 0H3.326c-.732 0-1.324.605-1.324 1.353v13.294c0 .748.593 1.353 1.324 1.353l4.676-4.013L12.676 16c.732 0 1.326-.605 1.326-1.353V1.353A1.336 1.336 0 0 0 12.677 0m-2.195 10L8.01 8.633L5.538 10l.473-2.896L4.01 5.055l2.764-.422L8.01 2l1.236 2.633l2.764.422l-2 2.05z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BotlMilk(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M10.137 15.984H4.885c-1.013 0-1.837-.747-1.837-1.666V7.022c0-.42.444-1.397 1.173-2.956c.422-.898.996-2.127 1.002-2.375c0-.124-.209-.352-.35-.503c-.238-.257-.484-.524-.33-.828c.148-.294.554-.335.974-.335H9.5c.53.004.91.066 1.053.354c.148.305-.107.577-.332.818c-.139.147-.348.369-.348.493c.002.25.562 1.494.975 2.401c.701 1.545 1.126 2.514 1.126 2.93v7.296c0 .92-.826 1.667-1.837 1.667M5.543 1.041c.219.242.453.585.453.939c0 .347-.324 1.113-.965 2.506c-.449.971-1.059 2.299-1.059 2.6v7.043c0 .533.475.966 1.055.966H9.99c.584 0 1.058-.433 1.058-.966V7.086c0-.301-.587-1.619-1.018-2.58c-.604-1.361-1.011-2.181-1.011-2.531c0-.358.309-.689.531-.93c.016-.019-.076-.068-.167-.068H5.627c-.053 0-.101.044-.084.064"/><path d="M9.963 13.092c0 .701-.521.887-1.164.887H6.235c-.643 0-1.163-.186-1.163-.887V6.538c0-.701.63 1.534 1.985.474c1.356-1.056 2.906-1.175 2.906-.474z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BotlTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 0h1.969v2.03H8zm1.886 3H8.06s.497 1.247-1.112 2.006C5.691 5.599 6 5.259 6 7.203V14.9c0 .608.521 1.1 1.16 1.1h3.679c.64 0 1.161-.492 1.161-1.1V7.203c-.001-2.029.187-1.436-.97-2.166C9.505 4.08 9.886 3 9.886 3M11 15h-1V7h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BowTie(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M.953 10.974c1.234 0 5.01-1.313 5.01-2.613v-.705c0-1.354-4.383-2.655-5.01-2.655C-.278 5-.26 10.974.953 10.974m14.103-.047c-1.237 0-5.022-1.303-5.022-2.593v-.699C10.034 6.29 14.429 5 15.056 5c1.235 0 1.215 5.927 0 5.927M7 7h1.984v1.984H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Box(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 2h16v2H1zm1 10a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V5H2zm4.98-5.041h4.082v1.104H6.98z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxDownload(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M16.644 4.236a1.022 1.022 0 0 0-1.433.113l-2.062 2.754c-.104-.016-8.143-.016-8.225-.006l-2.159-2.76a1.005 1.005 0 0 0-1.422-.092a1.02 1.02 0 0 0-.091 1.434l2.027 2.607a2.082 2.082 0 0 0-.196.873v4.746a2.09 2.09 0 0 0 2.089 2.088h7.661a2.09 2.09 0 0 0 2.089-2.088V9.159c0-.289-.059-.562-.165-.814l2-2.68a1.013 1.013 0 0 0-.113-1.429"/><path d="m11.781 2.057l-1.812.373L9.99.994a.978.978 0 0 0-.988-.988C8.45.008 8.002.457 8 1.008l.031 1.393l-1.781-.344a.698.698 0 0 0-.003.986l2.265 2.676c.271.271.712.27.985-.004l2.278-2.672a.698.698 0 0 0 .006-.986"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoxUpload(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M16.644 4.236a1.022 1.022 0 0 0-1.433.113l-2.062 2.754c-.104-.016-8.143-.016-8.225-.006l-2.159-2.76a1.005 1.005 0 0 0-1.422-.092a1.02 1.02 0 0 0-.091 1.434l2.027 2.607a2.082 2.082 0 0 0-.196.873v4.746a2.09 2.09 0 0 0 2.089 2.088h7.661a2.09 2.09 0 0 0 2.089-2.088V9.159c0-.289-.059-.562-.165-.814l2-2.68a1.013 1.013 0 0 0-.113-1.429"/><path d="m6.236 3.944l1.812-.373l-.021 1.436a.978.978 0 0 0 .989.988c.55-.002.999-.451 1.001-1.002L9.986 3.6l1.781.344a.697.697 0 0 0 .003-.986L9.505.282a.696.696 0 0 0-.985.004L6.242 2.958a.699.699 0 0 0-.006.986"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bread(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.942 9.189a8.683 8.683 0 0 1 1.173-.666a19.145 19.145 0 0 0-1.561-1.973a9.09 9.09 0 0 0-2.564.557c-1.673.643-2.354-.227-.771-1.311a8.358 8.358 0 0 1 1.685-.879a20.943 20.943 0 0 0-2.243-1.752a9.058 9.058 0 0 0-2.395.548c-1.671.642-2.354-.228-.769-1.311c.291-.199.598-.371.911-.532C3.859.668 1.6.515.523 1.697c-1.719 1.884 1.365 5.196 5.479 8.944c4.111 3.748 7.691 6.514 9.412 4.629c.973-1.068.767-3.001-.339-5.158c-.467.1-.925.219-1.362.388c-1.674.643-2.356-.226-.771-1.311"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Briefcase(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M6.959 3.079c0-.619.482-1.124 1.072-1.124h1.838c.59 0 1.184.505 1.184 1.124v-.058h.919v-.474c0-.86-.478-1.548-1.458-1.548H7.385c-.654 0-1.362.73-1.362 1.548v.474h.937v.058z"/><path d="M11 5.969v1.012h6V3.016H.985v3.965h6.003V5.969zM7 8.998v-.995H1.022V16H17V8.003h-5.987v.995H7.001zM8 7v.967h1.99V7z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BriefcasePerson(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M13.562 3.083H1.437c-.794 0-1.438.651-1.438 1.453V13.5c0 .804.644 1.454 1.438 1.454h12.125c.794 0 1.438-.65 1.438-1.454V4.536c0-.802-.644-1.453-1.438-1.453M7.504 5.906c.855 0 1.551.706 1.551 1.577c0 .872-.695 2.538-1.551 2.538S5.953 8.355 5.953 7.483c0-.871.695-1.577 1.551-1.577m3.5 7.125H4s-.119-3.076 1.677-3.076c.375.594.892 1.164 1.843 1.164c.949 0 1.401-.574 1.775-1.178c2.021 0 1.709 3.09 1.709 3.09"/><path d="M5.782 2.682c0-.428.295-.775.66-.775h2.093c.364 0 .659.348.659.775v.885h.781V2.376c0-.742-.514-1.347-1.144-1.347H6.15c-.632 0-1.147.604-1.147 1.347v1.191h.78v-.885z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Brightness(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.041 3.001C6.326 3.001 4.125 5.239 4.125 8c0 2.762 2.201 5 4.916 5c2.717 0 4.917-2.238 4.917-5c0-2.761-2.2-4.999-4.917-4.999m.042 8.226a3.227 3.227 0 1 1 0-6.453a3.227 3.227 0 0 1 0 6.453M8 0h1.958v2H8zm0 14h2v1.958H8zm7-7h1.958v1.958H15zM1 7h1.958v2H1zm13.691-5.207l1.326 1.327l-1.326 1.324l-1.325-1.328zM3.468 12.714l1.325 1.327l-1.367 1.366L2.1 14.08zm11.082.064l1.364 1.363l-1.326 1.326l-1.364-1.364zM3.443 1.775l1.301 1.302l-1.326 1.325l-1.3-1.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrushAndPencil(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M15.833 13.042c-.517.179-1.538.55-2.712-1.436c-1.174-1.984-2.726-1.523-2.726-1.523c-.087.015-.161.044-.238.07c-1.752-2.495-6.679-7.929-7.113-8.364C2.421 1.167 1.375.583.5 1.5c-.823.862-.401 1.987.222 2.61c.438.438 5.902 5.39 8.396 7.136a1.226 1.226 0 0 0-.076.337c0 3.646 3.729 3.453 4.5 3.312c1.375-.249 2.808-2.03 2.291-1.853"/><path d="m5.613 10.73l-.227.752l.486.486l-3.811 1.656l-.75-.719l1.721-3.627l-.018.706l.998-.014l.853-.845c-.53-.45-1.058-.906-1.562-1.347l-.834.835S.157 13.933 0 14.363c-.125.34.306.771.604.625c.48-.237 5.74-2.52 5.74-2.52l1.168-1.172a58.67 58.67 0 0 1-1.361-1.094zm6.534-8.81l1.259 1.236l-4.23 4.107c.378.46.734.902 1.051 1.311l4.261-4.275c.502-.502.53-1.226-.008-1.763L12.475.546c-.539-.537-1.383-.565-1.884-.064L6.71 4.37c.446.51.909 1.045 1.364 1.581z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrushOne(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m1.056 8.306l4.308-3.351l6.703 6.705l-3.35 4.308s.044-1.054-1.437-1.437c-1.39-.359-1.632-2.288-2.843-2.688c-1.211-.4-1.426-1.242-1.986-2.145S1.056 8.306 1.056 8.306m12.925-.476l-1.146-1.146c1.841-1.603 3.279-2.902 3.537-3.162A2.032 2.032 0 0 0 13.499.649c-.258.26-1.559 1.697-3.16 3.54L9.193 3.043a1.355 1.355 0 0 0-1.916-.001l-.957.957l6.704 6.705l.958-.959a1.355 1.355 0 0 0-.001-1.915m.265-5.053a.977.977 0 1 1 1.383-1.383a.976.976 0 0 1 0 1.383a.976.976 0 0 1-1.383 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrushTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.172 4.244c-.457.541-4.619 4.77-6.846 6.422l-.913-.867c1.29-2.511 4.741-7.292 5.196-7.832c.595-.704 2.037-1.215 2.807-.669c1.012.716.35 2.24-.244 2.946m-8.565 6.458s-2.092-.553-3.348 1.594C2.004 14.443.808 13.708.27 13.524c-.538-.184 1.256 2.27 2.093 2.27c.818 0 4.389 1.175 5.268-3.434c.018-.101.072-1.483-1.024-1.658"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BubbleChat(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="m2.701 14.927l1.051-3.998H2.299C1.059 10.929 0 9.944 0 8.672V2.303C0 1.033 1.011 0 2.252 0h11.496C14.989 0 16 1.033 16 2.303v6.369c0 1.271-.963 2.257-2.203 2.257H7.983zM.959 2.513v6.052c0 .88.703 1.452 1.566 1.452h2.297l-.77 2.869l3.482-2.87h5.98c.865 0 1.568-.571 1.568-1.451V2.513c0-.879-.703-1.595-1.568-1.595H2.525c-.863 0-1.566.716-1.566 1.595"/><path d="M12.625 8.946H3.344c-.996 0-1.295-.419-1.295-1.384V3.468c0-.963.361-1.426 1.357-1.426h9.219c.998 0 1.314.557 1.314 1.519v3.97c-.001.965-.316 1.415-1.314 1.415"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BubbleMessage(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.988 1.002c-4.387 0-7.947 2.619-7.947 5.852c0 2.95 2.968 5.383 6.824 5.786l-2.042 3.332l6.103-3.686c2.934-.86 5.012-2.967 5.012-5.433c0-3.232-3.559-5.851-7.95-5.851"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BubbleMessageDot(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.1 0H3.902C2.873 0 2.039.795 2.039 2.094v5.844c0 1.299.834 2.051 1.863 2.051h5.7l-1.57 4.98l4.474-4.98h1.595c1.029 0 1.864-.752 1.864-2.051V2.094C15.964.795 15.129 0 14.1 0M5.02 6.312a1.31 1.31 0 0 1-1.301-1.318c0-.728.582-1.317 1.301-1.317c.721 0 1.301.59 1.301 1.317A1.309 1.309 0 0 1 5.02 6.312M9 6.24c-.709 0-1.281-.558-1.281-1.246S8.291 3.749 9 3.749c.707 0 1.281.558 1.281 1.245c0 .689-.574 1.246-1.281 1.246m4 0c-.705 0-1.281-.558-1.281-1.246S12.295 3.749 13 3.749c.707 0 1.281.558 1.281 1.245c0 .689-.574 1.246-1.281 1.246"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BubbleMessageDotTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.019 1.04c-4.398 0-7.968 2.62-7.968 5.852c0 2.95 2.975 5.384 6.842 5.787l-2.048 3.25l6.119-3.603c2.942-.861 5.025-2.968 5.025-5.435c0-3.231-3.569-5.851-7.97-5.851M6 8H4V6h2zm4 0H8V6h2zm4 0h-2V6h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BubbleMessageHi(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.979.021c-4.398 0-7.968 2.792-7.968 6.235c0 3.143 2.976 5.734 6.841 6.164l-2.047 3.551l6.119-3.926c2.94-.917 5.023-3.161 5.023-5.789c0-3.444-3.568-6.235-7.968-6.235m.061 8.997H7.972V7.075h-1.94v1.943l-1.072.023V3.992l1.072-.023v2.018h1.94V3.969H9.04zm1.96 0H9.954v-3.08H11zm.031-3.977H9.954V3.937h1.077zm3 3.975h-1.063V7.963h1.063zm-.015-2h-1.047V3.954h1.047z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BubbleMessageQuote(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.009.042C3.615.042.052 2.826.052 6.263c0 3.136 2.972 5.722 6.832 6.151l-2.116 3.57l6.183-3.945c2.937-.914 5.018-3.154 5.018-5.776c-.001-3.437-3.566-6.221-7.96-6.221M7.051 7c0 1.333-1.655 2.062-3.084 2.062V7.945c.721 0 1.741-.354 1.741-.945H3.967V3.989h3.084zm5 0c0 1.333-1.655 2.062-3.084 2.062V7.945c.721 0 1.741-.354 1.741-.945H8.967V3.989h3.084z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BubbleMessageTalk(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.939 1C2.666 1 .009 2.987.009 5.438c0 2.236 2.215 4.082 5.092 4.387L3.88 12.26l4.249-2.7C10.318 8.906 12 7.309 12 5.438C12 2.988 9.213 1 5.939 1m10.008 8.89c0-1.124-1.062-2.288-2.289-2.868c-.344 1.95-1.924 3.745-4.417 4.447l-1.187.642c.454.34 1.01.611 1.634.788l3.638 1.971l-1.303-1.776c2.217-.225 3.924-1.571 3.924-3.204"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BubbleMessageText(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.01.078C3.606.078.031 2.759.031 6.065c0 3.018 2.98 5.508 6.852 5.92l-2.121 3.438l6.197-3.799c2.945-.88 5.031-3.035 5.031-5.559c0-3.306-3.574-5.987-7.98-5.987M10 9H4V8h6zm2-1.954H4V6h8zm0-2.055H4v-1h8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bucket(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" fill-rule="evenodd"><path d="M8.759 1.143c-.507-.509-2.356.521-4.13 2.302c-1.772 1.782-2.801 3.637-2.293 4.146c.506.508 2.354-.521 4.129-2.304C8.237 3.506 9.266 1.65 8.759 1.143"/><path fill="currentColor" d="m16.737 7.881l-.903-.896c-.327.91-.975 1.689-1.929 2.25c-1.017.6-2.293.916-3.691.916a9.576 9.576 0 0 1-2.378-.305a.593.593 0 1 1 .297-1.149c.695.18 1.398.271 2.091.271c1.175 0 2.239-.261 3.077-.753c.793-.468 1.314-1.111 1.507-1.861c.028-.114.023-.233.038-.35L8.862.062S7.173-.244 4.044 2.899c-3.133 3.145-3.01 4.935-3.01 4.935s7.16 7.189 7.877 7.907c.717.72 2.991-.208 5.339-2.564c2.348-2.363 3.093-4.688 2.487-5.296M2.336 7.59c-.508-.509.521-2.363 2.293-4.146C6.402 1.663 8.252.633 8.759 1.142c.507.508-.521 2.363-2.294 4.144C4.69 7.068 2.842 8.098 2.336 7.59"/><path fill="currentColor" d="M15.864 6.621c.69-2.685-1.775-5.647-5.498-6.605a.498.498 0 0 0-.604.355a.496.496 0 0 0 .356.602c3.063.787 5.157 3.049 4.833 5.136l.806.799c.034-.097.082-.187.107-.287"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bug(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16 9.984V9h-1.031L15 6.016h-2.252a5.35 5.35 0 0 0-1.281-1.683L11.028 4h1.957V1.985h.984v-.969h-1.953v.016H12v2.016h-1.406c-.067-.443-.303-.885-.609-1.215V0H9v1.224a2.795 2.795 0 0 0-1-.203c-.365 0-.7.087-1.016.209V0H6v1.854A2.212 2.212 0 0 0 5.411 3H3.984V1.031h-.016v-.016H2.015v.969H3v1.984h2.013l-.48.364a5.35 5.35 0 0 0-1.281 1.683H1.027V9H0v.984h1.984V7h.896a5.28 5.28 0 0 0-.173 1.321c0 .979.269 1.894.729 2.679l-2.409.026v3.016H.055V15H2v-3.031h2.194a5.269 5.269 0 0 0 2.793 1.554L8 5l1.029 8.519c1.08-.214 2.025-.771 2.762-1.534H14V15h1.969v-.984h-1V11h-2.406a5.262 5.262 0 0 0 .729-2.679c0-.457-.063-.945-.172-1.368h.896v3.031zm-10 .975H5v-1h1zm-1-3H4v-1h1zm7 0h-1v-1h1zm-2-3H9v-1h1zm-3.062 0h-1v-1h1zm4.062 6h-1v-1h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Building(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.961 3.031h-.992v-.997h-1.016V.031h-.938v2.003h-.957v.997h-1.023v2H9V16h6.982V5.031h-1.021zm-1.992.938h1v1h-1zm-2 .015h1v1h-1zM11 13.956h-1v-1h1zm0-1.999h-1v-1h1zm0-1.999h-1v-1h1zM11 8h-1V7h1zm2 5.956h-1v-1h1zm0-1.999h-1v-1h1zm0-1.999h-1v-1h1zM13 8h-1V7h1zm2 5.956h-1v-1h1zm0-1.999h-1v-1h1zm0-1.999h-1v-1h1zm0-2.009h-1v-1h1zm-8.039-4.9H5.924v-.997h-.949V.047h-.928v2.005H3.014v.997h-.979v1.998H1V16h6.982V5.047H6.961zm-2.004.912h1v1h-1zM3 3.961h1v1H3zm0 10.032H2v-1h1zm0-2.004H2v-1h1zm0-2.018H2v-1h1zM3 8H2V7h1zm2 5.969H4v-1h1zm0-1.995H4v-1h1zm0-1.993H4v-1h1zm0-2.006H4v-1h1zm2 5.984H6v-1h1zm0-1.998H6v-1h1zm0-1.998H6v-1h1zm0-2.008H6v-1h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BulletCheckedList(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.016 9h-3V6h3zM.953 8h1V6.969h-1zm2.016 4.986H.016V10h2.953zM.954 12h1v-1h-1zM5 3h10.977v.976H5zm0 4h10.977v.96H5zm0 4h10.977v.914H5zM1.366 5.295L.021 3.939l.635-.635l.71.696l2.036-2l.623.636z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BulletList(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.016 9h-3V6h3zM.953 8h1V6.969h-1zm2.016 4.986H.016V10h2.953zM.954 12h1v-1h-1zm2.062-7h-3V2h3zM.953 4h1V2.969h-1zM5 7h10.977v.96H5zm0-4h10.977v.96H5zm0 8h10.977v.914H5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BulletListTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 1h2v2H1zm0 5h2v2H1zm0 5h2v2H1zm5-9h11v1H6zm0 5h11v1H6zm0 5h11v1H6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bus(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" transform="translate(1 3)"><ellipse cx="3.492" cy="8.48" rx="1.492" ry="1.48"/><circle cx="12.485" cy="8.485" r="1.485"/><path d="M14.078.024h-7.91c-2.104 0-3.137.5-3.137 2.991c0 0-.888.993-1.586.993c-.698 0-1.445.443-1.445.985v3.131c0 .408.219.728.918.819a2.667 2.667 0 0 1-.049-.485c0-1.438 1.182-2.605 2.637-2.605S6.143 7.02 6.143 8.458c0 .176-.019.346-.053.512l3.822.002a2.669 2.669 0 0 1-.048-.477a2.644 2.644 0 0 1 2.642-2.643a2.643 2.643 0 0 1 2.641 2.643a2.6 2.6 0 0 1-.049.478h.887V3.009c-.001-1.863-.671-2.985-1.907-2.985M7.016 3.031H3.918c0-2.223.947-2.05 1.293-2.05L7.023.967zm5.027-.01H7.969V.978h4.074zm.931 0V.978h1.435c.692 0 .67 2.043.67 2.043z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ButtonArrowDown(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M7.984 16C3.595 16 .025 12.424.025 8.027c0-4.395 3.57-7.971 7.959-7.971c4.389 0 7.959 3.576 7.959 7.971c0 4.397-3.57 7.973-7.959 7.973M7.977 1.904c-3.363 0-6.102 2.732-6.102 6.086c0 3.357 2.739 6.09 6.102 6.09c3.365 0 6.102-2.733 6.102-6.09c-.001-3.354-2.738-6.086-6.102-6.086"/><path d="m5.047 8.051l2.984 3.904l2.916-3.905H8.944V4.786c0-.344-.324-.79-.955-.79c-.63 0-.976.483-.976.826v3.229z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ButtonArrowLeft(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M1 8.041C1 3.652 4.582.082 8.985.082c4.401 0 7.983 3.57 7.983 7.959c0 4.389-3.582 7.959-7.983 7.959C4.582 16 1 12.43 1 8.041m14.057 0c0-3.333-2.715-6.045-6.051-6.045c-3.337 0-6.052 2.712-6.052 6.045c0 3.333 2.716 6.045 6.052 6.045c3.337 0 6.051-2.712 6.051-6.045"/><path d="M8.975 5.02L5.071 8.022l3.905 2.951V8.97h3.14c.345 0 .791-.324.791-.955c0-.63-.483-.975-.826-.975H8.976V5.02z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ButtonArrowRight(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M16.969 8.041c0 4.389-3.582 7.959-7.986 7.959c-4.4 0-7.982-3.57-7.982-7.959c0-4.389 3.582-7.959 7.982-7.959c4.404 0 7.986 3.57 7.986 7.959m-14 0c0 3.305 2.712 5.996 6.045 5.996c3.333 0 6.047-2.691 6.047-5.996c0-3.305-2.713-5.996-6.047-5.996c-3.334 0-6.045 2.691-6.045 5.996"/><path d="m9.057 10.969l3.904-3.002l-3.905-2.951v2.003H5.793c-.346 0-.791.324-.791.955c0 .63.483.975.826.975h3.229z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ButtonArrowUp(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.984.057c4.389 0 7.959 3.562 7.959 7.941c0 4.377-3.57 7.939-7.959 7.939c-4.389 0-7.959-3.562-7.959-7.939c0-4.379 3.571-7.941 7.959-7.941m-.015 14.026c3.347 0 6.074-2.729 6.074-6.083c0-3.353-2.727-6.083-6.074-6.083c-3.349 0-6.073 2.73-6.073 6.083c0 3.354 2.724 6.083 6.073 6.083"/><path d="M11.975 7.906L8.973 4.002L6.022 7.907h2.003v3.264c0 .344.324.79.955.79c.63 0 .975-.483.975-.826V7.906z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ButtonBuy(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="m13.36 2.094l-5.36 1l-5.318-1C1.231 2.094.053 3.412.053 5.039v4.982c0 1.626 1.178 2.943 2.629 2.943L8 11.937l5.36 1.027c1.453 0 2.63-1.317 2.63-2.943V5.039c0-1.627-1.177-2.945-2.63-2.945M6 9.041h-.968V10H2V5h3.032v1H6zM10 10H7V5h1v4.025h1V5h1zm4-1.975L13.04 8v2h-1.08l.02-1.975H11V5h.98v3.025h1.06V5H14z"/><path d="M3 8v1.037h2.006V8zm0-2h2.026v1.029H3z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ButtonError(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M14.646 2.371c-3.111-3.11-8.177-3.111-11.288 0c-3.111 3.112-3.11 8.177 0 11.289c3.111 3.11 8.176 3.111 11.288 0c3.112-3.113 3.111-8.177 0-11.289M4.587 12.431C2.148 9.993 2.146 6.028 4.58 3.594c2.434-2.435 6.399-2.432 8.838.006c2.438 2.438 2.439 6.404.006 8.838c-2.436 2.434-6.4 2.431-8.837-.007"/><path d="M11.164 11.063a.557.557 0 0 1-.388-.141L9.009 9.157l-1.695 1.695a.622.622 0 0 1-.423.146a.842.842 0 0 1-.603-.265c-.221-.22-.27-.438-.271-.58a.554.554 0 0 1 .14-.4l1.724-1.725l-1.68-1.678c-.152-.153-.279-.627.12-1.025c.225-.225.446-.272.593-.272c.183 0 .32.072.387.141L9.009 6.9l1.696-1.696a.624.624 0 0 1 .424-.146a.84.84 0 0 1 .604.266c.378.379.302.81.131.98L10.138 8.03l1.737 1.736c.067.068.146.22.146.424a.834.834 0 0 1-.265.602c-.225.224-.445.271-.592.271"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ButtonHd(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M14.276 2L9.015 3.062L3.723 2C2.218 2 .999 3.243.999 4.777v5.41c0 1.534 1.219 2.777 2.724 2.777l5.292-1.012l5.261 1.012c1.504 0 2.724-1.243 2.724-2.777v-5.41C17 3.243 15.78 2 14.276 2m-6.231 8.032H6.957v-1.97H5.061v1.97H3.973V4.969h1.088v1.969h1.896V4.969h1.088zm3.955.029H9.953V4.96L12 4.959c2.303 0 2.072.812 2.072 2.552c0 1.738.335 2.55-2.072 2.55"/><path d="M12.197 6.006H11v2.979h1.197c.99-.016.792-.248.792-1.488c0-1.243.199-1.477-.792-1.491"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ButtonPlay(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.309 3L9.032 4.125L3.63 3C2.178 3 1.001 4.329 1.001 5.967v5.021c0 1.638 1.177 2.967 2.629 2.967l5.402-1.018l5.277 1.018c1.452 0 2.629-1.329 2.629-2.967V5.967C16.938 4.329 15.761 3 14.309 3m-3.48 6.117l-2.505 1.917a1.167 1.167 0 0 1-1.418-.004l.029-5.104c.419-.319.945-.222 1.368.101l2.521 1.93c.422.321.424.841.005 1.16"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ButtonPlus(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M9.021.114c-4.394 0-7.969 3.554-7.969 7.922s3.575 7.922 7.969 7.922c4.395 0 7.969-3.554 7.969-7.922S13.415.114 9.021.114M9.013 14.07c-3.329 0-6.04-2.724-6.04-6.071c0-3.347 2.711-6.071 6.04-6.071s6.039 2.725 6.039 6.071c0 3.348-2.71 6.071-6.039 6.071"/><path d="M8.996 11.969c-.654 0-.959-.437-.959-.731V8.983H5.813c-.265 0-.782-.298-.782-.986c0-.654.438-.959.732-.959h2.271V4.801c0-.264.298-.782.986-.782c.654 0 .959.438.959.732v2.285h2.236c.265 0 .783.299.783.987c0 .653-.438.958-.732.958H9.981v2.206c.001.263-.296.782-.985.782"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ButtonRemove(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.021.114C3.627.114.052 3.668.052 8.036s3.575 7.922 7.969 7.922c4.395 0 7.969-3.554 7.969-7.922S12.415.114 8.021.114M8.013 14.07c-3.329 0-6.04-2.724-6.04-6.071c0-3.347 2.711-6.071 6.04-6.071s6.039 2.725 6.039 6.071c0 3.348-2.71 6.071-6.039 6.071"/><path d="M4.813 8.982c-.265 0-.782-.298-.782-.986c0-.654.438-.959.732-.959l6.453-.002c.265 0 .783.299.783.987c0 .653-.438.958-.732.958z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ButtonSale(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M7 7h.968v1.939H7z"/><path d="M13.513 3.002L8.438 4.063L3.473 3.002C2.108 3.002 1 4.334 1 5.977v5.036c0 1.644 1.107 2.976 2.473 2.976l5.09-1.082l4.95 1.082c1.364 0 2.472-1.332 2.472-2.976V5.977c-.001-1.643-1.108-2.975-2.472-2.975M5.047 7.016H3.014v.969h1.002v.984h1.031v1.047H4.043v1H1.969V9.954h2.002v-.951h-.997v-.984H1.969V6.983h1.005V5.955h2.073zm3.966 4.031H7.969v-1.031h-.953v1.031H5.969V6.985h1V5.971h1.047v1.014h.997zm3.003-.031H9.969V5.954h1.047V9.97h1zm3.015-4H14.01v.938h.994v1.078h-.973v.933h1v1.052h-2.062V5.97h2.062z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ButtonSell(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.343 2.033L9.01 3.083l-5.375-1.05C2.18 2.033 1 3.359 1 4.996v5.018c0 1.637 1.18 2.964 2.635 2.964l5.375-1.02l5.333 1.02c1.456 0 2.636-1.327 2.636-2.964V4.996c0-1.637-1.18-2.963-2.636-2.963M6.042 6.004H4.021v.97h.995v1.009h1.026V9.08h-1v.978H2.98V8.939h1.989v-.907H3.98V7.023h-1V5.958h1V4.949h2.062zm3 .026H8.021v.962h.995v1.051h-.974v.873h1v1.141H6.98V4.979h2.062zm3 4.027H9.984V4.958h1.031v3.979h1.026zm3 0h-2.058V4.948h1.031V9h1.026z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ButtonStarburst(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.584 7.939c1.535-.506 2.469-1.008 2.401-1.378c-.068-.37-1.118-.53-2.736-.49c1.265-.99 1.967-1.776 1.774-2.101c-.189-.325-1.234-.121-2.742.458c.846-1.354 1.23-2.327.938-2.569c-.292-.242-1.203.299-2.416 1.352c.32-1.557.342-2.602-.018-2.73c-.357-.128-1.024.686-1.797 2.08C9.747.991 9.402 0 9.021 0c-.383 0-.727.99-.967 2.561C7.282 1.166 6.615.352 6.254.479c-.36.129-.337 1.174-.016 2.73c-1.212-1.051-2.124-1.592-2.417-1.352c-.293.24.092 1.216.937 2.57c-1.507-.58-2.553-.783-2.743-.458c-.191.325.51 1.111 1.777 2.102c-1.619-.04-2.672.119-2.738.49c-.067.368.867.872 2.4 1.378c-1.533.506-2.467 1.01-2.4 1.379c.067.369 1.119.53 2.738.489c-1.268.989-1.971 1.774-1.777 2.101c.19.326 1.236.123 2.741-.458c-.845 1.355-1.229 2.329-.937 2.571c.295.242 1.205-.301 2.417-1.351c-.321 1.556-.343 2.601.017 2.729c.359.129 1.028-.686 1.801-2.082c.24 1.57.584 2.561.967 2.561c.381 0 .726-.99.967-2.561c.772 1.395 1.44 2.209 1.799 2.081c.36-.128.338-1.173.016-2.729c1.213 1.051 2.124 1.593 2.416 1.351c.295-.242-.09-1.216-.936-2.569c1.506.581 2.551.782 2.742.458c.191-.326-.512-1.111-1.775-2.101c1.617.04 2.668-.12 2.735-.491c.068-.368-.866-.872-2.401-1.378"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ButtonTriangleUp(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M17.027 8a8 8 0 0 0-8-8a8 8 0 0 0-8 8a8 8 0 0 0 8 8a8 8 0 0 0 8-8M5.154 10c-.205-.186-.205-.774 0-.96l3.467-4.9a.566.566 0 0 1 .746 0l3.478 4.9c.207.185.207.773 0 .96h-7.69z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ButtonTv(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m14.338 3.045l-5.33 1.002l-5.361-1.002c-1.452 0-2.631 1.328-2.631 2.966v5.023C1.016 12.672 2.195 14 3.647 14l5.361-1.031L14.338 14c1.452 0 2.631-1.328 2.631-2.966V6.011c0-1.638-1.179-2.966-2.631-2.966M8.024 7.016H6.026v4.031H4.964V7.016h-1.98V6h5.04zm5.062 4.017h-1.127L9.962 5.965h1.3l1.268 3.666l1.231-3.666h1.294z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CabinCable(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M9.283 7H9V2H7.024v5h-.35C5.203 7 2.013 8.147 2.013 9.563v4.428C2.013 15.406 3.206 16 4.677 16h6.603c1.472 0 2.663-.594 2.663-2.009V9.563C13.943 8.147 10.752 7 9.283 7M5.03 12.046H2.951v-1.703c0-.385.454-1.343 1.51-1.343h.569zm5.016 2.286c0 .37-.294.669-.656.669H6.602a.662.662 0 0 1-.655-.669V9.573c0-.369.294-.573.655-.573H9.39c.362 0 .656.204.656.573zm2.995-2.286H10.94V9h.576c1.146 0 1.524.972 1.524 1.354v1.692z"/><path d="M1.682 4.93c-.367 0-.666-.273-.666-.612c0-.339.299-.613.666-.613c5.575 0 11.077-.985 14.358-2.573c.327-.158.728-.042.899.258c.171.299.045.67-.28.827C13.199 3.891 7.46 4.93 1.682 4.93"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cabinet(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" fill-rule="evenodd"><path d="M4 15.021h10V8.959H4zm1.042-4.98h7.938v3.917H5.042z"/><path fill="currentColor" d="M13.721.039H4.239c-.691 0-1.238.478-1.238 1.178v13.605c0 .614.531 1.178 1.238 1.178h9.482c.842 0 1.238-.548 1.238-1.178V1.217c0-.7-.568-1.178-1.238-1.178M14 15.021H4V8.959h10zm0-8H4V.959h10z"/><path fill="currentColor" d="M10.042 10.041v.999H7.938v-.999H5.042v3.917h7.937v-3.917zm0-7.041H7.958v-.958H5V6h8V2.042h-2.958z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calculator(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.689.041h-5.606v6.876h6.875V1.299c0-.694-.567-1.258-1.269-1.258m-.645 4.004h-3.106V2.938h3.106zM10 15.958h5.674c.71 0 1.284-.569 1.284-1.274V9.042H10zm1.953-4.989h3.07v1.062h-3.07zm0 1.984H15v1.094h-3.047zM7.898.041H2.326A1.25 1.25 0 0 0 1.083 1.3v5.623h6.815zM6 4.023h-.984v1.039H3.98V4.023H2.959V2.982H3.98V1.98h1.036v1.002H6zM1.083 14.676c0 .709.562 1.282 1.255 1.282h5.62V9H1.083zm2.36-3.967l1.046 1.047l1.048-1.049l.732.732l-1.048 1.049l1.061 1.061l-.735.736l-1.062-1.061l-1.022 1.021l-.732-.732l1.021-1.021l-1.045-1.047z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalculatorTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.58 0H4.322C3.592 0 3 .598 3 1.334v13.333C3 15.404 3.592 16 4.322 16h9.258c.729 0 1.42-.596 1.42-1.333V1.334C15 .598 14.31 0 13.58 0M7.021 14H4.987v-1h2.034zm0-5.979H4.987V7h2.034zM10 14H7.986v-1H10zm-2.979-3H4.987v-1h2.034zM10 11H7.986v-1H10zm0-3H7.986V7H10zm3 6h-2V9.979h2zm0-5.98h-2.014V7H13zM14.014 6H4V2h10.014z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarEmpty(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M4 0h.971v1.911H4zm7 0h1v1.906h-1z"/><path d="M15.976 4.959V2.456c0-.771-.606-1.398-1.354-1.398h-1.59v2.026H9.968V1.058H6.034v2.026H2.938V1.058H1.401c-.748 0-1.354.627-1.354 1.398v2.503zM.046 6.003v8.562c0 .772.606 1.399 1.354 1.399h13.221c.748 0 1.354-.627 1.354-1.399V6.003z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarOne(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M13 .031V1h-2V.031H5V1H3V.031H0V16h16V.031zm1.029 13.977H1.955v-9.07h12.074z"/><path d="M6.027 7.957h-.98v-.941h1.902v4.938h-.922zm4 0H9v-.941h1.953v4.938h-.926z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarThree(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" fill-rule="evenodd"><path d="M15 3.902H3c-1.027 0-1 .031-1 1.098v9c0 1.067-.027 1 1 1h12c1.027 0 1 .067 1-1V5c0-1.066.027-1.098-1-1.098M13.957 12H12v1.998h-.959V12H6.937v2h-.9v-2H4v-.952h2.037V9H4.029v-.929h2.008V6.065h.9v2.006h4.104V6.067H12v2.004h1.953V9H12v2.048h1.957z"/><path fill="currentColor" d="M12 9h1.953v-.929H12V6.067h-.959v2.004H6.937V6.065h-.9v2.006H4.029V9h2.008v2.048H4V12h2.037v2h.9v-2h4.104v1.998H12V12h1.957v-.952H12zm-.975 2.048H6.968V9h4.057z"/><path fill="currentColor" d="M14.77 1.051H13V2h-1v-.949H6V2H5v-.949H3.23C1.998 1.051 1 2.061 1 3.306v10.322c0 1.244.998 2.255 2.23 2.255h11.539c1.232 0 2.23-1.011 2.23-2.255V3.306c.001-1.245-.997-2.255-2.229-2.255M16 14c0 1.067.027 1-1 1H3c-1.027 0-1 .067-1-1V5h14zM5 0h.979v1H5zm7 0h1v1h-1z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Call(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.031 11.852c-.428-.539-1.123-1.32-1.718-1.394c-.362-.045-.778.255-1.188.538c-.08.04-.698.408-.773.43c-.396.113-1.241.146-1.752-.32c-.492-.45-1.27-1.283-1.898-2.046c-.6-.786-1.229-1.731-1.551-2.311c-.336-.601-.094-1.396.114-1.746c.038-.063.498-.536.601-.646l.015.018c.381-.32.78-.645.825-.997c.074-.586-.525-1.439-.953-1.979C5.325.858 4.662-.089 3.759.045c-.34.05-.633.169-.922.34L2.829.376a1.823 1.823 0 0 0-.048.037l-.025.013l.003.004c-.166.128-.64.482-.694.53c-.586.521-1.468 1.748-.786 3.955c.506 1.64 1.585 3.566 3.055 5.514l-.008.007c.072.094.146.179.221.27c.07.093.139.185.211.277l.01-.007c1.56 1.879 3.196 3.381 4.689 4.267c2.01 1.192 3.439.655 4.099.228c.062-.041.534-.408.694-.529l.004.004c.006-.006.01-.014.018-.02a3.27 3.27 0 0 0 .043-.033l-.006-.008c.242-.234.436-.484.57-.799c.351-.829-.42-1.693-.848-2.234"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CallForward(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="m14.252 8.774l2.512-2.139a.776.776 0 0 0 0-1.115l-2.512-2.139c-.34-.331-1.191-.452-1.191-.048v1.761H9.139c-.559 0-1.01.424-1.01.946s.451.946 1.01.946h3.922v1.74c0 .278.818.307 1.191.048"/><path d="M14.031 11.852c-.428-.539-1.123-1.32-1.718-1.394c-.362-.045-.778.255-1.188.538c-.08.04-.698.408-.773.43c-.396.113-1.241.146-1.752-.32c-.492-.45-1.27-1.283-1.898-2.046c-.6-.786-1.229-1.731-1.551-2.311c-.336-.601-.094-1.396.114-1.746c.038-.063.498-.536.601-.646l.015.018c.381-.32.78-.645.825-.997c.074-.586-.525-1.439-.953-1.979C5.325.858 4.662-.089 3.759.045c-.34.05-.633.169-.922.34L2.829.376l-.048.037l-.025.013l.003.004c-.166.128-.64.482-.694.53c-.586.521-1.468 1.748-.786 3.955c.506 1.64 1.585 3.566 3.055 5.514l-.008.007c.072.094.146.179.221.27c.07.093.139.185.211.277l.01-.007c1.56 1.879 3.196 3.381 4.689 4.267c2.01 1.192 3.439.655 4.099.228c.062-.041.534-.408.694-.529l.004.004c.006-.006.01-.014.018-.02a3.27 3.27 0 0 0 .043-.033l-.006-.008c.242-.234.436-.484.57-.799c.351-.829-.42-1.693-.848-2.234"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CallReply(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M10.878 3.259L8.366 5.398a.776.776 0 0 0 0 1.115l2.512 2.139c.339.331 1.122.452 1.122.048V6.939h3.991c.558 0 1.011-.424 1.011-.946s-.453-.946-1.011-.946H12v-1.74c0-.278-.748-.307-1.122-.048"/><path d="M14.031 11.852c-.428-.539-1.123-1.32-1.718-1.394c-.362-.045-.778.255-1.188.538c-.08.04-.698.408-.773.43c-.396.113-1.241.146-1.752-.32c-.492-.45-1.27-1.283-1.898-2.046c-.6-.786-1.229-1.731-1.551-2.311c-.336-.601-.094-1.396.114-1.746c.038-.063.498-.536.601-.646l.015.018c.381-.32.78-.645.825-.997c.074-.586-.525-1.439-.953-1.979C5.325.858 4.662-.089 3.759.045c-.34.05-.633.169-.922.34L2.829.376l-.048.037l-.025.013l.003.004c-.166.128-.64.482-.694.53c-.586.521-1.468 1.748-.786 3.955c.506 1.64 1.585 3.566 3.055 5.514l-.008.007c.072.094.146.179.221.27c.07.093.139.185.211.277l.01-.007c1.56 1.879 3.196 3.381 4.689 4.267c2.01 1.192 3.439.655 4.099.228c.062-.041.534-.408.694-.529l.004.004c.006-.006.01-.014.018-.02a3.27 3.27 0 0 0 .043-.033l-.006-.008c.242-.234.436-.484.57-.799c.351-.829-.42-1.693-.848-2.234"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Camera(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" transform="translate(.995 2.98)"><circle cx="7.958" cy="6.958" r="2.958"/><path d="M14.666 2.042h-3.713L9.937.031H6L5 2.042H1.333C.597 2.042 0 2.639 0 3.375v7.25c0 .736.597 1.334 1.333 1.334h13.333c.736 0 1.334-.598 1.334-1.334v-7.25c0-.736-.598-1.333-1.334-1.333M6.953.969h2.094v1.062H6.953zm1.049 10.064a4.052 4.052 0 0 1-4.055-4.048a4.052 4.052 0 0 1 4.055-4.048a4.051 4.051 0 0 1 4.055 4.048a4.051 4.051 0 0 1-4.055 4.048M14 4.031h-2.047V2.969H14z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraProjector(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.997 5.509A3.49 3.49 0 0 0 4.51 2.021a3.492 3.492 0 0 0-3.488 3.488c0 1.086.239 1.905 1.018 2.546l.007 6.154c0 .435.403.734.833.734h5.117zM4.5 7C3.675 7 3 6.327 3 5.5C3 4.674 3.675 4 4.5 4C5.326 4 6 4.674 6 5.5S5.326 7 4.5 7m11.812-1.968H9.006v5.938h7.306c.366 0 .662-.3.662-.667V5.698a.662.662 0 0 0-.662-.666m-.296 2.999h-2.047V6.953h2.047z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CameraSecurity(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M1 12h3.05v1.945H1zm3.16.227l2.028-5.768l1.183.416l-2.027 5.768z"/><path d="m2.595 3.217l-.486 1.664s-.294.774.427 1.014c.812.248 6.992 2.084 6.992 2.084l1.391-2.38z"/><path d="M13.709 4.971c.105-.039 2.249-1.2 2.249-1.2c-.162-.163-.257-.286-.495-.355L4.145.101a1.395 1.395 0 0 0-1.738.937l-.351 1.197s9.708 2.844 10.006 2.886l-1.019 1.798l2.946 1.025l.99-2.922c0-.001-1.11-.001-1.27-.051m-11.95 7.208l-.356-.529c1.944-1.198 2.61-4.396 2.59-6.771l.658-.006c.021 2.656-.726 5.971-2.892 7.306"/><path d="M2.029 12.996c0 1.645-.908 2.977-2.029 2.977v-5.952c1.121 0 2.029 1.332 2.029 2.975"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CanWater(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.168 5.069c-.506-2.648-2.262-4.037-4.626-4.037c-2.362 0-4.125 1.384-4.636 4.027c-.473.012-.849.314-.849.69v1.847L3.904 5.575L3.9 5.573c.314-.551.385-1.075.121-1.346c-.39-.398-1.344-.099-2.131.668c-.787.77-1.111 1.715-.723 2.113c.266.273.797.214 1.355-.092l3.534 6.101v2.257c0 .382.39.692.873.692h9.158c.483 0 .873-.311.873-.692V5.749c.002-.359-.348-.646-.792-.68M11.528 2c1.784 0 3.062.938 3.472 3H8c.41-2.062 1.744-3 3.527-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Candle(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M8 5h.926v1.823H8z"/><path d="M10.287 5.973c.11-2.604-.986-3.368-1.61-3.368c-.65 1.177-2.032 1.498-2.329 3.394C5.034 3.567 8.004 3.059 8.153.027c.252-.001 4.403 3.559 2.134 5.946m1.031 1.436s-.518.644-2.244.644c-1.727 0-3.37-.644-3.37-.644a.677.677 0 0 0-.681.672v7.226c0 .371.305.673.681.673h5.614a.678.678 0 0 0 .682-.673V8.081a.677.677 0 0 0-.682-.672m-.258 7.607H9.951V8.954h1.109z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Candy(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.484.037c-1.494 0-2.791.786-3.469 1.947h6.936C11.275.823 9.977.037 8.484.037m4.24 2.984H4.278c-.137 0-.248.102-.248.226v2.498c0 .124.111.225.248.225h8.446c.138 0 .248-.101.248-.225V3.247c0-.124-.111-.226-.248-.226m-.742 3.985H5.009c.584 1.052 1.742 1.802 2.99 1.997v6.95h.262v.031h.597v-.031h.126V9.021a3.98 3.98 0 0 0 2.998-2.015"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CandyStick(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.75.468A5.582 5.582 0 0 0 10.682.13A5.476 5.476 0 0 0 9.496.009A5.351 5.351 0 0 0 7.28.527a4.487 4.487 0 0 0-2.013 1.908C5.216 2.528.146 13.977.146 13.977c-.318.72.014 1.56.738 1.872a1.449 1.449 0 0 0 1.898-.743L7.751 3.844c.223-.503.723-.847 1.346-.952c.127-.021.256-.037.393-.037a2.735 2.735 0 0 1 1.797.66l.635.825c.195.434.23.899.045 1.318a1.41 1.41 0 0 0 .67 1.834c.025.013.045.028.07.039c.502.217 1.083.102 1.494-.229c.168-.137.312-.308.404-.513c.164-.374.264-.758.317-1.146c.281-2.067-.967-4.229-3.172-5.175"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Car(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<defs><path id="siGlyphCar0" d="M6.514 1.607C6.319 2.141 5.934 3 5.934 3h3.092V.957H7.544c-.28 0-.843.131-1.03.65m5.932-.64H9.968V3h4.086l-.484-1.447s-.41-.586-1.124-.586"/></defs><g fill="none" fill-rule="evenodd" transform="translate(1 4)"><circle cx="12.49" cy="6.49" r="1.49" fill="currentColor"/><circle cx="3.49" cy="6.49" r="1.49" fill="currentColor"/><use href="#siGlyphCar0"/><use href="#siGlyphCar0"/><path fill="currentColor" d="M16 4.836c0-.597-.078-1.137-.753-1.524l-.585-1.974C14.662.947 13.747 0 12.618 0H7.339C6.21 0 5.296 1.088 5.296 1.338L4.432 3l-.781.107C1.633 3.107.583 4.334.583 5.515L.024 6.984h1.007a2.454 2.454 0 0 1-.056-.518a2.508 2.508 0 0 1 2.519-2.498a2.508 2.508 0 0 1 2.518 2.498c0 .178-.02.351-.056.518h4.074a2.454 2.454 0 0 1-.056-.518a2.508 2.508 0 0 1 2.519-2.498a2.508 2.508 0 0 1 2.519 2.498c0 .178-.02.351-.056.518H16zM9.025 3H5.933s.385-.858.58-1.393c.188-.52.751-.65 1.03-.65h1.482zm.944 0V.967h2.478c.714 0 1.124.586 1.124.586L14.055 3z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CarGarage(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M.043 2.083v11.855h1.939V3.943h12.036v9.995h1.94V2.083z"/><path d="M3 5h9.956v.952H3zm0 2h9.957v.971H3zm8.97 5.96v.515c0 .271-.438.492-.975.492c-.536 0-.974-.221-.974-.492h.001v-.485H5.946v.498c0 .247-.434.449-.964.449s-.961-.202-.961-.449v-.509h.139c-1.287-.117-1.156-1.13-1.156-1.13c0-.73.152-1.901.996-2.402V8h1v1.192l.059-.008S7.096 9.001 8 9.001c.903 0 2.79.214 2.901.23l.098.018V8h1v1.457c.756.331 1 1.354 1 2.392c0 0-.027.93-1.03 1.111zm-7.001-.991h-1v-1h1zm3.99-.011h-2v-1h2zm3.01.011h-1v-1h1z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Casette(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" fill-rule="evenodd" transform="translate(1 3)"><path d="M4.629 4C3.178 4 2 5.189 2 6.656s1.178 2.656 2.629 2.656c1.453 0 2.629-1.189 2.629-2.656S6.082 4 4.629 4m6.981 0C10.169 4 9 5.189 9 6.656s1.169 2.656 2.61 2.656c1.442 0 2.611-1.189 2.611-2.656S13.053 4 11.61 4"/><path fill="currentColor" d="M14.654 0h-.347l-.843 1.954H2.61L1.756 0h-.36C.65 0 .046.624.046 1.392L.001 9.608c0 .768.604 1.392 1.35 1.392h13.32c.746 0 1.285-.624 1.285-1.392L16 1.392C16.002.624 15.398 0 14.654 0M4.5 9a2.5 2.5 0 1 1 .001-5.001A2.5 2.5 0 0 1 4.5 9m7 0a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5"/><path fill="currentColor" d="M3.529.908h8.707L12.927 0H3.033z"/><ellipse cx="4.453" cy="6.469" fill="currentColor" rx="1.453" ry="1.469"/><ellipse cx="11.453" cy="6.469" fill="currentColor" rx="1.453" ry="1.469"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CashierMachine(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.641 4.012h-1.656V2.957h1.352c.346 0 .625-.346.625-.771V.772c0-.427-.279-.771-.625-.771H9.646c-.346 0-.624.345-.624.771v1.414c0 .426.278.771.624.771h1.378v1.055H4.353c-.37 0-.67.3-.67.668l-1.67 8.631V16h13.998v-2.689L14.312 4.68a.67.67 0 0 0-.671-.668m-7.625 8.004H4.985v-1.047h1.031zm-1.032-2V8.969h1.031v1.047zm3.032 2.015H6.985v-1.062h1.031zm0-2.031H6.969V8.984h1.047zm2 2.029H8.969v-1.061h1.047zM8.969 10V8.969h1.047V10zm1.047-2H4.985V6h5.031zM13 8h-2.018V6.98H13z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Castle(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16 5h-1v1h-1v-.985L13 5v4h-2V4c.549 0 1-.434 1-.97c0 0-3.443-2.988-3.992-2.988h-.016c-.548 0-3.991 2.988-3.991 2.988c0 .536.493.97 1.041.97v5H2.98V6s-.457.055-.979.015V7h-1V6h-1v8.951c0 .556.489 1.007 1.042 1.007H7v-2.896c0-1.427 1.041-1.125 1.041-1.125s1.042-.302 1.042 1.125c0 0 .002 2.854 0 2.896H15c.553 0 1-.487 1-1.087zM1 10V9h1v1zm2 3H2v-1h1zm2 0H4v-1h1zm4-6H6.953V4.935H9zm3 6h-1v-1h1zm2 0h-1v-1h1zm1-3h-1.021V9H15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaterpillarMachine(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M13.017 10.012a1.978 1.978 0 0 0-1.983 1.97c0 1.087.89 1.97 1.983 1.97a1.977 1.977 0 0 0 1.981-1.97c0-1.087-.889-1.97-1.981-1.97m-6.033.05c-1.09 0-1.979.865-1.979 1.934c0 1.066.889 1.934 1.979 1.934c1.094 0 1.98-.867 1.98-1.934c.001-1.068-.886-1.934-1.98-1.934M3.276 6.496c0 1.379-1.837 2.496-3.276 2.496L.67 4c1.439 0 2.606 1.117 2.606 2.496"/><path d="M6.155 8.625L1.509 6.82l.414-1.361l4.645 1.805z"/><path d="M15.225 7.9c.031-.133-.267-.271-.267-.414V3.889c0-1.013-.501-1.836-1.507-1.836H9.16a1.832 1.832 0 0 0-1.824 1.836l-.672 3.597l.002.014c-.835.24-1.371.801-1.371 1.846c0 .044.008.083.009.127a3.014 3.014 0 0 1 1.681-.51a3.009 3.009 0 0 1 2.971 2.506c.027.119.044.227.052.312a3.016 3.016 0 0 1 3.011-2.811c1.227 0 2.28.731 2.756 1.776c.146-.407.173-.886.173-1.401c-.002-.665-.286-1.128-.723-1.445m-4.172-.838H8.938V3.355c0-.288.105-.4.234-.4h1.645c.13 0 .236.112.236.4zm2.965-1.003h-2.059V2.938h2.059z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CenteJustify(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M17 1.938c0 .518-.42.938-.938.938H1.98A.938.938 0 0 1 1.98 1h14.082c.518 0 .938.42.938.938m0 3c0 .518-.42.938-.938.938H1.98A.938.938 0 0 1 1.98 4h14.082c.518 0 .938.42.938.938m0 3c0 .518-.42.938-.938.938H1.98A.938.938 0 0 1 1.98 7h14.082c.518 0 .938.42.938.938m0 3c0 .518-.42.938-.938.938H1.98a.938.938 0 0 1 0-1.876h14.082c.518 0 .938.42.938.938m-3 3c0 .518-.42.938-.938.938H4.98a.938.938 0 0 1 0-1.876h8.082c.518 0 .938.42.938.938"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChairOne(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.377 7.984H6.641l-.592-6.641S5.704-.001 9.008-.001c3.306 0 2.96 1.344 2.96 1.344zM12.953 16H12.1l-1.026-2.984h1.879zm-7.055 0h-.896v-2.984h1.973zm7.078-4.778c0 .415-.252.75-.563.75H5.59c-.311 0-.562-.335-.562-.75c0 0 .429-2.186 1.134-2.186h5.682c.799 0 1.132 2.186 1.132 2.186m.896-2.867c0 .349-.19.631-.421.631c-.232 0-.42-.282-.42-.631V5.672c0-.349.188-.631.42-.631c.23 0 .421.282.421.631zm-8.925 0c0 .349-.209.631-.465.631c-.258 0-.466-.282-.466-.631V5.672c0-.349.208-.631.466-.631c.256 0 .465.282.465.631z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChairTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M8 14.095V11h1.938v3.085c1.739.205 3.031.972 3.031 1.886H5.045c0-.9 1.254-1.657 2.955-1.876m3.379-6.142H6.622L6.027 1.34S5.681.016 9 .016c3.32 0 2.975 1.324 2.975 1.324z"/><path d="M12.969 11.225c0 .416-.231.752-.629.752H5.66c-.396 0-.629-.336-.629-.752c0 0 .458-2.195 1.357-2.195h5.066c1.024-.001 1.515 2.195 1.515 2.195m1.006-3.715c0 .395-.213.475-.476.475s-.476-.08-.476-.475V4.485c0-.396.213-.453.476-.453s.476.058.476.453zm-9 0c0 .395-.213.475-.476.475s-.476-.08-.476-.475V4.485c0-.396.213-.453.476-.453s.476.058.476.453z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChampionCup(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.969 1H3.082v1.031h-3.09v4.836c0 1.869 2.086 3.407 4.133 3.567c.672.72 1.854 1.243 2.906 1.466v1.264c-1.006.309-2.803 1.035-2.906 1.826h6.904c-.104-.791-2.084-1.518-3.092-1.826V11.9c1.056-.223 2.291-.746 2.964-1.466c2.046-.16 4.04-1.698 4.04-3.567V2.031h-2.972zM.941 2.947H3.01v4.166c0 .822.212 1.604.585 2.308C2.047 8.928.941 7.658.941 6.172zm8.104 5.081l-1.544-.851l-1.546.851l.295-1.8L5 4.954l1.727-.263l.774-1.636l.772 1.636L10 4.954L8.75 6.228zm2.295 1.526c.375-.736.59-1.55.59-2.411V2.974h2.074V6.16c0 1.553-1.111 2.879-2.664 3.394"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartColumn(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M16 14.031H.984V1.016H.027L0 14.95h.027v.029L16 14.95z"/><path d="M4.958 8.021H2.016v4.964h2.942zm5.011-1.974H7.016v6.922h2.953zm4.984-2.016H12v8.947h2.953z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartColumnDecrease(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M0 16h15.938v.969H0zm14.906-7.058l-4.756-.875l1.609-1.609l-4.78-4.057L2.9 5.733a.5.5 0 1 1-.658-.753l4.745-3.914l5.479 4.685l1.684-1.684z"/><path d="M6 6h2.951v8.878H6zM1 9h2.982v5.878H1zm10 2h2.951v3.892H11z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartColumnIncrease(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 15h15.938v.969H1zM15.906.065L11.15.94l1.609 1.609l-4.78 4.057L3.9 3.274a.5.5 0 1 0-.658.753l4.745 3.914l5.479-4.685L15.15 4.94zM12 7h2.951v6.878H12zM2 8h2.982v5.878H2zm5 2h2.951v3.892H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartDecrease(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M17 14.051H2.885V.084H1.041L1 15.875h.041v.041L17 15.875z"/><path d="M16.816 2h-3.727a.13.13 0 0 0-.129.129l1.527 1.533l-3.476 4.463L8 6l-3.973 5.188l.062 1.75L8.061 8l2.949 2l4.36-5.449L16.813 6a.13.13 0 0 0 .129-.129V2.129A.125.125 0 0 0 16.816 2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChartPiece(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M9.392 1.049c-.349-.051-.562-.033-.925-.033c-4.095 0-7.424 3.334-7.424 7.43s3.329 7.431 7.424 7.431c4.094 0 7.423-3.335 7.423-7.431c0-.346.044-.598-.002-.929l-7.399.961z"/><path d="m11.375.047l-.945 6.539l6.613-.928C16.638 2.764 14.314.466 11.375.047"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Checked(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.432 6.189a1 1 0 0 1 1.415 0L6.968 8.31l6.179-6.179a.99.99 0 0 1 1.401.013l2.122 2.122a.991.991 0 0 1 .014 1.4l-9.022 9.021a.99.99 0 0 1-1.401-.014l-4.95-4.95a.998.998 0 0 1 0-1.413z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cheese(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.744 2.128c-.22.039-1.438.411-2.941 1.085c.301.464.383.945.158 1.266c-.346.5-1.295.416-2.119-.188c-.052-.036-.096-.079-.144-.118C3.92 5.454 1.031 6.841 1.031 6.841v2.594c1.207.125 2.154 1.123 2.164 2.354l7.752.698c.144-.975.73-1.706 1.434-1.706c.766 0 1.388.865 1.457 1.967l2.43.218l.242-2.24c-.727-.065-1.283-.397-1.283-.801c0-.435.647-.786 1.459-.81c0 0 .109-.961.248-2.273c.304-1.774-3.886-4.943-5.19-4.714m2.79 2.583c.463 0 .838.354.838.792c0 .435-.375.791-.838.791c-.461 0-.837-.356-.837-.791c0-.439.376-.792.837-.792m-3.003-.945c.566 0 1.025.312 1.025.697c0 .385-.459.696-1.025.696c-.566 0-1.024-.312-1.024-.696c0-.386.458-.697 1.024-.697m-5.014.924c.458 0 .827.348.827.779c0 .43-.369.777-.827.777c-.458 0-.827-.348-.827-.777c-.001-.431.369-.779.827-.779m-.011 6.435c-.887 0-1.604-.502-1.604-1.125c0-.621.717-1.125 1.604-1.125c.885 0 1.603.504 1.603 1.125c-.001.623-.718 1.125-1.603 1.125m9.525-3.094h-3.01c-.228.482-.714 1.016-1.506 1.016c-.794 0-1.285-.533-1.512-1.016h-8.02V6.937h8.051c.255-.447.678-.984 1.434-.984c.755 0 1.266.537 1.521.984h3.042z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cherry(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.74 9.09c-.117-.279-.24-.552-.354-.8c-.713-1.55-1.645-3.086-2.449-4.165c-.83-1.113-3.965-3.803-4.57-4.043l-.262.645c.492.194 1.488 1.064 1.488 1.064c.177.15.412 1.021.41 1.163a2.598 2.598 0 0 0 0 .266c.027.613-.078 1.288-.323 2.06c-.486 1.534-1.577 2.812-2.712 3.833c-.164-.023-.328-.05-.5-.05c-1.916 0-3.469 1.539-3.469 3.438c0 1.899 1.553 3.438 3.469 3.438s3.469-1.539 3.469-3.438c0-1.43-.881-2.654-2.135-3.172C6.884 8.292 7.89 7.023 8.371 5.501c.271-.855.387-1.612.354-2.314c-.003-.057 1.158.904 1.658 1.576c.783 1.05 1.676 2.324 2.388 3.869c.069.149.144.312.219.479c-1.688.232-2.991 1.653-2.991 3.39c0 1.898 1.553 3.438 3.469 3.438s3.469-1.539 3.469-3.438c.001-1.808-1.41-3.272-3.197-3.411"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChristmasMistletoe(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" transform="translate(0 3)"><circle cx="7.428" cy="9.428" r="1.428"/><ellipse cx="11.461" cy="8.493" rx="1.461" ry="1.493"/><ellipse cx="8.477" cy="5.493" rx="1.477" ry="1.493"/><path d="M7.079 2.363c-.074.025-.14.088-.2.145a1.833 1.833 0 0 1-.854.443c-.419.1-1.159-.211-1.522-.507c-.498 1.206-2.152 1.447-2.881.451c.103.571.149.824-.273 1.226c-.412.391-.839.62-1.339.815c.168-.065.325-.116.507-.122c.726-.028 1.482.411 1.908 1.002c.201.281.338.627.307.985c-.005.062-.062.21-.008.244c.078.049.277-.116.346-.147c.262-.117.557-.161.839-.141c.442.029.731.307.982.654c.149.204.353 1.272.788.85c.153-.149.194-.484.321-.677c.148-.227.374-.366.527-.569c.189-.25.019-.615-.096-.858a1.557 1.557 0 0 1-.126-.556c-.341-.225-1.189-.632-2.855-.619h-.002a.324.324 0 0 1-.322-.325a.324.324 0 0 1 .319-.33l.085-.001c1.424 0 2.325.301 2.844.565c.047-.181.105-.36.174-.525c.096-.232.275-.358.469-.488c.161-.109.322-.248.283-.478c-.061-.348-.206-.675-.221-1.037m8.93-2.199c-.528.172-1.02.273-1.603.248c-.599-.027-.725-.243-1.008-.725c.053 1.188-1.434 2.056-2.609 1.486c-.104.447-.499 1.141-.899 1.333a2.057 2.057 0 0 1-.974.214c-.084-.004-.177-.009-.252.021c-.072.027.525.669.634.724c.327.162.671.187.896.48c.08.104.158.226.229.35a11.926 11.926 0 0 1 2.483-2.525a.356.356 0 0 1 .474.045a.297.297 0 0 1-.048.439c-1.594 1.204-2.437 2.488-2.65 2.841a.697.697 0 0 0 .175.371c.291.31.671.171 1.046.275c.23.063.451.161.653.282c.188.114.307.304.542.318c.31.018.539-.121.654-.385c.158-.355-.061-.558-.08-.899c-.021-.412.033-.799.371-1.101c.214-.191.479-.346.766-.425c.075-.021.342-.025.372-.11c.022-.061-.118-.132-.161-.174c-.257-.243-.37-.582-.388-.916c-.038-.701.289-1.5.891-1.938a1.68 1.68 0 0 1 .486-.229"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChristmassBall(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.346 2.287a6.857 6.857 0 0 0-3.115-.121l8.605 8.547a6.767 6.767 0 0 0-.115-3.072c-.073-.269-5.123-5.287-5.375-5.354M2.178 13.826c2.132 2.118 5.215 2.66 7.711 1.602L.567 6.17C-.5 8.647.047 11.709 2.178 13.826M2.033 4.06c-.271.269-.503.56-.713.861l9.824 9.757a6.14 6.14 0 0 0 .867-.706a6.101 6.101 0 0 0 1.253-1.828l-9.39-9.328A6.06 6.06 0 0 0 2.033 4.06m7.925 6.912h1.085v1.055H9.958zm-2-2.008h1.085v1.098H7.958zm-2-1.986h1.085v1.076H5.958zm-2-2.07h1.085v1.139H3.958zm6.085-2.012l3.742 3.776l1.551-1.563l-3.744-3.777zM14.127.698c-.848 0-1.559.58-1.783 1.368l.598.593c-.002-.023-.008-.044-.008-.067a1.2 1.2 0 0 1 1.193-1.211a1.2 1.2 0 0 1 1.193 1.211c0 .668-.533 1.21-1.193 1.21c-.014 0-.025-.004-.041-.005l.596.593a1.886 1.886 0 0 0 1.312-1.799c0-1.044-.838-1.893-1.867-1.893"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChristmassEgg(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M12.764 3.12C11.707 1.23 9.989 0 8.045 0C6.139 0 4.446 1.188 3.384 3.021c1.468.581 3.099.908 4.817.908c1.619 0 3.16-.289 4.563-.809M2.223 10.865C2.818 14.28 5.198 16 8.045 16c2.801 0 5.143-1.663 5.787-4.961c-1.689.806-3.756 1.389-5.781 1.389c-2.17 0-4.056-.649-5.828-1.563m1.951-3.074c.723 0 1.418-.265 1.906-.727l.15-.141l.15.141c.486.463 1.181.727 1.904.727c.705 0 1.363-.241 1.855-.678l.145-.13l.146.13c.491.437 1.15.678 1.855.678c.611 0 1.195-.192 1.662-.533a9.9 9.9 0 0 0-.9-3.223a13.365 13.365 0 0 1-4.845.899c-1.821 0-3.552-.363-5.109-1.009a9.824 9.824 0 0 0-.915 3.053l.072.07c.49.473 1.191.743 1.924.743m8.068-2.977l1.049.894l-1.049.894l-1.049-.894zm-4.185.565l1.049.894l-1.049.894l-1.05-.894zm-4.2-.567l1.049.893l-1.049.894l-1.049-.894z"/><path d="M13.982 8.234c-.498.247-1.084.474-1.695.474c-.75 0-1.451-.28-2.002-.636c-.549.355-1.252.636-2 .636c-.766 0-1.502-.299-2.055-.673c-.555.374-1.292.673-2.057.673c-.764 0-1.494-.297-2.045-.669a7.143 7.143 0 0 0-.052.846c0 .367.032.711.082 1.043c1.774.789 3.687 1.366 5.892 1.366c2.064 0 4.163-.521 5.858-1.22c.067-.374.104-.768.104-1.189a6.366 6.366 0 0 0-.03-.651M5.945 10.3c-.502 0-.91-.296-.91-.661c0-.363.408-.66.91-.66s.91.297.91.66c0 .365-.408.661-.91.661m4.266.02c-.502 0-.91-.295-.91-.66c0-.364.408-.66.91-.66c.503 0 .911.296.911.66c0 .365-.408.66-.911.66"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChristmassHat(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.33.691c-.494.018-1.045.411-1.498.159c-2.182-1.212-6.19-1.157-7.729.963c-.832 1.146-.594 2.663-1.373 3.799c-1.061 1.549-2.763 2.326-2.699 4.593c.031 1.101 1.533.722 2.312.718a2420.625 2420.625 0 0 1 7.781-.017c2.919 0 1.506-1.327 1.157-3.068c-.166-.837-.69-1.951-.988-2.72c-.366-.938-1.649-3.04-.178-3.325c2.037-.396.203 1.445 3.218 2.177c.88.213 1.638-.739 1.638-1.647A1.66 1.66 0 0 0 15.883.772a1.425 1.425 0 0 0-.553-.081m-.405 13.978c0 .72-.624 1.303-1.39 1.303H2.422c-.766 0-1.389-.583-1.389-1.303v-1.303c0-.719.623-1.302 1.389-1.302h11.113c.766 0 1.39.583 1.39 1.302z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChristmassTree(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="m6.781 3.981l1.221-.936l1.217.936l-.465-1.516l1.219-.938H8.467L8.002.011l-.467 1.516H6.028l1.22.938z"/><path d="m10.984 11.979l1.934.016l-2.902-4.031H12L7.985 3.98L4.031 7.996h2.031l-3.047 3.953h2.016l-3.016 4.016l11.977-.017zm-4.968 3.032H4.969v-1.062h1.047zm1-4H5.969V9.949h1.047zm1.015-4H6.969V5.949h1.062zm.938 2.937h1.047v1.062H8.969zm1 5.063v-1.062h1.047v1.062z"/><path d="M12 7.995h.969v.984H12zm1 4h.969v.984H13zm-11 0h.969v.984H2zm1-4h.969v.984H3z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Circle(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<circle cx="9" cy="9" r="8" fill="currentColor" fill-rule="evenodd"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleBackward(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.031.021c-4.394 0-7.955 3.567-7.955 7.969s3.562 7.969 7.955 7.969c4.393 0 7.955-3.567 7.955-7.969S13.424.021 9.031.021M12.5 10.346c-.793-1.449-1.416-2.126-3.232-2.126l-.236.001v1.598a.488.488 0 0 1-.678.035L5.125 7.432a.552.552 0 0 1 0-.775l3.229-2.422c.215-.215.465-.25.678-.036v1.633l.256-.001c2.616 0 3.785 1.735 3.785 4.264c-.001.88-.406.558-.573.251"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleControlPad(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 .032a8 8 0 1 0 0 16a8 8 0 0 0 0-16M5 9.859a.755.755 0 0 1-.88 0L2.182 8.345c-.243-.19-.243-.498 0-.688L4.12 6.143a.75.75 0 0 1 .88 0zm1.142-5.74l1.514-1.937c.19-.243.498-.243.687 0L9.856 4.12a.744.744 0 0 1 0 .881H6.143a.754.754 0 0 1-.001-.881zm3.715 7.742L8.344 13.8c-.193.243-.5.243-.689 0l-1.512-1.939a.748.748 0 0 1 0-.879h3.714a.746.746 0 0 1 0 .879m2.024-2.004a.751.751 0 0 1-.881 0V6.143a.751.751 0 0 1 .88 0l1.937 1.513c.244.191.244.5 0 .689z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleDrashed(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.922 7.984c0-.385.04-.76.108-1.124l-.841-.319a6.94 6.94 0 0 0-.152 1.443c0 .671.101 1.317.277 1.932l.818-.369a6.102 6.102 0 0 1-.21-1.563m7.549 5.892A6.05 6.05 0 0 1 9 14.064c-.422 0-.834-.044-1.231-.126l-.315.83c.498.114 1.015.18 1.547.18c.637 0 1.252-.094 1.839-.254zm-6.75-8.892a6.132 6.132 0 0 1 1.844-2.01l-.367-.816a7.015 7.015 0 0 0-2.317 2.508zM14.21 11.1a6.114 6.114 0 0 1-1.904 1.979l.365.812a7.004 7.004 0 0 0 2.378-2.472zm-10.255.271l-.811.365a7.006 7.006 0 0 0 2.445 2.312l.318-.84a6.121 6.121 0 0 1-1.952-1.837m3.424-9.238A6.027 6.027 0 0 1 9 1.906c.394 0 .777.041 1.149.112l.318-.839a6.92 6.92 0 0 0-3.457.135zm6.684 2.495l.812-.365a7.01 7.01 0 0 0-2.534-2.385l-.318.84a6.105 6.105 0 0 1 2.04 1.91m1.63 1.46l-.816.367c.127.49.202 1 .202 1.529c0 .432-.048.852-.132 1.258l.83.314a6.965 6.965 0 0 0 .186-1.572a6.9 6.9 0 0 0-.27-1.896"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleError(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.618 2.397C10.513-.708 5.482-.713 2.383 2.386c-3.101 3.102-3.098 8.131.009 11.236c3.105 3.105 8.137 3.109 11.235.01c3.1-3.099 3.097-8.13-.009-11.235m-4.003 8.954L7.927 9.663l-1.688 1.688c-.689.689-1.207 1.289-2.029.468c-.82-.821-.223-1.339.469-2.029l1.688-1.687l-1.688-1.688c-.69-.689-1.289-1.207-.469-2.029c.822-.82 1.34-.221 2.029.469l1.688 1.686l1.688-1.686c.69-.689 1.205-1.29 2.027-.469c.822.822.223 1.34-.467 2.029L9.487 8.102l1.688 1.687c.689.691 1.289 1.209.467 2.03c-.82.821-1.337.221-2.027-.468"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleForward(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M.076 7.989c0 4.401 3.562 7.969 7.955 7.969c4.394 0 7.955-3.567 7.955-7.969S12.424.02 8.031.02C3.639.021.076 3.588.076 7.989m3.914 2.105c0-2.528 1.169-4.264 3.785-4.264l.256.001V4.198c.213-.214.463-.179.678.036l3.229 2.422a.552.552 0 0 1 0 .775L8.709 9.853a.488.488 0 0 1-.678-.035V8.22l-.236-.001c-1.816 0-2.439.677-3.232 2.126c-.167.308-.573.63-.573-.251"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleHelp(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9 0a8 8 0 0 0-8 8a8 8 0 0 0 8 8a8 8 0 0 0 8-8a8 8 0 0 0-8-8m0 14a1 1 0 1 1 0-2a1 1 0 0 1 0 2m1.647-5.757c-.473.365-.734.566-.734 1.147v.814c0 .459-.406.834-.907.834c-.502 0-.909-.375-.909-.834V9.39c0-1.357.831-1.998 1.38-2.422c.474-.366.734-.566.734-1.146c0-.654-.541-1.188-1.205-1.188c-.665 0-1.205.533-1.205 1.188c0 .461-.408.833-.909.833c-.5 0-.909-.372-.909-.833c0-1.574 1.357-2.854 3.023-2.854c1.665 0 3.021 1.279 3.021 2.854c0 1.356-.83 1.996-1.38 2.421"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleInfo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.969 0a7.938 7.938 0 1 0 0 15.876A7.938 7.938 0 0 0 8.97 0zM8 3h2v2H8zm2 8.765C10 12.447 9.554 13 9 13c-.553 0-1-.552-1-1.235v-4.53C8 6.555 8.447 6 9 6c.554 0 1 .553 1 1.235z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleLoadLeft(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.947 8.379c-.075-4.122-3.351-7.412-7.303-7.337c-3.917.073-6.853 2.89-6.83 6.964H.204c-.217.22-.217.876 0 1.096l1.911 1.758a.546.546 0 0 0 .785 0l1.92-1.758c.219-.22.219-.876 0-1.096H3.204C3.18 4.668 5.462 2.461 8.67 2.402c3.232-.062 5.912 2.63 5.975 6.002c.062 3.322-2.445 6.077-5.613 6.216v1.361c3.898-.139 6.991-3.521 6.915-7.602"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleLoadRight(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M.04 7.379C.115 3.257 3.392-.033 7.344.042c3.917.073 6.854 2.929 6.83 7.003h1.545a.572.572 0 0 1 0 .799L13.873 9.86a.546.546 0 0 1-.785 0l-1.725-2.016a.57.57 0 0 1 0-.799h1.422c.023-3.338-2.259-5.584-5.467-5.644c-3.233-.062-5.912 2.63-5.975 6.002c-.062 3.322 2.445 6.077 5.613 6.216v1.361C3.059 14.842-.035 11.46.04 7.379"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CirclePlus(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.021.097c-4.396 0-7.958 3.558-7.958 7.943c0 4.388 3.562 7.945 7.958 7.945c4.395 0 7.958-3.558 7.958-7.945c0-4.386-3.564-7.943-7.958-7.943m2.304 8.985h-1.237v1.237c0 .979.059 1.769-1.088 1.769c-1.144 0-1.088-.79-1.088-1.769V9.082H6.675c-.979 0-1.769.056-1.769-1.088c0-1.146.79-1.088 1.769-1.088h1.237V5.669C7.912 4.69 7.856 3.9 9 3.9c1.146 0 1.088.79 1.088 1.769v1.237h1.237c.979 0 1.769-.059 1.769 1.088c0 1.144-.79 1.088-1.769 1.088"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleRemove(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.021 1.097C3.625 1.097.063 4.655.063 9.04c0 4.388 3.562 7.945 7.958 7.945c4.395 0 7.958-3.558 7.958-7.945c0-4.386-3.564-7.943-7.958-7.943M10.271 10H5.729C4.772 10 4 10.05 4 9c0-1.053.772-1 1.728-1h4.544C11.228 8 12 7.946 12 9c0 1.051-.772 1-1.728 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleStar(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 .062c-4.419 0-8 3.559-8 7.947c0 4.39 3.581 7.949 8 7.949c4.418 0 8-3.56 8-7.949C16 3.621 12.418.062 8 .062m3.108 11.963L8.021 9.902l-3.088 2.123L6.112 8.59L3.024 6.465h3.817l1.18-3.435l1.18 3.435h3.816L9.93 8.59z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleTriangleDown(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 8.041a8 8 0 0 0 8 8a8 8 0 0 0 8-8a8 8 0 0 0-16 0M12.846 6c.205.185.205.772 0 .96l-3.467 4.9a.568.568 0 0 1-.746 0l-3.479-4.9c-.205-.187-.205-.774 0-.96z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleTriangleLeft(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9 0a8 8 0 0 0-8 8a8 8 0 0 0 8 8a8 8 0 0 0 8-8a8 8 0 0 0-8-8m2 11.846c-.186.205-.774.205-.96 0l-4.9-3.467a.563.563 0 0 1 0-.745l4.9-3.48c.185-.205.773-.205.96 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleTriangleRight(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9 16a8 8 0 0 0 8-8a8 8 0 0 0-8-8a8 8 0 0 0-8 8a8 8 0 0 0 8 8M7 4.154c.186-.205.775-.205.96 0l4.9 3.467a.567.567 0 0 1 0 .745l-4.9 3.48c-.185.205-.774.205-.96 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func City(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 8v8h3V8zm2 7H1.979v-2.021H3zm.021-3H1.98V9.979h1.041zM10 5V3H9V0H8v3H7v2H5v11h7V5zM7 15H6v-2h1zm0-2.958H6V10h1zm0-3H6V7h1zM9 15H8v-2h1zm0-2.958H8V10h1zm0-3H8V7h1zM11 15h-1v-2h1zm0-2.958h-1V10h1zm0-3h-1V7h1zM13 7v9h4v-5zm2.031 8.062H14v-1.094h1.031zm0-2H14v-1.094h1.031zm0-2H14V9.968h1.031z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clamp(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.297 1.938h5.656V.042H6.297A2.27 2.27 0 0 0 4 2.335v6.371c0 1.287 1.01 2.273 2.297 2.273H9v3.089h-.64v.815c-.825.266-1.39.739-1.39 1.116h4.982c0-.376-.559-.846-1.375-1.112v-.819h-.64v-3.089h1.999V9.04H9.937V6.948c.034-.01.062-.021.093-.031h.906v-.891H8.04v.891h.853c.036.012.066.023.105.034V9.04H6.295a.335.335 0 0 1-.334-.334V2.27a.337.337 0 0 1 .336-.332"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clapboard(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m2.042 7.954l14.916-3.523l-.98-4.056L1.063 3.898zM13.55 4.357l-1.448-2.312l2.01-.486l1.449 2.312zm-3.351.81L8.75 2.855l2.009-.485l1.45 2.312zm-3.351.81L5.4 3.665l2.01-.485l1.449 2.312zm-3.35.81L2.049 4.475l2.009-.485l1.449 2.312zm-1.477 1.25H17v7.95H2.021z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClapboardPlay(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m.979 8.954l14.916-3.523l-.98-4.056L0 4.898zm11.508-3.597l-1.448-2.312l2.01-.486l1.449 2.312zm-3.35.81L7.688 3.855l2.009-.485l1.45 2.312zm-3.352.81L4.337 4.665l2.01-.485l1.449 2.312zm-3.349.81L.987 5.475l2.009-.485l1.449 2.312zM1 9.037v8h15v-8zm6.904 6.13v-4.25l3.232 2.169z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clip(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M11.031 15.991h.953c-.504-1.792-2.192-9.907-2.219-10.054c.014-.449.008-1.28.303-1.574c.318-.314.9-.791.9-1.657c0-.81-.389-2.69-3.016-2.69c-2.631 0-2.914 1.881-2.914 2.69c0 .696.24 1.202.801 1.687c.448.441.371.791.371 1.514c0 .319-1.658 8.3-2.211 10.085h.953c.424-1.404 1.984-9.664 1.984-10.085c0-1.025-.107-1.561-.617-2.062c-.402-.349-.367-.657-.367-1.139c0-1.303.547-1.769 2-1.769c1.448 0 2.078.466 2.078 1.769c0 .523-.135.804-.477 1.143c-.412.411-.494 1.401-.518 2.062c-.002.13 1.61 8.667 1.996 10.08"/><path d="M7.746 9.023c-.549 2.39-1.412 5.619-1.775 6.965h4.014c-.33-1.354-1.227-4.612-1.775-6.965zm4.262-.015c.494 2.203 1.328 5.263 1.719 6.98h2.232V9.023zm-12 .029v6.925h2.227c.438-1.701 1.205-4.714 1.716-6.925z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clipboard(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M11.002 1.016v.937h3.014v13.089H2.971V1.953h2.982v-.906H2V16h13V1.016z"/><path d="M9.95 1.5C9.95.672 9.298 0 8.491 0c-.808 0-1.46.672-1.46 1.5c0 .176.035.343.09.5h-.057v1h2.873V2h-.076c.054-.157.089-.324.089-.5m-1.981.516V.959h1.047v1.057zM11.031 3v1.016H5.969V3h-1.94v11h8.951V3z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClipboardChecked(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M12.002 1.016v.937h3.014v13.089H3.971V1.953h2.982v-.906H3V16h13V1.016z"/><path d="M10.95 1.5c0-.828-.652-1.5-1.459-1.5c-.808 0-1.46.672-1.46 1.5c0 .176.035.343.09.5h-.057v1h2.873V2h-.076c.054-.157.089-.324.089-.5m-1.981.516V.959h1.047v1.057zm3.062 2H6.969V3H5.03v11h8.951V3h-1.95zM7.453 8.127l.762-.762l1.414 1.414l2.088-2.09l.811.811l-2.851 2.852z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clock(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M9.008 16.929c-4.385 0-7.95-3.563-7.95-7.942c0-4.381 3.565-7.944 7.95-7.944c4.384 0 7.95 3.563 7.95 7.944c0 4.378-3.566 7.942-7.95 7.942M8.973 2.863c-3.354 0-6.084 2.742-6.084 6.111c0 3.369 2.73 6.111 6.084 6.111c3.355 0 6.085-2.742 6.085-6.111c0-3.369-2.73-6.111-6.085-6.111"/><path d="M8 5h.959v3.978H8zm0 4h3.938v.979H8z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Close(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.053 10.938v-.906h.895v.906zM1.048 9.932V7.037h.905v2.895zm1.005-2.994v-.906h.895v.906zm5.968 4.041V10H9v.979zm1.016-1.038V7.037h.905v2.904zm-1.978 0V7.037h.873v2.904zM8 6.973v-.972h.951v.972zm-3.941 3.958V6.059h.883v3.989h1.031v.884zm10.999.016V6.042h1.9v.905h-1.04V8.1h1.061v.847H15.94v1.106h1.018v.894zm-3.979-.031v-.832h1.863v.832zm1.928-.969v-.921h.936v.921zm-.976-.983v-.943h.932v.943zm-.994-1.006v-.932h.926v.932zm1.042-1.042v-.848h1.863v.848z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloud(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.797 7.445c-.25 0-.487.068-.713.168c-.694-1.036-1.832-1.714-3.119-1.714c-.206 0-.407.022-.605.056c-.682-1.129-1.869-1.88-3.225-1.88c-1.91 0-3.49 1.491-3.779 3.438c-.018 0-.034-.004-.051-.004c-1.281 0-2.318 1.104-2.318 2.467h16c-.139-1.425-1.061-2.531-2.19-2.531"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudCloud(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.037 6.286c.958 0 1.803.543 2.32 1.234a1.42 1.42 0 0 1 .531-.111c.46 0 .869.225 1.166.581h3.887c0-.917-.702-1.661-1.57-1.661c-.011 0-.022.003-.034.003c-.196-1.31-1.265-2.315-2.559-2.315c-.917 0-1.72.507-2.182 1.267a2.533 2.533 0 0 0-.411-.038c-.732 0-1.391.236-1.86.754c.093.103.184.209.261.325a2.93 2.93 0 0 1 .451-.039m3.469 3.124c-.298-.394-.709-.644-1.172-.644c-.188 0-.365.049-.533.123c-.52-.765-1.369-1.264-2.332-1.264c-.155 0-.305.016-.453.041a2.775 2.775 0 0 0-.264-.36c-.525-.628-1.193-1.257-2.051-1.257c-1.43 0-2.709 1.33-2.925 2.768c-.013-.001-.024-.004-.038-.004c-.958 0-1.734 1.14-1.734 2.145h11.969c-.047-.476-.219-1.223-.467-1.548"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudDownload(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M6.982 8.016v1.987h1.002l-1.492 1.966l-1.476-1.966h1.017V8.016z"/><path d="M4.969 6.938h3.062v2.027h7.955C15.847 7.531 14.927 6.42 13.8 6.42c-.249 0-.486.066-.711.167c-.693-1.042-1.826-1.724-3.112-1.724c-.206 0-.407.022-.604.057c-.682-1.137-1.866-1.892-3.219-1.892c-1.906 0-3.484 1.501-3.771 3.46c-.018 0-.035-.006-.051-.006c-1.279 0-2.314 1.111-2.314 2.482h4.951z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudHeavyRain(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.801 5.42c-.249 0-.486.066-.711.167c-.693-1.042-1.826-1.724-3.112-1.724c-.206 0-.407.022-.604.057c-.682-1.137-1.866-1.892-3.219-1.892c-1.906 0-3.484 1.501-3.771 3.46c-.018 0-.035-.006-.051-.006c-1.279 0-2.314 1.111-2.314 2.482h15.969c-.14-1.433-1.06-2.544-2.187-2.544M2.387 13.967a.447.447 0 0 1-.148-.025c-.186-.068-.27-.255-.189-.413l2.208-4.275c.083-.156.301-.228.484-.159c.187.069.271.255.188.413l-2.205 4.273a.377.377 0 0 1-.338.186m3.036-.031a.427.427 0 0 1-.146-.025c-.184-.069-.267-.255-.187-.412l2.172-4.275c.082-.157.295-.229.479-.16c.183.069.266.255.186.412L5.755 13.75a.372.372 0 0 1-.332.186m2.942 0a.458.458 0 0 1-.15-.025c-.188-.068-.273-.252-.19-.408l2.241-4.234c.084-.155.305-.227.493-.158s.274.252.19.408l-2.24 4.233a.388.388 0 0 1-.344.184m3.004 0a.441.441 0 0 1-.147-.025c-.187-.068-.271-.252-.188-.408l2.207-4.234c.08-.155.299-.227.484-.158s.271.252.189.408l-2.207 4.233a.378.378 0 0 1-.338.184"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudPlus(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.916 9.941V7.937H10v-1h3.029v.969h1.033v2.031l2.947.014c-.14-1.43-1.062-2.539-2.189-2.539a1.74 1.74 0 0 0-.714.168c-.694-1.04-1.831-1.72-3.118-1.72c-.207 0-.408.022-.605.056C9.701 4.783 8.514 4.03 7.158 4.03c-1.91 0-3.49 1.497-3.779 3.451c-.018 0-.034-.006-.051-.006c-1.281 0-2.318 1.108-2.318 2.476z"/><path d="M12.969 9h-1.031v-.994h-.876V9H10.01v.959h1.052v.989h.876v-.989h1.031z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudRainHeavyRain(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.801 5.426c-.248 0-.486.066-.711.166c-.693-1.035-1.828-1.713-3.113-1.713c-.207 0-.408.022-.607.055c-.68-1.129-1.865-1.879-3.219-1.879c-1.908 0-3.486 1.49-3.773 3.438c-.019 0-.035-.004-.052-.004c-1.279 0-2.315 1.104-2.315 2.467H15.99c-.14-1.425-1.06-2.53-2.189-2.53M2.004 11.701a.515.515 0 1 1-.929-.448c.124-.257 1.11-1.115 1.11-1.115s-.058 1.306-.181 1.563m3.044-2.109a.516.516 0 0 1-.929-.448c.124-.256 1.11-1.114 1.11-1.114s-.057 1.304-.181 1.562m1.974 2.064a.515.515 0 0 1-.929-.448c.123-.257 1.11-1.114 1.11-1.114s-.057 1.304-.181 1.562M9.92 9.654a.515.515 0 1 1-.929-.448c.124-.258 1.11-1.115 1.11-1.115s-.057 1.306-.181 1.563m-1.918 5.024a.515.515 0 1 1-.929-.448c.124-.257 1.11-1.115 1.11-1.115s-.056 1.306-.181 1.563m-4 0a.515.515 0 1 1-.929-.448c.124-.257 1.11-1.115 1.11-1.115s-.056 1.306-.181 1.563m6.993-2.001a.516.516 0 0 1-.929-.448c.125-.256 1.109-1.114 1.109-1.114s-.056 1.305-.18 1.562m3.023-2.982a.514.514 0 1 1-.928-.448c.124-.258 1.109-1.115 1.109-1.115s-.057 1.306-.181 1.563"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudRemove(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M7.965 10.991V8.987l5.051-.034v2.031l2.994-.033c-.14-1.43-1.062-2.539-2.189-2.539a1.74 1.74 0 0 0-.714.168c-.694-1.04-1.831-1.72-3.118-1.72c-.207 0-.408.022-.605.056C8.702 5.783 7.515 5.03 6.159 5.03c-1.91 0-3.49 1.497-3.779 3.451c-.018 0-.034-.006-.051-.006c-1.281 0-2.318 1.108-2.318 2.476z"/><path d="M9 10h2.959v.99H9z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudSnow(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" transform="translate(1 3)"><path d="M13.803 3.424c-.25 0-.486.067-.711.168c-.693-1.043-1.827-1.725-3.11-1.725c-.207 0-.407.021-.604.056C8.698.785 7.513.03 6.162.03c-1.907 0-3.483 1.501-3.771 3.462c-.018 0-.034-.005-.051-.005c-1.277 0-2.312 1.112-2.312 2.484h15.959c-.138-1.434-1.059-2.547-2.184-2.547"/><circle cx="1.973" cy="7.973" r=".973"/><circle cx="5.962" cy="9.962" r=".962"/><circle cx="9.973" cy="7.973" r=".973"/><circle cx="13.962" cy="9.962" r=".962"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudSun(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.798 10.445c-.25 0-.487.066-.712.167c-.693-1.035-1.827-1.712-3.111-1.712c-.207 0-.408.022-.605.057c-.681-1.129-1.867-1.879-3.219-1.879c-1.906 0-3.483 1.49-3.771 3.436c-.018 0-.034-.006-.051-.006c-1.279 0-2.314 1.104-2.314 2.465h15.969c-.139-1.424-1.06-2.528-2.186-2.528m-3.829-2.271c.201-.035.403-.06.612-.06c1.015 0 1.933.447 2.62 1.171c.475-.558.77-1.283.77-2.088c0-1.747-1.355-3.163-3.025-3.163c-1.415 0-2.594 1.02-2.924 2.394c.806.327 1.492.943 1.947 1.746"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudThunder(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="m11.125 6.915l-.998 2.046h6.891c-.139-1.433-1.061-2.543-2.189-2.543c-.25 0-.487.067-.713.168c-.694-1.041-1.83-1.722-3.118-1.722c-.207 0-.407.021-.606.056c-.682-1.135-1.869-1.889-3.223-1.889c-1.912 0-3.492 1.498-3.781 3.455c-.018 0-.033-.005-.051-.005c-1.281 0-2.318 1.11-2.318 2.479h4.663l.944-2.058z"/><path d="M9.701 10.071H8.59l.648-2.051H7.381l-1.029 2.925l1.336-.013l-.657 3.037z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudUpload(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="m6.531 5.906l2.123 3.059h7.332C15.847 7.531 14.927 6.42 13.8 6.42c-.249 0-.486.066-.711.167c-.693-1.042-1.826-1.724-3.112-1.724c-.206 0-.407.022-.604.057c-.682-1.137-1.866-1.892-3.219-1.892c-1.906 0-3.484 1.501-3.771 3.46c-.018 0-.035-.006-.051-.006c-1.279 0-2.314 1.111-2.314 2.482h4.343z"/><path d="M6.016 11.973V9.984H5.014L6.506 8.02l1.476 1.964H6.965v1.989z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clover(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.538 7.99H9.035V4.461C9.035 2.565 10.655 1 12.55 1a3.432 3.432 0 0 1 3.429 3.435c0 1.897-1.548 3.553-3.441 3.553zm-4.59.024H4.43C2.535 8.014 1 6.333 1 4.435A3.434 3.434 0 0 1 4.429.999c1.894.001 3.52 1.539 3.52 3.436zm.072 4.543c0 1.897-1.743 3.436-3.611 3.436s-3.383-1.538-3.383-3.436s1.49-3.552 3.358-3.552c.956 0 3.636.025 3.636.025zm4.527 3.436c-1.887 0-3.513-1.548-3.513-3.446c0-.943.02-3.545.02-3.545h3.493c1.886 0 3.416 1.657 3.416 3.555c0 1.897-1.53 3.436-3.416 3.436"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CockPot(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.127 11.941H1.883C.529 11.941.016 7.004.016 7.004h9.979c0 .001-.268 4.937-1.868 4.937"/><path d="M15.984 7.486c0 .252-.276.456-.615.456H8.807c-.34 0-.614-.204-.614-.456s.274-.457.614-.457h6.562c.339 0 .615.205.615.457M4 4h1.922v1.297H4z"/><path d="M0 5h9.906v.969H0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cocktail(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M4.92 1.898C4.744.712 3.727 0 2.484 0A2.47 2.47 0 0 0 .01 2.462a2.468 2.468 0 0 0 2.954 2.414L.975 1.888z"/><path d="m9.966 9.001l6.066-5.975l-13-.026l6.042 6v4.031L7.027 15h5.01l-2.045-1.963zM5.188 4l8.594.04l-1.786 1.701c-.47.015-1.142-.087-2.058-.484c-1.219-.442-2.322.405-2.697.743z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Coconut(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.574 6.456c1.594-.299 3.322.218 3.322.218c-.576-1.486-2.635-1.486-2.635-1.486c-.125-1.008 2.635-2.521 2.635-2.521c-2.143-.744-4.035.918-4.035.918C10.531 2.261 12.507 0 12.507 0C9.624.195 8.884 3.294 8.884 3.294C7.565 2.025 4.023 2.133 4.023 2.133c1.483.135 2.802 2.507 2.802 2.507c-3.542 1.229-2.719 3.838-2.719 3.838c.857-.66 1.519-1.026 2.026-1.223c.428-.033 1.34-.117 1.206 1.041C6.489 13.584 8.457 16 8.457 16h3.541c-2.459-2.446-2.722-5.562-2.805-7.315c-.02.002-.005-1.907 2.381-2.229"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Code(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.74 11.993a.714.714 0 0 0 .493-.195l3.479-3.316a.644.644 0 0 0 0-.939l-3.479-3.318a.72.72 0 0 0-.984 0a.643.643 0 0 0 0 .938l2.987 2.85l-2.987 2.848a.643.643 0 0 0 0 .938a.713.713 0 0 0 .491.194m-7.439-.11a.727.727 0 0 1-.497-.19L.287 8.425a.62.62 0 0 1 0-.923l3.517-3.268a.74.74 0 0 1 .995 0a.624.624 0 0 1 0 .924l-3.02 2.805l3.02 2.804a.626.626 0 0 1 0 .926a.736.736 0 0 1-.498.19m2.384 3.006a.672.672 0 0 0 .503-.51L9.934 1.885a.636.636 0 0 0-.488-.768a.661.661 0 0 0-.77.514L5.93 14.125a.636.636 0 0 0 .755.764"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoffeeMachine(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M1.021 15.085v.879H17v-2.928H3.027c-1.109 0-2.006.917-2.006 2.049M6.035 12h.918c1.104 0 2-.926 2-2.067V8.037H4.035v.991h-1v.905h1c0 1.141.897 2.067 2 2.067"/><path d="M1.083.083v1.834c0 1.076.851.947 1.917.991v1.05h2.042v2H6v.917h.959V3.958h1.916V2.917H11v9.067h6V.083zm4.959 3.001H4.917V1.916h1.125zm2 0H6.917V1.916h1.125z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ColorPicker(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.308 4.434c.912-.911.81-2.495-.229-3.534c-1.041-1.039-2.624-1.142-3.534-.23l-1.181 1.183l3.764 3.764zm-9.84 9.842l6.147-6.148l.772.772l1.584-1.586l-5.309-5.309l-1.585 1.584l.773.773l-6.148 6.146l-1.682 4.693l.754.754zm3.448-8.848l1.635 1.636l-6.262 6.26l-2.594.96l.959-2.596z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ColumnDecrease(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 14.977c0 .537.275.974.615.974h13.77c.34 0 .615-.437.615-.974c0-.539-.275-.976-.615-.976H1.615c-.34 0-.615.437-.615.976m11.053-2.719c0 .377.365.686.816.686h1.223c.451 0 .815-.309.815-.686v-3.44c0-.378-.364-.686-.815-.686h-1.223c-.451 0-.816.308-.816.686zm-4.964.001c0 .378.365.685.815.685h1.224c.45 0 .815-.307.815-.685v-4.48c0-.378-.365-.685-.815-.685H7.904c-.45 0-.815.307-.815.685zm-5.027-.027c0 .375.367.68.824.68h1.236c.456 0 .824-.305.824-.68V5.727c0-.375-.368-.68-.824-.68H2.886c-.457 0-.824.305-.824.68zm13.204-5.311a.703.703 0 0 0 .705-.673a.698.698 0 0 0-.682-.719C10.285 5.352 5.521 3.424 2.215.236a.715.715 0 0 0-1.001.012a.69.69 0 0 0 .012.985C4.832 4.71 9.809 6.73 15.24 6.92a.235.235 0 0 0 .026.001"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ColumnIncrease(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16 14.977c0 .537-.275.974-.615.974H1.615c-.34 0-.615-.437-.615-.974c0-.539.275-.976.615-.976h13.77c.34 0 .615.437.615.976M4.947 12.258c0 .377-.366.686-.816.686H2.908c-.451 0-.815-.309-.815-.686v-3.44c0-.378.364-.686.815-.686h1.223c.45 0 .816.308.816.686zm4.964.001c0 .378-.365.685-.815.685H7.872c-.45 0-.815-.307-.815-.685v-4.48c0-.378.365-.685.815-.685h1.224c.45 0 .815.307.815.685zm5.027-.027c0 .375-.367.68-.824.68h-1.236c-.456 0-.824-.305-.824-.68V5.727c0-.375.368-.68.824-.68h1.236c.457 0 .824.305.824.68zM1.734 6.921a.703.703 0 0 1-.705-.673a.698.698 0 0 1 .682-.719c5.004-.177 9.768-2.105 13.074-5.293a.714.714 0 0 1 1 .012a.69.69 0 0 1-.011.985C12.168 4.71 7.191 6.73 1.76 6.92z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ColumnWave(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14 2.057c0-.584.447-1.057.999-1.057S16 1.473 16 2.057v12.719c0 .584-.449 1.057-1.001 1.057c-.552 0-.999-.473-.999-1.057zm-4 3c0-.584.449-1.057.999-1.057C11.552 4 12 4.473 12 5.057v9.719c0 .584-.448 1.057-1.001 1.057c-.55 0-.999-.473-.999-1.057zm-3.998 3C6.002 7.473 6.449 7 7 7c.552 0 1 .473 1 1.057v6.719c0 .584-.448 1.057-1 1.057c-.551 0-.998-.473-.998-1.057zm-3.959 3c0-.584.438-1.057.978-1.057S4 10.473 4 11.057v3.719c0 .584-.439 1.057-.979 1.057s-.978-.473-.978-1.057z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Comb(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m14.563 2.427l2.388 2.389c.283-.975-.717-2.714-2.49-4.486c-1.242-1.244-4.587 1.375-8.217 5.006c-3.629 3.629-6.168 6.892-4.926 8.135c1.738 1.737 3.533 2.775 4.524 2.48l-2.401-2.4c-.324-.324-.337-.834-.031-1.141l1.489-1.488l2.805 2.805c.333.333.825.378 1.103.102c.275-.276.23-.77-.102-1.102L5.899 9.922l2-1.999l2.805 2.804c.333.333.825.378 1.103.102c.275-.276.23-.77-.102-1.102L8.9 6.922l2-1.999l2.805 2.804c.333.333.825.378 1.103.102c.275-.276.23-.77-.102-1.102l-2.805-2.805l1.523-1.523c.306-.308.816-.295 1.139.028"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compass(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M9.021.035c-4.411 0-8 3.588-8 8c0 4.413 3.588 8 8 8c4.411 0 8-3.587 8-8c0-4.412-3.589-8-8-8M9.001 14A6.006 6.006 0 0 1 3 8c0-3.307 2.692-6 6-6c3.31 0 6 2.693 6 6c0 3.308-2.69 6-6 6z"/><path d="m6.042 6.021l2.021 3L12.98 11l-2.979-3.958z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cone(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.287 13.025h-2.013l-.528-2h-9.57l-.528 2H1.692c-.37 0-.671.271-.671.609v.684c0 .337.301.609.671.609h14.596c.37 0 .671-.272.671-.609v-.684c-.001-.337-.302-.609-.672-.609m-5.08-11.994h-4.41l-.759 2.921h5.931zm1.139 4.973H5.658l-.637 2.951h7.963z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CongratulationHat(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M8 13.099L4.125 10.64v3.154s0 .548.355.633c1.114.264 2.279.516 3.52.516c1.27 0 2.485-.191 3.604-.543c.271-.058.271-.551.271-.551V10.64zM1.246 8.066h-.48C.355 9.63.052 14.603.052 14.912c0 .028.007.057.013.084h1.883a.395.395 0 0 0 .011-.084c-.001-.308-.305-5.281-.713-6.846"/><path d="M8 1.062L0 6.043l1.33.999l5.555-1.258c.092-.529.528-.941 1.084-.941a1.125 1.125 0 0 1 0 2.25a1.07 1.07 0 0 1-.783-.352L2.374 7.83L8 11.702l8-5.659z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ConnectOne(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.969 11.082V4.918A2.47 2.47 0 1 0 12 2.5c0 1.213.877 2.217 2.031 2.425v6.15c-.32.058-.618.177-.883.345L5.533 3.848c.254-.388.404-.85.404-1.348a2.47 2.47 0 1 0-2.938 2.422v6.156a2.468 2.468 0 1 0 1 .014V4.908c.301-.066.58-.186.829-.351l7.599 7.556a2.47 2.47 0 1 0 2.542-1.031"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ConnectTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 13.043c-.258 0-.498.069-.71.182l-3.259-3.271c.207-.29.354-.623.422-.986h1.133a1.49 1.49 0 0 0 1.414.992c.829 0 1.5-.652 1.5-1.458c0-.806-.671-1.458-1.5-1.458a1.49 1.49 0 0 0-1.426 1.029h-1.119a2.466 2.466 0 0 0-.382-.954l3.29-3.302c.194.089.408.143.637.143c.829 0 1.5-.652 1.5-1.458c0-.806-.671-1.458-1.5-1.458c-.828 0-1.498.652-1.498 1.458c0 .261.075.503.199.715L9.945 6.486a2.482 2.482 0 0 0-.996-.419V4.886C9.556 4.7 10 4.155 10 3.502c0-.806-.671-1.458-1.5-1.458c-.828 0-1.498.652-1.498 1.458c0 .655.445 1.203 1.058 1.387v1.177c-.364.064-.7.205-.991.408l-3.26-3.271c.118-.209.191-.445.191-.7c0-.806-.671-1.458-1.5-1.458c-.828 0-1.498.652-1.498 1.458c0 .806.67 1.458 1.498 1.458c.234 0 .455-.057.652-.15l3.285 3.297c-.196.285-.33.613-.393.967H4.926a1.49 1.49 0 0 0-1.427-1.029c-.828 0-1.498.652-1.498 1.458c0 .806.67 1.458 1.498 1.458c.66 0 1.215-.416 1.414-.992h1.133c.068.369.221.707.434 1l-3.254 3.266a1.518 1.518 0 0 0-.727-.189c-.828 0-1.498.652-1.498 1.458c0 .806.67 1.458 1.498 1.458c.829 0 1.5-.652 1.5-1.458c0-.226-.057-.437-.15-.626l3.275-3.289c.277.184.594.309.935.369v1.159c-.612.184-1.058.731-1.058 1.387c0 .806.67 1.458 1.498 1.458c.829 0 1.5-.652 1.5-1.458c0-.652-.443-1.198-1.051-1.384v-1.163c.344-.062.663-.19.942-.377l3.269 3.281a1.408 1.408 0 0 0-.158.643c0 .806.67 1.458 1.498 1.458c.829 0 1.5-.652 1.5-1.458c0-.806-.67-1.462-1.499-1.462"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ContactBook(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.061 2.917h1.022v1.215H1.061v7.835h.988v1.117h-.988v2.797h11.897V0H1.061zm3.384 1.475c.03-.024.271-.204.357-.269l-.006-.002l.017-.007c.005-.006.021-.017.024-.02l.004.005c.146-.086.297-.146.469-.172c.462-.068.799.412 1.018.687c.219.272.523.706.486 1.003c-.022.179-.227.344-.421.506l-.007-.01c-.053.057-.287.296-.307.328c-.107.178-.23.581-.058.885c.163.295.485.773.79 1.172c.322.387.717.81.967 1.037c.262.237.693.22.895.162c.039-.01.354-.197.395-.217c.209-.144.42-.296.607-.273c.303.038.656.434.875.707c.219.274.613.712.433 1.132a1.237 1.237 0 0 1-.293.405l.004.004l-.022.016l-.006.008l-.004.001l-.354.269c-.336.217-1.066.488-2.092-.115c-.761-.449-1.596-1.211-2.393-2.163l-.004.003a6.555 6.555 0 0 1-.107-.141c-.037-.046-.076-.089-.114-.137l.004-.003c-.748-.987-1.298-1.965-1.556-2.796c-.347-1.12.102-1.741.399-2.005M0 3h.979v.992H0zm0 9h.977v.943H0zM14 2h.916v2.875H14zm0 9h.887v2.847H14zm0-5h.901v3.895H14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Contacter(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M3 5H2V4h1zm13.038-1.958H2.046c-.556 0-1.004.452-1.004 1.012v8.976c0 .56.448 1.012 1.004 1.012h13.992c.556 0 1.004-.451 1.004-1.012V4c0-.56-.448-.958-1.004-.958M3 13H2v-1h1zm11-2H4V6h10zm2 2h-1v-1h1zm0-8h-1V4h1z"/><path d="M5 7v3h5V7z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Contrast(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.502.115a7.42 7.42 0 1 0 .003 14.839A7.42 7.42 0 0 0 8.502.115m-.584 13.127V1.827c3.477 0 6.292 2.556 6.292 5.707s-2.815 5.708-6.292 5.708"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ControlPad(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M4.277 7.048c-2.344 0-3.242 1.958-3.242 4.372v.166c0 2.416.898 4.373 3.242 4.373h9.475c2.344 0 3.242-1.957 3.242-4.373v-.166c0-2.414-.898-4.372-3.242-4.372zM8 12H6v2.016H5V12H3v-1h2V8.984h1V11h2zm4-1h-1v-1h1zm2 2h-1v-1h1z"/><path d="M9 7.104H8V4l2.031.041L10 1.047h1V5H9z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Corkscrew(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.107.038H5.859C5.386.038 5 .448 5 .953v.085c0 .504.387.915.859.915h7.248c.476 0 .859-.411.859-.915V.953c.001-.505-.383-.915-.859-.915m-2.12 8.458c.012-.146-1.021-1.309-1.021-1.309l-.025-4.232h2.012V2h-4.91v.955h2.014v1.528s-1.015.886-1.017 1.024c-.005.137.981 1.326.981 1.326v4.338s-.951 1.215-.953 1.374c0 .159.957 1.329.957 1.329l.012 1.625c0 .277.215.5.48.5c.268 0 .451-.191.451-.469l-.002-5.636c.001.002 1.01-1.251 1.021-1.398"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CornerFlag(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M1.48 15.845a.43.43 0 0 1-.309-.128a.434.434 0 0 1-.001-.619l4.888-4.888V1.264c0-.241.184-.437.41-.437c.226 0 .408.195.408.437v8.856h9.581c.241 0 .438.172.438.383c0 .21-.196.381-.438.381h-4.815l-.028.198c-.312 2.172-2.19 3.81-4.371 3.81a4.392 4.392 0 0 1-2.994-1.175l-.158-.216l-.224.136l-2.077 2.078a.427.427 0 0 1-.31.13m3.122-2.942l.213.159c.684.62 1.535.955 2.429.955a3.565 3.565 0 0 0 3.474-2.856l.057-.276H6.623z"/><path d="M7.637 5.941c-.957.001-1.51-.868-1.524-1.044l.001-3.795c0-.22.146-.405.344-.452c.173.189.59.555 1.252.555c.495 0 1.021-.206 1.568-.612c.453-.339 1.633-.51 2.077-.51c1.007 0 1.623.851 1.646 1.023L13 4.899c0 .22-.147.404-.345.451a1.78 1.78 0 0 0-1.311-.579c-.515 0-1.783.214-2.345.637c-.47.355-.929.533-1.362.533"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoverFlow(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 4v8h7.894L13 4zm9 0v8h.916V4zm2-2h1v12h-1zM3 4h1v8H3zM1 2h.979v11.918H1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoverFood(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M16.471 12.076H1.49c-.271 0-.489.207-.489.461c0 .256.219.463.489.463h14.981c.271 0 .49-.207.49-.463c0-.254-.22-.461-.49-.461M8.987 4.201c-3.839 0-6.95 3.008-6.95 6.718h13.9c.001-3.71-3.112-6.718-6.95-6.718"/><path d="M8 3h1.969v1.969H8z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cow(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M12.953 15.079c0 .476-1.771.861-3.961.861c-2.188 0-3.962-.386-3.962-.861zM2.046 3.715C1.12 2.824.801.741 1.337.229c.533-.512 1.224.267 2.151 1.155c.927.89 2.643 2.025 2.107 2.537c-.533.514-2.621.682-3.549-.206m13.965-.041c.925-.891 1.244-2.974.708-3.486c-.533-.512-1.223.267-2.15 1.155c-.927.89-2.643 2.025-2.107 2.537c.533.514 2.621.682 3.549-.206M7 6h.984v.984H7zm3 0h.984v.984H10z"/><path d="M14.277 7.383c.21-.499.329-1.037.329-1.598c0-2.633-2.491-4.765-5.565-4.765c-3.076 0-5.566 2.132-5.566 4.765c0 .548.113 1.073.312 1.563c-1.031.399-1.76 1.343-1.76 2.442v.826c0 1.463 1.985 3.338 3.566 3.338h6.805c1.581 0 3.566-1.875 3.566-3.338V9.79c.001-1.073-.693-1.991-1.687-2.407m-6.23 4.664H6.953v-1.094h1.094zM5.841 6.876c0-1.655.509-2.999 1.137-2.999c.629 0 1.138 1.344 1.138 2.999a1.142 1.142 0 0 1-2.275 0m5.206 5.171H9.953v-1.094h1.094zM9.946 6.848c0-1.64.479-2.97 1.067-2.97c.59 0 1.067 1.33 1.067 2.97c.001 1.637-2.134 1.637-2.134 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cpu(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M12.83 2.041H5.254A2.206 2.206 0 0 0 3.042 4.24v7.561c0 1.215.99 2.199 2.212 2.199h7.576a2.204 2.204 0 0 0 2.212-2.199V4.24a2.205 2.205 0 0 0-2.212-2.199M5 12.958H4v-1h1zm0-9H4v-1h1zm6.958 6.07a.973.973 0 0 1-.975.972h-4.05a.973.973 0 0 1-.975-.972V5.973c0-.538.435-.973.976-.973h4.05c.538 0 .974.435.974.973zM14 12.958h-1v-1h1zm0-9h-1v-1h1z"/><path d="M7.72 6a.72.72 0 0 0-.72.722V9.18c0 .398.322.722.72.722h2.436c.396 0 .72-.323.72-.722V6.722a.721.721 0 0 0-.72-.722zM5 0h.871v1H5zm7 0h.871v1H12zM8 0h.871v1H8zM5 15h.959v.896H5zm7 0h.876v.896H12zm-4 0h.918v.896H8zm-7-4h1v.975H1zm0-7h1v.975H1zm0 3h1v.975H1zm15 4h.957v.975H16zm0-4h.957v.975H16zm0-3h.957v.975H16z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cran(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M15.939 7.002h-.916v2.01h-.94v1.968h2.855V9.012h-.999z"/><path d="M16.631 5.024L8.979 1.846v-.148c0-.366-.127-.661-.479-.661c-.353 0-.5.295-.5.661v3.331H5.98v-.243c0-.434-.267-.784-.596-.784H3.622c-.328 0-.596.351-.596.784v.243H1.511a.477.477 0 0 0-.47.486c0 .268.209.484.47.484h1.515v.175c0 .434.268.785.596.785h1.762c.329 0 .596-.352.596-.785V6h1.999v-.057H8v8.078h-.661c-.366 0-.665.339-.665.757v.576h-.002c-.366 0-.664.66-.664 1.037v.582h4.979v-.582c0-.377-.298-1.037-.664-1.037h-.002v-.576c0-.418-.299-.757-.665-.757h-.677V5.943h7.41c.01 0 .018-.006.027-.006a.687.687 0 0 0 .201-.045c.164-.066.374-.173.374-.359c0-.491-.02-.367-.36-.509m-7.652.005V2.875l5.126 2.154z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crop(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M1 3v2.036h1.978V3zm12 11v2.031h1.984V14zm-.01-9.015v4.976h1.979V3.039H7.047v1.946z"/><path d="M6.007 11.041V0H4v12.959h13v-1.918z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CrossHair(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.5 1C4.364 1 1 4.361 1 8.492s3.364 7.492 7.5 7.492c4.137 0 7.5-3.361 7.5-7.492S12.637 1 8.5 1m.469 13.945v-1.898h-.938v1.898A6.462 6.462 0 0 1 2.056 9h1.912V8H2.056a6.461 6.461 0 0 1 5.943-5.943v1.928h.969v-1.93A6.462 6.462 0 0 1 14.943 8h-1.928v.969h1.93a6.46 6.46 0 0 1-5.976 5.976"/><path d="M8.984 6.021h-.941v2.021H6.031v.942h2.012v1.975h.941V8.984h1.954v-.942H8.984z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CrossHairTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.979 10c1.093 0 1.979-.962 1.979-2.043c0-1.08-.886-1.957-1.979-1.957A1.968 1.968 0 0 0 7 7.957C7 9.037 7.886 10 8.979 10M8 0h1.986v4H8zm0 12v4.003h2.007V12zM1 7h3.984v1.969H1zm12 0v1.928h3.99V7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crown(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" transform="translate(1)"><path d="M12.971 15.23c0 .423-.417.768-.932.768H3.951c-.515 0-.932-.345-.932-.768v-.45c0-.422.417-.767.932-.767h8.088c.515 0 .932.345.932.767zm-.506-2.251H3.547L1.421 5.365l1.141-.709l3.011 4.251l1.772-5.865h1.418l1.594 6.22l3.188-4.606l1.046.709z"/><circle cx="14.968" cy="2.968" r=".968"/><circle cx="7.964" cy=".964" r=".964"/><ellipse cx=".983" cy="2.974" rx=".983" ry=".974"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cruise(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.143 12.728c.07.064.234.175.361.168c.08-.002.172-.058.264-.164c.301-.351.635-.538.992-.558c.541-.031.949.342 1.146.521c.27.242.385.263.422.258c.047-.004.426-.419.625-.545V6.059l-2.974.654V3.558c0-.343.289-.62.643-.62H8.02v3.14c.263-.058.694-.058.959 0v-3.14h2.407c.354 0 .642.277.642.62l.015 3.182l-3-.681v6.167c.077-.02.16-.033.25-.035c.267-.002.668.074 1.057.495c.131.138.268.215.34.214c.09-.005.273-.208.287-.222c.047-.047.475-.461 1.053-.447c.244.007.604.091.928.464c.139.159.271.254.367.26c.158.012.393-.204.492-.295c.225-.202.517-.421.875-.472c0 0 .57-2.447.137-3.531c-.23-.576-.285-1.078-1.508-1.501c-.111-.038-.227-.065-.339-.103V3.388c0-.755-.636-1.368-1.419-1.368H9.98V.021h-.978V2.02H7.98V.042h-.978V2.02H5.455c-.783 0-1.421.613-1.421 1.368v3.668c-.109.035-.221.062-.328.1c-1.216.423-1.237.96-1.505 1.52c-.514 1.075.285 3.79.285 3.79c.299-.211.574-.25.772-.237c.486.033.789.376.885.499m11.775 2.019c-.268-.44-.58-.683-.927-.72a.92.92 0 0 0-.253.014c-.367.064-.667.344-.898.599c-.102.114-.34.391-.504.375c-.096-.007-.234-.128-.375-.33c-.333-.473-.701-.581-.951-.589c-.594-.016-1.029.507-1.078.568a.35.35 0 0 0-.041.062c-.043.073-.161.212-.256.219c-.073.002-.213-.098-.344-.273c-.4-.532-.811-.629-1.082-.626a1.14 1.14 0 0 0-.26.043c-.173.05-.32.138-.443.233c-.206.16-.331.336-.355.371c-.018.023-.031.051-.045.08c-.109.169-.197.236-.243.241c-.044.006-.159-.02-.437-.328c-.204-.226-.625-.701-1.186-.663c-.366.027-.711.265-1.023.708c-.096.136-.189.207-.271.211c-.132.006-.299-.134-.372-.215c-.097-.156-.413-.591-.913-.632c-.205-.019-.489.031-.796.299a2.062 2.062 0 0 0-.299.326c-.095.125-.244.268-.374.271c-.16-.002-.33-.189-.375-.256c-.15-.216-.426-.26-.621-.096a.535.535 0 0 0-.09.689c.043.066.443.631 1.06.648c.274.008.684-.096 1.082-.627c.009-.012.017-.02.024-.031c.125-.155.248-.246.322-.242c.092.006.212.146.251.213a.365.365 0 0 0 .045.068c.048.061.484.602 1.078.568c.249-.007.616-.117.948-.589c.145-.202.28-.323.379-.329c.161-.005.399.258.504.373c.262.292.613.617 1.053.617c.031 0 .064-.001.096-.004c.199-.021.385-.117.559-.271c.131-.116.254-.261.367-.449c.007-.009.011-.019.016-.028c.017-.019.037-.042.061-.063c.074-.072.182-.15.285-.152h.003c.138 0 .276.148.37.272c.419.556.822.646 1.096.624c.5-.041.816-.475.914-.633c.055-.061.232-.232.369-.213c.085.001.176.073.273.209c.312.443.658.683 1.023.707c.258.02.489-.077.684-.207c.219-.145.395-.333.504-.453c.275-.309.396-.336.434-.328c.052.005.15.083.275.287c.137.227.414.287.617.135a.528.528 0 0 0 .12-.683"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cubic(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M17 13.316a.667.667 0 0 1-.667.667l-6.667 2A.666.666 0 0 1 9 15.316V8.65c0-.369.298-.667.666-.667l6.667-2c.368 0 .667.298.667.667zm-9.021 2c0 .369-.283.667-.635.667l-5.711-2c-.35 0-.633-.298-.633-.667V6.65c0-.369.283-.667.633-.667l5.711 2c.352 0 .635.298.635.667zM17 4.014c0 .318-.666.538-.666.538L9.666 6.741s-.564.22-1.006.22c-.441 0-.977-.232-.977-.232l-5.996-2.15s-.688-.215-.688-.564s.688-.586.688-.586l5.998-2.15s.621-.232.994-.232c.372 0 .99.222.99.222l6.662 2.188c.001-.002.669.238.669.557"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CupCake(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.807 8.328c-.39 0-.75-.091-1.073-.238c-.595-.271-1.044-.309-1.204-.895c-.246.691-1.008.797-1.996 1.019a4.271 4.271 0 0 1-.956.114c-.335 0-.653-.045-.956-.114c-1.059-.236-1.846-.407-2.029-1.169H4.55c-.12.648-.566.744-1.195 1.041a2.545 2.545 0 0 1-1.07.242a2.54 2.54 0 0 1-.844-.146l1.35 5.045c.104.522.558 1.736 1.52 1.736h6.555c.949 0 1.363-1.214 1.52-1.736l1.353-5.067a2.645 2.645 0 0 1-.932.168m-4.786 5.734H6.946V9.937h1.075zm-2.753-.093l-.872.109l-.605-3.943l.878-.11zm5.24-3.945l.879.11l-.605 3.943l-.873-.109zM7.599.061C3.527.061.174 2.518.174 4.696c0 .039-.024.305-.024.404C.15 6.151 1.159 7 2.403 7c.376 0 .727-.085 1.04-.224c.636-.282 1.087-.809 1.184-1.436c.146.732.933 1.328 1.994 1.554c.294.063.603.105.93.105c.327 0 .636-.042.93-.105c.962-.204 1.7-.711 1.941-1.349c.156.54.592.985 1.171 1.235c.314.135.664.219 1.043.219c1.262 0 2.283-.85 2.283-1.9l.003-.404C14.921 2.518 11.672.061 7.599.061"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Curtain(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M15.963 15.959h-.901V.953L.938.961v14.998H.043V.047h15.92z"/><path d="M1.295 11.253C7.052 9.019 5.889.954 5.889.954h-.873c.188 4.64-.879 7.296-3.041 8.888c.18-.253 2.119-1.623 2.01-8.888h-.953c-.023.995.265 7.181-2.463 9.601c0 0 1.581-2.226 1.369-9.601H.126v15.039h1.836c-.025-2.095-.267-3.503-.961-4.009c0 0 1.998-.048 2.047 4.009h.938c-.002 0 .264-3.946-2.691-4.74m13.422-.023C8.961 9.003 10.123.97 10.123.97H11c-.281 4.576.877 7.269 3.037 8.854c-.178-.253-2.209-1.621-1.96-8.854h.892c.025.994-.252 7.153 2.476 9.565c0 0-1.626-2.218-1.413-9.565h1.932v14.982h-1.912c.027-2.087.271-3.435.965-3.938c0 0-2.029-.104-2.078 3.938h-.906c-.002.001-.455-3.733 2.684-4.722"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CustomerSupport(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M9.005 12.094c-1.442 0-2.907-.856-3.762-1.989c-4.102 0-4.226 5.86-4.226 5.86h15.975s.322-5.886-4.291-5.886c-.854 1.147-2.254 2.015-3.696 2.015m3.436-7.025c0 1.686-1.539 4.91-3.439 4.91c-1.897 0-3.437-3.225-3.437-4.91s1.539-3.053 3.437-3.053c1.9.001 3.439 1.369 3.439 3.053"/><path d="M13.975 3.614c0-.316-.44-.571-.986-.573V2.39c0-.097.037-2.368-3.975-2.368c-4.01 0-3.973 2.271-3.973 2.368v.674c-.01 0-.018-.003-.027-.003c-.543 0-.98.251-.98.562v2.754c0 .309.438.561.98.561s.982-.252.982-.561V3.623c0-.043-.027-.082-.043-.122V2.75c0-.069-.288-1.783 3.061-1.783c3.35 0 3.002 1.714 3.002 1.783v.791c-.006.025-.025.047-.025.073v2.817c0 .317.443.574.992.574c.012 0 .021-.004.033-.004V8.03h-.984V9h1.953z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dashboard(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.521 15.979c-4.111 0-7.459-3.36-7.459-7.489C1.062 4.359 4.41 1 8.521 1s7.458 3.359 7.458 7.49c0 4.129-3.346 7.489-7.458 7.489M8.512 2.956c-3.052 0-5.534 2.486-5.534 5.545c0 3.058 2.482 5.545 5.534 5.545c3.051 0 5.532-2.487 5.532-5.545c0-3.059-2.482-5.545-5.532-5.545"/><path d="M8.965 10.062V7.639c0-.347-.213-.626-.477-.626s-.479.279-.479.626v2.423c-.387.192-.968.542-.968 1.438c0 .896.819 1.484 1.446 1.484s1.428-.589 1.428-1.484s-.56-1.245-.95-1.438M8 4h.959v.976H8zM5 5h.998v.979H5zm6 0h.998v.979H11z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DataArrowDown(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.997 10.04H2.023c-.555 0-1.003.833-1.003 1.859v1.241c0 1.026.448 1.859 1.003 1.859h13.974c.554 0 1.003-.833 1.003-1.859v-1.241c0-1.026-.449-1.859-1.003-1.859M12.5 14.104c-.885 0-1.604-.732-1.604-1.635s.719-1.636 1.604-1.636c.885 0 1.604.732 1.604 1.636c0 .902-.719 1.635-1.604 1.635m-.734-9.94H9.878V1.687c0-.557-.421-.633-.938-.633c-.519 0-.938.076-.938.633v2.477H6.213a.735.735 0 0 0 0 1.033l2.265 2.568a.725.725 0 0 0 1.027 0l2.261-2.568a.73.73 0 0 0 0-1.033"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Database(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" transform="translate(1)"><ellipse cx="6.43" cy="2.421" rx="6.43" ry="2.421"/><path d="M6.463 11.08c-3.516 0-6.447-.98-6.447-1.636v3.269c0 1.188 2.887 2.146 6.447 2.146s6.447-.959 6.447-2.146V9.444c0 .656-2.931 1.636-6.447 1.636"/><path d="M6.494 6.051C2.978 6.051.047 5.049.047 4.377v3.346c0 1.214 2.887 2.196 6.447 2.196s6.447-.982 6.447-2.196V4.377c0 .672-2.931 1.674-6.447 1.674"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabaseDownload(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><g transform="translate(2)"><ellipse cx="6.43" cy="2.421" rx="6.43" ry="2.421"/><path d="M6.494 6.051C2.978 6.051.047 5.049.047 4.377v3.346c0 1.214 2.887 2.196 6.447 2.196s6.447-.982 6.447-2.196V4.377c0 .672-2.931 1.674-6.447 1.674"/></g><path d="M11.014 12.003h2.02V10.5c-1.175.33-2.799.58-4.57.58c-3.516 0-6.447-.98-6.447-1.636v3.269c0 1.188 2.887 2.146 6.447 2.146c1.648 0 3.146-.207 4.287-.546zm3.968-.987v1.987h1.002l-1.492 1.966l-1.476-1.966h1.017v-1.987z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabaseError(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" transform="translate(2)"><ellipse cx="6.43" cy="2.421" rx="6.43" ry="2.421"/><path d="m8.575 11.832l1.014-1.013c-.931.159-1.997.261-3.126.261c-3.516 0-6.447-.98-6.447-1.636v3.269c0 1.188 2.887 2.146 6.447 2.146a17.9 17.9 0 0 0 2.605-.184l1.176-1.176zM6.494 9.919c3.561 0 6.447-.982 6.447-2.196V4.377c0 .672-2.932 1.674-6.447 1.674S.047 5.049.047 4.377v3.346c0 1.214 2.887 2.196 6.447 2.196"/><path d="m14.991 11.799l-.816-.816l-1.702 1.703l-1.668-1.667l-.815.813l1.668 1.668l-1.685 1.686l.814.816l1.687-1.687l1.702 1.703l.815-.814l-1.703-1.702z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabasePlus(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M7.43 5C10.981 5 14 3.837 14 2.5S10.981 0 7.43 0S1 1.163 1 2.5S3.879 5 7.43 5"/><path d="M7.494 9.919c3.561 0 6.447-.982 6.447-2.196V4.377c0 .672-2.932 1.674-6.447 1.674S1.047 5.049 1.047 4.377v3.346c0 1.214 2.887 2.196 6.447 2.196M11 13h4.915v.957H11z"/><path d="M1.016 9.444v3.269c0 1.188 2.887 2.146 6.447 2.146c.869 0 1.697-.059 2.455-.162v-2.739h2.04V10.52c-1.169.321-2.76.561-4.495.561c-3.516-.001-6.447-.981-6.447-1.637M13 11h.958v4.937H13z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabaseRemove(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M7.43 5c3.551 0 6.43-1.242 6.43-2.579C13.86 1.084 10.981 0 7.43 0C3.879 0 1 1.084 1 2.421C1 3.758 3.879 5 7.43 5m6.449 7.917a.716.716 0 0 0 .031-.204V9.444c0 .655-2.932 1.636-6.447 1.636s-6.447-.98-6.447-1.636v3.269c0 1.188 2.887 2.146 6.447 2.146c.885 0 1.728-.06 2.495-.166v-1.776z"/><path d="M7.494 9.919c3.561 0 6.447-.982 6.447-2.196V4.377c0 .672-2.932 1.674-6.447 1.674S1.047 5.049 1.047 4.377v3.346c0 1.214 2.887 2.196 6.447 2.196M11 14h4v.912h-4z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabaseShare(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M.73 6c-.37 0-.668.734-.668 1.54v.977c0 .805.299 1.461.668 1.461H7V14h2V9.978h6.394c.37 0 .606-.656.606-1.461V7.5c.002-.807-.235-1.5-.605-1.5zm13.332 3.031h-1.125V7.969h1.125zM.687 3.972h14.667c.368 0 .667-.665.667-1.485v-.991c0-.82-.299-1.485-.667-1.485H.687C.321.011.021.676.021 1.496v.991c0 .82.299 1.485.666 1.485m12.251-2.003h1v1h-1zM5 15h2v1H5zm-4 0h1.927v.919H1zm8 0h1.958v.938H9zm4 0h1.982v.888H13z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DatabaseUpload(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><g transform="translate(2 1)"><ellipse cx="6.43" cy="2.421" rx="6.43" ry="2.421"/><path d="M6.479 5.988C2.963 5.988.032 4.986.032 4.314V7.66c0 1.214 2.887 2.196 6.447 2.196s6.447-.982 6.447-2.196V4.314c0 .672-2.932 1.674-6.447 1.674"/></g><path d="M11.953 15.031H9.906l2.705-3.488c-1.131.276-2.587.475-4.164.475c-3.516 0-6.447-.98-6.447-1.636v3.269c0 1.188 2.887 2.146 6.447 2.146c1.32 0 2.484-.133 3.506-.359z"/><path d="M13.018 15.969v-1.988h-1.002l1.492-1.965l1.476 1.965h-1.017v1.988z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Delete(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m12.566 8l3.045-3.044c.42-.421.42-1.103 0-1.522L12.566.389a1.078 1.078 0 0 0-1.523 0L7.999 3.433L4.955.389a1.078 1.078 0 0 0-1.523 0L.388 3.434a1.074 1.074 0 0 0-.001 1.522L3.431 8L.387 11.044a1.075 1.075 0 0 0 .001 1.523l3.044 3.044c.42.421 1.102.421 1.523 0l3.044-3.044l3.044 3.044a1.076 1.076 0 0 0 1.523 0l3.045-3.044c.42-.421.42-1.103 0-1.523z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DeliciousCircle(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.916 9.021c0-4.369-3.555-7.91-7.937-7.91c-4.385 0-7.938 3.541-7.938 7.91c0 4.368 3.554 7.909 7.938 7.909c4.382.001 7.937-3.54 7.937-7.909m-8.013.063H2.976c0-3.339 2.734-6.044 6.107-6.044l.025.001v5.865h6.081c0 3.329-2.926 6.204-6.287 6.219z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Deny(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.016.06a7.97 7.97 0 1 0 .002 15.936A7.97 7.97 0 0 0 9.016.06M3.049 8.028a5.974 5.974 0 0 1 5.967-5.966c1.354 0 2.6.458 3.602 1.221l-8.347 8.348a5.931 5.931 0 0 1-1.222-3.603m5.967 5.966a5.921 5.921 0 0 1-3.447-1.105l8.309-8.309a5.93 5.93 0 0 1 1.104 3.448a5.974 5.974 0 0 1-5.966 5.966"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Desktop(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.089.064H1.913A.912.912 0 0 0 1 .975V11.05c0 .504.409.912.913.912h14.176A.911.911 0 0 0 17 11.05V.975a.912.912 0 0 0-.911-.911M2.971 9.369V2.575c0-.375.283-.677.633-.677h10.794c.351 0 .633.302.633.677v6.794c0 .373-.282.676-.633.676H3.604c-.35 0-.633-.303-.633-.676m11.476 4.714H9.911v-2.041H8.052v2.041H3.583c-.282 0-.511.792-.511 1.193c0 .403.229.623.511.623h10.864c.282 0 .511-.22.511-.623c0-.401-.229-1.193-.511-1.193"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DialNumber(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.878 3.639c0 .635-.449 1.152-1.004 1.152h-1.871C11.45 4.791 11 4.273 11 3.639V2.15c0-.635.45-1.15 1.003-1.15h1.871c.555 0 1.004.516 1.004 1.15zm-5 0c0 .635-.449 1.152-1.004 1.152H7.003C6.45 4.791 6 4.273 6 3.639V2.15C6 1.515 6.45 1 7.003 1h1.871c.555 0 1.004.516 1.004 1.15zm-5 0c0 .635-.449 1.152-1.004 1.152H2.003C1.45 4.791 1 4.273 1 3.639V2.15C1 1.515 1.45 1 2.003 1h1.871c.555 0 1.004.516 1.004 1.15zm10 5c0 .635-.449 1.152-1.004 1.152h-1.871C11.45 9.791 11 9.273 11 8.639V7.15c0-.635.45-1.15 1.003-1.15h1.871c.555 0 1.004.516 1.004 1.15zm-5 0c0 .635-.449 1.152-1.004 1.152H7.003C6.45 9.791 6 9.273 6 8.639V7.15C6 6.515 6.45 6 7.003 6h1.871c.555 0 1.004.516 1.004 1.15zm-5 0c0 .635-.449 1.152-1.004 1.152H2.003C1.45 9.791 1 9.273 1 8.639V7.15C1 6.515 1.45 6 2.003 6h1.871c.555 0 1.004.516 1.004 1.15zm10 5c0 .635-.449 1.152-1.004 1.152h-1.871c-.553 0-1.003-.518-1.003-1.152V12.15c0-.635.45-1.15 1.003-1.15h1.871c.555 0 1.004.516 1.004 1.15zm-5 0c0 .635-.449 1.152-1.004 1.152H7.003C6.45 14.791 6 14.273 6 13.639V12.15c0-.635.45-1.15 1.003-1.15h1.871c.555 0 1.004.516 1.004 1.15zm-5 0c0 .635-.449 1.152-1.004 1.152H2.003C1.45 14.791 1 14.273 1 13.639V12.15c0-.635.45-1.15 1.003-1.15h1.871c.555 0 1.004.516 1.004 1.15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DialNumberOne(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" transform="translate(1)"><circle cx="2" cy="3" r="2"/><circle cx="7" cy="3" r="2"/><circle cx="13" cy="2" r="2"/><circle cx="2" cy="8" r="2"/><circle cx="7" cy="8" r="2"/><circle cx="12" cy="8" r="2"/><circle cx="2" cy="13" r="2"/><circle cx="7" cy="13" r="2"/><circle cx="12" cy="13" r="2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Diamond(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.93 15.644L3.34 8.796a1.162 1.162 0 0 1-.002-1.641L7.904.323a1.162 1.162 0 0 1 1.642.002l4.54 6.85a1.16 1.16 0 0 1 .004 1.641l-4.518 6.83s-1.187.451-1.642-.002"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceFive(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.763 0H3.178c-1.18 0-2.135.966-2.135 2.155v11.69c0 1.189.955 2.154 2.135 2.154h11.585a2.145 2.145 0 0 0 2.136-2.154V2.155A2.147 2.147 0 0 0 14.763 0M5.002 6.153c-1.146 0-2.075-.951-2.075-2.123c0-1.172.929-2.123 2.075-2.123c1.146 0 2.076.951 2.076 2.123c0 1.172-.93 2.123-2.076 2.123m8 0c-1.146 0-2.075-.951-2.075-2.123c0-1.172.929-2.123 2.075-2.123c1.146 0 2.076.951 2.076 2.123c0 1.172-.93 2.123-2.076 2.123m0 8c-1.146 0-2.075-.951-2.075-2.123c0-1.172.929-2.123 2.075-2.123c1.146 0 2.076.951 2.076 2.123c0 1.172-.93 2.123-2.076 2.123m-4-4c-1.146 0-2.075-.951-2.075-2.123c0-1.172.929-2.123 2.075-2.123c1.146 0 2.076.951 2.076 2.123c0 1.172-.93 2.123-2.076 2.123m-4 4c-1.146 0-2.075-.951-2.075-2.123c0-1.172.929-2.123 2.075-2.123c1.146 0 2.076.951 2.076 2.123c0 1.172-.93 2.123-2.076 2.123"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceOne(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.763 0H3.178c-1.18 0-2.135.966-2.135 2.155v11.69c0 1.189.955 2.154 2.135 2.154h11.585a2.145 2.145 0 0 0 2.136-2.154V2.155A2.147 2.147 0 0 0 14.763 0M9.017 11.155c-1.696 0-3.069-1.406-3.069-3.14c0-1.734 1.373-3.14 3.069-3.14c1.694 0 3.069 1.406 3.069 3.14c0 1.734-1.375 3.14-3.069 3.14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceSix(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.763 0H3.178c-1.18 0-2.135.966-2.135 2.155v11.69c0 1.189.955 2.154 2.135 2.154h11.585a2.145 2.145 0 0 0 2.136-2.154V2.155A2.147 2.147 0 0 0 14.763 0M6.002 7.153c-1.146 0-2.075-.951-2.075-2.123c0-1.172.929-2.123 2.075-2.123c1.146 0 2.076.951 2.076 2.123c0 1.172-.93 2.123-2.076 2.123m6 0c-1.146 0-2.075-.951-2.075-2.123c0-1.172.929-2.123 2.075-2.123c1.146 0 2.076.951 2.076 2.123c0 1.172-.93 2.123-2.076 2.123m0 6c-1.146 0-2.075-.951-2.075-2.123c0-1.172.929-2.123 2.075-2.123c1.146 0 2.076.951 2.076 2.123c0 1.172-.93 2.123-2.076 2.123m-6 0c-1.146 0-2.075-.951-2.075-2.123c0-1.172.929-2.123 2.075-2.123c1.146 0 2.076.951 2.076 2.123c0 1.172-.93 2.123-2.076 2.123"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceThree(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.763 0H2.178C.998 0 .043.966.043 2.155v11.69c0 1.189.955 2.154 2.135 2.154h11.585a2.145 2.145 0 0 0 2.136-2.154V2.155A2.147 2.147 0 0 0 13.763 0M4.002 6.153c-1.146 0-2.075-.951-2.075-2.123c0-1.172.929-2.123 2.075-2.123c1.146 0 2.076.951 2.076 2.123c0 1.172-.93 2.123-2.076 2.123m8 8c-1.146 0-2.075-.951-2.075-2.123c0-1.172.929-2.123 2.075-2.123c1.146 0 2.076.951 2.076 2.123c0 1.172-.93 2.123-2.076 2.123m-4-4c-1.146 0-2.075-.951-2.075-2.123c0-1.172.929-2.123 2.075-2.123c1.146 0 2.076.951 2.076 2.123c0 1.172-.93 2.123-2.076 2.123"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiceTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.763 0H3.178c-1.18 0-2.135.966-2.135 2.155v11.69c0 1.189.955 2.154 2.135 2.154h11.585a2.145 2.145 0 0 0 2.136-2.154V2.155A2.147 2.147 0 0 0 14.763 0M6.002 10.153c-1.146 0-2.075-.951-2.075-2.123c0-1.172.929-2.123 2.075-2.123c1.146 0 2.076.951 2.076 2.123c0 1.172-.93 2.123-2.076 2.123m6 0c-1.146 0-2.075-.951-2.075-2.123c0-1.172.929-2.123 2.075-2.123c1.146 0 2.076.951 2.076 2.123c0 1.172-.93 2.123-2.076 2.123"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Disc(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" fill-rule="evenodd"><g transform="translate(3 3)"><path d="M7.466 1.066L8.984-.452l.869.87l-1.518 1.517zM.895 10.405l-.868-.87L1.545 8.02l.868.87z"/><ellipse cx="5.115" cy="5.141" rx="3.115" ry="3.141"/><path fill="currentColor" d="M4.988 3.022a1.977 1.977 0 1 0 0 3.953a1.978 1.978 0 0 0 1.979-1.976a1.98 1.98 0 0 0-1.979-1.977"/></g><path fill="currentColor" d="M15.958 7.958C15.958 3.516 12.385 0 7.979 0C3.573 0 0 3.516 0 7.958S3.572 16 7.979 16c4.564 0 7.979-3.69 7.979-8.042m-3.697-5.067l.869.869l-1.518 1.518l-.869-.869zM3.633 13.193l-.868-.869l1.518-1.518l.868.869zm1.252-5.208C4.885 6.25 6.28 4.843 8 4.843c1.719 0 3.115 1.406 3.115 3.142c0 1.733-1.396 3.14-3.115 3.14c-1.721 0-3.115-1.406-3.115-3.14"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiscAdd(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M11 12h5.938v1.969H11z"/><path d="M13 10h1.97v5.938H13zM8.988 6.022a1.977 1.977 0 1 0 0 3.953a1.978 1.978 0 0 0 1.979-1.976a1.98 1.98 0 0 0-1.979-1.977"/><path d="M16.958 7.969c0-4.413-3.573-7.906-7.979-7.906C4.573.063 1 3.556 1 7.969c0 4.412 3.572 7.989 7.979 7.989c.34 0 .674-.026 1.001-.071v-4.935a2.998 2.998 0 0 1-.979.173c-1.721 0-3.115-1.406-3.115-3.14c0-1.735 1.395-3.142 3.115-3.142c1.719 0 3.115 1.406 3.115 3.142c0 .356-.072.695-.182 1.015h4.933c.053-.348.091-.693.091-1.031M6.15 11.676l-1.518 1.518l-.868-.869l1.518-1.518zm6.462-6.399l-.869-.869l1.518-1.518l.869.869z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiscDeny(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M13.477 10.979a2.498 2.498 0 0 0 0 4.994a2.498 2.498 0 0 0 0-4.994m-1.643 2.677c0-1.018.816-1.843 1.82-1.843a1.8 1.8 0 0 1 .82.198l-2.443 2.475a1.853 1.853 0 0 1-.197-.83m1.494 1.504c-.25 0-.492-.055-.723-.159l2.377-2.41a1.826 1.826 0 0 1-1.654 2.569M7.987 6.049c-1.11 0-2.01.875-2.01 1.955c0 1.079.9 1.954 2.01 1.954c1.108 0 2.009-.875 2.009-1.954c0-1.08-.9-1.955-2.009-1.955"/><path d="M9.879 13.499c0-1.989 1.623-3.604 3.623-3.604c.775 0 1.489.247 2.078.659A7.88 7.88 0 0 0 16 8.029c0-4.396-3.579-7.96-7.992-7.96C3.594.069.016 3.633.016 8.029s3.578 7.96 7.992 7.96c.887 0 1.738-.153 2.535-.42a3.564 3.564 0 0 1-.664-2.07m2.544-10.727l.846.845l-1.787 1.788l-.845-.845zM3.548 13.338l-.845-.845l1.788-1.788l.845.846zM5 8a3 3 0 1 1 6 0a3 3 0 0 1-6 0"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiscDownload(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M7.987 6.042c-1.11 0-2.011.88-2.011 1.968c0 1.084.9 1.964 2.011 1.964c1.108 0 2.009-.88 2.009-1.964c0-1.088-.9-1.968-2.009-1.968"/><path d="m10.979 12.963l1.95-.008v-2.018h2.507a7.793 7.793 0 0 0 .565-2.896c0-4.396-3.582-7.958-8-7.958s-8 3.562-8 7.958c0 4.396 3.582 7.959 8 7.959a8.014 8.014 0 0 0 3.333-.73l.667-.301zm1.444-10.279l.846.846l-1.787 1.787l-.845-.845zM3.548 13.25l-.845-.846l1.788-1.787l.845.845zM8 11a3 3 0 1 1 0-6a3 3 0 0 1 0 6"/><path d="m14.521 15.958l1.373-1.937l-.921.003v-1.962h-.944v1.964l-.982.004z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiscError(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="m11.172 14.414l4.198-4.199l1.393 1.393l-4.2 4.199z"/><path d="m11.172 11.586l1.393-1.393l4.198 4.199l-1.393 1.393zM8.988 6.022a1.977 1.977 0 1 0 0 3.953a1.978 1.978 0 0 0 1.979-1.976a1.98 1.98 0 0 0-1.979-1.977"/><path d="M16.958 7.969c0-4.413-3.573-7.906-7.979-7.906C4.573.063 1 3.556 1 7.969c0 4.412 3.572 7.989 7.979 7.989c.34 0 .674-.026 1.001-.071v-4.935a2.998 2.998 0 0 1-.979.173c-1.721 0-3.115-1.406-3.115-3.14c0-1.735 1.395-3.142 3.115-3.142c1.719 0 3.115 1.406 3.115 3.142c0 .356-.072.695-.182 1.015h4.933c.053-.348.091-.693.091-1.031M6.15 11.676l-1.518 1.518l-.868-.869l1.518-1.518zm6.462-6.399l-.869-.869l1.518-1.518l.869.869z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiscPause(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" fill-rule="evenodd"><g transform="translate(3 3)"><path d="M7.466 1.064L8.984-.453l.869.869l-1.518 1.517zm-6.57 9.341l-.869-.87L1.546 8.02l.868.87z"/><ellipse cx="5.115" cy="5.141" rx="3.115" ry="3.141"/><path fill="currentColor" d="M4.988 3.022a1.977 1.977 0 1 0 0 3.953a1.978 1.978 0 0 0 1.979-1.976a1.98 1.98 0 0 0-1.979-1.977"/></g><path fill="currentColor" d="M15.301 10.953c.4-.991.657-1.967.657-2.984c0-4.413-3.573-7.906-7.979-7.906C3.573.063 0 3.556 0 7.969c0 4.412 3.572 7.989 7.979 7.989c.68 0 1.332-.1 1.959-.264v-4.741zm-3.04-8.062l.869.869l-1.518 1.518l-.869-.869zM3.633 13.193l-.868-.869l1.518-1.518l.868.869zM8 5a3 3 0 1 1-.002 6.002A3 3 0 0 1 8 5"/><path fill="currentColor" d="M11 12h2.009v4H11zm3 0h1.981v4H14z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiscPlay(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" fill-rule="evenodd"><g transform="translate(3 3)"><path d="M7.437 1.077L8.954-.441l.87.87l-1.518 1.517zM.883 10.403l-.868-.87l1.518-1.516l.868.87z"/><ellipse cx="5.115" cy="5.141" rx="3.115" ry="3.141"/><path fill="currentColor" d="M4.988 3.022a1.977 1.977 0 1 0 0 3.953a1.978 1.978 0 0 0 1.979-1.976a1.98 1.98 0 0 0-1.979-1.977"/></g><path d="M3.593 11.137h1.229v2.146H3.593z"/><ellipse cx="8.115" cy="8.142" rx="3.115" ry="3.142"/><path d="M11.114 2.744h1.229V4.89h-1.229z"/><path fill="currentColor" d="m11.991 10.911l2.384 1.845a8.022 8.022 0 0 0 1.583-4.787c0-4.413-3.573-7.906-7.979-7.906C3.573.063 0 3.556 0 7.969c0 4.412 3.572 7.989 7.979 7.989a7.976 7.976 0 0 0 4.013-1.068v-3.979zm.27-8.02l.869.869l-1.518 1.518l-.869-.869zM3.633 13.193l-.868-.869l1.518-1.518l.868.869zM8 11.125c-1.721 0-3.115-1.406-3.115-3.14C4.885 6.25 6.28 4.843 8 4.843c1.719 0 3.115 1.406 3.115 3.142c0 1.734-1.396 3.14-3.115 3.14"/><path fill="currentColor" d="M13.031 16v-2.969l1.938 1.5z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiscPlayTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.516 2.984L12.141 7h2.805L8.492.027L2.035 7h2.807zm0 10.035L4.809 9.031H2.003l6.468 6.903l6.468-6.903h-2.806z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiscRemove(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M13 14h3.938v1.969H13z"/><path d="M16.958 7.969c0-4.413-3.573-7.906-7.979-7.906C4.573.063 1 3.556 1 7.969c0 4.412 3.572 7.989 7.979 7.989a7.881 7.881 0 0 0 2.992-.601v-2.402h3.314c.941-1.319 1.673-3.255 1.673-4.986M4.633 13.193l-.868-.869l1.518-1.518l.868.869zM9 11.125c-1.721 0-3.115-1.406-3.115-3.14C5.885 6.25 7.28 4.843 9 4.843c1.719 0 3.115 1.406 3.115 3.142c0 1.734-1.396 3.14-3.115 3.14m3.612-5.848l-.869-.869l1.518-1.518l.869.869z"/><path d="M8.988 6.022a1.977 1.977 0 1 0 0 3.953a1.978 1.978 0 0 0 1.979-1.976a1.98 1.98 0 0 0-1.979-1.977"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiscStop(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" fill-rule="evenodd"><g transform="translate(3 3)"><path d="M7.465 1.063L8.982-.455l.87.87l-1.518 1.517zm-6.57 9.34l-.868-.87l1.518-1.516l.869.87z"/><ellipse cx="5.115" cy="5.141" rx="3.115" ry="3.141"/><path fill="currentColor" d="M4.988 3.022a1.977 1.977 0 1 0 0 3.953a1.978 1.978 0 0 0 1.979-1.976a1.98 1.98 0 0 0-1.979-1.977"/></g><path fill="currentColor" d="M15.301 10.953c.4-.991.657-1.967.657-2.984c0-4.413-3.573-7.906-7.979-7.906C3.573.063 0 3.556 0 7.969c0 4.412 3.572 7.989 7.979 7.989c.68 0 1.332-.1 1.959-.264l1.031-.257l-.031-4.484zm-3.04-8.062l.869.869l-1.518 1.518l-.869-.869zM3.633 13.193l-.868-.869l1.518-1.518l.868.869zM5 8.001a3 3 0 1 1 6 0a3.001 3.001 0 0 1-6 0"/><path fill="currentColor" d="M12 12h4.023v4H12z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiscUpload(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M7.917 6C6.857 6 6 6.863 6 7.93c0 1.065.857 1.928 1.917 1.928A1.921 1.921 0 0 0 9.832 7.93C9.832 6.864 8.974 6 7.917 6"/><path d="m10.979 15.11l3.531-4.607l.69.902a7.979 7.979 0 0 0 .799-3.459c0-4.447-3.582-7.946-8-7.946s-8 3.499-8 7.946c0 4.448 3.582 8.054 8 8.054c1.367 0 2.643-.3 3.766-.89zm1.444-12.419l.846.846l-1.788 1.787l-.844-.845zM3.548 13.258l-.845-.846l1.787-1.787l.846.845zM8 11a3 3 0 1 1 0-6a3 3 0 0 1 0 6"/><path d="m16 13.966l-1.489-1.86l-1.49 1.86h.997v2.003h.97v-2.003z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dish(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.979 4.06a3.933 3.933 0 0 0-3.926 3.929a3.933 3.933 0 0 0 3.926 3.931a3.933 3.933 0 0 0 3.926-3.931A3.933 3.933 0 0 0 8.979 4.06"/><path d="M9.007.072c-4.399 0-7.964 3.562-7.964 7.957a7.96 7.96 0 0 0 7.964 7.96c4.397 0 7.965-3.562 7.965-7.96c0-4.394-3.568-7.957-7.965-7.957m-.028 13.064a5.15 5.15 0 0 1-5.143-5.147a5.15 5.15 0 0 1 5.143-5.146a5.15 5.15 0 0 1 5.143 5.146a5.152 5.152 0 0 1-5.143 5.147"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiskAntenna(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M10.916 8.877c-2.558 2.559-6.535 2.734-8.877.391l9.269-9.27c2.342 2.344 2.169 6.318-.392 8.879M4.625 5.958c-.846-.845-.783-2.274.137-3.194c.92-.92 2.35-.98 3.193-.137z"/><path d="M13.644 15.997H6.378a.363.363 0 0 1-.318-.542l3.69-6.582a.366.366 0 0 1 .319-.186h.002a.363.363 0 0 1 .318.189l3.574 6.582a.364.364 0 0 1-.319.539"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dna(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m8.105 3.551l.05-.07C9.145 2.118 10 .94 10 .042H9c0 .242-.139.576-.352.959H6.334C6.129.628 6 .297 6 .042H5c0 .96.877 2.142 1.89 3.512a.012.012 0 0 0 .002.003C5.903 4.942 5 6.336 5 7.535c0 1.199.908 2.564 1.896 3.915l-.051.07C5.855 12.882 5 14.061 5 14.959h1c0-.247.144-.589.363-.979h2.291c.211.382.346.721.346.979h1c0-.961-.876-2.142-1.89-3.512l-.003-.004C9.097 10.059 10 8.665 10 7.467c0-1.199-.907-2.567-1.895-3.916m-.067-1.593c-.173.251-.356.507-.543.766c-.19-.259-.376-.515-.552-.766zM8.531 6H6.492c.273-.519.63-1.064 1.01-1.608c.389.542.755 1.088 1.029 1.608M9 7.467c0 .166-.035.347-.084.533H6.064A2.028 2.028 0 0 1 6 7.533c0-.172.037-.36.09-.555h2.839c.042.17.071.335.071.489m-2.039 5.575c.174-.251.357-.509.545-.767c.19.259.375.516.551.767zm-.504-4.063h2.061c-.273.524-.635 1.078-1.021 1.629c-.393-.549-.765-1.104-1.04-1.629"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Document(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M2 0v16h13V4.024L9 4V.062zm11 13H4v-1h9zm0-2H4v-1h9zm0-4v1H4V7z"/><path d="M10 0v2.844h4.752z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentArrowDown(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M7.984 0L2.016.052l.01 10.927l2.005-.01l-.004 1.984h2.02L4.031 16h9.938V6.012H7.984z"/><path d="M9 0v4.969h5zM4.017 14.059l-1.5 1.884l-1.5-1.884H2.03V12h.954v2.059z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentArrowLeft(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M10 0v4h4z"/><path d="M9 5V.042H3.033v11.02L4 10.258v1.659L6 12v3.042H4v.937h9.98V5z"/><path d="m2.9 12.004l-.003 1h2.044v.975H2.892l-.004.979l-1.846-1.452z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentArrowRight(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M10 4h4l-4-4z"/><path d="M9 5.042v-5H3v11.403l3.583 3.138l-1.971 1.396H14V5.042z"/><path d="m3.084 13.004l.003 1H1.043v.975h2.05l.004.979l1.845-1.452z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentArrowUp(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M7.984 0L2.016.052l.01 9.927l4.005 5.037H4v.953h1.031V16h8.938V6.012H7.984z"/><path d="M9 0v4.969h5zM1.016 13.953h.993v2.006h.966v-2.006h.978l-1.422-1.937z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentBackward(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M2.017.058v15.896h11.926V5.032H7.952V.058zm9.025 12.524c0 .367-.096.554-.288.554c-.205 0-.349-.255-.377-.306c-.519-.963-2.046-1.779-3.361-1.805v.905l-.039.038a.595.595 0 0 1-.429.155a.642.642 0 0 1-.393-.118l-2.108-1.603a.55.55 0 0 1 0-.756l2.097-1.613a.633.633 0 0 1 .447-.172c.158 0 .293.051.386.146l.039.039V9c1.873.026 4.026 1.582 4.026 3.582"/><path d="M9.004.016v3.978h4.946z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentBulletList(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M10.031.031v3.938h3.871z"/><path d="M9 5V0H2.042v16H14V5zM5 8H4V7h1zm0 2H4V9h1zm-.062 2h-1v-1h1zm1.02 0v-1h5v1zm5-2h-5V9h5zm0-2h-5V7h5z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentChecked(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M9.024.016v3.979h4.947z"/><path d="M8 5.062V0H2.008v12.745l.436.437l2.135-2.136l2.109 2.105l-2.801 2.802H14V5z"/><path d="m4.58 12.46l-2.136 2.136l-1.399-1.4l-.693.693l2.092 2.092l2.829-2.829z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentCopy(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M10.02.062H5.026V13h10.991V2.991h-5.995V.062zM6.998 10.096V8.937l7.02.096v.997zm.014-2.12v-.97h7.02v.97zm7.005-1.945H6.96V4.938h7.056z"/><path d="M11.06.062v1.844h4.752zm-8.997 2V16h10.886v-2.021H3.945V2.062z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentDoc(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M10.031.031v3.938h3.871zM4 9h.965v1.989H4zm4 0h.953v1.953H8z"/><path d="M8.938 5.092V.069H2.042v15.869h11.896V5.092zm-3.905 5.939v1.009H2.982V7.973l2.064.01v.971h.984l.016 2.078zm4.979.016h-.996v.984H7.969v-.984h-.994V8.985h.994V7.969h1.047v1.016h.996zm3.004-2.031h-1v1.969h1V12h-1.062v-.976h-.984V8.957h.984v-.986h1.062z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentEdit(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M7.016.016v3.968h1.263L9.932 2.34z"/><path d="M8.979 12.29c-1.511.585-3.757 1.467-4.021 1.594a1.266 1.266 0 0 1-.562.131a1.42 1.42 0 0 1-1.159-.631a1.386 1.386 0 0 1-.173-1.274c.088-.229.926-2.702 1.435-4.208l.078-.231l2.667-2.654H5.969V.012H.034V15.97h11.942V9.426l-2.793 2.786z"/><path d="m14.682 3.91l-1.535-1.514c-.41-.408-1.218-.577-1.607-.189L5.485 8.234s-1.321 3.913-1.446 4.242c-.099.262.231.59.46.475c.374-.186 4.105-1.631 4.105-1.631l6.071-6.057c.39-.388.418-.946.007-1.353m-6.661 6.944l-2.567.974l-.461-.425l1.225-2.965l.741.546h1.058z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentError(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M11.016 3.984h4.981L11.016.016zm-4.809 7.102L4.5 12.793l-1.707-1.707l-.707.707L3.793 13.5l-1.707 1.707l.707.707L4.5 14.207l1.707 1.707l.707-.707L5.207 13.5l1.707-1.707z"/><path d="M9.969 5.016V.011H4.034v10.4l.466.466l1.908-1.906l2.623 2.623L7.125 13.5l2.47 2.469h6.382V5.016z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentForward(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M2.017.058v15.896h12.006V5.032H7.952V.058zM8.922 9v-.953l.039-.039a.527.527 0 0 1 .386-.146c.174 0 .341.063.447.172l2.097 1.613a.55.55 0 0 1 0 .756l-2.108 1.603a.64.64 0 0 1-.393.118a.6.6 0 0 1-.429-.155l-.039-.038v-.905c-1.315.025-2.843.842-3.361 1.805c-.029.051-.173.306-.377.306c-.192 0-.288-.187-.288-.554c0-2.001 2.153-3.557 4.026-3.583"/><path d="M9.004.016v3.978h4.946z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentHelp(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.969.062H3.031v15.92h10.938V4.062h-5zm.047 12.972H7.954v-1.081h1.062zm1-7.05v1.004l1.001-.004v2.031h-1.001v.996h-1v1.011H7.969V8.968h2.016V7H7.016v1.031H5.969V6.969h1v-.984h3.047z"/><path d="M10.04.076v2.912H14z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentMusic(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M10.016 3.984h4.981L10.016.016z"/><path d="M8.969 5.016V.011H3.034v15.958h11.942V5.016zm3.058 6.642c-.045.379-.566.935-1.07 1.143c-.656.271-1.34.12-1.528-.336c-.188-.456.19-1.046.847-1.316c.231-.096.48-.15.709-.158V8.363l-2.973.699l.004 3.548c-.045.379-.649 1.02-1.153 1.228c-.655.271-1.339.12-1.529-.336c-.188-.456.192-1.046.848-1.316c.28-.116.528-.19.788-.166V8.142l5.062-1.188z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentPdf(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.964 0H2.016v15.958h11.942V5.033H8.964zM6.047 10.036H5.029v.985h-.998v1.025H2.969V7.954l2.061.014v.984h1.018zm4.906-2.083h2.078v1.078H12.02v.922h1.011v1.061H12.02v1.018h-1.067zm-1.92.018v.998h.973L10.02 11H9.016v1.016H6.969V7.961z"/><path d="M10.025.021v3.967h3.954zM4 9h.973v.961H4zm4 0h.969v1.983H8z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentPlus(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M9.927.062L4.073.114l-.021 9.833l1.969.015v2.017H8v3.083H6.062V16h9.907V6.012H9.927z"/><path d="M11 0v4.969h5zM6.979 13H4.992v-2h-.963v2H2.021v.986h2.008v2.008h.963v-2.008h1.987z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentRemove(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M15.969 16V6.012H10V.062L4 0v14l4-.031V16z"/><path d="M11 0v4.969h4.917zM2 15v.937h4.99V15z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentRss(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M7.958.042H2.02v15.911h11.959V4.042H7.958zm-3.979 13.02v-2.305a2.305 2.305 0 0 1 2.306 2.305zm3.615 0c0-2.07-1.57-3.659-3.614-3.659V7.83c2.718 0 5.177 2.48 5.177 5.232zm4.485-.025h-1.562c0-3.686-2.835-6.583-6.507-6.583V4.893c4.367.001 8.069 3.76 8.069 8.144"/><path d="M9.047.036v2.953h4.914z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentSearch(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M10.031.031v3.938h3.871zm-3.539 8.99a2.478 2.478 0 0 0-2.049 3.867l-2.385 2.385l.707.707l2.398-2.398c.385.246.838.394 1.328.394a2.479 2.479 0 0 0 2.477-2.477a2.479 2.479 0 0 0-2.476-2.478m.024 4.078a1.58 1.58 0 0 1-1.578-1.578c0-.87.708-1.578 1.578-1.578a1.58 1.58 0 0 1 1.578 1.578a1.58 1.58 0 0 1-1.578 1.578"/><path d="M9 5V0H2.042v13.876l1.059-1.059a3.602 3.602 0 0 1-.256-1.316a3.625 3.625 0 1 1 3.625 3.625c-.428 0-.832-.088-1.213-.223L4.222 16H14V5z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentStar(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M8 .058L2.008 0v15.954H14V5H7.977C7.964 5 8 .058 8 .058m1.052 9.165h3.056l-2.471 1.806l.943 2.924l-2.471-1.808l-2.473 1.808l.943-2.924l-2.471-1.806h3.056l.943-3.27z"/><path d="M9.024.016v3.979h4.947z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DocumentWarning(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M10.031.031v3.938h3.871z"/><path d="M8.994 0H3.042v16h11V5l-5.048.064zM4.958 13l3.48-7l3.52 7z"/><path d="M8 8v2.015h.977V8zm0 3h.875v1H8z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dog(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.348 6.007c-.026 0-.058.014-.086.017l.031-.03L.025 3.965l-.014.597l1.5 2.46c-.277.341-.473.706-.473.988v4.908H3v-1.655l2.98-1.334h3.019l.667 2.989h1.252v-5.59l-1.007-1.32zm10.404-2.384l-.416-1.385l-2.655 2.622l1.329 1.383l2.81.604l1.192-.872z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Door(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="m4.083.041l6.959 2.667v12.084l-6.709-1.709l-.208.834l7.803 2v-2.042l1.968-.001V.041z"/><path d="M9 8h.875v.875H9z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Downstair(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.969.047h-2.896v2h-2.011v2h-2.01v2H8.043v2H6.031v2h-2.01v2H1v2.918h2.896L17 2.965zM2.029 5.971l4.924.001l-2.154-1.414L7.98 1.377L6.654.051l-3.18 3.181l-1.448-2.205z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownwardsArrowToBar(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.152 11.004a.989.989 0 0 0-.988.99v1.979c0 .547.442.99.988.99h13.855a.99.99 0 0 0 .99-.99v-1.979a.99.99 0 0 0-.99-.99zm7.899-1.443a1.49 1.49 0 0 1-2.104 0L.504 3.116c-.582-.58-.838-2.102 1-2.102h12.988c1.902 0 1.582 1.521 1.002 2.102z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Drill(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.999 4a1 1 0 0 0-1-1h-1a1 1 0 0 0-1-1h-1c0-.553-.446-1-.998-1H5.009C1.019 1 .021 3.447.021 4v3C.021 8.983 1.438 9 2 9l-2 5a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1s-4 0 0-5h.007C6.559 9 9 9 9 7h1a1 1 0 0 0 1-1h1a1 1 0 0 0 1-1h3V4zm-10 0h1v3h-1zm-2 0h1v3h-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DropWater(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 10.333C3 13.463 5.427 16 8.418 16C11.41 16 14 13.463 14 10.333C14 7.204 8.418 0 8.418 0S3 7.204 3 10.333"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dropbox(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M7.98 1.41L2.615 4.36L.042 2.695L4.627.015zm7.999 1.285L12 0L9.016 1.388l4.944 3.139zM6.13 12.667L2 10.184v2.455L7.953 16v-5.675z"/><path d="M2.589 5.333L1 7.723l5.115 2.96l1.838-2.451zM9 10.065V16l5.979-3.167v-2.472l-4.134 2.047z"/><path d="m9.021 8.22l1.864 2.45l5.094-2.726l-1.515-2.624z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Drum(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" transform="translate(1)"><path d="M14.969 5.984c0 .509-.308.922-.687.922H.749c-.379 0-.687-.413-.687-.922c0-.508.308-.922.687-.922h13.533c.379 0 .687.415.687.922m0 9c0 .509-.308.922-.687.922H.749c-.379 0-.687-.413-.687-.922c0-.508.308-.922.687-.922h13.533c.379 0 .687.415.687.922m-1.996-6.922H1.996a.997.997 0 0 0-.996 1v2.844c0 .553.445 1 .996 1h10.977a.997.997 0 0 0 .996-1V9.062c0-.552-.446-1-.996-1"/><ellipse cx="5.439" cy="2.459" rx="1.439" ry="1.459"/><ellipse cx="9.439" cy="2.459" rx="1.439" ry="1.459"/><path d="M5.528 2.957a.49.49 0 0 1-.194-.041L1.251.79C1.033.694.942.457 1.052.262c.106-.195.37-.277.587-.18l4.085 2.126c.218.096.307.333.2.528a.45.45 0 0 1-.396.221m3.918 0a.49.49 0 0 0 .194-.041L13.723.79c.218-.096.308-.333.199-.528c-.107-.195-.37-.277-.587-.18L9.25 2.208c-.218.096-.307.333-.2.528a.45.45 0 0 0 .396.221"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Easal(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.957 11.062V3.107H9.98v-1h-.995V1c0-.553-.445-1-.996-1c-.55 0-.905.447-.905 1v1.107h-.995v1H3.023v7.955H1.085v.875l3.064.062l-.985 2.709c-.185.52-.002 1.047.516 1.232c.519.186.957-.15 1.143-.67l1.115-3.271h1.147v3c0 .553.355 1 .905 1a.997.997 0 0 0 .996-1v-3h1.089l1.064 3.334c.185.52.646.793 1.162.607c.519-.186.702-.691.518-1.213l-.956-2.729l3.033-.062v-.874zM3.95 3.979h2.073v.938h3.894v-.938h2.132v7.084H3.95z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Edit(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="m15.682 2.91l-1.535-1.514c-.41-.408-1.218-.577-1.607-.189L6.485 7.234s-1.321 3.913-1.446 4.242c-.099.262.231.59.46.475c.374-.186 4.105-1.631 4.105-1.631l6.071-6.057c.39-.388.418-.946.007-1.353M9.021 9.854l-2.567.974l-.461-.425l1.225-2.965l.741.546h1.058z"/><path d="M12.042 10.345v3.697H.958V3l8.183-.083l.812-.859H0V15h13l-.057-5.422z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Egg(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.262 7.956a2.411 2.411 0 0 1 2.031-.681c.439.056.395-.274.043-.402c-.893-.329-1.914-.149-2.607.547c-.69.694-.87 1.711-.549 2.604c.127.355.466.422.406-.012c-.099-.741.123-1.5.676-2.056"/><path d="M16.523 6.627c-.439-.798-1.008-1.199-1.881-1.638c-.764-.383-1.546-.753-2.232-1.243c-.637-.456-.727-.873-.982-1.51C10.895.904 8.571.17 7.128.063c-.56-.042-1.104.05-1.661.049C4.104.105 2.934.478 2.52 1.763c-.097.303.02.639.051.945a8.227 8.227 0 0 1 0 1.721c-.076.748-.186 1.602-.512 2.299c-.246.531-.614.926-.758 1.507c-.354 1.413-.404 3.023.521 4.277c.552.748 1.54 1.252 2.509 1.326c1.033.078 1.924.859 2.819 1.255c.81.356 1.6.619 2.482.785c1.552.294 2.979.102 3.919-1.111c.479-.615.663-1.336.883-2.044c.232-.747.725-1.579 1.402-2.079c.658-.482.992-1.155 1.078-1.927c.082-.738-.025-1.422-.391-2.09m-7.529-.721a3.118 3.118 0 0 1 3.131 3.121a3.145 3.145 0 0 1-3.145 3.131a3.115 3.115 0 0 1-3.131-3.121a3.147 3.147 0 0 1 3.145-3.131"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eject(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.152 1.004a.989.989 0 0 0-.988.99v1.979c0 .547.442.99.988.99h13.855a.99.99 0 0 0 .99-.99V1.994a.99.99 0 0 0-.99-.99zm7.899 14.557a1.49 1.49 0 0 1-2.104 0L1.504 9.116c-.582-.58-.838-2.102 1-2.102h12.988c1.902 0 1.582 1.521 1.002 2.102z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Electron(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" transform="translate(1)"><ellipse cx="7.956" cy="7.964" rx="1.956" ry="1.964"/><path d="M11.977 15.997c-2.42 0-5.461-1.544-7.938-4.03C.177 8.09-1.113 3.311 1.103 1.086C1.808.377 2.812.002 4.001.002c2.421 0 5.462 1.545 7.939 4.032c3.861 3.876 5.15 8.655 2.936 10.88c-.705.709-1.708 1.083-2.899 1.083M4.002 1.342c-.841 0-1.5.232-1.956.691c-1.63 1.636-.285 5.752 2.936 8.986c2.235 2.244 4.916 3.639 6.995 3.639c.842 0 1.5-.232 1.957-.691c1.631-1.637.285-5.752-2.936-8.986c-2.235-2.245-4.916-3.639-6.996-3.639"/><path d="M4.002 15.997c-1.189 0-2.193-.374-2.898-1.083C-1.113 12.689.177 7.91 4.04 4.033C6.517 1.547 9.558.002 11.978.002c1.191 0 2.193.375 2.9 1.084c2.214 2.223.924 7.003-2.936 10.881c-2.478 2.486-5.52 4.03-7.94 4.03m7.975-14.656c-2.079 0-4.76 1.395-6.995 3.639c-3.221 3.235-4.565 7.351-2.936 8.987c.456.458 1.115.691 1.956.691c2.08 0 4.76-1.395 6.996-3.639c3.221-3.235 4.564-7.352 2.935-8.987c-.457-.458-1.115-.691-1.956-.691"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ElevatorDown(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M2.334 14.943c-.736 0-1.334-.748-1.334-1.484s.598-1.334 1.334-1.334h2.184l8-5.999h3.148c.736 0 1.334.597 1.334 1.333s-.598 1.521-1.334 1.521h-2.184L5.47 14.943z"/><g transform="translate(4 1)"><ellipse cx="2.488" cy="1.5" rx="1.488" ry="1.5"/><path d="M2.482 4.042c-.789 0-1.43.64-1.43 1.428v1.856l-.904.938a.358.358 0 0 0 .228.632c.081 0 .16-.027.228-.083l3.305-2.937V5.47c0-.159-.031-.312-.081-.456a1.421 1.421 0 0 0-1.346-.972"/></g></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ElevatorUp(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M5.482 13.986H2.334C1.598 13.986 1 13.236 1 12.5s.598-1.44 1.334-1.44h2.143l8.041-6.043h3.148c.736 0 1.334.704 1.334 1.441c0 .735-.598 1.523-1.334 1.523h-2.184z"/><g transform="translate(5 1)"><ellipse cx="1.438" cy="1.479" rx="1.438" ry="1.479"/><path d="M3.924 4.155a.327.327 0 0 0-.471-.045l-.889.771c-.201-.5-.676-.85-1.231-.85c-.736 0-1.334.618-1.334 1.384v2.584L3.88 4.644a.359.359 0 0 0 .044-.489"/></g></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Emoticon(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.981.084a7.939 7.939 0 1 0 0 15.876a7.938 7.938 0 0 0 0-15.876M10.998 4a2 2 0 1 1 .003 3.999A2 2 0 0 1 10.998 4M5 4a2 2 0 1 1-.001 4.001A2 2 0 0 1 5 4m3 10c-3.013 0-5-1.899-5-4h10c0 2.101-1.988 4-5 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EndPage(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m8.327 11.886l4.447-4.94a.65.65 0 0 0-.002-.849l-2.841-.005V1.068c0-.553-.437-1-.976-1H7.004a.987.987 0 0 0-.976 1v5.02l-2.95-.005a.652.652 0 0 0 .004.848l4.485 4.954a.501.501 0 0 0 .76.001m5.591 2.948c0 .552-.437 1-.973 1H3.056c-.537 0-.973-.448-.973-1V14c0-.552.436-1 .973-1h9.889c.536 0 .973.448.973 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Erase(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.932 13.014L3.958 7.039L10.84.158c.322-.325.856-.314 1.191.022l4.762 4.759c.334.336.345.869.021 1.191zm-.969 1.096c-1.582 1.583-5.434.3-7.102-1.368c-1.666-1.667-.52-3.087 1.063-4.67z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Euro(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.629 13.113c-1.027.4-2.15.503-3.252.297a5.872 5.872 0 0 1-3.779-2.45h3.723l.667-1.952H5.757a5.258 5.258 0 0 1-.032-2.067h4.578l.665-1.925H6.501c1.207-1.847 3.516-2.885 5.902-2.438l.452-2.406C8.872-.575 5.034 1.541 3.634 5.017H1.682l-.666 1.925h2.128c-.086.7-.077 1.394.025 2.067H1.697l-.668 1.952h2.705c1.059 2.42 3.317 4.319 6.188 4.856a8.561 8.561 0 0 0 4.7-.43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Excavator(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.858 12.071H5.141c-.61 0-1.102.581-1.102 1.3v1.299c0 .718.491 1.299 1.102 1.299h9.717c.608 0 1.101-.581 1.101-1.299v-1.299c0-.719-.492-1.3-1.101-1.3m-7.879 3.034a1.123 1.123 0 0 1-1.115-1.131c0-.627.5-1.132 1.115-1.132c.613 0 1.113.505 1.113 1.132c0 .626-.5 1.131-1.113 1.131m3.007-.011a1.078 1.078 0 0 1-1.076-1.08c0-.597.482-1.081 1.076-1.081a1.08 1.08 0 0 1 1.076 1.081c0 .597-.483 1.08-1.076 1.08m3.045.032c-.631 0-1.145-.495-1.145-1.108c0-.615.514-1.111 1.145-1.111s1.145.496 1.145 1.111c0 .613-.514 1.108-1.145 1.108m-1.053-7.105l.021-2.014H7.417L6.584 3.91l-.004-.012l-.005.002l-2.586-2.696l.348-.9l-.864-.336l-2.258 6.283a2.368 2.368 0 0 0-.936 3.151c.6 1.146 2.303 1.926 3.44 1.315c.881-.472-.52-3.405-1.589-4.269l1.489-4.286l2.162 2.255l1.281 4.585v.996c0 1.607 1.031.944 1.031.944h6.031s.844.39.844-1.007V8.716c0-.867-.938-.708-.938-.708zm-.947 2.041H7.937V6.937h3.094z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Extinguisher(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.995 13.914v-6.83a3.002 3.002 0 0 0-2.03-2.959V2.969h2.032v-.941h-2.032v.035h-.905c-.005-.027-.022-.051-.025-.078c.06-.584.616-1.043 1.291-1.043h1.672V0h-1.672c-1.188 0-2.168.826-2.285 1.879L9.034 2H7.425c-1.51.064-3.355 1.002-3.385 4.078l-1.011 4.828h2.937L5.04 6.082C5.073 3.195 7.329 3.006 7.446 3h1.583v1.125s-1.932.541-1.932 2.959v6.809s-.467 1.951 2.66 1.951h.605c2.88 0 2.633-1.93 2.633-1.93"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeDrop(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.661 4.051c.234-1.077-.051-2.229-.86-3.046c-1.29-1.299-3.422-1.26-4.76.089C.704 2.443.663 4.589 1.954 5.889c.81.816 1.951 1.103 3.023.868l.602.586l-.58.583l.502.506l.546-.552l5.062 5.073h2.875v-2.766L8.777 5.129l.528-.533l-.5-.505l-.555.559zm5.355 6.371l.005 1.623l-1.804.009l-4.545-4.71l1.547-1.625zm2.969 4.418c0 .619-.44 1.122-.986 1.122c-.545 0-.985-.503-.985-1.122c0-.62.985-1.777.985-1.777s.986 1.157.986 1.777"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EyeGlass(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M1.376 8.967a.387.387 0 0 1-.095-.012a.353.353 0 0 1-.258-.432l2.282-2.842a.358.358 0 0 1 .088-.152l2.749-1.381a.372.372 0 0 1 .514-.014a.346.346 0 0 1 .016.498L3.984 5.946L1.726 8.704a.362.362 0 0 1-.35.263m15.189 0a.354.354 0 0 0 .35-.444l-2.282-2.842a.358.358 0 0 0-.088-.152l-2.749-1.381a.372.372 0 0 0-.514-.014a.347.347 0 0 0-.016.498l2.688 1.314l2.258 2.758c.043.159.19.263.353.263"/><path d="M4.986 12.832c.728 0 2.963-.758 3.274-2.938c.237-.017.473-.034.718-.034c.249 0 .482.024.721.042c.318 2.192 2.497 2.944 3.227 2.944c.786 0 4.043-.77 4.043-3.451c0-2.369-1.636-2.361-3.654-2.361c-1.725 0-3.168.113-3.552 1.523a7.496 7.496 0 0 0-.771-.045a7.26 7.26 0 0 0-.801.056c-.388-1.391-1.816-1.503-3.525-1.503c-2.008 0-3.635-.039-3.635 2.349c0 2.662 3.172 3.418 3.955 3.418"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Factory(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13 2v7L9 6v3L5 6v3L0 6v9h15V2zM3 12H1v-2h2zm4 0H5v-2h2zm4 0H9v-2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Faucet(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.957 14.119c0 1.022-.66 1.851-1.475 1.851c-.814 0-1.475-.828-1.475-1.851c0-1.021 1.475-3.119 1.475-3.119s1.475 2.098 1.475 3.119m-1.325-7.674c-1.205-1.208-2.851-1.362-4.633-1.382v-.014c0-.533-.373-.967-.834-.967H7.979V2.954c.708-.099 1.278-.47 1.514-.958c.262.116.62.189 1.016.189c.801 0 1.449-.294 1.449-.659c0-.364-.648-.659-1.449-.659a2.57 2.57 0 0 0-1.004.185C9.239.466 8.49.041 7.598.041c-.87 0-1.604.402-1.888.964c-.283-.12-.667-.194-1.093-.194c-.871 0-1.575.309-1.575.688c0 .38.704.687 1.575.687c.421 0 .801-.073 1.084-.19c.21.437.69.775 1.299.914v1.173H5.832c-.459 0-.832.434-.832.967v.01H1.436C.702 5.06.02 5.761.02 6.499c0 .737.596 1.464 1.33 1.464h3.653c.007.527.375.953.83.953h3.333c.451 0 .812-.417.828-.934c1.197.049 2.184.218 2.675.709c.277.278.497.648.617 1.268h2.613c-.157-1.405-.518-2.762-1.267-3.514"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Feather(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.247 6.038c1.415-2.604 1.945-4.637 1.625-4.939c-.274-.26-2.54.571-3.638 1.407c-.21.162-1.062 2.605-1.33 2.711c-.229.094-.194-1.478-.065-2.167c-1.065.495-2.854 1.772-3.842 2.657c-.077.07-.276 3.273-.368 3.369c-.188.188-.754-2.169-.912-1.967c-2.264 2.88-3.281 6.092-2.337 7.472C.478 16.033.026 16.973.026 16.973l.695-.027c.42-.719.837-1.394 1.252-2.048l.028.015c.998.098 3.428-.183 5.854-2.453c0 0-1.35-.771-1.13-1.002c.218-.232 2.313-.922 2.532-1.182c.785-.938 1.229-1.374 1.824-2.286c.157-.238-1.631.195-1.486-.04c.141-.228.469-1.387 2.652-1.912"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Female(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.982 9.965c2.243-.47 3.934-2.48 3.934-4.883c0-2.751-2.215-4.988-4.938-4.988c-2.721 0-4.936 2.237-4.936 4.988c0 2.412 1.704 4.43 3.959 4.888v2.073H5.062v1.925h2.939v2.001h1.98v-2.001h2.893v-1.925H9.981zm-4.048-5.01c0-1.768 1.367-3.205 3.045-3.205c1.68 0 3.045 1.438 3.045 3.205c0 1.767-1.365 3.203-3.045 3.203c-1.678 0-3.045-1.436-3.045-3.203"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fence(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.004 6.041h-.972L15.017 3l-1.47-3l-1.579 3l.015 3.041H9.962L9.942 3L8.525 0L7.007 3l.023 3.041H4.97L4.941 3L3.524 0L2.006 3l.023 3.041h-.967l.003.979h.967l.012 5.031L1.039 12v1.01h.993V16h2.935v-2.99H7.03V16h2.938v-2.99h2.015V16h3.052l-.02-2.99h.973V12h-.972V7.02h.973zm-11.02.931h2.058v5.079H4.96zM12.068 12H9.96l-.011-5.072h2.119z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FenceTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.863 6c.582 0 1.055-.294 1.055-.657V4.74c0-.363-.473-.657-1.055-.657h-1.904c0-.553-.301-1-.672-1h-.615c-.371 0-.672.447-.672 1H5.918c0-.553-.289-1-.644-1h-.59c-.355 0-.642.447-.642 1H2c-.553 0-1 .294-1 .657v.603C1 5.706 1.447 6 2 6h2.043v1.041H2c-.553 0-1 .282-1 .629v.576c0 .348.447.629 1 .629h2.043v1.208H2c-.553 0-1 .294-1 .657v.603c0 .363.447.657 1 .657h2.043c0 .553.286 1 .642 1h.59c.354 0 .644-.447.644-1H12c0 .553.301 1 .672 1h.615c.371 0 .672-.447.672-1h1.904c.582 0 1.055-.294 1.055-.657v-.603c0-.363-.473-.657-1.055-.657h-1.904V8.875h1.904c.582 0 1.055-.281 1.055-.629V7.67c0-.347-.473-.629-1.055-.629h-1.904V6zM12 6v1.041H5.918V6zm-6.082 4.083V8.875H12v1.208z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileBox(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M14.993.006H3.072l-2.06 8.755v5.17c0 1.334.472 2.028 1.804 2.028h12.28c1.246 0 1.885-.527 1.885-2.111V8.761zM11.016 10H6.985V8.969h4.031zM2.505 8.01L3.912.981h10.265L15.54 8.01z"/><path d="M5 4h7.947v.968H5zm1-2h5.947v.968H6zM4 6h9.951v.965H4z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileDownload(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M15.031 9.016v4H3v-4H1V16h15.938V9.016z"/><path d="m9.072 9.947l2.91-3.876l-2.014-.021V1.065H8.03V6.05h-2z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileUpload(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M15 9.047v4H3v-4H1V16h15.969V9.047z"/><path d="M8.997 1L6 4.963l2.016.021v4.985h1.937V4.984h2z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Film(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15 1h-.988v1h-2.017V1H4.101v1h-2.2V1H1v16h.901v-1h2.053v1h8.041v-1h2.017v1H15zM4 14H2v-2h2zm0-4H2V8h2zm0-4H2V4h2zm7 9H5v-5h6zm0-7H5V3h6zm3 6h-2v-2h2zm0-4h-2V8h2zm-2-4V4h2v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FilmThirtyFiveMm(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.027 4.038v8.946l6.95-.012V4.079zM6.348 15.969h7.269c.73 0 1.322-.551 1.322-1.229v-.698H5.026v.698c0 .678.591 1.229 1.322 1.229m7.285-14.942h-1.664V.581c0-.312-.234-.565-.524-.565H8.56c-.29 0-.524.254-.524.565v.446H6.365c-.731 0-1.322.542-1.322 1.209v.729h9.913v-.729c-.001-.667-.593-1.209-1.323-1.209M6.984 13.021l.011-9h-.964v1H4.953v-1h-.997c0 .211-.083 1.431-.083 1.431C3.809 7.876.029 8.967.029 10.01v3.012H1v-1.037h1.031v1.037h.938V12H4v1.022h.969v-1.006h1.062v1.006zM3.157 8.435c1.05-.677 1.452-2.409 1.452-2.409L6 6l-.063 5H1c0-.001.452-1.465 2.157-2.565"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Finder(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="m8.198 1l-.126.312c-1.32 3.25-1.764 4.497-1.912 6.143L6.111 8h2.988a.46.46 0 0 1 .444.583c-.256.949-.652 1.614-.759 3.286c-.023.382-.03.762-.02 1.136c.033 1.021.215.983.658 1.714l.146.24H16V1zm1.952 13c-.267-.576-.378-.287-.385-1.163c1.971-.428 2.629-1.334 3.274-2.044a.562.562 0 0 0-.032-.79a.546.546 0 0 0-.778.03c-.442.485-.773 1.23-2.417 1.638c.106-1.262.369-1.618.696-2.825a1.466 1.466 0 0 0-.249-1.271a1.469 1.469 0 0 0-1.16-.574H7.21c.18-1.635.615-2.3 1.664-5H15V14z"/><path d="M6.321 12.028c-.008.306.001.613.023.921c.352.023.709.041 1.075.041c.453 0 .902-.025 1.344-.064a7.469 7.469 0 0 1 .022-.932c-.446.045-.9.077-1.365.077c-.375 0-.741-.015-1.099-.043M11 4h.988v.938H11z"/><path d="M2.167 10.836a.566.566 0 0 1-.052-.791a.547.547 0 0 1 .779-.053c.893.792 1.603 1.697 3.719 1.919c.023-1.45.215-1.831.428-2.827H5.083a1 1 0 0 1-1-1c0-.183.027-2.203 1.864-7.084H0v14h7.021c-.297-.889-.336-.967-.388-1.965c-2.437-.228-3.43-1.279-4.466-2.199M2 4h1v1H2z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FingerUp(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.048 13.172v1.626c0 .629.478 1.141 1.066 1.141h4.788c.59 0 1.066-.512 1.066-1.141v-1.626c0-.631-.477-1.142-1.066-1.142H9.114c-.589 0-1.066.511-1.066 1.142m.94-9.098L8.999.941S8.797.02 8.125.02c-.747 0-1.107.358-1.107 1.002v3.87c0 2.329-.755 1.955-.755 1.955l-1.908-.785c-1.182 0-1.338.27-1.338.885c.399 1.045 4.003 2.857 4.003 2.857c1.133.334 1.68 1.133 1.68 1.133l5.274.062l-.019-.522c0-1.56 1.008-1.745 1.008-2.389V4.315c0-.643-.512-1.294-1.145-1.294h-.798v1h-1.061V2.625a.544.544 0 0 0-.058-.24c-.098-.209-.625-.367-.882-.367c-.331 0-.985.254-1.011.562l-.01 1.441z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fire(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.849 9.245c-.417-.518-2.513-2.1-1.658-4.219c0 0-1.908.725-1.149 2.953c-.35-.496-3.318-1.616-3.04-7.85c0 0-4.912 2.954-4.912 7.059c0 0 0 3.256 1.974 4.799c0 0-2.616-.43-3.317-4.557c-.176.494-4.756 5.288 1.74 8.446l3.489.083h3.351s7.372-2.493 3.522-6.714"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FireAlarm(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M10.986 9.977a4.936 4.936 0 0 0 4.93-4.941a4.936 4.936 0 0 0-4.93-4.941a4.936 4.936 0 0 0-4.931 4.941a4.937 4.937 0 0 0 4.931 4.941m0-7c1.135 0 2.055.921 2.055 2.059a2.056 2.056 0 1 1-4.109 0c0-1.138.92-2.059 2.054-2.059"/><path d="M14.283 11.018H7.681a.68.68 0 0 0-.678.684v.646c-1.326-.998-2.27-3.02-2.568-4.748c.323-.264.533-.656.533-1.1c0-.794-.658-1.438-1.469-1.438c-.812 0-1.469.644-1.469 1.438c0 .67.471 1.228 1.104 1.387c.43 2.389 1.877 4.898 3.869 6.052v1.316a.68.68 0 0 0 .678.683h6.602a.68.68 0 0 0 .679-.683v-3.554a.68.68 0 0 0-.679-.683"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FireHydrant(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 14.047h-.326c-.371 0-.674.457-.674 1.019v.893h7.987v-.893c0-.563-.303-1.019-.674-1.019h-.357V6H5zM4 4v.942L12 5V4zm3-2.858c-1.148.27-1.969.993-1.969 1.842h5.938c0-.854-.83-1.58-1.988-1.847V0H7zM3 8V7h.947v1h.027v.97h-.027v.938H3V8.97H2V8zm9.953.938v.968H12V7h.953v1h.989v.938z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FireWood(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.884 10.689c.231.518.033 1.064-.442 1.223L3.376 15.915c-.478.158-1.05-.133-1.282-.65c-.231-.517-.034-1.064.442-1.222l12.066-4.004c.476-.158 1.048.133 1.282.65m-1.282 5.226c.477.158 1.049-.133 1.282-.65c.231-.517.033-1.064-.442-1.222l-1.59-.527l-3.242 1.075zM3.375 10.039c-.478-.158-1.05.133-1.282.65c-.231.518-.034 1.064.442 1.223l1.592.527l3.24-1.076zm7.143-.086s3.934-1.43 1.881-3.851c-.604-.773-1.012-1.8-.557-3.016c0 0-1.348 1.265-.943 2.542c-.186-.283-2.047-1.012-1.898-4.587c0 0-2.256 1.778-2.256 3.956c0 0 .213 2.512 1.439 2.904c0 0-1.872-.474-2.248-2.842c-.092.285-2.409 2.859.764 4.818c0 0-1.169-.784-.599-2.263c1.153 1.95 3.46.563 3.46.563s-2.393-.551-1.092-3.673c.299 1.619 1.738 2.479 2.651 2.622c-.48-.966.348-.223.705.381c.47 1.17-1.307 2.446-1.307 2.446"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FirstAidBriefcase(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.82 3H3.18C1.975 3 1 3.99 1 5.209v7.58C1 14.01 1.975 15 3.18 15h11.64c1.205 0 2.18-.99 2.18-2.21V5.208C17.001 3.989 16.026 3 14.82 3m-4.785 7.016v1.953H8v-1.953H5.987V8H8V5.969h2.035V8h1.952v2.016zm1.937-8.047c0-1.038-.839-1.922-1.867-1.922h-2.19c-1.03 0-1.914.884-1.914 1.922l.971-.007c0-.598.363-1.026.956-1.026h2.167c.594 0 .984.429.984 1.026z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fish(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.537 7.348c-.65-1.275-3.148-2.226-6.143-2.226c-3.486 0-6.312 1.289-6.312 2.879s2.825 2.879 6.312 2.879c2.994 0 5.492-.951 6.143-2.226c.659.841 1.953 2.191 3.437 2.191V5.156c-1.484 0-2.778 1.351-3.437 2.192m-10.494.777H1V6.917h1.043z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flag(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 0h1.994v15.913H1zm3.056.52v7.575S5.667 6.664 9.244 8c3.576 1.338 4.305.974 5.712.742c0 0-2.048-.871-3.222-4.029c0 0 2.987-2.755 3.222-4.274c0 0-3.7 1.212-5.751.241C7.152-.293 4.994-.089 4.056.52"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Float(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9 .062c-4.398 0-7.979 3.576-7.979 7.969c0 4.395 3.581 7.97 7.979 7.97c4.4 0 7.979-3.575 7.979-7.97C16.979 3.638 13.4.062 9 .062M15.947 7h-2.133a4.958 4.958 0 0 0-4.229-3.896h.413v-2.05A7.03 7.03 0 0 1 15.947 7M4.994 8.021a3.972 3.972 0 0 1 3.965-3.979a3.972 3.972 0 0 1 3.965 3.979A3.972 3.972 0 0 1 8.959 12a3.972 3.972 0 0 1-3.965-3.979M8 1.055v2.05h.332a4.958 4.958 0 0 0-4.229 3.896H2.054A7.028 7.028 0 0 1 8 1.055M2.053 8.979h2.042A4.969 4.969 0 0 0 8 12.902v2.045a7.028 7.028 0 0 1-5.947-5.968m7.946 5.968v-2.055h-.021a4.973 4.973 0 0 0 3.845-3.913h2.126a7.029 7.029 0 0 1-5.95 5.968"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FloppyDisk(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M12.961.031H5.003v5.952h7.958zM12 5h-2V1h2z"/><path d="M14.988.031h-.909v7.07H3.941V.031H2.99c-1.105 0-2 .92-2 2.052v10.895c0 1.133.896 2.053 2 2.053h12c1.106 0 2-.92 2-2.053V2.083zM12.968 13h-8v-1h8zm0-2h-8v-1h8z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flower(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.164 8.404c-.171-.094-.522-.237-.982-.403c.46-.166.81-.308.981-.402c1.016-.549 1.121-2.16.235-3.598c-.887-1.437-2.429-2.157-3.444-1.609c-.173.095-.48.306-.865.597c.078-.458.119-.814.119-1c0-1.1-1.437-1.99-3.208-1.99c-1.771 0-3.21.89-3.21 1.99c0 .186.043.541.121.998a8.562 8.562 0 0 0-.863-.596C3.032 1.842 1.488 2.562.603 4c-.885 1.437-.779 3.048.235 3.597c.173.095.522.238.985.404c-.463.167-.814.31-.986.404c-1.016.549-1.121 2.158-.235 3.597c.886 1.437 2.429 2.157 3.444 1.608c.173-.093.479-.304.865-.595c-.078.457-.121.81-.121.997c0 1.099 1.436 1.989 3.21 1.989c1.771 0 3.208-.89 3.208-1.989c0-.187-.041-.542-.119-1.001c.385.293.693.505.866.598c1.016.549 2.558-.17 3.443-1.607c.887-1.439.781-3.049-.234-3.598M8 11.047a3.047 3.047 0 1 1 0-6.092a3.048 3.048 0 0 1 3.049 3.046A3.05 3.05 0 0 1 8 11.047"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlowerPot(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.25 5.965H.663L0 3.083h15.914zm-2.475 8.971H3.139L1.055 7.022h13.804z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folder(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.35 5L6.788 2.042H3.021v1.021H1.023v9.913h1.02l.002 1h14.902L16.968 5z"/><path d="M14.964 3.982h-5.61l-.33-.94h5.94z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderContact(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M7 8v1.957h2.014V8zm3-1h1v3h-1z"/><path d="M7.35 4L5.788 2.042H2.021v1.021H.023v9.913h1.02l.002 1h14.902L15.968 4zm3.643 6l-.018 1L7 11.067v-.99l-1-.021V7.972h1v-.98l3-.01V6H6v1h-.984v3.964h1v1.007h3.957V13H6v-1l-1 .058v-1.021H3.953V7.034h1.008V6.025H6V5h3.958v1.011H11v.981h.953V10z"/><path d="M12.958 2.979h-4.27l-.667-.958h4.937z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderError(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.35 5L6.788 2.042H3.021v1.021H1.023v9.913h1.02l.002 1h14.902L16.968 5zm3.568 6.109l-.828.829l-1.578-1.577l-1.578 1.577l-.83-.829l1.578-1.578l-1.578-1.578l.83-.828l1.578 1.576l1.576-1.576l.83.828l-1.578 1.578z"/><path d="M14.964 3.982v-.94h-5.94l.33.94z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderMusic(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M7.35 5L5.788 2.042H2.021v1.021H.023v9.913h1.02l.002 1h14.902L15.968 5zm2.495 6.959a1.977 1.977 0 0 1-.752.155c-.499 0-.9-.221-1.048-.576c-.225-.551.214-1.255.978-1.571c.287-.118.595-.184.864-.184l.056.001V7.318l-2.884.781l.001 3.521c-.054.468-.714 1.2-1.284 1.438a1.979 1.979 0 0 1-.749.153c-.5 0-.902-.221-1.051-.576c-.224-.551.217-1.255.981-1.571a2.045 2.045 0 0 1 .985-.143V7.215l5.116-1.279l.016 4.683c-.053.463-.64 1.097-1.229 1.34"/><path d="M13.964 3.982v-.94h-5.94l.33.94z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderOpen(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m8.03 4.042l-.802-1H1.015v1H.009V13h1.017l.005.984l13.619-.021l1.303-9.922H8.03zm5.79 8.999H1.711L3.1 4.953h11.932z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderPlus(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M5.787 2.042H2.02v1.021H.022v9.913h1.02l.002 1h14.902L15.967 5H7.349zm2.192 4.937h1.062V9h2.021v1.062H9.041v2.021H7.979v-2.021H5.958V9h2.021z"/><path d="M13.964 3.982v-.94h-5.94l.33.94z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderRemove(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.35 5L6.788 2.042H3.021v1.021H1.023v9.913h1.02l.002 1h14.902L16.968 5zm2.733 5.042H6.979V8.917h4.104z"/><path d="M14.964 3.982v-.94h-5.94l.33.94z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderRss(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.35 5L6.788 2.042H3.021v1.021H1.023v9.913h1.021l.001 1h14.902L16.968 5zm-.327 7.042h-1v-1h.17a.83.83 0 0 1 .83.83zm.97-.06c0-.94-1.018-2.023-1.966-2.023v-.95c1.54 0 2.948 1.445 2.948 2.973zm2.137.013c0-2.476-1.654-4.115-4.12-4.115v-.838c3.299 0 4.932 1.636 4.932 4.954z"/><path d="M14.964 3.982v-.94h-5.94l.33.94z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderSearch(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M13.959 1V.043H6.043l1 .957zM4.002 8.486a3.485 3.485 0 0 0 3.49 3.48a3.485 3.485 0 0 0 3.488-3.48a3.484 3.484 0 0 0-3.488-3.479a3.485 3.485 0 0 0-3.49 3.479m6.007.014c0 1.414-1.138 2.562-2.538 2.562c-1.402 0-2.539-1.147-2.539-2.562c0-1.416 1.137-2.564 2.539-2.564c1.4 0 2.538 1.148 2.538 2.564"/><path d="M3.652 0H1v1H0v12h1.016l.002 1h9.518l-1.299-1.279a4.627 4.627 0 0 1-6.4-4.263a4.625 4.625 0 0 1 4.627-4.616c2.551 0 4.629 2.07 4.629 4.616c0 .591-.123 1.151-.326 1.672L15.484 14h.42l.033-12H6z"/><path d="m14.021 14.029l-.937.979l-2.582-2.541l1.019-1.041z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderShare(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.413 2.034L8.575 0H2.02v2.034h-.979l.021 8.924h.978l-.019 1h7.021L9.021 13h.937l.021-1.042h5.664l1.397-9.924zm5.366 9.069H2.837l1.428-8.127h11.756zM1 15h1.979v.956H1zm4 0h1.979v.956H5zm7 0h1.979v.956H12zm3 0h1.979v.956H15zm-5 0v-1H9v1H8v1h3v-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FolderWarning(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.35 5L6.788 2.042H3.021v1.021H1.023v9.913h1.021l.001 1h14.902L16.968 5zm-3.329 8.021l4.5-7.041l4.459 7.041z"/><path d="M14.964 3.982v-.94h-5.94l.33.94zM9 8h1v1.956H9zm0 3h1v.973H9z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FootSign(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m5.428 13.572l-3.799-1.43l-.484 1.173s-.463 1.898 1.197 2.523c1.662.627 2.619-1.087 2.619-1.087zM8.207 3.193C6.565 2.534 4.26 3.979 3.463 5.8c-.328.75-.477 1.559-.601 2.357c-.059.381-.101.772-.176 1.152c-.096.477-.239.936-.381 1.399c-.197.643.02.751.619.861l2.098.792c.214.102.632.359.847.119c.19-.215.152-.607.279-.862c.164-.331.348-.668.551-.98C7.148 9.94 8 9.43 8.577 8.828c.715-.748 1.11-1.498 1.328-2.49c.29-1.321-.377-2.616-1.698-3.145m1.77 6.096l3.655 1.443s-.754 2.953-2.986 2.094c-2.232-.861-.669-3.537-.669-3.537M15.19.217c1.56.58 2.188 3.038 1.531 4.807c-.271.727-.703 1.379-1.146 2.014c-.212.302-.442.598-.643.911c-.247.39-.453.801-.661 1.211c-.291.57-.512.505-.996.206l-1.965-.773c-.213-.064-.669-.152-.657-.451c.011-.27.296-.52.378-.774c.106-.333.203-.685.269-1.028c.152-.768-.096-1.661-.092-2.441c0-.969.227-1.74.732-2.565c.677-1.106 1.992-1.582 3.25-1.117"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Forcus(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.999 4.875V3.01A1.01 1.01 0 0 1 4.01 1.999h1.908V.082H4.066a2.987 2.987 0 0 0-2.983 2.984v1.809zm9.084-2.876h1.948c.558 0 1.012.453 1.012 1.011v1.865h1.875V3.066A2.988 2.988 0 0 0 13.933.082h-1.85zM5.918 14.042H4.01a1.01 1.01 0 0 1-1.011-1.011v-1.948H1.083v1.85a2.987 2.987 0 0 0 2.983 2.984h1.852zm9.125-2.959v1.948c0 .558-.454 1.011-1.012 1.011h-1.948v1.875h1.85a2.988 2.988 0 0 0 2.985-2.984v-1.85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Forklift(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" transform="translate(0 2)"><circle cx="7.953" cy="8.953" r="1.953"/><circle cx="13.477" cy="9.477" r="1.477"/><path d="M13.477 6.925c1.23 0 2.252.882 2.475 2.044h.002l.047-7.024c.002-1.017-.91-1.917-1.943-1.919L9.656.015C8.625.013 8.142.734 7.779 1.852L6.674 4.476c0 .144-.719.25-.913.39c-.451.318-.743.78-.745 1.447c-.001.525.002 1.658.002 1.658l.168-.004a2.992 2.992 0 0 1 2.816-1.993c1.656 0 3 1.342 3 2.995l-.001.012H11a2.523 2.523 0 0 1 2.477-2.056m-1.49-5.091c.007-.931.146-.855.98-.853l1.042.002c.717.002.975.09.973.859l-.009 2.137l-2.991-.008l.006-2.137zM10.95 4l-2.989-.006l.653-2.08S8.887 1 9.716 1c.616 0 1.245.004 1.245.004L10.951 4zM3.033 2.375l-.018 6.656L.3 9.032c-.366-.001-.369.964-.002.965H3.35c.367.001.602-.297.604-.666l.019-6.954c.001-.367-.939-.37-.939-.002z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ForwardPage(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M15.798 3.516L11.892.266c-.277-.288-.593-.335-.871-.048c0 0 .037 2.19.006 2.19c-.68 0-1.083.113-1.7.314A6.329 6.329 0 0 0 7.119 3.99C5.843 5.113 5.026 6.392 5.026 8.203c0 1.182.523.752.742.339c1.035-1.945 2.923-2.928 5.281-2.928c.03 0-.001 2.146-.001 2.146c.278.287.628.253.869.047l3.881-3.25a.754.754 0 0 0 0-1.041"/><path d="M14 8.079v4.983H.938V3.937h5.215c.276-.312.621-.717 1.223-.919H.008v10.964H14.97V7.258z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fridge(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 .003c-.553 0-1 .444-1 .992V3.99c0 .548.447.992 1 .992h6c.553 0 1-.444 1-.992V.995a.996.996 0 0 0-1-.992zm6.042 2.038H9.958V.958h1.084zM5 5.99c-.553 0-1 .451-1 1.01v6.894c0 .558.447 1.01 1 1.01L5.021 16H6.98v-1.097h2.041V16h1.959v-1.097c.573 0 1.02-.452 1.02-1.01V6.999c0-.558-.447-1.009-1-1.009zm6.042 3.051H9.958V6.875h1.084z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fullscreen(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 5h12v8H3zm.918 9.938H1v-2.876h1v1.98h1.918zm13.082 0h-2.938v-.896H16v-1.984h1zm0-9.021h-1v-1.95h-1.943v-.946H17zM2 5.938H1V3h2.938v.938H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GameControll(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M6.352 7.083S6.021 6.041 4.35 6.041c-1.671 0-2.006 2.006-2.006 2.006l-1.336 5.974c0 1.106.835 2.005 2.005 2.005c1.039 0 1.859-1.714 2.304-2.902c.23.143.498.229.787.229c.773 0 1.533-.715 1.617-1.465h2.75c.084.75.584 1.465 1.356 1.465c.315 0 .608-.098.85-.264c.437 1.188 1.247 2.938 2.241 2.938c1.505 0 2.006-.898 2.006-2.005l-1.336-5.974s-.251-2.006-2.006-2.006s-1.881 1.042-1.881 1.042z"/><path d="M10.041 7.459h-1V5.307c-.01-.007-.161-.927.4-1.63c.37-.462.936-.696 1.683-.696c1.63 0 1.876-.436 1.876-1.938h1c0 1.835-.48 2.971-2.876 2.971c-.431 0-.737.107-.907.318c-.272.335-.184.875-.184.875z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GameOne(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m8.931 7.986l4.813-6.343A7.875 7.875 0 0 0 8.931-.001C4.551-.001 1 3.561 1 7.956c0 4.396 3.551 7.959 7.931 7.959c3.517 0 6.495-2.3 7.533-5.481zM7.499 5a1.5 1.5 0 1 1 .003-3.001A1.5 1.5 0 0 1 7.499 5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GasStation(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.475.07h-1.486l-1.988 2.335v8.927c0 .751-1 .793-1 .043L12 8.148c.005-.094.027-.934-.582-1.574c-.368-.385-.822-.539-1.446-.585V.053L2 0v14.018h-.001a.012.012 0 0 0-.008.003H1v1.938h9.979v-1.938l-.991.054c-.006 0-.009-.003-.009-.003H9.97V7.006c.323.039.543.073.715.253c.327.336.316.842.314.867v3.25c0 1.082.729 1.604 1.562 1.604c.875 0 1.438-.604 1.438-1.647V2.959h.984v-.984l.373-1.023h1.037c.268 0 .563-.134.563-.399a.482.482 0 0 0-.481-.483M2.958.939h6.073v5.092H2.958zm6.125 6.977V9H6.958V7.916zM2.958 8H5.02v1H2.958z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gate(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.589 7.076h-1.801l-.622-3.123h1.218c.278 0 .522-.152.522-.325V2.326c0-.173-.244-.325-.522-.325H2.538c-.278 0-.522.152-.522.325v1.302c0 .173.244.325.522.325h1.33l-.605 3.123H1.411c-.228 0-.411.141-.411.314v1.26c0 .174.184.314.411.314h1.486l-.831 4.288a.336.336 0 0 0 .245.407l1.301.324c.026.006.056.01.082.01c.154 0 .288-.105.332-.276l.921-4.753h7.163l.923 4.634a.336.336 0 0 0 .406.245l1.303-.324c.18-.045.289-.227.251-.385l-.83-4.17h1.426c.228 0 .411-.141.411-.314V7.39c0-.173-.184-.314-.411-.314m-10.276 0l.605-3.123h5.194l.622 3.123z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gear(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.969 8.969V7.011h-2.07a5.943 5.943 0 0 0-1.019-2.465l1.463-1.463l-1.414-1.414l-1.463 1.463a5.954 5.954 0 0 0-2.465-1.021V.042h-2v2.069a5.94 5.94 0 0 0-2.465 1.021L3.073 1.669L1.688 3.053l1.459 1.459a5.95 5.95 0 0 0-1.046 2.499H.032v1.958h2.063a5.943 5.943 0 0 0 1.026 2.507l-1.463 1.463l1.414 1.414l1.463-1.463c.72.513 1.557.867 2.465 1.021v2.069h2v-2.069a5.955 5.955 0 0 0 2.499-1.046l1.458 1.458l1.385-1.384l-1.463-1.463a5.95 5.95 0 0 0 1.026-2.507zm-7.948 2.093a3.084 3.084 0 0 1 0-6.166a3.084 3.084 0 0 1 0 6.166"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GearOne(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.887 10.025a3.39 3.39 0 0 0-.658-1.55l.766-.765l-.688-.687l-.771.771a3.434 3.434 0 0 0-1.542-.633V6.066h-.973v1.089a3.45 3.45 0 0 0-1.56.622l-.752-.753l-.688.688l.74.74c-.35.449-.586.985-.674 1.572H1.062v.974H2.08c.082.591.316 1.134.664 1.589l-.724.723l.688.687l.729-.729a3.44 3.44 0 0 0 1.583.642v1.048h.973v-1.054a3.448 3.448 0 0 0 1.567-.652l.746.747l.688-.688l-.746-.747c.338-.449.564-.983.645-1.564h1.059v-.974zm-3.391 2.27a1.782 1.782 0 0 1 0-3.565a1.783 1.783 0 1 1 0 3.565m8.535-7.925l1.09-.281l-.252-.979l-1.091.282a2.595 2.595 0 0 0-.461-.672a2.5 2.5 0 0 0-.647-.464l.301-1.079l-.973-.271l-.299 1.072a2.521 2.521 0 0 0-1.566.382l-.76-.776l-.721.707l.756.77a2.526 2.526 0 0 0-.441 1.575l-1.04.268l.252.977l1.038-.268c.117.243.266.475.465.678a2.5 2.5 0 0 0 .686.485l-.289 1.039l.971.271l.293-1.048c.542.033 1.092-.1 1.563-.415l.771.786l.72-.707l-.776-.791a2.537 2.537 0 0 0 .41-1.541m-2.517 1.617a1.49 1.49 0 1 1-.001-2.983a1.49 1.49 0 0 1 .001 2.983"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ghost(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m.008 7.91l.016 8.095L2.021 13l1.978 3.005L6.062 13l1.959 3.005L10.08 13l1.93 3.005L14.049 13l1.956 3.005l-.017-8.095C15.989 3.56 12.363.031 7.986.031C3.609.031.008 3.559.008 7.91m4.508.184a1.563 1.563 0 1 1 0-3.125a1.563 1.563 0 1 1 0 3.125m6.994-.028a1.584 1.584 0 0 1-1.588-1.579c0-.873.711-1.581 1.588-1.581c.878 0 1.588.708 1.588 1.581c0 .872-.71 1.579-1.588 1.579"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gift(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M0 4h16v3H0zm12.66 5.777a.681.681 0 0 1 0 1.029a.87.87 0 0 1-1.14 0L8.485 8.064h-.996l-3.035 2.742a.87.87 0 0 1-1.14 0a.683.683 0 0 1 0-1.029l1.898-1.713L1 8v8h14V8l-4.238.064z"/><path d="M11.928.097c-1.159 0-2.623.838-3.96 2.253C6.634.935 5.17.097 4.014.097c-.592 0-1.098.221-1.46.637c-.599.684-.678 2.021-.293 3.304h1.96s-.59-1.007-.428-1.884c.015-.018.078-.051.22-.051c.573 0 1.905.678 3.057 1.935h1.845c1.157-1.264 2.42-1.935 3.013-1.935c.176 0 .22.049.22.049c.124 1.138-.432 1.886-.432 1.886h2.038c.387-1.283.23-2.62-.365-3.304c-.364-.416-.868-.637-1.46-.637z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GlassWater(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m1.365 12.336l3.295 3.303a1.21 1.21 0 0 0 1.713.015l10.251-7.722a1.213 1.213 0 0 0-.014-1.713L10.765.364A1.209 1.209 0 0 0 9.054.35L1.351 10.622c-.47.47-.462 1.237.014 1.714m5.971-2.564c-1.59.008-3.354-1.247-3.354-1.247l2.677-3.43l2.659-3.408a1.168 1.168 0 0 1 1.631.014l4.393 4.323a1.12 1.12 0 0 1 .012 1.604l-4.565 3.428c-.001.001-1.6-1.293-3.453-1.284"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Global(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.048 0a8 8 0 1 0 .001 16.001A8 8 0 0 0 8.048 0M8 15A7 7 0 1 1 8 1a7 7 0 0 1 0 14"/><path d="M2.959 2.684c-.27 1.344.735 3.399 2.872 3.899c2.136.5 1.218.084.886-.583c-.334-.667-.035-1.167.482-1.333c.518-.167.387.293 1.155-.667c.194-2-1.472-.027-1.638-.86C8.366 1.223 5.95.5 5.249.75c-.701.25-2.02.589-2.29 1.934m3.914 5.01c-.124.055-.702.416-.901.666c-.199.25-.286.778 0 1c.286.223-.016 1.279.755 1.473c.771.194 1.543.307 1.572.917c.027.61-.072 2.027-.443 2.167c.5.277 1.717-1.195 2.145-1.973c.429-.777.572-1.889.543-2.167c-.029-.277.171-.86.257-1.61c.085-.75-.57-.8-.655-1.133c-.087-.334-.203-.701-3.273.66M13.125 4s-2.012.861-.644 2.14c1.366 1.277 2.062 2.49 1.995 2.768c-.062.276-1.38 1.706-.623 1.594c.761-.111 1.322-1.611 1.443-2.501c.119-.89-.924-2.639-.967-2.777c-.036-.141-.722-1.209-1.204-1.224"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Globe(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" transform="translate(3)"><path d="M5.031 13.225v1.434c-2.154.083-3.936.613-3.936 1.26a.05.05 0 0 0 .004.017l8.806.009s.005-.018.005-.025c0-.646-1.818-1.177-3.973-1.26v-1.48"/><ellipse cx="5.012" cy="7.031" rx="5.012" ry="5.031"/><path d="M5.04 13.971a6.882 6.882 0 0 1-4.341-1.549a.499.499 0 0 1-.071-.697a.49.49 0 0 1 .69-.072a5.811 5.811 0 0 0 3.722 1.33c3.263 0 5.916-2.678 5.916-5.967c0-2.918-2.151-5.453-5.005-5.896a5.918 5.918 0 0 0-.911-.07a.493.493 0 0 1-.491-.494c0-.273.22-.495.491-.495c.352 0 .708.027 1.061.081c3.328.518 5.837 3.475 5.837 6.875c0 3.834-3.094 6.954-6.898 6.954"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Glove(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.807 9.708c1.863-1.863 3.68-4.907.307-8.28a4.771 4.771 0 0 0-6.747 0L6.342 2.453v-.719c0-.958-.777-1.735-1.735-1.735C3.648 0 3 .777 3 1.735v4.188l-.847.719l-.613-.613L.313 7.256l7.974 7.974l1.227-1.227l-.549-.549z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GolfBall(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m6.031 12l2 1.969V16h.907v-2.031l2-1.969zm-3-6.547a5.452 5.452 0 1 0 10.906 0a5.452 5.452 0 1 0-10.906 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GolfFlag(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" transform="translate(1)"><ellipse cx="5.433" cy="14.98" rx="5.433" ry=".98"/><path d="M5 0h.956v12.959H5zm2.031.103s1.644-.42 4.45 1.461c2.806 1.882 2.249 3.512 4.103 3.646c.998-.066-.42 1.168-1.854.604c-2.781-1.093-3.6-2.637-6.699-1.21z"/><circle cx="13.945" cy="13.945" r="1.945"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Guitar(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m16.781.959l-.734-.736c-.312-.313-.676-.23-.952.046l-2.289 2.295c-.244.245-.349.529-.128.836L9.324 6.763c-1.125-.788-2.533-.816-3.359.013c-.243.244-.916 1.046-1.014 1.368c-.074.243-.727.647-1.033.643c-.789-.007-1.53.254-2.078.804c-1.247 1.252-1.016 3.519.521 5.062c1.539 1.54 3.798 1.774 5.045.521c.55-.55.811-1.292.803-2.085c-.004-.306.414-.964.664-1.043c.315-.099 1.248-.682 1.486-.919c.818-.821.797-2.212.033-3.334l3.215-3.458c.309.226.623.079.869-.167l2.289-2.294c.276-.278.331-.6.016-.915M7.512 11.133a1.599 1.599 0 0 1-1.605-1.592a1.6 1.6 0 0 1 1.605-1.593c.884 0 1.604.712 1.604 1.593a1.6 1.6 0 0 1-1.604 1.592"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hamburger(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.021 13.061c0 3.262 7.448 2.88 7.981 2.88c.533 0 7.98.531 7.98-2.88zM1 11h15.951v.951H1zm.318-1.009a.32.32 0 0 0 .179-.055L3 8.744l1.836 1.18a.312.312 0 0 0 .352-.002l1.828-1.179L8.801 9.91a.309.309 0 0 0 .349 0l1.833-1.168l1.808 1.168c.107.072.244.07.351-.002l1.867-1.152l1.469 1.182a.313.313 0 0 0 .445-.094a.34.34 0 0 0-.09-.463l-1.645-1.303a.313.313 0 0 0-.354-.001l-1.871 1.155l-1.807-1.168a.314.314 0 0 0-.349.001L8.972 9.232L7.187 8.064a.31.31 0 0 0-.351.002L5.009 9.244L3.173 8.065a.31.31 0 0 0-.352.001l-1.68 1.313a.34.34 0 0 0-.09.462c.062.098.164.15.267.15m15.479-5.033c0-.011.004-.021.002-.03C16.204.662 11.014.066 8.998.066s-7.117.597-7.8 4.862l.001.03zm-2.839-2h1.085V4.02h-1.085zm-11 0h1.073v1.058H2.958zm3-1h1.058V3H5.958zm5 0H12v1.058h-1.042zm-3 1H9.02V4.02H7.958zM8.979.974H10v1.053H8.979zM1 6h15.924v.902H1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hammer(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.216 14.949c.532.533.859.154 1.295-.281c.436-.436.815-.764.284-1.296c0 0-7.639-7.632-9.468-9.455L4.75 5.494zM2.048 7.015l.614-.604s-.271-.743.126-1.099s1.067-.136 1.067-.136L6.01 3.093s-.151-1.083.049-1.283c.2-.2 2.434-1.289 2.651-1.507l-.459-.459S5.123.219 4.784.558c-.199.2-1.689 1.704-2.751 2.766c0 0 .267.759-.083 1.109c-.351.351-1.141.099-1.141.099l-.623.623c-.263.265-.108.637.215.96l.686.686c.325.323.698.477.961.214"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HammerAndWrench(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.217 15.949c.531.533.859.154 1.295-.281c.436-.435.814-.764.283-1.296c0 0-7.402-7.395-9.23-9.218L4.987 6.731zM2.048 8.015l.614-.604s-.271-.743.126-1.099s1.067-.136 1.067-.136L6.01 4.093s-.151-1.083.049-1.283c.2-.2 2.435-1.289 2.651-1.507L8.251.844s-3.128.376-3.467.714c-.199.2-1.688 1.704-2.75 2.766c0 0 .266.759-.084 1.109c-.351.351-1.141.1-1.141.1l-.623.622c-.263.265-.108.637.215.96l.686.686c.325.324.698.477.961.214m9.685-1.5c1.077.511 2.428.354 3.322-.54c.69-.69.964-1.639.817-2.531l-1.521 1.519l-1.294.321l-1.462-1.442l.343-1.337l1.507-1.488c-.893-.146-1.902.065-2.592.756c-.895.895-1.017 2.279-.506 3.357l-.994.993l1.388 1.385zm-4.301 4.604L5.927 9.615l-1.308 1.309a.822.822 0 0 0-.181.266c-.077-.03-.12-.031-.147-.02a2.504 2.504 0 0 0-.791-.139c-1.381 0-2.5 1.105-2.5 2.469c0 1.364 1.119 2.469 2.5 2.469S6 14.864 6 13.5c0-.283-.059-.551-.146-.804c-.005-.024-.006-.046-.02-.081a.853.853 0 0 0 .288-.188zM3.5 14.938c-.812 0-1.469-.643-1.469-1.438c0-.795.656-1.438 1.469-1.438c.813 0 1.469.643 1.469 1.438c0 .795-.657 1.438-1.469 1.438"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hand(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 6.792c-.39-.193-.847-.089-1.023.231c0 0-1.096 2.399-1.734 2.297c-.344-.056-.619-.38-.742-1.005V1.771c0-.426-.439-.771-.98-.771c-.54 0-1.02.346-1.02.771v5.167h-1V.771c0-.426-.439-.771-.98-.771c-.54 0-1.02.346-1.02.771v6.167h-1V2.771c0-.426-.44-.771-.981-.771c-.54 0-1.019.346-1.019.771v8.415c0 2.584 1.729 4.721 5.678 4.721c4.883 0 6.205-8.188 6.205-8.188c.174-.32.003-.736-.384-.927"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandLamp(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m11.951 7.979l1.006.013l-1.039-.975V4.866C11.918 2.554 9.807.06 7.49.06h-.969c-2.316 0-4.438 2.547-4.438 4.858v2.117l-1.13.892l1.09.01l-.006 4.647l-.975 1.395l2.602.655l1.986.438l1.348.862l1.352-.858l2.006-.401l2.624-.613l-1.021-1.615zM6.035 13.98l-3.098-.578V7.917l3.098.372zm5.002-.49l-3.069.466V8.367l3.069-.357zM8.049 4.828c.135-.191.215-.418.215-.664c0-.67-.582-1.213-1.297-1.213s-1.295.543-1.295 1.213c0 .247.08.476.215.668L2.943 6.644V5.021C2.943 3.189 4.59.996 6.534.996H7.5c1.943 0 3.525 2.193 3.525 4.025v1.62z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandSwitch(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.484 3.062c-.82 0-1.484.637-1.484 1.422c0 .062.011.121.02.181L9.317 6.853c-.739-1.108-1.866-1.848-3.348-1.848l-.016.001l-.007-2.604c0-.715-.775-1.294-1.525-1.316c-.751.022-1.353.602-1.353 1.316v13.13c0 .715.602 1.294 1.353 1.316c.75-.022 1.525-.602 1.525-1.316v-2.604l.022.002c2.3 0 3.993-1.775 3.993-3.963c0-.309-.04-.607-.11-.895l3.955-2.332c.204.102.432.164.678.164c.82 0 1.484-.637 1.484-1.422c0-.785-.663-1.42-1.484-1.42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Handcuff(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M10.556 6.076c-.366-.167-1.679.441-1.524.829c.162.467.309 1.501.064 2.274c-.094.299-.139.543-.193.607c-.075.146.771 1.2 1.412 1.005c.221-.336.271-.61.42-1.081c.392-1.242.292-3.361-.179-3.634M6.852.066c-.41-.084-1.254-.104-1.484.14c-.023.025-.05.049-.06.079l-1.444 2.49a5.273 5.273 0 0 0-1.439.685A5.33 5.33 0 0 0 .284 6.301a5.399 5.399 0 0 0 .083 3.508c.227.617.564 1.188.997 1.684a5.275 5.275 0 0 0 2.364 1.568a5.21 5.21 0 0 0 2.97.074c-.635-.447-.979-.916-1.434-1.639a3.413 3.413 0 0 1-.991-.159a3.464 3.464 0 0 1-1.113-.613a3.545 3.545 0 0 1-1.067-1.453a3.575 3.575 0 0 1-.105-2.419A3.535 3.535 0 0 1 3.5 4.921a3.473 3.473 0 0 1 1.709-.534h.051c.26-.008.359.009.551.071c.104-.155.465-.667.91-.864l.498-2.487c.135-.3.415-.492.549-.653L6.981.098C6.944.086 6.898.076 6.852.066"/><path d="m13.524 4.241l-.895-2.475c-.002-.38-1.062-.684-1.454-.682l-1.564.005c-.357.001-1.375.263-1.543.6c-.015.03-.033.06-.033.092l-.623 2.48c-.455.3-.852.671-1.195 1.087A5.177 5.177 0 0 0 5.011 8.68a5.199 5.199 0 0 0 1.181 3.275c.416.512.928.944 1.506 1.277c.82.475 1.779.751 2.807.748c3.021-.011 5.465-2.398 5.455-5.336c-.009-1.835-.974-3.453-2.436-4.403m-2.555-2.257h1.062V3h-1.062zm-2 0h1.062V3H8.969zm1.453 10.266c-2.097 0-3.797-1.686-3.797-3.766s1.7-3.766 3.797-3.766c2.096 0 3.797 1.686 3.797 3.766s-1.701 3.766-3.797 3.766"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hanger(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.542 11.375S9.841 6.953 9.579 6.846V5.96c.827-.233 1.38-1.005 1.38-1.915c0-1.098-.885-1.99-1.971-1.99a1.982 1.982 0 0 0-1.969 1.99c0 .293.237.529.531.529c.293 0 .426-.342.426-.635c0-.512.545-.981 1.044-.981c.5 0 1.021.575 1.021 1.087c0 .511-.554.928-1.054.928c-.294 0-.473.237-.473.529v1.322c-.198.073-6.995 4.551-6.995 4.551c-.503.513-1.188 2.582 1.206 2.582h12.581c2.194 0 1.717-2.094 1.236-2.582m-.985 1.646H2.456c-.671 0-.652-.756-.347-1.086l6.349-4.016c.343-.182.774-.206 1.113-.004l6.37 4.021c.307.33.231 1.085-.384 1.085"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hardware(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.007 4.014a2.977 2.977 0 0 0 0 5.954a2.977 2.977 0 0 0 0-5.954m0 5.021a2.046 2.046 0 0 1 0-4.093a2.046 2.046 0 0 1 0 4.093M6 14h3.965v2H6zm0-3h3.964v1.959H6zM2.473 4.53l2.1-3.066l1.211.83l-2.1 3.065zm-2.55-1.626l2.1-3.066L3.23.665l-2.1 3.066zm11.673-1.733h3.734v3.457h-3.734z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hat(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M12.708 8.211c-.43.917-2.16 2.028-4.708 2.028c-2.532 0-4.333-1.097-4.712-2.013C1.375 8.312 0 9.456 0 10.136c0 1.128 3.581 2.728 8 2.728c4.418 0 8-1.6 8-2.728c0-.689-1.375-1.567-3.292-1.925"/><path d="M10.077 3.197c-.516 0-1.582.354-2.065.354c-.483 0-1.549-.354-2.064-.354c-1.114 0-1.81 1.372-1.907 1.897l-.02 2.427c.562.506 1.611 1.142 2.761 1.238a15.341 15.341 0 0 0 2.52 0c1.15-.097 2.074-.654 2.668-1.221l-.005-2.413c-.086-.546-.784-1.928-1.888-1.928"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HatChef(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4 15h10v1H4zm9.449-11.688A4.243 4.243 0 0 0 9.322.086a4.244 4.244 0 0 0-4.166 3.39a3.28 3.28 0 0 0-.833-.119a3.31 3.31 0 0 0-3.312 3.312c0 1.829 1.187 3.312 3.017 3.312c.019 0 .036-.005.056-.005v3.981h9.875v-3.384c1.883-.14 3.011-1.709 3.011-3.627a3.642 3.642 0 0 0-3.521-3.634M8.9 1.795c-.842.544-1.383 1.401-1.383 2.368H6.216c0-1.398 1.133-2.571 2.664-2.9c.277-.058.452.252.02.532"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Head(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.001 5.378S.812 1.177 5.145.305c4.333-.872 6.963 1.506 7.35 3.171c.387 1.665.575 2.518.387 3.171c-.232.317-.155.555.309 1.348c.464.793 1.393 1.665.077 1.823c-.851.159-.309.713-.309.713s.232.476-.464.555c0 0 .464.872.077.872s-.464.159-.542.476c-.077.317.077 2.061-2.631 1.03c-.619-.317-.892.977-1.238 1.744c-.208.461-.431 1.09-1.47.634c-.927-.407-2.995-1.378-3.771-2.705c-.424-.727.693-1.778.151-3.205c-.542-1.427-.712-.987-1.254-1.779c-.541-.793-.738-1.824-.816-2.775"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeadSet(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m15.665 6.733l-.01.002V5.632c0-3.066-2.601-5.561-5.798-5.561H8.195c-3.197 0-5.797 2.494-5.797 5.561V6.74c-.025 0-.05-.008-.076-.008c-.707 0-1.282.59-1.282 1.317v2.687c0 .728.575 1.314 1.282 1.314c.062 0 .125-.011.185-.019c.579 1.731 2.681 3.066 5.289 3.314c.278.367.398.631 1.249.631c1.092 0 1.913-.43 1.913-.961c0-.53-.821-.961-1.913-.961c-.844 0-.961.26-1.244.622c-2.355-.235-4.236-1.41-4.686-2.914c.447 0 .816-.608.816-1.027v-.122h.006V5.574c0-2.547 1.518-3.618 4.258-3.618h1.662C12.59 1.956 14 2.992 14 5.585v5.04h.014v.11c0 .728.752 1.221 1.466 1.221s1.478-.493 1.478-1.221V8.048a1.306 1.306 0 0 0-1.293-1.315"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heart(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.958 1.03a4.015 4.015 0 0 0-3.911 3.148A4.054 4.054 0 0 0 5.102 1.03c-2.237 0-4.05 1.824-4.05 4.072c0 6.496 8.005 9.838 8.005 9.838s7.912-3.258 7.912-9.879c0-2.228-1.795-4.031-4.011-4.031"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartDelete(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="m11.19 14.2l1.702-1.7l-1.703-1.699l.631-.632l1.701 1.701l1.7-1.7l.631.631l-1.7 1.699l1.7 1.7l-.631.631l-1.7-1.7l-1.699 1.7z"/><path d="m9.633 10.802l2.188-2.189l1.701 1.702l1.536-1.537c.551-1.094.911-2.326.911-3.716a4.019 4.019 0 0 0-4.011-4.031a4.015 4.015 0 0 0-3.911 3.148a4.054 4.054 0 0 0-3.945-3.148c-2.237 0-4.05 1.824-4.05 4.072c0 6.496 8.005 9.838 8.005 9.838s.729-.304 1.74-.902l1.539-1.538z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartPlus(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M11 12h4.958v.918H11z"/><path d="M13 10h.918v4.957H13z"/><path d="M11.917 8.958h3.055c.605-1.135.997-2.431.997-3.896a4.019 4.019 0 0 0-4.011-4.031a4.015 4.015 0 0 0-3.911 3.148a4.054 4.054 0 0 0-3.945-3.148c-2.237 0-4.05 1.824-4.05 4.072c0 6.496 8.005 9.838 8.005 9.838s.785-.324 1.86-.974v-3.05h2z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartRemove(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M9.917 10.917h3.696c1.318-1.519 2.355-3.464 2.355-5.855a4.019 4.019 0 0 0-4.011-4.031a4.015 4.015 0 0 0-3.911 3.148a4.054 4.054 0 0 0-3.945-3.148c-2.237 0-4.05 1.824-4.05 4.072c0 6.496 8.005 9.838 8.005 9.838s.785-.324 1.86-.974z"/><path d="M11 12h4.958v.918H11z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartSignal(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="m9.953 8.107l-1.304 2.385a.33.33 0 0 1-.331.17a.329.329 0 0 1-.287-.23l-.975-3.316l-1.213 3.946a.332.332 0 0 1-.317.242h-.009a.333.333 0 0 1-.319-.229L3.44 5.447L2.303 7.933c-.058.106-.126.219-.251.219H.492c1.686 4.589 7.555 6.897 7.555 6.897s5.845-2.295 7.511-6.942z"/><path d="m1.801 6.945l1.425-3.108a.328.328 0 0 1 .328-.181a.336.336 0 0 1 .29.238l1.647 5.511l1.221-4.16a.333.333 0 0 1 .319-.253h.004c.15 0 .283.1.323.244l1.089 3.868l1.007-1.928c.058-.11.173-.231.296-.231h6.006A7.584 7.584 0 0 0 16 5.044a4.03 4.03 0 0 0-7.963-.882A4.072 4.072 0 0 0 0 5.086c0 .672.09 1.256.244 1.859z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Helicopter(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M3.276 9.667c-1.255 0-2.275-1.038-2.275-2.312c0-1.275 1.021-2.312 2.275-2.312c.444 0 .874.13 1.244.376l-.364.545a1.58 1.58 0 0 0-.88-.266c-.893 0-1.619.743-1.619 1.657c0 .913.727 1.657 1.619 1.657zm12.978-5.615s-4.723.162-4.723.362v.132h-.043v-.132c0-.2-4.721-.362-4.721-.362c-.39 0-.705.162-.705.362v.205c0 .199.315.361.705.361c0 0 3.104-.107 4.277-.253v2.189h.922v-2.19c1.169.147 4.287.254 4.287.254c.39 0 .705-.162.705-.361v-.205c.001-.2-.314-.362-.704-.362m-1.476 10.896H8.863c-1.211 0-2.842-.904-2.842-1.62c0-.173.152-.312.34-.312c.188 0 .34.14.34.312c0 .309 1.314.788 2.162.788h5.915c.847 0 1.48-.475 1.48-.782c0-.173.153-.312.339-.312c.188 0 .341.14.341.312c0 .715-.948 1.614-2.16 1.614"/><path d="M12.59 6.125c-1.254 0-2.892.388-4.178.878H2.955a.626.626 0 0 0-.622.63l.698.228c0 .347.278.629.622.629l3.448.866c1.006 1.51 1.717 3.562 5.488 3.562c2.408 0 4.359-1.487 4.359-3.323c0-1.836-1.95-3.47-4.358-3.47m.559 3.928c-.681 0-1.229-.719-1.229-1.605c0-.885-.156-1.604.522-1.604c.683 0 2.726.719 2.726 1.604c0 1.724-1.338 1.605-2.019 1.605"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HelicopterPad(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M8 .055c-4.418 0-8 3.566-8 7.968c0 4.4 3.582 7.968 8 7.968s8-3.567 8-7.968C16 3.621 12.418.055 8 .055m.004 15.057c-3.934 0-7.121-3.181-7.121-7.105C.883 4.083 4.071.902 8.004.902c3.933 0 7.121 3.181 7.121 7.105c0 3.924-3.187 7.105-7.121 7.105"/><path d="M8.018 2.08c-3.264 0-5.91 2.654-5.91 5.927c0 3.273 2.646 5.927 5.91 5.927c3.264 0 5.911-2.654 5.911-5.927c0-3.273-2.648-5.927-5.911-5.927m2.059 8.039h-1.14V9.062H7.062v1.057H5.944V5.961h1.118v1.914h1.875V5.961h1.14z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HelmWheel(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.922 6.992c-.237 0-.621.076-.994.199a5.968 5.968 0 0 0-1.193-2.853c.354-.177.68-.396.848-.563c.412-.412.434-1.059.051-1.441c-.385-.385-1.031-.362-1.442.049c-.169.17-.39.499-.566.854a5.926 5.926 0 0 0-2.858-1.179c.126-.376.203-.763.203-1.002c0-.582-.441-1.055-.984-1.055c-.543 0-.984.473-.984 1.055c0 .24.078.63.204 1.007a5.97 5.97 0 0 0-2.854 1.184c-.177-.354-.396-.683-.566-.852c-.411-.412-1.057-.434-1.441-.05c-.384.384-.361 1.03.05 1.441c.169.169.499.39.854.566a5.97 5.97 0 0 0-1.185 2.856c-.373-.125-.76-.201-.997-.201c-.582 0-1.055.441-1.055.984c0 .543.473.984 1.055.984c.237 0 .622-.076.995-.2a5.946 5.946 0 0 0 1.186 2.856c-.351.176-.676.394-.842.561c-.412.412-.434 1.058-.051 1.441c.385.385 1.031.362 1.442-.049c.167-.167.383-.488.56-.837a5.949 5.949 0 0 0 2.859 1.179c-.123.37-.197.751-.197.986c0 .582.441 1.055.984 1.055c.543 0 .984-.473.984-1.055c0-.235-.074-.615-.196-.985a5.956 5.956 0 0 0 2.856-1.19c.176.348.391.67.557.836c.412.412 1.059.434 1.443.051c.384-.385.361-1.031-.051-1.442c-.166-.166-.488-.382-.836-.558a5.788 5.788 0 0 0 1.181-2.859c.37.121.749.195.983.195c.582 0 1.055-.441 1.055-.984c0-.543-.476-.984-1.058-.984m-3.229-2.113a4.822 4.822 0 0 1 1.12 2.709l-3.534.321a1.205 1.205 0 0 0-.293-.704zM8.839 7.058a1.325 1.325 0 0 0-.706-.3l.28-3.582a4.826 4.826 0 0 1 2.714 1.136zM7.627 3.174l.3 3.574c-.282.011-.54.107-.749.264L4.877 4.299a4.809 4.809 0 0 1 2.75-1.125M4.305 4.87l2.707 2.288c-.19.201-.311.461-.332.748l-3.505-.299a4.818 4.818 0 0 1 1.13-2.737m2.38 3.262c.032.282.157.536.354.729l-2.712 2.282a4.798 4.798 0 0 1-1.149-2.728zm.52.867c.205.146.455.233.728.243l-.3 3.574a4.826 4.826 0 0 1-2.734-1.105zm1.215 3.813l-.288-3.58c.262-.03.499-.132.689-.287l2.319 2.72a4.8 4.8 0 0 1-2.72 1.147m3.284-1.714L8.969 8.806a1.22 1.22 0 0 0 .307-.688l3.539.27a4.828 4.828 0 0 1-1.111 2.71"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Helmet(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.706 8.697h-.124c-.344-3.754-3.457-6.693-7.253-6.693c-4.026 0-7.288 3.305-7.288 7.383c0 2.579 1.192 4.589 5.218 4.589c3.146 0 1.98-2.803 5.315-4.026h4.132c.702 0 1.272-.271 1.272-.601v-.051c.001-.331-.57-.601-1.272-.601m-9.644 2.345H4.937V9.958h1.125z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hightheel(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.396 8.909s-.521-.239-.876 1.241c-.271 1.421-.538 1.734-.538 1.734H2.36s.007-.652.007-1.141s-.283-2.062-.987-3.219C.678 6.367.879 4.881 2.764 3.026c.074-.072.442-.211.902.231c.598.573 3.321 3.391 5.629 4.371c2.168.921 3.627-.121 3.627-.121s2.825.631 3.881 2.413c1.055 1.784-2.296 1.908-2.296 1.908s-4.169.24-5.175-.605c-.151-.192-4.123-2.473-4.936-2.314"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func History(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.479 1.046c3.859 0 6.819 3.192 7.166 6.985h1.324l-1.892 1.952l-2.065-1.952h1.338c-.33-3.229-2.746-5.958-5.936-5.958c-3.412 0-6.189 2.888-6.189 6.437c0 3.549 2.777 6.438 6.189 6.438c1.695 0 3.231-.713 4.35-1.865l.822.826a7.369 7.369 0 0 1-5.107 2.065c-4.092 0-7.419-3.349-7.419-7.464s3.327-7.464 7.419-7.464"/><path d="M8.058 4.03L8 8.953L12 9V8l-3.032.062V4.031z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hockey(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.938 15.969H2.032a1 1 0 0 1-1-1c0-.553.447-.916 1-.916h3.807c.062-.113.143-.305.178-.388L11.022.818a1.002 1.002 0 0 1 1.317-.517c.506.222.737.811.517 1.317l-5 12.75c-.251.592-.675 1.601-1.918 1.601M10 15h2.938v1H10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HomePage(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m7.65 4.158l-4.485 4.94a.648.648 0 0 0 .002.849l2.868.005v5.024c0 .553.439 1 .982 1h1.965a.99.99 0 0 0 .982-1v-5.02l2.811.005a.65.65 0 0 0-.004-.848L8.414 4.159a.506.506 0 0 0-.764-.001m6.268-2.324c0 .552-.437 1-.973 1H3.056c-.537 0-.973-.448-.973-1V1c0-.552.436-1 .973-1h9.889c.536 0 .973.448.973 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Horse(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.572 12.36c-1.889-2.274-.475-5.31-2.205-8.448a4.144 4.144 0 0 0-1.564-1.63a.975.975 0 0 0 .033-.03c.916-.928 1.412-1.903 1.106-2.178c-.305-.277-1.296.252-2.212 1.18c-.176.178-.33.356-.472.533c-.976-.081-2.003.242-2.895 1.095c0 0-3.909 3.697-4.949 4.697c-1.044.998.367 1.534.367 1.534s4.003-.967 4.833-.645c1.445.558-2.029 2.992-1.918 6.148c.066 2.044 5.214 1.205 8.562.225c1.825-.533 1.67-2.054 1.314-2.48z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HorseShoe(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.502 1.042c-4.545 0-7.424 4.236-7.424 9.462c0 2.405 1.523 5.554 2.245 6.265c.723.711 2.324-.822 2.05-1.323a8.363 8.363 0 0 1-1.014-3.969c0-3.647 1.605-6.604 4.143-6.604c2.535 0 4.142 2.957 4.142 6.604c0 1.486-.41 2.855-1.011 3.957c-.272.507 1.469 2.019 2.045 1.335c.577-.684 2.247-3.859 2.247-6.265c0-5.226-2.88-9.462-7.423-9.462"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hospital(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16 1H3v3.084H2v9.874c0 .547.447.989 1 .989h5V12h3.042v3l4.875-.053c.552 0 1.083-.442 1.083-.989V4h-1zM7.031 12.97H4v-1h3.031zm0-3H4v-1h3.031zm0-3H4v-1h3.031zM6.959 4h-3V3h3zM11 5.007h-1v.977H9v-.977H8V3.98h1v-.997h1v.997h1zm.953-2.028h3v1.047h-3zM15 12.97h-3v-1h3zm0-3h-3v-1h3zm0-3h-3v-1h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func House(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="m16.794 8.77l-3.81-3.906V2.017l-.968.022v1.728L9.502 1.245a.711.711 0 0 0-1.003 0L1.206 8.771a.713.713 0 0 0 0 1.002a.71.71 0 0 0 1.003-.001L9 2.982l6.793 6.79a.704.704 0 0 0 1.001.001a.715.715 0 0 0 0-1.003"/><path d="M4.043 9.532v5.69c0 .394.218.786.567.786h2.277l.064-3.993h4.074l-.002 3.993h2.303c.349 0 .632-.391.632-.786V9.531L9 4.625z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IceCream(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M9.513 16.906L12 8.042H7z"/><path d="M11.775 2.732C11.459 1.725 10.617 1 9.617 1c-1.09 0-1.996.856-2.233 2.002c-.021 0-.041-.008-.063-.008c-1.27 0-2.298 1.159-2.298 2.589c0 1.261.694 2.308 1.754 2.538l5.529.124C13.248 7.87 14 6.63 14 5.44c0-1.432-.98-2.592-2.225-2.708"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IdBadge(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.969 2.021H2.021C.911 2.021.01 2.916.01 4.016V11c0 1.101.9 1.995 2.011 1.995h.95v-1.042h2.027v1.042h5.961v-1.042h2.051v1.042h.959c1.11 0 2.01-.895 2.01-1.995V4.016c0-1.101-.9-1.995-2.01-1.995M5.457 3.969c.827 0 1.5.688 1.5 1.534c0 .847-.674 2.466-1.5 2.466c-.827 0-1.5-1.62-1.5-2.466c0-.846.673-1.534 1.5-1.534m3.537 7.05H1.987s-.12-3.028 1.679-3.028c.374.393.892.77 1.842.77c.951 0 1.404-.381 1.775-.778c2.025-.002 1.711 3.036 1.711 3.036m5.035-1.066H9.984v-1h4.045zm-2.06-1.984h-2v-1h2zm2.047-2.015H9.969v-1h4.047z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Image(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" transform="translate(0 2)"><path d="M13.438 11.944H2.557c-1.394 0-2.528-1.163-2.528-2.591v-6.75c0-1.43 1.135-2.591 2.528-2.591h10.881c1.393 0 2.527 1.161 2.527 2.591v6.75c0 1.428-1.135 2.591-2.527 2.591M2.237.979c-.7 0-1.272.614-1.272 1.371v7.318c0 .757.572 1.371 1.272 1.371h11.517c.702 0 1.273-.614 1.273-1.371V2.35c0-.757-.571-1.371-1.273-1.371z"/><ellipse cx="5.471" cy="3.461" rx="1.471" ry="1.461"/><path d="m11.234 3.037l2.76 6.951H2.021L5.497 5.98l3.117.944z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InColumns(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 1h4v12.998H1zm5 0h4v12.973H6zm5 0h4v13.019h-4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxDownload(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" fill-rule="evenodd"><path fill="currentColor" d="m6.198 4.256l2.19 2.513a.811.811 0 0 0 1.143 0l2.189-2.513c.314-.315.363-.876.049-1.19H9.965V1.144c0-.559-.433-1.01-.968-1.01c-.535 0-.969.451-.969 1.01v1.922H6.247c-.315.315-.364.874-.049 1.19"/><path d="M7 10h4.031v1.031H7z"/><path fill="currentColor" d="M14.993 1.006h-3.962v.975h3.146L15.54 9.01H2.505l1.407-7.029h3.057v-.975H3.073l-2.06 8.755v4.17c0 1.334.472 2.028 1.804 2.028h12.28c1.246 0 1.885-.527 1.885-2.111V9.761zM11.016 11H6.985V9.969h4.031z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InboxUpload(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M11.768 2.757L9.579.244a.811.811 0 0 0-1.143 0L6.247 2.757c-.315.315-.363.876-.049 1.19H8v1.922c0 .559.434 1.01.969 1.01c.535 0 .969-.451.969-1.01V3.947h1.781c.315-.315.364-.875.049-1.19"/><path d="M14.993 1.006h-3.29l.891.975h1.583L15.54 9.01H2.505l1.407-7.029h1.682l.766-.975H3.073l-2.06 8.755v4.17c0 1.334.472 2.028 1.804 2.028h12.28c1.246 0 1.885-.527 1.885-2.111V9.761zM11.016 11H6.985V9.969h4.031z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Infinity(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.273 4.042c-1.331 0-2.836.703-3.769 1.772c-.18.206-.775.189-.964-.023c-.936-1.056-2.436-1.749-3.763-1.749C1.705 4.042.02 5.822.02 8.01c0 2.189 1.685 3.97 3.758 3.97c1.345 0 2.865-.711 3.8-1.792c.173-.201.735-.205.908-.004c.931 1.083 2.448 1.796 3.788 1.796c2.067 0 3.746-1.78 3.746-3.97c0-2.188-1.68-3.968-3.746-3.968zm-8.78 7.045C2.05 11.087.876 9.707.876 8.01c0-1.695 1.175-3.076 2.619-3.076c1.534 0 3.476 1.448 3.476 3.076c0 1.629-1.942 3.077-3.476 3.077zm9-.032c-1.53 0-3.474-1.438-3.474-3.056c0-1.616 1.944-3.054 3.474-3.054c1.446 0 2.621 1.371 2.621 3.054c0 1.685-1.175 3.056-2.62 3.056z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InfinityTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.352 12.941c-2.383 0-4.32-1.992-4.32-4.441s1.938-4.441 4.32-4.441a4.29 4.29 0 0 1 3.424 1.73l-.972.951a2.968 2.968 0 0 0-2.452-1.315c-1.648 0-2.988 1.38-2.988 3.075c0 1.695 1.341 3.075 2.988 3.075c.868 0 2.119-1.929 3.126-3.478c1.349-2.078 2.622-4.039 4.173-4.039c2.377 0 4.31 1.986 4.31 4.428s-1.933 4.428-4.31 4.428c-1.339 0-2.6-.641-3.414-1.727l.963-.962a2.964 2.964 0 0 0 2.451 1.323c1.642 0 2.978-1.374 2.978-3.062c0-1.688-1.336-3.062-2.978-3.062c-.836 0-2.073 1.902-3.065 3.431c-1.365 2.102-2.654 4.086-4.234 4.086"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Info(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" transform="translate(4)"><ellipse cx="4.486" cy="1.525" rx="1.486" ry="1.525"/><path d="m4.479 13.325l1.373-8.622s.029-.699-.636-.699c-1.501 0-5.29.004-5.29.004s3.979.713 3.979 2.059c0 0-1.308 5.76-1.308 7.29c0 1.531.836 2.644 2.337 2.644c1.225 0 2.338-1.336 2.254-3.283c-1.197 1.836-2.709 2.102-2.709.607"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ipod(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" transform="translate(4)"><ellipse cx="4.469" cy="11.488" rx="1.469" ry="1.488"/><path d="M8.301.021H.779C.371.021.04.326.04.7v14.595c0 .375.33.679.739.679h7.522c.41 0 .739-.304.739-.679V.7c0-.374-.33-.679-.739-.679M4.5 14a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5M8 6H1V1h7z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Iron(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.317 8.083c.303-.62.636-1.136.636-1.946c0-.787-.226-1.354-.671-1.685c-.541-.4-1.809-.306-2.41-.231c-.086-.042-.2-.199-.302-.199l-6.558.607c-.73.151-1.455.883-1.455.883l-4.518 4.753c0 .374.056.692.669.692L13 10.943S12.491 8 14 8V4.867c.537-.043 1.613-.069 1.896.143c.261.194.393.572.393 1.127c0 .648-.254 1.168-.547 1.768c-.641 1.311-1.367 2.548.273 5.788c.06.116.175.184.295.184a.318.318 0 0 0 .153-.04a.344.344 0 0 0 .14-.461c-1.482-2.929-.865-3.943-.267-5.163zm-3.24-.944c-.316 0-2.145.149-2.145.14v-.554a.635.635 0 0 0-.627-.645h-.597a.635.635 0 0 0-.628.645v.554c0 .075-4.235.681-4.235.681l1.436-1.815s.446-.611.971-.611l5.34-.616c.269 0 .485.272.485.611z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func JumpBackward(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.994 1c0-.553-.442-1-.989-1h-1.979a.994.994 0 0 0-.99 1v14c0 .553.443 1 .99 1h1.979a.994.994 0 0 0 .989-1zM.438 9.052a1.49 1.49 0 0 1 0-2.104L6.882.506c.581-.582 2.103-.839 2.103 1v12.988c0 1.901-1.521 1.582-2.103 1.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func JumpDoublePageLeftRight(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16 8.062L14.172 5.25a.827.827 0 0 0-1.17 0v1.776h-2.058v-5.5s-.431-.5-.958-.5c-.529 0-.958.5-.958.5V14.5s.429.5.958.5c.527 0 .958-.5.958-.5V9h2.058v1.758a.83.83 0 0 0 1.17 0zM5.986 1.026c-.529 0-.958.5-.958.5v5.557L2.97 7.074V5.302a.83.83 0 0 0-1.17 0L0 8l1.801 2.75a.83.83 0 0 0 1.17 0V8.966l2.058-.008V14.5s.429.5.958.5c.527 0 .958-.5.958-.5V1.526c-.001 0-.431-.5-.959-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func JumpForward(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.002 1c0-.553.442-1 .989-1H4.97c.547 0 .989.447.989 1v14c0 .553-.442 1-.989 1H2.991a.994.994 0 0 1-.989-1zm8.111 14.495c-.582.581-2.103.9-2.103-1.001V1.506c0-1.839 1.521-1.582 2.103-1l6.444 6.442a1.49 1.49 0 0 1 0 2.104z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func JumpPageLeftRight(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.958 8.951V7.007H4.979v-1.94c0-.49-3.714 2.265-3.714 2.265a.946.946 0 0 0-.001 1.323s3.715 2.818 3.715 2.293V8.95h1.979zM11.002 7v1.973h2.046V11c0 .455 3.647-2.316 3.647-2.316a.871.871 0 0 0 0-1.261s-3.714-2.78-3.714-2.388L13.048 7zM8 0h2v16H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func JumpPageUpDown(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.979 12.021v-1.979H8.02v1.979H6.291a.851.851 0 0 0 0 1.255l2.056 2.459a.96.96 0 0 0 1.334.001l2.053-2.46a.85.85 0 0 0 0-1.255zM8 3.952v2.046h1.979V3.952h1.73c.365-.36.365.055 0-.307L9.675.305a.949.949 0 0 0-1.326 0l-2.098 3.34c-.365.361-.365-.054 0 .307zM1 7h16v2H1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Kette(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" fill-rule="evenodd"><path d="m2.4 4.945l.683 7.262l1.044-.001c.173-3.921.648-5.88 1.33-7.604l-2.975.337a.555.555 0 0 1-.082.006"/><path fill="currentColor" d="M4.051 16.081c0 .416.009.476.022.878h9.857c.013-.402.022-.462.022-.878v-.019H4.05zm.076-3.875l-1.044.001L2.4 4.945a.566.566 0 0 0 .082-.006l2.975-.337c-.682 1.724-1.157 3.683-1.33 7.604m8.947-7.878l-.506-.319c-.16-.417-.338-.597-.521-.915H6.379c-.056.077-4.027.509-4.027.509c-.715.025-1.289.627-1.286 1.414l.669 7.171c0 .749.846 1.357 1.583 1.357h.763c-.012.438-.001.95-.004 1.417h9.892c-.03-3.34-.173-5.917-.523-7.898c.005-.01 2.462-2.47 2.462-2.47c.152-.319-.419-.327-.908-.563c0 0-1.773-.02-1.926.297M7.031 1.997h4.914c-.735-.845-1.446-.955-2.361-.955c-.71 0-1.945.29-2.553.955"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Key(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.885.764C8.678 1.971 8.862 3.473 9.609 4.22l1.121 1.121l-9.496 9.497c-.246.246-.229.667.041.937s.689.288.937.04l1.387-1.386l1.556 1.555l.534-.532l.444-.446l.533-.532l-1.556-1.557l1.355-1.355l1.066 1.066l.979-.979l-1.065-1.065l4.263-4.265l1.121 1.121c.783.783 2.434.748 3.457-.275c1.023-1.023.82-2.748.111-3.458L13.176.488c-.592-.593-2.248-.767-3.291.276m5.42 3.969c.297.297-.162 1.12-.27 1.228c-.104.104-.926.562-1.224.265l-2.946-2.945c-.297-.297.162-1.118.268-1.224c.107-.107.929-.565 1.226-.27z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func KeyTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.797 1.272c-1.549-1.55-4.148-1.65-5.69-.107c-1.119 1.117-1.317 2.834-.812 4.233l-8.223 8.223s-.534 2.827 1.85 2.215c.058-.076 1.047-.586 1.047-.586l.023-2.289L6.215 13l.74-.875l.014-2.172h1.977l-.002-1.975l2.168-.008l.379-.258c1.439.602 3.197.162 4.321-.843c1.626-1.453 1.535-4.048-.015-5.597m-.781 1.759h-1.047V1.984h1.047z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Keyboard(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 4.031V11h14.997V4.031zM11.969 5h1.047v1.023h-1.047zm1.047 1.969v1.047h-1.031V6.969zm-3.049-1.99h1.05v1.06h-1.05zm1.055 1.996v1.062H9.979V6.975zm-3.04-1.996h1.033v1.036H7.982zm-.017 1.996h1.078v1.084H7.965zm-2.993-1.98h1.054v1.028H4.972zM7 6.988v1.049H5.985V6.988zm-1.969-.013v1.062H3.97V6.975zM2.969 4.984h1.062v1.047H2.969zm-1 1.995H3v1.037H1.969zM4 10H2V9h2zm8.021.021H4.969V8.968h7.052zM15 10h-2V9h2zm.016-2h-1.031V6.969h1.031zm0-1.977h-1.047V5h1.047z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Knife(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.344 5.262L2.746 2.664c-.454-.453-.943.032-1.525-.551C.639 1.531 1.294.745 1.7.34c.406-.405 1.025-.444 1.383-.087l3.275 3.276c.357.358-.657 2.092-1.014 1.733m11.594 10.363c-1.212 1.212-6.35-3.236-11.623-8.51l2.318-2.311c0 .001 9.885 10.24 9.305 10.821"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LadderPool(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M3.168 9.982c-.89 0-1.607-.338-1.968-.624a.437.437 0 0 1-.059-.634a.488.488 0 0 1 .66-.059c.053.041 1.26.946 2.87-.032c2.147-1.308 4.494-.063 4.593-.011c.018.009 1.94 1.013 3.676.007c2.269-1.32 3.867-.011 3.933.047a.436.436 0 0 1 .041.634a.486.486 0 0 1-.665.039c-.041-.035-1.154-.919-2.82.048c-2.213 1.284-4.53.06-4.629.008c-.02-.012-1.957-1.031-3.625-.013c-.711.434-1.4.59-2.007.59m.007 3.012c-1.148 0-1.91-.566-1.957-.604a.472.472 0 0 1 .566-.756c.057.042 1.279.919 2.903-.031c2.132-1.246 4.462-.062 4.56-.011c.018.009 1.958.984 3.709.007c2.254-1.258 3.833-.009 3.9.045a.476.476 0 0 1 .071.666a.47.47 0 0 1-.662.072c-.043-.034-1.17-.892-2.853.045c-2.195 1.225-4.498.058-4.597.007c-.018-.009-1.973-.998-3.655-.011c-.731.427-1.407.571-1.985.571m8.042 2.992a5.43 5.43 0 0 1-2.417-.584c-.02-.01-1.957-1.037-3.623-.011c-1.763 1.082-3.373.452-3.977-.034a.442.442 0 0 1-.059-.639a.482.482 0 0 1 .66-.059c.053.041 1.26.954 2.87-.033c2.147-1.32 4.494-.065 4.593-.011c.02.009 1.941 1.022 3.676.007c2.269-1.331 3.867-.009 3.933.046a.444.444 0 0 1 .041.641a.486.486 0 0 1-.665.04c-.041-.037-1.154-.926-2.82.047a4.312 4.312 0 0 1-2.212.59M10.31 4.972H4.665c-.276 0-.5-.21-.5-.469s.224-.469.5-.469h5.645c.276 0 .5.21.5.469s-.224.469-.5.469m0 2.028H4.665a.5.5 0 0 1 0-1h5.645a.5.5 0 0 1 0 1"/><path d="M4.485 9.708c-.276 0-.454-.224-.454-.5V2.281c-.033-.048-.379-.344-.888-.344h-.251c-.51 0-.81.296-.849.373a.485.485 0 0 1-.497.485a.51.51 0 0 1-.498-.515c0-.704.793-1.234 1.844-1.234h.251c1.05 0 1.842.53 1.842 1.234v6.927a.5.5 0 0 1-.5.501m6.015.021a.5.5 0 0 1-.5-.5V2.248c-.033-.048-.256-.311-.766-.311h-.252c-.509 0-.886.263-.925.34a.485.485 0 0 1-.497.485a.51.51 0 0 1-.497-.515c0-.704.792-1.234 1.843-1.234h.251c1.051 0 1.843.53 1.843 1.234v6.981a.5.5 0 0 1-.5.501"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lamp(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.655.029H7.312c-.765 0-3.272 6.932-3.272 6.932h10.941S12.379.029 11.655.029m-7.634 8v.929h4.981v5.084H6.396c-1.333 0-1.333 1.925-1.333 1.925h8.896s0-1.925-1.333-1.925H9.959V8.958h2.084v2h.875v-2h2.041v-.929z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LampDesk(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M13.521 7.989c-.258 0-.497.072-.708.19l-1.541-2.593l.376-.341l-1.579-1.743l-1.007 1.805l1.135 1.254l.358-.325l1.607 2.66a1.498 1.498 0 0 0 .243 1.569l-1.229 3.929a6.604 6.604 0 0 0-2.137-.361c-2.72 0-5.038-.375-5.974 1.916h11.946c-.544-1.33-1.562-.423-2.843-1.116l1.19-3.859c.053.006.104.017.16.017c.21 0 .409-.046.59-.126c.303-.134.548-.37.704-.662c.115-.212.186-.452.186-.711a1.488 1.488 0 0 0-1.477-1.503m-3.388-6.023L8.346-.009L3.113 2.266L7.354 6.95z"/><path d="M4.473 5.323c.999 1.104 1.882.47 1.882.47L4.177 3.386s-.706.833.296 1.937"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lavabo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M13.969 5.021H4.031c0 1.643 2.489 3.646 4.866 3.893l-.849 5.697c0 .738.548 1.386 1.286 1.386h3.299c.738 0 1.336-.598 1.336-1.336z"/><path d="M12.062 5.027V1.523c0-.396-.219-.537-.607-.537h-1.09l-.412.967l-.911-.284l.653-1.328a.487.487 0 0 1 .448-.296h1.499c.789 0 1.299.459 1.299 1.168v3.814z"/><path d="M11 3h1.223v.937H11z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LawHammer(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.44 7.389L9.662 2.61l.57-.569l-.917-.919l-4.247 4.245l.92.918l.53-.53l1.955 1.954l-8.535 8.535l.83.829l8.533-8.534l1.994 1.993l-.529.531l.918.917l4.254-4.253l-.917-.918z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutFour(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.956 6.326c0 .35-.299.632-.669.632H1.689c-.371 0-.67-.282-.67-.632V.633c0-.35.299-.633.67-.633h8.598c.37 0 .669.283.669.633zM17 15.333a.669.669 0 0 1-.671.667H7.714a.668.668 0 0 1-.671-.667V8.666c0-.368.3-.666.671-.666h8.615c.37 0 .671.298.671.666zm-11.024 0c0 .369-.3.667-.668.667H1.691a.668.668 0 0 1-.671-.667V8.666c0-.368.299-.666.671-.666h3.617c.368 0 .668.298.668.666zM17 6.345c0 .351-.324.634-.726.634h-3.528c-.398 0-.723-.283-.723-.634V.635c0-.352.324-.635.723-.635h3.528c.401 0 .726.283.726.635z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutOne(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M17 16H1V0h16zM2 15h14V1H2z"/><path d="M3 2h6v12H3zm7 0h5v6h-5zm0 7h5v5h-5z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutThree(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M17.021 16.979h-16V1h16zM2 16h14V2H2z"/><path d="M3 10v4.953h4.977V10zm7 0v4.992h5.002V10zM3 3v4.96h5.01V3zm7 0v4.963h5V3z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LayoutTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M17 16H1.021V.042H17zm-1.042-.957V1h-14v14.043z"/><path d="M3 2h5v12H3zm7 0h5v12h-5z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Leaf(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M12.295 14.201a5.16 5.16 0 0 0 .381-.354c-1.989-.903-3.498-1.999-4.858-3.108c-1.624-.004-3.298-.418-4.155-1.477c1.277.643 2.621.646 3.074.585c-3.839-3.466-4.902-6.855-4.902-6.855c1.314 2.06 2.701 3.652 4.059 4.916c-.569-1.088-.236-3.1-.236-3.1c1.107 3.898 1.237 3.96 1.164 3.994c.9.729 1.875 1.414 2.687 1.937c-.292-.68-.499-1.671-.472-3.369c0 0 .722 3.562 1.856 4.439c.903.521 1.698.925 2.314 1.232c.976-1.456.981-5.338-1.411-6.897C9.218 4.462 4.871 4.398.474.096c-1.315-1.287 1.129 10.036 4.67 13.193c2.176 1.945 5.008 2.701 7.151.912"/><path d="M11.266 14.064s1.18.613 2.534 1.211c1.354.598 2.003.477 2.079.625c.078.15.039-1.642.039-1.642s-1.828.433-3.863-.696z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftJustify(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M17 2.938c0 .518-.42.938-.938.938H1.98A.938.938 0 0 1 1.98 2h14.082c.518 0 .938.42.938.938m0 3c0 .518-.42.938-.938.938H1.98A.938.938 0 0 1 1.98 5h14.082c.518 0 .938.42.938.938m0 3c0 .518-.42.938-.938.938H1.98A.938.938 0 0 1 1.98 8h14.082c.518 0 .938.42.938.938m0 3c0 .518-.42.938-.938.938H1.98a.938.938 0 0 1 0-1.876h14.082c.518 0 .938.42.938.938m-5 3c0 .518-.42.938-.938.938H1.98a.938.938 0 0 1 0-1.876h9.082c.518 0 .938.42.938.938"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftwardsArrowToBar(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.994 1c0-.553-.442-1-.989-1H3.026a.994.994 0 0 0-.99 1v14c0 .553.443 1 .99 1h1.979a.994.994 0 0 0 .989-1zm1.444 8.052a1.49 1.49 0 0 1 0-2.104L13.882.506c.581-.582 2.103-.839 2.103 1v12.988c0 1.901-1.521 1.582-2.103 1.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightAlarm(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 12h11v2H2zm10-1H3V9.667C3 7.826 3.75 5 7.5 5S12 7.826 12 9.667zM7 2h1v2H7zM2.597 3.44l.707-.708l1.414 1.414l-.707.707zm10.413-.293l.707.707l-1.414 1.414l-.707-.707zM13 8h2v1h-2zM0 8h1.75v1H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightBulb(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13 3.991C13 1.422 10.729.015 8 .015S3 1.421 3 3.991c0 3.299 2.087 4.197 3.312 6.278c.264.449-.49.227-.235.683h3.818c.252-.452-.595-.229-.333-.676C10.782 8.192 13 7.375 13 3.991M6 12h3.953v.922H6zm3.969 2.968c0 .561-.434 1.017-.971 1.017H7.014c-.538 0-.972-.456-.972-1.017v-.947h3.927z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LightHouse(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.042 0S4.073 1.93 4.073 3h7.938c0-1.07-3.969-3-3.969-3M12 5v.982L16 7V4zM3.958 5L0 4v3l3.958-1zM5 9.969h6V9h1V8H4v1h1zM5 14h6v2H5zm0-3h6v2H5zm0-7h.968v3H5zm5 0h.968v3H10zM7 4h1.984v3H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Like(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.811 13.958H2.194a1.135 1.135 0 0 1-1.146-1.122V7.128c0-.62.515-1.123 1.146-1.123h2.617c.634 0 1.148.503 1.148 1.123v5.708c0 .62-.515 1.122-1.148 1.122m10.127-3.009v1.084h1.514c-.076 1.146-.658 1.897-1.74 1.897h-4.426c-.688 0-1.029-1.312-2.699-1.312l-.558.019V6.681s1.063-.166 1.419-1.291c0 0 1.451-3.961 2.57-4.356c.658 0 1.191-.047 1.191 1.125l-.595 1.814s-.353 2.032 2.14 2.032h2.145c.688 0 1.06.424 1.06 1.049c0 0 .014.357.007.896h-2.027v1.084h1.99a16.57 16.57 0 0 1-.218 1.916h-1.773z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LikeUnlike(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.811 2.035H2.194c-.632 0-1.146.502-1.146 1.123v5.707c0 .621.515 1.123 1.146 1.123h2.617c.634 0 1.148-.502 1.148-1.123V3.158c0-.621-.515-1.123-1.148-1.123m10.127 3.01V3.961h1.514c-.076-1.146-.658-1.898-1.74-1.898h-4.426c-.688 0-1.029 1.312-2.699 1.312l-.558-.018v5.955s1.063.166 1.419 1.291c0 0 1.451 3.961 2.57 4.357c.658 0 1.191.047 1.191-1.125l-.595-1.814s-.353-2.033 2.14-2.033h2.145c.688 0 1.06-.424 1.06-1.049c0 0 .014-.357.007-.895h-2.027V6.96h1.99a16.57 16.57 0 0 0-.218-1.916z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LineTwoAnglePoint(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.938.938h1.084v1.125H7.938zm1.083 14.124H7.937v-1.125h.073v.011h.917v-.011h.094zm.938-2.021H8.928V3h1.031V.041H7V3h1.011v10.041H7V16h2.959z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkOne(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M5 7h6.979v.959H5z"/><path d="M14.604 5h-4.219c-.743 0-1.35.591-1.35 1.318v2.344c0 .726.606 1.317 1.35 1.317h4.219c.744 0 1.348-.592 1.348-1.317V6.318C15.951 5.591 15.348 5 14.604 5m.42 3.679c0 .192-.171.348-.381.348h-4.318c-.21 0-.38-.155-.38-.348V6.302c0-.192.17-.349.38-.349h4.318c.21 0 .381.156.381.349zM6.621 5H2.377c-.749 0-1.357.591-1.357 1.318v2.344c0 .726.608 1.317 1.357 1.317h4.244c.748 0 1.356-.592 1.356-1.317V6.318C7.978 5.591 7.369 5 6.621 5m.428 3.662c0 .189-.172.343-.381.343H2.35c-.209 0-.381-.153-.381-.343V6.318c0-.189.172-.344.381-.344h4.318c.209 0 .381.154.381.344z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkThree(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.969 2c0-1.104-.894-2-1.996-2c-1.102 0-1.959.896-1.959 2c0 .723.408 1.332.986 1.683v3.338h-3.312C10.345 6.42 9.722 6 8.982 6c-.741 0-1.364.42-1.708 1.021H3a1 1 0 0 0-1 1v4.269c-.588.353-1.006.981-1.006 1.711a2 2 0 0 0 1.996 2a2 2 0 0 0 1.998-2c0-.719-.412-1.326-.988-1.678V8.938h3.252C7.59 9.562 8.225 10 8.982 10s1.391-.438 1.729-1.062H15a1 1 0 0 0 1-1V3.684c.574-.351.969-.961.969-1.684"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M1.711 15.919a1.707 1.707 0 0 1-.529-3.332l.32-.107l-.096-.321c-.281-.949-.081-1.882.534-2.494l.63-.627l.912.906l-.425.422c-.714.711-.655 1.583.159 2.396c.434.43.873.648 1.31.648c.389 0 .743-.165 1.086-.504l3.889-3.872l.914.909l-4.098 4.076c-.422.42-.989.642-1.645.642c-.295 0-.594-.046-.894-.138l-.324-.099l-.11.319c-.09.263-.227.488-.408.669a1.72 1.72 0 0 1-1.225.507m10.836-9.892l.195-.195c.24-.237.429-.66.481-1.078c.071-.562-.099-1.078-.479-1.457c-.333-.334-.777-.51-1.279-.51c-.502 0-1.043.189-1.315.463L6.485 6.895l-.93-.925l4.076-4.055c.428-.424 1.004-.648 1.665-.648c.302 0 .609.048.915.142l.328.101l.107-.323c.088-.266.227-.493.412-.678a1.73 1.73 0 0 1 2.438-.008a1.711 1.711 0 0 1-.008 2.426a1.704 1.704 0 0 1-.674.407l-.324.108l.102.327c.315.991.125 1.956-.502 2.582l-.609.605z"/><path d="M1.93 6.23c-.607-.605-.815-1.521-.554-2.452l.087-.313l-.308-.108A1.696 1.696 0 0 1 .529.537A1.727 1.727 0 0 1 2.967.529c.174.173.305.383.393.626l.111.31l.316-.091c.283-.083.569-.126.85-.126c.655 0 1.227.224 1.652.646l.689.688l-.903.9l-.381-.381c-.161-.159-.6-.53-1.222-.53c-.467 0-.911.208-1.322.618c-.605.602-.652 1.169-.584 1.538c.08.431.334.758.533.956l3.803 3.783l-.86.854zm12.318 9.683c-.459 0-.889-.178-1.213-.498a1.682 1.682 0 0 1-.414-.698l-.102-.325l-.328.094a3.007 3.007 0 0 1-.801.113c-.66 0-1.266-.241-1.705-.678l-.58-.579l.857-.852l.216.216c.249.247.76.412 1.269.412c.494 0 .928-.158 1.217-.447c.305-.303.473-.765.473-1.302c0-.509-.152-.988-.389-1.221l-3.667-3.65l.901-.897l4.035 4.013c.617.615.816 1.56.53 2.523l-.097.33l.33.103a1.684 1.684 0 0 1 1.203 1.617c0 .461-.184.894-.511 1.221a1.725 1.725 0 0 1-1.224.505"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Load(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.968 1.866c0 1.012-.437 4.061-.976 4.061c-.537 0-.977-3.049-.977-4.061c0-1.014.439-1.835.977-1.835c.539 0 .976.822.976 1.835M8.016 14.088c0-1.006.438-4.031.977-4.031c.538 0 .975 3.025.975 4.031c0 1.004-.437 1.818-.975 1.818c-.54 0-.977-.814-.977-1.818M2.859 7.03c1.025 0 4.113.431 4.113.961s-3.088.959-4.113.961C1.832 8.95 1 8.521 1 7.991s.832-.96 1.859-.961m12.264 1.952c-1.021 0-4.092-.44-4.092-.983c0-.543 3.07-.983 4.092-.983c1.019-.002 1.846.44 1.846.983c0 .543-.827.983-1.846.983m-.941-4.789c-.715.715-3.18 2.562-3.561 2.182c-.38-.381 1.465-2.848 2.18-3.562c.717-.717 1.609-.986 1.989-.607c.382.382.108 1.271-.608 1.987M3.816 11.846c.711-.712 3.16-2.542 3.541-2.16c.381.38-1.45 2.828-2.162 3.539c-.709.71-1.594.977-1.975.597c-.38-.383-.114-1.267.596-1.976m1.36-9.028c.726.725 2.604 3.213 2.229 3.588S4.544 4.9 3.817 4.177c-.726-.728-1.011-1.619-.636-1.994c.375-.375 1.268-.091 1.995.635m7.636 10.38c-.723-.722-2.582-3.205-2.198-3.589c.384-.385 2.867 1.476 3.588 2.198c.723.719.994 1.616.61 2.001c-.384.383-1.281.11-2-.61"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lock(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13 4.017C13 1.802 11.252 0 9.104 0H7.927C5.779 0 4.031 1.802 4.031 4.017V7.33a5.456 5.456 0 0 0 4.448 8.615a5.458 5.458 0 0 0 5.458-5.46A5.38 5.38 0 0 0 13 7.464zM9.104 2C10.15 2 11 2.904 11 4.017v1.646a5.398 5.398 0 0 0-2.521-.636c-.919 0-1.782.229-2.542.63V4.016C5.938 2.904 6.881 2 7.927 2zm-.625 9.846c.363 0 .693-.147.937-.379l2.578-.753a3.528 3.528 0 0 1-3.515 3.312a3.539 3.539 0 1 1 0-7.076A3.532 3.532 0 0 1 12 10.299l-2.573-.781a1.352 1.352 0 0 0-.947-.388a1.358 1.358 0 0 0-.001 2.716"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LockUnlock(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.927 0c-2.148 0-3.896 1.802-3.896 4.017v1.662a5.403 5.403 0 0 0-2.552-.651a5.457 5.457 0 0 0-5.458 5.458a5.458 5.458 0 1 0 10.916 0a5.38 5.38 0 0 0-1-3.106V4.018c0-1.112.943-2.017 1.989-2.017h.177c1.046 0 1.896.904 1.896 2.017v2.883h2V4.018C17 1.802 15.252 0 13.104 0zM7.416 11.467l2.578-.753a3.528 3.528 0 0 1-3.515 3.312a3.539 3.539 0 1 1 0-7.076A3.532 3.532 0 0 1 10 10.299l-2.573-.781a1.352 1.352 0 0 0-.947-.388a1.358 1.358 0 0 0-.001 2.716c.363 0 .693-.147.937-.379"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LouderSpeaker(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M.031 4.571V7.25C.031 8.118.697 9 1.521 9h6.447V3H1.521c-.824 0-1.49.703-1.49 1.571m7.032 11.367H4.025L2.062 9.062H5.1zM13 0L9 2.76v5.981L13 12zm0 3.989V8c1.086 0 1.969-.897 1.969-2.006c0-1.107-.883-2.005-1.969-2.005"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LouderTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" transform="translate(4)"><ellipse cx="4.962" cy="10.954" rx="1.962" ry="1.954"/><path d="M8.547.014H1.465C.675.014.031.612.031 1.351v13.295c0 .737.644 1.337 1.434 1.337h7.082c.789 0 1.433-.6 1.433-1.337V1.351C9.979.611 9.336.014 8.547.014M4.986 1.948c1.144 0 2.066.927 2.066 2.069a2.066 2.066 0 1 1-4.131 0c0-1.142.926-2.069 2.065-2.069m.006 12.083a3.038 3.038 0 0 1-3.047-3.033a3.038 3.038 0 0 1 3.047-3.031a3.039 3.039 0 0 1 3.047 3.031a3.04 3.04 0 0 1-3.047 3.033"/><ellipse cx="4.991" cy="3.978" rx=".991" ry=".978"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Magnet(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 0h3v2.953H1zm13 0h3v2.953h-3zM9.016 15.917c-4.334 0-7.982-3.344-7.982-7.454V4.031H3.97v3.886c0 2.908 2.23 5.13 5.047 5.13c2.826 0 5.016-2.315 5.016-5.13V4.031h2.936v4.432c-.002 4.04-3.576 7.454-7.953 7.454"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Magnifier(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M17 5.954C17 2.665 14.317 0 11.009 0C7.698 0 5.016 2.665 5.016 5.954c0 3.287 2.683 5.952 5.993 5.952c3.308 0 5.991-2.665 5.991-5.952m-11.066.065A5.082 5.082 0 0 1 11.026.943a5.08 5.08 0 0 1 5.088 5.076a5.08 5.08 0 0 1-5.088 5.075c-2.813 0-5.092-2.272-5.092-5.075m-3.112 9.945L1 14.142l4.037-4.038s.096.765.58 1.247c.482.484 1.242.576 1.242.576z"/><path d="M14.398 5.073c0 .572.44.356.44-.439c0-1.37-1.109-2.48-2.479-2.48c-.797 0-1.012.439-.439.439a2.479 2.479 0 0 1 2.478 2.48"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagnifierTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M16.025 7.5c0-4.143-3.356-7.5-7.499-7.5a7.499 7.499 0 0 0-7.499 7.5a7.5 7.5 0 0 0 7.5 7.5c2.219 0 7.5-.052 7.5-.052zm-7.553 5.529a5.506 5.506 0 1 1 .002-11.012a5.506 5.506 0 0 1-.002 11.012m6.487.929h-1v-1h1z"/><path d="M7.844 3.044c-2.119 0-3.839 1.616-3.839 3.608c0 .25.026.496.077.73c.186.84.529.691.529-.158c0-1.998 1.719-3.609 3.84-3.609c.905 0 .608-.571-.607-.571"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mail(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.304 3.059H1.701a.65.65 0 0 0-.648.648v8.617a.65.65 0 0 0 .648.648h14.603a.65.65 0 0 0 .649-.648V3.707a.65.65 0 0 0-.649-.648m-1.398 8.987l-2.884-3.403l-3.009 2.545L5.955 8.57l-2.862 3.477H1.847l3.189-4.353l-3.07-2.6l-.029-1.281l7.076 5.531l7.049-5.62v1.37l-3.017 2.6l3.124 4.338z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailBox(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.51.053a1.47 1.47 0 0 0-1.498 1.442c-.002.116.012.228.035.336L12.99 2.893C12.941.986 11.521.025 9.969.025c-1.283 0-2.926 1.128-2.926 3.133c0 .812.041 3.78.041 3.78h.988v3.567L4.578 6.938h1.379l-.039-3.864C5.918 1.153 6.834.025 8 .025H4c-1.906 0-3 1.034-3 3.038v3.875h2.41l4.662 4.732V16h1.846V6.938H13V4.184l1.535-1.541c.248.211.562.348.914.354a1.468 1.468 0 0 0 1.498-1.441A1.472 1.472 0 0 0 15.51.053m-4.467 3.114H8.918V2.001h2.125z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailEmpty(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.263 5.806L8.996.063L1.689 5.806c-.358 0-.648.3-.648.667v8.857c0 .367.29.667.648.667h14.573c.358 0 .648-.3.648-.667V6.473a.656.656 0 0 0-.647-.667m-1.279 9.225l-2.826-3.945l-3.162 2.622l-3.134-2.607l-2.878 3.931h-.995L5 10.207L2.045 7.156l.014-.669l6.938-5.228L15.95 6.52v.636L13.014 10.2l3.002 4.817z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailHasMail(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M3.004 0v7l5.027 4l4.924-4V3.031L8 3.026V0zm7.965 8.016h-6V6.969h6zM8.953 5v1h-4V5z"/><path d="M9 0v1.951h3.848zm6.34 4.963l-2.566-1.896v1.105l2.32 1.986v.654l-3.031 2.789L15.079 14l-1.094.015l-2.59-3.739l-3.394 2.786l-3.262-2.772l-2.707 3.726H.923l3.135-4.414l-3.15-2.82l.015-.69L5.19 2.251L5.174.892L.725 4.964c-.359 0-.707.291-.707.65v8.649c0 .36.348.706.707.706h14.617c.359 0 .65-.346.65-.706V5.614a.654.654 0 0 0-.652-.651"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailInbox(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" fill-rule="evenodd"><path fill="currentColor" d="M6.031 4.017L9 6.97l2.969-2.953H9.984V1.032H8.047v2.985z"/><path d="M16.062 7.094v-1.37l-7.048 5.62l-7.076-5.532l.029 1.282l3.07 2.599l-3.189 4.353h1.246l2.862-3.477l3.058 2.619l3.008-2.545l2.884 3.403l1.264-.015l-3.124-4.338z"/><path fill="currentColor" d="M16.304 5.059h-3.242l-4 4.004l-4.103-4.004H1.701a.65.65 0 0 0-.648.648v8.617a.65.65 0 0 0 .648.648h14.603a.65.65 0 0 0 .649-.648V5.707a.65.65 0 0 0-.649-.648m-1.398 8.987l-2.884-3.403l-3.009 2.545l-3.058-2.618l-2.862 3.477H1.847l3.189-4.353l-3.07-2.6l-.029-1.281l7.076 5.531l7.049-5.62v1.37l-3.017 2.6l3.124 4.338z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailSend(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" fill-rule="evenodd"><path fill="currentColor" d="M11.969 3.985L9 1.032L6.031 3.985h1.985V6.97h1.937V3.985z"/><path d="M16.062 7.094v-1.37l-7.048 5.62l-7.076-5.532l.029 1.282l3.07 2.599l-3.189 4.353h1.246l2.862-3.477l3.058 2.619l3.008-2.545l2.884 3.403l1.264-.015l-3.124-4.338z"/><path fill="currentColor" d="M16.304 5.059h-5.242v3.004H6.959V5.059H1.701a.65.65 0 0 0-.648.648v8.617a.65.65 0 0 0 .648.648h14.603a.65.65 0 0 0 .649-.648V5.707a.65.65 0 0 0-.649-.648m-1.398 8.987l-2.884-3.403l-3.009 2.545l-3.058-2.618l-2.862 3.477H1.847l3.189-4.353l-3.07-2.6l-.029-1.281l7.076 5.531l7.049-5.62v1.37l-3.017 2.6l3.124 4.338z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Male(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.875 5.884V.125h-5.762l2.223 2.221l-3.703 3.711a5.395 5.395 0 0 0-3.131-1.002c-3.004 0-5.447 2.452-5.447 5.467c0 3.016 2.443 5.468 5.447 5.468c3.002 0 5.444-2.452 5.444-5.468a5.45 5.45 0 0 0-1.011-3.161l3.702-3.712zM6.464 14.165a3.534 3.534 0 0 1-3.532-3.53a3.535 3.535 0 0 1 3.532-3.529a3.535 3.535 0 0 1 3.532 3.529a3.534 3.534 0 0 1-3.532 3.53"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ManDoctor(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.957 4.121c0-2.457-1.701-4.059-3.492-4.059c-1.794 0-3.457 1.586-3.457 4.059h-.001C5.014 5.779 6.569 8.92 8.483 8.92c1.913 0 3.464-3.141 3.474-4.799m1.542 6.92a2.463 2.463 0 0 0-2.464 2.46a2.463 2.463 0 0 0 4.926 0a2.462 2.462 0 0 0-2.462-2.46m.536 2.981V15h-1.053v-.978h-1.028v-1.033h1.028v-1.026h1.053v1.026h.989v1.033zm-4.637-1.02c0-1.241.545-2.354 1.406-3.109c.461-.406.434-.37.032-.175c-1.158.562-2.726.726-2.726.726s-2.908-.287-3.765-1.439C.24 9.005.116 14.969.116 14.969h9.775a4.16 4.16 0 0 1-.493-1.967"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapSquare(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.938 6.234v-3.05l-1.52 1.525zm-5.614-2.859L6.84 6.872l2.057 2.059l3.486-3.497zm2.809 2.832L9.682 9.671l2.076 2.062l3.451-3.465zM.042 10.792l3.883-3.917L1.82 4.771L.083 6.5zM.026 4.91l1.362-1.364l-1.362-1.36zM1.227.007a1.11 1.11 0 0 0-.814.36l2.062 2.068L4.903.007zm4.16 5.401L8.695 2.09L6.609 0h-.015L3.289 3.305zM14.883.062L8.485.054L9.581 1.15l.701-.701l3.537 3.537l2.148-2.155l.023-.659a1.11 1.11 0 0 0-1.107-1.11m-.042 15.879c.616 0 1.113-.5 1.113-1.117l.001-4.515l-5.631 5.632zM.177 14.855c0 .617.497 1.117 1.114 1.116l6.24.001l.806-.842l-3.522-3.535l.768-.731l3.522 3.534l1.169-1.179l-4.896-4.913l-5.202 5.201z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapThree(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m11.035 13.581l4.962 1.294V3.415l-4.962-1.353zm-5.014 1.384l3.922-1.384V2.062L6.021 3.879zM.042 13.118l4.95 1.847V3.879L.042 2.031z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MarkSnorker(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M13.555 2.037h-9.07c-1.89 0-3.424 1.402-3.424 3.132v1.605c0 1.73.871 3.133 2.761 3.133h1.862c1.439 0 2.67-.501 3.177-1.649c.071-.166.267-.176.335-.018c.5 1.16 1.729 1.667 3.167 1.667h1.854c1.883 0 2.746-1.402 2.746-3.133V5.169c0-1.73-1.527-3.132-3.408-3.132m2.49 4.919c0 1.157-.293 2.097-1.512 2.097h-1.971c-1.049 0-1.959-.395-2.167-1.368c0 0-.201-.647-1.391-.647c-1.188 0-1.354.644-1.354.644c-.209.978-1.131 1.372-2.186 1.372H3.488c-1.23 0-1.529-.939-1.529-2.097v-1.86c0-1.155 1-2.096 2.231-2.096h9.641c1.221 0 2.213.94 2.213 2.096v1.859z"/><path d="M12.743 7.98h1.36l.705-3.488a2.268 2.268 0 0 0-1.275-.415zm-.579 3.213l-.543 2.481c-.105.525-.857.92-1.375.815c-.504-.102-.949-.418-.869-.927c.929-.28 1.603-1.07 1.591-2.001c-.004-.308-.282-.554-.624-.55c-.34.004-.614.257-.609.565c.006.537-.473.98-1.066.986c-.595.008-1.084-.423-1.091-.961c-.004-.307-.283-.553-.623-.549c-.342.004-.613.258-.61.564c.012.949.73 1.735 1.695 1.977c-.06 1.131.83 1.945 1.957 2.172c1.238.25 2.68-.574 2.932-1.83l.596-2.738a210.32 210.32 0 0 1-1.361-.004M15.643.889a.713.713 0 0 0-.553-.843a.715.715 0 0 0-.836.563l-.213 1.178c.492.043.953.174 1.367.38z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MaskOne(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.435 4.102c-1.33-.415-3.647 1.139-5.439 1.139c-1.801 0-4.127-1.552-5.461-1.133c-1.924.605-2.502-.597-2.502.48c0 0-.193 6.048 3.346 6.886c2 .286 3.445 1.495 4.617 1.495c1.192 0 2.189-1.079 4.667-1.505c3.214-.439 3.298-6.876 3.298-6.876c0-1.082-.59.118-2.526-.486M12.83 8.525c-.854.295-1.686.339-2.256.165c-.354-.108-.615.073-.806.148c-.189.075-.189-1.3.599-1.952c.39-.325.553-.427 1.193-.648c1.385-.477 3.386.254 3.386.254s-.734 1.557-2.116 2.033M3.017 6.492s2.001-.73 3.386-.254c.642.222.805.323 1.193.648c.789.652.789 2.027.6 1.952c-.19-.075-.452-.257-.805-.148c-.572.174-1.403.13-2.258-.165c-1.382-.476-2.116-2.033-2.116-2.033"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MaskTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.836 2.867c.115-.773.125-1.334.125-1.334c0-1.036-.588.112-2.527-.466c-1.33-.397-3.65 1.09-5.443 1.09c-1.802 0-4.129-1.483-5.465-1.084c-1.924.58-2.503-.572-2.503.46c0 0-.19 5.63 3.353 6.433c2.002.273 3.443-1.057 4.615-1.057c1.194 0 2.195 1.454 4.676 1.047c.69-.091 1.226-.925 1.658-1.38l1.257 9.382zM7.081 5.208c-.153-.06-.364-.202-.647-.117c-.458.136-1.128.102-1.815-.129c-1.112-.372-1.702-2.043-1.702-2.043s1.61-.116 2.723.255c.516.175.648.255.96.507c.632.509.634 1.586.481 1.527m4.357-1.535c.313-.252.444-.332.96-.507c1.112-.371 2.723-.255 2.723-.255s-.59 1.671-1.703 2.043c-.686.23-1.355.265-1.813.129c-.283-.085-.494.058-.647.117c-.154.059-.152-1.017.48-1.527"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MedalStar(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="m8.007 5.996l1.071 3.02h3.54L9.74 11.318l1.094 3.489l-2.865-2.157l-2.864 2.157l1.094-3.489l-2.85-2.302h3.54z"/><path d="m8.007 1.947l2.276 6.039H11.5l.469-6.619S12.151.041 10.768.041H5.252c-1.361 0-1.235 1.305-1.235 1.305l.488 6.641h1.229z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicrophoneOne(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="m16.608 13.715l-5.567-6.416c-.488 1.203-1.523 2.176-2.705 2.723l6.416 5.525a1.323 1.323 0 0 0 1.856 0a1.281 1.281 0 0 0 0-1.832M8.736 1.49c-.42 1.39-1.244 2.744-2.389 3.89c-1.14 1.14-2.51 1.96-3.891 2.38a4.455 4.455 0 0 0 3.012 1.178c2.469 0 4.469-1.986 4.469-4.438a4.388 4.388 0 0 0-1.201-3.01"/><path d="M7.877.771A4.46 4.46 0 0 0 5.469.062C3 .062 1 2.048 1 4.5c0 .897.273 1.729.733 2.428c1.364-.314 2.759-1.106 3.907-2.255C6.79 3.523 7.56 2.15 7.877.771"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicrophoneTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="m13.984 7.031l-.04-2.007h-.936l.04 4.629c0 1.928-1.988 3.409-4.529 3.409s-4.535-1.481-4.535-3.409V5.031H3.01v1.985h-.995v.922h.995v2.111c0 1.698 1.741 3.25 3.989 3.75V16h2.979v-2.188c2.275-.485 4.005-2.048 4.005-3.762V7.953h1v-.922z"/><path d="M8.495 10.953c1.914 0 3.468-.869 3.468-1.94V7.016H5.025v1.997c0 1.071 1.554 1.94 3.47 1.94m3.484-7.525C11.979 2.086 10.421 1 8.5 1S5.021 2.086 5.021 3.428V5.99h6.957z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Microscope(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="m11.425 3.096l.408-1.001l-.431-.176l.614-1.509l-.883-.358l-.615 1.507l-.396-.162L8.09 6.386l1.709.696l1.123-2.754c1.349.623 2.289 1.692 2.618 3.013c.549 2.208-.786 4.518-3.085 5.684H3.862c-.502 0-.908.348-.908.777l-.906 1.379c0 .43.405.777.906.777h9.072c.502 0 .908-.348.908-.777v-1.379c0-.246-.142-.456-.35-.598c1.854-1.58 2.792-3.914 2.228-6.181c-.427-1.711-1.669-3.136-3.387-3.927"/><path d="m8.223 6.759l1.748.711l-.238.584l-1.747-.712z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mobile(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.939 2.042V.047l-.893.011l-.005 15.88h6.916V2.022zm1.088 9.979H6.984v-1.053h1.043zm0-2H6.968V8.984h1.059zm1.973 2H8.998v-1.068H10zm.016-2H8.998V8.968h1.018zm2 2h-1.042v-1.068h1.042zm-3.989 2H6.968v-1.053h1.059zm1.989 0H8.998v-1.068h1.018zm2.015 0h-1.089v-1.068h1.089zm0-4h-1.073V8.953h1.073zm.031-2H6.954V3.953h5.108z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mocule(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M17 9.935V6.086c-.043.027-.076.068-.117.099L14.69 3.956c.165-.287.267-.615.267-.97a1.954 1.954 0 1 0-3.908 0c0 .363.104.699.278.991l-2.336 2.3a1.964 1.964 0 0 0-.963-.26c-.375 0-.723.11-1.021.292L4.722 3.987c.16-.284.261-.608.261-.958c0-1.087-.89-1.968-1.985-1.968a1.976 1.976 0 0 0-1.986 1.968a1.977 1.977 0 0 0 3.014 1.677l2.289 2.325a1.933 1.933 0 0 0-.259.958c0 .36.104.694.273.985l-2.354 2.318a1.983 1.983 0 0 0-.963-.257c-1.088 0-1.971.868-1.971 1.94s.883 1.94 1.971 1.94s1.97-.868 1.97-1.94c0-.364-.108-.701-.285-.992l2.337-2.301c.293.172.63.277.994.277c.359 0 .691-.104.981-.271l2.308 2.344a1.9 1.9 0 0 0-.271.968c0 1.071.874 1.938 1.954 1.938A1.946 1.946 0 0 0 14.954 13a1.91 1.91 0 0 0-.284-.992l2.205-2.182c.043.035.079.079.125.109m-4.973 1.395L9.72 8.987c.173-.294.279-.632.279-.998c0-.374-.109-.72-.291-1.018l2.33-2.293c.285.163.611.264.963.264c.36 0 .693-.104.984-.275l2.262 2.297a2.4 2.4 0 0 0-.244 1.046c0 .379.093.727.24 1.036l-2.297 2.271a1.946 1.946 0 0 0-.947-.253c-.355 0-.685.102-.972.266"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Money(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.979 10.596c0-1.695-1.358-3.104-3.011-3.381V3.828c.504.156 1.025.195 1.235.67c.19.428.671.611 1.072.41c.403-.202.573-.714.384-1.143c-.473-1.066-1.542-1.494-2.676-1.688V0H8v2.066c-1.664.251-2.995 1.67-2.995 3.383c0 1.71 1.324 3.127 2.995 3.383v3.41c-.536-.148-1.033-.198-1.252-.695c-.188-.43-.668-.613-1.068-.412c-.404.201-.576.711-.388 1.141c.481 1.09 1.554 1.531 2.724 1.711V16h.938v-2.021c1.646-.272 3.025-1.685 3.025-3.383M6.617 5.449c0-.774.598-1.424 1.413-1.633v3.269c-.815-.211-1.413-.86-1.413-1.636m2.344 6.774V8.969c.815.225 1.41.865 1.41 1.627s-.595 1.402-1.41 1.627"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoneyBag(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M9 11h.922v1.959H9zM7 8h.938v1.974H7z"/><path d="M9.296 4.779c.858-1 1.47-3.257 1.47-4.074c0-1.726-1.029.44-2.298.44c-1.271 0-2.298-2.188-2.298-.44c0 .831.629 3.148 1.512 4.123c-2.608.74-4.664 4.494-4.664 7.794c0 2.918 2.455 3.364 5.482 3.364s5.481-.486 5.481-3.364c0-3.006-2.004-7.201-4.685-7.843m1.735 8.268h-1v.984h-1v1.031H7.969v-1.031h-1v-.984l.013-.016H5.969v-1.062h1.047v1.02l.016-.02h.938v-1.938H7v-.984H5.953V7.985H7V6.969h.969V5.938h1.062v1.031h1v1h.984v1.062H9.968V8.01l-.016.006H9.03v1.953h1.016v1.016h.984v2.062z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoneyCoin(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" transform="translate(1)"><ellipse cx="12.473" cy="5.973" rx="3.473" ry="1.973"/><path d="M12.525 9.081c-2.484 0-3.473-1.041-3.473-1.644v2.583c0 1.09 1.555 1.972 3.473 1.972c1.92 0 3.475-.882 3.475-1.972V7.499c0 .604-.99 1.582-3.475 1.582"/><path d="M12.525 13.072c-2.222 0-3.473-1.001-3.473-1.604v2.559c0 1.09 1.555 1.973 3.473 1.973c1.92 0 3.475-.883 3.475-1.973v-2.59c0 .603-1.252 1.635-3.475 1.635"/><ellipse cx="3.937" cy="1.973" rx="3.937" ry="1.973"/><path d="M4.062 5.081C1.247 5.081.125 4.04.125 3.437V6.02c0 1.09 1.763 1.972 3.937 1.972C6.238 7.992 8 7.11 8 6.02V3.499c0 .604-1.123 1.582-3.938 1.582"/><path d="M4.062 9.072C1.543 9.072.125 8.071.125 7.468v2.559C.125 11.117 1.888 12 4.062 12C6.238 12 8 11.117 8 10.027v-2.59c0 .603-1.42 1.635-3.938 1.635"/><path d="M4.062 13.072c-2.519 0-3.937-1.001-3.937-1.604v2.559C.125 15.117 1.888 16 4.062 16C6.238 16 8 15.117 8 14.027v-2.59c0 .603-1.42 1.635-3.938 1.635"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoneyThree(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M12.345 14.947H1.655c-.9 0-1.634-.703-1.634-1.57V5.616c0-.864.733-1.571 1.634-1.571h10.69c.901 0 1.635.707 1.635 1.571v7.761c-.001.867-.734 1.57-1.635 1.57M.995 6.414v6.211l1.369 1.432h9.33l1.327-1.385V6.46l-1.327-1.524h-9.33z"/><path d="M14.248 2.033H3.047l-.004.936h10.213l1.775 1.814v6.154l.896-.016V3.663c0-.87-.766-1.614-1.679-1.63M2 9h.953v.984H2zm9 0h.984v.953H11zM5 8h.969v.969H5zm4 2h.984v.969H9z"/><path d="M8 6.016h-.969v1.015H6v.938h1.031v1.062H6v.938h1.031v1.062H6v.938h1.031v1H8v-1h.984v-.938H8V9.969h.984v-.938H8V7.969h.984v-.938H8z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoonStar(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M10.705 13.274A6.888 6.888 0 0 1 6.334 1.065C2.748 1.892.072 5.099.072 8.936a8.084 8.084 0 0 0 8.084 8.085c3.838 0 7.043-2.676 7.871-6.263a6.868 6.868 0 0 1-5.322 2.516"/><path d="m12.719 1.021l1.025 2.203l2.293.352l-1.658 1.715l.391 2.42l-2.051-1.143l-2.051 1.143l.391-2.42l-1.661-1.715l2.294-.352z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Motobike(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M7.031 8.958h2.927c0-1.122-.657-1.166-1.464-1.166c-.808 0-1.463.044-1.463 1.166"/><path d="M10.617 1.042C10.179.448 9.397 0 8.496 0c-.902 0-1.684.448-2.119 1.042H4.022v.916h.021v.003h2.04C5.154 2.525 4 2.886 4 4.433s.744 1.498 1.016 1.9v5.28c0 1.515 1.047 2.053 2.047 2.246v.635c0 .428.483 1.506 1.445 1.506c.955 0 1.469-1.078 1.469-1.506v-.642c.988-.199 2.008-.744 2.008-2.239V6.209c.25-.407.995-.271.995-1.896s-1.215-1.816-2.099-2.354h2.099v-.916zM8.496.958c.864 0 1.563.695 1.563 1.553c0 .856-.699 1.552-1.563 1.552a1.556 1.556 0 0 1-1.562-1.552c0-.858.698-1.553 1.562-1.553m2.534 10.699c0 .576-.229.977-1.03 1.175v-2.816H7.039v2.816c-.833-.208-1.07-.633-1.07-1.219V7.304c.693.434 1.584.689 2.555.675c.959-.015 1.831-.29 2.507-.733z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mountain(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.012 1.062L4.035 8.87L2.709 7.569S-.305 14 .063 14h15.902v-.002l-3.338-7.1l-.983.612zM5.611 7.521L8.062 2.77l2.285 4.081l-.986.67l-1.34-1.288z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoviePlay(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M1 15h15l-.083-14H1zM11 2h1v1h-1zm3 2v8H3V4zM8 2h1v1H8zM5 2h1v1H5zM2 2h1v1H2zm.979 12h-1v-1h1zM6 14H5v-1h1zm3 0H8v-1h1zm3 0h-1v-1h1zm3 0h-1v-1h1zm0-11h-1V2h1z"/><path d="M7.003 9.865L7 6l3.949 2.015z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MovieProjector(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m16.443 9.06l-1.441.944v2.989l1.441.944c.287 0 .521-.359.521-.799V9.86c0-.441-.234-.8-.521-.8M9.986 5.975a2.976 2.976 0 0 0 2.967-2.983A2.976 2.976 0 0 0 9.986.008a2.976 2.976 0 0 0-2.969 2.984a2.977 2.977 0 0 0 2.969 2.983m0-5.003c1.108 0 2.006.905 2.006 2.02c0 1.116-.897 2.02-2.006 2.02a2.013 2.013 0 0 1-2.009-2.02A2.015 2.015 0 0 1 9.986.972M3.508 5.987a2.477 2.477 0 1 0 0-4.949a2.477 2.477 0 0 0 0 4.949m0-3.976a1.502 1.502 0 1 1 0 3.004a1.502 1.502 0 0 1 0-3.004m9.041 5.048H2.414C1.633 7.059 1 7.715 1 8.524v6.013c0 .809.633 1.464 1.414 1.464h10.135c.781 0 1.415-.655 1.415-1.464V8.525c0-.81-.634-1.466-1.415-1.466"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MultifunctionKnife(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" fill-rule="evenodd"><path fill="currentColor" d="M11.812 9.886c.551.095.998 2.334 1.808 1.267c.979-1.293 1.683-2.64 2.206-4.46c.958-3.324-3.986-.657-2.273-.827c1.338-.133-.314 3.588-1.68 3.071"/><path fill="currentColor" d="M9.381 1.198c-.07-.136-.348-.256-.521.08c-.405.786-1.002 1.469-.735 4.753L11.969 10s.074-3.674-2.588-8.802m2.739 10.178L6.81 5.791a2.347 2.347 0 0 1-.224.012c-1.28 0-2.319-1.087-2.319-2.426c0-.078.005-.156.012-.234l-.916-.958c-.566-.592-1.436-.53-1.828-.093L.137 3.533c-.371.371-.263 1.486.136 1.884l8.757 9.191c.566.593 1.567.508 2.235-.191l.671-.704c.668-.697.751-1.745.184-2.337m-1.749 2.322a.77.77 0 0 1-.76-.779a.77.77 0 0 1 .76-.781c.424 0 .764.35.764.781a.77.77 0 0 1-.764.779"/><path stroke="currentColor" d="M9.126 13.521H7.457c-1.206 0-1-1-1.999-1c-.999 0-1.17 2-2.001 2s-1.044-2-2-2s-.998 1-.998 1"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mushrooms(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.002 1C4.149 1 1.025 4.446 1.025 7.115c0 2.666 3.124 2.869 6.977 2.869c3.854 0 6.978-.203 6.978-2.869C14.979 4.446 11.855 1 8.002 1m-.08 10.169c-1.047 0-2.032-.019-2.892-.115c-.001.087-.005.172-.005.259c0 3.142-3.579 5.688 2.945 5.688c6.525 0 2.947-2.546 2.947-5.688c.001-.092-.003-.182-.004-.273c-.884.107-1.903.129-2.991.129"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Music(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.942.751L6.035.042v11.097a4.228 4.228 0 0 0-1.924.299c-1.594.651-2.422 2.217-1.965 3.312c.458 1.098 2.029 1.604 3.621.953c1.224-.5 2.073-1.451 2.184-2.362l-.008-9.817c2.627.798 4.52 1.673 4.52 4.284c0 .927 1.52 1.823 1.52-1.812c-.001-3.092-2.653-4.962-6.041-5.245"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicNote(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.021 2.188v9.174a4.101 4.101 0 0 0-1.792.301c-1.605.649-2.533 2.066-2.074 3.162c.465 1.099 2.139 1.459 3.743.809c1.233-.5 1.958-1.45 2.067-2.362l-.007-8.885l7.062-1.359v6.378a4.089 4.089 0 0 0-1.944.297c-1.605.65-2.532 2.067-2.072 3.163c.463 1.098 2.137 1.459 3.742.809c1.233-.501 2.09-1.451 2.201-2.362L16.958.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mustache(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.907 7.514c-1.561 1.638-2.747 1.19-4.694-.629c-1.611-1.503-2.934-.58-3.254.715c-.324-1.297-1.648-2.218-3.258-.712c-1.943 1.821-3.129 2.271-4.691.636c-.155 1.646 1.514 3.188 3.725 3.441c2.049.234 3.846-.719 4.225-2.168c.383 1.451 2.179 2.401 4.229 2.163c2.209-.257 3.875-1.799 3.718-3.446"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Network(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.783 10.094c-1.699.998-3.766 1.684-5.678 1.95a1.662 1.662 0 0 1-.684.934c.512 1.093 1.249 2.087 2.139 2.987a7.983 7.983 0 0 0 6.702-3.074c.029-.038.055-.08.083-.119c-.244-.914-.648-1.784-1.145-2.644c-.09.026-.176.046-.261.062c-.143.04-.291.068-.446.068a1.66 1.66 0 0 1-.71-.164M9.051 5.492a17.647 17.647 0 0 0-2.004-1.256a1.671 1.671 0 0 1-1.907.985c-.407 1.535-.624 3.162-.511 4.694a1.669 1.669 0 0 1 1.52 1.354c1.695-.279 3.47-.879 4.967-1.738a1.672 1.672 0 0 1-.297-.949c0-.413.156-.786.403-1.078c-.654-.736-1.389-1.443-2.171-2.012M4 9.989c-.137-1.634.104-3.392.541-5.032a1.673 1.673 0 0 1-.713-1.369c0-.197.039-.386.104-.562a18.102 18.102 0 0 0-1.974-.247c-.089.104-.185.204-.269.314a7.983 7.983 0 0 0-1.23 7.547a9.487 9.487 0 0 0 2.397.666A1.671 1.671 0 0 1 4 9.989m9.928-.3c-.029.037-.064.067-.096.1c.433.736.799 1.482 1.053 2.268a7.98 7.98 0 0 0 .832-6.122c-.09.133-.176.267-.271.396c-.436.601-.875 1.217-1.354 1.772c.045.152.076.311.076.479v.004c.084.374.013.779-.24 1.103M7.164 3.447c.799.414 1.584.898 2.33 1.44c.84.611 1.627 1.373 2.324 2.164c.207-.092.434-.145.676-.145c.5 0 .945.225 1.252.572c.404-.492.783-1.022 1.161-1.54c.194-.268.372-.543.544-.82A7.96 7.96 0 0 0 7.701.012c-.115.146-.229.29-.339.439c-.401.552-.739 1.08-1.04 1.637c.039.029.064.066.1.1c.417.276.697.734.742 1.259m-4.285 8.518a10.141 10.141 0 0 1-2.07-.487a7.954 7.954 0 0 0 5.806 4.397a10.969 10.969 0 0 1-1.753-2.66a1.675 1.675 0 0 1-1.983-1.25m1.635-9.723a1.32 1.32 0 0 1 1.199-.416C6.025 1.24 6.377.683 6.794.104a7.974 7.974 0 0 0-4.247 2.062c.59.066 1.176.14 1.761.252c.063-.064.133-.121.206-.176"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NetworkTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.494 15.958C4.361 15.958 1 12.622 1 8.52s3.361-7.438 7.494-7.438c4.133 0 7.495 3.336 7.495 7.438s-3.362 7.438-7.495 7.438m.016-14.02A6.58 6.58 0 0 0 1.938 8.51a6.58 6.58 0 0 0 6.572 6.573a6.58 6.58 0 0 0 6.573-6.573A6.58 6.58 0 0 0 8.51 1.938"/><path d="M8 2h.922v14.084H8z"/><path d="M1 8h13.96v.922H1zm1-3h12.406v.906H2zm0 6h12.406v.922H2z"/><path d="M8.317 15.854c-2.597-1.273-4.274-4.192-4.274-7.437c0-3.17 1.623-6.062 4.138-7.367l.461.887c-2.187 1.137-3.599 3.68-3.599 6.48c0 2.865 1.459 5.432 3.714 6.538z"/><path d="m8.74 15.789l-.469-.883c2.139-1.134 3.521-3.681 3.521-6.489c0-2.775-1.359-5.31-3.46-6.457l.479-.877c2.418 1.318 3.981 4.197 3.981 7.334c.001 3.175-1.591 6.068-4.052 7.372"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Newspaper(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.031 2.046v9.788c0 .386-1.062.389-1.062.016V6.063H1.022v6.918c-.004.104.031.991.979.991l14-.012s.973-.049.973-.96V2zm4.011 8.016h-2.07V6.937h2.07zM15 8H9V6.958h6zm0 2.083H8.981V8.939l6.019.02zm0 1.955H5.987v-1.08H15zm0-5.997H6V3.937h9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoDog(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="m13.125 5.149l-.342-1.128l-2.163 2.138l1.081 1.127l2.294.493l.97-.712zm-3.061 1.926H5.451c-.018 0-2.242-.952-2.576-1.096v.444l1.502 1.372c-.195.241-.332.503-.332.704v3.417l1.378.104v-1.207l1.878-.95h2.342l.469 2.054h.88V8.015z"/><path d="M9.01 15.958c-4.405 0-7.989-3.565-7.989-7.948C1.021 3.627 4.605.063 9.01.063c4.405 0 7.989 3.564 7.989 7.947s-3.584 7.948-7.989 7.948M8.979 1.073c-3.902 0-7.077 3.135-7.077 6.989s3.175 6.989 7.077 6.989s7.077-3.136 7.077-6.989c0-3.854-3.175-6.989-7.077-6.989"/><path d="m11.157 1.314l1.207.58l-6.348 13.213l-1.207-.58z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoSmoke(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M4 8h1.973v.965H4zm3 0h5.938v.923H7zm5.816-2.037c-.359-.27-.693-.332-1.094-.448c-.885-.256-1.281-1.3-1.322-2.114c-.117.546-.06 1.098.015 1.459c.11.542.414 1.051.903 1.324c.412.227.958.246 1.398.472c.426.217.78.509.907.692c.001 0-.229-.95-.807-1.385"/><path d="M8.995.067c-4.392 0-7.964 3.563-7.964 7.946c0 4.381 3.572 7.944 7.964 7.944c4.39 0 7.963-3.563 7.963-7.944c0-4.382-3.573-7.946-7.963-7.946m7.111 8.023a7.1 7.1 0 0 1-1.601 4.496L4.467 2.534A7.122 7.122 0 0 1 8.955.943c3.943 0 7.151 3.205 7.151 7.147M1.967 7.986c0-1.72.615-3.3 1.635-4.528l9.961 9.993a7.042 7.042 0 0 1-4.512 1.628c-3.907 0-7.084-3.183-7.084-7.093"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Note(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.938 0H3a2 2 0 0 0-2 2v2h16l-.062-2a2 2 0 0 0-2-2M1 8h4v2H1zm0 3h4v2H1zm4 5v-2H1c.066 1.045.927 2 1.987 2zM1 5h4v2H1zm5 0v2h11l-.062-2zm0 6v2h11l-.062-2zm8.938 5C16 16 16.935 15.045 17 14H6v2zM6 8v2h11l-.062-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoteTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.016 0v1.031h-1.062V0h-.895v1.031h-1.09V0h-.953v1.031H9.954V0h-.922v1.031H7.941V0h-.925v1.031H5.985V0h-.942v1.031H3.959V0H3v16h12.954V0zM5 6.958h9v1H5zm9 5H5v-1h9zM14 10H5V9h9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OldPhone(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M13.936 8.888C13.936 6.714 10.752 6 10.752 6c-.324 0-.748-.002-1.752 0V4.958l.993.013V3.083H9v1H6v-1H5v1.888l1.042-.013V6H4.176S.999 6.637.999 8.879c0 2.243-.944 5.779-.944 5.779s.736.048.736 1.31h13.333c0-1.208.858-1.315.858-1.315s-1.046-3.591-1.046-5.765m-6.415 6.195a3.605 3.605 0 1 1 .001-7.21a3.605 3.605 0 0 1-.001 7.21"/><path d="M1.273 4.879h1.305s1.266.17 1.305-.818H.049c0 .925 1.224.818 1.224.818M4 2.291C4 1.951 4.82 2 5.165 2h4.714c.342 0 1.136.049 1.136.389v.503h3.947c0-.01.002.008.002 0c0-2.938-5.77-2.841-7.793-2.841C5.15.051.027-.023.027 2.917H4c.002-.021 0-.812 0-.626m8.441 2.588h1.283s1.222.192 1.237-.818h-3.825c.048.88 1.305.818 1.305.818M8 11h2v.964H8z"/><ellipse cx="7.477" cy="11.492" rx="1.477" ry="1.492"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Open(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" fill-rule="evenodd"><path d="M6.969 7h-1v2.031h1v-.062H7V7z"/><path fill="currentColor" d="M5.969 7h1v-.979L5.031 6v5.958h.938v-2.02h1v-.907h-1zM7 7h1v1.969H7z"/><path d="M2.984 7.016v-.021H1v.021H.984v3.937H1v.078h1.984v-.078H3V7.016z"/><path fill="currentColor" d="M3 7h.947v3.938H3zM1 6h1.984v.943H1zm0 5h1.984v.953H1zM0 7h.949v3.938H0zm15.031-.969V7.5l-1.062-.875v-.594h-.953v5.938h.953V8.031l1.062.875v3.063h.922V6.031zm-4.947-.005H9.016v5.963h1.068v-.02h1.885v-.938H9.953V8.969h2.016v-.938H9.953V6.953h2.016v-.906h-1.885z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pallette(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.229 8.398c-1.439-1.932 5.643-7.369.205-7.369c-4.611 0-8.351 3.488-8.351 7.792c0 8.073 8.494 9.594 10.733 6.874c3.217-3.905-1.236-5.485-2.587-7.297M9.521 5.031a1.562 1.562 0 0 1 0-3.124a1.561 1.561 0 0 1 0 3.124M6.52 8.034a1.54 1.54 0 1 1 0-3.078a1.54 1.54 0 0 1 1.544 1.54A1.541 1.541 0 0 1 6.52 8.034m-.977 4.056a1.572 1.572 0 0 1-1.575-1.57c0-.867.704-1.57 1.575-1.57c.869 0 1.574.703 1.574 1.57c0 .866-.705 1.57-1.574 1.57m2.985 3.087a1.613 1.613 0 0 1-1.606-1.618a1.61 1.61 0 0 1 1.606-1.617a1.61 1.61 0 0 1 1.606 1.617a1.612 1.612 0 0 1-1.606 1.618m3.477-1.96a1.16 1.16 0 0 1-1.161-1.163a1.163 1.163 0 1 1 1.161 1.163"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaperClip(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.346 16C5.009 16 4 14.907 4 13.725V3.998C4 1.634 5.254 0 7.698 0h.67C11.045 0 12 1.56 12 3.998v7.007h-.954V3.998C11.046 2.414 10.409 1 8.367 1h-.685C5.872 1 5 2.318 5 3.998v9.727c0 .738.448 1.274 1.345 1.274h1.338c.852 0 1.379-.488 1.379-1.274V5.756c0-.531-.081-.716-1.119-.765c-1.059.052-.943.271-.943.765v4.254H5.999V5.756c0-1.121.636-1.696 1.944-1.758C9.249 4.058 10 4.616 10 5.756v7.969C10 14.947 8.966 16 7.682 16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaperPlane(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m17 1.042l-5.564 13.912l-3.478-3.477l.695 2.086l-1.623 1.395v-3.395l7.954-8.188l-8.937 6.594L1 8.694z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaperRoll(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M2.967 2C1.783 2 .861 2.524.365 3.476a3.313 3.313 0 0 0-.363 1.421v8.436s1.187-1.345 3.272 0c2.677 1.599 3.579-.771 4.737 0c1.845 1.43 2.915 0 2.915 0s-.013-6.852 0-7.031c.07-1.144.226-2.164.661-2.926c.485-.849 1.148-1.375 1.881-1.375H2.967z"/><path d="M15.029 3.713c-.297-.409-.643-.65-1.017-.65c-.53 0-1.011.482-1.364 1.26c-.314.696-.526 1.63-.58 2.676v1.004c.112 2.215.938 3.936 1.944 3.936c1.081 0 1.958-1.986 1.958-4.438c.001-1.606-.376-3.009-.941-3.788m-1.016 6.334c-.53 0-.966-.991-1.024-2.264v-.58c.029-.601.14-1.139.305-1.536c.187-.453.439-.729.72-.729c.195 0 .377.137.533.377c.298.446.496 1.256.496 2.179c-.001 1.41-.464 2.553-1.03 2.553"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaperShredder(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M12.969 3.003H7.985V0H3v6.982h9.969z"/><path d="M9.111 0v1.938h3.858zM3 13v2.986h.978V13zm2.013 0v3h.974v-3zM7 13v2.998h.998V13zm2 0v2.993h.986V13zm2 0v2.964h.963V13zM0 8.031V12h15.958V8.031zm2.969 2h-1V9h1zm1 0V9h1.062v1.031z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pause(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.002 1c0-.553.444-1 .993-1h3.972c.549 0 .993.447.993 1v14c0 .553-.444 1-.993 1H1.995a.996.996 0 0 1-.993-1zm8 0c0-.553.444-1 .993-1h3.972c.549 0 .993.447.993 1v14c0 .553-.444 1-.993 1H9.995a.996.996 0 0 1-.993-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pawn(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 16h12l-3-4V5l2-2V0h-1.938v1H9.937V0H8.062v1H5.937V0H4v3l2 2v7zm7-11h1v7h-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paypal(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M9.875.074H3.639L.965 12.917h3.31l.942-3.019h4.658c2.272 0 4.116-2.035 4.116-4.543v-.907c0-2.509-1.844-4.374-4.116-4.374m.25 5.138c0 1.033-.87 1.87-1.373 1.87l-2.928.012l.832-4.129h2.096c.503 0 1.373.838 1.373 1.871z"/><path d="M15.113 3.868c0 .77-.052 1.471-.052 2.11c0 3.211-1.308 5.104-4.222 5.104H6.557l-1.571 4.907h2.639l.943-3.021h4.281c2.273 0 4.115-1.721 4.115-4.213v-.902c.001-1.727-.549-3.226-1.851-3.985"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pen(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m15.224-.387l1.869 1.868l-3.035 3.035l-1.87-1.869zM11.252 3.48l1.876 1.876l-7.115 7.115l-1.876-1.876zM1.021 15.957l2.143-4.109l1.877 1.877zm10.534-5.996l-.468-.477l2.965-2.919l-.393-.434l.496-.448l.822.91z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PenNib(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m10.723.04l-.785 1.843l4.25 4.25l1.83-.773zM3.357 6.132s.94 4.213-3.294 8.447l.302.302l6.764-6.764a1.492 1.492 0 0 1-.145-.633a1.5 1.5 0 1 1 1.5 1.5c-.235 0-.455-.059-.654-.156l-6.758 6.76l.389.389c4.249-4.248 8.463-3.292 8.463-3.292l2.758-6.054l-3.295-3.295z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pencil(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.479 2.536L14.474.546c-.539-.537-1.383-.565-1.884-.064L3.47 9.616s-2.312 5.32-2.469 5.75c-.125.34.306.771.604.625c.48-.237 5.74-2.52 5.74-2.52l9.142-9.172c.502-.502.531-1.226-.008-1.763M2.312 13.906l1.721-3.627l-.018.706l.998-.014l9.524-9.442l1.259 1.236l-9.184 8.965l-.227.752l.486.486l-3.811 1.656z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Percent(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M4.541.068c-1.91 0-3.458 1.77-3.458 3.954s1.548 3.955 3.458 3.955C6.451 7.977 8 6.206 8 4.022S6.451.068 4.541.068m-.042 5.994c-.892 0-1.615-.938-1.615-2.095c0-1.155.724-2.093 1.615-2.093c.894 0 1.616.938 1.616 2.093c0 1.158-.722 2.095-1.616 2.095m9.025 1.989c-1.901 0-3.441 1.77-3.441 3.954s1.54 3.955 3.441 3.955c1.899 0 3.44-1.771 3.44-3.955s-1.54-3.954-3.44-3.954m-.047 6.043c-.885 0-1.602-.934-1.602-2.089c0-1.153.717-2.088 1.602-2.088c.885 0 1.602.935 1.602 2.088c-.001 1.155-.718 2.089-1.602 2.089m-8.743 1.898a.68.68 0 0 1-.341-.092a.666.666 0 0 1-.244-.919L12.646.365a.682.682 0 0 1 .594-.339a.7.7 0 0 1 .344.091c.325.188.434.6.243.919L5.331 15.652a.688.688 0 0 1-.597.34"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Person(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.518 10c-.402.548-.899 1.62-1.479 2.593c-.637 1.074-1.367-1.599-2.166-1.599c-.821 0-1.588 2.624-2.252 1.524c-.572-.947-1.066-1.967-1.456-2.491C1.122 10.027 1 15.914 1 15.914h15.745c0 .001.318-5.914-4.227-5.914m-.623-5.628c0 1.861-1.318 5.422-2.948 5.422C7.319 9.794 6 6.233 6 4.372C6 2.509 7.319 1 8.947 1c1.63-.001 2.948 1.508 2.948 3.372"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonChecked(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M10.926 4.121c0 1.68-1.309 4.891-2.922 4.891c-1.613 0-2.923-3.211-2.923-4.891c0-1.68 1.31-3.04 2.923-3.04s2.922 1.36 2.922 3.04m-1.893 8.675l2.393-2.421l1.279 1.335l1.643-1.662c-.631-.65-1.405-.998-2.669-.998c-.854 1.156-3.689 1.453-3.689 1.453s-2.899-.285-3.753-1.427c-4.093 0-4.217 5.91-4.217 5.91h11.1z"/><path d="m15.094 10.801l-2.381 2.381l-1.416-1.417l-.925.924l2.342 2.341l3.306-3.304z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonDoorMan(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.013 8.966c1.648 0 2.983-2.459 2.983-4.268c0-.239-.024-.471-.07-.695H6.098a3.53 3.53 0 0 0-.07.695c-.001 1.809 1.336 4.268 2.985 4.268M6 1h5.943v1.927H6zm6.042 9.634v5.328h4.93s.316-5.915-4.206-5.915c-.169.233-.434.42-.724.587M5.24 10C1.124 10 1 16 1 16h10v-4.985a10.93 10.93 0 0 1-1.986.434S6.1 11.16 5.24 10m4.776 5.031H8.969v-1.094h1.047zm.015-2H8.953v-1.062h1.078z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonError(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.972 4.071c0 1.664-1.332 4.846-2.976 4.846c-1.645 0-2.979-3.182-2.979-4.846c0-1.665 1.334-3.013 2.979-3.013c1.644.001 2.976 1.348 2.976 3.013m5.042 6.874l-.996-.996l-1.494 1.493l-1.493-1.493l-.996.996l1.493 1.493l-1.493 1.495l.996.996l1.493-1.495l1.494 1.495l.996-.996l-1.494-1.495zm-7.31-.615l-.708.15s-2.904-.283-3.76-1.416c-4.098 0-4.223 5.865-4.223 5.865h9.345l-.63-.65l1.951-1.963z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonMan(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m9.918 13.338l-.943 1.62l-.942-1.62l.488-3.296h.908zM6.534 15C4.833 12.746 5.378 9.224 5.21 9C1.123 8.999 1 15 1 15zm5.048 0H17c-.001 0 0-6.031-3.68-6.031c-.164.22-.035 3.759-1.738 6.031M9.008 8.941C7.39 8.941 6 5.732 6 4.064c0-1.67 1.389-3.006 3.008-3.006a3.016 3.016 0 0 1 2.994 3.006c0 1.668-1.374 4.877-2.994 4.877"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonPeople(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M10.247 11.46c-.646 1.026-1.389-1.53-2.2-1.53c-.834 0-1.612 2.51-2.287 1.459c-.58-.906-1.082-1.881-1.479-2.384c-4.106 0-4.19 5.966-4.19 5.966h15.993s.282-5.992-4.335-5.992c-.408.524-.914 1.551-1.502 2.482z"/><path d="M7.989 9C6.493 9 5.016 4.831 5.016 3.117C5.016 1.4 6.493.01 7.99.01c1.496 0 2.972 1.39 2.972 3.107c0 1.714-1.476 5.881-2.972 5.881z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonPlus(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M10.985 3.044c0 1.671-1.336 4.863-2.983 4.863c-1.648 0-2.986-3.192-2.986-4.863c0-1.669 1.338-3.023 2.986-3.023c1.647 0 2.983 1.354 2.983 3.023m4.973 9.977h-1.982v-1.983h-.972v1.983h-1.983v.971h1.983v1.983h.972v-1.983h1.982z"/><path d="M14.861 9.938c-.663-1.037-1.666-1.908-3.158-1.908c-.854 1.159-3.692 1.456-3.692 1.456S5.108 9.2 4.252 8.056c-4.096 0-4.221 5.923-4.221 5.923h9.906v-1.996h2.031V9.938z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonPrison(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M13.697 16.022H4.38a3.349 3.349 0 0 1-3.341-3.35V3.371A3.348 3.348 0 0 1 4.378.022h9.32a3.348 3.348 0 0 1 3.34 3.349v9.301a3.35 3.35 0 0 1-3.34 3.35zM4.107 1A2.113 2.113 0 0 0 2 3.114v9.77C2 14.05 2.945 15 4.107 15h9.785A2.114 2.114 0 0 0 16 12.884v-9.77A2.113 2.113 0 0 0 13.892 1z"/><path d="M4 1v13.691h1.04V1zm3 0v13.691h1.019V1zm3 0v13.691h1.038V1zm3 0v13.691h.918V1z"/><path d="M12.576 11.048c-.399.587-.895 1.459-1.471 2.498c-.632 1.146-1.359-.257-2.154-.257c-.816 0-1.58 1.351-2.24.176c-.568-1.011-1.061-1.826-1.449-2.389c-2.01 0-3.046 1.3-3.578 2.871c-.534 1.573 15.178 1.566 14.691-.013c-.485-1.58-1.537-2.886-3.799-2.886m-.599-4.894c0 1.661-1.323 4.838-2.955 4.838c-1.634 0-2.956-3.177-2.956-4.838c0-1.663 1.322-3.01 2.956-3.01c1.631.001 2.955 1.347 2.955 3.01"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonPublic(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.855 10.053c-.423.571-1.834.719-1.834.719s-1.441-.142-1.865-.706c-2.035 0-3.098 2.923-3.098 2.923h9.926c0-.001-.838-2.936-3.129-2.936m.088-4.135C9.943 6.977 9.062 9 7.978 9C6.89 9 6.011 6.977 6.011 5.918C6.011 4.859 6.89 4 7.978 4c1.084 0 1.965.859 1.965 1.918m4.161 2.103c-.371.575-1.527.881-1.527.881s-1.355-.301-1.729-.867c0 0-.012.664-.577 1.436c1.8-.232 2.578 1.503 2.578 1.503l3.131.006c-.001-.001.139-2.959-1.876-2.959m-.114-3.488c0 .848-.662 2.465-1.479 2.465c-.82 0-1.481-1.617-1.481-2.465c0-.846.663-1.533 1.481-1.533c.817 0 1.479.688 1.479 1.533M1.918 8.021c.378.571 1.549.875 1.549.875s1.373-.299 1.748-.861c0 0 .014.66.586 1.426c-1.824-.23-2.61 1.492-2.61 1.492l-3.17.005s-.143-2.937 1.897-2.937m.084-3.455c0 .828.664 2.411 1.479 2.411c.819 0 1.48-1.583 1.48-2.411a1.49 1.49 0 0 0-1.48-1.501a1.49 1.49 0 0 0-1.479 1.501"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonRemove(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M10.985 4.044c0 1.671-1.336 4.863-2.983 4.863c-1.648 0-2.986-3.192-2.986-4.863c0-1.669 1.338-3.023 2.986-3.023c1.647 0 2.983 1.354 2.983 3.023m-1.047 8.938h5.823c-.447-1.584-1.593-3.953-4.058-3.953c-.854 1.159-3.692 1.456-3.692 1.456s-2.903-.286-3.759-1.43c-4.096 0-4.221 5.923-4.221 5.923h9.906z"/><path d="M11 14h4.937v.972H11z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonTalk(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.524 12.285c-.86 0-2.015-.449-2.601-1.186C.114 11.099.029 15 .029 15h10.937c.001 0 .22-3.917-2.937-3.917c-.584.745-1.643 1.202-2.505 1.202m2.414-6.021c0 1.251-1.105 3.643-2.469 3.643C4.105 9.907 3 7.515 3 6.264C3 5.015 4.104 4 5.469 4c1.364 0 2.469 1.015 2.469 2.264M12.527.041c-1.91 0-3.461 1.158-3.461 2.59c0 1.306 1.294 2.382 2.971 2.561l-.986 1.812l2.756-1.969c1.277-.381 2.182-1.313 2.182-2.404c-.001-1.432-1.55-2.59-3.462-2.59"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.564 10.004c-.847 1.155-3.659 1.45-3.659 1.45s-2.876-.284-3.723-1.426C1.122 10.028 1 15.933 1 15.933h15.808c0 .001.319-5.929-4.244-5.929m-.746-5.984c0 1.669-1.303 4.857-2.908 4.857C7.303 8.877 6 5.689 6 4.02C6 2.353 7.303 1 8.91 1c1.605.001 2.908 1.353 2.908 3.02"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PersonWoman(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.677 11c-2.843 2.147-.724 3.424-3.653 3.424c-2.932 0-.604-1.191-3.76-3.398c-4.908 1.381-4.226 4.884-4.226 4.884h15.875c-.001 0 .651-3.921-4.236-4.91m-.103-4.002l-.575 1.88L13.883 9s-1.153-.85-1.309-2.002m-7.18.025l.538 1.893L4.045 9s1.169-.826 1.349-1.977m6.535-1.977c0 1.68-1.323 4.895-2.954 4.895c-1.632 0-2.954-3.215-2.954-4.895C6.021 3.363 7.343 2 8.975 2c1.63.001 2.954 1.364 2.954 3.046"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PetCarrier(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.504 3.037h-.535V2.022c0-.564-.446-1.022-.995-1.022H7.032c-.549 0-.994.458-.994 1.022v1.015h-.543c-3.682 0-4.494 11.789-4.494 11.789c0 .58.514 1.054 1.147 1.054h13.704c.634 0 1.148-.474 1.148-1.054c0 0-.883-11.789-4.496-11.789m-5.546 7.98V8.934H11v2.083zM11.04 12v2H6.988v-2zm-.009-6.083v2.104H6.958V5.917zm3.661 2.104h-2.755V5.917h2.125c.251.639.459 1.366.63 2.104m-8.661 0H3.285c.165-.729.367-1.473.615-2.104h2.131zm-.01.913v2.083H2.75c.086-.616.196-1.36.346-2.083zm5.958-.031H14.9c.153.731.271 1.489.359 2.113h-3.28zM6.977 2.185c0-.17.148-.309.33-.309H10.7c.182 0 .33.139.33.309v.853H6.977zm-4.46 11.041s.036-.45.133-1.226h3.371v2H3.365c-.469 0-.848-.349-.848-.774M14.629 14H11.98v-2h3.362c.1.768.141 1.233.141 1.233c-.001.422-.385.767-.854.767"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneFax(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M3.964 13.945c0 .565-.476 1.023-1.063 1.023h-.807c-.586 0-1.062-.458-1.062-1.023V1.054c0-.564.477-1.023 1.062-1.023h.807c.588 0 1.063.459 1.063 1.023zM15.881 0H6.034A1.03 1.03 0 0 0 5 1.023v12.891c0 .565.463 1.023 1.034 1.023h.924v-2.021H15v2.021h.881c.573 0 1.035-.458 1.035-1.023V1.023C16.916.459 16.454 0 15.881 0M8 12.021H6.969v-1.053H8zm0-2H6.969V8.968H8zm0-2H6.969V6.984H8zm2 4H8.969v-1.053H10zm0-2H8.969V8.953H10zm0-2H8.969V6.968H10zm2 4h-1.031v-1.037H12zm0-2h-1.031V8.984H12zm0-2h-1.031V6.953H12zm3.016.954v1.047h-2.047V8.975zm-2.047-.954V6.959h2.047v1.062zm2.047-3.99H6.969V1.969h8.047z"/><path d="m13.969 15.969l-.002-1.947H8.018l.002 1.947z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PhoneNumber(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.924 3.779C15.924-.037 9.787.09 7.637.09C5.485.09.038.136.038 3.953h3.909l.004-.957c0-.438.479-1.027 1.553-1.027h5.013c1.046 0 1.537.589 1.537 1.027c0 .154.008.927.008.957h3.86zM4.987 9.338a.628.628 0 0 1-.638.617h-.697a.628.628 0 0 1-.639-.617v-.676c0-.341.286-.617.639-.617h.697c.352 0 .638.276.638.617zm7.988.007a.637.637 0 0 1-.639.638h-.696a.637.637 0 0 1-.638-.638v-.697c0-.353.285-.638.638-.638h.696c.354 0 .639.285.639.638zm-3.989.004a.637.637 0 0 1-.638.638H7.65a.637.637 0 0 1-.637-.638v-.697c0-.353.285-.638.637-.638h.698c.353 0 .638.285.638.638zm-4.024 2.987a.633.633 0 0 1-.628.638h-.686a.633.633 0 0 1-.629-.638v-.697c0-.354.281-.639.629-.639h.686c.347 0 .628.285.628.639zm7.981.01a.621.621 0 0 1-.617.628h-.676a.621.621 0 0 1-.617-.628v-.686c0-.349.275-.629.617-.629h.676c.342 0 .617.28.617.629zm-3.961 0a.631.631 0 0 1-.637.628h-.697a.632.632 0 0 1-.638-.628v-.686c0-.349.285-.629.638-.629h.697c.353 0 .637.28.637.629zm-3.989 3.001a.628.628 0 0 1-.638.618h-.697a.628.628 0 0 1-.639-.618v-.676c0-.341.286-.617.639-.617h.697c.352 0 .638.276.638.617zm7.982.01a.638.638 0 0 1-.639.639h-.696a.637.637 0 0 1-.638-.639v-.697c0-.353.285-.638.638-.638h.696c.354 0 .639.285.639.638zm-3.993 0a.631.631 0 0 1-.627.639h-.686a.631.631 0 0 1-.627-.639v-.697c0-.353.279-.638.627-.638h.686c.348 0 .627.285.627.638zM.992 5.973h2.047s.707.228.887-.973H.067c.167 1.113.925.973.925.973M12.115 5c.209 1.053.963.973.963.973h1.891s.758.011.943-.973z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Piano(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 3v9h15.986V3zm2 1h2v1H2zm13.083 7.059H.992V7h1.091v2.938h.909V7h1.091v2.938h.909V7h1.091v2.938h.909V7h2.091v2.938h.909V7h1.091v2.938h.909V7h1.091v2.938h.909V7h1.091zM15 5H8V4h7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pick(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M9.168 5.921c.009.006 1.877 1.889 1.885 1.896l1.314-1.32c3.855 4.404 2.466 7.059 2.615 6.908c1.901-1.899.89-5.99-2.256-9.136C9.58 1.123 5.488.112 3.591 2.011c-.148.146 2.5-1.271 6.908 2.593zm4.972-4.317l.88.877l-.946.947l-.879-.877z"/><path d="M8.873 7.166L.289 15.752a.667.667 0 0 0-.009.943a.668.668 0 0 0 .944-.01l8.584-8.584z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Picture(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" transform="translate(1 1)"><path d="m13.842 4l-3.93-2.85a1.495 1.495 0 0 0-2.928.093L2.918 4H.021v9.96h15.938V4zM8.458 3c.556 0 1.034-.305 1.294-.753L11.895 4H5.047l2.166-1.664c.27.4.727.664 1.245.664m3.503 7.085L8.72 11.758l-3.727-4.72L1 13V5h14v8z"/><circle cx="12.963" cy="6.963" r=".963"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PictureCopy(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M14.969.031H5v2.938h1V1h8.027v7.001H11v2.951h3.969z"/><path d="M0 4v11h10V4zm8.967 8h-8V5h8z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PictureTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 0v16h16V0zm14 11H3V2h12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PiggyBank(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.041 6.33c0-.528-1.832-.95-1.93-.939c-1.005-1.909-3.206-3.228-5.762-3.358c-2.038-.104-3.416.358-4.387 1a3.55 3.55 0 0 0-1.084.001c-.47.081-.809.258-1.18.581c-.213.185-.59.58-.231.787c.259.15.633-.108.909-.124c.079-.006.152.009.225.022c-.697.907-.963 1.739-1.264 1.739C.054 6.039.04 6.762.04 7.345v1.333c0 .381.481 1.286.67 1.5c1.041 1.181 1.635 2.057 3.36 2.633c.03.663.463 1.189.997 1.189c.418 0 1.179-1.039 2.225-1.039c2.5 0 3.369 1.039 3.797 1.039c1.119 0 .996-1.354.992-1.402c1.957-.939 2.614-2.757 2.614-4.879c0-.15-.013-.297-.025-.444c.777-.076 1.371-.466 1.371-.945M4 7a1 1 0 1 1 0-2a1 1 0 0 1 0 2m7.274-2.274c-.432-.386-1.113-.576-1.782-.763c-.67-.186-1.226-.341-1.801-.234c-.744.139-.986-.216-.392-.447c.685-.266 1.605-.4 2.509-.148c.86.238 1.538.791 1.978 1.33c.421.521.04.76-.512.262"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pill(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m15.897 1.731l-.656-.657C13.887-.281 11.745-.341 10.46.944L1.957 9.446c-1.284 1.285-1.224 3.425.133 4.782l.654.654c1.357 1.357 3.498 1.418 4.783.134l8.503-8.505c1.284-1.282 1.224-3.423-.133-4.78m-4.811 8.433L6.841 5.917l4.208-4.208c.945-.944 2.532-.898 3.535.107l.604.601c.49.493.771 1.135.787 1.809a2.331 2.331 0 0 1-.68 1.729z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinLocation(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.469.021c-3.016 0-5.46 2.296-5.46 5.13c0 .732.166 1.428.458 2.057l5.002 8.668l5-8.668a4.84 4.84 0 0 0 .459-2.057c0-2.835-2.444-5.13-5.459-5.13m.023 9.021c-1.957 0-3.542-1.596-3.542-3.567c0-1.969 1.585-3.565 3.542-3.565c1.954 0 3.539 1.597 3.539 3.565c0 1.971-1.585 3.567-3.539 3.567"/><path d="M10.979 5.504A2.485 2.485 0 0 1 8.5 7.996a2.485 2.485 0 0 1-2.477-2.492A2.485 2.485 0 0 1 8.5 3.014a2.485 2.485 0 0 1 2.479 2.49"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinLocationAdd(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.475.031c-3.007 0-5.443 2.512-5.443 5.609c0 1.584 5.443 10.275 5.443 10.275s5.441-8.609 5.441-10.275c0-3.097-2.437-5.609-5.441-5.609m2.556 5.985h-2v2H8.969v-2h-2V4.954h2v-2h1.062v2h2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinLocationDelete(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.475.031c-3.007 0-5.443 2.512-5.443 5.609c0 1.584 5.443 10.275 5.443 10.275s5.441-8.609 5.441-10.275c0-3.097-2.437-5.609-5.441-5.609m3.474 7.867L9.914 8.932L7.5 6.518L5.086 8.932L4.051 7.898l2.414-2.414L4.051 3.07l1.035-1.035L7.5 4.449l2.414-2.414l1.035 1.035l-2.415 2.414z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinLocationLove(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.48 0C5.499 0 3.084 2.516 3.084 5.621c0 1.586 5.396 10.295 5.396 10.295s5.395-8.625 5.395-10.295C13.875 2.516 11.46 0 8.48 0m.029 8.25S5.861 7.015 5.861 4.612c0-.831.602-1.505 1.341-1.505c.637 0 1.166.498 1.305 1.164c.133-.665.661-1.164 1.295-1.164c.732 0 1.326.667 1.326 1.49c-.001 2.448-2.619 3.653-2.619 3.653"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinLocationMap(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="m0 11.438l3.938 1.468V1.062L0-.015zm11.506-3.433c-1.361 0-2.463 1.204-2.463 2.69c0 .384.076.747.207 1.079l2.264 4.152l2.245-4.152c.134-.332.208-.695.208-1.079c0-1.486-1.103-2.69-2.461-2.69m.016 4.041a1.5 1.5 0 1 1-.002-3a1.5 1.5 0 0 1 .002 3"/><path d="M16 .969L11.031.041v6.854c.156-.02.312-.046.475-.046c1.977 0 3.578 1.541 3.578 3.443c0 .492-.108.956-.301 1.382l-.385.629L16 13zm-8.074 9.323c0-1.37.838-2.544 2.043-3.098V.166l-4.938.896v11.844l3.681-.449l-.485-.783a3.344 3.344 0 0 1-.301-1.382"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinLocationOne(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M8 6h.956v7.944H8z"/><path d="M8.511.016C6.047.016 4.047 2.024 4.047 4.5c0 2.478 2 4.484 4.464 4.484c2.465 0 4.464-2.007 4.464-4.484S10.977.016 8.511.016m.159 1.626a3.008 3.008 0 0 0-3.018 2.997c0 .662-.648.346-.648-.645A3.008 3.008 0 0 1 8.022.997c.996 0 1.314.645.648.645m1.439 9.514v.926c2.477.248 3.729 1.062 3.729 1.47c0 .509-1.887 1.501-5.344 1.501c-3.459 0-5.346-.992-5.346-1.501c0-.438 1.342-1.181 3.758-1.421v-.927c-2.379.211-4.872.938-4.872 2.36c0 1.598 3.249 2.433 6.46 2.433c3.209 0 6.459-.835 6.459-2.433c0-1.462-2.56-2.203-4.844-2.408"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinLocationRemove(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.475.031c-3.007 0-5.443 2.512-5.443 5.609c0 1.584 5.443 10.275 5.443 10.275s5.441-8.609 5.441-10.275c0-3.097-2.437-5.609-5.441-5.609m2.556 5.985H5.969V4.954h5.062z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinLocationTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M8 .031c-2.717 0-4.92 2.119-4.92 4.733c0 .675.15 1.317.414 1.898l4.59 7.321l4.422-7.321a4.563 4.563 0 0 0 .412-1.898C12.918 2.15 10.717.031 8 .031m0 8.09a3.09 3.09 0 0 1-3.085-3.098A3.091 3.091 0 0 1 8 1.926a3.093 3.093 0 0 1 3.086 3.097A3.093 3.093 0 0 1 8 8.121m1.977-3.138a1.981 1.981 0 0 1-1.978 1.985a1.982 1.982 0 0 1-1.977-1.985A1.98 1.98 0 0 1 7.999 3c1.092 0 1.978.889 1.978 1.983"/><path d="M5.299 11.823c-1.717.364-2.43.842-2.43 1.379c0 .769 1.831 1.829 5.116 1.829c3.285 0 5.116-1.06 5.116-1.829c0-.535-.708-.999-2.401-1.379v-.82c1.865.366 3.254 1.101 3.254 2.199c0 1.584-3.076 2.77-5.969 2.77c-2.893 0-5.969-1.186-5.969-2.77c0-1.1 1.398-1.849 3.273-2.202z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinMap(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M11.953 8.979H6.025V7.668L7 6.991V4.926l-.975-.804V2.021h5.928v2.228l-.974.677v2.065l.975.634z"/><path d="M9.986 7.993H8.038v2.614l.989 5.372l.959-5.372zM6 0h5.959v.942H6z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PingPongRacket(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.297 10.802c-.451-.452-.447-.899-.372-1.144c.2-.67.474-1.415.632-2.234L8.53 12.451c.793-.162 1.517-.428 2.168-.621c.25-.074.723-.053 1.197.421c.473.475 2.906 3.642 2.906 3.642l2.164-2.163c0-.001-3.216-2.476-3.668-2.928m.021-4.636c-.01-1.285-.493-2.733-2.094-4.336a5.938 5.938 0 0 0-8.396 8.396c1.484 1.483 2.832 1.993 4.047 2.062zm-6.847 2.9a1.537 1.537 0 0 1-1.526-1.549c0-.854.684-1.547 1.526-1.547c.844 0 1.528.693 1.528 1.547c0 .857-.685 1.549-1.528 1.549"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pipe(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.274 4.644C9.75 6.47 8.318 7.874 6.03 8.494l-.729-2.08c-.43-1.237-5.646.584-5.213 1.821l1.133 3.242c.432 1.236 2.059 1.814 3.543 1.376c1.175-.347 4.087-1.353 7.528-6.349c1.125-1.635 2.566-2.66 3.698-3.006l-.515-1.473s-2.278.316-4.2 2.62z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pizza(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.863 2.581a7.436 7.436 0 0 0-1.82 4.363h.971a6.446 6.446 0 0 1 1.551-3.662zm1.403 1.401a5.466 5.466 0 0 0-1.26 2.962l4.221-.001zM6.969.042a7.435 7.435 0 0 0-4.363 1.8l.703.703a6.477 6.477 0 0 1 3.66-1.502zm0 1.989a5.507 5.507 0 0 0-2.958 1.216l2.958 2.958zM2.002 8.032a5.39 5.39 0 0 0 1.227 2.971l2.998-2.97zm-1.961 0a7.32 7.32 0 0 0 1.8 4.346l.687-.68a6.354 6.354 0 0 1-1.519-3.666zm3.926 3.7a5.499 5.499 0 0 0 3.002 1.252V8.731zM2.58 13.119a7.426 7.426 0 0 0 4.389 1.823v-.97a6.485 6.485 0 0 1-3.701-1.54zm10.563-.74a7.333 7.333 0 0 0 1.801-4.346h-.952a6.36 6.36 0 0 1-1.526 3.674zM8.756 8.033l3.009 2.98a5.39 5.39 0 0 0 1.234-2.98zm4.238-1.09a5.466 5.466 0 0 0-1.266-2.97l-2.971 2.97zm1.947 0a7.434 7.434 0 0 0-1.82-4.362l-.693.693a6.434 6.434 0 0 1 1.559 3.669zm-6.925 7.999a7.436 7.436 0 0 0 4.387-1.821l-.68-.68a6.493 6.493 0 0 1-3.707 1.532zm0-1.956a5.512 5.512 0 0 0 3.008-1.244l-3.008-3.01z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plugin(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.016 7.038h-5v2.931h5v.829l5.922 2.175v-8.93L5.016 6.219zm10.937-.054v-.968h-2.971V4.043h-.962v8.93h.962v-1.989h2.971v-.968h-2.971V6.984z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plus(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16 6h-4V2a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1v4H2a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h4v4a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-4h4a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Podium(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.432 14.035h-.479V5.72a.69.69 0 0 0-.688-.692L11 5.254V3.671l.068.027a.651.651 0 0 0 .854-.351a.646.646 0 0 0-.354-.847L8.984 1.354A1.472 1.472 0 0 0 7.515.033c-.819 0-1.484.658-1.484 1.469s.665 1.469 1.484 1.469c.392 0 .745-.153 1.011-.398l1.48.691v2.175l-5.272.976a.69.69 0 0 0-.689.693v6.93h-.539c-.279 0-.506.429-.506.96c0 .532.227.963.506.963h9.926c.279 0 .506-.431.506-.963c0-.534-.227-.963-.506-.963"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PokerFour(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.372.058H5.627a1.62 1.62 0 0 0-1.613 1.624v12.689a1.62 1.62 0 0 0 1.613 1.624h6.745a1.62 1.62 0 0 0 1.614-1.624V1.682A1.62 1.62 0 0 0 12.372.058M9.023 11S6.034 9.544 6.034 6.71c0-.979.678-1.776 1.513-1.776c.719 0 1.318.587 1.475 1.374c.15-.787.744-1.374 1.46-1.374c.828 0 1.499.787 1.499 1.76C11.98 9.58 9.023 11 9.023 11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PokerOne(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.494 1H5.412C4.082 1 3 1.098 3 2.445v13.1c0 1.348 1.082 1.445 2.412 1.445h6.082c1.33 0 2.412-.098 2.412-1.445v-13.1C13.906 1.098 12.824 1 11.494 1m-2.489 9.564v2.457H7.958v-2.424c-.489.268-1.19.336-1.631.336c-.728 0-1.317-.785-1.317-1.754c0-.967.59-1.75 1.317-1.75c.244 0 .573.092.902.246c-.133-.393-.216-.787-.216-1.078c0-.859.667-1.557 1.486-1.557c.817 0 1.482.697 1.482 1.557c0 .297-.084.701-.222 1.102c.34-.168.686-.27.938-.27c.713 0 1.291.783 1.291 1.75c0 .969-.578 1.754-1.291 1.754c-.447.001-1.197-.074-1.692-.369"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PokerThree(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.494 1H5.412C4.082 1 3 1.097 3 2.444v13.101c0 1.348 1.082 1.445 2.412 1.445h6.082c1.33 0 2.412-.098 2.412-1.445V2.444C13.906 1.097 12.824 1 11.494 1m-2.092 9.998h-.34v2.064H7.906v-2.064h-.258c-1.041.125-1.794.299-1.794-1.015c0-1.647 2.636-4.952 2.636-4.952s2.635 3.305 2.635 4.952c0 1.314-.682 1.14-1.723 1.015"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PokerTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.494 1.01H4.412C3.082 1.01 2 1.107 2 2.454v13.101C2 16.903 3.082 17 4.412 17h6.082c1.33 0 2.412-.098 2.412-1.445V2.454c0-1.348-1.082-1.444-2.412-1.444M7.489 13.135L4.748 9.041l2.824-4.115l2.742 4.093z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Print(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M6 12v3.98h5.993V12zM5 1h7.948v2.96H5z"/><path d="M1.041 5.016v9h3.91V11.01H13v3.006h4.041v-9zm2.975 2.013H2.969V5.953h1.047zm2-.06H4.969v-1h1.047z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Puzzle(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m14.992 6.567l-.01-2.817c0-.911-.799-1.724-1.709-1.724h-3.185C10.092.698 9.006.075 7.668.075c-1.337 0-2.419.623-2.419 1.917c0 .012-2.536.034-2.536.034c-.908 0-1.681.813-1.681 1.724l.035 2.711c1.322.016 2.39 1.056 2.39 2.341c0 1.284-1.067 2.323-2.39 2.34v3.208c0 .911.736 1.65 1.645 1.65h10.561c.909 0 1.745-.739 1.745-1.65l-.011-3.116c1.162-.139 1.969-1.125 1.969-2.333c0-1.215-.813-2.202-1.984-2.334"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuoteClose(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 14.969c-.553 0-1-.435-1-.969s.447-.969 1-.969c2.757 0 4-1.201 4-3.907V7.978H2.559a1.51 1.51 0 0 1-1.506-1.511V2.511A1.51 1.51 0 0 1 2.559.999h3.935C7.324.999 8 1.677 8 2.511v6.612c0 3.775-2.141 5.846-6 5.846m9 0c-.553 0-1-.435-1-.969s.447-.969 1-.969c2.757 0 4-1.201 4-3.907V7.947h-3.467a1.51 1.51 0 0 1-1.512-1.506V2.505c0-.83.678-1.506 1.512-1.506h3.955A1.51 1.51 0 0 1 17 2.505v6.618c0 3.775-2.141 5.846-6 5.846"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuoteOpen(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16 14.969c.552 0 1-.435 1-.969s-.448-.969-1-.969c-2.757 0-4-1.201-4-3.907V7.906h3.441c.83 0 1.506-.605 1.506-1.438V2.512A1.51 1.51 0 0 0 15.441 1h-3.935A1.51 1.51 0 0 0 10 2.512v6.612c0 3.775 2.141 5.845 6 5.845m-9 0c.552 0 1-.435 1-.969s-.448-.969-1-.969c-2.757 0-4-1.201-4-3.907V7.947h3.467a1.51 1.51 0 0 0 1.512-1.506V2.505A1.51 1.51 0 0 0 6.467.999H2.512A1.51 1.51 0 0 0 1 2.505v6.618c0 3.776 2.141 5.846 6 5.846"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Radio(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3 0h.938v4.93H3zm9.018 6.01v10h2.962v-10zm2.003 4.011h-1.042V9h1.042zm0-2h-1.042V7h1.042zm-13.01 7.995h10v-10h-10zm4.497-.954a3.548 3.548 0 0 1-2.562-1.094h5.125a3.548 3.548 0 0 1-2.563 1.094m3.521-3.093a3.54 3.54 0 0 1-.307 1.047h-6.43a3.57 3.57 0 0 1-.307-1.047zm-7.045-.953c.049-.37.147-.722.299-1.047H8.73c.153.325.252.677.301 1.047zm3.524-3.11c1.016 0 1.926.429 2.576 1.109H2.932a3.551 3.551 0 0 1 2.576-1.109"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Radioactive(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m10.295 6.346l2.734-5.187A7.542 7.542 0 0 0 9.035.004a7.55 7.55 0 0 0-3.994 1.155l2.727 5.17a2.607 2.607 0 0 1 1.246-.326c.468 0 .9.13 1.281.343m-.619 1.206a1.305 1.305 0 0 0-.662-.184c-.229 0-.439.063-.627.165c-.406.225-.686.65-.686 1.148c0 .028.008.055.01.083c.03.48.316.887.726 1.09c.175.086.37.14.577.14c.223 0 .428-.061.61-.156c.394-.209.665-.605.694-1.073c.003-.028.01-.055.01-.083c0-.485-.264-.903-.652-1.13m1.928 1.479a2.61 2.61 0 0 1-1.342 1.931l2.758 5.011c2.309-1.429 3.873-3.918 3.973-6.941zm-6.576 6.942l2.77-4.995a2.61 2.61 0 0 1-1.376-1.946H1.027c.1 3.023 1.676 5.512 4.001 6.941"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Record(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.98 5.004c-2.193 0-3.976 1.793-3.976 3.996c0 .732.211 1.409.555 2H8.354a3.942 3.942 0 0 0 .572-2.026c0-2.19-1.776-3.973-3.958-3.973c-2.182 0-3.957 1.782-3.957 3.973s1.775 3.973 3.957 3.973c.011 0 .031.053.031.053h8a4.001 4.001 0 0 0 3.958-4c0-2.203-1.784-3.996-3.977-3.996M5 11c-1.104 0-2-.897-2-2s.896-2 2-2c1.103 0 2 .897 2 2s-.897 2-2 2m8 0c-1.103 0-2-.897-2-1.999C11 7.898 11.898 7 13 7a2 2 0 0 1 0 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reduce(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.329 9.083a.69.69 0 0 0-.688.69l1.588 1.594L.25 14.269a.979.979 0 0 0 0 1.39a.984.984 0 0 0 1.391 0l2.976-2.9l1.579 1.585c.38 0 .687-.31.687-.69V9.773c0-.38-.307-.69-.687-.69zm2.244-5.826L1.786.209a.983.983 0 0 0-1.392 0a.944.944 0 0 0 0 1.357L3.18 4.645L1.62 6.199c0 .38.309.687.69.687h3.881c.381 0 .69-.307.69-.687V2.333a.686.686 0 0 0-.689-.688zm9.769 6.523a.69.69 0 0 0-.691-.687h-3.88a.688.688 0 0 0-.69.687v3.866c0 .381.31.688.69.688l1.595-1.587l2.968 2.976a.978.978 0 0 0 1.39 0a.983.983 0 0 0 0-1.389l-2.966-2.975zm-2.981-6.545L9.767 1.634a.69.69 0 0 0-.688.69v3.881c0 .381.309.69.688.69h3.867a.69.69 0 0 0 .687-.69L12.75 4.627l2.961-3.073a.982.982 0 0 0 0-1.39c-.383-.384-1.005-.116-1.39.268z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReelAudio(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M1.064 7.211A2.95 2.95 0 0 0 3 7.933a2.956 2.956 0 0 0 2.961-2.95c0-.506-.127-.981-.353-1.396A2.965 2.965 0 0 0 3 2.032A2.955 2.955 0 0 0 .039 4.983c0 .89.399 1.687 1.025 2.228m3.416-.367l-1.3-1.286a.582.582 0 0 0 .38-.46l1.332 1.318a2.316 2.316 0 0 1-.412.428M3.242 2.678c.197.021.389.072.574.141l-.39 1.791a.573.573 0 0 0-.434-.202c-.048 0-.091.017-.135.027zM.734 5.477l1.723-.596c-.006.035-.02.066-.02.102c0 .187.092.346.226.453L.93 6.035a2.218 2.218 0 0 1-.139-.309c-.028-.082-.039-.167-.057-.249m12.243 2.507c.721 0 1.382-.258 1.898-.684a2.952 2.952 0 0 0 1.07-2.276a2.965 2.965 0 0 0-2.969-2.961a2.969 2.969 0 0 0-2.97 2.961a2.967 2.967 0 0 0 2.971 2.96m1.529-1.093l-1.303-1.303a.592.592 0 0 0 .381-.466l1.334 1.335a2.35 2.35 0 0 1-.412.434m-1.244-4.219c.199.019.391.072.578.142l-.391 1.813a.579.579 0 0 0-.434-.205c-.051 0-.092.017-.136.028zm-2.465 3.109c-.025-.083-.037-.168-.055-.254l1.705-.606c-.006.036-.021.067-.021.104a.6.6 0 0 0 .224.462l-1.715.611c-.05-.102-.102-.205-.138-.317"/><path d="M10.028 8.53a4.486 4.486 0 0 1-1.56-3.403c0-.592.116-1.61.324-2.127H7.109c.208.518.324 1.535.324 2.127a4.508 4.508 0 0 1-4.511 4.504c-.86 0-2.24-.242-2.922-.661v3.637C0 13.442 1.177 14 1.917 14h12.067c.74 0 2.016-.558 2.016-1.393V8.97c-.685.417-2.162.662-3.024.662a4.485 4.485 0 0 1-2.948-1.102M3 11.957H2v-1h1zm1.96.001h-1v-1h1zM9 10H7V9h2zm3 1.969h-1v-1h1zm.953-1h1v1h-1z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ReelFilm(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M7.368 13.095c1.193-.188 2.796.098 3.817 1.004c1.021.906 2.453.801 3.815 1.029V16c-1.144-.193-3.37-.416-4.15-1.108c-.778-.69-2.062-1.09-3.267-.899c-1.096.173-1.319-.724-.215-.898"/><path d="M7.5.166C3.403.166.08 3.469.08 7.542s3.323 7.375 7.42 7.375c4.098 0 7.42-3.302 7.42-7.375S11.598.166 7.5.166M5 3a2 2 0 1 1-.001 4.001A2 2 0 0 1 5 3m0 9a2 2 0 1 1 .001-4.001A2 2 0 0 1 5 12m1.709-4.486c0-.455.38-.823.846-.823c.468 0 .848.368.848.823c0 .454-.38.823-.848.823a.836.836 0 0 1-.846-.823M10 12a2 2 0 1 1 .001-4.001A2 2 0 0 1 10 12m0-5a2 2 0 1 1 .001-4.001A2 2 0 0 1 10 7"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Remove(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 12a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Repeat(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.986 4.016H6.939L4.935 2.193a.645.645 0 0 0-.911 0v1.823H1.998a2 2 0 0 0-1.998 2v3.969a2 2 0 0 0 1.998 2h7.238l1.826 1.828a.646.646 0 0 0 .912 0v-1.828h2.012a2 2 0 0 0 1.998-2V6.016a2 2 0 0 0-1.998-2m-2.924 4.187L9.15 10.031H1.984V5.969h2.04v1.746a.631.631 0 0 0 .911 0l1.879-1.746h7.201l.002 4.062h-2.043V8.203a.632.632 0 0 0-.912 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ResizeInFrame(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M9.994 12.967c0 1.134-.92 2.054-2.055 2.054H4.03a2.055 2.055 0 0 1-2.054-2.054v-3.91c0-1.133.92-2.053 2.054-2.053h3.909c1.135 0 2.055.92 2.055 2.053zm1.784-8.863l-1.171-2.056a.55.55 0 0 0-.551.548L10.038 6.4a.55.55 0 0 0 .552.549l3.765.017a.55.55 0 0 0 .553-.548l-2.014-1.206l3.816-3.794a.768.768 0 0 0 .231-.551a.77.77 0 0 0-.231-.553a.794.794 0 0 0-1.117 0z"/><path d="M16.04 7.021v6.584c0 .795-.642 1.441-1.435 1.441H3.376a1.44 1.44 0 0 1-1.435-1.441V2.396c0-.794.644-1.44 1.435-1.44H9.98v-.93H3.057a2.054 2.054 0 0 0-2.049 2.053v11.832c0 1.131.92 2.053 2.049 2.053h11.881a2.055 2.055 0 0 0 2.051-2.053v-6.89z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ResizeOutFrame(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.994 12.967c0 1.134-.92 2.054-2.055 2.054H3.03a2.055 2.055 0 0 1-2.054-2.054v-3.91c0-1.133.92-2.053 2.054-2.053h3.909c1.135 0 2.055.92 2.055 2.053zm5.207-10.023L15.372 5a.55.55 0 0 0 .551-.548l.019-3.804a.55.55 0 0 0-.552-.549L11.625.082a.549.549 0 0 0-.553.548l2.014 1.206L9.27 5.63a.768.768 0 0 0-.231.551c0 .202.076.4.231.553a.794.794 0 0 0 1.117 0z"/><path d="M15.04 7.021v6.584a1.44 1.44 0 0 1-1.436 1.441H2.375a1.44 1.44 0 0 1-1.436-1.441V2.396c0-.794.645-1.44 1.436-1.44h6.604v-.93H2.056A2.054 2.054 0 0 0 .007 2.079v11.832c0 1.131.92 2.053 2.049 2.053h11.881a2.055 2.055 0 0 0 2.051-2.053v-6.89z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Retroeexcavadora(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" transform="translate(1 1)"><circle cx="12.473" cy="11.473" r="1.473"/><path d="M3.502 6.035C1.566 6.035 0 7.588 0 9.503s1.566 3.467 3.502 3.467c1.932 0 3.5-1.552 3.5-3.467s-1.568-3.468-3.5-3.468m0 5.127a1.668 1.668 0 0 1-1.677-1.659c0-.917.751-1.66 1.677-1.66c.924 0 1.675.743 1.675 1.66c0 .917-.751 1.659-1.675 1.659"/><path d="M15.332 6.68a.893.893 0 0 0-.891-.896l-.504-.07v-.732h.024V2.121h-.96V4h-.971v1.451l-2.436-.336C8.992 3.447 7.565.094 7.006.094l-5.45.669C.563.763.03 6 .03 6c.672-.389 2.542-1.062 3.375-1.062c2.526 0 4.641 2.16 4.641 4.688c0 .479-.079.902-.216 1.338h2.101c.146-1.245 1.107-2.129 2.393-2.129c1.283 0 2.33.89 2.477 2.135h.309c.492 0 .891-.321.891-.816zM2.938 4.031V1.937l3.125-.5l1 2.594z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ribbon(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m9.025 0l1.39 1.025h1.717l.53 1.657l1.39 1.025l-.531 1.659l.531 1.658l-1.39 1.024l-.53 1.658h-1.717l-1.39 1.025l-1.39-1.025H5.92l-.529-1.658l-1.389-1.024l.529-1.658l-.529-1.659l1.389-1.025l.529-1.657h1.715zM5.042 15.529l-1.04-2.682l-1.984 1.599L4.006 9l.675 1.577l2.317.014zm7.914-.021l1.075-2.661l1.948 1.578L14.027 9l-.675 1.577l-2.317.014z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightJustify(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16 1.938c0 .518-.42.938-.938.938H.98A.938.938 0 0 1 .98 1h14.082c.518 0 .938.42.938.938m0 3c0 .518-.42.938-.938.938H.98A.938.938 0 0 1 .98 4h14.082c.518 0 .938.42.938.938m0 3c0 .518-.42.938-.938.938H.98A.938.938 0 0 1 .98 7h14.082c.518 0 .938.42.938.938m0 3c0 .518-.42.938-.938.938H.98A.938.938 0 0 1 .98 10h14.082c.518 0 .938.42.938.938m0 3c0 .518-.42.938-.938.938H5.98a.938.938 0 0 1 0-1.876h9.082c.518 0 .938.42.938.938"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightwardsArrowToBar(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.002 1c0-.553.442-1 .989-1h1.979c.547 0 .989.447.989 1v14c0 .553-.442 1-.989 1h-1.979a.994.994 0 0 1-.989-1zM3.113 15.495c-.582.581-2.103.9-2.103-1.001V1.506c0-1.839 1.521-1.582 2.103-1l6.444 6.442a1.49 1.49 0 0 1 0 2.104z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ring(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.163 4.847c.492-.239.833-.738.833-1.322c0-.814-.66-1.474-1.474-1.474c-.771 0-1.4.598-1.462 1.356a6.484 6.484 0 0 0-1.681-.455L12.941.009h-6.9l2.61 2.943a6.443 6.443 0 0 0-1.692.46a1.462 1.462 0 0 0-1.46-1.33c-.813 0-1.474.651-1.474 1.458c0 .579.344 1.075.838 1.31A6.563 6.563 0 0 0 2.98 9.456c0 3.622 2.932 6.568 6.535 6.568s6.533-2.946 6.533-6.568a6.562 6.562 0 0 0-1.885-4.609m-4.647 9.178c-2.501 0-4.535-2.05-4.535-4.568s2.034-4.568 4.535-4.568c2.5 0 4.533 2.05 4.533 4.568s-2.033 4.568-4.533 4.568"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Road(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.073 0h-5.04v1.042H7.957V0H3.083l-2 16h14.922zM9 15H8v-3h1zm0-4.958H8V7h1zM7.958 4.959v-2h1v2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rocket(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 13c-.064.011-.103.027-.161.039c.701-2.108 1.133-5.049 1.133-7.175c0-4.315-2.932-5.837-3.965-5.837S5.04 1.548 5.04 5.864c0 2.125.432 5.063 1.132 7.171c-.057-.011-.11-.024-.173-.034c-.552 1.521-1.988 2.039-3.004 2.988l3.527.008c.276-.337.441-.72.476-1.127c.016.023.03.052.046.074h3.923c.028-.04.054-.089.082-.132c.031.412.192.802.467 1.143l3.471-.007C13.991 14.99 12.538 14.529 12 13m-3-3a1 1 0 1 1 0-2a1 1 0 0 1 0 2m0-4a2 2 0 1 1 .001-4.001A2 2 0 0 1 9 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Roller(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M7.5 15.938a.501.501 0 0 1-.5-.5v-3.906c0-.275.225-.5.5-.5H8V8.389c0-.248.179-.456.425-.494L16 5.93V3.032h-.5a.501.501 0 0 1-.5-.5c0-.275.225-.5.5-.5h1c.275 0 .5.225.5.5v3.857a.496.496 0 0 1-.423.494L9 8.849v2.183h.5c.275 0 .5.225.5.5v3.906c0 .275-.225.5-.5.5z"/><path d="M14 3V2a2 2 0 0 0-2-2H3a2 2 0 0 0-2 2v1a2 2 0 0 0 2 2h9a2 2 0 0 0 2-2"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rss(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M1 3c7.355 0 13 5.593 13 12.968h3C17 7.198 9.747 0 1 0z"/><path d="M1.052 8.99c4.008 0 6.957 2.9 6.957 6.919h2.95c0-5.346-4.578-9.847-9.907-9.847zm.01 6.91h3.91c0-2.491-1.43-3.797-3.91-3.82c-.014 0 0 3.82 0 3.82"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RugbyBall(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M4.516 8.938c.213 0 .414.043.599.119c.286-.365.591-.736.913-1.115a1.487 1.487 0 0 1-.074-.449a1.558 1.558 0 0 1 1.882-1.522c.362-.369.743-.74 1.143-1.114A1.575 1.575 0 0 1 10.5 2.922c.198 0 .387.041.561.109a44.852 44.852 0 0 1 3.967-2.9c-1.836-.252-6.53-.463-10.371 3.442C.647 7.655.659 13.079 1.419 14.917a32.274 32.274 0 0 1 1.932-3.399a1.536 1.536 0 0 1-.396-1.025a1.555 1.555 0 0 1 1.561-1.555"/><path d="M16.572.5c-1.559.761-3.159 1.917-4.706 3.254a1.575 1.575 0 0 1-1.366 2.34c-.316 0-.611-.098-.857-.262a48.64 48.64 0 0 0-.81.83c.153.241.245.524.245.83c0 .858-.699 1.555-1.562 1.555a1.54 1.54 0 0 1-.689-.166c-.297.348-.586.69-.863 1.029c.072.18.115.376.115.582c0 .858-.699 1.555-1.562 1.555c-.072 0-.141-.012-.211-.021c-1.188 1.597-2.01 2.888-2.261 3.475c1.771.788 7.329.867 11.362-3.215C17.49 8.153 17.264 2.07 16.572.5"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ruler(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 0v16h8V0zm1 2h2v1H6zm0 6h2v1H6zm2 7H6v-1h2zm1-3H6v-1h3zm0-6H6V5h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SafeBox(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" fill-rule="evenodd"><path d="M13.001.979H3.018c-.607 0-1.038.354-1.038 1.038v11.021c0 .684.5 1.004 1.06 1.004h9.94c.562 0 1.06-.375 1.06-1.047V2.015c0-.568-.54-1.036-1.039-1.036M13 13H3V2h10z"/><path fill="currentColor" d="M12.901.016H3.16C2.002.016 1 .907 1 2.005V13c0 1.099.871 1.989 2.028 1.989h.012V16H5v-1.011h6V16h2v-1.012c1.152-.005 1.998-.894 1.998-1.988V2.005c0-1.099-.938-1.989-2.097-1.989m1.138 12.979c0 .672-.497 1.047-1.06 1.047h-9.94c-.56 0-1.06-.32-1.06-1.004V2.017c0-.685.431-1.038 1.038-1.038H13c.499 0 1.038.467 1.038 1.035z"/><path fill="currentColor" d="M3 13h10V2H3zm5-7a2 2 0 0 1 2 2h1v1H9.723c-.347.595-.984 1-1.723 1a2 2 0 0 1 0-4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SafePin(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.756 12.992c-.066-.122-1.613-5.493-4.037-9.289c-.722-1.129-1.5-2.091-2.308-2.692l-.364.442c3.161 3.087 4.869 8.3 5.594 10.609c-.603-.048-1.614.069-2.063.518c-.442.44-.613 1.133-.57 1.712c-.026-.02-.056-.04-.077-.058c-1.338-1.046-5.919-5.068-8.672-7.611c.133-.352.266-.649.404-.911c.064-.121-.04-.382-.148-.425a1.365 1.365 0 0 1-.856-1.273c0-.758.601-1.373 1.345-1.373c.604 0 1.113.407 1.281.968c.035.116.33.199.479.112c.398-.234.861-.49 1.408-.811L6.076.817c-1.083-1.081-2.92-1-4.102.179C.792 2.175.712 4.009 1.796 5.091c0 0 1.988 1.914 2.244 2.151c2.936 2.727 7.797 6.938 9.768 8.402l.068.051c.435.275.959.363 1.457.263a2.08 2.08 0 0 0 1.629-1.627c.092-.456.018-.93-.206-1.339m-1.76 2.305a1.283 1.283 0 0 1-1.283-1.281a1.283 1.283 0 0 1 2.567 0c0 .707-.575 1.281-1.284 1.281"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Satellite(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.495 5.893c.337.338 1.689-1.016 1.352-1.352l-2.983-2.984c-.338-.338-1.69 1.016-1.354 1.352zm-9.207 5.004c-.216-.217-.691-.095-1.058.271c-.367.368-.488.841-.271 1.06l1.274 1.273c.217.218.689.096 1.056-.271c.369-.366.49-.842.273-1.058zM-.235 3.191l3.51-3.508l1.381 1.381l-3.51 3.51zm12.323 5.183l-1.431 1.428l-.739-.739l1.625-1.624a.7.7 0 0 0-.012-.992l-2.5-2.499a.7.7 0 0 0-.988-.009L6.419 5.564l-.735-.735l1.429-1.428l-1.395-1.394l-3.497 3.496l1.393 1.394l1.414-1.413l.735.735l-1.629 1.63a.7.7 0 0 0 .007.989l2.5 2.498a.701.701 0 0 0 .991.01l1.63-1.629l.739.739l-1.416 1.413l1.382 1.381l3.497-3.497zm-1.136 5.362l3.51-3.516l1.401 1.399l-3.51 3.516z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Saw(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m7.729 4.266l-1.79-2.974a.837.837 0 0 0-1.179-.016l.027 1.187L1.43 5.801L.237 5.775a.824.824 0 0 0 .015 1.171l2.992 1.778a.838.838 0 0 0 1.179.016l3.322-3.303a.824.824 0 0 0-.016-1.171M6.459 5.48L4.43 7.498c-.222.22-.631.171-.912-.107l-.82.008c-.281-.279-.33-.686-.109-.904l2.859-2.842c.224-.222.629-.174.91.106l-.008.817c.282.278.331.685.109.904m1.758 1.576L5.999 9.359l.569.594l.605-.626l1.296 1.345l-.604.626l.533.554l.604-.627l1.294 1.343l-.604.626l.535.556l.604-.626l1.293 1.343l-.604.627l.533.553l.605-.626l.297.344l-.605.626l.008.009l3.624-.923z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Scissor(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.958 4.146s-1.331-.81-1.534-.71L8.123 7.295L5.969 5.437c.071-.94-.715-2.052-2.033-2.732c-.647-.336-1.35-.521-1.978-.521c-.836 0-1.479.322-1.762.883c-.515 1.016.321 2.45 1.899 3.268c.647.335 1.35.52 1.979.52c.469 0 .866-.111 1.185-.298l1.579 1.527l-1.523 1.479c-.324-.21-.738-.335-1.234-.335c-.629 0-1.328.184-1.975.521c-1.559.812-2.373 2.232-1.852 3.234c.285.546.922.861 1.746.861c.628 0 1.329-.186 1.975-.523c1.255-.654 2.014-1.699 1.994-2.605l2.154-1.844l6.301 3.86c.203.099 1.534-.711 1.534-.711L9.43 8.084zM4.871 5.481c-.075.15-.227.256-.418.325c-.484.177-1.255.1-1.951-.26C1.493 5.024.878 4.111 1.159 3.553c.141-.277.497-.437.979-.437c.436 0 .929.133 1.392.372c.723.373 1.236.947 1.369 1.45c.052.198.051.385-.028.543M3.575 12.57c-.982.511-2.132.474-2.409-.057c-.293-.562.32-1.486 1.34-2.017c.469-.244.973-.379 1.418-.379c.231 0 .434.038.598.107c.18.074.316.185.391.328c.097.187.087.414.003.651c-.17.484-.662 1.012-1.341 1.367"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScissorLineCut(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" fill-rule="evenodd"><path d="M6.798 3.142c-.371-.222-.719-.264-.954-.115c-.474.298-.421 1.283.113 2.149c.37.597.927 1.011 1.38 1.098c.179.036.344.024.47-.057c.135-.085.222-.229.272-.404c.13-.446.001-1.124-.384-1.744c-.246-.398-.564-.726-.897-.927M5.167 8.268c-.342-.222-.791-.361-1.264-.393c-1.027-.066-1.923.396-1.957 1.012c-.032.579.833 1.186 1.823 1.248c.687.044 1.307-.148 1.659-.467c.173-.155.285-.34.297-.543a.761.761 0 0 0-.148-.469a1.463 1.463 0 0 0-.41-.388"/><path fill="currentColor" d="M16.922 9.938s-.449-.919-.651-.938l-6.515.062l-1.085-2.373c.487-.69.396-1.905-.304-3.036c-.343-.557-.799-1.021-1.278-1.312c-.643-.385-1.284-.434-1.762-.134c-.862.545-.881 2.031-.044 3.386c.342.556.796 1.022 1.279 1.311c.36.216.716.313 1.048.317l.796 1.854l-1.714.016c-.152-.311-.613-1.062-.994-1.289c-.482-.291-1.105-.639-1.757-.678c-1.572-.094-2.852.788-2.913 1.799c-.031.551.311 1.085.945 1.465c.483.289 1.107.47 1.757.509c1.267.075 2.331-.229 2.733-.935l2.367-.005l1.31 3h1.396l-1.34-3.003zM8.078 5.812c-.05.176-.137.319-.272.404c-.126.081-.291.093-.47.057c-.453-.087-1.01-.501-1.38-1.098c-.534-.866-.587-1.852-.113-2.149c.235-.148.583-.106.954.115c.333.201.651.529.896.927c.386.62.515 1.298.385 1.744M5.429 9.668c-.353.318-.973.511-1.659.467c-.99-.062-1.855-.669-1.823-1.248c.034-.615.93-1.078 1.957-1.012c.473.031.922.171 1.264.393c.179.117.315.25.41.389a.753.753 0 0 1 .148.469c-.012.202-.124.387-.297.542M10 13h1.984v1H10zm-4 0h1.984v.969H6zm8 0h1.984v.969H14zM2 13h1.984v.969H2z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScreenFul(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M4.021 5.001v5.938h7.957V5.001zm7.021 5.061H4.958V5.937h6.084z"/><path d="M6 7h3.917v1.938H6zm6-6l1.387 1.375l-1.385 1.381l1.224 1.216l1.387-1.381L16 4.965V1zM2.763 11.04l-1.376 1.359L0 11.035v3.934h4l-1.387-1.365l1.375-1.358zm-.002-6.069l1.226-1.215l-1.38-1.375L4 1H0v3.965l1.381-1.368zM13.27 11.08l-1.227 1.207l1.338 1.322L12 14.969h4v-3.934l-1.393 1.37z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScreenScale(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M5.021 5.001v5.938h7.957V5.001zm7.021 5.061H5.958V5.937h6.084z"/><path d="M7 7h3.917v1.938H7zm-6 4l1.387 1.375l-1.385 1.381l1.224 1.216l1.387-1.381L5 14.965V11zm14.763-9.96l-1.376 1.359L13 1.035v3.934h4l-1.387-1.365l1.375-1.358zm-.002 13.931l1.226-1.215l-1.38-1.375L17 11h-4v3.965l1.381-1.368zM2.27 1.08L1.043 2.287l1.338 1.322L1 4.969h4V1.035l-1.393 1.37z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Screw(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m9.984 6.943l-2.926-.671V5.014h2.926zm0 2.997l-2.926-.67V8.011l2.926.671zm0 3.029l-2.926-.671v-1.259l2.926.672zm-1.18 3.029h-.566l-1.18-1.93l2.926.671zM9 .099v1.933H7.938V.095c-1.656.21-2.926 1.337-2.926 2.706c0 1.521 6.957 1.521 6.957 0C11.969 1.438 10.647.315 9 .099"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ScrewDriver(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m14.146 12.334l-1.83-.599l-4.339-4.338l-.579.58l4.338 4.337l.637 1.872l3.067 1.832l.578-.579zM8.255 4.999L7.209 6.054c.157-.39.106-.819-.171-1.101L2.423.281C2.035-.11 1.356-.061.904.396L.393.912c-.45.457-.5 1.145-.111 1.537L4.896 7.12c.276.28.701.33 1.085.173l-.947.955l.715.72l3.22-3.249z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SdCard(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M5 9h5.906v4.938H5z"/><path d="M12.969 5.938h.969V1.063H5L2.022 6.215v1.754h1.009v1.053H2.022v7.916h11.915V9.022h-.969zM10 1.969h1V4h-1zm-2 0h1V4H8zM6 2h1.031v2.031H6zm6 13.025H3.993V7.937H12zM13 4h-1V1.969h1z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Seriff(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.717 5.374c-.692 0-1.255.553-1.278 1.241l-3.421-.503l-1.469-3.004A1.28 1.28 0 0 0 9 .672c-.708 0-1.283.575-1.283 1.282c0 .488.276.907.679 1.125L6.914 6.112l-3.354.493a1.28 1.28 0 0 0-2.559.051c0 .708.575 1.282 1.282 1.282c.333 0 .636-.131.863-.34l2.498 2.458l-.575 3.388c-.033-.002-.066-.01-.099-.01a1.284 1.284 0 0 0 0 2.565c.707 0 1.282-.575 1.282-1.282c0-.25-.074-.479-.197-.677l2.91-1.546l2.966 1.574c-.113.191-.183.41-.183.648a1.283 1.283 0 1 0 1.282-1.283c-.057 0-.11.01-.166.018l-.576-3.396l2.529-2.488a1.283 1.283 0 0 0 2.182-.912c0-.708-.574-1.281-1.282-1.281"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SewingMachine(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.875 2.125h-1.328C14.93 1.461 13.747 1 12.958 1V.042H12V1H4.042c-1.104 0-2 0-2 1v1.875c0 1.104.891 2 1.989 2V7H3v1h2v2h.958V5.875h.053c.5 0 .726-.415.837-.957l.139-.868c.019-.374-.019.255 0-.05H11v7.042H3.042a2 2 0 0 0-2 2v.916c0 1.104.896 1 2 1h10.916c1.104 0 2 .104 2-1V5.875h.917z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SewingRoll(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M12.982.986c0 .518-.41.938-.917.938h-10a.927.927 0 0 1-.917-.938c0-.519.41-.938.917-.938h10c.507 0 .917.42.917.938m0 14c0 .518-.41.938-.917.938h-10a.927.927 0 0 1-.917-.938c0-.519.41-.938.917-.938h10c.507 0 .917.42.917.938M3 3h7.917v.959H3zm0 2h7.917v.959H3zm0 2h7.917v.959H3zm0 2h7.917v.959H3zm0 2h7.917v.959H3zm0 2h7.917v.959H3z"/><path d="M16.836 15c-1.305-.208-1.679-2.154-1.78-2.98c-.088-.716-.181-4.562-2.383-5.545c-.612-.273-1.293-.402-1.878-.459c-.513-.052-.685-.039-.629-.098l-.748-.682c.294-.307.838-.28 1.418-.219c.734.07 1.518.212 2.21.529c2.621 1.203 2.949 5.732 3.024 6.354c.14 1.144.539 2.053.93 2.114z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareFive(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.021 4a1.97 1.97 0 0 0-1.706 1H9.708a1.985 1.985 0 0 0-3.697 1c0 .009.002.018.003.026l-2.052 1.23a1.97 1.97 0 0 0-.961-.257a2 2 0 1 0 2 2c0-.011-.003-.021-.003-.032l2.06-1.235c.281.152.6.247.942.247c.731 0 1.362-.396 1.708-.979h3.607c.344.596.976 1 1.706 1A1.99 1.99 0 0 0 17 6c0-1.105-.887-2-1.979-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareOne(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.969 10.062c-.83 0-1.578.342-2.117.891L5.893 8.474c.022-.146.045-.291.045-.442c0-.206-.021-.407-.06-.602l4.877-2.439a2.949 2.949 0 0 0 2.215 1.01a2.969 2.969 0 1 0-2.969-2.969c0 .029.008.058.009.087L4.81 5.72a2.943 2.943 0 0 0-1.84-.656a2.969 2.969 0 1 0 0 5.938c.758 0 1.442-.293 1.967-.761l5.09 2.545c-.008.083-.025.163-.025.247a2.969 2.969 0 1 0 5.938 0a2.973 2.973 0 0 0-2.971-2.971m.062-8.406a1.344 1.344 0 1 1 0 2.687a1.344 1.344 0 0 1 0-2.687m-10 7.75a1.405 1.405 0 1 1-.001-2.81a1.405 1.405 0 0 1 .001 2.81m10 5a1.405 1.405 0 1 1-.001-2.81a1.405 1.405 0 0 1 .001 2.81"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareThree(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.477 1.042c-1.375 0-2.47 1.062-2.47 2.442c0 .355.108.961.243 1.266L9.458 7.583L6.687 4.895a3.65 3.65 0 0 0 .304-1.328c0-1.38-1.098-2.536-2.474-2.536c-1.373 0-2.505 1.156-2.505 2.536c0 1.381 1.103 2.427 2.478 2.427c.495 0 .955-.15 1.342-.4l3.175 3.143v3.295C7.986 12.344 7 13.326 7 14.454c0 1.381 1.114 2.547 2.49 2.547c1.374 0 2.488-1.166 2.488-2.547c0-1.127-.981-2.108-2.002-2.422V8.675l3.164-3.081l-.023-.026c.396.264.869.419 1.379.419a2.495 2.495 0 0 0 2.488-2.501c0-1.382-1.132-2.444-2.507-2.444M4.486 5.023a1.52 1.52 0 0 1-1.514-1.526a1.52 1.52 0 0 1 1.514-1.526a1.52 1.52 0 0 1 1.513 1.526a1.52 1.52 0 0 1-1.513 1.526m6.554 9.493c0 .854-.692 1.547-1.548 1.547a1.547 1.547 0 1 1 0-3.094c.856 0 1.548.692 1.548 1.547m3.456-9.476c-.847 0-1.533-.706-1.533-1.579c0-.869.687-1.577 1.533-1.577c.846 0 1.534.708 1.534 1.577c0 .873-.688 1.579-1.534 1.579"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.025 5c-1.08.008-1.951.912-1.941 2.016c.004.529.213 1.007.545 1.362l-1.035 2.716a1.875 1.875 0 0 0-.562-.094c-.143 0-.281.018-.416.047L7.337 6.453c.404-.361.664-.885.664-1.475a1.97 1.97 0 0 0-1.959-1.979a1.97 1.97 0 0 0-1.957 1.979c0 .448.152.857.4 1.189L2.821 9.169a1.946 1.946 0 0 0-.789-.17c-1.088 0-1.969.896-1.969 2s.881 2 1.969 2c1.086 0 1.969-.896 1.969-2a2 2 0 0 0-.404-1.2l1.668-3.006a1.942 1.942 0 0 0 1.183.12l3.283 4.596a2.004 2.004 0 0 0-.668 1.49c0 1.104.881 2 1.969 2c1.086 0 1.969-.896 1.969-2c0-.538-.213-1.025-.553-1.385l1.031-2.709c.186.057.377.096.578.094c1.082-.008 1.951-.912 1.943-2.016c-.01-1.102-.893-1.992-1.975-1.983"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sheep(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.696 5.459c-.069-.02-.165-.262-.368-.542l.007-.02c.881-.114 1.553-.67 1.553-1.343c0-.672-.672-1.229-1.553-1.343c-.158-.64-.924-1.127-1.852-1.127c-.626 0-1.178.224-1.52.566c-.343-.343-.895-.566-1.52-.566c-.929 0-1.693.487-1.851 1.127c-.882.114-1.552.669-1.552 1.343c0 .673.67 1.229 1.552 1.343c.006.025.019.047.027.072c-.194.258-.298.472-.361.49c-1.013.296-2.292.732-2.136 1.271c.158.539 1.692.977 2.706.683c.157-.047-.821 3.324-.821 5.098c0 3.312 1.775 4.49 3.968 4.49s3.97-1.178 3.97-4.49c0-1.773-.979-5.145-.82-5.098c1.013.294 2.59-.144 2.748-.683c.158-.539-1.164-.975-2.177-1.271M7 9h1v1H7zm3 6H8v-1h2zm1-5h-1V9h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shield(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m3.068 1.976l.005 6.042c0 4.988 5.948 7.961 5.948 7.961s5.95-2.807 5.95-7.977l.001-6.026S12.308.036 9.021.036c-3.289 0-5.953 1.94-5.953 1.94m10.979 6.255c0 3.944-4.723 6.682-5.045 6.837v.015l-.015-.007l-.015.007v-.015c-.322-.155-5.045-2.893-5.045-6.837l-.004-5.664S6.181.992 8.972.979V.978l.015.001l.015-.001v.001c2.791.013 5.049 1.588 5.049 1.588zm-1.108.16c0 3.435-3.671 5.462-3.922 5.597V14s-.01-.004-.011-.006L8.994 14v-.013c-.251-.135-3.921-2.162-3.921-5.597l-.004-4.929s1.756-1.37 3.925-1.382h.023c2.17.012 3.924 1.382 3.924 1.382z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldPlus(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2 1.961v6.074c0 5.016 6.546 7.913 6.546 7.913s6.423-2.73 6.423-7.929c0-5.196.002-6.059.002-6.059S12.094.011 8.546.011S2 1.961 2 1.961m9.031 6.101H9V10H7.938V8.062H6V7h1.938V4.938H9V7h2.031z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldStar(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 0C4.688 0 2 1.958 2 1.958l.005 6.098C2.005 13.091 8.002 16 8.002 16S14 13.259 14 8.041V1.958S11.314 0 8 0m1.607 7.875l.955 2.939l-2.5-1.816l-2.502 1.816l.955-2.939l-2.5-1.816h3.091l.956-2.94l.955 2.94h3.091z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShieldTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.99.062c-3.307 0-5.988 1.944-5.988 1.944l.006 6.054c0 4.999 5.982 7.888 5.982 7.888s5.983-2.722 5.983-7.903c0-5.18.002-6.038.002-6.038S12.295.062 8.99.062m.063 13.994s-5.1-2.26-5.1-6.17l-.004-5.419S6.234.946 9.053.946z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shovel(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.732 2.509L13.495.274c-.431-.433-1.149-.415-1.603.038a.364.364 0 0 0-.092.159c-.559 2.235-.547 3.016-.454 3.323l-6.265 6.265l-1.919-1.917l-2.29 2.29c-.749.748-1.375 3.478-.077 4.775c1.297 1.297 4.024.668 4.771-.079l2.294-2.292l-1.879-1.878l6.284-6.283c.342.077 1.158.057 3.27-.47a.362.362 0 0 0 .159-.091c.453-.453.469-1.173.038-1.605m-.582.95c-1.103.311-2.385.587-2.669.533l-.435-.435c-.062-.266.216-1.561.53-2.671c.181-.134.413-.138.553.002l2.018 2.018c.138.139.134.371.003.553"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shower(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M6.004 3.051C4.857 3.194 3 4.767 3 6.953h6.953c0-2.157-1.778-3.743-2.955-3.896c-.009-1.562.007-2.14 3.47-2.14c3.22 0 3.547 1.425 3.547 4.431v10.061c0 .275.208.499.484.499a.5.5 0 0 0 .5-.499V5.348C15 2.148 14.936 0 10.469 0C6.59 0 6.044.831 6.004 3.051M3 8h.959v.916H3z"/><path d="M6 8h.959v.916H6zm3 0h.959v.916H9zm-5 2h.959v.916H4zm4 0h.959v.916H8z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignBoard(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M4.958 14.828C4.958 15.475 4.53 16 4 16c-.529 0-.958-.524-.958-1.172V1.256C3.042.61 3.471.085 4 .085c.53 0 .958.524.958 1.171z"/><path d="M15.691 2.042c.688 0 1.246.438 1.246.979c0 .541-.558.979-1.246.979H2.246C1.558 4 1 3.562 1 3.021c0-.541.558-.979 1.246-.979zm-1.698 4.002H7.965c-1.126 0-1.924-.091-1.924 1.036v4.865c0 1.127.913 1.035 2.039 1.035h5.818c1.126 0 2.039.092 2.039-1.035V7.08c.001-1.127-.818-1.036-1.944-1.036"/><path d="M7 3h.937v2.874H7zm7 0h.937v2.874H14z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignFoot(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M13.798 14.262c0 3.231-2.339.253-5.294.253c-2.951 0-5.403 2.979-5.403-.253c0-1.509 1.232-3.097 2.09-4.453c.978-1.547 1.685-2.766 3.259-2.766c1.58 0 2.364 1.294 3.345 2.852c.849 1.351 2.003 2.865 2.003 4.367"/><ellipse cx="5.91" cy="2.881" rx="1.91" ry="2.881"/><ellipse cx="10.936" cy="2.926" rx="1.936" ry="2.926"/><ellipse cx="1.871" cy="7.371" rx="1.885" ry="2.436" transform="rotate(-10.51 1.986 7.435)"/><path d="M12.115 7.305c-.345 1.326.201 2.504 1.214 2.627c1.014.126 2.116-.849 2.463-2.175c.345-1.328-.2-2.506-1.214-2.63c-1.016-.126-2.118.85-2.463 2.178"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignIn(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M11.959 4.917V1H1.096L9 3.666v12.251L1.219 14l-.215.829L9.959 17v-3.085h2v-2.873l-2.865-3z"/><path d="m11.1 8.102l2.87 2.931V8.968h2.046V6.98H13.97V5.068zM7 9h.958v.916H7z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignOut(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M9.995 4.917h1.992V0H1l8.026 2.666v12.251L1.321 13l-.215.829l8.89 2.153v-3.085l1.99.018V9.042h-1.99V4.917z"/><path d="m15.904 7l-2.87-2.932v1.987H11v1.916h2.034v2.062zM7 8h.958v.916H7z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignP(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M11.854 8H6.116C4.95 8 5.001 7 5.001 6.903V1c0-1 .083-1 1.115-1h5.738c1.166 0 1.115 0 1.115 1v5.903C12.969 8.06 13 8 11.854 8M6 1v6.031h6V1z"/><path d="M8 8h1.917v6.281H8zm1.969-5v-.969H8V6h1V4.984h.969v-.968H9V3zM10 3h.938v.969H10z"/><path d="M14 15.959H4c0-1.071 2.238-1.938 5-1.938c2.761 0 5 .867 5 1.938"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignRoadOne(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m13.447 3.902l3.365-1.34l-3.365-1.491H9.959V.167h-.902v.904H5.101v2.831h3.956v2.157H4.566L1.083 7.531l3.483 1.412h4.491v6.031h.902V8.943h3.017V6.059H9.959V3.902z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignRoadTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.633 3.938L16 2.562l-2.367-1.457l-4.654.418V.021H7v1.573l-3.979.427v2.833L7 4.428v3.063l-4.622-.437L0 8.375l2.378 1.479l4.622.41v5.674h1.979v-5.597l3.979.513V8.055l-3.979-.377V4.357z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignalOne(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.5 13.984a.499.499 0 0 1-.479-.358L6.647 8.984H4.993a.5.5 0 0 1-.486-.383l-.976-4.052l-1.048 4.06a.501.501 0 0 1-.484.375H0v-.953h1.61l1.452-5.625a.5.5 0 0 1 .484-.375h.004a.5.5 0 0 1 .482.383l1.352 5.617h1.635c.222 0 .417.146.479.358l1.005 3.346l1.02-3.348a.499.499 0 0 1 .479-.356h.687l1.347-2.736a.501.501 0 0 1 .445-.279h.004c.188 0 .359.104.444.271l1.41 2.744h1.63v.953h-1.936a.496.496 0 0 1-.444-.271l-1.095-2.131l-1.047 2.123a.497.497 0 0 1-.447.279h-.626l-1.396 4.644a.497.497 0 0 1-.478.356"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignalThree(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M17 12.965h-2a.997.997 0 0 1-.948-.684L12.998 9.19l-1.054 3.091a1 1 0 0 1-.948.684H8.998c-.449 0-.844-.3-.964-.732L6.499 6.769l-1.535 5.464a1 1 0 0 1-.964.732H1v-1.934h2.24l2.295-8.268a1 1 0 0 1 1.928 0l2.295 8.268h.518l1.773-5.316a1.001 1.001 0 0 1 1.897 0l1.774 5.316H17z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SignalTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M6.696 15.99a.671.671 0 0 1-.595-.354a.638.638 0 0 1 .29-.872c2.619-1.296 4.246-3.885 4.246-6.756c0-2.875-1.629-5.465-4.253-6.76a.64.64 0 0 1-.293-.872a.677.677 0 0 1 .899-.284c3.072 1.517 4.982 4.55 4.982 7.916c0 3.36-1.907 6.393-4.975 7.912a.668.668 0 0 1-.301.07"/><path d="M5.13 13.93a.634.634 0 0 1-.567-.348a.605.605 0 0 1-.042-.442a.63.63 0 0 1 .318-.388c1.773-.904 2.942-2.716 2.942-4.729c0-2.014-1.172-3.828-2.948-4.733a.615.615 0 0 1-.276-.829a.633.633 0 0 1 .567-.346c.102 0 .2.024.291.069a6.52 6.52 0 0 1 3.566 5.839a6.516 6.516 0 0 1-3.559 5.834a.632.632 0 0 1-.292.073"/><path d="M4.173 11.019c-.206 0-.4-.14-.491-.356c-.141-.32-.032-.713.241-.879c.558-.334.903-.998.903-1.732c0-.735-.347-1.4-.904-1.734c-.272-.164-.382-.558-.244-.877c.096-.221.284-.357.493-.357c.088 0 .173.023.252.07c.932.559 1.511 1.67 1.511 2.898c0 1.228-.579 2.338-1.51 2.896a.507.507 0 0 1-.251.071"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SiteMap(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.666 3.979h2.667A.662.662 0 0 0 11 3.324V.698a.662.662 0 0 0-.667-.656H7.666A.661.661 0 0 0 7 .698v2.626a.66.66 0 0 0 .666.655m-3.333 7.103H1.666a.657.657 0 0 0-.666.646v2.584c0 .355.298.646.666.646h2.667c.368 0 .667-.29.667-.646v-2.584a.657.657 0 0 0-.667-.646m6 0H7.666a.657.657 0 0 0-.666.646v2.584c0 .355.298.646.666.646h2.667c.368 0 .667-.29.667-.646v-2.584a.657.657 0 0 0-.667-.646m6 0h-2.667a.657.657 0 0 0-.666.646v2.584c0 .355.298.646.666.646h2.667c.368 0 .667-.29.667-.646v-2.584a.657.657 0 0 0-.667-.646M3.953 8.958h4.078v.98h1.938v-.98h4.073v.98h1.937V7.031h-6.01V5.016H8.031v2.015H2.042v2.907h1.911z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SiteMapRevert(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.334 12.02H7.667a.662.662 0 0 0-.667.655v2.626c0 .362.299.656.667.656h2.667a.662.662 0 0 0 .666-.656v-2.626a.66.66 0 0 0-.666-.655m3.333-7.104h2.667c.367 0 .666-.29.666-.646V1.686a.658.658 0 0 0-.666-.646h-2.667a.657.657 0 0 0-.667.646V4.27c0 .356.299.646.667.646m-6 0h2.667c.367 0 .666-.29.666-.646V1.686a.658.658 0 0 0-.666-.646H7.667A.657.657 0 0 0 7 1.686V4.27c0 .356.299.646.667.646m-6 0h2.667c.367 0 .666-.29.666-.646V1.686a.658.658 0 0 0-.666-.646H1.667A.657.657 0 0 0 1 1.686V4.27c0 .356.299.646.667.646m12.38 2.124H9.969v-.979H8.031v.979H3.958v-.979H2.021v2.906h6.01v2.015h1.938V8.967h5.989V6.061h-1.911z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Skateboard(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" transform="translate(0 5)"><path d="M15.252.268c-.548.953-1.697.857-2.978.857H3.699c-1.316 0-2.431.068-2.992-.938C.481-.219-.206.001.042.546c.649 1.42 2.046 2.402 3.663 2.402h8.573c1.573.001 2.938-.928 3.61-2.287c.284-.572-.392-.815-.636-.393"/><ellipse cx="11.984" cy="4.947" rx=".984" ry=".947"/><ellipse cx="3.97" cy="4.959" rx=".97" ry=".959"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Skull(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.828 12.836C15.715 11.589 17 9.804 17 7.85C17 3.993 15.063 0 9.031 0c-6.031 0-8 3.992-8 7.85c0 1.947 1.391 3.728 2.186 4.973c.309.479.451 1.872.451 1.983c0 .082.012.161.029.24c.09.625.631.927 1.297.95c0 0 .496.009.571 0c.757-.087 1.344-.411 1.423-1.049c.064.631.668 1.049 2.027 1.049s1.943-.396 2.023-1.008c.107.634.723.947 1.496 1.008c.053.004.575 0 .575 0c.704-.054 1.261-.438 1.261-1.133l-.011-.004c.001-.022.007-.043.007-.064c.001-.122.148-1.516.462-1.959M5.999 9a2 2 0 1 1 0-3.998A2 2 0 0 1 6 9zm3.011 3.848c-.824 0-1.493.439-1.493-.178c0-.616.669-1.891 1.493-1.891c.827 0 1.494 1.274 1.494 1.891c0 .618-.667.178-1.494.178M11.999 9A2 2 0 1 1 12 5a2 2 0 0 1 0 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sleep(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.953 14H1.027a.945.945 0 0 1-.867-.589a1.02 1.02 0 0 1 .168-1.063l5.531-6.375H.988C.465 5.973 0 5.572 0 5.028c0-.545.425-.985.947-.985h7.026c.374 0 .716.23.866.589c.15.359.085.774-.168 1.063l-5.492 6.387h4.773c.523 0 .947.389.947.933c.001.544-.422.985-.946.985m7.802-5.065H9.701a.64.64 0 0 1-.586-.391a.675.675 0 0 1 .107-.711l5.116-5.869h-5c-.354 0-.318-.93.035-.93h6.098c.253 0 .446.117.55.354a.675.675 0 0 1-.107.712L10.88 7.963h4.875c.354.001.354.972 0 .972"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SlideShow(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M7 0h1.906v2.219H7z"/><path d="M15.993 2.621c0-.334-.319-.605-.712-.605H.781c-.394 0-.712.271-.712.605v1.316h.962v8.049h4.434L3.17 14.724c-.161.162-.101.485.136.722l.43.429l3.327-3.77v3.863h1.851v-3.966l3.419 3.872l.429-.429c.237-.236.298-.56.136-.722l-2.365-2.738h4.406V3.936h1.056V2.621zm-1.962 8.485H1.959V3.937h12.072z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Smartphone(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.001 0H5C3.897 0 3.003 0 3.003 1v14c0 1.119.895.947 1.997.947h5.001c1.103 0 1.999.172 1.999-.947V1c0-1-.896-1-1.999-1M8.125 15.188h-1.23v-1.375h1.23zM11.037 13H4V1h7.037z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sms(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M2.002 11.984A1.988 1.988 0 0 1 .18 10.785c-.106-.248-.009-.546.221-.663c.229-.116.502-.01.611.239c.163.381.551.629.99.629c.59 0 1.068-.449 1.068-.999s-.479-.997-1.068-.997a1.988 1.988 0 0 1-1.984-1.99c0-1.096.887-1.988 1.977-1.988c.787 0 1.497.468 1.812 1.19c.108.25.011.547-.219.664c-.229.118-.501.011-.609-.237a1.064 1.064 0 0 0-.983-.623c-.586 0-1.061.447-1.061.994c0 .547.475.992 1.061.992a1.994 1.994 0 0 1 .006 3.988m11.995-.056a1.961 1.961 0 0 1-1.798-1.18c-.104-.244-.008-.536.218-.651c.226-.112.495-.009.603.235c.162.375.545.617.978.617c.581 0 1.054-.441 1.054-.98c0-.541-.473-.981-1.054-.981a1.955 1.955 0 0 1-.009-3.909c.777 0 1.479.46 1.787 1.17a.507.507 0 0 1-.216.652c-.224.116-.494.011-.601-.233a1.049 1.049 0 0 0-.971-.611c-.576 0-1.045.439-1.045.977s.469.976 1.045.976c1.089 0 1.967.879 1.967 1.961a1.96 1.96 0 0 1-1.958 1.957m-8.497.051c-.342 0-.482-.311-.482-.693V5.739c0-.302.154-.661.426-.661c.541 0 2.557 2.565 2.557 2.565s2-2.595 2.52-2.595c.273 0 .473.401.473.702v5.484c0 .382-.135.722-.477.722c-.343 0-.5-.311-.5-.692V6.891S8.2 8.985 8.001 8.985C7.809 8.961 5.985 6.86 5.985 6.86v4.394c-.001.382-.142.725-.485.725"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Snow(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m13.822 9.445l-.797-.461l-.507.879l-.934-.539l.508-.879l-.799-.461l-.508.879l-.804-.465V6.617l.804-.465l.508.88l.799-.462l-.508-.879l.934-.539l.507.879l.797-.459l-.507-.88l.88-.508l-.468-.813l-.881.51l-.508-.88l-.797.46l.508.879l-.934.539L10.607 4l-.798.461l.507.88l-.788.455l-1.559-.906v-.937h1.015v-.922H7.969V1.953h1.015v-.922H7.969V.047h-.938v.984H6.016v.922h1.015v1.078H6.016v.922h1.015v.937l-1.559.906l-.818-.472l.508-.879l-.797-.461l-.508.879l-.933-.539l.508-.879l-.799-.461l-.508.879l-.854-.492l-.468.813l.853.492l-.508.879l.799.461l.508-.879l.934.539l-.508.879l.797.461l.507-.879l.834.481v1.781l-.834.482l-.507-.88l-.797.46l.508.879l-.934.539L1.947 9l-.799.461l.508.88l-.853.492l.468.813l.854-.494l.508.88l.799-.462l-.508-.879l.933-.539l.508.879l.797-.459l-.508-.88l.832-.479l1.545.897v.921H6.016v.922h1.015v1.078H6.016v.922h1.015v1.016h.938v-1.016h1.015v-.922H7.969v-1.078h1.015v-.922H7.969v-.921l1.545-.897l.802.463l-.507.879l.798.461l.508-.879l.934.539l-.508.879l.797.461l.508-.879l.881.508l.468-.813l-.88-.508z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SoccerYard(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M5.969 3.085h-.885c-.588 0-1.063.45-1.063 1.007v3.865h1.964c.262-1.165 1.292-2.04 2.526-2.04c1.233 0 2.265.875 2.526 2.04h1.9V4.092c0-.557-.477-1.007-1.063-1.007h-.827v.962H5.969z"/><path d="M8.482 7.038c-.64 0-1.184.4-1.399.962h2.8a1.497 1.497 0 0 0-1.401-.962m.028 2.924c.656 0 1.206-.394 1.406-.937H7.104c.2.543.75.937 1.406.937"/><path d="M8.511 11.146a2.602 2.602 0 0 1-2.542-2.099H4.021v3.904c0 .555.476 1.007 1.065 1.007h.852v-1.021H11v1.021h.889c.588 0 1.064-.452 1.064-1.007V9.047h-1.9a2.604 2.604 0 0 1-2.542 2.099"/><path d="M14.439 1.042H2.498A.497.497 0 0 0 2 1.537v13.968c0 .273.223.495.498.495h11.941a.497.497 0 0 0 .498-.495V1.537a.495.495 0 0 0-.498-.495m-.418 14.02H9.979V14H7.02v1.062H2.937V1.937H7.02v1h2.933v-1h4.068z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sock(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M14.146 1.955h.22c.319 0 .578-.177.578-.396V.399c0-.218-.259-.395-.578-.395H7.58c-.319 0-.576.177-.576.395v1.16c0 .219.257.396.576.396zM11.99 13.156c0 .36.072.703.2 1.016c1.638-.412 1.847-2.04 1.735-3.539c-1.107.269-1.935 1.296-1.935 2.523m-8.908.954c0 1.368 1.226 1.742 2.392 1.827a2.634 2.634 0 0 0 .862-1.941a2.643 2.643 0 0 0-1.453-2.349c-1.71 1-1.801 1.71-1.801 2.463"/><path d="M13.856 9.961c-.102-.791-.248-1.755-.249-2.108c-.017-3.052.302-4.837.302-4.837l-5.81.041s.302 1.558.302 3.286s.709 3.241-2.155 4.612c-.188.09-.363.178-.529.264a3.4 3.4 0 0 1 .788 4.71c.948-.177 2.962-1.095 4.995-1.56a3.36 3.36 0 0 1-.222-1.178c0-1.59 1.109-2.921 2.578-3.23"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Socket(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.969 12.979a4.972 4.972 0 0 1-4.967-4.968A4.975 4.975 0 0 1 8.969 3.04a4.977 4.977 0 0 1 4.969 4.971c0 2.74-2.23 4.968-4.969 4.968m.039-9.076c-2.247 0-4.075 1.846-4.075 4.114s1.828 4.113 4.075 4.113c2.244 0 4.074-1.846 4.074-4.113c0-2.268-1.83-4.114-4.074-4.114"/><path d="M16.916 15.918H1V0h15.916zm-14.947-.887h14.062V.937H1.969z"/><path d="M7 7h.969v1.812H7zm3 0h.969v1.812H10z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SolarBlind(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><g transform="translate(3)"><path d="M11.502 12.02c-.25 0-.453-.161-.453-.36V1.041c0-.199.203-.36.453-.36s.454.161.454.36V11.66c0 .198-.204.36-.454.36"/><ellipse cx="11.45" cy="13.017" rx="1.45" ry="2.017"/><path d="M.906 9.664c-.25.055-.867.025-.867-.092V.443c0-.158 1.015-.158 1.015 0L.913 9.572c0 .041.081.072-.007.092M7 9.52V.284c0-.157 1.024.008 1.024.165v9.236C8.024 9.842 7 9.677 7 9.52"/></g><path d="M0 0h16v1.969H0zm1 3h12v1.021H1zm0 2h12v1.02H1zm0 2v.954h12V7zm0 2h12v1.993H1z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sos(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1.992 8c-.585 0-1.06-.444-1.06-.99s.475-.991 1.06-.991c.433 0 .816.243.981.619c.108.249.381.355.608.237a.513.513 0 0 0 .219-.66a1.969 1.969 0 0 0-1.809-1.188a1.982 1.982 0 0 0 .006 3.964c.591 0 1.067.448 1.067.997c0 .547-.477.994-1.067.994c-.438 0-.825-.246-.99-.624c-.106-.25-.381-.356-.608-.24c-.23.116-.327.413-.22.66a1.977 1.977 0 0 0 1.818 1.196A1.988 1.988 0 0 0 1.992 8m10.987-.012c-.578 0-1.049-.443-1.049-.987c0-.543.471-.986 1.049-.986c.431 0 .813.242.976.617c.106.247.378.354.604.236c.228-.116.323-.411.216-.659a1.956 1.956 0 0 0-1.795-1.183c-1.08 0-1.957.887-1.957 1.975c0 1.09.877 1.974 1.966 1.974c.585 0 1.059.446 1.059.991c0 .547-.474.991-1.059.991c-.434 0-.819-.244-.981-.622c-.108-.247-.377-.353-.605-.237c-.227.115-.323.41-.217.656a1.963 1.963 0 0 0 1.804 1.191a1.976 1.976 0 0 0 1.968-1.979a1.982 1.982 0 0 0-1.979-1.978m-5.45-2.962h-.074c-1.004 0-2.444.806-2.444 1.795v3.353c0 .991 1.44 1.796 2.444 1.796h.074c1.005 0 2.444-.805 2.444-1.796V6.821c.001-.989-1.439-1.795-2.444-1.795m1.506 4.927c0 .616-.895 1.117-1.52 1.117h-.047c-.624 0-1.521-.501-1.521-1.117V7.055c0-.615.897-1.117 1.521-1.117h.047c.625 0 1.52.502 1.52 1.117z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sound(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.987 14.868c0 .626-.679 1.132-1.516 1.132l-6.514-4.527c-.839 0-1.913-.508-1.913-1.133V5.682c0-.624 1.074-1.132 1.913-1.132L12.471.022c.837 0 1.516.508 1.516 1.133z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SoundMute(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m2.47 13.513l14-11.142l.873 1.097l-14 11.142zm2.483-3.601V5.01H3.128c-1.283 0-2.115 1.084-2.115 2.46v1.024c0 1.459.769 2.459 2.115 2.459h.386zm2.924 2.747l5.059 3.313c.586 0 1.06-.4 1.06-.895V7.919zM13.987.971c0-.507-.499-.92-1.115-.92L7.114 3.73C6.499 3.73 6 4.142 6 4.65v4.189l7.987-6.243z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpaceShip(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M11.023 8.441s1.299 1.967 2.578 2.594c.992.482-.109 3.392-.639 4.646c-.151.354-.357.344-.572.042c-.408-.573-1.22-1.755-2.502-1.755H5.134c-1.283 0-2.095 1.182-2.502 1.755c-.215.302-.423.312-.572-.042c-.53-1.255-1.631-4.164-.64-4.646c1.28-.627 2.579-2.594 2.579-2.594z"/><path d="M11.192 9.062C10.83 5.966 9.6 1.06 8.076.077a4.615 4.615 0 0 0-.539-.034c-.184 0-.362.013-.541.034c-1.524.983-2.755 5.89-3.116 8.985c-.07.604-.108 1.143-.108 1.565c0 1.525.489 2.645 1.26 3.357h5.01c.771-.713 1.26-1.832 1.26-3.357a14.257 14.257 0 0 0-.11-1.565M7.5 4a1.5 1.5 0 1 1-.001 3.001A1.5 1.5 0 0 1 7.5 4"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spanner(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m6.838 11.784l5.906-5.905c1.172.605 2.567.493 3.463-.402c.69-.691.924-1.682.716-2.638l-1.522 1.519l-1.356.266l-1.641-1.625l.282-1.396l1.509-1.49c-.955-.208-1.947.023-2.638.714c-.896.896-1.008 2.29-.402 3.464l-5.906 5.906c-1.173-.605-2.568-.492-3.465.402c-.688.691-.922 1.682-.715 2.637l1.523-1.519l1.355-.265l1.643 1.625l-.284 1.396l-1.509 1.49c.955.207 1.947-.023 2.637-.714c.896-.895 1.009-2.291.404-3.465"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpoonFork(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.978 4.072c0-1.846-1.173-4.01-2.481-4.01c-1.308 0-2.482 2.164-2.482 4.01c0 1.598 1.053 2.314 2.166 2.644c0 0 .279 4.538-.664 8.044c-.199.748.588 1.256 1.026 1.256c.431 0 1.212-.496 1.015-1.231c-.941-3.55-.714-8.04-.714-8.04c1.066-.373 2.134-1.118 2.134-2.673m3.836 2.533c1.074-.413 2.123-.64 2.123-2.506c0-1.01-.523-3.345-.906-4.033v3.988h-1.047v-4l-.969.001v3.999H8.953V.066c-.387.698-.93 3.007-.93 4.02c0 1.846 1.08 2.09 2.148 2.505c.041 1.28.097 5.234-.67 8.131c-.198.747.578 1.255 1.009 1.255c.425 0 1.192-.496.999-1.23c-.754-2.881-.746-6.779-.717-8.141z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spray(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m5.254.021l-.223.948l1.953.015zm7.709 7.003v-.95S11.3 3.04 10.936 3.04h-.967L9.985.015H8.016v1.966h1v1.06h-.951c-.363 0-2.059 3.033-2.059 3.033v9.269c0 .37.324.672.725.672h5.509c.399 0 .724-.302.724-.672V7.045h.018l-.019-.02zm-.947 8.007h-1.022V7h1.022zm-6.989-13l.232.914l1.679-.883z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Square(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M.002 1c0-.553.444-1 .993-1h13.972c.549 0 .993.447.993 1v14c0 .553-.444 1-.993 1H.995a.996.996 0 0 1-.993-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareChecked(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="m3.017 7.904l1.871-1.873l2.285 2.291l6.092-6.101l1.78 1.783l-7.961 7.973z"/><path d="M12.021 10.116v2.926H1.979V2.958h7.968l1.074-.933H1.002v11.936h11.977L13 9.083z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareDashedOne(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.048 1.976h1.993v1.988h.907v-2.92h-2.9zM.923 3.995V1.976h2.028v-.99H.027v3.009zm2.028 12.022H.964v-1.95H.027v2.928h2.924zM15 6v2.01h.989V6zm0 4v1.96h.982V10zM9 1v.965h1.984V1zM5 1v.973h1.995V1zm4 15v.973h1.958V16zm-4 0v.993h1.987V16zM0 6v1.99h.98V6zm0 4v1.961h.954V10zm15.006 4v2.073H13v.945h3.003V14z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareDashedTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10 .834h5.083v5.041h.834V0H10zM.834 5.875V.834h5.083V0H0v5.875zm14.249 4.25v5.041H10V16h5.917v-5.875zm-9.166 5.041H.834v-5.041H0V16h5.917z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareDelicious(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 0v16h16V0zm7.938 15.016V8.01H1.959V.958H9.01v7.011h7.006v7.047z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareEightAnglePoint(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16 4V1.041h-2.959v1.021H10V1.041H6.041v1.021H3V1.041H.041V4h1.011v3.041H.041V11h1.011v3.041H.041V17H3v-1.021h3.041V17H10v-1.021h3.041V17H16v-2.959h-1.031V11H16V7.041h-1.031V4zM9.062 1.938v1.125H6.978V1.938zm6 0v1.125h-1.084V1.938zm-14.083 0h1.084v1.125H.979zm-.01 6h1.125v2.125H.969zm1.093 8.146H.978v-1.125h1.084zm4.917-.022v-1.125h2.084v1.125zm6.989 0v-1.114h1.126v1.114zm1.126-5.958h-1.125V7.938h1.125zm-2.053-3.063V11h1.011v3.041h-1.011v1.021H10v-1.021H6.041v1.021H3v-1.021H1.969V11H3V7.041H1.969V4H3V2.979h3.041V4H10V2.979h3.041V4h1.011v3.041z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareFour(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8 4.769c0 1.187-.976 2.148-2.182 2.148H3.181c-1.204 0-2.18-.962-2.18-2.148V2.233c0-1.187.976-2.148 2.18-2.148h2.637C7.024.085 8 1.047 8 2.233zm0 9.049A2.181 2.181 0 0 1 5.818 16H3.181a2.181 2.181 0 0 1-2.18-2.182v-2.637c0-1.205.976-2.182 2.18-2.182h2.637C7.024 8.999 8 9.976 8 11.181zm9-9.049c0 1.187-.976 2.148-2.182 2.148h-2.637c-1.204 0-2.18-.962-2.18-2.148V2.233c0-1.187.976-2.148 2.18-2.148h2.637C16.024.085 17 1.047 17 2.233zm0 9.049A2.181 2.181 0 0 1 14.818 16h-2.637a2.181 2.181 0 0 1-2.18-2.182v-2.637c0-1.205.976-2.182 2.18-2.182h2.637c1.206 0 2.182.977 2.182 2.182z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareFourAnglePoint(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.959 3V.041H13v1.021H2.959V.041H0V3h1.011v10.041H0V16h2.959v-1.021H13V16h2.959v-2.959h-1.031V3zM13.938.938h1.084v1.125h-1.084zm-13 0h1.084v1.125H.938zm1.083 14.124H.937v-1.125h.073v.011h.917v-.011h.094zm13-1.124v1.125h-1.084v-1.125h.073v.011h.917v-.011zm-1.01-.897H13v1.021H2.959v-1.021H1.928V3h1.031V1.979H13V3h1.011z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquarePlus(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.664 0H1.336C.598 0 .002.598.002 1.335v13.329C.002 15.401.598 16 1.336 16h13.328c.738 0 1.336-.599 1.336-1.336V1.335C16 .598 15.402 0 14.664 0M9.023 9.016v2.55c0 .631-2.025.631-2.025 0v-2.55H4.421c-.633 0-.632-2.01 0-2.01h2.577V4.461c0-.633 2.025-.633 2.025 0v2.545h2.556c.633 0 .633 2.01 0 2.01z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SquareStar(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.628.021H2.352a1.33 1.33 0 0 0-1.33 1.331v13.275c0 .734.596 1.331 1.33 1.331h13.276c.734 0 1.33-.597 1.33-1.331V1.352a1.33 1.33 0 0 0-1.33-1.331m-3.23 13.484l-3.385-1.883l-3.385 1.883l.646-3.987l-2.737-2.824l3.783-.58l1.693-3.628l1.692 3.628l3.784.58l-2.738 2.824z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stamps(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M4 5h9.979v6H4z"/><path d="M16.969 11.984V11h-1.016V9.984h1.016V9h-1.016V7.984h1.016V7h-1.016V5.984h1.016V5h-1.016V3.984h1.016V3h-.984v-.984H15v1.016h-1.016V2.016H13v1.016h-1.016V2.016H11v1.016H9.984V2.016H9v1.016H7.984V2.016H7v1.016H5.984V2.016H5v1.016H3.984V2.016H3v1.016H2V4h-.984v.984H2V6h-.984v.984H2V8h-.984v.984H2V10h-.984v.984H2V12h-.984v.984H2v1h.984v-1H4v1h.984v-1H6v1h.984v-1H8v1h.984v-1H10v1h.984v-1H12v1h.984v-1H14v1h.984v-1h.969v-1zM15 12.016H2.969V3.985H15zM16 13h.984v.969H16zM1 2h.953v.984H1z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarCross(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.601 6.4L8 0L6.398 6.4L0 8l6.398 1.601L8 16l1.601-6.399L16 8zM8 9.334a1.333 1.333 0 1 1 .003-2.667A1.333 1.333 0 0 1 8 9.334"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarStick(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M14.775 1.206L13.73 4.289l1.946 2.607l-3.254-.041l-1.879 2.656l-.966-3.107l-3.106-.965l2.656-1.88L9.086.305l2.606 1.946z"/><path d="M1.852 15.533c-.39.391-.922.492-1.188.226l-.406-.405c-.266-.266-.164-.797.227-1.188L10.788 3.863c.391-.391.922-.492 1.188-.227l.406.406c.266.266.165.797-.227 1.188zm11.659-1.584s.162-1.059-.61-1.823c-.773-.762-1.87-.626-1.87-.626s1.266.02 1.895-.603c.631-.623.586-1.847.586-1.847s.133 1.35.634 1.847c.505.496 1.847.603 1.847.603s-1.261.102-1.819.652c-.56.553-.663 1.797-.663 1.797m-5 2s.162-1.059-.61-1.823c-.773-.762-1.87-.626-1.87-.626s1.266.02 1.895-.603c.631-.623.586-1.847.586-1.847s.133 1.35.634 1.847c.505.496 1.847.603 1.847.603s-1.261.102-1.819.652c-.56.553-.663 1.797-.663 1.797m-5-11s.162-1.059-.61-1.823c-.773-.762-1.87-.626-1.87-.626s1.266.02 1.895-.603c.631-.623.586-1.847.586-1.847s.133 1.35.634 1.847c.505.496 1.847.603 1.847.603s-1.261.102-1.819.652c-.56.553-.663 1.797-.663 1.797"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stele(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M1 14v1.99h15.003V14zm10-8h.905v.898H11zM5 6h.951v.905H5z"/><path d="M2 5.412v7.589h13V5.412C14.999.068 10.014.024 8.06.024C6.106.024 2 .068 2 5.412m4.967 4.574H5.894V7.954h-.887v2.003h-1.04V4.986h2.986v2.028H6.02v.86h.947zm1.986-.032h-1v-5h1zm2.085-1.958v1.958H10v-5h2.995L13 7.984z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stelescope(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="m2.606 3.764l-2.2-2.2L1.8.17L4 2.372z"/><path d="M1.082 3.647L4.33 6.896l2.477-2.477L3.559 1.17zm3.623 4.532l4.78 5.729l4.247-4.248l-5.729-4.78zm11.607 2.901l-5.15 5.15l-.81-.808l5.15-5.151z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StereoJack(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M4.929 5.762L3.86 6.827c-.316.313-.354.785-.088 1.05l1.334 1.33c.268.265.74.224 1.055-.088l1.07-1.067zm4.841-.238a.663.663 0 0 0-.02-.942l-1.336-1.33a.671.671 0 0 0-.945-.019L5.633 5.062l2.3 2.292zm-8.717 5.008l-.867.862c-.265.266-.252.71.033.994l.34.338c.283.283.73.297.996.031l.866-.861z"/><path d="M3.242 11.077L1.855 9.694l2.299-2.291l1.389 1.383zm12.432 3.897c-.523 0-1.794.083-2.864-.788c-1.214-.991-1.691-2.195-1.691-4.43c0-1.22.567-2.731 1.146-3.521c.502-.686.875-1.407.875-2.182c0-1.136-.172-1.812-.684-2.102c-.492-.279-1.127-.08-1.308-.014L8.534 4.906l-.518-.463l2.762-3.081c.047-.021 1.141-.497 2.062.026c.744.424 1.026 1.269 1.026 2.664c0 .975-.46 1.812-.999 2.549c-.568.779-1.023 2.135-1.023 3.154c0 2.03.391 3.067 1.439 3.928c1.132.929 2.559.623 2.572.623l.078.656a2.027 2.027 0 0 1-.259.012"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stethoscope(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.938 10A1.968 1.968 0 1 0 13 10c0 .736.408 1.37 1.007 1.708c-.558 1.911-1.947 2.979-3.573 2.979c-1.953 0-3.567-1.295-3.807-3.929c1.823-.133 4.331-1.249 4.331-3.555V3.624c0-.857-.356-1.657-.989-2.188a1.456 1.456 0 0 0-2.912.058A1.459 1.459 0 0 0 9.322 2.71c.185.248.694.59.694.914v3.579c0 1.671-2.783 2.827-3.752 2.827h-.797c-.654 0-3.498-.973-3.498-2.827V3.624c0-.32.598-.671.779-.918A1.459 1.459 0 1 0 3.541.024c-.791 0-1.432.631-1.455 1.417c-.622.529-1.07 1.333-1.07 2.183v3.579c0 2.356 2.365 3.469 4.207 3.56c.267 3.369 2.552 5.213 5.26 5.213c2.295 0 4.199-1.51 4.896-4.052A1.967 1.967 0 0 0 16.938 10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Store(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.021 7.028v5.014H4V7h-.965v8H15V7.028zM4 0h9.902v.929H4zm10.295 3.001l-.311-.955H4.021l-.336.955l-2.674 3.301l1.334.668h1.634V5.938h1.042V6.97h1.958V5.938h1.042V6.97h1.958V5.938h1.042V6.97h1.958V5.938h1.042V6.97h1.657l1.291-.668z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stove(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M1 0v3h14.958V0zm2 2H2V1h1zm2 0H4V1h1zM1 16h14.958V4.042H1zM5 6h7v1H5zM4 7.958h9v6H4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Strawberry(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.367 2.084c.055-.154.58-.761.58-1.167S8.726 0 8.049 0c-.682 0-.982.511-.982.948c0 .396.525.97.582 1.114c-7.258 0-5.432 5.919-5.432 5.919s2.031-3.882 3.84-4.237c-.002.096-.006.191-.006.287c0 2.734 1.996 4.951 1.996 4.951s1.809-2.217 1.809-4.951c0-.103-.004-.203-.006-.305c2.023.331 3.902 4.305 3.902 4.305s1.957-5.947-5.385-5.947"/><path d="M10.873 5.608c0 2.392-2.838 4.798-2.838 4.798S5.14 7.979 5.14 5.558c-1.193.947-2.016 2.833-2.589 3.588c0 0 2.302 6.854 5.476 6.854s5.446-6.729 5.446-6.729c-.455-.601-1.317-2.617-2.6-3.663"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StreetTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M17.02 14.046v-13h-16v13zM14 6.953h2v1h-2zM9 1.992h4V4.02H9zM9 5h4.021v2H9zm0 3h4.021v2H9zm0 2.958h4v2H9zM5 6.953h2V8H5zm-3 0h2V8H2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Strolley(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" transform="translate(1)"><ellipse cx="4.438" cy="14.459" rx="1.438" ry="1.459"/><ellipse cx="12.457" cy="14.419" rx="1.457" ry="1.419"/><path d="M15 10.016H3.953V3.778L1.781.41A.995.995 0 1 0 .174 1.584l1.849 2.901v5.531h-.055c-.553 0-1 .44-1 .984c0 .543.447.984 1 .984H15c.553 0 1-.441 1-.984a.991.991 0 0 0-1-.984"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StrolleyArrowDown(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><g transform="translate(1)"><ellipse cx="4.438" cy="14.459" rx="1.438" ry="1.459"/><ellipse cx="12.457" cy="14.419" rx="1.457" ry="1.419"/><path d="M15 10.016H3.953V3.778L1.781.41A.995.995 0 1 0 .174 1.584l1.849 2.901v5.531h-.055c-.553 0-1 .44-1 .984c0 .543.447.984 1 .984H15c.553 0 1-.441 1-.984a.991.991 0 0 0-1-.984"/></g><path d="m8.198 5.243l2.189 2.581a.795.795 0 0 0 1.143 0l2.188-2.581c.315-.324.364-.899.05-1.223h-1.804V2.045c0-.573-.434-1.037-.968-1.037c-.535 0-.969.464-.969 1.037V4.02H8.246c-.314.325-.363.899-.048 1.223"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StrolleyArrowUp(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><g transform="translate(1)"><ellipse cx="4.438" cy="14.459" rx="1.438" ry="1.459"/><ellipse cx="12.457" cy="14.419" rx="1.457" ry="1.419"/><path d="M15 10.016H3.953V3.778L1.781.41A.995.995 0 1 0 .174 1.584l1.849 2.901v5.531h-.055c-.553 0-1 .44-1 .984c0 .543.447.984 1 .984H15c.553 0 1-.441 1-.984a.991.991 0 0 0-1-.984"/></g><path d="m13.768 3.757l-2.189-2.513a.811.811 0 0 0-1.143 0L8.247 3.757c-.314.315-.363.876-.049 1.19H10v1.922c0 .559.434 1.01.969 1.01c.535 0 .969-.451.969-1.01V4.947h1.781c.315-.315.364-.875.049-1.19"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StrolleyPlus(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><g transform="translate(1)"><ellipse cx="4.438" cy="14.459" rx="1.438" ry="1.459"/><ellipse cx="12.457" cy="14.419" rx="1.457" ry="1.419"/><path d="M15 10.016H3.953V3.778L1.781.41A.995.995 0 1 0 .174 1.584l1.849 2.901v5.531h-.055c-.553 0-1 .44-1 .984c0 .543.447.984 1 .984H15c.553 0 1-.441 1-.984a.991.991 0 0 0-1-.984"/></g><path d="M12.959 4.025H9.063c-.558 0-1.01.434-1.01.969c0 .535.452.969 1.01.969h3.896c.559 0 1.01-.434 1.01-.969c0-.535-.451-.969-1.01-.969"/><path d="M11.979 6.942V3.046c0-.558-.434-1.01-.969-1.01c-.535 0-.969.452-.969 1.01v3.896c0 .559.434 1.01.969 1.01c.535 0 .969-.451.969-1.01"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StrolleyRemove(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><g transform="translate(1)"><circle cx="4.469" cy="14.469" r="1.469"/><ellipse cx="12.467" cy="14.45" rx="1.467" ry="1.45"/><path d="M15 10.016H3.953V3.778L1.781.41A.995.995 0 1 0 .174 1.584l1.849 2.901v5.531h-.055c-.553 0-1 .44-1 .984c0 .543.447.984 1 .984H15c.553 0 1-.441 1-.984a.991.991 0 0 0-1-.984"/></g><path d="M12.959 5.025H9.063c-.558 0-1.01.434-1.01.969c0 .535.452.969 1.01.969h3.896c.559 0 1.01-.434 1.01-.969c0-.535-.451-.969-1.01-.969"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Subway(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M5 15.979h-.979l1-4H6zm8 0h-.979l-1-4H12z"/><path d="M5 14h6.9v.926H5zM11.988.038H5.095c-2.242 0-3.054.688-3.054 2.775v7.784c0 1.848.813 2.441 2.984 2.441h7.725c1.574 0 2.291-.602 2.291-2.44V2.812c0-2.087-.81-2.775-3.053-2.775zM6.213 1h4.574c.117 0 .213.226.213.5c0 .273-.096.5-.213.5H6.213C6.096 2 6 1.773 6 1.5c0-.274.096-.5.213-.5m-.197 10.016H3.969V8.985h2.047zm7 0h-2.047V8.985h2.047zm.037-5.235c0 1.064-.166 1.257-1.128 1.257H5.091c-.964 0-1.127-.146-1.127-1.21V4.156c0-1.062.163-1.198 1.127-1.198h6.834c.962 0 1.128.198 1.128 1.261z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Suitcase(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5 3v12.969h7.99V3zm9.037.031V16h.711C15.992 16 17 15.021 17 13.812V5.216c0-1.205-1.008-2.185-2.252-2.185zM1 5.217v8.596c0 1.208 1.018 2.188 2.276 2.188h.685V3.032h-.685C2.018 3.031 1 4.011 1 5.217M9.968.047H8.019c-1.25 0-2.009.871-2.009 1.944h.961c0-.588.327-1.037 1.048-1.037h1.949c.719 0 1.057.45 1.057 1.037h.957c0-1.073-.763-1.944-2.014-1.944"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SuitcasePerson(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M2.437 3C1.643 3 .999 3.655.999 4.465v8.07c0 .81.644 1.465 1.438 1.465h13.125c.795 0 1.438-.655 1.438-1.465v-8.07C17 3.655 16.357 3 15.562 3zm6.56 1.932c.783 0 1.42.698 1.42 1.559s-.637 2.51-1.42 2.51c-.783 0-1.419-1.648-1.419-2.51c0-.861.636-1.559 1.419-1.559m3.934 7.185H5.072s-.134-3.171 1.883-3.171c.421.611 1.001 1.2 2.065 1.2c1.066 0 1.574-.594 1.992-1.214c2.268.001 1.919 3.185 1.919 3.185"/><path d="M7.143 2.617c0-.35.318-.635.711-.635h2.328c.392 0 .71.285.71.635v.48h1.025V2.12c0-.607-.554-1.103-1.233-1.103H7.278c-.682 0-1.234.495-1.234 1.103v.977h1.1v-.48z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Syringe(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" fill-rule="evenodd"><path fill="currentColor" d="m.032 4.258l.647.648L4.892.693L4.245.046z"/><path fill="currentColor" d="M1.844 3.207L3.45 4.813l1.4-1.4l-1.606-1.605z"/><path d="M5.931 5.487h1.104v6.062H5.931zm1.917-2.366h1.105v6.062H7.848z"/><path fill="currentColor" d="m13.548 15.369l1.749-.928l-2.629-2.731l-1.391 1.39zM6.79 1.974L1.633 7.13c-.131.131-.154.452-.02.585c.131.133.494.283.626.152l.513-.514l5.672 5.671l-.014.014l.671.673l.68-.68l.741.742l2.799-2.8l-.741-.741l.722-.721l-.672-.672l-.016.015l-5.671-5.672l.47-.47c.131-.132-.029-.504-.162-.637c-.131-.132-.309-.229-.441-.101M5.493 4.581l.781-.78l4.287 4.286l-.781.781zM3.372 6.703l.781-.781l4.286 4.287l-.78.78zm5.687 5.687l-.75-.751l2.888-2.888l.751.751z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tshirt(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M9.021 1.062c-.917 0-1.705-.27-2.04-1.062c-.206.012-.735.033-.935.067c.383 1.088 1.752 1.873 2.975 1.873s2.58-.785 2.963-1.873c-.199-.034-.719-.056-.922-.067c-.334.793-1.123 1.062-2.041 1.062"/><path d="M13.396.289c-.82 1.62-2.479 2.742-4.412 2.742c-1.924 0-3.572-1.11-4.398-2.717C2.33 1.33 1.043 3.83 1.043 6.466v.486h2.988v7.983h9.948V6.952H17v-.486c0-2.663-1.306-5.183-3.604-6.177"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tablet(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.995.043h-8.99A2.005 2.005 0 0 0 2 2.047v11.95C2 15.102 2.899 16 4.005 16h8.99A2.005 2.005 0 0 0 15 13.997V2.047A2.005 2.005 0 0 0 12.995.043m-2.912 14.999H6.93v-1.114h3.153zM13 13H4V1.97h9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tag(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m5.379.681l-.09.09l-.034-.035c-.854-.854-2.275-.818-3.173.08l-.255.254c-.896.897-.933 2.318-.078 3.173l.034.034l-.164.165c-.773.772-.801 1.999-.06 2.738l8.325 8.328c.742.74 1.967.712 2.742-.061l3.758-3.758c.772-.773.801-1.999.061-2.739L8.117.622c-.74-.74-1.967-.713-2.738.059m-.856 4.427a1.596 1.596 0 0 1-1.592-1.6c0-.881.714-1.597 1.592-1.597a1.599 1.599 0 0 1 0 3.197"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagOne(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.546 15.951c-.509.509-.515-4.982-.515-4.982s-5.493-.007-4.983-.516L10.061.439c.508-.508 1.357-.484 1.896.053l3.55 3.551c.539.539.561 1.388.054 1.896z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagOnePlus(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m13.613 1.047l-4.502.003l-7.779 7.77a1.14 1.14 0 0 0 0 1.607l5.244 5.242a1.135 1.135 0 0 0 1.605 0l7.787-7.707l.002-4.534zM4.984 9.014l1.023-1.023l-.941-.935l.973-.968l.939.938l1.025-1.021l.988.985l-1.024 1.021l.94.937l-.973.967l-.938-.937l-1.024 1.02zm2.137 4.543l-.789-.788l4.243-4.243l-.727-.725l.786-.788l1.514 1.514zM15.942.36h.984v1.796h-.984z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TagPrice(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M15.46 8.894L7.127.545c-.74-.74-1.967-.714-2.74.061l-.089.089l-.034-.034C3.409-.195 1.987-.16 1.088.74L.833.996C-.066 1.895-.1 3.319.755 4.175l.035.035l-.165.166c-.773.773-.801 2.002-.06 2.744l8.333 8.349c.74.741 1.969.714 2.74-.06l3.763-3.77c.774-.775.8-2.003.059-2.745M3.433 5.168c-.866 0-1.567-.734-1.569-1.641c0-.906.703-1.64 1.569-1.64S4.997 2.622 5 3.527c-.003.907-.703 1.641-1.567 1.641m9.001 7.909l-1.084-1.085c-.652.463-1.469.609-2.184.362a.633.633 0 0 1-.373-.333a.656.656 0 0 1-.02-.516a.683.683 0 0 1 .855-.41c.238.082.51.047.761-.068L9.136 9.769c-.953.682-2.219.662-2.992-.109c-.778-.777-.795-2.047-.104-2.998L4.978 5.599l.629-.629l1.069 1.065c.658-.461 1.476-.607 2.192-.359a.65.65 0 0 1 .392.85a.69.69 0 0 1-.857.409c-.238-.083-.508-.049-.758.063l1.262 1.257c.947-.665 2.198-.638 2.965.128c.767.767.797 2.013.133 2.955l1.08 1.076z"/><path d="M7.124 8.722c.268.267.698.295 1.081.107L7.017 7.636c-.189.385-.162.817.107 1.086m3.8.629c-.258-.259-.668-.292-1.041-.124l1.166 1.162c.167-.371.133-.781-.125-1.038"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tank(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M4 1V0h-.969v1L3 4.031h1.969V1zm5.031-1v1.031H8v4h2V0z"/><path d="M2 2h10.037v1H2zm0 6h4v4.004H2zm2.568-4H3.453C2.109 4 2 3.96 2 5.142v1.829h4V5.142C6 3.96 5.911 4 4.568 4M3.453 15.969h1.115c1.343 0 1.432.041 1.432-1.143v-1.822H2v1.822c0 1.184.109 1.143 1.453 1.143M9.557 4.051H8.465C7.15 4.051 7 4.01 7 5.193v1.812h4V5.193c0-1.183-.129-1.142-1.443-1.142M8.465 16.007h1.092c1.314 0 1.443.041 1.443-1.143v-1.851H7v1.851c0 1.184.149 1.143 1.465 1.143M7 8h4v4H7zM1 6v3.006h.969V6zm10 0v3.01h.969V6zm0 5v2.971h.969V11zM1 11v3.028h.969V11z"/><path d="M15.75 14.031c-.506 0-.955-.195-1.336-.58c-1.725-1.744-1.349-4.514-1.349-6.392c0-.729-.065-1.469-.065-2.024c0-1.348-.219-2.104-.962-2.104l-.007-.837S14 1.5 14 5.028c0 .562-.003 1.23-.012 1.966c-.041 3.481-.294 4.308 1.137 5.755c.191.193.389.283.625.283z"/><ellipse cx="15.484" cy="13.469" rx="1.484" ry="1.469"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Targer(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><ellipse cx="5.469" cy="9.441" rx="1.469" ry="1.441"/><path d="m14.994 3.671l-1.38 1.394l-2.149-.812l-5.746 5.691c-.117.109-.387.031-.554-.139c-.163-.168-.216-.416-.101-.528l5.699-5.721l-.826-2.103L11.396.025l.597 3.008c.004.004.012.006.016.012c.026.03 2.985.626 2.985.626"/><path d="M5.514 14.994C2.499 14.994.047 12.529.047 9.5c0-3.03 2.452-5.494 5.467-5.494c3.014 0 5.466 2.464 5.466 5.494c-.001 3.029-2.453 5.494-5.466 5.494m0-10.094C2.993 4.9.943 6.963.943 9.5c0 2.537 2.051 4.6 4.571 4.6c2.52 0 4.57-2.063 4.57-4.6c0-2.537-2.05-4.6-4.57-4.6"/><path d="M5.507 12.99a3.476 3.476 0 0 1-3.472-3.471a3.476 3.476 0 0 1 3.472-3.471a3.475 3.475 0 0 1 3.47 3.471a3.475 3.475 0 0 1-3.47 3.471m0-6.066A2.598 2.598 0 0 0 2.911 9.52a2.598 2.598 0 0 0 2.596 2.596A2.598 2.598 0 0 0 8.102 9.52a2.598 2.598 0 0 0-2.595-2.596m-3.593 8.86a.51.51 0 0 1-.722.013a.51.51 0 0 1 .014-.72l1.467-1.467c.204-.203.911.504.708.707zm7.904-.708a.51.51 0 0 1 .012.722a.508.508 0 0 1-.719-.014l-1.467-1.467c-.203-.204.504-.911.707-.708z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TeaCup(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 2h11.968v1.976H0zm9.469 11.935H2.54C.06 13.935.02 11.5.02 11.5V5.042h11.968v6.347s-.039 2.546-2.519 2.546m3.418-9.9h-.877v1.902c.678 0 2.095-.375 2.095 1.161v.778c0 1.149-.729 1.206-2.095 1.206v1.859h.877c1.709 0 3.092-1.091 3.092-2.436V6.468c0-1.343-1.383-2.433-3.092-2.433"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Teeth(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.715 9.02c0 3.025-.564 6.867-2.281 6.867c-2.814 0-1.518-5.886-3.943-5.886s-1.605 5.917-3.922 5.917c-1.629 0-2.277-3.916-2.277-6.898c0-1.807-2.393-4.771-.62-7.392C3.846-1.586 5.94.952 8.437.952c2.535 0 4.512-2.485 6.828.676c1.836 2.509-.55 5.616-.55 7.392"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TelephoneBox(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13 15.021V4h-1V2.917h.852C12.852 1.306 11.021.083 8 .083c-3.021 0-5 1.223-5 2.834h1.023V4H3v11.021H2V16h12v-.979zM5 3h6v1H5zm1 12.021H5v-2.04h1zM6 12H5v-2h1zm0-3H5V7h1zm0-3H5V5h1zm3 9.021H7v-2.04h2zm-2-3.063v-2h2v2zM7 9V7h2v2zm2-3.042H7v-1h2zm.958-.979H11v1.04H9.958zM11 7v2h-1V7zm.021 8.021H10v-2.04h1.021zm0-3.002H10V9.98h1.021z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Telescope(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="m13.058 1.4l.659 1.949l-.943.303l-.429-1.273l-8.593 2.906l.686 2.027l-.933.289l-.687-2.027l-.526.215c-.11.22-.132.503-.327.853l-2.004.682l.847 2.507l1.794-.645c.377.162.783.268 1.014.368l10.947-3.696l.216.637l1.183-.4l-1.719-5.093z"/><path d="M8.916 13.09V8.83c.391-.229.385-.699.387-1.187a1.32 1.32 0 0 0-1.312-1.326a1.32 1.32 0 0 0-1.324 1.316c-.002.494-.003.921.396 1.149v4.255c-1.896.102-3.07 1.287-3.07 1.931h7.999c-.001-.643-1.18-1.778-3.076-1.878"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Television(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.381 5.06H9.468l2.739-3.79l-1.012-.223L8.303 5.06h-.008L5.403 1.047l-1.014.223l2.739 3.79H.656a.582.582 0 0 0-.593.571v10.785c0 .317.265.572.593.572h14.723a.582.582 0 0 0 .594-.572V5.631a.578.578 0 0 0-.592-.571M13 14.67c0 .206-.152.372-.342.372H2.263c-.188 0-.341-.166-.341-.372V7.304c0-.204.152-.372.341-.372h10.395c.189 0 .342.168.342.372zm2 .372h-1V14h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TennisRacketBall(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" transform="translate(1)"><circle cx="13.953" cy="13.953" r="1.953"/><path d="M14.754 1.265C12.76-.733 9.038-.26 6.462 2.32c-1.567 1.564-2.351 3.555-2.28 5.321a.337.337 0 0 0-.106.327c.462 2.098-1.378 4.256-1.398 4.279a.304.304 0 0 0-.054.096l-2.38 2.38c-.288.288-.299.744-.024 1.02l.022.02c.272.275.73.265 1.017-.025l2.455-2.452c.167-.127 2.535-1.871 4.24-1.386a.338.338 0 0 0 .306-.062c1.797.111 3.841-.677 5.442-2.28c2.576-2.577 3.048-6.298 1.052-8.293M6.225 9.84a3.085 3.085 0 0 1-.604-.869l1.355-1.356l1.429 1.43l-1.37 1.371a3.082 3.082 0 0 1-.81-.576m1.459-2.931l1.414-1.416l1.43 1.43l-1.416 1.414zm1.427-2.814L7.652 2.636a7.137 7.137 0 0 1 1.709-1.121l1.166 1.165zm2.108-.723l1.43 1.43l-1.415 1.414l-1.429-1.428zm-4.274-.031l1.46 1.461L6.99 6.216L5.825 5.05c.275-.59.648-1.17 1.12-1.709m4.274 4.274l1.504 1.503a7.187 7.187 0 0 1-1.709 1.121L9.805 9.03zm.707-.706l1.415-1.416l1.208 1.209a7.082 7.082 0 0 1-1.12 1.709zm3-1.242l-.879-.879l1.008-1.009c.088.593.041 1.238-.129 1.888m-.776-3.753c.238.239.426.514.574.811l-1.37 1.37l-1.429-1.429l1.355-1.354c.32.15.615.347.87.602m-1.945-.913l-.972.973l-.836-.836a4.815 4.815 0 0 1 1.808-.137M5.449 6.087l.834.836l-.973.973a4.8 4.8 0 0 1 .139-1.809m-1.551 6.225c-.023-.033-.036-.071-.064-.1l-.021-.021c-.029-.029-.066-.041-.099-.064c.363-.543.835-1.393 1.051-2.375a4.005 4.005 0 0 0 1.491 1.52c-.934.226-1.799.69-2.358 1.04m4.192-1.567l1.008-1.009l.879.879c-.651.171-1.295.218-1.887.13"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TentCamp(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.021 14.023c-2.701-1.25-5.112-7.249-6.289-10.164l1.205-2.284a.355.355 0 0 0-.1-.454c-.138-.094-.316-.043-.4.114L9.436 3.132c-.225-.542-.375-.871-.436-.871c-.06 0-.207.319-.426.847L7.563 1.192c-.084-.157-.264-.208-.4-.114a.355.355 0 0 0-.1.454l1.215 2.302C7.107 6.733 4.699 12.739 2 14.013V12.02H1V15h4.666C8.268 13.082 9 6.926 9 6.926s.943 6.156 3.334 8.074L17 14.997V12.02h-.979z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TentOne(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16 9v2.131c-2.678-1.054-5.068-5.984-6.249-8.422L10.916.5a.357.357 0 0 0-.1-.455c-.137-.094-.316-.043-.4.115l-.992 1.88c-.218-.438-.364-.703-.424-.703s-.206.265-.424.703L7.584.16C7.5.002 7.32-.049 7.184.045a.357.357 0 0 0-.1.455l1.165 2.209c-1.186 2.447-3.59 7.406-6.28 8.434V9H1v6.977h.969v-.028h4v-4.193C8.321 9.757 9 6.005 9 6.005s.861 3.71 3.016 5.717v4.228H16v.028h1V9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TestTube(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M12.921.969c0-.937-.091-.938-1.05-.938H5.13c-.958 0-1.052.001-1.052.938h.953v2.94L2.019 14.238c0 .939.777 1.699 1.736 1.699h9.489c.958 0 1.736-.76 1.736-1.699L11.992 3.879l-.014-2.91zm1.048 12.637c.271.884-.203 1.435-1.432 1.435H4.562c-1.401 0-1.745-.593-1.432-1.435l2.786-9.652L5.905.941h5.125V3.93z"/><path d="M10.039 6.031h-3.02l-1.838 6.308c-.355 1.15-.24 1.584 1.408 1.584h3.879c1.633 0 1.67-.518 1.409-1.584z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TestTubeEmpty(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 1V.023H6V1h1.012v6L3 15s0 .962 1 .962h10c1 0 1-.962 1-.962l-4.042-8l-.02-6zm2 14.031H4L8 7V1h2v6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TestTubeTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M9.726.04L8.32 1.443l.703.703l-8.48 8.323c-.738.735-.674 1.996.143 2.811l2.031 2.025c.816.814 2.047.91 2.785.173l8.514-8.353l.684.684l1.408-1.404zm-6.37 14.358l-1.771-1.771c-.502-.501-.591-1.232-.195-1.623l8.35-8.225l3.59 3.592l-8.351 8.222c-.395.395-1.123.305-1.623-.195"/><path d="M5.273 12.78c-.34.336-.904.316-1.406-.184l-.376-.377c-.503-.504-.636-1.184-.298-1.518l3.858-3.858l3.203.909z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextSearch(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M0 0h13.931v.983H0zm0 2h13.931v.942H0zm0 12h10.958v.951H0zm8.49-9.946C6.01 4.054 4 6.047 4 8.506c0 2.459 2.01 4.452 4.49 4.452c2.48 0 4.489-1.993 4.489-4.452c0-2.459-2.008-4.452-4.489-4.452m0 7.964a3.54 3.54 0 1 1 0-7.08a3.54 3.54 0 1 1 0 7.08m7.448 2.593l-1.361 1.361l-2.996-2.996s.57-.073.931-.434c.361-.362.431-.928.431-.928z"/><path d="M8.677 6.43c.526 0 .329-.4-.403-.4a2.267 2.267 0 0 0-2.279 2.256c0 .725.404.921.404.4C6.398 7.44 7.418 6.43 8.677 6.43M0 4h3.973v.962H0zm0 2h3v.973H0zm0 2h2.98v.993H0zm0 2h3.02v.973H0zm0 2h4v.931H0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Thermometer(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.984 9.261V.735c0-.371-.371-.672-.83-.672h-1.31c-.458 0-.829.301-.829.672V9.26c-1.273.519-2.016 1.766-2.016 3.225a3.485 3.485 0 0 0 6.968 0c.002-1.459-.711-2.705-1.983-3.224M7.947.957H9.02V10H7.947z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThreeBall(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M4.518 7.071C2.046 7.071.04 9.066.04 11.529c0 2.461 2.006 4.456 4.478 4.456s4.476-1.995 4.476-4.456a4.468 4.468 0 0 0-4.476-4.458m.564 1.383a3.639 3.639 0 0 0-3.633 3.642c0 .822-.717.47-.717-.719a3.637 3.637 0 0 1 3.635-3.641c1.186 0 1.535.718.715.718"/><path d="M11.518 7.071a4.449 4.449 0 0 0-2.742.957c.131.157.25.322.361.493a3.6 3.6 0 0 1 2.23-.785c1.186 0 1.535.718.715.718c-.953 0-1.812.376-2.461.978c.268.646.418 1.354.418 2.098a5.462 5.462 0 0 1-1.264 3.499c.76.591 1.704.957 2.742.957c2.472 0 4.476-1.995 4.476-4.456a4.466 4.466 0 0 0-4.475-4.459m-3.5.126a5.57 5.57 0 0 1 4.678-1.104a4.414 4.414 0 0 0 .298-1.563c0-2.463-2.004-4.458-4.476-4.458S4.04 2.067 4.04 4.53c0 .51.104.992.263 1.448c.072-.003.143-.011.215-.011c1.326 0 2.542.462 3.5 1.23M8.367.736c1.186 0 1.535.718.715.718a3.639 3.639 0 0 0-3.633 3.642c0 .822-.717.47-.717-.719A3.637 3.637 0 0 1 8.367.736"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TicTacToe(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.492 9.019a3.47 3.47 0 1 0 .001 6.94a3.47 3.47 0 0 0-.001-6.94m.041 6.098a2.619 2.619 0 0 1-2.629-2.609a2.62 2.62 0 0 1 2.629-2.609c1.453 0 2.632 1.169 2.632 2.609c0 1.441-1.179 2.609-2.632 2.609M4.512.079a3.427 3.427 0 1 0 .002 6.854A3.427 3.427 0 0 0 4.512.079m0 6.016a2.589 2.589 0 0 1 0-5.176a2.59 2.59 0 0 1 2.589 2.587a2.591 2.591 0 0 1-2.589 2.589M8.006 9.97L7.1 9.063l-2.56 2.561l-2.561-2.561l-.907.907l2.561 2.561l-2.561 2.559l.907.906l2.561-2.558l2.56 2.558l.906-.906l-2.56-2.559zM17.057.926l-.912-.913l-2.577 2.577L10.992.013l-.911.913l2.576 2.576l-2.576 2.574l.911.912l2.576-2.575l2.577 2.575l.912-.912l-2.576-2.574z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimeGlass(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M12 2.026C12 .908 11.135 0 10.065 0h-3.12C5.877 0 5.01.908 5.01 2.026v3.536l2.021 2.427l-2.021 2.448v3.516c0 1.118.866 2.025 1.935 2.025h3.12c1.069 0 1.935-.907 1.935-2.025v-3.578L9.906 8.001L12 5.625zM11 5L9 7.989L11 11v3a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1v-3l2-3.011L6 5.02V2a1 1 0 0 1 1-1h3a1 1 0 0 1 1 1z"/><path d="M10 14H7v-2l1.5-1.979L10 12zm0-10L8.521 5.979L7 4z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TimeReload(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M9.549 1.046c-3.859 0-6.819 3.192-7.166 6.985H1.059l1.892 1.952l2.065-1.952H3.677c.331-3.229 2.747-5.958 5.937-5.958c3.412 0 6.189 2.888 6.189 6.437c0 3.549-2.777 6.438-6.189 6.438c-1.695 0-3.232-.713-4.35-1.865l-.821.826a7.364 7.364 0 0 0 5.106 2.065c4.092 0 7.419-3.349 7.419-7.464s-3.327-7.464-7.419-7.464"/><path d="M9 3.99V9h3.96V8H9.97V3.99z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Timer(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="m12.115 2.612l.525-.812l.85.554l.484-.75L11.527.011l-.483.75l.849.553l-.512.788A7.394 7.394 0 0 0 7.502.999C3.387.999.041 4.352.041 8.475c0 4.12 3.346 7.474 7.461 7.474c4.113 0 7.461-3.354 7.461-7.474a7.463 7.463 0 0 0-2.848-5.863M7.502 14.011c-3.047 0-5.527-2.488-5.527-5.544c0-3.058 2.48-5.545 5.527-5.545s5.527 2.487 5.527 5.545c0 3.055-2.48 5.544-5.527 5.544"/><path d="M7 4h1v5H7z"/><path d="M7 8h3v1H7z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Toilet(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M11.777 15.974h-6.29s3.013-5.98-1.474-5.98V8.001h11.966s.087 1.217-2.112 2.686c-3.387 2.26-2.09 5.287-2.09 5.287M8 6h7.979v.979H8zM3.012.009v6.974H7V.009zM6 3.036H5V1h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToolBox(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.988 9.031V8H5.031v1.031H2.938V8H1.014v6H17V8h-1.969v1.031zM5 5.986v.982h8v-.982h2v.982h1.987V4H1v2.968h1.974v-.982zm2.003-3.043H6.01V1.031h5.959v1.912h-.96v-.956H7.003z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrafficLight(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M9.671.017H7.296C6.578.017 6 .608 6 1.334v10.342c0 .727.578 1.317 1.296 1.317h2.375c.717 0 1.297-.591 1.297-1.317V1.334c0-.727-.58-1.317-1.297-1.317m-1.167 11.99a1.536 1.536 0 0 1-1.539-1.534c0-.846.69-1.535 1.539-1.535a1.535 1.535 0 1 1 0 3.069m.01-4.008a1.532 1.532 0 0 1-1.528-1.534a1.53 1.53 0 1 1 3.059 0c0 .847-.683 1.534-1.531 1.534m-.022-4.006c-.805 0-1.461-.668-1.461-1.491c0-.824.656-1.494 1.461-1.494c.812 0 1.464.67 1.464 1.494c0 .823-.652 1.491-1.464 1.491m4.149-.014h.524c.337 0 1.811-2.194 1.811-2.194c0-.418-.274-.76-.611-.76h-1.724c-.338 0-.611.342-.611.76v1.434c-.001.419.273.76.611.76m1.723 1.04h-1.726c-.336 0-.609.339-.609.756v1.423c0 .418.273.756.609.756h.524c.337 0 1.812-2.179 1.812-2.179c.001-.418-.275-.756-.61-.756m-.02 4.014h-1.725c-.336 0-.61.342-.61.761v1.433c0 .422.274.761.61.761h.525c.336 0 1.811-2.193 1.811-2.193c0-.42-.274-.762-.611-.762M4.355 1.024H2.64c-.334 0-.605.337-.605.749c0 0 1.463 2.163 1.799 2.163h.521c.335 0 .605-.337.605-.749V1.773c.001-.412-.27-.749-.605-.749m.02 4.015H2.66c-.334 0-.605.339-.605.756c0 0 1.463 2.178 1.799 2.178h.521c.334 0 .605-.338.605-.756V5.795c0-.417-.271-.756-.605-.756m.002 4.015H2.661c-.334 0-.606.332-.606.739c0 0 1.463 2.131 1.799 2.131h.523c.334 0 .604-.33.604-.739V9.793c0-.407-.27-.739-.604-.739M7 13h2.973v3.128H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Train(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M5.979 12.484a1.48 1.48 0 0 0-2.956 0l-.003.02a1.48 1.48 0 1 0 2.96 0zm5.019-.03a1.484 1.484 0 0 0-1.494-1.433c-.814 0-1.473.64-1.492 1.433a.197.197 0 0 0-.004.034c0 .81.67 1.467 1.496 1.467c.828 0 1.498-.657 1.498-1.467c0-.01-.004-.023-.004-.034m4 0a1.484 1.484 0 0 0-1.494-1.433c-.814 0-1.473.64-1.492 1.433a.197.197 0 0 0-.004.034c0 .81.67 1.467 1.496 1.467c.828 0 1.498-.657 1.498-1.467c0-.01-.004-.023-.004-.034"/><path d="M16.938 12.98c0-.965-.922-1.971-.922-1.971h-.047V8.015h-.984v-3l-1.962.024v3l-1.054-.024V6.036h-.953v1.979H7.969V4.953h1V3H1.016v1.938h1v5.125H1.03v2.922h.979a2.621 2.621 0 0 1-.048-.48c0-.012.005-.033.005-.033a2.54 2.54 0 0 1 2.535-2.513a2.54 2.54 0 0 1 2.537 2.513c0 .01.006.021.006.033c0 .164-.02.324-.049.48h8.975v-.004zM6 8.041H3.992V4.969H6z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrainRail(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M4.989 16h-.978l2-16h.978zm9 0h-.978l-2-16h.978z"/><path d="M3 13h12v.916H3zm0-3h12v.916H3zm1-3h10v.916H4zm0-3h10v.916H4zm1-3h8v.916H5z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trash(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m1.982 5.073l1.025 10.266c0 .366.307.661.684.661h7.58a.673.673 0 0 0 .684-.661L12.98 5.073zm6.051 8.995H6.961V6.989h1.072zm2 0H8.961l1-7.079h1.072zm-4 0H4.961l-1-7.079h1.072zm7.042-11.963H9.937V.709C9.937.317 9.481 0 9.081 0H5.986c-.4 0-.955.225-.955.615v1.396l-3.145.094a.717.717 0 0 0-.727.708v1.155H13.8V2.813a.715.715 0 0 0-.725-.708M5.947 1.44c0-.312.351-.565.783-.565h1.564c.432 0 .782.254.782.565v.665h-3.13z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tree(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m14.779 12.18l-2.984-3.679s1.601.436 1.775.436c.465 0 .195-.517 0-.714l-2.385-3.031s1.148-.274 1.565-.274c.418 0 .197-.517 0-.714L9.4.061a.505.505 0 0 0-.714 0l-3.395 4.1c-.198.197-.486.715 0 .715s1.622.316 1.622.316L4.325 8.079c-.198.197-.557.714 0 .714c.319 0 1.95-.291 1.95-.291l-2.958 3.687c-.197.196-.557.714 0 .714s4.691-1.007 4.691-1.007v3.045c0 .537.436.973.975.973c.537 0 1.015-.436 1.015-.973v-3.045s4.375.999 4.78.999c.405 0 .198-.519.001-.715"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trees(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.986 14.203v-2.278c1.52-.339 2.967-1.952 2.967-3.896c0-2.19-1.648-7.982-3.451-7.982c-1.803 0-3.455 5.792-3.455 7.982c0 1.901 1.513 3.489 2.984 3.874v2.21a62.349 62.349 0 0 0-3.011-.072a24.06 24.06 0 0 0-3.067.197v-2.672c1.058-.311 1.97-1.418 1.97-2.75c0-1.57-1.095-5.781-2.442-5.781c-1.35 0-2.442 4.211-2.442 5.781c0 1.354.905 2.48 1.993 2.769v2.793c-2.395.413-4.01 1.113-4.01 1.567h16c-.001-.456-1.627-1.531-4.036-1.742"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleDoubleArrowDown(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M10.044 16.565a1.49 1.49 0 0 1-2.104 0l-6.442-6.444c-.582-.581-.839-2.103 1-2.103h12.988c1.901 0 1.582 1.521 1.001 2.103z"/><path d="M10.044 9.565a1.49 1.49 0 0 1-2.104 0L1.498 3.121c-.582-.581-.839-2.103 1-2.103h12.988c1.901 0 1.582 1.521 1.001 2.103z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleDoubleArrowLeft(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M1.446 10.052a1.49 1.49 0 0 1 0-2.104L7.89 1.506c.581-.582 2.103-.839 2.103 1v12.988c0 1.901-1.521 1.582-2.103 1.001z"/><path d="M8.446 10.052a1.49 1.49 0 0 1 0-2.104l6.444-6.442c.581-.582 2.103-.839 2.103 1v12.988c0 1.901-1.521 1.582-2.103 1.001z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleDoubleArrowRight(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M9.113 15.495c-.582.581-2.103.9-2.103-1.001V1.506c0-1.839 1.521-1.582 2.103-1l6.444 6.442a1.49 1.49 0 0 1 0 2.104z"/><path d="M2.113 15.495c-.582.581-2.103.9-2.103-1.001V1.506c0-1.839 1.521-1.582 2.103-1l6.444 6.442a1.49 1.49 0 0 1 0 2.104z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleDoubleArrowUp(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M15.446 6.773c.581.581.9 2.103-1.001 2.103H1.457c-1.839 0-1.582-1.521-1-2.103L6.898.329a1.49 1.49 0 0 1 2.104 0z"/><path d="M15.446 13.773c.581.581.9 2.103-1.001 2.103H1.457c-1.839 0-1.582-1.521-1-2.103l6.441-6.444a1.49 1.49 0 0 1 2.104 0z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleDown(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.106 12.69a1.49 1.49 0 0 1-2.104 0L1.561 6.246c-.582-.581-.839-2.103 1-2.103h12.988c1.901 0 1.582 1.521 1 2.103z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleLeft(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M3.446 10.052a1.49 1.49 0 0 1 0-2.104L9.89 1.506c.581-.582 2.103-.839 2.103 1v12.988c0 1.901-1.521 1.582-2.103 1.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleRight(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M6.113 15.495c-.582.581-2.103.9-2.103-1.001V1.506c0-1.839 1.521-1.582 2.103-1l6.444 6.442a1.49 1.49 0 0 1 0 2.104z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TriangleUp(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M7.96 2.392a1.49 1.49 0 0 1 2.104 0l6.442 6.444c.582.581.839 2.103-1 2.103H2.518c-1.902 0-1.582-1.521-1.001-2.103z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrolleyArrowDown(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" fill-rule="evenodd"><path d="M3.17 9.994h1.873V7.959H2.861zM6 8h2.062v2.078H6zm-.969-3.029H2.52l.314 2.066h2.197zM12 7.109h1.746l.332-2.14H12zM9 8h2.062v2.094H9zm2.969 2.031h1.617l.289-2.047h-1.906z"/><path fill="currentColor" d="M4.039 13.484a1.437 1.437 0 1 0 2.768-.542H4.148c-.07.167-.109.351-.109.542m6.022.016c0 .821.65 1.485 1.453 1.485s1.451-.664 1.451-1.485c0-.197-.039-.385-.107-.557h-2.689c-.069.172-.108.36-.108.557m1.453-1.484c-.101 0-.199.011-.293.03h.586a1.42 1.42 0 0 0-.293-.03"/><path fill="currentColor" d="M11.221 12.046H3.564c-.319 0-.501.056-.501.423s.182.474.501.474h.583a1.435 1.435 0 0 1 2.659 0l3.362.001a1.46 1.46 0 0 1 1.053-.898"/><path fill="currentColor" d="M15.684 1.337c-.717 0-1.297.599-1.297 1.338l-.172 1.343h-3.172v.951h-.027v2.047H8.969v-.974h-.938v1.005H5.953V4.953h.005v-.936H1.557c-.318 0-.541.287-.541.654c0 .139.01.27.072.375L2.299 10.5c.04.172.168.299.331.369a.508.508 0 0 0 .265.083H13.36l-.174 1.093h-1.379c.48.101.871.443 1.051.897h1.496l1.16-8.948c.057.007.111.018.17.018c.715 0 1.295-.6 1.295-1.338c0-.738-.581-1.337-1.295-1.337M2.52 4.971h2.512v2.066H2.835zm2.523 5.023H3.17l-.309-2.035h2.182zm2.988.053H5.969V7.969h2.062zm3 0H8.969V7.953h2.062zm2.555-.016h-1.617V7.984h1.906zm.16-2.922H12V4.968h2.078zm-8.269 4.938c-.603 0-1.116.371-1.33.896h2.659a1.435 1.435 0 0 0-1.329-.896"/><path fill="currentColor" d="M11.807 12.046h-.586c-.48.101-.872.443-1.053.897h2.689a1.455 1.455 0 0 0-1.05-.897M9 3V1h-.979v2h-1l1.5 1.938L9.979 3z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrolleyArrowUp(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" fill-rule="evenodd"><path d="M2.17 9.994h1.873V7.959H1.861zM5 8h2.062v2.078H5zm-.969-3.029H1.52l.314 2.066h2.197zM11 7.109h1.746l.332-2.14H11zM8 8h2.062v2.094H8zm2.969 2.031h1.617l.289-2.047h-1.906z"/><path fill="currentColor" d="M3.039 13.484a1.437 1.437 0 1 0 2.768-.542H3.148c-.07.167-.109.351-.109.542m6.022.016c0 .821.65 1.485 1.453 1.485s1.451-.664 1.451-1.485c0-.197-.039-.385-.107-.557H9.169c-.069.172-.108.36-.108.557m1.453-1.484c-.101 0-.199.011-.293.03h.586a1.42 1.42 0 0 0-.293-.03"/><path fill="currentColor" d="M10.221 12.046H2.564c-.319 0-.501.056-.501.423s.182.474.501.474h.583a1.435 1.435 0 0 1 2.659 0l3.362.001a1.46 1.46 0 0 1 1.053-.898"/><path fill="currentColor" d="M14.684 1.337c-.717 0-1.297.599-1.297 1.338l-.172 1.343h-3.172v.951h-.027v2.047H7.969v-.974h-.938v1.005H4.953V4.953h.005v-.936H.557c-.318 0-.541.287-.541.654c0 .139.01.27.072.375L1.299 10.5c.04.172.168.299.331.369a.508.508 0 0 0 .265.083H12.36l-.174 1.093h-1.379c.48.101.871.443 1.051.897h1.496l1.16-8.948c.057.007.111.018.17.018c.715 0 1.295-.6 1.295-1.338c0-.738-.581-1.337-1.295-1.337M1.52 4.971h2.512v2.066H1.835zm2.523 5.023H2.17l-.309-2.035h2.182zm2.988.053H4.969V7.969h2.062zm3 0H7.969V7.953h2.062zm2.555-.016h-1.617V7.984h1.906zm.16-2.922H11V4.968h2.078zm-8.269 4.938c-.603 0-1.116.371-1.33.896h2.659a1.435 1.435 0 0 0-1.329-.896"/><path fill="currentColor" d="M10.807 12.046h-.586c-.48.101-.872.443-1.053.897h2.689a1.455 1.455 0 0 0-1.05-.897M7 2.938v2h.98v-2h1L7.48 1L6.021 2.938z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrolleyBriefcase(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" transform="translate(1)"><ellipse cx="4.471" cy="14.478" rx="1.471" ry="1.478"/><circle cx="12.481" cy="14.481" r="1.481"/><path d="M15.342 11.062H2.938V3.075L1.206.267A.68.68 0 0 0 .281.126a.63.63 0 0 0-.146.896L2.033 3.53v7.532s-.971-.167-.971.417c0 .521.565.469.934.469h13.346c.368 0 .609-.064.609-.413s-.241-.473-.609-.473M6.223 7.5v-.041z"/><path d="M6.889 9.297a.632.632 0 0 1-.625.639H4.681a.631.631 0 0 1-.624-.639V2.663c0-.353.279-.639.624-.639h1.583c.346 0 .625.286.625.639zm4.039-.025c0 .366-.282.663-.629.663H8.705c-.349 0-.629-.297-.629-.663V3.77c0-.367.28-.664.629-.664h1.594c.347 0 .629.297.629.664zm4.018-.013c0 .396-.285.719-.638.719h-1.617c-.354 0-.639-.322-.639-.719V4.801c0-.396.285-.717.639-.717h1.617c.353 0 .638.321.638.717z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrolleyError(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" fill-rule="evenodd"><g fill="currentColor" transform="translate(3 12)"><circle cx="1.437" cy="1.437" r="1.437"/><ellipse cx="7.452" cy="1.485" rx="1.452" ry="1.485"/></g><path d="M5 8h2.062v2.078H5zm0-3h2.078v2.094H5zm3 0h2.047v2.047H8zm5.078-.031H11v2.14h1.746zM8 8h2.062v2.094H8z"/><path fill="currentColor" d="M2.563 12.046c-.319 0-.501.056-.501.423s.182.474.501.474l9.416.001v-.897zm8.406-2.015V7.984h1.906l-.281 1.995h1.145l.775-5.984c.057.007.111.018.17.018c.715 0 1.295-.6 1.295-1.338c0-.739-.58-1.338-1.295-1.338c-.717 0-1.297.599-1.297 1.338l-.172 1.343H.557c-.318 0-.541.287-.541.654c0 .139.01.27.072.375l1.211 5.454c.04.172.168.299.331.369a.508.508 0 0 0 .265.083H11.98v-.922zM11 4.969h2.078l-.332 2.141H11zm-3.031 0h2.047v2.047H7.969zm-6.449.002h2.512v2.066H1.835zm2.523 5.023H2.17l-.309-2.035h2.182zm2.988.053H4.969V7.969h2.062zm0-3H4.953V4.953h2.078zm3 3H7.969V7.953h2.062zm5.963 1.664l-.701-.69l-.788.792l-.796-.792l-.691.702l.793.788l-.791.796l.701.689l.787-.791l.796.791l.69-.701l-.792-.788z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrolleyFull(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" transform="translate(0 2)"><path d="M5.8 9L3.451 0H.002l.025.938h2.438L5.065 10H15.97V9z"/><path d="M7 6h1.958v1.955H7zm3 0h1.953v2H10zm0-3h1.938v1.899H10zm0-3h1.953v1.941H10zm3 6h1.953v1.906H13zm0-3h1.938v1.899H13zm0-3h1.953v1.941H13zM7 3h1.938v1.899H7zm0-3h1.953v1.941H7z"/><ellipse cx="7.494" cy="10.471" rx="1.494" ry="1.471"/><circle cx="13.487" cy="10.487" r="1.487"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrolleyPlus(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" fill-rule="evenodd"><g fill="currentColor" transform="translate(4 12)"><circle cx="1.437" cy="1.437" r="1.437"/><ellipse cx="7.452" cy="1.485" rx="1.452" ry="1.485"/></g><path d="M6 8h2.062v2.078H6zm-.969-3.029H2.52l.314 2.066h2.197zM6 5h2.078v2.094H6zm3 0h2.047v2.047H9zm5.078-.031H12v2.14h1.746zM3.17 9.994h1.873V7.959H2.861zM9 8h2.062v2.094H9z"/><path fill="currentColor" d="M3.563 12.046c-.319 0-.501.056-.501.423s.182.474.501.474l9.416.001v-.897zm8.406-2.015V7.984h1.906l-.281 1.995h1.145l.775-5.984c.057.007.111.018.17.018c.715 0 1.295-.6 1.295-1.338c0-.739-.58-1.338-1.295-1.338c-.717 0-1.297.599-1.297 1.338l-.172 1.343H1.557c-.318 0-.541.287-.541.654c0 .139.01.27.072.375l1.211 5.454c.04.172.168.299.331.369a.508.508 0 0 0 .265.083H12.98v-.922zM12 4.969h2.078l-.332 2.141H12zm-3.031 0h2.047v2.047H8.969zm-6.449.002h2.512v2.066H2.835zm2.523 5.023H3.17l-.309-2.035h2.182zm2.988.053H5.969V7.969h2.062zm0-3H5.953V4.953h2.078zm3 3H8.969V7.953h2.062zM17 12h-1.062v-1H15v1h-1v.958h1v.917h.938v-.917H17z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrolleyRemove(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="none" fill-rule="evenodd"><g fill="currentColor" transform="translate(4 12)"><circle cx="1.437" cy="1.437" r="1.437"/><ellipse cx="7.452" cy="1.485" rx="1.452" ry="1.485"/></g><path fill="currentColor" d="M14 11h3v.958h-3z"/><path d="M6 8h2.062v2.078H6zm-.969-3.029H2.52l.314 2.066h2.197zM6 5h2.078v2.094H6zm3 0h2.047v2.047H9zm5.078-.031H12v2.14h1.746zM3.17 9.994h1.873V7.959H2.861zM9 8h2.062v2.094H9z"/><path fill="currentColor" d="M3.563 12.046c-.319 0-.501.056-.501.423s.182.474.501.474l9.416.001v-.897zm8.406-2.015V7.984h1.906l-.281 1.995h1.145l.775-5.984c.057.007.111.018.17.018c.715 0 1.295-.6 1.295-1.338c0-.739-.58-1.338-1.295-1.338c-.717 0-1.297.599-1.297 1.338l-.172 1.343H1.557c-.318 0-.541.287-.541.654c0 .139.01.27.072.375l1.211 5.454c.04.172.168.299.331.369a.508.508 0 0 0 .265.083H12.98v-.922zM12 4.969h2.078l-.332 2.141H12zm-3.031 0h2.047v2.047H8.969zm-6.449.002h2.512v2.066H2.835zm2.523 5.023H3.17l-.309-2.035h2.182zm2.988.053H5.969V7.969h2.062zm0-3H5.953V4.953h2.078zm3 3H8.969V7.953h2.062z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrolleyTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" transform="translate(1 1)"><circle cx="4.437" cy="12.437" r="1.437"/><ellipse cx="10.452" cy="12.485" rx="1.452" ry="1.485"/><path d="M14.684.337c-.717 0-1.297.599-1.297 1.338l-.172 1.343H.557c-.318 0-.541.287-.541.654c0 .139.01.27.072.375l1.211 5.454c.04.172.168.299.331.369a.508.508 0 0 0 .265.083H12.36l-.174 1.093H2.564c-.319 0-.501.056-.501.423s.182.474.501.474l10.79.001l1.16-8.948c.057.007.111.018.17.018c.715 0 1.295-.6 1.295-1.338c0-.74-.581-1.339-1.295-1.339m-1.938 5.772H11V3.968h2.078zM2.17 8.994l-.309-2.035h2.182v2.035zm5.799-5.025h2.047v2.047H7.969zm-.938 2.078H4.953V3.953h2.078zm-3-.01H1.834L1.52 3.971h2.512v2.066zm.938.932h2.062v2.078H4.969zm3-.016h2.062v2.094H7.969zm3 2.078V6.984h1.906l-.289 2.047z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Truck(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><g transform="translate(4 3)"><ellipse cx="1.438" cy="8.491" rx="1.438" ry="1.491"/><ellipse cx="9.463" cy="8.48" rx="1.463" ry="1.48"/><path d="M4.031.062v6.885h3.464c.47-.61 1.2-1.01 2.03-1.01c.831 0 1.562.399 2.031 1.01h1.398V.062z"/></g><path d="M8 11h2.916v.875H8zm8 0h.979v1H16zm-13.064.521a2.585 2.585 0 0 1 4.042-2.132V5.98c0-.527-.403-.954-.901-.954H2.946c-.499 0-1.056.846-1.056.846s-.841 1.505-.841 2.145v3.013c0 .542.396.943.901.943h1.032a2.614 2.614 0 0 1-.046-.452m.063-5.592l3.057.021l-.022 2.112H2.253c-.068-1.276.746-2.133.746-2.133"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trumpet(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m14.622 3.308l1.076-1.076l-.91-.91l-2.431 2.431l-.699-.698l-.658.66l.697.697l-1.305 1.305l-.667-.667l-.721.722l.666.667l-1.302 1.302l-.603-.603l-.691.691l.603.603L6.139 9.97c-1.369 1.369-4.321 2.119-5.1 1.34c0 0 .396 1.366 1.836 2.805c1.438 1.439 2.834 1.865 2.834 1.865c-.779-.778-.008-3.752 1.34-5.099l1.27-1.27l1.037 1.036c.231.231.54.342.875.342c.483 0 1.021-.231 1.463-.672L15.33 6.68c.427-.426.671-.959.671-1.463c0-.349-.119-.651-.341-.874zm.135 2.801l-3.636 3.637c-.45.447-.991.532-1.193.33L8.891 9.04l5.159-5.16l1.037 1.036c.086.086.104.212.104.302c-.001.282-.167.624-.434.891m2.134-4.807L15.703.114c-.092-.093-.299-.037-.461.126c-.162.163-.219.369-.127.462l1.189 1.188c.094.092.301.036.461-.126c.165-.164.219-.369.126-.462"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TurnOff(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M11.124 1.613v2.349c1.606.887 2.632 2.624 2.632 4.578c0 2.854-2.35 5.169-5.248 5.169c-2.896 0-5.246-2.314-5.246-5.169c0-2.017 1.012-3.749 2.696-4.601V1.611c-2.937.975-4.895 3.693-4.895 6.929c0 4.052 3.334 7.335 7.444 7.335c4.111 0 7.446-3.283 7.446-7.335c.001-3.177-1.973-5.902-4.829-6.927"/><path d="M8.46 7.873c.778 0 1.412-.48 1.412-1.075V1.115C9.872.521 9.238.04 8.46.04c-.779 0-1.412.481-1.412 1.075v5.683c0 .595.633 1.075 1.412 1.075"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwoArrowDown(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.002 11.854L3.008 8.109v3.94l4.994 3.943l4.965-4.042V8.032l-.012-.011zm0-7.784L3.008.119v3.939l4.994 3.965l4.965-4.064V.041l-.012-.01z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwoArrowInDownUp(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.771 4.096a.664.664 0 0 0 0-.946a.678.678 0 0 0-.954 0L9.989 4.503V1c0-.553-.444-1-.992-1a.995.995 0 0 0-.992 1v3.564L6.187 3.18a.674.674 0 0 0-1.15.473c0 .171.066.343.199.474l3.74 2.819zm0 7.78a.664.664 0 0 1 0 .946a.678.678 0 0 1-.954 0l-1.828-1.328V15c0 .553-.444 1-.992 1a.995.995 0 0 1-.992-1v-3.428l-1.818 1.266a.674.674 0 0 1-1.15-.473c0-.171.066-.343.199-.474l3.74-2.838z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwoArrowInLeftRight(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.125 5.229a.664.664 0 0 0-.946 0a.679.679 0 0 0 0 .954l1.353 1.828H2c-.553 0-1 .444-1 .992s.447.992 1 .992h3.594L4.21 11.812a.673.673 0 0 0 .472 1.15a.667.667 0 0 0 .475-.198l2.819-3.74zm7.696 0a.665.665 0 0 1 .947 0a.677.677 0 0 1 0 .954L12.44 8.011h3.531c.553 0 1 .444 1 .992a.995.995 0 0 1-1 .992h-3.453l1.266 1.817a.675.675 0 0 1-.473 1.15a.665.665 0 0 1-.474-.198l-2.838-3.74z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwoArrowLeft(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M5.146 8.027L8.89 3.032H4.95L1.007 8.027l4.042 4.964h3.918l.01-.011zm7.783 0l3.951-4.995h-3.939L8.976 8.027l4.064 4.964h3.918l.01-.011z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwoArrowRight(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.83 7.999L9.086 13h3.939l3.944-5.001l-4.042-4.969H9.009l-.011.011zm-7.783 0L1.096 13h3.938L9 7.999L4.935 3.03H1.018l-.01.011z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Typewriter(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M7 10h.938v.953H7zm2 0h.938v.953H9zm2 0h.938v.953H11z"/><path d="M2.062 14.98c0 .562.45 1.02 1.006 1.02h12.863c.556 0 1.006-.457 1.006-1.02l-1.938-6.021V6h-1.976c0 .011.005.021.005.031c0 .998-1.534.818-3.071.969h-.976c-1.528-.15-3.025.029-3.025-.969c0-.011.005-.021.005-.031h-1.94v2H1v2.938h1V9h2.009zm2.907-4.011h1v-.953h-1l.016-1.047h1.047v1h.938v-1h1.062v1h.938v-1h1.062v1h.938v-1h1.062v1h.938v-1h1.062v1.047h-1v.953h1v1.047h-1.047v-1h-.956v1h-1.051v-1h-.965v.953h.007v.047H8.979v-1h-.972v1H6.971v-1H6v1H4.969zM4.979 13h9.041v1H4.979zM16.021 3v1.049H14.94l-.002-3.027l-10.913.041v2.986H2.96V3h-.938v3h.938V4.951h13.062V6h.97V3z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UmberllaChair(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M9 0L1 6v1.984h.953v-1H4v1h1.969v-1h2.047v1h1v5.985h.948V7.984h.005v-1h2.078v1H14v-1h2.016v1H17V6z"/><path d="M3.229 11.016H1V12h1.678l3.342 2.994l1.011-.01v.985h.938v-.985h1.062v.985H10v-.985h4.047v.985H15v-.985h.984V14H6.562z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UmbrellaClose(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.031 10.938v1.045h.938L7.907 1.968V.43a.411.411 0 0 0-.824 0v1.538L3.985 11.983h.953v-1.045h1.078v1.045h1.016v3.549c0 .228.251.412.479.412c.228 0 .459-.185.459-.412v-3.549h.969v-1.045z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UmbrellaOpen(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.98 7.984c0-3.763-3.16-6.659-7.012-6.953V.468c0-.258-.17-.467-.447-.467a.475.475 0 0 0-.49.467v.555c-3.941.205-7.018 3.139-7.018 6.962h.955V6.969H4.03v1.016h1.938V6.969h1.08v1.016h1.904V6.969h1.094l-.016 1.016h1.949V6.969h1.02v1.016h1.938V6.969h1.078v1.016h.965zm-7.96 6.218c0 .479-.507.816-1.003.816c-.316 0-1.035-.016-1.035-.494a.474.474 0 0 0-.48-.466a.473.473 0 0 0-.478.466c0 .99.97 1.473 1.993 1.473c1.025 0 1.978-.803 1.978-1.795v-6.17h-.974z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UmbrellaSea(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M8.469 1L1 7v1.984h.953v-1H4v1h1.969v-1h2.047v7.047H5.002v1H3v.907h10.938v-.907H12v-1H8.964V8.984h.005v-1h2.078v1H13v-1h2.016v1H16V7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Upstair(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16.969 1.047h-2.896v2h-2.011v2h-2.01v2H8.043v2H6.031v2h-2.01v2H1v2.918h2.896L17 3.965zm-9.096.057H2.949l2.155 1.414l-3.182 3.18l1.326 1.326l3.181-3.181l1.448 2.205z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpwardsArrowToBar(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M15.002 4.913c.551 0 .995-.42.995-.938V2.102c0-.518-.444-.938-.995-.938H1.076c-.55 0-.995.42-.995.938v1.873c0 .519.445.938.995.938zM6.994 6.518a1.488 1.488 0 0 1 2.092 0l6.408 6.374c.579.573.835 2.078-.995 2.078H1.579c-1.891 0-1.573-1.505-.995-2.078z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpwardsArrowWithLoop(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M1 7h2.941v1.943H1z"/><path d="M14.02 7.049H8.017v1.909h6.052c.545 0 .988.44.988.979v3.129a.985.985 0 0 1-.988.979H7.958a.985.985 0 0 1-.988-.979V3.968h1.75c.392-.393.392-.638 0-1.029L6.708.291a1.003 1.003 0 0 0-1.42 0l-1.99 2.648c-.392.392-.392.638 0 1.029h1.729v9.053a2.977 2.977 0 0 0 2.98 2.967h6.014a2.976 2.976 0 0 0 2.98-2.967v-3.006a2.979 2.979 0 0 0-2.981-2.966"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Usb(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M13.074 3.098v1.84h.968V6.46L9.907 7.98V2.93L12 2.921L9.542 0L7.031 2.959h2.058v7.566l-4.13-1.099V7.92A1.495 1.495 0 0 0 4.5 4.999a1.495 1.495 0 0 0-.465 2.919V10l5.054 1.585v1.479A1.496 1.496 0 0 0 9.5 16a1.5 1.5 0 0 0 1.5-1.5a1.5 1.5 0 0 0-1.093-1.438V9.01l5.04-1.949V4.937h.97v-1.84z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func View(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" transform="translate(1 4)"><path d="M8 0C3.598 0 .031 2.66.031 3.969C.031 5.278 3.597 7.938 8 7.938c4.4 0 7.969-2.618 7.969-3.969S12.4 0 8 0m-.01 7.062C4.342 7.062 2.869 5.011 2.869 4C2.869 2.989 4.342.938 7.99.938c3.646 0 5.119 2.02 5.119 3.062c0 1.042-1.472 3.062-5.119 3.062"/><ellipse cx="7.932" cy="3.963" rx="1.932" ry="1.963"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WacomTablet(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M0 5v9h16V5zm5 2v5h9V7zM.959 7H2v1H.959zM2 12H.959v-1H2zM.953 10V8.979h2V10zM15 13H4V6h11z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wall(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12 3h4v2h-4zm-2 3h3v1.957h-3zm4 0h2v1.969h-2zM0 9h2v2H0zm3 0h4v2H3zm5 0h3v2H8zm4 0h4v2h-4zM5 6h4v1.957H5zM0 6h4v1.957H0zm10 6h3v1.957h-3zm4 0h2v1.969h-2zm-9 0h4v1.957H5zm-5 0h4v1.957H0zm8-9h3v2H8zM3 3h4v2H3zM0 3h2v2H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wallet(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.988 10.635V8.327c0-.749.573-1.358 1.279-1.358h4.697V5.531c0-.682-.404-1.252-.957-1.438v.001l-1.651-.014l-.02-.058H4.925l-.141.048l-2.614-.002l-.005-.021c-.636.12-1.148.696-1.148 1.437v7.953c0 .832.648 1.555 1.391 1.555h11.214c.743 0 1.343-.676 1.343-1.508v-1.453h-4.697c-.706 0-1.28-.647-1.28-1.396"/><path d="M15.996 8.061h-5.007c-.529 0-.958.468-.958 1.045v.816c0 .576.429 1.044.958 1.044h5.007c.529 0 .959-.468.959-1.044v-.816c0-.578-.43-1.045-.959-1.045m-2.965 1.955h-2.062V9h2.062zM4.926 4.022l6.631-2.269l.78 2.269h1.285c.135 0 .263.029.386.07L12.619.051L2.049 3.572l.117.476c.079-.015.16-.025.242-.025z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WashMachine(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M16.031.042h-15v3h15zm-12.989 2H1.938V.958h1.104zm2 0H3.938V.958h1.104z"/><path d="M1.021 3v13h14.958V3zm7.491 12.083c-3.068 0-5.555-2.5-5.555-5.583c0-3.084 2.486-5.584 5.555-5.584c3.067 0 5.554 2.5 5.554 5.584c-.001 3.083-2.487 5.583-5.554 5.583"/><path d="M8.516 5.083c-2.453 0-4.441 1.979-4.441 4.416c0 2.438 1.988 4.418 4.441 4.418c2.454 0 4.442-1.979 4.442-4.418c0-2.437-1.988-4.416-4.442-4.416m3.283 4.945c.079 0 .15-.021.218-.048a3.527 3.527 0 0 1-3.501 3.037a3.528 3.528 0 0 1-3.537-3.519c0-1.94 1.584-3.516 3.537-3.516a3.527 3.527 0 0 1 3.501 3.036a.53.53 0 1 0-.218 1.01"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WashMachineTwo(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M16.031.042h-15v3h15zM3 1.958H2v-1h1zm2 0H4v-1h1z"/><path d="M1.021 3v13h14.958V3zm7.436 12a5.5 5.5 0 1 1 0-11a5.5 5.5 0 0 1 0 11"/><path d="M8.516 5.083c-2.453 0-4.441 1.979-4.441 4.416c0 2.438 1.988 4.418 4.441 4.418c2.454 0 4.442-1.979 4.442-4.418c0-2.437-1.988-4.416-4.442-4.416m3.283 4.945c.079 0 .15-.021.218-.048c-.235 1.716-2.19 1.256-3.501.037c-1.531-1.424-3.537 1.425-3.537-.519c0-1.94 1.584-3.516 3.537-3.516a3.527 3.527 0 0 1 3.501 3.036a.53.53 0 1 0-.218 1.01"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Watch(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M8.969 6h-.953v2.941h.07v.043H10v-.953H8.969z"/><path d="M12.936 8.016a4.443 4.443 0 0 0-1.133-2.521A4.643 4.643 0 0 0 11 4.812V1H6v3.811a4.605 4.605 0 0 0-.859.76a4.41 4.41 0 0 0-1.117 2.92c0 .945.3 1.818.804 2.542c.315.453.715.843 1.173 1.154v3.767h5v-3.769c.415-.282.78-.631 1.081-1.032a4.458 4.458 0 0 0 .858-2.152h1.03v-.984zm-.936.62a3.47 3.47 0 0 1-.695 1.95A3.493 3.493 0 0 1 8.507 12a3.5 3.5 0 0 1-2.877-1.508A3.473 3.473 0 0 1 5 8.5c0-.878.337-1.673.875-2.287A3.482 3.482 0 0 1 8.507 5c1.025 0 1.94.45 2.58 1.153c.544.596.883 1.372.913 2.23v.254z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Webcam(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><g transform="translate(2 1)"><ellipse cx="5.967" cy="5.979" rx="1.967" ry="1.979"/><path d="M7.598.001c.002.035.014.068.014.105C7.61.788 6.887 1.34 5.994 1.34c-.892 0-1.615-.555-1.615-1.238c0-.035.01-.067.014-.102C1.864.714.003 3.072.001 5.88C-.003 9.259 2.678 11.999 5.985 12c3.31.003 5.994-2.732 5.996-6.108c.001-2.81-1.856-5.172-4.383-5.891M2 5.997a4 4 0 1 1 8 .008a4 4 0 1 1-8-.008"/></g><path d="M7.958 14.25a7.281 7.281 0 0 1-3.917-1.142l-1.824 1.811a.73.73 0 0 0 0 1.039l11.531-.052a.733.733 0 0 0 .002-1.039l-1.807-1.8a7.297 7.297 0 0 1-3.985 1.183"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WeightDown(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.905 5.01c0-1.088-.912-1.971-2.038-1.971h-.923c.034-.152.056-.32.056-.508C10 1.135 8.877 0 7.495 0S4.99 1.135 4.99 2.531c0 .188.026.355.068.508h-.922c-1.125 0-2.037.883-2.037 1.971L.083 13.985c0 1.09.912 1.972 2.039 1.972H12.88c1.126 0 2.037-.882 2.037-1.972zM5.824 2.531c0-.949.749-1.723 1.671-1.723c.921 0 1.67.773 1.67 1.723c0 .178-.034.346-.083.508H5.906a1.743 1.743 0 0 1-.082-.508m1.139 7.416V6.958h1.074v2.968l2.07.021l-2.648 3.092l-2.49-3.092z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WeightKilograms(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M9 9h.905v.92H9z"/><path d="M12.017 4.424c0-.732-.597-1.407-1.335-1.407h-.769A2.493 2.493 0 0 0 7.486 0a2.492 2.492 0 0 0-2.427 3.017h-.744c-.735 0-1.332.675-1.332 1.407l-1.975 9.25A1.33 1.33 0 0 0 2.344 15h10.313a1.33 1.33 0 0 0 1.333-1.326zM6 2.508C6 1.676 6.673 1 7.5 1C8.326 1 9 1.676 9 2.508A1.5 1.5 0 0 1 8.912 3H6.089A1.473 1.473 0 0 1 6 2.508m1.047 6.497H6.045v.948h1.002v1.067H5.982v-1.005h-.939v1.005H3.962V5.911h1.081v3.031h.939V7.937h1.065zm4 4.018h-1v1.023H8.985v-1.008l.064-.063l.892-.017l.012-1.937H7.986l-.008-3.069l3.07.013z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WeightUp(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M12.867 3.039h-.923c.034-.152.056-.32.056-.508C12 1.135 10.877 0 9.495 0S6.99 1.135 6.99 2.531c0 .188.026.355.068.508h-.922c-1.125 0-2.037.883-2.037 1.971l-2.016 8.975c0 1.09.79 1.972 1.917 1.972h11c1.126 0 1.917-.882 1.917-1.972L14.905 5.01c0-1.088-.912-1.971-2.038-1.971m-5.043-.508c0-.949.749-1.723 1.671-1.723c.921 0 1.67.773 1.67 1.723c0 .178-.034.346-.083.508H7.906a1.743 1.743 0 0 1-.082-.508m2.23 7.549v2.959H8.946v-2.938l-2.133-.021l2.73-3.06l2.564 3.06z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WheelChair(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" transform="translate(2 1)"><ellipse cx="6.49" cy="1.471" rx="1.49" ry="1.471"/><path d="M12.6 14.951a.98.98 0 0 1-.819-.442l-1.272-2.511H5.993a.993.993 0 0 1-.987-.999V5.064c0-.553.442-.998.987-.998c.543 0 .986.445.986.998v4.937h4.056c.329 0 .636.165.817.442l1.084 2.225l.654-.268a.98.98 0 0 1 1.323.449a1.004 1.004 0 0 1-.444 1.338l-1.431.659a.957.957 0 0 1-.438.105"/><path d="M5.381 15.974c-2.936 0-5.324-2.384-5.324-5.313c0-1.726.845-3.35 2.26-4.346a.667.667 0 0 1 .931.162a.666.666 0 0 1-.161.93a3.983 3.983 0 0 0-1.692 3.254a3.986 3.986 0 0 0 3.987 3.977c1.1 0 2.123-.437 2.885-1.23a.668.668 0 1 1 .964.925a5.292 5.292 0 0 1-3.85 1.641"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WheelSteel(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M7.999 16C3.589 16 0 12.41 0 8s3.588-8 7.999-8C12.411 0 16 3.59 16 8s-3.59 8-8.001 8M8 2C4.69 2 2 4.692 2 8s2.692 6 6 6s6-2.692 6-6s-2.69-6-6-6"/><path d="M7.992 6c-2.316 0-4.098.797-4.961 2.346c.037.473.154.928.336 1.352c1.221-.652 3.551 1.83 2.818 2.935c.58.204 1.141.355 1.797.355a5.94 5.94 0 0 0 1.887-.315c-.734-1.105 1.525-3.535 2.754-2.869a4.31 4.31 0 0 0 .344-1.396C12.095 6.875 10.293 6 7.992 6m.01 3.156c-.625 0-1.127-.51-1.127-1.141s.502-1.141 1.127-1.141c.619 0 1.123.51 1.123 1.141s-.504 1.141-1.123 1.141"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wieght(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M16 6.079V5h-1.955v3H4V5.041H2v1.011h-.961v5.906H2v.997h2V10h10.045v2.996H16v-1.038h1V6.079z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiOne(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" transform="translate(0 2)"><circle cx="7.989" cy="5.989" r="1.989"/><path d="M3.574 11.996C1.349 10.691.028 8.486.028 6.061c0-2.472 1.421-4.769 3.721-6.039l.428.73C2.168 1.876.93 3.894.93 6.061c0 2.128 1.151 4.063 3.09 5.215zm8.094-.717c2.095-1.079 3.387-3.092 3.387-5.298c0-2.162-1.254-4.155-3.284-5.249l.436-.695c2.308 1.236 3.734 3.498 3.734 5.944c0 2.499-1.472 4.784-3.854 6.005z"/><path d="M5.342 9.975C3.908 9.114 3.027 7.612 3.027 6.01c0-1.627.898-3.139 2.36-3.994l.434.76C4.645 3.477 3.923 4.7 3.923 6.01c0 1.296.713 2.514 1.872 3.218zm4.662-.731c1.277-.675 2.062-1.909 2.062-3.27c0-1.302-.734-2.511-1.93-3.2l.447-.728c1.465.841 2.366 2.327 2.366 3.928c0 1.664-.958 3.184-2.515 4.01z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WindTurbines(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M10.938 15.299c0 .352-1.058.639-1.442.639c-.385 0-1.442-.287-1.442-.639l.744-7.592c0-.353.313-.639.698-.639c.385 0 .698.286.698.639zM9.974 3.498c0 .28-.21.506-.468.506c-.26 0-.469-.226-.469-.506V.568c0-.279.209-.506.469-.506c.258 0 .468.227.468.506zm-5.13 5.441c-.265.183-.602.152-.756-.067c-.15-.219-.061-.547.203-.729l2.771-1.92c.264-.182.603-.153.755.066c.152.22.062.547-.203.73zm9.868-.99a.57.57 0 0 1 .188.783a.57.57 0 0 1-.791.154l-2.836-1.823a.57.57 0 0 1-.188-.784a.57.57 0 0 1 .789-.153zM9 5h.906v.938H9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Window(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m8.885 7.826l.896-6.236a8.133 8.133 0 0 1-.838-.439C6.855-.132 4.673-.085 3.052.257l-.984 6.566c1.524-.345 4.208-.586 6.817 1.003m7.292-.138l.763-6.141c-1.033.472-3.404 1.293-5.962.406l-.887 6.183c2.474 1.334 4.938.237 6.086-.448M8.674 9.199c-2.565-1.586-5.317-1.243-6.702-.913l-.907 5.735c1.159-.484 4.051-.753 6.935.487l.718-5.284zm6.564 6.538l.803-6.436c-.854.432-2.109.902-3.527.902c-.812 0-1.681-.16-2.553-.571l-.701 5.209c.066.039.135.066.204.107c2.12 1.257 4.257 1.15 5.774.789"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WoodStove(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M11.965 3.027V.072h-.949v2.955H3v11.957h2.016L5 15.978h.969v-.994h6.047v.994h.969v-.994H15V3.027zm1.05 7.02V12c0 1.002-.074 1.055-1.016 1.055H6.044c-.94 0-1.045-.053-1.045-1.055V6c0-1.001.104-1.033 1.045-1.033H12c.941 0 1.016.032 1.016 1.033v2.953h1.016v1.094z"/><path d="M11.974 7.031c0-.539-.401-.975-.896-.975H6.981c-.494 0-.897.436-.897.975v3.979c0 .537.403.975.897.975h4.097c.495 0 .896-.438.896-.975zM8 10H7V8h1zm2 0H8.964V8H10z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wrench(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="m7.982 11.19l4.177-4.185a3.604 3.604 0 0 0 3.795-.832c.698-.701 1.027-1.534 1.084-2.506l-.677.694l-2.212.74l-2.269-2.257l.776-2.174l.672-.67c-.946.078-1.799.36-2.487 1.048a3.623 3.623 0 0 0-.84 3.772L5.852 8.978c-1.266-.456-2.737-.117-3.765.914a3.596 3.596 0 0 0 0 5.078c1.387 1.39 3.654 1.37 5.066-.045c1.002-1.003 1.231-2.491.83-3.736zm-5.25 3.1a2.5 2.5 0 1 1 3.535-3.535a2.5 2.5 0 0 1-3.534 3.535z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WrenchScrewdriver(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="m5.811 11.799l5.852-5.852c1.162.601 2.544.489 3.432-.398c.684-.685.915-1.666.709-2.613l-1.509 1.506l-1.344.262l-1.626-1.609l.28-1.385L13.1.235c-.947-.207-1.93.022-2.613.706c-.889.888-.999 2.27-.398 3.433l-5.852 5.852c-1.163-.6-2.545-.488-3.433.398c-.684.686-.915 1.666-.708 2.613l1.508-1.505l1.343-.263l1.628 1.609l-.281 1.385l-1.495 1.476c.946.206 1.929-.022 2.613-.707c.887-.888.998-2.27.399-3.433"/><path d="m14.107 12.334l-1.83-.598L9.941 9.4l-.58.58l2.336 2.334l.637 1.872l3.067 1.833l.579-.579zm-6.83-7.227L2.478.29C2.075-.114 1.367-.064.9.408L.367.94c-.467.472-.52 1.179-.115 1.585l4.797 4.817z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Yen(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" fill-rule="evenodd" d="M14.722 1.625A.982.982 0 0 0 14.943.9a.983.983 0 0 0-.356-.669a.995.995 0 0 0-1.396.137L7.962 6.064L2.772.37a.994.994 0 0 0-1.757.527a.986.986 0 0 0 .219.724l5.798 6.4v1.994H4.03c-.564 0-1.018.463-1.018 1.036c0 .571.455.933 1.018.933h3.002v3c0 .555.42.99.956.99a.99.99 0 0 0 .987-.99v-3h2.973c.563 0 1.019-.361 1.019-.933c0-.573-.455-1.036-1.019-1.036H8.975V7.956z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YingYang(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd" transform="translate(1)"><path d="M7.984.053C3.599.053.045 3.614.045 8.006s3.555 7.953 7.939 7.953s7.939-3.561 7.939-7.953S12.369.053 7.984.053M7.49 2.045c.838 0 1.519.654 1.519 1.46c0 .806-.681 1.46-1.519 1.46c-.84 0-1.519-.654-1.519-1.46c0-.806.679-1.46 1.519-1.46m.545 12.863c-3.051 0-5.693-3.99.066-6.924c5.256-2.676 2.803-7.24-.066-6.976c4.375 0 6.939 3.111 6.939 6.95c.001 3.839-3.107 6.95-6.939 6.95"/><ellipse cx="8.493" cy="11.445" rx="1.493" ry="1.445"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Zip(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M12.945.021H4.054c-.57 0-1.034.45-1.034 1.004v13.972c0 .555.464 1.004 1.034 1.004h8.891c.57 0 1.034-.449 1.034-1.004V1.025c0-.554-.463-1.004-1.034-1.004m-2.914 2.026h-1v.906h1v1.094h-1v.906h1v1.094h-1v.922h.984v4.047H6.984V6.969h1l-.016-1.922h-1V3.953h1v-.906h-1V1.953h1v-1h2.062v1.094z"/><path d="M8 8h.953v.953H8z"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomIn(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M16.913 5.976c0-3.281-2.667-5.939-5.956-5.939C7.666.037 5 2.695 5 5.976c0 3.279 2.666 5.939 5.957 5.939c3.289 0 5.956-2.66 5.956-5.939m-11.004.031c0-2.79 2.271-5.048 5.079-5.048c2.805 0 5.078 2.258 5.078 5.048c0 2.788-2.273 5.047-5.078 5.047c-2.807 0-5.079-2.259-5.079-5.047M2.822 16L1 14.178l4.148-4.149s.086.773.57 1.256c.482.484 1.252.566 1.252.566z"/><path d="M13.254 5.031h-1.285V3.742c0-.41-.557-.742-.972-.742c-.415 0-.966.332-.966.742v1.289H8.74c-.41 0-.742.553-.742.968c0 .415.332.938.742.938h1.291v1.318c0 .41.551.743.966.743c.415 0 .972-.333.972-.743V6.937h1.285c.41 0 .742-.523.742-.938c0-.415-.332-.968-.742-.968"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomOut(children ...ElementRenderer) *SiGlyphIcon {
	return &SiGlyphIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<g fill="currentColor" fill-rule="evenodd"><path d="M17 6.01C17 2.725 14.33.063 11.035.063c-3.294 0-5.964 2.662-5.964 5.947c0 3.284 2.67 5.948 5.964 5.948C14.33 11.958 17 9.294 17 6.01M5.889 6c0-2.807 2.289-5.079 5.116-5.079c2.825 0 5.114 2.272 5.114 5.079c0 2.806-2.289 5.078-5.114 5.078c-2.827 0-5.116-2.272-5.116-5.078m-2.98 9.938l-1.823-1.823l4.148-4.148s.088.773.571 1.256c.483.484 1.252.566 1.252.566z"/><path d="M13.938 5.996c0 .523-.326.948-.729.948H8.771c-.402 0-.729-.425-.729-.948c0-.525.326-.95.729-.95h4.438c.403 0 .729.425.729.95"/></g>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
