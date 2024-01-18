package fontisto

import (
	. "github.com/delaneyj/gostar/elements"
)

const (
	IconifyVersion = "3.0.4"
	hAttr          = "1em"
	viewbox        = "0 0 24 24"
	fill           = "currentColor"
)

type FontistoIcon struct {
	*SVGSVGElement
}

type FontistoIconFn func(children ...ElementRenderer) *FontistoIcon

var IconLookup = map[string]FontistoIconFn{
	"acrobatReader":                    AcrobatReader,
	"adjust":                           Adjust,
	"adobe":                            Adobe,
	"aids":                             Aids,
	"airplay":                          Airplay,
	"algolia":                          Algolia,
	"amazon":                           Amazon,
	"ambulance":                        Ambulance,
	"americanExpress":                  AmericanExpress,
	"americanSignLanguageInterpreting": AmericanSignLanguageInterpreting,
	"ampproject":                       Ampproject,
	"anchor":                           Anchor,
	"android":                          Android,
	"angelist":                         Angelist,
	"angleDobuleDown":                  AngleDobuleDown,
	"angleDobuleLeft":                  AngleDobuleLeft,
	"angleDobuleRight":                 AngleDobuleRight,
	"angleDobuleUp":                    AngleDobuleUp,
	"angleDown":                        AngleDown,
	"angleLeft":                        AngleLeft,
	"angleRight":                       AngleRight,
	"angleUp":                          AngleUp,
	"angularjs":                        Angularjs,
	"appStore":                         AppStore,
	"apple":                            Apple,
	"appleMusic":                       AppleMusic,
	"applePay":                         ApplePay,
	"archive":                          Archive,
	"areaChart":                        AreaChart,
	"arrowDown":                        ArrowDown,
	"arrowDownL":                       ArrowDownL,
	"arrowExpand":                      ArrowExpand,
	"arrowH":                           ArrowH,
	"arrowLeft":                        ArrowLeft,
	"arrowLeftL":                       ArrowLeftL,
	"arrowMove":                        ArrowMove,
	"arrowResize":                      ArrowResize,
	"arrowReturnLeft":                  ArrowReturnLeft,
	"arrowReturnRight":                 ArrowReturnRight,
	"arrowRight":                       ArrowRight,
	"arrowRightL":                      ArrowRightL,
	"arrowSwap":                        ArrowSwap,
	"arrowUp":                          ArrowUp,
	"arrowUpL":                         ArrowUpL,
	"arrowV":                           ArrowV,
	"asterisk":                         Asterisk,
	"at":                               At,
	"atlassian":                        Atlassian,
	"atom":                             Atom,
	"audioDescription":                 AudioDescription,
	"automobile":                       Automobile,
	"aws":                              Aws,
	"babel":                            Babel,
	"backward":                         Backward,
	"baidu":                            Baidu,
	"ban":                              Ban,
	"bandage":                          Bandage,
	"barChart":                         BarChart,
	"batteryEmpty":                     BatteryEmpty,
	"batteryFull":                      BatteryFull,
	"batteryHalf":                      BatteryHalf,
	"batteryQuarter":                   BatteryQuarter,
	"batteryThreeQuarters":             BatteryThreeQuarters,
	"beachSlipper":                     BeachSlipper,
	"bedPatient":                       BedPatient,
	"behance":                          Behance,
	"bell":                             Bell,
	"bellAlt":                          BellAlt,
	"bicycle":                          Bicycle,
	"bing":                             Bing,
	"bitbucket":                        Bitbucket,
	"bitcoin":                          Bitcoin,
	"blind":                            Blind,
	"blogger":                          Blogger,
	"blood":                            Blood,
	"bloodDrop":                        BloodDrop,
	"bloodTest":                        BloodTest,
	"bluetoothB":                       BluetoothB,
	"bold":                             Bold,
	"bookmark":                         Bookmark,
	"bookmarkAlt":                      BookmarkAlt,
	"bower":                            Bower,
	"braille":                          Braille,
	"brokenLink":                       BrokenLink,
	"bug":                              Bug,
	"bus":                              Bus,
	"busTicket":                        BusTicket,
	"calculator":                       Calculator,
	"calendar":                         Calendar,
	"camera":                           Camera,
	"car":                              Car,
	"caretDown":                        CaretDown,
	"caretLeft":                        CaretLeft,
	"caretRight":                       CaretRight,
	"caretUp":                          CaretUp,
	"centerAlign":                      CenterAlign,
	"check":                            Check,
	"checkboxActive":                   CheckboxActive,
	"checkboxPassive":                  CheckboxPassive,
	"chrome":                           Chrome,
	"circleOnotch":                     CircleOnotch,
	"clock":                            Clock,
	"close":                            Close,
	"closeA":                           CloseA,
	"cloudDown":                        CloudDown,
	"cloudRefresh":                     CloudRefresh,
	"cloudUp":                          CloudUp,
	"cloudflare":                       Cloudflare,
	"cloudy":                           Cloudy,
	"cloudyGusts":                      CloudyGusts,
	"cocktail":                         Cocktail,
	"cocoapods":                        Cocoapods,
	"code":                             Code,
	"codepen":                          Codepen,
	"coffeescript":                     Coffeescript,
	"columns":                          Columns,
	"comment":                          Comment,
	"commenting":                       Commenting,
	"comments":                         Comments,
	"compass":                          Compass,
	"compassAlt":                       CompassAlt,
	"composer":                         Composer,
	"confused":                         Confused,
	"copy":                             Copy,
	"cpanel":                           Cpanel,
	"creditCard":                       CreditCard,
	"crop":                             Crop,
	"crosshairs":                       Crosshairs,
	"cssThree":                         CssThree,
	"cursor":                           Cursor,
	"curve":                            Curve,
	"dailymotion":                      Dailymotion,
	"database":                         Database,
	"date":                             Date,
	"dayCloudy":                        DayCloudy,
	"dayHaze":                          DayHaze,
	"dayLightning":                     DayLightning,
	"dayRain":                          DayRain,
	"daySnow":                          DaySnow,
	"daySunny":                         DaySunny,
	"deaf":                             Deaf,
	"delicious":                        Delicious,
	"deskpro":                          Deskpro,
	"desktop":                          Desktop,
	"deviantart":                       Deviantart,
	"digg":                             Digg,
	"dinnersClub":                      DinnersClub,
	"directionSign":                    DirectionSign,
	"discord":                          Discord,
	"discourse":                        Discourse,
	"discover":                         Discover,
	"dislike":                          Dislike,
	"disqus":                           Disqus,
	"dizzy":                            Dizzy,
	"dna":                              Dna,
	"doNotDisturb":                     DoNotDisturb,
	"docker":                           Docker,
	"doctor":                           Doctor,
	"dollar":                           Dollar,
	"download":                         Download,
	"dribbble":                         Dribbble,
	"dropbox":                          Dropbox,
	"drugPack":                         DrugPack,
	"earth":                            Earth,
	"edge":                             Edge,
	"eject":                            Eject,
	"electronjs":                       Electronjs,
	"ellipse":                          Ellipse,
	"email":                            Email,
	"ember":                            Ember,
	"envato":                           Envato,
	"equalizer":                        Equalizer,
	"eraser":                           Eraser,
	"euro":                             Euro,
	"export":                           Export,
	"expressionless":                   Expressionless,
	"eye":                              Eye,
	"facebook":                         Facebook,
	"famale":                           Famale,
	"favorite":                         Favorite,
	"fi":                               Fi,
	"file":                             File,
	"fileOne":                          FileOne,
	"film":                             Film,
	"filter":                           Filter,
	"fire":                             Fire,
	"firefox":                          Firefox,
	"firstAidAlt":                      FirstAidAlt,
	"fiveHundredPx":                    FiveHundredPx,
	"flag":                             Flag,
	"flash":                            Flash,
	"flickr":                           Flickr,
	"flipboard":                        Flipboard,
	"flotationRing":                    FlotationRing,
	"fog":                              Fog,
	"folder":                           Folder,
	"font":                             Font,
	"fontisto":                         Fontisto,
	"forward":                          Forward,
	"foursquare":                       Foursquare,
	"frowning":                         Frowning,
	"gbp":                              Gbp,
	"genderless":                       Genderless,
	"gg":                               Gg,
	"git":                              Git,
	"github":                           Github,
	"gitlab":                           Gitlab,
	"google":                           Google,
	"googleDrive":                      GoogleDrive,
	"googlePlay":                       GooglePlay,
	"googlePlus":                       GooglePlus,
	"googleWallet":                     GoogleWallet,
	"graphql":                          Graphql,
	"grunt":                            Grunt,
	"gulp":                             Gulp,
	"hackerNews":                       HackerNews,
	"hangout":                          Hangout,
	"hashtag":                          Hashtag,
	"headphone":                        Headphone,
	"heart":                            Heart,
	"heartAlt":                         HeartAlt,
	"heartEyes":                        HeartEyes,
	"heartbeat":                        Heartbeat,
	"heartbeatAlt":                     HeartbeatAlt,
	"helicopter":                       Helicopter,
	"helicopterAmbulance":              HelicopterAmbulance,
	"hexo":                             Hexo,
	"hipchat":                          Hipchat,
	"history":                          History,
	"holidayVillage":                   HolidayVillage,
	"home":                             Home,
	"horizon":                          Horizon,
	"horizonAlt":                       HorizonAlt,
	"hospital":                         Hospital,
	"hotAirBalloon":                    HotAirBalloon,
	"hotel":                            Hotel,
	"hotelAlt":                         HotelAlt,
	"hourglass":                        Hourglass,
	"hourglassEnd":                     HourglassEnd,
	"hourglassHalf":                    HourglassHalf,
	"hourglassStart":                   HourglassStart,
	"houzz":                            Houzz,
	"htmlFive":                         HtmlFive,
	"icq":                              Icq,
	"ifQuestionCircle":                 IfQuestionCircle,
	"ils":                              Ils,
	"imdb":                             Imdb,
	"import":                           Import,
	"indent":                           Indent,
	"info":                             Info,
	"injectionSyringe":                 InjectionSyringe,
	"inr":                              Inr,
	"instagram":                        Instagram,
	"internetExplorer":                 InternetExplorer,
	"intersex":                         Intersex,
	"invision":                         Invision,
	"island":                           Island,
	"italic":                           Italic,
	"iyzigo":                           Iyzigo,
	"java":                             Java,
	"jcb":                              Jcb,
	"jekyll":                           Jekyll,
	"jenkins":                          Jenkins,
	"jira":                             Jira,
	"joomla":                           Joomla,
	"jquery":                           Jquery,
	"jsfiddle":                         Jsfiddle,
	"json":                             Json,
	"justify":                          Justify,
	"key":                              Key,
	"keyboard":                         Keyboard,
	"kickstarter":                      Kickstarter,
	"krw":                              Krw,
	"laboratory":                       Laboratory,
	"language":                         Language,
	"laptop":                           Laptop,
	"laravel":                          Laravel,
	"laughing":                         Laughing,
	"leftAlign":                        LeftAlign,
	"less":                             Less,
	"lightbulb":                        Lightbulb,
	"lightning":                        Lightning,
	"lightnings":                       Lightnings,
	"like":                             Like,
	"line":                             Line,
	"lineChart":                        LineChart,
	"link":                             Link,
	"linkedin":                         Linkedin,
	"linux":                            Linux,
	"listOne":                          ListOne,
	"listTwo":                          ListTwo,
	"livestream":                       Livestream,
	"locked":                           Locked,
	"lowVision":                        LowVision,
	"mad":                              Mad,
	"magento":                          Magento,
	"magnet":                           Magnet,
	"male":                             Male,
	"map":                              Map,
	"mapMarker":                        MapMarker,
	"mapMarkerAlt":                     MapMarkerAlt,
	"mars":                             Mars,
	"marsDouble":                       MarsDouble,
	"marsStroke":                       MarsStroke,
	"marsStrokeH":                      MarsStrokeH,
	"marsStrokeV":                      MarsStrokeV,
	"mastercard":                       Mastercard,
	"maxcdn":                           Maxcdn,
	"medium":                           Medium,
	"meetup":                           Meetup,
	"mercury":                          Mercury,
	"messenger":                        Messenger,
	"meteor":                           Meteor,
	"metro":                            Metro,
	"mic":                              Mic,
	"microsoft":                        Microsoft,
	"minusA":                           MinusA,
	"mobile":                           Mobile,
	"mobileAlt":                        MobileAlt,
	"moneySymbol":                      MoneySymbol,
	"mongodb":                          Mongodb,
	"moreV":                            MoreV,
	"moreVa":                           MoreVa,
	"motorcycle":                       Motorcycle,
	"moveH":                            MoveH,
	"moveHa":                           MoveHa,
	"musicNote":                        MusicNote,
	"mysql":                            Mysql,
	"navIcon":                          NavIcon,
	"navIconA":                         NavIconA,
	"navIconGrid":                      NavIconGrid,
	"navIconGridA":                     NavIconGridA,
	"navIconList":                      NavIconList,
	"navIconListA":                     NavIconListA,
	"navigate":                         Navigate,
	"nervous":                          Nervous,
	"netflix":                          Netflix,
	"neuter":                           Neuter,
	"neutral":                          Neutral,
	"nginx":                            Nginx,
	"nightAltCloudy":                   NightAltCloudy,
	"nightAltLightning":                NightAltLightning,
	"nightAltRain":                     NightAltRain,
	"nightAltSnow":                     NightAltSnow,
	"nightClear":                       NightClear,
	"nodejs":                           Nodejs,
	"npm":                              Npm,
	"nurse":                            Nurse,
	"nursingHome":                      NursingHome,
	"odnoklassniki":                    Odnoklassniki,
	"onedrive":                         Onedrive,
	"onenote":                          Onenote,
	"openMouth":                        OpenMouth,
	"opencart":                         Opencart,
	"opera":                            Opera,
	"oracle":                           Oracle,
	"origin":                           Origin,
	"outdent":                          Outdent,
	"paperPlane":                       PaperPlane,
	"paperclip":                        Paperclip,
	"paragraph":                        Paragraph,
	"paralysisDisability":              ParalysisDisability,
	"parasol":                          Parasol,
	"passport":                         Passport,
	"passportAlt":                      PassportAlt,
	"paste":                            Paste,
	"pause":                            Pause,
	"paw":                              Paw,
	"paypal":                           Paypal,
	"paypalP":                          PaypalP,
	"payu":                             Payu,
	"periscope":                        Periscope,
	"person":                           Person,
	"persons":                          Persons,
	"phone":                            Phone,
	"photograph":                       Photograph,
	"php":                              Php,
	"picture":                          Picture,
	"pieChartOne":                      PieChartOne,
	"pieChartTwo":                      PieChartTwo,
	"pills":                            Pills,
	"pinboard":                         Pinboard,
	"pingdom":                          Pingdom,
	"pinterest":                        Pinterest,
	"plane":                            Plane,
	"planeTicket":                      PlaneTicket,
	"play":                             Play,
	"playList":                         PlayList,
	"playerSettings":                   PlayerSettings,
	"playstation":                      Playstation,
	"plusA":                            PlusA,
	"podcast":                          Podcast,
	"power":                            Power,
	"prescription":                     Prescription,
	"preview":                          Preview,
	"print":                            Print,
	"productHunt":                      ProductHunt,
	"propellerFour":                    PropellerFour,
	"propellerOne":                     PropellerOne,
	"propellerThree":                   PropellerThree,
	"propellerTwo":                     PropellerTwo,
	"pulse":                            Pulse,
	"python":                           Python,
	"qrcode":                           Qrcode,
	"question":                         Question,
	"quora":                            Quora,
	"quoteAleft":                       QuoteAleft,
	"quoteAright":                      QuoteAright,
	"quoteLeft":                        QuoteLeft,
	"quoteRight":                       QuoteRight,
	"radioBtnActive":                   RadioBtnActive,
	"radioBtnPassive":                  RadioBtnPassive,
	"rage":                             Rage,
	"rails":                            Rails,
	"rain":                             Rain,
	"rainbow":                          Rainbow,
	"rains":                            Rains,
	"random":                           Random,
	"raspberryPi":                      RaspberryPi,
	"react":                            React,
	"record":                           Record,
	"rectangle":                        Rectangle,
	"recycle":                          Recycle,
	"reddit":                           Reddit,
	"redis":                            Redis,
	"redo":                             Redo,
	"redux":                            Redux,
	"reply":                            Reply,
	"rightAlign":                       RightAlign,
	"rocket":                           Rocket,
	"room":                             Room,
	"rouble":                           Rouble,
	"rss":                              Rss,
	"ruby":                             Ruby,
	"safari":                           Safari,
	"saitBoat":                         SaitBoat,
	"sass":                             Sass,
	"saucelabs":                        Saucelabs,
	"save":                             Save,
	"saveOne":                          SaveOne,
	"scissors":                         Scissors,
	"scorp":                            Scorp,
	"search":                           Search,
	"sentry":                           Sentry,
	"share":                            Share,
	"shareA":                           ShareA,
	"shazam":                           Shazam,
	"shield":                           Shield,
	"ship":                             Ship,
	"shopify":                          Shopify,
	"shoppingBag":                      ShoppingBag,
	"shoppingBagOne":                   ShoppingBagOne,
	"shoppingBarcode":                  ShoppingBarcode,
	"shoppingBasket":                   ShoppingBasket,
	"shoppingBasketAdd":                ShoppingBasketAdd,
	"shoppingBasketRemove":             ShoppingBasketRemove,
	"shoppingPackage":                  ShoppingPackage,
	"shoppingPosMachine":               ShoppingPosMachine,
	"shoppingSale":                     ShoppingSale,
	"shoppingStore":                    ShoppingStore,
	"sinaWeibo":                        SinaWeibo,
	"sitemap":                          Sitemap,
	"skyatlas":                         Skyatlas,
	"skype":                            Skype,
	"slack":                            Slack,
	"slides":                           Slides,
	"slightlySmile":                    SlightlySmile,
	"smiley":                           Smiley,
	"smiling":                          Smiling,
	"snapchat":                         Snapchat,
	"snorkel":                          Snorkel,
	"snow":                             Snow,
	"snowflake":                        Snowflake,
	"snowflakeEight":                   SnowflakeEight,
	"snowflakeFive":                    SnowflakeFive,
	"snowflakeFour":                    SnowflakeFour,
	"snowflakeOne":                     SnowflakeOne,
	"snowflakeSeven":                   SnowflakeSeven,
	"snowflakeSix":                     SnowflakeSix,
	"snowflakeThree":                   SnowflakeThree,
	"snowflakeTwo":                     SnowflakeTwo,
	"snows":                            Snows,
	"soundcloud":                       Soundcloud,
	"sourcetree":                       Sourcetree,
	"spinner":                          Spinner,
	"spinnerCog":                       SpinnerCog,
	"spinnerFidget":                    SpinnerFidget,
	"spinnerRefresh":                   SpinnerRefresh,
	"spinnerRotateForward":             SpinnerRotateForward,
	"spotify":                          Spotify,
	"stackOverflow":                    StackOverflow,
	"star":                             Star,
	"starHalf":                         StarHalf,
	"steam":                            Steam,
	"stepBackwrad":                     StepBackwrad,
	"stepForward":                      StepForward,
	"stethoscope":                      Stethoscope,
	"stop":                             Stop,
	"stopwatch":                        Stopwatch,
	"strikethrough":                    Strikethrough,
	"stuckOutTongue":                   StuckOutTongue,
	"stumbleupon":                      Stumbleupon,
	"stylus":                           Stylus,
	"sublimeText":                      SublimeText,
	"subscript":                        Subscript,
	"subway":                           Subway,
	"suitcase":                         Suitcase,
	"suitcaseAlt":                      SuitcaseAlt,
	"sun":                              Sun,
	"sunglasses":                       Sunglasses,
	"sunglassesAlt":                    SunglassesAlt,
	"superscript":                      Superscript,
	"surgicalKnife":                    SurgicalKnife,
	"surprised":                        Surprised,
	"svn":                              Svn,
	"swarm":                            Swarm,
	"swift":                            Swift,
	"swimsuit":                         Swimsuit,
	"tableOne":                         TableOne,
	"tableTwo":                         TableTwo,
	"tablet":                           Tablet,
	"tabletAlt":                        TabletAlt,
	"tablets":                          Tablets,
	"taxi":                             Taxi,
	"ted":                              Ted,
	"telegram":                         Telegram,
	"tent":                             Tent,
	"tesla":                            Tesla,
	"testBottle":                       TestBottle,
	"testTube":                         TestTube,
	"testTubeAlt":                      TestTubeAlt,
	"textHeight":                       TextHeight,
	"textWidth":                        TextWidth,
	"thermometer":                      Thermometer,
	"thermometerAlt":                   ThermometerAlt,
	"ticket":                           Ticket,
	"ticketAlt":                        TicketAlt,
	"tinder":                           Tinder,
	"tl":                               Tl,
	"toggleOff":                        ToggleOff,
	"toggleOn":                         ToggleOn,
	"tongue":                           Tongue,
	"train":                            Train,
	"trainTicket":                      TrainTicket,
	"transgender":                      Transgender,
	"transgenderAlt":                   TransgenderAlt,
	"trash":                            Trash,
	"travis":                           Travis,
	"treehouse":                        Treehouse,
	"trello":                           Trello,
	"tripadvisor":                      Tripadvisor,
	"troy":                             Troy,
	"truck":                            Truck,
	"tty":                              Tty,
	"tumblr":                           Tumblr,
	"tv":                               Tv,
	"twitch":                           Twitch,
	"twitter":                          Twitter,
	"twoo":                             Twoo,
	"uber":                             Uber,
	"ubuntu":                           Ubuntu,
	"udacity":                          Udacity,
	"umbrella":                         Umbrella,
	"underline":                        Underline,
	"undo":                             Undo,
	"unity":                            Unity,
	"universalAcces":                   UniversalAcces,
	"unlocked":                         Unlocked,
	"unrealEngine":                     UnrealEngine,
	"upload":                           Upload,
	"usb":                              Usb,
	"userSecret":                       UserSecret,
	"venus":                            Venus,
	"venusDouble":                      VenusDouble,
	"venusMars":                        VenusMars,
	"viber":                            Viber,
	"vimeo":                            Vimeo,
	"vine":                             Vine,
	"visa":                             Visa,
	"visualStudio":                     VisualStudio,
	"vk":                               Vk,
	"volumeDown":                       VolumeDown,
	"volumeMute":                       VolumeMute,
	"volumeOff":                        VolumeOff,
	"volumeUp":                         VolumeUp,
	"vuejs":                            Vuejs,
	"wallet":                           Wallet,
	"webpack":                          Webpack,
	"webstorm":                         Webstorm,
	"wetransfer":                       Wetransfer,
	"whatsapp":                         Whatsapp,
	"wheelchair":                       Wheelchair,
	"wifi":                             Wifi,
	"wifiLogo":                         WifiLogo,
	"wikipedia":                        Wikipedia,
	"wind":                             Wind,
	"windows":                          Windows,
	"wink":                             Wink,
	"wix":                              Wix,
	"wordpress":                        Wordpress,
	"world":                            World,
	"worldO":                           WorldO,
	"xbox":                             Xbox,
	"yacht":                            Yacht,
	"yahoo":                            Yahoo,
	"yandex":                           Yandex,
	"yandexInternational":              YandexInternational,
	"yarn":                             Yarn,
	"yelp":                             Yelp,
	"yen":                              Yen,
	"youtubePlay":                      YoutubePlay,
	"zipperMouth":                      ZipperMouth,
	"zoom":                             Zoom,
	"zoomMinus":                        ZoomMinus,
	"zoomPlus":                         ZoomPlus,
}

func AcrobatReader(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23.678 15.404c-.712-.763-2.171-1.204-4.238-1.204a20.77 20.77 0 0 0-3.904.393l.131-.022a20.425 20.425 0 0 1-2.273-2.667l-.042-.062a22.788 22.788 0 0 1-1.363-2.176l-.061-.121c.663-1.807 1.098-3.895 1.208-6.069l.002-.051c0-1.678-.602-3.427-2.348-3.427a1.64 1.64 0 0 0-1.35.795L9.436.8c-.786 1.413-.431 4.508.92 7.565A68.022 68.022 0 0 1 8.65 12.98a56.821 56.821 0 0 1-1.929 4.114C2.802 18.687.267 20.544.03 21.998a1.609 1.609 0 0 0 .458 1.454c.392.34.907.547 1.47.547h.016h-.001c2.599 0 5.338-4.294 6.729-6.883c1.069-.361 2.137-.689 3.204-1.018a35.857 35.857 0 0 1 3.411-.768c2.744 2.508 5.163 2.91 6.379 2.91c1.497 0 2.031-.619 2.207-1.126a1.727 1.727 0 0 0-.255-1.749l.003.003l.02.04zm-1.39 1.058a1.28 1.28 0 0 1-1.397.909l.007.001a3.708 3.708 0 0 1-.626-.076l.024.004a9.321 9.321 0 0 1-3.921-2.12l.008.007c.895-.16 1.925-.251 2.976-.251h.015h-.001a7.996 7.996 0 0 1 1.867.153l-.051-.009c.492.106 1.274.436 1.099 1.385h.02zm-7.55-1.713c-.923.191-1.92.415-2.954.695a39.49 39.49 0 0 0-2.528.773a40.2 40.2 0 0 0 1.212-2.609c.429-1.023.783-2.077 1.139-3.056c.351.612.742 1.234 1.134 1.786a27.867 27.867 0 0 0 2.005 2.448l-.018-.02v-.02zm-4.665-13.53a.778.778 0 0 1 .679-.437h.001c.747 0 .89.871.89 1.565a19.484 19.484 0 0 1-1.004 5.121l.041-.137c-1.066-2.829-1.139-5.197-.602-6.11zM6.16 18.176c-1.818 3.05-3.562 4.946-4.63 4.946a.87.87 0 0 1-.529-.185l.002.001a.78.78 0 0 1-.248-.767l-.001.005c.214-1.094 2.243-2.622 5.41-4.003z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Adjust(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 20.501V3.5h-.064a8.269 8.269 0 0 0-4.24 1.162l.038-.021c-2.554 1.5-4.242 4.233-4.242 7.36s1.688 5.86 4.202 7.339l.04.022a8.219 8.219 0 0 0 4.2 1.14h.07H12zM24.001 12v.09c0 2.187-.598 4.235-1.64 5.988l.03-.054a12.045 12.045 0 0 1-4.311 4.337l-.056.031c-1.729 1.012-3.807 1.609-6.024 1.609s-4.295-.597-6.081-1.64l.057.031a12.045 12.045 0 0 1-4.337-4.311l-.031-.056C.596 16.296-.001 14.218-.001 12.001s.597-4.295 1.64-6.081l-.031.057A12.045 12.045 0 0 1 5.919 1.64l.056-.031C7.704.597 9.782 0 11.999 0s4.295.597 6.081 1.64l-.057-.031A12.045 12.045 0 0 1 22.36 5.92l.031.056a11.61 11.61 0 0 1 1.61 5.934v.095z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Adobe(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.097 0h10.025v24zm-7.063 0H0v24zm-.853 19.171l4.384-10.329l6.386 15.156h-4.184l-1.91-4.827z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Aids(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.47 13.321l-2.339 3.774l4.281 6.904l3.37-2.086zM13.555 0H2.225L-.001 3.97l5.313 8.58l2.339-3.781l-2.972-4.8h6.418L-.001 21.913l3.376 2.086L15.78 3.969z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Airplay(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.076 24v-.007a1.201 1.201 0 0 1-.701-2.066l.001-.001l7.076-8.262a1.2 1.2 0 0 1 1.896-.003l.002.003l7.076 8.261a1.2 1.2 0 0 1-.695 2.067h-.005v.007zm14.35-6l-2.057-2.4h7.025V2.4H2.4v13.2h7.024L7.367 18H2.4A2.402 2.402 0 0 1 0 15.6V2.401a2.402 2.402 0 0 1 2.4-2.4h23.999a2.402 2.402 0 0 1 2.4 2.4V15.6a2.402 2.402 0 0 1-2.4 2.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Algolia(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.267 8.09h.068a4.67 4.67 0 0 1 3.333 1.393a4.451 4.451 0 0 1 1.393 3.382v-.007v.013a4.677 4.677 0 0 1-1.419 3.361l-.001.001a4.582 4.582 0 0 1-3.296 1.393l-.083-.001h.004l-.079.001a4.581 4.581 0 0 1-3.295-1.392l-.001-.001a4.582 4.582 0 0 1-1.393-3.296l.001-.083v.004l-.001-.079c0-1.292.533-2.46 1.392-3.295l.001-.001a4.582 4.582 0 0 1 3.38-1.392h-.004zm3.375 2.999q.107-.054.054-.214a3.8 3.8 0 0 0-1.349-1.383l-.017-.01a3.967 3.967 0 0 0-1.891-.589h-.01l-.018-.001a.143.143 0 0 0-.142.162v-.001v3.589q0 .054.08.08a.425.425 0 0 0 .133.026h.001zM20.841.001h.027a3.03 3.03 0 0 1 2.195.937l.001.001c.576.554.934 1.331.934 2.192v.033v-.002v17.705a3.03 3.03 0 0 1-.937 2.195l-.001.001a3.033 3.033 0 0 1-2.196.938H3.133a3.03 3.03 0 0 1-2.195-.937l-.001-.001a3.033 3.033 0 0 1-.938-2.196v-.03v.001V3.133c0-.863.359-1.641.937-2.195L.937.937a3.033 3.033 0 0 1 2.196-.938h.029h-.001zM10.017 4.555v.858c0 .036.022.066.053.079h.001a.235.235 0 0 0 .107.026a6.753 6.753 0 0 1 2.076-.32h.014h-.001h.013c.742 0 1.456.117 2.125.334l-.049-.014a.156.156 0 0 0 .107-.053a.153.153 0 0 0 .054-.107v-.804c0-.293-.123-.558-.32-.745a1.027 1.027 0 0 0-.745-.32h-2.425a.939.939 0 0 0-.707.319l-.001.001a1.07 1.07 0 0 0-.298.743v.004zM5.463 6.536a.863.863 0 0 0-.32.671v.027v-.001v.005c0 .293.123.558.32.745l.429.429q.107.107.214-.054q.375-.48.777-.91a6.03 6.03 0 0 1 .918-.787l.02-.013q.16-.107.054-.214l-.429-.429a1.128 1.128 0 0 0-.733-.268h-.017h.001h-.017c-.28 0-.537.101-.735.269l.002-.001zm6.8 13.07h.068a6.618 6.618 0 0 0 3.391-.928l-.031.017a6.852 6.852 0 0 0 2.446-2.432l.017-.032a6.529 6.529 0 0 0 .914-3.353v-.101c0-1.234-.339-2.39-.928-3.378l.017.03a6.852 6.852 0 0 0-2.432-2.446l-.032-.017a6.536 6.536 0 0 0-3.346-.911h-.058h.003h-.055a6.574 6.574 0 0 0-3.377.927l.03-.017a6.772 6.772 0 0 0-2.446 2.458l-.017.033a6.629 6.629 0 0 0-.909 3.369v.034v-.002v.041c0 1.23.338 2.38.927 3.364l-.017-.03a6.852 6.852 0 0 0 2.432 2.446l.032.017a6.483 6.483 0 0 0 3.335.911h.04h-.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Amazon(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.394 19.768a.54.54 0 0 1 .352-.039l-.004-.001q.147.04.147.234a.762.762 0 0 1-.201.449c-.184.21-.376.4-.582.576l-.008.006c-.377.317-.795.617-1.235.885l-.045.025c-.531.338-1.151.669-1.796.955l-.092.036c-.713.309-1.556.577-2.428.761l-.089.016c-.915.209-1.965.328-3.043.328h-.032h.002a12.581 12.581 0 0 1-3.275-.436l.088.02a13.962 13.962 0 0 1-2.881-1.062l.081.037a17.952 17.952 0 0 1-2.365-1.43l.055.037a17.604 17.604 0 0 1-1.788-1.42l.015.013a12.824 12.824 0 0 1-1.11-1.154l-.015-.018a.57.57 0 0 1-.132-.217l-.001-.004a.193.193 0 0 1 .014-.161v.001a.272.272 0 0 1 .106-.093l.002-.001a.28.28 0 0 1 .124-.028l.032.002H.289a.28.28 0 0 1 .155.06H.443c1.103.71 2.419 1.45 3.777 2.117l.24.107c2.094.999 4.552 1.583 7.146 1.583c1.259 0 2.486-.138 3.666-.398l-.112.021a23.66 23.66 0 0 0 5.376-1.87l-.14.062zm2.772-1.54a1.486 1.486 0 0 1 .031.941l.003-.011a6.13 6.13 0 0 1-.397 1.414l.015-.041a4.057 4.057 0 0 1-1.134 1.656l-.005.004q-.228.187-.348.121t0-.32c.203-.44.406-.979.572-1.535l.024-.092q.314-1.025.087-1.319a.484.484 0 0 0-.204-.153l-.003-.001a1.18 1.18 0 0 0-.358-.08h-.004l-.395-.034a4.188 4.188 0 0 0-.48 0l.011-.001q-.294.014-.422.026l-.415.04l-.302.026q-.08.014-.174.02t-.147.014l-.114.014a.66.66 0 0 1-.093.006h-.134l-.04-.006l-.026-.02l-.02-.04q-.08-.214.63-.536a4.587 4.587 0 0 1 1.355-.397l.024-.003a6.164 6.164 0 0 1 1.475-.011l-.028-.003q.829.078 1.017.319zm-5.28-5.933c.004.31.07.604.186.871l-.006-.015c.12.297.264.553.436.787l-.007-.01q.248.335.502.616c.136.155.281.296.435.424l.006.005l.174.147l-3.04 3q-.536-.495-1.058-1.011t-.777-.783l-.254-.268a2.674 2.674 0 0 1-.329-.43l-.007-.012a5.199 5.199 0 0 1-1.291 1.363l-.015.01a5.46 5.46 0 0 1-1.665.84l-.039.01a7.58 7.58 0 0 1-1.86.308l-.016.001a4.724 4.724 0 0 1-1.873-.292l.033.011a5.58 5.58 0 0 1-1.585-.889l.011.009a3.868 3.868 0 0 1-1.103-1.488l-.009-.025a5.33 5.33 0 0 1-.416-2.083l.001-.098v-.059c0-.719.138-1.406.389-2.036l-.013.037c.234-.603.56-1.121.967-1.563l-.003.003a6.392 6.392 0 0 1 1.396-1.096l.03-.016a8.79 8.79 0 0 1 1.577-.743l.064-.02a12.581 12.581 0 0 1 1.653-.448l.088-.015a13.92 13.92 0 0 1 1.533-.242l.067-.006q.663-.06 1.333-.087V5.304a2.39 2.39 0 0 0-.29-1.31l.006.012a1.774 1.774 0 0 0-1.629-.709l.008-.001q-.08 0-.221.014a2.725 2.725 0 0 0-.561.167l.018-.007c-.286.106-.534.24-.759.402l.009-.006c-.29.228-.538.492-.742.788l-.008.012a4.51 4.51 0 0 0-.631 1.253l-.009.032l-3.938-.362c.002-.573.11-1.121.305-1.625l-.011.031a5.227 5.227 0 0 1 .903-1.522l-.006.007a6.5 6.5 0 0 1 1.421-1.257l.025-.015A7.096 7.096 0 0 1 9.269.34L9.32.328A9.332 9.332 0 0 1 11.8 0h.075h-.004l.101-.001c.828 0 1.627.123 2.381.351l-.058-.015c.661.188 1.238.47 1.752.836l-.017-.012c.42.314.779.681 1.075 1.096l.01.015c.245.331.449.712.593 1.122l.009.03c.101.293.162.63.167.981v.003zm-9 .282l-.001.074c0 .715.371 1.343.931 1.703l.008.005a2.122 2.122 0 0 0 1.877.29l-.015.004a2.263 2.263 0 0 0 1.523-1.632l.003-.015a4.45 4.45 0 0 0 .188-1.292v-.064v.003v-2.17a9.65 9.65 0 0 0-1.547.17l.06-.01a6.859 6.859 0 0 0-1.47.466l.044-.017a2.585 2.585 0 0 0-1.6 2.488v-.005z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ambulance(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m34.534 10.942l-5.662-8.223a1.02 1.02 0 0 0-.831-.436H23.02V1.001a.997.997 0 0 0-.997-.997H.995A.995.995 0 0 0 0 .999v18.563c0 .349.283.632.632.634h2.983v-.059a4.381 4.381 0 0 1 8.762 0v.062v-.003h7.065v-.059a4.381 4.381 0 0 1 8.762 0v.062v-.003h5.983a.634.634 0 0 0 .632-.634v-7.76a1.63 1.63 0 0 0-.286-.867l.004.006zM15.25 8.637h-2.698v2.697h-2.448V8.64H7.406V6.19h2.698V3.483h2.448V6.18h2.698zm17.903 2.59a.22.22 0 0 1-.195.116h-9.936V3.482h4.451a.942.942 0 0 1 .764.397l.002.003l4.906 7.12a.203.203 0 0 1 .009.217l.001-.001z"/><path fill="currentColor" d="M7.986 16.414h-.001a3.794 3.794 0 1 0 3.794 3.794v-.003a3.794 3.794 0 0 0-3.793-3.791m0 5.083a1.291 1.291 0 1 1 1.291-1.291v.001a1.29 1.29 0 0 1-1.29 1.29zm15.829-5.083h-.001a3.794 3.794 0 1 0 3.794 3.794v-.003a3.794 3.794 0 0 0-3.793-3.791m0 5.083a1.291 1.291 0 1 1 1.291-1.291v.001a1.29 1.29 0 0 1-1.29 1.29zM26.806 1.93v-.712a.667.667 0 0 0-.584-.72h-1.123a.66.66 0 0 0-.587.723v-.003v.712z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AmericanExpress(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M33.6 24H2.4A2.4 2.4 0 0 1 0 21.6V11.504h1.729l.39-.944h.873l.389.94h3.402v-.715l.304.72h1.766l.304-.732v.73h8.453V9.959h.16c.114.004.147.014.147.204v1.342h4.373v-.359c.447.227.974.36 1.533.36l.095-.001h-.005h1.84l.394-.94h.873l.385.94h3.546v-.894l.536.894h2.836V5.601h-2.807v.697l-.393-.697h-2.886v.697l-.361-.697h-3.897a3.47 3.47 0 0 0-1.709.353l.021-.009v-.344h-2.688v.344a1.627 1.627 0 0 0-1.149-.343h.006h-9.823l-.659 1.525l-.677-1.525H4.207v.697l-.341-.697h-2.64l-1.223 2.8v-6a2.4 2.4 0 0 1 2.4-2.4h31.2a2.4 2.4 0 0 1 2.4 2.4v10.48H34.13c-.03-.002-.066-.002-.101-.002c-.434 0-.837.13-1.173.353l.008-.005v-.346h-2.77a1.923 1.923 0 0 0-1.215.349l.006-.004v-.346h-4.946v.346a2.683 2.683 0 0 0-1.327-.346h-.039h.002h-3.263v.346a2.253 2.253 0 0 0-1.436-.345l.009-.001h-3.651l-.836.904l-.782-.904H7.162v5.908h5.352l.861-.918l.811.918h3.299v-1.383h.46a3.228 3.228 0 0 0 1.296-.217l-.022.008v1.594h2.72v-1.539h.131c.166 0 .183.006.183.174v1.366h8.266l.101.002c.474 0 .916-.142 1.284-.385l-.009.005v.378h2.622l.125.002c.491 0 .958-.101 1.382-.284l-.023.009v3.082a2.4 2.4 0 0 1-2.4 2.4zm-12.495-6.039h-1.018v-4.235h2.336a2.45 2.45 0 0 1 1.233.206l-.016-.006c.313.172.522.5.522.876l-.002.067v-.003l.001.038c0 .486-.293.904-.713 1.086l-.008.003c.201.072.369.195.494.354l.002.002a1.226 1.226 0 0 1 .168.779l.001-.006v.838h-1.016v-.53a1.212 1.212 0 0 0-.163-.834l.003.005a.98.98 0 0 0-.748-.187l.006-.001h-1.082v1.547zm0-3.36v.951h1.23a.97.97 0 0 0 .506-.09l-.006.003a.459.459 0 0 0 .21-.385l-.001-.023v.001a.405.405 0 0 0-.207-.38l-.002-.001a.966.966 0 0 0-.484-.08h.004zM12.08 17.96H8.073v-4.234h4.07l1.245 1.388l1.287-1.388h3.233c1.148 0 1.706.457 1.706 1.395c0 .955-.577 1.419-1.76 1.419h-1.265v1.419h-1.967l-1.246-1.399zm3.501-3.78l-1.554 1.67l1.554 1.724zm-6.499 2.055v.842h2.488l1.15-1.242l-1.106-1.234H9.081v.77h2.222v.863zm7.507-1.633v1.078h1.307c.4 0 .63-.204.63-.56c0-.34-.214-.519-.619-.519zm18.038 3.36h-1.954v-.91h1.946a.559.559 0 0 0 .411-.106l-.001.001a.372.372 0 0 0 .119-.273v-.016a.348.348 0 0 0-.123-.266a.513.513 0 0 0-.358-.094h.002l-.187-.006c-.914-.024-1.949-.052-1.949-1.305c0-.61.382-1.261 1.451-1.261h2.017v.902h-1.845a.693.693 0 0 0-.412.082l.004-.002a.329.329 0 0 0-.147.302v-.001v.011a.32.32 0 0 0 .22.304l.002.001a1.12 1.12 0 0 0 .392.048h-.003l.549.014a1.583 1.583 0 0 1 1.151.343l-.003-.002c.03.024.056.05.079.079l.001.001l.012 1.612a1.577 1.577 0 0 1-1.381.541zm-3.949 0h-1.972v-.91h1.962a.577.577 0 0 0 .415-.106l-.002.001a.373.373 0 0 0 .118-.273v-.01a.363.363 0 0 0-.123-.272a.526.526 0 0 0-.363-.094h.002l-.186-.006c-.911-.024-1.945-.052-1.945-1.305c0-.61.38-1.261 1.447-1.261h2.028v.902h-1.856a.686.686 0 0 0-.409.082l.004-.002a.354.354 0 0 0 .066.619l.002.001a1.134 1.134 0 0 0 .397.048h-.003l.545.014a1.609 1.609 0 0 1 1.158.344l-.003-.003a1.183 1.183 0 0 1 .302.901v-.005c.003.883-.532 1.333-1.587 1.333zm-2.578 0h-3.38v-4.237h3.377v.875H25.73v.77h2.312v.863H25.73v.842l2.37.004v.88zm1.97-7.286h-2.061l-.394-.944h-2.102l-.382.944h-1.184a2.18 2.18 0 0 1-1.472-.471l.005.003a2.122 2.122 0 0 1-.54-1.625l-.001.008a2.213 2.213 0 0 1 .548-1.657l-.002.002a2.033 2.033 0 0 1 1.545-.492l-.008-.001h.98v.903h-.96a.963.963 0 0 0-.78.251l.001-.001a1.393 1.393 0 0 0-.291.964v-.005a1.457 1.457 0 0 0 .281.998l-.003-.004a.983.983 0 0 0 .709.218h-.004h.454l1.431-3.33h1.52l1.713 4v-4h1.541l1.778 2.948V6.437h1.04v4.232h-1.444l-1.92-3.178v3.178zm-3.499-3.518l-.697 1.688h1.397zm-9.799 3.516H15.76V6.44h2.328a2.357 2.357 0 0 1 1.241.209l-.015-.006a.996.996 0 0 1 .514.871l-.002.07v-.003v.031a1.2 1.2 0 0 1-.706 1.094l-.008.003c.201.076.37.198.499.354l.002.002a1.213 1.213 0 0 1 .167.783l.001-.006v.831H18.76l-.004-.534v-.08a1.11 1.11 0 0 0-.163-.749l.003.005a.995.995 0 0 0-.744-.181l.006-.001h-1.085v1.54zm0-3.353v.94H18a.914.914 0 0 0 .505-.09l-.005.002a.437.437 0 0 0 .21-.373l-.001-.028v.001a.39.39 0 0 0-.211-.373l-.002-.001a1.025 1.025 0 0 0-.483-.08h.003zM6.056 10.674H4l-.389-.944H1.503l-.392.944h-1.1L1.824 6.44h1.503l1.721 4.007V6.44h1.651l1.324 2.871L9.24 6.44h1.685v4.232h-1.04L9.883 7.36l-1.467 3.313h-.888L6.057 7.355v3.318zM2.552 7.158l-.689 1.688h1.382zm18.888 3.515h-1.032V6.44h1.033v4.232zm-6.386 0H11.68V6.44h3.38v.88h-2.368v.763h2.311v.869H12.69v.846h2.368v.874z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AmericanSignLanguageInterpreting(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.531.095c.291.278.172.57-.669 1.622a5.587 5.587 0 0 0-1.167 1.881l-.013.039c-.093.348-.059.358.3.089a11.87 11.87 0 0 1 1.533-.959l.067-.032c.212-.111.664-.338.682-.338s.095-.042.198-.091a7.997 7.997 0 0 1 1.86-.65l.055-.01c.842-.096 1.375.304 1.178.88c-.087.255-.234.358-.833.575c-.355.13-.516.191-.633.24l-.36.16a8.917 8.917 0 0 0-1.072.536l.043-.024a8.78 8.78 0 0 0-.993.691l.017-.013a5.17 5.17 0 0 0-.664.63l-.005.006c.022.035.182.029.33-.014a9.847 9.847 0 0 1 2.92-.152l-.04-.003c2.08.259 4 1.586 3.452 2.393c-.24.358-.51.374-1.127.07a5.802 5.802 0 0 0-2.922-.779c-.566 0-1.112.08-1.63.229l.041-.01c-.62.182-.622.24-.008.267c.154.003.301.015.447.034l-.02-.002c.071.018.156.033.244.042l.007.001c.277.037.525.093.763.169l-.031-.009a4.76 4.76 0 0 1 1.477.667l-.017-.011c.6.446 1.11.971 1.526 1.566l.014.022c.658.886.737 1.806.167 1.994c-.358.118-.574-.039-.96-.7a3.222 3.222 0 0 0-2.283-1.777l-.02-.003a2.248 2.248 0 0 0-.822.029l.015-.003c-2.098.293-3.269 2.911-1.844 4.124a3.268 3.268 0 0 0 1.935.629c.995 0 1.887-.441 2.491-1.139l.004-.004c.669-.748.98-.927 1.342-.772c.784.337-.265 2.156-1.755 3.04c-.137.08-.741.371-.887.425l-.257.095c-.063.023-.174.056-.244.08l-.32.087a13.798 13.798 0 0 1-3.621.28l.035.002c-2.16.04-2.42.104-3.574.88a21.4 21.4 0 0 0-.971.689c-.303.223-.291.223-.632.026a4.259 4.259 0 0 1-2.274-4.027l-.001.012c.018-.295 0-.267.39-.539c.874-.61 2.376-1.749 2.44-1.85c.022-.034.223-1.026.254-1.241a1.22 1.22 0 0 1 .041-.195l-.002.008c.018-.056.034-.126.046-.197l.001-.009c.021-.123.094-.473.208-.996c.219-.921.449-1.677.717-2.414l-.048.151a5.33 5.33 0 0 0 .111-.27A15.508 15.508 0 0 1 7.851.713l.005-.005C8.594.095 9.27-.158 9.528.088zm18.245 6.098A4.433 4.433 0 0 1 30.01 9.1l.006.03a6.337 6.337 0 0 1 .055 1.261l.001-.016c-.022.146-.027.15-1.066.874c-.185.13-.704.518-.96.72c-.111.086-.31.24-.446.35c-.301.234-.4.339-.4.431c0 .038-.042.266-.091.507s-.107.524-.126.63a30.124 30.124 0 0 1-.591 2.533a14.048 14.048 0 0 1-1.009 2.528c-.461.923-.94 1.709-1.475 2.452l.035-.052l-.283.39c-1.17 1.61-2.8 2.685-3.171 2.09c-.18-.291-.052-.56.715-1.52c.305-.357.595-.752.855-1.167l.025-.042c.022-.042.067-.13.102-.193a2.36 2.36 0 0 0 .265-.787l.002-.013c-.019-.06-.097-.026-.298.127c-.661.485-1.41.934-2.2 1.311l-.085.036l-.338.16c-.293.135-1.007.411-1.257.486a1.79 1.79 0 0 1-.857.142h.006c-.468-.016-.738-.144-.875-.413a.685.685 0 0 1 .316-.861l.004-.002c.088-.044.2-.092.315-.134l.024-.008c.248-.087.622-.226.669-.248a4.22 4.22 0 0 1 .218-.092a8.355 8.355 0 0 0 2.77-1.756l-.003.003c.171-.167.211-.254.106-.23c-.356.089-.8.169-1.251.223l-.054.005a8.744 8.744 0 0 1-2.59-.146l.057.01c-1.93-.469-3.306-1.652-2.72-2.341c.271-.32.457-.308 1.217.063c.129.069.29.144.455.211l.034.012c.023.001.044.01.061.024c.227.099.504.194.789.269l.042.009c.43.128.923.202 1.434.202c.497 0 .978-.07 1.434-.201l-.037.009c.472-.138.529-.16.57-.214c.056-.073.005-.089-.293-.093a6.517 6.517 0 0 1-2.162-.451l.044.015a5.116 5.116 0 0 1-1.728-1.213l-.003-.003c-.08-.09-.16-.182-.181-.203a5.27 5.27 0 0 1-1.001-1.631l-.012-.036c-.16-.648.214-1.126.737-.93c.186.07.291.186.507.565c.941 1.642 2.19 2.24 3.665 1.746a2.801 2.801 0 0 0 1.904-3.218l.003.018c-.334-1.146-2.026-1.76-3.5-1.263a3.817 3.817 0 0 0-1.809 1.327l-.007.01c-.324.353-.473.435-.78.435c-.438 0-.63-.359-.428-.8a.863.863 0 0 0 .055-.138l.001-.006c.119-.275.25-.51.4-.73l-.01.015a4.76 4.76 0 0 1 1.292-1.263l.018-.011a6.906 6.906 0 0 1 3.241-1.033l.02-.001c.12-.012 1.254-.042 2.944-.08a4.61 4.61 0 0 0 2.15-.411l-.029.012a16.302 16.302 0 0 0 1.839-1.239l-.036.027c.191-.16.219-.155.525.022z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ampproject(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.96.66L13.62 9.9h2.16c.6 0 .84.48.48 1.08L8.58 23.52a12.42 12.42 0 0 0 3.418.48H12c6.628-.005 11.998-5.379 11.998-12.008c0-5.209-3.317-9.644-7.955-11.306L15.959.66zm-5.58 13.38H8.22c-.6 0-.84-.48-.48-1.08L15.42.48A12.42 12.42 0 0 0 12.002 0H12C5.372.005.002 5.379.002 12.008c0 5.209 3.317 9.644 7.955 11.306l.084.026z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Anchor(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.149 19.17C19.834 22.121 16.268 24 12.263 24c-.097 0-.194-.001-.29-.003h.014a12.513 12.513 0 0 1-10.134-4.8l-.021-.027L-.001 21v-5.994h5.994l-1.644 1.646c1.331 2.188 3.52 3.742 6.089 4.197l.054.008v-9.061c-2.606-.689-4.495-3.026-4.495-5.803a5.994 5.994 0 1 1 7.535 5.794l-.042.009v9.062a8.965 8.965 0 0 0 6.121-4.166l.022-.039l-1.652-1.647h5.994V21zM14.988 6.016a2.997 2.997 0 1 0-5.995-.001a2.997 2.997 0 0 0 5.995.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Android(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.11 5.12a.562.562 0 1 0 0-1.124H7.1a.517.517 0 0 0-.379.165a.575.575 0 0 0 0 .794c.095.102.23.166.38.166h.011h-.001zm6.087 0h.01c.15 0 .285-.064.379-.165a.575.575 0 0 0 0-.794a.518.518 0 0 0-.38-.166h-.011h.001a.562.562 0 1 0 0 1.124zM1.486 7.774c.811.004 1.468.66 1.471 1.471v6.23c0 .4-.163.763-.426 1.025a1.406 1.406 0 0 1-1.017.433h-.031h.002h-.008a1.478 1.478 0 0 1-1.478-1.478V9.224c0-.4.166-.761.433-1.019c.26-.266.623-.431 1.025-.431zm15.288.274v9.635c0 .429-.177.818-.462 1.096a1.511 1.511 0 0 1-1.089.462H15.2h.001h-1.082v3.274a1.486 1.486 0 0 1-2.972 0v-3.274h-1.99v3.282c0 .816-.662 1.478-1.478 1.478H7.65c-.4 0-.761-.166-1.019-.433a1.43 1.43 0 0 1-.433-1.026v-.028v.001l-.014-3.274H5.108c-.87 0-1.576-.706-1.576-1.576v-.011v.001v-9.606zm-3.346-5.84a6.391 6.391 0 0 1 2.452 2.19l.015.024c.579.866.923 1.932.923 3.077v.032v-.002H3.476v-.031c0-1.146.345-2.211.936-3.098l-.013.02a6.363 6.363 0 0 1 2.444-2.197l.036-.017L5.855.316a.19.19 0 0 1 .07-.288h.001a.2.2 0 0 1 .289.085l.001.001l1.04 1.904c.854-.383 1.85-.606 2.899-.606s2.045.223 2.945.624l-.046-.018l1.04-1.904a.2.2 0 0 1 .29-.086h-.001a.19.19 0 0 1 .072.289zm6.88 7.04v6.21c0 .816-.662 1.478-1.478 1.478h-.029c-.4 0-.761-.166-1.019-.433a1.43 1.43 0 0 1-.433-1.026v-.028v.001v-6.231c0-.399.166-.76.433-1.016a1.423 1.423 0 0 1 1.017-.426h.024h-.001h.028c.4 0 .763.163 1.025.426c.267.257.433.617.433 1.017z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Angelist(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.824 9.957q1.104-2.999 1.722-4.952c.262-.686.477-1.494.609-2.329l.009-.07l.002-.072c0-.25-.078-.482-.211-.673l.003.004a.681.681 0 0 0-.539-.265l-.031.001h.001q-.461 0-.939.766c-.402.703-.76 1.519-1.03 2.373l-.025.09l-1.647 4.762zm-2.191 4.432a10.407 10.407 0 0 1-1.405-.157l.062.009a7.03 7.03 0 0 1-1.26-.346l.049.016q.263.528.469 1.055c.113.274.228.621.323.975l.015.064c.263-.328.535-.621.827-.893l.005-.005c.273-.256.572-.492.889-.704l.025-.016zM8.772 9.397L6.993 4.256a15.267 15.267 0 0 0-1.111-2.641l.04.082q-.379-.601-.824-.601H5.08a.688.688 0 0 0-.541.263l-.001.001a1.08 1.08 0 0 0-.214.711v-.003c.088.972.292 1.868.6 2.713L4.9 4.707q.577 1.878 1.73 4.927a.554.554 0 0 1 .276-.254l.004-.001a1.307 1.307 0 0 1 .497-.074h-.003q.099 0 .396.017t.972.08zm-1.615 7.958a.678.678 0 0 0 .51-.254l.001-.001a.81.81 0 0 0 .23-.551v-.001a5.705 5.705 0 0 0-.468-1.493l.015.035c-.35-.871-.728-1.61-1.164-2.311l.036.062a5.325 5.325 0 0 0-.97-1.242l-.002-.002a1.409 1.409 0 0 0-.901-.42h-.005a1.154 1.154 0 0 0-.764.434l-.002.003a1.232 1.232 0 0 0-.419.811v.005c.055.445.201.847.419 1.199l-.007-.013c.345.639.712 1.189 1.122 1.702l-.017-.022a9.26 9.26 0 0 0 1.358 1.511l.009.008c.266.276.616.47 1.009.539l.011.002zm-4.943-.395q.23.281.64.824q1.088 1.499 2.01 1.499h.007a.871.871 0 0 0 .555-.199l-.001.001a.547.547 0 0 0 .247-.41v-.002a2.046 2.046 0 0 0-.334-.831l.004.007c-.294-.495-.593-.92-.921-1.321l.015.019a9.855 9.855 0 0 0-1.091-1.223l-.004-.004a1.229 1.229 0 0 0-.688-.383l-.008-.001c-.445.027-.83.265-1.06.613l-.003.005a2.277 2.277 0 0 0-.486 1.461v-.002c.016.549.139 1.064.348 1.533l-.01-.026c.271.639.6 1.189.992 1.688l-.012-.015a6.583 6.583 0 0 0 2.376 1.982l.038.017a7.096 7.096 0 0 0 3.098.701l.095-.001h-.005l.122.001a6.941 6.941 0 0 0 5.25-2.392l.007-.008a8.582 8.582 0 0 0 2.167-6.02v.014l.002-.186c0-.556-.061-1.097-.176-1.618l.009.05a1.683 1.683 0 0 0-.542-.946l-.002-.002a6.73 6.73 0 0 0-2.588-1l-.039-.005a17.824 17.824 0 0 0-4.011-.445h-.071h.004a1.37 1.37 0 0 0-.846.2l.006-.003a.816.816 0 0 0-.247.679v-.004a1.7 1.7 0 0 0 1.24 1.62l.012.003c1.034.329 2.223.518 3.457.518c.21 0 .419-.005.626-.016l-.029.001h.689c.144 0 .273.067.356.172l.001.001c.104.146.172.323.189.515v.004c-.33.252-.712.456-1.126.593l-.028.008c-.51.185-.951.41-1.357.682l.022-.014a5.957 5.957 0 0 0-1.577 1.695l-.015.025a3.644 3.644 0 0 0-.601 1.88v.007c.017.475.109.922.265 1.339l-.01-.029c.109.272.198.591.252.921l.004.026v.115l-.033.148a1.394 1.394 0 0 1-1.142-.839l-.003-.009a4.531 4.531 0 0 1-.419-2.143v.01v-.148a.898.898 0 0 1-.249.162l-.006.002a.721.721 0 0 1-.264.05H7.3c-.097 0-.191-.009-.283-.026l.009.001a2.659 2.659 0 0 1-.317-.08l.02.006c.028.096.054.216.072.338l.002.016c.013.08.022.174.025.269v.008c0 .416-.185.788-.478 1.039l-.002.002c-.29.271-.681.437-1.11.437H5.21h.001a3.125 3.125 0 0 1-2.107-1.011l-.002-.002a2.833 2.833 0 0 1-1.07-1.975l-.001-.011v-.014c0-.109.015-.214.044-.314l-.002.008a.551.551 0 0 1 .14-.24zm11.798-6.723a2.97 2.97 0 0 1 2.086 1.243l.006.009a5.866 5.866 0 0 1 .608 3.084l.001-.019a9.468 9.468 0 0 1-2.503 6.804l.006-.007a8.376 8.376 0 0 1-6.398 2.644h.013h-.007a8.246 8.246 0 0 1-3.023-.571l.056.019a7.401 7.401 0 0 1-2.483-1.544l.004.003a7.855 7.855 0 0 1-1.759-2.255l-.02-.043a5.504 5.504 0 0 1-.593-2.478v-.001a3.454 3.454 0 0 1 .617-2.202l-.007.011a3.19 3.19 0 0 1 1.91-1.052l.018-.002a8.016 8.016 0 0 1-.345-.91l-.017-.062a2.264 2.264 0 0 1-.116-.604v-.006a2.195 2.195 0 0 1 .78-1.517l.003-.003c.362-.429.88-.716 1.466-.773l.009-.001c.225.003.44.042.64.111l-.014-.004c.291.103.542.222.776.365l-.018-.01Q4.439 6.874 3.862 4.922c-.279-.786-.483-1.7-.573-2.648l-.004-.046A2.363 2.363 0 0 1 3.816.6l-.003.004A1.805 1.805 0 0 1 5.251.003h-.004q1.548 0 3.918 6.92q.412 1.187.626 1.829q.182-.51.511-1.45Q12.675.447 14.388.447l.071-.001c.505 0 .958.22 1.27.568l.001.002a2.205 2.205 0 0 1 .508 1.531v-.005c-.09.979-.285 1.879-.576 2.735l.024-.083q-.553 1.928-1.672 5.038z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleDobuleDown(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.613 13.935L0 2.344L2.352 0l9.261 9.239L20.874 0l2.352 2.344zm0 10.065L0 12.409l2.352-2.344l9.261 9.239l9.261-9.239l2.352 2.344z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleDobuleLeft(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.4 12L22.378 0L24.8 2.43L15.253 12l9.547 9.57L22.378 24zM0 12L11.978 0L14.4 2.43L4.853 12l9.547 9.57L11.978 24z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleDobuleRight(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.4 12L2.422 24L0 21.57L9.547 12L0 2.43L2.422 0zm10.4 0L12.822 24L10.4 21.57L19.947 12L10.4 2.43L12.822 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleDobuleUp(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m11.613 10.065l11.613 11.591L20.874 24l-9.261-9.239L2.352 24L0 21.656zm0-10.065l11.613 11.591l-2.352 2.344l-9.261-9.239l-9.261 9.239L0 11.591z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleDown(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.43 4.8L0 7.222L12 19.2L24 7.222L21.57 4.8L12 14.347z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleLeft(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.2 2.43L16.778 0L4.8 12l11.978 12l2.422-2.43L9.653 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleRight(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.8 21.57L7.222 24L19.2 12L7.222 0L4.8 2.43L14.347 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AngleUp(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.57 19.2L24 16.778L12 4.8L0 16.778L2.43 19.2L12 9.653z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Angularjs(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m11.172 24l-9.468-5.244L0 3.984L11.172 0l11.172 3.984l-1.704 14.772zm0-21.348l-6.984 15.66h2.604l1.404-3.504h5.928l1.404 3.504h2.604zm2.04 9.996h-4.08l2.04-4.908z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppStore(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.434 4.693a.988.988 0 0 1 .598-.458l.007-.002a.873.873 0 0 1 .754.099l-.003-.002c.219.139.382.349.458.597l.002.008a.895.895 0 0 1-.123.777l.002-.003l-4.21 7.307h3.077c.467 0 .87.274 1.057.671l.003.007a1.303 1.303 0 0 1 .02 1.313l.003-.007H5.105a.982.982 0 0 1-.699-.291a.958.958 0 0 1-.29-.688v-.027c0-.27.111-.513.29-.688a.98.98 0 0 1 .699-.291h.028h-.001h2.516l3.193-5.564l-.971-1.739a.894.894 0 0 1-.119-.781l-.002.006a1.08 1.08 0 0 1 .456-.602l.004-.003a.86.86 0 0 1 .757-.095l-.006-.002c.26.069.473.235.602.455l.003.005l.435.774zM7.257 17.66a1.2 1.2 0 0 1-.62.458l-.008.002a.882.882 0 0 1-.753-.099l.003.002a.967.967 0 0 1-.459-.598l-.001-.007a1.092 1.092 0 0 1 .099-.78l-.003.006l.726-1.21a1.71 1.71 0 0 1 1.932.577l.003.004zm11.661-4.644h.027c.273 0 .52.111.699.291c.179.175.29.418.29.688v.014v-.001v.014c0 .27-.111.513-.29.688a.98.98 0 0 1-.699.291h-.028h.001h-1.403l.968 1.645a.988.988 0 0 1 .071.782l.002-.007a1.08 1.08 0 0 1-.456.602l-.004.003a.867.867 0 0 1-.755.095l.006.002a.989.989 0 0 1-.602-.455l-.003-.005l-3.629-6.338a2.818 2.818 0 0 1-.384-2.05l-.003.018c.091-.523.352-.974.723-1.304l.002-.002l2.903 5.032zM11.999.001h.077c2.185 0 4.229.602 5.976 1.65l-.053-.029a12.233 12.233 0 0 1 4.348 4.322l.031.056a11.51 11.51 0 0 1 1.621 5.923v.081V12v.077c0 2.185-.602 4.229-1.65 5.976l.029-.053a12.233 12.233 0 0 1-4.322 4.348l-.056.031A11.51 11.51 0 0 1 12.077 24h-.081H12h-.077c-2.185 0-4.229-.602-5.976-1.65l.053.029a12.228 12.228 0 0 1-4.348-4.323L1.621 18A11.51 11.51 0 0 1 0 12.077v-.154c0-2.185.602-4.229 1.65-5.976L1.621 6a12.233 12.233 0 0 1 4.322-4.348l.056-.031A11.51 11.51 0 0 1 11.922 0h.081h-.004zM22.451 12v-.09c0-1.901-.52-3.681-1.426-5.205l.026.047a10.48 10.48 0 0 0-3.772-3.771l-.05-.027a10.17 10.17 0 0 0-5.183-1.402h-.075h.004h-.075c-1.899 0-3.677.521-5.197 1.429l.046-.026a10.477 10.477 0 0 0-3.771 3.771l-.027.05a10.169 10.169 0 0 0-1.403 5.186v.065v-.003v.075c0 1.899.521 3.677 1.429 5.197l-.026-.046a10.48 10.48 0 0 0 3.772 3.771l.05.027a10.169 10.169 0 0 0 5.186 1.403h.067h-.003h.075c1.899 0 3.677-.521 5.197-1.429l-.046.026a10.478 10.478 0 0 0 3.775-3.773l.027-.05a10.125 10.125 0 0 0 1.403-5.174v-.055v.003z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Apple(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.091 17.524a13.902 13.902 0 0 1-1.805 3.651l.03-.046q-1.861 2.827-3.706 2.827a6.649 6.649 0 0 1-2.064-.478l.044.016a6.254 6.254 0 0 0-2.167-.461h-.011a5.531 5.531 0 0 0-2.083.49l.035-.014a5.137 5.137 0 0 1-1.889.49L6.46 24q-2.192 0-4.342-3.735a14.628 14.628 0 0 1-2.12-7.236v-.02A8.524 8.524 0 0 1 1.644 7.59l-.017.024a5.005 5.005 0 0 1 4.061-2.08h.037h-.002a9.587 9.587 0 0 1 2.621.452l-.068-.019a8.888 8.888 0 0 0 1.95.428l.041.004a7.198 7.198 0 0 0 2.111-.508l-.048.017a7.944 7.944 0 0 1 2.48-.49h.085a5.25 5.25 0 0 1 3.019.949l-.017-.011c.572.416 1.066.89 1.488 1.424l.012.016a9.451 9.451 0 0 0-1.627 1.678l-.017.024a5.09 5.09 0 0 0-.938 2.958v.029v-.001v.062c0 1.181.373 2.275 1.007 3.171l-.012-.017a4.41 4.41 0 0 0 2.248 1.814l.031.01zM14.668.606a5.507 5.507 0 0 1-.432 1.998l.014-.036a5.94 5.94 0 0 1-1.34 1.992l-.001.001a4.102 4.102 0 0 1-1.53 1.031l-.028.009a7.66 7.66 0 0 1-1.475.243l-.025.001v-.028c0-1.372.42-2.647 1.139-3.701l-.015.023A5.754 5.754 0 0 1 14.547.005L14.581 0c.013.043.026.097.035.152l.001.008c.01.063.023.117.039.17l-.003-.01q0 .058.007.144t.007.142z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AppleMusic(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m24 6.124l.001-.097c0-.743-.088-1.465-.253-2.156l.013.063A4.942 4.942 0 0 0 21.598.903l-.02-.012a4.956 4.956 0 0 0-1.847-.723l-.03-.004a10.14 10.14 0 0 0-1.553-.15h-.011c-.04 0-.083-.01-.124-.013H5.988c-.152.01-.3.017-.455.026a6.96 6.96 0 0 0-2.242.415L3.34.427A5.033 5.033 0 0 0 .487 3.175l-.012.033c-.17.409-.297.885-.36 1.38l-.003.028c-.051.343-.087.751-.1 1.165v.016c0 .032-.007.062-.01.093v12.224c.01.14.017.283.027.424c.02.861.202 1.673.516 2.416l-.016-.043a5.01 5.01 0 0 0 3.199 2.792l.035.009c.377.111.817.192 1.271.227l.022.001c.555.053 1.11.06 1.667.06h11.028c.554 0 1.099-.037 1.633-.107l-.063.007a5.319 5.319 0 0 0 2.321-.823l-.021.013a5.078 5.078 0 0 0 1.867-2.176l.013-.032c.166-.383.295-.829.366-1.293l.004-.031a11.897 11.897 0 0 0 .129-2.05V6.127zm-6.424 3.99v5.712l.001.083c0 .407-.09.794-.252 1.14l.007-.017a2.13 2.13 0 0 1-1.373 1.137l-.015.003a4.483 4.483 0 0 1-1.06.173h-.01c-.029.002-.062.002-.096.002a1.871 1.871 0 0 1-.815-3.556l.011-.005c.293-.14.635-.252.991-.32l.027-.004c.378-.082.758-.153 1.134-.24a.621.621 0 0 0 .51-.513v-.003a.863.863 0 0 0 .02-.189V8.069a.739.739 0 0 0-.027-.19l.001.005a.29.29 0 0 0-.301-.234h.001c-.178.013-.34.036-.499.07l.024-.004q-1.14.225-2.28.456l-3.7.748c-.016 0-.032.01-.048.013a.452.452 0 0 0-.39.492v-.002v7.931l.001.095c0 .408-.079.797-.224 1.152l.007-.021a2.138 2.138 0 0 1-1.436 1.235l-.015.003a4.307 4.307 0 0 1-1.067.172h-.008a1.84 1.84 0 0 1-1.919-1.533l-.001-.011a1.867 1.867 0 0 1 1.141-2.071l.013-.004a5.678 5.678 0 0 1 1.072-.305l.036-.005c.287-.06.575-.116.86-.177a.7.7 0 0 0 .6-.693v-.022v.001v-9.04c0-.129.015-.254.044-.374l-.002.011a.696.696 0 0 1 .542-.517l.004-.001c.255-.066.515-.112.774-.165c.733-.15 1.466-.3 2.2-.444l2.27-.46c.67-.134 1.34-.27 2.01-.4c.181-.042.407-.079.637-.104l.027-.002a.493.493 0 0 1 .554.481c.008.067.012.144.012.222v5.733z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ApplePay(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M33.6 24H2.4A2.4 2.4 0 0 1 0 21.6V2.4A2.4 2.4 0 0 1 2.4 0h31.2A2.4 2.4 0 0 1 36 2.4v19.2a2.4 2.4 0 0 1-2.4 2.4m-5.807-7.11V18c.159.022.342.035.528.035h.016h-.001c1.394 0 2.056-.542 2.626-2.147l2.514-7.05h-1.454l-1.686 5.446h-.03L28.62 8.838h-1.496l2.425 6.713l-.13.408a1.14 1.14 0 0 1-1.209.96h.004a7.722 7.722 0 0 1-.42-.029zM5.713 8.054h-.026a2.438 2.438 0 0 0-2.058 1.241l-.006.012c-.844 1.452-.307 3.674.627 5.027c.438.64.918 1.266 1.551 1.266h.034c.265-.017.51-.086.731-.197l-.011.005c.267-.134.581-.213.913-.216h.001c.321.002.624.08.891.216l-.011-.005c.215.112.468.18.737.186h.029c.705-.013 1.147-.66 1.538-1.23c.279-.404.51-.869.672-1.366l.011-.038v-.01l-.018-.008A2.231 2.231 0 0 1 10 10.916v-.001a2.252 2.252 0 0 1 1.055-1.887l.009-.005l.018-.012a2.3 2.3 0 0 0-1.83-.994h-.109a3.339 3.339 0 0 0-1.239.307l.021-.009a1.858 1.858 0 0 1-.583.171l-.009.001a1.856 1.856 0 0 1-.626-.181l.011.005a2.716 2.716 0 0 0-.997-.259h-.008zM23.826 9.87c.88 0 1.366.412 1.366 1.159v.509l-1.786.106c-1.675.101-2.56.78-2.56 1.963a2.022 2.022 0 0 0 2.269 1.986l-.01.001l.065.001c.869 0 1.629-.468 2.041-1.167l.006-.011h.03v1.102h1.325v-4.586c0-1.33-1.061-2.188-2.703-2.188c-1.514 0-2.64.868-2.685 2.063h1.29a1.26 1.26 0 0 1 1.359-.939l-.006-.001zM14 6.304v9.216h1.431v-3.152h1.977a2.929 2.929 0 0 0 3.082-3.046v.006a2.9 2.9 0 0 0-3.04-3.024h.007zM9.253 5.6a2.235 2.235 0 0 0-1.477.76l-.002.003a2.112 2.112 0 0 0-.531 1.55v-.006c.034 0 .07.004.11.004a1.9 1.9 0 0 0 1.378-.707l.003-.004a2.23 2.23 0 0 0 .52-1.608v.007zm14.236 8.901c-.758 0-1.249-.365-1.249-.929c0-.582.47-.917 1.36-.97l1.59-.1v.521a1.55 1.55 0 0 1-1.708 1.477zm-6.418-3.33h-1.644V7.51h1.65a1.722 1.722 0 0 1 1.951 1.832v-.006a1.728 1.728 0 0 1-1.967 1.833z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Archive(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24.975 5.143H1.227A1.258 1.258 0 0 1 0 3.886v-.033v.002v-2.597C0 .574.546.018 1.225.001h23.75c.681.015 1.227.57 1.227 1.253v.034v-.002v2.569l.001.039a1.25 1.25 0 0 1-1.225 1.25h-.001zm.129 17.571V8.138a1.28 1.28 0 0 0-1.28-1.28H2.384a1.28 1.28 0 0 0-1.28 1.28V22.72c0 .707.573 1.28 1.28 1.28h21.434c.71 0 1.285-.576 1.286-1.286m-9.214-9h-5.569a.646.646 0 0 1-.64-.64v-.435a.646.646 0 0 1 .64-.64h5.571a.646.646 0 0 1 .64.64v.43a.645.645 0 0 1-.642.645"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AreaChart(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M32 22v2H0V0h2v22zM26 6l4 14H4v-9l7-9l9 9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDown(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.6 8.034L1.634 8l10.633 10.608L22.901 8l.034.034v5.319L12.268 24L1.602 13.353z"/><path fill="currentColor" d="M14.044 0v18.666h-3.555V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowDownL(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.68 18.24V0h-3.12v18.24H6.4L12.08 24l5.76-5.76z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowExpand(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.918 24L24 14.918V24zM.05 24v-9.082L9.132 24zm3.493-5.59l6.385-6.385l-6.435-6.436L0 9.082V0h9.082L5.711 3.371l6.438 6.436l6.315-6.315L14.968 0h9.082v9.082L20.678 5.71l-6.315 6.315l6.266 6.265l-2.218 2.218l-6.262-6.265l-6.389 6.386z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowH(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.235 13.234H4.765v3.494L0 11.964l4.765-4.765v3.494h14.47V7.199L24 11.964l-4.765 4.765z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeft(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.965 1.6l.035.034l-8.832 8.855H24v3.556H7.166L16 22.9l-.035.034h-5.319L0 12.268L10.647 1.601z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowLeftL(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24 10.64H5.76V6.4L0 12.16l5.76 5.76v-4.24H24z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowMove(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.566 4.541L12.025 0L7.484 4.541zM7.484 19.459L12.025 24l4.539-4.541z"/><path fill="currentColor" d="M13.194 4.49v14.969h-2.218V4.49zM4.541 7.485L0 12.026l4.541 4.541zm14.918 9.081L24 12.025l-4.541-4.541z"/><path fill="currentColor" d="M4.49 10.856h14.969v2.218H4.49z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowResize(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.327 7.418L7.418 18.327l-1.745-1.745L16.582 5.673l-2.4-2.4h6.545l.001 6.546zM3.272 20.727v-6.545l6.546 6.546z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowReturnLeft(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.935 20.822H24V7.414H5.746V3.2L0 8.946l5.746 5.746v-4.213h15.189v7.28H2.988v3.065z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowReturnRight(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.065 20.822H0V7.414h18.254V3.2L24 8.946l-5.746 5.746v-4.213H3.065v7.28h17.946v3.065z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRight(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.034 22.933L8 22.899l10.608-10.634L8 1.632l.034-.034h5.319L24 12.264L13.353 22.931z"/><path fill="currentColor" d="M0 10.488h18.666v3.555H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowRightL(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24 12.16L18.24 6.4v4.24H0v3.04h18.24v4.24z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowSwap(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 16.16l5.746-5.746v4.214H24v3.065H5.746v4.214zm18.254-7.28H0V5.814h18.254V1.6L24 7.346l-5.746 5.746z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUp(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.934 15.966L22.9 16L12.267 5.392L1.633 16l-.034-.034v-5.319L12.266 0l10.666 10.647z"/><path fill="currentColor" d="M10.49 24V5.334h3.555V24z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowUpL(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.92 5.76L12.16 0L6.4 5.76h4.24V24h3.04V5.76z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ArrowV(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.3 19.5v-15h2.4v15H17L12.5 24L8 19.5zM12.5 0L17 4.5H8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Asterisk(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.416 15.661L14.75 12l6.666-3.661a1.127 1.127 0 0 0 .43-1.554l.003.005l-.913-1.584a1.126 1.126 0 0 0-1.562-.397l.005-.003l-6.503 3.943l.163-7.603v-.022c0-.621-.503-1.125-1.124-1.125h-1.826c-.621 0-1.125.504-1.125 1.125v.025v-.001l.163 7.603l-6.504-3.942a1.124 1.124 0 0 0-1.554.394l-.003.005l-.913 1.581A1.123 1.123 0 0 0 .58 8.335l.006.003l6.666 3.661L.586 15.66a1.127 1.127 0 0 0-.43 1.554l-.003-.005l.913 1.582a1.126 1.126 0 0 0 1.562.397l-.005.003l6.503-3.943l-.163 7.603v.024c0 .621.504 1.125 1.125 1.125h1.826c.621 0 1.125-.504 1.125-1.125v-.025v.001l-.163-7.603l6.503 3.943a1.124 1.124 0 0 0 1.554-.394l.003-.005l.913-1.582a1.123 1.123 0 0 0-.427-1.546z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func At(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.596 20.093a11.435 11.435 0 0 1-4.48 2.943l-.08.025c-1.634.594-3.52.938-5.486.938h-.1h.005l-.189.001c-1.766 0-3.456-.327-5.012-.923l.096.032a11.476 11.476 0 0 1-3.93-2.474l.004.003a11.005 11.005 0 0 1-2.502-3.719l-.026-.073a12.447 12.447 0 0 1-.895-4.676l.001-.152v.008l-.001-.113c0-1.686.361-3.288 1.01-4.733l-.029.073a11.961 11.961 0 0 1 2.669-3.808l.004-.004A12.488 12.488 0 0 1 7.535.938l.084-.03C9.024.334 10.655.001 12.363.001h.098h-.005h.041c1.52 0 2.986.234 4.363.668l-.103-.028a11.174 11.174 0 0 1 3.724 1.951l-.023-.017a9.546 9.546 0 0 1 2.552 3.17l.025.055A9.935 9.935 0 0 1 24 10.364v-.012l.002.187a9.968 9.968 0 0 1-.559 3.301l.021-.07a7.526 7.526 0 0 1-1.443 2.494l.007-.009a5.856 5.856 0 0 1-2.015 1.485l-.037.015a5.973 5.973 0 0 1-2.409.498h-.018h.001a3.301 3.301 0 0 1-2.099-.617l.01.007a1.894 1.894 0 0 1-.782-1.534v-.015v.001h-.16a5.285 5.285 0 0 1-1.48 1.456l-.02.012a4.256 4.256 0 0 1-2.487.697h.007a4.182 4.182 0 0 1-3.418-1.447l-.005-.005a5.676 5.676 0 0 1-1.205-3.79v.013c.001-.96.167-1.88.473-2.735l-.018.057a7.648 7.648 0 0 1 1.312-2.37l-.011.014a6.69 6.69 0 0 1 1.985-1.643l.035-.017a5.369 5.369 0 0 1 2.612-.626h-.004a3.96 3.96 0 0 1 2.109.525l-.019-.01c.503.281.89.717 1.103 1.242l.006.017h.032l.262-1.291h2.903l-1.28 6.101c-.043.3-.102.632-.177 1.002c-.069.31-.11.666-.113 1.032v.002l-.001.063c0 .314.08.61.22.868l-.005-.01a.87.87 0 0 0 .834.369h-.004a2.439 2.439 0 0 0 2.099-1.341l.006-.014a6.858 6.858 0 0 0 .829-3.659l.001.016l.001-.155c0-1.183-.241-2.31-.676-3.335l.021.056a7.078 7.078 0 0 0-1.801-2.513l-.007-.006a7.747 7.747 0 0 0-2.698-1.517l-.055-.015a11.087 11.087 0 0 0-3.378-.515l-.115.001h.006l-.122-.001a9.55 9.55 0 0 0-3.707.744l.063-.024A8.63 8.63 0 0 0 5.715 5.34l-.003.003a9.081 9.081 0 0 0-1.854 2.938l-.021.062a10.524 10.524 0 0 0-.675 3.747c0 1.372.257 2.684.727 3.89l-.025-.073a8.406 8.406 0 0 0 1.968 2.917l.003.003a8.597 8.597 0 0 0 2.972 1.835l.06.019c1.114.406 2.4.641 3.741.641l.144-.001h-.007a11.131 11.131 0 0 0 4.62-.865l-.074.027a12.562 12.562 0 0 0 3.513-2.331l-.006.006zM12.391 8.356h-.038c-.504 0-.968.169-1.338.455l.005-.004a3.938 3.938 0 0 0-.987 1.128l-.01.019a5.782 5.782 0 0 0-.613 1.491l-.009.041a6.486 6.486 0 0 0-.214 1.609v.004c.002.293.031.578.085.854l-.005-.029c.057.306.171.578.331.817l-.005-.008a1.84 1.84 0 0 0 1.686.851h-.006l.078.001a2.53 2.53 0 0 0 1.431-.44l-.009.006a3.522 3.522 0 0 0 1.004-1.082l.009-.016c.248-.415.446-.895.567-1.405l.007-.035c.106-.421.171-.907.179-1.407v-.005c0-.36-.035-.712-.102-1.052l.006.034a2.7 2.7 0 0 0-.35-.917l.007.012a1.932 1.932 0 0 0-.661-.654l-.009-.005a1.926 1.926 0 0 0-.977-.262l-.066.001h.003z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Atlassian(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.353 23.99h-3.609a1.265 1.265 0 0 1-1.312-.785l-.003-.009c-1.335-2.671-2.787-5.301-3.996-8.065c-.895-1.885-1.417-4.095-1.417-6.428c0-2.992.86-5.783 2.345-8.14l-.037.063c.166-.282.521-.626.793-.626s.594.376.742.668q5.457 10.916 10.926 21.84c.543 1.043.293 1.471-.907 1.483zm-14.608 0H1.039c-1.042 0-1.273-.418-.814-1.336q2.807-5.551 5.562-11.081c.501-1.043.897-1.043 1.596-.177a12.19 12.19 0 0 1 2.948 7.971c0 1.148-.158 2.258-.453 3.312l.021-.086c-.345 1.325-.533 1.408-1.868 1.408z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Atom(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.819 12.629q-.132.263-.237.487t-.211.461q-.527 1.133-.948 2.304a14.97 14.97 0 0 0-.642 2.308l-.017.101q-.105.685-.145 1.37a4.476 4.476 0 0 0 .153 1.427l-.007-.032c.065.348.234.649.473.877l.001.001c.233.2.536.324.868.329h.001a1.6 1.6 0 0 0 .511-.028l-.01.002l.474-.105c.447-.13.835-.301 1.196-.514l-.024.013q.54-.316 1.067-.685a1.14 1.14 0 0 1 .255-.116l.008-.002a.468.468 0 0 1 .292.014l-.003-.001a.53.53 0 0 1 .351.392l.001.003a.437.437 0 0 1-.17.473l-.001.001c-.512.389-1.089.756-1.697 1.074l-.068.032c-.59.306-1.286.49-2.025.5h-.003a2.383 2.383 0 0 1-1.534-.444l.007.005a2.617 2.617 0 0 1-.917-1.285l-.005-.018a4.281 4.281 0 0 1-.263-1.307v-.009a8.225 8.225 0 0 1 .056-1.357l-.004.04c.169-1.402.498-2.678.972-3.881l-.037.106a31.92 31.92 0 0 1 1.706-3.711l-.086.174a.654.654 0 0 0 .079-.289v-.001a.904.904 0 0 0-.055-.296l.002.006q-.237-.658-.448-1.317l-.421-1.317a.245.245 0 0 0-.087-.175a.274.274 0 0 0-.199-.013l.002-.001q-.895.211-1.765.487a8.684 8.684 0 0 0-1.702.747l.042-.022a4.733 4.733 0 0 0-.601.415l.009-.007a7.14 7.14 0 0 0-.54.487a1.412 1.412 0 0 0-.421.89v.043c0 .345.119.663.319.913l-.002-.003c.188.293.415.541.678.746l.007.005q.395.303.816.566q.158.079.29.158l.263.158a.622.622 0 0 1 .271.325l.001.004a.45.45 0 0 1-.041.384l.001-.002a.477.477 0 0 1-.313.263l-.003.001a.6.6 0 0 1-.425-.054l.003.002a7.699 7.699 0 0 1-1.172-.696l.023.016a6.14 6.14 0 0 1-.996-.89l-.005-.006a3.74 3.74 0 0 1-.464-.682l-.01-.02a2.75 2.75 0 0 1-.261-.76l-.003-.017a2.253 2.253 0 0 1 .069-1.011l-.004.016a2.89 2.89 0 0 1 .466-.902l-.005.007c.215-.285.458-.532.729-.744l.009-.007c.261-.206.554-.394.865-.552l.03-.014a10.493 10.493 0 0 1 1.648-.716l.077-.023q.878-.29 1.804-.5c.039-.014.087-.027.137-.038l.008-.001c.058-.012.106-.026.153-.042l-.009.003q-.053-.237-.088-.474t-.066-.474a16.041 16.041 0 0 1-.196-1.94l-.001-.035a6.4 6.4 0 0 1 .235-1.994l-.011.045c.057-.285.139-.536.246-.773l-.009.022a2.83 2.83 0 0 1 .425-.676l-.003.004a2.358 2.358 0 0 1 1.177-.77l.017-.004a2.958 2.958 0 0 1 1.455.018L9.166.096a4.876 4.876 0 0 1 1.426.591l-.021-.012c.481.296.889.583 1.279.891l-.028-.021c.1.093.165.223.175.367v.002l.001.032a.461.461 0 0 1-.146.337a.502.502 0 0 1-.327.184h-.003a.524.524 0 0 1-.383-.106l.001.001q-.316-.211-.606-.421a6.835 6.835 0 0 0-.572-.377l-.033-.018a3.586 3.586 0 0 0-.699-.335l-.026-.008a2.463 2.463 0 0 0-.806-.132h-.025h.001h-.014c-.219 0-.426.054-.608.148l.007-.003a1.228 1.228 0 0 0-.458.429l-.003.005a2.042 2.042 0 0 0-.365.736L6.93 2.4a6.13 6.13 0 0 0-.155.797l-.003.033a8.488 8.488 0 0 0-.01 1.84l-.003-.036c.07.688.168 1.294.297 1.889l-.02-.111v.009a.14.14 0 0 0 .053.109c.034.025.077.04.124.04h.008q.737-.079 1.449-.145t1.422-.119a.958.958 0 0 0 .269-.081l-.006.002a.66.66 0 0 0 .209-.183l.001-.002a28.599 28.599 0 0 1 1.957-2.26l-.008.009a17.814 17.814 0 0 1 2.191-1.929l.047-.034a14.02 14.02 0 0 1 1.305-.886l.065-.036a6.445 6.445 0 0 1 1.481-.647l.047-.012a3.893 3.893 0 0 1 1.004-.184h.01a2.699 2.699 0 0 1 1.059.164l-.019-.006c.404.127.744.361.998.668l.003.004c.244.309.434.674.548 1.071l.005.022c.134.401.211.863.211 1.342v.012c0 .469-.029.932-.084 1.386l.005-.054q-.026.184-.053.351c-.022.138-.05.256-.084.372l.005-.021a.451.451 0 0 1-.196.328l-.002.001a.525.525 0 0 1-.385.088h.003a.514.514 0 0 1-.342-.183l-.001-.001a.505.505 0 0 1-.105-.398v.003q.026-.316.079-.645c.033-.194.052-.417.053-.645a5.79 5.79 0 0 0-.001-.795l.001.018a6.108 6.108 0 0 0-.112-.816l.007.04a1.5 1.5 0 0 0-.679-1.063l-.006-.004a1.782 1.782 0 0 0-1.33-.115l.013-.003a4.771 4.771 0 0 0-1.16.408l.027-.013q-.527.263-1.053.579a14.8 14.8 0 0 0-1.93 1.534l.007-.007q-.895.843-1.712 1.765q-.158.158-.316.342t-.29.342v.004a.052.052 0 0 1-.013.035a.101.101 0 0 0-.013.066v-.001h.29c1.496.019 2.947.124 4.373.31l-.186-.02c1.568.203 2.967.507 4.32.918l-.186-.049c.762.213 1.387.433 1.996.687l-.112-.042c.702.289 1.296.594 1.86.941l-.056-.032c.324.211.603.419.869.643l-.013-.011c.276.232.521.483.743.754l.008.01c.394.469.633 1.079.633 1.745l-.001.089v-.004a2.618 2.618 0 0 1-.841 1.723l-.002.002a5.343 5.343 0 0 1-1.133.864l-.026.014a8.954 8.954 0 0 1-1.278.593l-.065.021a.534.534 0 0 1-.372-.028l.003.001a.483.483 0 0 1-.262-.286l-.001-.003a.473.473 0 0 1 .001-.385l-.001.003a.573.573 0 0 1 .286-.301l.003-.001q.316-.158.645-.303a4.52 4.52 0 0 0 .633-.34l-.018.011c.185-.111.344-.224.494-.348l-.007.005q.224-.184.439-.395c.249-.24.414-.566.447-.929v-.014c0-.364-.128-.699-.342-.961l.002.003a3.422 3.422 0 0 0-.609-.654l-.006-.005a6.249 6.249 0 0 0-.722-.511l-.028-.016a10.552 10.552 0 0 0-1.915-.976l-.074-.025a19.122 19.122 0 0 0-1.965-.648l-.155-.036q-1.133-.29-2.277-.5a18.83 18.83 0 0 0-2.239-.286l-.065-.004q-.922-.053-1.857-.088t-1.857-.066h-.01a.69.69 0 0 0-.232.041l.005-.002a.283.283 0 0 0-.157.143l-.001.002L9.66 9.57q-.553.843-1.08 1.686a.247.247 0 0 0-.066.144v.001a.348.348 0 0 0 .013.174l-.001-.003q.474 1.08.988 2.133t1.119 2.067t1.277 1.975t1.435 1.856c.448.519.9.991 1.374 1.44l.009.008c.465.441.977.84 1.526 1.187l.041.024c.238.161.511.311.797.435l.033.013c.267.117.577.192.902.21h.007a1.392 1.392 0 0 0 .785-.162l-.007.004c.228-.133.412-.321.536-.546l.004-.007c.131-.208.235-.45.299-.707l.004-.017a6.63 6.63 0 0 0 .141-.767l.004-.036a10.322 10.322 0 0 0-.114-2.602l.009.061q-.211-1.251-.527-2.488q-.369-1.264-.843-2.475t-1.027-2.398a.819.819 0 0 1-.086-.192l-.002-.006a.304.304 0 0 1-.012-.2v-.004c0-.122.05-.232.132-.31a.541.541 0 0 1 .313-.158h.003a.511.511 0 0 1 .319.054l-.003-.001a.507.507 0 0 1 .236.26l.001.003q.184.421.395.856t.395.878q.553 1.343 1.001 2.738c.265.788.503 1.745.669 2.728l.016.116c.095.575.17 1.277.208 1.99l.002.052a7.372 7.372 0 0 1-.195 2.092l.01-.05q-.053.211-.132.439a3.743 3.743 0 0 1-.194.461l.01-.022c-.233.545-.638.98-1.145 1.244l-.014.007c-.319.157-.694.25-1.09.25c-.221 0-.435-.029-.639-.082l.017.004c-.53-.11-1-.297-1.425-.551l.021.012a11.699 11.699 0 0 1-1.276-.875l.025.019a15.09 15.09 0 0 1-2.013-1.906l-.015-.017q-.922-1.053-1.738-2.185a28.32 28.32 0 0 1-1.619-2.536l-.076-.145q-.777-1.383-1.435-2.857l-.053-.105zm-.526-4.636l.658 2.16q.395-.632.777-1.212t.803-1.185zm5.977 5.531h-.018c-.42 0-.801-.171-1.075-.448a1.362 1.362 0 0 1-.437-1.001l.002-.083v-.019c0-.415.172-.789.448-1.057a1.468 1.468 0 0 1 1.057-.448h.024h-.001h.018c.42 0 .801.171 1.075.448c.284.265.461.642.461 1.061v.02v-.001v.019c0 .418-.177.795-.46 1.06l-.001.001a1.51 1.51 0 0 1-1.075.448h-.015h.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func AudioDescription(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.32 16.326H7.04V14.96H3.996l-.872 1.366H0l.189-.255l3.274-4.548l3.091-4.302l1.386-.006h1.38zm-2.181-6.08l-.996 1.46l-.995 1.464l.989.019c.225 0 .434.005.605.005a3.702 3.702 0 0 0 .424-.011l-.014.001c.012-.012.014-.566.006-1.483zm4.836 6.08h-1.874V7.215h1.939a13.496 13.496 0 0 1 2.419.106l-.069-.008a3.666 3.666 0 0 1 2.49 2.028l.009.023a4.464 4.464 0 0 1 .526 2.455l.001-.014a4.222 4.222 0 0 1-.191 1.527l.008-.03a4.024 4.024 0 0 1-1.979 2.586l-.021.01c-.781.381-1.146.429-3.258.429zm.274-6.919v4.654l.866-.026c.994-.033 1.238-.119 1.634-.579a2.282 2.282 0 0 0 .52-1.729l.001.01a1.943 1.943 0 0 0-.181-1.118l.005.011c-.428-.921-.79-1.132-2.037-1.191l-.81-.032zm10.329 6.913c-.096 0-.141-.009-.155-.03s0-.042.01-.074c.506-1.32.803-2.846.814-4.44v-.005a13.27 13.27 0 0 0-.844-4.537l.03.092c-.014-.035-.016-.055-.006-.07c.023-.034.114-.041.357-.041c.314 0 .405.024.455.124c.239.654.445 1.435.581 2.239l.011.081c.107.635.168 1.366.168 2.111c0 1.419-.221 2.785-.63 4.068l.026-.095l-.182.553l-.416.02c-.103.002-.169.005-.219.005zm-4.232 0c-.055 0-.115 0-.195-.006l-.436-.02l.234-.351a7.383 7.383 0 0 0 1.134-3.079l.004-.038a10.77 10.77 0 0 0-.138-2.947l.01.066a7.853 7.853 0 0 0-1.171-2.614l.017.026l-.104-.143h.924l.176.274a8.4 8.4 0 0 1 .812 1.95l.014.061c.188.676.295 1.453.295 2.255a8.719 8.719 0 0 1-.871 3.811l.023-.052c-.344.7-.4.807-.73.807zm2.024 0a.581.581 0 0 1-.252-.026l.004.001c0-.024.135-.309.195-.43c.13-.26.262-.577.374-.904l.017-.056c.296-.922.466-1.983.466-3.084c0-1.497-.315-2.92-.883-4.207l.026.067a5.112 5.112 0 0 1-.195-.43a.585.585 0 0 1 .25-.027h-.003c.06 0 .124 0 .188.006l.436.019l.247.615c.442 1.159.698 2.5.698 3.9c0 1.465-.28 2.864-.789 4.148l.027-.076l-.18.453l-.436.02c-.064 0-.127.005-.188.005z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Automobile(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.031 15.388a.863.863 0 1 1 .002-1.726a.863.863 0 0 1-.002 1.726m0-2.88a2.014 2.014 0 1 0 .003 4.029a2.014 2.014 0 0 0-.004-4.029zm-14.137 2.88a.864.864 0 0 1 0-1.726a.863.863 0 0 1 0 1.726m0-2.88a2.014 2.014 0 1 0 2.014 2.014v-.001a2.014 2.014 0 0 0-2.014-2.014h-.002z"/><path fill="currentColor" d="m11.507 9.727l-.143-2.64h2.202l3.023 2.64zm-.854 0H7.191l1.207-2.64h2.254zm-4.695 0H4.894l2.014-2.64h.64zm17.46 2.16l-.18-.578a1.682 1.682 0 0 0-1.312-1.151l-.01-.001l-4.262-.71l-3.108-2.646a1.681 1.681 0 0 0-1.096-.404H7.206c-.585 0-1.103.288-1.42.729l-.004.005l-1.854 2.607l-1.954.2a1.283 1.283 0 0 0-1.112.973l-.002.009l-.292 1.259H.565a.558.558 0 0 0-.454.232l-.001.002a.556.556 0 0 0-.079.512l-.001-.004l.18.539a1.558 1.558 0 0 0 1.474 1.062h.554a2.656 2.656 0 1 1 5.311 0h8.826a2.656 2.656 0 0 1 2.656-2.656h.001a2.656 2.656 0 0 1 2.656 2.656h.458c.769 0 1.419-.509 1.631-1.209l.003-.012l.198-.666a.585.585 0 0 0-.093-.515l.001.001a.586.586 0 0 0-.467-.234z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Aws(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.248 6.685c-.656 2.694-.614 2.55-1.209 5.014a.43.43 0 0 1-.504.4h.003h-.675a.4.4 0 0 1-.463-.34v-.002c-.577-1.846-1.35-4.41-1.87-6.125c-.192-.633-.066-.618.436-.605c.174.005.342 0 .515 0a.39.39 0 0 1 .426.331v.002c.169.605.281 1.069 1.246 4.88c.018.075.042.146.066.215h.051c.023-.094.051-.182.075-.277q.548-2.313 1.092-4.63c.113-.48.314-.525.8-.525h.356c.32.005.422.07.502.389c.281 1.097 1.102 4.77 1.251 5.174c.24-.858-.084.37 1.336-5.108c.098-.379.192-.454.577-.454h.591c.253.005.328.084.267.333c-.113.446-.136.464-1.935 6.228c-.146.464-.197.506-.684.506h-.497c-.342 0-.431-.061-.515-.394c-.202-.759-1.092-4.485-1.238-5.01zm-6.364 5.211a.411.411 0 0 0 .661.112l.295-.197c.32-.211.342-.295.169-.633a2.626 2.626 0 0 1-.296-1.218v-.045v.002c0-.146.028-2.61-.042-3.13A2.03 2.03 0 0 0 5.2 4.976l-.014-.003a4.457 4.457 0 0 0-1.552-.139l.016-.001a4.67 4.67 0 0 0-2.025.547l.025-.012a.4.4 0 0 0-.191.227l-.001.003a1.86 1.86 0 0 0-.033.575l-.001-.008c.028.277.122.328.384.24c.24-.08.469-.182.708-.253a3.442 3.442 0 0 1 2.09-.074l-.024-.006c.317.095.564.334.668.639l.002.007a3.045 3.045 0 0 1 .11 1.141l.001-.012c0 .258-.005.258-.258.211a6.012 6.012 0 0 0-1.987-.144l.024-.002a2.53 2.53 0 0 0-1.787.933l-.004.005c-.26.36-.416.811-.416 1.298c0 .197.025.387.073.569l-.003-.016a1.761 1.761 0 0 0 1.225 1.422l.012.003a2.867 2.867 0 0 0 2.71-.45l-.006.005c.178-.131.338-.29.52-.446c.146.234.272.454.417.656zm-.72-2.887h.005a.201.201 0 0 1 .197.242V9.25c-.01.178.005.356-.014.534c-.05.523-.377.959-.83 1.164l-.009.004a2.181 2.181 0 0 1-1.232.234l.009.001a1.003 1.003 0 0 1-.917-1.24l-.001.007c.026-.548.46-.987 1.004-1.021h.003a4.642 4.642 0 0 1 1.824.085l-.032-.007zM21.51 12.06a2.051 2.051 0 0 0 1.367-2.357l.002.013a1.748 1.748 0 0 0-1.169-1.379l-.012-.004a20.053 20.053 0 0 1-2.383-.905l.129.053a.64.64 0 0 1-.356-.559v-.001a.81.81 0 0 1 .717-.927h.003a3.24 3.24 0 0 1 1.337-.015l-.02-.003c.342.066.67.197 1.003.295c.131.042.277.098.366-.075a.781.781 0 0 0-.338-1.019l-.004-.002a3.929 3.929 0 0 0-1.686-.373c-.546 0-1.066.11-1.54.308l.026-.01a1.93 1.93 0 0 0-.101 3.552l.012.005c.361.174.754.277 1.134.417s.759.272 1.13.426a.822.822 0 0 1 .059 1.404l-.003.002c-1.298.72-3.04-.113-3.243-.178c-.154-.051-.248.01-.295.174a.782.782 0 0 0 .539 1.063l.005.001a4.842 4.842 0 0 0 1.79.336c.546 0 1.072-.089 1.563-.253l-.035.01zm-7.634 7.625a16.544 16.544 0 0 0 6.763-2.244l-.076.042c.375-.22.726-.48 1.082-.736c.342-.24.15-.862-.53-.572c-2.344 1.029-5.066 1.687-7.925 1.821l-.051.002a22.28 22.28 0 0 1-4.393-.235l.123.016C5.752 17.33 2.945 16.321.44 14.852l.109.059a1.094 1.094 0 0 0-.297-.139l-.008-.002c-.23-.051-.361.22-.113.454a17.206 17.206 0 0 0 11.716 4.578h.014h-.001c.675-.038 1.35-.042 2.015-.117zm8.187-4.32c.694.038.909.277.736.946c-.178.694-.436 1.369-.651 2.053c-.042.136-.197.295-.038.413c.174.122.305-.047.422-.154a4.301 4.301 0 0 0 1.044-1.555l.01-.029a4.637 4.637 0 0 0 .408-1.919v-.047v.002c-.01-.29-.084-.413-.366-.492a3.891 3.891 0 0 0-.757-.148l-.016-.001a5.207 5.207 0 0 0-2.936.51l.03-.014a2.4 2.4 0 0 0-.412.28l.004-.003c-.051.042-.15.248.113.286a1.116 1.116 0 0 0 .283-.015l-.007.001a12.68 12.68 0 0 1 2.151-.112h-.019z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Babel(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.986 9.23v.024h.071v-.032q0-.012.024-.036v-.048a.068.068 0 0 0-.048.024a.097.097 0 0 0-.024.064v.007zm-.548-1.36v-.048h-.095v.048zm-1.623 3.054a3.416 3.416 0 0 0-.219-.292l.004.005l-.24-.286q0-.048.071-.107a1.84 1.84 0 0 1 .231-.151l.009-.005l.573-.501c.073-.113.134-.244.176-.382l.003-.011c.038-.128.06-.276.06-.428v-.014v.001v-.095a.5.5 0 0 0-.143-.346a1.216 1.216 0 0 0-.421-.272l-.008-.003a2.247 2.247 0 0 0-1.151-.239h.006a5.227 5.227 0 0 0-.991.188l.037-.009c-.349.101-.641.21-.923.338l.043-.018q-.119.095-.274.202a2.72 2.72 0 0 1-.329.195l-.016.007l-.024.048h.034q.012 0 .036-.024h.004a.05.05 0 0 1 .032.012c.008.009.012.02.012.032v.004l.024-.024h.024v.024q-.119.095-.226.191a2.246 2.246 0 0 1-.242.186l-.008.005l.024.048h-.024l-.048-.024l-.024.012a.106.106 0 0 1-.048.012v.024l.048.071h-.004a.05.05 0 0 1-.032-.012a.047.047 0 0 0-.032-.012h-.006a.24.24 0 0 0-.13.037l.001-.001a.2.2 0 0 0-.079.079l-.001.001l.024-.119c.044-.153.085-.347.116-.544l.004-.029l.095-.62v-.095a.763.763 0 0 1-.119-.095l-.095-.095v-.119h-.024q-.334.525-1.336 1.813q-.763.906-.811 1.002q-.143.167-.214.262a.338.338 0 0 0-.07.117l-.001.002a.129.129 0 0 0-.06.048a.16.16 0 0 0-.012.063v.009q-.024 0-.024-.012v-.012a.946.946 0 0 1-.177.175l-.002.002a.545.545 0 0 1-.223.107l-.004.001l-.191.024q-.024.024-.024.036v.06l.048-.024v.024l-.143.048l-.286.071h-.148q0 .024.06.012t.012.012l-.191.024a.262.262 0 0 1-.133.012h.002a.158.158 0 0 1-.06-.036h-.024l-.167.048q.024-.095-.549-.72a.226.226 0 0 1 .094-.155h.001c.091-.073.192-.14.3-.197l.01-.005l.787-.692c.195-.299.31-.665.31-1.058v-.017v.001L7.375 8.6a.767.767 0 0 0-.215-.466a1.786 1.786 0 0 0-.537-.365l-.011-.004a1.747 1.747 0 0 0-.574-.225l-.011-.002a4.102 4.102 0 0 0-.826-.081l-.097.001h.005c-.498.07-.931.162-1.353.28l.073-.017a8.543 8.543 0 0 0-1.212.427l.055-.022a3.018 3.018 0 0 1-.344.243l-.014.008q-.214.131-.454.274l-.048.071a.114.114 0 0 0 .049-.012H1.86a.106.106 0 0 1 .048-.012h.028a.05.05 0 0 1 .032.012c.007.009.012.02.012.032v.002l.048-.048h.048v.048a3.073 3.073 0 0 1-.291.257l-.007.005a2.696 2.696 0 0 1-.307.207l-.013.007l.024.071h-.024l-.071-.024v.004a.02.02 0 0 1-.021.021h-.004a.186.186 0 0 0-.072.024h.001v.024l.071.071h-.048a.068.068 0 0 1-.048-.024a.82.82 0 0 0-.242.121l.002-.002a.677.677 0 0 0-.166.188l-.002.003l.024.048a.873.873 0 0 1 .093-.078l.002-.002a.423.423 0 0 1 .116-.059l.003-.001v.071a.114.114 0 0 0-.049.012h.001A.215.215 0 0 0 1 9.724l.071.095c.066-.066.136-.129.209-.187l.005-.004a2.19 2.19 0 0 1 .23-.162l.01-.005h.013c.035 0 .067.009.096.025l-.001-.001l.036.024h.071c.179-.135.382-.266.594-.381l.026-.013c.184-.101.412-.209.646-.304l.046-.016l.048.048q-.119.191-.167.191v.009c0 .023.004.044.013.064v-.001c.01.018.022.034.036.048a12.533 12.533 0 0 1-.651 1.633l.033-.076a42.887 42.887 0 0 1-2.475 5.106l.113-.215c0 .018.005.034.012.049v-.001a.04.04 0 0 0 .036.024a.402.402 0 0 0 .156-.037l-.002.001a.428.428 0 0 0 .132-.08l.048.024v.048h.048l.071-.024l.012.012c.009.008.02.012.032.012h.004v.048a.439.439 0 0 1-.025.146l.001-.003a.58.58 0 0 1-.073.145l.001-.002a1.058 1.058 0 0 0-.117.196l-.003.007a1.07 1.07 0 0 0-.07.22l-.001.007v.024h.048c.114-.115.219-.24.314-.372l.006-.009c.092-.124.176-.265.244-.415l.006-.015c.386-.109.7-.218 1.005-.343l-.062.023q.442-.179.847-.37c.158-.001.31-.027.452-.074l-.01.003c.16-.051.298-.124.421-.217l-.003.003v-.048l-.167.048h-.024v-.024c.122-.018.23-.042.335-.075l-.015.004c.116-.036.215-.076.309-.124l-.011.005a13.71 13.71 0 0 1 1.213-.898l.052-.032c.494-.323.89-.758 1.16-1.269l.009-.019l.095-.024h.005c.012 0 .022.004.031.011c.009.008.02.012.032.012h.008a.05.05 0 0 0 .032-.012a.047.047 0 0 1 .032-.012h.004v-.012q0-.012.024-.012h.143l.24-.048h.095l.024.071l-1.025 1.613l-.24.262l-.119.358h.024l.119-.048l-.048.143l.048.096h-.01q-.012 0-.012.024v.024l.048-.024q.119-.071.095-.143h.143v.024q-.024 0-.048.036a.145.145 0 0 0-.024.08h.024c.306-.339.592-.71.849-1.103l.021-.035q.585-.859 1.134-1.574q-.024-.143 1.145-.406h.024a2.13 2.13 0 0 1-.197.893l.006-.013a.803.803 0 0 0-.07.209l-.001.005l-.048.24a1.267 1.267 0 0 0-.105.289l-.002.009q-.036.155-.08.298l-.262.93q-.119.31-.226.668a6.456 6.456 0 0 0-.172.694l-.007.045v.004a.05.05 0 0 0 .012.032c.009.008.02.012.032.012h.028l.048-.024h.024l-.024-.048h.124a.056.056 0 0 0 .055-.048a.474.474 0 0 1 .037-.098l-.001.003v-.018c0-.045.009-.088.025-.127l-.001.002a.277.277 0 0 1 .048-.071c.018-.09.039-.166.064-.24l-.004.013q.036-.107.08-.226c.02-.097.045-.181.075-.262l-.004.012c.033-.092.065-.167.101-.239l-.006.013c.039-.169.084-.313.139-.452l-.008.022q.06-.143.08-.143l.024.024l-.071.191q-.095.43-.202.787t-.226.692l.048.095h.024c.49-1.136.979-2.54 1.375-3.983l.056-.239a.627.627 0 0 1 .097-.229l-.001.002a.74.74 0 0 1 .165-.178l.002-.001l-.024-.048a.73.73 0 0 0 .083-.05l-.003.002a.42.42 0 0 0 .06-.048l-.048-.024a.7.7 0 0 0-.138-.047l-.005-.001l.31-1.336l.048-.048a.26.26 0 0 1 .07-.047l.002-.001v.071h-.004a.05.05 0 0 0-.032.012a.047.047 0 0 1-.032.012h-.004l.048.095q.095-.095.179-.167a.869.869 0 0 1 .174-.117l.005-.002h.007c.025 0 .047.009.064.024l.024.024h.048c.14-.108.297-.211.462-.301l.018-.009c.146-.083.319-.164.498-.231l.026-.009l.048.048q-.095.143-.143.143c0 .018.005.034.012.049l.012.023a5.671 5.671 0 0 1-.228.695l.014-.039q-.119.298-.262.585q-.43 1.074-.894 2t-.966 1.84a.574.574 0 0 1 .081-.12l-.001.001q.009 0 .036.024a.985.985 0 0 0 .125-.05l-.005.002l.095-.048h.024v.048h.048l.048-.024q0 .024.048.024v.048a.193.193 0 0 1-.025.096l.001-.001l-.048.095q-.048.071-.095.155a.68.68 0 0 0-.069.178l-.001.005v.048h.048q.119-.143.24-.298c.077-.099.141-.214.188-.336l.003-.009c.286-.07.527-.148.759-.24l-.039.014q.334-.131.644-.274c.12-.001.234-.023.34-.062l-.007.002c.119-.044.222-.104.312-.18l-.001.001v-.024l-.119.048h-.024v-.024a.619.619 0 0 0 .244-.05l-.004.002q.119-.048.214-.095q.334-.262.585-.442t.394-.274c.298-.204.551-.449.757-.731l.006-.009c.134-.181.215-.408.215-.654l-.001-.04v.002zm-.119.048v.17q0 .048-.024.06l-.024.012l-.046-.312l.071.036a.038.038 0 0 1 .023.034M1.217 9.517h-.024v-.048l.119-.024v.024a.068.068 0 0 1-.048.024a.064.064 0 0 0-.047.024M.62 15.505a.54.54 0 0 0-.076-.034l-.005-.001a.252.252 0 0 0-.079-.012H.459l-.018.012l-.012.012v-.098a.111.111 0 0 1 .012-.045v.001q.012-.024.06-.024c.016 0 .03.01.036.024a.16.16 0 0 1 .012.063v.009l.08-.119a.256.256 0 0 1 .079-.071l.001-.001l.071.024a.789.789 0 0 1-.062.196l.002-.005a.115.115 0 0 1-.1.071zm6.394-6.704q-.024 0-.06-.06t-.107-.202v-.071q.048 0 .095.071a.549.549 0 0 1 .071.212v.003zm-.072-.644l-.024.048l-.877-.358q.88.023.904.31zm-3.673 3.145H3.15v-.048l.119-.024zm-.095-1.908h.024v.095q-.024 0-.048.036a.548.548 0 0 1-.071.08v-.048a.574.574 0 0 0 .07-.078l.001-.002a.114.114 0 0 0 .024-.06zm-.263.62l.024.119h-.048v-.119zm-.143.214h.074v.013a.196.196 0 0 1-.025.096l.001-.001a.08.08 0 0 1-.067.036h-.029a.162.162 0 0 0 .036-.059v-.001a.25.25 0 0 0 .009-.069v-.012v.001zm-.095.31h.048v.024l-.071.095h-.048v-.024a.068.068 0 0 0 .048-.024a.097.097 0 0 0 .024-.064zm-.095.263l-.024.143H2.53l-.024-.143zm-1.004 2.843a.983.983 0 0 1-.062.233l.002-.007q-.036.08-.06.08v-.061c0-.052.013-.101.037-.144l-.001.002a.571.571 0 0 1 .084-.103m-.191.406v.024a.114.114 0 0 1-.012.049v-.001a.04.04 0 0 1-.036.024v-.048a.08.08 0 0 1 .047-.071h.001zm-.048.811h.095v.024a.068.068 0 0 1-.048.024a.068.068 0 0 0-.048.024h-.048v-.004a.05.05 0 0 1 .012-.032zm2.648-.811q-.262.119-.501.25c-.165.09-.306.181-.439.281l.009-.007q-.835.286-.835.358l-.549.226c-.192.077-.35.153-.503.238l.023-.012a.114.114 0 0 1-.06-.024a.42.42 0 0 1-.06-.048v-.002a.24.24 0 0 1 .037-.13l-.001.001a.346.346 0 0 1 .106-.106l.001-.001h.04c.039 0 .077.004.113.013l-.003-.001a.474.474 0 0 1 .098.037l-.003-.001q.095-.048.214-.095c.075-.03.164-.055.255-.07l.007-.001v-.048h-.191a2.22 2.22 0 0 1 .235-.125l.015-.006c.071-.034.162-.07.255-.101l.02-.006h.119v.024a.789.789 0 0 0-.196.062l.005-.002q-.071.036-.071.08v.004a.05.05 0 0 0 .012.032c.009.008.02.012.032.012h.004c.055-.025.102-.052.147-.082l-.003.002l.048-.036c.715-.292 1.327-.62 1.9-1.003l-.039.024v.048q0 .024-.06.08a1.34 1.34 0 0 1-.197.152l-.006.003zm-1.573.524c.078-.02.146-.045.21-.074l-.008.003a.59.59 0 0 0 .179-.119h.004a.05.05 0 0 1 .032.012l.036.036q-.119.048-.226.095t-.202.095H2.41zm2.528-1.956l-.787.62q-.191.143-.418.286a3.876 3.876 0 0 1-.464.252l-.025.01a8.329 8.329 0 0 1-1.703.817l-.06.018h-.024q.024-.167 1.241-2.553c.341-.054.638-.125.926-.215l-.046.012c.329-.1.598-.201.859-.317l-.048.019h.143a.351.351 0 0 0 .098-.013l-.003.001a.358.358 0 0 1 .095-.012h.001c.132 0 .259.022.377.062l-.008-.002a.76.76 0 0 1 .299.179l.024.119c-.069.187-.15.349-.246.499l.006-.01a.576.576 0 0 1-.237.225l-.003.002zm-1.526-.859h.119v.024h-.119zm1.766-.31h.095l.107.071q.036.024.036.048v.024h-.002a.24.24 0 0 1-.13-.037l.001.001a.553.553 0 0 1-.107-.08zm.262-1.407q-.167.119-.334.25a2.513 2.513 0 0 1-.321.219l-.013.007a3.33 3.33 0 0 0-.475.164l.022-.008q-.334.131-.93.394v-.004a.02.02 0 0 0-.021-.021h-.1v-.048c.018-.09.039-.166.064-.24l-.004.013a.748.748 0 0 1 .108-.204l-.001.002c.029-.152.061-.28.102-.405l-.006.023q.048-.143.095-.167l.64-1.455q-.024-.143.406-.214h.095v.095c.095-.024.213-.045.332-.059l.013-.001q.131-.012.179-.012a1.168 1.168 0 0 1 .593.038l-.008-.002c.121.04.21.145.226.273v.002h.048v-.167h.05c.064.032.12.072.167.119a.356.356 0 0 1 .095.165l.001.002v.008a.36.36 0 0 1-.037.161l.001-.002a.57.57 0 0 1-.107.143a.04.04 0 0 1-.036-.024a.16.16 0 0 1-.012-.063v-.009h-.071l-.024.191q-.382.573-.573.597a.957.957 0 0 1-.12.168l.001-.001q-.048.048-.071.048zm.406 1.001h-.009a.17.17 0 0 1-.064-.013h.001l-.071-.036v-.071h.071l.119.071a.114.114 0 0 1-.012.049v-.001a.04.04 0 0 1-.036.024zm.405.501l.071.036a.04.04 0 0 1 .024.036h-.048l-.048.024v-.119zm.024.24v-.024h.071zm.191.214a.351.351 0 0 0 .098-.013l-.003.001l.143-.036v.024l-.262.071v-.048zm.263.12c-.051 0-.1.004-.148.013l.005-.001a.803.803 0 0 1-.143.012v-.024l.24-.071l.095-.024l.024.048q-.048 0-.036.012t-.038.034zm.24-.072h-.028a.05.05 0 0 1-.032-.012l-.014-.012l.071-.024h.024l.012.012c.008.009.012.02.012.032v.004zm.644-.692v-.024l.119-.048v.048H7.61zm.692 4.247v-.071h.048l.012.012c.008.009.012.02.012.032v.028l-.071.024zm.501-4.08h-.071a.147.147 0 0 1-.07.036h-.001q-.048.012-.119.036l-.048.024h-.024l-.012-.012a.047.047 0 0 1-.018-.037l.31-.071l.095-.048l.048.024a.162.162 0 0 1-.059.036h-.001a.188.188 0 0 1-.06.012zm.24-.071v-.024h.004a.05.05 0 0 0 .032-.012l.036-.036l.214-.048v.024q-.095.024-.167.048l-.143.048zm.048-.48a.068.068 0 0 1-.048.024h-.004a.02.02 0 0 0-.021.021a.02.02 0 0 0-.021-.021h-.004a.186.186 0 0 0-.072.024h.001q-.024-.024.107-.071t.394-.119l.048-.024l.024.048c-.073.03-.168.063-.266.09l-.021.005a.668.668 0 0 1-.12.025h-.003zm.62-.668a.845.845 0 0 1-.097.338l.002-.004a.295.295 0 0 1-.213.143h-.002v-.024l-.167.071l-.048.024H9.09l-.024-.071q.167-.262.334-.501c.119-.17.237-.318.362-.459l-.004.005l.048.024v.028c0 .079-.009.155-.025.229l.001-.007a.905.905 0 0 1-.073.211l.002-.005zm1.097-.931h-.024v-.048h.095v.024a.114.114 0 0 0-.049.012h.001zm-.449 4.533a.2.2 0 0 1-.061-.012h.001a.252.252 0 0 0-.079-.012h-.001v-.071a.068.068 0 0 1 .024-.048c.019.001.036.01.048.024a.068.068 0 0 1 .024.048a.73.73 0 0 1 .05-.083l-.002.003a.184.184 0 0 1 .07-.06h.001l.048.024a1.222 1.222 0 0 1-.051.164l.003-.008q-.024.06-.071.06zm4.835-5.082l-.036-.048a.736.736 0 0 1-.078-.162l-.002-.006v-.024q.024 0 .06.048a.396.396 0 0 1 .06.141v.003zm-.048-.5l-.024.048l-.668-.286c.181.006.352.035.515.084l-.014-.004q.168.066.192.16zm-2.787 2.388h-.095v-.024h.095zm-.072-1.432h.024v.071q-.048 0-.095.095v-.048a.235.235 0 0 0 .036-.047v-.001zm-.191.454v.095h-.024v-.095zm-.119.167h.048q0 .119-.048.119h-.024c0-.018.005-.034.012-.049v.001a.157.157 0 0 0 .009-.054l-.001-.019v.001zm-.071.24h.024v.024l-.048.071h-.048v-.024q.048 0 .06-.024a.107.107 0 0 0 .008-.041v-.009zm-.071.215l-.024.095v-.095zm-.768 2.148q-.048.214-.095.24v-.048a.28.28 0 0 1 .025-.109l-.001.002a.154.154 0 0 1 .07-.084h.001zm-.144.311v.028a.05.05 0 0 1-.012.032l-.012.012h-.024v-.024l-.001-.01c0-.011.004-.022.012-.029a.044.044 0 0 1 .031-.013l.007.001zm-.024.62h.048v.024l-.048.024h-.048q0-.048.024-.048zm1.98-.62q-.191.095-.37.191a3.095 3.095 0 0 0-.355.221l.009-.006a3.806 3.806 0 0 0-.491.177l.026-.01q-.155.071-.155.095l-.406.179a3.55 3.55 0 0 0-.376.189l.018-.01a.235.235 0 0 0-.047-.036h-.001a.215.215 0 0 1-.048-.036v-.013c0-.035.009-.067.025-.096l-.001.001a.184.184 0 0 1 .07-.06h.001q.024-.024.036-.024h.013c.028 0 .055.004.081.013l-.002-.001c.023.008.044.02.06.036a.995.995 0 0 1 .172-.069l.007-.002q.08-.024.179-.048l.024-.048h-.143c.051-.033.11-.065.171-.092l.008-.003c.057-.026.124-.051.194-.069l.009-.002l.095-.024v.024a.715.715 0 0 0-.148.05l.005-.002a.165.165 0 0 0-.071.07v.005a.05.05 0 0 0 .012.032c.009.008.02.012.032.012h.004l.107-.071q.036-.024.036-.048a7.382 7.382 0 0 0 1.458-.756l-.026.017v.048l-.06.048c-.048.039-.09.078-.131.119zm-1.192.381h.008a.36.36 0 0 0 .161-.037l-.002.001c.056-.027.103-.054.148-.083l-.005.003h.004a.05.05 0 0 1 .032.012l.012.012a.85.85 0 0 0-.172.062l.005-.002a1.346 1.346 0 0 0-.148.083l.005-.003h-.024v-.048zm1.932-1.479l-.597.48a2.52 2.52 0 0 1-.309.22l-.011.006q-.179.107-.37.202a7.046 7.046 0 0 1-1.285.604l-.051.016h-.024q.024-.119.93-1.933q.358-.071.68-.167c.249-.075.459-.157.66-.252l-.028.012h.241c.098 0 .192.018.278.05l-.006-.002a.9.9 0 0 1 .252.144l-.002-.001l-.024.071a1.543 1.543 0 0 1-.159.377l.004-.007a.482.482 0 0 1-.179.174l-.002.001zm-1.168-.644h.095v.024h-.095zm1.36-.24h.068a.277.277 0 0 1 .071.048a.068.068 0 0 1 .024.048v.024h-.024a.15.15 0 0 1-.081-.024h.001a.184.184 0 0 1-.06-.07v-.001zm.191-1.073a1.47 1.47 0 0 1-.245.199l-.006.003c-.096.061-.178.12-.256.183l.005-.004q-.31.095-.573.191a4.696 4.696 0 0 0-.528.227l.027-.012l-.012-.012a.047.047 0 0 0-.032-.012h-.025v-.048q.024-.095.048-.167a.58.58 0 0 1 .073-.145l-.001.002q.048-.191.08-.298a.415.415 0 0 1 .061-.133l-.001.001l.501-1.098q0-.048.08-.08t.226-.08h.071v.071q.334-.048.406-.048a.875.875 0 0 1 .448.013l-.006-.002a.257.257 0 0 1 .179.201v.002h.048v-.095h.024a.256.256 0 0 1 .143.08a.2.2 0 0 1 .048.131v.004a.35.35 0 0 1-.025.13l.001-.002q-.024.06-.048.107a.114.114 0 0 1-.049-.012h.001q-.024-.012-.024-.06h-.043v.143a1.839 1.839 0 0 1-.252.347l.001-.001a.295.295 0 0 1-.177.107h-.002a1.21 1.21 0 0 1-.108.132l.001-.001a.102.102 0 0 1-.06.033h-.001zm.191.739v-.074h.048l.095.071v.004a.05.05 0 0 1-.012.032l-.012.012a.235.235 0 0 0-.047-.036h-.001a.106.106 0 0 0-.048-.012zm-8.278 4.32l-.048.024l-.024.024q0 .024.012.024h.04a.05.05 0 0 0 .032-.012q.012-.012-.012-.036zm-.168-.358v.024h.004a.05.05 0 0 0 .032-.012q.012-.012-.012-.036v-.024l-.013.011a.047.047 0 0 0-.012.032zm1.265-2.361l.024-.024h-.024zm0-.597l.048-.024l-.024-.048l-.048.024v.024zm12.547 1.785l-.119-.024v.071l.72-.071v-.024l-.573.048zm3.697-.453v.048h.071v-.024zm-.119 0a2.184 2.184 0 0 1-.333.07l-.012.001l-.37.048v.024h.071l.048-.024q.406-.024.501-.036t.095-.06zm-3.245-.787l-.048-.024a.418.418 0 0 0-.035.092l-.001.003a.358.358 0 0 0-.012.095v.024q.024 0 .048-.048a.67.67 0 0 0 .049-.137l.001-.005zm-.644-.358v-.024h-.024l-.048.048l.024.024h.004a.05.05 0 0 0 .032-.012a.052.052 0 0 0 .014-.034v-.002zm1.407 1.479v-.048h-.024v.048zm2.935-.668h.028a.05.05 0 0 0 .032-.012a.047.047 0 0 0 .012-.032v-.004l-.024-.048h-.014a.427.427 0 0 0-.144.025l.003-.001q-.06.024-.131.048v-.024l-.143.024l.001-.012a.096.096 0 0 0-.013-.048a.047.047 0 0 0-.032-.012h-.183q-.036 0-.06.048v-.024h-.095l-.048.024q0-.024-.012-.024h-.007l-.024.024v-.024h-.167v.048h.119q.024.024.036.024h.036l.048-.024v.024q-.095.024-.214.048a1.197 1.197 0 0 1-.24.024h-.052a.05.05 0 0 1-.032-.012a.047.047 0 0 1-.012-.032v-.004h.167l.071-.036a.106.106 0 0 1 .048-.012v-.048q-.119.024-.24.036t-.264.012h-.024v-.048c.1-.021.215-.034.332-.036h.001q0 .012.119-.036v.024h.024l.012-.012a.047.047 0 0 1 .032-.012h.104c.078 0 .155-.004.231-.013l-.009.001a1.62 1.62 0 0 0 .213-.038l-.011.002v.024l.048-.024h.048l-.001-.008c0-.024.015-.044.036-.052a.188.188 0 0 1 .06-.012v-.048c-.163.046-.354.08-.55.095l-.01.001h-.015c-.092 0-.181.009-.268.026l.009-.001h-.191q-.024 0-.024.012t-.024.012v-.012q0-.012-.024-.012l-.454.095v-.024l-.24.048a1.326 1.326 0 0 1-.253.024h-.01h.001v.004a.02.02 0 0 1-.021.021h-.008a.02.02 0 0 0-.021.021v.004q0-.024-.012-.024h-.202l-.143.048v-.024h-.024a.2.2 0 0 1-.095.024a.193.193 0 0 0-.096.025l.001-.001q-.024 0 0-.012t0-.012l-.048.024h-.214a1.893 1.893 0 0 0-.299.038l.013-.002a.106.106 0 0 1-.048.012h-.143v.024l-.179-.035q-.06-.012-.06-.036q.048-.214.119-.214h.024l-.024.048h.024a.235.235 0 0 1 .047-.036h.001a.04.04 0 0 0 .024-.036v-.048h-.048v-.143l.119-.167v.048h.048v-.048l-.048-.024a.235.235 0 0 1 .047-.036h.001a.215.215 0 0 0 .048-.036v-.024h-.071v-.048a1.53 1.53 0 0 0 .115-.229l.004-.011c.028-.068.053-.148.07-.231l.002-.009h.052a.05.05 0 0 0 .032-.012a.047.047 0 0 0 .012-.032v-.004h-.048v-.052a.05.05 0 0 1 .012-.032a.048.048 0 0 0 .016-.026a.312.312 0 0 0 .07-.105l.001-.002a.284.284 0 0 0 .024-.107v-.024h.024v.024a.21.21 0 0 0 .08-.094v-.001a.804.804 0 0 0 .059-.191l.001-.005q.119-.24.24-.525a5.05 5.05 0 0 0 .204-.556l.011-.041c.114-.273.236-.626.339-.987l.018-.074q.048-.226.095-.418h-.048v.048a.068.068 0 0 1-.024.048v-.004a.02.02 0 0 0-.021-.021h-.052v-.024c0-.064.014-.126.038-.181l-.001.003a.911.911 0 0 1 .082-.159l-.002.004h.024l.048.167h.048a.58.58 0 0 1 .013-.111l-.001.004a.583.583 0 0 0 .012-.107l-.012-.012a.047.047 0 0 0-.032-.012h-.004l-.048.024v-.24h-.024a.235.235 0 0 1-.036-.047v-.001a.215.215 0 0 0-.036-.048l-.24.024a.114.114 0 0 0-.012-.049v.001a.04.04 0 0 0-.036-.024a.162.162 0 0 1 .059-.036h.001q.036-.012.131-.036l.048.024l.048-.048q0 .024.012.024h.012l.048-.024q0 .024.012.024h.045v-.048h-.264l-.048.024v-.024q-1.79.143-2.052.167h-.24q0 .024-.012.024t-.036.024l-.071-.024l-.191.024h-.214a.235.235 0 0 1-.047.036h-.001a.106.106 0 0 1-.048.012l.501-.048a.235.235 0 0 1-.047.036h-.001a.16.16 0 0 1-.063.012h-.037a.05.05 0 0 1-.032-.012a.047.047 0 0 0-.032-.012h-.004q-.071.024-.167.048l-.191.048l-.071-.024l-.835.119a.114.114 0 0 1-.049-.012h.001a.106.106 0 0 0-.048-.012v.012q0 .012-.024.012a.58.58 0 0 0-.111.013l.004-.001a.432.432 0 0 0-.11.037l.002-.001l-.406.167l.024.048a.114.114 0 0 0 .049-.012h-.001a.215.215 0 0 0 .048-.036l.143.024l-.024.119a2.054 2.054 0 0 0-.146.109l.003-.002a.19.19 0 0 0-.071.149v.006q-.143.071-.226.119t-.08.071v.024h.024l.071-.024l.048.071v.095q0 .048-.978 2.099q-.43 1.002-.656 1.538t-.25.632a1.098 1.098 0 0 0-.245.11l.005-.003a.266.266 0 0 0-.119.129l-.001.002q0 .024.048.048a.266.266 0 0 0 .111.024h.009v.048l-.048.095v.095q0 .095.43.119l.012-.012a.047.047 0 0 1 .032-.012h.004a1.396 1.396 0 0 0-.035.195l-.001.007a.831.831 0 0 1-.038.185l.002-.006h.071a.507.507 0 0 0 .106-.152l.001-.003c.023-.049.044-.107.061-.168l.002-.008h.024c.019.001.036.01.048.024a.26.26 0 0 1 .047.07l.001.002a4.982 4.982 0 0 0-.371.823l-.011.036l.024.048h.024c.149-.424.318-.786.518-1.129l-.017.031q.525-.048 1.79-.191l.048-.024l.024.024h.071l.048-.048q0 .024.012.024h.012l.644-.048l.021.001a.285.285 0 0 0 .16-.049l-.001.001a.13.13 0 0 0 .06-.095v-.001q0-.024.071-.048c.068-.02.149-.037.233-.047l.007-.001c.018 0 .034.005.049.012h-.001c.014.007.03.012.048.012l.071-.024a.686.686 0 0 1 .215-.024h-.001v-.009c0-.04.009-.078.025-.112l-.001.002a.08.08 0 0 1 .071-.048v-.048a.351.351 0 0 0-.098.013l.003-.001a.474.474 0 0 0-.098.037l.003-.001a.418.418 0 0 0-.131-.024a.425.425 0 0 1-.134-.025l.003.001v.004a.02.02 0 0 1-.021.021h-.004a.068.068 0 0 0-.048.024l-.454-.024l-.048.024q0-.024-.048-.024h-.024l-.048.048q0-.024-.036-.036a.343.343 0 0 0-.093-.012h-.015h.001l-.048.024h-.119v-.024q.859-.095.859-.24h.004a.05.05 0 0 1 .032.012c.008.009.012.02.012.032v.004l.167-.024v-.048l-1.36.024q0 .024-.012.024h-.188a.17.17 0 0 0-.064.013h.001a.215.215 0 0 0-.048.036l-.071-.024a.8.8 0 0 0-.197.025l.005-.001l-.191.048a.351.351 0 0 1-.098-.013l.003.001a.358.358 0 0 0-.095-.012h-.095a.193.193 0 0 0-.096.025l.001-.001l-.31.024l-.143-.024c.031-.153.076-.289.136-.417l-.005.011q.131-.334.418-.93c.036-.106.076-.197.124-.283l-.004.009a.989.989 0 0 1 .167-.226h.048c.019.001.036.01.048.024l.191-.071h.119c.119 0 .234-.018.342-.05l-.009.002c.134-.039.244-.079.352-.126l-.018.007q0 .024.012.024h.107l.167-.024l.501-.048q.191-.024.37-.06t.346-.08l-.024-.048l-.025.001a.374.374 0 0 1-.089-.011l.003.001q-.036-.012-.036-.036a.08.08 0 0 0-.047-.071h-.001a.335.335 0 0 0-.125-.024l-.019.001h.001v-.048l.119-.048l.382-.024v-.048h-.24v-.048l.143-.024h.004a.05.05 0 0 0 .032-.012a.047.047 0 0 0 .012-.032V9.93h-.028c-.066 0-.13-.004-.194-.013l.007.001q-.071-.012-.071-.036a14.74 14.74 0 0 1-1.987.047l.03.001h-.004a.05.05 0 0 0-.032.012l-.036.036l-.071-.071c.155-.421.364-.784.625-1.104l-.005.007a.2.2 0 0 1-.012.061V8.87a.158.158 0 0 1-.036.06v.024h.048c.072-.149.139-.27.212-.387l-.01.017q.08-.13.155-.131q.31-.024.24-.06t.573-.131c.018 0 .034.005.049.012h-.001c.014.007.03.012.048.012c.4-.09.888-.166 1.386-.211l.045-.003v-.071l-.334.048v-.048a.114.114 0 0 0 .049-.012h-.001a.215.215 0 0 0 .048-.036v.024a.114.114 0 0 0 .049-.012h-.001l.022-.014h.143a.407.407 0 0 0 .122-.025l-.003.001a.438.438 0 0 1 .143-.024l.214-.024l.072-.072v-.024h-.057a.17.17 0 0 1-.064-.013h.001a.16.16 0 0 0-.063-.012h-.009v-.004a.05.05 0 0 1 .012-.032a.047.047 0 0 1 .032-.012h.006l.382-.048l.191-.048v.024h.028a.05.05 0 0 0 .032-.012a.047.047 0 0 1 .032-.012h.028l-.048.095a.2.2 0 0 0-.024.095l-.048.024a.114.114 0 0 0-.012-.049v.001a.215.215 0 0 0-.036-.048h-.024q-.358.954-.56 1.538t-.537 1.324a1.186 1.186 0 0 1-.111.305l.003-.007a.784.784 0 0 1-.081.133l.001-.002v.004a.05.05 0 0 0 .012.032l.012.012h.024v.057a.543.543 0 0 1-.025.162l.001-.004a.58.58 0 0 1-.073.145l.001-.002q-.143.43-.24.668a.694.694 0 0 1-.167.287l.048.262v.004a.05.05 0 0 1-.012.032l-.012.012h-.048l.001-.016a.138.138 0 0 0-.025-.08a.097.097 0 0 0-.064-.024h-.031a.871.871 0 0 0-.078.149l-.002.006c-.022.05-.043.109-.058.171l-.002.008h.048a.186.186 0 0 0 .071-.059v-.001a.145.145 0 0 0 .024-.08h.098l.048.024a.147.147 0 0 0-.036.07v.001a.474.474 0 0 1-.037.098l.001-.003l.048.048l-.071.143c.035 0 .068.009.096.025l-.001-.001a.277.277 0 0 1 .071.048a.54.54 0 0 1 .076-.034l.005-.002a.158.158 0 0 0 .06-.036h.597l.167-.048h.102c.061 0 .122-.004.181-.013l-.007.001a1.18 1.18 0 0 1 .174-.012h.053l.048-.024v.024a.2.2 0 0 0 .061-.012h-.001a.188.188 0 0 1 .06-.012h.005c.045 0 .088.004.131.013l-.004-.001a.583.583 0 0 0 .107.012v.024h-.119v.048a.114.114 0 0 0 .049-.012h-.001a.106.106 0 0 1 .048-.012l.012.012c.008.009.012.02.012.032v.004a.2.2 0 0 0-.061.012h.001a.252.252 0 0 1-.079.012h-.001l-.024-.024l-.071.024l-.048.024l-.071-.024v.012q0 .012-.024.012v.004a.05.05 0 0 0 .012.032c.009.008.02.012.032.012h.075l.382-.048l.143-.024l.012.012q.012.012.012-.012l.071-.024h.157a.342.342 0 0 0 .096-.013l-.003.001a1.51 1.51 0 0 1 .194-.035l.008-.001q0-.024.012-.024h.012v.024l.48-.096l-.024.024q.071-.024.06-.024t-.012-.024q0 .024.012.024h.08v-.048h-.167a.2.2 0 0 0-.061.012h.001a.188.188 0 0 1-.06.012h-.36v-.024l.191-.024h.071l.048-.024v.024c.121-.015.261-.024.402-.024h.016h-.001a7.512 7.512 0 0 0 1.162-.203l-.052.011l.008.001a.055.055 0 0 0 .052-.036q.012-.036.036-.131h-.071a1.06 1.06 0 0 0-.161.013l.006-.001a1.754 1.754 0 0 0-.192.039l.012-.003l-.012-.012a.047.047 0 0 1-.012-.032v-.028h.176a.193.193 0 0 0 .096-.025l-.001.001l.095-.048v.004a.02.02 0 0 1-.021.021h-.004zm-8.751.143h-.095c.016-.083.037-.155.063-.224l-.003.01c.028-.078.055-.142.085-.203l-.005.012l.048.024a.545.545 0 0 1-.037.195l.001-.004c-.019.05-.038.113-.054.178zm.095-.406v-.028a.05.05 0 0 1 .012-.032a.242.242 0 0 1 .058-.035l.002-.001v.048a.04.04 0 0 1-.024.036a.099.099 0 0 1-.049.012h-.002zm1.12.859v.024h-.143v-.024zm-.167-.024q0 .024-.036.048a.47.47 0 0 1-.104.047l-.003.001l-.167-.024h-.048v-.048zm1.12-5.177h-.095v-.028a.05.05 0 0 1 .012-.032a.047.047 0 0 1 .032-.012h.052l.048.024v.004a.05.05 0 0 1-.012.032a.047.047 0 0 1-.032.012zm1.264-.24h.024v.048h-.071v-.004a.05.05 0 0 1 .012-.032a.046.046 0 0 1 .03-.011h.005zm-.31.12l-.143-.024v-.024h.005c.061 0 .122-.004.181-.013L18 7.714a.37.37 0 0 0 .156-.061l-.001.001v.004a.05.05 0 0 0 .012.032c.009.008.02.012.032.012h.004v.048h-.071a.114.114 0 0 1-.049-.012l-.023-.012a.068.068 0 0 1-.048.024h-.051zm.143.453v-.024h.214v.024zm.215-.167v-.024h.095v.024zm.024-.358h-.024l.024-.048l.24-.024v.024a.986.986 0 0 1-.124.035l-.007.001a.51.51 0 0 1-.108.012zm.119.501v-.048h.101c.012 0 .022.005.03.012a.04.04 0 0 1 .012.03v.006zm.597-.835l-.119.048h-.261h.005c.078 0 .155-.004.231-.013l-.009.001q.107-.012.155-.012zm-.12.574v.024l-.191.024v-.024zm.287.144l-.036-.012a.587.587 0 0 1-.083-.038l.003.002l-.167.048h-.147a.05.05 0 0 1-.032-.012l-.012-.012h.167l.334-.071h.191v.024q-.071.024-.119.036a.473.473 0 0 1-.109.012h-.01h.001zm1.813-.48h.024a.873.873 0 0 1 .093-.078l.002-.002a.423.423 0 0 1 .116-.059l.003-.001a.963.963 0 0 1-.11.351l.003-.005a1.78 1.78 0 0 1-.182.278l.003-.004l-.119-.048c.027-.055.055-.122.077-.191l.003-.011q.036-.107.06-.202zm-.311.524h.024a.277.277 0 0 1-.048.071a.068.068 0 0 1-.048.024v-.071h.036a.067.067 0 0 0 .037-.022zm-.191.48v-.003c0-.064.013-.124.037-.179l-.001.003a.504.504 0 0 1 .107-.155V8.3a.603.603 0 0 1-.062.206l.002-.004a.491.491 0 0 1-.06.08v.095a.73.73 0 0 1-.05.083l.002-.003a.114.114 0 0 0-.024.06h-.024l-.026-.136a.068.068 0 0 0 .048-.024a.068.068 0 0 0 .024-.048zm-.071.24q-.024.071-.036.119a.474.474 0 0 1-.037.098l.001-.003l-.024-.024v-.048a.263.263 0 0 1 .048-.096a.068.068 0 0 1 .048-.024zm-1.455 4.176v-.024h.048v.071l-.024-.012a.038.038 0 0 1-.023-.035v-.002zm2.052.048v-.011q0-.012-.024-.012l-.119.024v-.024h.024l.262-.024v.024h-.001a.245.245 0 0 0-.081.013l.002-.001a.176.176 0 0 1-.058.009h-.005zm-.12-.382l-.143.024v-.024l.143-.024l.119-.024v.024h.024a.718.718 0 0 1 .116-.024h.003a.186.186 0 0 0 .072-.024h-.001l.012.012c.008.009.012.02.012.032v.004zm.549-.024a.182.182 0 0 0-.07-.024h-.096v-.024h.143l.024.024zm.167-.048v-.024h.004a.05.05 0 0 0 .032-.012l.012-.012h.024l.024.012a.04.04 0 0 1 .024.036zm.215-.024v-.024h.095v.024zm.119-.024v-.024l.024-.024l.048.048v-.024l.024.048zm1.79-.215v-.014q0-.012-.024-.012t-.012.012t-.012.012v.024l-.024.048l.071-.024zm.119.095h-.023v.048h.119q0-.024-.036-.024h-.06z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Backward(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m.511 13.019l16.695 10.836c.656.427 1.461-.136 1.461-1.02V16.15l11.872 7.705c.656.427 1.461-.136 1.461-1.02V1.166c0-.884-.805-1.447-1.461-1.02L18.667 7.85V1.166c0-.884-.805-1.447-1.461-1.02L.511 10.981a1.27 1.27 0 0 0-.003 2.037l.003.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Baidu(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.334 12.647c2.609-.562 2.249-3.683 2.175-4.365c-.128-1.05-1.366-2.883-3.043-2.739C.357 5.729.049 8.782.049 8.782c-.287 1.41.682 4.425 3.285 3.864zm4.843-5.239c1.44 0 2.599-1.659 2.599-3.709S9.619.001 8.181.001S5.571 1.65 5.571 3.7s1.17 3.709 2.609 3.709zm6.208.245c1.929.26 3.161-1.799 3.411-3.359A3.512 3.512 0 0 0 15.461.633l-.025-.008c-1.37-.316-3.059 1.873-3.229 3.299c-.18 1.749.25 3.489 2.169 3.736zm7.631 2.629a3.014 3.014 0 0 0-2.904-2.992h-.005c-2.299 0-2.609 2.119-2.609 3.619c0 1.43.118 3.419 2.984 3.359s2.542-3.239 2.542-3.989zm-2.899 6.533a26.925 26.925 0 0 1-4.678-4.735l-.047-.063c-2.359-3.676-5.713-2.179-6.832-.32a15.656 15.656 0 0 1-3.075 3.353l-.029.023c-.25.31-3.599 2.119-2.853 5.418s3.359 3.239 3.359 3.239a14.665 14.665 0 0 0 4.257-.329l-.098.019a8.507 8.507 0 0 1 1.835-.195c.827 0 1.627.115 2.385.33l-.062-.015s5.205 1.749 6.646-1.609c1.424-3.369-.81-5.108-.81-5.108z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ban(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.46 11.992v-.059a8.24 8.24 0 0 0-1.376-4.571l.019.031l-11.757 11.74a8.259 8.259 0 0 0 4.606 1.388h.026h-.001h.029a8.19 8.19 0 0 0 3.321-.699l-.052.021c3.074-1.315 5.188-4.314 5.188-7.807v-.045v.002zM4.884 16.654L16.657 4.897a8.112 8.112 0 0 0-4.607-1.42h-.074h.004h-.064a8.256 8.256 0 0 0-4.231 1.159l.038-.021C5.179 6.125 3.5 8.859 3.5 11.984c0 1.736.518 3.352 1.408 4.7l-.02-.032zm19.066-4.662v.035c0 1.678-.35 3.273-.981 4.718l.03-.076c-1.842 4.36-6.082 7.363-11.024 7.363S2.793 21.028.981 16.747l-.029-.078C.351 15.29.001 13.684.001 11.995S.351 8.7.982 7.244l-.03.077C2.794 2.961 7.034-.042 11.976-.042s9.182 3.004 10.994 7.285l.029.078a11.56 11.56 0 0 1 .952 4.631v.041v-.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bandage(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.815 2.296c.692-.689 1.646-1.116 2.7-1.116s2.009.426 2.7 1.116l.492.492c.689.692 1.116 1.646 1.116 2.7s-.426 2.009-1.116 2.7l-2.137 2.137l.838.839l2.138-2.138c.904-.905 1.464-2.155 1.464-3.536s-.559-2.631-1.464-3.536l-.492-.492c-.905-.905-2.155-1.464-3.536-1.464s-2.631.56-3.536 1.464L12.844 3.6l.838.839zm-5.482 17.268l-2.138 2.138c-.692.69-1.646 1.116-2.7 1.116s-2.009-.426-2.7-1.116l-.492-.491c-.689-.692-1.115-1.646-1.115-2.7s.426-2.008 1.116-2.701l2.137-2.137l-.838-.839l-2.138 2.138C.56 15.877.001 17.127.001 18.508s.56 2.631 1.464 3.536l.492.491C2.862 23.44 4.112 24 5.493 24s2.631-.56 3.536-1.465l2.138-2.138zm12.213-4.592L9.039 1.465C8.134.56 6.884 0 5.503 0S2.872.56 1.967 1.465l-.492.491C.57 2.861.011 4.111.011 5.492s.56 2.631 1.464 3.536l13.507 13.506c.905.905 2.155 1.465 3.536 1.465s2.631-.56 3.536-1.465l.492-.492a5 5 0 0 0 1.459-3.536c0-1.38-.557-2.629-1.459-3.536zM5.278 11.161L2.306 8.189c-.689-.692-1.116-1.646-1.116-2.7s.426-2.009 1.116-2.7l.492-.491c.692-.689 1.646-1.116 2.7-1.116s2.009.426 2.7 1.116L11.17 5.27zm16.428 10.043l-.492.492c-.692.689-1.646 1.115-2.7 1.115s-2.009-.426-2.701-1.116l-2.972-2.971l5.894-5.894l2.972 2.972c.691.692 1.119 1.648 1.119 2.704s-.428 2.011-1.119 2.704z"/><path fill="currentColor" d="M9.261 6.069a.558.558 0 0 0-.394-.958a.56.56 0 1 0 .394.958M5.286 2.094a.558.558 0 0 0 .394.958a.56.56 0 1 0-.394-.958m0 6.37a.558.558 0 0 0 .394.958a.56.56 0 1 0-.394-.958M2.098 5.276a.558.558 0 0 0 .394.958a.56.56 0 1 0-.394-.958m3.973.8a.558.558 0 0 0-.394-.958a.56.56 0 1 0 .394.958m15.851 11.86a.558.558 0 0 0-.958.394a.56.56 0 1 0 .958-.394m-3.983-3.188a.558.558 0 0 0 .394.958a.56.56 0 1 0-.394-.958m0 6.37a.558.558 0 0 0 .394.958a.56.56 0 1 0-.394-.958m-3.181-3.182a.558.558 0 0 0 .394.958a.56.56 0 1 0-.394-.958m3.968.795a.558.558 0 0 0-.394-.958a.56.56 0 1 0 .394.958"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BarChart(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 12v8H6v-8zm6-8v16h-4V4zm16 18v2H0V0h2v22zM22 8v12h-4V8zm6-6v18h-4V2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryEmpty(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.12 17.6A3.12 3.12 0 0 1 0 14.48V8.72A3.12 3.12 0 0 1 3.12 5.6h16.017a3.12 3.12 0 0 1 3.12 3.12v5.76a3.12 3.12 0 0 1-3.12 3.12zM1.108 8.72v5.762c0 1.111.9 2.011 2.011 2.011h16.018a2.01 2.01 0 0 0 2.009-2.011V8.72a2.012 2.012 0 0 0-2.009-2.01H3.119c-1.11 0-2.01.9-2.011 2.01m21.548.991a2 2 0 0 1 .014 3.775l-.014.004z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryFull(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.12 17.6A3.12 3.12 0 0 1 0 14.48V8.72A3.12 3.12 0 0 1 3.12 5.6h16.011a3.12 3.12 0 0 1 3.12 3.12v5.76a3.12 3.12 0 0 1-3.12 3.12zM1.108 8.72v5.76a2.012 2.012 0 0 0 2.011 2.01h16.012a2.011 2.011 0 0 0 2.009-2.01V8.72a2.012 2.012 0 0 0-2.009-2.01H3.119c-1.11 0-2.01.9-2.011 2.01m2.679 6.4a1.312 1.312 0 0 1-1.307-1.308V9.385a1.31 1.31 0 0 1 1.305-1.31h14.678a1.31 1.31 0 0 1 1.311 1.31v4.427a1.316 1.316 0 0 1-1.311 1.308zm18.87-5.41a2 2 0 0 1 .014 3.773l-.014.004z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryHalf(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.12 17.6A3.12 3.12 0 0 1 0 14.48V8.72A3.12 3.12 0 0 1 3.12 5.6h16.016a3.12 3.12 0 0 1 3.12 3.12v5.76a3.12 3.12 0 0 1-3.12 3.12zM1.108 8.72v5.76c0 1.11.9 2.01 2.01 2.01h16.018a2.01 2.01 0 0 0 2.006-2.01V8.72a2.012 2.012 0 0 0-2.009-2.01H3.118c-1.11 0-2.01.9-2.011 2.01zm2.678 6.4a1.31 1.31 0 0 1-1.31-1.311V9.386a1.31 1.31 0 0 1 1.31-1.311h6.03c.724 0 1.31.587 1.31 1.31v4.428a1.312 1.312 0 0 1-1.31 1.311zm18.871-5.41c.789.28 1.344 1.02 1.344 1.889s-.555 1.609-1.33 1.885l-.014.004z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryQuarter(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.12 17.6A3.12 3.12 0 0 1 0 14.48V8.72A3.12 3.12 0 0 1 3.12 5.6h16.016a3.12 3.12 0 0 1 3.12 3.12v5.76a3.12 3.12 0 0 1-3.12 3.12zM1.108 8.72v5.761c0 1.111.9 2.011 2.011 2.011h16.017a2.012 2.012 0 0 0 2.01-2.011V8.719c0-1.11-.9-2.01-2.01-2.011H3.119a2.01 2.01 0 0 0-2.011 2.011zm2.678 6.407a1.311 1.311 0 0 1-1.311-1.311v-4.43c0-.724.587-1.311 1.311-1.311h2.708c.724 0 1.31.587 1.31 1.311v4.429a1.316 1.316 0 0 1-1.311 1.311zm18.87-5.416c.789.28 1.344 1.02 1.344 1.889s-.555 1.609-1.33 1.885l-.014.004z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BatteryThreeQuarters(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.12 17.6A3.12 3.12 0 0 1 0 14.48V8.72A3.12 3.12 0 0 1 3.12 5.6h16.016a3.12 3.12 0 0 1 3.12 3.12v5.76a3.12 3.12 0 0 1-3.12 3.12zM1.108 8.72v5.76c0 1.11.9 2.01 2.01 2.01h16.018a2.01 2.01 0 0 0 2.006-2.01V8.72a2.012 2.012 0 0 0-2.009-2.01H3.118c-1.11 0-2.01.9-2.011 2.01zm2.678 6.4a1.31 1.31 0 0 1-1.31-1.311V9.386a1.31 1.31 0 0 1 1.31-1.311h10.986a1.31 1.31 0 0 1 1.308 1.31v4.428a1.316 1.316 0 0 1-1.308 1.307zm18.871-5.41c.789.28 1.344 1.02 1.344 1.889s-.555 1.609-1.33 1.885l-.014.004z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BeachSlipper(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.322 18.742q-.103-2.106-.042-4.211l.138-4.769c-.303-4.5-1.494-4.975-5.154-3.654l-.268.103C1.229 6.9.001 8.588.001 10.564c0 .337.036.667.104.984l-.006-.031l.426 2.044c.523 2.491 1.81 5.024 1.81 7.57a2.872 2.872 0 0 0 2.869 2.87h.365a2.88 2.88 0 0 0 2.87-3.021v.007zm-.674-3.509a.377.377 0 0 1-.365.289c-.027 0-.055-.006-.09-.006a.4.4 0 0 1-.247-.184l-.001-.002a.332.332 0 0 1-.026-.27l-.001.002a4.645 4.645 0 0 0 .375-1.844c0-.719-.161-1.401-.449-2.011l.012.029a4.442 4.442 0 0 0-2.329-1.958l-.031-.01c-.007 0-.007.007-.014.007a4.41 4.41 0 0 0-2.229 2.291l-.011.029a4.967 4.967 0 0 0-.235 1.526c0 .916.243 1.775.669 2.517l-.013-.025a.33.33 0 0 1 .006.27l.001-.002a.397.397 0 0 1-.217.213l-.003.001a.398.398 0 0 1-.127.021h-.004a.369.369 0 0 1-.348-.245l-.001-.003C.264 11.21 3.071 9.2 3.945 8.711a.457.457 0 0 1-.034-.164v-.001l-.001-.033a.51.51 0 1 1 1.004.126l.001-.004c1.038.432 3.806 2.077 2.733 6.598m2.258-.234a2.869 2.869 0 0 0 2.344 3.417l.017.002l.358.055a2.872 2.872 0 0 0 3.278-2.378l.002-.016c.392-2.512 2.058-4.817 2.96-7.2l.736-1.954c.191-.49.302-1.058.302-1.651a4.656 4.656 0 0 0-2.317-4.028l-.022-.012l-.254-.145c-3.414-1.872-4.666-1.583-5.657 2.814l-.598 4.728a53.855 53.855 0 0 1-.695 4.16zM15.301 3.33l-.001-.026c0-.044.005-.088.015-.129l-.001.004a.512.512 0 1 1 .99.223l.001-.004a.623.623 0 0 1-.064.163l.002-.003c.785.619 3.248 3.034.84 7.377a.365.365 0 0 1-.388.185h.002a.464.464 0 0 1-.127-.043l.003.001a.377.377 0 0 1-.185-.246v-.003a.33.33 0 0 1 .049-.269l-.001.001a5.038 5.038 0 0 0 1.097-3.15c0-.268-.021-.53-.061-.787l.004.029a4.37 4.37 0 0 0-1.835-2.626l-.016-.01c-.007 0-.007-.007-.014-.014a4.445 4.445 0 0 0-2.629 1.568l-.006.008a4.682 4.682 0 0 0-.796 2.622c0 .408.052.804.149 1.182l-.007-.033a.333.333 0 0 1-.07.262v-.001a.378.378 0 0 1-.274.145h-.091a.36.36 0 0 1-.31-.336v-.001c-.36-4.622 2.619-5.818 3.721-6.092z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BedPatient(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.102 11.147v1.731H1.904V6.672H0v12.414h1.904v-2.837h20.198v3.074H24v-8.178z"/><path fill="currentColor" d="M8.709 11.165c0 .564-.457 1.022-1.022 1.022H3.793a1.022 1.022 0 0 1-1.022-1.022v-.002c0-.564.457-1.022 1.022-1.022h3.894c.564 0 1.022.457 1.022 1.022zm11.034-4.001h-2.37V4.8h-1.68v2.365h-2.365v1.68h2.364v2.365h1.68V8.845h2.37z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Behance(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M27.618 3.676h-7.637V5.53h7.637zm-3.766 6.366a2.927 2.927 0 0 0-3.108 2.913h6.098a2.771 2.771 0 0 0-3.001-2.918l.01-.001zm.24 8.742a3.762 3.762 0 0 0 1.842-.49l-.019.01a2.23 2.23 0 0 0 1.131-1.285l.005-.016h3.302c-.702 2.668-3.092 4.604-5.935 4.604c-.157 0-.313-.006-.467-.018l.021.001a6.7 6.7 0 0 1-5.089-1.972a7.893 7.893 0 0 1-1.891-5.14c0-1.998.739-3.824 1.959-5.219l-.008.009a6.61 6.61 0 0 1 5.039-2.054h-.012a6.314 6.314 0 0 1 3.624 1.035l-.024-.015a6.309 6.309 0 0 1 2.271 2.636l.016.039a8.717 8.717 0 0 1 .756 3.578l-.001.134v-.007q0 .254-.03.703h-9.839a3.558 3.558 0 0 0 .86 2.567l-.004-.005a3.256 3.256 0 0 0 2.499.897l-.011.001zM4.14 18.032h4.423q3.064 0 3.064-2.494q0-2.69-2.974-2.69H4.14zm0-8.025h4.2a2.856 2.856 0 0 0 1.854-.551l-.008.006a2.041 2.041 0 0 0 .678-1.707l.001.008q0-2.152-2.839-2.153H4.14zM0 2.401h8.88l.126-.001c.775 0 1.532.076 2.264.222l-.074-.012a6.34 6.34 0 0 1 1.922.727l-.031-.016c.582.338 1.043.83 1.334 1.42l.009.02a5.039 5.039 0 0 1 .47 2.31v-.008l.002.131a4.093 4.093 0 0 1-2.545 3.79l-.027.01a4.496 4.496 0 0 1 2.562 1.706l.009.012a5.14 5.14 0 0 1 .866 3.056v-.009v.072a5.47 5.47 0 0 1-.379 2.005l.013-.037a4.599 4.599 0 0 1-.988 1.548l.001-.001c-.42.42-.905.774-1.439 1.046l-.032.015a6.905 6.905 0 0 1-1.763.62l-.046.008a9.788 9.788 0 0 1-1.955.192h-9.18z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bell(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.945 15.512c-.8-.786-1.619-1.6-1.619-5.44a7.835 7.835 0 0 0-6.539-7.717l-.046-.006a1.5 1.5 0 1 0-2.471.005l-.003-.005c-3.753.623-6.579 3.843-6.584 7.723v.001c0 3.84-.822 4.655-1.619 5.44A3.135 3.135 0 0 0 3.137 21h4.367a3 3 0 1 0 6 0h4.37a3.135 3.135 0 0 0 2.076-5.484l-.003-.003zm-9.441 6.613A1.127 1.127 0 0 1 9.379 21h2.251a1.127 1.127 0 0 1-1.126 1.125m7.36-3.376H3.138a.886.886 0 0 1-.625-1.509c1.34-1.34 2.418-2.612 2.418-7.17a5.572 5.572 0 0 1 11.144 0c0 4.578 1.089 5.84 2.418 7.17a.886.886 0 0 1-.625 1.509z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BellAlt(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.333 17.16c-1.04-1.04-2.339-2.341-2.339-7.409a7.505 7.505 0 0 0-6.22-7.393l-.045-.006a1.5 1.5 0 1 0-2.467.005l-.003-.005c-3.578.614-6.266 3.692-6.266 7.399c0 5.068-1.296 6.367-2.339 7.409A2.252 2.252 0 0 0 2.245 21h5.249a3 3 0 1 0 6 0h5.248a2.253 2.253 0 0 0 1.591-3.84m-9.84 4.965a.375.375 0 0 1 0 .75A1.877 1.877 0 0 1 8.618 21h.75a1.127 1.127 0 0 0 1.126 1.125z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bicycle(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M27.627 22.074a4.935 4.935 0 0 1-4.926-4.925a4.923 4.923 0 0 1 2.378-4.208l.022-.012l1.626 4.543c.136.376.489.64.904.64a.958.958 0 0 0 .327-.06l-.007.002a.962.962 0 0 0 .578-1.233l.002.007l-1.626-4.546c.215-.035.464-.057.717-.059h.003a4.935 4.935 0 0 1 4.926 4.925v.001a4.934 4.934 0 0 1-4.923 4.926h-.001zm-10.88-6.74l-3.124-6.263l8.186-.423zm-3.124.853a6.851 6.851 0 0 0-2.639-4.487l-.017-.012l.954-1.726l3.105 6.225zm-5.145 0l1.549-2.8a4.935 4.935 0 0 1 1.651 2.768l.005.032zM6.85 22.074a4.935 4.935 0 0 1-4.926-4.925a4.935 4.935 0 0 1 4.925-4.926c.536 0 1.051.086 1.533.245l-.035-.01l-2.337 4.231a.953.953 0 0 0 .016.956l-.002-.004c.17.285.476.472.826.473h4.828c-.47 2.274-2.456 3.959-4.835 3.96zm20.78-11.775h-.008c-.486 0-.96.051-1.418.147l.045-.008l-3.134-8.758a.965.965 0 0 0-.645-.599l-.007-.002l-3.84-1.048a.96.96 0 0 0-1.178.668l-.002.007a.96.96 0 0 0 .666 1.176l.007.002l3.36.916l1.39 3.873l-10.174.525l-.379-.758h1.11a.956.956 0 0 0 .959-.959V5.48a.96.96 0 0 0-.96-.96h-5.52a.96.96 0 0 0-.96.96v.001a.956.956 0 0 0 .959.959h2.258l.72 1.44l-1.593 2.88a6.695 6.695 0 0 0-2.437-.456h-.001a6.856 6.856 0 0 0-6.85 6.845v.001A6.857 6.857 0 0 0 6.846 24a6.854 6.854 0 0 0 6.768-5.851l.004-.037h2.956a.96.96 0 0 0 .763-.377l.002-.002l6.433-8.499l.666 1.862c-2.192 1.171-3.659 3.442-3.662 6.055v.001a6.856 6.856 0 0 0 6.848 6.849a6.856 6.856 0 0 0 6.848-6.848a6.856 6.856 0 0 0-6.848-6.847z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bing(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 0l4.795 1.686V18.56l6.753-3.9l-3.31-1.555l-2.09-5.2l10.64 3.738v5.435L4.795 24l-4.8-2.67V0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bitbucket(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.865 0H.854A.854.854 0 0 0 0 .843v.011c0 .048.004.095.012.141L.011.99l3.63 22.041c.096.55.567.964 1.136.97h17.424a.859.859 0 0 0 .847-.714l.001-.005l3.638-22.281a.856.856 0 0 0-.701-.98L25.981.02a.747.747 0 0 0-.129-.011h-.02h.001zm15.287 15.926H10.59L9.088 8.068h8.411z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bitcoin(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.024 7.385a3.42 3.42 0 0 1-1.87 3.712l-.02.009a4.338 4.338 0 0 1 2.518 1.479l.006.007c.445.664.711 1.481.711 2.36c0 .256-.023.508-.066.752l.004-.026a5.421 5.421 0 0 1-.483 1.837l.014-.034a3.988 3.988 0 0 1-.928 1.277l-.003.003a4.712 4.712 0 0 1-1.367.833l-.032.011a8.173 8.173 0 0 1-1.702.49l-.051.008c-.606.106-1.328.183-2.062.215l-.036.001v3.682H8.436v-3.62q-1.154 0-1.76-.014v3.634H4.455v-3.68q-.26 0-.779-.007t-.794-.007H-.001l.447-2.64h1.6a.771.771 0 0 0 .837-.734v-5.8h.23a1.466 1.466 0 0 0-.206-.015h-.025h.001V6.979a1.148 1.148 0 0 0-1.29-.978L1.599 6h-1.6V3.634l3.058.014q.923 0 1.399-.014V0h2.221v3.562q1.182-.029 1.76-.029V-.001h2.221v3.634c.751.061 1.44.176 2.108.344l-.089-.019a6.535 6.535 0 0 1 1.662.666l-.032-.017c.489.275.89.655 1.182 1.112l.008.013a3.64 3.64 0 0 1 .527 1.64l.001.012zm-3.101 7.861v-.038c0-.324-.08-.628-.221-.896l.005.01a2.165 2.165 0 0 0-.53-.66l-.003-.003a2.71 2.71 0 0 0-.81-.434l-.019-.006a6.536 6.536 0 0 0-.898-.259l-.046-.008a8.01 8.01 0 0 0-1.044-.128l-.023-.001q-.634-.043-.995-.043t-.93.014t-.685.014v4.872q.115 0 .534.007t.692.007t.765-.022t.844-.058t.822-.122c.31-.054.579-.124.839-.213l-.039.012c.268-.096.494-.198.71-.316l-.026.013c.221-.114.41-.259.568-.431l.001-.001c.144-.164.263-.355.349-.563l.005-.014a1.88 1.88 0 0 0 .136-.708v-.027v.001zM11.899 8.38v-.037c0-.293-.067-.57-.185-.818l.005.011a2.025 2.025 0 0 0-.439-.605l-.001-.001a2.141 2.141 0 0 0-.677-.395l-.015-.005a4.706 4.706 0 0 0-.761-.234l-.033-.006a7.431 7.431 0 0 0-.861-.114l-.027-.002q-.526-.043-.837-.036t-.779.014t-.57.007v4.428l.498.007q.426.007.67 0t.72-.029a6.12 6.12 0 0 0 .83-.085l-.036.005q.32-.058.742-.16c.264-.058.498-.15.711-.273l-.012.006a3.95 3.95 0 0 0 .538-.394l-.005.004a1.38 1.38 0 0 0 .386-.546l.003-.009c.084-.212.132-.458.132-.715v-.021v.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Blind(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.233.084c.515.219.912.632 1.105 1.145l.005.014a1.351 1.351 0 0 1-.003 1.049l.003-.009c-.148.465-.468.84-.882 1.057l-.01.005a1.81 1.81 0 0 1-1.396.106l.013.004A2.018 2.018 0 0 1 6.963 2.31l-.005-.014a1.351 1.351 0 0 1 .003-1.049l-.003.009A1.824 1.824 0 0 1 9.246.088zm-1.131 3.85a2.24 2.24 0 0 1 1.194 1.117l.006.014c.062.174.312 1.39.548 2.706l.437 2.393l1.145.999c.439.373.837.747 1.216 1.14l.005.005c.056.133.089.288.09.451v.001c.002.016.002.035.002.053c0 .196-.093.37-.237.48l-.001.001a.61.61 0 0 1-.544.236h.002h-.019a.93.93 0 0 1-.457-.12l.005.002c-.347-.236-2.72-2.338-2.823-2.497a2.087 2.087 0 0 1-.192-.513l-.003-.014c-.062-.27-.09-.326-.125-.222c-.166.527-.88 3.184-.88 3.266c.275.43.568.804.892 1.149l-.004-.004c.486.575.936 1.138 1.006 1.262c.111.202.118.52.118 3.835v3.614l-.16.236c-.185.284-.502.47-.862.47h-.019h.001a1.025 1.025 0 0 1-.86-.493l-.003-.005c-.139-.194-.144-.278-.166-3.426l-.021-3.225l-.916-1.082l-.916-1.075l-.11 1.581l-.617.902c-.34.493-1.173 1.725-1.852 2.731c-.48.742-.939 1.373-1.428 1.978l.034-.043a1.313 1.313 0 0 1-.481.132h-.004a.923.923 0 0 1-.969-.535l-.002-.006c-.305-.575-.263-.659 2.204-4.258l.96-1.394l.028-1.28a7.476 7.476 0 0 1 .367-2.582l-.015.053c.603-2.358 1.859-6.8 2-7.089a2.108 2.108 0 0 1 1.73-1.054h.005c.25.005.489.046.714.117l-.018-.005zm4.363 9.856c0 .152-.555 9.211-.597 9.67l-.042.541h-.27c-.236 0-.27-.021-.27-.16c-.007-.08.125-2.372.291-5.091c.285-4.8.298-4.938.437-4.966c.258-.056.451-.049.451.006"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Blogger(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.69 8.793a.84.84 0 0 0-.158.783L8.53 9.57c.05.279.261.498.531.559l.005.001a11.713 11.713 0 0 0 1.683.053l-.022.001h1.34c.141-.023.267-.061.386-.111l-.011.004a.868.868 0 0 0 .371-.712l-.001-.042v.002l.001-.035a.725.725 0 0 0-.425-.66l-.005-.002a.652.652 0 0 0-.36-.107h-1.351q-1.232 0-1.5.026a.782.782 0 0 0-.48.24zm3.323 5.032q-2.144 0-2.599.054a1.246 1.246 0 0 0-.621.19l.005-.003a.819.819 0 0 0-.238.342l-.002.006a.807.807 0 0 0-.053.404v-.004c.024.143.08.27.162.377l-.002-.002a.738.738 0 0 0 .343.266l.005.002a23.933 23.933 0 0 0 3.043.078l-.043.001l2.572-.054l.48-.054a.822.822 0 0 0 .429-.637v-.004a.731.731 0 0 0-.267-.696l-.001-.001a1.136 1.136 0 0 0-.616-.187h-.001q-.4-.026-2.599-.026zm11.949 7.021q.054-.483.054-9.163V4.721q-.054-1.447-.107-1.929a1.655 1.655 0 0 0-.219-.597l.004.007l-.054-.214A3.6 3.6 0 0 0 21.09.062l-.022-.003q-.32-.054-9.136-.054T2.85.059a3.9 3.9 0 0 0-1.743.962l.001-.001a3.594 3.594 0 0 0-.987 1.691l-.005.025A5.693 5.693 0 0 0 .01 4.25l-.001-.016v15.595c.011.542.05 1.062.115 1.573l-.008-.073c.108.437.283.82.517 1.161l-.008-.012c.247.351.55.644.898.873l.013.008c.345.235.748.419 1.179.53l.027.006c.456.057.995.095 1.541.104h15.446c.542-.011 1.062-.05 1.573-.115l-.073.008a3.594 3.594 0 0 0 1.873-1.09l.003-.003a3.298 3.298 0 0 0 .856-1.938l.001-.013zm-4.554-9.968q.054.32.026 2.492a16.62 16.62 0 0 1-.143 2.629l.009-.084a4.512 4.512 0 0 1-1.179 2.09a4.45 4.45 0 0 1-2.656 1.391l-.023.003q-.48.054-3.644.026t-3.537-.08a4.505 4.505 0 0 1-2.12-1.076l.004.003a4.449 4.449 0 0 1-1.357-2.007l-.009-.031a4.374 4.374 0 0 1-.187-1.23v-.003q-.026-.75-.026-3.322q.054-3.108.107-3.43a5.25 5.25 0 0 1 1.051-2.042l-.007.008a4.444 4.444 0 0 1 1.818-1.276l.031-.01a3.62 3.62 0 0 1 1.219-.319l.014-.001q.64-.054 2.357-.054q1.822.054 2.09.107c.859.29 1.594.749 2.198 1.341l-.001-.001a4.32 4.32 0 0 1 1.34 3.05v.004a1.69 1.69 0 0 0 .083.628l-.003-.012a.724.724 0 0 0 .291.318l.003.002a.734.734 0 0 0 .373.134h.002l.64.054c.287.007.562.046.826.113l-.025-.006a.84.84 0 0 1 .434.585z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Blood(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.127 15.608a5.217 5.217 0 0 0 5.211-5.211V5.911H1.916v4.486a5.218 5.218 0 0 0 5.21 5.211zM2.574 6.507h1.84v.406h-1.84zm0 1.475h1.84v.406h-1.84zm0 1.468h1.84v.406h-1.84zm0 1.475h1.84v.406h-1.84z"/><path fill="currentColor" d="M12.978 0H0v10.4a7.142 7.142 0 0 0 5.73 6.992l.046.008v.27c0 .21.227.375.51.375h.43v1.456h-.161a.277.277 0 0 0-.276.276v2.152c0 .152.123.274.274.274h.162V24h.83v-1.794h.162a.277.277 0 0 0 .276-.276v-2.155a.277.277 0 0 0-.276-.276h-.162v-1.456h.43c.286 0 .511-.166.511-.375v-.27c3.311-.661 5.772-3.542 5.776-6.999v-10.4zm0 10.4a5.851 5.851 0 0 1-11.7.004V1.28h11.7z"/><path fill="currentColor" d="M2.575 2.09h1.837v.406H2.575zm0 1.474h1.837v.406H2.575zm0 1.475h1.837v.406H2.575z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BloodDrop(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.647 13.663a1.923 1.923 0 0 0-.104-.239l.005.011c-.146-.341-1.32-2.609-1.516-2.918L7.581-.001l-5.59 10.776c-.19.31-1.232 2.32-1.376 2.66l-.057.126A7.456 7.456 0 0 0 0 16.417a7.582 7.582 0 1 0 15.164 0v-.006l.001-.101c0-.955-.19-1.866-.535-2.696l.017.047zm-7.062 8.355a5.585 5.585 0 0 1-4.899-8.269l-.015.029a6.382 6.382 0 0 0-.11 1.181v.001a6.7 6.7 0 0 0 6.696 6.696c.113 0 .234 0 .346-.006a5.53 5.53 0 0 1-2.003.368h-.023h.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BloodTest(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m17.355 20.131l2.616-.146l.153-2.615l-4.978-4.978l-2.767 2.762z"/><path fill="currentColor" d="M10.236 1.607L8.683 3.16L6.667 1.144a3.906 3.906 0 1 0-5.522 5.518l.002.002L3.164 8.68l-1.553 1.553L3.578 12.2l1.663-1.664l11.532 11.531l3.653-.207l2.136 2.136l1.45-1.45l-2.136-2.136l.207-3.653L10.552 5.226l1.664-1.664zm10.3 18.948l-3.266.186L6.158 9.629l3.453-3.453l11.112 11.112z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BluetoothB(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7.446 19.058l2.32-2.303l-2.32-2.303zm0-9.508l2.32-2.303l-2.32-2.303zM7.875 12l4.768 4.768L5.424 24v-9.52l-3.978 3.964L0 16.995L4.982 12L0 7.005l1.446-1.446l3.978 3.962V.001l7.216 7.232z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bold(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.22 7.091A7.1 7.1 0 0 0 14.127 0H2.181a2.182 2.182 0 0 0 0 4.364h.927v15.272h-.927a2.182 2.182 0 0 0 0 4.364H14.13a7.087 7.087 0 0 0 5.106-12.002l.002.002a7.04 7.04 0 0 0 1.981-4.906zm-7.09 2.727H7.473V4.363h6.655a2.727 2.727 0 0 1 0 5.454zm0 9.818H7.473v-5.454h6.655a2.727 2.727 0 0 1 0 5.454z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bookmark(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.08 1.6h8.88a.516.516 0 0 1 .48.478V20.24L7.6 16.56l-1.12-1.04l-1.12 1.04l-3.76 3.68V2.08a.516.516 0 0 1 .478-.48zm0-1.6A2.103 2.103 0 0 0 0 1.995V24l6.56-6.24L13.12 24V2.08A2.119 2.119 0 0 0 11.042 0h-.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BookmarkAlt(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 2.089V24l6.545-6.26L13.089 24V2.089A2.11 2.11 0 0 0 10.98 0l-.077.001h.004h-8.724L2.11 0A2.109 2.109 0 0 0 .001 2.088z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bower(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M26.779 11.722c-1.404-1.35-8.425-2.192-10.64-2.437a5.89 5.89 0 0 0 .264-.74l.011-.046c.255-.117.575-.239.904-.342l.062-.017c.041.122.235.586.341.806c4.456.123 4.684-3.31 4.865-4.252a5.313 5.313 0 0 1 1.695-3.43l.004-.003c-2.275-.664-5.551 1.024-6.648 3.547a6.475 6.475 0 0 0-1.193-.335l-.042-.006c-.695-2.591-3.022-4.468-5.788-4.468h-.031h.002a10.524 10.524 0 0 0-7.177 2.939l.004-.003A10.733 10.733 0 0 0 .973 6.414l-.027.069a11.769 11.769 0 0 0-.947 4.667v.107v-.005c0 5.946 4.059 11.157 6.352 11.157a2.26 2.26 0 0 0 2.06-1.408l.005-.015c.169.462.689 1.894.86 2.258c.253.54 1.421 1.007 1.934.446c.294.2.658.319 1.049.319c.594 0 1.123-.275 1.468-.705l.003-.004a1.865 1.865 0 0 0 2.407-1.375l.002-.012c.621-.033.926-.91.791-1.599a16.951 16.951 0 0 0-1.623-3.04l.039.061c.825.671 2.914.86 3.167 0c1.33 1.043 3.402.497 3.566-.353c1.616.42 3.468-.503 3.165-1.619c2.596-.18 2.264-2.94 1.536-3.64zM20.02 7.145a6.413 6.413 0 0 0-2.152-.433h-.01c-.861 0-1.387.488-2.198.488l-.102.002c-.286 0-.562-.042-.822-.121l.02.005a1.221 1.221 0 0 0 1.003.346l-.006.001a6.535 6.535 0 0 0 1.467-.327l-.046.014c.007.105.018.207.033.311a9.234 9.234 0 0 0-2.224.989l.04-.023a3.155 3.155 0 0 1-.28-1.311c0-.754.26-1.447.695-1.994l-.005.007a7.27 7.27 0 0 1 4.588 2.049l-.002-.002zm.966-.104l-.341-.32c-.329-.31-.686-.602-1.061-.869l-.031-.021a10 10 0 0 1 2.146-3.069l.002-.002a5.606 5.606 0 0 0-2.498 2.793l-.013.037a8.134 8.134 0 0 0-.833-.47l-.051-.023c.877-1.793 2.613-3.046 4.658-3.224l.022-.002c-1.365 1.238-.852 3.706-2.002 5.169zm-2.804 1.157a3.026 3.026 0 0 1-.284-1.18v-.006a4.7 4.7 0 0 1 .848.113l-.032-.006c-.024.152-.04.332-.045.514v.006c.048-.083.182-.369.237-.482a7.444 7.444 0 0 1 1.551.441l-.05-.018a4.22 4.22 0 0 1-2.211.618h-.015zm-8.045-2.22c-.552-.198-.552-.695 0-.893s1.244.051 1.244.446s-.693.644-1.244.446zm1.839.184v-.001a1.424 1.424 0 1 0-2.143 1.23l.007.004a1.428 1.428 0 0 0 2.138-1.232zm3.165-1.792a4.101 4.101 0 0 0-1.126 2.828c0 .854.26 1.648.706 2.305l-.009-.014a7.436 7.436 0 0 1-4.796 3.053l-.044.006a7.448 7.448 0 0 0 4.591-1.221l-.028.017c.367-.269.678-.588.93-.951l.009-.013c3.52.255 6.789.867 9.918 1.804l-.324-.083a1.2 1.2 0 0 1 .473.776l.001.007c-2.633-.369-7.38-.756-8.622-.822c.883.125 7.317 1.343 8.433 1.629a1.98 1.98 0 0 1-2.288.667l.014.005c.63.858-.593 1.887-2.297 1.321c.375.843-1.137 1.6-2.865.723c.022.843-2.14.94-2.993.008c.044.16.099.299.167.429l-.006-.012a4.41 4.41 0 0 1-4.356 3.984h-.002c-5.23-.028-9.458-4.274-9.458-9.507v-.064v.003c0-5.774 4.268-10.083 9.417-10.083h.026c2.08 0 3.851 1.32 4.521 3.169l.011.033z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Braille(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.324.254A1.462 1.462 0 0 1 2.789 1.9l.003-.01a1.466 1.466 0 0 1-1.378.972h-.02h.001A1.564 1.564 0 0 1 .044 1.774l-.003-.011A1.469 1.469 0 0 1 .715.157L.722.153a1.717 1.717 0 0 1 1.609.105zm6.911-.17c.503.24.85.735.874 1.313V1.4a1.746 1.746 0 0 1-.669 1.227l-.004.003a1.486 1.486 0 0 1-1.744-.141l.002.002a1.482 1.482 0 0 1-.453-1.213v.006c.043-.464.293-.861.656-1.103l.005-.003c.228-.113.497-.179.781-.179c.196 0 .385.031.562.09L9.232.085zm7.28.023c.439.237.74.68.773 1.196v.004a1.461 1.461 0 0 1-.609 1.284l-.004.003c-.227.161-.51.257-.815.257c-.197 0-.384-.04-.555-.113l.009.004a1.424 1.424 0 0 1-.852-1.06l-.001-.009a1.439 1.439 0 0 1 .705-1.521l.007-.004a1.812 1.812 0 0 1 1.358-.035l-.013-.004zm3.692.07a1.425 1.425 0 0 1 .704 1.321v-.004a1.425 1.425 0 0 1-.851 1.218l-.009.004a1.419 1.419 0 0 1-1.503-.205l.002.002a1.426 1.426 0 0 1-.471-1.45l-.002.01a1.91 1.91 0 0 1 .252-.58L18.325.5a1.576 1.576 0 0 1 1.893-.317l-.009-.004zM2.047 3.714a1.42 1.42 0 0 1 .798 1.609l.002-.01a1.42 1.42 0 0 1-1.384 1.112h-.016h.001A1.328 1.328 0 0 1 .041 5.342L.04 5.334a1.373 1.373 0 0 1 .326-1.281l-.001.001a1.37 1.37 0 0 1 1.687-.337l-.008-.004zm7.159-.046a1.43 1.43 0 0 1 .872 1.527l.001-.008A1.431 1.431 0 0 1 8.834 6.41l-.007.001a2.509 2.509 0 0 1-.625-.017l.014.002a1.968 1.968 0 0 1-.862-.741l-.005-.008a1.93 1.93 0 0 1-.019-1.221l-.004.014a1.49 1.49 0 0 1 1.896-.767l-.01-.003zM2.001 8.854a1.507 1.507 0 0 1 .835 1.674l.002-.01a1.358 1.358 0 0 1-1.405 1.091h.003l-.061.001a1.354 1.354 0 0 1-1.33-1.099l-.001-.008a1.45 1.45 0 0 1 .482-1.409l.002-.002a1.445 1.445 0 0 1 1.483-.237l-.01-.003zm3.661.039c.4.198.691.565.781 1.006l.002.009a1.433 1.433 0 0 1-.335 1.239l.001-.001a1.436 1.436 0 0 1-1.582.373l.01.003a1.443 1.443 0 0 1-.934-1.319v-.022c0-.506.262-.95.658-1.205l.006-.003a1.422 1.422 0 0 1 1.403-.076l-.008-.004zm3.537-.047c.429.219.745.606.864 1.072l.002.011a1.325 1.325 0 0 1-.273 1.123l.002-.002a1.313 1.313 0 0 1-1.165.56h.004a1.197 1.197 0 0 1-1.014-.439l-.002-.002a1.175 1.175 0 0 1-.379-1.012l-.001.006a1.24 1.24 0 0 1 .556-1.127l.005-.003A1.441 1.441 0 0 1 9.21 8.85l-.01-.003zm7.469.155a1.469 1.469 0 0 1 .484 1.826l.004-.009a1.307 1.307 0 0 1-1.297.789h.004h-.027a1.418 1.418 0 0 1-1.306-.865l-.003-.009a1.623 1.623 0 0 1 .357-1.619l-.001.001c.261-.219.601-.351.972-.351c.302 0 .583.088.82.24L16.669 9zm3.591-.015c.394.261.651.703.651 1.204l-.001.059v-.003a1.44 1.44 0 0 1-.735 1.204l-.007.004a1.446 1.446 0 0 1-1.567-.14l.003.002a1.453 1.453 0 0 1-.524-1.489l-.002.01c.136-.423.423-.765.796-.971l.009-.004a1.653 1.653 0 0 1 1.385.128zM2.318 12.639a1.438 1.438 0 0 1 .545 1.323l.001-.007a1.441 1.441 0 0 1-.84 1.14l-.009.004a1.443 1.443 0 0 1-1.216-.023l.008.004a1.444 1.444 0 0 1-.764-.924l-.002-.01a1.47 1.47 0 0 1 .689-1.621l.007-.004a1.716 1.716 0 0 1 1.588.122zm3.507-.046c.387.236.641.656.641 1.135l-.001.043v-.002l.002.066c0 .507-.268.952-.671 1.2l-.006.003a1.403 1.403 0 0 1-1.446.011l.007.004a1.46 1.46 0 0 1-.608-1.875l-.004.009a1.283 1.283 0 0 1 1.296-.805H5.03a1.157 1.157 0 0 1 .794.213l-.003-.002zm7.174-.023a1.355 1.355 0 0 1 .666 1.241v-.003a1.233 1.233 0 0 1-.463 1.082l-.003.002a1.434 1.434 0 0 1-1.837.021l.003.002a1.41 1.41 0 0 1-.477-1.605l-.003.01a1.41 1.41 0 0 1 1.379-.936h-.003l.055-.001c.251 0 .486.069.687.19l-.006-.003zm3.576-.039a1.465 1.465 0 0 1 .569 1.912l.004-.009a1.424 1.424 0 0 1-1.905.654l.008.004a1.442 1.442 0 0 1-.817-1.215v-.004a1.44 1.44 0 0 1 .665-1.302l.006-.003a1.666 1.666 0 0 1 1.481-.035l-.01-.004zM2.234 17.77a1.44 1.44 0 0 1 .64 1.327v-.005c-.048.534-.382.98-.845 1.188l-.009.004a1.365 1.365 0 0 1-1.593-.308l-.001-.001a1.394 1.394 0 0 1-.176-1.807l-.003.005c.204-.317.525-.544.9-.622l.009-.002a1.404 1.404 0 0 1 1.085.226l-.005-.003zm3.513-.069c.197.122.36.28.484.466l.004.006c.154.224.246.502.246.801c0 .346-.123.663-.328.91l.002-.002a1.426 1.426 0 0 1-1.672.4l.009.004a1.372 1.372 0 0 1-.874-1.493l-.001.008c.061-.487.355-.894.766-1.11l.008-.004a1.443 1.443 0 0 1 1.363.017l-.007-.004zm3.583-.015c.451.27.753.747.774 1.297v.003a1.996 1.996 0 0 1-.206.669l.005-.011c-.111.282-.334.5-.612.602l-.007.002a1.08 1.08 0 0 1-.796.155l.007.001a1.557 1.557 0 0 1-1.103-.736l-.004-.007a1.811 1.811 0 0 1-.019-1.372l-.004.012a1.52 1.52 0 0 1 1.975-.608l-.009-.004zm-7.322 3.552a1.552 1.552 0 0 1 .858 1.618l.001-.008a1.476 1.476 0 0 1-1.09 1.107l-.01.002a1.465 1.465 0 0 1-1.475-.503l-.002-.003a1.44 1.44 0 0 1-.288-.868c0-.232.054-.451.151-.645l-.004.008a1.46 1.46 0 0 1 1.269-.818h.002c.212.002.415.042.602.112l-.012-.004zm3.646.039a1.282 1.282 0 0 1 .813 1.284v-.004l.002.072c0 .36-.142.687-.374.927a1.478 1.478 0 0 1-1.053.438h-.017h.001a1.483 1.483 0 0 1-1.059-.461l-.001-.001a1.433 1.433 0 0 1-.353-.945c0-.326.108-.627.291-.868l-.003.004a1.438 1.438 0 0 1 1.761-.446l-.009-.004zm3.645.007a1.441 1.441 0 0 1 .654 1.89l.004-.009a1.336 1.336 0 0 1-1.133.813H8.82a2.456 2.456 0 0 1-.618-.017l.014.002a1.968 1.968 0 0 1-.862-.741l-.005-.008a1.907 1.907 0 0 1 .004-1.267l-.004.014c.167-.351.462-.619.823-.747l.01-.003a1.424 1.424 0 0 1 1.128.078l-.008-.004zm3.52-.046c.323.16.583.408.754.711l.005.009a2.069 2.069 0 0 1-.028 1.261l.004-.015c-.243.473-.727.79-1.285.79l-.048-.001h.002h-.003c-.434 0-.822-.196-1.081-.503l-.002-.002a1.407 1.407 0 0 1-.306-1.166l-.001.009c.084-.442.356-.807.729-1.015l.007-.004c.203-.114.446-.181.705-.181c.198 0 .387.039.559.111zm3.847.147a1.457 1.457 0 0 1 .476 1.835l.004-.009a1.435 1.435 0 0 1-1.609.743l.01.002a1.436 1.436 0 0 1-1.11-1.371v-.001a1.307 1.307 0 0 1 .804-1.313l.009-.003a1.508 1.508 0 0 1 1.423.121l-.006-.004zm3.607-.023a1.483 1.483 0 0 1 .391 2.009l.004-.006c-.258.37-.674.613-1.149.631h-.003l-.053.001c-.46 0-.87-.212-1.139-.545l-.002-.003a1.65 1.65 0 0 1-.286-1.114l-.001.007a1.6 1.6 0 0 1 .742-1.049l.007-.004a1.588 1.588 0 0 1 1.496.074l-.007-.004z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BrokenLink(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.184 20.239c-.61 0-1.104.494-1.104 1.104v1.586a1.104 1.104 0 1 0 2.208 0v-.033v.002v-1.555c0-.609-.494-1.104-1.103-1.104zm7.714-6.159h-1.555a1.105 1.105 0 0 0-.001 2.208h1.557a1.105 1.105 0 0 0 .001-2.208zm-2.579 4.678a1.104 1.104 0 0 0-1.559 1.561l-.001-.001l1.099 1.099a1.104 1.104 0 1 0 1.561-1.559l-.001-.001zM8.817 3.762c.61 0 1.104-.494 1.104-1.104V1.104a1.105 1.105 0 0 0-2.208-.001v1.556c0 .609.494 1.104 1.103 1.104zM2.658 9.92h.031a1.104 1.104 0 1 0 0-2.208h-.033h.002h-1.554a1.105 1.105 0 0 0-.001 2.208h.002zm1.024-4.677a1.104 1.104 0 1 0 1.561-1.559l-.001-.001l-1.099-1.099a1.104 1.104 0 1 0-1.561 1.559l.001.001zm10.66 12.221a1.104 1.104 0 0 0-1.561-1.559l.001-.001l-3.12 3.12a3.312 3.312 0 0 1-4.678-4.682l-.002.002l3.12-3.12a1.104 1.104 0 1 0-1.559-1.561l-.001.001l-3.12 3.12a5.52 5.52 0 0 0 3.9 9.422h.013a5.467 5.467 0 0 0 3.887-1.616zm6.243-14.052a5.467 5.467 0 0 0-3.887-1.616h-.013h.001h-.013a5.467 5.467 0 0 0-3.887 1.616l-3.12 3.12a1.104 1.104 0 1 0 1.559 1.561l.001-.001l3.12-3.12a3.312 3.312 0 1 1 4.683 4.679l-.002.002l-3.12 3.12a1.104 1.104 0 1 0 1.559 1.561l.001-.001l3.12-3.12c.996-.999 1.611-2.378 1.611-3.9s-.616-2.901-1.612-3.9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bug(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.304 2.305C5.719 4.007 6.158 6.727 6.158 6.727a6.453 6.453 0 0 0 4.81 1.92h-.01a6.455 6.455 0 0 0 4.799-1.919l.001-.001s.431-2.683-1.12-4.387c.971-.541 1.542-1.322 1.306-1.853c-.268-.599-1.45-.652-2.646-.119c-.43.19-.799.43-1.124.72l.004-.004a6.536 6.536 0 0 0-1.22-.114h-.001a7.083 7.083 0 0 0-1.192.104l.042-.006A4.148 4.148 0 0 0 8.738.379L8.711.368c-1.196-.531-2.38-.48-2.646.119c-.236.519.305 1.279 1.238 1.818zm12.028 12.39a5.84 5.84 0 0 0-.687-.103l-.025-.002c0-.064.011-.124.011-.19a11.79 11.79 0 0 0-.391-2.962l.018.082a3.707 3.707 0 0 0 1.782-.363l-.022.01c1.2-.534 1.949-1.45 1.68-2.045s-1.45-.652-2.646-.119a3.937 3.937 0 0 0-1.317.924l-.002.002c-.249-.6-.517-1.11-.824-1.592l.025.041a9.353 9.353 0 0 1-4.972 2.126l-.043.004v9.649a.96.96 0 0 1-1.92 0v-9.645a9.398 9.398 0 0 1-5.025-2.143l.015.012c-.271.42-.531.906-.752 1.413l-.026.067a4.09 4.09 0 0 0-1.232-.847l-.026-.01c-1.2-.531-2.38-.48-2.646.119s.487 1.513 1.68 2.045a3.752 3.752 0 0 0 1.622.363h.043h-.002c-.231.861-.366 1.85-.37 2.871v.003c0 .066.011.127.013.195c-.202.022-.41.053-.623.098c-1.558.32-2.752 1.128-2.664 1.793s1.424.945 2.986.62c.203-.042.4-.094.593-.152a10.768 10.768 0 0 0 1.349 3.384l-.027-.047a4.215 4.215 0 0 0-1.08.799l-.001.001c-.96.956-1.344 2.117-.866 2.596s1.639.09 2.595-.864c.255-.255.48-.54.668-.85l.011-.02a6.718 6.718 0 0 0 4.719 2.043h.006a6.741 6.741 0 0 0 4.781-2.095l.003-.003c.204.35.439.652.708.92c.954.956 2.117 1.344 2.596.866s.09-1.638-.866-2.593a4.338 4.338 0 0 0-1.128-.826l-.025-.011A10.575 10.575 0 0 0 18.325 17l.013-.071c.215.067.441.131.677.182c1.561.32 2.897.047 2.987-.62s-1.111-1.471-2.67-1.793z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Bus(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.275 19.36a.774.774 0 0 1 .772-.772h3.101a.774.774 0 0 1 .772.772a.774.774 0 0 1-.772.772h-3.101a.774.774 0 0 1-.772-.772m-6.174 0a.774.774 0 0 1-.772.772H4.228a.774.774 0 0 1-.772-.772a.774.774 0 0 1 .772-.772h3.101a.774.774 0 0 1 .772.772M4.36 2.25h13.727c.499 0 .903.404.903.903l-.462 10.349v.001a.903.903 0 0 1-.903.903h-12.8a.903.903 0 0 1-.903-.903v-.001L3.46 3.153c0-.499.405-.903.904-.903zm17.685 1.944h-1.531V1.051A1.05 1.05 0 0 0 19.467.003H2.987A1.05 1.05 0 0 0 1.941 1.05v3.143H.467a.463.463 0 0 0-.462.462v3.826c0 .255.207.461.462.462h1.474v15.055h3.845v-2.013h10.865v2.013h3.846v-2.839c.015-.066.024-.142.025-.22v-12h1.531a.463.463 0 0 0 .462-.462V4.651a.463.463 0 0 0-.462-.462z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func BusTicket(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.085 22.788l-.003.003a1.71 1.71 0 0 1-2.418 0l-.004-.004l-6.459-6.459a1.707 1.707 0 0 1 .006-2.407l.006-.006L13.915 1.213l.006-.006a1.707 1.707 0 0 1 2.407-.007l6.459 6.459l.003.003a1.71 1.71 0 0 1 0 2.418l-.004.004zm9.974-11.492l.12.12c.099.099.261.099.36 0s.099-.261 0-.36l-.12-.12c-.099-.099-.261-.099-.36 0s-.099.261 0 .36m-.84-.841l.12.12c.099.099.261.099.36 0s.099-.261 0-.36l-.12-.12c-.099-.099-.261-.099-.36 0s-.099.261 0 .36m.172-3.625l1.645 1.645c.099.099.261.099.36 0s.099-.261 0-.36L19.751 6.47c-.099-.099-.261-.099-.36 0s-.099.261 0 .36m-4.432-2.868l5.295 5.295c.099.099.261.099.36 0s.099-.261 0-.36l-5.295-5.295c-.099-.099-.261-.099-.36 0s-.099.261 0 .36m-.792 9.245a1.25 1.25 0 0 0 1.613-.127l.007-.007a1.272 1.272 0 0 0 .136-1.614l-.005-.006l.324-.324a1.14 1.14 0 0 0 .001-1.609l-1.211-1.211a2.383 2.383 0 0 0-3.362 0l-6.675 6.675a.26.26 0 0 0 0 .359l3.338 3.338a.26.26 0 0 0 .359 0l1.646-1.646a1.248 1.248 0 0 0 1.614-.127l.006-.006a1.273 1.273 0 0 0 .137-1.615l-.005-.006zm-3.554 4.85a.24.24 0 0 0 0-.34l-.01-.01a.26.26 0 0 0-.362.002l-2.053 1.957c-.099.099-.099.261 0 .36s.261.099.36 0zm2.629-3.014a.26.26 0 0 0 .359 0l.961-.961c.099-.099.099-.261 0-.36s-.261-.099-.36 0l-.961.961a.26.26 0 0 0 .003.363zm5.139-5.426l.12.12c.099.099.261.099.36 0s.099-.261 0-.36l-.12-.12c-.099-.099-.261-.099-.36 0s-.099.261 0 .36m-.841-.841l.12.12c.099.099.261.099.36 0s.099-.261 0-.36l-.12-.12c-.099-.099-.261-.099-.36 0s-.099.261 0 .36m-.84-.84l.12.12c.099.099.261.099.36 0s.099-.261 0-.36l-.12-.12c-.099-.099-.261-.099-.36 0s-.099.261 0 .36m-.84-.84l.12.12c.099.099.261.099.36 0s.099-.261 0-.36l-.119-.119c-.099-.099-.261-.099-.36 0s-.099.261 0 .36zm1.08-2.641l.468.468c.099.099.261.099.36 0s.099-.261 0-.36l-.468-.468c-.099-.099-.261-.099-.36 0s-.099.261 0 .36M4.767 17.229l.003.003a.25.25 0 0 0 .354 0l.003-.003l.42-.42c.099-.099.099-.261 0-.36s-.261-.099-.36 0l-.42.42a.261.261 0 0 0 0 .36M15.02 6.255l.12.12c.099.099.261.099.36 0s.099-.261 0-.36l-.12-.12c-.099-.099-.261-.099-.36 0s-.099.261 0 .36M3.662 16.724l.003.003a.25.25 0 0 0 .354 0l.003-.003l.721-.721c.099-.099.099-.261 0-.36s-.261-.099-.36 0l-.721.721a.261.261 0 0 0 0 .36M15.74 3.254l.469.469c.099.099.261.099.36 0s.099-.261 0-.36l-.469-.469c-.099-.099-.261-.099-.36 0s-.099.261 0 .36m-1.561 2.161l.12.12c.099.099.261.099.36 0s.099-.261 0-.36l-.12-.12c-.099-.099-.261-.099-.36 0s-.099.261 0 .36m-.84-.84l.12.12c.099.099.261.099.36 0s.099-.261 0-.36l-.121-.121c-.099-.099-.261-.099-.36 0s-.099.261 0 .36zm-.841-.84l.12.12c.099.099.261.099.36 0s.099-.261 0-.36l-.121-.121c-.099-.099-.261-.099-.36 0s-.099.261 0 .36zm1.885 8.943l-.048-.048c-.024-.024-.036-.06-.06-.084c-.012-.012-.012-.036-.024-.048c-.012-.036-.024-.048-.036-.085c0-.024-.012-.036-.012-.061a.121.121 0 0 1-.025-.069a.244.244 0 0 1-.012-.084c0-.024-.012-.036-.012-.06a.132.132 0 0 1 0-.096v-.001c0-.024-.012-.036 0-.048a.359.359 0 0 1 .023-.094l.001-.002c.012-.012 0-.024.012-.036c.012-.036.036-.06.048-.096c0-.024 0-.024.012-.036a.692.692 0 0 1 .22-.219l-.004.003c.024 0 .024 0 .036-.012c.036-.012.06-.036.096-.048c.012-.012.024 0 .036-.012a.35.35 0 0 0 .094-.023l.002-.001h.048c.036.012.073 0 .096 0s.036.012.061.012a.244.244 0 0 0 .084.012c.024 0 .048.024.073.024s.036.012.061.012s.048.024.084.036c.012.012.036.012.048.024a.152.152 0 0 0 .131.012h.001l.048.048a.764.764 0 0 1-1.076 1.079l-.002-.002zm.24-.24a.378.378 0 0 0 .529 0l.013-.012a.357.357 0 0 0-.014-.517l.001.001a.373.373 0 1 0-.528.528zm-4.082 4.082l-.048-.048c-.024-.024-.036-.06-.06-.084c-.012-.012-.012-.036-.024-.048c0-.024-.024-.048-.036-.084c0-.024-.012-.036-.012-.061a.122.122 0 0 1-.024-.072a.209.209 0 0 1-.012-.085c-.012-.012-.012-.036-.012-.06c-.012-.036.001-.072 0-.096v-.048a.359.359 0 0 1 .023-.094l.001-.002c.012-.012 0-.024.012-.036c.012-.036.036-.06.048-.096c0-.024 0-.024.012-.036a.692.692 0 0 1 .22-.219l-.004.003c.024 0 .024 0 .036-.012c.036-.012.06-.036.096-.048c.012-.012.024 0 .036-.012a.359.359 0 0 0 .094-.023l.002-.001h.144c.024 0 .036.012.06.012s.06.011.085.012s.048.024.072.024s.036.012.061.012s.048.024.084.036c.012.012.036.012.048.024a.152.152 0 0 0 .131.012h.001l.048.048a.764.764 0 0 1-1.076 1.079l-.002-.002zm.24-.24a.378.378 0 0 0 .529 0l.005-.005a.365.365 0 0 0 0-.517l-.006-.006a.373.373 0 1 0-.528.528m4.754-5.163a.826.826 0 0 0-.185-.085l.005.002a.5.5 0 0 0-.182-.06h.002c-.012-.012-.024 0-.036-.012c-.06-.011-.108-.012-.168-.024h-.049a.387.387 0 0 0-.166 0h-.002c-.012-.012-.024 0-.048 0a.484.484 0 0 0-.141.023l-.003.001c-.012.012-.024 0-.036.012c-.048.024-.108.036-.157.06c-.012.012-.036.012-.048.024c-.048.024-.108.06-.157.084l-.024.024a1.216 1.216 0 0 0-.293.294l.005-.006l-.026.026c-.024.048-.06.108-.084.157c-.012.012-.012.036-.024.048a.595.595 0 0 0-.06.159l.001-.003c-.012.012 0 .024-.012.036c0 .048-.024.096-.024.144c.012.012 0 .024-.001.048v.168c-.012.012 0 .024.001.048c.011.06.012.108.024.168c.012.012 0 .024.012.036c.008.054.023.11.046.164l.002.004a.59.59 0 0 0 .082.177l.002.003l-2.099 2.099a.831.831 0 0 0-.185-.086l.005.002a.66.66 0 0 0-.17-.048h.002c-.012-.012-.024 0-.036-.012a.942.942 0 0 0-.169-.024h.001c-.012.012-.024 0-.049 0a.768.768 0 0 0-.166 0h-.002c-.012-.012-.024 0-.048 0a.484.484 0 0 0-.141.023l-.003.001c-.012.012-.024 0-.036.012c-.048.024-.108.036-.157.06c-.012.012-.036.012-.048.024c-.048.024-.108.06-.157.084l-.025.025a1.216 1.216 0 0 0-.293.294l.005-.006l-.022.022c-.024.048-.06.108-.084.157c-.012.012-.012.036-.024.048a.595.595 0 0 0-.06.159l.001-.003c-.012.012 0 .024-.012.036c-.004.048-.023.096-.022.144c.012.012 0 .024 0 .048v.168c-.012.012 0 .024.001.048c.011.06.012.108.024.168c.012.012 0 .024.012.036c.013.053.033.112.056.171l.004.01c.016.06.043.121.081.177l.002.003l-1.466 1.466l-1.273-1.273l7.01-7.009l1.143-.446l.48.48l.001.001a.626.626 0 0 1 0 .886l-.001.001zm-1.513-1.704c-.024 0-.024 0-.036.012l-1.213-1.213a1.844 1.844 0 0 1 1.885.444l.001.001l.348.348zm-1.344 1.296l-1.344-1.344l.696-.696c.088-.084.179-.158.275-.223l-.011.008l1.32 1.32zM6.867 16.52l-1.344-1.344l.721-.721L7.589 15.8zm3.505-3.53l-1.344-1.344l1.933-1.933l1.345 1.345zm-2.439 2.44l-1.344-1.344l2.077-2.077l1.345 1.345z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calculator(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.502 0H1.498C.67 0 0 .671 0 1.498v21.004C0 23.33.671 24 1.498 24H19.5c.828 0 1.498-.671 1.498-1.498V1.498c0-.827-.67-1.497-1.497-1.498zM6 21H3v-3h3zm0-4.5H3v-2.998h3zM6 12H3V9.001h3zm6 9H9v-3h3zm0-4.5H9v-2.998h3zm0-4.5H9V9.001h3zm6 9h-3V9h3zm0-13.92H3V3h15z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Calendar(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.5 0h-3v3h3zm3 0H21v4.5h-6V0H9v4.5H3V0H1.499C.671 0 0 .671 0 1.499v21.002C0 23.329.671 24 1.499 24h21.002c.828 0 1.499-.671 1.499-1.499V1.499C24 .671 23.329 0 22.501 0zM21 21H3V7.5h18zM7.5 0h-3v3h3zm6 10.5h-3v3h3zm4.5 0h-3v3h3zM9 15H6v3h3zm0-4.5H6v3h3zm4.5 4.5h-3v3h3zm4.5 0h-3v3h3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Camera(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.846 9.692A4.161 4.161 0 0 1 18 13.846v.004a4.154 4.154 0 1 1-7.09-2.939a3.986 3.986 0 0 1 2.876-1.22h.063h-.003zm10.154-6h.055c1.002 0 1.908.414 2.554 1.081l.001.001a3.546 3.546 0 0 1 1.082 2.555v.058v-.003v12.924A3.693 3.693 0 0 1 24 24H3.692A3.693 3.693 0 0 1 0 20.308V7.33c0-1.002.414-1.908 1.081-2.554l.001-.001a3.546 3.546 0 0 1 2.555-1.082h.058h-.003h3.23l.735-1.962c.212-.507.557-.922.993-1.213l.01-.006A2.55 2.55 0 0 1 10.15 0h7.387a2.557 2.557 0 0 1 1.499.517L19.03.512c.445.297.791.712.996 1.201l.007.018l.735 1.962zM13.846 20.308l.091.001a6.204 6.204 0 0 0 4.472-1.896l.002-.002a6.205 6.205 0 0 0 1.897-4.474l-.001-.096v.005l.001-.092a6.204 6.204 0 0 0-1.896-4.472l-.002-.002c-1.167-1.172-2.781-1.897-4.565-1.897s-3.398.725-4.565 1.896a6.205 6.205 0 0 0-1.896 4.57v-.005l-.001.094c0 1.755.726 3.34 1.894 4.471l.002.002a6.205 6.205 0 0 0 4.474 1.897l.097-.001h-.005z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Car(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.515 13.754a.886.886 0 0 1-.88.88h-3.548a.885.885 0 0 1-.88-.88a.886.886 0 0 1 .88-.88h3.548a.888.888 0 0 1 .88.88m-12.376 0a.885.885 0 0 1-.88.88H3.711a.885.885 0 0 1-.88-.88a.886.886 0 0 1 .88-.88h3.548a.886.886 0 0 1 .879.88zm-1.84-8.167h11.404l1.399 3.562c-.022 0-.044-.004-.069-.004H4.899zm17.68 2.706a1.33 1.33 0 0 0-1.527-1.094l.008-.001l-2.183.356a.582.582 0 0 0-.094.026l.005-.002L18.782 4H5.216L3.81 7.578a.609.609 0 0 0-.087-.023l-2.185-.357A1.332 1.332 0 0 0 .019 8.286l-.001.008a1.331 1.331 0 0 0 1.088 1.519l.008.001l1.271.209a4.231 4.231 0 0 0-1.3 2.955v6.091h4.4v-2.3h12.429v2.3h4.4v-3.248c.018-.076.028-.163.028-.253v-2.586a4.244 4.244 0 0 0-1.213-2.876l.001.001l1.766-.29a1.329 1.329 0 0 0 1.092-1.527l.001.008z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretDown(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 6.4l12 12l12-12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretLeft(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m18.4 0l-12 12l12 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretRight(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m6.4 24l12-12l-12-12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CaretUp(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m24 18.4l-12-12l-12 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CenterAlign(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.334 2.666h26.665l.037.001a1.334 1.334 0 1 0 0-2.668L27.997 0h.002H1.334a1.334 1.334 0 0 0-.002 2.666zm3.555 2.667l-.037-.001a1.334 1.334 0 1 0 0 2.668l.039-.001h-.002h19.555a1.334 1.334 0 0 0 .002-2.666h-.002zm23.11 5.334H1.334a1.334 1.334 0 0 0-.002 2.666h26.667a1.334 1.334 0 0 0 .002-2.666zm0 10.666H1.334a1.334 1.334 0 0 0-.002 2.666h26.667a1.334 1.334 0 0 0 .002-2.666zm-3.555-2.667A1.334 1.334 0 0 0 24.446 16H4.889a1.334 1.334 0 0 0-.002 2.666h.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Check(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 10.909l4.364-4.364l8.727 8.727L28.364-.001l4.364 4.364l-19.636 19.636z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckboxActive(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24 24H0V0h18.4v2.4h-16v19.2h20v-8.8h2.4V24zM4.48 11.58l1.807-1.807l5.422 5.422l13.68-13.68L27.2 3.318L11.709 18.809z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CheckboxPassive(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24 24H0V0h24.8v24zm-1.6-2.4V2.4h-20v19.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Chrome(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m11.958 0l.131-.001c2.179 0 4.219.597 5.965 1.637L18 1.606a11.843 11.843 0 0 1 4.679 4.911l.031.067l-9.935-.52A6.543 6.543 0 0 0 8.811 7.08l.026-.016a5.806 5.806 0 0 0-2.468 3.032l-.012.04l-3.696-5.68A11.768 11.768 0 0 1 6.758 1.2l.07-.03a11.826 11.826 0 0 1 5.126-1.172h.002zM1.955 5.423l4.513 8.88a6.583 6.583 0 0 0 2.791 2.888l.035.017a5.748 5.748 0 0 0 2.792.713c.399 0 .789-.04 1.165-.117l-.037.006l-3.08 6.039a11.996 11.996 0 0 1-8.77-6.277l-.031-.066a11.627 11.627 0 0 1-1.332-5.441v-.138c0-2.423.73-4.674 1.982-6.548l-.027.043zm21.237 2.263a11.52 11.52 0 0 1 .8 4.155v.005l.001.172c0 1.405-.238 2.755-.677 4.011l.026-.085a11.873 11.873 0 0 1-2.065 3.662l.016-.02a12.031 12.031 0 0 1-3.237 2.772l-.056.031a11.684 11.684 0 0 1-5.962 1.615c-.248 0-.495-.008-.739-.023l.034.002l5.423-8.342a6.352 6.352 0 0 0 1.104-3.903l.001.014a5.831 5.831 0 0 0-1.433-3.723l.006.007zM12 7.954a4.046 4.046 0 0 1 4.044 4.044A4.046 4.046 0 0 1 12 16.042a4.046 4.046 0 0 1-4.044-4.044A4.046 4.046 0 0 1 12 7.954"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CircleOnotch(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24.126 11.937v.03c0 1.685-.353 3.287-.99 4.737l.03-.076c-1.228 2.913-3.499 5.184-6.334 6.383l-.078.029c-1.384.605-2.996.956-4.691.956s-3.307-.352-4.769-.986l.077.03c-2.913-1.228-5.184-3.499-6.382-6.334l-.029-.078a11.855 11.855 0 0 1-.956-4.703c0-2.029.501-3.941 1.387-5.618l-.032.066C3.134 2.998 6.373.599 10.208.009L10.275 0v3.183c-4.106.85-7.148 4.436-7.148 8.733c0 1.261.262 2.461.734 3.548l-.022-.058a8.867 8.867 0 0 0 4.696 4.732l.058.022c1.023.45 2.216.712 3.47.712s2.447-.262 3.526-.734l-.057.022a8.867 8.867 0 0 0 4.732-4.696l.022-.058a8.773 8.773 0 0 0 .705-3.477c0-4.298-3.038-7.886-7.083-8.736l-.058-.01V0c5.849.899 10.275 5.895 10.275 11.926v.012v-.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Clock(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 0C5.373 0 0 5.373 0 12s5.373 12 12 12s12-5.373 12-12C23.98 5.381 18.619.02 12.002 0zm0 21.6c-5.302 0-9.6-4.298-9.6-9.6S6.698 2.4 12 2.4s9.6 4.298 9.6 9.6c-.016 5.296-4.304 9.584-9.598 9.6zM12.6 6h-1.8v7.2l6.24 3.84l.96-1.56l-5.4-3.24z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Close(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.48 3.512a11.966 11.966 0 0 0-8.486-3.514C5.366-.002-.007 5.371-.007 11.999c0 3.314 1.344 6.315 3.516 8.487A11.966 11.966 0 0 0 11.995 24c6.628 0 12.001-5.373 12.001-12.001c0-3.314-1.344-6.315-3.516-8.487m-1.542 15.427a9.789 9.789 0 0 1-6.943 2.876c-5.423 0-9.819-4.396-9.819-9.819a9.789 9.789 0 0 1 2.876-6.943a9.786 9.786 0 0 1 6.942-2.876c5.422 0 9.818 4.396 9.818 9.818a9.785 9.785 0 0 1-2.876 6.942z"/><path fill="currentColor" d="m13.537 12l3.855-3.855a1.091 1.091 0 0 0-1.542-1.541l.001-.001l-3.855 3.855l-3.855-3.855A1.091 1.091 0 0 0 6.6 8.145l-.001-.001l3.855 3.855l-3.855 3.855a1.091 1.091 0 1 0 1.541 1.542l.001-.001l3.855-3.855l3.855 3.855a1.091 1.091 0 1 0 1.542-1.541l-.001-.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloseA(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24 2.4L21.6 0L12 9.6L2.4 0L0 2.4L9.6 12L0 21.6L2.4 24l9.6-9.6l9.6 9.6l2.4-2.4l-9.6-9.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudDown(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.947 17.678a.91.91 0 0 1 0-1.82h2.462a2.646 2.646 0 1 0-1.711-4.665l.004-.003a.911.911 0 0 1-1.178-1.389l.001-.001a4.45 4.45 0 0 1 2.095-.985l.027-.004A7.033 7.033 0 0 0 8.794 4.907l-.016.024a6.533 6.533 0 0 1 1.9 1.06l-.013-.01a.911.911 0 0 1-.562 1.626a.903.903 0 0 1-.576-.206l.002.001a4.663 4.663 0 0 0-2.946-1.04h-.014h.001a4.746 4.746 0 0 0 0 9.492a.91.91 0 0 1 0 1.82a6.567 6.567 0 1 1 0-13.134c.102 0 .202.008.304.014c1.543-2.737 4.431-4.555 7.743-4.555a8.859 8.859 0 0 1 8.859 8.857v.017c1.968.494 3.403 2.247 3.403 4.336a4.465 4.465 0 0 1-4.464 4.464zm-6.445 0a.896.896 0 0 1-.181-.019l.006.001l-.019-.006a.84.84 0 0 1-.156-.048l.006.002c-.005 0-.009-.005-.014-.007a.972.972 0 0 1-.147-.082l.003.002a.894.894 0 0 1-.141-.115l-2.584-2.582a.91.91 0 1 1 1.287-1.288l1.03 1.026v-3.984a.91.91 0 0 1 1.82 0v3.992l1.03-1.03a.91.91 0 1 1 1.287 1.288l-2.584 2.583l-.021.018a.907.907 0 0 1-.096.078l-.003.002l-.037.026a.827.827 0 0 1-.132.071l-.006.002a.871.871 0 0 1-.144.045l-.006.001l-.021.006a.879.879 0 0 1-.16.015h-.016z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudRefresh(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.327 16.753a4.342 4.342 0 0 1 2.231-3.793l.022-.011l-.734-.525a.91.91 0 1 1 1.061-1.479l-.002-.002l2.292 1.639c.016.011.029.026.044.039l.026.022c.014.012.03.022.043.035a.79.79 0 0 1 .064.077l.002.002a.8.8 0 0 1 .053.07l.002.003a.6.6 0 0 1 .04.076l.002.004c.014.03.018.031.026.048l.006.011l.008.022a.742.742 0 0 1 .024.075l.001.005c.007.027.016.045.021.068v.021a.64.64 0 0 1 .009.083v.002c0 .028.006.05.007.076v.016a1.068 1.068 0 0 1-.01.102l.001-.005v.011a.235.235 0 0 1-.006.054v-.002v.011l-.455 2.363a.911.911 0 0 1-1.072.72l.006.001a.912.912 0 0 1-.719-1.072l-.001.006l.194-1.009a2.518 2.518 0 1 0 3.679 2.235a.91.91 0 0 1 1.82 0a4.34 4.34 0 1 1-8.68 0zm10.63.925a.91.91 0 0 1 0-1.82h2.462a2.646 2.646 0 1 0-1.711-4.665l.004-.003a.91.91 0 0 1-1.177-1.388l.001-.001a4.438 4.438 0 0 1 2.095-.986l.027-.004A7.034 7.034 0 0 0 8.805 4.906l-.016.024a6.533 6.533 0 0 1 1.9 1.06l-.013-.01a.91.91 0 0 1-.569 1.622a.908.908 0 0 1-.571-.201l.002.001a4.665 4.665 0 0 0-2.947-1.04h-.014h.001a4.746 4.746 0 0 0 0 9.492h.63a.91.91 0 0 1 0 1.82h-.64a6.566 6.566 0 1 1 0-13.132c.102 0 .202.008.304.013C8.415 1.818 11.303 0 14.615 0a8.86 8.86 0 0 1 8.859 8.858v.017c1.968.494 3.403 2.247 3.403 4.336a4.465 4.465 0 0 1-4.464 4.464z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudUp(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.494 17.679a.91.91 0 0 1 0-1.82h3.92a2.646 2.646 0 1 0-1.711-4.665l.004-.003a.911.911 0 0 1-1.178-1.389l.001-.001a4.442 4.442 0 0 1 2.095-.986l.027-.004A7.034 7.034 0 0 0 8.797 4.908l-.016.025a6.533 6.533 0 0 1 1.9 1.06l-.013-.01a.911.911 0 0 1-.562 1.626a.903.903 0 0 1-.576-.206l.002.001a4.665 4.665 0 0 0-2.947-1.04h-.014h.001a4.746 4.746 0 0 0 0 9.492h2.093a.91.91 0 0 1 0 1.82H6.569a6.567 6.567 0 1 1 0-13.134c.102 0 .202.008.305.013C8.417 1.818 11.305 0 14.617 0a8.86 8.86 0 0 1 8.859 8.858v.018c1.968.494 3.403 2.247 3.403 4.336a4.465 4.465 0 0 1-4.464 4.464zm-5.966-.91v-4.187l-1.029 1.03a.91.91 0 1 1-1.287-1.288l2.583-2.584a.914.914 0 0 1 .19-.142l.005-.002c.017-.01.034-.018.051-.026a.873.873 0 0 1 .206-.079l.006-.001h.01a.88.88 0 0 1 .235-.011h-.003h.01c.017 0 .034.002.05.006h-.001c.086.01.164.03.238.061l-.006-.002a.948.948 0 0 1 .142.076l-.003-.002c.014.008.025.018.038.026a.999.999 0 0 1 .098.08l.022.018l2.585 2.584a.91.91 0 1 1-1.287 1.288l-1.03-1.03v4.19a.91.91 0 0 1-1.82 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloudflare(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.462 18.152a1.94 1.94 0 0 0-.206-1.718l.005.007a1.737 1.737 0 0 0-1.376-.677h-.006l-11.257-.146H8.62a.208.208 0 0 1-.171-.09v-.001a.254.254 0 0 1-.027-.203v.002a.31.31 0 0 1 .264-.202h.001l11.353-.146a4.076 4.076 0 0 0 3.309-2.461l.01-.027l.655-1.687a.361.361 0 0 0 .017-.222v.002c-.765-3.332-3.704-5.78-7.216-5.78a7.397 7.397 0 0 0-6.987 4.968l-.016.052a3.325 3.325 0 0 0-5.303 2.307l-.001.014a3.49 3.49 0 0 0 .088 1.186l-.005-.024a4.726 4.726 0 0 0-4.59 4.722v.006c.002.244.019.481.05.715l-.003-.029c.017.108.108.19.219.192h20.776a.283.283 0 0 0 .265-.2l.001-.002zm3.584-7.233c-.101 0-.21 0-.311.008a.186.186 0 0 0-.164.127v.001l-.439 1.528a1.94 1.94 0 0 0 .206 1.718l-.005-.007c.321.413.817.677 1.376.677h.006l2.4.146h.002c.071 0 .134.036.171.09v.001a.258.258 0 0 1 .027.203v-.002a.31.31 0 0 1-.264.202h-.001l-2.496.146a4.064 4.064 0 0 0-3.315 2.455l-.01.027l-.182.467a.132.132 0 0 0 .122.183h8.588a.225.225 0 0 0 .216-.163v-.002a6.149 6.149 0 0 0-5.921-7.809h-.005z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cloudy(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.581 6.371H6.6c1.122 0 2.152.392 2.96 1.047l-.009-.007a.907.907 0 0 0 1.279-.14l.001-.002a.907.907 0 0 0-.14-1.279l-.002-.001a6.511 6.511 0 0 0-1.846-1.038l-.046-.015c1.287-1.888 3.424-3.113 5.848-3.12h.001a7.06 7.06 0 0 1 7.05 7.001v.003a4.472 4.472 0 0 0-2.134.997l.006-.005a.908.908 0 0 0-.106 1.282l-.001-.001a.908.908 0 0 0 1.281.106l-.001.001a2.64 2.64 0 0 1 1.711-.626a2.656 2.656 0 0 1 2.653 2.652a2.656 2.656 0 0 1-2.653 2.653H6.581a4.762 4.762 0 0 1-4.757-4.756a4.762 4.762 0 0 1 4.757-4.749zm0 11.335h15.874a4.48 4.48 0 0 0 4.474-4.473a4.482 4.482 0 0 0-3.38-4.334l-.031-.007v-.018A8.88 8.88 0 0 0 14.646.001a8.907 8.907 0 0 0-7.737 4.516l-.023.044c-.102-.005-.204-.013-.306-.013a6.585 6.585 0 0 0-6.578 6.576a6.585 6.585 0 0 0 6.577 6.578z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CloudyGusts(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.256 19.953a.829.829 0 1 1 1.658 0a2.394 2.394 0 0 0 2.39 2.39a2.862 2.862 0 0 0 2.859-2.857a3.4 3.4 0 0 0-3.396-3.396H5.913a5.978 5.978 0 0 1 0-11.956h.069h-.004c.093 0 .185.006.278.01C7.661 1.654 10.29 0 13.305 0c1.646 0 3.178.493 4.454 1.34l-.03-.019a.828.828 0 1 1-.914 1.382l.003.002a6.353 6.353 0 0 0-3.513-1.049A6.408 6.408 0 0 0 7.508 5.33l-.017.039c-.004.009-.011.015-.015.023a.815.815 0 0 1-.109.167l.001-.001l-.013.017a.8.8 0 0 1-.151.128l-.003.002c-.017.012-.036.022-.054.033a.75.75 0 0 1-.14.062l-.006.002a.386.386 0 0 1-.045.017l-.003.001a.759.759 0 0 1-.191.027h-.026c-.023 0-.047-.006-.073-.007h-.051a4.32 4.32 0 1 0-.643 8.592h17.797a5.06 5.06 0 0 1 5.054 5.053A4.522 4.522 0 0 1 24.306 24a4.052 4.052 0 0 1-4.05-4.047m6.038-6.963a.829.829 0 1 1 0-1.658h4.442a2.727 2.727 0 0 0 2.723-2.723a2.275 2.275 0 0 0-2.271-2.271a1.88 1.88 0 0 0-1.877 1.876a.829.829 0 1 1-1.658 0a3.538 3.538 0 0 1 3.534-3.534c2.17 0 3.93 1.758 3.932 3.928a4.385 4.385 0 0 1-4.38 4.38zm-8.279-1.602a.829.829 0 1 1 0-1.658h4.442a2.726 2.726 0 0 0 2.721-2.726a2.273 2.273 0 0 0-2.27-2.272a1.878 1.878 0 0 0-1.877 1.878a.829.829 0 1 1-1.658 0a3.538 3.538 0 0 1 3.534-3.534a3.933 3.933 0 0 1 3.928 3.928a4.385 4.385 0 0 1-4.38 4.38z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cocktail(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23.812.678c-.054-.034-.107-.06-.16-.087a4.492 4.492 0 0 0-2.105-.516h-.008a4.602 4.602 0 0 0-3.965 2.278l-.012.022a4.429 4.429 0 0 0-.582 1.847l-.001.017h-9.6L6.305 1.65A2.664 2.664 0 0 0 3.229.071l.018-.004L.001.805l.201.886L3.448.953a1.76 1.76 0 0 1 2.014 1.035L5.466 2l.932 2.24H.302l10.193 10.193v7.463c-2.367.168-4 .952-4 1.925v.181h9.824v-.181c0-.972-1.636-1.76-4-1.925v-7.463l6.283-6.283c.197.167.417.322.649.458l.022.012A4.575 4.575 0 0 0 23.841.692L23.819.68zm.999 3.507l-2.494.007l1.241-2.16a3.305 3.305 0 0 1 1.251 2.135zm-2.038-2.609l-1.241 2.166l-1.254-2.153a3.285 3.285 0 0 1 1.264-.248c.443 0 .866.086 1.253.243zm-4.111 1.429c.218-.379.496-.698.825-.955l.007-.005l1.254 2.153l-2.488.007c.056-.449.198-.855.41-1.216zM3.359 5.506H6.92l.49 1.187H4.547zm14.894 1.188h-.813V6.7H8.387l-.49-1.187h11.534zm2.045 1.026l1.241-2.166l1.254 2.153c-.373.158-.807.25-1.262.25c-.444 0-.868-.087-1.255-.246zm4.111-1.422a3.297 3.297 0 0 1-.824.955l-.007.005l-1.254-2.153l2.488-.006a3.35 3.35 0 0 1-.411 1.214z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cocoapods(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24 24H0V0h24zM9.789 7.2a4.876 4.876 0 0 0-3.691 1.497l-.001.002a4.967 4.967 0 0 0-1.296 3.43v-.004l-.001.118a4.807 4.807 0 0 0 5.028 4.802h-.01l.126.002a4.942 4.942 0 0 0 3.073-1.066l-.01.008a4.274 4.274 0 0 0 1.489-2.514l.004-.026h-2.3a2.324 2.324 0 0 1-2.364 1.646h.006h-.047a2.749 2.749 0 0 1-2.007-.867l-.001-.001a3.056 3.056 0 0 1-.78-2.104v-.052c0-.769.277-1.473.736-2.018l-.004.005a2.66 2.66 0 0 1 2.084-.903h-.004l.08-.001c.595 0 1.141.208 1.57.556l-.005-.004c.395.347.67.822.761 1.36l.002.014h2.26a4.304 4.304 0 0 0-1.409-2.744l-.004-.004a4.85 4.85 0 0 0-3.289-1.13h.008zm7.334.219l-1.382.592l1.731 4.049l-1.732 4.05l1.382.59l1.162-2.737l.006.013l.804-1.92z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Code(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10 19.674l-.874.874a.55.55 0 0 1-.4.175a.55.55 0 0 1-.4-.175L.18 12.402a.546.546 0 0 1 0-.8l8.139-8.149a.55.55 0 0 1 .4-.175a.55.55 0 0 1 .4.175l.874.874a.546.546 0 0 1 0 .8l-6.87 6.87l6.87 6.87c.107.1.174.242.174.4v.011a.55.55 0 0 1-.168.397zM20.329 1.023l-6.52 22.566a.571.571 0 0 1-.268.339l-.003.001a.484.484 0 0 1-.413.043l.003.001l-1.084-.297a.571.571 0 0 1-.339-.268l-.001-.003a.522.522 0 0 1-.043-.432l-.001.004L18.18.411a.571.571 0 0 1 .268-.339l.003-.001a.484.484 0 0 1 .413-.043l-.003-.001l1.084.297c.147.043.267.14.339.268l.001.003a.522.522 0 0 1 .043.432zm11.484 11.379l-8.146 8.145a.547.547 0 0 1-.804 0l-.874-.874a.547.547 0 0 1 0-.804l6.87-6.87l-6.87-6.87a.547.547 0 0 1 0-.804l.874-.874a.547.547 0 0 1 .804 0l8.146 8.146a.547.547 0 0 1 0 .804z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Codepen(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m2.893 15.656l8.076 5.384v-4.808l-4.474-2.986zm-.831-1.928L4.647 12l-2.585-1.728zm10.969 7.312l8.076-5.384l-3.6-2.41l-4.476 2.986zM12 14.438L15.643 12L12 9.562L8.357 12zm-5.505-3.68l4.474-2.99V2.96L2.893 8.344zM19.353 12l2.585 1.728v-3.456zm-1.848-1.246l3.6-2.41l-8.074-5.384v4.808zM24 8.344v7.342c0 .346-.18.651-.452.825l-.004.002l-10.969 7.313a1.039 1.039 0 0 1-1.156-.002l.004.002L.454 16.514a.979.979 0 0 1-.456-.827l.001-.032v.002v-7.343c0-.346.18-.651.452-.825l.004-.002L11.424.175a1.039 1.039 0 0 1 1.156.002l-.004-.002l10.969 7.312a.979.979 0 0 1 .456.827L24 8.346z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Coffeescript(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.507 3.734a5.713 5.713 0 0 0 3.296-1.164l-.014.01a6.51 6.51 0 0 1 3.602-1.214h.01a5.69 5.69 0 0 1 1.972.128l-.039-.008q.734.21.794.63q.09.394-.42.644a3.346 3.346 0 0 1-1.25.314h-.009a2.28 2.28 0 0 1-1.483-.186l.013.006a1.041 1.041 0 0 1-.537-.623l-.002-.007q-.75.06-.944.33a.72.72 0 0 0-.164.512v-.002q.09.42.96.75c.484.152 1.04.24 1.616.24c.286 0 .568-.022.842-.063l-.031.004h.032c.889 0 1.715-.275 2.396-.743l-.014.009a1.392 1.392 0 0 0 .523-1.311l.001.008q-.18-.9-1.513-1.513a7.696 7.696 0 0 0-3.804-.4l.042-.005q-3.117.296-3.972 1.319a4.191 4.191 0 0 1-2.942 1.199h-.043h.002a3.074 3.074 0 0 1-1.356-.095l.022.006a.824.824 0 0 1-.583-.534l-.002-.006a.475.475 0 0 1 .311-.492l.003-.001c.249-.114.539-.194.843-.224l.011-.001a4.283 4.283 0 0 1 .969.049l-.025-.003c.284.036.543.105.785.205l-.021-.007a.675.675 0 0 0 .284-.196l.001-.001a.257.257 0 0 0 .045-.227v.002q-.09-.33-.659-.465A4.486 4.486 0 0 0 9.656.535l.019-.002q-1.559.15-1.814.69a1.495 1.495 0 0 0-.196.874v-.005a1.957 1.957 0 0 0 1.26 1.315l.014.004c.561.234 1.212.37 1.896.37c.235 0 .466-.016.692-.047l-.026.003zm14.06 4.738q-.96.239-2.069.42t-2.338.296q-1.229.12-2.608.197t-2.848.075q-1.529 0-2.908-.075t-2.623-.21t-2.353-.314t-2.069-.394a14.18 14.18 0 0 1-2.782-.961l.087.037a3.257 3.257 0 0 1-1.284-.978l-.005-.007c.115.867.267 1.623.464 2.362l-.029-.128c.23.885.461 1.6.726 2.298l-.051-.153a5.762 5.762 0 0 0-.86.634l.005-.005a6.662 6.662 0 0 0-.755.796l-.01.013a5.56 5.56 0 0 0-.974 1.804l-.011.04a5.903 5.903 0 0 0-.27 1.971v-.009c.035.666.225 1.281.535 1.818l-.01-.02a5.179 5.179 0 0 0 1.229 1.424l.01.008a4.222 4.222 0 0 0 1.646.819l.029.006a4.675 4.675 0 0 0 1.888.01l-.03.005c.298-.048.565-.127.817-.233l-.023.009q.405-.165.794-.285h-.032a4.184 4.184 0 0 1-1.51-.28l.029.01a5.447 5.447 0 0 1-1.405-.817l.01.008a5.172 5.172 0 0 1-1.202-1.27l-.012-.019a3.867 3.867 0 0 1-.613-1.656l-.002-.019a3.91 3.91 0 0 1 .005-1.69l-.005.026a4.51 4.51 0 0 1 .669-1.529l-.01.015q.12-.12.225-.239t.225-.239q.27.69.584 1.349t.68 1.29q.719 1.109 1.469 2.143t1.498 2.08c.188.346.375.761.533 1.191l.022.068c.127.337.246.749.335 1.172l.01.057c.337.456.733.844 1.183 1.161l.016.011c.448.312.965.567 1.519.739l.04.011c.613.219 1.34.402 2.09.516l.067.008c.672.105 1.447.165 2.236.165h.162a15.227 15.227 0 0 0 2.395-.192l-.087.011a13.62 13.62 0 0 0 2.345-.568l-.096.028a5.864 5.864 0 0 0 1.503-.747l-.02.013a4.795 4.795 0 0 0 1.172-1.167l.01-.015h.09q.12-.54.314-1.199c.149-.496.327-.923.543-1.327l-.019.038q.719-1.049 1.469-2.069t1.499-2.158a21.885 21.885 0 0 0 1.616-3.676l.047-.161a28.18 28.18 0 0 0 1.036-4.041l.028-.186a3.786 3.786 0 0 1-1.326 1.025l-.023.01a10.53 10.53 0 0 1-2.566.841l-.068.011zM5.752 6.402q.96.239 2.069.42t2.338.296q1.229.12 2.593.18t2.857.06q1.529 0 2.893-.06t2.593-.18t2.323-.314t2.083-.405a11.75 11.75 0 0 0 3.105-1.14l-.062.031q1.004-.599 1.004-1.199a1.17 1.17 0 0 0-.432-.853l-.002-.002a4.085 4.085 0 0 0-1.216-.725l-.028-.009c.12.083.219.188.293.31l.003.004a.756.756 0 0 1 .12.405q0 .63-.929 1.139c-.804.402-1.737.723-2.716.917l-.071.012q-1.739.36-3.957.615c-1.41.163-3.044.256-4.7.256l-.227-.001h.012h-.163c-1.658 0-3.296-.087-4.909-.257l.201.017a38.4 38.4 0 0 1-4.192-.643l.25.044a11.901 11.901 0 0 1-2.797-.991l.068.031q-.929-.509-.929-1.109V3.25c0-.187.05-.362.138-.513l-.003.005c.127-.196.292-.357.486-.476l.007-.004a8.818 8.818 0 0 0-1.758.875l.035-.021c-.341.2-.57.559-.584.973v.002q.06.63 1.094 1.229c.897.474 1.936.849 3.031 1.071l.075.013z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Columns(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M28 0H1.333C.596 0-.001.597-.001 1.334v21.333c0 .737.597 1.334 1.334 1.334H28c.737 0 1.334-.597 1.334-1.334V1.334C29.334.597 28.737 0 28 0M2.664 8.001h10.669v13.333H2.666zM16 21.333V8h10.667v13.333z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Comment(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14 2l-.128-.001c-2.098 0-4.102.399-5.942 1.124l.11-.038a11.285 11.285 0 0 0-4.4 2.922l-.007.007A5.884 5.884 0 0 0 2 9.994v.005a5.677 5.677 0 0 0 1.131 3.351l-.011-.015a9.584 9.584 0 0 0 3.096 2.719l.049.025l1.36.782l-.426 1.498A11.174 11.174 0 0 1 6.077 21.1l.029-.054a15.466 15.466 0 0 0 4.313-2.686l-.017.014l.672-.594l.89.094a17.07 17.07 0 0 0 2.028.125h.004l.128.001c2.098 0 4.102-.399 5.942-1.124l-.11.038a11.285 11.285 0 0 0 4.4-2.922l.007-.007c1.009-1.025 1.632-2.432 1.632-3.984s-.623-2.96-1.633-3.985l.001.001a11.245 11.245 0 0 0-4.329-2.904l-.078-.025c-1.73-.687-3.735-1.086-5.833-1.086l-.132.001h.007zm14 8a7.762 7.762 0 0 1-1.884 5.033l.009-.01a12.716 12.716 0 0 1-5.008 3.611l-.086.03c-2.023.846-4.374 1.337-6.839 1.337L13.99 20H14c-.8-.002-1.588-.047-2.363-.134l.097.009a17.021 17.021 0 0 1-7.069 3.756l-.118.026c-.503.145-1.107.266-1.726.339l-.055.005h-.08a.621.621 0 0 1-.422-.164a.81.81 0 0 1-.249-.424l-.001-.005v-.016a.201.201 0 0 1-.027-.102c0-.033.008-.063.021-.091l-.001.001a.406.406 0 0 0 .031-.159v-.002q-.008-.031.07-.149l.094-.141l.11-.133l.125-.141q.11-.125.484-.539l.539-.594q.164-.18.484-.617c.174-.231.343-.493.491-.767l.017-.033q.187-.359.422-.922c.137-.317.276-.712.39-1.117l.017-.07a11.573 11.573 0 0 1-3.844-3.405l-.024-.035A7.524 7.524 0 0 1-.001 9.999v-.002a7.762 7.762 0 0 1 1.884-5.033l-.009.01a12.716 12.716 0 0 1 5.008-3.611l.086-.03C8.991.487 11.342-.004 13.807-.004l.202.001h-.01l.192-.001c2.465 0 4.816.491 6.959 1.381l-.12-.044a12.734 12.734 0 0 1 5.078 3.622l.015.018a7.754 7.754 0 0 1 1.875 5.021v.003z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Commenting(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10 10a2 2 0 1 1-.586-1.414c.361.351.586.841.586 1.384v.032zm6 0a2 2 0 0 1-3.414 1.414a2 2 0 1 1 2.827-2.829l.001.001c.362.351.587.841.587 1.385v.031V10zm6 0a2 2 0 1 1-.586-1.414c.361.351.586.841.586 1.384v.032zm-8-8l-.128-.001c-2.098 0-4.102.399-5.942 1.124l.11-.038a11.285 11.285 0 0 0-4.4 2.922l-.007.007A5.884 5.884 0 0 0 2 9.994v.005a5.677 5.677 0 0 0 1.131 3.351l-.011-.015a9.584 9.584 0 0 0 3.096 2.719l.049.025l1.36.782l-.426 1.498A11.174 11.174 0 0 1 6.077 21.1l.029-.054a15.466 15.466 0 0 0 4.313-2.686l-.017.014l.672-.594l.89.094a17.07 17.07 0 0 0 2.028.125h.004l.128.001c2.098 0 4.102-.399 5.942-1.124l-.11.038a11.285 11.285 0 0 0 4.4-2.922l.007-.007c1.009-1.025 1.632-2.432 1.632-3.984s-.623-2.96-1.633-3.985l.001.001a11.245 11.245 0 0 0-4.329-2.904l-.078-.025c-1.73-.687-3.735-1.086-5.833-1.086l-.132.001h.007zm14 8a7.762 7.762 0 0 1-1.884 5.033l.009-.01a12.716 12.716 0 0 1-5.008 3.611l-.086.03c-2.023.846-4.374 1.337-6.839 1.337L13.99 20H14c-.8-.002-1.588-.047-2.363-.134l.097.009a17.021 17.021 0 0 1-7.069 3.756l-.118.026c-.503.145-1.107.266-1.726.339l-.055.005h-.08a.621.621 0 0 1-.422-.164a.81.81 0 0 1-.249-.424l-.001-.005v-.016a.201.201 0 0 1-.027-.102c0-.033.008-.063.021-.091l-.001.001a.406.406 0 0 0 .031-.159v-.002q-.008-.031.07-.149l.094-.141l.11-.133l.125-.141q.11-.125.484-.539l.539-.594q.164-.18.484-.617c.174-.231.343-.493.491-.767l.017-.033q.187-.359.422-.922c.137-.317.276-.712.39-1.117l.017-.07a11.573 11.573 0 0 1-3.844-3.405l-.024-.035A7.524 7.524 0 0 1-.001 9.999v-.002a7.405 7.405 0 0 1 1.128-3.914l-.018.031A10.74 10.74 0 0 1 4.06 2.941l.033-.022A14.994 14.994 0 0 1 8.454.814l.108-.027c1.603-.501 3.446-.789 5.357-.789h.085H14h.083c1.911 0 3.754.288 5.488.824l-.132-.035a15.195 15.195 0 0 1 4.517 2.165l-.048-.032a10.75 10.75 0 0 1 2.958 3.149l.026.046a7.36 7.36 0 0 1 1.107 3.882v.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Comments(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 2.182h-.105c-1.715 0-3.354.325-4.859.918l.09-.031a9.179 9.179 0 0 0-3.599 2.394l-.006.006A4.827 4.827 0 0 0 2.183 8.72v.005a4.619 4.619 0 0 0 .912 2.705l-.009-.012a7.892 7.892 0 0 0 2.5 2.229l.039.02l1.654.958l-.597 1.432q.579-.341 1.057-.665l.75-.528l.903.17c.783.15 1.684.237 2.606.24h.107c1.715 0 3.354-.325 4.859-.918l-.09.031a9.179 9.179 0 0 0 3.599-2.394l.006-.006c.827-.836 1.338-1.986 1.338-3.256s-.511-2.42-1.338-3.256a9.152 9.152 0 0 0-3.542-2.379l-.063-.021a13.04 13.04 0 0 0-4.79-.895h-.09h.005zM12 0l.151-.001c2.119 0 4.138.429 5.974 1.206l-.101-.038a10.884 10.884 0 0 1 4.358 3.161l.014.018c.996 1.174 1.602 2.706 1.602 4.38s-.606 3.207-1.61 4.39l.008-.01a10.866 10.866 0 0 1-4.299 3.152l-.073.026c-1.736.739-3.756 1.169-5.875 1.169l-.157-.001H12a17.26 17.26 0 0 1-3.106-.289l.106.016a14.174 14.174 0 0 1-4.64 2.158l-.1.022c-.389.1-.886.195-1.391.264l-.074.008h-.051a.535.535 0 0 1-.35-.136a.56.56 0 0 1-.196-.355v-.003a.333.333 0 0 1-.017-.107v-.006c0-.038.003-.076.009-.113l-.001.004a.375.375 0 0 1 .035-.104l-.001.002l.042-.086l.06-.094l.068-.086l.08-.086l.068-.08q.086-.102.392-.426t.443-.503t.383-.494c.152-.191.293-.406.415-.633l.011-.023q.179-.341.35-.75a10.034 10.034 0 0 1-3.303-2.985l-.021-.032A6.644 6.644 0 0 1-.003 8.724v-.002a6.852 6.852 0 0 1 1.609-4.391l-.009.011A10.866 10.866 0 0 1 5.896 1.19l.073-.026C7.705.425 9.725-.005 11.844-.005l.161.001h-.008zm14.01 19.925q.17.409.35.75c.133.25.274.465.433.665l-.007-.009q.247.315.383.494t.443.503q.306.32.392.426q.017.017.068.08t.08.086a.75.75 0 0 1 .066.083l.002.002a.87.87 0 0 1 .058.09l.002.004l.042.086l.034.102l.009.11l-.017.11a.634.634 0 0 1-.22.374l-.001.001a.513.513 0 0 1-.377.119h.002a13.591 13.591 0 0 1-6.243-2.482l.039.027c-.901.171-1.938.27-2.998.273h-.002a14.253 14.253 0 0 1-8.101-2.283l.056.034q.989.068 1.5.068h.067c1.855 0 3.644-.28 5.327-.801l-.128.034a15.069 15.069 0 0 0 4.544-2.229l-.044.03a11.487 11.487 0 0 0 3.24-3.563l.029-.055a8.67 8.67 0 0 0 1.142-4.327c0-.926-.143-1.818-.409-2.655l.017.062a10.181 10.181 0 0 1 3.457 3.004l.02.029a6.598 6.598 0 0 1 1.278 3.921a6.62 6.62 0 0 1-1.224 3.846l.014-.021a10.069 10.069 0 0 1-3.275 2.983l-.05.026z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Compass(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 12C0 5.373 5.373 0 12 0s12 5.373 12 12s-5.373 12-12 12S0 18.627 0 12m2.017 0c0 5.513 4.469 9.983 9.983 9.983s9.983-4.469 9.983-9.983c0-5.513-4.469-9.983-9.983-9.983c-5.513 0-9.983 4.469-9.983 9.983m8.278-.928v-.002a2.019 2.019 0 0 1 .219-.302l-.002.002c.02-.024.041-.046.062-.068c.071-.076.147-.146.228-.209l.004-.003c.027-.021.056-.037.083-.057s.075-.054.115-.078l6.483-3.795l-3.783 6.464v.002a1.975 1.975 0 0 1-.22.302l.002-.002c-.02.024-.041.046-.061.068a1.955 1.955 0 0 1-.229.208l-.004.003c-.027.021-.055.038-.083.057s-.075.054-.115.078l-6.483 3.796z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CompassAlt(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.983 11.186l1.482-2.65l-2.63 1.482l-.814-2.922l-.814 2.922l-2.651-1.482l1.461 2.65L7.095 12l2.922.814l-1.482 2.651l2.65-1.482l.814 2.922l.814-2.922l2.65 1.482l-1.482-2.63l2.922-.814zm-1.148.981c0 .021-.021.042-.021.062s-.021.062-.021.08c-.021.021-.021.042-.042.08c-.021.021-.021.042-.042.062a.464.464 0 0 1-.08.104l-.021.021c-.042.042-.08.062-.126.104c-.021.021-.042.021-.062.042s-.049.035-.079.041h-.001a.128.128 0 0 1-.07.021h-.011c-.021 0-.062.021-.08.021c-.042 0-.08.021-.126.021h-.08l-.021.001a.264.264 0 0 1-.106-.022l.002.001c-.021 0-.062-.021-.08-.021s-.062-.021-.08-.021c-.021-.021-.062-.021-.08-.042s-.042-.021-.062-.042a.55.55 0 0 1-.125-.104l-.021-.021c-.042-.042-.062-.08-.104-.126c-.021-.021-.021-.042-.042-.062s-.021-.062-.042-.08a.128.128 0 0 1-.021-.07v-.011c0-.021-.021-.042-.021-.062a.675.675 0 0 1 .001-.339l-.001.005c0-.021.021-.042.021-.062c0-.042.021-.062.042-.104c0-.021.021-.042.021-.062c.021-.021.021-.062.042-.08a.292.292 0 0 0 .062-.078l.001-.002c.042-.042.062-.08.104-.104c.021-.021.042-.042.062-.042c.042-.021.062-.042.104-.062c.021 0 .021-.021.042-.021c.042-.021.08-.021.126-.042h.042a.675.675 0 0 1 .339.001l-.005-.001c.021 0 .042.021.062.021c.042 0 .062.021.104.042c.021.021.042.021.08.042c.021.021.042.021.062.042a.91.91 0 0 1 .228.227l.002.003c.021.021.021.042.042.062s.035.049.041.079v.001c.013.02.021.044.021.07v.011c0 .021.021.042.021.08c.014.042.022.09.022.139l-.001.029v-.001l.001.035a.397.397 0 0 1-.024.138l.001-.003zm-1.231-6.344c0-.313 0-.606-.021-.856c.104.23.23.48.334.689l.418.751h.459V4.445h-.418v.584a5.06 5.06 0 0 0 .045.84l-.003-.026a6.184 6.184 0 0 0-.33-.702l.017.034l-.418-.73h-.522v1.962h.418v-.584zm-5.51 6.01c-.042.23-.08.459-.104.689c-.021-.23-.062-.438-.104-.668l-.171-.834h-.48l-.16.814c-.042.25-.104.48-.126.689c-.042-.209-.08-.459-.126-.689l-.149-.814h-.48l.459 1.962h.48l.188-.856a4.79 4.79 0 0 0 .102-.581l.002-.024c.021.23.062.418.08.606l.167.856h.48l.501-1.962h-.4zm12.96.313h.71v-.354h-.71v-.418h.751v-.355h-1.21v1.962h1.252v-.354h-.793zm-6.888 6.365c-.25-.104-.376-.146-.376-.271c0-.104.08-.188.292-.188c.151.006.294.035.428.083l-.01-.003l.104-.354a1.467 1.467 0 0 0-.498-.104h-.003c-.459 0-.751.25-.751.584c0 .292.209.459.542.584c.23.08.334.146.334.271s-.104.209-.313.209a1.176 1.176 0 0 1-.51-.128l.007.003l-.08.376c.16.079.348.126.546.126h.014h-.001c.542 0 .793-.271.793-.606c.007-.268-.163-.456-.518-.582"/><path fill="currentColor" d="M12 0C5.373 0 0 5.373 0 12s5.373 12 12 12s12-5.373 12-12C23.991 5.376 18.624.009 12.001 0zm10.059 13.106c-.534 4.693-4.205 8.386-8.842 8.948l-.049.005l-1.169-2.025l-1.169 2.024c-4.67-.57-8.331-4.238-8.886-8.862l-.005-.048L3.901 12l-1.962-1.127c.559-4.672 4.221-8.341 8.841-8.906l.049-.005l1.169 2.025l1.169-2.024c4.685.567 8.356 4.259 8.886 8.905l.004.047l-1.899 1.106z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Composer(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.558 23.914a.21.21 0 0 1 .022-.231c.046-.065.098-.122.149-.182c.238-.299.47-.635.68-.985l.024-.042c.096-.156.194-.292.301-.418l-.004.005a.855.855 0 0 1 .092-.089l.001-.001c.14-.108.264-.061.296.112a1.611 1.611 0 0 1-.011.442l.001-.009c-.011.11-.026.218-.036.328c-.004.032-.017.072.014.086s.061-.026.08-.05l.367-.456c.124-.16.253-.302.391-.434l.001-.001c.074-.067.154-.125.262-.072s.115.14.12.24c-.01.267-.047.52-.109.763l.005-.025a.998.998 0 0 0-.017.39l-.001-.006c.007.045.03.084.062.112a.207.207 0 0 1-.018.343l-.001.001a.432.432 0 0 1-.28.069h.002a.212.212 0 0 1-.187-.194v-.001a2.228 2.228 0 0 1 .013-.579l-.002.012a.926.926 0 0 0-.018-.343l.001.006c-.12.16-.235.302-.349.447c-.106.132-.214.25-.33.361l-.001.001a.908.908 0 0 1-.199.149l-.005.002a.198.198 0 0 1-.305-.151v-.001a1.482 1.482 0 0 1 .012-.422l-.001.009c0-.038.006-.074.011-.111c-.036-.01-.04.019-.053.034c-.209.283-.418.569-.626.853c-.043.058-.086.118-.169.122h-.028a.188.188 0 0 1-.157-.085v-.001zm11.063-.074c-.199-.1-.364-.24-.492-.409l-.003-.004c-.142-.176-.283-.35-.432-.538a3.98 3.98 0 0 0-.109.577l-.002.019c-.026.204-.106.295-.264.32c-.176.026-.264-.06-.223-.228c.119-.506.247-1.008.371-1.512a.406.406 0 0 1 .104-.187a.212.212 0 0 1 .296-.039c.048.032.065.005.094-.017a.534.534 0 0 1 .437-.144h-.003a.374.374 0 0 1 .338.372v.008a.738.738 0 0 1-.318.709l-.003.002c-.12.08-.115.08-.026.202c.108.151.227.282.36.398l.003.002a.25.25 0 0 0 .15.07h.005a.16.16 0 0 1 .156.124v.001a.19.19 0 0 1-.055.206a.36.36 0 0 1-.239.106h-.001a.33.33 0 0 1-.148-.037l.002.001zm-.471-1.76a.376.376 0 0 0-.313.237l-.001.003a1.469 1.469 0 0 0-.11.417l-.001.007c.199-.076.366-.197.495-.351l.002-.002a.39.39 0 0 0 .069-.129l.001-.003c.032-.106-.017-.178-.121-.178zM4.194 23.385c-.002-.027-.002-.058-.002-.09c0-.5.213-.949.553-1.263l.001-.001a.75.75 0 0 1 .67-.218l-.005-.001c.038.005.072.012.091-.038c.031-.09.105-.084.173-.058a.66.66 0 0 1 .341.535v.002a1.706 1.706 0 0 1-1.242 1.584l-.012.003a.668.668 0 0 1-.16.019c-.25-.004-.397-.174-.409-.474zm.711-1.023a1.487 1.487 0 0 0-.324.856v.004l-.001.025a.213.213 0 0 0 .255.209h-.001a.583.583 0 0 0 .325-.141l-.001.001c.233-.203.39-.49.422-.813v-.005c.005-.16-.033-.218-.144-.262a.51.51 0 0 0-.187-.037a.44.44 0 0 0-.342.163l-.001.001zm5.183 1.016a1.726 1.726 0 0 1 .544-1.349l.001-.001a.815.815 0 0 1 .333-.2l.006-.002a.24.24 0 0 0 .055-.02l-.001.001c.353-.213.708-.432 1.082-.66l-.56-.22c-.039-.014-.058.014-.08.033c-.22.172-.416.346-.6.532l-.001.001a5.8 5.8 0 0 1-.598.556l-.011.009a.148.148 0 0 0-.062.135v-.001a.658.658 0 0 1-.248.469l-.001.001a1.612 1.612 0 0 1-.517.321l-.011.004c-.133.053-.262.108-.394.16a.07.07 0 0 0-.054.067a5.714 5.714 0 0 1-.053.327a.278.278 0 0 1-.355.22l.002.001a.177.177 0 0 1-.119-.218v.001c.169-.523.313-1.056.53-1.562c.031-.072.103-.086.16-.137a1.482 1.482 0 0 0-.351-.103l-.009-.001a102.764 102.764 0 0 0-3.078-.617c-.218-.04-.444-.031-.667-.05c-.619-.053-1.241-.026-1.861-.04c-.25-.005-.497 0-.746-.012a.567.567 0 0 1-.117-.013l.004.001a.08.08 0 0 1-.059-.087c0-.029.039-.024.062-.026c.492-.054.986-.118 1.481-.151c.204-.014.409-.012.614-.01h.435c.493.002.976.037 1.449.103l-.056-.006c.16.017.324.029.487.043c.014.011.031.036.05.024s0-.031-.01-.046v-.01a.132.132 0 0 0-.075-.12h-.001a2.316 2.316 0 0 0-.667-.238l-.015-.002q-1.052-.32-2.122-.574a6.187 6.187 0 0 0-.971-.151l-.022-.001a.114.114 0 0 0-.099.015q.416.334.833.666c-.04.036-.074.01-.108 0c-.2-.083-.37-.173-.531-.277l.013.008a4.6 4.6 0 0 0-.774-.378l-.033-.011a.442.442 0 0 1-.082-.037l.002.001c-.048-.029-.038-.075-.031-.118s.047-.029.074-.029l.125-.001c.488 0 .969.025 1.444.075l-.059-.005c.314.026.622.106.932.16c.032.005.062.011.094.014c.046.005.118.048.133 0a.197.197 0 0 0-.048-.209a2.934 2.934 0 0 0-.226-.181l-.009-.006a.091.091 0 0 1-.049-.08c0-.02.007-.039.018-.054c.033-.062.054-.134.089-.198s.022-.101-.038-.146l-1.954-1.44c-.248-.182-.493-.364-.739-.542a.168.168 0 0 1-.08-.161v.001a.413.413 0 0 0-.039-.221l.001.003a.284.284 0 0 0-.183-.16h-.002c-.16-.05-.314-.115-.473-.166c-.067-.022-.08-.058-.08-.122c-.012-.262.033-.522.024-.782c-.019-.536-.05-1.069-.06-1.6c-.005-.202 0-.4-.012-.606c-.019-.415-.019-.83-.026-1.246v-.007c0-.042.013-.08.034-.112v.001a.19.19 0 0 0-.038-.209l-.199-.264a.174.174 0 0 1-.033-.161v.001c.048-.223.08-.451.126-.678c.01-.053-.022-.054-.058-.06c-.192-.029-.385-.06-.579-.084c-.065-.007-.072-.031-.067-.09q.093-.998.182-1.995v-.01c.014-.126.014-.126.137-.094q.417.105.833.212c.053.014.08.014.091-.055c.048-.274.114-.516.201-.749l-.009.029a.17.17 0 0 0 .01-.061a.177.177 0 0 0-.017-.077v.001c-.127-.293-.255-.583-.379-.88c-.022-.05-.039-.058-.086-.029s-.106.11-.16.08c-.043-.022-.055-.108-.08-.166L.84 6.452c-.037-.075-.098-.156-.085-.226s.12-.086.187-.125c.053-.031.031-.065.019-.103C.817 5.556.668 5.12.535 4.676S.269 3.77.153 3.312a5.848 5.848 0 0 1-.148-.775l-.003-.032c.002-.119.023-.233.06-.339l-.002.008c.01-.041.043-.034.069-.036a.38.38 0 0 1 .357.159l.001.001c.076.098.15.209.216.325l.007.014l.64 1.12c.033.06.047.053.084 0c.191-.266.355-.547.526-.828c.148-.24.348-.45.506-.682A.425.425 0 0 1 2.73 2.1h.002l.016-.002c.018 0 .034.006.046.016c.106.061.197.13.279.209c.036.046.074.032.118.019a10.373 10.373 0 0 0 2.204-.877l-.055.027c.061-.035.134-.071.209-.101l.012-.004c.154-.062.266.026.245.198a.6.6 0 0 1-.187.369c-.266.252-.53.511-.818.742a.113.113 0 0 0-.054.095c0 .015.003.03.009.043v-.001c.089.332.178.602.281.864l-.019-.055c.05.113.05.113.151.05q1.722-1.07 3.44-2.142c.115-.07.115-.07.199.029c.011.016.023.03.036.043c.058.05.033.08-.019.11c-.414.264-.821.53-1.232.795q-.506.33-1.015.658c-.026.017-.058.026-.08.065a.463.463 0 0 0 .088.014h.015a.31.31 0 0 1 .263.145l.001.001a.467.467 0 0 1 .152.304v.002l.001.024a.356.356 0 0 1-.11.258a9.225 9.225 0 0 0-.762.771l-.009.01a1.126 1.126 0 0 0-.098.143l-.003.005a17.365 17.365 0 0 0-.931 1.7l-.048.112a.843.843 0 0 1-.323.331l-.004.002a6.943 6.943 0 0 1-.53.288a.205.205 0 0 0-.139.215v-.001v.013a.339.339 0 0 1-.159.288l-.001.001a.18.18 0 0 0-.055.085v.001a7.79 7.79 0 0 0-.229.842l-.011.059c-.022.08.012.103.084.12c.382.09.763.191 1.145.283c.08.019.09.046.07.12c-.103.389-.221.773-.302 1.17l-.068.312c-.017.074-.053.089-.119.08c-.09-.012-.176-.017-.264-.026c-.031-.004-.06-.007-.053.039c.038.245.08.49.119.744a.159.159 0 0 0 .069-.095v-.001c.224-.484.486-.9.791-1.28l-.009.012c.278-.334.588-.64.898-.943c.293-.264.628-.491.99-.669l.026-.011c.072-.04.151-.065.22-.108c.33-.181.722-.289 1.14-.292h.001a4.059 4.059 0 0 1 .467-.017h-.004c.314.014.609.075.885.175l-.021-.007c.38.144.707.341.991.587l-.004-.003c.11.108.206.23.284.364l.005.008a.125.125 0 0 0 .161.06h-.001a2.196 2.196 0 0 1 1.135-.001l-.015-.003c.129.024.243.068.346.13l-.005-.003c.266.137.496.297.701.483l-.003-.003c.039.036.075.074.109.113l.002.002c.204.239.407.504.598.778l.022.034c.135.206.244.443.316.697l.004.018c.04-.122.08-.245.122-.367c.004-.014-.018-.038-.029-.058a28.887 28.887 0 0 0-.3-.502a.091.091 0 0 1-.02-.056c0-.024.009-.046.025-.062c.098-.163.192-.3.293-.432l-.008.011c.036-.04 0-.094-.024-.13s-.055 0-.08.014c-.233.12-.24.12-.365-.11l-.91-1.64c-.068-.119-.065-.122.067-.16c.37-.113.739-.226 1.112-.334c.072-.022.096-.043.07-.122q-.111-.356-.212-.716a.146.146 0 0 0-.093-.108h-.001c-.022-.007-.038-.022-.06-.029a.394.394 0 0 1-.305-.327V6.6a.116.116 0 0 0-.093-.093h-.001c-.05-.011-.098-.038-.151-.05a3.861 3.861 0 0 1-1.231-.591l.011.008a3.586 3.586 0 0 1-.364-.303l.001.001a.51.51 0 0 0-.115-.079l-.003-.001a9.194 9.194 0 0 0-1.42-.499l-.067-.015c-.257-.065-.226-.182-.16-.391a.455.455 0 0 1 .392-.281h.001a5.894 5.894 0 0 1 1.229-.204l.017-.001c.065 0 .08-.014.043-.074a1.124 1.124 0 0 0-.261-.354l-.001-.001a1.343 1.343 0 0 1-.114-.112l-.001-.001a8.265 8.265 0 0 0-1.339-1.156l-.027-.018a.447.447 0 0 1-.223-.309v-.003a.282.282 0 0 1 .09-.252a.45.45 0 0 1 .393-.144h-.002a.937.937 0 0 1 .464.172l-.003-.002c.3.195.595.4.896.598c.06.04.126.074.187.11c0 0 0 .005.005.007a.061.061 0 0 0 .036.022l.01.001l.01-.001a.095.095 0 0 0-.029-.039l-.012-.008c-.43-.429-.86-.859-1.304-1.274a.694.694 0 0 1-.23-.627V.64a.274.274 0 0 1 .196-.192h.002c.089-.007.134-.032.134-.14V.301a.29.29 0 0 1 .358-.282h-.002l.384.04a.579.579 0 0 1 .217.066l-.003-.002c.585.29 1.077.58 1.546.9l-.044-.028a6.71 6.71 0 0 1 1 .806L14.02 1.8c.01.025.034.043.062.043c-.017-.029-.029-.048-.038-.068c-.166-.42-.338-.84-.529-1.25a.408.408 0 0 1 .071-.43a.204.204 0 0 1 .198-.095h-.001a.697.697 0 0 1 .226.05l-.005-.002c.187.097.337.246.433.427l.003.005c.263.414.509.842.737 1.274c.198.364.369.786.49 1.228l.01.042c.119.446.209.973.252 1.514l.002.032c.001.011.002.02.005.03v-.001a1.49 1.49 0 0 1-.022.577l.002-.01a.398.398 0 0 0-.006.322l-.001-.003c.19.533.37 1.07.601 1.589c.062.137.065.137.204.094c.457-.142.91-.278 1.36-.422c.086-.026.101-.01.094.075a535.23 535.23 0 0 0-.139 1.818v.01a.135.135 0 0 1-.105.131h-.001c-.142.043-.278.104-.418.151c-.062.022-.067.044-.012.08c.08.086.154.178.234.262c.096.103.182.216.271.326c.08.098.072.139-.018.23c-.036.039-.08.08-.043.14a.964.964 0 0 1 .061.722l.002-.007a.404.404 0 0 0 .013.208l-.001-.003c.023.106.051.196.084.284l-.004-.013a.66.66 0 0 1-.006.463l.002-.005c-.152.503-.323.93-.526 1.339l.022-.049c-.127.314-.247.631-.372.948a5.396 5.396 0 0 0-.254 1.08l-.003.029c-.05.296-.155.577-.187.874a8.699 8.699 0 0 1-.074.581a1.408 1.408 0 0 1-.102.362l.004-.009l-.298.674a.64.64 0 0 0-.058.247v.001c-.007.197-.036.389-.034.586c0 .058 0 .086.08.068c.245-.058.493-.104.739-.154c.022-.005.053-.019.067.01s-.017.036-.033.048a4.214 4.214 0 0 0-.699.635l-.004.005c-.111.12-.223.24-.336.358a.06.06 0 0 1-.068.018c.107-.252.231-.469.378-.669l-.007.009c.005 0 .007 0 .01-.005a.02.02 0 0 0 .008-.016c0-.004-.001-.008-.003-.011s0 0-.005 0h-.002c-.005 0-.008.005-.01.01c-.545.221-1.083.466-1.625.692c-.66.274-1.294.605-1.926.938c-.029.014-.046.029-.041.062l.055.504a.98.98 0 0 0 .084-.038l-.004.002c.714-.386 1.453-.71 2.194-1.034c.48-.211.99-.353 1.481-.536c.205-.074.444-.13.691-.159l.015-.001c.004.006.011.011.014.018c-.168.068-.326.262-.486.262a1.271 1.271 0 0 0-.165.089l.005-.003a.707.707 0 0 0-.196.185l-.001.002a.09.09 0 0 0-.074.05v.001s-.019.014-.019.022h-.004v.005c-.05-.014-.24.187-.302.25l.024-.001l.025.001h-.001h.012a.227.227 0 0 0 .083-.015l-.002.001l.017-.008c.253-.065.551-.112.856-.133l.016-.001c.069-.005.137-.012.206-.019l.075.017c-1.064.274-2.125.56-3.164.918c-.951.304-1.77.666-2.54 1.105l.065-.034l-.038.026c.062.06.12.124.174.19l.003.004a.601.601 0 0 1 .079.518l.001-.004a1.711 1.711 0 0 1-1.219 1.375l-.012.003a.722.722 0 0 1-.16.02c-.258-.011-.405-.178-.417-.481zm.709-1.028a1.49 1.49 0 0 0-.331.877v.003l-.001.022a.21.21 0 0 0 .25.206h-.001a.567.567 0 0 0 .321-.137l-.001.001c.233-.204.39-.49.424-.813v-.006c0-.169-.036-.223-.148-.27a.508.508 0 0 0-.182-.034a.435.435 0 0 0-.327.151v.001zm-1.525-.136a2.14 2.14 0 0 0-.241.57l-.003.015c.184-.054.343-.136.484-.243l-.004.003a.685.685 0 0 0 .182-.182l.002-.002a.16.16 0 0 0-.045-.233h-.001a.31.31 0 0 0-.377.075zm.328-1.722c-.36.16-.734.284-1.102.421c-.029.01-.064.01-.08.046c.17.06.346.089.511.16c.434.182.795.364 1.141.568l-.043-.024a.098.098 0 0 0 .066.025a.102.102 0 0 0 .08-.039c.206-.194.425-.374.646-.552q.189-.15.374-.305l-.845-.349a.08.08 0 0 0-.073.009a.607.607 0 0 1-.397.046l.004.001a1.372 1.372 0 0 0-.174-.025h-.017a.241.241 0 0 0-.093.019zm-1.645.739l.862.214a.034.034 0 0 0 .012.012h.01l.004.001l.004-.001v-.01c0-.007 0-.01-.004-.012a.048.048 0 0 0-.028-.01h-.001c-.19-.089-.372-.197-.57-.274a.232.232 0 0 0-.089-.018a.329.329 0 0 0-.201.098zm-.511-2.08a.526.526 0 0 0 .001.489l-.001-.003a.179.179 0 0 1-.043.223c-.094.103-.185.206-.278.31c-.022.024-.047.043-.038.08c.026.097 0 .212.106.278c.031.022.018.058.018.089v.454c0 .033-.007.069.022.089a.624.624 0 0 0 .264.086h.003l.007.001l.007-.001h.01c.05.007.055-.072.086-.11c.014-.018.022-.043.036-.061c.192-.103.349-.263.55-.348a5.105 5.105 0 0 1 1.137-.467l.037-.009a2.531 2.531 0 0 0-.384-.029h-.007a2.4 2.4 0 0 0-.4.033l.014-.002c-.104.018-.226.03-.35.034h-.015l-.05.001c-.075 0-.148-.006-.22-.019l.008.001a.071.071 0 0 1-.068-.071v-.001c-.004-.033.01-.05.043-.058l.014-.004c.08-.016.153-.033.229-.05s.17-.04.255-.058c.145-.035.314-.059.487-.067h.006l.007.001l.007-.001h.007c.014 0 .029 0 .031-.012c.005-.026-.018-.025-.035-.024h-.022a1.203 1.203 0 0 1-.417-.194l.004.002a1.54 1.54 0 0 1-.658-.647l-.004-.008a2.213 2.213 0 0 0-.251-.338l.002.002l-.014-.017c-.051.154-.087.294-.145.423zm3.142 1.004c-.01.05-.014.098.058.115c.072.018.132.038.19.061l-.01-.004c.471.12.884.289 1.266.507l-.025-.013a.13.13 0 0 0 .082.029h.008a.226.226 0 0 0 .12-.037l-.001.001a.326.326 0 0 0 .182-.452l.001.002a.499.499 0 0 1 .081-.49l-.001.001a.137.137 0 0 0 .007-.15v.001c-.102-.247-.199-.497-.3-.744c-.01-.025-.012-.054-.05-.075l-.007.007l-.012-.007c-.017.026-.031.054-.048.08l-.4.48a1.148 1.148 0 0 1-.074.067a4.573 4.573 0 0 1-.685.394l-.029.012c.05 0 .099.005.149.01c.357.02.689.104.993.24l-.018-.007c.125.058.134.16.144.271c.004.046-.005.08-.039.08a.166.166 0 0 1-.045-.012h.001a10.63 10.63 0 0 0-1.352-.382l-.077-.014c-.017-.005-.043-.019-.062-.014h-.009c-.02.001-.034.01-.041.04zm-4.624-.014c.056.007.107.02.156.038l-.005-.002c.041.017.08.029.12.043c.065.026.133.041.205.065v.01h.01l.026.012c.005.001.009.002.012.005s.007 0 .011.004c.061.019.104.036.137.048a.16.16 0 0 0 .05.012h.008a.053.053 0 0 0 .025-.012a.027.027 0 0 0 .009-.012a.05.05 0 0 0 .009-.016v-.001c0-.007.007-.014.01-.024c.011-.036.024-.086.047-.16v-.014h-.007v-.024a.189.189 0 0 1 .079-.086h.001l.038-.026c.007-.005.014-.009.019-.015c.011-.007.014-.022.038-.024a.261.261 0 0 0 .168-.289v.002a1.386 1.386 0 0 0-.031-.186l.002.01a3.483 3.483 0 0 0-.07-.285a.45.45 0 0 1 .013-.171l-.001.003a.612.612 0 0 1 .09-.187l-.001.002c0-.004.007-.011.01-.017a.271.271 0 0 0 .018-.034l.001-.002a.285.285 0 0 0 .024-.053v-.002c.005 0 .012-.007.018-.014c.137-.213.122-.299-.068-.461c0 0-.004 0-.007-.005s-.026-.024-.038-.035a.391.391 0 0 0-.13-.132l-.002-.001c-.012-.024-.026-.026-.039-.038a1.214 1.214 0 0 0-.407-.267l-.008-.003a.51.51 0 0 0-.13-.066h-.01a.024.024 0 0 0-.014.005a.076.076 0 0 0-.029.045v.012l-.151.206c-.048.075-.098.14-.153.201l.001-.002a.132.132 0 0 0-.058.096v.001v-.004c-.12.151-.25.284-.379.418c-.058.007-.046-.025-.038-.05a.084.084 0 0 0 .005-.017c.014-.137.058-.269.06-.406a1.58 1.58 0 0 0 .04-.355l-.001-.063v-.009a.038.038 0 0 0 0-.038a1.702 1.702 0 0 0-.001-.429l.001.008l.002-.016l-.002-.016c0-.012-.007-.029-.011-.029v-.017c0-.159-.034-.31-.094-.446l.003.007a.148.148 0 0 0-.079-.036h-.001c-.133-.07-.257-.138-.4-.206h.005c-.018 0-.036-.018-.055-.025a.792.792 0 0 1-.126-.09l.001.001a1.027 1.027 0 0 1-.107-.1l-.001-.001a1.245 1.245 0 0 1-.05-.055a.118.118 0 0 1-.025-.025a1.538 1.538 0 0 0-.123-.14h-.003c.025.191-.007.4.19.521c.017.007.038.019.026.041c0 .007-.011.011-.018.017s-.016.009-.026.012H4.75c-.05.014-.091-.024-.14-.029v-.012c-.019 0-.039-.024-.058-.024a4.446 4.446 0 0 1-.357-.311l.001.001a2.338 2.338 0 0 1-.115-.168a1.203 1.203 0 0 1-.157-.287l-.003-.008a.117.117 0 0 0-.004-.027v-.016c-.017-.147-.048-.293-.072-.437v-.006a.158.158 0 0 0-.005-.038v.001c-.019-.144-.043-.288-.062-.432h-.004c0-.005.004-.017.004-.024v-.044c0-.146.009-.29.027-.431l-.002.017v-.034l-.001-.01l.001-.01c0-.006 0-.011.005-.018s0-.007 0-.01a.082.082 0 0 0 .007-.028v-.01l.001-.026l-.001-.028v.001a.23.23 0 0 0-.08.064v.001c-.022.022-.048.029-.065.08c-.024.042-.052.08-.08.122c-.074.11-.146.24-.209.334c-.007.024-.022.031-.029.043c-.074.11-.139.23-.211.35a.353.353 0 0 1-.052.131l.001-.001c-.033.062-.058.125-.089.185l-.094.184l-.022-.007l.038-.742v-.01a.14.14 0 0 0 .009-.031v-.001a.115.115 0 0 0 .009-.028v-.008a.163.163 0 0 1 .019-.055v.001c.052-.112.096-.243.128-.38l.003-.014a.754.754 0 0 0 .074-.222l.001-.004c.038-.076.069-.165.09-.259l.001-.008c0-.017.017-.036.024-.054l.055-.101c.05-.103.097-.207.137-.314c.004-.008.007-.022.011-.032s.005-.01.007-.01c.037-.062.072-.135.102-.21l.004-.01h.005c.031-.074.069-.151.101-.226c.05-.115.107-.213.174-.305l-.003.005c.042-.06.085-.112.13-.161l-.001.001v-.046c-.025-.137-.05-.266-.075-.41v-.01l-.001-.018l.001-.019v.001c-.022-.111-.041-.221-.062-.331c-.007-.036-.011-.072-.019-.111c0-.014-.006-.029-.006-.046c-.048-.228-.08-.457-.118-.684a.072.072 0 0 0-.072-.07h-.003a.079.079 0 0 1-.022-.005h.001h-.018a1.796 1.796 0 0 1-.241-.033l.011.002c-.074-.012-.137-.025-.233-.036v-.01h.014c-.145-.024-.277-.036-.444-.055h-.012a.162.162 0 0 0-.074-.018l-.021.001h.001c-.144-.024-.288-.041-.432-.06a2.099 2.099 0 0 0-.397-.057h-.005a.064.064 0 0 0-.033-.01c-.126-.031-.259-.033-.386-.062h-.006a.443.443 0 0 1-.09-.01l.003.001c-.08-.026-.096.01-.108.08a6.044 6.044 0 0 1-.089.425a.26.26 0 0 0 .048.24c.036.055.069.108.098.16a.544.544 0 0 1 .025.054l.001.003a.64.64 0 0 1 .048.244v.013v-.001c-.018.598.031 1.193.053 1.79c.01.286.029.572.043.858v.08c.011.154.012.306.034.458v.062a3.49 3.49 0 0 0-.073.405l-.002.017a.15.15 0 0 1-.005.037v-.001a.937.937 0 0 0-.052.441v-.009c0 .01.005.017.007.026l.007.031a.087.087 0 0 0 .04.062c.103.048.206.091.31.137a.055.055 0 0 0 .029.009a.055.055 0 0 0 .043-.021h.007c0-.012.022-.026.034-.039c.068-.049.127-.099.183-.154s.004-.007.007-.01a.094.094 0 0 1 .014-.019c.007-.004.011-.01.018-.014a.279.279 0 0 1 .054-.042l.001-.001a.124.124 0 0 1 .057-.022h.008c.007.002.012.006.017.01l.005.004c0 .005 0 .007.005.014l.001.016l-.001.017v-.001v.036a1.92 1.92 0 0 0-.165.423l-.003.014v.018a.122.122 0 0 0-.024.053v.001l-.001.013l.001.013v-.001c.001.005.002.01.005.015l.007.014c.073.114.144.212.22.306l-.004-.006v-.019a.21.21 0 0 0 .012.04v-.001a.143.143 0 0 0 .027.043l.021.024l.012.012c.116.107.246.202.385.282l.011.006c.025.014.041.029.061.041c.047.043.098.065.155.112l.008-.001c.02 0 .038.01.049.025c.061.051.128.099.2.14l.007.004c.025.011.025.024.039.024c.126.097.25.183.378.271v.005a.35.35 0 0 0 .108.079l.002.001c.13.118.293.213.414.32c.024.014.04.024.06.038s.029.01.043.034h-.006c.024.014.05.026.074.041c.09.087.192.162.302.224l.007.004c.024.01.026.014.039.039c.086.056.159.109.228.166l-.005-.004v-.005c.024.017.036.029.055.041c.16.089.151.23.112.379c0 .007 0 .017-.004.026s-.005.036-.007.055s0 .019 0 .029c-.01.089-.031.182.043.252c.007.005 0 .01.022.014c0 .014.025.026.036.038c.004.01.009.019.015.027c.019.031.041.06.055.084v.04a1.017 1.017 0 0 1 0 .119v-.003a.117.117 0 0 0 .113.148c.091.018.168.041.264.061l.022.001l.023-.001h-.001c.134.046.275.07.414.118zm3.768-.08c-.019 0-.038 0-.041.026l-.001.012c0 .038.028.07.065.075c.082.02.177.031.275.031h.004c.096 0 .191-.008.283-.023l-.01.001c.07-.01.061-.047.061-.09s-.026-.055-.058-.054l-.02-.002l-.021.002h.001a1.187 1.187 0 0 1-.239.024h-.261l-.014-.001l-.014.001h.001zm-4.525-3.754c.02.035.039.064.061.092l-.001-.002l.025.016a.483.483 0 0 0 .129.083l.003.001c.071.03.154.048.241.048h.012h-.001a.21.21 0 0 1 .219.123l.001.001a.124.124 0 0 1 .01.02v.001c.128.311.342.565.613.738l.006.004l.06.038a2.275 2.275 0 0 1 .301.196l.08.055c.106.08.2.161.288.247l.068.068q.133.134.259.276c.022.022.04.046.062.07c.043.046.084.094.126.142c.098.109.189.229.269.356l.007.011c.01.017.022.033.032.053c.029.053.06.117.086.183l.004.012c.051.116.131.21.231.277l.002.001c.05.038.106.072.16.106c.042.026.086.05.13.072c.062.031.12.067.183.101a.646.646 0 0 0 .139.059l.005.001l.029.007a2.146 2.146 0 0 0 .614.021l-.01.001c.06-.007.118-.017.175-.024s.108-.018.16-.026h.019c.055-.01.11-.019.16-.032q.206-.042.407-.103q.066-.018.13-.041a.245.245 0 0 0 .041-.012h-.001c.04-.014.08-.029.122-.046c.014-.004.029-.011.043-.017l.053-.022l.065-.029c.058-.026.113-.053.17-.084c.07-.039.14-.08.209-.122c.034-.023.069-.047.103-.072c.094-.067.177-.14.252-.22l.001-.001c.036-.038.07-.08.103-.119c.148-.191.283-.406.396-.634l.01-.022c.022-.04.043-.08.062-.125q.046-.095.089-.192l.022-.053c.028-.09.056-.205.077-.322l.003-.019a1.162 1.162 0 0 0-.103.041a2.55 2.55 0 0 1-.147.061a.581.581 0 0 1-.145.033h-.003l-.02.001l-.021-.001h.001l-.011.001a.11.11 0 0 1-.106-.08v-.001c-.019-.055.01-.08.046-.091l.031-.01l.019-.007l.034-.014l.036-.014c.178-.058.333-.134.475-.228l-.007.005q.13-.096.25-.204c.031-.029.062-.058.09-.089c.132-.135.24-.294.316-.47l.004-.01a.107.107 0 0 1 .007-.019v.001c.019-.038.041-.072.058-.11a.086.086 0 0 1 .081-.058h.006c.038 0 .08-.012.115-.012h.062a.627.627 0 0 0 .474-.307l.002-.003c.074-.108.139-.231.187-.363l.004-.012a.651.651 0 0 0 .024-.052l.001-.004a.774.774 0 0 0 .072-.236v-.011c0-.026.011-.062-.018-.072a.434.434 0 0 0-.138-.022h-.001c-.02 0-.039.002-.058.004h.002l-.027.005h.001c-.006.001-.01.004-.014.007s-.005 0-.005.005c-.014.018.019.06.01.089s0 .058-.01.084s-.019.058-.031.084a1.213 1.213 0 0 1-.058.124l.003-.006c-.082.144-.18.268-.292.376l-.001.001a.307.307 0 0 1 .001-.264l-.001.002l.011-.026c.025-.053.055-.103.084-.154c.01-.017.024-.031.034-.048c0 0 0-.004.004-.01l.001-.012l-.001-.012a.464.464 0 0 0-.148-.22l-.001-.001a.382.382 0 0 0-.065-.049l-.002-.001a.223.223 0 0 0-.059-.025h-.001c-.029 0-.06-.007-.09-.01a1.618 1.618 0 0 1-.261-.041l.011.002l-.022-.007q-.194-.061-.379-.14l-.122-.054a3.247 3.247 0 0 1-.479-.273l.011.007c-.108-.08-.219-.152-.329-.228a7.888 7.888 0 0 1-.219-.154l-.32-.24l-.022-.018l-.046-.033a.422.422 0 0 1-.046-.034a.236.236 0 0 1-.059-.059l-.001-.001l-.029-.039c-.01-.011-.019-.026-.029-.038a1.661 1.661 0 0 1-.37-.562l-.004-.011a2.44 2.44 0 0 1-.283.486l.005-.006a2.047 2.047 0 0 1-.14.161l.001-.001a1.653 1.653 0 0 1-.719.434l-.012.003c-.173.051-.373.08-.579.08h-.02h.001c.06-.048.112-.094.166-.137l.049-.046l.001-.001c.184-.139.346-.286.495-.446l.002-.002l.022-.014a.069.069 0 0 0 .019-.033v-.012h-.01a.092.092 0 0 0-.054.033c-.14.08-.284.16-.428.234l-.062.034l-.11.054a3.384 3.384 0 0 1-1.11.323l-.016.002c-.18.019-.36.039-.54.055s-.36.031-.54.043c-.12.007-.24.014-.36.019l-.027.001l-.028-.001h.001h-.105c-.036 0-.07 0-.106.004s-.059.026-.094.026c-.019.005-.036.007-.053.012a4.058 4.058 0 0 0-.881.325l.023-.011c-.074.038-.147.08-.218.122c-.046.031-.094.06-.139.094a1.336 1.336 0 0 0-.068.05a.448.448 0 0 0-.125.139l-.001.002a.118.118 0 0 0-.012.028v.001l-.005.015v-.001l-.007.032l-.001.025l.001.026v-.001a.369.369 0 0 0 .018.107l-.001-.003l.035.085c.026.065.05.13.073.197v-.006l.005.014v-.008c.125.257.251.474.391.68l-.012-.018zm7.397 2.742c0 .007 0 .014.01.019s0 .011 0 .017a.35.35 0 0 0 .037.1l-.001-.002c0 .014.012.01.017.034c.025.072.046.151.072.223s.034.16.106.216c.024.112.048.118.16.074a1.75 1.75 0 0 0 .187-.085l-.009.005c.12-.054.252-.125.379-.173h.017c.139-.062.278-.119.418-.182c.09-.038.176-.074.264-.113c.068-.038.15-.075.234-.106l.013-.004c.005 0 .019-.005.019-.007c.144-.068.302-.137.454-.205l.036-.022a.615.615 0 0 0 .214-.092l-.002.001l.094-.039l.132-.054c.005-.001.009-.002.012-.005v.005a.774.774 0 0 0 .103-.035l-.005.002l.06-.024c.036-.018.074-.034.113-.048l.058-.022a.08.08 0 0 0 .059-.077l-.001-.015v-.055c0-.253.015-.503.043-.748l-.003.03l.001-.036l-.001-.038v.002v-.035l.001-.018l-.001-.019v.001l-.001-.018l.001-.019v.001c0-.011.005-.022.007-.033l.005-.017v-.001a.152.152 0 0 1 .012-.032v.001l.007-.014c.064-.117.133-.263.194-.413l.011-.031c.082-.137.154-.296.209-.463l.005-.017c.058-.23.068-.473.106-.708c.041-.25.062-.504.118-.754c.068-.302.134-.608.202-.912a.06.06 0 0 0 .005-.019c.012-.038.022-.08.034-.12a.445.445 0 0 0 .036-.178v-.014v.001c0-.005.007-.01.007-.014v-.005l.007-.014c.014-.055.026-.113.041-.17c.062-.16.127-.331.19-.475a.086.086 0 0 1 .017-.036a1.82 1.82 0 0 0 .184-.418l.004-.014a.265.265 0 0 0 .071-.129v-.002s0-.004.004-.007s.005-.007.005-.01v-.005v.005c0-.007.012-.017.012-.024c.04-.077.083-.176.122-.276l.007-.022a.124.124 0 0 1 .019-.04c.054-.145.173-.27.17-.437a.127.127 0 0 0 .02-.068l-.001-.013v.001l.072-.31a.973.973 0 0 0 .019-.099l.001-.005a.148.148 0 0 0-.058-.15a4.452 4.452 0 0 0-.65-.337l-.03-.011a.285.285 0 0 1-.168-.093a.04.04 0 0 0 .004-.019a.252.252 0 0 1 .294-.049l-.001-.001c.031.012.067.02.104.02h.015h-.001c.007 0 .011.007.018.007s.012.007.019.007l.252.08h.01l.028.005h.007c.002 0 .012-.01.014-.022c.018-.127.097-.254 0-.382a.073.073 0 0 0-.015-.036a.271.271 0 0 1-.01-.415v-.011c0-.004 0-.006.005-.01a.078.078 0 0 1 .01-.015a.032.032 0 0 0 .014-.014c.002-.003.003-.007.003-.011s-.001-.008-.003-.011a.068.068 0 0 0-.014-.025c-.104-.11-.212-.22-.31-.334c-.025-.024-.032-.046-.054-.065a.319.319 0 0 0-.042-.03l-.002-.001c-.108-.149-.24-.199-.4-.08c-.108.024-.204.08-.32.118c-.024.019-.096.036-.12.055a.347.347 0 0 0-.062.019l.002-.001c-.175.053-.321.11-.462.177l.02-.009l-.008-.001l-.009.001c-.151.055-.302.11-.454.169l-.055.021l-.019.008l-.074.029c-.118.035-.408.146-.432.155c-.137.05-.142.054-.101.19a.261.261 0 0 1-.051.253c-.072.09-.133.19-.198.283l-.01.007l-.019.021c-.004.01-.01.018-.014.026a.344.344 0 0 0-.011.041v.007l.007.018c.003.005 0 .004 0 .007s0 .011.007.017s.012.026.017.038a1.2 1.2 0 0 0 .211.358l-.001-.002s0 .004.004.007a.026.026 0 0 1 .014.011c.022.043.043.084.06.127c.014.03.023.066.026.103v.001l.001.024c0 .06-.02.116-.054.161v-.001l-.005.005l-.01.01s-.005 0-.005.005a3.394 3.394 0 0 0-.141.442l-.006.026v.018c-.012.036-.025.072-.034.108a.482.482 0 0 0-.019.077v.029l-.001.027l.001.029v-.001a.363.363 0 0 0 .058.17l-.001-.001c.007.024.01.038.017.058c.044.116.087.263.121.412l.005.025c0 .007.01.01.01.01q.013.044.024.089l.014.062l.014.06a.423.423 0 0 1 .012.06v.002a.925.925 0 0 1 .012.223v-.003l-.001.018c0 .016.003.031.008.045v-.001c0 .004 0 .01.005.011a.08.08 0 0 0 .026.036a.104.104 0 0 0 .072.029h.017s.01.004.014.004c.147.055.293.106.44.16c.011.007.024.011.035.018c.032.022.066.041.097.062s.065.055.096.055c.012 0 .025.014.036.022c.062.063.13.119.204.167l.005.003a.527.527 0 0 0 .06.064c.062.07.118.149.19.219c0 .011.012.024.019.036c.038.05.072.101.108.151c.014.019.041.041.024.065s-.046 0-.065-.012a1.383 1.383 0 0 0-.269-.106l-.01-.002v-.004c-.007 0-.011-.007-.017-.007a1.689 1.689 0 0 0-.428-.112l-.009-.001v.012a7.106 7.106 0 0 0-.759-.039l-.172.002h.009c.01.039.034.025.054.048c.105.072.194.173.314.255c.024.007.018.011.024.017l.019.017c.05.056.105.104.166.145l.003.002c.01.024.019.024.029.033l.01.01c.075.083.148.174.214.269l.006.009s.005.005.005.007v-.017l.038.062a.27.27 0 0 0 .055.08a.106.106 0 0 0 .036.061c.065.119.123.257.166.402l.004.016l-.001.02l.001.021V15.9c.053.133.091.286.107.447l.001.007c0 .012.007.012.007.036H15.8c-.024.084.039.168.012.257a.07.07 0 0 1-.041-.025c-.007-.007-.012-.015-.017-.023v-.001l-.019-.041l-.033-.08a.359.359 0 0 0-.044-.085l.001.001c0-.024-.007-.019-.01-.029l-.01-.031a1.066 1.066 0 0 0-.171-.322l.002.002a.301.301 0 0 0-.11-.151l-.001-.001a1.248 1.248 0 0 0-.369-.355l-.005-.003h-.005l-.017-.012a.504.504 0 0 0-.078-.061l-.002-.001c-.024-.017-.05-.034-.074-.048s-.034-.022-.051-.031a.08.08 0 0 0-.04-.017c-.019 0-.029.017-.034.05c-.01.08-.022.16-.034.24l-.036.24v.017a1.524 1.524 0 0 0-.055.476v-.003v.019c0 .034 0 .07-.005.103a1.161 1.161 0 0 0-.004.342l-.001-.006a.12.12 0 0 0 .007.034v-.001v.017l-.001.088c0 .115.007.229.021.34l-.001-.013v.026l-.001.013l.001.013v-.001l.04.329a.695.695 0 0 1 .01.107v.001c0 .005.011.012.011.012c.005.048.012.091.018.134s.01.093.017.139c-.062-.062-.115-.108-.166-.16c-.014-.017-.026-.036-.041-.053a.522.522 0 0 0-.054-.06c0-.012-.014-.025-.022-.036a.318.318 0 0 1-.026-.035l-.001-.001a.267.267 0 0 1-.028-.056l-.001-.002l-.022-.062c-.031-.101-.07-.199-.106-.298l.001-.008l-.001-.009a2.928 2.928 0 0 1-.074-.78v.005h-.025c-.029.155-.058.32-.086.466h-.014a3.504 3.504 0 0 0-.051.609c0 .117.006.233.016.347l-.001-.015c0 .029.012.058.014.086l.022.173c.005.029.01.058.012.086a.115.115 0 0 1 .009.028v.005c0 .008.003.016.008.022c.005.024.007.05.012.074s.007.05.012.074c.004.013.007.03.009.046v.002c0 .007 0 .017.005.024c.003.02.007.037.012.053l-.001-.003c0 .014.01.034.014.05c.009.031.02.057.033.082l-.001-.002c.007 0 .011 0 .018.025c-.036.047-.047-.005-.072-.005v-.019h-.005a.85.85 0 0 1-.073-.086l-.002-.002c-.08-.132-.16-.271-.245-.392h-.005a3.198 3.198 0 0 0-.181-.473l.008.019c-.024-.08-.04-.16-.06-.24a.562.562 0 0 0-.037-.194l.001.004V17.3c-.007-.084-.022-.168-.033-.252a.837.837 0 0 1 .001-.333l-.001.005c.004-.029.017-.074-.022-.08l-.006-.001l-.006.001c-.026 0-.041.034-.058.055c-.05.068-.091.137-.137.205c-.01.017-.022.033-.034.05c-.024.033-.048.067-.074.098a.238.238 0 0 0-.043.047l-.001.001c-.012 0-.025.014-.048.022c-.048.05-.13.08-.146.168v.017a4.59 4.59 0 0 0-.149.437c-.004.019-.018.039-.018.058c-.048.084-.072.169-.111.24c-.017.047-.031.08-.047.115s-.032.08-.048.115s-.055.091-.075.139c0 .005 0 .012-.004.018s-.005.022-.007.033a.178.178 0 0 0 0 .071v-.001c.024.078.051.145.083.209l-.003-.007c.024.101.06.199.094.3h-.005l.013.016zm-6.72-1.618c-.008.22-.028.427-.059.631l.004-.031v.012l.024.007q.112-.223.223-.449a.046.046 0 0 0 .006-.023a.05.05 0 0 0-.018-.038l-.014-.014l-.112-.11a.05.05 0 0 0-.03-.017c-.01.001-.018.01-.02.032zm7.836-.141a.435.435 0 0 0 .048-.164v-.002c.029-.168.054-.334.086-.502c.011-.058.062-.134.019-.168c0 0-.005 0-.008-.005l-.011-.014c-.036-.029-.106.034-.16.053c-.033.012-.069.025-.067.08c.01.258.037.501.079.738l-.005-.032l.012-.023v-.001zM9.52 11.826c.011 0 .022.019.026.053a.73.73 0 0 1-.12.562l.002-.002a1.731 1.731 0 0 1-.645.481l-.011.004c-.406.214-.88.402-1.375.539l-.051.012c-.358.098-.728.151-1.088.24a5.75 5.75 0 0 0-1.36.389l.038-.014a1.128 1.128 0 0 0-.53.533l-.003.007a.517.517 0 0 0-.035.361l-.001-.004a.18.18 0 0 0 .01.037v-.001a2.33 2.33 0 0 1 .166-.312l-.006.01c.163-.214.401-.363.674-.41l.006-.001c.23-.041.464-.061.694-.09c.634-.08 1.274-.094 1.906-.187a3.4 3.4 0 0 0 1.455-.488l-.015.008c.342-.197.638-.411.908-.653l-.005.005a1.89 1.89 0 0 0 .591-.973l.003-.013a.55.55 0 0 0-.034-.416l.001.003a2.129 2.129 0 0 0-.32-.411l-.013-.013a2.145 2.145 0 0 0-1.499-.608h-.033h.002c-.218.007-.439.014-.66.017a2.698 2.698 0 0 0-.589.073l.018-.004a5.062 5.062 0 0 0-2.064 1.027l.008-.007c-.369.311-.702.626-1.017.959l-.005.006a.288.288 0 0 0-.04.052l-.001.001q-.32.511-.646 1.022a.58.58 0 0 0-.079.195l-.001.004a.146.146 0 0 0-.024.017v.006l.026-.019c0 .012-.005.025-.008.039c.055-.041.097-.068.13-.098a2.53 2.53 0 0 1 .942-.485l.018-.004a4.688 4.688 0 0 1 1.225-.239l.014-.001a10.554 10.554 0 0 0 1.828-.329l-.074.017a3.338 3.338 0 0 0 1.054-.429l-.014.008c.214-.145.401-.292.576-.452l-.004.003a.078.078 0 0 1 .047-.027zm4.04 2.88a3.102 3.102 0 0 0 .83-.034l-.019.003a.23.23 0 0 1-.031-.018l.006-.001l.006.001a1.05 1.05 0 0 0-.659-.178h.003a2.178 2.178 0 0 1-.712.043l.009.001c-.029-.004-.054-.012-.08-.014l-.075-.004c.012.006.025.011.036.018h-.017a.903.903 0 0 0 .448.16h.022c.084.003.164.011.241.023l-.012-.001zm-2.568-2.742c-.022.07-.043.14-.062.209a1.058 1.058 0 0 0 .019.715l-.002-.007v.01a.994.994 0 0 0 .202.28l.058.055c.245.237.529.434.843.58l.019.008c.249.119.545.228.851.311l.038.009a.99.99 0 0 0 .34.026h-.004c.288-.042.62-.067.959-.068h.062c.291 0 .577.017.859.05l-.034-.003a.155.155 0 0 0-.024-.019h-.001h.006a.713.713 0 0 0-.371-.16h-.003a14.3 14.3 0 0 1-1.669-.272l.094.018a1.507 1.507 0 0 1-.622-.29l.003.002q-.416-.35-.818-.714c-.13-.118-.264-.227-.39-.348a1.53 1.53 0 0 1-.295-.392l-.004-.008c-.007-.014-.014-.022-.022-.024h-.004c-.013-.002-.025.01-.03.03zm-3.574 1.18l-.523.086c-.4.046-.794.096-1.19.139a2.302 2.302 0 0 0-.758.174l.015-.006c-.18.084-.36.17-.54.253a.24.24 0 0 0-.166.275v-.001l-.001.009l.001.01c.005.022.018.026.043.01a.887.887 0 0 0 .064-.044l-.001.001a2.58 2.58 0 0 1 1.089-.485l.016-.003c.566-.106 1.138-.187 1.697-.327c.128-.021.244-.059.35-.111l-.007.003zm3.8-2.149v.266l-.001.009l.001.009l-.001.007l.001.007v.125c0 .004 0 .012-.005.017l-.005.015v-.001c0 .007 0 .012-.004.018l-.005.014v.001l-.007.017v.001a.172.172 0 0 0 .001.105v-.001v.01l-.001.005l.001.005v.004c.01.022.017.041.024.061l.036.094q.03.08.062.154c.02.041.047.076.08.105c.301.291.628.559.976.8l.026.017c.191.161.409.3.645.406l.017.007c.204.075.4.16.614.24l.166.054h.006c.18.059.388.098.603.11h.006c.031 0 .086.03.09-.04c0-.007 0-.014-.004-.022a.122.122 0 0 0-.018-.097l-.01-.014c0-.014-.007-.029-.01-.043a1.991 1.991 0 0 0-.766-1.165l-.006-.004l-.269-.206a3.116 3.116 0 0 0-.423-.283l-.016-.008c-.019-.01-.041-.022-.062-.031h-.001v.001c-.28-.127-.659-.199-.886-.313c-.005-.003-.011-.007-.018-.009h-.001a.133.133 0 0 1-.057-.038a.174.174 0 0 0-.012-.014c-.011-.012-.022-.022-.017-.024l.151.034c.122.02.23.048.334.084l-.014-.004l.099.026c.413.109.773.257 1.109.444l-.024-.012c.098.055.199.108.298.16a.074.074 0 0 0 .041.013l.01-.001a.037.037 0 0 0 .018-.014c.017-.021 0-.04-.014-.058l-.223-.314a.382.382 0 0 0-.116-.107l-.002-.001c-.221-.134-.442-.271-.663-.4a2.5 2.5 0 0 0-.26-.137l-.018-.007c-.154-.1-.343-.16-.546-.16h-.005a2.145 2.145 0 0 0-.384-.024h.003c-.041 0-.08 0-.122.005h-.075c-.08 0-.16.006-.24.004h-.008c-.078-.01-.105.022-.102.101zM7.974 8.702c-.137.043-.277.08-.415.118l-.118.036q-.087.03-.173.065a.468.468 0 0 0-.03.012h.001c-.038.015-.074.031-.111.047s-.058.033-.086.051c-.302.18-.564.372-.805.588l.005-.004l-.067.062l-.012.01c-.091.073-.16.142-.24.21l-.026.03a.926.926 0 0 0-.115.098c-.08.097-.166.19-.25.286s-.166.192-.247.288s-.16.194-.24.295l-.08.101c-.05.068-.098.137-.146.206a4.316 4.316 0 0 0-.25.41l-.012.025a1.375 1.375 0 0 0-.146.466l-.001.007c.098-.082.186-.17.266-.264l.003-.003c.16-.134.313-.283.464-.429l.065-.066l.01-.008c.072-.08.151-.152.236-.215l.004-.003c.296-.216.633-.408.992-.559l.033-.012l.018-.004c.062-.022.122-.049.181-.075c.111-.058.24-.106.375-.138l.012-.002c.209-.039.421-.072.631-.113a3.72 3.72 0 0 1 .803-.091h.001c.359.012.699.068 1.024.162l-.029-.007c.42.124.774.363 1.037.68l.003.004l.216.25c.011.014.018.031.033.036s.022-.005.03-.034c.061-.17.096-.366.096-.571v-.023v.001a2.718 2.718 0 0 0-.135-.636l.006.019a1.667 1.667 0 0 0-.463-.719l-.001-.001a1.99 1.99 0 0 0-.2-.162l-.006-.004c-.033-.026-.072-.047-.108-.072l-.112-.065c-.055-.034-.115-.062-.173-.091l-.058-.029c-.08-.038-.16-.072-.24-.108a1.641 1.641 0 0 0-.562-.165l-.007-.001h-.141c-.263 0-.516.041-.754.116l.018-.005zm4.073 2.736c.004-.015.021-.012.033-.018a.123.123 0 0 1-.029-.009h.001c-.006.002-.01.009-.005.03zm.443-.756c.125.023.236.058.34.104l-.009-.004c.269.123.491.242.704.372l-.029-.017a.982.982 0 0 0 .177.101l.006.003l.033.013a9.152 9.152 0 0 0-.56-.774a2.052 2.052 0 0 0-.422-.401l-.006-.004a.273.273 0 0 0-.111-.079l-.002-.001a5.25 5.25 0 0 0-.732-.251l-.039-.009l-.015-.005h.001a.936.936 0 0 0-.608.035l.006-.002c-.041.014-.07.033-.055.08a2.232 2.232 0 0 1 .06.708v-.007c0 .062.029.091.09.091c.084 0 .17.007.255 0q.16-.014.315-.014c.212.001.419.023.619.063l-.02-.003zm-11.455.32a.27.27 0 0 0 .022-.107v-.006c0-.005 0-.01-.004-.014l.001.006l-.001.006l.001.017a.21.21 0 0 1-.022.094l.001-.001zm-.454-2.5c0 .005-.002.01-.005.014c0 .01 0 .019-.004.032c-.05.48-.101.953-.154 1.429c-.007.06.007.086.05.098l.035.005h.001c.674.091 1.35.183 2.02.278l.676.096l1.04.142l.01.001l.01-.001c.04 0 .047-.038.058-.074l.278-.912c.022-.065 0-.084-.06-.101c-.646-.17-1.292-.338-1.938-.509L.66 8.491a.118.118 0 0 0-.039-.011H.617a.039.039 0 0 0-.035.022zm17.218-1.26l-4.254 1.246c-.072.019-.062.046-.034.097q.334.586.662 1.176c.025.043.043.061.08.055a.372.372 0 0 0 .042-.01l-.003.001q1.68-.59 3.37-1.17a.127.127 0 0 0 .094-.122v-.008l.043-.536c.019-.24.041-.48.062-.739v-.004l-.064.014zM6.32 3.495c-.342.237-.734.45-1.148.619l-.041.015a.85.85 0 0 0-.488.367l-.002.004c-.005.01-.022.014-.032.022c-.123.074-.267.13-.421.159l-.008.001l-.024.001l-.025-.001h.001h-.009a.146.146 0 0 0-.072.019h.001q-1.291.812-2.586 1.612c-.05.031-.06.055-.033.11c.16.332.302.673.492.99a.122.122 0 0 1 0 .154a1.959 1.959 0 0 0-.297.789l-.001.011c-.04.067-.033.108.055.119c.058.007.111.025.166.036c.442.118.88.23 1.33.334a.159.159 0 0 0 .113.021H3.29c.019-.007.018-.025.018-.039l.001-.009l-.001-.01l.274-.915a.23.23 0 0 1 .06-.125a.477.477 0 0 0 .19-.382c.004-.061.024-.118.029-.178v-.001a.17.17 0 0 1 .107-.158h.001c.336-.158.621-.374.852-.637l.003-.003a.76.76 0 0 0 .144-.202l.002-.004c.201-.43.409-.793.643-1.138l-.019.029c.257-.414.548-.773.878-1.09l.002-.002a.56.56 0 0 0 .159-.288l.001-.004a.22.22 0 0 0-.204-.265h-.001a.175.175 0 0 0-.109.039m7.937 3.229q.172.485.353.967c.022.062.043.08.113.061c.48-.147.965-.293 1.447-.435c.075-.022.086-.05.061-.122a22.155 22.155 0 0 0-.67-1.635a.24.24 0 0 1-.011-.211l-.001.002a1.44 1.44 0 0 0 .121-.808l.001.007c-.017-.139.01-.278-.007-.422a9.513 9.513 0 0 0-.27-1.183l.018.069a7.329 7.329 0 0 0-.5-1.218l.02.041a8.193 8.193 0 0 0-.456-.793l.02.032c-.084-.146-.183-.285-.262-.436a.671.671 0 0 0-.309-.328L13.921.31c-.041-.018-.091-.054-.122-.014a.184.184 0 0 0-.048.184V.479c.053.127.11.254.16.382c.151.3.293.655.405 1.024l.013.049c.035.125.064.275.083.429l.001.015c.004.055 0 .115-.058.137s-.065-.043-.094-.074c-.108-.118-.25-.235-.418-.166a.268.268 0 0 1-.162.014h.002c-.036-.007-.08-.017-.09-.058s.029-.065.06-.086a.37.37 0 0 0 .087-.056h-.001c.036-.033.062-.08.024-.125a1.995 1.995 0 0 0-.282-.328l-.001-.001a9.55 9.55 0 0 0-.716-.522a18.455 18.455 0 0 0-1.465-.804a.36.36 0 0 0-.186-.051h-.014h.001c-.046.005-.098 0-.12.055a.143.143 0 0 0 .025.151a.706.706 0 0 0 .17.158l.003.002c.233.118.428.234.614.361l-.019-.012c.36.274.72.554 1.066.845c.11.094.211.195.312.298l.181.187a.19.19 0 0 1 .043.071v.001l.001.005l-.001.005l.001.011l-.001.011a.095.095 0 0 1-.031.055c-.058.058-.116.029-.17-.005a.366.366 0 0 1-.059-.059l-.001-.001l-.029-.029c-.047-.048-.09-.096-.141-.142c-.019-.017-.041-.033-.062-.05c-.269-.202-.538-.4-.806-.602s-.538-.4-.804-.605a2.721 2.721 0 0 0-.463-.285l-.017-.007a.095.095 0 0 0-.137.024a.159.159 0 0 0-.024.19V.89c.055.11.13.203.219.277l.001.001c.126.104.24.209.347.319l.001.001c.504.56 1.018 1.11 1.506 1.68a.261.261 0 0 1 .062.132v.001c.014.097-.034.137-.126.106a.435.435 0 0 1-.053-.028l.002.001c-.7-.38-1.279-.74-1.838-1.128l.064.042a2.87 2.87 0 0 0-.509-.284l-.019-.008c-.058-.026-.12-.058-.168.014a.16.16 0 0 0 .043.197c.216.16.414.348.619.526c.487.375.908.799 1.268 1.273l.012.017a.282.282 0 0 0 .312.054l-.002.001a1.15 1.15 0 0 1 .648-.199h.02h-.001h.062c.014 0 .028.005.038.012a.06.06 0 0 1 .014.021c.026.058-.036.072-.069.094c-.216.142-.435.278-.651.422a.182.182 0 0 1-.167.028h.001a1.137 1.137 0 0 0-.309-.065h-.003c-.485.013-.949.06-1.402.138l.057-.008a.798.798 0 0 0-.295.082l.005-.002a.135.135 0 0 0-.065.145v-.001c.005.046.017.096.084.108c.278.053.554.113.833.17c.086.026.188.047.294.059l.009.001c.32.198.631.4.934.626c.171.134.378.23.603.273l.009.001c.106.017.209.047.313.074c.148.014.284.055.406.118l-.006-.003c.094.1.167.22.214.353l.002.007c.038.08.086.118.175.097a.122.122 0 0 1 .034-.004c.056-.005.081.03.099.086zM.349 2.587a.104.104 0 0 0-.068.113v-.001c.039.205.061.414.113.612c.228.862.516 1.705.79 2.553c.024.072.048.065.106.034l2.139-1.206l.278-.16a8.994 8.994 0 0 1-.859-.599l.019.015c-.022.098-.043.178-.06.257a.062.062 0 0 1-.055.055c-.069.01-.096-.05-.134-.09s-.068-.08-.046-.126c.061-.13.094-.27.147-.4s.054-.137.17-.05c.011.008.02.015.029.024c.245.257.53.473.847.637l.018.008a.354.354 0 0 0 .443-.024c.009-.007.02-.013.031-.019h.001a.038.038 0 0 0 .024-.035a.526.526 0 0 0-.104-.213l.001.001c-.176-.146-.336-.305-.502-.467a4.165 4.165 0 0 1-.637-.724l-.009-.014a1.041 1.041 0 0 0-.18-.228a.208.208 0 0 0-.334.047l-.001.001c-.134.25-.312.47-.451.72c-.151.271-.342.518-.48.8a.182.182 0 0 0 .001.161v-.001c.069.137.134.3.186.469l.006.023c.01.039.019.072-.034.091s-.062-.01-.08-.041c-.098-.166-.202-.328-.293-.497A15.084 15.084 0 0 0 .39 2.588l.036.057c-.016-.02-.022-.066-.051-.066a.048.048 0 0 0-.023.007zm4.086.4c-.187.086-.32.249-.5.35c-.047.026-.024.054 0 .08l.54.619a.271.271 0 0 1 .031.049l.001.002a.434.434 0 0 0 .048-.02l-.002.001l.257-.16c.029-.017.026-.038.017-.065c-.101-.276-.182-.56-.302-.83c-.014-.031-.026-.043-.043-.043a.11.11 0 0 0-.047.016h.001zm.827-1.298c-.036.019-.07.043-.106.062a9.774 9.774 0 0 1-.975.506l-.065.026c-.253.101-.502.205-.757.295c-.029.012-.09.01-.046.065l.358.464c.018.024.038.041.072.024c.189-.08.349-.168.501-.269l-.011.007c.37-.295.73-.601 1.076-.922c.054-.05.094-.101.08-.16c0-.058.017-.108-.031-.134a.04.04 0 0 0-.022-.006c-.026-.001-.049.032-.074.041zM12.8 23.788a1.552 1.552 0 0 1-.935-.381l.002.002a.287.287 0 0 1-.086-.394l-.001.001c.048-.086.122-.086.204-.08h.001c.066 0 .119.051.124.115a.237.237 0 0 0 .08.125c.104.099.229.177.367.228l.007.002a.197.197 0 0 0 .142.004h-.001c.08-.026.105-.08.058-.151c-.068-.098-.144-.192-.216-.286a2.633 2.633 0 0 1-.271-.401l-.007-.014a.416.416 0 0 1 .079-.525a.682.682 0 0 1 .928-.049l-.001-.001a.195.195 0 0 1 .089.195v-.001c-.029.16-.14.216-.27.115a.37.37 0 0 0-.362-.079l.003-.001c-.146.048-.187.139-.098.264c.112.16.24.32.357.475c.074.091.138.194.187.306l.003.009a.384.384 0 0 1-.359.521l-.021-.001h.001zm.64-.029a.149.149 0 0 1-.08-.183v.001c.098-.346.192-.696.298-1.04c.05-.16.115-.31.303-.358a.047.047 0 0 0 .024-.026a.391.391 0 0 1 .363-.248h.017h-.001a4.728 4.728 0 0 1 .839.002l-.018-.001c.125.004.151.08.08.184a.34.34 0 0 1-.189.129l-.002.001a2.135 2.135 0 0 1-.57.076l-.068-.001h.003h-.08a.46.46 0 0 0-.169.012l.003-.001c-.09.032-.04.147-.08.209a.79.79 0 0 0-.061.131l-.002.006c.16-.014.307-.029.454-.04a1.373 1.373 0 0 1 .194 0h-.005c.08 0 .111.048.104.126a.262.262 0 0 1-.174.225l-.002.001a1.2 1.2 0 0 1-.415.073l-.059-.001h.003c-.216 0-.234.014-.264.24c.186-.033.4-.052.618-.053h.001c.064.002.126.008.185.018l-.008-.001c.08.012.148.038.16.13l.002.032a.24.24 0 0 1-.119.207l-.001.001a.254.254 0 0 1-.118.029h-.054c-.336 0-.661.052-.965.148l.023-.006a.286.286 0 0 1-.08.012a.175.175 0 0 1-.093-.031h.001zm-10.955-.144a.632.632 0 0 1-.073-.485l-.001.004c.091-.455.321-.846.644-1.137l.002-.002a.939.939 0 0 1 .611-.24h.001a.774.774 0 0 1 .346.117l-.003-.002a.09.09 0 0 1 .05.08a.105.105 0 0 1-.007.036v-.001a.357.357 0 0 1-.147.202l-.001.001a.136.136 0 0 1-.152.011h.001a.389.389 0 0 0-.443.041h.001a.823.823 0 0 0-.243.279l-.002.004c-.08.138-.145.298-.185.468l-.002.012a.315.315 0 0 0-.003.125v-.002c.022.108.058.132.16.105c.214-.054.415-.146.622-.22a.51.51 0 0 1 .191-.046h.001c.07 0 .103.029.103.104v.003a.29.29 0 0 1-.08.201a.843.843 0 0 1-.222.166l-.005.002a3.01 3.01 0 0 1-.849.302l-.02.003a.377.377 0 0 1-.07.007h-.001a.248.248 0 0 1-.222-.139l-.001-.001zM8.6 19.092a.653.653 0 0 1-.075-.279v-.001c.025-.15.055-.279.093-.404l-.006.022c.096-.214.302-.257.486-.32c.209-.072.45-.116.7-.122h.003a.722.722 0 0 1 .639.296l.001.002c.043.058.036.08-.036.098a.077.077 0 0 1-.062-.012a1.398 1.398 0 0 0-.626-.146a.187.187 0 0 0-.102.011h.001c-.173.067-.346.134-.516.206a.488.488 0 0 0-.28.347l-.001.003c-.022.113-.05.21-.084.304l.004-.014c-.012.029-.026.065-.07.067h-.004c-.042.001-.05-.035-.066-.059zm.624-1.303a.439.439 0 0 0-.52-.031l.002-.001a.704.704 0 0 1-.565.079l.005.001a.178.178 0 0 1-.068-.032c-.086-.07-.086-.08.007-.142a.708.708 0 0 0 .102-.071l-.001.001a.183.183 0 0 1 .191-.038h-.001a.332.332 0 0 0 .267-.068l-.001.001a.55.55 0 0 1 .625.006l-.002-.001a.104.104 0 0 1 .025.015c.139.111.274.176.432.019c.065-.065.175-.038.264-.05c.223-.006.432-.056.622-.141l-.01.004a.049.049 0 0 1 .033-.013a.05.05 0 0 1 .047.034c.033.06.09.103.098.18c0 .036 0 .065-.038.08c-.091.036-.182.098-.278.01c-.031-.026-.065-.014-.098-.01a2.985 2.985 0 0 0-.671.221l.018-.008a.546.546 0 0 1-.214.047h-.003a.432.432 0 0 1-.27-.094l.001.001zm.024-.5a1.674 1.674 0 0 1-.221-.216l-.002-.002a.44.44 0 0 0-.388-.16h.002a.56.56 0 0 1 .198-.277l.001-.001a.416.416 0 0 0 .143-.382v.002a1.922 1.922 0 0 0-.107-.573l.004.013a.512.512 0 0 0-.224-.29l-.002-.001a.43.43 0 0 0-.159-.036h-.001c-.396-.054-.79-.134-1.188-.168a2.814 2.814 0 0 0-.941.076l.019-.004c-.019 0-.036.007-.055.01l-.02-.007c.126-.189.287-.346.474-.464l.006-.004a.101.101 0 0 1 .081-.01h-.001c.576.033 1.148.09 1.718.18a.402.402 0 0 0 .417-.154l.001-.001c.067-.084.135-.18.198-.278l.008-.014c.026-.048.065-.109.126-.09s.074.08.067.16a.943.943 0 0 1-.405.7l-.003.002c-.034.024-.034.038-.01.065c.185.237.298.538.299.866a.728.728 0 0 1-.082.418l.002-.004a.1.1 0 0 0 .031.143c.064.054.122.109.177.167l.001.001a.108.108 0 0 0 .16.01a.71.71 0 0 0 .112-.114l.001-.002c.107-.146.233-.27.377-.372l.005-.003a.209.209 0 0 1 .281.019c.032.038.072.084-.022.115a.725.725 0 0 0-.318.316l-.002.004c-.062.09-.118.187-.175.281a.378.378 0 0 1-.32.185a.395.395 0 0 1-.264-.107zm-3.778-1.2a.495.495 0 0 0-.224-.23l-.003-.001a.492.492 0 0 1-.277-.554l-.001.003a.163.163 0 0 1 .111-.116h.001c.043-.017.061.037.084.061c.065.086.112.19.134.303l.001.005a.352.352 0 0 0 .138.201l.001.001a.323.323 0 0 1 .127.322v-.002c-.012.026.031.091-.022.104h-.011c-.042-.002-.045-.064-.061-.097zm6.436.014c-.08-.014-.17-.018-.254-.025a.576.576 0 0 1-.322-.1l.002.001a.436.436 0 0 0-.234-.067h-.007a.832.832 0 0 1-.366-.148l.002.002a.194.194 0 0 1-.08-.157v-.014v.001c-.014-.16-.014-.16-.17-.199a.843.843 0 0 1-.507-.331l-.002-.002a.633.633 0 0 1-.113-.305v-.003a3.076 3.076 0 0 1-.018-.302c0-.029 0-.07.036-.08a.089.089 0 0 1 .091.038a.367.367 0 0 1 .093.168l.001.002c.016.073.033.134.055.193l-.003-.011c.069.223.262.235.442.27a5.68 5.68 0 0 0 1.901.041l-.031.004a.144.144 0 0 1 .146.033c.038.04.026.055-.017.07c-.094.033-.187.069-.284.103c-.262.105-.566.18-.883.212l-.014.001a2.087 2.087 0 0 1-.396-.058l.014.003a.208.208 0 0 0-.109-.005h.001c.016.088.071.16.145.201l.002.001a.675.675 0 0 0 .366.13h.002c.257-.003.505-.03.746-.079l-.026.004a.304.304 0 0 0 .113-.049l-.001.001q.204-.12.415-.223a.091.091 0 0 1 .118.012c-.223.22-.458.42-.696.616a.187.187 0 0 1-.126.054a.32.32 0 0 1-.033-.007h.001zM7.2 15.911a.463.463 0 0 0-.498-.095l.003-.001a.139.139 0 0 1-.139-.032c-.072-.054-.067-.101.014-.139c.1-.042.218-.078.34-.103l.013-.002c.098.007.187.018.274.034l-.015-.002a2.062 2.062 0 0 0 .902-.002l-.014.003a.828.828 0 0 1 .443.016l-.006-.002c.026.004.05.013.07.026h-.001c.033.032.089.061.08.108c-.008.034-.045.034-.08.034l-.015-.001l-.015.001h.001c-.142.019-.286.039-.43.054a.33.33 0 0 0-.114.041l.001-.001a1.796 1.796 0 0 1-.279.149l-.012.004a.603.603 0 0 1-.2.036a.49.49 0 0 1-.326-.125zM.982 9.76a.46.46 0 0 1-.012-.665a.304.304 0 0 1 .217-.058h-.001a.25.25 0 0 1 .271.186v.002a.444.444 0 0 1-.096.535a.294.294 0 0 1-.192.072a.276.276 0 0 1-.187-.072m16.181-1.853a.298.298 0 0 1 .284-.292h.001c.154 0 .23.08.235.24a.306.306 0 0 1-.265.302h-.01a.267.267 0 0 1-.246-.25v-.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Confused(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.316C6.099 1.316 1.316 6.099 1.316 12S6.099 22.684 12 22.684S22.684 17.901 22.684 12c-.012-5.896-4.788-10.672-10.683-10.684zm0 22.297C5.586 23.613.387 18.414.387 12S5.586.387 12 .387S23.613 5.586 23.613 12v.015c0 6.405-5.192 11.597-11.597 11.597z"/><path fill="currentColor" d="M12 24C5.386 23.966.034 18.614 0 12.003V12C.034 5.386 5.386.034 11.997 0H12c6.614.034 11.966 5.386 12 11.997V12c-.034 6.614-5.386 11.966-11.997 12zM12 .774C5.8.774.774 5.8.774 12S5.8 23.226 12 23.226S23.226 18.2 23.226 12C23.222 5.802 18.198.779 12.001.774zm0 22.297C5.886 23.071.929 18.114.929 12S5.886.929 12 .929c6.122.02 11.084 4.955 11.148 11.065V12c-.064 6.115-5.027 11.051-11.146 11.071zm0-21.368C6.313 1.703 1.703 6.313 1.703 12S6.313 22.297 12 22.297S22.297 17.687 22.297 12c-.026-5.676-4.62-10.271-10.294-10.297z"/><path fill="currentColor" d="M9.677 9.91a2.09 2.09 0 0 1-4.18 0v-.009c0-1.15.932-2.082 2.082-2.082h.018c1.15 0 2.082.932 2.082 2.082v.009zm8.904 0v.009c0 1.15-.932 2.082-2.082 2.082h-.009a2.09 2.09 0 0 1 0-4.18a2.132 2.132 0 0 1 2.09 2.088v.002zM7.355 18.581a.491.491 0 0 1-.464-.384l-.001-.003a.462.462 0 0 1 .078-.311l-.001.002a1.05 1.05 0 0 1 .304-.229l.006-.003l9.832-2.168a.462.462 0 0 1 .311.078l-.002-.001c.094.087.172.189.229.304l.003.006a.44.44 0 0 1-.306.541l-.003.001l-9.91 2.168h-.077z"/><path fill="currentColor" d="m7.355 18.968l-.032.001a.754.754 0 0 1-.741-.615l-.001-.005a.815.815 0 0 1 .079-.623l-.002.004c.145-.168.326-.3.533-.384l.009-.003l9.832-2.245a.815.815 0 0 1 .623.079l-.004-.002c.168.145.3.326.384.533l.003.009a.826.826 0 0 1-.614 1.006l-.006.001l-9.91 2.245zm9.91-3.097l-9.91 2.245l-.077.077l.077.077l9.91-2.245c0-.077.077-.077 0-.155z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Copy(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.092 13.455a1.08 1.08 0 0 0-.25-.692l.001.002v-.005q-.03-.037-.064-.07l-.009-.009l-.029-.029l-4.011-4.011V1.093c0-.603-.489-1.091-1.091-1.091H1.091C.488.002 0 .491 0 1.093v17.454c0 .603.489 1.091 1.091 1.091h3.273v3.273c0 .603.489 1.091 1.091 1.091h14.544c.603 0 1.091-.489 1.091-1.091v-9.422l.001-.033zm-8.001-5.366l4.275 4.275h-4.275zM2.182 17.454V2.182h12.364v4.276l-1.774-1.774a1.03 1.03 0 0 0-.069-.063l-.023-.018l-.053-.04l-.026-.022c-.019-.013-.039-.024-.058-.035l-.027-.012a1.205 1.205 0 0 0-.081-.039l-.034-.014l-.057-.021l-.039-.012l-.063-.016l-.029-.007a1.357 1.357 0 0 0-.092-.014H5.455c-.603 0-1.091.489-1.091 1.091v11.992zm4.363 4.363V6.545h4.363v6.909c0 .602.489 1.09 1.091 1.09h6.91v7.273z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cpanel(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.985 10.902h1.5l-.233.885a.552.552 0 0 1-.532.409h-.701a.702.702 0 0 0-.001 1.402h.427a.246.246 0 0 1 .237.313v-.002l-.268.98h-.465A1.95 1.95 0 0 1 .068 12.43l-.003.014c.2-.888.982-1.542 1.916-1.542zm.757 3.987l1.309-4.88a.546.546 0 0 1 .528-.409h1.398c.935 0 1.716.656 1.909 1.533l.002.013a2.02 2.02 0 0 1-1.937 2.449H4.749l.24-.893a.555.555 0 0 1 .533-.405h.413a.702.702 0 0 0 0-1.404H5.91h.001h-.742l-.967 3.585a.562.562 0 0 1-.533.4h-.927zm8.275-2.612c.011-.038.071-.198-.08-.198H8.774a.55.55 0 0 1-.529-.698l-.001.004l.131-.502h2.919a1.104 1.104 0 0 1 1.067 1.388l.002-.008l-.525 1.943a.897.897 0 0 1-.865.668h-.015h.001l-2.16-.007c-1.609 0-1.444-2.393.026-2.393h1.83l-.131.495a.562.562 0 0 1-.533.405h-.812a.148.148 0 1 0 0 .296h1.309c.173 0 .191-.146.206-.198l.32-1.193zm3.864-1.383h.011a1.943 1.943 0 0 1 1.872 2.466l.003-.014l-.33 1.242a.392.392 0 0 1-.379.293h-.739a.24.24 0 0 1-.229-.313v.002l.4-1.5a.697.697 0 0 0-.671-.88h-.75l-.645 2.4a.395.395 0 0 1-.381.293h-.72a.246.246 0 0 1-.237-.313v.002l.982-3.68h1.812zm3.794 1.416l-.3 1.126a.198.198 0 0 0 .191.251h2.75l-.214.788a.562.562 0 0 1-.533.4h-2.505a1.092 1.092 0 0 1-1.052-1.385l-.002.008l.405-1.498a1.505 1.505 0 0 1 1.448-1.11h1.778a1.096 1.096 0 0 1 1.057 1.388l.002-.008l-.12.45a.889.889 0 0 1-.858.66H18.71l.131-.487a.543.543 0 0 1 .525-.405h.827a.143.143 0 0 0 .139-.112v-.001l.026-.098a.15.15 0 0 0-.146-.191h-1.236l-.022-.001a.275.275 0 0 0-.27.224v.002zm2.633 2.565l1.335-4.992a.4.4 0 0 1 .386-.296h.717a.245.245 0 0 1 .235.313v-.002l-.994 3.68c-.2.75-.873 1.294-1.673 1.294h-.007z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CreditCard(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 21.5V12h36v9.5a2.507 2.507 0 0 1-2.5 2.5H2.466a2.4 2.4 0 0 1-1.731-.734l-.001-.001a2.401 2.401 0 0 1-.735-1.731v-.036zM10 18v2h6v-2zm-6 0v2h4v-2zM33.5 0A2.507 2.507 0 0 1 36 2.5V6H0V2.466A2.4 2.4 0 0 1 .734.735L.735.734a2.401 2.401 0 0 1 1.731-.735h.036H2.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crop(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.455 15.273h2.182V6.546a2.182 2.182 0 0 0-2.182-2.182H8.728v2.182h8.727zm-10.91 2.182V0H4.363v4.364H-.001v2.182h4.364v10.909c0 1.205.977 2.182 2.182 2.182h10.909v4.364h2.182v-4.364H24v-2.182z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Crosshairs(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.703 14H17a1.015 1.015 0 0 1-1-.999V11c.008-.549.451-.992.999-1h1.704a7.038 7.038 0 0 0-4.654-4.69l-.05-.013V7A1.015 1.015 0 0 1 13 8h-2.001a1.015 1.015 0 0 1-1-.999V5.297a7.038 7.038 0 0 0-4.69 4.654l-.013.05h1.703c.549.008.992.451 1 .999v2.001a1.015 1.015 0 0 1-.999 1H5.296a7.038 7.038 0 0 0 4.654 4.69l.05.013v-1.703c.008-.549.451-.992.999-1H13c.549.008.992.451 1 .999v1.704a7.038 7.038 0 0 0 4.69-4.654zM24 11v2a1.015 1.015 0 0 1-.999 1h-2.235c-.799 3.366-3.4 5.966-6.704 6.753l-.061.012v2.234a1.015 1.015 0 0 1-.999 1h-2.001a1.015 1.015 0 0 1-1-.999v-2.235c-3.366-.799-5.966-3.4-6.753-6.704L3.236 14H1.002a1.015 1.015 0 0 1-1-.999V11c.008-.549.451-.992.999-1h2.235c.799-3.366 3.4-5.966 6.704-6.753l.061-.012V1.001c.008-.549.451-.992.999-1h2.001c.549.008.992.451 1 .999v2.235c3.366.799 5.966 3.4 6.753 6.704l.012.061H23c.549.008.992.451 1 .999z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func CssThree(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.366 21.423L21.279 0H0l.025.278L1.91 21.426l.016.176L10.559 24h.136l8.653-2.4l.016-.176zM10.658 4.679h6.183l-.186 2.16l-6.1 2.712l.102.488h5.724l-.659 7.551L10.636 19l-5.084-1.408l-.32-3.616h2.093l.168 1.833l.016.178l3.085.82h.133l3.051-.848l.015-.177l.288-3.386l.023-.276H5.057l-.19-2.173l6.309-2.701l-.1-.49h-6.48l-.185-2.08h6.24z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Cursor(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 0l13.5 12l-5.04 2.24L12 22.5L9 24l-3.593-8.4L0 18z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Curve(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4 0h16a4 4 0 0 1 4 4v16a4 4 0 0 1-4 4H4a4 4 0 0 1-4-4V4a4 4 0 0 1 4-4"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dailymotion(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.068 11.313h-.023a3.1 3.1 0 1 0 .183 6.195h-.024a2.969 2.969 0 0 0 2.926-2.968l-.001-.076v.004l.002-.103a3.054 3.054 0 0 0-3.054-3.054h-.01h.001z"/><path fill="currentColor" d="M0 0v24h24V0zm20.693 20.807h-3.576V19.4a4.837 4.837 0 0 1-3.727 1.47h.011l-.104.001a5.646 5.646 0 0 1-3.83-1.49l.004.004A6.357 6.357 0 0 1 7.27 14.57l.001-.127v-.03a6.34 6.34 0 0 1 2.007-4.635l.003-.003a5.815 5.815 0 0 1 4.147-1.73h.041h-.002a4.186 4.186 0 0 1 3.526 1.578l.007.009V4.157l3.693-.765z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Database(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.286 10.286l.177.001c2.037 0 4.025-.21 5.944-.609l-.188.033a11.686 11.686 0 0 0 4.397-1.728l-.045.027v2.277q0 .924-1.379 1.714a12.162 12.162 0 0 1-3.676 1.242l-.073.01a26.28 26.28 0 0 1-4.976.463l-.189-.001h.01l-.18.001c-1.76 0-3.48-.169-5.146-.49l.17.027a12.39 12.39 0 0 1-3.818-1.285l.067.033q-1.379-.79-1.379-1.714V8.01A11.528 11.528 0 0 0 4.284 9.7l.071.01c1.73.367 3.719.577 5.756.577l.187-.001h-.01zm0 10.285l.177.001c2.037 0 4.025-.21 5.944-.609l-.188.033a11.686 11.686 0 0 0 4.397-1.728l-.045.027v2.277q0 .924-1.379 1.714a12.162 12.162 0 0 1-3.676 1.242l-.073.01a26.28 26.28 0 0 1-4.976.463L10.278 24h.01l-.18.001c-1.76 0-3.48-.169-5.146-.49l.17.027a12.409 12.409 0 0 1-3.818-1.285l.067.033q-1.38-.79-1.38-1.714v-2.277a11.528 11.528 0 0 0 4.282 1.69l.071.01c1.73.367 3.719.577 5.756.577l.187-.001zm0-5.143l.177.001c2.037 0 4.025-.21 5.944-.609l-.188.033a11.686 11.686 0 0 0 4.397-1.728l-.045.027v2.277q0 .924-1.379 1.714a12.162 12.162 0 0 1-3.676 1.242l-.073.01a26.28 26.28 0 0 1-4.976.463l-.189-.001h.01l-.18.001c-1.76 0-3.48-.169-5.146-.49l.17.027a12.409 12.409 0 0 1-3.818-1.285l.067.033q-1.38-.79-1.38-1.714v-2.277a11.528 11.528 0 0 0 4.282 1.69l.071.01c1.73.367 3.719.577 5.756.577l.187-.001zm0-15.429l.174-.001c1.761 0 3.483.169 5.15.491l-.17-.027c1.405.235 2.665.672 3.818 1.285l-.067-.033q1.379.79 1.379 1.714v1.714q0 .924-1.379 1.714a12.162 12.162 0 0 1-3.676 1.242l-.073.01a26.28 26.28 0 0 1-4.976.463l-.189-.001h.01l-.18.001c-1.76 0-3.48-.169-5.146-.49l.17.027a12.409 12.409 0 0 1-3.818-1.285l.067.033Q.002 6.066.002 5.142V3.428q0-.924 1.379-1.714A12.16 12.16 0 0 1 5.058.472l.073-.01A26.28 26.28 0 0 1 10.296 0h-.01z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Date(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22 2.25h-3.25V.75a.75.75 0 0 0-1.5-.001V2.25h-4.5V.75a.75.75 0 0 0-1.5-.001V2.25h-4.5V.75a.75.75 0 0 0-1.5-.001V2.25H2a2 2 0 0 0-2 1.999v17.75a2 2 0 0 0 2 2h20a2 2 0 0 0 2-2V4.249a2 2 0 0 0-2-1.999M22.5 22a.5.5 0 0 1-.499.5H2a.5.5 0 0 1-.5-.5V4.25a.5.5 0 0 1 .5-.499h3.25v1.5a.75.75 0 0 0 1.5.001V3.751h4.5v1.5a.75.75 0 0 0 1.5.001V3.751h4.5v1.5a.75.75 0 0 0 1.5.001V3.751H22a.5.5 0 0 1 .499.499z"/><path fill="currentColor" d="M5.25 9h3v2.25h-3zm0 3.75h3V15h-3zm0 3.75h3v2.25h-3zm5.25 0h3v2.25h-3zm0-3.75h3V15h-3zm0-3.75h3v2.25h-3zm5.25 7.5h3v2.25h-3zm0-3.75h3V15h-3zm0-3.75h3v2.25h-3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DayCloudy(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.106 23.999h-.067a6.106 6.106 0 0 1 0-12.212h.071h-.004c.094 0 .189.007.283.012c1.434-2.545 4.119-4.235 7.199-4.235a8.2 8.2 0 0 1 3.618.835l-.049-.021a5.983 5.983 0 0 1 2.019-1.43l.039-.015a5.842 5.842 0 0 1 2.316-.47a5.942 5.942 0 0 1 2.847 11.159l-.031.016A4.146 4.146 0 0 1 20.843 24h-.007zm-4.413-6.107a4.418 4.418 0 0 0 4.413 4.413h14.732a2.46 2.46 0 1 0-1.591-4.337l.004-.003a.846.846 0 1 1-1.095-1.291l.002-.001a4.128 4.128 0 0 1 1.948-.916l.025-.004a6.54 6.54 0 0 0-11.952-3.629l-.015.023a6.115 6.115 0 0 1 1.767.985l-.012-.009a.846.846 0 1 1-1.059 1.32l.002.001a4.342 4.342 0 0 0-2.741-.967h-.012h.001a4.418 4.418 0 0 0-4.416 4.414zm18.181-9.4c-.477.208-.886.48-1.238.813l.002-.002a8.223 8.223 0 0 1 3.187 6.499v.016c.48.12.902.307 1.28.556l-.016-.01c1.585-.641 2.682-2.167 2.682-3.95a4.25 4.25 0 0 0-4.01-4.243h-.011q-.119-.007-.238-.007h-.01a4.2 4.2 0 0 0-1.654.338l.027-.01zm8.747 13.102l-2.022-2.25a.846.846 0 1 1 1.259-1.131l.001.001l2.022 2.253a.846.846 0 1 1-1.259 1.131l-.001-.001zm4.413-7.71l-3.023-.164a.848.848 0 1 1 .094-1.691h-.002l3.023.164a.847.847 0 0 1-.045 1.691zm-5.78-6.62a.843.843 0 0 1 .065-1.195l.001-.001l2.253-2.022a.848.848 0 0 1 1.131 1.264l-.001.001l-2.254 2.025a.844.844 0 0 1-1.194-.064l-.001-.001zM15.18 6.61l-2.021-2.254a.846.846 0 1 1 1.259-1.131l.001.001l2.019 2.257a.846.846 0 1 1-1.259 1.131l-.001-.001zm6.757-1.893a.848.848 0 0 1-.801-.845l.001-.05v.002L21.301.8a.848.848 0 0 1 1.692.046l-.001.048V.892l-.164 3.024a.848.848 0 0 1-.845.801z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DayHaze(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.601 24a.79.79 0 0 1 0-1.58h7.405a.79.79 0 0 1 0 1.58zm-7.174-3.376a.79.79 0 0 1 0-1.58h21.754a.79.79 0 0 1 0 1.58zm4.013-3.372a.79.79 0 0 1 0-1.58h13.727a.79.79 0 0 1 0 1.58zm12.373-3.155H7.795a.79.79 0 0 1-.79-.79v-.056a6.299 6.299 0 0 1 12.598-.059v.003a.824.824 0 0 1 .007.104v.009a.79.79 0 0 1-.79.79zM8.651 12.514h9.305c-.399-2.238-2.33-3.915-4.653-3.915s-4.254 1.677-4.648 3.887l-.004.029zm13.893 1.583a.79.79 0 0 1 0-1.58h3.272a.79.79 0 0 1 0 1.58zm-21.754 0a.79.79 0 0 1 0-1.58h3.272a.79.79 0 0 1 0 1.58zm18.489-6.766a.787.787 0 0 1 0-1.118l2.313-2.314a.79.79 0 1 1 1.116 1.113l-.001.001l-2.313 2.314c-.143.143-.34.231-.559.231s-.416-.088-.559-.231zm-13.069 0L3.896 5.014a.79.79 0 0 1 .559-1.35c.218 0 .416.088.559.232L7.328 6.21a.79.79 0 1 1-1.116 1.118zm6.303-3.264V.79a.79.79 0 0 1 1.58 0v3.273a.79.79 0 0 1-1.58 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DayLightning(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.631 23.837a.552.552 0 0 1 0-.786l.68-.679h-.222a.52.52 0 0 1-.112-.012l.004.001a.471.471 0 0 1-.052-.015l.003.001a.587.587 0 0 1-.059-.019l.003.001a.415.415 0 0 1-.056-.03l.002.001l-.042-.022a.54.54 0 0 1-.153-.152l-.001-.002c-.009-.013-.015-.027-.022-.04a.612.612 0 0 1-.028-.052l-.002-.004a.634.634 0 0 1-.016-.051l-.001-.004a.363.363 0 0 1-.014-.046l-.001-.003a.53.53 0 0 1 .001-.221l-.001.004a.462.462 0 0 1 .016-.052l-.001.003a.524.524 0 0 1 .018-.059l-.001.004a.54.54 0 0 1 .031-.059l-.001.002a.376.376 0 0 1 .023-.041l-.001.002a.536.536 0 0 1 .07-.086l1.563-1.563a.554.554 0 0 1 .949.393a.55.55 0 0 1-.163.393l-.615.615h.222c.039 0 .076.004.113.012l-.004-.001c.017.003.032.009.049.014a.514.514 0 0 1 .059.019l-.004-.001a.43.43 0 0 1 .057.03l-.003-.001l.042.022a.598.598 0 0 1 .153.152l.001.002a.322.322 0 0 1 .02.037l.001.003a.33.33 0 0 1 .028.053l.001.003a.752.752 0 0 1 .016.05l.001.004c.005.012.01.029.014.045l.001.003a.53.53 0 0 1 0 .221l.001-.004c-.003.017-.01.033-.015.049l-.018.058l.001-.004a.54.54 0 0 1-.031.059l.001-.002c-.007.013-.013.027-.021.04a.671.671 0 0 1-.07.086l-1.627 1.627a.552.552 0 0 1-.786 0zm-2.214-2.943a.554.554 0 0 1 0-.786l1.745-1.745H4.916a.562.562 0 0 1-.11-.011l.004.001a.462.462 0 0 1-.052-.016l.003.001a.369.369 0 0 1-.056-.018l.003.001a.375.375 0 0 1-.06-.032l.002.001l-.038-.021a.575.575 0 0 1-.154-.153l-.001-.002l-.018-.033a.539.539 0 0 1-.032-.059l-.001-.003c-.006-.015-.009-.031-.015-.049c-.006-.015-.011-.034-.016-.053l-.001-.004a.54.54 0 0 1-.009-.085v-.046a.4.4 0 0 1 .027-.143l-.001.003c.005-.016.008-.031.015-.049a.395.395 0 0 1 .035-.065l-.001.002l.018-.033a.551.551 0 0 1 .071-.086l1.637-1.637H4.017a4.016 4.016 0 1 1 0-8.032c.064 0 .128.005.192.008c.944-1.668 2.705-2.778 4.725-2.784h.009c.85 0 1.654.197 2.368.548l-.032-.014a3.92 3.92 0 0 1 1.327-.94l.025-.01a3.902 3.902 0 0 1 3.39 7.017l-.02.01a2.727 2.727 0 0 1-2.296 4.198h-1.708l-.452.452h.223a.58.58 0 0 1 .111.011l-.004-.001c.02.005.036.01.052.016l-.003-.001c.022.005.04.011.057.018l-.003-.001a.46.46 0 0 1 .059.031l-.002-.001l.039.021a.571.571 0 0 1 .154.152l.001.002c.009.013.015.028.022.042l.027.05l.002.004a.474.474 0 0 1 .017.053l.001.004c.005.013.01.029.014.046l.001.003a.53.53 0 0 1 0 .222l.001-.004c-.003.017-.01.032-.014.049a.61.61 0 0 1-.019.06l.001-.004a.72.72 0 0 1-.029.056l.001-.003c-.008.014-.014.029-.023.043a.638.638 0 0 1-.068.084l-1.628 1.628a.554.554 0 0 1-.949-.393c0-.153.062-.292.163-.393l.681-.679h-.221a.52.52 0 0 1-.112-.012l.004.001c-.017-.003-.032-.01-.049-.014a.397.397 0 0 1-.058-.018l.003.001a.445.445 0 0 1-.058-.031l.003.002c-.013-.007-.027-.013-.04-.022a.554.554 0 0 1-.153-.152l-.001-.002c-.009-.014-.015-.028-.023-.042a.387.387 0 0 1-.026-.05L9.698 17a.434.434 0 0 1-.017-.054l-.001-.004a.37.37 0 0 1-.013-.046l-.001-.003a.53.53 0 0 1 0-.222l-.001.004c.003-.016.009-.032.014-.049a.535.535 0 0 1 .019-.06l-.001.004a.69.69 0 0 1 .029-.055l-.001.002c.008-.014.014-.029.023-.043a.712.712 0 0 1 .068-.084l.615-.615H7.74L6.265 17.25H7.51c.039 0 .076.004.112.012l-.004-.001c.016.003.03.009.046.013a.578.578 0 0 1 .062.019l-.004-.001q.025.012.049.026l.049.025a.57.57 0 0 1 .075.061l.01.008l.011.014c.02.021.039.044.056.068l.001.002q.015.024.027.049l.023.046l.001.002a.4.4 0 0 1 .019.061l.001.003c.004.013.009.026.012.04a.537.537 0 0 1 0 .222l.001-.004c-.003.014-.008.026-.012.039a.514.514 0 0 1-.021.069l.001-.004l-.024.049a.387.387 0 0 1-.028.05l.001-.002a.558.558 0 0 1-.068.083l-2.695 2.695a.552.552 0 0 1-.786 0zm-3.31-9.134a2.908 2.908 0 0 0 2.904 2.905h9.692a1.62 1.62 0 1 0-1.048-2.856l.002-.002a.556.556 0 0 1-.719-.848l.001-.001a2.707 2.707 0 0 1 1.276-.6l.017-.002a4.295 4.295 0 0 0-7.854-2.39l-.01.015c.441.16.822.377 1.159.647l-.008-.006a.555.555 0 1 1-.695.868l.001.001a2.853 2.853 0 0 0-1.802-.635h-.011h.001a2.907 2.907 0 0 0-2.905 2.903zm11.955-6.185a2.79 2.79 0 0 0-.813.534l.001-.001a5.395 5.395 0 0 1 2.092 4.268v.018c.314.077.59.198.838.359l-.011-.006a2.79 2.79 0 0 0-.869-5.381h-.007q-.076-.004-.152-.004h-.007c-.387 0-.756.079-1.091.223l.018-.007zm5.743 8.604l-1.328-1.478a.556.556 0 1 1 .827-.743v.001l1.325 1.475a.556.556 0 1 1-.827.743v-.001zm2.898-5.063l-1.984-.108a.556.556 0 1 1 .061-1.11h-.001l1.984.108a.556.556 0 0 1-.03 1.111zm-3.795-4.347a.553.553 0 0 1 .042-.784h.001l1.479-1.326a.556.556 0 1 1 .743.827h-.001l-1.479 1.328a.553.553 0 0 1-.784-.042v-.001zM9.98 4.34L8.652 2.86a.553.553 0 0 1 .413-.926c.164 0 .312.071.413.184v.001l1.328 1.479a.556.556 0 1 1-.827.743V4.34zm4.437-1.243a.556.556 0 0 1-.526-.555l.001-.033v.002l.107-1.985a.556.556 0 0 1 1.111.03l-.001.031V.586l-.107 1.985a.558.558 0 0 1-.555.526z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DayRain(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.75 23.818a.617.617 0 0 1 0-.876l3.205-3.205a.618.618 0 0 1 1.066.428a.616.616 0 0 1-.191.447l-3.206 3.205a.618.618 0 0 1-.876 0zm6.521-3.033a.617.617 0 0 1 0-.876l2.354-2.354H6.084l-3.23 3.23a.618.618 0 1 1-.876-.875l2.366-2.367a4.465 4.465 0 1 1 .246-8.924h-.006c.07 0 .138.005.207.009c1.054-1.858 3.016-3.093 5.268-3.099h.01c.947 0 1.843.22 2.639.611l-.035-.016a4.36 4.36 0 0 1 1.477-1.046l.028-.011a4.347 4.347 0 0 1 3.778 7.818l-.023.011a3.033 3.033 0 0 1-2.563 4.654h-.993l-3.23 3.23a.617.617 0 0 1-.876 0zm-8.915-7.698a3.232 3.232 0 0 0 3.228 3.229h10.779a1.8 1.8 0 1 0-1.164-3.173l.003-.002a.619.619 0 0 1-.801-.944l.001-.001a3.02 3.02 0 0 1 1.426-.671l.018-.003a4.785 4.785 0 0 0-8.744-2.655l-.011.017c.492.178.917.42 1.292.721l-.009-.007a.62.62 0 1 1-.774.967l.001.001a3.173 3.173 0 0 0-2.005-.707h-.009a3.233 3.233 0 0 0-3.231 3.229zM14.658 6.21c-.348.152-.647.35-.905.593l.002-.002a6.012 6.012 0 0 1 2.331 4.755v.011c.351.087.66.225.937.406l-.011-.007a3.116 3.116 0 0 0 1.96-2.889a3.11 3.11 0 0 0-2.932-3.105h-.008q-.085-.004-.17-.004h-.008c-.432 0-.843.089-1.216.249zm6.399 9.587l-1.479-1.647a.619.619 0 1 1 .921-.828l.001.001l1.479 1.648a.619.619 0 1 1-.921.828zm3.23-5.642l-2.212-.12a.62.62 0 0 1 .033-1.237l.035.001h-.002l2.211.12a.62.62 0 0 1-.033 1.237zm-4.229-4.84a.617.617 0 0 1 .047-.874h.001l1.648-1.479a.619.619 0 1 1 .828.921l-.001.001l-1.648 1.479a.617.617 0 0 1-.874-.047v-.001zm-8.834-.479L9.745 3.187a.619.619 0 1 1 .923-.828v.001l1.478 1.649a.621.621 0 0 1-.925.829l-.001-.001zm4.944-1.385a.62.62 0 0 1-.585-.654v.002l.12-2.212A.62.62 0 0 1 16.94.62l-.001.035V.653l-.12 2.212a.62.62 0 0 1-.618.586z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DaySnow(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.875 23.43a.57.57 0 1 1 1.14 0a.57.57 0 0 1-1.14 0m1.633-1.632a.57.57 0 1 1 1.14 0a.57.57 0 0 1-1.14 0m4.605-.744a.567.567 0 1 1-.001.002v-.004zm-8.772 0a.566.566 0 1 1-.001.004v-.004zm5.796-.887a.567.567 0 1 1 0 .005zm4.608-.744a.57.57 0 1 1 1.14 0a.57.57 0 0 1-1.14 0m-8.776 0a.57.57 0 1 1 1.14 0a.57.57 0 0 1-1.14 0m-.438-1.619h-.05a4.53 4.53 0 1 1 0-9.062h.053h-.003c.07 0 .14.006.21.009c1.064-1.888 3.057-3.143 5.342-3.143c.963 0 1.874.223 2.684.62l-.036-.016a4.409 4.409 0 1 1 5.36 6.856l-.023.012a3.076 3.076 0 0 1-2.6 4.72h-.005zm-3.275-4.531a3.278 3.278 0 0 0 3.274 3.275h10.929a1.825 1.825 0 1 0-1.181-3.217l.003-.002a.629.629 0 1 1-.815-.958l.001-.001a3.067 3.067 0 0 1 1.446-.681l.019-.003a4.852 4.852 0 0 0-8.868-2.693l-.011.017c.499.18.93.426 1.311.731l-.009-.007a.63.63 0 1 1-.787.981l.001.001a3.222 3.222 0 0 0-2.034-.718h-.009a3.278 3.278 0 0 0-3.272 3.274zm13.492-6.975a3.174 3.174 0 0 0-.918.601l.002-.002a6.098 6.098 0 0 1 2.364 4.823v.012c.356.089.669.228.95.412l-.012-.007a3.16 3.16 0 0 0 1.99-2.931a3.153 3.153 0 0 0-2.976-3.148h-.008q-.086-.004-.171-.004h-.008c-.438 0-.856.09-1.235.253l.02-.008zm9.765 4.002l-2.243-.121a.628.628 0 0 1-.594-.627l.001-.037v.002a.62.62 0 0 1 .664-.593h-.002l2.243.121a.628.628 0 0 1-.034 1.255zm-4.288-4.913a.624.624 0 0 1 .047-.886h.001l1.671-1.497a.628.628 0 1 1 .84.934h-.001l-1.672 1.5a.624.624 0 0 1-.886-.047V5.39zm-8.96-.485l-1.5-1.67a.628.628 0 0 1 .934-.839l.001.001l1.498 1.673a.628.628 0 0 1-.934.839l-.001-.001zm4.429-1.927L15.313.732A.628.628 0 1 1 16.551.52v.003l.381 2.243a.629.629 0 0 1-.51.724l-.004.001a.628.628 0 0 1-.724-.51l-.001-.004z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DaySunny(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.289 23.287v-2.952a.713.713 0 1 1 1.426 0v2.952a.713.713 0 1 1-1.426 0m8.19-2.804l-2.087-2.086a.713.713 0 1 1 1.008-1.008l2.085 2.089a.713.713 0 0 1-1.009 1.002l.001.001zm-15.962 0a.71.71 0 0 1 0-1.008l2.087-2.087a.713.713 0 1 1 1.008 1.008l-2.087 2.086a.71.71 0 0 1-1.008 0zm2.803-8.485a5.683 5.683 0 1 1 11.366 0a5.683 5.683 0 0 1-11.366 0m1.425 0a4.26 4.26 0 1 0 8.518 0a4.26 4.26 0 0 0-8.518 0m12.591.713a.713.713 0 1 1 0-1.426h2.952a.713.713 0 1 1 0 1.426zm-19.623 0a.713.713 0 1 1 0-1.426h2.953a.713.713 0 1 1 0 1.426zM17.39 6.608a.71.71 0 0 1 0-1.008l2.087-2.087a.713.713 0 1 1 1.008 1.008l-2.087 2.087a.71.71 0 0 1-1.008 0m-11.788 0L3.517 4.523a.713.713 0 1 1 1.008-1.008l2.087 2.087A.713.713 0 1 1 5.604 6.61zm5.685-2.944V.713a.713.713 0 1 1 1.426 0v2.952a.713.713 0 1 1-1.426 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Deaf(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m23.25.562l.56.56l-2.92 2.92l.771.771c.27.255.52.524.752.809l.012.015c0 .03-.052.052-.12.052s-1.258.232-2.658.524s-2.56.524-2.576.524c.03-.254.073-.478.128-.696l-.008.037c.075-.359.292-1.407.48-2.336s.374-1.85.419-2.051l.09-.359l1.55 1.55l1.446-1.446c.786-.786 1.453-1.44 1.475-1.44c.217.181.413.366.597.563l.003.004zM12.804 4.755a6.867 6.867 0 0 1 4.75 4.672l.012.048a5.044 5.044 0 0 1 .178 1.854l.001-.019a5.521 5.521 0 0 1-.704 3.147l.014-.028a10.35 10.35 0 0 1-2.002 2.601l-.005.005c-2.121 2.284-2.351 2.614-2.591 3.707a3.248 3.248 0 0 1-.899 1.811a2.917 2.917 0 0 1-2.806.775l.02.004c-.88-.254-2-1.266-2-1.798v-.019a.496.496 0 0 1 .499-.498l.035.001h-.002c.187 0 .307.09.673.486c.337.484.887.8 1.511.809h.001l.085.002c.432 0 .825-.169 1.116-.445l-.001.001c.4-.356.666-.854.719-1.414l.001-.009c.285-1.28.457-1.552 2.494-3.776c.756-.823 1.5-1.67 1.655-1.879a6.866 6.866 0 0 0 .984-2.055l.012-.049a7.175 7.175 0 0 0-.008-2.816l.008.046c-.587-2.264-2.466-3.964-4.785-4.28l-.03-.003a6.08 6.08 0 0 0-5.813 2.489l-.013.018a10.404 10.404 0 0 0-1.324 2.466l-.024.073a1.501 1.501 0 0 1-.265.512l.002-.003a.546.546 0 0 1-.415.096h.003a.552.552 0 0 1-.358-.222l-.001-.002c-.15-.202-.15-.218-.022-.64c.164-.469.322-.845.497-1.211l-.031.072c.909-2.108 2.622-3.724 4.743-4.482l.057-.018a5.62 5.62 0 0 1 1.787-.286c.183 0 .363.009.541.025l-.023-.002a5.268 5.268 0 0 1 1.759.237l-.037-.01zm-.8 2.24c.797.202 1.477.62 1.997 1.187l.003.003a4.324 4.324 0 0 1 1.235 2.12l.006.03c.299 1.25-.187 3.242-.854 3.474a.356.356 0 0 1-.414-.079c-.278-.202-.292-.48-.052-1.01c.202-.41.32-.893.32-1.404c0-.546-.135-1.06-.374-1.511l.008.018c-1.213-2.471-4.635-2.456-6.2.022c-.21.337-.352.486-.487.509a.58.58 0 0 1-.628-.334l-.001-.004c-.09-.179-.08-.247.052-.524a5.502 5.502 0 0 1 2.201-2.127l.03-.014a5.134 5.134 0 0 1 2.132-.456c.361 0 .713.037 1.053.106l-.034-.006zm-2.194 6.14c.488.26.876.653 1.123 1.131l.007.015c.086.241.135.518.135.808a1.817 1.817 0 0 1-.207.97l.005-.01a2.175 2.175 0 0 1-1.516 1.198l-.014.002a2.156 2.156 0 0 1-1.878-.514l.002.002a2.121 2.121 0 0 1-.735-1.48v-.006a2.12 2.12 0 0 1 .557-1.564l-.001.002a2.005 2.005 0 0 1 1.628-.734h-.004a1.816 1.816 0 0 1 .91.185zm-3.12 4.48c-.038.179-.285 1.378-.546 2.673s-.494 2.358-.509 2.374a8.915 8.915 0 0 1-.795-.751l-.005-.005l-.786-.786l-2.886 2.886l-.577-.584l-.584-.577l2.886-2.886l-.749-.75a9.016 9.016 0 0 1-.737-.78l-.011-.014c0-.022.951-.24 2.121-.48s2.298-.472 2.524-.524a3.06 3.06 0 0 1 .556-.102l.013-.001c.156-.008.156.014.088.306z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Delicious(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 12H0v12h12zM24 0H12v12h12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Deskpro(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m9.658 23.998l1.458-1.8l.123.002a4.18 4.18 0 0 0 1.616-.322l-.028.01a3.344 3.344 0 0 0 1.795-2.96v-.032v.002c-.005-.234-.038-.206-1.355-1.754h3.711a9.091 9.091 0 0 1-.5 3.612l.02-.064c-.851 1.734-2.487 2.97-4.432 3.25l-.031.004c-.408.051-1.538.056-2.377.051zm9.395-7.861a4.13 4.13 0 0 0 3.45-1.443l.005-.006l-8.227-.014l-.816 1.467zM17.009 5.641v7.892h-3.446l-1.533 2.602h-.305c-2.452 0-2.72-2.649-2.733-2.762a.473.473 0 0 0-.944.04l.002.042v-.002a4.448 4.448 0 0 0 1.239 2.678l-.001-.001H5.022c-1.195 0-5.194-1.256-5.016-5.344V.789c0-.75.455-.776.703-.788h3.844c.01 0 .014.005.023.005c.202-.018 2.35-.098 2.35 2.049a.474.474 0 1 0 .948 0c.002-.036.002-.078.002-.12c0-.744-.283-1.421-.746-1.931l.002.002h3.947c5.096-.018 5.93 3.722 5.93 5.635m-6.211 2.625l3 1.374c.621-2.134-1.978-3.362-3-1.374"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Desktop(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24.749 0H2.251a2.25 2.25 0 0 0-2.25 2.25v14.999a2.25 2.25 0 0 0 2.25 2.25h8.999l-.75 2.25H7.124a1.125 1.125 0 0 0 0 2.25h12.75a1.125 1.125 0 0 0 0-2.25h-3.375l-.75-2.249h8.999a2.25 2.25 0 0 0 2.25-2.249V2.25A2.25 2.25 0 0 0 24.748 0zm-.751 16.499H2.999V3h20.999z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Deviantart(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14.642 4.794l.23-.43V0h-4.367l-.436.44l-2.058 3.925l-.646.435H.015v5.993h4.04l.36.436l-4.18 7.981l-.24.43V24h4.37l.436-.44l2.07-3.925l.644-.436h7.35v-5.993h-4.05l-.36-.438l4.186-7.977z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Digg(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.84 8.364V4.802h2.394v11.77H0V8.365zm0 6.282v-4.359H2.39v4.359zm13.92-6.282H24v11.44h-6.234v-1.92h3.84v-1.313h-3.84zm3.84 6.282v-4.359h-1.44v4.359zm-11.062 1.92V8.364h6.234v11.44h-6.225v-1.92h3.84v-1.313zm2.39-6.282v4.359h1.453v-4.359zM9.6 4.8v2.39H7.172V4.8zm0 3.562v8.209H7.172V8.364z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DinnersClub(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.4 24A2.4 2.4 0 0 1 0 21.6V2.4A2.4 2.4 0 0 1 2.4 0h31.2A2.4 2.4 0 0 1 36 2.4v19.2a2.4 2.4 0 0 1-2.4 2.4zM9.741 5.6A8.557 8.557 0 0 0 7.2 11.848v-.008a8.355 8.355 0 0 0 2.526 5.99l.002.002a8.558 8.558 0 0 0 6.053 2.497h4.123a9.17 9.17 0 0 0 6.211-2.501l-.005.004a8.173 8.173 0 0 0 2.707-5.987v-.068a8.353 8.353 0 0 0-2.714-6.172l-.006-.006a9.003 9.003 0 0 0-2.796-1.748l-.061-.021a9.135 9.135 0 0 0-3.366-.631h-4.145a8.642 8.642 0 0 0-5.992 2.403zm-1.726 6.145a7.84 7.84 0 1 1 15.68-.006a7.84 7.84 0 0 1-15.68.009zm9.618 4.64c1.884-.734 3.196-2.534 3.196-4.64s-1.311-3.906-3.162-4.628l-.034-.012zm-6.746-4.64a4.975 4.975 0 0 0 3.154 4.628l.034.012V7.108c-1.879.739-3.186 2.537-3.188 4.64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DirectionSign(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.497 17.679h4.437V24H9.497zM1.906 4.054L-.001 5.961l1.907 1.906h12.026V4.053zm0 9.331l-1.907 1.906l1.907 1.9h12.026v-3.806zm19.616-4.662H9.736v3.814h11.786l1.907-1.907zM11.71 0L9.496 2.545v1.028h4.437V2.545z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Discord(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.93 11.4a1.23 1.23 0 1 1-1.225-1.331c.683.029 1.225.59 1.225 1.277l-.001.056v-.003zm-5.604-1.33a1.338 1.338 0 0 0-.005 2.664h.005a1.278 1.278 0 0 0 1.225-1.277l-.001-.056v.003l.002-.067c0-.685-.541-1.243-1.219-1.269h-.002zM21 2.472V24c-3.023-2.672-2.057-1.787-5.568-5.052l.636 2.22H2.459a2.467 2.467 0 0 1-2.46-2.466V2.466A2.466 2.466 0 0 1 2.458 0h16.081a2.467 2.467 0 0 1 2.46 2.466zm-3.42 11.376a16.004 16.004 0 0 0-1.77-7.086l.042.09a5.885 5.885 0 0 0-3.358-1.259l-.014-.001l-.168.192c1.15.312 2.15.837 3.002 1.535l-.014-.011a10.068 10.068 0 0 0-4.839-1.222c-1.493 0-2.911.321-4.189.898l.064-.026c-.444.204-.708.35-.708.35a8.036 8.036 0 0 1 3.1-1.56l.056-.012l-.12-.144A5.912 5.912 0 0 0 5.28 6.861l.012-.009c-1.052 2.036-1.686 4.437-1.728 6.982v.014a4.352 4.352 0 0 0 3.666 1.824h-.006s.444-.54.804-.996a3.75 3.75 0 0 1-2.093-1.406l-.007-.01c.176.124.468.284.49.3a8.638 8.638 0 0 0 4.188 1.067a8.73 8.73 0 0 0 3.36-.668l-.058.021c.528-.202.982-.44 1.404-.723l-.025.016a3.8 3.8 0 0 1-2.144 1.423l-.026.005c.36.456.792.972.792.972l.111.001a4.378 4.378 0 0 0 3.552-1.813l.009-.013z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Discourse(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.102 0h-.081C5.462 0 .13 5.252.001 11.779v.012L.007 24l12.097-.01c6.582-.055 11.897-5.404 11.897-11.995S18.686.056 12.109 0h-.005zM12 18.857h-.015a6.778 6.778 0 0 1-2.94-.666l.041.018l-4.345 1.077l1.227-4.018a6.78 6.78 0 0 1-.83-3.27A6.86 6.86 0 1 1 12 18.857"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Discover(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M33.6 24H2.4A2.4 2.4 0 0 1 0 21.6V2.4A2.4 2.4 0 0 1 2.4 0h31.2A2.4 2.4 0 0 1 36 2.4v19.2a2.4 2.4 0 0 1-2.4 2.4M19.112 9.6a2.574 2.574 0 1 0 0 5.148a2.574 2.574 0 0 0 0-5.148m2.334.144l2.034 4.914h.502l2.07-4.914h-1.013l-1.296 3.216l-1.28-3.219h-1.019zm-11.98 3.413l-.602.582c.329.55.922.913 1.599.913l.064-.001h-.003a1.545 1.545 0 0 0 1.653-1.541v-.034v.002c0-.76-.32-1.109-1.379-1.495c-.56-.209-.726-.344-.726-.602c0-.303.3-.532.697-.532a.976.976 0 0 1 .753.385l.002.002l.488-.64a2.07 2.07 0 0 0-1.389-.532h-.02h.001a1.405 1.405 0 0 0-1.494 1.369v.002c0 .666.298.998 1.178 1.316c.25.08.464.172.666.283l-.02-.01a.576.576 0 0 1 .281.494v.002a.68.68 0 0 1-.679.677l-.049-.002h.002h-.013c-.443 0-.825-.259-1.005-.633l-.003-.007zm5.555-3.52h-.015a2.55 2.55 0 0 0-1.792.732a2.498 2.498 0 0 0 1.758 4.273h.03h-.001c.427 0 .83-.103 1.186-.286l-.015.007v-1.097c-.274.32-.677.522-1.128.524l-.051.001a1.562 1.562 0 0 1-1.559-1.658v.004l-.001-.055c0-.87.699-1.577 1.566-1.589H15c.467.003.885.21 1.17.537l.002.002v-1.1a2.349 2.349 0 0 0-1.142-.293h-.008zm15.657 2.974h.122l1.293 1.926h1.15l-1.51-2.018a1.282 1.282 0 0 0 1.092-1.365v.004c0-.898-.618-1.414-1.695-1.414h-1.385v4.792h.934zM26.48 9.744v4.792h2.648v-.811H27.41v-1.294h1.651v-.812H27.41v-1.063h1.718v-.811zm-18.899 0v4.792h.933V9.744zm-4.381 0v4.792h1.367a2.49 2.49 0 0 0 1.714-.556l-.005.004a2.417 2.417 0 0 0 .87-1.839v-.027a2.382 2.382 0 0 0-2.581-2.374l.009-.001zm1.186 3.981h-.251v-3.17h.251a1.71 1.71 0 0 1 1.282.391l-.003-.002c.318.295.516.715.516 1.181v.012v-.001v.01c0 .471-.198.896-.515 1.196l-.001.001a1.763 1.763 0 0 1-1.287.382l.007.001zm26.565-1.775h-.272v-1.452h.287c.588 0 .899.246.899.711c0 .484-.316.74-.914.74z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dislike(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.355 23.783a3.262 3.262 0 0 1-.506-.299l.01.007l-.106-.071a1.018 1.018 0 0 1-.501-.599l-.002-.007a116.75 116.75 0 0 1-.518-1.964l-.066-.259a7.842 7.842 0 0 1-.244-1.171l-.005-.043c.172-1.574.501-3.011.976-4.378l-.042.138l-7.692.012h-.027A1.644 1.644 0 0 1 0 13.303l-.001.008a1.983 1.983 0 0 1 .897-1.785l.008-.005a2.178 2.178 0 0 1-.371-1.401v.007a1.846 1.846 0 0 1 1.135-1.871l.012-.004a2.385 2.385 0 0 1-.298-1.311v.006a1.905 1.905 0 0 1 .885-1.735l.008-.005a2.782 2.782 0 0 1-.224-1.462l-.001.013v-.185a1.834 1.834 0 0 1 2.007-1.719l-.007-.001h10.373c.051 0 .105 0 .16-.007a2.9 2.9 0 0 1 .214-.01h.016c.104 0 .204.014.299.04l-.008-.002c.186.063.348.147.493.253l-.005-.004l.1.063c.234.144.462.298.69.451l.102.072c.131.094.194.136.262.133c.24-.005.48-.007.72-.01v9.155l.013.769c-.451.288-.867.56-.993.664a7.407 7.407 0 0 1-.523.901l.017-.026c-.085.134-.17.265-.248.4l-1.729 2.89c-.08.136-.167.271-.254.407c-.164.26-.329.52-.465.79a.883.883 0 0 0-.079.521l-.001-.005l.006 3.12a1.525 1.525 0 0 1-.679 1.074l-.006.004a1.97 1.97 0 0 1-1.155.499h-.093a2.135 2.135 0 0 1-.931-.222l.012.006zm7.76-10.757V0h7.016c.17 0 .31.13.326.295v.001l1.218 12.368l.002.032a.329.329 0 0 1-.329.329zm3.624-1.76c0 .181.147.328.328.328h2.511a.329.329 0 1 0 0-.658h-2.51a.327.327 0 0 0-.328.327zm-1.873 0c0 .181.147.328.328.328h.774a.329.329 0 1 0 0-.658h-.774a.327.327 0 0 0-.327.327zm-.694-8a1.3 1.3 0 1 0 2.596.002a1.3 1.3 0 0 0-2.595-.003zm.657 0a.64.64 0 1 1 1.28.001a.64.64 0 0 1-1.279-.001v-.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Disqus(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.807 24h-.014a11.742 11.742 0 0 1-7.699-2.861l.015.013l-5.109.7l1.974-4.874a11.83 11.83 0 0 1-1.073-4.964C.901 5.416 6.22.06 12.804-.001h.006c6.627 0 12 5.373 12 12s-5.373 12-12 12zm6.503-12.034v-.034c0-3.463-2.443-5.931-6.654-5.931H8.111v12.002h4.479c4.242 0 6.719-2.574 6.719-6.037zm-6.606 3.089h-1.328V8.952h1.328a2.92 2.92 0 0 1 3.256 3.038v-.006v.031a2.914 2.914 0 0 1-3.264 3.036l.014.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dizzy(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.316C6.099 1.316 1.316 6.099 1.316 12S6.099 22.684 12 22.684S22.684 17.901 22.684 12c-.012-5.896-4.788-10.672-10.683-10.684zm0 22.297C5.586 23.613.387 18.414.387 12S5.586.387 12 .387S23.613 5.586 23.613 12v.015c0 6.405-5.192 11.597-11.597 11.597z"/><path fill="currentColor" d="M12 24C5.386 23.966.034 18.614 0 12.003V12C0 5.373 5.373 0 12 0s12 5.373 12 12c-.034 6.614-5.386 11.966-11.997 12zM12 .774C5.8.774.774 5.8.774 12S5.8 23.226 12 23.226S23.226 18.2 23.226 12C23.222 5.802 18.198.779 12.001.774zm0 22.297C5.886 23.071.929 18.114.929 12S5.886.929 12 .929S23.071 5.886 23.071 12S18.114 23.071 12 23.071m0-21.368C6.313 1.703 1.703 6.313 1.703 12S6.313 22.297 12 22.297S22.297 17.687 22.297 12c-.026-5.676-4.62-10.271-10.294-10.297z"/><path fill="currentColor" d="M16.413 17.187H7.51a.499.499 0 0 1-.464-.463v-.002a.499.499 0 0 1 .463-.464h8.905a.499.499 0 0 1 .464.463v.021c0 .246-.2.446-.446.446h-.02h.001z"/><path fill="currentColor" d="M16.413 17.574H7.51a.852.852 0 0 1 0-1.704h8.903a.852.852 0 0 1 0 1.704m-8.826-.929c-.077 0-.077.077 0 0c-.02.02-.033.047-.033.077s.013.058.033.077h8.903a.077.077 0 0 0 .077-.077c0-.077 0-.077-.077-.077zM8.671 9.91l1.084-1.084a.56.56 0 0 0 0-.774a.605.605 0 0 0-.852 0L7.819 9.136L6.735 8.052a.605.605 0 0 0-.852 0a.605.605 0 0 0 0 .852l1.084 1.084l-1.084 1.084a.605.605 0 0 0 0 .852a.56.56 0 0 0 .774 0l1.084-1.084l1.084 1.084a.56.56 0 0 0 .774 0a.605.605 0 0 0 0-.852zm8.361 0l1.084-1.084a.605.605 0 0 0 0-.852a.605.605 0 0 0-.852 0L16.18 9.058l-1.084-1.084a.605.605 0 0 0-.852 0a.56.56 0 0 0 0 .774l1.084 1.084l-1.084 1.084a.605.605 0 0 0 0 .852a.56.56 0 0 0 .774 0l1.084-1.084l1.084 1.084a.56.56 0 0 0 .774 0a.605.605 0 0 0 0-.852z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dna(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.47 12.354a6.912 6.912 0 0 1-2.392 1.14l-.048.011c1.428.65 2.405 2.063 2.409 3.704V24h2.293v-6.79a6.365 6.365 0 0 0-2.252-4.848zm-8.208-.724a6.921 6.921 0 0 1 2.392-1.14l.048-.011c-1.428-.65-2.405-2.063-2.409-3.704V-.001H0v6.778a6.365 6.365 0 0 0 2.252 4.845l.009.008z"/><path fill="currentColor" d="M10.431 0c0 7.18.026 6.822-.039 7.321a4.071 4.071 0 0 1-4.025 3.52v.006a6.385 6.385 0 0 0-6.362 6.361V24h2.293v-6.79a4.074 4.074 0 0 1 4.063-4.062h.001v-.006a6.38 6.38 0 0 0 6.362-6.361V0zM3.349.952h6.03a.48.48 0 0 0 .48-.48a.473.473 0 0 0-.473-.473H3.349a.48.48 0 1 0 .002.952h-.002z"/><path fill="currentColor" d="M3.349 3.202h6.03a.48.48 0 0 0 0-.96h-6.03a.48.48 0 0 0 0 .96m0 2.255h6.03a.48.48 0 0 0 0-.96h-6.03a.48.48 0 0 0-.48.48a.488.488 0 0 0 .48.48m6.393 1.774a.474.474 0 0 0-.474-.473H3.457a.481.481 0 0 0-.002.952h5.809a.48.48 0 0 0 .48-.48zm-4.95 1.712a.48.48 0 0 0 0 .96h3.149a.48.48 0 0 0 0-.96zm4.593 14.079h-6.03a.48.48 0 0 0 0 .96h6.03a.48.48 0 0 0 .48-.48a.487.487 0 0 0-.48-.48m0-2.248h-6.03a.48.48 0 0 0 0 .96h6.03a.48.48 0 0 0 0-.96m0-2.256h-6.03a.48.48 0 0 0 0 .96h6.03a.48.48 0 0 0 .48-.48a.487.487 0 0 0-.48-.48M2.983 16.75c0 .265.215.48.48.48h5.806a.48.48 0 0 0 0-.96H3.463a.48.48 0 0 0-.48.48m4.957-1.781a.48.48 0 0 0 0-.96H4.792a.48.48 0 0 0 0 .96z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DoNotDisturb(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.29 20.328c0-.11-.089-.177-.266-.177H8.87v.376h.11c.199 0 .31-.089.31-.199m-.287.442H8.87v.442h.133c.155 0 .31-.066.31-.221c.022-.155-.133-.222-.31-.222zm-1.371-.619c-.089 0-.133 0-.155.022v.442h.133c.177 0 .287-.089.287-.221c0-.178-.088-.244-.265-.244zm-6.526 0c-.089 0-.133 0-.155.022v1.04h.133a.486.486 0 0 0 .532-.556v.003a.466.466 0 0 0-.511-.51h.002zm4.361-3.539c-.509 0-.8.465-.8 1.084s.31 1.062.818 1.062s.8-.465.8-1.084c-.026-.575-.314-1.062-.818-1.062"/><path fill="currentColor" d="M5.154 0C2.301 0 0 1.99 0 4.446v.022a1.57 1.57 0 0 0 3.14 0v-.01c0-.308-.089-.595-.244-.837l.004.006a2.474 2.474 0 0 1 4.686 1.082v.001a2.488 2.488 0 0 1-2.201 2.455l-.011.001c-.133.044-.929.133-1.15.177h-.022C1.792 7.741 0 9.555 0 11.745v8.206c0 2.234 1.593 4.048 3.562 4.048h3.185c1.969 0 3.562-1.814 3.562-4.048V4.445c0-2.456-2.301-4.446-5.154-4.446zM.598 10.927a8.6 8.6 0 0 1 1.449-.088h-.009a2.972 2.972 0 0 1 1.935.515l-.01-.007a2.057 2.057 0 0 1 .8 1.776v-.006a2.381 2.381 0 0 1-.773 1.9l-.002.002a3.266 3.266 0 0 1-2.18.597l.012.001c-.553 0-.929-.044-1.194-.066v-4.623zm.682 7.411v.974H.598v-3.231h.885l.686 1.172c.187.309.372.673.531 1.051l.021.057h.022c-.042-.389-.066-.84-.066-1.297v-.032v.002v-.929h.684v3.207h-.8l-.702-1.238a8.557 8.557 0 0 1-.556-1.07l-.023-.058h-.019c.019.398.019.862.019 1.394zm.442 2.964a1.182 1.182 0 0 1-.663.201l-.07-.002h.003c-.177 0-.31-.022-.42-.022v-1.57a3.19 3.19 0 0 1 .491-.023h-.005a.946.946 0 0 1 .666.178l-.003-.002a.632.632 0 0 1 .265.578v-.003a.837.837 0 0 1-.262.663zm.863.176h-.35v-1.593h.354zm.708.023h-.007c-.157 0-.307-.033-.442-.092l.007.003l.066-.287a.947.947 0 0 0 .4.089c.177 0 .24-.066.24-.177c0-.089-.066-.155-.266-.221c-.266-.089-.442-.24-.442-.465c0-.266.221-.486.598-.486h.015c.14 0 .273.033.391.091l-.005-.002l-.086.286a.925.925 0 0 0-.331-.066H3.43c-.155 0-.24.066-.24.155c0 .11.089.155.287.221c.287.11.42.24.42.465c.042.266-.157.486-.599.486zm2.013-1.305h-.439v1.28h-.354v-1.28h-.421v-.31h1.217zm1.482.598c0 .486-.24.73-.663.73c-.4 0-.64-.221-.64-.73v-.886h.355v.907c0 .266.11.42.287.42c.199 0 .287-.133.287-.42v-.907h.354v.885zm-1.371-1.44l-.06.001a1.515 1.515 0 0 1-1.51-1.643v.006a1.586 1.586 0 0 1 1.583-1.68h.032h-.002l.072-.002a1.526 1.526 0 0 1 1.523 1.622v-.004a1.568 1.568 0 0 1-1.557 1.705l-.084-.002h.004zm2.521 2.123a1.963 1.963 0 0 1-.108-.363l-.002-.013c-.044-.199-.11-.24-.266-.266h-.108v.619H7.1v-1.546c.114-.015.246-.024.38-.024l.112.002h-.005a.772.772 0 0 1 .511.135l-.003-.002a.4.4 0 0 1 .156.317v.016v-.001a.442.442 0 0 1-.284.399l-.003.001a.43.43 0 0 1 .22.284l.001.003c.043.178.088.326.142.47l-.009-.028h-.376zm1.593-.133a.992.992 0 0 1-.623.155h.004c-.177 0-.31-.022-.4-.022v-1.569a2.543 2.543 0 0 1 .47-.022h-.005a.893.893 0 0 1 .469.091l-.005-.002a.376.376 0 0 1 .199.309v.001a.366.366 0 0 1-.263.331l-.003.001c.178.037.31.193.31.379l-.001.022v-.001a.312.312 0 0 1-.149.329l-.001.001zm.199-4.645h-.905v2.588h-.752v-2.586h-.885v-.621h2.544zm-2.298-1.061h-.024A2.197 2.197 0 0 1 5.22 13.24l-.001.009a2.27 2.27 0 0 1 2.261-2.455h.018h-.001h.047a2.194 2.194 0 0 1 2.187 2.375l.001-.008a2.234 2.234 0 0 1-2.221 2.478l-.084-.002h.004z"/><path fill="currentColor" d="M7.476 11.635c-.708 0-1.128.686-1.128 1.593c0 .929.442 1.57 1.15 1.57s1.128-.686 1.128-1.593c-.022-.841-.42-1.57-1.15-1.57m-3.87 1.504a1.34 1.34 0 0 0-1.489-1.482l.006-.001l-.056-.001a2.09 2.09 0 0 0-.444.047l.014-.003v3.076a1.874 1.874 0 0 0 .38.022h-.004a1.475 1.475 0 0 0 1.592-1.667z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Docker(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.21 10.923V7.764h3.533v3.159zM18.743 0v3.267H15.21V0zm4.177 7.77v3.159h-3.533V7.77zm-8.353-3.855v3.213h-3.533V3.915zm4.177 0v3.213H15.21V3.915zm14.834 5.348l.696.482l-.482.914a4.111 4.111 0 0 1-2.434 2.081l-.029.008a7.562 7.562 0 0 1-2.419.321h.01c-1.36 3.409-3.692 6.19-6.652 8.072l-.068.041c-2.937 1.771-6.484 2.819-10.276 2.819c-.2 0-.398-.003-.597-.009l.029.001q-5.89 0-8.889-3.534a10.867 10.867 0 0 1-2.072-4.152l-.016-.076a10.955 10.955 0 0 1-.377-2.881c0-.551.04-1.094.117-1.624l-.007.06h23.294l.148.002a5.24 5.24 0 0 0 2.451-.605l-.029.014a4.454 4.454 0 0 1-.64-2.803l-.002.019a5.803 5.803 0 0 1 .871-2.809l-.014.025l.536-.75l.75.536a5.097 5.097 0 0 1 2.189 3.339l.005.031a8.809 8.809 0 0 1 2.208-.05l-.037-.003c.647.024 1.245.22 1.753.542l-.014-.009zM6.215 7.764v3.159H2.681V7.764zm4.177 0v3.159H6.86V7.764zm4.177 0v3.159h-3.535V7.764zm-4.178-3.855v3.213H6.86V3.909z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Doctor(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m6.53 8.098l.14-.012a.305.305 0 0 0 .141-.053l-.001.001c.134.462.298.948.503 1.457c.263.666.522 1.213.812 1.741l-.04-.08c-.024.364-.053.738-.091 1.1a2.565 2.565 0 0 1-.129.627l.005-.018c-.012.005-.029 2.08-.029 2.08a2.868 2.868 0 0 0 2.198 2.787l.02.004a.384.384 0 0 1 .357-.246h.574a.39.39 0 0 1 .356.243l.001.003a2.877 2.877 0 0 0 2.229-2.789v-.001s-.035-2.066-.053-2.08a3.152 3.152 0 0 1-.122-.593l-.001-.015c-.035-.364-.058-.729-.091-1.1c.247-.446.506-.992.734-1.555l.038-.106c.205-.509.364-.994.503-1.457a.297.297 0 0 0 .139.053h.001l.141.012c.17.018.32-.122.334-.339l.152-1.931v-.002a.32.32 0 0 0-.279-.317h-.019a7.143 7.143 0 0 0-.242-2.999l.013.051A4.27 4.27 0 0 0 11.725.122l-.026-.004a5.847 5.847 0 0 0-.993-.112h-.021a5.49 5.49 0 0 0-1.038.118l.036-.006a4.295 4.295 0 0 0-3.114 2.419l-.011.027a7.047 7.047 0 0 0-.225 2.985l-.004-.037a.316.316 0 0 0-.282.313v.007l.152 1.931c.014.222.166.356.33.338z"/><path fill="currentColor" d="M21.416 20.878c-.07-3.04-.374-3.728-.538-4.194c-.065-.187-.118-1.451-2.206-2.271c-1.28-.504-2.932-.514-4.33-1.105v1.644a3.632 3.632 0 0 1-2.944 3.56l-.023.004a.384.384 0 0 1-.374.32h-.018v1.24a2.194 2.194 0 1 0 4.388 0v-.866a1.248 1.248 0 1 1 .588-.055l-.009.003v.965a2.774 2.774 0 0 1-5.548 0v-.05v.002v-1.251a.381.381 0 0 1-.35-.318v-.002a3.635 3.635 0 0 1-2.954-3.556v-1.657c-1.404.603-3.066.615-4.353 1.12c-2.094.819-2.142 2.08-2.206 2.27c-.16.468-.468 1.153-.538 4.195c-.012.4 0 1.013 1.206 1.549c2.626 1.03 6.009 1.35 9.344 1.58h.32c3.342-.228 6.72-.547 9.344-1.58c1.201-.533 1.212-1.142 1.201-1.546zm-14.681-1.24H5.489v1.251h-.89v-1.247H3.353v-.89h1.246v-1.246h.89v1.246h1.246z"/><path fill="currentColor" d="M16.225 17.965v-.001a.673.673 0 1 0 0 .001"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dollar(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12.496 15.87l.001.092a5.08 5.08 0 0 1-1.337 3.441l.003-.004a5.744 5.744 0 0 1-3.432 1.824l-.031.004v2.36a.413.413 0 0 1-.413.413H7.27h.001h-1.807a.434.434 0 0 1-.429-.429v-2.344a8.42 8.42 0 0 1-1.766-.434l.059.019a8.218 8.218 0 0 1-2.37-1.254l.02.015a7.476 7.476 0 0 1-.629-.509l.007.006q-.167-.16-.234-.24a.407.407 0 0 1-.025-.55l-.001.001l1.379-1.808a.415.415 0 0 1 .307-.16h.001a.331.331 0 0 1 .32.121l.001.001l.026.026a6.968 6.968 0 0 0 3.208 1.666l.046.009c.298.068.64.107.991.107l.062.001c.691 0 1.331-.216 1.858-.584l-.011.007a1.877 1.877 0 0 0 .824-1.637v.004c0-.263-.075-.508-.204-.715l.003.006a2.543 2.543 0 0 0-.447-.559l-.001-.001a3.693 3.693 0 0 0-.761-.493l-.022-.01q-.536-.274-.88-.429t-1.071-.435q-.522-.214-.824-.335t-.824-.355t-.837-.415t-.757-.475a4.765 4.765 0 0 1-.721-.57l.001.001a8.263 8.263 0 0 1-.57-.64l-.013-.016a3.267 3.267 0 0 1-.467-.756l-.008-.021a5.189 5.189 0 0 1-.275-.854l-.007-.036a4.549 4.549 0 0 1-.114-1.022V7.85c0-1.25.501-2.384 1.314-3.21l-.001.001a6.116 6.116 0 0 1 3.379-1.789l.036-.005V.428c0-.117.049-.223.127-.298a.413.413 0 0 1 .298-.127h1.827c.228 0 .413.185.413.413v.017v-.001v2.357c.55.058 1.052.167 1.533.323l-.053-.015c.459.14.845.293 1.216.471l-.052-.022c.328.163.607.33.871.516l-.02-.014q.4.282.522.388t.201.187a.385.385 0 0 1 .066.51l.001-.001l-1.082 1.956a.364.364 0 0 1-.307.214h-.001a.422.422 0 0 1-.362-.094h.001q-.04-.04-.194-.16t-.522-.355a6.665 6.665 0 0 0-.741-.411l-.042-.018a5.63 5.63 0 0 0-.956-.339l-.042-.009a4.34 4.34 0 0 0-1.144-.154h-.001a3.462 3.462 0 0 0-2.088.584l.013-.008a1.801 1.801 0 0 0-.686 2.142l-.004-.012c.092.22.226.406.394.555l.002.001c.161.151.333.296.513.431l.016.011c.215.148.462.287.721.403l.03.012q.502.234.81.362t.938.368q.71.268 1.085.422t1.018.469a8.04 8.04 0 0 1 1.041.587l-.029-.018c.31.22.58.44.837.675l-.006-.006c.277.242.513.522.701.835l.009.016c.166.293.309.634.413.991l.009.034c.109.365.171.785.171 1.219v.042v-.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Download(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 17.143V24h24v-6.857zm5.143 5.143H1.714v-1.714h3.429zM6.857 8.571h3.429V0h3.429v8.571h3.429l-5.143 5.143z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dribbble(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 21.44c-.541-2.962-1.298-5.573-2.287-8.067l.1.286h-.031l-.031.016q-.25.094-.672.258t-1.578.766c-.836.434-1.537.859-2.207 1.324l.066-.044a13.662 13.662 0 0 0-2.04 1.781l-.007.007a8.4 8.4 0 0 0-1.588 2.262l-.022.05l-.234-.172a10.001 10.001 0 0 0 6.454 2.344h.081H12h.005a10.2 10.2 0 0 0 4.062-.837l-.066.026zm-2.89-9.484q-.328-.766-.828-1.734c-3.057.923-6.57 1.454-10.207 1.454l-.325-.001h.017c-.01.086-.016.186-.016.287v.043v-.002v.042c0 1.315.252 2.57.711 3.721l-.024-.068a10.301 10.301 0 0 0 1.946 3.158l-.008-.009A13.2 13.2 0 0 1 6.31 16.24l-.005.005a14.014 14.014 0 0 1 2.183-1.916l.043-.029q1.078-.734 2.039-1.266c.432-.254.944-.506 1.474-.721l.081-.029l.578-.203l.203-.054c.079-.021.147-.046.212-.077l-.008.003zM11.44 8.64a53.493 53.493 0 0 0-3.913-6.047l.1.14c-2.866 1.381-4.962 3.97-5.644 7.088l-.012.068c3.37-.007 6.63-.463 9.729-1.31zm10.687 4.984a14.429 14.429 0 0 0-4.214-.614c-.768 0-1.522.059-2.259.172l.082-.01c.79 2.06 1.482 4.519 1.958 7.055l.042.273a10.33 10.33 0 0 0 2.866-2.92l.024-.04a9.975 9.975 0 0 0 1.49-3.856l.008-.058zM9.549 2.046a.043.043 0 0 0-.031.016a.044.044 0 0 1 .029-.014zm9.218 2.265a9.834 9.834 0 0 0-6.635-2.561l-.137.001h.007c-.861.003-1.695.111-2.493.312l.071-.015a43.86 43.86 0 0 1 3.723 5.735l.117.234c.794-.3 1.464-.618 2.102-.983l-.071.037a13.452 13.452 0 0 0 1.545-.986l-.037.026c.376-.287.71-.578 1.023-.891c.206-.197.398-.403.576-.621l.01-.012zm3.484 7.579a9.983 9.983 0 0 0-2.341-6.422l.013.016l-.016.016q-.141.187-.297.383a8.8 8.8 0 0 1-.674.689l-.006.006c-.339.324-.7.633-1.078.923l-.031.023a17.043 17.043 0 0 1-3.503 1.981l-.114.042q.39.828.687 1.484l.102.266c.036.104.076.191.121.275l-.005-.009q.56-.08 1.164-.11t1.149-.031q.547 0 1.078.023t1 .062t.88.086t.75.102l.57.094q.234.039.39.07zM24 12v.09c0 2.187-.598 4.234-1.64 5.987l.03-.054a12.045 12.045 0 0 1-4.311 4.337l-.056.03c-1.729 1.012-3.806 1.609-6.024 1.609s-4.295-.597-6.081-1.64l.057.031a12.045 12.045 0 0 1-4.337-4.311l-.03-.056C.596 16.294-.001 14.217-.001 11.999s.597-4.295 1.64-6.081l-.031.057a12.045 12.045 0 0 1 4.311-4.337l.056-.03C7.704.596 9.781-.001 11.999-.001s4.295.597 6.081 1.64l-.057-.031a12.045 12.045 0 0 1 4.337 4.311l.03.056A11.606 11.606 0 0 1 24 11.908v.096v-.005z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Dropbox(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m5.247 9.143l7.669 4.735l-5.309 4.424L0 13.35zm15.307 8.617v1.68l-7.606 4.549V24l-.015-.015l-.015.015v-.015l-7.591-4.549v-1.68l2.282 1.49l5.309-4.409v-.031l.015.015l.015-.015v.031l5.325 4.409zM7.606 0l5.309 4.424l-7.669 4.72L0 4.952zm12.979 9.143l5.247 4.207l-7.591 4.952l-5.325-4.424zM18.24 0l7.591 4.952l-5.247 4.191l-7.669-4.72z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func DrugPack(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.869.31L.311 7.869a1.049 1.049 0 0 0 0 1.49l14.33 14.331a1.046 1.046 0 0 0 1.488 0l7.559-7.559a1.049 1.049 0 0 0 0-1.49L9.358.311a1.049 1.049 0 0 0-1.49 0zm8.665 20.297a1.453 1.453 0 1 1-2.048-2.06a1.453 1.453 0 0 1 2.048 2.059zm-2.771-2.764a1.451 1.451 0 1 1-2.041-2.06a1.451 1.451 0 0 1 2.04 2.06zm-2.77-2.77a1.453 1.453 0 1 1-2.048-2.06a1.453 1.453 0 0 1 2.048 2.059zm-2.771-2.771a1.45 1.45 0 1 1 .431-1.031v.001a1.455 1.455 0 0 1-.43 1.03zm-2.764-2.77a1.45 1.45 0 1 1 .431-1.031v.001a1.455 1.455 0 0 1-.43 1.03zm15.145 7.005a1.45 1.45 0 1 1 .43-1.031v.002a1.455 1.455 0 0 1-.43 1.03zm-2.77-2.771a1.45 1.45 0 1 1 .43-1.031v.002a1.455 1.455 0 0 1-.43 1.03zm-2.771-2.77a1.45 1.45 0 1 1 .43-1.031v.002a1.455 1.455 0 0 1-.43 1.03zm-2.764-2.764a1.45 1.45 0 1 1 .431-1.031v.001a1.455 1.455 0 0 1-.43 1.03zm-2.77-2.77a1.45 1.45 0 1 1 .431-1.031v.001a1.455 1.455 0 0 1-.43 1.03z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Earth(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 0C5.373 0 0 5.373 0 12s5.373 12 12 12s12-5.373 12-12C23.992 5.376 18.624.008 12.001 0zm-.495 15.689c.068.203-.525.795-.575.829s-1.15-.575-1.337-.66a5.863 5.863 0 0 1-.599-1.413l-.01-.042l-.406-.389s-.93-.694-1.015-.727s-.254-1.336-.308-1.455a2.853 2.853 0 0 0-.516-.556l-.005-.004a5.292 5.292 0 0 1-.152-1.215v-.003c.043-.431.182-.822.395-1.162l-.007.011a2.009 2.009 0 0 0 .217-1.519l.003.014c-.102-.169-.765.08-.765.08a2.369 2.369 0 0 0-.543-.217l-.017-.004l-.343-.64s-2.044-.146-2.01-.196c1.947-2.732 5.103-4.492 8.671-4.492c.473 0 .939.031 1.396.091l-.054-.006h.11c.366.027.704.027.896.367c.038.066.724.669.707.669h-1.874l-.389.338l1.71 1.022c-.234.064-.71.212-1.03.314a1.977 1.977 0 0 0-.601.31l.005-.004l-.676.507l-.629.899l.916.677s1.066.812 1.12.846s.27-.677.27-.677q-.037-.822-.07-1.646l.507.169l.744-.068l.829.64c.143.224.285.484.407.756l.016.04c-.01.114-.09.735.054.758h.006a.148.148 0 0 0 .075-.021h-.001c.257-.125.501-.27.746-.415c.116-.069.364-.289.505-.289c-.061.25-.148.47-.261.674l.007-.014c-.026.04-.66 1.095-.93 1.543c-.085.14-.179.261-.287.369l-.88.896l.305.524l-.736.871a2.024 2.024 0 0 1-.199.204l-.002.002l-.789.699s-.862.203-.914.203s-.694.152-.694.152s-1.218-.118-1.387.524a3.562 3.562 0 0 0 .398 1.856l-.009-.02zm8.986-3.138c.154-.034-.41.72-.445.838a5.717 5.717 0 0 1-.615.72l-.96.12a2.011 2.011 0 0 1 1.797-1.651l.008-.001c.105-.01.184-.019.214-.026zm-5.445 3.005l-.449-.664c-.036-.054.179-.287.179-.287s1.705-.251 1.866-.234s.43.521.43.521s-.592.771-.79.771s-1.238-.107-1.238-.107zm.454 2.328s-.205.086-.41.188s-.462.188-.496.12a2.328 2.328 0 0 0-.452-.37l-.01-.006a4.903 4.903 0 0 0-.791-.467l-.03-.013a2.084 2.084 0 0 1-.681-.56l-.003-.004c-.044-.099.171-.393.205-.547s-.017-.613.086-.56s.325-.08.547 0c.25.079.463.214.632.391l.001.001c.103.224.247.412.425.562l.002.002c.188.12.53.256.513.48s.257.513.257.513zm.334 3.2c-.931-.413-1.164-.88-1.093-.969c.3-.266.575-.535.836-.818l.007-.008c.179-.251 1.113-1.31 1.113-1.31l-.287-.574l.8-.474c.237-.142.516-.238.816-.268l.008-.001l.78-.076q.07-.007.139-.009l1.258-.034s1.143-.296.969.018c-2.618 4.73-5.346 4.521-5.346 4.521z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Edge(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 10.659h.014a14.66 14.66 0 0 1 .82-3.34l-.033.101c.414-1.11.933-2.071 1.565-2.945l-.025.036a11.493 11.493 0 0 1 2.222-2.342l.024-.018A9.827 9.827 0 0 1 7.514.596l.07-.02a11.372 11.372 0 0 1 3.615-.577h.097h-.005l.175-.001c1.976 0 3.829.525 5.428 1.444l-.053-.028a10.588 10.588 0 0 1 3.914 4.015l.027.054a11.843 11.843 0 0 1 1.394 5.942v-.016v2.521H7.093v.068c0 .931.268 1.8.732 2.533l-.012-.019a4.523 4.523 0 0 0 1.804 1.631l.026.012c.738.36 1.593.628 2.494.758l.047.006c.473.075 1.019.118 1.574.118c.452 0 .897-.028 1.334-.084l-.052.005a14.196 14.196 0 0 0 2.889-.653l-.1.03a9.04 9.04 0 0 0 2.354-1.155l-.03.02v5.054a11.9 11.9 0 0 1-2.992 1.215l-.084.018a16.147 16.147 0 0 1-4.083.512h-.113h.006c-.071.002-.155.002-.24.002a11.55 11.55 0 0 1-4.08-.739l.08.026a9.01 9.01 0 0 1-4.157-3.307l-.02-.031a8.684 8.684 0 0 1-1.668-4.983v-.008a9.427 9.427 0 0 1 1.51-5.557l-.022.037a10.387 10.387 0 0 1 4.294-3.568l.066-.026a8.056 8.056 0 0 0-1.025 1.631l-.021.049a8.354 8.354 0 0 0-.61 2.089l-.006.049h8.514a5.365 5.365 0 0 0-.115-1.914l.008.037a3.538 3.538 0 0 0-.637-1.369l.006.009a4.735 4.735 0 0 0-.931-.882l-.014-.01a3.893 3.893 0 0 0-1.052-.542l-.028-.008q-.55-.174-1.006-.275a4.356 4.356 0 0 0-.737-.113l-.014-.001l-.295-.014a13.022 13.022 0 0 0-3.572.623l.092-.026a12.05 12.05 0 0 0-3.043 1.429l.046-.028c-.885.58-1.657 1.2-2.364 1.887l.004-.003a15.232 15.232 0 0 0-1.816 2.145l-.032.049z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eject(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.546 24.001A.546.546 0 0 1 0 23.455v-1.091c0-.301.245-.545.546-.545h22.909c.301 0 .545.244.546.545v1.091a.546.546 0 0 1-.546.546zm.063-4.364a.675.675 0 0 1-.499-1.052l-.002.002L11.498.291a.574.574 0 0 1 1-.003l.001.003l11.39 18.297a.675.675 0 0 1-.497 1.049h-.003z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Electronjs(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.827 5.156a10.076 10.076 0 0 0-4.185.04l.067-.012a3.749 3.749 0 0 0-2.471 1.702l-.009.016a3.322 3.322 0 0 0-.42 2.214l-.002-.019c.14.968.462 1.837.931 2.608l-.017-.03c.032.05.073.091.121.121l.002.001a.336.336 0 0 0 .162.041h.016h-.001a.292.292 0 0 0 .205-.095a.282.282 0 0 0 .095-.204v-.005a.31.31 0 0 0-.014-.092l.001.002a.205.205 0 0 1-.014-.068a6.136 6.136 0 0 1-.841-2.255l-.005-.036a2.834 2.834 0 0 1 .337-1.868l-.007.014A3.303 3.303 0 0 1 3.95 5.788l.02-.003a9.14 9.14 0 0 1 3.782.011l-.06-.011h.004c.014 0 .027.005.037.014l.014.014h.006a.314.314 0 0 0 .226-.095a.319.319 0 0 0 .095-.228v-.021a.33.33 0 0 0-.069-.202l.001.001a.334.334 0 0 0-.175-.108h-.002zm-4.473 8.727a22.426 22.426 0 0 0 2.166 1.999l.042.033c.726.587 1.536 1.15 2.386 1.653l.094.052a22.04 22.04 0 0 0 3.335 1.557l.159.052c.946.362 2.061.661 3.214.846l.099.013c.491.083 1.056.13 1.632.13c.44 0 .873-.028 1.298-.081l-.051.005a5.18 5.18 0 0 0 2.229-.79l-.021.012a.647.647 0 0 0 .094-.12l.002-.003a.28.28 0 0 0 .041-.147v-.006a.319.319 0 0 0-.095-.228a.319.319 0 0 0-.228-.095h-.004a.2.2 0 0 0-.081.028l.001-.001a.307.307 0 0 0-.08.055a4.61 4.61 0 0 1-2.032.694l-.019.001a9.857 9.857 0 0 1-2.758-.088l.059.008a17.13 17.13 0 0 1-3.227-.857l.118.039a18.642 18.642 0 0 1-3.329-1.588l.084.047a20.285 20.285 0 0 1-2.537-1.712l.042.032a18.583 18.583 0 0 1-2.13-1.966l-.011-.012a.358.358 0 0 0-.12-.079l-.002-.001a.32.32 0 0 0-.122-.027h-.005a.319.319 0 0 0-.228.095a.319.319 0 0 0-.095.228v.004c0 .04.01.077.028.11l-.001-.001zm15.792-.409a9.436 9.436 0 0 0 2.027-3.465l.019-.067a3.69 3.69 0 0 0 .205-1.226c0-.655-.167-1.271-.46-1.808l.01.02a3.453 3.453 0 0 0-1.641-1.437l-.023-.008a6.738 6.738 0 0 0-2.615-.518h-.032h.002h-.011a.272.272 0 0 0-.207.095a.342.342 0 0 0-.08.221v.012v-.001v.011c0 .084.03.161.081.221v-.001a.27.27 0 0 0 .207.095h.012h-.001c.05-.002.11-.002.169-.002c.784 0 1.533.156 2.215.439l-.038-.014a3.052 3.052 0 0 1 1.411 1.174l.007.012a3.32 3.32 0 0 1 .143 2.641l.007-.023a8.925 8.925 0 0 1-1.874 3.223l.006-.007l-.054.109a.231.231 0 0 0-.027.109v.004c0 .089.036.17.095.228a.319.319 0 0 0 .228.095h.012a.222.222 0 0 0 .129-.041h-.001a.692.692 0 0 1 .105-.066l.004-.002zm-5.264-8.208a20.854 20.854 0 0 0-3.194 1.019l.14-.051c-1.082.44-1.988.893-2.854 1.405l.1-.055a20.554 20.554 0 0 0-3.176 2.245l.027-.022A16.833 16.833 0 0 0 2.58 12.25l-.028.037a9.701 9.701 0 0 0-1.409 2.456l-.023.066a4.639 4.639 0 0 0-.339 1.755c0 .209.014.415.04.616l-.003-.024c.019.075.057.14.109.191c.053.05.124.08.203.08h.017h-.001h.006a.32.32 0 0 0 .227-.094a.319.319 0 0 0 .095-.228v-.058a4.259 4.259 0 0 1 .316-2.142l-.011.028a9.027 9.027 0 0 1 1.373-2.363l-.013.017a16.422 16.422 0 0 1 2.266-2.325l.025-.02a20.79 20.79 0 0 1 2.955-2.086l.099-.054c.714-.43 1.556-.855 2.429-1.22l.131-.048a18.185 18.185 0 0 1 2.702-.886l.134-.028a.665.665 0 0 0 .3-.139l-.001.001a.294.294 0 0 0 .08-.203v-.017v.001a.292.292 0 0 0-.095-.205a.282.282 0 0 0-.204-.095h-.08zM6.273 19.142a9.902 9.902 0 0 0 2.08 3.567l-.007-.008a3.747 3.747 0 0 0 2.722 1.295h.006a3.47 3.47 0 0 0 2.081-.715l-.008.006a6.584 6.584 0 0 0 1.729-1.954l.017-.032a.624.624 0 0 0 .039-.09l.002-.005a.288.288 0 0 0 .014-.09v-.01a.319.319 0 0 0-.095-.228a.319.319 0 0 0-.228-.095h-.006a.284.284 0 0 0-.149.042l.001-.001a.388.388 0 0 0-.121.121l-.001.002a6.159 6.159 0 0 1-1.528 1.777l-.013.01a2.862 2.862 0 0 1-1.73.614h-.002a3.21 3.21 0 0 1-2.328-1.167l-.004-.006a9.11 9.11 0 0 1-1.822-3.181l-.018-.065a.74.74 0 0 0-.135-.149l-.001-.001a.289.289 0 0 0-.186-.068h-.005a.292.292 0 0 0-.205.095a.319.319 0 0 0-.095.228v.113zm9.818-.572c.252-.779.471-1.718.615-2.681l.012-.1c.139-.9.218-1.938.218-2.994v-.035v.002v-.236c0-1.318-.119-2.608-.347-3.859l.02.131a16.37 16.37 0 0 0-.967-3.396l.04.11a10.01 10.01 0 0 0-1.448-2.53l.016.021a4.956 4.956 0 0 0-1.784-1.46l-.029-.013a.092.092 0 0 0-.066-.027h-.081a.34.34 0 0 0-.221.081h.001a.27.27 0 0 0-.095.207v.012v-.001v.008c0 .068.02.132.055.185l-.001-.001c.032.05.079.088.134.108l.002.001a4.49 4.49 0 0 1 1.642 1.365l.009.012c.516.68.956 1.458 1.286 2.293l.023.066c.362.89.665 1.938.858 3.023l.014.097c.2 1.072.314 2.305.314 3.565v.08v-.004v.18a19.85 19.85 0 0 1-.232 3.036l.014-.11a17.98 17.98 0 0 1-.69 2.936l.036-.127l.027-.08v.031c0 .089.036.17.095.228a.319.319 0 0 0 .228.095h.009a.294.294 0 0 0 .187-.068a.344.344 0 0 0 .108-.142l.001-.002zm6.055-.328v-.02c0-.415-.172-.79-.449-1.058a1.497 1.497 0 0 0-1.072-.45h-.035h.002h-.019c-.416 0-.791.173-1.058.45c-.277.268-.45.643-.45 1.058v.021v-.001v.032c0 .42.172.8.45 1.072c.268.278.643.451 1.059.451h.02h-.001h.033c.42 0 .799-.172 1.072-.45a1.5 1.5 0 0 0 .449-1.071v-.037v.002zm-.655 0v.02a.868.868 0 0 1-.259.62a.868.868 0 0 1-.62.259h-.021h.001h-.003a.879.879 0 0 1-.624-.259a.848.848 0 0 1-.273-.624v-.017v.001a.863.863 0 0 1 .272-.627a.86.86 0 0 1 .627-.273h.016c.246 0 .468.105.623.272l.001.001c.16.16.259.381.259.625v.002zM1.554 19.801h.019c.415 0 .79-.172 1.058-.449a1.5 1.5 0 0 0 .451-1.073v-.051c0-.416-.173-.791-.45-1.058a1.468 1.468 0 0 0-1.058-.45h-.021h.001h-.032c-.42 0-.8.172-1.072.45c-.277.268-.45.643-.45 1.058v.021v-.001v.032c0 .42.172.8.45 1.072a1.5 1.5 0 0 0 1.071.449h.035zm0-.655h-.02a.868.868 0 0 1-.62-.259a.868.868 0 0 1-.259-.62v-.023c0-.244.099-.464.259-.624a.848.848 0 0 1 .624-.273h.017h-.001c.247.001.47.105.627.272a.86.86 0 0 1 .273.627v.016a.848.848 0 0 1-.272.623l-.001.001a.88.88 0 0 1-.621.255h-.006zm9.52-16.064h.033c.42 0 .799-.172 1.072-.45c.278-.268.451-.643.451-1.059v-.052c0-.42-.172-.799-.45-1.072a1.497 1.497 0 0 0-1.072-.45h-.035h.002h-.019c-.416 0-.791.173-1.058.45c-.278.273-.45.652-.45 1.072v.035v-.002v.019c0 .416.173.791.45 1.058c.268.277.643.45 1.058.45zh-.001zm0-.654h-.007a.836.836 0 0 1-.606-.259a.836.836 0 0 1-.259-.606v-.028c0-.243.099-.462.259-.62a.836.836 0 0 1 .606-.259h.028c.243 0 .462.099.62.259c.16.158.259.377.259.62v.021v-.001v.007a.836.836 0 0 1-.259.606a.867.867 0 0 1-.621.26zm.246 11.264a.27.27 0 0 1-.118.027h-.147c-.303 0-.577-.125-.773-.327a1.074 1.074 0 0 1-.327-.774v-.041c0-.298.126-.566.327-.754l.001-.001c.198-.194.47-.314.77-.314h.022h-.001h.021c.262 0 .502.093.69.247l-.002-.001c.188.148.323.355.38.593l.001.007v.011c0 .044.005.087.014.129l-.001-.004c.009.033.014.07.014.109v.021c0 .262-.093.502-.247.69l.001-.002c-.155.192-.371.33-.619.382l-.007.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ellipse(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24 12c0 6.627-5.373 12-12 12S0 18.627 0 12S5.373 0 12 0s12 5.373 12 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Email(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16.484 11.976l6.151-5.344v10.627zm-7.926.905l2.16 1.875c.339.288.781.462 1.264.462h.017h-.001h.014c.484 0 .926-.175 1.269-.465l-.003.002l2.16-1.875l6.566 5.639H1.995zM1.986 5.365h20.03l-9.621 8.356a.612.612 0 0 1-.38.132h-.014h.001h-.014a.61.61 0 0 1-.381-.133l.001.001zm-.621 1.266l6.15 5.344l-6.15 5.28zm21.6-2.441c-.24-.12-.522-.19-.821-.19H1.859a1.87 1.87 0 0 0-.835.197l.011-.005A1.856 1.856 0 0 0 0 5.855v12.172a1.86 1.86 0 0 0 1.858 1.858h20.283a1.86 1.86 0 0 0 1.858-1.858V5.855c0-.727-.419-1.357-1.029-1.66l-.011-.005z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ember(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23.995 13.83v-.007a.277.277 0 0 0-.402-.247l.002-.001a1.878 1.878 0 0 1-1.095.4h-.003c-.514-.049-.353-1.2-.353-1.2s.113-1.054-.191-1.14s-.675.274-.675.274a3.74 3.74 0 0 0-.678 1.144l-.009.026l-.06.018a8.173 8.173 0 0 0-.014-1.443l.003.033c-.06-.131-.615-.12-.705.113a13.49 13.49 0 0 0-.552 2.448l-.008.072a4.185 4.185 0 0 1-1.594.848l-.03.007c-.758.12-.938-.353-.938-.353s2.055-.574 1.984-2.216s-1.658-1.035-1.84-.9a3.205 3.205 0 0 0-1.371 2.224l-.002.016c-.007.053-.026.282-.026.282a7.007 7.007 0 0 1-1.191.657l-.047.018s1.238-2.085-.274-3.034c-.686-.413-1.23.454-1.23.454s2.044-2.276 1.594-4.2c-.218-.915-.675-1.016-1.095-.866a1.922 1.922 0 0 0-.877.622l-.003.004a6.952 6.952 0 0 0-1.013 2.95l-.003.031c-.191 1.781-.473 3.941-.473 3.941a1.49 1.49 0 0 1-.748.398l-.01.002c-.364.018-.202-1.076-.202-1.076s.282-1.673.262-1.954s-.042-.435-.371-.533s-.694.32-.694.32s-.96 1.451-1.04 1.673l-.049.09l-.049-.06s.675-1.976.03-2.006s-1.069.705-1.069.705s-.735 1.23-.765 1.369l-.049-.06a9.314 9.314 0 0 0 .239-1.766l.001-.019a.32.32 0 0 0-.396-.282h.002s-.424-.049-.533.222c-.221.734-.416 1.63-.548 2.547l-.012.104a4.239 4.239 0 0 1-1.728.761l-.027.004c-.698.007-.626-.442-.626-.442s2.55-.874 1.852-2.595a1.271 1.271 0 0 0-1.194-.574h.005a1.8 1.8 0 0 0-1.545 1.236l-.004.013a3.767 3.767 0 0 0-.292 1.166l-.001.015a1.003 1.003 0 0 1-.71-.112l.005.003c-.24-.202-.375 0-.375 0s-.42.522 0 .682c.305.106.662.186 1.031.227l.023.002c.107.461.37.847.73 1.11l.005.004c.758.574 2.205-.049 2.205-.049l.596-.33s.018.547.454.626s.615.038 1.369-1.796c.442-.938.473-.885.473-.885l.049-.011a9.166 9.166 0 0 0-.209 2.253v-.013c.131.48.705.435.705.435s.311.09.56-.795c.228-.735.477-1.355.77-1.949l-.035.078h.054a6.323 6.323 0 0 0 .122 2.433l-.009-.043c.32.585 1.158.198 1.158.198a5.79 5.79 0 0 0 .697-.396l-.021.013a2.354 2.354 0 0 0 1.683.479l-.01.001a8.677 8.677 0 0 0 3.011-.994l-.045.023a1.8 1.8 0 0 0 1.537 1.001h.004a2.586 2.586 0 0 0 2.056-.699l-.001.001a.769.769 0 0 0 .449.696l.005.002c.465.191.776-.855.776-.855l.776-2.145h.071s.042 1.398.806 1.62a2.518 2.518 0 0 0 1.766-.518l-.006.004a.533.533 0 0 0 .198-.543l.001.004zm-21.669.198c.03-1.2.818-1.721 1.087-1.462s.173.825-.342 1.178s-.746.286-.746.286zm10.229-4.646s.72-1.864.885-.96s-1.5 3.607-1.5 3.607c.169-1.032.383-1.924.653-2.791zm.851 5.19a1.189 1.189 0 0 1-1.631.732l.007.003a3.48 3.48 0 0 1 .248-1.703l-.009.023c.371-1.241 1.249-.758 1.249-.758s.607.466.135 1.702zm3.172-.544a1.678 1.678 0 0 1 .307-1.152l-.003.005c.413-.758.735-.342.735-.342a.746.746 0 0 1-.049.961a1.192 1.192 0 0 1-.989.533z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Envato(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.781.274c-.88-.471-2.922-.327-5.324.383c-1.825 1.472-7.76 6.766-7.835 13.348a.49.49 0 0 1-.94.195l-.001-.003c-.566-1.316-1.026-3.512-.405-6.888a.492.492 0 0 0-.863-.401l-.001.001q-.324.4-.614.813c-3.787 5.469-1.079 12.48 3.378 14.914s11.081 1.894 14.272-3.943s.555-17.226-1.667-18.419"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Equalizer(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.5 24v-7.845c-1.306-.344-2.253-1.514-2.253-2.905s.947-2.562 2.232-2.901l.021-.005V0H24v10.344a3 3 0 0 1 .021 5.807l-.021.005v7.845zm-9.75 0V11.155c-1.306-.344-2.253-1.514-2.253-2.905s.947-2.562 2.232-2.901l.021-.005V0h1.5v5.344c1.306.344 2.253 1.514 2.253 2.905s-.947 2.562-2.232 2.901l-.021.005V24zm-10.5 0v-3.845C.944 19.811-.003 18.641-.003 17.25s.947-2.562 2.232-2.901l.021-.005V0h1.5v14.344c1.306.344 2.253 1.514 2.253 2.905s-.947 2.562-2.232 2.901l-.021.005V24z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eraser(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M25.199 21.599h-9.103l9.952-9.952a1.196 1.196 0 0 0 0-1.696l-9.6-9.6a1.196 1.196 0 0 0-1.696 0L.351 14.75a1.196 1.196 0 0 0 0 1.696l7.2 7.203c.217.217.517.351.848.351h16.834a1.2 1.2 0 1 0-.035-2.399h.002zM15.6 2.896l7.903 7.903l-6.943 6.943l-7.903-7.903z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Euro(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16.64 20.097l.597 2.71a.493.493 0 0 1-.053.385l.001-.002a.543.543 0 0 1-.286.246l-.004.001l-.086.017a.79.79 0 0 1-.174.059l-.005.001q-.11.026-.273.08t-.366.094q-.205.042-.434.086t-.511.086t-.571.08t-.622.051q-.333.017-.656.017a11.29 11.29 0 0 1-7.002-2.246l.03.021a11.042 11.042 0 0 1-4.039-5.914l-.018-.077H.549a.553.553 0 0 1-.546-.545V13.32a.553.553 0 0 1 .545-.546h1.125q-.034-.971.017-1.79H.527a.525.525 0 0 1-.525-.525v-.022v.001v-1.964c0-.29.235-.525.525-.525h.022h-.001h1.67a11.162 11.162 0 0 1 4.118-5.738l.033-.022A11.255 11.255 0 0 1 13.199 0h-.007h.066c1.151 0 2.268.143 3.335.412l-.094-.02c.142.046.26.136.339.254l.001.002a.557.557 0 0 1 .05.413l.001-.004l-.733 2.71a.5.5 0 0 1-.238.331l-.002.001a.485.485 0 0 1-.412.041l.003.001l-.068-.017q-.068-.017-.196-.042l-.298-.06l-.383-.06l-.443-.051l-.494-.042l-.503-.017l-.1-.001c-1.393 0-2.69.407-3.78 1.109l.028-.017A6.781 6.781 0 0 0 6.728 7.9l-.017.043h7.978a.56.56 0 0 1 .546.651v-.003l-.409 1.943a.5.5 0 0 1-.548.443h.002h-8.32a15.307 15.307 0 0 0 .002 1.832l-.002-.043h7.831c.17 0 .321.08.419.204l.001.001a.553.553 0 0 1 .102.464l.001-.004l-.409 1.909a.552.552 0 0 1-.527.443H6.784c1.036 2.558 3.5 4.33 6.378 4.33h.069h-.003q.307 0 .614-.026t.571-.06t.503-.08t.418-.086l.315-.08l.205-.051l.086-.034a.505.505 0 0 1 .445.036l-.002-.001c.134.077.23.208.258.362z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Export(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m26.444 16.617l.03-.039a1.225 1.225 0 0 0 .062-.094c.009-.016.018-.028.026-.044l.024-.05c.008-.017.015-.031.021-.047l.018-.05l.018-.051l.013-.05l.013-.054c.004-.018.006-.038.009-.057l.007-.048c0-.035.005-.071.005-.106q0-.053-.005-.106c0-.016-.004-.033-.007-.048s-.005-.039-.009-.057l-.013-.054l-.013-.049l-.018-.052c-.006-.018-.011-.033-.018-.049l-.021-.047a.557.557 0 0 0-.024-.05l-.026-.044l-.029-.048l-.034-.046c-.012-.015-.019-.027-.03-.039q-.033-.039-.069-.076l-2.733-2.732a1.091 1.091 0 1 0-1.542 1.541l.001.001l.865.865h-3.328V9.092a1.08 1.08 0 0 0-.249-.692l.001.002l-.004-.006a1.03 1.03 0 0 0-.063-.069l-.009-.009l-.028-.029L11.319.323A1.03 1.03 0 0 0 11.25.26l-.023-.018l-.053-.04l-.027-.019c-.019-.013-.039-.024-.058-.035l-.024-.014q-.039-.021-.081-.039L10.95.081L10.895.06l-.039-.012l-.063-.016l-.028-.007c-.031-.006-.062-.01-.093-.014H1.093A1.09 1.09 0 0 0 .002 1.092V22.91c0 .603.489 1.091 1.091 1.091h17.454c.603 0 1.091-.489 1.091-1.091v-5.891h3.329l-.865.865a1.091 1.091 0 1 0 1.541 1.542l.001-.001l2.728-2.722q.039-.044.073-.086zM11.636 3.724l4.275 4.275h-4.275zm5.818 18.093H2.182V2.181h7.273V9.09c0 .603.489 1.091 1.091 1.091h6.909v4.655h-4.759a1.091 1.091 0 0 0 0 2.182h.032h-.002h4.727z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Expressionless(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.316C6.099 1.316 1.316 6.099 1.316 12S6.099 22.684 12 22.684S22.684 17.901 22.684 12c-.012-5.896-4.788-10.672-10.683-10.684zm0 22.297C5.586 23.613.387 18.414.387 12S5.586.387 12 .387S23.613 5.586 23.613 12v.015c0 6.405-5.192 11.597-11.597 11.597z"/><path fill="currentColor" d="M12 24C5.386 23.966.034 18.614 0 12.003V12C0 5.373 5.373 0 12 0s12 5.373 12 12c-.034 6.614-5.386 11.966-11.997 12zM12 .774C5.8.774.774 5.8.774 12S5.8 23.226 12 23.226S23.226 18.2 23.226 12C23.222 5.802 18.198.779 12.001.774zm0 22.297C5.886 23.071.929 18.114.929 12S5.886.929 12 .929S23.071 5.886 23.071 12S18.114 23.071 12 23.071m0-21.368C6.313 1.703 1.703 6.313 1.703 12S6.313 22.297 12 22.297S22.297 17.687 22.297 12v-.005c0-5.684-4.608-10.292-10.292-10.292z"/><path fill="currentColor" d="M19.045 9.135a2.71 2.71 0 1 1-5.42 0h5.419zm-8.671 0a2.745 2.745 0 0 1-2.708 2.71h-.002a2.745 2.745 0 0 1-2.71-2.708v-.002zm6.039 8.052H7.51a.499.499 0 0 1-.464-.463v-.002a.499.499 0 0 1 .463-.464h8.905a.499.499 0 0 1 .464.463v.002a.41.41 0 0 1-.406.468c-.021 0-.041-.002-.061-.004z"/><path fill="currentColor" d="M16.413 17.574H7.51a.852.852 0 0 1 0-1.704h8.903a.852.852 0 0 1 0 1.704m-8.826-.851c-.02.02-.033.047-.033.077s.013.058.033.077h8.903a.077.077 0 0 0 .077-.077c0-.077 0-.077-.077-.077z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Eye(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 0C7.164 0 0 11.844 0 11.844S7.164 24 16 24s16-12.156 16-12.156S24.836 0 16 0m0 20a8 8 0 1 1 0-16a8 8 0 0 1 0 16"/><path fill="currentColor" d="M20 12.016a4 4 0 1 1-8 0a4 4 0 0 1 8 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Facebook(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.462.173v3.808h-2.265a2.11 2.11 0 0 0-1.675.521l.002-.002a2.368 2.368 0 0 0-.432 1.566v-.008v2.726h4.229l-.56 4.27H8.098V24H3.681V13.053H.001V8.784h3.68V5.639a5.56 5.56 0 0 1 1.502-4.162l-.003.003A5.418 5.418 0 0 1 9.185.002h-.013a24.124 24.124 0 0 1 3.406.185l-.117-.012z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Famale(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.041 16.683a14.884 14.884 0 0 1-.035-.72c2.549-.261 4.338-.872 4.338-1.585c-.007 0-.006-.03-.006-.041C16.432 12.619 19.99.417 13.367.663a3.344 3.344 0 0 0-2.196-.664h.008C2.208.677 6.175 12.202 4.13 14.377h-.004c.008.698 1.736 1.298 4.211 1.566c-.007.17-.022.381-.054.734C7.256 19.447.321 18.671.001 24h22.294c-.319-5.33-7.225-4.554-8.253-7.317z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Favorite(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 2.089v21.912l6.546-6.26l6.545 6.26V2.089A2.11 2.11 0 0 0 10.982 0l-.077.001h.004h-8.726L2.11 0A2.109 2.109 0 0 0 .001 2.088z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fi(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 24C5.373 24 0 18.627 0 12S5.373 0 12 0s12 5.373 12 12c-.008 6.624-5.376 11.992-11.999 12zm2.24-13.84l.514 7.475h3.032l.542-7.475l-.273.096a5.904 5.904 0 0 1-1.794.273l-.079-.001h.004h-.026a5.163 5.163 0 0 1-1.955-.381l.035.012zm-8.6-2.244l.96 9.721h3.128l.385-3.898l2.695-.674v-1.492H10.32l.16-1.476l3.032-.193l-.498-1.989zM16.267 6.36l-.065-.001c-.491 0-.94.179-1.285.476l.003-.002a1.59 1.59 0 0 0-.538 1.254v-.003l-.001.052c0 .48.211.911.545 1.206l.002.002c.364.3.834.481 1.348.481s.984-.182 1.351-.484l-.004.003c.336-.295.546-.725.546-1.204l-.001-.05v.002l.001-.047c0-.479-.211-.909-.544-1.202l-.002-.002a1.974 1.974 0 0 0-1.295-.482l-.064.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func File(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.638 8.945a1.08 1.08 0 0 0-.249-.692l.001.002l-.004-.006a1.03 1.03 0 0 0-.063-.069l-.009-.009l-.028-.029L11.463.321a.83.83 0 0 0-.07-.063l-.022-.02l-.054-.04L11.29.18l-.058-.036l-.024-.014c-.027-.015-.054-.027-.081-.039l-.033-.013l-.057-.021L11 .045l-.067-.017l-.026-.006a1.53 1.53 0 0 0-.094-.015h-9.72C.493.007.006.492.002 1.091v21.818C.002 23.512.491 24 1.093 24h17.454c.603 0 1.091-.489 1.091-1.091V8.974l.001-.029zM11.781 3.72l4.13 4.135h-4.13zM2.182 21.818V2.181h7.42v6.767c0 .603.489 1.091 1.091 1.091h6.767v11.779z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FileOne(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.638 8.945c0-.228-.07-.439-.19-.613l.002.004a1.516 1.516 0 0 0-.06-.08l-.005-.006a1.03 1.03 0 0 0-.063-.069l-.01-.01l-.028-.029L11.462.321a.83.83 0 0 0-.07-.063l-.022-.02l-.054-.04l-.027-.019a1.787 1.787 0 0 0-.058-.035l-.025-.015a1.205 1.205 0 0 0-.081-.039l-.034-.014l-.054-.02l-.039-.012l-.061-.016l-.031-.007a1.357 1.357 0 0 0-.092-.014H1.092C.492.007.005.491.001 1.09v21.818c0 .603.489 1.091 1.091 1.091h17.454c.603 0 1.091-.489 1.091-1.091V8.973zM11.781 3.72l4.13 4.135h-4.13zM2.182 21.818V2.181h7.42v6.767c0 .603.489 1.091 1.091 1.091h6.767v11.779z"/><path fill="currentColor" d="M13.454 17.454H6.15a1.091 1.091 0 0 0 0 2.182h.032h-.002h7.304a1.091 1.091 0 0 0 0-2.182h-.032zm0-4.363H6.15a1.091 1.091 0 0 0 0 2.182h.032h-.002h7.304a1.091 1.091 0 0 0 0-2.182h-.032z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Film(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 13.75C0 19.411 4.589 24 10.25 24S20.5 19.411 20.5 13.75a12.063 12.063 0 0 0-1.531-5.501L19 8.31a3.781 3.781 0 0 1-.781-2.309V6c0-1.75 1.846-3.5 3.78-3.5h.247a1.251 1.251 0 0 0 .002-2.5h-.252c-3.075 0-5.706 1.64-6.33 5a9.12 9.12 0 0 0-5.438-1.499l.019-.001C4.588 3.503.002 8.09.001 13.749zm9 0a1.25 1.25 0 1 1 1.25 1.251h-.003a1.25 1.25 0 0 1-1.246-1.25zm-6.5 0a2.247 2.247 0 1 1 0 .003zm11 0a2.247 2.247 0 1 1 0 .003zM8 19.25a2.25 2.25 0 1 1 2.25 2.25h-.004A2.25 2.25 0 0 1 8 19.25m0-11a2.247 2.247 0 1 1 4.494.002a2.247 2.247 0 0 1-4.494 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Filter(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m11.2 24l6.4-4v-6.4L28.8 0H0l11.2 13.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fire(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.535 8.998a16.637 16.637 0 0 0-2.513-2.734l-.02-.017a.383.383 0 0 0-.477-.077l.002-.001a.332.332 0 0 0-.174.374v-.002l-.014-.009c.499.919.796 2.012.804 3.173v.002c.008.71-.037 1.207-.038 1.216a4.32 4.32 0 0 1-.058.289c-.059.289-.453 1.141-1.011.907s-.592-1.129-.592-1.129h-.008a13.76 13.76 0 0 0-4.468-9.452l-.011-.01l-.005-.005l-.005-.005A10.687 10.687 0 0 0 8.306.324l-.05-.027a1.292 1.292 0 0 1-.145-.082l.006.003a1.424 1.424 0 0 0-.224-.115L7.882.099c-.05-.026-.102-.046-.156-.07a.421.421 0 0 0-.579.231l-.001.003a2.47 2.47 0 0 0 .196 1.16l-.006-.016C7.995 3.072 8.378 5 8.378 7.018c0 1.202-.136 2.372-.392 3.496l.02-.105a1.53 1.53 0 0 1-.942 1.054l-.01.003a1.272 1.272 0 0 1-1.317-.837l-.003-.009c-.026-.072-.058-.139-.078-.217v-.014c-.482-1.718.109-4.291.522-5.725c.277-1.061-1.025 0-1.025 0h-.005c-.256.216-.532.46-.822.73v-.01c-.367.336-.852.802-1.351 1.335c-.27.289-3.869 4.249-2.747 9.658c1.113 4.406 5.042 7.615 9.72 7.615c5.081 0 9.278-3.786 9.924-8.69l.005-.051c.024-.199.046-.397.056-.594c.005-.062.006-.127.009-.192a10.271 10.271 0 0 0-1.43-5.515l.026.048z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Firefox(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12.474 24l-.182.001c-2.534 0-4.886-.778-6.83-2.108l.042.027a12.68 12.68 0 0 1-4.521-5.425l-.032-.081a11.502 11.502 0 0 1-.925-4.139l-.001-.022a17.382 17.382 0 0 1 .381-4.707l-.021.115a17.65 17.65 0 0 1 1.58-4.41l-.046.099A11.506 11.506 0 0 1 4.386.005l.006-.006l-.152 3.882q.152-.194.939-.214t.967.214a4.65 4.65 0 0 1 2.187-1.895l.031-.011a7.624 7.624 0 0 1 3.228-.815h.012a9.01 9.01 0 0 0-1.629 2.011l-.023.041a3.823 3.823 0 0 0-.808 2.254v.005c.243.077.535.142.834.183l.029.003q.518.076.87.104t.939.055t.698.042q.207.069.131.629a2.308 2.308 0 0 1-.426 1.046l.004-.006a2.251 2.251 0 0 1-.227.255l-.001.001a3.937 3.937 0 0 1-.757.48l-.024.01a4.616 4.616 0 0 1-1.367.466l-.028.004l.212 2.616l-1.92-.926a1.725 1.725 0 0 0-.102 1.139l-.003-.012c.096.357.268.666.499.921l-.002-.002c.242.26.546.457.89.569l.015.004a2.277 2.277 0 0 0 1.134.088l-.014.002c.509-.093.963-.26 1.377-.492l-.022.012q.65-.352 1.154-.622a1.89 1.89 0 0 1 1.019-.24h-.004h.021c.467 0 .893.173 1.218.458l-.002-.002a.9.9 0 0 1 .269.904l.001-.006q-.014.028-.034.076a1.023 1.023 0 0 1-.119.174l.001-.001a.87.87 0 0 1-.245.212l-.004.002a1.388 1.388 0 0 1-.427.144l-.008.001a2.072 2.072 0 0 1-.651.013l.011.001a4.905 4.905 0 0 1-1.973 1.859l-.027.013a5.432 5.432 0 0 1-2.189.452c-.249 0-.493-.016-.733-.048l.028.003a5.543 5.543 0 0 0 2.207 1.132l.038.008a5.486 5.486 0 0 0 2.362.075l-.034.005a8.518 8.518 0 0 0 2.183-.742l-.049.022a7.07 7.07 0 0 0 1.773-1.212l-.004.003a5.456 5.456 0 0 0 1.098-1.411l.014-.029a5.531 5.531 0 0 0 .539-2.672v.012a8.08 8.08 0 0 0-.537-2.66l.019.055a4.221 4.221 0 0 0-1.086-1.719l-.002-.001c.723.3 1.347.669 1.91 1.111l-.017-.013c.475.41.84.933 1.056 1.529l.008.025A10.383 10.383 0 0 0 16.992.922l-.016-.013a10.4 10.4 0 0 1 5.674 3.834l.02.027a12.038 12.038 0 0 1 2.086 7.168v-.017l.001.12c0 1.217-.205 2.386-.583 3.474l.022-.074a11.75 11.75 0 0 1-1.729 3.32l.022-.031a14.386 14.386 0 0 1-2.583 2.685l-.029.022a11.975 11.975 0 0 1-7.405 2.56h-.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FirstAidAlt(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M25.259 4.206h-6.264V2.183A2.184 2.184 0 0 0 16.813.001h-6.015a2.184 2.184 0 0 0-2.182 2.182v2.023H2.344A2.342 2.342 0 0 0 .002 6.548v15.11A2.342 2.342 0 0 0 2.344 24h22.922a2.342 2.342 0 0 0 2.342-2.342V6.547a2.342 2.342 0 0 0-2.342-2.342h-.007zM10.483 2.182V2.18c0-.17.138-.308.308-.308h6.026c.17 0 .307.138.307.307v2.026h-6.64zm3.319 18.591a6.674 6.674 0 1 1-.001-13.345a6.674 6.674 0 0 1 .002 13.344z"/><path fill="currentColor" d="M15.446 9.374h-3.302v3.068H9.07v3.312h3.078v3.068h3.302V15.75h3.078v-3.312h-3.082z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FiveHundredPx(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m18.58 20.72l-.08.08c-1.971 1.981-4.7 3.206-7.714 3.206S5.043 22.78 3.072 20.8a10.739 10.739 0 0 1-3.059-5.925l-.008-.062q-.054-.375.64-.48q.683-.107.75.268a.077.077 0 0 1 .014.045v.009c.157.813.37 1.527.644 2.209l-.028-.08a9.083 9.083 0 0 0 2.035 3.025l.001.001a9.467 9.467 0 0 0 6.716 2.782a9.477 9.477 0 0 0 3.767-.776l-.062.024a9.611 9.611 0 0 0 3.014-2.036l.08-.08a.32.32 0 0 1 .337-.08l-.002-.001c.173.068.32.168.441.294q.482.499.228.782zm-6.32-8.24l-.88.88l.844.844q.282.282-.094.656a.642.642 0 0 1-.426.228h-.003a.344.344 0 0 1-.254-.133l-.001-.001l-.83-.817l-.88.88a.273.273 0 0 1-.18.068l-.022-.001h.001a.605.605 0 0 1-.414-.214l-.001-.001l-.026-.026a.545.545 0 0 1-.24-.385v-.003a.402.402 0 0 1 .107-.228l.88-.87l-.88-.88q-.214-.214.187-.602a.663.663 0 0 1 .411-.24h.004a.305.305 0 0 1 .175.068h-.001l.87.88l.87-.87q.24-.228.64.174q.354.368.139.594zm6.308.763v.02c0 2.16-.881 4.115-2.303 5.524l-.001.001a7.956 7.956 0 0 1-2.453 1.667l-.051.02c-.901.39-1.951.616-3.054.616s-2.152-.227-3.105-.636l.051.02a7.984 7.984 0 0 1-2.505-1.687a7.625 7.625 0 0 1-1.656-2.441l-.019-.05a2.141 2.141 0 0 1-.198-.521l-.003-.015h-.014q-.121-.362.576-.59q.67-.214.8.16A7.121 7.121 0 0 0 5.94 17.58l-.008-.009h.014v-4.566A4.298 4.298 0 0 1 7.31 9.9l.002-.002a4.65 4.65 0 0 1 3.392-1.379H10.7l.074-.001c1.287 0 2.45.528 3.285 1.379l.001.001a4.51 4.51 0 0 1 1.393 3.265l-.001.074v-.004a4.767 4.767 0 0 1-4.754 4.754h-.016a5.249 5.249 0 0 1-1.523-.224l.037.01q-.375-.147-.174-.817q.214-.683.59-.576l.187.04q.187.04.442.08c.118.022.256.036.396.04h.004l.064.001c.898 0 1.709-.367 2.293-.96a3.163 3.163 0 0 0 .965-2.277l-.001-.063v-.052c0-.88-.37-1.674-.963-2.234l-.001-.001a3.174 3.174 0 0 0-2.271-.952l-.077.001h.004l-.061-.001c-.928 0-1.76.414-2.32 1.068l-.003.004a3.14 3.14 0 0 0-.857 2.142v5.532a6.046 6.046 0 0 0 3.189.898h.055h-.003h.013c.89 0 1.738-.18 2.509-.505l-.043.016c2.346-.982 3.963-3.258 3.964-5.913l.001-.088a6.174 6.174 0 0 0-1.888-4.451l-.002-.002c-1.166-1.159-2.773-1.875-4.547-1.875s-3.381.716-4.547 1.875A9.697 9.697 0 0 0 5.1 9.844l-.018.024l-.026.026a1.476 1.476 0 0 1-.174.207a.84.84 0 0 1-.282.126l-.006.001a.808.808 0 0 1-.522-.042l.005.002a1.345 1.345 0 0 1-.493-.223l.004.003a.44.44 0 0 1-.207-.354V.506a.508.508 0 0 1 .492-.507h.018h-.001h11.746q.4 0 .4.737t-.4.737H4.783v6.469h.014a9.836 9.836 0 0 1 2.748-1.916l.06-.025a7.735 7.735 0 0 1 3.054-.616h.042h-.002h.046c1.087 0 2.121.227 3.057.635l-.049-.019c2.848 1.214 4.808 3.99 4.808 7.223v.024v-.001zm-.419-7.779a.324.324 0 0 1 .121.239v.005c0 .088-.028.17-.075.238l.001-.001q-.074.107-.221.282q-.348.348-.522.348a.291.291 0 0 1-.214-.093a10.507 10.507 0 0 0-2.703-1.752l-.067-.027a8.971 8.971 0 0 0-3.622-.75h-.079h.004a9.774 9.774 0 0 0-3.575.679l.067-.023q-.362.134-.602-.495a1.32 1.32 0 0 1-.107-.508v-.001a.312.312 0 0 1 .212-.267l.002-.001a9.718 9.718 0 0 1 4.015-.763h-.01h.016c1.523 0 2.972.315 4.286.884l-.07-.027a10.313 10.313 0 0 1 3.149 2.04l-.005-.004z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flag(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.705 3.07a.752.752 0 0 0-.545-.233h-.006c-.251.051-.474.14-.675.263l.009-.005c-.338.171-.698.362-1.076.574a7.493 7.493 0 0 1-1.29.559l-.055.016a4.513 4.513 0 0 1-1.436.258h-.004l-.05.001c-.373 0-.726-.086-1.04-.24l.014.006a20.206 20.206 0 0 0-2.524-1.031l-.154-.045a8.017 8.017 0 0 0-2.385-.355h-.065h.003a11.575 11.575 0 0 0-5.205 1.498l.057-.03c-.4.198-.682.351-.866.464l-.183-1.341a1.954 1.954 0 1 0-2.306.186l.008.005l2.64 19.385c.079.565.559.995 1.139.996h.001c.056 0 .111-.004.165-.011l-.006.001a1.153 1.153 0 0 0 .992-1.302l.001.006l-1.062-7.791a11.501 11.501 0 0 1 4.967-1.439l.031-.001h.009c.619 0 1.216.098 1.774.28l-.041-.011a7.492 7.492 0 0 1 1.452.606l-.039-.02c.374.211.812.409 1.268.569l.058.018c.449.168.968.266 1.509.27h.002a9.868 9.868 0 0 0 4.554-1.443l-.043.025c.185-.09.343-.202.48-.337a.664.664 0 0 0 .152-.473v.002v-9.335a.753.753 0 0 0-.23-.543z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flash(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 13.714h7.875L5.137 24l12.006-13.714H9.268L12 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flickr(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.501 11.65a5.25 5.25 0 1 1-10.5 0a5.25 5.25 0 0 1 10.5 0m13.5 0a5.25 5.25 0 1 1-10.5 0a5.25 5.25 0 0 1 10.5 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Flipboard(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h7.7v24H0zm8.5 8.5h7.8v7.8H8.5zm0-8.5H24v7.7H8.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func FlotationRing(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m23.278 7.834l.022-.007l-.253-.581c-.007-.007-.007-.022-.015-.03s-.015-.038-.022-.052l-.29-.64l-.015.007c-1.247-2.409-3.219-4.296-5.623-5.405l-.074-.031l.008-.022l-.67-.261l-.662-.253l-.007.022A11.85 11.85 0 0 0 11.998 0h-.004a11.91 11.91 0 0 0-4.19.757L7.886.73L7.878.708l-1.31.566l.008.022c-2.42 1.239-4.318 3.209-5.435 5.615l-.031.073l-.022-.007l-.522 1.325l.022.007a11.71 11.71 0 0 0-.589 3.704c0 1.461.262 2.86.742 4.153l-.027-.083l-.022.007l.275.648l.022.045c.007.007.007.022.015.03l.253.581l.022-.007c1.222 2.407 3.165 4.3 5.542 5.428l.072.031l-.007.022l.655.261l.662.268l.007-.022c1.137.396 2.448.625 3.812.625c1.474 0 2.886-.267 4.19-.756l-.082.027l.008.022l1.31-.566l-.008-.022c2.399-1.229 4.285-3.174 5.409-5.55l.031-.072l.022.007l.522-1.325l-.022-.007c.395-1.129.623-2.431.623-3.786c0-1.487-.275-2.909-.776-4.219l.027.081zm-15.816 6.18l-.03-.059a4.997 4.997 0 0 1 2.567-6.51l.032-.013a4.837 4.837 0 0 1 1.965-.41h.017h-.001c.64.001 1.252.121 1.814.34l-.034-.012a4.99 4.99 0 0 1 2.772 2.609l.013.031l.03.059c.256.579.405 1.254.405 1.964c0 .669-.132 1.307-.373 1.889l.012-.033c-.755 1.851-2.541 3.132-4.626 3.132a4.983 4.983 0 0 1-4.551-2.952l-.013-.032zm10.671-3.946L21.968 8.4c.401 1.068.634 2.302.634 3.591c0 1.16-.188 2.276-.536 3.319l.021-.074l-3.887-1.534c.15-.511.236-1.098.236-1.706c0-.689-.111-1.352-.316-1.973l.013.044zm-6.128-8.646h.006c1.125 0 2.209.176 3.226.501l-.075-.021l-.804 2.04l-.73 1.854a6.307 6.307 0 0 0-1.613-.204c-.67 0-1.317.102-1.924.291l.046-.012l-1.668-3.835a9.997 9.997 0 0 1 3.484-.614h.057h-.003zM1.93 8.832l3.894 1.531a6.368 6.368 0 0 0 .071 3.523l-.012-.045l-3.84 1.668a10.256 10.256 0 0 1-.607-3.512c0-1.131.179-2.219.511-3.239l-.021.074zm10.112 13.746a10.63 10.63 0 0 1-3.355-.543l.075.022l1.534-3.887a6.318 6.318 0 0 0 1.749.241c.67 0 1.316-.102 1.923-.291l-.046.012l1.668 3.842a10.435 10.435 0 0 1-3.531.602h-.019h.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fog(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.506 24a1.019 1.019 0 1 1 0-2.038h6.873a1.019 1.019 0 1 1 0 2.038zm3.263-5.491a1.019 1.019 0 1 1 0-2.038h9.728a1.019 1.019 0 1 1 0 2.038zm-2.04-5.49a1.019 1.019 0 1 1 0-2.038h13.399a1.019 1.019 0 1 1 0 2.038zM3.058 7.525a1.019 1.019 0 1 1 0-2.038h15.438a1.019 1.019 0 1 1 0 2.038zM1.019 2.039a1.019 1.019 0 1 1 0-2.038h23.187a1.019 1.019 0 1 1 0 2.038z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Folder(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m.061 8.78l4.801 14.399c.164.481.611.82 1.138.821h19.199c.048 0 .095-.003.141-.009l-.006.001l.036-.006a.785.785 0 0 0 .094-.018l.038-.009c.031-.008.062-.018.094-.028l.036-.013c.031-.012.061-.025.094-.041l.029-.014a.981.981 0 0 0 .114-.068l-.004.003a1.05 1.05 0 0 0 .106-.081l-.002.001l.028-.025l.067-.064l.031-.034a.794.794 0 0 0 .058-.072l.017-.021l.008-.012q.034-.048.063-.1l.007-.011a.952.952 0 0 0 .055-.115l.003-.008l.007-.019q.021-.052.036-.106c0-.012.007-.024.009-.037c.008-.03.014-.06.02-.094c0-.015.005-.03.007-.045l.008-.085c.002-.028 0-.034 0-.051V1.2A1.2 1.2 0 0 0 25.2 0h-9.6a1.2 1.2 0 0 0-1.2 1.2v1.2H6a1.2 1.2 0 0 0-1.2 1.2v3.6H1.2A1.2 1.2 0 0 0 .065 8.786l-.003-.008zM23.999 2.4v13.003l-2.462-7.385a1.202 1.202 0 0 0-1.138-.821H7.2V4.8h8.4a1.2 1.2 0 0 0 1.2-1.2V2.401zM2.865 9.6h16.67l4 12H6.865z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Font(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.106 0H1.895A1.895 1.895 0 0 0 0 1.895v4.924l-.001.053a1.895 1.895 0 0 0 3.79 0l-.001-.056v.003v-3.03h6.315v16.422H7.178a1.896 1.896 0 0 0-.002 3.79h9.645a1.896 1.896 0 0 0 .002-3.79h-2.929V3.789h6.315v3.03a1.896 1.896 0 0 0 3.79.002V1.894a1.895 1.895 0 0 0-1.894-1.895z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Fontisto(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 24C5.373 24 0 18.627 0 12S5.373 0 12 0s12 5.373 12 12c-.008 6.624-5.376 11.992-11.999 12zm2.24-13.84l.514 7.475h3.032l.542-7.475l-.273.096a5.904 5.904 0 0 1-1.794.273l-.079-.001h.004h-.026a5.163 5.163 0 0 1-1.955-.381l.035.012zm-8.6-2.244l.96 9.721h3.128l.385-3.898l2.695-.674v-1.492H10.32l.16-1.476l3.032-.193l-.498-1.989zM16.267 6.36l-.065-.001c-.491 0-.94.179-1.285.476l.003-.002a1.59 1.59 0 0 0-.538 1.254v-.003l-.001.052c0 .48.211.911.545 1.206l.002.002c.364.3.834.481 1.348.481s.984-.182 1.351-.484l-.004.003c.336-.295.546-.725.546-1.204l-.001-.05v.002l.001-.047c0-.479-.211-.909-.544-1.202l-.002-.002a1.974 1.974 0 0 0-1.295-.482l-.064.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Forward(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M31.489 10.982L14.796.146c-.657-.427-1.463.136-1.463 1.02V7.85L1.462.15C.805-.277-.001.286-.001 1.17v21.663c0 .884.805 1.447 1.463 1.02l11.871-7.705v6.685c0 .884.805 1.447 1.463 1.02l16.693-10.835a1.27 1.27 0 0 0 .003-2.037l-.003-.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Foursquare(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.785 6.258l.534-2.8a.643.643 0 0 0-.13-.578l.001.001a.628.628 0 0 0-.495-.241h-.015h.001H3.405a.717.717 0 0 0-.54.244l-.001.001a.776.776 0 0 0-.223.533v15.875q0 .101.086.014l4.196-5.075a2.12 2.12 0 0 1 .539-.475l.009-.005a1.617 1.617 0 0 1 .697-.108h-.005h3.465c.2 0 .382-.08.515-.209a.833.833 0 0 0 .258-.42l.001-.006q.346-1.874.534-2.754a.68.68 0 0 0-.166-.578v.001a.664.664 0 0 0-.526-.274H7.971a.937.937 0 0 1-.937-.937l.001-.032v.002v-.609l-.001-.032c0-.256.105-.487.274-.653a.949.949 0 0 1 .659-.265l.035.001H8h4.989a.805.805 0 0 0 .506-.195l-.001.001a.721.721 0 0 0 .29-.423zm3.273-3.2q-.216 1.053-.771 3.84t-1.002 5.046q-.447 2.257-.505 2.502q-.086.32-.13.469c-.062.183-.13.338-.21.485l.008-.016a1.437 1.437 0 0 1-.352.475l-.002.001c-.16.121-.343.223-.54.297l-.015.005a2.208 2.208 0 0 1-.791.144h-.047h.002h-3.915a.41.41 0 0 0-.312.144v.001q-.115.13-6.142 7.12c-.21.235-.507.389-.839.411h-.004a1.11 1.11 0 0 1-.706-.082l.007.003a1.352 1.352 0 0 1-.79-1.414l-.001.007V2.162A2.31 2.31 0 0 1 .552.681L.549.684A2.074 2.074 0 0 1 2.288.003L2.28.002h12.804a1.964 1.964 0 0 1 1.828.76l.004.005c.183.392.289.851.289 1.336c0 .342-.053.671-.151.98l.006-.023zm0 0l-2.278 11.39q.058-.245.505-2.502T16.287 6.9t.771-3.843z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Frowning(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.316C6.099 1.316 1.316 6.099 1.316 12S6.099 22.684 12 22.684S22.684 17.901 22.684 12c-.012-5.896-4.788-10.672-10.683-10.684zm0 22.297C5.586 23.613.387 18.414.387 12S5.586.387 12 .387S23.613 5.586 23.613 12v.015c0 6.405-5.192 11.597-11.597 11.597z"/><path fill="currentColor" d="M12 24C5.386 23.966.034 18.614 0 12.003V12C0 5.373 5.373 0 12 0s12 5.373 12 12c-.034 6.614-5.386 11.966-11.997 12zM12 .774C5.8.774.774 5.8.774 12S5.8 23.226 12 23.226S23.226 18.2 23.226 12C23.222 5.802 18.198.779 12.001.774zm0 22.297C5.886 23.071.929 18.114.929 12S5.886.929 12 .929S23.071 5.886 23.071 12S18.114 23.071 12 23.071m0-21.368C6.313 1.703 1.703 6.313 1.703 12S6.313 22.297 12 22.297S22.297 17.687 22.297 12v-.005c0-5.684-4.608-10.292-10.292-10.292z"/><path fill="currentColor" d="M9.677 9.91a2.09 2.09 0 0 1-4.18 0v-.009c0-1.15.932-2.082 2.082-2.082h.018c1.15 0 2.082.932 2.082 2.082v.009zm8.904 0v.009c0 1.15-.932 2.082-2.082 2.082h-.009a2.09 2.09 0 0 1 0-4.18a2.132 2.132 0 0 1 2.09 2.088v.002zm-1.936 8.903a.355.355 0 0 1-.387-.308v-.002c-.6-1.813-2.279-3.097-4.258-3.097s-3.658 1.285-4.249 3.066l-.009.032a.353.353 0 0 1-.23.232l-.002.001c-.041.037-.095.059-.155.059s-.114-.022-.155-.059a.44.44 0 0 1-.308-.545l-.001.003c.72-2.175 2.735-3.717 5.11-3.717s4.39 1.542 5.099 3.679l.011.038v.31l-.232.232a.285.285 0 0 1-.233.078h.001z"/><path fill="currentColor" d="m16.645 19.2l-.039.001a.769.769 0 0 1-.733-.537l-.001-.005c-.529-1.63-2.034-2.788-3.809-2.788l-.065.001h.003a4.096 4.096 0 0 0-3.862 2.758l-.009.029a.795.795 0 0 1-.383.462l-.004.002a1.49 1.49 0 0 1-.625.077h.006a.795.795 0 0 1-.462-.383l-.002-.004a.858.858 0 0 1-.076-.626l-.001.006c.743-2.308 2.872-3.948 5.384-3.948h.038h-.002h.036a5.648 5.648 0 0 1 5.373 3.909l.011.04a.58.58 0 0 1-.079.62l.001-.001c-.077.232-.31.31-.465.387h-.232v-.774l.077-.077c-.685-1.991-2.538-3.399-4.722-3.406h-.001a5.004 5.004 0 0 0-4.712 3.371l-.011.035v.077l.465.077l-.387-.155c.628-1.947 2.424-3.33 4.543-3.33l.107.001h-.005a4.908 4.908 0 0 1 4.635 3.295l.01.034z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gbp(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.386 17.2v6.277c0 .29-.235.525-.525.525h-.022h.001H.524a.525.525 0 0 1-.525-.525v-.022v.001v-2.56a.553.553 0 0 1 .545-.546h1.654v-6.526H.561a.537.537 0 0 1-.528-.537v-2.261c0-.29.235-.525.525-.525H.58H.579h1.619V6.7a6.162 6.162 0 0 1 2.098-4.801l.007-.006A7.678 7.678 0 0 1 9.678.001h-.015h.008c2.187 0 4.185.807 5.713 2.139l-.01-.009c.098.087.162.21.17.348v.001l.002.047a.525.525 0 0 1-.122.337l.001-.001l-1.756 2.165a.527.527 0 0 1-.373.205h-.002a.456.456 0 0 1-.392-.119a3.802 3.802 0 0 0-.427-.311l-.016-.009a6.035 6.035 0 0 0-1.133-.532l-.043-.013A4.421 4.421 0 0 0 9.7 3.942h-.003a3.337 3.337 0 0 0-2.339.803l.004-.004a2.69 2.69 0 0 0-.886 2.102v-.005v3.665h5.206a.537.537 0 0 1 .537.528v.019v-.001v2.23a.553.553 0 0 1-.545.546h-5.2v6.46h7.057v-3.092a.537.537 0 0 1 .528-.537h.019h-.001h2.78c.147 0 .28.061.374.16a.537.537 0 0 1 .157.38v.003z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Genderless(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.334 12v-.02a9.354 9.354 0 1 0-2.74 6.614a8.963 8.963 0 0 0 2.74-6.462l-.001-.139zM24 12v.036c0 1.67-.349 3.258-.977 4.695l.029-.075c-1.226 2.905-3.491 5.17-6.318 6.367l-.078.029C15.282 23.651 13.682 24 12 24s-3.283-.349-4.733-.978l.077.03c-2.905-1.226-5.17-3.491-6.367-6.318l-.029-.078C.349 15.282 0 13.682 0 12s.349-3.283.978-4.733l-.03.077C2.174 4.439 4.439 2.174 7.266.977l.078-.029C8.718.349 10.318 0 12 0s3.283.349 4.733.978l-.077-.03c2.905 1.226 5.17 3.491 6.367 6.318l.029.078c.599 1.362.948 2.95.948 4.62v.039z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gg(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12 10.286l6.857 6.857L12 24.001l-12-12l12-12l3 3l-1.714 1.714L12 3.429L3.429 12L12 20.571l3.446-3.446L10.285 12zM22.286 0l12 12l-12 12l-3-3L21 19.286l1.286 1.286l8.571-8.571l-8.571-8.571l-3.446 3.446l5.161 5.125l-1.714 1.714l-6.857-6.857z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Git(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.058 20.277q0-1.339-2.21-1.339q-2.116 0-2.116 1.393q0 1.353 2.303 1.353q2.022-.003 2.022-1.406zm-.79-10.112l.002-.083c0-.48-.15-.925-.406-1.29l.005.007a1.384 1.384 0 0 0-1.196-.549h.004q-1.661 0-1.661 1.942Q3.016 12 4.677 12q1.592.001 1.592-1.834zm3.6-4.339v2.706q-.48.16-1.058.294c.132.333.211.719.214 1.123v.002l.002.118a4.42 4.42 0 0 1-.986 2.79l.007-.008a4.287 4.287 0 0 1-2.615 1.503l-.025.003a1.562 1.562 0 0 0-.802.363l.002-.002a1.035 1.035 0 0 0-.26.781v-.004v.02c0 .266.116.505.301.669l.001.001c.217.19.476.337.762.425l.015.004q.475.154 1.051.294t1.152.342c.402.14.749.31 1.072.515l-.021-.012c.328.215.591.504.771.844l.006.012c.191.363.303.793.303 1.249l-.001.074v-.004q.003 4.073-4.855 4.073h-.046c-.6 0-1.185-.061-1.75-.177l.056.01a5.946 5.946 0 0 1-1.586-.564l.033.016a2.952 2.952 0 0 1-1.165-1.085l-.007-.013a3.19 3.19 0 0 1-.436-1.62l.001-.09v.004q0-2.21 2.438-3.014v-.054a1.843 1.843 0 0 1-.898-1.692v.005q0-1.46.844-1.835v-.054a2.955 2.955 0 0 1-1.592-1.436l-.008-.017a4.455 4.455 0 0 1-.64-2.212v-.005l-.001-.11a4.14 4.14 0 0 1 1.272-2.99l.001-.001a4.32 4.32 0 0 1 3.036-1.241l.117.002h-.006h.031c.866 0 1.678.234 2.376.642l-.022-.012a8.08 8.08 0 0 0 2.976-.651l-.053.02zm4.259 11.799h-2.971q.054-.602.054-1.794V7.675q0-1.259-.054-1.714h2.974q-.054.442-.054 1.661v8.21q0 1.192.054 1.794zm8.049-2.974v2.629a4.715 4.715 0 0 1-2.338.522h.008l-.085.001c-.486 0-.949-.099-1.37-.278l.023.009a2.54 2.54 0 0 1-.935-.667l-.002-.003a2.917 2.917 0 0 1-.524-1.024l-.005-.021a5.995 5.995 0 0 1-.246-1.206l-.002-.026q-.054-.59-.054-1.379V8.504h.026V8.45q-.094 0-.254-.014t-.24-.014a5.46 5.46 0 0 0-.823.085l.033-.005V5.958h1.286V4.94l.001-.111a7.29 7.29 0 0 0-.086-1.122l.005.041h3.04a18.336 18.336 0 0 0-.08 2.231v-.022h2.288v2.545q-.201 0-.582-.026t-.569-.026h-1.135v4.888q0 1.754 1.165 1.754h.036c.532 0 1.026-.165 1.432-.447l-.008.005zM14.464 1.969v.023c0 .517-.199.987-.524 1.339l.001-.001a1.659 1.659 0 0 1-1.263.583h-.024h.001h-.024c-.513 0-.973-.225-1.287-.581l-.002-.002a1.937 1.937 0 0 1-.536-1.341v-.042c0-.523.201-1 .53-1.356l-.001.001a1.7 1.7 0 0 1 1.288-.59h.033h-.002h.029c.508 0 .962.231 1.262.594l.002.003c.32.358.515.833.515 1.354v.02v-.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Github(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.301 0h.093c2.242 0 4.34.613 6.137 1.68l-.055-.031a12.351 12.351 0 0 1 4.449 4.422l.031.058a12.182 12.182 0 0 1 1.654 6.166c0 5.406-3.483 10-8.327 11.658l-.087.026a.724.724 0 0 1-.642-.113l.002.001a.624.624 0 0 1-.208-.466v-.014v.001l.008-1.226q.008-1.178.008-2.154a2.844 2.844 0 0 0-.833-2.274a10.918 10.918 0 0 0 1.718-.305l-.076.017a6.508 6.508 0 0 0 1.537-.642l-.031.017a4.52 4.52 0 0 0 1.292-1.058l.006-.007a4.9 4.9 0 0 0 .84-1.645l.009-.035a7.888 7.888 0 0 0 .329-2.281l-.001-.136v.007l.001-.072a4.73 4.73 0 0 0-1.269-3.23l.003.003c.168-.44.265-.948.265-1.479a4.25 4.25 0 0 0-.404-1.814l.011.026a2.095 2.095 0 0 0-1.31.181l.012-.005a8.622 8.622 0 0 0-1.512.726l.038-.022l-.609.384c-.922-.264-1.981-.416-3.075-.416s-2.153.152-3.157.436l.081-.02q-.256-.176-.681-.433a9.103 9.103 0 0 0-1.272-.595l-.066-.022A2.174 2.174 0 0 0 5.837 5.1l.013-.002a4.2 4.2 0 0 0-.393 1.788c0 .531.097 1.04.275 1.509l-.01-.029a4.723 4.723 0 0 0-1.265 3.303v-.004l-.001.13c0 .809.12 1.591.344 2.327l-.015-.057c.189.643.476 1.202.85 1.693l-.009-.013a4.35 4.35 0 0 0 1.267 1.062l.022.011c.432.252.933.465 1.46.614l.046.011c.466.125 1.024.227 1.595.284l.046.004c-.431.428-.718 1-.784 1.638l-.001.012a3.056 3.056 0 0 1-.699.236l-.021.004c-.256.051-.549.08-.85.08h-.066h.003a1.882 1.882 0 0 1-1.055-.348l.006.004a2.84 2.84 0 0 1-.881-.986l-.007-.015a2.603 2.603 0 0 0-.768-.827l-.009-.006a2.331 2.331 0 0 0-.776-.38l-.016-.004l-.32-.048a1.048 1.048 0 0 0-.471.074l.007-.003q-.128.072-.08.184c.039.086.087.16.145.225l-.001-.001c.061.072.13.135.205.19l.003.002l.112.08c.283.148.516.354.693.603l.004.006c.191.237.359.505.494.792l.01.024l.16.368c.135.402.38.738.7.981l.005.004c.3.234.662.402 1.057.478l.016.002c.33.064.714.104 1.106.112h.007c.045.002.097.002.15.002c.261 0 .517-.021.767-.062l-.027.004l.368-.064q0 .609.008 1.418t.008.873v.014c0 .185-.08.351-.208.466h-.001a.717.717 0 0 1-.645.111l.005.001C3.486 22.286.006 17.692.006 12.285c0-2.268.612-4.393 1.681-6.219l-.032.058a12.351 12.351 0 0 1 4.422-4.449l.058-.031a11.898 11.898 0 0 1 6.073-1.645h.098h-.005zm-7.64 17.666q.048-.112-.112-.192q-.16-.048-.208.032q-.048.112.112.192q.144.096.208-.032m.497.545q.112-.08-.032-.256q-.16-.144-.256-.048q-.112.08.032.256q.159.157.256.047zm.48.72q.144-.112 0-.304q-.128-.208-.272-.096q-.144.08 0 .288t.272.112m.672.673q.128-.128-.064-.304q-.192-.192-.32-.048q-.144.128.064.304q.192.192.32.044zm.913.4q.048-.176-.208-.256q-.24-.064-.304.112t.208.24q.24.097.304-.096m1.009.08q0-.208-.272-.176q-.256 0-.256.176q0 .208.272.176q.256.001.256-.175zm.929-.16q-.032-.176-.288-.144q-.256.048-.224.24t.288.128t.225-.224z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gitlab(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.906 24L.403 14.723a1.073 1.073 0 0 1-.351-.497l-.002-.008a.926.926 0 0 1 .002-.609l-.002.007l1.463-4.437zM5.293.354l2.874 8.823H1.512L4.335.354a.517.517 0 0 1 .49-.353h.015h-.001L4.865 0c.212 0 .388.151.427.351v.003zm2.874 8.823h9.479L12.907 24zm17.595 4.436a.926.926 0 0 1-.002.609l.002-.007a1.074 1.074 0 0 1-.351.503l-.002.002L12.906 24L24.3 9.177zM21.477.354L24.3 9.177h-6.655L20.519.354a.436.436 0 0 1 .455-.353h-.001h.014c.227 0 .419.146.489.349z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Google(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 10.282h11.328c.116.6.184 1.291.187 1.997v.003l.002.222c0 2.131-.527 4.139-1.457 5.901l.033-.069a10.332 10.332 0 0 1-4.004 4.137l-.051.027a11.844 11.844 0 0 1-6.051 1.5h.012h-.044c-1.672 0-3.263-.348-4.704-.975l.076.03c-2.902-1.219-5.164-3.482-6.354-6.306l-.029-.077C.346 15.293-.001 13.687-.001 12s.348-3.293.975-4.75l-.03.078C2.163 4.426 4.426 2.164 7.25.974l.077-.029a11.598 11.598 0 0 1 4.624-.944h.051h-.003l.199-.002c3.045 0 5.811 1.197 7.853 3.147l-.004-.004l-3.266 3.141a6.585 6.585 0 0 0-4.792-1.86h.009h-.047A7.137 7.137 0 0 0 8.24 5.457l.032-.018c-2.246 1.358-3.725 3.788-3.725 6.562s1.479 5.204 3.691 6.543l.034.019a7.096 7.096 0 0 0 3.679 1.016h.05h-.003h.083c.864 0 1.695-.137 2.474-.392l-.056.016a6.301 6.301 0 0 0 1.893-.95l-.017.012a6.837 6.837 0 0 0 1.268-1.264l.012-.016c.312-.393.582-.841.79-1.321l.015-.039a5.52 5.52 0 0 0 .346-1.184l.005-.035H12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleDrive(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.375 0h8.678l8.735 15.165H18.16zm1.446 16.393h16.607L23.089 24H6.482zM8.25 1.874l4.446 7.607L4.339 24L0 16.392z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GooglePlay(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m3.751.61l13.124 7.546l-2.813 2.813zM1.032 0l12.047 12L1.033 24a1.72 1.72 0 0 1-1.032-1.643v.003V1.641L0 1.576C0 .875.42.272 1.022.005zm19.922 10.594c.414.307.679.795.679 1.344l-.001.065V12a1.558 1.558 0 0 1-.629 1.403l-.004.003l-2.813 1.593L15.14 12l3.047-3.047zM3.751 23.39l10.312-10.359l2.813 2.813z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GooglePlus(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14.969 11.814l.001.146a8.015 8.015 0 0 1-.929 3.757l.021-.044a6.588 6.588 0 0 1-2.55 2.632l-.032.017a7.55 7.55 0 0 1-3.851.951h.008h-.022a7.457 7.457 0 0 1-2.995-.623l.048.019a7.515 7.515 0 0 1-4.044-4.013l-.018-.049c-.382-.876-.604-1.896-.604-2.969S.224 9.545.625 8.62l-.019.049a7.515 7.515 0 0 1 4.013-4.044l.049-.018a7.365 7.365 0 0 1 2.951-.606h.018h-.001L7.761 4c1.937 0 3.695.762 4.992 2.004l-.003-.003l-2.073 1.99a4.187 4.187 0 0 0-3.046-1.177h.006h-.031c-.864 0-1.671.24-2.36.657l.02-.011a4.865 4.865 0 0 0-.022 8.342l.022.012a4.508 4.508 0 0 0 2.339.646h.033h-.002h.052c.549 0 1.077-.088 1.572-.25l-.036.01a4.044 4.044 0 0 0 1.204-.607l-.011.008a4.39 4.39 0 0 0 .81-.807l.008-.01c.198-.25.369-.535.5-.841l.009-.024c.093-.22.17-.478.22-.746l.004-.024H7.634v-2.625h7.208c.078.382.124.821.126 1.27v.002zM24 10.541v2.189h-2.177v2.177h-2.189V12.73H17.46v-2.189h2.177V8.364h2.189v2.177z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func GoogleWallet(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.435 9h.028c.276 0 .519.137.667.346l.002.003a28.197 28.197 0 0 1 4.807 10.171l.041.195H5.006C3.95 15.877 2.263 12.528.04 9.609l.053.072a.394.394 0 0 1-.039-.45l-.001.002a.406.406 0 0 1 .389-.234H.441zm7.491 4.782a41.134 41.134 0 0 1-1.78 5.545l.1-.281c-.791-3.019-1.967-5.668-3.497-8.069l.068.114c.339-1.787.552-3.862.589-5.98v-.033c1.591 2.474 3.105 5.335 4.375 8.325l.144.381zm1.322-9.496c1.98 2.7 3.71 5.783 5.046 9.067l.104.288c1.253 3.033 2.13 6.547 2.461 10.219l.01.141h-6.04q-.549-8.906-7.406-19.714zM23.53 12v.173c0 3.803-.494 7.491-1.42 11.002l.067-.3c-.688-5.42-2.387-10.324-4.909-14.694l.101.189A40.649 40.649 0 0 0 15.877.262l.072.287a.426.426 0 0 1 .074-.382l-.001.001a.406.406 0 0 1 .329-.167h.015h-.001h4.81c.194 0 .373.066.516.176l-.002-.001a.775.775 0 0 1 .3.436l.001.005c.977 3.37 1.539 7.242 1.539 11.245v.147v-.008z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Graphql(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.731 2.751L17.666 5.6a2.138 2.138 0 1 1 2.07 3.548l-.015.003v5.7a2.14 2.14 0 1 1-2.098 3.502l-.002-.002l-4.905 2.832a2.14 2.14 0 1 1-4.079.054l-.004.015l-4.941-2.844a2.14 2.14 0 1 1-2.067-3.556l.015-.003V9.15a2.14 2.14 0 1 1 1.58-3.926l-.01-.005c.184.106.342.231.479.376l.001.001l4.938-2.85a2.14 2.14 0 1 1 4.096.021l.004-.015zm-.515.877a.766.766 0 0 1-.057.057l-.001.001l6.461 11.19c.026-.009.056-.016.082-.023V9.146a2.14 2.14 0 0 1-1.555-2.603l-.003.015l.019-.072zm-3.015.059l-.06-.06l-4.946 2.852A2.137 2.137 0 0 1 2.749 9.12l-.015.004l-.076.021v5.708l.084.023l6.461-11.19zm2.076.507a2.164 2.164 0 0 1-1.207-.004l.015.004l-6.46 11.189c.286.276.496.629.597 1.026l.003.015h12.911c.102-.413.313-.768.599-1.043l.001-.001L11.28 4.194zm.986 16.227l4.917-2.838a1.748 1.748 0 0 1-.038-.142H4.222l-.021.083l4.939 2.852c.39-.403.936-.653 1.54-.653c.626 0 1.189.268 1.581.696l.001.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Grunt(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.765 8.859a.655.655 0 0 1 .325-.558l.003-.002q-.328-.094-.609-.187a3.626 3.626 0 0 1 2.549.385l-.018-.009a5.17 5.17 0 0 1 1.354 1.118l.006.007c.191.26.469.446.791.514l.009.002a1.93 1.93 0 0 0 .621-.002l-.011.002l.282-.094v.004c0 .254-.118.48-.303.627l-.002.001c-.188.16-.434.258-.703.258h-.024h.001a1.71 1.71 0 0 1-.286 1.037l.004-.006l-.328.516a2.133 2.133 0 0 0 .091-1.232l.003.014l-.141-.375a1.529 1.529 0 0 1-1.51.278l.011.003a2.358 2.358 0 0 1-1.08-.561l.002.002a2.423 2.423 0 0 1-.609-1.592v-.001a2.16 2.16 0 0 0-.187.743v.007l-.14-.24a2.338 2.338 0 0 1-.098-.656v-.001zm1.359-.093a2.08 2.08 0 0 0-.069.946L4.053 9.7c.022.305.173.572.398.748l.002.002a1.247 1.247 0 0 0 1.038.28l-.008.001a1.4 1.4 0 0 0 .754-.234l-.005.003a5.034 5.034 0 0 0-.855-.835l-.012-.009c-.36-.29-.765-.583-1.185-.855zm3.701 10.965a.391.391 0 0 1 .047.331l.001-.003a3.212 3.212 0 0 1-.613.939l.002-.002a.83.83 0 0 1-.528.187h-.012h.001h-.011a.836.836 0 0 1-.529-.188l.002.001a2.448 2.448 0 0 1-.65-.874l-.006-.016a.319.319 0 0 1 .048-.375a.64.64 0 0 1 .471-.141h-.003h1.36a.572.572 0 0 1 .424.141l-.001-.001zm2.955-8.81h-.023c-.269 0-.515-.098-.705-.259l.002.001a.8.8 0 0 1-.305-.628v-.005h.047a.834.834 0 0 0 .468.142l.05-.001h-.002a1.418 1.418 0 0 0 1.122-.556l.002-.003c.382-.446.831-.82 1.335-1.111l.025-.013a3.661 3.661 0 0 1 2.554-.37l-.024-.004q-.282.094-.609.187a.656.656 0 0 1 .328.559c0 .234-.035.459-.098.673l.004-.017l-.141.234a2.641 2.641 0 0 0-.194-.768l.006.017a2.807 2.807 0 0 1-.614 1.599l.004-.006a2.34 2.34 0 0 1-1.062.557l-.016.003a1.567 1.567 0 0 1-1.502-.284l.003.002l-.141.375a1.674 1.674 0 0 0 .002.812l-.002-.012l.141.421l-.375-.516a1.705 1.705 0 0 1-.282-1.034v.004zm.841-.421v.047h-.047c.203.109.44.191.689.232l.014.002a1.474 1.474 0 0 0 1.081-.33l-.003.002c.227-.178.378-.444.4-.746v-.003a2.066 2.066 0 0 0-.074-.952l.004.015c-.477.307-.882.6-1.271.913l.029-.022a3.706 3.706 0 0 0-.814.831zm.609 9.231a.32.32 0 0 1 .046.377l.001-.002a2.092 2.092 0 0 1-.651.84l-.005.004a.737.737 0 0 1-.539.234h-.002a.737.737 0 0 1-.541-.235a3.48 3.48 0 0 1-.6-.869l-.009-.021a.348.348 0 0 1 .071-.328a.49.49 0 0 1 .402-.141h-.002h1.406c.158.019.302.069.429.144zM7.638 7.969c.358.232.795.371 1.264.375h.001a2.37 2.37 0 0 0 1.274-.381l-.009.006l-.421.421a1.34 1.34 0 0 1-.838.421l-.006.001a1.343 1.343 0 0 1-.843-.421l-.001-.001zm8.622 6.888c.277.225.499.507.65.829l.006.014c.255.501.404 1.093.404 1.72c0 .27-.028.534-.08.788l.004-.025a3.217 3.217 0 0 1-2.179 2.571l-.023.006a3.947 3.947 0 0 1-1.026 1.402l-.005.004a3.457 3.457 0 0 1-2.102.887l-.01.001a4.504 4.504 0 0 1-3.01.937H8.9a4.482 4.482 0 0 1-2.961-.896l.011.008a3.72 3.72 0 0 1-2.111-.94l.003.003a4.454 4.454 0 0 1-1.066-1.38l-.012-.026a3.219 3.219 0 0 1-2.2-2.559l-.002-.018a3.77 3.77 0 0 1 .101-1.9l-.007.027a3.23 3.23 0 0 1 .565-1.085l-.005.007l.328-.375c.068.431.208.817.408 1.165l-.009-.016c.198.353.44.655.723.911l.003.003c.123-.833.193-1.795.193-2.774v-.14c0-.308-.089-.595-.244-.837l.004.006a1.377 1.377 0 0 0-.325-.373l-.003-.002l-.375-.234a6.958 6.958 0 0 1-1.279-.811l.014.011A1.644 1.644 0 0 1 0 10.502v-.013c0-.314.106-.604.284-.835l-.002.003c.135-.162.289-.301.461-.417l.008-.005l.094-.094q.187-.141.187-.234c.104-.262.165-.566.165-.883l-.001-.057v.003a1.8 1.8 0 0 0-.074-.813l.004.013a3.455 3.455 0 0 1-.366-.589l-.009-.02a1.849 1.849 0 0 1-.185-1.254l-.002.012c.086-.351.329-.633.648-.771l.007-.003c.274-.13.594-.214.93-.234h.007l.047-.282c.026-.111.043-.24.047-.372v-.003c0-.19-.052-.368-.143-.52l.003.005a1.925 1.925 0 0 0-.788-.651l-.012-.005a1.815 1.815 0 0 0-.549-.186L.75 2.296l-.56-.094l.375-.469c.177-.137.377-.262.589-.366l.02-.009a2.917 2.917 0 0 1 1.52-.422h.029h-.001c.403.001.789.07 1.149.195l-.025-.008a4.695 4.695 0 0 1 2.333 1.81l.011.017l.282-.094a2.092 2.092 0 0 1 .236-1.131l-.005.011c.124-.314.307-.581.538-.799l.001-.001a.99.99 0 0 1 .532-.327l.007-.001v.141q.047.937.328 1.124A2.38 2.38 0 0 1 9.921.001l.016-.003a2.388 2.388 0 0 0 .1 1.749l-.006-.015c.242-.221.52-.409.823-.551l.02-.009c.149-.083.326-.134.514-.141h.236l-.282.609a1.776 1.776 0 0 0-.182.789c0 .135.015.267.043.394l-.002-.012l.421.141a4.685 4.685 0 0 1 2.311-1.817l.033-.01a3.523 3.523 0 0 1 1.147-.187h.004c.41 0 .805.069 1.172.195l-.025-.008c.271.105.503.247.706.424l-.003-.002q.282.141.282.187l.375.469l-.56.094c-.21.035-.398.1-.57.192l.01-.005a1.557 1.557 0 0 0-.796.649l-.004.007a.839.839 0 0 0-.093.52v-.004a.55.55 0 0 0 .049.378l-.001-.003v.282l.064-.001c.322 0 .623.087.882.24l-.008-.004c.327.141.569.422.654.766l.002.008a1.835 1.835 0 0 1-.192 1.251l.005-.009c-.09.224-.201.416-.333.592l.005-.007a1.375 1.375 0 0 0-.115.828l-.001-.008l-.001.064c0 .314.059.614.167.89l-.006-.017c.053.09.115.167.186.234l.001.001l.094.047c.173.142.327.295.464.463l.005.006c.176.228.282.518.282.832v.011v-.001a1.64 1.64 0 0 1-.652 1.262l-.004.003a6.847 6.847 0 0 1-1.225.781l-.04.018l-.375.187a2.86 2.86 0 0 0-.321.41l-.007.011a1.546 1.546 0 0 0-.234.822v.024v-.001l-.001.17c0 .965.068 1.914.2 2.842l-.012-.107a3.112 3.112 0 0 0 1.123-2.052l.001-.014zM14.246 1.783a3.33 3.33 0 0 0-1.772 1.483l-.009.016l.56.282c.244-.438.576-.799.973-1.071l.011-.007c.35-.263.792-.422 1.271-.422h.043h-.002h.018c.143 0 .282.017.415.05l-.012-.002c.132-.082.286-.16.447-.226l.022-.008a2.97 2.97 0 0 0-1.166-.234h-.018c-.28 0-.549.051-.797.145l.015-.005zm.983.843h-.005c-.354 0-.68.124-.935.33l.003-.002c-.303.262-.52.616-.607 1.019l-.002.012c.171.167.325.351.46.547l.008.012l.56-.141c.114-.24.182-.521.187-.818v-.002c.005-.298.073-.579.192-.832l-.005.012zM2.063 2.112c.15-.002.296-.019.436-.05l-.014.003h.041c.479 0 .921.158 1.277.426l-.005-.004c.409.279.74.64.977 1.062l.008.016l.56-.282A3.33 3.33 0 0 0 3.586 1.79l-.023-.007a2.212 2.212 0 0 0-.783-.141h-.017h.001h-.006c-.421 0-.821.086-1.185.242l.02-.007c.184.074.339.154.487.244l-.014-.008zm2.061 1.875a1.942 1.942 0 0 0-.584-1.03l-.001-.001a1.457 1.457 0 0 0-.923-.328h-.038h.002l.141.141c.107.248.175.537.187.839v.005c.006.292.074.566.192.812l-.005-.012l.56.141c.122-.223.279-.411.466-.564zm-1.502 8.199c.244.153.442.358.581.601l.004.008c.171.282.282.617.307.976v.007a22.223 22.223 0 0 1-.152 3.671l.011-.108c.543.2 1.17.32 1.824.328h.004a3.32 3.32 0 0 0 .274-.588l.007-.024l.516-1.312q.516-1.265.703-1.64a19.876 19.876 0 0 1-1.125 4.357l.046-.136c.168-.03.362-.047.559-.047h5.501c.176 0 .347.017.513.05l-.017-.003a19.165 19.165 0 0 1-1.061-4.06l-.014-.11q.187.328.703 1.593l.515 1.31c.082.235.177.436.29.625l-.009-.016a5.65 5.65 0 0 0 1.867-.34l-.039.013a21.295 21.295 0 0 1-.16-2.649c0-.321.007-.639.02-.957l-.002.045c.025-.366.135-.702.31-.994l-.006.01c.144-.251.341-.457.578-.605l.007-.004l.375-.234c.376-.189.698-.392.999-.621l-.016.011c.261-.169.439-.448.468-.769v-.004a1.003 1.003 0 0 0-.12-.565l.003.005a1.41 1.41 0 0 0-.301-.306l-.004-.003l-.141-.141c.029.264.079.502.148.732l-.007-.029l-.187-.234a6.318 6.318 0 0 1-.266-.608l-.016-.047a6.594 6.594 0 0 1-.276-1.131l-.005-.04a4.932 4.932 0 0 1 .026-1.08l-.003.025a3.55 3.55 0 0 1 .219-.984l-.008.024v.656l.094-.094c.137-.145.248-.316.324-.505l.004-.011a1.094 1.094 0 0 0 .116-.71l.001.007a.566.566 0 0 0-.324-.398l-.004-.001a1.773 1.773 0 0 0-.73-.117h.004a5.233 5.233 0 0 0-1.583.293l.037-.011a10.054 10.054 0 0 0-3.012 1.792l.013-.011l-.282.187l.187-.282a3.91 3.91 0 0 1 .655-.843l.001-.001a6.21 6.21 0 0 1 1.791-1.343l.036-.016a3.708 3.708 0 0 0-1.008-.693l-.023-.01a7.964 7.964 0 0 0-3.476-.703h.009a8.067 8.067 0 0 0-3.518.723l.051-.021a3.104 3.104 0 0 0-1.029.748l-.002.003a5.553 5.553 0 0 1 1.824 1.308l.004.004c.251.244.469.522.646.825l.01.018l.187.282l-.282-.187a9.987 9.987 0 0 0-2.927-1.751l-.069-.023a5.12 5.12 0 0 0-1.539-.281h-.008a1.803 1.803 0 0 0-.738.121l.012-.004a.566.566 0 0 0-.327.396l-.001.004a.994.994 0 0 0 .166.706l-.002-.003q.141.282.282.516l.094.094v-.656c.115.276.19.595.211.929v.008a5.016 5.016 0 0 1 .022 1.097l.001-.019c-.033.43-.134.828-.291 1.196l.009-.025a2.08 2.08 0 0 1-.286.664l.005-.008l-.187.234c.089-.191.142-.414.142-.65l-.001-.056v.003l-.14.133a1.022 1.022 0 0 0-.279.3l-.003.005a1.194 1.194 0 0 0-.094.563v-.004c.022.317.182.593.419.771l.003.002c.286.218.608.42.948.593l.036.017l.047.047q.234.146.326.192zm-1.078 6.798c.262.445.647.793 1.109 1.001l.015.006c.458.218.995.345 1.561.345h.036h-.002a1.39 1.39 0 0 1 .19-1.497l.42-.514a5.248 5.248 0 0 1-2.094-.481l.033.014a2.88 2.88 0 0 1-1.447-1.667l-.006-.02q-.234.234-.187 1.124A3.352 3.352 0 0 0 1.556 19l-.009-.018zm2.16 2.015c.343.53.82.946 1.384 1.207l.02.008a6.603 6.603 0 0 1-.549-1.034l-.017-.044a5.934 5.934 0 0 1-.88-.148l.041.009zm5.198 2.344q2.905 0 3.983-3.046a1.092 1.092 0 0 0-.05-1.059l.003.005a1.465 1.465 0 0 0-1.18-.35l.008-.001H6.141a1.465 1.465 0 0 0-1.173.352l.002-.001a1.181 1.181 0 0 0-.047 1.101l-.003-.007q1.078 3.003 3.983 3.003zm3.797-1.128a3.623 3.623 0 0 0 1.445-1.207l.008-.012a4.42 4.42 0 0 1-.891.094h-.001a6.203 6.203 0 0 1-.575 1.15l.015-.025zm3.561-3.233a3.565 3.565 0 0 0 .421-1.637v-.003q0-.937-.234-1.124a2.688 2.688 0 0 1-1.436 1.633l-.017.007a4.843 4.843 0 0 1-2.011.469h-.004l.375.516c.202.24.324.553.324.894c0 .218-.05.425-.14.609l.004-.008h.018c.572 0 1.114-.129 1.597-.361l-.022.01c.477-.214.862-.56 1.118-.992l.006-.011z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Gulp(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.288 18.319a9.777 9.777 0 0 1-3.89.61l.02.001a10.331 10.331 0 0 1-3.992-.633l.072.023l.703 1.173l.282 3.753q0 .328.868.539a8.742 8.742 0 0 0 1.955.216L5.448 24h-.007l.149.001c.687 0 1.356-.077 1.998-.224l-.06.012q.868-.211.868-.539l.234-3.753zm-8.585-6c.171-.571.411-1.07.716-1.52l-.012.019a3.132 3.132 0 0 1 1.156-1.071l.017-.008a.995.995 0 0 1 .896.049l-.005-.003a.94.94 0 0 1 .326.279l.002.003c.07.076.12.172.14.278l.001.004a.383.383 0 0 1-.025.307l.001-.002a.212.212 0 0 1-.188.118q-.118 0-.305-.328q-.282-.516-.8-.141a2.706 2.706 0 0 0-.793.945l-.007.015c-.215.375-.405.81-.547 1.267l-.013.047a1.447 1.447 0 0 0-.105.549c0 .147.021.288.061.422l-.003-.01a.275.275 0 0 0 .24.142c.297-.059.549-.22.725-.443l.002-.003c.239-.226.448-.48.624-.758l.01-.016l.234-.703q.094-.234.282-.234t.188.258a2.111 2.111 0 0 1-.194.74l.006-.013a6.135 6.135 0 0 0-.271.798l-.011.047l-.378 1.216a1.274 1.274 0 0 1-.121.263l.003-.006a.198.198 0 0 1-.159.118h-.009a.159.159 0 0 1-.132-.07v-.001a.467.467 0 0 1-.023-.285l-.001.003a9.23 9.23 0 0 1 .373-1.264l-.021.064l-.328.328c-.249.28-.594.47-.982.52l-.008.001l-.033.001a.665.665 0 0 1-.345-.096l.003.002l.375 3.705q.047.234 1.219.516c.85.209 1.826.328 2.83.328h.082h-.004h.078c1.004 0 1.98-.12 2.915-.345l-.084.017q1.126-.282 1.219-.516l.516-4.409l-.094.047c-.226.252-.49.463-.785.626l-.015.008a.984.984 0 0 1-.566.07l.006.001q-.285-.047-.285-.282c.098-.335.277-.62.515-.843l.001-.001l.141-.234q.234-.234.234-.352T9 12.389a.71.71 0 0 0-.375.188c-.15.138-.282.292-.394.46l-.006.009a4.416 4.416 0 0 0-.389.699l-.011.028l-.188.375l-.46 1.312a.62.62 0 0 1-.231.21l-.003.002a.213.213 0 0 1-.235 0h.001a.204.204 0 0 1-.091-.164l.047-.141c.122-.354.249-.648.394-.932l-.019.041l.328-.703q.047-.188.047-.141c-.183.2-.401.365-.644.487l-.013.006a.757.757 0 0 1-.474.07l.005.001a.303.303 0 0 1-.234-.232v-.002l-.047-.094c-.197.201-.47.326-.773.328q-.305-.047-.16-.703c-.174.41-.569.694-1.031.703h-.001a.401.401 0 0 1-.421-.232l-.001-.003a1.974 1.974 0 0 1 .23-1.09l-.005.011v-.047q.141-.328.328-.75q.141-.282.375-.188h.047q.141.094.047.375l-.141.188c-.11.205-.22.452-.314.708l-.014.043q-.094.422.07.422a.79.79 0 0 0 .356-.096l-.004.002c.209-.152.376-.349.489-.577l.004-.009c.107-.182.195-.393.254-.616l.004-.018l.094-.24q.094-.234.211-.234a.296.296 0 0 1 .211.093q.094.094 0 .328l-.328.75a1.736 1.736 0 0 0-.14.553v.007q0 .188.328.047c.126-.049.234-.112.33-.19l-.002.002q.047-.188.75-2.017l.469-1.173a.333.333 0 0 1 .318-.235h.011h-.001c.055.01.103.035.141.071q.094.07.047.258l-.847 2.026q-.328.844-.4 1.076t.028.24q.328 0 1.079-.8l.047-.047l.141-.328c.033-.162.098-.305.189-.428l-.002.003q.094-.094.141-.094q.188 0 .188.16v.007a.954.954 0 0 1-.049.304l.002-.007l.234-.234a1.006 1.006 0 0 1 .824-.352h-.003a.38.38 0 0 1 .353.378l-.001.023v-.001a2.256 2.256 0 0 1-.609 1.079l-.094.047a1.024 1.024 0 0 0-.139.274l-.002.007q0 .047.094.047c.129-.034.24-.1.328-.188l.8-.703l.703-6.379a8.93 8.93 0 0 1-3.022.516H7.69h.004a36.826 36.826 0 0 1-4.61-.006l.106.005h-.075A9.084 9.084 0 0 1 .03 6.428l.063.02zm10.179-6.286h-.047q0-.188-.868-.375a15.556 15.556 0 0 0-2.234-.28l-.042-.002l.469-1.92L10.834.594Q10.975.5 10.6.172a.527.527 0 0 0-.278-.159l-.004-.001q-.141-.023-.188.023L7.271 3.028l-.61 2.298l-1.219-.048a27.17 27.17 0 0 0-3.987.25l.14-.016q-1.594.234-1.594.516q0 .469 2.814.657a37.31 37.31 0 0 0 5.375-.007l-.121.007q2.814-.186 2.814-.656zm-3.28.234q0 .047-.188.118a1.238 1.238 0 0 1-.418.071h-.029h.001h-.028c-.15 0-.293-.026-.427-.073l.009.003q-.188-.07-.188-.118t.141-.094l-.047.047q0 .094.516.094t.516-.094l.047-.047q.091.044.091.09z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HackerNews(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h24v24H0zm12.8 13.446l4.339-8.303h-1.871q-2.143 4.018-2.839 5.786l-.375.96l-.32-.75a49.079 49.079 0 0 0-3.022-6.243l.129.243H6.857l4.286 8.2v5.52H12.8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hangout(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.951 0C4.455 0 0 4.455 0 9.951s4.455 9.951 9.951 9.951h.586V24c5.737-2.693 9.366-8.781 9.366-14.049v-.008C19.903 4.452 15.451 0 9.96 0h-.008zm-.585 10.537l-1.17 2.343h-1.76l1.17-2.342h-1.76V7.024h3.512v3.512zm4.683 0L12.88 12.88h-1.76l1.17-2.342h-1.76V7.024h3.512v3.512z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hashtag(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24.01 6.797a1.341 1.341 0 0 0-1.034-.487h-.035h.002h-3.441l.954-4.643A1.4 1.4 0 0 0 20.188.48l.002.003a1.345 1.345 0 0 0-1.036-.486h-.033h.002c-.9.029-1.643.665-1.835 1.51l-.003.013l-.259 1.262c-.008.032-.022.062-.029.094l-.708 3.435h-5.793l.954-4.643a1.4 1.4 0 0 0-.268-1.187l.002.003a1.345 1.345 0 0 0-1.036-.486h-.033h.002c-.9.029-1.643.665-1.835 1.51l-.003.013l-.27 1.313c0 .015-.01.028-.013.042l-.702 3.435H3.547c-.9.029-1.643.664-1.836 1.509l-.003.013a1.4 1.4 0 0 0 .267 1.188l-.002-.003c.248.298.619.487 1.035.487h.034h-.002h3.597l-1.022 4.97H1.864c-.9.029-1.643.665-1.836 1.51l-.003.013a1.404 1.404 0 0 0 .267 1.185L.29 17.18c.249.298.62.487 1.036.487h.034h-.002h3.597l-.96 4.668a1.4 1.4 0 0 0 .268 1.187l-.002-.003c.249.298.62.487 1.036.487h.033h-.002a1.947 1.947 0 0 0 1.836-1.51l.003-.013l.99-4.814h5.796l-.96 4.668a1.4 1.4 0 0 0 .268 1.187l-.002-.003c.249.298.62.486 1.036.486h.033h-.002a1.945 1.945 0 0 0 1.835-1.51l.003-.013l.99-4.814h3.6a1.947 1.947 0 0 0 1.836-1.509l.003-.013a1.388 1.388 0 0 0-.263-1.188l.002.003a1.348 1.348 0 0 0-1.037-.487h-.036h.002h-3.447l1.022-4.97h3.6a1.947 1.947 0 0 0 1.836-1.51l.003-.013a1.401 1.401 0 0 0-.268-1.188l.002.003zM8.819 14.472l1.022-4.97h5.796l-1.022 4.97z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Headphone(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.756 12.075v.019a.646.646 0 0 1-.621.645h-.001a.347.347 0 0 1-.077-.022l.003.001a2.84 2.84 0 0 0-2.346-1.308h-.002a3.037 3.037 0 0 0-2.904 3.163v-.006v6.282a3.028 3.028 0 0 0 2.893 3.15h.151c.194 0 .379-.044.544-.121l-.008.003v.006c.166-.075.332-.155.494-.241a6.082 6.082 0 0 0 1.507-1.191l.004-.004a6.6 6.6 0 0 0 1.535-2.996l.008-.044c.011-.048.027-.102.037-.15c.393-1.877.618-4.035.618-6.244c0-.692-.022-1.378-.066-2.059l.005.093a12.094 12.094 0 0 0-2.997-7.801l.012.014c0-.006-.006-.006-.011-.011a10.491 10.491 0 0 0-1.295-1.194l-.024-.018a9.212 9.212 0 0 0-1.47-.945l-.052-.025a9.393 9.393 0 0 0-4.39-1.072h-.016a9.453 9.453 0 0 0-4.434 1.097l.052-.025c-.581.304-1.079.626-1.545.987l.023-.017c-.488.376-.92.773-1.317 1.204l-.007.007l-.011.011a11.952 11.952 0 0 0-2.996 7.768v.008a31.93 31.93 0 0 0 .593 8.416l-.035-.206c.012.061.025.112.04.161l-.003-.011a6.63 6.63 0 0 0 1.549 3.045l-.005-.006c.431.469.929.866 1.481 1.179l.03.016c.136.08.298.161.466.231l.028.01v-.006c.157.074.341.117.535.118h.145a3.037 3.037 0 0 0 2.899-3.157v.006v-6.271a3.035 3.035 0 0 0-2.898-3.156h-.006a2.824 2.824 0 0 0-2.335 1.297l-.007.011c-.027 0-.054.021-.08.021a.646.646 0 0 1-.622-.645v-.02v.001c0-.016.016-.027.016-.037c.011-.289.027-.584.043-.874c.045-.877.169-1.701.367-2.495l-.019.089a9.194 9.194 0 0 1 1.71-3.611l-.015.021a7.946 7.946 0 0 1 .953-1.029c.085-.08.172-.15.257-.225a8.294 8.294 0 0 1 1.863-1.2l.051-.021a7.734 7.734 0 0 1 3.248-.715h.002c1.18 0 2.299.263 3.301.733l-.048-.02a8.372 8.372 0 0 1 1.925 1.231l-.011-.01c.085.075.172.15.252.225a8.033 8.033 0 0 1 .953 1.029a9.125 9.125 0 0 1 1.676 3.529l.012.062c.181.706.307 1.529.352 2.374l.001.032v.064c0 .285.018.567.052.843l-.003-.033c-.011.006-.011.016-.011.037z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heart(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 6.857l-.001-.097A6.76 6.76 0 0 1 6.863.001h-.005c3.428 0 6 2.572 6.857 3.428c.857-.856 3.428-3.428 6.857-3.428c5.143 0 6.857 3.428 6.857 6.857c0 6.857-12 16.285-13.715 17.143C11.999 23.144-.001 13.716-.001 6.858z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartAlt(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M27.547 7.39c-.298-2.281-1.514-6.24-6.272-7.182a10.32 10.32 0 0 0-1.953-.209h-.153a7.902 7.902 0 0 0-5.378 2.145l.004-.004A7.915 7.915 0 0 0 8.403-.001h-.127c-.7.009-1.376.085-2.03.221l.069-.012C.761 1.299.093 6.791.043 7.416v.021c-.286 2.997.473 10.598 13.472 16.437l.286.126l.278-.126C27.783 17.72 27.853 9.699 27.548 7.39zM13.24 3.521l.55.535l.55-.535a6.819 6.819 0 0 1 4.824-2.141h.125a9.951 9.951 0 0 1 1.781.191l-.063-.011c3.914.772 4.922 4.094 5.178 6.006a9.557 9.557 0 0 1-1.158 5.666l.025-.049c-1.849 3.607-5.638 6.736-11.262 9.309C1.861 17.042 1.145 10.252 1.396 7.569v-.028c.049-.522.598-5.082 5.172-5.986a9.61 9.61 0 0 1 1.7-.181h.141a6.818 6.818 0 0 1 4.829 2.143l.003.003z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartEyes(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.316C6.099 1.316 1.316 6.099 1.316 12S6.099 22.684 12 22.684S22.684 17.901 22.684 12c-.012-5.896-4.788-10.672-10.683-10.684zm0 22.297C5.586 23.613.387 18.414.387 12S5.586.387 12 .387S23.613 5.586 23.613 12v.015c0 6.405-5.192 11.597-11.597 11.597z"/><path fill="currentColor" d="M12 24C5.386 23.966.034 18.614 0 12.003V12C0 5.373 5.373 0 12 0s12 5.373 12 12c-.034 6.614-5.386 11.966-11.997 12zM12 .774C5.8.774.774 5.8.774 12S5.8 23.226 12 23.226S23.226 18.2 23.226 12C23.222 5.802 18.198.779 12.001.774zm0 22.297C5.886 23.071.929 18.114.929 12S5.886.929 12 .929c6.122.021 11.084 4.956 11.148 11.065V12c-.065 6.115-5.027 11.05-11.146 11.071zm0-21.368C6.313 1.703 1.703 6.313 1.703 12S6.313 22.297 12 22.297S22.297 17.687 22.297 12c-.026-5.676-4.62-10.271-10.294-10.297z"/><path fill="currentColor" d="M12 19.665h-.003a7.842 7.842 0 0 1-7.413-5.287l-.016-.055a.441.441 0 0 1 .85-.235l.001.003c.917 2.764 3.479 4.723 6.499 4.723l.086-.001H12h.011a6.933 6.933 0 0 0 6.555-4.674l.015-.049a.442.442 0 1 1 .851.235l.001-.003c-1.075 3.131-3.994 5.342-7.43 5.342H12z"/><path fill="currentColor" d="M12 20.052h-.005a8.23 8.23 0 0 1-7.797-5.594l-.017-.058a.857.857 0 0 1 1.624-.547l.002.006c.868 2.628 3.301 4.49 6.169 4.49h.026h-.001h.076a6.415 6.415 0 0 0 6.105-4.445l.013-.045a.858.858 0 0 1 1.624.548l.002-.006c-1.115 3.31-4.191 5.652-7.814 5.652h-.005zm0-.775h.068a7.321 7.321 0 0 0 6.962-5.058l.015-.052v-.077h-.077c-.992 2.947-3.729 5.032-6.954 5.032H12h.001h-.013a7.324 7.324 0 0 1-6.939-4.981l-.015-.052s-.077-.077-.077 0c0 0-.077.077 0 .077c.967 2.987 3.724 5.11 6.977 5.11h.072h-.004zM9.987 8.052a1.263 1.263 0 0 0-1.857 0l-.001.001l-.387.387l-.387-.387A1.316 1.316 0 1 0 5.496 9.91l.001.001l2.245 2.245l2.245-2.245a1.26 1.26 0 0 0 .001-1.857zm8.516 0a1.263 1.263 0 0 0-1.857 0l-.001.001l-.387.387l-.387-.387a1.316 1.316 0 1 0-1.859 1.857l.001.001l2.245 2.245l2.245-2.245a1.41 1.41 0 0 0-.002-1.86z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Heartbeat(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.702-.001H4.34a4.343 4.343 0 0 0-4.341 4.34v11.964a4.344 4.344 0 0 0 4.341 4.342h3.44v1.413h-4V24h15.493v-1.942h-4v-1.413h3.434a4.344 4.344 0 0 0 4.341-4.341V4.338A4.338 4.338 0 0 0 18.71 0h-.007zm1.449 16.312v.001c0 .8-.649 1.449-1.449 1.449H4.339c-.8 0-1.448-.648-1.448-1.448v-4.98h6.856l.622-2.03l2.88 6.54l1.305-4.544h5.594v5.014zm0-6.457h-6.674l-.507 1.774l-2.834-6.435l-1.449 4.695H2.89V4.345c0-.8.649-1.449 1.449-1.45h14.362c.8 0 1.449.649 1.449 1.449v.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HeartbeatAlt(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.909 10.928H24v2.02h-2.091zm-3.635 0h2.091v2.02h-2.091zm-6.022.001l-.657 2.215l-3.772-8.343l-1.954 6.17H0v2.02h7.346l.81-2.551l3.834 8.486l1.76-5.978h2.973v-2.02z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Helicopter(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.267 23.498a5.106 5.106 0 0 1-.044-1.068l-.001.015l.017-.844l1.07-.034l.502-.761c.217-.34.438-.634.678-.911l-.008.009l.16-.135h5.054c4.89 0 5.065.006 5.194.107c.235.266.456.559.654.868l.017.028l.535.794l1.07.034l.017.704c.028 1.026-.017 1.28-.24 1.504a.488.488 0 0 1-.448.185h.002a.496.496 0 0 1-.449-.184l-.001-.001c-.169-.16-1.746-2.434-1.82-2.62a42.085 42.085 0 0 0-2.911-.098c-.565 0-1.128.011-1.688.032l.081-.002l-4.485.006l-.872 1.296a17.106 17.106 0 0 1-1.058 1.474l.026-.034a.588.588 0 0 1-.38.138H3.93a.702.702 0 0 1-.663-.497zm7.42-5.093l-4.607-.017l-.897-.902l-.902-.902v-2.425c0-1.33.028-2.665.056-2.96a6.423 6.423 0 0 1 3.483-5.07l.037-.017c.46-.242.994-.433 1.556-.546l.038-.006a7.736 7.736 0 0 1 2.896.089l-.051-.01a6.556 6.556 0 0 1 2.886 1.646l-.001-.001a6.392 6.392 0 0 1 1.802 3.311l.007.041a21.914 21.914 0 0 1 .117 3.216l.001-.038l.028 2.755l-.924.924l-.924.918zm-.89-2.56a.701.701 0 0 0-.238.944l-.002-.003c.174.293.333.338 1.172.338h.766l.203-.186a.676.676 0 0 0-.104-1.108l-.003-.002c-.152-.096-.287-.113-.89-.113c-.65-.004-.729.007-.904.125zM6 10.05a5.343 5.343 0 0 0-.319 2.225l-.001-.016l-.006 1.04l.907.913l.911.909l.428-.011c.422-.006.422-.006.687-.265l.259-.259h3.662l.265.259l.259.259h.88l1.786-1.786v-1.03a5.27 5.27 0 0 0-.38-2.396l.013.035a5.033 5.033 0 0 0-3.704-3.035l-.031-.005a4.698 4.698 0 0 0-.902-.086h-.004a5.083 5.083 0 0 0-4.699 3.215l-.012.034zm3.431-5.916l.13-.874c.067-.48.13-.88.13-.89s-2.18-.011-4.845-.011H0V1.007h9.69V.866a1.08 1.08 0 0 1 .613-.808l.007-.003a1.296 1.296 0 0 1 .815.02l-.009-.003c.287.139.5.394.578.702l.002.007l.039.226h9.674v1.352h-9.696l.028.18c.074.434.203 1.409.203 1.498s-.107.096-1.256.096z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HelicopterAmbulance(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.861 12.033h-.014a2.25 2.25 0 1 1 1.928-3.407l.006.01h.684a2.86 2.86 0 0 0-2.614-1.707a2.854 2.854 0 1 0 2.4 4.399l.007-.011l-.617-.173a2.267 2.267 0 0 1-1.78.888H2.86z"/><path fill="currentColor" d="M23.552 12.254c-1.746-2.593-4.671-4.276-7.988-4.276l-.157.001h.008h-.066l-.178-.739l-.183-.011l-.32.017l-.1.772a3.918 3.918 0 0 0-2.484.987l.004-.004h-9.6l.622 1.334l7.6 2.14a4.04 4.04 0 0 0 3.367 3.757l.023.003l-.278.806h-2a.428.428 0 0 0 0 .856h10.32a.428.428 0 0 0 0-.856h-1.595l-.394-.806c2.012-.19 3.218-.751 3.68-1.718c.108-.26.171-.563.171-.88c0-.519-.168-.998-.453-1.387l.005.007zm-9.514.461v.939h-1.106v-.939H12v-1.106h.939v-.939h1.106v.939h.939v1.106zm.5 4.325l.259-.761h4.602l.373.761zm8.854-3.952a.671.671 0 0 1-.595.362h-3.6a3.784 3.784 0 0 1-3.741-3.783v-.031v.002v-.145c0-.106.006-.211.011-.32a.956.956 0 0 1 .951-.88h.03a.85.85 0 0 1 .088.005h-.003a9.105 9.105 0 0 1 6.779 4.068l.021.034a.64.64 0 0 1 .056.691z"/><path fill="currentColor" d="M21.918 7.742a.44.44 0 0 0 .495-.433V6.033a.42.42 0 0 0-.144-.316h-.001a.446.446 0 0 0-.341-.112h.002l-7.02.767l-7.019-.767a.446.446 0 0 0-.339.112a.422.422 0 0 0-.144.319V7.31a.441.441 0 0 0 .498.432h-.002l7.006-.951z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hexo(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.435 0L0 6.052v12l10.435 5.953l10.435-6.052V5.948zm4.803 17.109l-.939.521l-.939-.521V12.94H7.515v4.174l-.94.521l-.94-.521V6.887l.939-.521l.939.521v4.174h5.843V6.887l.94-.521l.938.521z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hipchat(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23.299 22.536a5.022 5.022 0 0 1-2.152-2.576l-.011-.035a.48.48 0 0 1 .189-.567l.002-.001c2.586-1.873 4.252-4.879 4.265-8.275v-.002C25.592 4.967 19.863 0 12.796 0C5.729.064 0 5.029 0 11.077c0 6.11 5.729 11.076 12.796 11.076l.148.001c.894 0 1.767-.094 2.609-.271l-.082.014a.49.49 0 0 1 .509.128a13.482 13.482 0 0 0 6.599 1.974l.023.001a.905.905 0 0 0 1.016-.89c.128-.193-.129-.385-.32-.574zm-3.439-1.018a.19.19 0 0 1 .065.125v.001a.136.136 0 0 1-.128.128a10.239 10.239 0 0 1-3.714-1.924l.018.015a.643.643 0 0 0-.644-.128l.005-.001a11.83 11.83 0 0 1-2.669.32h-.005c-5.919 0-10.695-3.945-10.695-8.851s4.776-8.851 10.695-8.851s10.696 3.945 10.696 8.851a8.322 8.322 0 0 1-4.289 7.045l-.043.022a.737.737 0 0 0-.383.573v.003a6.362 6.362 0 0 0 1.098 2.692l-.013-.02zm-.063-7.831v-.011a.5.5 0 0 0-.501-.501h-.007a.49.49 0 0 0-.32.128a9.362 9.362 0 0 1-6.008 2.167l-.174-.002h.009a9.995 9.995 0 0 1-6.194-2.191l.02.016a.666.666 0 0 0-.317-.128h-.003a.52.52 0 0 0-.508.577v-.002a.863.863 0 0 0 .318.624l.002.001a8.46 8.46 0 0 0 6.513 3.057l.178-.002h-.009h.065a8.098 8.098 0 0 0 6.672-3.103l.013-.018c.064-.192.252-.448.252-.64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func History(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23.998 11.978v.023C23.998 18.628 18.626 24 11.999 24a11.947 11.947 0 0 1-7.553-2.675l.022.018a1.162 1.162 0 0 1-.089-1.724l.545-.545a1.162 1.162 0 0 1 1.546-.094l-.002-.002a8.848 8.848 0 0 0 5.53 1.925A8.903 8.903 0 1 0 5.896 5.517l.005-.004l2.456 2.456a.774.774 0 0 1-.548 1.321H.775a.774.774 0 0 1-.774-.774V1.482A.775.775 0 0 1 1.322.934l2.389 2.389A11.955 11.955 0 0 1 12 0c6.619 0 11.987 5.36 11.999 11.976v.001zm-8.753 3.811l.475-.611a1.158 1.158 0 0 0-.202-1.628l-.003-.002l-1.969-1.532V6.968c0-.641-.52-1.162-1.162-1.162h-.774c-.641 0-1.162.52-1.162 1.162v6.56l3.165 2.461a1.162 1.162 0 0 0 1.63-.197z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HolidayVillage(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.912 18.228h2.091v2.091h-2.091zm7.228 0h2.091v2.091H28.14zm-1.754-7.416v-2.09h-2.898V9.74l1.081 1.073z"/><path fill="currentColor" d="M33.552 22.112v-4.409h.665l.618-.611l-7.322-7.322h-.909v1.269h-2.13L23.205 9.77h-3.516l-7.314 7.314l.618.618h.979v4.409h-1.911C9.203 19.377 8.443 13.176 8.71 8.563l.047-.094h.023c1.25 0 2.358.606 3.048 1.54l.007.01c.52.703.833 1.588.833 2.545c0 .66-.148 1.285-.414 1.844l.011-.026a3.842 3.842 0 0 0 1.518-.78l-.006.005a4.254 4.254 0 0 0 1.461-3.213c0-.955-.314-1.836-.844-2.546l.008.011A3.71 3.71 0 0 0 10.28 6.48l.026-.007c.265-.165.57-.311.891-.422l.032-.01a3.891 3.891 0 0 1 1.363-.241c1.083 0 2.064.433 2.781 1.134l-.001-.001a3.516 3.516 0 0 0-.093-1.685l.007.025a3.786 3.786 0 0 0-4.991-2.293l.025-.009a5.067 5.067 0 0 0-.702.303l.029-.014a4.551 4.551 0 0 0-.301-.909l.012.029C8.478.407 6.46-.525 4.847.29a2.989 2.989 0 0 0-1.09.95l-.006.01A3.664 3.664 0 0 1 6.53 3.003l.009.016a5.048 5.048 0 0 0-1.179-.417l-.034-.006C2.846 2.063.486 3.41.063 5.611a3.74 3.74 0 0 0 .109 1.812l-.007-.026a4.553 4.553 0 0 1 4.543-1.402l-.032-.007c.409.086.771.21 1.11.373l-.029-.013a4.114 4.114 0 0 0-4.117 2.04l-.011.02a4.59 4.59 0 0 0-.643 2.357c0 1.568.777 2.954 1.968 3.794l.015.01a4.028 4.028 0 0 0 1.709.648l.021.002a4.665 4.665 0 0 1-.671-2.423c0-.873.237-1.691.651-2.392l-.012.022a4.204 4.204 0 0 1 2.927-2.048l.025-.004c-1.417 6.453-1.809 10.949-.64 13.736H3.573v1.88h32.148v-1.88h-2.169zm-7.95-4.307h7.166v4.307h-7.166zm-10.861-.854l4.95-4.95l4.95 4.95v5.161h-5.435v-3.884h-2.749v3.884h-1.722v-5.161z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Home(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24 12h-3v12h-6v-9H9v9H3V12H0L12 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Horizon(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.976 16.585a.866.866 0 0 1-.224-.034l.006.002a.658.658 0 0 1-.062-.019l.003.001a.98.98 0 0 1-.209-.097l.004.002l-.022-.017a.854.854 0 0 1-.155-.137l-.001-.001l-.04-.045a.899.899 0 0 1-.118-.193l-.002-.006c0-.006-.006-.009-.009-.014c-.009-.021-.012-.042-.018-.062a.585.585 0 0 1-.025-.08l-.001-.006a1.003 1.003 0 0 1-.015-.095v-.004c0-.018-.007-.034-.008-.053v-.059a7.257 7.257 0 1 1 14.505-.342c0 .114 0 .225-.01.336v.064c0 .012-.005.027-.006.041a.89.89 0 0 1-.019.122l.001-.006a.924.924 0 0 1-.032.104l.002-.007c-.006.016-.008.032-.014.048l-.008.014a.815.815 0 0 1-.082.146l.002-.003v.005a.88.88 0 0 1-.554.374l-.006.001h-.016a.912.912 0 0 1-.113.012h-.109a.543.543 0 0 1-.062-.009l.003.001a.69.69 0 0 1-.096-.015l.004.001a.945.945 0 0 1-.105-.032l.006.002c-.017-.006-.034-.009-.05-.015c-1.722-.709-3.721-1.121-5.817-1.121l-.179.001h.009l-.17-.001a15.61 15.61 0 0 0-5.921 1.159l.104-.038c-.018.008-.038.011-.057.018s-.058.02-.088.027a1.077 1.077 0 0 1-.101.016h-.005c-.017 0-.033.006-.05.008h-.052zm1.006-2.231c1.574-.492 3.385-.776 5.261-.776h.088h-.005h.082c1.877 0 3.687.283 5.391.81l-.129-.034c-.475-2.555-2.687-4.464-5.344-4.464s-4.87 1.909-5.339 4.431l-.005.034zm15.99 1.886a.91.91 0 0 1 0-1.82h3.77a.91.91 0 0 1 0 1.82zm-25.062 0a.91.91 0 0 1 0-1.82h3.77a.91.91 0 0 1 0 1.82zm21.3-7.794a.91.91 0 0 1 0-1.288l2.666-2.666a.91.91 0 1 1 1.285 1.283l-.001.001l-2.666 2.666a.907.907 0 0 1-1.288 0zm-15.056 0L4.489 5.777a.91.91 0 1 1 1.287-1.288l2.666 2.666a.908.908 0 0 1-.629 1.567a.911.911 0 0 1-.657-.28zm7.261-3.76V.911a.91.91 0 0 1 1.82 0v3.77a.91.91 0 0 1-1.82 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HorizonAlt(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.658 19.847a7.254 7.254 0 1 1 11.2.18l.009-.01l-.017.016a.982.982 0 0 1-.094.089l-.002.001c-.016.013-.03.028-.047.04a.87.87 0 0 1-.116.07l-.005.002c-.015.008-.03.018-.045.025a.832.832 0 0 1-.162.054l-.006.001c-.024.006-.049.006-.072.01a1.105 1.105 0 0 1-.11.011h-.02a1.101 1.101 0 0 1-.131-.011l.005.001h-.022a.915.915 0 0 1-.177-.048l.006.002h-.005a12.505 12.505 0 0 0-4.52-.829c-1.63 0-3.189.306-4.622.864l.087-.03l-.026.007l-.036.01a.847.847 0 0 1-.263.044h-.001a.909.909 0 0 1-.807-.495l-.002-.005zm.234-4.52a5.417 5.417 0 0 0 .933 3.038l-.012-.019c1.349-.455 2.902-.718 4.516-.718s3.167.263 4.619.748l-.103-.03a5.434 5.434 0 1 0-9.951-3.021v.001zm16.08.91a.91.91 0 0 1 0-1.82h3.77a.91.91 0 0 1 0 1.82zm-25.062 0a.91.91 0 0 1 0-1.82h3.772a.91.91 0 0 1 0 1.82zm21.3-7.795a.907.907 0 0 1 0-1.288l2.666-2.666a.911.911 0 0 1 1.288 1.288l-2.665 2.666a.909.909 0 0 1-1.288 0zm-15.056 0L4.492 5.776a.91.91 0 1 1 1.287-1.288l2.666 2.666a.908.908 0 0 1-.629 1.567a.911.911 0 0 1-.657-.28zm7.261-3.76V.91a.91.91 0 0 1 1.82 0v3.77a.91.91 0 0 1-1.82 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hospital(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.962 0H2.036A2.037 2.037 0 0 0 0 2.036v19.928C0 23.088.912 24 2.036 24h19.929A2.035 2.035 0 0 0 24 21.965V2.037A2.036 2.036 0 0 0 21.964.002h-.001zM19.84 18.763h-4.29v-4.66H8.469v4.66h-4.32V5.219h4.32v4.66h7.081v-4.66h4.29z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HotAirBalloon(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.97 2.74a9.36 9.36 0 1 0 2.74 6.617v-.03a9.259 9.259 0 0 0-2.74-6.586zM2.903 9.356c0-4.474 2.781-8.193 6.452-8.988c3.671.795 6.452 4.521 6.452 8.988s-2.781 8.193-6.452 8.988c-3.675-.796-6.452-4.515-6.452-8.989z"/><path fill="currentColor" d="m12.225 9.355l.001-.185c0-3.193-1.077-6.135-2.888-8.481l.024.032C7.575 3.099 6.5 6.102 6.5 9.355s1.075 6.256 2.889 8.672l-.027-.037a13.785 13.785 0 0 0 2.864-8.444l-.001-.201zm-.021 9.675h-.739a.389.389 0 0 0-.388.373v1.041h-3.44v-1.04a.388.388 0 0 0-.388-.374h-.013h.001h-.744a.46.46 0 0 0-.459.449v.001l.326 4.072c.011.25.216.449.468.449h5.074c.25 0 .454-.2.459-.449v-.001l.326-4.072a.459.459 0 0 0-.459-.449l-.025.001h.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hotel(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M28.895 24H2.384A2.4 2.4 0 0 1 0 21.617V6.739a2.4 2.4 0 0 1 2.383-2.384h26.512a2.4 2.4 0 0 1 2.384 2.383v14.881A2.381 2.381 0 0 1 28.898 24zM2.384 5.945h-.007a.787.787 0 0 0-.787.787v14.895c0 .435.352.787.787.787h26.526a.787.787 0 0 0 .787-.787V6.732a.787.787 0 0 0-.787-.787h-.008zm20.281-4.038v.004a.854.854 0 0 1-.854.854H9.47a.854.854 0 0 1-.854-.854v-.008c0-.472.383-.854.854-.854h12.337a.898.898 0 0 1 .858.857zm2.988.636v.014c0 .712-.577 1.289-1.289 1.289h-.015h.001h-.014a1.289 1.289 0 0 1-1.289-1.289v-.015v.001v-1.254c0-.712.577-1.289 1.289-1.289h.015h-.001h.014c.712 0 1.289.577 1.289 1.289v.015v-.001zm-17.42 0v.014c0 .712-.577 1.289-1.289 1.289h-.015h.001h-.014a1.289 1.289 0 0 1-1.289-1.289v-.015v.001v-1.254C5.627.577 6.204 0 6.916 0h.015h-.001h.014c.712 0 1.289.577 1.289 1.289v.015v-.001zm21.585-.636v.004a.854.854 0 0 1-.854.854h-2.137a.854.854 0 0 1-.854-.854v-.008c0-.472.383-.854.854-.854h2.133a.898.898 0 0 1 .858.857zm-24.541 0v.004a.854.854 0 0 1-.854.854H2.317a.854.854 0 0 1-.854-.854v-.008c0-.472.383-.854.854-.854H4.45a.869.869 0 0 1 .826.858z"/><path fill="currentColor" d="M5.658 14.558v2h1.939v-2h1.176v5.214H7.597v-2.194H5.658v2.194H4.482v-5.214zm8.933 2.545a2.508 2.508 0 0 1-2.495 2.767l-.085-.001h.004h-.029a2.43 2.43 0 0 1-2.418-2.681l-.001.01a2.523 2.523 0 0 1 2.514-2.735h.031h-.002l.061-.001a2.427 2.427 0 0 1 2.417 2.65l.001-.009zm-3.751.063c0 1.018.48 1.748 1.271 1.748s1.24-.763 1.24-1.78c0-.954-.445-1.748-1.271-1.748c-.794-.002-1.24.762-1.24 1.778zm5.34-1.622h-1.398v-.986h4.006v.986h-1.43v4.228h-1.176zm6.326 2.002h-1.907v1.24h2.16v.954h-3.368v-5.182h3.21v.954h-2.034v1.081h1.907v.954zm1.049-2.988h1.176v4.228h2.066v.986h-3.242zm-7.915-3.336l-1.717.925l.32-1.907l-1.369-1.373l1.907-.286l.858-1.714l.858 1.717l1.907.286l-1.367 1.367l.32 1.907zm5.944.221l-1.303.699l.254-1.462l-1.05-1.017l1.43-.19l.667-1.335l.636 1.335l1.462.19l-1.049 1.018l.222 1.462zm3.782.255l-.858.445l.16-.954l-.699-.667l.954-.16l.445-.858l.414.858l.954.16l-.667.667l.16.954zm-15.639-.255l1.271.699l-.254-1.462l1.049-1.018l-1.43-.19l-.635-1.334l-.667 1.335l-1.431.19l1.049 1.018l-.254 1.462zm-3.814.255l.858.445l-.16-.954l.699-.667l-.954-.16l-.444-.858l-.414.858l-.954.16l.699.667l-.16.954z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HotelAlt(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.638 22.405v.001c0 .881-.714 1.595-1.595 1.595h-.006a1.594 1.594 0 0 1-1.594-1.594V1.595c0-.881.714-1.594 1.594-1.594h.006c.881 0 1.595.714 1.595 1.595zM2.014 2.686H.899A.903.903 0 0 0 .897 4.49h1.116zm23.59 3.264h-16A2.658 2.658 0 0 0 6.95 8.604v9.248a2.652 2.652 0 0 0 2.652 2.652h16.002a2.658 2.658 0 0 0 2.654-2.654V8.602a2.652 2.652 0 0 0-2.652-2.652zm1.28 11.9a1.28 1.28 0 0 1-1.275 1.275H9.607c-.703 0-1.273-.57-1.273-1.273V8.603a1.28 1.28 0 0 1 1.274-1.275h16a1.28 1.28 0 0 1 1.275 1.275zM12.101 5.437a.93.93 0 0 0 .93-.93V2.62a.93.93 0 0 0-1.86 0v.029v-.001v1.86a.93.93 0 0 0 .93.93zm11.009 0a.93.93 0 0 0 .93-.93v-1.86a.93.93 0 0 0-1.86 0v1.86c0 .514.417.93.93.93m-9.623-2.751h8.235V4.49h-8.235zm14.09.891a.892.892 0 0 0-.89-.89h-2.061v1.804h2.06a.893.893 0 0 0 .89-.89z"/><path fill="currentColor" d="M12.236 15.011h-.97v-.898h-.89v2.574h.89v-1.04h.97v1.04h.89v-2.574h-.89zm3.962-.601a1.887 1.887 0 0 0-1.093-.345c-.414 0-.797.132-1.11.357l.006-.004a1.233 1.233 0 0 0-.395.905l.002.077v-.003l-.001.057c0 .258.075.499.205.702l-.003-.005c.125.193.299.344.505.438l.007.003c.213.087.461.138.72.138l.084-.002h-.004l.051.001c.271 0 .529-.059.761-.165l-.011.005c.207-.099.375-.253.487-.444l.003-.005a1.45 1.45 0 0 0 .169-.682l-.001-.058v.003l.002-.074c0-.352-.147-.67-.383-.895zm-.657 1.556a.57.57 0 0 1-.405.169l-.038-.001h.002l-.027.001a.589.589 0 0 1-.414-.169a1.054 1.054 0 0 1 .002-1.124l-.003.004c.115-.105.268-.168.437-.168s.322.064.437.169h-.001a.715.715 0 0 1 .16.533v-.003a.871.871 0 0 1-.153.594l.002-.003zm1.21-1.22h.906v1.941h.89v-1.941h.906v-.634h-2.702zm3.994.851h1.38v-.53h-1.38v-.409h1.484v-.546h-2.374v2.574h2.422V16.1h-1.532zm2.894-1.484h-.89v2.574h2.278v-.634H23.64zm-13.03-2.085l.834-.442l.842.442l-.16-.93l.674-.658l-.93-.136l-.425-.85l-.417.85l-.93.136l.674.658zm3.087 0l.834-.442l.834.442l-.16-.93l.673-.658l-.93-.136l-.417-.85l-.417.85l-.938.136l.674.658zm3.079 0l.834-.442l.834.442l-.16-.93l.674-.658l-.93-.136l-.417-.85l-.417.85l-.938.136l.682.658zm3.079 0l.834-.442l.834.442l-.16-.93l.674-.658l-.93-.136l-.417-.85l-.417.85l-.938.136l.682.658zm3.912-2.574l-.417.85l-.93.136l.673.658l-.16.93l.834-.442l.834.442l-.16-.93l.682-.658l-.938-.136zM6.142 2.686h4.45V4.49h-4.45z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Hourglass(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.304 22a11.077 11.077 0 0 0-5.645-9.971L14.604 12c3.427-1.929 5.704-5.543 5.704-9.689c0-.109-.002-.218-.005-.327V2a.923.923 0 0 0 .899-1.003V1a.923.923 0 0 0-.898-1H.903a.923.923 0 0 0-.901 1.003V1a.923.923 0 0 0 .9 1h.001a11.077 11.077 0 0 0 5.645 9.971l.055.029C3.176 13.929.899 17.543.899 21.689c0 .109.002.218.005.327V22a1.007 1.007 0 0 0-.005 2h19.504a.923.923 0 0 0 .9-1.003V23a.945.945 0 0 0-1.003-1zM2.704 2h15.9c0 5-3.5 9-8 9s-7.9-4-7.9-9m0 20c0-5 3.5-9 8-9s8 4 8 9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HourglassEnd(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.403 22a11.078 11.078 0 0 0-5.644-9.971L14.704 12c3.427-1.929 5.704-5.543 5.704-9.689c0-.11-.002-.219-.005-.327V2a.923.923 0 0 0 .9-1.003V1a.923.923 0 0 0-.899-1H.904a.923.923 0 0 0-.901 1.003V1a.923.923 0 0 0 .9 1h.001a11.077 11.077 0 0 0 5.645 9.971l.055.029C3.177 13.929.9 17.543.9 21.689c0 .109.002.218.005.327V22A1.007 1.007 0 0 0 .9 24h19.504a1.007 1.007 0 0 0 .005-2h-.005zM2.704 2h15.9c0 5-3.5 9-8 9s-7.9-4-7.9-9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HourglassHalf(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.304 22a11.077 11.077 0 0 0-5.645-9.971L14.604 12c3.427-1.929 5.704-5.543 5.704-9.689c0-.109-.002-.218-.005-.327V2a.923.923 0 0 0 .899-1.003V1a.923.923 0 0 0-.898-1H.903a.923.923 0 0 0-.901 1.003V1a.923.923 0 0 0 .9 1h.001a11.077 11.077 0 0 0 5.645 9.971l.055.029C3.176 13.929.899 17.543.899 21.689c0 .109.002.218.005.327V22a1.007 1.007 0 0 0-.005 2h19.504a.923.923 0 0 0 .9-1.003V23a1.072 1.072 0 0 0-.996-1zm-1.8-20l.001.14c0 1.44-.333 2.802-.926 4.014l.024-.054H3.504a9.005 9.005 0 0 1-.901-3.958l.001-.149V2zm-15.9 20v-.042c0-1.429.294-2.79.825-4.024L3.404 18h14.199c.502 1.182.795 2.556.8 3.998V22z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HourglassStart(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.304 22a11.077 11.077 0 0 0-5.645-9.971L14.604 12c3.427-1.929 5.704-5.543 5.704-9.689c0-.109-.002-.218-.005-.327V2a.923.923 0 0 0 .899-1.003V1a.923.923 0 0 0-.898-1H.903a.923.923 0 0 0-.901 1.003V1a.923.923 0 0 0 .9 1h.001a11.077 11.077 0 0 0 5.645 9.971l.055.029C3.176 13.929.899 17.543.899 21.689c0 .109.002.218.005.327V22a1.007 1.007 0 0 0-.005 2h19.503a.923.923 0 0 0 .9-1.003V23a1 1 0 0 0-.999-1zm-17.7 0c0-5 3.5-9 8-9s8 4 8 9z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Houzz(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.926 15.991L13.853 12v7.995L6.926 24zM0 12v7.995l6.926-4.005zM6.926 0v7.995L0 12V4.005zm0 7.995l6.926-3.991V12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func HtmlFive(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16.95 7.035l.24-2.625H3.929l.705 8.01h9.18l-.33 3.42l-2.955.795l-2.94-.795l-.195-2.1H4.769l.33 4.17l5.43 1.5h.06v-.015l5.385-1.485l.75-8.16h-9.66l-.225-2.715zM0 0h21.12L19.2 21.57L10.53 24l-8.61-2.43z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Icq(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.209.014a3.856 3.856 0 0 0-2.22.481l.019-.01a3.846 3.846 0 0 0-1.923 2.274l-.007.027v.068l-.035.237l-.135-.237l-.338-.606A3.518 3.518 0 0 0 8.532.663L8.507.657a2.682 2.682 0 0 0-2.008.249l.014-.007A2.846 2.846 0 0 0 5.228 2.95l-.001.013a4.329 4.329 0 0 0 .516 2.894l-.011-.021l.238.339l-.815-.269a3.856 3.856 0 0 0-.946-.114c-.722 0-1.399.192-1.983.528l.019-.01A4.09 4.09 0 0 0 .262 8.482l-.009.028a2.974 2.974 0 0 0-.256 1.218c0 .546.145 1.059.398 1.502l-.008-.015c.202.35.462.643.772.875l.007.005c.401.391.924.659 1.507.741l.014.002l.236.067l-.505.442c-.42.372-.738.852-.908 1.397l-.006.022a2.233 2.233 0 0 0 .244 1.803l-.006-.01c.151.273.354.5.599.672l.006.004c.271.245.622.405 1.009.44h.007a4.364 4.364 0 0 0 2.398-.109l-.031.009l-.606 1.009l-.17.44a3.66 3.66 0 0 0-.197 1.198c0 .666.174 1.291.479 1.832l-.01-.019c.247.381.502.713.783 1.022l-.006-.006a3.59 3.59 0 0 0 1.499.772l.025.005c.35.109.752.172 1.168.172c.671 0 1.304-.163 1.862-.452l-.023.011a3.806 3.806 0 0 0 1.856-2.31l.006-.027l.032-.202a4.07 4.07 0 0 0 1.901 1.344l.029.008a2.224 2.224 0 0 0 1.971-.24l-.008.005a2.698 2.698 0 0 0 1.181-1.606l.004-.019a5.416 5.416 0 0 0-.007-2.37l.007.036c.308.224.669.398 1.06.5l.022.005c.376.143.81.226 1.263.226c.627 0 1.218-.158 1.733-.437l-.019.01a3.853 3.853 0 0 0 1.821-2.342l.006-.027c.101-.33.16-.71.16-1.104c0-.674-.171-1.307-.473-1.86l.01.02l-.172-.273a3.659 3.659 0 0 0-.94-1.076l-.009-.006a6.678 6.678 0 0 0-1.065-.558l-.046-.017a3.99 3.99 0 0 0 2.12-2.106l.01-.026a2.295 2.295 0 0 0-.181-1.971l.006.011v-.063a2.854 2.854 0 0 0-1.772-1.346l-.02-.004a6.255 6.255 0 0 0-1.613-.067l.024-.002l-1.118.237l.373-.846a3.67 3.67 0 0 0 .195-1.19c0-.71-.198-1.374-.542-1.939l.009.016l-.101-.172A3.543 3.543 0 0 0 16.004.152L15.98.148a4.32 4.32 0 0 0-.758-.133l-.017-.001zm-.575 1.427a2.198 2.198 0 0 1 .991.137l-.015-.005a2.206 2.206 0 0 1 1.518 1.099l.006.011c.17.339.27.739.27 1.162v.022v-.001l-.101.745c-.139.386-.31.719-.515 1.026l.011-.017l-3.044 3.724a5.55 5.55 0 0 0-.947-.561l-.035-.014l-.339-3.825l-.033-1.352l.032-.338a2.456 2.456 0 0 1 1.238-1.516l.014-.006c.274-.159.598-.267.943-.302l.01-.001zm-6.947.542h.003c.15 0 .292.037.416.103l-.005-.002a2.451 2.451 0 0 1 1.516 1.433l.006.017c.18.312.319.675.4 1.059l.004.024l.676 3.994l-.505.17l-3.012-3.282l-.17-.337a3.372 3.372 0 0 1-.371-1.804l-.001.011v-.037c0-.51.268-.957.671-1.209l.006-.003a.575.575 0 0 1 .363-.131zm12.243 5.08h.012c.155 0 .308.009.458.026l-.018-.002c.511.028.948.319 1.18.74l.004.007l.136.437l-.068.338a2.737 2.737 0 0 1-1.568 1.449l-.019.006a3.677 3.677 0 0 1-1.092.301l-.018.002l-4.161.505l-.068-.237l-.069-.101l.136-.169l3.418-2.842l.404-.168c.394-.182.854-.29 1.339-.293h.001zm-15.605.185c.18.01.349.037.512.08l-.018-.004l.916.37l2.977 2.169a2.935 2.935 0 0 0-.535.997l-.006.021l-4.873.606l-.337-.067a1.647 1.647 0 0 1-1.279-.87l-.004-.009a1.711 1.711 0 0 1-.062-1.534l-.004.011a2.392 2.392 0 0 1 1.299-1.414l.015-.006a2.502 2.502 0 0 1 1.406-.349h-.006zm7.226 2.715h.027a2.23 2.23 0 0 1 1.532.607l-.001-.001c.419.39.681.944.681 1.559v.037v-.002l.002.08c0 .599-.261 1.138-.676 1.507l-.002.002c-.39.417-.943.676-1.557.676h-.001l-.046.001a2.09 2.09 0 0 1-1.54-.675l-.001-.001a2.135 2.135 0 0 1-.676-1.562v-.029v.001a2.254 2.254 0 0 1 .676-1.59a2.293 2.293 0 0 1 1.557-.606h.033h-.002zm8.497 2.604l.339.035a2.13 2.13 0 0 1 1.246.969l.005.01c.315.368.506.849.506 1.375l-.002.086v-.004l-.101.676a2.122 2.122 0 0 1-1.138 1.448l-.012.006a2.196 2.196 0 0 1-1.91.167l.015.005a1.663 1.663 0 0 1-.907-.539l-.002-.002l-3.553-2.945l.237-.505l.101-.404l3.823-.338l1.352-.035zm-6.698 2.47l1.42 1.453l.878 1.083l.172.34a2.92 2.92 0 0 1 .402 1.769l.001-.011l.001.048c0 .497-.27.932-.671 1.165l-.006.003a1.025 1.025 0 0 1-.784.134l.007.001a2.435 2.435 0 0 1-1.517-1.47l-.005-.017l-.404-1.118l-.505-2.944l.573-.202l.44-.237zm-3.654.069c.302.179.656.309 1.033.368l.017.002c.067 1.151.167 2.369.335 3.688v1.354l-.033.336a2.435 2.435 0 0 1-1.237 1.381l-.015.006a2.428 2.428 0 0 1-1.879.233l.017.004a2.172 2.172 0 0 1-1.448-1.104l-.006-.012a2.59 2.59 0 0 1-.303-1.147v-.004l.172-.777c.076-.373.255-.696.505-.949l2.842-3.381z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func IfQuestionCircle(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.086 6.496h-.035a3.83 3.83 0 0 0-2.512.935l.005-.004a2.959 2.959 0 0 0-1.191 2.461v-.004h2.073a1.319 1.319 0 0 1 .5-1.135l.003-.002c.282-.27.665-.437 1.087-.437h.001l.075-.002c.383 0 .731.153.985.402c.255.257.412.611.412 1.001l-.001.046v-.002l.001.043c0 .392-.15.75-.396 1.017l.001-.001l-1.247 1.251a2.474 2.474 0 0 0-.545.863l-.006.017c-.109.222 0 .56 0 1.018v.861h1.412v-.589c.009-.407.162-.777.411-1.061l-.002.002c.08-.091.235-.203.367-.334s.32-.286.505-.463s.353-.32.467-.437c.171-.178.336-.368.49-.566l.012-.016a2.59 2.59 0 0 0 .569-1.626l-.002-.092v.004a2.898 2.898 0 0 0-.958-2.317l-.003-.002a3.644 3.644 0 0 0-2.488-.829h.008zm-.128 9.136h-.018c-.354 0-.675.144-.907.377a1.208 1.208 0 0 0-.386.887v.013v-.001v.012c0 .349.152.662.393.878l.001.001c.237.227.559.367.914.367h.02h-.001h.019c.354 0 .675-.144.906-.376c.238-.223.386-.539.386-.889v-.021c0-.349-.152-.663-.393-.879L12.891 16a1.309 1.309 0 0 0-.912-.368h-.022z"/><path fill="currentColor" d="M12 0C5.373 0 0 5.373 0 12s5.373 12 12 12s12-5.373 12-12S18.627 0 12 0m0 21.882c-5.458 0-9.882-4.425-9.882-9.882S6.543 2.118 12 2.118c5.458 0 9.882 4.425 9.882 9.882c0 5.458-4.425 9.882-9.882 9.882"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ils(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.909 8.454v8.475c0 .29-.235.525-.525.525h-.022h.001h-2.748a.525.525 0 0 1-.525-.525v-.022v.001v-8.454a4.656 4.656 0 0 0-4.639-4.64H3.81v19.657c0 .29-.235.525-.525.525h-.022h.001H.516a.525.525 0 0 1-.525-.525v-.022v.001V.524c0-.29.235-.525.525-.525h.022h-.001h7.972c1.541 0 2.984.421 4.219 1.155l-.038-.021a8.471 8.471 0 0 1 3.055 3.037l.021.04a8.178 8.178 0 0 1 1.142 4.194v.053zM23.454.546v15.063a8.232 8.232 0 0 1-1.155 4.219l.021-.038a8.471 8.471 0 0 1-3.037 3.055l-.04.021A8.18 8.18 0 0 1 15.062 24h-.067h.003h-7.93a.525.525 0 0 1-.525-.525v-.022v.001V7.07c0-.29.235-.525.525-.525h.022h-.001h2.748c.29 0 .525.235.525.525v.022v-.001v13.091h4.64a4.656 4.656 0 0 0 4.64-4.639V.521c0-.29.235-.525.525-.525h.022h-.001h2.748c.29 0 .525.235.525.525v.022v-.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Imdb(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.8 13.768v-1.875l.001-.044c0-.137-.02-.27-.057-.395l.002.01q-.054-.16-.32-.16t-.268.59v.16a8.109 8.109 0 0 0 .115 2.409l-.009-.051q.054.107.24.107a.227.227 0 0 0 .24-.159v-.002c.035-.125.055-.27.055-.418l-.001-.065v.003zm-4.607-3.322a1.274 1.274 0 0 0-.083-.545l.003.009q-.08-.16-.455-.16v4.5q.375 0 .455-.187a1.672 1.672 0 0 0 .08-.621v.004zM24 2.572v-.045c0-.696-.287-1.326-.749-1.776L23.25.75A2.472 2.472 0 0 0 21.473 0h-.047h.002H2.526A2.47 2.47 0 0 0 .749.75L.748.751a2.472 2.472 0 0 0-.75 1.777v.046v-.002v18.901c0 .696.287 1.326.749 1.776l.001.001c.451.463 1.08.75 1.777.75h.047h-.002h18.903c.696 0 1.326-.287 1.776-.749l.001-.001c.463-.451.75-1.08.75-1.777v-.045v.002zm-18.8 6v6.858H3.429V8.573zm6.054 0v6.858H9.751v-4.661l-.64 4.661H7.983l-.64-4.554v4.554H5.786V8.573h2.303q.16 1.018.429 3.161l.054.054l.375-3.214zm4.66 2.41v1.446a13.038 13.038 0 0 1-.115 2.21l.008-.067c-.1.383-.408.675-.794.749l-.006.001a13.951 13.951 0 0 1-2.327.107l.024.001h-.8V8.572h1.5a6.704 6.704 0 0 1 1.488.114l-.042-.007c.499.108.881.509.959 1.011l.001.007a4.29 4.29 0 0 1 .11 1.131v-.007zm4.608.96h-.058v1.929a2.996 2.996 0 0 1-.166 1.146l.006-.021a1.007 1.007 0 0 1-.968.479h.004h-.035c-.448 0-.852-.184-1.143-.48l-.054.429H16.5V8.567h1.66v2.25a1.472 1.472 0 0 1 1.21-.508h-.004a.987.987 0 0 1 .935.451l.002.004a2.245 2.245 0 0 1 .214 1.188l.001-.009z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Import(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.909 14.837h-3.272V9.091a1.08 1.08 0 0 0-.249-.692l.001.002l-.003-.006a1.03 1.03 0 0 0-.063-.069l-.009-.009l-.028-.029L11.32.322a.79.79 0 0 0-.072-.063l-.026-.02l-.053-.04l-.026-.019a1.787 1.787 0 0 0-.058-.035l-.024-.014Q11.022.11 10.98.092l-.034-.015l-.055-.02l-.039-.012l-.063-.016l-.029-.007c-.031-.006-.062-.01-.093-.014H1.092C.492.008.005.492.001 1.091v21.818C.001 23.512.49 24 1.092 24h17.454c.603 0 1.091-.489 1.091-1.091v-5.891h3.304a1.091 1.091 0 0 0 0-2.182zh.002zM11.636 3.724L15.911 8h-4.275zm5.818 18.093H2.182V2.181h7.273V9.09c0 .603.489 1.091 1.091 1.091h6.909v4.655h-4.276l.865-.865a1.091 1.091 0 0 0-1.542-1.541l.001-.001l-2.738 2.74a.94.94 0 0 0-.06.066l-.02.026c-.015.02-.031.039-.045.06l-.014.022a.976.976 0 0 0-.042.069l-.008.016a.85.85 0 0 0-.038.081v.011a1.327 1.327 0 0 0-.032.086v.012a.932.932 0 0 0-.023.091v.027c-.004.026-.009.052-.012.079a.918.918 0 0 0 0 .222v-.005c0 .027.007.052.012.079v.027c.006.032.015.062.023.093v.011a.874.874 0 0 0 .033.09v.01c.012.027.025.055.039.081l.008.015a.997.997 0 0 0 .042.07l.013.021l.046.063l.018.024c.021.026.043.05.066.074l.005.006l2.727 2.727a1.091 1.091 0 1 0 1.542-1.541l-.001-.001l-.865-.865h4.275z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Indent(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M27.999 21.333H1.334a1.334 1.334 0 0 0-.002 2.666h26.667a1.334 1.334 0 0 0 .002-2.666zm0-5.333H14.667a1.334 1.334 0 0 0-.002 2.666H28l.037.001a1.334 1.334 0 1 0 0-2.668l-.039.001H28zm0-5.333H14.667a1.334 1.334 0 0 0-.002 2.666H28l.037.001a1.334 1.334 0 1 0 0-2.668l-.039.001H28zM1.334 2.666h26.665l.037.001a1.334 1.334 0 1 0 0-2.668L27.997 0h.002H1.334a1.334 1.334 0 0 0-.002 2.666zm26.665 2.667H14.667l-.037-.001a1.334 1.334 0 1 0 0 2.668l.039-.001h-.002H28l.037.001a1.334 1.334 0 1 0 0-2.668l-.039.001H28zM.823 18.565a1.329 1.329 0 0 0 1.453-.289l5.333-5.333c.241-.241.39-.574.39-.942s-.149-.701-.39-.942L2.276 5.724A1.334 1.334 0 0 0 0 6.666v10.668c0 .552.336 1.026.814 1.228z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Info(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.436.006a2.24 2.24 0 0 1 2.408 2.354v-.006a3.156 3.156 0 0 1-3.151 3.01l-.065-.001h.003a2.151 2.151 0 0 1-2.367-2.398l-.001.01A3.087 3.087 0 0 1 8.44.004h-.005zM3.489 24c-1.268 0-2.199-.783-1.311-4.226l1.456-6.108c.254-.978.295-1.369 0-1.369a9.57 9.57 0 0 0-3.035 1.359l.033-.021l-.633-1.057c3.086-2.622 6.638-4.159 8.158-4.159c1.268 0 1.48 1.526.845 3.874l-1.666 6.421c-.296 1.135-.168 1.526.126 1.526a6.553 6.553 0 0 0 2.863-1.456l-.008.007l.72.979c-3.004 3.052-6.281 4.232-7.549 4.232z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InjectionSyringe(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.624 2.758L7.125 4.257L5.384 2.515l1.448-1.448L5.765-.001L0 5.764l1.067 1.067l1.448-1.448l1.742 1.742l-1.499 1.499l12.96 12.96l5.399.124l2.289 2.289l1.054-.114l-2.758-2.758l-.124-5.399zm0 1.708l6.584 6.584l-4.148 4.148l-1.676-1.674l2.08-2.08l-.88-.88l-2.08 2.08l-1.55-1.55l2.08-2.08l-.88-.88l-2.08 2.08L4.48 8.62z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Inr(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.307 5.83v1.759c0 .29-.235.525-.525.525h-.022h.001h-2.864a6.38 6.38 0 0 1-2.19 3.981l-.009.008a8.584 8.584 0 0 1-4.673 1.873l-.032.002q2.846 3.034 7.824 9.136a.465.465 0 0 1 .067.581l.001-.002a.496.496 0 0 1-.459.308L12.389 24h.002h-3.345a.5.5 0 0 1-.404-.204l-.001-.001Q3.425 17.539.152 14.062a.509.509 0 0 1-.154-.365v-.011v.001v-2.167a.553.553 0 0 1 .545-.546h1.909a7.664 7.664 0 0 0 3.668-.753l-.046.02a3.35 3.35 0 0 0 1.742-2.107l.005-.023H.52a.525.525 0 0 1-.525-.525v-.022v.001v-1.759c0-.29.235-.525.525-.525h.022h-.001h7.04Q6.61 3.355 3.013 3.355H.542a.553.553 0 0 1-.546-.545V.521c0-.29.235-.525.525-.525h.022h-.001h14.203c.29 0 .525.235.525.525v.022v-.001v1.759c0 .29-.235.525-.525.525h-.022h.001h-3.971a5.849 5.849 0 0 1 1.085 2.417l.006.037h2.935c.29 0 .525.235.525.525v.022v-.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Instagram(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16 12a4 4 0 1 0-1.172 2.829A3.843 3.843 0 0 0 16 12.06l-.001-.063zm2.16 0a6.135 6.135 0 1 1-1.797-4.359a5.922 5.922 0 0 1 1.798 4.256l-.001.109zm1.687-6.406v.002a1.44 1.44 0 1 1-.422-1.018c.256.251.415.601.415.988v.029v-.001zm-7.84-3.44l-1.195-.008q-1.086-.008-1.649 0t-1.508.047c-.585.02-1.14.078-1.683.17l.073-.01c-.425.07-.802.17-1.163.303l.043-.014a4.117 4.117 0 0 0-2.272 2.254l-.01.027a5.975 5.975 0 0 0-.284 1.083l-.005.037a11.76 11.76 0 0 0-.159 1.589l-.001.021q-.039.946-.047 1.508t0 1.649t.008 1.195t-.008 1.195t0 1.649t.047 1.508c.02.585.078 1.14.17 1.683l-.01-.073c.07.425.17.802.303 1.163l-.014-.043a4.117 4.117 0 0 0 2.254 2.272l.027.01c.318.119.695.219 1.083.284l.037.005c.469.082 1.024.14 1.588.159l.021.001q.946.039 1.508.047t1.649 0l1.188-.024l1.195.008q1.086.008 1.649 0t1.508-.047c.585-.02 1.14-.078 1.683-.17l-.073.01c.425-.07.802-.17 1.163-.303l-.043.014a4.117 4.117 0 0 0 2.272-2.254l.01-.027c.119-.318.219-.695.284-1.083l.005-.037c.082-.469.14-1.024.159-1.588l.001-.021q.039-.946.047-1.508t0-1.649t-.008-1.195t.008-1.195t0-1.649t-.047-1.508c-.02-.585-.078-1.14-.17-1.683l.01.073a6.254 6.254 0 0 0-.303-1.163l.014.043a4.117 4.117 0 0 0-2.254-2.272l-.027-.01a5.975 5.975 0 0 0-1.083-.284l-.037-.005a11.76 11.76 0 0 0-1.588-.159l-.021-.001q-.946-.039-1.508-.047t-1.649 0zM24 12q0 3.578-.08 4.953a6.64 6.64 0 0 1-6.985 6.968l.016.001q-1.375.08-4.953.08t-4.953-.08a6.64 6.64 0 0 1-6.968-6.985l-.001.016q-.08-1.375-.08-4.953t.08-4.953A6.64 6.64 0 0 1 7.061.079L7.045.078q1.375-.08 4.953-.08t4.953.08a6.64 6.64 0 0 1 6.968 6.985l.001-.016Q24 8.421 24 12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func InternetExplorer(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24.436 12.56v.051c0 .482-.035.956-.102 1.42l.006-.053H8.645l-.001.075c0 1.3.577 2.466 1.488 3.255l.005.005a5.058 5.058 0 0 0 3.442 1.344l.073-.001h-.004h.025c.918 0 1.781-.235 2.532-.648l-.027.014a4.77 4.77 0 0 0 1.85-1.76l.012-.022h5.768a10.786 10.786 0 0 1-2.33 3.837l.005-.005a10.964 10.964 0 0 1-8.022 3.478c-1.769 0-3.441-.418-4.921-1.161l.063.029a12.105 12.105 0 0 1-5.339 1.581l-.033.002q-3.232 0-3.232-3.586c.039-1.351.26-2.636.641-3.85l-.027.101a19.52 19.52 0 0 1 1.536-3.21l-.05.09c1.748-3.199 3.909-5.929 6.454-8.241l.026-.023Q6.07 6.36 2.757 10.109a10.718 10.718 0 0 1 3.842-6.112l.023-.018a10.622 10.622 0 0 1 6.739-2.393h.105h-.005q.409 0 .614.014C15.79.702 17.803.126 19.938.003l.039-.002h.068a6.57 6.57 0 0 1 1.559.187l-.045-.009a4.23 4.23 0 0 1 1.305.562l-.016-.01c.391.258.7.61.901 1.025l.007.015c.207.436.328.946.328 1.486l-.001.087V3.34a10.043 10.043 0 0 1-1.049 3.958l.026-.058a10.659 10.659 0 0 1 1.376 5.272v.053v-.003zm-.954-8.728a2.326 2.326 0 0 0-.719-1.799l-.001-.001a2.628 2.628 0 0 0-1.758-.67l-.116.002h.005a8.233 8.233 0 0 0-3.508.976l.044-.022a10.919 10.919 0 0 1 3.051 1.806l-.017-.014a11.126 11.126 0 0 1 2.298 2.62l.027.046a8.705 8.705 0 0 0 .692-2.919l.001-.025zM1.746 20.7a2.378 2.378 0 0 0 .662 1.808l-.001-.001a2.527 2.527 0 0 0 1.842.634h-.007a7.739 7.739 0 0 0 3.66-1.151l-.033.019a11.08 11.08 0 0 1-2.894-2.477l-.015-.019a10.613 10.613 0 0 1-1.852-3.266l-.023-.075a11.073 11.073 0 0 0-1.337 4.488zm6.872-9.722h9.927a4.337 4.337 0 0 0-1.535-3.227l-.006-.005c-.908-.804-2.11-1.295-3.426-1.295s-2.518.491-3.431 1.3l.005-.005a4.348 4.348 0 0 0-1.534 3.227z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Intersex(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.713.429V.413c0-.228.185-.413.413-.413h.016h-.001h3.864c.234 0 .445.098.595.254a.82.82 0 0 1 .254.595V4.73a.413.413 0 0 1-.413.413h-.017h.001h-.873a.413.413 0 0 1-.413-.413v-.017v.001V2.92l-3.404 3.415a7.459 7.459 0 0 1 1.688 4.741v.07v-.004l.001.1a7.414 7.414 0 0 1-1.98 5.054l.004-.005A7.476 7.476 0 0 1 8.6 18.8l-.032.003v1.768H9.87c.228 0 .413.185.413.413v.017V21v.873a.413.413 0 0 1-.413.413h-.017h.001h-1.286v1.302a.413.413 0 0 1-.413.413h-.017h.001h-.873a.413.413 0 0 1-.413-.413v-.017v.001v-1.286H5.551a.413.413 0 0 1-.413-.413v-.017v.001v-.873c0-.228.185-.413.413-.413h.017h-.001h1.288V18.8a7.481 7.481 0 0 1-3.644-1.394l.021.015a7.74 7.74 0 0 1-2.478-2.947l-.02-.046a7.461 7.461 0 0 1-.737-3.258c0-.227.01-.453.03-.675l-.002.029a7.474 7.474 0 0 1 2.144-4.735l-.001.001a7.671 7.671 0 0 1 5.548-2.363c.767 0 1.508.112 2.207.321l-.054-.014a7.856 7.856 0 0 1 2.666 1.399l-.015-.012l3.415-3.402h-1.81a.413.413 0 0 1-.413-.413v-.017v.001zm-6 16.714l.087.001a5.763 5.763 0 0 0 4.152-1.759l.001-.002a5.765 5.765 0 0 0 1.761-4.153l-.001-.092v.005l.001-.087a5.763 5.763 0 0 0-1.759-4.152l-.002-.001C10.87 5.814 9.37 5.141 7.713 5.141s-3.157.674-4.24 1.762a5.765 5.765 0 0 0-1.76 4.245v-.005l-.001.087c0 1.629.674 3.101 1.759 4.152l.002.001a5.765 5.765 0 0 0 4.153 1.761l.092-.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Invision(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.825 0H2.175A2.177 2.177 0 0 0 0 2.175v19.65A2.177 2.177 0 0 0 2.175 24h19.65A2.177 2.177 0 0 0 24 21.825V2.175A2.177 2.177 0 0 0 21.825 0M8.02 4.667a1.443 1.443 0 0 1 0 2.884v.027a1.454 1.454 0 1 1 0-2.91zm9.816 13.815a1.974 1.974 0 0 1-2.199-2.097v-.003c0-.382.055-.75.157-1.099l-.007.028l.71-2.545c.07-.221.11-.476.11-.739v-.037v.002a1.209 1.209 0 0 0-1.304-1.315h.004c-1.013 0-1.68.727-2.025 2.13l-1.381 5.535H9.488l.432-1.74a3.431 3.431 0 0 1-2.903 1.885H7.01a1.932 1.932 0 0 1-2.156-2.122l-.001.008c.004-.387.051-.76.137-1.119l-.007.034l1.1-4.49H4.376l.518-1.9h4.1l-1.62 6.4c-.084.293-.138.63-.15.979v.007c0 .41.2.53.517.6a2.71 2.71 0 0 0 2.554-1.821l.006-.019l1.06-4.25H9.638l.52-1.875h3.72l-.48 2.16a4.21 4.21 0 0 1 3.189-2.352l.023-.003a2.508 2.508 0 0 1 2.474 2.833l.001-.013v.004c0 .561-.088 1.101-.25 1.608l.01-.037l-.69 2.476a2.855 2.855 0 0 0-.1.654v.006c0 .434.18.645.494.645s.735-.24 1.2-1.56l.943.36c-.555 1.964-1.576 2.774-2.85 2.774z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Island(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1.258 4.956l-.16 1.162l1.13-.316l.441 1.088l.82-.838l.926.72l.292-1.135l1.162.16l-.316-1.13l1.088-.44l-.838-.82l.72-.926l-1.135-.292l.16-1.162l-1.13.316L3.977.255l-.825.838l-.926-.72l-.292 1.135l-1.162-.16l.316 1.13L0 2.919l.838.82l-.72.926zM2.687 1.73a1.946 1.946 0 1 1-1.209 2.485l-.004-.014a1.946 1.946 0 0 1 1.199-2.465l.014-.004zm2.004 18.852a.399.399 0 0 0 .284-.119l.246-.254c.434-.447 1.04-.724 1.711-.724s1.277.277 1.711.724l.001.001l.252.258a.399.399 0 0 0 .284.119h.001c.11 0 .21-.046.28-.119l.266-.27a2.382 2.382 0 0 1 1.703-.714c.673 0 1.282.279 1.716.727l.001.001l.24.246a.396.396 0 0 0 .568.003l.244-.252c.434-.447 1.04-.724 1.711-.724s1.277.277 1.711.724l.001.001l.25.254a.39.39 0 0 0 .568 0l.246-.254c.162-.164.346-.305.548-.418l.012-.006a13.083 13.083 0 0 0-4.309-2.637l-.091-.029c-.674-2.876-.08-6.654.846-9.511c.021-.022.04-.042.058-.064a3.425 3.425 0 0 1 2.615 3.323c0 .858-.316 1.642-.838 2.242l.004-.004l.073.001a3.418 3.418 0 1 0 0-6.836l-.111.002h.005c.261-.075.563-.127.875-.145h.011a3.695 3.695 0 0 1 3.366 1.626l.008.013a2.767 2.767 0 0 0 .273-1.367v.007a3.426 3.426 0 0 0-3.808-2.967l.016-.002a3.818 3.818 0 0 0-.666.092l.026-.005l.001-.086c0-.244-.024-.482-.069-.712l.004.023A3.111 3.111 0 0 0 14.073.056l.015-.001c-.44.088-.827.277-1.148.542l.004-.003a3.288 3.288 0 0 1 2.02 2.031l.007.023a4.76 4.76 0 0 0-.925-.598l-.028-.013c-2.027-.984-4.335-.417-5.155 1.271a2.89 2.89 0 0 0-.284 1.479l-.001-.009a4.16 4.16 0 0 1 2.21-.627c.718 0 1.395.18 1.986.497l-.023-.011a4.9 4.9 0 0 1 .863.539l-.011-.008a3.736 3.736 0 0 0-4.116 6.034l-.002-.002c.369.388.82.696 1.325.895l.025.009a3.735 3.735 0 0 1 3.959-4.888l-.017-.002c-2.038 3.902-3.261 6.946-3.465 9.291h-.05a10.717 10.717 0 0 0-2.956.419l.076-.019a10.652 10.652 0 0 1-2.484-4.129l-.021-.075c0-.016 0-.032.006-.045A1.822 1.822 0 0 1 8.41 14.48v-.006a1.822 1.822 0 1 0-2.187-2.914l.004-.003a2.3 2.3 0 0 1 .328-.34l.004-.003a1.96 1.96 0 0 1 1.296-.486c.24 0 .471.043.684.122l-.014-.004a1.501 1.501 0 0 0-.318-.674l.002.002a1.815 1.815 0 0 0-1.316-.562c-.484 0-.924.189-1.25.497l.001-.001c-.089.075-.17.154-.244.238l-.002.002a2.202 2.202 0 0 0-.272-.309l-.001-.001a1.65 1.65 0 0 0-1.196-.51a1.65 1.65 0 0 0-1.135.45l.001-.001c-.15.163-.261.363-.318.585l-.002.009A1.732 1.732 0 0 1 4 10.815l-.005-.004a2.604 2.604 0 0 0-.618.046l.016-.003a2.019 2.019 0 0 0-1.803 2.188l-.001-.008c.054.278.175.522.347.722l-.002-.002c.297-.73.946-1.26 1.731-1.383l.013-.002a2.342 2.342 0 0 1 .537-.043h-.004a1.99 1.99 0 1 0 1.033 3.846l-.014.003a1.996 1.996 0 0 1-.815-1.984l-.002.012a1.979 1.979 0 0 1 .944-1.366l.009-.005a16.413 16.413 0 0 0 1.315 4.837l-.042-.105a13.653 13.653 0 0 0-3.351 2.243l.011-.01c.429.146.794.381 1.09.684l.001.001a.485.485 0 0 0 .297.102z"/><path fill="currentColor" d="M1.202 21.61a.42.42 0 0 0 .297-.125a1.324 1.324 0 0 1 1.9-.001c.329.337.787.546 1.295.546s.966-.209 1.294-.546l.246-.252a.978.978 0 0 1 1.406 0l.252.258c.328.338.786.547 1.294.547c.504 0 .961-.207 1.288-.541l.268-.27a.975.975 0 0 1 1.401.005l.24.246c.326.34.785.552 1.293.552h.003c.508 0 .966-.211 1.293-.549l.001-.001l.244-.25a.98.98 0 0 1 1.406 0l.246.252c.328.337.786.546 1.293.546s.965-.209 1.293-.546l.247-.255a.978.978 0 0 1 1.406 0a.41.41 0 0 0 .592 0a.44.44 0 0 0 0-.61a1.804 1.804 0 0 0-2.589-.001l-.246.254a.978.978 0 0 1-1.406 0l-.246-.254a1.804 1.804 0 0 0-2.589-.001l-.244.25a.98.98 0 0 1-1.406-.003l-.24-.246a1.797 1.797 0 0 0-2.584-.011l-.266.27a.979.979 0 0 1-1.401-.005l-.252-.258c-.329-.337-.787-.546-1.295-.546s-.966.209-1.294.546l-.246.252a.975.975 0 0 1-.701.297H4.69a.967.967 0 0 1-.697-.296c-.391-.403-.937-.653-1.542-.653s-1.151.25-1.541.652l-.001.001a.439.439 0 0 0 .001.61a.4.4 0 0 0 .295.133z"/><path fill="currentColor" d="M21.696 22.584a1.804 1.804 0 0 0-2.589-.001l-.246.254a.98.98 0 0 1-1.406 0l-.247-.255a1.804 1.804 0 0 0-2.589-.001l-.244.25a.98.98 0 0 1-1.406-.003l-.24-.246a1.797 1.797 0 0 0-2.584-.011l-.266.27a.979.979 0 0 1-1.401-.005l-.252-.258a1.804 1.804 0 0 0-2.589-.001l-.246.254a.978.978 0 0 1-1.406 0c-.391-.403-.937-.653-1.542-.653s-1.151.25-1.541.652l-.001.001a.439.439 0 0 0 .001.61a.41.41 0 0 0 .594.002a1.324 1.324 0 0 1 1.9-.001c.329.337.787.546 1.295.546s.966-.209 1.294-.546l.246-.252a.98.98 0 0 1 1.406 0l.252.258a1.8 1.8 0 0 0 2.582.006l.268-.27a.975.975 0 0 1 1.401.005l.24.246c.326.34.785.552 1.293.552h.003c.507 0 .965-.21 1.292-.548v-.001l.244-.25a.978.978 0 0 1 1.406 0l.246.252c.328.337.786.546 1.293.546s.965-.209 1.293-.546l.246-.252a.978.978 0 0 1 1.406 0a.41.41 0 0 0 .592 0a.433.433 0 0 0 .002-.61z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Italic(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.182 0H5.423a1.091 1.091 0 0 0 0 2.182h.032h-.002h3.033L4.559 21.819h-3.5a1.091 1.091 0 0 0 0 2.182h.032h-.002h8.727a1.092 1.092 0 0 0 .002-2.182H6.783L10.71 2.182h3.469A1.092 1.092 0 0 0 14.181 0h-.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Iyzigo(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M33.6 24H2.4A2.4 2.4 0 0 1 0 21.6V2.4A2.4 2.4 0 0 1 2.4 0h31.2A2.4 2.4 0 0 1 36 2.4v19.2a2.4 2.4 0 0 1-2.4 2.4m-13.662-7.723l-.028.004a.516.516 0 0 0-.161.095l.001-.001c.03.086.047.185.047.288v.023v-.001v1.92a3.253 3.253 0 0 0 .013.425l-.001-.013a.59.59 0 0 0 .123.012h.005l.023.001a.482.482 0 0 0 .19-.039l-.003.001c.005-.043.008-.094.008-.144v-.034v.002v-.684a.629.629 0 0 0 .353.107h.015h-.001h.005a.669.669 0 0 0 .428-.154l-.001.001a1.242 1.242 0 0 0 .357-1.01v.005c0-.566-.258-.89-.708-.89a.723.723 0 0 0-.252.047l.005-.002a.796.796 0 0 0-.286.161l.001-.001c-.019-.077-.066-.12-.133-.12zm3.28 2.469h-.004a.254.254 0 0 0-.078.012h.002a.498.498 0 0 0-.025.109v.002c.018.066.11.136.312.136c.218 0 .42-.248.619-.758c.227-.63.433-1.36.54-1.8a.421.421 0 0 0 .02-.11v-.002a.183.183 0 0 0-.021-.087l.001.001l-.182-.008l-.022-.002a.154.154 0 0 0-.102.039l-.09.385l-.119.51c-.059.251-.121.51-.14.568l-.037-.146c-.026-.106-.062-.254-.1-.4c-.08-.289-.193-.667-.234-.8c-.018-.067-.058-.165-.08-.165c-.21 0-.287.023-.32.095c.173.48.462 1.43.498 1.553a.657.657 0 0 0 .082.193l-.002-.003a1.128 1.128 0 0 1-.278.656l.001-.001a.4.4 0 0 1-.146.027h-.015h.001h-.084zm-6.06 0h-.008a.25.25 0 0 0-.078.012h.002a.499.499 0 0 0-.024.109v.002c.018.066.111.136.313.136s.422-.255.618-.758c.23-.64.435-1.37.54-1.8a.421.421 0 0 0 .02-.11v-.002a.183.183 0 0 0-.021-.087l.001.001l-.182-.008l-.022-.002a.151.151 0 0 0-.101.039l-.094.4l-.118.498a9.79 9.79 0 0 1-.139.568l-.033-.13l-.103-.418c-.08-.289-.194-.667-.234-.8c-.017-.067-.058-.165-.08-.165c-.21 0-.288.023-.32.095c.173.48.462 1.43.498 1.553a.657.657 0 0 0 .082.193l-.002-.003a1.128 1.128 0 0 1-.278.656l.001-.001a.398.398 0 0 1-.145.027h-.016h.001h-.085zm14.522-3.025c-.213.014-.286.097-.31.165l-.021.37a.581.581 0 0 1-.086.027l-.004.001a.425.425 0 0 0-.162.063l.002-.001c-.034.031-.038.034-.038.145a.514.514 0 0 0 .25.04h-.002v1.018c0 .466.148.692.453.692h.02c.119 0 .232-.03.331-.082l-.004.002a.145.145 0 0 0 .061-.079v-.116a.586.586 0 0 0-.008-.094v.003a.75.75 0 0 1-.276.08h-.003c-.16 0-.206-.15-.223-.392v-.478c0-.177.004-.371.009-.56h.337a.876.876 0 0 0 .103-.009l-.004.001c.055-.012.08-.053.095-.263h-.528l.006-.419v-.113zm-3.285.48c-.361 0-.791.19-.791 1.092a1.065 1.065 0 0 0 .221.752l-.002-.002c.132.124.311.2.507.2l.056-.002h-.002c.111-.001.219-.016.322-.043l-.009.002c.112-.028.235-.114.24-.198a.774.774 0 0 0-.027-.17l.001.006a.945.945 0 0 1-.426.1l-.033-.001h.002c-.327 0-.48-.19-.486-.598h.229c.173 0 .344-.009.513-.027l-.021.002a.8.8 0 0 0 .292-.097l-.004.002c.016-.068.025-.147.025-.228v-.013v.001c0-.07-.004-.136-.009-.198v-.023a.56.56 0 0 0-.56-.56l-.034.001h.002zm-12.619 1.654h-.016a.368.368 0 0 0-.066.205c.041.102.258.177.514.177c.418 0 .688-.226.688-.577v-.001a.538.538 0 0 0-.238-.447l-.002-.001a3.968 3.968 0 0 0-.25-.186c-.063-.045-.123-.087-.166-.122a.268.268 0 0 1-.131-.209v-.001c0-.128.113-.214.28-.214h.004c.071 0 .139.014.202.039l-.004-.001a.762.762 0 0 0 .129.031h.004a.193.193 0 0 0 .118-.175c-.019-.058-.133-.149-.366-.165a1.069 1.069 0 0 0-.12-.009c-.38 0-.626.199-.626.507a.526.526 0 0 0 .18.387l.001.001l.099.086c.08.069.194.15.293.222c.16.112.218.195.218.298v.007a.24.24 0 0 1-.24.24h-.026a.85.85 0 0 1-.371-.059l.006.002a.265.265 0 0 0-.114-.032h-.001zm-2.542-1.657c-.361 0-.79.19-.79 1.092a1.07 1.07 0 0 0 .22.752l-.002-.002c.132.124.311.2.507.2l.055-.002h-.002c.112-.001.22-.016.323-.043l-.009.002c.112-.028.234-.114.24-.198a1.004 1.004 0 0 0-.026-.171l.001.007a.945.945 0 0 1-.426.1l-.032-.001h.002c-.328 0-.48-.19-.486-.598h.227c.173 0 .345-.009.514-.027l-.021.002a.83.83 0 0 0 .293-.097l-.004.002c.016-.069.025-.148.025-.228v-.012v.001c0-.063 0-.124-.007-.181c-.042-.395-.243-.597-.6-.597zm9.238 1.882a.358.358 0 0 0 .245.147h.008a.284.284 0 0 0 .213-.094c0-.036 0-.1.004-.153l-.08-.022a.268.268 0 0 1-.04-.142l.001-.026v.001v-.363l.004-.469v-.186c0-.382-.121-.554-.4-.572c-.07-.004-.08-.004-.144-.004a1.041 1.041 0 0 0-.613.201l.003-.002v.091a.342.342 0 0 0 .062.166l-.001-.001a.312.312 0 0 0 .122-.034l-.002.001c.127-.06.274-.099.43-.107h.003c.183 0 .197.114.197.288v.235a.595.595 0 0 0-.123-.013h-.012h.001h-.021c-.034 0-.067.002-.1.005h.004h-.014a.681.681 0 0 0-.681.658v.001a.496.496 0 0 0 .494.544h.022h-.001l.04.001c.095 0 .185-.02.266-.056l-.004.002a.658.658 0 0 0 .119-.094zm-7.486 0a.358.358 0 0 0 .245.147h.009a.282.282 0 0 0 .211-.094c0-.036 0-.099.004-.153l-.08-.022a.27.27 0 0 1-.04-.142l.001-.026v.001v-.362l.004-.47v-.182c0-.382-.121-.554-.4-.573c-.07-.004-.08-.004-.144-.004a1.041 1.041 0 0 0-.613.201l.003-.002l-.004.091a.363.363 0 0 0 .062.166l-.001-.001a.303.303 0 0 0 .121-.034l-.002.001c.127-.059.274-.097.429-.107h.004c.183 0 .198.114.198.289v.234a.57.57 0 0 0-.121-.012h-.015h.001h-.13a.681.681 0 0 0-.681.658v.001a.496.496 0 0 0 .494.544h.022h-.001l.04.001c.095 0 .185-.02.266-.056l-.004.002a.658.658 0 0 0 .119-.094zM29.5 16.283a.505.505 0 0 0-.194.1l.001-.001c.032.088.05.189.05.294v.015v-.001v.701c0 .32-.004.581-.01.8c.231 0 .273 0 .32-.041a.3.3 0 0 0 .045-.194v.001v-1.281a.484.484 0 0 1 .354-.173h.001c.24 0 .326.109.326.432v1.111c0 .054.012.106.034.152l-.001-.002c.24 0 .297-.013.338-.049c-.004-.19-.008-.378-.008-.59v-.63c0-.47-.094-.725-.586-.725a.779.779 0 0 0-.261.047l.005-.002a.838.838 0 0 0-.285.161l.001-.001c-.021-.082-.064-.124-.13-.124m-2.88.215c.206 0 .276.11.276.436v1.138a.27.27 0 0 0 .03.124l-.001-.002c.234 0 .289-.013.338-.049c-.008-.148-.008-.334-.008-.54v-.692c0-.35-.054-.678-.45-.709a.445.445 0 0 0-.065-.004h-.016h.001h-.006a.762.762 0 0 0-.25.042l.005-.002a.871.871 0 0 0-.299.166l.001-.001a.415.415 0 0 0-.358-.206h-.022a.813.813 0 0 0-.602.198l.001-.001c-.019-.084-.059-.123-.127-.123a.545.545 0 0 0-.195.099l.001-.001c.031.089.05.191.05.297v.012v-.001v.738c0 .31 0 .565-.004.774c.226 0 .27-.005.32-.041a.21.21 0 0 0 .043-.129l-.001-.025v.001v-1.333a.49.49 0 0 1 .357-.173h.001c.204 0 .277.114.277.437v.494c0 .31 0 .56-.008.77h.169a1.555 1.555 0 0 0 .147 0h-.003a.062.062 0 0 0 .054-.061v-1.464a.481.481 0 0 1 .34-.173h.001zM6.016 8.862h-.003c-.212 0-.418.026-.614.076l.017-.004l2.415 5.76l-1.208 2.873c.188.053.404.084.627.086h.001a1.338 1.338 0 0 0 1.25-.674l.003-.007l3.379-8.039a3.096 3.096 0 0 0-.608-.065h-.002a1.268 1.268 0 0 0-1.324.735l-.003.008l-1.297 3.113l-1.296-3.107a1.362 1.362 0 0 0-1.343-.754h.005zm7.725 1.378h2.132l-2.8 2.715a2.57 2.57 0 0 0-.573 2.226l-.003-.017h5.137a1.253 1.253 0 0 0-1.22-1.371H14.36l2.855-2.891a2.352 2.352 0 0 0 .418-2.053l.004.016h-5.168c.102 1.307 1.052 1.373 1.242 1.373zM29.27 8.862a3.051 3.051 0 0 0-2.934 3.153v-.005a2.942 2.942 0 1 0 5.869-.009l.001.01l.002-.097a3.05 3.05 0 0 0-2.932-3.048h-.006zm-5.746-.008a3.154 3.154 0 0 0-.005 6.303h.006a2.988 2.988 0 0 0 2.252-1.036l.003-.004a.78.78 0 0 0-.716-.475h-.014h.001c-.22.008-.425.066-.605.164l.007-.004a2.24 2.24 0 0 1-.909.217h-.024c-.501 0-.937-.277-1.164-.685l-.003-.007c-.174-.398-.276-.861-.276-1.348s.102-.95.285-1.37l-.009.022c.2-.401.607-.672 1.078-.672l.067.002h-.003c.355.01.688.091.989.23l-.016-.006c.175.09.38.146.598.154h.015a.764.764 0 0 0 .689-.433l.002-.005a2.97 2.97 0 0 0-2.252-1.044h-.001zm-3.6.007c-.067 0-1.65.012-1.65 1.084v5.211c.4 0 1.655-.076 1.655-1.04V8.864zm-15.071 0c-.07.003-1.653.019-1.653 1.086v5.211c.4 0 1.655-.076 1.655-1.04V8.864zM19.092 5.6a.825.825 0 0 0 .003 1.652a.826.826 0 0 0 0-1.651h-.002zm-15.064 0a.826.826 0 1 0 0 1.652a.826.826 0 0 0 0-1.652m16.43 12.366h-.009a.399.399 0 0 1-.297-.132v-1.178a.583.583 0 0 1 .361-.16h.002c.274 0 .407.202.407.618c0 .629-.24.853-.465.853zm1.688-.02h-.029c-.148-.009-.219-.105-.219-.292c0-.286.21-.353.368-.378c.04-.005.087-.008.135-.009h.017c.023 0 .045.002.066.005h-.003v.292l.001.044c0 .054-.003.108-.009.16l.001-.006a.373.373 0 0 1-.322.186h-.006zm-7.485 0h-.03c-.147-.009-.219-.105-.219-.292c0-.284.21-.352.368-.378c.04-.005.087-.008.134-.009h.016c.023 0 .045.002.067.005h-.003v.292l.001.044c0 .054-.003.108-.009.16l.001-.006a.374.374 0 0 1-.323.186h-.004zm13.472-.865h-.16c.018-.352.146-.58.342-.61c.215 0 .312.104.334.346l.004.09v.006a.41.41 0 0 1-.022.132l.001-.003a.213.213 0 0 1-.095.025a4.907 4.907 0 0 1-.403.013zm-15.161 0h-.16c.018-.354.146-.581.342-.61c.216 0 .312.103.332.346l.005.091v.008a.402.402 0 0 1-.021.13l.001-.003a.213.213 0 0 1-.095.025a4.935 4.935 0 0 1-.404.013m16.295-3.067c-.791 0-1.435-.906-1.435-2.021s.64-2.02 1.435-2.02s1.436.906 1.436 2.02s-.645 2.021-1.437 2.021z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Java(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.701 18.561s-.918.534.653.714c.575.085 1.239.134 1.913.134c1.084 0 2.138-.125 3.149-.363l-.093.018c.374.228.809.445 1.262.624l.059.02c-4.698 2.014-10.633-.117-6.942-1.148zm-.574-2.628s-1.029.762.542.924a19.461 19.461 0 0 0 6.541-.332l-.129.024c.275.258.604.463.968.596l.02.006c-5.68 1.661-12.008.131-7.942-1.218m4.839-4.458c1.158 1.333-.304 2.532-.304 2.532s2.939-1.52 1.59-3.418c-1.261-1.772-2.228-2.653 3.006-5.688c0 0-8.216 2.052-4.292 6.574"/><path fill="currentColor" d="M16.18 20.505s.678.56-.747.992c-2.712.822-11.287 1.07-13.67.033c-.856-.373.75-.89 1.254-.998c.232-.059.499-.093.774-.093h.057h-.003c-.952-.671-6.155 1.318-2.64 1.886c9.579 1.554 17.462-.7 14.978-1.82zM6.142 13.21s-4.362 1.036-1.545 1.412a33.229 33.229 0 0 0 5.908-.073l-.139.012c1.805-.152 3.618-.48 3.618-.48a7.533 7.533 0 0 0-1.126.605l.029-.018c-4.43 1.165-12.986.623-10.523-.569a8.14 8.14 0 0 1 3.734-.893h.046h-.002zm7.824 4.375c4.502-2.34 2.421-4.588.967-4.286a3.299 3.299 0 0 0-.539.146l.023-.007a.822.822 0 0 1 .379-.295l.006-.002c2.874-1.01 5.086 2.981-.928 4.56a.395.395 0 0 0 .089-.115l.001-.002zM11.252 0s2.494 2.494-2.366 6.33c-3.896 3.077-.889 4.831 0 6.836c-2.274-2.052-3.943-3.858-2.824-5.54c1.644-2.468 6.197-3.664 5.19-7.627z"/><path fill="currentColor" d="M6.585 23.925c4.32.277 10.96-.154 11.12-2.198c0 0-.302.775-3.572 1.391c-1.806.326-3.885.512-6.008.512c-1.739 0-3.448-.125-5.121-.366l.191.023s.553.458 3.393.64z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Jcb(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M33.6 24H2.4A2.4 2.4 0 0 1 0 21.6V2.4A2.4 2.4 0 0 1 2.4 0h31.2A2.4 2.4 0 0 1 36 2.4v19.2a2.4 2.4 0 0 1-2.4 2.4m-18.984-9.845v6.846h4.353a2.97 2.97 0 0 0 2.967-2.967V3.199h-4.352a2.97 2.97 0 0 0-2.968 2.967v3.88a5.404 5.404 0 0 1 3.571-.959l-.02-.001c.19 0 .394.005.605.014c.856.056 1.651.183 2.42.377l-.092-.02v1.256a5.597 5.597 0 0 0-2.23-.651l-.019-.001a4.406 4.406 0 0 0-.301-.01c-1.414 0-2.258.766-2.258 2.048a1.969 1.969 0 0 0 2.233 2.051l-.01.001c.106 0 .219-.005.336-.014a5.932 5.932 0 0 0 2.282-.668l-.032.016v1.256h-.009a12.48 12.48 0 0 1-2.282.354l-.038.002c-.211.01-.414.014-.605.014a5.425 5.425 0 0 1-3.568-.969zm8.183.8v6.032h4.353a2.97 2.97 0 0 0 2.967-2.967V3.2h-4.337a2.969 2.969 0 0 0-2.966 2.967v3.081h4.189l.053-.001c.085 0 .168.006.249.019l-.009-.001c1.016.052 1.646.582 1.646 1.386a1.395 1.395 0 0 1-1.351 1.35h-.002v.033c1.014.068 1.695.645 1.695 1.434c0 .874-.778 1.483-1.89 1.483zM6.4 14.725v6.24h4.353a2.97 2.97 0 0 0 2.966-2.964v-14.8H9.383a2.969 2.969 0 0 0-2.966 2.967v7.319c.742.397 1.62.642 2.552.669h.008a1.417 1.417 0 0 0 1.585-1.407l-.001-.063v.003v-3.456h2.543v3.44c0 1.612-1.229 2.43-3.651 2.43c-1.084 0-2.136-.137-3.139-.395l.087.019zm17.866-2.282v1.565h1.711c.025 0 .066-.004.106-.008s.08-.008.106-.008a.763.763 0 0 0 .603-.745v-.023v.001a.791.791 0 0 0-.598-.765l-.005-.001a.84.84 0 0 0-.172-.018l-.041.001h.002zm0-2.314v1.434h1.55l.034.001a.457.457 0 0 0 .129-.019l-.003.001a.71.71 0 0 0 .004-1.401l-.004-.001a.503.503 0 0 1-.07-.007h.004a.853.853 0 0 0-.095-.009h-.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Jekyll(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.845 24a2.845 2.845 0 0 1-2.741-3.611l-.005.02l-.015-.009l.09-.226l7-18.167A3.49 3.49 0 0 0 7.168.721l.003.022a.48.48 0 0 1 .014-.246L7.184.5l.014-.035V.458l.009-.019C7.315.147 7.657 0 8.215 0a7.702 7.702 0 0 1 2.607.592l-.051-.019a8.65 8.65 0 0 1 2.059 1.093l-.026-.018c.6.45.842.854.707 1.2l-.031.045l-.016.015a.755.755 0 0 1-.149.164l-.001.001a3.901 3.901 0 0 0-.831.929l-.009.016l-7.068 18.421l-.016-.006a2.878 2.878 0 0 1-2.542 1.561h-.007zM.658 20.282l-.02.05a2.363 2.363 0 0 0 1.348 3.012l.016.006c.245.098.53.154.828.154h.01h-.001h.005c.991 0 1.838-.617 2.177-1.488L5.026 22l.027-.061l6.959-18.09c.267-.443.593-.82.972-1.134l.007-.006l.016-.016c.012-.015.02-.015.02-.03c0-.06-.061-.27-.557-.645a7.908 7.908 0 0 0-1.845-.982l-.055-.018A7.354 7.354 0 0 0 8.22.479L8.198.478c-.39 0-.524.082-.545.126v.04a3.968 3.968 0 0 1-.039 1.541l.005-.026l-6.962 18.12zm8.95-11.507a7.078 7.078 0 0 1-1.809 1.491l-.034.018c-.946.252-1.718.51-2.47.809l.15-.053c-.478.281-.875.65-1.18 1.088l-.008.012l-3.215 8.36a1.934 1.934 0 0 0 1.112 2.415l.013.005a1.882 1.882 0 0 0 2.407-1.043l.005-.013zM4.672 18.76a.2.2 0 0 1 .265.099l.001.001a.201.201 0 1 1-.37.161l-.001-.001a.2.2 0 0 1 .104-.26l.001-.001zm-1.014-1.8a.3.3 0 0 1-.418-.276a.3.3 0 0 1 .572-.126l.001.002a.282.282 0 0 1-.153.389l-.002.001zm.286-1.1a.58.58 0 0 1 .283-.756l.004-.001a.575.575 0 0 1 .748.281l.001.003a.576.576 0 0 1-1.049.476l-.001-.003h.013zm2.428-2.26a.3.3 0 1 1-.153.395l-.001-.002a.3.3 0 0 1 .152-.394l.002-.001zm-1.29-1.375a.2.2 0 0 1 .281.183a.2.2 0 0 1-.381.086l-.001-.001a.2.2 0 0 1 .098-.266l.001-.001zM8.453 1.14c.1-.261.993-.162 1.995.226s1.729.909 1.63 1.17s-1 .164-2-.221s-1.734-.91-1.632-1.176z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Jenkins(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.698 19.922a4.456 4.456 0 0 1-.292 1.531l.01-.031L20.62 24H.839l-.328-.94a7.277 7.277 0 0 1-.483-1.45l-.009-.05a.823.823 0 0 1 .212-.751c.189-.166.405-.307.64-.415l.016-.007c.208-.098.378-.191.542-.292l-.024.014q.701-.469 1.873-1.172l1.031-.609c.381-.166.707-.387.985-.657l-.001.001q.187-.281-.281-.984l-.187-.328a3.86 3.86 0 0 1-.374-1.393l-.001-.013a2.613 2.613 0 0 1-1.822-1.162l-.006-.01a4.881 4.881 0 0 1-.7-1.799l-.004-.03a3.386 3.386 0 0 1 .196-2.131l-.009.022l.094-.14a1.15 1.15 0 0 0 .162-.367l.002-.008l.001-.048c0-.183-.027-.361-.076-.528l.003.013a2.721 2.721 0 0 1-.047-.417v-.005Q2.15 5.391 3.647 4.688a3.645 3.645 0 0 1 1.12-1.778l.004-.004A9.775 9.775 0 0 1 6.542 1.76l.056-.026l.703-.375a9.128 9.128 0 0 1 1.574-.73l.066-.02l.234-.094a9.847 9.847 0 0 1 3.183-.517c.348 0 .693.018 1.032.052l-.043-.004a7.26 7.26 0 0 1 3.868 1.282l-.024-.016c.529.429 1.004.874 1.445 1.35l.008.009c.46.456.85.98 1.155 1.556l.017.035a11.078 11.078 0 0 1 .94 4.746v-.012a17.697 17.697 0 0 1-.687 5.282l.031-.125c-.16.593-.336 1.092-.544 1.574l.028-.074l-.234.375q-.469.75-.399.984c.308.437.672.809 1.088 1.115l.013.009l.703.609q.893.753.94 1.175zM8.527 1.595a8.657 8.657 0 0 0-2.468 1.144l.031-.02a4.358 4.358 0 0 0-1.725 1.803l-.011.025c.25-.057.462-.121.667-.199l-.034.011c.201-.08.433-.131.676-.14h.238a.926.926 0 0 0 .428-.001l-.006.001q.234-.047.703-.843c.117-.224.242-.415.381-.595L7.4 2.79a2.72 2.72 0 0 1 .459-.392l.01-.006c.175-.12.328-.243.471-.377l-.002.002l.14-.047q.281-.094.281-.281q-.139-.141-.232-.094M3.605 5.95a2.65 2.65 0 0 0-.515 1.447v.006a10.438 10.438 0 0 0 .051 1.688l-.004-.047v.094a1.882 1.882 0 0 1 1.299-.28l-.01-.001c.453.057.849.262 1.148.563c.289.279.469.671.469 1.104v.022v-.001q.281 0 .328-.328a2.795 2.795 0 0 0-.147-.629l.006.02l-.094-.281a3.851 3.851 0 0 1-.046-1.095l-.001.016c.014-.417.056-.813.125-1.2l-.007.051q.117-.773.164-1.196v-.055c0-.4-.069-.784-.195-1.141l.007.024a3.802 3.802 0 0 0-2.576 1.217l-.002.003zm1.734 9.379c.16.623.336 1.146.545 1.651l-.03-.081c.196.51.432.951.716 1.356l-.013-.02c.332.435.8.751 1.339.887l.017.004c.542.112 1.188.196 1.845.233l.035.002a.843.843 0 0 1 .702-.422h.001a2.713 2.713 0 0 1 1.004.098l-.019-.005a5.594 5.594 0 0 1-1.866-1.535l-.009-.012a9.008 9.008 0 0 1-.825-1.004l-.019-.028a2.685 2.685 0 0 1-.421-1.208l-.001-.011l.14.14a18.68 18.68 0 0 0 1.698 2.121l-.011-.012a6.475 6.475 0 0 0 2.341 1.44l.045.014a4.004 4.004 0 0 0 2.513-.055l-.028.009a6.967 6.967 0 0 0 2.175-1.092l-.018.013c.393-.281.71-.642.932-1.062l.008-.017l.234-.328a9.66 9.66 0 0 0 1.438-3.309l.012-.065c.248-.995.39-2.138.39-3.314c0-.236-.006-.47-.017-.703l.001.033v-.469q-.047-.94-.094-1.359a3.205 3.205 0 0 0-.336-1.143l.008.018a1.993 1.993 0 0 0-1.111-.957l-.014-.004a1.136 1.136 0 0 0-1.313.212a.724.724 0 0 1 .347-.866l.004-.002a1.473 1.473 0 0 1 1.206-.114l-.01-.003l-.281-.328q-.563-.75-.89-1.125a4.033 4.033 0 0 0-1.104-.88l-.021-.01A7.616 7.616 0 0 0 13.671.99l-.04-.005a4.819 4.819 0 0 0-1.122-.129a4.92 4.92 0 0 0-2.003.423l.032-.012Q7.397 2.814 6.647 4.548a2.618 2.618 0 0 1 .234 1.322l.001-.009q-.047.422-.234 1.359t-.234 1.408c.018.416.085.809.196 1.183l-.009-.035c.099.325.166.701.187 1.089l.001.012q-.234.422-1.219.375a3.311 3.311 0 0 0-.383-1.095l.009.016a.982.982 0 0 0-.792-.562H4.4a1.502 1.502 0 0 0-1.289.422a1.852 1.852 0 0 0-.586 1.261v.004l-.001.077c0 .642.212 1.235.569 1.712l-.005-.007c.258.557.812.937 1.455.937l.043-.001h-.002c.216-.059.404-.156.566-.284l-.003.002c.153-.128.346-.213.558-.234h.004a.428.428 0 0 1 .071.402l.001-.003a1.19 1.19 0 0 1-.302.373l-.002.002a.7.7 0 0 0-.256.37l-.001.005c.003.244.045.478.121.696l-.005-.015zm10.406 5.858a.72.72 0 0 0-.022.356l-.001-.005a1.016 1.016 0 0 1-.025.406l.002-.007c.711.19 1.53.309 2.375.328h.012a1.685 1.685 0 0 0 .186-1.088l.001.01v-.046c0-.331-.032-.655-.094-.969l.005.032a1.176 1.176 0 0 0-.374-.749l-.001-.001a1.031 1.031 0 0 0-.802.05l.006-.003a4.81 4.81 0 0 0-.9.57l.009-.007q-.471.281-.616.375l.047.094c.051.11.085.237.094.371v.003a1.108 1.108 0 0 1 .896.097l-.005-.003h-.192c-.225 0-.434.07-.606.19l.003-.002zm-.795-.94l.422-.375c.174-.166.375-.307.595-.415l.014-.006a4.778 4.778 0 0 0-.948.148l.033-.007q-.633.14-.961.187a1.652 1.652 0 0 1-.901-.142l.01.004l-.094.094c.23.19.504.336.805.418l.015.004c.276.06.592.094.917.094l.096-.001h-.005zm-.563 1.547h-.047a.645.645 0 0 0 .681-.073l-.001.001a.658.658 0 0 0 .187-.541v.003a.531.531 0 0 0-.301-.467l-.003-.001a.61.61 0 0 0-.656.14a1.214 1.214 0 0 0 .049.571l-.003-.009c.027.15.059.278.099.403l-.006-.023zm-4.918-1.219c.009-.154.035-.298.076-.435l-.003.014a1.94 1.94 0 0 0 .072-.367l.001-.008a3.727 3.727 0 0 0-.61-.048l-.098.001h.005a3.609 3.609 0 0 1-1.289-.29l.023.009q-.281.281.516.68c.376.213.814.37 1.28.445l.022.003zM1.59 23.297h8.437l-.094-.328a7.412 7.412 0 0 1-.371-1.604l-.004-.037l-.328-.14a16.395 16.395 0 0 1-2.401-1.349l.057.037l-.375-.422q-.75-.891-.843-.846a20.971 20.971 0 0 0-4.816 2.981l.035-.028c.218.448.456 1.021.664 1.609zm11.624 0h.422l-.14-.187l-.187.14zm.703-1.172v-.469a2.192 2.192 0 0 0-1.078-.281c.142-.079.308-.137.484-.163l.008-.001c.202-.035.383-.092.553-.17l-.014.006v-.187a1.88 1.88 0 0 1-.801-.332l.005.004l-.234-.14a6.185 6.185 0 0 0-2.025-.744l-.038-.006q-.562-.093-.703.609a4.98 4.98 0 0 0-.069.839c0 .167.008.332.023.494l-.002-.021v.047c.074.525.205.999.389 1.446l-.014-.039l.187.469q.047.281.211.328l.057.001c.241 0 .47-.052.677-.145l-.01.004c.338-.107.626-.226.9-.367l-.032.015a6.25 6.25 0 0 0 .92-.685l-.006.005c.181-.191.382-.361.6-.508l.013-.008zm2.485 1.172l.047-.563a1.291 1.291 0 0 1-.751-.329l.001.001a1.267 1.267 0 0 0-.413-.209l-.009-.002h-.007a.772.772 0 0 0-.415.12l.003-.002a.748.748 0 0 1-.614.093l.005.001l-.281.328q.187.234.422.563h.796a.268.268 0 0 1 .094-.187a.293.293 0 0 1 .211-.094h.005c.075 0 .141.037.181.093v.001a.315.315 0 0 1 .072.187v.001zm.422 0h1.687a1.258 1.258 0 0 0-.694-.583l-.009-.003a1.168 1.168 0 0 0-.947.075l.006-.003a3.134 3.134 0 0 0-.044.513zm3.89-1.736l.047-.234c.09-.26.141-.559.141-.871l-.001-.073v.004q-.047-.324-.75-.886l-.231-.187l-.656-.656q-.566-.516-.895-.845a1.994 1.994 0 0 1-.509.511l-.007.004a2.27 2.27 0 0 0-.462.421l-.003.004a.691.691 0 0 1 .611-.061l-.005-.002a.703.703 0 0 1 .423.433l.001.005a3.384 3.384 0 0 1 .375 1.717v-.007a3.383 3.383 0 0 1-.29 1.592l.009-.021c.051.159.14.293.257.398l.001.001c.108.096.195.213.255.345l.003.006l-.094.14h1.313c.17-.471.332-1.053.454-1.649l.015-.086zM9.277 6.095c.204-.432.507-.788.882-1.049l.009-.006a1.807 1.807 0 0 1 1.192-.446h.028h-.001c.215.059.398.174.538.327l.001.001a.382.382 0 0 1 .023.517v-.001q-1.64-.375-2.531 1.547q-.417-.095-.14-.891zm8.624 3c.2-.074.431-.117.672-.117h.033h-.002a1.34 1.34 0 0 0 .71-.168l-.007.004l-.14-.328a3.071 3.071 0 0 1-.281-1.412v.005h.047l.234.469c.247.569.499 1.046.782 1.502l-.028-.049c-.113.267-.332.47-.602.561l-.007.002a1.508 1.508 0 0 1-.971.138l.009.001a.544.544 0 0 1-.446-.612v.003zm-4.499.094q-.94-1.453-.703-2.109q.094.187.281.609c.12.367.313.682.563.941l-.001-.001l.094.094c.102.099.195.206.277.321l.005.007a.845.845 0 0 1-.002.568l.002-.006a.81.81 0 0 1-.231.232l-.003.002a1.222 1.222 0 0 1-.32.185l-.008.003a2.299 2.299 0 0 1-1.352-.028l.016.005q-.723-.211-.352-.773c.172.014.331.047.483.098l-.014-.004c.207.049.448.083.695.094h.008a.882.882 0 0 0 .516-.235zm1.265 3.281c.044.109.077.235.093.367l.001.008c.019.158.069.3.143.427l-.003-.005a.487.487 0 0 0 .326.187h.002a1.948 1.948 0 0 0 1.087-.146l-.012.005c.42-.122.78-.265 1.123-.436l-.038.017l.563-.281a.853.853 0 0 0 .139-.709l.001.006a1.357 1.357 0 0 0-.373-.514l-.002-.002l-.328-.38a11.43 11.43 0 0 1-1.101-1.451l-.027-.046a3.743 3.743 0 0 1-.516-1.824v-.004q.281-.187.375.328v.14a12.83 12.83 0 0 0 2.203 2.953l.516.516a1.368 1.368 0 0 1-.144.711l.004-.008a1.463 1.463 0 0 0-.094.417v.005c-.471.32-1.01.618-1.576.865l-.065.025a4.24 4.24 0 0 1-1.617.315c-.19 0-.377-.012-.56-.036l.022.002a1.061 1.061 0 0 1-.351-.675v-.005a.913.913 0 0 1 .211-.774l-.001.001zm-9.422-.422a1.014 1.014 0 0 1-.372-.462l-.002-.007q-.187-.516-.446-.539t-.352.446a2.819 2.819 0 0 0-.069.629c0 .125.008.249.023.37l-.001-.014v.234a1.206 1.206 0 0 1-.329-.83l.001-.04v.002a1.923 1.923 0 0 1 .146-1.02l-.005.013q-.187-.094-.469.14l-.14.14q.094-.75.68-.843a1.064 1.064 0 0 1 1.029.443l.002.003a1.455 1.455 0 0 1 .255 1.346l.003-.01zm14.015 1.359a4.936 4.936 0 0 1-.939 1.265l-.001.001c-.417.39-.974.636-1.587.656h-.004a3.823 3.823 0 0 1 .001-.909l-.002.018a6.295 6.295 0 0 0 1.075-.296l-.044.014c.325-.128.603-.27.865-.434l-.021.012zm-6.654.563a8.048 8.048 0 0 0 2.864.517c.312 0 .62-.017.923-.051l-.037.003c.031.201.049.433.049.67l-.002.132v-.006v.14a6.874 6.874 0 0 1-2.389-.295l.049.013a1.947 1.947 0 0 1-1.45-1.111l-.005-.012zm4.312 2.203a.732.732 0 0 1-.417.514l-.005.002c-.418.138-.9.222-1.4.234h-.006a3.151 3.151 0 0 1-1.567-.242l.021.008a3.73 3.73 0 0 1-.882-.927l-.009-.014l-.189-.279l-.137-.14a.893.893 0 0 1-.255-.346l-.002-.006q-.023-.117.399-.117a2.103 2.103 0 0 0 1.627 1.215l.011.001c.523.073 1.128.115 1.742.115c.261 0 .521-.008.778-.022l-.036.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Jira(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.97 5.826a.455.455 0 0 0-.285-.1h-.015h.001a.727.727 0 0 0-.177.027l.005-.001c-.86.342-1.909.665-2.988.911l-.151.029a.626.626 0 0 0-.43.374l-.001.004c-.722 1.965-2.557 3.844-4.48 5.811a.505.505 0 0 1-.359.172h-.001a.47.47 0 0 1-.361-.169l-.001-.001C6.799 10.913 4.965 9.05 4.239 7.07a.648.648 0 0 0-.429-.375l-.004-.001A24.752 24.752 0 0 1 .499 5.683l.172.058a.36.36 0 0 0-.149-.031l-.028.001h.001h-.009a.474.474 0 0 0-.292.101l.001-.001a.509.509 0 0 0-.19.44v-.002a10.302 10.302 0 0 0 1.997 4.941l-.019-.026a36.782 36.782 0 0 0 3.909 4.441l.011.01c2.544 2.6 4.956 5.06 5.173 7.905a.501.501 0 0 0 .495.465h2.992a.534.534 0 0 0 .361-.157a.504.504 0 0 0 .131-.341v-.022v.001a11.318 11.318 0 0 0-2.221-5.87l.021.03c-.29-.42-.594-.837-.925-1.244a.368.368 0 0 1 .027-.476l.3-.306a36.575 36.575 0 0 0 3.852-4.359l.067-.094a10.135 10.135 0 0 0 1.975-4.866l.005-.05l.001-.024a.482.482 0 0 0-.186-.381zM5.857 17.843a.448.448 0 0 0-.334-.149h-.01h.001a.5.5 0 0 0-.405.232l-.001.002a10.917 10.917 0 0 0-1.993 5.514l-.002.037v.017c0 .132.05.253.132.344c.085.098.21.16.349.16h.013h-.001h2.993a.492.492 0 0 0 .491-.463v-.001a6.52 6.52 0 0 1 .855-2.704l-.017.031a.697.697 0 0 0-.06-.754l.001.002a27.47 27.47 0 0 0-2.012-2.272l.002.002z"/><path fill="currentColor" d="M10.757 7.229a.453.453 0 0 1 .398.653l.001-.003a8.18 8.18 0 0 1-1.749 2.559l-.001.001a.475.475 0 0 1-.327.129h-.021h.001h-.015a.49.49 0 0 1-.332-.13a8.09 8.09 0 0 1-1.73-2.511l-.02-.052a.496.496 0 0 1-.041-.184v-.002c0-.25.198-.453.446-.462h.001c.216 0 .766.058 1.706.058s1.49-.058 1.677-.058zM3.92 1.43a1.575 1.575 0 1 0 1.575 1.578A1.58 1.58 0 0 0 3.92 1.434zm10.321 0a1.575 1.575 0 1 0 1.575 1.575v-.001a1.58 1.58 0 0 0-1.575-1.57zM9.081 0a1.577 1.577 0 1 0 1.577 1.577v-.001A1.578 1.578 0 0 0 9.081 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Joomla(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.054 3.215V3.16c0-.871.359-1.658.937-2.221L.992.938A3.033 3.033 0 0 1 3.188 0h.029h-.001l.064-.001c.789 0 1.509.295 2.055.78L5.332.776a3.295 3.295 0 0 1 1.095 1.881l.003.021c.358-.091.77-.143 1.194-.143c.53 0 1.041.081 1.521.233l-.036-.01c.971.311 1.8.837 2.462 1.524l.002.002l-2.357 2.357a2.372 2.372 0 0 0-1.601-.749h-.032c-.53 0-1.007.227-1.339.588l-.001.001c-.346.352-.56.835-.56 1.368v.027v-.001v.03c0 .527.204 1.007.537 1.364l-.001-.001l5.36 5.36l-2.362 2.355l-5.36-5.36A4.748 4.748 0 0 1 2.47 9.191l-.006-.031a5.876 5.876 0 0 1 .063-2.88l-.01.041A2.931 2.931 0 0 1 .754 5.23l-.005-.007a3.141 3.141 0 0 1-.694-1.975v-.033v.002zM6.96 9.431l.8-.75l4.56-4.555a4.908 4.908 0 0 1 2.404-1.413l.034-.007a5.408 5.408 0 0 1 1.356-.168c.518 0 1.02.071 1.496.204l-.039-.009A3.102 3.102 0 0 1 18.636.809l.005-.004a3.155 3.155 0 0 1 2.173-.8h-.003h.015c.858 0 1.633.36 2.181.936l.001.001c.579.564.938 1.351.938 2.222l-.001.058v-.003l.001.06a3.03 3.03 0 0 1-.802 2.058l.002-.002a3.437 3.437 0 0 1-1.962 1.095l-.021.003a4.948 4.948 0 0 1 .045 2.82l.008-.035a6.125 6.125 0 0 1-1.553 2.624l-2.411-2.413c.461-.372.763-.924.8-1.548V7.86c0-.542-.226-1.03-.589-1.377l-.001-.001a1.897 1.897 0 0 0-1.378-.59h-.016h.001h-.015c-.542 0-1.03.226-1.377.589l-.001.001l-5.36 5.36zm14.253 8.194a3.188 3.188 0 0 1 1.983 1.042l.003.003c.496.539.801 1.262.801 2.056l-.001.067v-.003v.055a3.09 3.09 0 0 1-.937 2.221l-.001.001a3.092 3.092 0 0 1-2.222.938l-.058-.001h.003h-.017a3.182 3.182 0 0 1-2.025-.724l.006.004a3.014 3.014 0 0 1-1.122-1.836l-.003-.019a5.193 5.193 0 0 1-1.541.229c-.481 0-.946-.064-1.389-.184l.037.009a5.09 5.09 0 0 1-2.567-1.55l-.004-.005l2.357-2.41c.372.461.924.763 1.548.8h.021c.542 0 1.03-.226 1.377-.589l.001-.001c.364-.347.59-.836.59-1.378v-.023c0-.542-.226-1.03-.589-1.377l-.001-.001l-5.36-5.36l2.41-2.357l5.303 5.36a5.307 5.307 0 0 1 1.387 2.319l.009.038c.119.401.187.861.187 1.338s-.068.937-.196 1.372zm-4.5-2.835l-5.304 5.299a5.31 5.31 0 0 1-2.346 1.384l-.037.009c-.405.119-.871.188-1.353.188s-.947-.068-1.388-.196l.035.009A2.952 2.952 0 0 1 5.229 23.3l-.007.005a3.141 3.141 0 0 1-1.979.697h-.031h.002h-.055a3.09 3.09 0 0 1-2.221-.937l-.001-.001a3.092 3.092 0 0 1-.938-2.222v-.084c0-.752.263-1.443.702-1.985l-.005.006a2.925 2.925 0 0 1 1.748-1.095l.019-.003a5.547 5.547 0 0 1 .064-2.798l-.01.039a5.78 5.78 0 0 1 1.551-2.542l.002-.002l2.357 2.415c-.435.383-.717.93-.749 1.544v.03c0 .533.214 1.016.56 1.368c.338.363.82.59 1.354.59h.041h-.002l.049.001a1.76 1.76 0 0 0 1.316-.589l.002-.002l5.36-5.36z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Jquery(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.582 5.638c-2.206 3.17-1.931 7.294-.246 10.663c.04.08.08.16.123.24c.026.05.051.102.08.151c.015.03.033.059.049.087l.088.16l.16.275c.03.05.06.101.092.151q.091.147.186.292c.026.042.052.08.08.122c.088.131.178.262.27.39l.008.01c.014.021.031.041.046.062c.08.11.16.22.246.328l.093.118q.112.141.228.28l.087.104c.103.122.209.24.32.362l.006.006l.012.013c.106.115.214.228.32.34l.103.104a12.293 12.293 0 0 0 .364.352c.116.11.234.217.353.32c.002.002.003.004.006.005l.062.052q.16.137.32.271l.133.106q.132.106.267.209l.143.109c.099.073.199.145.3.214l.11.08l.03.022c.096.066.194.13.291.193c.042.028.08.057.126.08c.15.094.302.189.456.279l.126.071c.113.066.226.13.342.192c.062.034.126.066.189.098c.08.043.16.086.245.127l.057.027l.1.049q.192.092.388.18l.08.037q.224.098.453.189l.11.043c.141.054.284.108.427.16l.054.018c.16.055.32.106.48.156l.115.034c.128.046.292.094.461.132l.031.006c10.667 1.945 13.766-6.41 13.766-6.41c-2.602 3.39-7.222 4.285-11.6 3.289c-.166-.038-.326-.089-.489-.137l-.122-.037q-.24-.073-.47-.153l-.065-.023a13.565 13.565 0 0 1-.414-.154l-.116-.046c-.15-.059-.3-.122-.448-.186l-.09-.04c-.127-.058-.253-.115-.378-.175l-.109-.053a11.457 11.457 0 0 1-.292-.149c-.065-.034-.131-.066-.195-.102c-.118-.063-.234-.13-.35-.197l-.118-.066c-.154-.09-.306-.185-.456-.279c-.042-.026-.08-.054-.123-.08c-.109-.07-.218-.142-.325-.216c-.035-.023-.07-.05-.105-.074q-.155-.109-.307-.222l-.136-.103c-.093-.071-.186-.142-.277-.217l-.123-.099c-.116-.096-.231-.191-.345-.29l-.038-.032c-.122-.108-.24-.218-.362-.33l-.102-.098q-.133-.126-.26-.255l-.102-.102q-.16-.167-.32-.338l-.016-.017a14.24 14.24 0 0 1-.324-.369l-.083-.102a6.829 6.829 0 0 1-.233-.288a3.05 3.05 0 0 1-.086-.106c-.092-.119-.182-.24-.27-.358c-2.43-3.314-3.303-7.886-1.36-11.64z"/><path fill="currentColor" d="M8.429 2.966a7.592 7.592 0 0 0-1.202 4.121a7.63 7.63 0 0 0 .958 3.714l-.02-.039c.235.457.469.843.727 1.21l-.022-.032c.244.387.515.723.82 1.024l.001.001c.115.126.235.25.358.371l.094.094c.119.115.24.229.366.339c.005.004.01.01.015.013c.138.122.282.24.427.353l.097.076c.146.112.294.222.446.326l.013.01c.067.046.136.09.204.134c.033.022.064.044.097.064c.109.07.219.138.331.203l.046.027c.097.056.194.11.293.16c.034.019.069.036.104.054l.205.106l.03.014c.141.07.282.136.427.199c.031.014.062.026.094.039c.115.049.233.097.35.142l.15.055c.106.039.214.076.32.112l.146.046c.126.048.281.094.441.13l.024.004c8.236 1.365 10.138-4.977 10.138-4.977c-1.714 2.468-5.034 3.646-8.575 2.727a9.35 9.35 0 0 1-.606-.181c-.11-.035-.218-.074-.326-.113l-.147-.054a13.534 13.534 0 0 1-.351-.142c-.032-.014-.064-.026-.094-.04a9.92 9.92 0 0 1-.43-.199c-.073-.036-.144-.074-.215-.11l-.124-.064q-.138-.075-.273-.154l-.065-.037q-.167-.098-.33-.202l-.099-.066c-.072-.046-.143-.094-.214-.142a11.036 11.036 0 0 1-.445-.326l-.1-.08a9.972 9.972 0 0 1-3.34-4.723l-.02-.07a7.297 7.297 0 0 1-.373-2.325c0-1.338.353-2.594.972-3.679l-.019.037z"/><path fill="currentColor" d="M14.188.97a4.752 4.752 0 0 0-.804 2.659c0 .715.156 1.393.436 2.002l-.012-.03a6.558 6.558 0 0 0 3.699 3.484l.045.015c.068.026.135.048.205.071l.09.028c.078.029.176.058.276.081l.017.003c4.547.88 5.781-2.334 6.11-2.806c-1.081 1.556-2.896 1.929-5.124 1.389a5.513 5.513 0 0 1-.539-.167a6.342 6.342 0 0 1-.679-.283l.039.017a6.604 6.604 0 0 1-1.136-.697l.016.012c-1.994-1.513-3.232-4.4-1.931-6.749z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Jsfiddle(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M30.857 10.903a6.897 6.897 0 0 1 3.062 2.471l.015.023a6.673 6.673 0 0 1 1.173 3.796a6.709 6.709 0 0 1-2.02 4.806l-.001.001a6.683 6.683 0 0 1-4.776 2.001L28.194 24h.006q-.069 0-.197-.009t-.18-.009H6.96a7.433 7.433 0 0 1-4.937-2.152l.001.001a6.463 6.463 0 0 1-2.023-4.814v-.043a6.69 6.69 0 0 1 .96-3.462l-.017.03a7.072 7.072 0 0 1 2.487-2.502l.033-.018a4.67 4.67 0 0 1-.206-1.383v-.024v.001l-.001-.071c0-1.292.539-2.459 1.404-3.287l.002-.002a4.662 4.662 0 0 1 3.416-1.389h-.005h.03c1.104 0 2.121.374 2.93 1.003l-.011-.008a10.401 10.401 0 0 1 3.772-4.226l.042-.025A10.102 10.102 0 0 1 20.34-.001h.102h-.005h.057c3.805 0 7.134 2.036 8.957 5.078l.026.047a9.855 9.855 0 0 1 1.397 5.089v.075v-.004q0 .103-.009.309t-.009.308zm-22.835 4.56a4.097 4.097 0 0 0 1.434 3.304l.006.005a5.305 5.305 0 0 0 3.575 1.217h-.009h.035a5.731 5.731 0 0 0 4.079-1.697q-.274-.343-.814-.969l-.746-.866a3.493 3.493 0 0 1-2.464 1.114h-.005l-.057.001a2.349 2.349 0 0 1-1.546-.578l.003.002a2.022 2.022 0 0 1-.002-2.99l.002-.001c.404-.362.94-.584 1.528-.584h.043h-.002c.53.001 1.029.134 1.465.368l-.017-.008a4.78 4.78 0 0 1 1.249.941l.002.002q.56.583 1.114 1.286t1.183 1.406c.411.462.842.882 1.301 1.27l.019.015a5.565 5.565 0 0 0 1.624.931l.039.012c.605.228 1.305.36 2.035.36h.047h-.002l.148.002a5.184 5.184 0 0 0 3.365-1.234l-.008.006a4.453 4.453 0 0 0 1.433-3.276a4.45 4.45 0 0 0-1.438-3.281l-.003-.003a5.28 5.28 0 0 0-3.557-1.216h.008l-.128-.001a5.63 5.63 0 0 0-4.003 1.664l1.594 1.851a3.444 3.444 0 0 1 2.43-1.097h.031c.592 0 1.135.214 1.554.569l-.004-.003c.419.33.686.837.686 1.407v.035v-.002l.002.087c0 .583-.244 1.109-.635 1.481l-.001.001a2.257 2.257 0 0 1-1.529.593L23 17.585h.004a2.96 2.96 0 0 1-1.43-.368l.015.008a4.876 4.876 0 0 1-1.232-.941l-.002-.002q-.56-.583-1.12-1.286t-1.191-1.406a12.427 12.427 0 0 0-1.309-1.27l-.019-.016a5.703 5.703 0 0 0-1.615-.93l-.04-.013a5.588 5.588 0 0 0-1.992-.36h-.041h.002a5.353 5.353 0 0 0-3.558 1.216l.009-.007a3.994 3.994 0 0 0-1.459 3.258v-.008z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Json(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.012 23.968c.504-.001.998-.035 1.483-.1l-.057.006a11.775 11.775 0 0 0 3.222-.893l-.076.03a12.45 12.45 0 0 0 3.795-2.545l-.002.002a12.02 12.02 0 0 0 2.409-3.355l.031-.072c.487-1.01.845-2.182 1.013-3.415l.007-.061a12.547 12.547 0 0 0-.054-3.571l.009.071a11.908 11.908 0 0 0-.823-2.866l.03.078a11.052 11.052 0 0 0-.881-1.661l.027.044a12.337 12.337 0 0 0-4.74-4.268l-.07-.032a12.307 12.307 0 0 0-2.081-.845l-.091-.024h-.005c.119.063.24.132.345.2s.239.146.351.225a8.952 8.952 0 0 1 1.554 1.325l.005.005a9.803 9.803 0 0 1 2.201 3.972l.017.069c.261.888.44 1.919.498 2.981l.002.036a12.496 12.496 0 0 1-.052 2.286l.005-.057a8.718 8.718 0 0 1-1.25 3.637l.022-.038a7.531 7.531 0 0 1-1.881 2.07l-.019.014a5.444 5.444 0 0 1-3.43 1.207c-.356 0-.703-.034-1.04-.099l.034.005a3.934 3.934 0 0 1-.772-.268l.024.01a4.684 4.684 0 0 1-.793-.459l.014.01a6.108 6.108 0 0 1-1.237-1.098l-.007-.008a6.46 6.46 0 0 1-1.033-1.704l-.016-.043a7.238 7.238 0 0 1-.494-2.534v-.006a7.083 7.083 0 0 1 .872-3.707l-.019.037a6.462 6.462 0 0 1 1.641-1.91l.014-.011c.203-.159.431-.317.669-.462l.031-.017l.016-.007a4.987 4.987 0 0 0-1.455-.12l.014-.001a4.971 4.971 0 0 0-1.256.249l.035-.01a5.687 5.687 0 0 0-1.047.465l.028-.015a6.18 6.18 0 0 0-.746.505l.013-.01c-.211.18-.42.359-.615.555a7.821 7.821 0 0 0-2.017 4.214l-.005.042a13.51 13.51 0 0 0-.131 1.902c0 .369.014.735.043 1.097l-.003-.048c.095 1.486.433 2.869.976 4.143l-.031-.082a8.53 8.53 0 0 0 1.355 2.185l-.01-.012a8.174 8.174 0 0 0 3.078 2.211l.054.02c.616.26 1.33.45 2.076.536l.037.003c.054.01.116.015.179.015h.031h-.002l-.021-.032zm-2.933-.4a4.17 4.17 0 0 1-.469-.238l.019.01c-.149-.083-.3-.168-.449-.259a8.569 8.569 0 0 1-1.587-1.273l-.001-.001a9.495 9.495 0 0 1-2.232-4.014l-.015-.067a14.695 14.695 0 0 1-.494-3.136l-.001-.036a12.26 12.26 0 0 1 .051-2.137l-.005.055a9.034 9.034 0 0 1 1.147-3.624l-.023.044a6.942 6.942 0 0 1 1.767-2.039l.016-.012a6.188 6.188 0 0 1 1.59-.886l.043-.014a5.262 5.262 0 0 1 1.827-.321c.26 0 .515.018.765.054l-.029-.003c.361.049.684.128.994.236l-.036-.011c.029 0 .029 0 .045.03c.015.015.045.015.06.03c.045.016.1.045.165.074c.275.144.504.284.723.435l-.023-.015a6.334 6.334 0 0 1 2.084 2.462l.016.038c.393.864.625 1.873.631 2.936v.002l.002.174a7.104 7.104 0 0 1-.666 3.015l.018-.044a6.464 6.464 0 0 1-2.607 2.815l-.031.017c.09.03.18.045.271.075c.206.045.443.072.686.074h.002a5.645 5.645 0 0 0 3.94-1.647c.194-.179.371-.37.532-.574l.008-.011c.209-.244.411-.512.596-.793l.018-.03c.211-.326.416-.705.592-1.1l.022-.054c.21-.476.386-1.033.501-1.612l.008-.051a12.066 12.066 0 0 0 .146-3.068l.003.042a11.227 11.227 0 0 0-1.868-5.765l.026.041a9.142 9.142 0 0 0-.565-.731l.01.012a10.06 10.06 0 0 0-1.047-1.02l-.015-.012a8.474 8.474 0 0 0-1.319-.893l-.044-.022a9.516 9.516 0 0 0-1.623-.585l-.069-.015l-.3-.06c-.209-.03-.42-.044-.634-.06a8.619 8.619 0 0 0-1.044.018l.029-.002c-.777.054-1.498.174-2.193.357l.081-.018c-3.449.924-6.245 3.214-7.822 6.24l-.032.068a11.717 11.717 0 0 0-1.269 5.342c0 .552.038 1.095.11 1.626l-.007-.062c.129 1.103.403 2.109.805 3.048l-.027-.071c.261.641.538 1.183.854 1.698l-.029-.05a12.242 12.242 0 0 0 4.447 4.208l.063.032c.612.349 1.325.668 2.069.919l.088.026c.226.074.45.149.689.209z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Justify(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.334 2.666h26.665l.037.001a1.334 1.334 0 1 0 0-2.668L27.997 0h.002H1.334a1.334 1.334 0 0 0-.002 2.666zm26.665 2.667H1.334a1.334 1.334 0 0 0-.002 2.666h26.667l.037.001a1.334 1.334 0 1 0 0-2.668l-.039.001zm0 5.334H1.334a1.334 1.334 0 0 0-.002 2.666h26.667a1.334 1.334 0 0 0 .002-2.666zm0 10.666H1.334a1.334 1.334 0 0 0-.002 2.666h26.667a1.334 1.334 0 0 0 .002-2.666zm0-5.333H1.334a1.334 1.334 0 0 0-.002 2.666h26.667A1.334 1.334 0 0 0 28.001 16z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Key(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.815 14.632l-4.031 4.031H7.115v2.668H4.447v2.668H0v-4.447l9.368-9.368a7.41 7.41 0 0 1-.474-2.632a7.554 7.554 0 1 1 4.869 7.062zm7.532-9.31v-.003a2.668 2.668 0 1 0-2.669 2.668h.001a2.669 2.669 0 0 0 2.668-2.665"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Keyboard(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M33 24H3a3 3 0 0 1-3-3V3a3 3 0 0 1 3-3h30a3 3 0 0 1 3 3v18a3 3 0 0 1-3 3M8 7.25V4.749a.753.753 0 0 0-.749-.75H4.75a.753.753 0 0 0-.75.749v2.501a.753.753 0 0 0 .749.75H7.25A.753.753 0 0 0 8 7.25m6 0V4.749a.753.753 0 0 0-.749-.75H10.75a.753.753 0 0 0-.75.749v2.501a.753.753 0 0 0 .749.75h2.5a.753.753 0 0 0 .75-.749zm6 0V4.749a.753.753 0 0 0-.749-.75H16.75a.753.753 0 0 0-.75.749v2.501a.753.753 0 0 0 .749.75h2.5a.751.751 0 0 0 .75-.749zm6 0V4.749a.753.753 0 0 0-.749-.75H22.75a.753.753 0 0 0-.75.749v2.501a.753.753 0 0 0 .749.75h2.5a.753.753 0 0 0 .75-.749zm6 0V4.749a.753.753 0 0 0-.749-.75h-2.503a.752.752 0 0 0-.749.749v2.501a.752.752 0 0 0 .75.75h2.503a.751.751 0 0 0 .746-.749zm-24 6v-2.501a.753.753 0 0 0-.749-.75H4.75a.753.753 0 0 0-.75.749v2.5a.753.753 0 0 0 .749.75H7.25a.753.753 0 0 0 .75-.749zm6 0v-2.501a.753.753 0 0 0-.749-.75H10.75a.753.753 0 0 0-.75.749v2.5a.753.753 0 0 0 .749.75h2.5a.753.753 0 0 0 .75-.749zm6 0v-2.501a.753.753 0 0 0-.749-.75H16.75a.753.753 0 0 0-.75.749v2.5a.753.753 0 0 0 .749.75h2.5a.75.75 0 0 0 .75-.749zm6 0v-2.501a.753.753 0 0 0-.749-.75H22.75a.753.753 0 0 0-.75.749v2.5a.753.753 0 0 0 .749.75h2.5a.75.75 0 0 0 .75-.749zm6 0v-2.501a.752.752 0 0 0-.748-.75h-2.503a.751.751 0 0 0-.75.749v2.5a.752.752 0 0 0 .75.75h2.503a.752.752 0 0 0 .748-.749zm-24 6v-2.501a.753.753 0 0 0-.749-.75H4.75a.753.753 0 0 0-.75.749v2.5a.751.751 0 0 0 .749.75H7.25a.753.753 0 0 0 .75-.749zm18 0v-2.501a.753.753 0 0 0-.749-.75H10.75a.753.753 0 0 0-.75.749v2.5a.753.753 0 0 0 .749.75H25.25a.753.753 0 0 0 .75-.749zm6 0v-2.501a.753.753 0 0 0-.749-.75h-2.503a.751.751 0 0 0-.75.749v2.5a.752.752 0 0 0 .75.75h2.503a.752.752 0 0 0 .748-.749z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Kickstarter(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.879 4.397v3.801l4.444-6.318A4.233 4.233 0 0 1 14.218.337l.028-.01a3.836 3.836 0 0 1 2.248-.183l-.025-.005a4.102 4.102 0 0 1 1.956 1.046l-.001-.001a3.59 3.59 0 0 1 1.177 2.37l.001.013a3.613 3.613 0 0 1-.65 2.448l.008-.011l-4.015 5.729l4.872 6.157a3.454 3.454 0 0 1 .722 2.449l.001-.013a3.83 3.83 0 0 1-1.096 2.49v-.001a3.777 3.777 0 0 1-1.851 1.039l-.025.005a4.265 4.265 0 0 1-2.092-.034l.03.007a3.203 3.203 0 0 1-1.629-1.013l-.004-.004l-5.993-7.389v4.069a4.765 4.765 0 0 1-1.129 3.431l.006-.007a3.576 3.576 0 0 1-2.742 1.074h.008a3.634 3.634 0 0 1-2.968-1.254l-.004-.005a4.872 4.872 0 0 1-1.044-3.25v.01V4.342a4.616 4.616 0 0 1 1.051-3.114l-.006.008A3.614 3.614 0 0 1 3.98.004h-.009a3.753 3.753 0 0 1 2.785 1.072l-.001-.001a4.595 4.595 0 0 1 1.127 3.334l.001-.014z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Krw(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8.762 18.187l1.381-5.097h-2.71l1.28 5.114a.067.067 0 0 1 .017.046v.012a.07.07 0 0 0 .017.046a.385.385 0 0 1 .009-.062v.002a.341.341 0 0 0 .007-.059zm1.977-7.28l.597-2.182H6.358l.546 2.182zm3.273 0h2.37l-.597-2.182h-1.194zm7.654 7.295l1.33-5.114h-2.762l1.381 5.097a.385.385 0 0 0 .009.062v-.002a.113.113 0 0 0 .026.06a.226.226 0 0 1 .009-.053v.002a.25.25 0 0 0 .008-.049v-.001zm1.892-7.295l.56-2.182h-5.061l.579 2.182zm6.988.546v1.112c0 .29-.235.525-.525.525h-.022H30h-3.631l-2.795 10.5a.51.51 0 0 1-.498.41l-.032-.001h.001h-2.71l-.03.001a.509.509 0 0 1-.498-.406l-.001-.003l-2.83-10.5h-3.562l-2.846 10.502a.51.51 0 0 1-.498.41l-.032-.001h.001h-2.715a.506.506 0 0 1-.328-.12l.001.001a.498.498 0 0 1-.179-.286l-.001-.003l-2.727-10.5H.523a.525.525 0 0 1-.525-.525v-.022v.001v-1.112c0-.29.235-.525.525-.525h.022h-.001h2.983l-.56-2.182H.523a.525.525 0 0 1-.525-.525v-.022v.001v-1.112c0-.29.235-.525.525-.525h.022h-.001h1.854L.878.682A.485.485 0 0 1 .964.201L.963.202a.55.55 0 0 1 .425-.201l.025.001h-.001h2.336a.48.48 0 0 1 .529.406v.003l1.534 6.136h6.119L13.584.411a.51.51 0 0 1 .53-.409h-.001h2.148l.03-.001c.246 0 .451.175.498.406l.001.003l1.67 6.136h6.222L26.268.41a.481.481 0 0 1 .531-.409h-.002h2.35c.173 0 .327.08.427.204l.001.001a.486.486 0 0 1 .085.483l.001-.003L28.11 6.55h1.91c.29 0 .525.235.525.525v.022v-.001v1.112c0 .29-.235.525-.525.525h-.022h.001h-2.471l-.579 2.182h3.071c.29 0 .525.235.525.525v.022v-.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Laboratory(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.171 19.68L14.819 8.369V2.962h1.708V0H6.098v2.965H7.82v5.407L.454 19.68A2.792 2.792 0 0 0 2.791 24h17.034a2.8 2.8 0 0 0 2.34-4.331l.007.011zm-.905 2.302a1.633 1.633 0 0 1-1.434.854H2.791a1.635 1.635 0 0 1-1.37-2.531l-.004.006l7.549-11.6V2.96h4.686v5.754l7.541 11.6c.17.251.272.561.272.895c0 .285-.074.553-.204.785l.004-.008z"/><path fill="currentColor" d="M14.412 12.351H8.221l-5.655 8.698a.287.287 0 0 0-.012.299l-.001-.001c.05.087.142.145.248.146h17.032a.283.283 0 0 0 .247-.145l.001-.001a.283.283 0 0 0-.013-.298l.001.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Language(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.4 18.4H.9a.9.9 0 0 1-.9-.9V7.3a.9.9 0 0 1 .9-.9h10.5zm-4.525-2.72c.058.187.229.32.431.32h.854a.45.45 0 0 0 .425-.597l.001.003l-2.15-6.34a.451.451 0 0 0-.426-.306H4.791a.451.451 0 0 0-.425.302l-.001.003l-2.154 6.34a.45.45 0 0 0 .426.596h.856a.45.45 0 0 0 .431-.323l.001-.003l.342-1.193h2.258l.351 1.195zM5.41 10.414s.16.79.294 1.245l.406 1.408H4.68l.415-1.408c.131-.455.294-1.245.294-1.245zM23.1 18.4H12.6v-12h10.5a.9.9 0 0 1 .9.9v10.2a.9.9 0 0 1-.9.9m-1.35-8.55h-2.4v-.601a.45.45 0 0 0-.45-.45h-.601a.45.45 0 0 0-.45.45v.601h-2.4a.45.45 0 0 0-.45.45v.602c0 .248.201.45.45.45h4.281a5.861 5.861 0 0 1-1.126 1.621l.001-.001a7.105 7.105 0 0 1-.637-.764l-.014-.021a.452.452 0 0 0-.602-.129l.002-.001l-.273.16l-.24.146a.45.45 0 0 0-.139.642l-.001-.001c.253.359.511.674.791.969l-.004-.004c-.28.216-.599.438-.929.645l-.05.029a.45.45 0 0 0-.159.61l-.001-.002l.298.52a.451.451 0 0 0 .628.159l-.002.001c.507-.312.94-.619 1.353-.95l-.026.02c.387.313.82.62 1.272.901l.055.032a.449.449 0 0 0 .626-.158l.001-.002l.298-.52a.45.45 0 0 0-.153-.605l-.002-.001a11.5 11.5 0 0 1-1.004-.696l.027.02a6.708 6.708 0 0 0 1.586-2.572l.014-.047h.43a.45.45 0 0 0 .45-.45v-.602a.45.45 0 0 0-.45-.447h-.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Laptop(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.998 18.833H24.38c.843 0 1.526-.683 1.526-1.526V2.034a2.035 2.035 0 0 0-2.035-2.035H3.507a2.035 2.035 0 0 0-2.035 2.035v15.274a1.529 1.529 0 0 0 1.526 1.527zM4.02 2.549h19.336v13.736H4.02zm22.845 17.22H.509a.51.51 0 0 0-.509.509v1.886a2.097 2.097 0 0 0 2.295 1.825l-.009.001h22.801a2.094 2.094 0 0 0 2.285-1.815l.001-.01v-1.886a.51.51 0 0 0-.509-.509zm-9.996 2.62h-6.383a.51.51 0 1 1 0-1.02h.021h-.001h6.38a.51.51 0 1 1 0 1.02h-.021h.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Laravel(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M34.959 11.235h-.055q.438.384-.219.548l-4.766 1.259l4.326 5.862q.219.329.11.468a1.84 1.84 0 0 1-.536.346l-.012.005l-5.587 2.03q-5.314 1.863-5.807 2.028h-.055a2.125 2.125 0 0 1-.946.219h-.041h.002a1.252 1.252 0 0 1-.815-.596l-.003-.006q-.548-.712-3.999-6.957q-5.865 1.533-7.344 1.863a1.404 1.404 0 0 1-1.805-.809l-.003-.01l-1.425-3.127Q.446 2.47.172 1.758Q-.321.662.61.552l.712-.055Q7.025.06 7.677.005a1.497 1.497 0 0 1 .695.114l-.01-.004c.168.115.305.262.407.432l.003.006l8.492 14.19l10.628-2.52l-3.508-4.986q-.384-.548.493-.657l4.931-.819a1.04 1.04 0 0 1 .556.029l-.008-.002c.22.115.404.273.545.464l.003.004l2.136 2.63q1.753 2.185 1.918 2.35zM15.62 15.179q.274-.055.055-.329L7.726 1.099a.222.222 0 0 0-.22-.11h.001L1.7 1.482q-.055 0-.055.11c.001.08.021.155.056.221L1.7 1.81l7.177 14.792q.055.11.11.11h.274l3.889-.935q2.41-.599 2.47-.599zm16.764 3.123l-3.508-4.821a.476.476 0 0 0-.189-.191l-.002-.001a.402.402 0 0 0-.249.028l.003-.001l-10.407 2.74l3.452 5.807q.165.274.274.302a.502.502 0 0 0 .277-.028l-.003.001l10.244-3.508q.219-.055.219-.11a.26.26 0 0 0-.055-.137v.001zm.657-7.341q.438-.11.329-.274q-2.689-3.342-3.013-3.78a.394.394 0 0 0-.441-.219h.003l-3.944.712q-.219.055-.055.274l3.123 4.273q2.026-.493 3.999-.987z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Laughing(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.316C6.099 1.316 1.316 6.099 1.316 12S6.099 22.684 12 22.684S22.684 17.901 22.684 12c-.012-5.896-4.788-10.672-10.683-10.684zm0 22.297C5.586 23.613.387 18.414.387 12S5.586.387 12 .387S23.613 5.586 23.613 12v.015c0 6.405-5.192 11.597-11.597 11.597z"/><path fill="currentColor" d="M12 24C5.386 23.966.034 18.614 0 12.003V12C0 5.373 5.373 0 12 0s12 5.373 12 12c-.034 6.614-5.386 11.966-11.997 12zM12 .774C5.8.774.774 5.8.774 12S5.8 23.226 12 23.226S23.226 18.2 23.226 12C23.222 5.802 18.198.779 12.001.774zm0 22.297C5.886 23.071.929 18.114.929 12S5.886.929 12 .929S23.071 5.886 23.071 12S18.114 23.071 12 23.071m0-21.368C6.313 1.703 1.703 6.313 1.703 12S6.313 22.297 12 22.297S22.297 17.687 22.297 12v-.005c0-5.684-4.608-10.292-10.292-10.292z"/><path fill="currentColor" d="M18.116 14.245a6.116 6.116 0 1 1-12.232 0zm0-4.568h-1.935l.852-.619a.606.606 0 0 0 .077-.853l.001.001a.426.426 0 0 0-.379-.232h-.008a.57.57 0 0 0-.465.155l-2.168 1.703c-.09.112-.168.239-.228.376l-.004.011v.232c.03.306.285.542.596.542h.025h-.001h3.648a.61.61 0 0 0 .611-.611v-.009a.59.59 0 0 0-.155-.388a.562.562 0 0 0-.463-.31h-.002zm-8.051.62v-.155c0-.155-.077-.232-.232-.387L7.665 8.052a1.087 1.087 0 0 0-.46-.154h-.013a.424.424 0 0 0-.378.23l-.001.002a.61.61 0 0 0 .076.851l.001.001l.852.619H5.794a.53.53 0 0 0-.53.53v.013v-.001v.009a.61.61 0 0 0 .611.611h3.647c.31.155.542-.077.542-.465z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LeftAlign(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.334 2.666h26.665l.037.001a1.334 1.334 0 1 0 0-2.668L27.997 0h.002H1.334a1.334 1.334 0 0 0-.002 2.666zm0 5.333h19.555l.037.001a1.334 1.334 0 1 0 0-2.668l-.039.001h.002H1.334a1.334 1.334 0 0 0-.002 2.666zm26.665 2.668H1.334a1.334 1.334 0 0 0-.002 2.666h26.667a1.334 1.334 0 0 0 .002-2.666zm0 10.666H1.334a1.334 1.334 0 0 0-.002 2.666h26.667a1.334 1.334 0 0 0 .002-2.666zM1.334 18.666h19.555A1.334 1.334 0 0 0 20.891 16H1.334a1.334 1.334 0 0 0-.002 2.666z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Less(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.991 10.958v.005c0 .231.091.441.24.595a1 1 0 0 0 .773.262H24v1.05a1 1 0 0 0-.77.263l.001-.001a.875.875 0 0 0-.24.604v.015v-.001q0 .356.056.975t.056 1.069a1.884 1.884 0 0 1-.36 1.318l.004-.005a1.566 1.566 0 0 1-1.188.375l.007.001h-.75v-.936h.229v-.038a.68.68 0 0 0 .526-.169l-.001.001a1.013 1.013 0 0 0 .149-.661v.005l-.075-1.95a1.778 1.778 0 0 1 .211-.97l-.005.009c.152-.219.382-.376.649-.43l.007-.001v-.075a1.077 1.077 0 0 1-.654-.41l-.002-.003a1.752 1.752 0 0 1-.206-.982v.006l.075-1.95a1.135 1.135 0 0 0-.134-.662l.003.006a.667.667 0 0 0-.547-.168h.004h-.225v-.898h.75a1.57 1.57 0 0 1 1.183.377l-.002-.002a1.882 1.882 0 0 1 .355 1.321l.001-.008q0 .45-.056 1.069t-.059.994zm-3.976 1.2q1.614.525 1.614 1.801v.029a1.72 1.72 0 0 1-.637 1.338l-.003.002a2.674 2.674 0 0 1-1.809.543h.008a3.515 3.515 0 0 1-1.186-.233l.024.008a3.05 3.05 0 0 1-1.053-.603l.003.003l.787-1.162c.414.352.948.574 1.533.6h.005q.75 0 .75-.487q0-.225-.338-.45a4.235 4.235 0 0 0-.718-.292l-.031-.008l-.15-.075q-1.538-.6-1.538-1.726l-.001-.05c0-.518.24-.98.616-1.279l.003-.003a2.511 2.511 0 0 1 1.657-.506h-.006a3.404 3.404 0 0 1 2.051.755l-.006-.005l-.785 1.048a2.2 2.2 0 0 0-1.27-.487h-.005q-.64 0-.64.45a.52.52 0 0 0 .26.411l.003.001c.173.108.372.197.583.258l.017.004q.19.074.266.112zM5.549 14.372h.225l.223 1.275a2.37 2.37 0 0 1-.842.151l-.061-.001h.003q-1.614 0-1.614-2.026V8.145h-.526a.614.614 0 0 0-.507.169a1.132 1.132 0 0 0-.131.66v-.005l.038 1.914a1.872 1.872 0 0 1-.205 1.004l.005-.01a.916.916 0 0 1-.652.393h-.004v.075c.274.055.504.212.654.428l.002.003a1.755 1.755 0 0 1 .206.965v-.006q0 .375-.038 1.05v.9a1.23 1.23 0 0 0 .134.682l-.003-.007a.58.58 0 0 0 .51.187h-.003h.226v.9h-.754a1.568 1.568 0 0 1-1.184-.379l.002.002a1.882 1.882 0 0 1-.355-1.322l-.001.008q0-.45.056-1.069t.056-.975v-.022a.815.815 0 0 0-.24-.578a1.043 1.043 0 0 0-.772-.239h.004v-1.05a1 1 0 0 0 .77-.263l-.001.001a.853.853 0 0 0 .24-.595v-.006q0-.338-.056-.975T.898 8.9a1.884 1.884 0 0 1 .36-1.318l-.004.005a1.566 1.566 0 0 1 1.188-.375l-.007-.001h2.814v6.678c-.002.016-.002.034-.002.053c0 .12.036.231.097.324l-.001-.002a.266.266 0 0 0 .206.106zm8.702-2.214q1.614.525 1.614 1.801v.029a1.72 1.72 0 0 1-.637 1.338l-.003.002a2.674 2.674 0 0 1-1.809.543h.008a3.446 3.446 0 0 1-1.211-.234l.024.008a3.547 3.547 0 0 1-1.075-.605l.006.005l.787-1.162c.414.352.948.574 1.533.6h.005q.75 0 .75-.487q0-.225-.338-.45a4.235 4.235 0 0 0-.718-.292l-.031-.008l-.113-.075q-1.576-.6-1.576-1.726l-.001-.042c0-.525.25-.991.637-1.287l.004-.003a2.587 2.587 0 0 1 1.675-.506h-.006a3.247 3.247 0 0 1 2.012.755l-.005-.004l-.786 1.049a2.163 2.163 0 0 0-1.231-.487h-.007q-.675 0-.675.45a.46.46 0 0 0 .26.412l.003.001c.184.111.397.201.623.259l.017.004q.186.074.261.112zM8.437 9.57a2.254 2.254 0 0 1 1.854.803l.003.003a3.11 3.11 0 0 1 .655 2.051v-.006c-.003.241-.03.474-.079.698l.004-.023v.038H7.312c.034.388.221.726.499.958l.002.002a1.392 1.392 0 0 0 1.046.315l-.006.001a2.68 2.68 0 0 0 1.325-.382l-.012.007l.56 1.05a3.592 3.592 0 0 1-2.101.675a2.888 2.888 0 0 1-2.139-.825a3.065 3.065 0 0 1-.846-2.119l.002-.118v.006a3.028 3.028 0 0 1 .827-2.271l-.001.001a2.666 2.666 0 0 1 1.967-.864h.005zm-1.125 2.476H9.45q0-1.162-1.013-1.162a1.04 1.04 0 0 0-.731.3a1.386 1.386 0 0 0-.394.857z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lightbulb(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.45 13.447a7.764 7.764 0 0 0 2.291-5.517v-.07a7.904 7.904 0 0 0-2.291-5.575l.001.001A7.762 7.762 0 0 0 7.939-.001h-.071h.004h-.006A7.904 7.904 0 0 0 2.291 2.29l.001-.001A7.764 7.764 0 0 0 .001 7.806v.067v-.003v.006c0 2.174.875 4.143 2.291 5.575l-.001-.001a5.66 5.66 0 0 1 1.68 4.033v.023c0 .583.472 1.055 1.055 1.055h5.66a.972.972 0 0 0 .722-.32l.001-.001a1.03 1.03 0 0 0 .32-.747v-.005a5.835 5.835 0 0 1 1.717-4.044zm-1.539-1.502a7.695 7.695 0 0 0-2.213 4.468l-.004.039h-.537l1.36-6.546a1.08 1.08 0 0 0-2.108-.473l-.001.007l-.607 2.934l-.681-2.968a1.07 1.07 0 1 0-2.079.507l-.002-.007l1.502 6.509h-.5a7.946 7.946 0 0 0-2.218-4.508l.001.001a5.702 5.702 0 1 1 8.08.036zm.074 8.299v-.018c0-.583-.472-1.055-1.055-1.055h-.019h.001h-6.08a1.072 1.072 0 1 0 0 2.146h6.08a1.077 1.077 0 0 0 1.073-1.073m-1.968 1.61h-4.47a1.072 1.072 0 1 0 0 2.146h4.47a1.074 1.074 0 0 0 .001-2.146h-.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lightning(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.344 23.793a.698.698 0 0 1 0-.99l.858-.858h-.28a.725.725 0 0 1-.141-.014l.005.001c-.022-.004-.041-.012-.062-.018a.59.59 0 0 1-.074-.023l.003.001a.583.583 0 0 1-.071-.037l.003.002c-.017-.01-.036-.017-.052-.028a.705.705 0 0 1-.193-.192l-.002-.003c-.01-.016-.018-.034-.027-.05a.576.576 0 0 1-.035-.066l-.002-.004c-.009-.023-.015-.046-.022-.069s-.014-.04-.018-.062a.733.733 0 0 1 .001-.279l-.001.005c.004-.021.012-.041.018-.062a.707.707 0 0 1 .023-.073l-.002.005a.725.725 0 0 1 .039-.074l-.002.003c.009-.017.017-.034.027-.05a.731.731 0 0 1 .089-.108l1.97-1.97a.7.7 0 1 1 .993.989l-.776.776h.283c.048 0 .096.005.141.014l-.005-.001c.022.004.041.012.062.018a.526.526 0 0 1 .073.023l-.004-.001c.028.013.05.025.072.038l-.003-.002c.021.01.038.019.054.03l-.002-.001c.04.027.075.056.107.087a.846.846 0 0 1 .085.103l.002.003c.011.017.019.036.029.054c.011.018.023.04.033.063l.002.004a.575.575 0 0 1 .021.066l.001.005a.513.513 0 0 1 .017.058l.001.004a.743.743 0 0 1-.001.279l.001-.005a.67.67 0 0 1-.019.066l.001-.004a.57.57 0 0 1-.024.076l.002-.005a.668.668 0 0 1-.037.071l.002-.003c-.01.018-.017.036-.028.053a.673.673 0 0 1-.089.108l-2.053 2.053a.698.698 0 0 1-.99 0zm-2.791-3.711a.702.702 0 0 1 0-.992l2.201-2.2H6.185a.725.725 0 0 1-.141-.014l.005.001c-.021-.004-.04-.012-.062-.018a.537.537 0 0 1-.074-.024l.004.001a.566.566 0 0 1-.07-.037l.003.002c-.018-.01-.036-.017-.054-.029a.779.779 0 0 1-.107-.087a.846.846 0 0 1-.085-.103l-.002-.003l-.03-.054a.595.595 0 0 1-.033-.062l-.002-.004a.482.482 0 0 1-.021-.067l-.001-.004a.458.458 0 0 1-.017-.058l-.001-.004a.655.655 0 0 1 .001-.279l-.001.005c.004-.02.012-.039.017-.062a.638.638 0 0 1 .024-.077l-.002.005a.566.566 0 0 1 .037-.07l-.002.003a.545.545 0 0 1 .03-.055l-.001.002a.726.726 0 0 1 .088-.108l2.087-2.087H5.054a5.054 5.054 0 0 1 0-10.108c.078 0 .156.007.234.011C6.476 1.4 8.698 0 11.248 0a6.818 6.818 0 0 1 6.818 6.817v.013a3.449 3.449 0 0 1 2.622 3.338a3.441 3.441 0 0 1-3.437 3.438h-2.12l-.594.594h.28c.049 0 .097.005.143.015l-.005-.001a.524.524 0 0 1 .059.018l-.004-.001a.49.49 0 0 1 .079.025l-.004-.001a.702.702 0 0 1 .068.036l-.003-.002c.018.01.038.018.055.03a.705.705 0 0 1 .192.19l.002.003a.585.585 0 0 1 .028.051l.001.003c.011.018.023.04.033.063l.002.004c.009.023.015.047.022.071a.513.513 0 0 1 .017.058l.001.004a.655.655 0 0 1 0 .279l.001-.005c-.004.021-.012.039-.018.062a.57.57 0 0 1-.024.076l.002-.005a.668.668 0 0 1-.037.071l.002-.003c-.01.018-.017.036-.028.053a.791.791 0 0 1-.087.102l-2.054 2.053a.7.7 0 1 1-.991-.991l.858-.858h-.28a.725.725 0 0 1-.141-.014l.005.001c-.022-.004-.041-.012-.062-.018a.42.42 0 0 1-.14-.059l.002.001c-.017-.01-.036-.017-.052-.029a.702.702 0 0 1-.193-.191l-.002-.002a1.663 1.663 0 0 1-.026-.047l-.001-.003a.4.4 0 0 1-.061-.136l-.001-.003a.513.513 0 0 1-.017-.058l-.001-.004a.743.743 0 0 1 .001-.279l-.001.005a.47.47 0 0 1 .02-.066l-.001.004a.72.72 0 0 1 .022-.073l-.001.005a.725.725 0 0 1 .039-.074l-.002.003c.009-.016.017-.034.027-.05a.726.726 0 0 1 .088-.108l.8-.8H9.756l-1.882 1.882h1.569c.048 0 .096.005.141.014l-.005-.001c.021.004.041.012.062.018a.48.48 0 0 1 .073.023l-.003-.001a.759.759 0 0 1 .07.037l-.003-.001c.018.01.036.018.054.029a.705.705 0 0 1 .192.19l.002.002a.403.403 0 0 1 .027.051l.001.003c.011.018.023.04.033.063l.002.004a.586.586 0 0 1 .02.066l.001.005a.513.513 0 0 1 .017.058l.001.004a.655.655 0 0 1 0 .279l.001-.005c-.004.02-.012.039-.018.062a.57.57 0 0 1-.024.076l.002-.005a.566.566 0 0 1-.037.07l.002-.003a.711.711 0 0 1-.03.056l.001-.003a.726.726 0 0 1-.088.108l-3.393 3.396a.698.698 0 0 1-.99 0zM1.401 8.552a3.656 3.656 0 0 0 3.653 3.653h12.191a2.037 2.037 0 1 0-1.316-3.589l.003-.002a.701.701 0 0 1-.907-1.069l.001-.001a3.41 3.41 0 0 1 1.613-.758l.021-.003a5.414 5.414 0 0 0-9.89-3.004l-.012.019a5.041 5.041 0 0 1 1.462.816l-.01-.008a.701.701 0 1 1-.877 1.093l.001.001a3.588 3.588 0 0 0-2.268-.8h-.011h.001a3.656 3.656 0 0 0-3.655 3.653z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Lightnings(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.652 23.794a.7.7 0 0 1 0-.992l.859-.857h-.28a.648.648 0 0 1-.141-.015l.005.001c-.022-.004-.041-.012-.062-.018c-.028-.007-.051-.015-.074-.023l.005.002a.615.615 0 0 1-.073-.039l.003.002c-.017-.009-.035-.017-.051-.027a.723.723 0 0 1-.193-.193l-.002-.003c-.011-.016-.018-.034-.027-.05a.576.576 0 0 1-.035-.066l-.002-.004c-.007-.018-.014-.041-.021-.064l-.001-.005c-.006-.021-.014-.041-.018-.062a.655.655 0 0 1 0-.279l-.001.005a.657.657 0 0 1 .02-.066l-.001.004a.42.42 0 0 1 .06-.142l-.001.002a.419.419 0 0 1 .028-.052l-.001.002a.673.673 0 0 1 .089-.108l1.972-1.972a.702.702 0 1 1 .992.991l-.777.777h.28c.049 0 .096.005.142.014l-.005-.001a.47.47 0 0 1 .066.02l-.004-.001a.832.832 0 0 1 .074.022l-.005-.002a.832.832 0 0 1 .074.039l-.003-.002c.02.01.037.019.052.029l-.002-.001a.71.71 0 0 1 .194.192l.002.002l.025.047l.001.003a.594.594 0 0 1 .036.067l.002.004a.612.612 0 0 1 .02.064l.001.005c.006.02.014.04.019.062a.733.733 0 0 1-.001.279l.001-.005a.634.634 0 0 1-.02.065l.001-.004a.72.72 0 0 1-.022.073l.001-.004a.832.832 0 0 1-.039.074l.002-.003l-.028.052l.001-.002a.673.673 0 0 1-.089.108l-2.054 2.055a.7.7 0 0 1-.992 0zm-2.794-3.716a.7.7 0 0 1 0-.992l2.202-2.203h-1.57a.735.735 0 0 1-.142-.014l.005.001c-.022-.004-.041-.012-.062-.018c-.028-.007-.051-.015-.074-.023l.005.002a.725.725 0 0 1-.074-.039l.003.002c-.017-.009-.035-.017-.051-.028a.705.705 0 0 1-.193-.192l-.002-.003c-.01-.016-.018-.034-.028-.051a.576.576 0 0 1-.035-.066l-.002-.004a.555.555 0 0 1-.021-.065l-.001-.004a.513.513 0 0 1-.017-.058l-.001-.004a.655.655 0 0 1 0-.279l-.001.005a.67.67 0 0 1 .019-.066l-.001.004a.526.526 0 0 1 .023-.073l-.001.004a.725.725 0 0 1 .039-.074l-.002.003c.009-.016.017-.034.027-.05a.673.673 0 0 1 .089-.108l2.067-2.067H5.057a5.06 5.06 0 1 1 0-10.12c.078 0 .157.006.235.01C6.481 1.4 8.705 0 11.257 0a6.827 6.827 0 0 1 6.627 5.187l.01.047a4.82 4.82 0 0 1 8.054 3.4v.008a2.548 2.548 0 0 1-.771 4.977h-5.753l-.572.572h.281c.048 0 .096.005.141.014l-.005-.001a.47.47 0 0 1 .066.02l-.004-.001a.832.832 0 0 1 .074.022l-.005-.002a.725.725 0 0 1 .074.039l-.003-.002l.052.028l-.002-.001a.71.71 0 0 1 .194.192l.002.002c.009.014.017.03.025.047l.001.003a.594.594 0 0 1 .036.067l.002.004a.612.612 0 0 1 .02.064l.001.005c.006.021.014.041.019.062a.733.733 0 0 1-.001.279l.001-.005a.634.634 0 0 1-.02.065l.001-.003a.72.72 0 0 1-.022.073l.001-.005a.832.832 0 0 1-.039.074l.002-.003l-.028.052l.001-.002a.673.673 0 0 1-.089.108l-2.056 2.056a.702.702 0 1 1-.992-.991l.859-.857h-.28a.648.648 0 0 1-.141-.015l.005.001a.657.657 0 0 1-.066-.02l.004.001c-.028-.007-.051-.015-.074-.023l.005.002a.725.725 0 0 1-.074-.039l.003.002c-.017-.009-.034-.017-.05-.027a.705.705 0 0 1-.193-.192l-.002-.003a1.663 1.663 0 0 1-.026-.047l-.001-.003a.576.576 0 0 1-.035-.066l-.002-.004c-.007-.018-.014-.041-.021-.064l-.001-.005c-.006-.021-.014-.04-.018-.062a.743.743 0 0 1 .001-.279l-.001.005c.004-.022.012-.041.018-.062a.415.415 0 0 1 .06-.141l-.001.002c.009-.017.016-.034.027-.051a.673.673 0 0 1 .089-.108l.777-.777h-3.396l-1.863 1.862h1.571c.049 0 .096.005.142.014l-.005-.001a.634.634 0 0 1 .065.02l-.003-.001a.832.832 0 0 1 .074.022l-.005-.002a.832.832 0 0 1 .074.039l-.003-.002c.017.009.034.016.05.027a.699.699 0 0 1 .194.193l.002.003l.025.047l.001.003c.012.02.025.043.036.067l.002.004a.612.612 0 0 1 .02.064l.001.005a.503.503 0 0 1 .018.058l.001.004a.655.655 0 0 1 0 .279l.001-.005a.456.456 0 0 1-.02.065l.001-.004a.72.72 0 0 1-.022.073l.001-.005a.725.725 0 0 1-.039.074l.002-.003a.606.606 0 0 1-.029.052l.001-.002a.673.673 0 0 1-.089.108l-3.401 3.4a.7.7 0 0 1-.992 0zm8.435-13.185a3.44 3.44 0 0 1 2.415 3.279c0 .77-.253 1.481-.682 2.054l.006-.009h5.144a1.145 1.145 0 1 0-.741-2.019l.002-.001a.702.702 0 0 1-.906-1.072l.001-.001c.282-.24.622-.42.995-.518l.018-.004a3.425 3.425 0 0 0-3.416-3.218a3.417 3.417 0 0 0-2.828 1.495l-.008.012zM1.404 8.557a3.66 3.66 0 0 0 3.657 3.657h12.205a2.038 2.038 0 1 0-1.319-3.592l.003-.002a.702.702 0 0 1-.908-1.071l.001-.001a3.419 3.419 0 0 1 1.614-.759l.021-.003A5.42 5.42 0 0 0 6.777 3.78l-.012.019a5.076 5.076 0 0 1 1.464.816l-.01-.008a.702.702 0 1 1-.879 1.095l.001.001a3.597 3.597 0 0 0-2.271-.801h-.01h.001a3.662 3.662 0 0 0-3.66 3.658z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Like(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.547 24a.33.33 0 0 1-.327-.295v-.001L0 11.336l-.001-.031c0-.182.147-.329.329-.329h8.233v13.025zm3.36-3.265v.001a1.301 1.301 0 1 0 1.301-1.301h-.002c-.719 0-1.301.582-1.302 1.301zm.48-8c0 .181.147.328.329.328h.764a.328.328 0 0 0 0-.656h-.77a.329.329 0 0 0-.329.329v.003zm-3.61 0c0 .181.147.328.328.328H4.61a.328.328 0 0 0 0-.656H2.099a.328.328 0 0 0-.328.328v.003zm9.8 9.388a1.886 1.886 0 0 1-.494-.253l.006.004l-.1-.063a20.97 20.97 0 0 1-.69-.452l-.101-.071c-.132-.095-.194-.137-.262-.133c-.24.005-.48.007-.72.01V12.01l-.013-.77c.451-.287.867-.56.992-.664c.17-.342.338-.628.524-.902l-.017.027c.085-.133.17-.265.248-.396l1.724-2.894c.08-.136.167-.272.254-.408c.145-.216.299-.474.44-.74l.024-.049a.879.879 0 0 0 .079-.521l.001.005l-.006-3.12c.063-.454.319-.838.68-1.072l.006-.004a1.97 1.97 0 0 1 1.162-.5h.007l.063-.001c.341 0 .663.081.949.224l-.012-.006c.191.092.354.19.507.3l-.01-.007l.106.07c.241.127.421.341.501.6l.002.007c.187.696.358 1.329.517 1.964l.066.259c.101.34.188.751.244 1.172l.005.042c-.171 1.574-.5 3.01-.976 4.378l.042-.138l7.693-.011h.028a1.643 1.643 0 0 1 1.63 1.848l.001-.008a1.982 1.982 0 0 1-.894 1.781l-.008.005a2.171 2.171 0 0 1 .372 1.4v-.007a1.848 1.848 0 0 1-1.137 1.873l-.012.004a2.375 2.375 0 0 1 .299 1.31v-.006a1.905 1.905 0 0 1-.886 1.733l-.008.005a2.791 2.791 0 0 1 .224 1.459l.001-.013v.188a1.834 1.834 0 0 1-2.007 1.719l.007.001h-4.429l-.015.001l-.016-.001h.001h-5.91c-.051 0-.104 0-.16.008c-.068.004-.14.01-.214.01h-.015a1.13 1.13 0 0 1-.299-.04l.008.002zm-6.015-1.387a.64.64 0 1 1 1.281.001a.64.64 0 0 1-1.281-.001"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Line(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12.171 23.999l-.024.001a.376.376 0 0 1-.209-.064l.001.001c-.251-.197-.159-.676-.053-1.231a3.74 3.74 0 0 0 .098-1.4l.002.018a1.12 1.12 0 0 0-.984-.895h-.005a14.832 14.832 0 0 1-4.706-1.321l.091.038a13.208 13.208 0 0 1-1.842-1.042l.047.03a12.066 12.066 0 0 1-1.559-1.244l.006.005c-.46-.434-.877-.903-1.251-1.406l-.02-.029a9.324 9.324 0 0 1-.929-1.544l-.025-.057A8.258 8.258 0 0 1 0 10.266c0-.734.094-1.446.272-2.125l-.013.058a8.772 8.772 0 0 1 .77-1.974l-.023.047a9.855 9.855 0 0 1 1.198-1.754l-.011.013a11.3 11.3 0 0 1 1.543-1.506l.022-.017a12.621 12.621 0 0 1 1.831-1.221l.066-.033A13.448 13.448 0 0 1 7.734.838l.099-.03a14.25 14.25 0 0 1 2.318-.586L10.24.21a16.184 16.184 0 0 1 5.262.013L15.408.21c.922.152 1.732.36 2.514.63l-.107-.032c.847.288 1.566.608 2.251.983l-.073-.037c.722.393 1.343.808 1.922 1.273l-.025-.019c.578.464 1.09.963 1.552 1.507l.013.016c.435.51.827 1.081 1.156 1.691l.026.052c.301.551.558 1.191.737 1.863l.014.061a8.17 8.17 0 0 1 .26 2.069a8.338 8.338 0 0 1-.521 2.911l.019-.058a9.388 9.388 0 0 1-1.427 2.565l.015-.02a2.708 2.708 0 0 1-.239.298l.002-.002l-.008.009c-.428.508-.884.963-1.377 1.377l-.017.014a50.146 50.146 0 0 1-5.52 4.351l-.174.112a27.222 27.222 0 0 1-2.425 1.493l-.148.074a8.482 8.482 0 0 1-.916.428l-.06.021a2.031 2.031 0 0 1-.673.159h-.006zm5.983-16.742a.67.67 0 0 0-.668.668v4.958a.67.67 0 0 0 .668.668h2.586a.67.67 0 0 0 .668-.668v-.056a.67.67 0 0 0-.668-.668h-1.863v-1.057h1.862a.67.67 0 0 0 .668-.668v-.056a.67.67 0 0 0-.668-.668h-1.862V8.651h1.862a.67.67 0 0 0 .668-.668v-.057a.67.67 0 0 0-.668-.668zm-5.971-.012a.67.67 0 0 0-.668.668v4.958a.67.67 0 0 0 .668.668h.056a.67.67 0 0 0 .668-.668V9.905l2.434 3.325a.597.597 0 0 0 .05.068l-.001-.001a.557.557 0 0 0 .253.191l.004.001c.075.032.162.05.253.05h.059a.655.655 0 0 0 .315-.08l-.004.002a.47.47 0 0 0 .163-.121v-.001a.662.662 0 0 0 .193-.468V7.913a.668.668 0 0 0-.668-.668h-.059a.67.67 0 0 0-.668.668v2.911l-2.414-3.241a.673.673 0 0 0-.58-.342zm-2.334 0a.67.67 0 0 0-.668.668v4.958a.67.67 0 0 0 .668.668h.056a.67.67 0 0 0 .668-.668V7.913a.67.67 0 0 0-.668-.668zm-4.462 0a.67.67 0 0 0-.668.668v4.958a.67.67 0 0 0 .668.668h2.584a.67.67 0 0 0 .668-.668v-.05a.668.668 0 0 0-.668-.668H6.11V7.914a.67.67 0 0 0-.668-.668z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LineChart(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M32 22v2H0V0h2v22zM30 2.5v6.802a.498.498 0 0 1-.859.342l-1.89-1.89l-9.89 9.887a.49.49 0 0 1-.72 0L13.001 14l-6.5 6.5l-3-3l9.14-9.141a.49.49 0 0 1 .72 0L17.001 12l7.25-7.25l-1.89-1.89a.497.497 0 0 1 .342-.86h6.822a.48.48 0 0 1 .48.48v.021V2.5z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Link(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12.922 16.587l-3.671 3.671a3.896 3.896 0 0 1-5.504-5.509l-.002.002l3.671-3.671a1.3 1.3 0 0 0-1.837-1.835l.001-.001l-3.671 3.671a6.494 6.494 0 0 0 9.187 9.175l-.003.002l3.671-3.671a1.3 1.3 0 0 0-1.837-1.835l.001-.001zM24.007 6.489A6.494 6.494 0 0 0 12.921 1.9L9.25 5.571a1.3 1.3 0 1 0 1.835 1.837l.001-.001l3.671-3.671a3.896 3.896 0 0 1 5.504 5.509l.002-.002l-3.671 3.671a1.3 1.3 0 1 0 1.835 1.837l.001-.001l3.671-3.671a6.432 6.432 0 0 0 1.908-4.58V6.49z"/><path fill="currentColor" d="M7.412 16.592c.235.235.559.38.918.38s.683-.145.918-.38L16.59 9.25a1.3 1.3 0 0 0-1.837-1.835l.001-.001l-7.342 7.342c-.235.235-.38.559-.38.918s.145.683.38.918"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Linkedin(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.706 7.798V24H.311V7.798zm.343-5.002l.002.098c0 .749-.318 1.423-.826 1.895l-.002.001A3.067 3.067 0 0 1 3 5.59h.007h-.033a2.945 2.945 0 0 1-2.162-.801l.001.001a2.67 2.67 0 0 1-.815-1.998v.004l-.001-.069c0-.762.324-1.448.841-1.929L.84.797A3.066 3.066 0 0 1 3.045.001h-.006a2.983 2.983 0 0 1 2.177.795L5.214.794a2.72 2.72 0 0 1 .835 1.964V2.8v-.002zm19.062 11.92V24h-5.378v-8.665a4.702 4.702 0 0 0-.675-2.71l.012.021a2.325 2.325 0 0 0-2.076-.972h.008a2.634 2.634 0 0 0-1.73.568l.006-.004a3.494 3.494 0 0 0-1.032 1.375l-.008.023a3.862 3.862 0 0 0-.179 1.331v-.007v9.042H8.681q.033-6.523.033-10.578t-.016-4.839L8.682 7.8h5.378v2.354h-.033c.214-.345.435-.644.678-.924l-.008.009c.281-.309.583-.588.908-.838l.016-.012a4.212 4.212 0 0 1 1.392-.704l.03-.007a6.347 6.347 0 0 1 1.797-.254h.079h-.004a5.792 5.792 0 0 1 4.493 1.851l.003.004q1.702 1.856 1.702 5.436z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Linux(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.782 5.505a.242.242 0 0 0-.207.139l-.001.002q-.06.127-.114.127q-.067.014-.067-.067q0-.16.254-.201zm1.164.187q-.054.014-.154-.087a.223.223 0 0 0-.236-.06h.002q.32-.147.429.026a.08.08 0 0 1-.039.12h-.001zm-4.7 5.719a.066.066 0 0 0-.08.039a.992.992 0 0 0-.058.16l-.002.008a.789.789 0 0 1-.076.185l.002-.004a1.267 1.267 0 0 1-.135.175l.001-.001q-.134.147-.014.16q.054.014.167-.094a.688.688 0 0 0 .166-.236l.002-.004q.014-.04.026-.094a.475.475 0 0 1 .027-.083l-.001.003a.209.209 0 0 0 .02-.058v-.002a.22.22 0 0 0 .006-.053v-.04l-.014-.034zm11.45 4.808q0-.24-.737-.56q.054-.201.101-.368c.029-.099.053-.217.066-.338l.001-.01q.02-.181.04-.288a1.216 1.216 0 0 0 .006-.307v.005q-.014-.194-.014-.261a1.942 1.942 0 0 0-.05-.308l.003.013q-.047-.228-.054-.274t-.067-.334t-.074-.355a3.958 3.958 0 0 0-.638-1.391l.008.012a3.197 3.197 0 0 0-.952-.997l-.012-.007c.309.314.565.682.753 1.087l.01.024a5.058 5.058 0 0 1 .714 3.755l.005-.032a.686.686 0 0 1-.669.56q-.415.054-.515-.248a4.125 4.125 0 0 1-.107-1.128v-.024c0-.497-.056-.98-.163-1.445l.008.043a8.509 8.509 0 0 0-.28-.986l.019.062a3.674 3.674 0 0 0-.271-.629l.01.019q-.121-.207-.207-.328a1.301 1.301 0 0 0-.174-.2l-.001-.001l-.095-.094a8.517 8.517 0 0 0-.436-1.437l.021.058a3.164 3.164 0 0 0-.402-.759l.006.009a3.664 3.664 0 0 1-.305-.426l-.009-.016a1.387 1.387 0 0 1-.118-1.262l-.003.009a1.16 1.16 0 0 0 .059-.671l.001.008q-.074-.228-.596-.334a2.579 2.579 0 0 1-.609-.247l.013.007a2.762 2.762 0 0 0-.455-.208l-.02-.006q-.107-.014-.147-.348a1.317 1.317 0 0 1 .11-.691l-.003.008a.526.526 0 0 1 .479-.362h.001a.623.623 0 0 1 .682.396l.001.004a.989.989 0 0 1 .052.784l.002-.007q-.147.254-.026.355t.4.006q.174-.054.174-.48v-.498a3 3 0 0 0-.188-.69l.007.02a1.055 1.055 0 0 0-.281-.408l-.001-.001a1.214 1.214 0 0 0-.307-.198l-.008-.003a1.732 1.732 0 0 0-.35-.099l-.011-.002q-1.433.107-1.192 1.794a.6.6 0 0 1-.015.205l.001-.004a.605.605 0 0 0-.39-.141h-.006a1.904 1.904 0 0 0-.452.008l.01-.001q-.167.026-.207-.067a2.555 2.555 0 0 0-.221-1.222l.007.016a.714.714 0 0 0-.6-.455h-.02a.578.578 0 0 0-.537.364l-.001.004a1.98 1.98 0 0 0-.22.794v.006l-.001.068c0 .151.018.299.051.44l-.003-.013c.036.189.096.357.179.512l-.005-.01q.114.207.207.181a.365.365 0 0 0 .213-.185l.001-.002q.054-.121-.094-.107q-.094 0-.207-.194a.962.962 0 0 1-.127-.447v-.002a.798.798 0 0 1 .123-.499l-.002.003a.48.48 0 0 1 .458-.186H8.04c.17.011.31.124.361.279l.001.003c.08.152.127.332.127.522v.01c0 .1-.007.199-.021.295l.001-.011a1.498 1.498 0 0 0-.412.384l-.003.004c-.098.13-.221.235-.362.311l-.006.003q-.261.154-.274.167a.757.757 0 0 0-.206.356l-.001.005a.197.197 0 0 0 .1.239l.001.001a1.5 1.5 0 0 1 .335.26c.076.077.147.16.21.248l.004.006c.063.08.147.14.244.173l.004.001c.141.053.304.085.474.087h.001l.101.001c.452 0 .887-.074 1.293-.21l-.029.008q.026-.014.308-.094t.462-.141c.154-.054.284-.113.409-.181l-.013.007a.67.67 0 0 0 .28-.232l.001-.002q.121-.187.268-.107a.187.187 0 0 1 .087.112v.001a.194.194 0 0 1-.04.16a.363.363 0 0 1-.219.127h-.002a7.825 7.825 0 0 0-.807.308l.05-.02q-.489.207-.61.261a3.978 3.978 0 0 1-.911.303l-.027.005a3.57 3.57 0 0 1-1.084-.025l.022.003q-.134-.026-.121.026c.063.099.138.183.226.253l.002.002a1.224 1.224 0 0 0 .902.294h-.004c.176-.013.339-.046.494-.098l-.014.004c.191-.06.35-.124.503-.197l-.023.01q.228-.107.449-.234t.4-.228a2.95 2.95 0 0 1 .308-.152l.02-.008a.36.36 0 0 1 .237-.033h-.002a.18.18 0 0 1 .114.146v.001a.176.176 0 0 1-.014.061V7.73a.186.186 0 0 1-.053.067a.911.911 0 0 1-.077.058l-.003.002q-.04.026-.114.067t-.121.06t-.134.067a.8.8 0 0 1-.121.058l-.006.002c-.344.185-.64.38-.916.599l.012-.01c-.26.207-.551.399-.86.561l-.03.015a.677.677 0 0 1-.659.012l.004.002a3.353 3.353 0 0 1-.835-.962l-.009-.016q-.294-.415-.334-.294a.45.45 0 0 0-.014.114v.021v-.001a1.792 1.792 0 0 1-.205.767l.005-.01q-.201.422-.395.743a2.32 2.32 0 0 0-.28.763l-.002.014a1.133 1.133 0 0 0 .157.849l-.003-.005q-.308.08-.837 1.206a6.719 6.719 0 0 0-.633 1.845l-.007.043q-.026.24-.02.924a2.017 2.017 0 0 1-.077.805l.004-.014q-.107.32-.388.04a1.841 1.841 0 0 1-.48-1.244v-.016v.001a2.667 2.667 0 0 1 .057-.767l-.003.018q.054-.254-.014-.24a.137.137 0 0 0-.053.066v.001a2.422 2.422 0 0 0 .139 2.234l-.006-.011c.083.152.196.277.331.373l.003.002q.268.214.32.268c.426.424.881.821 1.359 1.187l.033.025q1.125.904 1.246 1.025a.733.733 0 0 1 .234.507v.002a.839.839 0 0 1-.798.88h-.002c.133.231.261.426.4.612l-.012-.016c.157.206.285.445.37.703l.005.017a3.53 3.53 0 0 1 .093.95v-.006q.616-.32.094-1.232a1.17 1.17 0 0 0-.142-.216l.002.002q-.087-.107-.127-.16t-.026-.08a.38.38 0 0 1 .172-.126l.003-.001a.256.256 0 0 1 .269.034q.616.696 2.223.48h.006c.96 0 1.814-.454 2.359-1.158l.005-.007q.308-.509.455-.4q.16.08.134.696c-.062.46-.17.876-.321 1.27l.013-.038a.904.904 0 0 0-.08.507v-.005q.04.194.32.207q.04-.254.194-1.031c.077-.337.141-.749.178-1.168l.003-.038l.001-.094c0-.315-.032-.622-.093-.919l.005.029a6.917 6.917 0 0 1-.101-1.201l.001-.103v.005l-.001-.044c0-.341.116-.655.311-.904l-.002.003a.845.845 0 0 1 .687-.24h-.004v-.012c0-.312.188-.579.458-.695l.005-.002a1.646 1.646 0 0 1 .981-.14l-.01-.001c.309.026.589.135.822.304l-.005-.003zM8.286 5.143a.715.715 0 0 0-.035-.405l.002.005q-.074-.174-.154-.201q-.121-.026-.121.094c.01.036.034.064.066.08h.001q.134 0 .094.201q-.04.268.107.268l.006.001a.035.035 0 0 0 .035-.035l-.001-.007zm5.611 2.64a.273.273 0 0 0-.087-.154a.388.388 0 0 0-.172-.067h-.002a.51.51 0 0 1-.196-.075l.002.001a.55.55 0 0 1-.127-.107l-.001-.001l-.094-.107l-.074-.087a.295.295 0 0 0-.053-.053h-.001q-.014-.006-.054.02q-.187.214.094.582c.122.193.302.341.515.419l.007.002a.194.194 0 0 0 .194-.106l.001-.001a.37.37 0 0 0 .047-.271v.002zm-2.383-2.852v-.004a.527.527 0 0 0-.213-.424l-.001-.001q-.08-.054-.121-.04l-.019-.001a.152.152 0 0 0-.089.028c-.008.006-.013.016-.013.027s.005.021.013.027a.165.165 0 0 0 .066.04h.001q.187.054.24.415q0 .04.107-.026q.028-.029.028-.04zm.72-3.12a.114.114 0 0 0-.034-.067a.496.496 0 0 0-.114-.091l-.002-.001l-.127-.08a.552.552 0 0 0-.316-.2l-.004-.001h-.001a.167.167 0 0 0-.153.1v.001a.294.294 0 0 0-.013.177v-.002a.313.313 0 0 1-.007.169l.001-.002a.443.443 0 0 1-.08.141v-.001a.627.627 0 0 0-.078.117l-.002.004q-.014.034.04.114c.014.013.033.021.054.021s.039-.008.054-.021q.054-.04.147-.121a.664.664 0 0 1 .196-.119l.005-.002a.262.262 0 0 1 .122-.013h-.001h.006c.069 0 .136-.01.199-.028l-.005.001a.171.171 0 0 0 .118-.094v-.001zM19.8 19.77c.159.092.296.2.414.327l.001.001a.577.577 0 0 1 .16.317v.003a.753.753 0 0 1-.035.307l.002-.005a.704.704 0 0 1-.206.294l-.001.001a4.212 4.212 0 0 1-.304.254l-.01.007a3.362 3.362 0 0 1-.382.239l-.018.009q-.248.134-.422.221t-.429.207t-.362.174c-.436.23-.81.479-1.158.76l.013-.01a13.19 13.19 0 0 0-1.018.863l.007-.006a1.524 1.524 0 0 1-.858.262l-.055-.001h.003a2.377 2.377 0 0 1-1.207-.2l.015.006a1.166 1.166 0 0 1-.394-.312l-.002-.002a1.707 1.707 0 0 1-.216-.332l-.005-.01a.599.599 0 0 0-.29-.259l-.004-.001a1.527 1.527 0 0 0-.614-.127h-.016h.001q-.59-.014-1.741-.014q-.254 0-.763.02t-.777.034a3.076 3.076 0 0 0-1.086.208l.021-.007a2.433 2.433 0 0 0-.725.404l.005-.004c-.168.147-.36.274-.567.375l-.015.007c-.193.098-.42.156-.661.156l-.062-.001h.003a6.08 6.08 0 0 1-1.518-.429l.04.015a12.693 12.693 0 0 0-1.864-.559l-.091-.017q-.254-.054-.683-.127t-.67-.121t-.529-.127a1.456 1.456 0 0 1-.454-.198l.006.003a.651.651 0 0 1-.226-.257l-.002-.004a1.215 1.215 0 0 1 .096-.897l-.003.007c.094-.207.176-.45.235-.703l.005-.027l.001-.063c0-.168-.02-.331-.057-.487l.003.014q-.067-.32-.134-.569a1.662 1.662 0 0 1-.06-.448l.001-.043v-.019a.48.48 0 0 1 .141-.34a1.335 1.335 0 0 1 .767-.187h-.004h.039c.275 0 .537-.059.773-.165l-.012.005c.221-.122.407-.278.557-.465l.003-.004a1.178 1.178 0 0 0 .16-.687v.004a1.129 1.129 0 0 1-.416 1.413l-.004.003a1.715 1.715 0 0 1-1.121.2l.01.001q-.455-.04-.576.134q-.174.201.067.763c.036.096.072.175.113.251l-.005-.011q.08.16.114.24c.025.065.046.142.059.221l.001.007a1.092 1.092 0 0 1 .013.299v-.004a1.762 1.762 0 0 1-.232.664l.004-.008c-.117.18-.187.401-.187.638v.002q.04.228.495.348q.268.08 1.132.248t1.333.274q.32.08.991.294t1.105.308a2.156 2.156 0 0 0 .753.053l-.009.001c.336-.027.635-.165.865-.376l-.001.001a.985.985 0 0 0 .308-.636v-.004l.001-.085c0-.249-.037-.489-.107-.716l.005.017a3.809 3.809 0 0 0-.264-.718l.01.022q-.134-.261-.268-.489a28.417 28.417 0 0 0-2.304-3.287l.041.052q-.91-.991-1.514-.536q-.147.121-.201-.201a2.129 2.129 0 0 1-.026-.516v.007c.005-.251.054-.49.138-.71l-.005.014c.097-.243.205-.45.329-.645l-.009.015c.106-.16.204-.344.286-.537l.009-.023q.107-.282.355-.964t.395-1.045c.136-.326.27-.595.419-.853l-.019.037a3.28 3.28 0 0 1 .524-.721l-.001.001a10.306 10.306 0 0 0 1.635-2.546l.026-.065q-.16-1.5-.214-4.151a4.906 4.906 0 0 1 .332-2.062l-.012.033A2.987 2.987 0 0 1 8.34.294l.018-.008a2.897 2.897 0 0 1 1.4-.281h-.006l.085-.001c.475 0 .935.066 1.37.19l-.035-.009c.455.117.853.309 1.202.563l-.01-.007c.539.42.958.969 1.216 1.601l.009.026a4.136 4.136 0 0 1 .394 1.984v-.009A8.65 8.65 0 0 0 14.4 7.27l-.017-.061a7.443 7.443 0 0 0 1.781 2.92l-.001-.001a7.513 7.513 0 0 1 1.315 2.133l.018.051c.337.735.614 1.593.788 2.487l.012.073a4.599 4.599 0 0 1 .066 1.145l.001-.014a2.333 2.333 0 0 1-.166.759l.006-.016q-.121.268-.268.294a.641.641 0 0 0-.313.252l-.002.003q-.181.228-.362.475a1.662 1.662 0 0 1-.533.444l-.009.004a1.538 1.538 0 0 1-.82.187h.004a1.91 1.91 0 0 1-.435-.07l.013.003a.665.665 0 0 1-.301-.181a2.636 2.636 0 0 1-.177-.202l-.004-.005a1.801 1.801 0 0 1-.15-.263l-.005-.011q-.094-.194-.121-.261q-.294-.495-.549-.4t-.375.656a2.65 2.65 0 0 0 .099 1.318l-.006-.019c.09.402.141.864.141 1.338c0 .451-.046.89-.135 1.315l.007-.042a1.683 1.683 0 0 0 .244 1.351l-.004-.006a1.1 1.1 0 0 0 .981.441h-.004c.44-.032.833-.207 1.141-.477l-.002.002a9.658 9.658 0 0 1 1.161-.867l.039-.023a8.425 8.425 0 0 1 1.324-.551l.062-.017a3.704 3.704 0 0 0 1.044-.497l-.013.008q.32-.248.248-.462a.775.775 0 0 0-.331-.38l-.004-.002a3.162 3.162 0 0 0-.667-.308l-.023-.007a1.111 1.111 0 0 1-.661-.633l-.003-.007a2.12 2.12 0 0 1-.202-.908l.001-.066v.003a.876.876 0 0 1 .208-.642l-.001.001c.006.272.045.532.112.78l-.005-.023c.054.209.12.389.202.561l-.008-.018c.076.146.167.272.275.382c.085.093.177.176.276.251l.005.004q.107.074.288.174t.218.125z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListOne(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.064 2.677h17.753l.037.001a1.332 1.332 0 1 0 0-2.664l-.039.001h.002H9.064a1.331 1.331 0 0 0-.002 2.662zm17.752 2.666H9.064a1.331 1.331 0 0 0-.002 2.662h17.754l.037.001a1.332 1.332 0 1 0 0-2.664l-.039.001zm0 5.326H9.064l-.037-.001a1.332 1.332 0 1 0 0 2.664l.039-.001h-.002h17.752l.037.001a1.332 1.332 0 1 0 0-2.664l-.039.001zm0 10.652H9.064a1.331 1.331 0 0 0-.002 2.662h17.754a1.331 1.331 0 0 0 .002-2.662zm0-5.327H9.064l-.037-.001a1.332 1.332 0 1 0 0 2.664l.039-.001h-.002h17.752l.037.001a1.332 1.332 0 1 0 0-2.664l-.039.001zM1.582 2.279v3.334a.648.648 0 0 0 .647.636l.041-.001h-.002l.032.001a.654.654 0 0 0 .654-.634V.627a.636.636 0 0 0-.636-.628h-.022h.001a.652.652 0 0 0-.537.286l-.001.002l-.006.009l-.579.82a.797.797 0 0 0-.209.504v.012c0 .347.273.63.616.646h.001zM.677 15.124h2.614a.67.67 0 0 0 .002-1.335H1.392v-.01c0-.2.457-.523.824-.781c.732-.515 1.639-1.157 1.639-2.27v-.016a1.837 1.837 0 0 0-1.906-1.836h.003A1.708 1.708 0 0 0 .13 10.535v.002a.661.661 0 0 0 .656.745h.009l.036.001a.666.666 0 0 0 .665-.642v-.001c0-.197.05-.418.443-.418l.038-.001c.281 0 .508.227.508.508v.017v-.001c0 .418-.549.818-1.079 1.2c-.655.473-1.399 1.008-1.399 1.836v.711a.664.664 0 0 0 .663.633h.008zm3.088 4.352a1.593 1.593 0 0 0-1.785-1.723l.008-.001a1.673 1.673 0 0 0-1.869 1.611v.002a.645.645 0 0 0 .64.722l.053-.002H.81a.641.641 0 0 0 .678-.64l-.001-.038v.002c0-.156.055-.335.487-.335c.314 0 .437.045.437.494s-.094.487-.465.487a.652.652 0 0 0 .003 1.302h.004c.49 0 .589.201.589.523v.138c0 .544-.209.647-.603.647c-.539 0-.582-.278-.582-.364v-.016a.634.634 0 0 0-.694-.631h.003a.627.627 0 0 0-.668.701v-.003a1.77 1.77 0 0 0 1.956 1.643l-.008.001A1.812 1.812 0 0 0 3.91 22.01l.001.008v-.138l.001-.067c0-.452-.179-.863-.469-1.165l.001.001c.204-.313.325-.697.325-1.109l-.001-.07v.003z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ListTwo(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.667 2.666H28A1.334 1.334 0 0 0 28.002 0H6.667a1.334 1.334 0 0 0-.002 2.666zm21.332 2.667H6.666l-.037-.001a1.334 1.334 0 1 0 0 2.668l.039-.001h-.002h21.333a1.334 1.334 0 0 0 .002-2.666zm0 5.334H6.666a1.334 1.334 0 0 0-.002 2.666h21.335a1.334 1.334 0 0 0 .002-2.666zm0 10.666H6.666a1.334 1.334 0 0 0-.002 2.666h21.335a1.334 1.334 0 0 0 .002-2.666zm0-5.333H6.666a1.334 1.334 0 0 0-.002 2.666h21.335A1.334 1.334 0 0 0 28.001 16zM1.334 0h-.002a1.334 1.334 0 1 0 .002 0m0 5.333h-.002a1.334 1.334 0 1 0 .002 0m0 5.334h-.002a1.334 1.334 0 1 0 .002 0m0 10.666h-.002a1.334 1.334 0 1 0 .002 0m0-5.333h-.002a1.334 1.334 0 1 0 .002 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Livestream(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M28.246 4.261h.055l.052.008l.054.015l.052.02l.051.027l.053.03l.047.04l.047.039l.05.061l.062.061l.04.061l.059.061l.101.16l.039.081l.039.101l.04.1l.042.1l.038.1l.078.242l.04.121l.042.12l.039.135l.038.135l.042.141l.04.135l.078.269l.04.162l.04.159l.039.182l.02.162l.039.18l.04.18l.023.182l.038.201l.02.199l.02.201l.042.201l.02.201l.022.199l.02.222l.02.221l.02.219l.02.222l.02.221l.022.241l.02.462l.022.241l.02.241l.04.485l.061.779l.022.262l.02.26v.522l.022.281l.02.281v.538l.02.28v.822l.02.564l.02.841v1.742l.02.502l.04.362l.062.241l.058.16l.063.135l.079.121l.061.059l.061.061l.059.061l.083.039l.12.062l.12.039l.135.04l.191.039l.26.015l.362-.022h.129v3.032l-.322.039l-.762.035l-.821-.039l-.603-.09l-.463-.12l-.342-.12l-.319-.16l-.318-.202l-.28-.248l-.18-.18l-.183-.201l-.229-.362l-.202-.421l-.16-.46l-.11-.538l-.073-.614l-.018-.682v-2.608l-.02-.563v-1.164l-.02-.538v-.393l-.022-.269v-.538l-.02-.261v-.25l-.011-.26l-.02-.261l-.019-.241v-.242l-.02-.241l-.02-.246v-.24l-.02-.242l-.023-.241l-.016-.241l-.019-.229l-.023-.229l-.02-.221l-.02-.227l-.02-.215l-.026-.219l-.02-.221l-.023-.202l-.042-.412l-.02-.19l-.02-.187l-.054-.369l-.022-.18l-.02-.182l-.04-.168l-.04-.178l-.039-.159l-.04-.162l-.042-.16l-.046-.135l-.039-.135l-.034-.135l-.061-.121l-.039-.135l-.039-.12l-.03-.121l-.039-.12l-.042-.121l-.04-.1l-.039-.102l-.039-.1l-.04-.079l-.038-.105l-.039-.081l-.059-.081l-.04-.061l-.048-.061l-.042-.059l-.055-.04l-.038-.039l-.062-.04l-.038-.02l-.062-.022l-.066-.016l-.059-.022l-.04-.02h3.394l-.069.042zM.011 0l1.573.024L3.05.059l1.306.02l1.145.04l.982.039l.864.039l.763.04l.681.039l.623.062l.563.059l.521.062l.48.059l.864.121l.381.082l.363.079l.657.164l.3.081l.281.081l.261.101l.261.079l.241.1l.219.102l.223.1l.199.1l.221.1l.199.121l.363.242l.16.118l.48.363l.135.121l.135.135l.135.121l.135.135l.242.269l.135.141l.242.269l.1.135l.121.141l.1.16l.1.135l.1.16l.101.162l.1.16l.1.162l.081.16l.082.16l.1.16l.242.48l.081.162l.061.159l.065.162l.081.182l.067.159l.081.182l.081.16l.059.16l.082.182l.073.182l.059.16l.069.18l.061.162l.061.179l.121.363l.061.18l.059.182l.04.18l.061.182l.053.182l.039.179l.052.182l.039.182l.055.18l.039.18l.061.182l.039.16l.061.182v.019l.039.162l.042.171l.039.18l.039.182l.078.205l.04.182l.039.179l.061.162l.039.18l.04.16l.04.182l.078.32l.04.162l.04.16l.039.162l.081.32l.078.32l.081.32l.019.162l.04.135l.062.3l.078.269l.04.135l.03.135l.039.135l.022.135l.039.135l.04.135l.02.135l.039.135l.04.121l.04.135l.02.12l.04.121l.02.121l.02.12l.02.121l.02.101l.022.121l.02.114l.022.101l.019.1l.022.105l.032.1l.02.1l.02.101l.02.092l.022.081l.031.1l.031.082l.022.087l.02.081l.022.081l.061.242l.02.061l.04.069l.04.121l.022.059l.039.062l.02.059l.051.101l.02.051l.04.082l.02.039l.02.04l.02.039l.022.04l.061.061l.022.02l.02.022l.039.022h.022l.02.02h-9.421l-.069-.02l-.077-.022l-.077-.019l-.081-.022l-.081-.02h-.022l-.07-.032l-.075-.02l-.077-.042l-.082-.04l-.059-.039l-.081-.04l-.061.027l-.082-.04l-.118-.121l-.082-.061l-.061-.061l-.081-.059l-.081-.081l-.081-.075l-.059-.081l-.082-.073l-.081-.061l-.081-.081l-.082-.081l-.059-.081h-.083l-.079-.081l-.061-.092l-.073-.087l-.16-.199l-.079-.101l-.061-.1l-.083-.136l-.079-.1l-.082-.121l-.081-.101l-.081-.118l-.162-.242l-.082-.121l-.079-.121l-.082-.121l-.081-.12l-.081-.135l-.079-.121l-.062-.135l-.242-.404l-.082-.135l-.079-.135l-.162-.269l-.082-.135l-.071-.194l-.082-.135l-.079-.135l-.082-.16l-.081-.162l-.081-.16l-.082-.162l-.086-.164l-.082-.16l-.079-.16l-.081-.16l-.182-.038l-.066-.17l-.079-.166l-.082-.166l-.079-.168l-.086-.168l-.081-.171l-.081-.16l-.086-.16l-.081-.16l-.085-.16l-.081-.16l-.081-.202l-.081-.18l-.085-.16l-.085-.162l-.081-.18l-.085-.16l-.082-.182l-.081-.16l-.081-.19l-.162-.361l-.078-.178l-.079-.182l-.082-.16l-.079-.18l-.079-.179l-.079-.182l-.081-.182l-.081-.18l-.16-.361l-.081-.201l-.081-.162l-.079-.186l-.077-.156l-.081-.16l-.098-.159l-.09-.182l-.1-.162l-.081-.159l-.089-.162l-.082-.16l-.079-.162l-.087-.155l-.081-.16l-.074-.19l-.167-.323l-.081-.16l-.082-.16l-.199-.32l-.182-.301l-.078-.16l-.1-.148l-.102-.162l-.079-.151l-.082-.135l-.168-.28l-.1-.143l-.082-.135l-.078-.141l-.082-.135l-.081-.141l-.085-.135l-.089-.121l-.079-.121l-.1-.121l-.102-.118l-.079-.12l-.101-.121l-.118-.157l-.09-.102l-.093-.114l-.079-.096l-.078-.135l-.1-.1l-.089-.074l-.101-.093l-.083-.087l-.101-.116l-.175-.162l-.092-.081l-.09-.069L1.086.56L.992.499L.906.425L.809.363L.724.304l-.1-.059l-.085-.04l-.092-.04l-.094-.04l-.096-.04L.16.043L.081.012L0 .004h.011z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Locked(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.5 6.5V10H2a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V12a2 2 0 0 0-2-2h-1.5V6.5a6.5 6.5 0 1 0-13 0M6 10V6.5a4 4 0 0 1 8 0V10zm2 5.5a2 2 0 1 1 3.092 1.676l-.008.005s.195 1.18.415 2.57v.001a.749.749 0 0 1-.749.749H9.248a.749.749 0 0 1-.749-.749v-.001l.415-2.57a2.002 2.002 0 0 1-.916-1.68z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func LowVision(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.168 5.658c.985.124 1.863.308 2.714.557l-.122-.03c.884.279 1.621.572 2.332.912l-.107-.046l.226.113l-.633.38c-.353.206-.654.38-.673.38a2.466 2.466 0 0 0-.328.192l.008-.006c-.16.107-.414.26-.566.346l-.28.154l-.434-.406a4.705 4.705 0 0 0-3.28-1.326a4.7 4.7 0 0 0-2.373.639l.022-.012a6.124 6.124 0 0 0-1.712 1.678l-.014.021a4.928 4.928 0 0 0-.438 3.799l-.009-.035c.034.096.068.219.096.345l.004.023c-.318.238-.68.453-1.063.629l-.037.015c-.014-.02-.1-.26-.194-.546a5.923 5.923 0 0 1-.286-1.836c0-1.325.427-2.55 1.151-3.545l-.012.017l.306-.414l-.346.138a15.706 15.706 0 0 0-4.916 3.359l-.501.508l.466.434a21.717 21.717 0 0 0 2.418 1.924l.074.048c.208.143.454.293.708.43l.045.022a.058.058 0 0 1 .054.053a7.528 7.528 0 0 1-1.188.726l-.045.02a16.154 16.154 0 0 1-1.794-1.281l.028.022a25.565 25.565 0 0 1-1.553-1.427l-.921-.921l.146-.207a26.124 26.124 0 0 1 2.475-2.472l.031-.026A14.593 14.593 0 0 1 9.95 5.704l.08-.01a16.743 16.743 0 0 1 3.195-.032l-.062-.004zm8.23.7a.781.781 0 0 1 .265.224l.001.002a.783.783 0 0 1-.048.729l.002-.003c-.066.073-2.985 1.806-6.48 3.851s-7.52 4.391-8.93 5.218c-2.046 1.193-2.613 1.499-2.806 1.499l-.04.001a.55.55 0 0 1-.532-.412l-.001-.004a.411.411 0 0 1 .027-.462l-.001.001c.106-.22.64-.546 5.25-3.238c2.826-1.646 6.87-4.012 8.996-5.25s3.913-2.258 3.985-2.258a.907.907 0 0 1 .32.105l-.005-.002zm-8.46 1.513a4.104 4.104 0 0 1 1.644.878l-.005-.004l.374.346l-.52.3c-.286.166-.546.306-.566.32a1.716 1.716 0 0 1-.457-.249l.004.003a2.34 2.34 0 0 0-1.576-.393l.01-.001a2.553 2.553 0 0 0-2.119 1.431l-.007.015a1.868 1.868 0 0 0-.219 1.108l-.001-.008v.66l-.5.286c-.161.108-.346.2-.543.268l-.017.005a2.311 2.311 0 0 1-.233-1.249l-.001.009l-.001-.087a3.59 3.59 0 0 1 1.149-2.637l.002-.002a3.712 3.712 0 0 1 3.605-.995l-.026-.006zm8.123 1.067a21.723 21.723 0 0 1 2.919 2.73l.019.022a21.788 21.788 0 0 1-2.005 1.983l-.027.023a16.214 16.214 0 0 1-8.411 3.815l-.085.01a15.435 15.435 0 0 1-4.564-.251l.1.017a13.609 13.609 0 0 1-2.158-.659l.093.033a22.604 22.604 0 0 1 2.071-1.209l.135-.063a.72.72 0 0 1 .268.139l-.001-.001a4.89 4.89 0 0 0 2.904.772h-.011a4.723 4.723 0 0 0 4.432-4.711l-.001-.084v.004l.006-.619l.546-.313a1.7 1.7 0 0 1 .581-.264l.012-.002a6.297 6.297 0 0 1 .104 1.865l.002-.025a7.46 7.46 0 0 1-.163 1.029l.01-.05a7.158 7.158 0 0 1-.891 1.899l.017-.026a1.546 1.546 0 0 0-.209.349l-.004.01c.28-.09.518-.192.745-.312l-.025.012a16.99 16.99 0 0 0 4.174-2.791l-.014.012l.659-.586l-.493-.466a22.474 22.474 0 0 0-1.949-1.595l-.064-.045c-.226-.147-.306-.234-.254-.286a8.252 8.252 0 0 1 1.051-.623l.049-.022c.165.082.308.174.439.281l-.004-.003zm-5.191 2.765a4.147 4.147 0 0 1-.477 1.709l.011-.023a3.858 3.858 0 0 1-4.936 1.716l.025.01c-.552-.24-.552-.24.147-.645l.645-.368l.706-.006a1.93 1.93 0 0 0 1.161-.225l-.01.005a2.541 2.541 0 0 0 1.194-1.236l.006-.016c.186-.414.214-.44.806-.793c.186-.12.406-.245.634-.356l.039-.017c.029 0 .049.113.049.246z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mad(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.316C6.099 1.316 1.316 6.099 1.316 12S6.099 22.684 12 22.684S22.684 17.901 22.684 12c-.012-5.896-4.788-10.672-10.683-10.684zm0 22.297C5.586 23.613.387 18.414.387 12S5.586.387 12 .387S23.613 5.586 23.613 12v.015c0 6.405-5.192 11.597-11.597 11.597z"/><path fill="currentColor" d="M12 24C5.386 23.966.034 18.614 0 12.003V12C0 5.373 5.373 0 12 0s12 5.373 12 12c-.034 6.614-5.386 11.966-11.997 12zM12 .774C5.8.774.774 5.8.774 12S5.8 23.226 12 23.226S23.226 18.2 23.226 12C23.222 5.802 18.198.779 12.001.774zm0 22.297C5.886 23.071.929 18.114.929 12S5.886.929 12 .929c6.122.021 11.084 4.956 11.148 11.065V12c-.065 6.115-5.027 11.05-11.146 11.071zm0-21.368C6.313 1.703 1.703 6.313 1.703 12S6.313 22.297 12 22.297S22.297 17.687 22.297 12c-.026-5.676-4.62-10.271-10.294-10.297z"/><path fill="currentColor" d="M9.677 9.91a2.09 2.09 0 0 1-4.18 0v-.009c0-1.15.932-2.082 2.082-2.082h.018c1.15 0 2.082.932 2.082 2.082v.009zm8.904 0v.009c0 1.15-.932 2.082-2.082 2.082h-.009a2.09 2.09 0 0 1 0-4.18a2.132 2.132 0 0 1 2.09 2.088v.002zm-1.936 8.903a.355.355 0 0 1-.387-.308v-.002c-.6-1.813-2.279-3.097-4.258-3.097s-3.658 1.285-4.249 3.066l-.009.032a.442.442 0 1 1-.851-.235l-.001.003c.72-2.175 2.735-3.717 5.11-3.717s4.39 1.542 5.099 3.679l.011.038a.44.44 0 0 1-.306.541l-.003.001h-.155z"/><path fill="currentColor" d="m16.645 19.2l-.039.001a.769.769 0 0 1-.733-.537l-.001-.005c-.529-1.63-2.034-2.788-3.809-2.788l-.065.001h.003a4.096 4.096 0 0 0-3.862 2.758l-.009.029a.839.839 0 0 1-1.09.541l.006.002a.795.795 0 0 1-.462-.383l-.002-.004a.858.858 0 0 1-.076-.626l-.001.006c.743-2.308 2.872-3.948 5.384-3.948h.037h-.002h.036a5.648 5.648 0 0 1 5.373 3.909l.011.04a.58.58 0 0 1-.079.62l.001-.001c-.077.232-.31.31-.465.387h-.155v-.774a.077.077 0 0 0 .077-.077c-.685-1.991-2.538-3.399-4.722-3.406H12a5.004 5.004 0 0 0-4.712 3.371l-.011.035v.077h.077c.628-1.947 2.424-3.33 4.543-3.33l.107.001h-.005a4.908 4.908 0 0 1 4.635 3.295l.01.034zM9.987 7.819a.279.279 0 0 1-.232-.078L7.2 6.193a1.05 1.05 0 0 1-.229-.304l-.003-.006a.462.462 0 0 1 .078-.311l-.001.002a1.05 1.05 0 0 1 .304-.229l.006-.003a.462.462 0 0 1 .311.078l-.002-.001l2.555 1.548c.094.087.172.189.229.304l.003.006a.462.462 0 0 1-.078.311l.001-.002a.426.426 0 0 1-.379.232h-.008z"/><path fill="currentColor" d="M9.987 8.206a1.085 1.085 0 0 1-.469-.158l.005.003l-2.555-1.548a.987.987 0 0 1-.385-.535l-.002-.007a.815.815 0 0 1 .079-.623l-.002.004a.987.987 0 0 1 .535-.385l.007-.002a.815.815 0 0 1 .623.079l-.004-.002l2.555 1.548a.987.987 0 0 1 .385.535l.002.007a.815.815 0 0 1-.079.623l.002-.004a.746.746 0 0 1-.69.465zm-2.477-2.4h-.077v.077l2.555 1.548l.077-.077zm6.503 2.013h-.008a.424.424 0 0 1-.378-.23l-.001-.002a.46.46 0 0 1-.077-.312v.002a.417.417 0 0 1 .23-.309l.002-.001l2.555-1.548a.46.46 0 0 1 .312-.077h-.002c.138.023.251.11.309.23l.001.002a.46.46 0 0 1 .077.312v-.002a.417.417 0 0 1-.23.309l-.002.001l-2.555 1.548a.285.285 0 0 1-.233.078h.001z"/><path fill="currentColor" d="M14.013 8.206h-.004a.81.81 0 0 1-.69-.384l-.002-.003a.923.923 0 0 1-.154-.623V7.2a.899.899 0 0 1 .384-.54l.003-.002l2.555-1.548a.81.81 0 0 1 .625-.076l-.006-.001c.216.087.397.219.54.385l.002.002a.81.81 0 0 1 .076.625l.001-.006a.899.899 0 0 1-.384.54l-.003.002l-2.555 1.548a.57.57 0 0 1-.39.077zm-.078-.774l2.632-1.471v-.077h-.077zl-.387.31z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Magento(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.7 19.368V7.8l-7.5-4.632L2.7 7.8v11.568L0 17.7V6.3L10.2 0l10.2 6.3v11.4zm-9 1.306l1.5.926l1.5-.926V7.262L15 9.3v11.735L10.2 24l-4.8-2.965V9.3l3.3-2.038z"/><path fill="currentColor" d="m17.7 7.8l-7.5-4.632L2.7 7.8v.009L0 6.3L10.2 0l10.2 6.3l-2.7 1.509zM15 9.32l-3.3 1.844V7.262L15 9.3zm-6.3 1.844L5.4 9.318V9.3l3.3-2.038z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Magnet(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14.13 22.716l-.001.047c0 .42.329.763.743.786h.002l5.433.447a.807.807 0 0 0 .886-.694v-.004l.424-4.479l-7.059-.586zM.742 18.813l.425 4.479a.807.807 0 0 0 .89.698h-.004l5.433-.446a.79.79 0 0 0 .745-.787l-.001-.047v.002l-.425-4.479zm-.7-7.385L.48 16.05l7.062-.581l-.437-4.617a3.81 3.81 0 0 1-.015-.339a3.979 3.979 0 0 1 4.096-3.844h-.006l.114-.002a3.977 3.977 0 0 1 3.975 3.839v.007c0 .114-.006.229-.015.34l-.436 4.617l7.062.581l.438-4.622c.027-.306.042-.612.042-.916C22.358 4.72 17.342 0 11.179 0S0 4.719 0 10.512c0 .304.014.606.042.916"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Male(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.145 16.629a23.876 23.876 0 0 1-.052-2.525l-.001.037a4.847 4.847 0 0 0 1.333-2.868l.002-.021c.339-.028.874-.358 1.03-1.666a1.217 1.217 0 0 0-.455-1.218l-.003-.002c.552-1.66 1.698-6.796-2.121-7.326C13.485.35 12.479 0 11.171 0c-5.233.096-5.864 3.951-4.72 8.366a1.222 1.222 0 0 0-.455 1.229l-.001-.008c.16 1.306.691 1.638 1.03 1.666a4.858 4.858 0 0 0 1.374 2.888a24.648 24.648 0 0 1-.058 2.569l.005-.081C7.308 19.413.32 18.631 0 24h22.458c-.322-5.369-7.278-4.587-8.314-7.371z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Map(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m26.602 23.85l-6.066-4.073l-6.073 4.073c-.132.091-.295.145-.472.145s-.34-.054-.474-.147l.003.002l-6.068-4.073l-6.067 4.073a.836.836 0 0 1-.928.011l.003.002a1.06 1.06 0 0 1-.46-.874v-.021v.001V5.393c0-.352.173-.663.439-.853l.003-.002L6.983.146c.132-.091.295-.145.472-.145s.34.054.474.147L7.926.146l6.066 4.073L20.059.146c.132-.091.295-.145.472-.145s.34.054.474.147l-.003-.002l6.541 4.392c.266.189.438.497.438.844l-.001.038v-.002v17.581c0 .361-.182.679-.459.869l-.004.002a.83.83 0 0 1-.45.131a.863.863 0 0 1-.472-.152l.003.002zM2 6.056v15.101l4.436-3.086V2.97zm12.993 15.045l4.56-3.062v-6.823c-.044.044-.08.087-.118.131c-.113.131-.225.262-.338.387l-.75-.662c.106-.125.218-.25.331-.381c.175-.194.35-.4.538-.606l.331.3V2.938L14.986 6v7.017a4.11 4.11 0 0 0 .435-.047l-.023.003l.16.987a6.7 6.7 0 0 1-.531.056h-.031zm-6.56-3.055l4.56 3.055V12.96l-.262.8a5.654 5.654 0 0 1-1.316-.625l.022.013l.531-.85c.295.191.633.358.991.483l.034.01V5.999l-4.56-3.061v6.304c.296.157.545.312.783.483l-.02-.014l-.581.812c-.056-.044-.118-.08-.181-.125zM22.246 9.28a3.579 3.579 0 0 0-.716.234l.023-.009v8.575l4.429 3.02V5.994l-4.436-3.033v5.467c.129-.044.297-.089.468-.125l.032-.006zm-5.735 3.348c.37-.184.689-.405.972-.665l-.003.003l.662.75a5.166 5.166 0 0 1-1.164.798l-.03.014zm-6.923-1.124c-.056-.075-.118-.144-.181-.218l.762-.65l.187.225c.218.256.418.5.64.72l-.687.725a9.551 9.551 0 0 1-.711-.782l-.014-.018zm-6.379.24l.013-.05c.136-.487.333-.912.586-1.297l-.011.017l.825.56c-.188.295-.34.636-.438 1L4.178 12l-.018.056zM24.114 9.7l-.64.65l-.712-.712l.65-.64l-.65-.64l.712-.712l.64.65l.64-.65l.712.712l-.65.64l.65.64l-.712.712zM4.497 9.53a3.54 3.54 0 0 1 .689-.509l.018-.009c.169-.093.368-.18.574-.249l.025-.007l.306.95c-.15.044.075.038-.062.113c-.334.132-.62.281-.888.456l.019-.012z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarker(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.075 23.52C1.264 13.642 0 12.629 0 9c0-4.971 4.029-9 9-9s9 4.029 9 9c0 3.629-1.264 4.64-8.075 14.516a1.126 1.126 0 0 1-1.847.004l-.002-.004z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MapMarkerAlt(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.075 23.52C1.264 13.642 0 12.629 0 9c0-4.971 4.029-9 9-9s9 4.029 9 9c0 3.629-1.264 4.64-8.075 14.516a1.126 1.126 0 0 1-1.847.004l-.002-.004zM9 12.75a3.75 3.75 0 1 0 0-7.5a3.75 3.75 0 0 0 0 7.5"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mars(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23 0c.549.008.992.451 1 .999V7.52a.48.48 0 0 1-.48.48h-.021h.001h-1.02a.48.48 0 0 1-.48-.48v-.021v.001v-4.094L16.031 9.39A8.623 8.623 0 0 1 18 14.893l-.001.112v-.006v.026a8.775 8.775 0 0 1-.733 3.522l.022-.057a8.93 8.93 0 0 1-4.742 4.778l-.058.022c-1.03.449-2.231.711-3.492.711s-2.462-.261-3.55-.733l.058.022a8.93 8.93 0 0 1-4.778-4.742l-.022-.058c-.449-1.03-.711-2.231-.711-3.492s.261-2.462.733-3.55l-.022.058a8.93 8.93 0 0 1 4.742-4.778l.058-.022a8.649 8.649 0 0 1 3.466-.711h.028h-.001l.106-.001c2.096 0 4.019.744 5.518 1.981l-.015-.012l5.969-5.969h-4.1a.48.48 0 0 1-.48-.48v-.021v.001v-1.02a.48.48 0 0 1 .48-.48h.021h-.001zM9 22h.015a7.015 7.015 0 1 0-4.96-2.054a6.722 6.722 0 0 0 4.846 2.055L9.006 22z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MarsDouble(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.574 5.571v-.016c0-.228.185-.413.413-.413h.017h-.001h3.857a.87.87 0 0 1 .857.857v3.873a.413.413 0 0 1-.413.413h-.017h.001h-.873a.413.413 0 0 1-.413-.413v-.017v.001v-1.794L20.6 11.477a7.643 7.643 0 0 1 1.689 4.809c0 .482-.044.955-.129 1.412l.007-.047a7.398 7.398 0 0 1-2.218 4.175l-.004.003a7.72 7.72 0 0 1-12.623-2.946l-.017-.054a7.434 7.434 0 0 1-3.003-.79l.043.02a7.909 7.909 0 0 1-2.371-1.775l-.006-.006a7.672 7.672 0 0 1-1.503-2.524L.448 13.7a7.5 7.5 0 0 1-.446-2.571c0-1.433.395-2.773 1.081-3.919l-.019.035a7.808 7.808 0 0 1 2.409-2.547l.029-.018a7.655 7.655 0 0 1 4.213-1.249c.766 0 1.506.112 2.205.319l-.055-.014a8.004 8.004 0 0 1 2.675 1.392l-.017-.013l3.415-3.402h-1.81a.413.413 0 0 1-.413-.413v-.017v.001v-.873c0-.228.185-.413.413-.413h.017h-.001h3.864c.234 0 .445.098.595.254a.82.82 0 0 1 .254.595v3.881a.413.413 0 0 1-.413.413h-.017h.001h-.873a.413.413 0 0 1-.413-.413v-.017v.001v-1.795l-3.401 3.415a7.808 7.808 0 0 1 1.225 2.209l.017.055a7.52 7.52 0 0 1 4.414 1.672l-.014-.011l3.415-3.402h-1.814a.413.413 0 0 1-.413-.413v-.017v.001zm-6.856 5.571q0-.268-.054-.777a5.876 5.876 0 0 0-3.623 2.001l-.007.008a5.74 5.74 0 0 0-1.459 3.915v-.004q0 .268.054.777a5.876 5.876 0 0 0 3.623-2.001l.007-.008a5.74 5.74 0 0 0 1.46-3.835l-.001-.08zm-12 0v.071c0 1.49.564 2.848 1.491 3.871l-.004-.005a5.859 5.859 0 0 0 3.666 1.992l.03.003a7.566 7.566 0 0 1-.04-.783v-.008l-.001-.09c0-1.862.685-3.564 1.816-4.868l-.008.009a7.538 7.538 0 0 1 4.481-2.624l.046-.007c-.95-2.116-3.038-3.562-5.464-3.562h-.015h.001L7.63 5.14a5.763 5.763 0 0 0-4.152 1.759l-.001.002a5.765 5.765 0 0 0-1.758 4.245v-.005zm12.857 11.143a6.012 6.012 0 0 0 6-5.999v-.072c0-1.49-.564-2.848-1.491-3.871l.004.005a5.859 5.859 0 0 0-3.666-1.992l-.03-.003a7.665 7.665 0 0 1-1.777 5.75l.009-.011a7.544 7.544 0 0 1-4.482 2.623l-.046.007c.95 2.116 3.038 3.562 5.464 3.562z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MarsStroke(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23 0c.549.008.992.451 1 .999V7.52a.48.48 0 0 1-.48.48h-.021h.001h-1.02a.48.48 0 0 1-.48-.48v-.021v.001v-4.094L18.672 6.75l2.187 2.187a.52.52 0 0 1 .141.356v.013a.465.465 0 0 1-.141.334l-.72.72a.465.465 0 0 1-.334.141h-.013a.524.524 0 0 1-.357-.141l-2.187-2.203l-1.218 1.234a8.623 8.623 0 0 1 1.969 5.503l-.001.112V15v.026a8.778 8.778 0 0 1-.733 3.522l.022-.057a8.93 8.93 0 0 1-4.742 4.778l-.058.022c-1.03.449-2.231.711-3.492.711s-2.462-.261-3.55-.733l.058.022a8.93 8.93 0 0 1-4.778-4.742l-.022-.058c-.449-1.03-.711-2.231-.711-3.492s.261-2.462.733-3.55l-.022.058a8.93 8.93 0 0 1 4.742-4.778l.058-.022a8.649 8.649 0 0 1 3.466-.711h.028h-.001l.106-.001c2.096 0 4.019.744 5.518 1.981l-.015-.012l1.218-1.218l-2.687-2.687a.52.52 0 0 1-.141-.356V3.69c0-.131.054-.249.141-.334l.72-.72a.465.465 0 0 1 .334-.141h.013c.138 0 .263.054.357.141l2.687 2.687l3.328-3.328h-4.1a.48.48 0 0 1-.48-.48v-.021v.001v-1.02a.48.48 0 0 1 .48-.48h.021h-.001zM9 22h.015a7.015 7.015 0 1 0-4.96-2.054a6.722 6.722 0 0 0 4.846 2.055L9.006 22z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MarsStrokeH(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M30.892 11.027a1.017 1.017 0 0 1 0 1.462l-4.777 4.777a.492.492 0 0 1-.73 0l-.731-.731a.491.491 0 0 1-.001-.73l3.006-3.006h-4.78v3.66a.5.5 0 0 1-.5.5h-.021h.001h-1.06a.5.5 0 0 1-.5-.5v-.021v.001v-3.64h-2.144a9.08 9.08 0 0 1-3.037 5.915l-.01.009a8.99 8.99 0 0 1-6.13 2.401l-.125-.001h.006h-.024A9.336 9.336 0 0 1 .068 10.645l-.004.045a9.097 9.097 0 0 1 1.45-4.025l-.021.034a9.463 9.463 0 0 1 3.014-2.935l.041-.023a9.05 9.05 0 0 1 4.014-1.306l.033-.002a9.364 9.364 0 0 1 8.367 3.848l.019.027a9.027 9.027 0 0 1 1.671 4.373l.003.039h2.146V7.06a.5.5 0 0 1 .5-.5h.021h-.001h1.06a.5.5 0 0 1 .5.5v.021v-.001v3.64h4.777l-3.006-3.006a.492.492 0 0 1 0-.73l.731-.731a.491.491 0 0 1 .73-.001zM9.359 19.038l.102.001a6.992 6.992 0 0 0 5.039-2.136l.002-.002c1.32-1.314 2.137-3.133 2.137-5.143s-.817-3.829-2.137-5.143a6.994 6.994 0 0 0-5.041-2.138l-.108.001h.005l-.102-.001a6.992 6.992 0 0 0-5.039 2.136l-.002.002c-1.32 1.314-2.137 3.133-2.137 5.143s.817 3.829 2.137 5.143a6.994 6.994 0 0 0 5.041 2.138l.108-.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MarsStrokeV(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.57 8.626c3.874.446 6.854 3.705 6.858 7.66v.001a7.71 7.71 0 0 1-11.916 6.461l.03.018a7.794 7.794 0 0 1-2.418-2.483l-.019-.034a7.478 7.478 0 0 1-1.077-3.308l-.002-.027a7.717 7.717 0 0 1 3.172-6.896l.022-.016a7.433 7.433 0 0 1 3.604-1.376l.032-.003V6.857H4.697a.413.413 0 0 1-.413-.413v-.017v.001v-.873c0-.228.185-.413.413-.413h.017h-.001h2.143v-2.21L5.624 4.164a.445.445 0 0 1-.306.121h-.01a.395.395 0 0 1-.286-.121l-.616-.616a.398.398 0 0 1-.121-.287v-.01c0-.118.046-.226.121-.306L7.112.253A.839.839 0 0 1 8.318.252l2.706 2.692c.075.08.121.187.121.306v.01a.395.395 0 0 1-.121.286l-.616.616a.398.398 0 0 1-.287.121h-.01a.447.447 0 0 1-.306-.121L8.573 2.93v2.21h2.159c.228 0 .413.185.413.413v.017v-.001v.873a.413.413 0 0 1-.413.413h-.017h.001h-2.145zm-.857 13.661a6.014 6.014 0 0 0 6-5.999v-.001a6 6 0 1 0-12 0a6.014 6.014 0 0 0 5.999 6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mastercard(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M33.6 24H2.4A2.4 2.4 0 0 1 0 21.6V2.4A2.4 2.4 0 0 1 2.4 0h31.2A2.4 2.4 0 0 1 36 2.4v19.2a2.4 2.4 0 0 1-2.4 2.4m-5.116-4.281h-.014a.988.988 0 0 0-.728.319l-.001.001a1.17 1.17 0 0 0 .001 1.579l-.001-.001c.182.197.441.32.729.32h.014h-.001a.769.769 0 0 0 .671-.31l.001-.002v.266h.453v-2.969h-.453v1.11a.792.792 0 0 0-.631-.314l-.042.001zm-4.25 0h-.023c-.289 0-.55.12-.735.313a1.19 1.19 0 0 0 .001 1.595l-.001-.001c.186.193.447.313.736.313h.024h-.001l.045.001a.749.749 0 0 0 .609-.312l.002-.002v.266h.47v-2.11h-.469v.25a.777.777 0 0 0-.658-.313h.002zm-2.048 0l-.042-.001a1.11 1.11 0 0 0 0 2.22l.044-.001h-.002l.041.001c.27 0 .517-.094.712-.252l-.002.002l-.219-.375a.875.875 0 0 1-.506.188h-.002a.65.65 0 0 1-.5-.188a.741.741 0 0 1 .001-.969l-.001.001a.61.61 0 0 1 .441-.188h.021h-.001a.898.898 0 0 1 .549.189l-.002-.002l.219-.375a1.165 1.165 0 0 0-.726-.251h-.025h.001zm-4.08 0h-.025a.992.992 0 0 0-.725.313a1.117 1.117 0 0 0-.297.802v-.002l-.001.046c0 .295.117.563.306.759c.19.189.452.305.741.305l.034-.001h-.002l.041.001c.313 0 .599-.113.821-.3l-.002.002l-.219-.343a.93.93 0 0 1-.608.234h-.001a.594.594 0 0 1-.64-.513v-.003h1.578v-.187a1.137 1.137 0 0 0-.283-.806l.001.001a.936.936 0 0 0-.696-.309h-.022h.001zm-3.077.063v.422h.437v.954c0 .518.247.781.734.781h.017c.202 0 .389-.06.546-.162l-.004.002l-.126-.391a.853.853 0 0 1-.406.11h-.001c-.198 0-.298-.115-.298-.343v-.95h.75v-.422h-.75v-.64h-.469v.64zm-1.967 1.53l-.203.36c.247.167.551.266.878.266h.03h-.002l.066.002c.239 0 .462-.071.649-.193l-.005.003a.59.59 0 0 0 .258-.487v-.014v.001c0-.352-.247-.56-.734-.625l-.218-.032c-.24-.041-.36-.114-.36-.218c0-.156.131-.234.391-.234h.005c.245 0 .475.064.674.176l-.007-.004l.188-.375a1.567 1.567 0 0 0-.863-.219h.003l-.054-.001a.999.999 0 0 0-.589.191l.003-.002a.607.607 0 0 0-.234.48v.021v-.001c0 .342.247.547.734.609l.203.032c.249.042.375.115.375.219c0 .176-.16.266-.485.266h-.009a1.2 1.2 0 0 1-.698-.222l.004.003zm-1.782-1.594h-.025a.99.99 0 0 0-.724.313a1.215 1.215 0 0 0 .001 1.596l-.001-.001a.992.992 0 0 0 .725.313h.026h-.001a.792.792 0 0 0 .672-.311l.001-.002v.266h.453v-2.11h-.453v.25a.846.846 0 0 0-.658-.313zh.001zm14.688.063v2.11h.453v-1.187c0-.362.152-.546.453-.546h.01c.099 0 .193.023.276.064l-.004-.002l.14-.438a.722.722 0 0 0-.298-.063l-.032.001h.002h-.017a.604.604 0 0 0-.528.31l-.002.003v-.25zm-6.39 0v2.11h.469v-1.187c0-.362.152-.546.453-.546h.01c.099 0 .193.023.276.064l-.004-.002l.141-.438a.78.78 0 0 0-.31-.063h-.019h.001h-.018a.602.602 0 0 0-.527.31l-.002.003v-.25zm-10.688.375c.311 0 .469.189.469.56v1.172h.469V20.56l.001-.033a.815.815 0 0 0-.228-.566a.889.889 0 0 0-.609-.24h-.009a.806.806 0 0 0-.732.372l-.002.003a.755.755 0 0 0-.705-.375h.002h-.009a.733.733 0 0 0-.599.311l-.002.002v-.25h-.467v2.11h.469V20.72c0-.373.174-.56.516-.56c.311 0 .469.189.469.56v1.172h.453V20.72c0-.374.174-.56.515-.56zM22.906 2h-.033a7.775 7.775 0 0 0-4.385 1.346l.027-.017a8.449 8.449 0 0 1 2.758 4.089l.016.06c.236.745.371 1.601.371 2.49c0 .879-.133 1.726-.379 2.524l.016-.06a8.29 8.29 0 0 1-2.766 4.118l-.016.013a7.833 7.833 0 0 0 4.385 1.328c1.12 0 2.185-.233 3.15-.653l-.051.02a7.893 7.893 0 0 0 4.207-4.175l.019-.051c.4-.91.633-1.971.633-3.086s-.233-2.176-.653-3.136l.02.05a7.877 7.877 0 0 0-4.183-4.205l-.051-.019a7.664 7.664 0 0 0-3.075-.633h-.011h.001zm-9.812 0h-.011a7.77 7.77 0 0 0-3.125.652l.05-.02a7.873 7.873 0 0 0-4.216 4.175l-.019.051c-.4.91-.633 1.971-.633 3.086s.233 2.176.653 3.136l-.02-.05a7.893 7.893 0 0 0 4.175 4.207l.051.019c.914.4 1.98.633 3.099.633a7.863 7.863 0 0 0 4.413-1.346l-.028.018a8.36 8.36 0 0 1-2.766-4.089l-.015-.059a8.372 8.372 0 0 1-.361-2.459a8.543 8.543 0 0 1 3.128-6.616l.015-.012a7.741 7.741 0 0 0-4.357-1.328h-.036h.002zM18 3.703a7.794 7.794 0 0 0-2.673 3.842l-.015.055a7.753 7.753 0 0 0-.36 2.359c0 .834.129 1.637.368 2.392l-.015-.056a7.724 7.724 0 0 0 2.677 3.879l.018.013a7.715 7.715 0 0 0 2.681-3.837l.015-.054a7.743 7.743 0 0 0 .353-2.337c0-.843-.132-1.654-.375-2.416l.015.056a7.813 7.813 0 0 0-2.67-3.884l-.018-.013zM28.546 21.5h-.017a.6.6 0 0 1-.443-.195a.72.72 0 0 1 .001-.955l-.001.001a.639.639 0 0 1 .922 0a.73.73 0 0 1-.001.961l.001-.001a.613.613 0 0 1-.444.189zh.001zm-4.265 0h-.022a.59.59 0 0 1-.431-.187a.764.764 0 0 1 .001-.97l-.001.001a.587.587 0 0 1 .431-.188h.023h-.001h.018c.175 0 .333.075.442.195c.112.121.18.283.18.461v.02v-.001l.001.027a.692.692 0 0 1-.173.459l.001-.001a.605.605 0 0 1-.436.184l-.034-.001h.002zm-12.938 0l-.029.001a.603.603 0 0 1-.439-.188a.764.764 0 0 1 .001-.97l-.001.001a.607.607 0 0 1 .471-.187h-.001h.023a.57.57 0 0 1 .43.195l.001.001a.757.757 0 0 1-.001.961l.001-.001a.59.59 0 0 1-.433.188h-.022h.001zm7.297-.86h-1.093a.55.55 0 0 1 1.094-.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Maxcdn(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M32.72 9.694L29.644 24h-6.262L26.72 8.4a1.763 1.763 0 0 0-.285-1.654l.003.004a1.89 1.89 0 0 0-1.563-.618h.007h-3.17L17.886 24h-6.262l3.825-17.869h-5.36L6.262 24H0L3.825 6.131L.96 0h23.975c1.268 0 2.471.28 3.55.781l-.052-.022a7.586 7.586 0 0 1 2.756 2.114l.011.014a7.354 7.354 0 0 1 1.511 3.11l.009.049a8.565 8.565 0 0 1-.011 3.705z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Medium(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24 24H0V0h24zM15.014 8.994v7.326c0 .198 0 .234-.127.362l-1.302 1.264v.27h6.32v-.27l-1.257-1.234a.376.376 0 0 1-.143-.364v.002v-9.07a.375.375 0 0 1 .143-.36l.001-.001l1.286-1.234v-.27h-4.456l-3.176 7.924l-3.609-7.924H4.019v.271l1.502 1.813a.627.627 0 0 1 .204.53v-.003v7.126a.823.823 0 0 1-.22.707l-1.69 2.054v.27h4.8v-.27l-1.691-2.054a.852.852 0 0 1-.233-.712v.004v-6.16l4.215 9.195h.49z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Meetup(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M28.595 23.062c-.38.374-.86.649-1.395.781l-.021.004a7.19 7.19 0 0 1-1.358.126a7.322 7.322 0 0 1-3.264-.762l.043.019c-4.365-1.92-3.401-6.785-1.486-10.139c.575-1.007 1.142-2.022 1.713-3.04c.449-.8 1.421-2.155 1.04-3.136c-.4-1.029-1.467-1.035-2.168-.168a15.796 15.796 0 0 0-1.669 2.901l-.04.102c-.507 1.058-3.04 6.618-3.04 6.618a10.41 10.41 0 0 1-2.107 2.925l-.002.002a2.197 2.197 0 0 1-1.381.484c-.441 0-.851-.129-1.196-.351l.009.005a1.514 1.514 0 0 1-.58-1.65l-.003.011c.527-3.022 5.111-10.054 1.95-10.55c-1.212-.19-1.541 1.158-1.914 2.019c-.618 1.422-1.089 2.902-1.749 4.307a23.946 23.946 0 0 0-1.706 4.858l-.034.169c-.32 1.386-.731 3.151-2.308 3.573c-4.32 1.154-5.63-1.696-5.63-1.697c-.705-2.24-.037-4.26.64-6.417c.525-1.666.838-3.385 1.502-5.006C3.626 6.16 4.807.128 9.081.514a8.514 8.514 0 0 1 3.287 1.221l-.035-.021c.856.499 1.508.766 2.505.228c.97-.522 1.414-1.495 2.57-1.829c1.238-.358 2.053.171 2.979.917c1.298 1.04 1.44.572 2.511.298a6.932 6.932 0 0 1 1.846-.244c.37 0 .733.028 1.087.083l-.04-.005c5.01.858 1.819 7.254.624 9.824c-.778 1.672-4.49 8.396-1.2 9.299c.992.272 2.271.148 3.098.86c.838.722.755 1.404.282 1.915z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mercury(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.12 4.232a7.793 7.793 0 0 1 3.101 2.788l.019.031a7.37 7.37 0 0 1 1.186 4.031v.063v-.003l.001.1a7.414 7.414 0 0 1-1.98 5.054l.004-.005A7.476 7.476 0 0 1 8.603 18.8l-.032.003v1.768h1.302c.228 0 .413.185.413.413v.017V21v.873a.413.413 0 0 1-.413.413h-.017h.001h-1.286v1.302a.413.413 0 0 1-.413.413h-.017h.001h-.873a.413.413 0 0 1-.413-.413v-.017v.001v-1.286H5.554a.413.413 0 0 1-.413-.413v-.017v.001v-.873c0-.228.185-.413.413-.413h.017h-.001h1.286V18.8a7.482 7.482 0 0 1-4.873-2.503l-.007-.008a7.409 7.409 0 0 1-1.977-5.05l.001-.1v-.055c0-1.498.442-2.892 1.203-4.06l-.018.029a7.786 7.786 0 0 1 3.083-2.799l.045-.02A6.757 6.757 0 0 1 1.274.624L1.259.577a.395.395 0 0 1 .048-.396l-.001.001a.41.41 0 0 1 .34-.181h.015h-.001h.943c.172 0 .317.111.369.266l.001.003a5.094 5.094 0 0 0 1.856 2.281l.018.011a4.957 4.957 0 0 0 2.819.871h.05h-.003h.047c1.053 0 2.03-.326 2.836-.882l-.017.011A5.079 5.079 0 0 0 12.442.303l.012-.033a.473.473 0 0 1 .499-.268h-.003h.832c.141 0 .265.071.339.18l.001.001a.4.4 0 0 1 .046.398l.001-.003a6.762 6.762 0 0 1-3.018 3.639l-.033.017zm-3.402 12.91l.087.001a5.763 5.763 0 0 0 4.152-1.759l.001-.002a5.765 5.765 0 0 0 1.761-4.153l-.001-.092v.005l.001-.087a5.763 5.763 0 0 0-1.759-4.152l-.002-.001C10.875 5.813 9.375 5.14 7.718 5.14s-3.157.674-4.24 1.762a5.765 5.765 0 0 0-1.76 4.245v-.005l-.001.087c0 1.629.674 3.101 1.759 4.152l.002.001a5.761 5.761 0 0 0 4.154 1.761zh-.004z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Messenger(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 0C5.373 0 0 4.974 0 11.111a10.8 10.8 0 0 0 4.438 8.633l.031.021V24l4.088-2.242c1.031.295 2.215.464 3.439.464H12c6.627 0 12-4.975 12-11.11S18.627 0 12 0m1.191 14.963L10.136 11.7l-5.963 3.26L10.732 8l3.131 3.259L19.752 8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Meteor(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 0l22.349 20.946a.928.928 0 0 1-.126 1.174a.977.977 0 0 1-1.264.083l.003.002zm6.638 2.099l17.349 15.95a.922.922 0 0 1-.125 1.174a.976.976 0 0 1-1.262.083l.003.002zM1.975 6.595l17.349 15.95a.931.931 0 0 1-.125 1.175a.973.973 0 0 1-1.262.081l.003.002zm10.273-2.692l12.128 11.145a.649.649 0 0 1-.088.821a.68.68 0 0 1-.882.058l.002.001zm-8.737 7.894l12.123 11.142a.648.648 0 0 1-.09.816a.683.683 0 0 1-.882.059l.002.001zM17.98 6.506l5.534 5.054a.289.289 0 0 1-.044.384a.351.351 0 0 1-.438.027l.001.001zM6.301 17.336l5.53 5.055a.29.29 0 0 1-.042.384a.348.348 0 0 1-.436.026l.001.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Metro(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m4.51 18.946l.713 2.69l1.879.68V24h2.546v-1.684l1.879-.68l.714-2.69z"/><path fill="currentColor" d="M14.962 12.17c0 .45-.365.814-.814.814H2.606a.81.81 0 0 1-.576-.24a.81.81 0 0 1-.24-.576l-.417-9.33c0-.45.365-.814.814-.814h12.374a.81.81 0 0 1 .576.24a.81.81 0 0 1 .24.576zm.352 2.782a.696.696 0 0 1-.694.696h-2.8a.698.698 0 0 1-.696-.696a.698.698 0 0 1 .696-.696h2.8a.697.697 0 0 1 .694.695zm-9.753 0a.698.698 0 0 1-.696.696h-2.8a.698.698 0 0 1-.696-.696a.698.698 0 0 1 .696-.696h2.8a.698.698 0 0 1 .696.695zM15.807 0H.943a.946.946 0 0 0-.944.944v18.135a2.56 2.56 0 0 0 2.56 2.56h1.739l-.836-3.478h9.795l-.835 3.481h1.74a2.56 2.56 0 0 0 2.56-2.56c.014-.064.023-17.051.023-17.051V.95a.947.947 0 0 0-.944-.944z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mic(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.4 16.8a4.805 4.805 0 0 0 4.8-4.8V4.8a4.8 4.8 0 1 0-9.6 0V12a4.805 4.805 0 0 0 4.8 4.8"/><path fill="currentColor" d="M16.8 12V9.6a1.2 1.2 0 1 0-2.4 0V12a6 6 0 1 1-12 0V9.6a1.2 1.2 0 1 0-2.4 0V12a8.406 8.406 0 0 0 7.154 8.298l.046.006V21.6H3.6a1.2 1.2 0 1 0 0 2.4h9.6a1.2 1.2 0 1 0 0-2.4H9.6v-1.296c4.09-.609 7.193-4.093 7.2-8.303z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Microsoft(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v10.719h10.719V0zm13.279 0v10.719h10.719V0zM0 13.281V24h10.719V13.281zm13.279 0V24h10.719V13.281z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MinusA(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.8 13.2v-2.4H1.2v2.4h12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mobile(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.75 0H2.25A2.25 2.25 0 0 0 0 2.25v19.5A2.25 2.25 0 0 0 2.25 24h10.5A2.25 2.25 0 0 0 15 21.75V2.25A2.25 2.25 0 0 0 12.75 0M7.5 22.5a1.498 1.498 0 1 1 .002-2.996A1.498 1.498 0 0 1 7.5 22.5h-.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MobileAlt(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.75 0H2.25A2.25 2.25 0 0 0 0 2.25v19.5A2.25 2.25 0 0 0 2.25 24h10.5A2.25 2.25 0 0 0 15 21.75V2.25A2.25 2.25 0 0 0 12.75 0M7.5 22.5a1.498 1.498 0 1 1 .002-2.996A1.498 1.498 0 0 1 7.5 22.5h-.001zm5.25-5.06a.564.564 0 0 1-.56.56H2.812a.56.56 0 0 1-.56-.56V2.813a.56.56 0 0 1 .56-.56h9.374a.564.564 0 0 1 .56.56z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoneySymbol(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.959 18.414a5.986 5.986 0 1 1-.001-.818l.001.019a5.986 5.986 0 1 1 .001.818zm.793-.374a5.181 5.181 0 1 0 10.363-.001a5.181 5.181 0 0 0-10.363.001M.806 18.014v.026a5.181 5.181 0 1 0 5.181-5.181a5.2 5.2 0 0 0-5.181 5.152v.002zM4.187 20a2.39 2.39 0 0 1-.479-1.143l-.001-.012h-.56v-.617h.48v-.43h-.48v-.617h.59c.101-.47.318-.879.619-1.21l-.002.002a2.694 2.694 0 0 1 2.023-.913h.019h-.001c.485.003.946.101 1.366.278l-.024-.009l-.24 1.02a2.133 2.133 0 0 0-.99-.24H6.48c-.427 0-.81.185-1.075.479l-.001.001c-.129.158-.23.345-.292.549l-.003.011h2.192v.618H4.965v.4h2.363v.618H5.1c.039.24.145.45.296.618l-.001-.001a1.493 1.493 0 0 0 1.158.457h-.004a2.358 2.358 0 0 0 1.034-.246l-.014.006l.19 1.053a2.999 2.999 0 0 1-1.395.349h-.001l-.069.001a2.686 2.686 0 0 1-2.104-1.015l-.004-.005zm11.786.886v-.698c.531-.214.91-.701.966-1.283v-.006a1.374 1.374 0 0 0-.028-.358l.002.009h-.859v-.912h.698a2.875 2.875 0 0 1-.08-.683v-.044a1.78 1.78 0 0 1 1.968-1.77l-.008-.001l.065-.001c.347 0 .675.079.968.221l-.013-.006l-.211.993a1.403 1.403 0 0 0-.756-.16h.004a.73.73 0 0 0-.805.809v-.003c0 .229.03.451.084.663l-.004-.019h1.208v.912h-1.076v.016c0 .214-.029.421-.084.618l.004-.016a1.528 1.528 0 0 1-.428.615l-.002.002v.027h2.309v1.074zM11.959 6.388a5.982 5.982 0 1 1-.001-.799l.001.018a5.958 5.958 0 1 1 .001.798l-.001-.018zm.793-.4a5.181 5.181 0 1 0 10.363.001a5.181 5.181 0 0 0-10.363-.002zm-11.946 0a5.18 5.18 0 0 0 10.362-.002A5.181 5.181 0 0 0 5.988.805A5.2 5.2 0 0 0 .806 5.986v.001zm4.832 3.52V8.376a1.905 1.905 0 0 1-.999-.353l.006.004a1.556 1.556 0 0 1-.592-1.222l.002-.069v.003h.886a.855.855 0 0 0 .692.912l.005.001V6.256a5.76 5.76 0 0 0-.35-.08a1.201 1.201 0 0 1-1.048-1.19v-.025c0-.428.211-.806.535-1.037l.004-.003a1.83 1.83 0 0 1 .851-.319l.008-.001v-1.13h.698v1.155c.347.027.661.154.917.352l-.004-.003c.328.249.538.64.538 1.079l-.001.051v-.002H6.9v-.021a.757.757 0 0 0-.555-.73l-.005-.001v1.235c.215.054.4.108.644.16c.291.072.534.244.696.476l.003.004a1.152 1.152 0 0 1 .241.752c0 .45-.234.845-.586 1.07l-.005.003a1.958 1.958 0 0 1-.989.32H6.34v1.127zm.697-1.828c.457-.054.672-.269.672-.645l.001-.038a.494.494 0 0 0-.348-.472l-.003-.001a1.37 1.37 0 0 0-.311-.106l-.009-.002zM5.128 4.832c0 .269.16.457.51.537V4.295c-.349.081-.51.242-.51.537M17.243 9.2V5.898c-.417.147-.828.299-1.243.449v-.532c.416-.152.833-.302 1.243-.459v-.282c-.414.141-.829.299-1.243.447v-.53c.414-.155.834-.303 1.243-.462V2.567h.668v1.718c.674-.244 1.354-.501 2.034-.746v.551c-.68.244-1.354.492-2.034.736v.284c.679-.246 1.36-.502 2.034-.746v.548c-.674.25-1.356.493-2.034.739v2.847a2.638 2.638 0 0 0 2.299-2.613h.694v.105a3.334 3.334 0 0 1-.353 1.399l.009-.019a3.432 3.432 0 0 1-.742.98l-.003.002a3.278 3.278 0 0 1-2.587.844l.015.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mongodb(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.562 9.518C9.299 3.958 6.32 2.131 6 1.432A8.712 8.712 0 0 1 5.289.06L5.267 0c0 .019 0 .031-.005.049v.144v-.001v.021l.001.011l-.001.011v.052l-.009.05V.4h-.005v.016h-.036v.057h-.006v.046h-.024v.064L5.159.6l-.005.007v.022h-.005v.018h-.006v.045h-.006v.019h-.005v.018h-.005v.022h-.045v.015h-.005v.019h-.005V.8h-.006v.023h-.005v.013a.194.194 0 0 0-.049.034l-.009.01c-.003.004 0 0 0 0v.058h-.005V.93h-.005v.01h-.005v.008h-.005V.97h-.061v.01h-.01V1h-.03v.01h-.005v.006h-.01v.01h-.03v.006h-.005v.058h-.006v.01h-.005v.006h-.005v.006l-.014.016l-.012.01a.416.416 0 0 0-.039.032l-.022.017l-.049.039l-.074.062c-.057.047-.117.1-.186.159c-.169.148-.37.338-.6.568l-.015.015a13.713 13.713 0 0 0-3.729 9.059v.018a10.303 10.303 0 0 0 .007 1.174l-.001-.03v.009a10.853 10.853 0 0 0 1.456 4.808l-.028-.052c.308.54.614.999.948 1.435l-.022-.03a12.916 12.916 0 0 0 2.483 2.503l.031.023a8.197 8.197 0 0 1 .4 2.79v-.011l.644.215a12.16 12.16 0 0 1 .059-2.582l-.006.061c.065-.257.186-.48.35-.664l-.001.002c.298-.213.559-.424.806-.651l-.006.006c.018-.019.028-.036.044-.054a11.395 11.395 0 0 0 3.614-8.337c0-.801-.082-1.582-.239-2.337l.013.074z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoreV(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h5.219v5.219H0zm0 9.39h5.219v5.219H0zm0 9.391h5.219V24H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoreVa(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.217 12a2.608 2.608 0 1 1-5.216 0a2.608 2.608 0 0 1 5.216 0m0-9.392a2.608 2.608 0 1 1-5.216 0a2.608 2.608 0 0 1 5.216 0m0 18.783a2.608 2.608 0 1 1-5.216 0a2.608 2.608 0 0 1 5.216 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Motorcycle(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.632 15.526a2.112 2.112 0 0 0-2.106 2.105v4.305a2.106 2.106 0 0 0 4.212 0v-.043v.002v-4.263a2.112 2.112 0 0 0-2.104-2.106z"/><path fill="currentColor" d="M16.263 2.631H12.21C11.719 1.094 10.303 0 8.631 0S5.544 1.094 5.06 2.604l-.007.027h-4a1.053 1.053 0 0 0 0 2.106h4.053c.268.899.85 1.635 1.615 2.096l.016.009c-2.871.867-4.929 3.48-4.947 6.577v5.528a1.753 1.753 0 0 0 1.736 1.737h1.422v-3a3.737 3.737 0 1 1 7.474 0v3h1.421a1.752 1.752 0 0 0 1.738-1.737v-5.474a6.855 6.855 0 0 0-4.899-6.567l-.048-.012a3.653 3.653 0 0 0 1.625-2.08l.007-.026h4.053a1.056 1.056 0 0 0 1.053-1.053a1.149 1.149 0 0 0-1.104-1.105h-.002zM8.631 5.84a2.106 2.106 0 1 1 2.106-2.106l.001.06c0 1.13-.916 2.046-2.046 2.046l-.063-.001h.003z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoveH(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.781 14.819V9.6H24v5.219zm-9.391 0V9.6h5.219v5.219zm-9.39 0V9.6h5.219v5.219z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MoveHa(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.783 12.208a2.608 2.608 0 1 1 5.215 0a2.608 2.608 0 0 1-5.215.001zm-9.392 0a2.609 2.609 0 1 1 5.217 0a2.609 2.609 0 0 1-5.217.001zm-9.391 0a2.609 2.609 0 1 1 2.609 2.61h-.001A2.608 2.608 0 0 1 0 12.21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func MusicNote(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.07.169a.913.913 0 0 0-.848-.114l.006-.002L6.151 3.991a.924.924 0 0 0-.613.869V17.7a3.847 3.847 0 0 0-1.846-.471h-.001a3.553 3.553 0 0 0-3.692 3.376v.008a3.555 3.555 0 0 0 3.699 3.385h-.007a3.553 3.553 0 0 0 3.692-3.376V7.731l9.23-3.223v8.973a3.86 3.86 0 0 0-1.846-.47h-.001a3.551 3.551 0 0 0-3.691 3.376v.008a3.554 3.554 0 0 0 3.699 3.385h-.007l.105.002a3.622 3.622 0 0 0 3.513-2.74l.005-.025a.908.908 0 0 0 .069-.34V.923a.921.921 0 0 0-.388-.752l-.003-.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Mysql(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m24.129 23.412l-.508-.484a6.66 6.66 0 0 0-.809-.891l-.005-.004q-.448-.407-.931-.774q-.387-.266-1.064-.641a1.6 1.6 0 0 1-.818-.824l-.004-.01l-.048-.024c.212-.021.406-.06.592-.115l-.023.006l.57-.157c.236-.074.509-.122.792-.133h.006c.298-.012.579-.06.847-.139l-.025.006q.194-.048.399-.109t.351-.109v-.169q-.145-.217-.351-.496a2.561 2.561 0 0 0-.443-.468l-.005-.004q-.629-.556-1.303-1.076a15.783 15.783 0 0 0-1.311-.916l-.068-.04a4.983 4.983 0 0 0-.825-.435l-.034-.012q-.448-.182-.883-.399a1.638 1.638 0 0 0-.327-.119l-.011-.002a.516.516 0 0 1-.29-.169l-.001-.001a2.999 2.999 0 0 1-.355-.609l-.008-.02q-.145-.339-.314-.651q-.363-.702-.702-1.427t-.651-1.452q-.217-.484-.399-.967a5.395 5.395 0 0 0-.461-.942l.013.023a17.366 17.366 0 0 0-1.331-1.961l.028.038a14.57 14.57 0 0 0-1.459-1.59l-.008-.007a17.619 17.619 0 0 0-1.632-1.356l-.049-.035q-.896-.651-1.96-1.282a3.7 3.7 0 0 0-.965-.393l-.026-.006l-1.113-.278l-.629-.048q-.314-.024-.629-.024a1.68 1.68 0 0 1-.387-.279a3.581 3.581 0 0 0-.353-.295l-.01-.007a11.686 11.686 0 0 0-2.043-.93L2.071.18A1.355 1.355 0 0 0 .9.096L.909.093a1.36 1.36 0 0 0-.795.84l-.003.01a1.515 1.515 0 0 0 .232 1.549l-.002-.003q.544.725.834 1.14q.217.291.448.605c.141.188.266.403.367.63l.008.021c.056.119.105.261.141.407l.003.016q.048.206.121.448q.217.556.411 1.14c.141.425.297.785.478 1.128l-.019-.04q.145.266.291.52T3.738 9a.868.868 0 0 0 .241.242l.003.002a.406.406 0 0 1 .169.313v.001a1.326 1.326 0 0 0-.217.586l-.001.006a4.227 4.227 0 0 1-.153.695l.008-.03a7.103 7.103 0 0 0-.351 2.231c0 .258.013.512.04.763l-.003-.031c.06.958.349 1.838.812 2.6l-.014-.025c.197.295.408.552.641.787a.914.914 0 0 0 1.106.203l-.005.002a.926.926 0 0 0 .617-.827v-.002c.048-.474.12-.898.219-1.312l-.013.067a.595.595 0 0 0 .038-.211l-.002-.045v.002q-.012-.109.133-.206v.048q.145.339.302.677t.326.677c.295.449.608.841.952 1.202l-.003-.003a7.74 7.74 0 0 0 1.127 1.001l.022.015c.212.162.398.337.566.528l.004.004c.158.186.347.339.56.454l.01.005v-.024h.048a.461.461 0 0 0-.18-.205l-.002-.001a1.829 1.829 0 0 1-.211-.136l.005.003q-.217-.217-.448-.484t-.423-.508q-.508-.702-.969-1.467t-.871-1.555q-.194-.387-.375-.798t-.351-.798a.997.997 0 0 1-.096-.334v-.005a.318.318 0 0 0-.168-.265l-.002-.001a2.947 2.947 0 0 1-.408.545l.001-.001a1.974 1.974 0 0 0-.382.58l-.005.013a4.272 4.272 0 0 0-.289 1.154l-.002.019q-.072.641-.145 1.318l-.048.024l-.024.024a.855.855 0 0 1-.59-.443l-.002-.005q-.182-.351-.326-.69a6.448 6.448 0 0 1-.423-2.144v-.009a6.218 6.218 0 0 1 .286-2.318l-.012.044q.072-.266.314-.896t.097-.871a.876.876 0 0 0-.265-.41l-.001-.001a3.31 3.31 0 0 1-.335-.335l-.003-.004q-.169-.244-.326-.52t-.278-.544a11.408 11.408 0 0 1-.474-1.353l-.022-.089a10.174 10.174 0 0 0-.546-1.503l.026.064a3.32 3.32 0 0 0-.39-.669l.006.008q-.244-.326-.436-.617q-.244-.314-.484-.605a3.414 3.414 0 0 1-.426-.657l-.009-.02a1.638 1.638 0 0 1-.119-.327l-.002-.011a.406.406 0 0 1 .049-.34l-.001.002a.303.303 0 0 1 .073-.145a.308.308 0 0 1 .143-.072h.002a.55.55 0 0 1 .536-.035l-.003-.001c.219.062.396.124.569.195l-.036-.013q.459.194.847.375c.298.142.552.292.792.459l-.018-.012q.194.121.387.266t.411.291h.339q.387 0 .822.037c.293.023.564.078.822.164l-.024-.007c.481.143.894.312 1.286.515l-.041-.019q.593.302 1.125.641c.589.367 1.098.743 1.577 1.154l-.017-.014c.5.428.954.867 1.38 1.331l.01.012c.416.454.813.947 1.176 1.464l.031.047c.334.472.671 1.018.974 1.584l.042.085a4.6 4.6 0 0 1 .234.536l.011.033q.097.278.217.57q.266.605.57 1.221t.57 1.198l.532 1.161c.187.406.396.756.639 1.079l-.011-.015c.203.217.474.369.778.422l.008.001c.368.092.678.196.978.319l-.047-.017c.143.065.327.134.516.195l.04.011c.212.065.396.151.565.259l-.009-.005c.327.183.604.363.868.559l-.021-.015q.411.302.822.57q.194.145.651.423t.484.52a11.202 11.202 0 0 0-1.834.087l.056-.006a5.959 5.959 0 0 0-1.479.39l.04-.014a2.556 2.556 0 0 1-.388.129l-.019.004a.312.312 0 0 0-.266.277v.001c.061.076.11.164.143.26l.002.006c.034.102.075.19.125.272l-.003-.006c.119.211.247.393.391.561l-.004-.005c.141.174.3.325.476.454l.007.005q.244.194.508.399c.161.126.343.25.532.362l.024.013c.284.174.614.34.958.479l.046.016c.374.15.695.324.993.531l-.016-.011q.291.169.58.375t.556.399c.073.072.137.152.191.239l.003.005a.573.573 0 0 0 .36.193h.003v-.048a.516.516 0 0 1-.184-.267l-.001-.004a.95.95 0 0 0-.112-.273l.002.004zM5.553 4.207q-.194 0-.363.012a1.285 1.285 0 0 0-.323.063l.009-.003v.024h.048q.097.145.244.326t.266.351l.387.798l.048-.024a.735.735 0 0 0 .252-.321l.002-.005c.052-.139.082-.301.082-.469l-.001-.054v.003a.504.504 0 0 1-.108-.154l-.001-.003l-.081-.182a.52.52 0 0 0-.214-.192l-.003-.001a.93.93 0 0 1-.244-.169"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NavIcon(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h27.65v5.219H0zm0 9.39h27.65v5.219H0zm0 9.391h27.65V24H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NavIconA(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.61 0h22.431a2.61 2.61 0 1 1 0 5.22H2.61a2.61 2.61 0 1 1 0-5.22m0 9.39h22.431a2.61 2.61 0 1 1 0 5.22H2.61a2.61 2.61 0 1 1 0-5.22m0 9.391h22.431a2.61 2.61 0 1 1 0 5.22H2.61a2.61 2.61 0 1 1 0-5.22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NavIconGrid(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h5.219v5.219H0zm9.39 0h5.219v5.219H9.39zm8.608 0h5.219v5.219h-5.219zM0 9.39h5.219v5.219H0zm9.39 0h5.219v5.219H9.39zm8.608 0h5.219v5.219h-5.219zM0 18.781h5.219V24H0zm9.39 0h5.219V24H9.39zm8.608 0h5.219V24h-5.219z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NavIconGridA(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 5.217a2.608 2.608 0 1 1 0-5.216a2.608 2.608 0 0 1 0 5.216m8.609 0a2.608 2.608 0 1 1 0-5.216a2.608 2.608 0 0 1 0 5.216m-18 0a2.608 2.608 0 1 1 0-5.216a2.608 2.608 0 0 1 0 5.216M12 14.608a2.608 2.608 0 1 1 0-5.216a2.608 2.608 0 0 1 0 5.216m8.609 0a2.608 2.608 0 1 1 0-5.216a2.608 2.608 0 0 1 0 5.216m-18 0a2.608 2.608 0 1 1 0-5.216a2.608 2.608 0 0 1 0 5.216M12 23.999a2.608 2.608 0 1 1 0-5.216A2.608 2.608 0 0 1 12 24zm8.609 0a2.608 2.608 0 1 1 0-5.216a2.608 2.608 0 0 1 0 5.216m-18 0a2.608 2.608 0 1 1 0-5.216a2.608 2.608 0 0 1 0 5.216"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NavIconList(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.184 0H27.65v5.219H7.184zm0 9.39H27.65v5.219H7.184zm0 9.391H27.65V24H7.184zM0 0h5.219v5.219H0zm0 9.39h5.219v5.219H0zm0 9.391h5.219V24H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NavIconListA(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.216 11.998a2.608 2.608 0 1 1-5.217 0a2.608 2.608 0 0 1 5.217 0m0-9.39a2.608 2.608 0 1 1-5.217 0a2.608 2.608 0 0 1 5.217 0m0 18.781a2.608 2.608 0 1 1-5.217 0a2.608 2.608 0 0 1 5.217 0M9.794 0h15.247a2.61 2.61 0 1 1 0 5.22H9.794a2.61 2.61 0 1 1 0-5.22m0 9.39h15.247a2.61 2.61 0 1 1 0 5.22H9.794a2.61 2.61 0 1 1 0-5.22m0 9.391h15.247a2.61 2.61 0 1 1 0 5.22H9.794a2.61 2.61 0 1 1 0-5.22"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Navigate(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23.968 0L0 10.286h13.68V24z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Nervous(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.316C6.099 1.316 1.316 6.099 1.316 12S6.099 22.684 12 22.684S22.684 17.901 22.684 12c-.012-5.896-4.788-10.672-10.683-10.684zm0 22.297C5.586 23.613.387 18.414.387 12S5.586.387 12 .387S23.613 5.586 23.613 12v.015c0 6.405-5.192 11.597-11.597 11.597z"/><path fill="currentColor" d="M12 24C5.386 23.966.034 18.614 0 12.003V12C0 5.373 5.373 0 12 0s12 5.373 12 12c-.034 6.614-5.386 11.966-11.997 12zM12 .774C5.8.774.774 5.8.774 12S5.8 23.226 12 23.226S23.226 18.2 23.226 12C23.222 5.802 18.198.779 12.001.774zm0 22.297C5.886 23.071.929 18.114.929 12S5.886.929 12 .929S23.071 5.886 23.071 12S18.114 23.071 12 23.071m0-21.368C6.313 1.703 1.703 6.313 1.703 12S6.313 22.297 12 22.297S22.297 17.687 22.297 12v-.005c0-5.684-4.608-10.292-10.292-10.292z"/><path fill="currentColor" d="M9.677 9.91a2.09 2.09 0 1 1-2.09-2.09h.009c1.15 0 2.082.932 2.082 2.082v.009zm8.904 0a2.09 2.09 0 1 1-2.09-2.09h.009c1.15 0 2.082.932 2.082 2.082v.009zm-3.871 6.967c0 2.013-1.161 3.639-2.71 3.639c-1.471 0-2.71-1.626-2.71-3.639s1.161-3.639 2.71-3.639c1.471 0 2.71 1.626 2.71 3.639M9.987 7.819a.279.279 0 0 1-.232-.078L7.2 6.193a1.05 1.05 0 0 1-.229-.304l-.003-.006a.462.462 0 0 1 .078-.311l-.001.002a1.05 1.05 0 0 1 .304-.229l.006-.003a.462.462 0 0 1 .311.078l-.002-.001l2.555 1.548c.094.087.172.189.229.304l.003.006a.462.462 0 0 1-.078.311l.001-.002a.426.426 0 0 1-.379.232h-.008z"/><path fill="currentColor" d="M9.987 8.206a1.085 1.085 0 0 1-.469-.158l.005.003l-2.555-1.548a.987.987 0 0 1-.385-.535l-.002-.007a.815.815 0 0 1 .079-.623l-.002.004a.987.987 0 0 1 .535-.385l.007-.002a.815.815 0 0 1 .623.079l-.004-.002l2.555 1.548a.987.987 0 0 1 .385.535l.002.007a.815.815 0 0 1-.079.623l.002-.004a.746.746 0 0 1-.69.465zm-2.477-2.4h-.077v.077l2.555 1.548l.077-.077zm6.503 2.013h-.008a.424.424 0 0 1-.378-.23l-.001-.002a.46.46 0 0 1-.077-.312v.002a.417.417 0 0 1 .23-.309l.002-.001l2.555-1.548a.46.46 0 0 1 .312-.077h-.002c.138.023.251.11.309.23l.001.002a.46.46 0 0 1 .077.312v-.002a.417.417 0 0 1-.23.309l-.002.001l-2.555 1.548a.285.285 0 0 1-.233.078h.001z"/><path fill="currentColor" d="M14.013 8.206h-.004a.81.81 0 0 1-.69-.384l-.002-.003a.923.923 0 0 1-.154-.623V7.2a.899.899 0 0 1 .384-.54l.003-.002l2.555-1.548a.81.81 0 0 1 .625-.076l-.006-.001a.899.899 0 0 1 .54.384l.002.003a.81.81 0 0 1 .076.625l.001-.006a.899.899 0 0 1-.384.54l-.003.002l-2.555 1.548a.57.57 0 0 1-.39.077zm-.078-.774l2.632-1.471l.387-.31l-.31.232l-.31-.387l.232.31l-2.555 1.548l-.387.31z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Netflix(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 0l4.751 13.462v-.007l.376 1.06c2.088 5.909 3.211 9.077 3.217 9.084s.32.02.7.04c1.156.05 2.59.18 3.677.31c.122.02.262.031.404.031l.08-.001h-.004L8.49 10.617l-.436-1.23l-2.423-6.851c-.46-1.3-.85-2.408-.87-2.45L4.732.001H.006z"/><path fill="currentColor" d="m8.511.012l-.01 5.307l-.008 5.307l-.437-1.232l-.567 11.81c.555 1.567.852 2.4.855 2.407s.32.024.7.042c1.157.05 2.59.18 3.681.31c.122.02.262.031.405.031l.079-.001h-.004c.013-.01.02-5.421.017-12.012L13.212.01H8.511zM0 0v11.992c0 6.595.007 11.997.015 12.002s.416-.03.907-.086s1.17-.13 1.51-.16c.518-.05 2.068-.15 2.248-.15c.052 0 .056-.27.063-5.081l.008-5.081l.38 1.06l.13.376l.57-11.802l-.19-.546L4.731 0H-.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Neuter(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.429 7.714l.001.1a7.414 7.414 0 0 1-1.98 5.054l.004-.005a7.476 7.476 0 0 1-4.848 2.509l-.032.003v8.212a.413.413 0 0 1-.413.413h-.017h.001h-.873a.413.413 0 0 1-.413-.413v-.017v.001v-8.196a7.482 7.482 0 0 1-4.873-2.503l-.007-.008A7.41 7.41 0 0 1-.001 7.81L0 7.709v-.018c0-1.073.224-2.094.629-3.019l-.019.049A7.649 7.649 0 0 1 4.671.629l.05-.019C5.604.225 6.633.001 7.714.001s2.11.224 3.043.628l-.05-.019a7.65 7.65 0 0 1 4.092 4.062l.019.05c.386.877.611 1.899.611 2.974v.02zm-7.715 6l.087.001a5.763 5.763 0 0 0 4.152-1.759l.001-.002a5.765 5.765 0 0 0 1.761-4.153l-.001-.092v.005l.001-.087a5.763 5.763 0 0 0-1.759-4.152l-.002-.001c-1.083-1.089-2.583-1.762-4.24-1.762s-3.157.674-4.24 1.762a5.765 5.765 0 0 0-1.76 4.245v-.005l-.001.087c0 1.629.674 3.101 1.759 4.152l.002.001a5.765 5.765 0 0 0 4.153 1.761l.092-.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Neutral(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.316C6.099 1.316 1.316 6.099 1.316 12S6.099 22.684 12 22.684S22.684 17.901 22.684 12c-.012-5.896-4.788-10.672-10.683-10.684zm0 22.297C5.586 23.613.387 18.414.387 12S5.586.387 12 .387S23.613 5.586 23.613 12v.015c0 6.405-5.192 11.597-11.597 11.597z"/><path fill="currentColor" d="M12 24C5.386 23.966.034 18.614 0 12.003V12C0 5.373 5.373 0 12 0s12 5.373 12 12c-.034 6.614-5.386 11.966-11.997 12zM12 .774C5.8.774.774 5.8.774 12S5.8 23.226 12 23.226S23.226 18.2 23.226 12C23.222 5.802 18.198.779 12.001.774zm0 22.297C5.886 23.071.929 18.114.929 12S5.886.929 12 .929S23.071 5.886 23.071 12S18.114 23.071 12 23.071m0-21.368C6.313 1.703 1.703 6.313 1.703 12S6.313 22.297 12 22.297S22.297 17.687 22.297 12v-.005c0-5.684-4.608-10.292-10.292-10.292z"/><path fill="currentColor" d="M9.677 9.91v.009c0 1.15-.932 2.082-2.082 2.082h-.009a2.09 2.09 0 0 1 0-4.18h.009c1.15 0 2.082.932 2.082 2.082v.009zm8.904 0a2.09 2.09 0 0 1-4.18 0v-.009c0-1.15.932-2.082 2.082-2.082h.009a2.132 2.132 0 0 1 2.09 2.088v.002zm-2.168 7.277H7.51a.465.465 0 1 1 0-.93h8.903a.499.499 0 0 1 .464.463v.021c0 .246-.2.446-.446.446h-.02h.001z"/><path fill="currentColor" d="M16.413 17.574H7.51a.852.852 0 0 1 0-1.704h8.903a.852.852 0 0 1 0 1.704m-8.826-.929c-.077 0-.077.077 0 0c-.02.02-.033.047-.033.077s.013.058.033.077h8.903a.077.077 0 0 0 .077-.077c0-.077 0-.077-.077-.077z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Nginx(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.395 0L0 6v12l10.395 6l10.4-6V6zm6 16.59a1.407 1.407 0 0 1-1.535 1.29h.006h-.026a2.342 2.342 0 0 1-1.771-.807l-.002-.003l-6-7.141v6.674c0 .703-.568 1.273-1.271 1.276h-.08a1.3 1.3 0 0 1-1.29-1.289V7.41a1.387 1.387 0 0 1 1.505-1.29h-.005h.039c.712 0 1.352.312 1.789.808l.002.003l5.97 7.141V7.411a1.3 1.3 0 0 1 1.289-1.29h.076a1.3 1.3 0 0 1 1.29 1.289v9.181z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightAltCloudy(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.57 10.402h.012c1.121 0 2.15.392 2.957 1.047l-.009-.007a.91.91 0 0 0 1.279-.139l.001-.002a.907.907 0 0 0-.14-1.279l-.002-.001a6.505 6.505 0 0 0-1.842-1.035l-.046-.015c1.285-1.881 3.417-3.101 5.834-3.108h.001a7.04 7.04 0 0 1 7.034 6.984v.003c-.82.145-1.54.495-2.129.995l.006-.005a.91.91 0 0 0-.105 1.282l-.001-.001a.91.91 0 0 0 1.282.105l-.001.001a2.628 2.628 0 0 1 1.707-.625a2.65 2.65 0 0 1 2.646 2.645a2.648 2.648 0 0 1-2.645 2.646H6.569a4.752 4.752 0 0 1-4.746-4.747a4.75 4.75 0 0 1 4.745-4.744zm17.892-8.503a6.497 6.497 0 0 0-.212.971l-.004.037a6.95 6.95 0 0 0 5.871 7.83l.038.004c.291.043.626.068.967.068h.062h-.003c-.985 2.036-2.974 3.447-5.306 3.593l-.018.001a4.477 4.477 0 0 0-2.349-1.483l-.031-.007v-.017a8.84 8.84 0 0 0-3.752-7.233l-.028-.019c.847-1.959 2.608-3.381 4.729-3.739zM6.57 21.714h15.84a4.47 4.47 0 0 0 4.466-4.467a4.491 4.491 0 0 0-.153-1.144l.007.031a8.17 8.17 0 0 0 6.646-6.117l.011-.056a.91.91 0 0 0-.219-.826l.001.001a.907.907 0 0 0-.811-.281l.005-.001l-.362.056a5.17 5.17 0 0 1-1.627.024l.028.003c-2.505-.362-4.408-2.494-4.408-5.07c0-.25.018-.495.052-.736l-.003.027a5.126 5.126 0 0 1 .597-1.786l-.013.026a.887.887 0 0 0 .136-.454V.918a.909.909 0 0 0-.86-.908h-.002a8.112 8.112 0 0 0-7.782 4.695l-.021.05a8.656 8.656 0 0 0-3.475-.715h-.009a8.893 8.893 0 0 0-7.721 4.511l-.023.044c-.102-.005-.203-.013-.305-.013a6.573 6.573 0 0 0-6.566 6.566a6.575 6.575 0 0 0 6.566 6.568z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightAltLightning(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.317 23.823a.62.62 0 0 1 0-.878l.761-.761H7.83a.657.657 0 0 1-.126-.013l.004.001c-.019-.003-.036-.011-.055-.016a.44.44 0 0 1-.064-.02l.003.001a.56.56 0 0 1-.065-.034l.003.002a.328.328 0 0 1-.046-.026l.002.001a.633.633 0 0 1-.172-.17l-.001-.002c-.01-.015-.017-.032-.026-.047s-.02-.035-.029-.055l-.002-.004c-.008-.021-.014-.043-.02-.064c-.005-.015-.011-.033-.015-.052l-.001-.003a.587.587 0 0 1 .001-.248l-.001.004c.004-.019.01-.035.016-.055s.011-.043.02-.063a.508.508 0 0 1 .033-.062l-.002.003c.009-.016.016-.033.026-.048a.713.713 0 0 1 .076-.093l1.746-1.746a.62.62 0 1 1 .878.878l-.687.687h.249c.042 0 .084.004.124.013l-.004-.001a.362.362 0 0 1 .058.018l-.004-.001a.483.483 0 0 1 .064.02l-.004-.001a.576.576 0 0 1 .066.035l-.003-.002c.015.008.029.014.044.023a.629.629 0 0 1 .171.17l.001.002c.01.015.017.031.026.047c.01.017.021.036.03.057l.001.003a.693.693 0 0 1 .019.059l.001.005l.015.051l.001.003a.552.552 0 0 1 0 .245l.001-.004c-.004.019-.011.038-.016.055a.426.426 0 0 1-.02.063l.001-.004a.453.453 0 0 1-.035.066l.001-.002c-.008.015-.014.029-.023.044a.68.68 0 0 1-.078.096l-1.819 1.819a.62.62 0 0 1-.878 0zm-2.474-3.29a.62.62 0 0 1 0-.878l1.95-1.95h-1.39a.638.638 0 0 1-.125-.013l.004.001c-.019-.004-.036-.011-.055-.016a.496.496 0 0 1-.065-.021l.003.001a.726.726 0 0 1-.063-.033l.003.002a.546.546 0 0 1-.049-.026l.002.001a.615.615 0 0 1-.171-.17l-.001-.002c-.01-.014-.016-.03-.025-.045a.37.37 0 0 1-.054-.121l-.001-.003a.634.634 0 0 1-.016-.051l-.001-.004a.577.577 0 0 1 .001-.246l-.001.004a.397.397 0 0 1 .018-.058l-.001.004a.634.634 0 0 1 .02-.065l-.001.004a.44.44 0 0 1 .035-.065l-.001.002l.025-.046l-.001.002a.68.68 0 0 1 .078-.096l1.828-1.828h-2.31a4.48 4.48 0 0 1 0-8.96c.07 0 .139.006.208.009c1.057-1.863 3.024-3.101 5.282-3.107h.007c.857 0 1.672.179 2.409.502l-.039-.015C13.233 1.316 15.15 0 17.374 0c.104 0 .207.003.31.009L17.67.008a.621.621 0 0 1 .588.62v.018v-.001a.619.619 0 0 1-.094.314l.001-.003a3.493 3.493 0 0 0 3.686 5.123l-.021.003l.248-.038a.621.621 0 0 1 .697.759l.001-.004a5.577 5.577 0 0 1-4.51 4.207l-.032.004a3.035 3.035 0 0 1-2.936 3.806H13.31l-.504.504h.248c.043 0 .085.005.126.013l-.004-.001a.362.362 0 0 1 .058.018l-.004-.001a.624.624 0 0 1 .065.021l-.004-.001a.44.44 0 0 1 .065.035l-.002-.001c.015.008.03.015.044.024a.629.629 0 0 1 .171.17l.001.002a.362.362 0 0 1 .024.044l.001.003c.01.016.021.036.03.056l.002.004a.693.693 0 0 1 .019.059l.001.004l.015.051l.001.003a.552.552 0 0 1 0 .245l.001-.004a.397.397 0 0 1-.018.058l.001-.003c-.005.017-.01.04-.018.059a.793.793 0 0 1-.035.066l.002-.003c-.008.015-.014.029-.023.044a.57.57 0 0 1-.079.095l-1.82 1.82a.62.62 0 1 1-.878-.878l.76-.76h-.248a.638.638 0 0 1-.125-.013l.004.001c-.019-.003-.037-.011-.055-.016a.452.452 0 0 1-.065-.021l.004.001a.56.56 0 0 1-.065-.034l.003.002c-.015-.008-.03-.015-.045-.025a.629.629 0 0 1-.171-.17l-.001-.002c-.01-.015-.017-.032-.026-.047a.711.711 0 0 1-.03-.056l-.002-.004a.833.833 0 0 1-.019-.059l-.001-.004l-.015-.051l-.001-.003a.636.636 0 0 1 0-.246l-.001.004c.004-.019.011-.037.016-.055s.011-.04.019-.06a.66.66 0 0 1 .035-.066l-.002.003c.008-.014.015-.03.024-.044a.68.68 0 0 1 .078-.096l.686-.686H8.547L6.9 16.462h1.391c.043 0 .085.004.126.013l-.004-.001c.019.004.036.011.055.016a.467.467 0 0 1 .066.021l-.004-.001a.726.726 0 0 1 .063.033l-.003-.002c.015.009.032.015.046.026a.615.615 0 0 1 .171.17l.001.002c.009.014.016.029.024.044a.539.539 0 0 1 .032.059l.001.004a.53.53 0 0 1 .018.056l.001.004c.005.019.013.036.016.055a.577.577 0 0 1-.001.246l.001-.004c-.004.019-.011.037-.016.055a.634.634 0 0 1-.02.065l.001-.004a.44.44 0 0 1-.035.065l.001-.002a.395.395 0 0 1-.025.046l.001-.002a.646.646 0 0 1-.078.098l-3.01 3.01a.62.62 0 0 1-.878 0zM1.241 10.335a3.241 3.241 0 0 0 3.237 3.238h10.811a1.804 1.804 0 1 0-.001-3.61c-.445 0-.852.161-1.167.428l.003-.002a.621.621 0 0 1-.803-.948l.001-.001a3.022 3.022 0 0 1 1.429-.672l.018-.003A4.798 4.798 0 0 0 6 6.101l-.011.017c.493.179.92.421 1.296.723l-.009-.007a.622.622 0 0 1-.777.967l.002.001a3.184 3.184 0 0 0-2.011-.709h-.009a3.241 3.241 0 0 0-3.24 3.241zm12.193-6.481a6.04 6.04 0 0 1 2.579 4.947v.012a3.064 3.064 0 0 1 1.62 1.011l.004.005a4.319 4.319 0 0 0 3.621-2.426l.011-.026h-.037a4.74 4.74 0 0 1-4.539-6.111l-.009.034a4.302 4.302 0 0 0-3.24 2.527l-.011.028z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightAltRain(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.241 23.795a.698.698 0 0 1 0-.99l3.626-3.625a.7.7 0 1 1 .99.991l-3.626 3.625a.698.698 0 0 1-.99 0zm7.376-3.43a.698.698 0 0 1 0-.99l2.663-2.663H6.881l-3.653 3.653a.701.701 0 0 1-.992-.991l2.677-2.677A5.049 5.049 0 0 1 5.048 6.6l.143.002h-.007c.079 0 .156.006.235.01C6.61 4.51 8.83 3.113 11.377 3.107h.007c.967 0 1.886.202 2.718.567l-.044-.017C15.064 1.482 17.227 0 19.736 0c.115 0 .229.003.342.009l-.016-.001c.37.02.663.326.663.7v.019v-.001a.7.7 0 0 1-.106.354l.002-.003a3.94 3.94 0 0 0 4.158 5.779l-.024.004l.279-.042a.7.7 0 0 1 .786.856l.001-.005c-.598 2.467-2.596 4.34-5.087 4.745l-.036.005a3.423 3.423 0 0 1-3.31 4.294H16.26l-3.653 3.654a.698.698 0 0 1-.99 0zM1.534 11.657a3.657 3.657 0 0 0 3.651 3.653h12.192a2.036 2.036 0 1 0-1.317-3.589l.003-.002a.701.701 0 0 1-.907-1.069l.001-.001a3.406 3.406 0 0 1 1.612-.758l.021-.003a5.411 5.411 0 0 0-9.891-3.003l-.012.019a5.052 5.052 0 0 1 1.462.815l-.01-.008a.701.701 0 1 1-.876 1.091l.001.001a3.594 3.594 0 0 0-2.268-.8h-.01h.001a3.656 3.656 0 0 0-3.654 3.655zm13.752-7.311c1.768 1.247 2.908 3.28 2.909 5.58v.013c.74.185 1.367.59 1.826 1.14l.004.006a4.87 4.87 0 0 0 4.084-2.736l.013-.029h-.051c-.261 0-.517-.019-.767-.056l.028.003c-2.613-.38-4.598-2.604-4.598-5.293c0-.538.079-1.058.227-1.547l-.01.038a4.849 4.849 0 0 0-3.654 2.85z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightAltSnow(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.651 23.354a.646.646 0 1 1 1.291 0a.646.646 0 0 1-1.291 0M8.5 21.505a.646.646 0 1 1 1.291 0a.646.646 0 0 1-1.291 0m5.218-.843a.646.646 0 1 1 1.291 0a.646.646 0 0 1-1.291 0m-9.94 0a.646.646 0 1 1 .646.646h-.001a.645.645 0 0 1-.645-.646m6.57-1.005a.646.646 0 1 1 1.291 0a.646.646 0 0 1-1.291 0m5.218-.843a.646.646 0 1 1 1.291 0a.646.646 0 0 1-1.291 0m-9.938 0a.646.646 0 1 1 1.291 0a.646.646 0 0 1-1.291 0m-.494-1.834h-.057a5.135 5.135 0 0 1 0-10.268h.06h-.003c.08 0 .159.006.238.01c1.206-2.14 3.463-3.561 6.053-3.561a6.91 6.91 0 0 1 2.77.576l-.045-.017C15.167 1.508 17.364 0 19.913 0c.118 0 .235.003.352.01l-.016-.001a.712.712 0 0 1 .581 1.063l.002-.003a4.003 4.003 0 0 0 4.209 5.898l-.024.004l.283-.043a.712.712 0 0 1 .8.869l.001-.005c-.608 2.506-2.638 4.409-5.168 4.821l-.037.005a3.479 3.479 0 0 1-3.365 4.361h-.012h.001zm-3.71-5.135a3.715 3.715 0 0 0 3.71 3.711h12.384a2.068 2.068 0 1 0-1.338-3.645l.003-.002a.712.712 0 1 1-.919-1.086l.002-.001a3.482 3.482 0 0 1 1.639-.771l.021-.003A5.5 5.5 0 0 0 6.877 6.997l-.012.019a5.163 5.163 0 0 1 1.485.829l-.01-.008a.713.713 0 1 1-.891 1.109l.001.001a3.653 3.653 0 0 0-2.305-.813h-.01h.001a3.715 3.715 0 0 0-3.713 3.712zm13.971-7.428c1.796 1.267 2.954 3.333 2.955 5.669v.013c.752.188 1.389.6 1.855 1.159l.005.006a4.947 4.947 0 0 0 4.149-2.78l.013-.029h-.044a5.43 5.43 0 0 1-5.2-7.004l-.01.038a4.93 4.93 0 0 0-3.712 2.896z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NightClear(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.524 14.721h.008c.644 0 1.275-.059 1.886-.172l-.063.01c-1.146 4.122-4.866 7.098-9.281 7.098h-.058h.003c-5.343-.006-9.673-4.336-9.678-9.679v-.001A9.76 9.76 0 0 1 9.37 2.665l.069-.017a10.013 10.013 0 0 0-.162 1.819v.007c.005 5.658 4.59 10.243 10.247 10.248h.001zM12.006.47a1.162 1.162 0 0 0-1.043-.465h.005C4.813.596.034 5.724 0 11.976v.003C.008 18.614 5.385 23.991 12.019 24h.061c6.243 0 11.367-4.786 11.905-10.889l.003-.045a1.168 1.168 0 0 0-.423-1.009l-.002-.002a1.164 1.164 0 0 0-1.084-.213l.008-.002l-.524.156a7.735 7.735 0 0 1-2.435.385h-.007a7.912 7.912 0 0 1-7.903-7.903V4.46c0-1.03.198-2.014.558-2.915l-.019.053a1.169 1.169 0 0 0-.155-1.134l.002.003z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Nodejs(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.639 24h-.003c-.339 0-.656-.091-.928-.251l.009.005l-2.937-1.737c-.438-.246-.224-.332-.08-.383a5.817 5.817 0 0 0 1.352-.618l-.024.014a.225.225 0 0 1 .219.018l-.001-.001l2.256 1.339a.288.288 0 0 0 .274-.001l-.002.001l8.794-5.077a.28.28 0 0 0 .134-.238V6.922a.282.282 0 0 0-.136-.239l-.001-.001l-8.791-5.072a.277.277 0 0 0-.272.001l.001-.001l-8.789 5.073a.281.281 0 0 0-.139.24v10.149c0 .101.055.188.137.234l.001.001l2.41 1.392c1.307.654 2.107-.116 2.107-.889V7.788c0-.14.114-.254.254-.254h1.119c.14 0 .254.113.254.254V17.81c0 1.745-.95 2.746-2.604 2.746l-.066.001a3.745 3.745 0 0 1-1.976-.56l-2.29-1.318a1.858 1.858 0 0 1-.922-1.605V6.923C0 6.243.367 5.65.913 5.328l.009-.005L9.717.241a1.923 1.923 0 0 1 1.858.005l-.01-.005l8.794 5.082c.555.327.921.92.923 1.6v10.15a1.86 1.86 0 0 1-.915 1.6l-.009.005l-8.792 5.078a1.823 1.823 0 0 1-.922.246h-.007z"/><path fill="currentColor" d="M13.356 17.009c-3.848 0-4.655-1.766-4.655-3.249v-.001c0-.14.113-.253.253-.253h1.14c.127 0 .232.093.252.214v.001c.171 1.158.683 1.742 3.01 1.742c1.853 0 2.64-.419 2.64-1.402c0-.567-.223-.987-3.102-1.269c-2.406-.24-3.894-.77-3.894-2.695c0-1.774 1.496-2.833 4-2.833c2.818 0 4.212.978 4.388 3.076l.001.022a.254.254 0 0 1-.254.254h-1.133a.253.253 0 0 1-.246-.198v-.002c-.274-1.218-.94-1.607-2.746-1.607c-2.023 0-2.258.705-2.258 1.232c0 .64.278.826 3.009 1.187c2.702.358 3.986.863 3.986 2.762c-.004 1.919-1.601 3.017-4.388 3.017z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Npm(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 12.534V9.871h-1.334v2.666zM24 7.2H0v8h6.666v1.334H12v-1.333h12zM6.666 8.534v5.337H5.333v-4H3.999v4H1.333V8.537zm6.667 0v5.337h-2.666v1.334H8.001V8.539zm9.333 0v5.337h-1.333v-4h-1.334v4h-1.334v-4h-1.333v4h-2.667V8.537z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Nurse(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.534 21.152a10.545 10.545 0 0 0-.51-3.901l.021.075c-.058-.172-.104-1.326-2.014-2.072c-1.406-.549-3.291-.454-4.688-1.406a2.203 2.203 0 0 1-.078-.311l-.002-.015a12.085 12.085 0 0 0 4.006-1.533l-.054.03c-2.012-.572-2.16-3.784-2.301-5.76c0-.558-.111-1.09-.312-1.575l.01.027a4.17 4.17 0 0 1-.03-.118L14.758.001H4.796l1.172 4.595c-.01.033-.018.07-.03.117a4.043 4.043 0 0 0-.302 1.547c-.142 2.022-.302 5.188-2.301 5.76a11.887 11.887 0 0 0 3.873 1.49l.076.012c-.024.13-.052.239-.086.346l.006-.02c-1.396.952-3.28.855-4.688 1.406C.607 16 .562 17.154.504 17.326c-.321.962-.506 2.07-.506 3.221c0 .213.006.424.019.634l-.001-.029c-.01.368 0 .923 1.101 1.406c3.261 1.429 8.64 1.44 8.661 1.44a27.858 27.858 0 0 0 8.857-1.498l-.196.058c1.098-.483 1.105-1.038 1.098-1.409zM13.807.738l-.975 3.834H6.709L5.734.738z"/><path fill="currentColor" d="M8.234 3.063h1.077v1.079h.919V3.063h1.079v-.917h-1.083V1.071h-.919V2.15H8.234z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func NursingHome(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m0 8.853l2.886 10.115c2.738-.403 5.899-.633 9.113-.633s6.375.23 9.467.675l-.353-.042l2.886-10.115C14.497 4.628 8.858 4.405 0 8.853m14.918 4.276h-2.071V15.2h-1.686v-2.071H9.09v-1.686h2.071V9.371h1.686v2.072h2.071z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Odnoklassniki(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.456 17.441a11.592 11.592 0 0 0 3.647-1.518l-.047.028a1.817 1.817 0 0 0-1.943-3.069l.008-.005c-1.184.751-2.625 1.196-4.17 1.196s-2.987-.446-4.202-1.215l.032.019a1.815 1.815 0 0 0-2.496.558l-.004.007v.005a1.812 1.812 0 0 0 .56 2.496l.007.004A11.354 11.354 0 0 0 4.37 17.42l.078.015L.979 20.9a1.796 1.796 0 0 0-.538 1.284c0 .489.195.932.51 1.256l.03.03c.326.327.777.529 1.275.529s.949-.202 1.275-.529l3.421-3.404l3.4 3.406a1.814 1.814 0 1 0 2.566-2.565zm-2.505-5.052a6.194 6.194 0 1 0-6.2-6.194v.006a6.2 6.2 0 0 0 6.2 6.188m0-8.758A2.565 2.565 0 1 1 4.386 6.2v-.002a2.567 2.567 0 0 1 2.565-2.567"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Onedrive(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.92 17.427a3.122 3.122 0 0 1-1.636.457h-.046h.002a3.254 3.254 0 0 1-2.407-5.428l-.003.003a3.18 3.18 0 0 1 2.044-1.045l.014-.001a3.261 3.261 0 0 1-.042-.532v-.055c0-1.095.447-2.086 1.169-2.8a4.061 4.061 0 0 1 2.842-1.155l.07-.001c.687 0 1.33.19 1.88.519L8.79 7.38a5.682 5.682 0 0 1 1.837-1.845l.024-.014a5.057 5.057 0 0 1 2.615-.72c1.264 0 2.421.459 3.312 1.22l-.007-.006a5.076 5.076 0 0 1 1.751 2.999l.005.031h-.285l-.079-.001a3.49 3.49 0 0 0-1.175.202l.024-.008a4.875 4.875 0 0 0-3.579-1.56l-.101.001h.005h-.047c-.662 0-1.295.121-1.879.342l.037-.012a5.364 5.364 0 0 0-2.657 2.167l-.013.022a4.43 4.43 0 0 0-.574 1.459l-.005.029a5.42 5.42 0 0 0-1.056.327L6.978 12a3.54 3.54 0 0 0-1.283.998l-.005.007c-.324.366-.58.8-.742 1.278l-.008.026a4.755 4.755 0 0 0-.256 1.544v.001l-.001.078c0 .571.111 1.117.312 1.617l-.01-.029l-.062-.091zm16.847-3.747q2.364.586 2.225 2.781t-2.552 2.472H8.7q-3.156-.416-3.1-3.263c.034-1.905 1.104-2.954 3.21-3.135a4.095 4.095 0 0 1 3.092-3.774l.028-.006a4.1 4.1 0 0 1 4.602 1.665l.009.015a2.234 2.234 0 0 1 1.661-.465l-.011-.001a4.03 4.03 0 0 1 1.719.447l-.023-.011a3.347 3.347 0 0 1 1.387 1.32l.009.016a3.89 3.89 0 0 1 .48 1.884v.038v-.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Onenote(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M25.35 10.078v3.803a.541.541 0 0 1-.535.539h-.825V9.538h.831c.146 0 .278.06.372.158a.529.529 0 0 1 .158.377v.005zm-.537 4.905a.54.54 0 0 1 .536.535v3.824a.53.53 0 0 1-.529.529h-.831v-4.892h.825zm0-10.893a.618.618 0 0 1 .407.159h-.001a.49.49 0 0 1 .186.381v3.814a.484.484 0 0 1-.185.38l-.001.001a.645.645 0 0 1-.406.143h-1.365v11.735a.445.445 0 0 1-.165.345l-.001.001a.644.644 0 0 1-.385.127h-7.092v-2.4h5.512v-1.142h-5.516v-1.348h5.512v-1.081h-5.512v-1.381h5.512V12.76h-5.512v-1.349h5.512v-1.065h-5.512V8.965h5.512v-1.08h-5.512v-1.35h5.512V5.408h-5.512V2.727h7.29c.07.038.129.085.178.142l.001.001c.097.094.162.22.18.361v.861h1.363zM14.445 0v24L0 21.506V2.559zm-4.121 7.114l-1.694.111v3.153q-.029 2.77 0 3.198l-3.378-6.16l-1.724.068v8.74l1.349.106V9.883l3.642 6.632l1.8.106V7.114z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func OpenMouth(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.316C6.099 1.316 1.316 6.099 1.316 12S6.099 22.684 12 22.684S22.684 17.901 22.684 12c-.012-5.896-4.788-10.672-10.683-10.684zm0 22.297C5.586 23.613.387 18.414.387 12S5.586.387 12 .387S23.613 5.586 23.613 12v.015c0 6.405-5.192 11.597-11.597 11.597z"/><path fill="currentColor" d="M12 24C5.386 23.966.034 18.614 0 12.003V12C0 5.373 5.373 0 12 0s12 5.373 12 12c-.034 6.614-5.386 11.966-11.997 12zM12 .774C5.8.774.774 5.8.774 12S5.8 23.226 12 23.226S23.226 18.2 23.226 12C23.222 5.802 18.198.779 12.001.774zm0 22.297C5.886 23.071.929 18.114.929 12S5.886.929 12 .929c6.122.02 11.084 4.955 11.148 11.065V12c-.064 6.115-5.027 11.051-11.146 11.071zm0-21.368C6.313 1.703 1.703 6.313 1.703 12S6.313 22.297 12 22.297S22.297 17.687 22.297 12c-.026-5.676-4.62-10.271-10.294-10.297z"/><path fill="currentColor" d="M9.677 9.91v.009c0 1.15-.932 2.082-2.082 2.082h-.018a2.082 2.082 0 0 1-2.082-2.082V9.91a2.09 2.09 0 0 1 4.18 0zm8.904 0v.009c0 1.15-.932 2.082-2.082 2.082h-.018a2.082 2.082 0 0 1-2.082-2.082v-.018c0-1.15.932-2.082 2.082-2.082h.009a2.132 2.132 0 0 1 2.09 2.088v.002zm-3.871 6.967c0 2.013-1.161 3.639-2.71 3.639s-2.71-1.626-2.71-3.639s1.161-3.639 2.71-3.639s2.71 1.626 2.71 3.639"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Opencart(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.06 21.626v.002A2.372 2.372 0 0 1 19.688 24h-.035A2.307 2.307 0 0 1 18 23.305a2.271 2.271 0 0 1-.703-1.645v-.037v.002a2.4 2.4 0 0 1 2.388-2.389h.036c.646 0 1.23.269 1.644.702l.001.001c.428.42.694 1.005.694 1.651v.039v-.002zm-10.842 0v.035c0 .646-.269 1.23-.702 1.644l-.001.001A2.307 2.307 0 0 1 8.864 24h-.037h.002h-.002a2.372 2.372 0 0 1-2.372-2.372v-.035c0-.647.266-1.233.695-1.653a2.271 2.271 0 0 1 1.645-.703h.037h-.002a2.4 2.4 0 0 1 2.389 2.388zM0 0a18.999 18.999 0 0 0 1.572 1.49l.028.023c.513.426 1.086.819 1.694 1.158l.058.03c.549.314 1.197.621 1.87.879l.099.033a15.37 15.37 0 0 0 2.295.641l.105.018q1.411.29 2.895.456t3.619.268t4.4.138t5.392.036q2.012 0 3.539.072c.968.04 1.872.126 2.761.257l-.141-.017c.683.096 1.29.236 1.875.422l-.08-.022c.4.113.746.31 1.03.573l-.002-.001c.197.191.326.452.347.742V7.2a1.704 1.704 0 0 1-.286.933l.004-.006c-.253.418-.523.78-.823 1.114l.006-.007q-.514.586-1.295 1.318t-1.68 1.513t-2.012 1.722q-2.678 2.273-4.14 3.575c.353-.599.72-1.116 1.126-1.598l-.014.018q.687-.839 1.36-1.527t1.368-1.426t1.2-1.325c.3-.338.561-.721.768-1.135l.014-.03a1.616 1.616 0 0 0 .187-1.022l.001.009a1.338 1.338 0 0 0-.651-.797l-.007-.003a4.932 4.932 0 0 0-1.657-.59l-.029-.004a21.734 21.734 0 0 0-2.898-.338l-.055-.002q-1.84-.116-4.4-.072q-2.432.029-4.546-.087T11.191 7.1a22.586 22.586 0 0 1-3.12-.632l.16.038a22.877 22.877 0 0 1-2.469-.804l.16.055a10.003 10.003 0 0 1-1.813-.929l.039.024a11.761 11.761 0 0 1-1.343-.975l.018.015a8.314 8.314 0 0 1-.969-1.02l-.013-.016a9.256 9.256 0 0 1-.709-.969l-.022-.038q-.24-.4-.579-.984T.001.003z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Opera(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.995 3.054a8.42 8.42 0 0 0-4.782-1.474h-.028h.001h-.018c-1.43 0-2.775.363-3.949 1.002l.043-.022a10.318 10.318 0 0 0-3.199 2.658l-.016.021A9.815 9.815 0 0 0 6.468 8.09l-.021.069a11.248 11.248 0 0 0-.649 3.549v.574c.03 1.284.266 2.503.675 3.64l-.025-.08a9.902 9.902 0 0 0 1.615 2.94l-.015-.02a10.301 10.301 0 0 0 3.16 2.653l.054.027a8.165 8.165 0 0 0 3.901.978h.024h-.001h.026a8.46 8.46 0 0 0 4.811-1.492l-.029.019a12.047 12.047 0 0 1-3.594 2.227l-.082.029c-1.264.506-2.73.8-4.264.8h-.059h.003q-.388 0-.576-.014a11.624 11.624 0 0 1-4.53-1.129l.07.03c-2.766-1.278-4.909-3.512-6.039-6.269l-.029-.079a11.723 11.723 0 0 1-.897-4.536c0-1.685.35-3.288.98-4.741l-.03.077C2.169 4.439 4.433 2.176 7.258.982l.077-.029a11.543 11.543 0 0 1 4.627-.951h.073a11.953 11.953 0 0 1 7.966 3.064l-.011-.01zM24 12v.061c0 1.737-.38 3.386-1.061 4.867l.03-.072a12.269 12.269 0 0 1-2.845 3.964l-.008.007a5.582 5.582 0 0 1-2.968.844h-.005a5.75 5.75 0 0 1-3.431-1.136l.015.011a6.53 6.53 0 0 0 3.378-3.083l.017-.037a10.748 10.748 0 0 0 1.333-5.436v.011l.002-.199c0-1.911-.492-3.706-1.356-5.267l.028.056a6.57 6.57 0 0 0-3.345-3.118l-.043-.016a5.794 5.794 0 0 1 3.4-1.112h.001a5.708 5.708 0 0 1 3.05.884l-.024-.014a11.964 11.964 0 0 1 2.789 3.876l.03.075a11.592 11.592 0 0 1 1.011 4.775v.063v-.003z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Oracle(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.361 12.443h1.592l-.832-1.354l-1.545 2.447h-.713l1.878-2.946c.049-.048.104-.091.162-.128l.004-.002a.392.392 0 0 1 .208-.059h.006a.352.352 0 0 1 .192.06l-.001-.001a.945.945 0 0 1 .166.13l1.878 2.946h-.713l-.333-.546h-1.592zm7.271.547v-2.542h-.594v2.789c0 .04.009.078.025.112l-.001-.002a.358.358 0 0 0 .071.095a.148.148 0 0 0 .079.071h.001a.346.346 0 0 0 .127.024h2.712l.333-.546zm-9.814-.476h.015c.281 0 .536-.113.722-.297a.966.966 0 0 0 .309-.709v-.044a.993.993 0 0 0-.309-.72a1.02 1.02 0 0 0-.722-.297h-2.63v3.089H5.8v-2.542h1.987c.134 0 .254.055.341.143a.48.48 0 0 1 .143.341v.016v-.001a.498.498 0 0 1-.142.345a.464.464 0 0 1-.346.155h-1.7l1.782 1.545h.855l-1.189-1.022zm-6.274 1.022h-.008c-.42 0-.8-.173-1.072-.451a1.464 1.464 0 0 1-.463-1.071v-.045c0-.422.178-.802.463-1.07l.001-.001a1.496 1.496 0 0 1 1.073-.451h1.839c.419 0 .797.173 1.068.451c.279.271.451.649.451 1.068v.026v-.001v.023c0 .42-.174.799-.453 1.07a1.483 1.483 0 0 1-1.068.451h-.026h.001zm1.76-.546h.024a.942.942 0 0 0 .688-.297a.976.976 0 0 0 .285-.69v-.037a.965.965 0 0 0-.285-.686a.965.965 0 0 0-.686-.285h-.028h.001H1.58a.976.976 0 0 0-.69.285a.946.946 0 0 0-.297.689v.025v-.001V12a.96.96 0 0 0 .296.694a.957.957 0 0 0 .694.297h.007zm11.287.546h-.025c-.419 0-.797-.173-1.068-.451a1.483 1.483 0 0 1-.451-1.068v-.05c0-.419.173-.797.451-1.068a1.483 1.483 0 0 1 1.068-.451h.026h-.001h2.138l-.357.546h-1.771a.976.976 0 0 0-.69.285a.946.946 0 0 0-.297.689v.025v-.001v.007a.96.96 0 0 0 .296.694a.957.957 0 0 0 .694.297h2.167l-.357.546zm7.247-.546h-.008a.91.91 0 0 1-.588-.215l.001.001a1.075 1.075 0 0 1-.355-.515l-.002-.008h2.542l.333-.546h-2.875c.065-.208.188-.381.351-.506l.002-.002a.94.94 0 0 1 .584-.202h1.77l.357-.546h-2.163c-.419 0-.797.173-1.068.451a1.483 1.483 0 0 0-.451 1.068v.026v-.001v.025c0 .419.173.797.451 1.068c.271.279.649.451 1.068.451h.026h-.001h1.827l.357-.546z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Origin(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.5 3.109a8.843 8.843 0 0 1 6.115 3.003l.009.01a9.078 9.078 0 0 1 1.695 2.9l.02.064c.309.853.488 1.837.488 2.863c0 .224-.009.447-.025.667l.002-.029a8.193 8.193 0 0 1-.321 1.906l.015-.058a9.124 9.124 0 0 1-.698 1.75l.024-.05a19.73 19.73 0 0 1-1.511 2.472l.041-.06a17.696 17.696 0 0 1-1.766 2.122l.004-.004a18.813 18.813 0 0 1-1.988 1.768l-.044.032c-.651.499-1.388.988-2.156 1.427l-.098.052l-.037.026a.196.196 0 0 1-.094.024l-.016-.001h.001a.297.297 0 0 1-.159-.099v-.001a.264.264 0 0 1-.062-.171v-.003c0-.03.004-.059.013-.086l-.001.002a.183.183 0 0 1 .037-.062c.255-.357.485-.763.67-1.193l.016-.042c.169-.394.298-.852.365-1.331l.003-.028a.119.119 0 0 0-.038-.085a.111.111 0 0 0-.084-.038h-.001a9.117 9.117 0 0 1-.973.065c-.221 0-.439-.015-.652-.043l.025.003a8.823 8.823 0 0 1-6.114-3.003l-.009-.01a9.041 9.041 0 0 1-1.695-2.9l-.02-.063a8.505 8.505 0 0 1-.488-2.867c0-.223.008-.444.025-.662l-.002.029a8.193 8.193 0 0 1 .321-1.906l-.015.058a8.547 8.547 0 0 1 .698-1.726l-.023.046c.484-.937.974-1.732 1.518-2.486l-.04.058A18.27 18.27 0 0 1 4.24 3.326l-.004.004a18.914 18.914 0 0 1 1.989-1.767l.044-.033c.652-.5 1.388-.989 2.156-1.428l.097-.051l.038-.025A.18.18 0 0 1 8.656 0h.014h-.001a.3.3 0 0 1 .158.1a.251.251 0 0 1 .062.167v.016a.228.228 0 0 1-.013.078V.359A.162.162 0 0 1 8.84.42a6.34 6.34 0 0 0-.671 1.195l-.016.042a5.321 5.321 0 0 0-.363 1.329l-.003.029a.125.125 0 0 0 .125.122h.001q.393-.049.8-.061a5.168 5.168 0 0 1 .826.038l-.025-.003zm-.807 12.368a3.307 3.307 0 0 0 2.523-.857l-.003.002a3.398 3.398 0 0 0 1.176-2.392v-.009a3.348 3.348 0 0 0-.858-2.537l.002.003a3.29 3.29 0 0 0-2.394-1.162h-.006a3.347 3.347 0 0 0-2.537.857l.003-.003a3.295 3.295 0 0 0-1.164 2.394v.006a3.37 3.37 0 0 0 .849 2.538l-.003-.003a3.27 3.27 0 0 0 2.406 1.162z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Outdent(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M27.999 21.333H1.334a1.334 1.334 0 0 0-.002 2.666h26.667a1.334 1.334 0 0 0 .002-2.666zm0-5.333H14.667l-.037-.001a1.334 1.334 0 1 0 0 2.668l.039-.001h-.002H28A1.334 1.334 0 0 0 28.002 16H28zm0-5.333H14.667l-.037-.001a1.334 1.334 0 1 0 0 2.668l.039-.001h-.002H28a1.334 1.334 0 0 0 .002-2.666H28zM1.334 2.666h26.665l.037.001a1.334 1.334 0 1 0 0-2.668L27.997 0h.002H1.334a1.334 1.334 0 0 0-.002 2.666zm26.665 2.667H14.667l-.037-.001a1.334 1.334 0 1 0 0 2.668l.039-.001h-.002H28l.037.001a1.334 1.334 0 1 0 0-2.668l-.039.001H28zm-22.275.39L.391 11.057c-.241.241-.39.574-.39.942s.149.701.39.942l5.333 5.333A1.334 1.334 0 0 0 8 17.332V6.666a1.334 1.334 0 0 0-2.276-.942z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaperPlane(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23.615.161a.837.837 0 0 1 .36.862l.001-.005l-3.426 20.56a.848.848 0 0 1-.424.6l-.004.002a.819.819 0 0 1-.406.107h-.01a.9.9 0 0 1-.326-.069l.006.002l-7.054-2.88l-3.989 4.377a.79.79 0 0 1-.604.281h-.026h.001h-.026a.76.76 0 0 1-.287-.056l.005.002a.808.808 0 0 1-.398-.311l-.002-.003a.856.856 0 0 1-.147-.482v-6.055L.539 14.51a.787.787 0 0 1-.535-.736a.78.78 0 0 1 .422-.788l.004-.002L22.705.134a.772.772 0 0 1 .912.027L23.615.16zm-4.578 20.065l2.96-17.709l-19.196 11.07l4.498 1.834L18.85 6.868l-6.4 10.668z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paperclip(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.818 0a5.455 5.455 0 0 0-5.455 5.455v9.818a3.273 3.273 0 1 0 6.546 0V9.608a1.092 1.092 0 0 0-2.182-.002v5.698a1.091 1.091 0 0 1-2.182 0v-.032v.002v-9.818a3.273 3.273 0 1 1 6.546 0v10.906a5.455 5.455 0 0 1-10.91 0V5.425a1.091 1.091 0 0 0-2.182 0v.032v-.002v10.906a7.636 7.636 0 0 0 15.272 0V5.455A5.461 5.461 0 0 0 9.816 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paragraph(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.545 0H6.181a6.182 6.182 0 0 0 0 12.364h2.546V22.94a1.091 1.091 0 0 0 2.182 0v-.032v.002V2.183h2.182v20.76a1.091 1.091 0 0 0 2.182 0v-.032v.002V2.184h3.304a1.091 1.091 0 0 0 0-2.182h-.032h.002zM6.182 10.182a4 4 0 0 1 0-8h2.546v8z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ParalysisDisability(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.878 17.986c-.924 2.33-3.15 3.951-5.758 3.968h-.005a6.08 6.08 0 0 1-3.62-10.965l.016-.012l-.355-2.191c-2.499 1.418-4.158 4.06-4.158 7.09a8.124 8.124 0 0 0 15.377 3.663l.021-.046z"/><path fill="currentColor" d="M16.248 15.07a1.747 1.747 0 0 0-.234-.194l-.006-.004c-.219-.15-.49-.24-.782-.24H8.895l-.819-5.047h6.758a.994.994 0 1 0 0-1.988H7.752l-.111-.672a1.411 1.411 0 1 0-2.783.47l-.001-.008l1.44 8.825a1.524 1.524 0 0 0 1.446 1.237h6.905l5.409 5.41a1.406 1.406 0 1 0 1.988-1.988zM8.345 2.691a2.691 2.691 0 1 1-5.383 0a2.691 2.691 0 0 1 5.383 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Parasol(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.397 13.829L0 23.226l.775.775l9.373-9.373a7.573 7.573 0 0 1-.74-.786l-.011-.014zM4.426 5.262a.534.534 0 0 0-.31.651l-.001-.004a6.399 6.399 0 0 0 3.263 3.819l.037.017c.069.04.153.064.242.064h.014h-.001h.001a.552.552 0 0 0 .313-.097l-.002.001a40.543 40.543 0 0 1 10.372-5.441l.287-.088A.095.095 0 0 0 18.625 4h-.001a22.419 22.419 0 0 0-4.968-.543h-.052h.003h-.006a24.15 24.15 0 0 0-9.336 1.865l.16-.059zm3.236 3.884a7.066 7.066 0 0 1-1.709-1.285l-.001-.001a5.139 5.139 0 0 1-1.18-1.993l-.01-.036c2.619-1.09 5.662-1.723 8.852-1.726h.001a23.38 23.38 0 0 1 3.001.196l-.117-.012c-3.375 1.369-6.285 3-8.955 4.941l.12-.083z"/><path fill="currentColor" d="M3.938 4.832c2.847-1.242 6.165-1.965 9.652-1.965c1.514 0 2.996.136 4.434.397l-.151-.023a.093.093 0 0 0 .104-.055v-.001a.1.1 0 0 0-.032-.12A15.067 15.067 0 0 0 8.799-.001c-2.253 0-4.391.491-6.313 1.371l.095-.039c-.543.312-.144 2.389.75 3.372a.543.543 0 0 0 .61.127l-.004.001zm16.053.519a.113.113 0 0 0-.086-.08h-.009a.105.105 0 0 0-.096.063v.001c-1.426 4.099-3.317 7.653-5.664 10.852l.08-.114a.54.54 0 0 0-.007.617l-.001-.002a8.789 8.789 0 0 0 3.702 3.293l.053.023c.054.02.117.032.183.032h.01a.549.549 0 0 0 .494-.325l.001-.003c1.192-2.743 1.885-5.938 1.885-9.295c0-1.794-.198-3.541-.573-5.222l.03.159zm-.199 7.151a23.846 23.846 0 0 1-1.753 7.02l.059-.157a8.22 8.22 0 0 1-3.288-2.939l-.02-.033a41.785 41.785 0 0 0 4.811-8.627l.103-.282a23.129 23.129 0 0 1 .082 5.109l.007-.092z"/><path fill="currentColor" d="M20.914 6.038a.098.098 0 0 0-.178.057l.001.016V6.11c.239 1.289.375 2.773.375 4.288c0 3.53-.741 6.887-2.075 9.925l.062-.159a.536.536 0 0 0 .181.646l.002.001c1.062.727 3.2 1.11 3.356.607c.008-.016.016-.04.024-.056c.832-1.818 1.318-3.943 1.318-6.182c0-3.451-1.153-6.632-3.094-9.18zm-1.56-1.403a.109.109 0 0 0-.097-.024h.001C15.137 5.998 11.56 7.862 8.34 10.189l.118-.081a.536.536 0 0 0-.182.644l-.001-.004a10.056 10.056 0 0 0 2.089 3.284l-.004-.004a8.048 8.048 0 0 0 2.765 1.795l.055.019a.529.529 0 0 0 .622-.19l.001-.002a40.047 40.047 0 0 0 5.493-10.629l.084-.285a.088.088 0 0 0 .007-.033a.086.086 0 0 0-.033-.067z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Passport(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.496 17.652a3.141 3.141 0 0 0-.415-.026l-.097.001h.005h-.346v.989h.386a1.615 1.615 0 0 0 .566-.06l-.012.003a.48.48 0 0 0 .216-.167l.001-.001a.496.496 0 0 0-.033-.587l.001.001a.427.427 0 0 0-.271-.152h-.002zm2.242 1.495h.948l-.474-1.294z"/><path fill="currentColor" d="M18.327 0H1.776C.795 0 0 .795 0 1.776v20.449c0 .981.795 1.776 1.776 1.776h16.552c.981 0 1.776-.795 1.776-1.776V1.776C20.104.795 19.309 0 18.328 0zm-7.92 2.72c.072.008.153.008.225.016c.242.024.457.12.628.267l-.002-.001c.373.329.682.72.913 1.159l.011.022c.119.212.238.466.339.729l.014.042c-.621.084-1.358.142-2.104.16h-.024zm2.419 8.405a18.984 18.984 0 0 0-2.399-.192h-.02V8.997h2.72c-.031.772-.14 1.503-.32 2.205l.016-.073zm-2.419.666a18.92 18.92 0 0 1 2.251.171l-.098-.011a6.594 6.594 0 0 1-.396.88l.018-.036a4 4 0 0 1-.919 1.177l-.005.004a1.115 1.115 0 0 1-.622.264h-.005a1.868 1.868 0 0 1-.223.016h-.002zm2.387 1.334c.167-.299.327-.652.459-1.018l.016-.051c.37.064.747.145 1.12.24h.008c-.08.097-.169.185-.249.266a5.724 5.724 0 0 1-1.724 1.191l-.036.015c.144-.186.278-.395.393-.615l.011-.023zm3.02-4.13a5.812 5.812 0 0 1-.874 2.595l.014-.024l-.097-.024a21.116 21.116 0 0 0-1.334-.306c.172-.661.286-1.429.319-2.218l.001-.022zm-5.407-3.029a21.607 21.607 0 0 0 2.506-.197l-.111.013c.179.651.298 1.405.329 2.182l.001.02h-2.72zm3.094-.297c.45-.08.88-.177 1.31-.298l.096-.024c.492.748.816 1.646.899 2.614l.001.021H13.83a10.787 10.787 0 0 0-.348-2.384l.017.075zm.835-1.061c-.362.089-.731.169-1.093.233a7.48 7.48 0 0 0-.461-1.036l.02.04a6.395 6.395 0 0 0-.422-.669l.013.019a5.863 5.863 0 0 1 1.76 1.205c.056.072.121.137.185.209zm-5.48-1.607c.169-.146.385-.242.622-.265h.005c.066-.009.144-.015.223-.016h.002v2.386a18.499 18.499 0 0 1-2.226-.171l.096.011c.119-.308.238-.562.371-.807l-.018.036c.239-.457.542-.845.904-1.169l.004-.003zm-2.258 8.236c-.458.088-.908.185-1.334.306l-.097.024a5.568 5.568 0 0 1-.859-2.551l-.001-.02h1.977c.022.807.134 1.578.328 2.316l-.016-.071zm-.876 1.068a13.3 13.3 0 0 1 1.12-.24c.15.419.31.771.494 1.109l-.02-.04c.14.253.275.466.422.669l-.013-.019a5.863 5.863 0 0 1-1.76-1.205c-.08-.089-.16-.177-.24-.274zm1.832-.354c.64-.093 1.386-.151 2.143-.16h.01v2.466a5.792 5.792 0 0 1-.538-.056a2.553 2.553 0 0 1-.317-.228l.004.003a4.091 4.091 0 0 1-.913-1.159l-.011-.022a6.389 6.389 0 0 1-.363-.793l-.016-.047zm2.154-1.021c-.869.017-1.706.086-2.526.206l.108-.013a9.664 9.664 0 0 1-.313-2.11l-.001-.019h2.732zM7.314 5.78a17.36 17.36 0 0 0 2.384.185h.01v2.017H6.976c.037-.797.157-1.549.352-2.27l-.017.074zm.409-2.587c-.134.185-.27.397-.393.618l-.017.033a6.4 6.4 0 0 0-.427.949l-.014.047a17.74 17.74 0 0 1-1.093-.233h-.008a5.68 5.68 0 0 1 .193-.209a5.927 5.927 0 0 1 1.72-1.183l.037-.015zM5.305 5.37c.426.113.86.209 1.31.298a10.474 10.474 0 0 0-.337 2.294l-.001.02H4.3a5.708 5.708 0 0 1 .921-2.657l-.013.021zm.9 13.62a1.147 1.147 0 0 1-.361.167l-.008.002a4.057 4.057 0 0 1-.636.049l-.1-.001h.005h-.458v1.318H3.94V17.03h1.133a3.65 3.65 0 0 1 .858.06l-.023-.004a.959.959 0 0 1 .505.343l.002.002c.126.193.201.429.201.683c0 .196-.044.381-.124.546l.003-.008a.842.842 0 0 1-.288.341l-.003.002zm3.021 1.534l-.306-.795h-1.4l-.29.795h-.747l1.36-3.495h.747l1.398 3.495zm3.64-.458a1.003 1.003 0 0 1-.468.391l-.007.002c-.206.082-.444.13-.694.13l-.073-.001h.003a1.544 1.544 0 0 1-1.025-.308l.004.003a1.338 1.338 0 0 1-.425-.888v-.004l.683-.064c.024.2.114.375.249.505a.94.94 0 0 0 1.033.014l-.004.002a.441.441 0 0 0 .177-.337v-.009a.312.312 0 0 0-.073-.201v.001a.661.661 0 0 0-.252-.151l-.005-.001c-.08-.032-.273-.08-.56-.153a1.944 1.944 0 0 1-.8-.349l.005.004a.903.903 0 0 1-.32-.689v-.025c0-.184.057-.354.155-.494l-.002.003c.108-.155.255-.277.427-.351l.006-.002c.191-.077.412-.121.644-.121l.049.001h-.002a1.438 1.438 0 0 1 .991.293l-.004-.003a.997.997 0 0 1 .346.755v.017v-.001l-.707.032a.623.623 0 0 0-.193-.385a.695.695 0 0 0-.397-.123l-.055.002h.002a.768.768 0 0 0-.483.13l.003-.002a.27.27 0 0 0-.113.217v.003c0 .084.041.159.104.206h.001c.188.111.407.193.64.231l.011.001c.291.063.546.149.787.259l-.024-.01c.159.085.289.205.383.35l.002.004a.997.997 0 0 1 .138.51l-.001.046v-.002v.027c0 .202-.06.39-.164.546l.002-.004zm3.254 0a1.003 1.003 0 0 1-.468.391l-.007.002c-.206.082-.444.13-.694.13l-.073-.001h.003a1.544 1.544 0 0 1-1.025-.308l.004.003a1.338 1.338 0 0 1-.425-.888v-.004l.683-.064c.024.2.114.375.249.505a.94.94 0 0 0 1.033.014l-.004.002a.441.441 0 0 0 .177-.337v-.009a.312.312 0 0 0-.073-.201v.001a.661.661 0 0 0-.252-.151l-.005-.001c-.08-.032-.273-.08-.56-.153a1.944 1.944 0 0 1-.8-.349l.005.004a.903.903 0 0 1-.32-.689v-.025c0-.184.057-.354.155-.494l-.002.003a.999.999 0 0 1 .427-.351l.006-.002c.191-.077.413-.121.644-.121l.049.001h-.002a1.438 1.438 0 0 1 .991.293l-.004-.003a.997.997 0 0 1 .346.755v.017v-.001l-.707.032a.623.623 0 0 0-.193-.385a.695.695 0 0 0-.397-.123l-.055.002h.002a.768.768 0 0 0-.483.13l.003-.002a.27.27 0 0 0-.113.217v.003c0 .084.041.159.104.206h.001c.187.111.406.192.639.231l.011.002c.291.063.546.149.787.259l-.024-.01c.159.085.289.205.383.35l.002.004a.997.997 0 0 1 .138.51l-.001.046v-.002l.002.059a.878.878 0 0 1-.165.514l.002-.003z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PassportAlt(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.36 19.68c-.24 0-.4.24-.4.64c0 .32.16.64.48.64c.24 0 .4-.24.4-.64c0-.32-.16-.64-.48-.64m-6.4 0c0 .08-.08.24-.08.32l-.08.4h.4l-.08-.4c-.08-.08-.08-.16-.16-.32m-1.52 0h-.16v.56h.16c.24 0 .32-.08.32-.32s-.08-.24-.32-.24m6.16 0h-.16v.56h.16c.24 0 .32-.08.32-.32s-.16-.24-.32-.24m3.52 0h-.16v.48h.16c.24 0 .32-.08.32-.24s-.16-.24-.32-.24m12-13.84L11.44.08a.554.554 0 0 0-.798.237l-.001.003V.4l-.24.4l-1.36 2.8h4.896c.698 0 1.264.566 1.264 1.264v.017v-.001v18.4c.08-.08.16-.16.16-.24l8-16.4a.57.57 0 0 0-.237-.798l-.003-.001zM9.2 14.64a6.04 6.04 0 0 0-1.665-.319l-.015-.001v2.32a3.008 3.008 0 0 0 1.675-1.979zm.24-4.96a5.278 5.278 0 0 1-1.829.321L7.516 10h.005v1.28h2.16a5.982 5.982 0 0 0-.251-1.642l.011.042zM7.52 7.12v2.16l.075.001c.579 0 1.131-.118 1.632-.331l-.027.01a2.616 2.616 0 0 0-1.662-1.834z"/><path fill="currentColor" d="M10.24 14.16c.29.159.514.409.637.711l.003.009a4.591 4.591 0 0 0 1.119-2.622L12 12.24h-1.52a5.276 5.276 0 0 1-.251 1.958zm-6.32 1.36c.252.224.542.412.859.552l.021.008c-.221-.3-.41-.643-.55-1.01l-.01-.03a1.019 1.019 0 0 0-.318.473zm6.4-7.6c-.3-.235-.639-.472-.991-.691L9.28 7.2c.285.345.527.741.708 1.169L10 8.4c.16 0 .32-.24.32-.48m-.88 8.16c.309-.147.573-.335.8-.56a.657.657 0 0 0-.317-.478l-.003-.002c-.082.399-.25.749-.484 1.045zm1.12-4.72h1.52a5.68 5.68 0 0 0-1.051-2.736l.011.016a1.575 1.575 0 0 1-.711.716l-.009.004c.08.64.16 1.28.24 2m-.88.88H7.52v1.2a5.498 5.498 0 0 1 1.956.413l-.036-.013a5.843 5.843 0 0 0 .24-1.596zm-4.96-4.8a2.65 2.65 0 0 0-.803.483l.003-.003a.69.69 0 0 0 .239.479l.001.001c.128-.371.319-.691.562-.963zm-.8 1.84a2.462 2.462 0 0 1-.635-.632L3.28 8.64a5.058 5.058 0 0 0-1.119 2.782l-.001.018H3.6a8.777 8.777 0 0 1 .335-2.222zm2.72 7.2v-2.16a9.42 9.42 0 0 0-1.667.337l.067-.017a2.647 2.647 0 0 0 1.582 1.834z"/><path fill="currentColor" d="M14.24 5.12v-.024a.618.618 0 0 0-.558-.615H.639c-.016-.002-.035-.002-.054-.002a.59.59 0 0 0-.586.645v-.002v18.249c0 .349.283.631.631.631h12.969a.69.69 0 0 0 .64-.638v-.002zM1.36 11.2a5.819 5.819 0 0 1 5.257-5.278l.023-.002v-.08h.88v.08a5.737 5.737 0 1 1-.9 11.439l.02.001c-2.977-.263-5.293-2.744-5.293-5.767c0-.138.005-.276.014-.411zm.08 9.36h-.16v.64h-.4v-1.84h.56a.77.77 0 0 1 .562.161L2 19.52c.16.08.16.24.16.4a.604.604 0 0 1-.16.4a.724.724 0 0 1-.54.24zm1.92.64l-.16-.48h-.48l-.16.48h-.4l.56-1.84h.56l.56 1.84zm1.12 0a1.01 1.01 0 0 1-.486-.083L4 21.12l.08-.32a1.006 1.006 0 0 0 .484.08H4.56c.16 0 .24-.08.24-.16s-.08-.16-.32-.24c-.24-.16-.48-.32-.48-.56c0-.32.24-.56.72-.56a1.01 1.01 0 0 1 .486.083L5.2 19.44l-.08.32a.596.596 0 0 0-.403-.08h.003c-.16 0-.24.08-.24.16s.08.16.32.24c.27.05.473.28.48.559v.001c0 .32-.24.56-.8.56m1.52 0a1.01 1.01 0 0 1-.486-.083l.006.003v-.32a1.006 1.006 0 0 0 .484.08H6c.16 0 .32-.08.32-.16s-.08-.16-.32-.24a.529.529 0 0 1-.481-.527l.001-.035v.002c0-.32.24-.56.72-.56a1.01 1.01 0 0 1 .486.083l-.006-.003l-.08.32a.596.596 0 0 0-.403-.08h.003c-.16 0-.24.08-.24.16s.08.16.32.24c.27.05.473.28.48.559v.001a.698.698 0 0 1-.804.56l.004.001zm2.16-.8a.53.53 0 0 1-.564.159l.004.001h-.16v.64h-.4v-1.84h.56a.77.77 0 0 1 .562.161l-.002-.001c.16.08.16.24.16.4c-.018.18-.075.343-.163.485zm2.08-.16a.844.844 0 0 1-.798.96H9.28a.888.888 0 0 1-.88-.88v-.08a.897.897 0 0 1 .796-.96h.088a.88.88 0 0 1 .876.797v.004zm1.2.96c-.08-.16-.08-.24-.16-.4c-.08-.24-.16-.32-.32-.32h-.16v.72h-.4v-1.84h.56a.671.671 0 0 1 .561.161l-.001-.001a.604.604 0 0 1 .16.399v.001a.542.542 0 0 1-.317.479l-.003.001c.16.08.16.16.24.32c.018.18.075.343.163.485l-.003-.005zm2-1.44h-.48v1.44h-.4v-1.44H12v-.32h1.36v.32z"/><path fill="currentColor" d="M3.92 14.16a13.941 13.941 0 0 1-.315-1.941L3.6 12.16H2.24a5.405 5.405 0 0 0 1.05 2.654l-.01-.014c.174-.253.387-.466.632-.635zm2.72-1.92H4.48c.001.571.089 1.122.251 1.639l-.011-.039a7.944 7.944 0 0 1 1.889-.398l.031-.002zm0-5.28c-.88.16-1.44 1.12-1.68 1.84c.491.196 1.06.337 1.653.398l.027.002zm0 4.4v-1.28a5.498 5.498 0 0 1-1.956-.413l.036.013a5.946 5.946 0 0 0-.24 1.679v.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paste(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.092 14.909a1.08 1.08 0 0 0-.25-.692l.001.002v-.005q-.03-.037-.064-.07l-.009-.009l-.029-.029l-4.011-4.011V1.82c0-.603-.489-1.091-1.091-1.091h-5.523A1.093 1.093 0 0 0 9.088.002H7.632c-.472 0-.874.3-1.025.72L6.605.73H1.092C.489.73.001 1.219.001 1.821v18.181c0 .603.489 1.091 1.091 1.091h3.273v1.818c0 .603.489 1.091 1.091 1.091H20c.603 0 1.091-.489 1.091-1.091v-7.972l.001-.028zm-8.001-5.366l4.275 4.275h-4.275zM2.182 18.909V2.913h2.245c.154.427.556.727 1.028.727h5.818c.472 0 .874-.3 1.025-.72l.002-.008h2.245v5.002l-1.774-1.777a1.03 1.03 0 0 0-.069-.063l-.023-.018l-.053-.04l-.026-.021a.715.715 0 0 0-.059-.036l-.027-.013a.86.86 0 0 0-.081-.039l-.033-.014l-.057-.02l-.038-.012l-.064-.016l-.028-.007c-.031-.006-.062-.01-.093-.014H5.454c-.603 0-1.091.489-1.091 1.091v11.992zm4.363 2.912V8h4.363v6.909c0 .602.489 1.09 1.091 1.09h6.91v5.818z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pause(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.801 23.999h-3.6a1.202 1.202 0 0 1-1.2-1.2v-21.6c0-.663.537-1.2 1.2-1.201h3.6a1.202 1.202 0 0 1 1.2 1.2v21.6c0 .663-.537 1.2-1.2 1.201m11.999 0h-3.6a1.2 1.2 0 0 1-1.2-1.2V1.198a1.2 1.2 0 0 1 1.2-1.2h3.6a1.2 1.2 0 0 1 1.2 1.2v21.6a1.2 1.2 0 0 1-1.2 1.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paw(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.189 5.376v.031a5.27 5.27 0 0 1-.308 1.779l.012-.037a3.468 3.468 0 0 1-.98 1.442l-.004.003c-.421.379-.981.61-1.595.61h-.048h.002a3.087 3.087 0 0 1-2.16-.899a5.611 5.611 0 0 1-1.427-2.079l-.013-.038a6.458 6.458 0 0 1-.469-2.355v-.035c0-.624.109-1.223.308-1.779l-.012.037a3.468 3.468 0 0 1 .98-1.442l.004-.003a2.38 2.38 0 0 1 1.65-.61h-.003a3.076 3.076 0 0 1 2.16.899a5.66 5.66 0 0 1 1.416 2.071l.013.038c.289.7.46 1.513.469 2.364v.003zm-5.345 7.548l.001.084a3.63 3.63 0 0 1-.665 2.1l.008-.012a2.15 2.15 0 0 1-1.765.924l-.099-.002h.005a3.346 3.346 0 0 1-2.214-.87l.003.003a6.063 6.063 0 0 1-2.117-4.454v-.008L0 10.607c0-.785.246-1.513.665-2.11l-.008.012a2.141 2.141 0 0 1 1.864-.93h-.005a3.346 3.346 0 0 1 2.214.87l-.003-.003a6.091 6.091 0 0 1 2.117 4.473v.007zm6.161-.422a6.991 6.991 0 0 1 3.998 1.531l-.014-.011a14.03 14.03 0 0 1 3.547 3.652l.032.051a7.336 7.336 0 0 1 1.439 3.957l.001.02l.002.087c0 .404-.099.786-.274 1.121l.006-.013a1.632 1.632 0 0 1-.748.699l-.01.004c-.291.139-.63.248-.984.309l-.024.003a7.045 7.045 0 0 1-1.116.086h-.075h.004a8.97 8.97 0 0 1-2.988-.726l.058.023a8.712 8.712 0 0 0-2.82-.701l-.031-.002c-1.103.082-2.125.33-3.075.719l.067-.024a9.862 9.862 0 0 1-3.105.694l-.028.001q-2.866.002-2.866-2.279a6.459 6.459 0 0 1 .892-3.022l-.017.03a13.458 13.458 0 0 1 2.184-3.012l-.004.004a12.865 12.865 0 0 1 2.868-2.249l.062-.033a6.103 6.103 0 0 1 3.003-.921h.012zm3.735-3.297h-.046a2.383 2.383 0 0 1-1.597-.612l.002.002a3.446 3.446 0 0 1-.976-1.422l-.008-.024a5.162 5.162 0 0 1-.297-1.742v-.033v.002c.009-.855.18-1.667.485-2.411l-.016.044A5.674 5.674 0 0 1 15.713.903L15.717.9a3.08 3.08 0 0 1 2.159-.898h.047c.614 0 1.174.232 1.597.612l-.002-.002c.438.383.776.869.976 1.422l.008.024c.191.522.301 1.125.301 1.753v.02v-.001a6.536 6.536 0 0 1-.485 2.403l.016-.044a5.62 5.62 0 0 1-1.436 2.113l-.004.004a3.088 3.088 0 0 1-2.159.896h-.001zm6.75-1.624l.092-.002c.731 0 1.376.366 1.762.925l.005.007a3.65 3.65 0 0 1 .657 2.099l-.001.086v-.004a6.056 6.056 0 0 1-2.109 4.454l-.008.007a3.344 3.344 0 0 1-2.209.867h-.002a2.143 2.143 0 0 1-1.854-.915l-.005-.007a3.618 3.618 0 0 1-.656-2.176v.004a6.092 6.092 0 0 1 2.109-4.474l.008-.007a3.345 3.345 0 0 1 2.205-.865h.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Paypal(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M33.6 24H2.4A2.4 2.4 0 0 1 0 21.6V2.4A2.4 2.4 0 0 1 2.4 0h31.2A2.4 2.4 0 0 1 36 2.4v19.2a2.4 2.4 0 0 1-2.4 2.4M14.032 10.325a.184.184 0 0 0-.142.07a.22.22 0 0 0-.063.15c0 .021.098.32.308.933l.454 1.326l.211.618q.355 1.04.371 1.106a14.561 14.561 0 0 0-1.259 1.828l-.038.07l-.002.025c0 .1.081.182.182.182l.027-.002h-.001h1.218a.335.335 0 0 0 .283-.159l.001-.001l4.032-5.819a.162.162 0 0 0 .033-.097l-.001-.015v.001a.223.223 0 0 0-.063-.15a.183.183 0 0 0-.141-.07h-1.219a.337.337 0 0 0-.284.159l-.001.001l-1.68 2.466l-.695-2.372a.334.334 0 0 0-.323-.254l-.026.001h.001zm13.978-.111h-.026c-.758 0-1.439.329-1.909.851l-.002.002a2.846 2.846 0 0 0-.8 1.982v.027v-.001c-.002.029-.002.064-.002.098c0 .537.209 1.026.549 1.389l-.001-.001a1.962 1.962 0 0 0 1.468.554h-.005c.337-.004.657-.073.949-.197l-.016.006c.295-.111.545-.284.742-.504l.002-.002c-.01.057-.021.105-.034.151l.002-.01a.829.829 0 0 0-.032.188v.002c0 .168.069.253.206.253h1.091a.32.32 0 0 0 .348-.299v-.001l.648-4.127a.209.209 0 0 0-.047-.177a.194.194 0 0 0-.157-.08h-1.203c-.147 0-.24.175-.27.522a1.659 1.659 0 0 0-1.509-.627l.008-.001zm-17.963 0h-.028c-.757 0-1.438.329-1.906.851l-.002.002a2.846 2.846 0 0 0-.8 1.982v.027v-.001a2.024 2.024 0 0 0 .548 1.486l-.001-.001c.354.344.838.556 1.372.556l.096-.002h-.005c.332-.005.647-.074.934-.197l-.016.006c.299-.114.553-.286.758-.505l.001-.001c-.037.098-.06.212-.063.33v.002c0 .168.07.253.206.253h1.09a.32.32 0 0 0 .348-.299v-.001l.649-4.127a.212.212 0 0 0-.048-.175a.194.194 0 0 0-.157-.08H11.82c-.147 0-.24.175-.269.522a1.688 1.688 0 0 0-1.51-.628l.008-.001zM32.612 8l-.016-.001a.19.19 0 0 0-.19.174v.001l-1.028 6.578a.19.19 0 0 0 .047.181a.2.2 0 0 0 .154.071h1.05l.032.002a.3.3 0 0 0 .3-.3v-.002l1.026-6.465l.001-.022a.263.263 0 0 0-.053-.158v.001a.193.193 0 0 0-.14-.061L33.78 8h.001zm-11.67 0a.308.308 0 0 0-.348.299l-1.027 6.452a.214.214 0 0 0 .047.174a.19.19 0 0 0 .156.08h1.305a.253.253 0 0 0 .248-.204v-.002l.285-1.834a.306.306 0 0 1 .111-.205h.001a.485.485 0 0 1 .237-.103h.003c.081-.015.174-.023.269-.024h.001c.084 0 .184.005.3.016a3.077 3.077 0 0 0 .382.02a2.9 2.9 0 0 0 1.976-.773l-.002.002a2.86 2.86 0 0 0 .774-2.135v.008a1.53 1.53 0 0 0-.605-1.35l-.004-.003a2.73 2.73 0 0 0-1.6-.419h.006zM2.96 8a.3.3 0 0 0-.332.298V8.3L1.6 14.752a.213.213 0 0 0 .047.174a.19.19 0 0 0 .156.08h1.204a.308.308 0 0 0 .348-.299l.285-1.739a.31.31 0 0 1 .11-.205h.001a.485.485 0 0 1 .237-.103h.003a1.58 1.58 0 0 1 .268-.024h.001c.084 0 .184.005.3.016a3.077 3.077 0 0 0 .381.02c.763 0 1.457-.293 1.977-.773l-.002.002a2.862 2.862 0 0 0 .774-2.135v.008a1.53 1.53 0 0 0-.605-1.35l-.004-.003a2.731 2.731 0 0 0-1.593-.42h.006zm24.984 5.835l-.048.001a1.08 1.08 0 0 1-.698-.255l.002.001a.887.887 0 0 1-.286-.653l.001-.045v-.027c0-.374.151-.713.396-.959c.241-.247.577-.4.949-.4h.033h-.002l.034-.001c.269 0 .514.099.702.262l-.001-.001a.917.917 0 0 1 .294.674l-.001.048v-.002v.024c0 .369-.153.703-.4.94a1.341 1.341 0 0 1-.949.391h-.025h.001zm-17.979 0l-.05.001c-.262 0-.503-.096-.687-.255l.001.001a.9.9 0 0 1-.278-.651l.001-.047v-.027c0-.374.151-.713.396-.959c.241-.247.576-.4.948-.4h.035h-.002l.034-.001c.269 0 .514.099.702.262l-.001-.001a.916.916 0 0 1 .293.673l-.001.049v-.002v.035c0 .369-.153.701-.4.938a1.374 1.374 0 0 1-.954.384h-.036h.002zm11.907-2.56l.269-1.691a.19.19 0 0 1 .206-.173h-.001h.285l.061-.001c.173 0 .344.012.511.035l-.019-.002c.155.028.29.098.396.198a.61.61 0 0 1 .191.444l-.001.038v-.002a1.022 1.022 0 0 1-.346.901l-.001.001a1.859 1.859 0 0 1-1.05.24h.006l-.505.016zm-17.979 0l.269-1.691a.189.189 0 0 1 .206-.173h-.001h.3a1.765 1.765 0 0 1 .927.186l-.01-.005c.189.121.253.378.19.766a.904.904 0 0 1-.408.726l-.003.002a3.326 3.326 0 0 1-1.484.188z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PaypalP(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.23 8.652a7.134 7.134 0 0 1-.063 2.779l.009-.047q-1.165 5.946-7.567 5.946h-.6a.868.868 0 0 0-.581.222l.001-.001a.973.973 0 0 0-.319.563l-.001.006l-.054.254l-.737 4.634l-.026.201a.96.96 0 0 1-.327.568l-.001.001a.892.892 0 0 1-.591.222H6.01l-.023.001a.533.533 0 0 1-.417-.2l-.001-.001a.6.6 0 0 1-.12-.484l-.001.004q.121-.75.355-2.25t.355-2.25t.362-2.24t.362-2.24a.518.518 0 0 1 .578-.495h-.002h1.753a13.659 13.659 0 0 0 3.25-.298l-.09.016a8.053 8.053 0 0 0 3.847-1.935l-.007.006a8.042 8.042 0 0 0 2.059-3.238l.017-.057c.181-.495.343-1.094.458-1.709l.011-.072a.22.22 0 0 1 .034-.102v.001a.05.05 0 0 1 .035-.015a.05.05 0 0 1 .013.002a.266.266 0 0 1 .08.048a3.36 3.36 0 0 1 1.309 2.141l.003.02zm-2.304-3.777a9.484 9.484 0 0 1-.638 3.225l.022-.065a6.538 6.538 0 0 1-4 4.204l-.045.014a10.294 10.294 0 0 1-3.375.56h-.001q0 .014-1.206.014l-1.206-.014a1.44 1.44 0 0 0-1.58 1.279l-.001.006q-.026.107-1.138 7.098a.14.14 0 0 1-.161.133h.001h-3.95a.647.647 0 0 1-.639-.74v.003L3.116.891c.044-.259.177-.482.366-.639l.002-.001c.181-.158.42-.254.681-.254h8.012c.473.019.921.082 1.353.184l-.048-.01a8.81 8.81 0 0 1 1.554.45L14.975.6a4.57 4.57 0 0 1 2.18 1.634l.009.013a4.468 4.468 0 0 1 .757 2.631v-.006z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Payu(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M33.6 24H2.4A2.4 2.4 0 0 1 0 21.6V2.4A2.4 2.4 0 0 1 2.4 0h31.2A2.4 2.4 0 0 1 36 2.4v19.2a2.4 2.4 0 0 1-2.4 2.4m-15.159-6.408c-.33 1.04-.672 1.55-1.72 1.658c-.209.018-.344.047-.42.148a.48.48 0 0 0-.031.394l-.001-.003l.028.127c.062.294.166.48.497.48c.034 0 .072 0 .112-.006c1.536-.1 2.358-.926 2.84-2.854l1.642-6.576a.54.54 0 0 0-.029-.456l.001.003a.536.536 0 0 0-.435-.128h.003h-.129a.583.583 0 0 0-.688.54v.002l-1.185 4.927c-.149.606-.356.72-.711.72c-.436 0-.61-.104-.784-.72l-1.342-4.927a.601.601 0 0 0-.706-.538l.004-.001h-.114a.526.526 0 0 0-.432.131a.53.53 0 0 0-.012.46l-.001-.003l1.36 4.97c.254.951.56 1.739 1.687 1.739h.031c.191 0 .376-.029.549-.084l-.013.004zm-6.938-4.32c-2.05 0-3.005.692-3.005 2.177c0 1.44.986 2.233 2.776 2.233c2.127 0 3.076-.724 3.076-2.348v-2.829c0-1.57-1.011-2.334-3.09-2.334h-.072a7.76 7.76 0 0 0-1.527.151l.049-.008c-.361.08-.512.176-.512.577v.114a.63.63 0 0 0 .074.346l-.002-.003a.312.312 0 0 0 .257.135l.023-.001h-.001a.88.88 0 0 0 .225-.034l-.006.002c.46-.098.988-.154 1.529-.154h.021h-.001c1.264 0 1.78.35 1.78 1.209v.766zm12.005-5.178c-.855 0-1.213.132-1.213.95v4.983a3.69 3.69 0 0 0 .539 2.044l-.009-.016c.677 1.062 2 1.624 3.84 1.624s3.175-.56 3.852-1.624a3.667 3.667 0 0 0 .529-2.033v.006v-3.84h.995a.386.386 0 0 0 .386-.386V7.821a.386.386 0 0 0-.386-.386h-1.957a.387.387 0 0 0-.386.387v.273h-.142c-.855 0-1.213.132-1.213.95v4.983a1.374 1.374 0 0 1-.194.782l.004-.006c-.24.367-.709.539-1.481.541s-1.242-.174-1.481-.541a1.387 1.387 0 0 1-.192-.709l.002-.07v.003V9.04c0-.818-.358-.95-1.212-.95zm-20.274 0c-1.13-.003-1.634.501-1.634 1.63v7.255c0 .437.14.577.577.577h.143c.437 0 .577-.14.577-.577V14.16h2.45c2.176 0 3.19-.96 3.19-3.034S7.52 8.091 5.347 8.091zM32.72 5.39a.287.287 0 0 0-.287.286v1.472c0 .158.129.287.287.287h1.454a.286.286 0 0 0 .286-.286V5.68a.289.289 0 0 0-.287-.287zM31.245 4a.195.195 0 0 0-.195.195v.999c0 .108.087.195.194.195h.987a.195.195 0 0 0 .195-.195v-.999A.194.194 0 0 0 32.232 4zM11.259 16.57c-1.015 0-1.508-.367-1.508-1.12c0-.83.495-1.153 1.766-1.153h1.58v.995c-.001.808-.298 1.278-1.838 1.278m-5.912-3.623H2.896V9.896c0-.423.16-.58.58-.58h1.871c1.2 0 1.894.296 1.894 1.809c0 1.182-.302 1.822-1.894 1.822"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Periscope(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.499 3.001a9.853 9.853 0 0 1 2.062 3.168l.024.066a9.4 9.4 0 0 1 .68 3.809v-.013a12.117 12.117 0 0 1-1.484 5.313l.032-.063a28.878 28.878 0 0 1-2.4 3.877l.056-.08q-1.078 1.406-2.25 2.672Q11.016 24 9.798 24q-1.125 0-2.909-1.594a24.712 24.712 0 0 1-2.955-3.266l-.045-.062a21.758 21.758 0 0 1-2.287-3.624l-.057-.126A11.712 11.712 0 0 1 .002 10.06l-.001-.028v-.076a9.75 9.75 0 0 1 1.361-4.985l-.025.045a10.334 10.334 0 0 1 3.518-3.631l.044-.025A9.126 9.126 0 0 1 9.719 0h.082h-.004h.008c1.305 0 2.543.285 3.656.796l-.055-.022a9.588 9.588 0 0 1 3.088 2.221l.006.006zM9.796 23.156q.656 0 2.086-1.313a23.076 23.076 0 0 0 2.834-3.214l.046-.066a24.666 24.666 0 0 0 2.516-4.117l.065-.149a10.398 10.398 0 0 0 1.078-4.245l.001-.02v-.039a9.302 9.302 0 0 0-1.197-4.579l.024.047a9.075 9.075 0 0 0-3.125-3.352l-.037-.022A7.886 7.886 0 0 0 9.817.845h-.019h.001h-.058a8.516 8.516 0 0 0-4.479 1.264l.037-.021a8.898 8.898 0 0 0-3.214 3.329l-.023.046a9.324 9.324 0 0 0-1.175 4.559v.063v-.003a10.459 10.459 0 0 0 1.108 4.278l-.028-.06a23.519 23.519 0 0 0 2.664 4.318l-.039-.053a20.79 20.79 0 0 0 2.934 3.265l.019.017q1.502 1.312 2.25 1.312zM15 9.7v.02a4.82 4.82 0 0 1-.715 2.532l.012-.021a5.368 5.368 0 0 1-1.873 1.862l-.025.014a4.984 4.984 0 0 1-2.568.703h-.035h.002a5.422 5.422 0 0 1-3.801-1.407l.005.004a5.029 5.029 0 0 1-1.687-3.357l-.001-.018A5.058 5.058 0 0 1 5.446 6.32l-.007.009v.007c0 .579.243 1.1.632 1.469l.001.001c.374.39.899.633 1.481.633h.02h-.001h.028a1.97 1.97 0 0 0 1.448-.632l.001-.001c.377-.374.61-.891.61-1.464v-.019c0-.419-.131-.806-.355-1.124l.004.006a2.39 2.39 0 0 0-.901-.794l-.014-.006a4.997 4.997 0 0 1 1.314-.171c.701 0 1.369.141 1.977.395l-.034-.013a5.328 5.328 0 0 1 2.425 1.996l.012.02c.573.849.915 1.895.915 3.022v.05v-.003z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Person(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.845 17.295a7.436 7.436 0 0 0-4.089-2.754l-.051-.011l-1.179 1.99a1.003 1.003 0 0 1-1 1c-.55 0-1-.45-1.525-1.774v-.032a1.25 1.25 0 1 0-2.5 0v.033v-.002c-.56 1.325-1.014 1.774-1.563 1.774a1.003 1.003 0 0 1-1-1l-1.142-1.994A7.47 7.47 0 0 0 .67 17.271l-.014.019a4.475 4.475 0 0 0-.655 2.197v.007c.005.15 0 .325 0 .5v2a2 2 0 0 0 2 2h15.5a2 2 0 0 0 2-2v-2c0-.174-.005-.35 0-.5a4.522 4.522 0 0 0-.666-2.221l.011.02zM4.5 5.29c0 2.92 1.82 7.21 5.25 7.21c3.37 0 5.25-4.29 5.25-7.21v-.065a5.25 5.25 0 1 0-10.5 0v.068z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Persons(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.313 17.295a7.436 7.436 0 0 0-4.089-2.754l-.051-.011l-1.179 1.99a1.003 1.003 0 0 1-1 1c-.55 0-1-.45-1.525-1.774v-.032a1.25 1.25 0 1 0-2.5 0v.033v-.002c-.56 1.325-1.014 1.774-1.563 1.774a1.003 1.003 0 0 1-1-1l-1.142-1.994a7.47 7.47 0 0 0-4.126 2.746l-.014.019a4.475 4.475 0 0 0-.655 2.197v.007c.005.15 0 .325 0 .5v2a2 2 0 0 0 2 2h15.5a2 2 0 0 0 2-2v-2c0-.174-.005-.35 0-.5a4.522 4.522 0 0 0-.666-2.221l.011.02zM7.968 5.29c0 2.92 1.82 7.21 5.25 7.21c3.37 0 5.25-4.29 5.25-7.21v-.065a5.25 5.25 0 1 0-10.5 0v.068zm11.234 1.72c0 1.902 1.186 4.698 3.42 4.698c2.195 0 3.42-2.795 3.42-4.698v-.052a3.421 3.421 0 0 0-6.842 0v.055v-.003zm-19.2 1.6c0 1.902 1.186 4.698 3.42 4.698c2.195 0 3.42-2.795 3.42-4.698v-.052a3.421 3.421 0 0 0-6.842 0v.055v-.003z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Phone(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.839 13.84c-2.372 2.378-5.12 4.648-6.209 3.56c-1.557-1.56-2.517-2.913-5.951-.155s-.796 4.598.712 6.106c1.74 1.74 8.226.091 14.64-6.32s8.06-12.898 6.32-14.64c-1.512-1.505-3.347-4.139-6.105-.711s-1.403 4.39.153 5.946c1.088 1.094-1.182 3.841-3.56 6.214"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Photograph(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12.586 0l-1.604 3.906H0v19.342h16.103v-.899l4.03 1.65l7.346-17.894zM1.588 18.467V5.541h12.926v12.926zm11.114-14.56l.736-1.79l11.958 4.906l-4.906 11.958l-4.379-1.798V3.908z"/><path fill="currentColor" d="M9.858 7.85a1.367 1.367 0 1 0 1.733.843l.003.01a1.368 1.368 0 0 0-1.746-.85z"/><path fill="currentColor" d="m12.632 9.601l-.589-.574l.503-.651l-.8-.202l.116-.814l-.79.225l-.31-.767l-.574.589l-.65-.503l-.202.8l-.814-.116l.225.79l-.767.31l.589.574l-.504.65l.8.202l-.116.814l.79-.225l.31.767l.573-.589l.651.504l.202-.8l.814.116l-.225-.79zm-1.883.837a1.367 1.367 0 1 1 .849-1.746l.003.01a1.374 1.374 0 0 1-.842 1.734zm-4.805-.697l-3.882 5.13v2.712h9.826zm8.161 3.348l-1-1.326l-2.4 3.178l1.704 2.247l.295.395h1.402zm6.647-4.828a2.223 2.223 0 0 0-.532-.063l-.077.001h.004a1.73 1.73 0 0 1 1.546-.174l-.012-.004a1.377 1.377 0 0 0-.296-.607l.002.002a1.66 1.66 0 0 0-2.341-.031c-.106.091-.2.19-.284.298l-.003.005a1.736 1.736 0 0 0-.238-.246l-.002-.002a1.976 1.976 0 0 0-1.935-.453l.014-.004v1.03c.373.087.697.26.962.498l-.002-.002c.12.105.226.222.316.351l.004.006a1.844 1.844 0 0 0-1.29-.434h.004v1.573a1.818 1.818 0 0 1 1.6-.011l-.011-.005v.046a12.456 12.456 0 0 1-1.613 2.988l.024-.035v2.774c.938-.954 1.558-2.821 2.1-5.587a1.995 1.995 0 0 1 .023 3.35l-.008.005a1.979 1.979 0 0 0 .878.079l-.01.001a2.001 2.001 0 0 0 .317-3.894l-.014-.004c.19.002.374.024.551.066l-.017-.003a2.25 2.25 0 0 1 1.692 1.441l.005.016c.176-.195.306-.436.37-.702l.002-.011a2.028 2.028 0 0 0-1.719-2.258z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Php(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 5.601h-.116c-1.61 0-3.18.175-4.69.507l.144-.027a16.125 16.125 0 0 0-3.91 1.343l.094-.042a8.123 8.123 0 0 0-2.57 1.93l-.007.008A3.6 3.6 0 0 0 0 11.684v.004c.019.914.374 1.741.946 2.367l-.002-.003a8.105 8.105 0 0 0 2.529 1.917l.048.021a15.7 15.7 0 0 0 3.71 1.282l.106.019c1.366.305 2.936.48 4.546.48h.123H12h.116c1.61 0 3.18-.175 4.69-.507l-.144.027a16.125 16.125 0 0 0 3.91-1.343l-.094.042a8.123 8.123 0 0 0 2.57-1.93l.007-.008A3.6 3.6 0 0 0 24 11.688v-.004a3.608 3.608 0 0 0-.947-2.371l.002.003a8.105 8.105 0 0 0-2.529-1.917l-.048-.021a15.7 15.7 0 0 0-3.71-1.282l-.106-.019a21.212 21.212 0 0 0-4.546-.48h-.123h.006zm-3.12 7.264c-.131.119-.28.221-.442.301l-.011.005a2.916 2.916 0 0 1-.482.179l-.021.005a1.723 1.723 0 0 1-.579.099h-.024h.001H5.35l-.32 1.963H3.583l1.28-6.675h2.773l.062-.001c.36 0 .706.063 1.026.179l-.021-.007c.295.108.546.276.748.489l.001.001c.175.223.3.493.354.789l.002.011a2.932 2.932 0 0 1-.015 1.059l.003-.019a2.82 2.82 0 0 1-.142.485l.007-.019q-.086.221-.184.417q-.122.196-.27.393a2.164 2.164 0 0 1-.317.343l-.003.002zm4.172.589l.565-2.822c.024-.107.038-.229.038-.355l-.002-.078v.004a.426.426 0 0 0-.111-.283a.671.671 0 0 0-.241-.134l-.005-.001a1.388 1.388 0 0 0-.418-.062l-.051.001h.002h-1.126l-.736 3.73H9.544l1.28-6.48h1.423l-.343 1.767h1.28l.073-.001c.331 0 .653.041.961.117l-.027-.006c.249.055.466.172.641.332l-.001-.001a.84.84 0 0 1 .306.498l.001.005a1.945 1.945 0 0 1-.04.787l.003-.014l-.589 2.994zm7.902-2.184c-.04.181-.082.328-.132.473l.009-.031c-.054.159-.12.297-.201.425l.005-.008a1.812 1.812 0 0 1-.248.408l.003-.004c-.098.122-.203.23-.317.329l-.003.003c-.131.119-.28.221-.442.301l-.011.005a2.916 2.916 0 0 1-.482.179l-.021.005a1.723 1.723 0 0 1-.579.099h-.024h.001h-1.972l-.343 1.959h-1.423l1.28-6.675h2.749l.073-.001c.365 0 .716.063 1.041.18l-.022-.007c.287.104.529.272.718.488l.002.002c.19.222.325.497.378.799l.002.01a2.763 2.763 0 0 1-.04 1.076l.004-.019zm-2.7-1.547h-.978l-.513 2.749h.908c.25 0 .496-.023.734-.066l-.025.004c.204-.036.386-.109.546-.212l-.006.003c.136-.122.25-.263.339-.421l.004-.008c.103-.188.18-.407.219-.638l.002-.012a1.877 1.877 0 0 0 .036-.649l.001.009a.812.812 0 0 0-.161-.419l.001.002a1.116 1.116 0 0 0-.409-.243l-.008-.002a1.982 1.982 0 0 0-.689-.096h.003zm-11.19 0h-.978l-.515 2.749h.91c.25 0 .496-.023.734-.066l-.025.004c.204-.036.386-.109.546-.212l-.006.003c.136-.122.25-.263.339-.421l.004-.008c.103-.188.18-.407.219-.638l.002-.012a1.877 1.877 0 0 0 .036-.649l.001.009a.812.812 0 0 0-.161-.419l.001.002a1.116 1.116 0 0 0-.409-.243l-.008-.002a1.982 1.982 0 0 0-.689-.096h.003z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Picture(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h32v24H0zm2 2v20h28V2zm2.012 18.054S6.092 6.008 9.181 6.008s4.442 9.908 6.926 9.908s2.298-3.871 4.116-3.871s8.08 8.01 8.08 8.01zM25 10a3 3 0 1 1 0-6a3 3 0 0 1 0 6"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PieChartOne(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 12h12c0-6.627-5.373-12-12-12z"/><path fill="currentColor" d="M10.5 13.5V.105C4.552.873.002 5.905 0 12c0 6.627 5.373 12 12 12c6.095-.003 11.127-4.552 11.889-10.44l.006-.06z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PieChartTwo(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 11.981C.004 5.612 4.965.403 11.233.002L11.268 0v11.981c0 .139.038.269.105.379l-.002-.003l5.989 10.378A11.831 11.831 0 0 1 12 23.999c-6.627 0-12-5.373-12-12zm13.318.752H24a12.006 12.006 0 0 1-5.291 9.231l-.043.027zm-.548-1.503V0c6.04.388 10.842 5.19 11.229 11.194l.002.035z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pills(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.294 0A5.296 5.296 0 0 0 0 5.294v13.411a5.294 5.294 0 1 0 10.588 0V5.294A5.308 5.308 0 0 0 5.295 0zm4.141 11.423H1.153V5.294a4.142 4.142 0 0 1 8.284 0zm12.802 3.16a5.689 5.689 0 0 0-9.265 6.523l-.015-.028zm-8.667 7.371a5.689 5.689 0 0 0 9.265-6.523l.015.028zM18.361.801h.005a3.68 3.68 0 0 1 2.865 5.99l.005-.007L17.187.993c.35-.121.752-.192 1.171-.193h.001zm0-.629h-.006c-.778 0-1.507.208-2.136.571l.021-.011l4.916 7.03A4.31 4.31 0 0 0 18.365.168h-.01h.001zm-2.862 2.007l4.049 5.791a3.68 3.68 0 0 1-4.044-5.798zm.08-.981a4.306 4.306 0 0 0 4.93 7.019l-.021.011z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pinboard(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.352 14.585L8.843 19.2l.72-4.062L3.428 7.57L0 7.752L7.58-.001v2.953l7.214 6.647l4.513-1.105l-4.689 4.982L24 23.999z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pingdom(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m14.874 19.213l9.898-4.228L16.005 24a17.875 17.875 0 0 0-1.174-4.902l.042.121zm14.974-8.31a54.969 54.969 0 0 0-19.153 2.494l.39-.106a18.115 18.115 0 0 1 3.105 4.399l.047.104c5.027-2.798 11.996-6.371 15.61-6.879zm-15.638.342l.257-.075a48.582 48.582 0 0 0 5.943-2.186l-.318.126a20.636 20.636 0 0 0 5.217-2.978l-.045.034a11.262 11.262 0 0 0 3.601-5.048l.024-.079c.139-.299.253-.647.327-1.01l.005-.031C26.704 3.43 16.734 3.79 16.734 3.79l2.351-2.156A32.707 32.707 0 0 0-.095 7.892L0 7.826c3.93.43 7.413 2.07 10.138 4.533l-.016-.014c1.351-.418 2.736-.716 4.087-1.093v-.007z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pinterest(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 8.61v-.051C0 7.521.199 6.53.56 5.621l-.019.054A7.56 7.56 0 0 1 2.038 3.27l-.005.005A10.201 10.201 0 0 1 6.822.391l.071-.016A11.166 11.166 0 0 1 9.797 0h.011h-.001h.073c1.516 0 2.95.355 4.222.985L14.046.96a8.068 8.068 0 0 1 3.169 2.762l.018.029a7.209 7.209 0 0 1 1.226 4.041l-.001.103V7.89v.016c0 .954-.1 1.885-.29 2.782l.015-.087a10.897 10.897 0 0 1-.894 2.617l.029-.064a8.957 8.957 0 0 1-1.447 2.168l.007-.008a6.289 6.289 0 0 1-2.051 1.47l-.04.016a6.493 6.493 0 0 1-2.65.556h-.08h.004a4.47 4.47 0 0 1-1.973-.473l.026.012a2.912 2.912 0 0 1-1.377-1.255l-.008-.015q-.144.56-.4 1.622t-.339 1.37q-.08.31-.296 1.024c-.112.4-.24.74-.393 1.066l.018-.042l-.462.902c-.224.437-.444.804-.686 1.154l.022-.034q-.361.526-.894 1.247l-.202.072l-.13-.144q-.216-2.265-.216-2.711c.013-1.062.125-2.089.328-3.083l-.018.105q.31-1.651.96-4.146t.75-2.928A5.45 5.45 0 0 1 5.31 8.65v.011c0-.851.279-1.636.75-2.27l-.007.01a2.26 2.26 0 0 1 1.904-1.053c.027-.002.058-.002.09-.002c.511 0 .969.226 1.279.584l.002.002c.306.378.492.864.492 1.394l-.002.089v-.004a9.12 9.12 0 0 1-.656 2.816l.022-.061a8.86 8.86 0 0 0-.632 2.665l-.002.032l-.001.048c0 .578.25 1.098.648 1.457l.002.002c.399.371.935.599 1.525.599l.05-.001h-.002h.024c.53 0 1.028-.134 1.464-.369l-.016.008a3.348 3.348 0 0 0 1.126-.972l.006-.009c.292-.396.561-.846.787-1.32l.021-.05c.219-.455.408-.986.537-1.54l.011-.053c.106-.425.208-.968.28-1.52l.009-.08c.058-.429.092-.927.094-1.433v-.002a4.898 4.898 0 0 0-1.575-3.884l-.003-.003a5.965 5.965 0 0 0-4.131-1.392h.013A6.636 6.636 0 0 0 4.6 4.218l.002-.002a6.29 6.29 0 0 0-1.933 4.747v-.01v.012c0 .433.066.85.188 1.243l-.008-.03c.102.357.235.669.401.958l-.011-.021q.209.346.39.656c.088.125.151.273.179.433l.001.007a3.513 3.513 0 0 1-.224 1.077l.008-.024q-.216.649-.534.649q-.029 0-.245-.043a2.967 2.967 0 0 1-1.304-.806l-.001-.002a4.207 4.207 0 0 1-.87-1.332l-.01-.028a8.45 8.45 0 0 1-.459-1.5l-.01-.058a7.548 7.548 0 0 1-.159-1.533z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Plane(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m24.794 16.522l-.281-2.748l-10.191-5.131s.091-1.742 0-4.31a7.092 7.092 0 0 0-1.839-4.339l.005.006h-.182a7.075 7.075 0 0 0-1.834 4.312l-.001.021c-.091 2.567 0 4.31 0 4.31L.281 13.774L0 16.522l6.889-2.074l3.491-.582a23.316 23.316 0 0 0 .733 7.143l-.036-.162l-2.728 1.095v1.798l3.52-.8c.155.312.3.566.456.812l-.021-.035v.282c.032-.046.062-.096.093-.143c.032.046.061.096.094.143v-.282c.135-.21.28-.464.412-.726l.023-.051l3.52.8v-1.798l-2.728-1.095c.463-1.733.728-3.723.728-5.774c0-.425-.011-.847-.034-1.266l.003.058l3.492.582l6.888 2.074z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlaneTicket(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.74 9.901l-1.188 1.188a.288.288 0 0 1-.239.072h-.001l-2.688-.624l-.084.084l1.968 1.032a.245.245 0 0 1 .062.392l-.005.005l-1.197 1.197l-.005.005a.239.239 0 0 1-.208.067h-.001l-.612-.084l.672.384a.185.185 0 0 1 .096.097l-.001-.001l.373.661l-.086-.614a.24.24 0 0 1 .067-.211l.005-.005l1.187-1.187l.005-.005a.239.239 0 0 1 .21-.067h.001a.214.214 0 0 1 .18.134l-.001-.001l1.034 1.97l.084-.084l-.626-2.69a.286.286 0 0 1 .072-.239l1.188-1.188a.219.219 0 0 0 .002-.286v-.001c-.044-.07-.192-.071-.264.001z"/><path fill="currentColor" d="m13.908 1.213l-1.392 1.392a1.07 1.07 0 0 1-1.514 1.514l-9.79 9.79l-.003.003a1.71 1.71 0 0 0 0 2.417l.004.004l6.456 6.456l.003.003a1.71 1.71 0 0 0 2.417 0l.004-.004l9.79-9.79a1.085 1.085 0 0 1 .06-1.451a1.07 1.07 0 0 1 1.448-.062l.003.003l1.392-1.392l.003-.003a1.71 1.71 0 0 0 0-2.417l-.004-.004l-6.456-6.456l-.003-.003a1.71 1.71 0 0 0-2.417 0l-.002.002zM6.744 12.984l1.92-1.92a.17.17 0 0 1 .248-.009a.17.17 0 0 1-.009.248l-1.92 1.92l-.007.007a.16.16 0 0 1-.227 0l-.007-.008a.17.17 0 0 1 .001-.24zm2.7-1.452a.17.17 0 0 1 .248-.009a.17.17 0 0 1-.009.248l-1.021 1.021a.17.17 0 0 1-.248.009a.17.17 0 0 1 .009-.248zm-1.719 7.911c.099.099.099.261 0 .36s-.261.099-.36 0l-3.168-3.168c-.099-.099-.099-.261 0-.36s.261-.099.36 0zm.06-1.308c.099.099.099.261 0 .36s-.261.099-.36 0l-1.92-1.92c-.099-.099-.099-.261 0-.36s.261-.099.36 0zm5.062-2.878l-1.915 1.915a.17.17 0 0 1-.248.009a.17.17 0 0 1 .009-.248l1.92-1.92l.007-.007a.16.16 0 0 1 .227 0l.007.008a.183.183 0 0 1 0 .239zm-1.531.211a.17.17 0 0 1-.24-.24l1.021-1.021a.17.17 0 0 1 .24.24zm3.084-4.907l-1.092 1.092l.624 2.688a.286.286 0 0 1-.072.239l-.432.432l-.005.005a.239.239 0 0 1-.21.067h-.001a.214.214 0 0 1-.18-.134l.001.001l-1.03-1.966l-.864.864l.084.636a.24.24 0 0 1-.067.211l-.005.005l-.324.324l-.005.005a.239.239 0 0 1-.21.067h-.001a.214.214 0 0 1-.18-.134l.001.001l-.504-.912l-.914-.506a.319.319 0 0 1-.133-.181l.001.002a.24.24 0 0 1 .067-.211l.005-.005l.324-.324l.005-.005a.239.239 0 0 1 .21-.067h.001l.636.084l.864-.864l-1.992-1.056a.245.245 0 0 1-.062-.392l.005-.005l.431-.431a.288.288 0 0 1 .239-.072h.001l2.688.624l1.092-1.092a.717.717 0 0 1 1.008 0a.729.729 0 0 1-.004 1.009zm4.608-.456c.099.099.099.261 0 .36s-.261.099-.36 0l-.12-.12c-.099-.099-.099-.261 0-.36s.261-.099.36 0zm-.84-.84c.099.099.099.261 0 .36s-.261.099-.36 0l-.12-.12c-.099-.099-.099-.261 0-.36s.261-.099.36 0zm-.84-.84c.099.099.099.261 0 .36s-.261.099-.36 0l-.121-.121c-.099-.099-.099-.261 0-.36s.261-.099.36 0zm-.839-.84c.099.099.099.261 0 .36s-.261.099-.36 0l-.121-.121c-.099-.099-.099-.261 0-.36s.261-.099.36 0zm-.84-.84c.099.099.099.261 0 .36s-.261.099-.36 0l-.121-.121c-.099-.099-.099-.261 0-.36s.261-.099.36 0zm-.84-.84c.099.099.099.261 0 .36s-.261.099-.36 0l-.121-.121c-.099-.099-.099-.261 0-.36s.261-.099.36 0zm-.84-.84c.099.099.099.261 0 .36s-.261.099-.36 0l-.121-.121c-.099-.099-.099-.261 0-.36s.261-.099.36 0zm.84-.983l.816-.816l-.289-.289l.192-.192l.804.804l-.192.192l-.288-.288l-.816.816zm.804.804l1.008-1.008l.227.227l-1.008 1.008zm5.601 3.585l-.192.192l-.288-.288l-.815.815l-.227-.227l.816-.816l-.288-.288l.192-.192zm-1.056-1.055l-.18.18l-.442-.442l-.216.216l.397.397l-.18.18l-.397-.397l-.24.24l.454.454l-.18.18l-.681-.681l1.008-1.008l.672.672zm-1.56-1.561l-.454.454l.777-.13l.252.252l-.732.108l-.228.948l-.251-.251l.168-.648l-.302.038l-.252.252l-.227-.227l1.008-1.008l.227.227zm-.732-.3l-.288-.096a.364.364 0 0 0-.093-.357l-.002-.002c-.144-.144-.372-.157-.588.06s-.204.444-.06.588l.002.002a.364.364 0 0 0 .355.093l.002-.001l.096.288a.68.68 0 0 1-.636-.181l-.025-.026l.001.001l-.015-.014a.677.677 0 0 1 0-.958l.04-.037l-.002.002a.67.67 0 0 1 .988-.046l.022.023l-.001-.001a.619.619 0 0 1 .205.66l-.001.003z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Play(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.001 1.165v21.669a1.275 1.275 0 0 0 1.891 1.017l-.006.003l21.442-10.8a1.172 1.172 0 0 0 .007-2.113l-.007-.003L1.886.138A1.273 1.273 0 0 0 .003 1.162v.004z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayList(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 24.001v-3.165h27.693v3.165zm0-7.912v-3.165h27.693v3.165zm0-7.055V.461A.505.505 0 0 1 .748.059L.745.058l8.483 4.273a.464.464 0 0 1 .003.836l-.003.001L.745 9.437a.534.534 0 0 1-.242.059h-.02A.484.484 0 0 1 0 9.035zm12.659-2.44V3.429h15.033v3.165z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlayerSettings(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.521 10.145a1.32 1.32 0 0 1-.97-.811l-.003-.009l-.609-1.482a1.297 1.297 0 0 1 .103-1.269l-.003.005l.988-1.318a.96.96 0 0 0-.049-1.164l.001.002l-1.082-1.085a.963.963 0 0 0-1.167-.043l.003-.002l-1.316.988a1.295 1.295 0 0 1-1.272.098l.008.003l-1.482-.609a1.308 1.308 0 0 1-.819-.963l-.001-.008l-.229-1.631a1.017 1.017 0 0 0-.841-.822l-.006-.001s-.26-.025-.776-.025s-.769.026-.769.026a1.017 1.017 0 0 0-.853.815l-.001.006l-.233 1.629a1.315 1.315 0 0 1-.811.968l-.009.003l-1.482.609a1.298 1.298 0 0 1-1.269-.104l.005.003l-1.318-.988a.963.963 0 0 0-1.165.046l.002-.001l-1.079 1.087a.958.958 0 0 0-.044 1.165l-.002-.003l.988 1.318a1.309 1.309 0 0 1 .099 1.272l.003-.008l-.611 1.482c-.162.419-.522.73-.963.819l-.008.001l-1.631.23a1.017 1.017 0 0 0-.817.849l-.001.006s-.029.256-.029.769s.027.77.027.77c.064.432.393.773.814.852l.006.001l1.629.232c.449.091.81.4.97.811l.003.009l.611 1.482a1.291 1.291 0 0 1-.105 1.269l.003-.005l-.988 1.318a.887.887 0 0 0-.13.999l-.002-.005c.175.217.677.752.678.754l.375.344c.204.188 1.042.449 1.372.203l1.317-.988a1.29 1.29 0 0 1 1.272-.097l-.008-.003l1.482.609c.42.161.731.522.819.963l.001.008l.229 1.626c.08.428.42.757.846.822l.006.001s.257.025.772.025s.77-.027.77-.027c.43-.067.768-.396.846-.816l.001-.006l.234-1.629c.09-.449.4-.81.811-.968l.009-.003l1.482-.611a1.291 1.291 0 0 1 1.269.105l-.005-.003l1.318.988a.962.962 0 0 0 1.166-.049l-.002.001l1.082-1.084a.966.966 0 0 0 .045-1.167l.002.003l-.988-1.318a1.29 1.29 0 0 1-.218-.722c0-.197.043-.383.121-.55l-.003.008l.609-1.482c.163-.42.524-.73.965-.819l.008-.001l1.627-.232c.428-.08.758-.42.822-.846l.001-.006s.027-.249.027-.765s-.027-.77-.027-.77a1.018 1.018 0 0 0-.818-.852l-.006-.001zM11.998 15a3 3 0 1 1 3.001-3.002v.001A3 3 0 0 1 11.998 15"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Playstation(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.646 5.077a6.515 6.515 0 0 1 .061 1.286l.001-.015c0 2.344-.007 4.688.006 7.032c.655.34 1.427.549 2.245.57h.007l.074.001a3.02 3.02 0 0 0 1.439-.362l-.016.008a3.117 3.117 0 0 0 1.143-1.096l.008-.014a4.846 4.846 0 0 0 .616-1.677l.004-.029a9.26 9.26 0 0 0 .108-2.112l.001.026a7.974 7.974 0 0 0-.692-3.171l.02.051a5.107 5.107 0 0 0-1.027-1.443l-.001-.001a6.73 6.73 0 0 0-1.298-1l-.031-.017a15.919 15.919 0 0 0-3.302-1.404l-.116-.03c-.75-.24-1.505-.47-2.262-.69c-1.323-.375-2.652-.735-4.005-.987v22.393l5.065 1.607q.006-9.414 0-18.827l-.001-.06c0-.36.112-.693.304-.967l-.004.005a.623.623 0 0 1 .6-.23l-.004-.001c.234.043.44.143.609.286l-.002-.002c.235.224.397.523.446.858l.001.008zM5.42 18.349c1.738-.621 3.475-1.249 5.216-1.866a2.867 2.867 0 0 0 .008-.357v-2.589c-2.202.777-4.4 1.563-6.602 2.342c-.537.196-1.082.37-1.608.594a9.67 9.67 0 0 0-1.795.943l.035-.022a1.633 1.633 0 0 0-.569.619l-.004.009a1.089 1.089 0 0 0-.03.854l-.003-.007c.139.322.354.589.624.785l.005.004c.47.338 1.02.601 1.613.756l.035.008c1.548.53 3.332.835 5.187.835c.117 0 .234-.001.351-.004h-.017a17.322 17.322 0 0 0 2.876-.308l-.108.017c.014-.146.006-.293.008-.439v-2.194c-.744.266-1.487.537-2.23.806a5.187 5.187 0 0 1-.837.26l-.037.007a6.627 6.627 0 0 1-1.351.139h-.033a3.6 3.6 0 0 1-1.198-.204l.025.008a.605.605 0 0 1-.344-.282l-.002-.003a.33.33 0 0 1 .109-.338h.001c.188-.165.413-.294.66-.372l.013-.004zm25.184-1.68a4.625 4.625 0 0 0-1.509-.831l-.033-.009c-.309-.098-.609-.22-.918-.314c-1.432-.449-3.079-.708-4.786-.708h-.109h.006c-.502.018-1.004.032-1.503.09c-1.52.165-2.908.484-4.228.946l.132-.04v3.04c1.829-.64 3.654-1.287 5.482-1.928a5.911 5.911 0 0 1 1.844-.288h.052h-.003l.06-.001c.412 0 .808.071 1.176.201l-.025-.008a.61.61 0 0 1 .343.276l.002.003a.38.38 0 0 1-.178.406l-.002.001a3.468 3.468 0 0 1-.99.453l-.025.006q-3.85 1.374-7.699 2.745c-.054.01-.033.076-.038.114v2.834q5.06-1.819 10.121-3.635c.665-.218 1.21-.437 1.738-.685l-.092.039a4.526 4.526 0 0 0 1.359-.959l.001-.001c.198-.209.32-.492.32-.804V17.6v.001a1.34 1.34 0 0 0-.497-.929l-.003-.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PlusA(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.2 1.2h-2.4v9.6H1.2v2.4h9.6v9.6h2.4v-9.6h9.6v-2.4h-9.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Podcast(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m10.064 11.853l.071.001a2.45 2.45 0 0 0 2.448-2.37V9.48a2.45 2.45 0 0 0-2.448-2.374l-.075.001h.004l-.071-.001a2.45 2.45 0 0 0-2.448 2.37v.004a2.45 2.45 0 0 0 2.448 2.374l.075-.001zm0-8.302A6.121 6.121 0 0 0 3.769 9.47v.01a5.715 5.715 0 0 0 1.75 4.102l.002.002a4.008 4.008 0 0 1 1.308-1.252l.018-.01a3.999 3.999 0 0 1-1.195-2.84v-.001a4.29 4.29 0 0 1 4.285-4.155l.131.002h-.006l.125-.002a4.287 4.287 0 0 1 4.285 4.147v.007a4.003 4.003 0 0 1-1.199 2.846l-.001.001a4.02 4.02 0 0 1 1.321 1.251l.009.015a5.724 5.724 0 0 0 1.76-4.111v-.001a6.123 6.123 0 0 0-6.304-5.929h.009zm0-3.55C4.505 0 0 4.245 0 9.48a9.368 9.368 0 0 0 4.754 8.055l.046.024a9.113 9.113 0 0 1-.094-1.307v-.001c.005-.309.039-.606.099-.894l-.005.031a7.524 7.524 0 0 1-2.925-5.906v-.003c0-4.259 3.666-7.712 8.188-7.712s8.188 3.453 8.188 7.712a7.52 7.52 0 0 1-2.919 5.905l-.017.013c.055.254.089.547.096.848v.005a9.472 9.472 0 0 1-.102 1.368l.006-.052c2.863-1.617 4.774-4.625 4.812-8.082v-.005c0-5.236-4.506-9.48-10.064-9.48zm0 13.039l-.098-.001a3.37 3.37 0 0 0-3.367 3.258v.005c0 1.802 1.553 7.694 3.464 7.694s3.464-5.896 3.464-7.694a3.37 3.37 0 0 0-3.47-3.263h.005z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Power(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.154 12.923v.033c0 1.542-.324 3.008-.907 4.334l.027-.069c-1.722 4.011-5.637 6.768-10.197 6.768S2.602 21.231.908 17.293l-.028-.072A10.603 10.603 0 0 1 0 12.956v-.069c0-3.595 1.729-6.785 4.4-8.786l.029-.02a1.813 1.813 0 0 1 1.388-.36l-.01-.001c.501.059.929.328 1.2.715l.004.005a1.744 1.744 0 0 1 .352 1.37l.001-.01a1.796 1.796 0 0 1-.709 1.215l-.005.003a7.394 7.394 0 0 0-2.969 5.925a7.397 7.397 0 1 0 11.844-5.912l-.019-.014a1.804 1.804 0 0 1-.713-1.209l-.001-.009a1.745 1.745 0 0 1 .356-1.363l-.003.004a1.71 1.71 0 0 1 1.203-.719l.008-.001a1.78 1.78 0 0 1 1.375.364l-.004-.003c2.702 2.021 4.432 5.213 4.432 8.809v.034v-.002zm-9.23-11.077v9.268a1.846 1.846 0 0 1-3.692 0v-.04v.002V1.83c0-.504.21-.959.547-1.282L9.78.547c.329-.338.789-.548 1.298-.548s.969.21 1.298.548c.338.324.548.778.548 1.282z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Prescription(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.76 2.852h-2.577v1.221h2.584c.655 0 1.186.531 1.186 1.186V21.6c0 .656-.532 1.187-1.187 1.187H2.401A1.187 1.187 0 0 1 1.214 21.6V5.259c0-.656.532-1.187 1.187-1.187h2.584v-1.22H2.408A2.412 2.412 0 0 0-.003 5.263v16.326a2.412 2.412 0 0 0 2.411 2.41h13.351a2.412 2.412 0 0 0 2.411-2.411V5.262a2.412 2.412 0 0 0-2.411-2.411z"/><path fill="currentColor" d="M12.605 2.225h-1.319a2.225 2.225 0 1 0-4.45 0h-1.31v3.057h7.073V2.225zm-2.258 0h-2.57v-.032a1.286 1.286 0 0 1 2.572 0v.034v-.002zm-4.4 6.287h9.458v1.17H5.947zm0 4.898h9.458v1.17H5.947zm0 4.84h9.458v1.176H5.947zM2.628 7.757h2.68v2.68h-2.68zm0 4.898h2.68v2.68h-2.68zm0 4.842h2.68v2.68h-2.68z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Preview(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m23.589 22.261l-2.102-2.101a4.978 4.978 0 0 0-3.115-7.594l-.033-.006V8.491a1.01 1.01 0 0 0-.233-.646l.001.002v-.005a.94.94 0 0 0-.06-.066l-.008-.009l-.027-.027L10.572.3a.91.91 0 0 0-.065-.059l-.026-.018a.938.938 0 0 0-.05-.038l-.025-.018l-.054-.034l-.023-.012q-.034-.02-.075-.037l-.032-.013l-.051-.018l-.036-.011l-.058-.015l-.028-.006a.81.81 0 0 0-.086-.013H1.015A1.019 1.019 0 0 0 0 1.02v20.377a1.02 1.02 0 0 0 1.015 1.019h16.31a4.924 4.924 0 0 0 2.74-.827l-.018.011l2.102 2.102a1.019 1.019 0 0 0 1.439-1.44l.001.001zm-3.325-4.827a2.943 2.943 0 1 1-5.887.001a2.943 2.943 0 0 1 5.887-.002zM10.868 3.478l3.993 3.994h-3.993zm-8.83-1.44h6.793v6.453c0 .563.456 1.019 1.019 1.019h6.453v3.05c-2.278.487-3.962 2.483-3.962 4.873c0 1.109.362 2.133.975 2.96l-.01-.013H2.037z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Print(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.909 4.363h-3.272V1.091C19.637.488 19.148 0 18.546 0H5.455c-.603 0-1.091.489-1.091 1.091v3.272H1.092c-.603 0-1.091.489-1.091 1.091v13.091c0 .603.489 1.091 1.091 1.091h3.272v3.273c0 .603.489 1.091 1.091 1.091h13.091c.603 0 1.091-.489 1.091-1.091v-3.273h3.273c.603 0 1.091-.489 1.091-1.091V5.454c0-.603-.489-1.091-1.091-1.091zM6.546 2.182h10.908v2.182H6.546zm10.908 19.636H6.546v-6.546h10.908zm4.363-4.363h-2.181v-3.273c0-.603-.489-1.091-1.091-1.091H5.454c-.603 0-1.091.489-1.091 1.091v3.273H2.181V6.546h19.637z"/><path fill="currentColor" d="M14.182 17.454H9.788a1.091 1.091 0 0 0 0 2.182h.032h-.002h4.394a1.091 1.091 0 0 0 0-2.182h-.032zm3.592-8.407a1.092 1.092 0 1 0 .772-.319h-.001a1.094 1.094 0 0 0-.771.32z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ProductHunt(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.387 10.21l.001-.045a1.69 1.69 0 0 0-.532-1.234l-.001-.001a1.755 1.755 0 0 0-1.237-.508H10.21v3.581h3.402c.489 0 .93-.204 1.243-.531l.001-.001c.328-.314.532-.755.532-1.244v-.018v.001zM24 12v-.077c0-2.185-.602-4.229-1.65-5.976l.029.053a12.228 12.228 0 0 0-4.323-4.348L18 1.621A11.5 11.5 0 0 0 12.077 0h-.081H12h-.077C9.738 0 7.694.602 5.947 1.65L6 1.621a12.228 12.228 0 0 0-4.348 4.323L1.621 6A11.5 11.5 0 0 0 0 11.923v.081V12v.077c0 2.185.602 4.229 1.65 5.976L1.621 18a12.228 12.228 0 0 0 4.323 4.348l.056.031A11.5 11.5 0 0 0 11.923 24h.081H12h.077c2.185 0 4.229-.602 5.976-1.65l-.053.029a12.228 12.228 0 0 0 4.348-4.323l.031-.056A11.5 11.5 0 0 0 24 12.077v-.081zm-6.194-1.79v.063a4.044 4.044 0 0 1-1.233 2.912l-.001.001a4.045 4.045 0 0 1-2.913 1.234l-.066-.001h.003h-3.387V18H7.79V6h5.806l.064-.001c1.142 0 2.174.472 2.911 1.232l.001.001a4.045 4.045 0 0 1 1.234 2.913l-.001.068z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PropellerFour(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 13.123a1.125 1.125 0 1 1 0-2.25a1.125 1.125 0 0 1 0 2.25m11.992-.54a4.646 4.646 0 0 0-3.653-4.288l-.031-.006a8.56 8.56 0 0 0-6.541 1.839l.016-.013a2.588 2.588 0 0 0-.563-.4l-.014-.007c1.625-1.535 4.339-3.075 4.715-5.118c.407-2.205-2.524-4.758-5.338-4.584a4.647 4.647 0 0 0-4.288 3.653l-.006.031a8.56 8.56 0 0 0 1.839 6.541l-.013-.016a2.622 2.622 0 0 0-.4.563l-.007.014C8.172 9.167 6.633 6.453 4.589 6.076C2.384 5.67-.169 8.6.005 11.414a4.647 4.647 0 0 0 3.654 4.288l.03.005a8.566 8.566 0 0 0 6.541-1.84l-.016.013c.167.156.357.291.563.4l.014.007c-1.625 1.536-4.34 3.075-4.716 5.119c-.406 2.205 2.524 4.758 5.339 4.584a4.646 4.646 0 0 0 4.288-3.653l.006-.031a8.562 8.562 0 0 0-1.84-6.541l.013.016c.156-.167.292-.357.4-.563l.007-.014c1.535 1.625 3.075 4.339 5.119 4.716c2.205.407 4.758-2.524 4.584-5.339z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PropellerOne(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.095 13.221a1.22 1.22 0 1 1 0-2.44a1.22 1.22 0 0 1 0 2.44m8.028-5.124c-1.731 1.473-3.041 2.684-5.262 3.377a2.823 2.823 0 0 0-1.389-1.92l-.014-.007c.866-2.231 1.464-5.216.994-6.625C11.754.827 9.485-.843 5.12.453C.331 1.876-.113 5.998 2.701 6.999c2.144.763 3.852 1.291 5.563 2.874a2.797 2.797 0 0 0-.987 2.126c0 .016.005.031.005.047c-2.364.367-5.242 1.341-6.226 2.451c-1.464 1.652-1.778 4.452 1.528 7.584c3.626 3.435 7.418 1.759 6.878-1.178c-.412-2.239-.809-3.982-.293-6.254c.275.103.593.165.925.168h.002a2.772 2.772 0 0 0 1.457-.414l-.012.007c1.499 1.864 3.784 3.874 5.238 4.171c2.164.443 4.745-.686 5.804-5.115c1.162-4.859-2.186-7.305-4.459-5.369z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PropellerThree(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.982 13.031a1.2 1.2 0 1 1 0-2.398a1.2 1.2 0 0 1 0 2.398m9.123-3.482a28.177 28.177 0 0 0-6.684 1.534l.196-.062a2.767 2.767 0 0 0-3.233-1.893l.018-.003a41.152 41.152 0 0 0-1.93-6.002l.101.279C4.27.124 2.285-.5.86.334c-1.608.941-.714 2.914.577 4.742a28.428 28.428 0 0 0 4.477 4.845l.038.032a2.738 2.738 0 0 0-.024 3.729l-.002-.002a42.157 42.157 0 0 0-3.982 4.364l-.067.087c-2.184 2.766-1.735 4.797-.301 5.615c1.619.923 2.878-.838 3.818-2.871c.837-1.825 1.519-3.952 1.94-6.171l.028-.176c.184.046.396.073.614.074h.001a2.765 2.765 0 0 0 2.618-1.899l.006-.02c1.633.506 3.62.952 5.655 1.252l.227.027c3.487.509 5.023-.896 5.013-2.547c-.01-1.863-2.166-2.075-4.398-1.871z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func PropellerTwo(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.809 13.832a1.233 1.233 0 1 1 .001-2.465a1.233 1.233 0 0 1-.001 2.465m7.005-3.266a12.383 12.383 0 0 1-4.071 2.01l-.088.021a2.84 2.84 0 0 0-1.859-2.656l-.02-.006c.112-.772.176-1.663.176-2.57a18.94 18.94 0 0 0-.437-4.055l.023.125C9.626.197 6.984-.585 4.789.396c-3.782 1.681-3.018 5.753-.241 7.159a12.507 12.507 0 0 1 3.85 2.595l-.001-.001a2.84 2.84 0 0 0-1.43 2.455c.003.176.021.346.053.511l-.003-.018a18.952 18.952 0 0 0-5.519 3.469l.016-.015c-2.349 2.41-1.704 5.088.241 6.503c3.347 2.435 6.488-.263 6.324-3.367a12.487 12.487 0 0 1 .341-4.717l-.02.087c.406.244.894.389 1.417.392h.001a2.812 2.812 0 0 0 1.827-.681l-.004.003a18.695 18.695 0 0 0 5.616 3.007l.133.038c3.261.829 5.256-1.068 5.512-3.46c.435-4.116-3.473-5.491-6.078-3.789z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Pulse(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8.606 11.366l2.16-7.008l4.619 10.497l1.022-3.546h11.12a9.91 9.91 0 0 0 .423-2.895c0-.42-.026-.833-.075-1.239l.005.049c-.281-2.127-1.432-6.101-6.094-7.014a10.518 10.518 0 0 0-1.948-.208h-.146a7.957 7.957 0 0 0-5.722 2.563l-.005.005A7.974 7.974 0 0 0 8.245.002h-.147c-.699.011-1.373.086-2.027.22l.07-.012C2.797.934.289 3.754.048 7.2l-.001.024a10.534 10.534 0 0 0 .438 4.216l-.02-.075z"/><path fill="currentColor" d="m17.647 12.955l-1.935 6.742l-4.674-10.619l-1.209 3.936H1.087C2.576 16.31 6.047 20.446 13.961 24c8.015-3.6 11.49-7.741 12.96-11.04z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Python(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.727 0a16.43 16.43 0 0 0-2.834.248l.098-.014C6.568.662 6.129 1.558 6.129 3.21v2.182h5.726v.727H3.981l-.066-.001A3.576 3.576 0 0 0 .408 8.999l-.004.023c-.256.872-.403 1.874-.403 2.91s.147 2.038.422 2.985l-.019-.076c.407 1.695 1.379 2.902 3.04 2.902h1.969v-2.616a3.64 3.64 0 0 1 3.574-3.557h5.722a2.885 2.885 0 0 0 2.863-2.885v-.026v.001v-5.452A3.204 3.204 0 0 0 14.724.233L14.71.232a17.319 17.319 0 0 0-2.879-.234h-.107h.005zM8.631 1.755h.017a1.091 1.091 0 1 1-1.091 1.094v-.008c0-.596.48-1.08 1.074-1.086"/><path fill="currentColor" d="M18.287 6.119v2.542a3.672 3.672 0 0 1-3.572 3.63H8.991A2.922 2.922 0 0 0 6.129 15.2v5.453c0 1.551 1.349 2.464 2.862 2.91c.855.277 1.839.437 2.86.437s2.005-.16 2.927-.456l-.068.019c1.44-.417 2.862-1.258 2.862-2.91v-2.184h-5.719v-.727h8.582c1.664 0 2.284-1.161 2.863-2.902c.28-.87.441-1.871.441-2.91s-.161-2.04-.46-2.979l.019.07c-.411-1.656-1.2-2.902-2.863-2.902zm-3.216 13.807h.017a1.091 1.091 0 1 1-1.091 1.091v-.011c0-.595.48-1.077 1.074-1.08"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Qrcode(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.455 17.455h2.182v2.182h-2.182z"/><path fill="currentColor" d="M13.091 24H24V13.091H13.091zm2.181-8.727h6.545v6.529h-6.545zm2.183-10.909h2.182v2.182h-2.182zM8.727 15.272H6.545V13.09H0v10.909h2.182v-6.545h2.182v2.182h6.545v-6.545H8.727z"/><path fill="currentColor" d="M4.364 21.818h2.182V24H4.364zm4.363 0h2.182V24H8.727zM4.364 4.364h2.182v2.182H4.364z"/><path fill="currentColor" d="M0 10.909h10.909V0H0zm2.182-8.727h6.545v6.545H2.182zM13.091 0v10.909H24V0zm8.727 8.727h-6.545V2.182h6.545z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Question(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.403 18.751v4.499c-.01.41-.34.74-.748.75H6.159a.768.768 0 0 1-.749-.748v-4.5c.01-.41.34-.739.749-.749h4.5c.41.01.74.34.75.749v.001zm5.923-11.247a6.306 6.306 0 0 1-.962 3.354l.015-.026a5.462 5.462 0 0 1-1.021 1.108l-.01.008c-.321.282-.672.55-1.042.794l-.036.022q-.413.253-1.144.665a3.71 3.71 0 0 0-1.275 1.204l-.009.014a2.535 2.535 0 0 0-.515 1.243l-.001.012a.978.978 0 0 1-.226.611l.001-.002a.652.652 0 0 1-.524.29h-4.5a.563.563 0 0 1-.479-.343l-.001-.004a1.394 1.394 0 0 1-.197-.702v-.845a4.356 4.356 0 0 1 1.219-2.935l-.001.001A7.945 7.945 0 0 1 9.251 9.96l.048-.02a4.627 4.627 0 0 0 1.574-1.049l.001-.001a2.094 2.094 0 0 0 .469-1.429v.005a1.695 1.695 0 0 0-.863-1.382l-.009-.004a3.436 3.436 0 0 0-2.018-.599h.003a3.53 3.53 0 0 0-2.039.552l.014-.009A12.825 12.825 0 0 0 4.45 8.149l-.025.035a.73.73 0 0 1-.581.3a.897.897 0 0 1-.472-.152l.003.002L.301 5.991a.732.732 0 0 1-.29-.464L.01 5.523a.747.747 0 0 1 .105-.527l-.002.003C1.77 2 4.912.003 8.522.003c.103 0 .205.002.307.005h-.015a8.362 8.362 0 0 1 3.074.602l-.057-.02a10.2 10.2 0 0 1 2.757 1.571l-.02-.016a7.838 7.838 0 0 1 1.966 2.349l.02.041c.483.857.768 1.881.769 2.971z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Quora(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23.254 19.014h-.053a3.884 3.884 0 0 1-.222 1.74l.009-.027a4.502 4.502 0 0 1-1.336 2.194l-.004.004a4.693 4.693 0 0 1-3.226 1.069h.011a4.832 4.832 0 0 1-3.117-1.025l.011.008a7.232 7.232 0 0 1-1.963-2.262l-.019-.037a9.916 9.916 0 0 1-2.732.375c-1.249 0-2.446-.226-3.55-.639l.07.023a10.596 10.596 0 0 1-5.017-3.642l-.019-.027A9.894 9.894 0 0 1 .004 10.66l.001-.111v.006l-.001-.131c0-1.997.578-3.859 1.575-5.427l-.024.041A10.79 10.79 0 0 1 5.44 1.366l.055-.028a10.402 10.402 0 0 1 5.137-1.339h.005c1.896.002 3.675.498 5.217 1.367l-.054-.028a10.467 10.467 0 0 1 3.885 3.6l.025.042a10.567 10.567 0 0 1 1.514 5.482v.093v-.005v.068c0 1.691-.395 3.289-1.099 4.707l.028-.062a10.94 10.94 0 0 1-2.872 3.464l-.023.017c.331.534.72.992 1.169 1.384l.007.006c.416.301.937.481 1.499.482l.076.002c.48 0 .912-.206 1.213-.534l.001-.001c.246-.289.403-.659.426-1.066v-.005zm-7.713-3.106c.583-1.426.922-3.08.922-4.814c0-.191-.004-.381-.012-.569l.001.027q0-4.285-1.42-6.374a4.969 4.969 0 0 0-4.433-2.091l.016-.001a4.894 4.894 0 0 0-4.382 2.076l-.01.016q-1.395 2.088-1.368 6.35t1.395 6.347a4.834 4.834 0 0 0 4.382 2.091l-.017.001l.107.001c.594 0 1.165-.099 1.698-.28l-.037.011a10.732 10.732 0 0 0-1.459-2.263l.013.016a3.172 3.172 0 0 0-2.304-1.027h-.002l-.053-.001c-.351 0-.683.08-.979.223l.013-.006l-.697-1.285a5.469 5.469 0 0 1 3.707-1.231h-.011a5.449 5.449 0 0 1 3.076.817l-.023-.013a7.27 7.27 0 0 1 1.859 1.951z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuoteAleft(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.948 6.27a5.801 5.801 0 0 1 1.904-4.551l.005-.004A6.824 6.824 0 0 1 23.619.001h-.011a7.208 7.208 0 0 1 5.676 2.298l.004.004a8.796 8.796 0 0 1 2.057 6.143l.001-.021a16.241 16.241 0 0 1-1.265 6.769l.04-.107a16.206 16.206 0 0 1-2.95 4.567l.011-.012a14.879 14.879 0 0 1-3.504 2.803l-.072.038a21.8 21.8 0 0 1-3.076 1.467l-.156.053l-3.134-5.29a8.091 8.091 0 0 0 3.317-2.482l.013-.017a6.874 6.874 0 0 0 1.565-3.641l.003-.032a4.842 4.842 0 0 1-3.571-1.708l-.006-.007a6.339 6.339 0 0 1-1.614-4.57l-.001.015zM0 6.27a5.807 5.807 0 0 1 1.905-4.551l.005-.004A6.819 6.819 0 0 1 6.672.001h-.011a7.204 7.204 0 0 1 5.676 2.298l.004.005a8.792 8.792 0 0 1 2.057 6.144l.001-.021a16.241 16.241 0 0 1-1.265 6.769l.04-.107a16.203 16.203 0 0 1-2.949 4.567l.011-.012a14.884 14.884 0 0 1-3.505 2.803l-.072.038a21.69 21.69 0 0 1-3.076 1.467l-.156.053l-3.133-5.29a8.093 8.093 0 0 0 3.317-2.484l.013-.017a6.874 6.874 0 0 0 1.565-3.641l.003-.032a4.836 4.836 0 0 1-3.569-1.708l-.006-.007A6.335 6.335 0 0 1 0 6.257l-.001.015z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuoteAright(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.4 17.73a5.801 5.801 0 0 1-1.904 4.55l-.005.004a6.824 6.824 0 0 1-4.762 1.714h.011A7.208 7.208 0 0 1 2.064 21.7l-.004-.004a8.796 8.796 0 0 1-2.057-6.143l-.001.021a16.241 16.241 0 0 1 1.265-6.769l-.04.107a16.187 16.187 0 0 1 2.951-4.567l-.011.012a14.879 14.879 0 0 1 3.504-2.803l.072-.038A21.8 21.8 0 0 1 10.819.049l.156-.053l3.134 5.29a8.091 8.091 0 0 0-3.317 2.482l-.013.017a6.874 6.874 0 0 0-1.565 3.641l-.003.032a4.842 4.842 0 0 1 3.571 1.708l.006.007a6.339 6.339 0 0 1 1.614 4.57l.001-.015zm16.948 0a5.807 5.807 0 0 1-1.905 4.551l-.005.004a6.819 6.819 0 0 1-4.762 1.714h.011a7.204 7.204 0 0 1-5.676-2.298l-.004-.005a8.792 8.792 0 0 1-2.057-6.144l-.001.021a16.241 16.241 0 0 1 1.265-6.769l-.04.107a16.203 16.203 0 0 1 2.949-4.567l-.011.012a14.884 14.884 0 0 1 3.505-2.803l.072-.038A21.69 21.69 0 0 1 27.765.048l.156-.053l3.133 5.29a8.093 8.093 0 0 0-3.317 2.484l-.013.017a6.874 6.874 0 0 0-1.565 3.641l-.003.032a4.836 4.836 0 0 1 3.569 1.708l.006.007a6.335 6.335 0 0 1 1.617 4.569l.001-.015z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuoteLeft(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M31.2 0H24l-4.8 9.6V24h14.4V9.6h-7.2zM12 0H4.8L0 9.6V24h14.4V9.6H7.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func QuoteRight(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.4 24h7.2l4.8-9.6V0H0v14.4h7.2zm19.2 0h7.2l4.8-9.6V0H19.2v14.4h7.2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RadioBtnActive(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 12C0 5.373 5.373 0 12 0s12 5.373 12 12s-5.373 12-12 12C5.376 23.992.008 18.624 0 12.001zm2.4 0c0 5.302 4.298 9.6 9.6 9.6s9.6-4.298 9.6-9.6s-4.298-9.6-9.6-9.6c-5.299.006-9.594 4.301-9.6 9.599zm4 0a5.6 5.6 0 1 1 11.2 0a5.6 5.6 0 0 1-11.2 0"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RadioBtnPassive(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 24C5.373 24 0 18.627 0 12S5.373 0 12 0s12 5.373 12 12c-.008 6.624-5.376 11.992-11.999 12zm0-21.6c-5.302 0-9.6 4.298-9.6 9.6s4.298 9.6 9.6 9.6s9.6-4.298 9.6-9.6c-.006-5.299-4.301-9.594-9.599-9.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rage(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.316C6.099 1.316 1.316 6.099 1.316 12S6.099 22.684 12 22.684S22.684 17.901 22.684 12c-.012-5.896-4.788-10.672-10.683-10.684zm0 22.297C5.586 23.613.387 18.414.387 12S5.586.387 12 .387S23.613 5.586 23.613 12v.015c0 6.405-5.192 11.597-11.597 11.597z"/><path fill="currentColor" d="M12 24C5.386 23.966.034 18.614 0 12.003V12C0 5.373 5.373 0 12 0s12 5.373 12 12s-5.373 12-12 12M12 .774C5.8.774.774 5.8.774 12S5.8 23.226 12 23.226S23.226 18.2 23.226 12C23.222 5.802 18.198.779 12.001.774zm0 22.374C5.886 23.148.929 18.191.929 12.077S5.886 1.006 12 1.006s11.071 4.957 11.071 11.071S18.114 23.148 12 23.148m0-21.445C6.313 1.703 1.703 6.313 1.703 12S6.313 22.297 12 22.297S22.297 17.687 22.297 12c-.026-5.676-4.62-10.271-10.294-10.297z"/><path fill="currentColor" d="M6.89 18.968a5.11 5.11 0 0 1 10.22 0H6.891zM9.677 9.91a2.09 2.09 0 1 1-2.09-2.09h.009c1.15 0 2.082.932 2.082 2.082v.009zm8.904 0v.009c0 1.15-.932 2.082-2.082 2.082h-.009a2.114 2.114 0 0 1-2.09-2.089v-.01c0-1.15.932-2.082 2.082-2.082h.009a2.132 2.132 0 0 1 2.09 2.088zM9.987 7.355a.279.279 0 0 1-.232-.078L7.2 5.729a1.05 1.05 0 0 1-.229-.304l-.003-.006a.462.462 0 0 1 .078-.311l-.001.002a1.05 1.05 0 0 1 .304-.229l.006-.003a.462.462 0 0 1 .311.078l-.002-.001l2.555 1.548c.094.087.172.189.229.304l.003.006a.462.462 0 0 1-.078.311l.001-.002a.426.426 0 0 1-.379.232h-.008z"/><path fill="currentColor" d="M9.987 7.742a1.085 1.085 0 0 1-.469-.158l.005.003l-2.555-1.548a.987.987 0 0 1-.385-.535l-.002-.007a.815.815 0 0 1 .079-.623l-.002.004a.987.987 0 0 1 .535-.385l.007-.002a.815.815 0 0 1 .623.079l-.004-.002l2.555 1.548a.987.987 0 0 1 .385.535l.002.007a.815.815 0 0 1-.079.623l.002-.004a.746.746 0 0 1-.69.465zm-2.477-2.4h-.077v.077l2.555 1.548l.077-.077zm6.503 2.013h-.008a.424.424 0 0 1-.378-.23l-.001-.002a.46.46 0 0 1-.077-.312v.002a.417.417 0 0 1 .23-.309l.002-.001l2.555-1.548a.46.46 0 0 1 .312-.077h-.002c.138.023.251.11.309.23l.001.002a.46.46 0 0 1 .077.312V5.42a.417.417 0 0 1-.23.309l-.002.001l-2.555 1.548a.285.285 0 0 1-.233.078h.001z"/><path fill="currentColor" d="M14.013 7.742h-.004a.81.81 0 0 1-.69-.384l-.002-.003a.81.81 0 0 1-.076-.625l-.001.006a.899.899 0 0 1 .384-.54l.003-.002l2.555-1.548a.81.81 0 0 1 .625-.076l-.006-.001a.899.899 0 0 1 .54.384l.002.003a.81.81 0 0 1 .076.625l.001-.006a.899.899 0 0 1-.384.54l-.003.002l-2.555 1.548a.962.962 0 0 1-.468.077zm-.078-.774l2.632-1.471V5.42l-.31-.387l.232.31l-2.555 1.548l-.387.31z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rails(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.89 14.987v1.164h1.982A1.166 1.166 0 0 0 24 14.987l-.001-.038v.002v-.45l.001-.047c0-.628-.502-1.139-1.127-1.153h-.989v-.54h1.958v-1.162h-1.879a1.188 1.188 0 0 0-1.136 1.187v.03v-.001v.405l-.001.035c0 .633.504 1.148 1.133 1.165h.002c1.36.006-.327 0 .934 0v.566l-2.006.006zm-10.623-.283s1.061-.097 1.061-1.55a1.493 1.493 0 0 0-1.277-1.587l-.008-.001H7.729v4.585h1.164v-1.106l1.006 1.106h1.722zm-.449-.939h-.926v-1.048h.934a.52.52 0 0 1 .26.524v-.003a.532.532 0 0 1-.264.526l-.003.001zm4.384-2.167h-1.18a1.177 1.177 0 0 0-1.127 1.176v.025v-.001v3.36h1.182v-.8h1.11v.8h1.144v-3.36a1.147 1.147 0 0 0-1.127-1.199h-.001zm-.015 2.45h-1.12v-1.114s0-.25.37-.25h.406c.327 0 .334.25.334.25v1.113zm1.68-2.45h1.23v4.553h-1.23zm2.96 3.376v-3.376h-1.223v4.553h2.88v-1.177z"/><path fill="currentColor" d="M.424 16.151h4.79s-.915-4.43 2.116-6.225c.661-.341 2.764-1.614 6.208 1.086c.109-.097.212-.174.212-.174s-3.153-3.337-6.663-2.964c-1.764.167-3.934 1.871-5.207 4.122a11.39 11.39 0 0 0-1.447 4.09l-.008.064z"/><path fill="currentColor" d="M.424 16.15h4.79s-.915-4.43 2.116-6.225c.661-.341 2.764-1.614 6.208 1.086c.109-.097.212-.174.212-.174s-3.153-3.336-6.663-2.963c-1.77.167-3.94 1.871-5.214 4.122a11.3 11.3 0 0 0-1.441 4.091zm9.981-7.915l.024-.431a2.695 2.695 0 0 0-.568-.22l-.02-.005l-.024.424c.2.07.394.148.588.231z"/><path fill="currentColor" d="m9.826 9.598l-.024.405c.217.008.423.037.622.084l-.022-.004l.024-.4a3.796 3.796 0 0 0-.588-.084zM7.584 7.592h.061L7.52 7.2c-.205 0-.407.014-.605.041l.023-.003l.115.379c.149-.016.322-.026.498-.026h.035zm.291 2.366l.139.444c.146-.079.321-.155.501-.217l.026-.008l-.134-.424a4.105 4.105 0 0 0-.559.217l.026-.011zM5.12 8.267l-.273-.444c-.151.08-.309.174-.473.277l.279.45c.16-.103.311-.2.466-.283zm1.242 2.894l.291.463c.113-.171.231-.32.36-.459l-.002.002l-.273-.438a3.432 3.432 0 0 0-.369.422zm-.88 2.07l.491.411c.028-.284.073-.54.135-.789l-.008.037l-.437-.366a5.68 5.68 0 0 0-.175.666l-.007.041zm-2.651-3.054L2.4 9.779c-.16.16-.309.32-.449.48l.466.424c.139-.19.274-.355.417-.514l-.005.006zM1 13.051l-.697-.272c-.115.277-.24.598-.303.772l.697.27c.079-.218.206-.533.303-.77m4.395 1.814c.009.295.035.575.078.849l-.005-.039l.727.277a9.332 9.332 0 0 1-.146-.836z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rain(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.119 23.764a.802.802 0 0 1 0-1.138l4.166-4.165a.803.803 0 0 1 1.386.557a.802.802 0 0 1-.248.581l-4.165 4.166a.802.802 0 0 1-1.138 0zm8.475-3.94a.802.802 0 0 1 0-1.138l3.059-3.056h-8.5l-4.198 4.195a.804.804 0 1 1-1.138-1.138l3.06-3.06h-.135a5.806 5.806 0 0 1 0-11.612h.067h-.003c.09 0 .18.007.269.011c1.364-2.42 3.917-4.027 6.846-4.027a7.832 7.832 0 0 1 7.832 7.831v.015a3.954 3.954 0 0 1 3.008 3.833a3.946 3.946 0 0 1-3.946 3.946h-.884l-4.198 4.198a.802.802 0 0 1-1.138 0zM1.609 9.82a4.201 4.201 0 0 0 4.196 4.197h14.004a2.338 2.338 0 1 0-.001-4.678c-.577 0-1.104.209-1.512.554l.003-.003a.806.806 0 0 1-1.045-1.227l.001-.001a3.924 3.924 0 0 1 1.852-.871l.024-.004A6.22 6.22 0 0 0 7.774 4.336l-.014.021a5.797 5.797 0 0 1 1.68.937l-.011-.009a.805.805 0 1 1-1.008 1.256l.001.001a4.127 4.127 0 0 0-2.606-.919h-.011h.001A4.2 4.2 0 0 0 1.607 9.82z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rainbow(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M31.558 16.69c0-8.212-6.657-14.869-14.869-14.869S1.82 8.478 1.82 16.69a.91.91 0 0 1-1.82 0v-.116C0 7.357 7.472-.116 16.69-.116s16.69 7.472 16.69 16.69v.122v-.006a.91.91 0 0 1-1.82 0zm-3.643 0c0-6.2-5.026-11.226-11.226-11.226S5.463 10.49 5.463 16.69a.91.91 0 0 1-1.82 0c0-7.206 5.841-13.047 13.047-13.047S29.737 9.484 29.737 16.69a.91.91 0 0 1-1.82 0zm-3.642 0a7.584 7.584 0 0 0-15.168 0a.91.91 0 0 1-1.82 0a9.405 9.405 0 0 1 18.81 0a.91.91 0 0 1-1.82 0zm-3.643 0v-.056a3.942 3.942 0 0 0-7.884 0v.059v-.003a.91.91 0 0 1-1.82 0a5.763 5.763 0 1 1 11.526 0a.91.91 0 0 1-1.82 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rains(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.524 23.758a.825.825 0 0 1 0-1.168l4.272-4.271a.825.825 0 1 1 1.167 1.167l-4.272 4.272a.824.824 0 0 1-1.166 0zm8.505-3.431a.825.825 0 0 1 0-1.168l3.138-3.138h-9.192l-4.304 4.305a.825.825 0 1 1-1.167-1.167l3.138-3.134H5.889a5.954 5.954 0 0 1 0-11.908h.069h-.004c.092 0 .184.007.276.012c1.398-2.482 4.016-4.131 7.02-4.131a8.033 8.033 0 0 1 7.796 6.098l.012.055a5.67 5.67 0 0 1 9.476 4.004v.009a3.004 3.004 0 0 1 2.102 2.861a2.998 2.998 0 0 1-2.998 2.998h-4.137l-4.304 4.301a.824.824 0 0 1-1.166 0zM1.651 10.068a4.309 4.309 0 0 0 4.302 4.304h14.372a2.398 2.398 0 1 0-1.557-4.227l.003-.003a.825.825 0 0 1-1.067-1.258l.001-.001a4.032 4.032 0 0 1 1.9-.894l.024-.004a6.376 6.376 0 0 0-11.653-3.54l-.014.022a5.934 5.934 0 0 1 1.723.961l-.012-.009A.825.825 0 1 1 8.64 6.707l.001.001a4.235 4.235 0 0 0-2.672-.943h-.012h.001a4.308 4.308 0 0 0-4.306 4.304zM21.523 8.11a4.048 4.048 0 0 1 2.841 3.858c0 .905-.298 1.741-.8 2.414l.008-.011h6.053a1.348 1.348 0 1 0-.873-2.377l.002-.002a.826.826 0 1 1-1.067-1.261l.002-.001a2.974 2.974 0 0 1 1.172-.609l.021-.005a4.03 4.03 0 0 0-7.348-2.027l-.009.014z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Random(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24.983 8.539V6.054h-4.902l-3.672 5.945l-2.099 3.414l-3.24 5.256c-.326.51-.889.844-1.53.845H0v-3.568h8.538L12.211 12l2.099-3.414l3.24-5.256a1.812 1.812 0 0 1 1.525-.845h5.904V0l7.417 4.27l-7.417 4.27z"/><path fill="currentColor" d="m12.902 6.316l-.63 1.022l-1.468 2.39l-2.265-3.675H.001V2.485h9.54a1.813 1.813 0 0 1 1.526.838l.004.007l1.836 2.985zM24.983 24v-2.485h-5.904a1.809 1.809 0 0 1-1.521-.838l-.004-.007l-1.836-2.985l.63-1.022l1.468-2.39l2.264 3.675h4.902v-2.485l7.417 4.27l-7.417 4.27z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RaspberryPi(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.656 17.338c-.857.989-1.334 2.79-.709 3.371c.6.449 2.2.391 3.385-1.23a3.026 3.026 0 0 0 .551-1.747a3.03 3.03 0 0 0-.477-1.637l.007.012c-.73-.555-1.778.164-2.757 1.243zm-8.057.3c-.908-1.04-2.088-1.658-2.851-1.2c-.51.382-.605 1.685.123 2.967c1.078 1.524 2.6 1.679 3.221 1.307c.659-.488.3-2.137-.493-3.075zm4.105 3.145c-1.1-.026-2.8.439-2.776 1.032c-.018.4 1.331 1.572 2.7 1.513c1.326.03 2.7-1.139 2.682-1.649s-1.5-.927-2.607-.884zM9.629 6.839c-1.275-.032-2.5.933-2.5 1.493c0 .68 1.008 1.376 2.51 1.394c1.543.01 2.518-.559 2.532-1.26c.016-.794-1.394-1.639-2.518-1.627zm-3.071.532c-2.135-.345-3.913.9-3.842 3.192c.07.884 4.63-3.041 3.843-3.177zm9.749 3.251c.071-2.277-1.709-3.521-3.844-3.176c-.787.135 3.772 4.061 3.844 3.176m.364.824c-1.239-.329-.42 5.049.588 4.615a2.875 2.875 0 0 0-.574-4.592l-.014-.007v-.015zm-14.9 4.675c1.007.45 1.827-4.929.589-4.6a2.895 2.895 0 0 0-1.483 2.525c0 .821.343 1.563.893 2.089l.001.001zm9.415-5.948a2.654 2.654 0 0 0-.457 3.751l-.004-.005a2.793 2.793 0 0 0 3.709 1.074l-.015.007a2.652 2.652 0 0 0 .458-3.751l.004.005a2.76 2.76 0 0 0-3.707-1.055l.015-.007zm-3.1.135A2.74 2.74 0 0 0 6.8 9.993c-1.02 0-1.911.548-2.395 1.366l-.007.013a2.653 2.653 0 0 0-.572 1.652c0 .855.402 1.616 1.027 2.105l.006.004a2.755 2.755 0 0 0 3.684-1.051l.007-.013a2.648 2.648 0 0 0 .564-1.64c0-.853-.4-1.612-1.024-2.1l-.006-.004zm4.369 7.162a2.646 2.646 0 0 0-2.788-2.5h.007a2.656 2.656 0 0 0-2.748 2.524v.035a2.649 2.649 0 0 0 2.79 2.5h-.007a2.628 2.628 0 0 0 2.742-2.511v-.034l.01-.015zM15.67 2.337a15.955 15.955 0 0 0-4.396 2.945l.007-.007c.377 1.5 2.344 1.558 3.063 1.512a.53.53 0 0 1-.314-.266l-.001-.003c.18-.12.821-.016 1.268-.255c-.171-.03-.252-.061-.329-.2a3.571 3.571 0 0 0 1.156-.472l-.015.009a.701.701 0 0 1-.47-.091l.003.002c.419-.179.779-.4 1.103-.664l-.008.006c-.2 0-.406 0-.466-.075a3.82 3.82 0 0 0 .875-.698l.002-.002c-.272.045-.39.016-.454-.03c.295-.226.544-.494.742-.798l.007-.012a.572.572 0 0 1-.525-.001l.003.002c.091-.194.47-.314.69-.779a.979.979 0 0 1-.493-.001l.007.002c.093-.328.242-.613.439-.859l-.004.005a9.146 9.146 0 0 1-1.15-.032l.035.003l.283-.285a2.209 2.209 0 0 0-1.251.126l.015-.005c-.149-.105 0-.255.185-.4c-.385.05-.734.138-1.065.262l.033-.011c-.164-.15.1-.285.24-.436a2.95 2.95 0 0 0-1.106.427l.012-.007c-.18-.165-.015-.314.1-.449c-.354.137-.657.33-.916.57l.002-.002c-.09-.1-.209-.179-.06-.449a2.388 2.388 0 0 0-.73.624l-.004.005c-.194-.134-.119-.3-.119-.449a7.996 7.996 0 0 0-.785.8l-.009.011c-.061-.031-.1-.15-.135-.346c-.779.75-1.889 2.623-.285 3.356a15.012 15.012 0 0 1 4.672-2.47l.107-.03l.041-.075zm-12.259 0a14.945 14.945 0 0 1 4.808 2.542l-.031-.024c1.6-.75.493-2.623-.282-3.356c-.041.194-.085.329-.135.359a8.487 8.487 0 0 0-.78-.803l-.008-.007c0 .15.077.33-.117.45a2.278 2.278 0 0 0-.71-.624L6.145.868c.149.256.025.33-.056.449a2.477 2.477 0 0 0-.883-.594L5.189.717c.12.149.3.3.12.465A2.81 2.81 0 0 0 4.245.765L4.228.763c.135.149.4.3.238.45A4.149 4.149 0 0 0 3.452.957L3.431.955c.181.15.342.289.192.4a2.128 2.128 0 0 0-.849-.173c-.143 0-.284.014-.419.04l.014-.002l.284.284a7.74 7.74 0 0 1-1.137.034l.016.001c.199.238.35.525.432.839l.003.015c-.045.045-.27.016-.483 0c.225.449.6.57.688.765a.551.551 0 0 1-.525-.001l.003.001c.219.307.465.573.741.805l.007.005a.67.67 0 0 1-.469.036l.005.001c.251.269.537.5.851.69l.018.01c-.06.07-.271.069-.479.075c.315.26.676.483 1.065.652l.029.011a.637.637 0 0 1-.469.105h.004c.327.21.708.371 1.116.46l.023.004a.41.41 0 0 1-.344.2c.449.254 1.078.135 1.258.27a.531.531 0 0 1-.31.269l-.004.001c.719.045 2.7-.015 3.072-1.514a15.673 15.673 0 0 0-4.292-2.877l-.099-.04l.04.016zM5.145.1c.243.033.463.103.664.205L5.797.3c.529-.17.65.063.91.159c.577-.12.752.141 1.029.419l.322-.009a3.636 3.636 0 0 1 1.451 2.044l.006.025A3.616 3.616 0 0 1 10.96.881l.012-.008l.321.007c.277-.28.452-.539 1.029-.418c.261-.1.38-.33.911-.166c.33-.1.62-.375 1.057-.045a.918.918 0 0 1 1.046.092l-.001-.001c.5-.06.653.061.774.21c.108 0 .809-.1 1.132.36c.81-.09 1.064.464.774.988a.675.675 0 0 1-.049.975l-.001.001c.15.269.062.553-.27.913a.743.743 0 0 1-.37.837l-.004.002c.06.51-.48.81-.629.914c-.061.3-.181.584-.8.734c-.089.449-.464.523-.824.614a3.808 3.808 0 0 1 2.187 3.744l.001-.013l.181.3a3.654 3.654 0 0 1 .671 5.517l.002-.003a9.813 9.813 0 0 1-.535 1.713l.024-.065a4.053 4.053 0 0 1-2.528 3.209l-.027.01a8.314 8.314 0 0 1-2.384 1.344l-.058.019a4.085 4.085 0 0 1-2.993 1.337h-.094a4.108 4.108 0 0 1-3.009-1.332l-.003-.003a8.442 8.442 0 0 1-2.463-1.378l.016.013a4.074 4.074 0 0 1-2.559-3.197l-.003-.022a9.46 9.46 0 0 1-.505-1.595l-.013-.067a3.652 3.652 0 0 1 .655-5.511l.013-.008l.172-.3a3.8 3.8 0 0 1 2.165-3.72l.023-.01c-.359-.09-.72-.165-.823-.615c-.615-.15-.735-.434-.8-.734c-.15-.105-.689-.4-.629-.928a.758.758 0 0 1-.374-.86l-.001.005c-.314-.346-.4-.645-.27-.915a.665.665 0 0 1-.045-.974c-.285-.525-.03-1.094.779-1C2.158.386 2.864.492 2.967.492c.121-.15.285-.285.779-.225A.942.942 0 0 1 4.8.169L4.795.167A.648.648 0 0 1 5.2.001h.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func React(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.902 11.974a2.41 2.41 0 1 1-4.82 0a2.41 2.41 0 0 1 4.82 0"/><path fill="currentColor" d="M13.491 17.494a27.278 27.278 0 0 1-8.749-1.206l.193.051a10.886 10.886 0 0 1-3.541-1.903l.021.016a3.414 3.414 0 0 1-1.41-2.467l-.001-.013c0-1.594 1.747-3.154 4.681-4.172c2.512-.827 5.403-1.304 8.405-1.304l.417.003h-.021a27.489 27.489 0 0 1 8.885 1.334l-.195-.055a11.058 11.058 0 0 1 3.433 1.814l-.024-.018a3.274 3.274 0 0 1 1.386 2.387l.001.013c0 1.656-1.954 3.332-5.103 4.374c-2.363.729-5.08 1.149-7.895 1.149c-.17 0-.339-.002-.508-.005h.025zm0-9.84c-.106-.002-.232-.002-.358-.002c-2.881 0-5.656.452-8.259 1.289l.191-.053c-2.698.941-3.908 2.228-3.908 3.087c0 .893 1.301 2.3 4.153 3.274c2.249.697 4.834 1.099 7.513 1.099c.235 0 .469-.003.702-.009l-.034.001a26.452 26.452 0 0 0 8.205-1.138l-.187.049c2.96-.984 4.32-2.391 4.32-3.28a2.247 2.247 0 0 0-.966-1.515l-.008-.005a9.806 9.806 0 0 0-2.97-1.574l-.07-.02c-2.368-.766-5.093-1.208-7.92-1.208c-.142 0-.283.001-.424.003h.021z"/><path fill="currentColor" d="m8.023 23.986l-.053.001c-.453 0-.878-.119-1.245-.327l.013.007c-1.378-.8-1.858-3.092-1.28-6.141c.697-3.128 1.827-5.894 3.344-8.4l-.069.124a27.856 27.856 0 0 1 5.425-6.872l.02-.018A10.943 10.943 0 0 1 17.365.332l.073-.027c.407-.196.885-.31 1.39-.31c.501 0 .976.113 1.4.314l-.02-.008c1.435.826 1.911 3.36 1.238 6.606c-.72 2.967-1.818 5.58-3.262 7.959l.07-.123c-1.452 2.63-3.209 4.882-5.266 6.819l-.014.013a10.872 10.872 0 0 1-3.321 2.082l-.074.026c-.46.181-.992.292-1.549.305h-.006zM9.228 9.539l.499.288c-1.381 2.28-2.465 4.926-3.106 7.74l-.034.18c-.533 2.809-.019 4.498.72 4.926c.191.102.417.161.657.161l.061-.001h-.003c.955 0 2.458-.605 4.196-2.122a26.398 26.398 0 0 0 4.972-6.404l.068-.135a26.071 26.071 0 0 0 3.025-7.307l.038-.182c.629-3.058.086-4.93-.686-5.378a2.246 2.246 0 0 0-1.814.092l.013-.006a9.858 9.858 0 0 0-2.907 1.846l.007-.006a26.608 26.608 0 0 0-5.136 6.468l-.069.132z"/><path fill="currentColor" d="M18.96 24c-1.306 0-2.96-.787-4.69-2.276a27.794 27.794 0 0 1-5.464-6.834l-.072-.137c-1.43-2.349-2.551-5.075-3.215-7.975l-.036-.185a10.75 10.75 0 0 1-.255-2.357c0-.527.037-1.044.108-1.551l-.007.058A3.271 3.271 0 0 1 6.701.344l.01-.007c1.43-.83 3.865.024 6.342 2.228a27.865 27.865 0 0 1 5.118 6.54l.072.138a26.769 26.769 0 0 1 3.242 7.798l.038.186c.163.709.257 1.524.257 2.36c0 .577-.044 1.143-.13 1.696l.008-.062a3.419 3.419 0 0 1-1.429 2.456l-.011.007a2.435 2.435 0 0 1-1.207.316l-.054-.001h.003zm-9.228-9.823a27.066 27.066 0 0 0 5.271 6.656l.02.018c2.166 1.863 3.884 2.266 4.628 1.834c.773-.446 1.339-2.276.754-5.233c-.702-2.89-1.787-5.434-3.219-7.741l.065.112a26.77 26.77 0 0 0-4.948-6.384l-.012-.011C9.957 1.354 8.066.888 7.293 1.334a2.234 2.234 0 0 0-.825 1.593v.007a9.854 9.854 0 0 0 .162 3.499l-.013-.066c.673 2.952 1.747 5.559 3.182 7.924l-.066-.117z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Record(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24 12c0-6.627-5.372-12-12-12C5.373 0 0 5.372 0 12c0 6.627 5.372 12 12 12c6.627 0 12-5.372 12-12"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rectangle(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0h24v24H0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Recycle(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m11.538 16.444l-.211 5.178l-.028.31l-5.91-.408a1.577 1.577 0 0 1-.943-.443a2.72 2.72 0 0 1-.654-.905l-.007-.017a2.935 2.935 0 0 1-.202-.759l-.002-.015a2.904 2.904 0 0 1 .06-.934l-.004.019q.106-.521.165-.774c.102-.368.205-.667.324-.959l-.021.059q.239-.647.267-.743q1.099.167 7.164.392zM6.091 8.199l2.533 5.333l-2.068-1.294a19.762 19.762 0 0 0-1.524 1.964l-.044.069c-.352.503-.692 1.08-.986 1.682l-.034.077q-.338.743-.555 1.33a4.631 4.631 0 0 0-.255.856l-.005.031l-.056.295l-2.673-5.023a1.494 1.494 0 0 1-.253-.786v-.003a1.638 1.638 0 0 1 .086-.672l-.003.012l.112-.253q.495-.886 1.604-2.641L.003 7.961zm17.32 7.275l-2.641 5.051a1.5 1.5 0 0 1-.509.652l-.004.003a1.5 1.5 0 0 1-.602.286l-.01.002l-.253.056q-.999.098-3.081.165L16.423 24l-3.236-5.164l2.971-5.094l.098 2.434a21.883 21.883 0 0 0 4.055.066l-.074.005a9.618 9.618 0 0 0 2.462-.485l-.068.02zM12.369 2.472q-.66.886-3.729 6.121L4.183 5.962l-.267-.165L7.082.788c.2-.298.49-.521.831-.632l.011-.003A2.536 2.536 0 0 1 9.061.014L9.05.013c.251.022.483.081.698.171L9.733.178c.227.092.419.192.601.306l-.016-.009c.218.146.409.299.585.466l-.002-.002q.338.31.507.485t.507.555q.341.38.454.493m9.216 4.319l2.983 5.108a1.797 1.797 0 0 1 .175 1.083l.001-.01a2.633 2.633 0 0 1-.393 1.052l.006-.01a2.443 2.443 0 0 1-.461.518l-.004.003a4.271 4.271 0 0 1-.517.383l-.018.01c-.194.115-.42.219-.656.301l-.027.008q-.429.155-.66.225t-.725.197l-.647.165q-.479-1.013-3.729-6.135l4.404-2.744zm-2.012-3.18l1.998-1.168l-3.095 5.249l-5.897-.281l2.125-1.21a20.074 20.074 0 0 0-1.109-2.443l.054.11a13.925 13.925 0 0 0-1.091-1.779l.028.041q-.485-.655-.908-1.126a5.696 5.696 0 0 0-.652-.648l-.008-.007l-.239-.183l5.695.014a1.47 1.47 0 0 1 .825.152l-.008-.004c.217.098.4.234.55.401l.001.002l.155.211q.549.854 1.576 2.669"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reddit(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M25.418 12v.03c0 .543-.156 1.05-.425 1.479l.007-.012a2.772 2.772 0 0 1-1.112 1.021l-.016.007c.108.403.17.865.17 1.343v.018v-.001a6.325 6.325 0 0 1-1.518 4.08l.007-.009a10.223 10.223 0 0 1-4.052 2.936l-.069.024c-1.635.686-3.535 1.085-5.529 1.085L12.728 24h.008l-.146.001c-1.991 0-3.888-.399-5.617-1.121l.096.036a10.26 10.26 0 0 1-4.101-2.944l-.013-.016a6.312 6.312 0 0 1-1.51-4.069v-.007c0-.47.056-.928.161-1.366l-.008.04a2.862 2.862 0 0 1-1.156-1.029l-.007-.011a2.79 2.79 0 0 1-.44-1.512c0-.777.314-1.481.823-1.991a2.704 2.704 0 0 1 1.952-.83h.05h-.003h.039c.799 0 1.519.343 2.019.889l.002.002a13.137 13.137 0 0 1 7.296-2.298h.008l1.646-7.39a.478.478 0 0 1 .211-.296l.002-.001a.46.46 0 0 1 .372-.071l-.003-.001l5.234 1.149c.174-.353.435-.639.757-.838l.009-.005c.319-.2.707-.319 1.123-.319c.585 0 1.116.235 1.501.617c.385.369.624.888.624 1.463v.036v-.002v.03c0 .578-.239 1.1-.624 1.472l-.001.001a2.114 2.114 0 0 1-1.504.624a2.12 2.12 0 0 1-1.497-.617a2.032 2.032 0 0 1-.617-1.461v-.038v.002l-4.738-1.05l-1.475 6.694c2.747.02 5.293.865 7.407 2.3l-.047-.03a2.808 2.808 0 0 1 2.031-.865c.78 0 1.486.317 1.997.83c.509.496.825 1.189.825 1.955v.039V12zM5.929 14.822v.032c0 .576.236 1.097.617 1.471a2.02 2.02 0 0 0 1.463.624h.036h-.002a2.13 2.13 0 0 0 2.128-2.128v-.034c0-.575-.239-1.094-.624-1.462l-.001-.001a2.055 2.055 0 0 0-1.471-.617h-.034h.002a2.132 2.132 0 0 0-2.114 2.113v.001zm11.489 5.036a.513.513 0 0 0 0-.738a.476.476 0 0 0-.341-.142h-.014h.001h-.008a.528.528 0 0 0-.361.142a3.541 3.541 0 0 1-1.694.876l-.023.004a9.256 9.256 0 0 1-4.604-.014l.064.014a3.55 3.55 0 0 1-1.721-.882l.002.002a.526.526 0 0 0-.361-.142h-.019a.478.478 0 0 0-.341.142a.468.468 0 0 0-.16.352v.014c0 .146.061.278.16.372a4.184 4.184 0 0 0 1.65.957l.03.008a8.042 8.042 0 0 0 1.695.414l.043.005q.666.064 1.29.064t1.29-.064a8.384 8.384 0 0 0 1.796-.437l-.058.019a4.214 4.214 0 0 0 1.685-.966l-.002.002zm-.042-2.908h.034c.575 0 1.094-.239 1.462-.624l.001-.001c.381-.374.617-.895.617-1.471v-.034v.002a2.132 2.132 0 0 0-2.113-2.114h-.033c-.576 0-1.097.236-1.471.617a2.02 2.02 0 0 0-.624 1.463v.036v-.002a2.13 2.13 0 0 0 2.128 2.128z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Redis(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M27.994 14.729c-.012.267-.365.566-1.091.945c-1.495.778-9.236 3.967-10.883 4.821a3.649 3.649 0 0 1-2.116.67a3.664 3.664 0 0 1-1.768-.452l.019.01c-1.304-.622-9.539-3.95-11.023-4.659c-.741-.35-1.119-.653-1.132-.933v2.83c0 .282.39.583 1.132.933c1.484.709 9.722 4.037 11.023 4.659a3.63 3.63 0 0 0 1.744.44c.795 0 1.531-.252 2.132-.681l-.011.008c1.647-.859 9.388-4.041 10.883-4.821c.76-.396 1.096-.7 1.096-.982v-2.791z"/><path fill="currentColor" d="M27.992 10.115c-.013.267-.365.565-1.09.944c-1.495.778-9.236 3.967-10.883 4.821a3.647 3.647 0 0 1-2.121.672c-.639 0-1.24-.163-1.763-.449l.019.01c-1.304-.627-9.539-3.955-11.023-4.664c-.741-.35-1.119-.653-1.132-.933v2.83c0 .282.39.583 1.132.933c1.484.709 9.721 4.037 11.023 4.659a3.64 3.64 0 0 0 1.749.442c.793 0 1.527-.251 2.128-.677l-.011.008c1.647-.859 9.388-4.043 10.883-4.821c.76-.397 1.096-.7 1.096-.984v-2.791z"/><path fill="currentColor" d="M27.992 5.329c.014-.285-.358-.534-1.107-.81C25.434 3.986 17.733.923 16.261.383A4.325 4.325 0 0 0 14.467 0c-.734 0-1.426.18-2.035.498l.024-.012c-1.731.622-9.924 3.835-11.381 4.405c-.729.287-1.086.552-1.073.834v2.83c0 .282.39.583 1.132.933c1.484.709 9.721 4.038 11.023 4.66a3.629 3.629 0 0 0 1.744.439c.795 0 1.531-.252 2.133-.68l-.011.008c1.647-.859 9.388-4.043 10.883-4.821c.76-.397 1.096-.7 1.096-.984V5.319h-.009zM10.025 8.013l6.488-.996l-1.96 2.874zm14.351-2.588l-4.253 1.68l-3.835-1.523l4.246-1.679l3.838 1.517zM13.111 2.64l-.628-1.157l1.958.765l1.846-.604l-.499 1.196l1.881.7l-2.426.252l-.543 1.311l-.879-1.457l-2.8-.252zM8.284 4.272c1.916 0 3.467.602 3.467 1.344S10.192 6.96 8.284 6.96S4.81 6.357 4.81 5.616s1.553-1.344 3.474-1.344"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Redo(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.428 6.572a1.094 1.094 0 1 0-1.547 1.545l.001.001A8.024 8.024 0 1 1 5.03 7.67l.012-.01l.515 2.586a1.094 1.094 0 0 0 1.98.398l.002-.004l3.577-5.354v-.004l.006-.01c.017-.027.034-.053.049-.082l.007-.014c.014-.027.027-.055.04-.086l.005-.015c.011-.028.021-.056.03-.086c0-.007 0-.014.005-.02a.903.903 0 0 0 .021-.086c0-.013 0-.027.006-.04s.008-.044.009-.066c0-.036.005-.072.005-.109v-.012c0-.032 0-.064-.005-.096l-.011-.075c0-.01 0-.02-.004-.03a.942.942 0 0 0-.024-.095v-.009a.99.99 0 0 0-.034-.096v-.004a1.38 1.38 0 0 0-.041-.086l-.004-.009a1.166 1.166 0 0 0-.045-.076l-.01-.016l-.046-.062l-.018-.024c-.015-.019-.033-.037-.05-.055l-.021-.024q-.038-.038-.078-.071l-.015-.011a.872.872 0 0 0-.071-.053L5.46.184a1.094 1.094 0 0 0-1.68 1.13l-.001-.007l.791 3.978C1.803 7.14.006 10.255.006 13.79c0 5.637 4.57 10.207 10.207 10.207S20.42 19.427 20.42 13.79c0-2.818-1.142-5.37-2.989-7.217z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Redux(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.54 16.745a1.756 1.756 0 0 0 1.582-1.854v.005a1.823 1.823 0 0 0-1.801-1.736h-.065a1.804 1.804 0 0 0-1.738 1.869v-.003c.026.473.22.897.522 1.217l-.001-.001a11.716 11.716 0 0 1-5.204 5.031l-.069.03a8.538 8.538 0 0 1-4.127 1.048c-.383 0-.761-.025-1.13-.073l.044.005a4.552 4.552 0 0 1-3.277-1.884l-.009-.014a4.671 4.671 0 0 1-.858-2.711c0-.837.218-1.623.601-2.305l-.012.024a7.519 7.519 0 0 1 2.194-2.562l.021-.014a9.756 9.756 0 0 1-.432-1.56l-.011-.066C-.913 14.54-.438 19.121.984 21.286a6.727 6.727 0 0 0 5.603 2.59h-.014l.119.001a8.26 8.26 0 0 0 1.88-.216l-.056.011a12.411 12.411 0 0 0 8.974-6.814l.032-.074z"/><path fill="currentColor" d="M23.18 12.799c-2.403-2.74-5.91-4.46-9.82-4.46a14.1 14.1 0 0 0-.357.005h.018h-.538a1.703 1.703 0 0 0-1.526-.95l-.059.001h.003h-.039c-.966.014-1.744.8-1.744 1.768l.002.085v-.004a1.822 1.822 0 0 0 1.8 1.738h.079a1.784 1.784 0 0 0 1.577-1.094l.004-.012h.585c2.554.003 4.924.787 6.884 2.128l-.042-.027a8.719 8.719 0 0 1 3.392 4.055l.021.058c.26.542.411 1.178.411 1.849c0 .71-.169 1.38-.47 1.972l.011-.025c-.777 1.586-2.379 2.658-4.231 2.658c-.068 0-.135-.001-.202-.004h.01a8.23 8.23 0 0 1-3.183-.707l.052.021c-.379.316-1.011.836-1.47 1.153a9.972 9.972 0 0 0 4.143.994h.012a7.052 7.052 0 0 0 6.223-3.381l.018-.031a7.1 7.1 0 0 0 .54-2.742c0-1.98-.796-3.774-2.085-5.079l.001.001l-.015.032z"/><path fill="currentColor" d="M6.844 17.316a1.82 1.82 0 0 0 1.803 1.738h.064a1.802 1.802 0 0 0-.047-3.601h-.067l-.033-.001a.658.658 0 0 0-.209.033l.005-.001a11.671 11.671 0 0 1-1.703-6.107c0-.364.016-.724.049-1.079l-.003.046A8.553 8.553 0 0 1 8.61 3.337l-.012.016a5.603 5.603 0 0 1 3.944-1.803h.009c3.412-.064 4.835 4.188 4.945 5.878l1.582.473C18.71 2.72 15.488 0 12.407 0a7.33 7.33 0 0 0-6.622 5.131l-.014.052a13.023 13.023 0 0 0-.692 4.229c0 2.588.741 5.003 2.021 7.044l-.032-.055a1.387 1.387 0 0 0-.219.922l-.001-.006z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Reply(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7 7.8v-3l-7 7l7 7v-3l-4-4zm6 1v-4l-7 7l7 7v-4.1c5 0 8.5 1.6 11 5.1c-1-5-4-10-11-11"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func RightAlign(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M27.999 21.333H1.334a1.334 1.334 0 0 0-.002 2.666h26.667a1.334 1.334 0 0 0 .002-2.666zm0-5.333H8.444a1.334 1.334 0 0 0-.002 2.666h19.557A1.334 1.334 0 0 0 28.001 16zm0-5.333H1.334a1.334 1.334 0 0 0-.002 2.666h26.667a1.334 1.334 0 0 0 .002-2.666zM1.334 2.666h26.665l.037.001a1.334 1.334 0 1 0 0-2.668L27.997 0h.002H1.334a1.334 1.334 0 0 0-.002 2.666zm26.665 2.667H8.444l-.037-.001a1.334 1.334 0 1 0 0 2.668l.039-.001h-.002h19.555a1.334 1.334 0 0 0 .002-2.666z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rocket(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 24h3.15c-.176-.634-.552-1.906-1.13-4.9zm10.62-4.9c-.578 2.994-.954 4.266-1.13 4.9h3.15l-2.018-4.9zm-4.303-6.438a2.297 2.297 0 1 1 0-4.594a2.297 2.297 0 0 1 0 4.594m4.525-2.976C10.594 2.437 6.317 0 6.317 0S2.039 2.436 1.792 9.686C1.562 16.406 4.06 24 4.06 24h4.514s2.498-7.594 2.268-14.314"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Room(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.684 19.504l-.623 1.002v.027h1.03v-1.029c0-.271.027-.569.027-.867h-.027c-.136.298-.244.569-.406.867zm7.314-7.965h-.813v-.542h.596l-.379-.379l.379-.379l.569.569l-.352.352zm-1.626 0h-.812v-.542h.813zm-1.626 0h-.812v-.542h.813zm-1.625 0h-.809v-.542h.813zm-1.626 0h-.811v-.542h.813zm-1.625 0h-.811v-.542h.813zm-1.626 0h-.812v-.542h.813zm-1.626 0h-.81v-.542h.813zm-1.625 0h-.81v-.542h.813zm-1.626 0h-.81v-.542h.813zm-1.625 0h-.81v-.542h.813zm-1.626 0h-.809v-.542h.812zm-1.626 0h-.809v-.542h.813zm-1.625 0h-.809v-.406l-.325.325l-.379-.378l.569-.569l.379.379l-.108.108h.68zm.031-1.219l-.379-.379l.569-.569l.379.379zm19.91-.271l-.569-.569l.379-.379l.569.569zM3.034 9.183L2.655 8.8l.569-.569l.379.379zm17.638-.271l-.569-.569l.379-.379l.569.569zM4.198 8.045l-.379-.379l.569-.569l.379.379zm15.305-.271l-.569-.569l.378-.378l.569.569zm-14.14-.867l-.379-.379l.568-.569l.379.379zm12.975-.267l-.569-.569l.379-.379l.569.569zm-11.81-.87l-.379-.379l.569-.569l.379.379zm10.646-.272l-.569-.569l.379-.379l.568.57zm-9.481-.866l-.381-.379l.569-.569l.379.379zm8.32-.271l-.573-.569l.379-.379l.569.569zm-7.155-.867l-.379-.379l.569-.569l.379.379zm5.986-.271l-.572-.569l.379-.379l.569.569zm-4.822-.866l-.379-.379l.569-.569l.379.379zm3.68-.271l-.565-.566l.379-.379l.569.569zm-2.542-.867L10.781.84l.596-.569l.379.379zm1.382-.271l-.569-.569l.352-.379l.569.569zM9.806 13.652c-.461 0-.731.434-.731 1.002c0 .596.271 1.002.731 1.002s.704-.434.704-1.002c.001-.542-.243-1.002-.704-1.002m-2.926 0a.774.774 0 0 0-.303.029l.005-.001v.813h.271c.325 0 .542-.16.542-.434c-.027-.271-.217-.406-.515-.406zm6.095 0c-.461 0-.731.434-.731 1.002c0 .596.271 1.002.731 1.002s.704-.434.704-1.002c.027-.542-.243-1.002-.704-1.002"/><path fill="currentColor" d="M22.618 10.564H1.341c-.741 0-1.342.601-1.342 1.342v.014v-.001v10.739C-.001 23.399.6 24 1.341 24h21.29c.741 0 1.342-.601 1.342-1.342v-.014v.001v-10.726a1.36 1.36 0 0 0-1.354-1.354zM1.436 22.537H.813V12.271h.623zm4.469-9.319c.269-.034.58-.054.895-.054a1.442 1.442 0 0 1 .98.243l-.005-.003a.721.721 0 0 1 .271.626v-.003a.783.783 0 0 1-.509.729l-.005.002c.217.08.325.271.406.569c.074.336.157.617.256.89l-.016-.05h-.704a4.3 4.3 0 0 1-.211-.673l-.006-.031c-.08-.379-.217-.461-.487-.487h-.215v1.165h-.677v-2.923zm3.657 9.21h-3.28v-.65l.596-.542c1.002-.894 1.49-1.436 1.52-1.978c0-.379-.217-.677-.758-.677a1.658 1.658 0 0 0-1.002.385l.003-.002l-.298-.786a2.516 2.516 0 0 1 1.493-.487h.028h-.001a1.44 1.44 0 0 1 1.626 1.428v.036v-.002a2.927 2.927 0 0 1-1.21 1.999l-.009.006l-.434.352v.027H9.57zM8.37 14.681a1.453 1.453 0 0 1 1.448-1.571h.015h-.001h.036a1.38 1.38 0 0 1 1.372 1.526l.001-.006a1.425 1.425 0 0 1-1.417 1.572l-.048-.001h.002h-.034a1.38 1.38 0 0 1-1.373-1.526v.006zm3.223 7.829l-.055.001c-.485 0-.942-.12-1.342-.333l.016.008l.217-.813a2.34 2.34 0 0 0 1.08.298h.003c.569 0 .84-.271.84-.623c0-.461-.461-.65-.921-.65h-.434v-.786h.434c.352 0 .813-.135.813-.542c0-.271-.217-.487-.678-.487a2.19 2.19 0 0 0-.986.277l.011-.006l-.217-.786a2.754 2.754 0 0 1 1.44-.352h-.004c1.002 0 1.544.514 1.544 1.165a1.129 1.129 0 0 1-.832 1.081l-.008.002v.027c.57.071 1.006.552 1.006 1.135v.003c-.03.786-.762 1.382-1.926 1.382zm1.354-6.312h-.02a1.394 1.394 0 0 1-1.388-1.525v.006a1.453 1.453 0 0 1 1.448-1.571h.015h-.001h.036a1.38 1.38 0 0 1 1.372 1.526l.001-.006a1.425 1.425 0 0 1-1.417 1.572zh.002zm4.741 5.12h-.568v1.11h-1.03v-1.111h-2.058v-.704l1.76-2.817h1.328v2.709h.569zm-.271-5.146l-.057-1.138c0-.352-.027-.786-.027-1.219c-.08.379-.217.813-.325 1.165l-.352 1.165h-.514l-.325-1.165c-.108-.352-.19-.786-.271-1.165c-.027.406-.027.867-.054 1.246l-.051 1.138h-.623l.19-2.98h.894l.298 1.002c.08.352.19.704.24 1.057c.08-.352.19-.731.271-1.083l.325-1.002h.894l.16 2.98zm5.742 6.258h-.623V12.19h.623z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rouble(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.778 7.449a3.628 3.628 0 0 0-1.107-2.761l-.001-.001a4.03 4.03 0 0 0-2.923-1.057h.009h-5.454v7.636h5.454a4.032 4.032 0 0 0 2.917-1.06l-.003.003a3.633 3.633 0 0 0 1.108-2.768v.007zm4.04 0a7.11 7.11 0 0 1-2.158 5.368l-.002.002a7.65 7.65 0 0 1-5.581 2.08h.015h-5.795v2.011h8.629c.29 0 .525.235.525.525v.022v-.001v2.203c0 .29-.235.525-.525.525h-.022h.001h-8.608v3.291a.537.537 0 0 1-.537.528H4.886a.525.525 0 0 1-.525-.525v-.022v.001v-3.273H.522a.525.525 0 0 1-.525-.525v-.022v.001v-2.203c0-.29.235-.525.525-.525h.022h-.001h3.818v-2.011H.522a.525.525 0 0 1-.525-.525v-.022v.001v-2.546a.537.537 0 0 1 .528-.537h.019h-.001h3.818V.529c0-.29.235-.525.525-.525h.022h-.001h9.187a7.653 7.653 0 0 1 5.57 2.084l-.004-.004a7.11 7.11 0 0 1 2.157 5.378v-.013z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Rss(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.017 12.754h.029c3.078 0 5.863 1.257 7.869 3.286l.001.001a11.159 11.159 0 0 1 3.274 7.911v.041v-.002h4.6v-.014C15.79 15.255 8.734 8.18.018 8.151H.015zm0-8.155c10.664.045 19.291 8.699 19.291 19.369v.033v-.002h4.602v-.026C23.91 10.764 13.227.049.029-.001H.024v4.6zm6.357 16.186v.002a3.186 3.186 0 1 1-3.186-3.186h.011h-.001a3.186 3.186 0 0 1 3.176 3.185z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ruby(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m6.28 7.906l-.014-.014l-2.96 2.96l7.186 7.174l2.96-2.946l4.226-4.226l-2.96-2.96V7.88H6.266z"/><path fill="currentColor" d="M10.466 0L0 6v12l10.466 6l10.466-6V6zm8.466 16.854l-8.466 4.88L2 16.854V7.12l8.466-4.88l8.466 4.88z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Safari(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.71 11.96v.016a.888.888 0 0 1-.222.588l.001-.001a.704.704 0 0 1-.542.255h-.015h.001h-.016a.888.888 0 0 1-.588-.222l.001.001a.704.704 0 0 1-.255-.542v-.028a.87.87 0 0 1 .229-.589l-.001.001a.724.724 0 0 1 .551-.254h.019c.223 0 .427.084.581.222l-.001-.001c.157.13.256.325.256.544v.013v-.001zm.2.777l4.69-7.782q-.121.107-.904.837t-1.68 1.56t-1.828 1.701t-1.567 1.48a7.132 7.132 0 0 0-.668.68l-.008.009l-4.674 7.768q.094-.094.898-.83t1.687-1.56q.88-.824 1.822-1.701t1.563-1.486q.63-.613.67-.676zM21.576 12v.061a9.299 9.299 0 0 1-1.416 4.947l.023-.039l-.228-.147q-.187-.121-.355-.221a.607.607 0 0 0-.217-.1l-.004-.001a.154.154 0 0 0-.174.175v-.001q0 .134.79.59a9.67 9.67 0 0 1-2.442 2.531l-.029.02a9.275 9.275 0 0 1-3.17 1.466l-.065.014l-.214-.895q-.014-.134-.201-.134a.12.12 0 0 0-.107.073v.001a.18.18 0 0 0-.026.129v-.001l.214.91a9.329 9.329 0 0 1-1.938.198h-.019h.001h-.026a9.392 9.392 0 0 1-4.996-1.43l.04.023q.014-.026.174-.274t.288-.449a.75.75 0 0 0 .126-.249l.001-.005a.154.154 0 0 0-.175-.174h.001q-.08 0-.228.194c-.1.133-.2.284-.29.441l-.011.021q-.154.268-.181.308a9.702 9.702 0 0 1-2.55-2.481l-.02-.03a9.29 9.29 0 0 1-1.461-3.21l-.013-.064l.922-.202a.186.186 0 0 0 .135-.179l-.001-.023v.001a.12.12 0 0 0-.073-.107H3.66a.208.208 0 0 0-.142-.026h.001l-.91.201a9.462 9.462 0 0 1-.185-1.86v-.065c0-1.861.544-3.594 1.482-5.05l-.022.037q.026.014.248.16t.4.254a.673.673 0 0 0 .23.106l.005.001q.174 0 .174-.16q0-.08-.167-.207a4.856 4.856 0 0 0-.413-.276l-.022-.012l-.268-.16a9.832 9.832 0 0 1 2.5-2.511l.032-.021a9.273 9.273 0 0 1 3.205-1.42l.063-.012l.201.898q.026.134.201.134a.12.12 0 0 0 .107-.073v-.001a.208.208 0 0 0 .026-.142v.001l-.201-.88a9.642 9.642 0 0 1 1.791-.178h.042c1.869 0 3.611.544 5.076 1.483l-.038-.023a3.554 3.554 0 0 0-.514.847l-.009.023q0 .174.16.174q.147 0 .64-.857a9.647 9.647 0 0 1 3.938 5.652l.014.066l-.75.16q-.134.026-.134.214a.12.12 0 0 0 .073.107h.001a.18.18 0 0 0 .129.026h-.001l.763-.174a9.5 9.5 0 0 1 .19 1.873zm1.138 0v-.035c0-1.493-.313-2.913-.877-4.197l.026.067c-1.093-2.588-3.111-4.606-5.63-5.673l-.069-.026c-1.229-.538-2.66-.85-4.165-.85s-2.937.313-4.234.877l.069-.027c-2.588 1.093-4.605 3.111-5.672 5.63l-.026.069c-.538 1.229-.85 2.66-.85 4.165s.313 2.937.877 4.234l-.027-.069c1.093 2.588 3.111 4.605 5.63 5.672l.069.026c1.229.538 2.66.85 4.165.85s2.937-.313 4.234-.877l-.069.027c2.588-1.093 4.605-3.111 5.672-5.629l.026-.069c.538-1.218.85-2.637.85-4.13v-.037zM24 12v.033c0 1.672-.35 3.263-.981 4.703l.029-.075c-1.222 2.903-3.485 5.166-6.311 6.359l-.077.029c-1.375.601-2.977.951-4.661.951s-3.286-.35-4.738-.981l.077.03c-2.903-1.222-5.166-3.485-6.359-6.311l-.029-.077C.349 15.286-.001 13.684-.001 12s.35-3.286.981-4.738l-.03.077C2.172 4.436 4.435 2.173 7.261.98l.077-.029C8.713.35 10.315 0 11.999 0s3.286.35 4.738.981l-.077-.03c2.903 1.222 5.166 3.485 6.359 6.311l.029.077c.601 1.364.951 2.955.951 4.627v.035v-.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SaitBoat(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.834 14.35a.416.416 0 0 0 .492.32h-.003l6.658-1.448a.404.404 0 0 0 .261-.185l.001-.002a.422.422 0 0 0 .054-.323l.001.003c-1.813-6.169-5.466-10.8-10.029-12.72H8.23a.415.415 0 0 0-.405.506L7.824.498zm-6.439 1.82l5.259-1.138a.419.419 0 0 0 .261-.178l.001-.002a.376.376 0 0 0 .054-.312l.001.003L8.296 6.83a.402.402 0 0 0-.335-.32h-.002a.23.23 0 0 0-.055-.007h-.015h.001a.414.414 0 0 0-.357.204l-.001.002l-.117.2c-1.558 2.646-2.909 4.935-3.52 8.795a.42.42 0 0 0 .123.365a.46.46 0 0 0 .38.098h-.003zm-.509 4.212a.35.35 0 0 0 .228-.09l.24-.24a2.425 2.425 0 0 1 3.418 0l.24.24a.31.31 0 0 0 .217.09h.01a.35.35 0 0 0 .228-.09l.255-.255a2.415 2.415 0 0 1 1.699-.696c.666 0 1.268.269 1.706.703l.234.234a.344.344 0 0 0 .227.097h.001a.35.35 0 0 0 .228-.09l.234-.234a2.425 2.425 0 0 1 3.418 0l.24.24c.056.055.132.09.217.09h.011h-.001a.35.35 0 0 0 .228-.09l.24-.24a2.39 2.39 0 0 1 .811-.532l.016-.006c.021-.042.048-.075.069-.117l2.488-5.073a.466.466 0 0 0-.049-.481l.001.001a.46.46 0 0 0-.459-.165l.003-.001l-17.666 3.832a.561.561 0 0 0-.399.755l-.001-.004l.565 1.434c.346.145.641.346.889.593a.239.239 0 0 0 .191.095l.025-.001h-.001zm16.535 2.109c-.34-.337-.807-.544-1.323-.544s-.984.208-1.323.545l-.24.24a.863.863 0 0 1-1.213 0l-.24-.24c-.34-.337-.807-.544-1.323-.544s-.984.208-1.323.545l-.234.234a.86.86 0 0 1-.605.248h-.006a.836.836 0 0 1-.602-.255l-.235-.235a1.882 1.882 0 0 0-1.327-.545c-.514 0-.979.205-1.32.538l-.255.255a.863.863 0 0 1-1.213 0l-.24-.24c-.34-.337-.807-.545-1.324-.545s-.984.208-1.324.545l-.24.24a.863.863 0 0 1-1.213 0c-.402-.396-.955-.641-1.565-.641s-1.162.245-1.565.641a.505.505 0 0 0 0 .72a.51.51 0 0 0 .358.151H.53a.501.501 0 0 0 .354-.145a1.196 1.196 0 0 1 1.696 0c.34.336.807.544 1.323.544s.984-.208 1.323-.544l.24-.24a.863.863 0 0 1 1.213 0l.24.24c.336.336.8.544 1.313.544h.011h-.001h.003c.512 0 .976-.205 1.314-.538l.262-.255a.863.863 0 0 1 1.213 0l.234.234c.338.34.806.55 1.323.551h.011c.513 0 .977-.208 1.312-.544l.234-.234a.863.863 0 0 1 1.213 0l.24.24c.34.336.807.544 1.323.544s.984-.208 1.323-.544l.24-.24a.863.863 0 0 1 1.213 0c.091.09.216.145.355.145h.011a.486.486 0 0 0 .352-.151a.5.5 0 0 0 .001-.719z"/><path fill="currentColor" d="M1.708 21.043h.004c.33 0 .628.134.844.351c.34.337.807.544 1.323.544s.984-.208 1.323-.545l.24-.24a.863.863 0 0 1 1.213 0l.24.24c.336.337.8.545 1.314.545h.01h-.001h.003c.512 0 .976-.205 1.314-.538l.262-.255a.863.863 0 0 1 1.22.007l.234.234c.338.339.806.55 1.323.551h.01c.513 0 .977-.208 1.313-.545l.234-.234a.863.863 0 0 1 1.213 0l.24.24c.34.337.807.544 1.323.544s.984-.208 1.323-.545l.24-.24a.863.863 0 0 1 1.213 0c.091.09.216.145.354.145h.011a.486.486 0 0 0 .352-.151a.505.505 0 0 0 0-.72c-.34-.337-.807-.544-1.323-.544s-.984.208-1.323.545l-.24.24a.863.863 0 0 1-1.213 0l-.24-.24c-.34-.337-.807-.544-1.323-.544s-.984.208-1.323.545l-.234.234a.86.86 0 0 1-.605.248h-.006a.836.836 0 0 1-.602-.255l-.235-.235a1.883 1.883 0 0 0-1.326-.544c-.514 0-.98.205-1.32.538l-.255.255a.863.863 0 0 1-1.213 0l-.24-.24c-.34-.337-.807-.545-1.324-.545s-.984.208-1.324.545l-.24.24a.863.863 0 0 1-1.213 0c-.402-.396-.955-.641-1.565-.641s-1.162.245-1.565.641a.5.5 0 0 0 0 .72a.51.51 0 0 0 .358.152H.51a.501.501 0 0 0 .354-.145c.211-.221.507-.359.837-.359h.005z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sass(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m27.533 13.796l.132-.002c.834 0 1.623.194 2.323.54l-.031-.014a3.908 3.908 0 0 1 1.541 1.335l.009.014a2.748 2.748 0 0 1 .474 1.68v-.006a2.074 2.074 0 0 1-.698 1.448l-.002.002a2.188 2.188 0 0 1-.785.495l-.015.005q-.25.05-.275-.075t.225-.275c.452-.274.779-.716.898-1.237l.003-.013a1.918 1.918 0 0 0-.2-1.006l.005.011a2.736 2.736 0 0 0-.94-1.094l-.01-.006a4.543 4.543 0 0 0-2.083-.699l-.016-.001a6.223 6.223 0 0 0-3.142.365l.042-.015c.145.339.236.732.249 1.144v.005a2.377 2.377 0 0 1-1.145 1.99l-.011.006c-.31.22-.664.413-1.041.563l-.035.012a1.787 1.787 0 0 1-.879.174h.005q-.949-.199-.45-1.65c.145-.421.348-.786.605-1.107l-.006.007c.329-.43.69-.807 1.089-1.141l.011-.009l-.098-.146a6.277 6.277 0 0 1-.486-1.054l-.014-.046q-.15-.5-.25-.949l-.05-.35l-.5.949q-.65 1.2-1.25 2.099l-.15.25c.168.408.292.881.348 1.375l.002.025a2.25 2.25 0 0 1-.942 1.996l-.007.005a8.993 8.993 0 0 1-.94.5l-.06.025a3.466 3.466 0 0 1-1.191.275h-.009a1.621 1.621 0 0 1-.81-.154l.01.004a.588.588 0 0 1-.149-.703l-.002.004c.4-.561.794-1.049 1.213-1.515l-.013.015l1.1-1.399l-.199-.45a6.597 6.597 0 0 1-.436-1.001l-.014-.048q-.15-.5-.25-.949l-.05-.35l-.5 1.2q-.599 1.399-1.1 2.399q-.747 1.5-1.25 2.3l-.1.15q-.599.949-1.15.949t-.747-.7a4.323 4.323 0 0 1-.152-1.147l.001-.108v.005l.1-.7v.05q-.4.949-.8 1.749a6.618 6.618 0 0 1-.563.919l.013-.019c-.206.216-.496.35-.817.35h-.008c-.021.002-.045.002-.069.002a.977.977 0 0 1-.657-.253l.001.001a2.972 2.972 0 0 1-.742-1.304l-.005-.021a4.552 4.552 0 0 1-.253-1.739v.01c.079-.796.238-1.525.471-2.22l-.021.071l-1.799 1.05v.05a4.282 4.282 0 0 1 .399 2.215l.001-.016a5.478 5.478 0 0 1-.913 2.77l.013-.02a4.99 4.99 0 0 1-2.314 1.962l-.033.012a4.484 4.484 0 0 1-1.825.382c-.418 0-.824-.056-1.208-.162l.032.007a1.375 1.375 0 0 1-.743-.593l-.003-.006a2.32 2.32 0 0 1-.403-1.196v-.006a3.286 3.286 0 0 1 .309-1.769l-.009.02a6.506 6.506 0 0 1 2.316-2.33l.03-.016a17.4 17.4 0 0 1 2.247-1.309l.106-.047l.25-.15q-.45-.35-1.65-1.2a25.85 25.85 0 0 1-2.874-2.221l.025.022a5.454 5.454 0 0 1-1.741-2.812l-.008-.038a3.213 3.213 0 0 1 .358-2.066l-.008.017a8.91 8.91 0 0 1 2.288-3.035l.012-.01a16.486 16.486 0 0 1 3.462-2.506l.087-.043A23.355 23.355 0 0 1 10.068.803l.17-.052a12.736 12.736 0 0 1 4.296-.747h.002a7.847 7.847 0 0 1 3.695.794l-.046-.021c.9.41 1.574 1.182 1.844 2.126l.006.024a4.493 4.493 0 0 1-.236 2.853l.011-.029a7.684 7.684 0 0 1-1.797 2.647l-.003.003a8.124 8.124 0 0 1-2.718 1.78l-.055.019a10.224 10.224 0 0 1-3.63.799l-.019.001a6.018 6.018 0 0 1-2.342-.21l.043.011a3.495 3.495 0 0 1-1.254-.703l.004.003a2.28 2.28 0 0 1-.594-.734l-.006-.013q-.15-.4-.025-.475t.175-.025l.25.25c.279.259.612.463.98.593l.02.006c.506.176 1.09.278 1.697.278c.195 0 .387-.01.576-.031l-.023.002a9.94 9.94 0 0 0 4.336-1.573l-.037.023a7.145 7.145 0 0 0 2.531-2.54l.018-.034c.301-.442.481-.988.481-1.576c0-.361-.068-.707-.192-1.025l.007.019a2.463 2.463 0 0 0-1.685-1.122l-.014-.002a7.971 7.971 0 0 0-3.349-.067l.05-.008a16.175 16.175 0 0 0-4.586 1.541l.091-.043q-6.148 3.1-6.349 6.349v.049c0 .873.343 1.666.902 2.252l-.001-.001a12.481 12.481 0 0 0 2.11 1.873l.039.026a16.661 16.661 0 0 1 1.898 1.649l.001.002l.1.1l3.2-1.749a5.405 5.405 0 0 1 1.308-1.334l.016-.011c.4-.336.908-.555 1.465-.599l.009-.001c.02-.002.043-.002.067-.002c.373 0 .693.225.831.547l.002.006a1.97 1.97 0 0 1 .149 1.06l.001-.01l-.097.493l.15-.1a.622.622 0 0 1 .629-.073l-.004-.002a.44.44 0 0 1 .319.528l.001-.003l-.15.55q-.45 1.799-.599 2.599a2.109 2.109 0 0 0-.074.705v-.006q.025.3.075.3t.199-.3l.15-.35q.05 0 0 .05l.55-1.15q1.95-4.298 2.099-4.8l.3-.949q.05-.1.4-.199a3.061 3.061 0 0 1 .856-.1h-.003q.7 0 .7.35l-.05.25a8.76 8.76 0 0 0-.309 1.087l-.011.063a2.944 2.944 0 0 0 .027 1.019l-.003-.019l.05.199c.102.417.257.783.46 1.117l-.01-.017q.599-.999 1.15-2.049c.29-.483.56-1.043.777-1.629l.023-.07c.075-.377.161-.695.266-1.005l-.016.056q.05-.15.427-.225c.238-.048.511-.075.79-.075h.036h-.002q.7 0 .747.35l-.1.199c-.119.336-.223.74-.293 1.155l-.006.045a3.56 3.56 0 0 0 .054 1.022l-.004-.023v.199c.115.436.286.819.51 1.166l-.01-.016l.15.35a6.503 6.503 0 0 1 2.855-.65h.039h-.002zM6.094 21.049a3.983 3.983 0 0 0 1.124-2.778a4.01 4.01 0 0 0-.182-1.197l.008.028l-.8.5c-.816.465-1.522.984-2.161 1.574l.008-.007a3.524 3.524 0 0 0-1.019 1.674l-.006.025q-.225.925.225 1.2c.183.084.398.133.623.133c.246 0 .479-.058.684-.162l-.009.004a3.735 3.735 0 0 0 1.497-.998l.003-.003zm6.697-4.604q.25-.65.55-1.55q.4-1.2.275-1.399t-.475-.05a2.05 2.05 0 0 0-.697.547l-.003.004a6.958 6.958 0 0 0-.582.866l-.018.034a4.722 4.722 0 0 0-.492 1.116l-.009.034a5.125 5.125 0 0 0-.299 2.011v-.012q.05.999.32 1.075t.625-.725c.229-.413.433-.893.585-1.396l.014-.054q.149-.354.204-.501zm5.546 2.65a2.112 2.112 0 0 0 1.249-1.741l.001-.008v-.05l-.55.7l-.999 1.1a.064.064 0 0 0-.025.05c0 .02.01.038.024.05q.048.049.3-.1zm4.299-.95a1.782 1.782 0 0 0 1.3-1.712v-.039v.002c-.01-.2-.046-.388-.105-.566l.004.015c-.389.346-.708.76-.939 1.227l-.011.023q-.451.899-.252 1.048z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Saucelabs(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.93 15.475h6.4l-3.586 6.762C2.287 20.09.02 16.314.02 12.009c0-6.627 5.373-12 12-12c1.782 0 3.474.389 4.995 1.086l-.075-.031L5.428 12.58h10.25l-7.539 7.614l3.115-5.876H3.8a8.507 8.507 0 0 1-.323-2.316V12a8.534 8.534 0 0 1 8.533-8.523h.012h-.001c.093 0 .187 0 .28.005l1.094-1.1a9.488 9.488 0 0 0-1.377-.098c-5.368 0-9.721 4.348-9.728 9.715v.04c0 1.236.235 2.417.663 3.501zM18.255 1.764c3.456 2.147 5.723 5.922 5.723 10.227c0 6.627-5.373 12-12 12c-1.783 0-3.476-.389-4.997-1.087l.075.031l11.43-11.513H8.201l7.656-7.616l-3.11 5.876h7.46c.192.692.308 1.489.319 2.311v.006a8.535 8.535 0 0 1-8.535 8.522h-.013h.001c-.1 0-.21 0-.314-.006l-1.084 1.1c.429.067.923.105 1.426.105c5.363 0 9.711-4.348 9.711-9.711a9.677 9.677 0 0 0-.672-3.556l.023.066h-6.394l3.581-6.756z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Save(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23.681 6.158L17.843.32a1.093 1.093 0 0 0-.771-.32H1.092C.489 0 .001.489.001 1.091v21.817c0 .603.489 1.091 1.091 1.091h21.817c.603 0 1.091-.489 1.091-1.091V6.928c0-.301-.122-.574-.32-.771zM6.549 2.182h6.546v5.819H6.546zm0 19.635v-5.818h10.905v5.818zm15.273 0h-2.185v-6.908c0-.603-.489-1.091-1.091-1.091H5.455c-.603 0-1.091.489-1.091 1.091v6.909H2.182V2.181h2.182V9.09c0 .603.489 1.091 1.091 1.091h8.728c.603 0 1.091-.489 1.091-1.091V2.181h1.344l5.199 5.199z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SaveOne(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23.681 6.158L17.843.32a1.093 1.093 0 0 0-.771-.32H1.092C.489 0 .001.489.001 1.091v21.817c0 .603.489 1.091 1.091 1.091h21.817c.603 0 1.091-.489 1.091-1.091V6.928c0-.301-.122-.574-.32-.771zM6.549 2.183h6.546v5.818H6.546zm-4.363 0h2.177v6.909c0 .603.489 1.091 1.091 1.091h8.728c.603 0 1.091-.489 1.091-1.091v-6.91h1.344l5.199 5.199v3.828l-1.43-1.435a1.088 1.088 0 0 0-1.542 0L14.8 13.818H5.454c-.603 0-1.091.489-1.091 1.091v6.909H2.181zM6.546 16h6.07l-2.841 2.845c-.197.197-.319.47-.32.771v2.207H6.543zm5.09 5.817v-1.753l7.979-7.976l1.754 1.754l-7.975 7.975zm6.909 0h-2.069l5.342-5.342v5.342z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Scissors(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.349 14.378a5.62 5.62 0 0 0-3.986-1.651a5.613 5.613 0 0 0-2.988.856l.024-.014l-2.081-3.075l5.954-8.792A1.092 1.092 0 0 0 17.464.476l-.002.004l-5.46 8.069L6.541.48a1.092 1.092 0 1 0-1.805 1.225l-.002-.003l5.95 8.793l-2.081 3.075a5.584 5.584 0 0 0-2.964-.841a5.64 5.64 0 1 0 4.567 2.331l.011.016l1.784-2.637l1.783 2.637a5.594 5.594 0 0 0-1.056 3.285a5.636 5.636 0 1 0 9.622-3.985zM8.078 20.802a3.455 3.455 0 1 1 0-4.884c.631.623 1.022 1.488 1.022 2.444s-.391 1.82-1.021 2.443zm12.727 0a3.455 3.455 0 1 1 0-4.884c.631.623 1.023 1.488 1.023 2.444s-.391 1.82-1.021 2.443z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Scorp(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M33.402 24H.027c0-.111-.018-.214-.018-.318q0-8.064-.007-16.14a.647.647 0 0 1 .267-.574l.002-.001Q4.827 3.535 9.376.091c.128-.098.223-.14.361 0q3.407 3.435 6.815 6.864c.01.024.019.053.024.083v.003c-.14-.061-.257-.104-.367-.153c-2.822-1.256-5.651-2.505-8.468-3.772a.326.326 0 0 0-.495.132l-.001.002q-2.338 3.389-4.696 6.759a.357.357 0 0 0 .005.502q3.408 4.463 6.799 8.943a.35.35 0 0 0 .515.111l-.001.001c2.645-1.206 5.296-2.4 7.941-3.606a.607.607 0 0 1 .603.02l-.003-.002q7.42 3.94 14.835 7.873c.055.03.104.061.153.091a.236.236 0 0 1 .005.048z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Search(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m22.241 24l-7.414-7.414a9.482 9.482 0 0 1-5.652 1.885h-.002l-.108.001A9.065 9.065 0 0 1 0 9.407l.001-.114v.006a9.298 9.298 0 0 1 18.596 0a9.8 9.8 0 0 1-1.904 5.682l.019-.027l7.414 7.414zM9.299 2.513a6.758 6.758 0 1 0 .029.001zH9.3z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sentry(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M26.745 22.818c.238-.357.38-.796.38-1.268s-.142-.911-.386-1.276l.005.008L15.758 1.181A2.464 2.464 0 0 0 13.558 0h.005a2.883 2.883 0 0 0-2.19 1.176l-.006.008l-3.55 6.252l.845.507c2.706 1.536 4.886 3.716 6.379 6.339l.043.083a17.173 17.173 0 0 1 2.363 7.549l.003.057h-2.532a15.01 15.01 0 0 0-2.072-6.493l.038.069a13.348 13.348 0 0 0-5.338-5.549l-.067-.034l-.847-.505l-3.379 5.746l.844.507c2.157 1.303 3.671 3.485 4.051 6.037l.006.046H2.576a.595.595 0 0 1-.339-.17a.237.237 0 0 1 0-.339l1.524-2.705a4.019 4.019 0 0 0-1.833-1.011l-.028-.006L.379 20.28c-.238.357-.38.796-.38 1.268s.142.911.386 1.277l-.005-.008a2.296 2.296 0 0 0 2.204 1.181l-.008.001h7.775v-1.017a11.474 11.474 0 0 0-1.386-5.301l.03.061a8.862 8.862 0 0 0-2.843-3.19l-.03-.02l1.188-2.194a15.523 15.523 0 0 1 3.849 4.162l.038.064a12.453 12.453 0 0 1 1.696 6.312l-.001.116v-.006v1.017h6.59v-1.144c0-3.576-1.006-6.917-2.75-9.756l.046.081c-1.459-2.72-3.523-4.944-6.021-6.551l-.069-.042l2.536-4.393a.524.524 0 0 1 .337-.169h.002c.17 0 .17 0 .339.17L24.889 21.3a.237.237 0 0 1 0 .339a.594.594 0 0 1-.336.169h-2.539v2.034h2.535a1.936 1.936 0 0 0 2.191-1.006l.005-.011z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Share(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.237 15.638h-.003a4.04 4.04 0 0 0-3.004 1.334l-.003.004l-8.948-4.348c0-.167.084-.418.084-.669a1.425 1.425 0 0 0-.087-.595l.003.01l8.948-4.348a4.04 4.04 0 0 0 3.007 1.338h.004a4.181 4.181 0 1 0-4.181-4.181a1.425 1.425 0 0 0 .087.595l-.003-.01l-8.948 4.348a4.04 4.04 0 0 0-3.007-1.338h-.004a4.181 4.181 0 0 0 0 8.362h.003a4.04 4.04 0 0 0 3.004-1.334l.003-.004l8.948 4.348c0 .167-.084.418-.084.669a4.181 4.181 0 1 0 8.363-.092a4.09 4.09 0 0 0-4.09-4.09l-.097.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShareA(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.408 22.528C2.281 15.77 6.08 10.995 11.624 8.434c1.678-.752 3.633-1.353 5.673-1.709l.151-.022c.462 0 .464-.014.464-3.352V0l11.446 11.446l-11.446 11.446V16.19H16.52a24.855 24.855 0 0 0-10.51 2.58l.145-.065a20.31 20.31 0 0 0-4.767 3.825l-.013.015l-1.374 1.454z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shazam(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.001 0c-6.627 0-12 5.373-12 12s5.373 12 12 12s12-5.373 12-12s-5.373-12-12-12M9.873 16.736h-.054a5.415 5.415 0 0 1-3.515-1.288l.008.007a5.76 5.76 0 0 1-2.028-4.006l-.001-.018a5.508 5.508 0 0 1 1.44-4.071l-.004.004C6.861 6.1 8.908 4.139 8.994 4.055a1.393 1.393 0 0 1 1.922 2.016c-.021.02-2.061 1.976-3.137 3.164a2.747 2.747 0 0 0-.719 2.035v-.008a3.006 3.006 0 0 0 1.065 2.082l.005.004a2.822 2.822 0 0 0 1.75.603c.687 0 1.317-.244 1.808-.65l-.005.004a25.588 25.588 0 0 0 1.38-1.386a1.393 1.393 0 0 1 2.047 1.889l.001-.001c-.035.035-.852.924-1.572 1.572a5.467 5.467 0 0 1-3.615 1.357H9.87zm8.41-.1c-1.143 1.262-3.189 3.225-3.276 3.309a1.393 1.393 0 0 1-1.922-2.016c.021-.02 2.063-1.977 3.137-3.166a2.747 2.747 0 0 0 .719-2.035v.008a2.998 2.998 0 0 0-1.065-2.08l-.005-.004a2.82 2.82 0 0 0-1.749-.604c-.687 0-1.317.245-1.808.651l.005-.004a23.972 23.972 0 0 0-1.38 1.384a1.392 1.392 0 1 1-2.047-1.887l-.001.001c.034-.037.85-.926 1.571-1.573a5.584 5.584 0 0 1 3.659-1.356a5.59 5.59 0 0 1 3.577 1.287l-.009-.008a5.761 5.761 0 0 1 2.027 4.007l.001.018a5.508 5.508 0 0 1-1.437 4.072l.004-.004z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shield(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17 13V3h-7v17.766a22.036 22.036 0 0 0 3.376-2.177l-.048.036Q17 15.751 17 13m3-12v12a7.15 7.15 0 0 1-.541 2.712l.017-.048a9.968 9.968 0 0 1-1.315 2.37l.018-.026a14.545 14.545 0 0 1-1.827 1.98l-.013.012a21.487 21.487 0 0 1-1.915 1.567l-.062.043q-.906.64-1.89 1.211t-1.398.774t-.664.313a.932.932 0 0 1-.818-.002l.005.002q-.25-.11-.664-.313t-1.398-.774t-1.89-1.211a21.613 21.613 0 0 1-1.996-1.624l.015.014a14.574 14.574 0 0 1-1.815-1.958l-.025-.034a9.813 9.813 0 0 1-1.273-2.277l-.024-.067a7.034 7.034 0 0 1-.523-2.663V1c.008-.549.451-.992.999-1h18.001c.549.008.992.451 1 .999V1z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ship(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.494 20.659a3.677 3.677 0 0 0-2.155 1.25l-.005.006a2.161 2.161 0 0 1-2.093.822l.013.002v-.013l-.08.002l-.084-.002h.004v.013a2.158 2.158 0 0 1-2.076-.819l-.004-.005a3.668 3.668 0 0 0-2.137-1.252l-.023-.004a3.527 3.527 0 0 0-4.062 2.284l-.008.025l1.182.395a2.294 2.294 0 0 1 2.683-1.473l-.015-.003a2.45 2.45 0 0 1 1.479.892l.004.005a3.294 3.294 0 0 0 2.692 1.212h-.006c.117 0 .24-.006.37-.014c.128.008.252.014.37.014a3.291 3.291 0 0 0 2.681-1.206l.005-.006a2.446 2.446 0 0 1 1.468-.894l.015-.002a2.29 2.29 0 0 1 2.663 1.46l.005.016l1.182-.395a3.534 3.534 0 0 0-4.093-2.305l.023-.004z"/><path fill="currentColor" d="M3.191 11.826a.923.923 0 0 1-.118-1.836h.005v.006l.554.1a.92.92 0 0 1 .48.807v.001a.925.925 0 0 1-.92.924zm1.131-6.198l2.72-.954h6.263l2.72.954V7.91l-5.851-.96l-5.851.96zm4.017-4.382h3.67l.589 2.181H7.75zm8.379 8.847l.554-.1v-.005a.925.925 0 0 1 .81.913a.923.923 0 1 1-1.846 0V10.9c0-.347.192-.649.475-.805l.005-.002zM5.1 19.656h.006c.323 0 .639.029.946.085l-.032-.005a4.6 4.6 0 0 1 2.661 1.519l.005.006c.154.19.34.347.55.465l.01.005v-.018l.903.08h.038l.907-.091v.032c.22-.123.406-.28.557-.465l.003-.003a4.587 4.587 0 0 1 2.637-1.521l.028-.004c.276-.052.594-.083.919-.084h.006c.338 0 .668.036.985.106l-.03-.006l4.143-11.134l-3.077-.506v-3.37L13.92 3.573L12.96.002H7.382l-.96 3.571l-3.345 1.174v3.37L0 8.623l4.143 11.134c.287-.064.617-.1.955-.1h.006z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Shopify(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.889 23.979l7.216-1.561S18.501 4.8 18.48 4.688a.222.222 0 0 0-.211-.192c-.1 0-1.929-.136-1.929-.136s-1.275-1.274-1.439-1.411a.377.377 0 0 0-.118-.073l-.003-.001l-.914 21.1h.023zm-3.627-12.674a4.203 4.203 0 0 0-1.77-.424h-.004c-1.447 0-1.5.906-1.5 1.141c0 1.232 3.24 1.715 3.24 4.629c0 2.3-1.44 3.76-3.406 3.76l-.133.002a4.677 4.677 0 0 1-3.405-1.465l-.002-.002l.646-2.086c.62.544 1.397.923 2.254 1.063l.026.003l.031.001a.945.945 0 0 0 .944-.932v-.001c0-1.619-2.654-1.694-2.654-4.359c-.038-2.235 1.567-4.416 4.823-4.416a3.964 3.964 0 0 1 1.899.372l-.024-.01l-.945 2.72l-.02.01zM9.722.83a.67.67 0 0 1 .402.136l-.002-.001C9.143 1.43 8.063 2.6 7.619 4.957a62.82 62.82 0 0 1-1.889.578C6.252 3.75 7.503.84 9.722.84zm1.23 2.949v.135c-.754.232-1.583.484-2.394.736a4.163 4.163 0 0 1 2.073-2.96l.021-.011a5.769 5.769 0 0 1 .309 2.11v-.01zm.539-2.234c.694.074 1.141.867 1.429 1.755c-.349.114-.735.231-1.158.366v-.252l.002-.132c0-.621-.1-1.22-.284-1.779l.011.04zm2.992 1.289c-.02 0-.06.021-.078.021s-.289.075-.714.21c-.423-1.233-1.176-2.37-2.508-2.37h-.115A1.647 1.647 0 0 0 9.821 0h-.004c-3.106 0-4.59 3.877-5.055 5.846c-1.194.365-2.063.636-2.16.674c-.675.213-.694.232-.772.87C1.755 7.852 0 21.453 0 21.453L13.561 24l.927-21.166z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBag(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.595 6.567c-.605 2.91-2.674 5.062-5.132 5.062s-4.531-2.16-5.134-5.07a1.863 1.863 0 0 1-1.306-1.767v-.001c0-1.028.832-1.861 1.859-1.864a1.864 1.864 0 0 1 1.862 1.861a1.85 1.85 0 0 1-.727 1.466l-.004.003c.441 2.106 1.818 3.649 3.447 3.649s3.014-1.551 3.45-3.666a1.852 1.852 0 0 1-.71-1.451a1.864 1.864 0 0 1 1.861-1.862a1.858 1.858 0 0 1 .546 3.636l-.013.003zm5.055 16.72L19.508.548a.573.573 0 0 0-.572-.545H1.72a.573.573 0 0 0-.572.543v.001L.001 23.4L0 23.427a.57.57 0 0 0 .161.398c.105.11.252.178.415.178h19.52a.574.574 0 0 0 .574-.574v-.001a.576.576 0 0 0-.018-.145l.001.004z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBagOne(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.596 23.997H11.545a.325.325 0 0 1-.325-.326v-.016v.001l.65-12.947a.325.325 0 0 1 .324-.306h9.752c.173 0 .314.135.324.305v.001l.647 12.882c.006.024.01.052.01.081v.001a.324.324 0 0 1-.324.324h-.001zm-9.098-10.878c.002.468.309.863.732.998l.007.002c.341 1.651 1.514 2.873 2.908 2.873s2.565-1.223 2.904-2.868a1.056 1.056 0 1 0-.955-.185l.002.002c-.248 1.198-1.029 2.077-1.955 2.077s-1.703-.874-1.953-2.067a1.052 1.052 0 1 0-1.696-.834v.003zM.325 22.72a.325.325 0 0 1-.324-.343v.001L.905 4.383a.325.325 0 0 1 .324-.306h2.204C4.124 1.677 5.968.001 8.112.001s3.986 1.677 4.677 4.074h1.995c.172 0 .314.135.323.305v.001l.261 5.209h-3.174c-.608 0-1.104.477-1.135 1.077v.003l-.605 12.051zM10.948 5.678a1.442 1.442 0 0 0 .784 2.65a1.44 1.44 0 0 0 .518-2.786l-.01-.003C12.004 2.881 10.243.81 8.113.81S4.215 2.885 3.982 5.548a1.44 1.44 0 1 0 1.299.127l-.006-.003c.168-2.002 1.376-3.56 2.838-3.56s2.673 1.563 2.838 3.569zM6.559 4.073h3.1a1.905 1.905 0 0 0-1.541-1.15l-.009-.001c-.7.077-1.28.525-1.545 1.14zm12.799 9.018a.382.382 0 1 1 .762 0a.382.382 0 0 1-.761 0zm-5.224 0a.382.382 0 1 1 .382.382a.382.382 0 0 1-.378-.382zm-2.976-6.146a.567.567 0 1 1 1.135 0a.567.567 0 0 1-1.135 0m-7.228 0a.57.57 0 1 1 1.137 0a.57.57 0 0 1-1.136 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBarcode(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M31.364 19.73h-.989V4.48h.989zm-2.078-3.761h-2.473V4.48h2.473zm-1.108 3.35v.002a.629.629 0 0 1-.234.49l-.001.001a.843.843 0 0 1-.55.202l-.031-.001h.001l-.034.001a.841.841 0 0 1-.556-.208l.001.001a.627.627 0 0 1-.225-.481v-.068h.558v.078a.25.25 0 0 0 .07.173a.265.265 0 0 0 .182.072l.017-.001h-.001l.016.001a.27.27 0 0 0 .182-.072a.247.247 0 0 0 .071-.174v-.589a.68.68 0 0 1-.407.137a.67.67 0 0 1-.53-.201a.717.717 0 0 1-.158-.449l.001-.045v.002v-.499a.624.624 0 0 1 .142-.365l-.001.001a.772.772 0 0 1 .348-.234l.005-.002a.942.942 0 0 1 .308-.051h.01l.032-.001c.214 0 .409.079.558.21l-.001-.001a.628.628 0 0 1 .225.482v.006zm-2.205-3.35h-1.485V4.48h1.485zm-.158 3.35c0 .202-.098.38-.249.491l-.002.001a.923.923 0 0 1-.579.202h-.027h.001h-.068a.918.918 0 0 1-.579-.205l.002.001a.615.615 0 0 1-.247-.49v-.267c0-.188.085-.356.218-.469l.001-.001a.792.792 0 0 1 .214-.14l.005-.002a.68.68 0 0 1-.334-.298l-.002-.003a.56.56 0 0 1-.06-.255v-.02v.001v-.175c0-.18.079-.342.205-.453l.001-.001a.852.852 0 0 1 .605-.208h-.002h.042l.044-.001c.216 0 .413.079.564.21l-.001-.001a.609.609 0 0 1 .205.457v.015v-.001v.171c0 .174-.077.33-.198.437l-.001.001a.724.724 0 0 1-.192.123l-.005.002a.822.822 0 0 1 .368.328l.002.004a.518.518 0 0 1 .069.259v.018v-.001zm-2.49-1.809l-.755 2.473h-.528l.755-2.473h-1.089v-.44h1.617zM21.017 4.48h1.98v11.489h-1.978zm-.048 13.314h-.562v-.078a.247.247 0 0 0-.269-.247h.001l-.017-.001a.256.256 0 0 0-.181.074a.25.25 0 0 0-.07.173v.588a.667.667 0 0 1 .404-.134h.004a.665.665 0 0 1 .532.201a.704.704 0 0 1 .159.447l-.001.045v-.002v.499a.633.633 0 0 1-.245.47l-.001.001a.852.852 0 0 1-.537.189h-.027h.001l-.034.001a.837.837 0 0 1-.555-.208l.001.001a.627.627 0 0 1-.225-.481v-1.6a.63.63 0 0 1 .127-.381l-.001.002a.748.748 0 0 1 .355-.253l.005-.001a.898.898 0 0 1 .313-.055h.011h-.001l.032-.001c.214 0 .409.079.558.21l-.001-.001a.628.628 0 0 1 .225.483v.005zm-2.364 1.493v.013a.658.658 0 0 1-.25.516l-.001.001a.907.907 0 0 1-.569.2l-.037-.001h.002l-.031.001a.844.844 0 0 1-.559-.21l.001.001a.628.628 0 0 1-.225-.483v-.074l.515-.032l.012.099a.26.26 0 0 0 .171.235l.002.001a.26.26 0 0 0 .098.019q.33 0 .33-.281v-.712a.251.251 0 0 0-.117-.211l-.001-.001a.252.252 0 0 0-.136-.039l-.017.001h.001l-.018-.001a.246.246 0 0 0-.192.091a.283.283 0 0 0-.058.167v.079h-.538v-1.59h1.574v.44h-1.036v.625a.514.514 0 0 1 .222-.143l.004-.001a.578.578 0 0 1 .171-.026h.008a.642.642 0 0 1 .522.202a.741.741 0 0 1 .148.445l-.002.049v-.002zM16.308 4.48h.989v11.489h-.989zm-.118 14.874h-.344v.627h-.538v-.627h-.952v-.442l.952-1.843h.538v1.843h.344zm-2.574-.069v.007a.638.638 0 0 1-.269.521l-.002.001a.986.986 0 0 1-.596.199h-.028h.001h-.026a.918.918 0 0 1-.579-.205l.002.001a.616.616 0 0 1-.248-.49v-.072l.512-.032l.011.102a.262.262 0 0 0 .096.179h.001c.055.044.126.07.202.07h.015h-.001q.37 0 .37-.281v-.171q0-.447-.575-.447v-.376a.645.645 0 0 0 .46-.142l-.001.001a.469.469 0 0 0 .137-.332l-.001-.022v.001v-.067c0-.191-.102-.285-.307-.285l-.024-.001a.292.292 0 0 0-.226.106a.336.336 0 0 0-.06.194v.018v-.001v.069h-.526v-.067l-.001-.028c0-.186.074-.354.194-.478a.818.818 0 0 1 .619-.223h-.002l.044-.001c.216 0 .413.079.564.21l-.001-.001a.609.609 0 0 1 .205.457v.015v-.001v.126a.632.632 0 0 1-.199.462a.809.809 0 0 1-.194.139l-.005.002a.822.822 0 0 1 .368.328l.002.004a.525.525 0 0 1 .069.26v.017v-.001zm-2.12-3.315h-1.485V4.481h1.485zm-.41 3.582v.43H9.375v-.403a.998.998 0 0 1 .118-.161l-.001.001c.037-.045.078-.094.118-.148l.139-.164q.118-.137.385-.439c.137-.15.26-.318.364-.498l.007-.014a.75.75 0 0 0 .099-.352v-.005a.31.31 0 0 0-.091-.22a.352.352 0 0 0-.243-.097h-.014h.001h-.014a.326.326 0 0 0-.304.209l-.001.002a.408.408 0 0 0-.031.155v.013v-.001v.112h-.53v-.077l-.001-.047c0-.235.106-.445.273-.584l.001-.001a.95.95 0 0 1 .623-.192h-.002h.021c.222 0 .426.082.582.217l-.001-.001a.65.65 0 0 1 .25.504c-.001.212-.061.41-.165.578l.003-.005c-.132.198-.268.37-.416.53l.003-.003q-.087.101-.261.318l-.097.118c-.062.078-.111.139-.143.184a.141.141 0 0 0-.022.04v.001zM8.924 15.97h-1.98V4.481h1.98zm-.288 4.014h-.538v-2.237h-.459v-.306a.896.896 0 0 0 .373-.12l-.004.002a.633.633 0 0 0 .25-.242l.002-.003h.376zm-2.644-.253H4.507V4.481h1.485zm6.749-15.25h2.473V15.97h-2.473zm5.828 0h1.485V15.97h-1.485zM34.82.006H.48a.495.495 0 0 0-.495.495v23.01c0 .273.221.495.495.495h34.34a.495.495 0 0 0 .495-.495V.501a.495.495 0 0 0-.495-.495"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBasket(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M25.714 10.286h.027c.465 0 .885.192 1.185.502c.31.301.502.721.502 1.186v.028v-.001v.027c0 .465-.192.885-.502 1.185c-.301.31-.721.502-1.186.502h-.028h.001h-.201l-1.541 8.869a1.655 1.655 0 0 1-.587 1.015l-.003.002c-.29.249-.67.401-1.085.401H5.129c-.415 0-.795-.152-1.087-.403l.002.002a1.654 1.654 0 0 1-.588-1.008l-.001-.01l-1.541-8.869h-.228c-.465 0-.885-.192-1.185-.502a1.647 1.647 0 0 1-.502-1.186v-.054c0-.465.192-.885.502-1.185c.301-.31.721-.502 1.186-.502h.028h-.001zM6.491 21a.823.823 0 0 0 .581-.3l.001-.001a.823.823 0 0 0 .208-.626v.003l-.429-5.572a.823.823 0 0 0-.3-.581l-.001-.001a.823.823 0 0 0-.626-.208h.003a.823.823 0 0 0-.581.3l-.001.001a.823.823 0 0 0-.208.626v-.003l.429 5.572a.827.827 0 0 0 .274.562l.001.001a.83.83 0 0 0 .572.228h.011h-.001zm5.505-.856v-5.58a.816.816 0 0 0-.254-.594a.82.82 0 0 0-.595-.254h-.016a.816.816 0 0 0-.594.254a.82.82 0 0 0-.254.595v5.588c0 .234.097.445.254.594a.82.82 0 0 0 .595.254h.016a.816.816 0 0 0 .594-.254a.82.82 0 0 0 .258-.598v-.006zm5.143 0v-5.58a.816.816 0 0 0-.254-.594a.82.82 0 0 0-.595-.254h-.016a.816.816 0 0 0-.594.254a.82.82 0 0 0-.254.595v5.588c0 .234.097.445.254.594a.82.82 0 0 0 .595.254h.016a.816.816 0 0 0 .594-.254a.82.82 0 0 0 .258-.598v-.006zm4.714.066l.429-5.572a.83.83 0 0 0-.208-.624l.001.001a.826.826 0 0 0-.58-.302h-.003a.83.83 0 0 0-.624.208l.001-.001a.826.826 0 0 0-.302.58v.003l-.429 5.572a.83.83 0 0 0 .208.624l-.001-.001a.826.826 0 0 0 .58.302h.08a.834.834 0 0 0 .573-.228a.83.83 0 0 0 .279-.564v-.003zM6.375 3.91L5.13 9.428H3.362l1.352-5.906c.163-.785.59-1.45 1.182-1.915l.007-.005A3.332 3.332 0 0 1 8.01.858h.038h-.002h2.236V.85c0-.234.097-.445.254-.594a.821.821 0 0 1 .596-.255h5.162c.234 0 .445.097.594.254a.82.82 0 0 1 .254.595v.008h2.273c.801 0 1.536.28 2.112.748l-.006-.005A3.262 3.262 0 0 1 22.706 3.5l.004.021l1.352 5.906h-1.768l-1.245-5.518a1.767 1.767 0 0 0-.606-.962l-.003-.003a1.638 1.638 0 0 0-1.046-.375h-2.255v.008a.816.816 0 0 1-.254.594a.82.82 0 0 1-.595.254h-5.158a.816.816 0 0 1-.594-.254a.82.82 0 0 1-.254-.595v-.008H8.029a1.64 1.64 0 0 0-1.048.377l.003-.002a1.76 1.76 0 0 0-.603.954l-.002.012z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBasketAdd(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.478 22.099v-.091H13.75v.091l.001.045a1.9 1.9 0 1 1-3.802 0l.001-.047v-.089h-.653a1.22 1.22 0 0 1-1.192-.96l-.001-.008L4.287 2.967l-3.261-.543A1.221 1.221 0 0 1 1.43.017L1.424.016l4.081.68c.497.085.887.461.991.943l.001.008l.692 3.264l11.905.992a6.51 6.51 0 0 0-.112 1.186v.004c0 .402.035.796.103 1.179l-.006-.04a1.397 1.397 0 0 0-.951 1.601l-.001-.009l.719 5.19c.044.333.196.625.418.845c.202.201.479.327.786.33h.001c.068 0 .135-.007.199-.019l-.007.001a1.389 1.389 0 0 0 1.013-1.605l.001.009l-.392-2.822a6.657 6.657 0 0 0 4.808 2.044c.043 0 .086 0 .128-.006l.093-.005l-.669 3.902a.407.407 0 0 1-.4.339H9.96l.327 1.547h14.157a1.221 1.221 0 0 1 .002 2.44H23.28v.091a1.9 1.9 0 1 1-3.802 0zm1.358 0a.539.539 0 1 0 1.07-.094v.003h-1.06a.514.514 0 0 0-.008.091zm-9.528 0a.539.539 0 1 0 1.07-.094v.003h-1.061a.494.494 0 0 0-.007.089v.002zm3.867-13.875a1.388 1.388 0 0 0-1.013 1.605l-.001-.009l.719 5.19c.044.333.196.625.418.845c.202.201.479.327.786.33h.001c.068 0 .135-.007.199-.019l-.007.001a1.389 1.389 0 0 0 1.013-1.605l.001.009l-.719-5.19a1.452 1.452 0 0 0-.418-.845a1.132 1.132 0 0 0-.786-.331h-.001c-.069 0-.136.007-.201.02l.007-.001zm-3.966 0a1.387 1.387 0 0 0-1.014 1.609l-.001-.009l.719 5.19c.044.333.196.625.418.845c.202.201.479.327.786.33h.001c.068 0 .135-.007.199-.019l-.007.001a1.388 1.388 0 0 0 1.014-1.605l.001.009l-.72-5.193a1.452 1.452 0 0 0-.418-.845a1.128 1.128 0 0 0-.786-.331H11.4c-.068 0-.134.007-.198.019zm8.789-1.128a5.678 5.678 0 1 1 11.36.002a5.678 5.678 0 0 1-11.362-.003zm1.627 0a4.051 4.051 0 1 0 8.106 0a4.051 4.051 0 0 0-8.108-.001zm3.24 1.621v-.808h-.808a.814.814 0 0 1 0-1.628h.808v-.808a.814.814 0 0 1 1.628 0v.808h.808a.814.814 0 0 1 0 1.628h-.808v.807a.814.814 0 0 1-1.628 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingBasketRemove(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.479 22.099v-.091H13.75v.091a1.9 1.9 0 1 1-3.802 0v-.091h-.651a1.222 1.222 0 0 1-1.193-.96l-.001-.008L4.286 2.967l-3.261-.543A1.221 1.221 0 0 1 1.429.017L1.423.016l4.081.68c.497.085.887.461.991.943l.001.008l.692 3.264l11.904.992c-.07.356-.11.766-.112 1.186v.002c0 .381.032.754.092 1.118l-.005-.039c-.29.112-.523.32-.663.585l-.003.007a1.536 1.536 0 0 0-.17.965l-.001-.008l.719 5.19c.044.333.196.625.418.845c.202.201.48.326.786.33h.001c.068 0 .135-.007.199-.019l-.007.001a1.389 1.389 0 0 0 1.013-1.605l.001.009l-.36-2.592a6.634 6.634 0 0 0 4.674 1.917h.002c.043 0 .086 0 .128-.006h.092l-.669 3.902a.407.407 0 0 1-.4.338H9.958l.327 1.547h14.157l.036-.001a1.221 1.221 0 0 1 0 2.442l-.038-.001h.002h-1.164v.091a1.9 1.9 0 0 1-3.802.002v-.002zm1.358 0a.54.54 0 1 0 1.071-.094v.003h-1.063a.494.494 0 0 0-.007.089v.002zm-9.528 0a.539.539 0 1 0 1.07-.094v.003h-1.062a.494.494 0 0 0-.007.089v.002zm3.972-13.972a1.387 1.387 0 0 0-1.016 1.601l-.001-.009l.719 5.19c.044.333.196.625.418.845c.202.201.479.327.786.33h.001c.068 0 .135-.007.199-.019l-.007.001a1.388 1.388 0 0 0 1.014-1.605l.001.009l-.719-5.19a1.452 1.452 0 0 0-.418-.845a1.128 1.128 0 0 0-.986-.308l.007-.001zm-3.967 0A1.387 1.387 0 0 0 10.3 9.728l-.001-.009l.719 5.19c.044.333.196.625.418.845c.202.201.48.326.786.33h.001c.068 0 .135-.007.199-.019l-.007.001a1.389 1.389 0 0 0 1.013-1.605l.001.009l-.72-5.185a1.452 1.452 0 0 0-.418-.845a1.128 1.128 0 0 0-.786-.331h-.003a1.05 1.05 0 0 0-.196.018l.007-.001zm8.684-1.032a5.678 5.678 0 1 1 11.36 0a5.678 5.678 0 0 1-11.362 0zm1.627 0a4.052 4.052 0 1 0 8.108 0a4.052 4.052 0 0 0-8.11 0zm2.43.813a.814.814 0 0 1 0-1.628h3.24a.814.814 0 0 1 0 1.628z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingPackage(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.416 24V12.902h5.68c.181 0 .328.147.328.328v10.114a.656.656 0 0 1-.656.656zM2.32 24a.656.656 0 0 1-.656-.656V13.23c0-.181.147-.328.328-.328h5.621V24zM.328 11.92A.328.328 0 0 1 0 11.592V7.561c0-.181.147-.328.328-.328h6.546C2.96 6.223 1.6 4.178 1.529 4.069a.656.656 0 0 1 1.099-.717l.002.002c.028.041 1.342 1.92 5.15 2.74c-1.273-.64-2.518-1.529-2.847-2.673A2.543 2.543 0 0 1 5.718.889l.003-.002a2.435 2.435 0 0 1 1.755-.888h.006c1.714 0 2.904 2.391 3.583 4.309c.749-1.87 2.037-4.252 3.74-4.252c.741.039 1.388.41 1.799.966l.005.006a2.54 2.54 0 0 1 .666 2.585l.005-.018c-.352 1.035-1.466 1.823-2.653 2.391c3.472-.872 4.675-2.61 4.69-2.633a.656.656 0 0 1 1.098.717l.002-.003c-.07.11-1.434 2.154-5.345 3.164h6.48c.181 0 .328.147.328.328v4.029a.328.328 0 0 1-.328.328zM6.677 1.788c-.65.69-.524 1.127-.48 1.27c.298 1.035 2.268 2.018 3.936 2.596c-.871-2.955-2.053-4.342-2.65-4.342a1.304 1.304 0 0 0-.804.473zm5.315 3.791c1.692-.501 3.698-1.389 4.043-2.406c.048-.142.194-.572-.422-1.291a1.189 1.189 0 0 0-.801-.513l-.007-.001c-.946 0-2.103 2.226-2.813 4.21z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingPosMachine(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.44 24A1.44 1.44 0 0 1 0 22.56V5.656c0-.795.645-1.44 1.44-1.44h1.706c.054 0 .106.013.151.036l-.002-.001V.168c0-.093.075-.168.168-.168h6.02c.093 0 .168.075.168.168v4.059a.337.337 0 0 1 .08-.01h1.515c.795 0 1.44.645 1.44 1.44v4.031a.326.326 0 0 1 .145-.033h.757c.186 0 .337.151.337.337v8.772a.337.337 0 0 1-.337.337h-.757a.33.33 0 0 1-.147-.034l.002.001v3.492a1.44 1.44 0 0 1-1.44 1.44zm7.324-3.231v1.134a.17.17 0 0 0 .169.169h1.334a.17.17 0 0 0 .169-.169v-1.134a.169.169 0 0 0-.169-.168H8.931a.17.17 0 0 0-.169.168zm-3.156 0v1.134a.17.17 0 0 0 .169.169H7.11a.17.17 0 0 0 .169-.169v-1.134a.169.169 0 0 0-.169-.168H5.776a.17.17 0 0 0-.169.168zm-3.156 0v1.135c0 .093.075.168.168.168h1.334a.17.17 0 0 0 .169-.169v-1.134a.169.169 0 0 0-.169-.168H2.62a.168.168 0 0 0-.168.168m6.311-2.571v1.134a.17.17 0 0 0 .169.169h1.334a.17.17 0 0 0 .169-.169v-1.135a.168.168 0 0 0-.168-.168H8.93a.168.168 0 0 0-.168.168zm-3.156 0v1.134a.17.17 0 0 0 .169.169H7.11a.17.17 0 0 0 .169-.169v-1.135a.168.168 0 0 0-.168-.168H5.775a.168.168 0 0 0-.168.168zm-3.156 0v1.135c0 .093.075.168.168.168h1.334a.17.17 0 0 0 .169-.169v-1.135a.168.168 0 0 0-.168-.168H2.619a.168.168 0 0 0-.168.168zm6.311-2.572v1.135c0 .093.075.168.168.168h1.336a.168.168 0 0 0 .168-.168v-1.135a.169.169 0 0 0-.169-.168H8.931a.17.17 0 0 0-.169.168m-3.156 0v1.135c0 .093.075.168.168.168H7.11a.168.168 0 0 0 .168-.168v-1.135a.169.169 0 0 0-.169-.168H5.775a.17.17 0 0 0-.169.168m-3.156 0v1.135c0 .093.075.168.168.168h1.336a.168.168 0 0 0 .168-.168v-1.135a.169.169 0 0 0-.169-.168H2.619a.168.168 0 0 0-.168.168zm-.21-6.713v4.912c0 .093.075.168.168.168h7.762a.168.168 0 0 0 .168-.168V8.913a.17.17 0 0 0-.169-.169h-7.76a.17.17 0 0 0-.169.169m-.504-2.682v1.19c0 .279.226.505.505.505h8.399a.505.505 0 0 0 .505-.505v-1.19a.505.505 0 0 0-.505-.505h-.99v1.01h.488v.178H2.746v-.178h.321a.505.505 0 0 0 .233-.057l-.003.001v-.898a.499.499 0 0 0-.231-.056h-.827a.505.505 0 0 0-.504.505z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingSale(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.666 23.111c-.25-.28-.534-.598-.75-.64a1.94 1.94 0 0 0-.98.315l.007-.005a2.845 2.845 0 0 1-1.154.378l-.012.001h-.01a.954.954 0 0 1-.367-.073l.006.002c-.491-.204-.69-.772-.88-1.322a2.058 2.058 0 0 0-.457-.894l.002.002a1.051 1.051 0 0 0-.557-.101h.004c-.146 0-.3.009-.454.018c-.166.01-.332.018-.49.018a1.447 1.447 0 0 1-1.061-.328l.002.002a1.956 1.956 0 0 1-.305-1.561l-.002.013a2.056 2.056 0 0 0-.084-1.021l.004.015a2.038 2.038 0 0 0-.878-.453l-.014-.003c-.55-.191-1.12-.389-1.322-.88a1.95 1.95 0 0 1 .313-1.543l-.004.006a2.05 2.05 0 0 0 .307-.973v-.005c-.042-.217-.36-.5-.64-.75c-.438-.388-.89-.791-.89-1.33s.452-.942.889-1.333c.28-.25.598-.534.64-.75a2.062 2.062 0 0 0-.315-.981l.005.008a1.953 1.953 0 0 1-.305-1.551l-.002.013c.204-.491.772-.69 1.322-.88c.346-.081.647-.24.894-.458l-.002.002a2.024 2.024 0 0 0 .078-1.019l.002.012a1.969 1.969 0 0 1 .31-1.555l-.004.006a1.45 1.45 0 0 1 1.064-.324H4.57c.16 0 .325.009.49.018c.154.009.308.018.454.018a1.052 1.052 0 0 0 .559-.103l-.006.003c.215-.246.374-.546.452-.878l.003-.014c.192-.55.39-1.12.88-1.322a.937.937 0 0 1 .361-.07h.011h-.001c.436.045.832.18 1.181.386L8.94 1.22a1.9 1.9 0 0 0 .969.311h.004c.217-.042.5-.36.75-.64C11.057.452 11.461 0 12 0s.942.452 1.333.89c.25.28.534.598.75.64a1.96 1.96 0 0 0 .98-.316l-.008.005A2.839 2.839 0 0 1 16.21.842l.012-.001h.01c.13 0 .254.026.367.073l-.006-.002c.491.204.69.772.88 1.322c.081.346.24.646.457.894l-.002-.002a1.043 1.043 0 0 0 .558.099h-.004c.146 0 .3-.009.454-.017c.166-.009.332-.018.49-.018a1.452 1.452 0 0 1 1.062.328l-.002-.002a1.962 1.962 0 0 1 .304 1.561l.002-.013a2.045 2.045 0 0 0 .085 1.021l-.004-.014c.245.216.546.375.878.452l.014.003c.55.192 1.12.39 1.322.88a1.962 1.962 0 0 1-.312 1.543l.004-.006a2.04 2.04 0 0 0-.31.967v.006c.042.217.36.5.64.75c.441.393.893.796.893 1.335s-.452.943-.89 1.334c-.28.25-.598.533-.64.75c.022.363.135.696.315.982l-.005-.009a1.95 1.95 0 0 1 .305 1.55l.002-.013c-.203.492-.772.69-1.322.88a2.05 2.05 0 0 0-.894.458l.002-.002a2.024 2.024 0 0 0-.078 1.019l-.002-.012a1.969 1.969 0 0 1-.31 1.554l.004-.006a1.45 1.45 0 0 1-1.063.326h.006c-.16 0-.325-.009-.49-.018a7.932 7.932 0 0 0-.454-.018a1.045 1.045 0 0 0-.559.103l.006-.003a2.047 2.047 0 0 0-.452.878l-.003.014c-.192.55-.39 1.12-.88 1.322a.937.937 0 0 1-.361.07h-.011h.001a2.877 2.877 0 0 1-1.18-.387l.013.007a1.89 1.89 0 0 0-.969-.31h-.004c-.217.042-.5.36-.75.64c-.395.438-.798.89-1.338.89s-.942-.452-1.334-.889zm3.532-15.802l-4.88 8.868a.166.166 0 0 0 .145.246h.847a.164.164 0 0 0 .144-.085v-.001l4.873-8.869a.166.166 0 0 0-.145-.246h-.839a.166.166 0 0 0-.145.086zm-.704 4.549a1.119 1.119 0 0 0-.406.863v2.401c0 .344.158.651.405.853l.002.002c.256.214.588.343.951.343s.695-.13.953-.345l-.002.002c.25-.203.408-.51.408-.854v-2.4c0-.344-.155-.653-.398-.859l-.002-.001a1.378 1.378 0 0 0-.923-.353h-.031h.002h-.031a1.4 1.4 0 0 0-.928.35l.002-.001zM9.25 7.551a1.119 1.119 0 0 0-.406.863v2.4c0 .345.157.654.405.857l.002.002c.256.217.589.348.953.348s.698-.132.955-.35l-.002.002c.252-.205.411-.514.411-.861V8.418c0-.347-.157-.656-.405-.862l-.002-.001a1.475 1.475 0 0 0-1.914.001l.002-.002zm4.978 7.701a.183.183 0 0 1-.069-.143v-2.401c0-.058.027-.109.068-.142a.286.286 0 0 1 .187-.07h.013h-.001h.012a.29.29 0 0 1 .188.069a.183.183 0 0 1 .069.142v2.402a.182.182 0 0 1-.068.142c-.054.044-.124.07-.2.07s-.146-.026-.201-.071h.001zm-4.24-4.306a.185.185 0 0 1-.069-.143v-2.4c0-.058.027-.109.068-.143c.054-.044.124-.07.2-.07s.146.026.201.071h-.001a.183.183 0 0 1 .069.143v2.401a.185.185 0 0 1-.068.143c-.054.044-.124.07-.2.07s-.146-.026-.201-.071z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ShoppingStore(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.482 24a.393.393 0 0 1-.389-.393v-1.783c0-.217.176-.393.393-.393h1.347V8.671a.79.79 0 0 1 .039-.244l-.002.006a3.265 3.265 0 0 1-1.8-1.287l-.007-.011A.39.39 0 0 1 0 6.921V.394C0 .177.176.001.393 0h27.014c.217 0 .393.176.393.393V6.92a.396.396 0 0 1-.064.215l.001-.002a3.49 3.49 0 0 1-2.687 1.468h-.008V21.43h1.091c.217 0 .393.176.393.393v1.783a.393.393 0 0 1-.393.393h-.001zm22.99-2.965V8.673c0-.07.009-.138.026-.202l-.001.006a3.498 3.498 0 0 1-1.53-.836l.002.002a3.82 3.82 0 0 1-2.68.976h.006a3.811 3.811 0 0 1-2.69-.996l.003.003a3.81 3.81 0 0 1-2.574.995l-.119-.002h.006a3.816 3.816 0 0 1-2.69-.996l.003.003a3.812 3.812 0 0 1-2.575.995l-.128-.002h.006a3.814 3.814 0 0 1-2.69-.996l.003.003a3.742 3.742 0 0 1-2.441.993h-.005v12.417H4.56v-5.253c0-.217.176-.393.393-.393h4.651c.217 0 .393.176.393.393v5.253zM19.295 7.833a3.028 3.028 0 0 0 2.2-.835l-.001.001a.364.364 0 0 1-.009-.083V.787h-4.41V6.92l.001.027l-.001.028v-.001a3.025 3.025 0 0 0 2.117.861l.11-.002h-.005zm-10.758 0l.106.002c.824 0 1.571-.328 2.118-.861l-.001.001l-.001-.027l.001-.028v.001V.788H6.312v6.133l.001.027l-.001.028v-.001a3.025 3.025 0 0 0 2.117.861l.112-.002h-.005zm2.873 10.495a.393.393 0 0 1-.393-.393v-6.821c0-.217.176-.393.393-.393h11.009c.217 0 .393.176.393.393v6.822a.393.393 0 0 1-.393.393h-.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SinaWeibo(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.44 22.251c-4.898.482-9.131-1.732-9.449-4.951s3.398-6.216 8.301-6.701s9.13 1.724 9.448 4.949s-3.398 6.219-8.298 6.699zm-1.291-3.818a1.858 1.858 0 0 1-2.265.735l.012.004a1.303 1.303 0 0 1-.497-1.966l-.002.003a1.861 1.861 0 0 1 2.221-.735l-.013-.004a1.306 1.306 0 0 1 .542 1.965l.003-.004zm1.564-2.004a.712.712 0 0 1-.854.31l.005.002a.48.48 0 0 1-.217-.723l-.001.002a.684.684 0 0 1 .832-.294l-.005-.002a.494.494 0 0 1 .221.741l.001-.002zm.217-3.349a5.34 5.34 0 0 0-5.969 2.581l-.014.027a3.641 3.641 0 0 0 2.295 5.173l.025.006a5.295 5.295 0 0 0 6.306-2.653l.014-.03a3.657 3.657 0 0 0-2.634-5.105l-.023-.004zm9.315-1.507c-.426-.123-.702-.222-.499-.757c.233-.428.371-.937.371-1.478s-.137-1.05-.379-1.494l.008.016c-.962-1.37-3.59-1.297-6.607-.037c0 0-.943.408-.703-.334c.463-1.499.388-2.739-.333-3.46c-1.65-1.657-5.999.046-9.718 3.784c-2.774 2.796-4.386 5.756-4.386 8.311c0 4.903 6.281 7.883 12.422 7.883c8.05 0 13.41-4.68 13.41-8.4c0-2.244-1.905-3.515-3.59-4.045v.012zm2.35-6.272a3.812 3.812 0 0 0-3.672-1.18l.025-.005a.965.965 0 1 0 .396 1.889l-.006.001A1.86 1.86 0 0 1 23.5 8.407l.004-.013a.97.97 0 0 0 .622 1.22l.007.002a1.002 1.002 0 0 0 1.22-.624l.002-.007c.118-.352.186-.757.186-1.179c0-.986-.373-1.884-.985-2.563l.003.003l.037.055zm2.978-2.703A7.814 7.814 0 0 0 21.749.002a7.85 7.85 0 0 0-1.684.182l.052-.01a1.126 1.126 0 0 0 .47 2.202l-.007.001a5.556 5.556 0 0 1 6.446 7.183l.011-.039a1.13 1.13 0 0 0 2.145.712l.002-.008c.24-.72.378-1.548.378-2.41a7.84 7.84 0 0 0-2.032-5.277l.006.007l.037.055z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sitemap(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.636 17.455V10.91H12V6.546h2.182V.001H7.637v6.545h2.182v4.364H2.183v6.545H.001V24h6.545v-6.545H4.364v-4.364h13.091v4.364h-2.182V24h6.545v-6.545z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Skyatlas(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m24.001 14.561l.001-.084a3.73 3.73 0 0 0-1.228-2.77l-.003-.003a4.216 4.216 0 0 0-2.895-1.144l-.099.001h.005h-.01c-.48 0-.945.064-1.387.185l.037-.009c.056-.227.088-.487.088-.754v-.039v.002a4.908 4.908 0 0 0-1.493-3.669l-.001-.001a5.09 5.09 0 0 0-3.591-1.475l-.153.002h.007h-.002c-.768 0-1.5.153-2.167.431l.037-.014a4.6 4.6 0 0 0-1.732 1.205l-.004.004a1.062 1.062 0 0 0-.263.543l-.001.006l-.001.026c0 .109.042.208.111.283c.06.066.145.107.24.11a1.36 1.36 0 0 0 .53-.266l-.002.002a5.628 5.628 0 0 1 1.064-.689l.033-.015a3.996 3.996 0 0 1 1.748-.396h.012h-.001l.102-.001a4.35 4.35 0 0 1 3.019 1.212L16 7.242a4.07 4.07 0 0 1 1.275 3.106v-.007a3.328 3.328 0 0 1-.115.705l.005-.023c-.057.2-.096.433-.109.673v.009c.001.158.068.3.176.4c.1.108.242.176.4.176c.275-.027.529-.089.766-.183l-.019.007a3.283 3.283 0 0 1 1.043-.263l.013-.001l.076-.001c.735 0 1.395.32 1.849.829l.002.002a2.57 2.57 0 0 1 .682 1.748v.034v-.002a2.336 2.336 0 0 1-.682 1.78a2.624 2.624 0 0 1-2.073.835h.007a4.913 4.913 0 0 1-3.169-1.278l.004.004a24.283 24.283 0 0 1-2.553-2.802l-.041-.054q-1.318-1.626-2.066-2.418a11.517 11.517 0 0 0-2.449-1.861l-.057-.03a6.75 6.75 0 0 0-3.195-.791h-.061h.003h-.071a5.77 5.77 0 0 0-2.838.741l.029-.015A5.568 5.568 0 0 0 .78 10.537l-.014.025A5.363 5.363 0 0 0 0 13.423v-.004l-.001.087a5.06 5.06 0 0 0 .804 2.746l-.012-.02a5.639 5.639 0 0 0 2.08 1.92l.03.015a5.858 5.858 0 0 0 2.802.703h.035h-.002a8.074 8.074 0 0 0 3.129-.66l-.052.02a10.095 10.095 0 0 0 2.762-1.66l-.015.012l.044-.044q.659-.571.659-.88v-.016c0-.105-.042-.2-.11-.27a.386.386 0 0 0-.27-.11h-.016h.001q-.176 0-.615.4l-.264.22a8.376 8.376 0 0 1-2.18 1.125l-.06.018a7.273 7.273 0 0 1-2.412.439h-.005a4.388 4.388 0 0 1-3.256-1.234l.001.001A3.979 3.979 0 0 1 2 13.503v-.021v.001a3.903 3.903 0 0 1 .992-2.751l-.003.004a4.041 4.041 0 0 1 3.13-1.23h-.01c.985 0 1.91.261 2.708.717l-.027-.014a8.425 8.425 0 0 1 2.102 1.706l.007.009q.659.703 1.89 2.264a27.744 27.744 0 0 0 1.902 2.232l-.012-.012a8.687 8.687 0 0 0 2.113 1.738l.041.022a5.493 5.493 0 0 0 2.705.703h.022h-.001l.09.001a4.382 4.382 0 0 0 3.054-1.233l-.001.001A4.056 4.056 0 0 0 24 14.661l-.001-.107v.005z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Skype(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M18.328 14.61v-.039c0-.504-.112-.982-.314-1.41l.009.02a3.463 3.463 0 0 0-.756-1.069l-.001-.001a4.837 4.837 0 0 0-1.112-.753l-.029-.013a8.335 8.335 0 0 0-1.227-.514l-.061-.018q-.602-.187-1.367-.359l-1.629-.374q-.469-.11-.687-.164a4.885 4.885 0 0 1-.581-.192l.034.012a1.569 1.569 0 0 1-.472-.252l.003.003a1.313 1.313 0 0 1-.254-.321l-.003-.007a.904.904 0 0 1-.117-.449v-.021v.001q0-1.2 2.25-1.2h.053c.41 0 .805.069 1.172.195l-.025-.008c.323.106.603.258.85.45l-.006-.005q.313.258.594.523c.182.172.387.322.609.445l.015.008c.213.118.466.187.736.187h.015h-.001l.06.001c.444 0 .844-.193 1.119-.5l.001-.001c.277-.308.446-.717.446-1.166v-.036v.002a1.986 1.986 0 0 0-.868-1.55l-.007-.004A5.98 5.98 0 0 0 14.6 4.985l-.041-.009a10.713 10.713 0 0 0-2.784-.359h-.063h.003h-.029c-.722 0-1.423.088-2.093.253l.06-.012a7.175 7.175 0 0 0-1.902.753l.035-.019a3.76 3.76 0 0 0-1.915 3.372v-.005l-.001.092c0 .566.109 1.106.308 1.601l-.01-.029c.196.473.495.87.87 1.176l.005.004c.361.295.773.549 1.216.744l.034.013c.451.194.989.369 1.545.496l.065.012l2.282.56c.66.128 1.247.321 1.797.58l-.046-.019c.302.184.501.511.501.886l-.001.055v-.003a1.171 1.171 0 0 1-.619 1.005l-.006.003a2.982 2.982 0 0 1-1.647.4h.007h-.062c-.492 0-.962-.092-1.394-.259l.026.009a3.264 3.264 0 0 1-1.02-.605l.004.003q-.383-.351-.711-.703a3.766 3.766 0 0 0-.705-.593l-.015-.009a1.526 1.526 0 0 0-.839-.25H7.45a1.46 1.46 0 0 0-1.179.467l-.001.001a1.734 1.734 0 0 0-.4 1.175v-.003q0 1.44 1.906 2.461a9.36 9.36 0 0 0 4.56 1.023h-.013h.032c.767 0 1.509-.105 2.212-.303l-.058.014a7.122 7.122 0 0 0 1.944-.853l-.029.017a4.072 4.072 0 0 0 1.373-1.441l.011-.02a4.06 4.06 0 0 0 .523-2.006v-.056v.003zM24 18a6 6 0 0 1-6 6h-.044a5.821 5.821 0 0 1-3.624-1.259l.012.009c-.703.156-1.512.247-2.342.25h-.033c-1.533 0-2.991-.319-4.312-.894l.07.027c-2.665-1.117-4.742-3.194-5.833-5.788l-.026-.071a10.634 10.634 0 0 1-.867-4.242v-.034V12c.003-.832.094-1.641.264-2.42l-.014.076A5.806 5.806 0 0 1 0 6.043v-.045V6a6 6 0 0 1 6-6h.043c1.37 0 2.629.471 3.625 1.26l-.012-.009a11.183 11.183 0 0 1 2.342-.25h.033c1.533 0 2.991.319 4.312.894l-.07-.027c2.665 1.117 4.742 3.194 5.833 5.788l.026.071c.548 1.251.867 2.709.867 4.242v.032v-.002a11.598 11.598 0 0 1-.263 2.42l.014-.076A5.808 5.808 0 0 1 24 17.955zv-.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Slack(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.909 11.192l.058-.001c.559 0 1.065.223 1.435.585c.37.357.599.857.599 1.411l-.001.056v-.003a1.841 1.841 0 0 1-1.329 1.872l-.013.003l-2.48.851l.808 2.409c.064.194.101.417.101.649v.03v-.002v.023c0 .567-.232 1.079-.606 1.448a1.96 1.96 0 0 1-1.432.62h-.026h.001h-.015c-.457 0-.88-.146-1.224-.394l.006.004a2.023 2.023 0 0 1-.768-1.026l-.004-.014l-.794-2.38l-4.471 1.529l.794 2.366c.071.202.113.435.115.677v.024c0 .567-.232 1.079-.606 1.448a1.975 1.975 0 0 1-1.442.623h-.032h.002h-.018c-.454 0-.873-.146-1.214-.394l.006.004a2.04 2.04 0 0 1-.761-1.026l-.004-.014l-.794-2.351l-2.202.766a2.41 2.41 0 0 1-.716.129h-.004l-.058.001a1.994 1.994 0 0 1-1.406-.578a1.957 1.957 0 0 1-.585-1.397l.001-.062v-.009c0-.457.15-.879.404-1.219l-.004.005c.253-.353.605-.62 1.017-.76l.015-.004l2.25-.765l-1.514-4.514l-2.25.779a2.313 2.313 0 0 1-.69.115h-.002l-.047.001a1.983 1.983 0 0 1-1.41-.585a1.951 1.951 0 0 1-.601-1.411v-.051c0-.457.15-.879.404-1.219l-.004.005c.253-.353.605-.62 1.017-.76l.015-.004l2.261-.762l-.764-2.295a2.116 2.116 0 0 1-.115-.677v-.034c0-.566.232-1.077.606-1.445a1.987 1.987 0 0 1 1.437-.613h.036h-.002h.018c.454 0 .873.146 1.214.394l-.006-.004c.356.253.624.609.761 1.026l.004.014l.779 2.308L12.13 5.08l-.779-2.308a2.116 2.116 0 0 1-.115-.677V2.07c0-.569.235-1.082.612-1.45c.363-.382.875-.62 1.442-.62h.023h-.001a2.128 2.128 0 0 1 2 1.413l.005.015l.765 2.32l2.336-.788a2.19 2.19 0 0 1 .617-.086h.045c.556 0 1.062.217 1.437.571l-.001-.001c.377.342.613.833.613 1.38v.043v-.002v.002c0 .449-.164.859-.435 1.175l.002-.002a2.34 2.34 0 0 1-1.051.737l-.016.005l-2.265.779l1.514 4.56l2.366-.808c.197-.072.425-.116.662-.121zm-11.452 3.779l4.471-1.514l-1.514-4.543l-4.471 1.543z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Slides(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 0v24h24V0zm14.176 15.794a3.294 3.294 0 0 1-2.237.871l-.077-.001h.004l-.133.001c-.788 0-1.541-.15-2.232-.422l.041.014c-.646-.27-.976-.525-.976-.78l-.076.1c.034-.223.12-.422.246-.588l-.002.003q.245-.39.44-.39c.121.058.224.12.321.189l-.007-.005c.24.144.517.271.809.364l.028.008c.367.118.789.186 1.227.186l.095-.001h-.005l.079.002a1.83 1.83 0 0 0 1.198-.444l-.002.002c.306-.277.497-.676.497-1.119l-.002-.074v.003a1.56 1.56 0 0 0-.437-1.171a4.304 4.304 0 0 0-1.712-.785l-.028-.005a4.468 4.468 0 0 1-1.896-.951l.006.005a2.35 2.35 0 0 1-.586-1.762l-.001.009v-.037c0-.772.348-1.462.896-1.922l.004-.003a3.336 3.336 0 0 1 2.349-.792h-.008a5.413 5.413 0 0 1 1.824.265l-.039-.011c.524.165.78.36.78.555a1.367 1.367 0 0 1-.229.649l.003-.005c-.15.27-.285.39-.39.39c-.017 0-.09-.045-.257-.12a2.983 2.983 0 0 0-1.431-.36h-.056h.003a2.156 2.156 0 0 0-1.298.335l.009-.005a1.06 1.06 0 0 0-.45.868v.018v-.001l-.001.046c0 .317.132.604.345.807c.331.257.722.453 1.147.565l.023.005c.96.225 1.798.653 2.5 1.238l-.01-.008a2.794 2.794 0 0 1 .66 2.035l.001-.01l.001.092c0 .85-.37 1.614-.957 2.139l-.003.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SlightlySmile(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.316C6.099 1.316 1.316 6.099 1.316 12S6.099 22.684 12 22.684S22.684 17.901 22.684 12c-.012-5.896-4.788-10.672-10.683-10.684zm0 22.297C5.586 23.613.387 18.414.387 12S5.586.387 12 .387S23.613 5.586 23.613 12v.015c0 6.405-5.192 11.597-11.597 11.597z"/><path fill="currentColor" d="M12 24C5.386 23.966.034 18.614 0 12.003V12C0 5.373 5.373 0 12 0s12 5.373 12 12c-.034 6.614-5.386 11.966-11.997 12zM12 .774C5.8.774.774 5.8.774 12S5.8 23.226 12 23.226S23.226 18.2 23.226 12C23.222 5.802 18.198.779 12.001.774zm0 22.297C5.886 23.071.929 18.114.929 12S5.886.929 12 .929S23.071 5.886 23.071 12S18.114 23.071 12 23.071m0-21.368C6.313 1.703 1.703 6.313 1.703 12S6.313 22.297 12 22.297S22.297 17.687 22.297 12v-.005c0-5.684-4.608-10.292-10.292-10.292z"/><path fill="currentColor" d="M9.677 9.91v.009c0 1.15-.932 2.082-2.082 2.082h-.018a2.082 2.082 0 0 1-2.082-2.082V9.91a2.09 2.09 0 0 1 4.18 0zm8.904 0v.009c0 1.15-.932 2.082-2.082 2.082h-.018a2.082 2.082 0 0 1-2.082-2.082v-.018c0-1.15.932-2.082 2.082-2.082h.009a2.132 2.132 0 0 1 2.09 2.088v.002zM12 18.194h-.013a3.42 3.42 0 0 1-3.231-2.299l-.007-.024a.441.441 0 0 1 .85-.235l.001.003a2.542 2.542 0 0 0 4.794.017l.006-.018a.44.44 0 0 1 .545-.309l-.003-.001a.44.44 0 0 1 .308.545l.001-.003a3.427 3.427 0 0 1-3.238 2.323h-.014z"/><path fill="currentColor" d="M12 18.581a3.914 3.914 0 0 1-3.63-2.605l-.008-.027a.857.857 0 0 1 1.624-.547l.002.006a2.174 2.174 0 0 0 2.011 1.471h.002a2.17 2.17 0 0 0 2.008-1.456l.005-.015a.795.795 0 0 1 .383-.462l.004-.002a.858.858 0 0 1 .626-.076l-.006-.001a.839.839 0 0 1 .541 1.09l.002-.006a3.749 3.749 0 0 1-3.561 2.632h-.001zm0-.775h.007a3 3 0 0 0 2.851-2.069l.006-.021v-.077h-.077a2.937 2.937 0 0 1-2.783 2.013h-.028a2.906 2.906 0 0 1-2.758-1.992l-.006-.021s-.077-.077-.077 0c0 0-.077.077 0 .077a3.004 3.004 0 0 0 2.857 2.09z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Smiley(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.316C6.099 1.316 1.316 6.099 1.316 12S6.099 22.684 12 22.684S22.684 17.901 22.684 12c-.012-5.896-4.788-10.672-10.683-10.684zm0 22.297C5.586 23.613.387 18.414.387 12S5.586.387 12 .387S23.613 5.586 23.613 12v.015c0 6.405-5.192 11.597-11.597 11.597z"/><path fill="currentColor" d="M12 24C5.373 24 0 18.627 0 12S5.373 0 12 0s12 5.373 12 12c-.034 6.614-5.386 11.966-11.997 12zM12 .774C5.8.774.774 5.8.774 12S5.8 23.226 12 23.226S23.226 18.2 23.226 12C23.222 5.802 18.198.779 12.001.774zm0 22.374C5.886 23.148.929 18.191.929 12.077S5.886 1.006 12 1.006s11.071 4.957 11.071 11.071S18.114 23.148 12 23.148m0-21.445C6.313 1.703 1.703 6.313 1.703 12S6.313 22.297 12 22.297S22.297 17.687 22.297 12v-.005c0-5.684-4.608-10.292-10.292-10.292z"/><path fill="currentColor" d="M18.116 14.245v.036a6.08 6.08 0 0 1-6.08 6.08h-.038H12h-.036a6.08 6.08 0 0 1-6.08-6.08v-.038v.002zM9.677 9.91v.009c0 1.15-.932 2.082-2.082 2.082h-.009a2.09 2.09 0 1 1 2.09-2.09zm8.904 0a2.09 2.09 0 0 1-4.18 0v-.009c0-1.15.932-2.082 2.082-2.082h.009a2.132 2.132 0 0 1 2.09 2.088v.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Smiling(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.316C6.099 1.316 1.316 6.099 1.316 12S6.099 22.684 12 22.684S22.684 17.901 22.684 12c-.012-5.896-4.788-10.672-10.683-10.684zm0 22.297C5.586 23.613.387 18.414.387 12S5.586.387 12 .387S23.613 5.586 23.613 12v.015c0 6.405-5.192 11.597-11.597 11.597z"/><path fill="currentColor" d="M12 24C5.386 23.966.034 18.614 0 12.003V12C0 5.373 5.373 0 12 0s12 5.373 12 12c-.034 6.614-5.386 11.966-11.997 12zM12 .774C5.8.774.774 5.8.774 12S5.8 23.226 12 23.226S23.226 18.2 23.226 12C23.222 5.802 18.198.779 12.001.774zm0 22.297C5.886 23.071.929 18.114.929 12S5.886.929 12 .929S23.071 5.886 23.071 12S18.114 23.071 12 23.071m0-21.368C6.313 1.703 1.703 6.313 1.703 12S6.313 22.297 12 22.297S22.297 17.687 22.297 12v-.005c0-5.684-4.608-10.292-10.292-10.292z"/><path fill="currentColor" d="M12 19.665h-.003a7.842 7.842 0 0 1-7.413-5.287l-.016-.055a.44.44 0 0 1 .306-.541l.003-.001h.31l.232.232c.917 2.764 3.479 4.723 6.499 4.723l.086-.001H12h.011a6.933 6.933 0 0 0 6.555-4.674l.015-.049a.44.44 0 0 1 .545-.309l-.003-.001a.353.353 0 0 1 .232.23l.001.002c.037.041.059.095.059.155s-.022.114-.059.155c-.983 3.163-3.884 5.419-7.312 5.419h-.045h.002z"/><path fill="currentColor" d="M12 20.052h-.005a8.23 8.23 0 0 1-7.797-5.594l-.017-.058a.837.837 0 0 1 .536-1.083l.006-.002a.58.58 0 0 1 .62.079l-.001-.001a.574.574 0 0 1 .387.461v.003c.868 2.628 3.301 4.49 6.169 4.49h.026h-.001h.076a6.415 6.415 0 0 0 6.105-4.445l.013-.045a.839.839 0 0 1 1.09-.541l-.006-.002a.795.795 0 0 1 .462.383l.002.004a.858.858 0 0 1 .076.626l.001-.006c-1.024 3.341-4.082 5.729-7.697 5.729h-.047H12zm0-.775h.068a7.321 7.321 0 0 0 6.962-5.058l.015-.052l.077-.465l-.155.387h-.077c-.992 2.947-3.729 5.032-6.954 5.032h-.014h.001h-.013a7.324 7.324 0 0 1-6.939-4.981l-.015-.052h-.077s-.077.077 0 .077c1.024 2.995 3.814 5.11 7.098 5.11h.026h-.001zm6.348-8.825h-3.716a.542.542 0 0 1 0-1.084h3.716a.547.547 0 0 1 .387.929c-.077.155-.232.155-.387.155"/><path fill="currentColor" d="M18.348 10.839h-3.716l-.044.001a.843.843 0 0 1-.652-.309l-.001-.002a1.147 1.147 0 0 1-.309-.693v-.004l-.001-.044c0-.263.121-.497.309-.652l.001-.001a1.89 1.89 0 0 1 .684-.307l.013-.002h3.716l.044-.001c.263 0 .497.121.652.309l.001.002c.173.184.287.425.309.693v.004a.843.843 0 0 1-.308.696l-.002.001a.84.84 0 0 1-.653.311l-.046-.001zm-3.716-1.162c-.077 0-.077 0-.155.077s-.077.077-.077.155s0 .077.077.155s.077.077.155.077h3.716c.077 0 .077 0 .155-.077s.077-.077.077-.155s0-.077-.077-.155a.264.264 0 0 0-.153-.077h-.001zm-5.264.775H5.643a.61.61 0 0 1-.611-.611v-.022a.53.53 0 0 1 .53-.53h.013h-.001h3.729a.53.53 0 0 1 .53.53v.013v-.001c.155.387-.155.619-.465.619z"/><path fill="currentColor" d="M9.368 10.839H5.652a1.024 1.024 0 0 1-1.006-1.005v-.001l-.001-.044c0-.263.121-.497.309-.652l.002-.001c.184-.173.425-.287.693-.309h3.72l.044-.001c.263 0 .497.121.652.309l.001.002c.173.184.287.425.309.693v.035a.976.976 0 0 1-.976.976l-.032-.001h.002zM5.652 9.677c-.077 0-.077 0-.155.077s-.077.077-.077.155a.246.246 0 0 0 .232.232h3.717a.246.246 0 0 0 .232-.232v-.001c0-.077 0-.077-.077-.155s-.077-.077-.155-.077z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Snapchat(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.015 0l.145-.002c1.344 0 2.595.4 3.64 1.087l-.025-.015c1.808 1.104 2.997 3.066 2.997 5.306a6 6 0 0 1-.022.521l.001-.022q0 .734-.141 2.984a.965.965 0 0 0 .437.11h.001c.295-.024.566-.099.814-.217l-.014.006c.234-.112.505-.188.792-.211l.008-.001c.328.004.631.109.88.285l-.005-.003a.84.84 0 0 1 .422.72a1.006 1.006 0 0 1-.487.841l-.005.003a3.23 3.23 0 0 1-1.056.487l-.023.005a3.818 3.818 0 0 0-1.095.462l.016-.009a.876.876 0 0 0-.492.74v.002a1.9 1.9 0 0 0 .192.683l-.005-.011c.417.905.954 1.68 1.602 2.346l-.002-.002a6.767 6.767 0 0 0 2.214 1.562l.043.017a6.65 6.65 0 0 0 1.205.352l.044.007c.249.033.44.244.44.5l-.002.049v-.002q0 1.094-3.422 1.61a1.637 1.637 0 0 0-.172.603v.006c-.032.27-.11.516-.224.739l.006-.012a.54.54 0 0 1-.479.29l-.039-.001h.002a7.126 7.126 0 0 1-1.012-.108l.043.006a6.767 6.767 0 0 0-.989-.101h-.082a5.08 5.08 0 0 0-.929.085l.031-.005c-.37.062-.7.185-.998.359l.014-.007a7.98 7.98 0 0 0-.926.608l.02-.014q-.422.32-.906.633a4.577 4.577 0 0 1-1.154.515l-.033.008c-.447.134-.96.211-1.492.211h-.058h.003h-.044a5.16 5.16 0 0 1-1.5-.221l.037.01a4.53 4.53 0 0 1-1.19-.534l.018.011q-.48-.313-.898-.633a7.843 7.843 0 0 0-.867-.573l-.039-.021a2.864 2.864 0 0 0-.958-.349l-.017-.002a5.176 5.176 0 0 0-.92-.08h-.068h.003a6.097 6.097 0 0 0-1.063.124l.04-.007c-.266.058-.579.1-.9.116l-.015.001a.566.566 0 0 1-.545-.286l-.001-.003a2.138 2.138 0 0 1-.217-.732l-.001-.01a1.707 1.707 0 0 0-.176-.635l.004.01Q0 19.674 0 18.581a.504.504 0 0 1 .435-.547h.003a7.04 7.04 0 0 0 1.297-.376l-.047.017A6.674 6.674 0 0 0 3.942 16.1l.003-.004a8.14 8.14 0 0 0 1.579-2.293l.021-.051c.099-.196.165-.424.187-.665l.001-.007a.875.875 0 0 0-.487-.74l-.005-.002a3.902 3.902 0 0 0-1.058-.455l-.028-.006a3.471 3.471 0 0 1-1.097-.5l.012.008a.979.979 0 0 1-.492-.819v-.001a.85.85 0 0 1 .403-.709l.004-.002a1.45 1.45 0 0 1 .859-.289h.001c.277.029.531.101.764.209l-.015-.006c.243.111.524.184.821.203h.027c.175 0 .34-.041.486-.113l-.006.003q-.141-2.218-.141-2.969a7.013 7.013 0 0 1 .437-2.861l-.016.048A6.668 6.668 0 0 1 8.86 1.002l.029-.017A8.387 8.387 0 0 1 12.862 0l.162.002h-.008z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Snorkel(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.856 11.37v-.564a6.044 6.044 0 0 0-6.04-6.04H6.04A6.044 6.044 0 0 0 0 10.806v.564a6.044 6.044 0 0 0 6.04 6.04h10.776a6.044 6.044 0 0 0 6.04-6.04M6.498 16.068a4.98 4.98 0 0 1 0-9.96h9.861a4.98 4.98 0 0 1 0 9.96a4.6 4.6 0 0 1-2.998-1.684l-.007-.009a2.145 2.145 0 0 0-1.561-.671c-.633 0-1.201.273-1.595.708l-.002.002a4.939 4.939 0 0 1-3.694 1.655h-.006z"/><path fill="currentColor" d="M26.113 0H24.23a.604.604 0 0 0-.602.602v3.2c0 .313.241.57.547.595h.002v11.546a4.443 4.443 0 0 1-4.438 4.438h-5.872a2.628 2.628 0 0 0-2.427-1.63a2.622 2.622 0 0 0 0 5.244h.007a2.612 2.612 0 0 0 2.414-1.615l.006-.017h5.872a6.435 6.435 0 0 0 6.429-6.428V4.398a.597.597 0 0 0 .542-.594V.596a.598.598 0 0 0-.596-.598z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Snow(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.651 23.256a.744.744 0 1 1 1.488 0a.744.744 0 0 1-1.488 0m2.13-2.129a.744.744 0 1 1 1.488 0a.744.744 0 0 1-1.488 0m6.009-.971a.744.744 0 1 1 1.488 0a.744.744 0 0 1-1.488 0m-11.447 0a.744.744 0 1 1 1.488 0a.744.744 0 0 1-1.488 0m7.567-1.158a.744.744 0 1 1 1.488 0a.744.744 0 0 1-1.488 0m6.009-.971a.744.744 0 1 1 1.488 0a.744.744 0 0 1-1.488 0m-11.447 0a.744.744 0 1 1 1.486 0a.744.744 0 0 1-1.485 0zm.44-2.112h-.065a5.912 5.912 0 0 1 0-11.824h.069h-.003c.091 0 .183.007.274.012C7.576 1.639 10.176.002 13.158.002a7.975 7.975 0 0 1 7.975 7.975v.015a4.027 4.027 0 0 1 3.063 3.904a4.019 4.019 0 0 1-4.019 4.019zm-4.273-5.914a4.278 4.278 0 0 0 4.273 4.274h14.261a2.382 2.382 0 1 0-1.541-4.199l.003-.003a.822.822 0 0 1-1.065-1.251l.001-.001a4.005 4.005 0 0 1 1.886-.887l.024-.004A6.333 6.333 0 0 0 7.916 4.417l-.014.022a5.9 5.9 0 0 1 1.711.954l-.012-.009a.82.82 0 1 1-1.026 1.279l.001.001a4.204 4.204 0 0 0-2.654-.936H5.91h.001a4.278 4.278 0 0 0-4.276 4.274z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Snowflake(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m21.188 17.029l-2.17-1.259l2.047-1.183c.44-.258-.486-1.851-.923-1.593l-2.973 1.714l-4.673-2.699l4.673-2.699l2.973 1.711c.44.258 1.366-1.349.923-1.593l-2.048-1.183l2.17-1.259c.44-.258-.486-1.851-.923-1.593l-2.17 1.259V4.285c0-.516-1.851-.516-1.851 0v3.444l-4.673 2.699V5.031l2.973-1.714c.44-.258-.486-1.851-.923-1.593l-2.036 1.167V.388c0-.516-1.851-.516-1.851 0v2.503L7.685 1.707c-.44-.258-1.366 1.349-.923 1.593l2.973 1.714V10.4L5.062 7.701V4.286c0-.516-1.851-.516-1.851 0v2.366L1.037 5.393C.597 5.135-.329 6.742.114 6.986l2.17 1.259L.236 9.428c-.44.258.486 1.851.923 1.593l2.973-1.714l4.673 2.716l-4.673 2.699l-2.973-1.714c-.44-.258-1.366 1.349-.923 1.593l2.048 1.183l-2.17 1.262c-.44.258.486 1.851.923 1.593l2.17-1.259v2.367c0 .516 1.851.516 1.851 0V16.3l4.673-2.699v5.386L6.76 20.703c-.44.258.486 1.851.923 1.593l2.048-1.183v2.502c0 .516 1.851.516 1.851 0v-2.503l2.048 1.183c.44.258 1.366-1.349.923-1.593l-2.973-1.714v-5.37l4.673 2.699v3.444c0 .516 1.851.516 1.851 0v-2.365l2.17 1.259c.444.227 1.354-1.366.914-1.623z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SnowflakeEight(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.662 14.947l-1.531.879l-1.394-.803l1.621-.939c.379-.213-.409-1.591-.788-1.364l-2.409 1.394l-1.939-1.121v-1.985l1.939-1.121l2.409 1.394c.379.213 1.167-1.151.788-1.364l-1.62-.939l1.394-.803l1.531.879c.273.167.863-.848.576-1l-.954-.546l.893-.515c.439-.257-.485-1.848-.924-1.591l-.894.515V4.811c0-.318-1.167-.318-1.167 0v1.772l-1.394.803V5.507c0-.439-1.575-.439-1.575 0v2.788L13.3 9.401l-1.727-.97V6.189l2.408-1.393c.379-.213-.409-1.591-.788-1.364l-1.619.937V2.764l1.53-.882c.273-.167-.303-1.167-.576-1l-.954.546V.386c0-.515-1.848-.515-1.848 0v1.03L8.772.87c-.273-.167-.863.848-.576 1l1.529.879v1.606l-1.622-.939c-.379-.213-1.167 1.151-.788 1.364l2.409 1.394v2.243l-1.727.969l-1.924-1.121V5.507c0-.439-1.575-.439-1.575 0v1.879l-1.394-.803V4.811c0-.318-1.167-.318-1.167 0v1.106l-.894-.515C.604 5.145-.321 6.752.119 6.993l.894.515l-.939.546c-.273.167.303 1.167.576 1l1.546-.879l1.394.803l-1.622.939c-.379.213.409 1.591.788 1.364l2.409-1.394l1.939 1.121v1.985l-1.939 1.121l-2.411-1.393c-.379-.213-1.167 1.151-.788 1.364l1.621.939l-1.394.803l-1.528-.879c-.273-.167-.863.848-.576 1l.954.546l-.894.515c-.439.257.485 1.848.924 1.591l.895-.514v1.106c0 .318 1.167.318 1.167 0v-1.774l1.396-.803v1.879c0 .439 1.575.439 1.575 0v-2.788L8.03 14.6l1.727.97v2.243l-2.409 1.394c-.379.213.409 1.591.788 1.364l1.621-.939v1.606l-1.546.894c-.273.167.303 1.167.576 1l.954-.546v1.03c0 .515 1.848.515 1.848 0v-1.03l.954.546c.273.167.863-.848.576-1l-1.531-.879v-1.606l1.621.939c.379.213 1.167-1.151.788-1.364l-2.409-1.394v-2.242l1.727-.97l1.924 1.121v2.758c0 .439 1.575.439 1.575 0v-1.879l1.394.803v1.773c0 .318 1.167.318 1.167 0v-1.105l.894.515c.439.257 1.364-1.35.924-1.591l-.894-.515l.954-.546c.259-.168-.318-1.168-.591-1.001zm-10-.591l-2.045-1.167v-2.378l2.045-1.151l2.045 1.151v2.379z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SnowflakeFive(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.752 17.796v2.5c0 .515 1.848.515 1.848 0v-3.56zm.379-3.606L1.04 12.402c-.439-.258-1.364 1.348-.924 1.591l2.17 1.243zM17.949 3.72c0-.515-1.848-.515-1.848 0v3.56l1.848-1.06zm-.378 6.09l3.091 1.788c.439.258 1.364-1.348.924-1.591l-2.17-1.243zM6.3 3.31l3.091 1.788V2.962l-2.17-1.243C6.783 1.446 5.859 3.052 6.3 3.31m8.183-1.605l-2.17 1.243v2.136l3.091-1.788c.439-.242-.487-1.847-.921-1.591m.921 19l-3.091-1.788v2.136l2.17 1.243c.434.258 1.36-1.348.921-1.591m-8.182 1.591l2.17-1.243v-2.136L6.3 20.705c-.441.242.482 1.848.922 1.591m2.712-7.607v8.924c0 .515 1.848.515 1.848 0v-8.924l-.924.515zm-1.879-2.136L.329 17.008c-.439.258.485 1.848.924 1.591l7.71-4.455l-.909-.515zm5.591-1.106l7.713-4.455c.439-.258-.485-1.848-.924-1.591l-7.697 4.455l.909.515zm7.727 5.561l-7.712-4.455v1.076l-.924.516L20.45 18.6c.439.258 1.364-1.334.924-1.591zM8.964 9.857L1.252 5.402C.813 5.144-.112 6.75.328 6.993l7.713 4.455v-1.076zm2.803-.547V.386c0-.515-1.848-.515-1.848 0V9.31l.924-.515zM1.04 11.599l3.091-1.788l-1.849-1.06l-2.17 1.243c-.435.258.488 1.848.928 1.606zM5.6 3.72c0-.515-1.848-.515-1.848 0v2.5L5.6 7.28zm15.062 8.697l-3.091 1.788l1.848 1.06l2.17-1.243c.436-.27-.488-1.861-.927-1.606zm-4.546 7.879c0 .515 1.848.515 1.848 0v-2.5l-1.848-1.06z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SnowflakeFour(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.18 11.042a.923.923 0 1 0 .932-1.594l-.005-.002l-2.052-1.185l2.173-1.262c.441-.258-.487-1.854-.927-1.596l-2.173 1.262V4.293c0-.517-1.854-.517-1.854 0v3.42l-2.948 1.702l-1.733-.973v-3.42l2.979-1.717c.441-.258-.487-1.854-.927-1.596l-2.052 1.185V.386c0-.517-1.854-.517-1.854 0v2.508L7.687 1.709c-.441-.258-1.366 1.353-.927 1.596l2.979 1.717v3.42l-1.733.973l-2.934-1.702V4.262c0-.517-1.854-.517-1.854 0v2.372L1.045 5.372C.604 5.114-.321 6.725.118 6.968L2.291 8.23L.239 9.415c-.441.258.487 1.854.927 1.596l2.979-1.717l2.948 1.702v1.991l-2.948 1.717l-2.979-1.717C.725 12.729-.2 14.34.239 14.583l2.052 1.185L.118 17.03c-.441.258.487 1.854.927 1.596l2.173-1.262v2.372c0 .517 1.854.517 1.854 0v-3.451l2.948-1.702l1.733.973v3.42l-2.979 1.717c-.441.258.487 1.854.927 1.596l2.052-1.185v2.511c0 .517 1.854.517 1.854 0v-2.508l2.052 1.185c.441.258 1.366-1.353.927-1.596l-2.979-1.717v-3.423l1.733-.973l2.948 1.702v3.451c0 .517 1.854.517 1.854 0v-2.372l2.173 1.262c.441.258 1.366-1.353.927-1.596l-2.173-1.262l2.052-1.185c.441-.258-.487-1.854-.927-1.596L17.2 14.704l-2.948-1.702v-1.991l2.964-1.702zm-7.463 2.157l-.927.517l-.198.106l-.927.517l-.927-.517l-.198-.106l-.927-.517v-2.385l.927-.517l.198-.106l.927-.517l.927.517l.198.106l.927.517v1.306z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SnowflakeOne(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m20.188 16.662l.8-.461a.194.194 0 0 0 .077-.076v-.001c.092-.153 0-.415-.108-.615c-.141-.246-.385-.495-.584-.369l-1.431.816l-.954-.554l.8-.461a.194.194 0 0 0 .077-.076v-.001a.375.375 0 0 0 .031-.152v-.002a1.186 1.186 0 0 0-.474-.813l-.003-.002a.232.232 0 0 0-.263-.015l.001-.001l-1.431.816l-1.123-.646l.566-.323a.342.342 0 0 0 .076-.076l.001-.001a.29.29 0 0 0 .03-.123a.91.91 0 0 0-.351-.614l-.002-.002a.212.212 0 0 0-.12-.036a.227.227 0 0 0-.094.021l.001-.001l-1.046.6l-2.569-1.477l2.568-1.476l1.046.6a.218.218 0 0 0 .213-.016h-.001a.877.877 0 0 0 .353-.612v-.01a.236.236 0 0 0-.031-.118l.001.001a.194.194 0 0 0-.076-.077h-.001l-.566-.323l1.123-.646l1.431.816a.231.231 0 0 0 .263-.016c.258-.191.435-.479.476-.81l.001-.006v-.002a.376.376 0 0 0-.032-.154l.001.002a.342.342 0 0 0-.076-.076l-.001-.001l-.8-.461l.954-.554l1.431.816c.2.123.446-.123.584-.369a.668.668 0 0 0 .106-.62l.001.005a.194.194 0 0 0-.076-.077h-.001l-.8-.461l1-.584c.062-.046.062-.338-.2-.784c-.246-.446-.508-.6-.584-.554l-1 .584v-.939a.224.224 0 0 0-.032-.109l.001.001a.707.707 0 0 0-.587-.212h.003c-.277 0-.615.077-.615.323v1.646l-.954.554v-.95a.17.17 0 0 0-.031-.097a.297.297 0 0 0-.121-.107l-.002-.001a1.1 1.1 0 0 0-.945.003l.007-.003a.329.329 0 0 0-.153.21v1.648l-1.123.646v-.661l.001-.014a.114.114 0 0 0-.031-.078a.264.264 0 0 0-.091-.091l-.001-.001a.936.936 0 0 0-.714.002l.006-.002a.207.207 0 0 0-.123.184v1.203l-2.569 1.477V7.695l1.046-.6a.222.222 0 0 0 .108-.19v-.011a.91.91 0 0 0-.351-.614l-.002-.002a.212.212 0 0 0-.12-.036a.227.227 0 0 0-.094.021l.001-.001l-.566.323V5.293l1.431-.816a.258.258 0 0 0 .123-.219v-.012v.001a1.186 1.186 0 0 0-.474-.813l-.003-.002a.232.232 0 0 0-.263-.015l.001-.001l-.8.461V2.785l1.431-.816c.2-.123.108-.446-.031-.693s-.385-.495-.584-.369l-.8.461V.21c0-.077-.262-.212-.769-.212s-.769.141-.769.212v1.154l-.8-.461c-.2-.123-.446.123-.584.369s-.231.566-.031.693l1.431.816v1.092l-.8-.461a.231.231 0 0 0-.263.016a1.189 1.189 0 0 0-.476.81l-.001.006c0 .153.077.212.123.231l1.431.816v1.292l-.566-.323a.218.218 0 0 0-.213.016h.001a.877.877 0 0 0-.353.612v.004l-.001.015a.21.21 0 0 0 .107.184l.001.001l1.046.6v2.971L7.326 9.186V7.983a.207.207 0 0 0-.122-.184l-.001-.001a.936.936 0 0 0-.714.002l.006-.002a.207.207 0 0 0-.123.184v.662l-1.123-.646V6.352a.231.231 0 0 0-.152-.212l-.002-.001a1.1 1.1 0 0 0-.945.003l.007-.003c-.123.077-.153.169-.141.212v.939l-.954-.554V5.093c0-.231-.338-.323-.615-.323s-.615.077-.615.323v.939l-1-.584c-.062-.046-.323.123-.584.554c-.246.446-.262.739-.2.784l1 .584l-.8.461c-.2.123-.108.446.031.693s.385.495.584.369l1.431-.816l.954.554l-.8.461a.258.258 0 0 0-.123.219v.012v-.001c.042.336.218.625.474.813l.003.002a.232.232 0 0 0 .263.015l-.001.001l1.431-.816l1.123.646l-.566.323a.222.222 0 0 0-.108.19v.011a.91.91 0 0 0 .351.614l.002.002a.212.212 0 0 0 .12.036a.227.227 0 0 0 .094-.021l-.001.001l1.046-.6l2.569 1.477l-2.569 1.477l-1.046-.6a.218.218 0 0 0-.213.016h.001a.877.877 0 0 0-.353.612v.004l-.001.015a.21.21 0 0 0 .107.184l.001.001l.566.323l-1.123.646l-1.431-.816a.231.231 0 0 0-.263.016a1.189 1.189 0 0 0-.476.81l-.001.006c0 .153.077.212.123.231l.8.461l-.954.554l-1.431-.816c-.2-.123-.446.123-.584.369s-.231.566-.031.693l.8.461l-1 .584c-.062.046-.062.338.2.784c.246.446.508.6.584.554l1-.584v.939c0 .231.338.323.615.323s.615-.077.615-.323v-1.646l.954-.554v.939a.24.24 0 0 0 .14.212l.002.001a1.1 1.1 0 0 0 .945-.003l-.007.003a.329.329 0 0 0 .153-.21v-1.648l1.123-.646v.661a.207.207 0 0 0 .122.184l.001.001a.936.936 0 0 0 .714-.002l-.006.002a.207.207 0 0 0 .123-.184v-1.203l2.554-1.476v2.971l-1.046.6a.222.222 0 0 0-.108.19v.011a.91.91 0 0 0 .351.614l.002.002c.031.02.069.032.11.032h.014h-.001h.007a.134.134 0 0 0 .085-.031l.566-.323v1.292l-1.458.799a.258.258 0 0 0-.123.219v.012v-.001c.042.336.218.625.474.813l.003.002a.339.339 0 0 0 .153.046h.007a.179.179 0 0 0 .102-.032h-.001l.8-.461v1.092l-1.431.816c-.2.123-.108.446.03.693c.123.2.292.4.477.4a.282.282 0 0 0 .109-.032l-.002.001l.8-.461v1.154c0 .077.262.212.769.212s.769-.141.769-.212v-1.154l.8.461a.267.267 0 0 0 .106.031h.001c.169 0 .354-.2.477-.4c.141-.246.231-.566.03-.693l-1.431-.816v-1.092l.8.461c.031.019.068.03.107.031a.34.34 0 0 0 .155-.047l-.002.001c.258-.191.435-.479.476-.81l.001-.006c0-.153-.077-.212-.123-.231l-1.431-.816v-1.292l.566.323a.13.13 0 0 0 .084.031h.008a.457.457 0 0 0 .126-.032l-.003.001a.877.877 0 0 0 .353-.612v-.004l.001-.015a.21.21 0 0 0-.107-.184l-.001-.001l-1.046-.6v-2.971l2.569 1.477v1.203a.207.207 0 0 0 .122.184l.001.001a.936.936 0 0 0 .714-.002l-.006.002a.264.264 0 0 0 .091-.091l.001-.001a.319.319 0 0 0 .031-.09v-.663l1.123.646v1.646c.002.098.065.18.152.212l.002.001a1.1 1.1 0 0 0 .945-.003l-.007.003a.45.45 0 0 0 .122-.107l.001-.001a.163.163 0 0 0 .031-.097v-.95l.954.554v1.646c0 .231.338.323.615.323a.707.707 0 0 0 .584-.212a.283.283 0 0 0 .03-.107v-.94l1 .584c.062.046.323-.123.584-.554c.246-.446.262-.739.2-.784z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SnowflakeSeven(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.805 17.884v2.535c0 .522 1.874.522 1.874 0v-3.611zm.384-3.656l-3.134-1.813c-.445-.261-1.383 1.367-.937 1.614l2.201 1.26zM18.217 6.13V3.595c0-.522-1.874-.522-1.874 0v3.611zm-.399 3.657l3.134 1.813c.445.261 1.383-1.367.937-1.614l-2.201-1.26zM6.402 3.196l3.134 1.813V2.843l-2.201-1.26c-.458-.277-1.394 1.352-.934 1.613zm6.084 1.797L15.62 3.18c.445-.261-.491-1.874-.937-1.614l-2.201 1.26zm3.135 15.842l-3.134-1.813v2.166l2.201 1.26c.44.261 1.379-1.367.934-1.613zm-6.1-1.814l-3.134 1.813c-.445.261.491 1.874.937 1.614l2.201-1.26zM1.07 11.6l3.134-1.813L2.33 8.712L.129 9.972c-.442.261.48 1.874.94 1.628zm4.609-7.989c0-.522-1.874-.522-1.874 0v2.535l1.874 1.075zm15.273 8.819l-3.134 1.813l1.874 1.075l2.201-1.26c.442-.274-.495-1.887-.94-1.628zm-4.609 7.99c0 .522 1.874.522 1.874 0v-2.535l-1.874-1.075z"/><path fill="currentColor" d="m20.476 16.594l.86-.491c.246-.139-.261-1.014-.507-.876l-1.367.783l-5.623-3.258v-1.506l5.623-3.242l1.367.783c.246.139.753-.737.507-.876l-.86-.491l1.106-.646c.062-.046.062-.338-.2-.783c-.246-.445-.507-.599-.587-.553l-1.106.646v-.999c0-.277-1.014-.277-1.014 0v1.582l-5.608 3.242l-1.306-.737V2.673l1.367-.783c.246-.139-.261-1.014-.507-.876l-.86.491V.215c0-.073-.261-.215-.768-.215s-.768.139-.768.215v1.29l-.86-.491c-.246-.139-.753.737-.507.876l1.367.783v6.499l-1.306.737l-5.608-3.242V5.085c0-.277-1.014-.277-1.014 0v.999l-1.106-.646c-.062-.046-.323.123-.587.553c-.246.445-.261.737-.2.783l1.106.646l-.86.491c-.246.139.261 1.014.507.876l1.367-.783l5.623 3.242v1.506l-5.623 3.242l-1.367-.783c-.246-.139-.753.737-.507.876l.86.491l-1.106.646c-.062.046-.062.338.2.783c.246.445.507.599.587.553l1.106-.646v.998c0 .277 1.014.277 1.014 0V17.33l5.608-3.242l1.306.737v6.499l-1.367.783c-.246.139.261 1.014.507.876l.86-.491v1.29c0 .073.261.215.768.215s.768-.139.768-.215v-1.29l.86.491c.246.139.753-.737.507-.876l-1.367-.783v-6.499l1.306-.737l5.608 3.242v1.582c0 .277 1.014.277 1.014 0v-.999l1.106.646c.062.046.323-.123.587-.553c.246-.445.261-.737.2-.783zm-9.463-2.535L9.23 13.045v-2.09l1.783-1.014l1.783 1.014v2.09z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SnowflakeSix(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m6.752 3.311l2.695 1.545h2.394l2.695-1.56c.44-.257-.485-1.848-.922-1.591L11.57 2.887v-2.5c0-.515-1.848-.515-1.848 0v2.5L7.677 1.705c-.455-.258-1.379 1.348-.924 1.606zm13.364 7.71c.44.257 1.364-1.348.922-1.591l-2.045-1.182l2.167-1.257c.44-.257-.485-1.848-.922-1.591l-2.167 1.257V4.279c0-.515-1.848-.515-1.848 0V7.4l1.197 2.076zM3.206 4.28v2.364L1.037 5.386C.597 5.129-.327 6.734.115 6.977l2.167 1.257L.237 9.416c-.44.257.485 1.848.922 1.591l2.695-1.56l1.197-2.076v-3.09c.004-.516-1.845-.516-1.845 0zm-2.045 8.697c-.44-.257-1.364 1.348-.922 1.591l2.045 1.182l-2.169 1.258c-.44.257.485 1.848.922 1.591l2.167-1.257v2.364c0 .515 1.848.515 1.848 0v-3.121l-1.197-2.076zm20.001 4.031l-2.167-1.257l2.045-1.182c.44-.257-.485-1.848-.922-1.591l-2.695 1.56l-1.201 2.075v3.121c0 .515 1.848.515 1.848 0v-2.362l2.167 1.257c.44.227 1.364-1.364.925-1.621m-6.622 3.696l-2.695-1.56H9.447l-2.695 1.56c-.44.257.485 1.848.922 1.591l2.045-1.182v2.499c0 .515 1.848.515 1.848 0v-2.5l2.045 1.182c.444.26 1.366-1.347.927-1.589zm.757-12.454l-2.015 1.167l-1.727-.969V6.084h-1.85v2.348l-1.727.969l-2.015-1.167l-.922 1.591l2.015 1.167v1.986l-2.032 1.166l.922 1.591l2.015-1.167l1.727.969v2.348h1.848v-2.348l1.727-.969l2.015 1.167l.922-1.591l-2.015-1.167v-1.986l2.031-1.167zm-4.652 6.182l-2.106-1.197v-2.47l2.106-1.197l2.11 1.197v2.469z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SnowflakeThree(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m16.232 9.832l-.923-1.592l-3.729 2.153V6.087H9.728v4.306L5.981 8.241l-.923 1.592l3.745 2.168l-3.745 2.153l.923 1.607l3.745-2.168v4.321h1.849v-4.321l3.729 2.168l.923-1.607l-3.729-2.153zM6.757 3.3l2.697 1.561h2.395L14.546 3.3c.44-.258-.485-1.849-.923-1.592L11.576 2.89V.388c0-.515-1.849-.515-1.849 0V2.89L7.68 1.708c-.453-.275-1.378 1.331-.923 1.592m14.418 13.703l-2.168-1.258l2.047-1.182c.44-.258-.485-1.849-.923-1.592l-2.697 1.561l-1.198 2.077v3.123c0 .515 1.849.515 1.849 0v-2.365l2.168 1.258c.436.227 1.36-1.365.921-1.622zM17.43 9.468l2.697 1.561c.44.258 1.365-1.348.923-1.592l-2.047-1.182l2.168-1.258c.44-.258-.485-1.849-.923-1.592L18.08 6.663V4.298c0-.515-1.849-.515-1.849 0v3.123zM.131 6.982L2.299 8.24L.252 9.422c-.44.258.485 1.849.923 1.592l2.697-1.561L5.07 7.376V4.253c0-.515-1.849-.515-1.849 0v2.365L1.053 5.36C.602 5.132-.322 6.724.132 6.982zm14.418 13.72l-2.697-1.561H9.453l-2.697 1.561c-.44.258.485 1.849.923 1.592l2.047-1.182v2.5c0 .515 1.849.515 1.849 0V21.11l2.047 1.182c.444.26 1.367-1.332.928-1.59zm-10.688-6.17l-2.697-1.561c-.44-.258-1.365 1.348-.923 1.592l2.047 1.182L.12 17.003c-.44.258.485 1.849.923 1.592l2.168-1.258v2.365c0 .515 1.849.515 1.849 0v-3.123z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SnowflakeTwo(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m22.122 14.326l-2.379 1.381l-1.228-.712l1.591-.921c.379-.212-.409-1.561-.788-1.348l-2.379 1.381l-3.586-2.097l3.607-2.091l2.379 1.381c.379.212 1.151-1.137.788-1.348l-1.591-.921l1.228-.712L22.143 9.7c.379.212 1.151-1.137.788-1.348l-1.591-.921l.712-.409c.44-.258-.485-1.849-.921-1.591l-.712.409V3.991c0-.424-1.561-.424-1.561 0v2.743l-1.228.712V5.597c0-.424-1.561-.424-1.561 0v2.729l-3.607 2.091V6.251l2.379-1.381c.379-.212-.409-1.561-.788-1.348l-1.591.921V3.019l2.379-1.381c.379-.212-.409-1.561-.788-1.348l-1.591.921V.378c0-.515-1.849-.515-1.849 0v.833L8.991.29C8.612.078 7.84 1.427 8.203 1.638l2.379 1.381v1.409l-1.591-.921c-.379-.212-1.151 1.137-.788 1.348l2.379 1.381v4.163L6.991 8.327V5.585c0-.424-1.561-.424-1.561 0v1.833l-1.228-.712V3.963c0-.424-1.561-.424-1.561 0v1.849l-.712-.409c-.44-.258-1.364 1.348-.921 1.591l.712.409l-1.596.924c-.379.212.409 1.561.788 1.348l2.379-1.381l1.228.712l-1.622.941c-.379.212.409 1.561.788 1.348l2.379-1.381l3.607 2.091l-3.607 2.072l-2.379-1.364c-.379-.212-1.151 1.137-.788 1.348l1.591.921l-1.228.712l-2.379-1.381C.502 14.101-.27 15.45.093 15.661l1.591.921l-.712.409c-.44.258.485 1.849.921 1.591l.712-.409v1.849c0 .424 1.561.424 1.561 0v-2.743l1.228-.712v1.849c0 .424 1.561.424 1.561 0v-2.742l3.607-2.091v4.183l-2.379 1.381c-.379.212.409 1.561.788 1.348l1.591-.921v1.409l-2.379 1.381c-.379.212.409 1.561.788 1.348l1.591-.921v.833c0 .515 1.849.515 1.849 0v-.833l1.591.921c.379.212 1.151-1.137.788-1.348l-2.379-1.381v-1.409l1.591.921c.379.212 1.151-1.137.788-1.348l-2.379-1.381V13.6l3.607 2.091v2.742c0 .424 1.561.424 1.561 0v-1.849l1.228.712v2.743c0 .424 1.561.424 1.561 0V18.19l.712.409c.44.258 1.364-1.348.921-1.591l-.712-.409l1.591-.921c.394-.215-.374-1.565-.757-1.353z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Snows(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.18 23.256a.744.744 0 1 1 1.488 0a.744.744 0 0 1-1.488 0m2.13-2.129a.743.743 0 1 1 0 .001zm6.009-.971a.744.744 0 1 1 1.488 0a.744.744 0 0 1-1.488 0m-11.447 0a.744.744 0 1 1 1.488 0a.744.744 0 0 1-1.488 0m7.567-1.158a.744.744 0 1 1 1.488 0a.744.744 0 0 1-1.488 0m6.01-.971a.743.743 0 1 1 0 .001zm-11.448 0a.744.744 0 1 1 1.488 0a.744.744 0 0 1-1.488 0m9.177-2.113H5.912a5.913 5.913 0 0 1 0-11.826c.091 0 .183.007.274.012c1.389-2.464 3.989-4.101 6.971-4.101A7.978 7.978 0 0 1 20.9 6.058l.011.055a5.631 5.631 0 0 1 9.411 3.976v.009a2.984 2.984 0 0 1 2.09 2.842a2.978 2.978 0 0 1-2.978 2.978h-.015h.001zm1.2-7.861a4.02 4.02 0 0 1 2.821 3.831c0 .9-.296 1.731-.797 2.401l.008-.011h6.012a1.338 1.338 0 1 0-.866-2.359l.002-.002a.82.82 0 0 1-1.058-1.252l.001-.001c.329-.28.726-.491 1.163-.605l.021-.005a4.001 4.001 0 0 0-3.991-3.761a3.997 3.997 0 0 0-3.305 1.747l-.009.014zM1.642 9.998a4.278 4.278 0 0 0 4.273 4.273h14.262a2.381 2.381 0 1 0-1.541-4.199l.003-.003a.82.82 0 1 1-1.061-1.25l.001-.001a3.985 3.985 0 0 1 1.887-.887l.024-.004A6.333 6.333 0 0 0 7.918 4.416l-.014.022a5.9 5.9 0 0 1 1.711.954l-.012-.009a.82.82 0 1 1-1.026 1.279l.001.001a4.204 4.204 0 0 0-2.654-.936h-.012h.001a4.278 4.278 0 0 0-4.276 4.274z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Soundcloud(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m11.408 18.708l.209-3.138l-.209-6.81a.37.37 0 0 0-.097-.221a.28.28 0 0 0-.207-.091h-.01a.285.285 0 0 0-.209.091a.298.298 0 0 0-.085.21v.011v-.001l-.182 6.81l.182 3.138a.34.34 0 0 0 .097.215c.05.053.122.086.201.086h.001a.286.286 0 0 0 .315-.285l-.001-.018v.001zm3.855-.378l.143-2.747l-.156-7.63v-.016a.342.342 0 0 0-.168-.295l-.002-.001a.36.36 0 0 0-.417.001l.001-.001a.345.345 0 0 0-.17.296v.017v-.001l-.013.078l-.131 7.54l.143 3.073v.021c0 .081.03.156.078.213a.37.37 0 0 0 .292.143h.009a.355.355 0 0 0 .26-.12a.308.308 0 0 0 .12-.246v-.016v.001zM1.656 13.903l.264 1.667l-.264 1.64q-.026.12-.12.12t-.12-.12L1.2 15.57l.221-1.667q.026-.12.12-.12t.115.12m1.12-1.028l.338 2.695l-.338 2.64a.127.127 0 0 1-.127.12h-.004q-.12 0-.12-.13l-.3-2.63l.3-2.695a.103.103 0 0 1 .101-.121l.019.002h-.001h.004c.068 0 .123.053.127.12zm1.197-.496l.325 3.19l-.325 3.086l.001.017c0 .07-.057.127-.127.127l-.017-.001h.001h-.007a.15.15 0 0 1-.149-.143l-.274-3.086l.274-3.19q.026-.156.156-.156t.143.157zm1.224-.091l.3 3.281l-.3 3.178a.174.174 0 0 1-.174.169h-.009l-.021.001a.15.15 0 0 1-.15-.15l.002-.021v.001l-.274-3.178l.274-3.281l-.001-.021a.15.15 0 0 1 .15-.15l.021.002h-.001h.008c.094 0 .171.075.174.169zm1.224.24l.274 3.047l-.274 3.203q-.026.209-.209.209h-.001a.186.186 0 0 1-.136-.059a.204.204 0 0 1-.059-.143v-.007l-.257-3.208l.26-3.047a.202.202 0 0 1 .195-.196l.014-.001c.107 0 .194.087.194.194v.002zm1.224-1.906l.274 4.948l-.274 3.203a.237.237 0 0 1-.065.163a.205.205 0 0 1-.156.072q-.209 0-.24-.24l-.24-3.203l.24-4.948q.026-.24.24-.24c.062 0 .118.028.155.072c.04.043.065.101.065.164zm1.224-1.119L9.12 15.6l-.247 3.178a.255.255 0 0 1-.508.001v-.001L8.153 15.6l.209-6.094l-.001-.017a.254.254 0 1 1 .508 0l-.001.018v-.001zm1.276-.521l.24 6.588l-.24 3.151v.012a.263.263 0 0 1-.263.263l-.025-.001h.001q-.247 0-.274-.274l-.209-3.151l.209-6.588V8.98c0-.078.033-.149.085-.199a.258.258 0 0 1 .188-.086h.005c.076 0 .144.033.191.086c.05.053.083.123.091.2zm2.527-.053l.196 6.641l-.196 3.112a.333.333 0 0 1-.325.325h-.005a.301.301 0 0 1-.216-.091a.362.362 0 0 1-.104-.239v-.001l-.182-3.112l.182-6.641v-.016c0-.088.037-.168.097-.224a.325.325 0 0 1 .228-.091a.326.326 0 0 1 .326.316v.017v-.001zm1.289.247l.182 6.407l-.182 3.073v.001a.35.35 0 0 1-.6.246a.39.39 0 0 1-.12-.246v-.002l-.156-3.073l.156-6.407a.401.401 0 0 1 .12-.26a.338.338 0 0 1 .244-.104h.008c.093 0 .177.04.235.104a.41.41 0 0 1 .11.259v.001zm2.76 6.407l-.182 3.008a.404.404 0 0 1-.69.286a.44.44 0 0 1-.131-.286v-.001l-.078-1.484l-.078-1.524l.156-8.28v-.04a.49.49 0 0 1 .156-.312a.41.41 0 0 1 .259-.091h.001c.074 0 .142.024.196.065l-.001-.001c.119.07.2.194.209.337v.001zm14.479-.247a3.68 3.68 0 0 1-3.679 3.672h-10.24a.504.504 0 0 1-.287-.143a.39.39 0 0 1-.12-.282V6.872q0-.3.36-.43a6.384 6.384 0 0 1 2.365-.445c1.7 0 3.247.652 4.406 1.719L24 7.712a6.28 6.28 0 0 1 2.082 4.188l.001.022a3.69 3.69 0 0 1 5.117 3.403v.009z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sourcetree(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 0h-.079c-1.66 0-3.239.349-4.667.978l.074-.029A12.269 12.269 0 0 0 3.52 3.523A12.219 12.219 0 0 0 .978 7.251l-.031.079A11.398 11.398 0 0 0 0 11.919v.086v-.004v.079c0 1.66.349 3.239.978 4.667l-.029-.074a12.276 12.276 0 0 0 2.572 3.807a12.224 12.224 0 0 0 3.729 2.542l.079.031c1.354.6 2.933.949 4.593.949h.083h-.004h.079c1.66 0 3.239-.349 4.667-.978l-.074.029a12.276 12.276 0 0 0 3.809-2.573a12.219 12.219 0 0 0 2.542-3.728l.031-.079c.6-1.354.949-2.932.949-4.593v-.158c0-1.66-.349-3.239-.978-4.667l.029.074a12.286 12.286 0 0 0-2.573-3.806A12.219 12.219 0 0 0 16.754.981L16.675.95C15.321.35 13.741 0 12.08 0h-.083zm.64 22.79v-2.087l5.193-2.633a.421.421 0 0 0 .154-.129l.001-.001a.335.335 0 0 0 .059-.191v-.011v.001v-1.186c.225-.082.412-.226.543-.412l.002-.004c.133-.179.214-.404.214-.648v-.118a1.146 1.146 0 0 0-.366-.746l-.001-.001a1.09 1.09 0 0 0-.75-.297h-.022h.001h-.02c-.308 0-.587.127-.786.332c-.205.2-.332.478-.332.787v.021v-.001v.059c.002.03.011.057.024.081v-.001c.013.222.096.423.227.583l-.001-.002c.133.163.304.289.501.364l.008.003v.949l-4.649 2.372v-1.942l2.039-.949a.533.533 0 0 0 .166-.13l.001-.001a.3.3 0 0 0 .071-.194v-3.233l3.793-2.134a.318.318 0 0 0 .142-.141l.001-.002a.407.407 0 0 0 .047-.189v-.901c.223-.079.409-.218.543-.397l.002-.003c.133-.177.214-.401.214-.644v-.04c0-.308-.127-.587-.332-.786a1.066 1.066 0 0 0-.775-.332H18.4a1.082 1.082 0 0 0-.71.366l-.001.001a1.113 1.113 0 0 0-.285.746v.025v-.001v.02c0 .243.08.466.216.646l-.002-.003c.137.182.322.321.538.397l.008.003v.688l-3.818 2.134a.51.51 0 0 0-.129.129l-.001.002a.338.338 0 0 0-.055.184v.017v-.001v3.2l-1.52.711v-6.972l3.2-1.566a.379.379 0 0 0 .153-.141l.001-.002a.347.347 0 0 0 .059-.189V7.04c.223-.079.409-.218.543-.397l.002-.003c.133-.177.214-.401.214-.644v-.021v.001v-.061a.078.078 0 0 0-.024-.057a1.127 1.127 0 0 0-.366-.719l-.001-.001a1.06 1.06 0 0 0-.738-.297h-.128a1.108 1.108 0 0 0-.719.378l-.001.001a1.103 1.103 0 0 0-.297.754v.123c.032.227.12.428.251.596l-.002-.003c.127.167.301.291.502.354l.007.002v.972l-2.656 1.304V5.288c.225-.082.412-.226.543-.412l.002-.004c.133-.179.214-.404.214-.648V4.2c0-.308-.127-.587-.332-.786a1.095 1.095 0 0 0-.787-.332h-.021h.001h-.02c-.308 0-.587.127-.786.332c-.205.2-.332.478-.332.787v.021v-.001v.118c.03.223.119.421.25.583l-.002-.002c.133.163.304.289.501.364l.008.003v6.569l-1.874-.996V7.865a.292.292 0 0 0-.048-.16l.001.001a.87.87 0 0 0-.095-.119l-1.306-.998a.825.825 0 0 0 .07-.196l.001-.006c.015-.067.024-.143.024-.222v-.024c0-.308-.127-.587-.332-.786a1.097 1.097 0 0 0-.786-.331H7.52h.001h-.02c-.308 0-.587.127-.786.332c-.205.2-.332.478-.332.787v.021v-.001v.024c0 .305.125.581.326.78c.2.205.478.332.787.332h.021h-.001h.015a.858.858 0 0 0 .288-.049l-.006.002c.11-.041.2-.081.287-.125l-.015.007l1.162.925v3.035a.29.29 0 0 0 .06.178l-.001-.001a.42.42 0 0 0 .152.129l.002.001l2.419 1.28V15.1l-4.055-1.874l.071-1.47v-.024a.29.29 0 0 0-.06-.178l.001.001a.42.42 0 0 0-.152-.129l-.002-.001l-1.851-.97a.522.522 0 0 0 .024-.157v-.2a1.127 1.127 0 0 0-.366-.719l-.001-.001a1.093 1.093 0 0 0-.752-.299h-.018h.001h-.007c-.305 0-.58.127-.775.332c-.208.2-.338.481-.338.792v.015v-.001v.118c.029.285.164.534.366.71l.001.001c.193.177.451.285.735.285h.131a1.45 1.45 0 0 0 .33-.083l-.01.003c.104-.04.195-.092.275-.156l-.002.002l1.707.88l-.047 1.47v.008c0 .075.017.145.048.208l-.001-.003a.305.305 0 0 0 .164.142l.002.001l4.577 2.134v6.869q-.308 0-.605-.024l-.605-.047l.071-4.364a.349.349 0 0 0-.06-.191l.001.001a.364.364 0 0 0-.175-.142l-.002-.001l-2.87-1.328v-.125c0-.305-.127-.58-.332-.775a1.095 1.095 0 0 0-.787-.332H6.49h.001h-.114a1.082 1.082 0 0 0-.71.366l-.001.001a1.113 1.113 0 0 0-.285.746v.025v-.001v.007c0 .305.127.58.332.775c.195.205.47.332.775.332h.039c.156 0 .305-.033.439-.094l-.007.003c.14-.067.261-.147.369-.242l-.002.002l2.656 1.21v4.008a10.611 10.611 0 0 1-3.534-1.343l.048.027a10.85 10.85 0 0 1-2.773-2.354l-.014-.017a11.1 11.1 0 0 1-1.824-3.112l-.026-.076a10.267 10.267 0 0 1-.676-3.699v-.111c0-1.494.314-2.915.88-4.201l-.026.067a11.06 11.06 0 0 1 2.324-3.44A11.019 11.019 0 0 1 7.73 2.065l.071-.028a10.269 10.269 0 0 1 4.139-.856h.061h-.003h.064c1.494 0 2.915.314 4.201.88l-.067-.026a11.078 11.078 0 0 1 3.44 2.32a11 11 0 0 1 2.296 3.369l.028.071c.54 1.218.854 2.639.854 4.134v.067v-.003v.064c0 1.444-.292 2.82-.82 4.072l.026-.069a11.1 11.1 0 0 1-2.175 3.373l.005-.006a10.852 10.852 0 0 1-3.172 2.32l-.065.028c-1.16.568-2.516.932-3.948 1.009l-.026.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spinner(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.818 6.664h-.001a1.847 1.847 0 1 1 1.306-.541a1.773 1.773 0 0 1-1.277.541h-.029zm-2.97 7.182h-.001a1.847 1.847 0 1 1 1.306-.541a1.773 1.773 0 0 1-1.278.541h-.031h.002zM12 3.692a1.847 1.847 0 1 1 1.306-.541a1.773 1.773 0 0 1-1.277.541zM4.818 21.029h-.001a1.847 1.847 0 1 1 1.306-.541a1.769 1.769 0 0 1-1.276.541h-.031zM19.182 7.125a2.308 2.308 0 1 1 0-4.615a2.308 2.308 0 0 1 0 4.615M12 24a1.847 1.847 0 1 1 1.306-.541a1.773 1.773 0 0 1-1.277.541zm10.154-9.231h-.048c-.75 0-1.428-.309-1.913-.807l-.001-.001c-.499-.503-.808-1.196-.808-1.961s.308-1.458.808-1.962a2.663 2.663 0 0 1 1.914-.808h.05h-.003h.048c.75 0 1.427.309 1.913.807l.001.001c.499.503.808 1.196.808 1.961s-.308 1.458-.808 1.962a2.664 2.664 0 0 1-1.915.809h-.049zm-2.971 7.643h-.05a3.097 3.097 0 0 1-2.236-.951l-.001-.001c-.584-.584-.945-1.391-.945-2.283s.361-1.698.945-2.283a3.106 3.106 0 0 1 2.234-.945h.054h-.003h.042c.877 0 1.67.362 2.237.944l.001.001c.588.582.952 1.39.952 2.283s-.364 1.7-.952 2.282a3.102 3.102 0 0 1-2.24.953h-.04z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpinnerCog(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.5 0h3l.75 4.5l1.28.53l3.97-2.56l2.03 2.03l-2.56 3.97l.53 1.28l4.499.75v3l-4.5.75l-.53 1.28l2.56 3.97l-2.03 2.03l-3.97-2.56l-1.28.53l-.75 4.499h-3l-.75-4.5l-1.28-.53l-3.97 2.56l-2.031-2.03l2.56-3.97l-.53-1.28l-4.499-.75v-3l4.5-.75l.53-1.28l-2.56-3.97l2.03-2.031l3.97 2.56l1.28-.53zM12 7.5a4.5 4.5 0 1 0 0 9a4.5 4.5 0 0 0 0-9"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpinnerFidget(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.87.12a4.238 4.238 0 0 0-2.226 1.454l-.007.009a4.071 4.071 0 0 0-.863 2.515c0 1.101.434 2.101 1.141 2.837l-.001-.001c.306.238.512.591.553.993v.006a7.743 7.743 0 0 1-.007 2.677l.007-.047a6.43 6.43 0 0 1-4.048 4.717l-.044.014l-.632.24l-.48-.133a3.916 3.916 0 0 0-1.134-.164c-.667 0-1.295.163-1.848.451l.022-.011a4.118 4.118 0 0 0-2.3 3.69a4.11 4.11 0 0 0 2.205 3.643l.023.011a4.114 4.114 0 0 0 5.856-2.62l.006-.029c.12-.494.222-.626.897-1.131c.62-.45 1.34-.813 2.115-1.052l.052-.014a6.53 6.53 0 0 1 1.72-.225c1.734 0 3.314.659 4.503 1.74l-.005-.005a.975.975 0 0 1 .372.676v.004a3.658 3.658 0 0 0 1.081 1.846l.003.002c.284.318.626.577 1.01.762l.019.008a3.588 3.588 0 0 0 2.055.409l-.015.001c.404 0 .793-.066 1.157-.188l-.026.007a4.197 4.197 0 0 0 2.691-2.951l.006-.029a5.1 5.1 0 0 0-.012-1.94l.005.032a4.224 4.224 0 0 0-2.757-2.899l-.03-.008a4.59 4.59 0 0 0-2.271-.042l.031-.006a1.494 1.494 0 0 1-1.307-.1l.007.004c-2.477-.935-4.206-3.286-4.206-6.041c0-.435.043-.86.125-1.271l-.007.041l.108-.566l.4-.427a3.993 3.993 0 0 0 1.192-2.676v-.008a4.103 4.103 0 0 0-3.06-4.201l-.029-.006a4.559 4.559 0 0 0-2.081.007l.031-.006zm1.76 2.546a1.613 1.613 0 0 1 .41 2.583a1.325 1.325 0 0 1-1.151.439l.006.001a1.325 1.325 0 0 1-1.142-.437l-.001-.001a1.517 1.517 0 0 1-.449-1.436l-.002.01a1.625 1.625 0 0 1 2.343-1.155l-.009-.004zm.337 9.39c.514.251.92.658 1.162 1.159l.006.015a1.716 1.716 0 0 1 .168 1.014l.001-.009a1.813 1.813 0 0 1-.197 1.081l.005-.01a2.466 2.466 0 0 1-1.803 1.364l-.015.002a2.48 2.48 0 0 1-2.612-1.369l-.006-.015a1.746 1.746 0 0 1-.185-1.063l-.001.009a1.74 1.74 0 0 1 .165-1.003l-.005.01a2.54 2.54 0 0 1 1.451-1.319l.018-.005a1.984 1.984 0 0 1 .889-.101l-.009-.001l.067-.001c.333 0 .646.089.915.244l-.009-.005zm-9.394 5.698c.443.137.8.441 1.004.837l.004.009a1.127 1.127 0 0 1 .173.819l.001-.007a1.58 1.58 0 0 1-.994 1.423l-.011.004a1.786 1.786 0 0 1-1.419-.101l.01.005a1.604 1.604 0 0 1-.806-1.554l-.001.007a1.45 1.45 0 0 1 .457-.98l.001-.001a1.54 1.54 0 0 1 1.591-.459l-.011-.003zm17.696.06c.398.175.718.467.922.833l.005.01c.076.195.12.421.12.657c0 .27-.057.526-.161.757l.005-.012a1.56 1.56 0 0 1-1.405.885l-.042-.001h.002a1.56 1.56 0 0 1-1.612-1.527v-.002a1.123 1.123 0 0 1 .183-.822l-.003.004a1.558 1.558 0 0 1 1.999-.779l-.01-.004zM5.062 6.082a10.763 10.763 0 0 0-2.195 3.944l-.02.076c-.217.789-.234 1.006-.03.4c.319-.864.67-1.598 1.077-2.295l-.037.068c.399-.627.823-1.172 1.294-1.673l-.006.006a3.02 3.02 0 0 0 .432-.513l.007-.011a.406.406 0 0 0-.237-.246l-.003-.001c-.11.066-.203.147-.28.243zm15.68.416a10.149 10.149 0 0 1 2.014 4.023l.015.07c.205.88.24.96.386.96c.234 0 .259-.193.09-.765a10.599 10.599 0 0 0-2.466-4.325l.005.005l-.4-.422zm-15.293.795a9.066 9.066 0 0 0-1.378 2.596l-.019.064c-.259.789-.325 1.168-.096.572c.426-1.086 1-2.02 1.708-2.835l-.01.012c.301-.349.349-.463.222-.59s-.21-.085-.427.18z"/><path fill="currentColor" d="M20.273 7.642a9.716 9.716 0 0 1 1.455 3.25l.014.067c.036.169.16.253.301.198c.15-.06.16-.193.012-.608a9.337 9.337 0 0 0-1.961-3.213l.005.005c-.102-.114-.024.018.174.301m-14.62.956a7.921 7.921 0 0 0-.775 1.869l-.013.057a3.14 3.14 0 0 0 .237-.382l.009-.018c.212-.368.424-.682.654-.979l-.014.019c.295-.4.379-.548.337-.62c-.11-.163-.296-.139-.434.054zm14.349.332c.146.285.293.631.416.988l.018.059c.193.584.331.72.512.494c.108-.133.108-.145-.08-.475a11.87 11.87 0 0 0-.978-1.413l.018.023c-.158-.193-.152-.162.094.325zm-8.516 12.79a.16.16 0 0 0 .042.222h.001c.09.09.198.096.981.06c.48-.018.969-.054 1.083-.08c.198-.036.205-.042.06-.084a2.704 2.704 0 0 0-.478-.048h-.002a8.885 8.885 0 0 1-.873-.096c-.703-.113-.735-.113-.814.026m-1.066.644c-.102.042-.102.234 0 .313a6.857 6.857 0 0 0 1.858.251c.176 0 .351-.007.524-.019l-.023.001a8.262 8.262 0 0 0 2.001-.265l-.057.013h-.045c-.224 0-.444.013-.661.039l.026-.003a9.86 9.86 0 0 1-2.755-.204l.065.012c-.686-.15-.837-.178-.933-.136zm-1.017.897c-.072.181.06.295.463.4c.803.214 1.726.337 2.677.337c.751 0 1.485-.077 2.193-.223l-.07.012c.62-.12.602-.174-.03-.08a10.992 10.992 0 0 1-4.34-.356l.078.019c-.741-.218-.923-.242-.971-.11z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpinnerRefresh(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m7.5 21l2.999-3v1.5a7.501 7.501 0 0 0 5.299-12.811l2.114-2.124A10.465 10.465 0 0 1 21 12.002C21 17.8 16.3 22.5 10.502 22.5H10.5V24zM0 12C.007 6.204 4.704 1.507 10.499 1.5h.001V0l3 3l-3 3V4.5h-.002a7.502 7.502 0 0 0-5.299 12.812l-2.112 2.124a10.397 10.397 0 0 1-3.088-7.407v-.03v.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SpinnerRotateForward(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23.314 8.518V.686l-2.84 2.84A11.962 11.962 0 0 0 11.981.004c-6.627 0-12 5.373-12 12s5.373 12 12 12c4.424 0 8.289-2.394 10.37-5.958l.031-.057l-2.662-1.536c-1.57 2.695-4.447 4.478-7.739 4.478a8.927 8.927 0 1 1 6.32-15.232l-2.82 2.82z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Spotify(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.61 16.906v-.007a.905.905 0 0 0-.464-.791l-.005-.002a13.236 13.236 0 0 0-6.999-1.797h.015a20.964 20.964 0 0 0-4.626.559l.142-.028a.752.752 0 0 0-.656.816v-.003v.012c0 .205.081.391.212.528c.132.14.318.227.525.227l.031-.001h-.001c.225-.034.42-.077.611-.133l-.032.008a18.536 18.536 0 0 1 3.786-.422h.01a11.736 11.736 0 0 1 6.258 1.639l-.054-.03c.146.097.323.159.513.172h.004a.734.734 0 0 0 .734-.734v-.017v.001zm1.5-3.36v-.032c0-.395-.219-.74-.542-.918l-.005-.003c-2.358-1.387-5.195-2.207-8.223-2.207c-.118 0-.237.001-.354.004h.018a17.46 17.46 0 0 0-4.858.687l.123-.031a.939.939 0 0 0-.749 1.004v-.004v.004c0 .516.418.934.934.934h.004c.218-.028.414-.072.603-.131l-.024.007a14.593 14.593 0 0 1 3.894-.516h.027h-.001a14.596 14.596 0 0 1 7.695 1.975l-.07-.038c.169.108.371.181.588.203h.01a.934.934 0 0 0 .934-.934v-.004zm1.69-3.874l.001-.052c0-.449-.251-.839-.62-1.039l-.006-.003a15.808 15.808 0 0 0-4.472-1.707l-.106-.019c-1.596-.372-3.429-.586-5.312-.586h-.05h.003a20.216 20.216 0 0 0-5.83.769L4.55 7a1.244 1.244 0 0 0-.6.398l-.002.002a1.135 1.135 0 0 0-.24.761v-.003v.029c0 .305.122.582.321.784c.196.203.471.328.775.328h.032h-.002c.23-.02.443-.064.645-.131l-.02.006a17.521 17.521 0 0 1 4.814-.578h-.014h.044c1.699 0 3.352.194 4.939.56l-.147-.029a13.65 13.65 0 0 1 4.028 1.519l-.067-.035c.178.111.393.18.623.187h.012c.304 0 .579-.122.778-.32c.205-.194.333-.469.333-.773l-.001-.035zM24 12v.09c0 2.187-.598 4.234-1.64 5.987l.03-.054a12.045 12.045 0 0 1-4.311 4.337l-.056.03c-1.729 1.012-3.806 1.609-6.024 1.609s-4.295-.597-6.081-1.64l.057.031a12.045 12.045 0 0 1-4.337-4.311l-.03-.056C.596 16.294-.001 14.217-.001 11.999s.597-4.295 1.64-6.081l-.031.057a12.045 12.045 0 0 1 4.311-4.337l.056-.03C7.704.596 9.781-.001 11.999-.001s4.295.597 6.081 1.64l-.057-.031a12.045 12.045 0 0 1 4.337 4.311l.03.056A11.606 11.606 0 0 1 24 11.908v.096v-.005z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StackOverflow(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.12 21.857H2.143v-6.428H0V24h19.259v-8.571H17.12zM4.504 14.839l.442-2.102l10.486 2.21l-.442 2.09zM5.883 9.83l.898-1.955l9.71 4.54l-.898 1.942zm2.692-4.768l1.366-1.647l8.218 6.87l-1.366 1.647zM13.888 0l6.388 8.585l-1.716 1.286l-6.386-8.585zM4.272 19.701v-2.13h10.714v2.13z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Star(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m19.555 23.411l-6.664-3.285a1.261 1.261 0 0 0-1.202.045l.006-.003l-6.416 3.75a.61.61 0 0 1-.902-.626v.003l.994-7.542a1.248 1.248 0 0 0-.392-1.082l-.001-.001l-4.571-4.247a1.265 1.265 0 0 1 .648-2.17l.007-.001l5.987-1.108c.421-.078.765-.355.935-.727l.003-.008L10.478.746a1.272 1.272 0 0 1 2.271-.087l.003.007l2.881 5.471c.197.365.558.62.981.666h.006l6.045.681a1.265 1.265 0 0 1 .811 2.119l.001-.001l-4.27 4.562a1.251 1.251 0 0 0-.315 1.116l-.001-.008l1.52 7.453a.61.61 0 0 1-.86.683l.004.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StarHalf(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.855.029a1.305 1.305 0 0 0-1.454.721l-.003.008l-2.511 5.669a.953.953 0 0 1-.805.728h-.005L1.084 8.289a1.66 1.66 0 0 0-.73.406l.001-.001a1.201 1.201 0 0 0 .079 1.78l.002.002l4.535 4.212a1.243 1.243 0 0 1 .405 1.058v-.005l-.972 7.532a.72.72 0 0 0 .083.408l-.002-.004a.58.58 0 0 0 .812.242l-.003.001l6.398-3.726c.081 0 .162-.081.243-.081V.027z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Steam(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24.72 7.094a3.812 3.812 0 0 1-6.508 2.695a3.812 3.812 0 1 1 5.387-5.391l.002.002a3.664 3.664 0 0 1 1.12 2.641v.055v-.003zM12.687 18.687v-.012a3.894 3.894 0 0 0-3.894-3.894H8.78h.001c-.299 0-.59.034-.87.099l.026-.005l1.625.656a3.042 3.042 0 0 1 1.704 1.644l.007.02c.164.356.26.772.26 1.21c0 .418-.087.816-.244 1.176l.007-.019a3.01 3.01 0 0 1-1.652 1.696l-.02.007c-.355.161-.77.254-1.206.254c-.422 0-.824-.088-1.188-.246l.019.007q-.328-.125-.969-.383l-.953-.383c.337.627.82 1.138 1.405 1.498l.017.01a3.749 3.749 0 0 0 1.999.571h.034h-.002h.012a3.894 3.894 0 0 0 3.894-3.894zM25.656 7.11a4.775 4.775 0 0 0-4.765-4.766a4.766 4.766 0 1 0 3.367 8.141a4.57 4.57 0 0 0 1.399-3.296l-.001-.083zm2.344 0a7.11 7.11 0 0 1-7.109 7.109h-.001L14.062 19.2a5.122 5.122 0 0 1-1.698 3.402l-.005.004A5.101 5.101 0 0 1 8.847 24h-.069h.004a5.313 5.313 0 0 1-5.181-4.152l-.007-.035L0 18.376v-6.703l6.08 2.453a5.052 5.052 0 0 1 2.664-.75h.041h-.002q.203 0 .547.031l4.438-6.359A7.137 7.137 0 0 1 20.89.001h.001a7.12 7.12 0 0 1 7.11 7.109v.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StepBackwrad(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.108 23.855L2.4 13.361v10.04a.599.599 0 0 1-.599.599H.597a.599.599 0 0 1-.599-.599V.599C-.002.268.266 0 .597 0h1.202c.331 0 .599.268.599.599v10.039L21.106.145a1.142 1.142 0 0 1 1.691 1.001v.024v-.001v21.68c0 .634-.511 1.149-1.143 1.155h-.001a1.108 1.108 0 0 1-.552-.152l.005.003z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StepForward(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 22.835V1.146A1.141 1.141 0 0 1 1.696.149L1.69.146L20.4 10.639V.601a.6.6 0 0 1 .599-.6h1.2a.6.6 0 0 1 .6.6v22.8A.599.599 0 0 1 22.2 24h-1.201a.599.599 0 0 1-.599-.599V13.362L1.69 23.856a1.1 1.1 0 0 1-.548.145A1.155 1.155 0 0 1 0 22.846zv.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stethoscope(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.267 14.312v-.007a2.48 2.48 0 1 0-2.854 2.452l.014.002v1.726a4.355 4.355 0 0 1-8.71 0v-2.462h.03a.753.753 0 0 0 .743-.629l.001-.004c3.374-.64 5.891-3.56 5.9-7.071V.454a.45.45 0 0 0-.453-.45h-2.307a.454.454 0 0 0-.453.453v.618c0 .25.203.452.453.453h1.252v6.781a5.695 5.695 0 0 1-4.389 5.536l-.039.008a.757.757 0 0 0-.706-.489H6.614a.763.763 0 0 0-.704.483l-.002.005c-2.542-.607-4.403-2.857-4.407-5.542V1.528h1.258c.25 0 .453-.203.453-.453V.457a.453.453 0 0 0-.453-.453H.453a.452.452 0 0 0-.452.452v7.859a7.21 7.21 0 0 0 5.825 7.065l.045.007c.056.345.34.609.69.633h.002v2.535a5.514 5.514 0 0 0 11.028 0v-.07v.004v-1.826a2.496 2.496 0 0 0 1.679-2.35zm-3.816 0a1.339 1.339 0 1 1 0 .002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stop(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.334 24H2.668a2.668 2.668 0 0 1-2.667-2.667V2.666A2.669 2.669 0 0 1 2.668 0h18.667A2.668 2.668 0 0 1 24 2.666v18.667A2.668 2.668 0 0 1 21.333 24z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stopwatch(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m18.241 6.523l1.45-1.657l-1.657-1.45l-1.45 1.553a10.446 10.446 0 0 0-4.811-1.961l-.055-.006V2.07h2.588V-.001H6.955V2.07h2.588v1.035C4.174 3.595 0 8.075 0 13.531c0 5.781 4.686 10.467 10.467 10.467s10.467-4.686 10.467-10.467c0-2.7-1.022-5.162-2.701-7.018zm-7.662 14.806a7.972 7.972 0 1 1 0-15.944a8.131 8.131 0 0 1 5.695 2.382a7.95 7.95 0 0 1 2.381 5.683v.012v-.001c-.173 4.347-3.711 7.812-8.07 7.869z"/><path fill="currentColor" d="M9.544 7.248h2.174v7.972H9.544z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Strikethrough(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M26.068 10.759H10.965a4.138 4.138 0 0 1 0-8.276h10.759a1.242 1.242 0 0 0 .002-2.482H10.965a6.62 6.62 0 0 0-5.156 10.771l-.01-.013H1.206a1.242 1.242 0 1 0 0 2.484l.037-.001h-.002h15.108a4.138 4.138 0 0 1 0 8.276H5.585A1.242 1.242 0 0 0 5.583 24h10.77A6.615 6.615 0 0 0 21.5 13.229l.01.013h4.597a1.242 1.242 0 1 0-.037-2.483h.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func StuckOutTongue(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.316C6.099 1.316 1.316 6.099 1.316 12S6.099 22.684 12 22.684S22.684 17.901 22.684 12c-.012-5.896-4.788-10.672-10.683-10.684zm0 22.297C5.586 23.613.387 18.414.387 12S5.586.387 12 .387S23.613 5.586 23.613 12v.015c0 6.405-5.192 11.597-11.597 11.597z"/><path fill="currentColor" d="M12 24C5.386 23.966.034 18.614 0 12.003V12C0 5.373 5.373 0 12 0s12 5.373 12 12c-.034 6.614-5.386 11.966-11.997 12zM12 .774C5.8.774.774 5.8.774 12S5.8 23.226 12 23.226S23.226 18.2 23.226 12C23.222 5.802 18.198.779 12.001.774zm0 22.297C5.886 23.071.929 18.114.929 12S5.886.929 12 .929S23.071 5.886 23.071 12S18.114 23.071 12 23.071m0-21.368C6.313 1.703 1.703 6.313 1.703 12S6.313 22.297 12 22.297S22.297 17.687 22.297 12v-.005c0-5.684-4.608-10.292-10.292-10.292z"/><path fill="currentColor" d="M9.677 9.91a2.09 2.09 0 1 1-2.09-2.09h.009c1.15 0 2.082.932 2.082 2.082v.009zm8.284.232h-1.703l.697-.542a.426.426 0 0 0 .232-.379v-.008a.57.57 0 0 0-.155-.465a.426.426 0 0 0-.379-.232h-.008a.57.57 0 0 0-.465.155l-2.013 1.548c-.09.112-.168.239-.228.376l-.004.011v.155c.03.306.285.542.596.542h.025h-.001h3.419a.53.53 0 0 0 .53-.53v-.013v.001a.516.516 0 0 0-.505-.621l-.039.001h.002zm-.464 4.335a.23.23 0 0 0-.155-.059a.23.23 0 0 0-.155.059a.353.353 0 0 0-.232.23l-.001.002c-.697 1.394-2.71 2.323-4.877 2.323c-2.245 0-4.258-.929-4.877-2.323a.469.469 0 1 0-.851.39l-.001-.003a5.48 5.48 0 0 0 3.525 2.548l.036.007c.077 2.09 1.006 3.716 2.168 3.716s2.09-1.626 2.168-3.716a5.49 5.49 0 0 0 3.548-2.53l.014-.025c0-.232-.077-.542-.31-.619"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stumbleupon(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.386 8.988V7.056a1.67 1.67 0 0 0-3.34 0v10.019l.001.094a6.6 6.6 0 0 1-2.06 4.799l-.003.003A7.031 7.031 0 0 1 0 16.98v-4.357h5.37v4.323c0 .453.188.862.491 1.153h.001c.303.299.719.483 1.179.483s.876-.184 1.179-.483c.303-.292.492-.701.492-1.154v-.035v.002V6.685c0-1.86.796-3.534 2.067-4.699l.005-.004A7.103 7.103 0 0 1 15.716.001c1.923 0 3.668.761 4.95 1.999l-.002-.002a6.443 6.443 0 0 1 2.063 4.732l-.001.086v-.004v2.23l-3.192.95zm8.677 3.634h5.37v4.354a7.033 7.033 0 0 1-11.985 4.984a6.648 6.648 0 0 1-2.063-4.821l.001-.104v.005v-4.387l2.145.998l3.192-.95v4.445c0 .45.188.857.491 1.145l.001.001c.303.299.719.483 1.179.483s.876-.184 1.179-.483c.303-.289.491-.695.491-1.145v-.026v.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Stylus(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.698 16.267c.622-.745.696-1.514.214-2.942c-.305-.903-.812-1.6-.439-2.16c.4-.598 1.239-.018.537.781l.14.098c.842.098 1.258-1.056.629-1.386c-1.66-.866-3.113.8-2.472 2.728c.274.818.659 1.685.348 2.374a1.543 1.543 0 0 1-1.125.95l-.01.002c-.726.037-.244-1.63.592-2.045c.074-.037.177-.086.08-.207a1.828 1.828 0 0 0-1.978 1.015l-.005.011c-1.014 1.934 1.922 2.649 3.49.781zm18.127-5.685c.24.586.598 1.166.385 1.68c-.177.439-.409.622-.666.666c-.36.061-.262-1.068.354-1.404c.055-.03.134-.177.061-.262a1.4 1.4 0 0 0-1.454.815l-.004.009c-.69 1.446 1.562 1.84 2.667.421c.439-.567.458-1.129.037-2.148c-.269-.647-.678-1.129-.421-1.556c.274-.452.934-.061.439.543l.11.061c.64.037.903-.818.415-1.031a1.572 1.572 0 0 0-1.915 2.216l-.004-.009z"/><path fill="currentColor" d="M13.671 9.035c-.446-.354-1.697.24-2.05 1.12a13.872 13.872 0 0 1-1.781 3.504l.029-.044c-.683.75-.75.171-.683-.262a13.126 13.126 0 0 1 1.728-4.099l-.031.052c-.202-.299-1.52-.256-2.435 1.166c-.342.537-1.12 2.326-1.99 3.735c-.189.305-.427.091-.244-.622a33.677 33.677 0 0 1 1.689-5.094l-.084.224c1.755-.387 3.795-.639 5.883-.7l.049-.001c.226-.061.378-.262 0-.274h-.153c-1.939 0-3.846.138-5.712.404l.214-.025a4.243 4.243 0 0 1 1.232-1.66l.008-.006a1.531 1.531 0 0 0-1.937.61l-.004.007a10.2 10.2 0 0 0-.676 1.176l-.027.062c-1.291.17-2.445.428-3.558.778l.147-.04c-.671.256-.598 1.068-.189.915a24.693 24.693 0 0 1 3.089-.899l.181-.034a24.847 24.847 0 0 0-1.567 4.696l-.033.174c-.378 2.136.946 2.124 1.593 1.28c.702-.922 2.166-4.16 2.392-4.504c.067-.116.16-.055.11.049c-1.636 3.265-1.495 4.529-.171 4.248a4.26 4.26 0 0 0 1.888-1.666l.011-.019c.055-.128.171-.116.146-.061c-1.04 2.691-2.356 4.87-3.241 5.554c-.806.617-1.404-.72 1.446-2.64c.421-.287.226-.678-.25-.543c-1.471.232-5.68 1.569-7.531 2.85c-.14.098-.269.177-.262.378c.006.116.207.074.305.012a19.182 19.182 0 0 1 6.478-2.447l.12-.018a.133.133 0 0 0 .057.013a.152.152 0 0 0 .042-.006h-.001c.104-.025.098.03.03.074a3.323 3.323 0 0 1-.318.167l-.024.01c-1.514.592-2.429 1.898-2.106 2.56c.274.574 1.76.366 2.459-.012c1.721-.934 2.972-2.765 3.826-5.291c.747-2.235 1.687-4.777 1.907-4.847zm10.001 5.17a41.424 41.424 0 0 0-11.785.882l.274-.051c-.794.207-.574.629-.171.55c.006 0 .177-.042.183-.042c2.191-.427 7.506-.8 10.607-.207c.374.066 1.491-1.051.891-1.13zm-9.219-.33c.781-.39 1.941-2.807 2.704-4.132c.055-.098.153-.018.098.049c-1.929 3.32-1.11 3.705-.348 3.656c1.019-.061 1.959-1.526 2.166-1.855c.086-.128.134-.025.086.067c-.049.153-.226.421-.39.787c-.232.518.012.72.214.812c.32.153 1.184.055 1.318-.48c-.866-.018 1.208-4.107 1.422-4.358a1.473 1.473 0 0 0-1.883.828l-.004.01c-.873 1.727-1.605 3.12-2.063 3.143c-.891.049 1.026-3.851 1.337-3.973c-.189-.274-1.404-.16-2.08.891c-.24.378-1.734 3.015-2.099 3.448c-.647.769-.696.11-.513-.659c.098-.394.2-.717.32-1.033l-.021.062a4.574 4.574 0 0 1 1.16-1.576l.005-.004c1.886-2.094 2.966-3.79 2.539-4.455c-.378-.592-1.642-.33-2.454.891c-1.495 2.24-2.874 5.31-3.051 6.714s.847 1.511 1.537 1.167m.793-4.107c.067-.153.11-.195.226-.452a21.354 21.354 0 0 1 2.139-3.828l-.045.068c.36-.378.866.134-.049 1.538a14.258 14.258 0 0 1-1.814 2.266l.002-.002v.006c-.171.189-.32.348-.39.439c-.05.062-.104.05-.067-.035z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SublimeText(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M.003 23.617V17.93a.72.72 0 0 1 .457-.654l.005-.002l7.453-2.361l-7.454-2.366a.706.706 0 0 1-.398-.377l-.002-.005a.434.434 0 0 1-.061-.222v-.014v.001v-5.737c0-.083.023-.161.064-.227l-.001.002a.7.7 0 0 1 .395-.379l.005-.002L18.014.023a.343.343 0 0 1 .461.365V.386v5.686a.724.724 0 0 1-.457.654l-.005.002l-7.375 2.338l7.378 2.339a.722.722 0 0 1 .462.656v5.691a.53.53 0 0 1-.011.106l.001-.003a.725.725 0 0 1-.45.558l-.005.002L.464 23.979a.429.429 0 0 1-.128.022a.345.345 0 0 1-.332-.387v.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Subscript(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23.476 22.872h-1.601v-.009c0-.169.386-.442.696-.66c.618-.435 1.385-.978 1.385-1.918v-.013c0-.857-.695-1.552-1.552-1.552l-.059.001h.003a1.444 1.444 0 0 0-1.54 1.402v.002a.578.578 0 1 0 1.149.09v-.002c0-.166.042-.353.375-.353l.032-.001c.237 0 .429.192.429.429v.014v-.001c0 .353-.464.691-.912 1.014c-.554.399-1.182.852-1.182 1.551v.601a.561.561 0 0 0 .56.536h2.215a.566.566 0 0 0 .002-1.128h-.002zM18.335.253a1.127 1.127 0 0 0-1.582.159l-.002.002l-7.378 9.055L1.995.414a1.124 1.124 0 0 0-2.012.692c0 .279.101.534.269.73l-.001-.002l7.671 9.414l-7.671 9.414a1.125 1.125 0 0 0 1.743 1.421l.001-.001l7.378-9.055l7.378 9.055a1.126 1.126 0 0 0 1.743-1.423l.002.002l-7.672-9.412l7.671-9.414a1.127 1.127 0 0 0-.159-1.582l-.002-.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Subway(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M1.338 24h3.318l1.883-3.771H4.793zm9.877-3.771L13.098 24h3.318l-3.456-3.771zm3.153-6.463a.679.679 0 0 1-.678.679H4.065a.679.679 0 0 1-.679-.679l-.347-7.782c0-.375.304-.678.679-.678h10.321c.375 0 .679.304.679.678zm.294 2.666a.582.582 0 0 1-.58.58h-2.328a.582.582 0 0 1-.58-.58a.582.582 0 0 1 .58-.58h2.331c.32.001.58.261.58.581zm-8.132 0a.582.582 0 0 1-.58.58H3.62a.582.582 0 0 1-.58-.58a.582.582 0 0 1 .58-.58h2.331c.32.001.58.261.58.581zm-.121-15.04h4.934l-.45 2.23H6.859zm9.979 5.236a1.29 1.29 0 0 0-.535.115l.008-.003V4.41a.787.787 0 0 0-.787-.787h-2.759l.45-2.23h1.928V.003H3.068v1.39h1.927l.45 2.23H2.681a.79.79 0 0 0-.787.787v2.344a1.304 1.304 0 0 0-.555-.126A1.34 1.34 0 0 0 .002 7.964v1.433c.003.737.6 1.334 1.337 1.336c.202-.001.393-.047.563-.129l-.008.003v8.054a.79.79 0 0 0 .787.787h12.394a.788.788 0 0 0 .786-.787v-8.045c.155.069.336.11.527.11a1.34 1.34 0 0 0 1.336-1.336V7.958a1.34 1.34 0 0 0-1.336-1.336z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Suitcase(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.184 3.979h-5.506v-1.44a2.546 2.546 0 0 0-2.541-2.542H10.75a2.546 2.546 0 0 0-2.542 2.541v1.44H2.702A2.705 2.705 0 0 0-.002 6.683v13.26a2.7 2.7 0 0 0 2.701 2.701h.796a1.36 1.36 0 1 0 2.712-.005v.005h12.456a1.356 1.356 0 1 0 2.712 0h.796a2.7 2.7 0 0 0 2.701-2.701V6.661a2.683 2.683 0 0 0-2.683-2.683h-.007zm-12.085-1.44c0-.36.292-.652.652-.652h3.386c.36 0 .652.292.652.652v1.44h-4.683v-1.44zm-4.527 17.22h-.001a2.201 2.201 0 1 1 2.201-2.201v.001c0 1.215-.985 2.2-2.2 2.201zm3.579-8.88l1.44-3.867l3.867 1.44l-1.44 3.868zm6.24 8.336l-.956-3.305l7.269-2.097l.956 3.305z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SuitcaseAlt(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.109 9.267h-2.281v-.689h-.998V1.164h1.259V0H4.4v1.164h1.255v7.414h-.998v.689H2.375a2.383 2.383 0 0 0-2.376 2.376v8.72a2.383 2.383 0 0 0 2.376 2.376h.689a1.26 1.26 0 1 0 2.52 0h8.294a1.26 1.26 0 1 0 2.52 0h.713a2.383 2.383 0 0 0 2.376-2.376v-8.72a2.383 2.383 0 0 0-2.376-2.376zM6.059 1.164h7.366v7.414h-.998v.689h-5.37v-.689h-.998zm-2.281 18.06L2 17.251l5.251-4.729l1.782 1.972zm7.272-8.126h2.851v2.662H11.05zm3.612 9.338a2.662 2.662 0 1 1 2.662-2.662v.005a2.658 2.658 0 0 1-2.657 2.657z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sun(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m24 14.358l-3.04-2.965l2.608-3.348l-4.114-1.051l.584-4.204l-4.088 1.142L14.35 0l-2.965 3.04L8.037.432L6.986 4.546l-4.204-.584l1.142 4.087l-3.932 1.596l3.04 2.966l-2.608 3.348l4.114 1.051l-.59 4.204l4.087-1.142l1.6 3.932l2.965-3.04l3.348 2.608l1.051-4.114l4.205.59l-1.142-4.087zm-9.719 4.302a7.04 7.04 0 1 1 4.378-8.99l.015.049c.24.679.378 1.461.378 2.276a7.042 7.042 0 0 1-4.722 6.649l-.049.015z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Sunglasses(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.316C6.099 1.316 1.316 6.099 1.316 12S6.099 22.684 12 22.684S22.684 17.901 22.684 12c-.012-5.896-4.788-10.672-10.683-10.684zm0 22.297C5.586 23.613.387 18.414.387 12S5.586.387 12 .387S23.613 5.586 23.613 12v.015c0 6.405-5.192 11.597-11.597 11.597z"/><path fill="currentColor" d="M12 24C5.386 23.966.034 18.614 0 12.003V12C0 5.373 5.373 0 12 0s12 5.373 12 12c-.034 6.614-5.386 11.966-11.997 12zM12 .774C5.8.774.774 5.8.774 12S5.8 23.226 12 23.226S23.226 18.2 23.226 12C23.222 5.802 18.198.779 12.001.774zm0 22.297C5.886 23.071.929 18.114.929 12S5.886.929 12 .929S23.071 5.886 23.071 12S18.114 23.071 12 23.071m0-21.368C6.313 1.703 1.703 6.313 1.703 12S6.313 22.297 12 22.297S22.297 17.687 22.297 12c-.026-5.676-4.62-10.271-10.294-10.297z"/><path fill="currentColor" d="M9.677 9.91v.009c0 1.15-.932 2.082-2.082 2.082h-.009a2.09 2.09 0 1 1 2.09-2.09zm8.904 0v.009c0 1.15-.932 2.082-2.082 2.082h-.009a2.09 2.09 0 0 1 0-4.18a2.132 2.132 0 0 1 2.09 2.088v.002zm-2.168 7.277H7.51a.499.499 0 0 1-.464-.463v-.002a.499.499 0 0 1 .463-.464h8.905a.499.499 0 0 1 .464.463v.021c0 .246-.2.446-.446.446h-.02h.001z"/><path fill="currentColor" d="M16.413 17.574H7.51a.852.852 0 0 1 0-1.704h8.903a.852.852 0 0 1 0 1.704m-8.826-.929c-.077 0-.077.077 0 0c-.02.02-.033.047-.033.077s.013.058.033.077h8.903a.077.077 0 0 0 .077-.077c0-.077 0-.077-.077-.077zm8.826-3.871a2.865 2.865 0 1 1 2.865-2.865v.007a2.858 2.858 0 0 1-2.858 2.858zm-8.826 0a2.908 2.908 0 0 1-2.864-2.862v-.009a2.858 2.858 0 0 1 2.858-2.858h.007a2.908 2.908 0 0 1 2.864 2.862v.002a2.963 2.963 0 0 1-2.86 2.864zm15.329-3.329h-2.787a3.829 3.829 0 0 0-3.712-3.329h-.004a3.758 3.758 0 0 0-3.714 3.311l-.002.018h-1.471a3.829 3.829 0 0 0-3.712-3.329H7.51a3.758 3.758 0 0 0-3.714 3.311l-.002.018H1.007v.929h2.787a3.829 3.829 0 0 0 3.712 3.329h.004a3.758 3.758 0 0 0 3.714-3.311l.002-.018h1.471a3.829 3.829 0 0 0 3.712 3.329h.004a3.758 3.758 0 0 0 3.714-3.311l.002-.018h2.787z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SunglassesAlt(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M32.094 17.878v-.018a.6.6 0 0 0-.048-.236l.001.004L25.87 3.684a1.076 1.076 0 0 0-.339-.414l-.003-.002l-4.309-3.07a1.023 1.023 0 0 0-.774-.175l.006-.001a1.003 1.003 0 0 0-.655.422l-.002.003a1.14 1.14 0 0 0-.165.923l-.002-.008c.079.321.288.581.567.727l.006.003l4.041 2.145c.199.108.356.273.45.474l.003.006l4.531 10.226a10.418 10.418 0 0 0-4.545-1.027l-.14.001h.007c-3.319 0-6.16 1.368-7.25 3.28a2.218 2.218 0 0 0-1.137-.32h-.001c-.432.015-.831.14-1.176.348l.011-.006c-1.082-1.933-3.93-3.31-7.267-3.31l-.129-.001c-1.632 0-3.178.368-4.559 1.027l.064-.028L7.625 4.708c.094-.208.252-.375.447-.477l.005-.003l4.041-2.145a1.133 1.133 0 0 0 .405-1.649l.002.004a1.007 1.007 0 0 0-1.427-.247l.003-.002l-4.309 3.07a1.01 1.01 0 0 0-.34.41l-.003.006L.272 17.618c-.009.018-.009.028-.018.046a3.415 3.415 0 0 0-.259 1.293v.001c0 2.783 3.458 5.04 7.73 5.04s7.73-2.256 7.73-5.04c0-.111-.01-.231-.018-.342a1.22 1.22 0 0 1 .725-.332h.014c.259.018.489.131.657.305c-.01.12-.018.25-.018.37c0 2.783 3.458 5.04 7.73 5.04s7.73-2.256 7.73-5.04a3.295 3.295 0 0 0-.194-1.103l.007.023zM7.813 16.01c-3.098 0-5.92 1.535-5.92 3.218a.416.416 0 1 1-.832 0c0-2.201 3.088-4.059 6.75-4.059a.421.421 0 0 1 0 .842zm16.847 0c-3.098 0-5.92 1.535-5.92 3.218a.416.416 0 1 1-.832 0c0-2.201 3.088-4.059 6.75-4.059a.421.421 0 0 1 0 .842z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Superscript(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23.477 4.149h-1.602V4.14c0-.169.386-.442.696-.66c.618-.435 1.385-.978 1.385-1.918v-.014c0-.857-.695-1.552-1.552-1.552l-.059.001h.003a1.444 1.444 0 0 0-1.54 1.402v.002a.578.578 0 1 0 1.149.09v-.002c0-.166.042-.353.375-.353l.032-.001c.237 0 .429.192.429.429v.014v-.001c0 .353-.464.691-.912 1.014c-.554.399-1.182.852-1.182 1.551v.601a.561.561 0 0 0 .56.536h2.215a.566.566 0 0 0 .002-1.128h-.002zm-5.141-2.396a1.127 1.127 0 0 0-1.582.159l-.002.002l-7.378 9.055l-7.378-9.055A1.126 1.126 0 0 0 .253 3.337l-.002-.002l7.671 9.415l-7.671 9.415a1.125 1.125 0 0 0 1.744 1.422l.001-.001l7.378-9.054l7.379 9.054a1.124 1.124 0 0 0 2.012-.692c0-.279-.101-.534-.269-.73l.001.002l-7.672-9.414l7.671-9.415a1.127 1.127 0 0 0-.159-1.582z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func SurgicalKnife(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M25.274 23.077h.967v.922h-.967zm1.942 0h.967v.922h-.967zm-7.762 0h.967v.922h-.967zm3.878 0h.967v.922h-.967zm-1.934 0h.967v.922h-.967zm-14.208 0c-.908.351-2.002.67-3.129.9l-.138.024h14.729v-.923zM18.344 10.14l7.898-7.898l-2.24-2.24L0 24a20.794 20.794 0 0 0 18.299-13.715zm5.92-9.239l.659.659L17.505 9l-.661-.659z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Surprised(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.316C6.099 1.316 1.316 6.099 1.316 12S6.099 22.684 12 22.684S22.684 17.901 22.684 12c-.012-5.896-4.788-10.672-10.683-10.684zm0 22.297C5.586 23.613.387 18.414.387 12S5.586.387 12 .387S23.613 5.586 23.613 12v.015c0 6.405-5.192 11.597-11.597 11.597z"/><path fill="currentColor" d="M12 24C5.386 23.966.034 18.614 0 12.003V12C0 5.373 5.373 0 12 0s12 5.373 12 12s-5.373 12-12 12M12 .774C5.8.774.774 5.8.774 12S5.8 23.226 12 23.226S23.226 18.2 23.226 12C23.222 5.802 18.198.779 12.001.774zm0 22.297C5.886 23.071.929 18.114.929 12S5.886.929 12 .929S23.071 5.886 23.071 12S18.114 23.071 12 23.071m0-21.368C6.313 1.703 1.703 6.313 1.703 12S6.313 22.297 12 22.297S22.297 17.687 22.297 12v-.005c0-5.684-4.608-10.292-10.292-10.292z"/><path fill="currentColor" d="M9.677 9.91v.009c0 1.15-.932 2.082-2.082 2.082h-.018a2.082 2.082 0 0 1-2.082-2.082V9.91a2.09 2.09 0 0 1 4.18 0zm8.904 0v.009c0 1.15-.932 2.082-2.082 2.082h-.009a2.09 2.09 0 1 1 2.09-2.09zm-1.471 7.045c0 1.471-2.245 2.632-5.032 2.632s-5.032-1.161-5.032-2.632s2.245-2.632 5.032-2.632s5.032 1.239 5.032 2.632"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Svn(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.726 17.352a.333.333 0 1 1-.666 0a.333.333 0 0 1 .666 0m-1.13 2.061a.333.333 0 1 1-.666 0a.333.333 0 0 1 .666 0m0-2.061a.333.333 0 1 1-.666 0a.333.333 0 0 1 .666 0m1.13 2.061a.333.333 0 1 1-.666 0a.333.333 0 0 1 .666 0"/><path fill="currentColor" d="M32.426 0c-.383.107-.773.275-1.078.328C20.194 2.501 9.55 5.112.001 8.89v6.184c.509-.04 8.506-.858 9.707-.955c2.564-.208 11.177-1.157 14.528-.925c.683.089 1.689.236 1.799.872c.06.347-.24.654-.436.836c-.055.059-.096.134-.182.16a19.611 19.611 0 0 1-5.438 2.525l-.141.035C14.234 19.46 7.532 21.011.645 21.985l-.646.075v1.941h.715c7.049-1.13 13.723-2.892 20.195-4.574c3.984-1.035 7.96-2.576 11.666-3.939V8.401c-.685.099-4.098.577-13.582 1.494c-7.68.743-9.798.53-10.382.348a.625.625 0 0 1-.06-.045l.001.001a.363.363 0 0 1-.181-.179l-.001-.002c-.219-.236-.037-.663.16-.8a.432.432 0 0 1 .197-.199l.003-.001a9.455 9.455 0 0 1 2.371-1.376l.064-.023c1.527-.69 3.405-1.377 5.336-1.938l.318-.079c.987-.287 1.975-.583 2.98-.854c3.356-.964 7.67-1.914 12.061-2.632l.713-.096V.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Swarm(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24.123 13.507c1.128 2.534.869 5.111-.579 5.755s-3.536-.887-4.664-3.421s-.869-5.111.579-5.755s3.536.887 4.664 3.421"/><path fill="currentColor" d="M16.563 16.5c-1.065-2.508-1.2-5.026-.506-6.64c-2.962.9-5.08 3.606-5.08 6.808l.002.164v-.008c0 3.962 2.998 7.174 6.697 7.174a6.474 6.474 0 0 0 4.936-2.326l.008-.01c-2.018.378-4.627-1.794-6.056-5.162zm-8.216-.631s-3.109 3.826-2.63 5.262s5.979 2.391 5.979 2.391a23.816 23.816 0 0 1-2.089-3.22l-.063-.128a27.804 27.804 0 0 1-1.163-4.112zm4.305-7.059S5.387-1.757.764 4.407c-2.604 3.474 1.935 5.446 6.384 5.724a8.944 8.944 0 0 0 5.541-1.343zm5.262-7.734c1.03.594 1.008 2.56-.049 4.39s-2.748 2.832-3.778 2.237c-1.03-.594-1.008-2.56.049-4.39s2.748-2.832 3.778-2.237"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Swift(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24.698 16.885c.063-.126.063-.252.126-.377c1.512-5.923-2.08-12.853-8.254-16.507a14.39 14.39 0 0 1 3.132 8.995a14.55 14.55 0 0 1-.313 3.009l.017-.095c-.087.39-.195.728-.33 1.052l.015-.041a5.633 5.633 0 0 1-.524-.327l.019.012C13.829 9.457 9.685 6.031 5.935 2.217l-.011-.011c2.438 3.602 4.965 6.75 7.722 9.678l-.036-.039C9.489 9.205 5.908 6.467 2.571 3.464l.076.067c.474.784.972 1.463 1.524 2.093l-.014-.017c3.3 4.423 6.99 8.272 11.099 11.626l.117.092c-3.213 1.955-7.687 2.08-12.224 0a18.078 18.078 0 0 1-3.195-1.923l.046.033c2.022 3.147 4.88 5.601 8.261 7.079l.12.047c1.693.839 3.687 1.331 5.796 1.331s4.103-.491 5.874-1.365l-.078.035h.064c.146-.065.271-.15.378-.252h-.001c1.512-.756 4.474-1.573 6.112 1.573c.44.882 1.259-3.213-1.827-6.994z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Swimsuit(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.867 12.65c-.16-2.031-.754-3.991-2.48-5.369a.8.8 0 0 0-.268-.141l-.006-.001a5.984 5.984 0 0 0-.96-1.795l.01.013a10.115 10.115 0 0 0-4.579-3.198l-.071-.021c.493-.16 1.194-.463 1.372-.96a.83.83 0 0 0-.19-.82a.813.813 0 0 0-.609-.362h-.003c-.564 0-.96.831-1.146 1.502c-.19-.671-.582-1.502-1.146-1.502a.796.796 0 0 0-.611.359l-.002.003a.828.828 0 0 0-.188.825l-.002-.006c.178.493.88.796 1.372.96a10.046 10.046 0 0 0-4.636 3.201l-.015.019a5.85 5.85 0 0 0-.938 1.74l-.012.042a.923.923 0 0 0-.275.144l.002-.002C.757 8.659.166 10.613.006 12.65v.024c0 .356.25.655.583.729l.005.001c.974.247 2.093.389 3.244.389c1.362 0 2.679-.198 3.921-.568l-.097.025a.863.863 0 0 0 .386-.202l-.001.001c.288.062.619.099.958.101h.029c.28 0 .554-.029.818-.085l-.026.005c.102.086.226.15.362.183l.006.001c1.145.345 2.461.543 3.823.543c1.151 0 2.27-.142 3.338-.409l-.095.02a.745.745 0 0 0 .604-.731v-.027v.001zM9.487.998c.196-.416.416-.659.6-.659c.107 0 .219.08.345.232a.504.504 0 0 1 .136.503l.001-.004c-.143.4-.898.677-1.372.808c.072-.334.172-.628.301-.906zm-2.173.071a.51.51 0 0 1 .136-.5c.124-.154.24-.232.344-.232c.184 0 .4.24.6.659c.119.252.219.546.286.852l.005.028a4.31 4.31 0 0 1-.54-.182l.028.01c-.474-.19-.782-.416-.862-.636zm2.32 11.237a.62.62 0 0 0-.041.256v-.001a3.826 3.826 0 0 1-1.313-.028l.024.004v-.008a.669.669 0 0 0-.043-.237l.002.005C7.337 9.921 5.697 7.872 3.375 7.13a.85.85 0 0 0-.223-.041h-.003c.22-.586.507-1.092.859-1.545l-.01.013a9.684 9.684 0 0 1 4.332-3.05l.068-.021a1.62 1.62 0 0 0-.468.857l-.002.01a.173.173 0 0 0 .136.196h.033c.082 0 .15-.058.164-.136v-.001c.093-.4.344-.728.682-.923l.007-.004c.342.202.591.529.687.916l.002.01a.16.16 0 0 0 .158.137h.038a.173.173 0 0 0 .138-.169l-.002-.028v.001a1.622 1.622 0 0 0-.468-.867h-.001a9.677 9.677 0 0 1 4.375 3.035l.014.017c.314.397.583.851.786 1.34l.014.038c.024.066.048.118.065.172a1.014 1.014 0 0 0-.232.044l.007-.002c-2.338.746-3.98 2.794-4.898 5.176zm-8.206 5.215v1.12c3.51.63 6.278 3.48 6.871 5.36h1.318c.594-1.877 3.36-4.734 6.871-5.36v-1.12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableOne(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M28 0H1.333C.596 0-.001.597-.001 1.334v21.333c0 .737.597 1.334 1.334 1.334H28c.737 0 1.334-.597 1.334-1.334V1.334C29.334.597 28.737 0 28 0M13.334 8.001v5.333H2.667V8.001zm2.666 0h10.667v5.333H16zM2.667 16h10.667v5.333H2.667zM16 21.333V16h10.667v5.333z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TableTwo(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M28 0H1.333C.596 0-.001.597-.001 1.334v21.333c0 .737.597 1.334 1.334 1.334H28c.737 0 1.334-.597 1.334-1.334V1.334C29.334.597 28.737 0 28 0M11.556 16v-2.667h6.222V16zm6.222 2.667v2.667h-6.222v-2.667zm0-10.667v2.667h-6.222V8.001zm2.667 0h6.222v2.667h-6.222zM8.889 8v2.667H2.667V8.001zm-6.222 5.334H8.89v2.667H2.667zm17.778 0h6.222v2.667h-6.222zM2.667 18.667H8.89v2.667H2.667zm17.778 2.666v-2.667h6.222v2.667z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tablet(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.24 0H1.952A2.117 2.117 0 0 0 0 2.257V2.25v19.499a2.115 2.115 0 0 0 1.942 2.25h14.297a2.116 2.116 0 0 0 1.949-2.257v.007V2.25A2.122 2.122 0 0 0 16.246 0l-.008-.001zM9.096 22.5a1.498 1.498 0 1 1 .002-2.996a1.498 1.498 0 0 1-.002 2.996h-.004zm7.781-5.062a.564.564 0 0 1-.56.56H1.875a.564.564 0 0 1-.56-.56V2.813a.564.564 0 0 1 .56-.56h14.438a.564.564 0 0 1 .56.56z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TabletAlt(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.308 0H1.954A2.121 2.121 0 0 0 0 2.257v-.006V21.75A2.12 2.12 0 0 0 1.948 24h14.356a2.12 2.12 0 0 0 1.954-2.256v.006V2.251a2.114 2.114 0 0 0-1.942-2.25h-.007zM9.131 22.5a1.498 1.498 0 1 1 .002-2.996a1.498 1.498 0 0 1-.002 2.996h-.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tablets(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 .001c-6.627 0-12 5.373-12 12s5.373 12 12 12s12-5.373 12-12s-5.373-12-12-12m0 1.614c4.997.005 9.168 3.533 10.164 8.234l.012.068H1.823C2.834 5.15 7.004 1.623 11.999 1.615zm0 20.772c-4.997-.005-9.169-3.533-10.165-8.234l-.012-.068h20.345c-1.003 4.769-5.173 8.298-10.168 8.301z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Taxi(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M25.678 24v-3.298H7.859V24H1.553v-8.728a6.072 6.072 0 0 1 1.863-4.238l.002-.002l-1.823-.3A1.904 1.904 0 0 1 .001 8.857a1.9 1.9 0 0 1 2.218-1.874l-.011-.002l3.127.51c.051.01.095.022.137.036l-.007-.002l2.014-5.126h6.434l.48-2.4h6l.48 2.4h6.062l2.015 5.13a.93.93 0 0 1 .12-.033l.008-.001l3.13-.51a1.906 1.906 0 0 1 2.18 1.559l.002.011a1.904 1.904 0 0 1-1.555 2.179l-.011.002l-2.532.415a6.075 6.075 0 0 1 1.739 4.115v3.714c0 .129-.015.254-.043.374l.002-.011v4.655zm-3.884-7.617a1.269 1.269 0 0 0 1.266 1.266h5.087a1.266 1.266 0 0 0 0-2.532h-5.085a1.27 1.27 0 0 0-1.266 1.266zm-17.745 0a1.268 1.268 0 0 0 1.266 1.266h5.087a1.266 1.266 0 0 0 0-2.532H5.316a1.27 1.27 0 0 0-1.266 1.265zm2.973-6.608h20.264c.034 0 .067.006.099.006l-2-5.106H9.031z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ted(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.431 10.223H0V8.002h7.53v2.221H5.1v6.446H2.431zm5.514-2.221h7.31v2.221h-4.64v1.086h4.64v2.063h-4.64v1.08h4.64v2.225h-7.31zm10.43 6.452h1.046c1.664 0 1.907-1.352 1.907-2.166a1.926 1.926 0 0 0-2.124-2.063l.008-.001h-.854v4.23zm-2.67-6.452h4.384c2.891 0 3.911 2.135 3.911 4.32c0 2.66-1.409 4.351-4.434 4.351h-3.862V8.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Telegram(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 0C5.373 0 0 5.373 0 12s5.373 12 12 12s12-5.373 12-12S18.627 0 12 0m5.894 8.221l-1.97 9.28c-.145.658-.537.818-1.084.508l-3-2.21l-1.446 1.394a.759.759 0 0 1-.6.295h-.005l.213-3.054l5.56-5.022c.24-.213-.054-.334-.373-.121L8.32 13.617l-2.96-.924c-.64-.203-.658-.64.135-.954l11.566-4.458c.538-.196 1.006.128.832.941z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tent(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23.36 21.44L13.174 3.996l2.101-3.6l-.686-.4l-1.872 3.213l-1.872-3.213l-.686.4l2.101 3.6L2.081 21.44H.001V24h25.446v-2.56zm-12.286-5.06a14.947 14.947 0 0 0 1.621-3.949l.023-.105c.966 3.835 3.282 7.017 6.401 9.085l.057.035H6.269a15.495 15.495 0 0 0 4.771-4.993l.039-.071z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tesla(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12 5.362l2.475-3.026c3.101.169 5.989.908 8.619 2.115l-.147-.061a7.552 7.552 0 0 1-3.18 2.42l-.051.018c-.146-1.439-1.154-1.79-4.354-1.79L12.001 24L8.62 5.034c-3.18 0-4.188.354-4.335 1.792a7.534 7.534 0 0 1-3.216-2.409L1.056 4.4c2.481-1.148 5.369-1.889 8.409-2.057l.061-.003zm0-3.9c.107-.002.233-.002.359-.002c3.967 0 7.738.84 11.144 2.352l-.175-.069a8.95 8.95 0 0 0 .65-1.331L24 2.349C20.449.886 16.329.026 12.01.001H12A32.625 32.625 0 0 0-.218 2.429L0 2.35c.224.555.449 1.015.703 1.458L.672 3.75C3.91 2.302 7.689 1.459 11.666 1.459l.352.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TestBottle(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.269 4.702H2.222A2.223 2.223 0 0 0 0 6.924v14.854C0 23.005.995 23.999 2.222 24h11.047a2.222 2.222 0 0 0 2.222-2.222V6.923a2.227 2.227 0 0 0-2.222-2.222zm.09 14.025H2.125V9.981H13.36zM2.314 3.925H13.18a.529.529 0 0 0 .526-.525V.526A.53.53 0 0 0 13.181 0H2.315a.531.531 0 0 0-.526.525V3.4c0 .29.235.525.525.525h.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TestTube(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.585 0H0v1.958h.73v17.978a4.063 4.063 0 1 0 8.126 0V1.958h.73zm-1.92 19.938a2.87 2.87 0 0 1-5.738.004v-1.044H3.9v-1.15H1.927v-2.025H3.9v-1.178H1.927V12.52H3.9v-1.178H1.927V9.317H3.9v-1.15H1.927V6.142H3.9v-1.15H1.927v-3.05h5.738zM13.129 0v1.958h.73v17.978a4.063 4.063 0 1 0 8.126 0V1.958h.73V0zm7.663 10.49h-5.737V9.326h1.973v-1.15h-1.973V6.151h1.973v-1.15h-1.973v-3.05h5.737z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TestTubeAlt(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.403 0h-9.586v1.958h.73v18.035a4.063 4.063 0 1 0 8.126 0v-.06v.003V1.958h.73zm-1.93 4.844h-5.735V1.952h5.737zM9.586 0H0v1.958h.73v18.035a4.063 4.063 0 1 0 8.126 0v-.06v.003V1.958h.73zm-1.92 19.938a2.87 2.87 0 0 1-5.736.005V1.961h5.737z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextHeight(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m22.111 14.056l-.657.704V8.658l.657.704a1.091 1.091 0 1 0 1.595-1.489l-2.558-2.739a.642.642 0 0 0-.046-.044l-.034-.031l-.051-.041l-.034-.026c-.02-.014-.041-.026-.062-.039l-.029-.018a.988.988 0 0 0-.08-.039l-.018-.009a1.059 1.059 0 0 0-.429-.088c-.118 0-.232.019-.339.053l.008-.002a1.038 1.038 0 0 0-.106.04l.007-.003l-.018.008a1.181 1.181 0 0 0-.08.039l-.028.017a1.345 1.345 0 0 0-.063.04l-.034.025c-.018.014-.035.026-.052.041l-.034.03a1.87 1.87 0 0 0-.046.044l-.013.012l-2.546 2.727a1.089 1.089 0 0 0 .811 1.82c.307 0 .585-.127.783-.332l.657-.704v6.103l-.657-.704a1.091 1.091 0 0 0-1.595 1.489l-.001-.001l2.559 2.741l.042.041l.038.034a.617.617 0 0 0 .048.038l.038.028a.565.565 0 0 0 .058.036l.034.02a.632.632 0 0 0 .074.036l.023.011c.03.013.061.025.093.035h.01c.097.032.208.05.324.051h.009c.156 0 .305-.033.439-.092l-.007.003l.012-.006c.029-.013.058-.026.085-.042l.024-.014c.023-.014.046-.027.067-.042l.03-.022l.056-.044l.03-.027a.73.73 0 0 0 .049-.047l.011-.011l2.546-2.727a1.089 1.089 0 0 0-.811-1.82c-.307 0-.585.127-.783.332zM12.727 4.8H1.091C.488 4.8 0 5.289 0 5.891v2.865a1.091 1.091 0 0 0 2.182 0v-.031v.002v-1.745h3.637v9.454H4.134a1.092 1.092 0 0 0-.002 2.182h5.553a1.092 1.092 0 0 0 .002-2.182H8V6.982h3.636v1.775a1.091 1.091 0 0 0 2.182 0v-.031v.002v-2.835a1.09 1.09 0 0 0-1.09-1.091z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TextWidth(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.528 21.103l.031-.034c.014-.016.027-.034.04-.051l.026-.034a.804.804 0 0 0 .039-.062l.018-.029c.014-.026.027-.051.039-.078l.009-.019a1.061 1.061 0 0 0 .088-.43c0-.118-.019-.232-.053-.338l.002.008a.994.994 0 0 0-.037-.099l-.009-.019a.945.945 0 0 0-.039-.078c-.005-.01-.012-.02-.018-.029s-.025-.042-.039-.062s-.017-.022-.026-.034l-.04-.051l-.031-.034c-.015-.015-.028-.031-.044-.045l-.012-.013l-2.727-2.546a1.091 1.091 0 0 0-1.492 1.588l-.001-.001l.704.657h-6.1l.704-.657a1.091 1.091 0 0 0-1.489-1.594l.001-.001L.331 19.58a1.528 1.528 0 0 0-.074.079l-.037.047a.364.364 0 0 0-.028.038l-.036.057l-.021.033a.834.834 0 0 0-.036.073l-.011.024c-.013.03-.024.061-.035.093v.009c-.032.097-.05.208-.051.324v.016c0 .153.032.298.088.43l-.003-.007l.007.014l.041.086l.015.025c.014.022.027.045.042.066l.023.031l.043.055l.028.032c.015.016.03.033.046.048l.011.011L3.07 23.71a1.091 1.091 0 0 0 1.489-1.594l-.704-.657h6.101l-.704.657a1.091 1.091 0 0 0 1.489 1.594l-.001.001l2.739-2.559a.665.665 0 0 0 .046-.046zM1.091 5.018c.603 0 1.091-.489 1.091-1.091V2.182h3.636v9.455H4.102a1.091 1.091 0 0 0 0 2.182h.032h-.002h5.582a1.091 1.091 0 0 0 0-2.182h-.032h.002h-1.686V2.182h3.636v1.776a1.091 1.091 0 0 0 2.182 0v-.032v.002v-2.835a1.09 1.09 0 0 0-1.09-1.091H1.09C.487.002-.001.491-.001 1.093v2.835c0 .603.489 1.091 1.091 1.091z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Thermometer(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.395 16.974c0-1.999.839-3.803 2.184-5.078l.003-.003V4.836a4.837 4.837 0 1 1 9.674 0v7.056a7.021 7.021 0 1 1-11.861 5.086zm4.201-12.137v7.501a.123.123 0 0 1-.006.039v-.001a1.218 1.218 0 0 1-.011.109l.001-.005a.939.939 0 0 1-.05.19l.002-.006a.864.864 0 0 1-.039.094l.002-.006a1.897 1.897 0 0 1-.046.088a.695.695 0 0 1-.056.085l.002-.002a.754.754 0 0 1-.06.073l.001-.001a1.051 1.051 0 0 1-.074.075l-.001.001c-.011.01-.02.022-.032.032a5.014 5.014 0 1 0 6.387.007l-.008-.007a.31.31 0 0 1-.031-.032c-.025-.022-.047-.046-.069-.07L26.507 13a1.022 1.022 0 0 1-.066-.079l-.002-.003a.823.823 0 0 1-.043-.066a.847.847 0 0 1-.055-.1l-.002-.006a.808.808 0 0 1-.064-.182l-.001-.007a.46.46 0 0 0-.011-.058l.001.003a.937.937 0 0 1-.015-.148V4.837a2.827 2.827 0 1 0-5.648.006v-.006zM0 16.974a6.98 6.98 0 0 1 2.186-5.079l.003-.003V4.837a4.836 4.836 0 1 1 9.672 0v7.056A7.021 7.021 0 1 1-.001 16.978zM4.201 4.837v7.501a.123.123 0 0 1-.006.039v-.001a.942.942 0 0 1-.01.108l.001-.005a.786.786 0 0 1-.049.188l.002-.006c-.011.027-.023.062-.036.088l-.047.088a.695.695 0 0 1-.056.085l.002-.002c-.02.027-.04.05-.06.073a1.051 1.051 0 0 1-.074.075l-.001.001c-.011.01-.02.022-.032.032a5.013 5.013 0 1 0 6.387.007l-.008-.007c-.011-.009-.019-.019-.028-.029a1.132 1.132 0 0 1-.077-.078l-.001-.001l-.062-.074c-.02-.025-.033-.049-.049-.074a.974.974 0 0 1-.051-.092l-.003-.006c-.011-.023-.019-.047-.028-.07a1.215 1.215 0 0 1-.036-.111l-.002-.007a.658.658 0 0 1-.012-.064a.827.827 0 0 1-.014-.139V4.835a2.827 2.827 0 1 0-5.648.006v-.006zm15.872 12.137a3.358 3.358 0 0 1 2.318-3.184l.024-.007V9.667a1.005 1.005 0 0 1 2.01 0v4.116a3.354 3.354 0 0 1 2.342 3.193a3.348 3.348 0 0 1-6.696 0zm-16.395 0a3.358 3.358 0 0 1 2.318-3.184l.024-.007V4.704a1.005 1.005 0 0 1 2.01 0v9.079a3.354 3.354 0 0 1 2.342 3.193a3.348 3.348 0 0 1-6.696 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ThermometerAlt(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m1.522 1.298l-.233.233c-.797.801-1.29 1.905-1.29 3.124s.493 2.323 1.29 3.124l12.577 12.577c.524.524 1.224.87 2.005.942l.013.001l4.014.362l1.974 1.974a1.247 1.247 0 1 0 1.765-1.759l-.001-.001l-1.974-1.974l-.362-4.014a3.316 3.316 0 0 0-.943-2.017L7.781 1.294A4.416 4.416 0 0 0 4.656.004a4.413 4.413 0 0 0-3.132 1.298zm18.013 13.381c.339.342.565.798.616 1.305l.001.009l.32 3.52l-.972.973l-3.52-.32a2.119 2.119 0 0 1-1.313-.617l-.452-.451l.856-.856l-.623-.625l-.855.856l-1.13-1.13l.856-.856l-.623-.625l-.855.856l-1.141-1.139l.855-.856l-.625-.625l-.855.856l-1.13-1.13l.856-.856l-.625-.625l-.855.856l-1.139-1.139l.855-.856l-.625-.625l-.855.856l-1.129-1.13l.855-.856l-.625-.625l-.855.856l-1.139-1.139l.855-.856l-.624-.625l-.856.856l-.943-.942c-.592-.593-.959-1.411-.959-2.315s.366-1.723.959-2.315l.233-.233c.593-.592 1.411-.959 2.315-.959s1.723.366 2.315.959z"/><path fill="currentColor" d="M4.389 4.397a.679.679 0 0 0 0 .952l13.28 13.28a.677.677 0 0 0 1.213-.416a.677.677 0 0 0-.26-.534l-.001-.001l-13.28-13.28c-.122-.121-.29-.195-.476-.195s-.353.074-.475.195z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ticket(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.817 5.669l4.504 4.501l-8.15 8.15l-4.501-4.504zm-3.006 13.944l8.8-8.8a.894.894 0 0 0 .001-1.28l-5.156-5.156a.926.926 0 0 0-1.28.001l-8.8 8.8a.894.894 0 0 0-.001 1.28l5.156 5.156a.926.926 0 0 0 1.28-.001m12.663-9.073L10.556 23.473c-.332.326-.787.527-1.289.527s-.957-.201-1.289-.527L6.184 21.68a2.74 2.74 0 0 0-3.875-3.874l.001-.001l-1.781-1.794c-.326-.332-.527-.787-.527-1.289s.201-.957.527-1.289L13.448.527C13.78.201 14.235 0 14.737 0s.957.201 1.289.527l1.781 1.781a2.74 2.74 0 1 0 3.874 3.874l.001-.001l1.794 1.781c.326.332.527.787.527 1.289s-.201.957-.527 1.289z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TicketAlt(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.396 11.484a1.073 1.073 0 0 0-1.451.059l-.001.001a1.085 1.085 0 0 0-.065 1.451l.003.003l-9.79 9.79l-.003.003a1.71 1.71 0 0 1-2.417 0l-.004-.004l-6.456-6.456l-.003-.003a1.71 1.71 0 0 1 0-2.417l.004-.004l9.79-9.79l.003.003c.416.416 1.09.416 1.506 0l.003-.003l.005-.005a1.065 1.065 0 0 0 0-1.506l-.003-.003l1.392-1.392l.003-.003a1.71 1.71 0 0 1 2.417 0l.004.004l6.456 6.456l.003.003a1.71 1.71 0 0 1 0 2.417l-.004.004zm-6.719 3.455a.262.262 0 0 0 .359 0l.637-.637l.003-.003a.25.25 0 0 0 0-.353l-.004-.004a.262.262 0 0 0-.359 0l-.638.638a.261.261 0 0 0 .001.358zm-1.236 1.237a.262.262 0 0 0 .359 0l.637-.637c.099-.099.099-.261 0-.36s-.261-.099-.36 0l-.637.637a.261.261 0 0 0 0 .36zm-3.48 3.479a.262.262 0 0 0 .36.002l2.015-2.015l.003-.003a.25.25 0 0 0 0-.353l-.004-.004a.26.26 0 0 0-.359 0L9.96 19.298a.262.262 0 0 0-.002.359zm7.741-8.891a1.052 1.052 0 1 0-.002 0zm-4.106 2.232l.003.003a.25.25 0 0 0 .354 0l.003-.003l.12-.12c.099-.099.099-.261 0-.36s-.261-.099-.36 0l-.12.12a.261.261 0 0 0 0 .36m-.84.84a.26.26 0 0 0 .359 0l.12-.12c.099-.099.099-.261 0-.36s-.261-.099-.36 0l-.118.118a.26.26 0 0 0-.002.362zm-.84.841a.262.262 0 0 0 .36 0l.12-.12c.099-.099.099-.261 0-.36s-.261-.099-.36 0l-.12.12a.26.26 0 0 0 .002.359zm-.84.84a.262.262 0 0 0 .359 0l.12-.12l.003-.003a.25.25 0 0 0 0-.353l-.004-.004a.26.26 0 0 0-.358.001l-.12.12a.26.26 0 0 0-.001.358zm-.84.839a.262.262 0 0 0 .359 0l.12-.12c.099-.099.099-.261 0-.36s-.261-.099-.36 0l-.12.12a.261.261 0 0 0 0 .36zm-.84.84a.262.262 0 0 0 .359 0l.12-.12c.099-.099.099-.261 0-.36s-.261-.099-.36 0l-.12.12a.261.261 0 0 0 0 .36zm-.84.84a.26.26 0 0 0 .359 0l.119-.119c.099-.099.099-.261 0-.36s-.261-.099-.36 0l-.119.119a.261.261 0 0 0 0 .36zm6.996-9.42a1.057 1.057 0 0 0 1.488 0l.007-.007a1.042 1.042 0 0 0 0-1.474l-.007-.007a1.052 1.052 0 1 0-1.488 1.488m-6.96 6.096l.013.013a.711.711 0 0 0 1.2-.646l-.001-.003l-.288.096a.364.364 0 0 1-.093.357l-.002.002a.385.385 0 0 1-.589-.062l.001.002c-.216-.216-.204-.444-.06-.588l.002-.002a.364.364 0 0 1 .355-.093l.002.001l.096-.288a.68.68 0 0 0-.636.181l-.024.026l.001-.001l-.021.02a.67.67 0 0 0 .045.987l-.002-.002zm2.304-1.295l.264-.264l-.948-.229l-.108-.732l-.252.252l.13.777l-.454-.454l-.227.227L10.307 14l.227-.227l-.25-.25l-.036-.276zm.624-2.28l.288-.288l.815.815l.227-.227l-.816-.816l.3-.3l-.192-.192l-.814.814zm-.108 1.739l.684-.684l-.18-.18l-.454.454l-.24-.24l.397-.397l-.195-.195l-.397.397l-.214-.214l.444-.444l-.18-.18l-.674.674zm-2.88 2.88l.227-.227l-1.008-1.008l-.227.227zm-1.894.262l.288-.288l.816.816l.227-.227l-.816-.816l.288-.288l-.192-.192l-.805.805zm6.778-9.538l.001.001a1.052 1.052 0 1 0-.001-.001m-6.768 6.768a.262.262 0 0 0 .359 0l1.92-1.92c.099-.099.099-.261 0-.36s-.261-.099-.36 0l-1.92 1.92a.261.261 0 0 0 0 .36zm-1.311-.062a.26.26 0 0 0 .359 0l3.156-3.156c.099-.099.099-.261 0-.36s-.261-.099-.36 0l-3.155 3.155a.261.261 0 0 0 0 .36zm12.708-2.797l-.005-.005a.537.537 0 0 1 0-.758l.005-.005a.543.543 0 1 1 0 .768m-2.136-2.136a.543.543 0 1 1 .768 0l-.005.005a.537.537 0 0 1-.758 0l-.002-.002zm-2.136-2.136a.543.543 0 1 1 .768 0l-.005.005a.537.537 0 0 1-.758 0l-.002-.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tinder(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.458 9.451c.044.072.122.12.212.12h.002c.06 0 .114-.022.157-.057l.013-.014c.387-.324.737-.656 1.064-1.01l.007-.007a7.493 7.493 0 0 0 1.873-4.969c0-1.02-.203-1.994-.571-2.881l.018.05A.5.5 0 0 1 9.382.09l.001-.001a.495.495 0 0 1 .616.034c10.875 10.114 8 17.818 7.785 18.337c-.87 3.141-4.335 5.414-8.444 5.53c-.138.008-.242.008-.363.008C4.125 23.998 0 21.009 0 17.191v-.06c0-5.3 4.8-10.522 5.009-10.744a.44.44 0 0 1 .528-.104l-.003-.001a.475.475 0 0 1 .291.438v.014v-.001a4.693 4.693 0 0 0 .643 2.721l-.012-.022v.018z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tl(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.488 0h2.411v6.2c2.434-.88 4.89-1.805 7.341-2.693V5.5c-2.454.88-4.889 1.775-7.341 2.655v1.027c2.451-.888 4.907-1.813 7.341-2.693v1.978c-2.433.905-4.894 1.78-7.341 2.67v10.277a9.526 9.526 0 0 0 7.53-5.681l.024-.062a9.315 9.315 0 0 0 .746-3.682v-.008h2.502v.38a12.025 12.025 0 0 1-1.279 5.052l.032-.069a12.318 12.318 0 0 1-2.676 3.536l-.009.008a11.83 11.83 0 0 1-8.023 3.117c-.443 0-.88-.024-1.311-.071l.053.005v-11.92c-1.506.53-2.99 1.083-4.488 1.62v-1.92c1.501-.548 3.006-1.091 4.488-1.658V9.042c-1.493.51-2.99 1.082-4.488 1.613v-1.91c1.491-.56 3.009-1.094 4.488-1.666z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToggleOff(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 11.617A6.818 6.818 0 0 1 6.813 4.8h10.371a6.817 6.817 0 1 1 0 13.634H6.818a6.818 6.818 0 0 1-6.817-6.813zm6.817 4.543a4.543 4.543 0 1 0-.003-9.085a4.543 4.543 0 0 0 .001 9.085h.003z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ToggleOn(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24 11.617a6.818 6.818 0 0 1-6.813 6.817H6.816a6.817 6.817 0 1 1 0-13.634h10.366a6.818 6.818 0 0 1 6.817 6.813zm-6.817-4.545a4.542 4.542 0 1 0 0 9.084a4.542 4.542 0 0 0 .002-9.084z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tongue(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.316C6.099 1.316 1.316 6.099 1.316 12S6.099 22.684 12 22.684S22.684 17.901 22.684 12c-.012-5.896-4.788-10.672-10.683-10.684zm0 22.297C5.586 23.613.387 18.414.387 12S5.586.387 12 .387S23.613 5.586 23.613 12v.015c0 6.405-5.192 11.597-11.597 11.597z"/><path fill="currentColor" d="M12 24C5.386 23.966.034 18.614 0 12.003V12C0 5.373 5.373 0 12 0s12 5.373 12 12c-.034 6.614-5.386 11.966-11.997 12zM12 .774C5.8.774.774 5.8.774 12S5.8 23.226 12 23.226S23.226 18.2 23.226 12C23.222 5.802 18.198.779 12.001.774zm0 22.297C5.886 23.071.929 18.114.929 12S5.886.929 12 .929S23.071 5.886 23.071 12S18.114 23.071 12 23.071m0-21.368C6.313 1.703 1.703 6.313 1.703 12S6.313 22.297 12 22.297S22.297 17.687 22.297 12c-.026-5.676-4.62-10.271-10.294-10.297z"/><path fill="currentColor" d="M9.677 9.91a2.09 2.09 0 1 1-2.09-2.09h.009c1.15 0 2.082.932 2.082 2.082v.009zm8.904 0a2.09 2.09 0 0 1-4.18 0v-.009c0-1.15.932-2.082 2.082-2.082h.009a2.132 2.132 0 0 1 2.09 2.088v.002zm-1.084 4.49a.23.23 0 0 0-.155-.059a.23.23 0 0 0-.155.059a.353.353 0 0 0-.232.23l-.001.002c-.697 1.394-2.71 2.323-4.877 2.323s-4.258-.929-4.877-2.323l-.232-.232h-.31a.489.489 0 0 0-.231.623l-.001-.003a5.48 5.48 0 0 0 3.525 2.548l.036.007c.077 2.09 1.006 3.716 2.168 3.716s2.09-1.626 2.168-3.716a5.49 5.49 0 0 0 3.548-2.53l.014-.025a1.07 1.07 0 0 0-.385-.618z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Train(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.002 15.088a.684.684 0 0 1 .683-.682h2.743a.684.684 0 0 1 .683.682a.684.684 0 0 1-.683.682H2.682a.684.684 0 0 1-.683-.682zm.8-13.1h12.151a.801.801 0 0 1 .8.799v.001l-.41 9.16a.801.801 0 0 1-.8.8H3.21a.801.801 0 0 1-.8-.799v-.001L2 2.788a.801.801 0 0 1 .8-.8zm12.889 13.1a.684.684 0 0 1-.683.682h-2.746a.684.684 0 0 1-.683-.682a.684.684 0 0 1 .683-.682h2.746c.376.002.68.306.682.682zM1.58 18.642h14.594a.928.928 0 0 0 .926-.926V.927a.928.928 0 0 0-.926-.926H1.579a.93.93 0 0 0-.926.926v16.789a.928.928 0 0 0 .926.926zM0 24h3.906l2.218-4.44H4.069zm13.686-4.44h-2.055L13.849 24h3.906z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TrainTicket(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.437 23.473a.556.556 0 0 0-.087-.112l-.022-.021l.001.001a.86.86 0 0 0-1.134-.063l-.003.003l-.982-.982a.86.86 0 0 0-.06-1.137a.86.86 0 0 0-1.134-.063l-.003.003l-.982-.982a.86.86 0 0 0-.06-1.137a.86.86 0 0 0-1.134-.063l-.003.003l-.982-.982a.86.86 0 0 0-.06-1.137a.86.86 0 0 0-1.134-.063l-.003.003l-.982-.982a.86.86 0 0 0-.06-1.137c-.024-.024-.059-.036-.084-.06L14.445.649c.01.024.026.049.047.07l.015.014l-.001-.001l.006.006a.838.838 0 0 0 1.186 0l.006-.006c.024-.024.024-.048.048-.072l.983.983a.145.145 0 0 0-.062.038l-.01.011a.888.888 0 0 0 .098 1.296l-.003-.003l.006.006a.838.838 0 0 0 1.186 0l.006-.006c.024-.024.024-.048.048-.072l.982.982a.145.145 0 0 0-.062.038l-.01.011a.847.847 0 1 0 1.198 1.198c.024-.024.024-.048.048-.072l.982.982a.145.145 0 0 0-.063.039l-.009.01a.847.847 0 1 0 1.198 1.198c.024-.024.024-.048.048-.072l.982.982a.145.145 0 0 0-.062.038l-.01.011l-.006.006a.838.838 0 0 0 0 1.186l.005.005l.019.02a.575.575 0 0 0 .116.09l-.003-.002zm9.26-10.889l.12.12a.254.254 0 0 0 .359-.359l-.122-.122a.254.254 0 0 0-.359.359zm-.012-4.754l2.108 2.108a.254.254 0 0 0 .359-.359l-2.108-2.108a.254.254 0 0 0-.359.359m-5.452-3.89l6.778 6.778a.254.254 0 0 0 .359-.359l-6.778-6.778a.254.254 0 0 0-.359.359m2.144 10.527h.288a.288.288 0 0 0 .264-.264v.001a.247.247 0 0 0-.275-.228h-2.132a.247.247 0 0 0-.251.251a.213.213 0 0 0 .109.205l-.001-.001l-2.503 2.503a.219.219 0 0 0-.204-.107a.246.246 0 0 0-.252.251l-.001 2.131a.247.247 0 0 0 .251.251a.288.288 0 0 0 .264-.264v.001v-.288zm2.479-2.719l.12.12a.254.254 0 0 0 .359-.359l-.12-.12a.254.254 0 0 0-.359.359m-.838-.838l.12.12a.254.254 0 0 0 .359-.359l-.12-.12a.254.254 0 0 0-.359.359m-.838-.838l.12.12a.254.254 0 0 0 .359-.359l-.121-.121a.254.254 0 0 0-.359.359zm-9.928 1.736l.036.036a1.698 1.698 0 0 0 .167 2.204l2.434 2.434a1 1 0 0 0 1.414 0l2.515-2.515a1 1 0 0 0 0-1.414l-2.434-2.434a1.697 1.697 0 0 0-2.197-.173l-.008.006l-.036-.036a.26.26 0 0 0-.359 0l-1.532 1.532a.26.26 0 0 0 0 .359zm9.087-2.571l.12.12a.254.254 0 0 0 .359-.359l-.121-.121a.254.254 0 0 0-.359.359zm-.838-.838l.119.119a.254.254 0 0 0 .359-.359l-.119-.119a.254.254 0 0 0-.359.359m-.839-.839l.12.12a.254.254 0 0 0 .359-.359l-.12-.12a.254.254 0 0 0-.359.359m1.676-2.994l.599.599a.254.254 0 0 0 .359-.359l-.599-.599a.254.254 0 0 0-.359.359m-2.514 2.156l.12.12a.254.254 0 0 0 .359-.359l-.12-.12a.254.254 0 0 0-.359.359m1.209-3.461l.596.596a.254.254 0 0 0 .359-.359l-.596-.596a.254.254 0 0 0-.359.359m-2.047 2.622l.12.12a.254.254 0 0 0 .359-.359l-.121-.121a.254.254 0 0 0-.359.359zm-.839-.838l.12.12a.254.254 0 0 0 .359-.359l-.121-.121a.254.254 0 0 0-.359.359zm-.212 13.146l-.012-.611l3.125-3.125l.611.012zm-.995-2.072a.493.493 0 0 1-.695 0l-1.113-1.113l3.21-3.21l1.113 1.113a.493.493 0 0 1 0 .695zm.838-2.203a.779.779 0 1 0 0-1.101a.788.788 0 0 0 0 1.101m-1.652 1.653l.001.001a.776.776 0 0 0 1.099 0l.001-.001c.299-.306.299-.795 0-1.101s-.001-.001-.001-.001a.776.776 0 0 0-1.099 0l-.001.001a.788.788 0 0 0 0 1.101m2-2.024a.271.271 0 1 1 .385 0l-.003.003a.268.268 0 0 1-.378 0zm-1.653 1.652a.272.272 0 1 1 .384 0l-.003.003a.268.268 0 0 1-.378 0zm-1.701-.55l-.958-.958l-.006-.006a1.178 1.178 0 0 1 0-1.666l.005-.005l1.532-1.532l.006-.006c.46-.46 1.206-.46 1.666 0l.005.005l.958.958z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Transgender(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.713.429V.413c0-.228.185-.413.413-.413h.016h-.001h3.864c.234 0 .445.098.595.254a.82.82 0 0 1 .254.595V4.73a.413.413 0 0 1-.413.413h-.017h.001h-.873a.413.413 0 0 1-.413-.413v-.017v.001V2.92l-3.404 3.415a7.459 7.459 0 0 1 1.688 4.741v.07v-.004l.001.1a7.414 7.414 0 0 1-1.98 5.054l.004-.005A7.476 7.476 0 0 1 8.6 18.8l-.032.003v1.768H9.87c.228 0 .413.185.413.413v.017V21v.873a.413.413 0 0 1-.413.413h-.017h.001h-1.286v1.302a.413.413 0 0 1-.413.413h-.017h.001h-.873a.413.413 0 0 1-.413-.413v-.017v.001v-1.286H5.551a.413.413 0 0 1-.413-.413v-.017v.001v-.873c0-.228.185-.413.413-.413h.017h-.001h1.288V18.8a7.481 7.481 0 0 1-3.644-1.394l.021.015a7.74 7.74 0 0 1-2.478-2.947l-.02-.046a7.461 7.461 0 0 1-.737-3.258c0-.227.01-.453.03-.675l-.002.029a7.474 7.474 0 0 1 2.144-4.735l-.001.001a7.671 7.671 0 0 1 5.548-2.363c.767 0 1.508.112 2.207.321l-.054-.014a7.856 7.856 0 0 1 2.666 1.399l-.015-.012l3.415-3.402h-1.81a.413.413 0 0 1-.413-.413v-.017v.001zm-6 16.714l.087.001a5.763 5.763 0 0 0 4.152-1.759l.001-.002a5.765 5.765 0 0 0 1.761-4.153l-.001-.092v.005l.001-.087a5.763 5.763 0 0 0-1.759-4.152l-.002-.001C10.87 5.814 9.37 5.141 7.713 5.141s-3.157.674-4.24 1.762a5.765 5.765 0 0 0-1.76 4.245v-.005l-.001.087c0 1.629.674 3.101 1.759 4.152l.002.001a5.765 5.765 0 0 0 4.153 1.761l.092-.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func TransgenderAlt(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.143.429V.413c0-.228.185-.413.413-.413h.016h-.001h3.864c.234 0 .445.098.595.254a.82.82 0 0 1 .254.595V4.73a.413.413 0 0 1-.413.413h-.017h.001h-.872a.413.413 0 0 1-.413-.413v-.017v.001V2.92l-3.402 3.415a7.459 7.459 0 0 1 1.688 4.741v.07v-.004l.001.098a7.414 7.414 0 0 1-1.982 5.056l.004-.005a7.473 7.473 0 0 1-4.848 2.505l-.032.003v1.768h1.302c.228 0 .413.185.413.413v.017v-.001v.873a.413.413 0 0 1-.413.413h-.017h.001h-1.286v1.302a.413.413 0 0 1-.413.413h-.017h.001h-.873a.413.413 0 0 1-.413-.413v-.017v.001v-1.282H8.982a.413.413 0 0 1-.413-.413v-.017v.001v-.873c0-.228.185-.413.413-.413h.017h-.001h1.286V18.8a7.482 7.482 0 0 1-4.873-2.503l-.007-.008a7.407 7.407 0 0 1-1.976-5.049l.001-.105v-.055c0-1.806.638-3.462 1.7-4.757l-.01.013l-.697-.71l-1.356 1.486a.399.399 0 0 1-.294.141h-.001l-.024.001a.445.445 0 0 1-.285-.102l.001.001l-.64-.59a.368.368 0 0 1-.141-.288l-.001-.022c0-.113.043-.216.114-.292l1.403-1.54l-1.486-1.5v1.81a.413.413 0 0 1-.413.413h-.017h.001H.41a.413.413 0 0 1-.413-.413v-.017v.001V.85c0-.234.098-.445.254-.595A.82.82 0 0 1 .846.001h3.88c.228 0 .413.185.413.413v.017V.43v.873a.413.413 0 0 1-.413.413h-.017h.001h-1.782l1.42 1.433L5.5 1.89a.399.399 0 0 1 .294-.141h.001l.024-.001c.108 0 .208.038.285.102l-.001-.001l.64.59c.086.068.14.171.141.288l.001.022a.426.426 0 0 1-.114.292L5.568 4.367l.763.75a7.459 7.459 0 0 1 4.741-1.688h.07h-.004h.064c1.806 0 3.463.638 4.758 1.701l-.013-.01l3.415-3.402h-1.811a.413.413 0 0 1-.413-.413v-.017v.001zm-6 16.714l.087.001a5.763 5.763 0 0 0 4.152-1.759l.001-.002a5.765 5.765 0 0 0 1.761-4.153l-.001-.092v.005l.001-.087a5.763 5.763 0 0 0-1.759-4.152l-.002-.001a5.765 5.765 0 0 0-4.153-1.761l-.092.001h.005l-.087-.001a5.763 5.763 0 0 0-4.152 1.759l-.001.002a5.765 5.765 0 0 0-1.76 4.245v-.005l-.001.087c0 1.629.674 3.101 1.759 4.152l.002.001a5.765 5.765 0 0 0 4.153 1.761l.092-.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trash(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.313 4V2.144C16.313.96 15.353 0 14.169 0H7.831A2.142 2.142 0 0 0 5.69 2.189v-.002V4H0v2h.575c.196.023.372.099.515.214l-.002-.002c.119.157.203.346.239.552l.001.008l1.187 15.106c.094 1.84.094 2.118 2.25 2.118h12.462c2.16 0 2.16-.275 2.25-2.113l1.187-15.1c.036-.217.12-.409.242-.572l-.002.003a.994.994 0 0 1 .508-.212h.58v-2h-5.687zM7 2.187c0-.6.487-.938 1.106-.938h5.734c.618 0 1.162.344 1.162.938V4h-8zM6.469 20l-.64-12h1.269l.656 12zm5.225 0H10.32V8h1.375zm3.85 0h-1.275l.656-12h1.269z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Travis(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.948 14.472c.029.814.465.806.984.789c.698-.013 1.081-.227 1.431-.058c-.006-.019-.115-.347-.598-.289a1.017 1.017 0 0 0 .063-.515l.001.005c-.064-.814-.537-1.461-1.057-1.443s-.851.693-.822 1.51zm.774-.994h.004a.264.264 0 1 1-.004-.001zm-7.89 1.93h.323c.52-.005.951.044.944-.772s-.37-1.476-.887-1.47s-1.022.673-.992 1.488c0 .172.033.335.093.485l-.003-.009a.83.83 0 0 0-.585.409l-.002.004a2.282 2.282 0 0 1 1.12-.134zm.447-1.733h.004a.264.264 0 1 1-.005 0zM10.5 4.986h1.459v-.602h.328v2.779h-.496v1.555h2.553V7.163h-.492V4.384h.323v.602h1.459V2.927h-5.136v2.059z"/><path fill="currentColor" d="M24.181 8.957q-.014-.264-.05-.527a3.253 3.253 0 0 0-.044-.262a2.555 2.555 0 0 0-.068-.288l.005.019l-.019-.069l-.048-.039a4.61 4.61 0 0 0-.852-.516l-.029-.012a9.182 9.182 0 0 0-.764-.317a10.935 10.935 0 0 0-3.556-4.981l-.023-.018C17.171.732 15.182-.001 13.022-.001S8.873.732 7.29 1.963l.021-.016a10.639 10.639 0 0 0-3.022 3.697l-.028.062c-.17.329-.35.739-.506 1.159l-.027.082a9.182 9.182 0 0 0-.764.317a4.719 4.719 0 0 0-.892.536l.011-.008l-.049.039l-.019.069q-.038.133-.064.269a6.69 6.69 0 0 0-.083.642h.004q-.006.073-.01.148a6.776 6.776 0 0 0 .238 2.144l-.011-.048q.073.258.171.507c.032.081.068.167.107.251c.019.042.041.081.062.125l.035.058c.014.023.023.04.043.069l.035.055l.056.029c.06.031.109.053.161.081l.156.071l.087.036A13.112 13.112 0 0 0 .874 14.86l-.031.051l-.846 1.336l1.189-1.043a19.53 19.53 0 0 1 2.587-1.808l.093-.05l.023.191l.06.484l.305 2.474a.167.167 0 0 0 .07.116h.001l.624.439l.031.106c.015.049.028.098.044.148l.143.43a9.775 9.775 0 0 0 1.311 2.648l-.021-.031l-.081.023c-.323-.15-1.459-.676-1.839-.848l-.797-.357l.472.735c.034.053.847 1.321 1.51 2.177a3.643 3.643 0 0 0 3.023 1.613h.056h-.003a5.92 5.92 0 0 0 .629-.036l-.026.002c.479-.056.838-.098 1.111-.134a7.39 7.39 0 0 0 2.637.477c.958 0 1.874-.179 2.716-.505l-.051.018c.51-.196.947-.416 1.358-.674l-.032.019l.023-.006c.656-.161 1.399-.335 1.767-.439c.059-.016.132-.034.209-.052a2.426 2.426 0 0 0 1.816-1.216l.006-.012c.459-.921 1.209-2.33 1.216-2.344l.351-.656l-.681.297c-.037.016-.901.393-1.351.628c-.015.009-.035.017-.052.026c.224-.44.433-.959.595-1.498l.018-.068l.247-1.017l.71-1.588s.199-1.735.203-1.734l.069-.664c.062-.027.441-.052.515-.081c.106-.042.211-.081.317-.128l.155-.068c.055-.026.104-.048.161-.081l.056-.029l.035-.055a.683.683 0 0 0 .05-.074l.002-.003l.034-.062c.02-.034.041-.075.06-.117l.003-.008q.058-.123.109-.251q.096-.25.169-.507a6.565 6.565 0 0 0 .225-2.113l.001.016zM7.986 13.241a15.688 15.688 0 0 1 1.595-.894l.098-.043q1.702-.129 3.41-.129c.298 0 .591.004.88.009a5.47 5.47 0 0 1-.611.434l-.025.014l-.735.445a6.327 6.327 0 0 0-.931.189l.044-.011a.115.115 0 0 0-.039.017h.001a.163.163 0 0 0-.075.094v.001l-.751 2.476l-4.015.893l-1.265-.887l-.368-2.979a40.977 40.977 0 0 1 2.831-.414q-.189.224-.357.466l-.685.986zm5.416 10.366a6.038 6.038 0 0 1-2.035-.224l.043.011c.055-.011.095-.017.144-.023a2.756 2.756 0 0 0 1.763-.948l.003-.004h.182c.087 0 .167-.007.252-.011c.223-.017.425-.046.623-.086l-.033.006l.062.062c.371.375.873.62 1.432.662h.007c-.717.3-1.548.502-2.418.563l-.025.001zm.059-1.451c.099-.222.175-.395.231-.525q.212.237.433.466a6.837 6.837 0 0 1-.385.047c-.081 0-.17.013-.249.013h-.03zm-1.969-5.111l.359-.081a.17.17 0 0 0 .123-.114v-.001l.75-2.465a2.273 2.273 0 0 1 .694.002l-.013-.002l.772 2.538a.168.168 0 0 0 .14.119h.001l.36.04l4.928.564l.018.001a.17.17 0 0 0 .097-.031l.495-.346l.627-.442a8.197 8.197 0 0 1-.09.351q-.07.223-.148.443a11.551 11.551 0 0 1-.707 1.635l.031-.065q-.108.007-.216.007c-.242 0-.484-.019-.752-.044a85.527 85.527 0 0 1-1.443-.155l-1.469-1.164c-.022-.017-.041-.035-.061-.052a.615.615 0 0 0-.439-.182h-.01a5.33 5.33 0 0 0-1.307.315l.037-.012a.91.91 0 0 0 .392-.661v-.004a2.774 2.774 0 0 1-1.437.628l-.014.001a2.332 2.332 0 0 1-1.599-.439l.007.005a.947.947 0 0 0 .273.494a6.403 6.403 0 0 0-.763-.047h-.001c-.097 0-.196 0-.29.007c-.549.029-1.322.858-1.863 1.693c-.266.081-1.1.349-2.063.632a10.843 10.843 0 0 1-1.146-2.025l-.027-.071a7.837 7.837 0 0 1-.267-.705l.231.161l.572.403l.191.133c.027.019.061.03.097.031l.019.001l.019-.001h-.001l4.892-1.09zm3.041-3.678a.17.17 0 0 0-.114-.113h-.001a5.738 5.738 0 0 0-1.024-.192l-.025-.002l.097-.007c1.17-.16 2.222-.42 3.227-.779l-.111.035c1.655.126 3.143.321 4.605.59l-.263-.04l-.361 2.534l-1.275.892l-4.008-.456zm8.799-3.414c-.044.323-.193.962-.193.962l-.115.019c-.06-.017-.624-.161-1.595-.343l.09-.02c.234-.057.431-.121.622-.199l-.03.011c.113-.047.205-.092.294-.143l-.014.007a.954.954 0 0 0 .241-.187c-.968.318-2.983.173-4.766-.011a46.912 46.912 0 0 0-4.901-.281h-.017a49.7 49.7 0 0 0-5.135.301l.221-.02c-1.785.185-3.79.329-4.765.011a.99.99 0 0 0 .237.185l.005.003c.078.043.174.089.272.129l.019.007c.161.067.358.132.561.183l.031.007c.071.016.143.032.215.045c-.887.168-1.398.302-1.456.317c0 0-.204.033-.249-.123a5.259 5.259 0 0 1-.189-.829l-.004-.031a6.994 6.994 0 0 1-.061-.938v-.044v.002v-.161c0-.11.008-.22.019-.33c.146-.206.335-.371.556-.485l.009-.004c.229-.136.514-.284.805-.417l.061-.025a10.817 10.817 0 0 1 1.755-.437l.067-.009C6.497 7 8.551 6.728 8.705 6.67s1.039-.743 1.194-.8c-.323.054-4.838.743-5.039.806C6.324 2.848 9.197.983 13.022.983s6.697 1.76 8.162 5.589c-1.37-.331-3.052-.615-4.766-.792l-.173-.014c.156.057 1.034.675 1.197.696a36.03 36.03 0 0 1 3.852.752l-.252-.054c.303.123.607.249.901.388c.352.158.636.306.912.467l-.046-.025c.229.119.418.284.561.485l.003.005c.015.161.02.327.022.491v.036c0 .333-.023.661-.066.982l.004-.037z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Treehouse(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M19.517 4.338c-.797-.448-2.069.275-2.84 1.617L15.276 8.39a3.024 3.024 0 0 0-.338 1.399c0 .653.204 1.259.552 1.756l-.007-.01l.041.061a18.593 18.593 0 0 0 1.803 2.095a1.592 1.592 0 1 1-2.732.889l-.001.008a1.927 1.927 0 0 0-.567-1.472c-.5-.557-1.523.495-1.854 1.592l-.016.06a7.933 7.933 0 0 0-.474 2.087l-.002.032c.037.059.071.119.099.18a1.792 1.792 0 1 1-2.431-.733l.009-.005l.12-.06c.183-.294.327-.634.414-.997l.005-.024l.297-.976a.922.922 0 0 0 .049-.156l.001-.006l.89-3.155l-.947 1.939c-.111-.359-.297-.376-.701-.09c-.24.181-.646.525-.841.662a3.486 3.486 0 0 0-.833 1.077l-.009.02c-.095.2-.233.366-.402.492l-.003.002a1.555 1.555 0 0 1-2.207-.221l-.002-.003a1.574 1.574 0 0 1 .892-2.523l.01-.002a12.972 12.972 0 0 0 2.604-1.527l-.033.024c.141-.092.264-.19.377-.298l-.001.001l2.176-1.652l-1.897 1.013a2.204 2.204 0 0 0-.433.018l.011-.001a3.76 3.76 0 0 0-1.57.356l.023-.01a.999.999 0 0 1-.283.267l-.004.002a1.406 1.406 0 1 1-1.64-2.284l.005-.003c.239-.168.535-.269.855-.271c.559.089 1.204.14 1.861.14c.387 0 .77-.018 1.148-.052l-.049.004l.285-.045a4.798 4.798 0 0 0 2.979-2.143l.012-.02l1.171-2.044c.646-1.127.601-2.358-.099-2.77L12.275.262c-.34-.17-.74-.269-1.164-.269s-.824.099-1.18.276l.015-.007L1.18 5.178A2.467 2.467 0 0 0-.007 7.155v9.644a2.545 2.545 0 0 0 1.155 1.992l.01.006l8.754 4.929c.34.171.74.271 1.164.271s.825-.1 1.179-.278l-.015.007l8.73-4.929a2.544 2.544 0 0 0 1.171-1.991V7.145a2.56 2.56 0 0 0-1.147-1.992l-.01-.006l-1.444-.814z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Trello(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.003 0h.04c.814 0 1.55.337 2.075.88l.001.001a2.88 2.88 0 0 1 .88 2.076v.047v-.002v18.042c0 .814-.337 1.55-.88 2.075l-.001.001a2.88 2.88 0 0 1-2.076.88h-.043h.002H2.955c-.814 0-1.55-.337-2.075-.88l-.001-.001a2.88 2.88 0 0 1-.88-2.076V21v.002V2.96c0-.814.337-1.55.88-2.075L.88.884a2.88 2.88 0 0 1 2.076-.88h.043h-.002zM10.448 18.162V4.537c0-.395-.164-.751-.428-1.004a1.386 1.386 0 0 0-.998-.423H4.543c-.395 0-.751.164-1.004.428a1.39 1.39 0 0 0-.429 1.004v13.619l-.001.05c0 .392.165.745.429.994l.001.001c.258.248.609.4.996.4h.023h-.001h4.468c.387 0 .738-.153.997-.401c.265-.25.43-.603.43-.995l-.001-.052v.003zm10.448-6V4.538c0-.395-.164-.751-.428-1.004a1.386 1.386 0 0 0-1.004-.429h-4.473c-.395 0-.751.164-1.004.428c-.265.254-.429.61-.429 1.004v.014v-.001v7.608l-.001.05c0 .392.165.745.429.994l.001.001c.258.248.609.4.996.4h.023h-.001h4.468c.387 0 .738-.153.997-.401c.265-.25.43-.603.43-.995l-.001-.048v.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tripadvisor(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.781 12.894v.016c0 .266-.11.505-.286.677a.928.928 0 0 1-.672.286h-.011h.001h-.002a.978.978 0 0 1-.978-.978v-.012c0-.264.11-.502.286-.672a.942.942 0 0 1 .677-.287h.017h-.001a.979.979 0 0 1 .969.968v.001zm12.019-.013a.979.979 0 1 1-1.958 0v-.017c0-.263.11-.501.286-.67a.99.99 0 0 1 1.386.001a.924.924 0 0 1 .287.671v.017v-.001zm-10.832.01a2.009 2.009 0 1 0-4.017.001a2.009 2.009 0 0 0 4.017 0zm12.009-.01a2.009 2.009 0 1 0-.589 1.422c.363-.353.589-.846.589-1.391zv.002zm-11.125.01a2.901 2.901 0 0 1-5.802 0v-.041c0-.788.328-1.499.854-2.005l.001-.001A2.8 2.8 0 0 1 5.912 10h.048h-.002h.041A2.78 2.78 0 0 1 8 10.846l.001.001c.524.508.85 1.218.85 2.005v.044v-.002zm12.02-.01v.004a2.901 2.901 0 1 1-.849-2.051c.524.508.849 1.218.849 2.004v.045zm-10.041.031l.001-.07a4.625 4.625 0 0 0-1.41-3.33l-.001-.001A4.623 4.623 0 0 0 6.09 8.099l-.074.001h.004h-.032c-.875 0-1.694.24-2.394.658l.021-.012c-1.443.85-2.397 2.397-2.397 4.166s.953 3.316 2.374 4.154l.022.012a4.63 4.63 0 0 0 2.373.646h.036h-.002l.07.001a4.625 4.625 0 0 0 3.33-1.41l.001-.001a4.623 4.623 0 0 0 1.412-3.331l-.001-.073v.004zm6.96-5.969c-1.66-.729-3.594-1.154-5.628-1.154l-.172.001H12a14.836 14.836 0 0 0-6.065 1.182l.097-.037h.017c.832 0 1.624.175 2.34.489l-.037-.015c.739.315 1.37.746 1.901 1.276a6.073 6.073 0 0 1 1.255 1.867l.015.039c.3.679.474 1.471.474 2.303v.018v-.001c0-1.619.641-3.088 1.682-4.168l-.002.002a6.107 6.107 0 0 1 1.811-1.264l.038-.016a5.652 5.652 0 0 1 2.252-.521h.007zm4.99 5.968a4.807 4.807 0 1 0-1.405 3.402a4.626 4.626 0 0 0 1.406-3.327l-.001-.079zm-2.77-5.895H24a5.528 5.528 0 0 0-.766 1.16l-.014.032a4.47 4.47 0 0 0-.412 1.163l-.005.03a5.923 5.923 0 0 1 1.144 3.515a5.97 5.97 0 0 1-10.594 3.776l-.008-.011c-.436.541-.879 1.15-1.291 1.781l-.054.087a8.695 8.695 0 0 0-.576-.882l.016.023q-.442-.63-.776-1.015a5.97 5.97 0 0 1-9.452-7.288l-.012.018a4.57 4.57 0 0 0-.429-1.218l.012.025a5.54 5.54 0 0 0-.789-1.2L0 7.019h3.802a12.375 12.375 0 0 1 3.61-1.611l.088-.02a16.569 16.569 0 0 1 4.425-.588h.079H12h.107c1.523 0 2.995.213 4.39.611l-.113-.027c1.378.383 2.584.942 3.672 1.665l-.047-.029z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Troy(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M33.6 24H2.4A2.4 2.4 0 0 1 0 21.6V2.4A2.4 2.4 0 0 1 2.4 0h31.2A2.4 2.4 0 0 1 36 2.4v19.2a2.4 2.4 0 0 1-2.4 2.4M24.911 8.679a7.332 7.332 0 0 0-.865.024l.027-.002c-.01.03.542 1.996 1.066 3.83l.064.224c1.114 3.896 1.114 3.896.937 4.274a1.67 1.67 0 0 1-.686.729l-.008.004l-.231.116l-.994.019c-.96.02-.998.024-1.12.111c-.228.164-.315.425-.489 1.467c-.09.55-.16.982-.16 1.006a2.441 2.441 0 0 0 .741.037l-.01.001c.214 0 .48 0 .812-.006a6.352 6.352 0 0 0 1.975-.132l-.043.008a4.621 4.621 0 0 0 2.55-1.852l.01-.016c.197-.286 5.257-9.732 5.257-9.814a3.979 3.979 0 0 0-.822-.029l.012-.001h-.594l-1.4.011l-.266.132a1.33 1.33 0 0 0-.385.274c-.067.08-.528 1.088-1.12 2.445a18.487 18.487 0 0 1-1.083 2.33l.049-.096c-.022-.046-.218-1.266-.378-2.282c-.187-1.218-.366-2.27-.4-2.346a.675.675 0 0 0-.349-.372l-.004-.002c-.151-.08-.223-.08-1.539-.095h-.553zm-3.77.131c-.043 0-.052.027-.062.071c-.027.123-.418 2.354-.418 2.386a.543.543 0 0 0 .148.117l.003.001c.41.281.69.725.746 1.237l.001.008a1.692 1.692 0 0 1-1.261 1.771l-.012.003c-.136.032-.19.067-.203.131c-.035.168-.418 2.357-.418 2.39c0 .006.023.015.179.015c.07 0 .16 0 .25-.007a4.205 4.205 0 0 0 3.9-3.417l.004-.026a5.42 5.42 0 0 0-.041-1.607l.005.032a4.177 4.177 0 0 0-2.423-2.964l-.027-.01a1.45 1.45 0 0 0-.364-.129l-.01-.002zM4.175 8.674c-.167 0-.603 0-.612.008s-.025.13-.058.32l-.137.758c-.104.588-.179 1.074-.167 1.082s.32.012.621.012h.596l-.012.091c0 .026-.037.211-.085.489l-.185 1.058a7.72 7.72 0 0 0-.258 2.515l-.001-.021c.182 1.082 1.114 1.766 2.624 1.925c.198.021.466.031.701.031a1.791 1.791 0 0 0 .528-.042l-.012.002c.022-.027.413-2.182.418-2.306c0-.052-.069-.068-.386-.088c-.778-.043-1.126-.297-1.126-.823c0-.16.367-2.381.457-2.763c.013-.059.032-.075.433-.075h.606a3.931 3.931 0 0 0 .691-.03l-.019.002c.025-.042.378-2 .378-2.099c0-.037-.198-.047-.847-.047h-.846l.107-.609c.195-1.063.149-1.32-.278-1.527c-.214-.107-.231-.107-1.152-.123l-.953-.012l-.024.111c-.012.064-.096.525-.183 1.03s-.171.96-.183 1.022l-.024.112zm6-.008l-.025.111c-.04.186-1.415 8.014-1.415 8.053a11.432 11.432 0 0 0 1.387.035l-.019.001h1.369l.04-.21c.025-.111.16-.871.302-1.686c.14-.8.297-1.6.342-1.75a2.611 2.611 0 0 1 1.727-1.805l.018-.005a2.318 2.318 0 0 1 .833-.09h-.007c.499 0 .53-.005.545-.08c.045-.195.452-2.57.445-2.593a.73.73 0 0 0-.292-.031h.003h-.021c-.16 0-.317.013-.47.038l.017-.002a3.452 3.452 0 0 0-1.603.813l.003-.003c-.292.27-.546.576-.756.912l-.011.019a.418.418 0 0 1-.094.144c.015-.157.037-.297.066-.435l-.004.024c.166-.885.076-1.192-.4-1.371a5.292 5.292 0 0 0-1.067-.071h.008zm9.704-.026h-.141c-.236 0-.467.022-.691.064l.023-.004a4.13 4.13 0 0 0-2.869 2.195l-.011.024a3.476 3.476 0 0 0-.429 1.84v-.007l-.002.139c0 .262.024.518.069.767l-.004-.026a4.183 4.183 0 0 0 1.879 2.674l.018.01c.276.177.595.325.933.427l.027.007c.025-.018.139-.633.247-1.233l.218-1.213l-.103-.08a1.981 1.981 0 0 1-.635-.721l-.005-.011a1.093 1.093 0 0 1-.147-.722l-.001.006a1.15 1.15 0 0 1 .131-.698l-.003.006c.228-.47.651-.815 1.161-.931l.011-.002c.08-.008.151-.031.151-.052c0-.054.4-2.314.422-2.394c-.015-.056-.07-.064-.249-.064z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Truck(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.449 0c-3.64 0-6.71.201-7.168.468c-.8.273-1.402 2.09-1.404 4.246v1.071h-.218a.533.533 0 0 0-.088-.135a.596.596 0 0 0 .146-.392v-.544a.49.49 0 0 0-.427-.535H.432a.492.492 0 0 0-.428.539v-.002v.544c0 .15.055.288.146.393l-.001-.001a.596.596 0 0 0-.146.392V8.73c-.002.015-.002.031-.002.048c0 .25.187.457.429.487h.859A.42.42 0 0 0 1.658 9l.001-.003h.218v4.822c0 .191.205.368.536.463v.626c-.473.061-.8.274-.8.518v5.36a2.962 2.962 0 0 0 .542 1.863l-.006-.009v.921c0 .243.197.44.44.44h1.798a.44.44 0 0 0 .44-.44v-.631h11.247v.631c0 .243.197.44.44.44h1.798a.44.44 0 0 0 .44-.44v-.92a2.955 2.955 0 0 0 .536-1.86v.007v-5.36c0-.244-.331-.457-.8-.518v-.626c.331-.095.535-.272.536-.463v-4.82h.218c.06.152.202.259.369.268h.855a.492.492 0 0 0 .428-.539v.002v-2.686a.599.599 0 0 0-.146-.393l.001.001a.596.596 0 0 0 .146-.392V4.72a.49.49 0 0 0-.427-.535h-.859a.492.492 0 0 0-.428.539v-.002v.544c0 .15.055.288.146.393l-.001-.001a.553.553 0 0 0-.087.131l-.001.003h-.218V4.716c0-2.167-.608-3.992-1.414-4.251c-.49-.266-3.545-.463-7.157-.463zM3.101 3.107h14.695c.259 0 .475.184.525.429l.001.003c.102.51.16 1.097.16 1.698v.038v-.002V9c0 .296-.24.536-.536.536H2.949A.536.536 0 0 1 2.413 9V5.238c0-.601.058-1.188.169-1.756l-.009.057a.536.536 0 0 1 .526-.432zm.651 11.786H4.9c.143 0 .259.112.267.253v.001l.191 3.482v.014c0 .148-.12.268-.268.268H3.752a.268.268 0 0 1-.268-.268v-3.482c0-.148.12-.268.268-.268m12.245 0h1.148c.148 0 .268.12.268.268v3.479c0 .148-.12.268-.268.268h-1.339a.268.268 0 0 1-.268-.268v-.015v.001l.191-3.482c.01-.14.126-.25.267-.25h.001zm-9.968.64h8.84c.074 0 .134.06.134.134v.16c0 .074-.06.134-.134.134h-8.84a.134.134 0 0 1-.134-.134v-.16c0-.074.06-.134.134-.134m.053.964h8.733c.074 0 .134.06.134.134v.16c0 .074-.06.134-.134.134H6.082a.134.134 0 0 1-.134-.134v-.16c0-.074.06-.134.134-.134m.054.964h8.626c.074 0 .134.06.134.134v.16c0 .074-.06.134-.134.134H6.136a.133.133 0 0 1-.133-.128v-.16c0-.074.06-.134.134-.134zm.054.964h8.518c.074 0 .134.06.134.134v.16c0 .074-.06.134-.134.134H6.19a.134.134 0 0 1-.134-.134v-.16c0-.074.06-.134.134-.134"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tty(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M13.09.147c2.033.147 3.892.778 5.497 1.778l-.051-.029a4.555 4.555 0 0 1 2.054 2.535l.009.032l.049.27h-6.03l-.098-.185a2.49 2.49 0 0 0-.439-.495l-.003-.003c-.59-.526-.761-.554-3.605-.56c-1.35 0-2.63.029-2.849.057a2.32 2.32 0 0 0-1.412.842l-.004.005l-.249.342h-5.96l.042-.292c.135-1.002 1.379-2.29 3.001-3.107C4.712.585 6.654.107 8.696.016L8.73.015a39.667 39.667 0 0 1 4.51.146L13.089.15zM5.787 5.899c.014.014.014.611-.007 1.322c-.035 1.251-.042 1.294-.206 1.43s-.313.135-2.688.135H.369l-.17-.178C.021 8.438.021 8.423.021 7.15V5.87h2.866c1.578 0 2.886.014 2.901.029zM20.64 7.108c0 1.228 0 1.244-.185 1.458l-.185.221h-5.049l-.413-.413V5.869h5.831zM3.04 9.809c.427.178.469.313.469 1.415v.988l-.442.413H1.823c-1.768 0-1.798-.021-1.798-1.44c0-1.002.098-1.273.512-1.394a9.55 9.55 0 0 1 2.553.027l-.051-.006zm4.266 0c.427.178.469.313.469 1.415v.988l-.413.413H6.111a6.99 6.99 0 0 1-1.455-.086l.04.005c-.327-.147-.406-.4-.406-1.36c0-1.01.098-1.28.512-1.401a9.574 9.574 0 0 1 2.555.027l-.051-.006zm4.267 0c.427.178.469.313.469 1.415v.988l-.413.413H9.04l-.48-.48v-.917c0-1.44.071-1.5 1.76-1.507a4.017 4.017 0 0 1 1.279.098l-.027-.006zm4.266 0c.427.178.469.313.469 1.415v.988l-.413.413h-2.588l-.48-.48v-.917c0-1.44.071-1.5 1.76-1.507a4.002 4.002 0 0 1 1.277.098l-.027-.006zm4.529.17l.206.227v2.005l-.413.413h-2.589l-.48-.48v-.917c0-.775.021-.946.142-1.138c.192-.32.434-.37 1.76-.356l1.166.014zM7.371 13.626c.362.185.406.32.406 1.36v.982l-.206.227l-.206.234l-1.2.021c-1.792.029-1.87-.035-1.87-1.507v-.91l.48-.48h1.223a6.33 6.33 0 0 1 1.409.086l-.038-.005zm4.316.029c.114.059.209.14.282.238l.002.002a5.616 5.616 0 0 1 .071 1.112v-.009v.967l-.206.227l-.206.235l-1.13.021c-1.84.042-1.941-.035-1.941-1.507v-.906l.48-.48h1.216a4.006 4.006 0 0 1 1.456.114l-.028-.006zm4.267 0c.114.059.209.14.282.238l.002.002a5.616 5.616 0 0 1 .071 1.112v-.009v.967l-.206.227l-.206.235l-1.13.021c-1.84.042-1.941-.035-1.941-1.507v-.906l.48-.48h1.216a4.006 4.006 0 0 1 1.456.114l-.028-.006zM3.305 17.587l.206.227v.982c-.007 1.424-.021 1.44-1.778 1.44H.51l-.48-.48v-.896c0-1.486.08-1.542 1.87-1.52l1.2.021zm4.266 0l.206.227v.967a5.766 5.766 0 0 1-.076 1.136l.005-.034c-.16.292-.419.348-1.714.348H4.776l-.48-.48v-.896c0-1.486.08-1.542 1.87-1.52l1.2.021zm4.267 0l.206.227v.967a5.766 5.766 0 0 1-.076 1.136l.005-.034c-.16.292-.419.348-1.714.348H9.04l-.48-.48v-.896c0-1.486.08-1.542 1.87-1.52l1.2.021zm4.266 0l.206.227v2.005l-.442.413h-2.56l-.48-.48v-.896c0-1.486.08-1.542 1.87-1.52l1.2.021zm4.252-.029l.221.213v2.048l-.442.413h-2.547l-.227-.206l-.234-.206l-.021-.924c-.042-1.514.042-1.586 1.849-1.557l1.188.014zm-3.626 3.698c.298.16.362.384.362 1.344s-.049 1.094-.427 1.294c-.178.092-1.04.106-6.328.106c-6.8 0-6.48.021-6.712-.462a4.227 4.227 0 0 1-.009-1.941l-.006.028c.206-.505-.256-.469 6.676-.469c5.26 0 6.269.015 6.44.101z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tumblr(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m12.634 19.166l1.154 3.418a3.204 3.204 0 0 1-1.578.947l-.022.005a7.765 7.765 0 0 1-2.547.462h-.006a8.354 8.354 0 0 1-2.807-.392l.059.016a6.244 6.244 0 0 1-2.068-1.077l.013.01a6.716 6.716 0 0 1-1.355-1.504l-.016-.025a6.107 6.107 0 0 1-.789-1.687l-.011-.044a6.237 6.237 0 0 1-.24-1.7V9.748H-.001V6.647a6.932 6.932 0 0 0 1.878-1.015l-.017.012a6.51 6.51 0 0 0 1.301-1.282l.012-.016c.313-.425.592-.908.817-1.42l.019-.049a7.99 7.99 0 0 0 .478-1.37l.012-.058c.085-.35.161-.785.211-1.229l.005-.051A.227.227 0 0 1 4.78.047a.153.153 0 0 1 .108-.05h3.52v6.115h4.8v3.634H8.391v7.495c0 .278.034.548.099.806l-.005-.023c.076.287.188.539.332.768l-.007-.012c.171.267.415.474.703.595l.01.004a2.675 2.675 0 0 0 1.182.201h-.007a4.704 4.704 0 0 0 1.964-.43l-.029.012z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Tv(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M27.75 0H2.25A2.25 2.25 0 0 0 0 2.25v15a2.25 2.25 0 0 0 2.25 2.25h11.489v1.499H6.24a1.5 1.5 0 0 0 0 3h18a1.5 1.5 0 0 0 0-3h-7.501v-1.5H27.75a2.25 2.25 0 0 0 2.25-2.25v-15a2.25 2.25 0 0 0-2.25-2.25zM27 16.5H3V3h24z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Twitch(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.481 6.259v6.259H9.39V6.259zm5.74 0v6.259H15.13V6.259zm0 10.962l3.649-3.663V2.091H3.649v15.13h4.702v3.13l3.13-3.13zM22.961 0v14.61l-6.259 6.259H12l-3.13 3.13H5.74V20.87H0V4.168L1.572 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Twitter(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M29.55 2.85a12.813 12.813 0 0 1-3.004 3.106l-.036.025q.018.262.018.787a17.203 17.203 0 0 1-.745 4.987l.032-.122a17.56 17.56 0 0 1-2.206 4.724l.04-.065a18.49 18.49 0 0 1-3.435 3.927l-.024.02a15.333 15.333 0 0 1-4.73 2.704l-.108.033c-1.765.648-3.803 1.022-5.928 1.022H9.29h.007h-.127c-3.41 0-6.584-1.015-9.234-2.76l.063.039c.419.048.904.075 1.396.075h.07h-.004l.126.001c2.807 0 5.386-.975 7.417-2.606l-.023.018a6.073 6.073 0 0 1-5.65-4.157l-.012-.043c.342.057.738.091 1.141.094h.003a6.26 6.26 0 0 0 1.637-.216l-.044.01a5.98 5.98 0 0 1-3.47-2.08l-.008-.011a5.816 5.816 0 0 1-1.379-3.773l.001-.084v.004v-.075a5.922 5.922 0 0 0 2.727.768h.011a6.094 6.094 0 0 1-1.953-2.129l-.016-.031a6 6 0 0 1-.731-2.889c0-1.126.306-2.18.84-3.084l-.015.028a17.29 17.29 0 0 0 5.425 4.427l.095.045c2.022 1.067 4.402 1.743 6.927 1.864l.038.001a6.548 6.548 0 0 1-.149-1.382v-.001a6.062 6.062 0 0 1 10.477-4.147l.003.003A11.857 11.857 0 0 0 28.772.415l-.053.03a5.913 5.913 0 0 1-2.635 3.323l-.028.015a12.045 12.045 0 0 0 3.569-.967l-.077.031z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Twoo(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.19 8.8c.054.01.106.023.16.03a2.81 2.81 0 0 1 1.446.582l-.006-.005a3.115 3.115 0 0 1 1.136 1.739l.004.021a2.933 2.933 0 0 1 .085 1.016l.001-.011a3.154 3.154 0 0 1-1.756 2.671l-.018.008c-.272.14-.589.235-.925.267l-.011.001a2.839 2.839 0 0 1-2.265-.729l.002.002c-.212-.185-.4-.405-.59-.61c-.314-.332-.626-.666-.939-.999c-.494-.526-.99-1.05-1.481-1.576a1.066 1.066 0 0 0-.672-.361l-.005-.001a1.056 1.056 0 0 0-1.063.537l-.003.005a1.176 1.176 0 0 0 .212 1.479l.001.001c.161.152.373.251.608.27h.004a1.114 1.114 0 0 0 1.175-.902l.001-.007c.021-.102.024-.21.037-.326l.112.118l1.36 1.365c.08.08.08.08.023.174a2.987 2.987 0 0 1-1.941 1.503l-.02.004a2.826 2.826 0 0 1-2.756-.798l-.001-.001a3.032 3.032 0 0 1-.758-1.259l-.006-.021a.24.24 0 0 1 .007-.174l-.001.002q.425-1.04.847-2.08c.067-.149.139-.339.199-.534l.01-.039a1.037 1.037 0 0 0-.007-.586l.002.007a.118.118 0 0 1 .052-.152h.001c.387-.302.862-.51 1.38-.581l.015-.002a2.812 2.812 0 0 1 2.274.667l-.003-.003c.243.222.468.45.68.69l.008.009q1.194 1.264 2.383 2.534c.16.185.379.315.628.361l.007.001a1.051 1.051 0 0 0 .959-.331l.001-.001c.195-.193.317-.459.324-.753v-.001a1.125 1.125 0 0 0-.926-1.178l-.007-.001a1.092 1.092 0 0 0-1.226.835l-.001.007c-.028.122-.034.25-.051.382l-.052-.05q-.708-.72-1.414-1.446c-.08-.08-.08-.08-.026-.174a3.014 3.014 0 0 1 1.877-1.489l.021-.005c.165-.043.338-.058.506-.084a.23.23 0 0 0 .061-.019l-.001.001zM7.757 11.326c.018-.033.03-.052.039-.073c.058-.137.118-.273.17-.412a.193.193 0 0 0 0-.121v.001c-.058-.154-.125-.305-.186-.465a1.028 1.028 0 0 1 .801-1.399l.006-.001a1.004 1.004 0 0 1 1.084.592l.003.006c.24.56.462 1.132.693 1.699l.068.154h.018c.013-.026.028-.05.038-.075c.23-.56.461-1.125.686-1.689a1.033 1.033 0 0 1 1.966.041l.002.007a1.006 1.006 0 0 1-.043.686l.003-.007q-.853 2.106-1.709 4.206a1.024 1.024 0 0 1-1.016.632h.003a.998.998 0 0 1-.886-.633l-.002-.007c-.116-.27-.223-.546-.334-.819c-.015-.036-.033-.07-.058-.122a.582.582 0 0 0-.037.075l-.002.005c-.115.283-.229.568-.345.85a1.014 1.014 0 0 1-1.023.649h.003a.985.985 0 0 1-.889-.633l-.002-.007c-.272-.651-.532-1.306-.8-1.959l-.527-1.299a.429.429 0 0 0-.257-.285l-.003-.001a.527.527 0 0 0-.13-.022h-.735c-.105 0-.105 0-.105.109v3.095a1.03 1.03 0 0 1-2.035.183l-.001-.007a1.075 1.075 0 0 1-.019-.187v-3.096c0-.094 0-.096-.094-.096H1.03a1.026 1.026 0 0 1-.126-2.045h.005a1.06 1.06 0 0 1 .188-.007h-.003c1.64 0 3.28.005 4.921 0h.011c.453 0 .84.28.999.676l.003.007c.221.57.458 1.134.688 1.7a.91.91 0 0 0 .048.094l-.002-.004z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Uber(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.178.001h.023c.495 0 .943.205 1.262.536v.001c.331.319.536.767.536 1.262v.024v-.001v20.38c0 .495-.205.943-.536 1.262h-.001a1.747 1.747 0 0 1-1.262.536h-.025h.001H1.797c-.495 0-.943-.205-1.262-.536v-.001a1.747 1.747 0 0 1-.536-1.262v-.025v.001v-20.38c0-.495.205-.943.536-1.262h.001A1.747 1.747 0 0 1 1.798 0h.025h-.001zM12.75 19.232a7.339 7.339 0 0 0 3.53-1.354l-.021.015a7.148 7.148 0 0 0 2.363-2.848l.018-.044c.401-.85.636-1.847.636-2.899c0-.281-.017-.558-.049-.831l.003.033a6.944 6.944 0 0 0-1.194-3.399l.015.024a7.297 7.297 0 0 0-2.586-2.338l-.039-.019a7.088 7.088 0 0 0-3.402-.858h-.025H12h-.028a7.135 7.135 0 0 0-3.438.876l.037-.019a7.317 7.317 0 0 0-2.609 2.332l-.016.025a6.831 6.831 0 0 0-1.177 3.297l-.002.026h5.036V10.17a.363.363 0 0 1 .365-.365h.011h-.001h3.653a.363.363 0 0 1 .365.365v.011v-.001v3.653a.363.363 0 0 1-.365.365h-.011h.001h-3.654a.363.363 0 0 1-.365-.365v-.011v.001v-1.072H4.767a7.339 7.339 0 0 0 1.354 3.53l-.015-.021a7.147 7.147 0 0 0 2.848 2.366l.044.018a6.963 6.963 0 0 0 2.932.637c.288 0 .573-.017.852-.051l-.034.003z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Ubuntu(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12.004 0h-.081c-1.661 0-3.242.345-4.674.968l.076-.029A12.12 12.12 0 0 0 3.51 3.507A12.112 12.112 0 0 0 .972 7.244l-.03.078a11.545 11.545 0 0 0-.939 4.598v.085v-.004v.081c0 1.661.345 3.242.968 4.674l-.029-.076a12.124 12.124 0 0 0 2.565 3.815a12.095 12.095 0 0 0 3.737 2.538l.078.03c1.356.593 2.937.939 4.598.939h.089h-.005h.081c1.661 0 3.242-.345 4.674-.968l-.076.029a12.12 12.12 0 0 0 3.815-2.568a12.095 12.095 0 0 0 2.538-3.737l.03-.078c.593-1.356.938-2.936.938-4.596v-.164c0-1.661-.345-3.242-.968-4.674l.029.076a12.12 12.12 0 0 0-2.568-3.815A12.112 12.112 0 0 0 16.76.969l-.078-.03a11.539 11.539 0 0 0-4.596-.938h-.088h.004zm2.691 4.123c.203-.366.543-.634.949-.738l.011-.002a1.555 1.555 0 0 1 1.241.177l-.007-.004c.366.203.634.543.738.949l.002.011a1.48 1.48 0 0 1-.177 1.215l.004-.006a1.516 1.516 0 0 1-.949.738l-.01.002a1.545 1.545 0 0 1-1.217-.152l.007.004a1.516 1.516 0 0 1-.738-.949l-.002-.011a1.607 1.607 0 0 1 .153-1.242zm-2.691.913h.017c.329 0 .653.022.97.066l-.037-.004c.347.047.652.112.949.196l-.047-.012c.053.346.179.655.363.921l-.005-.007c.196.285.448.517.742.685l.011.006c.274.156.598.258.943.284h.007a2.169 2.169 0 0 0 .965-.141l-.015.005a7.038 7.038 0 0 1 1.426 1.983l.018.042c.355.728.582 1.579.629 2.477l.001.016l-2.272.025a4.636 4.636 0 0 0-1.502-3.021l-.004-.003a4.455 4.455 0 0 0-3.098-1.248h-.066h.003H12c-.363 0-.716.045-1.053.13l.03-.006a5.56 5.56 0 0 0-.986.334l.035-.014l-1.111-2a6.62 6.62 0 0 1 1.422-.522l.046-.009a6.817 6.817 0 0 1 1.58-.182h.04h-.002zm-8.151 8.569h-.024c-.436 0-.83-.181-1.111-.471a1.545 1.545 0 0 1-.472-1.113v-.05c0-.435.18-.828.469-1.109a1.545 1.545 0 0 1 1.11-.469h.028h-.001h.017c.431 0 .821.18 1.097.468l.001.001c.289.281.469.675.469 1.11V12v-.001v.026c0 .435-.18.828-.469 1.109a1.514 1.514 0 0 1-1.1.471h-.012h.001zm1.407.148c.246-.215.449-.473.598-.763l.007-.014c.148-.28.234-.613.234-.965v-.019c0-.353-.087-.685-.24-.977l.006.012a2.6 2.6 0 0 0-.602-.775l-.003-.003a6.698 6.698 0 0 1 1.087-2.315l-.013.019a7.285 7.285 0 0 1 1.79-1.737l.025-.016l1.16 1.975a5.02 5.02 0 0 0-1.432 1.616l-.013.025a4.376 4.376 0 0 0-.555 2.146v.041v-.002v.04c0 .786.201 1.525.555 2.168l-.012-.023a4.787 4.787 0 0 0 1.442 1.632l.015.01l-1.161 1.976a7.28 7.28 0 0 1-1.801-1.732l-.014-.02a6.606 6.606 0 0 1-1.063-2.25l-.01-.046zm11.628 6.69a1.523 1.523 0 0 1-1.233.17l.011.003a1.609 1.609 0 0 1-.971-.733l-.004-.007a1.638 1.638 0 0 1-.146-1.246l-.003.011a1.52 1.52 0 0 1 .733-.956l.008-.004c.223-.132.491-.209.778-.209c.154 0 .302.022.443.064l-.011-.003c.417.107.757.375.956.733l.004.008a1.475 1.475 0 0 1 .17 1.22l.003-.011a1.52 1.52 0 0 1-.73.956zm-.024-3.481a2.114 2.114 0 0 0-.958-.135l.008-.001a2.273 2.273 0 0 0-.961.29l.011-.006a2.276 2.276 0 0 0-.748.684l-.005.008a2.167 2.167 0 0 0-.357.901l-.002.012a6.169 6.169 0 0 1-.866.181l-.035.004a6.91 6.91 0 0 1-.933.062H12h.001h-.024a6.935 6.935 0 0 1-1.641-.195l.048.01a6.757 6.757 0 0 1-1.507-.549l.039.018l1.111-2c.267.121.584.231.912.312l.038.008c.307.078.66.123 1.023.123h.064a4.46 4.46 0 0 0 3.1-1.249l-.001.001a4.626 4.626 0 0 0 1.505-3.01l.001-.018l2.272.025a6.527 6.527 0 0 1-.647 2.532l.017-.039a7.071 7.071 0 0 1-1.445 2.029l-.001.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Udacity(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.103 0L0 4.595v10.579A8.853 8.853 0 0 0 8.852 24h.026h-.001h.031c1.63 0 3.159-.433 4.478-1.191l-.044.023l6.562-3.729c2.604-1.337 4.358-3.998 4.374-7.07V.972L22.669.041l-6.575 3.673v11.57c0 .491-.057.97-.163 1.429l.008-.042a8.368 8.368 0 0 1-.602 1.735l.022-.05a6.333 6.333 0 0 1-2.443-.835l.029.016a7.763 7.763 0 0 1-1.384-1.037l.004.004a5.893 5.893 0 0 1-.776-.935l-.013-.021a5.402 5.402 0 0 1-.579-1.055l-.013-.036a6.6 6.6 0 0 1-.37-1.152l-.009-.047a5.981 5.981 0 0 1-.101-1.299v.008L9.712.922L8.104.003zm3.733 21.833a7.212 7.212 0 0 1-1.316.406l-.049.009a6.894 6.894 0 0 1-1.416.143a7.351 7.351 0 0 1-1.517-.161l.048.009a7.118 7.118 0 0 1-1.412-.44l.046.018a7.544 7.544 0 0 1-1.26-.684l.027.017a7.183 7.183 0 0 1-1.948-1.933l-.016-.025a7.38 7.38 0 0 1-.648-1.18l-.019-.049a7.08 7.08 0 0 1-.412-1.302l-.009-.049a7.523 7.523 0 0 1-.101-1.466v.01v-9.651l6.07-3.663v10.127a8.14 8.14 0 0 0 6.472 7.828l.053.009a5.648 5.648 0 0 1-.4.49l.004-.005a7.035 7.035 0 0 1-1.052.871l-.025.016a5.752 5.752 0 0 1-1.042.636l-.035.015zm10.52-8.505c-.099.459-.23.86-.396 1.242l.016-.042a7.412 7.412 0 0 1-1.752 2.314l-.008.007a6.12 6.12 0 0 1-2.894 1.45l-.039.007a8.507 8.507 0 0 0 .618-3.067V4.641l4.556-2.757v10.125l.002.161c0 .408-.038.807-.109 1.194l.006-.04z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Umbrella(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M17.852 11.059a3.236 3.236 0 0 0-2.431 1.1l-.003.004a3.233 3.233 0 0 0-4.865-.005l-.003.004a3.24 3.24 0 0 0-4.865-.004l-.003.004a3.24 3.24 0 0 0-2.434-1.103h-.005a3.21 3.21 0 0 0-1.614.432l.016-.008c.276-4.705 4.156-8.417 8.904-8.423h.001a8.93 8.93 0 0 1 8.903 8.399l.001.024a3.192 3.192 0 0 0-1.598-.423h-.005zm-6.484-9.597V.816a.817.817 0 0 0-1.634 0v.646C4.273 1.904.007 6.442-.001 11.977v2.333a.817.817 0 0 0 1.634 0a1.618 1.618 0 0 1 3.234 0a.817.817 0 0 0 1.634 0a1.618 1.618 0 0 1 3.234 0v6.439a3.254 3.254 0 0 0 3.25 3.25a3.254 3.254 0 0 0 3.25-3.25a.817.817 0 0 0-1.634 0a1.618 1.618 0 0 1-3.233 0V14.31a1.618 1.618 0 0 1 3.234 0a.817.817 0 0 0 1.634 0a1.618 1.618 0 0 1 3.234 0a.817.817 0 0 0 1.634 0v-2.332c-.008-5.536-4.273-10.073-9.696-10.514l-.038-.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Underline(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14 0a1.2 1.2 0 0 0-1.2 1.2v10.399a5.2 5.2 0 1 1-10.4 0V1.2A1.202 1.202 0 0 0 0 1.198v10.401a7.6 7.6 0 0 0 15.2 0V1.2A1.2 1.2 0 0 0 14.001 0zm0 21.6H1.2a1.202 1.202 0 0 0-.002 2.4H14a1.202 1.202 0 0 0 .002-2.4z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Undo(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.853 5.286l.791-3.978a1.094 1.094 0 0 0-1.688-1.12l.004-.002l-5.359 3.581a.727.727 0 0 0-.071.053l-.015.011q-.041.034-.078.071l-.027.025l-.05.055l-.018.024c-.015.021-.032.04-.046.062l-.01.016a.95.95 0 0 0-.045.076l-.004.009a1.118 1.118 0 0 0-.041.086v.004a.918.918 0 0 0-.034.096v.009a1.466 1.466 0 0 0-.024.095c0 .01 0 .02-.004.03c0 .025-.009.05-.011.075l-.005.096v.012c0 .036 0 .073.005.109c0 .022.006.044.009.066s0 .027.006.04a1.778 1.778 0 0 0 .026.106c.009.029.02.058.03.086l.005.015l.04.086l.007.014c.015.028.032.055.049.082l.006.01v.004l3.577 5.354a1.094 1.094 0 0 0 1.982-.387l.001-.007l.515-2.586a8.018 8.018 0 1 1-10.84.457A1.094 1.094 0 0 0 2.99 6.576l.001-.001a10.18 10.18 0 0 0-2.992 7.222c0 5.64 4.572 10.212 10.212 10.212s10.212-4.572 10.212-10.212c0-3.537-1.798-6.654-4.53-8.487l-.037-.023z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unity(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.266 12.001l4.2-7.249l2.03 7.253l-2.03 7.25l-4.2-7.25zm-2.047 1.177l4.201 7.254l-7.316-1.876l-5.285-5.378zm4.2-9.608l-4.2 7.253h-8.4l5.285-5.378l7.314-1.875zm6 5.963L20.853 0l-9.564 2.555l-1.416 2.489L7 5.023l-7 6.978l7 6.977l2.871-.022l1.418 2.489l9.564 2.554l2.56-9.531L21.96 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UniversalAcces(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M11.997 1.771C6.349 1.771 1.77 6.349 1.77 11.997s4.578 10.226 10.226 10.226c5.647 0 10.226-4.578 10.226-10.226c0-5.647-4.578-10.225-10.225-10.226m.198 2.252a1.45 1.45 0 1 1 0 2.9a1.45 1.45 0 0 1 0-2.9m5.307 3.668a.567.567 0 0 1-.364.223h-.003l-3.445.53a.107.107 0 0 0-.074.038c-.018.022.343 4.274.343 4.274l1.958 5.337a.772.772 0 0 1-.32 1.038l-.004.002a.625.625 0 0 1-.253.053h-.002a.833.833 0 0 1-.728-.513l-.002-.006s-2.508-5.691-2.522-5.734a.119.119 0 0 0-.112-.081a.11.11 0 0 0-.103.074v.001c-.014.041-2.522 5.734-2.522 5.734a.833.833 0 0 1-.727.518h-.004a.625.625 0 0 1-.256-.054l.004.002a.668.668 0 0 1-.364-.411l-.001-.005a.865.865 0 0 1 .042-.632l-.002.005s1.91-5.165 1.911-5.174l.355-4.363v-.01a.109.109 0 0 0-.092-.107h-.001l-3.36-.52a.561.561 0 1 1 .175-1.108h-.003l3.223.498h3.421a.093.093 0 0 0 .048 0h-.001l3.244-.497a.57.57 0 0 1 .64.464v.003a.547.547 0 0 1-.102.418l.001-.002z"/><path fill="currentColor" d="M12 0C5.373 0 0 5.373 0 12s5.373 12 12 12s12-5.373 12-12S18.627 0 12 0m0 22.975h-.004c-6.064 0-10.979-4.916-10.979-10.979S5.933 1.017 11.996 1.017c6.064 0 10.979 4.916 10.979 10.979V12C22.973 18.061 18.06 22.973 12 22.975"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Unlocked(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 22a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V12a2 2 0 0 0-2-2H6V6.5a4 4 0 0 1 8 0v.282a1.25 1.25 0 1 0 2.5 0v-.033v.002v-.25a6.5 6.5 0 1 0-13 0v3.5H2a2 2 0 0 0-2 2zm8-6.5a2 2 0 1 1 3.092 1.676l-.008.005s.195 1.18.415 2.57v.001a.749.749 0 0 1-.749.749H9.248a.749.749 0 0 1-.749-.749v-.001l.415-2.57a2.002 2.002 0 0 1-.916-1.68z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UnrealEngine(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 0C5.373 0 0 5.373 0 12s5.373 12 12 12s12-5.373 12-12S18.627 0 12 0m0 1.846c5.595.007 10.128 4.545 10.128 10.141c0 5.601-4.54 10.141-10.141 10.141S1.846 17.588 1.846 11.987c0-2.8 1.135-5.335 2.97-7.17a10.122 10.122 0 0 1 7.176-2.97zm0 2.77c-4.392.774-8.308 4.824-8.308 9.23c2.149-3.794 3.584-4.067 3.981-4.067s.606.206.606.663v5.654c0 .703-1.366.588-1.818.519C8.131 19.1 12 19.385 12 19.385l1.846-1.846l1.846.923c2.914-1.334 4.615-4.19 4.615-4.615a5.527 5.527 0 0 1-2.731 1.836l-.039.01c-.245 0-.923-.126-.923-.462V8.538c0-.581 1.342-2.354 1.846-3c-3.332.873-4.298 2.394-4.298 2.394s-.253-.548-1.24-.548c.501.473.838 1.114.922 1.832l.001.014v5.654a5.009 5.009 0 0 1-1.813.801l-.034.006c-.64 0-.952-.26-.952-.75s.029-6.634.029-6.634s-.923.339-.923-1.558c0-.949 1.846-2.135 1.846-2.135z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Upload(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 17.143V24h24v-6.857zm5.143 5.143H1.714v-1.714h3.429zm12-17.143h-3.429v8.571h-3.429V5.143H6.856L11.999 0z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Usb(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.67 2.064L4.921.313A.701.701 0 0 1 6.087.31l.002.003l1.334 1.835c.334.417.167.834-.417.834H6.005v9.84c0 .25.25.25.417.08l2.085-2.085c.311-.336.501-.788.501-1.284l-.001-.053v.003v-2l-.042.001a.96.96 0 0 1-.96-.96l.001-.044v.002V5.48c.01-.548.453-.989 1-.997h1.002l.042-.001c.53 0 .96.43.96.96l-.001.04V5.48v1.005a.96.96 0 0 1-.959 1.002l-.044-.001h.002v2l.002.103c0 .777-.32 1.479-.836 1.981l-.001.001l-2.586 2.586a1.46 1.46 0 0 0-.584 1.341l-.001-.007v4.087a2.252 2.252 0 1 1-1.016.003l.015-.003v-1.082a1.626 1.626 0 0 0-.581-1.415l-.003-.002l-2.586-2.586a2.754 2.754 0 0 1-.834-1.978v-.023v.001v-2.087c-.586-.197-1-.741-1-1.382v-.071a1.502 1.502 0 0 1 3.004 0v.036v-.002v.035c0 .642-.416 1.187-.993 1.38l-.01.003v2.086c0 .527.224 1.001.583 1.332l.001.001l2.085 2.085c.167.167.334.167.334-.08V2.982H3.999c-.58 0-.747-.417-.33-.918z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func UserSecret(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m8.312 22.15l1.385-6.461l-1.385-1.846l-1.846-.923zm3.692 0l1.846-9.229l-1.846.923l-1.385 1.846zm2.307-14.565a.357.357 0 0 0-.058-.087a3.904 3.904 0 0 0-1.401-.114l.016-.001a12.92 12.92 0 0 0-2.493.29l.084-.015a1.59 1.59 0 0 1-.616-.001l.01.002a12.347 12.347 0 0 0-2.396-.274h-.013a4.027 4.027 0 0 0-1.413.121l.028-.006a.346.346 0 0 0-.057.084l-.001.002q.029.26.058.39c.03.038.065.069.106.093l.002.001a.258.258 0 0 1 .108.149v.002c.034.075.07.174.102.274l.006.022q.08.24.101.296t.108.245c.044.1.085.18.129.258l-.007-.013q.036.058.13.202a.65.65 0 0 0 .171.193l.002.001q.08.05.202.137a.687.687 0 0 0 .248.114l.005.001q.13.029.296.058c.103.018.222.029.344.029h.011h-.001l.07.001c.284 0 .553-.067.791-.186l-.01.005c.197-.098.356-.246.466-.428l.003-.005c.08-.142.151-.307.204-.479l.005-.018c.047-.164.103-.305.172-.439l-.006.013a.283.283 0 0 1 .252-.18h.174a.282.282 0 0 1 .252.178l.001.002c.062.12.119.262.161.409l.004.017c.058.191.129.356.215.51l-.007-.013c.112.187.272.335.462.43l.006.003c.228.114.497.181.781.181c.025 0 .049 0 .074-.002h-.003h.01c.121 0 .24-.01.356-.031l-.012.002q.166-.029.296-.058a.71.71 0 0 0 .255-.117l-.002.001q.122-.086.202-.137a.648.648 0 0 0 .171-.191l.002-.003q.094-.144.13-.202c.038-.065.079-.145.115-.227l.007-.017q.086-.187.108-.245t.101-.296c.038-.122.074-.221.116-.316l-.008.02a.256.256 0 0 1 .107-.151l.001-.001a.37.37 0 0 0 .107-.093l.001-.001a3.17 3.17 0 0 0 .059-.376l.001-.014zm5.999 12.676A3.575 3.575 0 0 1 19.257 23a3.891 3.891 0 0 1-2.809.995h.009H3.855a3.888 3.888 0 0 1-2.803-.998l.004.003a3.575 3.575 0 0 1-1.05-2.748v.009q0-.88.065-1.701c.06-.684.157-1.302.292-1.905l-.018.095a9.898 9.898 0 0 1 .565-1.847l-.024.066a5.861 5.861 0 0 1 .923-1.501l-.007.008a4.118 4.118 0 0 1 1.324-1.067l.023-.011l-1.298-3.172h3.086a5.516 5.516 0 0 1-.32-1.845v-.002q0-.173.029-.462q-2.8-.577-2.8-1.385q0-.822 3.028-1.428c.215-.748.468-1.389.773-2l-.03.067A7.561 7.561 0 0 1 6.65.518l-.01.013a1.406 1.406 0 0 1 1.095-.534c.457.054.868.213 1.22.453L8.946.444c.343.234.754.393 1.198.446l.013.001a2.715 2.715 0 0 0 1.22-.453l-.009.006a2.7 2.7 0 0 1 1.2-.446l.013-.001c.443.002.837.209 1.094.531l.002.003a7.52 7.52 0 0 1 .997 1.597l.019.047c.274.543.527 1.184.722 1.85l.021.082q3.028.606 3.028 1.428q0 .808-2.8 1.385a5.565 5.565 0 0 1-.301 2.347l.012-.039h3.086l-1.178 3.248c.631.341 1.151.81 1.541 1.377l.01.015c.415.591.74 1.282.935 2.025l.01.045c.173.612.319 1.353.411 2.111l.007.074c.073.629.114 1.358.114 2.097v.048v-.003z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Venus(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.426 7.714l.001.1a7.414 7.414 0 0 1-1.98 5.054l.004-.005a7.476 7.476 0 0 1-4.848 2.509l-.032.003v3.482h3.016c.228 0 .413.185.413.413v.017v-.001v.873a.413.413 0 0 1-.413.413h-.017h.001h-3v3.016a.413.413 0 0 1-.413.413h-.017h.001h-.873a.413.413 0 0 1-.413-.413v-.017v.001v-3H3.84a.413.413 0 0 1-.413-.413v-.017v.001v-.873c0-.228.185-.413.413-.413h.017h-.001h3v-3.482a7.503 7.503 0 0 1-3.658-1.394l.022.015C1.26 12.576 0 10.295 0 7.719c0-.223.009-.444.028-.662l-.002.028a7.539 7.539 0 0 1 1.099-3.367l-.019.033a7.78 7.78 0 0 1 2.408-2.499l.03-.018A7.412 7.412 0 0 1 6.798.059l.034-.003a7.686 7.686 0 0 1 8.595 7.632v.03v-.002zm-13.716 0l-.001.087c0 1.629.674 3.101 1.759 4.152l.002.001c1.083 1.089 2.583 1.762 4.24 1.762s3.157-.674 4.24-1.762a5.765 5.765 0 0 0 1.761-4.153l-.001-.092v.005l.001-.087a5.763 5.763 0 0 0-1.759-4.152l-.002-.001c-1.083-1.089-2.583-1.762-4.24-1.762s-3.157.674-4.24 1.762a5.763 5.763 0 0 0-1.757 4.245v-.005z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VenusDouble(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23.968 7.087a7.716 7.716 0 0 1-3.171 6.894l-.022.015a7.433 7.433 0 0 1-3.604 1.376l-.032.003v3.482h3.015c.228 0 .413.185.413.413v.017v-.001v.873a.413.413 0 0 1-.413.413h-.017h.001h-2.999v3.015a.413.413 0 0 1-.413.413h-.017h.001h-.873a.413.413 0 0 1-.413-.413v-.017v.001v-3H8.568v3.015a.413.413 0 0 1-.413.413h-.017h.001h-.872a.413.413 0 0 1-.413-.413v-.017v.001v-2.999H3.838a.413.413 0 0 1-.413-.413v-.017v.001v-.873c0-.228.185-.413.413-.413h.017h-.001h3v-3.481a7.503 7.503 0 0 1-3.658-1.394l.022.015c-1.96-1.419-3.22-3.701-3.22-6.276c0-.223.009-.444.028-.663l-.002.029a7.514 7.514 0 0 1 2.2-4.773A7.366 7.366 0 0 1 6.916.045l.028-.002a7.403 7.403 0 0 1 5.073 1.274l-.025-.016a7.369 7.369 0 0 1 4.209-1.306c.297 0 .589.017.877.051l-.035-.003A7.378 7.378 0 0 1 21.76 2.31l.003.003a7.496 7.496 0 0 1 2.202 4.748zm-11.971 4.821a5.77 5.77 0 0 0 1.714-4.112l-.001-.086v.004l.001-.079a5.77 5.77 0 0 0-1.715-4.112a5.77 5.77 0 0 0-1.713 4.196v-.004l-.001.081a5.77 5.77 0 0 0 1.715 4.112m-4.286 1.806h.006c1.07 0 2.075-.283 2.942-.779l-.029.015a7.39 7.39 0 0 1-2.062-5.132l.001-.111v-.073c0-2.001.786-3.818 2.066-5.16l-.003.003a5.866 5.866 0 0 0-2.914-.764h-.006l-.087-.001a5.759 5.759 0 0 0-4.151 1.759l-.001.002a5.763 5.763 0 0 0-1.76 4.245v-.005l-.001.087c0 1.629.674 3.101 1.759 4.151l.002.002a5.763 5.763 0 0 0 4.153 1.761l.092-.001h-.005zm7.715 5.144v-3.482a7.61 7.61 0 0 1-3.455-1.279l.026.017a7.543 7.543 0 0 1-3.393 1.255l-.035.003v3.486zm.856-5.144l.087.001a5.759 5.759 0 0 0 4.151-1.759l.002-.002a5.763 5.763 0 0 0 1.761-4.153l-.001-.092v.005l.001-.087a5.759 5.759 0 0 0-1.759-4.151l-.002-.001a5.763 5.763 0 0 0-4.153-1.761l-.092.001h.005h-.006c-1.07 0-2.075.283-2.942.779l.029-.015a7.451 7.451 0 0 1 2.064 5.158v.084v-.004l.001.103a7.387 7.387 0 0 1-2.066 5.135l.002-.003a5.866 5.866 0 0 0 2.914.764h.004z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VenusMars(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.288.429V.413c0-.228.185-.413.413-.413h.017h-.001h3.864c.234 0 .445.098.595.254a.82.82 0 0 1 .254.595V4.73a.413.413 0 0 1-.413.413H27h.001h-.872a.413.413 0 0 1-.413-.413v-.017v.001V2.92l-3.402 3.415a7.643 7.643 0 0 1 1.689 4.809c0 .483-.044.955-.13 1.413l.007-.047a7.553 7.553 0 0 1-2.201 4.15l-.002.002a7.398 7.398 0 0 1-4.158 2.097l-.04.005a7.269 7.269 0 0 1-1.241.103a7.472 7.472 0 0 1-4.261-1.325l.025.016a7.475 7.475 0 0 1-3.396 1.239l-.033.003v1.768h1.302c.228 0 .413.185.413.413v.017v-.001v.873a.413.413 0 0 1-.413.413h-.017h.001h-1.286v1.306a.413.413 0 0 1-.413.413h-.016h.001h-.874a.413.413 0 0 1-.413-.413v-.017v.001v-1.286H5.556a.413.413 0 0 1-.413-.413v-.017v.001v-.873c0-.228.185-.413.413-.413h.017h-.001h1.286v-1.771a7.535 7.535 0 0 1-3.762-1.48l.019.014A7.72 7.72 0 0 1 .63 14.211l-.019-.05a7.39 7.39 0 0 1-.615-2.978c0-.4.031-.793.091-1.176l-.006.042A7.368 7.368 0 0 1 2.216 5.73l.001-.001a7.527 7.527 0 0 1 4.25-2.203l.042-.005a7.296 7.296 0 0 1 1.244-.103c1.587 0 3.059.49 4.273 1.328l-.025-.016a7.515 7.515 0 0 1 4.245-1.299h.042h-.002h.062a7.47 7.47 0 0 1 4.759 1.702l-.013-.01l3.415-3.402h-1.81a.413.413 0 0 1-.413-.413v-.017v.001zM12.002 15.335a5.77 5.77 0 0 0 1.715-4.113l-.001-.084v.004l.001-.079a5.77 5.77 0 0 0-1.715-4.112a5.77 5.77 0 0 0-1.714 4.197v-.004l-.001.079a5.77 5.77 0 0 0 1.715 4.112M1.717 11.143l-.001.087c0 1.629.674 3.101 1.759 4.152l.002.001a5.765 5.765 0 0 0 4.153 1.761l.092-.001h-.005h.033a5.813 5.813 0 0 0 2.915-.778l-.028.015a7.45 7.45 0 0 1-2.063-5.157v-.16c0-2.001.786-3.818 2.066-5.16l-.003.003a5.774 5.774 0 0 0-2.887-.763h-.035h.002l-.087-.001a5.763 5.763 0 0 0-4.152 1.759l-.001.002a5.765 5.765 0 0 0-1.76 4.245zm14.571 6l.087.001a5.763 5.763 0 0 0 4.152-1.759l.001-.002a5.765 5.765 0 0 0 1.761-4.153l-.001-.092v.005l.001-.087a5.763 5.763 0 0 0-1.759-4.152l-.002-.001a5.765 5.765 0 0 0-4.153-1.761l-.092.001h.005h-.033a5.813 5.813 0 0 0-2.915.778l.028-.015a7.448 7.448 0 0 1 2.062 5.156v.085v-.004v.08a7.452 7.452 0 0 1-2.066 5.16l.003-.003c.83.48 1.825.763 2.887.763h.036z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Viber(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.19 2.337c-.595-.545-3.005-2.297-8.372-2.32c0 0-6.333-.379-9.418 2.452C.684 4.185.083 6.698.014 9.815s-.146 8.959 5.485 10.547h.005l-.005 2.419s-.037.98.607 1.177c.778.244 1.238-.502 1.983-1.304c.408-.441.97-1.087 1.397-1.58a21.822 21.822 0 0 0 7.299-.557l-.15.032c.778-.253 5.18-.816 5.892-6.657c.744-6.03-.354-9.834-2.337-11.555m.652 11.115c-.607 4.875-4.173 5.185-4.829 5.395c-1.432.364-3.076.572-4.769.572c-.486 0-.968-.017-1.445-.051l.064.004s-2.438 2.939-3.197 3.703c-.248.248-.521.228-.516-.267c0-.323.019-4.018.019-4.018c-4.767-1.322-4.491-6.298-4.435-8.897s.544-4.735 1.997-6.169c2.612-2.367 7.989-2.016 7.989-2.016c4.543.019 6.718 1.388 7.224 1.847c1.674 1.435 2.527 4.867 1.898 9.896zm-6.511-3.788v.014a.304.304 0 0 1-.607.015v-.001a1.45 1.45 0 0 0-1.532-1.588h.004a.304.304 0 0 1 .016-.607h.017h-.001l.069-.001a2.037 2.037 0 0 1 2.033 2.175zm.951.531c.047-1.988-1.195-3.544-3.553-3.718a.304.304 0 1 1 .043-.606h-.001h.017a4.106 4.106 0 0 1 4.099 4.348l.001-.011a.304.304 0 0 1-.607-.008v-.008zm2.2.629v.002a.304.304 0 0 1-.607.003c-.024-3.822-2.573-5.903-5.662-5.925a.303.303 0 1 1 0-.606c3.459.024 6.239 2.411 6.267 6.525zm-.525 4.598v.009c-.506.891-1.453 1.875-2.428 1.561l-.009-.014a18.04 18.04 0 0 1-4.831-2.679l.041.03a12.352 12.352 0 0 1-1.969-1.963l-.019-.025a15.082 15.082 0 0 1-1.402-2.106l-.04-.079a13.659 13.659 0 0 1-1.193-2.513l-.03-.098c-.314-.975.665-1.922 1.561-2.428h.009a.851.851 0 0 1 1.12.181l.001.002s.581.693.83 1.036c.234.319.549.83.712 1.115a.976.976 0 0 1-.173 1.246l-.001.001l-.562.45a.821.821 0 0 0-.248.659v-.003a5.899 5.899 0 0 0 3.91 3.941l.042.011a.825.825 0 0 0 .656-.247l.45-.562a.977.977 0 0 1 1.251-.171l-.004-.002c.813.463 1.516.972 2.157 1.549l-.011-.01a.833.833 0 0 1 .176 1.11l.002-.003z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vimeo(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M27.727 5.561q-.17 4.022-5.658 11.094Q16.394 24 12.491 24q-2.42 0-4.09-4.48q-.75-2.726-2.25-8.214q-1.227-4.465-2.675-4.465a8.592 8.592 0 0 0-2.179 1.307l.015-.012l-1.313-1.672q.409-.358 1.84-1.645t2.215-1.968Q6.712.499 8.161.363a2.928 2.928 0 0 1 2.605.943l.002.003a6.77 6.77 0 0 1 1.377 3.438l.003.03q.75 4.891 1.125 6.357q.938 4.24 2.045 4.24q.869 0 2.625-2.744a10.937 10.937 0 0 0 1.844-4.119l.013-.073q.222-2.369-1.858-2.369a5.478 5.478 0 0 0-2.098.457l.035-.014q2.045-6.698 7.822-6.51q4.278.139 4.023 5.558z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vine(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M21.039 11.96a6.531 6.531 0 0 1-1.636.203a6.663 6.663 0 0 1-1.973-.296l.047.013a4.784 4.784 0 0 1-2.591-1.954l-.011-.018a6.209 6.209 0 0 1-.959-3.544v.01a3.685 3.685 0 0 1 .476-2.046l-.01.018c.242-.425.69-.708 1.205-.712h.028c.48 0 .898.262 1.12.651l.003.006a3.482 3.482 0 0 1 .411 1.873v-.01a9.528 9.528 0 0 1-.457 2.752l.019-.068c.3.475.758.826 1.299.983l.016.004c.423.139.909.22 1.414.22h.069h-.003l.712-.11a9.68 9.68 0 0 0 .822-3.832v-.005a6.9 6.9 0 0 0-1.056-4.192l.017.028A4.892 4.892 0 0 0 15.71.017l.017-.001a5.059 5.059 0 0 0-4.077 1.773l-.006.007a6.678 6.678 0 0 0-1.507 4.506V6.29l-.001.169c0 1.714.437 3.326 1.205 4.731l-.026-.052a7.624 7.624 0 0 0 3.162 3.157l.041.02a22.442 22.442 0 0 1-4.167 5.866l.002-.002a27.057 27.057 0 0 1-3.712-6.126l-.068-.175C5.337 10.881 4.464 7.407 4.119 3.779l-.011-.147H-.001c.477 4.241 1.482 8.114 2.955 11.747l-.106-.295c.908 2.372 2.11 4.42 3.592 6.231l-.031-.039a10.58 10.58 0 0 0 2.535 2.331l.04.025c.343.232.766.37 1.222.37c.419 0 .811-.117 1.144-.321l-.01.005a10.843 10.843 0 0 0 1.857-1.582l.006-.006a26.148 26.148 0 0 0 2.525-2.942l.05-.072a22.646 22.646 0 0 0 2.351-3.81l.06-.135a13.197 13.197 0 0 0 2.938-.346l-.089.018z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Visa(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M33.6 24H2.4A2.4 2.4 0 0 1 0 21.6V2.4A2.4 2.4 0 0 1 2.4 0h31.2A2.4 2.4 0 0 1 36 2.4v19.2a2.4 2.4 0 0 1-2.4 2.4m-15.76-9.238l-.359 2.25a6.84 6.84 0 0 0 2.903.531h-.011a5.167 5.167 0 0 0 3.275-.933l-.017.011a3.085 3.085 0 0 0 1.258-2.485v-.015v.001c0-1.1-.736-2.014-2.187-2.72a7.653 7.653 0 0 1-1.132-.672l.023.016a.754.754 0 0 1-.343-.592v-.002a.736.736 0 0 1 .379-.6l.004-.002a1.954 1.954 0 0 1 1.108-.257h-.006h.08l.077-.001c.644 0 1.255.139 1.806.388l-.028-.011l.234.125l.359-2.171a6.239 6.239 0 0 0-2.277-.422h-.049h.003a5.067 5.067 0 0 0-3.157.932l.016-.011a2.922 2.922 0 0 0-1.237 2.386v.005c-.01 1.058.752 1.972 2.266 2.72c.4.175.745.389 1.054.646l-.007-.006a.835.835 0 0 1 .297.608v.004c0 .319-.19.593-.464.716l-.005.002c-.3.158-.656.25-1.034.25h-.046h.002h-.075c-.857 0-1.669-.19-2.397-.53l.035.015l-.343-.172zm10.125 1.141h3.315q.08.343.313 1.5H34L31.906 7.372h-2a1.334 1.334 0 0 0-1.357.835l-.003.009l-3.84 9.187h2.72l.546-1.499zM14.891 7.372l-1.626 10.031h2.594l1.625-10.031zM4.922 9.419l2.11 7.968h2.734l4.075-10.015h-2.746l-2.534 6.844l-.266-1.391l-.904-4.609a1.042 1.042 0 0 0-1.177-.844l.006-.001H2.033l-.031.203c3.224.819 5.342 2.586 6.296 5.25A5.74 5.74 0 0 0 6.972 10.8l-.001-.001a6.103 6.103 0 0 0-2.007-1.368l-.04-.015zm25.937 4.421h-2.16q.219-.578 1.032-2.8l.046-.141l.16-.406c.066-.166.11-.302.14-.406l.188.859l.593 2.89z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VisualStudio(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m17.858 23.998l-9.771-9.484l-5.866 4.465L0 17.864V6.145l2.234-1.121l5.87 4.469L17.851 0l5.587 2.239V21.77L17.859 24zm-.563-16.186l-5.577 4.173l5.58 4.202zM2.788 9.497v5.016l2.787-2.525z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vk(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23.456 5.784a8.605 8.605 0 0 1-1.09 2.259l.019-.03q-.672 1.12-1.605 2.588q-.8 1.159-.847 1.2a1.28 1.28 0 0 0-.267.618l-.001.007a.897.897 0 0 0 .268.535l.4.446q3.21 3.299 3.611 4.548a.89.89 0 0 1-.112.829l.002-.003a.965.965 0 0 1-.784.289h.004h-2.636c-.337 0-.647-.118-.89-.314l.003.002a6.928 6.928 0 0 1-.951-.948l-.009-.012q-.691-.781-1.226-1.315q-1.782-1.694-2.63-1.694a.788.788 0 0 0-.516.135l.003-.002a.767.767 0 0 0-.16.584v-.004a12.532 12.532 0 0 0-.038 1.403v-.017v1.159a.78.78 0 0 1-.266.757l-.001.001a3.179 3.179 0 0 1-1.617.267l.013.001a8.323 8.323 0 0 1-4.275-1.268l.035.02A11.931 11.931 0 0 1 4.176 14.3l-.027-.042a26.36 26.36 0 0 1-2.471-3.992l-.07-.154A24.657 24.657 0 0 1 .375 7.31l-.06-.185a6.646 6.646 0 0 1-.31-1.535l-.002-.025q0-.758.892-.758h2.63a1.058 1.058 0 0 1 .739.225l-.002-.002c.2.219.348.488.421.788l.003.012a25.422 25.422 0 0 0 1.587 3.615l-.067-.137a14.56 14.56 0 0 0 1.623 2.576l-.023-.031q.8.982 1.248.982l.032.001a.4.4 0 0 0 .347-.2l.001-.002a1.783 1.783 0 0 0 .111-.787v.006v-3.879a3.211 3.211 0 0 0-.32-1.267l.008.019a2.956 2.956 0 0 0-.45-.695l.003.004a1.099 1.099 0 0 1-.311-.619l-.001-.006c0-.17.078-.323.2-.423l.001-.001a.678.678 0 0 1 .46-.178h4.154a.634.634 0 0 1 .559.222l.001.001a1.36 1.36 0 0 1 .159.763v-.005v5.173a.993.993 0 0 0 .136.584l-.002-.004a.401.401 0 0 0 .333.178h.001a.946.946 0 0 0 .471-.162l-.003.002c.272-.187.506-.4.709-.641l.004-.005a15.606 15.606 0 0 0 1.655-2.25l.039-.07c.344-.57.716-1.272 1.053-1.993l.062-.147l.446-.892a1.122 1.122 0 0 1 1.117-.759h-.003h2.631q1.066 0 .8.981z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeDown(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.881.133L7.09 5.188v13.625l8.791 5.055c.726.416 1.572-.202 1.572-1.145V1.279c0-.945-.847-1.564-1.572-1.145zM6 5.727l-4.926.01A1.184 1.184 0 0 0 .001 7v-.004v10.256a1.167 1.167 0 0 0 1.059 1.246h4.939V5.726zm20.289 6.307v.009c0 .876-.265 1.691-.719 2.367l.01-.015a3.975 3.975 0 0 1-1.85 1.551l-.026.009a.915.915 0 0 1-.383.082l-.036-.001h.002h-.012c-.289 0-.551-.118-.739-.308a1.018 1.018 0 0 1-.317-.738v-.047c0-.216.076-.413.202-.568l-.001.002c.137-.163.296-.302.475-.412l.008-.005q.284-.175.567-.384c.204-.159.367-.36.479-.591l.004-.01c.126-.277.199-.6.199-.941s-.073-.664-.205-.955l.006.015a1.682 1.682 0 0 0-.48-.598l-.004-.003q-.284-.209-.567-.384a1.967 1.967 0 0 1-.481-.414l-.002-.003a.894.894 0 0 1-.201-.567v-.047c0-.29.121-.552.316-.738c.189-.191.45-.309.739-.309h.012h-.001l.035-.001c.139 0 .27.03.388.085l-.006-.002a3.862 3.862 0 0 1 1.868 1.539l.009.016c.445.666.709 1.484.709 2.365v.001zm4.272 0v.037c0 1.743-.53 3.362-1.437 4.706l.019-.03a8.287 8.287 0 0 1-3.701 3.125l-.054.02a1.16 1.16 0 0 1-.417.082h-.019c-.294 0-.56-.121-.75-.317a1.024 1.024 0 0 1-.317-.741v-.012c0-.439.267-.815.648-.976l.007-.003a10.31 10.31 0 0 0 1.303-.759l-.035.023a6.382 6.382 0 0 0 .018-10.296l-.018-.012a9.95 9.95 0 0 0-1.208-.709l-.06-.027a1.062 1.062 0 0 1-.655-.979v-.003a1.083 1.083 0 0 1 1.067-1.068h.001c.156 0 .305.03.442.085l-.008-.003a8.286 8.286 0 0 1 3.735 3.113l.019.032a8.347 8.347 0 0 1 1.418 4.676v.039v-.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeMute(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.885.133L7.094 5.188v13.625l8.791 5.055c.726.416 1.572-.202 1.572-1.145V1.279c0-.945-.847-1.564-1.572-1.145zM6.003 5.727l-4.926.01A1.184 1.184 0 0 0 .004 7v-.004v10.256a1.167 1.167 0 0 0 1.059 1.246h4.939V5.726z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeOff(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.883.133L7.09 5.188v13.625l8.793 5.055c.724.416 1.571-.202 1.571-1.145V1.279c0-.945-.847-1.564-1.571-1.145zM6 5.615l-4.926.01A1.184 1.184 0 0 0 .001 6.888v-.004V17.14a1.167 1.167 0 0 0 1.059 1.246H6zm23.196 2.893l-.887-.884l-3.491 3.491l-3.492-3.491l-.884.884l3.491 3.491l-3.491 3.492l.884.887l3.491-3.491l3.492 3.491l.887-.887L25.702 12z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func VolumeUp(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.881.133L7.09 5.188v13.625l8.791 5.055c.726.416 1.572-.202 1.572-1.145V1.279c0-.945-.847-1.564-1.572-1.145zM6 5.727l-4.926.01A1.184 1.184 0 0 0 .001 7v-.004v10.256a1.167 1.167 0 0 0 1.059 1.246h4.939V5.726zm20.289 6.136v.009c0 .876-.265 1.691-.719 2.367l.01-.015a3.975 3.975 0 0 1-1.85 1.551l-.026.009a.915.915 0 0 1-.383.082l-.036-.001h.002h-.012c-.289 0-.551-.118-.739-.308a1.018 1.018 0 0 1-.317-.738v-.047c0-.216.076-.413.202-.568l-.001.002c.137-.163.296-.302.475-.412l.008-.005q.284-.175.567-.384c.204-.159.367-.36.479-.591l.004-.01c.126-.277.199-.6.199-.941s-.073-.664-.205-.955l.006.015a1.682 1.682 0 0 0-.48-.598l-.004-.003q-.284-.209-.567-.384a1.967 1.967 0 0 1-.481-.414l-.002-.003a.894.894 0 0 1-.201-.567v-.047c0-.29.121-.552.316-.738c.189-.191.45-.309.739-.309h.012h-.001l.035-.001c.139 0 .27.03.388.085l-.006-.002a3.862 3.862 0 0 1 1.868 1.539l.009.016c.445.666.709 1.484.709 2.365v.001zm4.272 0v.037c0 1.743-.53 3.362-1.437 4.706l.019-.03a8.287 8.287 0 0 1-3.701 3.125l-.054.02a1.16 1.16 0 0 1-.417.082h-.019c-.294 0-.56-.121-.75-.317a1.024 1.024 0 0 1-.317-.741v-.012c0-.439.267-.815.648-.976l.007-.003a10.31 10.31 0 0 0 1.303-.759l-.035.023a6.382 6.382 0 0 0 .018-10.296l-.018-.012a9.95 9.95 0 0 0-1.208-.709l-.06-.027a1.062 1.062 0 0 1-.655-.979v-.003a1.083 1.083 0 0 1 1.067-1.068h.001c.156 0 .305.03.442.085l-.008-.003a8.286 8.286 0 0 1 3.735 3.113l.019.032a8.347 8.347 0 0 1 1.418 4.676v.039v-.002zm4.271 0c-.008 5.249-3.161 9.76-7.676 11.749l-.082.032a1.083 1.083 0 0 1-1.502-.985v-.001a1.13 1.13 0 0 1 .648-.979l.007-.003c.1-.055.223-.114.351-.166l.025-.009c.152-.061.276-.12.395-.185l-.02.01c.531-.287.978-.569 1.405-.877l-.038.026a10.714 10.714 0 0 0 3.176-3.727l.028-.06c.726-1.403 1.151-3.063 1.151-4.822s-.426-3.419-1.179-4.882l.028.06a10.711 10.711 0 0 0-3.175-3.767l-.029-.02c-.39-.282-.837-.564-1.302-.818l-.066-.033a3.498 3.498 0 0 0-.35-.166l-.025-.009a3.594 3.594 0 0 1-.395-.185l.02.01a1.13 1.13 0 0 1-.654-.98v-.002A1.082 1.082 0 0 1 26.639 0h.001c.156 0 .305.03.442.085l-.008-.003c4.597 2.021 7.751 6.532 7.759 11.781v.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Vuejs(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M22.313 0h5.688l-14 24l-14-24h11l3 5.563L17.501 0zM3.5 2L14 20L24.5 2h-3.375L14 14.375L6.875 2z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wallet(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M23.596 13.43H11.562a2.46 2.46 0 0 0-2.639 2.231l-.001.009a2.472 2.472 0 0 0 2.683 2.239l-.009.001h12a.57.57 0 0 0 .621-.527v-3.424a.61.61 0 0 0-.604-.529h-.018zm-12.034 3.638a1.43 1.43 0 1 1 1.43-1.43v.004c0 .788-.639 1.426-1.426 1.426zM3.295 0v6.404h1.368V3.668c.138.034.298.056.462.062h.004a1.871 1.871 0 0 0 1.866-1.865a2.098 2.098 0 0 0-.066-.481l.003.015h7.492a2.324 2.324 0 0 0-.062.462v.004a1.871 1.871 0 0 0 1.865 1.866c.168-.006.329-.028.483-.066l-.016.003v2.736h1.401V0zm15.389 5.626h1.741a2.266 2.266 0 0 0-1.727-1.49l-.014-.002z"/><path fill="currentColor" d="M11.471 12.684h11.938V9.047a2.066 2.066 0 0 0-2.217-1.865h.007h-18.96a.097.097 0 0 0-.062-.031a.814.814 0 0 1-.466-.682v-.002a.886.886 0 0 1 .647-.776l.006-.001V4.105C1.057 4.167 0 6.032 0 8.364v13.772A2.067 2.067 0 0 0 2.215 24h-.007h18.96a2.065 2.065 0 0 0 2.207-1.857l.001-.008v-3.637H11.504c-1.927 0-3.482-1.306-3.482-2.891c0-1.616 1.554-2.922 3.45-2.922z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Webpack(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M20.952 18.5L11.238 24v-4.286l6.048-3.334zm.666-.595V6.385L18.07 8.456v7.381l3.548 2.048zM.666 18.5L10.38 24v-4.286L4.332 16.38zM0 17.905V6.385l3.548 2.071v7.381zM.405 5.643L10.381 0v4.143L3.952 7.691zm20.809 0L11.238 0v4.143l6.428 3.548zM10.381 18.738l-5.972-3.287V8.927l5.976 3.452zm.857 0l5.976-3.286V8.928l-5.976 3.452zM4.809 8.19l6-3.31l6 3.31l-6 3.453z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Webstorm(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M4.65 19.351h14.7V4.65H4.65zm6.9-1.726H5.926v-.945h5.625zm4.125-11.3a3.743 3.743 0 0 1 2.287.806l-.008-.006l-.7 1.05a2.753 2.753 0 0 0-1.567-.6h-.007c-.5 0-.8.225-.8.57c0 .4.271.555 1.319.81c1.231.345 1.951.8 1.951 1.875a1.98 1.98 0 0 1-2.261 1.948l.011.001a4.313 4.313 0 0 1-2.632-1.01l.007.005l.779-.976a2.87 2.87 0 0 0 1.841.75h.005c.555 0 .9-.225.9-.6c0-.346-.227-.525-1.23-.8c-1.246-.33-2.04-.676-2.04-1.9v-.044a1.933 1.933 0 0 1 2.154-1.874l-.009-.001zm-8.625.12l.944 3.6l1.05-3.6H10.1l1.051 3.6l.929-3.6h1.471l-1.8 6.225h-1.176l-1-3.6l-1.02 3.6H7.4L5.595 6.45z"/><path fill="currentColor" d="m24 8.13l-1.919-4.755L20.88.5L15.3 0l-.87.84l-1.11-.42L9.8 2.3L6 .045l-6 2.46l3.225 19.171L16.5 24l.84-4.26l.15-.091H4.35V4.35h15.3v14.025L24 15.8l-2.625-4.92z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wetransfer(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M2.705 15.571L.109 8.223A1.795 1.795 0 0 1 0 7.659v-.003a1.04 1.04 0 0 1 1.08-1.024h-.002h12.094c.566 0 1.025.459 1.026 1.025c-.009.18-.041.349-.093.509l.004-.014l-2.628 7.418a1.292 1.292 0 0 1-1.223.986h-.086a1.23 1.23 0 0 1-1.195-.938l-.002-.008l-1.166-3.552a.726.726 0 0 0-.691-.505l-.03.001h.001h-.023a.746.746 0 0 0-.704.5l-.002.005l-1.167 3.553A1.227 1.227 0 0 1 4 16.558h-.07a1.257 1.257 0 0 1-1.223-.977l-.002-.008zm20.218-1.668a.956.956 0 0 0-.677.128l.004-.002a4.063 4.063 0 0 1-2.133.64h-.003a2.839 2.839 0 0 1-2.959-2.594l-.001-.01a8.002 8.002 0 0 0 2.906.536c.164 0 .327-.005.489-.014l-.022.001h.056a5.617 5.617 0 0 0 2.765-.722l-.028.015a1.568 1.568 0 0 0 .637-1.735l.003.011a4.129 4.129 0 0 0-4.331-3.755h.01a4.88 4.88 0 0 0-5.027 5.196l-.001-.014l-.001.118a5.073 5.073 0 0 0 5.341 5.066l-.012.001a6.226 6.226 0 0 0 3.384-1.005l-.024.015c.232-.164.395-.411.446-.697l.001-.007a.99.99 0 0 0-.816-1.166l-.006-.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Whatsapp(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M15.271 13.21a7.014 7.014 0 0 1 1.543.7l-.031-.018c.529.235.986.51 1.403.833l-.015-.011c.02.061.032.13.032.203l-.001.032v-.001c-.015.429-.11.832-.271 1.199l.008-.021c-.231.463-.616.82-1.087 1.01l-.014.005a3.624 3.624 0 0 1-1.576.411h-.006a8.342 8.342 0 0 1-2.988-.982l.043.022a8.9 8.9 0 0 1-2.636-1.829l-.001-.001a20.473 20.473 0 0 1-2.248-2.794l-.047-.074a5.38 5.38 0 0 1-1.1-2.995l-.001-.013v-.124a3.422 3.422 0 0 1 1.144-2.447l.003-.003a1.17 1.17 0 0 1 .805-.341h.001c.101.003.198.011.292.025l-.013-.002c.087.013.188.021.292.023h.003a.642.642 0 0 1 .414.102l-.002-.001c.107.118.189.261.238.418l.002.008q.124.31.512 1.364c.135.314.267.701.373 1.099l.014.063a1.573 1.573 0 0 1-.533.889l-.003.002q-.535.566-.535.72a.436.436 0 0 0 .081.234l-.001-.001a7.03 7.03 0 0 0 1.576 2.119l.005.005a9.89 9.89 0 0 0 2.282 1.54l.059.026a.681.681 0 0 0 .339.109h.002q.233 0 .838-.752t.804-.752zm-3.147 8.216h.022a9.438 9.438 0 0 0 3.814-.799l-.061.024c2.356-.994 4.193-2.831 5.163-5.124l.024-.063c.49-1.113.775-2.411.775-3.775s-.285-2.662-.799-3.837l.024.062c-.994-2.356-2.831-4.193-5.124-5.163l-.063-.024c-1.113-.49-2.411-.775-3.775-.775s-2.662.285-3.837.799l.062-.024c-2.356.994-4.193 2.831-5.163 5.124l-.024.063a9.483 9.483 0 0 0-.775 3.787a9.6 9.6 0 0 0 1.879 5.72l-.019-.026l-1.225 3.613l3.752-1.194a9.45 9.45 0 0 0 5.305 1.612h.047zm0-21.426h.033c1.628 0 3.176.342 4.575.959L16.659.93c2.825 1.197 5.028 3.4 6.196 6.149l.029.076c.588 1.337.93 2.896.93 4.535s-.342 3.198-.959 4.609l.029-.074c-1.197 2.825-3.4 5.028-6.149 6.196l-.076.029c-1.327.588-2.875.93-4.503.93h-.034h.002h-.053c-2.059 0-3.992-.541-5.664-1.488l.057.03L-.001 24l2.109-6.279a11.505 11.505 0 0 1-1.674-6.01c0-1.646.342-3.212.959-4.631l-.029.075C2.561 4.33 4.764 2.127 7.513.959L7.589.93A11.178 11.178 0 0 1 12.092 0zh-.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wheelchair(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M6.742.18c.396.229.714.553.932.94l.006.012a2.038 2.038 0 0 1 .031 1.544l.005-.014A2.086 2.086 0 0 1 5.88 3.947h-.004a2.082 2.082 0 0 1-1.843-1.278l-.005-.014a2.04 2.04 0 0 1-.123-.705c0-.294.061-.575.172-.828l-.005.013A2.297 2.297 0 0 1 5.319.05l.016-.005A2.55 2.55 0 0 1 6.76.186L6.743.18zm.332 4.23c.439.12.8.399 1.021.771l.004.008c.176.404.331.882.439 1.377l.01.052c.145.614.296 1.242.332 1.386l.072.274h2.288c2.035 0 2.32.014 2.534.13c.317.15.537.458.56.82v.003a1.008 1.008 0 0 1-.373.923l-.002.002l-.195.181l-2.16.022c-1.184.014-2.16.029-2.16.036s.094.39.21.845l.195.83l2.24.036c1.913.029 2.267.05 2.418.15c.576.849 1.15 1.835 1.658 2.859l.068.151l1.546 2.88l.938-.31c1.047-.354 1.473-.4 1.805-.181c.278.184.463.49.48.84v.002a1.064 1.064 0 0 1-.398.878l-.002.002c-.466.227-1.022.441-1.597.609l-.077.019c-1.711.56-1.97.578-2.344.195c-.116-.116-.91-1.509-1.754-3.092l-1.546-2.88l-3.104-.036c-3.551-.043-3.205.036-3.465-.758c-.303-.91-1.754-7.074-1.754-7.457a1.646 1.646 0 0 1 .758-1.339l.007-.004a1.77 1.77 0 0 1 .981-.296c.1.028.223.053.349.069l.014.002zM4.461 9.506c.122.513.224.974.224 1.018a3.298 3.298 0 0 1-.572.616l-.005.004a6.183 6.183 0 0 0-2.023 4.581c0 .828.162 1.619.457 2.341l-.015-.041a6.637 6.637 0 0 0 2.557 3.016l.027.016a6.752 6.752 0 0 0 2.199.794l.041.006a7.369 7.369 0 0 0 2.635-.16l-.051.01a6.369 6.369 0 0 0 3.664-2.809l.016-.028c.181-.31.346-.534.375-.505s.289.49.578 1.026c.598 1.105.598 1.018-.022 1.726a8.313 8.313 0 0 1-5.058 2.794l-.046.006A8.306 8.306 0 0 1 .434 18.39l-.017-.058a8.1 8.1 0 0 1-.423-2.61c0-2.077.766-3.976 2.031-5.428l-.008.01a8.344 8.344 0 0 1 2.09-1.711l.04-.021c.145.27.255.583.313.915l.003.019z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wifi(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M16.807 0h-.045C10.572 0 4.942 2.4.752 6.319l.013-.012L0 7.02l3.862 3.826l.72-.66c3.201-2.952 7.494-4.763 12.21-4.763s9.009 1.81 12.222 4.774l-.012-.011l.72.66l3.862-3.826l-.765-.713A23.292 23.292 0 0 0 16.845 0h-.041h.002z"/><path fill="currentColor" d="M27.405 12.03A15.673 15.673 0 0 0 16.83 7.95h-.674l-.007.015a15.716 15.716 0 0 0-9.958 4.076l.013-.012l-.787.713l3.893 3.855l.72-.63c1.791-1.606 4.171-2.587 6.78-2.587s4.989.982 6.79 2.596l-.01-.008l.72.63l3.893-3.854z"/><path fill="currentColor" d="m16.815 24l5.475-5.415l-.87-.713a7.022 7.022 0 0 0-4.625-1.5h.013h-.066a7.603 7.603 0 0 0-4.567 1.515l.02-.014l-.862.713l.795.787l3.96 3.915z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WifiLogo(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M29.294 7.765a1.06 1.06 0 0 0-.001 2.118h.002a1.06 1.06 0 0 0 .001-2.118zm-1.059 7.412v.029a1.059 1.059 0 1 0 2.118 0v-.031v.001v-3.559a1.059 1.059 0 1 0-2.118 0v.031v-.002zm-2.47-7.412a1.06 1.06 0 0 1 .001 2.118h-3.884v1.412h2.47a1.06 1.06 0 0 1 .001 2.118h-2.472v1.794a1.059 1.059 0 1 1-2.118 0v-.031v.002v-6.359c0-.582.472-1.054 1.054-1.054h.005zm8.117 2.117v4.235a5.647 5.647 0 0 1-5.647 5.647h-2.118c-2.225 2.599-5.51 4.235-9.176 4.235S9.99 22.363 7.778 19.78l-.014-.016H5.646a5.647 5.647 0 0 1-5.647-5.647V9.882a5.647 5.647 0 0 1 5.647-5.647h2.118C9.989 1.636 13.274 0 16.94 0s6.951 1.636 9.163 4.219l.014.016h2.118a5.647 5.647 0 0 1 5.647 5.647M14.47 7.765a1.06 1.06 0 0 0-.001 2.118h.001a1.06 1.06 0 0 0 .001-2.118h-.002zm-1.059 7.412a1.06 1.06 0 0 0 2.118.001v-3.532a1.06 1.06 0 0 0-2.118-.001v.002zm-3.177.117l1.647-6a1.067 1.067 0 1 0-1.999-.712l-.001.007l-.706 2.588l-.706-2.588a1.105 1.105 0 0 0-2.116-.007l-.002.008l-.706 2.588l-.706-2.588a1.064 1.064 0 1 0-1.997.712l-.003-.007l1.647 6a1.07 1.07 0 0 0 2.117.005l.001-.006l.706-2.47l.706 2.47a1.062 1.062 0 0 0 2.117.005zm18-9.647h-7.425a3.163 3.163 0 0 0-3.163 3.163v.014v-.001v6.372a3.166 3.166 0 0 1-2.93 3.157l-.01.001h13.548a4.218 4.218 0 0 0 4.218-4.218v-.019v.001v-4.254a4.218 4.218 0 0 0-4.218-4.218h-.019h.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wikipedia(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m15.562 19.478l-3.07-7.238q-.258.51-1.652 3.182t-2.068 4.057a1.493 1.493 0 0 1-.292.004h.005l-.062.001c-.075 0-.15-.006-.222-.018l.008.001q-.854-2.01-2.662-6.114T2.843 7.144a4.935 4.935 0 0 0-.699-1.128l.007.008a5.747 5.747 0 0 0-1.063-1.035l-.015-.011a1.867 1.867 0 0 0-1.055-.447L.011 4.53q0-.052-.005-.25t-.005-.282h6.073v.521a2.98 2.98 0 0 0-.846.173l.021-.007a1.631 1.631 0 0 0-.692.447l-.001.001a.566.566 0 0 0-.102.669l-.002-.003q.271.614 2.255 5.2t2.453 5.626q.32-.635 1.458-2.776t1.364-2.578q-.198-.406-1.313-2.926T9.253 5.272q-.396-.72-2.094-.74v-.523l5.343.01v.49h-.04c-.345 0-.667.096-.942.263l.008-.005q-.349.24-.13.72q.344.73.906 1.974t.895 1.953q1.146-2.229 1.802-3.781q.25-.573-.103-.826a2.526 2.526 0 0 0-1.35-.276h.008c.007-.065.011-.14.011-.216V4.27v.002v-.25q.666 0 1.776-.005l2.834-.015v.51a2.738 2.738 0 0 0-1.253.351l.014-.007a2.596 2.596 0 0 0-.931.828l-.006.009l-2.219 4.604q.135.344 1.328 3.021t1.266 2.854L20.97 5.579a1.356 1.356 0 0 0-.511-.648l-.005-.003a2.038 2.038 0 0 0-.663-.324l-.015-.004a2.91 2.91 0 0 0-.572-.08h-.006v-.524l4.792.042l.01.021l-.01.458a2.252 2.252 0 0 0-2.089 1.495l-.005.016q-5.482 12.666-5.824 13.447z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wind(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M8.417 19.246a1.005 1.005 0 0 1 2.01 0a2.749 2.749 0 0 0 2.745 2.745a3.296 3.296 0 0 0 3.292-3.292a3.922 3.922 0 0 0-3.919-3.917H1.005a1.005 1.005 0 0 1 0-2.01h11.54a5.936 5.936 0 0 1 5.928 5.928a5.31 5.31 0 0 1-5.3 5.301a4.76 4.76 0 0 1-4.756-4.754zm5.702-8.015a1.005 1.005 0 0 1 0-2.01h6.156a3.922 3.922 0 0 0 3.918-3.92a3.296 3.296 0 0 0-3.292-3.292a2.749 2.749 0 0 0-2.745 2.745a1.005 1.005 0 0 1-2.01 0A4.76 4.76 0 0 1 20.901 0a5.31 5.31 0 0 1 5.301 5.3a5.936 5.936 0 0 1-5.928 5.929zm-13.114 0a1.005 1.005 0 0 1 0-2.01h6.158a3.922 3.922 0 0 0 3.917-3.92a3.296 3.296 0 0 0-3.292-3.29a2.749 2.749 0 0 0-2.745 2.745a1.005 1.005 0 0 1-2.01 0A4.761 4.761 0 0 1 7.788 0a5.31 5.31 0 0 1 5.301 5.3a5.936 5.936 0 0 1-5.932 5.929z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Windows(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.84 12.663v9.39L0 20.697v-8.034zm0-10.72v9.505H0V3.303zM24 12.663V24l-13.082-1.803v-9.534zM24 0v11.452H10.918V1.803z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wink(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 1.316C6.099 1.316 1.316 6.099 1.316 12S6.099 22.684 12 22.684S22.684 17.901 22.684 12c-.012-5.896-4.788-10.672-10.683-10.684zm0 22.297C5.586 23.613.387 18.414.387 12S5.586.387 12 .387S23.613 5.586 23.613 12v.015c0 6.405-5.192 11.597-11.597 11.597z"/><path fill="currentColor" d="M12 24C5.386 23.966.034 18.614 0 12.003V12C0 5.373 5.373 0 12 0s12 5.373 12 12c-.034 6.614-5.386 11.966-11.997 12zM12 .774C5.8.774.774 5.8.774 12S5.8 23.226 12 23.226S23.226 18.2 23.226 12C23.222 5.802 18.198.779 12.001.774zm0 22.297C5.886 23.071.929 18.114.929 12S5.886.929 12 .929S23.071 5.886 23.071 12S18.114 23.071 12 23.071m0-21.368C6.313 1.703 1.703 6.313 1.703 12S6.313 22.297 12 22.297S22.297 17.687 22.297 12v-.005c0-5.684-4.608-10.292-10.292-10.292z"/><path fill="currentColor" d="M9.677 9.91v.009c0 1.15-.932 2.082-2.082 2.082h-.009a2.09 2.09 0 1 1 2.09-2.09zM12 19.665a.465.465 0 1 1 0-.93h.011a6.933 6.933 0 0 0 6.555-4.674l.015-.049a.44.44 0 0 1 .545-.308l-.003-.001a.44.44 0 0 1 .309.545l.001-.003c-1.041 3.17-3.974 5.419-7.432 5.419z"/><path fill="currentColor" d="M12 20.052a.86.86 0 0 1-.852-.851a.916.916 0 0 1 .849-.851h.079a6.415 6.415 0 0 0 6.105-4.445l.013-.045a.795.795 0 0 1 .383-.462l.004-.002a.858.858 0 0 1 .626-.076l-.006-.001a.839.839 0 0 1 .541 1.09l.002-.006c-1.058 3.303-4.101 5.652-7.692 5.652H12h.003zm0-.852c-.02.02-.033.047-.033.077s.013.058.033.077h.017a7.393 7.393 0 0 0 7.014-5.058l.015-.052v-.077h-.077c-.992 2.947-3.729 5.032-6.954 5.032zm5.961-9.058h-1.703l.697-.542a.606.606 0 1 0-.775-.929l.001-.001l-2.013 1.548c-.09.112-.168.239-.228.376l-.004.011v.155c.03.306.285.542.596.542h.025h-.001h3.415a.61.61 0 0 0 .611-.611v-.009a.599.599 0 0 0-.596-.542h-.025h.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wix(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M3.438 13.605c.073-.281.149-.56.222-.842c.284-1.083.565-2.17.855-3.254c.105-.442.259-.831.459-1.19l-.012.024a1.426 1.426 0 0 1 1.016-.699l.008-.001a1.61 1.61 0 0 1 1.52.479l.001.001c.136.156.243.341.31.544l.003.011c.185.492.377 1.125.535 1.772l.025.124q.391 1.482.786 2.966c.006.023 0 .05.033.07c.056-.222.116-.442.172-.661l1.018-3.97c.062-.248.126-.495.191-.743c.049-.185.124-.346.222-.491l-.004.006a.981.981 0 0 1 .649-.422l.006-.001a5.298 5.298 0 0 1 1.044-.104h.001c.073 0 .086.017.07.086q-.466 1.779-.928 3.558q-.684 2.631-1.371 5.259c-.046.172-.089.347-.135.518c-.01.033-.014.062-.059.066a2.618 2.618 0 0 1-.99-.1l.018.005a.82.82 0 0 1-.31-.195a4.25 4.25 0 0 1-.335-.351l-.005-.006a2.684 2.684 0 0 1-.287-.569l-.006-.019a13.373 13.373 0 0 1-.679-2.192l-.017-.094c-.281-1.12-.568-2.24-.852-3.36a.531.531 0 0 0-.087-.186l.001.001a.25.25 0 0 0-.218-.129h-.005a.28.28 0 0 0-.239.189l-.001.002c-.096.304-.165.614-.248.922c-.334 1.288-.658 2.58-1.001 3.865c-.102.383-.271.743-.4 1.12c-.138.361-.369.66-.665.876l-.005.004a.97.97 0 0 1-.481.195h-.005a3.904 3.904 0 0 1-.675.029h.008c-.059 0-.08-.023-.099-.08q-.48-1.823-.968-3.643q-.631-2.384-1.27-4.777l-.238-.902c-.006-.03-.032-.073.03-.079a2.625 2.625 0 0 1 1.466.225l-.016-.007a.98.98 0 0 1 .487.581l.002.007c.172.56.301 1.126.446 1.691q.495 1.913.984 3.826a.243.243 0 0 0 .014.048l-.001-.001zm11.905 3.134l.278-.416q1.44-2.146 2.88-4.288a.083.083 0 0 0 .018-.051a.078.078 0 0 0-.025-.058q-1.194-1.774-2.385-3.548l-.72-1.083c-.01-.017-.033-.033-.026-.066l.489.03c.304.02.602.04.902.056c.244.03.461.129.636.276l-.002-.002c.158.12.291.264.396.426l.004.007c.518.786 1.03 1.576 1.569 2.346c.24.34.343.307.598.056c.126-.149.242-.316.342-.493l.008-.016c.446-.664.886-1.334 1.341-1.992a1.569 1.569 0 0 1 1.089-.65l.007-.001a8.387 8.387 0 0 1 1.227-.043h-.015c.08 0 .03.043.014.066c-.222.33-.442.661-.667.991q-1.22 1.822-2.438 3.633a.102.102 0 0 0-.023.064a.1.1 0 0 0 .029.071q1.546 2.299 3.089 4.598c.014.023.046.046.036.07s-.056.014-.08.014a9.047 9.047 0 0 1-1.167-.035l.037.003a1.747 1.747 0 0 1-1.066-.541l-.001-.001a3.617 3.617 0 0 1-.364-.527l-.01-.018c-.433-.647-.862-1.295-1.295-1.946a.907.907 0 0 0-.254-.262l-.003-.002c-.033-.03-.077-.048-.126-.048s-.093.018-.126.049a.719.719 0 0 0-.183.186l-.002.003c-.545.779-1.05 1.586-1.579 2.378a1.66 1.66 0 0 1-1.09.709l-.01.001a4.275 4.275 0 0 1-.925.052h.01zm-.704-7.981v.129q0 3.29.006 6.58l.001.039a1.2 1.2 0 0 1-.615 1.048l-.006.003a1.024 1.024 0 0 1-.445.155h-.005a4.65 4.65 0 0 1-.505.026h-.031h.002c-.109 0-.109 0-.109-.106c0-1.186 0-2.372-.006-3.558c0-.998-.01-1.992 0-2.99a.532.532 0 0 1 .269-.438l.003-.001c.146-.096.307-.16.459-.245c.304-.147.563-.301.807-.474l-.017.012c.056-.053.115-.112.194-.182zm-1.727.424v-.04c0-.282.036-.555.104-.815l-.005.022c.09-.326.276-.6.526-.798l.003-.002c.231-.161.51-.27.811-.303l.008-.001c.089-.017.178-.023.267-.036c.073-.014.093.026.096.089l.001.049c0 .119-.011.236-.032.35l.002-.012c-.066.403-.25.755-.516 1.028a1.482 1.482 0 0 1-.538.324l-.01.003a2.495 2.495 0 0 0-.615.325l.007-.005c-.026.02-.056.059-.093.04s-.017-.066-.02-.099s.003-.082.003-.118z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Wordpress(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 12v-.017c0-1.763.385-3.437 1.076-4.94l-.03.074l5.727 15.68a11.987 11.987 0 0 1-4.887-4.338l-.029-.048A11.543 11.543 0 0 1 0 12.098v-.103zm20.099-.609c0 .212-.015.421-.042.625l.003-.024a7.015 7.015 0 0 1-.171.822l.011-.049q-.117.468-.179.686t-.27.921q-.21.702-.273.905l-1.186 3.994l-4.338-12.889q.72-.047 1.374-.125a.458.458 0 0 0 .404-.286l.001-.003a.483.483 0 0 0-.04-.482l.001.002a.465.465 0 0 0-.447-.211h.002l-3.2.16q-1.17-.016-3.152-.16a.466.466 0 0 0-.321.081l.002-.001a.463.463 0 0 0-.178.231l-.001.003a.584.584 0 0 0-.023.292v-.004a.442.442 0 0 0 .14.257a.487.487 0 0 0 .303.125h.001l1.248.125l1.873 5.12l-2.622 7.865L4.65 6.38q.72-.047 1.374-.125a.458.458 0 0 0 .404-.286l.001-.003a.483.483 0 0 0-.04-.482l.001.002a.465.465 0 0 0-.447-.211h.002l-3.2.16q-.11 0-.359-.008t-.406-.008a12.04 12.04 0 0 1 4.222-3.925l.061-.031A11.857 11.857 0 0 1 11.998.004c1.575 0 3.08.304 4.458.855L16.375.83a12.092 12.092 0 0 1 3.73 2.333l-.009-.008H19.9c-.557 0-1.056.244-1.398.63l-.002.002c-.354.38-.571.892-.571 1.454v.03v-.002c0 .132.011.262.033.388l-.002-.014q.031.187.062.335c.033.137.076.257.13.37l-.005-.011q.094.21.141.328c.061.137.124.252.194.362l-.007-.011q.141.234.195.328t.226.374l.218.359c.595.94.958 2.079.982 3.301v.007zm-7.896 1.655l3.698 10.096c.012.066.04.125.08.173v-.001a11.74 11.74 0 0 1-3.978.686h-.028a11.54 11.54 0 0 1-3.441-.521l.082.022zm10.315-6.8A11.696 11.696 0 0 1 24 11.982v.025v-.001v.087c0 2.187-.603 4.233-1.653 5.981l.029-.053a12.088 12.088 0 0 1-4.297 4.315l-.056.031l3.667-10.586c.51-1.269.841-2.738.919-4.274l.001-.033v-.035a7.7 7.7 0 0 0-.1-1.242l.006.044z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func World(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 0h-.029c-6.627 0-12 5.373-12 12s5.373 12 12 12H12c6.627 0 12-5.373 12-12S18.627 0 12 0m0 23.04h-.029C5.874 23.04.931 18.097.931 12S5.874.96 11.971.96H12C18.097.96 23.04 5.903 23.04 12S18.097 23.04 12 23.04"/><path fill="currentColor" d="M15.825 7.142c-.99.25-2.136.412-3.314.45l-.026.001v3.929h3.9a19.875 19.875 0 0 0-.589-4.516zM12.479 1.99v4.725a15.337 15.337 0 0 0 3.209-.437l-.105.022c-.756-2.39-1.933-3.958-3.104-4.31m-4.09 4.304c.926.232 2 .384 3.102.42l.025.001V1.979c-1.171.34-2.366 1.91-3.127 4.315m10.799-1.368A10.039 10.039 0 0 0 14.417 2.2l-.069-.014a9.786 9.786 0 0 1 2.064 3.813l.016.069a15.153 15.153 0 0 0 2.84-1.182zm-1.921 6.596h4.806a10.003 10.003 0 0 0-2.293-5.937l.013.017c-.89.502-1.923.948-3.005 1.285l-.115.031c.353 1.375.568 2.958.593 4.588zm-9.681 0h3.934V7.593a16.207 16.207 0 0 1-3.481-.48l.112.024a18.935 18.935 0 0 0-.565 4.37zm8.799.957h-3.906v3.92c1.204.035 2.353.197 3.457.474l-.11-.023c.332-1.305.534-2.808.559-4.354v-.016zm-8.233 4.379c1-.255 2.156-.42 3.345-.455l.024-.001v-3.92H7.586c.023 1.564.228 3.069.594 4.51zm3.37 5.158v-4.737c-1.13.037-2.205.188-3.24.443l.107-.022c.762 2.406 1.956 3.979 3.133 4.316m2.826-.202a10.125 10.125 0 0 0 4.845-2.745l.002-.002a14.607 14.607 0 0 0-2.65-1.113l-.108-.029a9.867 9.867 0 0 1-2.098 3.9zM9.617 2.198a10.1 10.1 0 0 0-4.804 2.733l-.002.002c.775.424 1.678.808 2.622 1.102l.107.029a9.871 9.871 0 0 1 2.086-3.876l-.01.011zm7.056 14.885c1.199.365 2.235.812 3.208 1.357l-.081-.042a9.963 9.963 0 0 0 2.278-5.897l.001-.023h-4.806a21.08 21.08 0 0 1-.631 4.75zm-4.194.196v4.73c1.171-.352 2.348-1.92 3.104-4.315a14.76 14.76 0 0 0-3.081-.415l-.023-.001zM7.304 6.911A15.752 15.752 0 0 1 4.126 5.56l.079.041a9.963 9.963 0 0 0-2.278 5.897l-.001.023h4.782c.023-1.647.239-3.233.626-4.751zm-.601 5.568H1.921c.107 2.266.95 4.316 2.293 5.937l-.013-.017c.883-.499 1.91-.943 2.984-1.279l.114-.031a20.331 20.331 0 0 1-.595-4.591v-.019zm-1.897 6.588a10.056 10.056 0 0 0 4.742 2.72l.069.015A9.826 9.826 0 0 1 7.554 18l-.017-.07a15.486 15.486 0 0 0-2.814 1.178z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func WorldO(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24 12v-.006c0-3.551-1.546-6.74-4.001-8.933l-.012-.01a.596.596 0 0 0-.101-.087l-.002-.001A11.917 11.917 0 0 0 11.998.001c-3.032 0-5.8 1.126-7.91 2.984l.013-.011a.481.481 0 0 0-.07.065v.001A11.959 11.959 0 0 0 0 12.003c0 3.55 1.544 6.739 3.997 8.933l.012.01a.672.672 0 0 0 .106.097l.002.001a11.92 11.92 0 0 0 7.881 2.958a11.93 11.93 0 0 0 7.911-2.984l-.013.011a11.94 11.94 0 0 0 4.106-9.021v-.007zm-4.462 7.805a11.935 11.935 0 0 0-1.909-1.262l-.065-.032c.613-1.767.982-3.804 1.017-5.923v-.016h4.261a10.824 10.824 0 0 1-3.301 7.23zM12.572 18.3c1.283.069 2.482.351 3.588.81l-.072-.026c-.886 2.02-2.133 3.408-3.516 3.713zm0-1.144v-4.584h4.868a18.593 18.593 0 0 1-.976 5.578l.039-.131a11.637 11.637 0 0 0-3.903-.862l-.027-.001zm0-5.728V6.844a11.91 11.91 0 0 0 4.007-.891l-.079.029c.555 1.619.896 3.485.94 5.425v.021zm0-5.728V1.205c1.383.305 2.63 1.687 3.516 3.713c-1.034.43-2.233.711-3.487.781zm2.854-4a10.841 10.841 0 0 1 3.258 1.752l-.023-.018c-.443.348-.94.676-1.464.961l-.056.028a10.07 10.07 0 0 0-1.724-2.737l.009.011zm-4-.492V5.7a10.804 10.804 0 0 1-3.588-.81l.072.026c.89-2.02 2.135-3.407 3.518-3.712zM6.858 4.42a10.943 10.943 0 0 1-1.544-1.007l.024.018a10.744 10.744 0 0 1 3.158-1.712l.076-.023a10.02 10.02 0 0 0-1.689 2.658zm4.57 2.423v4.584H6.56c.044-1.961.385-3.827.979-5.577l-.039.131a11.65 11.65 0 0 0 3.901.861zm0 5.728v4.584a11.91 11.91 0 0 0-4.007.891l.079-.029c-.555-1.618-.896-3.485-.94-5.425v-.021zm0 5.728v4.495c-1.383-.305-2.63-1.687-3.516-3.713c1.034-.43 2.233-.71 3.487-.78zm-2.85 4a10.845 10.845 0 0 1-3.258-1.748l.024.018c.443-.348.94-.676 1.464-.961l.056-.028a9.843 9.843 0 0 0 1.723 2.733l-.009-.01zm8.564-2.72c.58.315 1.077.642 1.544 1.007l-.024-.018a10.744 10.744 0 0 1-3.158 1.712l-.076.023a10.014 10.014 0 0 0 1.689-2.657l.025-.065zm5.7-8.151h-4.261a19.663 19.663 0 0 0-1.058-6.078l.041.138a12.383 12.383 0 0 0 1.997-1.312l-.024.018a10.813 10.813 0 0 1 3.303 7.205zM4.462 4.195c.576.468 1.223.897 1.909 1.262l.065.032c-.613 1.767-.982 3.804-1.017 5.923v.016H1.157a10.824 10.824 0 0 1 3.301-7.23zm-3.304 8.377h4.261a19.663 19.663 0 0 0 1.058 6.078l-.041-.138c-.751.399-1.397.828-1.997 1.312l.024-.018a10.813 10.813 0 0 1-3.303-7.205l-.001-.028z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Xbox(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M24 12v-.002c0-3.618-1.606-6.861-4.144-9.054l-.015-.013a20.22 20.22 0 0 0-4.967 3.713l-.004.004c.044.046.087.085.131.132c3.719 4.012 7.106 9.73 6.546 12.471a11.85 11.85 0 0 0 2.452-7.246v-.006zM12.591 3.955a12.722 12.722 0 0 1 5.872-2.022l.048-.003A11.834 11.834 0 0 0 12 .001c-2.171 0-4.207.579-5.962 1.591l.058-.031c.658.567 2.837.781 5.484 2.4a.954.954 0 0 0 1.015-.007zM9.166 6.778c.046-.049.093-.09.138-.138a28.609 28.609 0 0 0-3.806-3.1l-.099-.064a1.845 1.845 0 0 0-1.516-.305l.013-.002A11.953 11.953 0 0 0 0 12.009c0 2.909 1.037 5.576 2.762 7.651l-.016-.02c-1.031-2.547 2.477-8.672 6.419-12.862zm2.918 2.42c-3.962 3.503-9.477 8.73-8.632 11.218A11.936 11.936 0 0 0 11.994 24a11.94 11.94 0 0 0 8.826-3.883l.008-.009c.486-2.618-4.755-7.337-8.744-10.91"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Yacht(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M0 19.216L5.414 24h15.882l2.274-3.195L.002 19.216zM13.985 0v19.073l7.309.503zm-.996 3.214L2.429 18.241l10.56.713z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Yahoo(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.818 13.803L11.005 24a8.65 8.65 0 0 0-1.505-.16h-.009a9.34 9.34 0 0 0-1.573.17l.059-.01l.187-10.197q-.577-.995-2.43-4.262t-3.12-5.402Q1.343 2.005.001 0a6.1 6.1 0 0 0 1.554.216h.004a7.6 7.6 0 0 0 1.653-.227L3.159 0q.909 1.6 1.926 3.31t2.409 3.988l2 3.274q.534-.88 1.579-2.56t1.694-2.74t1.514-2.538T15.823 0c.463.127.994.2 1.542.202h.001a6.977 6.977 0 0 0 1.692-.212L19.01 0q-.4.56-.866 1.28t-.714 1.132q-.252.418-.815 1.385t-.706 1.211q-2.106 3.574-5.091 8.795"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Yandex(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M7.083 14.8L2.985 24H-.001l4.5-9.834C2.385 13.092.974 11.148.974 7.552C.969 2.518 4.159.001 7.953.001h3.858v24H9.229v-9.2H7.083zM9.23 2.18H7.852c-2.08 0-4.097 1.378-4.097 5.372c0 3.858 1.847 5.1 4.097 5.1H9.23z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YandexInternational(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M5.2 24v-7.786L0 2.25h2.616L6.45 13.017L10.86-.001h2.405L7.607 16.302v7.697z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Yarn(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 24h-.007c-6.627 0-12-5.373-12-12c0-4.943 2.989-9.189 7.259-11.027l.078-.03A11.804 11.804 0 0 1 12-.003c6.629 0 12.003 5.374 12.003 12.003c0 4.942-2.987 9.187-7.254 11.027l-.078.03a11.728 11.728 0 0 1-4.647.943zm-3.28-4.601a1.912 1.912 0 0 0 1.457.49h-.007h.046a31.334 31.334 0 0 0 3.637-.337l-.172.022c.271-.061.508-.181.707-.345l-.003.002a13.008 13.008 0 0 0 3.065-1.423l-.053.032a4.789 4.789 0 0 1 1.788-.837l.033-.006a1.232 1.232 0 0 0 .945-1.367l.001.006a1.282 1.282 0 0 0-1.268-1.105l-.067.002h.003h-.021a5.045 5.045 0 0 0-2.568.941l.015-.01a5.192 5.192 0 0 1-.569.314l-.034.015a4.902 4.902 0 0 0-.285-2.05l.011.034a5.02 5.02 0 0 0-1.122-1.826l.002.002a6.482 6.482 0 0 0 1.123-2.535l.007-.042a6.41 6.41 0 0 0-.349-3.648l.015.043a.74.74 0 0 0-.458-.374l-.005-.001a.748.748 0 0 0-.217-.032h-.002a1.213 1.213 0 0 0-.338.054l.009-.002c-.464-.96-.63-1.071-.752-1.152a.96.96 0 0 0-.522-.16h-.001a.942.942 0 0 0-.337.062l.006-.002c-.48.241-.844.653-1.02 1.153l-.004.014a5.076 5.076 0 0 0-.125.283a3.4 3.4 0 0 0-2.298 1.103l-.002.003a1.33 1.33 0 0 1-.477.258l-.009.002h.005c-.4.142-.592.474-.82 1.079a2.572 2.572 0 0 0 .335 2.22l-.006-.009a5.168 5.168 0 0 0-1.307 1.706l-.013.031a5.446 5.446 0 0 0-.407 2.265v-.009a3.043 3.043 0 0 0-.878 1.692l-.002.017a3.162 3.162 0 0 0 .462 1.961l-.008-.014c.047.073.097.137.153.195a1.158 1.158 0 0 0 .006.408l-.001-.007c.068.333.281.605.568.752l.006.003c.367.198.804.316 1.267.32h.032c.299 0 .582-.069.834-.193l-.011.005zm1.466-.209c-.555 0-.934-.146-1.04-.4a.812.812 0 0 1 .439-1.064l.005-.002l-.01-.006a1.829 1.829 0 0 1-.256-.19l.002.002a1.002 1.002 0 0 1-.087-.106l-.002-.003c-.042-.056-.08-.104-.102-.104c-.006 0-.012.005-.017.014a2.66 2.66 0 0 0-.098.339l-.003.018c-.053.353-.19.667-.389.932l.004-.005a1.044 1.044 0 0 1-.708.277l-.058-.002h.003a2.133 2.133 0 0 1-.957-.246l.012.006a.386.386 0 0 1-.216-.268v-.003a.995.995 0 0 1 .254-.661l-.001.001a.366.366 0 0 1-.16.038a.41.41 0 0 1-.343-.214l-.001-.002a2.425 2.425 0 0 1-.342-1.526l-.001.011c.14-.599.458-1.108.895-1.484l.004-.003a4.857 4.857 0 0 1 .35-2.293l-.012.032c.4-.757.946-1.386 1.607-1.869l.015-.011c-.102-.117-.976-1.158-.626-2.094c.226-.61.328-.64.4-.667l.014-.005c.292-.091.542-.242.747-.441h-.001a2.714 2.714 0 0 1 2.008-.885h.029c.083 0 .164.006.244.016l-.009-.001c.053-.155.53-1.52 1.01-1.52a.275.275 0 0 1 .15.047l-.001-.001c.28.424.544.911.764 1.421l.024.061h.006a2.11 2.11 0 0 1 .638-.266l.014-.003c.038 0 .062.01.074.032c.202.448.327.969.345 1.518v.007a6.64 6.64 0 0 1-.086 1.662l.006-.04a6.148 6.148 0 0 1-1.112 2.443l.011-.015c-.092.134-.172.249-.233.352c-.021.034.041.09.152.193a4.194 4.194 0 0 1 1.182 1.775l.009.029a4.585 4.585 0 0 1 .155 2.391l.005-.029a.795.795 0 0 0-.022.32v-.004l.018.033h.022a3.6 3.6 0 0 0 1.729-.679l-.01.007a4.328 4.328 0 0 1 2.179-.824l.017-.001h.018a.58.58 0 0 1 .64.489v.003a.538.538 0 0 1-.411.591l-.004.001a5.377 5.377 0 0 0-2.048.946l.014-.01a12.06 12.06 0 0 1-2.91 1.337l-.087.023a.844.844 0 0 1-.5.282l-.005.001a30.06 30.06 0 0 1-3.254.294l-.059.002z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Yelp(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M9.206 17.662v1.7q-.014 3.908-.08 4.08a.853.853 0 0 1-.678.535h-.004l-.102.001a5.553 5.553 0 0 1-2.362-.524l.034.014a5.534 5.534 0 0 1-2.181-1.196l.005.005a1.002 1.002 0 0 1-.226-.474l-.001-.006a.773.773 0 0 1 .055-.353l-.002.006a3.85 3.85 0 0 1 .459-.633l-.003.004q.4-.495 2.422-2.891q.014 0 .8-.937a.907.907 0 0 1 .523-.327l.006-.001a1.173 1.173 0 0 1 .671.05l-.008-.003c.205.082.375.216.499.385l.002.003a.873.873 0 0 1 .171.522l-.001.043zm-1.995-3.36v.03a.939.939 0 0 1-.69.905l-.007.002l-1.606.522Q1.228 16.939 1 16.939a.854.854 0 0 1-.718-.475l-.002-.005a3.079 3.079 0 0 1-.227-.995v-.009a11.12 11.12 0 0 1 .018-2.279l-.005.05a4.249 4.249 0 0 1 .411-1.692l-.011.026a.792.792 0 0 1 .752-.428h-.002c1.109.356 2.004.706 2.875 1.1l-.171-.069l1.539.629l1.124.455a.931.931 0 0 1 .473.404l.002.004a1.161 1.161 0 0 1 .152.652v-.003zm11.056 3.976a5.362 5.362 0 0 1-1.228 2.158l.003-.003a5.512 5.512 0 0 1-1.788 1.686l-.026.014a.805.805 0 0 1-.845-.095l.002.001q-.187-.134-2.462-3.84l-.629-1.03a.903.903 0 0 1-.153-.62v.004c.025-.237.12-.448.262-.618l-.001.002a.93.93 0 0 1 1.117-.346l-.006-.002q.014.014 1.593.535q2.72.88 3.239 1.064c.239.068.448.162.639.281l-.011-.006a.808.808 0 0 1 .292.822l.001-.005zM9.273 9.819q.067 1.366-.72 1.633q-.776.227-1.526-.95l-5.059-8a.843.843 0 0 1 .254-.829l.001-.001A6.882 6.882 0 0 1 4.957.479l.044-.007A7.143 7.143 0 0 1 8.026.048L8.002.047a.82.82 0 0 1 .655.597l.001.006q.04.24.294 4.089t.322 5.082zm8.859 1.446a.813.813 0 0 1-.346.788l-.003.002q-.201.134-4.4 1.151q-.897.201-1.218.308l.014-.026a.92.92 0 0 1-.622-.056l.006.002a1.108 1.108 0 0 1-.493-.423l-.003-.005a.94.94 0 0 1 .002-1.166l-.002.002q.014-.014 1.004-1.366q1.673-2.289 2.008-2.73a2.88 2.88 0 0 1 .453-.519l.003-.002a.834.834 0 0 1 .875-.024l-.004-.002a5.426 5.426 0 0 1 1.633 1.761l.014.026c.5.629.875 1.38 1.076 2.2l.008.04z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Yen(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M10.278 24H7.34a.537.537 0 0 1-.537-.528v-.019v.001v-5.625H1.888a.537.537 0 0 1-.537-.528v-.019v.001v-1.756a.553.553 0 0 1 .545-.546H6.8v-1.449H1.885a.537.537 0 0 1-.537-.528v-.019v.001v-1.772a.553.553 0 0 1 .545-.546h3.648L.069.816A.572.572 0 0 1 .071.267L.069.27a.535.535 0 0 1 .466-.273H.55H.549h3.316c.213 0 .397.124.484.304l.001.003l3.665 7.244q.32.648.954 2.13q.17-.409.52-1.159t.469-1.04L13.214.32A.495.495 0 0 1 13.71 0h-.001h3.256a.552.552 0 0 1 .478.804l.001-.003l-5.335 9.87h3.665c.3.004.542.246.546.545v1.791a.537.537 0 0 1-.537.528h-4.96v1.449h4.943c.3.004.542.246.546.545v1.774a.537.537 0 0 1-.537.528h-4.952v5.625a.553.553 0 0 1-.545.546z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func YoutubePlay(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m13.544 16.419l9.216-4.762l-9.216-4.818zM17.068 0q3.2 0 6.181.086t4.371.181l1.39.08q.019 0 .32.029c.167.013.317.033.465.061l-.026-.004q.134.029.448.086c.205.035.388.088.561.158L30.76.671q.229.095.534.248c.225.114.419.238.599.377l-.008-.006c.203.158.383.323.549.502l.003.003c.102.107.199.222.288.342l.007.01c.202.318.388.686.537 1.071l.015.044c.24.559.417 1.207.501 1.885l.004.035q.152 1.219.24 2.6t.105 2.16v3.352l.001.263c0 1.853-.125 3.678-.367 5.465l.023-.208a7.612 7.612 0 0 1-.498 1.946l.018-.051a4.844 4.844 0 0 1-.62 1.186l.01-.015l-.266.32a4.536 4.536 0 0 1-.542.497l-.01.008a3.127 3.127 0 0 1-.571.353l-.019.008q-.305.143-.534.24a2.58 2.58 0 0 1-.527.15l-.016.002q-.314.057-.457.086t-.438.057t-.314.029q-4.775.372-11.937.372q-3.943-.038-6.848-.124t-3.819-.143l-.934-.08l-.686-.08a8.883 8.883 0 0 1-1.102-.204l.062.013a5.832 5.832 0 0 1-1.003-.415l.032.015a3.837 3.837 0 0 1-1.075-.78l-.001-.001a3.641 3.641 0 0 1-.288-.342l-.007-.01a6.192 6.192 0 0 1-.537-1.071l-.015-.044a7.021 7.021 0 0 1-.501-1.885l-.004-.035q-.152-1.219-.24-2.6t-.105-2.16v-3.352l-.001-.263c0-1.853.125-3.678.367-5.465l-.023.208A7.612 7.612 0 0 1 .84 3.238l-.018.051c.162-.447.368-.834.62-1.186l-.01.015l.266-.32c.169-.182.349-.347.542-.497l.01-.008c.172-.134.365-.257.57-.362l.021-.01q.305-.152.534-.248a2.58 2.58 0 0 1 .527-.15l.016-.002q.314-.057.448-.086c.121-.024.271-.044.424-.056l.015-.001q.305-.029.32-.029q4.783-.35 11.946-.35z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZipperMouth(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M12 0C5.342 0 0 5.342 0 12c0 6.581 5.419 12 12 12s12-5.419 12-12c0-6.658-5.342-12-12-12m0 1.703c5.729 0 10.297 4.645 10.297 10.297S17.652 22.297 12 22.297S1.703 17.652 1.703 12S6.348 1.703 12 1.703"/><path fill="currentColor" d="M7.587 12c1.161 0 2.09-.929 2.09-2.09s-.929-2.09-2.09-2.09s-2.09.929-2.09 2.09c0 1.161.929 2.09 2.09 2.09m8.903 0c1.161 0 2.09-.929 2.09-2.09s-.929-2.09-2.09-2.09s-2.09.929-2.09 2.09c0 1.161.929 2.09 2.09 2.09m-7.2 3.329l1.626 1.626l-1.626 1.626c-.31.31-.31.852 0 1.161s.852.31 1.161 0l1.626-1.626l1.626 1.626c.155.155.387.232.619.232s.465-.077.697-.155a.84.84 0 0 0 0-1.161l-1.626-1.626l1.626-1.626a.835.835 0 0 0 .232-.542c0-.465-.387-.852-.852-.929c-.232 0-.465.077-.619.232l-1.626 1.626l-1.626-1.626c-.155-.155-.387-.232-.619-.232s-.465.077-.619.232a.84.84 0 0 0 0 1.161z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func Zoom(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="m22.241 24l-7.414-7.414a9.482 9.482 0 0 1-5.652 1.885h-.002l-.108.001A9.065 9.065 0 0 1 0 9.407l.001-.114v.006a9.298 9.298 0 0 1 18.596 0a9.8 9.8 0 0 1-1.904 5.682l.019-.027l7.414 7.414zM9.299 2.513a6.758 6.758 0 1 0 .017.001z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomMinus(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.88 16.56a10.01 10.01 0 0 1-5.667 1.92H9.2l-.152.001A9.048 9.048 0 0 1 0 9.524v-.245a9.28 9.28 0 0 1 18.56 0a9.858 9.858 0 0 1-1.939 5.707l.019-.027l7.44 7.44l-1.76 1.6zM2.64 9.28a6.761 6.761 0 0 0 13.52-.04v-.055a6.513 6.513 0 0 0-1.998-4.703l-.002-.002A6.732 6.732 0 0 0 2.64 9.212v.072zm2.96.8v-1.6h7.44v1.6z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}

func ZoomPlus(children ...ElementRenderer) *FontistoIcon {
	return &FontistoIcon{
		SVGSVGElement: SVG_SVG(
			Text(`<path fill="currentColor" d="M14.88 16.56a10.01 10.01 0 0 1-5.667 1.92H9.2l-.152.001A9.048 9.048 0 0 1 0 9.524v-.245a9.28 9.28 0 0 1 18.56 0a9.858 9.858 0 0 1-1.939 5.707l.019-.027l7.44 7.44l-1.76 1.6zM2.64 9.28a6.761 6.761 0 0 0 13.52-.04v-.055a6.513 6.513 0 0 0-1.998-4.703l-.002-.002A6.732 6.732 0 0 0 2.64 9.212v.072zm5.92 3.733V10.08H5.6v-1.6h2.96V5.574h1.6V8.48h2.88v1.6h-2.88v2.934z"/>`),
		).HEIGHT(hAttr).VIEW_BOX(viewbox).IfChildren(len(children) > 0, children...),
	}
}
