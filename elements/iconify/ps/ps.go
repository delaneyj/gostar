package ps

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = ""
	hAttr          = "1em"
	viewbox        = "0 0 512 512"
	fill           = "currentColor"
)

type PsIcon struct {
	*SVGSVGElement
}

type PsIconFn func(children ...ElementRenderer) *PsIcon

var IconLookup = map[string]PsIconFn{
	"aim":                                    Aim,
	"aimAlt":                                 AimAlt,
	"airplane":                               Airplane,
	"alignCentered":                          AlignCentered,
	"alignJustified":                         AlignJustified,
	"alignLeft":                              AlignLeft,
	"alignRight":                             AlignRight,
	"amazon":                                 Amazon,
	"ambulance":                              Ambulance,
	"anchor":                                 Anchor,
	"anySolvent":                             AnySolvent,
	"anySolventWithoutTetrachlorethylene":    AnySolventWithoutTetrachlorethylene,
	"appStore":                               AppStore,
	"apple":                                  Apple,
	"apps":                                   Apps,
	"archive":                                Archive,
	"arrowBox":                               ArrowBox,
	"arrowDown":                              ArrowDown,
	"arrowLeft":                              ArrowLeft,
	"arrowRight":                             ArrowRight,
	"arrowUp":                                ArrowUp,
	"arto":                                   Arto,
	"asterisk":                               Asterisk,
	"attachment":                             Attachment,
	"aws":                                    Aws,
	"backpack":                               Backpack,
	"baidu":                                  Baidu,
	"bankSafe":                               BankSafe,
	"barCode":                                BarCode,
	"barrel":                                 Barrel,
	"basecamp":                               Basecamp,
	"battery":                                Battery,
	"batteryCharge":                          BatteryCharge,
	"bebo":                                   Bebo,
	"beer":                                   Beer,
	"behance":                                Behance,
	"bell":                                   Bell,
	"bike":                                   Bike,
	"bing":                                   Bing,
	"birthday":                               Birthday,
	"blaster":                                Blaster,
	"blip":                                   Blip,
	"blogger":                                Blogger,
	"bnter":                                  Bnter,
	"board":                                  Board,
	"boiledEgg":                              BoiledEgg,
	"boiledEggFinger":                        BoiledEggFinger,
	"bonnet":                                 Bonnet,
	"book":                                   Book,
	"bookTag":                                BookTag,
	"branch":                                 Branch,
	"brightkite":                             Brightkite,
	"brokenLink":                             BrokenLink,
	"browser":                                Browser,
	"bubble":                                 Bubble,
	"bug":                                    Bug,
	"building":                               Building,
	"bullLeft":                               BullLeft,
	"bullRight":                              BullRight,
	"burger":                                 Burger,
	"busLondon":                              BusLondon,
	"calendar":                               Calendar,
	"calendarGrid":                           CalendarGrid,
	"camera":                                 Camera,
	"car":                                    Car,
	"cart":                                   Cart,
	"cartSupermarket":                        CartSupermarket,
	"chat":                                   Chat,
	"chatAlt":                                ChatAlt,
	"check":                                  Check,
	"checkBox":                               CheckBox,
	"checkBoxEmpty":                          CheckBoxEmpty,
	"checked":                                Checked,
	"cinch":                                  Cinch,
	"clock":                                  Clock,
	"clothesWater":                           ClothesWater,
	"cloud":                                  Cloud,
	"cloudapp":                               Cloudapp,
	"clubsCard":                              ClubsCard,
	"cocktail":                               Cocktail,
	"code":                                   Code,
	"coffee":                                 Coffee,
	"coffeeHot":                              CoffeeHot,
	"coin":                                   Coin,
	"coins":                                  Coins,
	"compass":                                Compass,
	"contact":                                Contact,
	"contrast":                               Contrast,
	"cookie":                                 Cookie,
	"copy":                                   Copy,
	"coroflot":                               Coroflot,
	"couple":                                 Couple,
	"cpu":                                    Cpu,
	"creativeCommons":                        CreativeCommons,
	"creditCard":                             CreditCard,
	"crop":                                   Crop,
	"crown":                                  Crown,
	"cutlery":                                Cutlery,
	"daftPunk":                               DaftPunk,
	"dailybooth":                             Dailybooth,
	"dashboard":                              Dashboard,
	"dataBoard":                              DataBoard,
	"delete":                                 Delete,
	"delicious":                              Delicious,
	"designbump":                             Designbump,
	"designfloat":                            Designfloat,
	"designmoo":                              Designmoo,
	"deviantart":                             Deviantart,
	"diamondsCard":                           DiamondsCard,
	"digg":                                   Digg,
	"diggAlt":                                DiggAlt,
	"diigo":                                  Diigo,
	"disabled":                               Disabled,
	"doNotBleach":                            DoNotBleach,
	"doNotDry":                               DoNotDry,
	"doNotIron":                              DoNotIron,
	"doNotWash":                              DoNotWash,
	"doNotWring":                             DoNotWring,
	"dollarBill":                             DollarBill,
	"dollars":                                Dollars,
	"doubleArrow":                            DoubleArrow,
	"down":                                   Down,
	"downArrowCircle":                        DownArrowCircle,
	"download":                               Download,
	"downloadFromCloud":                      DownloadFromCloud,
	"dribbble":                               Dribbble,
	"dripDry":                                DripDry,
	"drop":                                   Drop,
	"dropbox":                                Dropbox,
	"drupal":                                 Drupal,
	"dry":                                    Dry,
	"dryFlat":                                DryFlat,
	"dryInTheShade":                          DryInTheShade,
	"dryNormalHightHeat":                     DryNormalHightHeat,
	"dryNormalLowHeat":                       DryNormalLowHeat,
	"dryNormalNoHeat":                        DryNormalNoHeat,
	"dzone":                                  Dzone,
	"ebay":                                   Ebay,
	"egg":                                    Egg,
	"eject":                                  Eject,
	"ember":                                  Ember,
	"enlarge":                                Enlarge,
	"etsy":                                   Etsy,
	"euroBill":                               EuroBill,
	"evernote":                               Evernote,
	"extinguisher":                           Extinguisher,
	"eye":                                    Eye,
	"facebook":                               Facebook,
	"facebookAlt":                            FacebookAlt,
	"facebookPlaces":                         FacebookPlaces,
	"facto":                                  Facto,
	"feedburner":                             Feedburner,
	"fiftyOneHundredTwenty":                  FiftyOneHundredTwenty,
	"file":                                   File,
	"film":                                   Film,
	"filter":                                 Filter,
	"fire":                                   Fire,
	"fish":                                   Fish,
	"flag":                                   Flag,
	"flagCorner":                             FlagCorner,
	"flagScout":                              FlagScout,
	"flickr":                                 Flickr,
	"folder":                                 Folder,
	"folkd":                                  Folkd,
	"forbidden":                              Forbidden,
	"formspring":                             Formspring,
	"forrst":                                 Forrst,
	"fortyOneHundredFive":                    FortyOneHundredFive,
	"forward":                                Forward,
	"foursquare":                             Foursquare,
	"friedEgg":                               FriedEgg,
	"friendfeed":                             Friendfeed,
	"friendster":                             Friendster,
	"fuck":                                   Fuck,
	"fullScreen":                             FullScreen,
	"gamepad":                                Gamepad,
	"gdgt":                                   Gdgt,
	"gift":                                   Gift,
	"girl":                                   Girl,
	"girlAngel":                              GirlAngel,
	"girlAngry":                              GirlAngry,
	"girlBigSmile":                           GirlBigSmile,
	"girlConfused":                           GirlConfused,
	"girlCry":                                GirlCry,
	"girlFlushed":                            GirlFlushed,
	"girlOmouth":                             GirlOmouth,
	"girlOpenMouth":                          GirlOpenMouth,
	"girlSad":                                GirlSad,
	"girlSadHunappy":                         GirlSadHunappy,
	"girlSleep":                              GirlSleep,
	"girlSmile":                              GirlSmile,
	"girlTwo":                                GirlTwo,
	"girlUser":                               GirlUser,
	"github":                                 Github,
	"githubAlt":                              GithubAlt,
	"globe":                                  Globe,
	"goodreads":                              Goodreads,
	"google":                                 Google,
	"googleBuzz":                             GoogleBuzz,
	"googleTalk":                             GoogleTalk,
	"gowalla":                                Gowalla,
	"gowallaAlt":                             GowallaAlt,
	"grooveshark":                            Grooveshark,
	"gun":                                    Gun,
	"guy":                                    Guy,
	"guyAngel":                               GuyAngel,
	"guyAngry":                               GuyAngry,
	"guyBigSmile":                            GuyBigSmile,
	"guyConfused":                            GuyConfused,
	"guyCry":                                 GuyCry,
	"guyFlushed":                             GuyFlushed,
	"guyHappy":                               GuyHappy,
	"guyOmouth":                              GuyOmouth,
	"guyOpenMouth":                           GuyOpenMouth,
	"guySad":                                 GuySad,
	"guySleep":                               GuySleep,
	"guySmile":                               GuySmile,
	"guyUser":                                GuyUser,
	"guyWrong":                               GuyWrong,
	"hackerNews":                             HackerNews,
	"hand":                                   Hand,
	"handPointerLeft":                        HandPointerLeft,
	"handPointerRight":                       HandPointerRight,
	"handPointerTop":                         HandPointerTop,
	"handWash":                               HandWash,
	"hangToDry":                              HangToDry,
	"hardDrive":                              HardDrive,
	"headphones":                             Headphones,
	"headset":                                Headset,
	"heartsCard":                             HeartsCard,
	"heat":                                   Heat,
	"helm":                                   Helm,
	"hiFive":                                 HiFive,
	"home":                                   Home,
	"hotdog":                                 Hotdog,
	"hourglass":                              Hourglass,
	"hungry":                                 Hungry,
	"hypeMachine":                            HypeMachine,
	"hyves":                                  Hyves,
	"icq":                                    Icq,
	"identi":                                 Identi,
	"image":                                  Image,
	"important":                              Important,
	"instapaper":                             Instapaper,
	"ipad":                                   Ipad,
	"iphone":                                 Iphone,
	"ipod":                                   Ipod,
	"ironAnyTemp":                            IronAnyTemp,
	"itunes":                                 Itunes,
	"iwatch":                                 Iwatch,
	"justice":                                Justice,
	"keyboard":                               Keyboard,
	"kik":                                    Kik,
	"krop":                                   Krop,
	"lab":                                    Lab,
	"label":                                  Label,
	"labelHogwarts":                          LabelHogwarts,
	"laptop":                                 Laptop,
	"last":                                   Last,
	"leaf":                                   Leaf,
	"left":                                   Left,
	"leftArrowCircle":                        LeftArrowCircle,
	"lego":                                   Lego,
	"lightning":                              Lightning,
	"link":                                   Link,
	"linkedin":                               Linkedin,
	"linkedinAlt":                            LinkedinAlt,
	"liquor":                                 Liquor,
	"livejournal":                            Livejournal,
	"lovedsgn":                               Lovedsgn,
	"mac":                                    Mac,
	"machineWash":                            MachineWash,
	"machineWashGentleOrDelicate":            MachineWashGentleOrDelicate,
	"machineWashPermanentPress":              MachineWashPermanentPress,
	"magnifyingGlass":                        MagnifyingGlass,
	"mail":                                   Mail,
	"mailBack":                               MailBack,
	"mailBill":                               MailBill,
	"mailStamp":                              MailStamp,
	"mailbox":                                Mailbox,
	"man":                                    Man,
	"maximumTempOneHundredFiftyThreeHundred": MaximumTempOneHundredFiftyThreeHundred,
	"maximumTempOneHundredTenTwoHundredThirty": MaximumTempOneHundredTenTwoHundredThirty,
	"maximumTempTwoHundredThreeHundredNinety":  MaximumTempTwoHundredThreeHundredNinety,
	"mayoHotdog":                 MayoHotdog,
	"meetup":                     Meetup,
	"megaphone":                  Megaphone,
	"metacafe":                   Metacafe,
	"mic":                        Mic,
	"micOff":                     MicOff,
	"milkshake":                  Milkshake,
	"ming":                       Ming,
	"minus":                      Minus,
	"minusBox":                   MinusBox,
	"minusCircle":                MinusCircle,
	"minusCircleOne":             MinusCircleOne,
	"misterWong":                 MisterWong,
	"mixx":                       Mixx,
	"mixxAlt":                    MixxAlt,
	"mobileme":                   Mobileme,
	"moon":                       Moon,
	"mouse":                      Mouse,
	"msnMessenger":               MsnMessenger,
	"music":                      Music,
	"musicScore":                 MusicScore,
	"myspace":                    Myspace,
	"myspaceAlt":                 MyspaceAlt,
	"newsvine":                   Newsvine,
	"next":                       Next,
	"ninetyFiveTwoHundred":       NinetyFiveTwoHundred,
	"noEye":                      NoEye,
	"nonChlorineBleachIfNeeded":  NonChlorineBleachIfNeeded,
	"official":                   Official,
	"openPadlock":                OpenPadlock,
	"openid":                     Openid,
	"organisation":               Organisation,
	"orkut":                      Orkut,
	"padlock":                    Padlock,
	"pandora":                    Pandora,
	"pant":                       Pant,
	"paperTablet":                PaperTablet,
	"path":                       Path,
	"paypal":                     Paypal,
	"pc":                         Pc,
	"pdiddy":                     Pdiddy,
	"pen":                        Pen,
	"penknife":                   Penknife,
	"peopleTeam":                 PeopleTeam,
	"petroleumSolventSteam":      PetroleumSolventSteam,
	"phone":                      Phone,
	"photobucket":                Photobucket,
	"piano":                      Piano,
	"picasa":                     Picasa,
	"picassa":                    Picassa,
	"piggyBank":                  PiggyBank,
	"piggyBankCoins":             PiggyBankCoins,
	"pin":                        Pin,
	"pinMap":                     PinMap,
	"pinboard":                   Pinboard,
	"ping":                       Ping,
	"pingchat":                   Pingchat,
	"pizza":                      Pizza,
	"plane":                      Plane,
	"play":                       Play,
	"playstation":                Playstation,
	"plixi":                      Plixi,
	"plurk":                      Plurk,
	"plus":                       Plus,
	"plusBox":                    PlusBox,
	"plusCircle":                 PlusCircle,
	"podcast":                    Podcast,
	"posterous":                  Posterous,
	"power":                      Power,
	"preston":                    Preston,
	"previous":                   Previous,
	"printer":                    Printer,
	"prisonSchoolBus":            PrisonSchoolBus,
	"promo":                      Promo,
	"pull":                       Pull,
	"puzzle":                     Puzzle,
	"qik":                        Qik,
	"quik":                       Quik,
	"quora":                      Quora,
	"quote":                      Quote,
	"radio":                      Radio,
	"radioEmpty":                 RadioEmpty,
	"ram":                        Ram,
	"random":                     Random,
	"rdio":                       Rdio,
	"readernaut":                 Readernaut,
	"reddit":                     Reddit,
	"resize":                     Resize,
	"retweet":                    Retweet,
	"retweetOne":                 RetweetOne,
	"rewind":                     Rewind,
	"right":                      Right,
	"rightArrowCircle":           RightArrowCircle,
	"road":                       Road,
	"robo":                       Robo,
	"rowSetting":                 RowSetting,
	"rss":                        Rss,
	"rssIcon":                    RssIcon,
	"safe":                       Safe,
	"saleTag":                    SaleTag,
	"save":                       Save,
	"scissors":                   Scissors,
	"scribd":                     Scribd,
	"seventyOneHundredSixty":     SeventyOneHundredSixty,
	"sharethis":                  Sharethis,
	"shield":                     Shield,
	"shoe":                       Shoe,
	"shoppingCart":               ShoppingCart,
	"sign":                       Sign,
	"simplenote":                 Simplenote,
	"sixtyOneHundredForty":       SixtyOneHundredForty,
	"skype":                      Skype,
	"slashdot":                   Slashdot,
	"slideshare":                 Slideshare,
	"smugmug":                    Smugmug,
	"sound":                      Sound,
	"soundDown":                  SoundDown,
	"soundLevelOne":              SoundLevelOne,
	"soundLevelTwo":              SoundLevelTwo,
	"soundPlus":                  SoundPlus,
	"soundcloud":                 Soundcloud,
	"spadesCard":                 SpadesCard,
	"spotify":                    Spotify,
	"squarespace":                Squarespace,
	"squidoo":                    Squidoo,
	"sreenshot":                  Sreenshot,
	"stats":                      Stats,
	"steam":                      Steam,
	"stethoscope":                Stethoscope,
	"store":                      Store,
	"stumbleupon":                Stumbleupon,
	"suitcase":                   Suitcase,
	"sun":                        Sun,
	"switch":                     Switch,
	"tacos":                      Tacos,
	"tag":                        Tag,
	"target":                     Target,
	"technorati":                 Technorati,
	"thirtyEighty":               ThirtyEighty,
	"threewords":                 Threewords,
	"ticket":                     Ticket,
	"token":                      Token,
	"triangle":                   Triangle,
	"tribe":                      Tribe,
	"tripit":                     Tripit,
	"trophy":                     Trophy,
	"truck":                      Truck,
	"truckOne":                   TruckOne,
	"tumbleDry":                  TumbleDry,
	"tumblr":                     Tumblr,
	"twitter":                    Twitter,
	"twitterAlt":                 TwitterAlt,
	"ufo":                        Ufo,
	"up":                         Up,
	"upArrowCircle":              UpArrowCircle,
	"upload":                     Upload,
	"uploadToCloud":              UploadToCloud,
	"user":                       User,
	"vcard":                      Vcard,
	"viddler":                    Viddler,
	"videoCamera":                VideoCamera,
	"vimeo":                      Vimeo,
	"virb":                       Virb,
	"wthree":                     Wthree,
	"wallet":                     Wallet,
	"wand":                       Wand,
	"warning":                    Warning,
	"watch":                      Watch,
	"waterTemperatureFifty":      WaterTemperatureFifty,
	"waterTemperatureForty":      WaterTemperatureForty,
	"waterTemperatureNinetyFive": WaterTemperatureNinetyFive,
	"waterTemperatureSeventy":    WaterTemperatureSeventy,
	"waterTemperatureSixty":      WaterTemperatureSixty,
	"waterTemperatureThirty":     WaterTemperatureThirty,
	"whatsapp":                   Whatsapp,
	"wikipedia":                  Wikipedia,
	"windows":                    Windows,
	"wists":                      Wists,
	"woman":                      Woman,
	"wordpress":                  Wordpress,
	"wordpressAlt":               WordpressAlt,
	"workCase":                   WorkCase,
	"world":                      World,
	"xing":                       Xing,
	"yahoo":                      Yahoo,
	"yahooBuzz":                  YahooBuzz,
	"yahooMessenger":             YahooMessenger,
	"yelp":                       Yelp,
	"youtube":                    Youtube,
	"youtubeAlt":                 YoutubeAlt,
	"zerply":                     Zerply,
	"zoomIn":                     ZoomIn,
	"zoomOut":                    ZoomOut,
	"zootool":                    Zootool,
	"zynga":                      Zynga,
}

func Aim(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M238 297q88 45 118 116q-51 41-60 49q-43-82-115-101q-60 77-149 101q-15-31-32-71q43-10 65-27.5t44-59.5q-2-53 12.5-96.5T167 145q13 4 40 18t29 15q-2 13 22 18q43 10 76-20l34 57q-7 7-18.5 16T304 270t-65 4q-1 9-1 23m8-164q-27 0-46.5-19.5T180 67t19.5-46T246 2t46.5 19T312 67t-19.5 46.5T246 133"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AimAlt(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M218 13q96 5 172 63.5T462 216q-3 64-63.5 109T242 371q-67 1-115-19q-3 0-3-2q1-1 4 0q29 7 46 3t22.5-12.5t.5-17.5q-4-8-27-14t-56-16t-55-26l1-1q3-8 4-10q2-3 4-3h64q2 0 4 3l7 16q4 11 15 12q12 0 15-11q1-4 0-7v-2l-1-1l-66-148q-2-5-6-5q-2 0-3 1t-1 2t-1 2q-1 1-57 126l-3 6l-9-12Q4 205 2 171q-5-69 57-115.5T218 13m4 253V127q0-16-17-16q-16 0-16 16v139q0 16 16 16q17 0 17-16m196 0V127q0-16-17-16q-9 0-12 4q-1 1-5 7l-48 108l-48-108q-1-2-2.5-4t-1.5-3q-4-4-13-4q-17 0-17 16v139q0 16 17 16t17-16v-65l33 67q4 10 14 10q11 0 17-11l32-66v65q0 16 17 16t17-16m-299-42H81q-2 0-2-3l19-44q0-1 1-1q0-1 1-1l2 1v1l19 44q1 3-2 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Airplane(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m489 230l-41-17v-42q0-43-53-43q-49 0-54 38l-23-10l2-49q0-34-21.5-70.5T245 0t-53 36.5t-21 72.5l2 49l-24 11q0-16-11.5-28.5T96 128q-53 0-53 43v51l-20 8Q0 242 0 269v77l85-20v15q0 22 22 22q9 0 15-6t6-16v-21h-9l62-13l5 86q-79 32-79 98v21h277v-21q0-71-79-101l4-89l84 19h-9v21q0 10 6 16t15 6q22 0 22-22v-12l85 19v-77q0-26-23-41m-105-59h21v23l-21-8zm-299 0h22v17l-22 13zM43 294v-23l132-66l4 60zm111 175q5-12 34-32l2 32zm79 0l-20-362q0-19 12-41.5T245 43t20 22t12 40l-19 364zm104 0h-38l2-34q29 13 36 34m132-175l-158-34l3-57l153 68v23z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignCentered(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M85 43q-21 0-21 21t21 21h214q21 0 21-21t-21-21zM21 149q0 22 22 22h298q22 0 22-22q0-9-6-15t-16-6H43q-10 0-16 6t-6 15m86 64q-22 0-22 22q0 9 6 15t16 6h170q10 0 16-6t6-15q0-22-22-22zM0 320q0 21 21 21h342q21 0 21-21t-21-21H21q-21 0-21 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignJustified(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 85h342q21 0 21-21t-21-21H21Q0 43 0 64t21 21m0 86h342q9 0 15-6t6-16q0-9-6-15t-15-6H21q-9 0-15 6t-6 15q0 10 6 16t15 6m0 85h342q9 0 15-6t6-15q0-10-6-16t-15-6H21q-9 0-15 6t-6 16q0 9 6 15t15 6m0 85h342q21 0 21-21t-21-21H21q-21 0-21 21t21 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignLeft(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 85h214q21 0 21-21t-21-21H21Q0 43 0 64t21 21m0 86h299q21 0 21-22q0-21-21-21H21q-9 0-15 6t-6 15q0 10 6 16t15 6m0 85h256q10 0 16-6t6-15q0-22-22-22H21q-9 0-15 6t-6 16q0 9 6 15t15 6m0 85h342q21 0 21-21t-21-21H21q-21 0-21 21t21 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AlignRight(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M363 85q21 0 21-21t-21-21H149q-21 0-21 21t21 21zM43 149q0 22 21 22h299q9 0 15-6t6-16q0-9-6-15t-15-6H64q-21 0-21 21m320 64H107q-22 0-22 22q0 9 6 15t16 6h256q9 0 15-6t6-15q0-10-6-16t-15-6M0 320q0 21 21 21h342q21 0 21-21t-21-21H21q-21 0-21 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Amazon(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M274 114q-27 2-37 4q-34 4-55 14q-54 20-54 80q0 37 21 57t53 20q18 0 35-5q23-7 47-31q15 21 28 32q6 3 12-1q5-4 19-16t19-16q6-6 0-13q-18-24-18-46V87.5l-.5-9.5l-1.5-8.5l-2-9l-2.5-7.5l-4-8l-5-7l-6.5-7Q294 6 248 6h-11q-87 6-100 76q-2 8 7 10l48 6q8-2 8-9q7-27 37-31h4q17 0 27 13q6 9 6 36zm0 50q0 37-9 53q-8 17-28 22l-3 1h-4q-13 0-21.5-10t-8.5-26q0-36 37-47q16-3 37-3zm158 241q14-12 22-31.5t8-34.5v-8l-1-3q-4-4-22.5-5.5T399 326q-9 2-20 10q-6 5 1 7q5-1 19-3q34-2 40 5q7 8-13 57q-2 4 .5 5t5.5-2M4 339q97 87 228 87q94 0 167-45q7-4 19-12q6-5 2.5-10t-9.5-2q-2 1-6 2t-6 2q-76 31-162 31q-123 0-228-60q-4-3-6.5 0t1.5 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ambulance(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M405 85h-42V43q0-18-13-30.5T320 0H43Q25 0 12.5 12.5T0 43v256q0 17 12.5 29.5T43 341h4q6 19 22.5 31t37.5 12q20 0 36.5-12t22.5-31h180q6 19 22.5 31t36.5 12q21 0 37.5-12t22.5-31h4q18 0 30.5-12.5T512 299v-77q0-17-13-30zM107 341q-8 0-15-6.5T85 320t7-14.5t15-6.5t14.5 6.5T128 320t-6.5 14.5T107 341m213-42H166q-6-19-22.5-31T107 256q-21 0-37.5 12T47 299h-4V43h277zm85 42q-8 0-14.5-6.5T384 320t6.5-14.5T405 299t15 6.5t7 14.5t-7 14.5t-15 6.5m64-42h-4q-6-19-22.5-31T405 256q-27 0-42 17V128h12l94 94zm-298-86h42v-42h43v-43h-43V85h-42v43h-43v43h43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Anchor(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M171 489q0 9 6 16t15 7q21 0 21-21v-2q66-8 113.5-54.5T382 322q3-19-19-23q-8-2-15.5 3.5T339 318q-6 50-41.5 85.5T213 446V256h64q10 0 16-6t6-15q0-22-22-22h-64v-44q28-7 46-30.5T277 85q0-35-25-60T192 0t-60 25t-25 60q0 30 18 53t46 31v44h-64q-22 0-22 22q0 9 6 15t16 6h64v190q-50-7-86-43t-40-85q-3-22-26-19q-19 4-19 23q10 67 57.5 113.5T171 489M149 85q0-17 13-29.5T192 43t30 12.5T235 85q0 18-13 30.5T192 128t-30-12.5T149 85"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AnySolvent(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m309 365l-15-51h-76l-15 51h-49l74-214h56l74 214zm-25-88q-14-43-24-79l-4-12q-2 5-28 91zM256 0Q150 0 75 75T0 256t75 181t181 75t181-75t75-181t-75-181T256 0m0 469q-88 0-150.5-62.5T43 256t62.5-150.5T256 43t150.5 62.5T469 256t-62.5 150.5T256 469"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AnySolventWithoutTetrachlorethylene(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M337 220q0 35-21 53q-22 19-62 19h-19v77h-45V156h68q37 0 60 17q19 12 19 47m-102 32h15q20 0 32-9q10-7 10-23q0-17-8-24q-14-8-28-8h-21zM256 0Q150 0 75 75T0 256t75 181t181 75t181-75t75-181t-75-181T256 0m0 469q-88 0-150.5-62.5T43 256t62.5-150.5T256 43t150.5 62.5T469 256t-62.5 150.5T256 469"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppStore(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 255v-63q0-10 10-10h111l-49 83H12q-10 0-10-10m78 128l-35 24q-3 1-4.5 1t-2-1.5t-.5-4.5l4-41q2-10 11-5l27 16q9 5 0 11m9-23l-28-16q-8-5-3-14l126-219q5-9 14-4l28 16q8 5 3 14L103 356q-5 9-14 4m96-95l48-83h28l39 83zm157 25q-13-6-80-163q-2-5-7.5-17T245 88.5t-9.5-23t-9-22.5t-6.5-19.5t-2.5-15T220 1q4-2 15 13.5T264 61t27 44q11 18 31 54.5t37 69t19 35.5q4 8-1.5 15T364 289q-12 7-22 1m35 58l-16-26q-5-10 4-15l16-8q10-6 15 5l13 25q6 10-5 16l-11 6q-8 7-16-3m61 67q-2-5-19-12t-23-16q-12-17 3-29q21-12 30.5 2.5T438 395zm24-160q0 10-10 10h-55q-2-6-4-9q-1-3-10-19q-4-8-14.5-28T354 182h98q10 0 10 10z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Apple(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M216 37q27-33 67-35q4 39-24 73q-29 36-68 33q-4-38 25-71m131 368q-12 18-20 28t-22 19.5t-28 9.5q-15 0-36.5-9.5T201 443q-19 0-40.5 9.5T126 462q-14 1-28.5-9T75 433t-22-29q-21-30-34-69.5T5.5 250T24 172q32-56 95-58q17 0 43 10t33 10q5 0 36-11.5t52-9.5q55 4 84 46q-5 3-14.5 10T330 198t-13 48q0 11 2.5 21.5t6 18.5t8 15t9.5 12.5t10 10t9.5 7.5t8 5t5.5 3l3 1l-3 10q-4 9-11.5 25T347 405"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Apps(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M133 45q0-17-12.5-29.5T91 3H48Q31 3 18 15.5T5 45v43q0 17 13 30t30 13h43q17 0 29.5-13T133 88zM91 88H48V45h43zm213-43q0-17-12.5-29.5T261 3h-42q-18 0-30.5 12.5T176 45v43q0 17 12.5 30t30.5 13h42q18 0 30.5-13T304 88zm-43 43h-42V45h42zm214-43q0-17-13-29.5T432 3h-43q-17 0-29.5 12.5T347 45v43q0 17 12.5 30t29.5 13h43q17 0 30-13t13-30zm-43 43h-43V45h43zM133 195q0-18-12.5-30.5T91 152H48q-17 0-30 12.5T5 195v42q0 18 13 30.5T48 280h43q17 0 29.5-12.5T133 237zm-42 42H48v-42h43zm213-42q0-18-12.5-30.5T261 152h-42q-18 0-30.5 12.5T176 195v42q0 18 12.5 30.5T219 280h42q18 0 30.5-12.5T304 237zm-43 42h-42v-42h42zM91 301H48q-17 0-30 13T5 344v43q0 17 13 29.5T48 429h43q17 0 29.5-12.5T133 387v-43q0-17-12.5-30T91 301m0 86H48v-43h43zm170-86h-42q-18 0-30.5 13T176 344v43q0 17 12.5 29.5T219 429h42q18 0 30.5-12.5T304 387v-43q0-17-12.5-30T261 301m0 86h-42v-43h42zm171-235h-43q-17 0-29.5 12.5T347 195v42q0 18 12.5 30.5T389 280h43q17 0 30-12.5t13-30.5v-42q0-18-13-30.5T432 152m0 85h-43v-42h43zm-43 192h43q17 0 30-12.5t13-29.5v-43q0-17-13-30t-30-13h-43q-17 0-29.5 13T347 344v43q0 17 12.5 29.5T389 429m0-85h43v43h-43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Archive(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469 3H43Q25 3 12.5 15.5T0 45v86q0 9 6 15t15 6h22v235q0 17 12.5 29.5T85 429h342q17 0 29.5-12.5T469 387V152h22q9 0 15-6t6-15V45q0-17-12.5-29.5T469 3M85 387V152h342v235zm384-278H43V45h426zm-140 90q-32 24-73 24t-73-24q-6-5-15.5-3.5T154 203q-5 7-4 16.5t8 13.5q46 32 98 32t98-32q7-4 8-13.5t-4-16.5q-4-6-13.5-7.5T329 199"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowBox(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 277q0 28 18 46t46 18h170q28 0 46-18t18-46V107q0-28-18-46t-46-18H67q-28 0-46 18T3 107zm42-170q0-22 22-22h170q22 0 22 22v170q0 22-22 22H67q-22 0-22-22zm171 42H88q0-1 16 20.5t32 43.5l16 22z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDown(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M152 0q-21 0-21 21v297l-94-77q-7-6-16-5t-14 7q-6 7-5 16t7 14l143 111l141-111q15-15 2-30q-16-14-30-2l-92 77V21q0-21-21-21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeft(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 192l111 141q15 15 30 2q14-16 2-30l-77-92h297q21 0 21-21t-21-21H66l77-94q6-7 5-16t-7-14q-17-12-30 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRight(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M243 47q-15 15-2 30l77 94H21q-21 0-21 21t21 21h297l-77 94q-6 7-5 16t7 14q7 6 16 5t14-7l111-143L273 51q-13-16-30-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUp(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M297 141q6-7 5-16t-7-14L152 0L11 111q-15 15-2 30q16 14 30 2l92-77v297q0 21 21 21t21-21V66l94 77q17 15 30-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Arto(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M262 462q-46 0-70-42q-12-20-32-20q-8 0-28 4q-18 4-28 4l-1 1l-1-1q-73-13-93-56q-15-36 12-90L218 35q44-33 77-33q31 0 46 24.5T357 77l-1 11q0 11-.5 32t-.5 40q0 17 3 29.5t6 16.5l3 4q12 8 15 21q2 21-14 37l-2 3h-4q-9 7-10 80q0 14-1 18q-1 54-30 76q-31 17-59 17m-102-79q31 0 47 28q21 34 55 34q23 0 50-15q21-16 22-62q0-10 1-18q2-88 22-95q9-15 7-21q0-5-7-10q-20-13-19-64l2-82q0-6-2-17t-13.5-26.5T295 19q-26 0-66 29L35 272q-3 4-6 11t-8 28t4 34q15 34 79 46q8 0 24-4q22-4 32-4m152-143l-28 1V112q0-1-.5-3.5t-3-7T274 95t-11-1t-17 8L71 308q-12 17 0 20q1 1 3.5 2t8.5 3.5t10-.5l7-7q6-6 12-11.5t18.5-14t26-15t34-12.5t43.5-9l1 120q3 3 8 7t18.5 4.5T287 384q0-16-.5-72t-.5-59q1 0 3.5-.5t9.5-1.5l14-2q14-4-1-9m-74 1q-40 3-56 4q-7 0-8-1.5t2-3.5l2-2l60-67z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Asterisk(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 213h128L49 305q-13 15 0 30q6 6 15 6t15-6l92-100v128q0 21 21 21t21-21V235l92 100q6 6 15 6t15-6q13-15 0-30l-100-92h128q21 0 21-21t-21-21H235l100-92q13-15 0-30q-15-13-30 0l-92 100V21q0-21-21-21t-21 21v128L79 49q-15-13-30 0q-13 15 0 30l100 92H21q-21 0-21 21t21 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Attachment(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M416 35Q382 1 333.5 1T250 35L32 254Q6 280 6 315t26 61t60 26q33 0 59-26l220-220q15-14 15-38q0-21-15-39q-16-16-38-16t-36 16L94 282q-15 15 0 30q14 16 30 0l200-203q8-7 15 0q7 8 0 15L122 344q-13 13-30 13t-30-13t-13-30t13-30L279 65q23-23 54-23t53 23q22 22 22 52.5T386 171L198 359q-14 14 0 30q16 14 30 0l188-188q34-35 34-83t-34-83"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Aws(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m2 87l204 67v292L3 377zm255 360l205-70V89l-204 64zm-25-328l188-61L232 1L39 58z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Backpack(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M280 53V43q0-18-12.5-30.5T237 0h-42q-18 0-30.5 12.5T152 43v10Q86 73 44.5 129T3 256v192q0 27 18 45.5T67 512h298q28 0 46-18.5t18-45.5V256q0-71-41.5-127T280 53m107 395q0 21-22 21H67q-22 0-22-21V256q0-70 50.5-120.5T216 85t120.5 50.5T387 256zm-86-192H131q-18 0-30.5 12.5T88 299v106q0 18 12.5 30.5T131 448h170q18 0 30.5-12.5T344 405V299q0-18-12.5-30.5T301 256m0 149H131v-42h170zm0-85H131v-21h170zM195 213h42v-21h22v-43h-22v-21h-42v21h-22v43h22z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Baidu(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M109 174q2 29-12 51t-36 23t-38.5-18.5T4 181t11.5-51.5T51 106q22-2 39 18t19 50M308 11q-21-5-41.5 11.5T239 67q-7 29 3 53t31 29t41-11.5T341 93q6-22-1.5-42t-17-29T308 11m66 128q-23 1-38 21t-15 52q1 31 16 47t39 16q53-2 52-65q0-26-12.5-43.5t-24-22.5t-17.5-5M160 2q-22 0-36.5 20.5T109 72t14.5 49.5T160 142q21 0 35.5-20.5T210 72t-14.5-49.5T160 2m-7 218q-25 37-70 76q-51 42-51 81q0 29 16.5 55T98 458q26 0 60-5.5t52-5.5t53.5 7.5T325 462q33 0 51.5-24.5T395 385q0-23-9.5-39.5T347 303q-47-39-81-87q-18-26-56-26q-36 0-57 30"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BankSafe(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 0H64Q37 0 18.5 18.5T0 64v405q0 18 12.5 30.5T43 512h426q18 0 30.5-12.5T512 469V64q0-27-18.5-45.5T448 0m21 469H43V64q0-21 21-21h384q21 0 21 21zM384 64H128q-17 0-30 12.5T85 107v42H64v43h21v128H64v43h21v42q0 18 13 30.5t30 12.5h256q27 0 45.5-18.5T448 384V128q0-27-18.5-45.5T384 64m21 320q0 21-21 21H128v-42h21v-43h-21V192h21v-43h-21v-42h256q21 0 21 21zm-70-100q6-13 6-28q0-20-12-36.5T299 196v-47q0-9-6-15t-16-6q-9 0-15 6t-6 15v47q-43 15-43 60q0 15 7 28l-43 42q-14 16 0 30q6 7 15 7t15-7l43-42q12 6 27 6q16 0 28-6l43 42q6 7 15 7q8 0 15-7q14-14 0-30zm-58-7q-21 0-21-21t21-21q22 0 22 21t-22 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarCode(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h43v427H0zm171 427V0H85v427zM213 0h43v427h-43zm86 0h42v427h-42zM0 469h512v43H0zM469 0h-42v427h85V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Barrel(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M351 0H125Q97 0 78 21Q44 60 24.5 122T5 254q0 81 22 148.5T84 497q15 15 41 15h228q31 0 49-23q73-87 73-231q0-155-77-239q-19-19-47-19m79 213h-83q-3-47-9-85h73q12 31 19 85M185 384q-7-36-7-85h126q-2 37-6 85zm102 43q-5 22-17 42h-64q-4-6-17-42zM69 393q-13-47-19-94h83q0 43 7 85H69zm224-265q5 27 9 85H178q4-58 9-85zm-98-43q3-11 15-42h60q10 19 15 42zm66 171q0 8-6.5 14.5T240 277t-14.5-6.5T219 256t6.5-14.5T240 235t14.5 6.5T261 256m-128-43H50q3-48 19-91v6h73q-6 38-9 85m214 86h83q-3 38-17 85h-73q1-17 3.5-45.5T347 299m19-250q18 18 25 36h-64q-8-30-12-42h36q12 0 15 6m-241-6h38q-2 6-6 20t-7 22H84q14-23 24-36q6-6 17-6m-13 422q-13-9-28-38h66q3 8 7 22t6 20h-38q-9 0-13-4m241 4h-34q2-6 6-20t7-22h62q-9 16-24 34q-5 8-17 8m79-211"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Basecamp(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 295q6 10 10 14q9 12 25 24q27 22 65 33q38 13 83 16q51 5 98-1q50-5 88-20t68-42q10-10 14-16q7-7 7-16q1-9-2-21q-1-9-5-22q-2-6-2-8q-12-50-34-94q-20-39-47-70q-28-32-61-50q-23-11-39-16q-14-4-22-5q-2-1-5-1h-28q-1 0-3 .5t-3 .5q-2 1-6 1t-6 1q-14 4-21 7q-25 11-40 22q-32 21-63 61q-26 35-46 82q-16 42-26 96q0 2-.5 6t-.5 7v9zm30-20q0-4 1-7l3-12q10-24 22-42q2-4 6-10t6-10q1-2 6-9t8-11q16-20 31-26q5 0 7.5-1t7 1.5t6 3.5t6 5t5.5 5t5 5.5t5 5.5q10 12 20 17q6 3 12 3q8-1 16-9q15-12 32-36q3-3 14.5-18t18.5-22q10-12 17-15q3-1 10 1q2 1 4 3l3 3l2 1q17 13 36 32q17 17 28 32q16 18 25 33q8 14 24 44q15 28 7 44q0 1-1 2q-17 28-60 45q-40 16-89 19q-50 3-93-1q-50-5-83-18q-39-16-58-40q-3-4-9-14q0-1-.5-1t-.5-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Battery(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 107q0-28-18.5-46T384 43H64q-27 0-45.5 18T0 107v170q0 28 18.5 46T64 341h320q27 0 45.5-18t18.5-46q27 0 45.5-18t18.5-46v-42q0-28-18.5-46T448 107m21 106q0 22-21 22q-17 0-30 12.5T405 277q0 22-21 22H64q-21 0-21-22V107q0-22 21-22h320q21 0 21 22q0 17 13 29.5t30 12.5q21 0 21 22zM85 128h235v128H85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryCharge(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 115q0-28-18.5-46T384 51H277v42h107q21 0 21 22q0 17 13 29.5t30 12.5q21 0 21 22v42q0 22-21 22q-17 0-30 12.5T405 285q0 22-21 22h-64v42h64q27 0 45.5-18t18.5-46q27 0 45.5-18t18.5-46v-42q0-28-18.5-46T448 115M43 285V115q0-22 21-22h85V51H64q-27 0-45.5 18T0 115v170q0 28 18.5 46T64 349h128v-42H64q-21 0-21-22m143-96l76-170l-38-17l-111 251l171-44l-77 196l40 15l109-271z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bebo(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M231 462h-46q-77 0-130.5-54T1 278V48q0-20 13.5-33T47 2t32.5 13.5T93 48v230q0 38 27 65t65 27h46q38 0 65-23.5t27-56.5t-14.5-57t-31.5-24h-92q-19 0-32.5-13.5T139 163t13.5-32.5T185 117h92q57 0 97.5 50.5T415 290q0 71-54 121.5T231 462"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Beer(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M431 107q25-1 42.5-19.5T491 43V0H107v21H21q-9 0-15 6T0 43q0 9 6 15t15 6h22v107q0 43 23.5 80t61.5 54l-21 122l-22 85h427l-21-85zm13 298H154l59-298h171zm4-362q0 9-6 15t-15 6H171q-10 0-16-6t-6-15zM85 171V64h26q6 17 21 29.5t34 13.5l-29 155q-52-31-52-91m52 298l6-21h311l7 21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Behance(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M418 89H302V61h116zM225 207q12 17 12 42q0 23-13 46q-12 16-20 22q-16 11-33 14q-16 4-40 4H2V49h138q53 0 74 30q13 17 13 44q0 28-13 42q-9 11-22 16q23 9 33 26M68 162h60q17 0 31-7q11-8 11-26q0-20-15-26q-15-5-34-5H68zm108 83q0-22-18-31q-12-5-29-5H68v77h60q17 0 29-5q19-10 19-36m285-47q2 20 1 41H313q1 31 21 43q13 8 30 8q19 0 30-10q6-5 11-14h55q-3 18-20 38q-29 29-78 29q-41 0-72-25t-31-82q0-54 28-82t73-28q26 0 49 9q22 11 35 31q12 16 17 42m-54 5q-2-22-15-32q-13-11-32-11q-20 0-32 12q-11 11-14 31z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bell(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M478 343q-3-1-8-3.5t-17.5-14t-22-26.5t-17.5-43.5t-8-63.5q0-37-11.5-64T359 84t-35-23t-36-16h-2q-2-2-9-2V0h-42v43q-7 0-7 2h-2q-19 7-28.5 11t-30 17t-31 26.5t-20 38T107 192q0 114-73 151l-4 3l-30 29v73h171q0 27 25 45.5t60 18.5t60-18.5t25-45.5h171v-73l-30-29zM256 469q-18 0-30.5-7T213 448h86q0 7-12.5 14t-30.5 7m213-64H43v-12l12-13q11-5 23.5-15t30-31t29-58t11.5-84q0-27 8-46t23.5-31T208 97.5T239 85h34q19 7 31 12.5t27.5 17.5t23.5 31t8 46q0 47 11.5 84t29 58t30 31t23.5 15l12 13z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bike(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M382 382q50 11 90-22t40-83q0-41-36-74q-32-32-92-32V64q0-27-18.5-45.5T320 0q-21 0-21 21q0 22 21 22q12 0 27.5 6T363 64v21H171V64h42q0-17-12.5-30T171 21h-64q0 12 14 25t28 18v43l-23 66q-55-13-107 51q-19 29-19 53q0 45 31 76t76 31q38 0 67-24t37-61h15q33 0 49-24l32-38q-13 30-4 68q6 27 28 48.5t51 28.5m-275-41q-28 0-46-18t-18-46q0-27 18-45.5t46-18.5h6l-28 86h81q-6 19-22.5 30.5T107 341m42-85v-21q18 12 22 21zm15-68l9-22l15 43q-23-21-24-21m209 85q7 8 16.5 8.5T405 275q12-10 2-28q-7-12-14-34q32-7 61 28q15 18 15 38q0 30-22 48.5T395 341q-20-3-34-16.5T343 292q-6-33 13-55q6 16 17 36m-130-15l-42-130h119zm77-45l21-42v21q-12 9-21 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bing(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 0v246q2 64 69 109t161 45q95 0 162.5-46.5T462 242t-67.5-111.5T232 84q-82 0-147 36V0zm230 137q61 0 104 31t43 74t-43 74t-104 31t-104-31t-43-74t43-74t104-31m-97 105q0-29 28.5-49t68.5-20t68.5 20t28.5 49t-28.5 49.5T232 312t-68.5-20.5T135 242"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Birthday(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 224q0-31-22-53t-53-22h-74V85h-43v64h-43V85h-42v64h-43V85h-43v64H75q-31 0-53 22T0 224q0 48 43 66v179H21q-9 0-15 6t-6 16q0 9 6 15t15 6h470q9 0 15-6t6-15q0-10-6-16t-15-6h-22V290q43-18 43-66M75 192h362q13 0 22.5 9.5T469 224t-9.5 22.5T437 256H75q-13 0-22.5-9.5T43 224t9.5-22.5T75 192m10 171h342v42H85zm0-22v-42h342v42zm0 128v-42h342v42zM277 32q0-14-6-23t-15-9t-15 9t-6 23t6 23t15 9t15-9t6-23m86 0q0-13-6.5-22.5T341 0q-8 0-14.5 9.5T320 32t6.5 22.5T341 64q9 0 15.5-9.5T363 32m-171 0q0-13-6.5-22.5T171 0q-9 0-15.5 9.5T149 32t6.5 22.5T171 64q8 0 14.5-9.5T192 32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Blaster(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M491 171h-22v-64q0-11-10-11q-11 0-11 11v64h-26q-11-47-50.5-80.5T277 47v-4q0-19-19-22L45 0Q31 0 26 11q-8 12 0 21l49 83q-32 38-32 88q0 41 21 70l-34 64Q0 381 0 427q0 30 27.5 57.5T85 512q21 0 32-20.5t22-62.5q1-6 2-9q6-31 8-36h22q26 0 49-21h15q76 0 131.5-43.5T427 213h21v64q0 11 11 11q10 0 10-11v-64h22q21 0 21-21t-21-21M162 55q-27 9-57 30L83 47zM98 412q-9 44-17 57q-13-3-25.5-16.5T43 427q0-33 21-64l28-54q16 14 34 24q-9 12-28 79m137-92q-84 0-128-55q-22-29-22-62q0-48 43.5-83T235 85q62 0 105.5 35t43.5 83t-43.5 82.5T235 320m100-160q-13 7-34.5 20T271 198q-1-3-17-55l-11-24l-17 15l-85 69q-11 6-2 19q7 11 19 4l64-45q8 42 15 60l8 32l24-19q44-33 79-71q5-5 2-14q-1-6-6-8.5t-9-.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Blip(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m2 408l33-34v-53l108 108H90l-34 33zm27-260v83q84 0 145 59q59 59 59 145h83q0-118-84.5-202.5T29 148M29 2v82q144 0 247 103.5T379 435h83q0-88-34.5-168T335 129T197 36.5T29 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Blogger(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M315 462q61 0 103.5-42.5T461 317l1-118l-1-6l-4-8l-6-5q-4-3-30.5-4t-32.5-6q-9-8-12-40q-4-27-13-49q-15-31-49-53.5T250 2H148Q87 2 44.5 44.5T2 147v170q0 60 42.5 102.5T148 462zM149 121h81q12 0 20 8t8 19q0 12-8 20t-20 8h-81q-11 0-19-8t-8-20q0-11 8-19t19-8m-27 194q0-11 8-19.5t19-8.5h165q11 0 19.5 8.5T342 315t-8.5 19.5T314 343H149q-11 0-19-8.5t-8-19.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bnter(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M38 321h17q4 21-9 43q5 0 13.5-2t25-12.5T105 321h321q15 0 25.5-10.5T462 285V57q0-15-10.5-26T426 20H38q-15 0-25.5 11T2 57v228q0 15 10.5 25.5T38 321m350-125q7 0 12 5t5 12t-5 12t-12 5t-12-5t-5-12q0-17 17-17m-104 52q1 1 1 4q0 18-13 30.5T241 295q-14 0-25.5-8.5T200 265zM78 196q17 0 17 17q0 7-5 12t-12 5t-12-5t-5-12t5-12t12-5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Board(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M491 320h-24l-17-196q-1-25-19.5-42.5T388 64H277V21q0-21-21-21t-21 21v43H124q-24 0-42.5 17.5T62 124L45 320H21q-9 0-15 6t-6 15q0 10 6 16t15 6h79L64 484q-3 8 2 16t13 10h6q13 0 22-15l38-132h90v85q0 21 21 21t21-21v-85h90l38 134q3 15 22 15h6q19-7 15-26l-36-123h79q9 0 15-6t6-16q0-9-6-15t-15-6m-404 0l18-194q0-8 6.5-13.5T126 107h262q9 0 15.5 5.5T410 126l17 194z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoiledEgg(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 256q0 49 29.5 89t74.5 54q-5 27-16 48.5T67 469H45q-9 0-15 6t-6 16q0 9 6 15t15 6h214q9 0 15-6t6-15q0-10-6-16t-15-6h-22q-13 0-24-21.5T197 399q48-14 76-58t28-106v-32q0-62-40-126Q214 0 152 0T43 77Q3 141 3 203zm175 213h-49q17-30 21-64h4q10 41 24 64m-26-106q-45 0-76-31t-31-76h212q-4 45-30 76t-75 31M45 203q0-50 34-105q36-55 73-55t73 55q34 55 34 105v10H45z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BoiledEggFinger(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 264q0 49 29.5 89t74.5 54q-5 27-16 48.5T67 477H45q-9 0-15 6t-6 16q0 9 6 15t15 6h214q9 0 15-6t6-15q0-10-6-16t-15-6h-22q-13 0-24-21.5T197 407q48-14 76-58t28-106v-30q0-12-8-56v-6l-22-34l-49 32l39-115q3-8-1-16.5T248 6q-19-7-28 13l-36 106l-32-15l-64 43l-51-34l-22 32v2Q3 184 3 211zm175 213h-49q17-30 21-64h4q10 41 24 64m-26-106q-45 0-76-31t-31-76h212q-4 45-30 76t-75 31M45 211q0-12 5-32l38 25l64-42l64 42l38-25q5 20 5 32v10H45z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bonnet(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M432 341v-42q0-82-55-137q-24-23-60-40q8-22 8-37q0-35-25-60T240 0t-60 25t-25 60q0 15 8 37q-52 24-83.5 72T48 299v42q-17 0-30 13T5 384v85q0 18 13 30.5T48 512h384q17 0 30-12.5t13-30.5v-85q0-17-13-30t-30-13M240 43q17 0 30 12.5T283 85q0 14-9 26q-22-4-34-4t-34 4q-9-12-9-26q0-17 13-29.5T240 43M91 299q0-52 31.5-92t81.5-53q35-11 74 0q34 9 69 38q42 44 42 107v42H91zm298 170h-42v-85h42zm-256 0H91v-85h42zm22-85h42v85h-42zm64 0h42v85h-42zm64 0h42v85h-42zm-235 0h21v85H48zm363 85v-85h21v85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Book(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M67 512h362v-43H67q-22 0-22-21t22-21h362V0H67Q39 0 21 18.5T3 64v363h4q-4 12-4 21q0 27 18 45.5T67 512M45 64q0-21 22-21h320v341H67q-8 0-22 4zm86 107h192v42H131zm0-86h192v43H131z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookTag(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M88 512h341v-43H88q-17 0-30-12.5T45 427q0-18 13-30.5T88 384h341V0H67Q39 0 21 18.5T3 64v363q0 35 25 60t60 25M45 64q0-21 22-21h320v298H88q-19 0-43 11zm22 363q0 21 21 21h341v-43H88q-21 0-21 22M344 85H109v128h235zm-43 86H152v-43h149z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Branch(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M299 0q-35 0-60.5 25T213 85t25.5 60.5T299 171t60-25.5T384 85t-25-60t-60-25m0 128q-18 0-30.5-12.5T256 85q0-17 12.5-29.5T299 43q17 0 29.5 12.5T341 85q0 18-12.5 30.5T299 128M149 299h-42V169q27-7 45.5-30.5T171 85q0-35-25.5-60T85 0T25 25T0 85q0 30 18 52.5T64 169v177q-27 6-45.5 29.5T0 429q0 35 25 60t60 25t60.5-25t25.5-60q0-30-18-52.5T107 346v-5h42q71 0 121-50t50-120h-43q0 53-37.5 90.5T149 299M43 85q0-17 12.5-29.5T85 43q18 0 30.5 12.5T128 85q0 18-12.5 30.5T85 128q-17 0-29.5-12.5T43 85m85 342q0 17-12.5 29.5T85 469q-17 0-29.5-12.5T43 427q0-18 12.5-30.5T85 384q18 0 30.5 12.5T128 427"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Brightkite(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M277 203q0 31-7 43q-4 7-14 9t-18 0l-7-2v-84q1-1 2.5-3t8.5-4t17-1q18 2 18 42M431 2v290q0 18-20 27q-13 6-115.5 42T97 430L1 462L85 29q1-3 3-7.5t13.5-12T129 2zM320 147q-9-24-29.5-30t-37.5 1.5t-18 16.5h-4V49l-56 9v216q4 3 11 8t29.5 11.5T262 296t39.5-17.5T320 254l4-12q15-44-4-95"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrokenLink(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m28 187l45 44q6 7 14 7q9 0 15-7q16-15 0-29l-44-45q-13-13-13-30t13-30l29-30q14-13 30.5-13T147 67l47 45q15 15 30 0t0-30l-45-47q-25-25-60-25T58 35L28 67Q2 93 2 127t26 60m399 226q25-25 25-61t-25-61l-45-45q-15-13-30 0q-16 16 0 30l45 45q13 15 13 30t-13 30l-30 30q-13 13-31 13t-29-13l-47-43q-14-16-30 0q-13 15 0 30l45 45q25 25 60 25t60-25zM294 148q7 7 15 7q9 0 15-7l86-85q13-15 0-30q-16-14-30 0l-86 85q-14 16 0 30M53 411q9 0 15-7l86-85q13-15 0-30q-16-14-30 0l-86 85q-14 16 0 30q7 7 15 7m101 40q2 2 6 2q16 0 19-15l22-64q3-7-1-15.5T188 347q-8-4-16.5 0T160 359l-21 64q-2 10 2.5 18.5T154 451m303-247q3-8-1-16.5T444 176l-64-21q-8-4-16.5 0T352 167t1 16.5t12 11.5l64 22q1 0 3 1t3 1q16 0 22-15m-190-92q4 0 6-2q19-7 13-28l-21-64q-9-19-28-13q-19 9-13 28l21 64q8 15 22 15M102 281q19-9 13-28q-7-19-28-13l-64 21q-19 10-12 28q7 15 21 15q4 0 6-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Browser(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469 3H43Q25 3 12.5 15.5T0 45v363q0 21 21 21h470q21 0 21-21V45q0-17-12.5-29.5T469 3m0 384H43V152h426zm0-278H43V45h426zm-277 86H85v149h107zm-43 106h-21v-64h21zm86-106h192v42H235zm0 85h149v43H235z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bubble(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M216 368q0-21-21-21h-86q-21 0-21 21t21 21h86q21 0 21-21m-85 43q-22 0-22 21t22 21q0 22 21 22t21-22q22 0 22-21t-22-21zm104-137q66-36 66-119q0-68-45-109T152 5T48 46T3 155q0 83 66 119q4 22 21.5 36.5T131 325h42q23 0 40.5-14.5T235 274m-40-28v15q0 22-22 22h-42q-22 0-22-22v-15l-12-4q-2-1-5.5-2.5t-12.5-9T63 214t-12.5-25t-5.5-34q0-50 32-78.5T152 48t75 28.5t32 78.5q0 12-2 22.5t-5.5 18.5t-8 15t-9 11.5t-9.5 8.5t-8.5 6t-6.5 3l-3 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bug(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M491 219h-90q2-22 2-54l77-57q7-5 8-14.5T484 78q-4-7-13.5-8T454 74l-55 42q-11-48-51.5-79.5T256 5t-91.5 31.5T111 116L55 74q-6-5-15.5-4T26 78q-5 6-4 15.5t8 14.5l77 57q0 7 1 25t1 29H21q-21 0-21 21t21 21h86q4 0 8-2q1 5 11 60l-77 77q-14 14 0 30q6 6 15 6t15-6l60-60q38 109 117 109q77 0 115-109l60 60q6 6 15 6t15-6q14-16 0-30l-77-77q9-44 11-60q1 0 3.5 1t4.5 1h86q21 0 21-21t-19-21M256 432q-56 0-87-115q-6-20-18-92q38-28 105-28q71 0 102 28q-9 85-34 146t-68 61m107-256q-46-21-107-21t-107 21v-21q0-46 34-77l11 23q6 11 19 11q7 0 9-2q8-3 10.5-11.5T230 82l-12-25q15-9 38-9q21 0 36 6l-13 26q-7 19 9 28q4 4 11 4q13 0 19-13l11-23q34 34 34 79z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Building(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M85 512h278q9 0 15-6t6-15V43q0-18-12.5-30.5T341 0H43Q25 0 12.5 12.5T0 43v448q0 9 6 15t15 6zm86-43h-43v-85h43zm85 0h-43v-85h43zM43 43h298v426h-42v-85q0-17-13-30t-30-13H128q-17 0-30 13t-13 30v85H43zm42 42h43v43H85zm86 0h42v43h-42zm85 0h43v43h-43zM85 171h43v42H85zm86 0h42v42h-42zm85 0h43v42h-43zM85 256h43v43H85zm86 0h42v43h-42zm85 0h43v43h-43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BullLeft(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0Q150 0 75 62.5T0 213q0 52 29.5 96.5T107 384q-3 8-6.5 21.5t-11 41.5t-6.5 46.5T94 512q36 0 67.5-21.5T212 448t23-21h21q106 0 181-63t75-151t-75-150.5T256 0m0 384h-21q-11 0-22 5.5t-15 9.5l-17 17q-25 31-51 45q5-25 17-62l11-32l-28-17q-87-57-87-137q0-70 62.5-120T256 43t150.5 50T469 213t-62.5 120.5T256 384"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BullRight(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0Q150 0 75 62.5T0 213t75 151t181 63h21q4 0 17 13.5t28.5 29t41.5 29t54 13.5q8 0 10-13t-1.5-32t-8.5-38t-9-32l-4-13q48-30 77.5-74.5T512 213q0-88-75-150.5T256 0m126 348l-28 17l11 32q14 48 17 62q-26-14-51-45q-2-1-8-7t-8.5-7.5t-8-5.5t-9-5.5t-9-3T277 384h-23q-88 0-150.5-50.5T41 213t62.5-120T254 43t150.5 50T467 213q3 81-85 135"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Burger(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M432 237v-51q21-11 21-37q0-62-43.5-105.5T304 0H176Q114 0 70.5 43.5T27 149q0 26 21 37v49q-18 3-30.5 17.5T5 286q0 20 14 35t33 18q-4 18-4 24q0 9 4 25q-25 12-25 39v21q0 27 18 45.5T91 512h298q28 0 46-18.5t18-45.5v-21q0-27-25-39q4-16 4-25q0-8-4-22q19-3 33-18t14-35q0-19-12.5-33.5T432 237M91 192h128v43H91zm170 0h128v43H261zM176 43h128q45 0 76 31t31 75H69q0-44 31-75t76-31m213 320q0 21-21 21H112q-21 0-21-21q0-22 21-22h256q21 0 21 22m22 85q0 21-22 21H91q-22 0-22-21v-21h342zm10-149H59q-11 0-11-11t11-11h362q11 0 11 11t-11 11M155 64h21v21h-21zm64 0h21v21h-21zm64 0h21v21h-21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BusLondon(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M427 3H85Q50 3 25 28T0 88v277q0 18 12.5 30.5T43 408h2q6 27 29.5 45.5T128 472q35 0 60-25t25-60t-25-60.5t-60-25.5q-30 0-52.5 18T45 365h-2V88q0-17 12.5-30T85 45h342q17 0 29.5 13T469 88v107h-85q-21 0-21 21v43q0 21 21 21h85v85h-2q-6-27-29.5-45.5T384 301t-52.5 18t-30.5 46h-88v43h88q6 27 29.5 45.5T384 472t52.5-18t30.5-46h2q18 0 30.5-12.5T512 365V88q0-35-25-60T427 3M128 344q17 0 30 12.5t13 30.5q0 17-13 29.5T128 429t-30-12.5T85 387q0-18 13-30.5t30-12.5m256 85q-17 0-30-12.5T341 387q0-18 13-30.5t30-12.5t30 12.5t13 30.5q0 17-13 29.5T384 429m64-320q0-9-6-15t-15-6H107q-18 0-30.5 12.5T64 131v42h384zM320 259v-43q0-21-21-21H85q-21 0-21 21v43q0 9 6 15t15 6h214q9 0 15-6t6-15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 43h-43V21q0-21-21-21t-21 21v22H149V21q0-21-21-21t-21 21v22H64q-27 0-45.5 19.5T0 109v358q0 19 12.5 32T43 512h426q18 0 30.5-13t12.5-32V109q0-27-18.5-46.5T448 43m21 426H43V171h426zm0-341H43v-19q0-10 6-17t15-7h384q9 0 15 7t6 17z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CalendarGrid(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 43h-43V21q0-21-21-21t-21 21v22H149V21q0-21-21-21t-21 21v22H64q-27 0-45.5 19.5T0 109v358q0 19 12.5 32T43 512h426q18 0 30.5-13t12.5-32V109q0-27-18.5-46.5T448 43M107 469H43v-42h64zm0-85H43v-43h64zm0-85H43v-43h64zm0-86H43v-42h64zm85 256h-43v-42h43zm0-85h-43v-43h43zm0-85h-43v-43h43zm0-86h-43v-42h43zm85 256h-42v-42h42zm0-85h-42v-43h42zm0-85h-42v-43h42zm0-86h-42v-42h42zm86 256h-43v-42h43zm0-85h-43v-43h43zm0-85h-43v-43h43zm0-86h-43v-42h43zM43 128v-19q0-10 6-17t15-7h384q21 0 21 22v21zm426 43v42h-64v-42zm-64 170h64v43h-64zm64 128h-64v-42h64zm0-170h-64v-43h64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Camera(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469 45H107V24q0-21-22-21q-21 0-21 21v21H43q-18 0-30.5 13T0 88v277q0 28 18.5 46T64 429h384q27 0 45.5-18t18.5-46V88q0-17-12.5-30T469 45M341 88h43v43h-43zM43 88h256v43H43zm21 299q-21 0-21-22V173h185q-57 40-57 107t57 107zm235-22q-35 0-60.5-25T213 280t25.5-60t60.5-25t60 25t25 60t-25 60t-60 25m170 0q0 22-21 22h-79q58-41 58-107t-58-107h100zm0-234h-42V88h42zm-170 85q-28 0-46 18.5T235 280t18 45.5t46 18.5q27 0 45.5-18.5T363 280t-18.5-45.5T299 216m0 85q-22 0-22-21t22-21q21 0 21 21t-21 21m-214 0h43v43H85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Car(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M363 256q-28 0-46 18.5T299 320t18 45.5t46 18.5q27 0 45.5-18.5T427 320t-18.5-45.5T363 256m0 85q-22 0-22-21t22-21q21 0 21 21t-21 21m134-40l11-41q5-16-3.5-31.5T480 209l-98-32l-34-102q-12-30-41-30H164q-13-2-24.5 5T124 70L85 177l-59 30q-28 18-22 49l11 43q-15 7-15 21q0 21 21 21h47q15 43 60 43q27 0 45.5-18.5T192 320t-18.5-45.5T128 256q-45 0-60 43h-8l-13-54l60-27q1 0 6-5h250q1 0 2 1t2 1l98 32l-13 52h-25v42h64q21 0 21-21q0-16-15-19m-369-2q21 0 21 21t-21 21t-21-21t21-21m85-128h-79l30-86h49zm43 0V85h49l28 86zm-64 128v42h107v-42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cart(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 69H137l-3-12q-5-23-22.5-37.5T70 5H21q-9 0-15 6T0 27q0 9 6 15t15 6h49q19 0 22 17l49 256q6 21 23.5 34t38.5 13h202q10 0 16-6t6-15q0-22-22-22H203q-14 0-20-12l-2-9h214q20 0 36.5-12t22.5-31l58-123v-5q0-27-18.5-45.5T448 69m-32 175v2q-3 15-21 15H173l-28-149h303q18 0 21 17zM256 432q0 18-12.5 30.5T213 475q-17 0-29.5-12.5T171 432t12.5-30.5T213 389q18 0 30.5 12.5T256 432m171 0q0 18-12.5 30.5T384 475t-30.5-12.5T341 432t12.5-30.5T384 389t30.5 12.5T427 432"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CartSupermarket(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 69H137l-3-12q-5-23-22.5-37.5T70 5H21q-9 0-15 6T0 27q0 9 6 15t15 6h49q19 0 22 17l49 256q6 21 23.5 34t38.5 13h202q10 0 16-6t6-15q0-22-22-22H203q-14 0-20-12l-2-9h214q20 0 36.5-12t22.5-31l58-123v-5q0-27-18.5-45.5T448 69M192 261h-19l-28-149h47zm85 0h-42V112h42zm86 0h-43V112h43zm53-17v2q-3 10-11 13V112h43q18 0 21 17zM256 432q0 18-12.5 30.5T213 475q-17 0-29.5-12.5T171 432t12.5-30.5T213 389q18 0 30.5 12.5T256 432m171 0q0 18-12.5 30.5T384 475t-30.5-12.5T341 432t12.5-30.5T384 389t30.5 12.5T427 432"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chat(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 262q0 43 24.5 81T90 405q-2 7-4.5 18t-7 34.5t-3.5 39T85 512q30 0 60.5-16t48.5-32t19-16q55 0 107-21q-6-2-22.5-12T277 405h-64q-18 0-38 20q-28 25-53 36l6-77l-17-15q-68-44-68-107q0-16 6-36q-4-6-5.5-18.5T42 185v-23l1-13Q0 195 0 262M299 0q-89 0-151.5 52T85 177q0 72 62 118t152 46q1 0 20.5 21.5t51.5 43t62 21.5q7 0 8.5-11t-1.5-26.5t-7-31.5t-7-27l-4-11q41-25 65.5-62.5T512 177q0-73-62.5-125T299 0m102 284l-28 17l11 32q2 5 5 17t6 19q-22-15-52-45q-23-25-42-25q-70 0-120.5-32.5T130 177q-1-56 48.5-95T299 43t120.5 39t49.5 95q0 63-68 107"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ChatAlt(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M463 226q6 20 6 36q0 63-68 107l-17 15q4 64 9 77q-12-6-54-36q-18-20-38-20h-64q-9 9-43 22q53 21 107 21q1 0 19 16t48.5 32t60.5 16q7 0 9-11t0-26.5t-4.5-31.5t-5.5-27l-3-11q89-56 89-143q0-67-43-113q-1 7-.5 20t-1 31.5T463 226M79 427q30 0 62-21.5t51.5-43T213 341q90 0 152-46t62-118q0-73-62.5-125T213 0Q125 0 62.5 52T0 177q0 86 90 143q-2 7-5.5 18T75 372.5t-5.5 39T79 427M43 177q0-56 49.5-95T213 43t121 39t50 95t-51 89t-120 33q-22 0-42 25q-35 35-52 45q1-2 11-36l11-32l-28-17q-70-44-70-107"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Check(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M405 491V192l-21-43l-21 43v299q0 21 21 21t21-21M491 0H21Q12 0 6 6T0 21v320h320v-42H43V107h426v192h-21v42h64V21q0-9-6-15t-15-6m-22 64H43V43h426zM85 149h235v43H85zm0 64h214v43H85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckBox(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M45 145v-38q0-8 7-15t15-7h170q5 0 9 2l34-29q-22-15-43-15H67q-28 0-46 18T3 107v170q0 26 18.5 45T67 341h170q27 0 45.5-19t18.5-45v-70l-42 38v32q0 8-7 15t-15 7H67q-8 0-15-7t-7-15zm314-43l-30-34l-28 28l-42 36l-81 71l-52-66l-34 25l77 105l90-79l42-37z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckBoxEmpty(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M237 341q28 0 46-18t18-46V107q0-28-18-46t-46-18H67q-28 0-46 18T3 107v170q0 28 18 46t46 18zM45 277V107q0-22 22-22h170q22 0 22 22v170q0 22-22 22H67q-22 0-22-22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Checked(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m189 395l263-290q13-13 13-30t-13-30l-34-34Q407 0 388 0q-20 0-30 11L189 183l-76-78Q97 85 83 85q-13 0-30 17l-34 37Q6 152 6 169.5T19 198zM83 139l106 106L386 41l34 36l-231 254L49 171z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cinch(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M468 98q0 30-25.5 54.5T377 188l-43 44v-37h-10q-59 0-101.5-28.5T180 98t42.5-68T324 2t101.5 28T468 98m-323 93q-57 0-97.5 40T7 327t40.5 95.5T145 462h65v-72h-40q-1-1-7.5 0t-17.5-2.5t-21.5-8.5t-18.5-19t-8-33q0-31 18.5-46.5T152 264h58v-72h-1l-1-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clock(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0Q150 0 75 75T0 256t75 181t181 75t181-75t75-181t-75-181T256 0m192 277h19q-8 76-61.5 130T277 467v-19q0-21-21-21t-21 21v19q-76-8-130-61.5T45 277h19q21 0 21-21t-21-21H45q6-76 60-130t130-60v19q0 21 21 21t21-21V45q76 8 130 61.5T467 235h-19q-21 0-21 21t21 21m-169-23l39-96q3-9-.5-17T305 130q-20-7-28 13l-47 117l92 137q5 8 17 8q5 0 13-4q19-10 6-30z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClothesWater(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M346 375h-52l-29-110q0-3-7-26l-4-26l-4 26q-2 8-7 26l-28 110h-49l-53-213h45l28 117q10 52 10 56q0-8 5-26q5-23 6-27l28-120h42l32 120q1 3 2 10t3 13q4 22 4 28q0-6 4-28q4-20 7-28l27-117h45zM256 0Q150 0 75 75T0 256q0 103 71.5 179.5T243 512h13q116-7 192.5-83T512 256q-15-110-84.5-183T256 0m175 384q-31 37-78 59.5T254 469h-11q-82 0-141-63.5T43 256q0-88 62.5-150.5T256 43q76 0 138 56.5T469 262q10 64-38 122"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloud(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M427 171v-11q0-65-47.5-112.5T267 0q-57 0-101 35.5T111 126q-48 6-79.5 42.5T0 254q0 53 38 91.5t90 38.5h299q37-8 61-38t24-69q0-38-24-68.5T427 171m0 170H128q-35 0-60-26t-25-61q0-31 25.5-54.5T128 171q4 0 11-11l10-32q8-39 41.5-62T267 43q48 0 82.5 34.5T384 160v32q0 9 6.5 15t14.5 6h13q22 5 36.5 23t14.5 41t-11 41.5t-31 22.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloudapp(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M100 358h-1h246q49 0 83-34.5t34-83.5t-34-83.5t-83-34.5h-7q-8-41-40-68.5T223 26q-49 0-83.5 34.5T105 144q0 1 .5 2.5t.5 2.5h-1q-43 0-73 30.5T2 254q0 41 28.5 71.5T100 358"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ClubsCard(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 512h256q27 0 45.5-18.5T384 448V64q0-27-18.5-45.5T320 0H64Q37 0 18.5 18.5T0 64v384q0 27 18.5 45.5T64 512M43 64q0-21 21-21h256q21 0 21 21v384q0 21-21 21H64q-21 0-21-21zm106 235q12 0 22-7v49h-22v22h86v-22h-22v-49q10 7 22 7q17 0 29.5-13t12.5-30t-12.5-30t-29.5-13q0-17-13-29.5T192 171t-30 12.5t-13 29.5q-17 0-29.5 13T107 256t12.5 30t29.5 13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cocktail(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M373 73q18-25 5-49q-11-21-37-21H41Q19 3 4 24Q-7 52 9 71l162 175v119q0 4-5 9t-20.5 9t-38.5 4q-22 0-22 21t22 21h170q22 0 22-21t-22-21q-23 0-38.5-4t-20.5-9t-5-9V246zM192 205L83 88h220zM341 45l-21 24v-2H64L45 45z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Code(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M312 3q-8-4-16 .5T286 18L169 404q-4 18 15 25q8 4 16-.5t10-14.5L325 28q7-18-13-25m45 320q14 14 30 0l106-107l-106-107q-16-14-30 0q-16 16 0 30l77 77l-77 77q-12 14 0 30m-250 0q16 14 30 0q16-16 0-30l-75-77l77-77q16-14 0-30q-14-14-30 0L3 216z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Coffee(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M301 21H45q-17 0-29.5 13T3 64v149q0 63 43.5 106.5T152 363h43q55 0 97-37t50-91h2q35 0 60-25.5t25-60.5t-25-60t-60-25q0-17-12.5-30T301 21m0 192q0 45-31 76t-75 31h-43q-45 0-76-31t-31-76V64h256zm86-64q0 18-13 30.5T344 192v-85q17 0 30 12.5t13 29.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CoffeeHot(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 213v150q0 62 43.5 105.5T149 512h43q56 0 97.5-36.5T339 384h2q35 0 60.5-25t25.5-60t-25.5-60.5T341 213q0-17-12.5-29.5T299 171H43q-18 0-30.5 12.5T0 213m341 43q18 0 30.5 12.5T384 299q0 17-12.5 29.5T341 341zm-42 107q0 44-31 75t-76 31h-43q-44 0-75-31t-31-75V213h256zM145 34q10-19-6-30q-17-11-28 5l-11 14Q82 52 85.5 85.5T113 143q6 6 15 6t15-6q13-15 0-30q-13-13-15-31.5t9-34.5zm75 53l8-8q14-14 2-30q-4-6-13.5-6.5T201 47l-9 8q-22 19-19 45q3 28 32 45q2 4 8 4q12 0 20-10q11-19-9-30q-11-7-11-11z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Coin(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0Q150 0 75 75T0 256t75 181t181 75t181-75t75-181t-75-181T256 0m0 469q-88 0-150.5-62.5T43 256t62.5-150.5T256 43t150.5 62.5T469 256t-62.5 150.5T256 469m0-405q-80 0-136 56T64 256t56 136t136 56t136-56t56-136t-56-136t-136-56m0 341q-62 0-105.5-43.5T107 256t43.5-105.5T256 107t105.5 43.5T405 256t-43.5 105.5T256 405m43-106h-22V171h-42l-30 23q-18 12-11 28q3 8 11.5 10.5T222 230l13-6v75h-22q-21 0-21 21t21 21h86q21 0 21-21t-21-21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Coins(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 88q-15 0-23 2q-23-40-62-63.5T171 3Q100 3 50 53T0 173q0 71 50 121t121 50q15 0 23-2q23 40 62 63.5t85 23.5q71 0 121-50t50-120q0-71-50-121T341 88M198 297q-16 4-27 4q-53 0-90.5-37.5T43 173t35-90.5T166 45q40 0 75.5 22t35.5 55q8 5 12 11.5t4 12.5t.5 13t.5 14l3 2q0 61-50 99q-20 17-49 23m141-102q27 0 46.5 18.5T405 259t-18 45.5t-46 18.5q-33 0-51-28q43-40 49-100m2 192q-66 0-104-56q3-1 9.5-4t9.5-4q31 44 85 44q45 0 76-30.5t31-75.5t-32-76t-77-31q0-9-4-21h6q53 0 90.5 37.5T469 261q0 52-37.5 89T341 387m-64-252h-6q-13-32-40-50t-60-18q-45 0-76 31t-31 75q0 45 31 76t76 31q17 0 25-2v2q83-37 83-105q2-4 0-20t-2-20m-89 100q-5 2-17 2q-28 0-46-18t-18-46q0-27 18-45.5t46-18.5q19 0 36.5 11t25.5 30h2q0 4 1 11.5t1 11.5q0 22-14 39.5T188 235"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compass(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0Q150 0 75 75T0 256t75 181t181 75t181-75t75-181t-75-181T256 0m0 469q-88 0-150.5-62.5T43 256t62.5-150.5T256 43t150.5 62.5T469 256t-62.5 150.5T256 469m-85-245v203l162-122q8-5 8-17V85L179 207q-8 5-8 17m42 19q0-9 7-13l91-85l-12 124q0 9-7 13l-91 85zm64 13q0 9-6 15t-15 6t-15-6t-6-15t6-15t15-6t15 6t6 15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Contact(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469 0h-64v4q-12-4-21-4H85Q68 0 55.5 12.5T43 43v64H21q-21 0-21 21t21 21h22v86H21q-21 0-21 21t21 21h22v86H21q-21 0-21 21t21 21h22v64q0 18 12.5 30.5T85 512h299q27 0 45.5-18.5T448 448v-85q27 0 45.5-18.5T512 299v-86h-4q4-12 4-21v-85h-4q4-14 4-22V43q0-18-12.5-30.5T469 0m-21 43h21v42q0 22-21 22zm0 106q9 0 21-4v47q0 21-21 21zm-43 299q0 21-21 21H85v-64h22q21 0 21-21t-21-21H85v-86h22q21 0 21-21t-21-21H85v-86h22q21 0 21-21t-21-21H85V43h299q21 0 21 21zm43-128v-64q9 0 21-4v47q0 21-21 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Contrast(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0Q150 0 75 75T0 256t75 181t181 75t181-75t75-181t-75-181T256 0m0 469q-88 0-150.5-62.5T43 256t62.5-150.5T256 43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cookie(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M216 3Q128 3 65.5 65.5T3 216t62.5 150.5T216 429t150.5-62.5T429 216T366.5 65.5T216 3m0 384q-70 0-120.5-50.5T45 216T95.5 95.5T216 45t120.5 50.5T387 216t-50.5 120.5T216 387m21-299h-21q-21 0-21 21q0 22 21 22h21q22 0 22-22q0-9-6-15t-16-6m-91 92q-16-16-30 0q-16 14 0 30l15 15q14 14 30 0q14-16 0-30zm85 91q-15-13-30 0q-14 16 0 30l15 15q15 15 30 0q14-14 0-30zm70-91q-14-16-30 0q-14 14 0 30l15 15q16 14 30 0q16-16 0-30z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copy(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M155 304H48V176h64q27 0 45.5-19t18.5-45V48h64v85l43-42V27q0-8-7-15t-15-7H114L5 112v213q0 8 7 15t15 7h128zM133 48v64q0 8-6.5 14.5T112 133H48zm320 85H306l-23 22l-43 42l-43 43v213q0 8 7 15t15 7h234q8 0 15-7t7-15V155q0-8-7-15t-15-7m-170 86l42-43v64q0 8-6.5 14.5T304 261h-64zm149 213H240V304h64q27 0 45.5-19t18.5-45v-64h64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Coroflot(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M164 303q0 8-6.5 14.5T143 324q-9 0-15-6.5t-6-14.5q0-9 6-15t15-6q8 0 14.5 6t6.5 15m233-23q0 46-32.5 78.5T286 391h-3q-15 32-45 51.5T171 462q-50 0-86-35.5T49 340q0-40 23.5-72t59.5-44q1-61-18-99T67 89q-16 1-26.5 1.5t-20-1t-14-5.5T3 74q2-9 18.5-9.5T40 52q1-9 0-16t-2-12t2-10T52 4q17-8 23.5 7.5T80 53q30-8 60.5 3T195 91t38.5 42.5T260 171q24-3 26-3q46 0 78.5 33t32.5 79m-237-62q4-1 11-1t21 2q17-28 51-42Q169 62 96 71q59 45 64 147m108 122q0-32-18.5-57T202 249q-13-4-22-5h-9q-40 0-68 28t-28 68t28 68t68 28q57 0 84-49l9-21q4-16 4-26m107-60q0-37-26-63.5T286 190q-44 0-71 35q35 14 57 45t22 70q0 17-4 29q36-1 60.5-27t24.5-62"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Couple(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 341v43h214v-43q0-38-24.5-67T133 237v-4q28-7 46-30.5t18-53.5v-42q0-35-25-60.5T112 21T52 46.5T27 107v42q0 30 18 53t46 31v4q-37 8-61.5 37T5 341m171 0H48q0-27 18.5-45.5T112 277t45.5 18.5T176 341M69 149v-42q0-18 13-30.5T112 64t30 12.5t13 30.5v42q0 18-13 30.5T112 192t-30-12.5T69 149m192 192v43h214v-43q0-38-24.5-67T389 237v-4q28-7 46-30.5t18-53.5v-42q0-35-25-60.5T368 21t-60 25.5t-25 60.5v42q0 30 18 53t46 31v4q-37 8-61.5 37T261 341m171 0H304q0-27 18.5-45.5T368 277t45.5 18.5T432 341M325 149v-42q0-18 13-30.5T368 64t30 12.5t13 30.5v42q0 18-13 30.5T368 192t-30-12.5t-13-30.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cpu(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M483 163q10 0 10-11t-10-11h-43v-21q0-27-18.5-45.5T376 56h-21V13q0-10-11-10t-11 10v43h-42V13q0-10-11-10t-11 10v43h-42V13q0-10-11-10t-11 10v43h-42V13q0-10-11-10t-11 10v43h-21q-27 0-45.5 18.5T56 120v21H13q-10 0-10 11t10 11h43v42H13q-10 0-10 11t10 11h43v42H13q-10 0-10 11t10 11h43v42H13q-10 0-10 11t10 11h43v21q0 27 18.5 45.5T120 440h21v43q0 10 11 10t11-10v-43h42v43q0 10 11 10t11-10v-43h42v43q0 10 11 10t11-10v-43h42v43q0 10 11 10t11-10v-43h21q27 0 45.5-18.5T440 376v-21h43q10 0 10-11t-10-11h-43v-42h43q10 0 10-11t-10-11h-43v-42h43q10 0 10-11t-10-11h-43v-42zm-86 213q0 21-21 21H120q-21 0-21-21V120q0-21 21-21h256q21 0 21 21zm-64-235H163q-22 0-22 22v170q0 22 22 22h170q22 0 22-22V163q0-22-22-22m-21 171H184V184h128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreativeCommons(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M234 232q0-33 19-52t48-19q43 0 61 33l-30 16q-8-19-26-19q-31 0-31 41t31 41q20 0 29-20l28 14q-20 36-60 36q-31 0-50-18.5T234 232m-64 71q40 0 60-36l-29-14q-8 20-28 20q-31 0-31-41t31-41q18 0 26 19l30-16q-18-33-61-33q-29 0-48 19t-19 52q0 34 19 52.5t50 18.5M69 394Q2 327 2 232T70 69Q135 2 232 2t164 67q66 66 66 163t-65 162q-71 68-165 68q-95 0-163-68M43 232q0 76 57 133q55 55 132 55q79 0 135-56q54-54 54-132q0-80-55-133q-56-56-134-56q-79 0-132 55q-57 59-57 134"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCard(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 0H64Q37 0 18.5 18.5T0 64v256q0 27 18.5 45.5T64 384h384q27 0 45.5-18.5T512 320V64q0-27-18.5-45.5T448 0m21 320q0 21-21 21H64q-21 0-21-21v-21h426zm0-64H43v-21h426zm0-64H43V64q0-21 21-21h384q21 0 21 21zM405 85h-85q-21 0-21 22q0 21 21 21h85q10 0 16-6t6-15q0-22-22-22m-128 22q0 8-6 14.5t-15 6.5t-15-6.5t-6-14.5q0-9 6-15.5t15-6.5t15 6.5t6 15.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crop(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M427 85H128V0H85v85H0v43h384v256H128V171H85v256h427v-43h-85zm-43 384h43v43h-43zM213 213h86v86h-86z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crown(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M48 475h384q17 0 30-13t13-30V5h-22q-35 0-60 25.5T368 91q0 59 58 81l-94 79l-71-122q18-6 30.5-22.5T304 69q0-27-18.5-45.5T240 5t-45.5 18.5T176 69q0 21 12 37.5t31 22.5l-69 122l-96-79q58-22 58-81q0-35-25-60.5T27 5H5v427q0 17 13 30t30 13M411 91q0-26 21-37v75q-21-13-21-38M240 48q21 0 21 21q0 9-6 15.5T240 91q-8 0-14.5-6.5T219 69q0-21 21-21M48 52q21 11 21 37q0 28-21 38zm0 169l113 94l79-139l79 139l113-94v147H48zm0 190h384v21H48z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cutlery(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M216 0h-21v459q0 23 15.5 38t37.5 15t37.5-15t15.5-38V282q22-19 22-47V107q0-45-31-76T216 0m43 459q0 10-11 10t-11-10V299h22zm21-224q0 9-6 15t-15 6h-22V47q19 6 31 22.5t12 37.5zM3 288v171q0 23 15.5 38T56 512t37.5-15t15.5-38V288q0-22-15.5-37.5T56 235t-37.5 15.5T3 288m42 0q0-11 11-11t11 11v171q0 10-11 10t-11-10zm43-139H67V0H45v149H24V0H3v171h42v64h22v-64h42V0H88z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DaftPunk(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M491 213v-27q21-12 21-37t-23-38q-5-38-35-64t-70-26t-70 26t-35 64q-23 13-23 38t21 37v27q0 18 13 30.5t30 12.5h43q-45 0-76 31t-31 76q0-45-31-76t-76-31h11l43-64h32v-6q21-12 21-37t-23-38q-5-38-35-64t-70-26t-70 26t-35 64Q0 124 0 149q0 18 12.5 30.5T43 192h10l43 64h11q-45 0-76 31T0 363v42h512v-42q0-45-31-76t-76-31h43q17 0 30-12.5t13-30.5M77 149l-11-21h124l-11 21zm51-85q11 0 15 2q-36 13-36 41H68q15-43 60-43M90 171h17q6 34 32 40v2h-22zm59 128q8 0 22 4v17H85v-17q14-4 22-4zM43 363q0-31 21-47v47zm149 0v-47q21 16 21 47zm107 0q0-31 21-47v47zm170 0h-21v-47q21 16 21 47m-42-60v17h-86v-17q14-4 22-4h42q8 0 22 4m0-154h21v22h-21zm21-21h-21V81q21 19 21 47M341 81v47h-21q0-28 21-47m-21 68h21v22h-21zm0 64v-21h21v21zm43 0V68q12-4 21-4t21 4v145zm64-21h21v21h-21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dailybooth(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M232 462q83 0 145-51l81 29l-34-81q38-59 38-127q0-95-67.5-162.5T232 2T69.5 69.5T2 232t67.5 162.5T232 462M105 172l47-11q24-44 73-44q29 0 48 15l47-11q8-1 15.5 4t9.5 14l31 134q5 22-14 27l-216 50q-19 5-25-17L90 199q-5-21 15-27m132 120q24 0 41-17t17-41t-17-41t-41-17t-41 17t-17 41t17 41t41 17m0-95q15 0 26 10.5t11 26.5q0 15-11 26t-26 11q-16 0-26.5-11T200 234q0-16 10.5-26.5T237 197"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dashboard(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m339 66l22-37Q305 8 256 8Q154 8 81 74T0 243q0 8 6 14.5t13 6.5h4q8 0 14-5.5t6-13.5q7-87 67.5-140.5T256 51q49 0 83 15m98 83q27 45 30 96q0 7 5.5 13t13.5 6h5q8 0 14-7t5-14q-6-75-49-135zM448 8q-17-9-30 6l-21 35l-22 36l-96 162q-12-4-23-4q-35 0-60 25t-25 60t25 60t60 25t60-25t25-60q0-38-25-60l94-155l21-37l23-38q13-20-6-30M256 371q-17 0-30-13t-13-30t13-30t30-13t30 13t13 30t-13 30t-30 13"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DataBoard(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M491 320h-24l-17-196q-1-25-19.5-42.5T388 64H277V21q0-21-21-21t-21 21v43H124q-24 0-42.5 17.5T62 124L45 320H21q-9 0-15 6t-6 15q0 10 6 16t15 6h79L64 484q-3 8 2 16t13 10h6q13 0 22-15l38-132h90v85q0 21 21 21t21-21v-85h90l38 134q3 15 22 15h6q19-7 15-26l-36-123h79q9 0 15-6t6-16q0-9-6-15t-15-6m-404 0l18-194q0-8 6.5-13.5T126 107h262q9 0 15.5 5.5T410 126l17 194zm237-183q-49 65-53 70q-35-24-47-32l-15-13l-77 103q-5 6-3.5 15.5T137 294q7 5 12 5q12 0 17-9l52-68q38 26 47.5 31t16.5 3l8-2l4-7q3-4 13.5-17.5T333 195t23-31q5-6 4-15.5t-8-14.5q-5-6-14-5t-14 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Delete(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M237 0h-42q-28 0-46 18.5T131 64v21H45q-17 0-29.5 13T3 128v21q0 18 12.5 30.5T45 192h2l39 284q1 16 13.5 26t29.5 10h174q17 0 29.5-10t13.5-26l39-284h2q17 0 29.5-12.5T429 149v-21q0-17-12.5-30T387 85h-86V64q0-27-18-45.5T237 0m-64 64q0-21 22-21h42q22 0 22 21v21h-86zm-44 405l-3-21h182l-2 21zm183-64H120L90 192h250zm75-277v21H45v-21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Delicious(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M444 142q-26-62-83-101q-30-19-60-29q-31-10-69-10q-48 0-90 18q-62 26-101 83q-19 31-29 61T2 232q0 48 18 90q26 62 83 101q31 19 61 29t68 10q48 0 90-18q62-26 101-83q19-30 29-60q10-31 10-69q0-48-18-90m-28 168q-26 56-73 87q-23 16-52 25q-27 9-59 9V232H33q0-42 15-78q26-56 73-87q23-16 52-25q27-9 59-9v199h199q0 42-15 78"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Designbump(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m258 278l-18 16V130h81L182 2L43 130h80v40l-17-12L5 252h51v210h92V252h57l-67-64v-79H82l100-92l100 92h-57v202l-66 59h56v92h92v-92h51z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Designfloat(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M263 118V73q46 11 77.5 42.5T383 196h-45q-17-59-75-78m-67 0V73Q95 93 73 196h45q15-57 78-78m-78 140H73q10 57 39.5 86t83.5 40v-45q-26-8-48-31t-30-50m313-29q0 85-59.5 144.5T228 433q-85 0-144.5-59.5T24 229T83.5 84T228 24q84 0 143.5 60T431 229m-20 0q0-76-53.5-130T228 45T98 98.5T44 229q0 76 54 130t130 54t129.5-54T411 229m-183 93q-39 0-66-27.5T135 229q0-39 27-66t66-28q38 1 65 28t28 66q-1 38-28 65.5T228 322m72-93q0-30-21-51.5T228 156t-51.5 21.5T155 229t21.5 51t51.5 21t51-21t21-51m-37 110v45q100-23 120-126h-45q-10 28-29.5 50.5T263 339m175 9q-10 0-16.5 6.5T415 371q0 9 6.5 15.5T438 393q9 0 15.5-6.5T460 371q0-10-6.5-16.5T438 348m-55 58q-12 0-20.5 8.5T354 435t8.5 20.5T383 464t20.5-8.5T412 435t-8.5-20.5T383 406m48-132q-14 0-14 14q0 13 14 13q13 0 13-13q0-14-13-14M31 428q-14 0-14 14q0 15 14 15q6 0 10.5-4.5T46 442q0-14-15-14m-4-77q-14 0-14 14t14 14q5 0 9.5-4t4.5-10t-4.5-10t-9.5-4m45 35q-8 0-13.5 6T53 405t5.5 13t13.5 6t13.5-6t5.5-13t-5.5-13t-13.5-6m366-282q-10 0-10 10q0 9 10 9t10-9q0-10-10-10M16 63Q0 63 0 80q0 16 16 16t16-16q0-17-16-17M53 3q-8 0-14 6t-6 15q0 8 6 14t14 6q9 0 15-6t6-14q0-9-6-15T53 3M21 130q-10 0-10 11q0 10 10 10q11 0 11-10q0-11-11-11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Designmoo(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M329 298q-3 0-6 4q-6 8-24 27q-33 31-45 31h-2q-9-3-7-13q0-7 4-21q7-21 16.5-44t21.5-50t19-43q28-62 32-72q6-18 6-29q0-23-23-28h-2q-1 0-4 14q0 2-6 22q-7 22-43 96l-5 9v-4q-2-12-14-20t-25-8q-29 0-63 32q-33 33-48 79q-9 29-11 53.5t9.5 44T148 397q25 0 52-34l3-3v3q0 13 7 23.5t20 10.5q27 0 77-57q16-18 23-30q3-4 3-7q0-5-4-5m-122 21q-13 22-30 37q-9 6-15 6q-19 0-17-28q2-62 48-117q20-22 32-22q18 0 18 21q-1 46-36 103M189 47l25-7q7-2 9-9l6-25q2-8 4 0l6 25q2 7 9 9l25 7q6 1 0 3l-25 7q-7 2-9 8l-6 26q-2 8-4 0l-6-26q-2-6-9-8l-25-7q-6-2 0-3m-74 183l-15-61q-5-18-23-23L9 128q-5-1-6.5-2.5t0-3.5t6.5-3l68-18q17-3 23-23l17-68q5-17 10 0l18 68q4 19 22 23l68 18q5 1 6.5 3t0 3.5t-6.5 2.5l-68 18q-19 6-22 23l-9 31q-10 14-21 30m303 18l40 11q9 1 3 4q-1 1-3 1l-40 11q-10 2-14 13l-10 41q-3 7-6 0l-10-41q-4-11-14-13l-40-11q-9-1-3-4q1-1 3-1l40-11q11-4 14-14l10-40q3-9 6 0l10 40q4 10 14 14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Deviantart(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M157 109h-4l-22-45l-65 8l26 55q-39 16-61 41.5t-26 51T2.5 268t7.5 38l5 14l212-61l-58-119q53-12 95.5-5t69.5 22l26 15l-72 19l-20-39q-35-5-66 0l49 100l211-62q-1-2-3-4.5t-9-10.5t-16-15t-24-16.5t-31.5-17t-40.5-14t-50-9.5t-60-1.5t-71 7.5M72 250q-2-5-4-14t7.5-31t35.5-39l32 65z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiamondsCard(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 512h256q27 0 45.5-18.5T384 448V64q0-27-18.5-45.5T320 0H64Q37 0 18.5 18.5T0 64v384q0 27 18.5 45.5T64 512M43 64q0-21 21-21h256q21 0 21 21v384q0 21-21 21H64q-21 0-21-21zm213 192l-64-85l-64 85l64 85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Digg(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M84 28v82H2v164h128V28zM48 233v-82h36v82zm92-123h46v164h-46zm0-82h46v46h-46zm189 82H206v164h77v36h-77v46h123zm-77 123v-82h31v82zm210-123H340v164h81v36h-81v46h122zm-41 123h-35v-82h35z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DiggAlt(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M462 93h-59V3H231v89h-59v116H32V92H2v117h30v31h54v54H63v32h77v55H63v-55H32v87h139v-87h-31v-32h-23v-54h87v172h232V233h26zm-91-59v114H262V34zm-36 347v-87h-28v87h-71V208h-33v-83h28v55h172v-55h27v84h-25v172zM140 92V62h32v30zm0-30H63V30h77zM63 92H32V62h31z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Diigo(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M134 177q27 0 50.5-10.5t36-20.5t31.5-29v345h-57q-1-14-1-42v-97q0-18 .5-54t.5-54q-4 1-14 4.5t-14 4.5q-55 14-95-10v248H14V117q27 29 54 45t66 15m91-100q15 0 26-11t11-27t-11-27t-26-11q-16 0-26.5 11T188 39t10.5 27T225 77M43 77q16 0 26.5-11T80 39T69.5 12T43 1Q28 1 17 12T6 39t11 27t26 11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Disabled(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M475 405h-5l-29-102q-7-21-24.5-34T379 256h-77q-45 0-60-43h105q21 0 21-21t-21-21H227l-15-43h5q27 0 45.5-18.5T281 64t-18.5-45.5T217 0q-28 0-46 18.5T153 64q0 3 1 8.5t1 8.5v11l44 134q11 32 39 52.5t62 20.5h79q14 0 21 15l38 134h37q9 0 15-6t6-15q0-10-6-16t-15-6M217 43q21 0 21 21t-21 21q-22 0-22-21q0-9 6.5-15t15.5-6m-20 469q72 0 127.5-48.5T387 343l-42-4q-7 56-49.5 93T197 469q-62 0-105.5-43.5T48 320q0-46 25-83.5t69-55.5l-17-38q-55 21-87.5 69.5T5 320q0 80 56 136t136 56"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoNotBleach(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m240 126l21 23l32-21l-36-53q-6-10-18-10t-18 10l-34 53l32 21zM101 341l66-100l-29-34L7 414q-7 10 0 21q1 0 5 7zm54 64l-32 43h234l-32-43zm315 11L340 209l-30 32l77 100l79 103q4-4 4-7q9-11 0-21m-2 90q15-15 2-30l-27-28l-39-43l-136-149l32-34l30-32L468 36q15-13-2-30q-3-6-13-6q-7 0-15 6L306 151l-30 32l-36 41l-36-41l-30-32L42 6q-8-6-15-6q-9 0-15 6q-15 15-2 30l138 154l30 32l32 34L74 405l-39 43l-25 28q-15 13 2 30q6 6 15 6q8 0 15-6l53-58l38-43l107-117l107 117l38 43l51 58q8 6 15 6q11 0 17-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoNotDry(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M471 32q2-2 4-2q-11-30-40-30H93Q64 0 53 28q1 0 2 1t2 1l36 32V43h342v19zm-36 309l42 34V137l-42 34zM57 480q-1 0-2 1t-2 1q11 30 40 30h342q29 0 40-30q-1 0-2-1t-2-1l-36-30v19H93v-19zm36-309l-42-34v238l42-34zM17 431q-7 5-8 14t5 16q4 8 15 8q9 0 13-4l9-6l42-34l171-139l171 139l42 34l9 6q4 4 13 4q10 0 17-8q5-7 4.5-16.5T514 431l-37-28l-42-34l-137-113l137-113l42-34l34-28q7-5 8-14t-5-16q-4-8-15-8q-9 0-13 4l-9 6l-42 34l-171 139L93 90L51 53l-9-6q-4-4-13-4q-10 0-17 8q-11 17 2 30l37 28l42 34l137 113L93 369l-42 34z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoNotIron(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m386 131l83 307q3 15 22 15h6q8-1 12.5-9t2.5-16l-85-310q-11-40-45.5-65.5T305 27H192q-21 0-21 21t21 21h113q27 0 50 17.5t31 44.5M21 453h320q22 0 22-21t-22-21H45q6-52 36-93.5t77-62.5q19-7 11-28q-10-19-28-10q-65 29-103 87T0 432q0 21 21 21m278-256q-22 0-22 22q0 9 6 15t16 6h42q10 0 16-6t6-15q0-22-22-22zM28 12q-16 14 0 30l426 448q8 6 15 6q9 0 15-6q16-16 0-30L58 12Q42-4 28 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoNotWash(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M66 359q3 16 15 30l25-38v-2L72 185l6 6q28 28 58 28v-43q-13 0-28-15q-20-22-47-28L51 86q-1-9-9-14t-17-3q-9 1-14 9T8 95zm209 9h-22l-30 43h82zm117-192q13 0 28 15q17 17 32 23l-30 137v2l25 38q11-10 15-29l58-267q2-9-3-17t-14-9q-9-2-17 3t-9 14l-17 84q-4-6-10-9q-1-1-8.5-9t-14-12.5T413 133zm-218-26l26 35l38 55l-89 128l-30 43l-21 32q-12 17 4 30q2 0 6 1t7 1q6 0 17-9l38-55l30-43l64-92l64 92l30 43l38 55q6 9 17 9q6 0 13-5q8-5 9.5-13t-5.5-16l-21-32l-30-43l-89-128l15-21l53-77l72-103q12-17-4-30q-2 0-6-1t-7-1q-11 0-17 9L287 167l-23 37l-49-71l-83-119q-6-9-17-9q-6 0-13 5t-8 13t4 16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoNotWring(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m260 149l30-38q-22-4-34-4q-18 0-34 6l30 36zm-4 207l-21 26q6 2 21 2t21-2zM73 380l100-124L51 107H0v277h43q8 0 30-4M43 149q22 0 40 13t26 34l17 49l-17 49q-8 21-26 34t-40 13zm416-40L339 256l103 124q16 4 29 4h43V107h-43q-8 0-12 2m10 232q-22 0-40-13t-26-34l-17-49l17-49q8-21 26-34t40-13zM356 166L465 34q5-6 4.5-15.5T463 4q-8-4-15-4q-12 0-17 9L329 134l-28 35l-45 53l-45-56l-28-34L81 9q-3-5-4-5q-1-1-9-4h-4q-9 0-13 4q-6 5-7 14.5T49 34l109 132l30 39l42 51l-38 45l-30 34L47 478q-11 15 2 30q8 4 15 4q12 0 17-9l113-138l28-34l34-43l34 43l28 34l113 138q7 9 17 9q9 0 13-4q6-5 7-14.5t-5-15.5L348 335l-28-34l-36-45l42-51z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DollarBill(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469 21H43q-18 0-30.5 13T0 64v256q0 17 12.5 30T43 363h426q18 0 30.5-13t12.5-30V64q0-17-12.5-30T469 21M43 64h42q0 17-12.5 30T43 107zm0 256v-43q17 0 29.5 13T85 320zm234-9v-76q0-22-21-22t-21 22v76q-22-17-22-55q0-26 13-45t30-19t30 19t13 45q0 35-22 55m192 9h-42q0-17 12.5-30t29.5-13zm0-85q-35 0-60 25t-25 60h-60q17-26 17-64q0-45-25-76t-60-31t-60 31t-25 76q0 38 17 64h-60q0-35-25-60t-60-25v-86q35 0 60-25t25-60h256q0 35 25 60t60 25zm0-128q-17 0-29.5-13T427 64h42zM341 85H171q-22 0-22 22q0 9 6 15t16 6h170q10 0 16-6t6-15q0-22-22-22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dollars(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M491 0H21Q12 0 6 6T0 21v256q0 10 6 16t15 6h470q9 0 15-6t6-16V21q0-9-6-15t-15-6M43 43h21q0 9-6 15t-15 6zm0 213v-21q21 0 21 21zm138 0h-74q0-27-18.5-45.5T43 192v-85q27 0 45.5-18.5T107 43h74q-32 45-32 106q0 60 32 107m75 0q-26 0-45-32t-19-75t19-74.5T256 43t45 31.5t19 74.5q0 44-18.5 75.5T256 256m213 0h-21q0-21 21-21zm0-64q-27 0-45.5 18.5T405 256h-72q32-47 32-107q0-61-32-106h72q0 27 18.5 45.5T469 107zm0-128q-9 0-15-6t-6-15h21zm0 299H350q13-18 13-43h-43q0 2-.5 5t-4 10.5t-9.5 13t-19 10t-31 4.5q-27 0-43-10.5T194.5 334t-2.5-14h-43q0 19 13 43H43v-43H0v43q0 17 12.5 29.5T43 405h426q18 0 30.5-12.5T512 363v-43h-43zm0 106H350q13-18 13-42h-43q0 6-2 13t-17.5 18t-44.5 11q-27 0-43-10t-18.5-18t-2.5-14h-43q0 18 13 42H43v-42H0v42q0 18 12.5 30.5T43 512h426q18 0 30.5-12.5T512 469v-42h-43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoubleArrow(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M165 10q-2-3-7-3h-19L11 95q-6 3-7.5 12T7 123q7 10 17 10q9 0 13-4l94-64v348l-94-64q-19-12-30 6q-11 20 6 30l128 85q2 0 2 3h22q2 0 2-3l128-85q19-10 6-30q-11-18-30-6l-93 64V67l93 64q5 2 9 2q9 0 17-8q11-20-6-30z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Down(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m192 284l186-162q15-17 2-30q-17-17-30-2L192 228L36 90Q23 75 6 92q-14 14 3 30z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownArrowCircle(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 512q106 0 181-75t75-181t-75-181T256 0T75 75T0 256t75 181t181 75m0-469q88 0 150.5 62.5T469 256t-62.5 150.5T256 469t-150.5-62.5T43 256t62.5-150.5T256 43m-13 315l2 3h22q2 0 2-3l64-42q17-12 6-30q-10-18-30-7l-30 22V171q0-22-21-22t-21 22v130l-30-22q-20-11-30 7q-10 19 6 30z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Download(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M491 299h-43l-58 42h79v128H43V341h79l-58-42H21q-8 0-14.5 6.5T0 320v149q0 18 12.5 30.5T43 512h426q18 0 30.5-12.5T512 469V320q0-8-6.5-14.5T491 299M256 0q-21 0-21 21v297l-94-77q-7-6-16-5t-14 7q-6 7-5 16t7 14l143 111l141-111q15-15 2-30q-16-14-30-2l-92 77V21q0-21-21-21m171 405q0 9-6.5 15.5T405 427q-8 0-14.5-6.5T384 405q0-8 6.5-14.5T405 384q9 0 15.5 6.5T427 405m-64 0q0 9-6.5 15.5T341 427q-8 0-14.5-6.5T320 405q0-8 6.5-14.5T341 384q9 0 15.5 6.5T363 405"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DownloadFromCloud(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M427 171v-11q0-65-47.5-112.5T267 0q-57 0-101 35.5T111 126q-48 6-79.5 42.5T0 254q0 53 38 91.5t90 38.5h299q37-8 61-38t24-69q0-38-24-68.5T427 171m0 170H128q-35 0-60-26t-25-61q0-31 25.5-54.5T128 171q4 0 11-11l10-32q8-39 41.5-62T267 43q48 0 82.5 34.5T384 160v32q0 9 6.5 15t14.5 6h13q22 5 36.5 23t14.5 41t-11 41.5t-31 22.5M307 218l-30 21V128q0-21-21-21t-21 21v109l-30-22q-20-11-30 7q-10 19 6 30l64 42q2 0 2 3h22q2 0 2-3l64-42q18-11 6-30q-17-15-34-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dribbble(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M232 2Q137 2 69.5 69.5T2 232t67.5 162.5T232 462t162.5-67.5T462 232T394.5 69.5T232 2m141 119q37 46 38 103q-39-9-73-9q-25 0-52 5q-4-10-12-28q60-27 99-71M232 52q62 0 112 40q-30 37-88 63q-26-50-62-98q20-5 38-5m-79 19q36 45 64 98q-69 20-142 20h-3q-10 0-14-1q19-80 95-117M52 232v-3q5 1 20 1q90 0 163-24q2 4 5.5 12t5.5 12q-57 19-102 60q-32 29-49 58q-43-51-43-116m180 180q-60 0-107-36q23-33 42-51q40-39 93-57q21 58 36 132q-30 12-64 12m102-32q-11-55-33-121q17-3 36-3h1q32 0 70 9q-12 71-74 115"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DripDry(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 43v426q0 18 12.5 30.5T45 512h342q17 0 29.5-12.5T429 469V43q0-18-12.5-30.5T387 0H45Q28 0 15.5 12.5T3 43m384 426H45V43h342zm-256-64q21 0 21-21V128q0-21-21-21q-22 0-22 21v256q0 21 22 21m85 0q21 0 21-21V128q0-21-21-21t-21 21v256q0 21 21 21m85 0q22 0 22-21V128q0-21-22-21q-21 0-21 21v256q0 21 21 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Drop(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m229 59l-9-17q-7-16-23.5-26.5T163 5q-39 0-58 37l-8 17q-2 5-7 15Q8 256 3 308q0 69 47 118t113 49q65 0 112.5-49T323 308q0-11-2.5-25t-9.5-33t-12.5-34t-16.5-40t-16.5-38t-19-41.5T229 59m-66 373q-48 0-83-36.5T45 308q0-42 90-230l8-17q8-13 20-13q11 0 19 13l8 17q90 182 90 230q0 52-34.5 88T163 432m10-85q-21 0-32-13q-10-13-10-49q2-8-4-16t-13-8q-8-1-16.5 5.5T88 283q0 56 21 79q25 27 64 27q22 0 22-21t-22-21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dropbox(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M375 161L238 77l87-72l137 85zm-51 159l132-78l-80-67l-129 78zm-107-67L88 175L8 242l132 78zM94 309v25l132 81V261l-87 72zm276 0l-45 24l-87-72v154l132-81zM226 77L140 5L2 90l87 71z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Drupal(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M309 87q-10-6-33-16.5T241 52q-15-10-49-50q-5 36-24 52q-14 11-56 33q-11 6-24.5 16.5t-34 31.5t-34 55T6 264q0 84 60.5 141T210 462t141.5-54.5T410 267q0-115-101-180m1 331q-15 15-46 18q-56 5-76-14q-6-5 0-9q3-3 6-3q2 0 4 1q15 12 51 12t57-14q6-4 6 1q2 5-2 8m-77-38q14-12 21-14q5-3 19-3t20 4q5 4 10 16q3 6-4 9q-3 2-6-5q-8-12-22-12q-12 0-28 12q-10 8-13 5q-3-6 3-12m161-84q0 32-15 54q-15 21-30 20q-7 0-32.5-25.5T280 318q-10 0-55.5 27T154 372q-28 0-43-10q-20-13-20-43q1-27 22.5-47t51.5-20q29-1 67 25t48 26q11 0 43-22t43-22q28 0 28 37"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dry(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M429 43q0-18-12.5-30.5T387 0H45Q28 0 15.5 12.5T3 43v426q0 18 12.5 30.5T45 512h342q17 0 29.5-12.5T429 469zm-42 426H45V43h342z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DryFlat(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M45 0Q28 0 15.5 12.5T3 43v426q0 18 12.5 30.5T45 512h342q17 0 29.5-12.5T429 469V43q0-18-12.5-30.5T387 0zm342 469H45V43h342zM88 277h256q21 0 21-21t-21-21H88q-21 0-21 21t21 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DryInTheShade(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 43v426q0 18 12.5 30.5T45 512h342q17 0 29.5-12.5T429 469V43q0-18-12.5-30.5T387 0H45Q28 0 15.5 12.5T3 43m98 0L45 98V43zM45 158L161 43h46L45 205zm342 311H45V265L267 43h120z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DryNormalHightHeat(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469 0H43Q25 0 12.5 12.5T0 43v426q0 18 12.5 30.5T43 512h426q18 0 30.5-12.5T512 469V43q0-18-12.5-30.5T469 0m0 469H43V43h426zm-213-23q79 0 135.5-56T448 254t-56-136t-136-56t-136 56t-56 136t56 136t136 56m0-339q61 0 105 43.5T405 256t-43.5 105.5T256 405t-105.5-43.5T107 256t43.5-105.5T256 107m-43 149q0 9-6 15t-15 6t-15-6t-6-15t6-15t15-6t15 6t6 15m64 0q0 9-6 15t-15 6t-15-6t-6-15t6-15t15-6t15 6t6 15m64 0q0 9-6 15t-15 6t-15-6t-6-15t6-15t15-6t15 6t6 15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DryNormalLowHeat(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469 0H43Q25 0 12.5 12.5T0 43v426q0 18 12.5 30.5T43 512h426q18 0 30.5-12.5T512 469V43q0-18-12.5-30.5T469 0m0 469H43V43h426zm-213-23q79 0 135.5-56T448 254t-56-136t-136-56t-136 56t-56 136t56 136t136 56m0-339q61 0 105 43.5T405 256t-43.5 105.5T256 405t-105.5-43.5T107 256t43.5-105.5T256 107m21 149q0 9-6 15t-15 6t-15-6t-6-15t6-15t15-6t15 6t6 15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DryNormalNoHeat(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469 0H43Q25 0 12.5 12.5T0 43v426q0 18 12.5 30.5T43 512h426q18 0 30.5-12.5T512 469V43q0-18-12.5-30.5T469 0m0 469H43V43h426zm-213-23q79 0 135.5-56T448 254t-56-136t-136-56t-136 56t-56 136t56 136t136 56"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dzone(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M426 74q-28 0-35 27L231 48q0-8-6-14t-14-6t-14.5 6t-6.5 14t6 14L27 256q-2-1-6-1q-19 0-19 19q0 8 5.5 13.5T21 293q11 0 17-11l153 37q0 1-.5 3.5t-.5 3.5q0 14 10 24t24 10t24-10t10-24q0-7-1-10l117-48q5 5 13 5q7 0 12.5-5.5T405 255q0-10-7-15l19-94q1 0 4 .5t5 .5q15 0 25.5-11t10.5-26t-10.5-25.5T426 74m-203-9l72 88l-73 30l-5-115zm-12 4l4 117l-33 13q0-1-6-7l28-124q2 1 7 1m-27 140v-4l32-13l4 100q-2 0-5 1t-4 1l-33-70q6-6 6-15m61 90q-9-7-19-7l-4-102l77-32l35 43zm59-143l88-35q0 2 5 11l-60 65zm86-41l-90 36l-73-89q2-3 4-9l159 53zM200 66l-28 124q-6-2-8-2q-8 0-14.5 6t-6.5 15q0 1 1 3v2L33 259q-1-1-2-1zM39 278q0-1 1-2v-2q0-6-3-9l110-45q5 9 17 9q4 0 10-2l33 70q-10 6-15 17zm216 34q-2-4-6-10l88-97l34 41q-2 6-2 9t2 9zm139-75q-1 0-3.5-.5t-3.5-.5q-9 0-13 6l-34-41l60-66q6 6 13 9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ebay(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m415 92l-36 77l-33-77h-51l9 17q-2 0-7.5-.5t-8.5-.5q-21 0-33 9q-14 8-14 31h35q0-20 14-20t14 19v10q-40 0-52 10q-10-8-23-8q-20 0-29 14h-1v-67h-35v61q-7-4-19-6q-22-5-49-5q-58 0-76 15T2 214q0 30 18 45t76 15q27 0 49-5q6-1 19-6v26h34v-14q9 17 30 17q34 0 36-50h8q23 0 33-19h1l2 17h33q-2-16-2-24v-34l15 29v64h46v-64l62-119zM96 176q14 0 21.5 5t8 9t.5 12H67q0-7 .5-11t7.5-9.5t21-5.5m49 57h-21q0 21-28 21q-29 0-29-34h97v13zm84-9q0 26-3 36t-12 10t-12-10t-3-36q0-12 .5-18t1.5-13t4.5-9.5t8.5-2.5t8.5 2.5t4.5 9.5t1.5 13t.5 18m58-4q-14 0-14-20q0-6 1.5-10t2.5-6t5.5-3.5l6-2l8-1l7.5-.5v22.5l-1.5 8l-3 6.5l-5 4.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Egg(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M277 224q0-37-17-83.5t-50-83T139 21T67.5 57.5t-50.5 83T0 224q0 58 40.5 98.5T139 363q57 0 97.5-40.5T277 224m-234 0q0-47 31.5-103.5T139 64t64.5 56.5T235 224q0 40-28 68t-68 28t-68-28t-28-68"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eject(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M365 288H67q-28 0-46 18.5T3 352v43q0 27 18 45.5T67 459h298q28 0 46-18.5t18-45.5v-43q0-27-18-45.5T365 288m22 107q0 9-6 15t-16 6H67q-10 0-16-6t-6-15v-43q0-21 22-21h298q22 0 22 21zM254 15Q241 0 218 0q-25 0-36 15L18 192q-25 22-11 49q10 26 38 26h342q26 0 40-26q11-28-8-49zm133 209H45v-4L210 43q2-1 6-1t4 1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ember(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M380 160q-3 11-11.5 20.5T354 194l-7 4q-27-68-62-117t-56-64L208 2q-2 22-17.5 43.5T160 78l-15 11l-41-21q3 16-1 32.5T94 126l-5 9q-38 25-51 67.5T30 276l6 30l-36-17q8 55 30.5 93T80 434.5t52.5 22T176 462l17-1q48 0 84.5-17t56-43.5t32-58.5t15.5-64t3-58.5t-2-43.5zM186 425q-39 0-65.5-19T85 360.5t-12-53t-1-45.5l2-18q2 7 7 12.5t9 7.5l4 2q16-41 37-70t34-38l13-9q1 13 10 26t18 19l9 7l25-13q-2 9 .5 19.5T246 224l3 5q23 15 30.5 40.5T284 313l-4 18l22-10q-4 33-18 56t-30 32t-31.5 13t-26.5 4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Enlarge(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M177 305L43 439v-55q0-21-22-21q-21 0-21 21v107q0 9 6 15t15 6h107q21 0 21-21q0-22-21-22H73l134-134q13-15 0-30q-15-13-30 0M491 0H384q-21 0-21 21q0 22 21 22h55L305 177q-13 15 0 30q6 6 15 6t15-6L469 73v55q0 21 22 21q21 0 21-21V21q0-9-6-15t-15-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Etsy(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M57 459q17-1 38.5-1t49.5.5t44 .5t44.5-.5t50.5-.5t39 1q5 0 15 2.5t14 2.5t13-3q4-6 5-13.5t0-17.5t0-16q1-9 6-24.5t7.5-24.5t-2.5-18t-20-12q-11 9-12 27.5t-5 26.5q-15 31-101 37q-101 8-115-19q-5-9-6.5-22t0-33t1.5-26q0-7-.5-22t-.5-23.5t1.5-19.5t4.5-19q17 1 48-2t54.5-3t38.5 9q10 6 15.5 25.5T304 295q8-2 10-8t-.5-21.5T311 246v-49q0-4 1-12.5t1.5-13t-.5-10t-5-8t-11-2.5q-9 7-11 25.5T274 202q-8 5-29 7q-60 7-120-5q-7-72 0-144q17-17 78-17q110 0 125 26q3 5 3.5 14.5t2.5 17t8 8.5q11 2 14-7t1.5-29.5T356 48q0-6 3.5-18.5T361 11q-3-4-7.5-6t-7-1.5t-8 1.5t-6.5 1q-36 5-81.5 6t-120 0T45 11q-2 0-8.5-.5T26 10t-9 2t-7 6q-2 12 6 16t20.5 6T54 48q10 10 12 40q2 16 2 40t-1 57t-1 52q0 20 1 54.5t1 58t-2 39.5q-2 25-12 35q-6 6-21 7t-25 5.5T0 456q5 4 12.5 5.5t13.5 1t16-2t15-1.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func EuroBill(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469 21H43q-18 0-30.5 13T0 64v256q0 17 12.5 30T43 363h426q18 0 30.5-13t12.5-30V64q0-17-12.5-30T469 21M149 320H43V64h106zm128 0h-21v-75q0-10 11-10q10 0 10 10zm128 0h-21v-75q0-10 11-10q10 0 10 10zm64 0h-21v-75q0-23-15-38t-38-15t-38.5 15.5T341 245v75h-21v-75q0-23-15-38t-38-15t-38.5 15.5T213 245v75h-21V192h277zm0-171H192V64h277zM85 299q22 0 22-22v-42q0-22-22-22q-9 0-15 6t-6 16v42q0 10 6 16t15 6m278-171h64q9 0 15-6t6-15q0-10-6-16t-15-6h-64q-22 0-22 22q0 9 6 15t16 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Evernote(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M96 101H22l93-94v76zm265-37q-6-10-30.5-20T290 34h-56Q216 2 174 2q-15 0-23.5 2.5T139 14t-3.5 9.5T135 36v65l-19 20H28Q3 138 3 173q0 18 3 41t11 52.5T43.5 318T88 344q39 6 64.5-2t31.5-16.5t6-13.5l1-52q3 5 8 13.5t22 22t34 13.5q28 0 44.5 13.5T316 354v42q-1 27-24 27h-49q-15-13-15-30q0-22 16-22l17 1v-36q-3 0-8 .5t-17.5 3.5t-22 8t-17.5 16.5t-8 27.5q0 38 20.5 54t48.5 16h50q4 0 10-2t21.5-13.5t27.5-30t21.5-56.5t9.5-88q0-37-1.5-67t-4.5-50.5t-5.5-35t-7-24.5t-7-15t-6.5-10t-4-6m-90 159q5-29 22-29q8 0 18 10t16 20l6 9q-53-10-62-10"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Extinguisher(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 159q12 14 27 3q31-24 54-26q9 26 42 38v7q-37 8-61 37t-24 67v188q0 19 14 33t33 14h119q20 0 33.5-14t13.5-33V285q0-38-24-67t-61-37v-7q13-3 23-15l118 58q2 4 8 4q9 0 11-2q11-7 11-19V29q0-12-11-19t-21 0L197 68q-15-17-42-17q-21 0-37.5 12.5T95 95q-49 5-83 30q-8 7-8.5 17t6.5 17m209-53l85-43v103l-85-43zm0 367q0 4-5 4H95q-4 0-4-4V307h128zm-5-209H95q6-19 22.5-31t37.5-12q20 0 36.5 12t22.5 31M155 93q9 0 15 6t6 16q0 9-6 15t-15 6q-10 0-16-6t-6-15q0-22 22-22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eye(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M264 43Q102 43 10 181q-6 11 0 22q92 138 254 138t254-138q6-11 0-22Q426 43 264 43m0 42q27 0 45.5 18.5T328 149q0 28-18.5 46T264 213t-45.5-18t-18.5-46q0-27 18.5-45.5T264 85m0 214q-43 0-80.5-12t-63-31t-40-34T55 192q50-64 113-90q-11 28-11 47q0 45 31 76t76 31t76-31t31-76q0-19-11-47q70 29 113 90q-80 107-209 107"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Facebook(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M51 91v63H4v78h47v230h95V232h65q6-37 8-78h-72v-53q0-6 6.5-12.5T168 82h52V2h-71q-28 0-48.5 8.5T71 29.5T57 55t-5.5 21.5T51 91"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FacebookAlt(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M462 78q0-30-23-53T386 2H78Q48 2 25 25T2 78v308q0 30 23 53t53 23h154V288h-56v-76h56v-30q0-39 25.5-68.5T318 84h62v76h-62q-5 0-9.5 6t-4.5 15v31h76v76h-76v174h82q30 0 53-23t23-53z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FacebookPlaces(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M311 116q-1-3-3.5-9.5T304 98q-21-49-61.5-72.5T158 2Q101 2 55 39.5T1 145v19l1 5v7q2 16 10.5 36.5t15.5 33t22.5 38T71 317q29 49 89 145q8-13 55-94q2-3 7.5-12t7.5-14q2-3 6.5-8.5t6.5-8.5q7-13 30.5-49.5T308 214t11-49v-22q0-11-8-27m-152 99q-39 0-54-38q-1-5-1-15v-13q0-25 17-38.5T161 97q25 0 41.5 17t16.5 42t-17.5 42t-42.5 17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Facto(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M333 272L370 6h92l-70 266zm20 39q-22 0-37 15t-15 36q0 22 15 37t37 15t37-15t15-37q0-21-15-36t-37-15m-243-27h158l16-64H123l27-139h174l16-77H78L2 407h84z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Feedburner(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M364 304q2-12 2-20q1-59-24.5-90.5T270 135q10 55-18.5 86.5T179 241q24-39 29-69t-3.5-46.5t-17.5-36T172.5 52T178 2q-44 31-80 104.5T62 242q2 37 7 62q-68 28-68 67q0 37 63 64t152 27t152-27t63-64q0-39-67-67M134 194q-11 52 12 89.5t56 37.5q63 0 88-100q36 27 39 72q3 49-37 83v1q-12 8-19 12q-1 1-2 1l-8 4l-2 1q-10 4-23 7q-3 0-4 1q-2 0-5 .5t-4 .5q-3 0-4 1h-26q-1 0-4-.5t-4-.5h-1q-8-2-10-2l-1-1l-9-3h-1q-30-12-47.5-39.5T101 293l-1 1q2-48 34-100"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FiftyOneHundredTwenty(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M503 21q-9-2-17 3t-9 14l-12 67q-5-3-15-13q-24-28-58-28t-58 28q-15 15-27 15q-11 0-28-15q-24-28-58-28q-35 0-57 28q-16 15-28 15t-28-15Q84 66 55 64l-4-26q-2-8-10-13t-16-4t-13.5 9.5T8 47l58 286q4 22 21.5 36.5T127 384h274q22 0 39.5-14.5T462 333l58-286q2-8-3.5-16.5T503 21m-81 303q-3 17-21 17H127q-18 0-21-17L63 109q3 2 7.5 6.5T78 122q25 27 58 27t58-27q14-15 27-15t28 15q25 27 58 27q35 0 57-27q15-15 28-15t28 15q19 19 36 25zm-137-89q0 8-6 14.5t-15 6.5t-15-6.5t-6-14.5q0-9 6-15.5t15-6.5t15 6.5t6 15.5m107 0q0 8-6.5 14.5T371 256q-9 0-15.5-6.5T349 235q0-9 6.5-15.5T371 213q8 0 14.5 6.5T392 235m-213 0q0 8-6.5 14.5T157 256q-8 0-14.5-6.5T136 235q0-9 6.5-15.5T157 213q9 0 15.5 6.5T179 235"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func File(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M32 512h343q13 0 22.5-9.5T407 480V32q-2-13-11.5-22.5T373 0H183L0 192v288q0 13 9.5 22.5T32 512M192 43v128q0 13-4 17t-17 4H43zM43 235h128q64 0 64-64V43h128v426H43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Film(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 512h384q27 0 45.5-18.5T512 448V64q0-27-18.5-45.5T448 0H64Q37 0 18.5 18.5T0 64v384q0 27 18.5 45.5T64 512M43 64q0-21 21-21h384q21 0 21 21v384q0 21-21 21H64q-21 0-21-21zm405 21q0 9-6.5 15.5T427 107q-9 0-15.5-6.5T405 85q0-8 6.5-14.5T427 64q8 0 14.5 6.5T448 85M192 448h128q27 0 45.5-18.5T384 384v-43q0-27-18.5-45.5T320 277H192q-27 0-45.5 18.5T128 341v43q0 27 18.5 45.5T192 448m-21-107q0-21 21-21h128q21 0 21 21v43q0 21-21 21H192q-21 0-21-21zm21-106h128q27 0 45.5-18.5T384 171v-43q0-27-18.5-45.5T320 64H192q-27 0-45.5 18.5T128 128v43q0 27 18.5 45.5T192 235m-21-107q0-21 21-21h128q21 0 21 21v43q0 21-21 21H192q-21 0-21-21zm277 43q0 8-6.5 14.5T427 192q-9 0-15.5-6.5T405 171q0-9 6.5-15.5T427 149q8 0 14.5 6.5T448 171m0 85q0 9-6.5 15t-14.5 6q-9 0-15.5-6t-6.5-15t6.5-15t15.5-6q8 0 14.5 6t6.5 15m0 85q0 9-6.5 15.5T427 363q-9 0-15.5-6.5T405 341q0-8 6.5-14.5T427 320q8 0 14.5 6.5T448 341m0 86q0 8-6.5 14.5T427 448q-9 0-15.5-6.5T405 427q0-9 6.5-15.5T427 405q8 0 14.5 6.5T448 427M107 85q0 9-6.5 15.5T85 107q-8 0-14.5-6.5T64 85q0-8 6.5-14.5T85 64q9 0 15.5 6.5T107 85m0 86q0 8-6.5 14.5T85 192q-8 0-14.5-6.5T64 171q0-9 6.5-15.5T85 149q9 0 15.5 6.5T107 171m0 85q0 9-6.5 15T85 277q-8 0-14.5-6T64 256t6.5-15t14.5-6q9 0 15.5 6t6.5 15m0 85q0 9-6.5 15.5T85 363q-8 0-14.5-6.5T64 341q0-8 6.5-14.5T85 320q9 0 15.5 6.5T107 341m0 86q0 8-6.5 14.5T85 448q-8 0-14.5-6.5T64 427q0-9 6.5-15.5T85 405q9 0 15.5 6.5T107 427"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Filter(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m169 399l43 32q10 9 25 9q6 0 20-4q23-12 23-39V246L419 73q17-21 4-45q-15-25-38-25H47Q18 3 9 26q-10 26 4 45l139 175v119q0 22 17 34M47 45h338L237 229v168l-42-32V229z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fire(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M93 219q0 46 28 74q32 32 81 32h7q57 0 89-32q30-30 30-76q0-33-16.5-64T273 100.5t-47.5-39t-40-25T166 27L117 5l19 49q7 25 4.5 35.5T125 118q-2 4-10 19.5t-11 24t-7 25t-4 32.5m69-75q21-31 23-55q100 60 100 130q0 31-17 44q-8 9-23 13q7-19 5.5-39.5t-4.5-32t-5-15.5q-9-20-28-13q-8 3-12 11.5t-1 16.5q17 53-9 77q-18 0-36.5-15T136 219q0-7 .5-13.5t1.5-11t3-10.5t3-8.5t4.5-8.5t4-7.5t4.5-8t5-7.5m245 160l-196 53l-197-53l-12 43l125 32L2 411l12 42l197-53l196 53l13-42l-126-32l126-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fish(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M493 99q-20-14-41-2q-40 17-66 59q-28-54-87-77V45q0-17-15-34q-16-11-37-8q-56 12-74 66q-72 15-122.5 60T0 235q0 72 70 121.5T224 406q56 0 98.5-24t63.5-66q26 42 66 60q14 4 17 4q13 0 24-6q19-13 19-36V133q0-25-19-34M224 65q5-12 32-20v20zM43 235q0-30 21.5-57t57.5-45q27 33 27 102t-27 103q-36-18-57.5-45.5T43 235m181 128q-24 0-62-8q30-51 30-120q0-68-30-119q24-9 62-9q62 0 100.5 35t38.5 93t-38.5 93T224 363m245-19q-64-42-64-105q0-33 19.5-63.5T469 135zM128 214q0 9-6.5 15t-14.5 6q-9 0-15.5-6T85 214t6.5-15t15.5-6q8 0 14.5 6t6.5 15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flag(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M283 128V64q-4-3-11.5-8.5T239 39t-50.5-15.5t-64 .5T48 49V21q0-9-6-15T27 0Q17 0 11 6T5 21v491h43V320q64-44 149-13v98q10 6 24 10.5t23.5 7T273 426t25 1h49q71 0 121-47l7-7V81l-37 32q-3 2-8 5t-21 9.5t-33 10t-42.5 2T283 128M48 277V98q110-64 192-13v192q-98-50-192 0m384 77q-118 69-192 30v-64h43V171q86 26 149-2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlagCorner(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M48 384h107v64h330l-72-171l72-170H325V43H48V21q0-9-6-15T27 0Q17 0 11 6T5 21v491h43zm277-235h96l-55 128l55 128H197v-21h128zM48 235V85h235v256H48z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlagScout(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M27 0Q17 0 11 6T5 21v491h43V384q73-2 136-18t105.5-40.5t76.5-54t54-58.5t33-52.5t17-38.5l5-15q-65 17-129 19t-110-5.5T147.5 99T84 74T48 53V21q0-9-6-15T27 0m21 235v-34l21 6v49q-17-6-21-9zm363-86q0 1-13 20t-16.5 24t-18.5 24.5t-24 27t-28.5 24t-36 24t-42.5 19t-52 16.5t-60.5 9t-71.5 4v-47h4q24 5 42-11.5t18-41.5v-26q0-17-10-31.5T76 164l-28-8v-54q88 39 140 49q99 19 223-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flickr(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M462 192q0 43-30 72.5T359 294t-73-29.5t-30-72.5t30-72.5T359 90t73 29.5t30 72.5M105 90q-43 0-73 29.5T2 192t30 72.5t73 29.5t73-29.5t30-72.5t-30-72.5T105 90"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folder(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 64H237q-10 0-17-9l-19-27Q183 0 147 0H64Q37 0 18.5 18.5T0 64v256q0 27 18.5 45.5T64 384h384q27 0 45.5-18.5T512 320V128q0-27-18.5-45.5T448 64m21 256q0 21-21 21H64q-21 0-21-21V64q0-21 21-21h83q9 0 17 8l17 26q21 30 56 30h211q21 0 21 21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folkd(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M303 98q-2 2-16 2q-1 2-2 8t-3 8q5 4 20 7.5t20 8.5q10-4 10.5-16.5T338 98q9-1 15 2.5t7.5 6.5l5.5 11q-2 2-23 2q-3 2-5.5 7t-4.5 7q2 4 9.5 10.5T352 157q46-12 74-16q2-2 6.5-12.5T443 116q23 4 25 39q-7 12-21.5 9.5T426 150q-28 5-69 14q-1 2-1 13.5t-2 16.5q9 6 19.5-.5T389 187q4 4 5 13.5t-2.5 16.5t-11.5 6q-10 0-12-20q-4 1-7.5-.5T354 201q-4 2-12.5 10.5T326 223q1 12 16 12.5t19 3.5q-1 23-30 23v-9.5l-.5-8.5l-3-7.5l-8.5-6.5q-7 1-18 7.5t-19 8.5q0 4 5 16q1 1 11.5.5T310 267q0 13-13.5 20.5T273 287q0-4 1.5-7t4-6t3.5-5q-5-6-9-18q-41 14-47 14q0 2-1.5 5.5t-.5 7.5q-1 5 5.5 7t6.5 5q3 15-14 18t-21-9q0-6 7-13t4-17q-5-2-20.5-3t-21.5-4q-3 2-6 7.5t-5 6.5q0 3 4 7t3 9q-11 5-24 0t-11-18q4-4 15.5-6t14.5-10q-20-24-23-37q-20 5-52.5 17T42 253q-8 23-9 23q-7 1-21 0q-7-10-3-27.5T26 230q5-1 8 4t6 5q17-4 48.5-13.5T136 212q-3-13 3.5-27.5T151 162t-.5-12t-28.5-2q-5-14 1.5-25t16.5-9q11 1 12.5 16.5T161 150q3-1 18-12t26-13q1-2-2-16q-1-3-8-3t-8-2V88q5-5 16.5-4t15.5 7q2 11-11 13q0 4 1.5 6.5l3 5l2.5 4.5q9-4 58-4q4-4 4-14q1-4-4-5q-7-1-5-6q1-12 10-13.5t17 5t8 15.5m19 103q5-3 6.5-11t.5-16t0-17q-17-14-27 7q-11 20 4 32q6 5 16 5m-103 25q17 2 34-12t15-29q-2-16-25-26t-40-2q-14 7-14.5 23t9.5 30.5t21 15.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Forbidden(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0Q150 0 75 75T0 256t75 181t181 75t181-75t75-181t-75-181T256 0M43 256q0-88 62.5-150.5T256 43q79 0 134 49L92 390q-49-55-49-134m213 213q-79 0-134-49l298-298q49 55 49 134q0 88-62.5 150.5T256 469"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Formspring(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M160 37Q206 2 272 2q49 0 87 19q10 5 10 17v5q-3 9-11 12q-2 1-7 1q-6 0-8-2q-25-12-54-15q-20-4-47-4q-56 0-104 23q14-14 22-21m232 297q0-9-8-15q-4-3-10-3q-11 0-15 7q-15 21-31 31q-5 3-14 9t-14 10l-6 4q-8 5-8 15v24l-57-24q-2-1-7-1h-14q-47 0-87.5-18.5T58 323h-1h1q-1 0-3-2q-18-27-18-56q0-52 50-89t120-37q51 0 95 22h1q19 9 33 27q12 15 12 34q0 30-30 52q-33 24-82 24q-16 0-23-1h-1q-34-6-60-21q-55-33-55-84v-7q-20 15-33 35q13 55 70 88l1 1q33 18 72 24h1q18 2 28 2q61 0 103-31q46-34 46-82q0-61-66-94l-1-1q-52-25-111-25q-84 0-144 46Q0 196 0 265q0 42 26 79l3 3q28 38 76 59.5T208 428h11l78 33q2 1 7 1q6 0 10-3q9-5 9-15v-42l1-1q4-3 25-16q20-13 40-40q3-6 3-11m-150-80v3zq18 2 25 2q33 0 62-12q7-12 7-22q0-8-5-20q-28 17-64 17q-7 0-19-2q-68-9-93-58q-17 4-34 12q14 32 46 54t75 29m158-99q0-54-58-83l-2-1q-46-22-98-22q-75 0-127 40q-17 14-29 29q53-27 119-29q24-4 37-4q43 0 83 19l1 1q25 14 33 34q23 20 34 48q7-17 7-32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Forrst(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 462h162v-95l-49-34l12-17l37 25v-74h50v47l36-19l11 19l-47 25v26l71-35l9 19l-80 40v73h172L192 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FortyOneHundredFive(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M51 38q-2-8-10-13t-16-4t-13.5 9.5T8 47l58 286q4 22 21.5 36.5T127 384h274q22 0 39.5-14.5T462 333l58-286q2-9-3-17t-14-9q-9-2-17 3t-9 14l-12 67q-5-3-15-13q-24-28-58-28t-58 28q-15 15-27 15q-11 0-28-15q-24-28-58-28q-35 0-57 28q-16 15-28 15t-28-15Q84 66 55 64zm27 84q25 27 58 27t58-27q14-15 27-15t28 15q25 27 58 27q35 0 57-27q15-15 28-15t28 15q19 19 36 25l-36 177q-3 17-22 17H127q-18 0-21-17L63 109q3 2 7.5 6.5T78 122m250 134q0 9-6.5 15t-14.5 6q-9 0-15.5-6t-6.5-15t6.5-15t15.5-6q8 0 14.5 6t6.5 15m-85 0q0 9-6.5 15t-15.5 6q-8 0-14.5-6t-6.5-15t6.5-15t14.5-6q9 0 15.5 6t6.5 15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Forward(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13 427q4 2 8 2q12 0 15-6l135-149v134q0 15 12 19q1 0 4 1t5 1q12 0 15-6l171-192q10-16 0-28L207 11q-7-14-24-6q-12 4-12 19v137L36 11Q29-3 13 5Q0 9 0 24v384q0 15 13 19M213 79l122 137l-122 137zM43 79l121 137L43 353z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Foursquare(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M462 3L290 179l-62-60l-26 25q-11-30-37.5-49T105 76q-43 0-73 29.5T2 176t29.5 70t71.5 30l125 121l143-139l-14-14l105-82zm-15 150L216 335l-87-77q-16 4-25 4q-37 0-62.5-25.5T16 176t25.5-60T104 91q36 0 62 25t26 60q0 18-8 36l35 57L447 35z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FriedEgg(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m439 13l-64 64q-17 17-10 40l-15 15q-61-47-137-47q-88 0-150.5 63T0 299t62.5 150.5T213 512t151-62.5T427 299q0-78-49-135l15-15h12q18 0 30-12l64-64q13-13 13-30t-13-30t-30-13t-30 13M213 469q-71 0-120.5-49.5T43 299t50-121t120-50q71 0 121 50t50 121q0 70-50 120t-121 50m0-234q-27 0-45.5 18T149 299q0 27 18.5 45.5T213 363q28 0 46-18.5t18-45.5q0-28-18-46t-46-18m0 85q-9 0-15-6t-6-15q0-10 6-16t15-6q22 0 22 22q0 9-6 15t-16 6m0-171q-37-2-77 22t-51 81q-9 39 0 72q9 39 37.5 69t67.5 34h45q69 0 102-54q15-22 15.5-54.5T335 262q-11-13-17-34q-8-31-35.5-55T213 149m90 203q-23 32-68 32h-41q-24-1-42.5-22T128 316q-5-26 0-56q14-68 81-68h4q28 0 44 15.5t20 33.5q10 30 26 51q8 8 8 27.5t-8 32.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Friendfeed(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M462 81h-68q-10 0-19 12.5t-9 29.5v40h81v81h-81v203h-81V244H163v203H83V244H2v-81h81v-40q0-50 31.5-86T192 1h68v80h-68q-11 0-20 12.5t-9 29.5v40h122v-40q0-50 32-86t77-36h68z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Friendster(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M190 41q24 20 32.5 52.5T222 156q-10 27-43.5 27T135 155q-10-27-5-59.5T155 41q20-18 35 0m83-35q-35 17-33 90q8 73 47 74q49-20 35-108q-3-23-15.5-41T273 6m162 53q-42-9-52 15q-11 21-26 81t-36 87q-24 29-58.5 44.5T193 296q-42-6-64-41q-12-20-27.5-72T68 108q-24-12-37.5-6t-20 15t-8 26.5T2 171t2 18v2q9 56 17 70q33 69 108.5 100T279 371q78-22 129.5-93.5T462 122q0-50-27-63"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fuck(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 276q0 58 36.5 102t93.5 51q68 8 118-36t50-111v-64q0-13-11-28t-31-15v107q0 47-35 79t-83 28q-41-7-68.5-39.5T45 276V152q-17 0-29.5 12.5T3 195zm192-145v85h42v-85q0-22-21-22t-21 22M131 24v192h42V24q0-21-21-21t-21 21M67 131v85h42v-85q0-22-21-22t-21 22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FullScreen(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 512h107q21 0 21-21q0-22-21-22H73l134-134q13-15 0-30q-15-13-30 0L43 439v-55q0-21-22-21q-21 0-21 21v107q0 9 6 15t15 6m470-149q-22 0-22 21v55L335 305q-15-13-30 0q-13 15 0 30l134 134h-55q-21 0-21 22q0 21 21 21h107q9 0 15-6t6-15V384q0-21-21-21m0-363H384q-21 0-21 21q0 22 21 22h55L305 177q-13 15 0 30q6 6 15 6t15-6L469 73v55q0 21 22 21q21 0 21-21V21q0-9-6-15t-15-6M21 149q22 0 22-21V73l134 134q6 6 15 6t15-6q13-15 0-30L73 43h55q21 0 21-22q0-21-21-21H21Q12 0 6 6T0 21v107q0 21 21 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gamepad(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 51v-8q0-10-6-16t-15-6h-43q-21 0-21 22v8q-30 16-43 34H192q-13-18-43-34v-8q0-22-21-22H85q-9 0-15 6t-6 16v8Q35 64 17.5 91T0 149q0 35 19 60L0 314v8q0 26 18 44t44 18h2q31 0 51-26l47-70q16 11 30 11q26 0 36-22h56q10 22 36 22q15 0 30-13l47 72q20 26 51 26h4q26 0 44-18t18-44v-8l-17-105q15-28 15-60q0-31-17.5-58T448 51m-43 34q28 0 46 18.5t18 45.5q0 25-19 47h-2q-17 17-43 17q-27 0-45.5-18T341 149q0-27 18.5-45.5T405 85m-298 0q27 0 45.5 18.5T171 149q0 28-18.5 46T107 213q-26 0-43-17h-2q-19-22-19-47q0-27 18-45.5T107 85M79 333q-3 8-15 8h-2q-19 0-19-19v-2l12-79q20 15 52 15q11 0 27-4zm205-98h-56q-10-22-36-22q21-27 21-64q0-15-2-21h90q-2 6-2 21q0 37 21 64q-26 0-36 22m166 106h-2q-9 0-15-6l-53-81q16 4 27 4q30 0 52-13l10 75v2q0 19-19 19m-66-170q0 9 6 15t15 6q10 0 16-6t6-15q9 0 15-6t6-16q0-9-6-15t-15-6q0-21-22-21q-21 0-21 21q-21 0-21 21q0 22 21 22m-256-64q0 8-6.5 14.5T107 128q-9 0-15.5-6.5T85 107q0-9 6.5-15.5T107 85q8 0 14.5 6.5T128 107m43 42q0 9-6.5 15.5T149 171q-8 0-14.5-6.5T128 149q0-8 6.5-14.5T149 128q9 0 15.5 6.5T171 149m-86 0q0 9-6 15.5T64 171t-15-6.5t-6-15.5q0-8 6-14.5t15-6.5t15 6.5t6 14.5m43 43q0 9-6.5 15t-14.5 6q-9 0-15.5-6T85 192t6.5-15t15.5-6q8 0 14.5 6t6.5 15m107-43h42v43h-42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gdgt(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1 25v276l23 23h253v69H24v46l23 23h276l23-23V324l69-138V48l-69 158V25L323 2H24zm69 230V71h207v184z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gift(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469 128h-76q12-20 12-43q0-35-25-60T320 0q-40 0-64 30q-24-30-64-30q-35 0-60 25t-25 60q0 23 12 43H43q-18 0-30.5 12.5T0 171v42q0 18 12.5 30.5T43 256v192q0 27 18 45.5t46 18.5h298q28 0 46-18.5t18-45.5V256q18 0 30.5-12.5T512 213v-42q0-18-12.5-30.5T469 128M192 469h-85q-22 0-22-21V256h107zm0-256H43v-42h149zm0-85q-17 0-30-12.5T149 85q0-17 13-29.5T192 43t30 12.5T235 85q0 18-13 30.5T192 128m85 341h-42V256h42zm0-256h-42v-42h42zm0-128q0-17 13-29.5T320 43t30 12.5T363 85q0 18-13 30.5T320 128t-30-12.5T277 85m150 363q0 21-22 21h-85V256h107zm42-235H320v-42h149z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Girl(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 192q0-80-56-136T256 0T120 56T64 192v11Q3 359 0 425q-2 44 137 44q53 43 119 43q70 0 119-43q137 0 137-44q-3-66-64-222zM256 469q-62 0-105.5-43.5T107 320V192q0-58 42-105q24 84 254 84q2 6 2 21v128q0 62-43.5 105.5T256 469m0-128q21 0 21-21v-21h-42v21q0 21 21 21m-64-106q0 8-6.5 14.5T171 256q-9 0-15.5-6.5T149 235q0-9 6.5-15.5T171 213q8 0 14.5 6.5T192 235m171 0q0 8-6.5 14.5T341 256q-8 0-14.5-6.5T320 235q0-9 6.5-15.5T341 213q9 0 15.5 6.5T363 235M235 405h42q18 0 30.5-12.5T320 363H192q0 17 12.5 29.5T235 405"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GirlAngel(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 192q0-56-32-107h32q21 0 21-21t-21-21h-73Q322 0 256 0T137 43H64q-21 0-21 21t21 21h32q-32 51-32 107v11Q3 359 0 425q-2 44 137 44q53 43 119 43q70 0 119-43q137 0 137-44q-3-66-64-222zM256 469q-62 0-105.5-43.5T107 320V192q0-58 42-105q24 84 254 84q2 6 2 21v128q0 62-43.5 105.5T256 469m-21-170q0 21 21 21t21-21v-22h-42zm21-86h-43q0 22-21 22q-5 0-8.5-1.5T178 231l-1-1q-6-6-6-17h-43q0 28 17 47q22 17 47 17q32 0 45-17q19-19 19-47m64 22q-5 0-8.5-1.5T306 231l-1-1q-6-6-6-17h-43q0 28 17 47q17 17 47 17q32 0 45-17q19-19 19-47h-43q0 22-21 22m-43 149q0 9-6 15t-15 6t-15-6t-6-15t6-15t15-6t15 6t6 15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GirlAngry(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 192q0-80-56-136T256 0T120 56T64 192v11Q3 359 0 425q-2 44 137 44q53 43 119 43q70 0 119-43q137 0 137-44q-3-66-64-222zM256 469q-62 0-105.5-43.5T107 320V192q0-58 42-105q24 84 254 84q2 6 2 21v128q0 62-43.5 105.5T256 469m0-128q21 0 21-21v-21h-42v21q0 21 21 21m21 22h-42q-18 0-30.5 12.5T192 405q0-8 25-14.5t39-6.5t39 6.5t25 14.5q0-17-12.5-29.5T277 363m-64-107v-43h-85l43 22v21q0 21 21 21t21-21m86-43v43q0 21 21 21t21-21v-21l43-22z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GirlBigSmile(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 192q0-80-56-136T256 0T120 56T64 192v11Q3 359 0 425q-2 44 137 44q53 43 119 43q70 0 119-43q137 0 137-44q-3-66-64-222zM256 469q-62 0-105.5-43.5T107 320V192q0-58 42-105q24 84 254 84q2 6 2 21v128q0 62-43.5 105.5T256 469m0-149q21 0 21-21v-22h-42v22q0 21 21 21m87 0q-20-3-23 19q0 4-3 13t-19 20.5t-42 11.5q-24 0-39.5-11.5t-20-21T192 337q-4-20-26-17q-8 2-13 9.5t-4 16.5q5 28 32 54.5t75 26.5q51 0 76.5-26.5T363 343q2-8-4-15t-16-8m-2-64q9 0 15.5-6t6.5-15v-22h-43v22q0 9 6 15t15 6m-170 0q8 0 14.5-6t6.5-15v-22h-43v22q0 9 6 15t16 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GirlConfused(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 192q0-80-56-136T256 0T120 56T64 192v11Q3 359 0 425q-2 44 137 44q53 43 119 43q70 0 119-43q137 0 137-44q-3-66-64-222zM256 469q-62 0-105.5-43.5T107 320V192q0-58 42-105q24 84 254 84q2 6 2 21v128q0 62-43.5 105.5T256 469m0-128q21 0 21-21v-21h-42v21q0 21 21 21m77 24l-22 10q-11 10-23 0q-31-16-62 0q-12 10-23 0l-22-10q-8-3-16-.5t-11 8.5q-3 8-.5 16.5T162 401l21 11q17 6 30 6q18 0 30-6q12-9 24 0q30 16 62 0l21-11q8-3 10.5-11.5T358 373q-3-6-11-8.5t-14 .5M192 235q0 8-6.5 14.5T171 256q-9 0-15.5-6.5T149 235q0-9 6.5-15.5T171 213q8 0 14.5 6.5T192 235m171 0q0 8-6.5 14.5T341 256q-8 0-14.5-6.5T320 235q0-9 6.5-15.5T341 213q9 0 15.5 6.5T363 235"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GirlCry(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 192q0-80-56-136T256 0T120 56T64 192v11Q3 359 0 425q-2 44 137 44q53 43 119 43q70 0 119-43q137 0 137-44q-3-66-64-222zM256 469q-62 0-105.5-43.5T107 320V192q0-58 42-105q24 84 254 84q2 6 2 21v128q0 62-43.5 105.5T256 469m0-149q21 0 21-21v-22h-42v22q0 21 21 21m85-107q-8 0-14.5 6.5T320 235q0 8 6 14.5t15 6.5t15.5-6t6.5-15t-6.5-15.5T341 213m-21 107q0 8 6 14.5t15 6.5q22 0 22-21v-21h-43zm-149-64q8 0 14.5-6t6.5-15t-6-15.5t-15-6.5t-15.5 6.5T149 235t6.5 15t15.5 6m21 64v-21h-43v21q0 8 6.5 14.5T171 341q21 0 21-21m64 43q-50 0-79 30q-6 6-6 14.5t6 14.5q15 15 30 0q19-17 49-17q27 0 49 17q6 7 15 7t15-7q6-6 6-14.5t-6-14.5q-29-30-79-30"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GirlFlushed(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 192q0-80-56-136T256 0T120 56T64 192v11Q3 359 0 425q-2 44 137 44q53 43 119 43q70 0 119-43q137 0 137-44q-3-66-64-222zM256 469q-62 0-105.5-43.5T107 320V192q0-58 42-105q24 84 254 84q2 6 2 21v128q0 62-43.5 105.5T256 469m-21-149q0 21 21 21t21-21v-21h-42zm42 43h-42q-18 0-30.5 12.5T192 405q0-8 25-14.5t39-6.5t39 6.5t25 14.5q0-17-12.5-29.5T277 363m-42-118q0-22-15.5-37.5T181 192t-38 15t-15 38t15.5 38.5T181 299q23 0 38.5-15.5T235 245m-54 32q-13 0-22.5-9.5T149 245t9.5-22.5T181 213t22.5 9.5T213 245t-9.5 22.5T181 277m150-85q-23 0-38.5 15.5T277 245q0 23 15.5 38.5T331 299q22 0 37.5-15.5T384 245t-15-38t-38-15m0 85q-13 0-22.5-9.5T299 245t9.5-22.5T331 213t22.5 9.5T363 245t-9.5 22.5T331 277"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GirlOmouth(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 203v-11q0-80-56-136T256 0T120 56T64 192v11Q3 359 0 425q-2 44 137 44q53 43 119 43q70 0 119-43q137 0 137-44q-3-66-64-222m-43 117q0 62-43.5 105.5T256 469t-105.5-43.5T107 320V192q0-58 42-105q24 84 254 84q2 6 2 21zm-149 0q21 0 21-21v-22h-42v22q0 21 21 21m-64-85q0 8-6.5 14.5T171 256q-9 0-15.5-6.5T149 235q0-9 6.5-15.5T171 213q8 0 14.5 6.5T192 235m171 0q0 8-6.5 14.5T341 256q-8 0-14.5-6.5T320 235q0-9 6.5-15.5T341 213q9 0 15.5 6.5T363 235m-64 170q0 18-12.5 30.5T256 448t-30.5-12.5T213 405q0-17 12.5-29.5T256 363t30.5 12.5T299 405"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GirlOpenMouth(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 192q0-80-56-136T256 0T120 56T64 192v11Q3 359 0 425q-2 44 137 44q53 43 119 43q70 0 119-43q137 0 137-44q-3-66-64-222zM256 469q-62 0-105.5-43.5T107 320V192q0-58 42-105q24 84 254 84q2 6 2 21v128q0 62-43.5 105.5T256 469m0-149q21 0 21-21v-22h-42v22q0 21 21 21m-64-85q0 8-6.5 14.5T171 256q-9 0-15.5-6.5T149 235q0-9 6.5-15.5T171 213q8 0 14.5 6.5T192 235m171 0q0 8-6.5 14.5T341 256q-8 0-14.5-6.5T320 235q0-9 6.5-15.5T341 213q9 0 15.5 6.5T363 235M213 384q0 17 13 30t30 13t30-13t13-30v-43h-86z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GirlSad(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 192q0-80-56-136T256 0T120 56T64 192v11Q3 359 0 425q-2 44 137 44q53 43 119 43q70 0 119-43q137 0 137-44q-3-66-64-222zM256 469q-62 0-105.5-43.5T107 320V192q0-58 42-105q24 84 254 84q2 6 2 21v128q0 62-43.5 105.5T256 469m0-128q21 0 21-21v-21h-42v21q0 21 21 21m-64-106q0 8-6.5 14.5T171 256q-9 0-15.5-6.5T149 235q0-9 6.5-15.5T171 213q8 0 14.5 6.5T192 235m171 0q0 8-6.5 14.5T341 256q-8 0-14.5-6.5T320 235q0-9 6.5-15.5T341 213q9 0 15.5 6.5T363 235m-86 128h-42q-18 0-30.5 12.5T192 405q0-8 25-14.5t39-6.5t39 6.5t25 14.5q0-17-12.5-29.5T277 363"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GirlSadHunappy(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 192q0-80-56-136T256 0T120 56T64 192v11Q3 359 0 425q-2 44 137 44q53 43 119 43q70 0 119-43q137 0 137-44q-3-66-64-222zM256 469q-62 0-105.5-43.5T107 320V192q0-58 42-105q24 84 256 84v149q0 62-43.5 105.5T256 469m181-83q0-1 1-2t1-2zm-181-45q21 0 21-21v-21h-42v21q0 21 21 21m0 22q-17 0-30 12.5T213 405h86q0-17-13-29.5T256 363m-64-128q0 8-6.5 14.5T171 256q-9 0-15.5-6.5T149 235q0-9 6.5-15.5T171 213q8 0 14.5 6.5T192 235m171 0q0 8-6.5 14.5T341 256q-8 0-14.5-6.5T320 235q0-9 6.5-15.5T341 213q9 0 15.5 6.5T363 235"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GirlSleep(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 192q0-80-56-136T256 0T120 56T64 192v11Q3 359 0 425q-2 44 137 44q53 43 119 43q70 0 119-43q137 0 137-44q-3-66-64-222zM256 469q-62 0-105.5-43.5T107 320V192q0-58 42-105q24 84 254 84q2 6 2 21v128q0 62-43.5 105.5T256 469m-21-170q0 21 21 21t21-21v-22h-42zm21-86h-43q0 22-21 22q-5 0-8.5-1.5T178 231l-1-1q-6-6-6-17h-43q0 28 17 47q22 17 47 17q32 0 45-17q19-19 19-47m64 22q-5 0-8.5-1.5T306 231l-1-1q-6-6-6-17h-43q0 28 17 47q17 17 47 17q32 0 45-17q19-19 19-47h-43q0 22-21 22m-43 149q0 9-6 15t-15 6t-15-6t-6-15t6-15t15-6t15 6t6 15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GirlSmile(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 320q21 0 21-21v-22h-42v22q0 21 21 21m192-128q0-80-56-136T256 0T120 56T64 192v11Q3 359 0 425q-2 44 137 44q53 43 119 43q70 0 119-43q137 0 137-44q-3-66-64-222zM256 469q-62 0-105.5-43.5T107 320V192q0-58 42-105q24 84 254 84q2 6 2 21v128q0 62-43.5 105.5T256 469m-64-234q0 8-6.5 14.5T171 256q-9 0-15.5-6.5T149 235q0-9 6.5-15.5T171 213q8 0 14.5 6.5T192 235m171 0q0 8-6.5 14.5T341 256q-8 0-14.5-6.5T320 235q0-9 6.5-15.5T341 213q9 0 15.5 6.5T363 235M256 405q38 0 59-16t23-32l3-16H171q0 3 .5 7t5 15t12.5 19.5t25.5 15.5t41.5 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GirlTwo(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 192q0-80-56-136T256 0T120 56T64 192v11Q3 359 0 425q-2 44 137 44q53 43 119 43q70 0 119-43q137 0 137-44q-3-66-64-222zM256 469q-62 0-105.5-43.5T107 320V192q0-58 42-105q24 84 254 84q2 6 2 21v128q0 62-43.5 105.5T256 469m0-128q21 0 21-21v-21h-42v21q0 21 21 21m0 64q17 0 30-12.5t13-29.5h-86q0 17 13 29.5t30 12.5m-64-170q0 8-6.5 14.5T171 256q-9 0-15.5-6.5T149 235q0-9 6.5-15.5T171 213q8 0 14.5 6.5T192 235m171 0q0 8-6.5 14.5T341 256q-8 0-14.5-6.5T320 235q0-9 6.5-15.5T341 213q9 0 15.5 6.5T363 235"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GirlUser(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 512h320v-64q0-72-52-124q18-4 26-10q26-18 26-58q0-68-43-134v-5q0-48-34.5-82.5T163 0T80 34.5T45 117v5Q3 193 3 256q0 36 25 58q10 6 26 10Q3 375 3 448zm85-395q0-9 4-21q58 46 145 53v11q0 31-22 53t-52 22q-31 0-53-22t-22-53zm64 192v-32h21v32q0 11-10 11q-11 0-11-11m119-30q-11 6-25 5t-24-5l-4 11q-2 0-4-2v-23q40-21 55-62q9 27 9 53q2 17-7 23m-36-172q-77-6-121-45q22-19 49-19t48 18.5t24 45.5M45 256q0-15 9-53q19 43 55 62v23q-2 0-4 2l-4-11q-30 13-49 0v-1q-1-1-2.5-3t-2.5-4t-1.5-6t-.5-9m0 192q0-36 19.5-67.5T116 333q14 30 47 30q31 0 47-30q30 16 50 48t20 67v21H45z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Github(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M253 3q-18 12-37.5 13t-46-5.5T139 3Q91-3 56.5 18.5T11 73.5t1 71.5t47 61q-19 10-17.5 46.5T63 296q-30 3-46.5 29T2 381.5T21 430q32 33 106 32.5T229 430q20-22 22-59t-15-58q-7-9-66.5-27.5T111 244q1-13 11-18.5t27.5-8.5t24.5-7q33-16 49-55.5t3-75.5q5-1 14-3.5t13-3.5zM87 148q-24-53 8-74q25-15 50 1q33 23 8 73q-13 9-33 9.5T87 148m94 217q2 23-22.5 33.5t-52 5.5T73 382q-6-21 1.5-30t30.5-15q24-8 49-1.5t27 29.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GithubAlt(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M338 13q12 29 4 53q18 15 24 47t-3 62q11 1 32.5 1.5t38 2T462 184q-67-8-88-7q-14 5-15 8q73 6 102 16q-77-13-102-14q-22 43-87 48q1 3 7.5 11.5T287 264t-.5 40.5T290 338q2 4 7 8t7 7q-8 8-22 1.5T266 338q-2-8 0-38t-6-32q0 4-2 40t3 39q1 3 5 8.5t3 10.5q-22 2-28.5-13t-6-46.5T233 264q-4 1-4 16q0 5 .5 16.5t.5 18t-1 16.5t-3 15.5t-6 11t-10 7.5t-15 1q-1-5 3-10.5t5-7.5q5-8 3-38.5t0-41.5q-9 6-7 33.5t-2 41.5q-4 9-15.5 12.5T160 355q1-5 7-9t8-8q4-1 2.5-20.5T177 292q-46 9-63-19q-1-3-5-12t-8-14q-2-2-9-7t-9-11q19-5 37.5 16.5T154 268q6 0 24-6q5-18 16-26q-70-11-89-49q-55 3-102 15q31-12 101-16q-1-4-3.5-6t-6.5-2.5t-7.5-1t-9 0t-7.5.5q-38 2-68 7q22-8 96-8q-9-24-4.5-58T115 70q-5-13-4-30t8-27q26 1 58 23q33-8 69-7q5 0 16.5 3t16.5 2t15-7t11-7q15-7 33-7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Globe(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M237 347q71 0 121-50.5T408 176T358 55.5T237 5q-70 0-120 50T67 176t50 121t120 50m0-299q53 0 90.5 37.5T365 176t-37.5 90.5T237 304t-90.5-38t-37.5-90t37.5-90T237 48m-21 361v23q0 9-5 14.5t-11 5.5l-5 1q-18 0-30.5 13T152 496h171q0-17-13-30t-30-13q-2 0-5.5-.5t-9.5-6t-6-14.5v-23q105-11 166-92q5-7 3.5-16t-7.5-14q-7-5-16.5-3.5T391 291q-57 77-154 77q-80 0-136-56T45 176q0-36 17-77q18-34 41-57q14-16 0-30q-15-15-30 0q-33 33-49 70q-21 42-21 94q0 91 61.5 158T216 409"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Goodreads(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M132 327q76-2 107-75h1v75q0 2-.5 10.5T239 353q-3 19-12 40q-12 22-32 34q-20 14-60 16q-37 0-63-20q-25-18-30-61H20q3 55 35 79q32 23 80 23q47 0 74-18q26-17 37-43q11-23 14-48q2-32 2-33V10h-22v69h-1q-13-37-42-57q-28-19-65-19q-62 1-93 48Q6 96 6 164q0 71 31 115q31 46 95 48M54 67q25-41 78-43q56 2 81 42q26 39 26 98t-26 98q-26 43-81 44q-51-2-78-43t-27-99q0-54 27-97"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Google(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M303 2H178Q110 2 70 35q-39 33-39 81q0 38 27.5 66.5T132 211q3 0 9-.5t10-.5q-6 14-6 23q0 16 18 41q-79 6-115 28q-47 27-47 74q0 36 33.5 61t96.5 25q74 0 116.5-34t42.5-80q0-20-8.5-36.5t-16-24.5t-25.5-24l-22-17q-16-11-16-26q0-12 17-29q17-13 25.5-21.5t17-26T270 105q0-45-43-82h37zm-53 370q0 29-23.5 47T161 437q-48 0-77.5-20.5T54 363q0-42 52-62q29-10 64-10q9 0 14 1q38 26 52 42t14 38m-49-195q-15 17-41 17q-35 0-55-35.5T85 86q0-28 13-45q16-19 42-19q34 0 55.5 36.5T217 133q0 28-16 44"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleBuzz(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M462 190q0 58-41 105L283 156l114-96q65 55 65 130m-294 39L383 50Q317 5 232 5q-62 0-114 24T35 94zM64 316l91-76L25 109Q2 146 2 190q0 74 62 126m345-8L270 167l-101 84l-2 2v1l-30 26h-1l-42 33l-16 14q6 4 10 6q-13 36-37 71q50 21 119-37q28 7 62 7q108 0 177-66"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleTalk(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M460 187q0-76-67.5-129.5T230 4T67.5 57.5T0 187q0 65 51.5 115.5T181 367q-4 51-21 97q63-16 126-99q76-15 125-64.5T460 187M329 300q-6 2-10 2h-18q59-43 59-115q0-79-70-121q6-2 19-2q46 0 78.5 34.5T420 183q0 44-26 77t-65 40"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gowalla(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M216 209q-1 9 0 47t-4 42q-60 7-83-59q-22-62 2-127q24-63 73-59q4 0 8 1q-12 53 24 66q28 12 51-11q13-13 14-32q1-31-26-50T219 4Q146-7 80.5 35T6 159q-8 81 19.5 120T126 353q-12 37-4 62t31.5 36t51 11.5T260 452t45-30q19-21 23.5-59t1.5-74.5t2-79.5zm-35 190q-24-16-10.5-41.5T208 329q7 0 8 4q5 59-35 66"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GowallaAlt(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M125 3q3-1 10-1q21 2 23 32q2 25-16 89q85 18 122 86q22 36 25 78q0 1-1 7t0 9t12 14q25 27 42 32q21 3 28.5 3t30.5-6.5t28-7.5q27-6 32 11q1 2 1 8q-4 20-40 38q-77 41-205 39q-17 16-51 24q-20 2-32 4q-56 2-70-12t-9-32q5-24 50-36q-1-3-4-9.5t-4-8.5q-15-41 2-80q-3-1-9-2.5t-9-2.5q-1 3-2 8.5t-3 7.5q-10 14-27.5 8T32 277q1-15 38-38q4-2 13-7t13-8q18-11 20-18q3-13-4-25l-7 7q-7 8-24 23q-25 20-38 24q-30 9-40-19q-1-4-1-10q8-44 74-76q-1-2-9-10.5T55 106T45 92.5T35.5 75T31 56q-2-25 16-36q6-3 11.5-3.5t11 2.5t9 5.5t9 7T95 38q11-27 30-35m102 177q-33-35-86-44q-13-4-13-5q-1-4 3-14q13-48 15-71q1-15-3.5-25.5T127 15q-14 7-23 35q7 6 13.5 21.5T122 101t-18 14t-18-26t4-38l-6-6q-5-6-9-9q-12-9-21-6q-8 3-11 15q-1 4 2 18q7 29 37 55q1 1 5.5 3.5t7 5T97 132q0 2-22 11q-4 2-8 4.5t-7.5 5T54 156q-26 19-34 32q-8 14-6 22q3 19 24 14q16-4 41-27l5-5l6-5.5l6-6.5q1-1 4.5-6t6.5-7.5t6-2.5q4 0 8.5 8.5T127 187q5 23-9 37q-6 5-20.5 13T77 249q-34 20-33 31q0 10 10 11.5t13-5.5q1-3-.5-11.5T68 266q2-2 18 2l10 2.5l8 1.5q2-2 7-9t8-11t9-7.5t12-3.5q12 0 17.5 15t-3.5 23q-4 4-9.5 5.5t-10 1.5t-12-.5t-10.5-.5q-11 16-9 47t15 47q16-2 45 0q-23-31-11-69t49-43q10-2 21 0t18 4t14.5 6.5t11 7t7.5 5.5l4 4q-5-69-50-114m-121-77q8 0 3-19q-2-6-9-18v-1q-3 6-2 22.5t8 15.5m12 171q24 0 29-6q1-8-5-14q-13 0-24 20m85 5q-31 5-40.5 37t10.5 57q2 2 7 6t8.5 7.5t3.5 5.5q-1 7-13 3q-26-7-50-4q25 29 83 31q9-7 11-20.5t-5-26.5l-2-4l-2-3l-2-2.5l-1-2.5v-3q1-3 16-3q37-4 47-48q-30-27-48-30q-11-2-23 0m26 90q14 20 0 50l-1 3q149-4 208-49q16-11 14-21q-2-8-39 4q-30 7-46 8q-14 1-27-3.5t-18.5-9t-18.5-17t-15-14.5q-2-4-2-1q-14 44-55 50m-109 31q-2-1-3-3t-2-3t-2.5-1.5t-3.5.5q-43 14-43 33q0 16 17 21q16 6 57 2q39-4 57-16q-52-7-77-33"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Grooveshark(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M232 2Q137 2 69.5 69.5T2 232t67.5 162.5T232 462t162.5-67.5T462 232T394.5 69.5T232 2m124 301q-8 0-17-9t-13-18l-5-9q-1-4-3.5-10t-13-22t-25-29.5T238 179t-59-18q2 33 .5 58t-6 39.5t-11 25t-14 13.5t-15 5t-14 1t-11.5 0q-16 1-25.5-19.5T73 232q0-66 46.5-112.5T232 73t111.5 46.5T391 232q1 16-11 43.5T356 303"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gun(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M128 323q0 8-6.5 14.5T107 344q-9 0-15.5-6.5T85 323q0-9 6.5-15.5T107 301q8 0 14.5 6.5T128 323M469 24H192v2q-42 11-64 41q-8 0-13.5-10.5T108 35l-1-11H64q0 24 11 49t34 34q0 2-1 6t-1 5q-46 21-64 70q-13 36-24 79.5T4 335t-4 28q0 30 13 45q15 21 51 21h64q29 0 46.5-17.5T192 365q0-24 4-64h39q29 0 52-18t31-46h23q51 0 73-42h45q10 0 10-11t-10-11h-34q0-3 1-10.5t1-10.5h85V3h-21q-22 0-22 21M192 71v119q-19-6-31-22.5T149 131q0-21 12-37.5T192 71m-64 316H64q-15 0-19-5q-2-4-2-15q12-89 40-166q10-25 28-38q12 41 53 62q-15 94-15 140q0 22-21 22m107-128h-32q4-18 4-22h64q-10 22-36 22m106-64H235V67h149v21h-96q-11 0-11 11q0 10 11 10h96v22h-96q-11 0-11 10q0 11 11 11h96q0 17-12.5 30T341 195m128-86h-42V67h42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Guy(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 171q-10-73-64.5-122T256 0T128.5 49.5T66 171h-2q-27 0-45.5 18T0 235q0 33 17.5 59T64 333q5 75 60.5 127T256 512t131.5-52T448 333q29-13 46.5-39t17.5-59q0-28-18.5-46T448 171M64 213v69q-21-18-21-47q0-22 21-22m192 256q-62 0-105.5-43.5T107 320V192q0-15 2-23q53-10 77-71q69 48 213 51q6 20 6 43v128q0 62-43.5 105.5T256 469m192-187v-69q21 0 21 22q0 29-21 47m-192 59q21 0 21-21v-21h-42v21q0 21 21 21m0 64q17 0 30-12.5t13-29.5h-86q0 17 13 29.5t30 12.5m-64-170q0 8-6.5 14.5T171 256q-9 0-15.5-6.5T149 235q0-9 6.5-15.5T171 213q8 0 14.5 6.5T192 235m171 0q0 8-6.5 14.5T341 256q-8 0-14.5-6.5T320 235q0-9 6.5-15.5T341 213q9 0 15.5 6.5T363 235"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GuyAngel(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 171q-6-46-32-86h32q21 0 21-21t-21-21h-73Q322 0 256 0T137 43H64q-21 0-21 21t21 21h32q-27 42-30 86h-2q-27 0-45.5 18T0 235q0 33 17.5 59T64 333q5 75 60.5 127T256 512t131.5-52T448 333q29-13 46.5-39t17.5-59q0-28-18.5-46T448 171M64 213v64q-21-16-21-42q0-22 21-22m192 256q-62 0-105.5-43.5T107 320V192q0-61 44-107h18q64 62 230 64q6 20 6 43v128q0 62-43.5 105.5T256 469m192-192v-64q21 0 21 22q0 26-21 42m-213 22q0 21 21 21t21-21v-22h-42zm21-86h-43q0 22-21 22q-5 0-8.5-1.5T178 231l-1-1q-6-6-6-17h-43q0 28 17 47q22 17 47 17q32 0 45-17q19-19 19-47m64 22q-5 0-8.5-1.5T306 231l-1-1q-6-6-6-17h-43q0 28 17 47q17 17 47 17q32 0 45-17q19-19 19-47h-43q0 22-21 22m-43 149q0 9-6 15t-15 6t-15-6t-6-15t6-15t15-6t15 6t6 15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GuyAngry(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 171q-10-73-64.5-122T256 0T128.5 49.5T66 171h-2q-27 0-45.5 18T0 235q0 33 17.5 59T64 333q5 75 60.5 127T256 512t131.5-52T448 333q29-13 46.5-39t17.5-59q0-28-18.5-46T448 171M64 213v69q-21-18-21-47q0-22 21-22m192 256q-62 0-105.5-43.5T107 320V192q0-15 2-23q53-10 77-71q69 48 213 51q6 20 6 43v128q0 62-43.5 105.5T256 469m192-187v-69q21 0 21 22q0 29-21 47m-192 59q21 0 21-21v-21h-42v21q0 21 21 21m21 22h-42q-18 0-30.5 12.5T192 405q0-8 25-14.5t39-6.5t39 6.5t25 14.5q0-17-12.5-29.5T277 363m-64-107v-43h-85l43 22v21q0 21 21 21t21-21m86-43v43q0 21 21 21t21-21v-21l43-22z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GuyBigSmile(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 171q-10-73-64.5-122T256 0T128.5 49.5T66 171h-2q-27 0-45.5 18T0 235q0 33 17.5 59T64 333q5 75 60.5 127T256 512t131.5-52T448 333q29-13 46.5-39t17.5-59q0-28-18.5-46T448 171M64 213v69q-21-18-21-47q0-22 21-22m192 256q-62 0-105.5-43.5T107 320V192q0-15 2-23q53-10 77-71q69 48 213 51q6 20 6 43v128q0 62-43.5 105.5T256 469m192-187v-69q21 0 21 22q0 29-21 47m-192 38q21 0 21-21v-22h-42v22q0 21 21 21m87 0q-20-3-23 19q0 4-3 13t-19 20.5t-42 11.5q-24 0-39.5-11.5t-20-21T192 337q-4-20-26-17q-8 2-13 9.5t-4 16.5q5 28 32 54.5t75 26.5q51 0 76.5-26.5T363 343q2-8-4-15t-16-8m-2-64q9 0 15.5-6t6.5-15v-22h-43v22q0 9 6 15t15 6m-170 0q8 0 14.5-6t6.5-15v-22h-43v22q0 9 6 15t16 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GuyConfused(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 171q-10-73-64.5-122T256 0T128.5 49.5T66 171h-2q-27 0-45.5 18T0 235q0 33 17.5 59T64 333q5 75 60.5 127T256 512t131.5-52T448 333q29-13 46.5-39t17.5-59q0-28-18.5-46T448 171M64 213v69q-21-18-21-47q0-22 21-22m192 256q-62 0-105.5-43.5T107 320V192q0-15 2-23q53-10 77-71q69 48 213 51q6 20 6 43v128q0 62-43.5 105.5T256 469m192-187v-69q21 0 21 22q0 29-21 47m-192 59q21 0 21-21v-21h-42v21q0 21 21 21m77 24l-22 10q-11 10-23 0q-31-16-62 0q-12 10-23 0l-22-10q-8-3-16-.5t-11 8.5q-3 8-.5 16.5T162 401l21 11q32 16 62 0q12-9 24 0q15 6 30 6q16 0 30-6l21-11q8-3 10.5-11.5T358 373q-3-6-11-8.5t-14 .5M192 235q0 8-6.5 14.5T171 256q-9 0-15.5-6.5T149 235q0-9 6.5-15.5T171 213q8 0 14.5 6.5T192 235m171 0q0 8-6.5 14.5T341 256q-8 0-14.5-6.5T320 235q0-9 6.5-15.5T341 213q9 0 15.5 6.5T363 235"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GuyCry(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 171q-10-73-64.5-122T256 0T128.5 49.5T66 171h-2q-27 0-45.5 18T0 235q0 33 17.5 59T64 333q5 75 60.5 127T256 512t131.5-52T448 333q29-13 46.5-39t17.5-59q0-28-18.5-46T448 171M64 213v69q-21-18-21-47q0-22 21-22m192 256q-62 0-105.5-43.5T107 320V192q0-15 2-23q53-10 77-71q69 48 213 51q6 20 6 43v128q0 62-43.5 105.5T256 469m192-187v-69q21 0 21 22q0 29-21 47m-192 38q21 0 21-21v-22h-42v22q0 21 21 21m85-107q-8 0-14.5 6.5T320 235q0 8 6 14.5t15 6.5t15.5-6t6.5-15t-6.5-15.5T341 213m-21 107q0 8 6 14.5t15 6.5q22 0 22-21v-21h-43zm-149-64q8 0 14.5-6t6.5-15t-6-15.5t-15-6.5t-15.5 6.5T149 235t6.5 15t15.5 6m21 64v-21h-43v21q0 8 6.5 14.5T171 341q21 0 21-21m64 43q-50 0-79 30q-6 6-6 14.5t6 14.5q15 15 30 0q19-17 49-17q27 0 49 17q6 7 15 7t15-7q6-6 6-14.5t-6-14.5q-29-30-79-30"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GuyFlushed(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 171q-10-73-64.5-122T256 0T128.5 49.5T66 171h-2q-27 0-45.5 18T0 235q0 33 17.5 59T64 333q5 75 60.5 127T256 512t131.5-52T448 333q29-13 46.5-39t17.5-59q0-28-18.5-46T448 171M64 213v69q-21-18-21-47q0-22 21-22m192 256q-62 0-105.5-43.5T107 320V192q0-15 2-23q53-10 77-71q69 48 213 51q6 20 6 43v128q0 62-43.5 105.5T256 469m192-187v-69q21 0 21 22q0 29-21 47m-213 38q0 21 21 21t21-21v-21h-42zm42 43h-42q-18 0-30.5 12.5T192 405q0-8 25-14.5t39-6.5t39 6.5t25 14.5q0-17-12.5-29.5T277 363m-42-118q0-22-15.5-37.5T181 192t-38 15t-15 38t15.5 38.5T181 299q23 0 38.5-15.5T235 245m-54 32q-13 0-22.5-9.5T149 245t9.5-22.5T181 213t22.5 9.5T213 245t-9.5 22.5T181 277m150-85q-23 0-38.5 15.5T277 245q0 23 15.5 38.5T331 299q22 0 37.5-15.5T384 245t-15-38t-38-15m0 85q-13 0-22.5-9.5T299 245t9.5-22.5T331 213t22.5 9.5T363 245t-9.5 22.5T331 277"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GuyHappy(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 171q-10-73-64.5-122T256 0T128.5 49.5T66 171h-2q-27 0-45.5 18T0 235q0 33 17.5 59T64 333q5 75 60.5 127T256 512t131.5-52T448 333q29-13 46.5-39t17.5-59q0-28-18.5-46T448 171M64 213v69q-21-18-21-47q0-22 21-22m192 256q-62 0-105.5-43.5T107 320V192q0-15 2-23q53-10 77-71q69 48 213 51q6 20 6 43v128q0 62-43.5 105.5T256 469m192-187v-69q21 0 21 22q0 29-21 47m-192 59q21 0 21-21v-21h-42v21q0 21 21 21m-64-106q0 8-6.5 14.5T171 256q-9 0-15.5-6.5T149 235q0-9 6.5-15.5T171 213q8 0 14.5 6.5T192 235m171 0q0 8-6.5 14.5T341 256q-8 0-14.5-6.5T320 235q0-9 6.5-15.5T341 213q9 0 15.5 6.5T363 235M235 405h42q18 0 30.5-12.5T320 363H192q0 17 12.5 29.5T235 405"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GuyOmouth(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 171q-10-73-64.5-122T256 0T128.5 49.5T66 171h-2q-27 0-45.5 18T0 235q0 33 17.5 59T64 333q5 75 60.5 127T256 512t131.5-52T448 333q29-13 46.5-39t17.5-59q0-28-18.5-46T448 171M64 213v69q-21-18-21-47q0-22 21-22m192 256q-62 0-105.5-43.5T107 320V192q0-69 53-115q70 70 239 70q6 23 6 45v128q0 62-43.5 105.5T256 469m192-187v-69q21 0 21 22q0 29-21 47m-192 38q21 0 21-21v-22h-42v22q0 21 21 21m-64-85q0 8-6.5 14.5T171 256q-9 0-15.5-6.5T149 235q0-9 6.5-15.5T171 213q8 0 14.5 6.5T192 235m171 0q0 8-6.5 14.5T341 256q-8 0-14.5-6.5T320 235q0-9 6.5-15.5T341 213q9 0 15.5 6.5T363 235m-64 170q0 18-12.5 30.5T256 448t-30.5-12.5T213 405q0-17 12.5-29.5T256 363t30.5 12.5T299 405"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GuyOpenMouth(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 171q-10-73-64.5-122T256 0T128.5 49.5T66 171h-2q-27 0-45.5 18T0 235q0 33 17.5 59T64 333q5 75 60.5 127T256 512t131.5-52T448 333q29-13 46.5-39t17.5-59q0-28-18.5-46T448 171M64 213v69q-21-18-21-47q0-22 21-22m192 256q-62 0-105.5-43.5T107 320V192q0-15 2-23q53-10 77-71q69 48 213 51q6 20 6 43v128q0 62-43.5 105.5T256 469m192-187v-69q21 0 21 22q0 29-21 47m-192 38q21 0 21-21v-22h-42v22q0 21 21 21m-64-85q0 8-6.5 14.5T171 256q-9 0-15.5-6.5T149 235q0-9 6.5-15.5T171 213q8 0 14.5 6.5T192 235m171 0q0 8-6.5 14.5T341 256q-8 0-14.5-6.5T320 235q0-9 6.5-15.5T341 213q9 0 15.5 6.5T363 235M213 384q0 17 13 30t30 13t30-13t13-30v-43h-86z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GuySad(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 171q-10-73-64.5-122T256 0T128.5 49.5T66 171h-2q-27 0-45.5 18T0 235q0 33 17.5 59T64 333q5 75 60.5 127T256 512t131.5-52T448 333q29-13 46.5-39t17.5-59q0-28-18.5-46T448 171M64 213v69q-21-18-21-47q0-22 21-22m192 256q-62 0-105.5-43.5T107 320V192q0-15 2-23q53-10 77-71q69 48 213 51q6 20 6 43v128q0 62-43.5 105.5T256 469m192-187v-69q21 0 21 22q0 29-21 47m-192 59q21 0 21-21v-21h-42v21q0 21 21 21m0 22q-17 0-30 12.5T213 405h86q0-17-13-29.5T256 363m-64-128q0 8-6.5 14.5T171 256q-9 0-15.5-6.5T149 235q0-9 6.5-15.5T171 213q8 0 14.5 6.5T192 235m171 0q0 8-6.5 14.5T341 256q-8 0-14.5-6.5T320 235q0-9 6.5-15.5T341 213q9 0 15.5 6.5T363 235"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GuySleep(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 171q-10-73-64.5-122T256 0T128.5 49.5T66 171h-2q-27 0-45.5 18T0 235q0 33 17.5 59T64 333q5 75 60.5 127T256 512t131.5-52T448 333q29-13 46.5-39t17.5-59q0-28-18.5-46T448 171M64 213v69q-21-18-21-47q0-22 21-22m192 256q-62 0-105.5-43.5T107 320V192q0-69 53-115q70 70 239 70q6 23 6 45v128q0 62-43.5 105.5T256 469m192-187v-69q21 0 21 22q0 29-21 47m-213 17q0 21 21 21t21-21v-22h-42zm21-86h-43q0 22-21 22q-5 0-8.5-1.5T178 231l-1-1q-6-6-6-17h-43q0 28 17 47q22 17 47 17q32 0 45-17q19-19 19-47m64 22q-5 0-8.5-1.5T306 231l-1-1q-6-6-6-17h-43q0 28 17 47q17 17 47 17q32 0 45-17q19-19 19-47h-43q0 22-21 22m-43 149q0 9-6 15t-15 6t-15-6t-6-15t6-15t15-6t15 6t6 15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GuySmile(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 171q-10-73-64.5-122T256 0T128.5 49.5T66 171h-2q-27 0-45.5 18T0 235q0 33 17.5 59T64 333q5 75 60.5 127T256 512t131.5-52T448 333q29-13 46.5-39t17.5-59q0-28-18.5-46T448 171M64 213v69q-21-18-21-47q0-22 21-22m192 256q-62 0-105.5-43.5T107 320V192q0-15 2-23q53-10 77-71q69 48 213 51q6 20 6 43v128q0 62-43.5 105.5T256 469m192-187v-69q21 0 21 22q0 29-21 47m-192 38q21 0 21-21v-22h-42v22q0 21 21 21m-64-85q0 8-6.5 14.5T171 256q-9 0-15.5-6.5T149 235q0-9 6.5-15.5T171 213q8 0 14.5 6.5T192 235m171 0q0 8-6.5 14.5T341 256q-8 0-14.5-6.5T320 235q0-9 6.5-15.5T341 213q9 0 15.5 6.5T363 235M256 405q38 0 59-16t23-32l3-16H171q0 3 .5 7t5 15t12.5 19.5t25.5 15.5t41.5 7"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GuyUser(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 448v64h320v-64q0-52-30-97.5T216 288v-23q29-15 46.5-43.5T280 160v-43q0-48-34.5-82.5T163 0T80 34.5T45 117v43q0 33 17.5 61.5T109 265v23q-46 17-76 62T3 448m277 0v21H45v-21q0-36 19.5-67.5T116 333q14 30 47 30q31 0 47-30q31 16 50.5 48t19.5 67m-45-341q-77-6-121-45q22-19 49-19t48 18.5t24 45.5M88 117q0-9 4-21q61 48 145 53v11q0 31-22 53t-52 22q-31 0-53-22t-22-53zm75 160h10v32q0 11-10 11q-11 0-11-11v-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GuyWrong(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 171q-10-73-64.5-122T256 0T128.5 49.5T66 171h-2q-27 0-45.5 18T0 235q0 33 17.5 59T64 333q5 75 60.5 127T256 512t131.5-52T448 333q29-13 46.5-39t17.5-59q0-28-18.5-46T448 171M64 213v69q-21-18-21-47q0-22 21-22m192 256q-62 0-105.5-43.5T107 320V192q0-15 2-23q53-10 77-71q69 48 213 51q6 20 6 43v128q0 62-43.5 105.5T256 469m192-187v-69q21 0 21 22q0 29-21 47m-192 59q21 0 21-21v-21h-42v21q0 21 21 21m-64-106q0 8-6.5 14.5T171 256q-9 0-15.5-6.5T149 235q0-9 6.5-15.5T171 213q8 0 14.5 6.5T192 235m171 0q0 8-6.5 14.5T341 256q-8 0-14.5-6.5T320 235q0-9 6.5-15.5T341 213q9 0 15.5 6.5T363 235m-86 128h-42q-18 0-30.5 12.5T192 405q0-8 25-14.5t39-6.5t39 6.5t25 14.5q0-17-12.5-29.5T277 363"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HackerNews(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 2v460h460V2zm250 265v113h-40V267L111 68h47l74 150l76-150h45z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hand(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 197v124q0 58 36.5 102t93.5 52q69 7 119-37.5T299 327V155q0-22-22-22q-9 0-15 6t-6 16v170q0 47-34.5 79T139 432q-42-4-69-36.5T43 321v-81q0-17-13-30T0 197M235 69q0-9-6-15t-16-6q-9 0-15 6t-6 15v192h43zm-64-42q0-22-22-22q-9 0-15 6t-6 16v234h43zm-64 64q0-22-22-22q-9 0-15 6t-6 16v170h43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandPointerLeft(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M237 341h43q68 0 112.5-50.5T427 173q-6-55-50.5-92.5T276 43H152q0 17 12.5 29.5T195 85h81q42 0 74 27.5t37 68.5q4 48-28 83t-79 35h-85q0 17 12.5 29.5T237 341m-85-64h64q21 0 21-21t-21-21h-64q-21 0-21 21t21 21m0-64h64q21 0 21-21t-21-21h-64q-21 0-21 21t21 21M3 128q0 21 21 21h192v-42H24q-21 0-21 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandPointerRight(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M152 341h43q17 0 29.5-12.5T237 299h-85q-47 0-79-35t-28-83q5-41 37-68.5T156 85h81q18 0 30.5-12.5T280 43H156Q98 43 54 80T3 173q-8 67 37 117.5T152 341m128-106h-64q-21 0-21 21t21 21h64q21 0 21-21t-21-21m0-64h-64q-21 0-21 21t21 21h64q21 0 21-21t-21-21m149-43q0-21-21-21H216v42h192q21 0 21-21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandPointerTop(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 152v124q0 58 37 102t93 51q68 8 118-36t50-111v-43q0-17-12.5-29.5T259 197v85q0 47-35 79t-83 28q-41-7-68.5-39.5T45 276v-81q0-18-12.5-30.5T3 152m234 85v-64q0-21-21-21t-21 21v64q0 22 21 22t21-22m-64-21v-64q0-21-21-21t-21 21v64q0 21 21 21t21-21M109 24q0-21-21-21T67 24v192h42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HandWash(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M58 461q4 22 21.5 36.5T119 512h274q22 0 39.5-14.5T454 461l58-265q2-8-3-16t-14-9q-9-2-17 3t-9 14l-57 264q-3 17-22 17H119q-18 0-21-17L43 188q-1-8-9.5-13.5T17 171q-8 1-13.5 9T0 196zm177-77q0 21 21 21t21-21V171h-42zm64-43q0 22 21 22t21-22V171h-42zm-171-85q9 0 15-6l28-28v98q0 21 21 21t21-21V117q0-36 27.5-58T303 45q27 4 43.5 25.5T363 119v180q0 21 21 21t21-21V117q0-51-38-86T277 0q-46 4-76 39.5T171 122v38l-58 58q-6 6-6 14.5t6 14.5q5 9 15 9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HangToDry(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M423 21Q412 0 387 0H45Q20 0 9 21H3v448q0 18 12.5 30.5T45 512h342q17 0 29.5-12.5T429 469V21zm-41 22q-13 37-59.5 61T216 128q-61 0-108-24T50 43zM45 469V111q65 60 171 60t171-60v358z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HardDrive(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M475 448V320L411 43Q401 0 368 0H112Q79 0 69 43L5 320v128q0 27 18.5 45.5T69 512h342q27 0 45.5-18.5T475 448M103 60q5-17 22-17h230q17 0 22 17l51 241q-7-2-17-2H69q-10 0-17 2zm329 388q0 21-21 21H69q-21 0-21-21v-85q0-10 6-16t15-6h342q9 0 15 6t6 16zm-43-43q0 9-6 15.5t-15 6.5t-15-6.5t-6-15.5q0-8 6-14.5t15-6.5t15 6.5t6 14.5m-85 0q0 9-6.5 15.5T283 427q-9 0-15.5-6.5T261 405q0-8 6.5-14.5T283 384q8 0 14.5 6.5T304 405"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Headphones(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M425 173q-7-65-55.5-108.5T256 21T142.5 64.5T87 173q-38 6-62.5 32.5T0 267q0 40 31 68t76 28h21V192q0-52 38-90t90-38t90 38t38 90v171h21q45 0 76-28t31-68q0-35-24.5-61.5T425 173M85 318q-19-5-30.5-19.5T43 267q0-37 42-52zm342 0V218q19 4 30.5 18.5T469 269q0 34-42 49"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Headset(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M425 151q-7-64-55.5-107.5T256 0T142.5 43.5T87 151q-38 7-62.5 33T0 245q0 40 31 68t76 28h21V171q0-53 37.5-90.5T256 43t90.5 37.5T384 171v149q0 86-68 105q-19-41-60-41q-27 0-45.5 18.5T192 448t18.5 45.5T256 512q21 0 38-12.5t22-32.5q43-9 70-36q35-35 39-92q37-6 61-32.5t24-61.5q1-35-22.5-61T425 151M85 297q-19-5-30.5-19.5T43 245q0-36 42-51zm171 172q-21 0-21-21t21-21t21 21t-21 21m171-172V194q19 5 30.5 19.5T469 245q0 37-42 52"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartsCard(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 512h256q27 0 45.5-18.5T384 448V64q0-27-18.5-45.5T320 0H64Q37 0 18.5 18.5T0 64v384q0 27 18.5 45.5T64 512M43 64q0-21 21-21h256q21 0 21 21v384q0 21-21 21H64q-21 0-21-21zm66 160q5 52 83 96q78-44 83-96v-11q0-17-12.5-29.5T233 171q-18 0-30.5 12.5T190 213q0-17-12.5-29.5T147 171q-17 0-29.5 12.5T105 213v5q1 1 2 3t2 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heat(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469 0H43Q25 0 12.5 12.5T0 43v426q0 18 12.5 30.5T43 512h426q18 0 30.5-12.5T512 469V43q0-18-12.5-30.5T469 0m0 469H43V43h426zm-213-23q79 0 135.5-56T448 254t-56-136t-136-56t-136 56t-56 136t56 136t136 56m0-339q61 0 105 43.5T405 256t-43.5 105.5T256 405t-105.5-43.5T107 256t43.5-105.5T256 107m-21 149q0 9-6.5 15t-15.5 6q-8 0-14.5-6t-6.5-15t6.5-15t14.5-6q9 0 15.5 6t6.5 15m85 0q0 9-6.5 15t-14.5 6q-9 0-15.5-6t-6.5-15t6.5-15t15.5-6q8 0 14.5 6t6.5 15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Helm(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M491 235h-45q-6-55-41-98l37-37q14-14 0-30q-16-14-30 0l-37 37q-43-35-98-41V21q0-21-21-21t-21 21v45q-59 6-98 41l-37-37q-14-14-30 0q-14 16 0 30l37 37q-35 43-41 98H21q-21 0-21 21t21 21h45q6 55 41 98l-37 37q-14 14 0 30q8 6 15 6q9 0 15-6l37-37q43 35 98 41v45q0 21 21 21t21-21v-45q55-6 98-41l37 37q6 6 15 6q7 0 15-6q14-16 0-30l-37-37q35-43 41-98h45q21 0 21-21t-21-21m-88 0h-96l68-69q22 29 28 69m-57-98l-69 68v-96q40 6 69 28m-111-28v96l-69-68q29-22 69-28m-98 57l68 69h-96q6-40 28-69m-28 111h96l-68 69q-22-29-28-69m57 98l69-68v96q-40-6-69-28m111 28v-96l69 68q-29 22-69 28m98-57l-68-69h96q-6 40-28 69"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HiFive(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m446 164l4-81H295l-23 126h1q-2 5-4 13v20h15v29q28 30 80 30q43 0 70.5-25t27.5-65q0-27-16-47m-81 117q-34 0-55-15t-21-34q0-13 8.5-21.5T320 202q20 0 20 19q0 3-1 7t-1 6q0 11 14 11q22 0 22-32q0-25-17-25q-13 0-15 9l-44-5l16-89h117l-2 45h-78l-4 23q9-14 41-14q26 0 41 16.5t15 37.5q0 30-21 50t-58 20m-76-59V96h-85v38h-28v75h15v13h-15v74h128v-74zm-65-106h45v29h-45zm60 160h-88v-34h15v-53h-15v-34h73v87h15zm-88 0v-34h15v-20h-16v-30q0-25-16-42.5T140 132q-13 0-25 5V83H2v75h15v64H2v74h209v-20zm-5 0h-80v-34h6v-35q0-20-6-20q-10 0-16 23v32h6v34H22v-34h15V138H22v-35h73v78q15-29 45-29q15 0 25 11.5t10 28.5v50h16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Home(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 7L6 257l30 30l28-28v226q0 18 12.5 30.5T107 528h298q18 0 30.5-12.5T448 485V259l30 30l30-30L271 25zm-43 478V336h86v149zm128 0V336q0-17-12.5-30T299 293h-86q-17 0-29.5 13T171 336v149h-64V217L256 67l149 150v268zm-85-341q-27 0-45.5 18.5T192 208t18.5 45.5T256 272t45.5-18.5T320 208t-18.5-45.5T256 144m0 85q-21 0-21-21t21-21t21 21t-21 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hotdog(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 384q0 35 25.5 60T91 469q31 43 85 43t85-43q35 0 60.5-25t25.5-60V149q0-31-21-54.5T274 66q-11-29-38-47.5T176 0t-60 18.5T78 66q-31 5-52 28.5T5 149zm299-235v235q0 26-21 36V113q21 11 21 36M176 43q27 0 45.5 18t18.5 46v298q0 28-18.5 46T176 469t-45.5-18t-18.5-46V107q0-28 18.5-46T176 43M48 149q0-25 21-36v307q-21-10-21-36zm115 209q-30 21-30 47q0 22 22 22q9 0 15-6t6-16q0-5 13-12q30-21 30-47q0-23-30-47q-13-8-13-13q0-2 13-13q30-18 30-47q0-7-2-13t-6.5-12t-6.5-8.5t-8-8.5l-7-5q-13-9-13-13q0-3 13-12l7-6q6-5 8-7.5t6.5-8.5t6.5-12t2-13q0-22-22-22q-9 0-15 6t-6 16q0 3-13 12l-7 5q-6 6-8 8.5t-6.5 8.5t-6.5 12t-2 13q0 23 30 47q13 8 13 13q0 2-13 13q-30 18-30 47q0 7 2 13t6.5 12t6.5 8.5t8 8.5l7 5q13 7 13 13q0 3-13 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hourglass(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M48 363v106H27q-22 0-22 22q0 9 6 15t16 6h298q10 0 16-6t6-15q0-22-22-22h-21V363q0-69-58-107q58-38 58-107V43h21q22 0 22-22q0-9-6-15t-16-6H27Q17 0 11 6T5 21q0 22 22 22h21v106q0 69 58 107q-58 38-58 107m43-214V43h170v106q0 35-25 60.5T176 235t-60-25.5T91 149m0 214q0-35 25-60.5t60-25.5t60 25.5t25 60.5v106H91zm149-256l-128 42v22q1 5 4.5 12t19.5 18.5t40 11.5t40-10.5t20-20.5l4-11zM112 448h128v-85l-128 21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hungry(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 171q0-47-23-86t-62-62v-2h-2Q319 0 277 0h-42q-39 0-84 21h-2v2q-39 23-62 62t-23 86q-27 0-45.5 18T0 235q0 33 17.5 59T64 333q5 75 60.5 127T256 512t131.5-52T448 333q29-13 46.5-39t17.5-59q0-28-18.5-46T448 171M64 282q-21-18-21-47q0-22 21-22zm341 38q0 62-43.5 105.5T256 469t-105.5-43.5T107 320V171q0-57 42-96v10q94 38 214 0V75q42 39 42 96zm43-38v-69q21 0 21 22q0 29-21 47m-171 81h-42q-22 0-22 21t22 21h42q22 0 22-21t-22-21m-42-86h-43q0 18 12.5 30.5T235 320zm42 43q18 0 30.5-12.5T320 277h-43zm34-90q6-4 9-4v9q0 9 6 15t15 6q10 0 16-6t6-15v-22h21v-42q-69 0-100 25zm-119-4q2 1 4.5 2t4.5 2l25-34q-35-25-98-25v42q13 0 21 2v20q0 9 6 15t16 6q9 0 15-6t6-15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HypeMachine(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M238 3Q143 3 75 53.5T7 176q0 77 78 129q-10 49-29 76q28 0 108-42q35 9 74 9q95 0 162.5-50.5T468 176q0-72-67.5-122.5T238 3m110 175q-3 9-4 11q-10 14-17 20q-14 15-36.5 33T253 270l-15 11q-57-38-90-72q-11-11-16-20q-2-3-5-11q-3-10-3-22q0-28 20-48.5T193 87q18 0 31.5 15t13.5 31q0-18 10-32t34-14q29 0 49.5 20t20.5 49q0 8-4 22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hyves(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M318 2h32q8-1 17 16.5t6 27.5q-2 5-16 11t-16 9q-4 5-5 15t-.5 32t.5 23v191q0 34 1 35q2 2 6 2.5t8.5.5h6.5q9 3 15.5 16.5t.5 23.5q-2 2-22 10q-40 20-76 17q-12-1-25-14.5T240 393q1-9 15-14q1 0 10-1.5t12-4.5q3-2 3-9.5t-1-18.5t-1-15v-24q0-20-1-23q-4-5-46 0q-52 8-82 17q-1 1-10 3t-15.5 5t-8.5 7q-3 5 0 54q0 22 1 23q2 4 12 4.5t11 .5q8 3 14.5 12t5.5 18q-1 16-29.5 24.5T71 462q-14 1-22-1q-10-2-19-13t-8-20q1-12 25-19q11-3 11-4q3-3 0-35q-4-32-6-111q0-21-1.5-56.5T49 145q0-35-1-36q-2-3-11-4.5T26 102Q11 95 3 75v-8q7-13 21-18.5T65 38q28-8 54-6t33 24q5 15 0 23q-4 6-20 9.5T114 97q-2 2-2 29q0 11-2 49t0 50q0 21 6 24q1 1 20-3q62-13 99-18q40-4 41-5q5-4 2-35q-4-37 0-74q3-31-1-35q-2-2-11-3.5T256 73q-33-14-25-41q3-6 16-13q9-5 71-17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Icq(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m453 255l-3-5q-8-12-18-21q-3-2-10.5-5.5T410 218q30-14 41-41q3-8-.5-21t-3.5-16v-2q0-16-35-26q-9-1-30-1l-22 5l7-17q11-30-6-59l-2-4q-15-23-45-31q-30-7-56 7q-36 19-36 44v1l-3 5l-3-5l-7-11q-15-23-41-31q-21-5-39 5q-20 12-25 39q-5 29 10 55l4 7l-15-5q-29-7-57 8q-28 13-39 42q-12 26 2 52q5 9 15 17q13 11 30 14l5 1l-10 9q-13 10-18 27q-5 19 5 34q5 9 12 13q7 7 19 9q24 5 47-2l-12 19l-3 9q-10 30 5 58q2 3 15 19q12 11 30 15q34 8 59-5q26-15 36-45l1-4q19 22 38 26q20 6 38-4q17-11 23-32q4-21 0-44q12 7 21 9q31 10 59-4q26-14 35-45q9-29-6-57m-96-108l8-3q16-8 35-6q18 3 24 15l2 8l-1 7q-7 16-31 28q-7 3-22 5l-81 10l-2-4l-1-2l3-4zM245 65q4-19 24-29q19-11 38-3q21 5 30 21q5 12 5 23l-2 14q-2 9-10 20l-59 71q-6-5-20-11l-6-73l-1-26zm-113 2q0-16 13-24q6-5 15-1q19 6 30 28q8 18 8 21l13 77l-10 3l-59-63l-3-7q-9-20-7-34M34 205q-7-15-1-29q8-21 26-28q19-9 37-5l18 7l58 42q-8 9-11 19l-95 12l-6-1q-16-3-26-17m188 191l-3 7q-6 18-25 26q-17 10-37 5t-28-21q-6-11-6-22l3-15q2-10 10-19l55-65q11 6 22 8q1 7 5 31.5t4 38.5zm-25-130q-14-11-14-30q0-17 14-31q11-11 31-11q19 0 30 11q13 13 13 31q0 20-13 30q-12 13-30 13q-21 0-31-13m122 114q-2 18-13 24q-6 4-15 2q-19-4-30-28l-8-22l-10-56l11-4l9-5l28 28l17 21l3 6q9 17 8 34m114-76q-5 20-23 28q-17 10-37 3q-11-3-18-10l-69-57l5-9l2-8l74-7h33q19 8 25 19q10 10 10 28z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Identi(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 212q-1 85 60.5 144.5T212 416q50 0 92-21l158 47l-51-165q11-33 11-66q0-85-61.5-145T212 6T63.5 66.5T2 212m352-1q0 27-9 49l27 75l2 6l9 25l-33-9l-72-22q-30 15-66 15q-59 0-101-40.5T69 211q0-57 42-98t101-41t100.5 40.5T354 211"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Image(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M528 67q0-28-18.5-46T464 3H80Q53 3 34.5 21T16 67v245l-9 6l9 11v36q0 28 18.5 46T80 429h384q27 0 45.5-18t18.5-46v-98l9-10l-9-9zm-43 298q0 22-21 22H80q-21 0-21-22v-32l128-91l134 89l24-36l-47-32l83-83l104 87zm0-151l-108-90l-118 118l-72-52l-128 92V67q0-22 21-22h384q21 0 21 22zM208 109q0 18-12.5 30.5T165 152q-17 0-29.5-12.5T123 109q0-17 12.5-29.5T165 67q18 0 30.5 12.5T208 109"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Important(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M43 280h42q18 0 30.5-12.5T128 237V45q0-17-12.5-29.5T85 3H43Q25 3 12.5 15.5T0 45v192q0 18 12.5 30.5T43 280m0-235h42v192H43zM0 387q0 17 12.5 29.5T43 429h42q18 0 30.5-12.5T128 387v-22q0-17-12.5-29.5T85 323H43q-18 0-30.5 12.5T0 365zm43-22h42v22H43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Instapaper(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M209 449q-42-3-54-12q-13-10-13-50V77q0-39 13-49q11-11 54-13V2H4v13q43 1 55 13q13 10 13 49v310q0 40-13 50q-12 9-55 12v13h205z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ipad(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 0H43Q25 0 12.5 12.5T0 43v426q0 18 12.5 30.5T43 512h298q18 0 30.5-12.5T384 469V43q0-18-12.5-30.5T341 0m0 469H43v-42h298zm0-85H43V107h298zm0-320H43V43h298zM213 448q0-9-6-15q-15-13-30 0q-6 6-6 15t6 15t15 6t15-6t6-15M64 128h43v43H64zm85 0h43v43h-43zm86 0h42v43h-42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Iphone(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M237 512q28 0 46-18.5t18-45.5V64q0-27-18-45.5T237 0H67Q39 0 21 18.5T3 64v384q0 27 18 45.5T67 512zM67 43h170q22 0 22 21H45q0-21 22-21m-22 64h214v277H45zm0 341v-21h214v21q0 21-22 21H67q-22 0-22-21m128 0q0-9-6-15q-15-13-30 0q-6 6-6 15t6 15t15 6t15-6t6-15M67 128h42v43H67zm85 0h43v43h-43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ipod(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M299 0H85Q50 0 25 25T0 85v342q0 35 25 60t60 25h214q35 0 60-25t25-60V85q0-35-25-60T299 0m42 427q0 17-12.5 29.5T299 469H85q-17 0-29.5-12.5T43 427V85q0-17 12.5-29.5T85 43h214q17 0 29.5 12.5T341 85zM64 256h256V85H64zm43-128h170v85H107zm85 149q-35 0-60 25.5T107 363t25 60t60 25t60-25t25-60t-25-60.5t-60-25.5m0 128q-17 0-30-12.5T149 363q0-18 13-30.5t30-12.5t30 12.5t13 30.5q0 17-13 29.5T192 405"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IronAnyTemp(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M235 216h106q10 0 16-6t6-15q0-22-22-22H235q-97 0-166 68.5T0 408q0 21 21 21h470q10 0 17-8q4-4 4-19L427 94q-11-40-45.5-65.5T305 3H192q-21 0-21 21t21 21h113q28 0 50.5 17.5T386 107l77 280H45q8-72 62.5-121.5T235 216"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Itunes(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M232 2Q137 2 69.5 69.5T2 232t67.5 162.5T232 462t162.5-67.5T462 232T394.5 69.5T232 2m0 414q-76 0-130-54T48 232t54-130t130-54t130 54t54 130t-54 130t-130 54m84-122q1 14-8.5 26T283 334q-14 1-25-8t-13-23q-1-15 9-26.5t24-13.5q8-1 18 2v-96l-93 17v121q1 15-8 26.5T171 347q-15 1-26-8t-13-23q-1-15 8.5-27t24.5-13q8-1 18 2V140l133-23z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Iwatch(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0Q139 0 90 49Q72 67 68 87q-29 5-48.5 28.5T0 169v174q0 31 18.5 54T66 425q4 19 21 38q46 49 160 49h9q96-1 161-53.5T501 337q11-41 11-81t-11-81q-19-70-85-122.5T256 0M83 384q-16 0-28-12.5T43 343V169q0-16 12-28.5T83 128h4q17 0 29 12t12 29v174q0 17-12 29t-29 12zm378-58q-8 29-29.5 59T363 441.5T256 469q-105 2-137-34q-5-5-8-13q26-8 43-30t17-49V169q0-28-16-49.5T113 90q2-6 9-11q32-36 134-36q86 0 137.5 46t67.5 97q8 28 8 70t-8 70m-333-91q0 17-12.5 29.5T85 277q-17 0-29.5-12.5T43 235q0-18 12.5-30.5T85 192q18 0 30.5 12.5T128 235"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Justice(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M262 247q16 16 30 0l90-89q13-15 0-30q-15-15-30 0l-90-90q16-14 0-29q-14-16-29 0l-90 87q-15 15 0 30q14 14 30 0l30 30l-60 59q-15-13-30 0L9 320q-15 15 0 30l27 34q16 16 30 0l105-105q6-6 6-14.5t-6-14.5l59-60l30 30q-12 12 2 27m-29-119l-30-30l30-30l89 90l-30 30l-30-30z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Keyboard(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469 64H43q-18 0-30.5 12.5T0 107v213q0 17 12.5 30T43 363h426q18 0 30.5-13t12.5-30V107q0-18-12.5-30.5T469 64M43 320V107h426v213zm85-64h256v43H128zm277-64h-42v43h85V128h-43zm0 64h43v43h-43zm-341 0h43v43H64zm0-64h85v43H64zm235 0h42v43h-42zm-64 0h42v43h-42zm-64 0h42v43h-42zm149-64h43v43h-43zm-64 0h43v43h-43zm-64 0h43v43h-43zm-64 0h43v43h-43zm-64 0h43v43H64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Kik(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 415V49q0-19 12.5-33T42 2t29 14t12 33v206l104-112q19-21 37-21q16 0 26 12.5t10 29.5q0 20-19 41l-66 64l79 120q10 16 10 31q0 18-11.5 30T225 462q-22 0-36-21l-73-114l-33 34v54q0 19-12 33t-29 14t-29.5-14T0 415m334-56q21 0 35.5-14.5T384 309q0-20-14.5-35T334 259q-20 0-35 15t-15 35q0 21 15 35.5t35 14.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Krop(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M148 281q1 7-4 13.5t-11 8.5q-6 0-6-3q-1-1 0-3t1-3q-2-7-92-76q-2 26-3 50.5t-1.5 35T31 316q-2 10-14 12q-10 3-13-8q-3-16 13-222l-1-4q-6-26 15-30q16-2 19 8q0 2-1 14.5T45 126t-5 55q28-47 99-137q10-13 16-14q6-2 8 6q2 7-5 12q-23 29-102 148q41 26 65 50t27 35m242-127q7 33-4 63t-30 34q-16 5-25-8q-1 81 2 95q1 2 3 3.5t2 2.5q2 7-7 9q-16 3-17-9q-7-20-1-152q-6 9-17 14q-5 65-40 73q-26 7-34-20q-10 22-25 25q-23 5-29-21q-4-24 3-67l-12-1q-7 55-18 58q-5 2-7-5v-4q1-3 2.5-8t2.5-12.5t2-18t1-21.5q-8-3-8-7q-2-6 3.5-14t14.5-9q10-4 12 6v11l14 1q1-4 6-11q4-2 5-2q4-1 9.5 1t6.5 6q1 5-5 14q-16 47-10 77q3 10 8 10q11-3 21-37q2-25 10-48q7-21 20-24q9-2 13 4q4-2 6-2q9-2 16 5t11 22q6-3 9.5-7.5T312 168l3-8q3-56 12-58t11 6q0 2-.5 8t-1.5 17.5t-1 22.5q13-32 23-33q10-2 19 6t13 25m-116 33q-1-5-3.5-7.5T265 178q-10 2-18 27t-3 46q1 12 11 10t17-27.5t2-46.5m96-31q-2-6-4-9t-5-2t-8.5 9.5t-11 28T334 226q3 14 15 10q14-3 20-34q4-28 1-46"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lab(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M69 512h342q26 0 44.5-17.5T475 450v-2q0-24-13-45L304 192V43h21q8 0 15-7t7-15t-7-14.5T325 0H155q-8 0-15 6.5T133 21t7 15t15 7h21v147L20 401Q7 422 7 444l-2 4q1 27 20 45.5T69 512m107-250l11-12q24-34 32-77V43h42v128q9 47 32 76l11 13l43 58H133zM54 427l58-79v15h267l47 64q2 2 6 19q0 9-6.5 16t-14.5 7H69q-8 0-14.5-7T48 446q0-13 6-19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Label(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M325 0H27Q17 0 11 6T5 21v465q0 9 6.5 15.5T27 508q7 0 15-5l121-106q4-4 15-4q7 0 15 4l122 106q4 5 15 5q8 0 14.5-6.5T351 486V21q-7-21-26-21m-21 439l-85-74q-15-15-43-15q-21 0-43 15l-85 74V43h256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LabelHogwarts(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m42 501l121-106q4-5 15-5q5 0 15 5l122 106q4 5 15 5q8 0 14.5-6.5T351 484V21q0-9-6-15t-15-6H27Q17 0 11 6T5 21v465q0 9 6.5 15.5T27 508q8 0 15-7M261 43h43v396l-43-38zm-128 0h86v322q-15-15-43-15q-21 0-43 15zm-85 0h43v358l-43 38z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Laptop(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469 64q0-27-18-45.5T405 0H107Q79 0 61 18.5T43 64v213H0v43q0 27 18.5 45.5T64 384h384q27 0 45.5-18.5T512 320v-43h-43zM85 64q0-21 22-21h298q22 0 22 21v213h-22V85q0-21-21-21H128q-21 0-21 21v192H85zM64 341q-21 0-21-21h170q0 21 22 21zm384 0H277q22 0 22-21h170q0 21-21 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Last(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M370 329q-63-1-91-66l-5-11l-44-101q-11-26-35.5-42.5T139 92q-41 0-70.5 29.5T39 192t29.5 70.5T139 292q61 0 88-53l18 41q-40 50-106 50q-57 0-97-40.5T2 192t40-97.5T139 54q89 0 127 85q1 2 7.5 17.5t18 42T311 244q7 16 12.5 25t17 16t27.5 7q25 1 41-12t16-33q0-9-2.5-15.5t-10-11.5t-14-7.5T378 205q-44-15-63-31.5T296 126q0-32 20-50.5T370 57q45 0 67 40l-29 15q-16-22-39-22q-16 0-26.5 10T332 125q0 6 1 11t5 9t6.5 6t10 5.5T365 161l13.5 4.5L393 170q37 12 53 29t16 49q0 35-26 58t-66 23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Leaf(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M176 512q21 0 21-21v-66q63-7 103-45q47-44 47-124q0-23-2-34q4-13-3-21l-2-3q-10-42-32-79t-49-61t-42.5-36T189 3l-2-1l-5-2h-12l-5 2q-3 1-15 8.5T117.5 35T77 74t-38.5 55T12 198l-2 3q-7 8-3 21q-2 11-2 34q0 80 47 124q40 38 103 45v66q0 21 21 21m21-452q36 28 60 59l-60 50zm0 164l84-66q6 11 17 49l-101 72zm0 107l107-77v2q0 63-34 94q-29 25-73 34zM155 60v109l-60-47q15-23 60-62m-84 98l84 66v55L54 207q7-25 17-49m11 192q-34-34-34-94v-2l107 77v51q-41-4-73-32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Left(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M181 384q7 0 15-4q17-17 2-30L60 192L196 36q15-13-2-30q-14-14-30 3L4 192l162 186q4 6 15 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftArrowCircle(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 512q106 0 181-75t75-181t-75-181T256 0T75 75T0 256t75 181t181 75m0-469q88 0 150.5 62.5T469 256t-62.5 150.5T256 469t-150.5-62.5T43 256t62.5-150.5T256 43M149 262q3 0 0 0q0 3 2 5q0 2 3 2l42 64q12 17 30 6q18-10 7-30l-22-30h130q22 0 22-21t-22-21H211l22-30q11-20-7-30q-17-13-30 2l-42 64q0 2-3 2v11q-2 2-2 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lego(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M88 448v21q0 18 12.5 30.5T131 512h170q18 0 30.5-12.5T344 469v-21q28 0 56.5-25.5T429 363V171q0-34-28.5-60T344 85h-21V64q0-27-18.5-45.5T259 0h-86q-27 0-45.5 18.5T109 64v21H88q-35 0-60 25.5T3 171v192q0 35 25 60t60 25m213 21H131v-21h170zM152 64q0-21 21-21h86q21 0 21 21v21H152zM45 171q0-18 13-30.5T88 128h256q12 0 27.5 13.5T387 171v192q0 16-15.5 29T344 405H88q-17 0-30-12.5T45 363zm128 53q0 14-9.5 23t-22.5 9t-22.5-9.5T109 224t9.5-22.5T141 192t22.5 9t9.5 23m150 0q0 13-9.5 22.5T291 256t-22.5-9t-9.5-23t9.5-23t22.5-9t22.5 9.5T323 224M216 363q44 0 75-13.5t40-24.5t11-18q8-19-11-28q-19-8-28 11q0 3-6 9t-28 13.5t-53 7.5t-53-7.5t-28-14.5l-6-8q-9-19-28-11q-19 9-11 28q2 7 11 18t40 24.5t75 13.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lightning(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M240 0L5 299l128 21l-42 192l256-320l-128-21zm28 222l-96 119l2-12l11-45l-45-7l-58-8l96-122l-2 19l-4 41l40 6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Link(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M420 390q26-26 26-61t-26-60l-45-45q-25-26-60-26t-61 26l-30-30q26-26 26-61t-26-60l-45-45Q154 2 119 2T58 28L28 58Q2 84 2 119t26 60l45 45q26 26 59 26q34 0 60-26l30 30q-26 26-26 61t26 60l45 45q26 26 59 26q34 0 60-26zM102 194l-44-45q-13-13-13-29.5T58 90l29-30q12-15 32-15q18 0 30 13l45 44q13 13 13 30t-13 30l-45-43q-14-14-30 0q-14 16 0 30l45 45q-13 13-31 13t-31-13m197 196l-45-44q-13-13-13-30t13-30l45 45q14 14 30 0q14-16 0-30l-45-45q12-13 30-13q17 0 29 13l45 45q13 13 13 29.5T388 361l-30 29q-11 13-28.5 13T299 390"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Linkedin(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M462 274v170h-98V285q0-67-50-67q-37 0-51 36q-3 7-3 24v166h-99q1-269 0-297h99v42l-1 1h1v-1q32-49 89-49q51 0 82 34t31 100M58 4Q33 4 17.5 18.5T2 55t15 37t39 15h1q25 0 40.5-15T113 55q-1-22-16-36.5T58 4M7 444h99V147H7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LinkedinAlt(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M434 5H41q-14 0-24 9.5T7 38v394q0 14 10 24t24 10h393q14 0 24-10t10-24V38q0-14-10-23.5T434 5M147 390H77V182h70zm-35-236q-18 0-28.5-10.5T73 118t11-25.5T112 82q18 0 28.5 10.5T151 118t-10.5 25.5T112 154m286 236h-70V279q0-47-35-47q-26 0-36 25q-2 5-2 17v116h-70q1-188 0-208h70v30q23-35 63-35q36 0 58 24t22 70zM255 213v-1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Liquor(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M359 226q49-23 49-66q0-31-22-53t-53-22q-16 0-34 7q2-2 2-7V21q0-21-21-21H152q-21 0-21 21v64q0 22 21 22q0 27-7.5 53T122 194q-63 31-94.5 83T7 386q8 53 47.5 87t97.5 39h141q53-7 88.5-37t45.5-78q7-41-10-88t-58-83m6-66q0 11-13 21.5T323 192v9q-2-1-5.5-3.5T312 194q-13-7-28-47q26-19 49-19q13 0 22.5 9.5T365 160M173 43h86v21h-86zm212 347q-7 34-32.5 55T289 469H154q-41-3-70.5-27.5T47 380q-7-39 14.5-80t77.5-67q31-15 43.5-52t12.5-74h42q0 21 2 32q0 1-1 3t-1 3h2q15 67 52 88q53 28 77 72.5t17 84.5M216 256q-54 0-91 28t-37 68t37 68t91 28t91-28t37-68t-37-68t-91-28m0 149q-35 0-60-15.5T131 352t25-37.5t60-15.5t60 15.5t25 37.5t-25 37.5t-60 15.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Livejournal(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M365 452h-1l-156-65q-2-2-2-3q-2-2 1-4q27-26 27-65q0-19-5-31v-1L115 130q-45 39-42 78q0 4-4 4q-4 2-4-3q-4-47 51-90q52-41 90-28q3 0 3 5q-2 3-6 2q-34-10-82 27l117 156q22 14 49.5 13t49.5-16q2-2 4 0q2 0 2 3l26 167q0 3-1 3q-1 1-3 1m-148-70l143 60l-24-154q-48 28-96 3q3 12 3 24q0 36-26 67M38 169q-4 0-4-4q-4-47 51-90q52-41 90-28q3 1 3 5q-2 3-6 3q-34-12-82 26q-51 41-48 83q0 4-4 4zm274 257q10-34 44-40l9 62zM130 20q-43 0-84.5 45T25 147l175 232l165 69l-28-181L162 36q-10-16-32-16m0-16q30 0 46 22l174 231q3 3 3 7l28 182q1 9-6 15q-5 3-10 3q-4 0-6-1l-166-69q-5-2-6-5L12 157q0-1-.5-1t-.5-1q-20-34 5-77q18-31 51.5-52.5T130 4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lovedsgn(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M334 2q-44-1-69 12q-32 16-56 72q-1-3-4-7q-11-20-17-29t-16.5-20T148 13Q118-2 75 4Q17 18 8 82q-11 41 23 92q24 37 28 43q7 10 35.5 47t38.5 52q49 77 61 107q1 2 7.5 19.5T209 462q20-88 118-204q114-131 73-205q-17-38-66-51"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mac(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M43 384h170q-8 32-26.5 58.5T149 469q-9 0-15 6t-6 16q0 9 6 15t15 6h214q9 0 15-6t6-15q0-10-6-16t-15-6q-19 0-37.5-26.5T299 384h170q18 0 30.5-12.5T512 341V43q0-18-12.5-30.5T469 0H43Q25 0 12.5 12.5T0 43v298q0 18 12.5 30.5T43 384m170 85q36-49 43-85q15 60 43 85zM43 43h426v213H43zm0 256h426v42H43zm234-86q0 9-6 15.5t-15 6.5t-15-6.5t-6-15.5q0-8 6-14.5t15-6.5t15 6.5t6 14.5m-85 0q0 9-6.5 15.5T171 235q-9 0-15.5-6.5T149 213q0-8 6.5-14.5T171 192q8 0 14.5 6.5T192 213m171 0q0 9-6.5 15.5T341 235q-8 0-14.5-6.5T320 213q0-8 6.5-14.5T341 192q9 0 15.5 6.5T363 213"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MachineWash(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M66 311q4 22 21.5 37t39.5 15h274q22 0 39.5-15t21.5-37l58-264q2-9-3-17t-14-9q-9-2-17 3t-9 14l-17 84q-1-2-4.5-4.5T450 113q-24-28-58-28t-58 28q-14 15-27 15t-28-15q-24-28-58-28q-35 0-57 28q-15 15-28 15t-28-15Q88 91 61 85L51 38q-2-8-10-13t-16-4t-13.5 9.5T8 47zm12-168q24 28 58 28t58-28q14-15 27-15t28 15q28 28 58 28q35 0 57-28q15-15 28-15t28 15q17 17 32 23l-30 137q-3 17-21 17H127q-18 0-21-17L72 137z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MachineWashGentleOrDelicate(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M66 295q4 22 21.5 37t39.5 15h274q22 0 39.5-15t21.5-37l58-264q2-9-3-17t-14-9q-9-2-17 3t-9 14l-17 84q-1-2-4.5-4.5T450 97q-24-28-58-28t-58 28q-14 15-27 15t-28-15q-24-28-58-28q-35 0-57 28q-15 15-28 15t-28-15Q88 75 61 69L51 22q-2-8-10-13T25 5t-13.5 9.5T8 31zm12-168q24 28 58 28t58-28q14-15 27-15t28 15q24 28 58 28q35 0 57-28q15-15 28-15t28 15q17 17 32 23l-30 137q-3 17-21 17H127q-18 0-21-17L72 121zm37 284h298q22 0 22-22q0-9-6-15t-16-6H115q-10 0-16 6t-6 15q0 22 22 22m0 64h298q22 0 22-22q0-9-6-15t-16-6H115q-10 0-16 6t-6 15q0 22 22 22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MachineWashPermanentPress(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M66 298q4 22 21.5 36.5T127 349h274q22 0 39.5-14.5T462 298l58-264q2-9-3-17t-14-9q-9-2-17 3t-9 14l-17 83q-1-1-4.5-3.5T450 100q-24-28-58-28t-58 28q-15 15-27 15q-11 0-28-15q-24-28-58-28q-35 0-57 28q-16 15-28 15t-28-15Q86 76 61 72L51 25q-1-8-9.5-13.5T25 8q-8 1-13.5 9.5T8 34zm12-168q25 27 58 27t58-27q14-15 27-15t28 15q25 27 58 27q35 0 57-27q15-15 28-15t28 15q17 17 32 23l-30 137q-3 17-21 17H127q-18 0-21-17L72 123q4 4 6 7m335 283q22 0 22-21t-22-21H115q-22 0-22 21t22 21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MagnifyingGlass(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M414 354q-18-18-41-11l-32-32q43-53 43-119q0-80-56-136T192 0T56 56T0 192t56 136t136 56q70 0 119-43l32 32q-6 24 11 41l85 85q13 13 30 13q18 0 30-13q13-13 13-30t-13-30zm-222-13q-62 0-105.5-43.5T43 192T86.5 86.5T192 43t105.5 43.5T341 192t-43.5 105.5T192 341"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mail(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469 21H43q-18 0-30.5 13T0 64v256q0 17 12.5 30T43 363h426q18 0 30.5-13t12.5-30V64q0-17-12.5-30T469 21m-40 43L256 166L83 64zM43 320V90l202 121q2 2 11 2t11-2L469 90v230z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailBack(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469 0H43Q25 0 12.5 12.5T0 43v298q0 18 12.5 30.5T43 384h426q18 0 30.5-12.5T512 341V43q0-18-12.5-30.5T469 0m-34 43L256 186L77 43zM43 341V70l121 99l-51 51q-14 14 0 30q6 6 15 6t15-6l55-56l58 47l58-47l55 56q6 6 15 6t15-6q14-16 0-30l-51-51l121-99v271z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailBill(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 0H64Q37 0 18.5 18.5T0 64v256q0 27 18.5 45.5T64 384h384q27 0 45.5-18.5T512 320V64q0-27-18.5-45.5T448 0m21 320q0 21-21 21H64q-21 0-21-21V64q0-21 21-21h21v85l43-21l43 21V43h277q21 0 21 21zM405 64h-42q-10 0-16 6t-6 15v64q0 22 22 22h42q22 0 22-22V85q0-9-6-15t-16-6m-21 149H277q-17 0-29.5 13T235 256v21q0 18 12.5 30.5T277 320h107q17 0 30-12.5t13-30.5v-21q0-17-13-30t-30-13m0 64H277v-21h107z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MailStamp(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469 0H43Q25 0 12.5 12.5T0 43v298q0 18 12.5 30.5T43 384h426q18 0 30.5-12.5T512 341V43q0-18-12.5-30.5T469 0M43 341V43h426v298zm341-149H256q-21 0-21 21q0 22 21 22h128q21 0 21-22q0-21-21-21m-64 64h-64q-21 0-21 21q0 22 21 22h64q21 0 21-22q0-21-21-21M427 64h-43q-21 0-21 21v64q0 22 21 22h43q9 0 15-6t6-16V85q0-9-6-15t-15-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mailbox(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M367 0H145q-14 0-25.5 9T105 32L0 299v170q0 18 12.5 30.5T43 512h426q18 0 30.5-12.5T512 469V299L407 32q-3-14-14.5-23T367 0m102 469H43V341h85l43 64h170l43-64h85zm-81-170q-29 0-38 23l-30 41H192l-30-41q-9-23-38-23H51l98-256h214l98 256z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Man(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M192 64q0 26-19 45t-45 19t-45-19t-19-45t19-45t45-19t45 19t19 45M64 256v235q0 9 6 15t15 6q10 0 16-6t6-15V363q0-22 19-22h2q6 0 11 2v-2h10v150q0 9 6 15t16 6q9 0 15-6t6-15V192q0-17-12.5-30T149 149h-42q-18 0-30.5 13T64 192zm77 87l8 20q0-10-8-20m-98-44V149q-18 0-30.5 13T0 192v149q17 0 30-12.5T43 299m213-107q0-17-12.5-30T213 149v150q0 17 13 29.5t30 12.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MaximumTempOneHundredFiftyThreeHundred(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M363 195q0-22-22-22H235q-97 0-166 68.5T0 408q0 21 21 21h470q10 0 17-8q4-4 4-19L427 94q-11-40-45.5-65.5T305 3H192q-21 0-21 21t21 21h113q28 0 50.5 17.5T386 107l77 280H45q8-72 62.5-121.5T235 216h106q10 0 16-6t6-15M235 301q0 9-6.5 15.5T213 323q-8 0-14.5-6.5T192 301q0-8 6.5-14.5T213 280q9 0 15.5 6.5T235 301m85 0q0 9-6.5 15.5T299 323q-9 0-15.5-6.5T277 301q0-8 6.5-14.5T299 280q8 0 14.5 6.5T320 301"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MaximumTempOneHundredTenTwoHundredThirty(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M363 195q0-22-22-22H235q-97 0-166 68.5T0 408q0 21 21 21h470q10 0 17-8q4-4 4-19L427 94q-11-40-45.5-65.5T305 3H192q-21 0-21 21t21 21h113q28 0 50.5 17.5T386 107l77 280H45q8-72 62.5-121.5T235 216h106q10 0 16-6t6-15m-86 106q0 9-6 15.5t-15 6.5t-15-6.5t-6-15.5q0-8 6-14.5t15-6.5t15 6.5t6 14.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MaximumTempTwoHundredThreeHundredNinety(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 429h470q10 0 17-8q4-4 4-19L427 94q-11-40-45.5-65.5T305 3H192q-21 0-21 21t21 21h113q28 0 50.5 17.5T386 107l77 280H45q8-72 62.5-121.5T235 216h106q10 0 16-6t6-15q0-22-22-22H235q-97 0-166 68.5T0 408q0 21 21 21m171-128q0 9-6.5 15.5T171 323q-9 0-15.5-6.5T149 301q0-8 6.5-14.5T171 280q8 0 14.5 6.5T192 301m85 0q0 9-6 15.5t-15 6.5t-15-6.5t-6-15.5q0-8 6-14.5t15-6.5t15 6.5t6 14.5m86 0q0 9-6.5 15.5T341 323q-8 0-14.5-6.5T320 301q0-8 6.5-14.5T341 280q9 0 15.5 6.5T363 301"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MayoHotdog(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M133 196q0 24 32 43q1 1 4.5 2t4.5 2q-4 4-9 4q-32 19-32 43q0 21 32 43q4 4 9 4q-4 4-9 4q-32 22-32 43t22 21q8 0 14.5-5.5T176 386q3-3 11-6q10-6 14.5-9t11-12.5T219 337q0-7-2-13.5t-6.5-11.5t-7.5-7t-9-7l-7-4q-4-4-9-4q4-4 9-4q10-6 14.5-9t11-12.5T219 243q0-23-32-42q-1-2-4.5-3t-4.5-2q4-4 9-4q32-19 32-43q0-9-6-15t-16-6q-8 0-14.5 6t-6.5 13q-1 1-11 7q-32 19-32 42M176 0q-21 0-21 21v2q-40 8-64 41q-35 0-60.5 25T5 149v214q0 35 25.5 60T91 448q24 33 64 41v2q0 21 21 21t21-21v-2q40-8 64-41q35 0 60.5-25t25.5-60V149q0-35-25.5-60T261 64q-24-33-64-41v-2q0-21-21-21M69 399q-21-11-21-36V149q0-25 21-36zm214-286q21 11 21 36v214q0 25-21 36zm-43 15v256q0 27-18.5 45.5T176 448t-45.5-18.5T112 384V128q0-27 18.5-45.5T176 64t45.5 18.5T240 128"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Meetup(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M212 199q-3 0-4-2q0-7 4-15l3-6v-3l1 3q3 10-1 23zm193-39l-1 1q-14 13-22 52q18-5 25-24q5-16-2-28zm-241 34h1q1 1 2 1l2-1q1-5-1-18q-7 8-8 17v1zm298 120q0 16-16 16H18q-16 0-16-16V70q0-16 16-16h428q16 0 16 16zm-322-90q-6-15-10-41q-9-38-14-50l-1 1h-8q-1 0-3 1t-3 1t-1-1.5t-1-1.5q-10 14-18 52q-4-7-7-14q-2-3-6.5-11.5T60 147h-4q-8 0-15 3v12q-2 39 5 59h3q3 0 15-4q0-3-.5-8t-.5-7q-2-3 0-8q2 2 5.5 8t5.5 8q2 4 6 10h10q2-1 7-1h2q3-6 3-26q0-16 4-25q7 26 18 59q3 0 9-2q2-1 5-1zm53 4q-2-4-7-4q-3 0-7 2q-3 1-8 1h-2q-5-2-7-5q-4-7-1-21q5 1 16 1h5l1 1q4-16-1-38q-1 0-2.5-.5T177 163q-6-2-10-2h-2q-17 3-21 34q-4 29 15 36q3 2 10 2q14 0 24-5m47-3q-1-3-6-5q-5 7-12 7q-9 0-14-7q-5-6-1-15q4 1 12 1h9q-1-9 1-23q2-12 3-20q-9-6-18-6q-23 19-22 56q1 8 10.5 13t19.5 5q12 0 18-6m54-35q-4-4-5-6q-6-6-11-8q-1 0-1 3l-3 5h-2v-1q0-17-1-27q-2-28 1-40q-4-6-17-6h-1q-1-2-1-3q-5 15-1 49q4 18 1 29q-3 0-5-1h-5q-3 0-5 2q-1 1-1 6q3 2 16 2l4 1q0 11 1 16q0 4 .5 11.5t.5 12.5h4q3-1 9-1q0-3-.5-9t-.5-8v-23q2-1 5-1.5t5-.5q11 0 13-2m57-10q-11-8-14-8h-2l-1 1v5q0 38-15 43q-1-6-2-17t-1-12v-12l-2-1q-2-1-2-2q-8-4-13-5q-3 40 8 53q5 5 12 5q3 0 11-2l2-2q3-6 8-10l3-6q10-16 8-30m72-25q-11-7-22-7q-13 0-21 11q-10 12-17 58q0 1-2 1q-1 0-2.5.5T356 220q7 8 9 32q2 19 5 25l2 1h15q-7-14-5-47q30-8 41-33q10-20 0-43"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Megaphone(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469 127V44q0-22-17-34q-15-14-38-7L30 118q-13 5-21.5 16.5T0 159v75q0 14 9 25.5T32 274l32 9v74q0 28 18.5 46t45.5 18h41q32 0 53-27l47-60l149 36q2 0 5.5 1t5.5 1q17 0 25-8q17-14 17-34v-84q19-6 31-22.5t12-36.5q0-21-12.5-37.5T469 127M43 159l42-13v98l-42-10zm145 211q-11 9-19 9h-41q-21 0-21-22v-64l115 30zm239-34l-318-85h19V133l299-87z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Metacafe(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m421 462l-68-49l-68 49l26-79l-67-49h83l26-79l26 79h83l-67 49zm-129-73l-99-71h122l38-116q32 99 38 116h65L353 2l-83 254H2l269 196q4-12 8.5-25t8-24.5T292 389"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mic(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M176 341q45 0 76-31t31-75V107q0-45-31-76T176 0t-76 31t-31 76v128q0 44 31 75t76 31m-64-234q0-28 18.5-46T176 43t45.5 18t18.5 46v128q0 27-18.5 45.5T176 299t-45.5-18.5T112 235zm235 128v-64q0-22-22-22q-9 0-15 6t-6 16v64q0 53-38 90.5T176 363t-90-37.5T48 235v-64q0-10-6-16t-15-6q-22 0-22 22v64q0 65 43 112.5T155 403v45q0 7-5.5 12t-11.5 7l-5 2q-17 0-29.5 13T91 512h170q0-17-12.5-30T219 469q-22-6-22-21v-45q64-8 107-55.5T347 235"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MicOff(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M133 469q-17 0-29.5 13T91 512h170q0-17-12.5-30T219 469q-22-6-22-21v-45q30-4 49-15l-27-32q-18 7-43 7q-52 0-90-37.5T48 235v-64q0-7-2-11l-7-6q-7-5-12-5q-8 0-15 7t-7 15v64q0 65 43 112.5T155 403v45q0 7-5.5 12t-11.5 7zm43-128q17 0 26-4l-35-38q-24-4-39.5-21.5T112 237l-43-49v47q0 44 31 75t76 31m107-104V107q0-45-31-76T176 0q-43 0-75 30l28 32q22-19 47-19q27 0 45.5 18t18.5 46v81zm42-88q-8 0-14.5 7t-6.5 15v64q0 15-2 23l34 39q11-37 11-62v-64q0-8-7-15t-15-7m5 271q6 7 17 7q7 0 15-5q16-16 2-29l-49-56l-30-32l-15-17l-32-36l-126-145l-36-41l-32-38q-10-7-17-7q-8 0-15 5q-17 15-2 29l59 69l43 47l102 115l28 32l15 17l28 32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Milkshake(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M304 149H197V98l96-57q17-11 9-30q-3-8-12.5-10T272 2L155 73v76H48q-17 0-30 13T5 192v64h26l38 218q1 16 14 27t29 11h132q16 0 28-10t15-26l34-220h26v-64q0-17-13-30t-30-13M91 213v-21h42v21zm64-21h42v21h-42zm64 0h42v21h-42zm85 21h-21v-21h21zM48 192h21v21H48zm196 277l-132-2l-38-211h204z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ming(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M398 313q0-13 9-22.5t23-9.5t23 9.5t9 22.5t-9.5 22.5T430 345t-22.5-9.5T398 313m32-133q14 0 23-9t9-23q0-13-9.5-22t-22.5-9t-22.5 9t-9.5 22q0 14 9 23t23 9m-122 7q0 13 9 22.5t23 9.5t23-9.5t9-22.5t-9-22.5t-23-9.5t-23 9.5t-9 22.5M148 398q-13 0-22.5 9.5T116 430q0 14 9 23t23 9t23-9t9-23q0-13-9.5-22.5T148 398m128-243q13 0 22.5-9.5T308 123t-9.5-22.5T276 91t-23 9.5t-10 22.5t9.5 22.5T276 155m64 88q-14 0-23 9.5t-9 22.5t9 22.5t23 9.5t23-9.5t9-22.5t-9-22.5t-23-9.5m-152 64q-13 0-22.5 9.5T156 339t9.5 22.5T188 371q14 0 23.5-9.5T221 339t-9.5-22.5T188 307m128 91q-13 0-23 9.5T283 430t9.5 22.5T316 462q13 0 22.5-9.5T348 430t-9.5-22.5T316 398m0-332q13 0 22.5-9.5T348 34t-9.5-22.5T316 2q-14 0-23.5 9.5T283 34t10 22.5t23 9.5M124 307q14 0 23-9.5t9-22.5t-9-22.5t-23-9.5t-23 9.5t-9 22.5t9 22.5t23 9.5M34 180q14 0 23-9t9-23q0-13-9.5-22T34 117t-22.5 9T2 148q0 14 9 23t23 9m0 165q13 0 22.5-9.5T66 313t-9-22.5t-23-9.5t-23 9.5T2 313t9.5 22.5T34 345m90-126q14 0 23-9.5t9-22.5t-9-22.5t-23-9.5t-23 9.5t-9 22.5t9 22.5t23 9.5m152 152q13 0 22.5-9.5T308 339t-9.5-22.5T276 307q-14 0-23.5 9.5T243 339t9.5 22.5T276 371M156 123q0 13 9.5 22.5T188 155q14 0 23.5-9.5T221 123t-10-22.5t-23-9.5t-22.5 9.5T156 123m-8-57q13 0 22.5-9.5T180 34q0-14-9-23t-23-9t-23 9t-9 23q0 13 9.5 22.5T148 66"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Minus(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M363 213q21 0 21-21t-21-21H21q-21 0-21 21t21 21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusBox(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M405 0H107Q62 0 31 31T0 107v298q0 45 31 76t76 31h298q45 0 76-31t31-76V107q0-45-31-76T405 0m64 405q0 28-18 46t-46 18H107q-28 0-46-18t-18-46V107q0-28 18-46t46-18h298q28 0 46 18t18 46zM363 235H149q-21 0-21 21t21 21h214q21 0 21-21t-21-21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusCircle(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0Q150 0 75 75T0 256t75 181t181 75t181-75t75-181t-75-181T256 0m0 469q-88 0-150.5-62.5T43 256t62.5-150.5T256 43t150.5 62.5T469 256t-62.5 150.5T256 469m107-234H149q-21 0-21 21t21 21h214q21 0 21-21t-21-21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusCircleOne(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0Q150 0 75 75T0 256t75 181t181 75t181-75t75-181t-75-181T256 0m0 469q-88 0-150.5-62.5T43 256t62.5-150.5T256 43t150.5 62.5T469 256t-62.5 150.5T256 469m107-234H149q-21 0-21 21t21 21h214q21 0 21-21t-21-21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MisterWong(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M117 223q-18 19-56.5 52T17 313q-9-4-14-14q52-44 100-88q5 3 14 12m182-5q-1 2-5.5 5.5T288 230q76 65 106 87q11-9 12-12q-16-12-51-42t-56-45M99 77H3v17q17 1 48 1h49q15 115 101 135v113Q43 356 14 463q6 1 59 1q28-78 128-89v88q1 1 20 1q0-82 1-92q117 6 148 92h37q-27-105-186-122V229q77-20 94-134h92V77h-46.5L314 76q0-41-14-72h-32q16 47 12 73q-13-1-67.5 0T140 75q0-23 13-71h-39q-10 18-15 73m45 18h136q-9 92-57 105q-14 3-26 0q-43-12-56-102q0-3 3-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mixx(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M168 327q-16 0-34-9q24 6 51-4q32-13 43-38q0 21-17.5 36T168 327m9 30q32 0 56.5-17.5T266 295q-24 44-79 52q-48 6-81-22q26 32 71 32m115-105v9q0 47-33.5 81T178 376q-33 0-61-18.5T75 309q11 36 40.5 58.5T183 390q47 0 79.5-33t32.5-79q0-14-3-26m66 37q0 73-51.5 125T182 466q-74 0-125.5-52T5 289q0-53 29-97t76-65L78 32q-3-8 1-15.5T91 7V6h6q15 0 20 15l22 97q21-6 43-6q73 0 124.5 52T358 289m-34 0q0-59-41.5-101T182 146q-17 0-35 5l27 112q4 22-2 23q0 1-1 1q-7 0-12-18l-38-109q-37 18-59.5 52.5T39 289q0 59 42 101t101 42t100.5-42T324 289"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MixxAlt(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M460 131V20H0v348h111V132h63v236h112V132h63v236h111z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mobileme(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M462 257q0 38-26.5 64.5T372 348H92q-37 0-63.5-26.5T2 257q0-34 22-59.5T78 167q-1-3-1-10q0-24 17-41t40-17q20 0 37 14q17-37 39.5-57T276 36q55 0 84.5 36t29.5 88v8q31 6 51.5 31t20.5 58"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Moon(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M299 235q-63 0-106.5-43.5T149 85q0-40 26-85q-74 7-124.5 62T0 192q0 80 56 136t136 56q75 0 130-50.5T384 209q-45 26-85 26M192 341q-62 0-105.5-43.5T43 192q0-79 64-122v15q0 80 56 136t136 56h15q-43 64-122 64"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mouse(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M176 0Q106 0 55.5 50T5 171v170q0 71 50.5 121T176 512t120.5-50T347 341V171q0-71-50.5-121T176 0m128 341q0 53-37.5 90.5T176 469t-90.5-37.5T48 341v-85h256zm0-128H48v-42q0-53 37.5-90.5T176 43t90.5 37.5T304 171zM176 85q-21 0-21 22v64q0 21 21 21t21-21v-64q0-22-21-22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MsnMessenger(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 251q3-32 21-56.5T54 162t23-11q-24-19-24-49q0-26 19-45t46-19q26 0 45 19t19 45q0 29-23 49q20 7 34 19q-89 45-99 142v2q-24-1-39-5.5T37 300l-3-4q-5-5-8-13l-13-6q-1-1-3-2t-5-8t-3-16m227-79q-10 4-22.5 11T173 207t-36.5 46.5T117 320q0 14 4 23.5t9 12.5l4 2q7 7 19 10q3 10 12 18q1 3 4.5 6.5T197 404t63 10h58q5 0 7-1q40-3 64.5-11t28.5-15l5-7v-1q4-4 6-11q28-5 32.5-36t-7.5-58q-12-37-40.5-63.5T350 171q35-28 35-72q0-39-28-66.5T289 5t-68 27.5T193 99q0 45 36 73"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Music(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M486 13Q463-5 431 2L218 64q-21 4-34 21.5T171 126v235q-26-20-64-20q-45 0-76 25.5T0 427t31 60t76 25q44 0 75-27t31-63v-1q-1-1-2-16h2V213q0-13 15-21l214-62q8-3 19 4q8 7 8 17v143q-29-17-64-17q-44 0-75 25.5T299 363t31 60t75 25q45 0 76-25t31-60q0-8-4-22h4V64q0-31-26-51M107 469q-26 0-45-12.5T43 427q0-18 19-30.5t45-12.5t45 12.5t19 30.5q0 17-19 29.5T107 469m298-64q-26 0-45-12.5T341 363q0-18 19-30.5t45-12.5t45 12.5t19 30.5q0 17-19 29.5T405 405m26-313l-213 62q-3 0-5 2v-30q0-14 15-21l214-62q15 0 19 4q8 5 8 17v28q-18-8-38 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicScore(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M491 0h-39q-21 7-34 24t-13 38v296q-29-17-64-17q-44 0-75 25.5T235 427t31 60t75 25q45 0 76-25t31-60q0-9-2-13q0-1 1-3.5t1-5.5V62q0-16 13-19h30q9 0 15-6t6-16q0-9-6-15t-15-6M341 469q-26 0-45-12.5T277 427q0-18 19-30.5t45-12.5t45 12.5t19 30.5q0 17-19 29.5T341 469M21 107h320q22 0 22-22q0-9-6-15t-16-6H21q-9 0-15 6T0 85q0 10 6 16t15 6m0 85h320q10 0 16-6t6-15q0-22-22-22H21q-9 0-15 6t-6 16q0 9 6 15t15 6m0 85h320q22 0 22-21t-22-21H21q-21 0-21 21t21 21m0 86h192l43-43H21q-9 0-15 6t-6 15q0 10 6 16t15 6m0 85h171v-43H21q-9 0-15 6t-6 16q0 9 6 15t15 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Myspace(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M462 206v60H2V118h59v88h342v-88h59z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MyspaceAlt(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M321 157q-31 0-54-22.5T244 80t22.5-54.5T321 3t54.5 22.5T398 80t-22.5 54.5T321 157m-63 74q0-32-22.5-54.5T181 154t-54 22.5t-22 54.5v138h153zm-77-88q24 0 40.5-16.5T238 87t-16.5-39.5T181 31q-23 0-39.5 16.5T125 87t16.5 39.5T181 143m245 132q0-43-31-74t-74-31t-74 31t-31 74q0 2 .5 4t.5 3h-1v182h210zM70 145q-26 0-44.5 18.5T7 208v113h126V212h-1v-4q0-26-18-44.5T70 145m0-8q19 0 32.5-13.5T116 91t-13.5-32.5T70 45q-20 0-33.5 13.5T23 91t14 32.5T70 137"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Newsvine(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m144 299l55-40q8 5 17 5q14 0 23.5-9.5T249 232q0-14-9.5-23.5T216 199q-13 0-22.5 9.5T184 232v4l-40 29v-91l55-40q7 4 17 4q14 0 23.5-9.5T249 106q0-14-9.5-23.5T216 73q-13 0-22.5 9.5T184 106v4l-40 29V2h-32v72L72 45v-6q0-14-9.5-23.5T40 6q-14 0-23.5 9.5T7 39t9.5 23.5T40 72q7 0 15-4l57 41v91l-40-30v-5q0-14-9.5-23.5T40 132q-14 0-23.5 9.5T7 165t9.5 23t23.5 9q9 0 16-4l56 41v97l-40-29v-6q0-14-9-23t-23-9t-23.5 9T7 296t9.5 23.5T40 329q6 0 16-4l56 40v97h32v-32l55-39q8 5 17 5q14 0 23.5-10t9.5-23t-9.5-23t-23.5-10q-13 0-22.5 9.5T184 363v4l-40 29z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Next(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M197 365q22 0 22-21V88q0-21-22-21q-21 0-21 21v73L42 11Q34-2 18 5Q5 9 5 24v384q0 15 13 19q4 2 9 2q11 0 15-6l134-149v70q0 21 21 21M48 353V79l122 137z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NinetyFiveTwoHundred(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M127 384h274q22 0 39.5-14.5T462 333l58-286q2-9-3-17t-14-9q-9-2-17 3t-9 14l-12 67q-5-3-15-13q-24-28-58-28t-58 28q-15 15-27 15q-11 0-28-15q-24-28-58-28q-35 0-57 28q-16 15-28 15t-28-15Q84 66 55 64l-4-26q-2-8-10-13t-16-4t-13.5 9.5T8 47l58 286q4 22 21.5 36.5T127 384M78 122q25 27 58 27t58-27q14-15 27-15t28 15q25 27 58 27q35 0 57-27q15-15 28-15t28 15q19 19 36 25l-36 177q-3 17-22 17H127q-18 0-21-17L63 109q3 2 7.5 6.5T78 122m207 70q0 9-6 15t-15 6t-15-6t-6-15t6-15t15-6t15 6t6 15m107 0q0 9-6.5 15t-14.5 6q-9 0-15.5-6t-6.5-15t6.5-15t15.5-6q8 0 14.5 6t6.5 15m-213 0q0 9-6.5 15t-15.5 6q-8 0-14.5-6t-6.5-15t6.5-15t14.5-6q9 0 15.5 6t6.5 15m106 85q0 9-6 15.5t-15 6.5t-15-6.5t-6-15.5q0-8 6-14.5t15-6.5t15 6.5t6 14.5m107 0q0 9-6.5 15.5T371 299q-9 0-15.5-6.5T349 277q0-8 6.5-14.5T371 256q8 0 14.5 6.5T392 277m-213 0q0 9-6.5 15.5T157 299q-8 0-14.5-6.5T136 277q0-8 6.5-14.5T157 256q9 0 15.5 6.5T179 277"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NoEye(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m411 113l-21 36q55 32 83 75q-13 20-33 39.5t-64 42t-95 25.5l-26 42h9q162 0 254-138q6-11 0-22q-3-6-12.5-19.5t-36-38.5t-58.5-42m-85 51l19-34l21-38l43-71l-36-21l-49 81q-23-6-54-6h-15Q100 79 10 213q-6 11 0 22q2 3 6.5 10.5t18 23T64 299t41.5 31t53.5 26l-42 71l36 21l49-81l24-38l25-41l26-43zm-94 73l-21 36l-28 45q-81-28-128-94q50-64 113-87q43-20 96-20q26 0 38 2l-4 7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NonChlorineBleachIfNeeded(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M473 373q7-11 0-21L259 11q-6-10-18-10t-18 10L10 352q-9 10 0 21q6 11 17 11h426q12 0 20-11M345 230l-73 111h-79l111-175zM240 62l41 64l-137 215H65zm83 279l47-72l45 72z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Official(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M347 307q-11-8-19.5-43T310 180t-18-75q-11-30-36-47t-50.5-23.5T178 26L159 5l-2-2q-2-3-10 0t-8 5v290q0 16-12 22.5T89 335t-47 23q-28 21-37 45.5T7 443q9 17 38 19t70-25q43-33 43-95V19q1 2 7 11q12 20 12.5 69.5T195 191q26 62 113 105q40 20 42 15q0-2-3-4m-34-29q-44-53-60-97q-33-85-50-108q-18-24-23-33q0-2 1-2q11 14 24 30q23 30 51 102q16 43 30 64q14 24 29 45z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OpenPadlock(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M192 43q35 0 60 25t25 60q0 21 22 21q21 0 21-21q0-52-38-90T192 0t-90 38t-38 90v64q-27 0-45.5 18.5T0 256v192q0 27 18.5 45.5T64 512h256q27 0 45.5-18.5T384 448V256q0-27-18.5-45.5T320 192H107v-64q0-35 25-60t60-25m128 192q21 0 21 21v192q0 21-21 21H64q-21 0-21-21V256q0-21 21-21zM171 380v25q0 22 21 22t21-22v-25q43-15 43-60q0-27-18.5-45.5T192 256t-45.5 18.5T128 320q0 20 12 36.5t31 23.5m21-81q21 0 21 21t-21 21t-21-21t21-21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Openid(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M417 171q-62-36-134-46V6l-70 45v73q-45 4-86 19t-68 37.5t-44 50T2 287t24 55t68.5 48.5T213 426l70-45V173q39 6 90 28l-40 25h129v-85zm-204-2v212q-60-8-96.5-31T75 298.5t7.5-56.5t48-48.5T213 169"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Organisation(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469 256v-21q0-28-18-46t-46-18H277v-43q18 0 30.5-12.5T320 85V43q0-18-12.5-30.5T277 0h-42q-18 0-30.5 12.5T192 43v42q0 18 12.5 30.5T235 128v43H107q-28 0-46 18t-18 46v21q-18 0-30.5 12.5T0 299v42q0 18 12.5 30.5T43 384h42q18 0 30.5-12.5T128 341v-42q0-18-12.5-30.5T85 256v-21q0-22 22-22h128v43q-18 0-30.5 12.5T192 299v42q0 18 12.5 30.5T235 384h42q18 0 30.5-12.5T320 341v-42q0-18-12.5-30.5T277 256v-43h128q22 0 22 22v21q-18 0-30.5 12.5T384 299v42q0 18 12.5 30.5T427 384h42q18 0 30.5-12.5T512 341v-42q0-18-12.5-30.5T469 256M235 43h42v42h-42zM85 341H43v-42h42zm192 0h-42v-42h42zm150 0v-42h42v42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Orkut(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M232 2Q137 2 69.5 69.5T2 232t67.5 162.5T232 462t162.5-67.5T462 232T394.5 69.5T232 2m0 387q-65 0-111-46T75 232t46-111t111-46t111 46t46 111t-46 111t-111 46"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Padlock(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M192 0q-52 0-90 38t-38 90v64q-27 0-45.5 18.5T0 256v192q0 27 18.5 45.5T64 512h256q27 0 45.5-18.5T384 448V256q0-27-18.5-45.5T320 192v-64q0-52-38-90T192 0m149 256v192q0 21-21 21H64q-21 0-21-21V256q0-21 21-21h256q21 0 21 21m-234-64v-64q0-35 25-60t60-25t60 25t25 60v64zm85 64q-27 0-45.5 18.5T128 320q0 20 12 36.5t31 23.5v25q0 22 21 22t21-22v-25q43-15 43-60q0-27-18.5-45.5T192 256m0 85q-21 0-21-21t21-21t21 21t-21 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pandora(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M321 40Q278 3 202 3H5v16q47 0 59 15t12 74v250q0 59-12 74.5T5 448v16h202v-16q-47 0-59-15.5T136 358v-96h66q76 0 119-36q42-38 42-93t-42-93M202 234h-66V32h66q45 0 70 27.5t25 73.5t-25 73.5t-70 27.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pant(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M45 512h26q34 0 43-36l38-224l41 226q3 16 14.5 26t27.5 10h24q17 0 29.5-12.5T301 471V43q0-18-12.5-30.5T259 0H45Q28 0 15.5 12.5T3 43v426q0 18 12.5 30.5T45 512m188-43l-55-320q-5-21-26-21t-26 21L71 469H45V149q41 0 62.5-20T131 87v-2h64v2q2 19 18 37t46 23v322zm26-364q-18-6-22-20h22zM45 43h214v21H45zm0 42h43q-6 22-43 22z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaperTablet(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M341 64h-21q-10-21-43-21h-64V21q0-21-21-21t-21 21v22h-64q-33 0-43 21H43q-18 0-30.5 13T0 109v360q0 18 12.5 30.5T43 512h298q18 0 30.5-13t12.5-32V109q0-19-12.5-32T341 64M117 85h150q10 0 10 11t-10 11H117q-10 0-10-11t10-11m224 384H43V107h21q1 5 4 12t14.5 18.5T107 149h170q36 0 43-42h21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Path(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M116 23Q22 61 5 153q-11 62 12 113q6 11 14 23q2 3 8 9q10-4 15-7q9-5 17.5-12T90 262.5t16-14.5q-39-51-10-102t90-65q56-14 114 9.5t74 75.5t-22 85q-19 17-46 22q-14 3-31 1q-5 0-17-2q-2-1-5-1.5t-5-1t-3.5-2t-2.5-3.5t-1-5V140q0-6-.5-6.5t-5.5-.5q-12-2-16-2q-39-2-57 1q-8 5-8 7v63q0 18 1.5 64t.5 70q-4 46-16 52q-10 4-27.5-3.5T92 376q0 5-1.5 25.5t1 34T106 454q42 17 76 5q37-13 53-48.5t10-75.5q61 18 119-8.5t85-81.5q18-39 11.5-83.5T426 84q-34-40-89-61T224 2.5T116 23"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paypal(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M291 2H89L0 412h118l29-136h85q62 0 107.5-36t58.5-99q14-66-21-102.5T291 2m-88 195h-39l25-112h58q31 0 40 29q-1 0-3.5-.5t-3.5-.5h-58zm84-56q-5 23-24 39t-41 17l14-61h52q0 3-1 5m144 28q9-44-6-78q32 39 20 101q-14 63-59.5 98.5T278 326h-85l-29 136H46l5-22h99l29-136h85q62 0 107.5-36t59.5-99"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pc(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M261 0H69Q42 0 23.5 18.5T5 64v384q0 27 18.5 45.5T69 512h192q28 0 46-18.5t18-45.5V64q0-27-18-45.5T261 0m22 448q0 21-22 21H69q-21 0-21-21V213h235zm0-256H48V64q0-21 21-21h192q22 0 22 21zm-22 64q0 9-6 15t-15 6t-15-6t-6-15t6-15t15-6t15 6t6 15M69 85h192v43H69zm171 75q0 11-11 11q-10 0-10-11t10-11q11 0 11 11M91 427h64v21H91zm85 0h64v21h-64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pdiddy(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 171q-10-73-64.5-122T256 0T128.5 49.5T66 171h-2q-27 0-45.5 18T0 235q0 33 19 60t49 38q12 72 64.5 125.5T256 512t123.5-53.5T444 333q29-12 48.5-39t19.5-59q0-28-18.5-46T448 171M151 85q107 40 210 0q36 34 42 86H109q6-49 42-86M64 213v69q-21-18-21-47q0-22 21-22m192 256q-55 0-97-45.5T109 320h74q30 0 41-30l6-36q5-21 26-21t26 21l6 36q11 30 41 30h74q-8 58-50 103.5T256 469m192-187v-69q21 0 21 22q0 29-21 47m-171 81h-42q-22 0-22 21t22 21h42q22 0 22-21t-22-21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pen(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M344 24Q323 3 289 3t-52 21L3 259v170h170l235-234q20-20 20.5-53T408 88zm-145 98l40 41l-108 108v-29l-41-9zm-43 265H62l20-20l-30-29l-7 6v-68l9-9l34 9v51l49 15l17 34l26-13zm34-34l-23-47l-6-3l110-110l41 40zm188-188l-38 38L229 92l38-38q9-9 22-9q11 0 25 9l64 64q7 7 7 23q0 17-7 24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Penknife(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M176 207q89 47 164 47q73 0 126-45l26-21L165 45Q139 0 91 0Q56 0 30.5 25T5 85v342q0 35 25.5 60T91 512q30 0 54-19.5t29-48.5l70-19l-2-11l49 30l28-51l49 29l28-51l51 30l38-70l-38-22l-15 34l-51-29l-26 49l-51-30l-28 51l-30-19l-12 19l-58 15zm0-49V96l233 100q-93 45-233-38m-43 269q0 17-12.5 29.5T91 469q-18 0-30.5-12.5T48 427V85q0-17 12.5-29.5T91 43q17 0 29.5 12.5T133 85zm-21-320q0 8-6.5 14.5T91 128q-9 0-15.5-6.5T69 107q0-9 6.5-15.5T91 85q8 0 14.5 6.5T112 107"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PeopleTeam(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M427 215v-6q19-6 30.5-22.5T469 149v-42q0-28-18-46t-46-18q-27 0-45.5 18T341 107v42q0 21 12 37.5t31 22.5v6q-29 7-53 28q-20-20-54-28v-6q19-6 31-22.5t12-37.5v-42q0-28-18.5-46T256 43t-45.5 18t-18.5 46v42q0 21 12 37.5t31 22.5v6q-29 7-54 28q-24-21-53-28v-6q19-6 31-22.5t12-37.5v-42q0-28-18.5-46T107 43q-28 0-46 18t-18 46v42q0 21 11.5 37.5T85 209v6q-37 8-61 37.5T0 320v64h512v-64q0-38-24-67.5T427 215M85 107q0-22 22-22q9 0 15 6t6 16v42q0 10-6 16t-15 6q-22 0-22-22zm86 234H43v-21q0-27 18-45.5t46-18.5t49 23q15 19 15 41zm64-234q0-22 21-22t21 22v42q0 22-21 22t-21-22zm85 234H192v-21q0-23 15-41q21-23 49-23t49 23q15 18 15 41zm64-234q0-10 6-16t15-6q22 0 22 22v42q0 22-22 22q-9 0-15-6t-6-16zm85 234H341v-21q0-22 15-41q21-23 49-23t46 18.5t18 45.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PetroleumSolventSteam(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M245 365h-44V151h121v37h-77v55h73v36h-73zm11 147q106 0 181-75t75-181t-75-181T256 0T75 75T0 256t75 181t181 75m0-469q88 0 150.5 62.5T469 256t-62.5 150.5T256 469t-150.5-62.5T43 256t62.5-150.5T256 43"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Phone(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M186 84q-2-34-27.5-57.5T98 3H75Q2 21 0 116q-2 53 18.5 114T85 340q87 89 235 89q35 0 59.5-24t25.5-61q0-21-11.5-41.5T361 271q-40-17-77-2q-19 7-30 20L141 163q2-1 23-17q25-28 22-62m-54 34q-13 13-32 13q-13 0-19 12q-4 16 4 24l156 171q7 10 19 6q11-3 17-13q10-16 24-21q23-11 42 0q10 4 16 15t4 19q-1 19-13.5 31T320 387q-129 0-203-77q-39-41-56.5-93.5T43 118q2-66 38-73h15q18 0 32 11.5T143 86q3 18-11 32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Photobucket(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M441 46H313q-37-20-81-20t-81 20h-45q0-8-6-14t-15-6H43q-8 0-14 6t-6 14q-9 0-15 6.5T2 67v250q0 8 6 14.5t15 6.5h128q37 20 81 20t81-20h128q9 0 15-6.5t6-14.5V67q0-8-6-14.5T441 46M23 317V67h98q-57 52-57 125t57 125zm284 0q-18 10-35 15h-2q-5 2-15 4h-5q-12 2-18 2t-18-2h-5q-10-2-15-4h-2q-17-5-35-15q-72-44-72-125t72-125q18-10 35-15h2q5-2 15-4h5q12-2 18-2t18 2h5q10 2 15 4h2q17 5 35 15q72 44 72 125t-72 125m134 0h-98q57-52 57-125T343 67h98zm-94-125h-18q0-40-28.5-68.5T232 95V78q48 0 81.5 33.5T347 192"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Piano(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469 0H43Q25 0 12.5 12.5T0 43v277q0 17 12.5 30T43 363h426q18 0 30.5-13t12.5-30V43q0-18-12.5-30.5T469 0M43 43h426v85H43zm0 277V171h42v149zm64 0v-43h21V171h21v149zm64 0v-43h21V171h21v149zm64 0V171h42v149zm64 0v-43h21V171h21v149zm64 0v-43h21V171h21v149zm64 0v-43h21V171h21v149zm21-235q0 9-6.5 15.5T427 107q-9 0-15.5-6.5T405 85q0-8 6.5-14.5T427 64q8 0 14.5 6.5T448 85m-64 0q0 9-6.5 15.5T363 107q-9 0-15.5-6.5T341 85q0-8 6.5-14.5T363 64q8 0 14.5 6.5T384 85M64 64h43v43H64zm64 0h128v21H128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Picasa(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M138 333h301q-26 55-76 89.5T253 462h-42q-40-3-73-19zM327 22Q281 2 232 2q-41 0-80 15q3 3 87.5 79.5T327 176zm-200 5q-58 30-91.5 85T2 232q0 28 8 60q3-2 102.5-92.5T214 107q-2-2-44-40.5T127 27m-14 403V231q-4 4-49 45t-46 42q30 73 95 112M351 36v272h98q13-35 13-76q0-60-29.5-112.5T351 36"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Picassa(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M138 333h301q-26 55-76 89.5T253 462h-42q-40-3-73-19zM327 22Q281 2 232 2q-41 0-80 15q3 3 87.5 79.5T327 176zm-200 5q-58 30-91.5 85T2 232q0 28 8 60q3-2 102.5-92.5T214 107q-2-2-44-40.5T127 27m-14 403V231q-4 4-49 45t-46 42q30 73 95 112M351 36v272h98q13-35 13-76q0-60-29.5-112.5T351 36"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PiggyBank(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M477 131h-32q-7 0-10-7q-11-16-20-21l-2-2V3h-21q-27 0-62 27q-2 0-2 3q-17-9-64-9q-35 0-85 17t-75 43L61 71L44 52q-14-16-30 0q-14 14 0 30l26 25l36 13q-25 34-25 85v32l19 175q1 16 13.5 27.5T113 451h29q32 0 41-30l15-43q33 10 72 7l11 36q11 30 41 30h29q17 0 29.5-11.5T394 412l9-85q29-29 40-60q3-8 13-8h4q24 0 42-17.5t18-42.5v-26q0-17-12.5-29.5T477 131m0 68q0 6-6.5 11.5T456 216q-36 0-53 36q-16 34-37 49l-6 5l-9 102h-29l-24-68l-17 2q-11 2-28 2q-29 0-64-11l-21-8l-26 83h-29L93 235v-30q0-42 26-74l17-20q49-44 128-44q39 0 53 8q18 10 41-13q2-1 6-4.5t5-3.5v64l6 6q5 6 11 11q9 9 10 11q17 27 47 27h34zm-128-68q0 8-6 14.5t-15 6.5t-15-6.5t-6-14.5q0-9 6-15.5t15-6.5t15 6.5t6 15.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PiggyBankCoins(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469 192h-32q-7 0-10-6q-9-14-20-22l-2-2V64h-21q-32 0-60 28l-2 2v-7q0-35-25.5-61T235 0q-35 0-60.5 25T149 87q0 13 2 20q-27 10-55 38l-43-15l-17-17q-14-14-30 0q-13 15 0 30l26 26l36 12q-25 38-25 86v32l19 175q2 16 14.5 27t28.5 11h29q13 0 25-8t16-22l15-43q33 10 72 7l13 36q12 28 41 28h30q16 0 28.5-11.5T388 471l9-85q29-29 40-60q3-8 13-8h4q24 0 42-18t18-42v-23q-2-18-14.5-30.5T469 192M237 43q18 0 31.5 13T282 87q0 19-13 32t-32 13q-19-2-32-14.5T192 87t13-31t32-13m232 217q0 7-5 12t-12 5h-4q-36 0-53 37q-9 22-34 49l-7 6l-11 100h-29l-24-68l-17 2q-11 2-28 2q-29 0-62-10l-21-9l-28 83h-29L85 297v-30q0-43 26-75q12-4 15-13v-4q13-13 45-30q26 30 64 30q47 0 72-38h2q14 7 41-11q2-1 6-4.5t7-4.5v64l6 7q6 6 11 10l5.5 5.5l4.5 5.5q18 28 47 28h32zm-128-68q0 9-6 15t-15 6t-15-6t-6-15t6-15t15-6t15 6t6 15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pin(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 277q-9 0-15 6t-6 16q0 9 6 15t15 6h128q10 0 16 6t6 15v150q0 21 21 21t21-21V341q0-9 6-15t16-6h128q9 0 15-6t6-15q0-10-6-16t-15-6h-22L277 43q22 0 22-22q0-9-6-15t-16-6H107Q97 0 91 6t-6 15q0 22 22 22L43 277zM149 43h86l64 234H85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PinMap(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M176 0Q101 0 53 53.5T5 192q0 95 154 311l17 24l17-24q154-216 154-311q0-85-48-138.5T176 0m0 454q-47-68-87.5-144.5T48 192q0-61 33.5-105T176 43t94.5 44T304 192q0 41-40.5 117.5T176 454m0-369q-35 0-60 25.5T91 171t25 60t60 25t60-25t25-60t-25-60.5T176 85m0 128q-17 0-30-12.5T133 171q0-18 13-30.5t30-12.5t30 12.5t13 30.5q0 17-13 29.5T176 213"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pinboard(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m280 260l90-95l-87 21L145 58V2L-1 150l66-3l118 145l-14 78l87-88l204 180z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ping(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 86q24-4 52-4q36 0 54 15q19 14 19 41q0 26-16 42q-22 20-59 20q-7 0-17-2v70H2zm33 86q5 1 17 1q41 0 41-34q0-31-38-31q-12 0-20 2zm103 96V135h34v133zm57-94q0-8-.5-21t-.5-18h29l2 20h1q13-23 43-23q20 0 34 14t14 43v79h-34v-75q0-34-26-34q-19 0-27 20q-1 3-1 11v78h-34zm266 75q0 41-20 60q-17 16-52 16q-28 0-47-11l8-25q16 10 39 10q39 0 39-40v-12h-1q-12 20-39 20q-25 0-41-18t-16-47q0-32 17.5-51t43.5-19t39 21l2-18h29q0 3-.5 15.5T461 173zm-34-61q0-6-1-9q-5-21-28-21q-15 0-24 11.5t-9 31.5q0 18 9 29.5t24 11.5q22 0 28-21q1-3 1-11zM155 59q-13 0-22 8.5T124 89t9 21.5t22 8.5t22-8.5t9-21.5t-9-21.5t-22-8.5m0 54q-10 0-17-7t-7-17t7-17t17-7t17 7t7 17t-7 17t-17 7m11-16q0 5-4 5q-5 2-5-4q0-5 4-5h2V81l-12 2v16q0 5-4 5q-5 0-5-4q0-5 4-5h2V77l18-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pingchat(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M143 330h139q0 56-40.5 95.5T143 465q-57 0-98-39.5T4 330q0-55 41-95t98-40zm322-193q0 55-40 93.5T329 269H193V130q2-53 41.5-89.5T329 4q56 0 96 39t40 94m-116 66q0-8-6-14t-15-6q-8 0-14 6t-6 14t6 14t14 6q9 0 15-6t6-14m0-134q0-8-6-14t-15-6q-8 0-14 6t-6 14v3l6 79q0 15 14 15q15 0 15-15l6-79z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pizza(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m2 400l21 9q100 42 201 42t201-42l21-9L224 5zm248-262q-26 10-26 38q0 17 12.5 30t30.5 13q15 0 23-9l58 103q-61 30-126 30T98 313L224 91zM62 379l15-28q74 36 145 36q75 0 145-36l15 28q-75 30-158.5 30T62 379m183-96q0 17-12.5 29.5T203 325q-18 0-30.5-12.5T160 283q0-18 12.5-30.5T203 240q17 0 29.5 12.5T245 283"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plane(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m489 230l-171-74l2-49q0-34-21.5-70.5T245 0t-53 36.5t-21 72.5l2 49l-150 75Q0 243 0 271v77l181-41l5 86q-79 32-79 98v21h277v-21q0-71-79-101l4-89l203 45v-77q0-24-23-39M43 294v-23l132-66l4 60zm111 175q5-12 34-32l2 32zm79 0l-20-362q0-19 12-41.5T245 43t20 22t12 40l-19 364zm104 0h-38l2-34q29 13 36 34m132-175l-158-34l3-57l153 68v23z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Play(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 427q4 2 9 2q11 0 15-6l170-192q12-16 0-28L42 11Q34-2 18 5Q5 9 5 24v384q0 15 13 19M48 79l122 137L48 353z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Playstation(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m250 371l-71-30V13q6 1 16.5 3t38 10T283 43t43 25t27 34q10 32 6 55.5T346 194t-16 14q-14 2-25-2.5t-15-9.5l-4-5V89l-35-15zm18-45v38q166-53 190-74q16-14-19-38q-42-30-78-25q-2 1-4 1q-46 3-88 21v35l106-26l38 17zm-106-3q-22 11-70 11q-39-1-65-12T2 291q0-21 48-42.5T161 217v40l-85 23q-14 18 5 19q20 0 40.5-2.5T152 291l10-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plixi(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M67 245q3 52 13 79q-24-5-32-7q-24-4-31-10q-16-12-16-30q0-16 10-26.5T36 238q11-1 31 7M292 58Q256-3 197 2q-30 2-58 27q-46 40-60 108q-7 33-5 77q0 28 2 40q6 60 23 103q17 47 46 80q2 2 6.5 7t8 8.5t6.5 5.5q19 10 36 0t17-30q0-10-12-25q-2-2-5-7.5t-5-8.5q-1-4-4-7q-21-26-36-99q-3-15-7-49q-1-11-1-34q0-75 28-110q14-19 30-18q23 0 35 35q19 58-8 120q-12 30-31 39q-16 8-39 3q-2 20 20 76q28 1 53-15q70-43 81-146q6-72-26-124"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plurk(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M289 4H6v461h117V338h283V4zm0 218H123V120h166z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plus(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 213h150v150q0 21 21 21t21-21V213h150q21 0 21-21t-21-21H213V21q0-21-21-21t-21 21v150H21q-21 0-21 21t21 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusBox(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M405 0H107Q62 0 31 31T0 107v298q0 45 31 76t76 31h298q45 0 76-31t31-76V107q0-45-31-76T405 0m64 405q0 28-18 46t-46 18H107q-28 0-46-18t-18-46V107q0-28 18-46t46-18h298q28 0 46 18t18 46zM363 235h-86v-86q0-21-21-21t-21 21v86h-86q-21 0-21 21t21 21h86v86q0 21 21 21t21-21v-86h86q21 0 21-21t-21-21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusCircle(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0Q150 0 75 75T0 256t75 181t181 75t181-75t75-181t-75-181T256 0m0 469q-88 0-150.5-62.5T43 256t62.5-150.5T256 43t150.5 62.5T469 256t-62.5 150.5T256 469m107-234h-86v-86q0-21-21-21t-21 21v86h-86q-21 0-21 21t21 21h86v86q0 21 21 21t21-21v-86h86q21 0 21-21t-21-21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Podcast(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 250v-9q0-99 67.5-169T232 2t162.5 70T462 241v9q-4-96-70.5-163T232 20T72.5 87T2 250m230-143q58 0 100 41t46 101v-9q0-62-42.5-106.5T232 89t-103.5 44.5T86 240v9q4-60 46-101t100-41m0 44q-19 0-33.5 15T184 201t14.5 35t33.5 15t33.5-15t14.5-35t-14.5-35t-33.5-15m0 104q-18 0-27.5 1.5t-21 8.5t-16.5 23t-5 42q0 55 20.5 93.5T232 462t49.5-38.5T302 330q0-26-5-42t-16.5-23t-21-8.5T232 255"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Posterous(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M61 462V288l2-3h5l1 1q30 50 98 50q60 0 103.5-44T314 164q0-72-38-117T177 2Q100 2 63 63l-2 1h-5L53 8H6q3 48 3 104v350zm0-320q0-7 5-24q7-34 33.5-55T159 42q44 0 72.5 35t28.5 90q0 58-28 93.5T157 296q-33 0-58.5-20T64 224q-3-11-3-25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Power(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M312 107q-19-9-28 10q-8 21 11 28q42 23 67 64t25 90q0 70-50 120t-121 50t-121-50t-50-120q0-47 25-88t67-64q18-10 11-30q-13-14-30-8q-53 27-84 78T3 299q0 88 62.5 150.5T216 512t150.5-62.5T429 299q0-61-31.5-113T312 107m-75 149V21q0-21-21-21t-21 21v235q0 21 21 21t21-21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Preston(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M510 371q-5-6-13.5-17.5T484 335q-3-3 0-6q20-65 9-122q-18-98-99.5-156.5T213 6Q127 19 68 82.5T4 233q-6 110 71 189q55 56 132 69.5T354 471q5-2 20.5-9t24.5-10q1 0 3-1t3-1q50-25 92-64q1-1 4-3.5t5-5t4-4.5zM369 190q8-18 26-34q11-12 30-15q17-2 36 13q9 6 13 19q24 74 2 145q0 2-2 4q-3-5-7-9.5t-8.5-10t-6.5-8.5q-48-61-81-91q-6-5-2-13m-68 290q-89 20-168-25T26 324v-6q24-92 91-139q63-43 135-34q29 3 68 21q11 9 11 20q0 47-9 83q-2 8-17 45q-8 16-28 64q0 1-1 4t-1 4q-51-22-72-38l6 6q38 34 60 47q2 0 2 4q0 33 26 47q5 5 21 9q4 1 11.5 2t11.5 2q-13 8-40 15m200-100q-37 22-81 30q-73 12-145-7h-2q0-6 2-8q0-1 1-4.5t1-4.5q107 39 231-8q-3 0-4.5 1t-2.5 1M245 277q10 8 22.5 8t22.5-10t10-22.5t-8-22.5q-1-1-2-3t-2-3q-2-7-6.5-21.5T275 181q-4-12-6-21q-2 2-2 4q-3 9-9.5 29.5T247 224q0 4-4 4q-10 3-28 7.5t-23 5.5q-40 11-43 11h7q52 12 85 21zm22-53q5 0 8.5 4t3.5 9t-4 9t-8 4q-5 0-9-4t-4-9q0-13 13-13m194 51q6 0 10-8q1-2 3.5-5.5t3.5-5.5l4-4q-2-2-4-5q-7-7-7-12q-1-6-4-20.5t-4-20.5q0-1-1-3.5t-1-4.5v2q-2 10-6.5 29t-6.5 28q-2 7 2 22q0 8 11 8m-4-40q3-3 6 0q2 2 2 6q0 9-6 9q-5 0-5-9q3-2 3-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Previous(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M27 67Q5 67 5 88v256q0 21 22 21q21 0 21-21v-73l134 150q8 6 15 6q5 0 9-2q13-3 13-19V24q0-15-13-19t-24 6L48 161V88q0-21-21-21m149 12v271L54 216z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Printer(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469 128h-42V0H85v128H43q-18 0-30.5 12.5T0 171v213h85v128h342V384h85V171q0-18-12.5-30.5T469 128M128 43h256v85H128zm256 426H128V299h256zm85-128h-42v-85H85v85H43V171h426zm-42-128q0 9-6.5 15.5T405 235q-8 0-14.5-6.5T384 213q0-8 6.5-14.5T405 192q9 0 15.5 6.5T427 213M149 320h214v43H149zm0 64h214v43H149z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PrisonSchoolBus(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469 149h-42q-22 0-22-21V85q0-17-12.5-29.5T363 43H43q-18 0-30.5 12.5T0 85v192q0 18 12.5 30.5T43 320h4q6 19 22.5 31t37.5 12q20 0 36.5-12.5T166 320h177q15 43 60 43q20 0 36.5-12.5T463 320h6q18 0 30.5-12.5T512 277v-85q0-17-12.5-30T469 149M107 320q-10 0-16-6t-6-15q0-22 22-22q9 0 15 6t6 16q0 9-6 15t-15 6m298 0q-9 0-15-6t-6-15q0-10 6-16t15-6q22 0 22 22q0 9-6 15t-16 6m60-43q-6-18-23-30t-37-12t-36.5 11.5T346 277H166q-6-18-22.5-30T107 235q-21 0-37.5 11.5T47 277h-4V85h320v22h-22q-21 0-21 21v21q0 10 6 16t15 6h39q18 21 47 21h27q9 11 15 15v70zM277 107H85q-21 0-21 21v21q0 10 6 16t15 6h192q22 0 22-22v-21q0-21-22-21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Promo(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M91 347v85l85-21l64 64l64-64l85 21v-85l86-22l-43-85l43-85l-86-22V48l-85 21l-64-64l-64 64l-85-21v85L5 155l43 85l-43 85zm-5-126l-19-39l34-8q15-3 23.5-14.5T133 133v-30l32 9h11q17 0 30-13l34-34l34 34q13 13 30 13q6 0 11-2l32-7v30q0 15 8.5 26.5T379 174l34 8l-19 39q-10 19 0 38l19 39l-34 8q-15 3-23.5 14.5T347 347v30l-32-9h-11q-17 0-30 13l-34 34l-34-34q-13-13-30-13q-6 0-11 2l-32 7v-30q0-15-8.5-26.5T101 306l-34-8l19-39q10-19 0-38m133 92l100-101q14-14 0-30q-15-13-30 0l-70 71l-28-28q-15-13-30 0q-13 15 0 30z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pull(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M45 512h342q17 0 29.5-12.5T429 469V192q0-48-31-83.5T321 66q2-4 2-13q0-22-15.5-37.5T269 0H163q-23 0-38.5 15.5T109 53q0 9 2 13q-46 7-77 42.5T3 192v277q0 18 12.5 30.5T45 512m0-85h43v42H45zm64-43V277h214v192H109zm278 85h-43v-42h43zM163 43h106q11 0 11 10q0 11-11 11H163q-11 0-11-11q0-10 11-10m-32 64h170q35 0 60.5 25t25.5 60v192h-43V203q0-11-11-11q-10 0-10 11v32H109v-32q0-11-10-11q-11 0-11 11v181H45V192q0-35 25.5-60t60.5-25"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Puzzle(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M69 341H47q-19 0-31.5 12.5T3 384v43q0 35 26 60t63 25h45q36 0 36-43v-21q0-8 10-14.5t22-6.5h24q12 0 22 6.5t10 14.5v21q0 43 36 43h68q28 0 47-18.5t19-45.5v-64q0-18-12.5-30.5T387 341h-24q-19 0-19-21v-21q0-22 19-22h24q17 0 30.5-12.5T431 235v-22q0-27-19-45.5T365 149h-66q24-17 24-42q0-45-31-76T216 0q-44 0-75.5 31T109 107q0 29 28 42H69q-27 0-46.5 18.5T3 213v22q0 17 12.5 29.5T47 277h24q17 0 17 22v21q0 21-19 21m0-106H45v-22q0-9 7-15t17-6h83q20 0 31.5-12t11.5-31q0-5-1.5-9t-3-6.5t-5.5-6t-5.5-4.5l-7.5-5l-5-3q-13-4-13-8q0-28 18.5-46T218 43q26 0 45 18.5t19 45.5q0 5-15 10l-6 3l-6 3l-5.5 3.5l-5 4.5l-3.5 4.5l-3 6l-1 7.5q0 43 47 43h79q10 0 17 6t7 15l-2 22h-24q-27 0-44.5 18T299 299v21q0 27 17.5 45.5T361 384h26v64q0 9-7 15t-17 6h-62v-21q0-27-23-45.5T227 384h-24q-27 0-49.5 18.5T131 448v21H92q-19 0-33-12.5T45 427l2-43h24q25 0 42.5-18.5T131 320v-21q0-28-17.5-46T69 235"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Qik(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M328 462h63q8-47 23-150q1-5 12-62t13-87q6-74-34-114.5T291 3Q155-8 71 76Q12 135 8 213q-4 71 46.5 115.5T193 366q20-1 41-11q11-4 26-16.5t21-15.5l-7.5 36l-3.5 35l3.5 35.5l18 22.5zm-65-205q-15 11-41.5 13.5T173 261q-33-25-23-77.5t40-75.5q16-11 41-14q31-4 51 9q30 23 21.5 77.5T263 257"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Quik(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M328 462h63q8-47 23-150q1-5 12-62t13-87q6-74-34-114.5T291 3Q155-8 71 76Q12 135 8 213q-4 71 46.5 115.5T193 366q20-1 41-11q11-4 26-16.5t21-15.5l-7.5 36l-3.5 35l3.5 35.5l18 22.5zm-65-205q-15 11-41.5 13.5T173 261q-33-25-23-77.5t40-75.5q16-11 41-14q31-4 51 9q30 23 21.5 77.5T263 257"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Quora(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M316 390q44-27 70.5-74T413 214q0-86-60-147T208 6T63 67T3 214t60 146.5T208 421q26 0 48-6q35 64 112 49v-35q-40-11-52-39m-2-149q0 52-29 89q-40-44-102-41v41q30 1 51 37q-12 3-24 3q-44 0-74.5-37.5T105 241v-54q0-53 31-91t74-38t73.5 38t30.5 91z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Quote(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M112 171q-37 0-64 21v-21q0-53 37.5-90.5T176 43q21 0 21-22q0-21-21-21q-70 0-120.5 50T5 171v106q0 45 31 76t76 31t76-31t31-76q0-44-31-75t-76-31m256 0q-37 0-64 21v-21q0-53 37.5-90.5T432 43q21 0 21-22q0-21-21-21q-70 0-120.5 50T261 171v106q0 45 31 76t76 31t76-31t31-76q0-44-31-75t-76-31"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Radio(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M152 341q62 0 105.5-43.5T301 192T257.5 86.5T152 43T46.5 86.5T3 192t43.5 105.5T152 341m0-256q45 0 76 31t31 76t-31 76t-76 31t-76-31t-31-76t31-76t76-31m43 107q0 18-12.5 30.5T152 235t-30.5-12.5T109 192t12.5-30.5T152 149t30.5 12.5T195 192"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RadioEmpty(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M301 192q0-62-43.5-105.5T152 43T46.5 86.5T3 192t43.5 105.5T152 341t105.5-43.5T301 192m-256 0q0-45 31-76t76-31t76 31t31 76t-31 76t-76 31t-76-31t-31-76"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ram(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 43H64q-27 0-45.5 18T0 107v192q0 17 12.5 29.5T43 341h426q18 0 30.5-12.5T512 299V107q0-28-18.5-46T448 43m-21 256v-64h-22v64h-21v-64h-21v64h-22v-64h-21v64h-43v-64q0-22-21-22t-21 22v64h-22v-64h-21v64h-21v-64h-22v64h-21v-64h-21v64H85v-64H64v64H43V107q0-22 21-22h384q21 0 21 22v192zM107 107H85q-21 0-21 21v64q0 21 21 21h22q21 0 21-21v-64q0-21-21-21m85 0h-21q-22 0-22 21v64q0 21 22 21h21q21 0 21-21v-64q0-21-21-21m149 0h-21q-21 0-21 21v64q0 21 21 21h21q22 0 22-21v-64q0-21-22-21m86 0h-22q-21 0-21 21v64q0 21 21 21h22q21 0 21-21v-64q0-21-21-21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Random(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M70 365q74 0 118-57q0-4-5-7l-19-38q-13 27-38.5 43.5T70 323H21q-8 0-14.5 6.5T0 344t6.5 14.5T21 365zM442 9q-16-14-30 0q-15 15 0 30l27 28h-83q-73 0-117 57q0 3 4 7l19 38q13-27 38.5-43.5T356 109h83l-27 28q-15 15 0 30q6 6 15 6q7 0 15-6l64-64q13-15 0-30zm0 256q-16-14-30 0q-15 15 0 30l27 28h-83q-30 0-56-16.5T260 263l-23-47l-24-47l-10-19q-18-38-54-60.5T70 67H21q-8 0-14.5 6.5T0 88t6.5 14.5T21 109h49q64 0 96 60l24 47l23 47l11 19q20 38 55.5 60.5T358 365h84l-28 28q-15 15 0 30q6 6 15 6q8 0 15-6l64-64q13-15 0-30z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rdio(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M400 165q2 18 2 27q0 79-58.5 135.5T202 384T60.5 328T2 192T60.5 56T202 0q32 0 62 9v117q-7-4-18.5-8.5t-41-4T151 131t-34.5 36.5T107 198l1 12q0 3 .5 7t5 15t12 19t23 15t36.5 7q37 0 64-19.5t35-38.5l9-19V21q23 11 46 30q59 36 115 35q1 0 2.5.5T459 88t2 4.5t1 7.5q0 16-12 29q-16 25-50 36"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Readernaut(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18 187q11-37 45-44q22-51 69-82.5T237 29q59 0 106.5 31.5T412 144h3V54q-17-7-17-25q0-11 8-19t19-8q12 0 20 8t8 19q0 21-21 26v98q16 11 23 34q9 45-6 56q-7 11-29 19q-2 9-6 18q-4 1-12.5 2.5T389 285q20-35 20-77q0-67-50.5-114.5T236 46T113 93.5T62 208q0 43 22 79q-4-1-11.5-2.5T61 282q-1-3-3-9.5t-3-9.5v-1q-10-4-18-8.5T26 246l-2-3q-7-5-8.5-19t.5-26zm214 200q-29-36-86.5-57.5T45 304l-43-4v61q41 0 80 10.5t65 25.5t46 29.5t29 25.5l10 10q3-4 9.5-11t28.5-24t47-30.5t64.5-24.5t80.5-11v-61q-7 0-18.5 1t-43 6.5t-59.5 14t-59.5 26T232 387m157-151q0 27-10 54q-104 24-147 76q-41-50-140-74q-11-27-11-56q0-63 45-107t109-44t109 44t45 107m-227-24q0-7-5-12t-12-5t-12 5t-5 12q0 17 17 17q7 0 12-5t5-12m41-43q0-11-7.5-18.5T177 143t-19 7.5t-8 18.5t8 18.5t19 7.5t18.5-7.5T203 169"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reddit(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M331 209q2 6 2 9q0 10-6 18q-7 9-15 10q-6 2-8 2q-10 0-19-7t-10-17q-1-2-1-5q0-9 5-17q6-8 14-11q3-1 10-1q9 0 17 5t11 14m-18 74q-6-2-8-2q-3 0-7 2q-33 19-68 19q-28 0-52-12q-1 0-8-5q-1-1-2.5-2t-2.5-1t-3-.5t-3-.5q-4 0-5 1h-1q-5 2-6 6q-2 4-2 8q0 5 2 7q2 3 5 5q33 23 77 23q36 0 71-16q3-2 11-6l2-1l3-3l3-6q0-1 1-3v-1q0-3-2-5q-1-3-5-7m-162-37h1q6 2 9 2q10 0 20-8q9-9 9-20v-2q0-11-9-20q-10-8-20-8h-6q-17 3-22 20q-2 6-2 8q0 10 6 18q5 7 14 10m311-72v3q0 17-9 31t-22 21q0 2 .5 6t.5 7q0 36-23 67q-37 48-110 67q-35 8-67 8q-53 0-100-19q-56-22-83-66q-16-25-16-57v-13q-14-9-21-20q-10-15-10-31q0-25 17-40q16-18 39-18h6q13 0 22 4q9 2 18 10l6-3q42-25 110-29q0-36 12-56q11-23 36-29q12-2 18-2q22 0 47 10q8-14 26-22q11-3 20-3q13 0 22 4q16 7 23 20q10 13 10 29q0 4-1 6q-1 20-18 35q-15 13-35 13h-7q-20-2-34-17q-15-15-15-35v-2q-20-10-38-10h-5q-17 2-23 16q-7 13-9 43q62 3 108 30q1 0 1 1l2 1q3-3 7-5q14-10 34-10q10 0 15 2h1q16 4 31 19q12 14 14 34M353 52v2q0 9 8 17q9 7 17 7h2q8 0 17-7q8-6 8-17v-2q0-11-8-17q-8-8-18-8q-4 0-6 1q-8 1-14 9q-6 6-6 15M79 152q-10-4-15-4h-3q-13 0-20 9q-9 7-9 19q-1 1-1 2q0 5 4 13q4 6 7 8q13-26 37-47m323 91q0-24-15-47q-27-37-81-54q-12-4-18-5q-27-6-56-6q-36 0-74 11q-54 17-81 54q-15 23-15 47q0 6 2 18q6 20 20 36q11 15 31 27q1 0 3.5 2t3.5 2q48 27 110 27q13 0 20-1q63-5 108-37q21-18 26-26q13-20 15-34q1-5 1-14m31-66q0-6-3-12q-4-8-12-13q-8-4-18-4q-8 0-15 4q24 21 37 47q3-2 7-8q4-10 4-14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Resize(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m44 514l135-135v56q0 21 21 21t21-21V328q0-21-21-21H93q-21 0-21 21t21 21h56L14 484q-14 14 0 30q8 6 15 6q9 0 15-6m335-165h56q21 0 21-21t-21-21H328q-21 0-21 21v107q0 21 21 21t21-21v-56l135 135q6 6 15 6q7 0 15-6q14-16 0-30zM484 14L349 149V93q0-21-21-21t-21 21v107q0 21 21 21h107q21 0 21-21t-21-21h-56L514 44q14-14 0-30q-16-14-30 0M14 14Q0 30 14 44l135 135H93q-21 0-21 21t21 21h107q21 0 21-21V93q0-21-21-21t-21 21v56L44 14Q30 0 14 14"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Retweet(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M94 328q-7 0-12-5.5T77 310V170H2L104 58l103 112h-76v103h81l53 55zm368-114h-75V74q0-7-5-12.5T370 56H199l53 55h81v103h-76l103 113z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RetweetOne(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m486 154l-30 21v-68q0-35-25-60.5T371 21H157q-35 0-60 25.5T72 107q0 9 6 15t15 6q10 0 16-6t6-15q0-18 12.5-30.5T157 64h214q17 0 29.5 12.5T413 107v66l-30-22q-19-11-29 7q-12 19 6 30l75 47l76-47q17-11 7-30q-15-15-32-4m-51 102q-10 0-16 6t-6 15q0 18-12.5 30.5T371 320H157q-17 0-29.5-12.5T115 277v-66l30 22q2 0 6 1t6 1q8 0 17-9q12-19-6-30l-75-47l-76 47q-17 11-7 30q11 17 30 7l30-22v66q0 35 25 60.5t60 25.5h214q35 0 60-25.5t25-60.5q1-9-4.5-15t-14.5-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rewind(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M177 423q6 6 15 6q4 0 9-2q12-3 12-19V271l135 150q6 6 15 6q4 0 8-2q13-3 13-19V24q0-15-13-19q-14-3-23 6L213 161V24q0-15-12-19q-15-3-24 6L6 203q-11 17 0 28zM341 79v271L220 216zm-170 0v271L49 216z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Right(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9 388q8 4 15 4q11 0 17-6l162-186L41 14Q26-1 11 12Q-4 29 9 42l137 156L9 354q-13 19 0 34"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightArrowCircle(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0Q150 0 75 75T0 256t75 181t181 75t181-75t75-181t-75-181T256 0m0 469q-88 0-150.5-62.5T43 256t62.5-150.5T256 43t150.5 62.5T469 256t-62.5 150.5T256 469m107-219q-3 0 0 0q0-3-2-5q0-2-3-2l-42-64q-12-17-30-6q-18 10-7 30l22 30H171q-22 0-22 21t22 21h130l-22 30q-11 20 7 30q19 10 30-6l42-64q0-3 3-3v-10q2 2 2-2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Road(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 495q14 17 34 17h412q22 0 34-17q14-18 6-38L367 30Q356 0 326 0H186q-13 0-25 8.5T145 30L9 457q-5 20 8 38M186 43h49v64q0 21 21 21t21-21V43h49l137 426H299v-85q0-21-22-21h-42q-22 0-22 21v85H51zm53 256h34q10 0 16.5-8t4.5-18l-15-85q-3-17-21-17h-6q-16 0-22 17l-15 85q0 10 7 18t17 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Robo(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M215 186q0-7 5-12t12-5t12 5t5 12t-5 11.5t-12 4.5t-12-4.5t-5-11.5M164 43v23h21V44q12-7 12-20q0-9-6.5-15.5T175 2t-15.5 6.5T153 24q0 13 11 19m-46 159q16 0 16-16q0-17-16-17q-7 0-12 5t-5 12t5 11.5t12 4.5m57 260q40 0 54-38H121q14 38 54 38M2 394q0-13 12-20v-94q0-14 9.5-24T47 244v-45q0-3 1-9t5.5-22T65 138.5t22-28T121 89q27-11 55-11t55 11q31 12 49.5 39.5T302 178l3 21v45q14 2 23.5 12t9.5 24v94q12 7 12 20q0 10-6.5 16.5T328 417t-15.5-6.5T306 394q0-13 12-20v-63q-5 2-13 4v33q-2 28-21.5 48T237 419H115q-27-3-46.5-23T47 348v-33q-8-2-13-4v63q12 7 12 20q0 10-6.5 16.5T24 417t-15.5-6.5T2 394m211-116q0-11-11-11q-10 0-10 11t10 11q11 0 11-11m26 0q0-11-11-11q-10 0-10 11t10 11q11 0 11-11m26 0q0-11-10-11q-11 0-11 11t11 11q10 0 10-11M86 189q0 12 6 19.5t12 9.5l5 1q9 3 31 0q3 0 18-1.5t17-2.5h2q2 1 17 2.5t18 1.5q22 3 31 0q3 0 6.5-1t10-9t6.5-20q0-32-23-51q-25-23-66-23h-2q-41 0-66 23q-10 8-16 21t-7 21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RowSetting(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m439 147l24-47l-51-51l-47 24q-21-13-56-24L292 0h-72l-17 49q-35 11-56 24l-47-24l-51 51l24 47q-13 21-24 56L0 220v72l49 17q5 27 24 56l-24 47l51 51l47-24q21 13 56 24l17 49h72l17-49q27-5 56-24l47 24l51-51l-24-47q13-21 24-56l49-17v-72l-49-17q-11-35-24-56m30 115l-42 15l-2 13q-7 35-26 62l-9 11l22 40l-9 9l-40-22l-11 7q-27 18-62 25l-13 5l-15 42h-12l-15-42l-13-2q-35-7-62-26l-11-9l-40 22l-9-9l22-40l-7-11q-18-27-25-62l-5-13l-42-15v-12l42-15l2-13q7-35 26-62l9-11l-22-40l9-9l40 22l11-7q27-18 62-25l13-5l15-42h12l15 42l13 2q35 7 62 26l11 9l40-22l9 9l-22 40l7 11q18 27 25 62l5 13l42 15zM256 149q-45 0-76 31t-31 76t31 76t76 31t76-31t31-76t-31-76t-76-31m0 171q-27 0-45.5-18.5T192 256t18.5-45.5T256 192t45.5 18.5T320 256t-18.5 45.5T256 320"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rss(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M106 358q18 18 18 43t-18 43t-43 18t-43-18Q2 427 2 401q0-27 18-43q18-18 43-18t43 18M2 158v88q88 0 152 64q63 61 63 152h89q0-126-90-214q-88-90-214-90M2 2v88q154 0 263 109t109 263h88q0-94-35.5-178T327 137q-63-64-147-99.5T2 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RssIcon(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M67 301q-28 0-46 18.5T3 365q0 28 18 46t46 18q27 0 45.5-18t18.5-46q0-27-18.5-45.5T67 301m0 86q-22 0-22-22q0-9 6-15t16-6q9 0 15 6t6 15q0 10-6 16t-15 6m0-235q-10 0-16 6t-6 15q0 22 22 22q71 0 120.5 49.5T237 365q0 22 22 22q9 0 15-6t6-16q0-88-62.5-150.5T67 152M67 3Q45 3 45 24t22 21q133 0 226.5 93.5T387 365q0 22 21 22t21-22q0-150-106-256T67 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Safe(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 0H64Q37 0 18.5 18.5T0 64v320q0 27 18.5 45.5T64 448v21q0 18 12.5 30.5T107 512h42q18 0 30.5-12.5T192 469v-21h149v21q0 18 13 30.5t30 12.5h43q17 0 29.5-12.5T469 469v-25q43-15 43-60V64q0-27-18.5-45.5T448 0M149 469h-42v-21h42zm278 0h-43v-21h43zm42-85q0 21-21 21H64q-21 0-21-21V64q0-21 21-21h384q21 0 21 21zM405 64H107q-18 0-30.5 12.5T64 107v234q0 18 12.5 30.5T107 384h298q18 0 30.5-12.5T448 341V107q0-18-12.5-30.5T405 64M107 341V107h298v234zm213-213q-27 0-45.5 18.5T256 192q0 20 12 36.5t31 23.5v47q0 21 21 21t21-21v-47q43-15 43-60q0-27-18.5-45.5T320 128m0 85q-21 0-21-21t21-21t21 21t-21 21M128 320h107V128H128zm43-149h21v106h-21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SaleTag(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m44 278l160 192q21 24 49 24q25 0 43-15l139-117q20-18 23-45q3-25-15-49L283 76q-26-34-68-43L76 1q-20-6-38 9L21 25Q6 36 6 61l8 143q5 43 30 74M68 42l141 34q22 5 44 27l160 192q8 8 5 18q0 6-9 17L270 447q-7 6-17 5.5t-15-7.5L78 253q-19-24-19-49L49 57zm57 70q0 9-6 15t-15 6t-15-6t-6-15t6-15t15-6t15 6t6 15m11 149l196 62h7q14 0 21-15q3-8-1-16t-12-11l-196-62q-8-4-16.5 0T123 231q-6 23 13 30m160 96q0 13-9 22.5t-23 9.5t-23-9.5t-9-22.5t9.5-22.5T264 325t22.5 9.5T296 357m-64-170q0 13-9.5 22.5T200 219t-22.5-9.5T168 187t9-22.5t23-9.5t23 9.5t9 22.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Save(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M69 5Q42 5 23.5 23.5T5 69v342q0 27 18.5 45.5T69 475h342q27 0 45.5-18.5T475 411V170L357 5zm150 43v64h42V48h22v85H133V48zm-86 384V304h214v128zm299-21q0 9-6 15t-15 6h-22V304q0-17-12.5-30T347 261H133q-17 0-29.5 13T91 304v128H69q-9 0-15-6t-6-15V69q0-9 6-15t15-6h22v85q0 18 12.5 30.5T133 176h150q17 0 29.5-12.5T325 133V48h11l96 134z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Scissors(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M171 344q0-11-9-38l66-41l-40-26l-54 32q-30-12-49-12q-35 0-60 25T0 344t25 60t60 25t60.5-25t25.5-60m-128 0q0-17 12.5-30T85 301q14 0 24 7l8 6l5 11q6 9 6 19q0 17-12.5 30T85 387q-17 0-29.5-13T43 344M480 67q-7 0-15 4l-158 96l41 26L512 88q-6-21-32-21M269 190l-107-64q9-27 9-38v-6q-4-34-28-56.5T85 3Q50 3 25 28T0 88q0 32 21 56t52 29h12q25 0 47-15l94 58l41 26l183 119q8 4 21 4q33 0 41-21L309 216zm-145-83l-7 11l-8 4q-9 9-24 9h-4q-17-4-27.5-15.5T43 88q0-17 12.5-30T85 45t29.5 11.5T128 86v2q0 7-4 19"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Scribd(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M269 412q-2 2-5.5 6t-5.5 6q-39 38-96 38q-55 0-96-29q-42-30-65-89l1-1q-2-8-2-13q0-21 15-35.5T50 280q21 0 35.5 14.5T100 330q0 19-15 36q31 34 75 34q32 0 54-19q22-20 22-49q0-18-12-36q-11-18-28-29q-18-11-53-26q-38-16-56-27q-22-15-34-30q-14-16-21-34q-7-15-7-36q0-47 37-80q40-32 92-32q34 0 72 15q28 11 52 37q17 15 17 37q0 20-14.5 34.5T245 140t-35.5-14.5T195 91q0-9 5-21q-15-7-41-7q-32 0-52 14q-20 15-20 37q0 20 17 35q18 15 60 33q43 19 65 33q20 13 38 34q16 21 23 40q0 2 1 2q-35 29-35 73q0 25 13 48m73-105q-24 0-41 17t-17 40q0 24 17 41t41 17t41-17t17-41q0-23-17-40t-41-17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SeventyOneHundredSixty(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M127 384h274q22 0 39.5-14.5T462 333l58-286q2-9-3-17t-14-9q-9-2-17 3t-9 14l-12 67q-5-3-15-13q-24-28-58-28t-58 28q-15 15-27 15q-11 0-28-15q-24-28-58-28q-35 0-57 28q-16 15-28 15t-28-15Q84 66 55 64l-4-26q-2-8-10-13t-16-4t-13.5 9.5T8 47l58 286q4 22 21.5 36.5T127 384M78 122q25 27 58 27t58-27q14-15 27-15t28 15q25 27 58 27q35 0 57-27q15-15 28-15t28 15q19 19 36 25l-36 177q-3 17-22 17H127q-18 0-21-17L63 109q3 2 7.5 6.5T78 122m207 70q0 9-6 15t-15 6t-15-6t-6-15t6-15t15-6t15 6t6 15m107 0q0 9-6.5 15t-14.5 6q-9 0-15.5-6t-6.5-15t6.5-15t15.5-6q8 0 14.5 6t6.5 15m-213 0q0 9-6.5 15t-15.5 6q-8 0-14.5-6t-6.5-15t6.5-15t14.5-6q9 0 15.5 6t6.5 15m64 85q0 9-6.5 15.5T221 299q-8 0-14.5-6.5T200 277q0-8 6.5-14.5T221 256q9 0 15.5 6.5T243 277m85 0q0 9-6.5 15.5T307 299q-9 0-15.5-6.5T285 277q0-8 6.5-14.5T307 256q8 0 14.5 6.5T328 277"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sharethis(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M164 232v5l165 82q22-19 52-19q34 0 57.5 23.5T462 381t-23.5 57.5T381 462t-57.5-23.5T300 381v-5l-165-82q-22 19-52 19q-34 0-57.5-23.5T2 232t23.5-57.5T83 151q30 0 52 19l165-82v-5q0-34 23.5-57.5T381 2t57.5 23.5T462 83t-23.5 57.5T381 164q-30 0-52-19l-165 82z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shield(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M320 0H64Q37 0 18.5 18.5T0 64v151q0 65 26.5 124T102 439l90 79l90-79q102-91 102-224V64q0-27-18.5-45.5T320 0m21 215q0 119-87 192l-62 56V43h128q21 0 21 21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shoe(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M392 195q-74 0-145-152q-2-3-3.5-6t-3-7t-2.5-6Q225 3 200 3H93Q76 3 63.5 15.5T51 45q0 86-17 116q-46 56-15 162H8v64q0 17 12.5 29.5T51 429h426q18 0 30.5-12.5T520 387v-64q0-19-6.5-39.5t-20-41t-40-34T392 195m21 44q27 6 43 26.5t18.5 34T477 323h-64zM68 186q25-31 25-141h107q1 3 4.5 9t4.5 8q0 3 4 7q-1 0-7 4l-21 21q-14 16 0 30q6 7 15 7t15-7l17-17q2 5 8.5 15.5T249 137l-21 21q-16 16 0 30q6 7 15 7q8 0 15-7l17-17q15 19 34 34l-17 17q-16 16 0 30q6 7 15 7q8 0 15-7l21-21q2-2 2-4q10 4 26 8v88H63v-7q-5-15-8.5-36T53 230.5T68 186m409 201H51v-22h426zm-320-86q28 0 46-18t18-46q0-27-18-45.5T157 173q-27 0-45.5 18.5T93 237q0 28 18.5 46t45.5 18m0-85q10 0 16 6t6 15q0 22-22 22q-9 0-15-6t-6-16q0-9 6-15t15-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingCart(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469 155H252L151 16q-6-10-17-11t-21 7q-15 16-2 32l79 111H43q-18 0-30.5 12.5T0 197v22q0 28 26 38l59 181q4 14 16 25.5t27 11.5h256q15 0 27-11.5t16-25.5l59-179q26-10 26-38v-22q0-19-12.5-31.5T469 155M43 197h179l15 20q2 0 2 2H43zm343 229v2q0 1-1 2t-1 2H130q-2 0-2-4L73 261h366zM275 219q7-12 7-22h187v22zM149 394q1 7 8 12t14 5h4q8-2 13-10t2-16l-21-85q-2-8-10-13t-16-2q-8 2-13 10t-2 15zm188 17h4q7 0 14-5t8-12l21-86q2-8-2.5-16t-12.5-9q-21-4-26 15l-21 85q-6 20 15 28m-81 0q21 0 21-22v-85q0-21-21-21t-21 21v85q0 22 21 22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sign(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469 171h-36L316 87q4-16 4-23q0-27-18.5-45.5T256 0t-45.5 18.5T192 64q0 7 4 23L79 171H43q-18 0-30.5 12.5T0 213v256q0 18 12.5 30.5T43 512h426q18 0 30.5-12.5T512 469V213q0-17-12.5-29.5T469 171M256 43q21 0 21 21t-21 21t-21-21t21-21m-32 76q16 9 32 9t32-9l73 52H151zM43 469V213h426v256zm362-192H107q-22 0-22 22q0 9 6 15t16 6h298q10 0 16-6t6-15q0-22-22-22m-64 86H171q-22 0-22 21t22 21h170q22 0 22-21t-22-21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Simplenote(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M65 394Q4 326 0 243v-19q1-36 17-74q17-47 50-75q0 27 1 36q9 45 37 73q14 13 54 35q20 11 59 30.5t59 30.5q51 27 60 58q13 49-30 92q-27 28-64 34h-31q-92-10-147-70m152-232q19 10 56.5 29.5T329 221q34 17 51 35q41 46 37 112q4-4 16-23q23-42 27-92q0-1 1-2v-36q-12-93-71.5-148.5T238 3h-12q-44 7-60 26q-16 18-18 42q-2 25 9 43.5t23.5 27T217 162"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SixtyOneHundredForty(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M164 92q-16 15-28 15t-28-15Q84 66 55 64l-4-26q-2-8-10-13t-16-4t-13.5 9.5T8 47l58 286q4 22 21.5 36.5T127 384h274q22 0 39.5-14.5T462 333l58-286q2-9-3-17t-14-9q-9-2-17 3t-9 14l-12 67q-5-3-15-13q-24-28-58-28t-58 28q-15 15-27 15q-11 0-28-15q-24-28-58-28q-35 0-57 28m143 57q35 0 57-27q15-15 28-15t28 15q19 19 36 25l-36 177q-3 17-22 17H127q-18 0-21-17L63 109q3 2 7.5 6.5T78 122q25 27 58 27t58-27q14-15 27-15t28 15q25 27 58 27m21 43q0 9-6.5 15t-14.5 6q-9 0-15.5-6t-6.5-15t6.5-15t15.5-6q8 0 14.5 6t6.5 15m-85 0q0 9-6.5 15t-15.5 6q-8 0-14.5-6t-6.5-15t6.5-15t14.5-6q9 0 15.5 6t6.5 15m85 85q0 9-6.5 15.5T307 299q-9 0-15.5-6.5T285 277q0-8 6.5-14.5T307 256q8 0 14.5 6.5T328 277m-85 0q0 9-6.5 15.5T221 299q-8 0-14.5-6.5T200 277q0-8 6.5-14.5T221 256q9 0 15.5 6.5T243 277"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Skype(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M441 267q3-15 3-33q0-88-63-151T230 20q-16 0-33 3q-32-19-68-19q-53 0-91 37.5T0 133q0 36 19 68q-2 11-2 33q0 89 62.5 151.5T230 448q18 0 34-3q30 19 68 19q53 0 91-37.5t38-91.5q0-37-20-68m-109 64q-13 18-40 31q-27 11-62 11q-43 0-70-15q-20-11-32-29q-13-17-13-36q0-9 8-17q7-7 19-7q9 0 16 5q7 6 12 17q4 10 11 21q7 8 18 14q12 5 30 5q25 0 41-11q16-10 16-27q0-11-8-21q-10-8-23-12t-36-9q-33-8-52-16q-21-9-34-25q-12-16-12-39t13-40q14-18 38-26q23-10 58-10q26 0 46 6q20 8 32 17q13 10 19 21q6 12 6 22t-8 18q-7 8-19 8q-11 0-16-5q-6-6-11-15q-8-16-19-23q-9-8-34-8q-22 0-36 9q-13 9-13 21q0 8 4 13q4 6 13 10q7 4 16 6q5 2 28 7q27 7 44 12q18 6 34 15q14 11 22 24q7 13 7 34q0 25-13 44"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Slashdot(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 459L209 2h88L93 459zm374-69q0-31-21.5-52T306 317q-31 0-52 21t-21 52q0 30 21 51t52 21q30 0 51.5-21t21.5-51"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Slideshare(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M455 142q-6 12-17 22t-20 15t-27 13.5t-26 12.5v73q0 17-15.5 43.5T301 348t-50-14t-17-27v-67q-8-2-14-6v73q0 10-13 19.5t-31 9.5q-19 0-46-20.5T103 278v-79q-66-26-81-45q-22-24-20-33q0-2 1-4q2-5 16 3t34.5 20T86 154q51 12 101 10q23-1 45 19l2 2v-9q0-14 44-14h73q8 0 43.5-14t46.5-25q13-15 20-8q5 6-6 27m-153-4q21 0 36-15t15-36t-15-36t-36-15t-36 15t-15 36t15 36t36 15m-134 0q21 0 36-15t15-36t-15-36t-36-15t-36 15t-15 36t15 36t36 15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Smugmug(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M188 36q0 14-13 24.5T143 71q-18 0-31-10.5T99 36t13-24t31-10q19 0 32 10t13 24M344 2q-19 0-32 10t-13 24t13 24.5T344 71q18 0 31-10.5T388 36t-13-24t-31-10m118 175q0 13-9 43t-30.5 72t-51.5 79.5t-78 64T190 462q-69 0-114.5-30.5t-60.5-76t-12-96T27 165q5-10 10.5-16.5T46 141l3-1h381q4 0 9 1.5t14 11t9 24.5m-46 9H67q-19 38-21 84t20 88q34 58 124 58q42 0 79.5-20t61.5-47.5t43.5-62T403 228t13-42"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sound(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 195v42q0 35 26.5 60.5T95 323h43q14 0 21 10l43 66q18 30 57 30h43q18 0 31.5-12.5T347 387V45q0-17-13-29.5T302 3h-43q-39 0-57 30l-43 66q-7 10-21 10H95q-37 0-63.5 25.5T5 195m43 0q0-18 13.5-30.5T95 152h43q36 0 57-30l43-66q7-11 21-11h45v342h-45q-14 0-21-11l-43-66q-21-30-57-30H95q-20 0-33.5-12.5T48 237z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SoundDown(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M90 323h42q14 0 22 10l42 66q20 30 58 30h43q19 0 31.5-12.5T341 387v-86q0-21-21-21t-21 21v86h-45q-14 0-21-11l-43-66q-17-30-58-30H90q-20 0-33.5-12.5T43 237v-42q0-18 13.5-30.5T90 152h42q37 0 58-30l43-66q7-11 21-11h45v86q0 21 21 21t21-21V45q0-17-12.5-29.5T297 3h-43q-38 0-58 30l-42 66q-8 10-22 10H90q-37 0-63.5 25.5T0 195v42q0 35 26.5 60.5T90 323m166-107q0 21 21 21h214q21 0 21-21t-21-21H277q-21 0-21 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SoundLevelOne(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 237q0 35 26.5 60.5T95 323h43q14 0 21 10l43 66q18 30 57 30h43q18 0 31.5-12.5T347 387V45q0-17-13-29.5T302 3h-43q-39 0-57 30l-43 66q-7 10-21 10H95q-37 0-63.5 25.5T5 195zm43-42q0-18 13.5-30.5T95 152h43q36 0 57-30l43-66q7-11 21-11h45v342h-45q-14 0-21-11l-43-66q-16-30-57-30H95q-20 0-33.5-12.5T48 237zm320-43v128q0 21 21 21q22 0 22-21V152q0-21-22-21q-21 0-21 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SoundLevelTwo(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5 237q0 35 26.5 60.5T95 323h43q14 0 21 10l43 66q18 30 57 30h43q18 0 31.5-12.5T347 387V45q0-17-13-29.5T302 3h-43q-39 0-57 30l-43 66q-7 10-21 10H95q-37 0-63.5 25.5T5 195zm43-42q0-18 13.5-30.5T95 152h43q36 0 57-30l43-66q7-11 21-11h45v342h-45q-14 0-21-11l-43-66q-16-30-57-30H95q-20 0-33.5-12.5T48 237zm363 85V152q0-21-22-21q-21 0-21 21v128q0 21 21 21q22 0 22-21m42-213q-21 0-21 21v256q0 21 21 21q22 0 22-21V88q0-21-22-21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SoundPlus(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M90 323h42q14 0 22 10l42 66q20 30 58 30h43q19 0 31.5-12.5T341 387v-86q0-21-21-21t-21 21v86h-45q-14 0-21-11l-43-66q-17-30-58-30H90q-20 0-33.5-12.5T43 237v-42q0-18 13.5-30.5T90 152h42q37 0 58-30l43-66q7-11 21-11h45v86q0 21 21 21t21-21V45q0-17-12.5-29.5T297 3h-43q-38 0-58 30l-42 66q-8 10-22 10H90q-37 0-63.5 25.5T0 195v42q0 35 26.5 60.5T90 323m401-128h-86v-86q0-21-21-21t-21 21v86h-86q-21 0-21 21t21 21h86v86q0 21 21 21t21-21v-86h86q21 0 21-21t-21-21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Soundcloud(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 238q0 24 18 39v-78Q2 215 2 238m37-49v98q5 2 15 2h5V187h-5q-5 0-15 2m43 6q-2 0-4-2v96h19V161q-12 15-15 34m34-52v146h20V133q-9 3-20 10m39-14v160h19V130q-4-1-12-1zm49 11q-11-5-20-8v157h29V129q-3 3-9 11m19-21v170h183q56-3 56-56q0-23-15.5-39.5T408 177q-10 0-22 5q-4-37-32-62t-66-25t-65 24"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpadesCard(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M64 512h256q27 0 45.5-18.5T384 448V64q0-27-18.5-45.5T320 0H64Q37 0 18.5 18.5T0 64v384q0 27 18.5 45.5T64 512M43 64q0-21 21-21h256q21 0 21 21v384q0 21-21 21H64q-21 0-21-21zm128 207v49h-22v21h86v-21h-22v-49q9 6 32.5 3t31.5-18q11-20-10.5-47T218 166l-26-17q-2 1-6 3.5t-15 10t-20 14.5t-21 18t-17 20t-9 21t3 20q8 15 31.5 18t32.5-3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spotify(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M88 408q59-85 162-85q92 0 151 71q67-67 67-162T400.5 69.5T237 2Q142 2 74.5 69.5T7 232q0 105 81 176m273-107q-5 11-17 6q-47-22-101-22q-55 0-103 23q-10 6-14-7q-3-11 5-17q53-25 112-25q62 0 111 25q12 5 7 17m19-59q-3 6-10 6q-3 0-4-1q-60-26-123-26q-64 0-122 26l-1 1h-3q-5 0-11-6l-2-8q0-6 5-9q63-29 134-29q72 0 135 29q5 3 5 9zM92 149q72-29 151-29t151 29q10 5 10 17q0 8-5.5 13.5T385 185q-3 0-4-1q-66-26-138-26q-73 0-137 26q-1 1-4 1q-8 0-14-5.5T82 166q0-10 10-17m158 250q-60 0-96 48q39 15 83 15q54 0 102-23q-35-40-89-40"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Squarespace(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M78 174L216 35q26-26 63-26t63 26q15 16-1 32q-15 16-31 0l-2-3q-13-11-30.5-10.5T248 67L106 209q-16 16-32 0t0-32q1-1 4-3m311 4q-16-16-32 0l-3 3l-138 139q-13 13-31.5 13T153 320q-16-16-32 0q-6 7-6 16t6 15q1 3 4 3q26 24 61.5 23.5T248 352l141-143q6-7 6-16t-6-15m-217 91l-3 3q-16 16 0 32q14 16 32 0l141-142q13-14 31.5-14t31.5 14q13 13 13 31.5T405 225L271 361q13 13 31 13t32-13l102-104q26-26 26-63.5T436 130t-63-26t-62 26zm-19-12l141-143q7-6 7-15.5T294 83q-6-7-15.5-7T263 83l-3 3l-138 139q-13 13-31.5 13T59 225t-13-31.5T59 162L193 26q-13-13-31-13t-32 13L28 130Q2 156 2 193.5T28 257t63 26t62-26"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Squidoo(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M393 368q-5 15-17.5 23t-25.5 4t-19-17.5t-1-28.5q4-15 16.5-23t25.5-4t19.5 17.5T393 368m-74 94q-31-4-53.5-27T237 379q-1-8-2-22t-3.5-20t-10.5-9q-10-8-51-1t-57 5q-27-4-43.5-23.5T52 257q0-12 5.5-40.5T60 170t-21-26q32 7 39 23q9 19 6.5 39.5t-1 43T97 280q11 7 56 2q7-1 24-4.5t27-3.5q-7-9-22-17.5t-17-9.5q-30-22-43-47q-7-13-15.5-55.5T85 90Q60 69 6 80q16-13 48-15t47 9q14 10 24 50t21 53q10 11 40 23q25 10 49 12q2-22 2-34.5t-10-24t-18.5-17.5t-22-15t-21.5-16Q123 68 128 2q9 50 31 68q7 6 24.5 12.5T206 92q66 39 78 104q6-4 21-25q16-23 16-31q1-21-24-44q-8-6-32-14.5t-33.5-22T227 18q1 3 4 14t6 17t8 10q9 7 33 13t33 12q46 28 45 66q0 14-14 41t-6 50q9 3 22.5 6t20 5t14.5 6t15 10q26 23 31.5 58.5t-7.5 67t-43.5 52T319 462m99-86q10-35-4-66.5T370 269q-31-9-60.5 9.5T270 332t4 66.5t45 40.5q31 10 59.5-9t39.5-54"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sreenshot(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469 213h-25q-13-54-52-93t-93-52V43q0-18-13-30.5T256 0t-30 12.5T213 43v25q-54 13-93 52t-52 93H43q-18 0-30.5 13T0 256t12.5 30T43 299h25q13 54 52 93t93 52v25q0 18 13 30.5t30 12.5t30-12.5t13-30.5v-25q54-13 93-52t52-93h25q18 0 30.5-13t12.5-30t-12.5-30t-30.5-13M299 399v-15q0-17-13-30t-30-13t-30 13t-13 30v15q-37-11-63-37t-37-63h15q17 0 30-13t13-30t-13-30t-30-13h-15q11-37 37-63t63-37v15q0 17 13 30t30 13t30-13t13-30v-15q75 23 100 100h-15q-17 0-30 13t-13 30t13 30t30 13h15q-11 37-37 63t-63 37"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stats(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21 512h86q9 0 15-6t6-15V341q0-17-12.5-29.5T85 299H43q-18 0-30.5 12.5T0 341v150q0 9 6 15t15 6m22-171h42v43H43zm0 86h42v42H43zm170 85h86q9 0 15-6t6-15V192q0-17-12.5-30T277 149h-42q-18 0-30.5 13T192 192v299q0 9 6 15t15 6m22-320h42v64h-42zm0 107h42v170h-42zM469 0h-42q-18 0-30.5 12.5T384 43v448q0 9 6 15t15 6h86q9 0 15-6t6-15V43q0-18-12.5-30.5T469 0m0 469h-42v-85h42zm0-128h-42V43h42z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Steam(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M428 135q0 18-12.5 30.5T385 178q-17 0-29.5-12.5T343 135q0-17 12.5-29.5T385 93q18 0 30.5 12.5T428 135m-17 72l-52 44q1 4 1 11q0 28-20 48t-48 20q-24 0-42.5-15T226 277L98 234q-17 12-36 12q-25 0-42.5-17.5T2 186t17.5-42.5T62 126q21 0 37 13.5t21 33.5l123 41q20-19 47-20l19-56v-3q0-32 22.5-54.5T386 58q31 0 53.5 22.5T462 135q0 25-14 44.5T411 207M99 166q-12-22-36-22q-18 0-30 12.5T21 186t12 29.5T63 228q4 0 12-2l-20-7q-12-6-16-19t2-25q7-12 20-16t25 2zm227-31q0 25 17 42.5t42 17.5t42.5-17.5T445 135t-17.5-42T385 76t-42 17t-17 42m1 133q0-19-13.5-32T281 223q-7 0-10 1l26 8q12 7 16 20t-3 25t-20 16t-25-3l-27-9q10 32 43 32q19 0 32.5-13t13.5-32"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stethoscope(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 107q-27 0-45.5 18T384 171q0 20 12.5 36.5T427 230v101q0 57-40.5 97.5T288 469t-98.5-40.5T149 331v-56q46-8 76.5-43.5T256 149V43q0-18-12.5-30.5T213 0h-21q-21 0-21 21q0 22 21 22h21v106q0 35-25 60.5T128 235t-60-25.5T43 149V43h21q21 0 21-22Q85 0 64 0H43Q25 0 12.5 12.5T0 43v106q0 47 30.5 82.5T107 275v56q0 75 53 128t128 53t128-53t53-128V230q19-6 31-22.5t12-36.5q0-28-18.5-46T448 107m0 85q-21 0-21-21q0-22 21-22t21 22q0 21-21 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Store(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M512 192q0-4-2-9L446 13q-4-13-19-13H85Q70 0 66 13L2 183q-2 5-2 9q0 34 21 51v226q-9 0-15 6t-6 16q0 9 6 15t15 6h470q9 0 15-6t6-15q0-10-6-16t-15-6V243q21-17 21-51M171 469h-43V320h43zm277 0H213V299q0-22-21-22h-85q-22 0-22 22v170H64V254q13-3 23-13q18 15 41 15q26 0 43-17q15 17 42 17q26 0 43-17q17 17 43 17q27 0 42-17q17 17 43 17q23 0 41-15q10 10 23 13zm11-256q-11 0-11-21t-21-21q-22 0-22 21t-21 21t-21-21t-22-21q-21 0-21 21t-21 21q-22 0-22-21t-21-21t-21 21t-22 21q-21 0-21-21t-21-21q-22 0-22 21t-21 21t-21-21t-22-21q-21 0-21 21t-11 21q-7 0-10-17l57-153h312l57 153q-3 17-10 17m-54 64H256q-21 0-21 22v106q0 22 21 22h149q22 0 22-22V299q0-22-22-22m-21 107H277v-64h107z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stumbleupon(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M462 334q0 32-23 55.5T383 413h-12q-33 0-56-23.5T292 334v-98q0-12-8-20.5t-19-8.5H158q-11 0-19.5 8t-8.5 19t8.5 19t20.5 9l16 1q32 1 54.5 24t21.5 53q-2 30-25.5 51.5T170 413H2v-51h168q12 0 20.5-7t8.5-17q1-9-7.5-16t-19.5-8l-16-1q-32-1-55-24.5T78 234q0-32 23.5-54.5T157 157h107q33 0 56 23t23 56v98q0 11 8 19.5t20 8.5h12q12 0 20-8.5t8-19.5V3h46q1 0 2.5.5t2.5.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Suitcase(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 107h-85V64q0-27-18.5-45.5T299 0h-86q-27 0-45.5 18.5T149 64v43H64q-27 0-45.5 18T0 171v277q0 27 18.5 45.5T64 512h384q27 0 45.5-18.5T512 448V171q0-28-18.5-46T448 107M192 64q0-21 21-21h86q21 0 21 21v43H192zM64 469q-21 0-21-21V171q0-22 21-22h21v320zm64 0V149h256v320zm341-21q0 21-21 21h-21V149h21q21 0 21 22z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sun(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 107q-62 0-105.5 43.5T107 256t43.5 105.5T256 405t105.5-43.5T405 256t-43.5-105.5T256 107m0 256q-45 0-76-31t-31-76t31-76t76-31t76 31t31 76t-31 76t-76 31m21-299V21q0-21-21-21t-21 21v43q0 21 21 21t21-21M85 256q0-21-21-21H21q-21 0-21 21t21 21h43q21 0 21-21m406-21h-43q-21 0-21 21t21 21h43q21 0 21-21t-21-21M256 427q-21 0-21 21v43q0 21 21 21t21-21v-43q0-21-21-21m164-305l22-22q14-14 0-30q-16-14-30 0l-22 22q-14 14 0 30q8 6 15 6q9 0 15-6m-313 6q7 0 15-6q14-16 0-30l-22-22q-14-14-30 0q-14 16 0 30l22 22q6 6 15 6m313 262q-14-14-30 0q-14 16 0 30l22 22q6 6 15 6q7 0 15-6q14-16 0-30zm-328 0l-22 22q-14 14 0 30q8 6 15 6q9 0 15-6l22-22q14-14 0-30q-16-14-30 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Switch(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M66 136h297q9 0 15-6t6-15q0-10-6-16t-15-6H66l34-51q7-6 6-15.5T98 12Q82 0 68 17L0 115l68 98q17 15 30 2q16-16 2-30zm348 47q-17 15-2 30l34 51H149q-9 0-15 6t-6 15q0 10 6 16t15 6h297l-34 51q-7 6-6 15.5t8 14.5q6 6 15.5 5t14.5-7l68-101l-68-98q-14-16-30-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tacos(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M299 149h21v22h-21zm-107 0h21v22h-21zm85 43h22v21h-22zm64 0h22v21h-22zm-192 0h22v21h-22zm64 0h22v21h-22zm288 4q7-24-6-49q-13-21-38-30l-30-6q-3-25-24-43q-25-21-55-10l-32 6l2 6q-32-6-62-6q-36 0-60 9v-9l-32-6q-26-7-55 10q-21 18-24 43l-27 6q-26 7-40 29.5T11 194l6 30Q0 270 0 320q0 17 12.5 30T43 363h426q18 0 30.5-13t12.5-30q0-53-19-94zM43 320q0-51 21-90q26-55 77.5-89T256 107q62 0 114 33.5t78 87.5q21 45 21 90H43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tag(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M162 483q18 15 42 15q32 0 51-24l163-194q26-32 29-75l11-143q0-24-15-36l-17-13q-19-14-38-8L247 41q-42 9-68 43L17 278q-19 23-15 49q3 28 23 45zM49 303l162-194q15-19 47-30l140-36l17 15l-10 143q-3 29-22 51L221 446q-6 8-15.5 8.5T189 449L53 333q-9-9-9-15q0-7 5-15m311-183q0 9-6.5 15t-14.5 6q-9 0-15.5-6t-6.5-15t6.5-15t15.5-6q8 0 14.5 6t6.5 15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Target(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0Q150 0 75 75T0 256t75 181t181 75t181-75t75-181t-75-181T256 0m0 469q-88 0-150.5-62.5T43 256t62.5-150.5T256 43t150.5 62.5T469 256t-62.5 150.5T256 469m0-362q-62 0-105.5 43.5T107 256t43.5 105.5T256 405t105.5-43.5T405 256t-43.5-105.5T256 107m0 256q-45 0-76-31t-31-76t31-76t76-31t76 31t31 76t-31 76t-76 31m43-107q0 18-12.5 30.5T256 299t-30.5-12.5T213 256t12.5-30.5T256 213t30.5 12.5T299 256"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Technorati(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M77 8q12-2 36-2h29q-40 51-44 95q-5 40 13 78.5t51 65.5q4 2 6 6t2 7.5t-.5 9.5t-.5 8q-7 65-8 79q13-8 45.5-30.5T258 292q117 22 204-33v91q1 26-13.5 49.5T410 433q-17 9-48 9H225l-137-1q-34 0-60-26T2 354V101Q0 69 22.5 41.5T77 8m120 59q-26 28-26 60q-3 46 39 78q37 28 87.5 30.5T389 215q55-31 55-84q2-48-42-79q-42-31-100-29q-64 2-105 44"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThirtyEighty(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M51 38q-2-8-10-13t-16-4t-13.5 9.5T8 47l58 286q4 22 21.5 36.5T127 384h274q22 0 39.5-14.5T462 333l58-286q2-9-3-17t-14-9q-9-2-17 3t-9 14l-12 67q-5-3-15-13q-24-28-58-28t-58 28q-15 15-27 15q-11 0-28-15q-24-28-58-28q-35 0-57 28q-16 15-28 15t-28-15Q84 66 55 64zm27 84q25 27 58 27t58-27q14-15 27-15t28 15q25 27 58 27q35 0 57-27q15-15 28-15t28 15q19 19 36 25l-36 177q-3 17-22 17H127q-18 0-21-17L63 109q3 2 7.5 6.5T78 122m207 113q0 8-6 14.5t-15 6.5t-15-6.5t-6-14.5q0-9 6-15.5t15-6.5t15 6.5t6 15.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Threewords(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M181 187L303 46V2H22v110h57V54h128l24-2v2q-11 7-18 16L101 203l15 33h10q10-1 17-1q44 0 73 22.5t29 61.5q0 37-27.5 61.5T150 405q-29 0-57-12.5T52 368l-14-13l-34 46q6 7 18 17t51 27t81 17q70 0 114-42.5T312 317t-39.5-92.5T181 187"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ticket(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M48 512h45q14 0 25.5-9t14.5-23q4-14 15.5-23t25.5-9t25.5 9t14.5 23q12 32 41 32h49q17 0 30-12.5t13-30.5V43q0-18-13-30.5T304 0h-45q-14 0-25.5 9T219 32q-12 32-43 32q-14 0-25.5-9T135 32Q125 0 93 0H48Q31 0 18 12.5T5 43v426q0 18 13 30.5T48 512m0-469l45 2q7 27 30.5 44.5T176 107t52.5-18T259 43h45v426l-45-2q-7-27-30.5-44.5T176 405t-52.5 18T93 469H48zm192 341q17 0 30-12.5t13-30.5V171q0-18-13-30.5T240 128H112q-17 0-30 12.5T69 171v170q0 18 13 30.5t30 12.5zM112 171h128v170H112z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Token(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0Q150 0 75 75T0 256q0 100 67.5 173T235 510v2h42v-2q93-8 159.5-74T510 277h2v-42h-2q-8-100-81-167.5T256 0m21 45q62 6 113 47l-44 45q-29-22-69-28zm-42 0v64q-35 3-69 28l-44-45q47-41 113-47M92 122l45 44q-25 34-28 69H45q3-63 47-113M45 277h64q6 33 25 64l-51 39q-35-50-38-103m190 190q-68-6-124-55l51-39q30 24 71 30v64zm21-104q-45 0-76-31t-31-76t31-76t76-31t76 31t31 76t-31 76t-76 31m21 104v-64q58-6 92-51l51 38q-56 71-143 77m169-113l-51-38q8-26 10-39h64q-6 41-23 77m-43-119q-6-40-28-69l45-44q41 51 47 113zm-104 21q0 18-12.5 30.5T256 299t-30.5-12.5T213 256t12.5-30.5T256 213t30.5 12.5T299 256"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Triangle(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M473 373q7-11 0-21L259 11q-6-10-18-10t-18 10L10 352q-9 10 0 21q3 11 17 11h426q12 0 20-11M65 341L240 62l175 279z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tribe(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M419 189q-4 0-12 2l-24-45q7-7 7-20q0-14-9-23t-23-9q-17 0-25 13L218 70v-5q0-13-9.5-22.5T186 33t-22.5 9.5T154 65v8l-72 31Q69 83 45 83q-18 0-30.5 12.5T2 126t12.5 30.5T45 169q9 0 21-5l71 87q-7 16-7 30q0 29 20.5 49.5T200 351q25 0 44-16.5t24-40.5l113-41q14 22 38 22q18 0 30.5-12.5T462 232t-12.5-30.5T419 189M87 133q0-1 .5-3t.5-3l74-40q9 11 24 11q11 0 22-9l19 6q-33 2-56.5 26.5T147 180q0 8 4 24zm72 92v-1h1zm96 14q-21-27-55-27q-8 0-20 2q-10-15-10-34q0-26 18.5-44t44.5-18t44.5 18.5T296 180q0 20-11.5 36T255 239m124-21l-104 37q43-25 43-75q0-40-30-65l38 13q1 13 10 21.5t22 8.5q5 0 8-1l24 43q-8 8-11 18"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tripit(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M286 212q-3-6-9-16l15-9l76-45q-1-3-1-8q0-12 8-20t20-8q11 0 19.5 8t8.5 20q0 11-8.5 19.5T395 162q-7 0-17-6q-8 5-18.5 11.5t-28 17T301 203zm12 64q-6 2-8 4l-3-2q-6-2-13-7q-8 8-13 11l12 7q6 4 8 5v5q0 4 1 7q2 9 10.5 14t18.5 2q9-4 13-11q3-4 3-11q0-5-1-7q-2-10-10.5-14.5T298 276m-23-47q-2-20-20-32t-41-10q-24 3-38.5 19T163 242q2 19 20 31t41 10q23-3 38-19t13-35m153-26q-27 0-34 26q-15 0-50.5-.5T289 228v7q0 3-2 11q20 0 55.5-.5t50.5-.5q7 28 35 28q14 0 24-10.5t10-24.5t-10-24.5t-24-10.5m-264 72q-1-1-3-4t-3-4q-17 10-36 23q-8-8-19-8q-12 2-19 9t-7 19q1 11 9 18.5t19 6.5q10 0 17.5-7.5T130 309v-1q0-4-1-6q2 0 40-22zM64 155q8 5 16 3q6-2 9-4l27 21l42 32q1-1 11-14q-45-35-72-51q0-1 .5-2.5t.5-2.5l-1-4q-1-9-8.5-14T72 116t-13 9q-4 5-4 12v4q3 7 9 14m151-26v44q2-1 8-1t8 1q0-32-1-44q34-7 34-40q0-16-12-28t-29-12t-29 12t-12 28q0 15 9.5 26t23.5 14m-65 105v-6q-17 0-46.5.5T61 229q-5-24-29-24q-13 0-21.5 9T2 235q0 13 8.5 21.5T32 265q21 0 28-20q58 0 90 1q-2-6 0-12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trophy(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469 107h-23v-7q21-10 21-36V43q0-18-12.5-30.5T425 0H85Q68 0 55.5 12.5T43 43v21q0 26 21 36v7H43q-18 0-30.5 12.5T0 149q0 53 37.5 90.5T128 277h41q11 10 23 15v94q0 54-21 75q-9 8-22 8h-21q-21 0-21 22q0 21 21 21h258q21 0 21-21q0-22-21-22h-21q-17 0-24-8q-21-19-21-75v-94q16-7 23-15h41q53 0 90.5-37.5T512 149q0-17-12.5-29.5T469 107M85 43h342v21H85zM43 149h34q13 40 47 86q-34-1-57.5-26.5T43 149m172 320q11-19 15-42h49q8 32 13 42zm62-85h-42v-77q6 2 21 2t21-2zm22-128h-86q-41-20-73.5-64.5T107 107h298q0 39-33 84t-73 65m89-21q30-36 47-86h34q0 34-23.5 59.5T388 235"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Truck(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M405 85h-42V43q0-18-13-30.5T320 0H43Q25 0 12.5 12.5T0 43v256q0 17 12.5 29.5T43 341h4q6 19 22.5 31t37.5 12q20 0 36.5-12t22.5-31h180q6 19 22.5 31t36.5 12q21 0 37.5-12t22.5-31h4q18 0 30.5-12.5T512 299v-77q0-17-13-30zM107 341q-8 0-15-6.5T85 320t7-14.5t15-6.5t14.5 6.5T128 320t-6.5 14.5T107 341m213-42H166q-6-19-22.5-31T107 256q-21 0-37.5 12T47 299h-4v-86h277zm0-214v86H43V43h277zm85 256q-8 0-14.5-6.5T384 320t6.5-14.5T405 299t15 6.5t7 14.5t-7 14.5t-15 6.5m64-42h-4q-6-19-22.5-31T405 256q-27 0-42 17V128h12l94 94z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TruckOne(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m480 147l-60-15l-55-89q-15-22-36-22H43q-17 0-30 13T0 64v256h26q6 17 23 30t36 13q21 0 37.5-12t22.5-31h198q15 43 60 43q20 0 36.5-13t23.5-30h6q18 0 30.5-12.5T512 277v-89q0-14-9-26t-23-15M85 320q-9 0-15-6t-6-15q0-10 6-16t15-6q22 0 22 22q0 9-6 15t-16 6m320 0q-9 0-15-6t-6-15q0-10 6-16t15-6q22 0 22 22q0 9-6 15t-16 6m64-43h-4q-6-17-23-29.5T405 235t-36.5 11.5T346 277H145q-6-18-23-30t-37-12q-27 0-42 17V64h286l12 21h-21q-21 0-21 22v42q0 22 21 22h83q1 0 3 1t4 1l59 15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TumbleDry(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M469 0H43Q25 0 12.5 12.5T0 43v426q0 18 12.5 30.5T43 512h426q18 0 30.5-12.5T512 469V43q0-18-12.5-30.5T469 0m0 469H43V43h426zm-213-23q79 0 135.5-56T448 254t-56-136t-136-56t-136 56t-56 136t56 136t136 56m0-339q61 0 105 43.5T405 256t-43.5 105.5T256 405t-105.5-43.5T107 256t43.5-105.5T256 107"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tumblr(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M192 378q-12-7-17-20q-3-7-3-49V190h108v-82H172V2h-66q-5 37-16 58q-12 24-31 40q-14 12-53 25v65h63v162q0 33 7 48q6 17 24 32t42 22q21 8 50 8q32 0 52-5q28-7 54-20v-72q-34 22-70 22q-20 0-36-9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Twitter(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M290 347H176q-24 0-40-16q-17-19-17-41v-40h162q22 0 38-16t16-38t-16-37q-15-16-38-16H119V59q0-24-16-40Q86 2 62 2Q37 2 21 18Q4 35 4 59v231q0 71 51 122q49 50 121 50h114q24 0 41-17t17-40q0-24-17-41q-19-17-41-17"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TwitterAlt(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M410 174q40-3 50-29q-8 5-24 7.5t-30-1.5q-2-8-2-11q-10-35-38-57.5T306 64q3-2 11-4l9-2q10-3 15.5-7t3.5-8q-1-2-5.5-2T329 42.5T317.5 46l-11 4l-6.5 2q23-8 25-19q-21 3-36 17q6-8 7-13q-40 25-73 111q-22-22-34-28q-46-26-119-52q-2 36 40 58q-8-1-29 3q7 36 52 46q-20 1-32 13q18 32 57 27q-25 12-16.5 27.5T173 255q-37 37-88 37T2 259q42 57 107 78.5t125.5 9.5T344 295.5t63-95.5q36 1 55-20q-31 4-52-6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ufo(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m516 211l6-10l-2-7q0-4-8.5-18.5T471 136t-81-40q-8-41-44-68.5T264 0t-82 27.5T138 96Q35 128 9 193l-1 1l-2 9l4 8l2 2H8q0 58 41 103q66 68 215 68t215-70q41-43 41-103zM264 43q33 0 57.5 19.5T349 111l-4 17h2q-9 43-83 43t-83-43h2l-4-17q3-29 27.5-48.5T264 43m-126 98q2 12 13 28l-30 29q-14 16 0 30q6 7 15 7t15-7l32-32q27 14 58 17v43q0 21 21 21t21-21v-45q34-3 58-17l32 32q7 7 15 7q6 0 15-7q14-14 0-30l-30-30q8-11 13-27q61 22 85 57q-17 28-68 54.5T264 277q-61 0-109.5-16.5t-68-31.5T55 198q24-36 83-57m126 200q-131 0-183-55q-7-7-15-19q85 53 198 53t198-53q-8 12-15 19q-52 55-183 55"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Up(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m44 273l156-139l156 137q4 4 15 4q10 0 17-6q13-15-2-30L200 79L14 241q-14 16-2 30q14 13 32 2"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UpArrowCircle(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0Q150 0 75 75T0 256t75 181t181 75t181-75t75-181t-75-181T256 0m0 469q-88 0-150.5-62.5T43 256t62.5-150.5T256 43t150.5 62.5T469 256t-62.5 150.5T256 469m13-315l-2-3h-22q-2 0-2 3l-64 42q-14 12-4 30t30 7l30-22v130q0 22 21 22t21-22V211l30 22q20 11 30-7q10-19-6-30z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Upload(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 384q21 0 21-21V66l94 77q7 6 16 5t14-7q6-7 5-16t-7-14L256 0L115 111q-15 15-2 30q16 14 30 2l92-77v297q0 21 21 21m171 21q0 9-6.5 15.5T405 427q-8 0-14.5-6.5T384 405q0-8 6.5-14.5T405 384q9 0 15.5 6.5T427 405m-64 0q0 9-6.5 15.5T341 427q-8 0-14.5-6.5T320 405q0-8 6.5-14.5T341 384q9 0 15.5 6.5T363 405m128-106H320v42h149v128H43V341h149v-42H21q-8 0-14.5 6.5T0 320v149q0 18 12.5 30.5T43 512h426q18 0 30.5-12.5T512 469V320q0-8-6.5-14.5T491 299"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UploadToCloud(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M427 171v-11q0-65-47.5-112.5T267 0q-57 0-101 35.5T111 126q-48 6-79.5 42.5T0 254q0 53 38 91.5t90 38.5h299q37-8 61-38t24-69q0-38-24-68.5T427 171m0 170H128q-35 0-60-26t-25-61q0-31 25.5-54.5T128 171q4 0 11-11l10-32q8-39 41.5-62T267 43q48 0 82.5 34.5T384 160v32q0 9 6.5 15t14.5 6h13q22 5 36.5 23t14.5 41t-11 41.5t-31 22.5M269 111l-2-2h-22q-2 0-2 2l-64 43q-8 5-9.5 13t3.5 16q10 18 30 7l30-21v108q0 22 21 22t21-22V169l30 21q20 11 30-7q11-19-6-29z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func User(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 448v64h320v-64q0-52-30-97.5T216 288v-23q29-15 46.5-43.5T280 160v-43q0-48-34.5-82.5T163 0T80 34.5T45 117v43q0 33 17.5 61.5T109 265v23q-46 17-76 62T3 448m277 0v21H45v-21q0-36 19.5-67.5T116 333q14 30 47 30q31 0 47-30q31 16 50.5 48t19.5 67M88 117q0-30 22-52t53-22q30 0 52 22t22 52v43q0 31-22 53t-52 22q-31 0-53-22t-22-53zm75 160h10v32q0 11-10 11q-11 0-11-11v-32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vcard(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M445 346q0-7-3-12t-5.5-6.5T426 323t-11-4l-77-31h-2q-1-1-2-1q-20-9-30-21v-1l-1-1q-3-5-6-4h-2l-1-2l-6-18h-1q-2-2-2-4l1-14l4-4q3-3 5-13t3-11t3-3.5t4.5-6.5t4.5-12q5-16 6.5-27t-1.5-15q-3 0-5 2q3-16 4-31q0-46-10-62q-7-12-25-23q-8-6-17-9q-13-4-37-3.5T187 7q-24 5-38 28q-11 18-11 66q0 3 2 11.5t3 11.5q1 4 2 6h-2q-3 4-1.5 15t6.5 27q2 8 4.5 12t4.5 6.5t3 3.5t3 11t5 13l4 4l1 12l1 7q-2-4-3-1t-3 9.5t-3 10.5l-1-1q-3-3-8 4q-6 12-31 22v2q-1 1-4 1l-76 31q-2 1-11 4t-11.5 4.5T17 334t-3 12L2 413h460z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Viddler(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M351 221q46 0 78.5-32.5T462 110t-32.5-78T351 0q-42 0-73 27t-37 67h-3q-6-40-37.5-67T128 0Q81 0 48.5 32T16 110t32 78t78 33h3v37L2 224v147l127-38v51h188l-55-163zm-223-80q-14 0-23-9t-9-22t9.5-22t22.5-9t22 9t9 22t-9 22t-22 9m194-31q0-13 9-22t22-9t22.5 9t9.5 22t-9 22t-23 9q-13 0-22-9t-9-22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VideoCamera(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M45 177q0 2-1 7.5t-1 7.5v149q0 18 12.5 30.5T85 384h171q17 0 30-12.5t13-30.5v-8l55 28q2 2 9 2q8 0 10-2q11-8 11-20V192q0-13-11-19q-11-7-21 0l-53 28v-9q0-2-1-7.5t-1-7.5q44-28 44-81q0-40-28-68T245 0q-45 0-74 36Q142 0 96 0Q56 0 28 28T0 96q0 52 45 81m296 49v81l-42-21v-39zm-85-34v149H85V192zM245 43q23 0 38.5 15.5T299 96t-15.5 37.5T245 149t-38-15.5T192 96t15-37.5T245 43M96 43q22 0 37.5 15.5T149 96t-15.5 37.5T96 149t-37.5-15.5T43 96t15.5-37.5T96 43m53 213q0 9-6 15t-15 6t-15-6t-6-15t6-15t15-6t15 6t6 15m86 0q0 9-6.5 15t-15.5 6q-8 0-14.5-6t-6.5-15t6.5-15t14.5-6q9 0 15.5 6t6.5 15m21-160q0 11-11 11q-10 0-10-11t10-11q11 0 11 11m-149 0q0 11-11 11T85 96t11-11t11 11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vimeo(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M395 0q-96-3-130 108q17-7 34-7q35 0 31 39q-1 22-31 70q-29 46-43 46q-19 0-34-71q-5-18-19-106q-12-79-66-73q-22 1-68 41q-11 10-33.5 30T2 107l21 28q33-22 36-22q24 0 45 75q6 23 18.5 68.5T141 325q29 75 68 75q66 0 159-123q91-116 94-184q5-90-67-93"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Virb(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M217 333H116L2 51h97l68 184l67-184h95zm245-233q0-20-13.5-34T415 52t-34 14t-14 34t14 33.5t34 13.5q19 0 33-13.5t14-33.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wthree(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m172 20l61 209l62-209h149v22l-54 108q37 12 53 38q19 30 19 66q0 48-23 78q-26 32-64 32q-27 0-51-19q-22-20-32-51l35-15q8 22 20 31q11 11 29 11q19 0 30-20q12-19 12-47q0-38-13-54q-19-26-48-26h-29v-8l72-107h-76l-86 287h-5l-63-213l-63 213h-4L2 20h44l61 209l41-140l-20-69z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wallet(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M405 0H85Q50 0 25 25T0 85v342q0 35 25 60t60 25h342q35 0 60-25t25-60V107q0-45-31-76T405 0M43 85q0-17 12.5-29.5T85 43h320q21 0 37.5 11.5T465 85H247l-70 58q-23 17-23 49t23 49l70 58h218q-6 19-22.5 30.5T405 341H85q-20 0-42 11zm42 299h299v21H49q11-21 36-21m384 43q0 17-12.5 29.5T427 469H85q-26 0-36-21h378v-66q22-6 42-19zm0-171H265l-60-47q-9-5-9-17t9-17l60-47h204zm-170-64q0 9-6.5 15t-15.5 6q-8 0-14.5-6t-6.5-15t6.5-15t14.5-6q9 0 15.5 6t6.5 15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wand(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M399 15Q386 2 369.5 2T340 15L37 316q-13 13-13 30t13 30l30 32q13 13 29.5 13t29.5-13l303-303q13-13 13-30t-13-30zm-91 152l-30-30l89-92l30 30zM67 3Q45 3 45 24v21H24Q3 45 3 67q0 21 21 21h21v21q0 22 22 22q9 0 15-6t6-16V88h21q10 0 16-6t6-15q0-22-22-22H88V24Q88 3 67 3m298 256q-21 0-21 21v21h-21q-22 0-22 22q0 9 6 15t16 6h21v21q0 10 6 16t15 6q22 0 22-22v-21h21q21 0 21-21q0-22-21-22h-21v-21q0-21-22-21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Warning(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 402q19 27 53 27h288q32 0 53-27q21-32 5-67l-75-162l-71-132Q247 7 208 7q-40 0-55 34L82 173L5 333q-13 39 5 69m36-47l75-160l70-133q6-12 19-12t19 12l69 133l76 162q9 14-2 23q-9 9-17 9H63q-8 0-17-9q-6-16 0-25m162-75q21 0 21-21v-86q0-21-21-21t-21 21v86q0 21 21 21m21 43q0 8-6 14.5t-15 6.5t-15-6.5t-6-14.5q0-9 6-15.5t15-6.5t15 6.5t6 15.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Watch(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3 256q0 39 20 73.5T75 384h-8v85q0 18 12.5 30.5T109 512h86q17 0 29.5-12.5T237 469v-85h-8q61-37 70-107h2q22 0 22-21t-22-21h-2q-9-68-70-107h8V43q0-18-12.5-30.5T195 0h-86Q92 0 79.5 12.5T67 43v85h8q-32 20-52 54.5T3 256m106 213v-21h86v21zm86-42h-86v-28q20 6 43 6q24 0 43-6zm64-171q0 45-31 76t-76 31t-76-31t-31-76t31-76t76-31t76 31t31 76M109 43h86v21h-86zm0 42h86v28q-19-6-43-6q-23 0-43 6zm64 107h-42v73l49 49l30-30l-37-37z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WaterTemperatureFifty(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M181 222q17 0 28 11q10 10 10 27q0 20-13 32q-10 11-36 11q-18 0-32-6v-22q2 1 7 3t8 4q6 2 17 2q24 0 24-19q0-20-24-20h-11q-2 0-4.5 1t-3.5 1l-11-6l5-62h66v22h-43l-2 23h2q4-2 13-2m138 17q0 29-10 47q-10 15-32 15q-24 0-32-15q-11-17-11-47q0-33 11-47q10-15 32-15q23 0 32 17q10 22 10 45m-59 0q0 20 4 32q4 11 13 11q7 0 13-11q4-12 4-32t-4-32q-6-11-13-11q-9 0-13 11q-4 12-4 32m72-34q0-8 5-15q4-8 10-11q8-4 15-4t15 4q7 4 11 11q4 8 4 15q0 5-4 15q-3 6-11 10q-7 5-15 5q-13 0-21-9q-9-8-9-21m17 0q0 5 5 8q5 5 8 5q4 0 9-5q4-4 4-8q0-5-4-9t-9-4q-4 0-8 4q-5 3-5 9M127 384h274q22 0 39.5-14.5T462 333l58-286q2-9-3-17t-14-9q-9-2-17 3t-9 14l-12 67q-5-3-15-13q-24-28-58-28t-58 28q-15 15-27 15q-11 0-28-15q-24-28-58-28q-35 0-57 28q-16 15-28 15t-28-15Q84 66 55 64l-4-26q-2-8-10-13t-16-4t-13.5 9.5T8 47l58 286q4 22 21.5 36.5T127 384M78 122q25 27 58 27t58-27q14-15 27-15t28 15q25 27 58 27q35 0 57-27q15-15 28-15t28 15q19 19 36 25l-36 177q-3 17-22 17H127q-18 0-21-17L63 109q3 2 7.5 6.5T78 122"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WaterTemperatureForty(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M226 275h-15v26h-26v-26h-51v-17l53-79h24v77h15zm-41-19v-47q-1 2-3 6.5t-3 6.5l-22 34zm134-17q0 29-10 47q-10 15-32 15q-24 0-32-15q-11-17-11-47q0-33 11-47q10-15 32-15q23 0 32 17q10 22 10 45m-59 0q0 20 4 32q4 11 13 11q7 0 13-11q4-12 4-32t-4-32q-6-11-13-11q-9 0-13 11q-4 12-4 32m72-34q0-8 5-15q4-8 10-11q8-4 15-4t15 4q7 4 11 11q4 8 4 15q0 5-4 15q-3 6-11 10q-7 5-15 5q-13 0-21-9q-9-8-9-21m17 0q0 5 5 8q5 5 8 5q4 0 9-5q4-4 4-8q0-5-4-9t-9-4q-4 0-8 4q-5 3-5 9M127 384h274q22 0 39.5-14.5T462 333l58-286q2-9-3-17t-14-9q-9-2-17 3t-9 14l-12 67q-5-3-15-13q-24-28-58-28t-58 28q-15 15-27 15q-11 0-28-15q-24-28-58-28q-35 0-57 28q-16 15-28 15t-28-15Q84 66 55 64l-4-26q-2-8-10-13t-16-4t-13.5 9.5T8 47l58 286q4 22 21.5 36.5T127 384M78 122q25 27 58 27t58-27q14-15 27-15t28 15q25 27 58 27q35 0 57-27q15-15 28-15t28 15q19 19 36 25l-36 177q-3 17-22 17H127q-18 0-21-17L63 109q3 2 7.5 6.5T78 122"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WaterTemperatureNinetyFive(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M221 230q0 39-15 54q-17 17-47 17q-12 0-17-2v-22q5 2 15 2q8 0 22-4q9-3 12-13q5-10 5-21q-2 2-5.5 6t-5.5 5q-4 2-15 2q-16 0-25-11q-9-8-9-23q0-19 11-32q7-11 30-11q14 0 23 6q13 11 15 20q6 24 6 27m-44-32q-6 0-13 7q-5 4-5 15q0 2 5 15q2 4 13 4q8 0 12-4q7-7 7-13q0-10-5-17q-3-7-14-7m102 24q17 0 28 11q10 10 10 27q0 20-12 32q-11 11-37 11q-18 0-32-6v-22q5 2 15 7q7 2 17 2q24 0 24-19q0-20-24-20h-10q-3 0-5.5 1t-3.5 1l-11-6l5-62h66v22h-43l-2 23h2q4-2 13-2m53-17q0-8 5-15q4-8 10-11q8-4 15-4t15 4q7 4 11 11q4 8 4 15q0 5-4 15q-3 6-11 10q-7 5-15 5q-13 0-21-9q-9-8-9-21m17 0q0 5 5 8q5 5 8 5q4 0 9-5q4-4 4-8q0-5-4-9t-9-4q-4 0-8 4q-5 3-5 9M503 21q-9-2-17 3t-9 14l-12 67q-5-3-15-13q-24-28-58-28t-58 28q-15 15-27 15q-11 0-28-15q-24-28-58-28q-35 0-57 28q-16 15-28 15t-28-15Q84 66 55 64l-4-26q-2-8-10-13t-16-4t-13.5 9.5T8 47l58 286q4 22 21.5 36.5T127 384h274q22 0 39.5-14.5T462 333l58-286q2-8-3.5-16.5T503 21m-81 303q-3 17-21 17H127q-18 0-21-17L63 109q3 2 7.5 6.5T78 122q25 27 58 27t58-27q14-15 27-15t28 15q25 27 58 27q35 0 57-27q15-15 28-15t28 15q19 19 36 25z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WaterTemperatureSeventy(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m149 301l45-100h-60v-22h87v17l-47 105zm170-62q0 29-10 47q-10 15-32 15q-24 0-32-15q-11-17-11-47q0-33 11-47q10-15 32-15q23 0 32 17q10 22 10 45m-59 0q0 20 4 32q4 11 13 11q7 0 13-11q4-12 4-32t-4-32q-6-11-13-11q-9 0-13 11q-4 12-4 32m72-34q0-8 5-15q4-8 10-11q8-4 15-4t15 4q7 4 11 11q4 8 4 15q0 5-4 15q-3 6-11 10q-7 5-15 5q-13 0-21-9q-9-8-9-21m17 0q0 5 5 8q5 5 8 5q4 0 9-5q4-4 4-8q0-5-4-9t-9-4q-4 0-8 4q-5 3-5 9M127 384h274q22 0 39.5-14.5T462 333l58-286q2-9-3-17t-14-9q-9-2-17 3t-9 14l-12 67q-5-3-15-13q-24-28-58-28t-58 28q-15 15-27 15q-11 0-28-15q-24-28-58-28q-35 0-57 28q-16 15-28 15t-28-15Q84 66 55 64l-4-26q-2-8-10-13t-16-4t-13.5 9.5T8 47l58 286q4 22 21.5 36.5T127 384M78 122q25 27 58 27t58-27q14-15 27-15t28 15q25 27 58 27q35 0 57-27q15-15 28-15t28 15q19 19 36 25l-36 177q-3 17-22 17H127q-18 0-21-17L63 109q3 2 7.5 6.5T78 122"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WaterTemperatureSixty(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M136 250q0-39 15-54q19-17 45-17q13 0 17 2v22q-4-2-15-2q-13 0-21 4q-10 4-13 8q-5 7-5 22h3q10-15 25-15q14 0 26 10q8 9 8 28q0 20-10 30q-13 11-30 11q-14 0-24-7q-13-9-15-15q-6-24-6-27m45 32q8 0 13-7q4-8 4-15q0-5-4-15q-3-4-13-4q-9 0-13 4q-6 6-6 13q0 9 4 17q4 7 15 7m138-43q0 29-10 47q-10 15-32 15q-24 0-32-15q-11-17-11-47q0-33 11-47q10-15 32-15q23 0 32 17q10 22 10 45m-59 0q0 20 4 32q4 11 13 11q7 0 13-11q4-12 4-32t-4-32q-6-11-13-11q-9 0-13 11q-4 12-4 32m72-34q0-8 5-15q4-8 10-11q8-4 15-4t15 4q7 4 11 11q4 8 4 15q0 5-4 15q-3 6-11 10q-7 5-15 5q-13 0-21-9q-9-8-9-21m17 0q0 5 5 8q5 5 8 5q4 0 9-5q4-4 4-8q0-5-4-9t-9-4q-4 0-8 4q-5 3-5 9M127 384h274q22 0 39.5-14.5T462 333l58-286q2-9-3-17t-14-9q-9-2-17 3t-9 14l-12 67q-5-3-15-13q-24-28-58-28t-58 28q-15 15-27 15q-11 0-28-15q-24-28-58-28q-35 0-57 28q-16 15-28 15t-28-15Q84 66 55 64l-4-26q-2-8-10-13t-16-4t-13.5 9.5T8 47l58 286q4 22 21.5 36.5T127 384M78 122q25 27 58 27t58-27q14-15 27-15t28 15q25 27 58 27q35 0 57-27q15-15 28-15t28 15q19 19 36 25l-36 177q-3 17-22 17H127q-18 0-21-17L63 109q3 2 7.5 6.5T78 122"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WaterTemperatureThirty(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M217 205q0 13-6 19q-15 12-20 13q17 3 22 8q8 10 8 20q0 18-12 27q-14 11-37 11q-20 0-34-6v-22q1 0 7 2.5t8 4.5q4 2 17 2q12 0 19-5q7-3 7-14q0-10-7-13q-10-5-21-5h-13v-19h9q13 0 21-4q6-3 6-13q0-13-17-13q-8 0-12 3q-5 2-13 6l-13-19q16-13 41-13q21 0 29 8q11 11 11 22m102 34q0 29-10 47q-10 15-32 15q-24 0-32-15q-11-17-11-47q0-33 11-47q10-15 32-15q23 0 32 17q10 22 10 45m-59 0q0 20 4 32q4 11 13 11q7 0 13-11q4-12 4-32t-4-32q-6-11-13-11q-9 0-13 11q-4 12-4 32m72-34q0-8 5-15q4-8 10-11q8-4 15-4t15 4q7 4 11 11q4 8 4 15q0 5-4 15q-3 6-11 10q-7 5-15 5q-13 0-21-9q-9-8-9-21m17 0q0 5 5 8q5 5 8 5q4 0 9-5q4-4 4-8q0-5-4-9t-9-4q-4 0-8 4q-5 3-5 9M127 384h274q22 0 39.5-14.5T462 333l58-286q2-9-3-17t-14-9q-9-2-17 3t-9 14l-12 67q-5-3-15-13q-24-28-58-28t-58 28q-15 15-27 15q-11 0-28-15q-24-28-58-28q-35 0-57 28q-16 15-28 15t-28-15Q84 66 55 64l-4-26q-2-8-10-13t-16-4t-13.5 9.5T8 47l58 286q4 22 21.5 36.5T127 384M78 122q25 27 58 27t58-27q14-15 27-15t28 15q25 27 58 27q35 0 57-27q15-15 28-15t28 15q19 19 36 25l-36 177q-3 17-22 17H127q-18 0-21-17L63 109q3 2 7.5 6.5T78 122"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Whatsapp(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M462 228q0 93-66 159t-160 66q-56 0-109-28L2 464l40-120q-32-54-32-116q0-93 66-158.5T236 4t160 65.5T462 228M236 39q-79 0-134.5 55.5T46 228q0 62 36 111l-24 70l74-23q49 31 104 31q79 0 134.5-55.5T426 228T370.5 94.5T236 39m114 241q-1-1-10-7q-3-1-19-8.5t-19-8.5q-9-3-13 2q-1 3-4.5 7.5t-7.5 9t-5 5.5q-4 6-12 1q-34-17-45-27q-7-7-13.5-15t-12-15t-5.5-8q-3-7 3-11q4-6 8-10l6-9q2-5-1-10q-4-13-17-41q-3-9-12-9h-11q-9 0-15 7q-19 19-19 45q0 24 22 57l2 3q2 3 4.5 6.5t7 9t9 10.5t10.5 11.5t13 12.5t14.5 11.5t16.5 10t18 8.5q16 6 27.5 10t18 5t9.5 1t7-1t5-1q9-1 21.5-9t15.5-17q8-21 3-26"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wikipedia(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M462 33h-95l-7 1v10q0 2 4 2q1 1 4 1l9 1q11 0 17 6q5 7 1 19l-84 220l-3-1l-54-121l1-2l44-90q7-14 13-22q4-9 19-10q3 0 3-3V34q-53-1-76 0v12q1 1 2 1l4 1q12 0 15 5t-2 20l-34 74l-30-69v-1q-11-23-6-27q3-1 8-2l4-1q3 0 3-3V34h-82v10q0 2 6 4q10 1 13.5 5.5T174 82l8 18l32 68q3 9 9 23l-46 101l-3-1Q103 125 82 70q-3-9-3-13q0-9 14-9h14q7 0 7-4v-9l-4-1H5l-3 1v9q0 2 6 4q15 0 22 5q5 7 11 24q13 33 54 129.5T150 339q13 30 29-1q22-46 55-121q6 15 17 42t19.5 47.5T284 338q14 32 29 0L417 79q6-15 14-23q9-8 28-9q3 0 3-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Windows(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M96 21q40-16 77-18q13 0 35 4q16 4 49 22l-42 142q-41-27-82-27q-35 0-77 16zm65 341q26-90 43-144q-9-6-17-10q-32-17-62-17q-1 0-4 .5t-4 .5q-35 3-63 14q-8 2-12 4L2 350q39-16 74-16q41 0 85 28m246-104q-42 13-75 13q-55 0-93-28l-40 141q50 29 83 29q37 0 86-20zm55-191q-38 16-77 16q-51 0-91-28l-41 142q40 27 89 27q38 0 79-18v-1h1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wists(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 0h123q2 43 11 109.5t11 95.5q5-31 13-109t10-96h124q1 43 5.5 105t5.5 99q14-65 34-204h123q-7 53-25 204.5T406 436H260q0-3-9-66.5T237 269q-4 18-15.5 77.5T204 436H58Q19 193 2 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Woman(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M192 64q0 26-19 45t-45 19t-45-19t-19-45t19-45t45-19t45 19t19 45M64 256v235q0 9 6 15t15 6q10 0 16-6t6-15v-43h42v43q0 9 6 15t16 6q9 0 15-6t6-15V192q0-17-12.5-30T149 149h-42q-18 0-30.5 13T64 192zM43 149q-18 0-30.5 13T0 192v149q17 0 30-12.5T43 299zm213 43q0-17-12.5-30T213 149v150q0 17 13 29.5t30 12.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wordpress(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2 232q0 95 67.5 162.5T232 462t162.5-67.5T462 232T394.5 69.5T232 2T69.5 69.5T2 232m17 0q0-44 19-87l101 278q-54-26-87-77.5T19 232m213 213q-32 0-60-9l64-185l65 179q1 0 1 1t1 2q-38 12-71 12m172-217q16-40 16-76q0-15-1-22q26 49 26 102q0 57-29 106.5T339 416zm-46-65q18 32 18 58t-17 69l-21 71l-77-229l25-2q5 0 7-4.5t-.5-8.5t-8.5-4q-34 2-56 2l-56-2q-6 0-8.5 4t-.5 8.5t7 4.5q11 2 23 2l33 91l-47 140l-77-231l24-2q6 0 7.5-4.5t-1-8.5t-7.5-4q-35 2-57 2H54q29-44 76-70t102-26q81 0 144 56h-3q-15 0-25.5 11.5T337 113q0 16 21 50"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WordpressAlt(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M28 119q29-54 83.5-86T232 1q83 0 150 59q-18-3-36 7q-24 14-26.5 42.5T345 157q32 20 32 76q0 10-14 42t-28 59l-14 27l-54-184q-2-13-2-17q0-8 5-13q5-7 8-7h26v-21H165v21h5q3 0 13 10q2 2 19 45l20 67l-43 100l-48-200q2-14 5-16q3-6 8-6h1v-21zm27 35q-11-14-23-14H20Q2 178 2 231q0 70 39 127t102 84zm375-36q8 33-5 73q-27 88-99 252q62-27 99-83.5T462 234q0-64-32-116M237 312l-59 144q30 7 54 7t56-7z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WorkCase(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M448 91h-85V69q0-27-18.5-45.5T299 5h-86q-27 0-45.5 18.5T149 69v22H64q-27 0-45.5 18T0 155v85q0 21 21 21v171q0 17 13 30t30 13h384q17 0 30-13t13-30V261q21 0 21-21v-85q0-28-18.5-46T448 91M192 69q0-9 6-15t15-6h86q9 0 15 6t6 15v22H192zM64 432V261h128v64q0 18 12.5 30.5T235 368h42q18 0 30.5-12.5T320 325v-64h128v171zm171-107v-64h42v64zm234-106H43v-64q0-22 21-22h384q21 0 21 22z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func World(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M256 0Q150 0 75 75T0 256t75 181t181 75t181-75t75-181t-75-181T256 0M109 331q21 53 21 96q-40-28-63.5-73.5T43 256q0-103 81-166q18 14 33 36.5t14 42.5q-4 31-41 61q-19 20-25.5 47.5T109 331m360-75q0 88-62.5 150.5T256 469q-44 0-83-17q5-65-26-136q-11-34 13-54q50-50 53-89q6-54-51-109q39-21 94-21q26 0 49 6l-43 43l-15 15l62 62l150 21q10 35 10 66M331 130l-24-23l43-43q55 26 87 81zm-62 90l-17 25q-20 32-5 64l9 20l15 72q3 23 28 30q4 2 10 2q15 0 28-11l24-23l2-4q15-56 34-75q27-19 15-68q-7-27-37-39l-51-21h-4q-35 0-51 28m19 49l17-26q4-8 13-8l43 15q8 4 8 8q2 7 2 24q-27 20-49 91l-11 11l-12-68l-13-26q-6-9 2-21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Xing(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M277 462L174 282L334 2h108L282 282l103 180zM108 323l80-132l-60-105H26l60 105L6 323z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Yahoo(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M373 122q-11 3-64 46t-57 57q-1 7-.5 39.5T253 307q3 1 32 1t33 1l-1 20H210q-3 0-43 1t-51 0l4-19h16l27-2l16-6q3-3 3-35.5t-2-42.5q-2-9-51.5-69T62 82H2V54h205v2h1l-1 5v21h-62q16 24 47.5 64.5T229 195l83-75h-49l-7-29h180l-2 2h1l-13 19l-5 8h-34q-8 2-10 2m59 178h-16l-18-2v30l14 1l15 1zm30-145q-24 0-60-6l1 129l27 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YahooBuzz(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M160 462q-64 0-109-45T6 309V2h88v307q0 27 19.5 46.5T160 375t46.5-19.5T226 309q0-28-19.5-47T160 243h-44v-88h44q64 0 109 45t45 109q0 63-45 108t-109 45"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YahooMessenger(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M232 462q95 0 162.5-67.5T462 232T394.5 69.5T232 2T69.5 69.5T2 232t67.5 162.5T232 462m59-356q12 0 21 8.5t9 21.5t-9 21.5t-21 8.5q-13 0-21.5-8.5T261 136t8.5-21.5T291 106m-114 0q12 0 21 8.5t9 21.5t-9 21.5t-21 8.5q-13 0-21.5-8.5T147 136t8.5-21.5T177 106M70 202h324v60h-1q-5 62-51 105.5T232 411t-110-43.5T71 262h-1zm302 54H88v-37h284z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Yelp(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M62 23q16-13 43-18t43-2q27 4 29 26l1 160q0 9-4.5 16.5T163 214q-15 4-26-16L53 61q-1-2-4.5-6t-4-6t0-6t5-9T62 23M34 326l85-29q21-10 23-26q2-18-19-27l-89-35q-9-4-17.5 3T5 233q-2 62 0 75q2 9 11 15.5t18 2.5m147 15q1-9-3.5-16t-10.5-8q-14-2-30 16l-59 71q-6 7-3.5 20T85 440l62 22q9 3 20.5-4.5T180 441zm147-16l-74-21q-21-7-30 0q-10 8 2 27l47 86q5 8 18.5 7t18.5-9q28-41 33-57q4-11-1-21t-14-12m18-120q-3-8-16.5-27T307 152q-21-18-35 0l-48 56q-15 18-4 34q9 16 32 11l82-15q20-5 12-33"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Youtube(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M325 339h-25v-15q0-12 12-12h1q12 0 12 12zm-95-32q-12 0-12 10v71q0 10 12 10t12-10v-71q0-10-12-10m154-40v136q0 25-18.5 42T321 462H63q-26 0-44.5-17T0 403V267q0-24 18.5-41.5T63 208h258q26 0 44.5 17.5T384 267M80 418V275h32v-21l-85-1v21h26v144zm96-122h-27v95q-2 5-7 6.5t-8-5.5v-19l-1-77h-26v100q2 17 7 20q9 6 22.5.5T155 402v16h21zm85 88v-64q0-19-12-26t-30 7v-47h-27v163h22l2-11q21 18 33.5 10t11.5-32m84-9h-20v14q0 11-11 11h-4q-11 0-11-11v-29h46v-17q0-26-1-33q-3-16-21.5-20t-30.5 5q-8 7-11 15q-3 10-3 27v38q0 37 28 43q24 5 35-16q7-11 4-27M242 169q2 5 7 8q4 3 11 3q6 0 10-3t7-9v10h30V51h-24v99q0 10-10 10q-9 0-9-10V51h-25v86q0 18 1 22q0 4 2 10m-90-71q0-18 3-29q3-9 11-16q8-6 20-6q11 0 18 4q7 3 11 10q4 5 6 13q1 5 1 21v32q0 20-1 26q-1 7-6 15q-3 5-11 11q-8 3-16 3q-10 0-18-3q-7-3-11-9q-4-8-5-14q-2-10-2-25zm23 50q0 5 3.5 9t8.5 4q12 0 12-13V81q0-13-12-13t-12 13zm-82 34h28V85l33-83h-30l-18 62L88 2H58l35 83z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YoutubeAlt(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M462 102q0-29-20-49t-48-20H70q-28 0-48 20T2 102v185q0 29 20 49t48 20h324q28 0 48-20t20-49zM186 273V99l132 87z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Zerply(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 366h14q64 0 98 17q13 7 36 24.5t34 23.5q45 28 76 31q38 4 56.5-30.5T332 360q-31 26-88 11q-13-4-32-14t-32-18l-12-8q-39-24-70-23q0-10 7-19q5-6 14.5-16.5T133 257q58-68 88-102q24-28 74-86l4-4l5-5.5l2-4.5q1-7 2.5-20t2.5-19q0-4-2-5.5T305 9h-2q-63 14-123 11q-45-2-68-12q-2 0-10-5q-7-4-12 1q-4 4-5 6q-21 28-23 51q0 15 1 18q1 1 1 2h1q21 29 121 10Q66 233 6 303v1l-4 10v44q1 4 5 8"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomIn(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M414 354q-18-18-41-11l-32-32q43-53 43-119q0-80-56-136T192 0T56 56T0 192t56 136t136 56q70 0 119-43l32 32q-6 24 11 41l85 85q13 13 30 13q18 0 30-13q13-13 13-30t-13-30zm-222-13q-62 0-105.5-43.5T43 192T86.5 86.5T192 43t105.5 43.5T341 192t-43.5 105.5T192 341m64-170h-43v-43q0-21-21-21t-21 21v43h-43q-21 0-21 21t21 21h43v43q0 21 21 21t21-21v-43h43q21 0 21-21t-21-21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomOut(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M192 384q70 0 119-43l32 32q-6 24 11 41l85 85q13 13 30 13q18 0 30-13q13-13 13-30t-13-30l-85-85q-18-18-41-11l-32-32q43-53 43-119q0-80-56-136T192 0T56 56T0 192t56 136t136 56m0-341q62 0 105.5 43.5T341 192t-43.5 105.5T192 341T86.5 297.5T43 192T86.5 86.5T192 43m-64 170h128q21 0 21-21t-21-21H128q-21 0-21 21t21 21"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Zootool(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23 329q22 0 47-23q53 101 77 128.5t60 27.5q39 0 72-24t33-51q0-19-13-26q-33 18-59 18q-18 0-37-20.5T132 260q81-72 129-142q50-33 50-65q0-14-9.5-26.5T279 14q-11 0-25 23Q167 2 121 2Q97 2 79.5 20.5T62 64q0 27 20 34q19-13 36-13q27 0 96 22q-15 22-49 58t-63 59q-20-17-37-17q-29 0-45 26.5T4 288q0 41 19 41"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Zynga(children ...ElementRenderer) *PsIcon {
	return &PsIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M316 3q32-3 57 14q13-1 21 12q14 28 6 42q-1 1-2.5-1t-3.5-4l-1-3q-3 12-10 17q1 24-12 50q-2 5 1 10q18 55 17 84q-2 34-19 53q2 19-10 37q-9 16-14 45q-5 21-1 29q14 18 13 22q0 4-7 8.5t-12 4.5q-17 0-21-12q-6-16-2-31q4-12 9-46t0-57q-9 4-23.5 6.5T278 286h-10q-6 31-13 49q-2 5-2 53l6 47q1 1 2.5 2t3.5 4.5t2 7.5q-1 13-20 13q-8 0-14.5-5t-6.5-13v-55q0-30-3-42q-12-48-12-66l-7-2q-14 16-24 26q5 15 5 21q0 2 7 6q18 9 17 16q0 6-10.5 7t-21.5-4q-1-1-5-12t-7.5-23.5T160 302q-3-9-1-14q8-20 7-29q-1-7-10.5-15t-17.5-8q-3 0-7.5 2.5t-8 5t-8.5 6.5t-8 6q-5 4-11.5 10t-11 10t-5.5 5q-2 2-17 28.5T42 347l-2 5q-1 5-2.5 9t-4.5 8.5t-8.5 7T13 379q-11 0-10-10q2-9 20-22q3-7 7.5-16l7.5-15l3-7l2-10l4-20.5l6.5-28.5l8.5-31q15-49 36-65q4-3 1-5q-4-3-10-10t-13-27.5T72 66q1-10 4 3q0 3 1 7q2 15 12.5 34t28.5 26l8-3q9-3 38-8.5t72-8.5l3-35q-7-6-8-16q-4 7-9 6t-3-35q1-17 24-21l16-3q2-2 5.5-4t18-5.5T316 3"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
